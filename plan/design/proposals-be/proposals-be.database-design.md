=========================================
# PROPOSALS-BE DATABASE DESIGN
* Skillsier Platform - Enterprise Scale
* PostgreSQL 16+
=========================================

- CRITICAL ALIGNMENT RULES:
- 1. Each domain folder in internal/domain/{domain}/ = ONE main table
- 2. Table names match domain folder names exactly
- 3. Sub-entities within domain create related tables with {domain}_{sub} naming
- 4. All domains from folder structure are covered
- 5. Rich, production-ready fields for large-scale application

```sql
-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";
CREATE EXTENSION IF NOT EXISTS "btree_gin";
CREATE EXTENSION IF NOT EXISTS "citext";
CREATE EXTENSION IF NOT EXISTS "btree_gist";

```
=========================================
##  SECTION 1: CORE PROPOSAL LIFECYCLE
```sql
-- Domain: internal/domain/proposal/
-- Entity: proposal/entity.go
-- =========================================

CREATE TABLE proposals (
    -- Primary Key
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    -- Idempotency
    idempotency_key UUID UNIQUE,

    -- References to Other Services
    job_id UUID NOT NULL,
    freelancer_id UUID NOT NULL,

    -- Optional Invitation/Template Context
    invite_id UUID,
    template_id UUID,

    -- Proposal Identity
    title VARCHAR(200) NOT NULL,

    -- Pricing
    pricing_model VARCHAR(20) NOT NULL CHECK (
        pricing_model IN ('FIXED', 'HOURLY', 'MILESTONE', 'RETAINER', 'SUBSCRIPTION')
    ),
    proposed_amount DECIMAL(12, 2) NOT NULL,
    currency CHAR(3) DEFAULT 'USD',
    hourly_rate DECIMAL(10, 2), -- if pricing_model = HOURLY

    -- Duration
    estimated_duration_value INTEGER,
    estimated_duration_unit VARCHAR(20), -- HOURS, DAYS, WEEKS, MONTHS

    -- Availability
    available_start_date DATE,
    available_hours_per_week INTEGER,

    -- Status & Lifecycle
    status VARCHAR(20) DEFAULT 'DRAFT' CHECK (
        status IN ('DRAFT', 'SUBMITTED', 'UNDER_REVIEW', 'SHORTLISTED', 'INTERVIEW_SCHEDULED',
                   'ACCEPTED', 'REJECTED', 'WITHDRAWN', 'EXPIRED', 'ARCHIVED')
    ),
    submission_status VARCHAR(20) CHECK (
        submission_status IN ('PENDING', 'PROCESSING', 'COMPLETED', 'FAILED')
    ),

    -- Lifecycle Timestamps
    submitted_at TIMESTAMPTZ,
    reviewed_at TIMESTAMPTZ,
    shortlisted_at TIMESTAMPTZ,
    accepted_at TIMESTAMPTZ,
    rejected_at TIMESTAMPTZ,
    withdrawn_at TIMESTAMPTZ,
    expired_at TIMESTAMPTZ,
    archived_at TIMESTAMPTZ,

    -- Rejection/Withdrawal Reasons
    rejection_reason VARCHAR(100),
    rejection_details TEXT,
    withdrawal_reason VARCHAR(100),
    withdrawal_details TEXT,

    -- Client Interaction
    client_viewed_at TIMESTAMPTZ,
    client_view_count INTEGER DEFAULT 0,
    client_last_viewed_at TIMESTAMPTZ,

    -- Quality Metrics
    quality_score DECIMAL(5, 2),
    completeness_score INTEGER DEFAULT 0 CHECK (completeness_score BETWEEN 0 AND 100),
    match_score DECIMAL(5, 2), -- AI-computed match with job requirements

    -- Engagement Tracking
    views_count INTEGER DEFAULT 0,
    shares_count INTEGER DEFAULT 0,
    bookmarks_count INTEGER DEFAULT 0,

    -- Flags
    is_featured BOOLEAN DEFAULT FALSE,
    is_boosted BOOLEAN DEFAULT FALSE,
    is_priority BOOLEAN DEFAULT FALSE,
    is_ai_generated BOOLEAN DEFAULT FALSE,
    is_team_proposal BOOLEAN DEFAULT FALSE,

    -- Spam & Fraud Detection
    spam_score DECIMAL(5, 2) DEFAULT 0,
    is_flagged_spam BOOLEAN DEFAULT FALSE,
    fraud_score DECIMAL(5, 2) DEFAULT 0,
    is_flagged_fraud BOOLEAN DEFAULT FALSE,

    -- Soft Delete
    is_deleted BOOLEAN DEFAULT FALSE,
    deleted_at TIMESTAMPTZ,
    deleted_by UUID,

    -- Audit
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    version INTEGER DEFAULT 1 NOT NULL,

    -- Constraints
    CONSTRAINT uk_proposals_job_freelancer UNIQUE (job_id, freelancer_id),
    CONSTRAINT chk_proposals_amounts CHECK (proposed_amount >= 0),
    CONSTRAINT chk_proposals_hourly CHECK (hourly_rate IS NULL OR hourly_rate >= 0)
);

-- Indexes for performance
CREATE INDEX idx_proposals_job ON proposals (job_id, status) WHERE is_deleted = FALSE;
CREATE INDEX idx_proposals_freelancer ON proposals (freelancer_id, status) WHERE is_deleted = FALSE;
CREATE INDEX idx_proposals_status ON proposals (status, submitted_at DESC) WHERE is_deleted = FALSE;
CREATE INDEX idx_proposals_submitted ON proposals (submitted_at DESC) WHERE status != 'DRAFT';
CREATE INDEX idx_proposals_quality ON proposals (quality_score DESC) WHERE status = 'SUBMITTED';
CREATE INDEX idx_proposals_match ON proposals (match_score DESC) WHERE status IN ('SUBMITTED', 'UNDER_REVIEW');
CREATE INDEX idx_proposals_spam ON proposals (spam_score DESC) WHERE is_flagged_spam = TRUE;
CREATE INDEX idx_proposals_idempotency ON proposals (idempotency_key) WHERE idempotency_key IS NOT NULL;

```
=========================================
##  SECTION 2: COVER LETTER
```sql
-- Domain: internal/domain/cover_letter/
-- Entity: cover_letter/entity.go
-- =========================================

CREATE TABLE cover_letters (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL UNIQUE,

    -- Content
    content TEXT NOT NULL,

    -- Metadata
    word_count INTEGER DEFAULT 0,
    char_count INTEGER DEFAULT 0,
    tone VARCHAR(20), -- PROFESSIONAL, CASUAL, FRIENDLY, FORMAL

    -- AI Analysis
    ai_sentiment_score DECIMAL(5, 2), -- -1.0 to 1.0
    ai_clarity_score DECIMAL(5, 2), -- 0 to 100
    ai_suggestions JSONB, -- Array of improvement suggestions

    -- Personalization
    personalization_score INTEGER DEFAULT 0 CHECK (personalization_score BETWEEN 0 AND 100),
    has_custom_greeting BOOLEAN DEFAULT FALSE,
    has_specific_references BOOLEAN DEFAULT FALSE,

    -- Quality Flags
    has_grammar_issues BOOLEAN DEFAULT FALSE,
    has_spelling_issues BOOLEAN DEFAULT FALSE,
    readability_score DECIMAL(5, 2), -- Flesch reading ease

    -- Versions (for drafts)
    version INTEGER DEFAULT 1,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_cover_letters_proposal FOREIGN KEY (proposal_id) REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_cover_letters_proposal ON cover_letters (proposal_id);

COMMENT ON TABLE cover_letters IS 'Cover letters - maps to internal/domain/cover_letter/entity.go';

```
=========================================
##  SECTION 3: ATTACHMENTS
```sql
-- Domain: internal/domain/attachment/
-- Entity: attachment/entity.go
-- =========================================

CREATE TABLE attachments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- File Reference (storage-be)
    file_id UUID NOT NULL,
    file_url TEXT NOT NULL,

    -- File Metadata
    file_name VARCHAR(255) NOT NULL,
    file_size BIGINT NOT NULL, -- bytes
    file_type VARCHAR(100) NOT NULL,
    mime_type VARCHAR(100),

    -- Security & Scanning
    virus_scan_status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        virus_scan_status IN ('PENDING', 'CLEAN', 'INFECTED', 'FAILED')
    ),
    virus_scan_result TEXT,
    scanned_at TIMESTAMPTZ,

    -- Content Analysis
    has_sensitive_data BOOLEAN DEFAULT FALSE,
    content_type VARCHAR(50), -- PORTFOLIO, RESUME, CERTIFICATE, SAMPLE_WORK

    -- Display Order
    display_order INTEGER DEFAULT 0,

    -- Status
    is_visible BOOLEAN DEFAULT TRUE,
    is_deleted BOOLEAN DEFAULT FALSE,
    deleted_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_attachments_proposal FOREIGN KEY (proposal_id) REFERENCES proposals(id) ON DELETE CASCADE,
    CONSTRAINT uk_attachments_file UNIQUE (proposal_id, file_id)
);

CREATE INDEX idx_attachments_proposal ON attachments (proposal_id, display_order);
CREATE INDEX idx_attachments_file ON attachments (file_id);
CREATE INDEX idx_attachments_scan ON attachments (virus_scan_status) WHERE virus_scan_status != 'CLEAN';

COMMENT ON TABLE attachments IS 'Proposal attachments - maps to internal/domain/attachment/entity.go';

```
=========================================
##  SECTION 4: QUESTION ANSWERS
```sql
-- Domain: internal/domain/question_answer/
-- Entity: question_answer/entity.go
-- =========================================

CREATE TABLE question_answers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- Question Reference (from jobs-be screening_questions)
    question_id UUID NOT NULL,
    question_text TEXT NOT NULL, -- Cached for historical accuracy
    question_type VARCHAR(20), -- TEXT, CHOICE, BOOLEAN, NUMERIC

    -- Answer
    answer_text TEXT,
    answer_choice VARCHAR(255), -- for CHOICE type
    answer_boolean BOOLEAN, -- for BOOLEAN type
    answer_numeric DECIMAL(10, 2), -- for NUMERIC type

    -- Metadata
    is_required BOOLEAN DEFAULT FALSE,
    is_answered BOOLEAN DEFAULT FALSE,
    answer_length INTEGER,

    -- Quality Analysis
    quality_score DECIMAL(5, 2),
    is_complete BOOLEAN DEFAULT FALSE,
    is_relevant BOOLEAN DEFAULT TRUE,

    -- Timestamps
    occurred_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    published_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_outbox_events_status ON outbox_events (status, next_attempt_at) WHERE status IN ('PENDING', 'FAILED');
CREATE INDEX idx_outbox_events_aggregate ON outbox_events (aggregate_type, aggregate_id);
CREATE INDEX idx_outbox_events_correlation ON outbox_events (correlation_id);
CREATE INDEX idx_outbox_events_topic ON outbox_events (topic, occurred_at DESC);

COMMENT ON TABLE outbox_events IS 'Transactional outbox for event publishing';

```
=========================================
##  SECTION 5: MILESTONES
```sql
-- Domain: internal/domain/milestone/
-- Entity: milestone/entity.go
-- =========================================

CREATE TABLE milestones (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- Milestone Details
    title VARCHAR(200) NOT NULL,
    description TEXT,

    -- Financials
    amount DECIMAL(12, 2) NOT NULL,
    currency CHAR(3) DEFAULT 'USD',

    -- Timeline
    estimated_duration_value INTEGER,
    estimated_duration_unit VARCHAR(20),
    due_date DATE,

    -- Order
    sequence_number INTEGER NOT NULL,

    -- Deliverables
    deliverables TEXT[],
    acceptance_criteria TEXT,

    -- Status
    status VARCHAR(20) DEFAULT 'PLANNED' CHECK (
        status IN ('PLANNED', 'ACTIVE', 'COMPLETED', 'CANCELLED')
    ),

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_milestones_proposal FOREIGN KEY (proposal_id) REFERENCES proposals(id) ON DELETE CASCADE,
    CONSTRAINT uk_milestones_sequence UNIQUE (proposal_id, sequence_number),
    CONSTRAINT chk_milestones_amount CHECK (amount >= 0)
);

CREATE INDEX idx_milestones_proposal ON milestones (proposal_id, sequence_number);

COMMENT ON TABLE milestones IS 'Proposal milestones - maps to internal/domain/milestone/entity.go';

```
=========================================
##  SECTION 6: BIDDING SYSTEM
```sql
-- Domain: internal/domain/bid/
-- Entity: bid/entity.go
-- =========================================

CREATE TABLE bids (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- Bid Details
    bid_amount DECIMAL(12, 2) NOT NULL,
    currency CHAR(3) DEFAULT 'USD',

    -- Bidding Context
    bid_type VARCHAR(20) DEFAULT 'STANDARD' CHECK (
        bid_type IN ('STANDARD', 'QUICK_BID', 'AUTO_BID', 'COUNTER_BID')
    ),

    -- Auto-Bidding Configuration
    is_auto_bid BOOLEAN DEFAULT FALSE,
    auto_bid_max_amount DECIMAL(12, 2),
    auto_bid_increment DECIMAL(12, 2),

    -- Auction Status
    is_winning_bid BOOLEAN DEFAULT FALSE,
    is_outbid BOOLEAN DEFAULT FALSE,
    outbid_at TIMESTAMPTZ,

    -- Ranking
    rank INTEGER,

    -- Validity
    valid_until TIMESTAMPTZ,
    is_expired BOOLEAN DEFAULT FALSE,

    -- Status
    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (
        status IN ('ACTIVE', 'WITHDRAWN', 'ACCEPTED', 'REJECTED', 'EXPIRED')
    ),

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_bids_proposal FOREIGN KEY (proposal_id) REFERENCES proposals(id) ON DELETE CASCADE,
    CONSTRAINT chk_bids_amount CHECK (bid_amount > 0)
);

CREATE INDEX idx_bids_proposal ON bids (proposal_id, bid_amount ASC);
CREATE INDEX idx_bids_winning ON bids (is_winning_bid, created_at DESC) WHERE is_winning_bid = TRUE;
CREATE INDEX idx_bids_status ON bids (status, created_at DESC);

COMMENT ON TABLE bids IS 'Proposal bids - maps to internal/domain/bid/entity.go';

```
=========================================
##  SECTION 7: BID STRATEGY
```sql
-- Domain: internal/domain/bid_strategy/
-- Entity: bid_strategy/entity.go
-- =========================================

CREATE TABLE bid_strategies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    freelancer_id UUID NOT NULL,

    -- Strategy Configuration
    strategy_name VARCHAR(100) NOT NULL,
    strategy_type VARCHAR(20) DEFAULT 'CONSERVATIVE' CHECK (
        strategy_type IN ('CONSERVATIVE', 'MODERATE', 'AGGRESSIVE', 'CUSTOM')
    ),

    -- Rules
    min_bid_amount DECIMAL(12, 2),
    max_bid_amount DECIMAL(12, 2),
    increment_percentage DECIMAL(5, 2),
    increment_amount DECIMAL(12, 2),

    -- Conditions
    conditions JSONB, -- Flexible rule conditions

    -- Time-based Rules
    time_of_day_rules JSONB,
    day_of_week_rules JSONB,

    -- Budget Allocation
    daily_budget DECIMAL(12, 2),
    monthly_budget DECIMAL(12, 2),
    spent_today DECIMAL(12, 2) DEFAULT 0,
    spent_this_month DECIMAL(12, 2) DEFAULT 0,

    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    is_paused BOOLEAN DEFAULT FALSE,
    paused_at TIMESTAMPTZ,
    paused_reason TEXT,

    -- Performance Tracking
    bids_placed INTEGER DEFAULT 0,
    bids_won INTEGER DEFAULT 0,
    total_spent DECIMAL(12, 2) DEFAULT 0,
    win_rate DECIMAL(5, 2),

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT chk_bid_strategies_amounts CHECK (max_bid_amount IS NULL OR min_bid_amount IS NULL OR max_bid_amount >= min_bid_amount)
);

CREATE INDEX idx_bid_strategies_freelancer ON bid_strategies (freelancer_id);
CREATE INDEX idx_bid_strategies_active ON bid_strategies (is_active) WHERE is_active = TRUE;

COMMENT ON TABLE bid_strategies IS 'Bid strategies - maps to internal/domain/bid_strategy/entity.go';

```
=========================================
##  SECTION 8: BID NOTIFICATIONS
```sql
-- Domain: internal/domain/bid_notification/
-- Entity: bid_notification/entity.go
-- =========================================

CREATE TABLE bid_notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    bid_id UUID NOT NULL,
    freelancer_id UUID NOT NULL,

    -- Notification Details
    notification_type VARCHAR(30) NOT NULL CHECK (
        notification_type IN ('OUTBID', 'WINNING', 'BID_ACCEPTED', 'BID_REJECTED', 'AUCTION_ENDING')
    ),

    -- Content
    title VARCHAR(200) NOT NULL,
    message TEXT,

    -- Priority
    priority VARCHAR(20) DEFAULT 'MEDIUM' CHECK (
        priority IN ('LOW', 'MEDIUM', 'HIGH', 'URGENT')
    ),

    -- Delivery Status
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMPTZ,
    is_sent BOOLEAN DEFAULT FALSE,
    sent_at TIMESTAMPTZ,

    -- Channels
    sent_email BOOLEAN DEFAULT FALSE,
    sent_push BOOLEAN DEFAULT FALSE,
    sent_sms BOOLEAN DEFAULT FALSE,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_bid_notifications_bid FOREIGN KEY (bid_id) REFERENCES bids(id) ON DELETE CASCADE
);

CREATE INDEX idx_bid_notifications_bid ON bid_notifications (bid_id);
CREATE INDEX idx_bid_notifications_freelancer ON bid_notifications (freelancer_id, is_read);
CREATE INDEX idx_bid_notifications_type ON bid_notifications (notification_type, created_at DESC);

COMMENT ON TABLE bid_notifications IS 'Bid notifications - maps to internal/domain/bid_notification/entity.go';

```
=========================================
##  SECTION 9: AUCTIONS
```sql
-- Domain: internal/domain/auction/
-- Entity: auction/entity.go
-- =========================================

CREATE TABLE auctions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    job_id UUID NOT NULL UNIQUE,

    -- Auction Configuration
    auction_type VARCHAR(20) DEFAULT 'OPEN' CHECK (
        auction_type IN ('OPEN', 'SEALED', 'DUTCH', 'REVERSE')
    ),

    -- Pricing
    starting_price DECIMAL(12, 2),
    reserve_price DECIMAL(12, 2),
    current_price DECIMAL(12, 2),
    currency CHAR(3) DEFAULT 'USD',

    -- Timing
    starts_at TIMESTAMPTZ NOT NULL,
    ends_at TIMESTAMPTZ NOT NULL,
    extended_until TIMESTAMPTZ,
    auto_extend_enabled BOOLEAN DEFAULT TRUE,
    extend_by_minutes INTEGER DEFAULT 10,

    -- Rules
    min_bid_increment DECIMAL(12, 2),
    max_bid_amount DECIMAL(12, 2),
    bid_visibility VARCHAR(20) DEFAULT 'PUBLIC' CHECK (
        bid_visibility IN ('PUBLIC', 'HIDDEN', 'RANK_ONLY')
    ),

    -- Status
    status VARCHAR(20) DEFAULT 'SCHEDULED' CHECK (
        status IN ('SCHEDULED', 'ACTIVE', 'ENDED', 'CANCELLED')
    ),

    -- Participation Stats
    bids_count INTEGER DEFAULT 0,
    unique_bidders INTEGER DEFAULT 0,

    -- Winning Bid
    winning_bid_id UUID,
    winning_freelancer_id UUID,
    winning_amount DECIMAL(12, 2),

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT chk_auctions_timing CHECK (ends_at > starts_at),
    CONSTRAINT chk_auctions_prices CHECK (reserve_price IS NULL OR starting_price IS NULL OR reserve_price >= starting_price)
);

CREATE INDEX idx_auctions_job ON auctions (job_id);
CREATE INDEX idx_auctions_status ON auctions (status, ends_at);
CREATE INDEX idx_auctions_active ON auctions (status, ends_at) WHERE status = 'ACTIVE';

COMMENT ON TABLE auctions IS 'Job auctions - maps to internal/domain/auction/entity.go';

```
=========================================
##  SECTION 10: BID ANOMALY DETECTION
```sql
-- Domain: internal/domain/bid_anomaly_detection/
-- Entity: bid_anomaly_detection/entity.go
-- =========================================

CREATE TABLE bid_anomalies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    bid_id UUID NOT NULL,
    proposal_id UUID NOT NULL,
    freelancer_id UUID NOT NULL,

    -- Anomaly Details
    anomaly_type VARCHAR(50) NOT NULL CHECK (
        anomaly_type IN ('UNUSUALLY_LOW', 'UNUSUALLY_HIGH', 'RAPID_BIDDING', 'PATTERN_VIOLATION',
                         'SUSPICIOUS_TIMING', 'COLLUSION_INDICATOR', 'FRAUD_INDICATOR')
    ),

    -- Severity
    severity VARCHAR(20) DEFAULT 'MEDIUM' CHECK (
        severity IN ('LOW', 'MEDIUM', 'HIGH', 'CRITICAL')
    ),
    confidence_score DECIMAL(5, 2), -- 0-100

    -- Analysis
    description TEXT,
    detected_patterns JSONB,
    risk_factors JSONB,

    -- Actions
    action_taken VARCHAR(50) CHECK (
        action_taken IN ('NONE', 'FLAGGED', 'BLOCKED', 'MANUAL_REVIEW', 'AUTO_REJECTED')
    ),
    action_reason TEXT,
    reviewed_by UUID,
    reviewed_at TIMESTAMPTZ,

    -- Resolution
    is_false_positive BOOLEAN,
    resolution_notes TEXT,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_bid_anomalies_bid FOREIGN KEY (bid_id) REFERENCES bids(id) ON DELETE CASCADE,
    CONSTRAINT fk_bid_anomalies_proposal FOREIGN KEY (proposal_id) REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_bid_anomalies_bid ON bid_anomalies (bid_id);
CREATE INDEX idx_bid_anomalies_proposal ON bid_anomalies (proposal_id);
CREATE INDEX idx_bid_anomalies_freelancer ON bid_anomalies (freelancer_id);
CREATE INDEX idx_bid_anomalies_severity ON bid_anomalies (severity, created_at DESC) WHERE severity IN ('HIGH', 'CRITICAL');

COMMENT ON TABLE bid_anomalies IS 'Bid anomaly detection - maps to internal/domain/bid_anomaly_detection/entity.go';

```
=========================================
##  SECTION 11: CONNECTS & BOOST SYSTEM
```sql
-- Domain: internal/domain/connect/
-- Entity: connect/entity.go
-- =========================================

CREATE TABLE connects (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    freelancer_id UUID NOT NULL,

    -- Balance
    total_connects INTEGER DEFAULT 0,
    available_connects INTEGER DEFAULT 0,
    reserved_connects INTEGER DEFAULT 0,

    -- Transactions Summary
    earned_total INTEGER DEFAULT 0,
    purchased_total INTEGER DEFAULT 0,
    spent_total INTEGER DEFAULT 0,
    refunded_total INTEGER DEFAULT 0,

    -- Subscription Benefits
    monthly_grant INTEGER DEFAULT 0,
    last_monthly_grant_at TIMESTAMPTZ,
    next_grant_at TIMESTAMPTZ,

    -- Expiration
    expiring_soon_count INTEGER DEFAULT 0,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT uk_connects_freelancer UNIQUE (freelancer_id),
    CONSTRAINT chk_connects_balance CHECK (available_connects >= 0 AND reserved_connects >= 0)
);

CREATE INDEX idx_connects_freelancer ON connects (freelancer_id);

COMMENT ON TABLE connects IS 'Freelancer connects balance - maps to internal/domain/connect/entity.go';

-- Connect Transactions Ledger
CREATE TABLE connect_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    freelancer_id UUID NOT NULL,

    -- Transaction Details
    transaction_type VARCHAR(30) NOT NULL CHECK (
        transaction_type IN ('GRANT', 'PURCHASE', 'SPEND', 'REFUND', 'EXPIRATION', 'ADMIN_ADJUSTMENT')
    ),
    amount INTEGER NOT NULL,

    -- Reference
    reference_type VARCHAR(30), -- PROPOSAL, SUBSCRIPTION, PURCHASE, etc
    reference_id UUID,

    -- Balance Snapshot
    balance_before INTEGER,
    balance_after INTEGER,

    -- Description
    description TEXT,

    -- Metadata
    metadata JSONB,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_connect_transactions_freelancer FOREIGN KEY (freelancer_id)
        REFERENCES connects(freelancer_id) ON DELETE CASCADE
);

CREATE INDEX idx_connect_transactions_freelancer ON connect_transactions (freelancer_id, created_at DESC);
CREATE INDEX idx_connect_transactions_type ON connect_transactions (transaction_type);
CREATE INDEX idx_connect_transactions_reference ON connect_transactions (reference_type, reference_id);

```
=========================================
##  SECTION 12: CONNECT REFUNDS
```sql
-- Domain: internal/domain/connect_refund/
-- Entity: connect_refund/entity.go
-- =========================================

CREATE TABLE connect_refunds (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,
    freelancer_id UUID NOT NULL,
    transaction_id UUID NOT NULL, -- original spend transaction

    -- Refund Details
    refund_amount INTEGER NOT NULL,
    refund_reason VARCHAR(50) NOT NULL CHECK (
        refund_reason IN ('JOB_CANCELLED', 'JOB_CLOSED_EARLY', 'INVALID_JOB', 'SYSTEM_ERROR', 'ADMIN_DECISION')
    ),

    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'APPROVED', 'REJECTED', 'PROCESSED')
    ),

    -- Processing
    processed_by UUID,
    processed_at TIMESTAMPTZ,
    processing_notes TEXT,

    -- Rejection
    rejection_reason TEXT,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_connect_refunds_proposal FOREIGN KEY (proposal_id) REFERENCES proposals(id) ON DELETE CASCADE,
    CONSTRAINT uk_connect_refunds_transaction UNIQUE (transaction_id)
);

CREATE INDEX idx_connect_refunds_proposal ON connect_refunds (proposal_id);
CREATE INDEX idx_connect_refunds_freelancer ON connect_refunds (freelancer_id);
CREATE INDEX idx_connect_refunds_status ON connect_refunds (status);

COMMENT ON TABLE connect_refunds IS 'Connect refunds - maps to internal/domain/connect_refund/entity.go';

```
=========================================
##  SECTION 13: BOOST SYSTEM
```sql
-- Domain: internal/domain/boost/
-- Entity: boost/entity.go
-- =========================================

CREATE TABLE boosts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- Boost Configuration
    boost_type VARCHAR(20) DEFAULT 'STANDARD' CHECK (
        boost_type IN ('STANDARD', 'PREMIUM', 'FEATURED')
    ),

    -- Cost
    boost_cost INTEGER NOT NULL, -- in connects or currency
    cost_type VARCHAR(20) DEFAULT 'CONNECTS' CHECK (
        cost_type IN ('CONNECTS', 'CURRENCY')
    ),
    currency CHAR(3),

    -- Duration & Validity
    duration_hours INTEGER NOT NULL,
    starts_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMPTZ NOT NULL,

    -- Status
    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (
        status IN ('ACTIVE', 'EXPIRED', 'CANCELLED')
    ),

    -- Performance Impact
    visibility_increase_percentage INTEGER,
    impressions_gained INTEGER DEFAULT 0,
    clicks_gained INTEGER DEFAULT 0,

    -- Placement
    placement_priority INTEGER DEFAULT 0,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_boosts_proposal FOREIGN KEY (proposal_id) REFERENCES proposals(id) ON DELETE CASCADE,
    CONSTRAINT chk_boosts_duration CHECK (duration_hours > 0),
    CONSTRAINT chk_boosts_timing CHECK (expires_at > starts_at)
);

CREATE INDEX idx_boosts_proposal ON boosts (proposal_id);
CREATE INDEX idx_boosts_active ON boosts (status, expires_at) WHERE status = 'ACTIVE';
CREATE INDEX idx_boosts_priority ON boosts (placement_priority DESC, starts_at DESC);

COMMENT ON TABLE boosts IS 'Proposal boosts - maps to internal/domain/boost/entity.go';

```
=========================================
##  SECTION 14: TEMPLATES & RATE CARDS
```sql
-- Domain: internal/domain/template/
-- Entity: template/entity.go
-- =========================================

CREATE TABLE templates (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    freelancer_id UUID NOT NULL,

    -- Template Identity
    template_name VARCHAR(200) NOT NULL,
    template_slug VARCHAR(250),

    -- Content
    title_template VARCHAR(200),
    cover_letter_template TEXT,

    -- Configuration
    default_pricing_model VARCHAR(20),
    default_amount DECIMAL(12, 2),
    default_currency CHAR(3) DEFAULT 'USD',
    default_duration_value INTEGER,
    default_duration_unit VARCHAR(20),

    -- Categories & Tags
    job_categories UUID[], -- Array of category IDs this template is for
    tags TEXT[],

    -- Usage Statistics
    usage_count INTEGER DEFAULT 0,
    success_rate DECIMAL(5, 2),
    last_used_at TIMESTAMPTZ,

    -- Settings
    is_default BOOLEAN DEFAULT FALSE,
    is_favorite BOOLEAN DEFAULT FALSE,

    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    is_archived BOOLEAN DEFAULT FALSE,
    archived_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT uk_templates_freelancer_slug UNIQUE (freelancer_id, template_slug)
);

CREATE INDEX idx_templates_freelancer ON templates (freelancer_id, is_archived);
CREATE INDEX idx_templates_active ON templates (freelancer_id) WHERE is_active = TRUE;

COMMENT ON TABLE templates IS 'Proposal templates - maps to internal/domain/template/entity.go';

```
=========================================
##  SECTION 15: RATE CARDS
```sql
-- Domain: internal/domain/rate_card/
-- Entity: rate_card/entity.go
-- =========================================

CREATE TABLE rate_cards (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    freelancer_id UUID NOT NULL,

    -- Rate Card Details
    name VARCHAR(100) NOT NULL,
    description TEXT,

    -- Pricing Tiers
    tier_name VARCHAR(50),
    base_rate DECIMAL(10, 2) NOT NULL,
    currency CHAR(3) DEFAULT 'USD',
    rate_type VARCHAR(20) DEFAULT 'HOURLY' CHECK (
        rate_type IN ('HOURLY', 'DAILY', 'WEEKLY', 'MONTHLY', 'FIXED')
    ),

    -- Service Scope
    included_services TEXT[],
    excluded_services TEXT[],
    deliverables TEXT[],

    -- Time Commitments
    min_hours INTEGER,
    max_hours INTEGER,
    response_time_hours INTEGER,
    turnaround_days INTEGER,

    -- Discounts & Premiums
    volume_discount_percentage DECIMAL(5, 2),
    rush_premium_percentage DECIMAL(5, 2),
    weekend_premium_percentage DECIMAL(5, 2),

    -- Validity
    valid_from DATE,
    valid_until DATE,

    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    is_public BOOLEAN DEFAULT FALSE,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT chk_rate_cards_hours CHECK (max_hours IS NULL OR min_hours IS NULL OR max_hours >= min_hours)
);

CREATE INDEX idx_rate_cards_freelancer ON rate_cards (freelancer_id, is_active);
CREATE INDEX idx_rate_cards_validity ON rate_cards (valid_from, valid_until) WHERE is_active = TRUE;

COMMENT ON TABLE rate_cards IS 'Rate cards - maps to internal/domain/rate_card/entity.go';

```
=========================================
##  SECTION 16: PROPOSAL PERFORMANCE ANALYTICS
```sql
-- Domain: internal/domain/performance/ (consolidated)
-- Entity: performance/entity.go
-- =========================================

CREATE TABLE proposal_performance (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL UNIQUE,

    -- View Metrics
    total_views INTEGER DEFAULT 0,
    unique_views INTEGER DEFAULT 0,
    client_views INTEGER DEFAULT 0,
    avg_view_duration_seconds INTEGER,

    -- Engagement Metrics
    clicks INTEGER DEFAULT 0,
    shares INTEGER DEFAULT 0,
    bookmarks INTEGER DEFAULT 0,
    downloads INTEGER DEFAULT 0,

    -- Conversion Metrics
    shortlist_rate DECIMAL(5, 2),
    interview_rate DECIMAL(5, 2),
    acceptance_rate DECIMAL(5, 2),

    -- Timing Metrics
    time_to_first_view_hours INTEGER,
    time_to_shortlist_hours INTEGER,
    time_to_response_hours INTEGER,

    -- Ranking & Visibility
    avg_search_position DECIMAL(5, 2),
    visibility_score INTEGER DEFAULT 0,

    -- Comparison Metrics
    percentile_rank INTEGER,
    better_than_percent DECIMAL(5, 2),

    -- Quality Indicators
    quality_score DECIMAL(5, 2),
    completeness_score INTEGER,
    professionalism_score DECIMAL(5, 2),

    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_performance_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_performance_proposal ON proposal_performance (proposal_id);
CREATE INDEX idx_proposal_performance_quality ON proposal_performance (quality_score DESC);

COMMENT ON TABLE proposal_performance IS 'Proposal performance analytics - maps to internal/domain/performance/entity.go';

```
=========================================
##  SECTION 17: SIMILARITY & DEDUPLICATION
```sql
-- Domain: internal/domain/similarity/ (consolidated)
-- Entity: similarity/entity.go
-- =========================================

CREATE TABLE proposal_similarity (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- Similarity Hash
    content_hash VARCHAR(64) NOT NULL, -- SHA-256 of normalized content
    simhash BIGINT, -- Locality-sensitive hash for near-duplicate detection

    -- Feature Vectors
    title_vector TEXT, -- Embedding vector for title
    cover_letter_vector TEXT, -- Embedding vector for cover letter

    -- Duplicate Detection
    is_potential_duplicate BOOLEAN DEFAULT FALSE,
    similar_proposals UUID[], -- Array of similar proposal IDs
    similarity_scores JSONB, -- Map of proposal_id -> similarity_score

    -- Template Detection
    is_template_based BOOLEAN DEFAULT FALSE,
    template_id UUID,
    template_deviation_score DECIMAL(5, 2),

    -- Content Analysis
    unique_content_percentage DECIMAL(5, 2),
    personalization_score DECIMAL(5, 2),

    computed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_similarity_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_similarity_proposal ON proposal_similarity (proposal_id);
CREATE INDEX idx_proposal_similarity_hash ON proposal_similarity (content_hash);
CREATE INDEX idx_proposal_similarity_simhash ON proposal_similarity (simhash);
CREATE INDEX idx_proposal_similarity_duplicates ON proposal_similarity (is_potential_duplicate)
    WHERE is_potential_duplicate = TRUE;

COMMENT ON TABLE proposal_similarity IS 'Proposal similarity analysis - maps to internal/domain/similarity/entity.go';

-- Duplicate Clusters
CREATE TABLE duplicate_clusters (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    -- Cluster Details
    cluster_hash VARCHAR(64) UNIQUE,
    proposal_ids UUID[] NOT NULL,

    -- Statistics
    cluster_size INTEGER DEFAULT 0,
    avg_similarity_score DECIMAL(5, 2),

    -- Resolution
    representative_proposal_id UUID, -- The "original" or best proposal in cluster
    resolution_status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        resolution_status IN ('PENDING', 'REVIEWED', 'LEGITIMATE', 'SPAM', 'IGNORED')
    ),

    reviewed_by UUID,
    reviewed_at TIMESTAMPTZ,
    review_notes TEXT,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_duplicate_clusters_hash ON duplicate_clusters (cluster_hash);
CREATE INDEX idx_duplicate_clusters_size ON duplicate_clusters (cluster_size DESC);
CREATE INDEX idx_duplicate_clusters_status ON duplicate_clusters (resolution_status);

```
=========================================
##  SECTION 18: PORTFOLIO INTEGRATION
```sql
-- Domain: internal/domain/portfolio/ (consolidated)
-- Entity: portfolio/entity.go
-- =========================================

CREATE TABLE proposal_portfolios (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- Portfolio Item Reference (from users-be)
    portfolio_item_id UUID NOT NULL,

    -- Display
    display_order INTEGER DEFAULT 0,
    is_featured BOOLEAN DEFAULT FALSE,

    -- Context
    relevance_score DECIMAL(5, 2),
    relevance_notes TEXT,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_portfolios_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE,
    CONSTRAINT uk_proposal_portfolios UNIQUE (proposal_id, portfolio_item_id)
);

CREATE INDEX idx_proposal_portfolios_proposal ON proposal_portfolios (proposal_id, display_order);
CREATE INDEX idx_proposal_portfolios_item ON proposal_portfolios (portfolio_item_id);

COMMENT ON TABLE proposal_portfolios IS 'Portfolio links - maps to internal/domain/portfolio/entity.go';

```
=========================================
##  SECTION 19: ENGAGEMENT & FOLLOW-UP
```sql
-- Domain: internal/domain/engagement/ (consolidated)
-- Entity: engagement/entity.go
-- =========================================

CREATE TABLE proposal_engagement (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- Engagement Events
    last_client_view_at TIMESTAMPTZ,
    last_client_action_at TIMESTAMPTZ,
    last_freelancer_update_at TIMESTAMPTZ,

    -- Follow-up Tracking
    follow_up_count INTEGER DEFAULT 0,
    last_follow_up_at TIMESTAMPTZ,
    next_follow_up_at TIMESTAMPTZ,

    -- Response Tracking
    pending_client_response BOOLEAN DEFAULT FALSE,
    client_response_deadline TIMESTAMPTZ,

    -- Interaction History
    messages_count INTEGER DEFAULT 0,
    questions_asked INTEGER DEFAULT 0,
    questions_answered INTEGER DEFAULT 0,

    -- Engagement Score
    engagement_score DECIMAL(5, 2),
    engagement_level VARCHAR(20) CHECK (
        engagement_level IN ('NONE', 'LOW', 'MEDIUM', 'HIGH', 'VERY_HIGH')
    ),

    -- Status
    is_stale BOOLEAN DEFAULT FALSE,
    days_without_activity INTEGER DEFAULT 0,

    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_engagement_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_engagement_proposal ON proposal_engagement (proposal_id);
CREATE INDEX idx_proposal_engagement_followup ON proposal_engagement (next_follow_up_at)
    WHERE next_follow_up_at IS NOT NULL;
CREATE INDEX idx_proposal_engagement_stale ON proposal_engagement (is_stale, days_without_activity DESC)
    WHERE is_stale = TRUE;

COMMENT ON TABLE proposal_engagement IS 'Engagement tracking - maps to internal/domain/engagement/entity.go';

```
=========================================
##  SECTION 20: SPAM DETECTION & MODERATION
```sql
-- Domain: internal/domain/spam_detection/
-- Entity: spam_detection/entity.go
-- =========================================

CREATE TABLE spam_detections (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- Detection Details
    spam_indicators JSONB, -- Array of detected spam signals
    spam_score DECIMAL(5, 2) NOT NULL,
    confidence_level DECIMAL(5, 2),

    -- Classification
    classification VARCHAR(20) CHECK (
        classification IN ('LEGITIMATE', 'SUSPICIOUS', 'LIKELY_SPAM', 'CONFIRMED_SPAM')
    ),

    -- Pattern Matching
    matched_patterns TEXT[],
    keyword_flags TEXT[],

    -- Behavioral Signals
    submission_velocity_flag BOOLEAN DEFAULT FALSE,
    duplicate_content_flag BOOLEAN DEFAULT FALSE,
    low_effort_flag BOOLEAN DEFAULT FALSE,

    -- Action Taken
    action_taken VARCHAR(50) CHECK (
        action_taken IN ('NONE', 'FLAGGED', 'QUARANTINED', 'AUTO_REJECTED', 'MANUAL_REVIEW')
    ),
    action_at TIMESTAMPTZ,

    -- Review
    reviewed_by UUID,
    reviewed_at TIMESTAMPTZ,
    is_false_positive BOOLEAN,
    review_notes TEXT,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_spam_detections_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_spam_detections_proposal ON spam_detections (proposal_id);
CREATE INDEX idx_spam_detections_score ON spam_detections (spam_score DESC);
CREATE INDEX idx_spam_detections_classification ON spam_detections (classification);

COMMENT ON TABLE spam_detections IS 'Spam detection - maps to internal/domain/spam_detection/entity.go';

```
=========================================
##  SECTION 21: FLAGGING SYSTEM
```sql
-- Domain: internal/domain/flag/
-- Entity: flag/entity.go
-- =========================================

CREATE TABLE proposal_flags (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- Reporter
    reported_by UUID NOT NULL,
    reporter_type VARCHAR(20) CHECK (
        reporter_type IN ('CLIENT', 'FREELANCER', 'ADMIN', 'SYSTEM')
    ),

    -- Flag Details
    flag_type VARCHAR(50) NOT NULL CHECK (
        flag_type IN ('SPAM', 'INAPPROPRIATE_CONTENT', 'PLAGIARISM', 'MISLEADING',
                      'OFFENSIVE', 'PRICING_ISSUE', 'OTHER')
    ),
    flag_reason TEXT NOT NULL,

    -- Evidence
    evidence_urls TEXT[],
    evidence_description TEXT,

    -- Priority
    priority VARCHAR(20) DEFAULT 'MEDIUM' CHECK (
        priority IN ('LOW', 'MEDIUM', 'HIGH', 'URGENT')
    ),

    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'UNDER_REVIEW', 'RESOLVED', 'DISMISSED', 'ESCALATED')
    ),

    -- Review
    reviewed_by UUID,
    reviewed_at TIMESTAMPTZ,
    review_decision VARCHAR(50),
    review_notes TEXT,

    -- Resolution
    resolution_action VARCHAR(50),
    resolved_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_flags_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_flags_proposal ON proposal_flags (proposal_id);
CREATE INDEX idx_proposal_flags_reporter ON proposal_flags (reported_by);
CREATE INDEX idx_proposal_flags_status ON proposal_flags (status, priority, created_at DESC);

COMMENT ON TABLE proposal_flags IS 'Proposal flags - maps to internal/domain/flag/entity.go';

```
=========================================
##  SECTION 22: COMPLIANCE
```sql
-- Domain: internal/domain/compliance/
-- Entity: compliance/entity.go
-- =========================================

CREATE TABLE proposal_compliance (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL UNIQUE,

    -- Compliance Checks
    terms_accepted BOOLEAN DEFAULT FALSE,
    terms_version VARCHAR(20),
    terms_accepted_at TIMESTAMPTZ,

    -- Content Validation
    has_prohibited_content BOOLEAN DEFAULT FALSE,
    prohibited_content_types TEXT[],

    -- Geographic Compliance
    complies_with_local_laws BOOLEAN DEFAULT TRUE,
    jurisdiction_checks JSONB,

    -- PII Detection
    contains_pii BOOLEAN DEFAULT FALSE,
    pii_types TEXT[], -- EMAIL, PHONE, ADDRESS, etc
    pii_redacted BOOLEAN DEFAULT FALSE,

    -- Export Control
    export_control_flag BOOLEAN DEFAULT FALSE,
    export_control_reason TEXT,

    -- Sanctions Screening
    sanctions_screened BOOLEAN DEFAULT FALSE,
    sanctions_clear BOOLEAN DEFAULT TRUE,
    sanctions_screening_at TIMESTAMPTZ,

    -- Legal Holds
    legal_hold BOOLEAN DEFAULT FALSE,
    legal_hold_reason TEXT,
    legal_hold_placed_at TIMESTAMPTZ,
    legal_hold_placed_by UUID,

    -- Audit Trail
    last_compliance_check_at TIMESTAMPTZ,
    next_compliance_check_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_compliance_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_compliance_proposal ON proposal_compliance (proposal_id);
CREATE INDEX idx_proposal_compliance_issues ON proposal_compliance (has_prohibited_content)
    WHERE has_prohibited_content = TRUE OR contains_pii = TRUE;

COMMENT ON TABLE proposal_compliance IS 'Compliance tracking - maps to internal/domain/compliance/entity.go';

```
=========================================
##  SECTION 23: CLIENT INTERACTION - INTERVIEWS
```sql
-- Domain: internal/domain/interview/
-- Entity: interview/entity.go
-- =========================================

CREATE TABLE interviews (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,
    job_id UUID NOT NULL,
    freelancer_id UUID NOT NULL,
    client_id UUID NOT NULL,

    -- Interview Details
    interview_type VARCHAR(20) DEFAULT 'VIDEO' CHECK (
        interview_type IN ('VIDEO', 'PHONE', 'IN_PERSON', 'CHAT')
    ),

    -- Scheduling
    scheduled_at TIMESTAMPTZ,
    duration_minutes INTEGER DEFAULT 30,
    timezone VARCHAR(50),

    -- Meeting Details
    meeting_url TEXT,
    meeting_id VARCHAR(100),
    meeting_password VARCHAR(50),
    location TEXT, -- for in-person interviews

    -- Status
    status VARCHAR(20) DEFAULT 'SCHEDULED' CHECK (
        status IN ('SCHEDULED', 'CONFIRMED', 'RESCHEDULED', 'COMPLETED', 'CANCELLED', 'NO_SHOW')
    ),

    -- Reminders
    reminder_sent BOOLEAN DEFAULT FALSE,
    reminder_sent_at TIMESTAMPTZ,

    -- Attendance
    freelancer_joined_at TIMESTAMPTZ,
    client_joined_at TIMESTAMPTZ,
    started_at TIMESTAMPTZ,
    ended_at TIMESTAMPTZ,
    actual_duration_minutes INTEGER,

    -- Outcome
    outcome VARCHAR(20) CHECK (
        outcome IN ('POSITIVE', 'NEGATIVE', 'NEUTRAL', 'PENDING')
    ),

    -- Notes
    client_notes TEXT,
    freelancer_notes TEXT,

    -- Cancellation
    cancelled_by UUID,
    cancellation_reason TEXT,
    cancelled_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_interviews_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_interviews_proposal ON interviews (proposal_id);
CREATE INDEX idx_interviews_freelancer ON interviews (freelancer_id, status);
CREATE INDEX idx_interviews_client ON interviews (client_id, status);
CREATE INDEX idx_interviews_scheduled ON interviews (scheduled_at) WHERE status = 'SCHEDULED';

COMMENT ON TABLE interviews IS 'Interviews - maps to internal/domain/interview/entity.go';

```
=========================================
##  SECTION 24: CLIENT FEEDBACK
```sql
-- Domain: internal/domain/feedback/
-- Entity: feedback/entity.go
-- =========================================

CREATE TABLE proposal_feedback (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,
    client_id UUID NOT NULL,

    -- Overall Assessment
    overall_rating INTEGER CHECK (overall_rating BETWEEN 1 AND 5),

    -- Detailed Ratings
    skills_match_rating INTEGER CHECK (skills_match_rating BETWEEN 1 AND 5),
    experience_rating INTEGER CHECK (experience_rating BETWEEN 1 AND 5),
    communication_rating INTEGER CHECK (communication_rating BETWEEN 1 AND 5),
    professionalism_rating INTEGER CHECK (professionalism_rating BETWEEN 1 AND 5),
    value_rating INTEGER CHECK (value_rating BETWEEN 1 AND 5),

    -- Written Feedback
    feedback_text TEXT,
    strengths TEXT,
    concerns TEXT,

    -- Decision
    would_hire BOOLEAN,
    decision_reason TEXT,

    -- Tags
    positive_tags TEXT[],
    negative_tags TEXT[],

    -- Visibility
    is_private BOOLEAN DEFAULT TRUE,
    is_shared_with_freelancer BOOLEAN DEFAULT FALSE,
    shared_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_feedback_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE,
    CONSTRAINT uk_proposal_feedback UNIQUE (proposal_id, client_id)
);

CREATE INDEX idx_proposal_feedback_proposal ON proposal_feedback (proposal_id);
CREATE INDEX idx_proposal_feedback_client ON proposal_feedback (client_id);
CREATE INDEX idx_proposal_feedback_rating ON proposal_feedback (overall_rating DESC);

COMMENT ON TABLE proposal_feedback IS 'Client feedback - maps to internal/domain/feedback/entity.go';

```
=========================================
##  SECTION 25: SHORTLISTING
```sql
-- Domain: internal/domain/shortlist/
-- Entity: shortlist/entity.go
-- =========================================

CREATE TABLE shortlists (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,
    job_id UUID NOT NULL,
    client_id UUID NOT NULL,

    -- Shortlist Details
    rank INTEGER,
    notes TEXT,
    tags TEXT[],

    -- Decision Factors
    strengths JSONB,
    concerns JSONB,
    decision_criteria JSONB,

    -- Comparison
    comparison_score DECIMAL(5, 2),
    compared_with UUID[], -- Other proposal IDs in shortlist

    -- Status
    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (
        status IN ('ACTIVE', 'REMOVED', 'MOVED_TO_INTERVIEW', 'SELECTED', 'DECLINED')
    ),

    -- Actions
    removed_reason TEXT,
    removed_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_shortlists_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE,
    CONSTRAINT uk_shortlists UNIQUE (proposal_id, job_id)
);

CREATE INDEX idx_shortlists_proposal ON shortlists (proposal_id);
CREATE INDEX idx_shortlists_job ON shortlists (job_id, rank);
CREATE INDEX idx_shortlists_client ON shortlists (client_id);

COMMENT ON TABLE shortlists IS 'Shortlists - maps to internal/domain/shortlist/entity.go';

```
=========================================
##  SECTION 26: CONVERSATION TRACKING
```sql
-- Domain: internal/domain/conversation/
-- Entity: conversation/entity.go
-- =========================================

CREATE TABLE proposal_conversations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL UNIQUE,

    -- Conversation Reference (from communications-be)
    conversation_id UUID NOT NULL,

    -- Participants
    freelancer_id UUID NOT NULL,
    client_id UUID NOT NULL,

    -- Activity
    messages_count INTEGER DEFAULT 0,
    last_message_at TIMESTAMPTZ,
    last_message_by UUID,

    -- Response Metrics
    freelancer_avg_response_time_minutes INTEGER,
    client_avg_response_time_minutes INTEGER,

    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    is_archived BOOLEAN DEFAULT FALSE,
    archived_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_conversations_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_conversations_proposal ON proposal_conversations (proposal_id);
CREATE INDEX idx_proposal_conversations_conv ON proposal_conversations (conversation_id);
CREATE INDEX idx_proposal_conversations_active ON proposal_conversations (is_active) WHERE is_active = TRUE;

COMMENT ON TABLE proposal_conversations IS 'Conversation tracking - maps to internal/domain/conversation/entity.go';

```
=========================================
##  SECTION 27: WORKFLOW & COLLABORATION - NEGOTIATION
```sql
-- Domain: internal/domain/negotiation/
-- Entity: negotiation/entity.go
-- =========================================

CREATE TABLE negotiations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- Current Offer
    current_amount DECIMAL(12, 2) NOT NULL,
    current_currency CHAR(3) DEFAULT 'USD',
    current_terms JSONB,

    -- Negotiation Status
    status VARCHAR(20) DEFAULT 'OPEN' CHECK (
        status IN ('OPEN', 'COUNTERED', 'ACCEPTED', 'DECLINED', 'EXPIRED')
    ),

    -- Last Action
    last_offer_by UUID, -- USER_ID who made last offer
    last_offer_at TIMESTAMPTZ,

    -- Limits
    freelancer_min_amount DECIMAL(12, 2),
    client_max_amount DECIMAL(12, 2),

    -- Rounds
    negotiation_rounds INTEGER DEFAULT 0,
    max_rounds INTEGER DEFAULT 5,

    -- Expiration
    expires_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_negotiations_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_negotiations_proposal ON negotiations (proposal_id);
CREATE INDEX idx_negotiations_status ON negotiations (status, updated_at DESC);

COMMENT ON TABLE negotiations IS 'Negotiations - maps to internal/domain/negotiation/entity.go';

-- Negotiation History
CREATE TABLE negotiation_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    negotiation_id UUID NOT NULL,

    -- Offer Details
    offered_by UUID NOT NULL,
    offered_amount DECIMAL(12, 2) NOT NULL,
    offered_currency CHAR(3) DEFAULT 'USD',
    offered_terms JSONB,

    -- Message
    message TEXT,

    -- Response
    response VARCHAR(20) CHECK (
        response IN ('PENDING', 'ACCEPTED', 'COUNTERED', 'DECLINED')
    ),
    response_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_negotiation_history_negotiation FOREIGN KEY (negotiation_id)
        REFERENCES negotiations(id) ON DELETE CASCADE
);

CREATE INDEX idx_negotiation_history_negotiation ON negotiation_history (negotiation_id, created_at DESC);

```
=========================================
##  SECTION 28: INVITATIONS
```sql
-- Domain: internal/domain/invite/
-- Entity: invite/entity.go
-- =========================================

CREATE TABLE invitations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    job_id UUID NOT NULL,
    freelancer_id UUID NOT NULL,
    client_id UUID NOT NULL,

    -- Invitation Details
    message TEXT,
    custom_terms JSONB,

    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'VIEWED', 'ACCEPTED', 'DECLINED', 'EXPIRED', 'WITHDRAWN')
    ),

    -- Tracking
    sent_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    viewed_at TIMESTAMPTZ,
    responded_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ,

    -- Response
    response_message TEXT,
    decline_reason VARCHAR(100),

    -- Resulting Proposal
    proposal_id UUID,

    -- Withdrawal
    withdrawn_by UUID,
    withdrawn_at TIMESTAMPTZ,
    withdrawal_reason TEXT,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT uk_invitations UNIQUE (job_id, freelancer_id)
);

CREATE INDEX idx_invitations_job ON invitations (job_id, status);
CREATE INDEX idx_invitations_freelancer ON invitations (freelancer_id, status);
CREATE INDEX idx_invitations_client ON invitations (client_id);
CREATE INDEX idx_invitations_pending ON invitations (status, expires_at) WHERE status = 'PENDING';

COMMENT ON TABLE invitations IS 'Job invitations - maps to internal/domain/invite/entity.go';

```
=========================================
##  SECTION 29: REVISION TRACKING
```sql
-- Domain: internal/domain/revision/
-- Entity: revision/entity.go
-- =========================================

CREATE TABLE proposal_revisions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- Revision Details
    revision_number INTEGER NOT NULL,

    -- Snapshot of Changed Fields
    changed_fields TEXT[] NOT NULL,
    old_values JSONB,
    new_values JSONB,

    -- Change Context
    change_reason VARCHAR(100),
    change_notes TEXT,

    -- Change Type
    change_type VARCHAR(20) CHECK (
        change_type IN ('MINOR_EDIT', 'MAJOR_EDIT', 'PRICING_CHANGE', 'CONTENT_CHANGE', 'AUTO_CORRECTION')
    ),

    -- Changed By
    changed_by UUID NOT NULL,
    changed_by_type VARCHAR(20) CHECK (
        changed_by_type IN ('FREELANCER', 'SYSTEM', 'ADMIN')
    ),

    -- Impact Assessment
    impact_level VARCHAR(20) CHECK (
        impact_level IN ('LOW', 'MEDIUM', 'HIGH')
    ),

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_revisions_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_revisions_proposal ON proposal_revisions (proposal_id, revision_number DESC);
CREATE INDEX idx_proposal_revisions_changed_by ON proposal_revisions (changed_by);

COMMENT ON TABLE proposal_revisions IS 'Proposal revision history - maps to internal/domain/revision/entity.go';

```
=========================================
##  SECTION 30: COLLABORATION (TEAM PROPOSALS)
```sql
-- Domain: internal/domain/collaboration/
-- Entity: collaboration/entity.go
-- =========================================

CREATE TABLE team_proposals (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL UNIQUE,

    -- Team Lead
    lead_freelancer_id UUID NOT NULL,

    -- Team Configuration
    team_name VARCHAR(200),
    total_members INTEGER DEFAULT 1,

    -- Revenue Split
    revenue_split_model VARCHAR(20) DEFAULT 'PERCENTAGE' CHECK (
        revenue_split_model IN ('PERCENTAGE', 'FIXED', 'MILESTONE_BASED', 'CUSTOM')
    ),
    split_configuration JSONB,

    -- Status
    is_active BOOLEAN DEFAULT TRUE,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_team_proposals_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_team_proposals_proposal ON team_proposals (proposal_id);
CREATE INDEX idx_team_proposals_lead ON team_proposals (lead_freelancer_id);

COMMENT ON TABLE team_proposals IS 'Team proposals - maps to internal/domain/collaboration/entity.go';

-- Team Members
CREATE TABLE team_members (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    team_proposal_id UUID NOT NULL,
    freelancer_id UUID NOT NULL,

    -- Role in Team
    role VARCHAR(100),
    responsibilities TEXT,

    -- Compensation
    revenue_percentage DECIMAL(5, 2),
    fixed_amount DECIMAL(12, 2),

    -- Status
    status VARCHAR(20) DEFAULT 'INVITED' CHECK (
        status IN ('INVITED', 'ACCEPTED', 'DECLINED', 'ACTIVE', 'REMOVED')
    ),

    -- Timestamps
    invited_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    responded_at TIMESTAMPTZ,
    joined_at TIMESTAMPTZ,
    removed_at TIMESTAMPTZ,
    removal_reason TEXT,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_team_members_team FOREIGN KEY (team_proposal_id)
        REFERENCES team_proposals(id) ON DELETE CASCADE,
    CONSTRAINT uk_team_members UNIQUE (team_proposal_id, freelancer_id)
);

CREATE INDEX idx_team_members_team ON team_members (team_proposal_id);
CREATE INDEX idx_team_members_freelancer ON team_members (freelancer_id, status);

```
=========================================
##  SECTION 31: LIFECYCLE MANAGEMENT - EXPIRATION
```sql
-- Domain: internal/domain/expiration/
-- Entity: expiration/entity.go
-- =========================================

CREATE TABLE proposal_expirations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL UNIQUE,

    -- Expiration Configuration
    expires_at TIMESTAMPTZ NOT NULL,
    expiration_reason VARCHAR(100),

    -- Warnings
    warning_sent BOOLEAN DEFAULT FALSE,
    warning_sent_at TIMESTAMPTZ,
    final_warning_sent BOOLEAN DEFAULT FALSE,
    final_warning_sent_at TIMESTAMPTZ,

    -- Extension
    can_be_extended BOOLEAN DEFAULT TRUE,
    extension_count INTEGER DEFAULT 0,
    max_extensions INTEGER DEFAULT 3,
    last_extended_at TIMESTAMPTZ,

    -- Status
    is_expired BOOLEAN DEFAULT FALSE,
    expired_at TIMESTAMPTZ,

    -- Actions on Expiry
    auto_archive BOOLEAN DEFAULT TRUE,
    notify_on_expiry BOOLEAN DEFAULT TRUE,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_expirations_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_expirations_proposal ON proposal_expirations (proposal_id);
CREATE INDEX idx_proposal_expirations_expires ON proposal_expirations (expires_at)
    WHERE is_expired = FALSE;

COMMENT ON TABLE proposal_expirations IS 'Expiration tracking - maps to internal/domain/expiration/entity.go';

```
=========================================
##  SECTION 32: WITHDRAWAL TRACKING
```sql
-- Domain: internal/domain/withdrawal/
-- Entity: withdrawal/entity.go
-- =========================================

CREATE TABLE proposal_withdrawals (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL UNIQUE,

    -- Withdrawal Details
    withdrawn_by UUID NOT NULL,
    withdrawal_reason VARCHAR(100) NOT NULL CHECK (
        withdrawal_reason IN ('FOUND_BETTER_OPPORTUNITY', 'NOT_INTERESTED', 'UNAVAILABLE',
                              'CLIENT_REQUEST', 'PRICING_ISSUE', 'OTHER')
    ),
    withdrawal_details TEXT,

    -- Impact
    connects_refunded BOOLEAN DEFAULT FALSE,
    refund_amount INTEGER,

    -- Status
    can_resubmit BOOLEAN DEFAULT TRUE,
    resubmitted BOOLEAN DEFAULT FALSE,
    resubmitted_at TIMESTAMPTZ,
    new_proposal_id UUID,

    withdrawn_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_withdrawals_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_withdrawals_proposal ON proposal_withdrawals (proposal_id);
CREATE INDEX idx_proposal_withdrawals_freelancer ON proposal_withdrawals (withdrawn_by);

COMMENT ON TABLE proposal_withdrawals IS 'Withdrawal tracking - maps to internal/domain/withdrawal/entity.go';

```
=========================================
##  SECTION 33: ARCHIVING
```sql
-- Domain: internal/domain/archive/
-- Entity: archive/entity.go
-- =========================================

CREATE TABLE proposal_archives (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL UNIQUE,

    -- Archive Details
    archived_by UUID NOT NULL,
    archive_reason VARCHAR(100) CHECK (
        archive_reason IN ('AUTO_EXPIRED', 'MANUAL', 'JOB_CLOSED', 'ACCEPTED_ELSEWHERE',
                           'WITHDRAWN', 'COMPLIANCE', 'OTHER')
    ),
    archive_notes TEXT,

    -- Retention
    retention_period_days INTEGER DEFAULT 365,
    delete_after TIMESTAMPTZ,

    -- Status
    can_restore BOOLEAN DEFAULT TRUE,
    restored BOOLEAN DEFAULT FALSE,
    restored_at TIMESTAMPTZ,
    restored_by UUID,

    archived_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_archives_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_archives_proposal ON proposal_archives (proposal_id);
CREATE INDEX idx_proposal_archives_delete ON proposal_archives (delete_after)
    WHERE delete_after IS NOT NULL;

COMMENT ON TABLE proposal_archives IS 'Archive tracking - maps to internal/domain/archive/entity.go';

```
=========================================
##  SECTION 34: PIPELINE TRACKING
```sql
-- Domain: internal/domain/pipeline/
-- Entity: pipeline/entity.go
-- =========================================

CREATE TABLE proposal_pipelines (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    freelancer_id UUID NOT NULL,

    -- Pipeline Stage
    stage VARCHAR(30) NOT NULL CHECK (
        stage IN ('DRAFT', 'SUBMITTED', 'UNDER_REVIEW', 'SHORTLISTED',
                  'INTERVIEWING', 'NEGOTIATING', 'FINALIZING', 'CLOSED')
    ),

    -- Proposals in Stage
    proposal_ids UUID[],
    proposals_count INTEGER DEFAULT 0,

    -- Stage Health
    avg_time_in_stage_hours INTEGER,
    stage_conversion_rate DECIMAL(5, 2),

    -- Metrics
    total_value DECIMAL(12, 2),
    currency CHAR(3) DEFAULT 'USD',

    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT uk_proposal_pipelines UNIQUE (freelancer_id, stage)
);

CREATE INDEX idx_proposal_pipelines_freelancer ON proposal_pipelines (freelancer_id);

COMMENT ON TABLE proposal_pipelines IS 'Pipeline tracking - maps to internal/domain/pipeline/entity.go';

```
=========================================
##  SECTION 35: PROPOSAL RECYCLING
```sql
-- Domain: internal/domain/recycling/
-- Entity: recycling/entity.go
-- =========================================

CREATE TABLE proposal_recycling (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    original_proposal_id UUID NOT NULL,
    new_proposal_id UUID,

    -- Recycling Details
    recycled_for_job_id UUID NOT NULL,

    -- Reuse Strategy
    reuse_strategy VARCHAR(20) CHECK (
        reuse_strategy IN ('FULL_COPY', 'TEMPLATE', 'PARTIAL', 'ADAPTED')
    ),

    -- Fields Reused
    reused_fields TEXT[],
    modified_fields TEXT[],

    -- Adaptations
    adaptations_made JSONB,

    -- Performance
    original_performance_score DECIMAL(5, 2),
    improvement_suggestions TEXT[],

    recycled_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_recycling_original FOREIGN KEY (original_proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE,
    CONSTRAINT fk_proposal_recycling_new FOREIGN KEY (new_proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_recycling_original ON proposal_recycling (original_proposal_id);
CREATE INDEX idx_proposal_recycling_new ON proposal_recycling (new_proposal_id);

COMMENT ON TABLE proposal_recycling IS 'Proposal recycling - maps to internal/domain/recycling/entity.go';

```
=========================================
##  SECTION 36: RECOMMENDATION ENGINE
```sql
-- Domain: internal/domain/recommendation/
-- Entity: recommendation/entity.go
-- =========================================

CREATE TABLE proposal_recommendations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- Recommendation Type
    recommendation_type VARCHAR(30) CHECK (
        recommendation_type IN ('PRICING_ADJUSTMENT', 'CONTENT_IMPROVEMENT', 'SKILL_HIGHLIGHT',
                                'PORTFOLIO_ADD', 'TIMING_SUGGESTION', 'BIDDING_STRATEGY')
    ),

    -- Recommendation Details
    title VARCHAR(200) NOT NULL,
    description TEXT,

    -- Impact
    expected_impact VARCHAR(20) CHECK (
        expected_impact IN ('LOW', 'MEDIUM', 'HIGH')
    ),
    confidence_score DECIMAL(5, 2),

    -- Suggested Action
    action_type VARCHAR(50),
    action_details JSONB,

    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'VIEWED', 'APPLIED', 'DISMISSED', 'EXPIRED')
    ),

    -- Tracking
    viewed_at TIMESTAMPTZ,
    applied_at TIMESTAMPTZ,
    dismissed_at TIMESTAMPTZ,

    -- Results (if applied)
    result_metrics JSONB,
    was_helpful BOOLEAN,

    expires_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_recommendations_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_recommendations_proposal ON proposal_recommendations (proposal_id);
CREATE INDEX idx_proposal_recommendations_status ON proposal_recommendations (status);
CREATE INDEX idx_proposal_recommendations_expires ON proposal_recommendations (expires_at)
    WHERE status = 'PENDING';

COMMENT ON TABLE proposal_recommendations IS 'AI recommendations - maps to internal/domain/recommendation/entity.go';

```
=========================================
##  SECTION 37: CONTEXT ENRICHMENT
```sql
-- Domain: internal/domain/context/
-- Entity: context/entity.go
-- =========================================

CREATE TABLE proposal_context (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL UNIQUE,

    -- Job Context
    job_requirements JSONB,
    job_skills JSONB,
    job_budget_range JSONB,
    client_preferences JSONB,

    -- Freelancer Context
    freelancer_skills JSONB,
    freelancer_experience_level VARCHAR(20),
    freelancer_success_rate DECIMAL(5, 2),
    freelancer_avg_rating DECIMAL(3, 2),

    -- Market Context
    avg_proposals_for_job INTEGER,
    competition_level VARCHAR(20) CHECK (
        competition_level IN ('LOW', 'MEDIUM', 'HIGH', 'VERY_HIGH')
    ),
    similar_jobs_pricing JSONB,

    -- Timing Context
    job_posted_at TIMESTAMPTZ,
    time_since_job_posted_hours INTEGER,
    proposal_rank_by_time INTEGER,

    -- Historical Context
    freelancer_previous_proposals_count INTEGER,
    freelancer_win_rate DECIMAL(5, 2),
    client_hire_history JSONB,

    computed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_context_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_context_proposal ON proposal_context (proposal_id);

COMMENT ON TABLE proposal_context IS 'Context enrichment - maps to internal/domain/context/entity.go';

```
=========================================
##  SECTION 38: URGENCY INDICATORS
```sql
-- Domain: internal/domain/urgency/
-- Entity: urgency/entity.go
-- =========================================

CREATE TABLE proposal_urgency (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL UNIQUE,

    -- Urgency Score
    urgency_score INTEGER DEFAULT 0 CHECK (urgency_score BETWEEN 0 AND 100),
    urgency_level VARCHAR(20) CHECK (
        urgency_level IN ('NONE', 'LOW', 'MEDIUM', 'HIGH', 'CRITICAL')
    ),

    -- Urgency Factors
    job_closing_soon BOOLEAN DEFAULT FALSE,
    high_competition BOOLEAN DEFAULT FALSE,
    client_active BOOLEAN DEFAULT FALSE,
    multiple_interviews_scheduled BOOLEAN DEFAULT FALSE,
    price_sensitive BOOLEAN DEFAULT FALSE,

    -- Timing
    job_closes_at TIMESTAMPTZ,
    hours_until_close INTEGER,

    -- Recommended Actions
    recommended_actions TEXT[],
    action_deadline TIMESTAMPTZ,

    computed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_urgency_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_urgency_proposal ON proposal_urgency (proposal_id);
CREATE INDEX idx_proposal_urgency_level ON proposal_urgency (urgency_level)
    WHERE urgency_level IN ('HIGH', 'CRITICAL');

COMMENT ON TABLE proposal_urgency IS 'Urgency tracking - maps to internal/domain/urgency/entity.go';

```
=========================================
##  SECTION 39: RISK ASSESSMENT
```sql
-- Domain: internal/domain/risk_assessment/
-- Entity: risk_assessment/entity.go
-- =========================================

CREATE TABLE proposal_risk_assessments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- Overall Risk
    overall_risk_score DECIMAL(5, 2) NOT NULL,
    risk_level VARCHAR(20) CHECK (
        risk_level IN ('LOW', 'MEDIUM', 'HIGH', 'CRITICAL')
    ),

    -- Risk Categories
    pricing_risk_score DECIMAL(5, 2),
    timeline_risk_score DECIMAL(5, 2),
    scope_risk_score DECIMAL(5, 2),
    client_risk_score DECIMAL(5, 2),

    -- Risk Factors
    risk_factors JSONB,
    red_flags TEXT[],

    -- Mitigation
    recommended_mitigations TEXT[],
    mitigation_notes TEXT,

    -- Confidence
    confidence_level DECIMAL(5, 2),

    -- Review
    requires_manual_review BOOLEAN DEFAULT FALSE,
    reviewed_by UUID,
    reviewed_at TIMESTAMPTZ,
    review_decision VARCHAR(20),

    assessed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_risk_assessments_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_risk_assessments_proposal ON proposal_risk_assessments (proposal_id);
CREATE INDEX idx_proposal_risk_assessments_level ON proposal_risk_assessments (risk_level, assessed_at DESC);

COMMENT ON TABLE proposal_risk_assessments IS 'Risk assessment - maps to internal/domain/risk_assessment/entity.go';

```
=========================================
##  SECTION 40: AI ASSIST & OPTIMIZATION
```sql
-- Domain: internal/domain/ai_assist/ (consolidated)
-- Entity: ai_assist/entity.go
-- =========================================

CREATE TABLE proposal_ai_suggestions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- Suggestion Type
    suggestion_type VARCHAR(50) NOT NULL CHECK (
        suggestion_type IN ('TITLE_OPTIMIZATION', 'COVER_LETTER_IMPROVEMENT', 'PRICING_RECOMMENDATION',
                           'SKILL_EMPHASIS', 'TONE_ADJUSTMENT', 'LENGTH_OPTIMIZATION', 'GRAMMAR_FIX')
    ),

    -- Original & Suggested
    original_text TEXT,
    suggested_text TEXT,

    -- Reasoning
    reasoning TEXT,
    expected_improvement VARCHAR(200),

    -- Confidence
    confidence_score DECIMAL(5, 2),

    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'ACCEPTED', 'REJECTED', 'MODIFIED', 'EXPIRED')
    ),

    -- User Action
    user_action_at TIMESTAMPTZ,
    user_feedback VARCHAR(20) CHECK (
        user_feedback IN ('HELPFUL', 'NOT_HELPFUL', 'PARTIALLY_HELPFUL')
    ),

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_ai_suggestions_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_ai_suggestions_proposal ON proposal_ai_suggestions (proposal_id, status);
CREATE INDEX idx_proposal_ai_suggestions_type ON proposal_ai_suggestions (suggestion_type);

COMMENT ON TABLE proposal_ai_suggestions IS 'AI suggestions - maps to internal/domain/ai_assist/entity.go';

-- AI Optimizations Applied
CREATE TABLE proposal_ai_optimizations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- Optimization Details
    optimization_type VARCHAR(50) NOT NULL,
    field_optimized VARCHAR(100),

    -- Changes Made
    before_value TEXT,
    after_value TEXT,

    -- Impact
    expected_impact_score DECIMAL(5, 2),
    actual_impact_score DECIMAL(5, 2),

    -- Metrics Before/After
    metrics_before JSONB,
    metrics_after JSONB,

    applied_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_ai_optimizations_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_ai_optimizations_proposal ON proposal_ai_optimizations (proposal_id);

```
=========================================
##  SECTION 41: SKILL MATCHING
```sql
-- Domain: internal/domain/skill_match/
-- Entity: skill_match/entity.go
-- =========================================

CREATE TABLE proposal_skill_matches (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL UNIQUE,

    -- Overall Match
    overall_match_score DECIMAL(5, 2),
    match_percentage INTEGER CHECK (match_percentage BETWEEN 0 AND 100),

    -- Required Skills Analysis
    required_skills_matched INTEGER DEFAULT 0,
    required_skills_total INTEGER DEFAULT 0,
    required_skills_match_rate DECIMAL(5, 2),

    -- Preferred Skills Analysis
    preferred_skills_matched INTEGER DEFAULT 0,
    preferred_skills_total INTEGER DEFAULT 0,

    -- Skill Details
    matched_skills JSONB, -- {skill_id: proficiency_level}
    missing_skills JSONB,
    exceeding_skills JSONB, -- Skills freelancer has beyond requirements

    -- Experience Level Match
    experience_match BOOLEAN DEFAULT FALSE,
    experience_gap VARCHAR(20), -- UNDER, OVER, PERFECT

    -- Suggestions
    skill_gap_suggestions TEXT[],
    learning_recommendations JSONB,

    computed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_skill_matches_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_skill_matches_proposal ON proposal_skill_matches (proposal_id);
CREATE INDEX idx_proposal_skill_matches_score ON proposal_skill_matches (overall_match_score DESC);

COMMENT ON TABLE proposal_skill_matches IS 'Skill matching - maps to internal/domain/skill_match/entity.go';

```
=========================================
##  SECTION 42: VIDEO INTRODUCTIONS
```sql
-- Domain: internal/domain/video_introduction/
-- Entity: video_introduction/entity.go
-- =========================================

CREATE TABLE video_introductions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL UNIQUE,

    -- Video Reference (storage-be)
    video_file_id UUID NOT NULL,
    video_url TEXT NOT NULL,
    thumbnail_url TEXT,

    -- Video Metadata
    duration_seconds INTEGER,
    file_size_bytes BIGINT,
    video_format VARCHAR(20),
    resolution VARCHAR(20),

    -- Processing Status
    processing_status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        processing_status IN ('PENDING', 'PROCESSING', 'COMPLETED', 'FAILED')
    ),
    transcoding_status VARCHAR(20),

    -- Transcript
    transcript TEXT,
    transcript_language VARCHAR(10),
    has_captions BOOLEAN DEFAULT FALSE,
    caption_file_url TEXT,

    -- Content Analysis
    audio_quality_score DECIMAL(5, 2),
    video_quality_score DECIMAL(5, 2),
    professionalism_score DECIMAL(5, 2),

    -- Compliance
    contains_inappropriate_content BOOLEAN DEFAULT FALSE,
    content_warnings TEXT[],

    -- Performance
    views_count INTEGER DEFAULT 0,
    avg_watch_percentage DECIMAL(5, 2),
    completion_rate DECIMAL(5, 2),

    uploaded_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    processed_at TIMESTAMPTZ,

    CONSTRAINT fk_video_introductions_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_video_introductions_proposal ON video_introductions (proposal_id);
CREATE INDEX idx_video_introductions_status ON video_introductions (processing_status);

COMMENT ON TABLE video_introductions IS 'Video introductions - maps to internal/domain/video_introduction/entity.go';

```
=========================================
##  SECTION 43: REFERENCES
```sql
-- Domain: internal/domain/reference/
-- Entity: reference/entity.go
-- =========================================

CREATE TABLE proposal_references (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- Reference Details
    reference_name VARCHAR(200),
    reference_title VARCHAR(200),
    reference_company VARCHAR(200),
    reference_relationship VARCHAR(100),

    -- Contact (hashed/encrypted)
    contact_email_hash VARCHAR(64),
    contact_phone_hash VARCHAR(64),

    -- Reference Content
    reference_text TEXT,
    reference_rating INTEGER CHECK (reference_rating BETWEEN 1 AND 5),

    -- Verification
    is_verified BOOLEAN DEFAULT FALSE,
    verified_at TIMESTAMPTZ,
    verification_method VARCHAR(50),

    -- Request Tracking
    verification_requested BOOLEAN DEFAULT FALSE,
    verification_request_sent_at TIMESTAMPTZ,
    verification_token VARCHAR(255),

    -- Visibility
    is_public BOOLEAN DEFAULT FALSE,
    show_contact_details BOOLEAN DEFAULT FALSE,

    -- Freshness
    reference_date DATE,
    is_recent BOOLEAN DEFAULT TRUE,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_proposal_references_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_proposal_references_proposal ON proposal_references (proposal_id);
CREATE INDEX idx_proposal_references_verified ON proposal_references (is_verified)
    WHERE is_verified = TRUE;

COMMENT ON TABLE proposal_references IS 'References - maps to internal/domain/reference/entity.go';

```
=========================================
##  SECTION 44: A/B TESTING & EXPERIMENTS
```sql
-- Domain: internal/domain/ab_testing/
-- Entity: ab_testing/entity.go
-- =========================================

CREATE TABLE proposal_experiments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    experiment_name VARCHAR(200) NOT NULL,
    experiment_type VARCHAR(50),

    -- Configuration
    variants JSONB NOT NULL, -- Array of variant configurations
    traffic_allocation JSONB, -- How traffic is split

    -- Status
    status VARCHAR(20) DEFAULT 'DRAFT' CHECK (
        status IN ('DRAFT', 'ACTIVE', 'PAUSED', 'COMPLETED', 'CANCELLED')
    ),

    -- Timing
    starts_at TIMESTAMPTZ,
    ends_at TIMESTAMPTZ,

    -- Target Criteria
    target_users JSONB,
    target_job_types JSONB,

    -- Success Metrics
    primary_metric VARCHAR(100),
    secondary_metrics TEXT[],

    -- Results
    results JSONB,
    winner_variant VARCHAR(50),
    confidence_level DECIMAL(5, 2),

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_proposal_experiments_status ON proposal_experiments (status);
CREATE INDEX idx_proposal_experiments_active ON proposal_experiments (starts_at, ends_at)
    WHERE status = 'ACTIVE';

-- Experiment Assignments
CREATE TABLE proposal_experiment_assignments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    experiment_id UUID NOT NULL,
    proposal_id UUID NOT NULL,

    -- Assignment
    variant VARCHAR(50) NOT NULL,
    assigned_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    -- Outcome Tracking
    converted BOOLEAN DEFAULT FALSE,
    converted_at TIMESTAMPTZ,
    conversion_value DECIMAL(12, 2),

    -- Metrics
    metrics JSONB,

    CONSTRAINT fk_proposal_experiment_assignments_experiment FOREIGN KEY (experiment_id)
        REFERENCES proposal_experiments(id) ON DELETE CASCADE,
    CONSTRAINT fk_proposal_experiment_assignments_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE,
    CONSTRAINT uk_proposal_experiment_assignments UNIQUE (experiment_id, proposal_id)
);

CREATE INDEX idx_proposal_experiment_assignments_experiment ON proposal_experiment_assignments (experiment_id);
CREATE INDEX idx_proposal_experiment_assignments_proposal ON proposal_experiment_assignments (proposal_id);

```
=========================================
##  SECTION 45: OUTBOX PATTERN FOR EVENTS
```sql
-- Domain: internal/domain/outbox/
-- Entity: outbox/entity.go
-- =========================================

CREATE TABLE outbox_events (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    -- Event Identity
    event_id VARCHAR(100) NOT NULL UNIQUE, -- ULID
    event_type VARCHAR(100) NOT NULL,
    event_version VARCHAR(10) DEFAULT 'v1',

    -- Aggregate
    aggregate_type VARCHAR(50) NOT NULL,
    aggregate_id UUID NOT NULL,

    -- Actor & Context
    actor_id UUID,
    tenant_id UUID,
    correlation_id UUID,
    causation_id UUID,

    -- Routing
    topic VARCHAR(100) NOT NULL,
    partition_key VARCHAR(255),

    -- Payload
    payload JSONB NOT NULL,
    schema_ref VARCHAR(64), -- Schema version hash

    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'PROCESSING', 'PUBLISHED', 'FAILED')
    ),

    -- Retry Logic
    attempts INTEGER DEFAULT 0,
    max_attempts INTEGER DEFAULT 3,
    last_attempt_at TIMESTAMPTZ,
    next_attempt_at TIMESTAMPTZ,
    error_message TEXT,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_question_answers_proposal FOREIGN KEY (proposal_id) REFERENCES proposals(id) ON DELETE CASCADE,
    CONSTRAINT uk_question_answers UNIQUE (proposal_id, question_id)
);

CREATE INDEX idx_question_answers_proposal ON question_answers (proposal_id);
CREATE INDEX idx_question_answers_question ON question_answers (question_id);

COMMENT ON TABLE question_answers IS 'Screening question answers - maps to internal/domain/question_answer/entity.go';

### ESSAM Start

-- Dead Letter Queue
CREATE TABLE outbox_dead_letter (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    event_id VARCHAR(100) NOT NULL,

    -- Original Event Data
    original_event JSONB NOT NULL,

    -- Failure Details
    failure_reason TEXT NOT NULL,
    attempts_made INTEGER,
    last_error TEXT,

    -- Resolution
    resolution_status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        resolution_status IN ('PENDING', 'RESOLVED', 'DISCARDED')
    ),
    resolved_at TIMESTAMPTZ,
    resolved_by UUID,
    resolution_notes TEXT,

    moved_to_dlq_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_outbox_dead_letter_status ON outbox_dead_letter (resolution_status);

    ```
=========================================
##  SECTION 46: AUDIT LOGS
-- =========================================

CREATE TABLE audit_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    -- Entity
    entity_type VARCHAR(100) NOT NULL,
    entity_id UUID NOT NULL,

    -- Action
    action VARCHAR(100) NOT NULL,
    action_category VARCHAR(50),

    -- Actor
    actor_user_id UUID,
    actor_type VARCHAR(20),
    actor_ip INET,
    actor_user_agent TEXT,

    -- Changes
    old_values JSONB,
    new_values JSONB,
    changed_fields TEXT[],

    -- Context
    request_id UUID,
    correlation_id UUID,
    session_id UUID,

    -- Compliance
    gdpr_relevant BOOLEAN DEFAULT FALSE,
    pii_accessed BOOLEAN DEFAULT FALSE,

    occurred_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_audit_logs_entity ON audit_logs (entity_type, entity_id, occurred_at DESC);
CREATE INDEX idx_audit_logs_actor ON audit_logs (actor_user_id, occurred_at DESC);
CREATE INDEX idx_audit_logs_action ON audit_logs (action, occurred_at DESC);
CREATE INDEX idx_audit_logs_compliance ON audit_logs (gdpr_relevant) WHERE gdpr_relevant = TRUE;

```
=========================================
##  SECTION 47: READ MODELS (CQRS PROJECTIONS)
```sql

-- =========================================

-- Proposal Read Model for Fast Queries
CREATE TABLE proposal_read_model (
    proposal_id UUID PRIMARY KEY,

    -- Core Info
    job_id UUID NOT NULL,
    freelancer_id UUID NOT NULL,
    status VARCHAR(20),

    -- Pricing
    proposed_amount DECIMAL(12, 2),
    currency CHAR(3),
    pricing_model VARCHAR(20),

    -- Freelancer Info (denormalized)
    freelancer_name VARCHAR(200),
    freelancer_title VARCHAR(200),
    freelancer_rating DECIMAL(3, 2),
    freelancer_success_rate DECIMAL(5, 2),

    -- Job Info (denormalized)
    job_title VARCHAR(200),
    job_status VARCHAR(20),
    client_id UUID,

    -- Metrics
    quality_score DECIMAL(5, 2),
    match_score DECIMAL(5, 2),
    views_count INTEGER,

    -- Timestamps
    submitted_at TIMESTAMPTZ,
    last_activity_at TIMESTAMPTZ,

    -- Search
    search_vector tsvector GENERATED ALWAYS AS (
        setweight(to_tsvector('english', COALESCE(job_title, '')), 'A') ||
        setweight(to_tsvector('english', COALESCE(freelancer_name, '')), 'B')
    ) STORED,

    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_proposal_read_job ON proposal_read_model (job_id, status);
CREATE INDEX idx_proposal_read_freelancer ON proposal_read_model (freelancer_id, status);
CREATE INDEX idx_proposal_read_search ON proposal_read_model USING gin(search_vector);
CREATE INDEX idx_proposal_read_quality ON proposal_read_model (quality_score DESC);
CREATE INDEX idx_proposal_read_submitted ON proposal_read_model (submitted_at DESC);

-- Freelancer Proposal Stats (Projection)
CREATE TABLE freelancer_proposal_stats (
    freelancer_id UUID PRIMARY KEY,

    -- Overall Stats
    total_proposals INTEGER DEFAULT 0,
    active_proposals INTEGER DEFAULT 0,
    successful_proposals INTEGER DEFAULT 0,

    -- Status Breakdown
    draft_count INTEGER DEFAULT 0,
    submitted_count INTEGER DEFAULT 0,
    shortlisted_count INTEGER DEFAULT 0,
    accepted_count INTEGER DEFAULT 0,
    rejected_count INTEGER DEFAULT 0,

    -- Success Metrics
    success_rate DECIMAL(5, 2),
    avg_quality_score DECIMAL(5, 2),
    avg_match_score DECIMAL(5, 2),

    -- Financial
    total_proposed_value DECIMAL(12, 2),
    avg_proposal_amount DECIMAL(12, 2),

    -- Response Metrics
    avg_response_time_hours INTEGER,
    response_rate DECIMAL(5, 2),

    -- Recent Activity
    last_proposal_at TIMESTAMPTZ,
    proposals_last_30_days INTEGER DEFAULT 0,

    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_freelancer_proposal_stats_success ON freelancer_proposal_stats (success_rate DESC);

-- Job Proposal Stats (Projection)
CREATE TABLE job_proposal_stats (
    job_id UUID PRIMARY KEY,

    -- Proposal Counts
    total_proposals INTEGER DEFAULT 0,
    qualified_proposals INTEGER DEFAULT 0,
    shortlisted_proposals INTEGER DEFAULT 0,

    -- Quality Distribution
    avg_quality_score DECIMAL(5, 2),
    avg_match_score DECIMAL(5, 2),
    high_quality_count INTEGER DEFAULT 0,

    -- Pricing Analysis
    avg_proposed_amount DECIMAL(12, 2),
    min_proposed_amount DECIMAL(12, 2),
    max_proposed_amount DECIMAL(12, 2),
    median_proposed_amount DECIMAL(12, 2),

    -- Timeline
    first_proposal_at TIMESTAMPTZ,
    last_proposal_at TIMESTAMPTZ,
    proposals_last_24h INTEGER DEFAULT 0,

    -- Competition
    competition_level VARCHAR(20),

    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_job_proposal_stats_competition ON job_proposal_stats (competition_level);

```
=========================================
##  SECTION 48: EXTERNAL REFERENCES
```sql

-- (Relations with other microservices)
-- =========================================

CREATE TABLE external_references (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    proposal_id UUID NOT NULL,

    -- External Service
    service_name VARCHAR(50) NOT NULL,
    reference_type VARCHAR(50) NOT NULL,
    reference_id UUID NOT NULL,

    -- Context
    context JSONB,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_external_references_proposal FOREIGN KEY (proposal_id)
        REFERENCES proposals(id) ON DELETE CASCADE
);

CREATE INDEX idx_external_references_proposal ON external_references (proposal_id);
CREATE INDEX idx_external_references_service ON external_references (service_name, reference_type, reference_id);

COMMENT ON TABLE external_references IS 'References to entities in other microservices';

```
=========================================
##  SECTION 49: NOTIFICATION PREFERENCES
```sql

-- =========================================

CREATE TABLE proposal_notification_preferences (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    freelancer_id UUID NOT NULL UNIQUE,

    -- Event Notifications
    notify_on_view BOOLEAN DEFAULT TRUE,
    notify_on_shortlist BOOLEAN DEFAULT TRUE,
    notify_on_interview BOOLEAN DEFAULT TRUE,
    notify_on_acceptance BOOLEAN DEFAULT TRUE,
    notify_on_rejection BOOLEAN DEFAULT FALSE,

    -- Bidding Notifications
    notify_on_outbid BOOLEAN DEFAULT TRUE,
    notify_on_winning_bid BOOLEAN DEFAULT TRUE,
    notify_on_auction_ending BOOLEAN DEFAULT TRUE,

    -- Recommendation Notifications
    notify_on_suggestions BOOLEAN DEFAULT TRUE,
    notify_on_job_match BOOLEAN DEFAULT TRUE,

    -- Channels
    email_enabled BOOLEAN DEFAULT TRUE,
    push_enabled BOOLEAN DEFAULT TRUE,
    sms_enabled BOOLEAN DEFAULT FALSE,

    -- Frequency
    digest_frequency VARCHAR(20) DEFAULT 'REAL_TIME' CHECK (
        digest_frequency IN ('REAL_TIME', 'HOURLY', 'DAILY', 'WEEKLY')
    ),
    quiet_hours_start TIME,
    quiet_hours_end TIME,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_proposal_notification_preferences_freelancer
    ON proposal_notification_preferences (freelancer_id);

```
=========================================
##  SECTION 50: PERFORMANCE INDEXES & VIEWS

```sql
CREATE INDEX idx_proposals_client_viewed ON proposals (client_viewed_at) WHERE client_viewed_at IS NOT NULL;
```
```sql
-- =========================================

-- Composite indexes for common queries
CREATE INDEX idx_proposals_job_status_quality
    ON proposals (job_id, status, quality_score DESC)
    WHERE is_deleted = FALSE;

CREATE INDEX idx_proposals_freelancer_submitted
    ON proposals (freelancer_id, submitted_at DESC)
    WHERE status IN ('SUBMITTED', 'UNDER_REVIEW', 'SHORTLISTED');

CREATE INDEX idx_proposals_active_boosted
    ON proposals (is_boosted, quality_score DESC)
    WHERE status = 'SUBMITTED' AND is_deleted = FALSE;

-- View for active proposals with key metrics
CREATE VIEW v_active_proposals AS
SELECT
    p.id,
    p.job_id,
    p.freelancer_id,
    p.title,
    p.status,
    p.proposed_amount,
    p.currency,
    p.quality_score,
    p.match_score,
    p.submitted_at,
    p.views_count,
    cl.word_count AS cover_letter_words,
    COUNT(DISTINCT a.id) AS attachments_count,
    COUNT(DISTINCT m.id) AS milestones_count,
    COALESCE(pe.engagement_score, 0) AS engagement_score
FROM proposals p
LEFT JOIN cover_letters cl ON p.id = cl.proposal_id
LEFT JOIN attachments a ON p.id = a.proposal_id AND a.is_deleted = FALSE
LEFT JOIN milestones m ON p.id = m.proposal_id
LEFT JOIN proposal_engagement pe ON p.id = pe.proposal_id
WHERE p.is_deleted = FALSE
    AND p.status IN ('SUBMITTED', 'UNDER_REVIEW', 'SHORTLISTED')
GROUP BY p.id, cl.word_count, pe.engagement_score;

-- View for proposal performance metrics
CREATE VIEW v_proposal_performance AS
SELECT
    p.id AS proposal_id,
    p.freelancer_id,
    p.job_id,
    p.status,
    p.submitted_at,
    pp.total_views,
    pp.unique_views,
    pp.engagement_score,
    psm.overall_match_score,
    psm.match_percentage,
    pra.overall_risk_score,
    pra.risk_level,
    EXTRACT(EPOCH FROM (CURRENT_TIMESTAMP - p.submitted_at))/3600 AS hours_since_submission
FROM proposals p
LEFT JOIN proposal_performance pp ON p.id = pp.proposal_id
LEFT JOIN proposal_skill_matches psm ON p.id = psm.proposal_id
LEFT JOIN proposal_risk_assessments pra ON p.id = pra.proposal_id
WHERE p.is_deleted = FALSE;

```
=========================================
##  SECTION 51: DATABASE FUNCTIONS & TRIGGERS
-- =========================================

```sql
-- Function to update proposal updated_at timestamp
CREATE FUNCTION update_proposal_updated_at()
RETURNS TRIGGER AS $
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$ LANGUAGE plpgsql;

-- Trigger for proposals table
CREATE TRIGGER trg_proposals_updated_at
    BEFORE UPDATE ON proposals
    FOR EACH ROW
    EXECUTE FUNCTION update_proposal_updated_at();

-- Function to update connect balance
CREATE FUNCTION update_connect_balance()
RETURNS TRIGGER AS $
BEGIN
    IF TG_OP = 'INSERT' THEN
        UPDATE connects
        SET available_connects = available_connects + NEW.amount,
            total_connects = total_connects + NEW.amount
        WHERE freelancer_id = NEW.freelancer_id;
    END IF;
    RETURN NEW;
END;
$ LANGUAGE plpgsql;

-- Function to track proposal submission
CREATE FUNCTION on_proposal_submitted()
RETURNS TRIGGER AS $
BEGIN
    IF NEW.status = 'SUBMITTED' AND OLD.status = 'DRAFT' THEN
        -- Update freelancer stats
        UPDATE freelancer_proposal_stats
        SET total_proposals = total_proposals + 1,
            submitted_count = submitted_count + 1,
            draft_count = draft_count - 1,
            last_proposal_at = NEW.submitted_at,
            proposals_last_30_days = proposals_last_30_days + 1
        WHERE freelancer_id = NEW.freelancer_id;

        -- Update job stats
        UPDATE job_proposal_stats
        SET total_proposals = total_proposals + 1,
            last_proposal_at = NEW.submitted_at,
            proposals_last_24h = proposals_last_24h + 1
        WHERE job_id = NEW.job_id;
    END IF;
    RETURN NEW;
END;
$ LANGUAGE plpgsql;

CREATE TRIGGER trg_proposal_submitted
    AFTER UPDATE ON proposals
    FOR EACH ROW
    WHEN (NEW.status = 'SUBMITTED' AND OLD.status = 'DRAFT')
    EXECUTE FUNCTION on_proposal_submitted();

```
=========================================
##  SECTION 52: TABLE COMMENTS
-- =========================================
```sql
COMMENT ON TABLE cover_letters IS 'Cover letters - maps to internal/domain/cover_letter/entity.go';
COMMENT ON TABLE attachments IS 'Proposal attachments - maps to internal/domain/attachment/entity.go';
COMMENT ON TABLE question_answers IS 'Question answers - maps to internal/domain/question_answer/entity.go';
COMMENT ON TABLE milestones IS 'Proposal milestones - maps to internal/domain/milestone/entity.go';
COMMENT ON TABLE bids IS 'Bids - maps to internal/domain/bid/entity.go';
COMMENT ON TABLE bid_strategies IS 'Bid strategies - maps to internal/domain/bid_strategy/entity.go';
COMMENT ON TABLE bid_notifications IS 'Bid notifications - maps to internal/domain/bid_notification/entity.go';
COMMENT ON TABLE auctions IS 'Auctions - maps to internal/domain/auction/entity.go';
COMMENT ON TABLE bid_anomalies IS 'Bid anomaly detection - maps to internal/domain/bid_anomaly_detection/entity.go';
COMMENT ON TABLE connects IS 'Connect balances - maps to internal/domain/connect/entity.go';
COMMENT ON TABLE connect_refunds IS 'Connect refunds - maps to internal/domain/connect_refund/entity.go';
COMMENT ON TABLE boosts IS 'Proposal boosts - maps to internal/domain/boost/entity.go';
COMMENT ON TABLE templates IS 'Proposal templates - maps to internal/domain/template/entity.go';
COMMENT ON TABLE rate_cards IS 'Rate cards - maps to internal/domain/rate_card/entity.go';
COMMENT ON TABLE proposal_performance IS 'Performance analytics - maps to internal/domain/performance/entity.go';
COMMENT ON TABLE proposal_similarity IS 'Similarity analysis - maps to internal/domain/similarity/entity.go';
COMMENT ON TABLE proposal_portfolios IS 'Portfolio links - maps to internal/domain/portfolio/entity.go';
COMMENT ON TABLE proposal_engagement IS 'Engagement tracking - maps to internal/domain/engagement/entity.go';
COMMENT ON TABLE spam_detections IS 'Spam detection - maps to internal/domain/spam_detection/entity.go';
COMMENT ON TABLE proposal_flags IS 'Flagging system - maps to internal/domain/flag/entity.go';
COMMENT ON TABLE proposal_compliance IS 'Compliance tracking - maps to internal/domain/compliance/entity.go';
COMMENT ON TABLE interviews IS 'Interviews - maps to internal/domain/interview/entity.go';
COMMENT ON TABLE proposal_feedback IS 'Client feedback - maps to internal/domain/feedback/entity.go';
COMMENT ON TABLE shortlists IS 'Shortlisting - maps to internal/domain/shortlist/entity.go';
COMMENT ON TABLE proposal_conversations IS 'Conversation tracking - maps to internal/domain/conversation/entity.go';
COMMENT ON TABLE negotiations IS 'Negotiations - maps to internal/domain/negotiation/entity.go';
COMMENT ON TABLE invitations IS 'Invitations - maps to internal/domain/invite/entity.go';
COMMENT ON TABLE proposal_revisions IS 'Revision history - maps to internal/domain/revision/entity.go';
COMMENT ON TABLE team_proposals IS 'Team proposals - maps to internal/domain/collaboration/entity.go';
COMMENT ON TABLE proposal_expirations IS 'Expiration tracking - maps to internal/domain/expiration/entity.go';
COMMENT ON TABLE proposal_withdrawals IS 'Withdrawal tracking - maps to internal/domain/withdrawal/entity.go';
COMMENT ON TABLE proposal_archives IS 'Archive tracking - maps to internal/domain/archive/entity.go';
COMMENT ON TABLE proposal_pipelines IS 'Pipeline tracking - maps to internal/domain/pipeline/entity.go';
COMMENT ON TABLE proposal_recycling IS 'Proposal recycling - maps to internal/domain/recycling/entity.go';
COMMENT ON TABLE proposal_recommendations IS 'AI recommendations - maps to internal/domain/recommendation/entity.go';
COMMENT ON TABLE proposal_context IS 'Context enrichment - maps to internal/domain/context/entity.go';
COMMENT ON TABLE proposal_urgency IS 'Urgency tracking - maps to internal/domain/urgency/entity.go';
COMMENT ON TABLE proposal_risk_assessments IS 'Risk assessment - maps to internal/domain/risk_assessment/entity.go';
COMMENT ON TABLE proposal_ai_suggestions IS 'AI suggestions - maps to internal/domain/ai_assist/entity.go';
COMMENT ON TABLE proposal_skill_matches IS 'Skill matching - maps to internal/domain/skill_match/entity.go';
COMMENT ON TABLE video_introductions IS 'Video introductions - maps to internal/domain/video_introduction/entity.go';
COMMENT ON TABLE proposal_references IS 'References - maps to internal/domain/reference/entity.go';

```
=========================================
##  SECTION 53: DATABASE STATISTICS
-- =========================================

```sql
CREATE VIEW v_table_sizes AS
SELECT
    schemaname,
    tablename,
    pg_size_pretty(pg_total_relation_size(schemaname||'.'||tablename)) AS size,
    pg_total_relation_size(schemaname||'.'||tablename) AS size_bytes
FROM pg_tables
WHERE schemaname = 'public'
ORDER BY pg_total_relation_size(schemaname||'.'||tablename) DESC;

CREATE VIEW v_index_usage AS
SELECT
    schemaname,
    tablename,
    indexname,
    idx_scan AS index_scans,
    idx_tup_read AS tuples_read,
    idx_tup_fetch AS tuples_fetched,
    pg_size_pretty(pg_relation_size(indexrelid)) AS index_size
FROM pg_stat_user_indexes
WHERE schemaname = 'public'
ORDER BY idx_scan DESC;
```
=========================================
## END OF PROPOSALS-BE DATABASE DESIGN
=========================================

## FINAL SUMMARY:
- Total Tables: 90+
- Total Indexes: 250+
- Total Domains Covered: 44 (all from proposals-be folder structure)
- Coverage: 100% of proposals-be folder structure
- Production ready for millions of proposals
- Full event sourcing with outbox pattern
- CQRS with read models
- Complete audit trails
- AI/ML integration ready
- Bidding & auction system
- Team collaboration support
- Full compliance tracking
- Risk assessment & fraud detection
- Multi-language support ready
- Enterprise-scale performance optimization

### ALIGNMENT WITH FOLDER STRUCTURE:
- ✅ proposal/ → proposals table
- ✅ cover_letter/ → cover_letters table
- ✅ attachment/ → attachments table
- ✅ question_answer/ → question_answers table
- ✅ milestone/ → milestones table
- ✅ bid/ → bids table
- ✅ bid_strategy/ → bid_strategies table
- ✅ bid_notification/ → bid_notifications table
- ✅ auction/ → auctions table
- ✅ bid_anomaly_detection/ → bid_anomalies table
- ✅ connect/ → connects table
- ✅ connect_refund/ → connect_refunds table
- ✅ boost/ → boosts table
- ✅ template/ → templates table
- ✅ rate_card/ → rate_cards table
- ✅ performance/ → proposal_performance table
- ✅ similarity/ → proposal_similarity table
- ✅ portfolio/ → proposal_portfolios table
- ✅ engagement/ → proposal_engagement table
- ✅ spam_detection/ → spam_detections table
- ✅ flag/ → proposal_flags table
- ✅ compliance/ → proposal_compliance table
- ✅ interview/ → interviews table
- ✅ feedback/ → proposal_feedback table
- ✅ shortlist/ → shortlists table
- ✅ conversation/ → proposal_conversations table
- ✅ negotiation/ → negotiations table
- ✅ invite/ → invitations table
- ✅ revision/ → proposal_revisions table
- ✅ collaboration/ → team_proposals table
- ✅ expiration/ → proposal_expirations table
- ✅ withdrawal/ → proposal_withdrawals table
- ✅ archive/ → proposal_archives table
- ✅ pipeline/ → proposal_pipelines table
- ✅ recycling/ → proposal_recycling table
- ✅ recommendation/ → proposal_recommendations table
- ✅ context/ → proposal_context table
- ✅ urgency/ → proposal_urgency table
- ✅ risk_assessment/ → proposal_risk_assessments table
- ✅ ai_assist/ → proposal_ai_suggestions, proposal_ai_optimizations tables
- ✅ skill_match/ → proposal_skill_matches table
- ✅ video_introduction/ → video_introductions table
- ✅ reference/ → proposal_references table
- ✅ ab_testing/ → proposal_experiments tables
- ✅ outbox/ → outbox_events table

All domains from the proposals-be folder structure are fully covered!