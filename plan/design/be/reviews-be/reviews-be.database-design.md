# REVIEWS-BE DATABASE DESIGN
**Skillsier Platform - Enterprise Scale (Upwork-like)**  
**PostgreSQL 16+**

---

## CRITICAL ALIGNMENT RULES:
1. Each domain folder in `internal/domain/{domain}/` = ONE main table
2. Table names follow domain folder names; when aggregated under Reviews, tables are prefixed with `review_` to reflect the domain boundary
3. Sub-entities within domain create related tables with `{domain}_{sub}` naming
4. All domains from folder structure are covered
5. Rich, production-ready fields for large-scale application

---

## GLOBAL EXTENSIONS

```sql
-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";        -- Trigram indexing for fuzzy search
CREATE EXTENSION IF NOT EXISTS "btree_gin";
CREATE EXTENSION IF NOT EXISTS "citext";         -- Case-insensitive text
CREATE EXTENSION IF NOT EXISTS "btree_gist";
```

---

=========================================
## SECTION 1: CORE REVIEW PRIMITIVES
=========================================

```sql
-- Domain: internal/domain/review/
-- Entity: review/entity.go
-- =========================================

-- Main Table: Reviews
CREATE TABLE reviews (
    -- Primary Key
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Idempotency
    idempotency_key UUID UNIQUE,
    
    -- Review Identity
    review_number VARCHAR(50) UNIQUE NOT NULL,       -- e.g., RVW-2025-001234
    
    -- Context
    contract_id UUID NOT NULL,
    job_id UUID,
    
    -- Parties
    reviewer_id UUID NOT NULL,
    reviewee_id UUID NOT NULL,
    
    -- Review Type
    review_type VARCHAR(30) NOT NULL CHECK (
        review_type IN ('CLIENT_TO_FREELANCER', 'FREELANCER_TO_CLIENT')
    ),
    
    -- Content
    title VARCHAR(300),
    body TEXT NOT NULL,
    body_length INTEGER,
    tags TEXT[],                                     -- ['professional', 'responsive', etc.]
    
    -- Rating Reference
    criteria_version_id UUID,
    overall_rating DECIMAL(3, 2) CHECK (
        overall_rating BETWEEN 1.0 AND 5.0
    ),
    
    -- Status Workflow
    status VARCHAR(20) DEFAULT 'DRAFT' CHECK (
        status IN ('DRAFT', 'SUBMITTED', 'PUBLISHED', 'WITHDRAWN', 'FLAGGED', 'HIDDEN')
    ),
    
    -- Lifecycle Timestamps
    drafted_at TIMESTAMPTZ,
    submitted_at TIMESTAMPTZ,
    published_at TIMESTAMPTZ,
    withdrawn_at TIMESTAMPTZ,
    last_edited_at TIMESTAMPTZ,
    
    -- Edit Window
    edit_allowed BOOLEAN DEFAULT TRUE,
    edit_deadline TIMESTAMPTZ,
    edit_count INTEGER DEFAULT 0,
    max_edits INTEGER DEFAULT 3,
    
    -- Response Permissions
    response_allowed BOOLEAN DEFAULT TRUE,
    response_deadline TIMESTAMPTZ,
    has_response BOOLEAN DEFAULT FALSE,
    
    -- Visibility
    is_public BOOLEAN DEFAULT TRUE,
    is_featured BOOLEAN DEFAULT FALSE,
    visibility VARCHAR(20) DEFAULT 'PUBLIC' CHECK (
        visibility IN ('PUBLIC', 'PRIVATE', 'CONTACTS_ONLY', 'HIDDEN')
    ),
    
    -- Engagement
    helpful_votes INTEGER DEFAULT 0,
    unhelpful_votes INTEGER DEFAULT 0,
    view_count INTEGER DEFAULT 0,
    
    -- Moderation
    is_flagged BOOLEAN DEFAULT FALSE,
    flag_count INTEGER DEFAULT 0,
    moderation_status VARCHAR(20) DEFAULT 'APPROVED' CHECK (
        moderation_status IN ('PENDING', 'APPROVED', 'REJECTED', 'UNDER_REVIEW')
    ),
    moderated_at TIMESTAMPTZ,
    moderated_by UUID,
    moderation_notes TEXT,
    
    -- Redaction
    has_redactions BOOLEAN DEFAULT FALSE,
    redacted_at TIMESTAMPTZ,
    
    -- Contract Context
    contract_value DECIMAL(15, 2),
    contract_duration_days INTEGER,
    milestones_completed INTEGER,
    
    -- Double-Blind Association
    double_blind_window_id UUID,
    
    -- Soft Delete
    is_deleted BOOLEAN DEFAULT FALSE,
    deleted_at TIMESTAMPTZ,
    deleted_by UUID,
    deletion_reason TEXT,
    
    -- Audit
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by UUID NOT NULL,
    
    -- Metadata
    metadata JSONB,
    
    -- Constraints
    CONSTRAINT chk_reviews_parties CHECK (reviewer_id != reviewee_id),
    CONSTRAINT chk_reviews_edit_deadline CHECK (edit_deadline IS NULL OR edit_deadline > submitted_at)
);

CREATE INDEX idx_reviews_contract ON reviews (contract_id);
CREATE INDEX idx_reviews_reviewer ON reviews (reviewer_id, created_at DESC);
CREATE INDEX idx_reviews_reviewee ON reviews (reviewee_id, published_at DESC);
CREATE INDEX idx_reviews_status ON reviews (status);
CREATE INDEX idx_reviews_type ON reviews (review_type);
CREATE INDEX idx_reviews_published ON reviews (published_at DESC) WHERE status = 'PUBLISHED';
CREATE INDEX idx_reviews_featured ON reviews (is_featured, published_at DESC) WHERE is_featured = TRUE;
CREATE INDEX idx_reviews_moderation ON reviews (moderation_status) WHERE moderation_status != 'APPROVED';
CREATE INDEX idx_reviews_flagged ON reviews (is_flagged, flag_count DESC) WHERE is_flagged = TRUE;
CREATE INDEX idx_reviews_double_blind ON reviews (double_blind_window_id) WHERE double_blind_window_id IS NOT NULL;
CREATE INDEX idx_reviews_deleted ON reviews (is_deleted, deleted_at DESC) WHERE is_deleted = TRUE;

COMMENT ON TABLE reviews IS 'Reviews - maps to internal/domain/review/entity.go';

-- Review Ratings (Per-Dimension Scores)
CREATE TABLE review_ratings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    review_id UUID NOT NULL,
    
    -- Dimension
    dimension_name VARCHAR(100) NOT NULL,
    dimension_weight DECIMAL(5, 2),
    
    -- Score
    raw_score DECIMAL(3, 2) NOT NULL CHECK (
        raw_score BETWEEN 1.0 AND 5.0
    ),
    normalized_score DECIMAL(5, 4),
    weighted_score DECIMAL(5, 4),
    
    -- Context
    criteria_version_id UUID,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_review_ratings_review FOREIGN KEY (review_id) 
        REFERENCES reviews(id) ON DELETE CASCADE,
    CONSTRAINT uq_review_ratings_dimension UNIQUE (review_id, dimension_name)
);

CREATE INDEX idx_review_ratings_review ON review_ratings (review_id);
CREATE INDEX idx_review_ratings_dimension ON review_ratings (dimension_name);

-- Review Responses
CREATE TABLE review_responses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    review_id UUID NOT NULL UNIQUE,
    
    -- Responder
    responder_id UUID NOT NULL,
    
    -- Content
    response_text TEXT NOT NULL,
    response_length INTEGER,
    
    -- Status
    is_published BOOLEAN DEFAULT TRUE,
    is_deleted BOOLEAN DEFAULT FALSE,
    
    -- Edit Tracking
    edit_count INTEGER DEFAULT 0,
    last_edited_at TIMESTAMPTZ,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ,
    
    CONSTRAINT fk_review_responses_review FOREIGN KEY (review_id) 
        REFERENCES reviews(id) ON DELETE CASCADE
);

CREATE INDEX idx_review_responses_review ON review_responses (review_id);
CREATE INDEX idx_review_responses_responder ON review_responses (responder_id);
CREATE INDEX idx_review_responses_deleted ON review_responses (is_deleted) WHERE is_deleted = TRUE;

-- Helpful Votes
CREATE TABLE review_helpful_votes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    review_id UUID NOT NULL,
    voter_id UUID NOT NULL,
    
    -- Vote
    is_helpful BOOLEAN NOT NULL,
    
    -- Change Tracking
    vote_changed BOOLEAN DEFAULT FALSE,
    original_vote BOOLEAN,
    changed_at TIMESTAMPTZ,
    
    -- Window Enforcement (24h)
    can_change BOOLEAN DEFAULT TRUE,
    change_deadline TIMESTAMPTZ,
    
    voted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_review_helpful_votes_review FOREIGN KEY (review_id) 
        REFERENCES reviews(id) ON DELETE CASCADE,
    CONSTRAINT uq_review_helpful_votes_voter UNIQUE (review_id, voter_id)
);

CREATE INDEX idx_review_helpful_votes_review ON review_helpful_votes (review_id);
CREATE INDEX idx_review_helpful_votes_voter ON review_helpful_votes (voter_id);
CREATE INDEX idx_review_helpful_votes_changeable ON review_helpful_votes (can_change, change_deadline) 
    WHERE can_change = TRUE;

-- Domain: internal/domain/review_draft/
-- Entity: review_draft/entity.go
-- =========================================

CREATE TABLE review_drafts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Idempotency
    idempotency_key UUID UNIQUE,
    
    -- Author
    author_id UUID NOT NULL,
    
    -- Context
    contract_id UUID NOT NULL,
    reviewee_id UUID NOT NULL,
    
    -- Draft Content
    title VARCHAR(300),
    body TEXT,
    tags TEXT[],
    
    -- Ratings Draft
    rating_map JSONB,                                -- {dimension: score}
    overall_rating DECIMAL(3, 2),
    
    -- Auto-Save
    auto_saved BOOLEAN DEFAULT FALSE,
    
    -- Expiry
    expires_at TIMESTAMPTZ,
    is_expired BOOLEAN DEFAULT FALSE,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT uq_review_drafts_author_contract UNIQUE (author_id, contract_id)
);

CREATE INDEX idx_review_drafts_author ON review_drafts (author_id, updated_at DESC);
CREATE INDEX idx_review_drafts_contract ON review_drafts (contract_id);
CREATE INDEX idx_review_drafts_expired ON review_drafts (is_expired, expires_at) WHERE is_expired = FALSE;

COMMENT ON TABLE review_drafts IS 'Review drafts - maps to internal/domain/review_draft/entity.go';

-- Domain: internal/domain/eligibility/
-- Entity: eligibility/entity.go
-- =========================================

CREATE TABLE review_eligibility (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Idempotency
    idempotency_key UUID UNIQUE,
    
    -- Parties
    reviewer_id UUID NOT NULL,
    reviewee_id UUID NOT NULL,
    contract_id UUID NOT NULL,
    
    -- Eligibility Result
    is_eligible BOOLEAN NOT NULL,
    not_before TIMESTAMPTZ,
    not_after TIMESTAMPTZ,
    
    -- Reasons
    eligibility_reason TEXT,
    ineligibility_reasons JSONB,                     -- Array of reasons
    
    -- Checks Performed
    contract_completed BOOLEAN,
    payment_verified BOOLEAN,
    within_time_window BOOLEAN,
    daily_limit_ok BOOLEAN,
    cooldown_ok BOOLEAN,
    
    -- Limits Tracking
    reviews_today INTEGER,
    daily_limit INTEGER DEFAULT 3,
    last_review_to_same_user TIMESTAMPTZ,
    cooldown_hours INTEGER DEFAULT 24,
    
    -- Decision Metadata
    checked_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    policy_version VARCHAR(50),
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT uq_review_eligibility_check UNIQUE (reviewer_id, contract_id, checked_at)
);

CREATE INDEX idx_review_eligibility_reviewer ON review_eligibility (reviewer_id, checked_at DESC);
CREATE INDEX idx_review_eligibility_contract ON review_eligibility (contract_id);
CREATE INDEX idx_review_eligibility_eligible ON review_eligibility (is_eligible);
CREATE INDEX idx_review_eligibility_reviewee ON review_eligibility (reviewer_id, reviewee_id);

COMMENT ON TABLE review_eligibility IS 'Review eligibility checks - maps to internal/domain/eligibility/entity.go';

-- Eligibility Policy Configuration
CREATE TABLE review_eligibility_policies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    policy_version VARCHAR(50) UNIQUE NOT NULL,
    
    -- Policy Rules
    requires_completed_contract BOOLEAN DEFAULT TRUE,
    requires_payment_verification BOOLEAN DEFAULT TRUE,
    
    -- Time Windows
    review_window_days INTEGER DEFAULT 30,
    min_hours_after_completion INTEGER DEFAULT 24,
    max_hours_after_completion INTEGER DEFAULT 720,  -- 30 days
    
    -- Limits
    daily_review_limit INTEGER DEFAULT 3,
    cooldown_hours INTEGER DEFAULT 24,
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    
    -- Audit
    created_by UUID,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    activated_at TIMESTAMPTZ
);

CREATE INDEX idx_review_eligibility_policies_version ON review_eligibility_policies (policy_version);
CREATE INDEX idx_review_eligibility_policies_active ON review_eligibility_policies (is_active) 
    WHERE is_active = TRUE;

```

---

=========================================
## SECTION 2: RATING & CRITERIA
=========================================

```sql
-- Domain: internal/domain/rating/
-- Entity: rating/entity.go
-- =========================================

CREATE TABLE rating_criteria (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Criteria Identity
    criteria_name VARCHAR(200) NOT NULL,
    category VARCHAR(100) NOT NULL,                  -- Job category
    
    -- Version Control
    version INTEGER NOT NULL DEFAULT 1,
    is_active BOOLEAN DEFAULT FALSE,
    
    -- Dimensions
    dimensions JSONB NOT NULL,                       -- Array of {name, weight, scale, description}
    dimension_count INTEGER NOT NULL,
    
    -- Weights
    total_weight DECIMAL(5, 2) DEFAULT 1.0,
    
    -- Scale
    min_score DECIMAL(3, 2) DEFAULT 1.0,
    max_score DECIMAL(3, 2) DEFAULT 5.0,
    
    -- Metadata
    description TEXT,
    
    -- Audit
    created_by UUID,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    activated_at TIMESTAMPTZ,
    deprecated_at TIMESTAMPTZ,
    
    CONSTRAINT uq_rating_criteria_category_version UNIQUE (category, version)
);

CREATE INDEX idx_rating_criteria_category ON rating_criteria (category, version DESC);
CREATE INDEX idx_rating_criteria_active ON rating_criteria (is_active) WHERE is_active = TRUE;
CREATE INDEX idx_rating_criteria_version ON rating_criteria (version);

COMMENT ON TABLE rating_criteria IS 'Rating criteria versions - maps to internal/domain/rating/entity.go';

-- Rating Aggregates (Per User)
CREATE TABLE rating_aggregates (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- User
    user_id UUID NOT NULL UNIQUE,
    user_type VARCHAR(30) NOT NULL CHECK (
        user_type IN ('FREELANCER', 'CLIENT')
    ),
    
    -- Overall Statistics
    total_reviews INTEGER DEFAULT 0,
    average_rating DECIMAL(3, 2),
    weighted_average_rating DECIMAL(5, 4),
    wilson_score DECIMAL(5, 4),                      -- Wilson score for ranking
    
    -- Rating Distribution
    rating_1_count INTEGER DEFAULT 0,
    rating_2_count INTEGER DEFAULT 0,
    rating_3_count INTEGER DEFAULT 0,
    rating_4_count INTEGER DEFAULT 0,
    rating_5_count INTEGER DEFAULT 0,
    
    -- Dimension Breakdown
    dimension_averages JSONB,                        -- {dimension: avg_score}
    
    -- Bayesian Prior
    prior_mean DECIMAL(3, 2) DEFAULT 3.5,
    prior_weight DECIMAL(5, 2) DEFAULT 5.0,
    
    -- Time Decay
    recent_rating_avg DECIMAL(3, 2),                 -- Last 90 days
    decay_factor DECIMAL(5, 4) DEFAULT 0.95,
    
    -- Last Updated
    last_review_at TIMESTAMPTZ,
    recalculated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_rating_aggregates_user ON rating_aggregates (user_id);
CREATE INDEX idx_rating_aggregates_type ON rating_aggregates (user_type);
CREATE INDEX idx_rating_aggregates_wilson ON rating_aggregates (wilson_score DESC);
CREATE INDEX idx_rating_aggregates_average ON rating_aggregates (average_rating DESC);

COMMENT ON TABLE rating_aggregates IS 'Rating aggregates per user';

```

---

=========================================
## SECTION 3: REPUTATION
=========================================

```sql
-- Domain: internal/domain/reputation/
-- Entity: reputation/entity.go
-- =========================================

CREATE TABLE reputation (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- User
    user_id UUID NOT NULL UNIQUE,
    user_type VARCHAR(30) NOT NULL CHECK (
        user_type IN ('FREELANCER', 'CLIENT')
    ),
    
    -- Reputation Score
    reputation_score DECIMAL(10, 4) NOT NULL,
    normalized_score DECIMAL(3, 2),                  -- 1-5 scale
    percentile DECIMAL(5, 2),                        -- User's percentile rank
    
    -- Score Components
    rating_component DECIMAL(10, 4),
    recency_component DECIMAL(10, 4),
    volume_component DECIMAL(10, 4),
    consistency_component DECIMAL(10, 4),
    
    -- Bayesian Calculation
    prior_mean DECIMAL(3, 2) DEFAULT 3.5,
    prior_weight DECIMAL(5, 2) DEFAULT 5.0,
    posterior_mean DECIMAL(5, 4),
    
    -- Time Decay
    decay_factor DECIMAL(5, 4) DEFAULT 0.95,
    half_life_days INTEGER DEFAULT 180,
    
    -- Recency Boost
    recency_boost DECIMAL(5, 4),
    recent_positive_trend BOOLEAN DEFAULT FALSE,
    
    -- Volume & Consistency
    total_reviews INTEGER DEFAULT 0,
    recent_reviews_90d INTEGER DEFAULT 0,
    consistency_score DECIMAL(5, 4),                 -- Variance measure
    
    -- Badges & Achievements
    is_top_rated BOOLEAN DEFAULT FALSE,
    top_rated_since TIMESTAMPTZ,
    
    -- Trend Analysis
    trend VARCHAR(20) CHECK (
        trend IN ('IMPROVING', 'STABLE', 'DECLINING')
    ),
    trend_slope DECIMAL(10, 6),
    
    -- Last Calculation
    calculated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    calculation_version VARCHAR(50),
    
    -- Audit
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_reputation_user ON reputation (user_id);
CREATE INDEX idx_reputation_score ON reputation (reputation_score DESC);
CREATE INDEX idx_reputation_normalized ON reputation (normalized_score DESC);
CREATE INDEX idx_reputation_percentile ON reputation (percentile DESC);
CREATE INDEX idx_reputation_top_rated ON reputation (is_top_rated) WHERE is_top_rated = TRUE;
CREATE INDEX idx_reputation_trend ON reputation (trend);

COMMENT ON TABLE reputation IS 'User reputation scores - maps to internal/domain/reputation/entity.go';

-- Reputation History (Time Series)
CREATE TABLE reputation_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Historical Score
    reputation_score DECIMAL(10, 4) NOT NULL,
    normalized_score DECIMAL(3, 2),
    
    -- Components
    rating_component DECIMAL(10, 4),
    recency_component DECIMAL(10, 4),
    volume_component DECIMAL(10, 4),
    
    -- Context
    total_reviews INTEGER,
    average_rating DECIMAL(3, 2),
    
    -- Timestamp
    recorded_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_reputation_history_user ON reputation_history (user_id, recorded_at DESC);
CREATE INDEX idx_reputation_history_recorded ON reputation_history (recorded_at DESC);

-- Top Rated Eligibility
CREATE TABLE reputation_top_rated_rules (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Rule Configuration
    rule_name VARCHAR(200) NOT NULL UNIQUE,
    
    -- Thresholds
    min_reputation_score DECIMAL(10, 4) DEFAULT 4.5,
    min_total_reviews INTEGER DEFAULT 50,
    min_recent_reviews_90d INTEGER DEFAULT 10,
    min_average_rating DECIMAL(3, 2) DEFAULT 4.7,
    max_consistency_variance DECIMAL(5, 4) DEFAULT 0.3,
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    
    -- Audit
    created_by UUID,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_reputation_top_rated_rules_active ON reputation_top_rated_rules (is_active) 
    WHERE is_active = TRUE;

```

---

=========================================
## SECTION 4: BADGE SYSTEM
=========================================

```sql
-- Domain: internal/domain/badge/
-- Entity: badge/entity.go
-- =========================================

CREATE TABLE badges (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Badge Identity
    badge_code VARCHAR(100) UNIQUE NOT NULL,
    badge_name VARCHAR(200) NOT NULL,
    description TEXT,
    
    -- Category
    category VARCHAR(100),                           -- 'PERFORMANCE', 'ACHIEVEMENT', 'VERIFICATION'
    
    -- Levels
    has_levels BOOLEAN DEFAULT FALSE,
    levels JSONB,                                    -- Array of {level, name, criteria}
    
    -- Visual
    icon_url TEXT,
    color VARCHAR(20),
    
    -- Criteria
    criteria_expression JSONB NOT NULL,              -- Rule-based criteria
    
    -- Status
    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (
        status IN ('ACTIVE', 'ARCHIVED', 'DEPRECATED')
    ),
    
    -- Display
    display_order INTEGER DEFAULT 100,
    is_featured BOOLEAN DEFAULT FALSE,
    
    -- Audit
    created_by UUID,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    archived_at TIMESTAMPTZ
);

CREATE INDEX idx_badges_code ON badges (badge_code);
CREATE INDEX idx_badges_category ON badges (category);
CREATE INDEX idx_badges_status ON badges (status);
CREATE INDEX idx_badges_featured ON badges (is_featured, display_order) WHERE is_featured = TRUE;

COMMENT ON TABLE badges IS 'Badge definitions - maps to internal/domain/badge/entity.go';

-- Domain: internal/domain/user_badge/
-- Entity: user_badge/entity.go
-- =========================================

CREATE TABLE user_badges (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Assignment
    user_id UUID NOT NULL,
    badge_id UUID NOT NULL,
    
    -- Level (if applicable)
    badge_level VARCHAR(50),                         -- 'BRONZE', 'SILVER', 'GOLD', 'PLATINUM'
    
    -- Status
    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (
        status IN ('ACTIVE', 'REVOKED', 'EXPIRED')
    ),
    
    -- Achievement
    achieved_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    achievement_reason TEXT,
    
    -- Renewal (for time-limited badges)
    is_renewable BOOLEAN DEFAULT FALSE,
    expires_at TIMESTAMPTZ,
    last_renewed_at TIMESTAMPTZ,
    renewal_count INTEGER DEFAULT 0,
    
    -- Revocation
    revoked_at TIMESTAMPTZ,
    revoked_by UUID,
    revocation_reason TEXT,
    
    -- Visibility
    is_visible BOOLEAN DEFAULT TRUE,
    display_order INTEGER,
    
    -- Audit
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_user_badges_badge FOREIGN KEY (badge_id) 
        REFERENCES badges(id) ON DELETE CASCADE,
    CONSTRAINT uq_user_badges_user_badge UNIQUE (user_id, badge_id, badge_level)
);

CREATE INDEX idx_user_badges_user ON user_badges (user_id, achieved_at DESC);
CREATE INDEX idx_user_badges_badge ON user_badges (badge_id);
CREATE INDEX idx_user_badges_status ON user_badges (status);
CREATE INDEX idx_user_badges_active ON user_badges (user_id, is_visible) 
    WHERE status = 'ACTIVE' AND is_visible = TRUE;
CREATE INDEX idx_user_badges_expires ON user_badges (expires_at) WHERE expires_at IS NOT NULL;

COMMENT ON TABLE user_badges IS 'User badge assignments - maps to internal/domain/user_badge/entity.go';

-- Badge Eligibility Checks
CREATE TABLE badge_eligibility_checks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Check Details
    user_id UUID NOT NULL,
    badge_id UUID NOT NULL,
    
    -- Result
    is_eligible BOOLEAN NOT NULL,
    eligibility_reason TEXT,
    
    -- Criteria Evaluation
    criteria_met JSONB,                              -- Which criteria passed/failed
    
    -- Action Taken
    action_taken VARCHAR(50),                        -- 'AWARDED', 'UPGRADED', 'NONE'
    
    checked_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_badge_eligibility_checks_badge FOREIGN KEY (badge_id) 
        REFERENCES badges(id) ON DELETE CASCADE
);

CREATE INDEX idx_badge_eligibility_checks_user ON badge_eligibility_checks (user_id, checked_at DESC);
CREATE INDEX idx_badge_eligibility_checks_badge ON badge_eligibility_checks (badge_id);
CREATE INDEX idx_badge_eligibility_checks_eligible ON badge_eligibility_checks (is_eligible);

```

---

=========================================
## SECTION 5: DOUBLE-BLIND & WINDOWS
=========================================

```sql
-- Domain: internal/domain/double_blind/
-- Entity: double_blind/entity.go
-- =========================================

CREATE TABLE double_blind_windows (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Idempotency
    idempotency_key UUID UNIQUE,
    
    -- Contract Context
    contract_id UUID NOT NULL UNIQUE,
    
    -- Parties
    client_id UUID NOT NULL,
    freelancer_id UUID NOT NULL,
    
    -- Window Configuration
    duration_hours INTEGER DEFAULT 168,              -- 7 days default
    grace_period_hours INTEGER DEFAULT 24,
    
    -- Window Status
    status VARCHAR(20) DEFAULT 'OPEN' CHECK (
        status IN ('SCHEDULED', 'OPEN', 'CLOSED', 'EXPIRED', 'CANCELLED')
    ),
    
    -- Timeline
    opens_at TIMESTAMPTZ NOT NULL,
    closes_at TIMESTAMPTZ NOT NULL,
    closed_at TIMESTAMPTZ,
    
    -- Submissions
    client_submitted BOOLEAN DEFAULT FALSE,
    client_submitted_at TIMESTAMPTZ,
    client_review_id UUID,
    
    freelancer_submitted BOOLEAN DEFAULT FALSE,
    freelancer_submitted_at TIMESTAMPTZ,
    freelancer_review_id UUID,
    
    both_submitted BOOLEAN DEFAULT FALSE,
    
    -- Publishing
    auto_publish BOOLEAN DEFAULT TRUE,
    published_at TIMESTAMPTZ,
    publishing_method VARCHAR(30),                   -- 'BOTH', 'SOLO_CLIENT', 'SOLO_FREELANCER', 'TIMEOUT'
    
    -- Reminders
    reminders_scheduled BOOLEAN DEFAULT FALSE,
    reminders_sent INTEGER DEFAULT 0,
    
    -- Audit
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by_system BOOLEAN DEFAULT TRUE
);

CREATE INDEX idx_double_blind_windows_contract ON double_blind_windows (contract_id);
CREATE INDEX idx_double_blind_windows_status ON double_blind_windows (status);
CREATE INDEX idx_double_blind_windows_closes ON double_blind_windows (closes_at) 
    WHERE status = 'OPEN';
CREATE INDEX idx_double_blind_windows_client ON double_blind_windows (client_id);
CREATE INDEX idx_double_blind_windows_freelancer ON double_blind_windows (freelancer_id);
CREATE INDEX idx_double_blind_windows_publishing ON double_blind_windows (both_submitted, published_at);

COMMENT ON TABLE double_blind_windows IS 'Double-blind review windows - maps to internal/domain/double_blind/entity.go';

-- Domain: internal/domain/review_reminder/
-- Entity: review_reminder/entity.go
-- =========================================

CREATE TABLE review_reminders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Association
    double_blind_window_id UUID NOT NULL,
    user_id UUID NOT NULL,
    
    -- Reminder Type
    reminder_type VARCHAR(30) CHECK (
        reminder_type IN ('WINDOW_OPENED', 'MIDPOINT', 'CLOSING_SOON', 'FINAL')
    ),
    
    -- Schedule
    scheduled_for TIMESTAMPTZ NOT NULL,
    
    -- Status
    status VARCHAR(20) DEFAULT 'SCHEDULED' CHECK (
        status IN ('SCHEDULED', 'SENT', 'FAILED', 'CANCELLED')
    ),
    
    -- Delivery
    sent_at TIMESTAMPTZ,
    delivery_channels TEXT[],                        -- ['EMAIL', 'IN_APP', 'PUSH']
    
    -- Retry
    retry_count INTEGER DEFAULT 0,
    max_retries INTEGER DEFAULT 3,
    last_error TEXT,
    
    -- Cancellation
    cancelled_at TIMESTAMPTZ,
    cancellation_reason TEXT,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_review_reminders_window FOREIGN KEY (double_blind_window_id) 
        REFERENCES double_blind_windows(id) ON DELETE CASCADE,
    CONSTRAINT uq_review_reminders UNIQUE (double_blind_window_id, user_id, reminder_type)
);

CREATE INDEX idx_review_reminders_window ON review_reminders (double_blind_window_id);
CREATE INDEX idx_review_reminders_user ON review_reminders (user_id);
CREATE INDEX idx_review_reminders_status ON review_reminders (status, scheduled_for);
CREATE INDEX idx_review_reminders_scheduled ON review_reminders (scheduled_for) 
    WHERE status = 'SCHEDULED';

COMMENT ON TABLE review_reminders IS 'Review reminders - maps to internal/domain/review_reminder/entity.go';

```

---

=========================================
## SECTION 6: PRIVATE FEEDBACK
=========================================

```sql
-- Domain: internal/domain/feedback/
-- Entity: feedback/entity.go
-- =========================================

CREATE TABLE private_feedback (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Idempotency
    idempotency_key UUID UNIQUE,
    
    -- Context
    contract_id UUID NOT NULL,
    reviewer_id UUID NOT NULL,
    reviewee_id UUID NOT NULL,
    
    -- Associated Review (optional)
    review_id UUID,
    
    -- Category
    category VARCHAR(100),                           -- 'CONCERN', 'PRAISE', 'ISSUE', 'OTHER'
    
    -- Content
    feedback_text TEXT NOT NULL,
    feedback_length INTEGER,
    
    -- NPS-Style Score (optional)
    nps_score INTEGER CHECK (
        nps_score BETWEEN 0 AND 10
    ),
    
    -- Confidentiality
    is_confidential BOOLEAN DEFAULT TRUE,
    visibility VARCHAR(20) DEFAULT 'ADMIN_ONLY' CHECK (
        visibility IN ('ADMIN_ONLY', 'INTERNAL', 'PLATFORM_TEAM')
    ),
    
    -- Review Status
    reviewed_by_admin BOOLEAN DEFAULT FALSE,
    reviewed_at TIMESTAMPTZ,
    reviewed_by UUID,
    admin_notes TEXT,
    
    -- Action Taken
    action_taken VARCHAR(100),
    action_notes TEXT,
    
    -- Soft Delete
    is_deleted BOOLEAN DEFAULT FALSE,
    deleted_at TIMESTAMPTZ,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_private_feedback_contract ON private_feedback (contract_id);
CREATE INDEX idx_private_feedback_reviewer ON private_feedback (reviewer_id);
CREATE INDEX idx_private_feedback_reviewee ON private_feedback (reviewee_id);
CREATE INDEX idx_private_feedback_review ON private_feedback (review_id) WHERE review_id IS NOT NULL;
CREATE INDEX idx_private_feedback_category ON private_feedback (category);
CREATE INDEX idx_private_feedback_reviewed ON private_feedback (reviewed_by_admin, reviewed_at);

COMMENT ON TABLE private_feedback IS 'Private feedback - maps to internal/domain/feedback/entity.go';

```

---

=========================================
## SECTION 7: SAFETY & GOVERNANCE
=========================================

```sql
-- Domain: internal/domain/flag/
-- Entity: flag/entity.go
-- =========================================

CREATE TABLE review_flags (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Idempotency
    idempotency_key UUID UNIQUE,
    
    -- Target
    review_id UUID NOT NULL,
    
    -- Reporter
    reporter_id UUID NOT NULL,
    reporter_type VARCHAR(30),                       -- 'USER', 'SYSTEM', 'ADMIN'
    
    -- Flag Reason
    flag_reason VARCHAR(100) NOT NULL CHECK (
        flag_reason IN ('SPAM', 'HARASSMENT', 'FAKE', 'INAPPROPRIATE', 'PII', 'COPYRIGHT', 'OTHER')
    ),
    custom_reason TEXT,
    
    -- Evidence
    evidence_text TEXT,
    evidence_urls TEXT[],
    
    -- Status
    status VARCHAR(20) DEFAULT 'OPEN' CHECK (
        status IN ('OPEN', 'UNDER_REVIEW', 'RESOLVED', 'DISMISSED')
    ),
    
    -- Resolution
    resolved_at TIMESTAMPTZ,
    resolved_by UUID,
    resolution_action VARCHAR(50),                   -- 'HIDE', 'REDACT', 'REMOVE', 'NO_ACTION'
    resolution_notes TEXT,
    
    -- Priority
    priority VARCHAR(20) DEFAULT 'NORMAL' CHECK (
        priority IN ('LOW', 'NORMAL', 'HIGH', 'CRITICAL')
    ),
    
    -- Auto-Flag Confidence (if system-flagged)
    confidence_score DECIMAL(5, 4),
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_review_flags_review FOREIGN KEY (review_id) 
        REFERENCES reviews(id) ON DELETE CASCADE
);

CREATE INDEX idx_review_flags_review ON review_flags (review_id);
CREATE INDEX idx_review_flags_reporter ON review_flags (reporter_id);
CREATE INDEX idx_review_flags_status ON review_flags (status);
CREATE INDEX idx_review_flags_open ON review_flags (status, priority, created_at DESC) 
    WHERE status IN ('OPEN', 'UNDER_REVIEW');
CREATE INDEX idx_review_flags_reason ON review_flags (flag_reason);

COMMENT ON TABLE review_flags IS 'Review flags - maps to internal/domain/flag/entity.go';

-- Domain: internal/domain/moderation/
-- Entity: moderation/entity.go
-- =========================================

CREATE TABLE review_moderation (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Target
    review_id UUID NOT NULL UNIQUE,
    
    -- State Machine
    moderation_state VARCHAR(30) DEFAULT 'PENDING' CHECK (
        moderation_state IN ('PENDING', 'AUTO_APPROVED', 'MANUAL_REVIEW', 'APPROVED', 
                             'REJECTED', 'HIDDEN', 'REMOVED')
    ),
    
    -- Auto-Flag Signals
    auto_flagged BOOLEAN DEFAULT FALSE,
    auto_flag_reasons JSONB,                         -- Array of detected issues
    auto_flag_confidence DECIMAL(5, 4),
    
    -- Heuristic Signals
    duplicate_content_detected BOOLEAN DEFAULT FALSE,
    burst_activity_detected BOOLEAN DEFAULT FALSE,
    suspicious_ip_detected BOOLEAN DEFAULT FALSE,
    sentiment_score DECIMAL(5, 4),
    
    -- Manual Review
    requires_manual_review BOOLEAN DEFAULT FALSE,
    manual_reviewer_id UUID,
    manual_reviewed_at TIMESTAMPTZ,
    manual_review_notes TEXT,
    
    -- Actions Taken
    actions_taken JSONB,                             -- Array of actions
    
    -- Appeals
    appeal_allowed BOOLEAN DEFAULT TRUE,
    has_active_appeal BOOLEAN DEFAULT FALSE,
    
    -- Audit
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    state_changed_at TIMESTAMPTZ,
    
    CONSTRAINT fk_review_moderation_review FOREIGN KEY (review_id) 
        REFERENCES reviews(id) ON DELETE CASCADE
);

CREATE INDEX idx_review_moderation_review ON review_moderation (review_id);
CREATE INDEX idx_review_moderation_state ON review_moderation (moderation_state);
CREATE INDEX idx_review_moderation_auto_flagged ON review_moderation (auto_flagged) 
    WHERE auto_flagged = TRUE;
CREATE INDEX idx_review_moderation_manual ON review_moderation (requires_manual_review) 
    WHERE requires_manual_review = TRUE;
CREATE INDEX idx_review_moderation_appeal ON review_moderation (has_active_appeal) 
    WHERE has_active_appeal = TRUE;

COMMENT ON TABLE review_moderation IS 'Review moderation - maps to internal/domain/moderation/entity.go';

-- Moderation State History
CREATE TABLE review_moderation_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    review_moderation_id UUID NOT NULL,
    
    -- State Change
    previous_state VARCHAR(30),
    new_state VARCHAR(30) NOT NULL,
    
    -- Actor
    changed_by UUID,
    changed_by_type VARCHAR(30),                     -- 'SYSTEM', 'MODERATOR', 'ADMIN'
    
    -- Reason
    change_reason TEXT,
    
    occurred_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_review_moderation_history_moderation FOREIGN KEY (review_moderation_id) 
        REFERENCES review_moderation(id) ON DELETE CASCADE
);

CREATE INDEX idx_review_moderation_history_moderation ON review_moderation_history (review_moderation_id, occurred_at DESC);

-- Domain: internal/domain/appeal/
-- Entity: appeal/entity.go
-- =========================================

CREATE TABLE review_appeals (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Idempotency
    idempotency_key UUID UNIQUE,
    
    -- Target
    review_id UUID NOT NULL,
    
    -- Appellant
    appellant_id UUID NOT NULL,
    
    -- Appeal Details
    appeal_reason TEXT NOT NULL,
    appeal_category VARCHAR(50),
    
    -- Evidence
    evidence_text TEXT,
    evidence_file_ids TEXT[],
    
    -- Status
    status VARCHAR(20) DEFAULT 'OPEN' CHECK (
        status IN ('OPEN', 'UNDER_REVIEW', 'APPROVED', 'DENIED', 'WITHDRAWN')
    ),
    
    -- Review
    reviewed_by UUID,
    reviewed_at TIMESTAMPTZ,
    
    -- Outcome
    outcome VARCHAR(20),
    outcome_reason TEXT,
    actions_taken JSONB,
    
    -- Timeline
    opened_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    resolved_at TIMESTAMPTZ,
    
    -- Withdrawal
    withdrawn_at TIMESTAMPTZ,
    withdrawal_reason TEXT,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_review_appeals_review FOREIGN KEY (review_id) 
        REFERENCES reviews(id) ON DELETE CASCADE
);

CREATE INDEX idx_review_appeals_review ON review_appeals (review_id);
CREATE INDEX idx_review_appeals_appellant ON review_appeals (appellant_id);
CREATE INDEX idx_review_appeals_status ON review_appeals (status);
CREATE INDEX idx_review_appeals_open ON review_appeals (status, opened_at DESC) 
    WHERE status IN ('OPEN', 'UNDER_REVIEW');

COMMENT ON TABLE review_appeals IS 'Review appeals - maps to internal/domain/appeal/entity.go';

-- Domain: internal/domain/evidence/
-- Entity: evidence/entity.go
-- =========================================

CREATE TABLE review_evidence (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Association
    review_id UUID,
    flag_id UUID,
    appeal_id UUID,
    
    -- Evidence Type
    evidence_type VARCHAR(50) CHECK (
        evidence_type IN ('DOCUMENT', 'SCREENSHOT', 'CHAT_LOG', 'VIDEO', 'AUDIO', 'OTHER')
    ),
    
    -- File Reference
    file_id UUID NOT NULL,                           -- References storage-be
    file_name VARCHAR(500),
    file_size_bytes BIGINT,
    mime_type VARCHAR(100),
    
    -- Content
    description TEXT,
    notes TEXT,
    
    -- Uploader
    uploaded_by UUID NOT NULL,
    
    -- Status
    is_verified BOOLEAN DEFAULT FALSE,
    verified_by UUID,
    verified_at TIMESTAMPTZ,
    
    -- Soft Delete
    is_deleted BOOLEAN DEFAULT FALSE,
    deleted_at TIMESTAMPTZ,
    deleted_by UUID,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_review_evidence_review FOREIGN KEY (review_id) 
        REFERENCES reviews(id) ON DELETE CASCADE,
    CONSTRAINT fk_review_evidence_flag FOREIGN KEY (flag_id) 
        REFERENCES review_flags(id) ON DELETE CASCADE,
    CONSTRAINT fk_review_evidence_appeal FOREIGN KEY (appeal_id) 
        REFERENCES review_appeals(id) ON DELETE CASCADE,
    CONSTRAINT chk_review_evidence_association CHECK (
        (review_id IS NOT NULL) OR (flag_id IS NOT NULL) OR (appeal_id IS NOT NULL)
    )
);

CREATE INDEX idx_review_evidence_review ON review_evidence (review_id) WHERE review_id IS NOT NULL;
CREATE INDEX idx_review_evidence_flag ON review_evidence (flag_id) WHERE flag_id IS NOT NULL;
CREATE INDEX idx_review_evidence_appeal ON review_evidence (appeal_id) WHERE appeal_id IS NOT NULL;
CREATE INDEX idx_review_evidence_uploader ON review_evidence (uploaded_by);

COMMENT ON TABLE review_evidence IS 'Evidence attachments - maps to internal/domain/evidence/entity.go';

-- Domain: internal/domain/redaction/
-- Entity: redaction/entity.go
-- =========================================

CREATE TABLE review_redactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Target
    review_id UUID NOT NULL,
    
    -- Redaction Details
    redaction_type VARCHAR(50) CHECK (
        redaction_type IN ('PII', 'PROFANITY', 'COPYRIGHT', 'SENSITIVE', 'MANUAL')
    ),
    
    -- What Was Redacted
    field_name VARCHAR(100),                         -- 'body', 'title', 'response'
    original_text TEXT,                              -- Encrypted
    redacted_text TEXT,
    
    -- Detection
    detected_by VARCHAR(30),                         -- 'SYSTEM', 'MODERATOR'
    detection_confidence DECIMAL(5, 4),
    
    -- Pattern Matched
    pattern_matched TEXT,
    
    -- Status
    status VARCHAR(20) DEFAULT 'APPLIED' CHECK (
        status IN ('PENDING', 'APPLIED', 'ROLLED_BACK')
    ),
    
    -- Rollback
    rolled_back_at TIMESTAMPTZ,
    rolled_back_by UUID,
    rollback_reason TEXT,
    
    -- Audit
    applied_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    applied_by UUID,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_review_redactions_review FOREIGN KEY (review_id) 
        REFERENCES reviews(id) ON DELETE CASCADE
);

CREATE INDEX idx_review_redactions_review ON review_redactions (review_id);
CREATE INDEX idx_review_redactions_type ON review_redactions (redaction_type);
CREATE INDEX idx_review_redactions_status ON review_redactions (status);

COMMENT ON TABLE review_redactions IS 'Review redactions - maps to internal/domain/redaction/entity.go';

-- Redaction Policies
CREATE TABLE review_redaction_policies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Policy Identity
    policy_name VARCHAR(200) NOT NULL UNIQUE,
    
    -- Rules
    patterns JSONB NOT NULL,                         -- Array of regex patterns
    replacement_text VARCHAR(100) DEFAULT '[REDACTED]',
    
    -- Scope
    applies_to_fields TEXT[],                        -- ['body', 'title', 'response']
    
    -- Detection
    auto_detect BOOLEAN DEFAULT TRUE,
    require_manual_review BOOLEAN DEFAULT FALSE,
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    
    -- Audit
    created_by UUID,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_review_redaction_policies_active ON review_redaction_policies (is_active) 
    WHERE is_active = TRUE;

-- Domain: internal/domain/compliance/
-- Entity: compliance/entity.go
-- =========================================

CREATE TABLE review_compliance (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Request Identity
    request_id VARCHAR(100) UNIQUE NOT NULL,
    request_type VARCHAR(30) NOT NULL CHECK (
        request_type IN ('ERASURE', 'EXPORT', 'ANONYMIZATION')
    ),
    
    -- Subject
    subject_id UUID NOT NULL,
    subject_type VARCHAR(30),                        -- 'USER', 'CONTRACT'
    
    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'IN_PROGRESS', 'COMPLETED', 'FAILED')
    ),
    
    -- Processing
    entities_affected JSONB,                         -- Reviews, ratings, badges, etc.
    total_entities INTEGER,
    processed_entities INTEGER DEFAULT 0,
    
    -- Results
    reviews_affected INTEGER DEFAULT 0,
    reviews_deleted INTEGER DEFAULT 0,
    reviews_anonymized INTEGER DEFAULT 0,
    ratings_deleted INTEGER DEFAULT 0,
    
    -- Export (for data export requests)
    export_file_url TEXT,
    export_format VARCHAR(20),                       -- 'JSON', 'PDF'
    export_expires_at TIMESTAMPTZ,
    
    -- Performance
    started_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    duration_seconds INTEGER,
    
    -- Error Handling
    error_message TEXT,
    
    -- Audit
    requested_by UUID,
    requested_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_review_compliance_request ON review_compliance (request_id);
CREATE INDEX idx_review_compliance_subject ON review_compliance (subject_id, subject_type);
CREATE INDEX idx_review_compliance_type ON review_compliance (request_type);
CREATE INDEX idx_review_compliance_status ON review_compliance (status);
CREATE INDEX idx_review_compliance_requested ON review_compliance (requested_at DESC);

COMMENT ON TABLE review_compliance IS 'Compliance requests - maps to internal/domain/compliance/entity.go';

-- Compliance Actions Log
CREATE TABLE review_compliance_actions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    compliance_request_id UUID NOT NULL,
    
    -- Action Details
    action_type VARCHAR(50) NOT NULL,                -- 'DELETE', 'ANONYMIZE', 'EXPORT'
    entity_type VARCHAR(30) NOT NULL,
    entity_id UUID NOT NULL,
    
    -- Result
    status VARCHAR(20) CHECK (
        status IN ('SUCCESS', 'FAILED', 'SKIPPED')
    ),
    error_message TEXT,
    
    performed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_review_compliance_actions_request FOREIGN KEY (compliance_request_id) 
        REFERENCES review_compliance(id) ON DELETE CASCADE
);

CREATE INDEX idx_review_compliance_actions_request ON review_compliance_actions (compliance_request_id);
CREATE INDEX idx_review_compliance_actions_entity ON review_compliance_actions (entity_type, entity_id);

-- Domain: internal/domain/audit_trail/
-- Entity: audit_trail/entity.go (implied)
-- =========================================

CREATE TABLE review_audit_trail (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Event Identity
    event_type VARCHAR(100) NOT NULL,
    
    -- Target
    entity_type VARCHAR(30) NOT NULL,
    entity_id UUID NOT NULL,
    
    -- Actor
    actor_id UUID,
    actor_type VARCHAR(30),                          -- 'USER', 'SYSTEM', 'ADMIN'
    
    -- Action
    action VARCHAR(100) NOT NULL,
    
    -- Changes
    previous_state JSONB,
    new_state JSONB,
    changes JSONB,                                   -- Diff of changes
    
    -- Context
    context JSONB,
    
    -- Immutability
    hash VARCHAR(64) NOT NULL,                       -- SHA256 of record + previous hash
    previous_hash VARCHAR(64),
    
    -- Timestamps
    occurred_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_review_audit_trail_entity ON review_audit_trail (entity_type, entity_id, occurred_at DESC);
CREATE INDEX idx_review_audit_trail_actor ON review_audit_trail (actor_id, occurred_at DESC);
CREATE INDEX idx_review_audit_trail_event ON review_audit_trail (event_type);
CREATE INDEX idx_review_audit_trail_occurred ON review_audit_trail (occurred_at DESC);

COMMENT ON TABLE review_audit_trail IS 'Immutable audit trail';

```

---

=========================================
## SECTION 8: STATS & EXPOSURE
=========================================

```sql
-- Domain: internal/domain/stats/
-- Entity: stats/entity.go (implied)
-- =========================================

CREATE TABLE review_stats (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- User
    user_id UUID NOT NULL UNIQUE,
    user_type VARCHAR(30) NOT NULL CHECK (
        user_type IN ('FREELANCER', 'CLIENT')
    ),
    
    -- Review Counts
    total_reviews_received INTEGER DEFAULT 0,
    total_reviews_given INTEGER DEFAULT 0,
    
    published_reviews_received INTEGER DEFAULT 0,
    published_reviews_given INTEGER DEFAULT 0,
    
    -- Rating Statistics
    average_rating_received DECIMAL(3, 2),
    average_rating_given DECIMAL(3, 2),
    
    -- Recent Activity (90 days)
    recent_reviews_received_90d INTEGER DEFAULT 0,
    recent_reviews_given_90d INTEGER DEFAULT 0,
    recent_average_rating_90d DECIMAL(3, 2),
    
    -- Response Statistics
    responses_given INTEGER DEFAULT 0,
    response_rate DECIMAL(5, 2),                     -- Percentage
    
    -- Engagement
    total_helpful_votes_received INTEGER DEFAULT 0,
    total_views_received INTEGER DEFAULT 0,
    
    -- Reputation Reference
    reputation_score DECIMAL(10, 4),
    reputation_percentile DECIMAL(5, 2),
    
    -- Badges Count
    active_badges_count INTEGER DEFAULT 0,
    
    -- Cache Control
    cache_key VARCHAR(100),
    cache_expires_at TIMESTAMPTZ,
    
    -- Last Updated
    last_review_received_at TIMESTAMPTZ,
    last_review_given_at TIMESTAMPTZ,
    recalculated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_review_stats_user ON review_stats (user_id);
CREATE INDEX idx_review_stats_type ON review_stats (user_type);
CREATE INDEX idx_review_stats_average ON review_stats (average_rating_received DESC);
CREATE INDEX idx_review_stats_recent ON review_stats (recent_reviews_received_90d DESC);
CREATE INDEX idx_review_stats_cache ON review_stats (cache_key, cache_expires_at);

COMMENT ON TABLE review_stats IS 'User review statistics';

-- Domain: internal/domain/featured/
-- Entity: featured/entity.go (implied)
-- =========================================

CREATE TABLE featured_reviews (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Review
    review_id UUID NOT NULL UNIQUE,
    
    -- Featured Configuration
    featured_reason TEXT,
    featured_category VARCHAR(50),                   -- 'QUALITY', 'HELPFUL', 'DETAILED'
    
    -- Duration
    duration_days INTEGER DEFAULT 30,
    
    -- Placement
    placement VARCHAR(50),                           -- 'HOME', 'CATEGORY', 'PROFILE'
    display_order INTEGER,
    
    -- Status
    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (
        status IN ('ACTIVE', 'EXPIRED', 'REMOVED')
    ),
    
    -- Timeline
    featured_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at TIMESTAMPTZ,
    expired_at TIMESTAMPTZ,
    
    -- Actor
    featured_by UUID NOT NULL,
    removed_by UUID,
    removal_reason TEXT,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_featured_reviews_review FOREIGN KEY (review_id) 
        REFERENCES reviews(id) ON DELETE CASCADE
);

CREATE INDEX idx_featured_reviews_review ON featured_reviews (review_id);
CREATE INDEX idx_featured_reviews_status ON featured_reviews (status);
CREATE INDEX idx_featured_reviews_active ON featured_reviews (status, display_order, featured_at DESC) 
    WHERE status = 'ACTIVE';
CREATE INDEX idx_featured_reviews_expires ON featured_reviews (expires_at) 
    WHERE status = 'ACTIVE' AND expires_at IS NOT NULL;
CREATE INDEX idx_featured_reviews_placement ON featured_reviews (placement, status);

COMMENT ON TABLE featured_reviews IS 'Featured reviews';

```

---

=========================================
## SECTION 9: OUTBOX & EVENTS
=========================================

```sql
-- Outbox Pattern for Event Publishing
CREATE TABLE outbox_events (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Event Identity
    event_id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    event_type VARCHAR(100) NOT NULL,
    event_version VARCHAR(10) DEFAULT 'v1',
    
    -- Aggregate
    aggregate_id UUID NOT NULL,
    aggregate_type VARCHAR(50) NOT NULL,
    
    -- Payload
    payload JSONB NOT NULL,
    
    -- Metadata
    correlation_id UUID,
    causation_id UUID,
    
    -- Actor
    actor_id UUID,
    actor_type VARCHAR(30),
    
    -- Partition
    partition_key VARCHAR(200) NOT NULL,
    
    -- Publishing Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'PUBLISHED', 'FAILED')
    ),
    
    -- Attempts
    publish_attempts INTEGER DEFAULT 0,
    max_attempts INTEGER DEFAULT 3,
    last_attempt_at TIMESTAMPTZ,
    next_attempt_at TIMESTAMPTZ,
    
    -- Error
    error_message TEXT,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    published_at TIMESTAMPTZ
);

CREATE INDEX idx_outbox_events_status ON outbox_events (status, next_attempt_at) 
    WHERE status = 'PENDING';
CREATE INDEX idx_outbox_events_aggregate ON outbox_events (aggregate_id, aggregate_type);
CREATE INDEX idx_outbox_events_created ON outbox_events (created_at DESC);
CREATE INDEX idx_outbox_events_partition ON outbox_events (partition_key);

COMMENT ON TABLE outbox_events IS 'Outbox pattern for reliable event publishing';

-- Dead Letter Queue
CREATE TABLE outbox_dead_letter (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Original Event
    original_event_id UUID NOT NULL,
    event_type VARCHAR(100) NOT NULL,
    event_payload JSONB NOT NULL,
    
    -- Failure Details
    failure_reason TEXT,
    failure_count INTEGER,
    last_error TEXT,
    
    -- Context
    correlation_id UUID,
    
    -- Timestamps
    failed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_outbox_dead_letter_event ON outbox_dead_letter (original_event_id);
CREATE INDEX idx_outbox_dead_letter_type ON outbox_dead_letter (event_type);
CREATE INDEX idx_outbox_dead_letter_failed ON outbox_dead_letter (failed_at DESC);

```

---

=========================================
## DATABASE VIEWS
=========================================

```sql
-- Active Reviews Summary
CREATE VIEW v_active_reviews AS
SELECT 
    r.id,
    r.review_number,
    r.contract_id,
    r.reviewer_id,
    r.reviewee_id,
    r.review_type,
    r.overall_rating,
    r.status,
    r.is_featured,
    r.helpful_votes,
    r.view_count,
    r.published_at
FROM reviews r
WHERE r.status = 'PUBLISHED' 
    AND r.is_deleted = FALSE
ORDER BY r.published_at DESC;

-- User Review Performance
CREATE VIEW v_user_review_performance AS
SELECT 
    rs.user_id,
    rs.user_type,
    rs.average_rating_received,
    rs.total_reviews_received,
    rs.recent_reviews_received_90d,
    ra.wilson_score,
    rep.reputation_score,
    rep.is_top_rated,
    rs.active_badges_count
FROM review_stats rs
LEFT JOIN rating_aggregates ra ON rs.user_id = ra.user_id
LEFT JOIN reputation rep ON rs.user_id = rep.user_id
ORDER BY rep.reputation_score DESC;

-- Moderation Queue
CREATE VIEW v_moderation_queue AS
SELECT 
    r.id AS review_id,
    r.review_number,
    r.reviewer_id,
    r.reviewee_id,
    rm.moderation_state,
    rm.auto_flagged,
    rm.requires_manual_review,
    rf.status AS flag_status,
    rf.flag_reason,
    rf.priority,
    rf.created_at AS flagged_at
FROM reviews r
INNER JOIN review_moderation rm ON r.id = rm.review_id
LEFT JOIN review_flags rf ON r.id = rf.review_id
WHERE rm.moderation_state IN ('PENDING', 'MANUAL_REVIEW', 'REJECTED')
    OR rf.status IN ('OPEN', 'UNDER_REVIEW')
ORDER BY 
    CASE rf.priority 
        WHEN 'CRITICAL' THEN 1 
        WHEN 'HIGH' THEN 2 
        WHEN 'NORMAL' THEN 3 
        ELSE 4 
    END,
    rf.created_at ASC;

-- Top Rated Users
CREATE VIEW v_top_rated_users AS
SELECT 
    r.user_id,
    r.user_type,
    r.reputation_score,
    r.normalized_score,
    r.percentile,
    r.is_top_rated,
    r.total_reviews,
    ra.average_rating,
    ra.wilson_score
FROM reputation r
INNER JOIN rating_aggregates ra ON r.user_id = ra.user_id
WHERE r.is_top_rated = TRUE
ORDER BY r.reputation_score DESC;

-- Featured Reviews Current
CREATE VIEW v_featured_reviews_current AS
SELECT 
    fr.id,
    fr.review_id,
    r.review_number,
    r.overall_rating,
    r.helpful_votes,
    r.reviewee_id,
    fr.featured_category,
    fr.placement,
    fr.featured_at,
    fr.expires_at
FROM featured_reviews fr
INNER JOIN reviews r ON fr.review_id = r.id
WHERE fr.status = 'ACTIVE'
    AND (fr.expires_at IS NULL OR fr.expires_at > CURRENT_TIMESTAMP)
ORDER BY fr.display_order, fr.featured_at DESC;

-- Double-Blind Window Status
CREATE VIEW v_double_blind_window_status AS
SELECT 
    dbw.id,
    dbw.contract_id,
    dbw.status,
    dbw.opens_at,
    dbw.closes_at,
    dbw.client_submitted,
    dbw.freelancer_submitted,
    dbw.both_submitted,
    dbw.published_at,
    CASE 
        WHEN dbw.both_submitted THEN 'BOTH_SUBMITTED'
        WHEN dbw.client_submitted THEN 'CLIENT_ONLY'
        WHEN dbw.freelancer_submitted THEN 'FREELANCER_ONLY'
        ELSE 'NONE'
    END AS submission_status,
    CASE 
        WHEN CURRENT_TIMESTAMP < dbw.opens_at THEN 'SCHEDULED'
        WHEN CURRENT_TIMESTAMP > dbw.closes_at THEN 'EXPIRED'
        ELSE 'OPEN'
    END AS window_state
FROM double_blind_windows dbw
ORDER BY dbw.closes_at DESC;

-- Badge Leaderboard
CREATE VIEW v_badge_leaderboard AS
SELECT 
    ub.user_id,
    COUNT(DISTINCT ub.badge_id) AS total_badges,
    COUNT(DISTINCT ub.badge_id) FILTER (WHERE ub.badge_level = 'PLATINUM') AS platinum_count,
    COUNT(DISTINCT ub.badge_id) FILTER (WHERE ub.badge_level = 'GOLD') AS gold_count,
    COUNT(DISTINCT ub.badge_id) FILTER (WHERE ub.badge_level = 'SILVER') AS silver_count,
    COUNT(DISTINCT ub.badge_id) FILTER (WHERE ub.badge_level = 'BRONZE') AS bronze_count,
    MAX(ub.achieved_at) AS latest_achievement
FROM user_badges ub
WHERE ub.status = 'ACTIVE' AND ub.is_visible = TRUE
GROUP BY ub.user_id
ORDER BY total_badges DESC, platinum_count DESC, gold_count DESC;

-- Database Health Overview
CREATE VIEW v_database_health AS
SELECT
    'Total Reviews' AS metric,
    COUNT(*) AS count
FROM reviews
WHERE is_deleted = FALSE
UNION ALL
SELECT
    'Published Reviews',
    COUNT(*)
FROM reviews
WHERE status = 'PUBLISHED' AND is_deleted = FALSE
UNION ALL
SELECT
    'Pending Moderation',
    COUNT(*)
FROM review_moderation
WHERE moderation_state IN ('PENDING', 'MANUAL_REVIEW')
UNION ALL
SELECT
    'Open Flags',
    COUNT(*)
FROM review_flags
WHERE status IN ('OPEN', 'UNDER_REVIEW')
UNION ALL
SELECT
    'Active Appeals',
    COUNT(*)
FROM review_appeals
WHERE status IN ('OPEN', 'UNDER_REVIEW')
UNION ALL
SELECT
    'Open Double-Blind Windows',
    COUNT(*)
FROM double_blind_windows
WHERE status = 'OPEN'
UNION ALL
SELECT
    'Active Badges Awarded',
    COUNT(*)
FROM user_badges
WHERE status = 'ACTIVE'
UNION ALL
SELECT
    'Pending Outbox Events',
    COUNT(*)
FROM outbox_events
WHERE status = 'PENDING';

-- Table Size Statistics
CREATE VIEW v_table_sizes AS
SELECT
    schemaname,
    tablename,
    pg_size_pretty(pg_total_relation_size(schemaname||'.'||tablename)) AS size,
    pg_total_relation_size(schemaname||'.'||tablename) AS size_bytes
FROM pg_tables
WHERE schemaname = 'public'
ORDER BY pg_total_relation_size(schemaname||'.'||tablename) DESC;

```

---

=========================================
## TABLE COMMENTS SUMMARY
=========================================

```sql
-- Core Review Primitives
COMMENT ON TABLE reviews IS 'Reviews - maps to internal/domain/review/entity.go';
COMMENT ON TABLE review_ratings IS 'Review ratings per dimension';
COMMENT ON TABLE review_responses IS 'Review responses from reviewees';
COMMENT ON TABLE review_helpful_votes IS 'Helpful votes tracking';
COMMENT ON TABLE review_drafts IS 'Review drafts - maps to internal/domain/review_draft/entity.go';
COMMENT ON TABLE review_eligibility IS 'Review eligibility checks - maps to internal/domain/eligibility/entity.go';
COMMENT ON TABLE review_eligibility_policies IS 'Eligibility policy configuration';

-- Rating & Criteria
COMMENT ON TABLE rating_criteria IS 'Rating criteria versions - maps to internal/domain/rating/entity.go';
COMMENT ON TABLE rating_aggregates IS 'Rating aggregates per user';

-- Reputation
COMMENT ON TABLE reputation IS 'User reputation scores - maps to internal/domain/reputation/entity.go';
COMMENT ON TABLE reputation_history IS 'Reputation history time series';
COMMENT ON TABLE reputation_top_rated_rules IS 'Top rated eligibility rules';

-- Badge System
COMMENT ON TABLE badges IS 'Badge definitions - maps to internal/domain/badge/entity.go';
COMMENT ON TABLE user_badges IS 'User badge assignments - maps to internal/domain/user_badge/entity.go';
COMMENT ON TABLE badge_eligibility_checks IS 'Badge eligibility check history';

-- Double-Blind & Windows
COMMENT ON TABLE double_blind_windows IS 'Double-blind review windows - maps to internal/domain/double_blind/entity.go';
COMMENT ON TABLE review_reminders IS 'Review reminders - maps to internal/domain/review_reminder/entity.go';

-- Private Feedback
COMMENT ON TABLE private_feedback IS 'Private feedback - maps to internal/domain/feedback/entity.go';

-- Safety & Governance
COMMENT ON TABLE review_flags IS 'Review flags - maps to internal/domain/flag/entity.go';
COMMENT ON TABLE review_moderation IS 'Review moderation - maps to internal/domain/moderation/entity.go';
COMMENT ON TABLE review_moderation_history IS 'Moderation state history';
COMMENT ON TABLE review_appeals IS 'Review appeals - maps to internal/domain/appeal/entity.go';
COMMENT ON TABLE review_evidence IS 'Evidence attachments - maps to internal/domain/evidence/entity.go';
COMMENT ON TABLE review_redactions IS 'Review redactions - maps to internal/domain/redaction/entity.go';
COMMENT ON TABLE review_redaction_policies IS 'Redaction policy rules';
COMMENT ON TABLE review_compliance IS 'Compliance requests - maps to internal/domain/compliance/entity.go';
COMMENT ON TABLE review_compliance_actions IS 'Compliance action log';
COMMENT ON TABLE review_audit_trail IS 'Immutable audit trail';

-- Stats & Exposure
COMMENT ON TABLE review_stats IS 'User review statistics';
COMMENT ON TABLE featured_reviews IS 'Featured reviews';

-- Outbox & Events
COMMENT ON TABLE outbox_events IS 'Outbox pattern for reliable event publishing';
COMMENT ON TABLE outbox_dead_letter IS 'Dead letter queue for failed events';

```

---

=========================================
## END OF REVIEWS-BE DATABASE DESIGN
=========================================

## FINAL SUMMARY:

**Total Tables:** 45+
**Total Indexes:** 200+
**Total Domains Covered:** 15 (all from reviews-be folder structure)
**Coverage:** 100% of reviews-be domain structure

### Key Features:
- ✅ Full event sourcing with outbox pattern
- ✅ CQRS with read models
- ✅ Double-blind review system
- ✅ Multi-dimensional rating criteria with versioning
- ✅ Advanced reputation algorithm (Bayesian + decay + recency)
- ✅ Badge automation with eligibility checks
- ✅ Comprehensive safety & moderation system
- ✅ PII redaction and compliance (GDPR)
- ✅ Appeal and evidence management
- ✅ Private feedback channel
- ✅ Helpful votes with 24h change window
- ✅ Featured reviews management
- ✅ Complete audit trails (immutable with hash chain)
- ✅ Production-ready for millions of reviews

### Alignment with Folder Structure:
- ✅ review/ → reviews, review_ratings, review_responses, review_helpful_votes tables
- ✅ review_draft/ → review_drafts table
- ✅ eligibility/ → review_eligibility, review_eligibility_policies tables
- ✅ rating/ → rating_criteria, rating_aggregates tables
- ✅ reputation/ → reputation, reputation_history, reputation_top_rated_rules tables
- ✅ badge/ → badges table
- ✅ user_badge/ → user_badges, badge_eligibility_checks tables
- ✅ double_blind/ → double_blind_windows table
- ✅ review_reminder/ → review_reminders table
- ✅ feedback/ → private_feedback table
- ✅ flag/ → review_flags table
- ✅ moderation/ → review_moderation, review_moderation_history tables
- ✅ appeal/ → review_appeals table
- ✅ evidence/ → review_evidence table
- ✅ redaction/ → review_redactions, review_redaction_policies tables
- ✅ compliance/ → review_compliance, review_compliance_actions tables
- ✅ audit_trail/ → review_audit_trail table
- ✅ stats/ → review_stats table
- ✅ featured/ → featured_reviews table

All domains from reviews-be folder structure are fully covered!
