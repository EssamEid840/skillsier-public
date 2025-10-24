```
fe/
├── .env.development  # Development environment
├── .env.example  # Environment variables template
├── .env.local  # Local environment variables
├── .env.local.example  # Local overrides template
├── .env.production  # Production environment
├── .env.staging  # Staging environment
├── .eslintrc.json  # Root ESLint config
│   # Web-specific ESLint rules
├── .gitignore  # Git ignore patterns
├── .nvmrc  # Node.js version (v20.x)
├── .prettierignore  # Prettier ignore patterns
├── .prettierrc  # Prettier configuration
├── pnm-workspace.yaml  # pnpm workspace configuration, perhaps typo in base pnpm-workspace.yaml
├── pnpm-lock.yaml  # Locked dependencies
├── pnpm-workspace.yaml  # pnpm workspace configuration
├── README.md  # Root README
│   # Web app documentation
├── tsconfig.base.json  # Base TypeScript configuration
│   # Shared compiler options
├── turbo.json  # Turborepo pipeline configuration
    # Build cache, task dependencies
```<!-- PART 1/11: fe/.github -->
```
fe/
├── .github/  # GitHub workflows
│   ├── actions/
│   │   ├── build-mobile/
│   │   │   └── action.yml
│   │   ├── build-web/
│   │   │   └── action.yml
│   │   ├── cache-deps/
│   │   └── action.yml
│   │   ├── cache-dependencies/
│   │   │   └── action.yml
│   │   ├── deploy/
│   │   │   └── action.yml
│   │   ├── deploy-preview/
│   │   │   └── action.yml
│   │   ├── setup-node/
│   │   │   └── action.yml
│   │   └── setup-pnpm/
│   │       └── action.yml
│   ├── CODEOWNERS  # Code ownership
│   └── workflows/
│       ├── accessibility.yml  # Accessibility checks
│       │   # - Run axe-core
│       │   # - Check WCAG compliance
│       ├── bundle-analysis.yml  # Bundle size checks
│       ├── cd-mobile-production.yml  # Mobile production deployment
│       │   # - Build
│       │   # - Submit to App Store/Play Store
│       ├── cd-mobile-staging.yml  # Mobile staging deployment
│       │   # - Build
│       │   # - Submit to TestFlight/Internal Testing
│       ├── cd-mobile.yml  # Mobile deployment
│       │   # - EAS build
│       │   # - Submit to stores
│       ├── cd-web-production.yml  # Web production deployment
│       │   # - Build
│       │   # - Deploy to production
│       │   # - Run smoke tests
│       │   # - Notify team
│       ├── cd-web-staging.yml  # Web staging deployment
│       │   # - Build
│       │   # - Deploy to staging
│       │   # - Run smoke tests
│       ├── cd-web.yml  # Web deployment
│       │   # - Build Next.js production
│       │   # - Deploy to K8s
│       ├── ci-mobile.yml  # Mobile CI pipeline
│       │   # - Lint
│       │   # - Type check
│       │   # - Unit tests
│       │   # - Build (iOS/Android)
│       ├── ci-web.yml  # Web CI pipeline
│       │   # - Lint
│       │   # - Type check
│       │   # - Unit tests
│       │   # - Build
│       │   # - Bundle size check
│       ├── ci.yml  # Continuous Integration
│       │   # - Lint, type-check, test
│       │   # - Build verification
│       │   # - Bundle size check
│       ├── dependency-review.yml  # Dependency review
│       │   # - Check for vulnerabilities
│       │   # - License compliance
│       ├── dependency-update.yml  # Dependabot automation
│       ├── dependabot.yml  # Automated dependency updates
│       ├── deploy-mobile-production.yml
│       ├── deploy-mobile-staging.yml
│       ├── deploy-web-production.yml
│       ├── deploy-web-staging.yml
│       ├── e2e-tests.yml  # E2E tests
│       │   # - Setup test environment
│       │   # - Run Playwright/Detox tests
│       │   # - Upload test results
│       ├── lighthouse.yml  # Performance audits
│       │   # - Run Lighthouse CI
│       │   # - Compare against budgets
│       │   # - Comment on PR
│       ├── lint.yml  # Linting
│       ├── release.yml  # Release automation
│       │   # - Create changelog
│       │   # - Tag release
│       │   # - Create GitHub release
│       ├── security-scan.yml  # Security scanning
│       ├── security.yml  # Security scanning
│       │   # - Dependency audit
│       │   # - SAST scan
│       │   # - License check
│       ├── test.yml  # Test runner
│       └── type-check.yml  # TypeScript checks
```
<!-- PART 2/11: fe/.husky -->
```
fe/
├── .husky/  # Git hooks for quality gates
│   ├── pre-commit  # Runs linting, type checking
│   │   # Validates no console.log in prod code
│   │   # Ensures all tests pass
│   └── pre-push  # Runs full test suite before push
│       # Bundle size check
```
<!-- PART 3/11: fe/.vscode -->
```
fe/
├── .vscode/  # VS Code workspace configuration
│   ├── extensions.json  # Recommended extensions list
│   │   # - ESLint, Prettier, Tailwind IntelliSense
│   │   # - i18n Ally, Error Lens
│   ├── launch.json  # Debug configurations
│   │   # - Next.js debug config
│   │   # - Expo debug config
│   │   # - Jest test debug config
│   └── settings.json  # Workspace settings
│       # - Format on save enabled
│       # - Auto import organization
│       # - Tailwind CSS IntelliSense config
```
<!-- PART 4/11: fe/apps/web -->
```
fe/
├── apps/
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
│       │   └── tests/
│       │       ├── auth/
│       │       │   ├── login.spec.ts
│       │       │   ├── oauth.spec.ts
│       │       │   ├── password-reset.spec.ts
│       │       │   └── register.spec.ts
│       │       ├── contracts/
│       │       │   ├── create-contract.spec.ts
│       │       │   ├── dispute.spec.ts
│       │       │   ├── submit-deliverable.spec.ts
│       │       │   └── time-tracking.spec.ts
│       │       ├── jobs/
│       │       │   ├── apply-to-job.spec.ts
│       │       │   ├── manage-jobs.spec.ts
│       │       │   ├── post-job.spec.ts
│       │       │   └── search-jobs.spec.ts
│       │       ├── messaging/
│       │       │   ├── notifications.spec.ts
│       │       │   ├── real-time-chat.spec.ts
│       │       │   └── send-message.spec.ts
│       │       ├── payments/
│       │       │   ├── add-payment-method.spec.ts
│       │       │   ├── escrow-funding.spec.ts
│       │       │   ├── release-payment.spec.ts
│       │       │   └── withdraw.spec.ts
│       │       └── proposals/
│       │           ├── create-proposal.spec.ts
│       │           ├── edit-proposal.spec.ts
│       │           └── withdraw-proposal.spec.ts
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
│       ├── middleware.ts  # Next.js middleware
│       │   # - CSP headers
│       │   # - CORS configuration
│       │   # - Rate limiting
│       │   # - Security headers (X-Frame-Options, etc.)
│       │   # - Security headers
│       │   # - CSRF protection
│       │   # - Rate limiting (client-side)
│       │   # - Auth checks
│       ├── public/  # Static assets
│       │   ├── animations/  # Lottie/animation files
│       │   │   ├── empty-state.json
│       │   │   ├── error.json
│       │   │   ├── loading.json
│       │   │   └── success.json
│       │   ├── fonts/  # Web fonts
│       │   │   ├── inter-var-latin.woff2  # Latin subset
│       │   │   ├── inter-var.woff2  # Variable font
│       │   │   └── noto-sans-arabic.woff2  # Arabic font
│       │   ├── images/
│       │   │   ├── payment-providers/  # Payment method icons
│       │   │   │   ├── mastercard.svg
│       │   │   │   ├── paypal.svg
│       │   │   │   ├── stripe.svg
│       │   │   │   └── visa.svg
│       │   │   ├── social/  # Social login icons
│       │   │   │   ├── github.svg
│       │   │   │   ├── google.svg
│       │   │   │   └── linkedin.svg
│       │   │   ├── favicon.ico
│       │   │   ├── hero-bg-mobile.webp  # Mobile hero
│       │   │   ├── hero-bg.webp  # Hero background
│       │   │   ├── logo-dark.svg  # Dark mode logo
│       │   │   ├── logo.svg  # Skillsier logo
│       │   │   ├── placeholder-avatar.png  # Default avatar
│       │   │   └── placeholder-company.png  # Default company logo
│       │   ├── locales/  # Public locale files (deprecated - moved to packages/shared)
│       │   ├── browserconfig.xml  # Windows tile config
│       │   ├── manifest.json  # PWA manifest
│       │   ├── robots.txt  # Search engine directives
│       │   ├── sitemap-dynamic.xml  # Dynamic routes sitemap
│       │   ├── sitemap.xml  # Static sitemap
│       │   └── sw.js  # Service worker (if using PWA)
│       ├── security/
│       │   ├── auth-guard.ts
│       │   ├── csrf-protection.ts
│       │   └── rate-limiter.ts
│       ├── src/
│       │   ├── app/
│       │   │   ├── [locale]/
│       │   │   │   ├── (admin)/
│       │   │   │   │   ├── audit/
│       │   │   │   │   │   ├── [logId]/
│       │   │   │   │   │   │   └── page.tsx  # Detailed audit log entry
│       │   │   │   │   │   │       # - Event details
│       │   │   │   │   │   │       # - Related entities
│       │   │   │   │   │   │       # - IP/device info
│       │   │   │   │   │   │       # BE: utility/audit
│       │   │   │   │   │   │       # GET /v1/audit/logs/{log_id}
│       │   │   │   │   │   └── page.tsx  # Audit logs
│       │   │   │   │   │       # - Log viewer
│       │   │   │   │   │       # - Advanced filters
│       │   │   │   │   │       # - Export logs
│       │   │   │   │   │       # BE: utility/audit
│       │   │   │   │   │       # GET /v1/audit/logs
│       │   │   │   │   ├── compliance/
│       │   │   │   │   │   ├── aml-kyc/
│       │   │   │   │   │   │   ├── monitoring/
│       │   │   │   │   │   │   │   └── page.tsx  # AML monitoring dashboard
│       │   │   │   │   │   │   │       # - Suspicious activity
│       │   │   │   │   │   │   │       # - Transaction patterns
│       │   │   │   │   │   │   │       # - Risk scores
│       │   │   │   │   │   │   │       # BE: admin-be/kyc_case, financial-be/transaction
│       │   │   │   │   │   │   │       # GET /v1/kyc/monitoring/suspicious-activity
│       │   │   │   │   │   │   ├── reports/
│       │   │   │   │   │   │   │   ├── [reportId]/
│       │   │   │   │   │   │   │   │   └── page.tsx  # SAR (Suspicious Activity Report) detail
│       │   │   │   │   │   │   │   │       # BE: admin-be/kyc_case
│       │   │   │   │   │   │   │   │       # GET /v1/kyc/reports/{report_id}
│       │   │   │   │   │   │   │   └── page.tsx  # AML reports list
│       │   │   │   │   │   │   │       # - Filed reports
│       │   │   │   │   │   │   │       # - Pending reports
│       │   │   │   │   │   │   │       # BE: admin-be/kyc_case
│       │   │   │   │   │   │   │       # GET /v1/kyc/reports
│       │   │   │   │   │   │   └── risk-assessment/
│       │   │   │   │   │   │       └── page.tsx  # Risk assessment tools
│       │   │   │   │   │   │           # - User risk profiles
│       │   │   │   │   │   │           # - Country risk matrix
│       │   │   │   │   │   │           # - Enhanced due diligence
│       │   │   │   │   │   │           # BE: admin-be/kyc_case
│       │   │   │   │   │   │           # GET /v1/kyc/risk-assessment
│       │   │   │   │   │   ├── data-retention/
│       │   │   │   │   │   │   ├── audit/
│       │   │   │   │   │   │   │   └── page.tsx  # Retention audit log
│       │   │   │   │   │   │   │       # - Deletion history
│       │   │   │   │   │   │   │       # - Policy compliance
│       │   │   │   │   │   │   │       # BE: utility/audit
│       │   │   │   │   │   │   │       # GET /v1/audit/retention
│       │   │   │   │   │   │   ├── policies/
│       │   │   │   │   │   │   │   └── page.tsx  # Retention policies
│       │   │   │   │   │   │   │       # - Policy definitions
│       │   │   │   │   │   │   │       # - Data categories
│       │   │   │   │   │   │   │       # - Retention periods
│       │   │   │   │   │   │   │       # BE: admin-be/data_retention (if exists) or utility/config
│       │   │   │   │   │   │   │       # GET /v1/retention/policies
│       │   │   │   │   │   │   └── schedule/
│       │   │   │   │   │   │       └── page.tsx  # Deletion schedule
│       │   │   │   │   │   │           # - Upcoming deletions
│       │   │   │   │   │   │           # - Retention expirations
│       │   │   │   │   │   │           # BE: admin-be/data_retention
│       │   │   │   │   │   │           # GET /v1/retention/schedule
│       │   │   │   │   │   ├── document-verification/
│       │   │   │   │   │   │   ├── automated-checks/
│       │   │   │   │   │   │   │   └── page.tsx  # Automated verification rules
│       │   │   │   │   │   │   │       # - OCR settings
│       │   │   │   │   │   │   │       # - Validation rules
│       │   │   │   │   │   │   │       # - ML model performance
│       │   │   │   │   │   │   │       # BE: admin-be/business_verification
│       │   │   │   │   │   │   │       # GET /v1/verification/automation-rules
│       │   │   │   │   │   │   ├── queue/
│       │   │   │   │   │   │   │   └── page.tsx  # Document verification queue
│       │   │   │   │   │   │   │       # - Pending documents
│       │   │   │   │   │   │   │       # - Priority sorting
│       │   │   │   │   │   │   │       # - Auto-verification status
│       │   │   │   │   │   │   │       # BE: admin-be/business_verification, storage-be/asset
│       │   │   │   │   │   │   │       # GET /v1/verification/documents/queue
│       │   │   │   │   │   │   └── [documentId]/
│       │   │   │   │   │   │       └── page.tsx  # Document review interface
│       │   │   │   │   │   │           # - Document viewer
│       │   │   │   │   │   │           # - Verification checks
│       │   │   │   │   │   │           # - Approve/reject/request-more
│       │   │   │   │   │   │           # BE: admin-be/business_verification, storage-be/asset
│       │   │   │   │   │   │           # GET /v1/verification/documents/{document_id}
│       │   │   │   │   │   │           # PUT /v1/verification/documents/{document_id}/review
│       │   │   │   │   │   ├── gdpr/
│       │   │   │   │   │   │   ├── consent-management/
│       │   │   │   │   │   │   │   └── page.tsx  # Consent logs & management
│       │   │   │   │   │   │   │       # - User consent history
│       │   │   │   │   │   │   │       # - Consent versions
│       │   │   │   │   │   │   │       # - Audit trail
│       │   │   │   │   │   │   │       # BE: users-be/consent, utility/audit
│       │   │   │   │   │   │   │       # GET /v1/users/consent-logs
│       │   │   │   │   │   │   ├── deletion-requests/
│       │   │   │   │   │   │   │   ├── [requestId]/
│       │   │   │   │   │   │   │   │   └── page.tsx  # Deletion request detail
│       │   │   │   │   │   │   │   │       # - Data preview
│       │   │   │   │   │   │   │   │       # - Retention check
│       │   │   │   │   │   │   │   │       # - Process deletion
│       │   │   │   │   │   │   │   │       # BE: admin-be/privacy, users-be/user
│       │   │   │   │   │   │   │   │       # GET /v1/privacy/deletion-requests/{request_id}
│       │   │   │   │   │   │   │   │       # POST /v1/privacy/deletion-requests/{request_id}/process
│       │   │   │   │   │   │   │   └── page.tsx  # Deletion requests queue
│       │   │   │   │   │   │   │       # BE: admin-be/privacy
│       │   │   │   │   │   │   │       # GET /v1/privacy/deletion-requests
│       │   │   │   │   │   │   ├── export-requests/
│       │   │   │   │   │   │   │   ├── [requestId]/
│       │   │   │   │   │   │   │   │   └── page.tsx  # Export request detail
│       │   │   │   │   │   │   │   │       # - Request review
│       │   │   │   │   │   │   │   │       # - Generate export
│       │   │   │   │   │   │   │   │       # - Approve/deny
│       │   │   │   │   │   │   │   │       # BE: admin-be/privacy, users-be/user
│       │   │   │   │   │   │   │   │       # GET /v1/privacy/export-requests/{request_id}
│       │   │   │   │   │   │   │   │       # POST /v1/privacy/export-requests/{request_id}/process
│       │   │   │   │   │   │   │   └── page.tsx  # Export requests queue
│       │   │   │   │   │   │   │       # BE: admin-be/privacy
│       │   │   │   │   │   │   │       # GET /v1/privacy/export-requests
│       │   │   │   │   │   │   └── reports/
│       │   │   │   │   │   │       └── page.tsx  # GDPR compliance reports
│       │   │   │   │   │   │           # - Processing activities
│       │   │   │   │   │   │           # - Data inventory
│       │   │   │   │   │   │           # - Breach reports
│       │   │   │   │   │   │           # BE: admin-be/privacy, utility/audit
│       │   │   │   │   │   │           # GET /v1/privacy/reports
│       │   │   │   │   │   ├── tax/
│       │   │   │   │   │   │   ├── forms/
│       │   │   │   │   │   │   │   └── page.tsx  # Tax form generation (1099, W-8BEN)
│       │   │   │   │   │   │   │       # BE: financial-be/tax
│       │   │   │   │   │   │   │       # GET /v1/tax/forms
│       │   │   │   │   │   │   └── reports/
│       │   │   │   │   │   │       └── page.tsx  # Tax compliance reports
│       │   │   │   │   │   │           # BE: financial-be/tax
│       │   │   │   │   │   │           # GET /v1/tax/reports
│       │   │   │   │   │   └── page.tsx  # Compliance dashboard
│       │   │   │   │   │       # - Overview
│       │   │   │   │   │       # - Pending tasks
│       │   │   │   │   │       # BE: admin-be/compliance
│       │   │   │   │   │       # GET /v1/compliance/dashboard
│       │   │   │   │   ├── content-moderation/
│       │   │   │   │   │   ├── flagged/
│       │   │   │   │   │   │   ├── [flagId]/
│       │   │   │   │   │   │   │   └── page.tsx  # Flagged content review
│       │   │   │   │   │   │   │       # BE: admin-be/moderation
│       │   │   │   │   │   │   │       # GET /v1/moderation/flags/{flag_id}
│       │   │   │   │   │   │   └── page.tsx  # Flagged content queue
│       │   │   │   │   │   │       # BE: admin-be/moderation
│       │   │   │   │   │   │       # GET /v1/moderation/flags
│       │   │   │   │   │   ├── rules/
│       │   │   │   │   │   │   └── page.tsx  # Moderation rules editor
│       │   │   │   │   │   │       # BE: admin-be/moderation
│       │   │   │   │   │   │       # GET /v1/moderation/rules
│       │   │   │   │   │   └── reports/
│       │   │   │   │   │       └── page.tsx  # Moderation reports
│       │   │   │   │   │           # BE: admin-be/moderation
│       │   │   │   │   │           # GET /v1/moderation/reports
│       │   │   │   │   ├── dashboard/
│       │   │   │   │   │   └── page.tsx  # Admin dashboard
│       │   │   │   │   │       # - Key metrics
│       │   │   │   │   │       # - Quick actions
│       │   │   │   │   │       # - Recent activity
│       │   │   │   │   │       # BE: admin-be/dashboard
│       │   │   │   │   │       # GET /v1/admin/dashboard
│       │   │   │   │   ├── disputes/
│       │   │   │   │   │   ├── [disputeId]/
│       │   │   │   │   │   │   ├── evidence/
│       │   │   │   │   │   │   │   └── page.tsx  # Dispute evidence review
│       │   │   │   │   │   │   │       # BE: admin-be/dispute
│       │   │   │   │   │   │   │       # GET /v1/disputes/{dispute_id}/evidence
│       │   │   │   │   │   │   ├── resolution/
│       │   │   │   │   │   │   │   └── page.tsx  # Dispute resolution interface
│       │   │   │   │   │   │   │       # BE: admin-be/dispute
│       │   │   │   │   │   │   │       # POST /v1/disputes/{dispute_id}/resolve
│       │   │   │   │   │   │   └── page.tsx  # Dispute detail
│       │   │   │   │   │   │       # BE: admin-be/dispute
│       │   │   │   │   │   │       # GET /v1/disputes/{dispute_id}
│       │   │   │   │   │   └── page.tsx  # Disputes queue
│       │   │   │   │   │       # BE: admin-be/dispute
│       │   │   │   │   │       # GET /v1/disputes
│       │   │   │   │   ├── financial-ops/
│       │   │   │   │   │   ├── fraud/
│       │   │   │   │   │   │   └── page.tsx  # Fraud detection dashboard
│       │   │   │   │   │   │       # BE: financial-be/fraud
│       │   │   │   │   │   │       # GET /v1/fraud/detections
│       │   │   │   │   │   ├── refunds/
│       │   │   │   │   │   │   ├── [refundId]/
│       │   │   │   │   │   │   │   └── page.tsx  # Refund detail
│       │   │   │   │   │   │   │       # BE: financial-be/refund
│       │   │   │   │   │   │   │       # GET /v1/refunds/{refund_id}
│       │   │   │   │   │   │   └── page.tsx  # Refund requests queue
│       │   │   │   │   │   │       # BE: financial-be/refund
│       │   │   │   │   │   │       # GET /v1/refunds/pending
│       │   │   │   │   │   ├── reports/
│       │   │   │   │   │   │   └── page.tsx  # Financial reports
│       │   │   │   │   │   │       # BE: financial-be/report
│       │   │   │   │   │   │       # GET /v1/financial/reports
│       │   │   │   │   │   └── transactions/
│       │   │   │   │   │       └── page.tsx  # Transaction monitoring
│       │   │   │   │   │           # BE: financial-be/transaction
│       │   │   │   │   │           # GET /v1/transactions
│       │   │   │   │   ├── incidents/
│       │   │   │   │   │   ├── [incidentId]/
│       │   │   │   │   │   │   ├── edit/
│       │   │   │   │   │   │   │   └── page.tsx  # Edit incident details
│       │   │   │   │   │   │   │       # - Update status
│       │   │   │   │   │   │   │       # - Add postmortem
│       │   │   │   │   │   │   │       # - Affected services
│       │   │   │   │   │   │   │       # BE: utility/status
│       │   │   │   │   │   │   │       # PUT /v1/status/incidents/{incident_id}
│       │   │   │   │   │   │   ├── timeline/
│       │   │   │   │   │   │   │   └── page.tsx  # Incident timeline
│       │   │   │   │   │   │   │       # - Event log
│       │   │   │   │   │   │   │       # - Update history
│       │   │   │   │   │   │   │       # BE: utility/status
│       │   │   │   │   │   │   │       # GET /v1/status/incidents/{incident_id}/timeline
│       │   │   │   │   │   │   └── page.tsx  # Incident detail
│       │   │   │   │   │   │       # - Current status
│       │   │   │   │   │   │       # - Impact assessment
│       │   │   │   │   │   │       # - Resolution steps
│       │   │   │   │   │   │       # BE: utility/status
│       │   │   │   │   │   │       # GET /v1/status/incidents/{incident_id}
│       │   │   │   │   │   ├── create/
│       │   │   │   │   │   │   └── page.tsx  # Create new incident
│       │   │   │   │   │   │       # - Incident type selection
│       │   │   │   │   │   │       # - Severity level
│       │   │   │   │   │   │       # - Affected services
│       │   │   │   │   │   │       # BE: utility/status
│       │   │   │   │   │   │       # POST /v1/status/incidents
│       │   │   │   │   │   ├── history/
│       │   │   │   │   │   │   └── page.tsx  # Historical incidents
│       │   │   │   │   │   │       # - Past incidents archive
│       │   │   │   │   │   │       # - Postmortems
│       │   │   │   │   │   │       # - Lessons learned
│       │   │   │   │   │   │       # BE: utility/status
│       │   │   │   │   │   │       # GET /v1/status/incidents/history
│       │   │   │   │   │   └── page.tsx  # Active incidents dashboard
│       │   │   │   │   │       # - Current incidents
│       │   │   │   │   │       # - Quick actions
│       │   │   │   │   │       # - Status board
│       │   │   │   │   │       # BE: utility/status
│       │   │   │   │   │       # GET /v1/status/incidents?status=active
│       │   │   │   │   ├── kyc/
│       │   │   │   │   │   ├── [caseId]/
│       │   │   │   │   │   │   ├── documents/
│       │   │   │   │   │   │   │   └── page.tsx  # KYC documents review
│       │   │   │   │   │   │   │       # BE: admin-be/kyc_case, storage-be/asset
│       │   │   │   │   │   │   │       # GET /v1/kyc/cases/{case_id}/documents
│       │   │   │   │   │   │   ├── history/
│       │   │   │   │   │   │   │   └── page.tsx  # KYC case history
│       │   │   │   │   │   │   │       # BE: admin-be/kyc_case
│       │   │   │   │   │   │   │       # GET /v1/kyc/cases/{case_id}/history
│       │   │   │   │   │   │   └── page.tsx  # KYC case detail
│       │   │   │   │   │   │       # BE: admin-be/kyc_case, users-be/user
│       │   │   │   │   │   │       # GET /v1/kyc/cases/{case_id}
│       │   │   │   │   │   ├── queue/
│       │   │   │   │   │   │   └── page.tsx  # KYC verification queue
│       │   │   │   │   │   │       # BE: admin-be/kyc_case
│       │   │   │   │   │   │       # GET /v1/kyc/cases/pending
│       │   │   │   │   │   └── reports/
│       │   │   │   │   │       └── page.tsx  # KYC reports
│       │   │   │   │   │           # BE: admin-be/kyc_case
│       │   │   │   │   │           # GET /v1/kyc/reports
│       │   │   │   │   ├── maintenance/
│       │   │   │   │   │   ├── [maintenanceId]/
│       │   │   │   │   │   │   ├── edit/
│       │   │   │   │   │   │   │   └── page.tsx  # Edit maintenance window
│       │   │   │   │   │   │   │       # BE: utility/status
│       │   │   │   │   │   │   │       # PUT /v1/status/maintenance/{maintenance_id}
│       │   │   │   │   │   │   └── page.tsx  # Maintenance detail
│       │   │   │   │   │   │       # BE: utility/status
│       │   │   │   │   │   │       # GET /v1/status/maintenance/{maintenance_id}
│       │   │   │   │   │   ├── schedule/
│       │   │   │   │   │   │   └── page.tsx  # Schedule maintenance
│       │   │   │   │   │   │       # - Date/time selection
│       │   │   │   │   │   │       # - Affected services
│       │   │   │   │   │   │       # - Notification plan
│       │   │   │   │   │   │       # BE: utility/status
│       │   │   │   │   │   │       # POST /v1/status/maintenance
│       │   │   │   │   │   └── page.tsx  # Maintenance calendar
│       │   │   │   │   │       # - Upcoming maintenance
│       │   │   │   │   │       # - Impact windows
│       │   │   │   │   │       # BE: utility/status
│       │   │   │   │   │       # GET /v1/status/maintenance
│       │   │   │   │   ├── platform-config/
│       │   │   │   │   │   ├── integrations/
│       │   │   │   │   │   │   ├── [integrationId]/
│       │   │   │   │   │   │   │   ├── configure/
│       │   │   │   │   │   │   │   │   └── page.tsx  # Configure integration
│       │   │   │   │   │   │   │   │       # BE: admin-be/integrations (if exists)
│       │   │   │   │   │   │   │   │       # PUT /v1/integrations/{integration_id}/config
│       │   │   │   │   │   │   │   ├── logs/
│       │   │   │   │   │   │   │   │   └── page.tsx  # Integration logs
│       │   │   │   │   │   │   │   │       # BE: utility/audit
│       │   │   │   │   │   │   │   │       # GET /v1/integrations/{integration_id}/logs
│       │   │   │   │   │   │   │   └── page.tsx  # Integration detail
│       │   │   │   │   │   │   │       # BE: admin-be/integrations
│       │   │   │   │   │   │   │       # GET /v1/integrations/{integration_id}
│       │   │   │   │   │   │   └── page.tsx  # Integrations management
│       │   │   │   │   │   │       # BE: admin-be/integrations
│       │   │   │   │   │   │       # GET /v1/integrations
│       │   │   │   │   │   ├── limits/
│       │   │   │   │   │   │   └── page.tsx  # Platform limits configuration
│       │   │   │   │   │   │       # - Rate limits
│       │   │   │   │   │   │       # - Upload limits
│       │   │   │   │   │   │       # - API quotas
│       │   │   │   │   │   │       # - Subscription limits
│       │   │   │   │   │   │       # BE: utility/config, admin-be/platform_config (if exists)
│       │   │   │   │   │   │       # GET /v1/config/limits
│       │   │   │   │   │   │       # PUT /v1/config/limits
│       │   │   │   │   │   ├── notifications/
│       │   │   │   │   │   │   ├── settings/
│       │   │   │   │   │   │   │   └── page.tsx  # Notification settings
│       │   │   │   │   │   │   │       # - Default preferences
│       │   │   │   │   │   │   │       # - Delivery channels
│       │   │   │   │   │   │   │       # - Retry policies
│       │   │   │   │   │   │   │       # BE: communications-be/config
│       │   │   │   │   │   │   │       # GET /v1/notifications/config
│       │   │   │   │   │   │   │       # PUT /v1/notifications/config
│       │   │   │   │   │   │   └── templates/
│       │   │   │   │   │   │       ├── [templateId]/
│       │   │   │   │   │   │       │   ├── edit/
│       │   │   │   │   │   │       │   │   └── page.tsx  # Edit notification template
│       │   │   │   │   │   │       │   │       # BE: communications-be/template
│       │   │   │   │   │   │       │   │       # PUT /v1/notifications/templates/{template_id}
│       │   │   │   │   │   │       │   ├── preview/
│       │   │   │   │   │   │       │   │   └── page.tsx  # Preview template
│       │   │   │   │   │   │       │   │       # BE: communications-be/template
│       │   │   │   │   │   │       │   │       # POST /v1/notifications/templates/{template_id}/preview
│       │   │   │   │   │   │       │   └── page.tsx  # Template detail
│       │   │   │   │   │   │       │       # BE: communications-be/template
│       │   │   │   │   │   │       │       # GET /v1/notifications/templates/{template_id}
│       │   │   │   │   │   │       └── page.tsx  # Template library
│       │   │   │   │   │   │           # BE: communications-be/template
│       │   │   │   │   │   │           # GET /v1/notifications/templates
│       │   │   │   │   │   ├── pricing/
│       │   │   │   │   │   │   └── page.tsx  # Pricing configuration
│       │   │   │   │   │   │       # - Commission rates
│       │   │   │   │   │   │       # - Subscription pricing
│       │   │   │   │   │   │       # - Regional pricing
│       │   │   │   │   │   │       # BE: financial-be/pricing_config (if exists)
│       │   │   │   │   │   │       # GET /v1/config/pricing
│       │   │   │   │   │   │       # PUT /v1/config/pricing
│       │   │   │   │   │   │       # Note: Requires change_approval
│       │   │   │   │   │   ├── security/
│       │   │   │   │   │   │   └── page.tsx  # Security configuration
│       │   │   │   │   │   │       # - Auth settings
│       │   │   │   │   │   │       # - Encryption keys
│       │   │   │   │   │   │       # - Compliance levels
│       │   │   │   │   │   │       # BE: utility/config, admin-be/security_config (if exists)
│       │   │   │   │   │   │       # GET /v1/config/security
│       │   │   │   │   │   │       # PUT /v1/config/security
│       │   │   │   │   │   └── skills/
│       │   │   │   │   │       └── page.tsx  # Skills taxonomy editor
│       │   │   │   │   │           # BE: admin-be/platform_config, users-be/capabilities
│       │   │   │   │   │           # GET /v1/config/skills
│       │   │   │   │   │           # PUT /v1/config/skills
│       │   │   │   │   ├── support/
│       │   │   │   │   │   ├── [ticketId]/
│       │   │   │   │   │   │   ├── escalate/
│       │   │   │   │   │   │   │   └── page.tsx  # Escalate ticket
│       │   │   │   │   │   │   │       # BE: admin-be/support
│       │   │   │   │   │   │   │       # POST /v1/support/tickets/{ticket_id}/escalate
│       │   │   │   │   │   │   └── page.tsx  # Ticket detail
│       │   │   │   │   │   │       # BE: admin-be/support
│       │   │   │   │   │   │       # GET /v1/support/tickets/{ticket_id}
│       │   │   │   │   │   ├── analytics/
│       │   │   │   │   │   │   └── page.tsx  # Support analytics
│       │   │   │   │   │   │       # BE: admin-be/support
│       │   │   │   │   │   │       # GET /v1/support/analytics
│       │   │   │   │   │   └── queue/
│       │   │   │   │   │       └── page.tsx  # Support tickets queue
│       │   │   │   │   │           # BE: admin-be/support
│       │   │   │   │   │           # GET /v1/support/tickets
│       │   │   │   │   ├── system-health/
│       │   │   │   │   │   ├── metrics/
│       │   │   │   │   │   │   └── page.tsx  # System metrics dashboard
│       │   │   │   │   │   │       # - CPU/Memory usage
│       │   │   │   │   │   │       # - Database performance
│       │   │   │   │   │   │       # - Queue depths
│       │   │   │   │   │   │       # BE: utility/metrics (or monitoring service)
│       │   │   │   │   │   │       # GET /v1/metrics/system
│       │   │   │   │   │   ├── services/
│       │   │   │   │   │   │   └── page.tsx  # Service health overview
│       │   │   │   │   │   │       # - All microservices status
│       │   │   │   │   │   │       # - Uptime metrics
│       │   │   │   │   │   │       # - Response times
│       │   │   │   │   │   │       # BE: utility/status
│       │   │   │   │   │   │       # GET /v1/status/services
│       │   │   │   │   │   └── page.tsx  # Health dashboard
│       │   │   │   │   │       # - Overall system health
│       │   │   │   │   │       # - Critical alerts
│       │   │   │   │   │       # - Performance trends
│       │   │   │   │   │       # BE: utility/status
│       │   │   │   │   │       # GET /v1/status/health
│       │   │   │   │   ├── users/
│       │   │   │   │   │   ├── [userId]/
│       │   │   │   │   │   │   ├── activity-log/
│       │   │   │   │   │   │   │   └── page.tsx  # User activity log
│       │   │   │   │   │   │   │       # - Login history
│       │   │   │   │   │   │   │       # - Action audit
│       │   │   │   │   │   │   │       # - IP addresses
│       │   │   │   │   │   │   │       # BE: utility/audit, users-be/user
│       │   │   │   │   │   │   │       # GET /v1/users/{user_id}/activity-log
│       │   │   │   │   │   │   ├── ban/
│       │   │   │   │   │   │   │   └── page.tsx  # Ban user
│       │   │   │   │   │   │   │       # - Ban duration
│       │   │   │   │   │   │   │       # - Reason
│       │   │   │   │   │   │   │       # - Notification
│       │   │   │   │   │   │   │       # BE: admin-be/user_management
│       │   │   │   │   │   │   │       # POST /v1/users/{user_id}/ban
│       │   │   │   │   │   │   ├── impersonate/
│       │   │   │   │   │   │   │   └── page.tsx  # Impersonate user
│       │   │   │   │   │   │   │       # - Start impersonation session
│       │   │   │   │   │   │   │       # - Audit trail
│       │   │   │   │   │   │   │       # BE: admin-be/user_management
│       │   │   │   │   │   │   │       # POST /v1/users/{user_id}/impersonate
│       │   │   │   │   │   │   ├── logs/
│       │   │   │   │   │   │   │   └── page.tsx  # User activity logs
│       │   │   │   │   │   │   │       # BE: utility/audit
│       │   │   │   │   │   │   │       # GET /v1/audit/users/{user_id}
│       │   │   │   │   │   │   └── page.tsx  # User detail admin view
│       │   │   │   │   │   │       # - Edit user data
│       │   │   │   │   │   │       # - Reset password
│       │   │   │   │   │   │       # - Verify account
│       │   │   │   │   │   │       # BE: admin-be/user_management
│       │   │   │   │   │   │       # GET /v1/users/{user_id}/admin
│       │   │   │   │   │   │       # PUT /v1/users/{user_id}/admin
│       │   │   │   │   │   ├── bulk-actions/
│       │   │   │   │   │   │   └── page.tsx  # Bulk user operations
│       │   │   │   │   │   │       # - Bulk suspend
│       │   │   │   │   │   │       # - Bulk notifications
│       │   │   │   │   │   │       # - Export CSV
│       │   │   │   │   │   │       # BE: admin-be/user_management
│       │   │   │   │   │   │       # POST /v1/users/bulk-actions
│       │   │   │   │   │   ├── reports/
│       │   │   │   │   │   │   └── page.tsx  # User reports
│       │   │   │   │   │   │       # - Flagged users
│       │   │   │   │   │   │       # - Abuse reports
│       │   │   │   │   │   │       # - Resolution
│       │   │   │   │   │   │       # BE: admin-be/reporting
│       │   │   │   │   │   │       # GET /v1/users/reports
│       │   │   │   │   │   ├── search/
│       │   │   │   │   │   │   └── page.tsx  # Admin user search
│       │   │   │   │   │   │       # - Advanced filters
│       │   │   │   │   │   │       # - Activity status
│       │   │   │   │   │   │       # - Role based
│       │   │   │   │   │   │       # BE: admin-be/user_management, search-be/query
│       │   │   │   │   │   │       # POST /v1/users/admin-search
│       │   │   │   │   │   └── suspended/
│       │   │   │   │   │       └── page.tsx  # Suspended users
│       │   │   │   │   │           # BE: admin-be/user
│       │   │   │   │   │           # GET /v1/admin/users/suspended
│       │   │   │   │   └── page.tsx  # Admin landing page
│       │   │   │   │       # - Navigation to sections
│       │   │   │   │       # - Role-based access
│       │   │   │   │       # BE: admin-be/dashboard
│       │   │   │   │       # GET /v1/admin
│       │   │   │   ├── (auth)/  # Auth screens
│       │   │   │   │   ├── _layout.tsx  # Auth layout
│       │   │   │   │   ├── callback.tsx  # OAuth callback
│       │   │   │   │   │   # BE: Keycloak token exchange
│       │   │   │   │   ├── forgot-password.tsx  # Password reset
│       │   │   │   │   │   # BE: users-be/security/recovery
│       │   │   │   │   │   # POST /v1/auth/forgot-password
│       │   │   │   │   ├── login.tsx  # Login screen
│       │   │   │   │   │   # - Email/password form
│       │   │   │   │   │   # - Social login (Google, Apple)
│       │   │   │   │   │   # - Biometric login (Face ID, Touch ID)
│       │   │   │   │   │   # BE: Keycloak OAuth2
│       │   │   │   │   │   # POST /v1/auth/login
│       │   │   │   │   └── register.tsx  # Registration screen
│       │   │   │   │       # BE: users-be/user
│       │   │   │   │       # POST /v1/users/register
│       │   │   │   ├── (dashboard)/  # Dashboard routes
│       │   │   │   │   ├── agency/
│       │   │   │   │   │   ├── overview/
│       │   │   │   │   │   │   └── page.tsx  # Agency dashboard
│       │   │   │   │   │   │       # - Revenue overview
│       │   │   │   │   │   │       # - Active projects
│       │   │   │   │   │   │       # - Team utilization
│       │   │   │   │   │   │       # - Client list
│       │   │   │   │   │   │       # BE: users-be/org, financial-be/analytics
│       │   │   │   │   │   │       # GET /v1/agencies/me/overview
│       │   │   │   │   │   ├── reporting/
│       │   │   │   │   │   │   ├── clients/
│       │   │   │   │   │   │   │   └── page.tsx  # Client reports
│       │   │   │   │   │   │   │       # - Performance by client
│       │   │   │   │   │   │   │       # - Billing history
│       │   │   │   │   │   │   │       # BE: users-be/org, financial-be/analytics
│       │   │   │   │   │   │   │       # GET /v1/agencies/reports/clients
│       │   │   │   │   │   │   ├── revenue/
│       │   │   │   │   │   │   │   └── page.tsx  # Revenue reports
│       │   │   │   │   │   │   │       # - Monthly/quarterly revenue
│       │   │   │   │   │   │   │       # - Profit margins
│       │   │   │   │   │   │   │       # - Forecasts
│       │   │   │   │   │   │   │       # BE: financial-be/analytics
│       │   │   │   │   │   │   │       # GET /v1/agencies/reports/revenue
│       │   │   │   │   │   │   ├── team/
│       │   │   │   │   │   │   │   └── page.tsx  # Team performance reports
│       │   │   │   │   │   │   │       # - Productivity
│       │   │   │   │   │   │   │       # - Utilization rates
│       │   │   │   │   │   │   │       # - Skill coverage
│       │   │   │   │   │   │   │       # BE: users-be/org
│       │   │   │   │   │   │   │       # GET /v1/agencies/reports/team
│       │   │   │   │   │   │   └── page.tsx  # Reporting dashboard
│       │   │   │   │   │   │       # - Custom reports builder
│       │   │   │   │   │   │       # - Export options
│       │   │   │   │   │   │       # BE: users-be/org, financial-be/analytics
│       │   │   │   │   │   │       # GET /v1/agencies/reports
│       │   │   │   │   │   ├── sub-accounts/
│       │   │   │   │   │   │   ├── [subAccountId]/
│       │   │   │   │   │   │   │   ├── settings/
│       │   │   │   │   │   │   │   │   └── page.tsx  # Sub-account settings
│       │   │   │   │   │   │   │   │       # BE: users-be/org
│       │   │   │   │   │   │   │   │       # PUT /v1/agencies/sub-accounts/{sub_account_id}
│       │   │   │   │   │   │   │   └── page.tsx  # Sub-account detail
│       │   │   │   │   │   │   │       # - Jobs posted
│       │   │   │   │   │   │   │       # - Contracts
│       │   │   │   │   │   │   │       # - Spending
│       │   │   │   │   │   │   │       # BE: users-be/org
│       │   │   │   │   │   │   │       # GET /v1/agencies/sub-accounts/{sub_account_id}
│       │   │   │   │   │   │   ├── create/
│       │   │   │   │   │   │   │   └── page.tsx  # Create sub-account
│       │   │   │   │   │   │   │       # BE: users-be/org
│       │   │   │   │   │   │   │       # POST /v1/agencies/sub-accounts
│       │   │   │   │   │   │   └── page.tsx  # All sub-accounts
│       │   │   │   │   │   │       # - List all
│       │   │   │   │   │   │       # - Manage
│       │   │   │   │   │   │       # BE: users-be/org
│       │   │   │   │   │   │       # GET /v1/agencies/sub-accounts
│       │   │   │   │   │   ├── talent-pool/
│       │   │   │   │   │   │   ├── [poolId]/
│       │   │   │   │   │   │   │   ├── members/
│       │   │   │   │   │   │   │   │   └── page.tsx  # Pool members
│       │   │   │   │   │   │   │   │       # - Freelancer list
│       │   │   │   │   │   │   │   │       # - Add/remove
│       │   │   │   │   │   │   │   │       # BE: users-be/org, search-be/talent-pool
│       │   │   │   │   │   │   │   │       # GET /v1/talent-pools/{pool_id}/members
│       │   │   │   │   │   │   │   └── page.tsx  # Talent pool detail
│       │   │   │   │   │   │   │       # BE: search-be/talent-pool
│       │   │   │   │   │   │   │       # GET /v1/talent-pools/{pool_id}
│       │   │   │   │   │   │   ├── create/
│       │   │   │   │   │   │   │   └── page.tsx  # Create talent pool
│       │   │   │   │   │   │   │       # BE: search-be/talent-pool
│       │   │   │   │   │   │   │       # POST /v1/talent-pools
│       │   │   │   │   │   │   └── page.tsx  # All talent pools
│       │   │   │   │   │   │       # BE: search-be/talent-pool
│       │   │   │   │   │   │       # GET /v1/talent-pools
│       │   │   │   │   │   ├── team/
│       │   │   │   │   │   │   ├── [memberId]/
│       │   │   │   │   │   │   │   └── page.tsx  # Team member detail
│       │   │   │   │   │   │   │       # - Performance metrics
│       │   │   │   │   │   │   │       # - Assignment history
│       │   │   │   │   │   │   │       # BE: users-be/org
│       │   │   │   │   │   │   │       # GET /v1/agencies/team/{member_id}
│       │   │   │   │   │   │   ├── invite/
│       │   │   │   │   │   │   │   └── page.tsx  # Invite to agency
│       │   │   │   │   │   │   │       # BE: users-be/org
│       │   │   │   │   │   │   │       # POST /v1/agencies/team/invite
│       │   │   │   │   │   │   └── page.tsx  # Agency team management
│       │   │   │   │   │   │       # - Member list
│       │   │   │   │   │   │       # - Roles/permissions
│       │   │   │   │   │   │       # - Performance tracking
│       │   │   │   │   │   │       # BE: users-be/org
│       │   │   │   │   │   │       # GET /v1/agencies/team
│       │   │   │   │   │   └── white-label/
│       │   │   │   │   │       └── page.tsx  # White-label settings
│       │   │   │   │   │           # - Branding
│       │   │   │   │   │           # - Custom domain
│       │   │   │   │   │           # - Logo/colors
│       │   │   │   │   │           # BE: users-be/org
│       │   │   │   │   │           # PUT /v1/agencies/white-label
│       │   │   │   │   ├── analytics/
│       │   │   │   │   │   ├── [reportId]/
│       │   │   │   │   │   │   └── page.tsx  # Custom report detail
│       │   │   │   │   │   │       # - Data visualization
│       │   │   │   │   │   │       # - Export options
│       │   │   │   │   │   │       # BE: financial-be/analytics
│       │   │   │   │   │   │       # GET /v1/analytics/reports/{report_id}
│       │   │   │   │   │   ├── custom-reports/
│       │   │   │   │   │   │   └── page.tsx  # Custom reports builder
│       │   │   │   │   │   │       # - Drag-drop metrics
│       │   │   │   │   │   │       # - Schedule reports
│       │   │   │   │   │   │       # BE: financial-be/analytics
│       │   │   │   │   │   │       # POST /v1/analytics/custom-reports
│       │   │   │   │   │   ├── performance/
│       │   │   │   │   │   │   └── page.tsx  # Platform performance analytics
│       │   │   │   │   │   │       # - Load times
│       │   │   │   │   │   │       # - Error rates
│       │   │   │   │   │   │       # - Device/browser stats
│       │   │   │   │   │   │       # BE: utility/metrics
│       │   │   │   │   │   │       # GET /v1/analytics/performance
│       │   │   │   │   │   ├── usage/
│       │   │   │   │   │   │   └── page.tsx  # Platform usage analytics
│       │   │   │   │   │   │       # - Active users
│       │   │   │   │   │   │       # - Feature adoption
│       │   │   │   │   │   │       # - Retention rates
│       │   │   │   │   │   │       # BE: financial-be/analytics
│       │   │   │   │   │   │       # GET /v1/analytics/usage
│       │   │   │   │   │   └── user-behavior/
│       │   │   │   │   │       └── page.tsx  # User behavior analytics
│       │   │   │   │   │           # - Funnels
│       │   │   │   │   │           # - Heatmaps (integrated)
│       │   │   │   │   │           # - Session recordings
│       │   │   │   │   │           # BE: utility/metrics
│       │   │   │   │   │           # GET /v1/analytics/user-behavior
│       │   │   │   │   ├── bidding/
│       │   │   │   │   │   ├── analytics/
│       │   │   │   │   │   │   └── page.tsx  # Bidding analytics
│       │   │   │   │   │   │       # - Win rate
│       │   │   │   │   │   │       # - Average bid amount
│       │   │   │   │   │   │       # - Competition analysis
│       │   │   │   │   │   │       # BE: proposals-be/bid-strategy
│       │   │   │   │   │   │       # GET /v1/bid-strategies/analytics
│       │   │   │   │   │   ├── auctions/
│       │   │   │   │   │   │   ├── [auctionId]/
│       │   │   │   │   │   │   │   └── page.tsx  # Auction participation
│       │   │   │   │   │   │   │       # - Real-time bidding
│       │   │   │   │   │   │   │       # - Bid history
│       │   │   │   │   │   │   │       # - Competitor activity
│       │   │   │   │   │   │   │       # BE: proposals-be/auction
│       │   │   │   │   │   │   │       # GET /v1/jobs/{job_id}/auction
│       │   │   │   │   │   │   │       # POST /v1/jobs/{job_id}/auction/bid
│       │   │   │   │   │   │   │       # WebSocket: Real-time updates
│       │   │   │   │   │   │   └── page.tsx  # Active auctions list
│       │   │   │   │   │   │       # BE: proposals-be/auction
│       │   │   │   │   │   │       # GET /v1/auctions/active
│       │   │   │   │   │   └── strategies/
│       │   │   │   │   │       ├── [strategyId]/
│       │   │   │   │   │       │   ├── edit/
│       │   │   │   │   │       │   │   └── page.tsx  # Edit bid strategy
│       │   │   │   │   │       │   │       # BE: proposals-be/bid-strategy
│       │   │   │   │   │       │   │       # PUT /v1/bid-strategies/{strategy_id}
│       │   │   │   │   │       │   └── page.tsx  # View bid strategy details
│       │   │   │   │   │       │       # BE: proposals-be/bid-strategy
│       │   │   │   │   │       │       # GET /v1/bid-strategies/{strategy_id}
│       │   │   │   │   │       ├── new/
│       │   │   │   │   │       │   └── page.tsx  # Create new bid strategy
│       │   │   │   │   │       │       # BE: proposals-be/bid-strategy
│       │   │   │   │   │       │       # POST /v1/bid-strategies
│       │   │   │   │   │       └── page.tsx  # Bid strategies list
│       │   │   │   │   │           # - Auto-bid rules
│       │   │   │   │   │           # - Price ranges
│       │   │   │   │   │           # - Category targeting
│       │   │   │   │   │           # BE: proposals-be/bid-strategy
│       │   │   │   │   │           # GET /v1/bid-strategies
│       │   │   │   │   ├── community/
│       │   │   │   │   │   ├── events/
│       │   │   │   │   │   │   ├── [eventId]/
│       │   │   │   │   │   │   │   ├── register/
│       │   │   │   │   │   │   │   │   └── page.tsx  # Event registration
│       │   │   │   │   │   │   │   │       # - RSVP form
│       │   │   │   │   │   │   │   │       # - Add to calendar
│       │   │   │   │   │   │   │   │       # BE: communications-be/events
│       │   │   │   │   │   │   │   │       # POST /v1/events/{event_id}/register
│       │   │   │   │   │   │   │   └── page.tsx  # Event detail
│       │   │   │   │   │   │   │       # - Info
│       │   │   │   │   │   │   │       # - Attendees
│       │   │   │   │   │   │   │       # - Check-in (QR code)
│       │   │   │   │   │   │   │       # BE: communications-be/events
│       │   │   │   │   │   │   │       # GET /v1/events/{event_id}
│       │   │   │   │   │   │   ├── my-events/
│       │   │   │   │   │   │   │   └── page.tsx  # My registered events
│       │   │   │   │   │   │   │       # - Attending
│       │   │   │   │   │   │   │       # - Past attended
│       │   │   │   │   │   │   │       # BE: communications-be/events
│       │   │   │   │   │   │   │       # GET /v1/users/me/events
│       │   │   │   │   │   │   └── upcoming/
│       │   │   │   │   │   │       └── page.tsx  # Upcoming events
│       │   │   │   │   │   │           # BE: communications-be/events
│       │   │   │   │   │   │           # GET /v1/events?status=upcoming
│       │   │   │   │   │   ├── forums/
│       │   │   │   │   │   │   ├── [forumId]/
│       │   │   │   │   │   │   │   ├── threads/
│       │   │   │   │   │   │   │   │   └── [threadId]/
│       │   │   │   │   │   │   │   │       ├── page.tsx  # Thread view
│       │   │   │   │   │   │   │   │       │   # - Posts
│       │   │   │   │   │   │   │   │       │   # - Reply
│       │   │   │   │   │   │   │   │       │   # - Voting
│       │   │   │   │   │   │   │   │       │   # BE: communications-be/forums (if exists) OR community service
│       │   │   │   │   │   │   │   │       │   # GET /v1/forums/{forum_id}/threads/{thread_id}
│       │   │   │   │   │   │   │   │       └── reply/
│       │   │   │   │   │   │   │   │           └── page.tsx  # Reply to thread
│       │   │   │   │   │   │   │   │               # BE: communications-be/forums
│       │   │   │   │   │   │   │   │               # POST /v1/threads/{thread_id}/replies
│       │   │   │   │   │   │   │   └── page.tsx  # Forum overview
│       │   │   │   │   │   │   │       # - Thread list
│       │   │   │   │   │   │   │       # - Create thread
│       │   │   │   │   │   │   │       # BE: communications-be/forums
│       │   │   │   │   │   │   │       # GET /v1/forums/{forum_id}/threads
│       │   │   │   │   │   │   └── page.tsx  # All forums
│       │   │   │   │   │   │       # - Categories
│       │   │   │   │   │   │       # - Popular threads
│       │   │   │   │   │   │       # BE: communications-be/forums
│       │   │   │   │   │   │       # GET /v1/forums
│       │   │   │   │   │   ├── groups/
│       │   │   │   │   │   │   ├── [groupId]/
│       │   │   │   │   │   │   │   ├── discussions/
│       │   │   │   │   │   │   │   │   └── page.tsx  # Group discussions
│       │   │   │   │   │   │   │   │       # - Posts feed
│       │   │   │   │   │   │   │   │       # - Comment
│       │   │   │   │   │   │   │   │       # BE: communications-be/discussions
│       │   │   │   │   │   │   │   │       # GET /v1/groups/{group_id}/discussions
│       │   │   │   │   │   │   │   ├── events/
│       │   │   │   │   │   │   │   │   ├── [eventId]/
│       │   │   │   │   │   │   │   │   │   └── page.tsx  # Event detail
│       │   │   │   │   │   │   │   │   │       # - RSVP
│       │   │   │   │   │   │   │   │   │       # - Calendar add
│       │   │   │   │   │   │   │   │   │       # BE: communications-be/events OR community service
│       │   │   │   │   │   │   │   │   │       # GET /v1/events/{event_id}
│       │   │   │   │   │   │   │   │   └── page.tsx  # Group events
│       │   │   │   │   │   │   │   │       # - Upcoming
│       │   │   │   │   │   │   │   │       # - Past
│       │   │   │   │   │   │   │   │       # BE: communications-be/events
│       │   │   │   │   │   │   │   │       # GET /v1/groups/{group_id}/events
│       │   │   │   │   │   │   │   ├── members/
│       │   │   │   │   │   │   │   │   └── page.tsx  # Group members
│       │   │   │   │   │   │   │   │       # - Member list
│       │   │   │   │   │   │   │   │       # - Invite
│       │   │   │   │   │   │   │   │       # - Roles
│       │   │   │   │   │   │   │   │       # BE: users-be/groups (if exists) OR community service
│       │   │   │   │   │   │   │   │       # GET /v1/groups/{group_id}/members
│       │   │   │   │   │   │   │   └── page.tsx  # Group overview
│       │   │   │   │   │   │   │       # - About
│       │   │   │   │   │   │   │       # - Join/leave
│       │   │   │   │   │   │   │       # - Activity feed
│       │   │   │   │   │   │   │       # BE: users-be/groups
│       │   │   │   │   │   │   │       # GET /v1/groups/{group_id}
│       │   │   │   │   │   │   ├── discover/
│       │   │   │   │   │   │   │   └── page.tsx  # Discover groups
│       │   │   │   │   │   │   │       # - Recommended
│       │   │   │   │   │   │   │       # - By interest
│       │   │   │   │   │   │   │       # BE: users-be/groups
│       │   │   │   │   │   │   │       # GET /v1/groups/discover
│       │   │   │   │   │   │   └── my-groups/
│       │   │   │   │   │   │       └── page.tsx  # My groups
│       │   │   │   │   │   │           # - Joined groups
│       │   │   │   │   │   │           # - Manage groups
│       │   │   │   │   │   │           # BE: users-be/groups
│       │   │   │   │   │   │           # GET /v1/users/me/groups
│       │   │   │   │   │   ├── mentors/
│       │   │   │   │   │   │   ├── [mentorId]/
│       │   │   │   │   │   │   │   └── page.tsx  # Mentor profile
│       │   │   │   │   │   │   │       # - Bio
│       │   │   │   │   │   │   │       # - Availability
│       │   │   │   │   │   │   │       # - Book session
│       │   │   │   │   │   │   │       # BE: users-be/mentor
│       │   │   │   │   │   │   │       # GET /v1/mentors/{mentor_id}
│       │   │   │   │   │   │   ├── apply/
│       │   │   │   │   │   │   │   └── page.tsx  # Apply to be mentor
│       │   │   │   │   │   │   │       # BE: users-be/mentor
│       │   │   │   │   │   │   │       # POST /v1/mentors/apply
│       │   │   │   │   │   │   ├── sessions/
│       │   │   │   │   │   │   │   └── page.tsx  # Mentoring sessions
│       │   │   │   │   │   │   │       # - Schedule session
│       │   │   │   │   │   │   │       # - Upcoming/past
│       │   │   │   │   │   │   │       # BE: users-be/mentor
│       │   │   │   │   │   │   │       # GET /v1/mentors/sessions
│       │   │   │   │   │   │   └── search/
│       │   │   │   │   │   │       └── page.tsx  # Search mentors
│       │   │   │   │   │   │           # BE: users-be/mentor
│       │   │   │   │   │   │           # GET /v1/mentors/search
│       │   │   │   │   │   └── page.tsx  # Community hub
│       │   │   │   │   │       # - Overview
│       │   │   │   │   │       # - Recent activity
│       │   │   │   │   │       # BE: communications-be/community
│       │   │   │   │   │       # GET /v1/community
│       │   │   │   │   ├── connects/
│       │   │   │   │   │   ├── purchase/
│       │   │   │   │   │   │   └── page.tsx  # Purchase connects
│       │   │   │   │   │   │       # - Select package
│       │   │   │   │   │   │       # - Payment processing
│       │   │   │   │   │   │       # BE: proposals-be/connect, financial-be/payment
│       │   │   │   │   │   │       # GET /v1/connects/packages
│       │   │   │   │   │   │       # POST /v1/connects/purchase
│       │   │   │   │   │   ├── usage/
│       │   │   │   │   │   │   └── page.tsx  # Connects usage analytics
│       │   │   │   │   │   │       # - Spending patterns
│       │   │   │   │   │   │       # - Refund history
│       │   │   │   │   │   │       # - ROI tracking
│       │   │   │   │   │   │       # BE: proposals-be/connect
│       │   │   │   │   │   │       # GET /v1/connects/usage-analytics
│       │   │   │   │   │   └── page.tsx  # Connects dashboard
│       │   │   │   │   │       # - Current balance
│       │   │   │   │   │       # - Transaction history
│       │   │   │   │   │       # - Refund requests
│       │   │   │   │   │       # BE: proposals-be/connect
│       │   │   │   │   │       # GET /v1/connects
│       │   │   │   │   │       # GET /v1/connects/balance
│       │   │   │   │   ├── deliverables/
│       │   │   │   │   │   ├── [contractId]/
│       │   │   │   │   │   │   ├── [deliverableId]/
│       │   │   │   │   │   │   │   ├── review/
│       │   │   │   │   │   │   │   │   └── page.tsx  # Review deliverable (client)
│       │   │   │   │   │   │   │   │       # - Approve/reject
│       │   │   │   │   │   │   │   │       # - Request changes
│       │   │   │   │   │   │   │   │       # - Add comments
│       │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable
│       │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/review
│       │   │   │   │   │   │   │   ├── revisions/
│       │   │   │   │   │   │   │   │   ├── [revisionId]/
│       │   │   │   │   │   │   │   │   │   └── page.tsx  # Revision detail
│       │   │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable
│       │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/revisions/{revision_id}
│       │   │   │   │   │   │   │   │   └── page.tsx  # Revision history
│       │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable
│       │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/revisions
│       │   │   │   │   │   │   │   └── page.tsx  # Deliverable detail
│       │   │   │   │   │   │   │       # - Status
│       │   │   │   │   │   │   │       # - Files
│       │   │   │   │   │   │   │       # - Comments
│       │   │   │   │   │   │   │       # BE: contracts-be/deliverable
│       │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}
│       │   │   │   │   │   │   ├── new/
│       │   │   │   │   │   │   │   └── page.tsx  # Create deliverable
│       │   │   │   │   │   │   │       # BE: contracts-be/deliverable
│       │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables
│       │   │   │   │   │   │   └── page.tsx  # Contract deliverables list
│       │   │   │   │   │   │       # BE: contracts-be/deliverable
│       │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables
│       │   │   │   │   │   └── page.tsx  # All deliverables overview
│       │   │   │   │   │       # BE: contracts-be/deliverable
│       │   │   │   │   │       # GET /v1/deliverables
│       │   │   │   │   ├── home/
│       │   │   │   │   │   └── page.tsx  # Dashboard home
│       │   │   │   │   │       # - Overview
│       │   │   │   │   │       # - Quick actions
│       │   │   │   │   │       # - Recent activity
│       │   │   │   │   │       # BE: multiple services
│       │   │   │   │   │       # GET /v1/dashboard
│       │   │   │   │   ├── invitations/
│       │   │   │   │   │   ├── received/
│       │   │   │   │   │   │   ├── [inviteId]/
│       │   │   │   │   │   │   │   └── page.tsx  # Invitation details
│       │   │   │   │   │   │   │       # - Job details
│       │   │   │   │   │   │   │       # - Accept/decline
│       │   │   │   │   │   │   │       # - Proposal draft
│       │   │   │   │   │   │   │       # BE: proposals-be/invite, jobs-be/job
│       │   │   │   │   │   │   │       # GET /v1/invites/{invite_id}
│       │   │   │   │   │   │   │       # POST /v1/invites/{invite_id}/accept
│       │   │   │   │   │   │   │       # POST /v1/invites/{invite_id}/decline
│       │   │   │   │   │   │   └── page.tsx  # Received invitations list
│       │   │   │   │   │   │       # BE: proposals-be/invite
│       │   │   │   │   │   │       # GET /v1/invites/received
│       │   │   │   │   │   ├── sent/
│       │   │   │   │   │   │   ├── [inviteId]/
│       │   │   │   │   │   │   │   └── page.tsx  # Sent invitation tracking
│       │   │   │   │   │   │   │       # - Delivery status
│       │   │   │   │   │   │   │       # - Response tracking
│       │   │   │   │   │   │   │       # BE: jobs-be/invitation
│       │   │   │   │   │   │   │       # GET /v1/jobs/{job_id}/invitations/{invite_id}
│       │   │   │   │   │   │   └── page.tsx  # Sent invitations list (client)
│       │   │   │   │   │   │       # BE: jobs-be/invitation
│       │   │   │   │   │   │       # GET /v1/jobs/{job_id}/invitations
│       │   │   │   │   │   └── page.tsx  # Invitations overview
│       │   │   │   │   │       # - Pending actions
│       │   │   │   │   │       # - Response rate (client)
│       │   │   │   │   │       # - Conversion metrics
│       │   │   │   │   │       # BE: proposals-be/invite OR jobs-be/invitation (based on role)
│       │   │   │   │   ├── learning/
│       │   │   │   │   │   ├── certifications/
│       │   │   │   │   │   │   ├── [certId]/
│       │   │   │   │   │   │   │   ├── verify/
│       │   │   │   │   │   │   │   │   └── page.tsx  # Verify certificate
│       │   │   │   │   │   │   │   │       # BE: utility-be/learning
│       │   │   │   │   │   │   │   │       # GET /v1/certifications/{cert_id}/verify
│       │   │   │   │   │   │   │   └── page.tsx  # Certificate detail
│       │   │   │   │   │   │   │       # - Download PDF
│       │   │   │   │   │   │   │       # - Share link
│       │   │   │   │   │   │   │       # - Add to profile
│       │   │   │   │   │   │   │       # BE: utility-be/learning
│       │   │   │   │   │   │   │       # GET /v1/certifications/{cert_id}
│       │   │   │   │   │   │   └── page.tsx  # All certifications
│       │   │   │   │   │   │       # - Earned certificates
│       │   │   │   │   │   │       # - Available certifications
│       │   │   │   │   │   │       # BE: utility-be/learning
│       │   │   │   │   │   │       # GET /v1/users/me/certifications
│       │   │   │   │   │   ├── courses/
│       │   │   │   │   │   │   ├── [courseId]/
│       │   │   │   │   │   │   │   ├── assessments/
│       │   │   │   │   │   │   │   │   └── [assessmentId]/
│       │   │   │   │   │   │   │   │       ├── results/
│       │   │   │   │   │   │   │   │       │   └── page.tsx  # View results
│       │   │   │   │   │   │   │   │       │       # BE: utility-be/learning
│       │   │   │   │   │   │   │   │       │       # GET /v1/assessments/{assessment_id}/results
│       │   │   │   │   │   │   │   │       ├── start/
│       │   │   │   │   │   │   │   │       │   └── page.tsx  # Start assessment
│       │   │   │   │   │   │   │   │       │       # BE: utility-be/learning
│       │   │   │   │   │   │   │   │       │       # POST /v1/assessments/{assessment_id}/start
│       │   │   │   │   │   │   │   │       └── submit/
│       │   │   │   │   │   │   │   │           └── page.tsx  # Submit answers
│       │   │   │   │   │   │   │   │               # BE: utility-be/learning
│       │   │   │   │   │   │   │   │               # POST /v1/assessments/{assessment_id}/submit
│       │   │   │   │   │   │   │   ├── lessons/
│       │   │   │   │   │   │   │   │   └── [lessonId]/
│       │   │   │   │   │   │   │   │       └── page.tsx  # Lesson view
│       │   │   │   │   │   │   │   │           # - Video/content player
│       │   │   │   │   │   │   │   │           # - Notes taking
│       │   │   │   │   │   │   │   │           # - Progress tracking
│       │   │   │   │   │   │   │   │           # BE: utility-be/learning OR external LMS
│       │   │   │   │   │   │   │   │           # GET /v1/courses/{course_id}/lessons/{lesson_id}
│       │   │   │   │   │   │   │   └── page.tsx  # Course overview
│       │   │   │   │   │   │   │       # - Syllabus
│       │   │   │   │   │   │   │       # - Progress
│       │   │   │   │   │   │   │       # - Enroll button
│       │   │   │   │   │   │   │       # BE: utility-be/learning
│       │   │   │   │   │   │   │       # GET /v1/courses/{course_id}
│       │   │   │   │   │   │   │       # POST /v1/courses/{course_id}/enroll
│       │   │   │   │   │   │   ├── browse/
│       │   │   │   │   │   │   │   └── page.tsx  # Browse all courses
│       │   │   │   │   │   │   │       # - Filter by skill
│       │   │   │   │   │   │   │       # - Search
│       │   │   │   │   │   │   │       # - Recommendations
│       │   │   │   │   │   │   │       # BE: utility-be/learning
│       │   │   │   │   │   │   │       # GET /v1/courses
│       │   │   │   │   │   │   └── my-courses/
│       │   │   │   │   │   │       └── page.tsx  # Enrolled courses
│       │   │   │   │   │   │           # - In progress
│       │   │   │   │   │   │           # - Completed
│       │   │   │   │   │   │           # - Certificates
│       │   │   │   │   │   │           # BE: utility-be/learning
│       │   │   │   │   │   │           # GET /v1/users/me/courses
│       │   │   │   │   │   ├── skill-tests/
│       │   │   │   │   │   │   ├── [testId]/
│       │   │   │   │   │   │   │   ├── instructions/
│       │   │   │   │   │   │   │   │   └── page.tsx  # Test instructions
│       │   │   │   │   │   │   │   │       # BE: users-be/capabilities
│       │   │   │   │   │   │   │   │       # GET /v1/skill-tests/{test_id}
│       │   │   │   │   │   │   │   ├── results/
│       │   │   │   │   │   │   │   │   └── page.tsx  # Test results
│       │   │   │   │   │   │   │   │       # - Score
│       │   │   │   │   │   │   │   │       # - Percentile
│       │   │   │   │   │   │   │   │       # - Badge earned
│       │   │   │   │   │   │   │   │       # BE: users-be/capabilities
│       │   │   │   │   │   │   │   │       # GET /v1/skill-tests/{test_id}/results
│       │   │   │   │   │   │   │   └── take-test/
│       │   │   │   │   │   │   │       └── page.tsx  # Take the test
│       │   │   │   │   │   │   │           # - Timed test interface
│       │   │   │   │   │   │   │           # - Submit answers
│       │   │   │   │   │   │   │           # BE: users-be/capabilities
│       │   │   │   │   │   │   │           # POST /v1/skill-tests/{test_id}/submit
│       │   │   │   │   │   │   └── page.tsx  # Available skill tests
│       │   │   │   │   │   │       # - Browse by skill
│       │   │   │   │   │   │       # - Test history
│       │   │   │   │   │   │       # - Top scores
│       │   │   │   │   │   │       # BE: users-be/capabilities
│       │   │   │   │   │   │       # GET /v1/skill-tests
│       │   │   │   │   │   └── page.tsx  # Learning dashboard
│       │   │   │   │   │       # - Recommendations
│       │   │   │   │   │       # - Progress overview
│       │   │   │   │   │       # BE: utility-be/learning
│       │   │   │   │   │       # GET /v1/learning/dashboard
│       │   │   │   │   ├── network/
│       │   │   │   │   │   ├── connections/
│       │   │   │   │   │   │   ├── [connectionId]/
│       │   │   │   │   │   │   │   └── page.tsx  # Connection detail
│       │   │   │   │   │   │   │       # - Shared projects
│       │   │   │   │   │   │   │       # - Endorsements
│       │   │   │   │   │   │   │       # BE: users-be/connections
│       │   │   │   │   │   │   │       # GET /v1/connections/{connection_id}
│       │   │   │   │   │   │   └── page.tsx  # Connections list
│       │   │   │   │   │   │       # - Manage connections
│       │   │   │   │   │   │       # - Pending requests
│       │   │   │   │   │   │       # BE: users-be/connections
│       │   │   │   │   │   │       # GET /v1/connections
│       │   │   │   │   │   ├── groups/
│       │   │   │   │   │   │   └── page.tsx  # User groups
│       │   │   │   │   │   │       # - Joined groups
│       │   │   │   │   │   │       # - Create/join
│       │   │   │   │   │   │       # BE: users-be/user_group
│       │   │   │   │   │   │       # GET /v1/groups
│       │   │   │   │   │   ├── recommendations/
│       │   │   │   │   │   │   └── page.tsx  # Connection recommendations
│       │   │   │   │   │   │       # BE: users-be/connections
│       │   │   │   │   │   │       # GET /v1/connections/recommendations
│       │   │   │   │   │   ├── referrals/
│       │   │   │   │   │   │   └── page.tsx  # Referral program
│       │   │   │   │   │   │       # - Referral code
│       │   │   │   │   │   │       # - Earnings from referrals
│       │   │   │   │   │   │       # BE: users-be/referral
│       │   │   │   │   │   │       # GET /v1/referrals
│       │   │   │   │   │   └── search/
│       │   │   │   │   │       └── page.tsx  # Network search
│       │   │   │   │   │           # - Find connections
│       │   │   │   │   │           # - Send requests
│       │   │   │   │   │           # BE: users-be/connections
│       │   │   │   │   │           # GET /v1/connections/search
│       │   │   │   │   ├── programs/
│       │   │   │   │   │   ├── plus/
│       │   │   │   │   │   │   └── page.tsx  # Freelancer Plus program
│       │   │   │   │   │   │       # - Benefits
│       │   │   │   │   │   │       # - Eligibility
│       │   │   │   │   │   │       # - Application
│       │   │   │   │   │   │       # BE: users-be/programs
│       │   │   │   │   │   │       # GET /v1/programs/plus
│       │   │   │   │   │   │       # POST /v1/programs/plus/apply
│       │   │   │   │   │   ├── talent-cloud/
│       │   │   │   │   │   │   └── page.tsx  # Talent Cloud program
│       │   │   │   │   │   │       # - Premium features
│       │   │   │   │   │   │       # - Dedicated support
│       │   │   │   │   │   │       # - Application
│       │   │   │   │   │   │       # BE: users-be/programs
│       │   │   │   │   │   │       # GET /v1/programs/talent-cloud
│       │   │   │   │   │   │       # POST /v1/programs/talent-cloud/apply
│       │   │   │   │   │   └── top-rated/
│       │   │   │   │   │       └── page.tsx  # Top Rated program
│       │   │   │   │   │           # - Badge levels
│       │   │   │   │   │           # - Requirements
│       │   │   │   │   │           # - Status checker
│       │   │   │   │   │           # BE: users-be/programs
│       │   │   │   │   │           # GET /v1/programs/top-rated
│       │   │   │   │   ├── reports/
│       │   │   │   │   │   ├── custom/
│       │   │   │   │   │   │   └── page.tsx  # Custom reports builder
│       │   │   │   │   │   │       # BE: financial-be/analytics
│       │   │   │   │   │   │       # POST /v1/analytics/custom-reports
│       │   │   │   │   │   ├── earnings/
│       │   │   │   │   │   │   └── page.tsx  # Earnings reports
│       │   │   │   │   │   │       # BE: financial-be/analytics
│       │   │   │   │   │   │       # GET /v1/analytics/earnings
│       │   │   │   │   │   ├── tax/
│       │   │   │   │   │   │   └── page.tsx  # Tax reports
│       │   │   │   │   │   │       # BE: financial-be/tax
│       │   │   │   │   │   │       # GET /v1/analytics/tax-reports
│       │   │   │   │   │   └── time-tracking/
│       │   │   │   │   │       └── page.tsx  # Time tracking reports
│       │   │   │   │   │           # BE: contracts-be/work_diary
│       │   │   │   │   │           # GET /v1/analytics/time-tracking
│       │   │   │   │   ├── search/
│       │   │   │   │   │   ├── advanced/
│       │   │   │   │   │   │   └── page.tsx  # Advanced search interface
│       │   │   │   │   │   │       # - Complex filters builder
│       │   │   │   │   │   │       # - Boolean operators
│       │   │   │   │   │   │       # - Saved search management
│       │   │   │   │   │   │       # BE: search-be/query
│       │   │   │   │   │   │       # POST /v1/search/advanced
│       │   │   │   │   │   ├── history/
│       │   │   │   │   │   │   └── page.tsx  # Search history
│       │   │   │   │   │   │       # BE: search-be/query
│       │   │   │   │   │   │       # GET /v1/search/history
│       │   │   │   │   │   ├── recommendations/
│       │   │   │   │   │   │   └── page.tsx  # Personalized recommendations
│       │   │   │   │   │   │       # - AI-powered job matches
│       │   │   │   │   │   │       # - Talent suggestions
│       │   │   │   │   │   │       # BE: search-be/recommendation
│       │   │   │   │   │   │       # GET /v1/recommendations/personalized
│       │   │   │   │   │   ├── saved/
│       │   │   │   │   │   │   ├── [searchId]/
│       │   │   │   │   │   │   │   ├── edit/
│       │   │   │   │   │   │   │   │   └── page.tsx  # Edit saved search
│       │   │   │   │   │   │   │   │       # BE: search-be/saved-search
│       │   │   │   │   │   │   │   │       # PUT /v1/search/saved-searches/{search_id}
│       │   │   │   │   │   │   │   └── results/
│       │   │   │   │   │   │   │       └── page.tsx  # View results from saved search
│       │   │   │   │   │   │   │           # BE: search-be/saved-search, search-be/query
│       │   │   │   │   │   │   │           # GET /v1/search/saved-searches/{search_id}/results
│       │   │   │   │   │   │   └── page.tsx  # Saved searches list (may be in combined, ensuring here)
│       │   │   │   │   │   │       # BE: search-be/saved-search
│       │   │   │   │   │   │       # GET /v1/search/saved-searches
│       │   │   │   │   │   ├── trending/
│       │   │   │   │   │   │   └── page.tsx  # Trending searches and jobs
│       │   │   │   │   │   │       # BE: search-be/trending
│       │   │   │   │   │   │       # GET /v1/trending/jobs
│       │   │   │   │   │   │       # GET /v1/trending/skills
│       │   │   │   │   │   └── page.tsx  # Search home
│       │   │   │   │   │       # BE: search-be
│       │   │   │   │   │       # POST /v1/search
│       │   │   │   │   ├── shortlists/
│       │   │   │   │   │   ├── [shortlistId]/
│       │   │   │   │   │   │   ├── edit/
│       │   │   │   │   │   │   │   └── page.tsx  # Edit shortlist
│       │   │   │   │   │   │   │       # BE: jobs-be/shortlist
│       │   │   │   │   │   │   │       # PUT /v1/jobs/{job_id}/shortlists/{shortlist_id}
│       │   │   │   │   │   │   └── page.tsx  # Shortlist details
│       │   │   │   │   │   │       # - View candidates
│       │   │   │   │   │   │       # - Send invitations
│       │   │   │   │   │   │       # - Compare profiles
│       │   │   │   │   │   │       # BE: jobs-be/shortlist
│       │   │   │   │   │   │       # GET /v1/jobs/{job_id}/shortlists/{shortlist_id}
│       │   │   │   │   │   ├── new/
│       │   │   │   │   │   │   └── page.tsx  # Create shortlist
│       │   │   │   │   │   │       # BE: jobs-be/shortlist
│       │   │   │   │   │   │       # POST /v1/jobs/{job_id}/shortlists
│       │   │   │   │   │   └── page.tsx  # Shortlists overview
│       │   │   │   │   │       # BE: jobs-be/shortlist
│       │   │   │   │   │       # GET /v1/jobs/{job_id}/shortlists
│       │   │   │   │   ├── talent/
│       │   │   │   │   │   ├── browse/
│       │   │   │   │   │   │   └── page.tsx  # Browse talent
│       │   │   │   │   │   │       # - Search freelancers
│       │   │   │   │   │   │       # - Filters (skills, rate, location)
│       │   │   │   │   │   │       # - Save to shortlist
│       │   │   │   │   │   │       # BE: search-be/query, users-be/profile
│       │   │   │   │   │   │       # POST /v1/search/freelancers
│       │   │   │   │   │   │       # GET /v1/search/freelancers?filters=...
│       │   │   │   │   │   ├── recommendations/
│       │   │   │   │   │   │   └── page.tsx  # AI-recommended talent for jobs
│       │   │   │   │   │   │       # BE: search-be/recommendation
│       │   │   │   │   │   │       # GET /v1/recommendations/talent?job_id={job_id}
│       │   │   │   │   │   ├── saved/
│       │   │   │   │   │   │   └── page.tsx  # Saved talent profiles
│       │   │   │   │   │   │       # BE: users-be/profile
│       │   │   │   │   │   │       # GET /v1/users/me/saved-profiles
│       │   │   │   │   │   └── shortlists/
│       │   │   │   │   │       ├── [shortlistId]/
│       │   │   │   │   │       │   ├── edit/
│       │   │   │   │   │       │   │   └── page.tsx  # Edit shortlist
│       │   │   │   │   │       │   │       # BE: jobs-be/shortlist
│       │   │   │   │   │       │   │       # PUT /v1/jobs/{job_id}/shortlists/{shortlist_id}
│       │   │   │   │   │       │   └── page.tsx  # Shortlist details
│       │   │   │   │   │       │       # - View candidates
│       │   │   │   │   │       │       # - Send invitations
│       │   │   │   │   │       │       # - Compare profiles
│       │   │   │   │   │       │       # BE: jobs-be/shortlist
│       │   │   │   │   │       │       # GET /v1/jobs/{job_id}/shortlists/{shortlist_id}
│       │   │   │   │   │       ├── new/
│       │   │   │   │   │       │   └── page.tsx  # Create shortlist
│       │   │   │   │   │       │       # BE: jobs-be/shortlist
│       │   │   │   │   │       │       # POST /v1/jobs/{job_id}/shortlists
│       │   │   │   │   │       └── page.tsx  # Shortlists overview
│       │   │   │   │   │           # BE: jobs-be/shortlist
│       │   │   │   │   │           # GET /v1/jobs/{job_id}/shortlists
│       │   │   │   │   ├── timesheets/
│       │   │   │   │   │   ├── [contractId]/
│       │   │   │   │   │   │   ├── [timesheetId]/
│       │   │   │   │   │   │   │   ├── edit/
│       │   │   │   │   │   │   │   │   └── page.tsx  # Edit timesheet
│       │   │   │   │   │   │   │   │       # BE: contracts-be/timesheet
│       │   │   │   │   │   │   │   │       # PUT /v1/contracts/{contract_id}/timesheets/{timesheet_id}
│       │   │   │   │   │   │   │   └── page.tsx  # Timesheet details
│       │   │   │   │   │   │   │       # - Hours breakdown
│       │   │   │   │   │   │   │       # - Approval status
│       │   │   │   │   │   │   │       # - Dispute options
│       │   │   │   │   │   │   │       # BE: contracts-be/timesheet
│       │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/timesheets/{timesheet_id}
│       │   │   │   │   │   │   ├── new/
│       │   │   │   │   │   │   │   └── page.tsx  # Create timesheet
│       │   │   │   │   │   │   │       # BE: contracts-be/timesheet
│       │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/timesheets
│       │   │   │   │   │   │   └── page.tsx  # Contract timesheets list
│       │   │   │   │   │   │       # BE: contracts-be/timesheet
│       │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/timesheets
│       │   │   │   │   │   ├── approve/
│       │   │   │   │   │   │   └── page.tsx  # Timesheets pending approval (client)
│       │   │   │   │   │   │       # BE: contracts-be/timesheet
│       │   │   │   │   │   │       # GET /v1/timesheets/pending-approval
│       │   │   │   │   │   └── page.tsx  # All timesheets overview
│       │   │   │   │   │       # BE: contracts-be/timesheet
│       │   │   │   │   │       # GET /v1/timesheets
│       │   │   │   │   ├── work-diary/
│       │   │   │   │   │   ├── [contractId]/
│       │   │   │   │   │   │   ├── calendar/
│       │   │   │   │   │   │   │   └── page.tsx  # Calendar view of work diary
│       │   │   │   │   │   │   │       # BE: contracts-be/work_diary
│       │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/work-diary/calendar
│       │   │   │   │   │   │   ├── screenshots/
│       │   │   │   │   │   │   │   └── page.tsx  # Screenshots management
│       │   │   │   │   │   │   │       # - View all screenshots
│       │   │   │   │   │   │   │       # - Delete sensitive ones
│       │   │   │   │   │   │   │       # - Privacy settings
│       │   │   │   │   │   │   │       # BE: contracts-be/work_diary, storage-be/asset
│       │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/work-diary/screenshots
│       │   │   │   │   │   │   └── page.tsx  # Work diary detail
│       │   │   │   │   │   │       # BE: contracts-be/work_diary
│       │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/work-diary
│       │   │   │   │   │   └── page.tsx  # Work diary overview (all contracts)
│       │   │   │   │   │       # BE: contracts-be/work_diary
│       │   │   │   │   │       # GET /v1/work-diary
│       │   │   │   │   └── page.tsx  # Dashboard home
│       │   │   │   │       # - Overview
│       │   │   │   │       # - Quick actions
│       │   │   │   │       # - Recent activity
│       │   │   │   │       # BE: multiple services
│       │   │   │   │       # GET /v1/dashboard
│       │   │   │   ├── (onboarding)/  # Onboarding flow
│       │   │   │   │   ├── _layout.tsx  # Onboarding layout
│       │   │   │   │   ├── complete.tsx  # Onboarding complete
│       │   │   │   │   ├── preferences.tsx  # Preferences
│       │   │   │   │   ├── profile.tsx  # Basic profile setup
│       │   │   │   │   ├── skills.tsx  # Skills selection
│       │   │   │   │   └── welcome.tsx  # Welcome screen
│       │   │   │   ├── accessibility-statement/
│       │   │   │   │   └── page.tsx  # Accessibility statement
│       │   │   │   │       # BE: CMS
│       │   │   │   ├── blog/
│       │   │   │   │   ├── [slug]/
│       │   │   │   │   │   └── page.tsx  # Blog post
│       │   │   │   │   │       # BE: CMS
│       │   │   │   │   └── page.tsx  # Blog home
│       │   │   │   │       # BE: CMS
│       │   │   │   ├── careers/
│       │   │   │   │   └── page.tsx  # Careers page
│       │   │   │   │       # BE: CMS
│       │   │   │   ├── contact/
│       │   │   │   │   └── page.tsx  # Contact us
│       │   │   │   │       # BE: communications-be/contact
│       │   │   │   ├── developers/
│       │   │   │   │   ├── api-docs/
│       │   │   │   │   │   └── page.tsx  # API documentation
│       │   │   │   │   │       # BE: Developer portal
│       │   │   │   │   ├── plugins/
│       │   │   │   │   │   └── page.tsx  # Plugins/integrations
│       │   │   │   │   │       # BE: Developer portal
│       │   │   │   │   ├── sample-code/
│       │   │   │   │   │   └── page.tsx  # Sample code/examples
│       │   │   │   │   │       # BE: Developer portal
│       │   │   │   │   ├── sdks/
│       │   │   │   │   │   └── page.tsx  # SDKs
│       │   │   │   │   │       # BE: Developer portal
│       │   │   │   │   └── page.tsx  # Developer portal home
│       │   │   │   │       # BE: Developer portal
│       │   │   │   ├── legal/
│       │   │   │   │   ├── ccpa/
│       │   │   │   │   │   └── page.tsx  # CCPA notice
│       │   │   │   │   │       # BE: CMS
│       │   │   │   │   ├── cookie-policy/
│       │   │   │   │   │   └── page.tsx  # Cookie policy
│       │   │   │   │   │       # BE: CMS
│       │   │   │   │   ├── dmca/
│       │   │   │   │   │   └── page.tsx  # DMCA policy
│       │   │   │   │   │       # BE: CMS
│       │   │   │   │   ├── gdpr/
│       │   │   │   │   │   └── page.tsx  # GDPR statement
│       │   │   │   │   │       # BE: CMS
│       │   │   │   │   ├── privacy-policy/
│       │   │   │   │   │   └── page.tsx  # Privacy policy
│       │   │   │   │   │       # BE: CMS
│       │   │   │   │   └── terms-of-service/
│       │   │   │   │       └── page.tsx  # Terms of service
│       │   │   │   │           # BE: CMS
│       │   │   │   ├── resources/
│       │   │   │   │   ├── case-studies/
│       │   │   │   │   │   ├── [slug]/
│       │   │   │   │   │   │   └── page.tsx  # Case study detail
│       │   │   │   │   │   │       # BE: CMS
│       │   │   │   │   │   └── page.tsx  # Case studies list
│       │   │   │   │   │       # BE: CMS
│       │   │   │   │   ├── guides/
│       │   │   │   │   │   ├── [slug]/
│       │   │   │   │   │   │   └── page.tsx  # Guide detail
│       │   │   │   │   │   │       # BE: CMS
│       │   │   │   │   │   └── page.tsx  # Guides list
│       │   │   │   │   │       # BE: CMS
│       │   │   │   │   ├── tools/
│       │   │   │   │   │   └── page.tsx  # Free tools/calculators
│       │   │   │   │   │       # BE: CMS/Utility
│       │   │   │   │   ├── webinars/
│       │   │   │   │   │   ├── [slug]/
│       │   │   │   │   │   │   └── page.tsx  # Webinar detail
│       │   │   │   │   │   │       # BE: CMS
│       │   │   │   │   │   └── page.tsx  # Webinars list
│       │   │   │   │   │       # BE: CMS
│       │   │   │   │   └── whitepapers/
│       │   │   │   │       ├── [slug]/
│       │   │   │   │       │   └── page.tsx  # Whitepaper detail
│       │   │   │   │       │       # BE: CMS
│       │   │   │   │       └── page.tsx  # Whitepapers list
│       │   │   │   │           # BE: CMS
│       │   │   │   ├── solutions/
│       │   │   │   │   ├── agencies/
│       │   │   │   │   │   └── page.tsx  # Solutions for agencies
│       │   │   │   │   │       # BE: CMS
│       │   │   │   │   ├── by-industry/
│       │   │   │   │   │   └── [industry]/
│       │   │   │   │   │       └── page.tsx  # Industry-specific solutions
│       │   │   │   │   │           # BE: CMS
│       │   │   │   │   ├── by-role/
│       │   │   │   │   │   └── [role]/
│       │   │   │   │   │       └── page.tsx  # Role-specific solutions
│       │   │   │   │   │           # BE: CMS
│       │   │   │   │   └── enterprise/
│       │   │   │   │       └── page.tsx  # Enterprise solutions
│       │   │   │   │           # BE: CMS
│       │   │   │   └── status/
│       │   │   │       └── page.tsx  # Platform status page
│       │   │   │           # BE: utility/status
│       │   │   │           # GET /v1/status
│       │   │   └── page.tsx  # Root page (redirects to /[locale])
│       │   ├── components/  # Web-specific components
│       │   │   ├── admin/
│       │   │   │   ├── ComplianceDashboard.tsx
│       │   │   │   ├── IncidentTimeline.tsx
│       │   │   │   ├── UserManagementTable.tsx
│       │   │   │   └── VerificationQueue.tsx
│       │   │   ├── agency/
│       │   │   │   ├── SubAccountCard.tsx
│       │   │   │   └── TalentPoolList.tsx
│       │   │   ├── community/
│       │   │   │   ├── EventCard.tsx
│       │   │   │   ├── ForumThread.tsx
│       │   │   │   └── GroupMemberList.tsx
│       │   │   ├── learning/
│       │   │   │   ├── CourseCard.tsx
│       │   │   │   ├── LessonPlayer.tsx
│       │   │   │   └── TestInterface.tsx
│       │   │   ├── marketing/
│       │   │   │   ├── CaseStudyCard.tsx
│       │   │   │   ├── GuideCard.tsx
│       │   │   │   └── WebinarCard.tsx
│       │   │   └── programs/
│       │   │       ├── BadgeLevel.tsx
│       │   │       └── ProgramApplicationForm.tsx
│       │   └── tests/
│       │       ├── factories/
│       │       │   ├── generators.ts  # Generate test data
│       │       │   ├── mock-api.ts  # Mock API responses
│       │       │   └── test-helpers.ts  # Helper functions
│       │       ├── fixtures/
│       │       │   ├── contract.ts  # Contract fixtures
│       │       │   ├── job.ts  # Job fixtures
│       │       │   ├── proposal.ts  # Proposal fixtures
│       │       │   └── user.ts  # User fixtures
│       │       ├── mocks/
│       │       │   ├── auth.ts  # Auth mocks
│       │       │   ├── contracts.ts  # Contract mocks
│       │       │   ├── financial.ts  # Financial mocks
│       │       │   ├── jobs.ts  # Job data mocks
│       │       │   ├── proposals.ts  # Proposal mocks
│       │       │   └── users.ts  # User data mocks
│       │       └── setup/
│       │           ├── jest.setup.ts  # Jest configuration
│       │           ├── msw.setup.ts  # MSW setup
│       │           └── testing-library.setup.ts  # Testing Library setup
│       ├── .bundlewatch.config.json  # Bundle size limits
│       ├── error.tsx  # Global error boundary
│       ├── globals.css  # Global styles
│       │   # - Tailwind base, components, utilities
│       │   # - CSS custom properties for theming
│       │   # - Dark mode variables
│       ├── jest.config.js  # Root Jest configuration
│       │   # Web Jest configuration
│       ├── layout.tsx  # Root layout
│       │   # - HTML lang attribute (i18n)
│       │   # - Head meta tags
│       │   # - Body layout
│       │   # - Font loading
│       ├── loading.tsx  # Global loading state
│       ├── next.config.js  # Next.js configuration
│       │   # - i18n config
│       │   # - Image optimization
│       │   # - Bundle analyzer
│       │   # - Security headers
│       │   # - Rewrites/redirects
│       ├── not-found.tsx  # 404 page
│       ├── package.json  # Web dependencies
│       ├── playwright.config.ts
│       ├── postcss.config.js  # PostCSS configuration
│       ├── README.md  # Web app documentation
│       ├── tailwind.config.js  # Tailwind configuration
│       │   # - Design tokens
│       │   # - Custom utilities
│       │   # - Plugin configuration
│       └── tsconfig.json  # Web TypeScript config
│           # Extends tsconfig.base.json
```
<!-- PART 5/11: fe/apps/mobile -->
```
fe/
├── apps/
│   └── mobile/  # React Native/Expo application
│       ├── app/  # Expo Router file-based routing
│       │   ├── (auth)/  # Auth screens
│       │   │   ├── _layout.tsx  # Auth layout
│       │   │   ├── callback.tsx  # OAuth callback
│       │   │   │   # BE: Keycloak token exchange
│       │   │   ├── forgot-password.tsx  # Password reset
│       │   │   │   # BE: users-be/security/recovery
│       │   │   │   # POST /v1/auth/forgot-password
│       │   │   ├── login.tsx  # Login screen
│       │   │   │   # - Email/password form
│       │   │   │   # - Social login (Google, Apple)
│       │   │   │   # - Biometric login (Face ID, Touch ID)
│       │   │   │   # BE: Keycloak OAuth2
│       │   │   │   # POST /v1/auth/login
│       │   │   └── register.tsx  # Registration screen
│       │   │       # BE: users-be/user
│       │   │       # POST /v1/users/register
│       │   ├── (onboarding)/  # Onboarding flow
│       │   │   ├── _layout.tsx  # Onboarding layout
│       │   │   ├── complete.tsx  # Onboarding complete
│       │   │   ├── preferences.tsx  # Preferences
│       │   │   ├── profile.tsx  # Basic profile setup
│       │   │   ├── skills.tsx  # Skills selection
│       │   │   └── welcome.tsx  # Welcome screen
│       │   ├── (tabs)/  # Bottom tabs navigation
│       │   │   ├── _layout.tsx  # Tabs layout
│       │   │   │   # - Bottom tab navigator
│       │   │   │   # - Tab bar icons
│       │   │   │   # - Badge indicators (messages, notifications)
│       │   │   ├── index.tsx  # Home tab / Dashboard
│       │   │   │   # - Dashboard overview
│       │   │   │   # - Quick actions
│       │   │   │   # - Recent activity
│       │   │   │   # BE: Multiple services (same as web dashboard)
│       │   │   ├── jobs.tsx  # Jobs tab
│       │   │   │   # - Job listings (browse/my-jobs based on role)
│       │   │   │   # - Search jobs
│       │   │   │   # - Saved jobs
│       │   │   │   # BE: jobs-be/job
│       │   │   │   # GET /v1/jobs/browse (freelancer)
│       │   │   │   # GET /v1/jobs/my-jobs (client)
│       │   │   ├── messages.tsx  # Messages tab
│       │   │   │   # - Conversation list
│       │   │   │   # - Real-time updates
│       │   │   │   # BE: communications-be/conversations
│       │   │   │   # GET /v1/conversations
│       │   │   │   # WebSocket connection
│       │   │   ├── more/
│       │   │   │   ├── _layout.tsx  # More menu stack
│       │   │   │   ├── about.tsx  # About app
│       │   │   │   │   # BE: none (static)
│       │   │   │   ├── account.tsx  # Account settings
│       │   │   │   │   # BE: users-be/account
│       │   │   │   ├── help.tsx  # Help center
│       │   │   │   │   # BE: CMS
│       │   │   │   └── index.tsx  # More menu home
│       │   │   ├── notifications/
│       │   │   │   ├── _layout.tsx  # Notifications stack
│       │   │   │   ├── [notificationId].tsx  # Notification detail
│       │   │   │   │   # BE: communications-be/notification
│       │   │   │   ├── index.tsx  # All notifications
│       │   │   │   │   # BE: communications-be/notification
│       │   │   │   └── settings.tsx  # Notification settings
│       │   │   │       # BE: communications-be/preferences
│       │   │   ├── profile.tsx  # Profile tab
│       │   │   │   # - Current user profile preview
│       │   │   │   # - Quick links to settings
│       │   │   │   # - Stats
│       │   │   │   # BE: users-be/profile
│       │   │   │   # GET /v1/users/me
│       │   │   ├── proposals.tsx  # Proposals tab
│       │   │   │   # - My proposals (freelancer)
│       │   │   │   # - Received proposals (client, redirects to job)
│       │   │   │   # BE: proposals-be
│       │   │   │   # GET /v1/proposals/my-proposals
│       │   │   └── search/
│       │   │       ├── _layout.tsx  # Search stack navigator
│       │   │       ├── filters.tsx  # Advanced filters
│       │   │       │   # BE: search-be/facets
│       │   │       ├── index.tsx  # Search home
│       │   │       │   # BE: search-be/query
│       │   │       ├── results.tsx  # Search results
│       │   │       │   # BE: search-be/query
│       │   │       └── saved.tsx  # Saved searches
│       │   │           # BE: search-be/saved-search
│       │   ├── +not-found.tsx  # 404 screen
│       │   ├── _layout.tsx  # Root layout
│       │   │   # - Auth provider
│       │   │   # - Query client provider
│       │   │   # - Theme provider
│       │   │   # - Error boundary
│       │   ├── contracts/
│       │   │   ├── [id]/
│       │   │   │   ├── index.tsx  # Contract overview
│       │   │   │   │   # BE: contracts-be/contract
│       │   │   │   │   # GET /v1/contracts/{contract_id}
│       │   │   │   ├── messages.tsx  # Contract messages
│       │   │   │   │   # BE: communications-be
│       │   │   │   │   # GET /v1/contracts/{contract_id}/conversation
│       │   │   │   ├── milestones.tsx  # Milestones
│       │   │   │   │   # BE: contracts-be/milestone
│       │   │   │   │   # GET /v1/contracts/{contract_id}/milestones
│       │   │   │   ├── timesheet.tsx  # Timesheet
│       │   │   │   │   # BE: contracts-be/timesheet
│       │   │   │   │   # GET /v1/contracts/{contract_id}/timesheets
│       │   │   │   └── work-diary.tsx  # Work diary
│       │   │   │       # BE: contracts-be/work_diary
│       │   │   │       # GET /v1/contracts/{contract_id}/work-diary
│       │   │   └── index.tsx  # Contracts list
│       │   │       # BE: contracts-be/contract
│       │   │       # GET /v1/contracts
│       │   ├── financials/
│       │   │   ├── invoices.tsx  # Invoices
│       │   │   │   # BE: financial-be/invoice
│       │   │   │   # GET /v1/invoices
│       │   │   ├── transactions.tsx  # Transaction history
│       │   │   │   # BE: financial-be/transaction
│       │   │   │   # GET /v1/transactions
│       │   │   └── wallet.tsx  # Wallet
│       │   │       # BE: financial-be/wallet
│       │   │       # GET /v1/wallet/balance
│       │   ├── index.tsx  # App entry point
│       │   │   # - Splash screen
│       │   │   # - Initial route determination
│       │   ├── jobs/
│       │   │   ├── [id].tsx  # Job detail
│       │   │   │   # BE: jobs-be/job
│       │   │   │   # GET /v1/jobs/{job_id}
│       │   │   ├── post.tsx  # Post job (client)
│       │   │   │   # BE: jobs-be/job
│       │   │   │   # POST /v1/jobs
│       │   │   └── search.tsx  # Job search
│       │   │       # BE: search-be
│       │   │       # POST /v1/search/jobs
│       │   ├── messages/
│       │   │   └── [conversationId].tsx  # Conversation thread
│       │   │       # - Message list
│       │   │       # - Message composer
│       │   │       # - Real-time updates
│       │   │       # BE: communications-be/messages
│       │   │       # GET /v1/conversations/{conversation_id}/messages
│       │   │       # POST /v1/messages
│       │   │       # WebSocket updates
│       │   ├── notifications/
│       │   │   └── index.tsx  # Notifications list
│       │   │       # BE: communications-be/notifications
│       │   │       # GET /v1/notifications
│       │   ├── offline/
│       │   │   ├── queue.tsx  # Offline actions queue
│       │   │   │   # - Pending uploads
│       │   │   │   # - Queued messages
│       │   │   │   # - Draft proposals
│       │   │   └── sync.tsx  # Sync status
│       │   │       # - Sync progress
│       │   │       # - Conflict resolution
│       │   ├── profile/
│       │   │   ├── edit/
│       │   │   │   ├── experience.tsx  # Edit experience
│       │   │   │   ├── index.tsx  # Edit profile
│       │   │   │   ├── portfolio.tsx  # Edit portfolio
│       │   │   │   └── skills.tsx  # Edit skills
│       │   │   └── [userId].tsx  # User profile (public view)
│       │   │       # BE: users-be/profile
│       │   │       # GET /v1/users/{user_id}/profile
│       │   ├── proposals/
│       │   │   ├── [id].tsx  # Proposal detail
│       │   │   │   # BE: proposals-be
│       │   │   │   # GET /v1/proposals/{proposal_id}
│       │   │   └── submit/
│       │   │       └── [jobId].tsx  # Submit proposal
│       │   │           # BE: proposals-be
│       │   │           # POST /v1/proposals
│       │   ├── reviews/
│       │   │   ├── create/
│       │   │   │   └── [contractId].tsx  # Create review
│       │   │   │       # BE: reviews-be/reviews
│       │   │   │       # POST /v1/reviews
│       │   │   └── index.tsx  # Reviews list
│       │   │       # BE: reviews-be/reviews
│       │   │       # GET /v1/reviews
│       │   ├── scanner/
│       │   │   ├── document.tsx  # Document scanner
│       │   │   │   # - Scan compliance docs
│       │   │   # - OCR processing
│       │   │   # BE: storage-be/asset, admin-be/business_verification
│       │   │   └── qr-code.tsx  # QR code scanner
│       │   │       # - Event check-in
│       │   │       # - Profile sharing
│       │   ├── settings/
│       │   │   ├── about.tsx  # About & support
│       │   │   ├── account.tsx  # Account settings
│       │   │   ├── index.tsx  # Settings menu
│       │   │   ├── notifications.tsx  # Notification settings
│       │   │   ├── privacy.tsx  # Privacy settings
│       │   │   └── security.tsx  # Security settings
│       │   ├── subscription/
│       │   │   ├── connects.tsx  # Connects management
│       │   │   │   # BE: subscriptions-be
│       │   │   ├── index.tsx  # Subscription overview
│       │   │   ├── plans.tsx  # Available plans
│       │   │   └── upgrade.tsx  # Upgrade plan
│       │   └── widgets/
│       │       ├── quick-actions.tsx  # Quick actions widget
│       │       │   # - Quick message
│       │       │   # - Quick proposal
│       │       └── time-tracker.tsx  # Home screen time tracker widget
│       │           # BE: contracts-be/work_diary
│       ├── assets/  # Mobile assets
│       │   ├── fonts/  # Custom fonts
│       │   ├── icons/  # App icons
│       │   ├── images/  # Images
│       └── splash/  # Splash screens
│       ├── src/
│       │   ├── components/  # Mobile-specific components
│       │   │   ├── Auth/
│       │   │   │   ├── BiometricButton.tsx  # Biometric auth button
│       │   │   │   ├── LoginForm.tsx  # Login form component
│       │   │   │   ├── RegisterForm.tsx  # Registration form
│       │   │   │   └── SocialButtons.tsx  # Social login buttons
│       │   │   ├── Common/
│       │   │   │   ├── EmptyState.tsx  # Empty state component
│       │   │   │   ├── ErrorBoundary.tsx  # Error boundary
│       │   │   │   ├── Loading.tsx  # Loading spinner
│       │   │   │   ├── OptimizedFlashList.tsx  # Optimized list (FlashList)
│       │   │   │   └── PullToRefresh.tsx  # Pull to refresh
│       │   │   ├── Contracts/
│       │   │   │   ├── ContractCard.tsx  # Contract card
│       │   │   │   ├── MilestoneItem.tsx  # Milestone list item
│       │   │   │   └── TimesheetEntry.tsx  # Timesheet entry
│       │   │   ├── Financial/
│       │   │   │   ├── InvoiceCard.tsx  # Invoice card
│       │   │   │   ├── TransactionItem.tsx  # Transaction list item
│       │   │   │   └── WalletCard.tsx  # Wallet balance card
│       │   │   ├── Jobs/
│       │   │   │   ├── JobCard.tsx  # Job card component
│       │   │   │   ├── JobDetail.tsx  # Job detail view
│       │   │   │   ├── JobFilters.tsx  # Job filters bottom sheet
│       │   │   │   └── JobList.tsx  # Job list
│       │   │   ├── Messages/
│       │   │   │   ├── ConversationCard.tsx  # Conversation list item
│       │   │   │   ├── MessageBubble.tsx  # Message bubble
│       │   │   │   ├── MessageComposer.tsx  # Message input
│       │   │   │   └── TypingIndicator.tsx  # Typing indicator
│       │   │   ├── Navigation/
│       │   │   │   ├── Header.tsx  # Screen header
│       │   │   │   └── TabBar.tsx  # Custom tab bar
│       │   │   ├── Profile/
│       │   │   │   ├── ExperienceItem.tsx  # Experience item
│       │   │   │   ├── PortfolioItem.tsx  # Portfolio item
│       │   │   │   ├── ProfileHeader.tsx  # Profile header
│       │   │   │   └── SkillTag.tsx  # Skill tag
│       │   │   └── Proposals/
│       │   │       ├── ProposalCard.tsx  # Proposal card...(truncated 149275 characters)...useSavedSearches.ts
│       │   │   │   │   │   └── useSearchSuggestions.ts
│       │   │   │   │   ├── queries/
│       │   │   │   │   │   ├── search-mutations.ts
│       │   │   │   │   │   └── search-queries.ts  # BE: search-be
│       │   │   │   │   └── types.ts
│       │   │   │   ├── storage/  # File storage feature
│       │   │   │   │   ├── api/
│       │   │   │   │   │   └── storage-api.ts  # BE: storage-be
│       │   │   │   │   │       # POST /v1/storage/upload
│       │   │   │   │   │       # GET /v1/storage/presign
│       │   │   │   │   ├── hooks/
│       │   │   │   │   │   ├── useFileDownload.ts
│       │   │   │   │   │   ├── usePresignedUrl.ts
│       │   │   │   │   │   └── useUpload.ts
│       │   │   │   │   └── types.ts
│       │   │   │   └── subscriptions/  # Subscriptions feature
│       │   │   │       ├── api/
│       │   │   │       │   └── subscriptions-api.ts
│       │   │   │       ├── hooks/
│       │   │   │       │   ├── useConnects.ts
│       │   │   │       │   ├── useEntitlements.ts
│       │   │   │       │   ├── usePlans.ts
│       │   │   │       │   ├── useSubscription.ts
│       │   │   │       │   └── useUpgrade.ts
│       │   │   │       ├── queries/
│       │   │   │       │   ├── subscription-mutations.ts
│       │   │   │       │   └── subscription-queries.ts  # BE: subscriptions-be
│       │   │   │       └── types.ts
│       │   │   ├── hooks/  # General hooks
│       │   │   │   ├── useClickOutside.ts
│       │   │   │   ├── useCopyToClipboard.ts
│       │   │   │   ├── useDebounce.ts
│       │   │   │   ├── useIntersectionObserver.ts
│       │   │   │   ├── useLocalStorage.ts
│       │   │   │   ├── useMediaQuery.ts
│       │   │   │   └── useToggle.ts
│       │   │   ├── i18n/  # Internationalization
│       │   │   │   ├── locales/  # Locale files
│       │   │   │   │   ├── ar/  # Arabic
│       │   │   │   │   ├── de/  # German
│       │   │   │   │   ├── en/
│       │   │   │   │   │   ├── admin.json
│       │   │   │   │   │   ├── auth.json
│       │   │   │   │   │   ├── common.json
│       │   │   │   │   │   ├── contracts.json
│       │   │   │   │   │   ├── errors.json
│       │   │   │   │   │   ├── financial.json
│       │   │   │   │   │   ├── jobs.json
│       │   │   │   │   │   ├── messages.json
│       │   │   │   │   │   ├── profile.json
│       │   │   │   │   │   ├── proposals.json
│       │   │   │   │   │   ├── reviews.json
│       │   │   │   │   │   ├── settings.json
│       │   │   │   │   │   └── subscription.json
│       │   │   │   │   ├── es/  # Spanish
│       │   │   │   │   ├── fr/  # French
│       │   │   │   │   ├── hi/  # Hindi
│       │   │   │   │   ├── ru/  # Russian
│       │   │   │   │   ├── tr/  # Turkish
│       │   │   │   │   └── zh/  # Chinese
│       │   │   │   └── config.ts  # i18n configuration
│       │   │   └── lib/  # Shared utilities
│       │   │       ├── api/
│       │   │       │   ├── client.ts  # Axios/Fetch client setup
│       │   │       │   │   # - Base URL
│       │   │       │   │   # - Auth interceptors
│       │   │       │   │   # - Error handling
│       │   │       │   │   # - Request/response logging
│       │   │       │   ├── endpoints.ts  # API endpoint constants
│       │   │       │   └── error-handler.ts  # Global error handler
│       │   │       ├── constants/
│       │   │       │   ├── api.ts  # API constants
│       │   │       │   ├── app.ts  # App constants
│       │   │       │   ├── permissions.ts  # RBAC permissions
│       │   │       │   └── routes.ts  # Route constants
│       │   │       ├── formatting/
│       │   │       │   ├── currency.ts  # Currency formatting
│       │   │       │   ├── date.ts  # Date formatting
│       │   │       │   ├── number.ts  # Number formatting
│       │   │       │   └── string.ts  # String utilities
│       │   │       ├── validation/
│       │   │       │   ├── schemas.ts  # Zod schemas
│       │   │       │   └── validators.ts  # Custom validators
│       │   │       └── websocket/
│       │   │           ├── client.ts  # WebSocket client
│       │   │           │   # WS: ws://communications-be/v1/realtime
│       │   │           │   # - Connection management
│       │   │           │   # - Reconnection logic
│       │   │           │   # - Event subscriptions
│       │   │           └── events.ts  # WebSocket event types
│       ├── package.json
│       ├── README.md
│       └── tsconfig.json
```
<!-- PART 6/11: fe/packages -->
```
fe/
├── packages/
├── types/  # TypeScript type definitions
│   # Shared TypeScript types
│   ├── src/
│   │   ├── api/  # API response types
│   │   │   ├── admin.ts
│   │   │   ├── common.ts  # Common types (pagination, filters, etc.)
│   │   │   ├── contracts.ts
│   │   │   ├── financial.ts
│   │   │   ├── jobs.ts
│   │   │   ├── messages.ts
│   │   │   ├── notifications.ts
│   │   │   ├── proposals.ts
│   │   │   ├── reviews.ts
│   │   │   ├── search.ts
│   │   │   ├── storage.ts
│   │   │   ├── subscriptions.ts
│   │   │   └── users.ts
│   │   ├── entities/  # Domain entities
│   │   │   ├── contract.ts
│   │   │   ├── invoice.ts
│   │   │   ├── job.ts
│   │   │   ├── message.ts
│   │   │   ├── notification.ts
│   │   │   ├── proposal.ts
│   │   │   ├── review.ts
│   │   │   ├── subscription.ts
│   │   │   ├── transaction.ts
│   │   │   └── user.ts
│   │   ├── enums/  # Enums
│   │   │   ├── contract-status.ts
│   │   │   ├── job-status.ts
│   │   │   ├── payment-status.ts
│   │   │   ├── proposal-status.ts
│   │   │   ├── review-rating.ts
│   │   │   └── user-type.ts
│   │   └── index.ts  # Export all types
│   ├── package.json
│   ├── README.md
│   └── tsconfig.json
├── ui/  # Cross-platform component library
│   # Cross-platform UI component library
│   ├── src/
│   │   ├── a11y/
│   │   │   ├── FocusTrap/
│   │   │   │   ├── FocusTrap.tsx
│   │   │   │   ├── FocusTrap.web.tsx
│   │   │   │   └── FocusTrap.types.ts
│   │   │   ├── SkipLink/
│   │   │   │   ├── SkipLink.tsx
│   │   │   │   ├── SkipLink.web.tsx
│   │   │   │   └── SkipLink.types.ts
│   │   │   ├── VisuallyHidden/
│   │   │   │   ├── VisuallyHidden.native.tsx
│   │   │   │   ├── VisuallyHidden.tsx
│   │   │   │   ├── VisuallyHidden.web.tsx
│   │   │   │   └── VisuallyHidden.types.ts
│   │   │   └── Announcer/
│   │   │       ├── LiveAnnouncer.native.tsx
│   │   │       ├── LiveAnnouncer.tsx  # Live region announcements
│   │   │       ├── LiveAnnouncer.web.tsx
│   │   │       └── LiveAnnouncer.types.ts
│   │   ├── auction/
│   │   │   ├── AuctionTimer.native.tsx
│   │   │   ├── AuctionTimer.tsx  # Countdown timer
│   │   │   ├── AuctionTimer.web.tsx
│   │   │   ├── BidHistoryChart.native.tsx
│   │   │   ├── BidHistoryChart.tsx  # Bid history visualization
│   │   │   ├── BidHistoryChart.web.tsx
│   │   │   ├── LiveBidFeed.native.tsx
│   │   │   ├── LiveBidFeed.tsx  # Real-time bid feed
│   │   │   └── LiveBidFeed.web.tsx
│   │   ├── charts/
│   │   │   ├── EarningsChart.native.tsx
│   │   │   ├── EarningsChart.tsx  # Earnings visualization
│   │   │   ├── EarningsChart.web.tsx
│   │   │   ├── PerformanceChart.native.tsx
│   │   │   ├── PerformanceChart.tsx  # Performance metrics
│   │   │   ├── PerformanceChart.web.tsx
│   │   │   ├── TrendChart.native.tsx
│   │   │   ├── TrendChart.tsx  # Trend visualization
│   │   │   └── TrendChart.web.tsx
│   │   ├── collaboration/
│   │   │   ├── CollaborationPanel.native.tsx
│   │   │   ├── CollaborationPanel.tsx  # Team collaboration
│   │   │   ├── CollaborationPanel.web.tsx
│   │   │   ├── GroupCard.native.tsx
│   │   │   ├── GroupCard.tsx  # User group card
│   │   │   ├── GroupCard.web.tsx
│   │   │   ├── MentorCard.native.tsx
│   │   │   ├── MentorCard.tsx  # Mentor profile card
│   │   │   └── MentorCard.web.tsx
│   │   ├── compliance/
│   │   │   ├── DocumentUploader.native.tsx
│   │   │   ├── DocumentUploader.tsx  # Compliance doc uploader
│   │   │   ├── DocumentUploader.web.tsx
│   │   │   ├── VerificationStatus.native.tsx
│   │   │   ├── VerificationStatus.tsx  # Verification status badge
│   │   │   └── VerificationStatus.web.tsx
│   │   ├── components/
│   │   │   ├── Accordion/
│   │   │   ├── Alert/
│   │   │   ├── Avatar/
│   │   │   ├── Badge/
│   │   │   ├── Breadcrumb/
│   │   │   ├── Button/
│   │   │   │   ├── Button.native.tsx  # Native-specific overrides
│   │   │   │   ├── Button.stories.tsx  # Storybook stories
│   │   │   │   ├── Button.test.tsx  # Component tests
│   │   │   │   ├── Button.tsx  # Base button component
│   │   │   │   └── Button.web.tsx  # Web-specific overrides
│   │   │   ├── Card/
│   │   │   ├── Checkbox/
│   │   │   ├── DataTable/
│   │   │   ├── DatePicker/
│   │   │   ├── Dropdown/
│   │   │   ├── FileUpload/
│   │   │   ├── Input/
│   │   │   │   ├── Input.native.tsx
│   │   │   │   ├── Input.test.tsx
│   │   │   │   ├── Input.tsx
│   │   │   │   └── Input.web.tsx
│   │   │   ├── Modal/
│   │   │   ├── Pagination/
│   │   │   ├── Popover/
│   │   │   ├── Progress/
│   │   │   ├── Radio/
│   │   │   ├── Rating/
│   │   │   ├── Select/
│   │   │   ├── Skeleton/
│   │   │   ├── Slider/
│   │   │   ├── Stepper/
│   │   │   ├── Switch/
│   │   │   ├── Tabs/
│   │   │   ├── Textarea/
│   │   │   ├── Timeline/
│   │   │   ├── TimePicker/
│   │   │   ├── Toast/
│   │   │   └── Tooltip/
│   │   ├── forms/  # Form components
│   │   │   ├── FormError/
│   │   │   ├── FormField/
│   │   │   ├── FormGroup/
│   │   │   ├── FormHelper/
│   │   │   └── FormLabel/
│   │   ├── icons/  # Icon components
│   │   │   └── index.ts  # Export all icons
│   │   ├── layout/  # Layout components
│   │   │   ├── Container/
│   │   │   ├── Divider/
│   │   │   ├── Grid/
│   │   │   ├── Spacer/
│   │   │   └── Stack/
│   │   ├── learning/
│   │   │   ├── AchievementBadge.native.tsx
│   │   │   ├── AchievementBadge.tsx  # Achievement badge
│   │   │   ├── AchievementBadge.web.tsx
│   │   │   ├── LearningPathCard.native.tsx
│   │   │   ├── LearningPathCard.tsx  # Learning path card
│   │   │   ├── LearningPathCard.web.tsx
│   │   │   ├── ProgressTracker.native.tsx
│   │   │   ├── ProgressTracker.tsx  # Progress visualization
│   │   │   └── ProgressTracker.web.tsx
│   │   ├── tracking/
│   │   │   ├── TimeTracker.native.tsx
│   │   │   ├── TimeTracker.tsx  # Time tracking widget
│   │   │   ├── TimeTracker.web.tsx
│   │   │   ├── TimesheetTable.native.tsx
│   │   │   ├── TimesheetTable.tsx  # Timesheet grid
│   │   │   ├── TimesheetTable.web.tsx
│   │   │   ├── WorkDiaryEntry.native.tsx
│   │   │   ├── WorkDiaryEntry.tsx  # Work diary card
│   │   │   └── WorkDiaryEntry.web.tsx
│   │   └── video/
│   │       ├── VideoPlayer.native.tsx
│   │       ├── VideoPlayer.tsx  # Video player
│   │       ├── VideoPlayer.web.tsx
│   │       ├── VideoUploader.native.tsx
│   │       ├── VideoUploader.tsx  # Video upload
│   │       └── VideoUploader.web.tsx
│   ├── package.json
│   ├── README.md
│   └── tsconfig.json
└── shared/  # Shared libraries
    ├── src/
        ├── accessibility/
        │   ├── a11y-utils.ts  # Accessibility utilities
        │   ├── aria-utils.ts  # ARIA utilities
        │   ├── focus-management.ts  # Focus management
        │   ├── keyboard-navigation.ts  # Keyboard navigation
        │   ├── screen-reader.ts  # Screen reader utilities
        │   └── testing/
        │       ├── a11y-test-utils.ts  # Testing utilities
        │       └── axe-config.ts  # axe-core configuration
        ├── features/
        │   ├── bidding/
        │   │   ├── api/
        │   │   │   ├── bid-api.ts  # Bid placement API
        │   │   │   │   # BE: proposals-be/bid
        │   │   │   └── bid-strategy-api.ts  # Bid strategy API
        │   │   │       # BE: proposals-be/bid-strategy
        │   │   ├── hooks/
        │   │   │   ├── useBidAnalytics.ts  # Bid analytics
        │   │   │   ├── useBidHistory.ts  # Bid history
        │   │   │   ├── useBidStrategies.ts  # List strategies
        │   │   │   ├── useBidStrategy.ts  # Bid strategy management
        │   │   │   └── usePlaceBid.ts  # Place bid
        │   │   ├── queries/
        │   │   │   ├── bidding-mutations.ts  # Bidding mutations
        │   │   │   └── bidding-queries.ts  # Bidding queries
        │   │   └── types.ts  # Bidding types
        │   ├── compliance/
        │   │   ├── api/
│   │   │   ├── compliance-api.ts  # Compliance API
│   │   │   │   # BE: users-be/compliance
│   │   │   └── tax-profile-api.ts  # Tax profile API
│   │   │       # BE: users-be/compliance, financial-be/tax
│   │   ├── hooks/
│   │   │   ├── useComplianceDocuments.ts  # Document management
│   │   │   ├── useComplianceProfile.ts  # Compliance profile
│   │   │   ├── useTaxProfile.ts  # Tax profile management
│   │   │   └── useTaxReports.ts  # Tax reports
│   │   ├── queries/
│   │   │   ├── compliance-mutations.ts  # Compliance mutations
│   │   │   └── compliance-queries.ts  # Compliance queries
│   │   └── types.ts  # Compliance types
│   ├── feature-flags/
│   │   ├── api/
│   │   │   └── flags-api.ts  # Feature flags API
│   │   │       # BE: utility/flags
│   │   ├── hooks/
│   │   │   ├── useFeatureFlag.ts  # Check single flag
│   │   │   ├── useFeatureFlagVariant.ts  # A/B test variant
│   │   │   └── useFeatureFlags.ts  # Get all flags
│   │   ├── queries/
│   │   │   └── flags-queries.ts  # Flag queries
│   │   ├── store/
│   │   │   └── flags-store.ts  # Flags state (Zustand)
│   │   └── types.ts  # Flag types
│   ├── invitations/
│   │   ├── api/
│   │   │   ├── job-invitations-api.ts  # Job invitations API (client)
│   │   │   │   # BE: jobs-be/invitation
│   │   │   └── proposal-invites-api.ts  # Proposal invites API (freelancer)
│   │   │       # BE: proposals-be/invite
│   │   ├── hooks/
│   │   │   ├── useAcceptInvite.ts  # Accept invite (freelancer)
│   │   │   ├── useDeclineInvite.ts  # Decline invite (freelancer)
│   │   │   ├── useInvitationAnalytics.ts  # Invitation metrics
│   │   │   ├── useInvitations.ts  # Invitations management
│   │   │   └── useSendInvitation.ts  # Send invitation (client)
│   │   ├── queries/
│   │   │   ├── invitations-mutations.ts  # Invitation mutations
│   │   │   └── invitations-queries.ts  # Invitation queries
│   │   └── types.ts  # Invitation types
│   ├── networking/
│   │   ├── api/
│   │   │   ├── connections-api.ts  # Connections API
│   │   │   │   # BE: users-be/connections
│   │   │   ├── groups-api.ts  # Groups API
│   │   │   │   # BE: users-be/user_group
│   │   │   ├── networking-api.ts  # Networking API
│   │   │   │   # BE: users-be/connections
│   │   │   └── referrals-api.ts  # Referrals API
│   │   │       # BE: users-be/referral
│   │   ├── hooks/
│   │   │   ├── useConnectionRequest.ts  # Send/accept/reject
│   │   │   ├── useConnections.ts  # Connections management
│   │   │   ├── useGroups.ts  # Groups management
│   │   │   ├── useNetworkRecommendations.ts  # Connection recommendations
│   │   │   └── useReferrals.ts  # Referral management
│   │   ├── queries/
│   │   │   ├── networking-mutations.ts  # Networking mutations
│   │   │   └── networking-queries.ts  # Networking queries
│   │   └── types.ts  # Networking types
│   ├── realtime/
│   │   ├── hooks/
│   │   │   ├── usePresence.ts  # User presence (online/offline)
│   │   │   ├── useRealtimeAuction.ts  # Real-time auction updates
│   │   │   ├── useRealtimeMessages.ts  # Real-time messages
│   │   │   ├── useRealtimeNotifications.ts  # Real-time notifications
│   │   │   └── useWebSocket.ts  # WebSocket connection
│   │   ├── store/
│   │   │   └── realtime-store.ts  # Real-time state (Zustand)
│   │   ├── websocket/
│   │   │   ├── client.ts  # WebSocket client
│   │   │   ├── heartbeat.ts  # Connection health
│   │   │   └── reconnection.ts  # Reconnection logic
│   │   └── types.ts  # Real-time types
│   ├── search/
│   │   ├── api/
│   │   │   ├── recommendations-api.ts  # Recommendations API
│   │   │   │   # BE: search-be/recommendation
│   │   │   ├── saved-searches-api.ts  # Saved searches API
│   │   │   │   # BE: search-be/saved-search
│   │   │   ├── search-api.ts  # Search API (already may exist, ensuring completeness)
│   │   │   │   # BE: search-be/query
│   │   │   └── trending-api.ts  # Trending API
│   │   │       # BE: search-be/trending
│   │   ├── hooks/
│   │   │   ├── useRecommendations.ts  # Recommendations
│   │   │   ├── useSavedSearches.ts  # Saved searches
│   │   │   ├── useSearch.ts  # Search execution
│   │   │   ├── useSearchHistory.ts  # Search history
│   │   │   ├── useSearchSuggestions.ts  # Auto-complete suggestions
│   │   │   └── useTrending.ts  # Trending items
│   │   ├── queries/
│   │   │   ├── search-mutations.ts  # Search mutations
│   │   │   └── search-queries.ts  # Search queries
│   │   ├── store/
│   │   │   └── search-store.ts  # Search UI state (filters, etc.)
│   │   └── types.ts  # Search types
│   ├── shortlists/
│   │   ├── api/
│   │   │   └── shortlists-api.ts  # Shortlists API
│   │   │       # BE: jobs-be/shortlist
│   │   ├── hooks/
│   │   │   ├── useAddToShortlist.ts  # Add candidate
│   │   │   ├── useRemoveFromShortlist.ts  # Remove candidate
│   │   │   ├── useShortlist.ts  # Single shortlist
│   │   │   └── useShortlists.ts  # Shortlists management
│   │   ├── queries/
│   │   │   ├── shortlists-mutations.ts  # Shortlist mutations
│   │   │   └── shortlists-queries.ts  # Shortlist queries
│   │   └── types.ts  # Shortlist types
│   ├── utils/
│   │   └── logger/
│   │       ├── formatters/
│   │       │   ├── json.ts  # JSON formatter
│   │       │   └── pretty.ts  # Pretty formatter
│   │       ├── log-levels.ts  # Log levels
│   │       ├── logger.ts  # Logger implementation
│   │       └── transports/
│   │           ├── console.ts  # Console transport
│   │           ├── file.ts  # File transport
│   │           └── remote.ts  # Remote logging service
│   ├── monitoring/
│   │   ├── analytics/
│   │   │   ├── analytics-config.ts  # Analytics configuration
│   │   │   ├── providers/
│   │   │   │   ├── amplitude.ts  # Amplitude
│   │   │   │   ├── google-analytics.ts  # GA4
│   │   │   │   └── mixpanel.ts  # Mixpanel
│   │   │   └── tracking/
│   │   │       ├── event-tracker.ts
│   │   │       ├── page-tracker.ts
│   │   │       └── user-tracker.ts
│   │   ├── performance/
│   │   │   ├── api-monitor.ts  # API performance monitoring
│   │   │   ├── bundle-monitor.ts  # Bundle size monitoring
│   │   │   ├── performance-observer.ts  # Performance API
│   │   │   └── web-vitals.ts  # Web Vitals monitoring
│   │   └── sentry/
│   │       ├── error-boundary.tsx  # Sentry error boundary
│   │       ├── sentry-config.ts  # Sentry configuration
│   │       ├── sentry-init.native.ts  # Mobile initialization
│   │       └── sentry-init.web.ts  # Web initialization
│   ├── security/
│   │   ├── auth/
│   │   │   ├── device-trust.ts  # Device trust
│   │   │   ├── session-manager.ts  # Session management
│   │   │   ├── token-manager.ts  # Token management
│   │   │   └── token-refresh.ts  # Auto token refresh
│   │   ├── encryption/
│   │   │   ├── crypto-utils.ts  # Encryption utilities
│   │   │   ├── key-management.ts  # Key management
│   │   │   └── secure-storage.ts  # Secure storage
│   │   ├── headers/
│   │   │   ├── cors.ts  # CORS configuration
│   │   │   ├── csp.ts  # Content Security Policy
│   │   │   └── security-headers.ts  # Security headers config
│   │   ├── monitoring/
│   │   │   ├── anomaly-detector.ts  # Anomaly detection
│   │   │   ├── breach-detector.ts  # Breach detection
│   │   │   └── security-monitor.ts  # Security monitoring
│   │   └── validation/
│   │       ├── input-validator.ts  # Input validation
│   │       ├── sanitizer.ts  # HTML/input sanitization
│   │       ├── sql-injection-guard.ts  # SQL injection protection
│   │       └── xss-protection.ts  # XSS protection
│   ├── utils/
│   │   └── logger/
│   │       ├── formatters/
│   │       │   ├── json.ts  # JSON formatter
│   │       │   └── pretty.ts  # Pretty formatter
│   │       ├── log-levels.ts  # Log levels
│   │       ├── logger.ts  # Logger implementation
│   │       └── transports/
│   │           ├── console.ts  # Console transport
│   │           ├── file.ts  # File transport
│   │           └── remote.ts  # Remote logging service
│   └── package.json
```
<!-- PART 7/11: fe/docs -->
```
fe/
├── docs/
│   ├── adr/  # Architecture Decision Records
│   │   ├── 001-monorepo-structure.md
│   │   ├── 002-state-management.md
│   │   ├── 003-authentication-approach.md
│   │   ├── 004-component-library.md
│   │   └── ...
│   ├── api/
│   │   ├── authentication.md
│   │   ├── endpoints/
│   │   │   ├── contracts.md
│   │   │   ├── financial.md
│   │   │   ├── jobs.md
│   │   │   ├── messages.md
│   │   │   ├── notifications.md
│   │   │   ├── payments.md
│   │   │   ├── proposals.md
│   │   │   └── users.md
│   │   ├── errors.md
│   │   ├── getting-started.md
│   │   ├── rate-limiting.md
│   │   └── webhooks.md
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
│   │   │   ├── button.md
│   │   │   ├── input.md
│   │   │   └── ...
│   │   ├── component-library.md
│   │   ├── design-system.md
│   │   ├── molecules/
│   │   │   ├── card.md
│   │   │   ├── form-field.md
│   │   │   └── ...
│   │   ├── organisms/
│   │   │   ├── header.md
│   │   │   ├── sidebar.md
│   │   │   └── ...
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
│   ├── guides/
│   │   ├── contributing/
│   │   │   ├── code-review.md
│   │   │   ├── getting-started.md
│   │   │   └── pull-requests.md
│   │   ├── deployment/
│   │   │   ├── ci-cd.md
│   │   │   ├── mobile.md
│   │   │   └── web.md
│   │   ├── development/
│   │   │   ├── coding-standards.md
│   │   │   ├── debugging.md
│   │   │   ├── git-workflow.md
│   │   │   └── testing.md
│   │   └── setup/
│   │       ├── environment-variables.md
│   │       ├── local-development.md
│   │       └── troubleshooting.md
│   └── api-integration/
│       ├── admin-be.md
│       ├── communications-be.md
│       ├── contracts-be.md
│       ├── financial-be.md
│       ├── jobs-be.md
│       ├── proposals-be.md
│       ├── reviews-be.md
│       ├── search-be.md
│       ├── storage-be.md
│       ├── subscriptions-be.md
│       ├── users-be.md
│       └── utility-services.md
```
<!-- PART 8/11: fe/scripts -->
```
fe/
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
│   │   └── verify-types.ts  # Type check
│   ├── dev/
│   │   ├── reset-db.sh  # Reset local DB
│   │   ├── start-all.sh  # Start all services
│   │   ├── start-mobile.sh  # Start mobile only
│   │   └── start-web.sh  # Start web only
│   ├── deploy/
│   │   ├── deploy-mobile-prod.sh
│   │   ├── deploy-mobile-staging.sh
│   │   ├── deploy-web-prod.sh
│   │   └── deploy-web-staging.sh
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
│   ├── build-all.sh  # Build all apps
│   ├── check-bundle-limits.ts  # Enforce limits
│   ├── check-bundle-size.sh  # Bundle size analysis
│   ├── clean.sh  # Clean build artifacts
│   ├── db-seed.sh  # Seed local dev data
│   ├── dev.sh  # Start all dev servers
│   ├── generate-bundle-report.ts  # Generate reports
│   ├── generate-types.sh  # Generate types from OpenAPI
│   ├── setup.sh  # Initial project setup
│   └── test-all.sh  # Run all tests
```
<!-- PART 9/11: fe/config -->
```
fe/
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
```
<!-- PART 10/11: fe/tests -->
```
fe/
├── tests/
│   ├── e2e/
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
│   │   │   └── users-api.test.ts
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
```
<!-- PART 11/11: root files -->
```
fe/
├── .env.development  # Development environment
├── .env.example  # Environment variables template
├── .env.local  # Local environment variables
├── .env.local.example  # Local overrides template
├── .env.production  # Production environment
├── .env.staging  # Staging environment
├── .eslintrc.json  # Root ESLint config
│   # Web-specific ESLint rules
├── .gitignore  # Git ignore patterns
├── .nvmrc  # Node.js version (v20.x)
├── .prettierignore  # Prettier ignore patterns
├── .prettierrc  # Prettier configuration
├── package.json  # Root package (workspace manager)
│   # Scripts: dev, build, test, lint, type-check
├── pnpm-lock.yaml  # Locked dependencies
├── pnpm-workspace.yaml  # pnpm workspace configuration
├── README.md  # Root README
│   # Web app documentation
├── tsconfig.base.json  # Base TypeScript configuration
│   # Shared compiler options
├── tsconfig.json  # Root TypeScript config
└── turbo.json  # Turborepo pipeline configuration
    # Build cache, task dependencies
```