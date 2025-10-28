# ADMIN-BE DATABASE DESIGN
**Skillsier Platform – Enterprise Scale (Upwork-like)**  
**PostgreSQL 16+**

---

## **CRITICAL ALIGNMENT RULES**
1. Each domain folder in `internal/domain/{domain}/` = ONE main table
2. Table names follow domain folder names; when aggregated under admin, tables are prefixed with `admin_` to reflect the domain boundary
3. Sub-entities within domain create related tables with `{domain}_{sub}` naming
4. All domains from folder structure are covered
5. Rich, production-ready fields for large-scale application

---

## **Global Extensions**

```sql
-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";
CREATE EXTENSION IF NOT EXISTS "btree_gin";
CREATE EXTENSION IF NOT EXISTS "citext";
CREATE EXTENSION IF NOT EXISTS "btree_gist";
```

---

## **SECTION 1: CORE ADMIN**

### Domain: admin_user/

```sql
-- =========================================
-- MAIN TABLE: admin_users
-- Domain: internal/domain/admin_user/
-- Entity: admin_user/entity.go
-- =========================================

CREATE TABLE admin_users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Identity
    keycloak_id UUID NOT NULL UNIQUE,
    email CITEXT NOT NULL UNIQUE,
    full_name VARCHAR(255) NOT NULL,
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('ACTIVE', 'INACTIVE', 'SUSPENDED', 'DELETED')
    ) DEFAULT 'ACTIVE',
    
    -- Roles (bitset stored as BIGINT for performance)
    roles BIGINT NOT NULL DEFAULT 0,
    -- Role bitset:
    -- 1 = SUPER_ADMIN
    -- 2 = MODERATOR
    -- 4 = SUPPORT
    -- 8 = COMPLIANCE
    -- 16 = FINANCE_ADMIN
    -- 32 = OPS_MANAGER
    -- 64 = CONTENT_MANAGER
    
    -- Permissions (bitset stored as BIGINT)
    permissions BIGINT NOT NULL DEFAULT 0,
    -- Permission bitset defined per role requirements
    
    -- Multi-Factor Authentication
    mfa_enabled BOOLEAN DEFAULT FALSE,
    mfa_secret_encrypted TEXT,
    mfa_backup_codes_encrypted TEXT,
    mfa_last_used_at TIMESTAMPTZ,
    
    -- Security
    password_hash TEXT NOT NULL,
    password_changed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    password_must_change BOOLEAN DEFAULT FALSE,
    failed_login_attempts INTEGER DEFAULT 0,
    locked_until TIMESTAMPTZ,
    
    -- Session Management
    last_login_at TIMESTAMPTZ,
    last_login_ip INET,
    last_login_user_agent TEXT,
    current_session_id UUID,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by UUID,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_by UUID,
    deleted_at TIMESTAMPTZ,
    
    -- Metadata
    preferences JSONB DEFAULT '{}'::JSONB,
    metadata JSONB DEFAULT '{}'::JSONB,
    
    -- Indexes for performance
    CONSTRAINT fk_admin_users_created_by FOREIGN KEY (created_by) REFERENCES admin_users(id),
    CONSTRAINT fk_admin_users_updated_by FOREIGN KEY (updated_by) REFERENCES admin_users(id)
);

CREATE INDEX idx_admin_users_keycloak ON admin_users (keycloak_id);
CREATE INDEX idx_admin_users_email ON admin_users USING gin (email gin_trgm_ops);
CREATE INDEX idx_admin_users_status ON admin_users (status) WHERE status != 'DELETED';
CREATE INDEX idx_admin_users_roles ON admin_users (roles);
CREATE INDEX idx_admin_users_last_login ON admin_users (last_login_at DESC);
CREATE INDEX idx_admin_users_created_at ON admin_users (created_at DESC);

COMMENT ON TABLE admin_users IS 'Admin users - maps to internal/domain/admin_user/entity.go';
COMMENT ON COLUMN admin_users.roles IS 'Bitset: 1=SUPER_ADMIN, 2=MODERATOR, 4=SUPPORT, 8=COMPLIANCE, 16=FINANCE_ADMIN, 32=OPS_MANAGER, 64=CONTENT_MANAGER';
COMMENT ON COLUMN admin_users.permissions IS 'Granular permission bitset per admin user';

-- =========================================
-- SUB-ENTITY: admin_user_roles
-- Domain: internal/domain/admin_user/
-- Entity: admin_user/role.go
-- =========================================

CREATE TABLE admin_user_roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Role Definition
    role_name VARCHAR(50) NOT NULL UNIQUE,
    role_code VARCHAR(20) NOT NULL UNIQUE,
    role_bit INTEGER NOT NULL UNIQUE, -- Bit position (0-63)
    
    -- Hierarchy
    hierarchy_level INTEGER NOT NULL,
    parent_role_id UUID,
    
    -- Permissions
    default_permissions BIGINT NOT NULL DEFAULT 0,
    
    -- Description
    display_name VARCHAR(100) NOT NULL,
    description TEXT,
    
    -- Restrictions
    restrictions JSONB DEFAULT '{}'::JSONB,
    can_assign_roles BIGINT DEFAULT 0, -- Bitset of roles this role can assign
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by UUID REFERENCES admin_users(id),
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_by UUID REFERENCES admin_users(id),
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    
    CONSTRAINT fk_admin_user_roles_parent FOREIGN KEY (parent_role_id) REFERENCES admin_user_roles(id),
    CONSTRAINT chk_admin_user_roles_bit_range CHECK (role_bit >= 0 AND role_bit < 64)
);

CREATE INDEX idx_admin_user_roles_active ON admin_user_roles (is_active) WHERE is_active = TRUE;
CREATE INDEX idx_admin_user_roles_hierarchy ON admin_user_roles (hierarchy_level);

COMMENT ON TABLE admin_user_roles IS 'Role definitions - maps to internal/domain/admin_user/role.go';

-- =========================================
-- SUB-ENTITY: admin_activity_logs
-- Domain: internal/domain/admin_user/
-- Entity: admin_user/activity_log.go
-- =========================================

CREATE TABLE admin_activity_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Actor
    admin_user_id UUID NOT NULL REFERENCES admin_users(id),
    admin_email VARCHAR(255) NOT NULL,
    admin_role VARCHAR(50) NOT NULL,
    
    -- Action
    action_type VARCHAR(100) NOT NULL,
    action_category VARCHAR(50) NOT NULL CHECK (
        action_category IN ('USER_ACTION', 'CONTENT_MODERATION', 'FINANCIAL', 'SYSTEM_CONFIG', 'SUPPORT', 'COMPLIANCE', 'SECURITY')
    ),
    action_description TEXT,
    
    -- Target
    resource_type VARCHAR(50),
    resource_id UUID,
    resource_identifier VARCHAR(255),
    
    -- Context
    correlation_id UUID,
    causation_id UUID,
    session_id UUID,
    request_id UUID,
    
    -- Request Details
    ip_address INET NOT NULL,
    user_agent TEXT,
    http_method VARCHAR(10),
    http_path VARCHAR(500),
    http_status_code INTEGER,
    
    -- Outcome
    success BOOLEAN NOT NULL,
    error_message TEXT,
    error_code VARCHAR(50),
    
    -- Changes (for audit)
    changes_before JSONB,
    changes_after JSONB,
    changes_diff JSONB,
    
    -- Compliance
    compliance_required BOOLEAN DEFAULT TRUE,
    data_classification VARCHAR(20) CHECK (
        data_classification IN ('PUBLIC', 'INTERNAL', 'SENSITIVE', 'CONFIDENTIAL')
    ),
    retention_years INTEGER DEFAULT 10,
    
    -- Timestamp
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

-- Partitioning by month for performance
CREATE INDEX idx_admin_activity_logs_admin ON admin_activity_logs (admin_user_id, created_at DESC);
CREATE INDEX idx_admin_activity_logs_action ON admin_activity_logs (action_type, created_at DESC);
CREATE INDEX idx_admin_activity_logs_resource ON admin_activity_logs (resource_type, resource_id);
CREATE INDEX idx_admin_activity_logs_created ON admin_activity_logs (created_at DESC);
CREATE INDEX idx_admin_activity_logs_correlation ON admin_activity_logs (correlation_id);
CREATE INDEX idx_admin_activity_logs_session ON admin_activity_logs (session_id);
CREATE INDEX idx_admin_activity_logs_ip ON admin_activity_logs (ip_address);

COMMENT ON TABLE admin_activity_logs IS 'Immutable audit trail - maps to internal/domain/admin_user/activity_log.go';
COMMENT ON COLUMN admin_activity_logs.retention_years IS 'Compliance retention: 10 years default';
```

---

## **SECTION 2: SUPPORT & CASEWORK**

### Domain: support_ticket/

```sql
-- =========================================
-- MAIN TABLE: support_tickets
-- Domain: internal/domain/support_ticket/
-- Entity: support_ticket/entity.go
-- =========================================

CREATE TABLE support_tickets (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Ticket Identification
    ticket_number VARCHAR(20) NOT NULL UNIQUE, -- e.g., TICK-2025-001234
    
    -- Requester
    requester_user_id UUID NOT NULL,
    requester_email CITEXT NOT NULL,
    requester_name VARCHAR(255) NOT NULL,
    requester_type VARCHAR(20) NOT NULL CHECK (
        requester_type IN ('CLIENT', 'FREELANCER', 'ADMIN', 'GUEST')
    ),
    
    -- Ticket Details
    subject VARCHAR(500) NOT NULL,
    description TEXT NOT NULL,
    
    -- Categorization
    category VARCHAR(50) NOT NULL CHECK (
        category IN ('BILLING', 'TECHNICAL', 'ACCOUNT', 'KYC', 'ABUSE', 'REFUND', 'DISPUTE', 'GENERAL', 'FEATURE_REQUEST', 'BUG_REPORT')
    ),
    sub_category VARCHAR(100),
    tags TEXT[], -- Array of tags
    
    -- Priority & Severity
    priority VARCHAR(20) NOT NULL CHECK (
        priority IN ('LOW', 'MEDIUM', 'HIGH', 'URGENT', 'CRITICAL')
    ) DEFAULT 'MEDIUM',
    severity_score INTEGER DEFAULT 50 CHECK (severity_score >= 0 AND severity_score <= 100),
    
    -- Status Management
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('OPEN', 'ASSIGNED', 'IN_PROGRESS', 'AWAITING_USER', 'AWAITING_INTERNAL', 'RESOLVED', 'CLOSED', 'REOPENED')
    ) DEFAULT 'OPEN',
    
    -- Assignment
    assigned_agent_id UUID REFERENCES admin_users(id),
    assigned_queue VARCHAR(50),
    assigned_at TIMESTAMPTZ,
    
    -- SLA Management
    sla_target_response_at TIMESTAMPTZ,
    sla_first_response_at TIMESTAMPTZ,
    sla_target_resolution_at TIMESTAMPTZ,
    sla_resolved_at TIMESTAMPTZ,
    sla_breached BOOLEAN DEFAULT FALSE,
    sla_breach_reason TEXT,
    
    -- Resolution
    resolution_type VARCHAR(50),
    resolution_notes TEXT,
    resolved_by UUID REFERENCES admin_users(id),
    
    -- Escalation
    escalation_level INTEGER DEFAULT 0,
    escalated_at TIMESTAMPTZ,
    escalated_by UUID REFERENCES admin_users(id),
    escalation_reason TEXT,
    
    -- Channel
    channel VARCHAR(20) CHECK (
        channel IN ('EMAIL', 'WEB', 'MOBILE', 'PHONE', 'CHAT')
    ) DEFAULT 'WEB',
    
    -- Satisfaction
    csat_score INTEGER CHECK (csat_score >= 1 AND csat_score <= 5),
    csat_feedback TEXT,
    csat_submitted_at TIMESTAMPTZ,
    
    -- Related Entities
    related_user_id UUID,
    related_job_id UUID,
    related_proposal_id UUID,
    related_contract_id UUID,
    related_payment_id UUID,
    related_dispute_id UUID,
    
    -- Timestamps
    opened_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    closed_at TIMESTAMPTZ,
    last_message_at TIMESTAMPTZ,
    last_updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    internal_notes TEXT
);

CREATE INDEX idx_support_tickets_number ON support_tickets (ticket_number);
CREATE INDEX idx_support_tickets_requester ON support_tickets (requester_user_id, opened_at DESC);
CREATE INDEX idx_support_tickets_assigned ON support_tickets (assigned_agent_id, status) WHERE assigned_agent_id IS NOT NULL;
CREATE INDEX idx_support_tickets_status ON support_tickets (status, priority, opened_at DESC);
CREATE INDEX idx_support_tickets_category ON support_tickets (category, sub_category);
CREATE INDEX idx_support_tickets_priority ON support_tickets (priority DESC, opened_at);
CREATE INDEX idx_support_tickets_sla_breach ON support_tickets (sla_breached) WHERE sla_breached = TRUE;
CREATE INDEX idx_support_tickets_opened_at ON support_tickets (opened_at DESC);
CREATE INDEX idx_support_tickets_tags ON support_tickets USING gin(tags);
CREATE INDEX idx_support_tickets_search ON support_tickets USING gin(to_tsvector('english', subject || ' ' || description));

COMMENT ON TABLE support_tickets IS 'Support tickets - maps to internal/domain/support_ticket/entity.go';
COMMENT ON COLUMN support_tickets.sla_target_response_at IS 'Target time for first response based on priority';
COMMENT ON COLUMN support_tickets.sla_target_resolution_at IS 'Target time for resolution based on priority';
```

### Domain: ticket_message/

```sql
-- =========================================
-- MAIN TABLE: ticket_messages
-- Domain: internal/domain/ticket_message/
-- Entity: ticket_message/entity.go
-- =========================================

CREATE TABLE ticket_messages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Ticket Reference
    ticket_id UUID NOT NULL REFERENCES support_tickets(id) ON DELETE CASCADE,
    
    -- Author
    author_id UUID NOT NULL,
    author_type VARCHAR(20) NOT NULL CHECK (
        author_type IN ('USER', 'AGENT', 'SYSTEM')
    ),
    author_name VARCHAR(255) NOT NULL,
    author_email CITEXT,
    
    -- Message Details
    body TEXT NOT NULL,
    body_format VARCHAR(20) DEFAULT 'PLAIN_TEXT' CHECK (
        body_format IN ('PLAIN_TEXT', 'MARKDOWN', 'HTML')
    ),
    
    -- Visibility
    visibility VARCHAR(20) NOT NULL CHECK (
        visibility IN ('PUBLIC', 'INTERNAL')
    ) DEFAULT 'PUBLIC',
    
    -- Type
    message_type VARCHAR(30) DEFAULT 'REPLY' CHECK (
        message_type IN ('REPLY', 'NOTE', 'STATUS_CHANGE', 'ASSIGNMENT', 'ESCALATION', 'AUTO_RESPONSE')
    ),
    
    -- Status Change (if applicable)
    old_status VARCHAR(20),
    new_status VARCHAR(20),
    
    -- Channel
    channel VARCHAR(20) CHECK (
        channel IN ('EMAIL', 'WEB', 'MOBILE', 'PHONE', 'CHAT')
    ),
    
    -- Email Integration
    email_message_id VARCHAR(255),
    email_in_reply_to VARCHAR(255),
    email_references TEXT,
    
    -- Editing
    edited BOOLEAN DEFAULT FALSE,
    edited_at TIMESTAMPTZ,
    edited_by UUID REFERENCES admin_users(id),
    original_body TEXT,
    
    -- Timestamps
    sent_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    read_at TIMESTAMPTZ,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_ticket_messages_ticket ON ticket_messages (ticket_id, sent_at DESC);
CREATE INDEX idx_ticket_messages_author ON ticket_messages (author_id, sent_at DESC);
CREATE INDEX idx_ticket_messages_visibility ON ticket_messages (ticket_id, visibility, sent_at DESC);
CREATE INDEX idx_ticket_messages_sent_at ON ticket_messages (sent_at DESC);
CREATE INDEX idx_ticket_messages_email ON ticket_messages (email_message_id) WHERE email_message_id IS NOT NULL;

COMMENT ON TABLE ticket_messages IS 'Ticket messages - maps to internal/domain/ticket_message/entity.go';

-- =========================================
-- SUB-ENTITY: ticket_message_attachments
-- Domain: internal/domain/ticket_message/
-- Entity: ticket_message/attachment.go
-- =========================================

CREATE TABLE ticket_message_attachments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Message Reference
    message_id UUID NOT NULL REFERENCES ticket_messages(id) ON DELETE CASCADE,
    ticket_id UUID NOT NULL REFERENCES support_tickets(id) ON DELETE CASCADE,
    
    -- File Details
    file_name VARCHAR(255) NOT NULL,
    file_size BIGINT NOT NULL,
    file_type VARCHAR(100) NOT NULL,
    mime_type VARCHAR(100) NOT NULL,
    
    -- Storage
    storage_path VARCHAR(500) NOT NULL,
    storage_provider VARCHAR(50) DEFAULT 'S3',
    storage_bucket VARCHAR(100),
    storage_key VARCHAR(500),
    
    -- Security
    checksum VARCHAR(64) NOT NULL, -- SHA-256
    virus_scan_status VARCHAR(20) CHECK (
        virus_scan_status IN ('PENDING', 'CLEAN', 'INFECTED', 'FAILED')
    ) DEFAULT 'PENDING',
    virus_scan_result TEXT,
    virus_scanned_at TIMESTAMPTZ,
    
    -- Access Control
    access_level VARCHAR(20) DEFAULT 'INTERNAL' CHECK (
        access_level IN ('PUBLIC', 'INTERNAL', 'RESTRICTED')
    ),
    
    -- Timestamps
    uploaded_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    uploaded_by UUID NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_ticket_message_attachments_message ON ticket_message_attachments (message_id);
CREATE INDEX idx_ticket_message_attachments_ticket ON ticket_message_attachments (ticket_id);
CREATE INDEX idx_ticket_message_attachments_virus ON ticket_message_attachments (virus_scan_status) WHERE virus_scan_status != 'CLEAN';
CREATE INDEX idx_ticket_message_attachments_checksum ON ticket_message_attachments (checksum);

COMMENT ON TABLE ticket_message_attachments IS 'Ticket message attachments - maps to internal/domain/ticket_message/attachment.go';
```

### Domain: support_agent/

```sql
-- =========================================
-- MAIN TABLE: support_agents
-- Domain: internal/domain/support_agent/
-- Entity: support_agent/entity.go
-- =========================================

CREATE TABLE support_agents (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Agent Reference
    admin_user_id UUID NOT NULL UNIQUE REFERENCES admin_users(id),
    
    -- Profile
    display_name VARCHAR(255) NOT NULL,
    agent_email CITEXT NOT NULL,
    
    -- Skills
    skills TEXT[], -- Array of skill tags
    skill_levels JSONB DEFAULT '{}'::JSONB, -- {"billing": 5, "technical": 4}
    languages TEXT[], -- Array of language codes
    
    -- Capacity
    max_concurrent_tickets INTEGER DEFAULT 10,
    current_ticket_count INTEGER DEFAULT 0,
    
    -- Queues
    assigned_queues TEXT[], -- Array of queue names
    primary_queue VARCHAR(50),
    
    -- Availability
    availability_status VARCHAR(20) NOT NULL CHECK (
        availability_status IN ('ONLINE', 'BUSY', 'AWAY', 'OFFLINE')
    ) DEFAULT 'OFFLINE',
    status_changed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    status_message VARCHAR(200),
    
    -- Auto-Away Settings
    auto_away_enabled BOOLEAN DEFAULT TRUE,
    auto_away_timeout_minutes INTEGER DEFAULT 15,
    last_activity_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    
    -- Performance Stats (rolling 30 days)
    stats_first_response_time_avg INTEGER, -- in seconds
    stats_resolution_time_avg INTEGER, -- in seconds
    stats_tickets_resolved_count INTEGER DEFAULT 0,
    stats_tickets_escalated_count INTEGER DEFAULT 0,
    stats_csat_avg DECIMAL(3,2),
    stats_csat_count INTEGER DEFAULT 0,
    stats_updated_at TIMESTAMPTZ,
    
    -- Schedule
    working_hours JSONB, -- {"monday": {"start": "09:00", "end": "17:00"}}
    timezone VARCHAR(50) DEFAULT 'UTC',
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_support_agents_admin_user ON support_agents (admin_user_id);
CREATE INDEX idx_support_agents_availability ON support_agents (availability_status, is_active) WHERE is_active = TRUE;
CREATE INDEX idx_support_agents_queues ON support_agents USING gin(assigned_queues);
CREATE INDEX idx_support_agents_skills ON support_agents USING gin(skills);
CREATE INDEX idx_support_agents_workload ON support_agents (current_ticket_count, max_concurrent_tickets);

COMMENT ON TABLE support_agents IS 'Support agent profiles - maps to internal/domain/support_agent/entity.go';
COMMENT ON COLUMN support_agents.stats_first_response_time_avg IS 'Average first response time in seconds (rolling 30 days)';
COMMENT ON COLUMN support_agents.stats_resolution_time_avg IS 'Average resolution time in seconds (rolling 30 days)';
```

---

## **SECTION 3: CONTENT & KNOWLEDGE**

### Domain: canned_response/

```sql
-- =========================================
-- MAIN TABLE: canned_responses
-- Domain: internal/domain/canned_response/
-- Entity: canned_response/entity.go
-- =========================================

CREATE TABLE canned_responses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Response Details
    title VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    
    -- Categorization
    category VARCHAR(100),
    tags TEXT[],
    
    -- Localization
    locale VARCHAR(10) NOT NULL DEFAULT 'en-US',
    
    -- Usage Tracking
    usage_count INTEGER DEFAULT 0,
    last_used_at TIMESTAMPTZ,
    
    -- Visibility
    visibility VARCHAR(20) DEFAULT 'ALL_AGENTS' CHECK (
        visibility IN ('ALL_AGENTS', 'QUEUE_SPECIFIC', 'ROLE_SPECIFIC', 'PRIVATE')
    ),
    allowed_roles BIGINT,
    allowed_queues TEXT[],
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    archived_at TIMESTAMPTZ,
    archived_by UUID REFERENCES admin_users(id),
    archive_reason TEXT,
    
    -- Versioning
    version INTEGER DEFAULT 1,
    previous_version_id UUID,
    
    -- Shortcuts
    shortcut_key VARCHAR(50),
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by UUID NOT NULL REFERENCES admin_users(id),
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_by UUID REFERENCES admin_users(id),
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_canned_responses_active ON canned_responses (is_active, category) WHERE is_active = TRUE;
CREATE INDEX idx_canned_responses_locale ON canned_responses (locale);
CREATE INDEX idx_canned_responses_category ON canned_responses (category);
CREATE INDEX idx_canned_responses_tags ON canned_responses USING gin(tags);
CREATE INDEX idx_canned_responses_shortcut ON canned_responses (shortcut_key) WHERE shortcut_key IS NOT NULL;
CREATE INDEX idx_canned_responses_search ON canned_responses USING gin(to_tsvector('english', title || ' ' || body));
CREATE INDEX idx_canned_responses_usage ON canned_responses (usage_count DESC);

COMMENT ON TABLE canned_responses IS 'Canned responses - maps to internal/domain/canned_response/entity.go';
```

### Domain: knowledge_base/

```sql
-- =========================================
-- MAIN TABLE: knowledge_base_articles
-- Domain: internal/domain/knowledge_base/
-- Entity: knowledge_base/entity.go
-- =========================================

CREATE TABLE knowledge_base_articles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Article Identification
    article_slug VARCHAR(200) NOT NULL,
    
    -- Content
    title VARCHAR(500) NOT NULL,
    body TEXT NOT NULL,
    summary TEXT,
    
    -- SEO
    meta_title VARCHAR(200),
    meta_description VARCHAR(500),
    meta_keywords TEXT[],
    
    -- Categorization
    category_id UUID,
    tags TEXT[],
    
    -- Localization
    locale VARCHAR(10) NOT NULL DEFAULT 'en-US',
    
    -- Publishing
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('DRAFT', 'IN_REVIEW', 'PUBLISHED', 'ARCHIVED')
    ) DEFAULT 'DRAFT',
    published_at TIMESTAMPTZ,
    published_by UUID REFERENCES admin_users(id),
    
    -- Versioning
    version INTEGER DEFAULT 1,
    current_version_id UUID,
    
    -- Visibility
    visibility VARCHAR(20) DEFAULT 'PUBLIC' CHECK (
        visibility IN ('PUBLIC', 'AUTHENTICATED', 'PREMIUM', 'INTERNAL')
    ),
    
    -- Analytics
    view_count INTEGER DEFAULT 0,
    helpful_count INTEGER DEFAULT 0,
    not_helpful_count INTEGER DEFAULT 0,
    
    -- Related Content
    related_article_ids UUID[],
    
    -- Author
    author_id UUID NOT NULL REFERENCES admin_users(id),
    last_edited_by UUID REFERENCES admin_users(id),
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    last_published_at TIMESTAMPTZ,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    
    CONSTRAINT uk_kb_articles_slug_locale UNIQUE (article_slug, locale)
);

CREATE INDEX idx_kb_articles_slug ON knowledge_base_articles (article_slug, locale);
CREATE INDEX idx_kb_articles_status ON knowledge_base_articles (status, published_at DESC);
CREATE INDEX idx_kb_articles_category ON knowledge_base_articles (category_id, status);
CREATE INDEX idx_kb_articles_locale ON knowledge_base_articles (locale, status);
CREATE INDEX idx_kb_articles_tags ON knowledge_base_articles USING gin(tags);
CREATE INDEX idx_kb_articles_published ON knowledge_base_articles (published_at DESC) WHERE status = 'PUBLISHED';
CREATE INDEX idx_kb_articles_search ON knowledge_base_articles USING gin(to_tsvector('english', title || ' ' || body));
CREATE INDEX idx_kb_articles_popularity ON knowledge_base_articles (view_count DESC, helpful_count DESC);

COMMENT ON TABLE knowledge_base_articles IS 'Knowledge base articles - maps to internal/domain/knowledge_base/entity.go';

-- =========================================
-- SUB-ENTITY: kb_article_versions
-- Domain: internal/domain/knowledge_base/
-- Entity: knowledge_base/version.go
-- =========================================

CREATE TABLE kb_article_versions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Article Reference
    article_id UUID NOT NULL REFERENCES knowledge_base_articles(id) ON DELETE CASCADE,
    
    -- Version Details
    version_number INTEGER NOT NULL,
    
    -- Snapshot
    title VARCHAR(500) NOT NULL,
    body TEXT NOT NULL,
    summary TEXT,
    
    -- Changes
    change_summary TEXT,
    diff_from_previous JSONB,
    
    -- Author
    created_by UUID NOT NULL REFERENCES admin_users(id),
    
    -- Timestamp
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    
    CONSTRAINT uk_kb_article_versions UNIQUE (article_id, version_number)
);

CREATE INDEX idx_kb_article_versions_article ON kb_article_versions (article_id, version_number DESC);
CREATE INDEX idx_kb_article_versions_created_at ON kb_article_versions (created_at DESC);

COMMENT ON TABLE kb_article_versions IS 'Article version history - maps to internal/domain/knowledge_base/version.go';

-- =========================================
-- SUB-ENTITY: kb_categories
-- Domain: internal/domain/knowledge_base/
-- Entity: knowledge_base/category.go
-- =========================================

CREATE TABLE kb_categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Category Details
    name VARCHAR(200) NOT NULL,
    slug VARCHAR(200) NOT NULL,
    description TEXT,
    
    -- Hierarchy
    parent_category_id UUID REFERENCES kb_categories(id),
    hierarchy_level INTEGER DEFAULT 0,
    hierarchy_path VARCHAR(500), -- e.g., "/billing/refunds"
    
    -- Display
    icon VARCHAR(50),
    color_hex VARCHAR(7),
    display_order INTEGER DEFAULT 0,
    
    -- Localization
    locale VARCHAR(10) NOT NULL DEFAULT 'en-US',
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by UUID REFERENCES admin_users(id),
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    
    CONSTRAINT uk_kb_categories_slug_locale UNIQUE (slug, locale)
);

CREATE INDEX idx_kb_categories_parent ON kb_categories (parent_category_id);
CREATE INDEX idx_kb_categories_active ON kb_categories (is_active, display_order) WHERE is_active = TRUE;
CREATE INDEX idx_kb_categories_locale ON kb_categories (locale);
CREATE INDEX idx_kb_categories_path ON kb_categories (hierarchy_path);

COMMENT ON TABLE kb_categories IS 'KB category tree - maps to internal/domain/knowledge_base/category.go';
```

### Domain: faq/

```sql
-- =========================================
-- MAIN TABLE: faqs
-- Domain: internal/domain/faq/
-- Entity: faq/entity.go
-- =========================================

CREATE TABLE faqs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- FAQ Content
    question VARCHAR(500) NOT NULL,
    answer TEXT NOT NULL,
    
    -- Categorization
    category VARCHAR(100),
    tags TEXT[],
    
    -- Localization
    locale VARCHAR(10) NOT NULL DEFAULT 'en-US',
    
    -- Ordering
    display_order INTEGER DEFAULT 0,
    
    -- Publishing
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('DRAFT', 'PUBLISHED', 'ARCHIVED')
    ) DEFAULT 'DRAFT',
    published_at TIMESTAMPTZ,
    
    -- Analytics
    view_count INTEGER DEFAULT 0,
    helpful_count INTEGER DEFAULT 0,
    not_helpful_count INTEGER DEFAULT 0,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by UUID NOT NULL REFERENCES admin_users(id),
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_by UUID REFERENCES admin_users(id),
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_faqs_status ON faqs (status, display_order);
CREATE INDEX idx_faqs_category ON faqs (category, display_order);
CREATE INDEX idx_faqs_locale ON faqs (locale, status);
CREATE INDEX idx_faqs_published ON faqs (published_at DESC) WHERE status = 'PUBLISHED';
CREATE INDEX idx_faqs_search ON faqs USING gin(to_tsvector('english', question || ' ' || answer));

COMMENT ON TABLE faqs IS 'FAQs - maps to internal/domain/faq/entity.go';
```

---

## **SECTION 4: SAFETY & MODERATION**

### Domain: moderation_queue/

```sql
-- =========================================
-- MAIN TABLE: moderation_queue_items
-- Domain: internal/domain/moderation_queue/
-- Entity: moderation_queue/entity.go
-- =========================================

CREATE TABLE moderation_queue_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Content Reference
    content_type VARCHAR(30) NOT NULL CHECK (
        content_type IN ('JOB', 'PROPOSAL', 'USER_PROFILE', 'REVIEW', 'MESSAGE', 'CONTRACT', 'ASSET', 'COMMENT')
    ),
    content_id UUID NOT NULL,
    content_owner_id UUID NOT NULL,
    content_snapshot JSONB, -- Snapshot of content at time of flagging
    
    -- Flag Details
    flag_reason VARCHAR(50) NOT NULL CHECK (
        flag_reason IN ('SPAM', 'HARASSMENT', 'INAPPROPRIATE_CONTENT', 'COPYRIGHT', 'FRAUD', 'FAKE_REVIEW', 'HATE_SPEECH', 'VIOLENCE', 'NUDITY', 'OTHER')
    ),
    flag_description TEXT,
    flag_weight INTEGER DEFAULT 50, -- 0-100 severity
    
    -- Reporter
    reporter_id UUID,
    reporter_type VARCHAR(20) CHECK (
        reporter_type IN ('USER', 'SYSTEM', 'AUTO_FLAG')
    ),
    
    -- Queue State
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('PENDING', 'ASSIGNED', 'IN_REVIEW', 'ACTIONED', 'DISMISSED', 'ESCALATED')
    ) DEFAULT 'PENDING',
    
    -- Assignment
    assigned_moderator_id UUID REFERENCES admin_users(id),
    assigned_at TIMESTAMPTZ,
    
    -- Review
    reviewed_at TIMESTAMPTZ,
    reviewed_by UUID REFERENCES admin_users(id),
    review_notes TEXT,
    
    -- Decision
    decision VARCHAR(30) CHECK (
        decision IN ('APPROVED', 'REMOVED', 'HIDDEN', 'WARNING_ISSUED', 'USER_SUSPENDED', 'USER_BANNED', 'NO_ACTION')
    ),
    decision_reason TEXT,
    actioned_at TIMESTAMPTZ,
    
    -- Priority
    priority VARCHAR(20) DEFAULT 'MEDIUM' CHECK (
        priority IN ('LOW', 'MEDIUM', 'HIGH', 'URGENT', 'CRITICAL')
    ),
    
    -- Auto-Flagging
    auto_flag_confidence DECIMAL(5,4), -- 0.0 to 1.0
    auto_flag_model_version VARCHAR(50),
    
    -- SLA
    sla_target_review_at TIMESTAMPTZ,
    sla_breached BOOLEAN DEFAULT FALSE,
    
    -- Timestamps
    queued_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    
    CONSTRAINT uk_moderation_queue_content UNIQUE (content_type, content_id)
);

CREATE INDEX idx_moderation_queue_status ON moderation_queue_items (status, priority, queued_at);
CREATE INDEX idx_moderation_queue_assigned ON moderation_queue_items (assigned_moderator_id, status) WHERE assigned_moderator_id IS NOT NULL;
CREATE INDEX idx_moderation_queue_content ON moderation_queue_items (content_type, content_id);
CREATE INDEX idx_moderation_queue_owner ON moderation_queue_items (content_owner_id);
CREATE INDEX idx_moderation_queue_priority ON moderation_queue_items (priority DESC, queued_at);
CREATE INDEX idx_moderation_queue_sla ON moderation_queue_items (sla_breached) WHERE sla_breached = TRUE;
CREATE INDEX idx_moderation_queue_reporter ON moderation_queue_items (reporter_id) WHERE reporter_id IS NOT NULL;

COMMENT ON TABLE moderation_queue_items IS 'Moderation queue - maps to internal/domain/moderation_queue/entity.go';
```

### Domain: user_action/

```sql
-- =========================================
-- MAIN TABLE: admin_user_actions
-- Domain: internal/domain/user_action/
-- Entity: user_action/entity.go
-- =========================================

CREATE TABLE admin_user_actions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Target User
    target_user_id UUID NOT NULL,
    target_user_email CITEXT NOT NULL,
    target_user_type VARCHAR(20) NOT NULL,
    
    -- Action Details
    action_type VARCHAR(30) NOT NULL CHECK (
        action_type IN ('SUSPEND', 'UNSUSPEND', 'BAN', 'UNBAN', 'WARN', 'VERIFY', 'REMOVE_VERIFICATION', 'RESTRICT', 'UNRESTRICT')
    ),
    action_reason_code VARCHAR(50) NOT NULL,
    action_reason_description TEXT NOT NULL,
    
    -- Duration (for temporary actions)
    duration_days INTEGER,
    effective_until TIMESTAMPTZ,
    
    -- Evidence
    evidence_refs JSONB, -- References to evidence files/tickets
    evidence_summary TEXT,
    
    -- Moderation Context
    related_moderation_item_id UUID REFERENCES moderation_queue_items(id),
    related_ticket_id UUID REFERENCES support_tickets(id),
    
    -- Taken By
    actioned_by UUID NOT NULL REFERENCES admin_users(id),
    actioned_by_role VARCHAR(50) NOT NULL,
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('ACTIVE', 'EXPIRED', 'REVERSED', 'SUPERSEDED')
    ) DEFAULT 'ACTIVE',
    
    -- Reversal
    reversed_at TIMESTAMPTZ,
    reversed_by UUID REFERENCES admin_users(id),
    reversal_reason TEXT,
    
    -- Notifications
    user_notified BOOLEAN DEFAULT FALSE,
    notified_at TIMESTAMPTZ,
    notification_channel VARCHAR(20),
    
    -- Appeal
    appeal_submitted BOOLEAN DEFAULT FALSE,
    appeal_decision VARCHAR(20) CHECK (
        appeal_decision IN ('PENDING', 'UPHELD', 'OVERTURNED', 'MODIFIED')
    ),
    
    -- Impact Tracking
    related_contracts_affected INTEGER DEFAULT 0,
    related_jobs_affected INTEGER DEFAULT 0,
    related_proposals_affected INTEGER DEFAULT 0,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at TIMESTAMPTZ,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_user_actions_target ON admin_user_actions (target_user_id, created_at DESC);
CREATE INDEX idx_admin_user_actions_type ON admin_user_actions (action_type, status);
CREATE INDEX idx_admin_user_actions_status ON admin_user_actions (status, effective_until);
CREATE INDEX idx_admin_user_actions_actioned_by ON admin_user_actions (actioned_by, created_at DESC);
CREATE INDEX idx_admin_user_actions_expires ON admin_user_actions (expires_at) WHERE expires_at IS NOT NULL AND status = 'ACTIVE';
CREATE INDEX idx_admin_user_actions_moderation ON admin_user_actions (related_moderation_item_id) WHERE related_moderation_item_id IS NOT NULL;

COMMENT ON TABLE admin_user_actions IS 'Admin actions on users - maps to internal/domain/user_action/entity.go';
```

### Domain: content_action/

```sql
-- =========================================
-- MAIN TABLE: admin_content_actions
-- Domain: internal/domain/content_action/
-- Entity: content_action/entity.go
-- =========================================

CREATE TABLE admin_content_actions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Content Reference
    content_type VARCHAR(30) NOT NULL CHECK (
        content_type IN ('JOB', 'PROPOSAL', 'REVIEW', 'MESSAGE', 'CONTRACT', 'ASSET', 'COMMENT', 'USER_PROFILE')
    ),
    content_id UUID NOT NULL,
    content_owner_id UUID NOT NULL,
    
    -- Action Details
    action_type VARCHAR(30) NOT NULL CHECK (
        action_type IN ('REMOVE', 'HIDE', 'APPROVE', 'REJECT', 'RESTORE', 'FLAG', 'FEATURE', 'UNFEATURE')
    ),
    action_reason_code VARCHAR(50) NOT NULL,
    action_reason_description TEXT NOT NULL,
    
    -- Scope
    action_scope VARCHAR(20) DEFAULT 'FULL' CHECK (
        action_scope IN ('FULL', 'PARTIAL', 'SHADOW')
    ),
    affected_fields TEXT[], -- For partial actions
    
    -- Moderation Context
    related_moderation_item_id UUID REFERENCES moderation_queue_items(id),
    related_ticket_id UUID REFERENCES support_tickets(id),
    
    -- Taken By
    actioned_by UUID NOT NULL REFERENCES admin_users(id),
    actioned_by_role VARCHAR(50) NOT NULL,
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('ACTIVE', 'REVERSED', 'SUPERSEDED')
    ) DEFAULT 'ACTIVE',
    
    -- Reversal
    reversed_at TIMESTAMPTZ,
    reversed_by UUID REFERENCES admin_users(id),
    reversal_reason TEXT,
    
    -- Notifications
    owner_notified BOOLEAN DEFAULT FALSE,
    notified_at TIMESTAMPTZ,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    
    CONSTRAINT uk_admin_content_actions_content UNIQUE (content_type, content_id, action_type, status)
);

CREATE INDEX idx_admin_content_actions_content ON admin_content_actions (content_type, content_id);
CREATE INDEX idx_admin_content_actions_owner ON admin_content_actions (content_owner_id, created_at DESC);
CREATE INDEX idx_admin_content_actions_type ON admin_content_actions (action_type, status);
CREATE INDEX idx_admin_content_actions_actioned_by ON admin_content_actions (actioned_by, created_at DESC);
CREATE INDEX idx_admin_content_actions_moderation ON admin_content_actions (related_moderation_item_id) WHERE related_moderation_item_id IS NOT NULL;

COMMENT ON TABLE admin_content_actions IS 'Admin actions on content - maps to internal/domain/content_action/entity.go';
```

---

## **SECTION 5: DISPUTES & CASES**

### Domain: dispute_resolution/

```sql
-- =========================================
-- MAIN TABLE: admin_dispute_cases
-- Domain: internal/domain/dispute_resolution/
-- Entity: dispute_resolution/entity.go
-- =========================================

CREATE TABLE admin_dispute_cases (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Case Identification
    case_number VARCHAR(20) NOT NULL UNIQUE, -- e.g., DSP-2025-001234
    
    -- Parties
    claimant_user_id UUID NOT NULL,
    respondent_user_id UUID NOT NULL,
    
    -- Related Entity
    related_entity_type VARCHAR(30) NOT NULL CHECK (
        related_entity_type IN ('CONTRACT', 'PAYMENT', 'JOB', 'PROPOSAL', 'REVIEW')
    ),
    related_entity_id UUID NOT NULL,
    
    -- Dispute Details
    dispute_type VARCHAR(50) NOT NULL CHECK (
        dispute_type IN ('PAYMENT_DISPUTE', 'QUALITY_DISPUTE', 'SCOPE_DISPUTE', 'CANCELLATION_DISPUTE', 'REFUND_DISPUTE', 'REVIEW_DISPUTE', 'OTHER')
    ),
    dispute_amount BIGINT,
    dispute_currency CHAR(3) DEFAULT 'USD',
    
    -- Claims
    claimant_claim TEXT NOT NULL,
    respondent_response TEXT,
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('OPENED', 'UNDER_REVIEW', 'EVIDENCE_GATHERING', 'AWAITING_DECISION', 'DECIDED', 'CLOSED', 'APPEALED')
    ) DEFAULT 'OPENED',
    
    -- Assignment
    assigned_arbitrator_id UUID REFERENCES admin_users(id),
    assigned_at TIMESTAMPTZ,
    
    -- Timeline
    state_timeline JSONB DEFAULT '[]'::JSONB, -- History of state changes
    
    -- Decision
    decision VARCHAR(30) CHECK (
        decision IN ('FAVOR_CLAIMANT', 'FAVOR_RESPONDENT', 'SPLIT_DECISION', 'DISMISSED', 'SETTLED')
    ),
    decision_rationale TEXT,
    decision_made_at TIMESTAMPTZ,
    decision_made_by UUID REFERENCES admin_users(id),
    
    -- Remedies
    remedy_type VARCHAR(30),
    remedy_amount BIGINT,
    remedy_currency CHAR(3) DEFAULT 'USD',
    remedy_description TEXT,
    
    -- Appeal
    appeal_submitted BOOLEAN DEFAULT FALSE,
    appeal_decision VARCHAR(20),
    appeal_decided_at TIMESTAMPTZ,
    
    -- SLA
    sla_target_decision_at TIMESTAMPTZ,
    sla_breached BOOLEAN DEFAULT FALSE,
    
    -- Timestamps
    opened_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    closed_at TIMESTAMPTZ,
    last_updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    internal_notes TEXT
);

CREATE INDEX idx_admin_dispute_cases_number ON admin_dispute_cases (case_number);
CREATE INDEX idx_admin_dispute_cases_claimant ON admin_dispute_cases (claimant_user_id, opened_at DESC);
CREATE INDEX idx_admin_dispute_cases_respondent ON admin_dispute_cases (respondent_user_id, opened_at DESC);
CREATE INDEX idx_admin_dispute_cases_status ON admin_dispute_cases (status, opened_at DESC);
CREATE INDEX idx_admin_dispute_cases_assigned ON admin_dispute_cases (assigned_arbitrator_id, status) WHERE assigned_arbitrator_id IS NOT NULL;
CREATE INDEX idx_admin_dispute_cases_related ON admin_dispute_cases (related_entity_type, related_entity_id);
CREATE INDEX idx_admin_dispute_cases_sla ON admin_dispute_cases (sla_breached) WHERE sla_breached = TRUE;

COMMENT ON TABLE admin_dispute_cases IS 'Dispute resolution cases - maps to internal/domain/dispute_resolution/entity.go';

-- =========================================
-- SUB-ENTITY: dispute_evidence
-- Domain: internal/domain/dispute_resolution/
-- Entity: dispute_resolution/evidence.go
-- =========================================

CREATE TABLE admin_dispute_evidence (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Case Reference
    dispute_case_id UUID NOT NULL REFERENCES admin_dispute_cases(id) ON DELETE CASCADE,
    
    -- Submitted By
    submitted_by_user_id UUID NOT NULL,
    submitted_by_party VARCHAR(20) NOT NULL CHECK (
        submitted_by_party IN ('CLAIMANT', 'RESPONDENT', 'ARBITRATOR')
    ),
    
    -- Evidence Details
    evidence_type VARCHAR(30) NOT NULL CHECK (
        evidence_type IN ('DOCUMENT', 'SCREENSHOT', 'MESSAGE_THREAD', 'CONTRACT_DOCUMENT', 'PAYMENT_PROOF', 'COMMUNICATION_LOG', 'OTHER')
    ),
    description TEXT NOT NULL,
    
    -- File Reference
    file_storage_path VARCHAR(500),
    file_name VARCHAR(255),
    file_size BIGINT,
    file_checksum VARCHAR(64),
    
    -- Integrity
    integrity_verified BOOLEAN DEFAULT FALSE,
    verified_at TIMESTAMPTZ,
    verified_by UUID REFERENCES admin_users(id),
    
    -- Timestamps
    submitted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_dispute_evidence_case ON admin_dispute_evidence (dispute_case_id, submitted_at DESC);
CREATE INDEX idx_dispute_evidence_submitter ON admin_dispute_evidence (submitted_by_user_id);
CREATE INDEX idx_dispute_evidence_party ON admin_dispute_evidence (dispute_case_id, submitted_by_party);

COMMENT ON TABLE admin_dispute_evidence IS 'Dispute evidence - maps to internal/domain/dispute_resolution/evidence.go';
```

### Domain: appeal/

```sql
-- =========================================
-- MAIN TABLE: admin_appeals
-- Domain: internal/domain/appeal/
-- Entity: appeal/entity.go
-- =========================================

CREATE TABLE admin_appeals (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Appeal Identification
    appeal_number VARCHAR(20) NOT NULL UNIQUE, -- e.g., APL-2025-001234
    
    -- Original Action Reference
    original_action_type VARCHAR(30) NOT NULL CHECK (
        original_action_type IN ('USER_SUSPENSION', 'USER_BAN', 'CONTENT_REMOVAL', 'DISPUTE_DECISION', 'VERIFICATION_DENIAL')
    ),
    original_action_id UUID NOT NULL,
    
    -- Appellant
    appellant_user_id UUID NOT NULL,
    appellant_email CITEXT NOT NULL,
    
    -- Appeal Details
    appeal_reason_code VARCHAR(50) NOT NULL,
    appeal_statement TEXT NOT NULL,
    supporting_evidence JSONB,
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('SUBMITTED', 'UNDER_REVIEW', 'AWAITING_INFO', 'DECIDED', 'CLOSED')
    ) DEFAULT 'SUBMITTED',
    
    -- Review
    assigned_reviewer_id UUID REFERENCES admin_users(id),
    assigned_at TIMESTAMPTZ,
    reviewed_at TIMESTAMPTZ,
    reviewer_notes TEXT,
    
    -- Decision
    decision VARCHAR(20) CHECK (
        decision IN ('UPHELD', 'OVERTURNED', 'MODIFIED', 'DISMISSED')
    ),
    decision_rationale TEXT,
    decision_made_at TIMESTAMPTZ,
    decision_made_by UUID REFERENCES admin_users(id),
    
    -- Modification Details (if applicable)
    modification_details JSONB,
    
    -- SLA
    sla_target_decision_at TIMESTAMPTZ,
    sla_breached BOOLEAN DEFAULT FALSE,
    
    -- Timestamps
    submitted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    closed_at TIMESTAMPTZ,
    last_updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_appeals_number ON admin_appeals (appeal_number);
CREATE INDEX idx_admin_appeals_appellant ON admin_appeals (appellant_user_id, submitted_at DESC);
CREATE INDEX idx_admin_appeals_status ON admin_appeals (status, submitted_at DESC);
CREATE INDEX idx_admin_appeals_assigned ON admin_appeals (assigned_reviewer_id, status) WHERE assigned_reviewer_id IS NOT NULL;
CREATE INDEX idx_admin_appeals_original_action ON admin_appeals (original_action_type, original_action_id);
CREATE INDEX idx_admin_appeals_sla ON admin_appeals (sla_breached) WHERE sla_breached = TRUE;

COMMENT ON TABLE admin_appeals IS 'Appeal cases - maps to internal/domain/appeal/entity.go';
```

### Domain: case_link/

```sql
-- =========================================
-- MAIN TABLE: admin_case_links
-- Domain: internal/domain/case_link/
-- Entity: case_link/entity.go
-- =========================================

CREATE TABLE admin_case_links (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Source Case
    source_case_type VARCHAR(30) NOT NULL CHECK (
        source_case_type IN ('TICKET', 'DISPUTE', 'APPEAL', 'FRAUD_REVIEW', 'KYC_CASE', 'REFUND_CASE')
    ),
    source_case_id UUID NOT NULL,
    
    -- Target Case
    target_case_type VARCHAR(30) NOT NULL CHECK (
        target_case_type IN ('TICKET', 'DISPUTE', 'APPEAL', 'FRAUD_REVIEW', 'KYC_CASE', 'REFUND_CASE')
    ),
    target_case_id UUID NOT NULL,
    
    -- Link Details
    link_type VARCHAR(30) NOT NULL CHECK (
        link_type IN ('RELATED', 'DUPLICATE', 'PARENT_CHILD', 'BLOCKS', 'BLOCKED_BY', 'CAUSES', 'CAUSED_BY')
    ),
    link_reason TEXT,
    
    -- Created By
    created_by UUID NOT NULL REFERENCES admin_users(id),
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    
    CONSTRAINT uk_admin_case_links UNIQUE (source_case_type, source_case_id, target_case_type, target_case_id, link_type)
);

CREATE INDEX idx_admin_case_links_source ON admin_case_links (source_case_type, source_case_id);
CREATE INDEX idx_admin_case_links_target ON admin_case_links (target_case_type, target_case_id);
CREATE INDEX idx_admin_case_links_type ON admin_case_links (link_type);

COMMENT ON TABLE admin_case_links IS 'Links between cases - maps to internal/domain/case_link/entity.go';
```

---

## **SECTION 6: COMPLIANCE & LEGAL**

### Domain: legal_hold/

```sql
-- =========================================
-- MAIN TABLE: admin_legal_holds
-- Domain: internal/domain/legal_hold/
-- Entity: legal_hold/entity.go
-- =========================================

CREATE TABLE admin_legal_holds (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Hold Identification
    hold_number VARCHAR(20) NOT NULL UNIQUE, -- e.g., LH-2025-001234
    
    -- Hold Details
    hold_name VARCHAR(255) NOT NULL,
    hold_reason TEXT NOT NULL,
    legal_matter_name VARCHAR(255),
    legal_matter_number VARCHAR(100),
    
    -- Scope
    scope_type VARCHAR(30) NOT NULL CHECK (
        scope_type IN ('USER', 'CONTRACT', 'CONTENT', 'COMMUNICATION', 'FINANCIAL', 'MULTI_ENTITY')
    ),
    scope_entity_ids UUID[], -- Array of affected entity IDs
    scope_description TEXT,
    
    -- Placed By
    placed_by UUID NOT NULL REFERENCES admin_users(id),
    placed_by_authority VARCHAR(100), -- e.g., "Legal Counsel", "Court Order"
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('ACTIVE', 'RELEASED', 'EXPIRED')
    ) DEFAULT 'ACTIVE',
    
    -- Release
    released_at TIMESTAMPTZ,
    released_by UUID REFERENCES admin_users(id),
    release_reason TEXT,
    release_authorization VARCHAR(255),
    
    -- Data Preservation
    data_preserved BOOLEAN DEFAULT FALSE,
    preservation_location VARCHAR(500),
    preservation_checksum VARCHAR(64),
    
    -- Timestamps
    placed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at TIMESTAMPTZ,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    internal_notes TEXT
);

CREATE INDEX idx_admin_legal_holds_number ON admin_legal_holds (hold_number);
CREATE INDEX idx_admin_legal_holds_status ON admin_legal_holds (status, placed_at DESC);
CREATE INDEX idx_admin_legal_holds_placed_by ON admin_legal_holds (placed_by, placed_at DESC);
CREATE INDEX idx_admin_legal_holds_scope ON admin_legal_holds USING gin(scope_entity_ids);
CREATE INDEX idx_admin_legal_holds_expires ON admin_legal_holds (expires_at) WHERE expires_at IS NOT NULL;

COMMENT ON TABLE admin_legal_holds IS 'Legal holds - maps to internal/domain/legal_hold/entity.go';

-- =========================================
-- SUB-ENTITY: legal_hold_export_jobs
-- Domain: internal/domain/legal_hold/
-- Entity: legal_hold/export_job.go
-- =========================================

CREATE TABLE admin_legal_hold_export_jobs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Hold Reference
    legal_hold_id UUID NOT NULL REFERENCES admin_legal_holds(id),
    
    -- Export Details
    export_name VARCHAR(255) NOT NULL,
    export_format VARCHAR(20) CHECK (
        export_format IN ('ZIP', 'PDF', 'CSV', 'JSON', 'PARQUET')
    ),
    
    -- Scope
    export_scope JSONB NOT NULL, -- Detailed scope of export
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('PENDING', 'IN_PROGRESS', 'COMPLETED', 'FAILED')
    ) DEFAULT 'PENDING',
    
    -- Progress
    total_records INTEGER,
    processed_records INTEGER DEFAULT 0,
    progress_percentage DECIMAL(5,2),
    
    -- Output
    output_file_path VARCHAR(500),
    output_file_size BIGINT,
    output_checksum VARCHAR(64),
    
    -- Requested By
    requested_by UUID NOT NULL REFERENCES admin_users(id),
    
    -- Timestamps
    requested_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    started_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    
    -- Error Details
    error_message TEXT,
    error_code VARCHAR(50),
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_legal_hold_export_jobs_hold ON admin_legal_hold_export_jobs (legal_hold_id, requested_at DESC);
CREATE INDEX idx_legal_hold_export_jobs_status ON admin_legal_hold_export_jobs (status, requested_at DESC);
CREATE INDEX idx_legal_hold_export_jobs_requested_by ON admin_legal_hold_export_jobs (requested_by);

COMMENT ON TABLE admin_legal_hold_export_jobs IS 'eDiscovery export jobs - maps to internal/domain/legal_hold/export_job.go';
```

### Domain: privacy_request/

```sql
-- =========================================
-- MAIN TABLE: admin_privacy_requests
-- Domain: internal/domain/privacy_request/
-- Entity: privacy_request/entity.go
-- =========================================

CREATE TABLE admin_privacy_requests (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Request Identification
    request_number VARCHAR(20) NOT NULL UNIQUE, -- e.g., DSAR-2025-001234
    
    -- Subject
    subject_user_id UUID,
    subject_email CITEXT NOT NULL,
    subject_name VARCHAR(255),
    
    -- Request Type
    request_type VARCHAR(20) NOT NULL CHECK (
        request_type IN ('ACCESS', 'ERASURE', 'RECTIFICATION', 'PORTABILITY', 'RESTRICTION', 'OBJECTION')
    ),
    
    -- Regulation
    regulation VARCHAR(20) NOT NULL CHECK (
        regulation IN ('GDPR', 'CCPA', 'LGPD', 'OTHER')
    ),
    
    -- Request Details
    request_details TEXT NOT NULL,
    request_scope JSONB, -- Specific data categories requested
    
    -- Identity Verification
    identity_verified BOOLEAN DEFAULT FALSE,
    identity_verification_method VARCHAR(50),
    identity_verified_at TIMESTAMPTZ,
    identity_verified_by UUID REFERENCES admin_users(id),
    identity_evidence_refs JSONB,
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('SUBMITTED', 'IDENTITY_VERIFICATION', 'APPROVED', 'IN_PROGRESS', 'FULFILLED', 'DENIED', 'WITHDRAWN')
    ) DEFAULT 'SUBMITTED',
    
    -- Assignment
    assigned_privacy_officer_id UUID REFERENCES admin_users(id),
    assigned_at TIMESTAMPTZ,
    
    -- Processing
    approved_at TIMESTAMPTZ,
    approved_by UUID REFERENCES admin_users(id),
    approval_notes TEXT,
    
    -- Denial (if applicable)
    denial_reason VARCHAR(100),
    denial_rationale TEXT,
    denied_at TIMESTAMPTZ,
    denied_by UUID REFERENCES admin_users(id),
    
    -- Fulfillment
    fulfilled_at TIMESTAMPTZ,
    fulfillment_method VARCHAR(50),
    fulfillment_location VARCHAR(500),
    fulfillment_checksum VARCHAR(64),
    
    -- SLA (30 days for GDPR)
    sla_target_fulfillment_at TIMESTAMPTZ,
    sla_breached BOOLEAN DEFAULT FALSE,
    
    -- Timestamps
    submitted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    closed_at TIMESTAMPTZ,
    last_updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    internal_notes TEXT
);

CREATE INDEX idx_admin_privacy_requests_number ON admin_privacy_requests (request_number);
CREATE INDEX idx_admin_privacy_requests_subject ON admin_privacy_requests (subject_user_id, submitted_at DESC);
CREATE INDEX idx_admin_privacy_requests_email ON admin_privacy_requests (subject_email);
CREATE INDEX idx_admin_privacy_requests_status ON admin_privacy_requests (status, submitted_at DESC);
CREATE INDEX idx_admin_privacy_requests_type ON admin_privacy_requests (request_type, status);
CREATE INDEX idx_admin_privacy_requests_assigned ON admin_privacy_requests (assigned_privacy_officer_id, status) WHERE assigned_privacy_officer_id IS NOT NULL;
CREATE INDEX idx_admin_privacy_requests_sla ON admin_privacy_requests (sla_breached) WHERE sla_breached = TRUE;

COMMENT ON TABLE admin_privacy_requests IS 'GDPR/CCPA DSAR requests - maps to internal/domain/privacy_request/entity.go';
```

### Domain: pii_access/

```sql
-- =========================================
-- MAIN TABLE: admin_pii_access_requests
-- Domain: internal/domain/pii_access/
-- Entity: pii_access/entity.go
-- =========================================

CREATE TABLE admin_pii_access_requests (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Request Identification
    request_number VARCHAR(20) NOT NULL UNIQUE, -- e.g., PII-2025-001234
    
    -- Requester
    requester_admin_id UUID NOT NULL REFERENCES admin_users(id),
    requester_role VARCHAR(50) NOT NULL,
    
    -- Purpose
    access_purpose VARCHAR(50) NOT NULL CHECK (
        access_purpose IN ('SUPPORT', 'INVESTIGATION', 'COMPLIANCE', 'LEGAL_REQUIREMENT', 'INCIDENT_RESPONSE', 'AUDIT')
    ),
    purpose_description TEXT NOT NULL,
    justification TEXT NOT NULL,
    
    -- Scope
    target_user_id UUID,
    target_data_types TEXT[], -- e.g., ["email", "phone", "address", "ssn"]
    scope_description TEXT,
    
    -- Related Case
    related_case_type VARCHAR(30),
    related_case_id UUID,
    
    -- Approval Workflow
    requires_approval BOOLEAN DEFAULT TRUE,
    approval_status VARCHAR(20) CHECK (
        approval_status IN ('PENDING', 'APPROVED', 'DENIED', 'EXPIRED')
    ) DEFAULT 'PENDING',
    
    approved_by UUID REFERENCES admin_users(id),
    approved_at TIMESTAMPTZ,
    approval_notes TEXT,
    
    denied_by UUID REFERENCES admin_users(id),
    denied_at TIMESTAMPTZ,
    denial_reason TEXT,
    
    -- Grant Details
    granted BOOLEAN DEFAULT FALSE,
    grant_expires_at TIMESTAMPTZ,
    grant_ttl_minutes INTEGER DEFAULT 30,
    
    -- Access Tracking
    access_count INTEGER DEFAULT 0,
    last_accessed_at TIMESTAMPTZ,
    access_log_refs UUID[],
    
    -- Masking Policy
    masking_policy_id UUID,
    fields_unmasked TEXT[],
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('PENDING', 'ACTIVE', 'EXPIRED', 'REVOKED', 'DENIED')
    ) DEFAULT 'PENDING',
    
    -- Revocation
    revoked_at TIMESTAMPTZ,
    revoked_by UUID REFERENCES admin_users(id),
    revocation_reason TEXT,
    
    -- Timestamps
    requested_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at TIMESTAMPTZ,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_pii_access_requester ON admin_pii_access_requests (requester_admin_id, requested_at DESC);
CREATE INDEX idx_admin_pii_access_target ON admin_pii_access_requests (target_user_id) WHERE target_user_id IS NOT NULL;
CREATE INDEX idx_admin_pii_access_status ON admin_pii_access_requests (status, requested_at DESC);
CREATE INDEX idx_admin_pii_access_expires ON admin_pii_access_requests (expires_at) WHERE status = 'ACTIVE';
CREATE INDEX idx_admin_pii_access_related_case ON admin_pii_access_requests (related_case_type, related_case_id);

COMMENT ON TABLE admin_pii_access_requests IS 'Break-glass PII access - maps to internal/domain/pii_access/entity.go';

-- =========================================
-- SUB-ENTITY: admin_pii_masking_policies
-- Domain: internal/domain/pii_access/
-- Entity: pii_access/policy.go
-- =========================================

CREATE TABLE admin_pii_masking_policies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Policy Details
    policy_name VARCHAR(100) NOT NULL UNIQUE,
    policy_description TEXT,
    
    -- Fields Configuration
    maskable_fields JSONB NOT NULL, -- {"email": "PARTIAL", "phone": "FULL", "ssn": "FULL"}
    redaction_rules JSONB NOT NULL,
    
    -- Applicability
    applies_to_roles BIGINT, -- Bitset of roles
    applies_to_purposes TEXT[],
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by UUID REFERENCES admin_users(id),
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_pii_masking_policies_active ON admin_pii_masking_policies (is_active) WHERE is_active = TRUE;

COMMENT ON TABLE admin_pii_masking_policies IS 'PII masking policies - maps to internal/domain/pii_access/policy.go';
```

### Domain: ip_claim/

```sql
-- =========================================
-- MAIN TABLE: admin_ip_claims
-- Domain: internal/domain/ip_claim/
-- Entity: ip_claim/entity.go
-- =========================================

CREATE TABLE admin_ip_claims (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Claim Identification
    claim_number VARCHAR(20) NOT NULL UNIQUE, -- e.g., DMCA-2025-001234
    
    -- Claimant
    claimant_name VARCHAR(255) NOT NULL,
    claimant_email CITEXT NOT NULL,
    claimant_organization VARCHAR(255),
    claimant_contact_details JSONB,
    
    -- Claim Type
    claim_type VARCHAR(20) NOT NULL CHECK (
        claim_type IN ('COPYRIGHT', 'TRADEMARK', 'PATENT', 'OTHER_IP')
    ),
    
    -- Infringing Content
    content_type VARCHAR(30) NOT NULL CHECK (
        content_type IN ('JOB', 'PROPOSAL', 'USER_PROFILE', 'PORTFOLIO_ITEM', 'MESSAGE', 'REVIEW')
    ),
    content_id UUID NOT NULL,
    content_owner_id UUID NOT NULL,
    content_snapshot JSONB,
    
    -- Claim Details
    claim_description TEXT NOT NULL,
    original_work_description TEXT,
    ownership_evidence_refs JSONB,
    
    -- Legal Basis
    legal_basis TEXT NOT NULL,
    sworn_statement BOOLEAN DEFAULT FALSE,
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('SUBMITTED', 'UNDER_REVIEW', 'CONTENT_REMOVED', 'COUNTER_CLAIMED', 'RESOLVED', 'DISMISSED')
    ) DEFAULT 'SUBMITTED',
    
    -- Review
    assigned_reviewer_id UUID REFERENCES admin_users(id),
    assigned_at TIMESTAMPTZ,
    reviewed_at TIMESTAMPTZ,
    reviewer_notes TEXT,
    
    -- Action Taken
    action_type VARCHAR(30) CHECK (
        action_type IN ('CONTENT_REMOVED', 'ACCOUNT_SUSPENDED', 'NO_ACTION', 'WARNING_ISSUED')
    ),
    action_taken_at TIMESTAMPTZ,
    action_taken_by UUID REFERENCES admin_users(id),
    
    -- Counter-Claim
    counter_claim_submitted BOOLEAN DEFAULT FALSE,
    counter_claim_statement TEXT,
    counter_claim_submitted_at TIMESTAMPTZ,
    
    -- Resolution
    resolution VARCHAR(50),
    resolution_notes TEXT,
    resolved_at TIMESTAMPTZ,
    
    -- Timestamps
    submitted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    closed_at TIMESTAMPTZ,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_ip_claims_number ON admin_ip_claims (claim_number);
CREATE INDEX idx_admin_ip_claims_claimant ON admin_ip_claims (claimant_email, submitted_at DESC);
CREATE INDEX idx_admin_ip_claims_content ON admin_ip_claims (content_type, content_id);
CREATE INDEX idx_admin_ip_claims_owner ON admin_ip_claims (content_owner_id);
CREATE INDEX idx_admin_ip_claims_status ON admin_ip_claims (status, submitted_at DESC);
CREATE INDEX idx_admin_ip_claims_assigned ON admin_ip_claims (assigned_reviewer_id, status) WHERE assigned_reviewer_id IS NOT NULL;

COMMENT ON TABLE admin_ip_claims IS 'IP/DMCA claims - maps to internal/domain/ip_claim/entity.go';
```

---

## **SECTION 7: IDENTITY & VERIFICATION**

### Domain: kyc_case/

```sql
-- =========================================
-- MAIN TABLE: admin_kyc_cases
-- Domain: internal/domain/kyc_case/
-- Entity: kyc_case/entity.go
-- =========================================

CREATE TABLE admin_kyc_cases (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Case Identification
    case_number VARCHAR(20) NOT NULL UNIQUE, -- e.g., KYC-2025-001234
    
    -- Subject
    subject_user_id UUID NOT NULL,
    subject_email CITEXT NOT NULL,
    subject_name VARCHAR(255) NOT NULL,
    
    -- KYC Level
    kyc_level VARCHAR(20) NOT NULL CHECK (
        kyc_level IN ('BASIC', 'INTERMEDIATE', 'ADVANCED', 'BUSINESS')
    ),
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('SUBMITTED', 'DOCUMENTS_REQUESTED', 'UNDER_REVIEW', 'APPROVED', 'REJECTED', 'EXPIRED')
    ) DEFAULT 'SUBMITTED',
    
    -- Document Requirements
    required_documents TEXT[], -- e.g., ["ID_CARD", "PROOF_OF_ADDRESS", "SELFIE"]
    submitted_documents UUID[], -- References to document IDs
    documents_complete BOOLEAN DEFAULT FALSE,
    
    -- Review
    assigned_reviewer_id UUID REFERENCES admin_users(id),
    assigned_at TIMESTAMPTZ,
    reviewed_at TIMESTAMPTZ,
    reviewer_notes TEXT,
    
    -- Decision
    decision VARCHAR(20) CHECK (
        decision IN ('APPROVED', 'REJECTED', 'REQUIRES_ADDITIONAL_INFO')
    ),
    decision_rationale TEXT,
    decision_made_at TIMESTAMPTZ,
    decision_made_by UUID REFERENCES admin_users(id),
    
    -- Rejection Details
    rejection_reasons TEXT[],
    rejection_notes TEXT,
    
    -- Verification Details
    identity_verified BOOLEAN DEFAULT FALSE,
    address_verified BOOLEAN DEFAULT FALSE,
    document_authenticity_verified BOOLEAN DEFAULT FALSE,
    
    -- Risk Assessment
    risk_score INTEGER CHECK (risk_score >= 0 AND risk_score <= 100),
    risk_factors JSONB,
    
    -- Third-Party Verification
    verification_provider VARCHAR(50),
    verification_reference VARCHAR(255),
    verification_result JSONB,
    
    -- Expiry
    expires_at TIMESTAMPTZ,
    
    -- SLA
    sla_target_decision_at TIMESTAMPTZ,
    sla_breached BOOLEAN DEFAULT FALSE,
    
    -- Timestamps
    submitted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    closed_at TIMESTAMPTZ,
    last_updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_kyc_cases_number ON admin_kyc_cases (case_number);
CREATE INDEX idx_admin_kyc_cases_subject ON admin_kyc_cases (subject_user_id, submitted_at DESC);
CREATE INDEX idx_admin_kyc_cases_status ON admin_kyc_cases (status, submitted_at DESC);
CREATE INDEX idx_admin_kyc_cases_level ON admin_kyc_cases (kyc_level, status);
CREATE INDEX idx_admin_kyc_cases_assigned ON admin_kyc_cases (assigned_reviewer_id, status) WHERE assigned_reviewer_id IS NOT NULL;
CREATE INDEX idx_admin_kyc_cases_sla ON admin_kyc_cases (sla_breached) WHERE sla_breached = TRUE;
CREATE INDEX idx_admin_kyc_cases_expires ON admin_kyc_cases (expires_at) WHERE expires_at IS NOT NULL;

COMMENT ON TABLE admin_kyc_cases IS 'KYC verification cases - maps to internal/domain/kyc_case/entity.go';

-- =========================================
-- SUB-ENTITY: kyc_documents
-- Domain: internal/domain/kyc_case/
-- Entity: kyc_case/document.go
-- =========================================

CREATE TABLE admin_kyc_documents (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Case Reference
    kyc_case_id UUID NOT NULL REFERENCES admin_kyc_cases(id) ON DELETE CASCADE,
    
    -- Document Details
    document_type VARCHAR(30) NOT NULL CHECK (
        document_type IN ('ID_CARD', 'PASSPORT', 'DRIVERS_LICENSE', 'PROOF_OF_ADDRESS', 'SELFIE', 'BANK_STATEMENT', 'UTILITY_BILL', 'OTHER')
    ),
    document_number VARCHAR(100),
    issuing_country CHAR(2),
    issuing_authority VARCHAR(255),
    issue_date DATE,
    expiry_date DATE,
    
    -- File Storage
    file_storage_path VARCHAR(500) NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_size BIGINT,
    file_type VARCHAR(100),
    file_checksum VARCHAR(64),
    
    -- Verification
    verification_status VARCHAR(20) CHECK (
        verification_status IN ('PENDING', 'VERIFIED', 'REJECTED', 'EXPIRED')
    ) DEFAULT 'PENDING',
    verification_method VARCHAR(50),
    verified_at TIMESTAMPTZ,
    verified_by UUID REFERENCES admin_users(id),
    rejection_reason TEXT,
    
    -- OCR & Data Extraction
    extracted_data JSONB,
    ocr_confidence_score DECIMAL(5,4),
    
    -- Security
    pii_encrypted BOOLEAN DEFAULT TRUE,
    access_restricted BOOLEAN DEFAULT TRUE,
    
    -- Timestamps
    uploaded_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_kyc_documents_case ON admin_kyc_documents (kyc_case_id, uploaded_at DESC);
CREATE INDEX idx_kyc_documents_type ON admin_kyc_documents (document_type, verification_status);
CREATE INDEX idx_kyc_documents_verification ON admin_kyc_documents (verification_status) WHERE verification_status != 'VERIFIED';

COMMENT ON TABLE admin_kyc_documents IS 'KYC documents - maps to internal/domain/kyc_case/document.go';
```

### Domain: business_verification/

```sql
-- =========================================
-- MAIN TABLE: admin_business_verifications
-- Domain: internal/domain/business_verification/
-- Entity: business_verification/entity.go
-- =========================================

CREATE TABLE admin_business_verifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Case Identification
    verification_number VARCHAR(20) NOT NULL UNIQUE, -- e.g., BIZ-2025-001234
    
    -- Business Owner
    business_owner_user_id UUID NOT NULL,
    
    -- Business Details
    business_name VARCHAR(255) NOT NULL,
    business_legal_name VARCHAR(255),
    business_type VARCHAR(50) CHECK (
        business_type IN ('SOLE_PROPRIETORSHIP', 'PARTNERSHIP', 'LLC', 'CORPORATION', 'NONPROFIT', 'OTHER')
    ),
    business_registration_number VARCHAR(100),
    tax_id VARCHAR(50),
    
    -- Business Address
    business_address_line1 VARCHAR(255),
    business_address_line2 VARCHAR(255),
    business_city VARCHAR(100),
    business_state VARCHAR(100),
    business_postal_code VARCHAR(20),
    business_country CHAR(2),
    
    -- Contact Details
    business_phone VARCHAR(30),
    business_email CITEXT,
    business_website VARCHAR(500),
    
    -- Industry
    industry_category VARCHAR(100),
    industry_description TEXT,
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('SUBMITTED', 'DOCUMENTS_REQUESTED', 'UNDER_REVIEW', 'APPROVED', 'REJECTED')
    ) DEFAULT 'SUBMITTED',
    
    -- Document Verification
    required_documents TEXT[], -- e.g., ["BUSINESS_LICENSE", "TAX_CERTIFICATE", "ARTICLES_OF_INCORPORATION"]
    submitted_documents UUID[],
    documents_complete BOOLEAN DEFAULT FALSE,
    
    -- Review
    assigned_reviewer_id UUID REFERENCES admin_users(id),
    assigned_at TIMESTAMPTZ,
    reviewed_at TIMESTAMPTZ,
    reviewer_notes TEXT,
    
    -- Decision
    decision VARCHAR(20) CHECK (
        decision IN ('APPROVED', 'REJECTED', 'REQUIRES_ADDITIONAL_INFO')
    ),
    decision_rationale TEXT,
    decision_made_at TIMESTAMPTZ,
    decision_made_by UUID REFERENCES admin_users(id),
    
    -- Rejection Details
    rejection_reasons TEXT[],
    rejection_notes TEXT,
    
    -- Third-Party Verification
    verification_provider VARCHAR(50),
    verification_reference VARCHAR(255),
    verification_result JSONB,
    
    -- SLA
    sla_target_decision_at TIMESTAMPTZ,
    sla_breached BOOLEAN DEFAULT FALSE,
    
    -- Timestamps
    submitted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    closed_at TIMESTAMPTZ,
    last_updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_business_verifications_number ON admin_business_verifications (verification_number);
CREATE INDEX idx_admin_business_verifications_owner ON admin_business_verifications (business_owner_user_id, submitted_at DESC);
CREATE INDEX idx_admin_business_verifications_status ON admin_business_verifications (status, submitted_at DESC);
CREATE INDEX idx_admin_business_verifications_assigned ON admin_business_verifications (assigned_reviewer_id, status) WHERE assigned_reviewer_id IS NOT NULL;
CREATE INDEX idx_admin_business_verifications_sla ON admin_business_verifications (sla_breached) WHERE sla_breached = TRUE;

COMMENT ON TABLE admin_business_verifications IS 'Business verification - maps to internal/domain/business_verification/entity.go';
```

### Domain: sanctions_screening/

```sql
-- =========================================
-- MAIN TABLE: admin_sanctions_screenings
-- Domain: internal/domain/sanctions_screening/
-- Entity: sanctions_screening/entity.go
-- =========================================

CREATE TABLE admin_sanctions_screenings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Screening Details
    screening_number VARCHAR(20) NOT NULL UNIQUE, -- e.g., SANC-2025-001234
    
    -- Subject
    subject_type VARCHAR(20) NOT NULL CHECK (
        subject_type IN ('USER', 'BUSINESS', 'TRANSACTION')
    ),
    subject_id UUID NOT NULL,
    subject_name VARCHAR(255) NOT NULL,
    
    -- Screening Configuration
    screening_lists TEXT[], -- e.g., ["OFAC", "UN", "EU", "INTERPOL"]
    screening_provider VARCHAR(50) NOT NULL,
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('PENDING', 'IN_PROGRESS', 'COMPLETED', 'FAILED')
    ) DEFAULT 'PENDING',
    
    -- Results
    match_found BOOLEAN DEFAULT FALSE,
    hit_count INTEGER DEFAULT 0,
    confidence_score DECIMAL(5,4),
    
    -- Screening Details
    screening_started_at TIMESTAMPTZ,
    screening_completed_at TIMESTAMPTZ,
    screening_duration_ms INTEGER,
    
    -- Provider Response
    provider_response JSONB,
    provider_reference VARCHAR(255),
    
    -- Triggered By
    triggered_by VARCHAR(30) CHECK (
        triggered_by IN ('USER_REGISTRATION', 'PROFILE_UPDATE', 'PERIODIC_CHECK', 'MANUAL_REQUEST', 'TRANSACTION')
    ),
    triggered_by_admin_id UUID REFERENCES admin_users(id),
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_sanctions_screenings_number ON admin_sanctions_screenings (screening_number);
CREATE INDEX idx_admin_sanctions_screenings_subject ON admin_sanctions_screenings (subject_type, subject_id);
CREATE INDEX idx_admin_sanctions_screenings_status ON admin_sanctions_screenings (status, created_at DESC);
CREATE INDEX idx_admin_sanctions_screenings_matches ON admin_sanctions_screenings (match_found) WHERE match_found = TRUE;

COMMENT ON TABLE admin_sanctions_screenings IS 'Sanctions screening runs - maps to internal/domain/sanctions_screening/entity.go';

-- =========================================
-- SUB-ENTITY: sanctions_hits
-- Domain: internal/domain/sanctions_screening/
-- Entity: sanctions_screening/hit.go
-- =========================================

CREATE TABLE admin_sanctions_hits (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Screening Reference
    screening_id UUID NOT NULL REFERENCES admin_sanctions_screenings(id) ON DELETE CASCADE,
    
    -- Hit Details
    list_name VARCHAR(100) NOT NULL, -- e.g., "OFAC SDN List"
    list_entry_id VARCHAR(255) NOT NULL,
    entity_name VARCHAR(255) NOT NULL,
    
    -- Match Details
    match_score DECIMAL(5,4) NOT NULL,
    match_criteria JSONB, -- Details of what matched (name, DOB, etc.)
    
    -- List Entry Details
    entry_type VARCHAR(50),
    entry_reason TEXT,
    entry_date DATE,
    entry_source VARCHAR(100),
    
    -- Disposition
    disposition VARCHAR(30) CHECK (
        disposition IN ('PENDING_REVIEW', 'FALSE_POSITIVE', 'TRUE_POSITIVE', 'ESCALATED', 'CLEARED')
    ) DEFAULT 'PENDING_REVIEW',
    
    -- Review
    reviewed_at TIMESTAMPTZ,
    reviewed_by UUID REFERENCES admin_users(id),
    review_notes TEXT,
    
    -- Disposition Details
    disposition_rationale TEXT,
    disposition_updated_at TIMESTAMPTZ,
    disposition_updated_by UUID REFERENCES admin_users(id),
    
    -- Timestamps
    detected_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_sanctions_hits_screening ON admin_sanctions_hits (screening_id);
CREATE INDEX idx_sanctions_hits_disposition ON admin_sanctions_hits (disposition) WHERE disposition != 'FALSE_POSITIVE';
CREATE INDEX idx_sanctions_hits_list ON admin_sanctions_hits (list_name, entity_name);
CREATE INDEX idx_sanctions_hits_detected ON admin_sanctions_hits (detected_at DESC);

COMMENT ON TABLE admin_sanctions_hits IS 'Sanctions screening hits - maps to internal/domain/sanctions_screening/hit.go';
```

---

## **SECTION 8: BILLING REMEDIES**

### Domain: refund_case/

```sql
-- =========================================
-- MAIN TABLE: admin_refund_cases
-- Domain: internal/domain/refund_case/
-- Entity: refund_case/entity.go
-- =========================================

CREATE TABLE admin_refund_cases (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Case Identification
    case_number VARCHAR(20) NOT NULL UNIQUE, -- e.g., REF-2025-001234
    
    -- Requester
    requester_user_id UUID NOT NULL,
    requester_email CITEXT NOT NULL,
    
    -- Payment Reference
    payment_id UUID NOT NULL,
    payment_amount BIGINT NOT NULL,
    payment_currency CHAR(3) DEFAULT 'USD',
    payment_date TIMESTAMPTZ NOT NULL,
    
    -- Related Entities
    related_contract_id UUID,
    related_job_id UUID,
    related_dispute_id UUID,
    
    -- Refund Details
    requested_amount BIGINT NOT NULL,
    refund_reason VARCHAR(50) NOT NULL CHECK (
        refund_reason IN ('DUPLICATE_CHARGE', 'UNAUTHORIZED_CHARGE', 'SERVICE_NOT_DELIVERED', 'POOR_QUALITY', 'CANCELLED_CONTRACT', 'TECHNICAL_ERROR', 'GOODWILL', 'OTHER')
    ),
    refund_description TEXT NOT NULL,
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('SUBMITTED', 'UNDER_REVIEW', 'APPROVED', 'REJECTED', 'PROCESSING', 'COMPLETED', 'FAILED')
    ) DEFAULT 'SUBMITTED',
    
    -- Review
    assigned_reviewer_id UUID REFERENCES admin_users(id),
    assigned_at TIMESTAMPTZ,
    reviewed_at TIMESTAMPTZ,
    reviewer_notes TEXT,
    
    -- Decision
    decision VARCHAR(20) CHECK (
        decision IN ('APPROVED_FULL', 'APPROVED_PARTIAL', 'REJECTED')
    ),
    approved_amount BIGINT,
    decision_rationale TEXT,
    decision_made_at TIMESTAMPTZ,
    decision_made_by UUID REFERENCES admin_users(id),
    
    -- Rejection Details
    rejection_reason VARCHAR(100),
    rejection_notes TEXT,
    
    -- Processing
    refund_method VARCHAR(30) CHECK (
        refund_method IN ('ORIGINAL_PAYMENT_METHOD', 'BANK_TRANSFER', 'WALLET_CREDIT', 'CHECK')
    ),
    refund_reference VARCHAR(255),
    refund_processed_at TIMESTAMPTZ,
    refund_completed_at TIMESTAMPTZ,
    
    -- Failure
    failure_reason TEXT,
    failed_at TIMESTAMPTZ,
    
    -- SLA
    sla_target_decision_at TIMESTAMPTZ,
    sla_breached BOOLEAN DEFAULT FALSE,
    
    -- Timestamps
    submitted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    closed_at TIMESTAMPTZ,
    last_updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_refund_cases_number ON admin_refund_cases (case_number);
CREATE INDEX idx_admin_refund_cases_requester ON admin_refund_cases (requester_user_id, submitted_at DESC);
CREATE INDEX idx_admin_refund_cases_payment ON admin_refund_cases (payment_id);
CREATE INDEX idx_admin_refund_cases_status ON admin_refund_cases (status, submitted_at DESC);
CREATE INDEX idx_admin_refund_cases_assigned ON admin_refund_cases (assigned_reviewer_id, status) WHERE assigned_reviewer_id IS NOT NULL;
CREATE INDEX idx_admin_refund_cases_sla ON admin_refund_cases (sla_breached) WHERE sla_breached = TRUE;

COMMENT ON TABLE admin_refund_cases IS 'Refund cases - maps to internal/domain/refund_case/entity.go';
```

### Domain: goodwill_credit/

```sql
-- =========================================
-- MAIN TABLE: admin_goodwill_credits
-- Domain: internal/domain/goodwill_credit/
-- Entity: goodwill_credit/entity.go
-- =========================================

CREATE TABLE admin_goodwill_credits (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Credit Details
    credit_number VARCHAR(20) NOT NULL UNIQUE, -- e.g., GW-2025-001234
    
    -- Recipient
    recipient_user_id UUID NOT NULL,
    recipient_email CITEXT NOT NULL,
    
    -- Amount
    credit_amount BIGINT NOT NULL,
    credit_currency CHAR(3) DEFAULT 'USD',
    
    -- Reason
    reason_code VARCHAR(50) NOT NULL CHECK (
        reason_code IN ('SERVICE_ISSUE', 'PLATFORM_ERROR', 'COMPENSATION', 'APOLOGY', 'RETENTION', 'PROMOTION', 'OTHER')
    ),
    reason_description TEXT NOT NULL,
    
    -- Context
    related_ticket_id UUID REFERENCES support_tickets(id),
    related_case_id UUID,
    related_case_type VARCHAR(30),
    
    -- Restrictions
    expiry_date DATE,
    usage_restrictions JSONB, -- e.g., {"usable_for": ["JOB_POSTS", "CONNECTS"], "min_transaction": 1000}
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('PENDING_APPROVAL', 'APPROVED', 'REJECTED', 'ISSUED', 'USED', 'EXPIRED', 'REVOKED')
    ) DEFAULT 'PENDING_APPROVAL',
    
    -- Approval
    requires_approval BOOLEAN DEFAULT TRUE,
    approved_by UUID REFERENCES admin_users(id),
    approved_at TIMESTAMPTZ,
    approval_notes TEXT,
    
    rejected_by UUID REFERENCES admin_users(id),
    rejected_at TIMESTAMPTZ,
    rejection_reason TEXT,
    
    -- Issuance
    issued_by UUID NOT NULL REFERENCES admin_users(id),
    issued_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Usage
    used_amount BIGINT DEFAULT 0,
    remaining_amount BIGINT,
    fully_used_at TIMESTAMPTZ,
    
    -- Revocation
    revoked_at TIMESTAMPTZ,
    revoked_by UUID REFERENCES admin_users(id),
    revocation_reason TEXT,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at TIMESTAMPTZ,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_goodwill_credits_number ON admin_goodwill_credits (credit_number);
CREATE INDEX idx_admin_goodwill_credits_recipient ON admin_goodwill_credits (recipient_user_id, created_at DESC);
CREATE INDEX idx_admin_goodwill_credits_status ON admin_goodwill_credits (status, created_at DESC);
CREATE INDEX idx_admin_goodwill_credits_issued_by ON admin_goodwill_credits (issued_by, created_at DESC);
CREATE INDEX idx_admin_goodwill_credits_expires ON admin_goodwill_credits (expires_at) WHERE status = 'ISSUED';

COMMENT ON TABLE admin_goodwill_credits IS 'Goodwill credits - maps to internal/domain/goodwill_credit/entity.go';
```

---

## **SECTION 9: CONFIGURATION & POLICY**

### Domain: system_config/

```sql
-- =========================================
-- MAIN TABLE: admin_system_configs
-- Domain: internal/domain/system_config/
-- Entity: system_config/entity.go
-- =========================================

CREATE TABLE admin_system_configs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Config Key
    config_key VARCHAR(200) NOT NULL UNIQUE,
    config_category VARCHAR(50) NOT NULL,
    
    -- Value
    config_value TEXT NOT NULL,
    config_value_type VARCHAR(20) CHECK (
        config_value_type IN ('STRING', 'INTEGER', 'FLOAT', 'BOOLEAN', 'JSON', 'ARRAY')
    ) DEFAULT 'STRING',
    
    -- Description
    display_name VARCHAR(255),
    description TEXT,
    
    -- Validation
    validation_rules JSONB,
    default_value TEXT,
    
    -- Environment
    environment VARCHAR(20) CHECK (
        environment IN ('ALL', 'PRODUCTION', 'STAGING', 'DEVELOPMENT')
    ) DEFAULT 'ALL',
    
    -- Feature Flag
    is_feature_flag BOOLEAN DEFAULT FALSE,
    flag_enabled BOOLEAN DEFAULT FALSE,
    flag_rollout_percentage INTEGER CHECK (flag_rollout_percentage >= 0 AND flag_rollout_percentage <= 100),
    flag_user_ids UUID[], -- Array of user IDs for targeted rollout
    
    -- Versioning
    version INTEGER DEFAULT 1,
    previous_value TEXT,
    
    -- Change Management
    requires_approval BOOLEAN DEFAULT FALSE,
    approved_by UUID REFERENCES admin_users(id),
    approved_at TIMESTAMPTZ,
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    is_sensitive BOOLEAN DEFAULT FALSE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by UUID REFERENCES admin_users(id),
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_by UUID REFERENCES admin_users(id),
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    tags TEXT[]
);

CREATE INDEX idx_admin_system_configs_key ON admin_system_configs (config_key);
CREATE INDEX idx_admin_system_configs_category ON admin_system_configs (config_category, is_active);
CREATE INDEX idx_admin_system_configs_feature_flags ON admin_system_configs (is_feature_flag) WHERE is_feature_flag = TRUE;
CREATE INDEX idx_admin_system_configs_environment ON admin_system_configs (environment);
CREATE INDEX idx_admin_system_configs_tags ON admin_system_configs USING gin(tags);

COMMENT ON TABLE admin_system_configs IS 'System configuration - maps to internal/domain/system_config/entity.go';

-- =========================================
-- SUB-ENTITY: config_change_history
-- Domain: internal/domain/system_config/
-- Entity: system_config/version.go
-- =========================================

CREATE TABLE admin_config_change_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Config Reference
    config_id UUID NOT NULL REFERENCES admin_system_configs(id),
    config_key VARCHAR(200) NOT NULL,
    
    -- Change Details
    version INTEGER NOT NULL,
    old_value TEXT,
    new_value TEXT NOT NULL,
    change_type VARCHAR(20) CHECK (
        change_type IN ('CREATED', 'UPDATED', 'DELETED', 'ENABLED', 'DISABLED')
    ),
    
    -- Change Reason
    change_reason TEXT,
    
    -- Changed By
    changed_by UUID NOT NULL REFERENCES admin_users(id),
    
    -- Timestamp
    changed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_config_change_history_config ON admin_config_change_history (config_id, changed_at DESC);
CREATE INDEX idx_config_change_history_key ON admin_config_change_history (config_key, changed_at DESC);
CREATE INDEX idx_config_change_history_changed_by ON admin_config_change_history (changed_by);

COMMENT ON TABLE admin_config_change_history IS 'Config change history - maps to internal/domain/system_config/version.go';
```

### Domain: policy_doc/

```sql
-- =========================================
-- MAIN TABLE: admin_policy_documents
-- Domain: internal/domain/policy_doc/
-- Entity: policy_doc/entity.go
-- =========================================

CREATE TABLE admin_policy_documents (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Policy Identification
    policy_code VARCHAR(50) NOT NULL UNIQUE,
    policy_name VARCHAR(255) NOT NULL,
    
    -- Type
    policy_type VARCHAR(30) CHECK (
        policy_type IN ('TERMS_OF_SERVICE', 'PRIVACY_POLICY', 'COMMUNITY_GUIDELINES', 'PAYMENT_TERMS', 'REFUND_POLICY', 'COOKIE_POLICY', 'OTHER')
    ),
    
    -- Content
    title VARCHAR(500) NOT NULL,
    content TEXT NOT NULL,
    summary TEXT,
    
    -- Versioning
    version VARCHAR(20) NOT NULL,
    version_number INTEGER NOT NULL,
    previous_version_id UUID,
    
    -- Publishing
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('DRAFT', 'IN_REVIEW', 'APPROVED', 'PUBLISHED', 'ARCHIVED')
    ) DEFAULT 'DRAFT',
    
    -- Effective Dates
    effective_from DATE NOT NULL,
    effective_until DATE,
    
    -- Localization
    locale VARCHAR(10) DEFAULT 'en-US',
    
    -- Legal Review
    legal_review_required BOOLEAN DEFAULT TRUE,
    legal_reviewed_by UUID REFERENCES admin_users(id),
    legal_reviewed_at TIMESTAMPTZ,
    legal_review_notes TEXT,
    
    -- Change Notice
    requires_user_consent BOOLEAN DEFAULT FALSE,
    consent_deadline DATE,
    change_summary TEXT,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by UUID REFERENCES admin_users(id),
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    published_at TIMESTAMPTZ,
    published_by UUID REFERENCES admin_users(id),
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_policy_documents_code ON admin_policy_documents (policy_code, version_number DESC);
CREATE INDEX idx_admin_policy_documents_type ON admin_policy_documents (policy_type, status);
CREATE INDEX idx_admin_policy_documents_status ON admin_policy_documents (status);
CREATE INDEX idx_admin_policy_documents_effective ON admin_policy_documents (effective_from, effective_until);
CREATE INDEX idx_admin_policy_documents_locale ON admin_policy_documents (locale);

COMMENT ON TABLE admin_policy_documents IS 'Policy documents - maps to internal/domain/policy_doc/entity.go';
```

### Domain: experiment/

```sql
-- =========================================
-- MAIN TABLE: admin_experiments
-- Domain: internal/domain/experiment/
-- Entity: experiment/entity.go
-- =========================================

CREATE TABLE admin_experiments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Experiment Details
    experiment_key VARCHAR(100) NOT NULL UNIQUE,
    experiment_name VARCHAR(255) NOT NULL,
    description TEXT,
    
    -- Hypothesis
    hypothesis TEXT,
    success_metrics TEXT[],
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('DRAFT', 'SCHEDULED', 'RUNNING', 'PAUSED', 'COMPLETED', 'CANCELLED')
    ) DEFAULT 'DRAFT',
    
    -- Variants
    control_variant JSONB NOT NULL,
    treatment_variants JSONB[] NOT NULL,
    
    -- Targeting
    target_user_percentage INTEGER CHECK (target_user_percentage >= 0 AND target_user_percentage <= 100),
    target_user_segments JSONB,
    excluded_user_ids UUID[],
    
    -- Allocation
    traffic_allocation JSONB NOT NULL, -- {"control": 50, "variant_a": 25, "variant_b": 25}
    
    -- Schedule
    scheduled_start_at TIMESTAMPTZ,
    scheduled_end_at TIMESTAMPTZ,
    actual_start_at TIMESTAMPTZ,
    actual_end_at TIMESTAMPTZ,
    
    -- Results
    results_summary JSONB,
    winning_variant VARCHAR(50),
    statistical_significance DECIMAL(5,4),
    
    -- Owner
    owner_id UUID NOT NULL REFERENCES admin_users(id),
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by UUID REFERENCES admin_users(id),
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    tags TEXT[]
);

CREATE INDEX idx_admin_experiments_key ON admin_experiments (experiment_key);
CREATE INDEX idx_admin_experiments_status ON admin_experiments (status, scheduled_start_at);
CREATE INDEX idx_admin_experiments_owner ON admin_experiments (owner_id);
CREATE INDEX idx_admin_experiments_schedule ON admin_experiments (scheduled_start_at, scheduled_end_at);
CREATE INDEX idx_admin_experiments_tags ON admin_experiments USING gin(tags);

COMMENT ON TABLE admin_experiments IS 'A/B experiments - maps to internal/domain/experiment/entity.go';
```

### Domain: throttle_policy/

```sql
-- =========================================
-- MAIN TABLE: admin_throttle_policies
-- Domain: internal/domain/throttle_policy/
-- Entity: throttle_policy/entity.go
-- =========================================

CREATE TABLE admin_throttle_policies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Policy Details
    policy_name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    
    -- Scope
    resource_type VARCHAR(50) NOT NULL, -- e.g., "API_ENDPOINT", "FEATURE", "ACTION"
    resource_identifier VARCHAR(255) NOT NULL,
    
    -- Limits
    window_type VARCHAR(20) CHECK (
        window_type IN ('SECOND', 'MINUTE', 'HOUR', 'DAY')
    ) DEFAULT 'MINUTE',
    window_size INTEGER NOT NULL,
    max_requests INTEGER NOT NULL,
    
    -- Targeting
    applies_to VARCHAR(20) CHECK (
        applies_to IN ('ALL_USERS', 'USER_TYPE', 'SUBSCRIPTION_TIER', 'SPECIFIC_USERS')
    ) DEFAULT 'ALL_USERS',
    user_type VARCHAR(20),
    subscription_tier VARCHAR(50),
    specific_user_ids UUID[],
    
    -- Behavior
    block_on_exceed BOOLEAN DEFAULT TRUE,
    retry_after_seconds INTEGER,
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by UUID REFERENCES admin_users(id),
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_by UUID REFERENCES admin_users(id),
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_throttle_policies_name ON admin_throttle_policies (policy_name);
CREATE INDEX idx_admin_throttle_policies_resource ON admin_throttle_policies (resource_type, resource_identifier);
CREATE INDEX idx_admin_throttle_policies_active ON admin_throttle_policies (is_active) WHERE is_active = TRUE;

COMMENT ON TABLE admin_throttle_policies IS 'Throttle policies - maps to internal/domain/throttle_policy/entity.go';

-- =========================================
-- SUB-ENTITY: throttle_policy_exceptions
-- Domain: internal/domain/throttle_policy/
-- Entity: throttle_policy/exception.go
-- =========================================

CREATE TABLE admin_throttle_policy_exceptions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Policy Reference
    policy_id UUID NOT NULL REFERENCES admin_throttle_policies(id),
    
    -- Exception Details
    exception_type VARCHAR(20) CHECK (
        exception_type IN ('USER', 'IP', 'API_KEY')
    ) NOT NULL,
    exception_value VARCHAR(255) NOT NULL, -- User ID, IP address, or API key
    
    -- Override Limits
    override_max_requests INTEGER,
    override_window_size INTEGER,
    
    -- Reason
    reason TEXT NOT NULL,
    
    -- Duration
    effective_from TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    effective_until TIMESTAMPTZ,
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    
    -- Created By
    created_by UUID NOT NULL REFERENCES admin_users(id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    
    CONSTRAINT uk_throttle_exception UNIQUE (policy_id, exception_type, exception_value)
);

CREATE INDEX idx_throttle_policy_exceptions_policy ON admin_throttle_policy_exceptions (policy_id);
CREATE INDEX idx_throttle_policy_exceptions_type ON admin_throttle_policy_exceptions (exception_type, exception_value);
CREATE INDEX idx_throttle_policy_exceptions_effective ON admin_throttle_policy_exceptions (effective_until) WHERE is_active = TRUE;

COMMENT ON TABLE admin_throttle_policy_exceptions IS 'Throttle exceptions - maps to internal/domain/throttle_policy/exception.go';
```

### Domain: quota_override/

```sql
-- =========================================
-- MAIN TABLE: admin_quota_overrides
-- Domain: internal/domain/quota_override/
-- Entity: quota_override/entity.go
-- =========================================

CREATE TABLE admin_quota_overrides (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Subject
    subject_user_id UUID NOT NULL,
    subject_email CITEXT NOT NULL,
    
    -- Feature/Resource
    feature_code VARCHAR(100) NOT NULL,
    feature_name VARCHAR(255) NOT NULL,
    
    -- Override Details
    override_type VARCHAR(20) CHECK (
        override_type IN ('INCREASE', 'DECREASE', 'UNLIMITED')
    ) NOT NULL,
    
    -- Limits
    original_limit INTEGER,
    override_limit INTEGER,
    
    -- Reason
    reason_code VARCHAR(50) NOT NULL,
    reason_description TEXT NOT NULL,
    justification TEXT,
    
    -- Related Case
    related_ticket_id UUID REFERENCES support_tickets(id),
    
    -- Duration
    effective_from TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    effective_until TIMESTAMPTZ NOT NULL,
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('ACTIVE', 'EXPIRED', 'REVOKED')
    ) DEFAULT 'ACTIVE',
    
    -- Revocation
    revoked_at TIMESTAMPTZ,
    revoked_by UUID REFERENCES admin_users(id),
    revocation_reason TEXT,
    
    -- Applied By
    applied_by UUID NOT NULL REFERENCES admin_users(id),
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_quota_overrides_subject ON admin_quota_overrides (subject_user_id, status);
CREATE INDEX idx_admin_quota_overrides_feature ON admin_quota_overrides (feature_code, status);
CREATE INDEX idx_admin_quota_overrides_status ON admin_quota_overrides (status, effective_until);
CREATE INDEX idx_admin_quota_overrides_expires ON admin_quota_overrides (effective_until) WHERE status = 'ACTIVE';

COMMENT ON TABLE admin_quota_overrides IS 'Quota overrides - maps to internal/domain/quota_override/entity.go';
```

---

## **SECTION 10: RISK, FRAUD & INCIDENTS**

### Domain: fraud_review/

```sql
-- =========================================
-- MAIN TABLE: admin_fraud_reviews
-- Domain: internal/domain/fraud_review/
-- Entity: fraud_review/entity.go
-- =========================================

CREATE TABLE admin_fraud_reviews (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Case Identification
    case_number VARCHAR(20) NOT NULL UNIQUE, -- e.g., FRD-2025-001234
    
    -- Subject
    subject_type VARCHAR(20) NOT NULL CHECK (
        subject_type IN ('USER', 'TRANSACTION', 'PAYMENT', 'ACCOUNT')
    ),
    subject_id UUID NOT NULL,
    subject_user_id UUID,
    
    -- Risk Signals
    risk_signals JSONB NOT NULL, -- Array of risk indicators
    risk_score INTEGER CHECK (risk_score >= 0 AND risk_score <= 100),
    
    -- Severity
    severity VARCHAR(20) NOT NULL CHECK (
        severity IN ('LOW', 'MEDIUM', 'HIGH', 'CRITICAL')
    ) DEFAULT 'MEDIUM',
    
    -- Trigger Reason
    trigger_reason VARCHAR(50) NOT NULL CHECK (
        trigger_reason IN ('VELOCITY_ALERT', 'PATTERN_ANOMALY', 'CHARGEBACK_RECEIVED', 'USER_REPORT', 'AUTO_FLAG', 'MANUAL_REVIEW')
    ),
    trigger_description TEXT,
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('PENDING', 'INVESTIGATING', 'CLEARED', 'CONFIRMED_FRAUD', 'ESCALATED', 'CLOSED')
    ) DEFAULT 'PENDING',
    
    -- Assignment
    assigned_investigator_id UUID REFERENCES admin_users(id),
    assigned_at TIMESTAMPTZ,
    
    -- Investigation
    investigation_notes TEXT[],
    evidence_refs JSONB,
    
    -- Decision
    decision VARCHAR(30) CHECK (
        decision IN ('CLEARED', 'FRAUD_CONFIRMED', 'SUSPICIOUS_PENDING', 'FALSE_POSITIVE')
    ),
    decision_rationale TEXT,
    decision_made_at TIMESTAMPTZ,
    decision_made_by UUID REFERENCES admin_users(id),
    
    -- Actions Taken
    actions_taken TEXT[],
    user_notified BOOLEAN DEFAULT FALSE,
    account_restricted BOOLEAN DEFAULT FALSE,
    
    -- SLA
    sla_target_review_at TIMESTAMPTZ,
    sla_breached BOOLEAN DEFAULT FALSE,
    
    -- Timestamps
    opened_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    closed_at TIMESTAMPTZ,
    last_updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_fraud_reviews_number ON admin_fraud_reviews (case_number);
CREATE INDEX idx_admin_fraud_reviews_subject ON admin_fraud_reviews (subject_type, subject_id);
CREATE INDEX idx_admin_fraud_reviews_user ON admin_fraud_reviews (subject_user_id) WHERE subject_user_id IS NOT NULL;
CREATE INDEX idx_admin_fraud_reviews_status ON admin_fraud_reviews (status, severity, opened_at DESC);
CREATE INDEX idx_admin_fraud_reviews_assigned ON admin_fraud_reviews (assigned_investigator_id, status) WHERE assigned_investigator_id IS NOT NULL;
CREATE INDEX idx_admin_fraud_reviews_sla ON admin_fraud_reviews (sla_breached) WHERE sla_breached = TRUE;

COMMENT ON TABLE admin_fraud_reviews IS 'Fraud investigation cases - maps to internal/domain/fraud_review/entity.go';
```

### Domain: risk_management/

```sql
-- =========================================
-- MAIN TABLE: admin_risk_holds
-- Domain: internal/domain/risk_management/
-- Entity: risk_management/hold.go
-- =========================================

CREATE TABLE admin_risk_holds (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Hold Details
    hold_type VARCHAR(30) NOT NULL CHECK (
        hold_type IN ('PAYMENT_HOLD', 'PAYOUT_HOLD', 'ACCOUNT_HOLD', 'TRANSACTION_HOLD')
    ),
    
    -- Subject
    subject_user_id UUID NOT NULL,
    subject_account_id UUID,
    
    -- Amount (if applicable)
    hold_amount BIGINT,
    hold_currency CHAR(3) DEFAULT 'USD',
    
    -- Reason
    reason_code VARCHAR(50) NOT NULL,
    reason_description TEXT NOT NULL,
    risk_factors JSONB,
    
    -- Related Entities
    related_transaction_id UUID,
    related_payment_id UUID,
    related_fraud_review_id UUID REFERENCES admin_fraud_reviews(id),
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('ACTIVE', 'RELEASED', 'EXPIRED', 'ESCALATED')
    ) DEFAULT 'ACTIVE',
    
    -- Release
    released_at TIMESTAMPTZ,
    released_by UUID REFERENCES admin_users(id),
    release_reason TEXT,
    
    -- Placed By
    placed_by UUID NOT NULL REFERENCES admin_users(id),
    
    -- Duration
    placed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at TIMESTAMPTZ,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_risk_holds_subject ON admin_risk_holds (subject_user_id, status);
CREATE INDEX idx_admin_risk_holds_type ON admin_risk_holds (hold_type, status);
CREATE INDEX idx_admin_risk_holds_status ON admin_risk_holds (status, placed_at DESC);
CREATE INDEX idx_admin_risk_holds_expires ON admin_risk_holds (expires_at) WHERE status = 'ACTIVE';

COMMENT ON TABLE admin_risk_holds IS 'Risk holds - maps to internal/domain/risk_management/hold.go';
```

### Domain: incident/

```sql
-- =========================================
-- MAIN TABLE: admin_incidents
-- Domain: internal/domain/incident/
-- Entity: incident/entity.go
-- =========================================

CREATE TABLE admin_incidents (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Incident Identification
    incident_number VARCHAR(20) NOT NULL UNIQUE, -- e.g., INC-2025-001234
    
    -- Incident Details
    title VARCHAR(500) NOT NULL,
    description TEXT NOT NULL,
    
    -- Severity
    severity VARCHAR(20) NOT NULL CHECK (
        severity IN ('SEV1', 'SEV2', 'SEV3', 'SEV4')
    ),
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('IDENTIFIED', 'INVESTIGATING', 'MITIGATING', 'RESOLVED', 'CLOSED')
    ) DEFAULT 'IDENTIFIED',
    
    -- Impact
    impact_description TEXT,
    affected_services TEXT[],
    affected_user_count INTEGER,
    
    -- Incident Commander
    commander_id UUID REFERENCES admin_users(id),
    assigned_at TIMESTAMPTZ,
    
    -- Timeline
    detected_at TIMESTAMPTZ NOT NULL,
    acknowledged_at TIMESTAMPTZ,
    mitigated_at TIMESTAMPTZ,
    resolved_at TIMESTAMPTZ,
    closed_at TIMESTAMPTZ,
    
    -- Duration
    time_to_detect_minutes INTEGER,
    time_to_acknowledge_minutes INTEGER,
    time_to_mitigate_minutes INTEGER,
    time_to_resolve_minutes INTEGER,
    
    -- Root Cause Analysis
    rca_completed BOOLEAN DEFAULT FALSE,
    rca_summary TEXT,
    root_cause TEXT,
    contributing_factors TEXT[],
    
    -- Action Items
    action_items JSONB,
    action_items_completed INTEGER DEFAULT 0,
    action_items_total INTEGER DEFAULT 0,
    
    -- Communications
    public_status_page_updated BOOLEAN DEFAULT FALSE,
    customer_notification_sent BOOLEAN DEFAULT FALSE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    tags TEXT[]
);

CREATE INDEX idx_admin_incidents_number ON admin_incidents (incident_number);
CREATE INDEX idx_admin_incidents_status ON admin_incidents (status, severity, detected_at DESC);
CREATE INDEX idx_admin_incidents_commander ON admin_incidents (commander_id, status) WHERE commander_id IS NOT NULL;
CREATE INDEX idx_admin_incidents_detected ON admin_incidents (detected_at DESC);
CREATE INDEX idx_admin_incidents_tags ON admin_incidents USING gin(tags);

COMMENT ON TABLE admin_incidents IS 'Operational incidents - maps to internal/domain/incident/entity.go';

-- =========================================
-- SUB-ENTITY: incident_timeline_events
-- Domain: internal/domain/incident/
-- Entity: incident/timeline_event.go
-- =========================================

CREATE TABLE admin_incident_timeline_events (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Incident Reference
    incident_id UUID NOT NULL REFERENCES admin_incidents(id) ON DELETE CASCADE,
    
    -- Event Details
    event_type VARCHAR(30) CHECK (
        event_type IN ('STATUS_CHANGE', 'INVESTIGATION_UPDATE', 'MITIGATION_APPLIED', 'COMMUNICATION', 'ACTION_TAKEN', 'NOTE')
    ) NOT NULL,
    description TEXT NOT NULL,
    
    -- Status Change (if applicable)
    old_status VARCHAR(20),
    new_status VARCHAR(20),
    
    -- Created By
    created_by UUID REFERENCES admin_users(id),
    created_by_name VARCHAR(255),
    
    -- Timestamp
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_incident_timeline_incident ON admin_incident_timeline_events (incident_id, created_at);
CREATE INDEX idx_incident_timeline_type ON admin_incident_timeline_events (event_type);

COMMENT ON TABLE admin_incident_timeline_events IS 'Incident timeline - maps to internal/domain/incident/timeline_event.go';
```

---

## **SECTION 11: INTEGRATIONS & BULK OPS**

### Domain: integrations_admin/

```sql
-- =========================================
-- MAIN TABLE: admin_integrations
-- Domain: internal/domain/integrations_admin/
-- Entity: integrations_admin/entity.go
-- =========================================

CREATE TABLE admin_integrations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Integration Details
    integration_name VARCHAR(255) NOT NULL UNIQUE,
    vendor_name VARCHAR(255) NOT NULL,
    
    -- Type
    integration_type VARCHAR(50) CHECK (
        integration_type IN ('PAYMENT_GATEWAY', 'IDENTITY_PROVIDER', 'ANALYTICS', 'MONITORING', 'COMMUNICATION', 'STORAGE', 'OTHER')
    ) NOT NULL,
    
    -- Configuration
    endpoints JSONB NOT NULL,
    scopes TEXT[],
    configuration JSONB,
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('ACTIVE', 'DISABLED', 'TESTING', 'DEPRECATED')
    ) DEFAULT 'TESTING',
    
    -- Health
    last_health_check_at TIMESTAMPTZ,
    health_status VARCHAR(20) CHECK (
        health_status IN ('HEALTHY', 'DEGRADED', 'DOWN', 'UNKNOWN')
    ),
    
    -- Credentials (encrypted)
    credentials_encrypted TEXT,
    credentials_last_rotated_at TIMESTAMPTZ,
    
    -- Added By
    added_by UUID NOT NULL REFERENCES admin_users(id),
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_integrations_name ON admin_integrations (integration_name);
CREATE INDEX idx_admin_integrations_vendor ON admin_integrations (vendor_name);
CREATE INDEX idx_admin_integrations_type ON admin_integrations (integration_type, status);
CREATE INDEX idx_admin_integrations_status ON admin_integrations (status);

COMMENT ON TABLE admin_integrations IS 'Third-party integrations - maps to internal/domain/integrations_admin/entity.go';

-- =========================================
-- SUB-ENTITY: admin_api_keys
-- Domain: internal/domain/integrations_admin/
-- Entity: integrations_admin/api_key.go
-- =========================================

CREATE TABLE admin_api_keys (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Integration Reference
    integration_id UUID REFERENCES admin_integrations(id),
    
    -- Key Details
    key_name VARCHAR(255) NOT NULL,
    key_hash VARCHAR(64) NOT NULL UNIQUE, -- SHA-256 hash of key
    key_prefix VARCHAR(20), -- First few chars for identification
    
    -- Scopes & Permissions
    scopes TEXT[],
    permissions JSONB,
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('ACTIVE', 'REVOKED', 'EXPIRED')
    ) DEFAULT 'ACTIVE',
    
    -- Usage
    last_used_at TIMESTAMPTZ,
    usage_count INTEGER DEFAULT 0,
    
    -- Expiry & Rotation
    expires_at TIMESTAMPTZ,
    rotation_policy VARCHAR(20) CHECK (
        rotation_policy IN ('NEVER', '30_DAYS', '90_DAYS', '180_DAYS', '365_DAYS')
    ),
    last_rotated_at TIMESTAMPTZ,
    
    -- Revocation
    revoked_at TIMESTAMPTZ,
    revoked_by UUID REFERENCES admin_users(id),
    revocation_reason TEXT,
    
    -- Issued By
    issued_by UUID NOT NULL REFERENCES admin_users(id),
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_api_keys_integration ON admin_api_keys (integration_id, status);
CREATE INDEX idx_admin_api_keys_hash ON admin_api_keys (key_hash);
CREATE INDEX idx_admin_api_keys_status ON admin_api_keys (status);
CREATE INDEX idx_admin_api_keys_expires ON admin_api_keys (expires_at) WHERE status = 'ACTIVE';

COMMENT ON TABLE admin_api_keys IS 'API keys - maps to internal/domain/integrations_admin/api_key.go';
```

### Domain: bulk_action/

```sql
-- =========================================
-- MAIN TABLE: admin_bulk_actions
-- Domain: internal/domain/bulk_action/
-- Entity: bulk_action/entity.go
-- =========================================

CREATE TABLE admin_bulk_actions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Action Identification
    action_number VARCHAR(20) NOT NULL UNIQUE, -- e.g., BULK-2025-001234
    
    -- Action Details
    action_type VARCHAR(50) NOT NULL CHECK (
        action_type IN ('USER_SUSPEND', 'USER_ACTIVATE', 'CONTENT_REMOVE', 'CONTENT_RESTORE', 'EMAIL_SEND', 'CREDIT_ISSUE', 'OTHER')
    ),
    action_description TEXT NOT NULL,
    
    -- Target Query
    target_query JSONB NOT NULL, -- Criteria for selecting targets
    target_count_preview INTEGER,
    target_ids UUID[], -- Populated after preview
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('DRAFT', 'PREVIEWED', 'APPROVED', 'EXECUTING', 'COMPLETED', 'FAILED', 'ROLLED_BACK')
    ) DEFAULT 'DRAFT',
    
    -- Execution
    batch_size INTEGER DEFAULT 100,
    total_batches INTEGER,
    completed_batches INTEGER DEFAULT 0,
    successful_count INTEGER DEFAULT 0,
    failed_count INTEGER DEFAULT 0,
    progress_percentage DECIMAL(5,2),
    
    -- Rollback
    supports_rollback BOOLEAN DEFAULT FALSE,
    rollback_plan JSONB,
    rolled_back_at TIMESTAMPTZ,
    rolled_back_by UUID REFERENCES admin_users(id),
    
    -- Safety Checks
    requires_approval BOOLEAN DEFAULT TRUE,
    dry_run_completed BOOLEAN DEFAULT FALSE,
    safety_checks_passed BOOLEAN DEFAULT FALSE,
    
    -- Approval
    approved_by UUID REFERENCES admin_users(id),
    approved_at TIMESTAMPTZ,
    approval_notes TEXT,
    
    -- Initiated By
    initiated_by UUID NOT NULL REFERENCES admin_users(id),
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    started_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    error_log TEXT
);

CREATE INDEX idx_admin_bulk_actions_number ON admin_bulk_actions (action_number);
CREATE INDEX idx_admin_bulk_actions_status ON admin_bulk_actions (status, created_at DESC);
CREATE INDEX idx_admin_bulk_actions_type ON admin_bulk_actions (action_type, status);
CREATE INDEX idx_admin_bulk_actions_initiated_by ON admin_bulk_actions (initiated_by, created_at DESC);

COMMENT ON TABLE admin_bulk_actions IS 'Bulk operations - maps to internal/domain/bulk_action/entity.go';
```

---

## **SECTION 12: SESSIONS & APPROVALS**

### Domain: admin_session/

```sql
-- =========================================
-- MAIN TABLE: admin_sessions
-- Domain: internal/domain/admin_session/
-- Entity: admin_session/entity.go
-- =========================================

CREATE TABLE admin_sessions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Session Details
    session_token_hash VARCHAR(64) NOT NULL UNIQUE,
    
    -- Admin Reference
    admin_user_id UUID NOT NULL REFERENCES admin_users(id),
    admin_email CITEXT NOT NULL,
    admin_role VARCHAR(50) NOT NULL,
    
    -- Access Type
    session_type VARCHAR(20) CHECK (
        session_type IN ('NORMAL', 'BREAK_GLASS', 'ELEVATED', 'JIT')
    ) DEFAULT 'NORMAL',
    
    -- Scope Grants
    granted_scopes JSONB, -- {"resources": ["pii.user.view", "user.suspend"]}
    granted_permissions BIGINT,
    
    -- Justification (for break-glass)
    access_reason VARCHAR(50) CHECK (
        access_reason IN ('SUPPORT', 'INCIDENT', 'INVESTIGATION', 'AUDIT', 'LEGAL_REQUIREMENT', 'EMERGENCY')
    ),
    justification TEXT,
    
    -- Related Case
    related_case_type VARCHAR(30),
    related_case_id UUID,
    
    -- Approval (for break-glass)
    requires_approval BOOLEAN DEFAULT FALSE,
    approved_by UUID REFERENCES admin_users(id),
    approved_at TIMESTAMPTZ,
    
    -- Session Management
    ip_address INET NOT NULL,
    user_agent TEXT,
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('ACTIVE', 'EXPIRED', 'TERMINATED', 'REVOKED')
    ) DEFAULT 'ACTIVE',
    
    -- Activity Tracking
    last_activity_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    activity_count INTEGER DEFAULT 0,
    
    -- Expiry
    started_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    terminated_at TIMESTAMPTZ,
    terminated_by UUID REFERENCES admin_users(id),
    termination_reason TEXT,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_sessions_admin_user ON admin_sessions (admin_user_id, status);
CREATE INDEX idx_admin_sessions_token ON admin_sessions (session_token_hash);
CREATE INDEX idx_admin_sessions_status ON admin_sessions (status, expires_at);
CREATE INDEX idx_admin_sessions_type ON admin_sessions (session_type, status);
CREATE INDEX idx_admin_sessions_expires ON admin_sessions (expires_at) WHERE status = 'ACTIVE';

COMMENT ON TABLE admin_sessions IS 'Admin sessions - maps to internal/domain/admin_session/entity.go';
COMMENT ON COLUMN admin_sessions.session_type IS 'BREAK_GLASS for emergency PII access, JIT for just-in-time elevated privileges';
```

### Domain: change_approval/

```sql
-- =========================================
-- MAIN TABLE: admin_change_approvals
-- Domain: internal/domain/change_approval/
-- Entity: change_approval/entity.go
-- =========================================

CREATE TABLE admin_change_approvals (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Change Request Details
    request_number VARCHAR(20) NOT NULL UNIQUE, -- e.g., CHG-2025-001234
    
    -- Change Type
    change_type VARCHAR(50) NOT NULL CHECK (
        change_type IN ('CONFIG_UPDATE', 'POLICY_CHANGE', 'PERMISSION_GRANT', 'SYSTEM_SETTING', 'FINANCIAL_ADJUSTMENT', 'USER_ACTION', 'OTHER')
    ),
    
    -- Resource Being Changed
    resource_type VARCHAR(50) NOT NULL,
    resource_id UUID,
    resource_identifier VARCHAR(255),
    
    -- Change Details
    change_description TEXT NOT NULL,
    change_diff JSONB, -- Before/after comparison
    risk_level VARCHAR(20) CHECK (
        risk_level IN ('LOW', 'MEDIUM', 'HIGH', 'CRITICAL')
    ) NOT NULL,
    risk_assessment TEXT,
    
    -- Requester
    requested_by UUID NOT NULL REFERENCES admin_users(id),
    requester_role VARCHAR(50),
    justification TEXT NOT NULL,
    
    -- Approval Policy
    required_approvals_count INTEGER DEFAULT 1,
    current_approvals_count INTEGER DEFAULT 0,
    required_approver_roles BIGINT, -- Bitset of required roles
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('PENDING', 'APPROVED', 'REJECTED', 'EXPIRED', 'IMPLEMENTED', 'ROLLED_BACK')
    ) DEFAULT 'PENDING',
    
    -- Implementation
    implemented_at TIMESTAMPTZ,
    implemented_by UUID REFERENCES admin_users(id),
    implementation_notes TEXT,
    
    -- Rollback
    rolled_back_at TIMESTAMPTZ,
    rolled_back_by UUID REFERENCES admin_users(id),
    rollback_reason TEXT,
    
    -- Timestamps
    requested_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at TIMESTAMPTZ,
    closed_at TIMESTAMPTZ,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_change_approvals_number ON admin_change_approvals (request_number);
CREATE INDEX idx_change_approvals_requester ON admin_change_approvals (requested_by, requested_at DESC);
CREATE INDEX idx_change_approvals_status ON admin_change_approvals (status, risk_level, requested_at DESC);
CREATE INDEX idx_change_approvals_resource ON admin_change_approvals (resource_type, resource_id);
CREATE INDEX idx_change_approvals_expires ON admin_change_approvals (expires_at) WHERE status = 'PENDING';

COMMENT ON TABLE admin_change_approvals IS 'Change approval workflow - maps to internal/domain/change_approval/entity.go';

-- =========================================
-- SUB-ENTITY: change_approval_decisions
-- Domain: internal/domain/change_approval/
-- Entity: change_approval/approval.go
-- =========================================

CREATE TABLE admin_change_approval_decisions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Change Request Reference
    change_approval_id UUID NOT NULL REFERENCES admin_change_approvals(id) ON DELETE CASCADE,
    
    -- Approver
    approver_id UUID NOT NULL REFERENCES admin_users(id),
    approver_role VARCHAR(50) NOT NULL,
    
    -- Decision
    decision VARCHAR(20) NOT NULL CHECK (
        decision IN ('APPROVED', 'REJECTED', 'CONDITIONAL_APPROVAL')
    ),
    rationale TEXT NOT NULL,
    conditions TEXT, -- For conditional approvals
    
    -- Timestamp
    decided_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_change_approval_decisions_request ON admin_change_approval_decisions (change_approval_id, decided_at);
CREATE INDEX idx_change_approval_decisions_approver ON admin_change_approval_decisions (approver_id);

COMMENT ON TABLE admin_change_approval_decisions IS 'Approval decisions - maps to internal/domain/change_approval/approval.go';
```

---

## **SECTION 13: REPORTING & METRICS**

### Domain: reporting/

```sql
-- =========================================
-- MAIN TABLE: admin_reports
-- Domain: internal/domain/reporting/
-- Entity: reporting/entity.go
-- =========================================

CREATE TABLE admin_reports (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Report Identification
    report_number VARCHAR(20) NOT NULL UNIQUE, -- e.g., RPT-2025-001234
    
    -- Report Details
    report_name VARCHAR(255) NOT NULL,
    report_type VARCHAR(50) CHECK (
        report_type IN ('PERFORMANCE', 'FINANCIAL', 'COMPLIANCE', 'OPERATIONAL', 'SECURITY', 'CUSTOM')
    ) NOT NULL,
    
    -- Parameters
    date_range_start DATE NOT NULL,
    date_range_end DATE NOT NULL,
    filters JSONB,
    
    -- Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('GENERATING', 'COMPLETED', 'FAILED', 'EXPIRED')
    ) DEFAULT 'GENERATING',
    
    -- Output
    output_format VARCHAR(20) CHECK (
        output_format IN ('PDF', 'CSV', 'EXCEL', 'JSON')
    ),
    output_file_path VARCHAR(500),
    output_file_size BIGINT,
    
    -- Generation
    generation_started_at TIMESTAMPTZ,
    generation_completed_at TIMESTAMPTZ,
    generation_duration_ms INTEGER,
    
    -- Error Details
    error_message TEXT,
    error_code VARCHAR(50),
    
    -- Generated By
    generated_by UUID NOT NULL REFERENCES admin_users(id),
    
    -- Schedule Reference (if scheduled)
    schedule_id UUID,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at TIMESTAMPTZ, -- Reports auto-deleted after expiry
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_reports_number ON admin_reports (report_number);
CREATE INDEX idx_admin_reports_generated_by ON admin_reports (generated_by, created_at DESC);
CREATE INDEX idx_admin_reports_type ON admin_reports (report_type, status);
CREATE INDEX idx_admin_reports_status ON admin_reports (status, created_at DESC);
CREATE INDEX idx_admin_reports_expires ON admin_reports (expires_at) WHERE status = 'COMPLETED';

COMMENT ON TABLE admin_reports IS 'Generated reports - maps to internal/domain/reporting/entity.go';

-- =========================================
-- SUB-ENTITY: admin_report_schedules
-- Domain: internal/domain/reporting/
-- Entity: reporting/schedule.go
-- =========================================

CREATE TABLE admin_report_schedules (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Schedule Details
    schedule_name VARCHAR(255) NOT NULL,
    report_type VARCHAR(50) NOT NULL,
    
    -- Configuration
    report_config JSONB NOT NULL, -- Filters, parameters, etc.
    output_format VARCHAR(20) NOT NULL,
    
    -- Frequency
    frequency VARCHAR(20) CHECK (
        frequency IN ('DAILY', 'WEEKLY', 'MONTHLY', 'QUARTERLY', 'ANNUAL')
    ) NOT NULL,
    day_of_week INTEGER CHECK (day_of_week >= 0 AND day_of_week <= 6),
    day_of_month INTEGER CHECK (day_of_month >= 1 AND day_of_month <= 31),
    time_of_day TIME,
    
    -- Recipients
    recipients JSONB NOT NULL, -- Array of email addresses
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    
    -- Execution Tracking
    last_executed_at TIMESTAMPTZ,
    next_execution_at TIMESTAMPTZ,
    execution_count INTEGER DEFAULT 0,
    last_execution_status VARCHAR(20),
    
    -- Created By
    created_by UUID NOT NULL REFERENCES admin_users(id),
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_admin_report_schedules_active ON admin_report_schedules (is_active, next_execution_at);
CREATE INDEX idx_admin_report_schedules_type ON admin_report_schedules (report_type);

COMMENT ON TABLE admin_report_schedules IS 'Scheduled reports - maps to internal/domain/reporting/schedule.go';
```

### Domain: metrics/

```sql
-- =========================================
-- MAIN TABLE: admin_platform_metrics
-- Domain: internal/domain/metrics/
-- Entity: metrics/entity.go
-- =========================================

CREATE TABLE admin_platform_metrics (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Metric Details
    metric_type VARCHAR(50) NOT NULL CHECK (
        metric_type IN ('TICKET', 'MODERATION', 'USER_ACTION', 'DISPUTE', 'PLATFORM_HEALTH', 'PERFORMANCE', 'SECURITY')
    ),
    metric_name VARCHAR(100) NOT NULL,
    
    -- Time Period
    period_type VARCHAR(20) CHECK (
        period_type IN ('HOURLY', 'DAILY', 'WEEKLY', 'MONTHLY')
    ) NOT NULL,
    period_start TIMESTAMPTZ NOT NULL,
    period_end TIMESTAMPTZ NOT NULL,
    
    -- Metric Values
    metric_value DECIMAL(15,4),
    metric_unit VARCHAR(50),
    metric_dimensions JSONB, -- Additional breakdowns
    
    -- Aggregates
    count INTEGER,
    sum BIGINT,
    avg DECIMAL(15,4),
    min DECIMAL(15,4),
    max DECIMAL(15,4),
    p50 DECIMAL(15,4),
    p95 DECIMAL(15,4),
    p99 DECIMAL(15,4),
    
    -- Comparison
    previous_period_value DECIMAL(15,4),
    change_percentage DECIMAL(5,2),
    
    -- Timestamps
    calculated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    
    CONSTRAINT uk_admin_platform_metrics UNIQUE (metric_type, metric_name, period_type, period_start)
);

CREATE INDEX idx_admin_platform_metrics_type ON admin_platform_metrics (metric_type, period_start DESC);
CREATE INDEX idx_admin_platform_metrics_name ON admin_platform_metrics (metric_name, period_type, period_start DESC);
CREATE INDEX idx_admin_platform_metrics_period ON admin_platform_metrics (period_type, period_start DESC);

COMMENT ON TABLE admin_platform_metrics IS 'Platform metrics - maps to internal/domain/metrics/entity.go';
```

---

## **SECTION 14: OUTBOX PATTERN**

```sql
-- =========================================
-- MAIN TABLE: outbox_events
-- SHARED TABLE (platform-shared/outbox/)
-- =========================================

CREATE TABLE outbox_events (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Event Identification
    event_id UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    event_type VARCHAR(100) NOT NULL,
    event_version VARCHAR(10) DEFAULT '1',
    
    -- Aggregate
    aggregate_type VARCHAR(50) NOT NULL,
    aggregate_id UUID NOT NULL,
    
    -- Correlation & Causation
    correlation_id UUID NOT NULL,
    causation_id UUID,
    
    -- Event Source
    event_source VARCHAR(50) DEFAULT 'admin-be',
    
    -- User Context
    user_context JSONB,
    
    -- Compliance Context
    compliance_context JSONB,
    
    -- Audit Metadata
    audit_metadata JSONB,
    
    -- Payload
    payload JSONB NOT NULL,
    
    -- Processing Status
    status VARCHAR(20) NOT NULL CHECK (
        status IN ('PENDING', 'PROCESSING', 'PUBLISHED', 'FAILED')
    ) DEFAULT 'PENDING',
    
    -- Retry Tracking
    retry_count INTEGER DEFAULT 0,
    max_retries INTEGER DEFAULT 3,
    last_error TEXT,
    
    -- Idempotency
    idempotency_key VARCHAR(255),
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    processed_at TIMESTAMPTZ,
    published_at TIMESTAMPTZ,
    
    -- Metadata
    metadata JSONB DEFAULT '{}'::JSONB,
    
    CONSTRAINT uk_outbox_idempotency UNIQUE (aggregate_id, event_type, idempotency_key)
);

CREATE INDEX idx_outbox_events_status ON outbox_events (status, created_at) WHERE status IN ('PENDING', 'FAILED');
CREATE INDEX idx_outbox_events_aggregate ON outbox_events (aggregate_type, aggregate_id);
CREATE INDEX idx_outbox_events_event_type ON outbox_events (event_type);
CREATE INDEX idx_outbox_events_correlation ON outbox_events (correlation_id);
CREATE INDEX idx_outbox_events_created_at ON outbox_events (created_at DESC);

COMMENT ON TABLE outbox_events IS 'Outbox pattern for reliable event publishing';
COMMENT ON COLUMN outbox_events.status IS 'PENDING: awaiting publish, PROCESSING: currently publishing, PUBLISHED: successfully published, FAILED: publish failed';
```

---

## **FINAL SUMMARY**

### **Database Coverage**
- **Total Sections:** 14
- **Total Domain Tables:** 60+
- **Total Sub-Entity Tables:** 25+
- **Total Indexes:** 400+
- **Coverage:** 100% of admin-be folder structure domains

### **Alignment with Architecture**
✅ **CRITICAL RULES FOLLOWED:**
1. Each domain folder = ONE main table
2. Tables prefixed with `admin_` to reflect domain boundary
3. Sub-entities use `{domain}_{sub}` naming convention
4. All domains from folder structure covered
5. Production-ready fields with comprehensive indexes

### **Key Features**
- **Audit Trail:** Immutable admin activity logs with 10-year retention
- **RBAC:** Bitset-based roles and permissions for performance
- **PII Protection:** Encrypted sensitive fields, break-glass access tracking
- **Compliance:** GDPR/CCPA DSAR support, legal holds, sanctions screening
- **SLA Tracking:** SLA targets and breach detection across all case types
- **Event Sourcing:** Outbox pattern with idempotency for reliable event publishing
- **Versioning:** Config changes, KB articles, policy documents all versioned
- **Multi-tier Workflow:** Approval workflows for sensitive operations
- **Comprehensive Moderation:** Queue-based moderation with auto-flagging
- **Risk Management:** Fraud detection, holds, incident management
- **Break-Glass Access:** PII access requests with approval and audit trail

### **Production Ready**
- Partitionable tables for high-volume data (activity logs, events)
- GIN indexes for JSONB and array columns
- Full-text search indexes for searchable content
- Composite indexes for common query patterns
- Check constraints for data integrity
- Foreign key relationships for referential integrity
- BTREE and GIN indexes optimized for query patterns

### **Security & Compliance**
- PII fields encrypted at rest
- No PII in event payloads (IDs only)
- Comprehensive audit trails (immutable)
- SOX compliant change tracking
- GDPR/CCPA data subject request handling
- Sanctions screening and watchlist checks
- Legal hold support for litigation
- Break-glass access with approval workflow

### **Scalability**
- Designed for millions of tickets, cases, and actions
- Partitioning strategy for time-series data
- Index optimization for read-heavy operations
- JSONB for flexible metadata without schema changes
- Efficient bitset storage for roles/permissions

---

**END OF ADMIN-BE DATABASE DESIGN**
