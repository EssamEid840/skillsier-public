-- USERS-BE DATABASE DESIGN
-- Skillsier Platform - Enterprise Scale
-- PostgreSQL 16+
-- =========================================
-- 
-- Design Principles:
-- 1. Optimized for large-scale (millions of users)
-- 2. Event-driven architecture with outbox pattern
-- 3. CQRS separation (write models + read projections)
-- 4. PII encryption at rest
-- 5. Soft deletes with audit trails
-- 6. Multi-tenant ready with proper indexing
-- 7. Supports both web and mobile applications
-- 8. Relations with other microservices via IDs only
-- =========================================

-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS "pg_trgm"; -- For full-text search
CREATE EXTENSION IF NOT EXISTS "btree_gin"; -- For composite indexes

-- =========================================
-- SECTION 1: CORE USER DOMAIN
-- =========================================

-- 1.1 Users Table (Core Identity)
CREATE TABLE users (
    -- Primary Key
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Keycloak Integration
    keycloak_id UUID NOT NULL UNIQUE,
    keycloak_created_at TIMESTAMPTZ,
    
    -- Basic Identity (PII - Encrypted)
    email VARCHAR(255) NOT NULL UNIQUE,
    email_verified BOOLEAN DEFAULT FALSE,
    email_verification_token VARCHAR(255),
    email_verification_sent_at TIMESTAMPTZ,
    phone VARCHAR(50),
    phone_verified BOOLEAN DEFAULT FALSE,
    phone_verification_code VARCHAR(10),
    phone_verification_sent_at TIMESTAMPTZ,
    
    -- Personal Info (PII)
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100),
    display_name VARCHAR(200),
    
    -- User Type & Status
    user_type VARCHAR(20) NOT NULL CHECK (user_type IN ('CLIENT', 'FREELANCER', 'HYBRID', 'ADMIN', 'AGENCY')),
    account_status VARCHAR(20) DEFAULT 'PENDING_VERIFICATION' CHECK (
        account_status IN ('PENDING_VERIFICATION', 'ACTIVE', 'SUSPENDED', 'BANNED', 'DEACTIVATED', 'DELETED')
    ),
    suspension_reason TEXT,
    suspension_expires_at TIMESTAMPTZ,
    banned_reason TEXT,
    banned_at TIMESTAMPTZ,
    banned_by UUID, -- admin user_id
    
    -- Profile Completion
    profile_completion_score INTEGER DEFAULT 0 CHECK (profile_completion_score BETWEEN 0 AND 100),
    onboarding_completed BOOLEAN DEFAULT FALSE,
    onboarding_step INTEGER DEFAULT 0,
    
    -- Security & Privacy
    two_factor_enabled BOOLEAN DEFAULT FALSE,
    two_factor_method VARCHAR(20), -- 'SMS', 'EMAIL', 'AUTHENTICATOR'
    security_questions_set BOOLEAN DEFAULT FALSE,
    last_password_change TIMESTAMPTZ,
    password_reset_token VARCHAR(255),
    password_reset_expires_at TIMESTAMPTZ,
    
    -- Activity Tracking
    last_login_at TIMESTAMPTZ,
    last_login_ip INET,
    last_active_at TIMESTAMPTZ,
    login_count INTEGER DEFAULT 0,
    failed_login_attempts INTEGER DEFAULT 0,
    account_locked_until TIMESTAMPTZ,
    
    -- Platform Metadata
    registration_source VARCHAR(50), -- 'WEB', 'MOBILE_IOS', 'MOBILE_ANDROID', 'API', 'SOCIAL_GOOGLE', 'SOCIAL_LINKEDIN'
    referral_code VARCHAR(50) UNIQUE,
    referred_by UUID, -- user_id of referrer
    
    -- Soft Delete & Audit
    is_deleted BOOLEAN DEFAULT FALSE,
    deleted_at TIMESTAMPTZ,
    deleted_by UUID,
    deletion_reason TEXT,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    version INTEGER DEFAULT 1 NOT NULL, -- Optimistic locking
    
    -- Constraints
    CONSTRAINT fk_users_referred_by FOREIGN KEY (referred_by) REFERENCES users(user_id),
    CONSTRAINT fk_users_banned_by FOREIGN KEY (banned_by) REFERENCES users(user_id),
    CONSTRAINT fk_users_deleted_by FOREIGN KEY (deleted_by) REFERENCES users(user_id)
);

-- Indexes for users table
CREATE INDEX idx_users_email ON users USING btree (email) WHERE is_deleted = FALSE;
CREATE INDEX idx_users_keycloak_id ON users USING btree (keycloak_id);
CREATE INDEX idx_users_user_type ON users USING btree (user_type) WHERE is_deleted = FALSE;
CREATE INDEX idx_users_account_status ON users USING btree (account_status) WHERE is_deleted = FALSE;
CREATE INDEX idx_users_created_at ON users USING btree (created_at DESC);
CREATE INDEX idx_users_last_active ON users USING btree (last_active_at DESC) WHERE is_deleted = FALSE;
CREATE INDEX idx_users_referral_code ON users USING btree (referral_code) WHERE referral_code IS NOT NULL;
CREATE INDEX idx_users_search ON users USING gin (
    to_tsvector('english', COALESCE(first_name, '') || ' ' || COALESCE(last_name, '') || ' ' || COALESCE(email, ''))
);

-- 1.2 User Profiles (Extended Information)
CREATE TABLE user_profiles (
    profile_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL UNIQUE,
    
    -- Professional Info
    professional_title VARCHAR(200), -- "Senior Full-Stack Developer", "UI/UX Designer"
    tagline VARCHAR(300), -- Short professional tagline
    bio TEXT, -- Long-form biography
    bio_headline VARCHAR(500), -- Short bio for listings
    
    -- Location & Demographics
    country_code CHAR(2), -- ISO 3166-1 alpha-2
    city VARCHAR(100),
    state_province VARCHAR(100),
    postal_code VARCHAR(20),
    timezone VARCHAR(50), -- IANA timezone (e.g., 'America/New_York')
    
    -- Contact Preferences
    preferred_language VARCHAR(10) DEFAULT 'en', -- ISO 639-1 code
    preferred_currency CHAR(3) DEFAULT 'USD', -- ISO 4217 code
    
    -- Media
    profile_picture_url TEXT,
    profile_picture_thumbnail_url TEXT,
    cover_image_url TEXT,
    video_intro_url TEXT, -- Professional introduction video
    
    -- Social Links
    website_url TEXT,
    linkedin_url TEXT,
    github_url TEXT,
    portfolio_url TEXT,
    twitter_url TEXT,
    behance_url TEXT,
    dribbble_url TEXT,
    
    -- Work Preferences (For Freelancers)
    hourly_rate DECIMAL(10, 2),
    hourly_rate_currency CHAR(3) DEFAULT 'USD',
    hourly_rate_visibility VARCHAR(20) DEFAULT 'PUBLIC' CHECK (
        hourly_rate_visibility IN ('PUBLIC', 'PRIVATE', 'CLIENTS_ONLY')
    ),
    
    -- Availability (References availability-be service)
    availability_status VARCHAR(20) DEFAULT 'AVAILABLE' CHECK (
        availability_status IN ('AVAILABLE', 'BUSY', 'NOT_AVAILABLE', 'VACATION')
    ),
    availability_hours_per_week INTEGER,
    available_from DATE,
    
    -- Search & Discovery
    search_visibility BOOLEAN DEFAULT TRUE,
    featured_profile BOOLEAN DEFAULT FALSE,
    profile_views_count INTEGER DEFAULT 0,
    
    -- Verification Status
    identity_verified BOOLEAN DEFAULT FALSE,
    identity_verified_at TIMESTAMPTZ,
    email_verified BOOLEAN DEFAULT FALSE,
    phone_verified BOOLEAN DEFAULT FALSE,
    payment_verified BOOLEAN DEFAULT FALSE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Foreign Keys
    CONSTRAINT fk_profiles_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Indexes for user_profiles
CREATE INDEX idx_profiles_user_id ON user_profiles (user_id);
CREATE INDEX idx_profiles_country ON user_profiles (country_code) WHERE search_visibility = TRUE;
CREATE INDEX idx_profiles_featured ON user_profiles (featured_profile) WHERE featured_profile = TRUE;
CREATE INDEX idx_profiles_hourly_rate ON user_profiles (hourly_rate) WHERE search_visibility = TRUE;
CREATE INDEX idx_profiles_search ON user_profiles USING gin (
    to_tsvector('english', COALESCE(professional_title, '') || ' ' || COALESCE(bio, '') || ' ' || COALESCE(tagline, ''))
);

-- 1.3 User Preferences (Application Settings)
CREATE TABLE user_preferences (
    preference_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL UNIQUE,
    
    -- Notification Preferences (Detailed)
    email_notifications_enabled BOOLEAN DEFAULT TRUE,
    push_notifications_enabled BOOLEAN DEFAULT TRUE,
    sms_notifications_enabled BOOLEAN DEFAULT FALSE,
    
    -- Email Categories
    email_job_alerts BOOLEAN DEFAULT TRUE,
    email_messages BOOLEAN DEFAULT TRUE,
    email_proposals BOOLEAN DEFAULT TRUE,
    email_contract_updates BOOLEAN DEFAULT TRUE,
    email_payments BOOLEAN DEFAULT TRUE,
    email_reviews BOOLEAN DEFAULT TRUE,
    email_marketing BOOLEAN DEFAULT FALSE,
    email_product_updates BOOLEAN DEFAULT TRUE,
    
    -- Notification Frequency
    notification_frequency VARCHAR(20) DEFAULT 'REAL_TIME' CHECK (
        notification_frequency IN ('REAL_TIME', 'HOURLY', 'DAILY', 'WEEKLY', 'NEVER')
    ),
    digest_enabled BOOLEAN DEFAULT FALSE,
    digest_schedule VARCHAR(20) DEFAULT 'DAILY' CHECK (
        digest_schedule IN ('DAILY', 'WEEKLY', 'BIWEEKLY', 'MONTHLY')
    ),
    digest_time TIME DEFAULT '09:00:00',
    
    -- Privacy Settings
    profile_visibility VARCHAR(20) DEFAULT 'PUBLIC' CHECK (
        profile_visibility IN ('PUBLIC', 'CONNECTIONS_ONLY', 'PRIVATE')
    ),
    show_email BOOLEAN DEFAULT FALSE,
    show_phone BOOLEAN DEFAULT FALSE,
    show_location BOOLEAN DEFAULT TRUE,
    show_earnings BOOLEAN DEFAULT FALSE,
    show_activity BOOLEAN DEFAULT TRUE,
    allow_search_engines BOOLEAN DEFAULT TRUE,
    
    -- Communication Preferences
    allow_direct_messages BOOLEAN DEFAULT TRUE,
    allow_job_invitations BOOLEAN DEFAULT TRUE,
    allow_connection_requests BOOLEAN DEFAULT TRUE,
    
    -- Platform Preferences
    theme VARCHAR(20) DEFAULT 'LIGHT' CHECK (theme IN ('LIGHT', 'DARK', 'AUTO')),
    language VARCHAR(10) DEFAULT 'en',
    date_format VARCHAR(20) DEFAULT 'MM/DD/YYYY',
    time_format VARCHAR(10) DEFAULT '12H' CHECK (time_format IN ('12H', '24H')),
    currency CHAR(3) DEFAULT 'USD',
    
    -- Accessibility
    accessibility_high_contrast BOOLEAN DEFAULT FALSE,
    accessibility_large_text BOOLEAN DEFAULT FALSE,
    accessibility_screen_reader_mode BOOLEAN DEFAULT FALSE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_preferences_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_preferences_user_id ON user_preferences (user_id);

-- =========================================
-- SECTION 2: CAPABILITIES DOMAIN
-- =========================================

-- 2.1 Skills Taxonomy (Master List - Single Source of Truth)
CREATE TABLE skills_taxonomy (
    skill_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Skill Identity
    skill_name VARCHAR(100) NOT NULL UNIQUE,
    skill_slug VARCHAR(100) NOT NULL UNIQUE,
    skill_description TEXT,
    
    -- Categorization (Hierarchical)
    parent_skill_id UUID, -- For sub-categories (e.g., React -> Frontend -> Web Dev)
    category VARCHAR(100), -- "Programming", "Design", "Marketing", "Writing"
    subcategory VARCHAR(100),
    level INTEGER DEFAULT 0, -- Hierarchy level (0 = root, 1 = category, 2 = subcategory)
    
    -- Metadata
    is_active BOOLEAN DEFAULT TRUE,
    is_featured BOOLEAN DEFAULT FALSE,
    popularity_score INTEGER DEFAULT 0, -- How many users have this skill
    demand_score INTEGER DEFAULT 0, -- How many jobs require this skill
    
    -- Search & Synonyms
    synonyms TEXT[], -- Alternative names ["ReactJS", "React.js"]
    related_skills UUID[], -- Related skill IDs
    
    -- External References
    linkedin_skill_id VARCHAR(100),
    indeed_skill_id VARCHAR(100),
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_skills_parent FOREIGN KEY (parent_skill_id) REFERENCES skills_taxonomy(skill_id)
);

CREATE INDEX idx_skills_name ON skills_taxonomy (skill_name);
CREATE INDEX idx_skills_slug ON skills_taxonomy (skill_slug);
CREATE INDEX idx_skills_category ON skills_taxonomy (category, subcategory) WHERE is_active = TRUE;
CREATE INDEX idx_skills_parent ON skills_taxonomy (parent_skill_id);
CREATE INDEX idx_skills_popularity ON skills_taxonomy (popularity_score DESC);
CREATE INDEX idx_skills_search ON skills_taxonomy USING gin (
    to_tsvector('english', skill_name || ' ' || COALESCE(skill_description, ''))
);

-- 2.2 User Skills (Freelancer Skills with Proficiency)
CREATE TABLE user_skills (
    user_skill_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    skill_id UUID NOT NULL,
    
    -- Proficiency
    proficiency_level VARCHAR(20) NOT NULL CHECK (
        proficiency_level IN ('BEGINNER', 'INTERMEDIATE', 'ADVANCED', 'EXPERT')
    ),
    years_of_experience INTEGER CHECK (years_of_experience >= 0),
    
    -- Verification
    is_verified BOOLEAN DEFAULT FALSE,
    verified_at TIMESTAMPTZ,
    verified_by VARCHAR(50), -- 'PLATFORM', 'TEST', 'ENDORSEMENT', 'CERTIFICATE'
    verification_score INTEGER CHECK (verification_score BETWEEN 0 AND 100),
    
    -- Usage Stats
    projects_count INTEGER DEFAULT 0, -- How many projects used this skill
    endorsements_count INTEGER DEFAULT 0,
    
    -- Display
    is_primary BOOLEAN DEFAULT FALSE, -- Featured/highlighted skill
    display_order INTEGER DEFAULT 0,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_user_skills_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT fk_user_skills_skill FOREIGN KEY (skill_id) REFERENCES skills_taxonomy(skill_id),
    CONSTRAINT uk_user_skills UNIQUE (user_id, skill_id)
);

CREATE INDEX idx_user_skills_user ON user_skills (user_id);
CREATE INDEX idx_user_skills_skill ON user_skills (skill_id);
CREATE INDEX idx_user_skills_proficiency ON user_skills (user_id, proficiency_level);
CREATE INDEX idx_user_skills_primary ON user_skills (user_id, is_primary) WHERE is_primary = TRUE;

-- 2.3 Specializations (Niche Expertise)
CREATE TABLE specializations (
    specialization_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Specialization Details
    title VARCHAR(200) NOT NULL, -- "React + TypeScript for FinTech Applications"
    description TEXT,
    
    -- Associated Skills
    primary_skills UUID[], -- Array of skill_ids
    secondary_skills UUID[],
    
    -- Industries/Domains
    industries TEXT[], -- ["FinTech", "Healthcare", "E-commerce"]
    target_clients TEXT[], -- ["Startups", "Enterprise", "Agencies"]
    
    -- Verification & Credibility
    is_verified BOOLEAN DEFAULT FALSE,
    verified_at TIMESTAMPTZ,
    evidence_urls TEXT[], -- Portfolio links, case studies
    
    -- Metrics
    projects_completed INTEGER DEFAULT 0,
    client_satisfaction DECIMAL(3, 2), -- Average rating for this specialization
    
    -- Display
    is_featured BOOLEAN DEFAULT FALSE,
    display_order INTEGER DEFAULT 0,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_specializations_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_specializations_user ON specializations (user_id);
CREATE INDEX idx_specializations_verified ON specializations (is_verified) WHERE is_verified = TRUE;
CREATE INDEX idx_specializations_featured ON specializations (user_id, is_featured) WHERE is_featured = TRUE;

-- 2.4 Service Catalog (Predefined Service Offerings)
CREATE TABLE service_catalog (
    service_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Service Details
    service_name VARCHAR(200) NOT NULL,
    service_slug VARCHAR(200) NOT NULL,
    short_description VARCHAR(500),
    full_description TEXT,
    
    -- Capabilities Reference (NO duplicate skill data)
    required_skills UUID[], -- skill_ids from skills_taxonomy
    specialization_id UUID, -- Optional reference to specialization
    
    -- Pricing
    base_price DECIMAL(10, 2) NOT NULL,
    currency CHAR(3) DEFAULT 'USD',
    pricing_model VARCHAR(20) CHECK (pricing_model IN ('FIXED', 'HOURLY', 'CUSTOM')),
    
    -- Service Scope
    estimated_delivery_days INTEGER,
    revisions_included INTEGER DEFAULT 0,
    
    -- Service Packages (if applicable)
    has_packages BOOLEAN DEFAULT FALSE,
    
    -- Media
    cover_image_url TEXT,
    gallery_urls TEXT[],
    
    -- Status & Visibility
    is_active BOOLEAN DEFAULT TRUE,
    is_featured BOOLEAN DEFAULT FALSE,
    
    -- Metrics
    orders_count INTEGER DEFAULT 0,
    average_rating DECIMAL(3, 2),
    views_count INTEGER DEFAULT 0,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_service_catalog_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT fk_service_catalog_specialization FOREIGN KEY (specialization_id) REFERENCES specializations(specialization_id)
);

CREATE INDEX idx_service_catalog_user ON service_catalog (user_id);
CREATE INDEX idx_service_catalog_active ON service_catalog (is_active) WHERE is_active = TRUE;
CREATE INDEX idx_service_catalog_featured ON service_catalog (is_featured) WHERE is_featured = TRUE;
CREATE INDEX idx_service_catalog_slug ON service_catalog (service_slug);

-- 2.5 Service Packages (Tiered Offerings)
CREATE TABLE service_packages (
    package_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    service_id UUID NOT NULL,
    
    -- Package Details
    package_tier VARCHAR(20) NOT NULL CHECK (package_tier IN ('BASIC', 'STANDARD', 'PREMIUM', 'CUSTOM')),
    package_name VARCHAR(100) NOT NULL,
    description TEXT,
    
    -- Pricing
    price DECIMAL(10, 2) NOT NULL,
    currency CHAR(3) DEFAULT 'USD',
    
    -- Deliverables
    delivery_days INTEGER NOT NULL,
    revisions_included INTEGER DEFAULT 0,
    features TEXT[], -- List of what's included
    
    -- Display
    is_popular BOOLEAN DEFAULT FALSE,
    display_order INTEGER DEFAULT 0,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_service_packages_service FOREIGN KEY (service_id) REFERENCES service_catalog(service_id) ON DELETE CASCADE,
    CONSTRAINT uk_service_packages UNIQUE (service_id, package_tier)
);

CREATE INDEX idx_service_packages_service ON service_packages (service_id);

-- =========================================
-- SECTION 3: EXPERIENCE & EDUCATION DOMAIN
-- =========================================

-- 3.1 Work Experience
CREATE TABLE work_experience (
    experience_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Company Information
    company_name VARCHAR(200) NOT NULL,
    company_logo_url TEXT,
    company_website TEXT,
    
    -- Position Details
    job_title VARCHAR(200) NOT NULL,
    employment_type VARCHAR(50) CHECK (
        employment_type IN ('FULL_TIME', 'PART_TIME', 'CONTRACT', 'FREELANCE', 'INTERNSHIP', 'VOLUNTEER')
    ),
    
    -- Duration
    start_date DATE NOT NULL,
    end_date DATE,
    is_current BOOLEAN DEFAULT FALSE,
    duration_months INTEGER, -- Calculated field
    
    -- Location
    location VARCHAR(200),
    is_remote BOOLEAN DEFAULT FALSE,
    
    -- Description
    description TEXT,
    responsibilities TEXT[],
    achievements TEXT[],
    
    -- Skills Used (References skills_taxonomy)
    skills_used UUID[], -- skill_ids
    
    -- Verification
    is_verified BOOLEAN DEFAULT FALSE,
    verified_by VARCHAR(100), -- LinkedIn, email verification, etc.
    
    -- Display
    display_order INTEGER DEFAULT 0,
    is_featured BOOLEAN DEFAULT FALSE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_experience_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT chk_experience_dates CHECK (end_date IS NULL OR end_date >= start_date)
);

CREATE INDEX idx_experience_user ON work_experience (user_id);
CREATE INDEX idx_experience_current ON work_experience (user_id, is_current) WHERE is_current = TRUE;
CREATE INDEX idx_experience_dates ON work_experience (start_date DESC, end_date DESC);

-- 3.2 Education
CREATE TABLE education (
    education_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Institution
    school_name VARCHAR(200) NOT NULL,
    school_logo_url TEXT,
    school_website TEXT,
    
    -- Degree Information
    degree_type VARCHAR(50), -- "Bachelor's", "Master's", "PhD", "Certificate", "Bootcamp"
    degree_name VARCHAR(200), -- "Computer Science", "Graphic Design"
    field_of_study VARCHAR(200),
    
    -- Duration
    start_year INTEGER,
    end_year INTEGER,
    graduation_date DATE,
    is_current BOOLEAN DEFAULT FALSE,
    
    -- Academic Details
    grade_gpa VARCHAR(20), -- "3.8/4.0", "First Class Honours"
    activities TEXT[], -- Clubs, organizations
    achievements TEXT[],
    
    -- Description
    description TEXT,
    
    -- Verification
    is_verified BOOLEAN DEFAULT FALSE,
    verification_document_url TEXT,
    
    -- Display
    display_order INTEGER DEFAULT 0,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_education_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT chk_education_years CHECK (end_year IS NULL OR end_year >= start_year)
);

CREATE INDEX idx_education_user ON education (user_id);
CREATE INDEX idx_education_current ON education (user_id, is_current) WHERE is_current = TRUE;

-- 3.3 Languages
CREATE TABLE user_languages (
    language_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Language Details
    language_code VARCHAR(10) NOT NULL, -- ISO 639-1
    language_name VARCHAR(100) NOT NULL, -- "English", "Spanish", "Mandarin"
    
    -- Proficiency
    proficiency_level VARCHAR(20) NOT NULL CHECK (
        proficiency_level IN ('BASIC', 'CONVERSATIONAL', 'FLUENT', 'NATIVE')
    ),
    
    -- Verification
    is_verified BOOLEAN DEFAULT FALSE,
    verification_method VARCHAR(50), -- "TEST", "CERTIFICATE", "ENDORSEMENT"
    
    -- Display
    is_primary BOOLEAN DEFAULT FALSE,
    display_order INTEGER DEFAULT 0,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_languages_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT uk_user_languages UNIQUE (user_id, language_code)
);

CREATE INDEX idx_languages_user ON user_languages (user_id);
CREATE INDEX idx_languages_primary ON user_languages (user_id, is_primary) WHERE is_primary = TRUE;

-- =========================================
-- SECTION 4: CREDENTIALS DOMAIN
-- =========================================

-- 4.1 External Certifications (AWS, Google, Microsoft, etc.)
CREATE TABLE external_certifications (
    certification_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Certification Details
    certification_name VARCHAR(200) NOT NULL,
    issuing_organization VARCHAR(200) NOT NULL,
    credential_id VARCHAR(200), -- Unique ID from issuer
    credential_url TEXT, -- Verification URL
    
    -- Dates
    issue_date DATE NOT NULL,
    expiry_date DATE,
    does_not_expire BOOLEAN DEFAULT FALSE,
    
    -- Associated Skills
    skills_demonstrated UUID[], -- skill_ids from skills_taxonomy
    
    -- Verification
    is_verified BOOLEAN DEFAULT FALSE,
    verified_at TIMESTAMPTZ,
    
    -- Media
    certificate_image_url TEXT,
    badge_image_url TEXT,
    
    -- Display
    display_order INTEGER DEFAULT 0,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_ext_cert_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_ext_cert_user ON external_certifications (user_id);
CREATE INDEX idx_ext_cert_org ON external_certifications (issuing_organization);
CREATE INDEX idx_ext_cert_expiry ON external_certifications (expiry_date) WHERE expiry_date IS NOT NULL;

-- 4.2 Platform Certifications (Skillsier-issued)
CREATE TABLE platform_certifications (
    platform_cert_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Certification Type
    certification_type VARCHAR(50) NOT NULL, -- "SKILL_TEST", "PROJECT_REVIEW", "INTERVIEW"
    skill_id UUID, -- Reference to skills_taxonomy
    
    -- Test/Assessment Details
    test_name VARCHAR(200),
    test_score INTEGER CHECK (test_score BETWEEN 0 AND 100),
    passing_score INTEGER,
    
    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'PASSED', 'FAILED', 'EXPIRED', 'REVOKED')
    ),
    
    -- Dates
    issued_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ,
    
    -- Verification
    certificate_number VARCHAR(100) UNIQUE,
    verification_url TEXT,
    
    -- Badge Display
    badge_image_url TEXT,
    badge_level VARCHAR(20) CHECK (badge_level IN ('BRONZE', 'SILVER', 'GOLD', 'PLATINUM')),
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_platform_cert_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT fk_platform_cert_skill FOREIGN KEY (skill_id) REFERENCES skills_taxonomy(skill_id)
);

CREATE INDEX idx_platform_cert_user ON platform_certifications (user_id);
CREATE INDEX idx_platform_cert_skill ON platform_certifications (skill_id);
CREATE INDEX idx_platform_cert_status ON platform_certifications (status, expires_at);

-- =========================================
-- SECTION 5: PORTFOLIOS & PROJECTS
-- =========================================

-- 5.1 Portfolio Items
CREATE TABLE portfolio_items (
    portfolio_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Project Details
    title VARCHAR(200) NOT NULL,
    description TEXT,
    project_url TEXT,
    
    -- Project Type
    project_type VARCHAR(50) CHECK (
        project_type IN ('PERSONAL', 'CLIENT_WORK', 'OPEN_SOURCE', 'ACADEMIC', 'COMPETITION')
    ),
    
    -- Media
    cover_image_url TEXT NOT NULL,
    images_urls TEXT[],
    video_url TEXT,
    
    -- Skills & Technologies
    skills_used UUID[], -- skill_ids
    technologies TEXT[], -- Freeform: "React 18", "AWS Lambda"
    
    -- Client Information (if applicable)
    client_name VARCHAR(200),
    project_duration VARCHAR(50), -- "2 months", "Ongoing"
    completion_date DATE,
    
    -- Metrics
    views_count INTEGER DEFAULT 0,
    likes_count INTEGER DEFAULT 0,
    
    -- Display & Visibility
    is_featured BOOLEAN DEFAULT FALSE,
    is_public BOOLEAN DEFAULT TRUE,
    display_order INTEGER DEFAULT 0,
    
    -- External Links
    github_url TEXT,
    live_demo_url TEXT,
    case_study_url TEXT,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_portfolio_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_portfolio_user ON portfolio_items (user_id);
CREATE INDEX idx_portfolio_featured ON portfolio_items (user_id, is_featured) WHERE is_featured = TRUE;
CREATE INDEX idx_portfolio_public ON portfolio_items (is_public) WHERE is_public = TRUE;

-- =========================================
-- SECTION 6: IDENTITY VERIFICATION DOMAIN
-- =========================================

-- 6.1 Identity Verification (KYC/KYB)
CREATE TABLE identity_verifications (
    verification_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Verification Type
    verification_type VARCHAR(20) NOT NULL CHECK (
        verification_type IN ('KYC_INDIVIDUAL', 'KYB_BUSINESS', 'ADDRESS', 'PAYMENT_METHOD')
    ),
    
    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'IN_REVIEW', 'APPROVED', 'REJECTED', 'EXPIRED', 'REQUIRES_RESUBMISSION')
    ),
    
    -- Document Information (Encrypted)
    document_type VARCHAR(50), -- "PASSPORT", "DRIVERS_LICENSE", "NATIONAL_ID", "BUSINESS_LICENSE"
    document_number_hash VARCHAR(255), -- Hashed for privacy
    document_country CHAR(2),
    document_expiry_date DATE,
    
    -- Uploaded Documents (References storage-be)
    document_front_storage_id UUID, -- File ID from storage-be
    document_back_storage_id UUID,
    selfie_storage_id UUID,
    address_proof_storage_id UUID,
    
    -- Review Information
    reviewed_by UUID, -- Admin user_id
    reviewed_at TIMESTAMPTZ,
    rejection_reason TEXT,
    requires_resubmission_reason TEXT,
    
    -- Approval Details
    approved_at TIMESTAMPTZ,
    verification_expires_at TIMESTAMPTZ,
    
    -- Third-Party Verification Service
    external_verification_provider VARCHAR(50), -- "ONFIDO", "JUMIO", "STRIPE_IDENTITY"
    external_verification_id VARCHAR(200),
    external_verification_status VARCHAR(50),
    
    -- Audit Trail
    submission_ip INET,
    submission_user_agent TEXT,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_identity_ver_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT fk_identity_ver_reviewer FOREIGN KEY (reviewed_by) REFERENCES users(user_id)
);

CREATE INDEX idx_identity_ver_user ON identity_verifications (user_id);
CREATE INDEX idx_identity_ver_status ON identity_verifications (status);
CREATE INDEX idx_identity_ver_type ON identity_verifications (verification_type, status);

-- =========================================
-- SECTION 7: TRUST & SAFETY DOMAIN
-- =========================================

-- 7.1 Trust Score
CREATE TABLE trust_scores (
    trust_score_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL UNIQUE,
    
    -- Overall Score (0-100)
    overall_score INTEGER DEFAULT 50 CHECK (overall_score BETWEEN 0 AND 100),
    
    -- Score Components
    identity_verification_score INTEGER DEFAULT 0 CHECK (identity_verification_score BETWEEN 0 AND 100),
    payment_verification_score INTEGER DEFAULT 0 CHECK (payment_verification_score BETWEEN 0 AND 100),
    profile_completion_score INTEGER DEFAULT 0 CHECK (profile_completion_score BETWEEN 0 AND 100),
    reputation_score INTEGER DEFAULT 0 CHECK (reputation_score BETWEEN 0 AND 100),
    activity_score INTEGER DEFAULT 0 CHECK (activity_score BETWEEN 0 AND 100),
    
    -- Trust Indicators
    email_verified BOOLEAN DEFAULT FALSE,
    phone_verified BOOLEAN DEFAULT FALSE,
    identity_verified BOOLEAN DEFAULT FALSE,
    payment_method_verified BOOLEAN DEFAULT FALSE,
    address_verified BOOLEAN DEFAULT FALSE,
    
    -- Platform Activity
    days_on_platform INTEGER DEFAULT 0,
    completed_projects INTEGER DEFAULT 0,
    cancellation_rate DECIMAL(5, 2) DEFAULT 0.00,
    response_rate DECIMAL(5, 2) DEFAULT 0.00,
    
    -- Risk Flags
    has_active_disputes BOOLEAN DEFAULT FALSE,
    has_payment_issues BOOLEAN DEFAULT FALSE,
    flagged_for_review BOOLEAN DEFAULT FALSE,
    
    -- Last Calculation
    last_calculated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_trust_score_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_trust_score_user ON trust_scores (user_id);
CREATE INDEX idx_trust_score_overall ON trust_scores (overall_score DESC);

-- 7.2 Background Checks
CREATE TABLE background_checks (
    background_check_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Check Type
    check_type VARCHAR(50) NOT NULL CHECK (
        check_type IN ('CRIMINAL', 'EMPLOYMENT', 'EDUCATION', 'CREDIT', 'REFERENCE', 'COMPREHENSIVE')
    ),
    
    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'IN_PROGRESS', 'COMPLETED', 'FAILED', 'EXPIRED')
    ),
    
    -- Provider Information
    provider_name VARCHAR(100), -- "CHECKR", "STERLING", "FIRST_ADVANTAGE"
    provider_report_id VARCHAR(200),
    provider_report_url TEXT,
    
    -- Results
    result VARCHAR(20) CHECK (result IN ('CLEAR', 'CONSIDER', 'SUSPENDED', 'FLAGGED')),
    findings TEXT[],
    report_summary TEXT,
    
    -- Dates
    initiated_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ,
    
    -- Consent
    user_consent_given BOOLEAN DEFAULT FALSE,
    consent_given_at TIMESTAMPTZ,
    consent_ip INET,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_background_check_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_background_check_user ON background_checks (user_id);
CREATE INDEX idx_background_check_status ON background_checks (status);
CREATE INDEX idx_background_check_expires ON background_checks (expires_at);

-- 7.3 Risk Assessment
CREATE TABLE risk_assessments (
    risk_assessment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Risk Level
    risk_level VARCHAR(20) DEFAULT 'MEDIUM' CHECK (
        risk_level IN ('LOW', 'MEDIUM', 'HIGH', 'CRITICAL')
    ),
    risk_score INTEGER CHECK (risk_score BETWEEN 0 AND 100),
    
    -- Risk Factors
    risk_factors JSONB, -- Detailed breakdown of risk factors
    
    -- Red Flags
    has_identity_concerns BOOLEAN DEFAULT FALSE,
    has_payment_fraud_indicators BOOLEAN DEFAULT FALSE,
    has_platform_abuse_history BOOLEAN DEFAULT FALSE,
    has_multiple_accounts BOOLEAN DEFAULT FALSE,
    suspicious_activity_detected BOOLEAN DEFAULT FALSE,
    
    -- Geographic Risk
    high_risk_country BOOLEAN DEFAULT FALSE,
    vpn_detected BOOLEAN DEFAULT FALSE,
    
    -- Behavioral Analysis
    login_pattern_anomaly BOOLEAN DEFAULT FALSE,
    transaction_pattern_anomaly BOOLEAN DEFAULT FALSE,
    
    -- Assessment Details
    assessed_by VARCHAR(50), -- "SYSTEM", "MANUAL", "ML_MODEL"
    assessment_reason TEXT,
    
    -- Actions Taken
    restrictions_applied TEXT[],
    monitoring_enabled BOOLEAN DEFAULT FALSE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at TIMESTAMPTZ,
    
    CONSTRAINT fk_risk_assessment_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_risk_assessment_user ON risk_assessments (user_id);
CREATE INDEX idx_risk_assessment_level ON risk_assessments (risk_level);
CREATE INDEX idx_risk_assessment_expires ON risk_assessments (expires_at);

-- =========================================
-- SECTION 8: COMPLIANCE & MODERATION DOMAIN
-- =========================================

-- 7.4 User Reports (Reports against users)
CREATE TABLE user_reports (
    report_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Reporter & Reported
    reported_user_id UUID NOT NULL,
    reporter_user_id UUID,
    reporter_type VARCHAR(20) CHECK (reporter_type IN ('USER', 'SYSTEM', 'ADMIN')),
    
    -- Report Details
    report_category VARCHAR(50) NOT NULL CHECK (
        report_category IN (
            'HARASSMENT', 'SPAM', 'FRAUD', 'FAKE_PROFILE', 'INAPPROPRIATE_CONTENT',
            'COPYRIGHT_VIOLATION', 'IMPERSONATION', 'SCAM', 'OTHER'
        )
    ),
    report_reason TEXT NOT NULL,
    evidence_urls TEXT[],
    
    -- Related Entities (References to other services)
    related_job_id UUID, -- From jobs-be
    related_proposal_id UUID, -- From proposals-be
    related_contract_id UUID, -- From contracts-be
    related_message_id UUID, -- From communications-be
    
    -- Status & Review
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'UNDER_REVIEW', 'RESOLVED', 'DISMISSED', 'ESCALATED')
    ),
    
    priority VARCHAR(20) DEFAULT 'MEDIUM' CHECK (priority IN ('LOW', 'MEDIUM', 'HIGH', 'URGENT')),
    
    -- Moderation
    assigned_to UUID, -- Moderator user_id
    reviewed_by UUID, -- Reviewer user_id
    reviewed_at TIMESTAMPTZ,
    resolution_notes TEXT,
    action_taken VARCHAR(100), -- "WARNING_ISSUED", "ACCOUNT_SUSPENDED", "CONTENT_REMOVED"
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    resolved_at TIMESTAMPTZ,
    
    CONSTRAINT fk_user_reports_reported FOREIGN KEY (reported_user_id) REFERENCES users(user_id),
    CONSTRAINT fk_user_reports_reporter FOREIGN KEY (reporter_user_id) REFERENCES users(user_id),
    CONSTRAINT fk_user_reports_assigned FOREIGN KEY (assigned_to) REFERENCES users(user_id),
    CONSTRAINT fk_user_reports_reviewed FOREIGN KEY (reviewed_by) REFERENCES users(user_id)
);

CREATE INDEX idx_user_reports_reported ON user_reports (reported_user_id);
CREATE INDEX idx_user_reports_reporter ON user_reports (reporter_user_id);
CREATE INDEX idx_user_reports_status ON user_reports (status);
CREATE INDEX idx_user_reports_priority ON user_reports (priority, status);
CREATE INDEX idx_user_reports_assigned ON user_reports (assigned_to) WHERE assigned_to IS NOT NULL;

-- 7.5 Moderation Actions (Admin actions on users)
CREATE TABLE moderation_actions (
    action_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Action Details
    action_type VARCHAR(50) NOT NULL CHECK (
        action_type IN (
            'WARNING', 'TEMPORARY_SUSPENSION', 'PERMANENT_BAN', 'CONTENT_REMOVAL',
            'FEATURE_RESTRICTION', 'ACCOUNT_REVIEW', 'DEMOTION', 'REINSTATEMENT'
        )
    ),
    
    -- Reason & Evidence
    reason TEXT NOT NULL,
    internal_notes TEXT, -- Not visible to user
    related_report_ids UUID[], -- References user_reports
    evidence_urls TEXT[],
    
    -- Action Parameters
    duration_days INTEGER, -- For temporary actions
    restrictions JSONB, -- {"messaging": false, "job_posting": false}
    expires_at TIMESTAMPTZ,
    
    -- Moderator Information
    actioned_by UUID NOT NULL,
    approved_by UUID, -- Senior moderator approval
    
    -- Status
    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (
        status IN ('ACTIVE', 'EXPIRED', 'REVOKED', 'APPEALED')
    ),
    
    -- User Notification
    user_notified BOOLEAN DEFAULT FALSE,
    notification_sent_at TIMESTAMPTZ,
    
    -- Appeal
    appeal_allowed BOOLEAN DEFAULT TRUE,
    appeal_deadline TIMESTAMPTZ,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_moderation_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT fk_moderation_actioned_by FOREIGN KEY (actioned_by) REFERENCES users(user_id),
    CONSTRAINT fk_moderation_approved_by FOREIGN KEY (approved_by) REFERENCES users(user_id)
);

CREATE INDEX idx_moderation_user ON moderation_actions (user_id);
CREATE INDEX idx_moderation_type ON moderation_actions (action_type);
CREATE INDEX idx_moderation_status ON moderation_actions (status, expires_at);
CREATE INDEX idx_moderation_actioned_by ON moderation_actions (actioned_by);

-- 7.6 Moderation Appeals
CREATE TABLE moderation_appeals (
    appeal_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    moderation_action_id UUID NOT NULL,
    user_id UUID NOT NULL,
    
    -- Appeal Details
    appeal_reason TEXT NOT NULL,
    supporting_evidence TEXT[],
    evidence_urls TEXT[],
    
    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'UNDER_REVIEW', 'APPROVED', 'DENIED', 'WITHDRAWN')
    ),
    
    -- Review
    reviewed_by UUID,
    reviewed_at TIMESTAMPTZ,
    review_decision TEXT,
    decision_reason TEXT,
    
    -- Outcome
    action_modified BOOLEAN DEFAULT FALSE,
    new_action_type VARCHAR(50),
    new_duration_days INTEGER,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    resolved_at TIMESTAMPTZ,
    
    CONSTRAINT fk_appeals_action FOREIGN KEY (moderation_action_id) REFERENCES moderation_actions(action_id),
    CONSTRAINT fk_appeals_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT fk_appeals_reviewed_by FOREIGN KEY (reviewed_by) REFERENCES users(user_id),
    CONSTRAINT uk_appeals_action UNIQUE (moderation_action_id)
);

CREATE INDEX idx_appeals_action ON moderation_appeals (moderation_action_id);
CREATE INDEX idx_appeals_user ON moderation_appeals (user_id);
CREATE INDEX idx_appeals_status ON moderation_appeals (status);

-- =========================================
-- SECTION 9: SCORING DOMAINS
-- =========================================

-- 9.1 User Metrics (Raw Data Collection)
CREATE TABLE user_metrics (
    metric_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Performance Metrics
    total_jobs_completed INTEGER DEFAULT 0,
    total_earnings DECIMAL(12, 2) DEFAULT 0,
    total_hours_worked DECIMAL(10, 2) DEFAULT 0,
    
    -- Quality Metrics
    average_rating DECIMAL(3, 2) DEFAULT 0.00,
    total_reviews INTEGER DEFAULT 0,
    positive_reviews INTEGER DEFAULT 0,
    negative_reviews INTEGER DEFAULT 0,
    
    -- Response & Delivery
    average_response_time_hours DECIMAL(8, 2), -- Hours to first response
    median_delivery_time_days DECIMAL(8, 2),
    on_time_delivery_rate DECIMAL(5, 2) DEFAULT 100.00,
    
    -- Reliability
    completion_rate DECIMAL(5, 2) DEFAULT 100.00, -- % of accepted projects completed
    cancellation_rate DECIMAL(5, 2) DEFAULT 0.00,
    dispute_rate DECIMAL(5, 2) DEFAULT 0.00,
    
    -- Client Satisfaction
    client_satisfaction_score DECIMAL(5, 2) DEFAULT 0.00, -- 0-100
    rehire_rate DECIMAL(5, 2) DEFAULT 0.00,
    
    -- Activity
    proposals_sent INTEGER DEFAULT 0,
    proposals_accepted INTEGER DEFAULT 0,
    proposal_acceptance_rate DECIMAL(5, 2) DEFAULT 0.00,
    
    -- Engagement
    profile_views INTEGER DEFAULT 0,
    job_invitations_received INTEGER DEFAULT 0,
    
    -- Time-based Stats
    days_since_last_project INTEGER DEFAULT 0,
    active_projects_count INTEGER DEFAULT 0,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    last_calculated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_user_metrics_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT uk_user_metrics_user UNIQUE (user_id)
);

CREATE INDEX idx_user_metrics_user ON user_metrics (user_id);
CREATE INDEX idx_user_metrics_rating ON user_metrics (average_rating DESC);
CREATE INDEX idx_user_metrics_completion ON user_metrics (completion_rate DESC);

-- 9.2 Metric History (Time-series data)
CREATE TABLE metric_history (
    history_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Metric Information
    metric_name VARCHAR(100) NOT NULL, -- "average_rating", "completion_rate"
    metric_value DECIMAL(12, 2) NOT NULL,
    metric_unit VARCHAR(20), -- "percentage", "decimal", "count", "hours"
    
    -- Context
    measurement_period VARCHAR(20), -- "DAILY", "WEEKLY", "MONTHLY", "QUARTERLY"
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    
    -- Timestamp
    recorded_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_metric_history_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_metric_history_user ON metric_history (user_id, metric_name, recorded_at DESC);
CREATE INDEX idx_metric_history_period ON metric_history (period_start, period_end);

-- 9.3 Reputation Score
CREATE TABLE reputation_scores (
    reputation_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL UNIQUE,
    
    -- Overall Reputation (0-100)
    overall_reputation INTEGER DEFAULT 50 CHECK (overall_reputation BETWEEN 0 AND 100),
    
    -- Component Scores (weighted)
    quality_score INTEGER DEFAULT 50 CHECK (quality_score BETWEEN 0 AND 100),
    reliability_score INTEGER DEFAULT 50 CHECK (reliability_score BETWEEN 0 AND 100),
    professionalism_score INTEGER DEFAULT 50 CHECK (professionalism_score BETWEEN 0 AND 100),
    communication_score INTEGER DEFAULT 50 CHECK (communication_score BETWEEN 0 AND 100),
    
    -- Score Breakdown (detailed)
    rating_component DECIMAL(5, 2) DEFAULT 0.00, -- Contribution from reviews
    completion_component DECIMAL(5, 2) DEFAULT 0.00, -- Contribution from completion rate
    response_component DECIMAL(5, 2) DEFAULT 0.00, -- Contribution from response time
    tenure_component DECIMAL(5, 2) DEFAULT 0.00, -- Contribution from platform tenure
    
    -- Reputation Tier
    reputation_tier VARCHAR(20) DEFAULT 'NEWCOMER' CHECK (
        reputation_tier IN ('NEWCOMER', 'RISING_TALENT', 'ESTABLISHED', 'TOP_RATED', 'ELITE')
    ),
    
    -- Badge Display
    display_badge BOOLEAN DEFAULT FALSE,
    badge_earned_at TIMESTAMPTZ,
    
    -- Calculation Metadata
    data_points_count INTEGER DEFAULT 0, -- Number of projects/reviews used
    confidence_level DECIMAL(5, 2) DEFAULT 0.00, -- Statistical confidence
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    last_calculated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_reputation_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_reputation_user ON reputation_scores (user_id);
CREATE INDEX idx_reputation_overall ON reputation_scores (overall_reputation DESC);
CREATE INDEX idx_reputation_tier ON reputation_scores (reputation_tier);

-- 9.4 Success Score (Prediction/ML-based)
CREATE TABLE success_scores (
    success_score_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL UNIQUE,
    
    -- Predictive Scores (0-100)
    overall_success_score INTEGER DEFAULT 50 CHECK (overall_success_score BETWEEN 0 AND 100),
    job_win_probability DECIMAL(5, 2) DEFAULT 0.00, -- Likelihood of winning proposals
    client_satisfaction_prediction DECIMAL(5, 2) DEFAULT 0.00,
    project_completion_prediction DECIMAL(5, 2) DEFAULT 0.00,
    
    -- Category-specific Success Scores
    success_by_category JSONB, -- {"web_dev": 85, "mobile_dev": 70}
    
    -- ML Model Information
    model_version VARCHAR(50),
    model_confidence DECIMAL(5, 2),
    
    -- Feature Importance
    top_success_factors TEXT[], -- Factors contributing to score
    improvement_areas TEXT[], -- Areas for improvement
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    last_calculated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_success_score_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_success_score_user ON success_scores (user_id);
CREATE INDEX idx_success_score_overall ON success_scores (overall_success_score DESC);

-- =========================================
-- SECTION 10: ACHIEVEMENTS & BADGES DOMAIN
-- =========================================

-- 10.1 Achievement Definitions (Master List)
CREATE TABLE achievement_definitions (
    achievement_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Achievement Identity
    achievement_code VARCHAR(100) NOT NULL UNIQUE, -- "FIRST_PROJECT", "TOP_RATED_2024"
    achievement_name VARCHAR(200) NOT NULL,
    achievement_description TEXT,
    
    -- Category
    category VARCHAR(50) NOT NULL CHECK (
        category IN ('MILESTONE', 'SKILL', 'QUALITY', 'COMMUNITY', 'SPECIAL', 'SEASONAL')
    ),
    
    -- Criteria
    criteria JSONB NOT NULL, -- {"projects_completed": 1} or complex conditions
    
    -- Rewards
    reputation_points INTEGER DEFAULT 0,
    platform_credits DECIMAL(10, 2) DEFAULT 0,
    
    -- Display
    badge_icon_url TEXT,
    badge_color VARCHAR(20),
    rarity VARCHAR(20) CHECK (rarity IN ('COMMON', 'UNCOMMON', 'RARE', 'EPIC', 'LEGENDARY')),
    
    -- Visibility
    is_active BOOLEAN DEFAULT TRUE,
    is_public BOOLEAN DEFAULT TRUE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_achievement_defs_code ON achievement_definitions (achievement_code);
CREATE INDEX idx_achievement_defs_category ON achievement_definitions (category) WHERE is_active = TRUE;

-- 10.2 User Achievements (Earned Badges)
CREATE TABLE user_achievements (
    user_achievement_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    achievement_id UUID NOT NULL,
    
    -- Progress
    progress_current INTEGER DEFAULT 0,
    progress_target INTEGER,
    progress_percentage DECIMAL(5, 2) DEFAULT 0.00,
    
    -- Status
    status VARCHAR(20) DEFAULT 'IN_PROGRESS' CHECK (
        status IN ('LOCKED', 'IN_PROGRESS', 'UNLOCKED', 'CLAIMED', 'EXPIRED')
    ),
    
    -- Unlock Details
    unlocked_at TIMESTAMPTZ,
    claimed_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ,
    
    -- Display
    is_featured BOOLEAN DEFAULT FALSE,
    display_on_profile BOOLEAN DEFAULT TRUE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_user_achievements_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT fk_user_achievements_achievement FOREIGN KEY (achievement_id) REFERENCES achievement_definitions(achievement_id),
    CONSTRAINT uk_user_achievements UNIQUE (user_id, achievement_id)
);

CREATE INDEX idx_user_achievements_user ON user_achievements (user_id);
CREATE INDEX idx_user_achievements_status ON user_achievements (user_id, status);
CREATE INDEX idx_user_achievements_featured ON user_achievements (user_id, is_featured) WHERE is_featured = TRUE;

-- =========================================
-- SECTION 11: CONNECTIONS & NETWORKING DOMAIN
-- =========================================

-- 11.1 User Connections (Professional Network)
CREATE TABLE user_connections (
    connection_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Connection Parties
    user_id UUID NOT NULL,
    connected_user_id UUID NOT NULL,
    
    -- Connection Type
    connection_type VARCHAR(20) DEFAULT 'CONNECTION' CHECK (
        connection_type IN ('CONNECTION', 'FOLLOW', 'BLOCK', 'FAVORITE')
    ),
    
    -- Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'ACCEPTED', 'DECLINED', 'BLOCKED')
    ),
    
    -- Request Details
    requested_by UUID NOT NULL,
    request_message TEXT,
    
    -- Response
    responded_at TIMESTAMPTZ,
    decline_reason VARCHAR(100),
    
    -- Interaction Stats
    messages_exchanged INTEGER DEFAULT 0,
    projects_together INTEGER DEFAULT 0,
    last_interaction_at TIMESTAMPTZ,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_connections_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT fk_connections_connected_user FOREIGN KEY (connected_user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT fk_connections_requested_by FOREIGN KEY (requested_by) REFERENCES users(user_id),
    CONSTRAINT uk_user_connections UNIQUE (user_id, connected_user_id),
    CONSTRAINT chk_connections_different_users CHECK (user_id != connected_user_id)
);

CREATE INDEX idx_connections_user ON user_connections (user_id, status);
CREATE INDEX idx_connections_connected_user ON user_connections (connected_user_id, status);
CREATE INDEX idx_connections_type ON user_connections (connection_type);

-- 11.2 Endorsements (Skill Endorsements)
CREATE TABLE skill_endorsements (
    endorsement_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Endorsement Parties
    endorsed_user_id UUID NOT NULL,
    endorser_user_id UUID NOT NULL,
    
    -- Skill Reference
    user_skill_id UUID NOT NULL, -- References user_skills table
    skill_id UUID NOT NULL, -- References skills_taxonomy
    
    -- Endorsement Details
    endorsement_text TEXT,
    relationship VARCHAR(50) CHECK (
        relationship IN ('COLLEAGUE', 'CLIENT', 'MANAGER', 'TEAM_MEMBER', 'PARTNER', 'OTHER')
    ),
    
    -- Project Context (Optional)
    project_reference UUID, -- From contracts-be or portfolio
    worked_together BOOLEAN DEFAULT FALSE,
    
    -- Visibility
    is_public BOOLEAN DEFAULT TRUE,
    is_featured BOOLEAN DEFAULT FALSE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_endorsements_endorsed FOREIGN KEY (endorsed_user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT fk_endorsements_endorser FOREIGN KEY (endorser_user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT fk_endorsements_user_skill FOREIGN KEY (user_skill_id) REFERENCES user_skills(user_skill_id) ON DELETE CASCADE,
    CONSTRAINT fk_endorsements_skill FOREIGN KEY (skill_id) REFERENCES skills_taxonomy(skill_id),
    CONSTRAINT uk_endorsements UNIQUE (endorsed_user_id, endorser_user_id, skill_id),
    CONSTRAINT chk_endorsements_different_users CHECK (endorsed_user_id != endorser_user_id)
);

CREATE INDEX idx_endorsements_endorsed ON skill_endorsements (endorsed_user_id);
CREATE INDEX idx_endorsements_endorser ON skill_endorsements (endorser_user_id);
CREATE INDEX idx_endorsements_skill ON skill_endorsements (skill_id);

-- =========================================
-- SECTION 12: AVAILABILITY DOMAIN (LOCAL CACHE)
-- =========================================

-- 12.1 Availability Status (Read Model - Synced from availability-be)
CREATE TABLE availability_status (
    availability_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL UNIQUE,
    
    -- Current Status
    status VARCHAR(20) DEFAULT 'AVAILABLE' CHECK (
        status IN ('AVAILABLE', 'BUSY', 'PARTIALLY_AVAILABLE', 'NOT_AVAILABLE', 'VACATION')
    ),
    
    -- Capacity
    hours_per_week INTEGER,
    max_concurrent_projects INTEGER,
    current_active_projects INTEGER DEFAULT 0,
    
    -- Availability Window
    available_from DATE,
    available_until DATE,
    next_available_date DATE,
    
    -- Preferences
    preferred_project_duration VARCHAR(50), -- "SHORT_TERM", "LONG_TERM", "BOTH"
    minimum_project_budget DECIMAL(10, 2),
    preferred_work_times JSONB, -- {"timezone": "EST", "hours": "9-17"}
    
    -- Last Sync (from availability-be)
    last_synced_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    sync_source VARCHAR(50) DEFAULT 'availability-be',
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_availability_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_availability_user ON availability_status (user_id);
CREATE INDEX idx_availability_status ON availability_status (status) WHERE status IN ('AVAILABLE', 'PARTIALLY_AVAILABLE');

-- =========================================
-- SECTION 13: NOTIFICATIONS & PREFERENCES DOMAIN
-- =========================================

-- 13.1 Notification Settings (Granular Control)
CREATE TABLE notification_settings (
    setting_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL UNIQUE,
    
    -- Channel Enablement
    email_enabled BOOLEAN DEFAULT TRUE,
    push_enabled BOOLEAN DEFAULT TRUE,
    sms_enabled BOOLEAN DEFAULT FALSE,
    in_app_enabled BOOLEAN DEFAULT TRUE,
    
    -- Category Settings (JSON for flexibility)
    job_notifications JSONB DEFAULT '{"email": true, "push": true, "sms": false}'::jsonb,
    proposal_notifications JSONB DEFAULT '{"email": true, "push": true, "sms": false}'::jsonb,
    message_notifications JSONB DEFAULT '{"email": true, "push": true, "sms": false}'::jsonb,
    contract_notifications JSONB DEFAULT '{"email": true, "push": true, "sms": false}'::jsonb,
    payment_notifications JSONB DEFAULT '{"email": true, "push": true, "sms": true}'::jsonb,
    review_notifications JSONB DEFAULT '{"email": true, "push": true, "sms": false}'::jsonb,
    security_notifications JSONB DEFAULT '{"email": true, "push": true, "sms": true}'::jsonb,
    marketing_notifications JSONB DEFAULT '{"email": false, "push": false, "sms": false}'::jsonb,
    
    -- Digest Settings
    daily_digest_enabled BOOLEAN DEFAULT FALSE,
    daily_digest_time TIME DEFAULT '09:00:00',
    weekly_digest_enabled BOOLEAN DEFAULT FALSE,
    weekly_digest_day INTEGER DEFAULT 1 CHECK (weekly_digest_day BETWEEN 0 AND 6), -- 0=Sunday
    
    -- Quiet Hours
    quiet_hours_enabled BOOLEAN DEFAULT FALSE,
    quiet_hours_start TIME DEFAULT '22:00:00',
    quiet_hours_end TIME DEFAULT '08:00:00',
    quiet_hours_timezone VARCHAR(50),
    
    -- Smart Notifications
    intelligent_grouping BOOLEAN DEFAULT TRUE,
    priority_notifications_only BOOLEAN DEFAULT FALSE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_notification_settings_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_notification_settings_user ON notification_settings (user_id);

-- =========================================
-- SECTION 14: EVENT SOURCING & OUTBOX PATTERN
-- =========================================

-- 14.1 Outbox Table (Transactional Outbox Pattern)
CREATE TABLE outbox_events (
    event_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Event Identity
    aggregate_id UUID NOT NULL, -- user_id, profile_id, etc.
    aggregate_type VARCHAR(100) NOT NULL, -- "user", "profile", "skill"
    event_type VARCHAR(200) NOT NULL, -- "user.created.v1", "profile.updated.v1"
    event_version VARCHAR(20) DEFAULT 'v1',
    
    -- Event Payload (Non-PII only - IDs and codes)
    payload JSONB NOT NULL,
    
    -- Event Metadata
    correlation_id UUID, -- For tracing across services
    causation_id UUID, -- Parent event that caused this
    
    -- Actor Information
    actor_user_id UUID,
    actor_type VARCHAR(20) CHECK (actor_type IN ('USER', 'SYSTEM', 'ADMIN', 'SERVICE')),
    
    -- Idempotency
    idempotency_key VARCHAR(255) UNIQUE,
    
    -- Processing Status
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (
        status IN ('PENDING', 'PROCESSING', 'PUBLISHED', 'FAILED', 'DEAD_LETTER')
    ),
    
    -- Publishing Details
    published_at TIMESTAMPTZ,
    publish_attempts INTEGER DEFAULT 0,
    last_attempt_at TIMESTAMPTZ,
    error_message TEXT,
    
    -- Kafka Topic
    topic_name VARCHAR(200) NOT NULL, -- "users.events", "profiles.events"
    partition_key VARCHAR(255), -- For Kafka partitioning
    
    -- Metadata
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    -- Indexes for efficient polling
    CONSTRAINT fk_outbox_actor FOREIGN KEY (actor_user_id) REFERENCES users(user_id)
);

CREATE INDEX idx_outbox_status ON outbox_events (status, created_at) WHERE status IN ('PENDING', 'FAILED');
CREATE INDEX idx_outbox_aggregate ON outbox_events (aggregate_id, aggregate_type, created_at DESC);
CREATE INDEX idx_outbox_event_type ON outbox_events (event_type, created_at DESC);
CREATE INDEX idx_outbox_correlation ON outbox_events (correlation_id);
CREATE INDEX idx_outbox_idempotency ON outbox_events (idempotency_key) WHERE idempotency_key IS NOT NULL;

-- 14.2 Event Store (Full Event History - Optional)
CREATE TABLE event_store (
    event_store_id BIGSERIAL PRIMARY KEY,
    event_id UUID NOT NULL UNIQUE,
    
    -- Event Identity
    aggregate_id UUID NOT NULL,
    aggregate_type VARCHAR(100) NOT NULL,
    event_type VARCHAR(200) NOT NULL,
    event_version VARCHAR(20) DEFAULT 'v1',
    sequence_number BIGINT NOT NULL, -- Per-aggregate sequence
    
    -- Event Data
    event_data JSONB NOT NULL,
    metadata JSONB,
    
    -- Causation
    correlation_id UUID,
    causation_id UUID,
    
    -- Actor
    actor_user_id UUID,
    actor_type VARCHAR(20),
    
    -- Timestamp
    occurred_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT uk_event_store_aggregate_seq UNIQUE (aggregate_id, sequence_number)
);

CREATE INDEX idx_event_store_aggregate ON event_store (aggregate_id, sequence_number);
CREATE INDEX idx_event_store_type ON event_store (event_type, occurred_at DESC);
CREATE INDEX idx_event_store_occurred ON event_store (occurred_at DESC);

-- 14.3 Consumed Events Tracking (Idempotency for Event Consumers)
CREATE TABLE consumed_events (
    consumed_event_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Event Identity
    event_id UUID NOT NULL UNIQUE, -- From source service
    event_type VARCHAR(200) NOT NULL,
    source_service VARCHAR(100) NOT NULL, -- "jobs-be", "contracts-be"
    
    -- Consumer Information
    consumer_name VARCHAR(100) NOT NULL, -- "JobEventConsumer", "ContractEventConsumer"
    
    -- Processing
    processed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    processing_duration_ms INTEGER,
    
    -- Idempotency
    idempotency_key VARCHAR(255),
    
    -- Event Data (for debugging)
    event_payload JSONB,
    
    CONSTRAINT uk_consumed_events_event UNIQUE (event_id, consumer_name)
);

CREATE INDEX idx_consumed_events_event_id ON consumed_events (event_id);
CREATE INDEX idx_consumed_events_type ON consumed_events (event_type, source_service);
CREATE INDEX idx_consumed_events_processed ON consumed_events (processed_at DESC);

-- =========================================
-- SECTION 15: READ MODELS / PROJECTIONS
-- =========================================

-- 15.1 User Read Model (Optimized for Queries)
CREATE TABLE user_read_model (
    user_id UUID PRIMARY KEY,
    
    -- Basic Info
    full_name VARCHAR(300),
    display_name VARCHAR(200),
    email VARCHAR(255),
    user_type VARCHAR(20),
    account_status VARCHAR(20),
    
    -- Profile
    professional_title VARCHAR(200),
    tagline VARCHAR(300),
    profile_picture_url TEXT,
    country_code CHAR(2),
    city VARCHAR(100),
    
    -- Stats (Denormalized)
    profile_completion_score INTEGER,
    reputation_score INTEGER,
    trust_score INTEGER,
    total_projects INTEGER,
    average_rating DECIMAL(3, 2),
    
    -- Availability
    availability_status VARCHAR(20),
    hourly_rate DECIMAL(10, 2),
    
    -- Skills (Top 5 for quick display)
    top_skills JSONB, -- [{"skill_name": "React", "proficiency": "EXPERT"}]
    
    -- Verification
    email_verified BOOLEAN,
    identity_verified BOOLEAN,
    
    -- Search Optimization
    search_vector tsvector,
    
    -- Timestamps
    last_active_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_user_read_model FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_user_read_search ON user_read_model USING gin(search_vector);
CREATE INDEX idx_user_read_reputation ON user_read_model (reputation_score DESC) WHERE account_status = 'ACTIVE';
CREATE INDEX idx_user_read_rating ON user_read_model (average_rating DESC) WHERE account_status = 'ACTIVE';
CREATE INDEX idx_user_read_type_status ON user_read_model (user_type, account_status);

-- 15.2 User Stats Read Model (Aggregated Statistics)
CREATE TABLE user_stats_read_model (
    user_id UUID PRIMARY KEY,
    
    -- Project Stats
    total_projects INTEGER DEFAULT 0,
    active_projects INTEGER DEFAULT 0,
    completed_projects INTEGER DEFAULT 0,
    cancelled_projects INTEGER DEFAULT 0,
    
    -- Financial Stats
    total_earnings DECIMAL(12, 2) DEFAULT 0,
    pending_earnings DECIMAL(12, 2) DEFAULT 0,
    total_hours_worked DECIMAL(10, 2) DEFAULT 0,
    
    -- Performance Stats
    completion_rate DECIMAL(5, 2) DEFAULT 100.00,
    on_time_delivery_rate DECIMAL(5, 2) DEFAULT 100.00,
    response_rate DECIMAL(5, 2) DEFAULT 0.00,
    
    -- Review Stats
    total_reviews INTEGER DEFAULT 0,
    average_rating DECIMAL(3, 2) DEFAULT 0.00,
    five_star_reviews INTEGER DEFAULT 0,
    four_star_reviews INTEGER DEFAULT 0,
    three_star_reviews INTEGER DEFAULT 0,
    two_star_reviews INTEGER DEFAULT 0,
    one_star_reviews INTEGER DEFAULT 0,
    
    -- Engagement Stats
    profile_views INTEGER DEFAULT 0,
    proposal_sent INTEGER DEFAULT 0,
    proposals_accepted INTEGER DEFAULT 0,
    invitations_received INTEGER DEFAULT 0,
    
    -- Timeline Stats
    days_on_platform INTEGER DEFAULT 0,
    days_since_last_project INTEGER,
    last_project_date DATE,
    
    -- Timestamps
    last_calculated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_user_stats_read_model FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_user_stats_earnings ON user_stats_read_model (total_earnings DESC);
CREATE INDEX idx_user_stats_rating ON user_stats_read_model (average_rating DESC);

-- 15.3 User Search Index (Elasticsearch/PostgreSQL Hybrid)
CREATE TABLE user_search_index (
    user_id UUID PRIMARY KEY,
    
    -- Searchable Fields
    full_name_normalized VARCHAR(300),
    professional_title_normalized VARCHAR(200),
    bio_text TEXT,
    skills_text TEXT, -- Concatenated skill names
    
    -- Filters
    user_type VARCHAR(20),
    account_status VARCHAR(20),
    country_code CHAR(2),
    hourly_rate DECIMAL(10, 2),
    availability_status VARCHAR(20),
    
    -- Scores for Ranking
    reputation_score INTEGER,
    profile_completion_score INTEGER,
    success_score INTEGER,
    
    -- Boolean Filters
    is_verified BOOLEAN,
    is_top_rated BOOLEAN,
    is_available BOOLEAN,
    
    -- Full-Text Search Vector
    search_vector tsvector,
    
    -- Timestamps
    indexed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_user_search_index FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_search_index_vector ON user_search_index USING gin(search_vector);
CREATE INDEX idx_search_index_filters ON user_search_index (user_type, account_status, is_available);
CREATE INDEX idx_search_index_location ON user_search_index (country_code) WHERE is_available = TRUE;
CREATE INDEX idx_search_index_rate ON user_search_index (hourly_rate) WHERE is_available = TRUE;

-- =========================================
-- SECTION 16: AUDIT & COMPLIANCE
-- =========================================

-- 16.1 Audit Log (Comprehensive Activity Tracking)
CREATE TABLE audit_logs (
    audit_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Entity Information
    entity_type VARCHAR(100) NOT NULL, -- "user", "profile", "skill"
    entity_id UUID NOT NULL,
    
    -- Action Details
    action VARCHAR(100) NOT NULL, -- "CREATE", "UPDATE", "DELETE", "LOGIN", "EXPORT"
    action_category VARCHAR(50), -- "AUTHENTICATION", "PROFILE_CHANGE", "SECURITY"
    
    -- Actor
    actor_user_id UUID,
    actor_type VARCHAR(20) CHECK (actor_type IN ('USER', 'SYSTEM', 'ADMIN', 'API')),
    actor_ip INET,
    actor_user_agent TEXT,
    
    -- Changes (for UPDATE actions)
    old_values JSONB,
    new_values JSONB,
    changed_fields TEXT[],
    
    -- Context
    request_id UUID,
    session_id UUID,
    api_endpoint VARCHAR(255),
    http_method VARCHAR(10),
    
    -- Compliance
    gdpr_relevant BOOLEAN DEFAULT FALSE,
    pii_accessed BOOLEAN DEFAULT FALSE,
    
    -- Timestamp
    occurred_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_audit_actor FOREIGN KEY (actor_user_id) REFERENCES users(user_id)
);

CREATE INDEX idx_audit_entity ON audit_logs (entity_type, entity_id, occurred_at DESC);
CREATE INDEX idx_audit_actor ON audit_logs (actor_user_id, occurred_at DESC);
CREATE INDEX idx_audit_action ON audit_logs (action, occurred_at DESC);
CREATE INDEX idx_audit_occurred ON audit_logs (occurred_at DESC);
CREATE INDEX idx_audit_gdpr ON audit_logs (gdpr_relevant) WHERE gdpr_relevant = TRUE;

-- 16.2 Data Access Log (PII Access Tracking)
CREATE TABLE data_access_logs (
    access_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Accessed User Data
    user_id UUID NOT NULL,
    data_type VARCHAR(100) NOT NULL, -- "EMAIL", "PHONE", "ADDRESS", "SSN", "FULL_PROFILE"
    data_fields TEXT[], -- Specific fields accessed
    
    -- Accessor
    accessor_user_id UUID,
    accessor_role VARCHAR(50),
    access_reason TEXT,
    
    -- Context
    access_method VARCHAR(50), -- "API", "ADMIN_PANEL", "EXPORT", "REPORT"
    ip_address INET,
    user_agent TEXT,
    
    -- Compliance
    legal_basis VARCHAR(100), -- "CONSENT", "CONTRACT", "LEGITIMATE_INTEREST", "LEGAL_OBLIGATION"
    consent_id UUID, -- Reference to user consent record
    
    -- Timestamp
    accessed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_data_access_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT fk_data_access_accessor FOREIGN KEY (accessor_user_id) REFERENCES users(user_id)
);

CREATE INDEX idx_data_access_user ON data_access_logs (user_id, accessed_at DESC);
CREATE INDEX idx_data_access_accessor ON data_access_logs (accessor_user_id, accessed_at DESC);
CREATE INDEX idx_data_access_type ON data_access_logs (data_type, accessed_at DESC);

-- 16.3 Data Retention Policy
CREATE TABLE data_retention_policies (
    policy_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- Policy Details
    data_type VARCHAR(100) NOT NULL UNIQUE, -- "USER_DATA", "AUDIT_LOGS", "EVENTS"
    retention_period_days INTEGER NOT NULL,
    
    -- Deletion Rules
    deletion_method VARCHAR(50) CHECK (deletion_method IN ('SOFT_DELETE', 'HARD_DELETE', 'ANONYMIZE')),
    anonymization_rules JSONB,
    
    -- Compliance
    legal_basis TEXT,
    applicable_regulations TEXT[], -- ["GDPR", "CCPA", "PIPEDA"]
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- 16.4 User Consent Records
CREATE TABLE user_consents (
    consent_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Consent Type
    consent_type VARCHAR(100) NOT NULL, -- "TERMS_OF_SERVICE", "PRIVACY_POLICY", "MARKETING", "DATA_PROCESSING"
    
    -- Consent Details
    consent_version VARCHAR(50) NOT NULL, -- Version of terms/policy
    consent_text TEXT,
    consent_given BOOLEAN NOT NULL,
    
    -- Context
    consent_method VARCHAR(50), -- "SIGNUP", "UPDATE", "EXPLICIT_ACTION"
    ip_address INET,
    user_agent TEXT,
    
    -- Withdrawal
    withdrawn_at TIMESTAMPTZ,
    withdrawal_reason TEXT,
    
    -- Timestamps
    given_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at TIMESTAMPTZ,
    
    CONSTRAINT fk_user_consents_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_user_consents_user ON user_consents (user_id);
CREATE INDEX idx_user_consents_type ON user_consents (consent_type, consent_given);
CREATE INDEX idx_user_consents_expires ON user_consents (expires_at) WHERE expires_at IS NOT NULL;

-- =========================================
-- SECTION 17: MICROSERVICES RELATIONS
-- =========================================

-- 17.1 Cross-Service References (External Entity IDs)
CREATE TABLE external_references (
    reference_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- External Service
    service_name VARCHAR(100) NOT NULL, -- "jobs-be", "contracts-be", "financial-be"
    entity_type VARCHAR(100) NOT NULL, -- "job", "contract", "wallet"
    entity_id UUID NOT NULL, -- ID from external service
    
    -- Reference Metadata
    reference_context JSONB, -- Additional context about the relation
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_external_refs_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT uk_external_refs UNIQUE (service_name, entity_type, entity_id)
);

CREATE INDEX idx_external_refs_user ON external_references (user_id, service_name);
CREATE INDEX idx_external_refs_entity ON external_references (service_name, entity_type, entity_id);

-- =========================================
-- SECTION 18: SESSIONS & SECURITY
-- =========================================

-- 18.1 User Sessions (Active Sessions Tracking)
CREATE TABLE user_sessions (
    session_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Session Details
    device_type VARCHAR(50), -- "WEB", "MOBILE_IOS", "MOBILE_ANDROID", "DESKTOP"
    device_name VARCHAR(200),
    device_id VARCHAR(255),
    
    -- Network
    ip_address INET NOT NULL,
    user_agent TEXT,
    location_country CHAR(2),
    location_city VARCHAR(100),
    
    -- Session Tokens
    refresh_token_hash VARCHAR(255),
    access_token_jti VARCHAR(255), -- JWT ID
    
    -- Session Status
    is_active BOOLEAN DEFAULT TRUE,
    last_activity_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    
    -- Security
    suspicious_activity BOOLEAN DEFAULT FALSE,
    forced_logout BOOLEAN DEFAULT FALSE,
    logout_reason TEXT,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    terminated_at TIMESTAMPTZ,
    
    CONSTRAINT fk_sessions_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_sessions_user ON user_sessions (user_id, is_active);
CREATE INDEX idx_sessions_token ON user_sessions (refresh_token_hash) WHERE is_active = TRUE;
CREATE INDEX idx_sessions_device ON user_sessions (device_id) WHERE device_id IS NOT NULL;
CREATE INDEX idx_sessions_expires ON user_sessions (expires_at) WHERE is_active = TRUE;

-- 18.2 Security Events
CREATE TABLE security_events (
    event_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Event Type
    event_type VARCHAR(100) NOT NULL CHECK (
        event_type IN (
            'LOGIN_SUCCESS', 'LOGIN_FAILED', 'PASSWORD_CHANGED', 'PASSWORD_RESET_REQUESTED',
            'EMAIL_CHANGED', 'PHONE_CHANGED', '2FA_ENABLED', '2FA_DISABLED',
            'SUSPICIOUS_ACTIVITY', 'ACCOUNT_LOCKED', 'SESSION_HIJACK_DETECTED'
        )
    ),
    
    -- Event Details
    event_description TEXT,
    severity VARCHAR(20) CHECK (severity IN ('LOW', 'MEDIUM', 'HIGH', 'CRITICAL')),
    
    -- Context
    ip_address INET,
    user_agent TEXT,
    device_id VARCHAR(255),
    location_country CHAR(2),
    
    -- Response
    action_taken VARCHAR(100), -- "NONE", "EMAIL_SENT", "ACCOUNT_LOCKED", "2FA_REQUIRED"
    
    -- Timestamp
    occurred_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_security_events_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_security_events_user ON security_events (user_id, occurred_at DESC);
CREATE INDEX idx_security_events_type ON security_events (event_type, occurred_at DESC);
CREATE INDEX idx_security_events_severity ON security_events (severity, occurred_at DESC);

-- =========================================
-- SECTION 19: FUTURE-PROOFING & EXTENSIBILITY
-- =========================================

-- 19.1 Custom User Fields (Flexible Schema)
CREATE TABLE custom_user_fields (
    field_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Field Definition
    field_key VARCHAR(100) NOT NULL,
    field_type VARCHAR(50) CHECK (field_type IN ('STRING', 'NUMBER', 'BOOLEAN', 'DATE', 'JSON')),
    
    -- Field Value
    string_value TEXT,
    number_value DECIMAL(20, 6),
    boolean_value BOOLEAN,
    date_value DATE,
    json_value JSONB,
    
    -- Metadata
    field_category VARCHAR(100), -- "CUSTOM_PROFILE", "PREFERENCES", "METADATA"
    is_searchable BOOLEAN DEFAULT FALSE,
    is_public BOOLEAN DEFAULT FALSE,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_custom_fields_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT uk_custom_fields UNIQUE (user_id, field_key)
);

CREATE INDEX idx_custom_fields_user ON custom_user_fields (user_id);
CREATE INDEX idx_custom_fields_key ON custom_user_fields (field_key);
CREATE INDEX idx_custom_fields_searchable ON custom_user_fields (field_key, is_searchable) WHERE is_searchable = TRUE;

-- 19.2 Feature Flags (Per-User Feature Control)
CREATE TABLE user_feature_flags (
    flag_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    
    -- Feature Definition
    feature_key VARCHAR(100) NOT NULL,
    is_enabled BOOLEAN DEFAULT FALSE,
    
    -- Rollout Control
    rollout_percentage INTEGER CHECK (rollout_percentage BETWEEN 0 AND 100),
    
    -- Context
    enabled_by VARCHAR(50), -- "ADMIN", "SYSTEM", "AB_TEST", "SUBSCRIPTION"
    enable_reason TEXT,
    
    -- Expiration
    expires_at TIMESTAMPTZ,
    
    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_feature_flags_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT uk_feature_flags UNIQUE (user_id, feature_key)
);

CREATE INDEX idx_feature_flags_user ON user_feature_flags (user_id);
CREATE INDEX idx_feature_flags_key ON user_feature_flags (feature_key, is_enabled);

-- =========================================
-- TRIGGERS & FUNCTIONS
-- =========================================

-- Trigger function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$ LANGUAGE plpgsql;

-- Apply updated_at trigger to all relevant tables
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_user_profiles_updated_at BEFORE UPDATE ON user_profiles
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_user_preferences_updated_at BEFORE UPDATE ON user_preferences
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Function to update search vector
CREATE OR REPLACE FUNCTION update_user_search_vector()
RETURNS TRIGGER AS $
BEGIN
    NEW.search_vector := 
        setweight(to_tsvector('english', COALESCE(NEW.full_name_normalized, '')), 'A') ||
        setweight(to_tsvector('english', COALESCE(NEW.professional_title_normalized, '')), 'B') ||
        setweight(to_tsvector('english', COALESCE(NEW.bio_text, '')), 'C') ||
        setweight(to_tsvector('english', COALESCE(NEW.skills_text, '')), 'B');
    RETURN NEW;
END;
$ LANGUAGE plpgsql;

CREATE TRIGGER update_search_vector_trigger
    BEFORE INSERT OR UPDATE ON user_search_index
    FOR EACH ROW EXECUTE FUNCTION update_user_search_vector();

-- =========================================
-- VIEWS FOR COMMON QUERIES
-- =========================================

-- View: Complete User Profile
CREATE OR REPLACE VIEW v_user_complete_profile AS
SELECT 
    u.user_id,
    u.email,
    u.first_name,
    u.last_name,
    u.user_type,
    u.account_status,
    u.created_at,
    u.last_active_at,
    p.professional_title,
    p.tagline,
    p.bio,
    p.profile_picture_url,
    p.country_code,
    p.city,
    p.hourly_rate,
    p.availability_status,
    ts.overall_score as trust_score,
    rs.overall_reputation as reputation_score,
    um.average_rating,
    um.total_jobs_completed,
    um.completion_rate
FROM users u
LEFT JOIN user_profiles p ON u.user_id = p.user_id
LEFT JOIN trust_scores ts ON u.user_id = ts.user_id
LEFT JOIN reputation_scores rs ON u.user_id = rs.user_id
LEFT JOIN user_metrics um ON u.user_id = um.user_id
WHERE u.is_deleted = FALSE;

-- View: Active Freelancers
CREATE OR REPLACE VIEW v_active_freelancers AS
SELECT 
    u.user_id,
    u.first_name || ' ' || u.last_name as full_name,
    p.professional_title,
    p.hourly_rate,
    p.country_code,
    p.availability_status,
    rs.overall_reputation,
    um.average_rating,
    um.total_jobs_completed,
    ARRAY_AGG(DISTINCT st.skill_name) FILTER (WHERE st.skill_id IS NOT NULL) as skills
FROM users u
INNER JOIN user_profiles p ON u.user_id = p.user_id
LEFT JOIN reputation_scores rs ON u.user_id = rs.user_id
LEFT JOIN user_metrics um ON u.user_id = um.user_id
LEFT JOIN user_skills us ON u.user_id = us.user_id
LEFT JOIN skills_taxonomy st ON us.skill_id = st.skill_id
WHERE u.user_type IN ('FREELANCER', 'HYBRID')
    AND u.account_status = 'ACTIVE'
    AND u.is_deleted = FALSE
    AND p.search_visibility = TRUE
GROUP BY u.user_id, p.professional_title, p.hourly_rate, p.country_code, 
         p.availability_status, rs.overall_reputation, um.average_rating, um.total_jobs_completed;

-- =========================================
-- COMMENTS FOR DOCUMENTATION
-- =========================================

COMMENT ON TABLE users IS 'Core user identity and authentication data';
COMMENT ON TABLE user_profiles IS 'Extended user profile information for professional presentation';
COMMENT ON TABLE skills_taxonomy IS 'Master list of all skills in the platform - single source of truth';
COMMENT ON TABLE user_skills IS 'Skills associated with users including proficiency levels';
COMMENT ON TABLE outbox_events IS 'Transactional outbox for reliable event publishing to Kafka';
COMMENT ON TABLE user_read_model IS 'Optimized read model for user queries and search';
COMMENT ON TABLE audit_logs IS 'Comprehensive audit trail for compliance and debugging';

-- =========================================
-- END OF USERS-BE DATABASE DESIGN
-- =========================================

-- Database Statistics Query
-- Run this to get overview of database structure:
/*
SELECT 
    schemaname,
    tablename,
    pg_size_pretty(pg_total_relation_size(schemaname||'.'||tablename)) AS size
FROM pg_tables
WHERE schemaname = 'public'
ORDER BY pg    -- Provider Information
