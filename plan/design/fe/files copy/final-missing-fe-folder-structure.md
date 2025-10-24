# Final Missing Frontend Folder Structure for Skillsier Application
## Completing All Requirements from fe-folder-structure-prompt.md

> **Note**: This document contains ONLY the folder structure elements that are:
> 1. Required by `fe-folder-structure-prompt.md`
> 2. NOT present in `combined-folder-structure.md`
> 3. NOT present in `missing-fe-folder-structure.md`
> 4. NOT present in `additional-missing-fe-folder-structure.md`
> 5. NOT present in `remaining-missing-fe-folder-structure.md`

---

## I. Critical Missing Admin Routes

### 1. Admin Incidents & Status Management

```
apps/web/src/app/[locale]/(admin)/
│
├── incidents/
│   ├── [incidentId]/
│   │   ├── edit/
│   │   │   └── page.tsx  # Edit incident details
│   │   │       # - Update status
│   │   │       # - Add postmortem
│   │   │       # - Affected services
│   │   │       # BE: utility/status
│   │   │       # PUT /v1/status/incidents/{incident_id}
│   │   ├── timeline/
│   │   │   └── page.tsx  # Incident timeline
│   │   │       # - Event log
│   │   │       # - Update history
│   │   │       # BE: utility/status
│   │   │       # GET /v1/status/incidents/{incident_id}/timeline
│   │   └── page.tsx  # Incident detail
│   │       # - Current status
│   │       # - Impact assessment
│   │       # - Resolution steps
│   │       # BE: utility/status
│   │       # GET /v1/status/incidents/{incident_id}
│   ├── create/
│   │   └── page.tsx  # Create new incident
│   │       # - Incident type selection
│   │       # - Severity level
│   │       # - Affected services
│   │       # BE: utility/status
│   │       # POST /v1/status/incidents
│   ├── history/
│   │   └── page.tsx  # Historical incidents
│   │       # - Past incidents archive
│   │       # - Postmortems
│   │       # - Lessons learned
│   │       # BE: utility/status
│   │       # GET /v1/status/incidents/history
│   └── page.tsx  # Active incidents dashboard
│       # - Current incidents
│       # - Quick actions
│       # - Status board
│       # BE: utility/status
│       # GET /v1/status/incidents?status=active
│
├── maintenance/
│   ├── [maintenanceId]/
│   │   ├── edit/
│   │   │   └── page.tsx  # Edit maintenance window
│   │   │       # BE: utility/status
│   │   │       # PUT /v1/status/maintenance/{maintenance_id}
│   │   └── page.tsx  # Maintenance detail
│   │       # BE: utility/status
│   │       # GET /v1/status/maintenance/{maintenance_id}
│   ├── schedule/
│   │   └── page.tsx  # Schedule maintenance
│   │       # - Date/time selection
│   │       # - Affected services
│   │       # - Notification plan
│   │       # BE: utility/status
│   │       # POST /v1/status/maintenance
│   └── page.tsx  # Maintenance calendar
│       # - Upcoming maintenance
│       # - Impact windows
│       # BE: utility/status
│       # GET /v1/status/maintenance
│
└── system-health/
    ├── services/
    │   └── page.tsx  # Service health overview
    │       # - All microservices status
    │       # - Uptime metrics
    │       # - Response times
    │       # BE: utility/status
    │       # GET /v1/status/services
    ├── metrics/
    │   └── page.tsx  # System metrics dashboard
    │       # - CPU/Memory usage
    │       # - Database performance
    │       # - Queue depths
    │       # BE: utility/metrics (or monitoring service)
    │       # GET /v1/metrics/system
    └── page.tsx  # Health dashboard
        # - Overall system health
        # - Critical alerts
        # - Performance trends
        # BE: utility/status
        # GET /v1/status/health
```

### 2. Admin Compliance & Privacy Tools

```
apps/web/src/app/[locale]/(admin)/
│
├── compliance/
│   ├── gdpr/
│   │   ├── export-requests/
│   │   │   ├── [requestId]/
│   │   │   │   └── page.tsx  # Export request detail
│   │   │   │       # - Request review
│   │   │   │       # - Generate export
│   │   │   │       # - Approve/deny
│   │   │   │       # BE: admin-be/privacy, users-be/user
│   │   │   │       # GET /v1/privacy/export-requests/{request_id}
│   │   │   │       # POST /v1/privacy/export-requests/{request_id}/process
│   │   │   └── page.tsx  # Export requests queue
│   │   │       # BE: admin-be/privacy
│   │   │       # GET /v1/privacy/export-requests
│   │   ├── deletion-requests/
│   │   │   ├── [requestId]/
│   │   │   │   └── page.tsx  # Deletion request detail
│   │   │   │       # - Data preview
│   │   │   │       # - Retention check
│   │   │   │       # - Process deletion
│   │   │   │       # BE: admin-be/privacy, users-be/user
│   │   │   │       # GET /v1/privacy/deletion-requests/{request_id}
│   │   │   │       # POST /v1/privacy/deletion-requests/{request_id}/process
│   │   │   └── page.tsx  # Deletion requests queue
│   │   │       # BE: admin-be/privacy
│   │   │       # GET /v1/privacy/deletion-requests
│   │   ├── consent-management/
│   │   │   └── page.tsx  # Consent logs & management
│   │   │       # - User consent history
│   │   │       # - Consent versions
│   │   │       # - Audit trail
│   │   │       # BE: users-be/consent, utility/audit
│   │   │       # GET /v1/users/consent-logs
│   │   └── reports/
│   │       └── page.tsx  # GDPR compliance reports
│   │           # - Processing activities
│   │           # - Data inventory
│   │           # - Breach reports
│   │           # BE: admin-be/privacy, utility/audit
│   │           # GET /v1/privacy/reports
│   │
│   ├── aml-kyc/
│   │   ├── monitoring/
│   │   │   └── page.tsx  # AML monitoring dashboard
│   │   │       # - Suspicious activity
│   │   │       # - Transaction patterns
│   │   │       # - Risk scores
│   │   │       # BE: admin-be/kyc_case, financial-be/transaction
│   │   │       # GET /v1/kyc/monitoring/suspicious-activity
│   │   ├── reports/
│   │   │   ├── [reportId]/
│   │   │   │   └── page.tsx  # SAR (Suspicious Activity Report) detail
│   │   │   │       # BE: admin-be/kyc_case
│   │   │   │       # GET /v1/kyc/reports/{report_id}
│   │   │   └── page.tsx  # AML reports list
│   │   │       # - Filed reports
│   │   │       # - Pending reports
│   │   │       # BE: admin-be/kyc_case
│   │   │       # GET /v1/kyc/reports
│   │   └── risk-assessment/
│   │       └── page.tsx  # Risk assessment tools
│   │           # - User risk profiles
│   │           # - Country risk matrix
│   │           # - Enhanced due diligence
│   │           # BE: admin-be/kyc_case
│   │           # GET /v1/kyc/risk-assessment
│   │
│   ├── data-retention/
│   │   ├── policies/
│   │   │   └── page.tsx  # Retention policies
│   │   │       # - Policy definitions
│   │   │       # - Data categories
│   │   │       # - Retention periods
│   │   │       # BE: admin-be/data_retention (if exists) or utility/config
│   │   │       # GET /v1/retention/policies
│   │   ├── schedule/
│   │   │   └── page.tsx  # Deletion schedule
│   │   │       # - Upcoming deletions
│   │   │       # - Retention expirations
│   │   │       # BE: admin-be/data_retention
│   │   │       # GET /v1/retention/schedule
│   │   └── audit/
│   │       └── page.tsx  # Retention audit log
│   │           # - Deletion history
│   │           # - Policy compliance
│   │           # BE: utility/audit
│   │           # GET /v1/audit/retention
│   │
│   └── document-verification/
│       ├── queue/
│       │   └── page.tsx  # Document verification queue
│       │       # - Pending documents
│       │       # - Priority sorting
│       │       # - Auto-verification status
│       │       # BE: admin-be/business_verification, storage-be/asset
│       │       # GET /v1/verification/documents/queue
│       ├── [documentId]/
│       │   └── page.tsx  # Document review interface
│       │       # - Document viewer
│       │       # - Verification checks
│       │       # - Approve/reject/request-more
│       │       # BE: admin-be/business_verification, storage-be/asset
│       │       # GET /v1/verification/documents/{document_id}
│       │       # PUT /v1/verification/documents/{document_id}/review
│       └── automated-checks/
│           └── page.tsx  # Automated verification rules
│               # - OCR settings
│               # - Validation rules
│               # - ML model performance
│               # BE: admin-be/business_verification
│               # GET /v1/verification/automation-rules
```

### 3. Admin Platform Configuration

```
apps/web/src/app/[locale]/(admin)/
│
├── platform-config/
│   ├── limits/
│   │   └── page.tsx  # Platform limits configuration
│   │       # - Rate limits
│   │       # - Upload limits
│   │       # - API quotas
│   │       # - Subscription limits
│   │       # BE: utility/config, admin-be/platform_config (if exists)
│   │       # GET /v1/config/limits
│   │       # PUT /v1/config/limits
│   │
│   ├── pricing/
│   │   └── page.tsx  # Pricing configuration
│   │       # - Commission rates
│   │       # - Subscription pricing
│   │       # - Regional pricing
│   │       # BE: financial-be/pricing_config (if exists)
│   │       # GET /v1/config/pricing
│   │       # PUT /v1/config/pricing
│   │       # Note: Requires change_approval
│   │
│   ├── notifications/
│   │   ├── templates/
│   │   │   ├── [templateId]/
│   │   │   │   ├── edit/
│   │   │   │   │   └── page.tsx  # Edit notification template
│   │   │   │   │       # BE: communications-be/template
│   │   │   │   │       # PUT /v1/notifications/templates/{template_id}
│   │   │   │   ├── preview/
│   │   │   │   │   └── page.tsx  # Preview template
│   │   │   │   │       # BE: communications-be/template
│   │   │   │   │       # POST /v1/notifications/templates/{template_id}/preview
│   │   │   │   └── page.tsx  # Template detail
│   │   │   │       # BE: communications-be/template
│   │   │   │       # GET /v1/notifications/templates/{template_id}
│   │   │   └── page.tsx  # Template library
│   │   │       # BE: communications-be/template
│   │   │       # GET /v1/notifications/templates
│   │   └── settings/
│   │       └── page.tsx  # Notification settings
│   │           # - Default preferences
│   │           # - Delivery channels
│   │           # - Retry policies
│   │           # BE: communications-be/config
│   │           # GET /v1/notifications/config
│   │           # PUT /v1/notifications/config
│   │
│   ├── integrations/
│   │   ├── [integrationId]/
│   │   │   ├── configure/
│   │   │   │   └── page.tsx  # Configure integration
│   │   │   │       # BE: admin-be/integrations (if exists)
│   │   │   │       # PUT /v1/integrations/{integration_id}/config
│   │   │   ├── logs/
│   │   │   │   └── page.tsx  # Integration logs
│   │   │   │       # BE: utility/audit
│   │   │   │       # GET /v1/integrations/{integration_id}/logs
│   │   │   └── page.tsx  # Integration detail
│   │   │       # BE: admin-be/integrations
│   │   │       # GET /v1/integrations/{integration_id}
│   │   └── page.tsx  # Integrations list
│   │       # - Payment providers
│   │       # - Email services
│   │       # - Storage providers
│   │       # - Auth providers
│   │       # BE: admin-be/integrations
│   │       # GET /v1/integrations
│   │
│   └── localization/
│       ├── languages/
│       │   └── page.tsx  # Language management
│       │       # - Enabled languages
│       │       # - Default language
│       │       # - RTL settings
│       │       # BE: utility/i18n
│       │       # GET /v1/config/languages
│       │       # PUT /v1/config/languages
│       └── regions/
│           └── page.tsx  # Regional settings
│               # - Timezone defaults
│               # - Currency settings
│               # - Date/time formats
│               # BE: utility/config
│               # GET /v1/config/regions
│               # PUT /v1/config/regions
```

---

## II. Missing User Journey Routes

### 1. Client-Specific Routes Not Covered

```
apps/web/src/app/[locale]/(dashboard)/
│
├── vendor-management/
│   ├── preferred/
│   │   ├── [vendorId]/
│   │   │   ├── performance/
│   │   │   │   └── page.tsx  # Vendor performance metrics
│   │   │   │       # - Success rate
│   │   │   │       # - Average delivery time
│   │   │   │       # - Quality scores
│   │   │   │       # BE: users-be/org (vendor subdomain), contracts-be/contract
│   │   │   │       # GET /v1/vendors/{vendor_id}/performance
│   │   │   ├── history/
│   │   │   │   └── page.tsx  # Work history with vendor
│   │   │   │       # - Past contracts
│   │   │   │       # - Total spend
│   │   │   │       # - Reviews given
│   │   │   │       # BE: contracts-be/contract, financial-be/invoice
│   │   │   │       # GET /v1/vendors/{vendor_id}/history
│   │   │   └── page.tsx  # Vendor detail
│   │   │       # BE: users-be/org (vendor subdomain)
│   │   │       # GET /v1/vendors/{vendor_id}
│   │   └── page.tsx  # Preferred vendors list
│   │       # - Star ratings
│   │       # - Quick invite
│   │       # BE: users-be/org
│   │       # GET /v1/vendors/preferred
│   │
│   ├── blacklist/
│   │   ├── [userId]/
│   │   │   └── page.tsx  # Blacklist entry detail
│   │   │       # - Reason for blacklist
│   │   │       # - Remove option
│   │   │       # BE: users-be/org (blacklist subdomain)
│   │   │       # GET /v1/vendors/blacklist/{user_id}
│   │   │       # DELETE /v1/vendors/blacklist/{user_id}
│   │   └── page.tsx  # Blacklisted vendors
│   │       # BE: users-be/org
│   │       # GET /v1/vendors/blacklist
│   │
│   └── compliance-docs/
│       ├── [vendorId]/
│       │   └── page.tsx  # Vendor compliance documents
│       │       # - W-9/W-8BEN forms
│       │       # - Insurance certificates
│       │       # - Background checks
│       │       # BE: admin-be/business_verification, storage-be/asset
│       │       # GET /v1/vendors/{vendor_id}/compliance-docs
│       └── page.tsx  # Compliance tracking
│           # - Expiring documents
│           # - Missing documents
│           # BE: admin-be/business_verification
│           # GET /v1/vendors/compliance-status
│
├── spend-analytics/
│   ├── by-category/
│   │   └── page.tsx  # Spending by category
│   │       # - Job category breakdown
│   │       # - Trend analysis
│   │       # BE: financial-be/analytics (if exists) or financial-be/invoice
│   │       # GET /v1/analytics/spend/by-category
│   ├── by-department/
│   │   └── page.tsx  # Spending by department
│   │       # - Cost center breakdown
│   │       # - Budget vs actual
│   │       # BE: financial-be/budget, financial-be/invoice
│   │       # GET /v1/analytics/spend/by-department
│   ├── by-vendor/
│   │   └── page.tsx  # Spending by vendor
│   │       # - Top vendors
│   │       # - Vendor concentration risk
│   │       # BE: financial-be/invoice, users-be/profile
│   │       # GET /v1/analytics/spend/by-vendor
│   ├── forecasting/
│   │   └── page.tsx  # Spend forecasting
│   │       # - Projected spend
│   │       # - Budget burn rate
│   │       # - Alerts for overages
│   │       # BE: financial-be/forecast (if exists) or financial-be/analytics
│   │       # GET /v1/analytics/spend/forecast
│   └── page.tsx  # Spend analytics dashboard
│       # - Overview metrics
│       # - Charts & visualizations
│       # BE: financial-be/analytics
│       # GET /v1/analytics/spend
│
└── sourcing/
    ├── campaigns/
    │   ├── [campaignId]/
    │   │   ├── edit/
    │   │   │   └── page.tsx  # Edit sourcing campaign
    │   │   │       # BE: jobs-be/campaign (if exists) or jobs-be/job
    │   │   │       # PUT /v1/sourcing/campaigns/{campaign_id}
    │   │   ├── analytics/
    │   │   │   └── page.tsx  # Campaign analytics
    │   │   │       # - Reach
    │   │   │       # - Engagement
    │   │   │       # - Conversions
    │   │   │       # BE: jobs-be/analytics
    │   │   │       # GET /v1/sourcing/campaigns/{campaign_id}/analytics
    │   │   └── page.tsx  # Campaign detail
    │   │       # BE: jobs-be/campaign
    │   │       # GET /v1/sourcing/campaigns/{campaign_id}
    │   ├── create/
    │   │   └── page.tsx  # Create sourcing campaign
    │   │       # - Target criteria
    │   │       # - Budget allocation
    │   │       # - Messaging
    │   │       # BE: jobs-be/campaign
    │   │       # POST /v1/sourcing/campaigns
    │   └── page.tsx  # Campaigns list
    │       # BE: jobs-be/campaign
    │       # GET /v1/sourcing/campaigns
    │
    ├── talent-pools/
    │   ├── [poolId]/
    │   │   ├── edit/
    │   │   │   └── page.tsx  # Edit talent pool
    │   │   │       # BE: users-be/talent_pool (if exists) or search-be/saved-search
    │   │   │       # PUT /v1/sourcing/talent-pools/{pool_id}
    │   │   ├── members/
    │   │   │   └── page.tsx  # Pool members
    │   │   │       # - Add/remove members
    │   │   │       # - Bulk actions
    │   │   │       # BE: users-be/talent_pool
    │   │   │       # GET /v1/sourcing/talent-pools/{pool_id}/members
    │   │   └── page.tsx  # Talent pool detail
    │   │       # BE: users-be/talent_pool
    │   │       # GET /v1/sourcing/talent-pools/{pool_id}
    │   ├── create/
    │   │   └── page.tsx  # Create talent pool
    │   │       # BE: users-be/talent_pool
    │   │       # POST /v1/sourcing/talent-pools
    │   └── page.tsx  # Talent pools list
    │       # BE: users-be/talent_pool
    │       # GET /v1/sourcing/talent-pools
    │
    └── invitations/
        ├── [invitationId]/
        │   └── page.tsx  # Invitation detail
        │       # - View status
        │       # - Resend invitation
        │       # BE: communications-be/invitation (if exists) or jobs-be/job
        │       # GET /v1/sourcing/invitations/{invitation_id}
        └── page.tsx  # Sent invitations
            # - Invitation history
            # - Response rates
            # BE: communications-be/invitation
            # GET /v1/sourcing/invitations
```

### 2. Freelancer-Specific Routes Not Covered

```
apps/web/src/app/[locale]/(dashboard)/
│
├── availability/
│   ├── calendar/
│   │   └── page.tsx  # Availability calendar
│   │       # - Mark available/busy
│   │       # - Recurring patterns
│   │       # - Sync with external calendar
│   │       # BE: users-be/availability (if exists) or users-be/profile
│   │       # GET /v1/users/me/availability
│   │       # PUT /v1/users/me/availability
│   ├── settings/
│   │   └── page.tsx  # Availability settings
│   │       # - Working hours
│   │       # - Timezone preferences
│   │       # - Auto-reply settings
│   │       # BE: users-be/settings
│   │       # PUT /v1/users/me/availability-settings
│   └── page.tsx  # Availability dashboard
│       # - Current status
│       # - Upcoming commitments
│       # BE: users-be/availability
│       # GET /v1/users/me/availability-overview
│
├── job-alerts/
│   ├── [alertId]/
│   │   ├── edit/
│   │   │   └── page.tsx  # Edit job alert
│   │   │       # BE: search-be/alert (if exists) or search-be/saved-search
│   │   │       # PUT /v1/job-alerts/{alert_id}
│   │   ├── history/
│   │   │   └── page.tsx  # Alert history
│   │   │       # - Jobs matched
│   │   │       # - Notifications sent
│   │   │       # BE: search-be/alert
│   │   │       # GET /v1/job-alerts/{alert_id}/history
│   │   └── page.tsx  # Alert detail
│   │       # BE: search-be/alert
│   │       # GET /v1/job-alerts/{alert_id}
│   ├── create/
│   │   └── page.tsx  # Create job alert
│   │       # - Search criteria
│   │       # - Notification frequency
│   │       # - Delivery channel
│   │       # BE: search-be/alert
│   │       # POST /v1/job-alerts
│   └── page.tsx  # Job alerts list
│       # - Active alerts
│       # - Pause/resume
│       # BE: search-be/alert
│       # GET /v1/job-alerts
│
├── learning/
│   ├── courses/
│   │   ├── [courseId]/
│   │   │   ├── lessons/
│   │   │   │   └── [lessonId]/
│   │   │   │       └── page.tsx  # Lesson content
│   │   │   │           # BE: learning-be/lesson (if exists) or external LMS
│   │   │   │           # GET /v1/learning/courses/{course_id}/lessons/{lesson_id}
│   │   │   ├── progress/
│   │   │   │   └── page.tsx  # Course progress
│   │   │   │       # BE: learning-be/progress
│   │   │   │       # GET /v1/learning/courses/{course_id}/progress
│   │   │   └── page.tsx  # Course detail
│   │   │       # BE: learning-be/course
│   │   │       # GET /v1/learning/courses/{course_id}
│   │   └── page.tsx  # Courses catalog
│   │       # BE: learning-be/course
│   │       # GET /v1/learning/courses
│   ├── assessments/
│   │   ├── [assessmentId]/
│   │   │   ├── take/
│   │   │   │   └── page.tsx  # Take assessment
│   │   │   │       # BE: learning-be/assessment
│   │   │   │       # POST /v1/learning/assessments/{assessment_id}/submit
│   │   │   ├── results/
│   │   │   │   └── page.tsx  # Assessment results
│   │   │   │       # BE: learning-be/assessment
│   │   │   │       # GET /v1/learning/assessments/{assessment_id}/results
│   │   │   └── page.tsx  # Assessment detail
│   │   │       # BE: learning-be/assessment
│   │   │       # GET /v1/learning/assessments/{assessment_id}
│   │   └── page.tsx  # Assessments list
│   │       # BE: learning-be/assessment
│   │       # GET /v1/learning/assessments
│   └── achievements/
│       └── page.tsx  # Learning achievements
│           # - Certificates earned
│           # - Badges
│           # - Skill verifications
│           # BE: learning-be/achievement
│           # GET /v1/learning/achievements
│
└── reputation/
    ├── overview/
    │   └── page.tsx  # Reputation overview
    │       # - Overall score
    │       # - Score breakdown
    │       # - Recent changes
    │       # BE: reviews-be/reputation (if exists) or reviews-be/review
    │       # GET /v1/reputation/overview
    ├── reviews-received/
    │   └── page.tsx  # Reviews received
    │       # - Client reviews
    │       # - Response option
    │       # BE: reviews-be/review
    │       # GET /v1/reviews/received
    ├── reviews-given/
    │   └── page.tsx  # Reviews given
    │       # - Reviews I left
    │       # - Edit option (if within timeframe)
    │       # BE: reviews-be/review
    │       # GET /v1/reviews/given
    ├── disputes/
    │   └── page.tsx  # Review disputes
    │       # - Disputed reviews
    │       # - Appeal status
    │       # BE: reviews-be/dispute (if exists) or admin-be/case_mgmt
    │       # GET /v1/reviews/disputes
    └── badges/
        └── page.tsx  # Achievement badges
            # - Top rated badge
            # - Rising talent
            # - Expert verified
            # BE: reviews-be/badge (if exists) or users-be/profile
            # GET /v1/reputation/badges
```

---

## III. Missing Mobile App Routes

### 1. Mobile-Specific Dashboard Routes

```
apps/mobile/app/
│
├── (dashboard)/
│   ├── quick-actions/
│   │   └── index.tsx  # Quick actions screen
│   │       # - Quick message
│   │       # - Quick proposal
│   │       # - Quick time entry
│   │       # - Quick invoice
│   │       # BE: Multiple services
│   │
│   ├── activity-feed/
│   │   └── index.tsx  # Activity feed
│   │       # - Recent activities
│   │       # - Notifications inline
│   │       # - Quick actions from feed
│   │       # BE: communications-be/notification, utility/activity
│   │       # GET /v1/activity/feed
│   │
│   ├── today/
│   │   └── index.tsx  # Today view
│   │       # - Today's schedule
│   │       # - Pending tasks
│   │       # - Quick metrics
│   │       # BE: contracts-be/work_diary, communications-be/notification
│   │       # GET /v1/today/overview
│   │
│   └── widgets/
│       ├── earnings/
│       │   └── index.tsx  # Earnings widget
│       │       # BE: financial-be/wallet
│       ├── time-tracker/
│       │   └── index.tsx  # Time tracker widget
│       │       # BE: contracts-be/work_diary
│       └── notifications/
│           └── index.tsx  # Notifications widget
│               # BE: communications-be/notification
│
├── camera/
│   ├── document-scan/
│   │   └── index.tsx  # Document scanner
│   │       # - ID verification
│   │       # - Contract documents
│   │       # - Expense receipts
│   │       # BE: storage-be/asset, admin-be/business_verification
│   │       # POST /v1/storage/uploads (scan)
│   ├── qr-scan/
│   │   └── index.tsx  # QR code scanner
│   │       # - Profile sharing
│   │       # - Event check-in
│   │       # BE: users-be/profile (QR data)
│   └── photo-upload/
│       └── index.tsx  # Photo upload
│           # - Portfolio photos
│           # - Work progress photos
│           # BE: storage-be/asset
│           # POST /v1/storage/uploads
│
├── offline/
│   ├── queue/
│   │   └── index.tsx  # Offline queue
│   │       # - Pending uploads
│   │       # - Queued messages
│   │       # - Draft proposals
│   │       # BE: None (local storage)
│   ├── sync/
│   │   └── index.tsx  # Sync status
│   │       # - Sync progress
│   │       # - Conflict resolution
│   │       # BE: Multiple services (sync endpoints)
│   └── settings/
│       └── index.tsx  # Offline settings
│           # - Auto-sync preferences
│           # - Download for offline
│           # BE: None (local storage)
│
└── accessibility/
    ├── voice-commands/
    │   └── index.tsx  # Voice commands
    │       # - Voice-to-text
    │       # - Command shortcuts
    └── screen-reader/
        └── index.tsx  # Screen reader optimized view
```

### 2. Mobile Settings & Preferences

```
apps/mobile/app/
│
├── settings/
│   ├── app-preferences/
│   │   ├── theme/
│   │   │   └── index.tsx  # Theme settings
│   │   │       # - Light/dark/auto
│   │   │       # - Accent color
│   │   ├── language/
│   │   │   └── index.tsx  # Language preferences
│   │   │       # - App language
│   │   │       # - Content language
│   │   ├── notifications/
│   │   │   ├── channels/
│   │   │   │   └── index.tsx  # Notification channels
│   │   │   ├── do-not-disturb/
│   │   │   │   └── index.tsx  # DND settings
│   │   │   └── index.tsx  # Notification preferences
│   │   │       # BE: communications-be/preferences
│   │   │       # GET/PUT /v1/notifications/preferences
│   │   └── data-usage/
│   │       └── index.tsx  # Data usage settings
│   │           # - Download on WiFi only
│   │           # - Auto-play videos
│   │           # - Image quality
│   │
│   ├── biometric/
│   │   └── index.tsx  # Biometric authentication
│   │       # - Face ID / Touch ID
│   │       # - Setup
│   │       # BE: users-be/auth (device trust)
│   │       # POST /v1/auth/device-trust
│   │
│   ├── storage/
│   │   ├── cache/
│   │   │   └── index.tsx  # Cache management
│   │   │       # - Clear cache
│   │   │       # - Cache size
│   │   └── downloads/
│   │       └── index.tsx  # Downloaded files
│   │           # - Manage downloads
│   │           # - Clear downloads
│   │
│   └── advanced/
│       ├── developer/
│       │   └── index.tsx  # Developer options
│       │       # - API endpoint override
│       │       # - Debug logging
│       │       # - Performance monitoring
│       └── experiments/
│           └── index.tsx  # Experimental features
│               # - Beta features toggle
│               # BE: utility/flags
│               # GET /v1/flags/user
```

---

## IV. Missing Shared Features

### 1. Feature Flags & Experiments (packages/shared/src/)

```
packages/shared/src/
│
├── features/
│   ├── flags/
│   │   ├── api/
│   │   │   └── flags-api.ts  # Feature flags API
│   │   │       # BE: utility/flags
│   │   │       # GET /v1/flags/user
│   │   │       # GET /v1/flags/organization
│   │   ├── hooks/
│   │   │   ├── useFeatureFlag.ts  # Single flag hook
│   │   │   ├── useFeatureFlags.ts  # Multiple flags
│   │   │   └── useFeatureFlagVariant.ts  # A/B variant
│   │   ├── providers/
│   │   │   └── FeatureFlagsProvider.tsx  # Context provider
│   │   ├── types.ts  # Flag types
│   │   └── utils.ts  # Flag utilities
│   │
│   └── experiments/
│       ├── api/
│       │   └── experiments-api.ts  # Experiments API
│       │       # BE: utility/experiments (if exists) or utility/flags
│       │       # GET /v1/experiments/active
│       │       # POST /v1/experiments/track-event
│       ├── hooks/
│       │   ├── useExperiment.ts  # Experiment hook
│       │   └── useExperimentTracking.ts  # Track experiment events
│       ├── providers/
│       │   └── ExperimentsProvider.tsx  # Experiments context
│       └── types.ts  # Experiment types
```

### 2. Real-time & WebSocket Features

```
packages/shared/src/
│
├── features/
│   ├── realtime/
│   │   ├── api/
│   │   │   └── websocket-client.ts  # WebSocket client
│   │   │       # BE: communications-be/websocket (if exists)
│   │   │       # WS: wss://api.skillsier.com/v1/ws
│   │   ├── hooks/
│   │   │   ├── useWebSocket.ts  # WebSocket hook
│   │   │   ├── usePresence.ts  # User presence
│   │   │   ├── useTypingIndicator.ts  # Typing indicator
│   │   │   └── useRealtimeNotifications.ts  # Real-time notifications
│   │   ├── providers/
│   │   │   └── WebSocketProvider.tsx  # WebSocket context
│   │   ├── types.ts  # WebSocket message types
│   │   └── utils.ts  # Connection management
│   │
│   └── presence/
│       ├── api/
│       │   └── presence-api.ts  # Presence API
│       │       # BE: communications-be/presence (if exists)
│       │       # POST /v1/presence/heartbeat
│       │       # GET /v1/presence/users/{user_id}
│       ├── hooks/
│       │   ├── useOnlineStatus.ts  # Online status
│       │   ├── useLastSeen.ts  # Last seen time
│       │   └── usePresenceSubscription.ts  # Subscribe to presence
│       └── types.ts  # Presence types
```

### 3. Geolocation & Maps Features

```
packages/shared/src/
│
├── features/
│   ├── geolocation/
│   │   ├── api/
│   │   │   └── geolocation-api.ts  # Geolocation API
│   │   │       # BE: utility/geo (if exists)
│   │   │       # GET /v1/geo/location
│   │   │       # GET /v1/geo/timezone
│   │   ├── hooks/
│   │   │   ├── useGeolocation.ts  # Get current location
│   │   │   ├── useTimezone.ts  # Detect timezone
│   │   │   ├── useCountryDetection.ts  # Detect country
│   │   │   └── useDistanceCalculation.ts  # Calculate distance
│   │   ├── types.ts  # Geolocation types
│   │   └── utils.ts  # Geo utilities
│   │
│   └── maps/
│       ├── components/
│       │   ├── Map.tsx  # Map component (web)
│       │   ├── Map.native.tsx  # Map component (mobile)
│       │   ├── Marker.tsx  # Map marker
│       │   └── LocationPicker.tsx  # Location picker
│       └── types.ts  # Map types
```

### 4. Content Moderation & Safety

```
packages/shared/src/
│
├── features/
│   ├── moderation/
│   │   ├── api/
│   │   │   └── moderation-api.ts  # Moderation API
│   │   │       # BE: admin-be/moderation (if exists)
│   │   │       # POST /v1/moderation/report
│   │   │       # POST /v1/moderation/content-check
│   │   ├── hooks/
│   │   │   ├── useContentModeration.ts  # Check content
│   │   │   ├── useReporting.ts  # Report content
│   │   │   └── useBlockUser.ts  # Block/unblock users
│   │   ├── types.ts  # Moderation types
│   │   └── utils.ts  # Content validation
│   │
│   └── safety/
│       ├── components/
│       │   ├── ReportModal.tsx  # Report modal
│       │   ├── BlockConfirmation.tsx  # Block confirmation
│       │   └── SafetyNotice.tsx  # Safety notice banner
│       └── utils.ts  # Safety utilities
```

### 5. Gamification & Achievements

```
packages/shared/src/
│
├── features/
│   ├── gamification/
│   │   ├── api/
│   │   │   └── gamification-api.ts  # Gamification API
│   │   │       # BE: users-be/gamification (if exists) or reviews-be/reputation
│   │   │       # GET /v1/gamification/points
│   │   │       # GET /v1/gamification/achievements
│   │   │       # GET /v1/gamification/leaderboard
│   │   ├── hooks/
│   │   │   ├── usePoints.ts  # User points
│   │   │   ├── useAchievements.ts  # Achievements
│   │   │   ├── useBadges.ts  # Badges
│   │   │   └── useLeaderboard.ts  # Leaderboard
│   │   ├── components/
│   │   │   ├── PointsDisplay.tsx  # Points display
│   │   │   ├── BadgeCollection.tsx  # Badge collection
│   │   │   ├── AchievementToast.tsx  # Achievement notification
│   │   │   └── LeaderboardWidget.tsx  # Leaderboard widget
│   │   └── types.ts  # Gamification types
│   │
│   └── achievements/
│       ├── definitions/
│       │   ├── freelancer-achievements.ts  # Freelancer achievements
│       │   ├── client-achievements.ts  # Client achievements
│       │   └── platform-achievements.ts  # Platform-wide achievements
│       └── utils.ts  # Achievement utilities
```

### 6. Performance Monitoring & Analytics

```
packages/shared/src/
│
├── features/
│   ├── performance/
│   │   ├── api/
│   │   │   └── performance-api.ts  # Performance API
│   │   │       # BE: utility/metrics (if exists) or utility/analytics
│   │   │       # POST /v1/metrics/performance
│   │   ├── hooks/
│   │   │   ├── usePerformanceMonitor.ts  # Monitor performance
│   │   │   ├── useWebVitals.ts  # Web Vitals
│   │   │   └── useErrorTracking.ts  # Error tracking
│   │   ├── utils/
│   │   │   ├── performance-observer.ts  # Performance observer
│   │   │   ├── web-vitals-reporter.ts  # Web Vitals reporter
│   │   │   └── error-reporter.ts  # Error reporter
│   │   └── types.ts  # Performance types
│   │
│   └── analytics/
│       ├── api/
│       │   └── analytics-api.ts  # Analytics API
│       │       # BE: utility/analytics
│       │       # POST /v1/analytics/events
│       │       # POST /v1/analytics/page-view
│       ├── hooks/
│       │   ├── useAnalytics.ts  # Track events
│       │   ├── usePageView.ts  # Track page views
│       │   └── useConversionTracking.ts  # Track conversions
│       ├── events/
│       │   ├── user-events.ts  # User events
│       │   ├── job-events.ts  # Job events
│       │   ├── proposal-events.ts  # Proposal events
│       │   ├── contract-events.ts  # Contract events
│       │   ├── payment-events.ts  # Payment events
│       │   └── system-events.ts  # System events
│       ├── utils/
│       │   ├── event-builder.ts  # Build events
│       │   ├── anonymize.ts  # Anonymize PII
│       │   └── batch-sender.ts  # Batch events
│       └── types.ts  # Analytics types
```

---

## V. Missing UI Components

### 1. Advanced UI Components (packages/ui/src/)

```
packages/ui/src/components/
│
├── Charts/
│   ├── AreaChart/
│   │   ├── AreaChart.tsx  # Base area chart
│   │   ├── AreaChart.web.tsx  # Web implementation
│   │   ├── AreaChart.native.tsx  # Native implementation
│   │   └── AreaChart.types.ts  # Types
│   ├── BarChart/
│   │   ├── BarChart.tsx
│   │   ├── BarChart.web.tsx
│   │   ├── BarChart.native.tsx
│   │   └── BarChart.types.ts
│   ├── LineChart/
│   │   ├── LineChart.tsx
│   │   ├── LineChart.web.tsx
│   │   ├── LineChart.native.tsx
│   │   └── LineChart.types.ts
│   ├── PieChart/
│   │   ├── PieChart.tsx
│   │   ├── PieChart.web.tsx
│   │   ├── PieChart.native.tsx
│   │   └── PieChart.types.ts
│   └── Sparkline/
│       ├── Sparkline.tsx
│       ├── Sparkline.web.tsx
│       ├── Sparkline.native.tsx
│       └── Sparkline.types.ts
│
├── Calendar/
│   ├── Calendar.tsx  # Full calendar
│   ├── Calendar.web.tsx
│   ├── Calendar.native.tsx
│   ├── DatePicker/
│   │   ├── DatePicker.tsx
│   │   ├── DatePicker.web.tsx
│   │   └── DatePicker.native.tsx
│   ├── DateRangePicker/
│   │   ├── DateRangePicker.tsx
│   │   ├── DateRangePicker.web.tsx
│   │   └── DateRangePicker.native.tsx
│   ├── TimePicker/
│   │   ├── TimePicker.tsx
│   │   ├── TimePicker.web.tsx
│   │   └── TimePicker.native.tsx
│   └── Calendar.types.ts
│
├── Editor/
│   ├── RichTextEditor/
│   │   ├── RichTextEditor.tsx
│   │   ├── RichTextEditor.web.tsx  # Web (e.g., TipTap/Slate)
│   │   ├── RichTextEditor.native.tsx  # Native implementation
│   │   └── RichTextEditor.types.ts
│   ├── MarkdownEditor/
│   │   ├── MarkdownEditor.tsx
│   │   ├── MarkdownEditor.web.tsx
│   │   ├── MarkdownEditor.native.tsx
│   │   └── MarkdownEditor.types.ts
│   └── CodeEditor/
│       ├── CodeEditor.tsx
│       ├── CodeEditor.web.tsx  # Web (e.g., Monaco/CodeMirror)
│       ├── CodeEditor.native.tsx
│       └── CodeEditor.types.ts
│
├── FileUpload/
│   ├── FileUpload.tsx  # Base file upload
│   ├── FileUpload.web.tsx  # Drag & drop support
│   ├── FileUpload.native.tsx  # Document picker
│   ├── ImageUpload/
│   │   ├── ImageUpload.tsx
│   │   ├── ImageUpload.web.tsx
│   │   ├── ImageUpload.native.tsx  # Camera/gallery
│   │   └── ImageCropper.tsx  # Image cropping
│   ├── MultiFileUpload/
│   │   ├── MultiFileUpload.tsx
│   │   ├── MultiFileUpload.web.tsx
│   │   └── MultiFileUpload.native.tsx
│   └── FileUpload.types.ts
│
├── DataDisplay/
│   ├── Table/
│   │   ├── Table.tsx  # Base table
│   │   ├── Table.web.tsx  # Full featured table
│   │   ├── DataTable/
│   │   │   ├── DataTable.tsx  # With sorting/filtering
│   │   │   └── DataTable.types.ts
│   │   ├── VirtualTable/
│   │   │   ├── VirtualTable.tsx  # Virtualized table
│   │   │   └── VirtualTable.types.ts
│   │   └── Table.types.ts
│   ├── List/
│   │   ├── List.tsx
│   │   ├── List.web.tsx
│   │   ├── List.native.tsx
│   │   ├── VirtualList/
│   │   │   ├── VirtualList.tsx
│   │   │   ├── VirtualList.web.tsx
│   │   │   └── VirtualList.native.tsx  # FlashList
│   │   └── List.types.ts
│   ├── Timeline/
│   │   ├── Timeline.tsx
│   │   ├── Timeline.web.tsx
│   │   ├── Timeline.native.tsx
│   │   └── Timeline.types.ts
│   └── Kanban/
│       ├── KanbanBoard.tsx
│       ├── KanbanBoard.web.tsx  # Drag & drop
│       ├── KanbanBoard.native.tsx  # Touch gestures
│       └── KanbanBoard.types.ts
│
├── Navigation/
│   ├── Breadcrumb/
│   │   ├── Breadcrumb.tsx
│   │   ├── Breadcrumb.web.tsx
│   │   └── Breadcrumb.types.ts
│   ├── Tabs/
│   │   ├── Tabs.tsx
│   │   ├── Tabs.web.tsx
│   │   ├── Tabs.native.tsx
│   │   └── Tabs.types.ts
│   ├── Stepper/
│   │   ├── Stepper.tsx
│   │   ├── Stepper.web.tsx
│   │   ├── Stepper.native.tsx
│   │   └── Stepper.types.ts
│   └── Pagination/
│       ├── Pagination.tsx
│       ├── Pagination.web.tsx
│       ├── Pagination.native.tsx
│       └── Pagination.types.ts
│
├── Feedback/
│   ├── Alert/
│   │   ├── Alert.tsx
│   │   ├── Alert.web.tsx
│   │   ├── Alert.native.tsx
│   │   └── Alert.types.ts
│   ├── Notification/
│   │   ├── Notification.tsx
│   │   ├── Notification.web.tsx  # Toast notifications
│   │   ├── Notification.native.tsx
│   │   └── Notification.types.ts
│   ├── Progress/
│   │   ├── ProgressBar/
│   │   │   ├── ProgressBar.tsx
│   │   │   ├── ProgressBar.web.tsx
│   │   │   └── ProgressBar.native.tsx
│   │   ├── ProgressCircle/
│   │   │   ├── ProgressCircle.tsx
│   │   │   ├── ProgressCircle.web.tsx
│   │   │   └── ProgressCircle.native.tsx
│   │   └── Progress.types.ts
│   ├── Skeleton/
│   │   ├── Skeleton.tsx
│   │   ├── Skeleton.web.tsx
│   │   ├── Skeleton.native.tsx
│   │   └── Skeleton.types.ts
│   └── EmptyState/
│       ├── EmptyState.tsx
│       ├── EmptyState.web.tsx
│       ├── EmptyState.native.tsx
│       └── EmptyState.types.ts
│
└── Forms/
    ├── Select/
    │   ├── Select.tsx
    │   ├── Select.web.tsx
    │   ├── Select.native.tsx
    │   ├── MultiSelect/
    │   │   ├── MultiSelect.tsx
    │   │   ├── MultiSelect.web.tsx
    │   │   └── MultiSelect.native.tsx
    │   └── Select.types.ts
    ├── Checkbox/
    │   ├── Checkbox.tsx
    │   ├── Checkbox.web.tsx
    │   ├── Checkbox.native.tsx
    │   └── Checkbox.types.ts
    ├── Radio/
    │   ├── RadioGroup.tsx
    │   ├── RadioGroup.web.tsx
    │   ├── RadioGroup.native.tsx
    │   └── Radio.types.ts
    ├── Switch/
    │   ├── Switch.tsx
    │   ├── Switch.web.tsx
    │   ├── Switch.native.tsx
    │   └── Switch.types.ts
    ├── Slider/
    │   ├── Slider.tsx
    │   ├── Slider.web.tsx
    │   ├── Slider.native.tsx
    │   ├── RangeSlider/
    │   │   ├── RangeSlider.tsx
    │   │   ├── RangeSlider.web.tsx
    │   │   └── RangeSlider.native.tsx
    │   └── Slider.types.ts
    └── FormField/
        ├── FormField.tsx  # Form field wrapper
        ├── FormField.web.tsx
        ├── FormField.native.tsx
        ├── FormLabel.tsx
        ├── FormError.tsx
        ├── FormHelperText.tsx
        └── FormField.types.ts
```

---

## VI. Missing Type Definitions

### 1. Additional Domain Types (packages/types/src/)

```
packages/types/src/
│
├── domains/
│   ├── activity/
│   │   └── index.ts  # Activity feed types
│   │       # - Activity event
│   │       # - Activity feed item
│   │       # - Activity filters
│   ├── achievements/
│   │   └── index.ts  # Achievement types
│   │       # - Achievement definition
│   │       # - User achievement
│   │       # - Badge
│   │       # - Points
│   ├── compliance/
│   │   └── index.ts  # Compliance types
│   │       # - KYC case
│   │       # - Business verification
│   │       # - Document verification
│   │       # - Privacy request
│   ├── experiments/
│   │   └── index.ts  # Experiment types
│   │       # - Experiment definition
│   │       # - Variant
│   │       # - User assignment
│   ├── flags/
│   │   └── index.ts  # Feature flag types
│   │       # - Flag definition
│   │       # - Flag value
│   │       # - Targeting rules
│   ├── incidents/
│   │   └── index.ts  # Incident types
│   │       # - Incident
│   │       # - Maintenance window
│   │       # - System health
│   ├── learning/
│   │   └── index.ts  # Learning types
│   │       # - Course
│   │       # - Lesson
│   │       # - Assessment
│   │       # - Progress
│   ├── moderation/
│   │   └── index.ts  # Moderation types
│   │       # - Report
│   │       # - Moderation action
│   │       # - Content check
│   ├── presence/
│   │   └── index.ts  # Presence types
│   │       # - Online status
│   │       # - Last seen
│   │       # - Typing indicator
│   └── sourcing/
│       └── index.ts  # Sourcing types
│           # - Campaign
│           # - Talent pool
│           # - Invitation
│
├── api/
│   ├── webhooks/
│   │   └── index.ts  # Webhook types
│   │       # - Webhook event
│   │       # - Webhook subscription
│   │       # - Webhook delivery
│   ├── realtime/
│   │   └── index.ts  # Real-time types
│   │       # - WebSocket message
│   │       # - Presence update
│   │       # - Event subscription
│   └── performance/
│       └── index.ts  # Performance types
│           # - Web Vitals
│           # - Performance metric
│           # - Error event
│
└── shared/
    ├── geolocation/
    │   └── index.ts  # Geolocation types
    │       # - Coordinates
    │       # - Location
    │       # - Distance
    └── analytics/
        └── index.ts  # Analytics types
            # - Event
            # - Page view
            # - Conversion
```

---

## VII. Missing API Clients

### 1. Additional API Clients (packages/shared/src/api/)

```
packages/shared/src/api/
│
├── activity/
│   ├── activity-client.ts  # Activity API client
│   │   # BE: utility/activity (if exists) or communications-be/notification
│   │   # GET /v1/activity/feed
│   │   # GET /v1/activity/user/{user_id}
│   └── types.ts
│
├── compliance/
│   ├── compliance-client.ts  # Compliance API client
│   │   # BE: admin-be/privacy, admin-be/kyc_case
│   │   # GET /v1/privacy/export-requests
│   │   # POST /v1/privacy/export-requests/{id}/process
│   │   # GET /v1/privacy/deletion-requests
│   │   # POST /v1/privacy/deletion-requests/{id}/process
│   └── types.ts
│
├── experiments/
│   ├── experiments-client.ts  # Experiments API client
│   │   # BE: utility/experiments (if exists) or utility/flags
│   │   # GET /v1/experiments/active
│   │   # POST /v1/experiments/{id}/track
│   └── types.ts
│
├── flags/
│   ├── flags-client.ts  # Feature flags API client
│   │   # BE: utility/flags
│   │   # GET /v1/flags/user
│   │   # GET /v1/flags/organization
│   └── types.ts
│
├── incidents/
│   ├── incidents-client.ts  # Incidents API client
│   │   # BE: utility/status
│   │   # GET /v1/status/incidents
│   │   # POST /v1/status/incidents
│   │   # PUT /v1/status/incidents/{id}
│   │   # GET /v1/status/maintenance
│   └── types.ts
│
├── learning/
│   ├── learning-client.ts  # Learning API client
│   │   # BE: learning-be (if exists) or external LMS
│   │   # GET /v1/learning/courses
│   │   # GET /v1/learning/courses/{id}
│   │   # POST /v1/learning/assessments/{id}/submit
│   └── types.ts
│
├── moderation/
│   ├── moderation-client.ts  # Moderation API client
│   │   # BE: admin-be/moderation (if exists)
│   │   # POST /v1/moderation/report
│   │   # POST /v1/moderation/content-check
│   └── types.ts
│
├── presence/
│   ├── presence-client.ts  # Presence API client
│   │   # BE: communications-be/presence (if exists)
│   │   # POST /v1/presence/heartbeat
│   │   # GET /v1/presence/users/{id}
│   └── types.ts
│
├── sourcing/
│   ├── sourcing-client.ts  # Sourcing API client
│   │   # BE: jobs-be/campaign (if exists) or jobs-be/job
│   │   # GET /v1/sourcing/campaigns
│   │   # POST /v1/sourcing/campaigns
│   │   # GET /v1/sourcing/talent-pools
│   └── types.ts
│
└── webhooks/
    ├── webhooks-client.ts  # Webhooks API client
    │   # BE: utility/webhooks (if exists) or admin-be
    │   # GET /v1/webhooks
    │   # POST /v1/webhooks
    │   # PUT /v1/webhooks/{id}
    │   # DELETE /v1/webhooks/{id}
    └── types.ts
```

---

## VIII. Testing Infrastructure

### 1. Test Utilities (packages/shared/src/testing/)

```
packages/shared/src/testing/
│
├── mocks/
│   ├── api/
│   │   ├── users-mock.ts  # Mock users API
│   │   ├── jobs-mock.ts  # Mock jobs API
│   │   ├── proposals-mock.ts  # Mock proposals API
│   │   ├── contracts-mock.ts  # Mock contracts API
│   │   ├── financial-mock.ts  # Mock financial API
│   │   ├── communications-mock.ts  # Mock communications API
│   │   └── ...  # Other service mocks
│   ├── data/
│   │   ├── users.ts  # Mock user data
│   │   ├── jobs.ts  # Mock job data
│   │   ├── proposals.ts  # Mock proposal data
│   │   └── ...  # Other mock data
│   └── handlers/
│       ├── auth-handlers.ts  # MSW auth handlers
│       ├── users-handlers.ts  # MSW users handlers
│       └── ...  # Other MSW handlers
│
├── factories/
│   ├── user-factory.ts  # User factory
│   ├── job-factory.ts  # Job factory
│   ├── proposal-factory.ts  # Proposal factory
│   └── ...  # Other factories
│
├── fixtures/
│   ├── auth.ts  # Auth fixtures
│   ├── users.ts  # User fixtures
│   ├── jobs.ts  # Job fixtures
│   └── ...  # Other fixtures
│
├── setup/
│   ├── test-setup.ts  # Test environment setup
│   ├── msw-setup.ts  # MSW server setup
│   └── test-providers.tsx  # Test providers wrapper
│
└── utils/
    ├── render-with-providers.tsx  # Custom render
    ├── wait-for-async.ts  # Async utilities
    └── accessibility-checker.ts  # A11y testing utils
```

### 2. E2E Test Structure

```
tests/
│
├── e2e/
│   ├── web/
│   │   ├── auth/
│   │   │   ├── login.spec.ts
│   │   │   ├── register.spec.ts
│   │   │   └── password-reset.spec.ts
│   │   ├── freelancer/
│   │   │   ├── profile.spec.ts
│   │   │   ├── proposals.spec.ts
│   │   │   ├── contracts.spec.ts
│   │   │   └── earnings.spec.ts
│   │   ├── client/
│   │   │   ├── job-posting.spec.ts
│   │   │   ├── talent-search.spec.ts
│   │   │   ├── hiring.spec.ts
│   │   │   └── payments.spec.ts
│   │   └── admin/
│   │       ├── kyc.spec.ts
│   │       ├── moderation.spec.ts
│   │       └── financial-ops.spec.ts
│   │
│   ├── mobile/
│   │   ├── auth/
│   │   ├── freelancer/
│   │   ├── client/
│   │   └── offline/
│   │
│   └── fixtures/
│       ├── users.json
│       ├── jobs.json
│       └── ...
│
├── integration/
│   ├── features/
│   │   ├── auth.test.ts
│   │   ├── jobs.test.ts
│   │   ├── proposals.test.ts
│   │   └── ...
│   └── api/
│       ├── users-api.test.ts
│       ├── jobs-api.test.ts
│       └── ...
│
└── performance/
    ├── lighthouse/
    │   ├── home.spec.ts
    │   ├── dashboard.spec.ts
    │   └── job-detail.spec.ts
    └── load/
        ├── api-load.spec.ts
        └── user-simulation.spec.ts
```

---

## IX. Documentation & Scripts

### 1. Documentation Structure

```
docs/
│
├── architecture/
│   ├── frontend-architecture.md
│   ├── microservices-integration.md
│   ├── state-management.md
│   ├── routing-strategy.md
│   ├── authentication-flow.md
│   └── data-fetching-patterns.md
│
├── api-integration/
│   ├── users-be.md
│   ├── jobs-be.md
│   ├── proposals-be.md
│   ├── contracts-be.md
│   ├── financial-be.md
│   ├── communications-be.md
│   ├── search-be.md
│   ├── storage-be.md
│   ├── reviews-be.md
│   ├── admin-be.md
│   ├── subscriptions-be.md
│   └── utility-services.md
│
├── features/
│   ├── authentication.md
│   ├── job-management.md
│   ├── proposal-system.md
│   ├── contract-management.md
│   ├── payment-system.md
│   ├── messaging.md
│   ├── notifications.md
│   ├── search-discovery.md
│   └── admin-tools.md
│
├── components/
│   ├── component-library.md
│   ├── design-system.md
│   ├── theming.md
│   └── accessibility.md
│
├── guides/
│   ├── getting-started.md
│   ├── development-workflow.md
│   ├── testing-guide.md
│   ├── deployment.md
│   ├── troubleshooting.md
│   └── contributing.md
│
└── adr/  # Architecture Decision Records
    ├── 001-monorepo-structure.md
    ├── 002-state-management.md
    ├── 003-authentication-approach.md
    ├── 004-component-library.md
    └── ...
```

### 2. Scripts & Automation

```
scripts/
│
├── dev/
│   ├── start-all.sh  # Start all services
│   ├── start-web.sh  # Start web only
│   ├── start-mobile.sh  # Start mobile only
│   └── reset-db.sh  # Reset local DB
│
├── build/
│   ├── build-all.sh  # Build all apps
│   ├── build-web.sh  # Build web
│   ├── build-mobile.sh  # Build mobile
│   └── analyze-bundle.sh  # Bundle analysis
│
├── test/
│   ├── test-all.sh  # Run all tests
│   ├── test-unit.sh  # Unit tests
│   ├── test-integration.sh  # Integration tests
│   ├── test-e2e-web.sh  # E2E web tests
│   ├── test-e2e-mobile.sh  # E2E mobile tests
│   └── test-coverage.sh  # Coverage report
│
├── ci/
│   ├── pre-commit.sh  # Pre-commit hook
│   ├── pre-push.sh  # Pre-push hook
│   ├── verify-types.sh  # Type check
│   └── lint-all.sh  # Lint all code
│
├── deploy/
│   ├── deploy-web-staging.sh
│   ├── deploy-web-prod.sh
│   ├── deploy-mobile-staging.sh
│   └── deploy-mobile-prod.sh
│
├── maintenance/
│   ├── clean-all.sh  # Clean all artifacts
│   ├── update-deps.sh  # Update dependencies
│   ├── generate-types.sh  # Generate API types
│   └── sync-translations.sh  # Sync i18n
│
└── utils/
    ├── check-ports.sh  # Check port availability
    ├── setup-env.sh  # Setup environment
    └── doctor.sh  # Environment health check
```

---

## X. CI/CD Configuration

### 1. GitHub Actions Workflows

```
.github/
│
├── workflows/
│   ├── ci-web.yml  # Web CI pipeline
│   │   # - Lint
│   │   # - Type check
│   │   # - Unit tests
│   │   # - Build
│   │   # - Bundle size check
│   │
│   ├── ci-mobile.yml  # Mobile CI pipeline
│   │   # - Lint
│   │   # - Type check
│   │   # - Unit tests
│   │   # - Build (iOS/Android)
│   │
│   ├── cd-web-staging.yml  # Web staging deployment
│   │   # - Build
│   │   # - Deploy to staging
│   │   # - Run smoke tests
│   │
│   ├── cd-web-production.yml  # Web production deployment
│   │   # - Build
│   │   # - Deploy to production
│   │   # - Run smoke tests
│   │   # - Notify team
│   │
│   ├── cd-mobile-staging.yml  # Mobile staging deployment
│   │   # - Build
│   │   # - Submit to TestFlight/Internal Testing
│   │
│   ├── cd-mobile-production.yml  # Mobile production deployment
│   │   # - Build
│   │   # - Submit to App Store/Play Store
│   │
│   ├── e2e-tests.yml  # E2E tests
│   │   # - Setup test environment
│   │   # - Run Playwright/Detox tests
│   │   # - Upload test results
│   │
│   ├── lighthouse.yml  # Performance audits
│   │   # - Run Lighthouse CI
│   │   # - Compare against budgets
│   │   # - Comment on PR
│   │
│   ├── accessibility.yml  # Accessibility checks
│   │   # - Run axe-core
│   │   # - Check WCAG compliance
│   │
│   ├── security.yml  # Security scanning
│   │   # - Dependency audit
│   │   # - SAST scan
│   │   # - License check
│   │
│   ├── dependency-review.yml  # Dependency review
│   │   # - Check for vulnerabilities
│   │   # - License compliance
│   │
│   └── release.yml  # Release automation
│       # - Create changelog
│       # - Tag release
│       # - Create GitHub release
│
├── actions/
│   ├── setup-node/
│   │   └── action.yml
│   ├── cache-deps/
│   │   └── action.yml
│   ├── build-web/
│   │   └── action.yml
│   ├── build-mobile/
│   │   └── action.yml
│   └── deploy-preview/
│       └── action.yml
│
└── CODEOWNERS  # Code ownership
```

---

## XI. Missing Environment Configuration

### 1. Environment Files

```
fe/
│
├── .env.example  # Template
├── .env.development  # Local development
├── .env.staging  # Staging environment
├── .env.production  # Production environment
│
└── config/
    ├── environments/
    │   ├── development.ts  # Dev config
    │   ├── staging.ts  # Staging config
    │   └── production.ts  # Production config
    │
    ├── feature-flags.ts  # Feature flags config
    │
    └── constants/
        ├── api-endpoints.ts  # API endpoint constants
        ├── app-config.ts  # App configuration
        ├── performance-budgets.ts  # Performance budgets
        └── third-party-keys.ts  # Third-party service keys
```

---

## XII. Monitoring & Observability Setup

### 1. Monitoring Configuration

```
packages/shared/src/monitoring/
│
├── sentry/
│   ├── sentry-config.ts  # Sentry configuration
│   ├── sentry-init.web.ts  # Web initialization
│   ├── sentry-init.native.ts  # Mobile initialization
│   └── error-boundary.tsx  # Sentry error boundary
│
├── analytics/
│   ├── analytics-config.ts  # Analytics configuration
│   ├── providers/
│   │   ├── google-analytics.ts  # GA4
│   │   ├── mixpanel.ts  # Mixpanel
│   │   └── amplitude.ts  # Amplitude
│   └── tracking/
│       ├── page-tracker.ts
│       ├── event-tracker.ts
│       └── user-tracker.ts
│
├── performance/
│   ├── web-vitals.ts  # Web Vitals monitoring
│   ├── performance-observer.ts  # Performance API
│   ├── bundle-monitor.ts  # Bundle size monitoring
│   └── api-monitor.ts  # API performance monitoring
│
└── logging/
    ├── logger.ts  # Application logger
    ├── log-levels.ts  # Log level configuration
    └── transports/
        ├── console-transport.ts
        ├── remote-transport.ts
        └── file-transport.ts  # Mobile only
```

---

## XIII. Security Infrastructure

### 1. Security Configuration

```
packages/shared/src/security/
│
├── auth/
│   ├── token-manager.ts  # Token management
│   ├── token-refresh.ts  # Auto token refresh
│   ├── session-manager.ts  # Session management
│   └── device-trust.ts  # Device trust
│
├── encryption/
│   ├── crypto-utils.ts  # Encryption utilities
│   ├── secure-storage.ts  # Secure storage
│   └── key-management.ts  # Key management
│
├── validation/
│   ├── input-validator.ts  # Input validation
│   ├── sanitizer.ts  # HTML/input sanitization
│   ├── xss-protection.ts  # XSS protection
│   └── sql-injection-guard.ts  # SQL injection protection
│
├── headers/
│   ├── security-headers.ts  # Security headers config
│   ├── csp.ts  # Content Security Policy
│   └── cors.ts  # CORS configuration
│
└── monitoring/
    ├── security-monitor.ts  # Security monitoring
    ├── breach-detector.ts  # Breach detection
    └── anomaly-detector.ts  # Anomaly detection
```

### 2. Web Security Middleware

```
apps/web/
│
├── middleware.ts  # Next.js middleware
│   # - Security headers
│   # - CSRF protection
│   # - Rate limiting (client-side)
│   # - Auth checks
│
└── security/
    ├── csrf-protection.ts
    ├── rate-limiter.ts
    └── auth-guard.ts
```

---

## XIV. Accessibility Infrastructure

### 1. Accessibility Utilities

```
packages/shared/src/accessibility/
│
├── a11y-utils.ts  # Accessibility utilities
├── screen-reader.ts  # Screen reader utilities
├── focus-management.ts  # Focus management
├── keyboard-navigation.ts  # Keyboard navigation
├── aria-utils.ts  # ARIA utilities
│
└── testing/
    ├── a11y-test-utils.ts  # Testing utilities
    └── axe-config.ts  # axe-core configuration
```

### 2. Accessibility Components

```
packages/ui/src/a11y/
│
├── SkipLink/
│   ├── SkipLink.tsx
│   ├── SkipLink.web.tsx
│   └── SkipLink.types.ts
│
├── VisuallyHidden/
│   ├── VisuallyHidden.tsx
│   ├── VisuallyHidden.web.tsx
│   ├── VisuallyHidden.native.tsx
│   └── VisuallyHidden.types.ts
│
├── FocusTrap/
│   ├── FocusTrap.tsx
│   ├── FocusTrap.web.tsx
│   └── FocusTrap.types.ts
│
└── Announcer/
    ├── LiveAnnouncer.tsx  # Live region announcements
    ├── LiveAnnouncer.web.tsx
    ├── LiveAnnouncer.native.tsx
    └── LiveAnnouncer.types.ts
```

---

## XV. Summary of Missing Microservice Endpoints

Based on this comprehensive folder structure, here are the **missing backend endpoints** that need to be implemented or clarified:

### 1. **Utility Services (Missing Domains)**
- `utility/activity` - Activity feed service
- `utility/experiments` - A/B testing service
- `utility/webhooks` - Webhook management service
- `utility/geo` - Geolocation service

### 2. **Communications-BE (Missing Domains)**
- `communications-be/presence` - User presence/online status
- `communications-be/websocket` - WebSocket real-time connections
- `communications-be/invitation` - Direct invitations to jobs

### 3. **Admin-BE (Missing Domains)**
- `admin-be/privacy` - GDPR/privacy request management
- `admin-be/data_retention` - Data retention policies
- `admin-be/moderation` - Content moderation
- `admin-be/integrations` - Third-party integrations management
- `admin-be/platform_config` - Platform-wide configuration

### 4. **Users-BE (Missing Domains)**
- `users-be/availability` - Freelancer availability calendar
- `users-be/talent_pool` - Client talent pool management
- `users-be/gamification` - Points, badges, achievements

### 5. **Jobs-BE (Missing Domains)**
- `jobs-be/campaign` - Sourcing campaigns
- `jobs-be/template` - Job posting templates

### 6. **Search-BE (Missing Domains)**
- `search-be/alert` - Job alerts with notifications

### 7. **Financial-BE (Missing Domains)**
- `financial-be/pricing_config` - Platform pricing configuration
- `financial-be/forecast` - Spend forecasting
- `financial-be/budget` - Budget management and tracking

### 8. **New Microservice Needed**
- `learning-be` - Learning platform integration
  - Courses
  - Lessons
  - Assessments
  - Progress tracking
  - Achievements

### 9. **Reviews-BE (Missing Domains)**
- `reviews-be/reputation` - Reputation scoring system
- `reviews-be/badge` - Achievement badges
- `reviews-be/dispute` - Review dispute resolution

---

## XVI. Backend Integration Notes

### Required Backend Enhancements:

1. **Real-time Infrastructure**
   - WebSocket server for live updates
   - Presence/typing indicators
   - Live notifications

2. **Privacy & Compliance**
   - GDPR export/deletion automation
   - Data retention scheduler
   - Consent management system

3. **Platform Configuration**
   - Feature flags service
   - A/B testing framework
   - Platform limits/quotas

4. **Advanced Analytics**
   - Activity feed aggregation
   - Performance metrics
   - User behavior tracking

5. **Talent Sourcing**
   - Campaign management
   - Talent pool curation
   - Direct invitation system

6. **Learning Integration**
   - Course catalog
   - Assessment engine
   - Progress tracking
   - Certificate issuance

---

## Conclusion

This document completes the **final missing frontend folder structure** for the Skillsier application. All routes, components, features, API clients, and utilities are now mapped to their corresponding backend microservices with proper endpoint documentation.

### Coverage Summary:
- ✅ All admin routes completed
- ✅ All user journey routes completed
- ✅ All mobile-specific features completed
- ✅ All shared features and utilities completed
- ✅ All UI components cataloged
- ✅ All type definitions completed
- ✅ All API clients documented
- ✅ Testing infrastructure defined
- ✅ CI/CD pipelines configured
- ✅ Monitoring and observability setup
- ✅ Security infrastructure completed
- ✅ Accessibility infrastructure completed
- ✅ Missing backend endpoints identified

### Next Steps:
1. Review and implement missing backend endpoints
2. Generate OpenAPI specs for all services
3. Implement frontend components following this structure
4. Setup CI/CD pipelines
5. Configure monitoring and observability
6. Conduct security audit
7. Perform accessibility audit
