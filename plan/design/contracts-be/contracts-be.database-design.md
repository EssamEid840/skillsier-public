# CONTRACTS-BE DATABASE DESIGN
- Skillsier Platform - Enterprise Scale
- PostgreSQL 16+

## CRITICAL ALIGNMENT RULES:
- 1. Each domain folder in internal/domain/{domain}/ = ONE main table
- 2. Table names follow domain folder names; when aggregated under Contracts, tables are prefixed with contract_ to reflect the domain boundary

- 3. Sub-entities within domain create related tables with {domain}_{sub} naming
- 4. All domains from folder structure are covered
- 5. Rich, production-ready fields for large-scale application

```
-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";
CREATE EXTENSION IF NOT EXISTS "btree_gin";
CREATE EXTENSION IF NOT EXISTS "citext";
CREATE EXTENSION IF NOT EXISTS "btree_gist";

```
=========================================
##  SECTION 1: CORE CONTRACT DOMAIN
```sql
-- Domain: internal/domain/contract/
-- Entity: contract/entity.go
-- =========================================

CREATE TABLE contracts (
    -- Primary Key
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Idempotency
    idempotency_key UUID UNIQUE,
    
    -- References to Other Services
    job_id UUID, -- NULL for direct contracts
    proposal_id UUID, -- NULL for direct contracts
    client_id UUID NOT NULL,
    freelancer_id UUID NOT NULL,
    
    -- Contract Identity
    contract_number VARCHAR(50) UNIQUE NOT NULL, -- e.g., CTR-2025-001234
    title VARCHAR(300) NOT NULL,
    description TEXT,
    
    -- Contract Type
    contract_type VARCHAR(30) NOT NULL CHECK (
        contract_type IN ('FIXED_PRICE', 'HOURLY', 'MILESTONE_BASED', 'RETAINER', 'TIME_AND_MATERIALS', 'RECURRING')
    ),
    
    -- Financial Terms
    total_contract_value DECIMAL(15, 2),
    currency CHAR(3) DEFAULT 'USD' NOT NULL,
    hourly_rate DECIMAL(10, 2), -- for hourly contracts
    payment_frequency VARCHAR(20), -- WEEKLY, BIWEEKLY, MONTHLY, ON_MILESTONE
    
    -- Timeline
    start_date DATE NOT NULL,
    end_date DATE,
    expected_duration_days INTEGER,
    actual_duration_days INTEGER,
    is_indefinite BOOLEAN DEFAULT FALSE,
    
    -- Status & Lifecycle
    status VARCHAR(20) DEFAULT 'DRAFT' CHECK (
        status IN ('DRAFT', 'PENDING_SIGNATURE', 'ACTIVE', 'PAUSED', 'COMPLETED', 
                   'TERMINATED', 'DISPUTED', 'CANCELLED', 'EXPIRED')
    ),
    
    -- Lifecycle Timestamps
    drafted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    signed_at TIMESTAMPTZ,
    activated_at TIMESTAMPTZ,
    paused_at TIMESTAMPTZ,
    resumed_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    terminated_at TIMESTAMPTZ,
    cancelled_at TIMESTAMPTZ,
    
    -- Pause Management
    pause_reason TEXT,
    pause_requested_by UUID,
    max_pause_days INTEGER DEFAULT 30,
    total_paused_days INTEGER DEFAULT 0,
    
    -- Termination
    termination_type VARCHAR(20), -- MUTUAL, UNILATERAL, BREACH, FORCE_MAJEURE
    termination_reason TEXT,
    terminated_by UUID,
    early_termination_penalty DECIMAL(12, 2),
    
    -- Work Arrangement
    work_hours_per_week INTEGER,
    weekly_limit_hours INTEGER, -- For hourly contracts
    requires_nda BOOLEAN DEFAULT FALSE,
    nda_signed_at TIMESTAMPTZ,
    
    -- Deliverables
    expected_deliverables TEXT[],
    acceptance_criteria TEXT,
    
    -- Communication Preferences
    preferred_communication_channel VARCHAR(30),
    meeting_frequency VARCHAR(30),
    reporting_requirements TEXT,
    
    -- Performance Tracking
    performance_rating DECIMAL(3, 2), -- 0.00 to 5.00
    client_satisfaction_score INTEGER CHECK (client_satisfaction_score BETWEEN 1 AND 5),
    freelancer_satisfaction_score INTEGER CHECK (freelancer_satisfaction_score BETWEEN 1 AND 5),
    
    -- Financial Tracking
    total_paid DECIMAL(15, 2) DEFAULT 0,
    total_pending DECIMAL(15, 2) DEFAULT 0,
    total_hours_worked DECIMAL(10, 2) DEFAULT 0,
    total_hours_approved DECIMAL(10, 2) DEFAULT 0,
    
    -- Budget Management
    budget_consumed_percentage DECIMAL(5, 2) DEFAULT 0,
    is_over_budget BOOLEAN DEFAULT FALSE,
    
    -- Compliance & Legal
    requires_background_check BOOLEAN DEFAULT FALSE,
    background_check_completed BOOLEAN DEFAULT FALSE,
    background_check_completed_at TIMESTAMPTZ,
    
    ip_ownership VARCHAR(20) DEFAULT 'CLIENT' CHECK (
        ip_ownership IN ('CLIENT', 'FREELANCER', 'SHARED', 'NEGOTIATED')
    ),
    confidentiality_level VARCHAR(20) DEFAULT 'STANDARD',
    
    -- Escrow Integration
    escrow_account_id UUID, -- Reference to financial-be
    escrow_status VARCHAR(20),
    funds_held DECIMAL(15, 2) DEFAULT 0,
    
    -- Auto-Renewal (for recurring contracts)
    auto_renew BOOLEAN DEFAULT FALSE,
    auto_renew_notification_days INTEGER DEFAULT 30,
    renewal_count INTEGER DEFAULT 0,
    
    -- Quality & Risk
    quality_score DECIMAL(5, 2),
    risk_score DECIMAL(5, 2),
    compliance_score DECIMAL(5, 2),
    
    -- Template
    created_from_template_id UUID,
    
    -- Collaboration
    allows_subcontracting BOOLEAN DEFAULT FALSE,
    max_subcontractors INTEGER,
    
    -- Notifications
    notify_client BOOLEAN DEFAULT TRUE,
    notify_freelancer BOOLEAN DEFAULT TRUE,
    
    -- Metadata
    tags TEXT[],
    custom_fields JSONB,
    notes TEXT,
    
    -- Audit
    created_by UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    version INTEGER DEFAULT 1 NOT NULL,
    
    -- Soft Delete
    is_deleted BOOLEAN DEFAULT FALSE,
    deleted_at TIMESTAMPTZ,
    deleted_by UUID,
    
    -- Constraints
    CONSTRAINT chk_contracts_dates CHECK (end_date IS NULL OR end_date >= start_date),
    CONSTRAINT chk_contracts_value CHECK (total_contract_value IS NULL OR total_contract_value >= 0),
    CONSTRAINT chk_contracts_hourly CHECK (hourly_rate IS NULL OR hourly_rate >= 0)
);

-- Indexes
CREATE INDEX idx_contracts_client ON contracts (client_id, status) WHERE is_deleted = FALSE;
CREATE INDEX idx_contracts_freelancer ON contracts (freelancer_id, status) WHERE is_deleted = FALSE;
CREATE INDEX idx_contracts_status ON contracts (status, activated_at DESC) WHERE is_deleted = FALSE;
CREATE INDEX idx_contracts_job ON contracts (job_id) WHERE job_id IS NOT NULL;
CREATE INDEX idx_contracts_proposal ON contracts (proposal_id) WHERE proposal_id IS NOT NULL;
CREATE INDEX idx_contracts_number ON contracts (contract_number);
CREATE INDEX idx_contracts_dates ON contracts (start_date, end_date) WHERE is_deleted = FALSE;
CREATE INDEX idx_contracts_type ON contracts (contract_type, status);
CREATE INDEX idx_contracts_escrow ON contracts (escrow_account_id) WHERE escrow_account_id IS NOT NULL;
CREATE INDEX idx_contracts_template ON contracts (created_from_template_id) WHERE created_from_template_id IS NOT NULL;

COMMENT ON TABLE contracts IS 'Core contracts - maps to internal/domain/contract/entity.go';

-- Contract State History
CREATE TABLE contract_state_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- State Transition
    from_status VARCHAR(20) NOT NULL,
    to_status VARCHAR(20) NOT NULL,
    
    -- Context
    reason TEXT,
    changed_by UUID NOT NULL,
    changed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Additional Data
    metadata JSONB,
    
    CONSTRAINT fk_contract_state_history_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_contract_state_history_contract ON contract_state_history (contract_id, changed_at DESC);

```
=========================================
##  SECTION 2: STATEMENT OF WORK (SOW)
```sql
-- Domain: internal/domain/sow/
-- Entity: sow/entity.go
-- =========================================

CREATE TABLE statements_of_work (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- SOW Identity
    sow_number VARCHAR(50) UNIQUE NOT NULL,
    title VARCHAR(300) NOT NULL,
    version INTEGER DEFAULT 1 NOT NULL,
    
    -- Content
    scope_of_work TEXT NOT NULL,
    objectives TEXT[],
    deliverables JSONB, -- Structured deliverables list
    timeline_details TEXT,
    
    -- Responsibilities
    client_responsibilities TEXT,
    freelancer_responsibilities TEXT,
    
    -- Acceptance Criteria
    acceptance_criteria TEXT,
    testing_requirements TEXT,
    quality_standards TEXT,
    
    -- Dependencies & Assumptions
    dependencies TEXT[],
    assumptions TEXT[],
    constraints TEXT[],
    risks TEXT[],
    
    -- Status
    status VARCHAR(20) DEFAULT 'DRAFT' CHECK (
        status IN ('DRAFT', 'PENDING_APPROVAL', 'APPROVED', 'REJECTED', 'SUPERSEDED', 'ARCHIVED')
    ),
    
    -- Approval Workflow
    requires_client_approval BOOLEAN DEFAULT TRUE,
    requires_freelancer_approval BOOLEAN DEFAULT TRUE,
    client_approved BOOLEAN DEFAULT FALSE,
    client_approved_by UUID,
    client_approved_at TIMESTAMPTZ,
    freelancer_approved BOOLEAN DEFAULT FALSE,
    freelancer_approved_by UUID,
    freelancer_approved_at TIMESTAMPTZ,
    
    -- Rejection
    rejection_reason TEXT,
    rejected_by UUID,
    rejected_at TIMESTAMPTZ,
    
    -- Versioning
    superseded_by UUID, -- Points to newer SOW version
    effective_from DATE,
    effective_until DATE,
    
    -- Documents
    document_urls TEXT[],
    
    created_by UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_statements_of_work_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE,
    CONSTRAINT fk_statements_of_work_superseded FOREIGN KEY (superseded_by) 
        REFERENCES statements_of_work(id)
);

CREATE INDEX idx_statements_of_work_contract ON statements_of_work (contract_id, version DESC);
CREATE INDEX idx_statements_of_work_status ON statements_of_work (status);
CREATE INDEX idx_statements_of_work_number ON statements_of_work (sow_number);

COMMENT ON TABLE statements_of_work IS 'Statements of Work - maps to internal/domain/sow/entity.go';

```
=========================================
##  SECTION 3: FINANCIAL HOLDS
```sql
-- Domain: internal/domain/hold/
-- Entity: hold/entity.go
-- =========================================

CREATE TABLE financial_holds (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- Hold Details
    hold_type VARCHAR(30) NOT NULL CHECK (
        hold_type IN ('RISK_ASSESSMENT', 'DISPUTE', 'FRAUD_INVESTIGATION', 'COMPLIANCE', 
                      'PAYMENT_VERIFICATION', 'CHARGEBACK', 'MANUAL_REVIEW')
    ),
    
    -- Amount
    hold_amount DECIMAL(15, 2) NOT NULL,
    currency CHAR(3) DEFAULT 'USD',
    
    -- Reason
    reason TEXT NOT NULL,
    risk_level VARCHAR(20) CHECK (
        risk_level IN ('LOW', 'MEDIUM', 'HIGH', 'CRITICAL')
    ),
    
    -- Status
    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (
        status IN ('ACTIVE', 'RELEASED', 'ESCALATED', 'CANCELLED')
    ),
    
    -- Timeline
    placed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at TIMESTAMPTZ,
    released_at TIMESTAMPTZ,
    
    -- Resolution
    resolution_type VARCHAR(30),
    resolution_notes TEXT,
    resolved_by UUID,
    
    -- Impact
    affects_payout BOOLEAN DEFAULT TRUE,
    blocks_new_work BOOLEAN DEFAULT FALSE,
    
    -- Reference
    reference_type VARCHAR(30), -- MILESTONE, TIMESHEET, TRANSACTION
    reference_id UUID,
    
    -- Audit
    placed_by UUID NOT NULL,
    
    CONSTRAINT fk_financial_holds_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE,
    CONSTRAINT chk_financial_holds_amount CHECK (hold_amount >= 0)
);

CREATE INDEX idx_financial_holds_contract ON financial_holds (contract_id, status);
CREATE INDEX idx_financial_holds_status ON financial_holds (status, placed_at DESC);
CREATE INDEX idx_financial_holds_expires ON financial_holds (expires_at) WHERE status = 'ACTIVE';
CREATE INDEX idx_financial_holds_reference ON financial_holds (reference_type, reference_id);

COMMENT ON TABLE financial_holds IS 'Financial holds - maps to internal/domain/hold/entity.go';

```
=========================================
##  SECTION 4: MILESTONES
```sql
-- Domain: internal/domain/milestone/
-- Entity: milestone/entity.go
-- =========================================

CREATE TABLE milestones (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- Milestone Identity
    milestone_number INTEGER NOT NULL,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    
    -- Financial
    amount DECIMAL(12, 2) NOT NULL,
    currency CHAR(3) DEFAULT 'USD',
    
    -- Timeline
    due_date DATE NOT NULL,
    estimated_hours DECIMAL(8, 2),
    
    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'IN_PROGRESS', 'SUBMITTED', 'UNDER_REVIEW', 'REVISION_REQUESTED',
                   'APPROVED', 'REJECTED', 'PAID', 'CANCELLED', 'DISPUTED')
    ),
    
    -- Deliverables
    expected_deliverables TEXT[],
    deliverable_format VARCHAR(100),
    acceptance_criteria TEXT,
    
    -- Submission
    submitted_at TIMESTAMPTZ,
    submitted_by UUID,
    submission_notes TEXT,
    
    -- Review
    reviewed_at TIMESTAMPTZ,
    reviewed_by UUID,
    review_notes TEXT,
    review_rating INTEGER CHECK (review_rating BETWEEN 1 AND 5),
    
    -- Approval
    approved_at TIMESTAMPTZ,
    approved_by UUID,
    approval_notes TEXT,
    
    -- Rejection/Revision
    rejection_reason TEXT,
    rejected_at TIMESTAMPTZ,
    revision_count INTEGER DEFAULT 0,
    max_revisions INTEGER DEFAULT 3,
    
    -- Payment
    payment_released_at TIMESTAMPTZ,
    payment_transaction_id UUID, -- Reference to financial-be
    
    -- Late Penalties
    is_late BOOLEAN DEFAULT FALSE,
    days_late INTEGER DEFAULT 0,
    late_penalty_amount DECIMAL(12, 2),
    
    -- Dependencies
    depends_on_milestone_ids UUID[],
    blocks_milestone_ids UUID[],
    
    -- Quality Metrics
    quality_score DECIMAL(5, 2),
    completion_percentage INTEGER DEFAULT 0 CHECK (completion_percentage BETWEEN 0 AND 100),
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_milestones_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE,
    CONSTRAINT uk_milestones_number UNIQUE (contract_id, milestone_number),
    CONSTRAINT chk_milestones_amount CHECK (amount >= 0)
);

CREATE INDEX idx_milestones_contract ON milestones (contract_id, milestone_number);
CREATE INDEX idx_milestones_status ON milestones (status, due_date);
CREATE INDEX idx_milestones_due ON milestones (due_date) WHERE status IN ('PENDING', 'IN_PROGRESS');
CREATE INDEX idx_milestones_payment ON milestones (payment_transaction_id) WHERE payment_transaction_id IS NOT NULL;

COMMENT ON TABLE milestones IS 'Contract milestones - maps to internal/domain/milestone/entity.go';

-- Milestone Activities Log
CREATE TABLE milestone_activities (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    milestone_id UUID NOT NULL,
    
    -- Activity
    activity_type VARCHAR(50) NOT NULL,
    description TEXT,
    
    -- Actor
    performed_by UUID NOT NULL,
    performed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Changes
    old_values JSONB,
    new_values JSONB,
    
    CONSTRAINT fk_milestone_activities_milestone FOREIGN KEY (milestone_id) 
        REFERENCES milestones(id) ON DELETE CASCADE
);

CREATE INDEX idx_milestone_activities_milestone ON milestone_activities (milestone_id, performed_at DESC);

```
=========================================
##  SECTION 5: DELIVERABLES
```sql
-- Domain: internal/domain/deliverable/
-- Entity: deliverable/entity.go
-- =========================================

CREATE TABLE deliverables (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    milestone_id UUID NOT NULL,
    
    -- Deliverable Details
    title VARCHAR(200) NOT NULL,
    description TEXT,
    deliverable_type VARCHAR(50) CHECK (
        deliverable_type IN ('DOCUMENT', 'CODE', 'DESIGN', 'VIDEO', 'REPORT', 'OTHER')
    ),
    
    -- File References (storage-be)
    file_id UUID,
    file_url TEXT,
    file_name VARCHAR(255),
    file_size BIGINT,
    file_mime_type VARCHAR(100),
    
    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'SUBMITTED', 'UNDER_REVIEW', 'REVISION_REQUESTED', 
                   'APPROVED', 'REJECTED')
    ),
    
    -- Submission
    submitted_at TIMESTAMPTZ,
    submitted_by UUID,
    submission_notes TEXT,
    
    -- Review
    reviewed_at TIMESTAMPTZ,
    reviewed_by UUID,
    review_notes TEXT,
    review_rating INTEGER CHECK (review_rating BETWEEN 1 AND 5),
    
    -- Revision
    revision_number INTEGER DEFAULT 0,
    revision_notes TEXT,
    previous_version_id UUID, -- Points to previous deliverable version
    
    -- Approval
    approved_at TIMESTAMPTZ,
    approved_by UUID,
    
    -- Quality Metrics
    quality_score DECIMAL(5, 2),
    meets_requirements BOOLEAN,
    
    -- Virus Scanning
    virus_scan_status VARCHAR(20) DEFAULT 'PENDING',
    scanned_at TIMESTAMPTZ,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_deliverables_milestone FOREIGN KEY (milestone_id) 
        REFERENCES milestones(id) ON DELETE CASCADE,
    CONSTRAINT fk_deliverables_previous FOREIGN KEY (previous_version_id) 
        REFERENCES deliverables(id)
);

CREATE INDEX idx_deliverables_milestone ON deliverables (milestone_id, revision_number DESC);
CREATE INDEX idx_deliverables_status ON deliverables (status);
CREATE INDEX idx_deliverables_file ON deliverables (file_id) WHERE file_id IS NOT NULL;

COMMENT ON TABLE deliverables IS 'Milestone deliverables - maps to internal/domain/deliverable/entity.go';

```
=========================================
##  SECTION 6: TIME TRACKING - TIMESHEETS
```sql
-- Domain: internal/domain/timesheet/
-- Entity: timesheet/entity.go
-- =========================================

CREATE TABLE timesheets (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- Timesheet Period
    week_start_date DATE NOT NULL,
    week_end_date DATE NOT NULL,
    
    -- Hours
    total_hours DECIMAL(8, 2) NOT NULL DEFAULT 0,
    billable_hours DECIMAL(8, 2) NOT NULL DEFAULT 0,
    non_billable_hours DECIMAL(8, 2) DEFAULT 0,
    overtime_hours DECIMAL(8, 2) DEFAULT 0,
    
    -- Financial
    hourly_rate DECIMAL(10, 2) NOT NULL,
    total_amount DECIMAL(12, 2) NOT NULL,
    currency CHAR(3) DEFAULT 'USD',
    
    -- Status
    status VARCHAR(20) DEFAULT 'DRAFT' CHECK (
        status IN ('DRAFT', 'SUBMITTED', 'UNDER_REVIEW', 'APPROVED', 'REJECTED', 
                   'PAID', 'DISPUTED')
    ),
    
    -- Submission
    submitted_at TIMESTAMPTZ,
    submitted_by UUID NOT NULL,
    submission_notes TEXT,
    
    -- Review
    reviewed_at TIMESTAMPTZ,
    reviewed_by UUID,
    review_notes TEXT,
    
    -- Approval
    approved_at TIMESTAMPTZ,
    approved_by UUID,
    approved_hours DECIMAL(8, 2), -- May differ from submitted hours
    approved_amount DECIMAL(12, 2),
    
    -- Rejection
    rejection_reason TEXT,
    rejected_at TIMESTAMPTZ,
    
    -- Payment
    payment_released_at TIMESTAMPTZ,
    payment_transaction_id UUID,
    
    -- Validation
    has_work_diary BOOLEAN DEFAULT FALSE,
    work_diary_coverage_percentage DECIMAL(5, 2),
    manual_time_entries_count INTEGER DEFAULT 0,
    
    -- Dispute
    is_disputed BOOLEAN DEFAULT FALSE,
    dispute_reason TEXT,
    disputed_at TIMESTAMPTZ,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_timesheets_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE,
    CONSTRAINT uk_timesheets_period UNIQUE (contract_id, week_start_date),
    CONSTRAINT chk_timesheets_hours CHECK (total_hours >= 0),
    CONSTRAINT chk_timesheets_dates CHECK (week_end_date > week_start_date)
);

CREATE INDEX idx_timesheets_contract ON timesheets (contract_id, week_start_date DESC);
CREATE INDEX idx_timesheets_status ON timesheets (status, submitted_at DESC);
CREATE INDEX idx_timesheets_pending_approval ON timesheets (status, submitted_at) 
    WHERE status = 'SUBMITTED';

COMMENT ON TABLE timesheets IS 'Weekly timesheets - maps to internal/domain/timesheet/entity.go';

-- Time Entries (Daily breakdown)
CREATE TABLE time_entries (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    timesheet_id UUID NOT NULL,
    
    -- Entry Details
    work_date DATE NOT NULL,
    start_time TIME,
    end_time TIME,
    hours_worked DECIMAL(5, 2) NOT NULL,
    
    -- Description
    task_description TEXT NOT NULL,
    task_category VARCHAR(100),
    
    -- Classification
    is_billable BOOLEAN DEFAULT TRUE,
    is_overtime BOOLEAN DEFAULT FALSE,
    is_manual_entry BOOLEAN DEFAULT FALSE, -- vs auto-tracked
    
    -- Work Diary Reference
    has_work_diary_entry BOOLEAN DEFAULT FALSE,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_time_entries_timesheet FOREIGN KEY (timesheet_id) 
        REFERENCES timesheets(id) ON DELETE CASCADE,
    CONSTRAINT chk_time_entries_hours CHECK (hours_worked >= 0 AND hours_worked <= 24)
);

CREATE INDEX idx_time_entries_timesheet ON time_entries (timesheet_id, work_date);

```
=========================================
##  SECTION 7: WORK DIARY
```sql
-- Domain: internal/domain/work_diary/
-- Entity: work_diary/entity.go
-- =========================================

CREATE TABLE work_diary_entries (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    time_entry_id UUID, -- Optional link to time_entry
    
    -- Entry Details
    work_date DATE NOT NULL,
    time_slot TIME NOT NULL, -- 10-minute slots
    
    -- Activity Tracking
    activity_level INTEGER CHECK (activity_level BETWEEN 0 AND 100),
    keyboard_strokes INTEGER DEFAULT 0,
    mouse_clicks INTEGER DEFAULT 0,
    mouse_movement INTEGER DEFAULT 0,
    
    -- Screenshots
    screenshot_count INTEGER DEFAULT 0,
    screenshot_urls TEXT[],
    blurred_screenshot_urls TEXT[],
    
    -- Application Tracking
    active_application VARCHAR(200),
    active_window_title VARCHAR(500),
    url_visited TEXT,
    
    -- Work Description
    task_description TEXT,
    
    -- Status
    is_productive BOOLEAN DEFAULT TRUE,
    is_manual BOOLEAN DEFAULT FALSE,
    is_verified BOOLEAN DEFAULT FALSE,
    
    -- Privacy
    screenshot_deleted BOOLEAN DEFAULT FALSE,
    screenshot_deleted_at TIMESTAMPTZ,
    screenshot_deleted_reason TEXT,
    blur_applied BOOLEAN DEFAULT FALSE,
    
    -- Metadata
    device_info JSONB,
    timezone VARCHAR(50),
    
    recorded_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_work_diary_entries_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE,
    CONSTRAINT fk_work_diary_entries_time_entry FOREIGN KEY (time_entry_id) 
        REFERENCES time_entries(id) ON DELETE SET NULL
);

CREATE INDEX idx_work_diary_entries_contract ON work_diary_entries (contract_id, work_date, time_slot);
CREATE INDEX idx_work_diary_entries_date ON work_diary_entries (work_date, time_slot);
CREATE INDEX idx_work_diary_entries_time_entry ON work_diary_entries (time_entry_id) 
    WHERE time_entry_id IS NOT NULL;

COMMENT ON TABLE work_diary_entries IS 'Work diary tracking - maps to internal/domain/work_diary/entity.go';

-- Work Diary Activity Summary (daily aggregate)
CREATE TABLE work_diary_daily_summaries (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    work_date DATE NOT NULL,
    
    -- Aggregated Metrics
    total_tracked_minutes INTEGER DEFAULT 0,
    total_productive_minutes INTEGER DEFAULT 0,
    total_idle_minutes INTEGER DEFAULT 0,
    avg_activity_level DECIMAL(5, 2),
    
    -- Activity Counts
    total_keyboard_strokes INTEGER DEFAULT 0,
    total_mouse_clicks INTEGER DEFAULT 0,
    total_screenshots INTEGER DEFAULT 0,
    
    -- Applications
    top_applications JSONB,
    application_usage_breakdown JSONB,
    
    computed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_work_diary_daily_summaries_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE,
    CONSTRAINT uk_work_diary_daily_summaries UNIQUE (contract_id, work_date)
);

CREATE INDEX idx_work_diary_daily_summaries_contract ON work_diary_daily_summaries (contract_id, work_date DESC);

```
=========================================
##  SECTION 8: CONTRACT TEMPLATES
```sql
-- Domain: internal/domain/template/
-- Entity: template/entity.go
-- =========================================

CREATE TABLE contract_templates (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Template Identity
    template_name VARCHAR(200) NOT NULL,
    template_slug VARCHAR(250) UNIQUE NOT NULL,
    template_category VARCHAR(50) CHECK (
        template_category IN ('STANDARD', 'NDA', 'HOURLY', 'FIXED_PRICE', 'RETAINER', 'CUSTOM')
    ),
    
    -- Content
    title_template VARCHAR(300),
    description_template TEXT,
    
    -- Default Terms
    default_contract_type VARCHAR(30),
    default_payment_frequency VARCHAR(20),
    default_payment_terms TEXT,
    
    -- Clauses
    standard_clauses JSONB, -- Array of reusable clauses
    required_clauses TEXT[],
    optional_clauses TEXT[],
    
    -- Legal Terms
    ip_rights_clause TEXT,
    confidentiality_clause TEXT,
    termination_clause TEXT,
    dispute_resolution_clause TEXT,
    liability_clause TEXT,
    
    -- Configuration
    allows_customization BOOLEAN DEFAULT TRUE,
    requires_legal_review BOOLEAN DEFAULT FALSE,
    
    -- Visibility
    is_public BOOLEAN DEFAULT FALSE,
    is_system_template BOOLEAN DEFAULT FALSE,
    owner_id UUID, -- NULL for system templates
    
    -- Usage Statistics
    usage_count INTEGER DEFAULT 0,
    last_used_at TIMESTAMPTZ,
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    is_archived BOOLEAN DEFAULT FALSE,
    archived_at TIMESTAMPTZ,
    
    -- Versioning
    version INTEGER DEFAULT 1,
    
    created_by UUID,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT chk_contract_templates_owner CHECK (
        (is_system_template = TRUE AND owner_id IS NULL) OR 
        (is_system_template = FALSE AND owner_id IS NOT NULL)
    )
);

CREATE INDEX idx_contract_templates_slug ON contract_templates (template_slug);
CREATE INDEX idx_contract_templates_category ON contract_templates (template_category, is_active);
CREATE INDEX idx_contract_templates_owner ON contract_templates (owner_id) WHERE owner_id IS NOT NULL;
CREATE INDEX idx_contract_templates_public ON contract_templates (is_public, is_active) WHERE is_public = TRUE;

COMMENT ON TABLE contract_templates IS 'Contract templates - maps to internal/domain/template/entity.go';

```
=========================================
##  SECTION 9: CONTRACT AMENDMENTS
```sql
-- Domain: internal/domain/amendment/
-- Entity: amendment/entity.go
-- =========================================

CREATE TABLE contract_amendments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- Amendment Identity
    amendment_number INTEGER NOT NULL,
    amendment_type VARCHAR(30) CHECK (
        amendment_type IN ('SCOPE_CHANGE', 'BUDGET_INCREASE', 'TIMELINE_EXTENSION', 
                          'RATE_ADJUSTMENT', 'DELIVERABLE_CHANGE', 'TERMS_MODIFICATION', 'OTHER')
    ),
    
    -- Description
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    justification TEXT,
    
    -- Changes
    changes JSONB NOT NULL, -- Structured change details
    fields_changed TEXT[],
    
    -- Financial Impact
    budget_impact DECIMAL(12, 2),
    currency CHAR(3) DEFAULT 'USD',
    
    -- Timeline Impact
    timeline_extension_days INTEGER,
    new_end_date DATE,
    
    -- Status
    status VARCHAR(20) DEFAULT 'DRAFT' CHECK (
        status IN ('DRAFT', 'PENDING_APPROVAL', 'APPROVED', 'REJECTED', 'IMPLEMENTED', 'CANCELLED')
    ),
    
    -- Approval Workflow
    requires_client_approval BOOLEAN DEFAULT TRUE,
    requires_freelancer_approval BOOLEAN DEFAULT TRUE,
    client_approved BOOLEAN DEFAULT FALSE,
    client_approved_by UUID,
    client_approved_at TIMESTAMPTZ,
    freelancer_approved BOOLEAN DEFAULT FALSE,
    freelancer_approved_by UUID,
    freelancer_approved_at TIMESTAMPTZ,
    
    -- Implementation
    implemented_at TIMESTAMPTZ,
    implemented_by UUID,
    
    -- Rejection
    rejection_reason TEXT,
    rejected_by UUID,
    rejected_at TIMESTAMPTZ,
    
    -- Documents
    supporting_documents TEXT[],
    
    proposed_by UUID NOT NULL,
    proposed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_contract_amendments_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE,
    CONSTRAINT uk_contract_amendments_number UNIQUE (contract_id, amendment_number)
);

CREATE INDEX idx_contract_amendments_contract ON contract_amendments (contract_id, amendment_number DESC);
CREATE INDEX idx_contract_amendments_status ON contract_amendments (status);
CREATE INDEX idx_contract_amendments_pending ON contract_amendments (status, proposed_at) 
    WHERE status = 'PENDING_APPROVAL';

COMMENT ON TABLE contract_amendments IS 'Contract amendments - maps to internal/domain/amendment/entity.go';

```
=========================================
##  SECTION 10: DISPUTES
```sql
-- Domain: internal/domain/dispute/
-- Entity: dispute/entity.go
-- =========================================

CREATE TABLE disputes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- Dispute Identity
    dispute_number VARCHAR(50) UNIQUE NOT NULL,
    
    -- Type & Subject
    dispute_type VARCHAR(30) NOT NULL CHECK (
        dispute_type IN ('MILESTONE', 'TIMESHEET', 'DELIVERABLE', 'PAYMENT', 'QUALITY', 
                        'SCOPE', 'CONDUCT', 'OTHER')
    ),
    subject_type VARCHAR(30), -- MILESTONE, TIMESHEET, etc.
    subject_id UUID,
    
    -- Description
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    reason TEXT NOT NULL,
    
    -- Parties
    opened_by UUID NOT NULL,
    respondent_id UUID NOT NULL,
    
    -- Status
    status VARCHAR(20) DEFAULT 'OPEN' CHECK (
        status IN ('OPEN', 'UNDER_REVIEW', 'AWAITING_RESPONSE', 'MEDIATION', 'ARBITRATION',
                   'RESOLVED', 'CLOSED', 'ESCALATED')
    ),
    
    -- Priority
    priority VARCHAR(20) DEFAULT 'MEDIUM' CHECK (
        priority IN ('LOW', 'MEDIUM', 'HIGH', 'URGENT')
    ),
    severity VARCHAR(20) CHECK (
        severity IN ('MINOR', 'MODERATE', 'MAJOR', 'CRITICAL')
    ),
    
    -- Financial Impact
    disputed_amount DECIMAL(12, 2),
    currency CHAR(3) DEFAULT 'USD',
    financial_hold_id UUID,
    
    -- Timeline
    opened_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    response_due_at TIMESTAMPTZ,
    responded_at TIMESTAMPTZ,
    resolution_deadline TIMESTAMPTZ,
    resolved_at TIMESTAMPTZ,
    closed_at TIMESTAMPTZ,
    
    -- Mediation
    mediator_assigned UUID,
    mediation_started_at TIMESTAMPTZ,
    
    -- Resolution
    resolution_type VARCHAR(30) CHECK (
        resolution_type IN ('MUTUAL_AGREEMENT', 'MEDIATION', 'ARBITRATION', 'ADMIN_DECISION', 
                           'WITHDRAWN', 'TIMEOUT')
    ),
    resolution_description TEXT,
    resolution_decision TEXT,
    
    -- Financial Resolution
    refund_amount DECIMAL(12, 2),
    refund_to VARCHAR(20), -- CLIENT, FREELANCER
    payment_adjustment DECIMAL(12, 2),
    
    -- Escalation
    escalation_count INTEGER DEFAULT 0,
    last_escalated_at TIMESTAMPTZ,
    escalation_reason TEXT,
    
    -- Satisfaction
    opener_satisfaction_score INTEGER CHECK (opener_satisfaction_score BETWEEN 1 AND 5),
    respondent_satisfaction_score INTEGER CHECK (respondent_satisfaction_score BETWEEN 1 AND 5),
    
    resolved_by UUID,
    
    CONSTRAINT fk_disputes_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_disputes_contract ON disputes (contract_id, opened_at DESC);
CREATE INDEX idx_disputes_status ON disputes (status, priority DESC, opened_at DESC);
CREATE INDEX idx_disputes_number ON disputes (dispute_number);
CREATE INDEX idx_disputes_subject ON disputes (subject_type, subject_id) 
    WHERE subject_type IS NOT NULL AND subject_id IS NOT NULL;
CREATE INDEX idx_disputes_open ON disputes (status, opened_at) WHERE status IN ('OPEN', 'UNDER_REVIEW');

COMMENT ON TABLE disputes IS 'Contract disputes - maps to internal/domain/dispute/entity.go';

-- Dispute Evidence
CREATE TABLE dispute_evidence (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    dispute_id UUID NOT NULL,
    
    -- Evidence Details
    evidence_type VARCHAR(30) CHECK (
        evidence_type IN ('DOCUMENT', 'SCREENSHOT', 'MESSAGE_THREAD', 'VIDEO', 'AUDIO', 'OTHER')
    ),
    title VARCHAR(200) NOT NULL,
    description TEXT,
    
    -- File Reference
    file_id UUID,
    file_url TEXT,
    file_name VARCHAR(255),
    file_size BIGINT,
    
    -- Context
    relevance_score DECIMAL(5, 2),
    
    uploaded_by UUID NOT NULL,
    uploaded_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_dispute_evidence_dispute FOREIGN KEY (dispute_id) 
        REFERENCES disputes(id) ON DELETE CASCADE
);

CREATE INDEX idx_dispute_evidence_dispute ON dispute_evidence (dispute_id, uploaded_at DESC);

-- Dispute Messages/Responses
CREATE TABLE dispute_messages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    dispute_id UUID NOT NULL,
    
    -- Message Details
    message_type VARCHAR(30) CHECK (
        message_type IN ('RESPONSE', 'CLARIFICATION', 'COUNTER_ARGUMENT', 'MEDIATOR_NOTE', 
                        'RESOLUTION_PROPOSAL', 'ADMIN_COMMENT')
    ),
    message_text TEXT NOT NULL,
    
    -- Attachments
    attachment_urls TEXT[],
    
    -- Sender
    sent_by UUID NOT NULL,
    sender_role VARCHAR(20), -- OPENER, RESPONDENT, MEDIATOR, ADMIN
    
    -- Visibility
    is_private BOOLEAN DEFAULT FALSE, -- Private to mediators/admins only
    
    sent_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_dispute_messages_dispute FOREIGN KEY (dispute_id) 
        REFERENCES disputes(id) ON DELETE CASCADE
);

CREATE INDEX idx_dispute_messages_dispute ON dispute_messages (dispute_id, sent_at DESC);

```
=========================================
##  SECTION 11: BUDGET TRACKING
```sql
-- Domain: internal/domain/budget/
-- Entity: budget/entity.go
-- =========================================

CREATE TABLE contract_budgets (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL UNIQUE,
    
    -- Budget Configuration
    total_budget DECIMAL(15, 2) NOT NULL,
    currency CHAR(3) DEFAULT 'USD',
    
    -- Thresholds
    warning_threshold_percentage DECIMAL(5, 2) DEFAULT 80.00,
    critical_threshold_percentage DECIMAL(5, 2) DEFAULT 95.00,
    
    -- Spending Tracking
    committed_amount DECIMAL(15, 2) DEFAULT 0,
    spent_amount DECIMAL(15, 2) DEFAULT 0,
    pending_amount DECIMAL(15, 2) DEFAULT 0,
    remaining_amount DECIMAL(15, 2),
    
    -- Percentage Calculations
    budget_consumed_percentage DECIMAL(5, 2) DEFAULT 0,
    
    -- Status
    is_over_budget BOOLEAN DEFAULT FALSE,
    over_budget_amount DECIMAL(15, 2),
    
    -- Alerts
    warning_triggered BOOLEAN DEFAULT FALSE,
    warning_triggered_at TIMESTAMPTZ,
    critical_triggered BOOLEAN DEFAULT FALSE,
    critical_triggered_at TIMESTAMPTZ,
    
    -- Variance Analysis
    budget_variance DECIMAL(15, 2),
    variance_percentage DECIMAL(5, 2),
    
    -- Forecast
    projected_total_cost DECIMAL(15, 2),
    projected_completion_date DATE,
    burn_rate_per_week DECIMAL(12, 2),
    
    -- Budget Adjustments
    adjustment_count INTEGER DEFAULT 0,
    last_adjustment_at TIMESTAMPTZ,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_contract_budgets_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE,
    CONSTRAINT chk_contract_budgets_total CHECK (total_budget > 0)
);

CREATE INDEX idx_contract_budgets_contract ON contract_budgets (contract_id);
CREATE INDEX idx_contract_budgets_overbudget ON contract_budgets (is_over_budget) 
    WHERE is_over_budget = TRUE;

COMMENT ON TABLE contract_budgets IS 'Budget tracking - maps to internal/domain/budget/entity.go';

-- Budget History/Adjustments
CREATE TABLE budget_adjustments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    budget_id UUID NOT NULL,
    
    -- Adjustment Details
    adjustment_type VARCHAR(30) CHECK (
        adjustment_type IN ('INCREASE', 'DECREASE', 'REALLOCATION', 'CORRECTION')
    ),
    
    -- Amounts
    previous_amount DECIMAL(15, 2) NOT NULL,
    new_amount DECIMAL(15, 2) NOT NULL,
    adjustment_amount DECIMAL(15, 2) NOT NULL,
    
    -- Reason
    reason TEXT NOT NULL,
    justification TEXT,
    
    -- Approval
    requires_approval BOOLEAN DEFAULT TRUE,
    approved BOOLEAN DEFAULT FALSE,
    approved_by UUID,
    approved_at TIMESTAMPTZ,
    
    adjusted_by UUID NOT NULL,
    adjusted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_budget_adjustments_budget FOREIGN KEY (budget_id) 
        REFERENCES contract_budgets(id) ON DELETE CASCADE
);

CREATE INDEX idx_budget_adjustments_budget ON budget_adjustments (budget_id, adjusted_at DESC);

```
=========================================
##  SECTION 12: REMINDERS & NOTIFICATIONS
```sql
-- Domain: internal/domain/reminder/
-- Entity: reminder/entity.go
-- =========================================

CREATE TABLE contract_reminders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- Reminder Type
    reminder_type VARCHAR(30) NOT NULL CHECK (
        reminder_type IN ('MILESTONE_DUE', 'TIMESHEET_SUBMISSION', 'APPROVAL_PENDING', 
                         'PAYMENT_DUE', 'CONTRACT_ENDING', 'REVIEW_REQUEST', 'DOCUMENT_EXPIRY')
    ),
    
    -- Target
    recipient_ids UUID[] NOT NULL,
    
    -- Trigger Conditions
    trigger_condition JSONB,
    trigger_date TIMESTAMPTZ,
    trigger_days_before INTEGER, -- Days before event
    
    -- Message
    title VARCHAR(200) NOT NULL,
    message TEXT NOT NULL,
    
    -- Status
    status VARCHAR(20) DEFAULT 'SCHEDULED' CHECK (
        status IN ('SCHEDULED', 'SENT', 'FAILED', 'CANCELLED', 'EXPIRED')
    ),
    
    -- Delivery
    sent_at TIMESTAMPTZ,
    delivery_method VARCHAR(20) DEFAULT 'EMAIL', -- EMAIL, PUSH, SMS, IN_APP
    delivery_status JSONB,
    
    -- Response Tracking
    acknowledged BOOLEAN DEFAULT FALSE,
    acknowledged_by UUID,
    acknowledged_at TIMESTAMPTZ,
    
    -- Recurrence
    is_recurring BOOLEAN DEFAULT FALSE,
    recurrence_pattern VARCHAR(50), -- DAILY, WEEKLY, MONTHLY
    next_occurrence_at TIMESTAMPTZ,
    
    -- Priority
    priority VARCHAR(20) DEFAULT 'MEDIUM' CHECK (
        priority IN ('LOW', 'MEDIUM', 'HIGH', 'URGENT')
    ),
    
    scheduled_by UUID NOT NULL,
    scheduled_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_contract_reminders_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_contract_reminders_contract ON contract_reminders (contract_id);
CREATE INDEX idx_contract_reminders_trigger ON contract_reminders (trigger_date, status) 
    WHERE status = 'SCHEDULED';
CREATE INDEX idx_contract_reminders_type ON contract_reminders (reminder_type, status);

COMMENT ON TABLE contract_reminders IS 'Contract reminders - maps to internal/domain/reminder/entity.go';

```
=========================================
##  SECTION 13: AUDIT & COMPLIANCE
```sql
-- Domain: internal/domain/audit/
-- Entity: audit/entity.go
-- =========================================

CREATE TABLE contract_audit_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Entity
    entity_type VARCHAR(50) NOT NULL,
    entity_id UUID NOT NULL,
    contract_id UUID, -- For easy filtering
    
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
    financial_record BOOLEAN DEFAULT FALSE,
    legal_record BOOLEAN DEFAULT FALSE,
    
    -- Risk
    risk_level VARCHAR(20),
    requires_review BOOLEAN DEFAULT FALSE,
    
    occurred_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_contract_audit_logs_entity ON contract_audit_logs (entity_type, entity_id, occurred_at DESC);
CREATE INDEX idx_contract_audit_logs_contract ON contract_audit_logs (contract_id, occurred_at DESC) 
    WHERE contract_id IS NOT NULL;
CREATE INDEX idx_contract_audit_logs_actor ON contract_audit_logs (actor_user_id, occurred_at DESC);
CREATE INDEX idx_contract_audit_logs_action ON contract_audit_logs (action, occurred_at DESC);
CREATE INDEX idx_contract_audit_logs_compliance ON contract_audit_logs (occurred_at DESC) 
    WHERE gdpr_relevant = TRUE OR financial_record = TRUE OR legal_record = TRUE;

COMMENT ON TABLE contract_audit_logs IS 'Comprehensive audit trail';

```
=========================================
##  SECTION 14: DIRECT CONTRACT INVITATIONS
```sql
-- Domain: internal/domain/invitation/
-- Entity: invitation/entity.go
-- =========================================

CREATE TABLE contract_invitations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Invitation Details
    client_id UUID NOT NULL,
    freelancer_id UUID NOT NULL,
    
    -- Proposed Terms
    proposed_title VARCHAR(300) NOT NULL,
    proposed_description TEXT,
    proposed_contract_type VARCHAR(30) NOT NULL,
    proposed_rate DECIMAL(10, 2),
    proposed_budget DECIMAL(15, 2),
    currency CHAR(3) DEFAULT 'USD',
    proposed_start_date DATE,
    proposed_duration_days INTEGER,
    
    -- Invitation Message
    invitation_message TEXT,
    
    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'VIEWED', 'ACCEPTED', 'DECLINED', 'EXPIRED', 'WITHDRAWN')
    ),
    
    -- Timeline
    sent_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    viewed_at TIMESTAMPTZ,
    responded_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ NOT NULL,
    
    -- Response
    response_message TEXT,
    decline_reason VARCHAR(100),
    
    -- Result
    contract_id UUID, -- Created contract if accepted
    
    -- Withdrawal
    withdrawn_by UUID,
    withdrawn_at TIMESTAMPTZ,
    withdrawal_reason TEXT,
    
    CONSTRAINT uk_contract_invitations UNIQUE (client_id, freelancer_id, sent_at),
    CONSTRAINT chk_contract_invitations_expires CHECK (expires_at > sent_at)
);

CREATE INDEX idx_contract_invitations_client ON contract_invitations (client_id, status);
CREATE INDEX idx_contract_invitations_freelancer ON contract_invitations (freelancer_id, status);
CREATE INDEX idx_contract_invitations_pending ON contract_invitations (status, expires_at) 
    WHERE status = 'PENDING';
CREATE INDEX idx_contract_invitations_contract ON contract_invitations (contract_id) 
    WHERE contract_id IS NOT NULL;

COMMENT ON TABLE contract_invitations IS 'Direct contract invitations - maps to internal/domain/invitation/entity.go';

```
=========================================
##  SECTION 15: RATE CARDS
```sql
-- Domain: internal/domain/rate_card/
-- Entity: rate_card/entity.go
-- =========================================

CREATE TABLE contract_rate_cards (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    freelancer_id UUID NOT NULL,
    
    -- Rate Card Details
    name VARCHAR(100) NOT NULL,
    description TEXT,
    
    -- Base Rates
    standard_hourly_rate DECIMAL(10, 2) NOT NULL,
    currency CHAR(3) DEFAULT 'USD',
    
    -- Tier-based Rates
    tier_configuration JSONB, -- Structured tier definitions
    
    -- Premium Rates
    rush_rate_multiplier DECIMAL(5, 2) DEFAULT 1.5,
    weekend_rate_multiplier DECIMAL(5, 2) DEFAULT 1.3,
    holiday_rate_multiplier DECIMAL(5, 2) DEFAULT 2.0,
    overtime_rate_multiplier DECIMAL(5, 2) DEFAULT 1.5,
    
    -- Volume Discounts
    volume_discount_rules JSONB,
    
    -- Services Included
    included_services TEXT[],
    excluded_services TEXT[],
    
    -- Minimum Commitments
    minimum_hours_per_week INTEGER,
    minimum_contract_duration_days INTEGER,
    
    -- Validity
    effective_from DATE NOT NULL,
    effective_until DATE,
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    is_default BOOLEAN DEFAULT FALSE,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT chk_contract_rate_cards_rate CHECK (standard_hourly_rate > 0)
);

CREATE INDEX idx_contract_rate_cards_freelancer ON contract_rate_cards (freelancer_id, is_active);
CREATE INDEX idx_contract_rate_cards_validity ON contract_rate_cards (effective_from, effective_until) 
    WHERE is_active = TRUE;

COMMENT ON TABLE contract_rate_cards IS 'Rate cards - maps to internal/domain/rate_card/entity.go';

```
=========================================
##  SECTION 16: CONTRACT ANALYTICS
```sql
-- Domain: internal/domain/analytics/
-- Entity: analytics/entity.go
-- =========================================

CREATE TABLE contract_analytics (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL UNIQUE,
    
    -- Performance Metrics
    on_time_delivery_rate DECIMAL(5, 2),
    milestone_completion_rate DECIMAL(5, 2),
    budget_adherence_rate DECIMAL(5, 2),
    quality_score_avg DECIMAL(5, 2),
    
    -- Time Metrics
    avg_milestone_delay_days DECIMAL(5, 2),
    avg_response_time_hours DECIMAL(8, 2),
    avg_approval_time_hours DECIMAL(8, 2),
    
    -- Communication Metrics
    total_messages_exchanged INTEGER DEFAULT 0,
    avg_messages_per_week DECIMAL(5, 2),
    avg_client_response_hours DECIMAL(8, 2),
    avg_freelancer_response_hours DECIMAL(8, 2),
    
    -- Financial Metrics
    payment_promptness_score DECIMAL(5, 2),
    invoice_dispute_rate DECIMAL(5, 2),
    total_revenue DECIMAL(15, 2),
    
    -- Work Pattern Analysis
    most_productive_day_of_week VARCHAR(20),
    most_productive_hour_range VARCHAR(20),
    avg_daily_hours DECIMAL(5, 2),
    consistency_score DECIMAL(5, 2),
    
    -- Risk Indicators
    dispute_count INTEGER DEFAULT 0,
    revision_count INTEGER DEFAULT 0,
    deadline_miss_count INTEGER DEFAULT 0,
    budget_overrun_count INTEGER DEFAULT 0,
    
    -- Satisfaction
    overall_satisfaction_score DECIMAL(5, 2),
    would_rehire_score DECIMAL(5, 2),
    
    -- Engagement
    engagement_score DECIMAL(5, 2),
    collaboration_quality_score DECIMAL(5, 2),
    
    last_computed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_contract_analytics_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_contract_analytics_contract ON contract_analytics (contract_id);
CREATE INDEX idx_contract_analytics_quality ON contract_analytics (quality_score_avg DESC);
CREATE INDEX idx_contract_analytics_satisfaction ON contract_analytics (overall_satisfaction_score DESC);

COMMENT ON TABLE contract_analytics IS 'Contract analytics - maps to internal/domain/analytics/entity.go';

```
=========================================
##  SECTION 17: COMPLIANCE & LEGAL
```sql
-- Domain: internal/domain/compliance/
-- Entity: compliance/entity.go
-- =========================================

CREATE TABLE contract_compliance (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL UNIQUE,
    
    -- Legal Requirements
    requires_legal_review BOOLEAN DEFAULT FALSE,
    legal_review_completed BOOLEAN DEFAULT FALSE,
    legal_reviewed_by UUID,
    legal_reviewed_at TIMESTAMPTZ,
    legal_review_notes TEXT,
    
    -- Tax Compliance
    tax_form_required VARCHAR(20), -- W9, 1099, etc.
    tax_form_received BOOLEAN DEFAULT FALSE,
    tax_form_received_at TIMESTAMPTZ,
    tax_id_verified BOOLEAN DEFAULT FALSE,
    
    -- Geographic Compliance
    work_location_country CHAR(2),
    work_location_state VARCHAR(100),
    complies_with_local_laws BOOLEAN DEFAULT TRUE,
    jurisdiction_restrictions TEXT[],
    
    -- Labor Laws
    labor_law_classification VARCHAR(50), -- EMPLOYEE, CONTRACTOR, etc.
    misclassification_risk_level VARCHAR(20),
    
    -- Export Control
    export_control_check_required BOOLEAN DEFAULT FALSE,
    export_control_approved BOOLEAN DEFAULT FALSE,
    export_control_reviewed_at TIMESTAMPTZ,
    
    -- Sanctions Screening
    sanctions_screened BOOLEAN DEFAULT FALSE,
    sanctions_clear BOOLEAN DEFAULT TRUE,
    sanctions_screening_at TIMESTAMPTZ,
    sanctions_notes TEXT,
    
    -- Data Protection
    data_protection_impact_assessment BOOLEAN DEFAULT FALSE,
    gdpr_compliant BOOLEAN DEFAULT TRUE,
    ccpa_compliant BOOLEAN DEFAULT TRUE,
    data_processing_agreement_signed BOOLEAN DEFAULT FALSE,
    
    -- IP Protection
    ip_agreement_signed BOOLEAN DEFAULT FALSE,
    ip_agreement_date DATE,
    ip_transfer_documented BOOLEAN DEFAULT FALSE,
    
    -- Insurance
    requires_insurance BOOLEAN DEFAULT FALSE,
    insurance_verified BOOLEAN DEFAULT FALSE,
    insurance_expiry_date DATE,
    insurance_policy_number VARCHAR(100),
    
    -- Audit Trail
    last_compliance_check_at TIMESTAMPTZ,
    next_compliance_check_at TIMESTAMPTZ,
    compliance_officer_id UUID,
    
    -- Risk Assessment
    compliance_risk_score DECIMAL(5, 2),
    non_compliance_issues TEXT[],
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_contract_compliance_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_contract_compliance_contract ON contract_compliance (contract_id);
CREATE INDEX idx_contract_compliance_review ON contract_compliance (legal_review_completed, requires_legal_review);
CREATE INDEX idx_contract_compliance_sanctions ON contract_compliance (sanctions_clear) 
    WHERE sanctions_clear = FALSE;

COMMENT ON TABLE contract_compliance IS 'Compliance tracking - maps to internal/domain/compliance/entity.go';

```
=========================================
##  SECTION 18: CONTRACT FEEDBACK
```sql
-- Domain: internal/domain/feedback/
-- Entity: feedback/entity.go
-- =========================================

CREATE TABLE contract_feedback (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- Feedback Provider
    provided_by UUID NOT NULL,
    provider_role VARCHAR(20) CHECK (
        provider_role IN ('CLIENT', 'FREELANCER')
    ),
    
    -- Overall Rating
    overall_rating INTEGER NOT NULL CHECK (overall_rating BETWEEN 1 AND 5),
    
    -- Detailed Ratings
    communication_rating INTEGER CHECK (communication_rating BETWEEN 1 AND 5),
    quality_rating INTEGER CHECK (quality_rating BETWEEN 1 AND 5),
    professionalism_rating INTEGER CHECK (professionalism_rating BETWEEN 1 AND 5),
    timeliness_rating INTEGER CHECK (timeliness_rating BETWEEN 1 AND 5),
    value_rating INTEGER CHECK (value_rating BETWEEN 1 AND 5),
    
    -- Written Feedback
    feedback_text TEXT,
    strengths TEXT,
    areas_for_improvement TEXT,
    
    -- Recommendations
    would_work_again BOOLEAN,
    would_recommend BOOLEAN,
    recommendation_likelihood INTEGER CHECK (recommendation_likelihood BETWEEN 1 AND 10),
    
    -- Tags
    positive_tags TEXT[],
    negative_tags TEXT[],
    
    -- Visibility
    is_public BOOLEAN DEFAULT FALSE,
    is_anonymous BOOLEAN DEFAULT FALSE,
    
    -- Response
    has_response BOOLEAN DEFAULT FALSE,
    response_text TEXT,
    responded_at TIMESTAMPTZ,
    
    -- Moderation
    is_flagged BOOLEAN DEFAULT FALSE,
    flag_reason TEXT,
    moderation_status VARCHAR(20) DEFAULT 'APPROVED',
    
    -- Review Integration
    review_id UUID, -- Reference to reviews-be
    
    provided_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_contract_feedback_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE,
    CONSTRAINT uk_contract_feedback UNIQUE (contract_id, provided_by)
);

CREATE INDEX idx_contract_feedback_contract ON contract_feedback (contract_id);
CREATE INDEX idx_contract_feedback_provider ON contract_feedback (provided_by);
CREATE INDEX idx_contract_feedback_rating ON contract_feedback (overall_rating DESC);
CREATE INDEX idx_contract_feedback_public ON contract_feedback (is_public, overall_rating DESC) 
    WHERE is_public = TRUE;

COMMENT ON TABLE contract_feedback IS 'Contract feedback - maps to internal/domain/feedback/entity.go';

```
=========================================
##  SECTION 19: CONTRACT SEARCH INDEX
```sql
-- Domain: internal/domain/search/
-- Entity: search/entity.go
-- =========================================

CREATE TABLE contract_search_index (
    contract_id UUID PRIMARY KEY,
    
    -- Core Info
    contract_number VARCHAR(50),
    title VARCHAR(300),
    contract_type VARCHAR(30),
    status VARCHAR(20),
    
    -- Participants (denormalized)
    client_id UUID,
    client_name VARCHAR(200),
    freelancer_id UUID,
    freelancer_name VARCHAR(200),
    
    -- Financial
    total_value DECIMAL(15, 2),
    currency CHAR(3),
    
    -- Timeline
    start_date DATE,
    end_date DATE,
    
    -- Performance
    quality_score DECIMAL(5, 2),
    performance_rating DECIMAL(3, 2),
    
    -- Status Flags
    is_active BOOLEAN,
    is_disputed BOOLEAN,
    is_overbudget BOOLEAN,
    
    -- Tags
    tags TEXT[],
    
    -- Full-Text Search
    search_vector tsvector GENERATED ALWAYS AS (
        setweight(to_tsvector('english', COALESCE(contract_number, '')), 'A') ||
        setweight(to_tsvector('english', COALESCE(title, '')), 'B') ||
        setweight(to_tsvector('english', COALESCE(client_name, '')), 'C') ||
        setweight(to_tsvector('english', COALESCE(freelancer_name, '')), 'C')
    ) STORED,
    
    indexed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);


CREATE INDEX idx_contract_search_vector ON contract_search_index USING gin(search_vector);
CREATE INDEX idx_contract_search_index_client ON contract_search_index (client_id);
CREATE INDEX idx_contract_search_index_freelancer ON contract_search_index (freelancer_id);
CREATE INDEX idx_contract_search_index_status ON contract_search_index (status, start_date DESC);
CREATE INDEX idx_contract_search_index_type ON contract_search_index (contract_type);

COMMENT ON TABLE contract_search_index IS 'Contract search index - maps to internal/domain/search/entity.go';

```
=========================================
##  SECTION 20: INSURANCE
```sql
-- Domain: internal/domain/insurance/
-- Entity: insurance/entity.go
-- =========================================

CREATE TABLE contract_insurance (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- Insurance Policy
    policy_number VARCHAR(100) UNIQUE NOT NULL,
    policy_type VARCHAR(30) CHECK (
        policy_type IN ('PROFESSIONAL_LIABILITY', 'GENERAL_LIABILITY', 'ERRORS_OMISSIONS', 
                       'CYBER_LIABILITY', 'WORKERS_COMP', 'OTHER')
    ),
    
    -- Provider
    insurance_provider VARCHAR(200) NOT NULL,
    provider_contact TEXT,
    
    -- Coverage
    coverage_amount DECIMAL(15, 2) NOT NULL,
    currency CHAR(3) DEFAULT 'USD',
    deductible_amount DECIMAL(12, 2),
    coverage_details TEXT,
    
    -- Validity
    effective_from DATE NOT NULL,
    effective_until DATE NOT NULL,
    
    -- Status
    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (
        status IN ('ACTIVE', 'EXPIRED', 'CANCELLED', 'SUSPENDED')
    ),
    
    -- Verification
    is_verified BOOLEAN DEFAULT FALSE,
    verified_by UUID,
    verified_at TIMESTAMPTZ,
    verification_documents TEXT[],
    
    -- Claims
    claims_filed INTEGER DEFAULT 0,
    has_active_claims BOOLEAN DEFAULT FALSE,
    
    -- Reminders
    renewal_reminder_sent BOOLEAN DEFAULT FALSE,
    expiry_warning_sent BOOLEAN DEFAULT FALSE,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_contract_insurance_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE,
    CONSTRAINT chk_contract_insurance_dates CHECK (effective_until > effective_from)
);

CREATE INDEX idx_contract_insurance_contract ON contract_insurance (contract_id);
CREATE INDEX idx_contract_insurance_policy ON contract_insurance (policy_number);
CREATE INDEX idx_contract_insurance_expiry ON contract_insurance (effective_until, status) 
    WHERE status = 'ACTIVE';

COMMENT ON TABLE contract_insurance IS 'Contract insurance - maps to internal/domain/insurance/entity.go';

-- Insurance Claims
CREATE TABLE insurance_claims (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    insurance_id UUID NOT NULL,
    contract_id UUID NOT NULL,
    
    -- Claim Details
    claim_number VARCHAR(100) UNIQUE NOT NULL,
    claim_type VARCHAR(30),
    incident_description TEXT NOT NULL,
    
    -- Amount
    claim_amount DECIMAL(12, 2) NOT NULL,
    approved_amount DECIMAL(12, 2),
    currency CHAR(3) DEFAULT 'USD',
    
    -- Status
    status VARCHAR(20) DEFAULT 'SUBMITTED' CHECK (
        status IN ('SUBMITTED', 'UNDER_REVIEW', 'APPROVED', 'PARTIALLY_APPROVED', 
                  'DENIED', 'SETTLED', 'WITHDRAWN')
    ),
    
    -- Timeline
    incident_date DATE NOT NULL,
    filed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    resolved_at TIMESTAMPTZ,
    
    -- Resolution
    denial_reason TEXT,
    settlement_notes TEXT,
    
    -- Documents
    supporting_documents TEXT[],
    
    filed_by UUID NOT NULL,
    
    CONSTRAINT fk_insurance_claims_insurance FOREIGN KEY (insurance_id) 
        REFERENCES contract_insurance(id) ON DELETE CASCADE,
    CONSTRAINT fk_insurance_claims_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_insurance_claims_insurance ON insurance_claims (insurance_id);
CREATE INDEX idx_insurance_claims_contract ON insurance_claims (contract_id);
CREATE INDEX idx_insurance_claims_status ON insurance_claims (status, filed_at DESC);

```
=========================================
##  SECTION 21: ESCROW MANAGEMENT
```sql
-- Domain: internal/domain/escrow/
-- Entity: escrow/entity.go
-- =========================================

CREATE TABLE contract_escrow (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- Escrow Account (in financial-be)
    escrow_account_id UUID NOT NULL UNIQUE,
    
    -- Balances
    total_funded DECIMAL(15, 2) NOT NULL DEFAULT 0,
    total_held DECIMAL(15, 2) NOT NULL DEFAULT 0,
    total_released DECIMAL(15, 2) DEFAULT 0,
    total_refunded DECIMAL(15, 2) DEFAULT 0,
    available_balance DECIMAL(15, 2) DEFAULT 0,
    currency CHAR(3) DEFAULT 'USD',
    
    -- Status
    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (
        status IN ('PENDING', 'ACTIVE', 'FROZEN', 'CLOSED')
    ),
    
    -- Funding
    funding_complete BOOLEAN DEFAULT FALSE,
    funded_at TIMESTAMPTZ,
    
    -- Release Rules
    auto_release_enabled BOOLEAN DEFAULT TRUE,
    release_requires_approval BOOLEAN DEFAULT TRUE,
    release_delay_days INTEGER DEFAULT 0,
    
    -- Holds
    active_holds_count INTEGER DEFAULT 0,
    total_held_by_holds DECIMAL(15, 2) DEFAULT 0,
    
    -- Activity Tracking
    last_transaction_at TIMESTAMPTZ,
    transaction_count INTEGER DEFAULT 0,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_contract_escrow_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_contract_escrow_contract ON contract_escrow (contract_id);
CREATE INDEX idx_contract_escrow_account ON contract_escrow (escrow_account_id);
CREATE INDEX idx_contract_escrow_status ON contract_escrow (status);

COMMENT ON TABLE contract_escrow IS 'Escrow management - maps to internal/domain/escrow/entity.go';

-- Escrow Transactions Log
CREATE TABLE escrow_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    escrow_id UUID NOT NULL,
    
    -- Transaction Details
    transaction_type VARCHAR(30) NOT NULL CHECK (
        transaction_type IN ('FUND', 'HOLD', 'RELEASE', 'REFUND', 'FEE', 'ADJUSTMENT')
    ),
    amount DECIMAL(12, 2) NOT NULL,
    currency CHAR(3) DEFAULT 'USD',
    
    -- Reference
    reference_type VARCHAR(30), -- MILESTONE, TIMESHEET, DISPUTE
    reference_id UUID,
    
    -- Description
    description TEXT,
    
    -- Balance Snapshot
    balance_before DECIMAL(15, 2) NOT NULL,
    balance_after DECIMAL(15, 2) NOT NULL,
    
    -- Financial Integration
    financial_transaction_id UUID, -- Reference to financial-be
    
    -- Status
    status VARCHAR(20) DEFAULT 'COMPLETED',
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_escrow_transactions_escrow FOREIGN KEY (escrow_id) 
        REFERENCES contract_escrow(id) ON DELETE CASCADE
);

CREATE INDEX idx_escrow_transactions_escrow ON escrow_transactions (escrow_id, created_at DESC);
CREATE INDEX idx_escrow_transactions_reference ON escrow_transactions (reference_type, reference_id);

```
=========================================
##  SECTION 22: TERMINATION
```sql
-- Domain: internal/domain/termination/
-- Entity: termination/entity.go
-- =========================================

CREATE TABLE contract_terminations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL UNIQUE,
    
    -- Termination Details
    termination_type VARCHAR(30) NOT NULL CHECK (
        termination_type IN ('MUTUAL', 'UNILATERAL_CLIENT', 'UNILATERAL_FREELANCER', 
                            'BREACH', 'FORCE_MAJEURE', 'ADMIN_FORCED')
    ),
    reason_category VARCHAR(50),
    reason_details TEXT NOT NULL,
    
    -- Initiator
    initiated_by UUID NOT NULL,
    initiated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Notice Period
    notice_period_days INTEGER,
    notice_date DATE,
    effective_date DATE NOT NULL,
    
    -- Status
    status VARCHAR(20) DEFAULT 'PROPOSED' CHECK (
        status IN ('PROPOSED', 'NEGOTIATING', 'AGREED', 'CONTESTED', 'FINALIZED')
    ),
    
    -- Negotiation
    requires_negotiation BOOLEAN DEFAULT TRUE,
    negotiation_deadline TIMESTAMPTZ,
    
    -- Approval (for mutual termination)
    client_approved BOOLEAN DEFAULT FALSE,
    client_approved_at TIMESTAMPTZ,
    freelancer_approved BOOLEAN DEFAULT FALSE,
    freelancer_approved_at TIMESTAMPTZ,
    
    -- Financial Settlement
    final_payment_due DECIMAL(12, 2),
    early_termination_penalty DECIMAL(12, 2),
    outstanding_balance DECIMAL(12, 2),
    refund_due DECIMAL(12, 2),
    settlement_amount DECIMAL(12, 2),
    currency CHAR(3) DEFAULT 'USD',
    
    -- Settlement Status
    settlement_completed BOOLEAN DEFAULT FALSE,
    settlement_completed_at TIMESTAMPTZ,
    settlement_transaction_id UUID,
    
    -- Work Completion
    work_completion_percentage DECIMAL(5, 2),
    deliverables_handed_over BOOLEAN DEFAULT FALSE,
    handover_completed_at TIMESTAMPTZ,
    
    -- Post-Termination
    non_compete_applies BOOLEAN DEFAULT FALSE,
    non_compete_duration_days INTEGER,
    confidentiality_remains BOOLEAN DEFAULT TRUE,
    ip_transfer_completed BOOLEAN DEFAULT FALSE,
    
    -- Documentation
    termination_agreement_url TEXT,
    termination_agreement_signed BOOLEAN DEFAULT FALSE,
    termination_documents TEXT[],
    
    -- Dispute
    contested BOOLEAN DEFAULT FALSE,
    contest_reason TEXT,
    contested_at TIMESTAMPTZ,
    
    finalized_at TIMESTAMPTZ,
    
    CONSTRAINT fk_contract_terminations_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE,
    CONSTRAINT chk_contract_terminations_dates CHECK (effective_date >= notice_date)
);

CREATE INDEX idx_contract_terminations_contract ON contract_terminations (contract_id);
CREATE INDEX idx_contract_terminations_status ON contract_terminations (status, initiated_at DESC);
CREATE INDEX idx_contract_terminations_effective ON contract_terminations (effective_date);

COMMENT ON TABLE contract_terminations IS 'Contract terminations - maps to internal/domain/termination/entity.go';

```
=========================================
##  SECTION 23: COLLABORATION WORKSPACE
```sql
-- Domain: internal/domain/workspace/
-- Entity: workspace/entity.go
-- =========================================

CREATE TABLE contract_workspaces (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL UNIQUE,
    
    -- Workspace Configuration
    workspace_name VARCHAR(200) NOT NULL,
    description TEXT,
    
    -- Members
    member_ids UUID[] NOT NULL,
    
    -- Access Control
    access_level VARCHAR(20) DEFAULT 'PRIVATE' CHECK (
        access_level IN ('PRIVATE', 'CONTRACT_PARTIES', 'TEAM', 'PUBLIC')
    ),
    
    -- Features
    file_sharing_enabled BOOLEAN DEFAULT TRUE,
    commenting_enabled BOOLEAN DEFAULT TRUE,
    version_control_enabled BOOLEAN DEFAULT TRUE,
    real_time_collaboration BOOLEAN DEFAULT FALSE,
    
    -- Storage
    storage_used_bytes BIGINT DEFAULT 0,
    storage_limit_bytes BIGINT,
    
    -- Activity
    total_documents INTEGER DEFAULT 0,
    total_comments INTEGER DEFAULT 0,
    last_activity_at TIMESTAMPTZ,
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    is_archived BOOLEAN DEFAULT FALSE,
    archived_at TIMESTAMPTZ,
    
    created_by UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_contract_workspaces_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_contract_workspaces_contract ON contract_workspaces (contract_id);

COMMENT ON TABLE contract_workspaces IS 'Collaboration workspaces - maps to internal/domain/workspace/entity.go';

-- Workspace Documents
CREATE TABLE workspace_documents (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    workspace_id UUID NOT NULL,
    
    -- Document Details
    document_name VARCHAR(255) NOT NULL,
    document_type VARCHAR(50),
    description TEXT,
    
    -- File Reference (storage-be)
    file_id UUID NOT NULL,
    file_url TEXT NOT NULL,
    file_size BIGINT NOT NULL,
    file_mime_type VARCHAR(100),
    
    -- Version
    version INTEGER DEFAULT 1,
    previous_version_id UUID,
    
    -- Access
    visibility VARCHAR(20) DEFAULT 'ALL_MEMBERS' CHECK (
        visibility IN ('ALL_MEMBERS', 'CLIENT_ONLY', 'FREELANCER_ONLY', 'CUSTOM')
    ),
    
    -- Status
    is_locked BOOLEAN DEFAULT FALSE,
    locked_by UUID,
    locked_at TIMESTAMPTZ,
    
    -- Activity
    download_count INTEGER DEFAULT 0,
    comment_count INTEGER DEFAULT 0,
    last_accessed_at TIMESTAMPTZ,
    
    uploaded_by UUID NOT NULL,
    uploaded_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_workspace_documents_workspace FOREIGN KEY (workspace_id) 
        REFERENCES contract_workspaces(id) ON DELETE CASCADE,
    CONSTRAINT fk_workspace_documents_previous FOREIGN KEY (previous_version_id) 
        REFERENCES workspace_documents(id)
);

CREATE INDEX idx_workspace_documents_workspace ON workspace_documents (workspace_id, uploaded_at DESC);
CREATE INDEX idx_workspace_documents_file ON workspace_documents (file_id);

-- Workspace Comments
CREATE TABLE workspace_comments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    document_id UUID NOT NULL,
    
    -- Comment Details
    comment_text TEXT NOT NULL,
    
    -- Threading
    parent_comment_id UUID,
    
    -- Reactions
    reactions JSONB, -- emoji reactions
    
    -- Status
    is_resolved BOOLEAN DEFAULT FALSE,
    resolved_by UUID,
    resolved_at TIMESTAMPTZ,
    
    -- Edit History
    is_edited BOOLEAN DEFAULT FALSE,
    edited_at TIMESTAMPTZ,
    
    commented_by UUID NOT NULL,
    commented_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_workspace_comments_document FOREIGN KEY (document_id) 
        REFERENCES workspace_documents(id) ON DELETE CASCADE,
    CONSTRAINT fk_workspace_comments_parent FOREIGN KEY (parent_comment_id) 
        REFERENCES workspace_comments(id)
);

CREATE INDEX idx_workspace_comments_document ON workspace_comments (document_id, commented_at DESC);
CREATE INDEX idx_workspace_comments_parent ON workspace_comments (parent_comment_id) 
    WHERE parent_comment_id IS NOT NULL;

```
=========================================
##  SECTION 24: RECURRING CONTRACTS
```sql
-- Domain: internal/domain/recurring/
-- Entity: recurring/entity.go
-- =========================================

CREATE TABLE recurring_contracts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    base_contract_id UUID NOT NULL UNIQUE,
    
    -- Recurrence Configuration
    recurrence_pattern VARCHAR(30) NOT NULL CHECK (
        recurrence_pattern IN ('WEEKLY', 'BIWEEKLY', 'MONTHLY', 'QUARTERLY', 'ANNUALLY')
    ),
    recurrence_interval INTEGER DEFAULT 1,
    
    -- Auto-Renewal
    auto_renew BOOLEAN DEFAULT TRUE,
    requires_approval BOOLEAN DEFAULT FALSE,
    notification_days_before INTEGER DEFAULT 30,
    
    -- Terms
    renews_indefinitely BOOLEAN DEFAULT FALSE,
    max_renewals INTEGER,
    current_renewal_count INTEGER DEFAULT 0,
    
    -- Next Renewal
    next_renewal_date DATE NOT NULL,
    last_renewed_at TIMESTAMPTZ,
    
    -- Status
    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (
        status IN ('ACTIVE', 'PAUSED', 'CANCELLED', 'COMPLETED')
    ),
    
    -- Cancellation
    cancel_notice_period_days INTEGER DEFAULT 30,
    cancellation_deadline DATE,
    cancelled_at TIMESTAMPTZ,
    cancellation_reason TEXT,
    
    -- Performance Tracking
    total_renewal_value DECIMAL(15, 2) DEFAULT 0,
    currency CHAR(3) DEFAULT 'USD',
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_recurring_contracts_base FOREIGN KEY (base_contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_recurring_contracts_base ON recurring_contracts (base_contract_id);
CREATE INDEX idx_recurring_contracts_renewal ON recurring_contracts (next_renewal_date, status) 
    WHERE status = 'ACTIVE';

COMMENT ON TABLE recurring_contracts IS 'Recurring contracts - maps to internal/domain/recurring/entity.go';

-- Renewal History
CREATE TABLE contract_renewal_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    recurring_contract_id UUID NOT NULL,
    
    -- Renewed Contract
    new_contract_id UUID NOT NULL,
    
    -- Renewal Details
    renewal_number INTEGER NOT NULL,
    renewal_date DATE NOT NULL,
    
    -- Terms
    renewal_value DECIMAL(12, 2),
    currency CHAR(3) DEFAULT 'USD',
    terms_changed BOOLEAN DEFAULT FALSE,
    changes_summary TEXT,
    
    -- Approval
    approved_by_client BOOLEAN DEFAULT TRUE,
    approved_by_freelancer BOOLEAN DEFAULT TRUE,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_contract_renewal_history_recurring FOREIGN KEY (recurring_contract_id) 
        REFERENCES recurring_contracts(id) ON DELETE CASCADE,
    CONSTRAINT fk_contract_renewal_history_contract FOREIGN KEY (new_contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_contract_renewal_history_recurring ON contract_renewal_history (recurring_contract_id, renewal_number DESC);

```
=========================================
##  SECTION 25: E-SIGNATURES
```sql
-- Domain: internal/domain/signature/
-- Entity: signature/entity.go
-- =========================================

CREATE TABLE contract_signatures (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- Signer Details
    signer_id UUID NOT NULL,
    signer_role VARCHAR(20) NOT NULL CHECK (
        signer_role IN ('CLIENT', 'FREELANCER', 'WITNESS', 'GUARANTOR')
    ),
    signer_name VARCHAR(200) NOT NULL,
    signer_email CITEXT NOT NULL,
    
    -- Signature
    signature_type VARCHAR(20) DEFAULT 'ELECTRONIC' CHECK (
        signature_type IN ('ELECTRONIC', 'DIGITAL', 'BIOMETRIC', 'TYPED')
    ),
    signature_image_url TEXT,
    signature_data TEXT, -- Encrypted signature data
    
    -- Certificate (for digital signatures)
    certificate_id VARCHAR(100),
    certificate_issuer VARCHAR(200),
    certificate_valid_until DATE,
    
    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'SENT', 'VIEWED', 'SIGNED', 'DECLINED', 'EXPIRED')
    ),
    
    -- Timeline
    invitation_sent_at TIMESTAMPTZ,
    document_viewed_at TIMESTAMPTZ,
    signed_at TIMESTAMPTZ,
    declined_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ,
    
    -- Verification
    ip_address INET,
    user_agent TEXT,
    geolocation JSONB,
    device_info JSONB,
    
    -- Legal Binding
    is_legally_binding BOOLEAN DEFAULT TRUE,
    consent_given BOOLEAN DEFAULT FALSE,
    consent_timestamp TIMESTAMPTZ,
    
    -- Decline Reason
    decline_reason TEXT,
    
    -- Audit Trail
    verification_hash VARCHAR(64), -- SHA-256 of signature + metadata
    blockchain_hash VARCHAR(66), -- Optional blockchain anchoring
    
    -- Reminders
    reminder_count INTEGER DEFAULT 0,
    last_reminder_sent_at TIMESTAMPTZ,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_contract_signatures_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_contract_signatures_contract ON contract_signatures (contract_id);
CREATE INDEX idx_contract_signatures_signer ON contract_signatures (signer_id, status);
CREATE INDEX idx_contract_signatures_pending ON contract_signatures (status, invitation_sent_at) 
    WHERE status IN ('PENDING', 'SENT', 'VIEWED');

COMMENT ON TABLE contract_signatures IS 'E-signatures - maps to internal/domain/signature/entity.go';

```
=========================================
##  SECTION 26: INVOICING
```sql
-- Domain: internal/domain/invoice/
-- Entity: invoice/entity.go
-- =========================================

CREATE TABLE invoices (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- Invoice Identity
    invoice_number VARCHAR(50) UNIQUE NOT NULL,
    invoice_type VARCHAR(30) DEFAULT 'STANDARD' CHECK (
        invoice_type IN ('STANDARD', 'MILESTONE', 'TIMESHEET', 'RETAINER', 'ADJUSTMENT', 'CREDIT_NOTE')
    ),
    
    -- Reference
    reference_type VARCHAR(30), -- MILESTONE, TIMESHEET, CONTRACT
    reference_id UUID,
    
    -- Parties
    billed_to UUID NOT NULL, -- Client
    billed_by UUID NOT NULL, -- Freelancer
    
    -- Financial Details
    subtotal DECIMAL(12, 2) NOT NULL,
    tax_total DECIMAL(12, 2) DEFAULT 0,
    discount_amount DECIMAL(12, 2) DEFAULT 0,
    total_amount DECIMAL(12, 2) NOT NULL,
    currency CHAR(3) DEFAULT 'USD',
    
    -- Payment Terms
    payment_terms VARCHAR(100),
    due_date DATE NOT NULL,
    net_days INTEGER DEFAULT 30,
    
    -- Status
    status VARCHAR(20) DEFAULT 'DRAFT' CHECK (
        status IN ('DRAFT', 'ISSUED', 'SENT', 'VIEWED', 'PARTIALLY_PAID', 'PAID', 
                   'OVERDUE', 'CANCELLED', 'REFUNDED', 'DISPUTED')
    ),
    
    -- Timeline
    issued_at TIMESTAMPTZ,
    sent_at TIMESTAMPTZ,
    viewed_at TIMESTAMPTZ,
    paid_at TIMESTAMPTZ,
    cancelled_at TIMESTAMPTZ,
    
    -- Payment Tracking
    amount_paid DECIMAL(12, 2) DEFAULT 0,
    amount_due DECIMAL(12, 2),
    last_payment_at TIMESTAMPTZ,
    
    -- Overdue Management
    is_overdue BOOLEAN DEFAULT FALSE,
    days_overdue INTEGER DEFAULT 0,
    late_fee_amount DECIMAL(12, 2),
    
    -- Document
    invoice_pdf_url TEXT,
    invoice_html TEXT,
    
    -- Notes
    notes TEXT,
    internal_notes TEXT,
    
    -- Cancellation
    cancellation_reason TEXT,
    cancelled_by UUID,
    
    -- Integration
    external_invoice_id VARCHAR(100), -- For payment gateway
    payment_link_url TEXT,
    
    created_by UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_invoices_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE,
    CONSTRAINT chk_invoices_amounts CHECK (total_amount >= 0 AND subtotal >= 0)
);

CREATE INDEX idx_invoices_contract ON invoices (contract_id, issued_at DESC);
CREATE INDEX idx_invoices_number ON invoices (invoice_number);
CREATE INDEX idx_invoices_status ON invoices (status, due_date);
CREATE INDEX idx_invoices_billed_to ON invoices (billed_to, status);
CREATE INDEX idx_invoices_billed_by ON invoices (billed_by, status);
CREATE INDEX idx_invoices_overdue ON invoices (is_overdue, due_date) WHERE is_overdue = TRUE;
CREATE INDEX idx_invoices_reference ON invoices (reference_type, reference_id);

COMMENT ON TABLE invoices IS 'Contract invoices - maps to internal/domain/invoice/entity.go';

-- Invoice Line Items
CREATE TABLE invoice_line_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    invoice_id UUID NOT NULL,
    
    -- Line Item Details
    line_number INTEGER NOT NULL,
    item_type VARCHAR(30) CHECK (
        item_type IN ('SERVICE', 'PRODUCT', 'HOUR', 'MILESTONE', 'ADJUSTMENT', 'FEE', 'TAX')
    ),
    description TEXT NOT NULL,
    
    -- Quantity & Pricing
    quantity DECIMAL(10, 2) DEFAULT 1,
    unit_price DECIMAL(12, 2) NOT NULL,
    amount DECIMAL(12, 2) NOT NULL,
    
    -- Tax
    is_taxable BOOLEAN DEFAULT TRUE,
    tax_rate DECIMAL(5, 2),
    tax_amount DECIMAL(12, 2) DEFAULT 0,
    
    -- Reference
    reference_type VARCHAR(30), -- MILESTONE, TIME_ENTRY, DELIVERABLE
    reference_id UUID,
    
    -- Metadata
    metadata JSONB,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_invoice_line_items_invoice FOREIGN KEY (invoice_id) 
        REFERENCES invoices(id) ON DELETE CASCADE,
    CONSTRAINT uk_invoice_line_items UNIQUE (invoice_id, line_number),
    CONSTRAINT chk_invoice_line_items_amount CHECK (amount >= 0)
);

CREATE INDEX idx_invoice_line_items_invoice ON invoice_line_items (invoice_id, line_number);

-- Invoice Payment History
CREATE TABLE invoice_payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    invoice_id UUID NOT NULL,
    
    -- Payment Details
    payment_amount DECIMAL(12, 2) NOT NULL,
    currency CHAR(3) DEFAULT 'USD',
    payment_method VARCHAR(50),
    
    -- Reference to financial-be
    transaction_id UUID,
    
    -- Status
    status VARCHAR(20) DEFAULT 'COMPLETED',
    
    -- Notes
    notes TEXT,
    
    paid_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_invoice_payments_invoice FOREIGN KEY (invoice_id) 
        REFERENCES invoices(id) ON DELETE CASCADE
);

CREATE INDEX idx_invoice_payments_invoice ON invoice_payments (invoice_id, paid_at DESC);

```
=========================================
##  SECTION 27: SERVICE LEVEL AGREEMENTS (SLA)
```sql
-- Domain: internal/domain/sla/
-- Entity: sla/entity.go
-- =========================================

CREATE TABLE slas (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- SLA Details
    name VARCHAR(200) NOT NULL,
    description TEXT,
    sla_type VARCHAR(30) CHECK (
        sla_type IN ('RESPONSE_TIME', 'DELIVERY_TIME', 'QUALITY', 'AVAILABILITY', 'UPTIME', 'CUSTOM')
    ),
    
    -- Effective Period
    effective_from DATE NOT NULL,
    effective_until DATE,
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    
    -- Monitoring
    monitoring_frequency VARCHAR(20) DEFAULT 'DAILY' CHECK (
        monitoring_frequency IN ('HOURLY', 'DAILY', 'WEEKLY', 'MONTHLY', 'REAL_TIME')
    ),
    
    -- Penalties & Rewards
    has_penalties BOOLEAN DEFAULT TRUE,
    has_rewards BOOLEAN DEFAULT FALSE,
    penalty_cap DECIMAL(12, 2),
    reward_cap DECIMAL(12, 2),
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_slas_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_slas_contract ON slas (contract_id, is_active);

COMMENT ON TABLE slas IS 'Service Level Agreements - maps to internal/domain/sla/entity.go';

-- SLA Metrics/Targets
CREATE TABLE sla_metrics (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sla_id UUID NOT NULL,
    
    -- Metric Definition
    metric_name VARCHAR(100) NOT NULL,
    metric_type VARCHAR(50) NOT NULL,
    metric_description TEXT,
    
    -- Target
    target_value DECIMAL(12, 2) NOT NULL,
    threshold_value DECIMAL(12, 2), -- Warning threshold
    unit VARCHAR(50) NOT NULL, -- MINUTES, HOURS, PERCENTAGE, COUNT
    
    -- Measurement
    measurement_method VARCHAR(100),
    aggregation_method VARCHAR(30) CHECK (
        aggregation_method IN ('AVERAGE', 'MEDIAN', 'P95', 'P99', 'MAX', 'MIN', 'SUM')
    ),
    
    -- Conditions
    applies_to TEXT, -- What this metric applies to
    exclusions TEXT[], -- What to exclude from measurement
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_sla_metrics_sla FOREIGN KEY (sla_id) 
        REFERENCES slas(id) ON DELETE CASCADE
);

CREATE INDEX idx_sla_metrics_sla ON sla_metrics (sla_id);

-- SLA Breaches
CREATE TABLE sla_breaches (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sla_id UUID NOT NULL,
    metric_id UUID,
    
    -- Breach Details
    metric_type VARCHAR(50) NOT NULL,
    target_value DECIMAL(12, 2) NOT NULL,
    actual_value DECIMAL(12, 2) NOT NULL,
    variance DECIMAL(12, 2),
    
    -- Severity
    severity VARCHAR(20) CHECK (
        severity IN ('MINOR', 'MODERATE', 'MAJOR', 'CRITICAL')
    ),
    
    -- Impact
    impact_description TEXT,
    affected_deliverables UUID[],
    
    -- Timeline
    breach_period_start TIMESTAMPTZ NOT NULL,
    breach_period_end TIMESTAMPTZ,
    detected_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Resolution
    status VARCHAR(20) DEFAULT 'OPEN' CHECK (
        status IN ('OPEN', 'ACKNOWLEDGED', 'INVESTIGATING', 'RESOLVED', 'WAIVED')
    ),
    resolution_notes TEXT,
    resolved_at TIMESTAMPTZ,
    resolved_by UUID,
    
    -- Penalties
    penalty_type VARCHAR(30) CHECK (
        penalty_type IN ('FINANCIAL', 'SERVICE_CREDIT', 'PERFORMANCE_SCORE', 'NONE')
    ),
    penalty_amount DECIMAL(12, 2),
    penalty_applied BOOLEAN DEFAULT FALSE,
    
    -- Waiver
    is_waived BOOLEAN DEFAULT FALSE,
    waived_by UUID,
    waiver_reason TEXT,
    
    CONSTRAINT fk_sla_breaches_sla FOREIGN KEY (sla_id) 
        REFERENCES slas(id) ON DELETE CASCADE,
    CONSTRAINT fk_sla_breaches_metric FOREIGN KEY (metric_id) REFERENCES sla_metrics(id) ON DELETE SET NULL
);

CREATE INDEX idx_sla_breaches_sla ON sla_breaches (sla_id, detected_at DESC);
CREATE INDEX idx_sla_breaches_status ON sla_breaches (status, severity DESC);
CREATE INDEX idx_sla_breaches_open ON sla_breaches (status, detected_at) WHERE status IN ('OPEN', 'ACKNOWLEDGED');

-- SLA Performance Records
CREATE TABLE sla_performance_records (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sla_id UUID NOT NULL,
    metric_id UUID NOT NULL,
    
    -- Measurement Period
    period_start TIMESTAMPTZ NOT NULL,
    period_end TIMESTAMPTZ NOT NULL,
    
    -- Performance
    measured_value DECIMAL(12, 2) NOT NULL,
    target_value DECIMAL(12, 2) NOT NULL,
    performance_percentage DECIMAL(5, 2),
    is_met BOOLEAN NOT NULL,
    
    -- Context
    sample_size INTEGER,
    measurement_method VARCHAR(100),
    
    recorded_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_sla_performance_records_sla FOREIGN KEY (sla_id) 
        REFERENCES slas(id) ON DELETE CASCADE,
    CONSTRAINT fk_sla_performance_records_metric FOREIGN KEY (metric_id) 
        REFERENCES sla_metrics(id) ON DELETE CASCADE
);

CREATE INDEX idx_sla_performance_records_sla ON sla_performance_records (sla_id, period_start DESC);
CREATE INDEX idx_sla_performance_records_metric ON sla_performance_records (metric_id, period_start DESC);

```
=========================================
##  SECTION 28: AGENCY CONTRACTS
```sql
-- Domain: internal/domain/agency/
-- Entity: agency/entity.go
-- =========================================

CREATE TABLE agency_contracts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL UNIQUE,
    agency_id UUID NOT NULL, -- Reference to users-be agency
    
    -- Agency Details
    agency_name VARCHAR(200) NOT NULL,
    agency_representative_id UUID,
    
    -- Team Configuration
    team_size INTEGER DEFAULT 0,
    max_team_size INTEGER,
    
    -- Billing Model
    billing_model VARCHAR(30) DEFAULT 'AGGREGATE' CHECK (
        billing_model IN ('AGGREGATE', 'INDIVIDUAL', 'HYBRID')
    ),
    
    -- Financial
    agency_fee_percentage DECIMAL(5, 2),
    agency_fee_amount DECIMAL(12, 2),
    total_team_cost DECIMAL(15, 2) DEFAULT 0,
    
    -- Management
    lead_member_id UUID,
    requires_agency_approval BOOLEAN DEFAULT TRUE,
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_agency_contracts_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_agency_contracts_contract ON agency_contracts (contract_id);
CREATE INDEX idx_agency_contracts_agency ON agency_contracts (agency_id);

COMMENT ON TABLE agency_contracts IS 'Agency contracts - maps to internal/domain/agency/entity.go';

-- Agency Team Members
CREATE TABLE agency_members (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    agency_contract_id UUID NOT NULL,
    user_id UUID NOT NULL, -- Freelancer
    
    -- Role
    role VARCHAR(100) NOT NULL,
    responsibilities TEXT,
    
    -- Billing
    bill_rate DECIMAL(10, 2),
    currency CHAR(3) DEFAULT 'USD',
    
    -- Revenue Split
    split_percentage DECIMAL(5, 2),
    split_amount DECIMAL(12, 2),
    
    -- Timeline
    joined_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    left_at TIMESTAMPTZ,
    
    -- Status
    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (
        status IN ('INVITED', 'ACTIVE', 'INACTIVE', 'REMOVED')
    ),
    
    -- Performance
    hours_worked DECIMAL(10, 2) DEFAULT 0,
    amount_earned DECIMAL(12, 2) DEFAULT 0,
    
    CONSTRAINT fk_agency_members_agency_contract FOREIGN KEY (agency_contract_id) 
        REFERENCES agency_contracts(id) ON DELETE CASCADE,
    CONSTRAINT uk_agency_members UNIQUE (agency_contract_id, user_id)
);

CREATE INDEX idx_agency_members_agency_contract ON agency_members (agency_contract_id);
CREATE INDEX idx_agency_members_user ON agency_members (user_id);

-- Agency Billing Splits
CREATE TABLE agency_billing_splits (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    agency_contract_id UUID NOT NULL,
    
    -- Period
    billing_period_start DATE NOT NULL,
    billing_period_end DATE NOT NULL,
    
    -- Total Revenue
    total_revenue DECIMAL(15, 2) NOT NULL,
    currency CHAR(3) DEFAULT 'USD',
    
    -- Agency Fee
    agency_fee DECIMAL(12, 2) NOT NULL,
    
    -- Member Splits
    member_splits JSONB NOT NULL, -- {member_id: amount}
    
    -- Status
    status VARCHAR(20) DEFAULT 'CALCULATED' CHECK (
        status IN ('CALCULATED', 'APPROVED', 'DISTRIBUTED')
    ),
    
    calculated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    approved_at TIMESTAMPTZ,
    distributed_at TIMESTAMPTZ,
    
    CONSTRAINT fk_agency_billing_splits_agency_contract FOREIGN KEY (agency_contract_id) 
        REFERENCES agency_contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_agency_billing_splits_agency_contract ON agency_billing_splits (agency_contract_id, billing_period_start DESC);

```
=========================================
##  SECTION 29: INTELLECTUAL PROPERTY RIGHTS
```sql
-- Domain: internal/domain/ip_rights/
-- Entity: ip_rights/entity.go
-- =========================================

CREATE TABLE ip_rights (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- IP Assignment
    assignment_type VARCHAR(30) NOT NULL CHECK (
        assignment_type IN ('FULL_TRANSFER', 'LICENSE', 'SHARED_OWNERSHIP', 'RETAINED', 'WORK_FOR_HIRE')
    ),
    
    -- Transfer Details
    transfer_date DATE,
    effective_date DATE NOT NULL,
    
    -- Scope
    ip_scope TEXT NOT NULL,
    included_works TEXT[],
    excluded_works TEXT[],
    
    -- Rights Granted
    rights_granted TEXT[],
    restrictions TEXT[],
    
    -- Geographic Scope
    territory VARCHAR(100) DEFAULT 'WORLDWIDE',
    geographic_restrictions TEXT[],
    
    -- Duration
    perpetual BOOLEAN DEFAULT TRUE,
    term_start_date DATE,
    term_end_date DATE,
    
    -- License Terms
    license_terms TEXT,
    sublicensing_allowed BOOLEAN DEFAULT FALSE,
    modification_allowed BOOLEAN DEFAULT TRUE,
    commercial_use_allowed BOOLEAN DEFAULT TRUE,
    
    -- Attribution
    requires_attribution BOOLEAN DEFAULT FALSE,
    attribution_text TEXT,
    
    -- Protection
    confidentiality_applies BOOLEAN DEFAULT TRUE,
    non_compete_applies BOOLEAN DEFAULT FALSE,
    non_compete_duration_months INTEGER,
    
    -- Documentation
    ip_documentation_urls TEXT[],
    signed_agreement_url TEXT,
    
    -- Status
    status VARCHAR(20) DEFAULT 'DRAFT' CHECK (
        status IN ('DRAFT', 'PENDING_SIGNATURE', 'ACTIVE', 'EXPIRED', 'TERMINATED')
    ),
    
    -- Notes
    notes TEXT,
    special_conditions TEXT,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_ip_rights_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_ip_rights_contract ON ip_rights (contract_id);
CREATE INDEX idx_ip_rights_status ON ip_rights (status);

COMMENT ON TABLE ip_rights IS 'IP rights management - maps to internal/domain/ip_rights/entity.go';

-- IP Licenses
CREATE TABLE ip_licenses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ip_rights_id UUID NOT NULL,
    
    -- License Details
    license_type VARCHAR(30) CHECK (
        license_type IN ('EXCLUSIVE', 'NON_EXCLUSIVE', 'SOLE', 'SUBLICENSE')
    ),
    license_name VARCHAR(200),
    
    -- Scope
    scope_description TEXT,
    permitted_uses TEXT[],
    prohibited_uses TEXT[],
    
    -- Geographic
    territory VARCHAR(100) DEFAULT 'WORLDWIDE',
    
    -- Duration
    term_start DATE NOT NULL,
    term_end DATE,
    is_perpetual BOOLEAN DEFAULT FALSE,
    
    -- Financial
    license_fee DECIMAL(12, 2),
    royalty_percentage DECIMAL(5, 2),
    currency CHAR(3) DEFAULT 'USD',
    
    -- Sublicensing
    sublicensing_allowed BOOLEAN DEFAULT FALSE,
    sublicense_approval_required BOOLEAN DEFAULT TRUE,
    
    -- Termination
    termination_conditions TEXT[],
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_ip_licenses_ip_rights FOREIGN KEY (ip_rights_id) 
        REFERENCES ip_rights(id) ON DELETE CASCADE
);

CREATE INDEX idx_ip_licenses_ip_rights ON ip_licenses (ip_rights_id);

-- IP Transfers
CREATE TABLE ip_transfers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ip_rights_id UUID NOT NULL,
    
    -- Transfer Details
    transfer_type VARCHAR(30) CHECK (
        transfer_type IN ('ASSIGNMENT', 'SALE', 'GIFT', 'INHERITANCE')
    ),
    
    -- Parties
    transferor_id UUID NOT NULL,
    transferee_id UUID NOT NULL,
    
    -- Consideration
    consideration_amount DECIMAL(12, 2),
    currency CHAR(3) DEFAULT 'USD',
    
    -- Transfer Date
    transfer_date DATE NOT NULL,
    effective_date DATE,
    
    -- Documentation
    transfer_agreement_url TEXT,
    registered BOOLEAN DEFAULT FALSE,
    registration_date DATE,
    registration_number VARCHAR(100),
    
    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'COMPLETED', 'CANCELLED')
    ),
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_ip_transfers_ip_rights FOREIGN KEY (ip_rights_id) 
        REFERENCES ip_rights(id) ON DELETE CASCADE
);

CREATE INDEX idx_ip_transfers_ip_rights ON ip_transfers (ip_rights_id);

```
=========================================
##  SECTION 30: NON-DISCLOSURE AGREEMENTS (NDA)
```sql
-- Domain: internal/domain/nda/
-- Entity: nda/entity.go
-- =========================================

CREATE TABLE ndas (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- NDA Type
    nda_type VARCHAR(30) DEFAULT 'MUTUAL' CHECK (
        nda_type IN ('UNILATERAL', 'MUTUAL', 'MULTILATERAL')
    ),
    
    -- Effective Period
    effective_from DATE NOT NULL,
    effective_until DATE,
    is_perpetual BOOLEAN DEFAULT FALSE,
    
    -- Scope
    confidential_information_definition TEXT NOT NULL,
    exclusions TEXT[],
    permitted_disclosures TEXT[],
    
    -- Obligations
    use_restrictions TEXT NOT NULL,
    return_destruction_requirements TEXT,
    notification_requirements TEXT,
    
    -- Legal Terms
    governing_law VARCHAR(100),
    jurisdiction VARCHAR(100),
    
    -- Status
    status VARCHAR(20) DEFAULT 'DRAFT' CHECK (
        status IN ('DRAFT', 'PENDING_SIGNATURE', 'ACTIVE', 'EXPIRED', 'TERMINATED', 'BREACHED')
    ),
    
    -- Signatures
    signed_at TIMESTAMPTZ,
    signed_by_all_parties BOOLEAN DEFAULT FALSE,
    
    -- Breach
    breached_at TIMESTAMPTZ,
    breach_description TEXT,
    
    -- Documents
    nda_document_url TEXT,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_ndas_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_ndas_contract ON ndas (contract_id);
CREATE INDEX idx_ndas_status ON ndas (status);

COMMENT ON TABLE ndas IS 'Non-Disclosure Agreements - maps to internal/domain/nda/entity.go';

-- NDA Parties
CREATE TABLE nda_parties (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nda_id UUID NOT NULL,
    
    -- Party Details
    party_user_id UUID NOT NULL,
    party_role VARCHAR(30) CHECK (
        party_role IN ('DISCLOSING', 'RECEIVING', 'BOTH')
    ),
    
    -- Signature
    has_signed BOOLEAN DEFAULT FALSE,
    signed_at TIMESTAMPTZ,
    signature_id UUID, -- Reference to contract_signatures
    
    -- Obligations
    specific_obligations TEXT[],
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_nda_parties_nda FOREIGN KEY (nda_id) 
        REFERENCES ndas(id) ON DELETE CASCADE,
    CONSTRAINT uk_nda_parties UNIQUE (nda_id, party_user_id)
);

CREATE INDEX idx_nda_parties_nda ON nda_parties (nda_id);

-- NDA Breaches
CREATE TABLE nda_breaches (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nda_id UUID NOT NULL,
    
    -- Breach Details
    breach_type VARCHAR(30) CHECK (
        breach_type IN ('UNAUTHORIZED_DISCLOSURE', 'UNAUTHORIZED_USE', 'FAILURE_TO_RETURN', 
                       'FAILURE_TO_NOTIFY', 'OTHER')
    ),
    description TEXT NOT NULL,
    
    -- Parties Involved
    breaching_party_id UUID NOT NULL,
    affected_parties UUID[],
    
    -- Impact
    severity VARCHAR(20) CHECK (
        severity IN ('MINOR', 'MODERATE', 'MAJOR', 'CRITICAL')
    ),
    impact_assessment TEXT,
    estimated_damages DECIMAL(12, 2),
    
    -- Discovery
    discovered_at TIMESTAMPTZ NOT NULL,
    reported_by UUID NOT NULL,
    reported_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Investigation
    investigation_status VARCHAR(20) DEFAULT 'OPEN' CHECK (
        investigation_status IN ('OPEN', 'INVESTIGATING', 'CONFIRMED', 'UNFOUNDED', 'RESOLVED')
    ),
    investigation_notes TEXT,
    
    -- Resolution
    resolution VARCHAR(20) CHECK (
        resolution IN ('SETTLEMENT', 'LITIGATION', 'ARBITRATION', 'MUTUAL_AGREEMENT', 'DISMISSED')
    ),
    resolution_description TEXT,
    resolved_at TIMESTAMPTZ,
    
    -- Remedial Actions
    remedial_actions_taken TEXT[],
    
    CONSTRAINT fk_nda_breaches_nda FOREIGN KEY (nda_id) 
        REFERENCES ndas(id) ON DELETE CASCADE
);

CREATE INDEX idx_nda_breaches_nda ON nda_breaches (nda_id);
CREATE INDEX idx_nda_breaches_status ON nda_breaches (investigation_status);

```
=========================================
##  SECTION 31: CONTRACT NEGOTIATIONS
```sql
-- Domain: internal/domain/negotiation/
-- Entity: negotiation/entity.go
-- =========================================

CREATE TABLE contract_negotiations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL UNIQUE,
    
    -- Negotiation Configuration
    max_rounds INTEGER DEFAULT 5,
    current_round INTEGER DEFAULT 1,
    
    -- Status
    status VARCHAR(20) DEFAULT 'OPEN' CHECK (
        status IN ('OPEN', 'IN_PROGRESS', 'ACCEPTED', 'REJECTED', 'EXPIRED', 'CANCELLED')
    ),
    
    -- Timeline
    started_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at TIMESTAMPTZ,
    concluded_at TIMESTAMPTZ,
    
    -- Parties
    initiator_id UUID NOT NULL,
    respondent_id UUID NOT NULL,
    
    -- Current Terms
    current_offer_id UUID,
    
    -- Outcome
    final_terms JSONB,
    accepted_offer_id UUID,
    
    -- Rejection
    rejection_reason TEXT,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_contract_negotiations_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_contract_negotiations_contract ON contract_negotiations (contract_id);
CREATE INDEX idx_contract_negotiations_status ON contract_negotiations (status);

COMMENT ON TABLE contract_negotiations IS 'Contract negotiations - maps to internal/domain/negotiation/entity.go';

-- Negotiation Offers
CREATE TABLE negotiation_offers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    negotiation_id UUID NOT NULL,
    
    -- Offer Details
    round_number INTEGER NOT NULL,
    offer_type VARCHAR(30) CHECK (
        offer_type IN ('INITIAL', 'COUNTER', 'FINAL')
    ),
    
    -- Proposed Terms
    terms_json JSONB NOT NULL,
    
    -- Key Terms
    proposed_rate DECIMAL(10, 2),
    proposed_budget DECIMAL(15, 2),
    proposed_timeline_days INTEGER,
    proposed_milestones JSONB,
    
    -- Changes from Previous
    changes_summary TEXT,
    changes_json JSONB,
    
    -- Proposer
    proposed_by UUID NOT NULL,
    proposed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Validity
    expires_at TIMESTAMPTZ,
    
    -- Response
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'ACCEPTED', 'COUNTERED', 'REJECTED', 'EXPIRED', 'WITHDRAWN')
    ),
    responded_at TIMESTAMPTZ,
    response_message TEXT,
    
    -- Notes
    proposer_notes TEXT,
    
    CONSTRAINT fk_negotiation_offers_negotiation FOREIGN KEY (negotiation_id) 
        REFERENCES contract_negotiations(id) ON DELETE CASCADE
);

CREATE INDEX idx_negotiation_offers_negotiation ON negotiation_offers (negotiation_id, round_number DESC);
CREATE INDEX idx_negotiation_offers_status ON negotiation_offers (status);

-- Negotiation Counter Offers
CREATE TABLE negotiation_counter_offers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    original_offer_id UUID NOT NULL,
    
    -- Counter Offer
    counter_terms_json JSONB NOT NULL,
    counter_message TEXT,
    
    -- Countered By
    countered_by UUID NOT NULL,
    countered_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_negotiation_counter_offers_offer FOREIGN KEY (original_offer_id) 
        REFERENCES negotiation_offers(id) ON DELETE CASCADE
);

CREATE INDEX idx_negotiation_counter_offers_offer ON negotiation_counter_offers (original_offer_id);

-- Negotiation History
CREATE TABLE negotiation_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    negotiation_id UUID NOT NULL,
    
    -- Event
    event_type VARCHAR(50) NOT NULL,
    event_description TEXT,
    
    -- Actor
    actor_id UUID NOT NULL,
    
    -- Changes
    old_values JSONB,
    new_values JSONB,
    
    occurred_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_negotiation_history_negotiation FOREIGN KEY (negotiation_id) 
        REFERENCES contract_negotiations(id) ON DELETE CASCADE
);

CREATE INDEX idx_negotiation_history_negotiation ON negotiation_history (negotiation_id, occurred_at DESC);




```
=========================================
##  SECTION 32: CONTRACT REPORTS
```sql
-- Domain: internal/domain/report/
-- Entity: report/entity.go
-- =========================================

CREATE TABLE contract_reports (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- Report Details
    report_type VARCHAR(50) NOT NULL CHECK (
        report_type IN ('PERFORMANCE', 'FINANCIAL', 'TIMESHEET_SUMMARY', 'MILESTONE_STATUS', 
                       'BUDGET_ANALYSIS', 'CUSTOM', 'AUDIT', 'COMPLIANCE')
    ),
    report_name VARCHAR(200) NOT NULL,
    
    -- Period
    period_start DATE,
    period_end DATE,
    
    -- Format
    format VARCHAR(20) DEFAULT 'PDF' CHECK (
        format IN ('PDF', 'XLSX', 'CSV', 'JSON', 'HTML')
    ),
    
    -- Filters & Parameters
    filters JSONB,
    parameters JSONB,
    
    -- Generation
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'GENERATING', 'COMPLETED', 'FAILED', 'EXPIRED')
    ),
    
    -- Output
    file_url TEXT,
    file_size BIGINT,
    
    -- Metadata
    total_pages INTEGER,
    data_points_count INTEGER,
    
    -- Timeline
    generated_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ,
    
    -- Error Handling
    error_message TEXT,
    retry_count INTEGER DEFAULT 0,
    
    -- Access
    is_public BOOLEAN DEFAULT FALSE,
    access_count INTEGER DEFAULT 0,
    last_accessed_at TIMESTAMPTZ,
    
    created_by UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_contract_reports_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_contract_reports_contract ON contract_reports (contract_id, created_at DESC);
CREATE INDEX idx_contract_reports_type ON contract_reports (report_type);
CREATE INDEX idx_contract_reports_status ON contract_reports (status, created_at DESC);

COMMENT ON TABLE contract_reports IS 'Contract reports - maps to internal/domain/report/entity.go';

-- Report Schedules
CREATE TABLE report_schedules (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- Schedule Details
    report_type VARCHAR(50) NOT NULL,
    report_name VARCHAR(200) NOT NULL,
    
    -- Cadence
    cadence VARCHAR(20) NOT NULL CHECK (
        cadence IN ('DAILY', 'WEEKLY', 'BIWEEKLY', 'MONTHLY', 'QUARTERLY', 'ANNUALLY')
    ),
    
    -- Schedule
    day_of_week INTEGER CHECK (day_of_week BETWEEN 0 AND 6),
    day_of_month INTEGER CHECK (day_of_month BETWEEN 1 AND 31),
    time_of_day TIME DEFAULT '09:00:00',
    timezone VARCHAR(50) DEFAULT 'UTC',
    
    -- Next Run
    next_run_at TIMESTAMPTZ NOT NULL,
    last_run_at TIMESTAMPTZ,
    
    -- Recipients
    recipients UUID[] NOT NULL,
    
    -- Configuration
    format VARCHAR(20) DEFAULT 'PDF',
    filters JSONB,
    parameters JSONB,
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    is_paused BOOLEAN DEFAULT FALSE,
    
    -- Stats
    run_count INTEGER DEFAULT 0,
    success_count INTEGER DEFAULT 0,
    failure_count INTEGER DEFAULT 0,
    
    created_by UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_report_schedules_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_report_schedules_contract ON report_schedules (contract_id);
CREATE INDEX idx_report_schedules_next_run ON report_schedules (next_run_at, is_active) 
    WHERE is_active = TRUE AND is_paused = FALSE;

-- Report Runs (History)
CREATE TABLE report_runs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    schedule_id UUID NOT NULL,
    report_id UUID,
    
    -- Run Details
    run_type VARCHAR(20) CHECK (
        run_type IN ('SCHEDULED', 'MANUAL')
    ),
    
    -- Status
    status VARCHAR(20) DEFAULT 'RUNNING' CHECK (
        status IN ('RUNNING', 'COMPLETED', 'FAILED')
    ),
    
    -- Timing
    started_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    completed_at TIMESTAMPTZ,
    duration_seconds INTEGER,
    
    -- Error
    error_message TEXT,
    
    CONSTRAINT fk_report_runs_schedule FOREIGN KEY (schedule_id) 
        REFERENCES report_schedules(id) ON DELETE CASCADE,
    CONSTRAINT fk_report_runs_report FOREIGN KEY (report_id) 
        REFERENCES contract_reports(id) ON DELETE SET NULL
);

CREATE INDEX idx_report_runs_schedule ON report_runs (schedule_id, started_at DESC);

```
=========================================
##  SECTION 33: CONTRACT ATTACHMENTS
```sql
-- Domain: internal/domain/attachment/
-- Entity: attachment/entity.go
-- =========================================

CREATE TABLE contract_attachments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- Attachment Type
    attachment_type VARCHAR(50) CHECK (
        attachment_type IN ('CONTRACT_DOCUMENT', 'SUPPORTING_DOCUMENT', 'LEGAL_DOCUMENT', 
                           'REFERENCE', 'SPECIFICATION', 'PROPOSAL', 'OTHER')
    ),
    attachment_category VARCHAR(50),
    
    -- File Details
    file_id UUID NOT NULL, -- Reference to storage-be
    file_url TEXT NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_size BIGINT NOT NULL,
    file_mime_type VARCHAR(100),
    
    -- Versioning
    version INTEGER DEFAULT 1,
    is_current BOOLEAN DEFAULT TRUE,
    previous_version_id UUID,
    
    -- Description
    title VARCHAR(200),
    description TEXT,
    
    -- Security
    is_confidential BOOLEAN DEFAULT FALSE,
    access_level VARCHAR(20) DEFAULT 'CONTRACT_PARTIES' CHECK (
        access_level IN ('PUBLIC', 'CONTRACT_PARTIES', 'CLIENT_ONLY', 'FREELANCER_ONLY', 'ADMIN_ONLY')
    ),
    
    -- Virus Scanning
    virus_scan_status VARCHAR(20) DEFAULT 'PENDING',
    scanned_at TIMESTAMPTZ,
    
    -- Metadata
    tags TEXT[],
    metadata JSONB,
    
    -- Activity
    download_count INTEGER DEFAULT 0,
    last_downloaded_at TIMESTAMPTZ,
    
    -- Status
    is_archived BOOLEAN DEFAULT FALSE,
    archived_at TIMESTAMPTZ,
    
    uploaded_by UUID NOT NULL,
    uploaded_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_contract_attachments_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE,
    CONSTRAINT fk_contract_attachments_previous FOREIGN KEY (previous_version_id) 
        REFERENCES contract_attachments(id)
);

CREATE INDEX idx_contract_attachments_contract ON contract_attachments (contract_id, uploaded_at DESC);
CREATE INDEX idx_contract_attachments_file ON contract_attachments (file_id);
CREATE INDEX idx_contract_attachments_type ON contract_attachments (attachment_type);
CREATE INDEX idx_contract_attachments_current ON contract_attachments (contract_id, is_current) 
    WHERE is_current = TRUE;

COMMENT ON TABLE contract_attachments IS 'Generic contract attachments - maps to internal/domain/attachment/entity.go';

```
=========================================
##  SECTION 34: CONTRACT RENEWALS (EXPLICIT WORKFLOW)
```sql
-- Domain: internal/domain/renewal/
-- Entity: renewal/entity.go
-- =========================================

CREATE TABLE contract_renewals (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- Renewal Type
    renewal_type VARCHAR(30) CHECK (
        renewal_type IN ('STANDARD', 'EXTENSION', 'RENEGOTIATION', 'EARLY_RENEWAL')
    ),
    
    -- Proposed Terms
    proposed_start_date DATE NOT NULL,
    proposed_end_date DATE,
    proposed_duration_days INTEGER,
    proposed_rate DECIMAL(10, 2),
    proposed_budget DECIMAL(15, 2),
    proposed_terms JSONB,
    
    -- Terms Changes
    has_changes BOOLEAN DEFAULT FALSE,
    changes_summary TEXT,
    changes_detail JSONB,
    
    -- Request
    requested_by UUID NOT NULL,
    requested_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    request_reason TEXT,
    
    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'UNDER_REVIEW', 'APPROVED', 'REJECTED', 'CANCELLED', 'EXPIRED')
    ),
    
    -- Approval
    requires_client_approval BOOLEAN DEFAULT TRUE,
    requires_freelancer_approval BOOLEAN DEFAULT TRUE,
    
    approved_by_client BOOLEAN DEFAULT FALSE,
    client_approved_by UUID,
    client_approved_at TIMESTAMPTZ,
    client_approval_notes TEXT,
    
    approved_by_freelancer BOOLEAN DEFAULT FALSE,
    freelancer_approved_by UUID,
    freelancer_approved_at TIMESTAMPTZ,
    freelancer_approval_notes TEXT,
    
    -- Rejection
    rejected_by UUID,
    rejected_at TIMESTAMPTZ,
    rejection_reason TEXT,
    
    -- Result
    new_contract_id UUID, -- Created contract if approved
    
    -- Timeline
    decision_deadline TIMESTAMPTZ,
    expires_at TIMESTAMPTZ,
    
    -- Cancellation
    cancelled_by UUID,
    cancelled_at TIMESTAMPTZ,
    cancellation_reason TEXT,
    
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_contract_renewals_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE,
    CONSTRAINT fk_contract_renewals_new_contract FOREIGN KEY (new_contract_id) 
        REFERENCES contracts(id) ON DELETE SET NULL
);

CREATE INDEX idx_contract_renewals_contract ON contract_renewals (contract_id, requested_at DESC);
CREATE INDEX idx_contract_renewals_status ON contract_renewals (status);
CREATE INDEX idx_contract_renewals_pending ON contract_renewals (status, decision_deadline) 
    WHERE status IN ('PENDING', 'UNDER_REVIEW');

COMMENT ON TABLE contract_renewals IS 'Explicit renewal workflow - maps to internal/domain/renewal/entity.go';

```
=========================================
##  SECTION 35: CONTRACT PAUSES
```sql
-- Domain: internal/domain/pause/
-- Entity: pause/entity.go
-- =========================================

CREATE TABLE contract_pauses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- Pause Details
    pause_type VARCHAR(30) CHECK (
        pause_type IN ('PLANNED', 'EMERGENCY', 'MUTUAL_AGREEMENT', 'CLIENT_REQUEST', 'FREELANCER_REQUEST')
    ),
    reason TEXT NOT NULL,
    
    -- Timeline
    started_at TIMESTAMPTZ NOT NULL,
    planned_end_at TIMESTAMPTZ,
    ended_at TIMESTAMPTZ,
    
    -- Duration
    duration_days INTEGER,
    
    -- Request
    requested_by UUID NOT NULL,
    requested_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Approval
    requires_approval BOOLEAN DEFAULT TRUE,
    approved_by UUID,
    approved_at TIMESTAMPTZ,
    
    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'APPROVED', 'ACTIVE', 'ENDED', 'REJECTED', 'CANCELLED')
    ),
    
    -- Impact
    affects_deadline BOOLEAN DEFAULT TRUE,
    deadline_extension_days INTEGER,
    affects_payment BOOLEAN DEFAULT FALSE,
    
    -- Resume
    resumed_by UUID,
    resume_reason TEXT,
    
    -- Notes
    notes TEXT,
    
    CONSTRAINT fk_contract_pauses_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_contract_pauses_contract ON contract_pauses (contract_id, started_at DESC);
CREATE INDEX idx_contract_pauses_status ON contract_pauses (status);
CREATE INDEX idx_contract_pauses_active ON contract_pauses (status, started_at) 
    WHERE status = 'ACTIVE';

COMMENT ON TABLE contract_pauses IS 'Contract pause history - maps to internal/domain/pause/entity.go';

```
=========================================
##  SECTION 36: WORKROOM (ENHANCED WORKSPACE)
```sql
-- Domain: internal/domain/workroom/
-- Entity: workroom/entity.go
-- =========================================

-- Workroom Tasks
CREATE TABLE workroom_tasks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    workspace_id UUID NOT NULL,
    
    -- Task Details
    title VARCHAR(200) NOT NULL,
    description TEXT,
    
    -- Priority
    priority VARCHAR(20) DEFAULT 'MEDIUM' CHECK (
        priority IN ('LOW', 'MEDIUM', 'HIGH', 'URGENT')
    ),
    
    -- Status
    status VARCHAR(20) DEFAULT 'TODO' CHECK (
        status IN ('TODO', 'IN_PROGRESS', 'IN_REVIEW', 'COMPLETED', 'CANCELLED')
    ),
    
    -- Assignment
    assigned_to UUID[],
    
    -- Timeline
    due_date DATE,
    started_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    
    -- Dependencies
    depends_on_task_ids UUID[],
    blocks_task_ids UUID[],
    
    -- Checklist
    checklist_items JSONB, -- {item: text, completed: boolean}
    
    -- Attachments
    attachment_ids UUID[],
    
    -- Activity
    comment_count INTEGER DEFAULT 0,
    
    created_by UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_workroom_tasks_workspace FOREIGN KEY (workspace_id) 
        REFERENCES contract_workspaces(id) ON DELETE CASCADE
);

CREATE INDEX idx_workroom_tasks_workspace ON workroom_tasks (workspace_id);
CREATE INDEX idx_workroom_tasks_status ON workroom_tasks (status, due_date);
CREATE INDEX idx_workroom_tasks_assigned ON workroom_tasks USING gin(assigned_to);

COMMENT ON TABLE workroom_tasks IS 'Workroom tasks - maps to internal/domain/workroom/tasks.go';

-- Workroom Task Assignees
CREATE TABLE workroom_task_assignees (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    task_id UUID NOT NULL,
    user_id UUID NOT NULL,
    
    -- Assignment
    assigned_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    assigned_by UUID NOT NULL,
    
    -- Status
    acceptance_status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        acceptance_status IN ('PENDING', 'ACCEPTED', 'DECLINED')
    ),
    responded_at TIMESTAMPTZ,
    
    CONSTRAINT fk_workroom_task_assignees_task FOREIGN KEY (task_id) 
        REFERENCES workroom_tasks(id) ON DELETE CASCADE,
    CONSTRAINT uk_workroom_task_assignees UNIQUE (task_id, user_id)
);

CREATE INDEX idx_workroom_task_assignees_task ON workroom_task_assignees (task_id);
CREATE INDEX idx_workroom_task_assignees_user ON workroom_task_assignees (user_id);

-- Workroom Notes
CREATE TABLE workroom_notes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    workspace_id UUID NOT NULL,
    
    -- Note Details
    title VARCHAR(200),
    content TEXT NOT NULL,
    
    -- Format
    content_format VARCHAR(20) DEFAULT 'MARKDOWN' CHECK (
        content_format IN ('PLAIN', 'MARKDOWN', 'HTML')
    ),
    
    -- Category
    category VARCHAR(50),
    tags TEXT[],
    
    -- Pinning
    is_pinned BOOLEAN DEFAULT FALSE,
    pinned_at TIMESTAMPTZ,
    
    -- Visibility
    is_private BOOLEAN DEFAULT FALSE,
    visible_to UUID[], -- If private, who can see it
    
    -- Attachments
    attachment_ids UUID[],
    
    -- Activity
    view_count INTEGER DEFAULT 0,
    last_viewed_at TIMESTAMPTZ,
    
    created_by UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_workroom_notes_workspace FOREIGN KEY (workspace_id) 
        REFERENCES contract_workspaces(id) ON DELETE CASCADE
);

CREATE INDEX idx_workroom_notes_workspace ON workroom_notes (workspace_id, created_at DESC);
CREATE INDEX idx_workroom_notes_pinned ON workroom_notes (workspace_id, is_pinned) 
    WHERE is_pinned = TRUE;

COMMENT ON TABLE workroom_notes IS 'Workroom notes - maps to internal/domain/workroom/notes.go';

```
=========================================
##  SECTION 37: PERFORMANCE TRACKING (ENHANCED)
```sql
-- Domain: internal/domain/performance/
-- Entity: performance/entity.go
-- =========================================

-- Performance KPIs
CREATE TABLE performance_kpis (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- KPI Definition
    kpi_name VARCHAR(100) NOT NULL,
    kpi_type VARCHAR(50) NOT NULL,
    kpi_description TEXT,
    
    -- Target
    target_value DECIMAL(12, 2) NOT NULL,
    unit VARCHAR(50) NOT NULL,
    
    -- Measurement
    measurement_frequency VARCHAR(20) CHECK (
        measurement_frequency IN ('DAILY', 'WEEKLY', 'MONTHLY', 'QUARTERLY', 'MILESTONE')
    ),
    
    -- Thresholds
    excellent_threshold DECIMAL(12, 2),
    good_threshold DECIMAL(12, 2),
    acceptable_threshold DECIMAL(12, 2),
    
    -- Weight (for overall score)
    weight DECIMAL(5, 2) DEFAULT 1.00,
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_performance_kpis_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_performance_kpis_contract ON performance_kpis (contract_id, is_active);

COMMENT ON TABLE performance_kpis IS 'Performance KPIs - maps to internal/domain/performance/kpis.go';

-- Performance Records
CREATE TABLE performance_records (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kpi_id UUID NOT NULL,
    contract_id UUID NOT NULL,
    
    -- Measurement Period
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    
    -- Measured Value
    measured_value DECIMAL(12, 2) NOT NULL,
    target_value DECIMAL(12, 2) NOT NULL,
    
    -- Performance
    achievement_percentage DECIMAL(5, 2),
    performance_level VARCHAR(20) CHECK (
        performance_level IN ('EXCELLENT', 'GOOD', 'ACCEPTABLE', 'BELOW_TARGET', 'POOR')
    ),
    
    -- Context
    measurement_method VARCHAR(100),
    sample_size INTEGER,
    notes TEXT,
    
    recorded_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    recorded_by UUID,
    
    CONSTRAINT fk_performance_records_kpi FOREIGN KEY (kpi_id) 
        REFERENCES performance_kpis(id) ON DELETE CASCADE,
    CONSTRAINT fk_performance_records_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_performance_records_kpi ON performance_records (kpi_id, period_start DESC);
CREATE INDEX idx_performance_records_contract ON performance_records (contract_id, period_start DESC);

COMMENT ON TABLE performance_records IS 'Performance measurements - maps to internal/domain/performance/records.go';

-- Performance Scores
CREATE TABLE performance_scores (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- Score Period
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    
    -- Overall Score
    overall_score DECIMAL(5, 2) NOT NULL CHECK (overall_score BETWEEN 0 AND 100),
    grade VARCHAR(2) CHECK (grade IN ('A+', 'A', 'A-', 'B+', 'B', 'B-', 'C+', 'C', 'C-', 'D', 'F')),
    
    -- Component Scores
    quality_score DECIMAL(5, 2),
    timeliness_score DECIMAL(5, 2),
    communication_score DECIMAL(5, 2),
    professionalism_score DECIMAL(5, 2),
    
    -- KPI Breakdown
    kpi_scores JSONB, -- {kpi_id: score}
    
    -- Trend
    trend VARCHAR(20) CHECK (
        trend IN ('IMPROVING', 'STABLE', 'DECLINING')
    ),
    change_from_previous DECIMAL(5, 2),
    
    -- Context
    total_kpis_measured INTEGER,
    kpis_above_target INTEGER,
    kpis_below_target INTEGER,
    
    calculated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_performance_scores_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE,
    CONSTRAINT uk_performance_scores UNIQUE (contract_id, period_start, period_end)
);

CREATE INDEX idx_performance_scores_contract ON performance_scores (contract_id, period_start DESC);
CREATE INDEX idx_performance_scores_overall ON performance_scores (overall_score DESC);

COMMENT ON TABLE performance_scores IS 'Aggregated performance scores - maps to internal/domain/performance/scores.go';

-- Performance Benchmarks
CREATE TABLE performance_benchmarks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Benchmark Context
    contract_type VARCHAR(30),
    industry VARCHAR(100),
    skill_category VARCHAR(100),
    
    -- KPI
    kpi_type VARCHAR(50) NOT NULL,
    
    -- Benchmark Values
    percentile_10 DECIMAL(12, 2),
    percentile_25 DECIMAL(12, 2),
    percentile_50 DECIMAL(12, 2), -- Median
    percentile_75 DECIMAL(12, 2),
    percentile_90 DECIMAL(12, 2),
    average_value DECIMAL(12, 2),
    
    -- Sample
    sample_size INTEGER NOT NULL,
    
    -- Period
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    
    -- Status
    is_current BOOLEAN DEFAULT TRUE,
    
    calculated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT chk_performance_benchmarks_sample CHECK (sample_size > 0)
);

CREATE INDEX idx_performance_benchmarks_type ON performance_benchmarks (kpi_type, is_current);
CREATE INDEX idx_performance_benchmarks_context ON performance_benchmarks (contract_type, industry);

COMMENT ON TABLE performance_benchmarks IS 'Industry benchmarks - maps to internal/domain/performance/benchmarks.go';

```
=========================================
##  SECTION 38: DIRECT CONTRACTS (ENHANCED)
```sql
-- Domain: internal/domain/direct_contract/
-- Entity: direct_contract/entity.go
-- =========================================

CREATE TABLE direct_contracts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    invitation_id UUID UNIQUE, -- Link to contract_invitations
    contract_id UUID UNIQUE, -- Created contract
    
    -- Parties
    client_id UUID NOT NULL,
    freelancer_id UUID NOT NULL,
    
    -- Flow Status
    flow_status VARCHAR(20) DEFAULT 'INITIATED' CHECK (
        flow_status IN ('INITIATED', 'INVITATION_SENT', 'VIEWED', 'NEGOTIATING', 
                       'TERMS_AGREED', 'SIGNING', 'ACTIVATED', 'DECLINED', 'EXPIRED', 'CANCELLED')
    ),
    
    -- Invitation Token
    invitation_token VARCHAR(100) UNIQUE,
    token_expires_at TIMESTAMPTZ,
    
    -- Terms
    proposed_terms JSONB NOT NULL,
    agreed_terms JSONB,
    
    -- Timeline
    initiated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    invitation_sent_at TIMESTAMPTZ,
    viewed_at TIMESTAMPTZ,
    terms_agreed_at TIMESTAMPTZ,
    signing_started_at TIMESTAMPTZ,
    activated_at TIMESTAMPTZ,
    
    -- Decline/Cancel
    declined_at TIMESTAMPTZ,
    decline_reason TEXT,
    cancelled_at TIMESTAMPTZ,
    cancellation_reason TEXT,
    
    -- Metadata
    source VARCHAR(50), -- REFERRAL, NETWORK, DIRECT
    metadata JSONB,
    
    CONSTRAINT fk_direct_contracts_invitation FOREIGN KEY (invitation_id) 
        REFERENCES contract_invitations(id) ON DELETE SET NULL,
    CONSTRAINT fk_direct_contracts_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE SET NULL
);

CREATE INDEX idx_direct_contracts_client ON direct_contracts (client_id, flow_status);
CREATE INDEX idx_direct_contracts_freelancer ON direct_contracts (freelancer_id, flow_status);
CREATE INDEX idx_direct_contracts_token ON direct_contracts (invitation_token) 
    WHERE invitation_token IS NOT NULL;

COMMENT ON TABLE direct_contracts IS 'Direct contract flow tracking - maps to internal/domain/direct_contract/entity.go';

```
=========================================
##  SECTION 39: OUTBOX PATTERN FOR EVENTS
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
    
    -- Timestamps
    occurred_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    published_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_outbox_events_status ON outbox_events (status, next_attempt_at) 
    WHERE status IN ('PENDING', 'FAILED');
CREATE INDEX idx_outbox_events_aggregate ON outbox_events (aggregate_type, aggregate_id);
CREATE INDEX idx_outbox_events_correlation ON outbox_events (correlation_id);
CREATE INDEX idx_outbox_events_topic ON outbox_events (topic, occurred_at DESC);

COMMENT ON TABLE outbox_events IS 'Transactional outbox for event publishing';

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
##  SECTION 40: READ MODELS (CQRS PROJECTIONS)
=========================================

```sql

-- Contract Read Model for Fast Queries
CREATE TABLE contract_read_model (
    contract_id UUID PRIMARY KEY,
    
    -- Core Info
    contract_number VARCHAR(50),
    title VARCHAR(300),
    contract_type VARCHAR(30),
    status VARCHAR(20),
    
    -- Participants (denormalized)
    client_id UUID,
    client_name VARCHAR(200),
    freelancer_id UUID,
    freelancer_name VARCHAR(200),
    
    -- Financial
    total_value DECIMAL(15, 2),
    total_paid DECIMAL(15, 2),
    currency CHAR(3),
    
    -- Timeline
    start_date DATE,
    end_date DATE,
    activated_at TIMESTAMPTZ,
    
    -- Performance
    quality_score DECIMAL(5, 2),
    performance_rating DECIMAL(3, 2),
    completion_percentage DECIMAL(5, 2),
    
    -- Status Flags
    is_active BOOLEAN,
    is_paused BOOLEAN,
    is_disputed BOOLEAN,
    is_overbudget BOOLEAN,
    has_open_milestones BOOLEAN,
    
    -- Counts
    milestone_count INTEGER DEFAULT 0,
    completed_milestone_count INTEGER DEFAULT 0,
    dispute_count INTEGER DEFAULT 0,
    
    -- Search
    search_vector tsvector GENERATED ALWAYS AS (
        setweight(to_tsvector('english', COALESCE(contract_number, '')), 'A') ||
        setweight(to_tsvector('english', COALESCE(title, '')), 'B') ||
        setweight(to_tsvector('english', COALESCE(client_name, '')), 'C') ||
        setweight(to_tsvector('english', COALESCE(freelancer_name, '')), 'C')
    ) STORED,
    
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_contract_read_search ON contract_read_model USING gin(search_vector);
CREATE INDEX idx_contract_read_client ON contract_read_model (client_id, status);
CREATE INDEX idx_contract_read_freelancer ON contract_read_model (freelancer_id, status);
CREATE INDEX idx_contract_read_status ON contract_read_model (status, start_date DESC);
CREATE INDEX idx_contract_read_active ON contract_read_model (is_active) WHERE is_active = TRUE;

-- User Contract Stats (Projection)
CREATE TABLE user_contract_stats (
    user_id UUID PRIMARY KEY,
    user_type VARCHAR(20), -- CLIENT, FREELANCER
    
    -- Overall Stats
    total_contracts INTEGER DEFAULT 0,
    active_contracts INTEGER DEFAULT 0,
    completed_contracts INTEGER DEFAULT 0,
    terminated_contracts INTEGER DEFAULT 0,
    
    -- Financial
    total_contract_value DECIMAL(15, 2) DEFAULT 0,
    total_paid DECIMAL(15, 2) DEFAULT 0,
    total_earned DECIMAL(15, 2) DEFAULT 0,
    
    -- Performance
    avg_quality_score DECIMAL(5, 2),
    avg_performance_rating DECIMAL(3, 2),
    on_time_completion_rate DECIMAL(5, 2),
    
    -- Disputes
    total_disputes INTEGER DEFAULT 0,
    disputes_won INTEGER DEFAULT 0,
    disputes_lost INTEGER DEFAULT 0,
    
    -- Recent Activity
    last_contract_at TIMESTAMPTZ,
    contracts_last_30_days INTEGER DEFAULT 0,
    
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_user_contract_stats_performance ON user_contract_stats (avg_quality_score DESC);

```
=========================================
##  SECTION 41: EXTERNAL REFERENCES

```sql

-- (Relations with other microservices)
-- =========================================

CREATE TABLE external_references (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL,
    
    -- External Service
    service_name VARCHAR(50) NOT NULL,
    reference_type VARCHAR(50) NOT NULL,
    reference_id UUID NOT NULL,
    
    -- Context
    context JSONB,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_external_references_contract FOREIGN KEY (contract_id) 
        REFERENCES contracts(id) ON DELETE CASCADE
);

CREATE INDEX idx_external_references_contract ON external_references (contract_id);
CREATE INDEX idx_external_references_service ON external_references (service_name, reference_type, reference_id);

COMMENT ON TABLE external_references IS 'References to entities in other microservices';

```
=========================================
##  SECTION 42: DATABASE FUNCTIONS & TRIGGERS


```sql

-- =========================================

-- Function to update updated_at timestamp
CREATE FUNCTION update_updated_at()
RETURNS TRIGGER AS $
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$ LANGUAGE plpgsql;

-- Trigger for contracts table
CREATE TRIGGER trg_contracts_updated_at
    BEFORE UPDATE ON contracts
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();

-- Function to update contract financial totals
CREATE FUNCTION update_contract_financial_totals()
RETURNS TRIGGER AS $
BEGIN
    IF TG_OP = 'INSERT' OR TG_OP = 'UPDATE' THEN
        UPDATE contracts
        SET total_paid = (
            SELECT COALESCE(SUM(amount), 0)
            FROM milestones
            WHERE contract_id = NEW.contract_id 
            AND status = 'PAID'
        ),
        total_pending = (
            SELECT COALESCE(SUM(amount), 0)
            FROM milestones
            WHERE contract_id = NEW.contract_id 
            AND status IN ('SUBMITTED', 'UNDER_REVIEW', 'APPROVED')
        )
        WHERE id = NEW.contract_id;
    END IF;
    RETURN NEW;
END;
$ LANGUAGE plpgsql;

CREATE TRIGGER trg_milestone_financial_update
    AFTER INSERT OR UPDATE OF status, amount ON milestones
    FOR EACH ROW
    EXECUTE FUNCTION update_contract_financial_totals();

```
=========================================
##  SECTION 43: PERFORMANCE VIEWS

```sql

-- =========================================

-- View for active contracts with key metrics
CREATE VIEW v_active_contracts AS
SELECT 
    c.id,
    c.contract_number,
    c.title,
    c.status,
    c.contract_type,
    c.client_id,
    c.freelancer_id,
    c.total_contract_value,
    c.currency,
    c.start_date,
    c.end_date,
    cb.budget_consumed_percentage,
    cb.is_over_budget,
    COUNT(DISTINCT m.id) AS total_milestones,
    COUNT(DISTINCT CASE WHEN m.status = 'COMPLETED' THEN m.id END) AS completed_milestones,
    COUNT(DISTINCT d.id) AS dispute_count,
    ca.quality_score,
    ca.on_time_delivery_rate,
    c.activated_at,
    c.updated_at
FROM contracts c
LEFT JOIN contract_budgets cb ON c.id = cb.contract_id
LEFT JOIN milestones m ON c.id = m.contract_id
LEFT JOIN disputes d ON c.id = d.contract_id AND d.status IN ('OPEN', 'UNDER_REVIEW')
LEFT JOIN contract_analytics ca ON c.id = ca.contract_id
WHERE c.status = 'ACTIVE' AND c.is_deleted = FALSE
GROUP BY c.id, cb.budget_consumed_percentage, cb.is_over_budget, ca.quality_score, ca.on_time_delivery_rate;

-- View for contracts requiring attention
CREATE VIEW v_contracts_requiring_attention AS
SELECT 
    c.id AS contract_id,
    c.contract_number,
    c.title,
    c.status,
    'OVERDUE_MILESTONE' AS attention_type,
    m.title AS attention_subject,
    m.due_date AS attention_date,
    'HIGH' AS priority
FROM contracts c
INNER JOIN milestones m ON c.id = m.contract_id
WHERE c.status = 'ACTIVE' 
    AND m.status IN ('PENDING', 'IN_PROGRESS')
    AND m.due_date < CURRENT_DATE
    AND c.is_deleted = FALSE

UNION ALL

SELECT 
    c.id AS contract_id,
    c.contract_number,
    c.title,
    c.status,
    'BUDGET_WARNING' AS attention_type,
    'Budget threshold exceeded' AS attention_subject,
    CURRENT_TIMESTAMP AS attention_date,
    'MEDIUM' AS priority
FROM contracts c
INNER JOIN contract_budgets cb ON c.id = cb.contract_id
WHERE c.status = 'ACTIVE'
    AND cb.warning_triggered = TRUE
    AND c.is_deleted = FALSE

UNION ALL

SELECT 
    c.id AS contract_id,
    c.contract_number,
    c.title,
    c.status,
    'PENDING_APPROVAL' AS attention_type,
    'Timesheet pending approval' AS attention_subject,
    t.submitted_at AS attention_date,
    'MEDIUM' AS priority
FROM contracts c
INNER JOIN timesheets t ON c.id = t.contract_id
WHERE t.status = 'SUBMITTED'
    AND c.is_deleted = FALSE

UNION ALL

SELECT 
    c.id AS contract_id,
    c.contract_number,
    c.title,
    c.status,
    'OPEN_DISPUTE' AS attention_type,
    d.title AS attention_subject,
    d.opened_at AS attention_date,
    'URGENT' AS priority
FROM contracts c
INNER JOIN disputes d ON c.id = d.contract_id
WHERE d.status IN ('OPEN', 'UNDER_REVIEW')
    AND c.is_deleted = FALSE

UNION ALL

SELECT 
    c.id AS contract_id,
    c.contract_number,
    c.title,
    c.status,
    'EXPIRING_SOON' AS attention_type,
    'Contract expiring soon' AS attention_subject,
    c.end_date AS attention_date,
    'MEDIUM' AS priority
FROM contracts c
WHERE c.status = 'ACTIVE'
    AND c.end_date IS NOT NULL
    AND c.end_date <= CURRENT_DATE + INTERVAL '7 days'
    AND c.is_deleted = FALSE;

```
=========================================
##  SECTION 44: TABLE COMMENTS



-- =========================================

```sql
COMMENT ON TABLE contracts IS 'Core contracts - maps to internal/domain/contract/entity.go';
COMMENT ON TABLE contract_state_history IS 'Contract state transitions';
COMMENT ON TABLE statements_of_work IS 'Statements of Work - maps to internal/domain/sow/entity.go';
COMMENT ON TABLE financial_holds IS 'Financial holds - maps to internal/domain/hold/entity.go';
COMMENT ON TABLE milestones IS 'Contract milestones - maps to internal/domain/milestone/entity.go';
COMMENT ON TABLE milestone_activities IS 'Milestone activity log';
COMMENT ON TABLE deliverables IS 'Milestone deliverables - maps to internal/domain/deliverable/entity.go';
COMMENT ON TABLE timesheets IS 'Weekly timesheets - maps to internal/domain/timesheet/entity.go';
COMMENT ON TABLE time_entries IS 'Daily time entries';
COMMENT ON TABLE work_diary_entries IS 'Work diary tracking - maps to internal/domain/work_diary/entity.go';
COMMENT ON TABLE work_diary_daily_summaries IS 'Daily work diary aggregates';
COMMENT ON TABLE contract_templates IS 'Contract templates - maps to internal/domain/template/entity.go';
COMMENT ON TABLE contract_amendments IS 'Contract amendments - maps to internal/domain/amendment/entity.go';
COMMENT ON TABLE disputes IS 'Contract disputes - maps to internal/domain/dispute/entity.go';
COMMENT ON TABLE dispute_evidence IS 'Dispute evidence attachments';
COMMENT ON TABLE dispute_messages IS 'Dispute communication thread';
COMMENT ON TABLE contract_budgets IS 'Budget tracking - maps to internal/domain/budget/entity.go';
COMMENT ON TABLE budget_adjustments IS 'Budget adjustment history';
COMMENT ON TABLE contract_reminders IS 'Contract reminders - maps to internal/domain/reminder/entity.go';
COMMENT ON TABLE contract_invitations IS 'Direct contract invitations - maps to internal/domain/invitation/entity.go';
COMMENT ON TABLE contract_rate_cards IS 'Rate cards - maps to internal/domain/rate_card/entity.go';
COMMENT ON TABLE contract_analytics IS 'Contract analytics - maps to internal/domain/analytics/entity.go';
COMMENT ON TABLE contract_compliance IS 'Compliance tracking - maps to internal/domain/compliance/entity.go';
COMMENT ON TABLE contract_feedback IS 'Contract feedback - maps to internal/domain/feedback/entity.go';
COMMENT ON TABLE contract_insurance IS 'Contract insurance - maps to internal/domain/insurance/entity.go';
COMMENT ON TABLE insurance_claims IS 'Insurance claims';
COMMENT ON TABLE contract_escrow IS 'Escrow management - maps to internal/domain/escrow/entity.go';
COMMENT ON TABLE escrow_transactions IS 'Escrow transaction ledger';
COMMENT ON TABLE contract_terminations IS 'Contract terminations - maps to internal/domain/termination/entity.go';
COMMENT ON TABLE contract_workspaces IS 'Collaboration workspaces - maps to internal/domain/workspace/entity.go';
COMMENT ON TABLE workspace_documents IS 'Workspace documents';
COMMENT ON TABLE workspace_comments IS 'Document comments';
COMMENT ON TABLE recurring_contracts IS 'Recurring contracts - maps to internal/domain/recurring/entity.go';
COMMENT ON TABLE contract_renewal_history IS 'Renewal history';
COMMENT ON TABLE contract_signatures IS 'E-signatures - maps to internal/domain/signature/entity.go';

```
=========================================
##  SECTION 45: DATABASE STATISTICS
```sql

-- =========================================

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

-- Database health check view
CREATE VIEW v_database_health AS
SELECT
    'Active Contracts' AS metric,
    COUNT(*) AS count
FROM contracts
WHERE status = 'ACTIVE' AND is_deleted = FALSE
UNION ALL
SELECT
    'Pending Approvals',
    COUNT(*)
FROM milestones
WHERE status IN ('SUBMITTED', 'UNDER_REVIEW')
UNION ALL
SELECT
    'Open Disputes',
    COUNT(*)
FROM disputes
WHERE status IN ('OPEN', 'UNDER_REVIEW')
UNION ALL
SELECT
    'Overdue Milestones',
    COUNT(*)
FROM milestones
WHERE status IN ('PENDING', 'IN_PROGRESS')
    AND due_date < CURRENT_DATE
UNION ALL
SELECT
    'Contracts Over Budget',
    COUNT(*)
FROM contract_budgets
WHERE is_over_budget = TRUE
UNION ALL
SELECT
    'Pending Signatures',
    COUNT(*)
FROM contract_signatures
WHERE status IN ('PENDING', 'SENT');
```
=========================================
## END OF CONTRACTS-BE DATABASE DESIGN
=========================================

## FINAL SUMMARY:
- Total Tables: 70+
- Total Indexes: 200+
- Total Domains Covered: 26 (all from contracts-be folder structure)
- Coverage: 100% of contracts-be folder structure
- Production ready for millions of contracts
- Full event sourcing with outbox pattern
- CQRS with read models
- Complete audit trails
- Financial holds & escrow management
- Work diary & time tracking
- Dispute resolution system
- E-signature integration
- Compliance tracking
- Multi-language support ready
- Enterprise-scale performance optimization

## ALIGNMENT WITH FOLDER STRUCTURE:
- ✅ contract/ → contracts table
- ✅ sow/ → statements_of_work table
- ✅ hold/ → financial_holds table
- ✅ milestone/ → milestones table
- ✅ deliverable/ → deliverables table
- ✅ timesheet/ → timesheets, time_entries tables
- ✅ work_diary/ → work_diary_entries, work_diary_daily_summaries tables
- ✅ template/ → contract_templates table
- ✅ amendment/ → contract_amendments table
- ✅ dispute/ → disputes, dispute_evidence, dispute_messages tables
- ✅ budget/ → contract_budgets, budget_adjustments tables
- ✅ reminder/ → contract_reminders table
- ✅ audit/ → contract_audit_logs table
- ✅ invitation/ → contract_invitations table
- ✅ rate_card/ → contract_rate_cards table
- ✅ analytics/ → contract_analytics table
- ✅ compliance/ → contract_compliance table
- ✅ feedback/ → contract_feedback table
- ✅ search/ → contract_search_index table
- ✅ insurance/ → contract_insurance, insurance_claims tables
- ✅ escrow/ → contract_escrow, escrow_transactions tables
- ✅ termination/ → contract_terminations table
- ✅ workspace/ → contract_workspaces, workspace_documents, workspace_comments tables
- ✅ recurring/ → recurring_contracts, contract_renewal_history tables
- ✅ signature/ → contract_signatures table
- ✅ outbox/ → outbox_events, outbox_dead_letter tables
 
## ADDITIONAL FEATURES:
- ✅ Contract state history tracking
- ✅ Milestone activities log
- ✅ Work diary daily summaries for analytics
- ✅ Budget adjustment history
- ✅ Contract read models for CQRS
- ✅ User contract statistics projections
- ✅ External service references
- ✅ Comprehensive views for dashboards
- ✅ Database health monitoring
- ✅ Automated triggers for financial calculations

## INTEGRATION POINTS:
- financial-be: Escrow, payments, holds
- users-be: Client and freelancer references
- jobs-be: Job references (for contracts from jobs)
- proposals-be: Proposal references (for contracts from proposals)
- reviews-be: Feedback integration
- communications-be: Notifications and reminders
- storage-be: File and document storage
- admin-be: Dispute resolution and compliance

## PERFORMANCE OPTIMIZATIONS:
- Partial indexes on frequently filtered columns
- GIN indexes for full-text search
- Materialized views for complex aggregations (can be added)
- Composite indexes for common query patterns
- JSONB indexes for flexible metadata queries
- Partitioning strategy ready (by created_at for contracts)

## COMPLIANCE & LEGAL:
- 10-year data retention support
- GDPR/CCPA compliance fields
- Audit trail for all changes
- Legal review tracking
- Tax compliance fields
- Export control and sanctions screening
- Data protection impact assessments

## FINANCIAL SAFEGUARDS:
- Escrow account integration
- Financial holds for disputes
- Budget tracking with alerts
- Payment reconciliation
- Milestone-based payment releases
- Multi-currency support

All domains from the contracts-be folder structure are fully covered!



```sql

CREATE TRIGGER trg_milestones_updated_at
    BEFORE UPDATE ON milestones
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();



CREATE TRIGGER trg_deliverables_updated_at
    BEFORE UPDATE ON deliverables
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();



CREATE TRIGGER trg_timesheets_updated_at
    BEFORE UPDATE ON timesheets
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();



CREATE TRIGGER trg_slas_updated_at
    BEFORE UPDATE ON slas
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();



CREATE TRIGGER trg_contract_renewals_updated_at
    BEFORE UPDATE ON contract_renewals
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();



CREATE TRIGGER trg_contract_workspaces_updated_at
    BEFORE UPDATE ON contract_workspaces
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();



CREATE TRIGGER trg_workroom_tasks_updated_at
    BEFORE UPDATE ON workroom_tasks
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();



CREATE TRIGGER trg_workroom_notes_updated_at
    BEFORE UPDATE ON workroom_notes
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();



CREATE TRIGGER trg_agency_contracts_updated_at
    BEFORE UPDATE ON agency_contracts
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();



CREATE TRIGGER trg_performance_kpis_updated_at
    BEFORE UPDATE ON performance_kpis
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();

```