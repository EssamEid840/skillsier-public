fe/
├── .husky/
│   ├── pre-commit
│   └── pre-push
├── .vscode/
│   ├── extensions.json
│   ├── launch.json
│   └── settings.json
├── .github/
│   ├── actions/
│   │   ├── build-mobile/
│   │   │   └── action.yml
│   │   ├── build-web/
│   │   │   └── action.yml
│   │   ├── cache-deps/
│   │   │   └── action.yml
│   │   ├── deploy-preview/
│   │   │   └── action.yml
│   │   └── setup-node/
│   │       └── action.yml
│   ├── workflows/
│   │   ├── accessibility.yml  # Accessibility checks
│   │   │   # - Run axe-core
│   │   │   # - Check WCAG compliance
│   │   ├── cd-mobile-production.yml  # Mobile production deployment
│   │   │   # - Build
│   │   │   # - Submit to App Store/Play Store
│   │   ├── cd-mobile-staging.yml  # Mobile staging deployment
│   │   │   # - Build
│   │   │   # - Submit to TestFlight/Internal Testing
│   │   ├── cd-web-production.yml  # Web production deployment
│   │   │   # - Build
│   │   │   # - Deploy to production
│   │   │   # - Run smoke tests
│   │   │   # - Notify team
│   │   ├── cd-web-staging.yml  # Web staging deployment
│   │   │   # - Build
│   │   │   # - Deploy to staging
│   │   │   # - Run smoke tests
│   │   ├── ci-mobile.yml  # Mobile CI pipeline
│   │   │   # - Lint
│   │   │   # - Type check
│   │   │   # - Unit tests
│   │   │   # - Build (iOS/Android)
│   │   ├── ci-web.yml  # Web CI pipeline
│   │   │   # - Lint
│   │   │   # - Type check
│   │   │   # - Unit tests
│   │   │   # - Build
│   │   │   # - Bundle size check
│   │   ├── dependency-review.yml  # Dependency review
│   │   │   # - Check for vulnerabilities
│   │   │   # - License compliance
│   │   ├── e2e-tests.yml  # E2E tests
│   │   │   # - Setup test environment
│   │   │   # - Run Playwright/Detox tests
│   │   │   # - Upload test results
│   │   ├── lighthouse.yml  # Performance audits
│   │   │   # - Run Lighthouse CI
│   │   │   # - Compare against budgets
│   │   │   # - Comment on PR
│   │   ├── release.yml  # Release automation
│   │   │   # - Create changelog
│   │   │   # - Tag release
│   │   │   # - Create GitHub release
│   │   └── security.yml  # Security scanning
│   │       # - Dependency audit
│   │       # - SAST scan
│   │       # - License check
│   └── CODEOWNERS  # Code ownership
├── apps/
│   ├── mobile/
│   │   ├── app/
│   │   │   ├── (auth)/
│   │   │   │   ├── _layout.tsx
│   │   │   │   ├── callback.tsx
│   │   │   │   ├── login.tsx
│   │   │   │   └── register.tsx
│   │   │   ├── (dashboard)/
│   │   │   │   ├── activity-feed/
│   │   │   │   │   └── index.tsx  # Activity feed
│   │   │   │   │       # - Recent activities
│   │   │   │   │       # - Notifications inline
│   │   │   │   │       # - Quick actions from feed
│   │   │   │   │       # BE: communications-be/notification, utility/activity
│   │   │   │   │       # GET /v1/activity/feed
│   │   │   │   ├── today/
│   │   │   │   │   └── index.tsx  # Today view
│   │   │   │   │       # - Today's schedule
│   │   │   │   │       # - Pending tasks
│   │   │   │   │       # - Quick metrics
│   │   │   │   │       # BE: contracts-be/work_diary, communications-be/notification
│   │   │   │   │       # GET /v1/today/overview
│   │   │   │   ├── widgets/
│   │   │   │   │   ├── earnings/
│   │   │   │   │   │   └── index.tsx  # Earnings widget
│   │   │   │   │   │       # BE: financial-be/wallet
│   │   │   │   │   ├── notifications/
│   │   │   │   │   │   └── index.tsx  # Notifications widget
│   │   │   │   │   │       # BE: communications-be/notification
│   │   │   │   │   └── time-tracker/
│   │   │   │   │       └── index.tsx  # Time tracker widget
│   │   │   │   │           # BE: contracts-be/work_diary
│   │   │   │   └── quick-actions/
│   │   │   │       └── index.tsx  # Quick actions screen
│   │   │   │           # - Quick message
│   │   │   │           # - Quick proposal
│   │   │   │           # - Quick time entry
│   │   │   │           # - Quick invoice
│   │   │   │           # BE: Multiple services
│   │   │   ├── accessibility/
│   │   │   │   ├── screen-reader/
│   │   │   │   │   └── index.tsx  # Screen reader optimized view
│   │   │   │   └── voice-commands/
│   │   │   │       └── index.tsx  # Voice commands
│   │   │   │           # - Voice-to-text
│   │   │   │           # - Command shortcuts
│   │   │   ├── camera/
│   │   │   │   ├── document-scan/
│   │   │   │   │   └── index.tsx  # Document scanner
│   │   │   │   │       # - ID verification
│   │   │   │   │       # - Contract documents
│   │   │   │   │       # - Expense receipts
│   │   │   │   │       # BE: storage-be/asset, admin-be/business_verification
│   │   │   │   │       # POST /v1/storage/uploads (scan)
│   │   │   │   ├── photo-upload/
│   │   │   │   │   └── index.tsx  # Photo upload
│   │   │   │   │       # - Portfolio photos
│   │   │   │   │       # - Work progress photos
│   │   │   │   │       # BE: storage-be/asset
│   │   │   │   │       # POST /v1/storage/uploads
│   │   │   │   └── qr-scan/
│   │   │   │       └── index.tsx  # QR code scanner
│   │   │   │           # - Profile sharing
│   │   │   │           # - Event check-in
│   │   │   │           # BE: users-be/profile (QR data)
│   │   │   ├── offline/
│   │   │   │   ├── queue/
│   │   │   │   │   └── index.tsx  # Offline queue
│   │   │   │   │       # - Pending uploads
│   │   │   │   │       # - Queued messages
│   │   │   │   │       # - Draft proposals
│   │   │   │   │       # BE: None (local storage)
│   │   │   │   ├── settings/
│   │   │   │   │   └── index.tsx  # Offline settings
│   │   │   │   │       # - Auto-sync preferences
│   │   │   │   │       # - Download for offline
│   │   │   │   │       # BE: None (local storage)
│   │   │   │   └── sync/
│   │   │   │       └── index.tsx  # Sync status
│   │   │   │           # - Sync progress
│   │   │   │           # - Conflict resolution
│   │   │   │           # BE: Multiple services (sync endpoints)
│   │   │   ├── settings/
│   │   │   │   ├── advanced/
│   │   │   │   │   ├── developer/
│   │   │   │   │   │   └── index.tsx  # Developer options
│   │   │   │   │   │       # - API endpoint override
│   │   │   │   │   │       # - Debug logging
│   │   │   │   │   │       # - Performance monitoring
│   │   │   │   │   └── experiments/
│   │   │   │   │       └── index.tsx  # Experimental features
│   │   │   │   │           # - Beta features toggle
│   │   │   │   │           # BE: utility/flags
│   │   │   │   │           # GET /v1/flags/user
│   │   │   │   ├── app-preferences/
│   │   │   │   │   ├── data-usage/
│   │   │   │   │   │   └── index.tsx  # Data usage settings
│   │   │   │   │   │       # - Download on WiFi only
│   │   │   │   │   │       # - Auto-play videos
│   │   │   │   │   │       # - Image quality
│   │   │   │   │   ├── language/
│   │   │   │   │   │   └── index.tsx  # Language preferences
│   │   │   │   │   │       # - App language
│   │   │   │   │   │       # - Content language
│   │   │   │   │   ├── notifications/
│   │   │   │   │   │   ├── channels/
│   │   │   │   │   │   │   └── index.tsx  # Notification channels
│   │   │   │   │   │   ├── do-not-disturb/
│   │   │   │   │   │   │   └── index.tsx  # DND settings
│   │   │   │   │   │   └── index.tsx  # Notification preferences
│   │   │   │   │   │       # BE: communications-be/preferences
│   │   │   │   │   │       # GET/PUT /v1/notifications/preferences
│   │   │   │   │   └── theme/
│   │   │   │   │       └── index.tsx  # Theme settings
│   │   │   │   │           # - Light/dark/auto
│   │   │   │   │           # - Accent color
│   │   │   │   ├── biometric/
│   │   │   │   │   └── index.tsx  # Biometric authentication
│   │   │   │   │       # - Face ID / Touch ID
│   │   │   │   │       # - Setup
│   │   │   │   │       # BE: users-be/auth (device trust)
│   │   │   │   │       # POST /v1/auth/device-trust
│   │   │   │   └── storage/
│   │   │   │       ├── cache/
│   │   │   │       │   └── index.tsx  # Cache management
│   │   │   │       │       # - Clear cache
│   │   │   │       │       # - Cache size
│   │   │   │       └── downloads/
│   │   │   │           └── index.tsx  # Downloaded files
│   │   │   │               # - Manage downloads
│   │   │   │               # - Clear downloads
│   │   │   ├── +not-found.tsx
│   │   │   ├── _layout.tsx
│   │   │   └── index.tsx
│   │   └── src/
│   └── web/
│       ├── security/
│       │   ├── auth-guard.ts
│       │   ├── csrf-protection.ts
│       │   └── rate-limiter.ts
│       ├── middleware.ts  # Next.js middleware
│       │   # - Security headers
│       │   # - CSRF protection
│       │   # - Rate limiting (client-side)
│       │   # - Auth checks
│       └── src/
│           └── app/
│               └── [locale]/
│                   ├── (admin)/
│                   │   ├── compliance/
│                   │   │   ├── aml-kyc/
│                   │   │   │   ├── monitoring/
│                   │   │   │   │   └── page.tsx  # AML monitoring dashboard
│                   │   │   │   │       # - Suspicious activity
│                   │   │   │   │       # - Transaction patterns
│                   │   │   │   │       # - Risk scores
│                   │   │   │   │       # BE: admin-be/kyc_case, financial-be/transaction
│                   │   │   │   │       # GET /v1/kyc/monitoring/suspicious-activity
│                   │   │   │   ├── reports/
│                   │   │   │   │   ├── [reportId]/
│                   │   │   │   │   │   └── page.tsx  # SAR (Suspicious Activity Report) detail
│                   │   │   │   │   │       # BE: admin-be/kyc_case
│                   │   │   │   │   │       # GET /v1/kyc/reports/{report_id}
│                   │   │   │   │   └── page.tsx  # AML reports list
│                   │   │   │   │       # - Filed reports
│                   │   │   │   │       # - Pending reports
│                   │   │   │   │       # BE: admin-be/kyc_case
│                   │   │   │   │       # GET /v1/kyc/reports
│                   │   │   │   └── risk-assessment/
│                   │   │   │       └── page.tsx  # Risk assessment tools
│                   │   │   │           # - User risk profiles
│                   │   │   │           # - Country risk matrix
│                   │   │   │           # - Enhanced due diligence
│                   │   │   │           # BE: admin-be/kyc_case
│                   │   │   │           # GET /v1/kyc/risk-assessment
│                   │   │   ├── data-retention/
│                   │   │   │   ├── audit/
│                   │   │   │   │   └── page.tsx  # Retention audit log
│                   │   │   │   │       # - Deletion history
│                   │   │   │   │       # - Policy compliance
│                   │   │   │   │       # BE: utility/audit
│                   │   │   │   │       # GET /v1/audit/retention
│                   │   │   │   ├── policies/
│                   │   │   │   │   └── page.tsx  # Retention policies
│                   │   │   │   │       # - Policy definitions
│                   │   │   │   │       # - Data categories
│                   │   │   │   │       # - Retention periods
│                   │   │   │   │       # BE: admin-be/data_retention (if exists) or utility/config
│                   │   │   │   │       # GET /v1/retention/policies
│                   │   │   │   └── schedule/
│                   │   │   │       └── page.tsx  # Deletion schedule
│                   │   │   │           # - Upcoming deletions
│                   │   │   │           # - Retention expirations
│                   │   │   │           # BE: admin-be/data_retention
│                   │   │   │           # GET /v1/retention/schedule
│                   │   │   ├── document-verification/
│                   │   │   │   ├── automated-checks/
│                   │   │   │   │   └── page.tsx  # Automated verification rules
│                   │   │   │   │       # - OCR settings
│                   │   │   │   │       # - Validation rules
│                   │   │   │   │       # - ML model performance
│                   │   │   │   │       # BE: admin-be/business_verification
│                   │   │   │   │       # GET /v1/verification/automation-rules
│                   │   │   │   ├── queue/
│                   │   │   │   │   └── page.tsx  # Document verification queue
│                   │   │   │   │       # - Pending documents
│                   │   │   │   │       # - Priority sorting
│                   │   │   │   │       # - Auto-verification status
│                   │   │   │   │       # BE: admin-be/business_verification, storage-be/asset
│                   │   │   │   │       # GET /v1/verification/documents/queue
│                   │   │   │   └── [documentId]/
│                   │   │   │       └── page.tsx  # Document review interface
│                   │   │   │           # - Document viewer
│                   │   │   │           # - Verification checks
│                   │   │   │           # - Approve/reject/request-more
│                   │   │   │           # BE: admin-be/business_verification, storage-be/asset
│                   │   │   │           # GET /v1/verification/documents/{document_id}
│                   │   │   │           # PUT /v1/verification/documents/{document_id}/review
│                   │   │   ├── gdpr/
│                   │   │   │   ├── consent-management/
│                   │   │   │   │   └── page.tsx  # Consent logs & management
│                   │   │   │   │       # - User consent history
│                   │   │   │   │       # - Consent versions
│                   │   │   │   │       # - Audit trail
│                   │   │   │   │       # BE: users-be/consent, utility/audit
│                   │   │   │   │       # GET /v1/users/consent-logs
│                   │   │   │   ├── deletion-requests/
│                   │   │   │   │   ├── [requestId]/
│                   │   │   │   │   │   └── page.tsx  # Deletion request detail
│                   │   │   │   │   │       # - Data preview
│                   │   │   │   │   │       # - Retention check
│                   │   │   │   │   │       # - Process deletion
│                   │   │   │   │   │       # BE: admin-be/privacy, users-be/user
│                   │   │   │   │   │       # GET /v1/privacy/deletion-requests/{request_id}
│                   │   │   │   │   │       # POST /v1/privacy/deletion-requests/{request_id}/process
│                   │   │   │   │   └── page.tsx  # Deletion requests queue
│                   │   │   │   │       # BE: admin-be/privacy
│                   │   │   │   │       # GET /v1/privacy/deletion-requests
│                   │   │   │   ├── export-requests/
│                   │   │   │   │   ├── [requestId]/
│                   │   │   │   │   │   └── page.tsx  # Export request detail
│                   │   │   │   │   │       # - Request review
│                   │   │   │   │   │       # - Generate export
│                   │   │   │   │   │       # - Approve/deny
│                   │   │   │   │   │       # BE: admin-be/privacy, users-be/user
│                   │   │   │   │   │       # GET /v1/privacy/export-requests/{request_id}
│                   │   │   │   │   │       # POST /v1/privacy/export-requests/{request_id}/process
│                   │   │   │   │   └── page.tsx  # Export requests queue
│                   │   │   │   │       # BE: admin-be/privacy
│                   │   │   │   │       # GET /v1/privacy/export-requests
│                   │   │   │   └── reports/
│                   │   │   │       └── page.tsx  # GDPR compliance reports
│                   │   │   │           # - Processing activities
│                   │   │   │           # - Data inventory
│                   │   │   │           # - Breach reports
│                   │   │   │           # BE: admin-be/privacy, utility/audit
│                   │   │   │           # GET /v1/privacy/reports
│                   │   │   ├── incidents/
│                   │   │   │   ├── [incidentId]/
│                   │   │   │   │   ├── edit/
│                   │   │   │   │   │   └── page.tsx  # Edit incident details
│                   │   │   │   │   │       # - Update status
│                   │   │   │   │   │       # - Add postmortem
│                   │   │   │   │   │       # - Affected services
│                   │   │   │   │   │       # BE: utility/status
│                   │   │   │   │   │       # PUT /v1/status/incidents/{incident_id}
│                   │   │   │   │   ├── timeline/
│                   │   │   │   │   │   └── page.tsx  # Incident timeline
│                   │   │   │   │   │       # - Event log
│                   │   │   │   │   │       # - Update history
│                   │   │   │   │   │       # BE: utility/status
│                   │   │   │   │   │       # GET /v1/status/incidents/{incident_id}/timeline
│                   │   │   │   │   └── page.tsx  # Incident detail
│                   │   │   │   │       # - Current status
│                   │   │   │   │       # - Impact assessment
│                   │   │   │   │       # - Resolution steps
│                   │   │   │   │       # BE: utility/status
│                   │   │   │   │       # GET /v1/status/incidents/{incident_id}
│                   │   │   │   ├── create/
│                   │   │   │   │   └── page.tsx  # Create new incident
│                   │   │   │   │       # - Incident type selection
│                   │   │   │   │       # - Severity level
│                   │   │   │   │       # - Affected services
│                   │   │   │   │       # BE: utility/status
│                   │   │   │   │       # POST /v1/status/incidents
│                   │   │   │   ├── history/
│                   │   │   │   │   └── page.tsx  # Historical incidents
│                   │   │   │   │       # - Past incidents archive
│                   │   │   │   │       # - Postmortems
│                   │   │   │   │       # - Lessons learned
│                   │   │   │   │       # BE: utility/status
│                   │   │   │   │       # GET /v1/status/incidents/history
│                   │   │   │   └── page.tsx  # Active incidents dashboard
│                   │   │   │       # - Current incidents
│                   │   │   │       # - Quick actions
│                   │   │   │       # - Status board
│                   │   │   │       # BE: utility/status
│                   │   │   │       # GET /v1/status/incidents?status=active
│                   │   │   ├── maintenance/
│                   │   │   │   ├── [maintenanceId]/
│                   │   │   │   │   ├── edit/
│                   │   │   │   │   │   └── page.tsx  # Edit maintenance window
│                   │   │   │   │   │       # BE: utility/status
│                   │   │   │   │   │       # PUT /v1/status/maintenance/{maintenance_id}
│                   │   │   │   │   └── page.tsx  # Maintenance detail
│                   │   │   │   │       # BE: utility/status
│                   │   │   │   │       # GET /v1/status/maintenance/{maintenance_id}
│                   │   │   │   ├── schedule/
│                   │   │   │   │   └── page.tsx  # Schedule maintenance
│                   │   │   │   │       # - Date/time selection
│                   │   │   │   │       # - Affected services
│                   │   │   │   │       # - Notification plan
│                   │   │   │   │       # BE: utility/status
│                   │   │   │   │       # POST /v1/status/maintenance
│                   │   │   │   └── page.tsx  # Maintenance calendar
│                   │   │   │       # - Upcoming maintenance
│                   │   │   │       # - Impact windows
│                   │   │   │       # BE: utility/status
│                   │   │   │       # GET /v1/status/maintenance
│                   │   │   └── system-health/
│                   │   │       ├── metrics/
│                   │   │       │   └── page.tsx  # System metrics dashboard
│                   │   │       │       # - CPU/Memory usage
│                   │   │       │       # - Database performance
│                   │   │       │       # - Queue depths
│                   │   │       │       # BE: utility/metrics (or monitoring service)
│                   │   │       │       # GET /v1/metrics/system
│                   │   │       ├── services/
│                   │   │       │   └── page.tsx  # Service health overview
│                   │   │       │       # - All microservices status
│                   │   │       │       # - Uptime metrics
│                   │   │       │       # - Response times
│                   │   │       │       # BE: utility/status
│                   │   │       │       # GET /v1/status/services
│                   │   │       └── page.tsx  # Health dashboard
│                   │   │           # - Overall system health
│                   │   │           # - Critical alerts
│                   │   │           # - Performance trends
│                   │   │           # BE: utility/status
│                   │   │           # GET /v1/status/health
│                   │   └── platform-config/
│                   │       ├── integrations/
│                   │       │   ├── [integrationId]/
│                   │       │   │   ├── configure/
│                   │       │   │   │   └── page.tsx  # Configure integration
│                   │       │   │   │       # BE: admin-be/integrations (if exists)
│                   │       │   │   │       # PUT /v1/integrations/{integration_id}/config
│                   │       │   │   ├── logs/
│                   │       │   │   │   └── page.tsx  # Integration logs
│                   │       │   │   │       # BE: utility/audit
│                   │       │   │   │       # GET /v1/integrations/{integration_id}/logs
│                   │       │   │   └── page.tsx  # Integration detail
│                   │       │   │       # BE: admin-be/integrations
│                   │       │   │       # GET /v1/integrations/{integration_id}
│                   │       │   └── page.tsx  # Integrations list
│                   │       │       # - Payment providers
│                   │       │       # - Email services
│                   │       │       # - Storage providers
│                   │       │       # - Auth providers
│                   │       │       # BE: admin-be/integrations
│                   │       │       # GET /v1/integrations
│                   │       ├── localization/
│                   │       │   ├── languages/
│                   │       │   │   └── page.tsx  # Language management
│                   │       │   │       # - Enabled languages
│                   │       │   │       # - Default language
│                   │       │   │       # - RTL settings
│                   │       │   │       # BE: utility/i18n
│                   │       │   │       # GET /v1/config/languages
│                   │       │   │       # PUT /v1/config/languages
│                   │       │   └── regions/
│                   │       │       └── page.tsx  # Regional settings
│                   │       │           # - Timezone defaults
│                   │       │           # - Currency settings
│                   │       │           # - Date/time formats
│                   │       │           # BE: utility/config
│                   │       │           # GET /v1/config/regions
│                   │       │           # PUT /v1/config/regions
│                   │       ├── notifications/
│                   │       │   ├── settings/
│                   │       │   │   └── page.tsx  # Notification settings
│                   │       │   │       # - Default preferences
│                   │       │   │       # - Delivery channels
│                   │       │   │       # - Retry policies
│                   │       │   │       # BE: communications-be/config
│                   │       │   │       # GET /v1/notifications/config
│                   │       │   │       # PUT /v1/notifications/config
│                   │       │   └── templates/
│                   │       │       ├── [templateId]/
│                   │       │       │   ├── edit/
│                   │       │       │   │   └── page.tsx  # Edit notification template
│                   │       │       │   │       # BE: communications-be/template
│                   │       │       │   │       # PUT /v1/notifications/templates/{template_id}
│                   │       │       │   ├── preview/
│                   │       │       │   │   └── page.tsx  # Preview template
│                   │       │       │   │       # BE: communications-be/template
│                   │       │       │   │       # POST /v1/notifications/templates/{template_id}/preview
│                   │       │       │   └── page.tsx  # Template detail
│                   │       │       │       # BE: communications-be/template
│                   │       │       │       # GET /v1/notifications/templates/{template_id}
│                   │       │       └── page.tsx  # Template library
│                   │       │           # BE: communications-be/template
│                   │       │           # GET /v1/notifications/templates
│                   │       └── pricing/
│                   │           └── page.tsx  # Pricing configuration
│                   │               # - Commission rates
│                   │               # - Subscription pricing
│                   │               # - Regional pricing
│                   │               # BE: financial-be/pricing_config (if exists)
│                   │               # GET /v1/config/pricing
│                   │               # PUT /v1/config/pricing
│                   │               # Note: Requires change_approval
│                   └── (dashboard)/
│                       ├── availability/
│                       │   ├── calendar/
│                       │   │   └── page.tsx  # Availability calendar
│                       │   │       # - Mark available/busy
│                       │   │       # - Recurring patterns
│                       │   │       # - Sync with external calendar
│                       │   │       # BE: users-be/availability (if exists) or users-be/profile
│                       │   │       # GET /v1/users/me/availability
│                       │   │       # PUT /v1/users/me/availability
│                       │   ├── settings/
│                       │   │   └── page.tsx  # Availability settings
│                       │   │       # - Working hours
│                       │   │       # - Timezone preferences
│                       │   │       # - Auto-reply settings
│                       │   │       # BE: users-be/settings
│                       │   │       # PUT /v1/users/me/availability-settings
│                       │   └── page.tsx  # Availability dashboard
│                       │       # - Current status
│                       │       # - Upcoming commitments
│                       │       # BE: users-be/availability
│                       │       # GET /v1/users/me/availability-overview
│                       ├── job-alerts/
│                       │   ├── [alertId]/
│                       │   │   ├── edit/
│                       │   │   │   └── page.tsx  # Edit job alert
│                       │   │   │       # BE: search-be/alert (if exists) or search-be/saved-search
│                       │   │   │       # PUT /v1/job-alerts/{alert_id}
│                       │   │   ├── history/
│                       │   │   │   └── page.tsx  # Alert history
│                       │   │   │       # - Jobs matched
│                       │   │   │       # - Notifications sent
│                       │   │   │       # BE: search-be/alert
│                       │   │   │       # GET /v1/job-alerts/{alert_id}/history
│                       │   │   └── page.tsx  # Alert detail
│                       │   │       # BE: search-be/alert
│                       │   │       # GET /v1/job-alerts/{alert_id}
│                       │   ├── create/
│                       │   │   └── page.tsx  # Create job alert
│                       │   │       # - Search criteria
│                       │   │       # - Notification frequency
│                       │   │       # - Delivery channel
│                       │   │       # BE: search-be/alert
│                       │   │       # POST /v1/job-alerts
│                       │   └── page.tsx  # Job alerts list
│                       │       # - Active alerts
│                       │       # - Pause/resume
│                       │       # BE: search-be/alert
│                       │       # GET /v1/job-alerts
│                       ├── learning/
│                       │   ├── achievements/
│                       │   │   └── page.tsx  # Learning achievements
│                       │   │       # - Certificates earned
│                       │   │       # - Badges
│                       │   │       # - Skill verifications
│                       │   │       # BE: learning-be/achievement
│                       │   │       # GET /v1/learning/achievements
│                       │   ├── assessments/
│                       │   │   ├── [assessmentId]/
│                       │   │   │   ├── results/
│                       │   │   │   │   └── page.tsx  # Assessment results
│                       │   │   │   │       # BE: learning-be/assessment
│                       │   │   │   │       # GET /v1/learning/assessments/{assessment_id}/results
│                       │   │   │   ├── take/
│                       │   │   │   │   └── page.tsx  # Take assessment
│                       │   │   │   │       # BE: learning-be/assessment
│                       │   │   │   │       # POST /v1/learning/assessments/{assessment_id}/submit
│                       │   │   │   └── page.tsx  # Assessment detail
│                       │   │   │       # BE: learning-be/assessment
│                       │   │   │       # GET /v1/learning/assessments/{assessment_id}
│                       │   │   └── page.tsx  # Assessments list
│                       │   │       # BE: learning-be/assessment
│                       │   │       # GET /v1/learning/assessments
│                       │   └── courses/
│                       │       ├── [courseId]/
│                       │       │   ├── lessons/
│                       │       │   │   └── [lessonId]/
│                       │       │   │       └── page.tsx  # Lesson content
│                       │       │   │           # BE: learning-be/lesson (if exists) or external LMS
│                       │       │   │           # GET /v1/learning/courses/{course_id}/lessons/{lesson_id}
│                       │       │   ├── progress/
│                       │       │   │   └── page.tsx  # Course progress
│                       │       │   │       # BE: learning-be/progress
│                       │       │   │       # GET /v1/learning/courses/{course_id}/progress
│                       │       │   └── page.tsx  # Course detail
│                       │       │       # BE: learning-be/course
│                       │       │       # GET /v1/learning/courses/{course_id}
│                       │       └── page.tsx  # Courses catalog
│                       │           # BE: learning-be/course
│                       │           # GET /v1/learning/courses
│                       ├── reputation/
│                       │   ├── badges/
│                       │   │   └── page.tsx  # Achievement badges
│                       │   │       # - Top rated badge
│                       │   │       # - Rising talent
│                       │   │       # - Expert verified
│                       │   │       # BE: reviews-be/badge (if exists) or users-be/profile
│                       │   │       # GET /v1/reputation/badges
│                       │   ├── disputes/
│                       │   │   └── page.tsx  # Review disputes
│                       │   │       # - Disputed reviews
│                       │   │       # - Appeal status
│                       │   │       # BE: reviews-be/dispute (if exists) or admin-be/case_mgmt
│                       │   │       # GET /v1/reviews/disputes
│                       │   ├── reviews-given/
│                       │   │   └── page.tsx  # Reviews given
│                       │   │       # - Reviews I left
│                       │   │       # - Edit option (if within timeframe)
│                       │   │       # BE: reviews-be/review
│                       │   │       # GET /v1/reviews/given
│                       │   ├── reviews-received/
│                       │   │   └── page.tsx  # Reviews received
│                       │   │       # - Client reviews
│                       │   │       # - Response option
│                       │   │       # BE: reviews-be/review
│                       │   │       # GET /v1/reviews/received
│                       │   └── overview/
│                       │       └── page.tsx  # Reputation overview
│                       │           # - Overall score
│                       │           # - Score breakdown
│                       │           # - Recent changes
│                       │           # BE: reviews-be/reputation (if exists) or reviews-be/review
│                       │           # GET /v1/reputation/overview
│                       ├── sourcing/
│                       │   ├── campaigns/
│                       │   │   ├── [campaignId]/
│                       │   │   │   ├── analytics/
│                       │   │   │   │   └── page.tsx  # Campaign analytics
│                       │   │   │   │       # - Reach
│                       │   │   │   │       # - Engagement
│                       │   │   │   │       # - Conversions
│                       │   │   │   │       # BE: jobs-be/analytics
│                       │   │   │   │       # GET /v1/sourcing/campaigns/{campaign_id}/analytics
│                       │   │   │   ├── edit/
│                       │   │   │   │   └── page.tsx  # Edit sourcing campaign
│                       │   │   │   │       # BE: jobs-be/campaign (if exists) or jobs-be/job
│                       │   │   │   │       # PUT /v1/sourcing/campaigns/{campaign_id}
│                       │   │   │   └── page.tsx  # Campaign detail
│                       │   │   │       # BE: jobs-be/campaign
│                       │   │   │       # GET /v1/sourcing/campaigns/{campaign_id}
│                       │   │   ├── create/
│                       │   │   │   └── page.tsx  # Create sourcing campaign
│                       │   │   │       # - Target criteria
│                       │   │   │       # - Budget allocation
│                       │   │   │       # - Messaging
│                       │   │   │       # BE: jobs-be/campaign
│                       │   │   │       # POST /v1/sourcing/campaigns
│                       │   │   └── page.tsx  # Campaigns list
│                       │   │       # BE: jobs-be/campaign
│                       │   │       # GET /v1/sourcing/campaigns
│                       │   ├── invitations/
│                       │   │   ├── [invitationId]/
│                       │   │   │   └── page.tsx  # Invitation detail
│                       │   │   │       # - View status
│                       │   │   │       # - Resend invitation
│                       │   │   │       # BE: communications-be/invitation (if exists) or jobs-be/job
│                       │   │   │       # GET /v1/sourcing/invitations/{invitation_id}
│                       │   │   └── page.tsx  # Sent invitations
│                       │   │       # - Invitation history
│                       │   │       # - Response rates
│                       │   │       # BE: communications-be/invitation
│                       │   │       # GET /v1/sourcing/invitations
│                       │   └── talent-pools/
│                       │       ├── [poolId]/
│                       │       │   ├── edit/
│                       │       │   │   └── page.tsx  # Edit talent pool
│                       │       │   │       # BE: users-be/talent_pool (if exists) or search-be/saved-search
│                       │       │   │       # PUT /v1/sourcing/talent-pools/{pool_id}
│                       │       │   ├── members/
│                       │       │   │   └── page.tsx  # Pool members
│                       │       │   │       # - Add/remove members
│                       │       │   │       # - Bulk actions
│                       │       │   │       # BE: users-be/talent_pool
│                       │       │   │       # GET /v1/sourcing/talent-pools/{pool_id}/members
│                       │       │   └── page.tsx  # Talent pool detail
│                       │       │       # BE: users-be/talent_pool
│                       │       │       # GET /v1/sourcing/talent-pools/{pool_id}
│                       │       ├── create/
│                       │       │   └── page.tsx  # Create talent pool
│                       │       │       # BE: users-be/talent_pool
│                       │       │       # POST /v1/sourcing/talent-pools
│                       │       └── page.tsx  # Talent pools list
│                       │           # BE: users-be/talent_pool
│                       │           # GET /v1/sourcing/talent-pools
│                       ├── spend-analytics/
│                       │   ├── by-category/
│                       │   │   └── page.tsx  # Spending by category
│                       │   │       # - Job category breakdown
│                       │   │       # - Trend analysis
│                       │   │       # BE: financial-be/analytics (if exists) or financial-be/invoice
│                       │   │       # GET /v1/analytics/spend/by-category
│                       │   ├── by-department/
│                       │   │   └── page.tsx  # Spending by department
│                       │   │       # - Cost center breakdown
│                       │   │       # - Budget vs actual
│                       │   │       # BE: financial-be/budget, financial-be/invoice
│                       │   │       # GET /v1/analytics/spend/by-department
│                       │   ├── by-vendor/
│                       │   │   └── page.tsx  # Spending by vendor
│                       │   │       # - Top vendors
│                       │   │       # - Vendor concentration risk
│                       │   │       # BE: financial-be/invoice, users-be/profile
│                       │   │       # GET /v1/analytics/spend/by-vendor
│                       │   ├── forecasting/
│                       │   │   └── page.tsx  # Spend forecasting
│                       │   │       # - Projected spend
│                       │   │       # - Budget burn rate
│                       │   │       # - Alerts for overages
│                       │   │       # BE: financial-be/forecast (if exists) or financial-be/analytics
│                       │   │       # GET /v1/analytics/spend/forecast
│                       │   └── page.tsx  # Spend analytics dashboard
│                       │       # - Overview metrics
│                       │       # - Charts & visualizations
│                       │       # BE: financial-be/analytics
│                       │       # GET /v1/analytics/spend
│                       └── search/
│                           ├── freelancers/
│                           │   └── page.tsx  # Advanced freelancer search (client)
│                           │       # - Search by skills
│                           └── jobs/
│                               └── page.tsx  # Advanced job search
│                                   # - Full-text search
│                                   # - Faceted filters
│                                   # - Autocomplete suggestions
│                                   # - Search history
│                                   # - Save search
│                                   # BE: search-be/query
│                                   # POST /v1/search/jobs
│                                   # Body: { query, filters: {...}, sort, page }
│                                   # BE: search-be/suggestions
│                                   # GET /v1/suggestions?q={query}
├── config/
│   └── environments/
│       ├── development.ts  # Dev config
│       ├── production.ts  # Production config
│       └── staging.ts  # Staging config
├── docs/
│   ├── adr/
│   │   ├── 001-monorepo-structure.md
│   │   ├── 002-state-management.md
│   │   ├── 003-authentication-approach.md
│   │   ├── 004-component-library.md
│   │   └── ...
│   ├── api-integration/
│   │   ├── admin-be.md
│   │   ├── communications-be.md
│   │   ├── contracts-be.md
│   │   ├── financial-be.md
│   │   ├── jobs-be.md
│   │   ├── proposals-be.md
│   │   ├── reviews-be.md
│   │   ├── search-be.md
│   │   ├── storage-be.md
│   │   ├── subscriptions-be.md
│   │   ├── users-be.md
│   │   └── utility-services.md
│   ├── architecture/
│   │   ├── authentication-flow.md
│   │   ├── data-fetching-patterns.md
│   │   ├── frontend-architecture.md
│   │   ├── microservices-integration.md
│   │   ├── routing-strategy.md
│   │   └── state-management.md
│   ├── components/
│   │   ├── accessibility.md
│   │   ├── component-library.md
│   │   ├── design-system.md
│   │   └── theming.md
│   ├── features/
│   │   ├── admin-tools.md
│   │   ├── authentication.md
│   │   ├── contract-management.md
│   │   ├── job-management.md
│   │   ├── messaging.md
│   │   ├── notifications.md
│   │   ├── payment-system.md
│   │   ├── proposal-system.md
│   │   └── search-discovery.md
│   └── guides/
│       ├── contributing.md
│       ├── deployment.md
│       ├── development-workflow.md
│       ├── getting-started.md
│       ├── testing-guide.md
│       └── troubleshooting.md
├── packages/
│   ├── shared/
│   │   └── src/
│   │       ├── accessibility/
│   │       │   ├── a11y-utils.ts  # Accessibility utilities
│   │       │   ├── aria-utils.ts  # ARIA utilities
│   │       │   ├── focus-management.ts  # Focus management
│   │       │   ├── keyboard-navigation.ts  # Keyboard navigation
│   │       │   ├── screen-reader.ts  # Screen reader utilities
│   │       │   └── testing/
│   │       │       ├── a11y-test-utils.ts  # Testing utilities
│   │       │       └── axe-config.ts  # axe-core configuration
│   │       ├── api/
│   │       │   ├── activity/
│   │       │   │   ├── activity-client.ts  # Activity API client
│   │       │   │   │   # BE: utility/activity (if exists) or communications-be/notification
│   │       │   │   │   # GET /v1/activity/feed
│   │       │   │   │   # GET /v1/activity/user/{user_id}
│   │       │   │   └── types.ts
│   │       │   ├── compliance/
│   │       │   │   ├── compliance-client.ts  # Compliance API client
│   │       │   │   │   # BE: admin-be/privacy, admin-be/kyc_case
│   │       │   │   │   # GET /v1/privacy/export-requests
│   │       │   │   │   # POST /v1/privacy/export-requests/{id}/process
│   │       │   │   │   # GET /v1/privacy/deletion-requests
│   │       │   │   │   # POST /v1/privacy/deletion-requests/{id}/process
│   │       │   │   └── types.ts
│   │       │   ├── experiments/
│   │       │   │   ├── experiments-client.ts  # Experiments API client
│   │       │   │   │   # BE: utility/experiments (if exists) or utility/flags
│   │       │   │   │   # GET /v1/experiments/active
│   │       │   │   │   # POST /v1/experiments/{id}/track
│   │       │   │   └── types.ts
│   │       │   ├── flags/
│   │       │   │   ├── flags-client.ts  # Feature flags API client
│   │       │   │   │   # BE: utility/flags
│   │       │   │   │   # GET /v1/flags/user
│   │       │   │   │   # GET /v1/flags/organization
│   │       │   │   └── types.ts
│   │       │   ├── incidents/
│   │       │   │   ├── incidents-client.ts  # Incidents API client
│   │       │   │   │   # BE: utility/status
│   │       │   │   │   # GET /v1/status/incidents
│   │       │   │   │   # POST /v1/status/incidents
│   │       │   │   │   # PUT /v1/status/incidents/{id}
│   │       │   │   │   # GET /v1/status/maintenance
│   │       │   │   └── types.ts
│   │       │   ├── learning/
│   │       │   │   ├── learning-client.ts  # Learning API client
│   │       │   │   │   # BE: learning-be (if exists) or external LMS
│   │       │   │   │   # GET /v1/learning/courses
│   │       │   │   │   # GET /v1/learning/courses/{id}
│   │       │   │   │   # POST /v1/learning/assessments/{id}/submit
│   │       │   │   └── types.ts
│   │       │   ├── moderation/
│   │       │   │   ├── moderation-client.ts  # Moderation API client
│   │       │   │   │   # BE: admin-be/moderation (if exists)
│   │       │   │   │   # POST /v1/moderation/report
│   │       │   │   │   # POST /v1/moderation/content-check
│   │       │   │   └── types.ts
│   │       │   ├── presence/
│   │       │   │   ├── presence-client.ts  # Presence API client
│   │       │   │   │   # BE: communications-be/presence (if exists)
│   │       │   │   │   # POST /v1/presence/heartbeat
│   │       │   │   │   # GET /v1/presence/users/{id}
│   │       │   │   └── types.ts
│   │       │   ├── sourcing/
│   │       │   │   ├── sourcing-client.ts  # Sourcing API client
│   │       │   │   │   # BE: jobs-be/campaign (if exists) or jobs-be/job
│   │       │   │   │   # GET /v1/sourcing/campaigns
│   │       │   │   │   # POST /v1/sourcing/campaigns
│   │       │   │   │   # GET /v1/sourcing/talent-pools
│   │       │   │   └── types.ts
│   │       │   └── webhooks/
│   │       │       ├── types.ts
│   │       │       └── webhooks-client.ts  # Webhooks API client
│   │       │           # BE: utility/webhooks (if exists) or admin-be
│   │       │           # GET /v1/webhooks
│   │       │           # POST /v1/webhooks
│   │       │           # PUT /v1/webhooks/{id}
│   │       │           # DELETE /v1/webhooks/{id}
│   │       ├── features/
│   │       │   ├── experiments/
│   │       │   │   ├── api/
│   │       │   │   │   └── experiments-api.ts  # Experiments API
│   │       │   │   │       # BE: utility/experiments (if exists) or utility/flags
│   │       │   │   │       # GET /v1/experiments/active
│   │       │   │   │       # POST /v1/experiments/track-event
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useExperiment.ts  # Experiment hook
│   │       │   │   │   └── useExperimentTracking.ts  # Track experiment events
│   │       │   │   ├── providers/
│   │       │   │   │   └── ExperimentsProvider.tsx  # Experiments context
│   │       │   │   └── types.ts  # Experiment types
│   │       │   ├── flags/
│   │       │   │   ├── api/
│   │       │   │   │   └── flags-api.ts  # Feature flags API
│   │       │   │   │       # BE: utility/flags
│   │       │   │   │       # GET /v1/flags/user
│   │       │   │   │       # GET /v1/flags/organization
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useFeatureFlag.ts  # Single flag hook
│   │       │   │   │   ├── useFeatureFlags.ts  # Multiple flags
│   │       │   │   │   └── useFeatureFlagVariant.ts  # A/B variant
│   │       │   │   ├── providers/
│   │       │   │   │   └── FeatureFlagsProvider.tsx  # Context provider
│   │       │   │   ├── types.ts  # Flag types
│   │       │   │   └── utils.ts  # Flag utilities
│   │       │   ├── geolocation/
│   │       │   │   ├── api/
│   │       │   │   │   └── geolocation-api.ts  # Geolocation API
│   │       │   │   │       # BE: utility/geo (if exists)
│   │       │   │   │       # GET /v1/geo/location
│   │       │   │   │       # GET /v1/geo/timezone
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useCountryDetection.ts  # Detect country
│   │       │   │   │   ├── useDistanceCalculation.ts  # Calculate distance
│   │       │   │   │   ├── useGeolocation.ts  # Get current location
│   │       │   │   │   └── useTimezone.ts  # Detect timezone
│   │       │   │   ├── types.ts  # Geolocation types
│   │       │   │   └── utils.ts  # Geo utilities
│   │       │   ├── moderation/
│   │       │   │   ├── api/
│   │       │   │   │   └── moderation-api.ts  # Moderation API
│   │       │   │   │       # BE: admin-be/moderation (if exists)
│   │       │   │   │       # POST /v1/moderation/report
│   │       │   │   │       # POST /v1/moderation/content-check
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useBlockUser.ts  # Block/unblock users
│   │       │   │   │   ├── useContentModeration.ts  # Check content
│   │       │   │   │   └── useReporting.ts  # Report content
│   │       │   │   ├── types.ts  # Moderation types
│   │       │   │   └── utils.ts  # Content validation
│   │       │   ├── presence/
│   │       │   │   ├── api/
│   │       │   │   │   └── presence-api.ts  # Presence API
│   │       │   │   │       # BE: communications-be/presence (if exists)
│   │       │   │   │       # POST /v1/presence/heartbeat
│   │       │   │   │       # GET /v1/presence/users/{user_id}
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useLastSeen.ts  # Last seen time
│   │       │   │   │   ├── useOnlineStatus.ts  # Online status
│   │       │   │   │   └── usePresenceSubscription.ts  # Subscribe to presence
│   │       │   │   └── types.ts  # Presence types
│   │       │   ├── realtime/
│   │       │   │   ├── api/
│   │       │   │   │   └── websocket-client.ts  # WebSocket client
│   │       │   │   │       # BE: communications-be/websocket (if exists)
│   │       │   │   │       # WS: wss://api.skillsier.com/v1/ws
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── usePresence.ts  # User presence
│   │       │   │   │   ├── useRealtimeNotifications.ts  # Real-time notifications
│   │       │   │   │   ├── useTypingIndicator.ts  # Typing indicator
│   │       │   │   │   └── useWebSocket.ts  # WebSocket hook
│   │       │   │   ├── providers/
│   │       │   │   │   └── WebSocketProvider.tsx  # WebSocket context
│   │       │   │   ├── types.ts  # WebSocket message types
│   │       │   │   └── utils.ts  # Connection management
│   │       │   ├── safety/
│   │       │   │   ├── components/
│   │       │   │   │   ├── BlockConfirmation.tsx  # Block confirmation
│   │       │   │   │   ├── ReportModal.tsx  # Report modal
│   │       │   │   │   └── SafetyNotice.tsx  # Safety notice banner
│   │       │   │   └── utils.ts  # Safety utilities
│   │       │   ├── gamification/
│   │       │   │   ├── api/
│   │       │   │   │   └── gamification-api.ts  # Gamification API
│   │       │   │   │       # BE: users-be/gamification (if exists) or reviews-be/reputation
│   │       │   │   │       # GET /v1/gamification/points
│   │       │   │   │       # GET /v1/gamification/achievements
│   │       │   │   │       # GET /v1/gamification/leaderboard
│   │       │   │   ├── components/
│   │       │   │   │   ├── AchievementToast.tsx  # Achievement notification
│   │       │   │   │   ├── BadgeCollection.tsx  # Badge collection
│   │       │   │   │   ├── LeaderboardWidget.tsx  # Leaderboard widget
│   │       │   │   │   └── PointsDisplay.tsx  # Points display
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useAchievements.ts  # Achievements
│   │       │   │   │   ├── useBadges.ts  # Badges
│   │       │   │   │   ├── useLeaderboard.ts  # Leaderboard
│   │       │   │   │   └── usePoints.ts  # User points
│   │       │   │   └── types.ts  # Gamification types
│   │       │   ├── achievements/
│   │       │   │   ├── definitions/
│   │       │   │   │   ├── client-achievements.ts  # Client achievements
│   │       │   │   │   ├── freelancer-achievements.ts  # Freelancer achievements
│   │       │   │   │   └── platform-achievements.ts  # Platform-wide achievements
│   │       │   │   └── utils.ts  # Achievement utilities
│   │       │   ├── performance/
│   │       │   │   ├── api/
│   │       │   │   │   └── performance-api.ts  # Performance API
│   │       │   │   │       # BE: utility/metrics (if exists) or utility/analytics
│   │       │   │   │       # POST /v1/metrics/performance
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useErrorTracking.ts  # Error tracking
│   │       │   │   │   ├── usePerformanceMonitor.ts  # Monitor performance
│   │       │   │   │   └── useWebVitals.ts  # Web Vitals
│   │       │   │   ├── types.ts  # Performance types
│   │       │   │   └── utils/
│   │       │   │       ├── error-reporter.ts  # Error reporter
│   │       │   │       ├── performance-observer.ts  # Performance observer
│   │       │   │       └── web-vitals-reporter.ts  # Web Vitals reporter
│   │       │   └── analytics/
│   │       │       ├── api/
│   │       │       │   └── analytics-api.ts  # Analytics API
│   │       │       │       # BE: utility/analytics
│   │       │       │       # POST /v1/analytics/events
│   │       │       │       # POST /v1/analytics/page-view
│   │       │       ├── events/
│   │       │       │   ├── contract-events.ts  # Contract events
│   │       │       │   ├── job-events.ts  # Job events
│   │       │       │   ├── payment-events.ts  # Payment events
│   │       │       │   ├── proposal-events.ts  # Proposal events
│   │       │       │   ├── system-events.ts  # System events
│   │       │       │   └── user-events.ts  # User events
│   │       │       ├── hooks/
│   │       │       │   ├── useAnalytics.ts  # Track events
│   │       │       │   ├── useConversionTracking.ts  # Track conversions
│   │       │       │   └── usePageView.ts  # Track page views
│   │       │       ├── types.ts  # Analytics types
│   │       │       └── utils/
│   │       │           ├── anonymize.ts  # Anonymize PII
│   │       │           ├── batch-sender.ts  # Batch events
│   │       │           └── event-builder.ts  # Build events
│   │       ├── monitoring/
│   │       │   ├── analytics/
│   │       │   │   ├── analytics-config.ts  # Analytics configuration
│   │       │   │   ├── providers/
│   │       │   │   │   ├── amplitude.ts  # Amplitude
│   │       │   │   │   ├── google-analytics.ts  # GA4
│   │       │   │   │   └── mixpanel.ts  # Mixpanel
│   │       │   │   └── tracking/
│   │       │   │       ├── event-tracker.ts
│   │       │   │       ├── page-tracker.ts
│   │       │   │       └── user-tracker.ts
│   │       │   ├── logging/
│   │       │   │   ├── log-levels.ts  # Log level configuration
│   │       │   │   ├── logger.ts  # Application logger
│   │       │   │   └── transports/
│   │       │   │       ├── console-transport.ts
│   │       │   │       ├── file-transport.ts  # Mobile only
│   │       │   │       └── remote-transport.ts
│   │       │   ├── performance/
│   │       │   │   ├── api-monitor.ts  # API performance monitoring
│   │       │   │   ├── bundle-monitor.ts  # Bundle size monitoring
│   │       │   │   ├── performance-observer.ts  # Performance API
│   │       │   │   └── web-vitals.ts  # Web Vitals monitoring
│   │       │   └── sentry/
│   │       │       ├── error-boundary.tsx  # Sentry error boundary
│   │       │       ├── sentry-config.ts  # Sentry configuration
│   │       │       ├── sentry-init.native.ts  # Mobile initialization
│   │       │       └── sentry-init.web.ts  # Web initialization
│   │       ├── security/
│   │       │   ├── auth/
│   │       │   │   ├── device-trust.ts  # Device trust
│   │       │   │   ├── session-manager.ts  # Session management
│   │       │   │   ├── token-manager.ts  # Token management
│   │       │   │   └── token-refresh.ts  # Auto token refresh
│   │       │   ├── encryption/
│   │       │   │   ├── crypto-utils.ts  # Encryption utilities
│   │       │   │   ├── key-management.ts  # Key management
│   │       │   │   └── secure-storage.ts  # Secure storage
│   │       │   ├── headers/
│   │       │   │   ├── cors.ts  # CORS configuration
│   │       │   │   ├── csp.ts  # Content Security Policy
│   │       │   │   └── security-headers.ts  # Security headers config
│   │       │   ├── monitoring/
│   │       │   │   ├── anomaly-detector.ts  # Anomaly detection
│   │       │   │   ├── breach-detector.ts  # Breach detection
│   │       │   │   └── security-monitor.ts  # Security monitoring
│   │       │   └── validation/
│   │       │       ├── input-validator.ts  # Input validation
│   │       │       ├── sanitizer.ts  # HTML/input sanitization
│   │       │       ├── sql-injection-guard.ts  # SQL injection protection
│   │       │       └── xss-protection.ts  # XSS protection
│   │       └── testing/
│   │           ├── factories/
│   │           │   ├── job-factory.ts  # Job factory
│   │           │   ├── proposal-factory.ts  # Proposal factory
│   │           │   └── user-factory.ts  # User factory
│   │           ├── fixtures/
│   │           │   ├── auth.ts  # Auth fixtures
│   │           │   ├── jobs.ts  # Job fixtures
│   │           │   └── users.ts  # User fixtures
│   │           ├── mocks/
│   │           │   ├── api/
│   │           │   │   ├── communications-mock.ts  # Mock communications API
│   │           │   │   ├── contracts-mock.ts  # Mock contracts API
│   │           │   │   ├── financial-mock.ts  # Mock financial API
│   │           │   │   ├── jobs-mock.ts  # Mock jobs API
│   │           │   │   ├── proposals-mock.ts  # Mock proposals API
│   │           │   │   └── users-mock.ts  # Mock users API
│   │           │   ├── data/
│   │           │   │   ├── jobs.ts  # Mock job data
│   │           │   │   ├── proposals.ts  # Mock proposal data
│   │           │   │   └── users.ts  # Mock user data
│   │           │   └── handlers/
│   │           │       ├── auth-handlers.ts  # MSW auth handlers
│   │           │       └── users-handlers.ts  # MSW users handlers
│   │           ├── setup/
│   │           │   ├── msw-setup.ts  # MSW server setup
│   │           │   ├── test-providers.tsx  # Test providers wrapper
│   │           │   └── test-setup.ts  # Test environment setup
│   │           └── utils/
│   │               ├── accessibility-checker.ts  # A11y testing utils
│   │               ├── render-with-providers.tsx  # Custom render
│   │               └── wait-for-async.ts  # Async utilities
│   ├── types/
│   │   └── src/
│   │       ├── api/
│   │       │   ├── performance/
│   │       │   │   └── index.ts  # Performance types
│   │       │   │       # - Web Vitals
│   │       │   │       # - Performance metric
│   │       │   │       # - Error event
│   │       │   └── webhooks/
│   │       │       └── index.ts  # Webhook types
│   │       │           # - Webhook event
│   │       │           # - Webhook subscription
│   │       │           # - Webhook delivery
│   │       ├── domains/
│   │       │   ├── achievements/
│   │       │   │   └── index.ts  # Achievement types
│   │       │   │       # - Achievement definition
│   │       │   │       # - User achievement
│   │       │   │       # - Badge
│   │       │   │       # - Points
│   │       │   ├── activity/
│   │       │   │   └── index.ts  # Activity feed types
│   │       │   │       # - Activity event
│   │       │   │       # - Activity feed item
│   │       │   │       # - Activity filters
│   │       │   ├── compliance/
│   │       │   │   └── index.ts  # Compliance types
│   │       │   │       # - KYC case
│   │       │   │       # - Business verification
│   │       │   │       # - Document verification
│   │       │   │       # - Privacy request
│   │       │   ├── experiments/
│   │       │   │   └── index.ts  # Experiment types
│   │       │   │       # - Experiment definition
│   │       │   │       # - Variant
│   │       │   │       # - User assignment
│   │       │   ├── flags/
│   │       │   │   └── index.ts  # Feature flag types
│   │       │   │       # - Flag definition
│   │       │   │       # - Flag value
│   │       │   │       # - Targeting rules
│   │       │   ├── incidents/
│   │       │   │   └── index.ts  # Incident types
│   │       │   │       # - Incident
│   │       │   │       # - Maintenance window
│   │       │   │       # - System health
│   │       │   ├── learning/
│   │       │   │   └── index.ts  # Learning types
│   │       │   │       # - Course
│   │       │   │       # - Lesson
│   │       │   │       # - Assessment
│   │       │   │       # - Progress
│   │       │   ├── moderation/
│   │       │   │   └── index.ts  # Moderation types
│   │       │   │       # - Report
│   │       │   │       # - Moderation action
│   │       │   │       # - Content check
│   │       │   ├── presence/
│   │       │   │   └── index.ts  # Presence types
│   │       │   │       # - Online status
│   │       │   │       # - Last seen
│   │       │   │       # - Typing indicator
│   │       │   └── sourcing/
│   │       │       └── index.ts  # Sourcing types
│   │       │           # - Campaign
│   │       │           # - Talent pool
│   │       │           # - Invitation
│   │       └── shared/
│   │           ├── analytics/
│   │           │   └── index.ts  # Analytics types
│   │           │       # - Event
│   │           │       # - Page view
│   │           │       # - Conversion
│   │           └── geolocation/
│   │               └── index.ts  # Geolocation types
│   │                   # - Coordinates
│   │                   # - Location
│   │                   # - Distance
│   └── ui/
│       └── src/
│           ├── a11y/
│           │   ├── Announcer/
│           │   │   ├── LiveAnnouncer.native.tsx
│           │   │   ├── LiveAnnouncer.tsx  # Live region announcements
│           │   │   ├── LiveAnnouncer.types.ts
│           │   │   └── LiveAnnouncer.web.tsx
│           │   ├── FocusTrap/
│           │   │   ├── FocusTrap.tsx
│           │   │   ├── FocusTrap.types.ts
│           │   │   └── FocusTrap.web.tsx
│           │   ├── SkipLink/
│           │   │   ├── SkipLink.tsx
│           │   │   ├── SkipLink.types.ts
│           │   │   └── SkipLink.web.tsx
│           │   └── VisuallyHidden/
│           │       ├── VisuallyHidden.native.tsx
│           │       ├── VisuallyHidden.tsx
│           │       ├── VisuallyHidden.types.ts
│           │       └── VisuallyHidden.web.tsx
│           └── components/
│               ├── Calendar/
│               │   ├── Calendar.native.tsx
│               │   ├── Calendar.tsx  # Full calendar
│               │   ├── Calendar.types.ts
│               │   ├── Calendar.web.tsx
│               │   ├── DatePicker/
│               │   │   ├── DatePicker.native.tsx
│               │   │   ├── DatePicker.tsx
│               │   │   └── DatePicker.web.tsx
│               │   ├── DateRangePicker/
│               │   │   ├── DateRangePicker.native.tsx
│               │   │   ├── DateRangePicker.tsx
│               │   │   └── DateRangePicker.web.tsx
│               │   └── TimePicker/
│               │       ├── TimePicker.native.tsx
│               │       ├── TimePicker.tsx
│               │       └── TimePicker.web.tsx
│               ├── DataDisplay/
│               │   ├── Kanban/
│               │   │   ├── KanbanBoard.native.tsx  # Touch gestures
│               │   │   ├── KanbanBoard.tsx
│               │   │   ├── KanbanBoard.types.ts
│               │   │   └── KanbanBoard.web.tsx  # Drag & drop
│               │   ├── List/
│               │   │   ├── List.native.tsx
│               │   │   ├── List.tsx
│               │   │   ├── List.types.ts
│               │   │   └── List.web.tsx
│               │   ├── Table/
│               │   │   ├── DataTable/
│               │   │   │   ├── DataTable.tsx  # With sorting/filtering
│               │   │   │   └── DataTable.types.ts
│               │   │   ├── Table.tsx  # Base table
│               │   │   ├── Table.types.ts
│               │   │   ├── Table.web.tsx  # Full featured table
│               │   │   └── VirtualTable/
│               │   │       ├── VirtualTable.tsx  # Virtualized table
│               │   │       └── VirtualTable.types.ts
│               │   └── Timeline/
│               │       ├── Timeline.native.tsx
│               │       ├── Timeline.tsx
│               │       ├── Timeline.types.ts
│               │       └── Timeline.web.tsx
│               ├── Editor/
│               │   ├── CodeEditor/
│               │   │   ├── CodeEditor.native.tsx
│               │   │   ├── CodeEditor.tsx
│               │   │   └── CodeEditor.web.tsx  # Web (e.g., Monaco/CodeMirror)
│               │   ├── MarkdownEditor/
│               │   │   ├── MarkdownEditor.native.tsx
│               │   │   ├── MarkdownEditor.tsx
│               │   │   └── MarkdownEditor.web.tsx
│               │   └── RichTextEditor/
│               │       ├── RichTextEditor.native.tsx  # Native implementation
│               │       ├── RichTextEditor.tsx
│               │       └── RichTextEditor.web.tsx  # Web (e.g., TipTap/Slate)
│               ├── Feedback/
│               │   ├── Alert/
│               │   │   ├── Alert.native.tsx
│               │   │   ├── Alert.tsx
│               │   │   └── Alert.web.tsx
│               │   ├── EmptyState/
│               │   │   ├── EmptyState.native.tsx
│               │   │   ├── EmptyState.tsx
│               │   │   └── EmptyState.web.tsx
│               │   ├── Notification/
│               │   │   ├── Notification.native.tsx
│               │   │   ├── Notification.tsx
│               │   │   └── Notification.web.tsx  # Toast notifications
│               │   └── Progress/
│               │       ├── Progress.types.ts
│               │       ├── ProgressBar/
│               │       │   ├── ProgressBar.native.tsx
│               │       │   ├── ProgressBar.tsx
│               │       │   └── ProgressBar.web.tsx
│               │       └── ProgressCircle/
│               │           ├── ProgressCircle.native.tsx
│               │           ├── ProgressCircle.tsx
│               │           └── ProgressCircle.web.tsx
│               ├── FileUpload/
│               │   ├── FileUpload.native.tsx  # Document picker
│               │   ├── FileUpload.tsx  # Base file upload
│               │   ├── FileUpload.web.tsx  # Drag & drop support
│               │   ├── FileUpload.types.ts
│               │   └── ImageUpload/
│               │       ├── ImageCropper.tsx  # Image cropping
│               │       ├── ImageUpload.native.tsx  # Camera/gallery
│               │       ├── ImageUpload.tsx
│               │       └── ImageUpload.web.tsx
│               ├── Forms/
│               │   ├── Checkbox/
│               │   │   ├── Checkbox.native.tsx
│               │   │   ├── Checkbox.tsx
│               │   │   └── Checkbox.web.tsx
│               │   ├── FormField/
│               │   │   ├── FormError.tsx
│               │   │   ├── FormField.native.tsx
│               │   │   ├── FormField.tsx  # Form field wrapper
│               │   │   ├── FormField.types.ts
│               │   │   ├── FormField.web.tsx
│               │   │   ├── FormHelperText.tsx
│               │   │   └── FormLabel.tsx
│               │   ├── Radio/
│               │   │   ├── Radio.types.ts
│               │   │   ├── RadioGroup.native.tsx
│               │   │   ├── RadioGroup.tsx
│               │   │   └── RadioGroup.web.tsx
│               │   ├── Select/
│               │   │   ├── MultiSelect/
│               │   │   │   ├── MultiSelect.native.tsx
│               │   │   │   ├── MultiSelect.tsx
│               │   │   │   └── MultiSelect.web.tsx
│               │   │   ├── Select.native.tsx
│               │   │   ├── Select.tsx
│               │   │   ├── Select.types.ts
│               │   │   └── Select.web.tsx
│               │   ├── Slider/
│               │   │   ├── RangeSlider/
│               │   │   │   ├── RangeSlider.native.tsx
│               │   │   │   ├── RangeSlider.tsx
│               │   │   │   └── RangeSlider.web.tsx
│               │   │   ├── Slider.native.tsx
│               │   │   ├── Slider.tsx
│               │   │   └── Slider.web.tsx
│               │   └── Switch/
│               │       ├── Switch.native.tsx
│               │       ├── Switch.tsx
│               │       └── Switch.web.tsx
│               └── Navigation/
│                   ├── Breadcrumb/
│                   │   ├── Breadcrumb.tsx
│                   │   ├── Breadcrumb.types.ts
│                   │   └── Breadcrumb.web.tsx
│                   ├── Pagination/
│                   │   ├── Pagination.native.tsx
│                   │   ├── Pagination.tsx
│                   │   └── Pagination.web.tsx
│                   ├── Stepper/
│                   │   ├── Stepper.native.tsx
│                   │   ├── Stepper.tsx
│                   │   └── Stepper.web.tsx
│                   └── Tabs/
│                       ├── Tabs.native.tsx
│                       ├── Tabs.tsx
│                       └── Tabs.web.tsx
├── scripts/
│   ├── build/
│   │   ├── analyze-bundle.sh  # Bundle analysis
│   │   ├── build-all.sh  # Build all apps
│   │   ├── build-mobile.sh  # Build mobile
│   │   └── build-web.sh  # Build web
│   ├── ci/
│   │   ├── lint-all.sh  # Lint all code
│   │   ├── pre-commit.sh  # Pre-commit hook
│   │   ├── pre-push.sh  # Pre-push hook
│   │   └── verify-types.sh  # Type check
│   ├── deploy/
│   │   ├── deploy-mobile-prod.sh
│   │   ├── deploy-mobile-staging.sh
│   │   ├── deploy-web-prod.sh
│   │   └── deploy-web-staging.sh
│   ├── dev/
│   │   ├── reset-db.sh  # Reset local DB
│   │   ├── start-all.sh  # Start all services
│   │   ├── start-mobile.sh  # Start mobile only
│   │   └── start-web.sh  # Start web only
│   ├── maintenance/
│   │   ├── clean-all.sh  # Clean all artifacts
│   │   ├── generate-types.sh  # Generate API types
│   │   └── update-deps.sh  # Update dependencies
│   ├── test/
│   │   ├── test-all.sh  # Run all tests
│   │   ├── test-coverage.sh  # Coverage report
│   │   ├── test-e2e-mobile.sh  # E2E mobile tests
│   │   ├── test-e2e-web.sh  # E2E web tests
│   │   ├── test-integration.sh  # Integration tests
│   │   └── test-unit.sh  # Unit tests
│   └── utils/
│       ├── check-ports.sh  # Check port availability
│       ├── doctor.sh  # Environment health check
│       └── setup-env.sh  # Setup environment
├── tests/
│   ├── e2e/
│   │   ├── fixtures/
│   │   │   ├── jobs.json
│   │   │   ├── users.json
│   │   │   └── ...
│   │   ├── mobile/
│   │   │   ├── auth/
│   │   │   ├── client/
│   │   │   ├── freelancer/
│   │   │   └── offline/
│   │   └── web/
│   │       ├── admin/
│   │       │   ├── financial-ops.spec.ts
│   │       │   ├── kyc.spec.ts
│   │       │   └── moderation.spec.ts
│   │       ├── client/
│   │       │   ├── hiring.spec.ts
│   │       │   ├── job-posting.spec.ts
│   │       │   ├── payments.spec.ts
│   │       │   └── talent-search.spec.ts
│   │       └── freelancer/
│   │           ├── contracts.spec.ts
│   │           ├── earnings.spec.ts
│   │           ├── profile.spec.ts
│   │           └── proposals.spec.ts
│   ├── integration/
│   │   ├── api/
│   │   │   ├── jobs-api.test.ts
│   │   │   ├── users-api.test.ts
│   │   │   └── ...
│   │   └── features/
│   │       ├── auth.test.ts
│   │       ├── jobs.test.ts
│   │       ├── proposals.test.ts
│   │       └── ...
│   └── performance/
│       ├── lighthouse/
│       │   ├── dashboard.spec.ts
│       │   ├── home.spec.ts
│       │   └── job-detail.spec.ts
│       └── load/
│           ├── api-load.spec.ts
│           └── user-simulation.spec.ts
├── .env.development  # Local development
├── .env.example  # Template
├── .env.production  # Production environment
├── .env.staging  # Staging environment
└── config/
    ├── constants/
    │   ├── api-endpoints.ts  # API endpoint constants
    │   ├── app-config.ts  # App configuration
    │   ├── performance-budgets.ts  # Performance budgets
    │   └── third-party-keys.ts  # Third-party service keys
    └── feature-flags.ts  # Feature flags config
