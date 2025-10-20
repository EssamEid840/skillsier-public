# Users-BE “Fix Pack” — Integrity, Security, and Scale

Huge improvement — this is genuinely production‑grade. You covered the domains, added outbox/CQRS, audit, KYC, moderation, capacity, and a ton of pragmatic fields. I’d ship this with a small “fix pack” to tighten integrity, security, and scale.

## High‑priority integrity fixes

### 1) Arrays that should be relations
Anywhere you use `UUID[]` or `TEXT[]` to point at other rows, you lose referential integrity and query power. Make junction tables (and drop the arrays):

- `service_catalog.required_skill_ids` → **service_required_skills**`(service_id, skill_id UNIQUE)`  
- `specializations.primary_skill_ids`, `secondary_skill_ids` → **specialization_skills**`(specialization_id, skill_id, kind CHECK IN ('PRIMARY','SECONDARY'))`  
- `experience.skill_ids`, `education.skill_ids`, `portfolios.skill_ids` → **experience_skills**, **education_skills**, **portfolio_skills**  
- `moderation_actions.related_report_ids` → **moderation_action_reports**`(action_id, report_id UNIQUE)`  

> And anywhere else you have `*_ids` in an array.

### 2) Missing `ON DELETE` behavior
Add explicit actions on every FK that points to `users` or other core rows (you already did many). E.g., endpoints like `endorsements.skill_id` (CASCADE?), `connections.requested_by` (SET NULL?), etc. Be deliberate: relationships, reports, sessions, outbox, logs.

### 3) Uniqueness semantics
A few should be unique but aren’t:

- `service_catalog`: `UNIQUE (user_id, service_slug)` (global slug collisions across all users are unlikely what you want).  
- `profiles.custom_url_slug` is unique globally — that’s fine if it’s a public vanity slug; otherwise consider `(user_id, custom_url_slug)` + unique global in a separate “vanity routes” table.

### 4) Time‑overlap constraints
For vacation windows, rate history, mentorship sessions, etc., prevent overlaps per user with exclusion constraints (needs **btree_gist** instead of **btree_gin**):

- `availability_vacations`:  
  `EXCLUDE USING gist (user_id WITH =, daterange(start_date, end_date, '[]') WITH &&)`
- `rate_history`: same pattern on `(user_id, daterange(effective_at, effective_until))`

---

## Security & PII

### Token hygiene
Don’t store verification/reset tokens raw. Store **hashed tokens** + created/expiry in dedicated tables with TTL indexes; keep only last N active. E.g., `email_verification_tokens`, `password_reset_tokens` with `(user_id, token_hash, purpose, expires_at)`. Then drop `users.password_reset_token`, `email_verification_token`, `phone_verification_code`.

### Emails & case‑insensitivity
Use **citext** for `users.email` (and any email in other tables) to avoid case traps. Add pg_trgm GIN on `citext(email)` for fuzzy lookup if needed.

### RLS on PII tables
Turn on **Row‑Level Security** and add policies on `users`, `profiles`, `payment_methods`, `tax_profiles`, `identity_verifications`, `sessions`, `devices`, etc. Admin/service roles get broader policies; users can only see self.

### JSONB shape checks
For fields that gate logic (e.g., `notification_settings.*_json`, `communication_channels.channels_json`, `outbox_events.payload`), add `CHECK (jsonb_typeof(...) = 'object')` and consider JSON Schema validation in app.

---

## Performance/scale

### Partition hot, append‑only tables
Monthly partition by `occurred_at/created_at` for:

- `audit_logs`, `security_events`, `outbox_events`, `data_access_logs`, possibly `user_reports`, `moderation_actions`.

Add local indexes only on needed columns; keep parent with no data.

### Search vector as generated column
Replace the trigger with a stored **generated** column, simpler and faster to keep in sync:

```sql
ALTER TABLE user_read_model
ADD COLUMN search_vector tsvector
GENERATED ALWAYS AS (
  setweight(to_tsvector('english', coalesce(full_name,'')), 'A') ||
  setweight(to_tsvector('english', coalesce(professional_title,'')), 'B') ||
  setweight(to_tsvector('english', coalesce(tagline,'')), 'C')
) STORED;
CREATE INDEX idx_user_read_search ON user_read_model USING gin (search_vector);
DROP TRIGGER IF EXISTS update_search_vector_trigger ON user_read_model;
```

If you have non-English users, switch to `to_tsvector('simple', …)` or multi‑column per language, and pick the config at query time.

### Prune redundant indexes
You’ve got healthy coverage; just watch for overlap like `(status)` + `(status, expires_at)` on several tables — keep the composite and drop the single if reads need both.

### Partial indexes where filters are common
You already use them; consider adding:

- `sessions (user_id) WHERE is_active`
- `moderation_actions (status) WHERE status IN ('ACTIVE','UNDER_APPEAL')`
- `availability (status) WHERE status IN ('AVAILABLE','PARTIALLY_AVAILABLE')` (you have this — good)

### Foreign keys on big, write‑heavy paths
Where event streams write fast (e.g., `outbox_events.actor_user_id`), consider **deferrable FKs** or documenting that some columns are **soft references** (no FK) to avoid lock amplification — but do it intentionally.

---

## Consistency & modeling

- **Enum types over VARCHAR checks.** For fields like `user_type`, `account_status`, etc., Postgres enums give smaller storage and better constraints. If you prefer flexibility, keep VARCHAR + CHECK, but define central domain types to reuse.
- **Naming & domain alignment.** You asked for `{domain}_{sub}` for sub‑entities. A few tables are generic names in public (e.g., `preferences`, `connections`, `availability`) — they do map to their domains in comments, but the name doesn’t reflect it. If strict alignment matters, rename to:
  - `profile_preferences` (or move entirely to `user_settings` + `privacy_settings`, which you added)
  - `network_connections`, `network_endorsements`, etc.
- You already introduced domain‑specific `user_settings`, `privacy_settings`; consider migrating off **preferences** (it overlaps).
- **Duplication cleanup.** You noted this in the “RESOLUTION” section but didn’t apply it:
  - If `two_factor_auth` is source of truth, remove `users.two_factor_*` (or keep a single `two_factor_enabled` cache + trigger to sync).
  - If **freelancers** owns rates, drop `profiles.hourly_rate*` (or keep a denormalized copy with a materialized view/trigger).
  - Prefer one of **user_statistics** vs **user_metrics** (you chose metrics; drop statistics).
- **Slug & URL hygiene.** Add `CHECK (service_slug ~* '^[a-z0-9]+(?:-[a-z0-9]+)*$')` (and similar for other slugs) to keep routing clean.

---

## Concrete DDL patch (drop‑in)

Here are representative snippets you can apply now; replicate for the other array/FK spots:

```sql
-- 0) Prereqs
CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS btree_gist;

-- 1) emails as citext
ALTER TABLE users ALTER COLUMN email TYPE citext;
CREATE UNIQUE INDEX IF NOT EXISTS uq_users_email ON users ((lower(email))) WHERE is_deleted = FALSE;

-- 2) unique per-user service slug
ALTER TABLE service_catalog ADD CONSTRAINT uk_service_slug UNIQUE (user_id, service_slug);

-- 3) service skills junction
CREATE TABLE service_required_skills (
  service_id UUID NOT NULL REFERENCES service_catalog(id) ON DELETE CASCADE,
  skill_id   UUID NOT NULL REFERENCES skills_taxonomy(id),
  PRIMARY KEY (service_id, skill_id)
);
-- migrate existing data (pseudo):
-- INSERT INTO service_required_skills SELECT id, unnest(required_skill_ids) FROM service_catalog WHERE required_skill_ids IS NOT NULL;
ALTER TABLE service_catalog DROP COLUMN IF EXISTS required_skill_ids;

-- 4) vacations: prevent overlaps
ALTER TABLE availability_vacations
  ADD COLUMN period daterange GENERATED ALWAYS AS (daterange(start_date, end_date, '[]')) STORED;
CREATE INDEX ON availability_vacations USING gist (user_id, period);
ALTER TABLE availability_vacations
  ADD CONSTRAINT no_overlapping_vacations EXCLUDE USING gist (user_id WITH =, period WITH &&);

-- 5) sessions active partial index (fast lookups)
CREATE INDEX IF NOT EXISTS idx_sessions_user_active ON sessions (user_id) WHERE is_active = TRUE;

-- 6) RLS example (apply to other PII tables similarly)
ALTER TABLE users ENABLE ROW LEVEL SECURITY;
CREATE POLICY users_self ON users
  USING (id = current_setting('app.current_user_id')::uuid)
  WITH CHECK (id = current_setting('app.current_user_id')::uuid);

-- 7) Replace trigger VECTOR with generated column (if you keep english)
ALTER TABLE user_read_model DROP COLUMN IF EXISTS search_vector;
ALTER TABLE user_read_model
ADD COLUMN search_vector tsvector GENERATED ALWAYS AS (
  setweight(to_tsvector('english', coalesce(full_name,'')), 'A') ||
  setweight(to_tsvector('english', coalesce(professional_title,'')), 'B') ||
  setweight(to_tsvector('english', coalesce(tagline,'')), 'C')
) STORED;
CREATE INDEX idx_user_read_search ON user_read_model USING gin (search_vector);

-- 8) token tables (sketch)
CREATE TABLE email_verification_tokens (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  token_hash bytea NOT NULL,
  purpose VARCHAR(20) NOT NULL CHECK (purpose IN ('EMAIL_VERIFY')),
  created_at timestamptz NOT NULL DEFAULT now(),
  expires_at timestamptz NOT NULL
);
CREATE INDEX ON email_verification_tokens (user_id, expires_at);
```

---

## Observability & ops

- Add `pg_stat_statements` (if not already) and a slow query log baseline in staging.
- Add autovacuum scale settings for the largest, churny tables (events, logs).
- Materialize a couple of read models you’ll hit often (`v_active_freelancers`, `v_user_network_summary`) if they become heavy joins; refresh on a schedule or via outbox consumers.

---

## Migration plan (pragmatic)

1. **Deploy** new tables & constraints additively.  
2. **Backfill** junction tables from arrays.  
3. **Switch reads** to new relations.  
4. **Drop** array columns.  
5. **Introduce RLS** with `app.current_user_id` set by the app (start permissive in admin).  
6. **Partition** logs/events and move old data.  
7. **Remove duplicated columns** (`two_factor_*`, `profiles.hourly_rate*`) or wire sync triggers if you want an interim denorm.


--------------------------------------------------------------------------------------------------------------------------------
--------------------------------------------------------------------------------------------------------------------------------
++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++



# Users-BE Database Design — Updated Sections Only

This document contains **only the modified/new sections** you should include directly in your initial schema. No ALTER/DROP needed.

---

## Extensions (updated)

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";
CREATE EXTENSION IF NOT EXISTS "btree_gin";
CREATE EXTENSION IF NOT EXISTS "citext";
CREATE EXTENSION IF NOT EXISTS "btree_gist";
```
---

## Section 1: Core User Domain (updated `users`)

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    -- Keycloak
    keycloak_id UUID NOT NULL UNIQUE,
    keycloak_created_at TIMESTAMPTZ,

    -- Identity (PII - email is CITEXT for case-insensitive unique)
    email CITEXT NOT NULL UNIQUE,
    email_verified BOOLEAN DEFAULT FALSE,
    phone VARCHAR(50),
    phone_verified BOOLEAN DEFAULT FALSE,

    -- Personal Info
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100),
    display_name VARCHAR(200),

    -- User Type & Status
    user_type VARCHAR(20) NOT NULL CHECK (user_type IN ('CLIENT','FREELANCER','HYBRID','ADMIN','AGENCY')),
    account_status VARCHAR(20) DEFAULT 'PENDING_VERIFICATION' CHECK (
        account_status IN ('PENDING_VERIFICATION','ACTIVE','SUSPENDED','BANNED','DEACTIVATED','DELETED')
    ),
    suspension_reason TEXT,
    suspension_expires_at TIMESTAMPTZ,
    banned_reason TEXT,
    banned_at TIMESTAMPTZ,
    banned_by UUID,

    -- Profile Completion
    profile_completion_score INTEGER DEFAULT 0 CHECK (profile_completion_score BETWEEN 0 AND 100),
    onboarding_completed BOOLEAN DEFAULT FALSE,
    onboarding_step INTEGER DEFAULT 0,
    onboarding_data JSONB,

    -- Security & Privacy (summary flags only; see dedicated tables for secrets/tokens)
    two_factor_enabled BOOLEAN DEFAULT FALSE,
    two_factor_method VARCHAR(20),
    two_factor_secret VARCHAR(255),
    security_questions_set BOOLEAN DEFAULT FALSE,
    last_password_change TIMESTAMPTZ,

    -- Activity
    last_login_at TIMESTAMPTZ,
    last_login_ip INET,
    last_login_device VARCHAR(255),
    last_login_user_agent TEXT,
    last_active_at TIMESTAMPTZ,
    login_count INTEGER DEFAULT 0,
    failed_login_attempts INTEGER DEFAULT 0,
    account_locked_until TIMESTAMPTZ,
    last_failed_login_at TIMESTAMPTZ,

    -- Platform Metadata
    registration_source VARCHAR(50),
    registration_ip INET,
    registration_user_agent TEXT,
    referral_code VARCHAR(50) UNIQUE,
    referred_by UUID,
    referral_count INTEGER DEFAULT 0,

    -- Feature Flags
    beta_features_enabled BOOLEAN DEFAULT FALSE,
    feature_flags JSONB,

    -- Soft Delete & Audit
    is_deleted BOOLEAN DEFAULT FALSE,
    deleted_at TIMESTAMPTZ,
    deleted_by UUID,
    deletion_reason TEXT,

    -- GDPR/Compliance
    data_processing_consent BOOLEAN DEFAULT FALSE,
    data_processing_consent_date TIMESTAMPTZ,
    marketing_consent BOOLEAN DEFAULT FALSE,
    marketing_consent_date TIMESTAMPTZ,

    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    version INTEGER DEFAULT 1 NOT NULL,

    CONSTRAINT fk_users_referred_by FOREIGN KEY (referred_by) REFERENCES users(id),
    CONSTRAINT fk_users_banned_by FOREIGN KEY (banned_by) REFERENCES users(id),
    CONSTRAINT fk_users_deleted_by FOREIGN KEY (deleted_by) REFERENCES users(id)
);

CREATE INDEX idx_users_email ON users (email) WHERE is_deleted = FALSE;
CREATE INDEX idx_users_keycloak_id ON users (keycloak_id);
CREATE INDEX idx_users_user_type ON users (user_type) WHERE is_deleted = FALSE;
CREATE INDEX idx_users_account_status ON users (account_status) WHERE is_deleted = FALSE;
CREATE INDEX idx_users_created_at ON users (created_at DESC);
CREATE INDEX idx_users_last_active ON users (last_active_at DESC) WHERE is_deleted = FALSE;
CREATE INDEX idx_users_referral_code ON users (referral_code) WHERE referral_code IS NOT NULL;
CREATE INDEX idx_users_phone ON users (phone) WHERE phone IS NOT NULL AND is_deleted = FALSE;
```
---

## Section 4: Service Catalog (updated) + junction

```sql
CREATE TABLE service_catalog (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,

    service_name VARCHAR(200) NOT NULL,
    service_slug VARCHAR(200) NOT NULL, -- unique per user
    short_description VARCHAR(500),
    full_description TEXT,

    -- Capabilities Reference (moved to junction table)
    specialization_id UUID,

    -- Pricing
    base_price DECIMAL(10, 2) NOT NULL,
    currency CHAR(3) DEFAULT 'USD',
    pricing_model VARCHAR(20) CHECK (pricing_model IN ('FIXED','HOURLY','CUSTOM')),

    -- Scope
    estimated_delivery_days INTEGER,
    revisions_included INTEGER DEFAULT 0,
    express_delivery_available BOOLEAN DEFAULT FALSE,
    express_delivery_days INTEGER,
    express_delivery_fee DECIMAL(10, 2),

    -- Service Packages
    has_packages BOOLEAN DEFAULT FALSE,

    -- Media
    cover_image_url TEXT,
    gallery_urls TEXT[],
    video_url TEXT,

    -- Requirements & Deliverables
    client_requirements TEXT[],
    deliverables TEXT[],

    -- Status & Visibility
    is_active BOOLEAN DEFAULT TRUE,
    is_featured BOOLEAN DEFAULT FALSE,
    approval_status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        approval_status IN ('PENDING','APPROVED','REJECTED','SUSPENDED')
    ),

    -- Metrics
    orders_count INTEGER DEFAULT 0,
    average_rating DECIMAL(3, 2),
    views_count INTEGER DEFAULT 0,
    favorite_count INTEGER DEFAULT 0,

    -- SEO
    seo_keywords TEXT[],
    seo_title VARCHAR(200),
    seo_description TEXT,

    -- FAQ
    faq JSONB,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    last_ordered_at TIMESTAMPTZ,

    CONSTRAINT fk_service_catalog_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_service_catalog_specialization FOREIGN KEY (specialization_id) REFERENCES specializations(id),
    CONSTRAINT uk_service_slug_per_user UNIQUE (user_id, service_slug),
    CONSTRAINT chk_service_slug_format CHECK (service_slug ~ '^[a-z0-9]+(?:-[a-z0-9]+)*$')
);

CREATE INDEX idx_service_catalog_user ON service_catalog (user_id);
CREATE INDEX idx_service_catalog_active ON service_catalog (is_active) WHERE is_active = TRUE;
CREATE INDEX idx_service_catalog_featured ON service_catalog (is_featured) WHERE is_featured = TRUE;
CREATE INDEX idx_service_catalog_slug ON service_catalog (service_slug);
CREATE INDEX idx_service_catalog_approval ON service_catalog (approval_status);

-- required skills junction
CREATE TABLE service_required_skills (
  service_id UUID NOT NULL REFERENCES service_catalog(id) ON DELETE CASCADE,
  skill_id   UUID NOT NULL REFERENCES skills_taxonomy(id),
  PRIMARY KEY (service_id, skill_id)
);
```
---

## Sections 3/5/6/9: Arrays → Junctions (updated)

### Specializations (updated) + `specialization_skills`

```sql
CREATE TABLE specializations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,

    title VARCHAR(200) NOT NULL,
    description TEXT,

    niche_expertise TEXT,
    industries TEXT[],
    target_clients TEXT[],
    unique_value_proposition TEXT,

    is_verified BOOLEAN DEFAULT FALSE,
    verified_at TIMESTAMPTZ,
    evidence_urls TEXT[],
    verification_notes TEXT,

    projects_completed INTEGER DEFAULT 0,
    client_satisfaction DECIMAL(3, 2),
    success_rate DECIMAL(5, 2),

    is_featured BOOLEAN DEFAULT FALSE,
    display_order INTEGER DEFAULT 0,
    show_on_profile BOOLEAN DEFAULT TRUE,

    seo_keywords TEXT[],

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_specializations_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_specializations_user ON specializations (user_id);
CREATE INDEX idx_specializations_verified ON specializations (is_verified) WHERE is_verified = TRUE;
CREATE INDEX idx_specializations_featured ON specializations (user_id, is_featured) WHERE is_featured = TRUE;

CREATE TABLE specialization_skills (
  specialization_id UUID NOT NULL REFERENCES specializations(id) ON DELETE CASCADE,
  skill_id          UUID NOT NULL REFERENCES skills_taxonomy(id),
  kind              VARCHAR(20) NOT NULL CHECK (kind IN ('PRIMARY','SECONDARY')),
  PRIMARY KEY (specialization_id, skill_id, kind)
);
```

### Experience (updated) + `experience_skills`

```sql
CREATE TABLE experience (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,

    company_name VARCHAR(200) NOT NULL,
    company_logo_url TEXT,
    company_website TEXT,
    company_size VARCHAR(50),
    company_industry VARCHAR(100),

    job_title VARCHAR(200) NOT NULL,
    department VARCHAR(100),
    employment_type VARCHAR(50) CHECK (
        employment_type IN ('FULL_TIME','PART_TIME','CONTRACT','FREELANCE','INTERNSHIP','VOLUNTEER')
    ),

    start_date DATE NOT NULL,
    end_date DATE,
    is_current BOOLEAN DEFAULT FALSE,
    duration_months INTEGER,

    location VARCHAR(200),
    city VARCHAR(100),
    country_code CHAR(2),
    is_remote BOOLEAN DEFAULT FALSE,

    description TEXT,
    responsibilities TEXT[],
    achievements TEXT[],
    key_projects TEXT[],

    technologies_used TEXT[],

    team_size INTEGER,
    managed_team BOOLEAN DEFAULT FALSE,
    direct_reports INTEGER,

    is_verified BOOLEAN DEFAULT FALSE,
    verified_by VARCHAR(100),
    verification_contact_email VARCHAR(255),
    verification_contact_phone VARCHAR(50),

    display_order INTEGER DEFAULT 0,
    is_featured BOOLEAN DEFAULT FALSE,
    show_on_profile BOOLEAN DEFAULT TRUE,

    media_urls TEXT[],

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_experience_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT chk_experience_dates CHECK (end_date IS NULL OR end_date >= start_date)
);

CREATE INDEX idx_experience_user ON experience (user_id);
CREATE INDEX idx_experience_current ON experience (user_id, is_current) WHERE is_current = TRUE;
CREATE INDEX idx_experience_dates ON experience (start_date DESC, end_date DESC);
CREATE INDEX idx_experience_company ON experience (company_name);

CREATE TABLE experience_skills (
  experience_id UUID NOT NULL REFERENCES experience(id) ON DELETE CASCADE,
  skill_id      UUID NOT NULL REFERENCES skills_taxonomy(id),
  PRIMARY KEY (experience_id, skill_id)
);
```

### Education (updated) + `education_skills`

```sql
CREATE TABLE education (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,

    school_name VARCHAR(200) NOT NULL,
    school_logo_url TEXT,
    school_website TEXT,
    school_type VARCHAR(50),

    degree_type VARCHAR(50),
    degree_name VARCHAR(200),
    field_of_study VARCHAR(200),
    major VARCHAR(100),
    minor VARCHAR(100),

    start_year INTEGER,
    start_month INTEGER,
    end_year INTEGER,
    end_month INTEGER,
    graduation_date DATE,
    is_current BOOLEAN DEFAULT FALSE,

    location VARCHAR(200),
    city VARCHAR(100),
    country_code CHAR(2),

    grade_gpa VARCHAR(20),
    grade_scale VARCHAR(20),
    honors TEXT[],
    activities TEXT[],
    achievements TEXT[],
    thesis_title TEXT,
    thesis_url TEXT,

    description TEXT,
    coursework TEXT[],

    is_verified BOOLEAN DEFAULT FALSE,
    verification_document_url TEXT,
    diploma_url TEXT,
    transcript_url TEXT,

    display_order INTEGER DEFAULT 0,
    show_on_profile BOOLEAN DEFAULT TRUE,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_education_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT chk_education_years CHECK (end_year IS NULL OR end_year >= start_year)
);

CREATE INDEX idx_education_user ON education (user_id);
CREATE INDEX idx_education_current ON education (user_id, is_current) WHERE is_current = TRUE;
CREATE INDEX idx_education_school ON education (school_name);
CREATE INDEX idx_education_degree ON education (degree_type, field_of_study);

CREATE TABLE education_skills (
  education_id UUID NOT NULL REFERENCES education(id) ON DELETE CASCADE,
  skill_id     UUID NOT NULL REFERENCES skills_taxonomy(id),
  PRIMARY KEY (education_id, skill_id)
);
```

### Portfolios (updated) + `portfolio_skills`

```sql
CREATE TABLE portfolios (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,

    title VARCHAR(200) NOT NULL,
    description TEXT,
    short_description VARCHAR(500),
    project_url TEXT,

    project_type VARCHAR(50) CHECK (
        project_type IN ('PERSONAL','CLIENT_WORK','OPEN_SOURCE','ACADEMIC','COMPETITION','FREELANCE')
    ),

    cover_image_url TEXT NOT NULL,
    thumbnail_url TEXT,
    images_urls TEXT[],
    video_url TEXT,
    video_thumbnail_url TEXT,

    technologies TEXT[],
    tools_used TEXT[],

    client_name VARCHAR(200),
    client_industry VARCHAR(100),
    client_company_size VARCHAR(50),
    project_duration VARCHAR(50),
    project_budget_range VARCHAR(50),
    completion_date DATE,

    role VARCHAR(100),
    team_size INTEGER,
    my_contribution TEXT,
    challenges_overcome TEXT,
    results_achieved TEXT[],

    views_count INTEGER DEFAULT 0,
    likes_count INTEGER DEFAULT 0,
    comments_count INTEGER DEFAULT 0,
    shares_count INTEGER DEFAULT 0,

    is_featured BOOLEAN DEFAULT FALSE,
    is_public BOOLEAN DEFAULT TRUE,
    is_draft BOOLEAN DEFAULT FALSE,
    display_order INTEGER DEFAULT 0,

    github_url TEXT,
    live_demo_url TEXT,
    case_study_url TEXT,
    behance_url TEXT,
    dribbble_url TEXT,

    seo_keywords TEXT[],
    awards TEXT[],
    featured_in TEXT[],

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    published_at TIMESTAMPTZ,

    CONSTRAINT fk_portfolios_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_portfolios_user ON portfolios (user_id);
CREATE INDEX idx_portfolios_featured ON portfolios (user_id, is_featured) WHERE is_featured = TRUE;
CREATE INDEX idx_portfolios_public ON portfolios (is_public) WHERE is_public = TRUE;
CREATE INDEX idx_portfolios_project_type ON portfolios (project_type);

CREATE TABLE portfolio_skills (
  portfolio_id UUID NOT NULL REFERENCES portfolios(id) ON DELETE CASCADE,
  skill_id     UUID NOT NULL REFERENCES skills_taxonomy(id),
  PRIMARY KEY (portfolio_id, skill_id)
);
```
---

## Section 14.2: Moderation Actions (updated) + mapping

```sql
CREATE TABLE moderation_actions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,

    action_type VARCHAR(50) NOT NULL CHECK (
        action_type IN (
            'WARNING','TEMPORARY_SUSPENSION','PERMANENT_BAN','CONTENT_REMOVAL',
            'FEATURE_RESTRICTION','ACCOUNT_REVIEW','DEMOTION','REINSTATEMENT'
        )
    ),

    reason TEXT NOT NULL,
    internal_notes TEXT,
    evidence_urls TEXT[],

    duration_days INTEGER,
    restrictions JSONB,
    expires_at TIMESTAMPTZ,

    severity VARCHAR(20) CHECK (severity IN ('LOW','MEDIUM','HIGH','CRITICAL')),

    actioned_by UUID NOT NULL,
    approved_by UUID,
    approval_required BOOLEAN DEFAULT FALSE,

    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (
        status IN ('ACTIVE','EXPIRED','REVOKED','APPEALED','UNDER_APPEAL')
    ),

    user_notified BOOLEAN DEFAULT FALSE,
    notification_sent_at TIMESTAMPTZ,
    notification_method VARCHAR(50),

    appeal_allowed BOOLEAN DEFAULT TRUE,
    appeal_deadline TIMESTAMPTZ,

    is_automatic BOOLEAN DEFAULT FALSE,
    automatic_trigger VARCHAR(100),

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_moderation_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_moderation_actioned_by FOREIGN KEY (actioned_by) REFERENCES users(id),
    CONSTRAINT fk_moderation_approved_by FOREIGN KEY (approved_by) REFERENCES users(id)
);

CREATE INDEX idx_moderation_user ON moderation_actions (user_id);
CREATE INDEX idx_moderation_type ON moderation_actions (action_type);
CREATE INDEX idx_moderation_status ON moderation_actions (status, expires_at);
CREATE INDEX idx_moderation_actioned_by ON moderation_actions (actioned_by);
CREATE INDEX idx_moderation_severity ON moderation_actions (severity);

CREATE TABLE moderation_action_reports (
  action_id UUID NOT NULL REFERENCES moderation_actions(id) ON DELETE CASCADE,
  report_id UUID NOT NULL REFERENCES user_reports(id) ON DELETE CASCADE,
  PRIMARY KEY (action_id, report_id)
);
```
---

## Sections 44 & 42: No-overlap periods (updated)

```sql
CREATE TABLE availability_vacations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,

    start_date DATE NOT NULL,
    end_date DATE NOT NULL,

    vacation_type VARCHAR(50),
    notes TEXT,

    status VARCHAR(20) DEFAULT 'SCHEDULED' CHECK (
        status IN ('SCHEDULED','ACTIVE','COMPLETED','CANCELLED')
    ),

    auto_responder_enabled BOOLEAN DEFAULT FALSE,
    auto_responder_message TEXT,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    period DATERANGE GENERATED ALWAYS AS (daterange(start_date, end_date, '[]')) STORED,

    CONSTRAINT fk_availability_vacations_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT chk_vacation_dates CHECK (end_date >= start_date),
    CONSTRAINT no_overlapping_vacations EXCLUDE USING gist (user_id WITH =, period WITH &&)
);

CREATE INDEX idx_availability_vacations_user ON availability_vacations (user_id);
CREATE INDEX idx_availability_vacations_dates ON availability_vacations (start_date, end_date);
CREATE INDEX idx_availability_vacations_status ON availability_vacations (status);
CREATE INDEX idx_availability_vacations_user_period ON availability_vacations USING gist (user_id, period);
```

```sql
CREATE TABLE rate_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,

    rate_amount DECIMAL(10, 2) NOT NULL,
    rate_currency CHAR(3) DEFAULT 'USD',
    rate_type VARCHAR(20) DEFAULT 'HOURLY' CHECK (rate_type IN ('HOURLY','DAILY','FIXED')),

    effective_at DATE NOTNULL,
    effective_until DATE,

    change_reason VARCHAR(100),
    notes TEXT,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    effective_period DATERANGE GENERATED ALWAYS AS (daterange(effective_at, effective_until, '[]')) STORED,

    CONSTRAINT fk_rate_history_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT no_overlapping_rates EXCLUDE USING gist (user_id WITH =, effective_period WITH &&)
);

CREATE INDEX idx_rate_history_user ON rate_history (user_id, effective_at DESC);
CREATE INDEX idx_rate_history_effective ON rate_history (effective_at, effective_until);
CREATE INDEX idx_rate_history_user_period ON rate_history USING gist (user_id, effective_period);
```
> **Note:** Replace `NOTNULL` with `NOT NULL` if your editor doesn't auto-correct.

---

## Section 22: User Read Model (updated — generated `search_vector`)

```sql
CREATE TABLE user_read_model (
    user_id UUID PRIMARY KEY,

    full_name VARCHAR(300),
    display_name VARCHAR(200),
    email VARCHAR(255),
    user_type VARCHAR(20),
    account_status VARCHAR(20),

    professional_title VARCHAR(200),
    tagline VARCHAR(300),
    profile_picture_url TEXT,
    country_code CHAR(2),
    city VARCHAR(100),

    profile_completion_score INTEGER,
    reputation_score INTEGER,
    trust_score INTEGER,
    total_projects INTEGER,
    average_rating DECIMAL(3, 2),

    availability_status VARCHAR(20),
    hourly_rate DECIMAL(10, 2),

    top_skills JSONB,

    email_verified BOOLEAN,
    identity_verified BOOLEAN,

    search_vector tsvector GENERATED ALWAYS AS (
        setweight(to_tsvector('english', COALESCE(full_name,'')), 'A') ||
        setweight(to_tsvector('english', COALESCE(professional_title,'')), 'B') ||
        setweight(to_tsvector('english', COALESCE(tagline,'')), 'C')
    ) STORED,

    last_active_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user_read_model FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_user_read_search ON user_read_model USING gin(search_vector);
CREATE INDEX idx_user_read_reputation ON user_read_model (reputation_score DESC) WHERE account_status = 'ACTIVE';
```
---

## Section 19: Sessions (extra index only)

```sql
CREATE INDEX idx_sessions_user_active ON sessions (user_id) WHERE is_active = TRUE;
```
---

## New: Security Token Tables (centralized token storage)

```sql
CREATE TABLE email_verification_tokens (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  token_hash BYTEA NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  expires_at TIMESTAMPTZ NOT NULL
);
CREATE INDEX idx_evt_user_exp ON email_verification_tokens (user_id, expires_at);

CREATE TABLE phone_verification_codes (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  code_hash BYTEA NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  expires_at TIMESTAMPTZ NOT NULL
);
CREATE INDEX idx_pvc_user_exp ON phone_verification_codes (user_id, expires_at);

CREATE TABLE password_reset_tokens (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  token_hash BYTEA NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  expires_at TIMESTAMPTZ NOT NULL,
  attempts INTEGER NOT NULL DEFAULT 0
);
CREATE INDEX idx_prt_user_exp ON password_reset_tokens (user_id, expires_at);
```
