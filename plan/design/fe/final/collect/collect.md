```
fe/
├── .github/
│   ├── actions/
│   │   ├── build-mobile/
│   │   │   └── action.yml
│   │   ├── build-web/
│   │   │   └── action.yml
│   │   ├── cache-dependencies/
│   │   │   └── action.yml
│   │   ├── cache-deps/
│   │   │   └── action.yml
│   │   ├── deploy/
│   │   │   └── action.yml
│   │   ├── deploy-preview/
│   │   │   └── action.yml
│   │   ├── setup-node/
│   │   │   └── action.yml
│   │   └── setup-pnpm/
│   │       └── action.yml
│   ├── workflows/
│   │   ├── accessibility.yml  # Accessibility checks
│   │   │   # - Run axe-core
│   │   │   # - Check WCAG compliance
│   │   ├── bundle-analysis.yml  # Bundle size checks
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
│   │   ├── ci.yml  # Main CI pipeline
│   │   ├── dependency-review.yml  # Dependency review
│   │   │   # - Check for vulnerabilities
│   │   │   # - License compliance
│   │   ├── dependency-update.yml  # Dependabot automation
│   │   ├── deploy-mobile-production.yml
│   │   ├── deploy-mobile-staging.yml
│   │   ├── deploy-web-production.yml
│   │   ├── deploy-web-staging.yml
│   │   ├── e2e-tests.yml  # E2E tests
│   │   │   # - Setup test environment
│   │   │   # - Run Playwright/Detox tests
│   │   │   # - Upload test results
│   │   ├── lighthouse.yml  # Performance checks
│   │   │   # Performance audits
│   │   │   # - Run Lighthouse CI
│   │   │   # - Compare against budgets
│   │   │   # - Comment on PR
│   │   ├── lint.yml  # Linting
│   │   ├── release.yml  # Release automation
│   │   │   # - Create changelog
│   │   │   # - Tag release
│   │   │   # - Create GitHub release
│   │   ├── security-scan.yml  # Security scanning
│   │   ├── security.yml  # Security scanning
│   │   │   # - Dependency audit
│   │   │   # - SAST scan
│   │   │   # - License check
│   │   ├── test.yml  # Test runner
│   │   └── type-check.yml  # TypeScript checks
│   └── CODEOWNERS  # Code ownership
├── apps/
│   ├── mobile/
│   │   └── app/
│   │       ├── (dashboard)/
│   │       │   ├── activity-feed/
│   │       │   │   └── index.tsx  # Activity feed
│   │       │   │       # - Recent activities
│   │       │   │       # - Notifications inline
│   │       │   │       # - Quick actions from feed
│   │       │   │       # BE: communications-be/notification, utility/activity
│   │       │   │       # GET /v1/activity/feed
│   │       │   ├── quick-actions/
│   │       │   │   └── index.tsx  # Quick actions screen
│   │       │   │       # - Quick message
│   │       │   │       # - Quick proposal
│   │       │   │       # - Quick time entry
│   │       │   │       # - Quick invoice
│   │       │   │       # BE: Multiple services
│   │       │   ├── today/
│   │       │   │   └── index.tsx  # Today view
│   │       │   │       # - Today's schedule
│   │       │   │       # - Pending tasks
│   │       │   │       # - Quick metrics
│   │       │   │       # BE: contracts-be/work_diary, communications-be/notification
│   │       │   │       # GET /v1/today/overview
│   │       │   └── widgets/
│   │       │       ├── earnings/
│   │       │       │   └── index.tsx  # Earnings widget
│   │       │       │       # BE: financial-be/wallet
│   │       │       ├── notifications/
│   │       │       │   └── index.tsx  # Notifications widget
│   │       │       │       # BE: communications-be/notification
│   │       │       └── time-tracker/
│   │       │           └── index.tsx  # Time tracker widget
│   │       │               # BE: contracts-be/work_diary
│   │       ├── (tabs)/
│   │       │   ├── more/
│   │       │   │   ├── _layout.tsx  # More menu stack
│   │       │   │   ├── about.tsx  # About app
│   │       │   │   │   # BE: none (static)
│   │       │   │   ├── account.tsx  # Account settings
│   │       │   │   │   # BE: users-be/account
│   │       │   │   ├── help.tsx  # Help center
│   │       │   │   │   # BE: CMS
│   │       │   │   └── index.tsx  # More menu home
│   │       │   ├── notifications/
│   │       │   │   ├── [notificationId].tsx  # Notification detail
│   │       │   │   │   # BE: communications-be/notification
│   │       │   │   ├── _layout.tsx  # Notifications stack
│   │       │   │   ├── index.tsx  # All notifications
│   │       │   │   │   # BE: communications-be/notification
│   │       │   │   └── settings.tsx  # Notification settings
│   │       │   │       # BE: communications-be/preferences
│   │       │   └── search/
│   │       │       ├── _layout.tsx  # Search stack navigator
│   │       │       ├── filters.tsx  # Advanced filters
│   │       │       │   # BE: search-be/facets
│   │       │       ├── index.tsx  # Search home
│   │       │       │   # BE: search-be/query
│   │       │       ├── results.tsx  # Search results
│   │       │       │   # BE: search-be/query
│   │       │       └── saved.tsx  # Saved searches
│   │       │           # BE: search-be/saved-search
│   │       ├── accessibility/
│   │       │   ├── screen-reader/
│   │       │   │   └── index.tsx  # Screen reader optimized view
│   │       │   └── voice-commands/
│   │       │       └── index.tsx  # Voice commands
│   │       │           # - Voice-to-text
│   │       │           # - Command shortcuts
│   │       ├── camera/
│   │       │   ├── document-scan/
│   │       │   │   └── index.tsx  # Document scanner
│   │       │   │       # - ID verification
│   │       │   │       # - Contract documents
│   │       │   │       # - Expense receipts
│   │       │   │       # BE: storage-be/asset, admin-be/business_verification
│   │       │   │       # POST /v1/storage/uploads (scan)
│   │       │   ├── photo-upload/
│   │       │   │   └── index.tsx  # Photo upload
│   │       │   │       # - Portfolio photos
│   │       │   │       # - Work progress photos
│   │       │   │       # BE: storage-be/asset
│   │       │   │       # POST /v1/storage/uploads
│   │       │   └── qr-scan/
│   │       │       └── index.tsx  # QR code scanner
│   │       │           # - Profile sharing
│   │       │           # - Event check-in
│   │       │           # BE: users-be/profile (QR data)
│   │       ├── mobile-settings/
│   │       │   ├── advanced/
│   │       │   │   ├── developer-mode.tsx  # Developer options
│   │       │   │   ├── diagnostics.tsx  # Diagnostics tools
│   │       │   │   └── logs.tsx  # Error logs
│   │       │   ├── app-settings/
│   │       │   │   ├── cache.tsx  # Clear cache
│   │       │   │   ├── data-usage.tsx  # Data usage settings
│   │       │   │   ├── index.tsx  # App settings overview
│   │       │   │   ├── language.tsx  # Language selection
│   │       │   │   ├── notifications.tsx  # Notification preferences
│   │       │   │   └── theme.tsx  # Theme selection
│   │       │   ├── offline/
│   │       │   │   ├── offline-mode.tsx  # Offline capabilities
│   │       │   │   ├── storage.tsx  # Storage management
│   │       │   │   └── sync-frequency.tsx  # Sync settings
│   │       │   └── security/
│   │       │       ├── auto-lock.tsx  # Auto-lock timer
│   │       │       ├── biometric.tsx  # Biometric settings
│   │       │       ├── index.tsx  # Security settings
│   │       │       └── pin.tsx  # PIN setup
│   │       ├── offline/
│   │       │   ├── queue/
│   │       │   │   └── index.tsx  # Offline queue
│   │       │   │       # - Pending uploads
│   │       │   │       # - Queued messages
│   │       │   │       # - Draft proposals
│   │       │   │       # BE: None (local storage)
│   │       │   ├── settings/
│   │       │   │   └── index.tsx  # Offline settings
│   │       │   │       # - Auto-sync preferences
│   │       │   │       # - Download for offline
│   │       │   │       # BE: None (local storage)
│   │       │   ├── sync/
│   │       │   │   └── index.tsx  # Sync status
│   │       │   │       # - Sync progress
│   │       │   │       # - Conflict resolution
│   │       │   │       # BE: Multiple services (sync endpoints)
│   │       │   ├── queue.tsx  # Offline actions queue
│   │       │   │   # - Pending uploads
│   │       │   │   # - Queued messages
│   │       │   │   # - Draft proposals
│   │       │   └── sync.tsx  # Sync status
│   │       │       # - Sync progress
│   │       │       # - Conflict resolution
│   │       ├── onboarding/
│   │       │   ├── biometric/
│   │       │   │   └── setup.tsx  # Biometric auth setup
│   │       │   │       # - Face ID
│   │       │   │       # - Touch ID
│   │       │   │       # BE: users-be/auth
│   │       │   │       # POST /v1/auth/biometric/enable
│   │       │   ├── intro/
│   │       │   │   └── index.tsx  # App intro screens
│   │       │   │       # - Swipeable intro
│   │       │   │       # - Feature highlights
│   │       │   ├── permissions/
│   │       │   │   ├── camera.tsx  # Camera permission
│   │       │   │   ├── index.tsx  # All permissions
│   │       │   │   ├── location.tsx  # Location permission
│   │       │   │   └── notifications.tsx  # Notification permission
│   │       │   └── tutorial/
│   │       │       ├── apply.tsx  # Tutorial: Apply
│   │       │       ├── browse-jobs.tsx  # Tutorial: Browse jobs
│   │       │       ├── index.tsx  # Interactive tutorial
│   │       │       └── time-tracking.tsx  # Tutorial: Time tracking
│   │       ├── quick-actions/
│   │       │   ├── quick-apply/
│   │       │   │   └── [jobId].tsx  # Quick apply to job
│   │       │   │       # - Pre-filled proposal
│   │       │   │       # - One-tap apply
│   │       │   │       # - Connect deduction
│   │       │   │       # BE: proposals-be/proposal
│   │       │   │       # POST /v1/proposals/quick-apply
│   │       │   ├── quick-message/
│   │       │   │   └── [userId].tsx  # Quick message to user
│   │       │   │       # - Template messages
│   │       │   │       # - Voice-to-text
│   │       │   │       # BE: communications-be/conversation
│   │       │   │       # POST /v1/conversations/{conversation_id}/messages/quick
│   │       │   └── quick-time-entry/
│   │       │       └── [contractId].tsx  # Quick time logging
│   │       │           # - Timer widget
│   │       │           # - Quick notes
│   │       │           # - One-tap submit
│   │       │           # BE: contracts-be/work_diary
│   │       │           # POST /v1/work-diary/quick-entry
│   │       ├── scanner/
│   │       │   ├── document.tsx  # Document scanner
│   │       │   │   # - Scan compliance docs
│   │       │   │   # - OCR processing
│   │       │   │   # BE: storage-be/asset, admin-be/business_verification
│   │       │   └── qr-code.tsx  # QR code scanner
│   │       │       # - Event check-in
│   │       │       # - Profile sharing
│   │       ├── settings/
│   │       │   ├── advanced/
│   │       │   │   ├── developer/
│   │       │   │   │   └── index.tsx  # Developer options
│   │       │   │   │       # - API endpoint override
│   │       │   │   │       # - Debug logging
│   │       │   │   │       # - Performance monitoring
│   │       │   │   └── experiments/
│   │       │   │       └── index.tsx  # Experimental features
│   │       │   │           # - Beta features toggle
│   │       │   │           # BE: utility/flags
│   │       │   │           # GET /v1/flags/user
│   │       │   ├── app-preferences/
│   │       │   │   ├── data-usage/
│   │       │   │   │   └── index.tsx  # Data usage settings
│   │       │   │   │       # - Download on WiFi only
│   │       │   │   │       # - Auto-play videos
│   │       │   │   │       # - Image quality
│   │       │   │   ├── language/
│   │       │   │   │   └── index.tsx  # Language preferences
│   │       │   │   │       # - App language
│   │       │   │   │       # - Content language
│   │       │   │   ├── notifications/
│   │       │   │   │   ├── channels/
│   │       │   │   │   │   └── index.tsx  # Notification channels
│   │       │   │   │   ├── do-not-disturb/
│   │       │   │   │   │   └── index.tsx  # DND settings
│   │       │   │   │   └── index.tsx  # Notification preferences
│   │       │   │   │       # BE: communications-be/preferences
│   │       │   │   │       # GET/PUT /v1/notifications/preferences
│   │       │   │   └── theme/
│   │       │   │       └── index.tsx  # Theme settings
│   │       │   │           # - Light/dark/auto
│   │       │   │           # - Accent color
│   │       │   ├── biometric/
│   │       │   │   └── index.tsx  # Biometric authentication
│   │       │   │       # - Face ID / Touch ID
│   │       │   │       # - Setup
│   │       │   │       # BE: users-be/auth (device trust)
│   │       │   │       # POST /v1/auth/device-trust
│   │       │   └── storage/
│   │       │       ├── cache/
│   │       │       │   └── index.tsx  # Cache management
│   │       │       │       # - Clear cache
│   │       │       │       # - Cache size
│   │       │       └── downloads/
│   │       │           └── index.tsx  # Downloaded files
│   │       │               # - Manage downloads
│   │       │               # - Clear downloads
│   │       └── widgets/
│   │           ├── quick-actions.tsx  # Quick actions widget
│   │           │   # - Quick message
│   │           │   # - Quick proposal
│   │           └── time-tracker.tsx  # Home screen time tracker widget
│   │               # BE: contracts-be/work_diary
│   └── web/
│       ├── e2e/
│       │   ├── fixtures/
│       │   │   ├── test-data.ts
│       │   │   ├── test-jobs.json
│       │   │   └── test-users.json
│       │   ├── helpers/
│       │   │   ├── assertions.ts
│       │   │   ├── auth.ts
│       │   │   └── navigation.ts
│       │   ├── tests/
│       │   │   ├── auth/
│       │   │   │   ├── login.spec.ts
│       │   │   │   ├── oauth.spec.ts
│       │   │   │   ├── password-reset.spec.ts
│       │   │   │   └── register.spec.ts
│       │   │   ├── contracts/
│       │   │   │   ├── create-contract.spec.ts
│       │   │   │   ├── dispute.spec.ts
│       │   │   │   ├── submit-deliverable.spec.ts
│       │   │   │   └── time-tracking.spec.ts
│       │   │   ├── jobs/
│       │   │   │   ├── apply-to-job.spec.ts
│       │   │   │   ├── manage-jobs.spec.ts
│       │   │   │   ├── post-job.spec.ts
│       │   │   │   └── search-jobs.spec.ts
│       │   │   ├── messaging/
│       │   │   │   ├── notifications.spec.ts
│       │   │   │   ├── real-time-chat.spec.ts
│       │   │   │   └── send-message.spec.ts
│       │   │   ├── payments/
│       │   │   │   ├── add-payment-method.spec.ts
│       │   │   │   ├── escrow-funding.spec.ts
│       │   │   │   ├── release-payment.spec.ts
│       │   │   │   └── withdraw.spec.ts
│       │   │   └── proposals/
│       │   │       ├── create-proposal.spec.ts
│       │   │       ├── edit-proposal.spec.ts
│       │   │       └── withdraw-proposal.spec.ts
│       │   └── playwright.config.ts
│       ├── lib/
│       │   ├── image-optimization/
│       │   │   ├── loader.ts  # Custom image loader
│       │   │   ├── placeholder.ts  # Blur placeholders
│       │   │   └── responsive.ts  # Responsive images
│       │   └── security/
│       │       ├── cors.ts  # CORS configuration
│       │       ├── csp.ts  # Content Security Policy
│       │       ├── headers.ts  # Security headers
│       │       └── rate-limiter.ts  # Rate limiting
│       ├── security/
│       │   ├── auth-guard.ts
│       │   ├── csrf-protection.ts
│       │   └── rate-limiter.ts
│       ├── src/
│       │   └── app/
│       │       └── [locale]/
│       │           ├── (admin)/
│       │           │   ├── compliance/
│       │           │   │   ├── aml-kyc/
│       │           │   │   │   ├── monitoring/
│       │           │   │   │   │   └── page.tsx  # AML monitoring dashboard
│       │           │   │   │   │       # - Suspicious activity
│       │           │   │   │   │       # - Transaction patterns
│       │           │   │   │   │       # - Risk scores
│       │           │   │   │   │       # BE: admin-be/kyc_case, financial-be/transaction
│       │           │   │   │   │       # GET /v1/kyc/monitoring/suspicious-activity
│       │           │   │   │   ├── reports/
│       │           │   │   │   │   ├── [reportId]/
│       │           │   │   │   │   │   └── page.tsx  # SAR (Suspicious Activity Report) detail
│       │           │   │   │   │   │       # BE: admin-be/kyc_case
│       │           │   │   │   │   │       # GET /v1/kyc/reports/{report_id}
│       │           │   │   │   │   └── page.tsx  # AML reports list
│       │           │   │   │   │       # - Filed reports
│       │           │   │   │   │       # - Pending reports
│       │           │   │   │   │       # BE: admin-be/kyc_case
│       │           │   │   │   │       # GET /v1/kyc/reports
│       │           │   │   │   └── risk-assessment/
│       │           │   │   │       └── page.tsx  # Risk assessment tools
│       │           │   │   │           # - User risk profiles
│       │           │   │   │           # - Country risk matrix
│       │           │   │   │           # - Enhanced due diligence
│       │           │   │   │           # BE: admin-be/kyc_case
│       │           │   │   │           # GET /v1/kyc/risk-assessment
│       │           │   │   ├── data-retention/
│       │           │   │   │   ├── audit/
│       │           │   │   │   │   └── page.tsx  # Retention audit log
│       │           │   │   │   │       # - Deletion history
│       │           │   │   │   │       # - Policy compliance
│       │           │   │   │   │       # BE: utility/audit
│       │           │   │   │   │       # GET /v1/audit/retention
│       │           │   │   │   ├── policies/
│       │           │   │   │   │   └── page.tsx  # Retention policies
│       │           │   │   │   │       # - Policy definitions
│       │           │   │   │   │       # - Data categories
│       │           │   │   │   │       # - Retention periods
│       │           │   │   │   │       # BE: admin-be/data_retention (if exists) or utility/config
│       │           │   │   │   │       # GET /v1/retention/policies
│       │           │   │   │   └── schedule/
│       │           │   │   │       └── page.tsx  # Deletion schedule
│       │           │   │   │           # - Upcoming deletions
│       │           │   │   │           # - Retention expirations
│       │           │   │   │           # BE: admin-be/data_retention
│       │           │   │   │           # GET /v1/retention/schedule
│       │           │   │   ├── document-verification/
│       │           │   │   │   ├── [documentId]/
│       │           │   │   │   │   └── page.tsx  # Document review interface
│       │           │   │   │   │       # - Document viewer
│       │           │   │   │   │       # - Verification checks
│       │           │   │   │   │       # - Approve/reject/request-more
│       │           │   │   │   │       # BE: admin-be/business_verification, storage-be/asset
│       │           │   │   │   │       # GET /v1/verification/documents/{document_id}
│       │           │   │   │   │       # PUT /v1/verification/documents/{document_id}/review
│       │           │   │   │   ├── automated-checks/
│       │           │   │   │   │   └── page.tsx  # Automated verification rules
│       │           │   │   │   │       # - OCR settings
│       │           │   │   │   │       # - Validation rules
│       │           │   │   │   │       # - ML model performance
│       │           │   │   │   │       # BE: admin-be/business_verification
│       │           │   │   │   │       # GET /v1/verification/automation-rules
│       │           │   │   │   └── queue/
│       │           │   │   │       └── page.tsx  # Document verification queue
│       │           │   │   │           # - Pending documents
│       │           │   │   │           # - Priority sorting
│       │           │   │   │           # - Auto-verification status
│       │           │   │   │           # BE: admin-be/business_verification, storage-be/asset
│       │           │   │   │           # GET /v1/verification/documents/queue
│       │           │   │   └── gdpr/
│       │           │   │       ├── consent-management/
│       │           │   │       │   └── page.tsx  # Consent logs & management
│       │           │   │       │       # - User consent history
│       │           │   │       │       # - Consent versions
│       │           │   │       │       # - Audit trail
│       │           │   │       │       # BE: users-be/consent, utility/audit
│       │           │   │       │       # GET /v1/users/consent-logs
│       │           │   │       ├── deletion-requests/
│       │           │   │       │   ├── [requestId]/
│       │           │   │       │   │   └── page.tsx  # Deletion request detail
│       │           │   │       │   │       # - Data preview
│       │           │   │       │   │       # - Retention check
│       │           │   │       │   │       # - Process deletion
│       │           │   │       │   │       # BE: admin-be/privacy, users-be/user
│       │           │   │       │   │       # GET /v1/privacy/deletion-requests/{request_id}
│       │           │   │       │   │       # POST /v1/privacy/deletion-requests/{request_id}/process
│       │           │   │       │   └── page.tsx  # Deletion requests queue
│       │           │   │       │       # BE: admin-be/privacy
│       │           │   │       │       # GET /v1/privacy/deletion-requests
│       │           │   │       ├── export-requests/
│       │           │   │       │   ├── [requestId]/
│       │           │   │       │   │   └── page.tsx  # Export request detail
│       │           │   │       │   │       # - Request review
│       │           │   │       │   │       # - Generate export
│       │           │   │       │   │       # - Approve/deny
│       │           │   │       │   │       # BE: admin-be/privacy, users-be/user
│       │           │   │       │   │       # GET /v1/privacy/export-requests/{request_id}
│       │           │   │       │   │       # POST /v1/privacy/export-requests/{request_id}/process
│       │           │   │       │   └── page.tsx  # Export requests queue
│       │           │   │       │       # BE: admin-be/privacy
│       │           │   │       │       # GET /v1/privacy/export-requests
│       │           │   │       └── reports/
│       │           │   │           └── page.tsx  # GDPR compliance reports
│       │           │   │               # - Processing activities
│       │           │   │               # - Data inventory
│       │           │   │               # - Breach reports
│       │           │   │               # BE: admin-be/privacy, utility/audit
│       │           │   │               # GET /v1/privacy/reports
│       │           │   ├── incidents/
│       │           │   │   ├── [incidentId]/
│       │           │   │   │   ├── edit/
│       │           │   │   │   │   └── page.tsx  # Edit incident details
│       │           │   │   │   │       # - Update status
│       │           │   │   │   │       # - Add postmortem
│       │           │   │   │   │       # - Affected services
│       │           │   │   │   │       # BE: utility/status
│       │           │   │   │   │       # PUT /v1/status/incidents/{incident_id}
│       │           │   │   │   ├── timeline/
│       │           │   │   │   │   └── page.tsx  # Incident timeline
│       │           │   │   │   │       # - Event log
│       │           │   │   │   │       # - Update history
│       │           │   │   │   │       # BE: utility/status
│       │           │   │   │   │       # GET /v1/status/incidents/{incident_id}/timeline
│       │           │   │   │   └── page.tsx  # Incident detail
│       │           │   │   │       # - Current status
│       │           │   │   │       # - Impact assessment
│       │           │   │   │       # - Resolution steps
│       │           │   │   │       # BE: utility/status
│       │           │   │   │       # GET /v1/status/incidents/{incident_id}
│       │           │   │   ├── create/
│       │           │   │   │   └── page.tsx  # Create new incident
│       │           │   │   │       # - Incident type selection
│       │           │   │   │       # - Severity level
│       │           │   │   │       # - Affected services
│       │           │   │   │       # BE: utility/status
│       │           │   │   │       # POST /v1/status/incidents
│       │           │   │   ├── history/
│       │           │   │   │   └── page.tsx  # Historical incidents
│       │           │   │   │       # - Past incidents archive
│       │           │   │   │       # - Postmortems
│       │           │   │   │       # - Lessons learned
│       │           │   │   │       # BE: utility/status
│       │           │   │   │       # GET /v1/status/incidents/history
│       │           │   │   └── page.tsx  # Active incidents dashboard
│       │           │   │       # - Current incidents
│       │           │   │       # - Quick actions
│       │           │   │       # - Status board
│       │           │   │       # BE: utility/status
│       │           │   │       # GET /v1/status/incidents?status=active
│       │           │   ├── maintenance/
│       │           │   │   ├── [maintenanceId]/
│       │           │   │   │   ├── edit/
│       │           │   │   │   │   └── page.tsx  # Edit maintenance window
│       │           │   │   │   │       # BE: utility/status
│       │           │   │   │   │       # PUT /v1/status/maintenance/{maintenance_id}
│       │           │   │   │   └── page.tsx  # Maintenance detail
│       │           │   │   │       # BE: utility/status
│       │           │   │   │       # GET /v1/status/maintenance/{maintenance_id}
│       │           │   │   ├── schedule/
│       │           │   │   │   └── page.tsx  # Schedule maintenance
│       │           │   │   │       # - Date/time selection
│       │           │   │   │       # - Affected services
│       │           │   │   │       # - Notification plan
│       │           │   │   │       # BE: utility/status
│       │           │   │   │       # POST /v1/status/maintenance
│       │           │   │   └── page.tsx  # Maintenance calendar
│       │           │   │       # - Upcoming maintenance
│       │           │   │       # - Impact windows
│       │           │   │       # BE: utility/status
│       │           │   │       # GET /v1/status/maintenance
│       │           │   ├── platform-config/
│       │           │   │   ├── integrations/
│       │           │   │   │   ├── [integrationId]/
│       │           │   │   │   │   ├── configure/
│       │           │   │   │   │   │   └── page.tsx  # Configure integration
│       │           │   │   │   │   │       # BE: admin-be/integrations (if exists)
│       │           │   │   │   │   │       # PUT /v1/integrations/{integration_id}/config
│       │           │   │   │   │   ├── logs/
│       │           │   │   │   │   │   └── page.tsx  # Integration logs
│       │           │   │   │   │   │       # BE: utility/audit
│       │           │   │   │   │   │       # GET /v1/integrations/{integration_id}/logs
│       │           │   │   │   │   └── page.tsx  # Integration detail
│       │           │   │   │   │       # BE: admin-be/integrations
│       │           │   │   │   │       # GET /v1/integrations/{integration_id}
│       │           │   │   │   └── page.tsx  # Integrations list
│       │           │   │   │       # - Payment providers
│       │           │   │   │       # - Email services
│       │           │   │   │       # - Storage providers
│       │           │   │   │       # - Auth providers
│       │           │   │   │       # BE: admin-be/integrations
│       │           │   │   │       # GET /v1/integrations
│       │           │   │   ├── limits/
│       │           │   │   │   └── page.tsx  # Platform limits configuration
│       │           │   │   │       # - Rate limits
│       │           │   │   │       # - Upload limits
│       │           │   │   │       # - API quotas
│       │           │   │   │       # - Subscription limits
│       │           │   │   │       # BE: utility/config, admin-be/platform_config (if exists)
│       │           │   │   │       # GET /v1/config/limits
│       │           │   │   │       # PUT /v1/config/limits
│       │           │   │   ├── localization/
│       │           │   │   │   ├── languages/
│       │           │   │   │   │   └── page.tsx  # Language management
│       │           │   │   │   │       # - Enabled languages
│       │           │   │   │   │       # - Default language
│       │           │   │   │   │       # - RTL settings
│       │           │   │   │   │       # BE: utility/i18n
│       │           │   │   │   │       # GET /v1/config/languages
│       │           │   │   │   │       # PUT /v1/config/languages
│       │           │   │   │   └── regions/
│       │           │   │   │       └── page.tsx  # Regional settings
│       │           │   │   │           # - Timezone defaults
│       │           │   │   │           # - Currency settings
│       │           │   │   │           # - Date/time formats
│       │           │   │   │           # BE: utility/config
│       │           │   │   │           # GET /v1/config/regions
│       │           │   │   │           # PUT /v1/config/regions
│       │           │   │   ├── notifications/
│       │           │   │   │   ├── settings/
│       │           │   │   │   │   └── page.tsx  # Notification settings
│       │           │   │   │   │       # - Default preferences
│       │           │   │   │   │       # - Delivery channels
│       │           │   │   │   │       # - Retry policies
│       │           │   │   │   │       # BE: communications-be/config
│       │           │   │   │   │       # GET /v1/notifications/config
│       │           │   │   │   │       # PUT /v1/notifications/config
│       │           │   │   │   └── templates/
│       │           │   │   │       ├── [templateId]/
│       │           │   │   │       │   ├── edit/
│       │           │   │   │       │   │   └── page.tsx  # Edit notification template
│       │           │   │   │       │   │       # BE: communications-be/template
│       │           │   │   │       │   │       # PUT /v1/notifications/templates/{template_id}
│       │           │   │   │       │   ├── preview/
│       │           │   │   │       │   │   └── page.tsx  # Preview template
│       │           │   │   │       │   │       # BE: communications-be/template
│       │           │   │   │       │   │       # POST /v1/notifications/templates/{template_id}/preview
│       │           │   │   │       │   └── page.tsx  # Template detail
│       │           │   │   │       │       # BE: communications-be/template
│       │           │   │   │       │       # GET /v1/notifications/templates/{template_id}
│       │           │   │   │       └── page.tsx  # Template library
│       │           │   │   │           # BE: communications-be/template
│       │           │   │   │           # GET /v1/notifications/templates
│       │           │   │   └── pricing/
│       │           │   │       └── page.tsx  # Pricing configuration
│       │           │   │           # - Commission rates
│       │           │   │           # - Subscription pricing
│       │           │   │           # - Regional pricing
│       │           │   │           # BE: financial-be/pricing_config (if exists)
│       │           │   │           # GET /v1/config/pricing
│       │           │   │           # PUT /v1/config/pricing
│       │           │   │           # Note: Requires change_approval
│       │           │   └── system-health/
│       │           │       ├── metrics/
│       │           │       │   └── page.tsx  # System metrics dashboard
│       │           │       │       # - CPU/Memory usage
│       │           │       │       # - Database performance
│       │           │       │       # - Queue depths
│       │           │       │       # BE: utility/metrics (or monitoring service)
│       │           │       │       # GET /v1/metrics/system
│       │           │       ├── services/
│       │           │       │   └── page.tsx  # Service health overview
│       │           │       │       # - All microservices status
│       │           │       │       # - Uptime metrics
│       │           │       │       # - Response times
│       │           │       │       # BE: utility/status
│       │           │       │       # GET /v1/status/services
│       │           │       └── page.tsx  # Health dashboard
│       │           │           # - Overall system health
│       │           │           # - Critical alerts
│       │           │           # - Performance trends
│       │           │           # BE: utility/status
│       │           │           # GET /v1/status/health
│       │           ├── (dashboard)/
│       │           │   ├── agency/
│       │           │   │   ├── overview/
│       │           │   │   │   └── page.tsx  # Agency dashboard
│       │           │   │   │       # - Revenue overview
│       │           │   │   │       # - Active projects
│       │           │   │   │       # - Team utilization
│       │           │   │   │       # - Client list
│       │           │   │   │       # BE: users-be/org, financial-be/analytics
│       │           │   │   │       # GET /v1/agencies/me/overview
│       │           │   │   ├── reporting/
│       │           │   │   │   ├── clients/
│       │           │   │   │   │   └── page.tsx  # Client reports
│       │           │   │   │   │       # - Per-client spending
│       │           │   │   │   │       # - Project status
│       │           │   │   │   │       # BE: financial-be/analytics
│       │           │   │   │   │       # GET /v1/agencies/reports/clients
│       │           │   │   │   ├── financial/
│       │           │   │   │   │   └── page.tsx  # Financial reports
│       │           │   │   │   │       # - Revenue by period
│       │           │   │   │   │       # - Margins
│       │           │   │   │   │       # - Forecasts
│       │           │   │   │   │       # BE: financial-be/analytics
│       │           │   │   │   │       # GET /v1/agencies/reports/financial
│       │           │   │   │   └── team/
│       │           │   │   │       └── page.tsx  # Team performance
│       │           │   │   │           # - Utilization rates
│       │           │   │   │           # - Project assignments
│       │           │   │   │           # BE: users-be/org, jobs-be/analytics
│       │           │   │   │           # GET /v1/agencies/reports/team
│       │           │   │   ├── sub-accounts/
│       │           │   │   │   ├── [subAccountId]/
│       │           │   │   │   │   ├── settings/
│       │           │   │   │   │   │   └── page.tsx  # Sub-account settings
│       │           │   │   │   │   │       # BE: users-be/org
│       │           │   │   │   │   │       # PUT /v1/agencies/sub-accounts/{sub_account_id}
│       │           │   │   │   │   └── page.tsx  # Sub-account detail
│       │           │   │   │   │       # - Jobs posted
│       │           │   │   │   │       # - Contracts
│       │           │   │   │   │       # - Spending
│       │           │   │   │   │       # BE: users-be/org
│       │           │   │   │   │       # GET /v1/agencies/sub-accounts/{sub_account_id}
│       │           │   │   │   ├── create/
│       │           │   │   │   │   └── page.tsx  # Create sub-account
│       │           │   │   │   │       # BE: users-be/org
│       │           │   │   │   │       # POST /v1/agencies/sub-accounts
│       │           │   │   │   └── page.tsx  # All sub-accounts
│       │           │   │   │       # - List all
│       │           │   │   │       # - Manage
│       │           │   │   │       # BE: users-be/org
│       │           │   │   │       # GET /v1/agencies/sub-accounts
│       │           │   │   ├── talent-pool/
│       │           │   │   │   ├── [poolId]/
│       │           │   │   │   │   ├── members/
│       │           │   │   │   │   │   └── page.tsx  # Pool members
│       │           │   │   │   │   │       # - Freelancer list
│       │           │   │   │   │   │       # - Add/remove
│       │           │   │   │   │   │       # BE: users-be/org, search-be/talent-pool
│       │           │   │   │   │   │       # GET /v1/talent-pools/{pool_id}/members
│       │           │   │   │   │   └── page.tsx  # Talent pool detail
│       │           │   │   │   │       # BE: search-be/talent-pool
│       │           │   │   │   │       # GET /v1/talent-pools/{pool_id}
│       │           │   │   │   ├── create/
│       │           │   │   │   │   └── page.tsx  # Create talent pool
│       │           │   │   │   │       # BE: search-be/talent-pool
│       │           │   │   │   │       # POST /v1/talent-pools
│       │           │   │   │   └── page.tsx  # All talent pools
│       │           │   │   │       # BE: search-be/talent-pool
│       │           │   │   │       # GET /v1/talent-pools
│       │           │   │   └── white-label/
│       │           │   │       └── page.tsx  # White-label settings
│       │           │   │           # - Branding
│       │           │   │           # - Custom domain
│       │           │   │           # - Logo/colors
│       │           │   │           # BE: users-be/org
│       │           │   │           # PUT /v1/agencies/white-label
│       │           │   ├── analytics/
│       │           │   │   ├── custom-reports/
│       │           │   │   │   ├── [reportId]/
│       │           │   │   │   │   ├── edit/
│       │           │   │   │   │   │   └── page.tsx  # Edit custom report
│       │           │   │   │   │   │       # BE: financial-be/reports
│       │           │   │   │   │   │       # PUT /v1/reports/custom/{report_id}
│       │           │   │   │   │   └── page.tsx  # View custom report
│       │           │   │   │   │       # BE: financial-be/reports
│       │           │   │   │   │       # GET /v1/reports/custom/{report_id}
│       │           │   │   │   ├── new/
│       │           │   │   │   │   └── page.tsx  # Create custom report
│       │           │   │   │   │       # BE: financial-be/reports
│       │           │   │   │   │       # POST /v1/reports/custom
│       │           │   │   │   └── page.tsx  # Custom reports list
│       │           │   │   │       # BE: financial-be/reports
│       │           │   │   │       # GET /v1/reports/custom
│       │           │   │   ├── earnings/
│       │           │   │   │   ├── forecast/
│       │           │   │   │   │   └── page.tsx  # Earnings forecast
│       │           │   │   │   │       # - Projected income
│       │           │   │   │   │       # - Pipeline value
│       │           │   │   │   │       # BE: financial-be/analytics
│       │           │   │   │   │       # GET /v1/analytics/earnings/forecast
│       │           │   │   │   └── page.tsx  # Earnings analytics
│       │           │   │   │       # - Monthly trends
│       │           │   │   │       # - Year-over-year comparison
│       │           │   │   │       # - Client breakdown
│       │           │   │   │       # BE: financial-be/analytics
│       │           │   │   │       # GET /v1/analytics/earnings
│       │           │   │   ├── market-insights/
│       │           │   │   │   └── page.tsx  # Market insights
│       │           │   │   │       # - Skill demand trends
│       │           │   │   │       # - Rate benchmarks
│       │           │   │   │       # - Competition analysis
│       │           │   │   │       # BE: search-be/analytics, jobs-be/analytics
│       │           │   │   │       # GET /v1/analytics/market-insights
│       │           │   │   └── performance/
│       │           │   │       └── page.tsx  # Performance analytics
│       │           │   │           # - Response time metrics
│       │           │   │           # - Proposal success rate
│       │           │   │           # - Client satisfaction
│       │           │   │           # BE: users-be/analytics, proposals-be/performance
│       │           │   │           # GET /v1/users/me/analytics/performance
│       │           │   ├── availability/
│       │           │   │   ├── calendar/
│       │           │   │   │   └── page.tsx  # Availability calendar
│       │           │   │   │       # - Mark available/busy
│       │           │   │   │       # - Recurring patterns
│       │           │   │   │       # - Sync with external calendar
│       │           │   │   │       # BE: users-be/availability (if exists) or users-be/profile
│       │           │   │   │       # GET /v1/users/me/availability
│       │           │   │   │       # PUT /v1/users/me/availability
│       │           │   │   ├── settings/
│       │           │   │   │   └── page.tsx  # Availability settings
│       │           │   │   │       # - Working hours
│       │           │   │   │       # - Timezone preferences
│       │           │   │   │       # - Auto-reply settings
│       │           │   │   │       # BE: users-be/settings
│       │           │   │   │       # PUT /v1/users/me/availability-settings
│       │           │   │   └── page.tsx  # Availability dashboard
│       │           │   │       # - Current status
│       │           │   │       # - Upcoming commitments
│       │           │   │       # BE: users-be/availability
│       │           │   │       # GET /v1/users/me/availability-overview
│       │           │   ├── bidding/
│       │           │   │   ├── analytics/
│       │           │   │   │   └── page.tsx  # Bidding analytics
│       │           │   │   │       # - Win rate
│       │           │   │   │       # - Average bid amount
│       │           │   │   │       # - Competition analysis
│       │           │   │   │       # BE: proposals-be/bid-strategy
│       │           │   │   │       # GET /v1/bid-strategies/analytics
│       │           │   │   ├── auctions/
│       │           │   │   │   ├── [auctionId]/
│       │           │   │   │   │   └── page.tsx  # Auction participation
│       │           │   │   │   │       # - Real-time bidding
│       │           │   │   │   │       # - Bid history
│       │           │   │   │   │       # - Competitor activity
│       │           │   │   │   │       # BE: proposals-be/auction
│       │           │   │   │   │       # GET /v1/jobs/{job_id}/auction
│       │           │   │   │   │       # POST /v1/jobs/{job_id}/auction/bid
│       │           │   │   │   │       # WebSocket: Real-time updates
│       │           │   │   │   └── page.tsx  # Active auctions list
│       │           │   │   │       # BE: proposals-be/auction
│       │           │   │   │       # GET /v1/auctions/active
│       │           │   │   └── strategies/
│       │           │   │       ├── [strategyId]/
│       │           │   │       │   ├── edit/
│       │           │   │       │   │   └── page.tsx  # Edit bid strategy
│       │           │   │       │   │       # BE: proposals-be/bid-strategy
│       │           │   │       │   │       # PUT /v1/bid-strategies/{strategy_id}
│       │           │   │       │   └── page.tsx  # View bid strategy details
│       │           │   │       │       # BE: proposals-be/bid-strategy
│       │           │   │       │       # GET /v1/bid-strategies/{strategy_id}
│       │           │   │       ├── new/
│       │           │   │       │   └── page.tsx  # Create new bid strategy
│       │           │   │       │       # BE: proposals-be/bid-strategy
│       │           │   │       │       # POST /v1/bid-strategies
│       │           │   │       └── page.tsx  # Bid strategies list
│       │           │   │           # - Auto-bid rules
│       │           │   │           # - Price ranges
│       │           │   │           # - Category targeting
│       │           │   │           # BE: proposals-be/bid-strategy
│       │           │   │           # GET /v1/bid-strategies
│       │           │   ├── community/
│       │           │   │   ├── events/
│       │           │   │   │   ├── [eventId]/
│       │           │   │   │   │   ├── register/
│       │           │   │   │   │   │   └── page.tsx  # Event registration
│       │           │   │   │   │   │       # - RSVP form
│       │           │   │   │   │   │       # - Add to calendar
│       │           │   │   │   │   │       # BE: communications-be/events
│       │           │   │   │   │   │       # POST /v1/events/{event_id}/register
│       │           │   │   │   │   └── page.tsx  # Event detail
│       │           │   │   │   │       # - Info
│       │           │   │   │   │       # - Attendees
│       │           │   │   │   │       # - Check-in (QR code)
│       │           │   │   │   │       # BE: communications-be/events
│       │           │   │   │   │       # GET /v1/events/{event_id}
│       │           │   │   │   ├── my-events/
│       │           │   │   │   │   └── page.tsx  # My registered events
│       │           │   │   │   │       # - Attending
│       │           │   │   │   │       # - Past attended
│       │           │   │   │   │       # BE: communications-be/events
│       │           │   │   │   │       # GET /v1/users/me/events
│       │           │   │   │   └── upcoming/
│       │           │   │   │       └── page.tsx  # Upcoming events
│       │           │   │   │           # BE: communications-be/events
│       │           │   │   │           # GET /v1/events?status=upcoming
│       │           │   │   ├── forums/
│       │           │   │   │   ├── [forumId]/
│       │           │   │   │   │   ├── threads/
│       │           │   │   │   │   │   └── [threadId]/
│       │           │   │   │   │   │       ├── reply/
│       │           │   │   │   │   │       │   └── page.tsx  # Reply to thread
│       │           │   │   │   │   │       │       # BE: communications-be/forums
│       │           │   │   │   │   │       │       # POST /v1/threads/{thread_id}/replies
│       │           │   │   │   │   │       └── page.tsx  # Thread view
│       │           │   │   │   │   │           # - Posts
│       │           │   │   │   │   │           # - Reply
│       │           │   │   │   │   │           # - Voting
│       │           │   │   │   │   │           # BE: communications-be/forums (if exists) OR community service
│       │           │   │   │   │   │           # GET /v1/forums/{forum_id}/threads/{thread_id}
│       │           │   │   │   │   └── page.tsx  # Forum overview
│       │           │   │   │   │       # - Thread list
│       │           │   │   │   │       # - Create thread
│       │           │   │   │   │       # BE: communications-be/forums
│       │           │   │   │   │       # GET /v1/forums/{forum_id}/threads
│       │           │   │   │   └── page.tsx  # All forums
│       │           │   │   │       # - Categories
│       │           │   │   │       # - Popular threads
│       │           │   │   │       # BE: communications-be/forums
│       │           │   │   │       # GET /v1/forums
│       │           │   │   └── groups/
│       │           │   │       ├── [groupId]/
│       │           │   │       │   ├── discussions/
│       │           │   │       │   │   └── page.tsx  # Group discussions
│       │           │   │       │   │       # - Posts feed
│       │           │   │       │   │       # - Comment
│       │           │   │       │   │       # BE: communications-be/discussions
│       │           │   │       │   │       # GET /v1/groups/{group_id}/discussions
│       │           │   │       │   ├── events/
│       │           │   │       │   │   ├── [eventId]/
│       │           │   │       │   │   │   └── page.tsx  # Event detail
│       │           │   │       │   │   │       # - RSVP
│       │           │   │       │   │   │       # - Calendar add
│       │           │   │       │   │   │       # BE: communications-be/events OR community service
│       │           │   │       │   │   │       # GET /v1/events/{event_id}
│       │           │   │       │   │   └── page.tsx  # Group events
│       │           │   │       │   │       # - Upcoming
│       │           │   │       │   │       # - Past
│       │           │   │       │   │       # BE: communications-be/events
│       │           │   │       │   │       # GET /v1/groups/{group_id}/events
│       │           │   │       │   ├── members/
│       │           │   │       │   │   └── page.tsx  # Group members
│       │           │   │       │   │       # - Member list
│       │           │   │       │   │       # - Invite
│       │           │   │       │   │       # - Roles
│       │           │   │       │   │       # BE: users-be/groups (if exists) OR community service
│       │           │   │       │   │       # GET /v1/groups/{group_id}/members
│       │           │   │       │   └── page.tsx  # Group overview
│       │           │   │       │       # - About
│       │           │   │       │       # - Join/leave
│       │           │   │       │       # - Activity feed
│       │           │   │       │       # BE: users-be/groups
│       │           │   │       │       # GET /v1/groups/{group_id}
│       │           │   │       ├── discover/
│       │           │   │       │   └── page.tsx  # Discover groups
│       │           │   │       │       # - Recommended
│       │           │   │       │       # - By interest
│       │           │   │       │       # BE: users-be/groups
│       │           │   │       │       # GET /v1/groups/discover
│       │           │   │       └── my-groups/
│       │           │   │           └── page.tsx  # My groups
│       │           │   │               # - Joined groups
│       │           │   │               # - Manage groups
│       │           │   │               # BE: users-be/groups
│       │           │   │               # GET /v1/users/me/groups
│       │           │   ├── compliance/
│       │           │   │   ├── documents/
│       │           │   │   │   ├── [documentId]/
│       │           │   │   │   │   └── page.tsx  # Compliance document details
│       │           │   │   │   │       # BE: storage-be/asset, admin-be/business_verification
│       │           │   │   │   │       # GET /v1/compliance/documents/{document_id}
│       │           │   │   │   ├── upload/
│       │           │   │   │   │   └── page.tsx  # Upload compliance documents
│       │           │   │   │   │       # BE: storage-be/asset, admin-be/business_verification
│       │           │   │   │   │       # POST /v1/compliance/documents/upload
│       │           │   │   │   └── page.tsx  # Compliance documents list
│       │           │   │   │       # BE: admin-be/business_verification
│       │           │   │   │       # GET /v1/compliance/documents
│       │           │   │   ├── reports/
│       │           │   │   │   ├── tax-summary/
│       │           │   │   │   │   └── page.tsx  # Annual tax summary
│       │           │   │   │   │       # BE: financial-be/tax
│       │           │   │   │   │       # GET /v1/tax/reports/annual-summary
│       │           │   │   │   └── page.tsx  # Compliance reports
│       │           │   │   │       # - Income reports
│       │           │   │   │       # - Tax withholding
│       │           │   │   │       # - Payment history
│       │           │   │   │       # BE: financial-be/reports
│       │           │   │   │       # GET /v1/reports/compliance
│       │           │   │   └── tax-profile/
│       │           │   │       ├── edit/
│       │           │   │       │   └── page.tsx  # Edit tax profile
│       │           │   │       │       # BE: users-be/compliance, financial-be/tax
│       │           │   │       │       # PUT /v1/users/me/compliance/tax-profile
│       │           │   │       └── page.tsx  # Tax profile overview
│       │           │   │           # - Tax ID
│       │           │   │           # - Tax forms
│       │           │   │           # - Withholding settings
│       │           │   │           # BE: users-be/compliance
│       │           │   │           # GET /v1/users/me/compliance/tax-profile
│       │           │   ├── connects/
│       │           │   │   ├── purchase/
│       │           │   │   │   └── page.tsx  # Purchase connects
│       │           │   │   │       # - Select package
│       │           │   │   │       # - Payment processing
│       │           │   │   │       # BE: proposals-be/connect, financial-be/payment
│       │           │   │   │       # GET /v1/connects/packages
│       │           │   │   │       # POST /v1/connects/purchase
│       │           │   │   ├── usage/
│       │           │   │   │   └── page.tsx  # Connects usage analytics
│       │           │   │   │       # - Spending patterns
│       │           │   │   │       # - Refund history
│       │           │   │   │       # - ROI tracking
│       │           │   │   │       # BE: proposals-be/connect
│       │           │   │   │       # GET /v1/connects/usage-analytics
│       │           │   │   └── page.tsx  # Connects dashboard
│       │           │   │       # - Current balance
│       │           │   │       # - Transaction history
│       │           │   │       # - Refund requests
│       │           │   │       # BE: proposals-be/connect
│       │           │   │       # GET /v1/connects
│       │           │   │       # GET /v1/connects/balance
│       │           │   ├── deliverables/
│       │           │   │   ├── [contractId]/
│       │           │   │   │   ├── [deliverableId]/
│       │           │   │   │   │   ├── review/
│       │           │   │   │   │   │   └── page.tsx  # Review deliverable (client)
│       │           │   │   │   │   │       # - Approve/reject
│       │           │   │   │   │   │       # - Request changes
│       │           │   │   │   │   │       # - Add comments
│       │           │   │   │   │   │       # BE: contracts-be/deliverable
│       │           │   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/review
│       │           │   │   │   │   ├── revisions/
│       │           │   │   │   │   │   ├── [revisionId]/
│       │           │   │   │   │   │   │   └── page.tsx  # Revision detail
│       │           │   │   │   │   │   │       # BE: contracts-be/deliverable
│       │           │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/revisions/{revision_id}
│       │           │   │   │   │   │   └── page.tsx  # Revision history
│       │           │   │   │   │   │       # BE: contracts-be/deliverable
│       │           │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/revisions
│       │           │   │   │   │   ├── upload/
│       │           │   │   │   │   │   └── page.tsx  # Upload new version
│       │           │   │   │   │   │       # BE: contracts-be/deliverable, storage-be/asset
│       │           │   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/upload
│       │           │   │   │   │   └── page.tsx  # Deliverable details
│       │           │   │   │   │       # - File viewer
│       │           │   │   │   │       # - Download
│       │           │   │   │   │       # - Metadata
│       │           │   │   │   │       # - Comments thread
│       │           │   │   │   │       # BE: contracts-be/deliverable, storage-be/asset
│       │           │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}
│       │           │   │   │   ├── new/
│       │           │   │   │   │   └── page.tsx  # Submit new deliverable
│       │           │   │   │   │       # BE: contracts-be/deliverable, storage-be/asset
│       │           │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables
│       │           │   │   │   └── page.tsx  # Contract deliverables list
│       │           │   │   │       # BE: contracts-be/deliverable
│       │           │   │   │       # GET /v1/contracts/{contract_id}/deliverables
│       │           │   │   ├── pending-review/
│       │           │   │   │   └── page.tsx  # Deliverables pending client review
│       │           │   │   │       # BE: contracts-be/deliverable
│       │           │   │   │       # GET /v1/deliverables/pending-review
│       │           │   │   └── page.tsx  # All deliverables overview
│       │           │   │       # BE: contracts-be/deliverable
│       │           │   │       # GET /v1/deliverables
│       │           │   ├── invitations/
│       │           │   │   ├── received/
│       │           │   │   │   ├── [inviteId]/
│       │           │   │   │   │   └── page.tsx  # Invitation details
│       │           │   │   │   │       # - Job details
│       │           │   │   │   │       # - Accept/decline
│       │           │   │   │   │       # - Proposal draft
│       │           │   │   │   │       # BE: proposals-be/invite, jobs-be/job
│       │           │   │   │   │       # GET /v1/invites/{invite_id}
│       │           │   │   │   │       # POST /v1/invites/{invite_id}/accept
│       │           │   │   │   │       # POST /v1/invites/{invite_id}/decline
│       │           │   │   │   └── page.tsx  # Received invitations list
│       │           │   │   │       # BE: proposals-be/invite
│       │           │   │   │       # GET /v1/invites/received
│       │           │   │   ├── sent/
│       │           │   │   │   ├── [inviteId]/
│       │           │   │   │   │   └── page.tsx  # Sent invitation tracking
│       │           │   │   │   │       # - Delivery status
│       │           │   │   │   │       # - Response tracking
│       │           │   │   │   │       # BE: jobs-be/invitation
│       │           │   │   │   │       # GET /v1/jobs/{job_id}/invitations/{invite_id}
│       │           │   │   │   └── page.tsx  # Sent invitations list (client)
│       │           │   │   │       # BE: jobs-be/invitation
│       │           │   │   │       # GET /v1/jobs/{job_id}/invitations
│       │           │   │   └── page.tsx  # Invitations overview
│       │           │   │       # - Pending actions
│       │           │   │       # - Response rate (client)
│       │           │   │       # - Conversion metrics
│       │           │   │       # BE: proposals-be/invite OR jobs-be/invitation (based on role)
│       │           │   ├── job-alerts/
│       │           │   │   ├── [alertId]/
│       │           │   │   │   ├── edit/
│       │           │   │   │   │   └── page.tsx  # Edit job alert
│       │           │   │   │   │       # BE: search-be/alert (if exists) or search-be/saved-search
│       │           │   │   │   │       # PUT /v1/job-alerts/{alert_id}
│       │           │   │   │   ├── history/
│       │           │   │   │   │   └── page.tsx  # Alert history
│       │           │   │   │   │       # - Jobs matched
│       │           │   │   │   │       # - Notifications sent
│       │           │   │   │   │       # BE: search-be/alert
│       │           │   │   │   │       # GET /v1/job-alerts/{alert_id}/history
│       │           │   │   │   └── page.tsx  # Alert detail
│       │           │   │   │       # BE: search-be/alert
│       │           │   │   │       # GET /v1/job-alerts/{alert_id}
│       │           │   │   ├── create/
│       │           │   │   │   └── page.tsx  # Create job alert
│       │           │   │   │       # - Search criteria
│       │           │   │   │       # - Notification frequency
│       │           │   │   │       # - Delivery channel
│       │           │   │   │       # BE: search-be/alert
│       │           │   │   │       # POST /v1/job-alerts
│       │           │   │   └── page.tsx  # Job alerts list
│       │           │   │       # - Active alerts
│       │           │   │       # - Pause/resume
│       │           │   │       # BE: search-be/alert
│       │           │   │       # GET /v1/job-alerts
│       │           │   ├── learning/
│       │           │   │   ├── achievements/
│       │           │   │   │   └── page.tsx  # Achievements & badges
│       │           │   │   │       # - Earned badges
│       │           │   │   │       # - Progress to next level
│       │           │   │   │       # - Leaderboard
│       │           │   │   │       # BE: users-be/achievement
│       │           │   │   │       # GET /v1/users/me/achievements
│       │           │   │   │       # Learning achievements
│       │           │   │   │       # - Certificates earned
│       │           │   │   │       # - Badges
│       │           │   │   │       # - Skill verifications
│       │           │   │   │       # BE: learning-be/achievement
│       │           │   │   │       # GET /v1/learning/achievements
│       │           │   │   ├── assessments/
│       │           │   │   │   ├── [assessmentId]/
│       │           │   │   │   │   ├── results/
│       │           │   │   │   │   │   └── page.tsx  # Assessment results
│       │           │   │   │   │   │       # BE: learning-be/assessment
│       │           │   │   │   │   │       # GET /v1/learning/assessments/{assessment_id}/results
│       │           │   │   │   │   ├── take/
│       │           │   │   │   │   │   └── page.tsx  # Take assessment
│       │           │   │   │   │   │       # BE: learning-be/assessment
│       │           │   │   │   │   │       # POST /v1/learning/assessments/{assessment_id}/submit
│       │           │   │   │   │   └── page.tsx  # Assessment detail
│       │           │   │   │   │       # BE: learning-be/assessment
│       │           │   │   │   │       # GET /v1/learning/assessments/{assessment_id}
│       │           │   │   │   └── page.tsx  # Assessments list
│       │           │   │   │       # BE: learning-be/assessment
│       │           │   │   │       # GET /v1/learning/assessments
│       │           │   │   ├── certifications/
│       │           │   │   │   ├── [certId]/
│       │           │   │   │   │   ├── verify/
│       │           │   │   │   │   │   └── page.tsx  # Verify certificate
│       │           │   │   │   │   │       # BE: utility-be/learning
│       │           │   │   │   │   │       # GET /v1/certifications/{cert_id}/verify
│       │           │   │   │   │   └── page.tsx  # Certificate detail
│       │           │   │   │   │       # - Download PDF
│       │           │   │   │   │       # - Share link
│       │           │   │   │   │       # - Add to profile
│       │           │   │   │   │       # BE: utility-be/learning
│       │           │   │   │   │       # GET /v1/certifications/{cert_id}
│       │           │   │   │   └── page.tsx  # Manage certifications
│       │           │   │   │       # - Upload certificates
│       │           │   │   │       # - Verification status
│       │           │   │   │       # - Expiry tracking
│       │           │   │   │       # BE: users-be/credential
│       │           │   │   │       # GET /v1/users/me/credentials?type=certification
│       │           │   │   │       # All certifications
│       │           │   │   │       # - Earned certificates
│       │           │   │   │       # - Available certifications
│       │           │   │   │       # BE: utility-be/learning
│       │           │   │   │       # GET /v1/users/me/certifications
│       │           │   │   ├── courses/
│       │           │   │   │   ├── [courseId]/
│       │           │   │   │   │   ├── assessments/
│       │           │   │   │   │   │   └── [assessmentId]/
│       │           │   │   │   │   │       ├── results/
│       │           │   │   │   │   │       │   └── page.tsx  # View results
│       │           │   │   │   │   │       │       # BE: utility-be/learning
│       │           │   │   │   │   │       │       # GET /v1/assessments/{assessment_id}/results
│       │           │   │   │   │   │       ├── start/
│       │           │   │   │   │   │       │   └── page.tsx  # Start assessment
│       │           │   │   │   │   │       │       # BE: utility-be/learning
│       │           │   │   │   │   │       │       # POST /v1/assessments/{assessment_id}/start
│       │           │   │   │   │   │       └── submit/
│       │           │   │   │   │   │           └── page.tsx  # Submit answers
│       │           │   │   │   │   │               # BE: utility-be/learning
│       │           │   │   │   │   │               # POST /v1/assessments/{assessment_id}/submit
│       │           │   │   │   │   ├── lessons/
│       │           │   │   │   │   │   └── [lessonId]/
│       │           │   │   │   │   │       └── page.tsx  # Lesson view
│       │           │   │   │   │   │           # - Video/content player
│       │           │   │   │   │   │           # - Notes taking
│       │           │   │   │   │   │           # - Progress tracking
│       │           │   │   │   │   │           # BE: utility-be/learning OR external LMS
│       │           │   │   │   │   │           # GET /v1/courses/{course_id}/lessons/{lesson_id}
│       │           │   │   │   │   │           # Lesson content
│       │           │   │   │   │   │           # BE: learning-be/lesson (if exists) or external LMS
│       │           │   │   │   │   │           # GET /v1/learning/courses/{course_id}/lessons/{lesson_id}
│       │           │   │   │   │   ├── progress/
│       │           │   │   │   │   │   └── page.tsx  # Course progress
│       │           │   │   │   │   │       # BE: learning-be/progress
│       │           │   │   │   │   │       # GET /v1/learning/courses/{course_id}/progress
│       │           │   │   │   │   └── page.tsx  # Course overview
│       │           │   │   │   │       # - Syllabus
│       │           │   │   │   │       # - Progress
│       │           │   │   │   │       # - Enroll button
│       │           │   │   │   │       # BE: utility-be/learning
│       │           │   │   │   │       # GET /v1/courses/{course_id}
│       │           │   │   │   │       # POST /v1/courses/{course_id}/enroll
│       │           │   │   │   │       # Course detail
│       │           │   │   │   │       # BE: learning-be/course
│       │           │   │   │   │       # GET /v1/learning/courses/{course_id}
│       │           │   │   │   ├── browse/
│       │           │   │   │   │   └── page.tsx  # Browse all courses
│       │           │   │   │   │       # - Filter by skill
│       │           │   │   │   │       # - Search
│       │           │   │   │   │       # - Recommendations
│       │           │   │   │   │       # BE: utility-be/learning
│       │           │   │   │   │       # GET /v1/courses
│       │           │   │   │   ├── my-courses/
│       │           │   │   │   │   └── page.tsx  # Enrolled courses
│       │           │   │   │   │       # - In progress
│       │           │   │   │   │       # - Completed
│       │           │   │   │   │       # - Certificates
│       │           │   │   │   │       # BE: utility-be/learning
│       │           │   │   │   │       # GET /v1/users/me/courses
│       │           │   │   │   └── page.tsx  # Courses catalog
│       │           │   │   │       # BE: learning-be/course
│       │           │   │   │       # GET /v1/learning/courses
│       │           │   │   ├── mentorship/
│       │           │   │   │   ├── [sessionId]/
│       │           │   │   │   │   └── page.tsx  # Mentorship session details
│       │           │   │   │   │       # BE: users-be/mentorship
│       │           │   │   │   │       # GET /v1/users/me/mentorship/{session_id}
│       │           │   │   │   ├── find-mentor/
│       │           │   │   │   │   └── page.tsx  # Find a mentor
│       │           │   │   │   │       # BE: users-be/mentorship, search-be/query
│       │           │   │   │   │       # POST /v1/search/mentors
│       │           │   │   │   ├── my-mentees/
│       │           │   │   │   │   └── page.tsx  # Manage mentees
│       │           │   │   │   │       # BE: users-be/mentorship
│       │           │   │   │   │       # GET /v1/users/me/mentorship/mentees
│       │           │   │   │   └── page.tsx  # Mentorship dashboard
│       │           │   │   │       # BE: users-be/mentorship
│       │           │   │   │       # GET /v1/users/me/mentorship
│       │           │   │   ├── paths/
│       │           │   │   │   ├── [pathId]/
│       │           │   │   │   │   ├── progress/
│       │           │   │   │   │   │   └── page.tsx  # Learning path progress
│       │           │   │   │   │   │       # BE: users-be/learning_path
│       │           │   │   │   │   │       # GET /v1/users/me/learning-path/{path_id}/progress
│       │           │   │   │   │   └── page.tsx  # Learning path details
│       │           │   │   │   │       # - Courses
│       │           │   │   │   │       # - Milestones
│       │           │   │   │   │       # - Resources
│       │           │   │   │   │       # BE: users-be/learning_path
│       │           │   │   │   │       # GET /v1/users/me/learning-path/{path_id}
│       │           │   │   │   ├── discover/
│       │           │   │   │   │   └── page.tsx  # Discover learning paths
│       │           │   │   │   │       # BE: users-be/learning_path
│       │           │   │   │   │       # GET /v1/learning-paths/discover
│       │           │   │   │   └── page.tsx  # My learning paths
│       │           │   │   │       # BE: users-be/learning_path
│       │           │   │   │       # GET /v1/users/me/learning-path
│       │           │   │   └── skill-tests/
│       │           │   │       ├── [testId]/
│       │           │   │       │   ├── instructions/
│       │           │   │       │   │   └── page.tsx  # Test instructions
│       │           │   │       │   │       # BE: users-be/capabilities
│       │           │   │       │   │       # GET /v1/skill-tests/{test_id}
│       │           │   │       │   ├── results/
│       │           │   │       │   │   └── page.tsx  # Test results
│       │           │   │       │   │       # - Score
│       │           │   │       │   │       # - Percentile
│       │           │   │       │   │       # - Badge earned
│       │           │   │       │   │       # BE: users-be/capabilities
│       │           │   │       │   │       # GET /v1/skill-tests/{test_id}/results
│       │           │   │       │   └── take-test/
│       │           │   │       │       └── page.tsx  # Take the test
│       │           │   │       │           # - Timed test interface
│       │           │   │       │           # - Submit answers
│       │           │   │       │           # BE: users-be/capabilities
│       │           │   │       │           # POST /v1/skill-tests/{test_id}/submit
│       │           │   │       └── page.tsx  # Available skill tests
│       │           │   │           # - Browse by skill
│       │           │   │           # - Test history
│       │           │   │           # - Top scores
│       │           │   │           # BE: users-be/capabilities
│       │           │   │           # GET /v1/skill-tests
│       │           │   ├── network/
│       │           │   │   ├── connections/
│       │           │   │   │   ├── [userId]/
│       │           │   │   │   │   └── page.tsx  # Connection profile view
│       │           │   │   │   │       # BE: users-be/profile, users-be/connection
│       │           │   │   │   │       # GET /v1/users/{user_id}
│       │           │   │   │   │       # GET /v1/users/me/connections/{user_id}
│       │           │   │   │   ├── pending/
│       │           │   │   │   │   └── page.tsx  # Pending connection requests
│       │           │   │   │   │       # BE: users-be/connection
│       │           │   │   │   │       # GET /v1/users/me/connections/pending
│       │           │   │   │   └── page.tsx  # Connections list
│       │           │   │   │       # BE: users-be/connection
│       │           │   │   │       # GET /v1/users/me/connections
│       │           │   │   ├── groups/
│       │           │   │   │   ├── [groupId]/
│       │           │   │   │   │   ├── members/
│       │           │   │   │   │   │   └── page.tsx  # Group members
│       │           │   │   │   │   │       # BE: users-be/user_group
│       │           │   │   │   │   │       # GET /v1/groups/{group_id}/members
│       │           │   │   │   │   └── page.tsx  # Group details
│       │           │   │   │   │       # - Posts
│       │           │   │   │   │       # - Events
│       │           │   │   │   │       # - Resources
│       │           │   │   │   │       # BE: users-be/user_group
│       │           │   │   │   │       # GET /v1/groups/{group_id}
│       │           │   │   │   ├── discover/
│       │           │   │   │   │   └── page.tsx  # Discover groups
│       │           │   │   │   │       # BE: users-be/user_group
│       │           │   │   │   │       # GET /v1/groups/discover
│       │           │   │   │   └── page.tsx  # My groups
│       │           │   │   │       # BE: users-be/user_group
│       │           │   │   │       # GET /v1/users/me/groups
│       │           │   │   ├── recommendations/
│       │           │   │   │   └── page.tsx  # Connection recommendations
│       │           │   │   │       # - People you may know
│       │           │   │   │       # - Similar professionals
│       │           │   │   │       # BE: search-be/recommendation
│       │           │   │   │       # GET /v1/recommendations/connections
│       │           │   │   └── referrals/
│       │           │   │       ├── dashboard/
│       │           │   │       │   └── page.tsx  # Referral dashboard
│       │           │   │       │       # - Total referrals
│       │           │   │       │       # - Earnings
│       │           │   │       │       # - Conversion rate
│       │           │   │       │       # BE: users-be/referral
│       │           │   │       │       # GET /v1/users/me/referral-code
│       │           │   │       │       # GET /v1/referrals/analytics
│       │           │   │       └── page.tsx  # Referrals overview
│       │           │   │           # - Share referral code
│       │           │   │           # - Track referrals
│       │           │   │           # BE: users-be/referral
│       │           │   │           # GET /v1/referrals
│       │           │   ├── reputation/
│       │           │   │   ├── badges/
│       │           │   │   │   └── page.tsx  # Achievement badges
│       │           │   │   │       # - Top rated badge
│       │           │   │   │       # - Rising talent
│       │           │   │   │       # - Expert verified
│       │           │   │   │       # BE: reviews-be/badge (if exists) or users-be/profile
│       │           │   │   │       # GET /v1/reputation/badges
│       │           │   │   ├── disputes/
│       │           │   │   │   └── page.tsx  # Review disputes
│       │           │   │   │       # - Disputed reviews
│       │           │   │   │       # - Appeal status
│       │           │   │   │       # BE: reviews-be/dispute (if exists) or admin-be/case_mgmt
│       │           │   │   │       # GET /v1/reviews/disputes
│       │           │   │   ├── overview/
│       │           │   │   │   └── page.tsx  # Reputation overview
│       │           │   │   │       # - Overall score
│       │           │   │   │       # - Score breakdown
│       │           │   │   │       # - Recent changes
│       │           │   │   │       # BE: reviews-be/reputation (if exists) or reviews-be/review
│       │           │   │   │       # GET /v1/reputation/overview
│       │           │   │   ├── reviews-given/
│       │           │   │   │   └── page.tsx  # Reviews given
│       │           │   │   │       # - Reviews I left
│       │           │   │   │       # - Edit option (if within timeframe)
│       │           │   │   │       # BE: reviews-be/review
│       │           │   │   │       # GET /v1/reviews/given
│       │           │   │   └── reviews-received/
│       │           │   │       └── page.tsx  # Reviews received
│       │           │   │           # - Client reviews
│       │           │   │           # - Response option
│       │           │   │           # BE: reviews-be/review
│       │           │   │           # GET /v1/reviews/received
│       │           │   ├── reviews/
│       │           │   │   ├── analytics/
│       │           │   │   │   └── page.tsx  # Review analytics
│       │           │   │   │       # - Rating trends
│       │           │   │   │       # - Response rate
│       │           │   │   │       # - Sentiment analysis
│       │           │   │   │       # BE: reviews-be/review
│       │           │   │   │       # GET /v1/users/me/reviews/analytics
│       │           │   │   ├── disputes/
│       │           │   │   │   ├── [disputeId]/
│       │           │   │   │   │   └── page.tsx  # Review dispute details
│       │           │   │   │   │       # - Evidence submission
│       │           │   │   │   │       # - Admin review status
│       │           │   │   │   │       # BE: reviews-be/review, admin-be/case_mgmt
│       │           │   │   │   │       # GET /v1/reviews/{review_id}/dispute
│       │           │   │   │   └── page.tsx  # Review disputes list
│       │           │   │   │       # BE: reviews-be/review
│       │           │   │   │       # GET /v1/reviews/disputes
│       │           │   │   ├── given/
│       │           │   │   │   ├── [reviewId]/
│       │           │   │   │   │   ├── edit/
│       │           │   │   │   │   │   └── page.tsx  # Edit given review
│       │           │   │   │   │   │       # BE: reviews-be/review
│       │           │   │   │   │   │       # PUT /v1/reviews/{review_id}
│       │           │   │   │   │   └── page.tsx  # Given review details
│       │           │   │   │   │       # BE: reviews-be/review
│       │           │   │   │   │       # GET /v1/reviews/{review_id}
│       │           │   │   │   └── page.tsx  # Given reviews list
│       │           │   │   │       # BE: reviews-be/review
│       │           │   │   │       # GET /v1/users/me/reviews/given
│       │           │   │   ├── pending/
│       │           │   │   │   ├── [contractId]/
│       │           │   │   │   │   └── page.tsx  # Leave review form
│       │           │   │   │   │       # BE: reviews-be/review, contracts-be/contract
│       │           │   │   │   │       # GET /v1/contracts/{contract_id}
│       │           │   │   │   │       # POST /v1/reviews
│       │           │   │   │   └── page.tsx  # Pending reviews to complete
│       │           │   │   │       # BE: reviews-be/review
│       │           │   │   │       # GET /v1/reviews/pending
│       │           │   │   └── received/
│       │           │   │       ├── [reviewId]/
│       │           │   │       │   ├── respond/
│       │           │   │       │   │   └── page.tsx  # Respond to review
│       │           │   │       │   │       # BE: reviews-be/review
│       │           │   │       │   │       # POST /v1/reviews/{review_id}/response
│       │           │   │       │   └── page.tsx  # Review details
│       │           │   │       │       # BE: reviews-be/review
│       │           │   │       │       # GET /v1/reviews/{review_id}
│       │           │   │       └── page.tsx  # Received reviews list
│       │           │   │           # BE: reviews-be/review
│       │           │   │           # GET /v1/users/me/reviews/received
│       │           │   ├── search/
│       │           │   │   ├── advanced/
│       │           │   │   │   └── page.tsx  # Advanced search interface
│       │           │   │   │       # - Complex filters builder
│       │           │   │   │       # - Boolean operators
│       │           │   │   │       # - Saved search management
│       │           │   │   │       # BE: search-be/query
│       │           │   │   │       # POST /v1/search/advanced
│       │           │   │   ├── history/
│       │           │   │   │   └── page.tsx  # Search history
│       │           │   │   │       # BE: search-be/query
│       │           │   │   │       # GET /v1/search/history
│       │           │   │   ├── recommendations/
│       │           │   │   │   └── page.tsx  # Personalized recommendations
│       │           │   │   │       # - AI-powered job matches
│       │           │   │   │       # - Talent suggestions
│       │           │   │   │       # BE: search-be/recommendation
│       │           │   │   │       # GET /v1/recommendations/personalized
│       │           │   │   ├── saved/
│       │           │   │   │   ├── [searchId]/
│       │           │   │   │   │   ├── edit/
│       │           │   │   │   │   │   └── page.tsx  # Edit saved search
│       │           │   │   │   │   │       # BE: search-be/saved-search
│       │           │   │   │   │   │       # PUT /v1/search/saved-searches/{search_id}
│       │           │   │   │   │   └── results/
│       │           │   │   │   │       └── page.tsx  # View results from saved search
│       │           │   │   │   │           # BE: search-be/saved-search, search-be/query
│       │           │   │   │   │           # GET /v1/search/saved-searches/{search_id}/results
│       │           │   │   │   └── page.tsx  # Saved searches list (may be in combined, ensuring here)
│       │           │   │   └── trending/
│       │           │   │       └── page.tsx  # Trending searches and jobs
│       │           │   │           # BE: search-be/trending
│       │           │   │           # GET /v1/trending/jobs
│       │           │   │           # GET /v1/trending/skills
│       │           │   ├── sourcing/
│       │           │   │   ├── campaigns/
│       │           │   │   │   ├── [campaignId]/
│       │           │   │   │   │   ├── analytics/
│       │           │   │   │   │   │   └── page.tsx  # Campaign analytics
│       │           │   │   │   │   │       # - Reach
│       │           │   │   │   │   │       # - Engagement
│       │           │   │   │   │   │       # - Conversions
│       │           │   │   │   │   │       # BE: jobs-be/analytics
│       │           │   │   │   │   │       # GET /v1/sourcing/campaigns/{campaign_id}/analytics
│       │           │   │   │   │   ├── edit/
│       │           │   │   │   │   │   └── page.tsx  # Edit sourcing campaign
│       │           │   │   │   │   │       # BE: jobs-be/campaign (if exists) or jobs-be/job
│       │           │   │   │   │   │       # PUT /v1/sourcing/campaigns/{campaign_id}
│       │           │   │   │   │   └── page.tsx  # Campaign detail
│       │           │   │   │   │       # BE: jobs-be/campaign
│       │           │   │   │   │       # GET /v1/sourcing/campaigns/{campaign_id}
│       │           │   │   │   ├── create/
│       │           │   │   │   │   └── page.tsx  # Create sourcing campaign
│       │           │   │   │   │       # - Target criteria
│       │           │   │   │   │       # - Budget allocation
│       │           │   │   │   │       # - Messaging
│       │           │   │   │   │       # BE: jobs-be/campaign
│       │           │   │   │   │       # POST /v1/sourcing/campaigns
│       │           │   │   │   └── page.tsx  # Campaigns list
│       │           │   │   │       # BE: jobs-be/campaign
│       │           │   │   │       # GET /v1/sourcing/campaigns
│       │           │   │   ├── invitations/
│       │           │   │   │   ├── [invitationId]/
│       │           │   │   │   │   └── page.tsx  # Invitation detail
│       │           │   │   │   │       # - View status
│       │           │   │   │   │       # - Resend invitation
│       │           │   │   │   │       # BE: communications-be/invitation (if exists) or jobs-be/job
│       │           │   │   │   │       # GET /v1/sourcing/invitations/{invitation_id}
│       │           │   │   │   └── page.tsx  # Sent invitations
│       │           │   │   │       # - Invitation history
│       │           │   │   │       # - Response rates
│       │           │   │   │       # BE: communications-be/invitation
│       │           │   │   │       # GET /v1/sourcing/invitations
│       │           │   │   └── talent-pools/
│       │           │   │       ├── [poolId]/
│       │           │   │       │   ├── edit/
│       │           │   │       │   │   └── page.tsx  # Edit talent pool
│       │           │   │       │   │       # BE: users-be/talent_pool (if exists) or search-be/saved-search
│       │           │   │       │   │       # PUT /v1/sourcing/talent-pools/{pool_id}
│       │           │   │       │   ├── members/
│       │           │   │       │   │   └── page.tsx  # Pool members
│       │           │   │       │   │       # - Add/remove members
│       │           │   │       │   │       # - Bulk actions
│       │           │   │       │   │       # BE: users-be/talent_pool
│       │           │   │       │   │       # GET /v1/sourcing/talent-pools/{pool_id}/members
│       │           │   │       │   └── page.tsx  # Talent pool detail
│       │           │   │       │       # BE: users-be/talent_pool
│       │           │   │       │       # GET /v1/sourcing/talent-pools/{pool_id}
│       │           │   │       ├── create/
│       │           │   │       │   └── page.tsx  # Create talent pool
│       │           │   │       │       # BE: users-be/talent_pool
│       │           │   │       │       # POST /v1/sourcing/talent-pools
│       │           │   │       └── page.tsx  # Talent pools list
│       │           │   │           # BE: users-be/talent_pool
│       │           │   │           # GET /v1/sourcing/talent-pools
│       │           │   ├── specialized/
│       │           │   │   ├── plus/
│       │           │   │   │   ├── subscribe/
│       │           │   │   │   │   └── page.tsx  # Subscribe to Plus
│       │           │   │   │   │       # - Plans
│       │           │   │   │   │       # - Payment
│       │           │   │   │   │       # BE: subscriptions-be/subscription (financial-be)
│       │           │   │   │   │       # POST /v1/subscriptions/plus
│       │           │   │   │   └── page.tsx  # Plus membership overview
│       │           │   │   │       # - Benefits
│       │           │   │   │       # - Pricing
│       │           │   │   │       # - Comparison
│       │           │   │   │       # BE: subscriptions-be/plans
│       │           │   │   │       # GET /v1/subscriptions/plans/plus
│       │           │   │   ├── talent-cloud/
│       │           │   │   │   ├── projects/
│       │           │   │   │   │   └── page.tsx  # Talent Cloud exclusive projects
│       │           │   │   │   │       # - High-value projects
│       │           │   │   │   │       # - Direct invites
│       │           │   │   │   │       # BE: jobs-be/job (with exclusive flag)
│       │           │   │   │   │       # GET /v1/jobs/talent-cloud
│       │           │   │   │   └── page.tsx  # Talent Cloud info
│       │           │   │   │       # - Qualification
│       │           │   │   │       # - Apply
│       │           │   │   │       # - Status
│       │           │   │   │       # BE: users-be/badges
│       │           │   │   │       # GET /v1/programs/talent-cloud
│       │           │   │   └── top-rated/
│       │           │   │       ├── application/
│       │           │   │       │   └── page.tsx  # Apply for Top Rated
│       │           │   │       │       # - Eligibility check
│       │           │   │       │       # - Submit application
│       │           │   │       │       # BE: users-be/badges OR reviews-be/reputation
│       │           │   │       │       # POST /v1/users/me/top-rated/apply
│       │           │   │       └── page.tsx  # Top Rated program info
│       │           │   │           # - Benefits
│       │           │   │           # - Requirements
│       │           │   │           # - Application status
│       │           │   │           # BE: users-be/badges
│       │           │   │           # GET /v1/programs/top-rated
│       │           │   ├── spend-analytics/
│       │           │   │   ├── by-category/
│       │           │   │   │   └── page.tsx  # Spending by category
│       │           │   │   │       # - Job category breakdown
│       │           │   │   │       # - Trend analysis
│       │           │   │   │       # BE: financial-be/analytics (if exists) or financial-be/invoice
│       │           │   │   │       # GET /v1/analytics/spend/by-category
│       │           │   │   ├── by-department/
│       │           │   │   │   └── page.tsx  # Spending by department
│       │           │   │   │       # - Cost center breakdown
│       │           │   │   │       # - Budget vs actual
│       │           │   │   │       # BE: financial-be/budget, financial-be/invoice
│       │           │   │   │       # GET /v1/analytics/spend/by-department
│       │           │   │   ├── by-vendor/
│       │           │   │   │   └── page.tsx  # Spending by vendor
│       │           │   │   │       # - Top vendors
│       │           │   │   │       # - Vendor concentration risk
│       │           │   │   │       # BE: financial-be/invoice, users-be/profile
│       │           │   │   │       # GET /v1/analytics/spend/by-vendor
│       │           │   │   ├── forecasting/
│       │           │   │   │   └── page.tsx  # Spend forecasting
│       │           │   │   │       # - Projected spend
│       │           │   │   │       # - Budget burn rate
│       │           │   │   │       # - Alerts for overages
│       │           │   │   │       # BE: financial-be/forecast (if exists) or financial-be/analytics
│       │           │   │   │       # GET /v1/analytics/spend/forecast
│       │           │   │   └── page.tsx  # Spend analytics dashboard
│       │           │   │       # - Overview metrics
│       │           │   │       # - Charts & visualizations
│       │           │   │       # BE: financial-be/analytics
│       │           │   │       # GET /v1/analytics/spend
│       │           │   ├── talent/
│       │           │   │   ├── browse/
│       │           │   │   │   └── page.tsx  # Browse talent
│       │           │   │   │       # - Search freelancers
│       │           │   │   │       # - Filters (skills, rate, location)
│       │           │   │   │       # - Save to shortlist
│       │           │   │   │       # BE: search-be/query, users-be/profile
│       │           │   │   │       # POST /v1/search/freelancers
│       │           │   │   │       # GET /v1/search/freelancers?filters=...
│       │           │   │   ├── recommendations/
│       │           │   │   │   └── page.tsx  # AI-recommended talent for jobs
│       │           │   │   │       # BE: search-be/recommendation
│       │           │   │   │       # GET /v1/recommendations/talent?job_id={job_id}
│       │           │   │   ├── saved/
│       │           │   │   │   └── page.tsx  # Saved talent profiles
│       │           │   │   │       # BE: users-be/profile
│       │           │   │   │       # GET /v1/users/me/saved-profiles
│       │           │   │   └── shortlists/
│       │           │   │       ├── [shortlistId]/
│       │           │   │       │   ├── edit/
│       │           │   │       │   │   └── page.tsx  # Edit shortlist
│       │           │   │       │   │       # BE: jobs-be/shortlist
│       │           │   │       │   │       # PUT /v1/jobs/{job_id}/shortlists/{shortlist_id}
│       │           │   │       │   └── page.tsx  # Shortlist details
│       │           │   │       │       # - View candidates
│       │           │   │       │       # - Send invitations
│       │           │   │       │       # - Compare profiles
│       │           │   │       │       # BE: jobs-be/shortlist
│       │           │   │       │       # GET /v1/jobs/{job_id}/shortlists/{shortlist_id}
│       │           │   │       ├── new/
│       │           │   │       │   └── page.tsx  # Create shortlist
│       │           │   │       │       # BE: jobs-be/shortlist
│       │           │   │       │       # POST /v1/jobs/{job_id}/shortlists
│       │           │   │       └── page.tsx  # Shortlists overview
│       │           │   │           # BE: jobs-be/shortlist
│       │           │   │           # GET /v1/jobs/{job_id}/shortlists
│       │           │   ├── timesheets/
│       │           │   │   ├── [contractId]/
│       │           │   │   │   ├── [timesheetId]/
│       │           │   │   │   │   ├── edit/
│       │           │   │   │   │   │   └── page.tsx  # Edit timesheet
│       │           │   │   │   │   │       # BE: contracts-be/timesheet
│       │           │   │   │   │   │       # PUT /v1/contracts/{contract_id}/timesheets/{timesheet_id}
│       │           │   │   │   │   └── page.tsx  # Timesheet details
│       │           │   │   │   │       # - Hours breakdown
│       │           │   │   │   │       # - Approval status
│       │           │   │   │   │       # - Dispute options
│       │           │   │   │   │       # BE: contracts-be/timesheet
│       │           │   │   │   │       # GET /v1/contracts/{contract_id}/timesheets/{timesheet_id}
│       │           │   │   │   ├── new/
│       │           │   │   │   │   └── page.tsx  # Create timesheet
│       │           │   │   │   │       # BE: contracts-be/timesheet
│       │           │   │   │   │       # POST /v1/contracts/{contract_id}/timesheets
│       │           │   │   │   └── page.tsx  # Contract timesheets list
│       │           │   │   │       # BE: contracts-be/timesheet
│       │           │   │   │       # GET /v1/contracts/{contract_id}/timesheets
│       │           │   │   ├── approve/
│       │           │   │   │   └── page.tsx  # Timesheets pending approval (client)
│       │           │   │   │       # BE: contracts-be/timesheet
│       │           │   │   │       # GET /v1/timesheets/pending-approval
│       │           │   │   └── page.tsx  # All timesheets overview
│       │           │   │       # BE: contracts-be/timesheet
│       │           │   │       # GET /v1/timesheets
│       │           │   ├── vendor-management/
│       │           │   │   ├── blacklist/
│       │           │   │   │   ├── [userId]/
│       │           │   │   │   │   └── page.tsx  # Blacklist entry detail
│       │           │   │   │   │       # - Reason for blacklist
│       │           │   │   │   │       # - Remove option
│       │           │   │   │   │       # BE: users-be/org (blacklist subdomain)
│       │           │   │   │   │       # GET /v1/vendors/blacklist/{user_id}
│       │           │   │   │   │       # DELETE /v1/vendors/blacklist/{user_id}
│       │           │   │   │   └── page.tsx  # Blacklisted vendors
│       │           │   │   │       # BE: users-be/org
│       │           │   │   │       # GET /v1/vendors/blacklist
│       │           │   │   ├── compliance-docs/
│       │           │   │   │   ├── [vendorId]/
│       │           │   │   │   │   └── page.tsx  # Vendor compliance documents
│       │           │   │   │   │       # - W-9/W-8BEN forms
│       │           │   │   │   │       # - Insurance certificates
│       │           │   │   │   │       # - Background checks
│       │           │   │   │   │       # BE: admin-be/business_verification, storage-be/asset
│       │           │   │   │   │       # GET /v1/vendors/{vendor_id}/compliance-docs
│       │           │   │   │   └── page.tsx  # Compliance tracking
│       │           │   │   │       # - Expiring documents
│       │           │   │   │       # - Missing documents
│       │           │   │   │       # BE: admin-be/business_verification
│       │           │   │   │       # GET /v1/vendors/compliance-status
│       │           │   │   └── preferred/
│       │           │   │       ├── [vendorId]/
│       │           │   │       │   ├── history/
│       │           │   │       │   │   └── page.tsx  # Work history with vendor
│       │           │   │       │   │       # - Past contracts
│       │           │   │       │   │       # - Total spend
│       │           │   │       │   │       # - Reviews given
│       │           │   │       │   │       # BE: contracts-be/contract, financial-be/invoice
│       │           │   │       │   │       # GET /v1/vendors/{vendor_id}/history
│       │           │   │       │   ├── performance/
│       │           │   │       │   │   └── page.tsx  # Vendor performance metrics
│       │           │   │       │   │       # - Success rate
│       │           │   │       │   │       # - Average delivery time
│       │           │   │       │   │       # - Quality scores
│       │           │   │       │   │       # BE: users-be/org (vendor subdomain), contracts-be/contract
│       │           │   │       │   │       # GET /v1/vendors/{vendor_id}/performance
│       │           │   │       │   └── page.tsx  # Vendor detail
│       │           │   │       │       # BE: users-be/org (vendor subdomain)
│       │           │   │       │       # GET /v1/vendors/{vendor_id}
│       │           │   │       └── page.tsx  # Preferred vendors list
│       │           │   │           # - Star ratings
│       │           │   │           # - Quick invite
│       │           │   │           # BE: users-be/org
│       │           │   │           # GET /v1/vendors/preferred
│       │           │   └── work-diary/
│       │           │       ├── [contractId]/
│       │           │       │   ├── calendar/
│       │           │       │   │   └── page.tsx  # Calendar view of work diary
│       │           │       │   │       # BE: contracts-be/work_diary
│       │           │       │   │       # GET /v1/contracts/{contract_id}/work-diary/calendar
│       │           │       │   ├── screenshots/
│       │           │       │   │   └── page.tsx  # Screenshots management
│       │           │       │   │       # - View all screenshots
│       │           │       │   │       # - Delete sensitive ones
│       │           │       │   │       # - Privacy settings
│       │           │       │   │       # BE: contracts-be/work_diary, storage-be/asset
│       │           │       │   │       # GET /v1/contracts/{contract_id}/work-diary/screenshots
│       │           │       │   └── page.tsx  # Work diary detail
│       │           │       │       # BE: contracts-be/work_diary
│       │           │       │       # GET /v1/contracts/{contract_id}/work-diary
│       │           │       └── page.tsx  # Work diary overview (all contracts)
│       │           │           # BE: contracts-be/work_diary
│       │           │           # GET /v1/work-diary
│       │           ├── (public)/
│       │           │   ├── developers/
│       │           │   │   ├── api/
│       │           │   │   │   ├── authentication/
│       │           │   │   │   │   └── page.tsx  # OAuth & API keys
│       │           │   │   │   ├── endpoints/
│       │           │   │   │   │   ├── [category]/
│       │           │   │   │   │   │   └── page.tsx  # Endpoint category
│       │           │   │   │   │   └── page.tsx  # All endpoints
│       │           │   │   │   ├── getting-started/
│       │           │   │   │   │   └── page.tsx  # API getting started
│       │           │   │   │   ├── rate-limits/
│       │           │   │   │   │   └── page.tsx  # Rate limiting info
│       │           │   │   │   ├── webhooks/
│       │           │   │   │   │   └── page.tsx  # Webhook documentation
│       │           │   │   │   └── page.tsx  # API overview
│       │           │   │   ├── community/
│       │           │   │   │   ├── forum/
│       │           │   │   │   │   └── page.tsx  # Developer forum
│       │           │   │   │   └── github/
│       │           │   │   │       └── page.tsx  # GitHub repos
│       │           │   │   ├── plugins/
│       │           │   │   │   ├── shopify/
│       │           │   │   │   │   └── page.tsx  # Shopify app
│       │           │   │   │   ├── wordpress/
│       │           │   │   │   │   └── page.tsx  # WordPress plugin
│       │           │   │   │   └── page.tsx  # All plugins/integrations
│       │           │   │   ├── sample-code/
│       │           │   │   │   └── page.tsx  # Code examples
│       │           │   │   └── sdks/
│       │           │   │       ├── javascript/
│       │           │   │       │   └── page.tsx  # JavaScript SDK
│       │           │   │       ├── python/
│       │           │   │       │   └── page.tsx  # Python SDK
│       │           │   │       ├── ruby/
│       │           │   │       │   └── page.tsx  # Ruby SDK
│       │           │   │       └── page.tsx  # All SDKs
│       │           │   ├── legal/
│       │           │   │   ├── accessibility/
│       │           │   │   │   └── page.tsx  # Accessibility statement
│       │           │   │   ├── compliance/
│       │           │   │   │   ├── aml/
│       │           │   │   │   │   └── page.tsx  # AML policy
│       │           │   │   │   ├── kyc/
│       │           │   │   │   │   └── page.tsx  # KYC policy
│       │           │   │   │   └── page.tsx  # Compliance overview
│       │           │   │   ├── dmca/
│       │           │   │   │   └── page.tsx  # DMCA policy
│       │           │   │   ├── licenses/
│       │           │   │   │   └── page.tsx  # Open source licenses
│       │           │   │   ├── privacy/
│       │           │   │   │   ├── ccpa/
│       │           │   │   │   │   └── page.tsx  # CCPA information
│       │           │   │   │   ├── cookies/
│       │           │   │   │   │   └── page.tsx  # Cookie Policy
│       │           │   │   │   ├── gdpr/
│       │           │   │   │   │   └── page.tsx  # GDPR information
│       │           │   │   │   └── policy/
│       │           │   │   │       └── page.tsx  # Privacy Policy
│       │           │   │   └── terms/
│       │           │   │       ├── client/
│       │           │   │       │   └── page.tsx  # Client Agreement
│       │           │   │       ├── freelancer/
│       │           │   │       │   └── page.tsx  # Freelancer Agreement
│       │           │   │       ├── service/
│       │           │   │       │   └── page.tsx  # Terms of Service
│       │           │   │       └── page.tsx  # Legal terms overview
│       │           │   ├── resources/
│       │           │   │   ├── api/
│       │           │   │   │   ├── changelog/
│       │           │   │   │   │   └── page.tsx  # API changelog
│       │           │   │   │   ├── docs/
│       │           │   │   │   │   └── page.tsx  # API documentation
│       │           │   │   │   └── reference/
│       │           │   │   │       └── page.tsx  # API reference
│       │           │   │   ├── case-studies/
│       │           │   │   │   ├── [slug]/
│       │           │   │   │   │   └── page.tsx  # Case study detail
│       │           │   │   │   └── page.tsx  # All case studies
│       │           │   │   ├── guides/
│       │           │   │   │   ├── [slug]/
│       │           │   │   │   │   └── page.tsx  # Guide detail
│       │           │   │   │   └── page.tsx  # All guides
│       │           │   │   │       # - Getting started
│       │           │   │   │       # - Best practices
│       │           │   │   │       # - How-tos
│       │           │   │   ├── tools/
│       │           │   │   │   ├── budget-estimator/
│       │           │   │   │   │   └── page.tsx  # Project budget estimator
│       │           │   │   │   ├── rate-calculator/
│       │           │   │   │   │   └── page.tsx  # Hourly rate calculator
│       │           │   │   │   └── roi-calculator/
│       │           │   │   │       └── page.tsx  # ROI calculator
│       │           │   │   ├── webinars/
│       │           │   │   │   ├── [slug]/
│       │           │   │   │   │   ├── register/
│       │           │   │   │   │   │   └── page.tsx  # Webinar registration
│       │           │   │   │   │   │       # BE: communications-be/events
│       │           │   │   │   │   │       # POST /v1/webinars/{webinar_id}/register
│       │           │   │   │   │   └── page.tsx  # Webinar detail
│       │           │   │   │   ├── on-demand/
│       │           │   │   │   │   └── page.tsx  # On-demand recordings
│       │           │   │   │   └── upcoming/
│       │           │   │   │       └── page.tsx  # Upcoming webinars
│       │           │   │   └── whitepapers/
│       │           │   │       ├── [slug]/
│       │           │   │       │   └── page.tsx  # Whitepaper detail
│       │           │   │       │       # - Download gated content
│       │           │   │       │       # BE: communications-be/leads
│       │           │   │       │       # POST /v1/leads/download
│       │           │   │       └── page.tsx  # All whitepapers
│       │           │   └── solutions/
│       │           │       ├── agencies/
│       │           │       │   └── page.tsx  # Agency solutions
│       │           │       │       # - White-label
│       │           │       │       # - Multi-client
│       │           │       │       # - Revenue share
│       │           │       ├── by-industry/
│       │           │       │   ├── finance/
│       │           │       │   │   └── page.tsx  # Finance solutions
│       │           │       │   ├── marketing/
│       │           │       │   │   └── page.tsx  # Marketing solutions
│       │           │       │   ├── tech/
│       │           │       │   │   └── page.tsx  # Tech industry solutions
│       │           │       │   └── page.tsx  # Browse by industry
│       │           │       ├── by-role/
│       │           │       │   ├── designers/
│       │           │       │   │   └── page.tsx  # Designer hiring solutions
│       │           │       │   ├── developers/
│       │           │       │   │   └── page.tsx  # Developer hiring solutions
│       │           │       │   ├── writers/
│       │           │       │   │   └── page.tsx  # Writer hiring solutions
│       │           │       │   └── page.tsx  # Browse by role
│       │           │       └── enterprise/
│       │           │           ├── contact-sales/
│       │           │           │   └── page.tsx  # Enterprise contact form
│       │           │           │       # BE: communications-be/contact
│       │           │           │       # POST /v1/contact/enterprise
│       │           │           └── page.tsx  # Enterprise solutions
│       │           │               # - MSA contracts
│       │           │               # - Dedicated support
│       │           │               # - Volume pricing
│       │           ├── developers/
│       │           │   ├── api-reference/
│       │           │   │   ├── [endpoint]/
│       │           │   │   │   └── page.tsx  # API endpoint reference
│       │           │   │   │       # BE: static (OpenAPI spec)
│       │           │   │   └── page.tsx  # API reference home
│       │           │   │       # BE: static (OpenAPI spec)
│       │           │   ├── docs/
│       │           │   │   ├── [section]/
│       │           │   │   │   └── page.tsx  # Documentation section
│       │           │   │   │       # BE: static or CMS
│       │           │   │   │       # GET /v1/content/docs/{section}
│       │           │   │   └── page.tsx  # API documentation home
│       │           │   │       # BE: static
│       │           │   ├── sandbox/
│       │           │   │   └── page.tsx  # API sandbox/playground
│       │           │   │       # BE: developer API
│       │           │   │       # POST /v1/developer/sandbox/execute
│       │           │   ├── sdks/
│       │           │   │   └── page.tsx  # SDK downloads and docs
│       │           │   │       # BE: static
│       │           │   └── webhooks/
│       │           │       └── page.tsx  # Webhooks documentation
│       │           │           # BE: static
│       │           ├── enterprise/
│       │           │   ├── case-studies/
│       │           │   │   └── page.tsx  # Enterprise case studies
│       │           │   │       # BE: CMS
│       │           │   │       # GET /v1/content/case-studies?type=enterprise
│       │           │   ├── contact/
│       │           │   │   └── page.tsx  # Enterprise contact/demo request
│       │           │   │       # BE: communications-be
│       │           │   │       # POST /v1/contact/enterprise
│       │           │   ├── pricing/
│       │           │   │   └── page.tsx  # Enterprise pricing
│       │           │   │       # BE: financial-be/subscription
│       │           │   │       # GET /v1/subscriptions/plans?type=enterprise
│       │           │   └── solutions/
│       │           │       ├── managed-services/
│       │           │       │   └── page.tsx  # Managed services offering
│       │           │       │       # BE: none (marketing content)
│       │           │       ├── staffing/
│       │           │       │   └── page.tsx  # Enterprise staffing solutions
│       │           │       │       # BE: none (marketing content)
│       │           │       └── page.tsx  # Enterprise solutions overview
│       │           │           # BE: none (marketing content)
│       │           ├── legal/
│       │           │   ├── accessibility/
│       │           │   │   └── page.tsx  # Accessibility statement
│       │           │   │       # BE: none (static content)
│       │           │   ├── compliance/
│       │           │   │   ├── ccpa/
│       │           │   │   │   └── page.tsx  # CCPA compliance
│       │           │   │   │       # BE: none (static content)
│       │           │   │   ├── gdpr/
│       │           │   │   │   └── page.tsx  # GDPR compliance
│       │           │   │   │       # BE: none (static content)
│       │           │   │   └── page.tsx  # Compliance overview
│       │           │   │       # BE: none (static content)
│       │           │   ├── dmca/
│       │           │   │   └── page.tsx  # DMCA policy
│       │           │   │       # BE: none (static content)
│       │           │   ├── ip-policy/
│       │           │   │   └── page.tsx  # Intellectual property policy
│       │           │   │       # BE: none (static content)
│       │           │   ├── privacy/
│       │           │   │   ├── cookie-policy/
│       │           │   │   │   └── page.tsx  # Cookie policy
│       │           │   │   │       # BE: none (static content)
│       │           │   │   ├── data-processing/
│       │           │   │   │   └── page.tsx  # Data processing agreement
│       │           │   │   │       # BE: none (static content)
│       │           │   │   └── page.tsx  # Privacy policy
│       │           │   │       # BE: none (static content)
│       │           │   └── terms/
│       │           │       ├── client/
│       │           │       │   └── page.tsx  # Client terms of service
│       │           │       │       # BE: none (static content)
│       │           │       ├── freelancer/
│       │           │       │   └── page.tsx  # Freelancer terms of service
│       │           │       │       # BE: none (static content with version from CMS)
│       │           │       └── page.tsx  # General terms
│       │           │           # BE: none (static content)
│       │           ├── resources/
│       │           │   ├── blog/
│       │           │   │   ├── [postId]/
│       │           │   │   │   └── page.tsx  # Blog post
│       │           │   │   │       # BE: CMS
│       │           │   │   │       # GET /v1/content/blog/{post_id}
│       │           │   │   ├── category/
│       │           │   │   │   └── [categoryId]/
│       │           │   │   │       └── page.tsx  # Blog category
│       │           │   │   │           # BE: CMS
│       │           │   │   │           # GET /v1/content/blog?category={category_id}
│       │           │   │   └── page.tsx  # Blog home
│       │           │   │       # BE: CMS
│       │           │   │       # GET /v1/content/blog
│       │           │   ├── case-studies/
│       │           │   │   ├── [caseStudyId]/
│       │           │   │   │   └── page.tsx  # Case study detail
│       │           │   │   │       # BE: CMS
│       │           │   │   │       # GET /v1/content/case-studies/{case_study_id}
│       │           │   │   └── page.tsx  # Case studies list
│       │           │   │       # BE: CMS
│       │           │   │       # GET /v1/content/case-studies
│       │           │   ├── faq/
│       │           │   │   └── page.tsx  # Frequently asked questions
│       │           │   │       # BE: CMS
│       │           │   │       # GET /v1/content/faq
│       │           │   ├── guides/
│       │           │   │   ├── [guideId]/
│       │           │   │   │   └── page.tsx  # Guide detail
│       │           │   │   │       # BE: CMS or static
│       │           │   │   │       # GET /v1/content/guides/{guide_id}
│       │           │   │   ├── client/
│       │           │   │   │   └── page.tsx  # Client guides
│       │           │   │   │       # BE: CMS
│       │           │   │   │       # GET /v1/content/guides?category=client
│       │           │   │   ├── freelancer/
│       │           │   │   │   └── page.tsx  # Freelancer guides
│       │           │   │   │       # BE: CMS
│       │           │   │   │       # GET /v1/content/guides?category=freelancer
│       │           │   │   └── page.tsx  # All guides
│       │           │   │       # BE: CMS
│       │           │   │       # GET /v1/content/guides
│       │           │   ├── tutorials/
│       │           │   │   ├── [tutorialId]/
│       │           │   │   │   └── page.tsx  # Tutorial detail
│       │           │   │   │       # BE: CMS
│       │           │   │   │       # GET /v1/content/tutorials/{tutorial_id}
│       │           │   │   └── page.tsx  # Tutorials list
│       │           │   │       # BE: CMS
│       │           │   │       # GET /v1/content/tutorials
│       │           │   └── webinars/
│       │           │       ├── [webinarId]/
│       │           │       │   └── page.tsx  # Webinar detail & registration
│       │           │       │       # BE: CMS + registration system
│       │           │       │       # GET /v1/content/webinars/{webinar_id}
│       │           │       │       # POST /v1/webinars/{webinar_id}/register
│       │           │       └── page.tsx  # Upcoming webinars
│       │           │           # BE: CMS
│       │           │           # GET /v1/content/webinars
│       │           ├── security/
│       │           │   ├── bug-bounty/
│       │           │   │   └── page.tsx  # Bug bounty program
│       │           │   │       # BE: none (static content)
│       │           │   ├── certifications/
│       │           │   │   └── page.tsx  # Security certifications (SOC2, ISO, etc.)
│       │           │   │       # BE: none (static content)
│       │           │   ├── overview/
│       │           │   │   └── page.tsx  # Security overview
│       │           │   │       # BE: none (static content)
│       │           │   └── responsible-disclosure/
│       │           │       └── page.tsx  # Responsible disclosure policy
│       │           │           # BE: none (static content)
│       │           ├── status/
│       │           │   ├── current/
│       │           │   │   └── page.tsx  # Current system status
│       │           │   │       # BE: utility/status
│       │           │   │       # GET /v1/status/current
│       │           │   ├── history/
│       │           │   │   └── page.tsx  # Status history
│       │           │   │       # BE: utility/status
│       │           │   │       # GET /v1/status/history
│       │           │   └── subscribe/
│       │           │       └── page.tsx  # Subscribe to status updates
│       │           │           # BE: communications-be
│       │           │           # POST /v1/notifications/status-subscribe
│       │           └── transparency/
│       │               └── page.tsx  # Transparency report
│       │                   # - User statistics
│       │                   # - Moderation actions
│       │                   # - Government requests
│       │                   # BE: admin-be/reporting
│       │                   # GET /v1/public/transparency-report
│       └── middleware.ts  # Next.js middleware
│           # - CSP headers
│           # - CORS configuration
│           # - Rate limiting
│           # - Security headers (X-Frame-Options, etc.)
│           # - Security headers
│           # - CSRF protection
│           # - Rate limiting (client-side)
│           # - Auth checks
├── config/
│   ├── constants/
│   │   ├── api-endpoints.ts  # API endpoint constants
│   │   ├── app-config.ts  # App configuration
│   │   ├── performance-budgets.ts  # Performance budgets
│   │   └── third-party-keys.ts  # Third-party service keys
│   ├── environments/
│   │   ├── development.ts  # Dev config
│   │   ├── production.ts  # Production config
│   │   └── staging.ts  # Staging config
│   ├── feature-flags.ts  # Feature flags config
│   └── index.ts  # Config loader
├── docs/
│   ├── adr/  # Architecture Decision Records
│   │   ├── ...
│   │   ├── 001-monorepo-structure.md
│   │   ├── 002-state-management.md
│   │   ├── 003-authentication-approach.md
│   │   └── 004-component-library.md
│   ├── api/
│   │   ├── endpoints/
│   │   │   ├── contracts.md
│   │   │   ├── jobs.md
│   │   │   ├── messages.md
│   │   │   ├── notifications.md
│   │   │   ├── payments.md
│   │   │   ├── proposals.md
│   │   │   └── users.md
│   │   ├── authentication.md
│   │   ├── errors.md
│   │   ├── getting-started.md
│   │   ├── rate-limiting.md
│   │   └── webhooks.md
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
│   │   ├── backend-integration/
│   │   │   ├── api-patterns.md  # API patterns
│   │   │   ├── caching.md  # Caching strategy
│   │   │   ├── error-handling.md  # Error handling
│   │   │   └── real-time.md  # WebSocket/SSE
│   │   ├── diagrams/
│   │   │   ├── auth-flow.mmd
│   │   │   ├── data-flow.mmd
│   │   │   └── system-architecture.mmd
│   │   ├── frontend/
│   │   │   ├── authentication.md  # Auth flow
│   │   │   ├── routing.md  # Routing strategy
│   │   │   ├── state-management.md  # State management
│   │   │   └── structure.md  # Folder structure
│   │   ├── mobile/
│   │   │   ├── architecture.md  # Mobile architecture
│   │   │   ├── offline-support.md  # Offline capabilities
│   │   │   └── performance.md  # Performance optimization
│   │   ├── authentication-flow.md
│   │   ├── data-fetching-patterns.md
│   │   ├── frontend-architecture.md
│   │   ├── microservices-integration.md
│   │   ├── overview.md  # System overview
│   │   ├── routing-strategy.md
│   │   └── state-management.md
│   ├── components/
│   │   ├── atoms/
│   │   │   ├── ...
│   │   │   ├── button.md
│   │   │   └── input.md
│   │   ├── molecules/
│   │   │   ├── ...
│   │   │   ├── card.md
│   │   │   └── form-field.md
│   │   ├── organisms/
│   │   │   ├── ...
│   │   │   ├── header.md
│   │   │   └── sidebar.md
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
│       ├── contributing/
│       │   ├── code-review.md
│       │   ├── getting-started.md
│       │   └── pull-requests.md
│       ├── deployment/
│       │   ├── ci-cd.md
│       │   ├── mobile.md
│       │   └── web.md
│       ├── development/
│       │   ├── coding-standards.md
│       │   ├── debugging.md
│       │   ├── git-workflow.md
│       │   └── testing.md
│       ├── setup/
│       │   ├── environment-variables.md
│       │   ├── local-development.md
│       │   └── troubleshooting.md
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
│   │       │   ├── testing/
│   │       │   │   ├── a11y-test-utils.ts  # Testing utilities
│   │       │   │   └── axe-config.ts  # axe-core configuration
│   │       │   ├── a11y-utils.ts  # Accessibility utilities
│   │       │   ├── aria-utils.ts  # ARIA utilities
│   │       │   ├── focus-management.ts  # Focus management
│   │       │   ├── keyboard-navigation.ts  # Keyboard navigation
│   │       │   └── screen-reader.ts  # Screen reader utilities
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
│   │       │   ├── achievements/
│   │       │   │   ├── definitions/
│   │       │   │   │   ├── client-achievements.ts  # Client achievements
│   │       │   │   │   ├── freelancer-achievements.ts  # Freelancer achievements
│   │       │   │   │   └── platform-achievements.ts  # Platform-wide achievements
│   │       │   │   └── utils.ts  # Achievement utilities
│   │       │   ├── analytics/
│   │       │   │   ├── api/
│   │       │   │   │   └── analytics-api.ts  # Analytics API
│   │       │   │   │       # BE: utility/analytics
│   │       │   │   │       # POST /v1/analytics/events
│   │       │   │   │       # POST /v1/analytics/page-view
│   │       │   │   ├── events/
│   │       │   │   │   ├── contract-events.ts  # Contract events
│   │       │   │   │   ├── job-events.ts  # Job events
│   │       │   │   │   ├── payment-events.ts  # Payment events
│   │       │   │   │   ├── proposal-events.ts  # Proposal events
│   │       │   │   │   ├── system-events.ts  # System events
│   │       │   │   │   └── user-events.ts  # User events
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useAnalytics.ts  # Track events
│   │       │   │   │   ├── useConversionTracking.ts  # Track conversions
│   │       │   │   │   └── usePageView.ts  # Track page views
│   │       │   │   ├── utils/
│   │       │   │   │   ├── anonymize.ts  # Anonymize PII
│   │       │   │   │   ├── batch-sender.ts  # Batch events
│   │       │   │   │   └── event-builder.ts  # Build events
│   │       │   │   └── types.ts  # Analytics types
│   │       │   ├── auctions/
│   │       │   │   ├── api/
│   │       │   │   │   └── auctions-api.ts  # Auctions API client
│   │       │   │   │       # BE: proposals-be/auction
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useActiveAuctions.ts  # Active auctions list
│   │       │   │   │   ├── useAuction.ts  # Single auction
│   │       │   │   │   ├── useAuctionBid.ts  # Place bid
│   │       │   │   │   └── useAuctionHistory.ts  # Bid history
│   │       │   │   ├── queries/
│   │       │   │   │   ├── auctions-mutations.ts  # Auction mutations
│   │       │   │   │   └── auctions-queries.ts  # Auction queries
│   │       │   │   ├── store/
│   │       │   │   │   └── auction-store.ts  # Real-time auction state (Zustand)
│   │       │   │   └── types.ts  # Auction types
│   │       │   ├── bidding/
│   │       │   │   ├── api/
│   │       │   │   │   ├── bid-api.ts  # Bid placement API
│   │       │   │   │   │   # BE: proposals-be/bid
│   │       │   │   │   └── bid-strategy-api.ts  # Bid strategy API
│   │       │   │   │       # BE: proposals-be/bid-strategy
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useBidAnalytics.ts  # Bid analytics
│   │       │   │   │   ├── useBidHistory.ts  # Bid history
│   │       │   │   │   ├── useBidStrategies.ts  # List strategies
│   │       │   │   │   ├── useBidStrategy.ts  # Bid strategy management
│   │       │   │   │   └── usePlaceBid.ts  # Place bid
│   │       │   │   ├── queries/
│   │       │   │   │   ├── bidding-mutations.ts  # Bidding mutations
│   │       │   │   │   └── bidding-queries.ts  # Bidding queries
│   │       │   │   └── types.ts  # Bidding types
│   │       │   ├── collaboration/
│   │       │   │   ├── api/
│   │       │   │   │   └── collaboration-api.ts  # Collaboration API
│   │       │   │   │       # BE: communications-be/realtime
│   │       │   │   ├── components/
│   │       │   │   │   ├── ActiveUsers/
│   │       │   │   │   │   ├── ActiveUsers.tsx
│   │       │   │   │   │   └── ActiveUsers.types.ts
│   │       │   │   │   ├── CollaboratorCursor/
│   │       │   │   │   │   ├── CollaboratorCursor.tsx
│   │       │   │   │   │   └── CollaboratorCursor.types.ts
│   │       │   │   │   └── PresenceIndicator/
│   │       │   │   │       ├── PresenceIndicator.tsx
│   │       │   │   │       └── PresenceIndicator.types.ts
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useCollaboration.ts  # Collaboration session
│   │       │   │   │   ├── useCursors.ts  # Cursor tracking
│   │       │   │   │   ├── usePresence.ts  # User presence
│   │       │   │   │   └── useSharedState.ts  # Shared state sync
│   │       │   │   ├── providers/
│   │       │   │   │   └── CollaborationProvider.tsx  # Collab context
│   │       │   │   └── types.ts  # Collaboration types
│   │       │   ├── compliance/
│   │       │   │   ├── api/
│   │       │   │   │   ├── compliance-api.ts  # Compliance API
│   │       │   │   │   │   # BE: users-be/compliance
│   │       │   │   │   └── tax-profile-api.ts  # Tax profile API
│   │       │   │   │       # BE: users-be/compliance, financial-be/tax
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useComplianceDocuments.ts  # Document management
│   │       │   │   │   ├── useComplianceProfile.ts  # Compliance profile
│   │       │   │   │   ├── useTaxProfile.ts  # Tax profile management
│   │       │   │   │   └── useTaxReports.ts  # Tax reports
│   │       │   │   ├── queries/
│   │       │   │   │   ├── compliance-mutations.ts  # Compliance mutations
│   │       │   │   │   └── compliance-queries.ts  # Compliance queries
│   │       │   │   └── types.ts  # Compliance types
│   │       │   ├── connects/
│   │       │   │   ├── api/
│   │       │   │   │   └── connects-api.ts  # Connects API client
│   │       │   │   │       # BE: proposals-be/connect
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useConnectPackages.ts  # Available packages
│   │       │   │   │   ├── useConnectRefund.ts  # Request refund
│   │       │   │   │   ├── useConnects.ts  # Connects balance and history
│   │       │   │   │   └── usePurchaseConnects.ts  # Purchase connects
│   │       │   │   ├── queries/
│   │       │   │   │   ├── connects-mutations.ts  # Connect mutations
│   │       │   │   │   └── connects-queries.ts  # Connect queries
│   │       │   │   └── types.ts  # Connect types
│   │       │   ├── deliverables/
│   │       │   │   ├── api/
│   │       │   │   │   └── deliverables-api.ts  # Deliverables API client
│   │       │   │   │       # BE: contracts-be/deliverable, storage-be/asset
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useDeliverable.ts  # Single deliverable
│   │       │   │   │   ├── useDeliverableRevisions.ts  # Revision management
│   │       │   │   │   ├── useDeliverables.ts  # Deliverables list
│   │       │   │   │   ├── useReviewDeliverable.ts  # Review deliverable (client)
│   │       │   │   │   └── useUploadDeliverable.ts  # Upload deliverable
│   │       │   │   ├── queries/
│   │       │   │   │   ├── deliverables-mutations.ts  # Deliverable mutations
│   │       │   │   │   └── deliverables-queries.ts  # Deliverable queries
│   │       │   │   └── types.ts  # Deliverable types
│   │       │   ├── experiments/
│   │       │   │   ├── api/
│   │       │   │   │   └── experiments-api.ts  # Experiments API
│   │       │   │   │       # BE: utility-be/experiments
│   │       │   │   │       # GET /v1/experiments/active
│   │       │   │   │       # POST /v1/experiments/{experiment_id}/track
│   │       │   │   │       # BE: utility/experiments (if exists) or utility/flags
│   │       │   │   │       # POST /v1/experiments/track-event
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useExperiment.ts  # Get experiment variant
│   │       │   │   │   │   # Experiment hook
│   │       │   │   │   ├── useExperimentTracking.ts  # Track experiment events
│   │       │   │   │   └── useFeatureVariant.ts  # Feature variant
│   │       │   │   ├── providers/
│   │       │   │   │   ├── ExperimentProvider.tsx  # Experiment context
│   │       │   │   │   └── ExperimentsProvider.tsx  # Experiments context
│   │       │   │   ├── utils/
│   │       │   │   │   ├── experiment-storage.ts  # Store assignments
│   │       │   │   │   ├── tracking.ts  # Track events
│   │       │   │   │   └── variant-assignment.ts  # Assign variant
│   │       │   │   └── types.ts  # Experiment types
│   │       │   ├── feature-flags/
│   │       │   │   ├── api/
│   │       │   │   │   └── flags-api.ts  # Feature flags API
│   │       │   │   │       # BE: utility/flags
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useFeatureFlag.ts  # Check single flag
│   │       │   │   │   ├── useFeatureFlags.ts  # Get all flags
│   │       │   │   │   └── useFeatureFlagVariant.ts  # A/B test variant
│   │       │   │   ├── queries/
│   │       │   │   │   └── flags-queries.ts  # Flag queries
│   │       │   │   ├── store/
│   │       │   │   │   └── flags-store.ts  # Flags state (Zustand)
│   │       │   │   └── types.ts  # Flag types
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
│   │       │   ├── i18n/
│   │       │   │   ├── components/
│   │       │   │   │   ├── FormattedMessage/
│   │       │   │   │   │   └── FormattedMessage.tsx
│   │       │   │   │   ├── LocaleSwitcher/
│   │       │   │   │   │   ├── LocaleSwitcher.native.tsx
│   │       │   │   │   │   ├── LocaleSwitcher.tsx
│   │       │   │   │   │   └── LocaleSwitcher.web.tsx
│   │       │   │   │   └── TranslationProvider/
│   │       │   │   │       └── TranslationProvider.tsx
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useCurrencyFormat.ts  # Currency formatting
│   │       │   │   │   ├── useDateFormat.ts  # Date formatting
│   │       │   │   │   ├── useNumberFormat.ts  # Number formatting
│   │       │   │   │   ├── useRTL.ts  # RTL detection
│   │       │   │   │   └── useTranslation.ts  # (already exists, enhanced)
│   │       │   │   ├── tools/
│   │       │   │   │   ├── currency-formatter.ts  # Currency formatting
│   │       │   │   │   ├── missing-keys-detector.ts  # Detect missing keys
│   │       │   │   │   ├── pluralization-helper.ts  # Pluralization
│   │       │   │   │   └── translation-manager.ts  # Manage translations
│   │       │   │   └── utils/
│   │       │   │       ├── fallback-loader.ts  # Fallback translations
│   │       │   │       ├── interpolation.ts  # Message interpolation
│   │       │   │       └── locale-detection.ts  # Auto-detect locale
│   │       │   ├── interviews/
│   │       │   │   ├── api/
│   │       │   │   │   └── interviews-api.ts  # Interviews API client
│   │       │   │   │       # BE: proposals-be/interview
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useInterview.ts  # Single interview
│   │       │   │   │   ├── useInterviewFeedback.ts  # Interview feedback
│   │       │   │   │   ├── useInterviews.ts  # Interviews list
│   │       │   │   │   └── useScheduleInterview.ts  # Schedule interview
│   │       │   │   ├── queries/
│   │       │   │   │   ├── interviews-mutations.ts  # Interview mutations
│   │       │   │   │   └── interviews-queries.ts  # Interview queries
│   │       │   │   └── types.ts  # Interview types
│   │       │   ├── invitations/
│   │       │   │   ├── api/
│   │       │   │   │   ├── job-invitations-api.ts  # Job invitations API (client)
│   │       │   │   │   │   # BE: jobs-be/invitation
│   │       │   │   │   └── proposal-invites-api.ts  # Proposal invites API (freelancer)
│   │       │   │   │       # BE: proposals-be/invite
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useAcceptInvite.ts  # Accept invite (freelancer)
│   │       │   │   │   ├── useDeclineInvite.ts  # Decline invite (freelancer)
│   │       │   │   │   ├── useInvitationAnalytics.ts  # Invitation metrics
│   │       │   │   │   ├── useInvitations.ts  # Invitations management
│   │       │   │   │   └── useSendInvitation.ts  # Send invitation (client)
│   │       │   │   ├── queries/
│   │       │   │   │   ├── invitations-mutations.ts  # Invitation mutations
│   │       │   │   │   └── invitations-queries.ts  # Invitation queries
│   │       │   │   └── types.ts  # Invitation types
│   │       │   ├── learning/
│   │       │   │   ├── api/
│   │       │   │   │   ├── learning-paths-api.ts  # Learning paths API
│   │       │   │   │   │   # BE: users-be/learning_path
│   │       │   │   │   └── mentorship-api.ts  # Mentorship API
│   │       │   │   │       # BE: users-be/mentorship
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useAchievements.ts  # Achievements/badges
│   │       │   │   │   ├── useLearningPath.ts  # Single learning path
│   │       │   │   │   ├── useLearningPaths.ts  # Learning paths list
│   │       │   │   │   ├── useLearningProgress.ts  # Track progress
│   │       │   │   │   └── useMentorship.ts  # Mentorship management
│   │       │   │   ├── queries/
│   │       │   │   │   ├── learning-mutations.ts  # Learning mutations
│   │       │   │   │   └── learning-queries.ts  # Learning queries
│   │       │   │   └── types.ts  # Learning types
│   │       │   ├── maps/
│   │       │   │   ├── components/
│   │       │   │   │   ├── LocationPicker.tsx  # Location picker
│   │       │   │   │   ├── Map.native.tsx  # Map component (mobile)
│   │       │   │   │   ├── Map.tsx  # Map component (web)
│   │       │   │   │   └── Marker.tsx  # Map marker
│   │       │   │   └── types.ts  # Map types
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
│   │       │   ├── negotiations/
│   │       │   │   ├── api/
│   │       │   │   │   └── negotiations-api.ts  # Negotiations API client
│   │       │   │   │       # BE: proposals-be/negotiation
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useNegotiation.ts  # Single negotiation
│   │       │   │   │   ├── useNegotiationHistory.ts  # Negotiation history
│   │       │   │   │   └── useNegotiationOffer.ts  # Make/accept/reject offer
│   │       │   │   ├── queries/
│   │       │   │   │   ├── negotiations-mutations.ts  # Negotiation mutations
│   │       │   │   │   └── negotiations-queries.ts  # Negotiation queries
│   │       │   │   └── types.ts  # Negotiation types
│   │       │   ├── networking/
│   │       │   │   ├── api/
│   │       │   │   │   ├── connections-api.ts  # Connections API
│   │       │   │   │   │   # BE: users-be/connection
│   │       │   │   │   ├── groups-api.ts  # Groups API
│   │       │   │   │   │   # BE: users-be/user_group
│   │       │   │   │   └── referrals-api.ts  # Referrals API
│   │       │   │   │       # BE: users-be/referral
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useConnectionRequest.ts  # Send/accept/reject
│   │       │   │   │   ├── useConnections.ts  # Connections management
│   │       │   │   │   ├── useGroups.ts  # Groups management
│   │       │   │   │   ├── useNetworkRecommendations.ts  # Connection recommendations
│   │       │   │   │   └── useReferrals.ts  # Referral management
│   │       │   │   ├── queries/
│   │       │   │   │   ├── networking-mutations.ts  # Networking mutations
│   │       │   │   │   └── networking-queries.ts  # Networking queries
│   │       │   │   └── types.ts  # Networking types
│   │       │   ├── performance/
│   │       │   │   ├── api/
│   │       │   │   │   └── performance-api.ts  # Performance API
│   │       │   │   │       # BE: utility/metrics (if exists) or utility/analytics
│   │       │   │   │       # POST /v1/metrics/performance
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useErrorTracking.ts  # Error tracking
│   │       │   │   │   ├── usePerformanceMonitor.ts  # Monitor performance
│   │       │   │   │   └── useWebVitals.ts  # Web Vitals
│   │       │   │   ├── utils/
│   │       │   │   │   ├── error-reporter.ts  # Error reporter
│   │       │   │   │   ├── performance-observer.ts  # Performance observer
│   │       │   │   │   └── web-vitals-reporter.ts  # Web Vitals reporter
│   │       │   │   └── types.ts  # Performance types
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
│   │       │   │   │   ├── usePresence.ts  # User presence (online/offline)
│   │       │   │   │   │   # User presence
│   │       │   │   │   ├── useRealtimeAuction.ts  # Real-time auction updates
│   │       │   │   │   ├── useRealtimeMessages.ts  # Real-time messages
│   │       │   │   │   ├── useRealtimeNotifications.ts  # Real-time notifications
│   │       │   │   │   ├── useTypingIndicator.ts  # Typing indicator
│   │       │   │   │   └── useWebSocket.ts  # WebSocket connection
│   │       │   │   │       # WebSocket hook
│   │       │   │   ├── providers/
│   │       │   │   │   └── WebSocketProvider.tsx  # WebSocket context
│   │       │   │   ├── store/
│   │       │   │   │   └── realtime-store.ts  # Real-time state (Zustand)
│   │       │   │   ├── websocket/
│   │       │   │   │   ├── client.ts  # WebSocket client
│   │       │   │   │   ├── heartbeat.ts  # Connection health
│   │       │   │   │   └── reconnection.ts  # Reconnection logic
│   │       │   │   ├── types.ts  # Real-time types
│   │       │   │   │   # WebSocket message types
│   │       │   │   └── utils.ts  # Connection management
│   │       │   ├── safety/
│   │       │   │   ├── components/
│   │       │   │   │   ├── BlockConfirmation.tsx  # Block confirmation
│   │       │   │   │   ├── ReportModal.tsx  # Report modal
│   │       │   │   │   └── SafetyNotice.tsx  # Safety notice banner
│   │       │   │   └── utils.ts  # Safety utilities
│   │       │   ├── search/
│   │       │   │   ├── api/
│   │       │   │   │   ├── recommendations-api.ts  # Recommendations API
│   │       │   │   │   │   # BE: search-be/recommendation
│   │       │   │   │   ├── saved-searches-api.ts  # Saved searches API
│   │       │   │   │   │   # BE: search-be/saved-search
│   │       │   │   │   ├── search-api.ts  # Search API (already may exist, ensuring completeness)
│   │       │   │   │   │   # BE: search-be/query
│   │       │   │   │   └── trending-api.ts  # Trending API
│   │       │   │   │       # BE: search-be/trending
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useRecommendations.ts  # Recommendations
│   │       │   │   │   ├── useSavedSearches.ts  # Saved searches
│   │       │   │   │   ├── useSearch.ts  # Search execution
│   │       │   │   │   ├── useSearchHistory.ts  # Search history
│   │       │   │   │   ├── useSearchSuggestions.ts  # Auto-complete suggestions
│   │       │   │   │   └── useTrending.ts  # Trending items
│   │       │   │   ├── queries/
│   │       │   │   │   ├── search-mutations.ts  # Search mutations
│   │       │   │   │   └── search-queries.ts  # Search queries
│   │       │   │   ├── store/
│   │       │   │   │   └── search-store.ts  # Search UI state (filters, etc.)
│   │       │   │   └── types.ts  # Search types
│   │       │   ├── shortlists/
│   │       │   │   ├── api/
│   │       │   │   │   └── shortlists-api.ts  # Shortlists API
│   │       │   │   │       # BE: jobs-be/shortlist
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useAddToShortlist.ts  # Add candidate
│   │       │   │   │   ├── useRemoveFromShortlist.ts  # Remove candidate
│   │       │   │   │   ├── useShortlist.ts  # Single shortlist
│   │       │   │   │   └── useShortlists.ts  # Shortlists management
│   │       │   │   ├── queries/
│   │       │   │   │   ├── shortlists-mutations.ts  # Shortlist mutations
│   │       │   │   │   └── shortlists-queries.ts  # Shortlist queries
│   │       │   │   └── types.ts  # Shortlist types
│   │       │   ├── video/
│   │       │   │   ├── api/
│   │       │   │   │   └── video-api.ts  # Video call API
│   │       │   │   │       # BE: communications-be/video
│   │       │   │   │       # POST /v1/video/rooms
│   │       │   │   │       # GET /v1/video/rooms/{room_id}/token
│   │       │   │   ├── components/
│   │       │   │   │   ├── ParticipantGrid/
│   │       │   │   │   │   ├── ParticipantGrid.tsx
│   │       │   │   │   │   └── ParticipantGrid.types.ts
│   │       │   │   │   ├── VideoControls/
│   │       │   │   │   │   ├── VideoControls.tsx
│   │       │   │   │   │   └── VideoControls.types.ts
│   │       │   │   │   └── VideoRoom/
│   │       │   │   │       ├── VideoRoom.native.tsx
│   │       │   │   │       ├── VideoRoom.tsx
│   │       │   │   │       ├── VideoRoom.types.ts
│   │       │   │   │       └── VideoRoom.web.tsx
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useRecording.ts  # Call recording
│   │       │   │   │   ├── useScreenShare.ts  # Screen sharing
│   │       │   │   │   ├── useVideoCall.ts  # Video call management
│   │       │   │   │   └── useVideoDevices.ts  # Device selection
│   │       │   │   └── types.ts  # Video types
│   │       │   ├── webhooks/
│   │       │   │   ├── api/
│   │       │   │   │   └── webhooks-api.ts  # Webhooks API client
│   │       │   │   │       # BE: utility-be/webhooks OR admin-be
│   │       │   │   │       # GET /v1/webhooks
│   │       │   │   │       # POST /v1/webhooks
│   │       │   │   │       # PUT /v1/webhooks/{webhook_id}
│   │       │   │   │       # DELETE /v1/webhooks/{webhook_id}
│   │       │   │   │       # POST /v1/webhooks/{webhook_id}/test
│   │       │   │   ├── components/
│   │       │   │   │   ├── EventSelector/
│   │       │   │   │   │   ├── EventSelector.tsx
│   │       │   │   │   │   └── EventSelector.types.ts
│   │       │   │   │   ├── WebhookForm/
│   │       │   │   │   │   ├── WebhookForm.tsx
│   │       │   │   │   │   ├── WebhookForm.types.ts
│   │       │   │   │   │   └── WebhookForm.web.tsx
│   │       │   │   │   └── WebhookLogs/
│   │       │   │   │       ├── WebhookLogs.tsx
│   │       │   │   │       └── WebhookLogs.types.ts
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useWebhook.ts  # Single webhook
│   │       │   │   │   ├── useWebhookLogs.ts  # Webhook delivery logs
│   │       │   │   │   ├── useWebhooks.ts  # List webhooks
│   │       │   │   │   └── useWebhookTest.ts  # Test webhook
│   │       │   │   ├── queries/
│   │       │   │   │   ├── webhooks-mutations.ts  # Webhook mutations
│   │       │   │   │   └── webhooks-queries.ts  # Webhook queries
│   │       │   │   └── types.ts  # Webhook types
│   │       │   └── work-tracking/
│   │       │       ├── api/
│   │       │       │   ├── timesheet-api.ts  # Timesheet API
│   │       │       │   │   # BE: contracts-be/timesheet
│   │       │       │   └── work-diary-api.ts  # Work diary API
│   │       │       │       # BE: contracts-be/work_diary
│   │       │       ├── hooks/
│   │       │       │   ├── useApproveTimesheet.ts  # Approve timesheet (client)
│   │       │       │   ├── useTimesheet.ts  # Timesheet management
│   │       │       │   ├── useTimeTracking.ts  # Real-time time tracking
│   │       │       │   └── useWorkDiary.ts  # Work diary entries
│   │       │       ├── queries/
│   │       │       │   ├── timesheet-mutations.ts  # Timesheet mutations
│   │       │       │   ├── timesheet-queries.ts  # Timesheet queries
│   │       │       │   ├── work-diary-mutations.ts  # Work diary mutations
│   │       │       │   └── work-diary-queries.ts  # Work diary queries
│   │       │       ├── store/
│   │       │       │   └── time-tracker-store.ts  # Time tracker state (Zustand)
│   │       │       └── types.ts  # Work tracking types
│   │       ├── monitoring/
│   │       │   ├── analytics/
│   │       │   │   ├── providers/
│   │       │   │   │   ├── amplitude.ts  # Amplitude
│   │       │   │   │   ├── google-analytics.ts  # GA4
│   │       │   │   │   └── mixpanel.ts  # Mixpanel
│   │       │   │   ├── tracking/
│   │       │   │   │   ├── event-tracker.ts
│   │       │   │   │   ├── page-tracker.ts
│   │       │   │   │   └── user-tracker.ts
│   │       │   │   ├── analytics-config.ts  # Analytics configuration
│   │       │   │   ├── event-tracker.ts  # Event tracking wrapper
│   │       │   │   ├── google-analytics.ts  # GA4
│   │       │   │   ├── mixpanel.ts  # Mixpanel
│   │       │   │   └── segment.ts  # Segment
│   │       │   ├── error-tracking/
│   │       │   │   ├── error-boundary.tsx  # Error boundary HOC
│   │       │   │   ├── error-filters.ts  # Filter errors
│   │       │   │   ├── error-reporter.ts  # Custom error reporter
│   │       │   │   └── sentry.ts  # Sentry configuration
│   │       │   ├── logging/
│   │       │   │   ├── transports/
│   │       │   │   │   ├── console-transport.ts
│   │       │   │   │   ├── file-transport.ts  # Mobile only
│   │       │   │   │   └── remote-transport.ts
│   │       │   │   ├── log-levels.ts  # Log level configuration
│   │       │   │   └── logger.ts  # Application logger
│   │       │   ├── performance/
│   │       │   │   ├── api-monitor.ts  # API performance monitoring
│   │       │   │   ├── bundle-monitor.ts  # Bundle size monitoring
│   │       │   │   ├── metrics.ts  # Custom metrics
│   │       │   │   ├── performance-observer.ts  # Performance monitoring
│   │       │   │   │   # Performance API
│   │       │   │   └── web-vitals.ts  # Core Web Vitals
│   │       │   │       # Web Vitals monitoring
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
│   │       │   ├── validation/
│   │       │   │   ├── input-validator.ts  # Input validation
│   │       │   │   ├── sanitizer.ts  # HTML/input sanitization
│   │       │   │   ├── sql-injection-guard.ts  # SQL injection protection
│   │       │   │   └── xss-protection.ts  # XSS protection
│   │       │   ├── csrf.ts  # CSRF token management
│   │       │   ├── encryption.ts  # Client-side encryption
│   │       │   │   # - Encrypt sensitive data
│   │       │   │   # - Decrypt data
│   │       │   ├── permissions.ts  # Permission checks
│   │       │   │   # - Role-based access
│   │       │   │   # - Feature flags
│   │       │   │   # - Resource ownership
│   │       │   ├── sanitization.ts  # Input sanitization
│   │       │   │   # - HTML sanitization
│   │       │   │   # - SQL injection prevention
│   │       │   │   # - XSS prevention
│   │       │   └── validation.ts  # Input validation
│   │       │       # - Schema validation
│   │       │       # - Type checking
│   │       ├── testing/
│   │       │   ├── factories/
│   │       │   │   ├── ...  # Other factories
│   │       │   │   ├── job-factory.ts  # Job factory
│   │       │   │   ├── proposal-factory.ts  # Proposal factory
│   │       │   │   └── user-factory.ts  # User factory
│   │       │   ├── fixtures/
│   │       │   │   ├── ...  # Other fixtures
│   │       │   │   ├── auth.ts  # Auth fixtures
│   │       │   │   ├── contract.ts  # Contract fixtures
│   │       │   │   ├── job.ts  # Job fixtures
│   │       │   │   ├── jobs.ts  # Job fixtures
│   │       │   │   ├── proposal.ts  # Proposal fixtures
│   │       │   │   ├── user.ts  # User fixtures
│   │       │   │   └── users.ts  # User fixtures
│   │       │   ├── mocks/
│   │       │   │   ├── api/
│   │       │   │   │   ├── ...  # Other service mocks
│   │       │   │   │   ├── communications-mock.ts  # Mock communications API
│   │       │   │   │   ├── contracts-mock.ts  # Mock contracts API
│   │       │   │   │   ├── financial-mock.ts  # Mock financial API
│   │       │   │   │   ├── jobs-mock.ts  # Mock jobs API
│   │       │   │   │   ├── proposals-mock.ts  # Mock proposals API
│   │       │   │   │   └── users-mock.ts  # Mock users API
│   │       │   │   ├── data/
│   │       │   │   │   ├── ...  # Other mock data
│   │       │   │   │   ├── jobs.ts  # Mock job data
│   │       │   │   │   ├── proposals.ts  # Mock proposal data
│   │       │   │   │   └── users.ts  # Mock user data
│   │       │   │   ├── handlers/
│   │       │   │   │   ├── ...  # Other MSW handlers
│   │       │   │   │   ├── auth-handlers.ts  # MSW auth handlers
│   │       │   │   │   └── users-handlers.ts  # MSW users handlers
│   │       │   │   ├── auth.ts  # Auth mocks
│   │       │   │   ├── contracts.ts  # Contract mocks
│   │       │   │   ├── financial.ts  # Financial mocks
│   │       │   │   ├── jobs.ts  # Job data mocks
│   │       │   │   ├── proposals.ts  # Proposal mocks
│   │       │   │   └── users.ts  # User data mocks
│   │       │   ├── setup/
│   │       │   │   ├── jest.setup.ts  # Jest configuration
│   │       │   │   ├── msw-setup.ts  # MSW server setup
│   │       │   │   ├── msw.setup.ts  # MSW setup
│   │       │   │   ├── test-providers.tsx  # Test providers wrapper
│   │       │   │   ├── test-setup.ts  # Test environment setup
│   │       │   │   └── testing-library.setup.ts  # Testing Library setup
│   │       │   └── utils/
│   │       │       ├── accessibility-checker.ts  # A11y testing utils
│   │       │       ├── mock-api.ts  # Mock API responses
│   │       │       ├── mock-providers.tsx  # Mock context providers
│   │       │       ├── render-with-providers.tsx  # Custom render
│   │       │       ├── render.tsx  # Custom render with providers
│   │       │       ├── test-data-generators.ts  # Generate test data
│   │       │       ├── test-helpers.ts  # Helper functions
│   │       │       └── wait-for-async.ts  # Async utilities
│   │       └── utils/
│   │           └── logger/
│   │               ├── formatters/
│   │               │   ├── json.ts  # JSON formatter
│   │               │   └── pretty.ts  # Pretty formatter
│   │               ├── transports/
│   │               │   ├── console.ts  # Console transport
│   │               │   ├── file.ts  # File transport
│   │               │   └── remote.ts  # Remote logging service
│   │               ├── log-levels.ts  # Log levels
│   │               └── logger.ts  # Logger implementation
│   ├── types/
│   │   └── src/
│   │       ├── api/
│   │       │   ├── performance/
│   │       │   │   └── index.ts  # Performance types
│   │       │   │       # - Web Vitals
│   │       │   │       # - Performance metric
│   │       │   │       # - Error event
│   │       │   ├── realtime/
│   │       │   │   └── index.ts  # Real-time types
│   │       │   │       # - WebSocket message
│   │       │   │       # - Presence update
│   │       │   │       # - Event subscription
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
│           ├── auction/
│           │   ├── AuctionTimer.native.tsx
│           │   ├── AuctionTimer.tsx  # Countdown timer
│           │   ├── AuctionTimer.web.tsx
│           │   ├── BidHistoryChart.native.tsx
│           │   ├── BidHistoryChart.tsx  # Bid history visualization
│           │   ├── BidHistoryChart.web.tsx
│           │   ├── LiveBidFeed.native.tsx
│           │   ├── LiveBidFeed.tsx  # Real-time bid feed
│           │   └── LiveBidFeed.web.tsx
│           ├── charts/
│           │   ├── EarningsChart.native.tsx
│           │   ├── EarningsChart.tsx  # Earnings visualization
│           │   ├── EarningsChart.web.tsx
│           │   ├── PerformanceChart.native.tsx
│           │   ├── PerformanceChart.tsx  # Performance metrics
│           │   ├── PerformanceChart.web.tsx
│           │   ├── TrendChart.native.tsx
│           │   ├── TrendChart.tsx  # Trend visualization
│           │   └── TrendChart.web.tsx
│           ├── collaboration/
│           │   ├── CollaborationPanel.native.tsx
│           │   ├── CollaborationPanel.tsx  # Team collaboration
│           │   ├── CollaborationPanel.web.tsx
│           │   ├── GroupCard.native.tsx
│           │   ├── GroupCard.tsx  # User group card
│           │   ├── GroupCard.web.tsx
│           │   ├── MentorCard.native.tsx
│           │   ├── MentorCard.tsx  # Mentor profile card
│           │   └── MentorCard.web.tsx
│           ├── compliance/
│           │   ├── DocumentUploader.native.tsx
│           │   ├── DocumentUploader.tsx  # Compliance doc uploader
│           │   ├── DocumentUploader.web.tsx
│           │   ├── VerificationStatus.native.tsx
│           │   ├── VerificationStatus.tsx  # Verification status badge
│           │   └── VerificationStatus.web.tsx
│           ├── components/
│           │   ├── Accessibility/
│           │   │   ├── FocusTrap/
│           │   │   │   ├── FocusTrap.tsx
│           │   │   │   ├── FocusTrap.types.ts
│           │   │   │   └── FocusTrap.web.tsx
│           │   │   ├── KeyboardShortcuts/
│           │   │   │   ├── ShortcutDialog/
│           │   │   │   │   └── ShortcutDialog.tsx
│           │   │   │   ├── KeyboardShortcuts.tsx
│           │   │   │   └── KeyboardShortcuts.types.ts
│           │   │   ├── LiveRegion/
│           │   │   │   ├── LiveRegion.native.tsx
│           │   │   │   ├── LiveRegion.tsx
│           │   │   │   ├── LiveRegion.types.ts
│           │   │   │   └── LiveRegion.web.tsx
│           │   │   ├── ScreenReaderOnly/
│           │   │   │   ├── ScreenReaderOnly.tsx
│           │   │   │   └── ScreenReaderOnly.types.ts
│           │   │   └── SkipLinks/
│           │   │       ├── SkipLinks.tsx
│           │   │       ├── SkipLinks.types.ts
│           │   │       └── SkipLinks.web.tsx
│           │   ├── AI/
│           │   │   ├── AIAssistant/
│           │   │   │   ├── ChatInterface/
│           │   │   │   │   └── ChatInterface.tsx
│           │   │   │   ├── AIAssistant.native.tsx
│           │   │   │   ├── AIAssistant.tsx
│           │   │   │   ├── AIAssistant.types.ts
│           │   │   │   └── AIAssistant.web.tsx
│           │   │   ├── AutoComplete/
│           │   │   │   ├── SmartAutoComplete.native.tsx
│           │   │   │   ├── SmartAutoComplete.tsx
│           │   │   │   ├── SmartAutoComplete.types.ts
│           │   │   │   └── SmartAutoComplete.web.tsx
│           │   │   ├── ContentGenerator/
│           │   │   │   ├── TemplateSelector/
│           │   │   │   │   └── TemplateSelector.tsx
│           │   │   │   ├── ContentGenerator.tsx
│           │   │   │   └── ContentGenerator.types.ts
│           │   │   └── SmartSuggestions/
│           │   │       ├── SuggestionCard/
│           │   │       │   └── SuggestionCard.tsx
│           │   │       ├── SmartSuggestions.tsx
│           │   │       └── SmartSuggestions.types.ts
│           │   ├── Calendar/
│           │   │   ├── DatePicker/
│           │   │   │   ├── DatePicker.native.tsx
│           │   │   │   ├── DatePicker.tsx
│           │   │   │   └── DatePicker.web.tsx
│           │   │   ├── DateRangePicker/
│           │   │   │   ├── DateRangePicker.native.tsx
│           │   │   │   ├── DateRangePicker.tsx
│           │   │   │   └── DateRangePicker.web.tsx
│           │   │   ├── TimePicker/
│           │   │   │   ├── TimePicker.native.tsx
│           │   │   │   ├── TimePicker.tsx
│           │   │   │   └── TimePicker.web.tsx
│           │   │   ├── Calendar.native.tsx
│           │   │   ├── Calendar.tsx  # Full calendar
│           │   │   ├── Calendar.types.ts
│           │   │   └── Calendar.web.tsx
│           │   ├── Charts/
│           │   │   ├── AreaChart/
│           │   │   │   ├── AreaChart.native.tsx  # Native implementation
│           │   │   │   ├── AreaChart.tsx  # Base area chart
│           │   │   │   ├── AreaChart.types.ts  # Types
│           │   │   │   └── AreaChart.web.tsx  # Web implementation
│           │   │   ├── BarChart/
│           │   │   │   ├── BarChart.native.tsx
│           │   │   │   ├── BarChart.tsx
│           │   │   │   ├── BarChart.types.ts
│           │   │   │   └── BarChart.web.tsx
│           │   │   ├── FunnelChart/
│           │   │   │   ├── FunnelChart.tsx
│           │   │   │   └── FunnelChart.types.ts
│           │   │   ├── GanttChart/
│           │   │   │   ├── Task/
│           │   │   │   │   └── Task.tsx
│           │   │   │   ├── GanttChart.tsx
│           │   │   │   ├── GanttChart.types.ts
│           │   │   │   └── GanttChart.web.tsx
│           │   │   ├── HeatMap/
│           │   │   │   ├── HeatMap.tsx
│           │   │   │   ├── HeatMap.types.ts
│           │   │   │   └── HeatMap.web.tsx  # D3/Recharts
│           │   │   ├── LineChart/
│           │   │   │   ├── LineChart.native.tsx
│           │   │   │   ├── LineChart.tsx
│           │   │   │   ├── LineChart.types.ts
│           │   │   │   └── LineChart.web.tsx
│           │   │   ├── OrgChart/
│           │   │   │   ├── Node/
│           │   │   │   │   └── OrgNode.tsx
│           │   │   │   ├── OrgChart.tsx
│           │   │   │   ├── OrgChart.types.ts
│           │   │   │   └── OrgChart.web.tsx
│           │   │   ├── PieChart/
│           │   │   │   ├── PieChart.native.tsx
│           │   │   │   ├── PieChart.tsx
│           │   │   │   ├── PieChart.types.ts
│           │   │   │   └── PieChart.web.tsx
│           │   │   └── Sparkline/
│           │   │       ├── Sparkline.native.tsx
│           │   │       ├── Sparkline.tsx
│           │   │       ├── Sparkline.types.ts
│           │   │       └── Sparkline.web.tsx
│           │   ├── DataDisplay/
│           │   │   ├── Kanban/
│           │   │   │   ├── KanbanBoard.native.tsx  # Touch gestures
│           │   │   │   ├── KanbanBoard.tsx
│           │   │   │   ├── KanbanBoard.types.ts
│           │   │   │   └── KanbanBoard.web.tsx  # Drag & drop
│           │   │   ├── List/
│           │   │   │   ├── VirtualList/
│           │   │   │   │   ├── VirtualList.native.tsx  # FlashList
│           │   │   │   │   ├── VirtualList.tsx
│           │   │   │   │   └── VirtualList.web.tsx
│           │   │   │   ├── List.native.tsx
│           │   │   │   ├── List.tsx
│           │   │   │   ├── List.types.ts
│           │   │   │   └── List.web.tsx
│           │   │   ├── Table/
│           │   │   │   ├── DataTable/
│           │   │   │   │   ├── DataTable.tsx  # With sorting/filtering
│           │   │   │   │   └── DataTable.types.ts
│           │   │   │   ├── VirtualTable/
│           │   │   │   │   ├── VirtualTable.tsx  # Virtualized table
│           │   │   │   │   └── VirtualTable.types.ts
│           │   │   │   ├── Table.tsx  # Base table
│           │   │   │   ├── Table.types.ts
│           │   │   │   └── Table.web.tsx  # Full featured table
│           │   │   └── Timeline/
│           │   │       ├── Timeline.native.tsx
│           │   │       ├── Timeline.tsx
│           │   │       ├── Timeline.types.ts
│           │   │       └── Timeline.web.tsx
│           │   ├── Editor/
│           │   │   ├── CodeEditor/
│           │   │   │   ├── CodeEditor.native.tsx
│           │   │   │   ├── CodeEditor.tsx
│           │   │   │   ├── CodeEditor.types.ts
│           │   │   │   └── CodeEditor.web.tsx  # Web (e.g., Monaco/CodeMirror)
│           │   │   ├── MarkdownEditor/
│           │   │   │   ├── MarkdownEditor.native.tsx
│           │   │   │   ├── MarkdownEditor.tsx
│           │   │   │   ├── MarkdownEditor.types.ts
│           │   │   │   └── MarkdownEditor.web.tsx
│           │   │   └── RichTextEditor/
│           │   │       ├── RichTextEditor.native.tsx  # Native implementation
│           │   │       ├── RichTextEditor.tsx
│           │   │       ├── RichTextEditor.types.ts
│           │   │       └── RichTextEditor.web.tsx  # Web (e.g., TipTap/Slate)
│           │   ├── Feedback/
│           │   │   ├── Alert/
│           │   │   │   ├── Alert.native.tsx
│           │   │   │   ├── Alert.tsx
│           │   │   │   ├── Alert.types.ts
│           │   │   │   └── Alert.web.tsx
│           │   │   ├── EmptyState/
│           │   │   │   ├── EmptyState.native.tsx
│           │   │   │   ├── EmptyState.tsx
│           │   │   │   ├── EmptyState.types.ts
│           │   │   │   └── EmptyState.web.tsx
│           │   │   ├── Notification/
│           │   │   │   ├── Notification.native.tsx
│           │   │   │   ├── Notification.tsx
│           │   │   │   ├── Notification.types.ts
│           │   │   │   └── Notification.web.tsx  # Toast notifications
│           │   │   ├── Progress/
│           │   │   │   ├── ProgressBar/
│           │   │   │   │   ├── ProgressBar.native.tsx
│           │   │   │   │   ├── ProgressBar.tsx
│           │   │   │   │   └── ProgressBar.web.tsx
│           │   │   │   ├── ProgressCircle/
│           │   │   │   │   ├── ProgressCircle.native.tsx
│           │   │   │   │   ├── ProgressCircle.tsx
│           │   │   │   │   └── ProgressCircle.web.tsx
│           │   │   │   └── Progress.types.ts
│           │   │   └── Skeleton/
│           │   │       ├── Skeleton.native.tsx
│           │   │       ├── Skeleton.tsx
│           │   │       ├── Skeleton.types.ts
│           │   │       └── Skeleton.web.tsx
│           │   ├── FileUpload/
│           │   │   ├── ImageUpload/
│           │   │   │   ├── ImageCropper.tsx  # Image cropping
│           │   │   │   ├── ImageUpload.native.tsx  # Camera/gallery
│           │   │   │   ├── ImageUpload.tsx
│           │   │   │   └── ImageUpload.web.tsx
│           │   │   ├── MultiFileUpload/
│           │   │   │   ├── MultiFileUpload.native.tsx
│           │   │   │   ├── MultiFileUpload.tsx
│           │   │   │   └── MultiFileUpload.web.tsx
│           │   │   ├── FileUpload.native.tsx  # Document picker
│           │   │   ├── FileUpload.tsx  # Base file upload
│           │   │   ├── FileUpload.types.ts
│           │   │   └── FileUpload.web.tsx  # Drag & drop support
│           │   ├── Form/
│           │   │   ├── CodeEditor/
│           │   │   │   ├── LanguageSelector/
│           │   │   │   │   └── LanguageSelector.tsx
│           │   │   │   ├── CodeEditor.tsx
│           │   │   │   ├── CodeEditor.types.ts
│           │   │   │   └── CodeEditor.web.tsx  # Monaco/CodeMirror
│           │   │   ├── DateRangePicker/
│           │   │   │   ├── DateRangePicker.native.tsx
│           │   │   │   ├── DateRangePicker.tsx
│           │   │   │   ├── DateRangePicker.types.ts
│           │   │   │   └── DateRangePicker.web.tsx
│           │   │   ├── MarkdownEditor/
│           │   │   │   ├── Preview/
│           │   │   │   │   └── MarkdownPreview.tsx
│           │   │   │   ├── MarkdownEditor.tsx
│           │   │   │   ├── MarkdownEditor.types.ts
│           │   │   │   └── MarkdownEditor.web.tsx
│           │   │   ├── RichTextEditor/
│           │   │   │   ├── Toolbar/
│           │   │   │   │   ├── Toolbar.tsx
│           │   │   │   │   └── Toolbar.types.ts
│           │   │   │   ├── RichTextEditor.native.tsx  # Limited rich text
│           │   │   │   ├── RichTextEditor.tsx
│           │   │   │   ├── RichTextEditor.types.ts
│           │   │   │   └── RichTextEditor.web.tsx  # Quill/TipTap
│           │   │   └── SignaturePad/
│           │   │       ├── SignaturePad.native.tsx  # React Native Canvas
│           │   │       ├── SignaturePad.tsx
│           │   │       ├── SignaturePad.types.ts
│           │   │       └── SignaturePad.web.tsx  # Canvas-based
│           │   ├── Forms/
│           │   │   ├── Checkbox/
│           │   │   │   ├── Checkbox.native.tsx
│           │   │   │   ├── Checkbox.tsx
│           │   │   │   ├── Checkbox.types.ts
│           │   │   │   └── Checkbox.web.tsx
│           │   │   ├── FormField/
│           │   │   │   ├── FormError.tsx
│           │   │   │   ├── FormField.native.tsx
│           │   │   │   ├── FormField.tsx  # Form field wrapper
│           │   │   │   ├── FormField.types.ts
│           │   │   │   ├── FormField.web.tsx
│           │   │   │   ├── FormHelperText.tsx
│           │   │   │   └── FormLabel.tsx
│           │   │   ├── Radio/
│           │   │   │   ├── Radio.types.ts
│           │   │   │   ├── RadioGroup.native.tsx
│           │   │   │   ├── RadioGroup.tsx
│           │   │   │   └── RadioGroup.web.tsx
│           │   │   ├── Select/
│           │   │   │   ├── MultiSelect/
│           │   │   │   │   ├── MultiSelect.native.tsx
│           │   │   │   │   ├── MultiSelect.tsx
│           │   │   │   │   └── MultiSelect.web.tsx
│           │   │   │   ├── Select.native.tsx
│           │   │   │   ├── Select.tsx
│           │   │   │   ├── Select.types.ts
│           │   │   │   └── Select.web.tsx
│           │   │   ├── Slider/
│           │   │   │   ├── RangeSlider/
│           │   │   │   │   ├── RangeSlider.native.tsx
│           │   │   │   │   ├── RangeSlider.tsx
│           │   │   │   │   └── RangeSlider.web.tsx
│           │   │   │   ├── Slider.native.tsx
│           │   │   │   ├── Slider.tsx
│           │   │   │   ├── Slider.types.ts
│           │   │   │   └── Slider.web.tsx
│           │   │   └── Switch/
│           │   │       ├── Switch.native.tsx
│           │   │       ├── Switch.tsx
│           │   │       ├── Switch.types.ts
│           │   │       └── Switch.web.tsx
│           │   └── Navigation/
│           │       ├── Breadcrumb/
│           │       │   ├── Breadcrumb.tsx
│           │       │   ├── Breadcrumb.types.ts
│           │       │   └── Breadcrumb.web.tsx
│           │       ├── Pagination/
│           │       │   ├── Pagination.native.tsx
│           │       │   ├── Pagination.tsx
│           │       │   ├── Pagination.types.ts
│           │       │   └── Pagination.web.tsx
│           │       ├── Stepper/
│           │       │   ├── Stepper.native.tsx
│           │       │   ├── Stepper.tsx
│           │       │   ├── Stepper.types.ts
│           │       │   └── Stepper.web.tsx
│           │       └── Tabs/
│           │           ├── Tabs.native.tsx
│           │           ├── Tabs.tsx
│           │           ├── Tabs.types.ts
│           │           └── Tabs.web.tsx
│           ├── learning/
│           │   ├── AchievementBadge.native.tsx
│           │   ├── AchievementBadge.tsx  # Achievement badge
│           │   ├── AchievementBadge.web.tsx
│           │   ├── LearningPathCard.native.tsx
│           │   ├── LearningPathCard.tsx  # Learning path card
│           │   ├── LearningPathCard.web.tsx
│           │   ├── ProgressTracker.native.tsx
│           │   ├── ProgressTracker.tsx  # Progress visualization
│           │   └── ProgressTracker.web.tsx
│           ├── tracking/
│           │   ├── TimesheetTable.native.tsx
│           │   ├── TimesheetTable.tsx  # Timesheet grid
│           │   ├── TimesheetTable.web.tsx
│           │   ├── TimeTracker.native.tsx
│           │   ├── TimeTracker.tsx  # Time tracking widget
│           │   ├── TimeTracker.web.tsx
│           │   ├── WorkDiaryEntry.native.tsx
│           │   ├── WorkDiaryEntry.tsx  # Work diary card
│           │   └── WorkDiaryEntry.web.tsx
│           └── video/
│               ├── VideoPlayer.native.tsx
│               ├── VideoPlayer.tsx  # Video player
│               ├── VideoPlayer.web.tsx
│               ├── VideoUploader.native.tsx
│               ├── VideoUploader.tsx  # Video upload
│               └── VideoUploader.web.tsx
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
│   │   ├── sync-translations.sh  # Sync i18n
│   │   └── update-deps.sh  # Update dependencies
│   ├── test/
│   │   ├── test-all.sh  # Run all tests
│   │   ├── test-coverage.sh  # Coverage report
│   │   ├── test-e2e-mobile.sh  # E2E mobile tests
│   │   ├── test-e2e-web.sh  # E2E web tests
│   │   ├── test-integration.sh  # Integration tests
│   │   └── test-unit.sh  # Unit tests
│   ├── utils/
│   │   ├── check-ports.sh  # Check port availability
│   │   ├── doctor.sh  # Environment health check
│   │   └── setup-env.sh  # Setup environment
│   ├── analyze-bundle.ts  # Bundle size analysis
│   ├── check-bundle-limits.ts  # Enforce limits
│   └── generate-bundle-report.ts  # Generate reports
├── tests/
│   ├── e2e/
│   │   ├── fixtures/
│   │   │   ├── ...
│   │   │   ├── jobs.json
│   │   │   └── users.json
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
│   │       ├── auth/
│   │       │   ├── login.spec.ts
│   │       │   ├── password-reset.spec.ts
│   │       │   └── register.spec.ts
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
│   │   │   ├── ...
│   │   │   ├── jobs-api.test.ts
│   │   │   └── users-api.test.ts
│   │   └── features/
│   │       ├── ...
│   │       ├── auth.test.ts
│   │       ├── jobs.test.ts
│   │       └── proposals.test.ts
│   └── performance/
│       ├── lighthouse/
│       │   ├── dashboard.spec.ts
│       │   ├── home.spec.ts
│       │   └── job-detail.spec.ts
│       └── load/
│           ├── api-load.spec.ts
│           └── user-simulation.spec.ts
├── .bundlewatch.config.json  # Bundle size limits
├── .env.development  # Development environment
│   # Local development
├── .env.example  # Example environment variables
│   # Template
├── .env.local  # Local development (git-ignored)
├── .env.production  # Production environment
└── .env.staging  # Staging environment
```
