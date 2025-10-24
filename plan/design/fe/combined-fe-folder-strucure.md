```
fe/
└── fe/
    ├── .github/  # GitHub workflows
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
    │   │   ├── notify-deployment/
    │   │   │   └── action.yml
    │   │   ├── setup-node/
    │   │   │   └── action.yml
    │   │   └── setup-pnpm/
    │   │       └── action.yml
    │   ├── environments/
    │   │   ├── development.yml  # Development environment config
    │   │   ├── production.yml  # Production environment config
    │   │   └── staging.yml  # Staging environment config
    │   ├── secrets/
    │   │   └── README.md  # Secrets documentation
    │   ├── workflows/
    │   │   ├── accessibility.yml  # Accessibility checks
    │   │   │   # - Run axe-core
    │   │   │   # - Check WCAG compliance
    │   │   │   # A11y checks
    │   │   ├── bundle-analysis.yml  # Bundle size checks
    │   │   ├── cd-mobile-production.yml  # Mobile production deployment
    │   │   │   # - Build
    │   │   │   # - Submit to App Store/Play Store
    │   │   ├── cd-mobile-staging.yml  # Mobile staging deployment
    │   │   │   # - Build
    │   │   │   # - Submit to TestFlight/Internal Testing
    │   │   ├── cd-mobile.yml  # Mobile deployment
    │   │   │   # - EAS build
    │   │   │   # - Submit to stores
    │   │   ├── cd-web-production.yml  # Web production deployment
    │   │   │   # - Build
    │   │   │   # - Deploy to production
    │   │   │   # - Run smoke tests
    │   │   │   # - Notify team
    │   │   ├── cd-web-staging.yml  # Web staging deployment
    │   │   │   # - Build
    │   │   │   # - Deploy to staging
    │   │   │   # - Run smoke tests
    │   │   ├── cd-web.yml  # Web deployment
    │   │   │   # - Build Next.js production
    │   │   │   # - Deploy to K8s
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
    │   │   ├── ci.yml  # Continuous Integration
    │   │   │   # - Lint, type-check, test
    │   │   │   # - Build verification
    │   │   │   # - Bundle size check
    │   │   │   # Main CI pipeline
    │   │   │   # Continuous integration
    │   │   │   # - Lint
    │   │   │   # - Type check
    │   │   │   # - Unit tests
    │   │   │   # - Build
    │   │   ├── dependabot.yml  # Automated dependency updates
    │   │   ├── dependency-review.yml  # Dependency review
    │   │   │   # - Check for vulnerabilities
    │   │   │   # - License compliance
    │   │   │   # Dependency checks
    │   │   ├── dependency-update.yml  # Dependabot automation
    │   │   │   # Automated dependency updates
    │   │   ├── deploy-mobile-production.yml
    │   │   ├── deploy-mobile-staging.yml
    │   │   ├── deploy-web-production.yml
    │   │   ├── deploy-web-staging.yml
    │   │   ├── e2e-tests.yml  # E2E tests
    │   │   │   # - Setup test environment
    │   │   │   # - Run Playwright/Detox tests
    │   │   │   # - Upload test results
    │   │   │   # E2E test pipeline
    │   │   ├── lighthouse.yml  # Performance audits
    │   │   │   # - Run Lighthouse CI
    │   │   │   # - Compare against budgets
    │   │   │   # - Comment on PR
    │   │   │   # Performance checks
    │   │   ├── lint.yml  # Linting
    │   │   ├── mobile-build-android.yml  # Android build
    │   │   ├── mobile-build-ios.yml  # iOS build
    │   │   ├── performance-tests.yml  # Performance testing
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
    │   │   ├── type-check.yml  # TypeScript checks
    │   │   ├── visual-regression.yml  # Visual regression tests
    │   │   ├── web-deploy-production.yml  # Web production deployment
    │   │   └── web-deploy-staging.yml  # Web staging deployment
    │   └── CODEOWNERS  # Code ownership
    ├── .husky/  # Git hooks for quality gates
    │   ├── pre-commit  # Runs linting, type checking
    │   │   # Validates no console.log in prod code
    │   │   # Ensures all tests pass
    │   └── pre-push  # Runs full test suite before push
    │       # Bundle size check
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
    ├── apps/  # Application workspaces
    │   ├── mobile/  # React Native/Expo application
    │   │   ├── app/  # Expo Router file-based routing
    │   │   │   ├── (auth)/  # Auth screens
    │   │   │   │   ├── onboarding/
    │   │   │   │   │   └── kyc/
    │   │   │   │   │       └── index.tsx  # KYC onboarding (mobile)
    │   │   │   │   │           # - Upload ID docs, liveness
    │   │   │   │   │           # - Track review status
    │   │   │   │   │           # BE: admin-be/kyc_case
    │   │   │   │   │           # POST /v1/kyc/cases
    │   │   │   │   │           # GET  /v1/kyc/cases/{id}
    │   │   │   │   ├── _layout.tsx  # Auth layout
    │   │   │   │   ├── callback.tsx  # OAuth callback
    │   │   │   │   │   # BE: Keycloak token exchange
    │   │   │   │   ├── forgot-password.tsx  # Password reset
    │   │   │   │   │   # BE: users-be/security/recovery
    │   │   │   │   │   # POST /v1/auth/forgot-password
    │   │   │   │   │   # Password reset flow
    │   │   │   │   │   # - Email input
    │   │   │   │   │   # - Code verification
    │   │   │   │   │   # BE: users-be/auth
    │   │   │   │   ├── login.tsx  # Login screen
    │   │   │   │   │   # - Email/password form
    │   │   │   │   │   # - Social login (Google, Apple)
    │   │   │   │   │   # - Biometric login (Face ID, Touch ID)
    │   │   │   │   │   # BE: Keycloak OAuth2
    │   │   │   │   │   # POST /v1/auth/login
    │   │   │   │   ├── register.tsx  # Registration screen
    │   │   │   │   │   # BE: users-be/user
    │   │   │   │   │   # POST /v1/users/register
    │   │   │   │   ├── reset-password.tsx  # Reset password with code
    │   │   │   │   │   # - New password input
    │   │   │   │   │   # - Confirmation
    │   │   │   │   │   # BE: users-be/auth
    │   │   │   │   │   # POST /v1/auth/reset-password
    │   │   │   │   └── verify-email.tsx  # Email verification
    │   │   │   │       # - Code input
    │   │   │   │       # - Resend code
    │   │   │   │       # BE: users-be/auth
    │   │   │   │       # POST /v1/auth/verify-email
    │   │   │   ├── (billing)/
    │   │   │   │   ├── payout-methods/
    │   │   │   │   │   └── index.tsx  # Payout methods (mobile)
    │   │   │   │   │       # - Add/edit bank/card
    │   │   │   │   │       # BE: financial-be/payout_method
    │   │   │   │   │       # GET/POST/DELETE /v1/payout-methods
    │   │   │   │   └── wallet/
    │   │   │   │       └── index.tsx  # Wallet (mobile)
    │   │   │   │           # - Balance, transactions
    │   │   │   │           # - Request payout
    │   │   │   │           # BE: financial-be/wallet|payout
    │   │   │   │           # GET  /v1/wallet
    │   │   │   │           # GET  /v1/wallet/transactions
    │   │   │   │           # POST /v1/payouts
    │   │   │   ├── (contracts)/
    │   │   │   │   └── offers/
    │   │   │   │       └── [id]/
    │   │   │   │           └── index.tsx  # Offer (mobile - accept/decline)
    │   │   │   │               # - Review terms, accept/decline
    │   │   │   │               # BE: contracts-be/offer (proposed if missing)
    │   │   │   │               # GET  /v1/offers/{id}
    │   │   │   │               # POST /v1/offers/{id}/accept|decline
    │   │   │   ├── (dashboard)/
    │   │   │   │   ├── activity-feed/
    │   │   │   │   │   └── index.tsx  # Activity feed
    │   │   │   │   │       # - Recent activities
    │   │   │   │   │       # - Notifications inline
    │   │   │   │   │       # - Quick actions from feed
    │   │   │   │   │       # BE: communications-be/notification, utility/activity
    │   │   │   │   │       # GET /v1/activity/feed
    │   │   │   │   ├── quick-actions/
    │   │   │   │   │   └── index.tsx  # Quick actions screen
    │   │   │   │   │       # - Quick message
    │   │   │   │   │       # - Quick proposal
    │   │   │   │   │       # - Quick time entry
    │   │   │   │   │       # - Quick invoice
    │   │   │   │   │       # BE: Multiple services
    │   │   │   │   ├── today/
    │   │   │   │   │   └── index.tsx  # Today view
    │   │   │   │   │       # - Today's schedule
    │   │   │   │   │       # - Pending tasks
    │   │   │   │   │       # - Quick metrics
    │   │   │   │   │       # BE: contracts-be/work_diary, communications-be/notification
    │   │   │   │   │       # GET /v1/today/overview
    │   │   │   │   └── widgets/
    │   │   │   │       ├── earnings/
    │   │   │   │       │   └── index.tsx  # Earnings widget
    │   │   │   │       │       # BE: financial-be/wallet
    │   │   │   │       ├── notifications/
    │   │   │   │       │   └── index.tsx  # Notifications widget
    │   │   │   │       │       # BE: communications-be/notification
    │   │   │   │       └── time-tracker/
    │   │   │   │           └── index.tsx  # Time tracker widget
    │   │   │   │               # BE: contracts-be/work_diary
    │   │   │   ├── (inbox)/
    │   │   │   │   ├── messages/
    │   │   │   │   │   ├── [conversationId]/
    │   │   │   │   │   │   └── index.tsx  # Messages (mobile - thread)
    │   │   │   │   │   │       # - View/send messages, attachments
    │   │   │   │   │   │       # BE: communications-be/conversation|message
    │   │   │   │   │   │       # GET  /v1/conversations/{id}
    │   │   │   │   │   │       # GET  /v1/conversations/{id}/messages
    │   │   │   │   │   │       # POST /v1/messages
    │   │   │   │   │   └── index.tsx  # Messages (mobile - inbox)
    │   │   │   │   │       # - List user conversations
    │   │   │   │   │       # BE: communications-be/conversation
    │   │   │   │   │       # GET /v1/conversations?mine=1
    │   │   │   │   └── proposals/
    │   │   │   │       └── index.tsx  # Proposals inbox (mobile)
    │   │   │   │           # - List/detail, statuses
    │   │   │   │           # BE: proposals-be/proposal
    │   │   │   │           # GET /v1/proposals?mine=1
    │   │   │   ├── (market)/
    │   │   │   │   └── jobs/
    │   │   │   │       └── index.tsx  # Job discovery (mobile)
    │   │   │   │           # - Feed, trending, filters
    │   │   │   │           # BE: search-be/feed|trending|query
    │   │   │   │           # GET /v1/feed
    │   │   │   │           # GET /v1/trending
    │   │   │   ├── (onboarding)/  # Onboarding flow
    │   │   │   │   ├── _layout.tsx  # Onboarding layout
    │   │   │   │   ├── complete.tsx  # Onboarding complete
    │   │   │   │   ├── preferences.tsx  # Preferences
    │   │   │   │   ├── profile.tsx  # Basic profile setup
    │   │   │   │   ├── skills.tsx  # Skills selection
    │   │   │   │   └── welcome.tsx  # Welcome screen
    │   │   │   ├── (tabs)/  # Bottom tabs navigation
    │   │   │   │   ├── admin/
    │   │   │   │   │   ├── kyc/
    │   │   │   │   │   │   ├── [kycId].tsx  # KYC case detail (mobile)
    │   │   │   │   │   │   │   # BE: admin-be/kyc-case
    │   │   │   │   │   │   │   # GET /v1/admin/kyc/{kyc_id}
    │   │   │   │   │   │   └── index.tsx  # KYC queue (mobile)
    │   │   │   │   │   │       # BE: admin-be/kyc-case
    │   │   │   │   │   │       # GET /v1/admin/kyc
    │   │   │   │   │   ├── moderation/
    │   │   │   │   │   │   ├── [reportId].tsx  # Report detail (mobile)
    │   │   │   │   │   │   │   # BE: admin-be/moderation
    │   │   │   │   │   │   │   # GET /v1/admin/reports/{report_id}
    │   │   │   │   │   │   └── index.tsx  # Moderation queue (mobile)
    │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │       # GET /v1/admin/reports
    │   │   │   │   │   ├── users/
    │   │   │   │   │   │   ├── [userId].tsx  # User detail (mobile admin)
    │   │   │   │   │   │   │   # BE: users-be/profile
    │   │   │   │   │   │   │   # GET /v1/admin/users/{user_id}
    │   │   │   │   │   │   └── index.tsx  # Users search (mobile)
    │   │   │   │   │   │       # BE: search-be/query
    │   │   │   │   │   │       # POST /v1/admin/users/search
    │   │   │   │   │   ├── _layout.tsx  # Admin tab layout
    │   │   │   │   │   └── index.tsx  # Admin dashboard (mobile)
    │   │   │   │   │       # BE: admin-be/analytics
    │   │   │   │   │       # GET /v1/admin/dashboard
    │   │   │   │   ├── browse/
    │   │   │   │   │   ├── freelancers/
    │   │   │   │   │   │   ├── [userId].tsx  # Freelancer profile (mobile)
    │   │   │   │   │   │   │   # BE: users-be/profile
    │   │   │   │   │   │   │   # GET /v1/users/{user_id}
    │   │   │   │   │   │   ├── filters.tsx  # Freelancer search filters
    │   │   │   │   │   │   └── index.tsx  # Freelancers list (mobile)
    │   │   │   │   │   │       # BE: search-be/query
    │   │   │   │   │   │       # GET /v1/search/freelancers
    │   │   │   │   │   ├── jobs/
    │   │   │   │   │   │   ├── [jobId].tsx  # Job detail (mobile)
    │   │   │   │   │   │   │   # BE: jobs-be/job
    │   │   │   │   │   │   │   # GET /v1/jobs/{job_id}
    │   │   │   │   │   │   ├── filters.tsx  # Job search filters
    │   │   │   │   │   │   └── index.tsx  # Jobs list (mobile)
    │   │   │   │   │   │       # BE: search-be/query
    │   │   │   │   │   │       # GET /v1/search/jobs
    │   │   │   │   │   ├── _layout.tsx  # Browse tab layout
    │   │   │   │   │   └── categories.tsx  # Browse categories
    │   │   │   │   │       # BE: jobs-be/categories
    │   │   │   │   │       # GET /v1/jobs/categories
    │   │   │   │   ├── contracts/
    │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   ├── chat.tsx  # Contract chat (mobile)
    │   │   │   │   │   │   │   # BE: communications-be/conversation
    │   │   │   │   │   │   │   # GET /v1/contracts/{contract_id}/messages
    │   │   │   │   │   │   ├── deliverables.tsx  # Deliverables (mobile)
    │   │   │   │   │   │   │   # BE: contracts-be/deliverable
    │   │   │   │   │   │   │   # GET /v1/contracts/{contract_id}/deliverables
    │   │   │   │   │   │   ├── details.tsx  # Contract details (mobile)
    │   │   │   │   │   │   │   # BE: contracts-be/contract
    │   │   │   │   │   │   │   # GET /v1/contracts/{contract_id}
    │   │   │   │   │   │   ├── disputes.tsx  # Contract disputes (mobile)
    │   │   │   │   │   │   │   # BE: contracts-be/dispute
    │   │   │   │   │   │   │   # GET /v1/contracts/{contract_id}/disputes
    │   │   │   │   │   │   ├── milestones.tsx  # Milestones (mobile)
    │   │   │   │   │   │   │   # BE: contracts-be/milestone
    │   │   │   │   │   │   │   # GET /v1/contracts/{contract_id}/milestones
    │   │   │   │   │   │   └── workdiary.tsx  # Work diary (mobile)
    │   │   │   │   │   │       # BE: contracts-be/workdiary
    │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/workdiary
    │   │   │   │   │   ├── _layout.tsx  # Contracts tab layout
    │   │   │   │   │   ├── active.tsx  # Active contracts (mobile)
    │   │   │   │   │   │   # BE: contracts-be/contract
    │   │   │   │   │   │   # GET /v1/contracts?status=active
    │   │   │   │   │   ├── completed.tsx  # Completed contracts (mobile)
    │   │   │   │   │   │   # BE: contracts-be/contract
    │   │   │   │   │   │   # GET /v1/contracts?status=completed
    │   │   │   │   │   └── index.tsx  # All contracts (mobile)
    │   │   │   │   │       # BE: contracts-be/contract
    │   │   │   │   │       # GET /v1/contracts
    │   │   │   │   ├── dashboard/
    │   │   │   │   │   ├── _layout.tsx  # Dashboard tab layout
    │   │   │   │   │   ├── analytics.tsx  # User analytics (mobile)
    │   │   │   │   │   │   # BE: users-be/analytics
    │   │   │   │   │   │   # GET /v1/users/me/analytics
    │   │   │   │   │   ├── earnings.tsx  # Earnings overview (mobile)
    │   │   │   │   │   │   # BE: financial-be/reports
    │   │   │   │   │   │   # GET /v1/reports/earnings
    │   │   │   │   │   ├── index.tsx  # Main dashboard (mobile)
    │   │   │   │   │   │   # BE: Multiple services
    │   │   │   │   │   │   # Aggregated dashboard data
    │   │   │   │   │   └── notifications.tsx  # Notifications (mobile)
    │   │   │   │   │       # BE: communications-be/notification
    │   │   │   │   │       # GET /v1/notifications
    │   │   │   │   ├── financial/
    │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   ├── [disputeId].tsx  # Dispute detail (mobile)
    │   │   │   │   │   │   │   # BE: contracts-be/dispute
    │   │   │   │   │   │   │   # GET /v1/disputes/{dispute_id}
    │   │   │   │   │   │   └── index.tsx  # Disputes list (mobile)
    │   │   │   │   │   │       # BE: contracts-be/dispute
    │   │   │   │   │   │       # GET /v1/disputes
    │   │   │   │   │   ├── invoices/
    │   │   │   │   │   │   ├── [invoiceId].tsx  # Invoice detail (mobile)
    │   │   │   │   │   │   │   # BE: financial-be/invoice
    │   │   │   │   │   │   │   # GET /v1/invoices/{invoice_id}
    │   │   │   │   │   │   └── index.tsx  # Invoices list (mobile)
    │   │   │   │   │   │       # BE: financial-be/invoice
    │   │   │   │   │   │       # GET /v1/invoices
    │   │   │   │   │   ├── payouts/
    │   │   │   │   │   │   ├── [payoutId].tsx  # Payout detail (mobile)
    │   │   │   │   │   │   │   # BE: financial-be/payout
    │   │   │   │   │   │   │   # GET /v1/payouts/{payout_id}
    │   │   │   │   │   │   ├── index.tsx  # Payouts list (mobile)
    │   │   │   │   │   │   │   # BE: financial-be/payout
    │   │   │   │   │   │   │   # GET /v1/payouts
    │   │   │   │   │   │   └── request.tsx  # Request payout (mobile)
    │   │   │   │   │   │       # BE: financial-be/payout
    │   │   │   │   │   │       # POST /v1/payouts
    │   │   │   │   │   ├── _layout.tsx  # Financial tab layout
    │   │   │   │   │   ├── transactions.tsx  # Transaction history (mobile)
    │   │   │   │   │   │   # BE: financial-be/transaction
    │   │   │   │   │   │   # GET /v1/transactions
    │   │   │   │   │   └── wallet.tsx  # Wallet (mobile)
    │   │   │   │   │       # BE: financial-be/wallet
    │   │   │   │   │       # GET /v1/wallet
    │   │   │   │   ├── jobs/
    │   │   │   │   │   ├── [jobId]/
    │   │   │   │   │   │   ├── apply.tsx  # Apply to job (mobile)
    │   │   │   │   │   │   │   # BE: proposals-be/proposal
    │   │   │   │   │   │   │   # POST /v1/proposals
    │   │   │   │   │   │   ├── details.tsx  # Job details (mobile)
    │   │   │   │   │   │   │   # BE: jobs-be/job
    │   │   │   │   │   │   │   # GET /v1/jobs/{job_id}
    │   │   │   │   │   │   └── proposals.tsx  # Job proposals (client view, mobile)
    │   │   │   │   │   │       # BE: proposals-be/proposal
    │   │   │   │   │   │       # GET /v1/jobs/{job_id}/proposals
    │   │   │   │   │   ├── _layout.tsx  # Jobs tab layout
    │   │   │   │   │   ├── create.tsx  # Create job (mobile)
    │   │   │   │   │   │   # BE: jobs-be/job
    │   │   │   │   │   │   # POST /v1/jobs
    │   │   │   │   │   ├── drafts.tsx  # Job drafts (mobile)
    │   │   │   │   │   │   # BE: jobs-be/draft
    │   │   │   │   │   │   # GET /v1/jobs/drafts
    │   │   │   │   │   ├── index.tsx  # Browse jobs (mobile)
    │   │   │   │   │   │   # BE: search-be/query
    │   │   │   │   │   │   # GET /v1/search/jobs
    │   │   │   │   │   ├── my-jobs.tsx  # My posted jobs (mobile)
    │   │   │   │   │   │   # BE: jobs-be/job
    │   │   │   │   │   │   # GET /v1/jobs/my-jobs
    │   │   │   │   │   └── saved.tsx  # Saved jobs (mobile)
    │   │   │   │   │       # BE: jobs-be/saved
    │   │   │   │   │       # GET /v1/jobs/saved
    │   │   │   │   ├── messages/
    │   │   │   │   │   ├── [conversationId]/
    │   │   │   │   │   │   ├── chat.tsx  # Chat interface (mobile)
    │   │   │   │   │   │   │   # BE: communications-be/conversation
    │   │   │   │   │   │   │   # GET /v1/conversations/{conversation_id}/messages
    │   │   │   │   │   │   │   # POST /v1/conversations/{conversation_id}/messages
    │   │   │   │   │   │   └── details.tsx  # Conversation details (mobile)
    │   │   │   │   │   │       # BE: communications-be/conversation
    │   │   │   │   │   │       # GET /v1/conversations/{conversation_id}
    │   │   │   │   │   ├── _layout.tsx  # Messages tab layout
    │   │   │   │   │   ├── archived.tsx  # Archived conversations (mobile)
    │   │   │   │   │   │   # BE: communications-be/conversation
    │   │   │   │   │   │   # GET /v1/conversations?archived=true
    │   │   │   │   │   ├── compose.tsx  # New message (mobile)
    │   │   │   │   │   │   # BE: communications-be/conversation
    │   │   │   │   │   │   # POST /v1/conversations
    │   │   │   │   │   └── index.tsx  # Conversations list (mobile)
    │   │   │   │   │       # BE: communications-be/conversation
    │   │   │   │   │       # GET /v1/conversations
    │   │   │   │   ├── more/
    │   │   │   │   │   ├── _layout.tsx  # More menu stack
    │   │   │   │   │   ├── about.tsx  # About app
    │   │   │   │   │   │   # BE: none (static)
    │   │   │   │   │   ├── account.tsx  # Account settings
    │   │   │   │   │   │   # BE: users-be/account
    │   │   │   │   │   ├── help.tsx  # Help center
    │   │   │   │   │   │   # BE: CMS
    │   │   │   │   │   └── index.tsx  # More menu home
    │   │   │   │   ├── notifications/
    │   │   │   │   │   ├── [notificationId].tsx  # Notification detail
    │   │   │   │   │   │   # BE: communications-be/notification
    │   │   │   │   │   ├── _layout.tsx  # Notifications stack
    │   │   │   │   │   ├── index.tsx  # All notifications
    │   │   │   │   │   │   # BE: communications-be/notification
    │   │   │   │   │   └── settings.tsx  # Notification settings
    │   │   │   │   │       # BE: communications-be/preferences
    │   │   │   │   ├── profile/
    │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   ├── basic.tsx  # Edit basic info (mobile)
    │   │   │   │   │   │   │   # BE: users-be/profile
    │   │   │   │   │   │   │   # PUT /v1/users/me/profile
    │   │   │   │   │   │   ├── education.tsx  # Edit education (mobile)
    │   │   │   │   │   │   │   # BE: users-be/profile
    │   │   │   │   │   │   │   # PUT /v1/users/me/education
    │   │   │   │   │   │   ├── experience.tsx  # Edit experience (mobile)
    │   │   │   │   │   │   │   # BE: users-be/profile
    │   │   │   │   │   │   │   # PUT /v1/users/me/experience
    │   │   │   │   │   │   ├── portfolio.tsx  # Edit portfolio (mobile)
    │   │   │   │   │   │   │   # BE: users-be/portfolio, storage-be/asset
    │   │   │   │   │   │   │   # PUT /v1/users/me/portfolio
    │   │   │   │   │   │   └── skills.tsx  # Edit skills (mobile)
    │   │   │   │   │   │       # BE: users-be/profile
    │   │   │   │   │   │       # PUT /v1/users/me/skills
    │   │   │   │   │   ├── portfolio/
    │   │   │   │   │   │   ├── [itemId].tsx  # Portfolio item detail (mobile)
    │   │   │   │   │   │   │   # BE: users-be/portfolio
    │   │   │   │   │   │   │   # GET /v1/users/me/portfolio/{item_id}
    │   │   │   │   │   │   ├── add.tsx  # Add portfolio item (mobile)
    │   │   │   │   │   │   │   # BE: users-be/portfolio, storage-be/asset
    │   │   │   │   │   │   │   # POST /v1/users/me/portfolio
    │   │   │   │   │   │   └── index.tsx  # Portfolio list (mobile)
    │   │   │   │   │   │       # BE: users-be/portfolio
    │   │   │   │   │   │       # GET /v1/users/me/portfolio
    │   │   │   │   │   ├── _layout.tsx  # Profile tab layout
    │   │   │   │   │   ├── index.tsx  # View profile (mobile)
    │   │   │   │   │   │   # BE: users-be/profile
    │   │   │   │   │   │   # GET /v1/users/me
    │   │   │   │   │   ├── reviews.tsx  # User reviews (mobile)
    │   │   │   │   │   │   # BE: reviews-be/review
    │   │   │   │   │   │   # GET /v1/users/me/reviews
    │   │   │   │   │   └── settings.tsx  # Profile settings (mobile)
    │   │   │   │   │       # BE: users-be/preferences
    │   │   │   │   │       # GET /v1/users/me/preferences
    │   │   │   │   ├── proposals/
    │   │   │   │   │   ├── [proposalId]/
    │   │   │   │   │   │   ├── details.tsx  # Proposal details (mobile)
    │   │   │   │   │   │   │   # BE: proposals-be/proposal
    │   │   │   │   │   │   │   # GET /v1/proposals/{proposal_id}
    │   │   │   │   │   │   ├── edit.tsx  # Edit proposal (mobile)
    │   │   │   │   │   │   │   # BE: proposals-be/proposal
    │   │   │   │   │   │   │   # PUT /v1/proposals/{proposal_id}
    │   │   │   │   │   │   └── withdraw.tsx  # Withdraw proposal (mobile)
    │   │   │   │   │   │       # BE: proposals-be/proposal
    │   │   │   │   │   │       # POST /v1/proposals/{proposal_id}/withdraw
    │   │   │   │   │   ├── _layout.tsx  # Proposals tab layout
    │   │   │   │   │   ├── active.tsx  # Active proposals (mobile)
    │   │   │   │   │   │   # BE: proposals-be/proposal
    │   │   │   │   │   │   # GET /v1/proposals?status=active
    │   │   │   │   │   ├── archived.tsx  # Archived proposals (mobile)
    │   │   │   │   │   │   # BE: proposals-be/proposal
    │   │   │   │   │   │   # GET /v1/proposals?status=archived
    │   │   │   │   │   ├── drafts.tsx  # Proposal drafts (mobile)
    │   │   │   │   │   │   # BE: proposals-be/draft
    │   │   │   │   │   │   # GET /v1/proposals/drafts
    │   │   │   │   │   └── index.tsx  # All proposals (mobile)
    │   │   │   │   │       # BE: proposals-be/proposal
    │   │   │   │   │       # GET /v1/proposals
    │   │   │   │   ├── search/
    │   │   │   │   │   ├── _layout.tsx  # Search stack navigator
    │   │   │   │   │   │   # Search tab layout
    │   │   │   │   │   ├── filters.tsx  # Advanced filters
    │   │   │   │   │   │   # BE: search-be/facets
    │   │   │   │   │   │   # Search filters (mobile)
    │   │   │   │   │   ├── freelancers.tsx  # Search freelancers (mobile)
    │   │   │   │   │   │   # BE: search-be/query
    │   │   │   │   │   │   # GET /v1/search/freelancers
    │   │   │   │   │   ├── index.tsx  # Search home
    │   │   │   │   │   │   # BE: search-be/query
    │   │   │   │   │   ├── jobs.tsx  # Search jobs (mobile)
    │   │   │   │   │   │   # BE: search-be/query
    │   │   │   │   │   │   # GET /v1/search/jobs
    │   │   │   │   │   ├── portfolios.tsx  # Search portfolios (mobile)
    │   │   │   │   │   │   # BE: search-be/query
    │   │   │   │   │   │   # GET /v1/search/portfolios
    │   │   │   │   │   ├── results.tsx  # Search results
    │   │   │   │   │   │   # BE: search-be/query
    │   │   │   │   │   ├── saved-searches.tsx  # Saved searches (mobile)
    │   │   │   │   │   │   # BE: search-be/saved-search
    │   │   │   │   │   │   # GET /v1/search/saved-searches
    │   │   │   │   │   └── saved.tsx  # Saved searches
    │   │   │   │   │       # BE: search-be/saved-search
    │   │   │   │   ├── settings/
    │   │   │   │   │   ├── account/
    │   │   │   │   │   │   ├── close.tsx  # Close account (mobile)
    │   │   │   │   │   │   │   # BE: users-be/account
    │   │   │   │   │   │   │   # POST /v1/users/me/close-account
    │   │   │   │   │   │   ├── email.tsx  # Change email (mobile)
    │   │   │   │   │   │   │   # BE: users-be/account
    │   │   │   │   │   │   │   # PUT /v1/users/me/email
    │   │   │   │   │   │   ├── password.tsx  # Change password (mobile)
    │   │   │   │   │   │   │   # BE: users-be/auth
    │   │   │   │   │   │   │   # PUT /v1/auth/password
    │   │   │   │   │   │   └── phone.tsx  # Change phone (mobile)
    │   │   │   │   │   │       # BE: users-be/account
    │   │   │   │   │   │       # PUT /v1/users/me/phone
    │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   ├── payment-methods.tsx  # Payment methods (mobile)
    │   │   │   │   │   │   │   # BE: financial-be/payment-method
    │   │   │   │   │   │   │   # GET /v1/payment-methods
    │   │   │   │   │   │   └── subscription.tsx  # Subscription (mobile)
    │   │   │   │   │   │       # BE: financial-be/subscription
    │   │   │   │   │   │       # GET /v1/subscriptions
    │   │   │   │   │   ├── _layout.tsx  # Settings tab layout
    │   │   │   │   │   ├── index.tsx  # Settings overview (mobile)
    │   │   │   │   │   ├── notifications.tsx  # Notification settings (mobile)
    │   │   │   │   │   │   # BE: communications-be/preferences
    │   │   │   │   │   │   # GET /v1/notifications/preferences
    │   │   │   │   │   ├── privacy.tsx  # Privacy settings (mobile)
    │   │   │   │   │   │   # BE: users-be/privacy
    │   │   │   │   │   │   # GET /v1/users/me/privacy
    │   │   │   │   │   └── security.tsx  # Security settings (mobile)
    │   │   │   │   │       # - Two-factor auth
    │   │   │   │   │       # - Active sessions
    │   │   │   │   │       # BE: users-be/mfa, users-be/sessions
    │   │   │   │   │       # GET /v1/users/me/mfa
    │   │   │   │   │       # GET /v1/users/me/sessions
    │   │   │   │   ├── _layout.tsx  # Tabs layout
    │   │   │   │   │   # - Bottom tab navigator
    │   │   │   │   │   # - Tab bar icons
    │   │   │   │   │   # - Badge indicators (messages, notifications)
    │   │   │   │   ├── index.tsx  # Home tab / Dashboard
    │   │   │   │   │   # - Dashboard overview
    │   │   │   │   │   # - Quick actions
    │   │   │   │   │   # - Recent activity
    │   │   │   │   │   # BE: Multiple services (same as web dashboard)
    │   │   │   │   ├── jobs.tsx  # Jobs tab
    │   │   │   │   │   # - Job listings (browse/my-jobs based on role)
    │   │   │   │   │   # - Search jobs
    │   │   │   │   │   # - Saved jobs
    │   │   │   │   │   # BE: jobs-be/job
    │   │   │   │   │   # GET /v1/jobs/browse (freelancer)
    │   │   │   │   │   # GET /v1/jobs/my-jobs (client)
    │   │   │   │   ├── messages.tsx  # Messages tab
    │   │   │   │   │   # - Conversation list
    │   │   │   │   │   # - Real-time updates
    │   │   │   │   │   # BE: communications-be/conversations
    │   │   │   │   │   # GET /v1/conversations
    │   │   │   │   │   # WebSocket connection
    │   │   │   │   ├── profile.tsx  # Profile tab
    │   │   │   │   │   # - Current user profile preview
    │   │   │   │   │   # - Quick links to settings
    │   │   │   │   │   # - Stats
    │   │   │   │   │   # BE: users-be/profile
    │   │   │   │   │   # GET /v1/users/me
    │   │   │   │   └── proposals.tsx  # Proposals tab
    │   │   │   │       # - My proposals (freelancer)
    │   │   │   │       # - Received proposals (client, redirects to job)
    │   │   │   │       # BE: proposals-be
    │   │   │   │       # GET /v1/proposals/my-proposals
    │   │   │   ├── (work)/
    │   │   │   │   └── contracts/
    │   │   │   │       └── [id]/
    │   │   │   │           ├── milestones/
    │   │   │   │           │   └── index.tsx  # Milestones (mobile)
    │   │   │   │           │       # - View progress, submit for review
    │   │   │   │           │       # BE: contracts-be/milestone
    │   │   │   │           │       # GET  /v1/contracts/{id}/milestones
    │   │   │   │           └── work-diary/
    │   │   │   │               └── index.tsx  # Work diary & timesheets (mobile)
    │   │   │   │                   # - Log time/screenshots
    │   │   │   │                   # BE: contracts-be/workdiary|timesheet
    │   │   │   │                   # GET  /v1/contracts/{id}/work-diary
    │   │   │   │                   # POST /v1/contracts/{id}/timesheets
    │   │   │   ├── accessibility/
    │   │   │   │   ├── screen-reader/
    │   │   │   │   │   └── index.tsx  # Screen reader optimized view
    │   │   │   │   └── voice-commands/
    │   │   │   │       └── index.tsx  # Voice commands
    │   │   │   │           # - Voice-to-text
    │   │   │   │           # - Command shortcuts
    │   │   │   ├── billing/
    │   │   │   │   └── subscriptions/
    │   │   │   │       └── index.tsx  # Subscriptions (mobile)
    │   │   │   │           # - View current plan and status
    │   │   │   │           # - Upgrade/downgrade/cancel (with proration)
    │   │   │   │           # - View invoices/history
    │   │   │   │           # BE: financial-be/subscription
    │   │   │   │           # GET  /v1/subscriptions/me
    │   │   │   │           # POST /v1/subscriptions/change
    │   │   │   │           # POST /v1/subscriptions/cancel
    │   │   │   │           # GET  /v1/subscriptions/{id}/invoices
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
    │   │   │   ├── contracts/
    │   │   │   │   ├── [id]/
    │   │   │   │   │   ├── index.tsx  # Contract overview
    │   │   │   │   │   │   # BE: contracts-be/contract
    │   │   │   │   │   │   # GET /v1/contracts/{contract_id}
    │   │   │   │   │   ├── messages.tsx  # Contract messages
    │   │   │   │   │   │   # BE: communications-be
    │   │   │   │   │   │   # GET /v1/contracts/{contract_id}/conversation
    │   │   │   │   │   ├── milestones.tsx  # Milestones
    │   │   │   │   │   │   # BE: contracts-be/milestone
    │   │   │   │   │   │   # GET /v1/contracts/{contract_id}/milestones
    │   │   │   │   │   ├── timesheet.tsx  # Timesheet
    │   │   │   │   │   │   # BE: contracts-be/timesheet
    │   │   │   │   │   │   # GET /v1/contracts/{contract_id}/timesheets
    │   │   │   │   │   └── work-diary.tsx  # Work diary
    │   │   │   │   │       # BE: contracts-be/work_diary
    │   │   │   │   │       # GET /v1/contracts/{contract_id}/work-diary
    │   │   │   │   └── index.tsx  # Contracts list
    │   │   │   │       # BE: contracts-be/contract
    │   │   │   │       # GET /v1/contracts
    │   │   │   ├── financials/
    │   │   │   │   ├── invoices.tsx  # Invoices
    │   │   │   │   │   # BE: financial-be/invoice
    │   │   │   │   │   # GET /v1/invoices
    │   │   │   │   ├── transactions.tsx  # Transaction history
    │   │   │   │   │   # BE: financial-be/transaction
    │   │   │   │   │   # GET /v1/transactions
    │   │   │   │   └── wallet.tsx  # Wallet
    │   │   │   │       # BE: financial-be/wallet
    │   │   │   │       # GET /v1/wallet/balance
    │   │   │   ├── jobs/
    │   │   │   │   ├── [id].tsx  # Job detail
    │   │   │   │   │   # BE: jobs-be/job
    │   │   │   │   │   # GET /v1/jobs/{job_id}
    │   │   │   │   ├── post.tsx  # Post job (client)
    │   │   │   │   │   # BE: jobs-be/job
    │   │   │   │   │   # POST /v1/jobs
    │   │   │   │   └── search.tsx  # Job search
    │   │   │   │       # BE: search-be
    │   │   │   │       # POST /v1/search/jobs
    │   │   │   ├── messages/
    │   │   │   │   └── [conversationId].tsx  # Conversation thread
    │   │   │   │       # - Message list
    │   │   │   │       # - Message composer
    │   │   │   │       # - Real-time updates
    │   │   │   │       # BE: communications-be/messages
    │   │   │   │       # GET /v1/conversations/{conversation_id}/messages
    │   │   │   │       # POST /v1/messages
    │   │   │   │       # WebSocket updates
    │   │   │   ├── mobile-settings/
    │   │   │   │   ├── app-settings.tsx  # App-specific settings
    │   │   │   │   │   # - Cache management
    │   │   │   │   │   # - Offline mode
    │   │   │   │   │   # - Data usage
    │   │   │   │   ├── biometric.tsx  # Biometric authentication
    │   │   │   │   │   # - FaceID/TouchID setup
    │   │   │   │   │   # BE: users-be/auth
    │   │   │   │   │   # POST /v1/auth/biometric/register
    │   │   │   │   └── haptics.tsx  # Haptic feedback settings
    │   │   │   ├── notifications/
    │   │   │   │   └── index.tsx  # Notifications list
    │   │   │   │       # BE: communications-be/notifications
    │   │   │   │       # GET /v1/notifications
    │   │   │   │       # Notifications center (mobile)
    │   │   │   │       # - In-app alerts list, mark read
    │   │   │   │       # BE: communications-be/notification
    │   │   │   │       # GET  /v1/notifications?mine=1
    │   │   │   │       # POST /v1/notifications/{id}/read
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
    │   │   │   │   ├── sync/
    │   │   │   │   │   └── index.tsx  # Sync status
    │   │   │   │   │       # - Sync progress
    │   │   │   │   │       # - Conflict resolution
    │   │   │   │   │       # BE: Multiple services (sync endpoints)
    │   │   │   │   ├── queue.tsx  # Offline actions queue
    │   │   │   │   │   # - Pending uploads
    │   │   │   │   │   # - Queued messages
    │   │   │   │   │   # - Draft proposals
    │   │   │   │   └── sync.tsx  # Sync status
    │   │   │   │       # - Sync progress
    │   │   │   │       # - Conflict resolution
    │   │   │   ├── onboarding/
    │   │   │   │   ├── biometric/
    │   │   │   │   │   └── setup.tsx  # Biometric auth setup
    │   │   │   │   │       # - Face ID
    │   │   │   │   │       # - Touch ID
    │   │   │   │   │       # BE: users-be/auth
    │   │   │   │   │       # POST /v1/auth/biometric/enable
    │   │   │   │   ├── intro/
    │   │   │   │   │   └── index.tsx  # App intro screens
    │   │   │   │   │       # - Swipeable intro
    │   │   │   │   │       # - Feature highlights
    │   │   │   │   ├── permissions/
    │   │   │   │   │   ├── camera.tsx  # Camera permission
    │   │   │   │   │   ├── index.tsx  # All permissions
    │   │   │   │   │   ├── location.tsx  # Location permission
    │   │   │   │   │   └── notifications.tsx  # Notification permission
    │   │   │   │   ├── profile-setup/
    │   │   │   │   │   ├── basic-info.tsx  # Basic information
    │   │   │   │   │   │   # BE: users-be/profile
    │   │   │   │   │   │   # POST /v1/onboarding/profile
    │   │   │   │   │   ├── photo.tsx  # Profile photo
    │   │   │   │   │   │   # BE: storage-be/asset, users-be/profile
    │   │   │   │   │   │   # POST /v1/storage/uploads
    │   │   │   │   │   │   # PUT /v1/users/me/photo
    │   │   │   │   │   ├── preferences.tsx  # Initial preferences
    │   │   │   │   │   │   # BE: users-be/preferences
    │   │   │   │   │   │   # POST /v1/onboarding/preferences
    │   │   │   │   │   └── skills.tsx  # Skills selection
    │   │   │   │   │       # BE: users-be/skills
    │   │   │   │   │       # POST /v1/onboarding/skills
    │   │   │   │   ├── tutorial/
    │   │   │   │   │   ├── apply.tsx  # Tutorial: Apply
    │   │   │   │   │   ├── browse-jobs.tsx  # Tutorial: Browse jobs
    │   │   │   │   │   ├── index.tsx  # Interactive tutorial
    │   │   │   │   │   └── time-tracking.tsx  # Tutorial: Time tracking
    │   │   │   │   ├── _layout.tsx  # Onboarding layout
    │   │   │   │   ├── complete.tsx  # Onboarding complete
    │   │   │   │   │   # - Success message
    │   │   │   │   │   # - Next steps
    │   │   │   │   │   # BE: users-be/onboarding
    │   │   │   │   │   # POST /v1/onboarding/complete
    │   │   │   │   ├── features.tsx  # Features showcase
    │   │   │   │   ├── notifications-setup.tsx  # Notification preferences
    │   │   │   │   │   # BE: communications-be/preferences
    │   │   │   │   │   # POST /v1/notifications/preferences
    │   │   │   │   ├── permissions.tsx  # Request permissions
    │   │   │   │   │   # - Notifications
    │   │   │   │   │   # - Camera
    │   │   │   │   │   # - Location
    │   │   │   │   │   # Permission requests
    │   │   │   │   │   # - Camera access
    │   │   │   │   │   # - Location (optional)
    │   │   │   │   │   # BE: None (device-level)
    │   │   │   │   ├── profile-setup.tsx  # Quick profile setup
    │   │   │   │   │   # BE: users-be/profile
    │   │   │   │   │   # POST /v1/users/me/profile
    │   │   │   │   ├── user-type.tsx  # Select user type
    │   │   │   │   │   # - Freelancer
    │   │   │   │   │   # - Client
    │   │   │   │   │   # BE: None (local state)
    │   │   │   │   ├── verification.tsx  # Identity verification
    │   │   │   │   │   # - Document upload
    │   │   │   │   │   # - Selfie verification
    │   │   │   │   │   # BE: admin-be/kyc
    │   │   │   │   │   # POST /v1/kyc/submit
    │   │   │   │   └── welcome.tsx  # Welcome screen
    │   │   │   │       # - App introduction
    │   │   │   │       # - Value propositions
    │   │   │   │       # BE: None (static)
    │   │   │   ├── profile/
    │   │   │   │   ├── edit/
    │   │   │   │   │   ├── experience.tsx  # Edit experience
    │   │   │   │   │   ├── index.tsx  # Edit profile
    │   │   │   │   │   ├── portfolio.tsx  # Edit portfolio
    │   │   │   │   │   └── skills.tsx  # Edit skills
    │   │   │   │   └── [userId].tsx  # User profile (public view)
    │   │   │   │       # BE: users-be/profile
    │   │   │   │       # GET /v1/users/{user_id}/profile
    │   │   │   ├── proposals/
    │   │   │   │   ├── submit/
    │   │   │   │   │   └── [jobId].tsx  # Submit proposal
    │   │   │   │   │       # BE: proposals-be
    │   │   │   │   │       # POST /v1/proposals
    │   │   │   │   └── [id].tsx  # Proposal detail
    │   │   │   │       # BE: proposals-be
    │   │   │   │       # GET /v1/proposals/{proposal_id}
    │   │   │   ├── quick-actions/
    │   │   │   │   ├── quick-apply/
    │   │   │   │   │   └── [jobId].tsx  # Quick apply to job
    │   │   │   │   │       # - Pre-filled proposal
    │   │   │   │   │       # - One-tap apply
    │   │   │   │   │       # - Connect deduction
    │   │   │   │   │       # BE: proposals-be/proposal
    │   │   │   │   │       # POST /v1/proposals/quick-apply
    │   │   │   │   ├── quick-message/
    │   │   │   │   │   └── [userId].tsx  # Quick message to user
    │   │   │   │   │       # - Template messages
    │   │   │   │   │       # - Voice-to-text
    │   │   │   │   │       # BE: communications-be/conversation
    │   │   │   │   │       # POST /v1/conversations/{conversation_id}/messages/quick
    │   │   │   │   ├── quick-time-entry/
    │   │   │   │   │   └── [contractId].tsx  # Quick time logging
    │   │   │   │   │       # - Timer widget
    │   │   │   │   │       # - Quick notes
    │   │   │   │   │       # - One-tap submit
    │   │   │   │   │       # BE: contracts-be/work_diary
    │   │   │   │   │       # POST /v1/work-diary/quick-entry
    │   │   │   │   ├── quick-apply.tsx  # Quick proposal submission
    │   │   │   │   │   # - Minimal form
    │   │   │   │   │   # - Draft save
    │   │   │   │   │   # BE: proposals-be/proposal
    │   │   │   │   │   # POST /v1/proposals/quick
    │   │   │   │   │   # Quick job application
    │   │   │   │   │   # - Saved templates
    │   │   │   │   │   # - Quick submit
    │   │   │   │   │   # POST /v1/proposals/quick-apply
    │   │   │   │   ├── quick-invoice.tsx  # Quick invoice creation
    │   │   │   │   │   # - Templates
    │   │   │   │   │   # - Recent clients
    │   │   │   │   │   # - Quick send
    │   │   │   │   │   # BE: financial-be/invoice
    │   │   │   │   │   # POST /v1/invoices/quick-create
    │   │   │   │   ├── quick-message.tsx  # Quick message
    │   │   │   │   │   # - Contact selection
    │   │   │   │   │   # - Quick send
    │   │   │   │   │   # BE: communications-be/message
    │   │   │   │   │   # POST /v1/messages
    │   │   │   │   │   # Quick messaging
    │   │   │   │   │   # - Contact from anywhere
    │   │   │   │   │   # - Voice messages
    │   │   │   │   │   # - Quick replies
    │   │   │   │   │   # BE: communications-be/messages
    │   │   │   │   │   # POST /v1/messages/quick-send
    │   │   │   │   └── quick-time-entry.tsx  # Quick time logging
    │   │   │   │       # - Current contract
    │   │   │   │       # - Duration input
    │   │   │   │       # BE: contracts-be/workdiary
    │   │   │   │       # POST /v1/contracts/{contract_id}/workdiary/quick
    │   │   │   │       # - One-tap timer
    │   │   │   │       # - Recent tasks
    │   │   │   │       # - Auto-fill
    │   │   │   │       # POST /v1/contracts/time-entry/quick
    │   │   │   ├── reviews/
    │   │   │   │   ├── create/
    │   │   │   │   │   └── [contractId].tsx  # Create review
    │   │   │   │   │       # BE: reviews-be/reviews
    │   │   │   │   │       # POST /v1/reviews
    │   │   │   │   └── index.tsx  # Reviews list
    │   │   │   │       # BE: reviews-be/reviews
    │   │   │   │       # GET /v1/reviews
    │   │   │   ├── scanner/
    │   │   │   │   ├── document.tsx  # Document scanner
    │   │   │   │   │   # - Scan compliance docs
    │   │   │   │   │   # - OCR processing
    │   │   │   │   │   # BE: storage-be/asset, admin-be/business_verification
    │   │   │   │   └── qr-code.tsx  # QR code scanner
    │   │   │   │       # - Event check-in
    │   │   │   │       # - Profile sharing
    │   │   │   ├── search/
    │   │   │   │   ├── alerts/
    │   │   │   │   │   └── index.tsx  # Saved search alerts (mobile)
    │   │   │   │   │       # - Create/manage alerts
    │   │   │   │   │       # BE: search-be/alert (proposed if missing)
    │   │   │   │   │       # GET/POST/DELETE /v1/search/alerts
    │   │   │   │   └── saved/
    │   │   │   │       └── index.tsx  # Saved searches (mobile)
    │   │   │   │           # - List/run/delete saved queries
    │   │   │   │           # BE: search-be/saved_query (proposed if missing)
    │   │   │   │           # GET/POST/DELETE /v1/search/saved
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
    │   │   │   │   ├── app-settings/
    │   │   │   │   │   ├── appearance.tsx  # Appearance settings
    │   │   │   │   │   │   # - Theme (light/dark/auto)
    │   │   │   │   │   │   # - Font size
    │   │   │   │   │   │   # - Display density
    │   │   │   │   │   │   # BE: None (local storage)
    │   │   │   │   │   ├── biometric-auth.tsx  # Biometric authentication
    │   │   │   │   │   │   # - Face ID
    │   │   │   │   │   │   # - Fingerprint
    │   │   │   │   │   │   # - Setup/disable
    │   │   │   │   │   │   # BE: None (device-level)
    │   │   │   │   │   ├── haptics.tsx  # Haptic feedback
    │   │   │   │   │   │   # - Enable/disable
    │   │   │   │   │   │   # - Intensity
    │   │   │   │   │   │   # BE: None (local storage)
    │   │   │   │   │   ├── offline-mode.tsx  # Offline settings
    │   │   │   │   │   │   # - Auto-sync
    │   │   │   │   │   │   # - Storage limits
    │   │   │   │   │   │   # - Download quality
    │   │   │   │   │   │   # BE: None (local storage)
    │   │   │   │   │   └── quick-actions.tsx  # Quick action config
    │   │   │   │   │       # - Customize quick actions
    │   │   │   │   │       # - Shortcuts
    │   │   │   │   │       # BE: None (local storage)
    │   │   │   │   ├── biometric/
    │   │   │   │   │   └── index.tsx  # Biometric authentication
    │   │   │   │   │       # - Face ID / Touch ID
    │   │   │   │   │       # - Setup
    │   │   │   │   │       # BE: users-be/auth (device trust)
    │   │   │   │   │       # POST /v1/auth/device-trust
    │   │   │   │   ├── mobile-specific/
    │   │   │   │   │   ├── battery-optimization.tsx  # Battery settings
    │   │   │   │   │   │   # - Power saving mode
    │   │   │   │   │   │   # - Background sync
    │   │   │   │   │   │   # BE: None (local storage)
    │   │   │   │   │   ├── cache-management.tsx  # Cache management
    │   │   │   │   │   │   # - Clear cache
    │   │   │   │   │   │   # - Cache size
    │   │   │   │   │   │   # BE: None (local storage)
    │   │   │   │   │   └── data-usage.tsx  # Data usage settings
    │   │   │   │   │       # - WiFi only mode
    │   │   │   │   │       # - Image quality
    │   │   │   │   │       # - Video autoplay
    │   │   │   │   │       # BE: None (local storage)
    │   │   │   │   ├── privacy/
    │   │   │   │   │   └── gdpr/
    │   │   │   │   │       └── index.tsx  # GDPR (mobile)
    │   │   │   │   │           # - Request data export
    │   │   │   │   │           # - Request account erasure
    │   │   │   │   │           # - Track request status
    │   │   │   │   │           # BE: users-be/privacy
    │   │   │   │   │           # POST /v1/privacy/export
    │   │   │   │   │           # POST /v1/privacy/erase
    │   │   │   │   │           # GET  /v1/privacy/requests/{id}
    │   │   │   │   ├── storage/
    │   │   │   │   │   ├── cache/
    │   │   │   │   │   │   └── index.tsx  # Cache management
    │   │   │   │   │   │       # - Clear cache
    │   │   │   │   │   │       # - Cache size
    │   │   │   │   │   └── downloads/
    │   │   │   │   │       └── index.tsx  # Downloaded files
    │   │   │   │   │           # - Manage downloads
    │   │   │   │   │           # - Clear downloads
    │   │   │   │   ├── about.tsx  # About & support
    │   │   │   │   ├── account.tsx  # Account settings
    │   │   │   │   ├── index.tsx  # Settings menu
    │   │   │   │   ├── notifications.tsx  # Notification settings
    │   │   │   │   ├── privacy.tsx  # Privacy settings
    │   │   │   │   └── security.tsx  # Security settings
    │   │   │   ├── subscription/
    │   │   │   │   ├── connects.tsx  # Connects management
    │   │   │   │   │   # BE: subscriptions-be
    │   │   │   │   ├── index.tsx  # Subscription overview
    │   │   │   │   ├── plans.tsx  # Available plans
    │   │   │   │   └── upgrade.tsx  # Upgrade plan
    │   │   │   ├── support/
    │   │   │   │   └── tickets/
    │   │   │   │       └── index.tsx  # Support tickets (mobile)
    │   │   │   │           # - Create ticket, reply with attachments
    │   │   │   │           # - View ticket list and details
    │   │   │   │           # BE: admin-be/support_ticket
    │   │   │   │           # GET  /v1/support/tickets (mine)
    │   │   │   │           # POST /v1/support/tickets
    │   │   │   │           # POST /v1/support/tickets/{id}/messages
    │   │   │   │           # BE: storage-be/asset (uploads)
    │   │   │   │           # POST /v1/storage/uploads (signed URL) → PUT file → POST /v1/storage/commit
    │   │   │   ├── widgets/
    │   │   │   │   ├── quick-actions.tsx  # Quick actions widget
    │   │   │   │   │   # - Quick message
    │   │   │   │   │   # - Quick proposal
    │   │   │   │   └── time-tracker.tsx  # Home screen time tracker widget
    │   │   │   │       # BE: contracts-be/work_diary
    │   │   │   ├── +not-found.tsx  # 404 screen
    │   │   │   ├── _layout.tsx  # Root layout
    │   │   │   │   # - Auth provider
    │   │   │   │   # - Query client provider
    │   │   │   │   # - Theme provider
    │   │   │   │   # - Error boundary
    │   │   │   └── index.tsx  # App entry point
    │   │   │       # - Splash screen
    │   │   │       # - Initial route determination
    │   │   ├── assets/  # Mobile assets
    │   │   │   ├── fonts/  # Custom fonts
    │   │   │   ├── icons/  # App icons
    │   │   │   ├── images/  # Images
    │   │   │   └── splash/  # Splash screens
    │   │   ├── offline/
    │   │   │   ├── queue.tsx  # Offline actions queue
    │   │   │   │   # - Pending uploads
    │   │   │   │   # - Queued messages
    │   │   │   │   # - Draft proposals
    │   │   │   └── sync.tsx  # Sync status
    │   │   │       # - Sync progress
    │   │   │       # - Conflict resolution
    │   │   ├── scanner/
    │   │   │   ├── document.tsx  # Document scanner
    │   │   │   │   # - Scan compliance docs
    │   │   │   │   # - OCR processing
    │   │   │   │   # BE: storage-be/asset, admin-be/business_verification
    │   │   │   └── qr-code.tsx  # QR code scanner
    │   │   │       # - Event check-in
    │   │   │       # - Profile sharing
    │   │   ├── src/
    │   │   │   ├── components/  # Mobile-specific components
    │   │   │   │   ├── Auth/
    │   │   │   │   │   ├── BiometricButton.tsx  # Biometric auth button
    │   │   │   │   │   ├── LoginForm.tsx  # Login form component
    │   │   │   │   │   ├── RegisterForm.tsx  # Registration form
    │   │   │   │   │   └── SocialButtons.tsx  # Social login buttons
    │   │   │   │   ├── Common/
    │   │   │   │   │   ├── EmptyState.tsx  # Empty state component
    │   │   │   │   │   ├── ErrorBoundary.tsx  # Error boundary
    │   │   │   │   │   ├── Loading.tsx  # Loading spinner
    │   │   │   │   │   ├── OptimizedFlashList.tsx  # Optimized list (FlashList)
    │   │   │   │   │   └── PullToRefresh.tsx  # Pull to refresh
    │   │   │   │   ├── Contracts/
    │   │   │   │   │   ├── ContractCard.tsx  # Contract card
    │   │   │   │   │   ├── MilestoneItem.tsx  # Milestone list item
    │   │   │   │   │   └── TimesheetEntry.tsx  # Timesheet entry
    │   │   │   │   ├── Financial/
    │   │   │   │   │   ├── InvoiceCard.tsx  # Invoice card
    │   │   │   │   │   ├── TransactionItem.tsx  # Transaction list item
    │   │   │   │   │   └── WalletCard.tsx  # Wallet balance card
    │   │   │   │   ├── Jobs/
    │   │   │   │   │   ├── JobCard.tsx  # Job card component
    │   │   │   │   │   ├── JobDetail.tsx  # Job detail view
    │   │   │   │   │   ├── JobFilters.tsx  # Job filters bottom sheet
    │   │   │   │   │   └── JobList.tsx  # Job list
    │   │   │   │   ├── Messages/
    │   │   │   │   │   ├── ConversationCard.tsx  # Conversation list item
    │   │   │   │   │   ├── MessageBubble.tsx  # Message bubble
    │   │   │   │   │   ├── MessageComposer.tsx  # Message input
    │   │   │   │   │   └── TypingIndicator.tsx  # Typing indicator
    │   │   │   │   ├── Navigation/
    │   │   │   │   │   ├── Header.tsx  # Screen header
    │   │   │   │   │   └── TabBar.tsx  # Custom tab bar
    │   │   │   │   ├── Profile/
    │   │   │   │   │   ├── ExperienceItem.tsx  # Experience item
    │   │   │   │   │   ├── PortfolioItem.tsx  # Portfolio item
    │   │   │   │   │   ├── ProfileHeader.tsx  # Profile header
    │   │   │   │   │   └── SkillTag.tsx  # Skill tag
    │   │   │   │   ├── Proposals/
    │   │   │   │   │   ├── ProposalCard.tsx  # Proposal card
    │   │   │   │   │   ├── ProposalForm.tsx  # Proposal submission form
    │   │   │   │   │   └── ProposalList.tsx  # Proposal list
    │   │   │   │   └── UI/
    │   │   │   │       ├── Avatar.tsx  # Avatar component
    │   │   │   │       ├── Badge.tsx  # Badge component
    │   │   │   │       ├── BottomSheet.tsx  # Bottom sheet modal
    │   │   │   │       ├── Button.tsx  # Button component
    │   │   │   │       ├── Card.tsx  # Card component
    │   │   │   │       ├── Input.tsx  # Input component
    │   │   │   │       └── SearchBar.tsx  # Search bar
    │   │   │   ├── hooks/  # Mobile-specific hooks
    │   │   │   │   ├── useAppState.ts  # App state (foreground/background)
    │   │   │   │   ├── useBiometricAuth.ts  # Biometric authentication
    │   │   │   │   ├── useCamera.ts  # Camera access
    │   │   │   │   ├── useHighFPSAnimation.ts  # High FPS animations
    │   │   │   │   ├── useKeyboard.ts  # Keyboard handling
    │   │   │   │   ├── useLocation.ts  # Geolocation
    │   │   │   │   ├── useNetworkStatus.ts  # Network connectivity
    │   │   │   │   ├── useOrientation.ts  # Device orientation
    │   │   │   │   └── usePushNotifications.ts  # Push notifications
    │   │   │   ├── lib/
    │   │   │   │   ├── i18n/
    │   │   │   │   │   └── index.ts  # i18n configuration (mobile)
    │   │   │   │   │       # Uses packages/shared i18n resources
    │   │   │   │   ├── analytics.ts  # Mobile analytics
    │   │   │   │   │   # - Event tracking
    │   │   │   │   │   # - Screen tracking
    │   │   │   │   ├── deeplink.ts  # Deep link handling
    │   │   │   │   ├── error-tracking.ts  # Error tracking (Sentry)
    │   │   │   │   ├── keycloak-mobile.ts  # Keycloak mobile config
    │   │   │   │   │   # - OAuth2 with PKCE
    │   │   │   │   │   # - Token storage (SecureStore)
    │   │   │   │   │   # - Refresh token flow
    │   │   │   │   ├── performance.ts  # Performance utilities
    │   │   │   │   │   # - FPS monitoring
    │   │   │   │   │   # - Memory optimization
    │   │   │   │   ├── push-notifications.ts  # Push notification setup
    │   │   │   │   │   # - Firebase/Expo notifications
    │   │   │   │   │   # - Token registration
    │   │   │   │   │   # BE: communications-be/push
    │   │   │   │   │   # POST /v1/push-tokens
    │   │   │   │   ├── storage.ts  # Secure storage wrapper
    │   │   │   │   │   # - Token storage
    │   │   │   │   │   # - Biometric keys
    │   │   │   │   └── utils.ts  # General utilities
    │   │   │   ├── stores/  # Mobile-specific Zustand stores
    │   │   │   │   ├── biometric-store.ts  # Biometric settings
    │   │   │   │   ├── camera-store.ts  # Camera state
    │   │   │   │   └── offline-queue-store.ts  # Offline action queue
    │   │   │   └── types/  # Mobile-specific types
    │   │   │       ├── biometric.ts  # Biometric types
    │   │   │       └── navigation.ts  # Navigation types
    │   │   ├── widgets/
    │   │   │   ├── quick-actions.tsx  # Quick actions widget
    │   │   │   │   # - Quick message
    │   │   │   │   # - Quick proposal
    │   │   │   └── time-tracker.tsx  # Home screen time tracker widget
    │   │   │       # BE: contracts-be/work_diary
    │   │   ├── .env  # Environment variables
    │   │   ├── .eslintrc.json  # ESLint config
    │   │   ├── app.json  # Expo config
    │   │   │   # - App name, slug
    │   │   │   # - Icons, splash screens
    │   │   │   # - Permissions
    │   │   │   # - Build settings (EAS)
    │   │   ├── babel.config.js  # Babel config
    │   │   ├── eas.json  # EAS Build config
    │   │   │   # - Development build
    │   │   │   # - Preview build
    │   │   │   # - Production build
    │   │   │   # - Credentials
    │   │   ├── global.css  # Global styles (NativeWind)
    │   │   ├── index.js  # Entry point
    │   │   ├── metro.config.js  # Metro bundler config
    │   │   │   # - Monorepo support
    │   │   │   # - Asset resolution
    │   │   ├── package.json  # Mobile dependencies
    │   │   │   # - Expo SDK
    │   │   │   # - React Native
    │   │   │   # - Expo Router
    │   │   │   # - NativeWind
    │   │   │   # - TanStack Query
    │   │   │   # - Zustand
    │   │   ├── tailwind.config.js  # Tailwind config (NativeWind)
    │   │   ├── tsconfig.json  # TypeScript config
    │   │   └── │
    │   ├── web/  # Next.js web application
    │   │   ├── e2e/
    │   │   │   ├── auth/
    │   │   │   │   ├── login.spec.ts
    │   │   │   │   ├── password-reset.spec.ts
    │   │   │   │   └── register.spec.ts
    │   │   │   ├── config/
    │   │   │   │   ├── playwright.config.ts
    │   │   │   │   └── test-ids.ts
    │   │   │   ├── contracts/
    │   │   │   │   ├── contract-creation.spec.ts
    │   │   │   │   └── contract-execution.spec.ts
    │   │   │   ├── fixtures/
    │   │   │   │   ├── contracts.json
    │   │   │   │   ├── jobs.json
    │   │   │   │   ├── test-data.ts
    │   │   │   │   ├── test-jobs.json
    │   │   │   │   ├── test-users.json
    │   │   │   │   └── users.json
    │   │   │   ├── helpers/
    │   │   │   │   ├── assertions.ts
    │   │   │   │   ├── auth.ts
    │   │   │   │   └── navigation.ts
    │   │   │   ├── jobs/
    │   │   │   │   ├── job-application.spec.ts
    │   │   │   │   ├── job-posting.spec.ts
    │   │   │   │   └── job-search.spec.ts
    │   │   │   ├── page-objects/
    │   │   │   │   ├── auth/
    │   │   │   │   │   ├── LoginPage.ts
    │   │   │   │   │   └── RegisterPage.ts
    │   │   │   │   ├── dashboard/
    │   │   │   │   │   └── DashboardPage.ts
    │   │   │   │   └── jobs/
    │   │   │   │       ├── JobDetailPage.ts
    │   │   │   │       └── JobListPage.ts
    │   │   │   ├── payments/
    │   │   │   │   ├── escrow.spec.ts
    │   │   │   │   └── payment-methods.spec.ts
    │   │   │   ├── proposals/
    │   │   │   │   ├── proposal-review.spec.ts
    │   │   │   │   └── proposal-submission.spec.ts
    │   │   │   ├── tests/
    │   │   │   │   ├── auth/
    │   │   │   │   │   ├── login.spec.ts
    │   │   │   │   │   ├── oauth.spec.ts
    │   │   │   │   │   ├── password-reset.spec.ts
    │   │   │   │   │   ├── register.spec.ts
    │   │   │   │   │   └── registration.spec.ts
    │   │   │   │   ├── contracts/
    │   │   │   │   │   ├── contract-creation.spec.ts
    │   │   │   │   │   ├── create-contract.spec.ts
    │   │   │   │   │   ├── deliverable-submission.spec.ts
    │   │   │   │   │   ├── dispute.spec.ts
    │   │   │   │   │   ├── milestone-tracking.spec.ts
    │   │   │   │   │   ├── submit-deliverable.spec.ts
    │   │   │   │   │   └── time-tracking.spec.ts
    │   │   │   │   ├── jobs/
    │   │   │   │   │   ├── apply-to-job.spec.ts
    │   │   │   │   │   ├── job-application.spec.ts
    │   │   │   │   │   ├── job-posting.spec.ts
    │   │   │   │   │   ├── job-search.spec.ts
    │   │   │   │   │   ├── manage-jobs.spec.ts
    │   │   │   │   │   ├── post-job.spec.ts
    │   │   │   │   │   └── search-jobs.spec.ts
    │   │   │   │   ├── messaging/
    │   │   │   │   │   ├── notifications.spec.ts
    │   │   │   │   │   ├── real-time-chat.spec.ts
    │   │   │   │   │   └── send-message.spec.ts
    │   │   │   │   ├── payments/
    │   │   │   │   │   ├── add-payment-method.spec.ts
    │   │   │   │   │   ├── escrow-funding.spec.ts
    │   │   │   │   │   ├── invoice-generation.spec.ts
    │   │   │   │   │   ├── payment-methods.spec.ts
    │   │   │   │   │   ├── payout-processing.spec.ts
    │   │   │   │   │   ├── release-payment.spec.ts
    │   │   │   │   │   └── withdraw.spec.ts
    │   │   │   │   └── proposals/
    │   │   │   │       ├── create-proposal.spec.ts
    │   │   │   │       ├── edit-proposal.spec.ts
    │   │   │   │       ├── proposal-acceptance.spec.ts
    │   │   │   │       ├── proposal-negotiation.spec.ts
    │   │   │   │       ├── proposal-submission.spec.ts
    │   │   │   │       └── withdraw-proposal.spec.ts
    │   │   │   └── playwright.config.ts
    │   │   ├── lib/
    │   │   │   ├── image-optimization/
    │   │   │   │   ├── loader.ts  # Custom image loader
    │   │   │   │   ├── placeholder.ts  # Blur placeholders
    │   │   │   │   └── responsive.ts  # Responsive images
    │   │   │   └── security/
    │   │   │       ├── cors.ts  # CORS configuration
    │   │   │       ├── csp.ts  # Content Security Policy
    │   │   │       ├── headers.ts  # Security headers
    │   │   │       └── rate-limiter.ts  # Rate limiting
    │   │   ├── security/
    │   │   │   ├── cors/
    │   │   │   │   ├── cors-config.ts  # CORS configuration
    │   │   │   │   └── origin-validation.ts  # Origin validation
    │   │   │   ├── headers/
    │   │   │   │   ├── csp.ts  # Content Security Policy
    │   │   │   │   ├── hsts.ts  # HTTP Strict Transport Security
    │   │   │   │   ├── permissions-policy.ts  # Permissions Policy
    │   │   │   │   └── x-frame-options.ts  # X-Frame-Options
    │   │   │   ├── rate-limiting/
    │   │   │   │   ├── ddos-protection.ts  # DDoS protection
    │   │   │   │   └── rate-limiter.ts  # Rate limiting
    │   │   │   ├── auth-guard.ts
    │   │   │   ├── cors.ts  # CORS configuration
    │   │   │   ├── csp.ts  # Content Security Policy
    │   │   │   ├── csrf-protection.ts
    │   │   │   ├── headers.ts  # Security headers config
    │   │   │   └── rate-limiter.ts
    │   │   ├── src/
    │   │   │   ├── app/
    │   │   │   │   ├── (admin)/
    │   │   │   │   │   ├── currency/
    │   │   │   │   │   │   ├── conversions/
    │   │   │   │   │   │   │   └── page.tsx  # Currency conversions
    │   │   │   │   │   │   │       # - Conversion history
    │   │   │   │   │   │   │       # - Spread analysis
    │   │   │   │   │   │   │       # BE: financial-be/currency
    │   │   │   │   │   │   │       # GET /v1/admin/currency/conversions
    │   │   │   │   │   │   ├── hedging/
    │   │   │   │   │   │   │   └── page.tsx  # Currency hedging
    │   │   │   │   │   │   │       # BE: financial-be/currency
    │   │   │   │   │   │   │       # GET /v1/admin/currency/hedging
    │   │   │   │   │   │   └── rates/
    │   │   │   │   │   │       ├── manual-override/
    │   │   │   │   │   │       │   └── page.tsx  # Manual rate override
    │   │   │   │   │   │       │       # BE: financial-be/currency
    │   │   │   │   │   │       │       # POST /v1/admin/currency/rates/override
    │   │   │   │   │   │       └── page.tsx  # Exchange rates
    │   │   │   │   │   │           # BE: financial-be/currency
    │   │   │   │   │   │           # GET /v1/admin/currency/rates
    │   │   │   │   │   ├── fees/
    │   │   │   │   │   │   ├── overrides/
    │   │   │   │   │   │   │   └── page.tsx  # Fee overrides
    │   │   │   │   │   │   │       # - Custom fee arrangements
    │   │   │   │   │   │   │       # - Volume discounts
    │   │   │   │   │   │   │       # BE: financial-be/fees
    │   │   │   │   │   │   │       # GET /v1/admin/fees/overrides
    │   │   │   │   │   │   │       # POST /v1/admin/fees/overrides
    │   │   │   │   │   │   ├── promotions/
    │   │   │   │   │   │   │   ├── [promotionId]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Promotion detail
    │   │   │   │   │   │   │   │       # BE: financial-be/fees
    │   │   │   │   │   │   │   │       # GET /v1/admin/fees/promotions/{promotion_id}
    │   │   │   │   │   │   │   └── page.tsx  # Fee promotions
    │   │   │   │   │   │   │       # BE: financial-be/fees
    │   │   │   │   │   │   │       # GET /v1/admin/fees/promotions
    │   │   │   │   │   │   │       # POST /v1/admin/fees/promotions
    │   │   │   │   │   │   └── structures/
    │   │   │   │   │   │       ├── [structureId]/
    │   │   │   │   │   │       │   └── page.tsx  # Fee structure detail
    │   │   │   │   │   │       │       # BE: financial-be/fees
    │   │   │   │   │   │       │       # GET /v1/admin/fees/structures/{structure_id}
    │   │   │   │   │   │       └── page.tsx  # Fee structures
    │   │   │   │   │   │           # BE: financial-be/fees
    │   │   │   │   │   │           # GET /v1/admin/fees/structures
    │   │   │   │   │   │           # POST /v1/admin/fees/structures
    │   │   │   │   │   └── financial/
    │   │   │   │   │       ├── fraud/
    │   │   │   │   │       │   ├── cases/
    │   │   │   │   │       │   │   ├── [caseId]/
    │   │   │   │   │       │   │   │   ├── evidence/
    │   │   │   │   │       │   │   │   │   └── page.tsx  # Case evidence
    │   │   │   │   │       │   │   │   │       # BE: financial-be/fraud, storage-be
    │   │   │   │   │       │   │   │   │       # GET /v1/admin/fraud/cases/{case_id}/evidence
    │   │   │   │   │       │   │   │   ├── investigation/
    │   │   │   │   │       │   │   │   │   └── page.tsx  # Case investigation
    │   │   │   │   │       │   │   │   │       # BE: financial-be/fraud
    │   │   │   │   │       │   │   │   │       # GET /v1/admin/fraud/cases/{case_id}/investigation
    │   │   │   │   │       │   │   │   └── resolution/
    │   │   │   │   │       │   │   │       └── page.tsx  # Case resolution
    │   │   │   │   │       │   │   │           # BE: financial-be/fraud
    │   │   │   │   │       │   │   │           # POST /v1/admin/fraud/cases/{case_id}/resolve
    │   │   │   │   │       │   │   └── page.tsx  # Fraud cases
    │   │   │   │   │       │   │       # BE: financial-be/fraud
    │   │   │   │   │       │   │       # GET /v1/admin/fraud/cases
    │   │   │   │   │       │   ├── detection/
    │   │   │   │   │       │   │   ├── alerts/
    │   │   │   │   │       │   │   │   ├── [alertId]/
    │   │   │   │   │       │   │   │   │   └── page.tsx  # Alert investigation
    │   │   │   │   │       │   │   │   │       # BE: financial-be/fraud
    │   │   │   │   │       │   │   │   │       # GET /v1/admin/fraud/alerts/{alert_id}
    │   │   │   │   │       │   │   │   │       # POST /v1/admin/fraud/alerts/{alert_id}/investigate
    │   │   │   │   │       │   │   │   └── page.tsx  # Fraud alerts
    │   │   │   │   │       │   │   │       # BE: financial-be/fraud
    │   │   │   │   │       │   │   │       # GET /v1/admin/fraud/alerts
    │   │   │   │   │       │   │   ├── ml-models/
    │   │   │   │   │       │   │   │   └── page.tsx  # ML fraud models
    │   │   │   │   │       │   │   │       # - Model performance
    │   │   │   │   │       │   │   │       # - Model tuning
    │   │   │   │   │       │   │   │       # BE: financial-be/fraud
    │   │   │   │   │       │   │   │       # GET /v1/admin/fraud/ml-models
    │   │   │   │   │       │   │   └── rules/
    │   │   │   │   │       │   │       └── page.tsx  # Fraud detection rules
    │   │   │   │   │       │   │           # BE: financial-be/fraud
    │   │   │   │   │       │   │           # GET /v1/admin/fraud/rules
    │   │   │   │   │       │   │           # POST /v1/admin/fraud/rules
    │   │   │   │   │       │   ├── patterns/
    │   │   │   │   │       │   └── page.tsx  # Fraud pattern analysis
    │   │   │   │   │       │       # BE: financial-be/fraud
    │   │   │   │   │       │       # GET /v1/admin/fraud/patterns
    │   │   │   │   │       ├── reserves/
    │   │   │   │   │       │   ├── adjustments/
    │   │   │   │   │       │   │   └── page.tsx  # Reserve adjustments
    │   │   │   │   │       │   │       # BE: financial-be/reserves
    │   │   │   │   │       │   │       # GET /v1/admin/reserves/adjustments
    │   │   │   │   │       │   │       # POST /v1/admin/reserves/adjustments
    │   │   │   │   │       │   ├── calculation/
    │   │   │   │   │       │   │   └── page.tsx  # Reserve calculation
    │   │   │   │   │       │   │       # - Reserve requirements
    │   │   │   │   │       │   │       # - Rolling reserves
    │   │   │   │   │       │   │       # BE: financial-be/reserves
    │   │   │   │   │       │   │       # GET /v1/admin/reserves/calculation
    │   │   │   │   │       │   └── releases/
    │   │   │   │   │       │       └── page.tsx  # Reserve releases
    │   │   │   │   │       │           # BE: financial-be/reserves
    │   │   │   │   │       │           # GET /v1/admin/reserves/releases
    │   │   │   │   │       └── settlement/
    │   │   │   │   │           ├── batches/
    │   │   │   │   │           │   ├── [batchId]/
    │   │   │   │   │           │   │   ├── approve/
    │   │   │   │   │           │   │   │   └── page.tsx  # Approve settlement
    │   │   │   │   │           │   │   │       # BE: financial-be/settlement, admin-be/change-approval
    │   │   │   │   │           │   │   │       # POST /v1/admin/settlement/batches/{batch_id}/approve
    │   │   │   │   │           │   │   ├── review/
    │   │   │   │   │           │   │   │   └── page.tsx  # Review settlement batch
    │   │   │   │   │           │   │   │       # BE: financial-be/settlement
    │   │   │   │   │           │   │   │       # GET /v1/admin/settlement/batches/{batch_id}
    │   │   │   │   │           │   │   └── page.tsx  # Batch detail
    │   │   │   │   │           │   │       # BE: financial-be/settlement
    │   │   │   │   │           │   │       # GET /v1/admin/settlement/batches/{batch_id}
    │   │   │   │   │           │   └── page.tsx  # Settlement batches
    │   │   │   │   │           │       # BE: financial-be/settlement
    │   │   │   │   │           │       # GET /v1/admin/settlement/batches
    │   │   │   │   │           ├── holds/
    │   │   │   │   │           │   ├── [holdId]/
    │   │   │   │   │           │   │   ├── release/
    │   │   │   │   │           │   │   │   └── page.tsx  # Release hold
    │   │   │   │   │           │   │   │       # BE: financial-be/settlement
    │   │   │   │   │           │   │   │       # POST /v1/admin/settlement/holds/{hold_id}/release
    │   │   │   │   │           │   │   └── page.tsx  # Hold detail
    │   │   │   │   │           │   │       # BE: financial-be/settlement
    │   │   │   │   │           │   │       # GET /v1/admin/settlement/holds/{hold_id}
    │   │   │   │   │           │   └── page.tsx  # Payment holds
    │   │   │   │   │           │       # BE: financial-be/settlement
    │   │   │   │   │           │       # GET /v1/admin/settlement/holds
    │   │   │   │   │           └── rules/
    │   │   │   │   │               └── page.tsx  # Settlement rules
    │   │   │   │   │                   # - Auto-hold rules
    │   │   │   │   │                   # - Risk thresholds
    │   │   │   │   │                   # BE: financial-be/settlement
    │   │   │   │   │                   # GET /v1/admin/settlement/rules
    │   │   │   │   │                   # PUT /v1/admin/settlement/rules
    │   │   │   │   ├── [locale]/  # Internationalized routing (en, ar, zh, hi, de, fr, tr, es, ru)
    │   │   │   │   │   ├── (admin)/
    │   │   │   │   │   │   ├── admin/
    │   │   │   │   │   │   │   ├── audit-logs/
    │   │   │   │   │   │   │   │   ├── [logId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Audit log details
    │   │   │   │   │   │   │   │   │       # BE: utility-be/audit
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/audit-logs/{log_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Audit logs viewer
    │   │   │   │   │   │   │   │       # - System-wide audit trail
    │   │   │   │   │   │   │   │       # - Filter by entity/action
    │   │   │   │   │   │   │   │       # - Export capabilities
    │   │   │   │   │   │   │   │       # BE: utility-be/audit
    │   │   │   │   │   │   │   │       # GET /v1/admin/audit-logs
    │   │   │   │   │   │   │   ├── break-glass/
    │   │   │   │   │   │   │   │   ├── active/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Active admin sessions
    │   │   │   │   │   │   │   │   │       # - Time-boxed access monitoring
    │   │   │   │   │   │   │   │   │       # - Force termination
    │   │   │   │   │   │   │   │   │       # BE: admin-be/admin_session
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/break-glass/active
    │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve break-glass requests
    │   │   │   │   │   │   │   │   │       # - Two-person rule
    │   │   │   │   │   │   │   │   │       # - Request review
    │   │   │   │   │   │   │   │   │       # BE: admin-be/admin_session
    │   │   │   │   │   │   │   │   │       # POST /v1/admin/break-glass/approve/{request_id}
    │   │   │   │   │   │   │   │   └── request/
    │   │   │   │   │   │   │   │       └── page.tsx  # Request break-glass access
    │   │   │   │   │   │   │   │           # - Justification
    │   │   │   │   │   │   │   │           # - Duration request
    │   │   │   │   │   │   │   │           # BE: admin-be/admin_session
    │   │   │   │   │   │   │   │           # POST /v1/admin/break-glass/request
    │   │   │   │   │   │   │   ├── business-verification/
    │   │   │   │   │   │   │   │   ├── [verificationId]/
    │   │   │   │   │   │   │   │   │   ├── reverify/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Request reverification
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/business_verification
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/business-verification/{verification_id}/reverify
    │   │   │   │   │   │   │   │   │   ├── review/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Review business verification
    │   │   │   │   │   │   │   │   │   │       # - Company evidence review
    │   │   │   │   │   │   │   │   │   │       # - Decision making
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/business_verification
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/business-verification/{verification_id}/review
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Verification details
    │   │   │   │   │   │   │   │   │       # BE: admin-be/business_verification
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/business-verification/{verification_id}
    │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending verifications
    │   │   │   │   │   │   │   │   │       # BE: admin-be/business_verification
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/business-verification?status=pending
    │   │   │   │   │   │   │   │   └── page.tsx  # Business verification dashboard
    │   │   │   │   │   │   │   │       # BE: admin-be/business_verification
    │   │   │   │   │   │   │   │       # GET /v1/admin/business-verification
    │   │   │   │   │   │   │   ├── communications-ops/
    │   │   │   │   │   │   │   │   ├── broadcasts/
    │   │   │   │   │   │   │   │   │   ├── [broadcastId]/
    │   │   │   │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Broadcast analytics
    │   │   │   │   │   │   │   │   │   │   │       # BE: communications-be/broadcast
    │   │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/broadcasts/{broadcast_id}/analytics
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit broadcast
    │   │   │   │   │   │   │   │   │   │   │       # BE: communications-be/broadcast
    │   │   │   │   │   │   │   │   │   │   │       # PUT /v1/admin/broadcasts/{broadcast_id}
    │   │   │   │   │   │   │   │   │   │   ├── schedule/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Schedule broadcast
    │   │   │   │   │   │   │   │   │   │   │       # BE: communications-be/broadcast
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/broadcasts/{broadcast_id}/schedule
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Broadcast details
    │   │   │   │   │   │   │   │   │   │       # BE: communications-be/broadcast
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/broadcasts/{broadcast_id}
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create broadcast
    │   │   │   │   │   │   │   │   │   │       # - Target audience
    │   │   │   │   │   │   │   │   │   │       # - Message content
    │   │   │   │   │   │   │   │   │   │       # - Delivery channels
    │   │   │   │   │   │   │   │   │   │       # BE: communications-be/broadcast
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/broadcasts
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Broadcasts dashboard
    │   │   │   │   │   │   │   │   │       # BE: communications-be/broadcast
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/broadcasts
    │   │   │   │   │   │   │   │   ├── campaigns/
    │   │   │   │   │   │   │   │   │   ├── [campaignId]/
    │   │   │   │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Campaign analytics
    │   │   │   │   │   │   │   │   │   │   │       # BE: communications-be/campaign
    │   │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/campaigns/{campaign_id}/analytics
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit campaign
    │   │   │   │   │   │   │   │   │   │   │       # BE: communications-be/campaign
    │   │   │   │   │   │   │   │   │   │   │       # PUT /v1/admin/campaigns/{campaign_id}
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Campaign details
    │   │   │   │   │   │   │   │   │   │       # BE: communications-be/campaign
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/campaigns/{campaign_id}
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create campaign
    │   │   │   │   │   │   │   │   │   │       # BE: communications-be/campaign
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/campaigns
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Campaigns dashboard
    │   │   │   │   │   │   │   │   │       # BE: communications-be/campaign
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/campaigns
    │   │   │   │   │   │   │   │   ├── rate-limits/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Communication rate limits
    │   │   │   │   │   │   │   │   │       # - Configure limits
    │   │   │   │   │   │   │   │   │       # - Monitor usage
    │   │   │   │   │   │   │   │   │       # BE: communications-be/rate_limit
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/rate-limits
    │   │   │   │   │   │   │   │   │       # PUT /v1/admin/rate-limits
    │   │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │   │       ├── [templateId]/
    │   │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit template
    │   │   │   │   │   │   │   │       │   │       # BE: communications-be/template
    │   │   │   │   │   │   │   │       │   │       # PUT /v1/admin/templates/{template_id}
    │   │   │   │   │   │   │   │       │   ├── logs/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Webhook delivery logs
    │   │   │   │   │   │   │   │       │   │       # BE: communications-be/template
    │   │   │   │   │   │   │   │       │   │       # GET /v1/developer/webhooks/{webhook_id}/logs
    │   │   │   │   │   │   │   │       │   ├── test/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Test template
    │   │   │   │   │   │   │   │       │   │       # BE: communications-be/template
    │   │   │   │   │   │   │   │       │   │       # POST /v1/admin/templates/{template_id}/test
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Template details
    │   │   │   │   │   │   │   │       │       # BE: communications-be/template
    │   │   │   │   │   │   │   │       │       # GET /v1/admin/templates/{template_id}
    │   │   │   │   │   │   │   │       ├── create/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Create template
    │   │   │   │   │   │   │   │       │       # BE: communications-be/template
    │   │   │   │   │   │   │   │       │       # POST /v1/admin/templates
    │   │   │   │   │   │   │   │       └── page.tsx  # Templates dashboard
    │   │   │   │   │   │   │   │           # BE: communications-be/template
    │   │   │   │   │   │   │   │           # GET /v1/admin/templates
    │   │   │   │   │   │   │   ├── dashboard/
    │   │   │   │   │   │   │   │   └── page.tsx  # Admin dashboard
    │   │   │   │   │   │   │   │       # - Key metrics
    │   │   │   │   │   │   │   │       # - Pending actions
    │   │   │   │   │   │   │   │       # - System alerts
    │   │   │   │   │   │   │   │       # BE: admin-be/dashboard
    │   │   │   │   │   │   │   │       # GET /v1/admin/dashboard
    │   │   │   │   │   │   │   ├── financial-ops/
    │   │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   │   │   │   ├── mediate/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Mediate payment dispute
    │   │   │   │   │   │   │   │   │   │   │       # BE: admin-be/dispute_resolution
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/financial-disputes/{dispute_id}/mediate
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute details
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/dispute_resolution
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/financial-disputes/{dispute_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Financial disputes
    │   │   │   │   │   │   │   │   │       # BE: admin-be/dispute_resolution
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/financial-disputes
    │   │   │   │   │   │   │   │   ├── payouts/
    │   │   │   │   │   │   │   │   │   ├── [payoutId]/
    │   │   │   │   │   │   │   │   │   │   ├── retry/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Retry failed payout
    │   │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payout
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/payouts/{payout_id}/retry
    │   │   │   │   │   │   │   │   │   │   ├── review/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Review payout
    │   │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payout
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/payouts/{payout_id}/review
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Payout details
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payout
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/payouts/{payout_id}
    │   │   │   │   │   │   │   │   │   ├── failed/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Failed payouts
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payout
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/payouts?status=failed
    │   │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending payouts
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payout
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/payouts?status=pending
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Payouts dashboard
    │   │   │   │   │   │   │   │   │       # BE: financial-be/payout
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/payouts
    │   │   │   │   │   │   │   │   ├── reconciliation/
    │   │   │   │   │   │   │   │   │   ├── [reconciliationId]/
    │   │   │   │   │   │   │   │   │   │   ├── resolve/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Resolve reconciliation
    │   │   │   │   │   │   │   │   │   │   │       # BE: financial-be/reconciliation
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/reconciliation/{reconciliation_id}/resolve
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reconciliation details
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/reconciliation
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/reconciliation/{reconciliation_id}
    │   │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending reconciliations
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/reconciliation
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/reconciliation?status=pending
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Reconciliation dashboard
    │   │   │   │   │   │   │   │   │       # BE: financial-be/reconciliation
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/reconciliation
    │   │   │   │   │   │   │   │   └── tax-forms/
    │   │   │   │   │   │   │   │       ├── [formId]/
    │   │   │   │   │   │   │   │       │   ├── review/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Review tax form
    │   │   │   │   │   │   │   │       │   │       # BE: financial-be/tax
    │   │   │   │   │   │   │   │       │   │       # POST /v1/admin/tax-forms/{form_id}/review
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Tax form details
    │   │   │   │   │   │   │   │       │       # BE: financial-be/tax
    │   │   │   │   │   │   │   │       │       # GET /v1/admin/tax-forms/{form_id}
    │   │   │   │   │   │   │   │       ├── generate/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Generate tax forms
    │   │   │   │   │   │   │   │       │       # - Bulk 1099 generation
    │   │   │   │   │   │   │   │       │       # - Tax year selection
    │   │   │   │   │   │   │   │       │       # BE: financial-be/tax
    │   │   │   │   │   │   │   │       │       # POST /v1/admin/tax-forms/generate
    │   │   │   │   │   │   │   │       └── page.tsx  # Tax forms dashboard
    │   │   │   │   │   │   │   │           # BE: financial-be/tax
    │   │   │   │   │   │   │   │           # GET /v1/admin/tax-forms
    │   │   │   │   │   │   │   ├── goodwill-credits/
    │   │   │   │   │   │   │   │   ├── [creditId]/
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve goodwill credit
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/goodwill_credit
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/goodwill-credits/{credit_id}/approve
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Goodwill credit details
    │   │   │   │   │   │   │   │   │       # BE: admin-be/goodwill_credit
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/goodwill-credits/{credit_id}
    │   │   │   │   │   │   │   │   ├── issue/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Issue goodwill credit
    │   │   │   │   │   │   │   │   │       # - User selection
    │   │   │   │   │   │   │   │   │       # - Amount and reason
    │   │   │   │   │   │   │   │   │       # BE: admin-be/goodwill_credit
    │   │   │   │   │   │   │   │   │       # POST /v1/admin/goodwill-credits
    │   │   │   │   │   │   │   │   └── page.tsx  # Goodwill credits dashboard
    │   │   │   │   │   │   │   │       # BE: admin-be/goodwill_credit
    │   │   │   │   │   │   │   │       # GET /v1/admin/goodwill-credits
    │   │   │   │   │   │   │   ├── kyc-cases/
    │   │   │   │   │   │   │   │   ├── [caseId]/
    │   │   │   │   │   │   │   │   │   ├── documents/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Case documents viewer
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc_case, storage-be/asset
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/kyc-cases/{case_id}/documents
    │   │   │   │   │   │   │   │   │   ├── reopen/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reopen KYC case
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc_case
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/kyc-cases/{case_id}/reopen
    │   │   │   │   │   │   │   │   │   ├── review/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Review KYC case
    │   │   │   │   │   │   │   │   │   │       # - Document verification
    │   │   │   │   │   │   │   │   │   │       # - Approve/reject/escalate
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc_case
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/kyc-cases/{case_id}/review
    │   │   │   │   │   │   │   │   │   └── page.tsx  # KYC case details
    │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc_case
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/kyc-cases/{case_id}
    │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # KYC queue
    │   │   │   │   │   │   │   │   │       # - Prioritization
    │   │   │   │   │   │   │   │   │       # - Assignment
    │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc_case
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/kyc-cases/queue
    │   │   │   │   │   │   │   │   └── page.tsx  # KYC cases dashboard
    │   │   │   │   │   │   │   │       # BE: admin-be/kyc_case
    │   │   │   │   │   │   │   │       # GET /v1/admin/kyc-cases
    │   │   │   │   │   │   │   ├── moderation/
    │   │   │   │   │   │   │   │   ├── actions/
    │   │   │   │   │   │   │   │   │   ├── [actionId]/
    │   │   │   │   │   │   │   │   │   │   ├── appeal/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Review appeal
    │   │   │   │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/moderation/actions/{action_id}/appeal
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Action details
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/moderation/actions/{action_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Moderation actions
    │   │   │   │   │   │   │   │   │       # - Warnings
    │   │   │   │   │   │   │   │   │       # - Suspensions
    │   │   │   │   │   │   │   │   │       # - Bans
    │   │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/moderation/actions
    │   │   │   │   │   │   │   │   ├── patterns/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Abuse patterns detection
    │   │   │   │   │   │   │   │   │       # - Pattern analysis
    │   │   │   │   │   │   │   │   │       # - Risk scoring
    │   │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/moderation/patterns
    │   │   │   │   │   │   │   │   └── reports/
    │   │   │   │   │   │   │   │       ├── [reportId]/
    │   │   │   │   │   │   │   │       │   ├── review/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Review report
    │   │   │   │   │   │   │   │       │   │       # - Content review
    │   │   │   │   │   │   │   │       │   │       # - Take action
    │   │   │   │   │   │   │   │       │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │       │   │       # POST /v1/admin/moderation/reports/{report_id}/review
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Report details
    │   │   │   │   │   │   │   │       │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │       │       # GET /v1/admin/moderation/reports/{report_id}
    │   │   │   │   │   │   │   │       ├── queue/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Moderation queue
    │   │   │   │   │   │   │   │       │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │       │       # GET /v1/admin/moderation/reports/queue
    │   │   │   │   │   │   │   │       └── page.tsx  # Reports dashboard
    │   │   │   │   │   │   │   │           # BE: admin-be/moderation
    │   │   │   │   │   │   │   │           # GET /v1/admin/moderation/reports
    │   │   │   │   │   │   │   ├── search-quality/
    │   │   │   │   │   │   │   │   ├── blacklists/
    │   │   │   │   │   │   │   │   │   ├── [blacklistId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Blacklist entry details
    │   │   │   │   │   │   │   │   │   │       # BE: search-be/admin
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/search/blacklists/{blacklist_id}
    │   │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Add blacklist entry
    │   │   │   │   │   │   │   │   │   │       # BE: search-be/admin
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/search/blacklists
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Blacklist management
    │   │   │   │   │   │   │   │   │       # BE: search-be/admin
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/search/blacklists
    │   │   │   │   │   │   │   │   ├── boosts/
    │   │   │   │   │   │   │   │   │   ├── [boostId]/
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit boost rule
    │   │   │   │   │   │   │   │   │   │   │       # BE: search-be/admin
    │   │   │   │   │   │   │   │   │   │   │       # PUT /v1/admin/search/boosts/{boost_id}
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Boost details
    │   │   │   │   │   │   │   │   │   │       # BE: search-be/admin
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/search/boosts/{boost_id}
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create boost rule
    │   │   │   │   │   │   │   │   │   │       # BE: search-be/admin
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/search/boosts
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Boost rules management
    │   │   │   │   │   │   │   │   │       # BE: search-be/admin
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/search/boosts
    │   │   │   │   │   │   │   │   ├── reindex/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Reindex operations
    │   │   │   │   │   │   │   │   │       # - Full reindex
    │   │   │   │   │   │   │   │   │       # - Selective reindex
    │   │   │   │   │   │   │   │   │       # - Progress monitoring
    │   │   │   │   │   │   │   │   │       # BE: search-be/admin
    │   │   │   │   │   │   │   │   │       # POST /v1/admin/search/reindex
    │   │   │   │   │   │   │   │   └── synonyms/
    │   │   │   │   │   │   │   │       ├── [synonymId]/
    │   │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit synonym
    │   │   │   │   │   │   │   │       │   │       # BE: search-be/admin
    │   │   │   │   │   │   │   │       │   │       # PUT /v1/admin/search/synonyms/{synonym_id}
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Synonym details
    │   │   │   │   │   │   │   │       │       # BE: search-be/admin
    │   │   │   │   │   │   │   │       │       # GET /v1/admin/search/synonyms/{synonym_id}
    │   │   │   │   │   │   │   │       ├── create/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Create synonym
    │   │   │   │   │   │   │   │       │       # BE: search-be/admin
    │   │   │   │   │   │   │   │       │       # POST /v1/admin/search/synonyms
    │   │   │   │   │   │   │   │       └── page.tsx  # Synonyms management
    │   │   │   │   │   │   │   │           # BE: search-be/admin
    │   │   │   │   │   │   │   │           # GET /v1/admin/search/synonyms
    │   │   │   │   │   │   │   ├── system/
    │   │   │   │   │   │   │   │   ├── configuration/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # System configuration
    │   │   │   │   │   │   │   │   │       # - Global settings
    │   │   │   │   │   │   │   │   │       # - Environment variables
    │   │   │   │   │   │   │   │   │       # BE: utility-be/config
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/system/config
    │   │   │   │   │   │   │   │   │       # PUT /v1/admin/system/config
    │   │   │   │   │   │   │   │   ├── experiments/
    │   │   │   │   │   │   │   │   │   ├── [experimentId]/
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit experiment
    │   │   │   │   │   │   │   │   │   │   │       # BE: utility-be/experiments
    │   │   │   │   │   │   │   │   │   │   │       # PUT /v1/admin/experiments/{experiment_id}
    │   │   │   │   │   │   │   │   │   │   ├── results/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Experiment results
    │   │   │   │   │   │   │   │   │   │   │       # BE: utility-be/experiments
    │   │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/experiments/{experiment_id}/results
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Experiment details
    │   │   │   │   │   │   │   │   │   │       # BE: utility-be/experiments
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/experiments/{experiment_id}
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create A/B experiment
    │   │   │   │   │   │   │   │   │   │       # BE: utility-be/experiments
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/experiments
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Experiments dashboard
    │   │   │   │   │   │   │   │   │       # BE: utility-be/experiments
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/experiments
    │   │   │   │   │   │   │   │   ├── feature-flags/
    │   │   │   │   │   │   │   │   │   ├── [flagId]/
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit feature flag
    │   │   │   │   │   │   │   │   │   │   │       # BE: utility-be/feature_flags
    │   │   │   │   │   │   │   │   │   │   │       # PUT /v1/admin/feature-flags/{flag_id}
    │   │   │   │   │   │   │   │   │   │   ├── rollout/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Manage rollout
    │   │   │   │   │   │   │   │   │   │   │       # BE: utility-be/feature_flags
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/feature-flags/{flag_id}/rollout
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Feature flag details
    │   │   │   │   │   │   │   │   │   │       # BE: utility-be/feature_flags
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/feature-flags/{flag_id}
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create feature flag
    │   │   │   │   │   │   │   │   │   │       # BE: utility-be/feature_flags
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/feature-flags
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Feature flags dashboard
    │   │   │   │   │   │   │   │   │       # BE: utility-be/feature_flags
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/feature-flags
    │   │   │   │   │   │   │   │   ├── health/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # System health dashboard
    │   │   │   │   │   │   │   │   │       # - Service status
    │   │   │   │   │   │   │   │   │       # - Resource usage
    │   │   │   │   │   │   │   │   │       # - Error rates
    │   │   │   │   │   │   │   │   │       # BE: utility-be/health
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/system/health
    │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # System logs viewer
    │   │   │   │   │   │   │   │   │       # - Real-time logs
    │   │   │   │   │   │   │   │   │       # - Search and filter
    │   │   │   │   │   │   │   │   │       # BE: utility-be/logs
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/system/logs
    │   │   │   │   │   │   │   │   └── metrics/
    │   │   │   │   │   │   │   │       └── page.tsx  # System metrics
    │   │   │   │   │   │   │   │           # - Performance metrics
    │   │   │   │   │   │   │   │           # - Custom dashboards
    │   │   │   │   │   │   │   │           # BE: utility-be/metrics
    │   │   │   │   │   │   │   │           # GET /v1/admin/system/metrics
    │   │   │   │   │   │   │   └── two-person-rules/
    │   │   │   │   │   │   │       ├── [ruleId]/
    │   │   │   │   │   │   │       │   ├── approve/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Approve rule change
    │   │   │   │   │   │   │       │   │       # BE: admin-be/change_approval
    │   │   │   │   │   │   │       │   │       # POST /v1/admin/two-person-rules/{rule_id}/approve
    │   │   │   │   │   │   │       │   └── page.tsx  # Rule details
    │   │   │   │   │   │   │       │       # BE: admin-be/change_approval
    │   │   │   │   │   │   │       │       # GET /v1/admin/two-person-rules/{rule_id}
    │   │   │   │   │   │   │       ├── pending/
    │   │   │   │   │   │   │       │   └── page.tsx  # Pending approvals
    │   │   │   │   │   │   │       │       # BE: admin-be/change_approval
    │   │   │   │   │   │   │       │       # GET /v1/admin/two-person-rules?status=pending
    │   │   │   │   │   │   │       └── page.tsx  # Two-person rules dashboard
    │   │   │   │   │   │   │           # BE: admin-be/change_approval
    │   │   │   │   │   │   │           # GET /v1/admin/two-person-rules
    │   │   │   │   │   │   ├── audit-reporting/
    │   │   │   │   │   │   │   └── page.tsx  # Audit & reporting
    │   │   │   │   │   │   │       # - System logs, CSV/BI exports
    │   │   │   │   │   │   │       # BE: utility/audit | financial-be/reports
    │   │   │   │   │   │   │       # GET /v1/admin/audit
    │   │   │   │   │   │   │       # GET /v1/admin/reports?type=…
    │   │   │   │   │   │   ├── business-verification/
    │   │   │   │   │   │   │   └── page.tsx  # Business Verification cases
    │   │   │   │   │   │   │       # - Queue, details, evidence review
    │   │   │   │   │   │   │       # - Approve/Reject with notes
    │   │   │   │   │   │   │       # BE: admin-be/business_verification
    │   │   │   │   │   │   │       # GET  /v1/admin/business-verifications
    │   │   │   │   │   │   │       # POST /v1/admin/business-verifications/{id}/approve
    │   │   │   │   │   │   │       # POST /v1/admin/business-verifications/{id}/reject
    │   │   │   │   │   │   ├── communications/
    │   │   │   │   │   │   │   ├── broadcasts/
    │   │   │   │   │   │   │   │   └── page.tsx  # Broadcast messages
    │   │   │   │   │   │   │   │       # - Create/send announcements
    │   │   │   │   │   │   │   │       # - Rate limits & compliance
    │   │   │   │   │   │   │   │       # BE: communications-be/broadcast
    │   │   │   │   │   │   │   │       # POST /v1/broadcasts
    │   │   │   │   │   │   │   │       # GET  /v1/broadcasts
    │   │   │   │   │   │   │   ├── campaigns/
    │   │   │   │   │   │   │   │   └── page.tsx  # Multi-step campaigns
    │   │   │   │   │   │   │   │       # - Audience + schedule + analytics
    │   │   │   │   │   │   │   │       # BE: communications-be/campaign
    │   │   │   │   │   │   │   │       # GET/POST/PUT/DELETE /v1/campaigns
    │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │       └── page.tsx  # Templates management
    │   │   │   │   │   │   │           # - Email/SMS/in-app templates
    │   │   │   │   │   │   │           # BE: communications-be/template
    │   │   │   │   │   │   │           # GET/POST/PUT/DELETE /v1/templates
    │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   ├── aml-kyc/
    │   │   │   │   │   │   │   │   ├── monitoring/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # AML monitoring dashboard
    │   │   │   │   │   │   │   │   │       # - Suspicious activity
    │   │   │   │   │   │   │   │   │       # - Transaction patterns
    │   │   │   │   │   │   │   │   │       # - Risk scores
    │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc_case, financial-be/transaction
    │   │   │   │   │   │   │   │   │       # GET /v1/kyc/monitoring/suspicious-activity
    │   │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   │   ├── [reportId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # SAR (Suspicious Activity Report) detail
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc_case
    │   │   │   │   │   │   │   │   │   │       # GET /v1/kyc/reports/{report_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # AML reports list
    │   │   │   │   │   │   │   │   │       # - Filed reports
    │   │   │   │   │   │   │   │   │       # - Pending reports
    │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc_case
    │   │   │   │   │   │   │   │   │       # GET /v1/kyc/reports
    │   │   │   │   │   │   │   │   └── risk-assessment/
    │   │   │   │   │   │   │   │       └── page.tsx  # Risk assessment tools
    │   │   │   │   │   │   │   │           # - User risk profiles
    │   │   │   │   │   │   │   │           # - Country risk matrix
    │   │   │   │   │   │   │   │           # - Enhanced due diligence
    │   │   │   │   │   │   │   │           # BE: admin-be/kyc_case
    │   │   │   │   │   │   │   │           # GET /v1/kyc/risk-assessment
    │   │   │   │   │   │   │   ├── data-retention/
    │   │   │   │   │   │   │   │   ├── audit/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Retention audit log
    │   │   │   │   │   │   │   │   │       # - Deletion history
    │   │   │   │   │   │   │   │   │       # - Policy compliance
    │   │   │   │   │   │   │   │   │       # BE: utility/audit
    │   │   │   │   │   │   │   │   │       # GET /v1/audit/retention
    │   │   │   │   │   │   │   │   ├── policies/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Retention policies
    │   │   │   │   │   │   │   │   │       # - Policy definitions
    │   │   │   │   │   │   │   │   │       # - Data categories
    │   │   │   │   │   │   │   │   │       # - Retention periods
    │   │   │   │   │   │   │   │   │       # BE: admin-be/data_retention (if exists) or utility/config
    │   │   │   │   │   │   │   │   │       # GET /v1/retention/policies
    │   │   │   │   │   │   │   │   └── schedule/
    │   │   │   │   │   │   │   │       └── page.tsx  # Deletion schedule
    │   │   │   │   │   │   │   │           # - Upcoming deletions
    │   │   │   │   │   │   │   │           # - Retention expirations
    │   │   │   │   │   │   │   │           # BE: admin-be/data_retention
    │   │   │   │   │   │   │   │           # GET /v1/retention/schedule
    │   │   │   │   │   │   │   ├── document-verification/
    │   │   │   │   │   │   │   │   ├── [documentId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Document review interface
    │   │   │   │   │   │   │   │   │       # - Document viewer
    │   │   │   │   │   │   │   │   │       # - Verification checks
    │   │   │   │   │   │   │   │   │       # - Approve/reject/request-more
    │   │   │   │   │   │   │   │   │       # BE: admin-be/business_verification, storage-be/asset
    │   │   │   │   │   │   │   │   │       # GET /v1/verification/documents/{document_id}
    │   │   │   │   │   │   │   │   │       # PUT /v1/verification/documents/{document_id}/review
    │   │   │   │   │   │   │   │   ├── automated-checks/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Automated verification rules
    │   │   │   │   │   │   │   │   │       # - OCR settings
    │   │   │   │   │   │   │   │   │       # - Validation rules
    │   │   │   │   │   │   │   │   │       # - ML model performance
    │   │   │   │   │   │   │   │   │       # BE: admin-be/business_verification
    │   │   │   │   │   │   │   │   │       # GET /v1/verification/automation-rules
    │   │   │   │   │   │   │   │   └── queue/
    │   │   │   │   │   │   │   │       └── page.tsx  # Document verification queue
    │   │   │   │   │   │   │   │           # - Pending documents
    │   │   │   │   │   │   │   │           # - Priority sorting
    │   │   │   │   │   │   │   │           # - Auto-verification status
    │   │   │   │   │   │   │   │           # BE: admin-be/business_verification, storage-be/asset
    │   │   │   │   │   │   │   │           # GET /v1/verification/documents/queue
    │   │   │   │   │   │   │   ├── gdpr/
    │   │   │   │   │   │   │   │   ├── consent-management/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Consent logs & management
    │   │   │   │   │   │   │   │   │       # - User consent history
    │   │   │   │   │   │   │   │   │       # - Consent versions
    │   │   │   │   │   │   │   │   │       # - Audit trail
    │   │   │   │   │   │   │   │   │       # BE: users-be/consent, utility/audit
    │   │   │   │   │   │   │   │   │       # GET /v1/users/consent-logs
    │   │   │   │   │   │   │   │   ├── deletion-requests/
    │   │   │   │   │   │   │   │   │   ├── [requestId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Deletion request detail
    │   │   │   │   │   │   │   │   │   │       # - Data preview
    │   │   │   │   │   │   │   │   │   │       # - Retention check
    │   │   │   │   │   │   │   │   │   │       # - Process deletion
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/privacy, users-be/user
    │   │   │   │   │   │   │   │   │   │       # GET /v1/privacy/deletion-requests/{request_id}
    │   │   │   │   │   │   │   │   │   │       # POST /v1/privacy/deletion-requests/{request_id}/process
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Deletion requests queue
    │   │   │   │   │   │   │   │   │       # BE: admin-be/privacy
    │   │   │   │   │   │   │   │   │       # GET /v1/privacy/deletion-requests
    │   │   │   │   │   │   │   │   ├── export-requests/
    │   │   │   │   │   │   │   │   │   ├── [requestId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Export request detail
    │   │   │   │   │   │   │   │   │   │       # - Request review
    │   │   │   │   │   │   │   │   │   │       # - Generate export
    │   │   │   │   │   │   │   │   │   │       # - Approve/deny
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/privacy, users-be/user
    │   │   │   │   │   │   │   │   │   │       # GET /v1/privacy/export-requests/{request_id}
    │   │   │   │   │   │   │   │   │   │       # POST /v1/privacy/export-requests/{request_id}/process
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Export requests queue
    │   │   │   │   │   │   │   │   │       # BE: admin-be/privacy
    │   │   │   │   │   │   │   │   │       # GET /v1/privacy/export-requests
    │   │   │   │   │   │   │   │   └── reports/
    │   │   │   │   │   │   │   │       └── page.tsx  # GDPR compliance reports
    │   │   │   │   │   │   │   │           # - Processing activities
    │   │   │   │   │   │   │   │           # - Data inventory
    │   │   │   │   │   │   │   │           # - Breach reports
    │   │   │   │   │   │   │   │           # BE: admin-be/privacy, utility/audit
    │   │   │   │   │   │   │   │           # GET /v1/privacy/reports
    │   │   │   │   │   │   │   ├── incidents/
    │   │   │   │   │   │   │   │   ├── [incidentId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit incident details
    │   │   │   │   │   │   │   │   │   │       # - Update status
    │   │   │   │   │   │   │   │   │   │       # - Add postmortem
    │   │   │   │   │   │   │   │   │   │       # - Affected services
    │   │   │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/status/incidents/{incident_id}
    │   │   │   │   │   │   │   │   │   ├── timeline/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Incident timeline
    │   │   │   │   │   │   │   │   │   │       # - Event log
    │   │   │   │   │   │   │   │   │   │       # - Update history
    │   │   │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │   │   │       # GET /v1/status/incidents/{incident_id}/timeline
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Incident detail
    │   │   │   │   │   │   │   │   │       # - Current status
    │   │   │   │   │   │   │   │   │       # - Impact assessment
    │   │   │   │   │   │   │   │   │       # - Resolution steps
    │   │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │   │       # GET /v1/status/incidents/{incident_id}
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create new incident
    │   │   │   │   │   │   │   │   │       # - Incident type selection
    │   │   │   │   │   │   │   │   │       # - Severity level
    │   │   │   │   │   │   │   │   │       # - Affected services
    │   │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │   │       # POST /v1/status/incidents
    │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Historical incidents
    │   │   │   │   │   │   │   │   │       # - Past incidents archive
    │   │   │   │   │   │   │   │   │       # - Postmortems
    │   │   │   │   │   │   │   │   │       # - Lessons learned
    │   │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │   │       # GET /v1/status/incidents/history
    │   │   │   │   │   │   │   │   └── page.tsx  # Active incidents dashboard
    │   │   │   │   │   │   │   │       # - Current incidents
    │   │   │   │   │   │   │   │       # - Quick actions
    │   │   │   │   │   │   │   │       # - Status board
    │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │       # GET /v1/status/incidents?status=active
    │   │   │   │   │   │   │   ├── maintenance/
    │   │   │   │   │   │   │   │   ├── [maintenanceId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit maintenance window
    │   │   │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/status/maintenance/{maintenance_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Maintenance detail
    │   │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │   │       # GET /v1/status/maintenance/{maintenance_id}
    │   │   │   │   │   │   │   │   ├── schedule/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Schedule maintenance
    │   │   │   │   │   │   │   │   │       # - Date/time selection
    │   │   │   │   │   │   │   │   │       # - Affected services
    │   │   │   │   │   │   │   │   │       # - Notification plan
    │   │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │   │       # POST /v1/status/maintenance
    │   │   │   │   │   │   │   │   └── page.tsx  # Maintenance calendar
    │   │   │   │   │   │   │   │       # - Upcoming maintenance
    │   │   │   │   │   │   │   │       # - Impact windows
    │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │       # GET /v1/status/maintenance
    │   │   │   │   │   │   │   └── system-health/
    │   │   │   │   │   │   │       ├── metrics/
    │   │   │   │   │   │   │       │   └── page.tsx  # System metrics dashboard
    │   │   │   │   │   │   │       │       # - CPU/Memory usage
    │   │   │   │   │   │   │       │       # - Database performance
    │   │   │   │   │   │   │       │       # - Queue depths
    │   │   │   │   │   │   │       │       # BE: utility/metrics (or monitoring service)
    │   │   │   │   │   │   │       │       # GET /v1/metrics/system
    │   │   │   │   │   │   │       ├── services/
    │   │   │   │   │   │   │       │   └── page.tsx  # Service health overview
    │   │   │   │   │   │   │       │       # - All microservices status
    │   │   │   │   │   │   │       │       # - Uptime metrics
    │   │   │   │   │   │   │       │       # - Response times
    │   │   │   │   │   │   │       │       # BE: utility/status
    │   │   │   │   │   │   │       │       # GET /v1/status/services
    │   │   │   │   │   │   │       └── page.tsx  # Health dashboard
    │   │   │   │   │   │   │           # - Overall system health
    │   │   │   │   │   │   │           # - Critical alerts
    │   │   │   │   │   │   │           # - Performance trends
    │   │   │   │   │   │   │           # BE: utility/status
    │   │   │   │   │   │   │           # GET /v1/status/health
    │   │   │   │   │   │   ├── financial/
    │   │   │   │   │   │   │   └── tax/
    │   │   │   │   │   │   │       └── page.tsx  # Tax ops (admin)
    │   │   │   │   │   │   │           # - Forms queues, reports
    │   │   │   │   │   │   │           # BE: financial-be/tax
    │   │   │   │   │   │   │           # GET /v1/tax/forms
    │   │   │   │   │   │   │           # GET /v1/tax/reports
    │   │   │   │   │   │   ├── kyc-cases/
    │   │   │   │   │   │   │   └── page.tsx  # KYC case management
    │   │   │   │   │   │   │       # - Triage, document review
    │   │   │   │   │   │   │       # - Approve/Reject/Request-info
    │   │   │   │   │   │   │       # BE: admin-be/kyc_case
    │   │   │   │   │   │   │       # GET  /v1/admin/kyc/cases
    │   │   │   │   │   │   │       # POST /v1/admin/kyc/cases/{id}/approve
    │   │   │   │   │   │   │       # POST /v1/admin/kyc/cases/{id}/reject
    │   │   │   │   │   │   │       # POST /v1/admin/kyc/cases/{id}/request-info
    │   │   │   │   │   │   ├── moderation/
    │   │   │   │   │   │   │   ├── actions/
    │   │   │   │   │   │   │   │   └── page.tsx  # Enforcement actions
    │   │   │   │   │   │   │   │       # - Warn / Suspend / Ban
    │   │   │   │   │   │   │   │       # BE: users-be/moderation
    │   │   │   │   │   │   │   │       # POST /v1/admin/warning
    │   │   │   │   │   │   │   │       # POST /v1/admin/suspension
    │   │   │   │   │   │   │   │       # POST /v1/admin/ban
    │   │   │   │   │   │   │   ├── appeals/
    │   │   │   │   │   │   │   │   └── page.tsx  # Appeals review
    │   │   │   │   │   │   │   │       # - View appeals, decide outcomes
    │   │   │   │   │   │   │   │       # BE: users-be/appeal
    │   │   │   │   │   │   │   │       # GET  /v1/admin/appeals
    │   │   │   │   │   │   │   │       # POST /v1/admin/appeals/{id}/decide
    │   │   │   │   │   │   │   └── review-queue/
    │   │   │   │   │   │   │       └── page.tsx  # Moderation queue
    │   │   │   │   │   │   │           # - Content & report review
    │   │   │   │   │   │   │           # BE: users-be/moderation
    │   │   │   │   │   │   │           # GET /v1/admin/moderation/queue
    │   │   │   │   │   │   ├── ops/
    │   │   │   │   │   │   │   ├── admin-session/
    │   │   │   │   │   │   │   │   └── page.tsx  # JIT “break-glass” admin session
    │   │   │   │   │   │   │   │       # - Request time-boxed access
    │   │   │   │   │   │   │   │       # - Two-person approval
    │   │   │   │   │   │   │   │       # - Full audit trail
    │   │   │   │   │   │   │   │       # BE: admin-be/admin_session
    │   │   │   │   │   │   │   │       # POST /v1/admin/sessions {reason,duration}
    │   │   │   │   │   │   │   │       # GET  /v1/admin/sessions
    │   │   │   │   │   │   │   │       # POST /v1/admin/sessions/{id}/approve
    │   │   │   │   │   │   │   └── change-approval/
    │   │   │   │   │   │   │       └── page.tsx  # Two-person change approvals
    │   │   │   │   │   │   │           # - Risky change queue
    │   │   │   │   │   │   │           # - Approve / Rollback
    │   │   │   │   │   │   │           # BE: admin-be/change_approval
    │   │   │   │   │   │   │           # GET  /v1/admin/change-approvals
    │   │   │   │   │   │   │           # POST /v1/admin/change-approvals/{id}/approve
    │   │   │   │   │   │   │           # POST /v1/admin/change-approvals/{id}/rollback
    │   │   │   │   │   │   ├── platform-config/
    │   │   │   │   │   │   │   ├── integrations/
    │   │   │   │   │   │   │   │   ├── [integrationId]/
    │   │   │   │   │   │   │   │   │   ├── configure/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Configure integration
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/integrations (if exists)
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/integrations/{integration_id}/config
    │   │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Integration logs
    │   │   │   │   │   │   │   │   │   │       # BE: utility/audit
    │   │   │   │   │   │   │   │   │   │       # GET /v1/integrations/{integration_id}/logs
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Integration detail
    │   │   │   │   │   │   │   │   │       # BE: admin-be/integrations
    │   │   │   │   │   │   │   │   │       # GET /v1/integrations/{integration_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Integrations list
    │   │   │   │   │   │   │   │       # - Payment providers
    │   │   │   │   │   │   │   │       # - Email services
    │   │   │   │   │   │   │   │       # - Storage providers
    │   │   │   │   │   │   │   │       # - Auth providers
    │   │   │   │   │   │   │   │       # BE: admin-be/integrations
    │   │   │   │   │   │   │   │       # GET /v1/integrations
    │   │   │   │   │   │   │   ├── localization/
    │   │   │   │   │   │   │   │   ├── languages/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Language management
    │   │   │   │   │   │   │   │   │       # - Enabled languages
    │   │   │   │   │   │   │   │   │       # - Default language
    │   │   │   │   │   │   │   │   │       # - RTL settings
    │   │   │   │   │   │   │   │   │       # BE: utility/i18n
    │   │   │   │   │   │   │   │   │       # GET /v1/config/languages
    │   │   │   │   │   │   │   │   │       # PUT /v1/config/languages
    │   │   │   │   │   │   │   │   └── regions/
    │   │   │   │   │   │   │   │       └── page.tsx  # Regional settings
    │   │   │   │   │   │   │   │           # - Timezone defaults
    │   │   │   │   │   │   │   │           # - Currency settings
    │   │   │   │   │   │   │   │           # - Date/time formats
    │   │   │   │   │   │   │   │           # BE: utility/config
    │   │   │   │   │   │   │   │           # GET /v1/config/regions
    │   │   │   │   │   │   │   │           # PUT /v1/config/regions
    │   │   │   │   │   │   │   ├── notifications/
    │   │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Notification settings
    │   │   │   │   │   │   │   │   │       # - Default preferences
    │   │   │   │   │   │   │   │   │       # - Delivery channels
    │   │   │   │   │   │   │   │   │       # - Retry policies
    │   │   │   │   │   │   │   │   │       # BE: communications-be/config
    │   │   │   │   │   │   │   │   │       # GET /v1/notifications/config
    │   │   │   │   │   │   │   │   │       # PUT /v1/notifications/config
    │   │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │   │       ├── [templateId]/
    │   │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit notification template
    │   │   │   │   │   │   │   │       │   │       # BE: communications-be/template
    │   │   │   │   │   │   │   │       │   │       # PUT /v1/notifications/templates/{template_id}
    │   │   │   │   │   │   │   │       │   ├── preview/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Preview template
    │   │   │   │   │   │   │   │       │   │       # BE: communications-be/template
    │   │   │   │   │   │   │   │       │   │       # POST /v1/notifications/templates/{template_id}/preview
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Template detail
    │   │   │   │   │   │   │   │       │       # BE: communications-be/template
    │   │   │   │   │   │   │   │       │       # GET /v1/notifications/templates/{template_id}
    │   │   │   │   │   │   │   │       └── page.tsx  # Template library
    │   │   │   │   │   │   │   │           # BE: communications-be/template
    │   │   │   │   │   │   │   │           # GET /v1/notifications/templates
    │   │   │   │   │   │   │   └── pricing/
    │   │   │   │   │   │   │       └── page.tsx  # Pricing configuration
    │   │   │   │   │   │   │           # - Commission rates
    │   │   │   │   │   │   │           # - Subscription pricing
    │   │   │   │   │   │   │           # - Regional pricing
    │   │   │   │   │   │   │           # BE: financial-be/pricing_config (if exists)
    │   │   │   │   │   │   │           # GET /v1/config/pricing
    │   │   │   │   │   │   │           # PUT /v1/config/pricing
    │   │   │   │   │   │   │           # Note: Requires change_approval
    │   │   │   │   │   │   ├── refund-cases/
    │   │   │   │   │   │   │   └── page.tsx  # Refund & goodwill credits
    │   │   │   │   │   │   │       # - Intake, investigation, decision
    │   │   │   │   │   │   │       # BE: admin-be/refund_case | admin-be/goodwill_credit
    │   │   │   │   │   │   │       # GET  /v1/admin/refunds
    │   │   │   │   │   │   │       # POST /v1/admin/refunds/{id}/approve
    │   │   │   │   │   │   │       # POST /v1/admin/refunds/{id}/deny
    │   │   │   │   │   │   ├── search/
    │   │   │   │   │   │   │   └── explain/
    │   │   │   │   │   │   │       └── page.tsx  # Search explainability (admin)
    │   │   │   │   │   │   │           # - "Why this result" tooling
    │   │   │   │   │   │   │           # BE: search-be/explainability
    │   │   │   │   │   │   │           # GET /v1/explain/{docId}
    │   │   │   │   │   │   ├── search-quality/
    │   │   │   │   │   │   │   ├── facets-filters/
    │   │   │   │   │   │   │   │   └── page.tsx  # Facets & filters config
    │   │   │   │   │   │   │   │       # - Define schemas & options
    │   │   │   │   │   │   │   │       # BE: search-be/facets | search-be/filters
    │   │   │   │   │   │   │   │       # GET /v1/search/admin/facets
    │   │   │   │   │   │   │   │       # PUT /v1/search/admin/facets
    │   │   │   │   │   │   │   │       # GET /v1/search/admin/filters
    │   │   │   │   │   │   │   │       # PUT /v1/search/admin/filters
    │   │   │   │   │   │   │   ├── indexing/
    │   │   │   │   │   │   │   │   └── page.tsx  # Indexing & backfills
    │   │   │   │   │   │   │   │       # - Trigger reindex/backfills
    │   │   │   │   │   │   │   │       # - Monitor progress/errors
    │   │   │   │   │   │   │   │       # BE: search-be/index
    │   │   │   │   │   │   │   │       # POST /v1/search/admin/reindex
    │   │   │   │   │   │   │   │       # GET  /v1/search/admin/index/jobs
    │   │   │   │   │   │   │   ├── language/
    │   │   │   │   │   │   │   │   └── page.tsx  # Language analyzers
    │   │   │   │   │   │   │   │       # - Stopwords, stemming, per-locale
    │   │   │   │   │   │   │   │       # BE: search-be/language
    │   │   │   │   │   │   │   │       # GET /v1/search/admin/languages
    │   │   │   │   │   │   │   │       # PUT /v1/search/admin/languages
    │   │   │   │   │   │   │   ├── metrics/
    │   │   │   │   │   │   │   │   └── page.tsx  # Search metrics
    │   │   │   │   │   │   │   │       # - Query logs, CTR, latency, drift
    │   │   │   │   │   │   │   │       # BE: search-be/metrics
    │   │   │   │   │   │   │   │       # GET /v1/search/admin/metrics?range=…
    │   │   │   │   │   │   │   ├── rewrites/
    │   │   │   │   │   │   │   │   └── page.tsx  # Query rewrites & synonyms
    │   │   │   │   │   │   │   │       # - Edit dictionaries
    │   │   │   │   │   │   │   │       # - A/B preview before publish
    │   │   │   │   │   │   │   │       # BE: search-be/rewrite
    │   │   │   │   │   │   │   │       # GET/PUT /v1/search/admin/rewrites
    │   │   │   │   │   │   │   │       # POST    /v1/search/admin/preview
    │   │   │   │   │   │   │   └── speller/
    │   │   │   │   │   │   │       └── page.tsx  # Speller configuration
    │   │   │   │   │   │   │           # - Dictionaries & thresholds
    │   │   │   │   │   │   │           # BE: search-be/speller
    │   │   │   │   │   │   │           # GET/PUT /v1/search/admin/speller
    │   │   │   │   │   │   ├── status-incidents/
    │   │   │   │   │   │   │   └── page.tsx  # Incidents & maintenance
    │   │   │   │   │   │   │       # - Open/resolve incidents
    │   │   │   │   │   │   │       # - Post maintenance notes
    │   │   │   │   │   │   │       # BE: utility/status | communications-be/broadcast
    │   │   │   │   │   │   │       # GET  /v1/status
    │   │   │   │   │   │   │       # POST /v1/broadcasts
    │   │   │   │   │   │   └── storage/
    │   │   │   │   │   │       └── page.tsx  # Storage lifecycle (admin)
    │   │   │   │   │   │           # - Retention, soft-delete, restore
    │   │   │   │   │   │           # BE: storage-be/lifecycle
    │   │   │   │   │   │           # GET /v1/storage/lifecycle
    │   │   │   │   │   ├── (auth)/  # Authentication pages (no dashboard layout)
    │   │   │   │   │   │   ├── callback/
    │   │   │   │   │   │   │   └── page.tsx  # OAuth callback handler
    │   │   │   │   │   │   │       # - Keycloak callback
    │   │   │   │   │   │   │       # - Google OAuth callback
    │   │   │   │   │   │   │       # - GitHub OAuth callback
    │   │   │   │   │   │   │       # - LinkedIn OAuth callback
    │   │   │   │   │   │   │       # - Token exchange
    │   │   │   │   │   │   │       # BE: Keycloak token exchange
    │   │   │   │   │   │   │       # POST /oauth2/token (Keycloak)
    │   │   │   │   │   │   │       # Returns: access_token, refresh_token
    │   │   │   │   │   │   ├── forgot-password/
    │   │   │   │   │   │   │   └── page.tsx  # Password reset request
    │   │   │   │   │   │   │       # - Email input
    │   │   │   │   │   │   │       # - Send reset link
    │   │   │   │   │   │   │       # BE: users-be/security/recovery
    │   │   │   │   │   │   │       # POST /v1/auth/forgot-password
    │   │   │   │   │   │   │       # Body: { email }
    │   │   │   │   │   │   │       # Returns: reset_email_sent
    │   │   │   │   │   │   ├── login/
    │   │   │   │   │   │   │   └── page.tsx  # Login page
    │   │   │   │   │   │   │       # - Email/password form
    │   │   │   │   │   │   │       # - Social login buttons (Google, GitHub, LinkedIn)
    │   │   │   │   │   │   │       # - Remember me option
    │   │   │   │   │   │   │       # - Forgot password link
    │   │   │   │   │   │   │       # - Register link
    │   │   │   │   │   │   │       # BE: Keycloak OAuth2 flow
    │   │   │   │   │   │   │       # POST /v1/auth/login (users-be)
    │   │   │   │   │   │   │       # Returns: JWT access_token, refresh_token
    │   │   │   │   │   │   │       # Social: OAuth2 redirect to Keycloak
    │   │   │   │   │   │   ├── mfa/
    │   │   │   │   │   │   │   ├── setup/
    │   │   │   │   │   │   │   │   └── page.tsx  # MFA setup
    │   │   │   │   │   │   │   │       # - QR code display
    │   │   │   │   │   │   │   │       # - Backup codes
    │   │   │   │   │   │   │   │       # - Verify setup
    │   │   │   │   │   │   │   │       # BE: users-be/security/mfa
    │   │   │   │   │   │   │   │       # POST /v1/auth/mfa/setup
    │   │   │   │   │   │   │   │       # GET /v1/auth/mfa/qrcode
    │   │   │   │   │   │   │   │       # POST /v1/auth/mfa/verify-setup
    │   │   │   │   │   │   │   └── verify/
    │   │   │   │   │   │   │       └── page.tsx  # MFA verification
    │   │   │   │   │   │   │           # - OTP input
    │   │   │   │   │   │   │           # - Backup code option
    │   │   │   │   │   │   │           # - Trust device option
    │   │   │   │   │   │   │           # BE: users-be/security/mfa
    │   │   │   │   │   │   │           # POST /v1/auth/mfa/verify
    │   │   │   │   │   │   │           # Body: { code, trust_device }
    │   │   │   │   │   │   ├── register/
    │   │   │   │   │   │   │   ├── verification/
    │   │   │   │   │   │   │   │   └── page.tsx  # Email verification callback
    │   │   │   │   │   │   │   │       # - Verify token from email link
    │   │   │   │   │   │   │   │       # - Success/error messaging
    │   │   │   │   │   │   │   │       # - Auto-redirect to onboarding
    │   │   │   │   │   │   │   │       # BE: users-be/user
    │   │   │   │   │   │   │   │       # POST /v1/users/verify-email
    │   │   │   │   │   │   │   │       # Body: { token }
    │   │   │   │   │   │   │   │       # Returns: verification_status
    │   │   │   │   │   │   │   │       # Publishes: UserVerified event
    │   │   │   │   │   │   │   └── page.tsx  # Registration page
    │   │   │   │   │   │   │       # - User type selection (Freelancer/Client)
    │   │   │   │   │   │   │       # - Email, password, name fields
    │   │   │   │   │   │   │       # - Terms acceptance
    │   │   │   │   │   │   │       # - Social registration
    │   │   │   │   │   │   │       # BE: users-be/user
    │   │   │   │   │   │   │       # POST /v1/users/register
    │   │   │   │   │   │   │       # Body: { email, password, first_name, last_name, user_type, terms_accepted }
    │   │   │   │   │   │   │       # Returns: user_id, verification_email_sent
    │   │   │   │   │   │   │       # Publishes: UserCreated event
    │   │   │   │   │   │   ├── reset-password/
    │   │   │   │   │   │   │   └── page.tsx  # Password reset form
    │   │   │   │   │   │   │       # - New password input
    │   │   │   │   │   │   │       # - Confirm password
    │   │   │   │   │   │   │       # - Token validation
    │   │   │   │   │   │   │       # BE: users-be/security/recovery
    │   │   │   │   │   │   │       # POST /v1/auth/reset-password
    │   │   │   │   │   │   │       # Body: { token, new_password }
    │   │   │   │   │   │   │       # Returns: password_reset_success
    │   │   │   │   │   │   └── layout.tsx  # Auth pages layout
    │   │   │   │   │   │       # - Minimal layout (logo + form)
    │   │   │   │   │   │       # - No header/footer
    │   │   │   │   │   │       # - Language switcher
    │   │   │   │   │   ├── (billing)/
    │   │   │   │   │   │   ├── invoices/
    │   │   │   │   │   │   │   └── page.tsx  # Invoices
    │   │   │   │   │   │   │       # - List/pay/download invoices
    │   │   │   │   │   │   │       # BE: financial-be/invoice|payment
    │   │   │   │   │   │   │       # GET  /v1/invoices
    │   │   │   │   │   │   │       # POST /v1/invoices/{id}/pay
    │   │   │   │   │   │   ├── payouts/
    │   │   │   │   │   │   │   └── page.tsx  # Payouts
    │   │   │   │   │   │   │       # - Destinations & requests
    │   │   │   │   │   │   │       # BE: financial-be/payout
    │   │   │   │   │   │   │       # GET/POST /v1/payouts
    │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   └── page.tsx  # Billing reports (client)
    │   │   │   │   │   │   │       # - Spend by team/project, exports
    │   │   │   │   │   │   │       # BE: financial-be/reports
    │   │   │   │   │   │   │       # GET /v1/reports?scope=client
    │   │   │   │   │   │   ├── tax-forms/
    │   │   │   │   │   │   │   └── page.tsx  # Tax forms (freelancer/client)
    │   │   │   │   │   │   │       # - W-8/W-9/VAT info, upload forms
    │   │   │   │   │   │   │       # BE: financial-be/tax_form (proposed if missing)
    │   │   │   │   │   │   │       # GET/POST /v1/tax/forms (mine)
    │   │   │   │   │   │   └── wallet/
    │   │   │   │   │   │       └── page.tsx  # Wallet
    │   │   │   │   │   │           # - Balance & transactions
    │   │   │   │   │   │           # BE: financial-be/wallet
    │   │   │   │   │   │           # GET /v1/wallet
    │   │   │   │   │   ├── (client)/
    │   │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   │   └── profile/
    │   │   │   │   │   │   │       └── page.tsx  # Billing profile (client)
    │   │   │   │   │   │   │           # - Payment methods, tax/VAT
    │   │   │   │   │   │   │           # BE: users-be/financial_profile|payment_method
    │   │   │   │   │   │   │           # GET/PUT /v1/financial-profile
    │   │   │   │   │   │   │           # CRUD    /v1/payment-methods
    │   │   │   │   │   │   ├── jobs/
    │   │   │   │   │   │   │   ├── manage/
    │   │   │   │   │   │   │   │   └── page.tsx  # Job management
    │   │   │   │   │   │   │   │       # - Eligibility, budget, visibility
    │   │   │   │   │   │   │   │       # BE: jobs-be/job (+ core handlers)
    │   │   │   │   │   │   │   │       # GET/PATCH /v1/jobs/{id}/eligibility|budget|visibility
    │   │   │   │   │   │   │   └── post/
    │   │   │   │   │   │   │       └── page.tsx  # Post a job
    │   │   │   │   │   │   │           # - Draft, attachments, publish
    │   │   │   │   │   │   │           # BE: jobs-be/job | storage-be/asset
    │   │   │   │   │   │   │           # POST /v1/jobs
    │   │   │   │   │   │   │           # PATCH /v1/jobs/{id}
    │   │   │   │   │   │   │           # POST /v1/jobs/{id}/publish
    │   │   │   │   │   │   ├── offers/
    │   │   │   │   │   │   │   └── new/
    │   │   │   │   │   │   │       └── [proposalId]/
    │   │   │   │   │   │   │           └── page.tsx  # Create offer from proposal
    │   │   │   │   │   │   │               # - Terms, milestones, start date
    │   │   │   │   │   │   │               # BE: contracts-be/offer (proposed if missing)
    │   │   │   │   │   │   │               # POST /v1/offers
    │   │   │   │   │   │   ├── org/
    │   │   │   │   │   │   │   └── page.tsx  # Organization & teams
    │   │   │   │   │   │   │       # - Members, roles, invites
    │   │   │   │   │   │   │       # BE: users-be/org|team|invite
    │   │   │   │   │   │   │       # GET/POST/PATCH /v1/orgs, /v1/orgs/{id}/members, /v1/invites
    │   │   │   │   │   │   ├── shortlists/
    │   │   │   │   │   │   │   └── page.tsx  # Talent shortlists
    │   │   │   │   │   │   │       # - Create/manage candidate lists
    │   │   │   │   │   │   │       # BE: proposals-be/shortlist (proposed)
    │   │   │   │   │   │   │       # GET/POST/DELETE /v1/shortlists
    │   │   │   │   │   │   └── talent/
    │   │   │   │   │   │       └── invite/
    │   │   │   │   │   │           └── [userId]/
    │   │   │   │   │   │               └── page.tsx  # Invite freelancer to apply
    │   │   │   │   │   │                   # - Send invite, track status
    │   │   │   │   │   │                   # BE: proposals-be/invite
    │   │   │   │   │   │                   # POST /v1/invites
    │   │   │   │   │   ├── (contracts)/
    │   │   │   │   │   │   ├── contracts/
    │   │   │   │   │   │   │   └── [id]/
    │   │   │   │   │   │   │       └── page.tsx  # Contract detail & SOW
    │   │   │   │   │   │   │           # - Sign/accept, milestones
    │   │   │   │   │   │   │           # BE: contracts-be/contract|sow
    │   │   │   │   │   │   │           # GET /v1/contracts/{id}
    │   │   │   │   │   │   │           # POST /v1/contracts/{id}/sign
    │   │   │   │   │   │   ├── escrow/
    │   │   │   │   │   │   │   └── [contractId]/
    │   │   │   │   │   │   │       ├── fund/
    │   │   │   │   │   │   │       │   └── page.tsx  # Fund escrow (client)
    │   │   │   │   │   │   │       │       # - Amount, payment method
    │   │   │   │   │   │   │       │       # BE: financial-be/escrow
    │   │   │   │   │   │   │       │       # POST /v1/escrow/{contractId}/fund
    │   │   │   │   │   │   │       └── release/
    │   │   │   │   │   │   │           └── page.tsx  # Release escrow (client)
    │   │   │   │   │   │   │               # - Partial/full release
    │   │   │   │   │   │   │               # BE: financial-be/escrow
    │   │   │   │   │   │   │               # POST /v1/escrow/{contractId}/release
    │   │   │   │   │   │   └── workroom/
    │   │   │   │   │   │       └── [id]/
    │   │   │   │   │   │           ├── approvals/
    │   │   │   │   │   │           │   └── page.tsx  # Timesheet approvals (client)
    │   │   │   │   │   │           │       # - Approve/reject, comments
    │   │   │   │   │   │           │       # BE: contracts-be/timesheet_approval (proposed if missing)
    │   │   │   │   │   │           │       # POST /v1/contracts/{id}/timesheets/{tsId}/approve|reject
    │   │   │   │   │   │           ├── files/
    │   │   │   │   │   │           │   └── page.tsx  # Workroom files tab
    │   │   │   │   │   │           │       # - Shared files, versions
    │   │   │   │   │   │           │       # BE: storage-be/asset
    │   │   │   │   │   │           │       # GET /v1/contracts/{id}/files
    │   │   │   │   │   │           └── milestones/
    │   │   │   │   │   │               └── page.tsx  # Milestones (web)
    │   │   │   │   │   │                   # - Create/edit/submit, approvals
    │   │   │   │   │   │                   # BE: contracts-be/milestone
    │   │   │   │   │   │                   # GET/POST/PATCH /v1/contracts/{id}/milestones
    │   │   │   │   │   ├── (dashboard)/  # Main authenticated dashboard
    │   │   │   │   │   │   ├── admin/  # Admin panel (SUPER_ADMIN only)
    │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   ├── platform/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Platform-wide analytics
    │   │   │   │   │   │   │   │   │       # - User growth
    │   │   │   │   │   │   │   │   │       # - Transaction volume
    │   │   │   │   │   │   │   │   │       # - System health
    │   │   │   │   │   │   │   │   │       # BE: admin-be/analytics
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/analytics/platform
    │   │   │   │   │   │   │   │   └── page.tsx  # Analytics overview
    │   │   │   │   │   │   │   │       # BE: admin-be/analytics
    │   │   │   │   │   │   │   │       # GET /v1/admin/analytics
    │   │   │   │   │   │   │   ├── audit-logs/
    │   │   │   │   │   │   │   │   └── page.tsx  # Audit logs
    │   │   │   │   │   │   │   │       # - All admin actions
    │   │   │   │   │   │   │   │       # - Filter by admin
    │   │   │   │   │   │   │   │       # - Filter by action type
    │   │   │   │   │   │   │   │       # - Search logs
    │   │   │   │   │   │   │   │       # BE: admin-be/audit
    │   │   │   │   │   │   │   │       # GET /v1/admin/audit-logs
    │   │   │   │   │   │   │   ├── audits/
    │   │   │   │   │   │   │   │   ├── [auditId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Audit detail
    │   │   │   │   │   │   │   │   │       # - Full audit trail
    │   │   │   │   │   │   │   │   │       # - User actions
    │   │   │   │   │   │   │   │   │       # BE: admin-be/audit
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/audits/{audit_id}
    │   │   │   │   │   │   │   │   ├── exports/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Export audit logs
    │   │   │   │   │   │   │   │   │       # - Date range
    │   │   │   │   │   │   │   │   │       # - Format selection
    │   │   │   │   │   │   │   │   │       # BE: admin-be/audit
    │   │   │   │   │   │   │   │   │       # POST /v1/admin/audits/export
    │   │   │   │   │   │   │   │   └── page.tsx  # Audit logs
    │   │   │   │   │   │   │   │       # - System-wide audit trail
    │   │   │   │   │   │   │   │       # - Filter capabilities
    │   │   │   │   │   │   │   │       # BE: admin-be/audit
    │   │   │   │   │   │   │   │       # GET /v1/admin/audits
    │   │   │   │   │   │   │   ├── business-verification/
    │   │   │   │   │   │   │   │   ├── [caseId]/
    │   │   │   │   │   │   │   │   │   ├── review/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Review business documents
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/business_verification
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/business-verification/{case_id}/review
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Business verification case
    │   │   │   │   │   │   │   │   │       # BE: admin-be/business_verification
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/business-verification/{case_id}
    │   │   │   │   │   │   │   │   ├── [verificationId]/
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve business
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/business-verification
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/business-verifications/{verification_id}/approve
    │   │   │   │   │   │   │   │   │   ├── documents/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Verification documents
    │   │   │   │   │   │   │   │   │   │       # - View uploaded docs
    │   │   │   │   │   │   │   │   │   │       # - Request additional
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/business-verification, storage-be/asset
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/business-verifications/{verification_id}/documents
    │   │   │   │   │   │   │   │   │   ├── reject/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reject business
    │   │   │   │   │   │   │   │   │   │       # - Rejection reason
    │   │   │   │   │   │   │   │   │   │       # - Feedback
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/business-verification
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/business-verifications/{verification_id}/reject
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Business verification detail
    │   │   │   │   │   │   │   │   │       # - Company info
    │   │   │   │   │   │   │   │   │       # - Documents
    │   │   │   │   │   │   │   │   │       # - Decision history
    │   │   │   │   │   │   │   │   │       # BE: admin-be/business-verification
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/business-verifications/{verification_id}
    │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending verifications
    │   │   │   │   │   │   │   │   │       # BE: admin-be/business-verification
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/business-verifications?status=pending
    │   │   │   │   │   │   │   │   └── page.tsx  # Business verification queue
    │   │   │   │   │   │   │   │       # BE: admin-be/business_verification
    │   │   │   │   │   │   │   │       # GET /v1/admin/business-verification/cases
    │   │   │   │   │   │   │   │       # Business verifications queue
    │   │   │   │   │   │   │   │       # BE: admin-be/business-verification
    │   │   │   │   │   │   │   │       # GET /v1/admin/business-verifications
    │   │   │   │   │   │   │   ├── change-approvals/
    │   │   │   │   │   │   │   │   ├── [approvalId]/
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve change
    │   │   │   │   │   │   │   │   │   │       # - Second approver
    │   │   │   │   │   │   │   │   │   │       # - Apply change
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/change-approval
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/change-approvals/{approval_id}/approve
    │   │   │   │   │   │   │   │   │   ├── reject/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reject change
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/change-approval
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/change-approvals/{approval_id}/reject
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Change approval detail
    │   │   │   │   │   │   │   │   │       # - Change details
    │   │   │   │   │   │   │   │   │       # - Risk assessment
    │   │   │   │   │   │   │   │   │       # - Approval history
    │   │   │   │   │   │   │   │   │       # BE: admin-be/change-approval
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/change-approvals/{approval_id}
    │   │   │   │   │   │   │   │   ├── [requestId]/
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve/reject change
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/change_approval
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/change-approvals/{request_id}/approve
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/change-approvals/{request_id}/reject
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Change request detail
    │   │   │   │   │   │   │   │   │       # BE: admin-be/change_approval
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/change-approvals/{request_id}
    │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending approvals
    │   │   │   │   │   │   │   │   │       # BE: admin-be/change-approval
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/change-approvals?status=pending
    │   │   │   │   │   │   │   │   └── page.tsx  # Change approval queue (Two-Person Rule)
    │   │   │   │   │   │   │   │       # - Pending approvals
    │   │   │   │   │   │   │   │       # - My requests
    │   │   │   │   │   │   │   │       # - History
    │   │   │   │   │   │   │   │       # BE: admin-be/change_approval
    │   │   │   │   │   │   │   │       # GET /v1/admin/change-approvals
    │   │   │   │   │   │   │   │       # Change approvals list
    │   │   │   │   │   │   │   │       # - Two-person rule changes
    │   │   │   │   │   │   │   │       # - Risky operations
    │   │   │   │   │   │   │   │       # BE: admin-be/change-approval
    │   │   │   │   │   │   │   ├── communications/
    │   │   │   │   │   │   │   │   ├── broadcasts/
    │   │   │   │   │   │   │   │   │   ├── [broadcastId]/
    │   │   │   │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Broadcast analytics
    │   │   │   │   │   │   │   │   │   │   │       # - Delivery rates
    │   │   │   │   │   │   │   │   │   │   │       # - Engagement metrics
    │   │   │   │   │   │   │   │   │   │   │       # BE: communications-be/broadcast
    │   │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/broadcasts/{broadcast_id}/analytics
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Broadcast detail
    │   │   │   │   │   │   │   │   │   │       # BE: communications-be/broadcast
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/broadcasts/{broadcast_id}
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create broadcast
    │   │   │   │   │   │   │   │   │   │       # - Compose message
    │   │   │   │   │   │   │   │   │   │       # - Target audience
    │   │   │   │   │   │   │   │   │   │       # - Schedule
    │   │   │   │   │   │   │   │   │   │       # BE: communications-be/broadcast
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/broadcasts
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Broadcasts list
    │   │   │   │   │   │   │   │   │       # BE: communications-be/broadcast
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/broadcasts
    │   │   │   │   │   │   │   │   ├── campaigns/
    │   │   │   │   │   │   │   │   │   ├── [campaignId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Campaign detail
    │   │   │   │   │   │   │   │   │   │       # BE: communications-be/campaign
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/campaigns/{campaign_id}
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create campaign
    │   │   │   │   │   │   │   │   │   │       # BE: communications-be/campaign
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/campaigns
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Campaigns list
    │   │   │   │   │   │   │   │   │       # BE: communications-be/campaign
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/campaigns
    │   │   │   │   │   │   │   │   ├── rate-limits/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Communication rate limits
    │   │   │   │   │   │   │   │   │       # - Per-user limits
    │   │   │   │   │   │   │   │   │       # - Global throttling
    │   │   │   │   │   │   │   │   │       # BE: communications-be/preferences
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/communications/rate-limits
    │   │   │   │   │   │   │   │   │       # PUT /v1/admin/communications/rate-limits
    │   │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │   │       ├── [templateId]/
    │   │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit template
    │   │   │   │   │   │   │   │       │   │       # BE: communications-be/template
    │   │   │   │   │   │   │   │       │   │       # PUT /v1/admin/templates/{template_id}
    │   │   │   │   │   │   │   │       │   ├── preview/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Preview template
    │   │   │   │   │   │   │   │       │   │       # BE: communications-be/template
    │   │   │   │   │   │   │   │       │   │       # POST /v1/admin/templates/{template_id}/preview
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Template detail
    │   │   │   │   │   │   │   │       │       # BE: communications-be/template
    │   │   │   │   │   │   │   │       │       # GET /v1/admin/templates/{template_id}
    │   │   │   │   │   │   │   │       ├── create/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Create template
    │   │   │   │   │   │   │   │       │       # BE: communications-be/template
    │   │   │   │   │   │   │   │       │       # POST /v1/admin/templates
    │   │   │   │   │   │   │   │       └── page.tsx  # Templates list
    │   │   │   │   │   │   │   │           # BE: communications-be/template
    │   │   │   │   │   │   │   │           # GET /v1/admin/templates
    │   │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   │   ├── aml-kyc/
    │   │   │   │   │   │   │   │   │   ├── monitoring/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # AML monitoring dashboard
    │   │   │   │   │   │   │   │   │   │       # - Suspicious activity
    │   │   │   │   │   │   │   │   │   │       # - Transaction patterns
    │   │   │   │   │   │   │   │   │   │       # - Risk scores
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc_case, financial-be/transaction
    │   │   │   │   │   │   │   │   │   │       # GET /v1/kyc/monitoring/suspicious-activity
    │   │   │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   │   │   ├── [reportId]/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # SAR (Suspicious Activity Report) detail
    │   │   │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc_case
    │   │   │   │   │   │   │   │   │   │   │       # GET /v1/kyc/reports/{report_id}
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # AML reports list
    │   │   │   │   │   │   │   │   │   │       # - Filed reports
    │   │   │   │   │   │   │   │   │   │       # - Pending reports
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc_case
    │   │   │   │   │   │   │   │   │   │       # GET /v1/kyc/reports
    │   │   │   │   │   │   │   │   │   └── risk-assessment/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Risk assessment tools
    │   │   │   │   │   │   │   │   │           # - User risk profiles
    │   │   │   │   │   │   │   │   │           # - Country risk matrix
    │   │   │   │   │   │   │   │   │           # - Enhanced due diligence
    │   │   │   │   │   │   │   │   │           # BE: admin-be/kyc_case
    │   │   │   │   │   │   │   │   │           # GET /v1/kyc/risk-assessment
    │   │   │   │   │   │   │   │   ├── data-retention/
    │   │   │   │   │   │   │   │   │   ├── audit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Retention audit log
    │   │   │   │   │   │   │   │   │   │       # - Deletion history
    │   │   │   │   │   │   │   │   │   │       # - Policy compliance
    │   │   │   │   │   │   │   │   │   │       # BE: utility/audit
    │   │   │   │   │   │   │   │   │   │       # GET /v1/audit/retention
    │   │   │   │   │   │   │   │   │   ├── policies/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Retention policies
    │   │   │   │   │   │   │   │   │   │       # - Policy definitions
    │   │   │   │   │   │   │   │   │   │       # - Data categories
    │   │   │   │   │   │   │   │   │   │       # - Retention periods
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/data_retention (if exists) or utility/config
    │   │   │   │   │   │   │   │   │   │       # GET /v1/retention/policies
    │   │   │   │   │   │   │   │   │   └── schedule/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Deletion schedule
    │   │   │   │   │   │   │   │   │           # - Upcoming deletions
    │   │   │   │   │   │   │   │   │           # - Retention expirations
    │   │   │   │   │   │   │   │   │           # BE: admin-be/data_retention
    │   │   │   │   │   │   │   │   │           # GET /v1/retention/schedule
    │   │   │   │   │   │   │   │   ├── document-verification/
    │   │   │   │   │   │   │   │   │   ├── [documentId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Document review interface
    │   │   │   │   │   │   │   │   │   │       # - Document viewer
    │   │   │   │   │   │   │   │   │   │       # - Verification checks
    │   │   │   │   │   │   │   │   │   │       # - Approve/reject/request-more
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/business_verification, storage-be/asset
    │   │   │   │   │   │   │   │   │   │       # GET /v1/verification/documents/{document_id}
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/verification/documents/{document_id}/review
    │   │   │   │   │   │   │   │   │   ├── automated-checks/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Automated verification rules
    │   │   │   │   │   │   │   │   │   │       # - OCR settings
    │   │   │   │   │   │   │   │   │   │       # - Validation rules
    │   │   │   │   │   │   │   │   │   │       # - ML model performance
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/business_verification
    │   │   │   │   │   │   │   │   │   │       # GET /v1/verification/automation-rules
    │   │   │   │   │   │   │   │   │   └── queue/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Document verification queue
    │   │   │   │   │   │   │   │   │           # - Pending documents
    │   │   │   │   │   │   │   │   │           # - Priority sorting
    │   │   │   │   │   │   │   │   │           # - Auto-verification status
    │   │   │   │   │   │   │   │   │           # BE: admin-be/business_verification, storage-be/asset
    │   │   │   │   │   │   │   │   │           # GET /v1/verification/documents/queue
    │   │   │   │   │   │   │   │   └── gdpr/
    │   │   │   │   │   │   │   │       ├── consent-management/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Consent logs & management
    │   │   │   │   │   │   │   │       │       # - User consent history
    │   │   │   │   │   │   │   │       │       # - Consent versions
    │   │   │   │   │   │   │   │       │       # - Audit trail
    │   │   │   │   │   │   │   │       │       # BE: users-be/consent, utility/audit
    │   │   │   │   │   │   │   │       │       # GET /v1/users/consent-logs
    │   │   │   │   │   │   │   │       ├── deletion-requests/
    │   │   │   │   │   │   │   │       │   ├── [requestId]/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Deletion request detail
    │   │   │   │   │   │   │   │       │   │       # - Data preview
    │   │   │   │   │   │   │   │       │   │       # - Retention check
    │   │   │   │   │   │   │   │       │   │       # - Process deletion
    │   │   │   │   │   │   │   │       │   │       # BE: admin-be/privacy, users-be/user
    │   │   │   │   │   │   │   │       │   │       # GET /v1/privacy/deletion-requests/{request_id}
    │   │   │   │   │   │   │   │       │   │       # POST /v1/privacy/deletion-requests/{request_id}/process
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Deletion requests queue
    │   │   │   │   │   │   │   │       │       # BE: admin-be/privacy
    │   │   │   │   │   │   │   │       │       # GET /v1/privacy/deletion-requests
    │   │   │   │   │   │   │   │       ├── export-requests/
    │   │   │   │   │   │   │   │       │   ├── [requestId]/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Export request detail
    │   │   │   │   │   │   │   │       │   │       # - Request review
    │   │   │   │   │   │   │   │       │   │       # - Generate export
    │   │   │   │   │   │   │   │       │   │       # - Approve/deny
    │   │   │   │   │   │   │   │       │   │       # BE: admin-be/privacy, users-be/user
    │   │   │   │   │   │   │   │       │   │       # GET /v1/privacy/export-requests/{request_id}
    │   │   │   │   │   │   │   │       │   │       # POST /v1/privacy/export-requests/{request_id}/process
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Export requests queue
    │   │   │   │   │   │   │   │       │       # BE: admin-be/privacy
    │   │   │   │   │   │   │   │       │       # GET /v1/privacy/export-requests
    │   │   │   │   │   │   │   │       └── reports/
    │   │   │   │   │   │   │   │           └── page.tsx  # GDPR compliance reports
    │   │   │   │   │   │   │   │               # - Processing activities
    │   │   │   │   │   │   │   │               # - Data inventory
    │   │   │   │   │   │   │   │               # - Breach reports
    │   │   │   │   │   │   │   │               # BE: admin-be/privacy, utility/audit
    │   │   │   │   │   │   │   │               # GET /v1/privacy/reports
    │   │   │   │   │   │   │   ├── configurations/
    │   │   │   │   │   │   │   │   ├── experiments/
    │   │   │   │   │   │   │   │   │   ├── [experimentId]/
    │   │   │   │   │   │   │   │   │   │   ├── results/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Experiment results
    │   │   │   │   │   │   │   │   │   │   │       # - A/B test metrics
    │   │   │   │   │   │   │   │   │   │   │       # - Statistical significance
    │   │   │   │   │   │   │   │   │   │   │       # BE: admin-be/experiments
    │   │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/experiments/{experiment_id}/results
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Experiment detail
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/experiments
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/experiments/{experiment_id}
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create experiment
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/experiments
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/experiments
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Experiments list
    │   │   │   │   │   │   │   │   │       # BE: admin-be/experiments
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/experiments
    │   │   │   │   │   │   │   │   ├── feature-flags/
    │   │   │   │   │   │   │   │   │   ├── [flagId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Feature flag detail
    │   │   │   │   │   │   │   │   │   │       # - Toggle flag
    │   │   │   │   │   │   │   │   │   │       # - Rollout percentage
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/config (or utility-be/flags)
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/feature-flags/{flag_id}
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/admin/feature-flags/{flag_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Feature flags list
    │   │   │   │   │   │   │   │   │       # BE: admin-be/config
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/feature-flags
    │   │   │   │   │   │   │   │   └── page.tsx  # System configurations
    │   │   │   │   │   │   │   │       # BE: admin-be/config
    │   │   │   │   │   │   │   │       # GET /v1/admin/configurations
    │   │   │   │   │   │   │   ├── content-moderation/
    │   │   │   │   │   │   │   │   ├── [reportId]/
    │   │   │   │   │   │   │   │   │   ├── actions/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Moderation actions
    │   │   │   │   │   │   │   │   │   │       # - Warning
    │   │   │   │   │   │   │   │   │   │       # - Suspension
    │   │   │   │   │   │   │   │   │   │       # - Ban
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/reports/{report_id}/actions
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Report detail
    │   │   │   │   │   │   │   │   │       # - Content review
    │   │   │   │   │   │   │   │   │       # - Reporter info
    │   │   │   │   │   │   │   │   │       # - History
    │   │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/reports/{report_id}
    │   │   │   │   │   │   │   │   ├── appeals/
    │   │   │   │   │   │   │   │   │   ├── [appealId]/
    │   │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve appeal
    │   │   │   │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/appeals/{appeal_id}/approve
    │   │   │   │   │   │   │   │   │   │   ├── reject/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reject appeal
    │   │   │   │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/appeals/{appeal_id}/reject
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Appeal detail
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/appeals/{appeal_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Appeals queue
    │   │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/appeals
    │   │   │   │   │   │   │   │   ├── queue/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Moderation queue
    │   │   │   │   │   │   │   │   │       # - Pending reports
    │   │   │   │   │   │   │   │   │       # - Priority sorting
    │   │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/reports?status=pending
    │   │   │   │   │   │   │   │   └── rules/
    │   │   │   │   │   │   │   │       └── page.tsx  # Moderation rules
    │   │   │   │   │   │   │   │           # - Auto-moderation rules
    │   │   │   │   │   │   │   │           # - Keyword filters
    │   │   │   │   │   │   │   │           # BE: admin-be/moderation
    │   │   │   │   │   │   │   │           # GET /v1/admin/moderation/rules
    │   │   │   │   │   │   │   │           # PUT /v1/admin/moderation/rules
    │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   │   │   ├── assign/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Assign dispute to admin
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/disputes
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/disputes/{dispute_id}/assign
    │   │   │   │   │   │   │   │   │   ├── escalate/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Escalate dispute
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/disputes
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/disputes/{dispute_id}/escalate
    │   │   │   │   │   │   │   │   │   └── resolve/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Resolve dispute
    │   │   │   │   │   │   │   │   │           # - Resolution decision
    │   │   │   │   │   │   │   │   │           # - Financial settlement
    │   │   │   │   │   │   │   │   │           # - Explanation
    │   │   │   │   │   │   │   │   │           # BE: admin-be/disputes
    │   │   │   │   │   │   │   │   │           # POST /v1/admin/disputes/{dispute_id}/resolve
    │   │   │   │   │   │   │   │   │           # Publishes: DisputeResolved event
    │   │   │   │   │   │   │   │   └── page.tsx  # Disputes management
    │   │   │   │   │   │   │   │       # - Open disputes
    │   │   │   │   │   │   │   │       # - Assigned to me
    │   │   │   │   │   │   │   │       # - Resolved disputes
    │   │   │   │   │   │   │   │       # BE: admin-be/disputes
    │   │   │   │   │   │   │   │       # GET /v1/admin/disputes
    │   │   │   │   │   │   │   ├── financial-ops/
    │   │   │   │   │   │   │   │   ├── chargebacks/
    │   │   │   │   │   │   │   │   │   ├── [chargebackId]/
    │   │   │   │   │   │   │   │   │   │   ├── dispute/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute chargeback
    │   │   │   │   │   │   │   │   │   │   │       # BE: financial-be/chargeback, admin-be/financial-ops
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/chargebacks/{chargeback_id}/dispute
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Chargeback detail
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/chargeback
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/chargebacks/{chargeback_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Chargebacks list
    │   │   │   │   │   │   │   │   │       # BE: financial-be/chargeback
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/chargebacks
    │   │   │   │   │   │   │   │   ├── goodwill-credits/
    │   │   │   │   │   │   │   │   │   ├── [creditId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Goodwill credit detail
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/goodwill-credit
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/goodwill-credits/{credit_id}
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve goodwill credit
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/goodwill-credit
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/goodwill-credits/{credit_id}/approve
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create goodwill credit
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/goodwill-credit
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/goodwill-credits
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Goodwill credits list
    │   │   │   │   │   │   │   │   │       # BE: admin-be/goodwill-credit
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/goodwill-credits
    │   │   │   │   │   │   │   │   ├── payouts/
    │   │   │   │   │   │   │   │   │   ├── [payoutId]/
    │   │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve payout
    │   │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payout
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/payouts/{payout_id}/approve
    │   │   │   │   │   │   │   │   │   │   ├── hold/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Hold payout
    │   │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payout
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/payouts/{payout_id}/hold
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Payout detail
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payout
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/payouts/{payout_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Payouts review queue
    │   │   │   │   │   │   │   │   │       # BE: financial-be/payout
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/payouts?status=pending
    │   │   │   │   │   │   │   │   ├── reconciliation/
    │   │   │   │   │   │   │   │   │   ├── [reconId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reconciliation detail
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/reconciliation
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/reconciliation/{recon_id}
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Start reconciliation
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/reconciliation
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/reconciliation
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Reconciliation reports
    │   │   │   │   │   │   │   │   │       # BE: financial-be/reconciliation
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/reconciliation
    │   │   │   │   │   │   │   │   └── refund-cases/
    │   │   │   │   │   │   │   │       ├── [caseId]/
    │   │   │   │   │   │   │   │       │   ├── approve/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Approve refund
    │   │   │   │   │   │   │   │       │   │       # BE: admin-be/refund-case
    │   │   │   │   │   │   │   │       │   │       # POST /v1/admin/refund-cases/{case_id}/approve
    │   │   │   │   │   │   │   │       │   ├── investigation/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Investigation notes
    │   │   │   │   │   │   │   │       │   │       # BE: admin-be/refund-case
    │   │   │   │   │   │   │   │       │   │       # GET /v1/admin/refund-cases/{case_id}/investigation
    │   │   │   │   │   │   │   │       │   │       # POST /v1/admin/refund-cases/{case_id}/investigation
    │   │   │   │   │   │   │   │       │   ├── reject/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Reject refund
    │   │   │   │   │   │   │   │       │   │       # BE: admin-be/refund-case
    │   │   │   │   │   │   │   │       │   │       # POST /v1/admin/refund-cases/{case_id}/reject
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Refund case detail
    │   │   │   │   │   │   │   │       │       # BE: admin-be/refund-case
    │   │   │   │   │   │   │   │       │       # GET /v1/admin/refund-cases/{case_id}
    │   │   │   │   │   │   │   │       └── page.tsx  # Refund cases queue
    │   │   │   │   │   │   │   │           # BE: admin-be/refund-case
    │   │   │   │   │   │   │   │           # GET /v1/admin/refund-cases
    │   │   │   │   │   │   │   ├── incidents/
    │   │   │   │   │   │   │   │   ├── [incidentId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit incident details
    │   │   │   │   │   │   │   │   │   │       # - Update status
    │   │   │   │   │   │   │   │   │   │       # - Add postmortem
    │   │   │   │   │   │   │   │   │   │       # - Affected services
    │   │   │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/status/incidents/{incident_id}
    │   │   │   │   │   │   │   │   │   ├── timeline/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Incident timeline
    │   │   │   │   │   │   │   │   │   │       # - Event log
    │   │   │   │   │   │   │   │   │   │       # - Update history
    │   │   │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │   │   │       # GET /v1/status/incidents/{incident_id}/timeline
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Incident detail
    │   │   │   │   │   │   │   │   │       # - Current status
    │   │   │   │   │   │   │   │   │       # - Impact assessment
    │   │   │   │   │   │   │   │   │       # - Resolution steps
    │   │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │   │       # GET /v1/status/incidents/{incident_id}
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create new incident
    │   │   │   │   │   │   │   │   │       # - Incident type selection
    │   │   │   │   │   │   │   │   │       # - Severity level
    │   │   │   │   │   │   │   │   │       # - Affected services
    │   │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │   │       # POST /v1/status/incidents
    │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Historical incidents
    │   │   │   │   │   │   │   │   │       # - Past incidents archive
    │   │   │   │   │   │   │   │   │       # - Postmortems
    │   │   │   │   │   │   │   │   │       # - Lessons learned
    │   │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │   │       # GET /v1/status/incidents/history
    │   │   │   │   │   │   │   │   └── page.tsx  # Active incidents dashboard
    │   │   │   │   │   │   │   │       # - Current incidents
    │   │   │   │   │   │   │   │       # - Quick actions
    │   │   │   │   │   │   │   │       # - Status board
    │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │       # GET /v1/status/incidents?status=active
    │   │   │   │   │   │   │   ├── kyc/
    │   │   │   │   │   │   │   │   ├── [caseId]/
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve KYC
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc_case
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/kyc/cases/{case_id}/approve
    │   │   │   │   │   │   │   │   │   │       # Publishes: KYCApproved event
    │   │   │   │   │   │   │   │   │   └── reject/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Reject KYC
    │   │   │   │   │   │   │   │   │           # - Rejection reason
    │   │   │   │   │   │   │   │   │           # - Required actions
    │   │   │   │   │   │   │   │   │           # BE: admin-be/kyc_case
    │   │   │   │   │   │   │   │   │           # POST /v1/admin/kyc/cases/{case_id}/reject
    │   │   │   │   │   │   │   │   ├── [kycId]/
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve KYC
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc-case
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/kyc/{kyc_id}/approve
    │   │   │   │   │   │   │   │   │   ├── documents/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # KYC documents review
    │   │   │   │   │   │   │   │   │   │       # - ID verification
    │   │   │   │   │   │   │   │   │   │       # - Address proof
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc-case, storage-be/asset
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/kyc/{kyc_id}/documents
    │   │   │   │   │   │   │   │   │   ├── reject/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reject KYC
    │   │   │   │   │   │   │   │   │   │       # - Rejection reason
    │   │   │   │   │   │   │   │   │   │       # - Request resubmission
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc-case
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/kyc/{kyc_id}/reject
    │   │   │   │   │   │   │   │   │   ├── reopen/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reopen KYC case
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc-case
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/kyc/{kyc_id}/reopen
    │   │   │   │   │   │   │   │   │   └── page.tsx  # KYC case detail
    │   │   │   │   │   │   │   │   │       # - User information
    │   │   │   │   │   │   │   │   │       # - Documents
    │   │   │   │   │   │   │   │   │       # - Decision history
    │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc-case
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/kyc/{kyc_id}
    │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending KYC cases
    │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc-case
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/kyc?status=pending
    │   │   │   │   │   │   │   │   ├── rejected/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Rejected KYC cases
    │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc-case
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/kyc?status=rejected
    │   │   │   │   │   │   │   │   └── page.tsx  # KYC cases queue
    │   │   │   │   │   │   │   │       # - Pending cases
    │   │   │   │   │   │   │   │       # - Approved/rejected cases
    │   │   │   │   │   │   │   │       # BE: admin-be/kyc_case
    │   │   │   │   │   │   │   │       # GET /v1/admin/kyc/cases
    │   │   │   │   │   │   │   │       # KYC queue
    │   │   │   │   │   │   │   │       # - Triage queue
    │   │   │   │   │   │   │   │       # - Priority sorting
    │   │   │   │   │   │   │   │       # BE: admin-be/kyc-case
    │   │   │   │   │   │   │   │       # GET /v1/admin/kyc
    │   │   │   │   │   │   │   ├── maintenance/
    │   │   │   │   │   │   │   │   ├── [maintenanceId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit maintenance window
    │   │   │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/status/maintenance/{maintenance_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Maintenance detail
    │   │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │   │       # GET /v1/status/maintenance/{maintenance_id}
    │   │   │   │   │   │   │   │   ├── schedule/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Schedule maintenance
    │   │   │   │   │   │   │   │   │       # - Date/time selection
    │   │   │   │   │   │   │   │   │       # - Affected services
    │   │   │   │   │   │   │   │   │       # - Notification plan
    │   │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │   │       # POST /v1/status/maintenance
    │   │   │   │   │   │   │   │   └── page.tsx  # Maintenance calendar
    │   │   │   │   │   │   │   │       # - Upcoming maintenance
    │   │   │   │   │   │   │   │       # - Impact windows
    │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │       # GET /v1/status/maintenance
    │   │   │   │   │   │   │   ├── moderation/
    │   │   │   │   │   │   │   │   ├── jobs/
    │   │   │   │   │   │   │   │   │   ├── [jobId]/
    │   │   │   │   │   │   │   │   │   │   └── review/
    │   │   │   │   │   │   │   │   │   │       └── page.tsx  # Review flagged job
    │   │   │   │   │   │   │   │   │   │           # - Job content
    │   │   │   │   │   │   │   │   │   │           # - Flag reason
    │   │   │   │   │   │   │   │   │   │           # - Actions: Approve/Remove/Hide
    │   │   │   │   │   │   │   │   │   │           # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │   │           # POST /v1/admin/moderation/jobs/{job_id}/review
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Flagged jobs
    │   │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/moderation/jobs
    │   │   │   │   │   │   │   │   ├── messages/
    │   │   │   │   │   │   │   │   │   └── [messageId]/
    │   │   │   │   │   │   │   │   │       └── review/
    │   │   │   │   │   │   │   │   │           └── page.tsx  # Review flagged message
    │   │   │   │   │   │   │   │   │               # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │               # POST /v1/admin/moderation/messages/{message_id}/review
    │   │   │   │   │   │   │   │   ├── proposals/
    │   │   │   │   │   │   │   │   │   └── [proposalId]/
    │   │   │   │   │   │   │   │   │       └── review/
    │   │   │   │   │   │   │   │   │           └── page.tsx  # Review flagged proposal
    │   │   │   │   │   │   │   │   │               # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │               # POST /v1/admin/moderation/proposals/{proposal_id}/review
    │   │   │   │   │   │   │   │   └── reviews/
    │   │   │   │   │   │   │   │       └── [reviewId]/
    │   │   │   │   │   │   │   │           └── review/
    │   │   │   │   │   │   │   │               └── page.tsx  # Review flagged review
    │   │   │   │   │   │   │   │                   # BE: admin-be/moderation
    │   │   │   │   │   │   │   │                   # POST /v1/admin/moderation/reviews/{review_id}/review
    │   │   │   │   │   │   │   ├── platform-config/
    │   │   │   │   │   │   │   │   ├── integrations/
    │   │   │   │   │   │   │   │   │   ├── [integrationId]/
    │   │   │   │   │   │   │   │   │   │   ├── configure/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Configure integration
    │   │   │   │   │   │   │   │   │   │   │       # BE: admin-be/integrations (if exists)
    │   │   │   │   │   │   │   │   │   │   │       # PUT /v1/integrations/{integration_id}/config
    │   │   │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Integration logs
    │   │   │   │   │   │   │   │   │   │   │       # BE: utility/audit
    │   │   │   │   │   │   │   │   │   │   │       # GET /v1/integrations/{integration_id}/logs
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Integration detail
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/integrations
    │   │   │   │   │   │   │   │   │   │       # GET /v1/integrations/{integration_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Integrations list
    │   │   │   │   │   │   │   │   │       # - Payment providers
    │   │   │   │   │   │   │   │   │       # - Email services
    │   │   │   │   │   │   │   │   │       # - Storage providers
    │   │   │   │   │   │   │   │   │       # - Auth providers
    │   │   │   │   │   │   │   │   │       # BE: admin-be/integrations
    │   │   │   │   │   │   │   │   │       # GET /v1/integrations
    │   │   │   │   │   │   │   │   ├── limits/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Platform limits configuration
    │   │   │   │   │   │   │   │   │       # - Rate limits
    │   │   │   │   │   │   │   │   │       # - Upload limits
    │   │   │   │   │   │   │   │   │       # - API quotas
    │   │   │   │   │   │   │   │   │       # - Subscription limits
    │   │   │   │   │   │   │   │   │       # BE: utility/config, admin-be/platform_config (if exists)
    │   │   │   │   │   │   │   │   │       # GET /v1/config/limits
    │   │   │   │   │   │   │   │   │       # PUT /v1/config/limits
    │   │   │   │   │   │   │   │   ├── localization/
    │   │   │   │   │   │   │   │   │   ├── languages/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Language management
    │   │   │   │   │   │   │   │   │   │       # - Enabled languages
    │   │   │   │   │   │   │   │   │   │       # - Default language
    │   │   │   │   │   │   │   │   │   │       # - RTL settings
    │   │   │   │   │   │   │   │   │   │       # BE: utility/i18n
    │   │   │   │   │   │   │   │   │   │       # GET /v1/config/languages
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/config/languages
    │   │   │   │   │   │   │   │   │   └── regions/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Regional settings
    │   │   │   │   │   │   │   │   │           # - Timezone defaults
    │   │   │   │   │   │   │   │   │           # - Currency settings
    │   │   │   │   │   │   │   │   │           # - Date/time formats
    │   │   │   │   │   │   │   │   │           # BE: utility/config
    │   │   │   │   │   │   │   │   │           # GET /v1/config/regions
    │   │   │   │   │   │   │   │   │           # PUT /v1/config/regions
    │   │   │   │   │   │   │   │   ├── notifications/
    │   │   │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Notification settings
    │   │   │   │   │   │   │   │   │   │       # - Default preferences
    │   │   │   │   │   │   │   │   │   │       # - Delivery channels
    │   │   │   │   │   │   │   │   │   │       # - Retry policies
    │   │   │   │   │   │   │   │   │   │       # BE: communications-be/config
    │   │   │   │   │   │   │   │   │   │       # GET /v1/notifications/config
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/notifications/config
    │   │   │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │   │   │       ├── [templateId]/
    │   │   │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit notification template
    │   │   │   │   │   │   │   │   │       │   │       # BE: communications-be/template
    │   │   │   │   │   │   │   │   │       │   │       # PUT /v1/notifications/templates/{template_id}
    │   │   │   │   │   │   │   │   │       │   ├── preview/
    │   │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Preview template
    │   │   │   │   │   │   │   │   │       │   │       # BE: communications-be/template
    │   │   │   │   │   │   │   │   │       │   │       # POST /v1/notifications/templates/{template_id}/preview
    │   │   │   │   │   │   │   │   │       │   └── page.tsx  # Template detail
    │   │   │   │   │   │   │   │   │       │       # BE: communications-be/template
    │   │   │   │   │   │   │   │   │       │       # GET /v1/notifications/templates/{template_id}
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Template library
    │   │   │   │   │   │   │   │   │           # BE: communications-be/template
    │   │   │   │   │   │   │   │   │           # GET /v1/notifications/templates
    │   │   │   │   │   │   │   │   └── pricing/
    │   │   │   │   │   │   │   │       └── page.tsx  # Pricing configuration
    │   │   │   │   │   │   │   │           # - Commission rates
    │   │   │   │   │   │   │   │           # - Subscription pricing
    │   │   │   │   │   │   │   │           # - Regional pricing
    │   │   │   │   │   │   │   │           # BE: financial-be/pricing_config (if exists)
    │   │   │   │   │   │   │   │           # GET /v1/config/pricing
    │   │   │   │   │   │   │   │           # PUT /v1/config/pricing
    │   │   │   │   │   │   │   │           # Note: Requires change_approval
    │   │   │   │   │   │   │   ├── refunds/
    │   │   │   │   │   │   │   │   ├── [caseId]/
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve refund
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/refund_case
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/refunds/{case_id}/approve
    │   │   │   │   │   │   │   │   │   └── reject/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Reject refund
    │   │   │   │   │   │   │   │   │           # BE: admin-be/refund_case
    │   │   │   │   │   │   │   │   │           # POST /v1/admin/refunds/{case_id}/reject
    │   │   │   │   │   │   │   │   └── page.tsx  # Refund cases
    │   │   │   │   │   │   │   │       # - Pending refund requests
    │   │   │   │   │   │   │   │       # - Processed refunds
    │   │   │   │   │   │   │   │       # BE: admin-be/refund_case
    │   │   │   │   │   │   │   │       # GET /v1/admin/refunds
    │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   ├── financial/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Financial reports
    │   │   │   │   │   │   │   │   │       # BE: admin-be/reports
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/reports/financial
    │   │   │   │   │   │   │   │   ├── moderation/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Moderation reports
    │   │   │   │   │   │   │   │   │       # BE: admin-be/reports
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/reports/moderation
    │   │   │   │   │   │   │   │   ├── users/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # User reports
    │   │   │   │   │   │   │   │   │       # BE: admin-be/reports
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/reports/users
    │   │   │   │   │   │   │   │   └── page.tsx  # Admin reports
    │   │   │   │   │   │   │   │       # - Platform metrics
    │   │   │   │   │   │   │   │       # - User growth
    │   │   │   │   │   │   │   │       # - Revenue reports
    │   │   │   │   │   │   │   │       # - Moderation stats
    │   │   │   │   │   │   │   │       # BE: admin-be/reports
    │   │   │   │   │   │   │   │       # GET /v1/admin/reports
    │   │   │   │   │   │   │   ├── search-quality/
    │   │   │   │   │   │   │   │   ├── blacklists/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Search blacklists
    │   │   │   │   │   │   │   │   │       # - Blacklisted terms
    │   │   │   │   │   │   │   │   │       # - Blocked users
    │   │   │   │   │   │   │   │   │       # BE: search-be/admin
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/search/blacklists
    │   │   │   │   │   │   │   │   │       # PUT /v1/admin/search/blacklists
    │   │   │   │   │   │   │   │   ├── boosts/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Search boosts
    │   │   │   │   │   │   │   │   │       # - Boosted content
    │   │   │   │   │   │   │   │   │       # - Boost rules
    │   │   │   │   │   │   │   │   │       # BE: search-be/admin
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/search/boosts
    │   │   │   │   │   │   │   │   │       # PUT /v1/admin/search/boosts
    │   │   │   │   │   │   │   │   ├── reindex/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Reindex controls
    │   │   │   │   │   │   │   │   │       # - Trigger reindex
    │   │   │   │   │   │   │   │   │       # - Monitor progress
    │   │   │   │   │   │   │   │   │       # BE: search-be/admin
    │   │   │   │   │   │   │   │   │       # POST /v1/admin/search/reindex
    │   │   │   │   │   │   │   │   └── synonyms/
    │   │   │   │   │   │   │   │       └── page.tsx  # Search synonyms
    │   │   │   │   │   │   │   │           # - Synonym management
    │   │   │   │   │   │   │   │           # - Query expansion rules
    │   │   │   │   │   │   │   │           # BE: search-be/admin
    │   │   │   │   │   │   │   │           # GET /v1/admin/search/synonyms
    │   │   │   │   │   │   │   │           # PUT /v1/admin/search/synonyms
    │   │   │   │   │   │   │   ├── sessions/
    │   │   │   │   │   │   │   │   ├── [sessionId]/
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve JIT session
    │   │   │   │   │   │   │   │   │   │       # - Second approver required
    │   │   │   │   │   │   │   │   │   │       # - Time-box access
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/admin-session
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/sessions/{session_id}/approve
    │   │   │   │   │   │   │   │   │   ├── revoke/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Revoke session
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/admin-session
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/sessions/{session_id}/revoke
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Session detail
    │   │   │   │   │   │   │   │   │       # - Session info
    │   │   │   │   │   │   │   │   │       # - Actions performed
    │   │   │   │   │   │   │   │   │       # BE: admin-be/admin-session
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/sessions/{session_id}
    │   │   │   │   │   │   │   │   ├── request/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Request JIT session
    │   │   │   │   │   │   │   │   │       # - Justification
    │   │   │   │   │   │   │   │   │       # - Requested duration
    │   │   │   │   │   │   │   │   │       # BE: admin-be/admin-session
    │   │   │   │   │   │   │   │   │       # POST /v1/admin/sessions/request
    │   │   │   │   │   │   │   │   └── page.tsx  # Admin sessions list
    │   │   │   │   │   │   │   │       # - Active sessions
    │   │   │   │   │   │   │   │       # - Session history
    │   │   │   │   │   │   │   │       # BE: admin-be/admin-session
    │   │   │   │   │   │   │   │       # GET /v1/admin/sessions
    │   │   │   │   │   │   │   ├── system/
    │   │   │   │   │   │   │   │   ├── announcements/
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create announcement
    │   │   │   │   │   │   │   │   │   │       # BE: communications-be/announcements
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/announcements
    │   │   │   │   │   │   │   │   │   └── page.tsx  # System announcements
    │   │   │   │   │   │   │   │   │       # BE: communications-be/announcements
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/announcements
    │   │   │   │   │   │   │   │   ├── feature-flags/
    │   │   │   │   │   │   │   │   │   ├── [flagId]/
    │   │   │   │   │   │   │   │   │   │   └── edit/
    │   │   │   │   │   │   │   │   │   │       └── page.tsx  # Edit feature flag
    │   │   │   │   │   │   │   │   │   │           # BE: subscriptions-be/feature_toggles
    │   │   │   │   │   │   │   │   │   │           # PUT /v1/admin/feature-flags/{flag_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Feature flags management
    │   │   │   │   │   │   │   │   │       # - List flags
    │   │   │   │   │   │   │   │   │       # - Toggle flags
    │   │   │   │   │   │   │   │   │       # - Rollout percentage
    │   │   │   │   │   │   │   │   │       # BE: subscriptions-be/feature_toggles
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/feature-flags
    │   │   │   │   │   │   │   │   ├── maintenance/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Maintenance mode
    │   │   │   │   │   │   │   │   │       # - Enable/disable maintenance
    │   │   │   │   │   │   │   │   │       # - Maintenance message
    │   │   │   │   │   │   │   │   │       # BE: admin-be/system
    │   │   │   │   │   │   │   │   │       # POST /v1/admin/system/maintenance
    │   │   │   │   │   │   │   │   └── page.tsx  # System settings
    │   │   │   │   │   │   │   │       # - Feature flags
    │   │   │   │   │   │   │   │       # - System configuration
    │   │   │   │   │   │   │   │       # BE: admin-be/system
    │   │   │   │   │   │   │   │       # GET /v1/admin/system/config
    │   │   │   │   │   │   │   ├── system-health/
    │   │   │   │   │   │   │   │   ├── incidents/
    │   │   │   │   │   │   │   │   │   ├── [incidentId]/
    │   │   │   │   │   │   │   │   │   │   ├── postmortem/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Incident postmortem
    │   │   │   │   │   │   │   │   │   │   │       # BE: admin-be/incidents (or utility-be/status)
    │   │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/incidents/{incident_id}/postmortem
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/incidents/{incident_id}/postmortem
    │   │   │   │   │   │   │   │   │   │   ├── resolve/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Resolve incident
    │   │   │   │   │   │   │   │   │   │   │       # BE: admin-be/incidents
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/incidents/{incident_id}/resolve
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Incident detail
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/incidents
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/incidents/{incident_id}
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create incident
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/incidents
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/incidents
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Incidents list
    │   │   │   │   │   │   │   │   │       # BE: admin-be/incidents
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/incidents
    │   │   │   │   │   │   │   │   ├── maintenance/
    │   │   │   │   │   │   │   │   │   ├── [maintenanceId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Maintenance detail
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/maintenance
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/maintenance/{maintenance_id}
    │   │   │   │   │   │   │   │   │   ├── schedule/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Schedule maintenance
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/maintenance
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/maintenance
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Maintenance windows
    │   │   │   │   │   │   │   │   │       # BE: admin-be/maintenance
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/maintenance
    │   │   │   │   │   │   │   │   ├── metrics/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # System metrics dashboard
    │   │   │   │   │   │   │   │   │       # - CPU/Memory usage
    │   │   │   │   │   │   │   │   │       # - Database performance
    │   │   │   │   │   │   │   │   │       # - Queue depths
    │   │   │   │   │   │   │   │   │       # BE: utility/metrics (or monitoring service)
    │   │   │   │   │   │   │   │   │       # GET /v1/metrics/system
    │   │   │   │   │   │   │   │   ├── services/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Service health overview
    │   │   │   │   │   │   │   │   │       # - All microservices status
    │   │   │   │   │   │   │   │   │       # - Uptime metrics
    │   │   │   │   │   │   │   │   │       # - Response times
    │   │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │   │       # GET /v1/status/services
    │   │   │   │   │   │   │   │   ├── status/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # System status dashboard
    │   │   │   │   │   │   │   │   │       # - Service health
    │   │   │   │   │   │   │   │   │       # - Uptime metrics
    │   │   │   │   │   │   │   │   │       # BE: admin-be/system (or utility-be/status)
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/system/status
    │   │   │   │   │   │   │   │   └── page.tsx  # Health dashboard
    │   │   │   │   │   │   │   │       # - Overall system health
    │   │   │   │   │   │   │   │       # - Critical alerts
    │   │   │   │   │   │   │   │       # - Performance trends
    │   │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │   │       # GET /v1/status/health
    │   │   │   │   │   │   │   ├── users/
    │   │   │   │   │   │   │   │   ├── [userId]/
    │   │   │   │   │   │   │   │   │   ├── ban/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Ban user
    │   │   │   │   │   │   │   │   │   │       # - Ban reason
    │   │   │   │   │   │   │   │   │   │       # - Permanent/temporary
    │   │   │   │   │   │   │   │   │   │       # - Related accounts
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/users
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/users/{user_id}/ban
    │   │   │   │   │   │   │   │   │   │       # Publishes: UserBanned event
    │   │   │   │   │   │   │   │   │   │       # - Duration
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/account, admin-be/moderation
    │   │   │   │   │   │   │   │   │   ├── contracts/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # User contracts
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/contract
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/users/{user_id}/contracts
    │   │   │   │   │   │   │   │   │   ├── financials/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # User financial history
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/reports
    │   │   │   │   │   │   │   │   │   │       # GET /v1/admin/users/{user_id}/financials
    │   │   │   │   │   │   │   │   │   ├── impersonate/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Impersonate user
    │   │   │   │   │   │   │   │   │   │       # - Requires approval
    │   │   │   │   │   │   │   │   │   │       # - Audit trail
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/admin-session
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/users/{user_id}/impersonate
    │   │   │   │   │   │   │   │   │   ├── suspend/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Suspend user
    │   │   │   │   │   │   │   │   │   │       # - Suspension reason
    │   │   │   │   │   │   │   │   │   │       # - Duration
    │   │   │   │   │   │   │   │   │   │       # - Notify user
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/users
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/users/{user_id}/suspend
    │   │   │   │   │   │   │   │   │   │       # Publishes: UserSuspended event
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/account, admin-be/moderation
    │   │   │   │   │   │   │   │   │   ├── unban/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Unban user
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/account
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/users/{user_id}/unban
    │   │   │   │   │   │   │   │   │   ├── verify/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Verify user (manual)
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/users
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/users/{user_id}/verify
    │   │   │   │   │   │   │   │   │   │       # Manually verify user
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/account
    │   │   │   │   │   │   │   │   │   ├── warn/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Warn user
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/users
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/users/{user_id}/warn
    │   │   │   │   │   │   │   │   │   ├── warning/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Issue warning
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │   │       # POST /v1/admin/users/{user_id}/warning
    │   │   │   │   │   │   │   │   │   └── page.tsx  # User detail
    │   │   │   │   │   │   │   │   │       # - Full profile
    │   │   │   │   │   │   │   │   │       # - Activity history
    │   │   │   │   │   │   │   │   │       # - Actions menu
    │   │   │   │   │   │   │   │   │       # BE: users-be/profile
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/users/{user_id}
    │   │   │   │   │   │   │   │   ├── banned/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Banned users
    │   │   │   │   │   │   │   │   │       # BE: users-be/account
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/users?status=banned
    │   │   │   │   │   │   │   │   ├── bulk-actions/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Bulk user actions
    │   │   │   │   │   │   │   │   │       # BE: admin-be/users
    │   │   │   │   │   │   │   │   │       # POST /v1/admin/users/bulk-action
    │   │   │   │   │   │   │   │   ├── search/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Search users
    │   │   │   │   │   │   │   │   │       # - Advanced search
    │   │   │   │   │   │   │   │   │       # - Bulk actions
    │   │   │   │   │   │   │   │   │       # BE: search-be/query
    │   │   │   │   │   │   │   │   │       # POST /v1/admin/users/search
    │   │   │   │   │   │   │   │   ├── suspended/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Suspended users
    │   │   │   │   │   │   │   │   │       # BE: users-be/account
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/users?status=suspended
    │   │   │   │   │   │   │   │   └── page.tsx  # Users management
    │   │   │   │   │   │   │   │       # - User list
    │   │   │   │   │   │   │   │       # - Search users
    │   │   │   │   │   │   │   │       # - Filter by status/type
    │   │   │   │   │   │   │   │       # - Bulk actions
    │   │   │   │   │   │   │   │       # BE: admin-be/users
    │   │   │   │   │   │   │   │       # GET /v1/admin/users?filters={...}
    │   │   │   │   │   │   │   ├── layout.tsx  # Admin layout (RBAC guard)
    │   │   │   │   │   │   │   │   # - Only ADMIN, SUPER_ADMIN, MODERATOR
    │   │   │   │   │   │   │   │   # - Admin navigation sidebar
    │   │   │   │   │   │   │   └── page.tsx  # Admin dashboard
    │   │   │   │   │   │   │       # - Key metrics
    │   │   │   │   │   │   │       # - Pending moderation queue
    │   │   │   │   │   │   │       # - Recent admin actions
    │   │   │   │   │   │   │       # - System alerts
    │   │   │   │   │   │   │       # BE: admin-be/dashboard
    │   │   │   │   │   │   │       # GET /v1/admin/dashboard
    │   │   │   │   │   │   │       # - Quick actions
    │   │   │   │   │   │   │       # - Recent activities
    │   │   │   │   │   │   │       # BE: admin-be/analytics
    │   │   │   │   │   │   ├── agency/
    │   │   │   │   │   │   │   ├── overview/
    │   │   │   │   │   │   │   │   └── page.tsx  # Agency dashboard
    │   │   │   │   │   │   │   │       # - Revenue overview
    │   │   │   │   │   │   │   │       # - Active projects
    │   │   │   │   │   │   │   │       # - Team utilization
    │   │   │   │   │   │   │   │       # - Client list
    │   │   │   │   │   │   │   │       # BE: users-be/org, financial-be/analytics
    │   │   │   │   │   │   │   │       # GET /v1/agencies/me/overview
    │   │   │   │   │   │   │   ├── reporting/
    │   │   │   │   │   │   │   │   ├── clients/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Client reports
    │   │   │   │   │   │   │   │   │       # - Per-client spending
    │   │   │   │   │   │   │   │   │       # - Project status
    │   │   │   │   │   │   │   │   │       # BE: financial-be/analytics
    │   │   │   │   │   │   │   │   │       # GET /v1/agencies/reports/clients
    │   │   │   │   │   │   │   │   ├── financial/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Financial reports
    │   │   │   │   │   │   │   │   │       # - Revenue by period
    │   │   │   │   │   │   │   │   │       # - Margins
    │   │   │   │   │   │   │   │   │       # - Forecasts
    │   │   │   │   │   │   │   │   │       # BE: financial-be/analytics
    │   │   │   │   │   │   │   │   │       # GET /v1/agencies/reports/financial
    │   │   │   │   │   │   │   │   └── team/
    │   │   │   │   │   │   │   │       └── page.tsx  # Team performance
    │   │   │   │   │   │   │   │           # - Utilization rates
    │   │   │   │   │   │   │   │           # - Project assignments
    │   │   │   │   │   │   │   │           # BE: users-be/org, jobs-be/analytics
    │   │   │   │   │   │   │   │           # GET /v1/agencies/reports/team
    │   │   │   │   │   │   │   ├── sub-accounts/
    │   │   │   │   │   │   │   │   ├── [subAccountId]/
    │   │   │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Sub-account settings
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/org
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/agencies/sub-accounts/{sub_account_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Sub-account detail
    │   │   │   │   │   │   │   │   │       # - Jobs posted
    │   │   │   │   │   │   │   │   │       # - Contracts
    │   │   │   │   │   │   │   │   │       # - Spending
    │   │   │   │   │   │   │   │   │       # BE: users-be/org
    │   │   │   │   │   │   │   │   │       # GET /v1/agencies/sub-accounts/{sub_account_id}
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create sub-account
    │   │   │   │   │   │   │   │   │       # BE: users-be/org
    │   │   │   │   │   │   │   │   │       # POST /v1/agencies/sub-accounts
    │   │   │   │   │   │   │   │   └── page.tsx  # All sub-accounts
    │   │   │   │   │   │   │   │       # - List all
    │   │   │   │   │   │   │   │       # - Manage
    │   │   │   │   │   │   │   │       # BE: users-be/org
    │   │   │   │   │   │   │   │       # GET /v1/agencies/sub-accounts
    │   │   │   │   │   │   │   ├── talent-pool/
    │   │   │   │   │   │   │   │   ├── [poolId]/
    │   │   │   │   │   │   │   │   │   ├── members/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Pool members
    │   │   │   │   │   │   │   │   │   │       # - Freelancer list
    │   │   │   │   │   │   │   │   │   │       # - Add/remove
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/org, search-be/talent-pool
    │   │   │   │   │   │   │   │   │   │       # GET /v1/talent-pools/{pool_id}/members
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Talent pool detail
    │   │   │   │   │   │   │   │   │       # BE: search-be/talent-pool
    │   │   │   │   │   │   │   │   │       # GET /v1/talent-pools/{pool_id}
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create talent pool
    │   │   │   │   │   │   │   │   │       # BE: search-be/talent-pool
    │   │   │   │   │   │   │   │   │       # POST /v1/talent-pools
    │   │   │   │   │   │   │   │   └── page.tsx  # All talent pools
    │   │   │   │   │   │   │   │       # BE: search-be/talent-pool
    │   │   │   │   │   │   │   │       # GET /v1/talent-pools
    │   │   │   │   │   │   │   └── white-label/
    │   │   │   │   │   │   │       └── page.tsx  # White-label settings
    │   │   │   │   │   │   │           # - Branding
    │   │   │   │   │   │   │           # - Custom domain
    │   │   │   │   │   │   │           # - Logo/colors
    │   │   │   │   │   │   │           # BE: users-be/org
    │   │   │   │   │   │   │           # PUT /v1/agencies/white-label
    │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   ├── clients/
    │   │   │   │   │   │   │   │   └── page.tsx  # Client analytics (freelancer)
    │   │   │   │   │   │   │   │       # - Top clients
    │   │   │   │   │   │   │   │       # - Client retention rate
    │   │   │   │   │   │   │   │       # - Repeat business
    │   │   │   │   │   │   │   │       # - Client lifetime value
    │   │   │   │   │   │   │   │       # BE: users-be/analytics
    │   │   │   │   │   │   │   │       # GET /v1/users/me/analytics/clients
    │   │   │   │   │   │   │   ├── custom-reports/
    │   │   │   │   │   │   │   │   ├── [reportId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit custom report
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/reports
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/reports/custom/{report_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # View custom report
    │   │   │   │   │   │   │   │   │       # BE: financial-be/reports
    │   │   │   │   │   │   │   │   │       # GET /v1/reports/custom/{report_id}
    │   │   │   │   │   │   │   │   ├── new/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create custom report
    │   │   │   │   │   │   │   │   │       # BE: financial-be/reports
    │   │   │   │   │   │   │   │   │       # POST /v1/reports/custom
    │   │   │   │   │   │   │   │   └── page.tsx  # Custom reports list
    │   │   │   │   │   │   │   │       # BE: financial-be/reports
    │   │   │   │   │   │   │   │       # GET /v1/reports/custom
    │   │   │   │   │   │   │   ├── earnings/
    │   │   │   │   │   │   │   │   ├── forecast/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Earnings forecast
    │   │   │   │   │   │   │   │   │       # - Projected earnings
    │   │   │   │   │   │   │   │   │       # - Based on current trends
    │   │   │   │   │   │   │   │   │       # - Scenario analysis
    │   │   │   │   │   │   │   │   │       # BE: users-be/analytics
    │   │   │   │   │   │   │   │   │       # GET /v1/users/me/analytics/earnings/forecast
    │   │   │   │   │   │   │   │   │       # - Projected income
    │   │   │   │   │   │   │   │   │       # - Pipeline value
    │   │   │   │   │   │   │   │   │       # BE: financial-be/analytics
    │   │   │   │   │   │   │   │   │       # GET /v1/analytics/earnings/forecast
    │   │   │   │   │   │   │   │   └── page.tsx  # Earnings analytics
    │   │   │   │   │   │   │   │       # - Monthly earnings
    │   │   │   │   │   │   │   │       # - Yearly earnings
    │   │   │   │   │   │   │   │       # - Client breakdown
    │   │   │   │   │   │   │   │       # - Tax estimates
    │   │   │   │   │   │   │   │       # BE: users-be/analytics
    │   │   │   │   │   │   │   │       # GET /v1/users/me/analytics/earnings
    │   │   │   │   │   │   │   │       # - Monthly trends
    │   │   │   │   │   │   │   │       # - Year-over-year comparison
    │   │   │   │   │   │   │   │       # BE: financial-be/analytics
    │   │   │   │   │   │   │   │       # GET /v1/analytics/earnings
    │   │   │   │   │   │   │   ├── freelancers/
    │   │   │   │   │   │   │   │   └── page.tsx  # Freelancer analytics (client)
    │   │   │   │   │   │   │   │       # - Top freelancers
    │   │   │   │   │   │   │   │       # - Freelancer retention
    │   │   │   │   │   │   │   │       # - Performance scores
    │   │   │   │   │   │   │   │       # - Cost efficiency
    │   │   │   │   │   │   │   │       # BE: users-be/analytics
    │   │   │   │   │   │   │   │       # GET /v1/users/me/analytics/freelancers
    │   │   │   │   │   │   │   ├── market-insights/
    │   │   │   │   │   │   │   │   └── page.tsx  # Market insights
    │   │   │   │   │   │   │   │       # - Skill demand trends
    │   │   │   │   │   │   │   │       # - Rate benchmarks
    │   │   │   │   │   │   │   │       # - Competition analysis
    │   │   │   │   │   │   │   │       # BE: search-be/analytics, jobs-be/analytics
    │   │   │   │   │   │   │   │       # GET /v1/analytics/market-insights
    │   │   │   │   │   │   │   ├── performance/
    │   │   │   │   │   │   │   │   └── page.tsx  # Performance analytics
    │   │   │   │   │   │   │   │       # - Response time metrics
    │   │   │   │   │   │   │   │       # - Proposal success rate
    │   │   │   │   │   │   │   │       # - Client satisfaction
    │   │   │   │   │   │   │   │       # BE: users-be/analytics, proposals-be/performance
    │   │   │   │   │   │   │   │       # GET /v1/users/me/analytics/performance
    │   │   │   │   │   │   │   ├── projects/
    │   │   │   │   │   │   │   │   └── page.tsx  # Project analytics
    │   │   │   │   │   │   │   │       # - Completion rates
    │   │   │   │   │   │   │   │       # - Average duration
    │   │   │   │   │   │   │   │       # - Budget adherence
    │   │   │   │   │   │   │   │       # - Success metrics
    │   │   │   │   │   │   │   │       # BE: users-be/analytics
    │   │   │   │   │   │   │   │       # GET /v1/users/me/analytics/projects
    │   │   │   │   │   │   │   ├── spending/
    │   │   │   │   │   │   │   │   ├── forecast/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Spending forecast
    │   │   │   │   │   │   │   │   │       # - Projected spending
    │   │   │   │   │   │   │   │   │       # - Based on active contracts
    │   │   │   │   │   │   │   │   │       # - Budget overrun warnings
    │   │   │   │   │   │   │   │   │       # BE: users-be/analytics
    │   │   │   │   │   │   │   │   │       # GET /v1/users/me/analytics/spending/forecast
    │   │   │   │   │   │   │   │   └── page.tsx  # Spending analytics (client)
    │   │   │   │   │   │   │   │       # - Monthly spending
    │   │   │   │   │   │   │   │       # - Yearly spending
    │   │   │   │   │   │   │   │       # - Freelancer breakdown
    │   │   │   │   │   │   │   │       # - Project breakdown
    │   │   │   │   │   │   │   │       # BE: users-be/analytics
    │   │   │   │   │   │   │   │       # GET /v1/users/me/analytics/spending
    │   │   │   │   │   │   │   └── page.tsx  # Analytics dashboard
    │   │   │   │   │   │   │       # - Overview metrics
    │   │   │   │   │   │   │       # - Customizable widgets
    │   │   │   │   │   │   │       # - Time range selector
    │   │   │   │   │   │   │       # - Export reports
    │   │   │   │   │   │   │       # BE: users-be/analytics
    │   │   │   │   │   │   │       # GET /v1/users/me/analytics/dashboard
    │   │   │   │   │   │   ├── availability/
    │   │   │   │   │   │   │   ├── calendar/
    │   │   │   │   │   │   │   │   └── page.tsx  # Availability calendar
    │   │   │   │   │   │   │   │       # - Mark available/busy
    │   │   │   │   │   │   │   │       # - Recurring patterns
    │   │   │   │   │   │   │   │       # - Sync with external calendar
    │   │   │   │   │   │   │   │       # BE: users-be/availability (if exists) or users-be/profile
    │   │   │   │   │   │   │   │       # GET /v1/users/me/availability
    │   │   │   │   │   │   │   │       # PUT /v1/users/me/availability
    │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   └── page.tsx  # Availability settings
    │   │   │   │   │   │   │   │       # - Working hours
    │   │   │   │   │   │   │   │       # - Timezone preferences
    │   │   │   │   │   │   │   │       # - Auto-reply settings
    │   │   │   │   │   │   │   │       # BE: users-be/settings
    │   │   │   │   │   │   │   │       # PUT /v1/users/me/availability-settings
    │   │   │   │   │   │   │   └── page.tsx  # Availability dashboard
    │   │   │   │   │   │   │       # - Current status
    │   │   │   │   │   │   │       # - Upcoming commitments
    │   │   │   │   │   │   │       # BE: users-be/availability
    │   │   │   │   │   │   │       # GET /v1/users/me/availability-overview
    │   │   │   │   │   │   ├── bidding/
    │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   └── page.tsx  # Bidding analytics
    │   │   │   │   │   │   │   │       # - Win rate
    │   │   │   │   │   │   │   │       # - Average bid amount
    │   │   │   │   │   │   │   │       # - Competition analysis
    │   │   │   │   │   │   │   │       # BE: proposals-be/bid-strategy
    │   │   │   │   │   │   │   │       # GET /v1/bid-strategies/analytics
    │   │   │   │   │   │   │   ├── auctions/
    │   │   │   │   │   │   │   │   ├── [auctionId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Auction participation
    │   │   │   │   │   │   │   │   │       # - Real-time bidding
    │   │   │   │   │   │   │   │   │       # - Bid history
    │   │   │   │   │   │   │   │   │       # - Competitor activity
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/auction
    │   │   │   │   │   │   │   │   │       # GET /v1/jobs/{job_id}/auction
    │   │   │   │   │   │   │   │   │       # POST /v1/jobs/{job_id}/auction/bid
    │   │   │   │   │   │   │   │   │       # WebSocket: Real-time updates
    │   │   │   │   │   │   │   │   └── page.tsx  # Active auctions list
    │   │   │   │   │   │   │   │       # BE: proposals-be/auction
    │   │   │   │   │   │   │   │       # GET /v1/auctions/active
    │   │   │   │   │   │   │   └── strategies/
    │   │   │   │   │   │   │       ├── [strategyId]/
    │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit bid strategy
    │   │   │   │   │   │   │       │   │       # BE: proposals-be/bid-strategy
    │   │   │   │   │   │   │       │   │       # PUT /v1/bid-strategies/{strategy_id}
    │   │   │   │   │   │   │       │   └── page.tsx  # View bid strategy details
    │   │   │   │   │   │   │       │       # BE: proposals-be/bid-strategy
    │   │   │   │   │   │   │       │       # GET /v1/bid-strategies/{strategy_id}
    │   │   │   │   │   │   │       ├── new/
    │   │   │   │   │   │   │       │   └── page.tsx  # Create new bid strategy
    │   │   │   │   │   │   │       │       # BE: proposals-be/bid-strategy
    │   │   │   │   │   │   │       │       # POST /v1/bid-strategies
    │   │   │   │   │   │   │       └── page.tsx  # Bid strategies list
    │   │   │   │   │   │   │           # - Auto-bid rules
    │   │   │   │   │   │   │           # - Price ranges
    │   │   │   │   │   │   │           # - Category targeting
    │   │   │   │   │   │   │           # BE: proposals-be/bid-strategy
    │   │   │   │   │   │   │           # GET /v1/bid-strategies
    │   │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   │   ├── budgets/
    │   │   │   │   │   │   │   │   └── page.tsx  # Budgets
    │   │   │   │   │   │   │   │       # - Create/manage budgets & alerts
    │   │   │   │   │   │   │   │       # BE: financial-be/budget
    │   │   │   │   │   │   │   │       # GET  /v1/budgets
    │   │   │   │   │   │   │   │       # POST /v1/budgets
    │   │   │   │   │   │   │   └── subscriptions/
    │   │   │   │   │   │   │       ├── history/
    │   │   │   │   │   │   │       │   └── page.tsx  # Subscription history
    │   │   │   │   │   │   │       │       # - Invoices & entitlement history
    │   │   │   │   │   │   │       │       # BE: financial-be/subscription
    │   │   │   │   │   │   │       │       # GET /v1/subscriptions/{id}/invoices
    │   │   │   │   │   │   │       └── manage/
    │   │   │   │   │   │   │           └── page.tsx  # Manage subscription
    │   │   │   │   │   │   │               # - Plan selection, upgrade/downgrade
    │   │   │   │   │   │   │               # - Cancel & dunning status
    │   │   │   │   │   │   │               # BE: financial-be/subscription
    │   │   │   │   │   │   │               # GET  /v1/subscriptions/me
    │   │   │   │   │   │   │               # POST /v1/subscriptions/change
    │   │   │   │   │   │   │               # POST /v1/subscriptions/cancel
    │   │   │   │   │   │   ├── community/
    │   │   │   │   │   │   │   ├── events/
    │   │   │   │   │   │   │   │   ├── [eventId]/
    │   │   │   │   │   │   │   │   │   ├── register/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Event registration
    │   │   │   │   │   │   │   │   │   │       # - RSVP form
    │   │   │   │   │   │   │   │   │   │       # - Add to calendar
    │   │   │   │   │   │   │   │   │   │       # BE: communications-be/events
    │   │   │   │   │   │   │   │   │   │       # POST /v1/events/{event_id}/register
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Event detail
    │   │   │   │   │   │   │   │   │       # - Info
    │   │   │   │   │   │   │   │   │       # - Attendees
    │   │   │   │   │   │   │   │   │       # - Check-in (QR code)
    │   │   │   │   │   │   │   │   │       # BE: communications-be/events
    │   │   │   │   │   │   │   │   │       # GET /v1/events/{event_id}
    │   │   │   │   │   │   │   │   ├── my-events/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # My registered events
    │   │   │   │   │   │   │   │   │       # - Attending
    │   │   │   │   │   │   │   │   │       # - Past attended
    │   │   │   │   │   │   │   │   │       # BE: communications-be/events
    │   │   │   │   │   │   │   │   │       # GET /v1/users/me/events
    │   │   │   │   │   │   │   │   └── upcoming/
    │   │   │   │   │   │   │   │       └── page.tsx  # Upcoming events
    │   │   │   │   │   │   │   │           # BE: communications-be/events
    │   │   │   │   │   │   │   │           # GET /v1/events?status=upcoming
    │   │   │   │   │   │   │   ├── forums/
    │   │   │   │   │   │   │   │   ├── [forumId]/
    │   │   │   │   │   │   │   │   │   ├── threads/
    │   │   │   │   │   │   │   │   │   │   └── [threadId]/
    │   │   │   │   │   │   │   │   │   │       ├── reply/
    │   │   │   │   │   │   │   │   │   │       │   └── page.tsx  # Reply to thread
    │   │   │   │   │   │   │   │   │   │       │       # BE: communications-be/forums
    │   │   │   │   │   │   │   │   │   │       │       # POST /v1/threads/{thread_id}/replies
    │   │   │   │   │   │   │   │   │   │       └── page.tsx  # Thread view
    │   │   │   │   │   │   │   │   │   │           # - Posts
    │   │   │   │   │   │   │   │   │   │           # - Reply
    │   │   │   │   │   │   │   │   │   │           # - Voting
    │   │   │   │   │   │   │   │   │   │           # BE: communications-be/forums (if exists) OR community service
    │   │   │   │   │   │   │   │   │   │           # GET /v1/forums/{forum_id}/threads/{thread_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Forum overview
    │   │   │   │   │   │   │   │   │       # - Thread list
    │   │   │   │   │   │   │   │   │       # - Create thread
    │   │   │   │   │   │   │   │   │       # BE: communications-be/forums
    │   │   │   │   │   │   │   │   │       # GET /v1/forums/{forum_id}/threads
    │   │   │   │   │   │   │   │   └── page.tsx  # All forums
    │   │   │   │   │   │   │   │       # - Categories
    │   │   │   │   │   │   │   │       # - Popular threads
    │   │   │   │   │   │   │   │       # BE: communications-be/forums
    │   │   │   │   │   │   │   │       # GET /v1/forums
    │   │   │   │   │   │   │   └── groups/
    │   │   │   │   │   │   │       └── [groupId]/
    │   │   │   │   │   │   │           ├── discussions/
    │   │   │   │   │   │   │           │   └── page.tsx  # Group discussions
    │   │   │   │   │   │   │           │       # - Posts feed
    │   │   │   │   │   │   │           │       # - Comment
    │   │   │   │   │   │   │           │       # BE: communications-be/discussions
    │   │   │   │   │   │   │           │       # GET /v1/groups/{group_id}/discussions
    │   │   │   │   │   │   │           ├── events/
    │   │   │   │   │   │   │           │   ├── [eventId]/
    │   │   │   │   │   │   │           │   │   └── page.tsx  # Event detail
    │   │   │   │   │   │   │           │   │       # - RSVP
    │   │   │   │   │   │   │           │   │       # - Calendar add
    │   │   │   │   │   │   │           │   │       # BE: communications-be/events OR community service
    │   │   │   │   │   │   │           │   │       # GET /v1/events/{event_id}
    │   │   │   │   │   │   │           │   └── page.tsx  # Group events
    │   │   │   │   │   │   │           │       # - Upcoming
    │   │   │   │   │   │   │           │       # - Past
    │   │   │   │   │   │   │           │       # BE: communications-be/events
    │   │   │   │   │   │   │           │       # GET /v1/groups/{group_id}/events
    │   │   │   │   │   │   │           ├── members/
    │   │   │   │   │   │   │           │   └── page.tsx  # Group members
    │   │   │   │   │   │   │           │       # - Member list
    │   │   │   │   │   │   │           │       # - Invite
    │   │   │   │   │   │   │           │       # - Roles
    │   │   │   │   │   │   │           │       # BE: users-be/groups (if exists) OR community service
    │   │   │   │   │   │   │           │       # GET /v1/groups/{group_id}/members
    │   │   │   │   │   │   │           └── page.tsx  # Group overview
    │   │   │   │   │   │   │               # - About
    │   │   │   │   │   │   │               # - Join/leave
    │   │   │   │   │   │   │               # - Activity feed
    │   │   │   │   │   │   │               # BE: users-be/groups
    │   │   │   │   │   │   │               # GET /v1/groups/{group_id}
    │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   ├── documents/
    │   │   │   │   │   │   │   │   ├── [documentId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Compliance document details
    │   │   │   │   │   │   │   │   │       # BE: storage-be/asset, admin-be/business_verification
    │   │   │   │   │   │   │   │   │       # GET /v1/compliance/documents/{document_id}
    │   │   │   │   │   │   │   │   ├── upload/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Upload compliance documents
    │   │   │   │   │   │   │   │   │       # BE: storage-be/asset, admin-be/business_verification
    │   │   │   │   │   │   │   │   │       # POST /v1/compliance/documents/upload
    │   │   │   │   │   │   │   │   └── page.tsx  # Compliance documents list
    │   │   │   │   │   │   │   │       # BE: admin-be/business_verification
    │   │   │   │   │   │   │   │       # GET /v1/compliance/documents
    │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   ├── tax-summary/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Annual tax summary
    │   │   │   │   │   │   │   │   │       # BE: financial-be/tax
    │   │   │   │   │   │   │   │   │       # GET /v1/tax/reports/annual-summary
    │   │   │   │   │   │   │   │   └── page.tsx  # Compliance reports
    │   │   │   │   │   │   │   │       # - Income reports
    │   │   │   │   │   │   │   │       # - Tax withholding
    │   │   │   │   │   │   │   │       # - Payment history
    │   │   │   │   │   │   │   │       # BE: financial-be/reports
    │   │   │   │   │   │   │   │       # GET /v1/reports/compliance
    │   │   │   │   │   │   │   └── tax-profile/
    │   │   │   │   │   │   │       ├── edit/
    │   │   │   │   │   │   │       │   └── page.tsx  # Edit tax profile
    │   │   │   │   │   │   │       │       # BE: users-be/compliance, financial-be/tax
    │   │   │   │   │   │   │       │       # PUT /v1/users/me/compliance/tax-profile
    │   │   │   │   │   │   │       └── page.tsx  # Tax profile overview
    │   │   │   │   │   │   │           # - Tax ID
    │   │   │   │   │   │   │           # - Tax forms
    │   │   │   │   │   │   │           # - Withholding settings
    │   │   │   │   │   │   │           # BE: users-be/compliance
    │   │   │   │   │   │   │           # GET /v1/users/me/compliance/tax-profile
    │   │   │   │   │   │   ├── connects/
    │   │   │   │   │   │   │   ├── purchase/
    │   │   │   │   │   │   │   │   └── page.tsx  # Purchase connects
    │   │   │   │   │   │   │   │       # - Select package
    │   │   │   │   │   │   │   │       # - Payment processing
    │   │   │   │   │   │   │   │       # BE: proposals-be/connect, financial-be/payment
    │   │   │   │   │   │   │   │       # GET /v1/connects/packages
    │   │   │   │   │   │   │   │       # POST /v1/connects/purchase
    │   │   │   │   │   │   │   ├── usage/
    │   │   │   │   │   │   │   │   └── page.tsx  # Connects usage analytics
    │   │   │   │   │   │   │   │       # - Spending patterns
    │   │   │   │   │   │   │   │       # - Refund history
    │   │   │   │   │   │   │   │       # - ROI tracking
    │   │   │   │   │   │   │   │       # BE: proposals-be/connect
    │   │   │   │   │   │   │   │       # GET /v1/connects/usage-analytics
    │   │   │   │   │   │   │   └── page.tsx  # Connects dashboard
    │   │   │   │   │   │   │       # - Current balance
    │   │   │   │   │   │   │       # - Transaction history
    │   │   │   │   │   │   │       # - Refund requests
    │   │   │   │   │   │   │       # BE: proposals-be/connect
    │   │   │   │   │   │   │       # GET /v1/connects
    │   │   │   │   │   │   │       # GET /v1/connects/balance
    │   │   │   │   │   │   ├── contracts/  # Contracts management
    │   │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   │   ├── amendments/
    │   │   │   │   │   │   │   │   │   ├── [amendmentId]/
    │   │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve/reject amendment
    │   │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/amendment
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/amendments/{amendment_id}/approve
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/amendments/{amendment_id}/reject
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Amendment detail
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/amendment
    │   │   │   │   │   │   │   │   │   │       # GET /v1/amendments/{amendment_id}
    │   │   │   │   │   │   │   │   │   │       # - View changes
    │   │   │   │   │   │   │   │   │   │       # - Approval status
    │   │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/amendments/{amendment_id}
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create amendment
    │   │   │   │   │   │   │   │   │   │       # - Modify terms
    │   │   │   │   │   │   │   │   │   │       # - Add milestones
    │   │   │   │   │   │   │   │   │   │       # - Change scope
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/amendment
    │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/amendments
    │   │   │   │   │   │   │   │   │   ├── propose/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Propose amendment
    │   │   │   │   │   │   │   │   │   │       # - Change description
    │   │   │   │   │   │   │   │   │   │       # - Updated terms
    │   │   │   │   │   │   │   │   │   │       # - Reason
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/amendment
    │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/amendments
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Contract amendments list
    │   │   │   │   │   │   │   │   │       # - Proposed amendments
    │   │   │   │   │   │   │   │   │       # - Accepted amendments
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/amendment
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/amendments
    │   │   │   │   │   │   │   │   │       # Amendments list
    │   │   │   │   │   │   │   │   ├── audit-trail/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Complete contract audit trail
    │   │   │   │   │   │   │   │   │       # - All changes
    │   │   │   │   │   │   │   │   │       # - Approval history
    │   │   │   │   │   │   │   │   │       # - Access logs
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/contract, utility-be/audit
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/audit-trail
    │   │   │   │   │   │   │   │   ├── change-orders/
    │   │   │   │   │   │   │   │   │   ├── [orderId]/
    │   │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve change order
    │   │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/change_order
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/change-orders/{order_id}/approve
    │   │   │   │   │   │   │   │   │   │   ├── reject/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reject change order
    │   │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/change_order
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/change-orders/{order_id}/reject
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Change order details
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/change_order
    │   │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/change-orders/{order_id}
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create change order
    │   │   │   │   │   │   │   │   │   │       # - Scope modifications
    │   │   │   │   │   │   │   │   │   │       # - Budget adjustments
    │   │   │   │   │   │   │   │   │   │       # - Timeline changes
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/change_order
    │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/change-orders
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Change orders list
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/change_order
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/change-orders
    │   │   │   │   │   │   │   │   ├── complete/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Complete contract
    │   │   │   │   │   │   │   │   │       # - Confirm all deliverables received
    │   │   │   │   │   │   │   │   │       # - Leave review
    │   │   │   │   │   │   │   │   │       # - Final payment release
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/contract
    │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/complete
    │   │   │   │   │   │   │   │   │       # Publishes: ContractCompleted event
    │   │   │   │   │   │   │   │   │       # Redirects to: /reviews/create
    │   │   │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Contract compliance tracking
    │   │   │   │   │   │   │   │   │       # - KPIs monitoring
    │   │   │   │   │   │   │   │   │       # - SLA compliance
    │   │   │   │   │   │   │   │   │       # - Penalty tracking
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/compliance
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/compliance
    │   │   │   │   │   │   │   │   ├── deliverables/
    │   │   │   │   │   │   │   │   │   ├── [deliverableId]/
    │   │   │   │   │   │   │   │   │   │   ├── accept/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Accept deliverable
    │   │   │   │   │   │   │   │   │   │   │       # - Review checklist
    │   │   │   │   │   │   │   │   │   │   │       # - Release milestone payment
    │   │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/accept
    │   │   │   │   │   │   │   │   │   │   ├── reject/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reject deliverable
    │   │   │   │   │   │   │   │   │   │   │       # - Provide feedback
    │   │   │   │   │   │   │   │   │   │   │       # - Request revisions
    │   │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/reject
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Deliverable detail
    │   │   │   │   │   │   │   │   │   │       # - File preview
    │   │   │   │   │   │   │   │   │   │       # - Download
    │   │   │   │   │   │   │   │   │   │       # - Approval status
    │   │   │   │   │   │   │   │   │   │       # - Feedback
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable
    │   │   │   │   │   │   │   │   │   │       # GET /v1/deliverables/{deliverable_id}
    │   │   │   │   │   │   │   │   │   │       # BE: storage-be
    │   │   │   │   │   │   │   │   │   │       # GET /v1/storage/download/{file_id}
    │   │   │   │   │   │   │   │   │   │       # - Files and description
    │   │   │   │   │   │   │   │   │   │       # - Review history
    │   │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}
    │   │   │   │   │   │   │   │   │   ├── submit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Submit deliverable
    │   │   │   │   │   │   │   │   │   │       # - Upload files
    │   │   │   │   │   │   │   │   │   │       # - Add notes
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable, storage-be/asset
    │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables
    │   │   │   │   │   │   │   │   │   └── page.tsx  # All deliverables
    │   │   │   │   │   │   │   │   │       # - List all submitted deliverables
    │   │   │   │   │   │   │   │   │       # - Status (pending, approved, rejected)
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables
    │   │   │   │   │   │   │   │   │       # Deliverables list
    │   │   │   │   │   │   │   │   ├── details/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Full contract details
    │   │   │   │   │   │   │   │   │       # - Contract terms
    │   │   │   │   │   │   │   │   │       # - Scope of work (SOW)
    │   │   │   │   │   │   │   │   │       # - Payment terms
    │   │   │   │   │   │   │   │   │       # - Deadlines
    │   │   │   │   │   │   │   │   │       # - Special clauses
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/contract
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/details
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/sow
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/sow
    │   │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   │   │   │   ├── escalate/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Escalate to admin/mediation
    │   │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/dispute
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/disputes/{dispute_id}/escalate
    │   │   │   │   │   │   │   │   │   │   │       # BE: admin-be
    │   │   │   │   │   │   │   │   │   │   │       # Creates mediation case
    │   │   │   │   │   │   │   │   │   │   ├── respond/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Respond to dispute
    │   │   │   │   │   │   │   │   │   │   │       # - Response message
    │   │   │   │   │   │   │   │   │   │   │       # - Additional evidence
    │   │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/dispute
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/disputes/{dispute_id}/responses
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute detail
    │   │   │   │   │   │   │   │   │   │       # - Dispute timeline
    │   │   │   │   │   │   │   │   │   │       # - Messages/responses
    │   │   │   │   │   │   │   │   │   │       # - Evidence
    │   │   │   │   │   │   │   │   │   │       # - Admin notes (if assigned)
    │   │   │   │   │   │   │   │   │   │       # - Resolution status
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/dispute
    │   │   │   │   │   │   │   │   │   │       # GET /v1/disputes/{dispute_id}
    │   │   │   │   │   │   │   │   │   ├── open/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Open a dispute
    │   │   │   │   │   │   │   │   │   │       # - Dispute reason
    │   │   │   │   │   │   │   │   │   │       # - Description
    │   │   │   │   │   │   │   │   │   │       # - Evidence upload
    │   │   │   │   │   │   │   │   │   │       # - Desired resolution
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/dispute
    │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/disputes
    │   │   │   │   │   │   │   │   │   │       # BE: storage-be/uploads
    │   │   │   │   │   │   │   │   │   │       # Publishes: DisputeOpened event
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Disputes list
    │   │   │   │   │   │   │   │   │       # - Active disputes
    │   │   │   │   │   │   │   │   │       # - Resolved disputes
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/dispute
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/disputes
    │   │   │   │   │   │   │   │   ├── extensions/
    │   │   │   │   │   │   │   │   │   ├── [extensionId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Extension request detail
    │   │   │   │   │   │   │   │   │   │       # - Approve/reject
    │   │   │   │   │   │   │   │   │   │       # - Negotiation
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/extension
    │   │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/extensions/{extension_id}
    │   │   │   │   │   │   │   │   │   ├── request/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Request extension
    │   │   │   │   │   │   │   │   │   │       # - New deadline
    │   │   │   │   │   │   │   │   │   │       # - Justification
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/extension
    │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/extensions
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Extensions list
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/extension
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/extensions
    │   │   │   │   │   │   │   │   ├── feedback/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Ongoing feedback
    │   │   │   │   │   │   │   │   │       # - Mid-contract feedback
    │   │   │   │   │   │   │   │   │       # - Performance notes
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/feedback
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/feedback
    │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/feedback
    │   │   │   │   │   │   │   │   ├── invoices/
    │   │   │   │   │   │   │   │   │   ├── [invoiceId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Invoice detail
    │   │   │   │   │   │   │   │   │   │       # - Line items
    │   │   │   │   │   │   │   │   │   │       # - Payment status
    │   │   │   │   │   │   │   │   │   │       # - Download PDF
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/invoice
    │   │   │   │   │   │   │   │   │   │       # GET /v1/invoices/{invoice_id}
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create invoice
    │   │   │   │   │   │   │   │   │   │       # - Add line items
    │   │   │   │   │   │   │   │   │   │       # - Tax calculation
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/invoice
    │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/invoices
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Invoices list
    │   │   │   │   │   │   │   │   │       # BE: financial-be/invoice
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/invoices
    │   │   │   │   │   │   │   │   ├── messages/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Contract-specific messages
    │   │   │   │   │   │   │   │   │       # - Threaded conversation
    │   │   │   │   │   │   │   │   │       # - File sharing
    │   │   │   │   │   │   │   │   │       # - Quick links to milestones/deliverables
    │   │   │   │   │   │   │   │   │       # BE: communications-be/conversations
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/conversation
    │   │   │   │   │   │   │   │   ├── milestones/
    │   │   │   │   │   │   │   │   │   ├── [milestoneId]/
    │   │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve milestone (client)
    │   │   │   │   │   │   │   │   │   │   │       # - Review deliverables
    │   │   │   │   │   │   │   │   │   │   │       # - Accept/request changes
    │   │   │   │   │   │   │   │   │   │   │       # - Approval notes
    │   │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/milestone
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/milestones/{milestone_id}/approve
    │   │   │   │   │   │   │   │   │   │   │       # Publishes: MilestoneApproved event
    │   │   │   │   │   │   │   │   │   │   │       # Triggers: Escrow release (financial-be)
    │   │   │   │   │   │   │   │   │   │   ├── dispute/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute milestone
    │   │   │   │   │   │   │   │   │   │   │       # - Reason for dispute
    │   │   │   │   │   │   │   │   │   │   │       # - Evidence upload
    │   │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/dispute
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/milestones/{milestone_id}/dispute
    │   │   │   │   │   │   │   │   │   │   └── submit/
    │   │   │   │   │   │   │   │   │   │       └── page.tsx  # Submit deliverables (freelancer)
    │   │   │   │   │   │   │   │   │   │           # - Upload files
    │   │   │   │   │   │   │   │   │   │           # - Completion notes
    │   │   │   │   │   │   │   │   │   │           # BE: contracts-be/deliverable
    │   │   │   │   │   │   │   │   │   │           # POST /v1/milestones/{milestone_id}/deliverables
    │   │   │   │   │   │   │   │   │   │           # BE: storage-be/uploads
    │   │   │   │   │   │   │   │   │   │           # Publishes: MilestoneCompleted event
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create milestone (if contract allows)
    │   │   │   │   │   │   │   │   │   │       # - Milestone title
    │   │   │   │   │   │   │   │   │   │       # - Description
    │   │   │   │   │   │   │   │   │   │       # - Amount
    │   │   │   │   │   │   │   │   │   │       # - Due date
    │   │   │   │   │   │   │   │   │   │       # - Deliverables
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/milestone
    │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/milestones
    │   │   │   │   │   │   │   │   │   │       # Publishes: MilestoneCreated event
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Milestones list
    │   │   │   │   │   │   │   │   │       # - All milestones
    │   │   │   │   │   │   │   │   │       # - Status (pending, in_progress, completed)
    │   │   │   │   │   │   │   │   │       # - Amount & due date
    │   │   │   │   │   │   │   │   │       # - Approval status
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/milestone
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/milestones
    │   │   │   │   │   │   │   │   ├── pause/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pause contract
    │   │   │   │   │   │   │   │   │       # - Reason
    │   │   │   │   │   │   │   │   │       # - Expected resume date
    │   │   │   │   │   │   │   │   │       # - Notify other party
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/contract
    │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/pause
    │   │   │   │   │   │   │   │   │       # Publishes: ContractPaused event
    │   │   │   │   │   │   │   │   │       # - Reason selection
    │   │   │   │   │   │   │   │   │       # - Estimated resume date
    │   │   │   │   │   │   │   │   ├── payments/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Contract payments
    │   │   │   │   │   │   │   │   │       # - Payment schedule
    │   │   │   │   │   │   │   │   │       # - Escrow status
    │   │   │   │   │   │   │   │   │       # - Released payments
    │   │   │   │   │   │   │   │   │       # - Pending payments
    │   │   │   │   │   │   │   │   │       # BE: financial-be/escrow
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/escrow
    │   │   │   │   │   │   │   │   │       # BE: financial-be/payment
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/payments
    │   │   │   │   │   │   │   │   ├── resume/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Resume paused contract
    │   │   │   │   │   │   │   │   │       # - Confirm resume
    │   │   │   │   │   │   │   │   │       # - Adjust timeline
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/contract
    │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/resume
    │   │   │   │   │   │   │   │   ├── sow/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit SOW (before signing)
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/sow
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/contracts/{contract_id}/sow
    │   │   │   │   │   │   │   │   │   └── page.tsx  # SOW detail view
    │   │   │   │   │   │   │   │   │       # - Scope of work
    │   │   │   │   │   │   │   │   │       # - Deliverables
    │   │   │   │   │   │   │   │   │       # - Timeline
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/sow
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/sow
    │   │   │   │   │   │   │   │   ├── terminate/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Terminate contract
    │   │   │   │   │   │   │   │   │       # - Termination reason
    │   │   │   │   │   │   │   │   │       # - Early termination terms
    │   │   │   │   │   │   │   │   │       # - Final deliverables
    │   │   │   │   │   │   │   │   │       # - Escrow settlement
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/termination
    │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/terminate
    │   │   │   │   │   │   │   │   │       # Publishes: ContractTerminated event
    │   │   │   │   │   │   │   │   │       # - Final settlement
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/contract
    │   │   │   │   │   │   │   │   ├── timesheet/
    │   │   │   │   │   │   │   │   │   ├── [timesheetId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Timesheet detail
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/timesheet
    │   │   │   │   │   │   │   │   │   │       # GET /v1/timesheets/{timesheet_id}
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve timesheet (client)
    │   │   │   │   │   │   │   │   │   │       # - Review hours
    │   │   │   │   │   │   │   │   │   │       # - Approve/request changes
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/timesheet
    │   │   │   │   │   │   │   │   │   │       # POST /v1/timesheets/{timesheet_id}/approve
    │   │   │   │   │   │   │   │   │   ├── submit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Submit timesheet (freelancer)
    │   │   │   │   │   │   │   │   │   │       # - Hours worked per day
    │   │   │   │   │   │   │   │   │   │       # - Task descriptions
    │   │   │   │   │   │   │   │   │   │       # - Billable/non-billable
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/timesheet
    │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/timesheets
    │   │   │   │   │   │   │   │   │   │       # Publishes: TimesheetSubmitted event
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Timesheet view (hourly contracts)
    │   │   │   │   │   │   │   │   │       # - Weekly/monthly view
    │   │   │   │   │   │   │   │   │       # - Total hours
    │   │   │   │   │   │   │   │   │       # - Approval status
    │   │   │   │   │   │   │   │   │       # - Submit for review
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/timesheet
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/timesheets
    │   │   │   │   │   │   │   │   ├── work-diary/
    │   │   │   │   │   │   │   │   │   ├── [date]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Work diary for specific date
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/work_diary
    │   │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/work-diary?date={date}
    │   │   │   │   │   │   │   │   │   ├── add-entry/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Add work diary entry (freelancer)
    │   │   │   │   │   │   │   │   │   │       # - Date & time
    │   │   │   │   │   │   │   │   │   │       # - Hours worked
    │   │   │   │   │   │   │   │   │   │       # - Description
    │   │   │   │   │   │   │   │   │   │       # - Upload screenshot (optional)
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/work_diary
    │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/work-diary/entries
    │   │   │   │   │   │   │   │   │   │       # BE: storage-be/uploads
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Work diary overview
    │   │   │   │   │   │   │   │   │       # - Daily activity logs
    │   │   │   │   │   │   │   │   │       # - Screenshots (if enabled)
    │   │   │   │   │   │   │   │   │       # - Productivity metrics
    │   │   │   │   │   │   │   │   │       # - Calendar view
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/work_diary
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/work-diary
    │   │   │   │   │   │   │   │   ├── workdiary/
    │   │   │   │   │   │   │   │   │   ├── manual-time/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Add manual time entry
    │   │   │   │   │   │   │   │   │   │       # - Date/time range
    │   │   │   │   │   │   │   │   │   │       # - Description
    │   │   │   │   │   │   │   │   │   │       # - Approval required
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/workdiary
    │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/workdiary/manual-time
    │   │   │   │   │   │   │   │   │   └── reports/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Work diary reports
    │   │   │   │   │   │   │   │   │           # - Weekly summaries
    │   │   │   │   │   │   │   │   │           # - Time breakdowns
    │   │   │   │   │   │   │   │   │           # BE: contracts-be/workdiary
    │   │   │   │   │   │   │   │   │           # GET /v1/contracts/{contract_id}/workdiary/reports
    │   │   │   │   │   │   │   │   └── page.tsx  # Contract overview
    │   │   │   │   │   │   │   │       # - Contract details
    │   │   │   │   │   │   │   │       # - Parties involved
    │   │   │   │   │   │   │   │       # - Budget/rate
    │   │   │   │   │   │   │   │       # - Timeline
    │   │   │   │   │   │   │   │       # - Status
    │   │   │   │   │   │   │   │       # - Quick actions (message, submit work, etc.)
    │   │   │   │   │   │   │   │       # BE: contracts-be/contract
    │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}
    │   │   │   │   │   │   │   ├── active/
    │   │   │   │   │   │   │   │   └── page.tsx  # Active contracts only
    │   │   │   │   │   │   │   │       # BE: contracts-be/contract
    │   │   │   │   │   │   │   │       # GET /v1/contracts?status=active
    │   │   │   │   │   │   │   ├── archived/
    │   │   │   │   │   │   │   │   └── page.tsx  # Archived contracts
    │   │   │   │   │   │   │   │       # - Completed contracts
    │   │   │   │   │   │   │   │       # - Historical data
    │   │   │   │   │   │   │   │       # BE: contracts-be/contract
    │   │   │   │   │   │   │   │       # GET /v1/contracts?status=archived
    │   │   │   │   │   │   │   ├── benchmarking/
    │   │   │   │   │   │   │   │   └── page.tsx  # Contract performance benchmarking
    │   │   │   │   │   │   │   │       # - Industry comparisons
    │   │   │   │   │   │   │   │       # - Performance metrics
    │   │   │   │   │   │   │   │       # - Best practices
    │   │   │   │   │   │   │   │       # BE: contracts-be/analytics
    │   │   │   │   │   │   │   │       # GET /v1/contracts/benchmarking
    │   │   │   │   │   │   │   ├── calendar/
    │   │   │   │   │   │   │   │   └── page.tsx  # Contracts calendar view
    │   │   │   │   │   │   │   │       # - Milestone timeline
    │   │   │   │   │   │   │   │       # - Payment schedules
    │   │   │   │   │   │   │   │       # - Deliverable deadlines
    │   │   │   │   │   │   │   │       # BE: contracts-be/contract, contracts-be/milestone
    │   │   │   │   │   │   │   │       # GET /v1/contracts/calendar
    │   │   │   │   │   │   │   ├── completed/
    │   │   │   │   │   │   │   │   └── page.tsx  # Completed contracts
    │   │   │   │   │   │   │   │       # BE: contracts-be/contract
    │   │   │   │   │   │   │   │       # GET /v1/contracts?status=completed
    │   │   │   │   │   │   │   ├── paused/
    │   │   │   │   │   │   │   │   └── page.tsx  # Paused contracts list
    │   │   │   │   │   │   │   │       # - Temporarily paused
    │   │   │   │   │   │   │   │       # - Resume options
    │   │   │   │   │   │   │   │       # BE: contracts-be/contract
    │   │   │   │   │   │   │   │       # GET /v1/contracts?status=paused
    │   │   │   │   │   │   │   ├── recurring/
    │   │   │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   │   │   └── renew/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Renew recurring contract
    │   │   │   │   │   │   │   │   │           # BE: contracts-be/recurring
    │   │   │   │   │   │   │   │   │           # POST /v1/contracts/{contract_id}/renew
    │   │   │   │   │   │   │   │   └── page.tsx  # Recurring contracts
    │   │   │   │   │   │   │   │       # - List recurring contracts
    │   │   │   │   │   │   │   │       # - Renewal schedule
    │   │   │   │   │   │   │   │       # BE: contracts-be/recurring
    │   │   │   │   │   │   │   │       # GET /v1/contracts/recurring
    │   │   │   │   │   │   │   ├── templates/
    │   │   │   │   │   │   │   │   ├── [templateId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit template
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/template
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/contract-templates/{template_id}
    │   │   │   │   │   │   │   │   │   │       # Edit contract template
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/contracts/templates/{template_id}
    │   │   │   │   │   │   │   │   │   └── use/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Use template for new contract
    │   │   │   │   │   │   │   │   │           # BE: contracts-be/template
    │   │   │   │   │   │   │   │   │           # GET /v1/contracts/templates/{template_id}
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create contract template
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/template
    │   │   │   │   │   │   │   │   │       # POST /v1/contract-templates
    │   │   │   │   │   │   │   │   └── page.tsx  # Contract templates (for recurring work)
    │   │   │   │   │   │   │   │       # BE: contracts-be/template
    │   │   │   │   │   │   │   │       # GET /v1/contract-templates
    │   │   │   │   │   │   │   │       # Contract templates list
    │   │   │   │   │   │   │   │       # - Standard contract templates
    │   │   │   │   │   │   │   │       # - Create from existing
    │   │   │   │   │   │   │   │       # GET /v1/contracts/templates
    │   │   │   │   │   │   │   └── page.tsx  # Contracts list
    │   │   │   │   │   │   │       # - Active contracts
    │   │   │   │   │   │   │       # - Completed contracts
    │   │   │   │   │   │   │       # - Filter by status
    │   │   │   │   │   │   │       # - Sort options
    │   │   │   │   │   │   │       # BE: contracts-be/contract
    │   │   │   │   │   │   │       # GET /v1/contracts?status=active
    │   │   │   │   │   │   ├── dashboard/  # Dashboard home (role-based view)
    │   │   │   │   │   │   │   └── page.tsx  # Main dashboard
    │   │   │   │   │   │   │       # Freelancer view:
    │   │   │   │   │   │   │       # - Active proposals
    │   │   │   │   │   │   │       # - Active contracts
    │   │   │   │   │   │   │       # - Earnings overview
    │   │   │   │   │   │   │       # - Job recommendations
    │   │   │   │   │   │   │       # - Profile completion score
    │   │   │   │   │   │   │       # Client view:
    │   │   │   │   │   │   │       # - Active jobs
    │   │   │   │   │   │   │       # - Spending overview
    │   │   │   │   │   │   │       # - Recent proposals
    │   │   │   │   │   │   │       # - Talent recommendations
    │   │   │   │   │   │   │       # BE: users-be/user
    │   │   │   │   │   │   │       # GET /v1/users/me
    │   │   │   │   │   │   │       # BE: Multiple services for dashboard data
    │   │   │   │   │   │   │       # GET /v1/analytics/dashboard (analytics service)
    │   │   │   │   │   │   │       # GET /v1/jobs/my-jobs (jobs-be)
    │   │   │   │   │   │   │       # GET /v1/proposals/my-proposals (proposals-be)
    │   │   │   │   │   │   │       # GET /v1/contracts/active (contracts-be)
    │   │   │   │   │   │   ├── deliverables/
    │   │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   │   ├── [deliverableId]/
    │   │   │   │   │   │   │   │   │   ├── review/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Review deliverable (client)
    │   │   │   │   │   │   │   │   │   │       # - Approve/reject
    │   │   │   │   │   │   │   │   │   │       # - Request changes
    │   │   │   │   │   │   │   │   │   │       # - Add comments
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable
    │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/review
    │   │   │   │   │   │   │   │   │   ├── revisions/
    │   │   │   │   │   │   │   │   │   │   ├── [revisionId]/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Revision detail
    │   │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable
    │   │   │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/revisions/{revision_id}
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Revision history
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable
    │   │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/revisions
    │   │   │   │   │   │   │   │   │   ├── upload/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Upload new version
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable, storage-be/asset
    │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/upload
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Deliverable details
    │   │   │   │   │   │   │   │   │       # - File viewer
    │   │   │   │   │   │   │   │   │       # - Download
    │   │   │   │   │   │   │   │   │       # - Metadata
    │   │   │   │   │   │   │   │   │       # - Comments thread
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable, storage-be/asset
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}
    │   │   │   │   │   │   │   │   ├── new/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Submit new deliverable
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable, storage-be/asset
    │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables
    │   │   │   │   │   │   │   │   └── page.tsx  # Contract deliverables list
    │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable
    │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables
    │   │   │   │   │   │   │   ├── pending-review/
    │   │   │   │   │   │   │   │   └── page.tsx  # Deliverables pending client review
    │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable
    │   │   │   │   │   │   │   │       # GET /v1/deliverables/pending-review
    │   │   │   │   │   │   │   └── page.tsx  # All deliverables overview
    │   │   │   │   │   │   │       # BE: contracts-be/deliverable
    │   │   │   │   │   │   │       # GET /v1/deliverables
    │   │   │   │   │   │   ├── financial/
    │   │   │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Billing history
    │   │   │   │   │   │   │   │   │       # - Past invoices
    │   │   │   │   │   │   │   │   │       # - Payment receipts
    │   │   │   │   │   │   │   │   │       # BE: financial-be/billing
    │   │   │   │   │   │   │   │   │       # GET /v1/billing/history
    │   │   │   │   │   │   │   │   ├── payment-methods/
    │   │   │   │   │   │   │   │   │   ├── [methodId]/
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit payment method
    │   │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payment-method
    │   │   │   │   │   │   │   │   │   │   │       # PUT /v1/payment-methods/{method_id}
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Payment method detail
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payment-method
    │   │   │   │   │   │   │   │   │   │       # GET /v1/payment-methods/{method_id}
    │   │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Add payment method
    │   │   │   │   │   │   │   │   │   │       # - Card/bank details
    │   │   │   │   │   │   │   │   │   │       # - Verification
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payment-method
    │   │   │   │   │   │   │   │   │   │       # POST /v1/payment-methods
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Payment methods list
    │   │   │   │   │   │   │   │   │       # - Manage cards/banks
    │   │   │   │   │   │   │   │   │       # - Set default
    │   │   │   │   │   │   │   │   │       # BE: financial-be/payment-method
    │   │   │   │   │   │   │   │   │       # GET /v1/payment-methods
    │   │   │   │   │   │   │   │   └── subscriptions/
    │   │   │   │   │   │   │   │       └── page.tsx  # Subscription billing
    │   │   │   │   │   │   │   │           # - Current plan
    │   │   │   │   │   │   │   │           # - Billing cycle
    │   │   │   │   │   │   │   │           # - Upgrade/downgrade
    │   │   │   │   │   │   │   │           # BE: financial-be/subscription
    │   │   │   │   │   │   │   │           # GET /v1/subscriptions/billing
    │   │   │   │   │   │   │   ├── chargebacks/
    │   │   │   │   │   │   │   │   ├── [chargebackId]/
    │   │   │   │   │   │   │   │   │   ├── respond/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Respond to chargeback
    │   │   │   │   │   │   │   │   │   │       # - Upload evidence
    │   │   │   │   │   │   │   │   │   │       # - Add defense statement
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/chargeback
    │   │   │   │   │   │   │   │   │   │       # POST /v1/financial/chargebacks/{chargeback_id}/respond
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Chargeback details
    │   │   │   │   │   │   │   │   │       # BE: financial-be/chargeback
    │   │   │   │   │   │   │   │   │       # GET /v1/financial/chargebacks/{chargeback_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Chargebacks list
    │   │   │   │   │   │   │   │       # - Active disputes
    │   │   │   │   │   │   │   │       # - Resolution status
    │   │   │   │   │   │   │   │       # BE: financial-be/chargeback
    │   │   │   │   │   │   │   │       # GET /v1/financial/chargebacks
    │   │   │   │   │   │   │   ├── cost-centers/
    │   │   │   │   │   │   │   │   ├── [centerId]/
    │   │   │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Cost center analytics
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/cost_center, financial-be/analytics
    │   │   │   │   │   │   │   │   │   │       # GET /v1/financial/cost-centers/{center_id}/analytics
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit cost center
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/cost_center
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/financial/cost-centers/{center_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Cost center details
    │   │   │   │   │   │   │   │   │       # BE: financial-be/cost_center
    │   │   │   │   │   │   │   │   │       # GET /v1/financial/cost-centers/{center_id}
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create cost center
    │   │   │   │   │   │   │   │   │       # BE: financial-be/cost_center
    │   │   │   │   │   │   │   │   │       # POST /v1/financial/cost-centers
    │   │   │   │   │   │   │   │   └── page.tsx  # Cost centers list
    │   │   │   │   │   │   │   │       # - Department/project budgets
    │   │   │   │   │   │   │   │       # - Spend tracking
    │   │   │   │   │   │   │   │       # BE: financial-be/cost_center
    │   │   │   │   │   │   │   │       # GET /v1/financial/cost-centers
    │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   │   │   ├── evidence/
    │   │   │   │   │   │   │   │   │   │   ├── submit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Submit evidence
    │   │   │   │   │   │   │   │   │   │   │       # - Upload documents
    │   │   │   │   │   │   │   │   │   │   │       # - Add description
    │   │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/dispute, storage-be/asset
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/disputes/{dispute_id}/evidence
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Evidence list
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/dispute
    │   │   │   │   │   │   │   │   │   │       # GET /v1/disputes/{dispute_id}/evidence
    │   │   │   │   │   │   │   │   │   ├── messages/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute messages
    │   │   │   │   │   │   │   │   │   │       # - Communication thread
    │   │   │   │   │   │   │   │   │   │       # - Mediation chat
    │   │   │   │   │   │   │   │   │   │       # BE: communications-be/conversation
    │   │   │   │   │   │   │   │   │   │       # GET /v1/disputes/{dispute_id}/messages
    │   │   │   │   │   │   │   │   │   ├── resolution/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute resolution
    │   │   │   │   │   │   │   │   │   │       # - Accept/reject resolution
    │   │   │   │   │   │   │   │   │   │       # - Final settlement
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/dispute
    │   │   │   │   │   │   │   │   │   │       # POST /v1/disputes/{dispute_id}/resolution
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute detail
    │   │   │   │   │   │   │   │   │       # - Status and timeline
    │   │   │   │   │   │   │   │   │       # - Evidence
    │   │   │   │   │   │   │   │   │       # - Messages
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/dispute
    │   │   │   │   │   │   │   │   │       # GET /v1/disputes/{dispute_id}
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create dispute
    │   │   │   │   │   │   │   │   │       # - Select contract
    │   │   │   │   │   │   │   │   │       # - Issue description
    │   │   │   │   │   │   │   │   │       # - Initial evidence
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/dispute
    │   │   │   │   │   │   │   │   │       # POST /v1/disputes
    │   │   │   │   │   │   │   │   └── page.tsx  # Disputes list
    │   │   │   │   │   │   │   │       # - Open disputes
    │   │   │   │   │   │   │   │       # - Resolved disputes
    │   │   │   │   │   │   │   │       # BE: contracts-be/dispute
    │   │   │   │   │   │   │   │       # GET /v1/disputes
    │   │   │   │   │   │   │   ├── escrow/
    │   │   │   │   │   │   │   │   ├── [escrowId]/
    │   │   │   │   │   │   │   │   │   ├── fund/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Fund escrow
    │   │   │   │   │   │   │   │   │   │       # - Select payment method
    │   │   │   │   │   │   │   │   │   │       # - Amount confirmation
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/escrow
    │   │   │   │   │   │   │   │   │   │       # POST /v1/escrow/{escrow_id}/fund
    │   │   │   │   │   │   │   │   │   ├── release/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Release escrow funds
    │   │   │   │   │   │   │   │   │   │       # - Release to freelancer
    │   │   │   │   │   │   │   │   │   │       # - Partial/full release
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/escrow
    │   │   │   │   │   │   │   │   │   │       # POST /v1/escrow/{escrow_id}/release
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Escrow detail
    │   │   │   │   │   │   │   │   │       # - Balance and transactions
    │   │   │   │   │   │   │   │   │       # - Release history
    │   │   │   │   │   │   │   │   │       # BE: financial-be/escrow
    │   │   │   │   │   │   │   │   │       # GET /v1/escrow/{escrow_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Escrow accounts list
    │   │   │   │   │   │   │   │       # - Active escrows
    │   │   │   │   │   │   │   │       # - Transaction history
    │   │   │   │   │   │   │   │       # BE: financial-be/escrow
    │   │   │   │   │   │   │   │       # GET /v1/escrow
    │   │   │   │   │   │   │   ├── forecasting/
    │   │   │   │   │   │   │   │   └── page.tsx  # Revenue/expense forecasting
    │   │   │   │   │   │   │   │       # - Projected earnings
    │   │   │   │   │   │   │   │       # - Cash flow predictions
    │   │   │   │   │   │   │   │       # - Scenario planning
    │   │   │   │   │   │   │   │       # BE: financial-be/analytics
    │   │   │   │   │   │   │   │       # GET /v1/financial/forecasts
    │   │   │   │   │   │   │   ├── invoices/
    │   │   │   │   │   │   │   │   ├── [invoiceId]/
    │   │   │   │   │   │   │   │   │   ├── download/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Download invoice PDF
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/invoice
    │   │   │   │   │   │   │   │   │   │       # GET /v1/invoices/{invoice_id}/download
    │   │   │   │   │   │   │   │   │   ├── pay/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Pay invoice
    │   │   │   │   │   │   │   │   │   │       # - Select payment method
    │   │   │   │   │   │   │   │   │   │       # - Confirm payment
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payment
    │   │   │   │   │   │   │   │   │   │       # POST /v1/invoices/{invoice_id}/pay
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Invoice detail (already in combined, but completion)
    │   │   │   │   │   │   │   │   └── overdue/
    │   │   │   │   │   │   │   │       └── page.tsx  # Overdue invoices
    │   │   │   │   │   │   │   │           # - Payment reminders
    │   │   │   │   │   │   │   │           # - Late fees
    │   │   │   │   │   │   │   │           # BE: financial-be/invoice
    │   │   │   │   │   │   │   │           # GET /v1/invoices?status=overdue
    │   │   │   │   │   │   │   ├── payment-methods/
    │   │   │   │   │   │   │   │   ├── [methodId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit payment method
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payment_method
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/financial/payment-methods/{method_id}
    │   │   │   │   │   │   │   │   │   ├── remove/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Remove payment method
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payment_method
    │   │   │   │   │   │   │   │   │   │       # DELETE /v1/financial/payment-methods/{method_id}
    │   │   │   │   │   │   │   │   │   └── verify/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Verify payment method
    │   │   │   │   │   │   │   │   │           # - Micro-deposit verification
    │   │   │   │   │   │   │   │   │           # - Card verification
    │   │   │   │   │   │   │   │   │           # BE: financial-be/payment_method
    │   │   │   │   │   │   │   │   │           # POST /v1/financial/payment-methods/{method_id}/verify
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add payment method
    │   │   │   │   │   │   │   │   │       # - Credit/debit card
    │   │   │   │   │   │   │   │   │       # - Bank account (ACH)
    │   │   │   │   │   │   │   │   │       # - Digital wallet
    │   │   │   │   │   │   │   │   │       # BE: financial-be/payment_method
    │   │   │   │   │   │   │   │   │       # POST /v1/financial/payment-methods
    │   │   │   │   │   │   │   │   └── page.tsx  # Payment methods list
    │   │   │   │   │   │   │   │       # BE: financial-be/payment_method
    │   │   │   │   │   │   │   │       # GET /v1/financial/payment-methods
    │   │   │   │   │   │   │   ├── payout-methods/
    │   │   │   │   │   │   │   │   ├── [methodId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit payout method
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payout_method
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/financial/payout-methods/{method_id}
    │   │   │   │   │   │   │   │   │   ├── remove/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Remove payout method
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payout_method
    │   │   │   │   │   │   │   │   │   │       # DELETE /v1/financial/payout-methods/{method_id}
    │   │   │   │   │   │   │   │   │   └── verify/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Verify payout method
    │   │   │   │   │   │   │   │   │           # BE: financial-be/payout_method
    │   │   │   │   │   │   │   │   │           # POST /v1/financial/payout-methods/{method_id}/verify
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add payout method
    │   │   │   │   │   │   │   │   │       # - Bank account
    │   │   │   │   │   │   │   │   │       # - PayPal
    │   │   │   │   │   │   │   │   │       # - Wire transfer
    │   │   │   │   │   │   │   │   │       # BE: financial-be/payout_method
    │   │   │   │   │   │   │   │   │       # POST /v1/financial/payout-methods
    │   │   │   │   │   │   │   │   └── page.tsx  # Payout methods list
    │   │   │   │   │   │   │   │       # BE: financial-be/payout_method
    │   │   │   │   │   │   │   │       # GET /v1/financial/payout-methods
    │   │   │   │   │   │   │   ├── payouts/
    │   │   │   │   │   │   │   │   ├── [payoutId]/
    │   │   │   │   │   │   │   │   │   ├── details/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Payout transaction detail
    │   │   │   │   │   │   │   │   │   │       # - Transaction breakdown
    │   │   │   │   │   │   │   │   │   │       # - Tax withholdings
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payout
    │   │   │   │   │   │   │   │   │   │       # GET /v1/payouts/{payout_id}
    │   │   │   │   │   │   │   │   │   └── receipt/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Payout receipt
    │   │   │   │   │   │   │   │   │           # - Download receipt
    │   │   │   │   │   │   │   │   │           # - Tax information
    │   │   │   │   │   │   │   │   │           # BE: financial-be/payout
    │   │   │   │   │   │   │   │   │           # GET /v1/payouts/{payout_id}/receipt
    │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending payouts
    │   │   │   │   │   │   │   │   │       # - In-process withdrawals
    │   │   │   │   │   │   │   │   │       # - Estimated dates
    │   │   │   │   │   │   │   │   │       # BE: financial-be/payout
    │   │   │   │   │   │   │   │   │       # GET /v1/payouts?status=pending
    │   │   │   │   │   │   │   │   └── schedule/
    │   │   │   │   │   │   │   │       └── page.tsx  # Schedule payout
    │   │   │   │   │   │   │   │           # - Set payout frequency
    │   │   │   │   │   │   │   │           # - Minimum threshold
    │   │   │   │   │   │   │   │           # BE: financial-be/payout
    │   │   │   │   │   │   │   │           # POST /v1/payouts/schedule
    │   │   │   │   │   │   │   ├── reconciliation/
    │   │   │   │   │   │   │   │   └── page.tsx  # Financial reconciliation
    │   │   │   │   │   │   │   │       # - Match transactions
    │   │   │   │   │   │   │   │       # - Resolve discrepancies
    │   │   │   │   │   │   │   │       # - Bank statement upload
    │   │   │   │   │   │   │   │       # BE: financial-be/reconciliation
    │   │   │   │   │   │   │   │       # GET /v1/financial/reconciliation
    │   │   │   │   │   │   │   │       # POST /v1/financial/reconciliation/upload
    │   │   │   │   │   │   │   ├── refunds/
    │   │   │   │   │   │   │   │   ├── [refundId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Refund detail
    │   │   │   │   │   │   │   │   │       # - Refund status
    │   │   │   │   │   │   │   │   │       # - Processing timeline
    │   │   │   │   │   │   │   │   │       # BE: financial-be/refund
    │   │   │   │   │   │   │   │   │       # GET /v1/refunds/{refund_id}
    │   │   │   │   │   │   │   │   ├── request/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Request refund
    │   │   │   │   │   │   │   │   │       # - Select transaction
    │   │   │   │   │   │   │   │   │       # - Reason and evidence
    │   │   │   │   │   │   │   │   │       # BE: financial-be/refund
    │   │   │   │   │   │   │   │   │       # POST /v1/refunds
    │   │   │   │   │   │   │   │   └── page.tsx  # Refunds list
    │   │   │   │   │   │   │   │       # - Requested refunds
    │   │   │   │   │   │   │   │       # - Completed refunds
    │   │   │   │   │   │   │   │       # BE: financial-be/refund
    │   │   │   │   │   │   │   │       # GET /v1/refunds
    │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   ├── earnings/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Earnings reports
    │   │   │   │   │   │   │   │   │       # - Period selection
    │   │   │   │   │   │   │   │   │       # - Breakdown by project
    │   │   │   │   │   │   │   │   │       # - Export options
    │   │   │   │   │   │   │   │   │       # BE: financial-be/reports
    │   │   │   │   │   │   │   │   │       # GET /v1/reports/earnings
    │   │   │   │   │   │   │   │   ├── expenses/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Expense reports
    │   │   │   │   │   │   │   │   │       # - Platform fees
    │   │   │   │   │   │   │   │   │       # - Service charges
    │   │   │   │   │   │   │   │   │       # BE: financial-be/reports
    │   │   │   │   │   │   │   │   │       # GET /v1/reports/expenses
    │   │   │   │   │   │   │   │   ├── tax/
    │   │   │   │   │   │   │   │   │   ├── 1099/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # 1099 tax forms
    │   │   │   │   │   │   │   │   │   │       # - Annual 1099s
    │   │   │   │   │   │   │   │   │   │       # - Download PDFs
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/tax
    │   │   │   │   │   │   │   │   │   │       # GET /v1/tax/forms/1099
    │   │   │   │   │   │   │   │   │   ├── vat/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # VAT reports
    │   │   │   │   │   │   │   │   │   │       # - VAT breakdown
    │   │   │   │   │   │   │   │   │   │       # - Export for filing
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/tax
    │   │   │   │   │   │   │   │   │   │       # GET /v1/tax/reports/vat
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Tax reports overview
    │   │   │   │   │   │   │   │   │       # BE: financial-be/tax
    │   │   │   │   │   │   │   │   │       # GET /v1/tax/reports
    │   │   │   │   │   │   │   │   └── page.tsx  # Financial reports overview
    │   │   │   │   │   │   │   │       # - Quick stats
    │   │   │   │   │   │   │   │       # - Report categories
    │   │   │   │   │   │   │   │       # BE: financial-be/reports
    │   │   │   │   │   │   │   │       # GET /v1/reports
    │   │   │   │   │   │   │   ├── tax/
    │   │   │   │   │   │   │   │   ├── forms/
    │   │   │   │   │   │   │   │   │   ├── w9/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # W-9 form management
    │   │   │   │   │   │   │   │   │   │       # - Submit W-9
    │   │   │   │   │   │   │   │   │   │       # - Update information
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/tax
    │   │   │   │   │   │   │   │   │   │       # POST /v1/tax/forms/w9
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Tax forms list
    │   │   │   │   │   │   │   │   │       # BE: financial-be/tax
    │   │   │   │   │   │   │   │   │       # GET /v1/tax/forms
    │   │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Tax settings
    │   │   │   │   │   │   │   │   │       # - Tax residency
    │   │   │   │   │   │   │   │   │       # - Withholding preferences
    │   │   │   │   │   │   │   │   │       # BE: financial-be/tax
    │   │   │   │   │   │   │   │   │       # PUT /v1/tax/settings
    │   │   │   │   │   │   │   │   └── page.tsx  # Tax overview
    │   │   │   │   │   │   │   │       # - Tax liability
    │   │   │   │   │   │   │   │       # - Forms required
    │   │   │   │   │   │   │   │       # BE: financial-be/tax
    │   │   │   │   │   │   │   │       # GET /v1/tax/overview
    │   │   │   │   │   │   │   └── wallet/
    │   │   │   │   │   │   │       ├── add-funds/
    │   │   │   │   │   │   │       │   └── page.tsx  # Add funds to wallet
    │   │   │   │   │   │   │       │       # - Amount input
    │   │   │   │   │   │   │       │       # - Payment method selection
    │   │   │   │   │   │   │       │       # BE: financial-be/wallet
    │   │   │   │   │   │   │       │       # POST /v1/wallet/add-funds
    │   │   │   │   │   │   │       └── withdraw/
    │   │   │   │   │   │   │           └── page.tsx  # Withdraw from wallet
    │   │   │   │   │   │   │               # - Withdrawal amount
    │   │   │   │   │   │   │               # - Destination account
    │   │   │   │   │   │   │               # BE: financial-be/wallet, financial-be/payout
    │   │   │   │   │   │   │               # POST /v1/wallet/withdraw
    │   │   │   │   │   │   ├── financials/  # Financial management
    │   │   │   │   │   │   │   ├── escrow/
    │   │   │   │   │   │   │   │   ├── [escrowId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Escrow detail
    │   │   │   │   │   │   │   │   │       # - Related contract
    │   │   │   │   │   │   │   │   │       # - Amount held
    │   │   │   │   │   │   │   │   │       # - Release schedule
    │   │   │   │   │   │   │   │   │       # - Transaction history
    │   │   │   │   │   │   │   │   │       # BE: financial-be/escrow
    │   │   │   │   │   │   │   │   │       # GET /v1/escrow/{escrow_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Escrow overview
    │   │   │   │   │   │   │   │       # - Active escrow accounts
    │   │   │   │   │   │   │   │       # - Total amount in escrow
    │   │   │   │   │   │   │   │       # - Pending releases
    │   │   │   │   │   │   │   │       # BE: financial-be/escrow
    │   │   │   │   │   │   │   │       # GET /v1/escrow
    │   │   │   │   │   │   │   ├── invoices/
    │   │   │   │   │   │   │   │   ├── [invoiceId]/
    │   │   │   │   │   │   │   │   │   ├── pay/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Pay invoice (client)
    │   │   │   │   │   │   │   │   │   │       # - Invoice summary
    │   │   │   │   │   │   │   │   │   │       # - Payment method selection
    │   │   │   │   │   │   │   │   │   │       # - Process payment
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payment
    │   │   │   │   │   │   │   │   │   │       # POST /v1/invoices/{invoice_id}/pay
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Invoice detail
    │   │   │   │   │   │   │   │   │       # - Invoice information
    │   │   │   │   │   │   │   │   │       # - Line items
    │   │   │   │   │   │   │   │   │       # - Tax details
    │   │   │   │   │   │   │   │   │       # - Payment status
    │   │   │   │   │   │   │   │   │       # - Download PDF
    │   │   │   │   │   │   │   │   │       # BE: financial-be/invoice
    │   │   │   │   │   │   │   │   │       # GET /v1/invoices/{invoice_id}
    │   │   │   │   │   │   │   │   │       # GET /v1/invoices/{invoice_id}/pdf
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create invoice (manual invoicing)
    │   │   │   │   │   │   │   │   │       # - Client selection
    │   │   │   │   │   │   │   │   │       # - Line items
    │   │   │   │   │   │   │   │   │       # - Tax settings
    │   │   │   │   │   │   │   │   │       # - Due date
    │   │   │   │   │   │   │   │   │       # - Notes
    │   │   │   │   │   │   │   │   │       # BE: financial-be/invoice
    │   │   │   │   │   │   │   │   │       # POST /v1/invoices
    │   │   │   │   │   │   │   │   └── page.tsx  # Invoices list
    │   │   │   │   │   │   │   │       # - Sent invoices (freelancer)
    │   │   │   │   │   │   │   │       # - Received invoices (client)
    │   │   │   │   │   │   │   │       # - Filter by status (paid, pending, overdue)
    │   │   │   │   │   │   │   │       # BE: financial-be/invoice
    │   │   │   │   │   │   │   │       # GET /v1/invoices
    │   │   │   │   │   │   │   ├── payment-methods/
    │   │   │   │   │   │   │   │   ├── [methodId]/
    │   │   │   │   │   │   │   │   │   ├── delete/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Delete payment method
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/payment_method
    │   │   │   │   │   │   │   │   │   │       # DELETE /v1/payment-methods/{method_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Payment method detail
    │   │   │   │   │   │   │   │   │       # BE: financial-be/payment_method
    │   │   │   │   │   │   │   │   │       # GET /v1/payment-methods/{method_id}
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add payment method
    │   │   │   │   │   │   │   │   │       # - Card details (Stripe Elements)
    │   │   │   │   │   │   │   │   │       # - PayPal connection
    │   │   │   │   │   │   │   │   │       # - Bank account (ACH)
    │   │   │   │   │   │   │   │   │       # - Set as default
    │   │   │   │   │   │   │   │   │       # BE: financial-be/payment_method
    │   │   │   │   │   │   │   │   │       # POST /v1/payment-methods
    │   │   │   │   │   │   │   │   └── page.tsx  # Payment methods list
    │   │   │   │   │   │   │   │       # - Saved credit cards
    │   │   │   │   │   │   │   │       # - PayPal accounts
    │   │   │   │   │   │   │   │       # - Bank accounts
    │   │   │   │   │   │   │   │       # - Default payment method
    │   │   │   │   │   │   │   │       # BE: financial-be/payment_method
    │   │   │   │   │   │   │   │       # GET /v1/payment-methods
    │   │   │   │   │   │   │   ├── payout-methods/
    │   │   │   │   │   │   │   │   ├── [methodId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Payout method detail
    │   │   │   │   │   │   │   │   │       # BE: financial-be/payout_method
    │   │   │   │   │   │   │   │   │       # GET /v1/payout-methods/{method_id}
    │   │   │   │   │   │   │   │   │       # DELETE /v1/payout-methods/{method_id}
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add payout method
    │   │   │   │   │   │   │   │   │       # - Bank account details
    │   │   │   │   │   │   │   │   │       # - PayPal email
    │   │   │   │   │   │   │   │   │       # - Wire transfer info
    │   │   │   │   │   │   │   │   │       # - Tax forms (W-9, W-8BEN)
    │   │   │   │   │   │   │   │   │       # BE: financial-be/payout_method
    │   │   │   │   │   │   │   │   │       # POST /v1/payout-methods
    │   │   │   │   │   │   │   │   └── page.tsx  # Payout methods list (freelancer)
    │   │   │   │   │   │   │   │       # - Bank accounts
    │   │   │   │   │   │   │   │       # - PayPal
    │   │   │   │   │   │   │   │       # - Wire transfer details
    │   │   │   │   │   │   │   │       # BE: financial-be/payout_method
    │   │   │   │   │   │   │   │       # GET /v1/payout-methods
    │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   ├── earnings/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Detailed earnings report
    │   │   │   │   │   │   │   │   │       # - By project
    │   │   │   │   │   │   │   │   │       # - By client
    │   │   │   │   │   │   │   │   │       # - By time period
    │   │   │   │   │   │   │   │   │       # BE: financial-be/reports
    │   │   │   │   │   │   │   │   │       # GET /v1/reports/earnings/detailed
    │   │   │   │   │   │   │   │   ├── spending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Detailed spending report
    │   │   │   │   │   │   │   │   │       # - By project
    │   │   │   │   │   │   │   │   │       # - By freelancer
    │   │   │   │   │   │   │   │   │       # - By category
    │   │   │   │   │   │   │   │   │       # BE: financial-be/reports
    │   │   │   │   │   │   │   │   │       # GET /v1/reports/spending/detailed
    │   │   │   │   │   │   │   │   └── page.tsx  # Financial reports
    │   │   │   │   │   │   │   │       # - Earnings report (freelancer)
    │   │   │   │   │   │   │   │       # - Spending report (client)
    │   │   │   │   │   │   │   │       # - Tax report
    │   │   │   │   │   │   │   │       # - Date range selection
    │   │   │   │   │   │   │   │       # - Export options
    │   │   │   │   │   │   │   │       # BE: financial-be/reports
    │   │   │   │   │   │   │   │       # GET /v1/reports/earnings
    │   │   │   │   │   │   │   │       # GET /v1/reports/spending
    │   │   │   │   │   │   │   ├── tax/
    │   │   │   │   │   │   │   │   ├── forms/
    │   │   │   │   │   │   │   │   │   ├── upload/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Upload tax form
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/tax
    │   │   │   │   │   │   │   │   │   │       # POST /v1/tax/forms
    │   │   │   │   │   │   │   │   │   │       # BE: storage-be/uploads
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Tax forms list
    │   │   │   │   │   │   │   │   │       # - W-9, 1099, W-8BEN, etc.
    │   │   │   │   │   │   │   │   │       # - Download forms
    │   │   │   │   │   │   │   │   │       # BE: financial-be/tax
    │   │   │   │   │   │   │   │   │       # GET /v1/tax/forms
    │   │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Tax settings
    │   │   │   │   │   │   │   │   │       # - Tax information
    │   │   │   │   │   │   │   │   │       # - VAT reverse charge
    │   │   │   │   │   │   │   │   │       # - Tax exemptions
    │   │   │   │   │   │   │   │   │       # BE: financial-be/tax
    │   │   │   │   │   │   │   │   │       # PUT /v1/tax/settings
    │   │   │   │   │   │   │   │   └── page.tsx  # Tax information
    │   │   │   │   │   │   │   │       # - Tax forms
    │   │   │   │   │   │   │   │       # - Tax ID
    │   │   │   │   │   │   │   │       # - VAT/GST number
    │   │   │   │   │   │   │   │       # - Tax residency
    │   │   │   │   │   │   │   │       # BE: financial-be/tax
    │   │   │   │   │   │   │   │       # GET /v1/tax/info
    │   │   │   │   │   │   │   ├── transactions/
    │   │   │   │   │   │   │   │   ├── [transactionId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Transaction detail
    │   │   │   │   │   │   │   │   │       # - Full transaction info
    │   │   │   │   │   │   │   │   │       # - Related contract/job
    │   │   │   │   │   │   │   │   │       # - Receipt download
    │   │   │   │   │   │   │   │   │       # BE: financial-be/transaction
    │   │   │   │   │   │   │   │   │       # GET /v1/transactions/{transaction_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Transaction history
    │   │   │   │   │   │   │   │       # - All transactions
    │   │   │   │   │   │   │   │       # - Filter by type (payment, payout, refund, etc.)
    │   │   │   │   │   │   │   │       # - Filter by date range
    │   │   │   │   │   │   │   │       # - Search by description
    │   │   │   │   │   │   │   │       # - Export to CSV
    │   │   │   │   │   │   │   │       # BE: financial-be/transaction
    │   │   │   │   │   │   │   │       # GET /v1/transactions?filters={...}
    │   │   │   │   │   │   │   ├── wallet/
    │   │   │   │   │   │   │   │   ├── add-funds/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add funds (client)
    │   │   │   │   │   │   │   │   │       # - Amount input
    │   │   │   │   │   │   │   │   │       # - Payment method selection
    │   │   │   │   │   │   │   │   │       # - Payment processing
    │   │   │   │   │   │   │   │   │       # BE: financial-be/wallet
    │   │   │   │   │   │   │   │   │       # POST /v1/wallet/add-funds
    │   │   │   │   │   │   │   │   │       # BE: financial-be/payment
    │   │   │   │   │   │   │   │   │       # POST /v1/payments (Stripe/PayPal integration)
    │   │   │   │   │   │   │   │   ├── withdraw/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Withdraw funds (freelancer)
    │   │   │   │   │   │   │   │   │       # - Amount input
    │   │   │   │   │   │   │   │   │       # - Payout method selection
    │   │   │   │   │   │   │   │   │       # - Tax information
    │   │   │   │   │   │   │   │   │       # - Withdrawal fees
    │   │   │   │   │   │   │   │   │       # BE: financial-be/payout
    │   │   │   │   │   │   │   │   │       # POST /v1/payouts/request
    │   │   │   │   │   │   │   │   │       # Publishes: PayoutRequested event
    │   │   │   │   │   │   │   │   └── page.tsx  # Wallet details
    │   │   │   │   │   │   │   │       # - Available balance
    │   │   │   │   │   │   │   │       # - Pending balance
    │   │   │   │   │   │   │   │       # - Escrow balance
    │   │   │   │   │   │   │   │       # - Add funds button (client)
    │   │   │   │   │   │   │   │       # - Withdraw button (freelancer)
    │   │   │   │   │   │   │   │       # - Transaction history
    │   │   │   │   │   │   │   │       # BE: financial-be/wallet
    │   │   │   │   │   │   │   │       # GET /v1/wallet
    │   │   │   │   │   │   │   └── page.tsx  # Financial overview
    │   │   │   │   │   │   │       # - Wallet balance
    │   │   │   │   │   │   │       # - Pending payments
    │   │   │   │   │   │   │       # - Recent transactions
    │   │   │   │   │   │   │       # - Earnings chart (freelancer)
    │   │   │   │   │   │   │       # - Spending chart (client)
    │   │   │   │   │   │   │       # BE: financial-be/wallet
    │   │   │   │   │   │   │       # GET /v1/wallet/balance
    │   │   │   │   │   │   │       # BE: financial-be/transaction
    │   │   │   │   │   │   │       # GET /v1/transactions/recent
    │   │   │   │   │   │   ├── invitations/
    │   │   │   │   │   │   │   ├── received/
    │   │   │   │   │   │   │   │   ├── [inviteId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Invitation details
    │   │   │   │   │   │   │   │   │       # - Job details
    │   │   │   │   │   │   │   │   │       # - Accept/decline
    │   │   │   │   │   │   │   │   │       # - Proposal draft
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/invite, jobs-be/job
    │   │   │   │   │   │   │   │   │       # GET /v1/invites/{invite_id}
    │   │   │   │   │   │   │   │   │       # POST /v1/invites/{invite_id}/accept
    │   │   │   │   │   │   │   │   │       # POST /v1/invites/{invite_id}/decline
    │   │   │   │   │   │   │   │   └── page.tsx  # Received invitations list
    │   │   │   │   │   │   │   │       # BE: proposals-be/invite
    │   │   │   │   │   │   │   │       # GET /v1/invites/received
    │   │   │   │   │   │   │   ├── sent/
    │   │   │   │   │   │   │   │   ├── [inviteId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Sent invitation tracking
    │   │   │   │   │   │   │   │   │       # - Delivery status
    │   │   │   │   │   │   │   │   │       # - Response tracking
    │   │   │   │   │   │   │   │   │       # BE: jobs-be/invitation
    │   │   │   │   │   │   │   │   │       # GET /v1/jobs/{job_id}/invitations/{invite_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Sent invitations list (client)
    │   │   │   │   │   │   │   │       # BE: jobs-be/invitation
    │   │   │   │   │   │   │   │       # GET /v1/jobs/{job_id}/invitations
    │   │   │   │   │   │   │   └── page.tsx  # Invitations overview
    │   │   │   │   │   │   │       # - Pending actions
    │   │   │   │   │   │   │       # - Response rate (client)
    │   │   │   │   │   │   │       # - Conversion metrics
    │   │   │   │   │   │   │       # BE: proposals-be/invite OR jobs-be/invitation (based on role)
    │   │   │   │   │   │   ├── job-alerts/
    │   │   │   │   │   │   │   ├── [alertId]/
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit job alert
    │   │   │   │   │   │   │   │   │       # BE: search-be/alert (if exists) or search-be/saved-search
    │   │   │   │   │   │   │   │   │       # PUT /v1/job-alerts/{alert_id}
    │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Alert history
    │   │   │   │   │   │   │   │   │       # - Jobs matched
    │   │   │   │   │   │   │   │   │       # - Notifications sent
    │   │   │   │   │   │   │   │   │       # BE: search-be/alert
    │   │   │   │   │   │   │   │   │       # GET /v1/job-alerts/{alert_id}/history
    │   │   │   │   │   │   │   │   └── page.tsx  # Alert detail
    │   │   │   │   │   │   │   │       # BE: search-be/alert
    │   │   │   │   │   │   │   │       # GET /v1/job-alerts/{alert_id}
    │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   └── page.tsx  # Create job alert
    │   │   │   │   │   │   │   │       # - Search criteria
    │   │   │   │   │   │   │   │       # - Notification frequency
    │   │   │   │   │   │   │   │       # - Delivery channel
    │   │   │   │   │   │   │   │       # BE: search-be/alert
    │   │   │   │   │   │   │   │       # POST /v1/job-alerts
    │   │   │   │   │   │   │   └── page.tsx  # Job alerts list
    │   │   │   │   │   │   │       # - Active alerts
    │   │   │   │   │   │   │       # - Pause/resume
    │   │   │   │   │   │   │       # BE: search-be/alert
    │   │   │   │   │   │   │       # GET /v1/job-alerts
    │   │   │   │   │   │   ├── jobs/  # Jobs management
    │   │   │   │   │   │   │   ├── [jobId]/
    │   │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Job analytics (client)
    │   │   │   │   │   │   │   │   │       # - Views
    │   │   │   │   │   │   │   │   │       # - Proposals received
    │   │   │   │   │   │   │   │   │       # - Proposal conversion rate
    │   │   │   │   │   │   │   │   │       # - Time to hire
    │   │   │   │   │   │   │   │   │       # BE: jobs-be/analytics
    │   │   │   │   │   │   │   │   │       # GET /v1/jobs/{job_id}/analytics
    │   │   │   │   │   │   │   │   │       # Job analytics detail
    │   │   │   │   │   │   │   │   │       # - View count trends
    │   │   │   │   │   │   │   │   │       # - Proposal funnel
    │   │   │   │   │   │   │   │   │       # - Demographic insights
    │   │   │   │   │   │   │   │   │       # - Conversion metrics
    │   │   │   │   │   │   │   │   ├── bidding/
    │   │   │   │   │   │   │   │   │   ├── place-bid/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Place/update bid (freelancer)
    │   │   │   │   │   │   │   │   │   │       # - Current bid amount
    │   │   │   │   │   │   │   │   │   │       # - Minimum bid
    │   │   │   │   │   │   │   │   │   │       # - Place new bid
    │   │   │   │   │   │   │   │   │   │       # - Bid increment rules
    │   │   │   │   │   │   │   │   │   │       # - Outbid warning
    │   │   │   │   │   │   │   │   │   │       # BE: proposals-be/bidding
    │   │   │   │   │   │   │   │   │   │       # POST /v1/jobs/{job_id}/bids
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/bids/{bid_id}
    │   │   │   │   │   │   │   │   │   │       # Publishes: BidPlaced, BidUpdated, OutbidAlert events
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Active bids on job (client view)
    │   │   │   │   │   │   │   │   │       # - Real-time bid updates
    │   │   │   │   │   │   │   │   │       # - Current lowest bid
    │   │   │   │   │   │   │   │   │       # - Bid history
    │   │   │   │   │   │   │   │   │       # - Accept bid
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/bidding
    │   │   │   │   │   │   │   │   │       # GET /v1/jobs/{job_id}/bids
    │   │   │   │   │   │   │   │   │       # WebSocket: ws://proposals-be/v1/jobs/{job_id}/bids
    │   │   │   │   │   │   │   │   ├── close/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Close job
    │   │   │   │   │   │   │   │   │       # - Reason for closing
    │   │   │   │   │   │   │   │   │       # - Notify applicants
    │   │   │   │   │   │   │   │   │       # BE: jobs-be/job
    │   │   │   │   │   │   │   │   │       # POST /v1/jobs/{job_id}/close
    │   │   │   │   │   │   │   │   │       # Publishes: JobClosed event
    │   │   │   │   │   │   │   │   ├── duplicate/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Duplicate job form
    │   │   │   │   │   │   │   │   │       # - Pre-filled from original
    │   │   │   │   │   │   │   │   │       # - Modify and post
    │   │   │   │   │   │   │   │   │       # BE: jobs-be/job
    │   │   │   │   │   │   │   │   │       # GET /v1/jobs/{job_id} (to fetch original)
    │   │   │   │   │   │   │   │   │       # POST /v1/jobs (to create duplicate)
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit job (client only)
    │   │   │   │   │   │   │   │   │       # - Same form as post job
    │   │   │   │   │   │   │   │   │       # - Cannot edit if has accepted proposals
    │   │   │   │   │   │   │   │   │       # BE: jobs-be/job
    │   │   │   │   │   │   │   │   │       # PUT /v1/jobs/{job_id}
    │   │   │   │   │   │   │   │   │       # Publishes: JobUpdated event
    │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Job edit history
    │   │   │   │   │   │   │   │   │       # - Version timeline
    │   │   │   │   │   │   │   │   │       # - Change diff viewer
    │   │   │   │   │   │   │   │   │       # - Revert capability
    │   │   │   │   │   │   │   │   │       # BE: jobs-be/audit
    │   │   │   │   │   │   │   │   │       # GET /v1/jobs/{job_id}/history
    │   │   │   │   │   │   │   │   ├── invite/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Invite freelancers (client)
    │   │   │   │   │   │   │   │   │       # - Search freelancers
    │   │   │   │   │   │   │   │   │       # - Send invitation with message
    │   │   │   │   │   │   │   │   │       # BE: jobs-be/invitations
    │   │   │   │   │   │   │   │   │       # POST /v1/jobs/{job_id}/invitations
    │   │   │   │   │   │   │   │   │       # BE: search-be
    │   │   │   │   │   │   │   │   │       # POST /v1/search/freelancers
    │   │   │   │   │   │   │   │   │       # BE: communications-be
    │   │   │   │   │   │   │   │   │       # Publishes: JobInvitationSent event
    │   │   │   │   │   │   │   │   ├── proposals/
    │   │   │   │   │   │   │   │   │   ├── [proposalId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Proposal detail
    │   │   │   │   │   │   │   │   │   │       # - Full proposal view
    │   │   │   │   │   │   │   │   │   │       # - Freelancer profile preview
    │   │   │   │   │   │   │   │   │   │       # - Accept/Reject buttons
    │   │   │   │   │   │   │   │   │   │       # - Shortlist button
    │   │   │   │   │   │   │   │   │   │       # - Message freelancer
    │   │   │   │   │   │   │   │   │   │       # BE: proposals-be
    │   │   │   │   │   │   │   │   │   │       # GET /v1/proposals/{proposal_id}
    │   │   │   │   │   │   │   │   │   │       # POST /v1/proposals/{proposal_id}/accept
    │   │   │   │   │   │   │   │   │   │       # POST /v1/proposals/{proposal_id}/reject
    │   │   │   │   │   │   │   │   │   │       # POST /v1/proposals/{proposal_id}/shortlist
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Proposals received (client)
    │   │   │   │   │   │   │   │   │       # - List all proposals
    │   │   │   │   │   │   │   │   │       # - Filter (all/shortlisted/archived)
    │   │   │   │   │   │   │   │   │       # - Sort (date, rate, rating)
    │   │   │   │   │   │   │   │   │       # - Proposal cards with key info
    │   │   │   │   │   │   │   │   │       # BE: proposals-be
    │   │   │   │   │   │   │   │   │       # GET /v1/proposals?job_id={job_id}
    │   │   │   │   │   │   │   │   ├── repost/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Repost expired job
    │   │   │   │   │   │   │   │   │       # - Review and update
    │   │   │   │   │   │   │   │   │       # - Set new deadline
    │   │   │   │   │   │   │   │   │       # BE: jobs-be/job
    │   │   │   │   │   │   │   │   │       # POST /v1/jobs/{job_id}/repost
    │   │   │   │   │   │   │   │   └── page.tsx  # Job detail page
    │   │   │   │   │   │   │   │       # - Full job description
    │   │   │   │   │   │   │   │       # - Client info
    │   │   │   │   │   │   │   │       # - Skills required
    │   │   │   │   │   │   │   │       # - Budget/rate
    │   │   │   │   │   │   │   │       # - Proposals count
    │   │   │   │   │   │   │   │       # - Similar jobs
    │   │   │   │   │   │   │   │       # - "Submit Proposal" button (freelancer)
    │   │   │   │   │   │   │   │       # - Save job button
    │   │   │   │   │   │   │   │       # BE: jobs-be/job
    │   │   │   │   │   │   │   │       # GET /v1/jobs/{job_id}
    │   │   │   │   │   │   │   │       # BE: proposals-be
    │   │   │   │   │   │   │   │       # GET /v1/proposals/count?job_id={job_id}
    │   │   │   │   │   │   │   │       # BE: search-be/similarity
    │   │   │   │   │   │   │   │       # GET /v1/similarity/jobs/{job_id}
    │   │   │   │   │   │   │   ├── archived/
    │   │   │   │   │   │   │   │   └── page.tsx  # Archived jobs list
    │   │   │   │   │   │   │   │       # - View archived jobs
    │   │   │   │   │   │   │   │       # - Restore capability
    │   │   │   │   │   │   │   │       # BE: jobs-be/job
    │   │   │   │   │   │   │   │       # GET /v1/jobs/archived
    │   │   │   │   │   │   │   ├── batch/
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Bulk job creation
    │   │   │   │   │   │   │   │   │       # - CSV upload
    │   │   │   │   │   │   │   │   │       # - Template selection
    │   │   │   │   │   │   │   │   │       # - Preview and publish
    │   │   │   │   │   │   │   │   │       # BE: jobs-be/job
    │   │   │   │   │   │   │   │   │       # POST /v1/jobs/batch
    │   │   │   │   │   │   │   │   └── manage/
    │   │   │   │   │   │   │   │       └── page.tsx  # Batch operations
    │   │   │   │   │   │   │   │           # - Bulk edit
    │   │   │   │   │   │   │   │           # - Bulk status change
    │   │   │   │   │   │   │   │           # - Bulk archive
    │   │   │   │   │   │   │   │           # BE: jobs-be/job
    │   │   │   │   │   │   │   │           # PATCH /v1/jobs/batch
    │   │   │   │   │   │   │   ├── browse/
    │   │   │   │   │   │   │   │   └── page.tsx  # Job listings with filters
    │   │   │   │   │   │   │   │       # - Category filters
    │   │   │   │   │   │   │   │       # - Budget range
    │   │   │   │   │   │   │   │       # - Experience level
    │   │   │   │   │   │   │   │       # - Job type (fixed/hourly)
    │   │   │   │   │   │   │   │       # - Location preferences
    │   │   │   │   │   │   │   │       # - Skills required
    │   │   │   │   │   │   │   │       # - Posted date
    │   │   │   │   │   │   │   │       # - Saved jobs indicator
    │   │   │   │   │   │   │   │       # - "Best Matches" tab
    │   │   │   │   │   │   │   │       # BE: jobs-be/job
    │   │   │   │   │   │   │   │       # GET /v1/jobs/browse?filters={...}
    │   │   │   │   │   │   │   │       # BE: search-be/query
    │   │   │   │   │   │   │   │       # POST /v1/search/jobs (for advanced search)
    │   │   │   │   │   │   │   ├── categories/
    │   │   │   │   │   │   │   │   ├── [categoryId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Jobs in category
    │   │   │   │   │   │   │   │   │       # BE: jobs-be/job
    │   │   │   │   │   │   │   │   │       # GET /v1/jobs?category_id={category_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Browse by category
    │   │   │   │   │   │   │   │       # - Category grid
    │   │   │   │   │   │   │   │       # - Subcategories
    │   │   │   │   │   │   │   │       # BE: jobs-be/categories
    │   │   │   │   │   │   │   │       # GET /v1/jobs/categories
    │   │   │   │   │   │   │   ├── drafts/
    │   │   │   │   │   │   │   │   ├── [draftId]/
    │   │   │   │   │   │   │   │   │   └── edit/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Edit draft
    │   │   │   │   │   │   │   │   │           # BE: jobs-be/draft
    │   │   │   │   │   │   │   │   │           # PUT /v1/jobs/drafts/{draft_id}
    │   │   │   │   │   │   │   │   │           # DELETE /v1/jobs/drafts/{draft_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Job drafts list
    │   │   │   │   │   │   │   │       # BE: jobs-be/draft
    │   │   │   │   │   │   │   │       # GET /v1/jobs/drafts
    │   │   │   │   │   │   │   ├── insights/
    │   │   │   │   │   │   │   │   └── page.tsx  # Market insights for job posting
    │   │   │   │   │   │   │   │       # - Recommended rates
    │   │   │   │   │   │   │   │       # - Competition analysis
    │   │   │   │   │   │   │   │       # - Time-to-hire estimates
    │   │   │   │   │   │   │   │       # - Skill demand trends
    │   │   │   │   │   │   │   │       # BE: jobs-be/analytics, search-be/trending
    │   │   │   │   │   │   │   │       # GET /v1/jobs/market-insights
    │   │   │   │   │   │   │   │       # GET /v1/trending/skills
    │   │   │   │   │   │   │   ├── invitations/
    │   │   │   │   │   │   │   │   └── page.tsx  # Job invitations received
    │   │   │   │   │   │   │   │       # BE: jobs-be/invitations
    │   │   │   │   │   │   │   │       # GET /v1/jobs/invitations
    │   │   │   │   │   │   │   ├── my-jobs/
    │   │   │   │   │   │   │   │   ├── active/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Active jobs only
    │   │   │   │   │   │   │   │   │       # BE: jobs-be/job
    │   │   │   │   │   │   │   │   │       # GET /v1/jobs/my-jobs?status=active
    │   │   │   │   │   │   │   │   ├── closed/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Closed jobs
    │   │   │   │   │   │   │   │   │       # BE: jobs-be/job
    │   │   │   │   │   │   │   │   │       # GET /v1/jobs/my-jobs?status=closed
    │   │   │   │   │   │   │   │   └── page.tsx  # All posted jobs
    │   │   │   │   │   │   │   │       # - Active jobs
    │   │   │   │   │   │   │   │       # - Closed jobs
    │   │   │   │   │   │   │   │       # - Drafts
    │   │   │   │   │   │   │   │       # BE: jobs-be/job
    │   │   │   │   │   │   │   │       # GET /v1/jobs/my-jobs?status=active
    │   │   │   │   │   │   │   ├── post/
    │   │   │   │   │   │   │   │   └── page.tsx  # Post a new job (client)
    │   │   │   │   │   │   │   │       # - Job title
    │   │   │   │   │   │   │   │       # - Job description (rich text editor)
    │   │   │   │   │   │   │   │       # - Category selection
    │   │   │   │   │   │   │   │       # - Required skills (autocomplete)
    │   │   │   │   │   │   │   │       # - Experience level
    │   │   │   │   │   │   │   │       # - Job type (fixed/hourly)
    │   │   │   │   │   │   │   │       # - Budget/rate
    │   │   │   │   │   │   │   │       # - Duration
    │   │   │   │   │   │   │   │       # - Attachments
    │   │   │   │   │   │   │   │       # - Screening questions
    │   │   │   │   │   │   │   │       # - Visibility (public/private/invited)
    │   │   │   │   │   │   │   │       # - Save as draft
    │   │   │   │   │   │   │   │       # BE: jobs-be/job
    │   │   │   │   │   │   │   │       # POST /v1/jobs
    │   │   │   │   │   │   │   │       # Body: { title, description, category_id, skills, budget, ... }
    │   │   │   │   │   │   │   │       # BE: jobs-be/attachments
    │   │   │   │   │   │   │   │       # POST /v1/jobs/{job_id}/attachments
    │   │   │   │   │   │   │   │       # BE: jobs-be/screening
    │   │   │   │   │   │   │   │       # POST /v1/jobs/{job_id}/screening-questions
    │   │   │   │   │   │   │   │       # BE: storage-be/uploads
    │   │   │   │   │   │   │   │       # Publishes: JobPosted event
    │   │   │   │   │   │   │   ├── recommendations/
    │   │   │   │   │   │   │   │   └── page.tsx  # Recommended jobs (freelancer)
    │   │   │   │   │   │   │   │       # - ML-powered job recommendations
    │   │   │   │   │   │   │   │       # - Based on skills, history, preferences
    │   │   │   │   │   │   │   │       # BE: search-be/recommendations
    │   │   │   │   │   │   │   │       # GET /v1/recommendations/jobs
    │   │   │   │   │   │   │   ├── saved/
    │   │   │   │   │   │   │   │   └── page.tsx  # Saved/bookmarked jobs
    │   │   │   │   │   │   │   │       # BE: jobs-be/saved_jobs
    │   │   │   │   │   │   │   │       # GET /v1/jobs/saved
    │   │   │   │   │   │   │   │       # DELETE /v1/jobs/saved/{job_id}
    │   │   │   │   │   │   │   ├── saved-searches/
    │   │   │   │   │   │   │   │   └── page.tsx  # Saved job searches
    │   │   │   │   │   │   │   │       # - Manage saved searches
    │   │   │   │   │   │   │   │       # - Email alerts toggle
    │   │   │   │   │   │   │   │       # BE: search-be/saved-search
    │   │   │   │   │   │   │   │       # GET /v1/search/saved-searches?type=jobs
    │   │   │   │   │   │   │   ├── scheduling/
    │   │   │   │   │   │   │   │   └── page.tsx  # Schedule job postings
    │   │   │   │   │   │   │   │       # - Future publish dates
    │   │   │   │   │   │   │   │       # - Auto-repost settings
    │   │   │   │   │   │   │   │       # - Expiration reminders
    │   │   │   │   │   │   │   │       # BE: jobs-be/job
    │   │   │   │   │   │   │   │       # POST /v1/jobs/{job_id}/schedule
    │   │   │   │   │   │   │   ├── templates/
    │   │   │   │   │   │   │   │   ├── [templateId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit job template
    │   │   │   │   │   │   │   │   │   │       # BE: jobs-be/template
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/jobs/templates/{template_id}
    │   │   │   │   │   │   │   │   │   └── use/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Use template to create job
    │   │   │   │   │   │   │   │   │           # BE: jobs-be/template, jobs-be/job
    │   │   │   │   │   │   │   │   │           # GET /v1/jobs/templates/{template_id}
    │   │   │   │   │   │   │   │   │           # POST /v1/jobs (using template data)
    │   │   │   │   │   │   │   │   └── page.tsx  # Job templates list
    │   │   │   │   │   │   │   │       # - Saved templates
    │   │   │   │   │   │   │   │       # - Create new template
    │   │   │   │   │   │   │   │       # BE: jobs-be/template
    │   │   │   │   │   │   │   │       # GET /v1/jobs/templates
    │   │   │   │   │   │   │   └── page.tsx  # Jobs list (role-based)
    │   │   │   │   │   │   │       # Freelancer view: Browse available jobs
    │   │   │   │   │   │   │       # Client view: My posted jobs
    │   │   │   │   │   │   │       # - Filters (category, budget, skills, etc.)
    │   │   │   │   │   │   │       # - Search bar
    │   │   │   │   │   │   │       # - Sort options
    │   │   │   │   │   │   │       # - Pagination
    │   │   │   │   │   │   │       # BE: jobs-be/job
    │   │   │   │   │   │   │       # GET /v1/jobs?filters=...&page=1&limit=20
    │   │   │   │   │   │   │       # Freelancer: GET /v1/jobs/browse
    │   │   │   │   │   │   │       # Client: GET /v1/jobs/my-jobs
    │   │   │   │   │   │   ├── learning/
    │   │   │   │   │   │   │   ├── achievements/
    │   │   │   │   │   │   │   │   └── page.tsx  # Achievements & badges
    │   │   │   │   │   │   │   │       # - Earned badges
    │   │   │   │   │   │   │   │       # - Progress to next level
    │   │   │   │   │   │   │   │       # - Leaderboard
    │   │   │   │   │   │   │   │       # BE: users-be/achievement
    │   │   │   │   │   │   │   │       # GET /v1/users/me/achievements
    │   │   │   │   │   │   │   │       # Learning achievements
    │   │   │   │   │   │   │   │       # - Certificates earned
    │   │   │   │   │   │   │   │       # - Badges
    │   │   │   │   │   │   │   │       # - Skill verification
    │   │   │   │   │   │   │   │       # - Skill verifications
    │   │   │   │   │   │   │   │       # BE: learning-be/achievement
    │   │   │   │   │   │   │   │       # GET /v1/learning/achievements
    │   │   │   │   │   │   │   ├── assessments/
    │   │   │   │   │   │   │   │   ├── [assessmentId]/
    │   │   │   │   │   │   │   │   │   ├── results/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Assessment results
    │   │   │   │   │   │   │   │   │   │       # BE: learning-be/assessment
    │   │   │   │   │   │   │   │   │   │       # GET /v1/learning/assessments/{assessment_id}/results
    │   │   │   │   │   │   │   │   │   ├── take/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Take assessment
    │   │   │   │   │   │   │   │   │   │       # BE: learning-be/assessment
    │   │   │   │   │   │   │   │   │   │       # POST /v1/learning/assessments/{assessment_id}/submit
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Assessment detail
    │   │   │   │   │   │   │   │   │       # BE: learning-be/assessment
    │   │   │   │   │   │   │   │   │       # GET /v1/learning/assessments/{assessment_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Assessments list
    │   │   │   │   │   │   │   │       # BE: learning-be/assessment
    │   │   │   │   │   │   │   │       # GET /v1/learning/assessments
    │   │   │   │   │   │   │   ├── certifications/
    │   │   │   │   │   │   │   │   ├── [certId]/
    │   │   │   │   │   │   │   │   │   ├── verify/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Verify certificate
    │   │   │   │   │   │   │   │   │   │       # BE: utility-be/learning
    │   │   │   │   │   │   │   │   │   │       # GET /v1/certifications/{cert_id}/verify
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Certificate detail
    │   │   │   │   │   │   │   │   │       # - Download PDF
    │   │   │   │   │   │   │   │   │       # - Share link
    │   │   │   │   │   │   │   │   │       # - Add to profile
    │   │   │   │   │   │   │   │   │       # BE: utility-be/learning
    │   │   │   │   │   │   │   │   │       # GET /v1/certifications/{cert_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Manage certifications
    │   │   │   │   │   │   │   │       # - Upload certificates
    │   │   │   │   │   │   │   │       # - Verification status
    │   │   │   │   │   │   │   │       # - Expiry tracking
    │   │   │   │   │   │   │   │       # BE: users-be/credential
    │   │   │   │   │   │   │   │       # GET /v1/users/me/credentials?type=certification
    │   │   │   │   │   │   │   │       # All certifications
    │   │   │   │   │   │   │   │       # - Earned certificates
    │   │   │   │   │   │   │   │       # - Available certifications
    │   │   │   │   │   │   │   │       # BE: utility-be/learning
    │   │   │   │   │   │   │   │       # GET /v1/users/me/certifications
    │   │   │   │   │   │   │   ├── courses/
    │   │   │   │   │   │   │   │   ├── [courseId]/
    │   │   │   │   │   │   │   │   │   ├── assessments/
    │   │   │   │   │   │   │   │   │   │   └── [assessmentId]/
    │   │   │   │   │   │   │   │   │   │       ├── results/
    │   │   │   │   │   │   │   │   │   │       │   └── page.tsx  # View results
    │   │   │   │   │   │   │   │   │   │       │       # BE: utility-be/learning
    │   │   │   │   │   │   │   │   │   │       │       # GET /v1/assessments/{assessment_id}/results
    │   │   │   │   │   │   │   │   │   │       ├── start/
    │   │   │   │   │   │   │   │   │   │       │   └── page.tsx  # Start assessment
    │   │   │   │   │   │   │   │   │   │       │       # BE: utility-be/learning
    │   │   │   │   │   │   │   │   │   │       │       # POST /v1/assessments/{assessment_id}/start
    │   │   │   │   │   │   │   │   │   │       └── submit/
    │   │   │   │   │   │   │   │   │   │           └── page.tsx  # Submit answers
    │   │   │   │   │   │   │   │   │   │               # BE: utility-be/learning
    │   │   │   │   │   │   │   │   │   │               # POST /v1/assessments/{assessment_id}/submit
    │   │   │   │   │   │   │   │   │   ├── lessons/
    │   │   │   │   │   │   │   │   │   │   └── [lessonId]/
    │   │   │   │   │   │   │   │   │   │       └── page.tsx  # Lesson content
    │   │   │   │   │   │   │   │   │   │           # BE: learning-be/lesson (if exists) or external LMS
    │   │   │   │   │   │   │   │   │   │           # GET /v1/learning/courses/{course_id}/lessons/{lesson_id}
    │   │   │   │   │   │   │   │   │   │           # Lesson view
    │   │   │   │   │   │   │   │   │   │           # - Video/content player
    │   │   │   │   │   │   │   │   │   │           # - Notes taking
    │   │   │   │   │   │   │   │   │   │           # - Progress tracking
    │   │   │   │   │   │   │   │   │   │           # BE: utility-be/learning OR external LMS
    │   │   │   │   │   │   │   │   │   │           # GET /v1/courses/{course_id}/lessons/{lesson_id}
    │   │   │   │   │   │   │   │   │   ├── progress/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Course progress
    │   │   │   │   │   │   │   │   │   │       # BE: learning-be/progress
    │   │   │   │   │   │   │   │   │   │       # GET /v1/learning/courses/{course_id}/progress
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Course detail
    │   │   │   │   │   │   │   │   │       # BE: learning-be/course
    │   │   │   │   │   │   │   │   │       # GET /v1/learning/courses/{course_id}
    │   │   │   │   │   │   │   │   │       # Course overview
    │   │   │   │   │   │   │   │   │       # - Syllabus
    │   │   │   │   │   │   │   │   │       # - Progress
    │   │   │   │   │   │   │   │   │       # - Enroll button
    │   │   │   │   │   │   │   │   │       # BE: utility-be/learning
    │   │   │   │   │   │   │   │   │       # GET /v1/courses/{course_id}
    │   │   │   │   │   │   │   │   │       # POST /v1/courses/{course_id}/enroll
    │   │   │   │   │   │   │   │   ├── browse/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Browse all courses
    │   │   │   │   │   │   │   │   │       # - Filter by skill
    │   │   │   │   │   │   │   │   │       # - Search
    │   │   │   │   │   │   │   │   │       # - Recommendations
    │   │   │   │   │   │   │   │   │       # BE: utility-be/learning
    │   │   │   │   │   │   │   │   │       # GET /v1/courses
    │   │   │   │   │   │   │   │   ├── my-courses/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Enrolled courses
    │   │   │   │   │   │   │   │   │       # - In progress
    │   │   │   │   │   │   │   │   │       # - Completed
    │   │   │   │   │   │   │   │   │       # - Certificates
    │   │   │   │   │   │   │   │   │       # BE: utility-be/learning
    │   │   │   │   │   │   │   │   │       # GET /v1/users/me/courses
    │   │   │   │   │   │   │   │   └── page.tsx  # Courses catalog
    │   │   │   │   │   │   │   │       # BE: learning-be/course
    │   │   │   │   │   │   │   │       # GET /v1/learning/courses
    │   │   │   │   │   │   │   ├── mentorship/
    │   │   │   │   │   │   │   │   ├── [sessionId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Mentorship session details
    │   │   │   │   │   │   │   │   │       # BE: users-be/mentorship
    │   │   │   │   │   │   │   │   │       # GET /v1/users/me/mentorship/{session_id}
    │   │   │   │   │   │   │   │   ├── find-mentor/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Find a mentor
    │   │   │   │   │   │   │   │   │       # BE: users-be/mentorship, search-be/query
    │   │   │   │   │   │   │   │   │       # POST /v1/search/mentors
    │   │   │   │   │   │   │   │   ├── my-mentees/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Manage mentees
    │   │   │   │   │   │   │   │   │       # BE: users-be/mentorship
    │   │   │   │   │   │   │   │   │       # GET /v1/users/me/mentorship/mentees
    │   │   │   │   │   │   │   │   └── page.tsx  # Mentorship dashboard
    │   │   │   │   │   │   │   │       # BE: users-be/mentorship
    │   │   │   │   │   │   │   │       # GET /v1/users/me/mentorship
    │   │   │   │   │   │   │   ├── paths/
    │   │   │   │   │   │   │   │   ├── [pathId]/
    │   │   │   │   │   │   │   │   │   ├── progress/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Learning path progress
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/learning_path
    │   │   │   │   │   │   │   │   │   │       # GET /v1/users/me/learning-path/{path_id}/progress
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Learning path details
    │   │   │   │   │   │   │   │   │       # - Courses
    │   │   │   │   │   │   │   │   │       # - Milestones
    │   │   │   │   │   │   │   │   │       # - Resources
    │   │   │   │   │   │   │   │   │       # BE: users-be/learning_path
    │   │   │   │   │   │   │   │   │       # GET /v1/users/me/learning-path/{path_id}
    │   │   │   │   │   │   │   │   ├── discover/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Discover learning paths
    │   │   │   │   │   │   │   │   │       # BE: users-be/learning_path
    │   │   │   │   │   │   │   │   │       # GET /v1/learning-paths/discover
    │   │   │   │   │   │   │   │   └── page.tsx  # My learning paths
    │   │   │   │   │   │   │   │       # BE: users-be/learning_path
    │   │   │   │   │   │   │   │       # GET /v1/users/me/learning-path
    │   │   │   │   │   │   │   └── skill-tests/
    │   │   │   │   │   │   │       ├── [testId]/
    │   │   │   │   │   │   │       │   ├── instructions/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Test instructions
    │   │   │   │   │   │   │       │   │       # BE: users-be/capabilities
    │   │   │   │   │   │   │       │   │       # GET /v1/skill-tests/{test_id}
    │   │   │   │   │   │   │       │   ├── results/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Test results
    │   │   │   │   │   │   │       │   │       # - Score
    │   │   │   │   │   │   │       │   │       # - Percentile
    │   │   │   │   │   │   │       │   │       # - Badge earned
    │   │   │   │   │   │   │       │   │       # BE: users-be/capabilities
    │   │   │   │   │   │   │       │   │       # GET /v1/skill-tests/{test_id}/results
    │   │   │   │   │   │   │       │   └── take-test/
    │   │   │   │   │   │   │       │       └── page.tsx  # Take the test
    │   │   │   │   │   │   │       │           # - Timed test interface
    │   │   │   │   │   │   │       │           # - Submit answers
    │   │   │   │   │   │   │       │           # BE: users-be/capabilities
    │   │   │   │   │   │   │       │           # POST /v1/skill-tests/{test_id}/submit
    │   │   │   │   │   │   │       └── page.tsx  # Available skill tests
    │   │   │   │   │   │   │           # - Browse by skill
    │   │   │   │   │   │   │           # - Test history
    │   │   │   │   │   │   │           # - Top scores
    │   │   │   │   │   │   │           # BE: users-be/capabilities
    │   │   │   │   │   │   │           # GET /v1/skill-tests
    │   │   │   │   │   │   ├── messages/  # Messaging
    │   │   │   │   │   │   │   ├── [conversationId]/
    │   │   │   │   │   │   │   │   ├── archive/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Archive conversation
    │   │   │   │   │   │   │   │   │       # BE: communications-be/conversations
    │   │   │   │   │   │   │   │   │       # POST /v1/conversations/{conversation_id}/archive
    │   │   │   │   │   │   │   │   ├── info/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Conversation info
    │   │   │   │   │   │   │   │   │       # - Participants
    │   │   │   │   │   │   │   │   │       # - Related job/contract
    │   │   │   │   │   │   │   │   │       # - Shared files
    │   │   │   │   │   │   │   │   │       # - Search in conversation
    │   │   │   │   │   │   │   │   │       # BE: communications-be/conversations
    │   │   │   │   │   │   │   │   │       # GET /v1/conversations/{conversation_id}/info
    │   │   │   │   │   │   │   │   └── page.tsx  # Conversation thread
    │   │   │   │   │   │   │   │       # - Message history
    │   │   │   │   │   │   │   │       # - Real-time new messages
    │   │   │   │   │   │   │   │       # - Message composer
    │   │   │   │   │   │   │   │       # - File attachments
    │   │   │   │   │   │   │   │       # - Typing indicators
    │   │   │   │   │   │   │   │       # - Read receipts
    │   │   │   │   │   │   │   │       # - Quick actions (block, report)
    │   │   │   │   │   │   │   │       # BE: communications-be/messages
    │   │   │   │   │   │   │   │       # GET /v1/conversations/{conversation_id}/messages
    │   │   │   │   │   │   │   │       # POST /v1/messages
    │   │   │   │   │   │   │   │       # WebSocket: Real-time message delivery
    │   │   │   │   │   │   │   │       # Publishes: MessageSent event
    │   │   │   │   │   │   │   │       # Conversation
    │   │   │   │   │   │   │   │       # - Thread, read receipts, uploads
    │   │   │   │   │   │   │   │       # BE: communications-be/conversation | message
    │   │   │   │   │   │   │   │       # GET  /v1/conversations/{id}
    │   │   │   │   │   │   │   │       # GET  /v1/conversations/{id}/messages
    │   │   │   │   │   │   │   ├── archived/
    │   │   │   │   │   │   │   │   └── page.tsx  # Archived conversations
    │   │   │   │   │   │   │   │       # BE: communications-be/conversations
    │   │   │   │   │   │   │   │       # GET /v1/conversations?archived=true
    │   │   │   │   │   │   │   │       # - View archived messages
    │   │   │   │   │   │   │   │       # - Unarchive capability
    │   │   │   │   │   │   │   │       # BE: communications-be/conversation
    │   │   │   │   │   │   │   ├── compose/
    │   │   │   │   │   │   │   │   └── page.tsx  # Compose new message
    │   │   │   │   │   │   │   │       # - Recipient search
    │   │   │   │   │   │   │   │       # - Subject and body
    │   │   │   │   │   │   │   │       # - Attachments
    │   │   │   │   │   │   │   │       # BE: communications-be/conversation
    │   │   │   │   │   │   │   │       # POST /v1/conversations
    │   │   │   │   │   │   │   ├── drafts/
    │   │   │   │   │   │   │   │   └── page.tsx  # Message drafts
    │   │   │   │   │   │   │   │       # - Saved drafts
    │   │   │   │   │   │   │   │       # - Resume editing
    │   │   │   │   │   │   │   │       # BE: communications-be/draft
    │   │   │   │   │   │   │   │       # GET /v1/messages/drafts
    │   │   │   │   │   │   │   ├── new/
    │   │   │   │   │   │   │   │   └── page.tsx  # Start new conversation
    │   │   │   │   │   │   │   │       # - Select recipient (search users)
    │   │   │   │   │   │   │   │       # - Initial message
    │   │   │   │   │   │   │   │       # BE: communications-be/conversations
    │   │   │   │   │   │   │   │       # POST /v1/conversations
    │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   └── notifications/
    │   │   │   │   │   │   │   │       └── page.tsx  # Notification preferences
    │   │   │   │   │   │   │   │           # - Email/SMS/push, digests
    │   │   │   │   │   │   │   │           # BE: communications-be/notification
    │   │   │   │   │   │   │   │           # GET/PUT /v1/notifications/preferences
    │   │   │   │   │   │   │   ├── starred/
    │   │   │   │   │   │   │   │   └── page.tsx  # Starred messages
    │   │   │   │   │   │   │   │       # - Important messages
    │   │   │   │   │   │   │   │       # - Quick access
    │   │   │   │   │   │   │   │       # BE: communications-be/conversation
    │   │   │   │   │   │   │   │       # GET /v1/conversations?starred=true
    │   │   │   │   │   │   │   └── page.tsx  # Inbox
    │   │   │   │   │   │   │       # - User conversations list
    │   │   │   │   │   │   │       # BE: communications-be/conversation
    │   │   │   │   │   │   │       # GET /v1/conversations?mine=1
    │   │   │   │   │   │   ├── network/
    │   │   │   │   │   │   │   ├── connections/
    │   │   │   │   │   │   │   │   ├── [userId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Connection profile view
    │   │   │   │   │   │   │   │   │       # BE: users-be/profile, users-be/connection
    │   │   │   │   │   │   │   │   │       # GET /v1/users/{user_id}
    │   │   │   │   │   │   │   │   │       # GET /v1/users/me/connections/{user_id}
    │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending connection requests
    │   │   │   │   │   │   │   │   │       # BE: users-be/connection
    │   │   │   │   │   │   │   │   │       # GET /v1/users/me/connections/pending
    │   │   │   │   │   │   │   │   └── page.tsx  # Connections list
    │   │   │   │   │   │   │   │       # BE: users-be/connection
    │   │   │   │   │   │   │   │       # GET /v1/users/me/connections
    │   │   │   │   │   │   │   ├── groups/
    │   │   │   │   │   │   │   │   ├── [groupId]/
    │   │   │   │   │   │   │   │   │   ├── members/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Group members
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/user_group
    │   │   │   │   │   │   │   │   │   │       # GET /v1/groups/{group_id}/members
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Group details
    │   │   │   │   │   │   │   │   │       # - Posts
    │   │   │   │   │   │   │   │   │       # - Events
    │   │   │   │   │   │   │   │   │       # - Resources
    │   │   │   │   │   │   │   │   │       # BE: users-be/user_group
    │   │   │   │   │   │   │   │   │       # GET /v1/groups/{group_id}
    │   │   │   │   │   │   │   │   ├── discover/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Discover groups
    │   │   │   │   │   │   │   │   │       # BE: users-be/user_group
    │   │   │   │   │   │   │   │   │       # GET /v1/groups/discover
    │   │   │   │   │   │   │   │   └── page.tsx  # My groups
    │   │   │   │   │   │   │   │       # BE: users-be/user_group
    │   │   │   │   │   │   │   │       # GET /v1/users/me/groups
    │   │   │   │   │   │   │   ├── recommendations/
    │   │   │   │   │   │   │   │   └── page.tsx  # Connection recommendations
    │   │   │   │   │   │   │   │       # - People you may know
    │   │   │   │   │   │   │   │       # - Similar professionals
    │   │   │   │   │   │   │   │       # BE: search-be/recommendation
    │   │   │   │   │   │   │   │       # GET /v1/recommendations/connections
    │   │   │   │   │   │   │   └── referrals/
    │   │   │   │   │   │   │       ├── dashboard/
    │   │   │   │   │   │   │       │   └── page.tsx  # Referral dashboard
    │   │   │   │   │   │   │       │       # - Total referrals
    │   │   │   │   │   │   │       │       # - Earnings
    │   │   │   │   │   │   │       │       # - Conversion rate
    │   │   │   │   │   │   │       │       # BE: users-be/referral
    │   │   │   │   │   │   │       │       # GET /v1/users/me/referral-code
    │   │   │   │   │   │   │       │       # GET /v1/referrals/analytics
    │   │   │   │   │   │   │       └── page.tsx  # Referrals overview
    │   │   │   │   │   │   │           # - Share referral code
    │   │   │   │   │   │   │           # - Track referrals
    │   │   │   │   │   │   │           # BE: users-be/referral
    │   │   │   │   │   │   │           # GET /v1/referrals
    │   │   │   │   │   │   ├── notifications/  # Notifications center
    │   │   │   │   │   │   │   ├── [notificationId]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Notification detail (redirects to relevant page)
    │   │   │   │   │   │   │   │       # - Mark as read
    │   │   │   │   │   │   │   │       # - Navigate to related entity
    │   │   │   │   │   │   │   │       # BE: communications-be/notifications
    │   │   │   │   │   │   │   │       # POST /v1/notifications/{notif_id}/read
    │   │   │   │   │   │   │   ├── all/
    │   │   │   │   │   │   │   │   └── page.tsx  # All notifications
    │   │   │   │   │   │   │   │       # - Complete history
    │   │   │   │   │   │   │   │       # - Filter options
    │   │   │   │   │   │   │   │       # BE: communications-be/notification
    │   │   │   │   │   │   │   │       # GET /v1/notifications
    │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   ├── channels/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Notification channels
    │   │   │   │   │   │   │   │   │       # - Email settings
    │   │   │   │   │   │   │   │   │       # - Push notification settings
    │   │   │   │   │   │   │   │   │       # - SMS settings
    │   │   │   │   │   │   │   │   │       # BE: communications-be/preferences
    │   │   │   │   │   │   │   │   │       # GET /v1/notifications/channels
    │   │   │   │   │   │   │   │   │       # PUT /v1/notifications/channels
    │   │   │   │   │   │   │   │   └── page.tsx  # Notification settings
    │   │   │   │   │   │   │   │       # - Email notifications
    │   │   │   │   │   │   │   │       # - Push notifications
    │   │   │   │   │   │   │   │       # - In-app notifications
    │   │   │   │   │   │   │   │       # - Notification preferences by type
    │   │   │   │   │   │   │   │       # - Frequency settings
    │   │   │   │   │   │   │   │       # - Quiet hours
    │   │   │   │   │   │   │   │       # BE: communications-be/preferences
    │   │   │   │   │   │   │   │       # GET /v1/notifications/preferences
    │   │   │   │   │   │   │   │       # PUT /v1/notifications/preferences
    │   │   │   │   │   │   │   │       # Notification preferences (completion)
    │   │   │   │   │   │   │   ├── unread/
    │   │   │   │   │   │   │   │   └── page.tsx  # Unread notifications only
    │   │   │   │   │   │   │   │       # BE: communications-be/notifications
    │   │   │   │   │   │   │   │       # GET /v1/notifications?unread=true
    │   │   │   │   │   │   │   │       # BE: communications-be/notification
    │   │   │   │   │   │   │   │       # GET /v1/notifications?read=false
    │   │   │   │   │   │   │   └── page.tsx  # Notifications center
    │   │   │   │   │   │   │       # - All notifications list
    │   │   │   │   │   │   │       # - Unread indicator
    │   │   │   │   │   │   │       # - Mark all as read
    │   │   │   │   │   │   │       # - Filter by type
    │   │   │   │   │   │   │       # - Real-time updates
    │   │   │   │   │   │   │       # BE: communications-be/notifications
    │   │   │   │   │   │   │       # GET /v1/notifications
    │   │   │   │   │   │   │       # POST /v1/notifications/read-all
    │   │   │   │   │   │   │       # WebSocket: ws://communications-be/v1/notifications
    │   │   │   │   │   │   ├── organization/  # Organization management (for clients)
    │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   └── page.tsx  # Organization analytics
    │   │   │   │   │   │   │   │       # - Hiring metrics
    │   │   │   │   │   │   │   │       # - Freelancer performance
    │   │   │   │   │   │   │   │       # - Cost per hire
    │   │   │   │   │   │   │   │       # - Time to hire
    │   │   │   │   │   │   │   │       # BE: analytics-be
    │   │   │   │   │   │   │   │       # GET /v1/analytics/organization/{org_id}
    │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Organization billing
    │   │   │   │   │   │   │   │   │       # - Billing profile
    │   │   │   │   │   │   │   │   │       # - Tax information
    │   │   │   │   │   │   │   │   │       # - Payment methods
    │   │   │   │   │   │   │   │   │       # BE: financial-be/billing_profile
    │   │   │   │   │   │   │   │   │       # GET /v1/organizations/{org_id}/billing-profile
    │   │   │   │   │   │   │   │   └── page.tsx  # Organization settings
    │   │   │   │   │   │   │   │       # - Company name
    │   │   │   │   │   │   │   │       # - Industry
    │   │   │   │   │   │   │   │       # - Company size
    │   │   │   │   │   │   │   │       # - Website
    │   │   │   │   │   │   │   │       # - Logo
    │   │   │   │   │   │   │   │       # BE: users-be/organization
    │   │   │   │   │   │   │   │       # PATCH /v1/organizations/{org_id}
    │   │   │   │   │   │   │   ├── spending/
    │   │   │   │   │   │   │   │   ├── budgets/
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create budget
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/budget
    │   │   │   │   │   │   │   │   │   │       # POST /v1/organizations/{org_id}/budgets
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Budget management
    │   │   │   │   │   │   │   │   │       # - Set budgets
    │   │   │   │   │   │   │   │   │       # - Budget alerts
    │   │   │   │   │   │   │   │   │       # BE: financial-be/budget
    │   │   │   │   │   │   │   │   │       # GET /v1/organizations/{org_id}/budgets
    │   │   │   │   │   │   │   │   └── page.tsx  # Spending overview
    │   │   │   │   │   │   │   │       # - Total spending
    │   │   │   │   │   │   │   │       # - By project
    │   │   │   │   │   │   │   │       # - By freelancer
    │   │   │   │   │   │   │   │       # - By time period
    │   │   │   │   │   │   │   │       # BE: financial-be/reports
    │   │   │   │   │   │   │   │       # GET /v1/organizations/{org_id}/spending
    │   │   │   │   │   │   │   ├── team/
    │   │   │   │   │   │   │   │   ├── [memberId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit member
    │   │   │   │   │   │   │   │   │   │       # - Change role
    │   │   │   │   │   │   │   │   │   │       # - Update permissions
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/team
    │   │   │   │   │   │   │   │   │   │       # PATCH /v1/organizations/{org_id}/members/{member_id}
    │   │   │   │   │   │   │   │   │   └── remove/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Remove member
    │   │   │   │   │   │   │   │   │           # BE: users-be/team
    │   │   │   │   │   │   │   │   │           # DELETE /v1/organizations/{org_id}/members/{member_id}
    │   │   │   │   │   │   │   │   ├── invite/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Invite team member
    │   │   │   │   │   │   │   │   │       # - Email address
    │   │   │   │   │   │   │   │   │       # - Role selection
    │   │   │   │   │   │   │   │   │       # - Permissions
    │   │   │   │   │   │   │   │   │       # BE: users-be/team
    │   │   │   │   │   │   │   │   │       # POST /v1/organizations/{org_id}/members/invite
    │   │   │   │   │   │   │   │   │       # BE: communications-be
    │   │   │   │   │   │   │   │   │       # Sends invitation email
    │   │   │   │   │   │   │   │   ├── roles/
    │   │   │   │   │   │   │   │   │   ├── [roleId]/
    │   │   │   │   │   │   │   │   │   │   └── edit/
    │   │   │   │   │   │   │   │   │   │       └── page.tsx  # Edit role
    │   │   │   │   │   │   │   │   │   │           # BE: users-be/role
    │   │   │   │   │   │   │   │   │   │           # PUT /v1/organizations/{org_id}/roles/{role_id}
    │   │   │   │   │   │   │   │   │   │           # DELETE /v1/organizations/{org_id}/roles/{role_id}
    │   │   │   │   │   │   │   │   │   └── create/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Create custom role
    │   │   │   │   │   │   │   │   │           # - Role name
    │   │   │   │   │   │   │   │   │           # - Permissions selection
    │   │   │   │   │   │   │   │   │           # BE: users-be/role
    │   │   │   │   │   │   │   │   │           # POST /v1/organizations/{org_id}/roles
    │   │   │   │   │   │   │   │   └── page.tsx  # Team members list
    │   │   │   │   │   │   │   │       # - Active members
    │   │   │   │   │   │   │   │       # - Pending invitations
    │   │   │   │   │   │   │   │       # - Roles
    │   │   │   │   │   │   │   │       # BE: users-be/team
    │   │   │   │   │   │   │   │       # GET /v1/organizations/{org_id}/members
    │   │   │   │   │   │   │   └── page.tsx  # Organization overview
    │   │   │   │   │   │   │       # - Company details
    │   │   │   │   │   │   │       # - Team members
    │   │   │   │   │   │   │       # - Spending overview
    │   │   │   │   │   │   │       # - Active contracts
    │   │   │   │   │   │   │       # BE: users-be/organization
    │   │   │   │   │   │   │       # GET /v1/organizations/{org_id}
    │   │   │   │   │   │   ├── portfolio/
    │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   └── page.tsx  # Portfolio analytics
    │   │   │   │   │   │   │   │       # - View count trends
    │   │   │   │   │   │   │   │       # - Engagement metrics
    │   │   │   │   │   │   │   │       # - Profile strength score
    │   │   │   │   │   │   │   │       # BE: users-be/analytics
    │   │   │   │   │   │   │   │       # GET /v1/users/me/portfolio/analytics
    │   │   │   │   │   │   │   ├── import/
    │   │   │   │   │   │   │   │   └── page.tsx  # Import portfolio
    │   │   │   │   │   │   │   │       # - LinkedIn import
    │   │   │   │   │   │   │   │       # - Behance import
    │   │   │   │   │   │   │   │       # - GitHub import
    │   │   │   │   │   │   │   │       # BE: users-be/portfolio, storage-be/asset
    │   │   │   │   │   │   │   │       # POST /v1/users/me/portfolio/import
    │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │       └── page.tsx  # Portfolio templates
    │   │   │   │   │   │   │           # - Pre-designed layouts
    │   │   │   │   │   │   │           # - Customize template
    │   │   │   │   │   │   │           # BE: users-be/portfolio
    │   │   │   │   │   │   │           # GET /v1/portfolio/templates
    │   │   │   │   │   │   ├── profile/  # Current user profile management
    │   │   │   │   │   │   │   ├── availability/
    │   │   │   │   │   │   │   │   └── page.tsx  # Availability management
    │   │   │   │   │   │   │   │       # - Calendar view
    │   │   │   │   │   │   │   │       # - Set available hours
    │   │   │   │   │   │   │   │       # - Time zone
    │   │   │   │   │   │   │   │       # - Vacation mode
    │   │   │   │   │   │   │   │       # - Max concurrent projects
    │   │   │   │   │   │   │   │       # BE: users-be/availability
    │   │   │   │   │   │   │   │       # GET /v1/users/{id}/availability
    │   │   │   │   │   │   │   │       # POST /v1/users/{id}/availability
    │   │   │   │   │   │   │   │       # PATCH /v1/users/{id}/availability
    │   │   │   │   │   │   │   ├── certifications/
    │   │   │   │   │   │   │   │   ├── [certId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit certification
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/certifications
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/users/me/certifications/{cert_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Certification detail
    │   │   │   │   │   │   │   │   │       # BE: users-be/certifications
    │   │   │   │   │   │   │   │   │       # GET /v1/users/me/certifications/{cert_id}
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add certification
    │   │   │   │   │   │   │   │   │       # - Certification name
    │   │   │   │   │   │   │   │   │       # - Issuing organization
    │   │   │   │   │   │   │   │   │       # - Issue date
    │   │   │   │   │   │   │   │   │       # - Expiry date (if any)
    │   │   │   │   │   │   │   │   │       # - Credential ID
    │   │   │   │   │   │   │   │   │       # - Credential URL
    │   │   │   │   │   │   │   │   │       # - Certificate upload
    │   │   │   │   │   │   │   │   │       # BE: users-be/credentials
    │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/certifications
    │   │   │   │   │   │   │   │   │       # BE: storage-be/uploads
    │   │   │   │   │   │   │   │   │       # - Credential URL/ID
    │   │   │   │   │   │   │   │   │       # - Upload certificate
    │   │   │   │   │   │   │   │   │       # BE: users-be/certifications, storage-be/asset
    │   │   │   │   │   │   │   │   │       # POST /v1/users/me/certifications
    │   │   │   │   │   │   │   │   ├── verify/
    │   │   │   │   │   │   │   │   │   ├── [certificationId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Verification request
    │   │   │   │   │   │   │   │   │   │       # - Submit for verification
    │   │   │   │   │   │   │   │   │   │       # - Upload proof
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/credentials
    │   │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/certifications/{cert_id}/verify
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Verify certifications
    │   │   │   │   │   │   │   │   │       # - Verification requests
    │   │   │   │   │   │   │   │   │       # - Badge display
    │   │   │   │   │   │   │   │   │       # BE: users-be/certifications
    │   │   │   │   │   │   │   │   │       # POST /v1/users/me/certifications/{cert_id}/verify
    │   │   │   │   │   │   │   │   └── page.tsx  # Certifications list
    │   │   │   │   │   │   │   │       # - External certifications
    │   │   │   │   │   │   │   │       # - Platform certifications
    │   │   │   │   │   │   │   │       # - Badges earned
    │   │   │   │   │   │   │   │       # BE: users-be/credentials
    │   │   │   │   │   │   │   │       # GET /v1/users/{id}/credentials
    │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   └── page.tsx  # Edit profile form
    │   │   │   │   │   │   │   │       # - Basic info (name, title, bio)
    │   │   │   │   │   │   │   │       # - Profile photo upload
    │   │   │   │   │   │   │   │       # - Location
    │   │   │   │   │   │   │   │       # - Languages
    │   │   │   │   │   │   │   │       # - Hourly rate (freelancer)
    │   │   │   │   │   │   │   │       # - Professional headline
    │   │   │   │   │   │   │   │       # BE: users-be/profile
    │   │   │   │   │   │   │   │       # PATCH /v1/users/{id}/profile
    │   │   │   │   │   │   │   │       # BE: storage-be/uploads
    │   │   │   │   │   │   │   │       # POST /v1/storage/upload (photo)
    │   │   │   │   │   │   │   ├── education/
    │   │   │   │   │   │   │   │   ├── [educationId]/
    │   │   │   │   │   │   │   │   │   └── edit/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Edit education form
    │   │   │   │   │   │   │   │   │           # BE: users-be/education
    │   │   │   │   │   │   │   │   │           # PUT /v1/users/{id}/education/{edu_id}
    │   │   │   │   │   │   │   │   │           # DELETE /v1/users/{id}/education/{edu_id}
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add education form
    │   │   │   │   │   │   │   │   │       # - School/university
    │   │   │   │   │   │   │   │   │       # - Degree/qualification
    │   │   │   │   │   │   │   │   │       # - Field of study
    │   │   │   │   │   │   │   │   │       # - Graduation year
    │   │   │   │   │   │   │   │   │       # - GPA (optional)
    │   │   │   │   │   │   │   │   │       # - Description
    │   │   │   │   │   │   │   │   │       # BE: users-be/education
    │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/education
    │   │   │   │   │   │   │   │   └── page.tsx  # Education list
    │   │   │   │   │   │   │   │       # BE: users-be/education
    │   │   │   │   │   │   │   │       # GET /v1/users/{id}/education
    │   │   │   │   │   │   │   ├── experience/
    │   │   │   │   │   │   │   │   ├── [experienceId]/
    │   │   │   │   │   │   │   │   │   └── edit/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Edit experience form
    │   │   │   │   │   │   │   │   │           # BE: users-be/experience
    │   │   │   │   │   │   │   │   │           # PUT /v1/users/{id}/experience/{exp_id}
    │   │   │   │   │   │   │   │   │           # DELETE /v1/users/{id}/experience/{exp_id}
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add experience form
    │   │   │   │   │   │   │   │   │       # - Company name
    │   │   │   │   │   │   │   │   │       # - Position title
    │   │   │   │   │   │   │   │   │       # - Start/end dates
    │   │   │   │   │   │   │   │   │       # - Current position checkbox
    │   │   │   │   │   │   │   │   │       # - Description (rich text)
    │   │   │   │   │   │   │   │   │       # - Skills used
    │   │   │   │   │   │   │   │   │       # BE: users-be/experience
    │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/experience
    │   │   │   │   │   │   │   │   └── page.tsx  # Work experience list
    │   │   │   │   │   │   │   │       # - List all experience entries
    │   │   │   │   │   │   │   │       # - Add new experience button
    │   │   │   │   │   │   │   │       # - Edit/delete actions
    │   │   │   │   │   │   │   │       # BE: users-be/experience
    │   │   │   │   │   │   │   │       # GET /v1/users/{id}/experience
    │   │   │   │   │   │   │   ├── languages/
    │   │   │   │   │   │   │   │   └── page.tsx  # Language proficiency
    │   │   │   │   │   │   │   │       # - Add languages
    │   │   │   │   │   │   │   │       # - Proficiency levels
    │   │   │   │   │   │   │   │       # BE: users-be/profile
    │   │   │   │   │   │   │   │       # PUT /v1/users/me/languages
    │   │   │   │   │   │   │   ├── portfolio/
    │   │   │   │   │   │   │   │   ├── [portfolioId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit portfolio item
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/portfolio
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/users/{id}/portfolio/{item_id}
    │   │   │   │   │   │   │   │   │   │       # DELETE /v1/users/{id}/portfolio/{item_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Portfolio item detail
    │   │   │   │   │   │   │   │   │       # BE: users-be/portfolio
    │   │   │   │   │   │   │   │   │       # GET /v1/users/{id}/portfolio/{item_id}
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add portfolio item
    │   │   │   │   │   │   │   │   │       # - Project title
    │   │   │   │   │   │   │   │   │       # - Description
    │   │   │   │   │   │   │   │   │       # - Media upload (images/videos)
    │   │   │   │   │   │   │   │   │       # - Project URL
    │   │   │   │   │   │   │   │   │       # - Skills used
    │   │   │   │   │   │   │   │   │       # - Client (optional)
    │   │   │   │   │   │   │   │   │       # - Completion date
    │   │   │   │   │   │   │   │   │       # BE: users-be/portfolio
    │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/portfolio
    │   │   │   │   │   │   │   │   │       # BE: storage-be/uploads
    │   │   │   │   │   │   │   │   │       # POST /v1/storage/upload (media)
    │   │   │   │   │   │   │   │   ├── reorder/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Reorder portfolio items
    │   │   │   │   │   │   │   │   │       # - Drag & drop interface
    │   │   │   │   │   │   │   │   │       # - Set featured items
    │   │   │   │   │   │   │   │   │       # BE: users-be/portfolio
    │   │   │   │   │   │   │   │   │       # PATCH /v1/users/{id}/portfolio/reorder
    │   │   │   │   │   │   │   │   │       # Body: { item_ids: string[] }
    │   │   │   │   │   │   │   │   └── page.tsx  # Portfolio items list
    │   │   │   │   │   │   │   │       # - Grid/list view
    │   │   │   │   │   │   │   │       # - Featured items
    │   │   │   │   │   │   │   │       # - Reorder items (drag & drop)
    │   │   │   │   │   │   │   │       # BE: users-be/portfolio
    │   │   │   │   │   │   │   │       # GET /v1/users/{id}/portfolio
    │   │   │   │   │   │   │   ├── portfolio-items/
    │   │   │   │   │   │   │   │   └── [itemId]/
    │   │   │   │   │   │   │   │       ├── analytics/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Portfolio item analytics
    │   │   │   │   │   │   │   │       │       # - View count
    │   │   │   │   │   │   │   │       │       # - Engagement rate
    │   │   │   │   │   │   │   │       │       # BE: users-be/analytics
    │   │   │   │   │   │   │   │       │       # GET /v1/users/me/portfolio/{item_id}/analytics
    │   │   │   │   │   │   │   │       └── page.tsx  # Portfolio item detail
    │   │   │   │   │   │   │   ├── references/
    │   │   │   │   │   │   │   │   ├── [referenceId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Reference detail
    │   │   │   │   │   │   │   │   │       # - Reference content
    │   │   │   │   │   │   │   │   │       # - Relationship details
    │   │   │   │   │   │   │   │   │       # BE: users-be/references
    │   │   │   │   │   │   │   │   │       # GET /v1/users/me/references/{reference_id}
    │   │   │   │   │   │   │   │   ├── request/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Request reference
    │   │   │   │   │   │   │   │   │       # - Select contact
    │   │   │   │   │   │   │   │   │       # - Reference type
    │   │   │   │   │   │   │   │   │       # BE: users-be/references
    │   │   │   │   │   │   │   │   │       # POST /v1/users/me/references/request
    │   │   │   │   │   │   │   │   └── page.tsx  # References list
    │   │   │   │   │   │   │   │       # - Manage references
    │   │   │   │   │   │   │   │       # - Request new
    │   │   │   │   │   │   │   │       # BE: users-be/references
    │   │   │   │   │   │   │   │       # GET /v1/users/me/references
    │   │   │   │   │   │   │   ├── reviews/
    │   │   │   │   │   │   │   │   └── page.tsx  # Reviews (from profile)
    │   │   │   │   │   │   │   │       # - View received/given
    │   │   │   │   │   │   │   │       # - Submit review
    │   │   │   │   │   │   │   │       # BE: reviews-be/review
    │   │   │   │   │   │   │   │       # GET  /v1/reviews?subject_id=…
    │   │   │   │   │   │   │   │       # POST /v1/reviews
    │   │   │   │   │   │   │   ├── service-catalog/
    │   │   │   │   │   │   │   │   ├── [serviceId]/
    │   │   │   │   │   │   │   │   │   └── edit/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Edit service
    │   │   │   │   │   │   │   │   │           # BE: users-be/service_catalog
    │   │   │   │   │   │   │   │   │           # PUT /v1/users/{id}/services/{service_id}
    │   │   │   │   │   │   │   │   │           # DELETE /v1/users/{id}/services/{service_id}
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add service
    │   │   │   │   │   │   │   │   │       # - Service name
    │   │   │   │   │   │   │   │   │       # - Description
    │   │   │   │   │   │   │   │   │       # - Capabilities required
    │   │   │   │   │   │   │   │   │       # - Delivery time
    │   │   │   │   │   │   │   │   │       # - Pricing
    │   │   │   │   │   │   │   │   │       # - Packages (Basic/Standard/Premium)
    │   │   │   │   │   │   │   │   │       # BE: users-be/service_catalog
    │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/services
    │   │   │   │   │   │   │   │   └── page.tsx  # Service catalog management (freelancer)
    │   │   │   │   │   │   │   │       # - List offered services
    │   │   │   │   │   │   │   │       # - Service packages
    │   │   │   │   │   │   │   │       # - Pricing tiers
    │   │   │   │   │   │   │   │       # BE: users-be/service_catalog
    │   │   │   │   │   │   │   │       # GET /v1/users/{id}/service-catalog
    │   │   │   │   │   │   │   ├── skills/
    │   │   │   │   │   │   │   │   ├── specializations/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Specializations & niche expertise
    │   │   │   │   │   │   │   │   │       # - Add specializations
    │   │   │   │   │   │   │   │   │       # - Verification status
    │   │   │   │   │   │   │   │   │       # - Niche expertise tags
    │   │   │   │   │   │   │   │   │       # BE: users-be/capabilities
    │   │   │   │   │   │   │   │   │       # GET /v1/users/{id}/specializations
    │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/specializations
    │   │   │   │   │   │   │   │   └── page.tsx  # Skills management
    │   │   │   │   │   │   │   │       # - List current skills with levels
    │   │   │   │   │   │   │   │       # - Add new skills (autocomplete)
    │   │   │   │   │   │   │   │       # - Edit skill proficiency
    │   │   │   │   │   │   │   │       # - Remove skills
    │   │   │   │   │   │   │   │       # - Primary skills (max 5)
    │   │   │   │   │   │   │   │       # BE: users-be/capabilities
    │   │   │   │   │   │   │   │       # GET /v1/users/{id}/skills
    │   │   │   │   │   │   │   │       # POST /v1/users/{id}/skills
    │   │   │   │   │   │   │   │       # PUT /v1/users/{id}/skills/{skill_id}
    │   │   │   │   │   │   │   │       # DELETE /v1/users/{id}/skills/{skill_id}
    │   │   │   │   │   │   │   │       # BE: search-be/taxonomy
    │   │   │   │   │   │   │   │       # GET /v1/taxonomy/skills (autocomplete)
    │   │   │   │   │   │   │   ├── social-links/
    │   │   │   │   │   │   │   │   └── page.tsx  # Social media links
    │   │   │   │   │   │   │   │       # - LinkedIn, Twitter, GitHub
    │   │   │   │   │   │   │   │       # - Portfolio websites
    │   │   │   │   │   │   │   │       # BE: users-be/profile
    │   │   │   │   │   │   │   │       # PUT /v1/users/me/social-links
    │   │   │   │   │   │   │   ├── verification/
    │   │   │   │   │   │   │   │   ├── identity/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # ID verification
    │   │   │   │   │   │   │   │   │       # - Upload ID document
    │   │   │   │   │   │   │   │   │       # - Selfie verification
    │   │   │   │   │   │   │   │   │       # - Address proof
    │   │   │   │   │   │   │   │   │       # BE: users-be/verification/kyc
    │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/kyc/submit
    │   │   │   │   │   │   │   │   │       # BE: storage-be/uploads
    │   │   │   │   │   │   │   │   │       # BE: admin-be/kyc_case (creates case)
    │   │   │   │   │   │   │   │   ├── phone/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Phone verification
    │   │   │   │   │   │   │   │   │       # - Enter phone number
    │   │   │   │   │   │   │   │   │       # - Receive OTP
    │   │   │   │   │   │   │   │   │       # - Verify OTP
    │   │   │   │   │   │   │   │   │       # BE: users-be/verification
    │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/verify-phone/send
    │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/verify-phone/verify
    │   │   │   │   │   │   │   │   └── page.tsx  # Verification status
    │   │   │   │   │   │   │   │       # - Email verified
    │   │   │   │   │   │   │   │       # - Phone verified
    │   │   │   │   │   │   │   │       # - ID verification status
    │   │   │   │   │   │   │   │       # - Payment method verified
    │   │   │   │   │   │   │   │       # BE: users-be/verification
    │   │   │   │   │   │   │   │       # GET /v1/users/{id}/verification-status
    │   │   │   │   │   │   │   ├── visibility/
    │   │   │   │   │   │   │   │   └── page.tsx  # Profile visibility settings
    │   │   │   │   │   │   │   │       # - Public/private toggle
    │   │   │   │   │   │   │   │       # - Search visibility
    │   │   │   │   │   │   │   │       # - Profile sections visibility
    │   │   │   │   │   │   │   │       # BE: users-be/profile
    │   │   │   │   │   │   │   │       # PUT /v1/users/me/visibility
    │   │   │   │   │   │   │   └── page.tsx  # Profile overview / public view
    │   │   │   │   │   │   │       # - Profile header (photo, name, title, location)
    │   │   │   │   │   │   │       # - Stats (rating, jobs completed, earnings)
    │   │   │   │   │   │   │       # - Skills showcase
    │   │   │   │   │   │   │       # - Portfolio highlights
    │   │   │   │   │   │   │       # - Recent reviews
    │   │   │   │   │   │   │       # - Availability calendar
    │   │   │   │   │   │   │       # BE: users-be/profile
    │   │   │   │   │   │   │       # GET /v1/users/{id}/profile
    │   │   │   │   │   │   │       # BE: users-be/capabilities
    │   │   │   │   │   │   │       # GET /v1/users/{id}/skills
    │   │   │   │   │   │   │       # BE: users-be/portfolio
    │   │   │   │   │   │   │       # GET /v1/users/{id}/portfolio
    │   │   │   │   │   │   │       # BE: reviews-be
    │   │   │   │   │   │   │       # GET /v1/reviews?user_id={id}&limit=5
    │   │   │   │   │   │   │       # BE: users-be/availability
    │   │   │   │   │   │   │       # GET /v1/users/{id}/availability
    │   │   │   │   │   │   ├── proposals/  # Proposals management
    │   │   │   │   │   │   │   ├── [proposalId]/
    │   │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Proposal performance analytics
    │   │   │   │   │   │   │   │   │       # - View metrics
    │   │   │   │   │   │   │   │   │       # - Comparison with others
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/analytics
    │   │   │   │   │   │   │   │   │       # GET /v1/proposals/{proposal_id}/analytics
    │   │   │   │   │   │   │   │   ├── bidding/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Bid status for proposal
    │   │   │   │   │   │   │   │   │       # - Your current bid
    │   │   │   │   │   │   │   │   │       # - Current lowest bid
    │   │   │   │   │   │   │   │   │       # - Update bid
    │   │   │   │   │   │   │   │   │       # - Bid history
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/bidding
    │   │   │   │   │   │   │   │   │       # GET /v1/proposals/{proposal_id}/bid
    │   │   │   │   │   │   │   │   ├── collaborators/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Team collaboration on proposal
    │   │   │   │   │   │   │   │   │       # - Invite team members
    │   │   │   │   │   │   │   │   │       # - Internal notes
    │   │   │   │   │   │   │   │   │       # - Review assignments
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/proposal, users-be/team
    │   │   │   │   │   │   │   │   │       # GET /v1/proposals/{proposal_id}/collaborators
    │   │   │   │   │   │   │   │   │       # POST /v1/proposals/{proposal_id}/collaborators
    │   │   │   │   │   │   │   │   ├── compare/
    │   │   │   │   │   │   │   │   │   ├── [compareWith]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Compare two proposals side-by-side
    │   │   │   │   │   │   │   │   │   │       # - Terms comparison
    │   │   │   │   │   │   │   │   │   │       # - Pricing comparison
    │   │   │   │   │   │   │   │   │   │       # - Deliverables comparison
    │   │   │   │   │   │   │   │   │   │       # - Timeline comparison
    │   │   │   │   │   │   │   │   │   │       # BE: proposals-be/proposal
    │   │   │   │   │   │   │   │   │   │       # GET /v1/proposals/{proposal_id}
    │   │   │   │   │   │   │   │   │   │       # GET /v1/proposals/{compare_with}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Compare with other proposals
    │   │   │   │   │   │   │   │   │       # - Side-by-side comparison
    │   │   │   │   │   │   │   │   │       # - Highlight differences
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/proposal
    │   │   │   │   │   │   │   │   │       # GET /v1/proposals/compare?ids=...
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit proposal
    │   │   │   │   │   │   │   │   │       # - Only if status = DRAFT or PENDING
    │   │   │   │   │   │   │   │   │       # - Update cover letter, rate, timeline
    │   │   │   │   │   │   │   │   │       # BE: proposals-be
    │   │   │   │   │   │   │   │   │       # PUT /v1/proposals/{proposal_id}
    │   │   │   │   │   │   │   │   ├── feedback/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Proposal feedback
    │   │   │   │   │   │   │   │   │       # - Client feedback
    │   │   │   │   │   │   │   │   │       # - Improvement suggestions
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/proposal
    │   │   │   │   │   │   │   │   │       # GET /v1/proposals/{proposal_id}/feedback
    │   │   │   │   │   │   │   │   ├── milestones/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit proposal milestones
    │   │   │   │   │   │   │   │   │   │       # - Add/remove milestones
    │   │   │   │   │   │   │   │   │   │       # - Adjust payment schedule
    │   │   │   │   │   │   │   │   │   │       # BE: proposals-be/milestone
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/proposals/{proposal_id}/milestones
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Milestones overview
    │   │   │   │   │   │   │   │   │       # - Timeline view
    │   │   │   │   │   │   │   │   │       # - Payment breakdown
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/milestone
    │   │   │   │   │   │   │   │   │       # GET /v1/proposals/{proposal_id}/milestones
    │   │   │   │   │   │   │   │   ├── negotiation/
    │   │   │   │   │   │   │   │   │   ├── counter-offer/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create counter-offer
    │   │   │   │   │   │   │   │   │   │       # - Modify terms
    │   │   │   │   │   │   │   │   │   │       # - Adjust pricing
    │   │   │   │   │   │   │   │   │   │       # - Add conditions
    │   │   │   │   │   │   │   │   │   │       # BE: proposals-be/proposal
    │   │   │   │   │   │   │   │   │   │       # POST /v1/proposals/{proposal_id}/counter-offer
    │   │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Negotiation timeline
    │   │   │   │   │   │   │   │   │   │       # - All offers/counter-offers
    │   │   │   │   │   │   │   │   │   │       # - Changes tracking
    │   │   │   │   │   │   │   │   │   │       # - Decision points
    │   │   │   │   │   │   │   │   │   │       # BE: proposals-be/proposal
    │   │   │   │   │   │   │   │   │   │       # GET /v1/proposals/{proposal_id}/negotiation-history
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Active negotiation dashboard
    │   │   │   │   │   │   │   │   │       # - Current offer status
    │   │   │   │   │   │   │   │   │       # - Next steps
    │   │   │   │   │   │   │   │   │       # - Quick actions
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/proposal
    │   │   │   │   │   │   │   │   │       # GET /v1/proposals/{proposal_id}/negotiation-status
    │   │   │   │   │   │   │   │   ├── negotiations/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Proposal negotiations history
    │   │   │   │   │   │   │   │   │       # - Negotiation timeline
    │   │   │   │   │   │   │   │   │       # - Counter-offers
    │   │   │   │   │   │   │   │   │       # - Terms evolution
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/negotiation
    │   │   │   │   │   │   │   │   │       # GET /v1/proposals/{proposal_id}/negotiations
    │   │   │   │   │   │   │   │   ├── questions/
    │   │   │   │   │   │   │   │   │   ├── [questionId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Answer single question
    │   │   │   │   │   │   │   │   │   │       # BE: proposals-be/question
    │   │   │   │   │   │   │   │   │   │       # GET /v1/proposals/{proposal_id}/questions/{question_id}
    │   │   │   │   │   │   │   │   │   │       # POST /v1/proposals/{proposal_id}/questions/{question_id}/answer
    │   │   │   │   │   │   │   │   │   └── page.tsx  # All proposal questions
    │   │   │   │   │   │   │   │   │       # - Q&A thread
    │   │   │   │   │   │   │   │   │       # - Ask new question
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/question
    │   │   │   │   │   │   │   │   │       # GET /v1/proposals/{proposal_id}/questions
    │   │   │   │   │   │   │   │   │       # POST /v1/proposals/{proposal_id}/questions
    │   │   │   │   │   │   │   │   ├── revise/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Revise proposal
    │   │   │   │   │   │   │   │   │       # - Update terms
    │   │   │   │   │   │   │   │   │       # - Adjust bid
    │   │   │   │   │   │   │   │   │       # - Add clarifications
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/proposal
    │   │   │   │   │   │   │   │   │       # PUT /v1/proposals/{proposal_id}/revise
    │   │   │   │   │   │   │   │   ├── versions/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Proposal revision history
    │   │   │   │   │   │   │   │   │       # - Version timeline
    │   │   │   │   │   │   │   │   │       # - Change tracking
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/audit
    │   │   │   │   │   │   │   │   │       # GET /v1/proposals/{proposal_id}/versions
    │   │   │   │   │   │   │   │   ├── withdraw/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Withdraw proposal
    │   │   │   │   │   │   │   │   │       # - Confirmation dialog
    │   │   │   │   │   │   │   │   │       # - Reason for withdrawal
    │   │   │   │   │   │   │   │   │       # - Connects refund info
    │   │   │   │   │   │   │   │   │       # BE: proposals-be
    │   │   │   │   │   │   │   │   │       # POST /v1/proposals/{proposal_id}/withdraw
    │   │   │   │   │   │   │   │   │       # Publishes: ProposalWithdrawn event
    │   │   │   │   │   │   │   │   └── page.tsx  # Proposal detail
    │   │   │   │   │   │   │   │       # - View submitted proposal
    │   │   │   │   │   │   │   │       # - Proposal status
    │   │   │   │   │   │   │   │       # - Client messages/feedback
    │   │   │   │   │   │   │   │       # - Withdraw option (if pending)
    │   │   │   │   │   │   │   │       # BE: proposals-be
    │   │   │   │   │   │   │   │       # GET /v1/proposals/{proposal_id}
    │   │   │   │   │   │   │   ├── accepted/
    │   │   │   │   │   │   │   │   └── page.tsx  # Accepted proposals
    │   │   │   │   │   │   │   │       # BE: proposals-be
    │   │   │   │   │   │   │   │       # GET /v1/proposals/my-proposals?status=accepted
    │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   └── page.tsx  # Proposal analytics (freelancer)
    │   │   │   │   │   │   │   │       # - Total proposals submitted
    │   │   │   │   │   │   │   │       # - Acceptance rate
    │   │   │   │   │   │   │   │       # - Average response time
    │   │   │   │   │   │   │   │       # - Connects spent
    │   │   │   │   │   │   │   │       # BE: proposals-be/analytics
    │   │   │   │   │   │   │   │       # GET /v1/proposals/analytics
    │   │   │   │   │   │   │   ├── archived/
    │   │   │   │   │   │   │   │   └── page.tsx  # Archived proposals
    │   │   │   │   │   │   │   │       # - Old proposals
    │   │   │   │   │   │   │   │       # - Historical reference
    │   │   │   │   │   │   │   │       # BE: proposals-be/proposal
    │   │   │   │   │   │   │   │       # GET /v1/proposals?status=archived
    │   │   │   │   │   │   │   ├── benchmarking/
    │   │   │   │   │   │   │   │   └── page.tsx  # Benchmark against market
    │   │   │   │   │   │   │   │       # - Rate comparisons
    │   │   │   │   │   │   │   │       # - Proposal quality metrics
    │   │   │   │   │   │   │   │       # - Success factors
    │   │   │   │   │   │   │   │       # BE: proposals-be/analytics, search-be/trending
    │   │   │   │   │   │   │   │       # GET /v1/proposals/benchmarking
    │   │   │   │   │   │   │   ├── declined/  # (already listed above; kept once)
    │   │   │   │   │   │   │   │   └── page.tsx  # Declined proposals
    │   │   │   │   │   │   │   │       # - Feedback from client
    │   │   │   │   │   │   │   │       # - Learn from rejections
    │   │   │   │   │   │   │   │       # BE: proposals-be/proposal
    │   │   │   │   │   │   │   │       # GET /v1/proposals?status=declined
    │   │   │   │   │   │   │   ├── drafts/
    │   │   │   │   │   │   │   │   ├── [draftId]/
    │   │   │   │   │   │   │   │   │   └── edit/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Edit proposal draft
    │   │   │   │   │   │   │   │   │           # BE: proposals-be/draft
    │   │   │   │   │   │   │   │   │           # PUT /v1/proposals/drafts/{draft_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Proposal drafts list
    │   │   │   │   │   │   │   │       # BE: proposals-be/draft
    │   │   │   │   │   │   │   │       # GET /v1/proposals/drafts
    │   │   │   │   │   │   │   ├── insights/
    │   │   │   │   │   │   │   │   ├── pricing-analysis/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pricing competitiveness
    │   │   │   │   │   │   │   │   │       # - Market rate comparison
    │   │   │   │   │   │   │   │   │       # - Your pricing vs. competitors
    │   │   │   │   │   │   │   │   │       # - Pricing recommendations
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/analytics, search-be/market-data
    │   │   │   │   │   │   │   │   │       # GET /v1/proposals/insights/pricing
    │   │   │   │   │   │   │   │   │       # GET /v1/market/rates
    │   │   │   │   │   │   │   │   ├── response-time/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Response time analytics
    │   │   │   │   │   │   │   │   │       # - Average response time
    │   │   │   │   │   │   │   │   │       # - Impact on success rate
    │   │   │   │   │   │   │   │   │       # - Optimization tips
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/analytics
    │   │   │   │   │   │   │   │   │       # GET /v1/proposals/insights/response-time
    │   │   │   │   │   │   │   │   ├── win-rate/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Win rate analytics
    │   │   │   │   │   │   │   │   │       # - Success by job type
    │   │   │   │   │   │   │   │   │       # - Success by budget range
    │   │   │   │   │   │   │   │   │       # - Improvement suggestions
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/analytics
    │   │   │   │   │   │   │   │   │       # GET /v1/proposals/insights/win-rate
    │   │   │   │   │   │   │   │   └── page.tsx  # Proposal insights
    │   │   │   │   │   │   │   │       # - Win rate analysis
    │   │   │   │   │   │   │   │       # - Optimal pricing insights
    │   │   │   │   │   │   │   │       # - Response time impact
    │   │   │   │   │   │   │   │       # - Proposal quality score
    │   │   │   │   │   │   │   │       # BE: proposals-be/analytics
    │   │   │   │   │   │   │   │       # GET /v1/proposals/insights
    │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   └── page.tsx  # Pending proposals
    │   │   │   │   │   │   │   │       # BE: proposals-be
    │   │   │   │   │   │   │   │       # GET /v1/proposals/my-proposals?status=pending
    │   │   │   │   │   │   │   ├── portfolio-showcases/
    │   │   │   │   │   │   │   │   ├── [showcaseId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit showcase
    │   │   │   │   │   │   │   │   │   │       # BE: proposals-be/showcase
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/proposals/showcases/{showcase_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Showcase detail
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/showcase
    │   │   │   │   │   │   │   │   │       # GET /v1/proposals/showcases/{showcase_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Manage showcases
    │   │   │   │   │   │   │   │       # - Create showcase
    │   │   │   │   │   │   │   │       # - Link to proposals
    │   │   │   │   │   │   │   │       # BE: proposals-be/showcase
    │   │   │   │   │   │   │   │       # GET /v1/proposals/showcases
    │   │   │   │   │   │   │   │       # POST /v1/proposals/showcases
    │   │   │   │   │   │   │   ├── rejected/
    │   │   │   │   │   │   │   │   └── page.tsx  # Rejected proposals
    │   │   │   │   │   │   │   │       # BE: proposals-be
    │   │   │   │   │   │   │   │       # GET /v1/proposals/my-proposals?status=rejected
    │   │   │   │   │   │   │   ├── submit/
    │   │   │   │   │   │   │   │   └── [jobId]/
    │   │   │   │   │   │   │   │       └── page.tsx  # Submit proposal (freelancer)
    │   │   │   │   │   │   │   │           # - Cover letter (required)
    │   │   │   │   │   │   │   │           # - Proposed rate/budget
    │   │   │   │   │   │   │   │           # - Proposed timeline
    │   │   │   │   │   │   │   │           # - Answer screening questions
    │   │   │   │   │   │   │   │           # - Attachments (portfolio samples)
    │   │   │   │   │   │   │   │           # - Milestones (for fixed-price)
    │   │   │   │   │   │   │   │           # - Terms acceptance
    │   │   │   │   │   │   │   │           # - Connects deduction warning
    │   │   │   │   │   │   │   │           # BE: proposals-be
    │   │   │   │   │   │   │   │           # POST /v1/proposals
    │   │   │   │   │   │   │   │           # Body: { job_id, cover_letter, rate, timeline, ... }
    │   │   │   │   │   │   │   │           # BE: subscriptions-be/connects
    │   │   │   │   │   │   │   │           # POST /v1/connects/deduct
    │   │   │   │   │   │   │   │           # BE: storage-be/uploads
    │   │   │   │   │   │   │   │           # Publishes: ProposalSubmitted event
    │   │   │   │   │   │   │   ├── templates/
    │   │   │   │   │   │   │   │   ├── [templateId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit template
    │   │   │   │   │   │   │   │   │   │       # BE: proposals-be/templates
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/proposals/templates/{template_id}
    │   │   │   │   │   │   │   │   │   │       # DELETE /v1/proposals/templates/{template_id}
    │   │   │   │   │   │   │   │   │   │       # Edit proposal template
    │   │   │   │   │   │   │   │   │   │       # BE: proposals-be/template
    │   │   │   │   │   │   │   │   │   └── use/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Use template for new proposal
    │   │   │   │   │   │   │   │   │           # BE: proposals-be/template
    │   │   │   │   │   │   │   │   │           # GET /v1/proposals/templates/{template_id}
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create template
    │   │   │   │   │   │   │   │   │       # - Template name
    │   │   │   │   │   │   │   │   │       # - Cover letter template
    │   │   │   │   │   │   │   │   │       # - Default rate/terms
    │   │   │   │   │   │   │   │   │       # BE: proposals-be/templates
    │   │   │   │   │   │   │   │   │       # POST /v1/proposals/templates
    │   │   │   │   │   │   │   │   └── page.tsx  # Proposal templates
    │   │   │   │   │   │   │   │       # - List saved templates
    │   │   │   │   │   │   │   │       # - Create new template
    │   │   │   │   │   │   │   │       # BE: proposals-be/templates
    │   │   │   │   │   │   │   │       # GET /v1/proposals/templates
    │   │   │   │   │   │   │   │       # Proposal templates list
    │   │   │   │   │   │   │   │       # - Manage templates
    │   │   │   │   │   │   │   │       # - Create template from proposal
    │   │   │   │   │   │   │   │       # BE: proposals-be/template
    │   │   │   │   │   │   │   ├── withdrawn/
    │   │   │   │   │   │   │   │   └── page.tsx  # Withdrawn proposals
    │   │   │   │   │   │   │   │       # - Self-withdrawn proposals
    │   │   │   │   │   │   │   │       # - Reasons tracking
    │   │   │   │   │   │   │   │       # BE: proposals-be/proposal
    │   │   │   │   │   │   │   │       # GET /v1/proposals?status=withdrawn
    │   │   │   │   │   │   │   └── page.tsx  # Proposals list
    │   │   │   │   │   │   │       # Freelancer: Submitted proposals
    │   │   │   │   │   │   │       # Client: Received proposals (redirect to jobs)
    │   │   │   │   │   │   │       # BE: proposals-be
    │   │   │   │   │   │   │       # GET /v1/proposals/my-proposals
    │   │   │   │   │   │   ├── reputation/
    │   │   │   │   │   │   │   ├── badges/
    │   │   │   │   │   │   │   │   └── page.tsx  # Achievement badges
    │   │   │   │   │   │   │   │       # - Top rated badge
    │   │   │   │   │   │   │   │       # - Rising talent
    │   │   │   │   │   │   │   │       # - Expert verified
    │   │   │   │   │   │   │   │       # BE: reviews-be/badge (if exists) or users-be/profile
    │   │   │   │   │   │   │   │       # GET /v1/reputation/badges
    │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   └── page.tsx  # Review disputes
    │   │   │   │   │   │   │   │       # - Disputed reviews
    │   │   │   │   │   │   │   │       # - Appeal status
    │   │   │   │   │   │   │   │       # BE: reviews-be/dispute (if exists) or admin-be/case_mgmt
    │   │   │   │   │   │   │   │       # GET /v1/reviews/disputes
    │   │   │   │   │   │   │   ├── overview/
    │   │   │   │   │   │   │   │   └── page.tsx  # Reputation overview
    │   │   │   │   │   │   │   │       # - Overall score
    │   │   │   │   │   │   │   │       # - Score breakdown
    │   │   │   │   │   │   │   │       # - Recent changes
    │   │   │   │   │   │   │   │       # BE: reviews-be/reputation (if exists) or reviews-be/review
    │   │   │   │   │   │   │   │       # GET /v1/reputation/overview
    │   │   │   │   │   │   │   ├── reviews-given/
    │   │   │   │   │   │   │   │   └── page.tsx  # Reviews given
    │   │   │   │   │   │   │   │       # - Reviews I left
    │   │   │   │   │   │   │   │       # - Edit option (if within timeframe)
    │   │   │   │   │   │   │   │       # BE: reviews-be/review
    │   │   │   │   │   │   │   │       # GET /v1/reviews/given
    │   │   │   │   │   │   │   └── reviews-received/
    │   │   │   │   │   │   │       └── page.tsx  # Reviews received
    │   │   │   │   │   │   │           # - Client reviews
    │   │   │   │   │   │   │           # - Response option
    │   │   │   │   │   │   │           # BE: reviews-be/review
    │   │   │   │   │   │   │           # GET /v1/reviews/received
    │   │   │   │   │   │   ├── reviews/  # Reviews & ratings
    │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   └── page.tsx  # Review analytics
    │   │   │   │   │   │   │   │       # - Rating trends
    │   │   │   │   │   │   │   │       # - Response rate
    │   │   │   │   │   │   │   │       # - Sentiment analysis
    │   │   │   │   │   │   │   │       # BE: reviews-be/review
    │   │   │   │   │   │   │   │       # GET /v1/users/me/reviews/analytics
    │   │   │   │   │   │   │   ├── badges/
    │   │   │   │   │   │   │   │   └── page.tsx  # Badges & achievements
    │   │   │   │   │   │   │   │       # - Earned badges
    │   │   │   │   │   │   │   │       # - Badge criteria
    │   │   │   │   │   │   │   │       # - Progress towards badges
    │   │   │   │   │   │   │   │       # BE: reviews-be/badges
    │   │   │   │   │   │   │   │       # GET /v1/reviews/badges?user_id={current_user}
    │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create review
    │   │   │   │   │   │   │   │   │       # - Rating (1-5 stars)
    │   │   │   │   │   │   │   │   │       # - Multiple criteria ratings:
    │   │   │   │   │   │   │   │   │       # - Quality of work
    │   │   │   │   │   │   │   │   │       # - Communication
    │   │   │   │   │   │   │   │   │       # - Professionalism
    │   │   │   │   │   │   │   │   │       # - Deadlines
    │   │   │   │   │   │   │   │   │       # - Written feedback (required)
    │   │   │   │   │   │   │   │   │       # - Recommend to others?
    │   │   │   │   │   │   │   │   │       # - Skills demonstrated
    │   │   │   │   │   │   │   │   │       # - Make public/private
    │   │   │   │   │   │   │   │   │       # BE: reviews-be/reviews
    │   │   │   │   │   │   │   │   │       # POST /v1/reviews
    │   │   │   │   │   │   │   │   │       # Body: { contract_id, rating, criteria_ratings, feedback, ... }
    │   │   │   │   │   │   │   │   │       # Publishes: ReviewSubmitted event
    │   │   │   │   │   │   │   │   │       # Triggers: Reputation update (users-be)
    │   │   │   │   │   │   │   │   └── pending/
    │   │   │   │   │   │   │   │       └── page.tsx  # Pending reviews
    │   │   │   │   │   │   │   │           # - Contracts awaiting review
    │   │   │   │   │   │   │   │           # - Reminders
    │   │   │   │   │   │   │   │           # BE: reviews-be/reviews
    │   │   │   │   │   │   │   │           # GET /v1/reviews/pending
    │   │   │   │   │   │   │   ├── dispute/
    │   │   │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Review dispute detail
    │   │   │   │   │   │   │   │   │       # - Dispute reason
    │   │   │   │   │   │   │   │   │       # - Resolution status
    │   │   │   │   │   │   │   │   │       # BE: reviews-be/dispute
    │   │   │   │   │   │   │   │   │       # GET /v1/reviews/disputes/{dispute_id}
    │   │   │   │   │   │   │   │   ├── submit/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute a review
    │   │   │   │   │   │   │   │   │       # - Select review
    │   │   │   │   │   │   │   │   │       # - Dispute reason
    │   │   │   │   │   │   │   │   │       # - Evidence
    │   │   │   │   │   │   │   │   │       # BE: reviews-be/dispute
    │   │   │   │   │   │   │   │   │       # POST /v1/reviews/{review_id}/dispute
    │   │   │   │   │   │   │   │   └── page.tsx  # Review disputes list
    │   │   │   │   │   │   │   │       # BE: reviews-be/dispute
    │   │   │   │   │   │   │   │       # GET /v1/reviews/disputes
    │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Review dispute details
    │   │   │   │   │   │   │   │   │       # - Evidence submission
    │   │   │   │   │   │   │   │   │       # - Admin review status
    │   │   │   │   │   │   │   │   │       # BE: reviews-be/review, admin-be/case_mgmt
    │   │   │   │   │   │   │   │   │       # GET /v1/reviews/{review_id}/dispute
    │   │   │   │   │   │   │   │   └── page.tsx  # Review disputes list
    │   │   │   │   │   │   │   │       # BE: reviews-be/review
    │   │   │   │   │   │   │   │       # GET /v1/reviews/disputes
    │   │   │   │   │   │   │   ├── given/
    │   │   │   │   │   │   │   │   ├── [reviewId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit given review
    │   │   │   │   │   │   │   │   │   │       # BE: reviews-be/review
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/reviews/{review_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Given review details
    │   │   │   │   │   │   │   │   │       # BE: reviews-be/review
    │   │   │   │   │   │   │   │   │       # GET /v1/reviews/{review_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Reviews given list
    │   │   │   │   │   │   │   │       # BE: reviews-be/reviews
    │   │   │   │   │   │   │   │       # GET /v1/reviews?user_id={current_user}&type=given
    │   │   │   │   │   │   │   │       # Given reviews list
    │   │   │   │   │   │   │   │       # BE: reviews-be/review
    │   │   │   │   │   │   │   │       # GET /v1/users/me/reviews/given
    │   │   │   │   │   │   │   │       # Reviews given by user
    │   │   │   │   │   │   │   │       # - All reviews posted
    │   │   │   │   │   │   │   │       # - Edit capability (time-limited)
    │   │   │   │   │   │   │   │       # GET /v1/reviews/given
    │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Leave review form
    │   │   │   │   │   │   │   │   │       # BE: reviews-be/review, contracts-be/contract
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}
    │   │   │   │   │   │   │   │   │       # POST /v1/reviews
    │   │   │   │   │   │   │   │   └── page.tsx  # Pending reviews to complete
    │   │   │   │   │   │   │   │       # BE: reviews-be/review
    │   │   │   │   │   │   │   │       # GET /v1/reviews/pending
    │   │   │   │   │   │   │   │       # Pending reviews to write
    │   │   │   │   │   │   │   │       # - Completed contracts
    │   │   │   │   │   │   │   │       # - Review prompts
    │   │   │   │   │   │   │   ├── received/
    │   │   │   │   │   │   │   │   ├── [reviewId]/
    │   │   │   │   │   │   │   │   │   ├── respond/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Respond to review
    │   │   │   │   │   │   │   │   │   │       # - Public response
    │   │   │   │   │   │   │   │   │   │       # - Character limit
    │   │   │   │   │   │   │   │   │   │       # BE: reviews-be/reviews
    │   │   │   │   │   │   │   │   │   │       # POST /v1/reviews/{review_id}/respond
    │   │   │   │   │   │   │   │   │   │       # Publishes: ReviewResponded event
    │   │   │   │   │   │   │   │   │   │       # BE: reviews-be/review
    │   │   │   │   │   │   │   │   │   │       # POST /v1/reviews/{review_id}/response
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Review detail
    │   │   │   │   │   │   │   │   │       # - Full review content
    │   │   │   │   │   │   │   │   │       # - Reviewer info
    │   │   │   │   │   │   │   │   │       # - Related contract
    │   │   │   │   │   │   │   │   │       # - Response option
    │   │   │   │   │   │   │   │   │       # BE: reviews-be/reviews
    │   │   │   │   │   │   │   │   │       # GET /v1/reviews/{review_id}
    │   │   │   │   │   │   │   │   │       # Review details
    │   │   │   │   │   │   │   │   │       # BE: reviews-be/review
    │   │   │   │   │   │   │   │   └── page.tsx  # Reviews received list
    │   │   │   │   │   │   │   │       # - All reviews received
    │   │   │   │   │   │   │   │       # - Filter by rating
    │   │   │   │   │   │   │   │       # - Filter by contract
    │   │   │   │   │   │   │   │       # BE: reviews-be/reviews
    │   │   │   │   │   │   │   │       # GET /v1/reviews?user_id={current_user}&type=received
    │   │   │   │   │   │   │   │       # Received reviews list
    │   │   │   │   │   │   │   │       # BE: reviews-be/review
    │   │   │   │   │   │   │   │       # GET /v1/users/me/reviews/received
    │   │   │   │   │   │   │   │       # Reviews received
    │   │   │   │   │   │   │   │       # - All reviews about user
    │   │   │   │   │   │   │   │       # - Response capability
    │   │   │   │   │   │   │   │       # GET /v1/reviews/received
    │   │   │   │   │   │   │   ├── stats/
    │   │   │   │   │   │   │   │   └── page.tsx  # Detailed statistics
    │   │   │   │   │   │   │   │       # - Rating breakdown
    │   │   │   │   │   │   │   │       # - Review trends over time
    │   │   │   │   │   │   │   │       # - Category-specific ratings
    │   │   │   │   │   │   │   │       # - Comparison to platform average
    │   │   │   │   │   │   │   │       # BE: reviews-be/stats
    │   │   │   │   │   │   │   │       # GET /v1/reviews/stats/detailed?user_id={current_user}
    │   │   │   │   │   │   │   └── page.tsx  # Reviews overview
    │   │   │   │   │   │   │       # - Reviews received
    │   │   │   │   │   │   │       # - Reviews given
    │   │   │   │   │   │   │       # - Overall rating stats
    │   │   │   │   │   │   │       # - Badges earned
    │   │   │   │   │   │   │       # BE: reviews-be/reviews
    │   │   │   │   │   │   │       # GET /v1/reviews?user_id={current_user}
    │   │   │   │   │   │   │       # BE: reviews-be/stats
    │   │   │   │   │   │   │       # GET /v1/reviews/stats?user_id={current_user}
    │   │   │   │   │   │   ├── search/  # Advanced search functionality
    │   │   │   │   │   │   │   ├── advanced/
    │   │   │   │   │   │   │   │   └── page.tsx  # Advanced search interface
    │   │   │   │   │   │   │   │       # - Complex filters builder
    │   │   │   │   │   │   │   │       # - Boolean operators
    │   │   │   │   │   │   │   │       # - Saved search management
    │   │   │   │   │   │   │   │       # BE: search-be/query
    │   │   │   │   │   │   │   │       # POST /v1/search/advanced
    │   │   │   │   │   │   │   ├── freelancers/
    │   │   │   │   │   │   │   │   ├── advanced/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Advanced freelancer search
    │   │   │   │   │   │   │   │   │       # - Multiple filters
    │   │   │   │   │   │   │   │   │       # - Boolean operators
    │   │   │   │   │   │   │   │   │       # - Saved search option
    │   │   │   │   │   │   │   │   │       # BE: search-be/query
    │   │   │   │   │   │   │   │   │       # POST /v1/search/freelancers/advanced
    │   │   │   │   │   │   │   │   └── page.tsx  # Advanced freelancer search (client)
    │   │   │   │   │   │   │   │       # - Search by skills
    │   │   │   │   │   │   │   │       # - Experience level
    │   │   │   │   │   │   │   │       # - Hourly rate range
    │   │   │   │   │   │   │   │       # - Location
    │   │   │   │   │   │   │   │       # - Availability
    │   │   │   │   │   │   │   │       # - Rating
    │   │   │   │   │   │   │   │       # - Portfolio keywords
    │   │   │   │   │   │   │   │       # BE: search-be/query
    │   │   │   │   │   │   │   │       # POST /v1/search/freelancers
    │   │   │   │   │   │   │   │       # Body: { query, filters: {...}, sort, page }
    │   │   │   │   │   │   │   │       # Basic freelancer search
    │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   └── page.tsx  # Search history
    │   │   │   │   │   │   │   │       # BE: search-be/query
    │   │   │   │   │   │   │   │       # GET /v1/search/history
    │   │   │   │   │   │   │   ├── jobs/
    │   │   │   │   │   │   │   │   ├── advanced/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Advanced job search
    │   │   │   │   │   │   │   │   │       # - Complex filters
    │   │   │   │   │   │   │   │   │       # - Save search
    │   │   │   │   │   │   │   │   │       # - Full-text search
    │   │   │   │   │   │   │   │   │       # - Faceted filters
    │   │   │   │   │   │   │   │   │       # - Autocomplete suggestions
    │   │   │   │   │   │   │   │   │       # - Search history
    │   │   │   │   │   │   │   │   │       # BE: search-be/query
    │   │   │   │   │   │   │   │   │       # POST /v1/search/jobs/advanced
    │   │   │   │   │   │   │   │   │       # BE: search-be/suggestions
    │   │   │   │   │   │   │   │   │       # GET /v1/suggestions?q={query}
    │   │   │   │   │   │   │   │   └── page.tsx  # Advanced job search
    │   │   │   │   │   │   │   │       # - Full-text search
    │   │   │   │   │   │   │   │       # - Faceted filters
    │   │   │   │   │   │   │   │       # - Autocomplete suggestions
    │   │   │   │   │   │   │   │       # - Search history
    │   │   │   │   │   │   │   │       # - Save search
    │   │   │   │   │   │   │   │       # BE: search-be/query
    │   │   │   │   │   │   │   │       # POST /v1/search/jobs
    │   │   │   │   │   │   │   │       # Body: { query, filters: {...}, sort, page }
    │   │   │   │   │   │   │   │       # BE: search-be/suggestions
    │   │   │   │   │   │   │   │       # GET /v1/suggestions?q={query}
    │   │   │   │   │   │   │   │       # Basic job search
    │   │   │   │   │   │   │   ├── portfolio/
    │   │   │   │   │   │   │   │   └── page.tsx  # Search portfolios
    │   │   │   │   │   │   │   │       # - Search by project keywords
    │   │   │   │   │   │   │   │       # - Filter by skills used
    │   │   │   │   │   │   │   │       # - Filter by industry
    │   │   │   │   │   │   │   │       # BE: search-be/portfolio
    │   │   │   │   │   │   │   │       # POST /v1/search/portfolios
    │   │   │   │   │   │   │   ├── portfolios/
    │   │   │   │   │   │   │   │   └── page.tsx  # Search portfolios
    │   │   │   │   │   │   │   │       # - Search by category
    │   │   │   │   │   │   │   │       # - Filter by tags
    │   │   │   │   │   │   │   │       # BE: search-be/query
    │   │   │   │   │   │   │   │       # GET /v1/search/portfolios
    │   │   │   │   │   │   │   ├── recommendations/
    │   │   │   │   │   │   │   │   └── page.tsx  # Personalized recommendations
    │   │   │   │   │   │   │   │       # - AI-powered job matches
    │   │   │   │   │   │   │   │       # - Talent suggestions
    │   │   │   │   │   │   │   │       # BE: search-be/recommendation
    │   │   │   │   │   │   │   │       # GET /v1/recommendations/personalized
    │   │   │   │   │   │   │   ├── saved/
    │   │   │   │   │   │   │   │   ├── [searchId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit saved search
    │   │   │   │   │   │   │   │   │   │       # BE: search-be/saved-search
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/search/saved-searches/{search_id}
    │   │   │   │   │   │   │   │   │   └── results/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # View results from saved search
    │   │   │   │   │   │   │   │   │           # BE: search-be/saved-search, search-be/query
    │   │   │   │   │   │   │   │   │           # GET /v1/search/saved-searches/{search_id}/results
    │   │   │   │   │   │   │   │   └── page.tsx  # Saved searches list (may be in combined, ensuring here)
    │   │   │   │   │   │   │   │       # Saved searches list
    │   │   │   │   │   │   │   ├── saved-searches/
    │   │   │   │   │   │   │   │   ├── [searchId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit saved search
    │   │   │   │   │   │   │   │   │   │       # BE: search-be/saved_search
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/saved-searches/{search_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Execute saved search
    │   │   │   │   │   │   │   │   │       # BE: search-be/saved_search
    │   │   │   │   │   │   │   │   │       # GET /v1/saved-searches/{search_id}/results
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create saved search
    │   │   │   │   │   │   │   │   │       # - Name the search
    │   │   │   │   │   │   │   │   │       # - Set alert frequency
    │   │   │   │   │   │   │   │   │       # - Save filters
    │   │   │   │   │   │   │   │   │       # BE: search-be/saved_search
    │   │   │   │   │   │   │   │   │       # POST /v1/saved-searches
    │   │   │   │   │   │   │   │   └── page.tsx  # Saved searches list
    │   │   │   │   │   │   │   │       # - List all saved searches
    │   │   │   │   │   │   │   │       # - Email alerts toggle
    │   │   │   │   │   │   │   │       # - Edit search
    │   │   │   │   │   │   │   │       # - Delete search
    │   │   │   │   │   │   │   │       # BE: search-be/saved_search
    │   │   │   │   │   │   │   │       # GET /v1/saved-searches
    │   │   │   │   │   │   │   └── trending/
    │   │   │   │   │   │   │       └── page.tsx  # Trending searches and jobs
    │   │   │   │   │   │   │           # BE: search-be/trending
    │   │   │   │   │   │   │           # GET /v1/trending/jobs
    │   │   │   │   │   │   │           # GET /v1/trending/skills
    │   │   │   │   │   │   ├── settings/  # User settings
    │   │   │   │   │   │   │   ├── accessibility/
    │   │   │   │   │   │   │   │   └── page.tsx  # Accessibility preferences
    │   │   │   │   │   │   │   │       # - Screen reader settings
    │   │   │   │   │   │   │   │       # - Keyboard shortcuts
    │   │   │   │   │   │   │   │       # - High contrast mode
    │   │   │   │   │   │   │   │       # - Motion reduction
    │   │   │   │   │   │   │   │       # BE: users-be/preferences
    │   │   │   │   │   │   │   │       # PUT /v1/users/me/accessibility
    │   │   │   │   │   │   │   ├── account/
    │   │   │   │   │   │   │   │   ├── close/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Close account
    │   │   │   │   │   │   │   │   │       # - Confirmation steps
    │   │   │   │   │   │   │   │   │       # - Data retention options
    │   │   │   │   │   │   │   │   │       # - Final warning
    │   │   │   │   │   │   │   │   │       # BE: users-be/account
    │   │   │   │   │   │   │   │   │       # POST /v1/users/me/close-account
    │   │   │   │   │   │   │   │   ├── data-export/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Export user data (GDPR)
    │   │   │   │   │   │   │   │   │       # - Request export
    │   │   │   │   │   │   │   │   │       # - Download data
    │   │   │   │   │   │   │   │   │       # BE: users-be/account
    │   │   │   │   │   │   │   │   │       # POST /v1/users/me/data-export
    │   │   │   │   │   │   │   │   ├── deactivate/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Deactivate account
    │   │   │   │   │   │   │   │   │       # - Temporary suspension
    │   │   │   │   │   │   │   │   │       # - Reactivation info
    │   │   │   │   │   │   │   │   │       # BE: users-be/account
    │   │   │   │   │   │   │   │   │       # POST /v1/users/me/deactivate
    │   │   │   │   │   │   │   │   ├── delete/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Delete account
    │   │   │   │   │   │   │   │   │       # - Reason for deletion
    │   │   │   │   │   │   │   │   │       # - Data export option
    │   │   │   │   │   │   │   │   │       # - Confirmation (password + checkbox)
    │   │   │   │   │   │   │   │   │       # - GDPR compliance
    │   │   │   │   │   │   │   │   │       # BE: users-be/account
    │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/delete-account
    │   │   │   │   │   │   │   │   │       # Publishes: AccountDeleted event
    │   │   │   │   │   │   │   │   ├── email/
    │   │   │   │   │   │   │   │   │   └── change/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Change email
    │   │   │   │   │   │   │   │   │           # - New email input
    │   │   │   │   │   │   │   │   │           # - Password confirmation
    │   │   │   │   │   │   │   │   │           # - Verification email sent
    │   │   │   │   │   │   │   │   │           # BE: users-be/user
    │   │   │   │   │   │   │   │   │           # POST /v1/users/{id}/change-email
    │   │   │   │   │   │   │   │   ├── phone/
    │   │   │   │   │   │   │   │   │   └── change/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Change phone number
    │   │   │   │   │   │   │   │   │           # - New phone input
    │   │   │   │   │   │   │   │   │           # - OTP verification
    │   │   │   │   │   │   │   │   │           # BE: users-be/user
    │   │   │   │   │   │   │   │   │           # POST /v1/users/{id}/change-phone
    │   │   │   │   │   │   │   │   ├── reactivate/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Reactivate account
    │   │   │   │   │   │   │   │   │       # - Restore account
    │   │   │   │   │   │   │   │   │       # - Update information
    │   │   │   │   │   │   │   │   │       # BE: users-be/account
    │   │   │   │   │   │   │   │   │       # POST /v1/users/me/reactivate
    │   │   │   │   │   │   │   │   └── username/
    │   │   │   │   │   │   │   │       └── change/
    │   │   │   │   │   │   │   │           └── page.tsx  # Change username
    │   │   │   │   │   │   │   │               # - New username
    │   │   │   │   │   │   │   │               # - Availability check
    │   │   │   │   │   │   │   │               # BE: users-be/user
    │   │   │   │   │   │   │   │               # POST /v1/users/{id}/change-username
    │   │   │   │   │   │   │   ├── advanced/
    │   │   │   │   │   │   │   │   └── page.tsx  # Advanced settings
    │   │   │   │   │   │   │   │       # - Debug mode
    │   │   │   │   │   │   │   │       # - Experimental APIs
    │   │   │   │   │   │   │   │       # - Performance monitoring opt-in
    │   │   │   │   │   │   │   │       # BE: users-be/preferences
    │   │   │   │   │   │   │   │       # GET /v1/settings/advanced
    │   │   │   │   │   │   │   │       # PUT /v1/settings/advanced
    │   │   │   │   │   │   │   ├── api/
    │   │   │   │   │   │   │   │   ├── documentation/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # API documentation
    │   │   │   │   │   │   │   │   │       # - API reference
    │   │   │   │   │   │   │   │   │       # - Examples
    │   │   │   │   │   │   │   │   │       # BE: None (static content)
    │   │   │   │   │   │   │   │   └── page.tsx  # API settings (developer section)
    │   │   │   │   │   │   │   ├── authorized-apps/
    │   │   │   │   │   │   │   │   ├── [appId]/
    │   │   │   │   │   │   │   │   │   └── revoke/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Revoke app access
    │   │   │   │   │   │   │   │   │           # BE: users-be/oauth_token
    │   │   │   │   │   │   │   │   │           # DELETE /v1/settings/authorized-apps/{app_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Authorized apps list
    │   │   │   │   │   │   │   │       # - Third-party app permissions
    │   │   │   │   │   │   │   │       # - Last used dates
    │   │   │   │   │   │   │   │       # BE: users-be/oauth_token
    │   │   │   │   │   │   │   │       # GET /v1/settings/authorized-apps
    │   │   │   │   │   │   │   ├── blocked-users/
    │   │   │   │   │   │   │   │   └── page.tsx  # Blocked users management
    │   │   │   │   │   │   │   │       # - View blocked users
    │   │   │   │   │   │   │   │       # - Unblock users
    │   │   │   │   │   │   │   │       # BE: users-be/blocked
    │   │   │   │   │   │   │   │       # GET /v1/users/me/blocked-users
    │   │   │   │   │   │   │   │       # DELETE /v1/users/me/blocked-users/{user_id}
    │   │   │   │   │   │   │   ├── developer/
    │   │   │   │   │   │   │   │   ├── api-keys/
    │   │   │   │   │   │   │   │   │   ├── [keyId]/
    │   │   │   │   │   │   │   │   │   │   ├── regenerate/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Regenerate API key
    │   │   │   │   │   │   │   │   │   │   │       # BE: users-be/api_key
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/developer/api-keys/{key_id}/regenerate
    │   │   │   │   │   │   │   │   │   │   ├── revoke/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Revoke API key
    │   │   │   │   │   │   │   │   │   │   │       # BE: users-be/api_key
    │   │   │   │   │   │   │   │   │   │   │       # DELETE /v1/developer/api-keys/{key_id}
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # API key details
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/api_key
    │   │   │   │   │   │   │   │   │   │       # GET /v1/developer/api-keys/{key_id}
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create API key
    │   │   │   │   │   │   │   │   │   │       # - Key name
    │   │   │   │   │   │   │   │   │   │       # - Permissions/scopes
    │   │   │   │   │   │   │   │   │   │       # - Expiration
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/developer
    │   │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/api-keys
    │   │   │   │   │   │   │   │   │   │       # - Name and permissions
    │   │   │   │   │   │   │   │   │   │       # - Expiration settings
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/api_key
    │   │   │   │   │   │   │   │   │   │       # POST /v1/developer/api-keys
    │   │   │   │   │   │   │   │   │   └── page.tsx  # API keys list
    │   │   │   │   │   │   │   │   │       # - Active keys
    │   │   │   │   │   │   │   │   │       # - Create new key
    │   │   │   │   │   │   │   │   │       # - Revoke key
    │   │   │   │   │   │   │   │   │       # BE: users-be/developer
    │   │   │   │   │   │   │   │   │       # GET /v1/users/{id}/api-keys
    │   │   │   │   │   │   │   │   │       # BE: users-be/api_key
    │   │   │   │   │   │   │   │   │       # GET /v1/developer/api-keys
    │   │   │   │   │   │   │   │   ├── oauth-apps/
    │   │   │   │   │   │   │   │   │   ├── [appId]/
    │   │   │   │   │   │   │   │   │   │   ├── credentials/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth credentials
    │   │   │   │   │   │   │   │   │   │   │       # - Client ID/Secret
    │   │   │   │   │   │   │   │   │   │   │       # - Regenerate secret
    │   │   │   │   │   │   │   │   │   │   │       # BE: users-be/oauth_app
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/developer/oauth-apps/{app_id}/regenerate-secret
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit OAuth app
    │   │   │   │   │   │   │   │   │   │   │       # BE: users-be/oauth_app
    │   │   │   │   │   │   │   │   │   │   │       # PUT /v1/developer/oauth-apps/{app_id}
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth app detail
    │   │   │   │   │   │   │   │   │   │       # - Client ID/secret
    │   │   │   │   │   │   │   │   │   │       # - Edit/delete
    │   │   │   │   │   │   │   │   │   │       # - Usage stats
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/developer
    │   │   │   │   │   │   │   │   │   │       # GET /v1/users/{id}/oauth-apps/{app_id}
    │   │   │   │   │   │   │   │   │   │       # OAuth app details
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/oauth_app
    │   │   │   │   │   │   │   │   │   │       # GET /v1/developer/oauth-apps/{app_id}
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create OAuth app
    │   │   │   │   │   │   │   │   │   │       # - App name
    │   │   │   │   │   │   │   │   │   │       # - Redirect URIs
    │   │   │   │   │   │   │   │   │   │       # - Scopes
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/developer
    │   │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/oauth-apps
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/oauth_app
    │   │   │   │   │   │   │   │   │   │       # POST /v1/developer/oauth-apps
    │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth apps list
    │   │   │   │   │   │   │   │   │       # BE: users-be/oauth_app
    │   │   │   │   │   │   │   │   │       # GET /v1/developer/oauth-apps
    │   │   │   │   │   │   │   │   └── page.tsx  # Developer settings
    │   │   │   │   │   │   │   │       # - API keys
    │   │   │   │   │   │   │   │       # - OAuth applications
    │   │   │   │   │   │   │   │       # - API usage stats
    │   │   │   │   │   │   │   │       # BE: users-be/developer
    │   │   │   │   │   │   │   │       # GET /v1/users/{id}/developer
    │   │   │   │   │   │   │   │       # Developer settings hub
    │   │   │   │   │   │   │   │       # BE: None (navigation)
    │   │   │   │   │   │   │   ├── devices/
    │   │   │   │   │   │   │   │   └── page.tsx  # Connected devices
    │   │   │   │   │   │   │   │       # - Active sessions
    │   │   │   │   │   │   │   │       # - Device trust management
    │   │   │   │   │   │   │   │       # - Logout devices
    │   │   │   │   │   │   │   │       # BE: users-be/sessions
    │   │   │   │   │   │   │   │       # GET /v1/users/me/devices
    │   │   │   │   │   │   │   │       # DELETE /v1/users/me/devices/{device_id}
    │   │   │   │   │   │   │   ├── integrations/
    │   │   │   │   │   │   │   │   ├── available/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Available integrations
    │   │   │   │   │   │   │   │   │       # - Browse integrations
    │   │   │   │   │   │   │   │   │       # - Connect new
    │   │   │   │   │   │   │   │   │       # BE: users-be/integrations
    │   │   │   │   │   │   │   │   │       # GET /v1/integrations/available
    │   │   │   │   │   │   │   │   ├── calendar/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Calendar integration
    │   │   │   │   │   │   │   │   │       # - Google Calendar
    │   │   │   │   │   │   │   │   │       # - Outlook Calendar
    │   │   │   │   │   │   │   │   │       # - Sync settings
    │   │   │   │   │   │   │   │   │       # BE: users-be/integrations
    │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/integrations/calendar
    │   │   │   │   │   │   │   │   ├── connected/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Connected integrations
    │   │   │   │   │   │   │   │   │       # - Active integrations
    │   │   │   │   │   │   │   │   │       # - Disconnect options
    │   │   │   │   │   │   │   │   │       # BE: users-be/integrations
    │   │   │   │   │   │   │   │   │       # GET /v1/users/me/integrations
    │   │   │   │   │   │   │   │   ├── slack/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Slack integration
    │   │   │   │   │   │   │   │   │       # - Connect Slack workspace
    │   │   │   │   │   │   │   │   │       # - Notification channels
    │   │   │   │   │   │   │   │   │       # BE: users-be/integrations
    │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/integrations/slack
    │   │   │   │   │   │   │   │   └── webhooks/
    │   │   │   │   │   │   │   │       ├── [webhookId]/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Webhook detail
    │   │   │   │   │   │   │   │       │       # - Delivery logs
    │   │   │   │   │   │   │   │       │       # - Test webhook
    │   │   │   │   │   │   │   │       │       # - Edit/delete
    │   │   │   │   │   │   │   │       │       # BE: users-be/integrations
    │   │   │   │   │   │   │   │       │       # GET /v1/users/{id}/webhooks/{webhook_id}
    │   │   │   │   │   │   │   │       │       # PUT /v1/users/{id}/webhooks/{webhook_id}
    │   │   │   │   │   │   │   │       │       # DELETE /v1/users/{id}/webhooks/{webhook_id}
    │   │   │   │   │   │   │   │       └── create/
    │   │   │   │   │   │   │   │           └── page.tsx  # Create webhook
    │   │   │   │   │   │   │   │               # - Webhook URL
    │   │   │   │   │   │   │   │               # - Secret key
    │   │   │   │   │   │   │   │               # - Events to subscribe
    │   │   │   │   │   │   │   │               # BE: users-be/integrations
    │   │   │   │   │   │   │   │               # POST /v1/users/{id}/webhooks
    │   │   │   │   │   │   │   ├── labs/
    │   │   │   │   │   │   │   │   └── page.tsx  # Experimental features
    │   │   │   │   │   │   │   │       # - Beta feature toggles
    │   │   │   │   │   │   │   │       # - Early access programs
    │   │   │   │   │   │   │   │       # BE: utility-be/feature_flags
    │   │   │   │   │   │   │   │       # GET /v1/labs/features
    │   │   │   │   │   │   │   │       # PUT /v1/labs/features/{feature_id}/toggle
    │   │   │   │   │   │   │   ├── login-history/
    │   │   │   │   │   │   │   │   └── page.tsx  # Login history
    │   │   │   │   │   │   │   │       # - Recent logins
    │   │   │   │   │   │   │   │       # - Location tracking
    │   │   │   │   │   │   │   │       # - Security alerts
    │   │   │   │   │   │   │   │       # BE: users-be/audit
    │   │   │   │   │   │   │   │       # GET /v1/users/me/login-history
    │   │   │   │   │   │   │   ├── notifications/
    │   │   │   │   │   │   │   │   ├── digest/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Email digest settings
    │   │   │   │   │   │   │   │   │       # - Daily/weekly/monthly
    │   │   │   │   │   │   │   │   │       # - Content preferences
    │   │   │   │   │   │   │   │   │       # BE: communications-be/preferences
    │   │   │   │   │   │   │   │   │       # PUT /v1/notifications/digest-preferences
    │   │   │   │   │   │   │   │   ├── email/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Email notification settings
    │   │   │   │   │   │   │   │   │       # - Per-category toggles
    │   │   │   │   │   │   │   │   │       # - Digest preferences
    │   │   │   │   │   │   │   │   │       # BE: communications-be/preferences
    │   │   │   │   │   │   │   │   │       # PUT /v1/notifications/email-preferences
    │   │   │   │   │   │   │   │   └── push/
    │   │   │   │   │   │   │   │       └── page.tsx  # Push notification settings
    │   │   │   │   │   │   │   │           # - Device management
    │   │   │   │   │   │   │   │           # - Per-category toggles
    │   │   │   │   │   │   │   │           # BE: communications-be/preferences
    │   │   │   │   │   │   │   │           # PUT /v1/notifications/push-preferences
    │   │   │   │   │   │   │   ├── preferences/
    │   │   │   │   │   │   │   │   ├── accessibility/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Accessibility settings
    │   │   │   │   │   │   │   │   │       # - Screen reader optimizations
    │   │   │   │   │   │   │   │   │       # - High contrast mode
    │   │   │   │   │   │   │   │   │       # - Font size
    │   │   │   │   │   │   │   │   │       # - Keyboard shortcuts
    │   │   │   │   │   │   │   │   │       # - Motion preferences
    │   │   │   │   │   │   │   │   │       # BE: users-be/preferences
    │   │   │   │   │   │   │   │   │       # PATCH /v1/users/{id}/preferences/accessibility
    │   │   │   │   │   │   │   │   ├── appearance/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Appearance preferences
    │   │   │   │   │   │   │   │   │       # - Theme selection
    │   │   │   │   │   │   │   │   │       # - Color customization
    │   │   │   │   │   │   │   │   │       # - Layout preferences
    │   │   │   │   │   │   │   │   │       # BE: users-be/preferences
    │   │   │   │   │   │   │   │   │       # PUT /v1/users/me/preferences/appearance
    │   │   │   │   │   │   │   │   ├── language/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Language settings
    │   │   │   │   │   │   │   │   │       # - Interface language
    │   │   │   │   │   │   │   │   │       # - Content languages
    │   │   │   │   │   │   │   │   │       # BE: users-be/preferences
    │   │   │   │   │   │   │   │   │       # PATCH /v1/users/{id}/preferences/language
    │   │   │   │   │   │   │   │   │       # Language preferences
    │   │   │   │   │   │   │   │   │       # - Content language
    │   │   │   │   │   │   │   │   │       # PUT /v1/users/me/preferences/language
    │   │   │   │   │   │   │   │   └── timezone/
    │   │   │   │   │   │   │   │       └── page.tsx  # Timezone settings
    │   │   │   │   │   │   │   │           # - Timezone selection
    │   │   │   │   │   │   │   │           # - Time format (12/24 hour)
    │   │   │   │   │   │   │   │           # BE: users-be/preferences
    │   │   │   │   │   │   │   │           # PUT /v1/users/me/preferences/timezone
    │   │   │   │   │   │   │   ├── privacy/
    │   │   │   │   │   │   │   │   ├── activity/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Activity privacy
    │   │   │   │   │   │   │   │   │       # - Who can see activity
    │   │   │   │   │   │   │   │   │       # - Activity history settings
    │   │   │   │   │   │   │   │   │       # BE: users-be/privacy
    │   │   │   │   │   │   │   │   │       # PUT /v1/users/me/privacy/activity
    │   │   │   │   │   │   │   │   ├── blocked-users/
    │   │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Block user
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/privacy
    │   │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/blocked-users
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Blocked users list
    │   │   │   │   │   │   │   │   │       # - List blocked users
    │   │   │   │   │   │   │   │   │       # - Unblock option
    │   │   │   │   │   │   │   │   │       # BE: users-be/privacy
    │   │   │   │   │   │   │   │   │       # GET /v1/users/{id}/blocked-users
    │   │   │   │   │   │   │   │   │       # DELETE /v1/users/{id}/blocked-users/{user_id}
    │   │   │   │   │   │   │   │   ├── data-export/
    │   │   │   │   │   │   │   │   │   ├── request/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Request new data export
    │   │   │   │   │   │   │   │   │   │       # - Select data categories
    │   │   │   │   │   │   │   │   │   │       # - Format (JSON, CSV, PDF)
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/privacy
    │   │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/data-export/request
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Data export (GDPR)
    │   │   │   │   │   │   │   │   │       # - Request data export
    │   │   │   │   │   │   │   │   │       # - Export history
    │   │   │   │   │   │   │   │   │       # - Download exports
    │   │   │   │   │   │   │   │   │       # BE: users-be/privacy
    │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/data-export
    │   │   │   │   │   │   │   │   │       # GET /v1/users/{id}/data-exports
    │   │   │   │   │   │   │   │   ├── data-sharing/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Data sharing preferences
    │   │   │   │   │   │   │   │   │       # - Analytics opt-in/out
    │   │   │   │   │   │   │   │   │       # - Third-party sharing
    │   │   │   │   │   │   │   │   │       # BE: users-be/privacy
    │   │   │   │   │   │   │   │   │       # PUT /v1/users/me/privacy/data-sharing
    │   │   │   │   │   │   │   │   ├── gdpr/
    │   │   │   │   │   │   │   │   │   ├── delete/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # GDPR — Delete account
    │   │   │   │   │   │   │   │   │   │       # - Submit erasure request
    │   │   │   │   │   │   │   │   │   │       # - Track status
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/privacy
    │   │   │   │   │   │   │   │   │   │       # POST /v1/privacy/erase
    │   │   │   │   │   │   │   │   │   │       # GET  /v1/privacy/requests/{id}
    │   │   │   │   │   │   │   │   │   └── export/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # GDPR — Data export
    │   │   │   │   │   │   │   │   │           # - Request export
    │   │   │   │   │   │   │   │   │           # - Track status & download link
    │   │   │   │   │   │   │   │   │           # BE: users-be/privacy
    │   │   │   │   │   │   │   │   │           # POST /v1/privacy/export
    │   │   │   │   │   │   │   │   │           # GET  /v1/privacy/requests/{id}
    │   │   │   │   │   │   │   │   └── profile-visibility/
    │   │   │   │   │   │   │   │       └── page.tsx  # Profile visibility settings
    │   │   │   │   │   │   │   │           # - Search engine visibility
    │   │   │   │   │   │   │   │           # - Profile sections
    │   │   │   │   │   │   │   │           # BE: users-be/privacy
    │   │   │   │   │   │   │   │           # PUT /v1/users/me/privacy/profile-visibility
    │   │   │   │   │   │   │   ├── security/
    │   │   │   │   │   │   │   │   ├── login-history/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Login history
    │   │   │   │   │   │   │   │   │       # - Recent logins
    │   │   │   │   │   │   │   │   │       # - Failed attempts
    │   │   │   │   │   │   │   │   │       # - Device/location info
    │   │   │   │   │   │   │   │   │       # BE: users-be/security/audit
    │   │   │   │   │   │   │   │   │       # GET /v1/users/{id}/login-history
    │   │   │   │   │   │   │   │   ├── password/
    │   │   │   │   │   │   │   │   │   └── change/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Change password
    │   │   │   │   │   │   │   │   │           # - Current password
    │   │   │   │   │   │   │   │   │           # - New password
    │   │   │   │   │   │   │   │   │           # - Password strength meter
    │   │   │   │   │   │   │   │   │           # BE: users-be/security
    │   │   │   │   │   │   │   │   │           # POST /v1/users/{id}/change-password
    │   │   │   │   │   │   │   │   ├── sessions/
    │   │   │   │   │   │   │   │   │   ├── revoke-all/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Revoke all sessions (except current)
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/security/session
    │   │   │   │   │   │   │   │   │   │       # POST /v1/users/{id}/sessions/revoke-all
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Active sessions
    │   │   │   │   │   │   │   │   │       # - List all active sessions
    │   │   │   │   │   │   │   │   │       # - Device info
    │   │   │   │   │   │   │   │   │       # - Location
    │   │   │   │   │   │   │   │   │       # - Last active
    │   │   │   │   │   │   │   │   │       # - Revoke session
    │   │   │   │   │   │   │   │   │       # BE: users-be/security/session
    │   │   │   │   │   │   │   │   │       # GET /v1/users/{id}/sessions
    │   │   │   │   │   │   │   │   │       # DELETE /v1/users/{id}/sessions/{session_id}
    │   │   │   │   │   │   │   │   └── two-factor/
    │   │   │   │   │   │   │   │       ├── disable/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Disable 2FA
    │   │   │   │   │   │   │   │       │       # - Password confirmation
    │   │   │   │   │   │   │   │       │       # - 2FA code verification
    │   │   │   │   │   │   │   │       │       # BE: users-be/security/mfa
    │   │   │   │   │   │   │   │       │       # POST /v1/users/{id}/mfa/disable
    │   │   │   │   │   │   │   │       └── enable/
    │   │   │   │   │   │   │   │           └── page.tsx  # Enable 2FA
    │   │   │   │   │   │   │   │               # - QR code scan
    │   │   │   │   │   │   │   │               # - Verify setup
    │   │   │   │   │   │   │   │               # - Save backup codes
    │   │   │   │   │   │   │   │               # BE: users-be/security/mfa
    │   │   │   │   │   │   │   │               # POST /v1/users/{id}/mfa/enable
    │   │   │   │   │   │   │   ├── two-factor/
    │   │   │   │   │   │   │   │   ├── backup-codes/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Backup codes
    │   │   │   │   │   │   │   │   │       # - Generate codes
    │   │   │   │   │   │   │   │   │       # - Download codes
    │   │   │   │   │   │   │   │   │       # BE: users-be/mfa
    │   │   │   │   │   │   │   │   │       # POST /v1/users/me/mfa/backup-codes
    │   │   │   │   │   │   │   │   ├── disable/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Disable 2FA
    │   │   │   │   │   │   │   │   │       # - Verification required
    │   │   │   │   │   │   │   │   │       # - Confirmation
    │   │   │   │   │   │   │   │   │       # BE: users-be/mfa
    │   │   │   │   │   │   │   │   │       # POST /v1/users/me/mfa/disable
    │   │   │   │   │   │   │   │   └── methods/
    │   │   │   │   │   │   │   │       ├── app/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Authenticator app setup
    │   │   │   │   │   │   │   │       │       # - QR code
    │   │   │   │   │   │   │   │       │       # - Verification
    │   │   │   │   │   │   │   │       │       # BE: users-be/mfa
    │   │   │   │   │   │   │   │       │       # POST /v1/users/me/mfa/app
    │   │   │   │   │   │   │   │       └── sms/
    │   │   │   │   │   │   │   │           └── page.tsx  # SMS 2FA setup
    │   │   │   │   │   │   │   │               # - Phone verification
    │   │   │   │   │   │   │   │               # - Test code
    │   │   │   │   │   │   │   │               # BE: users-be/mfa
    │   │   │   │   │   │   │   │               # POST /v1/users/me/mfa/sms
    │   │   │   │   │   │   │   └── page.tsx  # Settings overview
    │   │   │   │   │   │   │       # - Quick access to all settings
    │   │   │   │   │   │   │       # - Profile completion indicator
    │   │   │   │   │   │   │       # - Recently changed settings
    │   │   │   │   │   │   │       # BE: users-be/preferences
    │   │   │   │   │   │   │       # GET /v1/users/{id}/preferences
    │   │   │   │   │   │   ├── sourcing/
    │   │   │   │   │   │   │   ├── campaigns/
    │   │   │   │   │   │   │   │   ├── [campaignId]/
    │   │   │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Campaign analytics
    │   │   │   │   │   │   │   │   │   │       # - Reach
    │   │   │   │   │   │   │   │   │   │       # - Engagement
    │   │   │   │   │   │   │   │   │   │       # - Conversions
    │   │   │   │   │   │   │   │   │   │       # BE: jobs-be/analytics
    │   │   │   │   │   │   │   │   │   │       # GET /v1/sourcing/campaigns/{campaign_id}/analytics
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit sourcing campaign
    │   │   │   │   │   │   │   │   │   │       # BE: jobs-be/campaign (if exists) or jobs-be/job
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/sourcing/campaigns/{campaign_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Campaign detail
    │   │   │   │   │   │   │   │   │       # BE: jobs-be/campaign
    │   │   │   │   │   │   │   │   │       # GET /v1/sourcing/campaigns/{campaign_id}
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create sourcing campaign
    │   │   │   │   │   │   │   │   │       # - Target criteria
    │   │   │   │   │   │   │   │   │       # - Budget allocation
    │   │   │   │   │   │   │   │   │       # - Messaging
    │   │   │   │   │   │   │   │   │       # BE: jobs-be/campaign
    │   │   │   │   │   │   │   │   │       # POST /v1/sourcing/campaigns
    │   │   │   │   │   │   │   │   └── page.tsx  # Campaigns list
    │   │   │   │   │   │   │   │       # BE: jobs-be/campaign
    │   │   │   │   │   │   │   │       # GET /v1/sourcing/campaigns
    │   │   │   │   │   │   │   ├── invitations/
    │   │   │   │   │   │   │   │   ├── [invitationId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Invitation detail
    │   │   │   │   │   │   │   │   │       # - View status
    │   │   │   │   │   │   │   │   │       # - Resend invitation
    │   │   │   │   │   │   │   │   │       # BE: communications-be/invitation (if exists) or jobs-be/job
    │   │   │   │   │   │   │   │   │       # GET /v1/sourcing/invitations/{invitation_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Sent invitations
    │   │   │   │   │   │   │   │       # - Invitation history
    │   │   │   │   │   │   │   │       # - Response rates
    │   │   │   │   │   │   │   │       # BE: communications-be/invitation
    │   │   │   │   │   │   │   │       # GET /v1/sourcing/invitations
    │   │   │   │   │   │   │   └── talent-pools/
    │   │   │   │   │   │   │       ├── [poolId]/
    │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit talent pool
    │   │   │   │   │   │   │       │   │       # BE: users-be/talent_pool (if exists) or search-be/saved-search
    │   │   │   │   │   │   │       │   │       # PUT /v1/sourcing/talent-pools/{pool_id}
    │   │   │   │   │   │   │       │   ├── members/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Pool members
    │   │   │   │   │   │   │       │   │       # - Add/remove members
    │   │   │   │   │   │   │       │   │       # - Bulk actions
    │   │   │   │   │   │   │       │   │       # BE: users-be/talent_pool
    │   │   │   │   │   │   │       │   │       # GET /v1/sourcing/talent-pools/{pool_id}/members
    │   │   │   │   │   │   │       │   └── page.tsx  # Talent pool detail
    │   │   │   │   │   │   │       │       # BE: users-be/talent_pool
    │   │   │   │   │   │   │       │       # GET /v1/sourcing/talent-pools/{pool_id}
    │   │   │   │   │   │   │       ├── create/
    │   │   │   │   │   │   │       │   └── page.tsx  # Create talent pool
    │   │   │   │   │   │   │       │       # BE: users-be/talent_pool
    │   │   │   │   │   │   │       │       # POST /v1/sourcing/talent-pools
    │   │   │   │   │   │   │       └── page.tsx  # Talent pools list
    │   │   │   │   │   │   │           # BE: users-be/talent_pool
    │   │   │   │   │   │   │           # GET /v1/sourcing/talent-pools
    │   │   │   │   │   │   ├── specialized/
    │   │   │   │   │   │   │   ├── plus/
    │   │   │   │   │   │   │   │   ├── subscribe/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Subscribe to Plus
    │   │   │   │   │   │   │   │   │       # - Plans
    │   │   │   │   │   │   │   │   │       # - Payment
    │   │   │   │   │   │   │   │   │       # BE: subscriptions-be/subscription (financial-be)
    │   │   │   │   │   │   │   │   │       # POST /v1/subscriptions/plus
    │   │   │   │   │   │   │   │   └── page.tsx  # Plus membership overview
    │   │   │   │   │   │   │   │       # - Benefits
    │   │   │   │   │   │   │   │       # - Pricing
    │   │   │   │   │   │   │   │       # - Comparison
    │   │   │   │   │   │   │   │       # BE: subscriptions-be/plans
    │   │   │   │   │   │   │   │       # GET /v1/subscriptions/plans/plus
    │   │   │   │   │   │   │   ├── talent-cloud/
    │   │   │   │   │   │   │   │   ├── projects/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Talent Cloud exclusive projects
    │   │   │   │   │   │   │   │   │       # - High-value projects
    │   │   │   │   │   │   │   │   │       # - Direct invites
    │   │   │   │   │   │   │   │   │       # BE: jobs-be/job (with exclusive flag)
    │   │   │   │   │   │   │   │   │       # GET /v1/jobs/talent-cloud
    │   │   │   │   │   │   │   │   └── page.tsx  # Talent Cloud info
    │   │   │   │   │   │   │   │       # - Qualification
    │   │   │   │   │   │   │   │       # - Apply
    │   │   │   │   │   │   │   │       # - Status
    │   │   │   │   │   │   │   │       # BE: users-be/badges
    │   │   │   │   │   │   │   │       # GET /v1/programs/talent-cloud
    │   │   │   │   │   │   │   └── top-rated/
    │   │   │   │   │   │   │       ├── application/
    │   │   │   │   │   │   │       │   └── page.tsx  # Apply for Top Rated
    │   │   │   │   │   │   │       │       # - Eligibility check
    │   │   │   │   │   │   │       │       # - Submit application
    │   │   │   │   │   │   │       │       # BE: users-be/badges OR reviews-be/reputation
    │   │   │   │   │   │   │       │       # POST /v1/users/me/top-rated/apply
    │   │   │   │   │   │   │       └── page.tsx  # Top Rated program info
    │   │   │   │   │   │   │           # - Benefits
    │   │   │   │   │   │   │           # - Requirements
    │   │   │   │   │   │   │           # - Application status
    │   │   │   │   │   │   │           # BE: users-be/badges
    │   │   │   │   │   │   │           # GET /v1/programs/top-rated
    │   │   │   │   │   │   ├── spend-analytics/
    │   │   │   │   │   │   │   ├── by-category/
    │   │   │   │   │   │   │   │   └── page.tsx  # Spending by category
    │   │   │   │   │   │   │   │       # - Job category breakdown
    │   │   │   │   │   │   │   │       # - Trend analysis
    │   │   │   │   │   │   │   │       # BE: financial-be/analytics (if exists) or financial-be/invoice
    │   │   │   │   │   │   │   │       # GET /v1/analytics/spend/by-category
    │   │   │   │   │   │   │   ├── by-department/
    │   │   │   │   │   │   │   │   └── page.tsx  # Spending by department
    │   │   │   │   │   │   │   │       # - Cost center breakdown
    │   │   │   │   │   │   │   │       # - Budget vs actual
    │   │   │   │   │   │   │   │       # BE: financial-be/budget, financial-be/invoice
    │   │   │   │   │   │   │   │       # GET /v1/analytics/spend/by-department
    │   │   │   │   │   │   │   ├── by-vendor/
    │   │   │   │   │   │   │   │   └── page.tsx  # Spending by vendor
    │   │   │   │   │   │   │   │       # - Top vendors
    │   │   │   │   │   │   │   │       # - Vendor concentration risk
    │   │   │   │   │   │   │   │       # BE: financial-be/invoice, users-be/profile
    │   │   │   │   │   │   │   │       # GET /v1/analytics/spend/by-vendor
    │   │   │   │   │   │   │   ├── forecasting/
    │   │   │   │   │   │   │   │   └── page.tsx  # Spend forecasting
    │   │   │   │   │   │   │   │       # - Projected spend
    │   │   │   │   │   │   │   │       # - Budget burn rate
    │   │   │   │   │   │   │   │       # - Alerts for overages
    │   │   │   │   │   │   │   │       # BE: financial-be/forecast (if exists) or financial-be/analytics
    │   │   │   │   │   │   │   │       # GET /v1/analytics/spend/forecast
    │   │   │   │   │   │   │   └── page.tsx  # Spend analytics dashboard
    │   │   │   │   │   │   │       # - Overview metrics
    │   │   │   │   │   │   │       # - Charts & visualizations
    │   │   │   │   │   │   │       # BE: financial-be/analytics
    │   │   │   │   │   │   │       # GET /v1/analytics/spend
    │   │   │   │   │   │   ├── subscription/  # Subscription management
    │   │   │   │   │   │   │   ├── addons/
    │   │   │   │   │   │   │   │   └── [addonId]/
    │   │   │   │   │   │   │   │       ├── cancel/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Cancel addon
    │   │   │   │   │   │   │   │       │       # BE: subscriptions-be/addons
    │   │   │   │   │   │   │   │       │       # DELETE /v1/subscriptions/{sub_id}/addons/{addon_id}
    │   │   │   │   │   │   │   │       └── purchase/
    │   │   │   │   │   │   │   │           └── page.tsx  # Purchase addon
    │   │   │   │   │   │   │   │               # BE: subscriptions-be/addons
    │   │   │   │   │   │   │   │               # POST /v1/subscriptions/{sub_id}/addons
    │   │   │   │   │   │   │   ├── billing-history/
    │   │   │   │   │   │   │   │   ├── [invoiceId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Invoice detail
    │   │   │   │   │   │   │   │   │       # - Invoice details
    │   │   │   │   │   │   │   │   │       # - Download PDF
    │   │   │   │   │   │   │   │   │       # BE: subscriptions-be/invoices
    │   │   │   │   │   │   │   │   │       # GET /v1/invoices/{invoice_id}
    │   │   │   │   │   │   │   │   │       # GET /v1/invoices/{invoice_id}/pdf
    │   │   │   │   │   │   │   │   └── page.tsx  # Billing history
    │   │   │   │   │   │   │   │       # - Past invoices
    │   │   │   │   │   │   │   │       # - Payment history
    │   │   │   │   │   │   │   │       # - Download invoices
    │   │   │   │   │   │   │   │       # BE: subscriptions-be/invoices
    │   │   │   │   │   │   │   │       # GET /v1/subscriptions/{sub_id}/invoices
    │   │   │   │   │   │   │   ├── cancel/
    │   │   │   │   │   │   │   │   └── page.tsx  # Cancel subscription
    │   │   │   │   │   │   │   │       # - Cancellation reason
    │   │   │   │   │   │   │   │       # - Feedback
    │   │   │   │   │   │   │   │       # - Immediate vs. end of period
    │   │   │   │   │   │   │   │       # - Refund eligibility
    │   │   │   │   │   │   │   │       # - Data retention info
    │   │   │   │   │   │   │   │       # BE: subscriptions-be/subscriptions
    │   │   │   │   │   │   │   │       # POST /v1/subscriptions/{sub_id}/cancel
    │   │   │   │   │   │   │   │       # Publishes: SubscriptionCancelled event
    │   │   │   │   │   │   │   ├── connects/
    │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Connects usage history
    │   │   │   │   │   │   │   │   │       # - Connects spent (proposals)
    │   │   │   │   │   │   │   │   │       # - Connects added (purchases/plan)
    │   │   │   │   │   │   │   │   │       # - Balance over time
    │   │   │   │   │   │   │   │   │       # BE: subscriptions-be/connects
    │   │   │   │   │   │   │   │   │       # GET /v1/connects/transactions
    │   │   │   │   │   │   │   │   ├── purchase/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Purchase connects
    │   │   │   │   │   │   │   │   │       # - Select package
    │   │   │   │   │   │   │   │   │       # - Pricing options
    │   │   │   │   │   │   │   │   │       # - Bulk discounts
    │   │   │   │   │   │   │   │   │       # - Payment method
    │   │   │   │   │   │   │   │   │       # BE: subscriptions-be/connects
    │   │   │   │   │   │   │   │   │       # POST /v1/connects/purchase
    │   │   │   │   │   │   │   │   │       # Body: { package_id, quantity }
    │   │   │   │   │   │   │   │   │       # Publishes: ConnectsPurchased event
    │   │   │   │   │   │   │   │   └── page.tsx  # Connects overview
    │   │   │   │   │   │   │   │       # - Current balance
    │   │   │   │   │   │   │   │       # - Usage history
    │   │   │   │   │   │   │   │       # - Included in plan
    │   │   │   │   │   │   │   │       # - Purchase more
    │   │   │   │   │   │   │   │       # BE: subscriptions-be/connects
    │   │   │   │   │   │   │   │       # GET /v1/connects/balance
    │   │   │   │   │   │   │   │       # GET /v1/connects/usage
    │   │   │   │   │   │   │   ├── downgrade/
    │   │   │   │   │   │   │   │   └── page.tsx  # Downgrade plan
    │   │   │   │   │   │   │   │       # - Select new plan
    │   │   │   │   │   │   │   │       # - Effective date (end of billing period)
    │   │   │   │   │   │   │   │       # - Feature comparison
    │   │   │   │   │   │   │   │       # - Confirm downgrade
    │   │   │   │   │   │   │   │       # BE: subscriptions-be/subscriptions
    │   │   │   │   │   │   │   │       # POST /v1/subscriptions/downgrade
    │   │   │   │   │   │   │   │       # Publishes: SubscriptionDowngraded event
    │   │   │   │   │   │   │   ├── plans/
    │   │   │   │   │   │   │   │   ├── compare/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Plan comparison
    │   │   │   │   │   │   │   │   │       # - Side-by-side comparison
    │   │   │   │   │   │   │   │   │       # - Feature highlights
    │   │   │   │   │   │   │   │   │       # BE: subscriptions-be/plans
    │   │   │   │   │   │   │   │   │       # GET /v1/plans/compare
    │   │   │   │   │   │   │   │   └── page.tsx  # Available plans
    │   │   │   │   │   │   │   │       # - Plan comparison table
    │   │   │   │   │   │   │   │       # - Feature matrix
    │   │   │   │   │   │   │   │       # - Pricing tiers
    │   │   │   │   │   │   │   │       # - Billing periods (monthly, annual)
    │   │   │   │   │   │   │   │       # - Free trial info
    │   │   │   │   │   │   │   │       # BE: subscriptions-be/plans
    │   │   │   │   │   │   │   │       # GET /v1/plans
    │   │   │   │   │   │   │   ├── reactivate/
    │   │   │   │   │   │   │   │   └── page.tsx  # Reactivate subscription
    │   │   │   │   │   │   │   │       # - Select plan
    │   │   │   │   │   │   │   │       # - Payment method
    │   │   │   │   │   │   │   │       # BE: subscriptions-be/subscriptions
    │   │   │   │   │   │   │   │       # POST /v1/subscriptions/reactivate
    │   │   │   │   │   │   │   ├── trial/
    │   │   │   │   │   │   │   │   ├── convert/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Convert trial to paid
    │   │   │   │   │   │   │   │   │       # - Select plan
    │   │   │   │   │   │   │   │   │       # - Payment method
    │   │   │   │   │   │   │   │   │       # - Apply promotion code
    │   │   │   │   │   │   │   │   │       # BE: subscriptions-be/trials
    │   │   │   │   │   │   │   │   │       # POST /v1/trials/{trial_id}/convert
    │   │   │   │   │   │   │   │   └── page.tsx  # Trial status
    │   │   │   │   │   │   │   │       # - Trial end date
    │   │   │   │   │   │   │   │       # - Days remaining
    │   │   │   │   │   │   │   │       # - Trial features
    │   │   │   │   │   │   │   │       # - Upgrade prompt
    │   │   │   │   │   │   │   │       # BE: subscriptions-be/trials
    │   │   │   │   │   │   │   │       # GET /v1/trials/current
    │   │   │   │   │   │   │   ├── upgrade/
    │   │   │   │   │   │   │   │   ├── confirm/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Confirm upgrade
    │   │   │   │   │   │   │   │   │       # - Upgrade summary
    │   │   │   │   │   │   │   │   │       # - Payment processing
    │   │   │   │   │   │   │   │   │       # BE: subscriptions-be/subscriptions
    │   │   │   │   │   │   │   │   │       # POST /v1/subscriptions/{sub_id}/confirm-upgrade
    │   │   │   │   │   │   │   │   └── page.tsx  # Upgrade plan
    │   │   │   │   │   │   │   │       # - Select new plan
    │   │   │   │   │   │   │   │       # - Billing period
    │   │   │   │   │   │   │   │       # - Proration calculation
    │   │   │   │   │   │   │   │       # - Payment method
    │   │   │   │   │   │   │   │       # - Confirm upgrade
    │   │   │   │   │   │   │   │       # BE: subscriptions-be/subscriptions
    │   │   │   │   │   │   │   │       # POST /v1/subscriptions/upgrade
    │   │   │   │   │   │   │   │       # Body: { plan_id, billing_period }
    │   │   │   │   │   │   │   │       # Publishes: SubscriptionUpgraded event
    │   │   │   │   │   │   │   ├── usage/
    │   │   │   │   │   │   │   │   └── page.tsx  # Usage statistics
    │   │   │   │   │   │   │   │       # - Jobs posted
    │   │   │   │   │   │   │   │       # - Proposals submitted
    │   │   │   │   │   │   │   │       # - Storage used
    │   │   │   │   │   │   │   │       # - API calls
    │   │   │   │   │   │   │   │       # - Usage vs. limits
    │   │   │   │   │   │   │   │       # BE: subscriptions-be/usage
    │   │   │   │   │   │   │   │       # GET /v1/subscriptions/usage
    │   │   │   │   │   │   │   └── page.tsx  # Subscription overview
    │   │   │   │   │   │   │       # - Current plan details
    │   │   │   │   │   │   │       # - Usage stats
    │   │   │   │   │   │   │       # - Next billing date
    │   │   │   │   │   │   │       # - Upgrade/downgrade options
    │   │   │   │   │   │   │       # - Connects balance
    │   │   │   │   │   │   │       # BE: subscriptions-be/subscriptions
    │   │   │   │   │   │   │       # GET /v1/subscriptions/current
    │   │   │   │   │   │   │       # BE: subscriptions-be/entitlements
    │   │   │   │   │   │   │       # GET /v1/entitlements
    │   │   │   │   │   │   │       # BE: subscriptions-be/connects
    │   │   │   │   │   │   │       # GET /v1/connects/balance
    │   │   │   │   │   │   ├── subscriptions/
    │   │   │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   │   │   └── page.tsx  # Subscription billing details
    │   │   │   │   │   │   │   │       # - Current charges
    │   │   │   │   │   │   │   │       # - Next billing date
    │   │   │   │   │   │   │   │       # - Payment history
    │   │   │   │   │   │   │   │       # BE: financial-be/subscription
    │   │   │   │   │   │   │   │       # GET /v1/subscriptions/billing
    │   │   │   │   │   │   │   ├── cancel/
    │   │   │   │   │   │   │   │   └── page.tsx  # Cancel subscription
    │   │   │   │   │   │   │   │       # - Cancellation reasons
    │   │   │   │   │   │   │   │       # - Retention offers
    │   │   │   │   │   │   │   │       # - Confirm cancellation
    │   │   │   │   │   │   │   │       # BE: financial-be/subscription
    │   │   │   │   │   │   │   │       # POST /v1/subscriptions/cancel
    │   │   │   │   │   │   │   ├── change-plan/
    │   │   │   │   │   │   │   │   └── page.tsx  # Change subscription plan
    │   │   │   │   │   │   │   │       # - Plan comparison
    │   │   │   │   │   │   │   │       # - Proration calculation
    │   │   │   │   │   │   │   │       # - Confirm change
    │   │   │   │   │   │   │   │       # BE: financial-be/subscription
    │   │   │   │   │   │   │   │       # POST /v1/subscriptions/change-plan
    │   │   │   │   │   │   │   ├── features/
    │   │   │   │   │   │   │   │   └── page.tsx  # Subscription features
    │   │   │   │   │   │   │   │       # - Active features
    │   │   │   │   │   │   │   │       # - Feature limits
    │   │   │   │   │   │   │   │       # - Usage tracking
    │   │   │   │   │   │   │   │       # BE: financial-be/subscription
    │   │   │   │   │   │   │   │       # GET /v1/subscriptions/features
    │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   └── page.tsx  # Subscription history
    │   │   │   │   │   │   │   │       # - Past subscriptions
    │   │   │   │   │   │   │   │       # - Billing history
    │   │   │   │   │   │   │   │       # BE: financial-be/subscription
    │   │   │   │   │   │   │   │       # GET /v1/subscriptions/history
    │   │   │   │   │   │   │   ├── pause/
    │   │   │   │   │   │   │   │   └── page.tsx  # Pause subscription
    │   │   │   │   │   │   │   │       # - Pause duration
    │   │   │   │   │   │   │   │       # - Resume date
    │   │   │   │   │   │   │   │       # BE: financial-be/subscription
    │   │   │   │   │   │   │   │       # POST /v1/subscriptions/pause
    │   │   │   │   │   │   │   ├── reactivate/
    │   │   │   │   │   │   │   │   └── page.tsx  # Reactivate subscription
    │   │   │   │   │   │   │   │       # - Choose plan
    │   │   │   │   │   │   │   │       # - Payment method
    │   │   │   │   │   │   │   │       # BE: financial-be/subscription
    │   │   │   │   │   │   │   │       # POST /v1/subscriptions/reactivate
    │   │   │   │   │   │   │   ├── upgrade/
    │   │   │   │   │   │   │   │   └── page.tsx  # Upgrade subscription
    │   │   │   │   │   │   │   └── usage/
    │   │   │   │   │   │   │       └── page.tsx  # Subscription usage metrics
    │   │   │   │   │   │   │           # - Connects used
    │   │   │   │   │   │   │           # - Posts remaining
    │   │   │   │   │   │   │           # - Feature usage
    │   │   │   │   │   │   │           # BE: financial-be/subscription
    │   │   │   │   │   │   │           # GET /v1/subscriptions/usage
    │   │   │   │   │   │   ├── talent/
    │   │   │   │   │   │   │   ├── browse/
    │   │   │   │   │   │   │   │   └── page.tsx  # Browse talent
    │   │   │   │   │   │   │   │       # - Search freelancers
    │   │   │   │   │   │   │   │       # - Filters (skills, rate, location)
    │   │   │   │   │   │   │   │       # - Save to shortlist
    │   │   │   │   │   │   │   │       # BE: search-be/query, users-be/profile
    │   │   │   │   │   │   │   │       # POST /v1/search/freelancers
    │   │   │   │   │   │   │   │       # GET /v1/search/freelancers?filters=...
    │   │   │   │   │   │   │   ├── recommendations/
    │   │   │   │   │   │   │   │   └── page.tsx  # AI-recommended talent for jobs
    │   │   │   │   │   │   │   │       # BE: search-be/recommendation
    │   │   │   │   │   │   │   │       # GET /v1/recommendations/talent?job_id={job_id}
    │   │   │   │   │   │   │   ├── saved/
    │   │   │   │   │   │   │   │   └── page.tsx  # Saved talent profiles
    │   │   │   │   │   │   │   │       # BE: users-be/profile
    │   │   │   │   │   │   │   │       # GET /v1/users/me/saved-profiles
    │   │   │   │   │   │   │   └── shortlists/
    │   │   │   │   │   │   │       ├── [shortlistId]/
    │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit shortlist
    │   │   │   │   │   │   │       │   │       # BE: jobs-be/shortlist
    │   │   │   │   │   │   │       │   │       # PUT /v1/jobs/{job_id}/shortlists/{shortlist_id}
    │   │   │   │   │   │   │       │   └── page.tsx  # Shortlist details
    │   │   │   │   │   │   │       │       # - View candidates
    │   │   │   │   │   │   │       │       # - Send invitations
    │   │   │   │   │   │   │       │       # - Compare profiles
    │   │   │   │   │   │   │       │       # BE: jobs-be/shortlist
    │   │   │   │   │   │   │       │       # GET /v1/jobs/{job_id}/shortlists/{shortlist_id}
    │   │   │   │   │   │   │       ├── new/
    │   │   │   │   │   │   │       │   └── page.tsx  # Create shortlist
    │   │   │   │   │   │   │       │       # BE: jobs-be/shortlist
    │   │   │   │   │   │   │       │       # POST /v1/jobs/{job_id}/shortlists
    │   │   │   │   │   │   │       └── page.tsx  # Shortlists overview
    │   │   │   │   │   │   │           # BE: jobs-be/shortlist
    │   │   │   │   │   │   │           # GET /v1/jobs/{job_id}/shortlists
    │   │   │   │   │   │   ├── teams/
    │   │   │   │   │   │   │   ├── [teamId]/
    │   │   │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Team compliance dashboard
    │   │   │   │   │   │   │   │   │       # - Document status
    │   │   │   │   │   │   │   │   │       # - Training completion
    │   │   │   │   │   │   │   │   │       # - Policy acknowledgments
    │   │   │   │   │   │   │   │   │       # BE: users-be/team, admin-be/business_verification
    │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/compliance
    │   │   │   │   │   │   │   │   ├── hierarchy/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Team hierarchy management
    │   │   │   │   │   │   │   │   │       # - Organizational chart
    │   │   │   │   │   │   │   │   │       # - Reporting structure
    │   │   │   │   │   │   │   │   │       # - Role relationships
    │   │   │   │   │   │   │   │   │       # BE: users-be/team
    │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/hierarchy
    │   │   │   │   │   │   │   │   ├── performance/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Team performance metrics
    │   │   │   │   │   │   │   │   │       # - Productivity stats
    │   │   │   │   │   │   │   │   │       # - Quality metrics
    │   │   │   │   │   │   │   │   │       # - Member contributions
    │   │   │   │   │   │   │   │   │       # BE: users-be/team, contracts-be/analytics
    │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/performance
    │   │   │   │   │   │   │   │   └── spending-controls/
    │   │   │   │   │   │   │   │       └── page.tsx  # Team spending controls
    │   │   │   │   │   │   │   │           # - Approval workflows
    │   │   │   │   │   │   │   │           # - Spending limits
    │   │   │   │   │   │   │   │           # - Auto-approval rules
    │   │   │   │   │   │   │   │           # BE: users-be/team, financial-be/budget
    │   │   │   │   │   │   │   │           # GET /v1/teams/{team_id}/spending-controls
    │   │   │   │   │   │   │   │           # PUT /v1/teams/{team_id}/spending-controls
    │   │   │   │   │   │   │   └── integrations/
    │   │   │   │   │   │   │       ├── [integrationId]/
    │   │   │   │   │   │   │       │   ├── configure/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Configure integration
    │   │   │   │   │   │   │       │   │       # BE: users-be/integration
    │   │   │   │   │   │   │       │   │       # PUT /v1/teams/integrations/{integration_id}
    │   │   │   │   │   │   │       │   ├── logs/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Integration logs
    │   │   │   │   │   │   │       │   │       # BE: users-be/integration, utility-be/audit
    │   │   │   │   │   │   │       │   │       # GET /v1/teams/integrations/{integration_id}/logs
    │   │   │   │   │   │   │       │   └── page.tsx  # Integration details
    │   │   │   │   │   │   │       │       # BE: users-be/integration
    │   │   │   │   │   │   │       │       # GET /v1/teams/integrations/{integration_id}
    │   │   │   │   │   │   │       ├── available/
    │   │   │   │   │   │   │       │   └── page.tsx  # Available integrations
    │   │   │   │   │   │   │       │       # - Slack, JIRA, etc.
    │   │   │   │   │   │   │       │       # - Feature descriptions
    │   │   │   │   │   │   │       │       # BE: users-be/integration
    │   │   │   │   │   │   │       │       # GET /v1/teams/integrations/available
    │   │   │   │   │   │   │       └── page.tsx  # Active integrations list
    │   │   │   │   │   │   │           # BE: users-be/integration
    │   │   │   │   │   │   │           # GET /v1/teams/integrations
    │   │   │   │   │   │   ├── timesheets/
    │   │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   │   ├── [timesheetId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit timesheet
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/timesheet
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/contracts/{contract_id}/timesheets/{timesheet_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Timesheet details
    │   │   │   │   │   │   │   │   │       # - Hours breakdown
    │   │   │   │   │   │   │   │   │       # - Approval status
    │   │   │   │   │   │   │   │   │       # - Dispute options
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/timesheet
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/timesheets/{timesheet_id}
    │   │   │   │   │   │   │   │   ├── new/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create timesheet
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/timesheet
    │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/timesheets
    │   │   │   │   │   │   │   │   └── page.tsx  # Contract timesheets list
    │   │   │   │   │   │   │   │       # BE: contracts-be/timesheet
    │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/timesheets
    │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   └── page.tsx  # Timesheets pending approval (client)
    │   │   │   │   │   │   │   │       # BE: contracts-be/timesheet
    │   │   │   │   │   │   │   │       # GET /v1/timesheets/pending-approval
    │   │   │   │   │   │   │   └── page.tsx  # All timesheets overview
    │   │   │   │   │   │   │       # BE: contracts-be/timesheet
    │   │   │   │   │   │   │       # GET /v1/timesheets
    │   │   │   │   │   │   ├── vendor-management/
    │   │   │   │   │   │   │   ├── blacklist/
    │   │   │   │   │   │   │   │   ├── [userId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Blacklist entry detail
    │   │   │   │   │   │   │   │   │       # - Reason for blacklist
    │   │   │   │   │   │   │   │   │       # - Remove option
    │   │   │   │   │   │   │   │   │       # BE: users-be/org (blacklist subdomain)
    │   │   │   │   │   │   │   │   │       # GET /v1/vendors/blacklist/{user_id}
    │   │   │   │   │   │   │   │   │       # DELETE /v1/vendors/blacklist/{user_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Blacklisted vendors
    │   │   │   │   │   │   │   │       # BE: users-be/org
    │   │   │   │   │   │   │   │       # GET /v1/vendors/blacklist
    │   │   │   │   │   │   │   ├── compliance-docs/
    │   │   │   │   │   │   │   │   ├── [vendorId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Vendor compliance documents
    │   │   │   │   │   │   │   │   │       # - W-9/W-8BEN forms
    │   │   │   │   │   │   │   │   │       # - Insurance certificates
    │   │   │   │   │   │   │   │   │       # - Background checks
    │   │   │   │   │   │   │   │   │       # BE: admin-be/business_verification, storage-be/asset
    │   │   │   │   │   │   │   │   │       # GET /v1/vendors/{vendor_id}/compliance-docs
    │   │   │   │   │   │   │   │   └── page.tsx  # Compliance tracking
    │   │   │   │   │   │   │   │       # - Expiring documents
    │   │   │   │   │   │   │   │       # - Missing documents
    │   │   │   │   │   │   │   │       # BE: admin-be/business_verification
    │   │   │   │   │   │   │   │       # GET /v1/vendors/compliance-status
    │   │   │   │   │   │   │   └── preferred/
    │   │   │   │   │   │   │       ├── [vendorId]/
    │   │   │   │   │   │   │       │   ├── history/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Work history with vendor
    │   │   │   │   │   │   │       │   │       # - Past contracts
    │   │   │   │   │   │   │       │   │       # - Total spend
    │   │   │   │   │   │   │       │   │       # - Reviews given
    │   │   │   │   │   │   │       │   │       # BE: contracts-be/contract, financial-be/invoice
    │   │   │   │   │   │   │       │   │       # GET /v1/vendors/{vendor_id}/history
    │   │   │   │   │   │   │       │   ├── performance/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Vendor performance metrics
    │   │   │   │   │   │   │       │   │       # - Success rate
    │   │   │   │   │   │   │       │   │       # - Average delivery time
    │   │   │   │   │   │   │       │   │       # - Quality scores
    │   │   │   │   │   │   │       │   │       # BE: users-be/org (vendor subdomain), contracts-be/contract
    │   │   │   │   │   │   │       │   │       # GET /v1/vendors/{vendor_id}/performance
    │   │   │   │   │   │   │       │   └── page.tsx  # Vendor detail
    │   │   │   │   │   │   │       │       # BE: users-be/org (vendor subdomain)
    │   │   │   │   │   │   │       │       # GET /v1/vendors/{vendor_id}
    │   │   │   │   │   │   │       └── page.tsx  # Preferred vendors list
    │   │   │   │   │   │   │           # - Star ratings
    │   │   │   │   │   │   │           # - Quick invite
    │   │   │   │   │   │   │           # BE: users-be/org
    │   │   │   │   │   │   │           # GET /v1/vendors/preferred
    │   │   │   │   │   │   ├── work-diary/
    │   │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   │   ├── calendar/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Calendar view of work diary
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/work_diary
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/work-diary/calendar
    │   │   │   │   │   │   │   │   ├── screenshots/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Screenshots management
    │   │   │   │   │   │   │   │   │       # - View all screenshots
    │   │   │   │   │   │   │   │   │       # - Delete sensitive ones
    │   │   │   │   │   │   │   │   │       # - Privacy settings
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/work_diary, storage-be/asset
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/work-diary/screenshots
    │   │   │   │   │   │   │   │   └── page.tsx  # Work diary detail
    │   │   │   │   │   │   │   │       # BE: contracts-be/work_diary
    │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/work-diary
    │   │   │   │   │   │   │   └── page.tsx  # Work diary overview (all contracts)
    │   │   │   │   │   │   │       # BE: contracts-be/work_diary
    │   │   │   │   │   │   │       # GET /v1/work-diary
    │   │   │   │   │   │   └── layout.tsx  # Dashboard layout
    │   │   │   │   │   │       # - Authenticated header with user menu
    │   │   │   │   │   │       # - Sidebar navigation (responsive)
    │   │   │   │   │   │       # - Notification bell with unread count
    │   │   │   │   │   │       # - Messages indicator
    │   │   │   │   │   │       # - Breadcrumbs
    │   │   │   │   │   │       # - Footer
    │   │   │   │   │   │       # BE: communications-be
    │   │   │   │   │   │       # GET /v1/notifications/unread-count
    │   │   │   │   │   │       # WebSocket: ws://communications-be/v1/realtime
    │   │   │   │   │   ├── (delivery)/
    │   │   │   │   │   │   └── deliverables/
    │   │   │   │   │   │       └── [contractId]/
    │   │   │   │   │   │           └── page.tsx  # Deliverables & versions
    │   │   │   │   │   │               # - Upload, request changes
    │   │   │   │   │   │               # BE: contracts-be/deliverable | storage-be/asset
    │   │   │   │   │   │               # GET/POST /v1/contracts/{id}/deliverables
    │   │   │   │   │   ├── (freelancer)/
    │   │   │   │   │   │   ├── portfolio/
    │   │   │   │   │   │   │   └── manage/
    │   │   │   │   │   │   │       └── page.tsx  # Manage portfolio
    │   │   │   │   │   │   │           # - Add/edit items, media uploads
    │   │   │   │   │   │   │           # BE: users-be/portfolio | storage-be/asset
    │   │   │   │   │   │   │           # GET/POST/PATCH /v1/portfolio
    │   │   │   │   │   │   └── proposals/
    │   │   │   │   │   │       └── new/
    │   │   │   │   │   │           └── page.tsx  # New proposal (freelancer)
    │   │   │   │   │   │               # - Compose, attach files
    │   │   │   │   │   │               # BE: proposals-be/proposal | storage-be/asset
    │   │   │   │   │   │               # POST /v1/proposals
    │   │   │   │   │   ├── (onboarding)/  # Onboarding flow (post-registration)
    │   │   │   │   │   │   ├── client/  # Client onboarding
    │   │   │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   │   │   └── page.tsx  # Billing setup
    │   │   │   │   │   │   │   │       # - Billing address
    │   │   │   │   │   │   │   │       # - Tax information
    │   │   │   │   │   │   │   │       # - VAT/GST number
    │   │   │   │   │   │   │   │       # - Payment method
    │   │   │   │   │   │   │   │       # BE: financial-be/billing_profile
    │   │   │   │   │   │   │   │       # POST /v1/billing-profiles
    │   │   │   │   │   │   │   │       # BE: financial-be/payment_method
    │   │   │   │   │   │   │   │       # POST /v1/payment-methods
    │   │   │   │   │   │   │   ├── company/
    │   │   │   │   │   │   │   │   └── page.tsx  # Company info
    │   │   │   │   │   │   │   │       # - Company name
    │   │   │   │   │   │   │   │       # - Industry
    │   │   │   │   │   │   │   │       # - Company size
    │   │   │   │   │   │   │   │       # - Website
    │   │   │   │   │   │   │   │       # - Logo upload
    │   │   │   │   │   │   │   │       # BE: users-be/organization
    │   │   │   │   │   │   │   │       # POST /v1/organizations
    │   │   │   │   │   │   │   │       # BE: storage-be/uploads
    │   │   │   │   │   │   │   ├── preferences/
    │   │   │   │   │   │   │   │   └── page.tsx  # Hiring preferences
    │   │   │   │   │   │   │   │       # - Typical project types
    │   │   │   │   │   │   │   │       # - Budget ranges
    │   │   │   │   │   │   │   │       # - Notification settings
    │   │   │   │   │   │   │   │       # BE: users-be/preferences
    │   │   │   │   │   │   │   │       # POST /v1/users/{id}/preferences
    │   │   │   │   │   │   │   │       # Publishes: ClientProfileCompleted event
    │   │   │   │   │   │   │   ├── team/
    │   │   │   │   │   │   │   │   └── page.tsx  # Team setup (optional)
    │   │   │   │   │   │   │   │       # - Invite team members
    │   │   │   │   │   │   │   │       # - Assign roles
    │   │   │   │   │   │   │   │       # BE: users-be/team
    │   │   │   │   │   │   │   │       # POST /v1/organizations/{id}/members
    │   │   │   │   │   │   │   │       # BE: communications-be
    │   │   │   │   │   │   │   │       # POST /v1/invitations/team
    │   │   │   │   │   │   │   └── verification/
    │   │   │   │   │   │   │       └── page.tsx  # Business verification
    │   │   │   │   │   │   │           # - Upload business documents
    │   │   │   │   │   │   │           # - Company registration
    │   │   │   │   │   │   │           # - Tax documents
    │   │   │   │   │   │   │           # BE: admin-be/business_verification
    │   │   │   │   │   │   │           # POST /v1/admin/business-verification
    │   │   │   │   │   │   │           # BE: storage-be/uploads
    │   │   │   │   │   │   ├── freelancer/  # Freelancer onboarding
    │   │   │   │   │   │   │   ├── experience/
    │   │   │   │   │   │   │   │   └── page.tsx  # Work experience
    │   │   │   │   │   │   │   │       # - Add previous positions
    │   │   │   │   │   │   │   │       # - Company, role, dates, description
    │   │   │   │   │   │   │   │       # BE: users-be/experience
    │   │   │   │   │   │   │   │       # POST /v1/users/{id}/experience
    │   │   │   │   │   │   │   ├── portfolio/
    │   │   │   │   │   │   │   │   └── page.tsx  # Portfolio items
    │   │   │   │   │   │   │   │       # - Upload work samples
    │   │   │   │   │   │   │   │       # - Project descriptions
    │   │   │   │   │   │   │   │       # - Links to live work
    │   │   │   │   │   │   │   │       # BE: users-be/portfolio
    │   │   │   │   │   │   │   │       # POST /v1/users/{id}/portfolio
    │   │   │   │   │   │   │   │       # BE: storage-be/uploads
    │   │   │   │   │   │   │   │       # POST /v1/storage/upload (signed URL)
    │   │   │   │   │   │   │   ├── preferences/
    │   │   │   │   │   │   │   │   └── page.tsx  # Job preferences
    │   │   │   │   │   │   │   │       # - Job categories
    │   │   │   │   │   │   │   │       # - Work availability
    │   │   │   │   │   │   │   │       # - Preferred job types
    │   │   │   │   │   │   │   │       # - Notification settings
    │   │   │   │   │   │   │   │       # BE: users-be/preferences
    │   │   │   │   │   │   │   │       # POST /v1/users/{id}/preferences
    │   │   │   │   │   │   │   │       # Publishes: FreelancerProfileCompleted event
    │   │   │   │   │   │   │   ├── profile/
    │   │   │   │   │   │   │   │   └── page.tsx  # Basic profile
    │   │   │   │   │   │   │   │       # - Professional title
    │   │   │   │   │   │   │   │       # - Bio
    │   │   │   │   │   │   │   │       # - Location
    │   │   │   │   │   │   │   │       # - Profile photo
    │   │   │   │   │   │   │   │       # BE: users-be/profile
    │   │   │   │   │   │   │   │       # PATCH /v1/users/{id}/profile
    │   │   │   │   │   │   │   ├── rates/
    │   │   │   │   │   │   │   │   └── page.tsx  # Rate setting
    │   │   │   │   │   │   │   │       # - Hourly rate
    │   │   │   │   │   │   │   │       # - Preferred project budget range
    │   │   │   │   │   │   │   │       # - Currency
    │   │   │   │   │   │   │   │       # BE: users-be/freelancer
    │   │   │   │   │   │   │   │       # PATCH /v1/users/{id}/rates
    │   │   │   │   │   │   │   └── skills/
    │   │   │   │   │   │   │       └── page.tsx  # Skills selection
    │   │   │   │   │   │   │           # - Skill search & autocomplete
    │   │   │   │   │   │   │           # - Skill level (Beginner/Intermediate/Expert)
    │   │   │   │   │   │   │           # - Primary skills (minimum 3)
    │   │   │   │   │   │   │           # BE: users-be/capabilities
    │   │   │   │   │   │   │           # POST /v1/users/{id}/skills
    │   │   │   │   │   │   │           # Body: { skills: [{ skill_id, level }] }
    │   │   │   │   │   │   ├── welcome/
    │   │   │   │   │   │   │   └── page.tsx  # Welcome message
    │   │   │   │   │   │   │       # - User type confirmation
    │   │   │   │   │   │   │       # - Next steps preview
    │   │   │   │   │   │   │       # BE: users-be/user
    │   │   │   │   │   │   │       # GET /v1/users/me
    │   │   │   │   │   │   └── layout.tsx  # Onboarding layout
    │   │   │   │   │   │       # - Progress indicator
    │   │   │   │   │   │       # - Skip options
    │   │   │   │   │   │       # - Help sidebar
    │   │   │   │   │   ├── (public)/  # Public pages (no auth required)
    │   │   │   │   │   │   ├── about/
    │   │   │   │   │   │   │   ├── leadership/
    │   │   │   │   │   │   │   │   └── page.tsx  # Leadership team
    │   │   │   │   │   │   │   │       # BE: None (static)
    │   │   │   │   │   │   │   ├── press/
    │   │   │   │   │   │   │   │   └── page.tsx  # Press releases
    │   │   │   │   │   │   │   │       # BE: None (static or CMS)
    │   │   │   │   │   │   │   └── page.tsx  # About page
    │   │   │   │   │   │   │       # - Company story
    │   │   │   │   │   │   │       # - Team showcase
    │   │   │   │   │   │   │       # - Mission & values
    │   │   │   │   │   │   │       # BE: None (static)
    │   │   │   │   │   │   ├── blog/
    │   │   │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Blog post detail
    │   │   │   │   │   │   │   ├── category/
    │   │   │   │   │   │   │   │   └── [category]/
    │   │   │   │   │   │   │   │       └── page.tsx  # Category listing
    │   │   │   │   │   │   │   │           # BE: CMS or separate service
    │   │   │   │   │   │   │   └── page.tsx  # Blog listing
    │   │   │   │   │   │   ├── careers/
    │   │   │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Individual job posting
    │   │   │   │   │   │   │   │       # BE: None (static or CMS)
    │   │   │   │   │   │   │   └── page.tsx  # Careers overview
    │   │   │   │   │   │   ├── case-studies/
    │   │   │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Case study detail
    │   │   │   │   │   │   │   │       # BE: None (static or CMS)
    │   │   │   │   │   │   │   └── page.tsx  # Case studies list
    │   │   │   │   │   │   │       # BE: None (static or CMS)
    │   │   │   │   │   │   ├── contact/
    │   │   │   │   │   │   │   └── page.tsx  # Contact form
    │   │   │   │   │   │   │       # - Support inquiries
    │   │   │   │   │   │   │       # - Partnership requests
    │   │   │   │   │   │   │       # - Media contacts
    │   │   │   │   │   │   │       # BE: communications-be/messages
    │   │   │   │   │   │   │       # POST /v1/contact
    │   │   │   │   │   │   │       # Body: { name, email, subject, message, type }
    │   │   │   │   │   │   ├── developers/
    │   │   │   │   │   │   │   ├── api/
    │   │   │   │   │   │   │   │   ├── authentication/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth & API keys
    │   │   │   │   │   │   │   │   ├── endpoints/
    │   │   │   │   │   │   │   │   │   ├── [category]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Endpoint category
    │   │   │   │   │   │   │   │   │   └── page.tsx  # All endpoints
    │   │   │   │   │   │   │   │   ├── getting-started/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # API getting started
    │   │   │   │   │   │   │   │   ├── rate-limits/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Rate limiting info
    │   │   │   │   │   │   │   │   ├── webhooks/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Webhook documentation
    │   │   │   │   │   │   │   │   └── page.tsx  # API overview
    │   │   │   │   │   │   │   │       # API documentation
    │   │   │   │   │   │   │   │       # BE: None (static)
    │   │   │   │   │   │   │   ├── changelog/
    │   │   │   │   │   │   │   │   └── page.tsx  # API changelog
    │   │   │   │   │   │   │   │       # BE: None (static)
    │   │   │   │   │   │   │   ├── community/
    │   │   │   │   │   │   │   │   ├── forum/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Developer forum
    │   │   │   │   │   │   │   │   └── github/
    │   │   │   │   │   │   │   │       └── page.tsx  # GitHub repos
    │   │   │   │   │   │   │   ├── plugins/
    │   │   │   │   │   │   │   │   ├── shopify/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Shopify app
    │   │   │   │   │   │   │   │   └── wordpress/
    │   │   │   │   │   │   │   │       └── page.tsx  # WordPress plugin
    │   │   │   │   │   │   │   ├── sdks/
    │   │   │   │   │   │   │   │   ├── javascript/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # JavaScript SDK
    │   │   │   │   │   │   │   │   ├── python/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Python SDK
    │   │   │   │   │   │   │   │   ├── ruby/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Ruby SDK
    │   │   │   │   │   │   │   │   └── page.tsx  # All SDKs
    │   │   │   │   │   │   │   └── page.tsx  # Developer portal
    │   │   │   │   │   │   │       # BE: None (static)
    │   │   │   │   │   │   ├── enterprise/
    │   │   │   │   │   │   │   ├── contact/
    │   │   │   │   │   │   │   │   └── page.tsx  # Enterprise contact
    │   │   │   │   │   │   │   │       # BE: communications-be/messages
    │   │   │   │   │   │   │   │       # POST /v1/contact/enterprise
    │   │   │   │   │   │   │   ├── demo/
    │   │   │   │   │   │   │   │   └── page.tsx  # Request demo
    │   │   │   │   │   │   │   │       # BE: communications-be/messages
    │   │   │   │   │   │   │   │       # POST /v1/contact/demo
    │   │   │   │   │   │   │   └── page.tsx  # Enterprise solutions
    │   │   │   │   │   │   │       # BE: None (static)
    │   │   │   │   │   │   ├── help/
    │   │   │   │   │   │   │   ├── [category]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Help category
    │   │   │   │   │   │   │   ├── article/
    │   │   │   │   │   │   │   │   └── [slug]/
    │   │   │   │   │   │   │   │       └── page.tsx  # Help article
    │   │   │   │   │   │   │   │           # BE: communications-be/knowledge-base
    │   │   │   │   │   │   │   │           # GET /v1/kb/articles
    │   │   │   │   │   │   │   │           # GET /v1/kb/articles/{slug}
    │   │   │   │   │   │   │   └── page.tsx  # Help center home
    │   │   │   │   │   │   ├── how-it-works/
    │   │   │   │   │   │   │   ├── clients/
    │   │   │   │   │   │   │   │   └── page.tsx  # For clients
    │   │   │   │   │   │   │   │       # - Posting jobs
    │   │   │   │   │   │   │   │       # - Hiring process
    │   │   │   │   │   │   │   │       # - Payment & escrow
    │   │   │   │   │   │   │   │       # BE: None (static)
    │   │   │   │   │   │   │   ├── freelancers/
    │   │   │   │   │   │   │   │   └── page.tsx  # For freelancers
    │   │   │   │   │   │   │   │       # - Getting started
    │   │   │   │   │   │   │   │       # - Finding jobs
    │   │   │   │   │   │   │   │       # - Earning money
    │   │   │   │   │   │   │   │       # BE: None (static)
    │   │   │   │   │   │   │   └── page.tsx  # How it works overview
    │   │   │   │   │   │   ├── legal/
    │   │   │   │   │   │   │   ├── accessibility/
    │   │   │   │   │   │   │   │   └── page.tsx  # Accessibility statement
    │   │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   │   ├── aml/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # AML policy
    │   │   │   │   │   │   │   │   ├── kyc/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # KYC policy
    │   │   │   │   │   │   │   │   └── page.tsx  # Compliance information
    │   │   │   │   │   │   │   │       # BE: None (static/versioned)
    │   │   │   │   │   │   │   │       # Compliance overview
    │   │   │   │   │   │   │   ├── cookies/
    │   │   │   │   │   │   │   │   └── page.tsx  # Cookie Policy
    │   │   │   │   │   │   │   │       # Cookie policy
    │   │   │   │   │   │   │   │       # BE: None (static)
    │   │   │   │   │   │   │   ├── dmca/
    │   │   │   │   │   │   │   │   └── page.tsx  # DMCA policy
    │   │   │   │   │   │   │   │       # BE: None (static)
    │   │   │   │   │   │   │   ├── licenses/
    │   │   │   │   │   │   │   │   └── page.tsx  # Open source licenses
    │   │   │   │   │   │   │   ├── privacy/
    │   │   │   │   │   │   │   │   ├── ccpa/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # CCPA information
    │   │   │   │   │   │   │   │   ├── cookies/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Cookie Policy
    │   │   │   │   │   │   │   │   ├── gdpr/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # GDPR information
    │   │   │   │   │   │   │   │   ├── policy/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Privacy Policy
    │   │   │   │   │   │   │   │   └── page.tsx  # Privacy Policy
    │   │   │   │   │   │   │   │       # Privacy policy
    │   │   │   │   │   │   │   │       # BE: None (static)
    │   │   │   │   │   │   │   └── terms/
    │   │   │   │   │   │   │       ├── client/
    │   │   │   │   │   │   │       │   └── page.tsx  # Client Agreement
    │   │   │   │   │   │   │       ├── freelancer/
    │   │   │   │   │   │   │       │   └── page.tsx  # Freelancer Agreement
    │   │   │   │   │   │   │       ├── service/
    │   │   │   │   │   │   │       │   └── page.tsx  # Terms of Service
    │   │   │   │   │   │   │       └── page.tsx  # Terms of Service
    │   │   │   │   │   │   │           # Legal terms overview
    │   │   │   │   │   │   │           # Terms of service
    │   │   │   │   │   │   │           # BE: None (static)
    │   │   │   │   │   │   ├── partners/
    │   │   │   │   │   │   │   ├── become-partner/
    │   │   │   │   │   │   │   │   └── page.tsx  # Partner application
    │   │   │   │   │   │   │   │       # BE: communications-be/messages
    │   │   │   │   │   │   │   │       # POST /v1/partners/apply
    │   │   │   │   │   │   │   ├── directory/
    │   │   │   │   │   │   │   │   └── page.tsx  # Partners directory
    │   │   │   │   │   │   │   │       # BE: None (static or CMS)
    │   │   │   │   │   │   │   └── page.tsx  # Partners program
    │   │   │   │   │   │   │       # BE: None (static)
    │   │   │   │   │   │   ├── pricing/
    │   │   │   │   │   │   │   ├── compare/
    │   │   │   │   │   │   │   │   └── page.tsx  # Plan comparison
    │   │   │   │   │   │   │   │       # BE: financial-be/subscription
    │   │   │   │   │   │   │   │       # GET /v1/subscriptions/plans/compare
    │   │   │   │   │   │   │   ├── enterprise/
    │   │   │   │   │   │   │   │   └── page.tsx  # Enterprise pricing
    │   │   │   │   │   │   │   │       # BE: None (static)
    │   │   │   │   │   │   │   └── page.tsx  # Pricing page
    │   │   │   │   │   │   │       # - Plan comparison
    │   │   │   │   │   │   │       # - Feature matrix
    │   │   │   │   │   │   │       # - FAQ
    │   │   │   │   │   │   │       # BE: subscriptions-be/plans
    │   │   │   │   │   │   │       # GET /v1/plans?public=true
    │   │   │   │   │   │   │       # Returns: plans with public pricing
    │   │   │   │   │   │   │       # Pricing overview
    │   │   │   │   │   │   │       # BE: financial-be/subscription
    │   │   │   │   │   │   │       # GET /v1/subscriptions/plans
    │   │   │   │   │   │   ├── resources/
    │   │   │   │   │   │   │   ├── api/
    │   │   │   │   │   │   │   │   ├── changelog/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # API changelog
    │   │   │   │   │   │   │   │   ├── docs/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # API documentation
    │   │   │   │   │   │   │   │   └── reference/
    │   │   │   │   │   │   │   │       └── page.tsx  # API reference
    │   │   │   │   │   │   │   ├── case-studies/
    │   │   │   │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Case study detail
    │   │   │   │   │   │   │   │   └── page.tsx  # All case studies
    │   │   │   │   │   │   │   ├── guides/
    │   │   │   │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Guide detail
    │   │   │   │   │   │   │   │   │       # BE: None (static or CMS)
    │   │   │   │   │   │   │   │   └── page.tsx  # All guides
    │   │   │   │   │   │   │   │       # - Getting started
    │   │   │   │   │   │   │   │       # - Best practices
    │   │   │   │   │   │   │   │       # - How-tos
    │   │   │   │   │   │   │   │       # Guides list
    │   │   │   │   │   │   │   │       # BE: None (static or CMS)
    │   │   │   │   │   │   │   ├── tools/
    │   │   │   │   │   │   │   │   ├── budget-estimator/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Project budget estimator
    │   │   │   │   │   │   │   │   ├── rate-calculator/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Hourly rate calculator
    │   │   │   │   │   │   │   │   └── roi-calculator/
    │   │   │   │   │   │   │   │       └── page.tsx  # ROI calculator
    │   │   │   │   │   │   │   ├── tutorials/
    │   │   │   │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Tutorial detail
    │   │   │   │   │   │   │   │   │       # BE: None (static or CMS)
    │   │   │   │   │   │   │   │   └── page.tsx  # Tutorials list
    │   │   │   │   │   │   │   │       # BE: None (static or CMS)
    │   │   │   │   │   │   │   ├── webinars/
    │   │   │   │   │   │   │   │   ├── [id]/
    │   │   │   │   │   │   │   │   │   ├── register/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Register for webinar
    │   │   │   │   │   │   │   │   │   │       # BE: communications-be/events (if exists)
    │   │   │   │   │   │   │   │   │   │       # POST /v1/webinars/{id}/register
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Webinar detail
    │   │   │   │   │   │   │   │   │       # BE: None (static or CMS)
    │   │   │   │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   │   │   │   ├── register/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Webinar registration
    │   │   │   │   │   │   │   │   │   │       # BE: communications-be/events
    │   │   │   │   │   │   │   │   │   │       # POST /v1/webinars/{webinar_id}/register
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Webinar detail
    │   │   │   │   │   │   │   │   ├── on-demand/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # On-demand recordings
    │   │   │   │   │   │   │   │   ├── upcoming/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Upcoming webinars
    │   │   │   │   │   │   │   │   └── page.tsx  # Webinars list
    │   │   │   │   │   │   │   │       # BE: None (static or CMS)
    │   │   │   │   │   │   │   └── page.tsx  # Resources hub
    │   │   │   │   │   │   │       # BE: None (static)
    │   │   │   │   │   │   ├── security/
    │   │   │   │   │   │   │   └── page.tsx  # Security information
    │   │   │   │   │   │   │       # - Security practices
    │   │   │   │   │   │   │       # - Compliance certifications
    │   │   │   │   │   │   │       # - Vulnerability disclosure
    │   │   │   │   │   │   │       # BE: None (static)
    │   │   │   │   │   │   ├── solutions/
    │   │   │   │   │   │   │   ├── agencies/
    │   │   │   │   │   │   │   │   └── page.tsx  # Agency solutions
    │   │   │   │   │   │   │   │       # - White-label
    │   │   │   │   │   │   │   │       # - Multi-client
    │   │   │   │   │   │   │   │       # - Revenue share
    │   │   │   │   │   │   │   ├── by-industry/
    │   │   │   │   │   │   │   │   ├── finance/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Finance solutions
    │   │   │   │   │   │   │   │   ├── marketing/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Marketing solutions
    │   │   │   │   │   │   │   │   ├── tech/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Tech industry solutions
    │   │   │   │   │   │   │   │   └── page.tsx  # Browse by industry
    │   │   │   │   │   │   │   ├── by-role/
    │   │   │   │   │   │   │   │   ├── designers/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Designer hiring solutions
    │   │   │   │   │   │   │   │   ├── developers/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Developer hiring solutions
    │   │   │   │   │   │   │   │   ├── writers/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Writer hiring solutions
    │   │   │   │   │   │   │   │   └── page.tsx  # Browse by role
    │   │   │   │   │   │   │   └── enterprise/
    │   │   │   │   │   │   │       ├── contact-sales/
    │   │   │   │   │   │   │       │   └── page.tsx  # Enterprise contact form
    │   │   │   │   │   │   │       │       # BE: communications-be/contact
    │   │   │   │   │   │   │       │       # POST /v1/contact/enterprise
    │   │   │   │   │   │   │       └── page.tsx  # Enterprise solutions
    │   │   │   │   │   │   │           # - MSA contracts
    │   │   │   │   │   │   │           # - Dedicated support
    │   │   │   │   │   │   │           # - Volume pricing
    │   │   │   │   │   │   ├── status/
    │   │   │   │   │   │   │   └── page.tsx  # System status page
    │   │   │   │   │   │   │       # - Service status
    │   │   │   │   │   │   │       # - Incident history
    │   │   │   │   │   │   │       # BE: admin-be/system (or utility-be/status)
    │   │   │   │   │   │   │       # GET /v1/status
    │   │   │   │   │   │   ├── trust-safety/
    │   │   │   │   │   │   │   └── page.tsx  # Trust & Safety
    │   │   │   │   │   │   │       # - Security measures
    │   │   │   │   │   │   │       # - Payment protection
    │   │   │   │   │   │   │       # - Dispute resolution
    │   │   │   │   │   │   │       # BE: None (static)
    │   │   │   │   │   │   ├── layout.tsx  # Public pages layout
    │   │   │   │   │   │   │   # - Public header (no auth UI)
    │   │   │   │   │   │   │   # - Footer
    │   │   │   │   │   │   │   # - Language switcher
    │   │   │   │   │   │   ├── page.tsx  # Homepage / Landing
    │   │   │   │   │   │   │   # - Hero section
    │   │   │   │   │   │   │   # - Features overview
    │   │   │   │   │   │   │   # - Stats showcase
    │   │   │   │   │   │   │   # - Testimonials
    │   │   │   │   │   │   │   # - CTA sections
    │   │   │   │   │   │   │   # BE: None (static/cached content)
    │   │   │   │   │   │   └── sitemap.xml  # Dynamic sitemap
    │   │   │   │   │   │       # BE: Multiple services for dynamic content
    │   │   │   │   │   ├── (reviews)/
    │   │   │   │   │   │   └── reviews/
    │   │   │   │   │   │       └── new/
    │   │   │   │   │   │           └── [contractId]/
    │   │   │   │   │   │               └── page.tsx  # Write a review
    │   │   │   │   │   │                   # - Post-review, edit within window
    │   │   │   │   │   │                   # BE: reviews-be/review
    │   │   │   │   │   │                   # POST /v1/reviews
    │   │   │   │   │   ├── (social)/
    │   │   │   │   │   │   └── network/
    │   │   │   │   │   │       └── page.tsx  # Network & groups
    │   │   │   │   │   │           # - Connections, groups, referrals
    │   │   │   │   │   │           # BE: users-be/professional_network|user_group|referral
    │   │   │   │   │   │           # GET/POST /v1/network|/v1/groups|/v1/referrals
    │   │   │   │   │   ├── (support)/
    │   │   │   │   │   │   └── disputes/
    │   │   │   │   │   │       └── [contractId]/
    │   │   │   │   │   │           └── page.tsx  # Dispute center
    │   │   │   │   │   │               # - Open case, attach evidence
    │   │   │   │   │   │               # BE: contracts-be/dispute | financial-be/refund | admin-be/refund_case
    │   │   │   │   │   │               # POST /v1/contracts/{id}/disputes
    │   │   │   │   │   ├── contracts/
    │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   ├── change-requests/
    │   │   │   │   │   │   │   │   ├── [requestId]/
    │   │   │   │   │   │   │   │   │   ├── approval-chain/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Change request approvals
    │   │   │   │   │   │   │   │   │   │       # - Approval hierarchy
    │   │   │   │   │   │   │   │   │   │       # - Pending approvals
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/change-request
    │   │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/change-requests/{request_id}/approvals
    │   │   │   │   │   │   │   │   │   ├── impact-analysis/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Change impact analysis
    │   │   │   │   │   │   │   │   │   │       # - Cost impact
    │   │   │   │   │   │   │   │   │   │       # - Timeline impact
    │   │   │   │   │   │   │   │   │   │       # - Resource impact
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/change-request
    │   │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/change-requests/{request_id}/impact
    │   │   │   │   │   │   │   │   │   └── negotiate/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Negotiate change request
    │   │   │   │   │   │   │   │   │           # - Counter proposals
    │   │   │   │   │   │   │   │   │           # - Impact analysis
    │   │   │   │   │   │   │   │   │           # BE: contracts-be/change-request
    │   │   │   │   │   │   │   │   │           # POST /v1/contracts/{contract_id}/change-requests/{request_id}/negotiate
    │   │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │   │       └── page.tsx  # Change request templates
    │   │   │   │   │   │   │   │           # BE: contracts-be/change-request
    │   │   │   │   │   │   │   │           # GET /v1/contracts/change-requests/templates
    │   │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   │   ├── audits/
    │   │   │   │   │   │   │   │   │   ├── [auditId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Audit detail
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/compliance
    │   │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/compliance/audits/{audit_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Audits list
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/compliance
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/compliance/audits
    │   │   │   │   │   │   │   │   ├── documents/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Compliance documents
    │   │   │   │   │   │   │   │   │       # - Certifications
    │   │   │   │   │   │   │   │   │       # - Insurance documents
    │   │   │   │   │   │   │   │   │       # - Legal requirements
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/compliance, storage-be/asset
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/compliance/documents
    │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/compliance/documents
    │   │   │   │   │   │   │   │   └── reports/
    │   │   │   │   │   │   │   │       └── page.tsx  # Compliance reports
    │   │   │   │   │   │   │   │           # - Generate reports
    │   │   │   │   │   │   │   │           # - Export for auditors
    │   │   │   │   │   │   │   │           # BE: contracts-be/compliance
    │   │   │   │   │   │   │   │           # GET /v1/contracts/{contract_id}/compliance/reports
    │   │   │   │   │   │   │   │           # POST /v1/contracts/{contract_id}/compliance/reports/generate
    │   │   │   │   │   │   │   ├── deliverables/
    │   │   │   │   │   │   │   │   ├── [deliverableId]/
    │   │   │   │   │   │   │   │   │   ├── approvals/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approval workflow
    │   │   │   │   │   │   │   │   │   │       # - Multi-step approvals
    │   │   │   │   │   │   │   │   │   │       # - Stakeholder sign-offs
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable
    │   │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/approvals
    │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/approve
    │   │   │   │   │   │   │   │   │   ├── feedback/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Deliverable feedback
    │   │   │   │   │   │   │   │   │   │       # - Annotate files
    │   │   │   │   │   │   │   │   │   │       # - Comment threads
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/deliverable, communications-be/comment
    │   │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/feedback
    │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/feedback
    │   │   │   │   │   │   │   │   │   └── versions/
    │   │   │   │   │   │   │   │   │       ├── [versionId]/
    │   │   │   │   │   │   │   │   │       │   └── page.tsx  # Specific version detail
    │   │   │   │   │   │   │   │   │       │       # BE: contracts-be/deliverable
    │   │   │   │   │   │   │   │   │       │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/versions/{version_id}
    │   │   │   │   │   │   │   │   │       └── compare/
    │   │   │   │   │   │   │   │   │           └── page.tsx  # Compare versions
    │   │   │   │   │   │   │   │   │               # BE: contracts-be/deliverable
    │   │   │   │   │   │   │   │   │               # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/versions/compare
    │   │   │   │   │   │   │   │   └── bulk-operations/
    │   │   │   │   │   │   │   │       └── page.tsx  # Bulk deliverable actions
    │   │   │   │   │   │   │   │           # BE: contracts-be/deliverable
    │   │   │   │   │   │   │   │           # POST /v1/contracts/{contract_id}/deliverables/bulk
    │   │   │   │   │   │   │   ├── knowledge-transfer/
    │   │   │   │   │   │   │   │   ├── checklist/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Transfer checklist
    │   │   │   │   │   │   │   │   │       # - Required items
    │   │   │   │   │   │   │   │   │       # - Completion status
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/knowledge-transfer
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/knowledge-transfer/checklist
    │   │   │   │   │   │   │   │   ├── documentation/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Transfer documentation
    │   │   │   │   │   │   │   │   │       # - Create docs
    │   │   │   │   │   │   │   │   │       # - Track completeness
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/knowledge-transfer, storage-be/asset
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/knowledge-transfer/documentation
    │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/knowledge-transfer/documentation
    │   │   │   │   │   │   │   │   └── sessions/
    │   │   │   │   │   │   │   │       ├── [sessionId]/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Session detail
    │   │   │   │   │   │   │   │       │       # BE: contracts-be/knowledge-transfer
    │   │   │   │   │   │   │   │       │       # GET /v1/contracts/{contract_id}/knowledge-transfer/sessions/{session_id}
    │   │   │   │   │   │   │   │       └── page.tsx  # Sessions list
    │   │   │   │   │   │   │   │           # - Schedule sessions
    │   │   │   │   │   │   │   │           # - Session recordings
    │   │   │   │   │   │   │   │           # BE: contracts-be/knowledge-transfer
    │   │   │   │   │   │   │   │           # GET /v1/contracts/{contract_id}/knowledge-transfer/sessions
    │   │   │   │   │   │   │   ├── milestones/
    │   │   │   │   │   │   │   │   ├── [milestoneId]/
    │   │   │   │   │   │   │   │   │   ├── review/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Review milestone
    │   │   │   │   │   │   │   │   │   │       # - Approve/reject
    │   │   │   │   │   │   │   │   │   │       # - Request changes
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/milestone
    │   │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/milestones/{milestone_id}/review
    │   │   │   │   │   │   │   │   │   ├── revisions/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Milestone revision history
    │   │   │   │   │   │   │   │   │   │       # BE: contracts-be/milestone
    │   │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/milestones/{milestone_id}/revisions
    │   │   │   │   │   │   │   │   │   └── submission/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Submit milestone
    │   │   │   │   │   │   │   │   │           # - Upload deliverables
    │   │   │   │   │   │   │   │   │           # - Add notes
    │   │   │   │   │   │   │   │   │           # BE: contracts-be/milestone, storage-be/asset
    │   │   │   │   │   │   │   │   │           # POST /v1/contracts/{contract_id}/milestones/{milestone_id}/submit
    │   │   │   │   │   │   │   │   └── reorder/
    │   │   │   │   │   │   │   │       └── page.tsx  # Reorder milestones
    │   │   │   │   │   │   │   │           # - Drag and drop
    │   │   │   │   │   │   │   │           # - Adjust dependencies
    │   │   │   │   │   │   │   │           # BE: contracts-be/milestone
    │   │   │   │   │   │   │   │           # PUT /v1/contracts/{contract_id}/milestones/reorder
    │   │   │   │   │   │   │   ├── quality/
    │   │   │   │   │   │   │   │   ├── metrics/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Quality metrics
    │   │   │   │   │   │   │   │   │       # - Code quality
    │   │   │   │   │   │   │   │   │       # - Deliverable quality
    │   │   │   │   │   │   │   │   │       # - Process quality
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/quality
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/quality/metrics
    │   │   │   │   │   │   │   │   ├── reviews/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Quality reviews
    │   │   │   │   │   │   │   │   │       # - Schedule reviews
    │   │   │   │   │   │   │   │   │       # - Review results
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/quality
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/quality/reviews
    │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/quality/reviews
    │   │   │   │   │   │   │   │   └── standards/
    │   │   │   │   │   │   │   │       └── page.tsx  # Quality standards
    │   │   │   │   │   │   │   │           # - Set standards
    │   │   │   │   │   │   │   │           # - Compliance tracking
    │   │   │   │   │   │   │   │           # BE: contracts-be/quality
    │   │   │   │   │   │   │   │           # GET /v1/contracts/{contract_id}/quality/standards
    │   │   │   │   │   │   │   ├── risks/
    │   │   │   │   │   │   │   │   ├── monitoring/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Risk monitoring
    │   │   │   │   │   │   │   │   │       # - Active risks
    │   │   │   │   │   │   │   │   │       # - Risk trends
    │   │   │   │   │   │   │   │   │       # - Alerts
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/risk
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/risks/monitoring
    │   │   │   │   │   │   │   │   ├── register/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Risk register
    │   │   │   │   │   │   │   │   │       # - Identify risks
    │   │   │   │   │   │   │   │   │       # - Risk assessment
    │   │   │   │   │   │   │   │   │       # - Mitigation plans
    │   │   │   │   │   │   │   │   │       # BE: contracts-be/risk
    │   │   │   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/risks
    │   │   │   │   │   │   │   │   │       # POST /v1/contracts/{contract_id}/risks
    │   │   │   │   │   │   │   │   └── reports/
    │   │   │   │   │   │   │   │       └── page.tsx  # Risk reports
    │   │   │   │   │   │   │   │           # BE: contracts-be/risk
    │   │   │   │   │   │   │   │           # GET /v1/contracts/{contract_id}/risks/reports
    │   │   │   │   │   │   │   └── work-diary/
    │   │   │   │   │   │   │       ├── activity-levels/
    │   │   │   │   │   │   │       │   └── page.tsx  # Activity level tracking
    │   │   │   │   │   │   │       │       # - Keyboard/mouse metrics
    │   │   │   │   │   │   │       │       # - Focus time
    │   │   │   │   │   │   │       │       # - Idle time detection
    │   │   │   │   │   │   │       │       # BE: contracts-be/workdiary
    │   │   │   │   │   │   │       │       # GET /v1/contracts/{contract_id}/workdiary/activity
    │   │   │   │   │   │   │       ├── bulk-entry/
    │   │   │   │   │   │   │       │   └── page.tsx  # Bulk time entry
    │   │   │   │   │   │   │       │       # - Multiple entries at once
    │   │   │   │   │   │   │       │       # - Copy from previous week
    │   │   │   │   │   │   │       │       # BE: contracts-be/workdiary
    │   │   │   │   │   │   │       │       # POST /v1/contracts/{contract_id}/workdiary/bulk
    │   │   │   │   │   │   │       ├── corrections/
    │   │   │   │   │   │   │       │   └── page.tsx  # Time entry corrections
    │   │   │   │   │   │   │       │       # - Request corrections
    │   │   │   │   │   │   │       │       # - Approve corrections
    │   │   │   │   │   │   │       │       # BE: contracts-be/workdiary
    │   │   │   │   │   │   │       │       # POST /v1/contracts/{contract_id}/workdiary/corrections
    │   │   │   │   │   │   │       │       # GET /v1/contracts/{contract_id}/workdiary/corrections
    │   │   │   │   │   │   │       └── screenshots/
    │   │   │   │   │   │   │           └── page.tsx  # Screenshot gallery
    │   │   │   │   │   │   │               # - View all screenshots
    │   │   │   │   │   │   │               # - Delete/flag screenshots
    │   │   │   │   │   │   │               # BE: contracts-be/workdiary, storage-be/asset
    │   │   │   │   │   │   │               # GET /v1/contracts/{contract_id}/workdiary/screenshots
    │   │   │   │   │   │   └── benchmarking/
    │   │   │   │   │   │       ├── costs/
    │   │   │   │   │   │       │   └── page.tsx  # Cost benchmarking
    │   │   │   │   │   │       │       # - Rate comparison
    │   │   │   │   │   │       │       # - Budget efficiency
    │   │   │   │   │   │       │       # BE: contracts-be/analytics, financial-be/analytics
    │   │   │   │   │   │       │       # GET /v1/contracts/benchmarks/costs
    │   │   │   │   │   │       ├── performance/
    │   │   │   │   │   │       │   └── page.tsx  # Performance benchmarking
    │   │   │   │   │   │       │       # - Compare against similar contracts
    │   │   │   │   │   │       │       # - Industry standards
    │   │   │   │   │   │       │       # BE: contracts-be/analytics
    │   │   │   │   │   │       │       # GET /v1/contracts/benchmarks/performance
    │   │   │   │   │   │       └── quality/
    │   │   │   │   │   │           └── page.tsx  # Quality benchmarking
    │   │   │   │   │   │               # - Deliverable quality comparison
    │   │   │   │   │   │               # - Client satisfaction
    │   │   │   │   │   │               # BE: contracts-be/analytics, reviews-be/analytics
    │   │   │   │   │   │               # GET /v1/contracts/benchmarks/quality
    │   │   │   │   │   ├── developers/
    │   │   │   │   │   │   ├── api-reference/
    │   │   │   │   │   │   │   ├── [endpoint]/
    │   │   │   │   │   │   │   │   └── page.tsx  # API endpoint reference
    │   │   │   │   │   │   │   │       # BE: static (OpenAPI spec)
    │   │   │   │   │   │   │   └── page.tsx  # API reference home
    │   │   │   │   │   │   │       # BE: static (OpenAPI spec)
    │   │   │   │   │   │   ├── docs/
    │   │   │   │   │   │   │   ├── [section]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Documentation section
    │   │   │   │   │   │   │   │       # BE: static or CMS
    │   │   │   │   │   │   │   │       # GET /v1/content/docs/{section}
    │   │   │   │   │   │   │   └── page.tsx  # API documentation home
    │   │   │   │   │   │   │       # BE: static
    │   │   │   │   │   │   ├── sandbox/
    │   │   │   │   │   │   │   └── page.tsx  # API sandbox/playground
    │   │   │   │   │   │   │       # BE: developer API
    │   │   │   │   │   │   │       # POST /v1/developer/sandbox/execute
    │   │   │   │   │   │   ├── sdks/
    │   │   │   │   │   │   │   └── page.tsx  # SDK downloads and docs
    │   │   │   │   │   │   │       # BE: static
    │   │   │   │   │   │   └── webhooks/
    │   │   │   │   │   │       └── page.tsx  # Webhooks documentation
    │   │   │   │   │   │           # BE: static
    │   │   │   │   │   ├── enterprise/
    │   │   │   │   │   │   ├── case-studies/
    │   │   │   │   │   │   │   └── page.tsx  # Enterprise case studies
    │   │   │   │   │   │   │       # BE: CMS
    │   │   │   │   │   │   │       # GET /v1/content/case-studies?type=enterprise
    │   │   │   │   │   │   ├── contact/
    │   │   │   │   │   │   │   └── page.tsx  # Enterprise contact/demo request
    │   │   │   │   │   │   │       # BE: communications-be
    │   │   │   │   │   │   │       # POST /v1/contact/enterprise
    │   │   │   │   │   │   ├── pricing/
    │   │   │   │   │   │   │   └── page.tsx  # Enterprise pricing
    │   │   │   │   │   │   │       # BE: financial-be/subscription
    │   │   │   │   │   │   │       # GET /v1/subscriptions/plans?type=enterprise
    │   │   │   │   │   │   └── solutions/
    │   │   │   │   │   │       ├── managed-services/
    │   │   │   │   │   │       │   └── page.tsx  # Managed services offering
    │   │   │   │   │   │       │       # BE: none (marketing content)
    │   │   │   │   │   │       ├── staffing/
    │   │   │   │   │   │       │   └── page.tsx  # Enterprise staffing solutions
    │   │   │   │   │   │       │       # BE: none (marketing content)
    │   │   │   │   │   │       └── page.tsx  # Enterprise solutions overview
    │   │   │   │   │   │           # BE: none (marketing content)
    │   │   │   │   │   ├── financial/
    │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   ├── kpis/
    │   │   │   │   │   │   │   │   ├── custom/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Custom KPI builder
    │   │   │   │   │   │   │   │   │       # BE: financial-be/analytics
    │   │   │   │   │   │   │   │   │       # POST /v1/analytics/kpis/custom
    │   │   │   │   │   │   │   │   └── dashboard/
    │   │   │   │   │   │   │   │       └── page.tsx  # Financial KPI dashboard
    │   │   │   │   │   │   │   │           # - Custom KPIs
    │   │   │   │   │   │   │   │           # - Benchmarking
    │   │   │   │   │   │   │   │           # BE: financial-be/analytics
    │   │   │   │   │   │   │   │           # GET /v1/analytics/kpis/dashboard
    │   │   │   │   │   │   │   ├── margins/
    │   │   │   │   │   │   │   │   └── page.tsx  # Margin analysis
    │   │   │   │   │   │   │   │       # - Gross margin
    │   │   │   │   │   │   │   │       # - Operating margin
    │   │   │   │   │   │   │   │       # - Net margin
    │   │   │   │   │   │   │   │       # BE: financial-be/analytics
    │   │   │   │   │   │   │   │       # GET /v1/analytics/margins
    │   │   │   │   │   │   │   └── profitability/
    │   │   │   │   │   │   │       ├── by-client/
    │   │   │   │   │   │   │       │   └── page.tsx  # Client profitability
    │   │   │   │   │   │   │       │       # BE: financial-be/analytics
    │   │   │   │   │   │   │       │       # GET /v1/analytics/profitability/by-client
    │   │   │   │   │   │   │       ├── by-project/
    │   │   │   │   │   │   │       │   └── page.tsx  # Project profitability
    │   │   │   │   │   │   │       │       # BE: financial-be/analytics
    │   │   │   │   │   │   │       │       # GET /v1/analytics/profitability/by-project
    │   │   │   │   │   │   │       └── by-service/
    │   │   │   │   │   │   │           └── page.tsx  # Service line profitability
    │   │   │   │   │   │   │               # BE: financial-be/analytics
    │   │   │   │   │   │   │               # GET /v1/analytics/profitability/by-service
    │   │   │   │   │   │   ├── budgets/
    │   │   │   │   │   │   │   ├── [budgetId]/
    │   │   │   │   │   │   │   │   ├── adjustments/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Budget adjustments
    │   │   │   │   │   │   │   │   │       # - Reallocation requests
    │   │   │   │   │   │   │   │   │       # - Approval workflow
    │   │   │   │   │   │   │   │   │       # BE: financial-be/budget
    │   │   │   │   │   │   │   │   │       # GET /v1/budgets/{budget_id}/adjustments
    │   │   │   │   │   │   │   │   │       # POST /v1/budgets/{budget_id}/adjustments
    │   │   │   │   │   │   │   │   ├── allocations/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Budget allocations
    │   │   │   │   │   │   │   │   │       # - Department allocations
    │   │   │   │   │   │   │   │   │       # - Project allocations
    │   │   │   │   │   │   │   │   │       # BE: financial-be/budget
    │   │   │   │   │   │   │   │   │       # GET /v1/budgets/{budget_id}/allocations
    │   │   │   │   │   │   │   │   │       # POST /v1/budgets/{budget_id}/allocations
    │   │   │   │   │   │   │   │   ├── forecasts/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Budget forecasts
    │   │   │   │   │   │   │   │   │       # - Rolling forecasts
    │   │   │   │   │   │   │   │   │       # - Projection scenarios
    │   │   │   │   │   │   │   │   │       # BE: financial-be/budget
    │   │   │   │   │   │   │   │   │       # GET /v1/budgets/{budget_id}/forecasts
    │   │   │   │   │   │   │   │   └── variance/
    │   │   │   │   │   │   │   │       └── page.tsx  # Budget variance
    │   │   │   │   │   │   │   │           # - Actual vs. planned
    │   │   │   │   │   │   │   │           # - Variance analysis
    │   │   │   │   │   │   │   │           # BE: financial-be/budget
    │   │   │   │   │   │   │   │           # GET /v1/budgets/{budget_id}/variance
    │   │   │   │   │   │   │   ├── consolidation/
    │   │   │   │   │   │   │   │   └── page.tsx  # Budget consolidation
    │   │   │   │   │   │   │   │       # - Roll-up view
    │   │   │   │   │   │   │   │       # - Cross-department view
    │   │   │   │   │   │   │   │       # BE: financial-be/budget
    │   │   │   │   │   │   │   │       # GET /v1/budgets/consolidation
    │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │       └── page.tsx  # Budget templates
    │   │   │   │   │   │   │           # - Standard templates
    │   │   │   │   │   │   │           # - Create from template
    │   │   │   │   │   │   │           # BE: financial-be/budget
    │   │   │   │   │   │   │           # GET /v1/budgets/templates
    │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   ├── aml/
    │   │   │   │   │   │   │   │   ├── monitoring/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # AML monitoring
    │   │   │   │   │   │   │   │   │       # - Transaction monitoring
    │   │   │   │   │   │   │   │   │       # - Suspicious activity
    │   │   │   │   │   │   │   │   │       # BE: financial-be/compliance
    │   │   │   │   │   │   │   │   │       # GET /v1/compliance/aml/monitoring
    │   │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # AML reports
    │   │   │   │   │   │   │   │   │       # - SAR filing
    │   │   │   │   │   │   │   │   │       # - Compliance reports
    │   │   │   │   │   │   │   │   │       # BE: financial-be/compliance
    │   │   │   │   │   │   │   │   │       # GET /v1/compliance/aml/reports
    │   │   │   │   │   │   │   │   │       # POST /v1/compliance/aml/reports/sar
    │   │   │   │   │   │   │   │   └── page.tsx  # AML dashboard
    │   │   │   │   │   │   │   │       # BE: financial-be/compliance
    │   │   │   │   │   │   │   │       # GET /v1/compliance/aml
    │   │   │   │   │   │   │   ├── audits/
    │   │   │   │   │   │   │   │   ├── findings/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Audit findings
    │   │   │   │   │   │   │   │   │       # BE: financial-be/compliance
    │   │   │   │   │   │   │   │   │       # GET /v1/compliance/audits/findings
    │   │   │   │   │   │   │   │   ├── remediation/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Remediation tracking
    │   │   │   │   │   │   │   │   │       # BE: financial-be/compliance
    │   │   │   │   │   │   │   │   │       # GET /v1/compliance/audits/remediation
    │   │   │   │   │   │   │   │   └── schedule/
    │   │   │   │   │   │   │   │       └── page.tsx  # Audit schedule
    │   │   │   │   │   │   │   │           # BE: financial-be/compliance
    │   │   │   │   │   │   │   │           # GET /v1/compliance/audits/schedule
    │   │   │   │   │   │   │   ├── kyc/
    │   │   │   │   │   │   │   │   ├── due-diligence/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Enhanced due diligence
    │   │   │   │   │   │   │   │   │       # BE: financial-be/compliance
    │   │   │   │   │   │   │   │   │       # GET /v1/compliance/kyc/due-diligence
    │   │   │   │   │   │   │   │   └── verification/
    │   │   │   │   │   │   │   │       └── page.tsx  # KYC verification
    │   │   │   │   │   │   │   │           # BE: financial-be/compliance, admin-be/kyc
    │   │   │   │   │   │   │   │           # GET /v1/compliance/kyc/verification
    │   │   │   │   │   │   │   └── sanctions/
    │   │   │   │   │   │   │       ├── alerts/
    │   │   │   │   │   │   │       │   └── page.tsx  # Sanctions alerts
    │   │   │   │   │   │   │       │       # BE: financial-be/compliance
    │   │   │   │   │   │   │       │       # GET /v1/compliance/sanctions/alerts
    │   │   │   │   │   │   │       └── screening/
    │   │   │   │   │   │   │           └── page.tsx  # Sanctions screening
    │   │   │   │   │   │   │               # - Watchlist screening
    │   │   │   │   │   │   │               # - PEP screening
    │   │   │   │   │   │   │               # BE: financial-be/compliance
    │   │   │   │   │   │   │               # POST /v1/compliance/sanctions/screen
    │   │   │   │   │   │   ├── credit-management/
    │   │   │   │   │   │   │   ├── collections/
    │   │   │   │   │   │   │   │   ├── actions/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Collection actions
    │   │   │   │   │   │   │   │   │       # - Reminders
    │   │   │   │   │   │   │   │   │       # - Escalations
    │   │   │   │   │   │   │   │   │       # BE: financial-be/credit, communications-be
    │   │   │   │   │   │   │   │   │       # POST /v1/credit/collections/actions
    │   │   │   │   │   │   │   │   ├── aging/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Aging analysis
    │   │   │   │   │   │   │   │   │       # - 30/60/90 day buckets
    │   │   │   │   │   │   │   │   │       # - Collection priority
    │   │   │   │   │   │   │   │   │       # BE: financial-be/credit
    │   │   │   │   │   │   │   │   │       # GET /v1/credit/collections/aging
    │   │   │   │   │   │   │   │   └── page.tsx  # Collections dashboard
    │   │   │   │   │   │   │   │       # BE: financial-be/credit
    │   │   │   │   │   │   │   │       # GET /v1/credit/collections
    │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   └── page.tsx  # Credit disputes
    │   │   │   │   │   │   │   │       # BE: financial-be/credit
    │   │   │   │   │   │   │   │       # GET /v1/credit/disputes
    │   │   │   │   │   │   │   ├── limits/
    │   │   │   │   │   │   │   │   ├── [clientId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Client credit limit
    │   │   │   │   │   │   │   │   │       # - Set/adjust limit
    │   │   │   │   │   │   │   │   │       # - Usage tracking
    │   │   │   │   │   │   │   │   │       # BE: financial-be/credit
    │   │   │   │   │   │   │   │   │       # GET /v1/credit/limits/{client_id}
    │   │   │   │   │   │   │   │   │       # PUT /v1/credit/limits/{client_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # All credit limits
    │   │   │   │   │   │   │   │       # BE: financial-be/credit
    │   │   │   │   │   │   │   │       # GET /v1/credit/limits
    │   │   │   │   │   │   │   └── scoring/
    │   │   │   │   │   │   │       └── page.tsx  # Credit scoring
    │   │   │   │   │   │   │           # - Credit assessment
    │   │   │   │   │   │   │           # - Risk rating
    │   │   │   │   │   │   │           # BE: financial-be/credit
    │   │   │   │   │   │   │           # GET /v1/credit/scoring
    │   │   │   │   │   │   │           # POST /v1/credit/scoring/assess
    │   │   │   │   │   │   ├── forecasting/
    │   │   │   │   │   │   │   ├── expenses/
    │   │   │   │   │   │   │   │   └── page.tsx  # Expense forecast
    │   │   │   │   │   │   │   │       # BE: financial-be/forecasting
    │   │   │   │   │   │   │   │       # GET /v1/forecasting/expenses
    │   │   │   │   │   │   │   ├── profitability/
    │   │   │   │   │   │   │   │   └── page.tsx  # Profitability forecast
    │   │   │   │   │   │   │   │       # - Margin analysis
    │   │   │   │   │   │   │   │       # - Break-even analysis
    │   │   │   │   │   │   │   │       # BE: financial-be/forecasting
    │   │   │   │   │   │   │   │       # GET /v1/forecasting/profitability
    │   │   │   │   │   │   │   ├── revenue/
    │   │   │   │   │   │   │   │   ├── models/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Revenue models
    │   │   │   │   │   │   │   │   │       # - Historical models
    │   │   │   │   │   │   │   │   │       # - Predictive models
    │   │   │   │   │   │   │   │   │       # BE: financial-be/forecasting
    │   │   │   │   │   │   │   │   │       # GET /v1/forecasting/revenue/models
    │   │   │   │   │   │   │   │   ├── scenarios/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Revenue scenarios
    │   │   │   │   │   │   │   │   │       # - Best/worst/likely cases
    │   │   │   │   │   │   │   │   │       # - Sensitivity analysis
    │   │   │   │   │   │   │   │   │       # BE: financial-be/forecasting
    │   │   │   │   │   │   │   │   │       # GET /v1/forecasting/revenue/scenarios
    │   │   │   │   │   │   │   │   │       # POST /v1/forecasting/revenue/scenarios
    │   │   │   │   │   │   │   │   └── page.tsx  # Revenue forecast
    │   │   │   │   │   │   │   │       # BE: financial-be/forecasting
    │   │   │   │   │   │   │   │       # GET /v1/forecasting/revenue
    │   │   │   │   │   │   │   └── validation/
    │   │   │   │   │   │   │       └── page.tsx  # Forecast validation
    │   │   │   │   │   │   │           # - Accuracy tracking
    │   │   │   │   │   │   │           # - Model refinement
    │   │   │   │   │   │   │           # BE: financial-be/forecasting
    │   │   │   │   │   │   │           # GET /v1/forecasting/validation
    │   │   │   │   │   │   ├── reconciliation/
    │   │   │   │   │   │   │   ├── bank/
    │   │   │   │   │   │   │   │   ├── [accountId]/
    │   │   │   │   │   │   │   │   │   ├── auto-match/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Auto-matching
    │   │   │   │   │   │   │   │   │   │       # - Rule-based matching
    │   │   │   │   │   │   │   │   │   │       # - AI-powered suggestions
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/reconciliation
    │   │   │   │   │   │   │   │   │   │       # POST /v1/reconciliation/bank/{account_id}/auto-match
    │   │   │   │   │   │   │   │   │   ├── discrepancies/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Discrepancy resolution
    │   │   │   │   │   │   │   │   │   │       # BE: financial-be/reconciliation
    │   │   │   │   │   │   │   │   │   │       # GET /v1/reconciliation/bank/{account_id}/discrepancies
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Account reconciliation
    │   │   │   │   │   │   │   │   │       # BE: financial-be/reconciliation
    │   │   │   │   │   │   │   │   │       # GET /v1/reconciliation/bank/{account_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Bank reconciliation list
    │   │   │   │   │   │   │   │       # BE: financial-be/reconciliation
    │   │   │   │   │   │   │   │       # GET /v1/reconciliation/bank
    │   │   │   │   │   │   │   ├── intercompany/
    │   │   │   │   │   │   │   │   └── page.tsx  # Intercompany reconciliation
    │   │   │   │   │   │   │   │       # BE: financial-be/reconciliation
    │   │   │   │   │   │   │   │       # GET /v1/reconciliation/intercompany
    │   │   │   │   │   │   │   ├── merchant/
    │   │   │   │   │   │   │   │   └── page.tsx  # Merchant account reconciliation
    │   │   │   │   │   │   │   │       # - Payment processor reconciliation
    │   │   │   │   │   │   │   │       # - Fee reconciliation
    │   │   │   │   │   │   │   │       # BE: financial-be/reconciliation
    │   │   │   │   │   │   │   │       # GET /v1/reconciliation/merchant
    │   │   │   │   │   │   │   └── reports/
    │   │   │   │   │   │   │       └── page.tsx  # Reconciliation reports
    │   │   │   │   │   │   │           # - Status reports
    │   │   │   │   │   │   │           # - Exception reports
    │   │   │   │   │   │   │           # BE: financial-be/reconciliation
    │   │   │   │   │   │   │           # GET /v1/reconciliation/reports
    │   │   │   │   │   │   ├── risk-management/
    │   │   │   │   │   │   │   ├── exposure/
    │   │   │   │   │   │   │   │   ├── credit/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Credit exposure
    │   │   │   │   │   │   │   │   │       # - Counterparty risk
    │   │   │   │   │   │   │   │   │       # - Concentration risk
    │   │   │   │   │   │   │   │   │       # BE: financial-be/risk
    │   │   │   │   │   │   │   │   │       # GET /v1/risk/exposure/credit
    │   │   │   │   │   │   │   │   ├── currency/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Currency exposure
    │   │   │   │   │   │   │   │   │       # - FX risk analysis
    │   │   │   │   │   │   │   │   │       # - Hedging strategies
    │   │   │   │   │   │   │   │   │       # BE: financial-be/risk
    │   │   │   │   │   │   │   │   │       # GET /v1/risk/exposure/currency
    │   │   │   │   │   │   │   │   └── operational/
    │   │   │   │   │   │   │   │       └── page.tsx  # Operational risk
    │   │   │   │   │   │   │   │           # - Process risks
    │   │   │   │   │   │   │   │           # - System risks
    │   │   │   │   │   │   │   │           # BE: financial-be/risk
    │   │   │   │   │   │   │   │           # GET /v1/risk/exposure/operational
    │   │   │   │   │   │   │   ├── limits/
    │   │   │   │   │   │   │   │   └── page.tsx  # Risk limits
    │   │   │   │   │   │   │   │       # - Set risk limits
    │   │   │   │   │   │   │   │       # - Monitor breaches
    │   │   │   │   │   │   │   │       # BE: financial-be/risk
    │   │   │   │   │   │   │   │       # GET /v1/risk/limits
    │   │   │   │   │   │   │   │       # PUT /v1/risk/limits
    │   │   │   │   │   │   │   └── reports/
    │   │   │   │   │   │   │       ├── stress-testing/
    │   │   │   │   │   │   │       │   └── page.tsx  # Stress testing
    │   │   │   │   │   │   │       │       # - Scenario analysis
    │   │   │   │   │   │   │       │       # - Impact assessment
    │   │   │   │   │   │   │       │       # BE: financial-be/risk
    │   │   │   │   │   │   │       │       # GET /v1/risk/reports/stress-testing
    │   │   │   │   │   │   │       └── var/
    │   │   │   │   │   │   │           └── page.tsx  # Value at Risk
    │   │   │   │   │   │   │               # BE: financial-be/risk
    │   │   │   │   │   │   │               # GET /v1/risk/reports/var
    │   │   │   │   │   │   └── treasury/
    │   │   │   │   │   │       ├── cash-flow/
    │   │   │   │   │   │       │   ├── analysis/
    │   │   │   │   │   │       │   │   └── page.tsx  # Cash flow analysis
    │   │   │   │   │   │       │   │       # - Historical trends
    │   │   │   │   │   │       │   │       # - Variance analysis
    │   │   │   │   │   │       │   │       # BE: financial-be/treasury
    │   │   │   │   │   │       │   │       # GET /v1/treasury/cash-flow/analysis
    │   │   │   │   │   │       │   ├── forecast/
    │   │   │   │   │   │       │   │   └── page.tsx  # Cash flow forecasting
    │   │   │   │   │   │       │   │       # - 30/60/90 day forecast
    │   │   │   │   │   │       │   │       # - Scenario planning
    │   │   │   │   │   │       │   │       # BE: financial-be/treasury
    │   │   │   │   │   │       │   │       # GET /v1/treasury/cash-flow/forecast
    │   │   │   │   │   │       │   └── reports/
    │   │   │   │   │   │       │       └── page.tsx  # Cash flow reports
    │   │   │   │   │   │       │           # BE: financial-be/treasury
    │   │   │   │   │   │       │           # GET /v1/treasury/cash-flow/reports
    │   │   │   │   │   │       ├── dashboard/
    │   │   │   │   │   │       │   └── page.tsx  # Treasury dashboard
    │   │   │   │   │   │       │       # - Cash position
    │   │   │   │   │   │       │       # - Liquidity forecast
    │   │   │   │   │   │       │       # - Working capital
    │   │   │   │   │   │       │       # BE: financial-be/treasury
    │   │   │   │   │   │       │       # GET /v1/treasury/dashboard
    │   │   │   │   │   │       ├── investments/
    │   │   │   │   │   │       │   └── page.tsx  # Short-term investments
    │   │   │   │   │   │       │       # - Investment tracking
    │   │   │   │   │   │       │       # - Returns analysis
    │   │   │   │   │   │       │       # BE: financial-be/treasury
    │   │   │   │   │   │       │       # GET /v1/treasury/investments
    │   │   │   │   │   │       └── liquidity/
    │   │   │   │   │   │           ├── alerts/
    │   │   │   │   │   │           │   └── page.tsx  # Liquidity alerts
    │   │   │   │   │   │           │       # - Threshold monitoring
    │   │   │   │   │   │           │       # - Alert configuration
    │   │   │   │   │   │           │       # BE: financial-be/treasury
    │   │   │   │   │   │           │       # GET /v1/treasury/liquidity/alerts
    │   │   │   │   │   │           │       # PUT /v1/treasury/liquidity/alerts
    │   │   │   │   │   │           └── management/
    │   │   │   │   │   │               └── page.tsx  # Liquidity management
    │   │   │   │   │   │                   # - Current ratio
    │   │   │   │   │   │                   # - Quick ratio
    │   │   │   │   │   │                   # - Reserve requirements
    │   │   │   │   │   │                   # BE: financial-be/treasury
    │   │   │   │   │   │                   # GET /v1/treasury/liquidity
    │   │   │   │   │   ├── legal/
    │   │   │   │   │   │   ├── accessibility/
    │   │   │   │   │   │   │   └── page.tsx  # Accessibility statement
    │   │   │   │   │   │   │       # BE: none (static content)
    │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   ├── ccpa/
    │   │   │   │   │   │   │   │   └── page.tsx  # CCPA compliance
    │   │   │   │   │   │   │   │       # BE: none (static content)
    │   │   │   │   │   │   │   ├── gdpr/
    │   │   │   │   │   │   │   │   └── page.tsx  # GDPR compliance
    │   │   │   │   │   │   │   │       # BE: none (static content)
    │   │   │   │   │   │   │   └── page.tsx  # Compliance overview
    │   │   │   │   │   │   │       # BE: none (static content)
    │   │   │   │   │   │   ├── dmca/
    │   │   │   │   │   │   │   └── page.tsx  # DMCA policy
    │   │   │   │   │   │   │       # BE: none (static content)
    │   │   │   │   │   │   ├── ip-policy/
    │   │   │   │   │   │   │   └── page.tsx  # Intellectual property policy
    │   │   │   │   │   │   │       # BE: none (static content)
    │   │   │   │   │   │   ├── privacy/
    │   │   │   │   │   │   │   ├── cookie-policy/
    │   │   │   │   │   │   │   │   └── page.tsx  # Cookie policy
    │   │   │   │   │   │   │   │       # BE: none (static content)
    │   │   │   │   │   │   │   ├── data-processing/
    │   │   │   │   │   │   │   │   └── page.tsx  # Data processing agreement
    │   │   │   │   │   │   │   │       # BE: none (static content)
    │   │   │   │   │   │   │   └── page.tsx  # Privacy policy
    │   │   │   │   │   │   │       # BE: none (static content)
    │   │   │   │   │   │   └── terms/
    │   │   │   │   │   │       ├── client/
    │   │   │   │   │   │       │   └── page.tsx  # Client terms of service
    │   │   │   │   │   │       │       # BE: none (static content)
    │   │   │   │   │   │       ├── freelancer/
    │   │   │   │   │   │       │   └── page.tsx  # Freelancer terms of service
    │   │   │   │   │   │       │       # BE: none (static content with version from CMS)
    │   │   │   │   │   │       └── page.tsx  # General terms
    │   │   │   │   │   │           # BE: none (static content)
    │   │   │   │   │   ├── notifications/
    │   │   │   │   │   │   └── page.tsx  # Notifications center (web)
    │   │   │   │   │   │       # - In-app notifications, bulk mark read
    │   │   │   │   │   │       # BE: communications-be/notification
    │   │   │   │   │   │       # GET  /v1/notifications?mine=1
    │   │   │   │   │   │       # POST /v1/notifications/{id}/read
    │   │   │   │   │   ├── pricing/
    │   │   │   │   │   │   └── page.tsx  # Pricing (public)
    │   │   │   │   │   │       # - Plans overview, compare
    │   │   │   │   │   │       # BE: financial-be/subscription (plans)
    │   │   │   │   │   │       # GET /v1/subscriptions/plans
    │   │   │   │   │   ├── resources/
    │   │   │   │   │   │   ├── blog/
    │   │   │   │   │   │   │   ├── [postId]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Blog post
    │   │   │   │   │   │   │   │       # BE: CMS
    │   │   │   │   │   │   │   │       # GET /v1/content/blog/{post_id}
    │   │   │   │   │   │   │   ├── category/
    │   │   │   │   │   │   │   │   └── [categoryId]/
    │   │   │   │   │   │   │   │       └── page.tsx  # Blog category
    │   │   │   │   │   │   │   │           # BE: CMS
    │   │   │   │   │   │   │   │           # GET /v1/content/blog?category={category_id}
    │   │   │   │   │   │   │   └── page.tsx  # Blog home
    │   │   │   │   │   │   │       # BE: CMS
    │   │   │   │   │   │   │       # GET /v1/content/blog
    │   │   │   │   │   │   ├── case-studies/
    │   │   │   │   │   │   │   ├── [caseStudyId]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Case study detail
    │   │   │   │   │   │   │   │       # BE: CMS
    │   │   │   │   │   │   │   │       # GET /v1/content/case-studies/{case_study_id}
    │   │   │   │   │   │   │   └── page.tsx  # Case studies list
    │   │   │   │   │   │   │       # BE: CMS
    │   │   │   │   │   │   │       # GET /v1/content/case-studies
    │   │   │   │   │   │   ├── faq/
    │   │   │   │   │   │   │   └── page.tsx  # Frequently asked questions
    │   │   │   │   │   │   │       # BE: CMS
    │   │   │   │   │   │   │       # GET /v1/content/faq
    │   │   │   │   │   │   ├── guides/
    │   │   │   │   │   │   │   ├── [guideId]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Guide detail
    │   │   │   │   │   │   │   │       # BE: CMS or static
    │   │   │   │   │   │   │   │       # GET /v1/content/guides/{guide_id}
    │   │   │   │   │   │   │   ├── client/
    │   │   │   │   │   │   │   │   └── page.tsx  # Client guides
    │   │   │   │   │   │   │   │       # BE: CMS
    │   │   │   │   │   │   │   │       # GET /v1/content/guides?category=client
    │   │   │   │   │   │   │   ├── freelancer/
    │   │   │   │   │   │   │   │   └── page.tsx  # Freelancer guides
    │   │   │   │   │   │   │   │       # BE: CMS
    │   │   │   │   │   │   │   │       # GET /v1/content/guides?category=freelancer
    │   │   │   │   │   │   │   └── page.tsx  # All guides
    │   │   │   │   │   │   │       # BE: CMS
    │   │   │   │   │   │   │       # GET /v1/content/guides
    │   │   │   │   │   │   ├── tutorials/
    │   │   │   │   │   │   │   ├── [tutorialId]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Tutorial detail
    │   │   │   │   │   │   │   │       # BE: CMS
    │   │   │   │   │   │   │   │       # GET /v1/content/tutorials/{tutorial_id}
    │   │   │   │   │   │   │   └── page.tsx  # Tutorials list
    │   │   │   │   │   │   │       # BE: CMS
    │   │   │   │   │   │   │       # GET /v1/content/tutorials
    │   │   │   │   │   │   └── webinars/
    │   │   │   │   │   │       ├── [webinarId]/
    │   │   │   │   │   │       │   └── page.tsx  # Webinar detail & registration
    │   │   │   │   │   │       │       # BE: CMS + registration system
    │   │   │   │   │   │       │       # GET /v1/content/webinars/{webinar_id}
    │   │   │   │   │   │       │       # POST /v1/webinars/{webinar_id}/register
    │   │   │   │   │   │       └── page.tsx  # Upcoming webinars
    │   │   │   │   │   │           # BE: CMS
    │   │   │   │   │   │           # GET /v1/content/webinars
    │   │   │   │   │   ├── search/
    │   │   │   │   │   │   ├── alerts/
    │   │   │   │   │   │   │   └── page.tsx  # Search alerts (web)
    │   │   │   │   │   │   │       # - Manage email/push alerts for queries
    │   │   │   │   │   │   │       # BE: search-be/alert (proposed if missing)
    │   │   │   │   │   │   │       # GET/POST/DELETE /v1/search/alerts
    │   │   │   │   │   │   ├── assist/
    │   │   │   │   │   │   │   └── page.tsx  # Search assist
    │   │   │   │   │   │   │       # - Suggestions, speller, rewrites, languages
    │   │   │   │   │   │   │       # BE: search-be/suggestions|speller|rewrites|languages
    │   │   │   │   │   │   │       # GET /v1/suggestions
    │   │   │   │   │   │   │       # GET /v1/speller
    │   │   │   │   │   │   │       # GET /v1/rewrites
    │   │   │   │   │   │   │       # GET /v1/languages
    │   │   │   │   │   │   ├── feed/
    │   │   │   │   │   │   │   └── page.tsx  # Personalized feed
    │   │   │   │   │   │   │       # - Jobs/talent recommendations
    │   │   │   │   │   │   │       # BE: search-be/feed|similarity|trending
    │   │   │   │   │   │   │       # GET /v1/feed
    │   │   │   │   │   │   │       # GET /v1/trending
    │   │   │   │   │   │   ├── portfolios/
    │   │   │   │   │   │   │   └── page.tsx  # Portfolio search
    │   │   │   │   │   │   │       # - Filter by skills/tags
    │   │   │   │   │   │   │       # BE: search-be/portfolio_index
    │   │   │   │   │   │   │       # GET /v1/search/portfolios
    │   │   │   │   │   │   └── saved/
    │   │   │   │   │   │       └── page.tsx  # Saved searches (web)
    │   │   │   │   │   │           # - List/create/delete saved queries
    │   │   │   │   │   │           # BE: search-be/saved_query (proposed if missing)
    │   │   │   │   │   │           # GET/POST/DELETE /v1/search/saved
    │   │   │   │   │   ├── security/
    │   │   │   │   │   │   ├── bug-bounty/
    │   │   │   │   │   │   │   └── page.tsx  # Bug bounty program
    │   │   │   │   │   │   │       # BE: none (static content)
    │   │   │   │   │   │   ├── certifications/
    │   │   │   │   │   │   │   └── page.tsx  # Security certifications (SOC2, ISO, etc.)
    │   │   │   │   │   │   │       # BE: none (static content)
    │   │   │   │   │   │   ├── overview/
    │   │   │   │   │   │   │   └── page.tsx  # Security overview
    │   │   │   │   │   │   │       # BE: none (static content)
    │   │   │   │   │   │   └── responsible-disclosure/
    │   │   │   │   │   │       └── page.tsx  # Responsible disclosure policy
    │   │   │   │   │   │           # BE: none (static content)
    │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   ├── account/
    │   │   │   │   │   │   │   └── saved-items/
    │   │   │   │   │   │   │       └── page.tsx  # Saved items
    │   │   │   │   │   │   │           # - Saved jobs/talent, remove
    │   │   │   │   │   │   │           # BE: users-be/saved_item
    │   │   │   │   │   │   │           # GET/DELETE /v1/saved-items
    │   │   │   │   │   │   ├── advanced/
    │   │   │   │   │   │   │   ├── audit-logs/
    │   │   │   │   │   │   │   │   ├── export/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Export audit logs
    │   │   │   │   │   │   │   │   │       # BE: admin-be/audit
    │   │   │   │   │   │   │   │   │       # POST /v1/settings/audit-logs/export
    │   │   │   │   │   │   │   │   └── page.tsx  # Audit log settings
    │   │   │   │   │   │   │   │       # - Retention period
    │   │   │   │   │   │   │   │       # - Log level
    │   │   │   │   │   │   │   │       # BE: admin-be/audit
    │   │   │   │   │   │   │   │       # GET /v1/settings/audit-logs
    │   │   │   │   │   │   │   │       # PUT /v1/settings/audit-logs
    │   │   │   │   │   │   │   ├── data-retention/
    │   │   │   │   │   │   │   │   └── page.tsx  # Data retention policies
    │   │   │   │   │   │   │   │       # - Retention rules
    │   │   │   │   │   │   │   │       # - Deletion schedules
    │   │   │   │   │   │   │   │       # BE: admin-be/data-governance
    │   │   │   │   │   │   │   │       # GET /v1/settings/data-retention
    │   │   │   │   │   │   │   │       # PUT /v1/settings/data-retention
    │   │   │   │   │   │   │   ├── ip-whitelist/
    │   │   │   │   │   │   │   │   └── page.tsx  # IP whitelist
    │   │   │   │   │   │   │   │       # - Add/remove IPs
    │   │   │   │   │   │   │   │       # - CIDR ranges
    │   │   │   │   │   │   │   │       # BE: admin-be/security
    │   │   │   │   │   │   │   │       # GET /v1/settings/ip-whitelist
    │   │   │   │   │   │   │   │       # PUT /v1/settings/ip-whitelist
    │   │   │   │   │   │   │   ├── rate-limiting/
    │   │   │   │   │   │   │   │   └── page.tsx  # Rate limit configuration
    │   │   │   │   │   │   │   │       # - API rate limits
    │   │   │   │   │   │   │   │       # - Custom rules
    │   │   │   │   │   │   │   │       # BE: admin-be/security
    │   │   │   │   │   │   │   │       # GET /v1/settings/rate-limits
    │   │   │   │   │   │   │   │       # PUT /v1/settings/rate-limits
    │   │   │   │   │   │   │   └── session-management/
    │   │   │   │   │   │   │       └── page.tsx  # Session settings
    │   │   │   │   │   │   │           # - Timeout duration
    │   │   │   │   │   │   │           # - Concurrent sessions
    │   │   │   │   │   │   │           # - Force logout
    │   │   │   │   │   │   │           # BE: admin-be/security
    │   │   │   │   │   │   │           # GET /v1/settings/sessions
    │   │   │   │   │   │   │           # PUT /v1/settings/sessions
    │   │   │   │   │   │   ├── developer/
    │   │   │   │   │   │   │   ├── api-keys/
    │   │   │   │   │   │   │   │   ├── [keyId]/
    │   │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # API key usage logs
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/api-keys
    │   │   │   │   │   │   │   │   │   │       # GET /v1/developer/api-keys/{key_id}/logs
    │   │   │   │   │   │   │   │   │   ├── permissions/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Key permissions
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/api-keys
    │   │   │   │   │   │   │   │   │   │       # GET /v1/developer/api-keys/{key_id}/permissions
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/developer/api-keys/{key_id}/permissions
    │   │   │   │   │   │   │   │   │   └── rotate/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Rotate API key
    │   │   │   │   │   │   │   │   │           # BE: admin-be/api-keys
    │   │   │   │   │   │   │   │   │           # POST /v1/developer/api-keys/{key_id}/rotate
    │   │   │   │   │   │   │   │   └── page.tsx  # API keys management
    │   │   │   │   │   │   │   │       # BE: admin-be/api-keys
    │   │   │   │   │   │   │   │       # GET /v1/developer/api-keys
    │   │   │   │   │   │   │   │       # POST /v1/developer/api-keys
    │   │   │   │   │   │   │   ├── documentation/
    │   │   │   │   │   │   │   │   └── page.tsx  # API documentation
    │   │   │   │   │   │   │   │       # - Interactive docs
    │   │   │   │   │   │   │   │       # - Code samples
    │   │   │   │   │   │   │   │       # BE: None (static with API client examples)
    │   │   │   │   │   │   │   ├── oauth-apps/
    │   │   │   │   │   │   │   │   ├── [appId]/
    │   │   │   │   │   │   │   │   │   ├── authorizations/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # App authorizations
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/oauth
    │   │   │   │   │   │   │   │   │   │       # GET /v1/developer/oauth-apps/{app_id}/authorizations
    │   │   │   │   │   │   │   │   │   ├── credentials/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth credentials
    │   │   │   │   │   │   │   │   │   │       # - Client ID
    │   │   │   │   │   │   │   │   │   │       # - Client secret rotation
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/oauth
    │   │   │   │   │   │   │   │   │   │       # GET /v1/developer/oauth-apps/{app_id}/credentials
    │   │   │   │   │   │   │   │   │   │       # POST /v1/developer/oauth-apps/{app_id}/rotate-secret
    │   │   │   │   │   │   │   │   │   ├── scopes/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth scopes
    │   │   │   │   │   │   │   │   │   │       # BE: admin-be/oauth
    │   │   │   │   │   │   │   │   │   │       # GET /v1/developer/oauth-apps/{app_id}/scopes
    │   │   │   │   │   │   │   │   │   │       # PUT /v1/developer/oauth-apps/{app_id}/scopes
    │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth app detail
    │   │   │   │   │   │   │   │   │       # BE: admin-be/oauth
    │   │   │   │   │   │   │   │   │       # GET /v1/developer/oauth-apps/{app_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # OAuth apps
    │   │   │   │   │   │   │   │       # BE: admin-be/oauth
    │   │   │   │   │   │   │   │       # GET /v1/developer/oauth-apps
    │   │   │   │   │   │   │   │       # POST /v1/developer/oauth-apps
    │   │   │   │   │   │   │   ├── sandbox/
    │   │   │   │   │   │   │   │   ├── environments/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Sandbox environments
    │   │   │   │   │   │   │   │   │       # BE: admin-be/sandbox
    │   │   │   │   │   │   │   │   │       # GET /v1/developer/sandbox/environments
    │   │   │   │   │   │   │   │   │       # POST /v1/developer/sandbox/environments
    │   │   │   │   │   │   │   │   └── test-data/
    │   │   │   │   │   │   │   │       └── page.tsx  # Test data management
    │   │   │   │   │   │   │   │           # BE: admin-be/sandbox
    │   │   │   │   │   │   │   │           # GET /v1/developer/sandbox/test-data
    │   │   │   │   │   │   │   │           # POST /v1/developer/sandbox/test-data
    │   │   │   │   │   │   │   └── webhooks/
    │   │   │   │   │   │   │       ├── [webhookId]/
    │   │   │   │   │   │   │       │   ├── deliveries/
    │   │   │   │   │   │   │       │   │   ├── [deliveryId]/
    │   │   │   │   │   │   │       │   │   │   └── page.tsx  # Delivery detail
    │   │   │   │   │   │   │       │   │   │       # BE: admin-be/webhooks
    │   │   │   │   │   │   │       │   │   │       # GET /v1/developer/webhooks/{webhook_id}/deliveries/{delivery_id}
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Webhook deliveries
    │   │   │   │   │   │   │       │   │       # - Delivery history
    │   │   │   │   │   │   │       │   │       # - Retry failed deliveries
    │   │   │   │   │   │   │       │   │       # BE: admin-be/webhooks
    │   │   │   │   │   │   │       │   │       # GET /v1/developer/webhooks/{webhook_id}/deliveries
    │   │   │   │   │   │   │       │   ├── events/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Webhook event config
    │   │   │   │   │   │   │       │   │       # BE: admin-be/webhooks
    │   │   │   │   │   │   │       │   │       # GET /v1/developer/webhooks/{webhook_id}/events
    │   │   │   │   │   │   │       │   │       # PUT /v1/developer/webhooks/{webhook_id}/events
    │   │   │   │   │   │   │       │   ├── test/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Test webhook
    │   │   │   │   │   │   │       │   │       # BE: admin-be/webhooks
    │   │   │   │   │   │   │       │   │       # POST /v1/developer/webhooks/{webhook_id}/test
    │   │   │   │   │   │   │       │   └── page.tsx  # Webhook detail
    │   │   │   │   │   │   │       │       # BE: admin-be/webhooks
    │   │   │   │   │   │   │       │       # GET /v1/developer/webhooks/{webhook_id}
    │   │   │   │   │   │   │       └── page.tsx  # Webhooks management
    │   │   │   │   │   │   │           # BE: admin-be/webhooks
    │   │   │   │   │   │   │           # GET /v1/developer/webhooks
    │   │   │   │   │   │   │           # POST /v1/developer/webhooks
    │   │   │   │   │   │   ├── labs/
    │   │   │   │   │   │   │   └── page.tsx  # Experimental features
    │   │   │   │   │   │   │       # - Beta features toggle
    │   │   │   │   │   │   │       # - A/B test participation
    │   │   │   │   │   │   │       # BE: admin-be/feature-flags
    │   │   │   │   │   │   │       # GET /v1/settings/labs
    │   │   │   │   │   │   │       # PUT /v1/settings/labs
    │   │   │   │   │   │   └── security/
    │   │   │   │   │   │       ├── devices/
    │   │   │   │   │   │       │   └── page.tsx  # Devices & sessions
    │   │   │   │   │   │       │       # - Active sessions, revoke device
    │   │   │   │   │   │       │       # BE: users-be/session
    │   │   │   │   │   │       │       # GET/DELETE /v1/sessions
    │   │   │   │   │   │       ├── mfa/
    │   │   │   │   │   │       │   └── page.tsx  # MFA settings
    │   │   │   │   │   │       │       # - Enroll/disable TOTP/SMS
    │   │   │   │   │   │       │       # BE: users-be/mfa
    │   │   │   │   │   │       │       # GET/POST/DELETE /v1/mfa
    │   │   │   │   │   │       └── password/
    │   │   │   │   │   │           └── page.tsx  # Change password
    │   │   │   │   │   │               # - Update password
    │   │   │   │   │   │               # BE: users-be/account
    │   │   │   │   │   │               # POST /v1/account/password
    │   │   │   │   │   ├── status/
    │   │   │   │   │   │   ├── current/
    │   │   │   │   │   │   │   └── page.tsx  # Current system status
    │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │       # GET /v1/status/current
    │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   └── page.tsx  # Status history
    │   │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │   │       # GET /v1/status/history
    │   │   │   │   │   │   ├── subscribe/
    │   │   │   │   │   │   │   └── page.tsx  # Subscribe to status updates
    │   │   │   │   │   │   │       # BE: communications-be
    │   │   │   │   │   │   │       # POST /v1/notifications/status-subscribe
    │   │   │   │   │   │   └── page.tsx  # Status page (public)
    │   │   │   │   │   │       # - Current incidents & components
    │   │   │   │   │   │       # BE: utility/status
    │   │   │   │   │   │       # GET /v1/status
    │   │   │   │   │   ├── teams/
    │   │   │   │   │   │   ├── [teamId]/
    │   │   │   │   │   │   │   ├── capacity/
    │   │   │   │   │   │   │   │   ├── forecasting/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Capacity forecasting
    │   │   │   │   │   │   │   │   │       # BE: users-be/team
    │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/capacity/forecast
    │   │   │   │   │   │   │   │   ├── planning/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Capacity planning
    │   │   │   │   │   │   │   │   │       # - Resource allocation
    │   │   │   │   │   │   │   │   │       # - Utilization forecasting
    │   │   │   │   │   │   │   │   │       # BE: users-be/team, contracts-be
    │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/capacity/planning
    │   │   │   │   │   │   │   │   └── utilization/
    │   │   │   │   │   │   │   │       └── page.tsx  # Utilization tracking
    │   │   │   │   │   │   │   │           # - Current utilization
    │   │   │   │   │   │   │   │           # - Historical trends
    │   │   │   │   │   │   │   │           # BE: users-be/team
    │   │   │   │   │   │   │   │           # GET /v1/teams/{team_id}/capacity/utilization
    │   │   │   │   │   │   │   ├── hierarchy/
    │   │   │   │   │   │   │   │   ├── org-chart/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Organization chart
    │   │   │   │   │   │   │   │   │       # - Visual hierarchy
    │   │   │   │   │   │   │   │   │       # - Drag-and-drop reorg
    │   │   │   │   │   │   │   │   │       # BE: users-be/team
    │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/hierarchy
    │   │   │   │   │   │   │   │   ├── reporting-lines/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Reporting structure
    │   │   │   │   │   │   │   │   │       # BE: users-be/team
    │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/reporting-lines
    │   │   │   │   │   │   │   │   └── page.tsx  # Team hierarchy overview
    │   │   │   │   │   │   │   │       # BE: users-be/team
    │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/hierarchy
    │   │   │   │   │   │   │   ├── knowledge-base/
    │   │   │   │   │   │   │   │   ├── articles/
    │   │   │   │   │   │   │   │   │   ├── [articleId]/
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit article
    │   │   │   │   │   │   │   │   │   │   │       # BE: users-be/knowledge-base
    │   │   │   │   │   │   │   │   │   │   │       # PUT /v1/teams/{team_id}/kb/articles/{article_id}
    │   │   │   │   │   │   │   │   │   │   ├── versions/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Article versions
    │   │   │   │   │   │   │   │   │   │   │       # BE: users-be/knowledge-base
    │   │   │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/kb/articles/{article_id}/versions
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Article detail
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/knowledge-base
    │   │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/kb/articles/{article_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Articles list
    │   │   │   │   │   │   │   │   │       # BE: users-be/knowledge-base
    │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/kb/articles
    │   │   │   │   │   │   │   │   ├── categories/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # KB categories
    │   │   │   │   │   │   │   │   │       # BE: users-be/knowledge-base
    │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/kb/categories
    │   │   │   │   │   │   │   │   └── search/
    │   │   │   │   │   │   │   │       └── page.tsx  # KB search
    │   │   │   │   │   │   │   │           # BE: users-be/knowledge-base, search-be
    │   │   │   │   │   │   │   │           # GET /v1/teams/{team_id}/kb/search
    │   │   │   │   │   │   │   ├── policies/
    │   │   │   │   │   │   │   │   ├── [policyId]/
    │   │   │   │   │   │   │   │   │   ├── attestations/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Policy attestations
    │   │   │   │   │   │   │   │   │   │       # - Track acknowledgments
    │   │   │   │   │   │   │   │   │   │       # - Compliance tracking
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/policy
    │   │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/policies/{policy_id}/attestations
    │   │   │   │   │   │   │   │   │   ├── versions/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Policy versions
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/policy
    │   │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/policies/{policy_id}/versions
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Policy detail
    │   │   │   │   │   │   │   │   │       # BE: users-be/policy
    │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/policies/{policy_id}
    │   │   │   │   │   │   │   │   └── page.tsx  # Policies list
    │   │   │   │   │   │   │   │       # BE: users-be/policy
    │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/policies
    │   │   │   │   │   │   │   ├── procurement/
    │   │   │   │   │   │   │   │   ├── contracts/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Procurement contracts
    │   │   │   │   │   │   │   │   │       # BE: users-be/procurement, contracts-be
    │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/procurement/contracts
    │   │   │   │   │   │   │   │   ├── requests/
    │   │   │   │   │   │   │   │   │   ├── [requestId]/
    │   │   │   │   │   │   │   │   │   │   ├── approval/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve procurement
    │   │   │   │   │   │   │   │   │   │   │       # BE: users-be/procurement
    │   │   │   │   │   │   │   │   │   │   │       # POST /v1/teams/{team_id}/procurement/requests/{request_id}/approve
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Request detail
    │   │   │   │   │   │   │   │   │   │       # BE: users-be/procurement
    │   │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/procurement/requests/{request_id}
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Procurement requests
    │   │   │   │   │   │   │   │   │       # BE: users-be/procurement
    │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/procurement/requests
    │   │   │   │   │   │   │   │   │       # POST /v1/teams/{team_id}/procurement/requests
    │   │   │   │   │   │   │   │   └── vendors/
    │   │   │   │   │   │   │   │       ├── evaluation/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Vendor evaluation
    │   │   │   │   │   │   │   │       │       # BE: users-be/procurement
    │   │   │   │   │   │   │   │       │       # GET /v1/teams/{team_id}/procurement/vendors/evaluation
    │   │   │   │   │   │   │   │       ├── preferred/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Preferred vendors
    │   │   │   │   │   │   │   │       │       # BE: users-be/procurement
    │   │   │   │   │   │   │   │       │       # GET /v1/teams/{team_id}/procurement/vendors/preferred
    │   │   │   │   │   │   │   │       └── page.tsx  # Vendor management
    │   │   │   │   │   │   │   │           # BE: users-be/procurement
    │   │   │   │   │   │   │   │           # GET /v1/teams/{team_id}/procurement/vendors
    │   │   │   │   │   │   │   ├── training/
    │   │   │   │   │   │   │   │   ├── certifications/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Team certifications
    │   │   │   │   │   │   │   │   │       # - Track certifications
    │   │   │   │   │   │   │   │   │       # - Expiration alerts
    │   │   │   │   │   │   │   │   │       # BE: users-be/training
    │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/training/certifications
    │   │   │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Training compliance
    │   │   │   │   │   │   │   │   │       # - Completion tracking
    │   │   │   │   │   │   │   │   │       # - Compliance reports
    │   │   │   │   │   │   │   │   │       # BE: users-be/training
    │   │   │   │   │   │   │   │   │       # GET /v1/teams/{team_id}/training/compliance
    │   │   │   │   │   │   │   │   └── programs/
    │   │   │   │   │   │   │   │       ├── [programId]/
    │   │   │   │   │   │   │   │       │   ├── enroll/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Enroll in program
    │   │   │   │   │   │   │   │       │   │       # BE: users-be/training
    │   │   │   │   │   │   │   │       │   │       # POST /v1/teams/{team_id}/training/programs/{program_id}/enroll
    │   │   │   │   │   │   │   │       │   ├── progress/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Training progress
    │   │   │   │   │   │   │   │       │   │       # BE: users-be/training
    │   │   │   │   │   │   │   │       │   │       # GET /v1/teams/{team_id}/training/programs/{program_id}/progress
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Program detail
    │   │   │   │   │   │   │   │       │       # BE: users-be/training
    │   │   │   │   │   │   │   │       │       # GET /v1/teams/{team_id}/training/programs/{program_id}
    │   │   │   │   │   │   │   │       └── page.tsx  # Training programs
    │   │   │   │   │   │   │   │           # BE: users-be/training
    │   │   │   │   │   │   │   │           # GET /v1/teams/{team_id}/training/programs
    │   │   │   │   │   │   │   └── workflows/
    │   │   │   │   │   │   │       ├── [workflowId]/
    │   │   │   │   │   │   │       │   ├── analytics/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Workflow analytics
    │   │   │   │   │   │   │       │   │       # - Completion rates
    │   │   │   │   │   │   │       │   │       # - Bottlenecks
    │   │   │   │   │   │   │       │   │       # BE: users-be/workflow
    │   │   │   │   │   │   │       │   │       # GET /v1/teams/{team_id}/workflows/{workflow_id}/analytics
    │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit workflow
    │   │   │   │   │   │   │       │   │       # BE: users-be/workflow
    │   │   │   │   │   │   │       │   │       # PUT /v1/teams/{team_id}/workflows/{workflow_id}
    │   │   │   │   │   │   │       │   └── page.tsx  # Workflow detail
    │   │   │   │   │   │   │       │       # BE: users-be/workflow
    │   │   │   │   │   │   │       │       # GET /v1/teams/{team_id}/workflows/{workflow_id}
    │   │   │   │   │   │   │       └── page.tsx  # Workflows list
    │   │   │   │   │   │   │           # BE: users-be/workflow
    │   │   │   │   │   │   │           # GET /v1/teams/{team_id}/workflows
    │   │   │   │   │   │   └── cross-team/
    │   │   │   │   │   │       ├── benchmarking/
    │   │   │   │   │   │       │   └── page.tsx  # Inter-team benchmarking
    │   │   │   │   │   │       │       # - Performance comparison
    │   │   │   │   │   │       │       # - Best practices sharing
    │   │   │   │   │   │       │       # BE: users-be/team
    │   │   │   │   │   │       │       # GET /v1/teams/benchmarking
    │   │   │   │   │   │       └── collaboration/
    │   │   │   │   │   │           └── page.tsx  # Cross-team collaboration
    │   │   │   │   │   │               # - Shared projects
    │   │   │   │   │   │               # - Resource sharing
    │   │   │   │   │   │               # BE: users-be/team
    │   │   │   │   │   │               # GET /v1/teams/collaboration
    │   │   │   │   │   ├── transparency/
    │   │   │   │   │   │   └── page.tsx  # Transparency report
    │   │   │   │   │   │       # - User statistics
    │   │   │   │   │   │       # - Moderation actions
    │   │   │   │   │   │       # - Government requests
    │   │   │   │   │   │       # BE: admin-be/reporting
    │   │   │   │   │   │       # GET /v1/public/transparency-report
    │   │   │   │   │   └── layout.tsx  # Root layout for locale
    │   │   │   │   │       # - HTML lang attribute (i18n)
    │   │   │   │   │       # - Head meta tags
    │   │   │   │   │       # - Body layout
    │   │   │   │   │       # - Font loading
    │   │   │   │   ├── │/
    │   │   │   │   │   ├── operations/
    │   │   │   │   │   │   └── search-quality/
    │   │   │   │   │   │       ├── boosts/
    │   │   │   │   │   │       │   └── page.tsx  # Search boosts
    │   │   │   │   │   │       │       # - Term boosts
    │   │   │   │   │   │       │       # - Document boosts
    │   │   │   │   │   │       │       # BE: search-be/admin
    │   │   │   │   │   │       │       # GET /v1/admin/search/boosts
    │   │   │   │   │   │       │       # POST /v1/admin/search/boosts
    │   │   │   │   │   │       ├── reindexing/
    │   │   │   │   │   │       │   ├── schedule/
    │   │   │   │   │   │       │   │   └── page.tsx  # Reindex scheduling
    │   │   │   │   │   │       │   │       # BE: search-be/admin
    │   │   │   │   │   │       │   │       # POST /v1/admin/search/reindex/schedule
    │   │   │   │   │   │       │   └── status/
    │   │   │   │   │   │       │       └── page.tsx  # Reindex status
    │   │   │   │   │   │       │           # BE: search-be/admin
    │   │   │   │   │   │       │           # GET /v1/admin/search/reindex/status
    │   │   │   │   │   │       ├── relevance/
    │   │   │   │   │   │       │   ├── metrics/
    │   │   │   │   │   │       │   │   └── page.tsx  # Relevance metrics
    │   │   │   │   │   │       │   │       # - NDCG scores
    │   │   │   │   │   │       │   │       # - Click-through rates
    │   │   │   │   │   │       │   │       # BE: search-be/admin
    │   │   │   │   │   │       │   │       # GET /v1/admin/search/relevance/metrics
    │   │   │   │   │   │       │   ├── testing/
    │   │   │   │   │   │       │   │   └── page.tsx  # Relevance testing
    │   │   │   │   │   │       │   │       # - A/B testing
    │   │   │   │   │   │       │   │       # - Test queries
    │   │   │   │   │   │       │   │       # BE: search-be/admin
    │   │   │   │   │   │       │   │       # POST /v1/admin/search/relevance/test
    │   │   │   │   │   │       │   └── tuning/
    │   │   │   │   │   │       │       └── page.tsx  # Relevance tuning
    │   │   │   │   │   │       │           # - Weights adjustment
    │   │   │   │   │   │       │           # - Field boosting
    │   │   │   │   │   │       │           # BE: search-be/admin
    │   │   │   │   │   │       │           # GET /v1/admin/search/relevance/tuning
    │   │   │   │   │   │       │           # PUT /v1/admin/search/relevance/tuning
    │   │   │   │   │   │       └── synonyms/
    │   │   │   │   │   │           ├── [synonymId]/
    │   │   │   │   │   │           │   └── page.tsx  # Synonym detail
    │   │   │   │   │   │           │       # BE: search-be/admin
    │   │   │   │   │   │           │       # GET /v1/admin/search/synonyms/{synonym_id}
    │   │   │   │   │   │           │       # PUT /v1/admin/search/synonyms/{synonym_id}
    │   │   │   │   │   │           └── page.tsx  # Synonym management
    │   │   │   │   │   │               # BE: search-be/admin
    │   │   │   │   │   │               # GET /v1/admin/search/synonyms
    │   │   │   │   │   │               # POST /v1/admin/search/synonyms
    │   │   │   │   │   ├── trust-safety/
    │   │   │   │   │   │   ├── content-moderation/
    │   │   │   │   │   │   │   ├── appeals/
    │   │   │   │   │   │   │   │   ├── [appealId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Appeal review
    │   │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/moderation/appeals/{appeal_id}
    │   │   │   │   │   │   │   │   │       # POST /v1/admin/moderation/appeals/{appeal_id}/decide
    │   │   │   │   │   │   │   │   └── page.tsx  # Appeals queue
    │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │       # GET /v1/admin/moderation/appeals
    │   │   │   │   │   │   │   ├── automation/
    │   │   │   │   │   │   │   │   ├── actions/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Automated actions
    │   │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/moderation/automation/actions
    │   │   │   │   │   │   │   │   └── rules/
    │   │   │   │   │   │   │   │       └── page.tsx  # Auto-moderation rules
    │   │   │   │   │   │   │   │           # BE: admin-be/moderation
    │   │   │   │   │   │   │   │           # GET /v1/admin/moderation/automation/rules
    │   │   │   │   │   │   │   │           # POST /v1/admin/moderation/automation/rules
    │   │   │   │   │   │   │   ├── ml-assistance/
    │   │   │   │   │   │   │   │   ├── accuracy/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Model accuracy
    │   │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/moderation/ml-accuracy
    │   │   │   │   │   │   │   │   ├── predictions/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # ML predictions
    │   │   │   │   │   │   │   │   │       # - Auto-classification
    │   │   │   │   │   │   │   │   │       # - Confidence scores
    │   │   │   │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   │   │   │   │       # GET /v1/admin/moderation/ml-predictions
    │   │   │   │   │   │   │   │   └── training/
    │   │   │   │   │   │   │   │       └── page.tsx  # Model training
    │   │   │   │   │   │   │   │           # - Feedback loop
    │   │   │   │   │   │   │   │           # - Model retraining
    │   │   │   │   │   │   │   │           # BE: admin-be/moderation
    │   │   │   │   │   │   │   │           # POST /v1/admin/moderation/ml-training
    │   │   │   │   │   │   │   └── queue/
    │   │   │   │   │   │   │       ├── categories/
    │   │   │   │   │   │   │       │   └── [category]/
    │   │   │   │   │   │   │       │       └── page.tsx  # Category-specific queue
    │   │   │   │   │   │   │       │           # BE: admin-be/moderation
    │   │   │   │   │   │   │       │           # GET /v1/admin/moderation/queue?category={category}
    │   │   │   │   │   │   │       ├── priority/
    │   │   │   │   │   │   │       │   └── page.tsx  # Priority queue
    │   │   │   │   │   │   │       │       # BE: admin-be/moderation
    │   │   │   │   │   │   │       │       # GET /v1/admin/moderation/queue?priority=high
    │   │   │   │   │   │   │       └── page.tsx  # Moderation queue
    │   │   │   │   │   │   │           # BE: admin-be/moderation
    │   │   │   │   │   │   │           # GET /v1/admin/moderation/queue
    │   │   │   │   │   │   ├── risk-scoring/
    │   │   │   │   │   │   │   ├── models/
    │   │   │   │   │   │   │   │   └── page.tsx  # Risk scoring models
    │   │   │   │   │   │   │   │       # - User risk scores
    │   │   │   │   │   │   │   │       # - Transaction risk scores
    │   │   │   │   │   │   │   │       # BE: admin-be/risk-scoring
    │   │   │   │   │   │   │   │       # GET /v1/admin/risk-scoring/models
    │   │   │   │   │   │   │   ├── monitoring/
    │   │   │   │   │   │   │   │   └── page.tsx  # Risk monitoring
    │   │   │   │   │   │   │   │       # BE: admin-be/risk-scoring
    │   │   │   │   │   │   │   │       # GET /v1/admin/risk-scoring/monitoring
    │   │   │   │   │   │   │   └── thresholds/
    │   │   │   │   │   │   │       └── page.tsx  # Risk thresholds
    │   │   │   │   │   │   │           # BE: admin-be/risk-scoring
    │   │   │   │   │   │   │           # GET /v1/admin/risk-scoring/thresholds
    │   │   │   │   │   │   │           # PUT /v1/admin/risk-scoring/thresholds
    │   │   │   │   │   │   └── watchlists/
    │   │   │   │   │   │       ├── custom/
    │   │   │   │   │   │       │   └── page.tsx  # Custom watchlists
    │   │   │   │   │   │       │       # - Internal blacklist
    │   │   │   │   │   │       │       # - High-risk users
    │   │   │   │   │   │       │       # BE: admin-be/watchlists
    │   │   │   │   │   │       │       # GET /v1/admin/watchlists/custom
    │   │   │   │   │   │       │       # POST /v1/admin/watchlists/custom
    │   │   │   │   │   │       ├── global/
    │   │   │   │   │   │       │   └── page.tsx  # Global watchlists
    │   │   │   │   │   │       │       # - OFAC
    │   │   │   │   │   │       │       # - UN sanctions
    │   │   │   │   │   │       │       # - PEP lists
    │   │   │   │   │   │       │       # BE: admin-be/watchlists
    │   │   │   │   │   │       │       # GET /v1/admin/watchlists/global
    │   │   │   │   │   │       └── monitoring/
    │   │   │   │   │   │           └── page.tsx  # Watchlist monitoring
    │   │   │   │   │   │               # - Screening results
    │   │   │   │   │   │               # - False positives
    │   │   │   │   │   │               # BE: admin-be/watchlists
    │   │   │   │   │   │               # GET /v1/admin/watchlists/monitoring
    │   │   │   │   │   └── │  # conflict: file also existed
    │   │   │   │   └── │
    │   │   │   └── features/
    │   │   │       ├── budget/
    │   │   │       │   └── api/
    │   │   │       │       └── mutations.ts  # Budget mutations
    │   │   │       │           # - Create/update budgets
    │   │   │       │           # - Invalidate ['budget','list']
    │   │   │       │           # BE: financial-be/budget
    │   │   │       │           # POST /v1/budgets
    │   │   │       ├── interviews/
    │   │   │       │   └── api/
    │   │   │       │       └── mutations.ts  # Interview scheduling
    │   │   │       │           # - Create/reschedule/cancel
    │   │   │       │           # BE: proposals-be/interview
    │   │   │       │           # POST /v1/proposals/{id}/interviews
    │   │   │       ├── moderation/
    │   │   │       │   └── api/
    │   │   │       │       └── mutations.ts  # Moderation actions
    │   │   │       │           # - Warn / Suspend / Ban users
    │   │   │       │           # BE: users-be/moderation
    │   │   │       │           # POST /v1/admin/warning
    │   │   │       │           # POST /v1/admin/suspension
    │   │   │       │           # POST /v1/admin/ban
    │   │   │       ├── notifications/
    │   │   │       │   └── api/
    │   │   │       │       ├── mutations.ts  # Notifications mutations
    │   │   │       │       │   # - Mark read / mark all
    │   │   │       │       │   # BE: communications-be/notification
    │   │   │       │       │   # POST /v1/notifications/{id}/read
    │   │   │       │       └── queries.ts  # Notifications queries
    │   │   │       │           # - List notifications
    │   │   │       │           # BE: communications-be/notification
    │   │   │       │           # GET /v1/notifications?mine=1
    │   │   │       ├── offers/
    │   │   │       │   └── api/
    │   │   │       │       └── mutations.ts  # Offer create/respond
    │   │   │       │           # - Create, accept, decline
    │   │   │       │           # BE: contracts-be/offer (proposed if missing)
    │   │   │       │           # POST /v1/offers; POST /v1/offers/{id}/accept|decline
    │   │   │       ├── reviews/
    │   │   │       │   ├── api/
    │   │   │       │   │   ├── mutations.ts  # Review mutations
    │   │   │       │   │   │   # - Submit / edit review
    │   │   │       │   │   │   # - Invalidate ['reviews','list',subjectId]
    │   │   │       │   │   │   # BE: reviews-be/review
    │   │   │       │   │   │   # POST /v1/reviews
    │   │   │       │   │   └── queries.ts  # Review queries
    │   │   │       │   │       # - List reviews for subject
    │   │   │       │   │       # BE: reviews-be/review
    │   │   │       │   │       # GET /v1/reviews?subject_id=…
    │   │   │       │   └── components/
    │   │   │       │       └── ReviewForm.tsx  # Review form (presentational)
    │   │   │       │           # - Consumes typed props only
    │   │   │       │           # BE: none (presentational)
    │   │   │       ├── search/
    │   │   │       │   └── api/
    │   │   │       │       ├── alerts.ts  # Search alerts API
    │   │   │       │       │   # BE: search-be/alert (proposed if missing)
    │   │   │       │       │   # GET/POST/DELETE /v1/search/alerts
    │   │   │       │       └── saved.ts  # Saved queries API
    │   │   │       │           # BE: search-be/saved_query (proposed if missing)
    │   │   │       │           # GET/POST/DELETE /v1/search/saved
    │   │   │       ├── subscriptions/
    │   │   │       │   └── api/
    │   │   │       │       ├── mutations.ts  # Subscription mutations
    │   │   │       │       │   # - Change plan / Cancel
    │   │   │       │       │   # - Invalidate ['subscriptions','me']
    │   │   │       │       │   # BE: financial-be/subscription
    │   │   │       │       │   # POST /v1/subscriptions/change
    │   │   │       │       │   # POST /v1/subscriptions/cancel
    │   │   │       │       └── queries.ts  # Subscription queries
    │   │   │       │           # - Current subscription
    │   │   │       │           # - Invoices history
    │   │   │       │           # BE: financial-be/subscription
    │   │   │       │           # GET /v1/subscriptions/me
    │   │   │       │           # GET /v1/subscriptions/{id}/invoices
    │   │   │       └── support/
    │   │   │           └── api/
    │   │   │               ├── mutations.ts  # Support ticket mutations
    │   │   │               │   # - Create ticket, add message, attach files
    │   │   │               │   # - Invalidate ['support','tickets']
    │   │   │               │   # BE: admin-be/support_ticket
    │   │   │               │   # POST /v1/support/tickets
    │   │   │               │   # POST /v1/support/tickets/{id}/messages
    │   │   │               │   # BE: storage-be/asset (uploads)
    │   │   │               │   # POST /v1/storage/uploads (signed URL) → PUT file → POST /v1/storage/commit
    │   │   │               └── queries.ts  # Support ticket queries
    │   │   │                   # - List & detail
    │   │   │                   # BE: admin-be/support_ticket
    │   │   │                   # GET /v1/support/tickets
    │   │   │                   # GET /v1/support/tickets/{id}
    │   │   └── middleware.ts  # Next.js middleware
    │   │       # - CSP headers
    │   │       # - CORS configuration
    │   │       # - Rate limiting
    │   │       # - Security headers (X-Frame-Options, etc.)
    │   │       # - Security headers
    │   │       # - CSRF protection
    │   │       # - Rate limiting (client-side)
    │   │       # - Auth checks
    │   │       # Next.js middleware for security headers
    │   │       # Next.js middleware for security
    │   └── │
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
    │   ├── .env.development
    │   ├── .env.production
    │   ├── .env.staging
    │   ├── feature-flags.ts  # Feature flags config
    │   └── index.ts  # Config loader
    ├── deploy/  # Deployment configurations
    │   ├── docker/
    │   │   ├── mobile-build.Dockerfile  # EAS build container
    │   │   ├── nginx.conf  # NGINX config for static serving
    │   │   ├── web.Dockerfile  # Multi-stage Next.js build
    │   │   └── web.dockerignore
    │   └── k8s/  # Kubernetes manifests
    │       ├── mobile-api/  # Mobile API gateway (if separate)
    │       ├── monitoring/
    │       │   ├── grafana-dashboard.yaml  # Dashboards
    │       │   └── prometheus.yaml  # Metrics scraping
    │       └── web/
    │           ├── configmap.yaml  # Environment config
    │           ├── deployment.yaml  # Web app deployment
    │           ├── hpa.yaml  # Horizontal Pod Autoscaler
    │           ├── ingress.yaml  # NGINX ingress
    │           └── service.yaml  # ClusterIP service
    ├── docs/  # Documentation
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
    │   │   ├── microservices/
    │   │   │   ├── admin-be.md
    │   │   │   ├── communications-be.md
    │   │   │   ├── contracts-be.md
    │   │   │   ├── financial-be.md
    │   │   │   ├── jobs-be.md
    │   │   │   ├── proposals-be.md
    │   │   │   ├── reviews-be.md
    │   │   │   ├── search-be.md
    │   │   │   ├── storage-be.md
    │   │   │   ├── subscriptions-be.md
    │   │   │   └── users-be.md
    │   │   ├── authentication.md  # API authentication
    │   │   ├── error-handling.md  # Error handling
    │   │   ├── errors.md
    │   │   ├── getting-started.md
    │   │   ├── introduction.md  # API integration guide
    │   │   ├── rate-limiting.md  # Rate limiting
    │   │   └── webhooks.md
    │   ├── api-integration/
    │   │   ├── authentication/
    │   │   │   ├── keycloak-integration.md
    │   │   │   ├── refresh-flow.md
    │   │   │   └── token-management.md
    │   │   ├── best-practices/
    │   │   │   ├── api-client-patterns.md
    │   │   │   ├── caching-strategies.md
    │   │   │   └── rate-limiting.md
    │   │   ├── error-handling/
    │   │   │   ├── error-codes.md
    │   │   │   ├── fallback-mechanisms.md
    │   │   │   └── retry-strategies.md
    │   │   ├── microservices/
    │   │   │   ├── admin-be-integration.md
    │   │   │   ├── communications-be-integration.md
    │   │   │   ├── contracts-be-integration.md
    │   │   │   ├── financial-be-integration.md
    │   │   │   ├── jobs-be-integration.md
    │   │   │   ├── proposals-be-integration.md
    │   │   │   ├── reviews-be-integration.md
    │   │   │   ├── search-be-integration.md
    │   │   │   ├── storage-be-integration.md
    │   │   │   ├── subscriptions-be-integration.md
    │   │   │   └── users-be-integration.md
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
    │   │   ├── authentication.md  # Auth flow
    │   │   ├── data-fetching-patterns.md
    │   │   ├── data-fetching.md  # Data fetching patterns
    │   │   ├── frontend-architecture.md  # FE architecture details
    │   │   ├── microservices-integration.md
    │   │   ├── overview.md  # System overview
    │   │   │   # System architecture
    │   │   ├── performance.md  # Performance optimization
    │   │   ├── routing-strategy.md
    │   │   ├── routing.md  # Routing strategy
    │   │   └── state-management.md  # State management patterns
    │   ├── components/
    │   │   ├── atoms/
    │   │   │   ├── ...
    │   │   │   ├── button.md
    │   │   │   └── input.md
    │   │   ├── design-system/
    │   │   │   ├── breakpoints.md
    │   │   │   ├── color-system.md
    │   │   │   ├── design-tokens.md
    │   │   │   ├── spacing.md
    │   │   │   └── typography.md
    │   │   ├── molecules/
    │   │   │   ├── ...
    │   │   │   ├── card.md
    │   │   │   └── form-field.md
    │   │   ├── organisms/
    │   │   │   ├── ...
    │   │   │   ├── header.md
    │   │   │   └── sidebar.md
    │   │   ├── storybook/
    │   │   │   ├── setup.md
    │   │   │   ├── testing-stories.md
    │   │   │   └── writing-stories.md
    │   │   ├── usage-guidelines/
    │   │   │   ├── accessibility.md
    │   │   │   ├── component-composition.md
    │   │   │   ├── responsive-design.md
    │   │   │   └── theming-guide.md
    │   │   ├── accessibility.md
    │   │   ├── component-library.md
    │   │   ├── design-system.md
    │   │   ├── design-tokens.md
    │   │   ├── overview.md  # Component library overview
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
    │   │   ├── setup/
    │   │   │   ├── environment-variables.md
    │   │   │   ├── local-development.md
    │   │   │   └── troubleshooting.md
    │   │   ├── contributing.md
    │   │   ├── deployment.md
    │   │   ├── development-workflow.md
    │   │   ├── development.md
    │   │   ├── getting-started.md
    │   │   ├── testing-guide.md
    │   │   ├── testing.md
    │   │   └── troubleshooting.md
    │   ├── ARCHITECTURE.md  # System architecture
    │   ├── CONTRIBUTING.md  # Contribution guidelines
    │   ├── DEPLOYMENT.md  # Deployment procedures
    │   ├── MICROSERVICES_MAPPING.md  # BE microservices integration
    │   ├── PERFORMANCE.md  # Performance optimization guide
    │   ├── README.md  # Project overview
    │   ├── SETUP.md  # Development setup guide
    │   ├── STATE_MANAGEMENT.md  # TanStack Query + Zustand patterns
    │   └── TESTING.md  # Testing strategy
    ├── packages/  # Shared libraries
    │   ├── config/  # Shared configurations
    │   │   ├── eslint-config/
    │   │   │   ├── index.js  # Base ESLint config
    │   │   │   ├── next.js  # Next.js-specific rules
    │   │   │   ├── package.json
    │   │   │   └── react.js  # React-specific rules
    │   │   ├── tailwind-config/
    │   │   │   ├── themes/
    │   │   │   │   ├── dark.js
    │   │   │   │   └── light.js
    │   │   │   ├── index.js  # Base Tailwind config
    │   │   │   │   # - Design tokens
    │   │   │   │   # - Color palette
    │   │   │   │   # - Spacing scale
    │   │   │   │   # - Typography
    │   │   │   │   # - Breakpoints
    │   │   │   │   # - Animations
    │   │   │   └── package.json
    │   │   └── typescript-config/
    │   │       ├── base.json  # Base TS config
    │   │       ├── nextjs.json  # Next.js TS config
    │   │       ├── package.json
    │   │       └── react-native.json  # React Native TS config
    │   ├── shared/  # Business logic, hooks, utilities
    │   │   # Shared business logic, hooks, utilities
    │   │   ├── src/
    │   │   │   ├── accessibility/
    │   │   │   │   ├── testing/
    │   │   │   │   │   ├── a11y-test-utils.ts  # Testing utilities
    │   │   │   │   │   └── axe-config.ts  # axe-core configuration
    │   │   │   │   ├── a11y-utils.ts  # Accessibility utilities
    │   │   │   │   ├── aria-utils.ts  # ARIA utilities
    │   │   │   │   ├── focus-management.ts  # Focus management
    │   │   │   │   ├── keyboard-navigation.ts  # Keyboard navigation
    │   │   │   │   └── screen-reader.ts  # Screen reader utilities
    │   │   │   ├── activity/
    │   │   │   │   ├── activity-client.ts  # Activity API client
    │   │   │   │   │   # BE: utility/activity (if exists) or communications-be/notification
    │   │   │   │   │   # GET /v1/activity/feed
    │   │   │   │   │   # GET /v1/activity/user/{user_id}
    │   │   │   │   └── types.ts
    │   │   │   ├── analytics/
    │   │   │   │   ├── batching/
    │   │   │   │   │   ├── event-batcher.ts  # Event batching
    │   │   │   │   │   ├── flush-strategies.ts  # Flush strategies
    │   │   │   │   │   └── queue-manager.ts  # Queue management
    │   │   │   │   ├── collectors/
    │   │   │   │   │   ├── error-collector.ts  # Error tracking
    │   │   │   │   │   ├── interaction-collector.ts  # Interaction tracking
    │   │   │   │   │   ├── page-view-collector.ts  # Page view collection
    │   │   │   │   │   └── performance-collector.ts  # Performance metrics
    │   │   │   │   ├── enrichment/
    │   │   │   │   │   ├── context-enricher.ts  # Event context enrichment
    │   │   │   │   │   ├── session-enricher.ts  # Session data enrichment
    │   │   │   │   │   └── user-enricher.ts  # User data enrichment
    │   │   │   │   ├── events/
    │   │   │   │   │   ├── auth-events.ts  # Auth analytics
    │   │   │   │   │   ├── contract-events.ts  # Contract analytics
    │   │   │   │   │   ├── job-events.ts  # Job analytics
    │   │   │   │   │   ├── payment-events.ts  # Payment analytics
    │   │   │   │   │   ├── proposal-events.ts  # Proposal analytics
    │   │   │   │   │   └── user-events.ts  # User behavior
    │   │   │   │   ├── privacy/
    │   │   │   │   │   ├── consent-manager.ts  # Consent management
    │   │   │   │   │   ├── data-minimization.ts  # Data minimization
    │   │   │   │   │   └── pii-scrubber.ts  # PII removal
    │   │   │   │   ├── sampling/
    │   │   │   │   │   ├── strategies/
    │   │   │   │   │   │   ├── adaptive-sampling.ts
    │   │   │   │   │   │   ├── percentage-sampling.ts
    │   │   │   │   │   │   └── user-based-sampling.ts
    │   │   │   │   │   └── sampler.ts  # Event sampling
    │   │   │   │   ├── trackers/
    │   │   │   │   │   ├── conversion-tracker.ts
    │   │   │   │   │   ├── interaction-tracker.ts
    │   │   │   │   │   └── page-view-tracker.ts
    │   │   │   │   └── utils/
    │   │   │   │       ├── anonymize.ts  # PII anonymization
    │   │   │   │       └── event-builder.ts
    │   │   │   ├── api/
    │   │   │   │   ├── activity/
    │   │   │   │   │   ├── activity-client.ts  # Activity API client
    │   │   │   │   │   │   # BE: utility/activity (if exists) or communications-be/notification
    │   │   │   │   │   │   # GET /v1/activity/feed
    │   │   │   │   │   │   # GET /v1/activity/user/{user_id}
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── communications/
    │   │   │   │   │   ├── client.ts  # Communications API client
    │   │   │   │   │   │   # - Conversations, messages, notifications
    │   │   │   │   │   │   # BE: communications-be
    │   │   │   │   │   ├── conversations.ts  # /v1/conversations, /v1/messages
    │   │   │   │   │   └── notifications.ts  # /v1/notifications
    │   │   │   │   ├── compliance/
    │   │   │   │   │   ├── compliance-client.ts  # Compliance API client
    │   │   │   │   │   │   # BE: admin-be/privacy, admin-be/kyc_case
    │   │   │   │   │   │   # GET /v1/privacy/export-requests
    │   │   │   │   │   │   # POST /v1/privacy/export-requests/{id}/process
    │   │   │   │   │   │   # GET /v1/privacy/deletion-requests
    │   │   │   │   │   │   # POST /v1/privacy/deletion-requests/{id}/process
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── contracts/
    │   │   │   │   │   ├── client.ts  # Contracts API client
    │   │   │   │   │   │   # - Contract, SOW, milestones, diary, deliverables, disputes, escrow, offers
    │   │   │   │   │   │   # BE: contracts-be | financial-be/escrow
    │   │   │   │   │   ├── contracts.ts  # /v1/contracts
    │   │   │   │   │   ├── deliverables.ts  # /v1/contracts/{id}/deliverables
    │   │   │   │   │   ├── diary.ts  # /v1/contracts/{id}/work-diary|timesheets
    │   │   │   │   │   ├── disputes.ts  # /v1/contracts/{id}/disputes
    │   │   │   │   │   ├── escrow.ts  # /v1/escrow/{contractId}/fund|release
    │   │   │   │   │   ├── milestones.ts  # /v1/contracts/{id}/milestones
    │   │   │   │   │   └── offers.ts  # /v1/offers (proposed)
    │   │   │   │   ├── experiments/
    │   │   │   │   │   ├── experiments-client.ts  # Experiments API client
    │   │   │   │   │   │   # BE: utility/experiments (if exists) or utility/flags
    │   │   │   │   │   │   # GET /v1/experiments/active
    │   │   │   │   │   │   # POST /v1/experiments/{id}/track
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── financial/
    │   │   │   │   │   ├── client.ts  # Financial API client
    │   │   │   │   │   │   # - Base URL, interceptors, retries
    │   │   │   │   │   │   # BE: financial-be (all)
    │   │   │   │   │   │   # Propagates trace & idempotency keys
    │   │   │   │   │   ├── invoices.ts  # Invoices endpoints
    │   │   │   │   │   │   # - GET /v1/invoices; POST /v1/invoices/{id}/pay
    │   │   │   │   │   ├── payouts.ts  # Payouts endpoints
    │   │   │   │   │   │   # - GET/POST /v1/payouts
    │   │   │   │   │   ├── refunds.ts  # Refunds endpoints
    │   │   │   │   │   │   # - GET/POST /v1/refunds
    │   │   │   │   │   ├── subscription.ts  # Subscription endpoints
    │   │   │   │   │   │   # - GET /v1/subscriptions/me; POST /v1/subscriptions/change|cancel
    │   │   │   │   │   └── wallet.ts  # Wallet endpoints
    │   │   │   │   │       # - GET /v1/wallet; GET /v1/wallet/transactions
    │   │   │   │   ├── flags/
    │   │   │   │   │   ├── flags-client.ts  # Feature flags API client
    │   │   │   │   │   │   # BE: utility/flags
    │   │   │   │   │   │   # GET /v1/flags/user
    │   │   │   │   │   │   # GET /v1/flags/organization
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── incidents/
    │   │   │   │   │   ├── incidents-client.ts  # Incidents API client
    │   │   │   │   │   │   # BE: utility/status
    │   │   │   │   │   │   # GET /v1/status/incidents
    │   │   │   │   │   │   # POST /v1/status/incidents
    │   │   │   │   │   │   # PUT /v1/status/incidents/{id}
    │   │   │   │   │   │   # GET /v1/status/maintenance
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── jobs/
    │   │   │   │   │   ├── client.ts  # Jobs API client
    │   │   │   │   │   │   # - Post/edit/publish jobs; invites
    │   │   │   │   │   │   # BE: jobs-be
    │   │   │   │   │   ├── invites.ts  # /v1/invites
    │   │   │   │   │   └── jobs.ts  # /v1/jobs + subroutes
    │   │   │   │   ├── learning/
    │   │   │   │   │   ├── learning-client.ts  # Learning API client
    │   │   │   │   │   │   # BE: learning-be (if exists) or external LMS
    │   │   │   │   │   │   # GET /v1/learning/courses
    │   │   │   │   │   │   # GET /v1/learning/courses/{id}
    │   │   │   │   │   │   # POST /v1/learning/assessments/{id}/submit
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── moderation/
    │   │   │   │   │   ├── moderation-client.ts  # Moderation API client
    │   │   │   │   │   │   # BE: admin-be/moderation (if exists)
    │   │   │   │   │   │   # POST /v1/moderation/report
    │   │   │   │   │   │   # POST /v1/moderation/content-check
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── presence/
    │   │   │   │   │   ├── presence-client.ts  # Presence API client
    │   │   │   │   │   │   # BE: communications-be/presence (if exists)
    │   │   │   │   │   │   # POST /v1/presence/heartbeat
    │   │   │   │   │   │   # GET /v1/presence/users/{id}
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── proposals/
    │   │   │   │   │   ├── client.ts  # Proposals API client
    │   │   │   │   │   │   # - Retries, error mapping
    │   │   │   │   │   │   # BE: proposals-be
    │   │   │   │   │   ├── interviews.ts  # Interviews scheduling
    │   │   │   │   │   │   # - POST /v1/proposals/{id}/interviews
    │   │   │   │   │   └── proposals.ts  # Proposals CRUD
    │   │   │   │   │       # - POST /v1/proposals; GET /v1/proposals?mine=1
    │   │   │   │   ├── search/
    │   │   │   │   │   ├── alerts.ts  # /v1/search/alerts (proposed)
    │   │   │   │   │   ├── client.ts  # Search API client
    │   │   │   │   │   │   # - Query/feed/portfolios + assist + saved/alerts
    │   │   │   │   │   │   # BE: search-be
    │   │   │   │   │   ├── feed.ts  # /v1/feed|trending|similarity
    │   │   │   │   │   ├── portfolios.ts  # /v1/search/portfolios
    │   │   │   │   │   ├── query.ts  # /v1/search
    │   │   │   │   │   └── saved.ts  # /v1/search/saved (proposed)
    │   │   │   │   ├── sourcing/
    │   │   │   │   │   ├── sourcing-client.ts  # Sourcing API client
    │   │   │   │   │   │   # BE: jobs-be/campaign (if exists) or jobs-be/job
    │   │   │   │   │   │   # GET /v1/sourcing/campaigns
    │   │   │   │   │   │   # POST /v1/sourcing/campaigns
    │   │   │   │   │   │   # GET /v1/sourcing/talent-pools
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── storage/
    │   │   │   │   │   ├── assets.ts  # /v1/storage/uploads|commit|assets
    │   │   │   │   │   └── client.ts  # Storage API client
    │   │   │   │   │       # - Signed upload, commit, list
    │   │   │   │   │       # BE: storage-be
    │   │   │   │   ├── users/
    │   │   │   │   │   ├── client.ts  # Users API client
    │   │   │   │   │   │   # - Org/team, profile, security, saved items
    │   │   │   │   │   │   # BE: users-be
    │   │   │   │   │   ├── org.ts  # /v1/orgs, /v1/invites
    │   │   │   │   │   ├── profile.ts  # /v1/profile, /v1/users/{handle}
    │   │   │   │   │   ├── saved.ts  # /v1/saved-items
    │   │   │   │   │   └── security.ts  # /v1/sessions, /v1/mfa
    │   │   │   │   └── webhooks/
    │   │   │   │       ├── types.ts
    │   │   │   │       └── webhooks-client.ts  # Webhooks API client
    │   │   │   │           # BE: utility/webhooks (if exists) or admin-be
    │   │   │   │           # GET /v1/webhooks
    │   │   │   │           # POST /v1/webhooks
    │   │   │   │           # PUT /v1/webhooks/{id}
    │   │   │   │           # DELETE /v1/webhooks/{id}
    │   │   │   ├── compliance/
    │   │   │   │   ├── compliance-client.ts  # Compliance API client
    │   │   │   │   │   # BE: admin-be/privacy, admin-be/kyc_case
    │   │   │   │   │   # GET /v1/privacy/export-requests
    │   │   │   │   │   # POST /v1/privacy/export-requests/{id}/process
    │   │   │   │   │   # GET /v1/privacy/deletion-requests
    │   │   │   │   │   # POST /v1/privacy/deletion-requests/{id}/process
    │   │   │   │   └── types.ts
    │   │   │   ├── experiments/
    │   │   │   │   ├── ab-testing/
    │   │   │   │   │   ├── bucketing.ts  # User bucketing
    │   │   │   │   │   ├── experiment-engine.ts  # A/B test engine
    │   │   │   │   │   │   # BE: admin-be/experiments
    │   │   │   │   │   │   # GET /v1/experiments/active
    │   │   │   │   │   │   # POST /v1/experiments/{id}/track
    │   │   │   │   │   ├── tracking.ts  # Experiment tracking
    │   │   │   │   │   └── variant-selection.ts  # Variant assignment
    │   │   │   │   ├── analytics/
    │   │   │   │   │   ├── conversion-tracking.ts  # Conversion tracking
    │   │   │   │   │   ├── metric-collection.ts  # Metric collection
    │   │   │   │   │   └── statistical-analysis.ts  # Statistical analysis
    │   │   │   │   ├── feature-variants/
    │   │   │   │   │   ├── rollout-strategies.ts  # Rollout strategies
    │   │   │   │   │   ├── targeting.ts  # User targeting
    │   │   │   │   │   └── variant-manager.ts  # Feature variants
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useExperiment.ts  # Experiment hook
    │   │   │   │   │   ├── useFeatureVariant.ts  # Feature variant hook
    │   │   │   │   │   └── useVariant.ts  # Variant hook
    │   │   │   │   ├── experiments-client.ts  # Experiments API client
    │   │   │   │   │   # BE: utility/experiments (if exists) or utility/flags
    │   │   │   │   │   # GET /v1/experiments/active
    │   │   │   │   │   # POST /v1/experiments/{id}/track
    │   │   │   │   └── types.ts
    │   │   │   ├── features/  # Feature-specific shared code
    │   │   │   │   ├── achievements/
    │   │   │   │   │   ├── definitions/
    │   │   │   │   │   │   ├── client-achievements.ts  # Client achievements
    │   │   │   │   │   │   ├── freelancer-achievements.ts  # Freelancer achievements
    │   │   │   │   │   │   └── platform-achievements.ts  # Platform-wide achievements
    │   │   │   │   │   └── utils.ts  # Achievement utilities
    │   │   │   │   ├── admin/  # Admin feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── admin-api.ts  # Admin API client
    │   │   │   │   │   │   │   # BE: admin-be/*
    │   │   │   │   │   │   ├── analytics-api.ts  # Admin analytics API
    │   │   │   │   │   │   │   # BE: admin-be/analytics
    │   │   │   │   │   │   ├── kyc-api.ts  # KYC management API
    │   │   │   │   │   │   │   # BE: admin-be/kyc-case
    │   │   │   │   │   │   └── moderation-api.ts  # Moderation API
    │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useAdminSession.ts  # JIT admin session hook
    │   │   │   │   │   │   ├── useAdminUsers.ts
    │   │   │   │   │   │   ├── useDisputes.ts
    │   │   │   │   │   │   ├── useKYCCases.ts
    │   │   │   │   │   │   ├── useKycCases.ts  # KYC cases management
    │   │   │   │   │   │   ├── useModeration.ts  # Moderation actions
    │   │   │   │   │   │   └── useSystemHealth.ts  # System health monitoring
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── admin-mutations.ts  # Admin mutations
    │   │   │   │   │   │   └── admin-queries.ts  # BE: admin-be
    │   │   │   │   │   │       # Admin queries
    │   │   │   │   │   ├── admin-api.ts
    │   │   │   │   │   └── types.ts  # Admin types
    │   │   │   │   ├── analytics/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── analytics-api.ts  # Analytics API
    │   │   │   │   │   │       # BE: utility/analytics
    │   │   │   │   │   │       # POST /v1/analytics/events
    │   │   │   │   │   │       # POST /v1/analytics/page-view
    │   │   │   │   │   │       # Analytics API client
    │   │   │   │   │   │       # BE: Multiple services (/analytics endpoints)
    │   │   │   │   │   ├── events/
    │   │   │   │   │   │   ├── contract-events.ts  # Contract events
    │   │   │   │   │   │   ├── job-events.ts  # Job events
    │   │   │   │   │   │   ├── payment-events.ts  # Payment events
    │   │   │   │   │   │   ├── proposal-events.ts  # Proposal events
    │   │   │   │   │   │   ├── system-events.ts  # System events
    │   │   │   │   │   │   └── user-events.ts  # User events
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useAnalytics.ts  # Track events
    │   │   │   │   │   │   ├── useContractAnalytics.ts  # Contract analytics
    │   │   │   │   │   │   ├── useConversionTracking.ts  # Track conversions
    │   │   │   │   │   │   ├── useJobAnalytics.ts  # Job analytics
    │   │   │   │   │   │   ├── usePageView.ts  # Track page views
    │   │   │   │   │   │   ├── useProfileAnalytics.ts  # Profile analytics
    │   │   │   │   │   │   └── useRevenueAnalytics.ts  # Revenue analytics
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   └── analytics-queries.ts  # Analytics queries
    │   │   │   │   │   ├── utils/
    │   │   │   │   │   │   ├── anonymize.ts  # Anonymize PII
    │   │   │   │   │   │   ├── batch-sender.ts  # Batch events
    │   │   │   │   │   │   └── event-builder.ts  # Build events
    │   │   │   │   │   └── types.ts  # Analytics types
    │   │   │   │   ├── auctions/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── auctions-api.ts  # Auctions API client
    │   │   │   │   │   │       # BE: proposals-be/auction
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useActiveAuctions.ts  # Active auctions list
    │   │   │   │   │   │   ├── useAuction.ts  # Single auction
    │   │   │   │   │   │   ├── useAuctionBid.ts  # Place bid
    │   │   │   │   │   │   └── useAuctionHistory.ts  # Bid history
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── auctions-mutations.ts  # Auction mutations
    │   │   │   │   │   │   └── auctions-queries.ts  # Auction queries
    │   │   │   │   │   ├── store/
    │   │   │   │   │   │   └── auction-store.ts  # Real-time auction state (Zustand)
    │   │   │   │   │   └── types.ts  # Auction types
    │   │   │   │   ├── auth/  # Authentication
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useAuth.ts  # Main auth hook
    │   │   │   │   │   │   │   # - isAuthenticated
    │   │   │   │   │   │   │   # - user
    │   │   │   │   │   │   │   # - login, logout
    │   │   │   │   │   │   │   # - refresh token
    │   │   │   │   │   │   ├── useKeycloak.ts  # Keycloak integration
    │   │   │   │   │   │   ├── usePermissions.ts  # RBAC permissions
    │   │   │   │   │   │   └── useSession.ts  # Session management
    │   │   │   │   │   ├── stores/
    │   │   │   │   │   │   └── auth-store.ts  # Zustand auth store
    │   │   │   │   │   │       # - user
    │   │   │   │   │   │       # - tokens
    │   │   │   │   │   │       # - login/logout actions
    │   │   │   │   │   │       # - Persisted to storage
    │   │   │   │   │   ├── utils/
    │   │   │   │   │   │   ├── rbac.ts  # RBAC utilities
    │   │   │   │   │   │   ├── session.ts  # Session utilities
    │   │   │   │   │   │   └── token.ts  # Token utilities
    │   │   │   │   │   └── types.ts  # Auth types
    │   │   │   │   ├── bidding/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── bid-api.ts  # Bid placement API
    │   │   │   │   │   │   │   # BE: proposals-be/bid
    │   │   │   │   │   │   └── bid-strategy-api.ts  # Bid strategy API
    │   │   │   │   │   │       # BE: proposals-be/bid-strategy
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useBidAnalytics.ts  # Bid analytics
    │   │   │   │   │   │   ├── useBidHistory.ts  # Bid history
    │   │   │   │   │   │   ├── useBidStrategies.ts  # List strategies
    │   │   │   │   │   │   ├── useBidStrategy.ts  # Bid strategy management
    │   │   │   │   │   │   └── usePlaceBid.ts  # Place bid
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── bidding-mutations.ts  # Bidding mutations
    │   │   │   │   │   │   └── bidding-queries.ts  # Bidding queries
    │   │   │   │   │   └── types.ts  # Bidding types
    │   │   │   │   ├── collaboration/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── collaboration-api.ts  # Collaboration API
    │   │   │   │   │   ├── components/
    │   │   │   │   │   │   ├── ActiveUsers/
    │   │   │   │   │   │   │   ├── ActiveUsers.tsx
    │   │   │   │   │   │   │   └── ActiveUsers.types.ts
    │   │   │   │   │   │   ├── CollaboratorCursor/
    │   │   │   │   │   │   │   ├── CollaboratorCursor.tsx
    │   │   │   │   │   │   │   └── CollaboratorCursor.types.ts
    │   │   │   │   │   │   └── PresenceIndicator/
    │   │   │   │   │   │       ├── PresenceIndicator.tsx
    │   │   │   │   │   │       └── PresenceIndicator.types.ts
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useCollaboration.ts  # Collaboration session
    │   │   │   │   │   │   ├── useCursors.ts  # Cursor tracking
    │   │   │   │   │   │   ├── usePresence.ts  # User presence
    │   │   │   │   │   │   └── useSharedState.ts  # Shared state sync
    │   │   │   │   │   ├── providers/
    │   │   │   │   │   │   └── CollaborationProvider.tsx  # Collab context
    │   │   │   │   │   └── types.ts  # Collaboration types
    │   │   │   │   ├── compliance/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── compliance-api.ts  # Compliance API
    │   │   │   │   │   │   │   # BE: users-be/compliance
    │   │   │   │   │   │   └── tax-profile-api.ts  # Tax profile API
    │   │   │   │   │   │       # BE: users-be/compliance, financial-be/tax
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useComplianceDocuments.ts  # Document management
    │   │   │   │   │   │   ├── useComplianceProfile.ts  # Compliance profile
    │   │   │   │   │   │   ├── useTaxProfile.ts  # Tax profile management
    │   │   │   │   │   │   └── useTaxReports.ts  # Tax reports
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── compliance-mutations.ts  # Compliance mutations
    │   │   │   │   │   │   └── compliance-queries.ts  # Compliance queries
    │   │   │   │   │   └── types.ts  # Compliance types
    │   │   │   │   ├── connects/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── connects-api.ts  # Connects API client
    │   │   │   │   │   │       # BE: proposals-be/connect
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useConnectPackages.ts  # Available packages
    │   │   │   │   │   │   ├── useConnectRefund.ts  # Request refund
    │   │   │   │   │   │   ├── useConnects.ts  # Connects balance and history
    │   │   │   │   │   │   └── usePurchaseConnects.ts  # Purchase connects
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── connects-mutations.ts  # Connect mutations
    │   │   │   │   │   │   └── connects-queries.ts  # Connect queries
    │   │   │   │   │   └── types.ts  # Connect types
    │   │   │   │   ├── contracts/  # Contracts feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── contracts-api.ts
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useContract.ts
    │   │   │   │   │   │   ├── useContracts.ts
    │   │   │   │   │   │   ├── useDeliverables.ts
    │   │   │   │   │   │   ├── useDisputes.ts
    │   │   │   │   │   │   ├── useMilestones.ts
    │   │   │   │   │   │   ├── useTimesheets.ts
    │   │   │   │   │   │   └── useWorkDiary.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── contracts-mutations.ts
    │   │   │   │   │   │   └── contracts-queries.ts  # BE: contracts-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── deliverables/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── deliverables-api.ts  # Deliverables API client
    │   │   │   │   │   │       # BE: contracts-be/deliverable, storage-be/asset
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useDeliverable.ts  # Single deliverable
    │   │   │   │   │   │   ├── useDeliverableRevisions.ts  # Revision management
    │   │   │   │   │   │   ├── useDeliverables.ts  # Deliverables list
    │   │   │   │   │   │   ├── useReviewDeliverable.ts  # Review deliverable (client)
    │   │   │   │   │   │   └── useUploadDeliverable.ts  # Upload deliverable
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── deliverables-mutations.ts  # Deliverable mutations
    │   │   │   │   │   │   └── deliverables-queries.ts  # Deliverable queries
    │   │   │   │   │   └── types.ts  # Deliverable types
    │   │   │   │   ├── disputes/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── disputes-api.ts  # Disputes API client
    │   │   │   │   │   │       # BE: contracts-be/dispute
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useDispute.ts  # Single dispute
    │   │   │   │   │   │   ├── useDisputeEvidence.ts  # Evidence management
    │   │   │   │   │   │   ├── useDisputeResolution.ts  # Resolution actions
    │   │   │   │   │   │   └── useDisputes.ts  # Disputes list
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── disputes-mutations.ts  # Dispute mutations
    │   │   │   │   │   │   └── disputes-queries.ts  # Dispute queries
    │   │   │   │   │   └── types.ts  # Dispute types
    │   │   │   │   ├── events/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── events-api.ts  # Events API client
    │   │   │   │   │   │       # BE: communications-be/events (if exists)
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── usePresence.ts  # User presence tracking
    │   │   │   │   │   │   ├── useRealTimeUpdates.ts  # Real-time data sync
    │   │   │   │   │   │   ├── useTypingIndicator.ts  # Typing indicators
    │   │   │   │   │   │   └── useWebSocket.ts  # WebSocket connection
    │   │   │   │   │   ├── providers/
    │   │   │   │   │   │   └── WebSocketProvider.tsx  # WebSocket context
    │   │   │   │   │   └── types.ts  # Event types
    │   │   │   │   ├── experiments/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── experiments-api.ts  # Experiments API
    │   │   │   │   │   │       # BE: utility-be/experiments
    │   │   │   │   │   │       # GET /v1/experiments/active
    │   │   │   │   │   │       # POST /v1/experiments/{experiment_id}/track
    │   │   │   │   │   │       # BE: utility/experiments (if exists) or utility/flags
    │   │   │   │   │   │       # POST /v1/experiments/track-event
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useExperiment.ts  # Get experiment variant
    │   │   │   │   │   │   │   # Experiment hook
    │   │   │   │   │   │   │   # A/B test variant
    │   │   │   │   │   │   ├── useExperimentTracking.ts  # Track experiment events
    │   │   │   │   │   │   └── useFeatureVariant.ts  # Feature variant
    │   │   │   │   │   ├── providers/
    │   │   │   │   │   │   ├── ExperimentProvider.tsx  # Experiment context
    │   │   │   │   │   │   └── ExperimentsProvider.tsx  # Experiments context
    │   │   │   │   │   ├── utils/
    │   │   │   │   │   │   ├── experiment-context.ts  # Experiment context
    │   │   │   │   │   │   ├── experiment-storage.ts  # Store assignments
    │   │   │   │   │   │   ├── tracking.ts  # Track events
    │   │   │   │   │   │   └── variant-assignment.ts  # Assign variant
    │   │   │   │   │   │       # Variant logic
    │   │   │   │   │   └── types.ts  # Experiment types
    │   │   │   │   ├── feature-flags/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── flags-api.ts  # Feature flags API
    │   │   │   │   │   │       # BE: utility/flags
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useFeatureFlag.ts  # Check single flag
    │   │   │   │   │   │   ├── useFeatureFlags.ts  # Get all flags
    │   │   │   │   │   │   └── useFeatureFlagVariant.ts  # A/B test variant
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   └── flags-queries.ts  # Flag queries
    │   │   │   │   │   ├── store/
    │   │   │   │   │   │   └── flags-store.ts  # Flags state (Zustand)
    │   │   │   │   │   └── types.ts  # Flag types
    │   │   │   │   ├── financial/  # Financial feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── financial-api.ts
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useEscrow.ts
    │   │   │   │   │   │   ├── useInvoices.ts
    │   │   │   │   │   │   ├── usePaymentMethods.ts
    │   │   │   │   │   │   ├── usePayoutMethods.ts
    │   │   │   │   │   │   ├── useTransactions.ts
    │   │   │   │   │   │   └── useWallet.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── financial-mutations.ts
    │   │   │   │   │   │   └── financial-queries.ts  # BE: financial-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── flags/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── flags-api.ts  # Feature flags API
    │   │   │   │   │   │       # BE: utility/flags
    │   │   │   │   │   │       # GET /v1/flags/user
    │   │   │   │   │   │       # GET /v1/flags/organization
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useFeatureFlag.ts  # Single flag hook
    │   │   │   │   │   │   ├── useFeatureFlags.ts  # Multiple flags
    │   │   │   │   │   │   └── useFeatureFlagVariant.ts  # A/B variant
    │   │   │   │   │   ├── providers/
    │   │   │   │   │   │   └── FeatureFlagsProvider.tsx  # Context provider
    │   │   │   │   │   ├── types.ts  # Flag types
    │   │   │   │   │   └── utils.ts  # Flag utilities
    │   │   │   │   ├── gamification/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── achievements-api.ts  # Achievements API
    │   │   │   │   │   │   │   # BE: users-be/achievement
    │   │   │   │   │   │   ├── badges-api.ts  # Badges API
    │   │   │   │   │   │   │   # BE: users-be/badge
    │   │   │   │   │   │   ├── gamification-api.ts  # Gamification API
    │   │   │   │   │   │   │   # BE: users-be/gamification (if exists) or reviews-be/reputation
    │   │   │   │   │   │   │   # GET /v1/gamification/points
    │   │   │   │   │   │   │   # GET /v1/gamification/achievements
    │   │   │   │   │   │   │   # GET /v1/gamification/leaderboard
    │   │   │   │   │   │   └── leaderboards-api.ts  # Leaderboards API
    │   │   │   │   │   │       # BE: users-be/leaderboard
    │   │   │   │   │   ├── components/
    │   │   │   │   │   │   ├── AchievementToast.tsx  # Achievement notification
    │   │   │   │   │   │   ├── BadgeCollection.tsx  # Badge collection
    │   │   │   │   │   │   ├── LeaderboardWidget.tsx  # Leaderboard widget
    │   │   │   │   │   │   └── PointsDisplay.tsx  # Points display
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useAchievements.ts  # Achievements
    │   │   │   │   │   │   │   # User achievements
    │   │   │   │   │   │   ├── useBadges.ts  # Badges
    │   │   │   │   │   │   │   # User badges
    │   │   │   │   │   │   ├── useLeaderboard.ts  # Leaderboard
    │   │   │   │   │   │   │   # Leaderboard data
    │   │   │   │   │   │   └── usePoints.ts  # User points
    │   │   │   │   │   │       # Points/reputation
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── gamification-mutations.ts  # Gamification mutations
    │   │   │   │   │   │   └── gamification-queries.ts  # Gamification queries
    │   │   │   │   │   └── types.ts  # Gamification types
    │   │   │   │   ├── geolocation/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── geolocation-api.ts  # Geolocation API
    │   │   │   │   │   │       # BE: utility/geo (if exists)
    │   │   │   │   │   │       # GET /v1/geo/location
    │   │   │   │   │   │       # GET /v1/geo/timezone
    │   │   │   │   │   │       # BE: utility-be/geolocation (if exists)
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useCountry.ts  # Country detection
    │   │   │   │   │   │   ├── useCountryDetection.ts  # Detect country
    │   │   │   │   │   │   ├── useDistanceCalculation.ts  # Calculate distance
    │   │   │   │   │   │   ├── useGeolocation.ts  # Get current location
    │   │   │   │   │   │   │   # Device location
    │   │   │   │   │   │   └── useTimezone.ts  # Detect timezone
    │   │   │   │   │   │       # User timezone
    │   │   │   │   │   ├── utils/
    │   │   │   │   │   │   ├── distance-calculator.ts  # Distance calculations
    │   │   │   │   │   │   └── geocoding.ts  # Geocoding utilities
    │   │   │   │   │   ├── types.ts  # Geolocation types
    │   │   │   │   │   └── utils.ts  # Geo utilities
    │   │   │   │   ├── i18n/
    │   │   │   │   │   ├── components/
    │   │   │   │   │   │   ├── FormattedMessage/
    │   │   │   │   │   │   │   └── FormattedMessage.tsx
    │   │   │   │   │   │   ├── LocaleSwitcher/
    │   │   │   │   │   │   │   ├── LocaleSwitcher.native.tsx
    │   │   │   │   │   │   │   ├── LocaleSwitcher.tsx
    │   │   │   │   │   │   │   └── LocaleSwitcher.web.tsx
    │   │   │   │   │   │   └── TranslationProvider/
    │   │   │   │   │   │       └── TranslationProvider.tsx
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useCurrencyFormat.ts  # Currency formatting
    │   │   │   │   │   │   ├── useDateFormat.ts  # Date formatting
    │   │   │   │   │   │   ├── useNumberFormat.ts  # Number formatting
    │   │   │   │   │   │   ├── useRTL.ts  # RTL detection
    │   │   │   │   │   │   └── useTranslation.ts  # (already exists, enhanced)
    │   │   │   │   │   ├── tools/
    │   │   │   │   │   │   ├── currency-formatter.ts  # Currency formatting
    │   │   │   │   │   │   ├── missing-keys-detector.ts  # Detect missing keys
    │   │   │   │   │   │   ├── pluralization-helper.ts  # Pluralization
    │   │   │   │   │   │   └── translation-manager.ts  # Manage translations
    │   │   │   │   │   └── utils/
    │   │   │   │   │       ├── fallback-loader.ts  # Fallback translations
    │   │   │   │   │       ├── interpolation.ts  # Message interpolation
    │   │   │   │   │       └── locale-detection.ts  # Auto-detect locale
    │   │   │   │   ├── interviews/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── interviews-api.ts  # Interviews API client
    │   │   │   │   │   │       # BE: proposals-be/interview
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useInterview.ts  # Single interview
    │   │   │   │   │   │   ├── useInterviewFeedback.ts  # Interview feedback
    │   │   │   │   │   │   ├── useInterviews.ts  # Interviews list
    │   │   │   │   │   │   └── useScheduleInterview.ts  # Schedule interview
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── interviews-mutations.ts  # Interview mutations
    │   │   │   │   │   │   └── interviews-queries.ts  # Interview queries
    │   │   │   │   │   └── types.ts  # Interview types
    │   │   │   │   ├── invitations/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── job-invitations-api.ts  # Job invitations API (client)
    │   │   │   │   │   │   │   # BE: jobs-be/invitation
    │   │   │   │   │   │   └── proposal-invites-api.ts  # Proposal invites API (freelancer)
    │   │   │   │   │   │       # BE: proposals-be/invite
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useAcceptInvite.ts  # Accept invite (freelancer)
    │   │   │   │   │   │   ├── useDeclineInvite.ts  # Decline invite (freelancer)
    │   │   │   │   │   │   ├── useInvitationAnalytics.ts  # Invitation metrics
    │   │   │   │   │   │   ├── useInvitations.ts  # Invitations management
    │   │   │   │   │   │   └── useSendInvitation.ts  # Send invitation (client)
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── invitations-mutations.ts  # Invitation mutations
    │   │   │   │   │   │   └── invitations-queries.ts  # Invitation queries
    │   │   │   │   │   └── types.ts  # Invitation types
    │   │   │   │   ├── jobs/  # Jobs feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── jobs-api.ts  # Jobs API client
    │   │   │   │   │   │       # Axios/Fetch wrapper
    │   │   │   │   │   │       # All jobs-be endpoints
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useCreateJob.ts  # Create job mutation
    │   │   │   │   │   │   ├── useDeleteJob.ts  # Delete job mutation
    │   │   │   │   │   │   ├── useJob.ts  # Single job query
    │   │   │   │   │   │   ├── useJobFilters.ts  # Job filters state
    │   │   │   │   │   │   ├── useJobs.ts  # List jobs query
    │   │   │   │   │   │   ├── useSaveJob.ts  # Save/bookmark job
    │   │   │   │   │   │   └── useUpdateJob.ts  # Update job mutation
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── jobs-mutations.ts  # TanStack Query mutations
    │   │   │   │   │   │   │   # BE: jobs-be
    │   │   │   │   │   │   │   # POST /v1/jobs
    │   │   │   │   │   │   │   # PUT /v1/jobs/{id}
    │   │   │   │   │   │   └── jobs-queries.ts  # TanStack Query queries
    │   │   │   │   │   │       # BE: jobs-be
    │   │   │   │   │   │       # GET /v1/jobs
    │   │   │   │   │   │       # GET /v1/jobs/{id}
    │   │   │   │   │   ├── utils/
    │   │   │   │   │   │   ├── job-helpers.ts  # Job utility functions
    │   │   │   │   │   │   └── job-validation.ts  # Job validation (Zod)
    │   │   │   │   │   └── types.ts  # Jobs types
    │   │   │   │   ├── learning/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── learning-paths-api.ts  # Learning paths API
    │   │   │   │   │   │   │   # BE: users-be/learning_path
    │   │   │   │   │   │   └── mentorship-api.ts  # Mentorship API
    │   │   │   │   │   │       # BE: users-be/mentorship
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useAchievements.ts  # Achievements/badges
    │   │   │   │   │   │   ├── useLearningPath.ts  # Single learning path
    │   │   │   │   │   │   ├── useLearningPaths.ts  # Learning paths list
    │   │   │   │   │   │   ├── useLearningProgress.ts  # Track progress
    │   │   │   │   │   │   └── useMentorship.ts  # Mentorship management
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── learning-mutations.ts  # Learning mutations
    │   │   │   │   │   │   └── learning-queries.ts  # Learning queries
    │   │   │   │   │   └── types.ts  # Learning types
    │   │   │   │   ├── messages/  # Messaging feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── messages-api.ts
    │   │   │   │   │   │   └── websocket.ts  # WebSocket client
    │   │   │   │   │   │       # WS: ws://communications-be/v1/realtime
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useConversation.ts
    │   │   │   │   │   │   ├── useConversations.ts
    │   │   │   │   │   │   ├── useMessages.ts
    │   │   │   │   │   │   ├── useRealtimeMessages.ts  # WebSocket
    │   │   │   │   │   │   └── useSendMessage.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── messages-mutations.ts
    │   │   │   │   │   │   └── messages-queries.ts  # BE: communications-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── moderation/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── moderation-api.ts  # Moderation API
    │   │   │   │   │   │       # BE: admin-be/moderation (if exists)
    │   │   │   │   │   │       # POST /v1/moderation/report
    │   │   │   │   │   │       # POST /v1/moderation/content-check
    │   │   │   │   │   │       # BE: admin-be/moderation
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useBlockUser.ts  # Block/unblock users
    │   │   │   │   │   │   ├── useContentModeration.ts  # Check content
    │   │   │   │   │   │   ├── useContentStatus.ts  # Content status check
    │   │   │   │   │   │   ├── useModerationActions.ts  # Moderation actions
    │   │   │   │   │   │   ├── useReportContent.ts  # Report content
    │   │   │   │   │   │   └── useReporting.ts  # Report content
    │   │   │   │   │   ├── utils/
    │   │   │   │   │   │   ├── content-validator.ts  # Content validation
    │   │   │   │   │   │   └── profanity-filter.ts  # Profanity filtering (client-side)
    │   │   │   │   │   ├── types.ts  # Moderation types
    │   │   │   │   │   └── utils.ts  # Content validation
    │   │   │   │   ├── negotiations/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── negotiations-api.ts  # Negotiations API client
    │   │   │   │   │   │       # BE: proposals-be/negotiation
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useNegotiation.ts  # Single negotiation
    │   │   │   │   │   │   ├── useNegotiationHistory.ts  # Negotiation history
    │   │   │   │   │   │   └── useNegotiationOffer.ts  # Make/accept/reject offer
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── negotiations-mutations.ts  # Negotiation mutations
    │   │   │   │   │   │   └── negotiations-queries.ts  # Negotiation queries
    │   │   │   │   │   └── types.ts  # Negotiation types
    │   │   │   │   ├── networking/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── connections-api.ts  # Connections API
    │   │   │   │   │   │   │   # BE: users-be/connection
    │   │   │   │   │   │   ├── groups-api.ts  # Groups API
    │   │   │   │   │   │   │   # BE: users-be/user_group
    │   │   │   │   │   │   └── referrals-api.ts  # Referrals API
    │   │   │   │   │   │       # BE: users-be/referral
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useConnectionRequest.ts  # Send/accept/reject
    │   │   │   │   │   │   ├── useConnections.ts  # Connections management
    │   │   │   │   │   │   ├── useGroups.ts  # Groups management
    │   │   │   │   │   │   ├── useNetworkRecommendations.ts  # Connection recommendations
    │   │   │   │   │   │   └── useReferrals.ts  # Referral management
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── networking-mutations.ts  # Networking mutations
    │   │   │   │   │   │   └── networking-queries.ts  # Networking queries
    │   │   │   │   │   └── types.ts  # Networking types
    │   │   │   │   ├── notifications/  # Notifications feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── notifications-api.ts
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useMarkAsRead.ts
    │   │   │   │   │   │   ├── useNotifications.ts
    │   │   │   │   │   │   ├── useRealtimeNotifications.ts  # WebSocket
    │   │   │   │   │   │   └── useUnreadCount.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── notifications-mutations.ts
    │   │   │   │   │   │   └── notifications-queries.ts  # BE: communications-be
    │   │   │   │   │   ├── store.ts  # Notifications store (Zustand)
    │   │   │   │   │   │   # - Unread counts, last fetched
    │   │   │   │   │   │   # BE: none (consumes API in features/notifications)
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── offline/
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useNetworkStatus.ts  # Network status
    │   │   │   │   │   │   ├── useOfflineQueue.ts  # Offline action queue
    │   │   │   │   │   │   ├── useOfflineStorage.ts  # Local storage
    │   │   │   │   │   │   └── useOfflineSync.ts  # Data synchronization
    │   │   │   │   │   ├── store/
    │   │   │   │   │   │   ├── offline-store.ts  # Offline state management
    │   │   │   │   │   │   └── sync-store.ts  # Sync state management
    │   │   │   │   │   └── types.ts  # Offline types
    │   │   │   │   ├── organizations/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── budgets-api.ts  # Budget management API
    │   │   │   │   │   │   │   # BE: financial-be/budget
    │   │   │   │   │   │   ├── organizations-api.ts  # Organizations API
    │   │   │   │   │   │   │   # BE: users-be/org
    │   │   │   │   │   │   └── vendors-api.ts  # Vendor management API
    │   │   │   │   │   │       # BE: users-be/org (vendor subdomain)
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useBudgets.ts  # Budget management
    │   │   │   │   │   │   ├── useOrganization.ts  # Organization details
    │   │   │   │   │   │   ├── useTeamMembers.ts  # Team member management
    │   │   │   │   │   │   └── useVendors.ts  # Vendor management
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── organizations-mutations.ts  # Org mutations
    │   │   │   │   │   │   └── organizations-queries.ts  # Org queries
    │   │   │   │   │   └── types.ts  # Organization types
    │   │   │   │   ├── performance/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── performance-api.ts  # Performance API
    │   │   │   │   │   │       # BE: utility/metrics (if exists) or utility/analytics
    │   │   │   │   │   │       # POST /v1/metrics/performance
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useErrorTracking.ts  # Error tracking
    │   │   │   │   │   │   ├── usePerformanceMonitor.ts  # Monitor performance
    │   │   │   │   │   │   └── useWebVitals.ts  # Web Vitals
    │   │   │   │   │   ├── utils/
    │   │   │   │   │   │   ├── error-reporter.ts  # Error reporter
    │   │   │   │   │   │   ├── performance-observer.ts  # Performance observer
    │   │   │   │   │   │   └── web-vitals-reporter.ts  # Web Vitals reporter
    │   │   │   │   │   └── types.ts  # Performance types
    │   │   │   │   ├── presence/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── presence-api.ts  # Presence API
    │   │   │   │   │   │       # BE: communications-be/presence (if exists)
    │   │   │   │   │   │       # POST /v1/presence/heartbeat
    │   │   │   │   │   │       # GET /v1/presence/users/{user_id}
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useLastSeen.ts  # Last seen time
    │   │   │   │   │   │   ├── useOnlineStatus.ts  # Online status
    │   │   │   │   │   │   └── usePresenceSubscription.ts  # Subscribe to presence
    │   │   │   │   │   └── types.ts  # Presence types
    │   │   │   │   ├── profile/  # Profile feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── profile-api.ts
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useEducation.ts
    │   │   │   │   │   │   ├── useExperience.ts
    │   │   │   │   │   │   ├── usePortfolio.ts
    │   │   │   │   │   │   ├── useProfile.ts
    │   │   │   │   │   │   ├── useServiceCatalog.ts
    │   │   │   │   │   │   ├── useSkills.ts
    │   │   │   │   │   │   └── useUpdateProfile.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── profile-mutations.ts
    │   │   │   │   │   │   └── profile-queries.ts  # BE: users-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── proposals/  # Proposals feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── proposals-api.ts
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useBidding.ts  # Bidding hooks
    │   │   │   │   │   │   ├── useProposal.ts
    │   │   │   │   │   │   ├── useProposals.ts
    │   │   │   │   │   │   ├── useSubmitProposal.ts
    │   │   │   │   │   │   └── useWithdrawProposal.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── proposals-mutations.ts
    │   │   │   │   │   │   └── proposals-queries.ts  # BE: proposals-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── realtime/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── websocket-client.ts  # WebSocket client
    │   │   │   │   │   │       # BE: communications-be/websocket (if exists)
    │   │   │   │   │   │       # WS: wss://api.skillsier.com/v1/ws
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── usePresence.ts  # User presence (online/offline)
    │   │   │   │   │   │   │   # User presence
    │   │   │   │   │   │   ├── useRealtimeAuction.ts  # Real-time auction updates
    │   │   │   │   │   │   ├── useRealtimeMessages.ts  # Real-time messages
    │   │   │   │   │   │   ├── useRealtimeNotifications.ts  # Real-time notifications
    │   │   │   │   │   │   ├── useTypingIndicator.ts  # Typing indicator
    │   │   │   │   │   │   └── useWebSocket.ts  # WebSocket connection
    │   │   │   │   │   │       # WebSocket hook
    │   │   │   │   │   ├── providers/
    │   │   │   │   │   │   └── WebSocketProvider.tsx  # WebSocket context
    │   │   │   │   │   ├── store/
    │   │   │   │   │   │   └── realtime-store.ts  # Real-time state (Zustand)
    │   │   │   │   │   ├── websocket/
    │   │   │   │   │   │   ├── client.ts  # WebSocket client
    │   │   │   │   │   │   ├── heartbeat.ts  # Connection health
    │   │   │   │   │   │   └── reconnection.ts  # Reconnection logic
    │   │   │   │   │   ├── types.ts  # Real-time types
    │   │   │   │   │   │   # WebSocket message types
    │   │   │   │   │   └── utils.ts  # Connection management
    │   │   │   │   ├── referrals/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── referrals-api.ts  # Referrals API (enhanced)
    │   │   │   │   │   │       # BE: users-be/referral
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useReferralCode.ts  # User referral code
    │   │   │   │   │   │   ├── useReferralProgram.ts  # Referral program details
    │   │   │   │   │   │   ├── useReferralStats.ts  # Referral statistics
    │   │   │   │   │   │   └── useRewards.ts  # Rewards management
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── referrals-mutations.ts  # Referral mutations
    │   │   │   │   │   │   └── referrals-queries.ts  # Referral queries
    │   │   │   │   │   └── types.ts  # Referral types
    │   │   │   │   ├── reviews/  # Reviews feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── reviews-api.ts
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useBadges.ts
    │   │   │   │   │   │   ├── useCreateReview.ts
    │   │   │   │   │   │   ├── useReviews.ts
    │   │   │   │   │   │   └── useReviewStats.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── reviews-mutations.ts
    │   │   │   │   │   │   └── reviews-queries.ts  # BE: reviews-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── safety/
    │   │   │   │   │   ├── components/
    │   │   │   │   │   │   ├── BlockConfirmation.tsx  # Block confirmation
    │   │   │   │   │   │   ├── ReportModal.tsx  # Report modal
    │   │   │   │   │   │   └── SafetyNotice.tsx  # Safety notice banner
    │   │   │   │   │   └── utils.ts  # Safety utilities
    │   │   │   │   ├── search/  # Search feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── recommendations-api.ts  # Recommendations API
    │   │   │   │   │   │   │   # BE: search-be/recommendation
    │   │   │   │   │   │   ├── saved-searches-api.ts  # Saved searches API
    │   │   │   │   │   │   │   # BE: search-be/saved-search
    │   │   │   │   │   │   ├── search-api.ts  # Search API (already may exist, ensuring completeness)
    │   │   │   │   │   │   │   # BE: search-be/query
    │   │   │   │   │   │   │   # Search API
    │   │   │   │   │   │   └── trending-api.ts  # Trending API
    │   │   │   │   │   │       # BE: search-be/trending
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useFreelancerSearch.ts
    │   │   │   │   │   │   ├── useJobSearch.ts
    │   │   │   │   │   │   ├── useRecommendations.ts  # Recommendations
    │   │   │   │   │   │   ├── useSavedSearches.ts  # Saved searches
    │   │   │   │   │   │   ├── useSearch.ts  # Search execution
    │   │   │   │   │   │   ├── useSearchHistory.ts  # Search history
    │   │   │   │   │   │   ├── useSearchSuggestions.ts  # Auto-complete suggestions
    │   │   │   │   │   │   └── useTrending.ts  # Trending items
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── search-mutations.ts  # Search mutations
    │   │   │   │   │   │   └── search-queries.ts  # BE: search-be
    │   │   │   │   │   │       # Search queries
    │   │   │   │   │   ├── store/
    │   │   │   │   │   │   └── search-store.ts  # Search UI state (filters, etc.)
    │   │   │   │   │   └── types.ts  # Search types
    │   │   │   │   ├── shortlists/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── shortlists-api.ts  # Shortlists API
    │   │   │   │   │   │       # BE: jobs-be/shortlist
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useAddToShortlist.ts  # Add candidate
    │   │   │   │   │   │   ├── useRemoveFromShortlist.ts  # Remove candidate
    │   │   │   │   │   │   ├── useShortlist.ts  # Single shortlist
    │   │   │   │   │   │   └── useShortlists.ts  # Shortlists management
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── shortlists-mutations.ts  # Shortlist mutations
    │   │   │   │   │   │   └── shortlists-queries.ts  # Shortlist queries
    │   │   │   │   │   └── types.ts  # Shortlist types
    │   │   │   │   ├── storage/  # File storage feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── storage-api.ts  # BE: storage-be
    │   │   │   │   │   │       # POST /v1/storage/upload
    │   │   │   │   │   │       # GET /v1/storage/presign
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useFileDownload.ts
    │   │   │   │   │   │   ├── usePresignedUrl.ts
    │   │   │   │   │   │   └── useUpload.ts
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── subscriptions/  # Subscriptions feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── subscriptions-api.ts  # Subscriptions API
    │   │   │   │   │   │       # BE: financial-be/subscription
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useConnects.ts
    │   │   │   │   │   │   ├── useEntitlements.ts  # Feature entitlements
    │   │   │   │   │   │   ├── usePlans.ts  # Subscription plans
    │   │   │   │   │   │   ├── useSubscription.ts  # Current subscription
    │   │   │   │   │   │   ├── useUpgrade.ts
    │   │   │   │   │   │   └── useUsage.ts  # Usage metrics
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── subscriptions-mutations.ts  # Subscription mutations
    │   │   │   │   │   │   └── subscriptions-queries.ts  # BE: subscriptions-be
    │   │   │   │   │   │       # Subscription queries
    │   │   │   │   │   └── types.ts  # Subscription types
    │   │   │   │   ├── tax/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── tax-api.ts  # Tax API client
    │   │   │   │   │   │       # BE: financial-be/tax
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useTaxForms.ts  # Tax forms management
    │   │   │   │   │   │   ├── useTaxReports.ts  # Tax reports
    │   │   │   │   │   │   └── useTaxSettings.ts  # Tax settings
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── tax-mutations.ts  # Tax mutations
    │   │   │   │   │   │   └── tax-queries.ts  # Tax queries
    │   │   │   │   │   └── types.ts  # Tax types
    │   │   │   │   ├── video/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── video-api.ts  # Video call API
    │   │   │   │   │   │       # BE: communications-be/video
    │   │   │   │   │   │       # POST /v1/video/rooms
    │   │   │   │   │   │       # GET /v1/video/rooms/{room_id}/token
    │   │   │   │   │   ├── components/
    │   │   │   │   │   │   ├── ParticipantGrid/
    │   │   │   │   │   │   │   ├── ParticipantGrid.tsx
    │   │   │   │   │   │   │   └── ParticipantGrid.types.ts
    │   │   │   │   │   │   ├── VideoControls/
    │   │   │   │   │   │   │   ├── VideoControls.tsx
    │   │   │   │   │   │   │   └── VideoControls.types.ts
    │   │   │   │   │   │   └── VideoRoom/
    │   │   │   │   │   │       ├── VideoRoom.native.tsx
    │   │   │   │   │   │       ├── VideoRoom.tsx
    │   │   │   │   │   │       ├── VideoRoom.types.ts
    │   │   │   │   │   │       └── VideoRoom.web.tsx
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useRecording.ts  # Call recording
    │   │   │   │   │   │   ├── useScreenShare.ts  # Screen sharing
    │   │   │   │   │   │   ├── useVideoCall.ts  # Video call management
    │   │   │   │   │   │   └── useVideoDevices.ts  # Device selection
    │   │   │   │   │   └── types.ts  # Video types
    │   │   │   │   ├── webhooks/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── webhooks-api.ts  # Webhooks API client
    │   │   │   │   │   │       # BE: utility-be/webhooks OR admin-be
    │   │   │   │   │   │       # GET /v1/webhooks
    │   │   │   │   │   │       # POST /v1/webhooks
    │   │   │   │   │   │       # PUT /v1/webhooks/{webhook_id}
    │   │   │   │   │   │       # DELETE /v1/webhooks/{webhook_id}
    │   │   │   │   │   │       # POST /v1/webhooks/{webhook_id}/test
    │   │   │   │   │   ├── components/
    │   │   │   │   │   │   ├── EventSelector/
    │   │   │   │   │   │   │   ├── EventSelector.tsx
    │   │   │   │   │   │   │   └── EventSelector.types.ts
    │   │   │   │   │   │   ├── WebhookForm/
    │   │   │   │   │   │   │   ├── WebhookForm.tsx
    │   │   │   │   │   │   │   ├── WebhookForm.types.ts
    │   │   │   │   │   │   │   └── WebhookForm.web.tsx
    │   │   │   │   │   │   └── WebhookLogs/
    │   │   │   │   │   │       ├── WebhookLogs.tsx
    │   │   │   │   │   │       └── WebhookLogs.types.ts
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useWebhook.ts  # Single webhook
    │   │   │   │   │   │   ├── useWebhookLogs.ts  # Webhook delivery logs
    │   │   │   │   │   │   ├── useWebhooks.ts  # List webhooks
    │   │   │   │   │   │   └── useWebhookTest.ts  # Test webhook
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── webhooks-mutations.ts  # Webhook mutations
    │   │   │   │   │   │   └── webhooks-queries.ts  # Webhook queries
    │   │   │   │   │   └── types.ts  # Webhook types
    │   │   │   │   └── work-tracking/
    │   │   │   │       ├── api/
    │   │   │   │       │   ├── timesheet-api.ts  # Timesheet API
    │   │   │   │       │   │   # BE: contracts-be/timesheet
    │   │   │   │       │   └── work-diary-api.ts  # Work diary API
    │   │   │   │       │       # BE: contracts-be/work_diary
    │   │   │   │       ├── hooks/
    │   │   │   │       │   ├── useApproveTimesheet.ts  # Approve timesheet (client)
    │   │   │   │       │   ├── useTimesheet.ts  # Timesheet management
    │   │   │   │       │   ├── useTimeTracking.ts  # Real-time time tracking
    │   │   │   │       │   └── useWorkDiary.ts  # Work diary entries
    │   │   │   │       ├── queries/
    │   │   │   │       │   ├── timesheet-mutations.ts  # Timesheet mutations
    │   │   │   │       │   ├── timesheet-queries.ts  # Timesheet queries
    │   │   │   │       │   ├── work-diary-mutations.ts  # Work diary mutations
    │   │   │   │       │   └── work-diary-queries.ts  # Work diary queries
    │   │   │   │       ├── store/
    │   │   │   │       │   └── time-tracker-store.ts  # Time tracker state (Zustand)
    │   │   │   │       └── types.ts  # Work tracking types
    │   │   │   ├── flags/
    │   │   │   │   ├── flags-client.ts  # Feature flags API client
    │   │   │   │   │   # BE: utility/flags
    │   │   │   │   │   # GET /v1/flags/user
    │   │   │   │   │   # GET /v1/flags/organization
    │   │   │   │   └── types.ts
    │   │   │   ├── gamification/
    │   │   │   │   ├── achievements/
    │   │   │   │   │   ├── achievement-engine.ts  # Achievement system
    │   │   │   │   │   │   # BE: users-be/achievements
    │   │   │   │   │   │   # GET /v1/achievements
    │   │   │   │   │   │   # POST /v1/achievements/{id}/claim
    │   │   │   │   │   ├── achievement-notifier.ts  # Achievement notifications
    │   │   │   │   │   ├── achievement-tracker.ts  # Progress tracking
    │   │   │   │   │   └── types.ts  # Achievement types
    │   │   │   │   ├── badges/
    │   │   │   │   │   ├── badge-criteria.ts  # Badge criteria
    │   │   │   │   │   ├── badge-display.ts  # Badge rendering
    │   │   │   │   │   └── badge-system.ts  # Badge management
    │   │   │   │   │       # BE: users-be/badges
    │   │   │   │   │       # GET /v1/badges
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useAchievements.ts  # Achievements hook
    │   │   │   │   │   ├── useBadges.ts  # Badges hook
    │   │   │   │   │   ├── useLeaderboard.ts  # Leaderboard hook
    │   │   │   │   │   └── usePoints.ts  # Points hook
    │   │   │   │   ├── leaderboards/
    │   │   │   │   │   ├── leaderboard-engine.ts  # Leaderboard system
    │   │   │   │   │   │   # BE: users-be/leaderboards
    │   │   │   │   │   │   # GET /v1/leaderboards/{type}
    │   │   │   │   │   ├── ranking-algorithms.ts  # Ranking logic
    │   │   │   │   │   └── time-windows.ts  # Time-based rankings
    │   │   │   │   └── points/
    │   │   │   │       ├── earning-rules.ts  # Point earning rules
    │   │   │   │       ├── points-system.ts  # Points management
    │   │   │   │       │   # BE: users-be/points
    │   │   │   │       │   # GET /v1/points/balance
    │   │   │   │       │   # POST /v1/points/earn
    │   │   │   │       └── redemption.ts  # Point redemption
    │   │   │   ├── geolocation/
    │   │   │   │   ├── browser-api/
    │   │   │   │   │   ├── fallback.ts  # Fallback strategies
    │   │   │   │   │   ├── geolocation-api.ts  # Browser Geolocation API
    │   │   │   │   │   └── permissions.ts  # Permission handling
    │   │   │   │   ├── country-detection/
    │   │   │   │   │   ├── currency-mapping.ts  # Country to currency
    │   │   │   │   │   ├── detector.ts  # Country detection
    │   │   │   │   │   └── locale-mapping.ts  # Country to locale
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useCountry.ts  # Country hook
    │   │   │   │   │   ├── useGeolocation.ts  # Geolocation hook
    │   │   │   │   │   └── useTimezone.ts  # Timezone hook
    │   │   │   │   ├── ip-detection/
    │   │   │   │   │   ├── cache.ts  # Location cache
    │   │   │   │   │   └── ip-resolver.ts  # IP-based location
    │   │   │   │   │       # BE: utility-be/geolocation
    │   │   │   │   │       # GET /v1/geo/ip-lookup
    │   │   │   │   └── timezone/
    │   │   │   │       ├── dst-handler.ts  # Daylight saving handling
    │   │   │   │       ├── timezone-converter.ts  # Timezone conversion
    │   │   │   │       └── timezone-detector.ts  # Timezone detection
    │   │   │   ├── hooks/  # General hooks
    │   │   │   │   ├── useClickOutside.ts
    │   │   │   │   ├── useCopyToClipboard.ts
    │   │   │   │   ├── useDebounce.ts
    │   │   │   │   ├── useIntersectionObserver.ts
    │   │   │   │   ├── useLocalStorage.ts
    │   │   │   │   ├── useMediaQuery.ts
    │   │   │   │   └── useToggle.ts
    │   │   │   ├── i18n/  # Internationalization
    │   │   │   │   ├── locales/  # Locale files
    │   │   │   │   │   ├── ar/  # Arabic
    │   │   │   │   │   ├── de/  # German
    │   │   │   │   │   ├── en/
    │   │   │   │   │   │   ├── admin.json
    │   │   │   │   │   │   ├── auth.json
    │   │   │   │   │   │   ├── common.json
    │   │   │   │   │   │   ├── contracts.json
    │   │   │   │   │   │   ├── errors.json
    │   │   │   │   │   │   ├── financial.json
    │   │   │   │   │   │   ├── jobs.json
    │   │   │   │   │   │   ├── messages.json
    │   │   │   │   │   │   ├── profile.json
    │   │   │   │   │   │   ├── proposals.json
    │   │   │   │   │   │   ├── reviews.json
    │   │   │   │   │   │   ├── settings.json
    │   │   │   │   │   │   └── subscription.json
    │   │   │   │   │   ├── es/  # Spanish
    │   │   │   │   │   ├── fr/  # French
    │   │   │   │   │   ├── hi/  # Hindi
    │   │   │   │   │   ├── ru/  # Russian
    │   │   │   │   │   ├── tr/  # Turkish
    │   │   │   │   │   └── zh/  # Chinese
    │   │   │   │   └── config.ts  # i18n configuration
    │   │   │   ├── incidents/
    │   │   │   │   ├── incidents-client.ts  # Incidents API client
    │   │   │   │   │   # BE: utility/status
    │   │   │   │   │   # GET /v1/status/incidents
    │   │   │   │   │   # POST /v1/status/incidents
    │   │   │   │   │   # PUT /v1/status/incidents/{id}
    │   │   │   │   │   # GET /v1/status/maintenance
    │   │   │   │   └── types.ts
    │   │   │   ├── learning/
    │   │   │   │   ├── learning-client.ts  # Learning API client
    │   │   │   │   │   # BE: learning-be (if exists) or external LMS
    │   │   │   │   │   # GET /v1/learning/courses
    │   │   │   │   │   # GET /v1/learning/courses/{id}
    │   │   │   │   │   # POST /v1/learning/assessments/{id}/submit
    │   │   │   │   └── types.ts
    │   │   │   ├── lib/  # Shared utilities
    │   │   │   │   ├── api/
    │   │   │   │   │   ├── client.ts  # Axios/Fetch client setup
    │   │   │   │   │   │   # - Base URL
    │   │   │   │   │   │   # - Auth interceptors
    │   │   │   │   │   │   # - Error handling
    │   │   │   │   │   │   # - Request/response logging
    │   │   │   │   │   ├── endpoints.ts  # API endpoint constants
    │   │   │   │   │   └── error-handler.ts  # Global error handler
    │   │   │   │   ├── constants/
    │   │   │   │   │   ├── api.ts  # API constants
    │   │   │   │   │   ├── app.ts  # App constants
    │   │   │   │   │   ├── permissions.ts  # RBAC permissions
    │   │   │   │   │   └── routes.ts  # Route constants
    │   │   │   │   ├── formatting/
    │   │   │   │   │   ├── currency.ts  # Currency formatting
    │   │   │   │   │   ├── date.ts  # Date formatting
    │   │   │   │   │   ├── number.ts  # Number formatting
    │   │   │   │   │   └── string.ts  # String utilities
    │   │   │   │   ├── validation/
    │   │   │   │   │   ├── schemas.ts  # Zod schemas
    │   │   │   │   │   └── validators.ts  # Custom validators
    │   │   │   │   └── websocket/
    │   │   │   │       ├── client.ts  # WebSocket client
    │   │   │   │       │   # WS: ws://communications-be/v1/realtime
    │   │   │   │       │   # - Connection management
    │   │   │   │       │   # - Reconnection logic
    │   │   │   │       │   # - Event subscriptions
    │   │   │   │       └── events.ts  # WebSocket event types
    │   │   │   ├── logging/
    │   │   │   │   ├── context/
    │   │   │   │   │   ├── request-context.ts  # Request context
    │   │   │   │   │   ├── trace-context.ts  # Trace context
    │   │   │   │   │   └── user-context.ts  # User context
    │   │   │   │   ├── formatters/
    │   │   │   │   │   ├── json-formatter.ts
    │   │   │   │   │   ├── pretty-formatter.ts
    │   │   │   │   │   └── structured-formatter.ts
    │   │   │   │   ├── transports/
    │   │   │   │   │   ├── console-transport.ts
    │   │   │   │   │   ├── file-transport.ts
    │   │   │   │   │   └── remote-transport.ts
    │   │   │   │   ├── log-levels.ts  # Log level configuration
    │   │   │   │   └── logger.ts  # Base logger
    │   │   │   ├── moderation/
    │   │   │   │   ├── auto-moderation/
    │   │   │   │   │   ├── action-executor.ts  # Automated actions
    │   │   │   │   │   ├── escalation.ts  # Escalation logic
    │   │   │   │   │   └── rule-engine.ts  # Auto-moderation rules
    │   │   │   │   ├── content-validation/
    │   │   │   │   │   ├── format-validator.ts  # Format validation
    │   │   │   │   │   ├── link-validator.ts  # Link validation
    │   │   │   │   │   ├── profanity-filter.ts  # Profanity detection
    │   │   │   │   │   └── spam-detector.ts  # Spam detection
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useContentValidation.ts  # Validation hook
    │   │   │   │   │   ├── useModerationStatus.ts  # Status hook
    │   │   │   │   │   └── useReporting.ts  # Reporting hook
    │   │   │   │   ├── reporting/
    │   │   │   │   │   ├── evidence-collection.ts  # Evidence gathering
    │   │   │   │   │   ├── report-submission.ts  # Report submission
    │   │   │   │   │   │   # BE: admin-be/moderation
    │   │   │   │   │   │   # POST /v1/reports
    │   │   │   │   │   └── report-types.ts  # Report categories
    │   │   │   │   ├── moderation-client.ts  # Moderation API client
    │   │   │   │   │   # BE: admin-be/moderation (if exists)
    │   │   │   │   │   # POST /v1/moderation/report
    │   │   │   │   │   # POST /v1/moderation/content-check
    │   │   │   │   └── types.ts
    │   │   │   ├── monitoring/
    │   │   │   │   ├── analytics/
    │   │   │   │   │   ├── providers/
    │   │   │   │   │   │   ├── amplitude.ts  # Amplitude
    │   │   │   │   │   │   ├── google-analytics.ts  # GA4
    │   │   │   │   │   │   └── mixpanel.ts  # Mixpanel
    │   │   │   │   │   ├── tracking/
    │   │   │   │   │   │   ├── event-tracker.ts
    │   │   │   │   │   │   ├── page-tracker.ts
    │   │   │   │   │   │   └── user-tracker.ts
    │   │   │   │   │   ├── analytics-config.ts  # Analytics configuration
    │   │   │   │   │   ├── event-tracker.ts  # Event tracking wrapper
    │   │   │   │   │   ├── event-tracking.ts  # Event tracking setup
    │   │   │   │   │   ├── google-analytics.ts  # GA4
    │   │   │   │   │   ├── mixpanel-config.ts  # Mixpanel configuration
    │   │   │   │   │   ├── mixpanel.ts  # Mixpanel
    │   │   │   │   │   ├── segment-config.ts  # Segment configuration
    │   │   │   │   │   └── segment.ts  # Segment
    │   │   │   │   ├── error-tracking/
    │   │   │   │   │   ├── custom-integrations.ts  # Custom integrations
    │   │   │   │   │   ├── error-boundary-setup.ts  # Error boundary setup
    │   │   │   │   │   ├── error-boundary.tsx  # Error boundary HOC
    │   │   │   │   │   ├── error-enrichment.ts  # Context enrichment
    │   │   │   │   │   ├── error-filters.ts  # Filter errors
    │   │   │   │   │   ├── error-reporter.ts  # Custom error reporter
    │   │   │   │   │   ├── sentry-config.ts  # Sentry configuration
    │   │   │   │   │   └── sentry.ts  # Sentry configuration
    │   │   │   │   ├── logging/
    │   │   │   │   │   ├── transports/
    │   │   │   │   │   │   ├── console-transport.ts
    │   │   │   │   │   │   ├── file-transport.ts  # Mobile only
    │   │   │   │   │   │   └── remote-transport.ts
    │   │   │   │   │   ├── log-levels.ts  # Log level configuration
    │   │   │   │   │   └── logger.ts  # Application logger
    │   │   │   │   ├── performance/
    │   │   │   │   │   ├── api-monitor.ts  # API performance monitoring
    │   │   │   │   │   ├── bundle-monitor.ts  # Bundle size monitoring
    │   │   │   │   │   ├── custom-metrics.ts  # Custom performance metrics
    │   │   │   │   │   ├── metrics.ts  # Custom metrics
    │   │   │   │   │   ├── performance-observer.ts  # Performance API
    │   │   │   │   │   │   # Performance monitoring
    │   │   │   │   │   ├── profiling.ts  # Performance profiling
    │   │   │   │   │   ├── web-vitals-reporter.ts  # Web Vitals reporting
    │   │   │   │   │   └── web-vitals.ts  # Web Vitals monitoring
    │   │   │   │   │       # Core Web Vitals
    │   │   │   │   ├── sentry/
    │   │   │   │   │   ├── error-boundary.tsx  # Sentry error boundary
    │   │   │   │   │   ├── sentry-config.ts  # Sentry configuration
    │   │   │   │   │   ├── sentry-init.native.ts  # Mobile initialization
    │   │   │   │   │   └── sentry-init.web.ts  # Web initialization
    │   │   │   │   ├── analytics.ts  # Analytics setup
    │   │   │   │   ├── logger.ts  # Logging configuration
    │   │   │   │   ├── sentry.ts  # Error tracking setup
    │   │   │   │   └── web-vitals.ts  # Performance monitoring
    │   │   │   ├── offline/
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useOfflineQueue.ts  # Offline queue hook
    │   │   │   │   │   ├── useOfflineStorage.ts  # Offline storage hook
    │   │   │   │   │   ├── useOnlineStatus.ts  # Online status hook
    │   │   │   │   │   └── useSyncStatus.ts  # Sync status hook
    │   │   │   │   ├── network/
    │   │   │   │   │   ├── bandwidth-estimator.ts  # Bandwidth estimation
    │   │   │   │   │   ├── connection-quality.ts  # Connection quality
    │   │   │   │   │   └── network-monitor.ts  # Network status monitoring
    │   │   │   │   ├── queue/
    │   │   │   │   │   ├── conflict-resolution.ts  # Conflict handling
    │   │   │   │   │   ├── operation-queue.ts  # Offline operation queue
    │   │   │   │   │   ├── queue-processor.ts  # Queue processing
    │   │   │   │   │   └── retry-strategies.ts  # Retry logic
    │   │   │   │   ├── storage/
    │   │   │   │   │   ├── async-storage.ts  # React Native AsyncStorage
    │   │   │   │   │   ├── cache-manager.ts  # Cache management
    │   │   │   │   │   └── indexed-db.ts  # IndexedDB wrapper
    │   │   │   │   └── sync/
    │   │   │   │       ├── sync-strategies/
    │   │   │   │       │   ├── conflict-merge.ts  # Merge strategies
    │   │   │   │       │   ├── full-sync.ts  # Full synchronization
    │   │   │   │       │   └── incremental-sync.ts  # Incremental sync
    │   │   │   │       ├── change-detection.ts  # Change detection
    │   │   │   │       ├── sync-engine.ts  # Synchronization engine
    │   │   │   │       └── version-tracking.ts  # Version management
    │   │   │   ├── performance/
    │   │   │   │   ├── custom-metrics/
    │   │   │   │   │   ├── component-render-time.ts  # Component perf
    │   │   │   │   │   ├── route-change-duration.ts  # Route timing
    │   │   │   │   │   └── time-to-interactive.ts  # Custom TTI
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useAnalyticsTracking.ts  # Analytics events
    │   │   │   │   │   ├── useErrorTracking.ts  # Error monitoring
    │   │   │   │   │   └── usePerformanceMetrics.ts  # Web vitals tracking
    │   │   │   │   ├── monitoring/
    │   │   │   │   │   ├── long-tasks.ts  # Long task detection
    │   │   │   │   │   ├── memory-usage.ts  # Memory monitoring
    │   │   │   │   │   └── performance-observer.ts  # Performance Observer API
    │   │   │   │   ├── resource-timing/
    │   │   │   │   │   ├── api-timing.ts  # API request timing
    │   │   │   │   │   ├── asset-timing.ts  # Asset load timing
    │   │   │   │   │   └── navigation-timing.ts  # Navigation timing
    │   │   │   │   ├── utils/
    │   │   │   │   │   ├── error-reporter.ts  # Error reporting
    │   │   │   │   │   ├── metrics-collector.ts  # Metrics collection
    │   │   │   │   │   └── trace-headers.ts  # Distributed tracing
    │   │   │   │   ├── web-vitals/
    │   │   │   │   │   ├── collectors/
    │   │   │   │   │   │   ├── cls-collector.ts  # Cumulative Layout Shift
    │   │   │   │   │   │   ├── fid-collector.ts  # First Input Delay
    │   │   │   │   │   │   ├── inp-collector.ts  # Interaction to Next Paint
    │   │   │   │   │   │   ├── lcp-collector.ts  # Largest Contentful Paint
    │   │   │   │   │   │   └── ttfb-collector.ts  # Time to First Byte
    │   │   │   │   │   ├── attribution.ts  # Performance attribution
    │   │   │   │   │   └── reporting.ts  # Web Vitals reporting
    │   │   │   │   └── types.ts  # Performance types
    │   │   │   ├── presence/
    │   │   │   │   ├── presence-client.ts  # Presence API client
    │   │   │   │   │   # BE: communications-be/presence (if exists)
    │   │   │   │   │   # POST /v1/presence/heartbeat
    │   │   │   │   │   # GET /v1/presence/users/{id}
    │   │   │   │   └── types.ts
    │   │   │   ├── realtime/
    │   │   │   │   ├── live-updates/
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useLiveDocument.ts  # Live document hook
    │   │   │   │   │   │   └── useLiveQuery.ts  # Live query hook
    │   │   │   │   │   ├── event-stream.ts  # Server-sent events
    │   │   │   │   │   └── polling-fallback.ts  # Polling fallback
    │   │   │   │   ├── presence/
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useOnlineUsers.ts  # Online users hook
    │   │   │   │   │   │   └── usePresence.ts  # Presence hook
    │   │   │   │   │   ├── presence-tracker.ts  # User presence tracking
    │   │   │   │   │   │   # BE: communications-be/presence
    │   │   │   │   │   │   # POST /v1/presence/update
    │   │   │   │   │   └── types.ts  # Presence types
    │   │   │   │   ├── typing-indicators/
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   └── useTypingIndicator.ts  # Typing hook
    │   │   │   │   │   └── typing-manager.ts  # Typing indicator management
    │   │   │   │   │       # BE: communications-be/typing
    │   │   │   │   │       # POST /v1/conversations/{id}/typing
    │   │   │   │   └── websocket/
    │   │   │   │       ├── connection-manager.ts  # WebSocket connection management
    │   │   │   │       │   # - Connection pooling
    │   │   │   │       │   # - Reconnection logic
    │   │   │   │       │   # - Heartbeat monitoring
    │   │   │   │       │   # BE: communications-be/websocket
    │   │   │   │       ├── error-recovery.ts  # Error handling & recovery
    │   │   │   │       ├── message-handler.ts  # Message routing
    │   │   │   │       │   # - Message type handling
    │   │   │   │       │   # - Event dispatch
    │   │   │   │       └── subscription-manager.ts  # Subscription management
    │   │   │   │           # - Channel subscriptions
    │   │   │   │           # - Topic filtering
    │   │   │   ├── security/
    │   │   │   │   ├── auth/
    │   │   │   │   │   ├── device-trust.ts  # Device trust
    │   │   │   │   │   ├── session-manager.ts  # Session management
    │   │   │   │   │   ├── token-manager.ts  # Token management
    │   │   │   │   │   └── token-refresh.ts  # Auto token refresh
    │   │   │   │   ├── csrf/
    │   │   │   │   │   ├── csrf-token.ts  # CSRF token management
    │   │   │   │   │   ├── csrf-validation.ts  # CSRF validation
    │   │   │   │   │   └── double-submit-cookie.ts  # Double submit pattern
    │   │   │   │   ├── encryption/
    │   │   │   │   │   ├── client-encryption.ts  # Client-side encryption
    │   │   │   │   │   ├── crypto-utils.ts  # Encryption utilities
    │   │   │   │   │   ├── key-management.ts  # Key management
    │   │   │   │   │   └── secure-storage.ts  # Secure storage
    │   │   │   │   ├── headers/
    │   │   │   │   │   ├── cors.ts  # CORS configuration
    │   │   │   │   │   ├── csp.ts  # Content Security Policy
    │   │   │   │   │   └── security-headers.ts  # Security headers config
    │   │   │   │   ├── monitoring/
    │   │   │   │   │   ├── anomaly-detector.ts  # Anomaly detection
    │   │   │   │   │   ├── breach-detector.ts  # Breach detection
    │   │   │   │   │   └── security-monitor.ts  # Security monitoring
    │   │   │   │   ├── sanitization/
    │   │   │   │   │   ├── html-sanitizer.ts  # HTML sanitization
    │   │   │   │   │   ├── input-sanitizer.ts  # Input sanitization
    │   │   │   │   │   ├── sql-injection-prevention.ts  # SQL injection prevention
    │   │   │   │   │   └── xss-prevention.ts  # XSS prevention
    │   │   │   │   ├── validation/
    │   │   │   │   │   ├── file-validation.ts  # File validation
    │   │   │   │   │   ├── input-validation.ts  # Input validation
    │   │   │   │   │   ├── input-validator.ts  # Input validation
    │   │   │   │   │   ├── sanitizer.ts  # HTML/input sanitization
    │   │   │   │   │   ├── schema-validation.ts  # Schema validation
    │   │   │   │   │   ├── sql-injection-guard.ts  # SQL injection protection
    │   │   │   │   │   └── xss-protection.ts  # XSS protection
    │   │   │   │   ├── csrf.ts  # CSRF token management
    │   │   │   │   │   # CSRF protection
    │   │   │   │   ├── encryption.ts  # Client-side encryption
    │   │   │   │   │   # - Encrypt sensitive data
    │   │   │   │   │   # - Decrypt data
    │   │   │   │   ├── permissions.ts  # Permission checks
    │   │   │   │   │   # - Role-based access
    │   │   │   │   │   # - Feature flags
    │   │   │   │   │   # - Resource ownership
    │   │   │   │   ├── sanitization.ts  # Input sanitization
    │   │   │   │   │   # - HTML sanitization
    │   │   │   │   │   # - SQL injection prevention
    │   │   │   │   │   # - XSS prevention
    │   │   │   │   └── validation.ts  # Input validation
    │   │   │   ├── sourcing/
    │   │   │   │   ├── sourcing-client.ts  # Sourcing API client
    │   │   │   │   │   # BE: jobs-be/campaign (if exists) or jobs-be/job
    │   │   │   │   │   # GET /v1/sourcing/campaigns
    │   │   │   │   │   # POST /v1/sourcing/campaigns
    │   │   │   │   │   # GET /v1/sourcing/talent-pools
    │   │   │   │   └── types.ts
    │   │   │   ├── testing/
    │   │   │   │   ├── factories/
    │   │   │   │   │   ├── job-factory.ts  # Job factory
    │   │   │   │   │   ├── proposal-factory.ts  # Proposal factory
    │   │   │   │   │   └── user-factory.ts  # User factory
    │   │   │   │   ├── fixtures/
    │   │   │   │   │   ├── auth.ts  # Auth fixtures
    │   │   │   │   │   ├── contract.ts  # Contract fixtures
    │   │   │   │   │   ├── job.ts  # Job fixtures
    │   │   │   │   │   ├── jobs.ts  # Job fixtures
    │   │   │   │   │   ├── proposal.ts  # Proposal fixtures
    │   │   │   │   │   ├── user.ts  # User fixtures
    │   │   │   │   │   └── users.ts  # User fixtures
    │   │   │   │   ├── mock-api/
    │   │   │   │   │   ├── handlers.ts  # MSW handlers
    │   │   │   │   │   └── server.ts  # MSW server setup
    │   │   │   │   ├── mock-data/
    │   │   │   │   │   ├── factories/
    │   │   │   │   │   │   ├── job-factory.ts  # Job factory
    │   │   │   │   │   │   ├── proposal-factory.ts  # Proposal factory
    │   │   │   │   │   │   └── user-factory.ts  # User factory
    │   │   │   │   │   ├── contracts.ts  # Mock contract data
    │   │   │   │   │   ├── jobs.ts  # Mock job data
    │   │   │   │   │   ├── messages.ts  # Mock message data
    │   │   │   │   │   ├── payments.ts  # Mock payment data
    │   │   │   │   │   ├── proposals.ts  # Mock proposal data
    │   │   │   │   │   └── users.ts  # Mock user data
    │   │   │   │   ├── mocks/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── responses/
    │   │   │   │   │   │   │   ├── error-responses.ts
    │   │   │   │   │   │   │   └── success-responses.ts
    │   │   │   │   │   │   ├── communications-mock.ts  # Mock communications API
    │   │   │   │   │   │   ├── contracts-mock.ts  # Mock contracts API
    │   │   │   │   │   │   ├── financial-mock.ts  # Mock financial API
    │   │   │   │   │   │   ├── handlers.ts  # MSW handlers
    │   │   │   │   │   │   ├── jobs-mock.ts  # Mock jobs API
    │   │   │   │   │   │   ├── proposals-mock.ts  # Mock proposals API
    │   │   │   │   │   │   ├── server.ts  # MSW server
    │   │   │   │   │   │   └── users-mock.ts  # Mock users API
    │   │   │   │   │   ├── data/
    │   │   │   │   │   │   ├── jobs.ts  # Mock job data
    │   │   │   │   │   │   ├── proposals.ts  # Mock proposal data
    │   │   │   │   │   │   └── users.ts  # Mock user data
    │   │   │   │   │   ├── handlers/
    │   │   │   │   │   │   ├── auth-handlers.ts  # MSW auth handlers
    │   │   │   │   │   │   └── users-handlers.ts  # MSW users handlers
    │   │   │   │   │   ├── storage/
    │   │   │   │   │   │   └── local-storage-mock.ts
    │   │   │   │   │   ├── websocket/
    │   │   │   │   │   │   └── websocket-mock.ts
    │   │   │   │   │   ├── auth.ts  # Auth mocks
    │   │   │   │   │   ├── contracts.ts  # Contract mocks
    │   │   │   │   │   ├── financial.ts  # Financial mocks
    │   │   │   │   │   ├── jobs.ts  # Job data mocks
    │   │   │   │   │   ├── proposals.ts  # Proposal mocks
    │   │   │   │   │   └── users.ts  # User data mocks
    │   │   │   │   ├── providers/
    │   │   │   │   │   └── TestProviders.tsx  # Test provider wrapper
    │   │   │   │   ├── setup/
    │   │   │   │   │   ├── global-setup.ts  # Global test setup
    │   │   │   │   │   ├── jest.setup.ts  # Jest configuration
    │   │   │   │   │   ├── msw-setup.ts  # MSW server setup
    │   │   │   │   │   ├── msw.setup.ts  # MSW setup
    │   │   │   │   │   ├── test-env.ts  # Test environment
    │   │   │   │   │   ├── test-providers.tsx  # Test providers wrapper
    │   │   │   │   │   ├── test-setup.ts  # Test environment setup
    │   │   │   │   │   └── testing-library.setup.ts  # Testing Library setup
    │   │   │   │   ├── test-utils/
    │   │   │   │   │   ├── custom-matchers.ts  # Custom Jest matchers
    │   │   │   │   │   ├── mock-providers.tsx  # Mock providers
    │   │   │   │   │   ├── render-utils.tsx  # React Testing Library utils
    │   │   │   │   │   └── test-wrapper.tsx  # Test wrapper component
    │   │   │   │   ├── utils/
    │   │   │   │   │   ├── accessibility-checker.ts  # A11y testing utils
    │   │   │   │   │   ├── mock-api.ts  # Mock API responses
    │   │   │   │   │   ├── mock-providers.tsx  # Mock context providers
    │   │   │   │   │   ├── render-with-providers.tsx  # Custom render
    │   │   │   │   │   ├── render.tsx  # Custom render with providers
    │   │   │   │   │   ├── test-data-generators.ts  # Generate test data
    │   │   │   │   │   ├── test-helpers.ts  # Helper functions
    │   │   │   │   │   └── wait-for-async.ts  # Async utilities
    │   │   │   │   └── test-utils.ts  # Common test utilities
    │   │   │   ├── utils/
    │   │   │   │   └── logger/
    │   │   │   │       ├── formatters/
    │   │   │   │       │   ├── json.ts  # JSON formatter
    │   │   │   │       │   └── pretty.ts  # Pretty formatter
    │   │   │   │       ├── transports/
    │   │   │   │       │   ├── console.ts  # Console transport
    │   │   │   │       │   ├── file.ts  # File transport
    │   │   │   │       │   └── remote.ts  # Remote logging service
    │   │   │   │       ├── log-levels.ts  # Log levels
    │   │   │   │       └── logger.ts  # Logger implementation
    │   │   │   └── webhooks/
    │   │   │       ├── types.ts
    │   │   │       └── webhooks-client.ts  # Webhooks API client
    │   │   │           # BE: utility/webhooks (if exists) or admin-be
    │   │   │           # GET /v1/webhooks
    │   │   │           # POST /v1/webhooks
    │   │   │           # PUT /v1/webhooks/{id}
    │   │   │           # DELETE /v1/webhooks/{id}
    │   │   ├── (none)
    │   │   ├── package.json
    │   │   ├── README.md
    │   │   └── tsconfig.json
    │   ├── testing/
    │   │   ├── factories/
    │   │   │   ├── ...  # Other factories
    │   │   │   ├── job-factory.ts  # Job factory
    │   │   │   ├── proposal-factory.ts  # Proposal factory
    │   │   │   └── user-factory.ts  # User factory
    │   │   ├── fixtures/
    │   │   │   ├── ...  # Other fixtures
    │   │   │   ├── auth.ts  # Auth fixtures
    │   │   │   ├── jobs.ts  # Job fixtures
    │   │   │   └── users.ts  # User fixtures
    │   │   ├── mocks/
    │   │   │   ├── api/
    │   │   │   │   ├── ...  # Other service mocks
    │   │   │   │   ├── communications-mock.ts  # Mock communications API
    │   │   │   │   ├── contracts-mock.ts  # Mock contracts API
    │   │   │   │   ├── financial-mock.ts  # Mock financial API
    │   │   │   │   ├── jobs-mock.ts  # Mock jobs API
    │   │   │   │   ├── proposals-mock.ts  # Mock proposals API
    │   │   │   │   └── users-mock.ts  # Mock users API
    │   │   │   ├── data/
    │   │   │   │   ├── ...  # Other mock data
    │   │   │   │   ├── jobs.ts  # Mock job data
    │   │   │   │   ├── proposals.ts  # Mock proposal data
    │   │   │   │   └── users.ts  # Mock user data
    │   │   │   └── handlers/
    │   │   │       ├── ...  # Other MSW handlers
    │   │   │       ├── auth-handlers.ts  # MSW auth handlers
    │   │   │       └── users-handlers.ts  # MSW users handlers
    │   │   ├── setup/
    │   │   │   ├── msw-setup.ts  # MSW server setup
    │   │   │   ├── test-providers.tsx  # Test providers wrapper
    │   │   │   └── test-setup.ts  # Test environment setup
    │   │   └── utils/
    │   │       ├── accessibility-checker.ts  # A11y testing utils
    │   │       ├── render-with-providers.tsx  # Custom render
    │   │       └── wait-for-async.ts  # Async utilities
    │   ├── types/  # TypeScript type definitions
    │   │   # Shared TypeScript types
    │   │   ├── src/
    │   │   │   ├── api/  # API response types
    │   │   │   │   ├── performance/
    │   │   │   │   │   └── index.ts  # Performance types
    │   │   │   │   │       # - Web Vitals
    │   │   │   │   │       # - Performance metric
    │   │   │   │   │       # - Error event
    │   │   │   │   ├── realtime/
    │   │   │   │   │   └── index.ts  # Real-time types
    │   │   │   │   │       # - WebSocket message
    │   │   │   │   │       # - Presence update
    │   │   │   │   │       # - Event subscription
    │   │   │   │   ├── webhooks/
    │   │   │   │   │   └── index.ts  # Webhook types
    │   │   │   │   │       # - Webhook event
    │   │   │   │   │       # - Webhook subscription
    │   │   │   │   │       # - Webhook delivery
    │   │   │   │   ├── admin.ts
    │   │   │   │   ├── common.ts  # Common types (pagination, filters, etc.)
    │   │   │   │   ├── contracts.ts
    │   │   │   │   ├── financial.ts
    │   │   │   │   ├── jobs.ts
    │   │   │   │   ├── messages.ts
    │   │   │   │   ├── notifications.ts
    │   │   │   │   ├── proposals.ts
    │   │   │   │   ├── review.ts
    │   │   │   │   ├── search.ts
    │   │   │   │   ├── storage.ts
    │   │   │   │   ├── subscriptions.ts
    │   │   │   │   └── users.ts
    │   │   │   ├── domains/
    │   │   │   │   ├── achievements/
    │   │   │   │   │   └── index.ts  # Achievement types
    │   │   │   │   │       # - Achievement definition
    │   │   │   │   │       # - User achievement
    │   │   │   │   │       # - Badge
    │   │   │   │   │       # - Points
    │   │   │   │   ├── activity/
    │   │   │   │   │   └── index.ts  # Activity feed types
    │   │   │   │   │       # - Activity event
    │   │   │   │   │       # - Activity feed item
    │   │   │   │   │       # - Activity filters
    │   │   │   │   ├── compliance/
    │   │   │   │   │   └── index.ts  # Compliance types
    │   │   │   │   │       # - KYC case
    │   │   │   │   │       # - Business verification
    │   │   │   │   │       # - Document verification
    │   │   │   │   │       # - Privacy request
    │   │   │   │   ├── experiments/
    │   │   │   │   │   └── index.ts  # Experiment types
    │   │   │   │   │       # - Experiment definition
    │   │   │   │   │       # - Variant
    │   │   │   │   │       # - User assignment
    │   │   │   │   ├── flags/
    │   │   │   │   │   └── index.ts  # Feature flag types
    │   │   │   │   │       # - Flag definition
    │   │   │   │   │       # - Flag value
    │   │   │   │   │       # - Targeting rules
    │   │   │   │   ├── incidents/
    │   │   │   │   │   └── index.ts  # Incident types
    │   │   │   │   │       # - Incident
    │   │   │   │   │       # - Maintenance window
    │   │   │   │   │       # - System health
    │   │   │   │   ├── learning/
    │   │   │   │   │   └── index.ts  # Learning types
    │   │   │   │   │       # - Course
    │   │   │   │   │       # - Lesson
    │   │   │   │   │       # - Assessment
    │   │   │   │   │       # - Progress
    │   │   │   │   ├── moderation/
    │   │   │   │   │   └── index.ts  # Moderation types
    │   │   │   │   │       # - Report
    │   │   │   │   │       # - Moderation action
    │   │   │   │   │       # - Content check
    │   │   │   │   ├── presence/
    │   │   │   │   │   └── index.ts  # Presence types
    │   │   │   │   │       # - Online status
    │   │   │   │   │       # - Last seen
    │   │   │   │   │       # - Typing indicator
    │   │   │   │   └── sourcing/
    │   │   │   │       └── index.ts  # Sourcing types
    │   │   │   │           # - Campaign
    │   │   │   │           # - Talent pool
    │   │   │   │           # - Invitation
    │   │   │   ├── entities/  # Domain entities
    │   │   │   │   ├── contracts/
    │   │   │   │   │   ├── dispute.ts  # Dispute entity types
    │   │   │   │   │   ├── milestone.ts  # Milestone entity types
    │   │   │   │   │   ├── offer.ts  # Offer entity types
    │   │   │   │   │   └── timesheet.ts  # Timesheet entity types
    │   │   │   │   ├── financial/
    │   │   │   │   │   ├── escrow.ts  # Escrow DTOs
    │   │   │   │   │   ├── report.ts  # Billing report row DTOs
    │   │   │   │   │   └── tax.ts  # Tax DTOs
    │   │   │   │   ├── search/
    │   │   │   │   │   └── saved_query.ts  # Saved query/alert DTOs
    │   │   │   │   ├── contract.ts
    │   │   │   │   ├── invoice.ts
    │   │   │   │   ├── job.ts
    │   │   │   │   ├── message.ts
    │   │   │   │   ├── notification.ts  # Notification DTOs
    │   │   │   │   ├── proposal.ts
    │   │   │   │   ├── review.ts
    │   │   │   │   ├── subscription.ts
    │   │   │   │   ├── transaction.ts
    │   │   │   │   └── user.ts
    │   │   │   ├── enums/  # Enums
    │   │   │   │   ├── contract-status.ts
    │   │   │   │   ├── job-status.ts
    │   │   │   │   ├── payment-status.ts
    │   │   │   │   ├── proposal-status.ts
    │   │   │   │   ├── review-rating.ts
    │   │   │   │   └── user-type.ts
    │   │   │   ├── shared/
    │   │   │   │   ├── analytics/
    │   │   │   │   │   └── index.ts  # Analytics types
    │   │   │   │   │       # - Event
    │   │   │   │   │       # - Page view
    │   │   │   │   │       # - Conversion
    │   │   │   │   └── geolocation/
    │   │   │   │       └── index.ts  # Geolocation types
    │   │   │   │           # - Coordinates
    │   │   │   │           # - Location
    │   │   │   │           # - Distance
    │   │   │   └── index.ts  # Export all types
    │   │   ├── package.json
    │   │   ├── README.md
    │   │   └── tsconfig.json
    │   └── ui/  # Cross-platform component library
    │       # Cross-platform UI component library
    │       ├── src/
    │       │   ├── a11y/
    │       │   │   ├── Announcer/
    │       │   │   │   ├── LiveAnnouncer.native.tsx
    │       │   │   │   ├── LiveAnnouncer.tsx  # Live region announcements
    │       │   │   │   ├── LiveAnnouncer.types.ts
    │       │   │   │   └── LiveAnnouncer.web.tsx
    │       │   │   ├── FocusTrap/
    │       │   │   │   ├── FocusTrap.tsx
    │       │   │   │   ├── FocusTrap.types.ts
    │       │   │   │   └── FocusTrap.web.tsx
    │       │   │   ├── SkipLink/
    │       │   │   │   ├── SkipLink.tsx
    │       │   │   │   ├── SkipLink.types.ts
    │       │   │   │   └── SkipLink.web.tsx
    │       │   │   └── VisuallyHidden/
    │       │   │       ├── VisuallyHidden.native.tsx
    │       │   │       ├── VisuallyHidden.tsx
    │       │   │       ├── VisuallyHidden.types.ts
    │       │   │       └── VisuallyHidden.web.tsx
    │       │   ├── accessibility/
    │       │   │   ├── FocusTrap/
    │       │   │   │   ├── FocusTrap.native.tsx
    │       │   │   │   ├── FocusTrap.tsx
    │       │   │   │   └── FocusTrap.web.tsx
    │       │   │   ├── ScreenReaderAnnouncer/
    │       │   │   │   ├── ScreenReaderAnnouncer.native.tsx
    │       │   │   │   ├── ScreenReaderAnnouncer.tsx
    │       │   │   │   └── ScreenReaderAnnouncer.web.tsx
    │       │   │   └── SkipLinks/
    │       │   │       ├── SkipLinks.tsx
    │       │   │       └── SkipLinks.web.tsx
    │       │   ├── ai/
    │       │   │   ├── AIAssistant/
    │       │   │   │   ├── AIAssistant.native.tsx
    │       │   │   │   ├── AIAssistant.tsx
    │       │   │   │   └── AIAssistant.web.tsx
    │       │   │   ├── AutoComplete/
    │       │   │   │   ├── AIAutoComplete.native.tsx
    │       │   │   │   ├── AIAutoComplete.tsx
    │       │   │   │   └── AIAutoComplete.web.tsx
    │       │   │   └── SmartSuggestions/
    │       │   │       ├── SmartSuggestions.native.tsx
    │       │   │       ├── SmartSuggestions.tsx
    │       │   │       └── SmartSuggestions.web.tsx
    │       │   ├── auction/
    │       │   │   ├── AuctionTimer.native.tsx
    │       │   │   ├── AuctionTimer.tsx  # Countdown timer
    │       │   │   ├── AuctionTimer.web.tsx
    │       │   │   ├── BidHistoryChart.native.tsx
    │       │   │   ├── BidHistoryChart.tsx  # Bid history visualization
    │       │   │   ├── BidHistoryChart.web.tsx
    │       │   │   ├── LiveBidFeed.native.tsx
    │       │   │   ├── LiveBidFeed.tsx  # Real-time bid feed
    │       │   │   └── LiveBidFeed.web.tsx
    │       │   ├── charts/
    │       │   │   ├── EarningsChart.native.tsx
    │       │   │   ├── EarningsChart.tsx  # Earnings visualization
    │       │   │   ├── EarningsChart.web.tsx
    │       │   │   ├── PerformanceChart.native.tsx
    │       │   │   ├── PerformanceChart.tsx  # Performance metrics
    │       │   │   ├── PerformanceChart.web.tsx
    │       │   │   ├── TrendChart.native.tsx
    │       │   │   ├── TrendChart.tsx  # Trend visualization
    │       │   │   └── TrendChart.web.tsx
    │       │   ├── collaboration/
    │       │   │   ├── CollaborationPanel.native.tsx
    │       │   │   ├── CollaborationPanel.tsx  # Team collaboration
    │       │   │   ├── CollaborationPanel.web.tsx
    │       │   │   ├── GroupCard.native.tsx
    │       │   │   ├── GroupCard.tsx  # User group card
    │       │   │   ├── GroupCard.web.tsx
    │       │   │   ├── MentorCard.native.tsx
    │       │   │   ├── MentorCard.tsx  # Mentor profile card
    │       │   │   └── MentorCard.web.tsx
    │       │   ├── compliance/
    │       │   │   ├── DocumentUploader.native.tsx
    │       │   │   ├── DocumentUploader.tsx  # Compliance doc uploader
    │       │   │   ├── DocumentUploader.web.tsx
    │       │   │   ├── VerificationStatus.native.tsx
    │       │   │   ├── VerificationStatus.tsx  # Verification status badge
    │       │   │   └── VerificationStatus.web.tsx
    │       │   ├── components/
    │       │   │   ├── Accessibility/
    │       │   │   │   ├── FocusTrap/
    │       │   │   │   │   ├── FocusTrap.tsx
    │       │   │   │   │   ├── FocusTrap.types.ts
    │       │   │   │   │   └── FocusTrap.web.tsx
    │       │   │   │   ├── KeyboardShortcuts/
    │       │   │   │   │   ├── ShortcutDialog/
    │       │   │   │   │   │   └── ShortcutDialog.tsx
    │       │   │   │   │   ├── KeyboardShortcuts.tsx
    │       │   │   │   │   └── KeyboardShortcuts.types.ts
    │       │   │   │   ├── LiveRegion/
    │       │   │   │   │   ├── LiveRegion.native.tsx
    │       │   │   │   │   ├── LiveRegion.tsx
    │       │   │   │   │   ├── LiveRegion.types.ts
    │       │   │   │   │   └── LiveRegion.web.tsx
    │       │   │   │   └── ScreenReaderOnly/
    │       │   │   │       ├── ScreenReaderOnly.tsx
    │       │   │   │       └── ScreenReaderOnly.types.ts
    │       │   │   ├── Accordion/
    │       │   │   ├── AI/
    │       │   │   │   ├── AIAssistant/
    │       │   │   │   │   ├── chat/
    │       │   │   │   │   │   ├── ChatBubble.tsx
    │       │   │   │   │   │   └── ChatInput.tsx
    │       │   │   │   │   ├── ChatInterface/
    │       │   │   │   │   │   └── ChatInterface.tsx
    │       │   │   │   │   ├── AIAssistant.native.tsx
    │       │   │   │   │   ├── AIAssistant.tsx  # AI chat assistant
    │       │   │   │   │   ├── AIAssistant.types.ts
    │       │   │   │   │   └── AIAssistant.web.tsx
    │       │   │   │   ├── AutoComplete/
    │       │   │   │   │   ├── AIAutoComplete.native.tsx
    │       │   │   │   │   ├── AIAutoComplete.tsx  # AI-powered autocomplete
    │       │   │   │   │   ├── AIAutoComplete.types.ts
    │       │   │   │   │   ├── AIAutoComplete.web.tsx
    │       │   │   │   │   ├── SmartAutoComplete.native.tsx
    │       │   │   │   │   ├── SmartAutoComplete.tsx
    │       │   │   │   │   ├── SmartAutoComplete.types.ts
    │       │   │   │   │   └── SmartAutoComplete.web.tsx
    │       │   │   │   ├── ContentGeneration/
    │       │   │   │   │   ├── templates/
    │       │   │   │   │   │   └── GenerationTemplates.tsx
    │       │   │   │   │   ├── ContentGeneration.types.ts
    │       │   │   │   │   └── ContentGenerator.tsx  # AI content generation
    │       │   │   │   └── SmartSuggestions/
    │       │   │   │       ├── SuggestionCard/
    │       │   │   │       │   └── SuggestionCard.tsx
    │       │   │   │       ├── SmartSuggestions.native.tsx
    │       │   │   │       ├── SmartSuggestions.tsx  # AI suggestions
    │       │   │   │       ├── SmartSuggestions.types.ts
    │       │   │   │       └── SmartSuggestions.web.tsx
    │       │   │   ├── Alert/
    │       │   │   ├── auction/
    │       │   │   │   ├── AuctionTimer.native.tsx
    │       │   │   │   ├── AuctionTimer.tsx  # Countdown timer
    │       │   │   │   ├── AuctionTimer.web.tsx
    │       │   │   │   ├── BidHistoryChart.native.tsx
    │       │   │   │   ├── BidHistoryChart.tsx  # Bid history visualization
    │       │   │   │   ├── BidHistoryChart.web.tsx
    │       │   │   │   ├── LiveBidFeed.native.tsx
    │       │   │   │   ├── LiveBidFeed.tsx  # Real-time bid feed
    │       │   │   │   └── LiveBidFeed.web.tsx
    │       │   │   ├── Avatar/
    │       │   │   ├── Badge/
    │       │   │   ├── Breadcrumb/
    │       │   │   ├── Button/
    │       │   │   │   ├── Button.native.tsx  # Native-specific overrides
    │       │   │   │   ├── Button.stories.tsx  # Storybook stories
    │       │   │   │   ├── Button.test.tsx  # Component tests
    │       │   │   │   ├── Button.tsx  # Base button component
    │       │   │   │   └── Button.web.tsx  # Web-specific overrides
    │       │   │   ├── Calendar/
    │       │   │   │   ├── DatePicker/
    │       │   │   │   │   ├── DatePicker.native.tsx
    │       │   │   │   │   ├── DatePicker.tsx
    │       │   │   │   │   └── DatePicker.web.tsx
    │       │   │   │   ├── DateRangePicker/
    │       │   │   │   │   ├── DateRangePicker.native.tsx
    │       │   │   │   │   ├── DateRangePicker.tsx
    │       │   │   │   │   └── DateRangePicker.web.tsx
    │       │   │   │   ├── TimePicker/
    │       │   │   │   │   ├── TimePicker.native.tsx
    │       │   │   │   │   ├── TimePicker.tsx
    │       │   │   │   │   └── TimePicker.web.tsx
    │       │   │   │   ├── Calendar.native.tsx
    │       │   │   │   ├── Calendar.tsx  # Full calendar
    │       │   │   │   ├── Calendar.types.ts
    │       │   │   │   └── Calendar.web.tsx
    │       │   │   ├── Card/
    │       │   │   ├── charts/
    │       │   │   │   ├── EarningsChart.native.tsx
    │       │   │   │   ├── EarningsChart.tsx  # Earnings visualization
    │       │   │   │   ├── EarningsChart.web.tsx
    │       │   │   │   ├── PerformanceChart.native.tsx
    │       │   │   │   ├── PerformanceChart.tsx  # Performance metrics
    │       │   │   │   ├── PerformanceChart.web.tsx
    │       │   │   │   ├── TrendChart.native.tsx
    │       │   │   │   ├── TrendChart.tsx  # Trend visualization
    │       │   │   │   └── TrendChart.web.tsx
    │       │   │   ├── Charts/
    │       │   │   │   ├── FunnelChart/
    │       │   │   │   │   ├── FunnelChart.tsx
    │       │   │   │   │   └── FunnelChart.types.ts
    │       │   │   │   ├── GanttChart/
    │       │   │   │   │   ├── Task/
    │       │   │   │   │   │   └── Task.tsx
    │       │   │   │   │   ├── GanttChart.tsx
    │       │   │   │   │   └── GanttChart.types.ts
    │       │   │   │   ├── HeatMap/
    │       │   │   │   │   ├── HeatMap.tsx
    │       │   │   │   │   ├── HeatMap.types.ts
    │       │   │   │   │   └── HeatMap.web.tsx  # D3/Recharts
    │       │   │   │   └── OrgChart/
    │       │   │   │       ├── Node/
    │       │   │   │       │   └── OrgNode.tsx
    │       │   │   │       ├── OrgChart.tsx
    │       │   │   │       ├── OrgChart.types.ts
    │       │   │   │       └── OrgChart.web.tsx
    │       │   │   ├── Checkbox/
    │       │   │   ├── CodeEditor/
    │       │   │   │   ├── CodeEditor.native.tsx
    │       │   │   │   ├── CodeEditor.tsx
    │       │   │   │   ├── CodeEditor.types.ts
    │       │   │   │   └── CodeEditor.web.tsx  # Web (e.g., Monaco/CodeMirror)
    │       │   │   ├── collaboration/
    │       │   │   │   ├── CollaborationPanel.native.tsx
    │       │   │   │   ├── CollaborationPanel.tsx  # Team collaboration
    │       │   │   │   ├── CollaborationPanel.web.tsx
    │       │   │   │   ├── GroupCard.native.tsx
    │       │   │   │   ├── GroupCard.tsx  # User group card
    │       │   │   │   ├── GroupCard.web.tsx
    │       │   │   │   ├── MentorCard.native.tsx
    │       │   │   │   ├── MentorCard.tsx  # Mentor profile card
    │       │   │   │   └── MentorCard.web.tsx
    │       │   │   ├── compliance/
    │       │   │   │   ├── DocumentUploader.native.tsx
    │       │   │   │   ├── DocumentUploader.tsx  # Compliance doc uploader
    │       │   │   │   ├── DocumentUploader.web.tsx
    │       │   │   │   ├── VerificationStatus.native.tsx
    │       │   │   │   ├── VerificationStatus.tsx  # Verification status badge
    │       │   │   │   └── VerificationStatus.web.tsx
    │       │   │   ├── DataDisplay/
    │       │   │   │   ├── Kanban/
    │       │   │   │   │   ├── KanbanBoard.native.tsx  # Touch gestures
    │       │   │   │   │   ├── KanbanBoard.tsx
    │       │   │   │   │   ├── KanbanBoard.types.ts
    │       │   │   │   │   └── KanbanBoard.web.tsx  # Drag & drop
    │       │   │   │   ├── List/
    │       │   │   │   │   ├── VirtualList/
    │       │   │   │   │   │   ├── VirtualList.native.tsx  # FlashList
    │       │   │   │   │   │   ├── VirtualList.tsx
    │       │   │   │   │   │   └── VirtualList.web.tsx
    │       │   │   │   │   ├── List.native.tsx
    │       │   │   │   │   ├── List.tsx
    │       │   │   │   │   ├── List.types.ts
    │       │   │   │   │   └── List.web.tsx
    │       │   │   │   ├── Table/
    │       │   │   │   │   ├── DataTable/
    │       │   │   │   │   │   ├── DataTable.tsx  # With sorting/filtering
    │       │   │   │   │   │   └── DataTable.types.ts
    │       │   │   │   │   ├── VirtualTable/
    │       │   │   │   │   │   ├── VirtualTable.tsx  # Virtualized table
    │       │   │   │   │   │   └── VirtualTable.types.ts
    │       │   │   │   │   ├── Table.tsx  # Base table
    │       │   │   │   │   ├── Table.types.ts
    │       │   │   │   │   └── Table.web.tsx  # Full featured table
    │       │   │   │   └── Timeline/
    │       │   │   │       ├── Timeline.native.tsx
    │       │   │   │       ├── Timeline.tsx
    │       │   │   │       ├── Timeline.types.ts
    │       │   │   │       └── Timeline.web.tsx
    │       │   │   ├── DataTable/
    │       │   │   ├── DatePicker/
    │       │   │   ├── Dropdown/
    │       │   │   ├── Editor/
    │       │   │   │   ├── CodeEditor/
    │       │   │   │   │   ├── CodeEditor.native.tsx
    │       │   │   │   │   ├── CodeEditor.tsx
    │       │   │   │   │   └── CodeEditor.web.tsx  # Web (e.g., Monaco/CodeMirror)
    │       │   │   │   ├── MarkdownEditor/
    │       │   │   │   │   ├── MarkdownEditor.native.tsx
    │       │   │   │   │   ├── MarkdownEditor.tsx
    │       │   │   │   │   └── MarkdownEditor.web.tsx
    │       │   │   │   └── RichTextEditor/
    │       │   │   │       ├── RichTextEditor.native.tsx  # Native implementation
    │       │   │   │       ├── RichTextEditor.tsx
    │       │   │   │       └── RichTextEditor.web.tsx  # Web (e.g., TipTap/Slate)
    │       │   │   ├── Feedback/
    │       │   │   │   ├── Alert/
    │       │   │   │   │   ├── Alert.native.tsx
    │       │   │   │   │   ├── Alert.tsx
    │       │   │   │   │   ├── Alert.types.ts
    │       │   │   │   │   └── Alert.web.tsx
    │       │   │   │   ├── EmptyState/
    │       │   │   │   │   ├── EmptyState.native.tsx
    │       │   │   │   │   ├── EmptyState.tsx
    │       │   │   │   │   ├── EmptyState.types.ts
    │       │   │   │   │   └── EmptyState.web.tsx
    │       │   │   │   ├── Notification/
    │       │   │   │   │   ├── Notification.native.tsx
    │       │   │   │   │   ├── Notification.tsx
    │       │   │   │   │   ├── Notification.types.ts
    │       │   │   │   │   └── Notification.web.tsx  # Toast notifications
    │       │   │   │   ├── Progress/
    │       │   │   │   │   ├── ProgressBar/
    │       │   │   │   │   │   ├── ProgressBar.native.tsx
    │       │   │   │   │   │   ├── ProgressBar.tsx
    │       │   │   │   │   │   └── ProgressBar.web.tsx
    │       │   │   │   │   ├── ProgressCircle/
    │       │   │   │   │   │   ├── ProgressCircle.native.tsx
    │       │   │   │   │   │   ├── ProgressCircle.tsx
    │       │   │   │   │   │   └── ProgressCircle.web.tsx
    │       │   │   │   │   └── Progress.types.ts
    │       │   │   │   └── Skeleton/
    │       │   │   │       ├── Skeleton.native.tsx
    │       │   │   │       ├── Skeleton.tsx
    │       │   │   │       ├── Skeleton.types.ts
    │       │   │   │       └── Skeleton.web.tsx
    │       │   │   ├── FileUpload/
    │       │   │   │   ├── ImageUpload/
    │       │   │   │   │   ├── ImageCropper.tsx  # Image cropping
    │       │   │   │   │   ├── ImageUpload.native.tsx  # Camera/gallery
    │       │   │   │   │   ├── ImageUpload.tsx
    │       │   │   │   │   └── ImageUpload.web.tsx
    │       │   │   │   ├── MultiFileUpload/
    │       │   │   │   │   ├── MultiFileUpload.native.tsx
    │       │   │   │   │   ├── MultiFileUpload.tsx
    │       │   │   │   │   └── MultiFileUpload.web.tsx
    │       │   │   │   ├── FileUpload.native.tsx  # Document picker
    │       │   │   │   ├── FileUpload.tsx  # Base file upload
    │       │   │   │   ├── FileUpload.types.ts
    │       │   │   │   └── FileUpload.web.tsx  # Drag & drop support
    │       │   │   ├── FileUploader/
    │       │   │   │   ├── FileUploader.native.tsx  # Uploader (native)
    │       │   │   │   │   # - Signed URL upload flow
    │       │   │   │   │   # BE: storage-be/asset
    │       │   │   │   │   # POST /v1/storage/uploads (signed) → PUT file → POST /v1/storage/commit
    │       │   │   │   ├── FileUploader.tsx  # Uploader (shared)
    │       │   │   │   │   # - Abstraction wrapper
    │       │   │   │   │   # BE: storage-be/asset (same flow)
    │       │   │   │   └── FileUploader.web.tsx  # Uploader (web)
    │       │   │   │       # - Drag & drop, previews
    │       │   │   │       # BE: storage-be/asset (same flow)
    │       │   │   ├── Form/
    │       │   │   │   ├── CodeEditor/
    │       │   │   │   │   ├── LanguageSelector/
    │       │   │   │   │   │   └── LanguageSelector.tsx
    │       │   │   │   │   ├── syntax-highlighting/
    │       │   │   │   │   │   ├── languages.ts
    │       │   │   │   │   │   └── themes.ts
    │       │   │   │   │   ├── CodeEditor.native.tsx  # Native code editor
    │       │   │   │   │   ├── CodeEditor.tsx  # Base code editor
    │       │   │   │   │   ├── CodeEditor.types.ts
    │       │   │   │   │   └── CodeEditor.web.tsx  # Monaco/CodeMirror
    │       │   │   │   ├── DateRangePicker/
    │       │   │   │   │   ├── presets/
    │       │   │   │   │   │   └── DatePresets.tsx
    │       │   │   │   │   ├── DateRangePicker.native.tsx
    │       │   │   │   │   ├── DateRangePicker.tsx
    │       │   │   │   │   ├── DateRangePicker.types.ts
    │       │   │   │   │   └── DateRangePicker.web.tsx
    │       │   │   │   ├── MarkdownEditor/
    │       │   │   │   │   ├── Preview/
    │       │   │   │   │   │   └── MarkdownPreview.tsx
    │       │   │   │   │   ├── preview/
    │       │   │   │   │   │   └── MarkdownPreview.tsx
    │       │   │   │   │   ├── MarkdownEditor.native.tsx
    │       │   │   │   │   ├── MarkdownEditor.tsx
    │       │   │   │   │   ├── MarkdownEditor.types.ts
    │       │   │   │   │   └── MarkdownEditor.web.tsx
    │       │   │   │   ├── RichTextEditor/
    │       │   │   │   │   ├── Toolbar/
    │       │   │   │   │   │   ├── Toolbar.tsx
    │       │   │   │   │   │   └── Toolbar.types.ts
    │       │   │   │   │   ├── toolbar/
    │       │   │   │   │   │   ├── FormatButtons.tsx
    │       │   │   │   │   │   ├── InsertButtons.tsx
    │       │   │   │   │   │   └── Toolbar.tsx
    │       │   │   │   │   ├── RichTextEditor.native.tsx  # Limited rich text
    │       │   │   │   │   │   # Native implementation
    │       │   │   │   │   ├── RichTextEditor.tsx  # Base editor
    │       │   │   │   │   ├── RichTextEditor.types.ts
    │       │   │   │   │   └── RichTextEditor.web.tsx  # Web (TipTap/Slate)
    │       │   │   │   └── SignaturePad/
    │       │   │   │       ├── SignaturePad.native.tsx  # React Native Canvas
    │       │   │   │       │   # React Native Skia
    │       │   │   │       ├── SignaturePad.tsx
    │       │   │   │       ├── SignaturePad.types.ts
    │       │   │   │       └── SignaturePad.web.tsx  # Canvas-based
    │       │   │   ├── Forms/
    │       │   │   │   ├── Checkbox/
    │       │   │   │   │   ├── Checkbox.native.tsx
    │       │   │   │   │   ├── Checkbox.tsx
    │       │   │   │   │   ├── Checkbox.types.ts
    │       │   │   │   │   └── Checkbox.web.tsx
    │       │   │   │   ├── FormField/
    │       │   │   │   │   ├── FormError.tsx
    │       │   │   │   │   ├── FormField.native.tsx
    │       │   │   │   │   ├── FormField.tsx  # Form field wrapper
    │       │   │   │   │   ├── FormField.types.ts
    │       │   │   │   │   ├── FormField.web.tsx
    │       │   │   │   │   ├── FormHelperText.tsx
    │       │   │   │   │   └── FormLabel.tsx
    │       │   │   │   ├── Radio/
    │       │   │   │   │   ├── Radio.types.ts
    │       │   │   │   │   ├── RadioGroup.native.tsx
    │       │   │   │   │   ├── RadioGroup.tsx
    │       │   │   │   │   └── RadioGroup.web.tsx
    │       │   │   │   ├── Select/
    │       │   │   │   │   ├── MultiSelect/
    │       │   │   │   │   │   ├── MultiSelect.native.tsx
    │       │   │   │   │   │   ├── MultiSelect.tsx
    │       │   │   │   │   │   └── MultiSelect.web.tsx
    │       │   │   │   │   ├── Select.native.tsx
    │       │   │   │   │   ├── Select.tsx
    │       │   │   │   │   ├── Select.types.ts
    │       │   │   │   │   └── Select.web.tsx
    │       │   │   │   ├── Slider/
    │       │   │   │   │   ├── RangeSlider/
    │       │   │   │   │   │   ├── RangeSlider.native.tsx
    │       │   │   │   │   │   ├── RangeSlider.tsx
    │       │   │   │   │   │   └── RangeSlider.web.tsx
    │       │   │   │   │   ├── Slider.native.tsx
    │       │   │   │   │   ├── Slider.tsx
    │       │   │   │   │   ├── Slider.types.ts
    │       │   │   │   │   └── Slider.web.tsx
    │       │   │   │   └── Switch/
    │       │   │   │       ├── Switch.native.tsx
    │       │   │   │       ├── Switch.tsx
    │       │   │   │       ├── Switch.types.ts
    │       │   │   │       └── Switch.web.tsx
    │       │   │   ├── Input/
    │       │   │   │   ├── Input.native.tsx
    │       │   │   │   ├── Input.test.tsx
    │       │   │   │   ├── Input.tsx
    │       │   │   │   └── Input.web.tsx
    │       │   │   ├── learning/
    │       │   │   │   ├── AchievementBadge.native.tsx
    │       │   │   │   ├── AchievementBadge.tsx  # Achievement badge
    │       │   │   │   ├── AchievementBadge.web.tsx
    │       │   │   │   ├── LearningPathCard.native.tsx
    │       │   │   │   ├── LearningPathCard.tsx  # Learning path card
    │       │   │   │   ├── LearningPathCard.web.tsx
    │       │   │   │   ├── ProgressTracker.native.tsx
    │       │   │   │   ├── ProgressTracker.tsx  # Progress visualization
    │       │   │   │   └── ProgressTracker.web.tsx
    │       │   │   ├── LoadingSpinner/
    │       │   │   │   ├── LoadingSpinner.native.tsx  # Spinner (native)
    │       │   │   │   ├── LoadingSpinner.tsx  # Spinner (shared)
    │       │   │   │   └── LoadingSpinner.web.tsx  # Spinner (web)
    │       │   │   │       # BE: none (presentational)
    │       │   │   ├── Modal/
    │       │   │   │   ├── Modal.native.tsx  # Modal (native)
    │       │   │   │   ├── Modal.tsx  # Modal (shared)
    │       │   │   │   └── Modal.web.tsx  # Modal (web)
    │       │   │   │       # BE: none (presentational)
    │       │   │   ├── Navigation/
    │       │   │   │   ├── Breadcrumb/
    │       │   │   │   │   ├── Breadcrumb.tsx
    │       │   │   │   │   ├── Breadcrumb.types.ts
    │       │   │   │   │   └── Breadcrumb.web.tsx
    │       │   │   │   ├── Pagination/
    │       │   │   │   │   ├── Pagination.native.tsx
    │       │   │   │   │   ├── Pagination.tsx
    │       │   │   │   │   ├── Pagination.types.ts
    │       │   │   │   │   └── Pagination.web.tsx
    │       │   │   │   ├── Stepper/
    │       │   │   │   │   ├── Stepper.native.tsx
    │       │   │   │   │   ├── Stepper.tsx
    │       │   │   │   │   ├── Stepper.types.ts
    │       │   │   │   │   └── Stepper.web.tsx
    │       │   │   │   └── Tabs/
    │       │   │   │       ├── Tabs.native.tsx
    │       │   │   │       ├── Tabs.tsx
    │       │   │   │       ├── Tabs.types.ts
    │       │   │   │       └── Tabs.web.tsx
    │       │   │   ├── Pagination/
    │       │   │   ├── Popover/
    │       │   │   ├── Progress/
    │       │   │   ├── Radio/
    │       │   │   ├── Rating/
    │       │   │   ├── Select/
    │       │   │   ├── Skeleton/
    │       │   │   ├── Slider/
    │       │   │   ├── Stepper/
    │       │   │   ├── Switch/
    │       │   │   ├── Tabs/
    │       │   │   ├── Textarea/
    │       │   │   ├── Timeline/
    │       │   │   ├── TimePicker/
    │       │   │   ├── Toast/
    │       │   │   │   ├── Toast.native.tsx  # Toasts (native)
    │       │   │   │   ├── Toast.tsx  # Toasts (shared)
    │       │   │   │   └── Toast.web.tsx  # Toasts (web)
    │       │   │   │       # BE: none (presentational)
    │       │   │   ├── Tooltip/
    │       │   │   ├── tracking/
    │       │   │   │   ├── TimesheetTable.native.tsx
    │       │   │   │   ├── TimesheetTable.tsx  # Timesheet grid
    │       │   │   │   ├── TimesheetTable.web.tsx
    │       │   │   │   ├── TimeTracker.native.tsx
    │       │   │   │   ├── TimeTracker.tsx  # Time tracking widget
    │       │   │   │   ├── TimeTracker.web.tsx
    │       │   │   │   ├── WorkDiaryEntry.native.tsx
    │       │   │   │   ├── WorkDiaryEntry.tsx  # Work diary card
    │       │   │   │   └── WorkDiaryEntry.web.tsx
    │       │   │   ├── video/
    │       │   │   │   ├── VideoPlayer.native.tsx
    │       │   │   │   ├── VideoPlayer.tsx  # Video player
    │       │   │   │   ├── VideoPlayer.web.tsx
    │       │   │   │   ├── VideoUploader.native.tsx
    │       │   │   │   ├── VideoUploader.tsx  # Video upload
    │       │   │   │   └── VideoUploader.web.tsx
    │       │   │   └── Visualization/
    │       │   │       ├── GanttChart/
    │       │   │       │   ├── timeline/
    │       │   │       │   │   ├── Timeline.tsx
    │       │   │       │   │   └── TimelineItem.tsx
    │       │   │       │   ├── GanttChart.native.tsx
    │       │   │       │   ├── GanttChart.tsx
    │       │   │       │   ├── GanttChart.types.ts
    │       │   │       │   └── GanttChart.web.tsx
    │       │   │       ├── HeatMap/
    │       │   │       │   ├── HeatMap.native.tsx
    │       │   │       │   ├── HeatMap.tsx
    │       │   │       │   ├── HeatMap.types.ts
    │       │   │       │   └── HeatMap.web.tsx
    │       │   │       ├── KanbanBoard/
    │       │   │       │   ├── card/
    │       │   │       │   │   └── KanbanCard.tsx
    │       │   │       │   ├── column/
    │       │   │       │   │   └── KanbanColumn.tsx
    │       │   │       │   ├── KanbanBoard.native.tsx
    │       │   │       │   ├── KanbanBoard.tsx
    │       │   │       │   ├── KanbanBoard.types.ts
    │       │   │       │   └── KanbanBoard.web.tsx
    │       │   │       ├── OrgChart/
    │       │   │       │   ├── node/
    │       │   │       │   │   └── OrgNode.tsx
    │       │   │       │   ├── OrgChart.native.tsx
    │       │   │       │   ├── OrgChart.tsx
    │       │   │       │   ├── OrgChart.types.ts
    │       │   │       │   └── OrgChart.web.tsx
    │       │   │       └── TreeView/
    │       │   │           ├── node/
    │       │   │           │   └── TreeNode.tsx
    │       │   │           ├── TreeView.native.tsx
    │       │   │           ├── TreeView.tsx
    │       │   │           ├── TreeView.types.ts
    │       │   │           └── TreeView.web.tsx
    │       │   ├── forms/  # Form components
    │       │   │   ├── CodeEditor/
    │       │   │   │   ├── CodeEditor.native.tsx
    │       │   │   │   ├── CodeEditor.tsx
    │       │   │   │   └── CodeEditor.web.tsx
    │       │   │   ├── DateRangePicker/
    │       │   │   │   ├── DateRangePicker.native.tsx
    │       │   │   │   ├── DateRangePicker.tsx
    │       │   │   │   └── DateRangePicker.web.tsx
    │       │   │   ├── FormError/
    │       │   │   ├── FormField/
    │       │   │   ├── FormGroup/
    │       │   │   ├── FormHelper/
    │       │   │   ├── FormLabel/
    │       │   │   ├── MarkdownEditor/
    │       │   │   │   ├── MarkdownEditor.native.tsx
    │       │   │   │   ├── MarkdownEditor.tsx
    │       │   │   │   └── MarkdownEditor.web.tsx
    │       │   │   ├── RichTextEditor/
    │       │   │   │   ├── RichTextEditor.native.tsx
    │       │   │   │   ├── RichTextEditor.tsx
    │       │   │   │   └── RichTextEditor.web.tsx
    │       │   │   └── SignatureInput/
    │       │   │       ├── SignatureInput.native.tsx
    │       │   │       ├── SignatureInput.tsx
    │       │   │       └── SignatureInput.web.tsx
    │       │   ├── icons/  # Icon components
    │       │   │   └── index.ts  # Export all icons
    │       │   ├── layout/  # Layout components
    │       │   │   ├── Container/
    │       │   │   ├── Divider/
    │       │   │   ├── Grid/
    │       │   │   ├── Spacer/
    │       │   │   └── Stack/
    │       │   ├── learning/
    │       │   │   ├── AchievementBadge.native.tsx
    │       │   │   ├── AchievementBadge.tsx  # Achievement badge
    │       │   │   ├── AchievementBadge.web.tsx
    │       │   │   ├── LearningPathCard.native.tsx
    │       │   │   ├── LearningPathCard.tsx  # Learning path card
    │       │   │   ├── LearningPathCard.web.tsx
    │       │   │   ├── ProgressTracker.native.tsx
    │       │   │   ├── ProgressTracker.tsx  # Progress visualization
    │       │   │   └── ProgressTracker.web.tsx
    │       │   ├── tracking/
    │       │   │   ├── TimesheetTable.native.tsx
    │       │   │   ├── TimesheetTable.tsx  # Timesheet grid
    │       │   │   ├── TimesheetTable.web.tsx
    │       │   │   ├── TimeTracker.native.tsx
    │       │   │   ├── TimeTracker.tsx  # Time tracking widget
    │       │   │   ├── TimeTracker.web.tsx
    │       │   │   ├── WorkDiaryEntry.native.tsx
    │       │   │   ├── WorkDiaryEntry.tsx  # Work diary card
    │       │   │   └── WorkDiaryEntry.web.tsx
    │       │   ├── video/
    │       │   │   ├── VideoPlayer.native.tsx
    │       │   │   ├── VideoPlayer.tsx  # Video player
    │       │   │   ├── VideoPlayer.web.tsx
    │       │   │   ├── VideoUploader.native.tsx  # Video upload
    │       │   │   ├── VideoUploader.tsx  # Video upload
    │       │   │   └── VideoUploader.web.tsx
    │       │   └── visualization/
    │       │       ├── Gantt/
    │       │       │   ├── GanttChart.native.tsx
    │       │       │   ├── GanttChart.tsx
    │       │       │   └── GanttChart.web.tsx
    │       │       ├── Heatmap/
    │       │       │   ├── Heatmap.native.tsx
    │       │       │   ├── Heatmap.tsx
    │       │       │   └── Heatmap.web.tsx
    │       │       ├── Kanban/
    │       │       │   ├── KanbanBoard.native.tsx
    │       │       │   ├── KanbanBoard.tsx
    │       │       │   └── KanbanBoard.web.tsx
    │       │       └── OrgChart/
    │       │           ├── OrganizationChart.native.tsx
    │       │           ├── OrganizationChart.tsx
    │       │           └── OrganizationChart.web.tsx
    │       ├── package.json
    │       ├── README.md
    │       └── tsconfig.json
    ├── public/  # Static assets
    │   ├── animations/  # Lottie/animation files
    │   │   ├── empty-state.json
    │   │   ├── error.json
    │   │   ├── loading.json
    │   │   └── success.json
    │   ├── fonts/  # Web fonts
    │   │   ├── inter-var-latin.woff2  # Latin subset
    │   │   ├── inter-var.woff2  # Variable font
    │   │   └── noto-sans-arabic.woff2  # Arabic font
    │   ├── images/
    │   │   ├── payment-providers/  # Payment method icons
    │   │   │   ├── mastercard.svg
    │   │   │   ├── paypal.svg
    │   │   │   ├── stripe.svg
    │   │   │   └── visa.svg
    │   │   ├── social/  # Social login icons
    │   │   │   ├── github.svg
    │   │   │   ├── google.svg
    │   │   │   └── linkedin.svg
    │   │   ├── favicon.ico
    │   │   ├── hero-bg-mobile.webp  # Mobile hero
    │   │   ├── hero-bg.webp  # Hero background
    │   │   ├── logo-dark.svg  # Dark mode logo
    │   │   ├── logo.svg  # Skillsier logo
    │   │   ├── placeholder-avatar.png  # Default avatar
    │   │   └── placeholder-company.png  # Default company logo
    │   ├── locales/  # Public locale files (deprecated - moved to packages/shared)
    │   ├── browserconfig.xml  # Windows tile config
    │   ├── manifest.json  # PWA manifest
    │   ├── robots.txt  # Search engine directives
    │   ├── sitemap-dynamic.xml  # Dynamic routes sitemap
    │   ├── sitemap.xml  # Static sitemap
    │   └── sw.js  # Service worker (if using PWA)
    ├── scripts/  # Automation scripts
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
    ├── security/
    │   ├── auth-guard.ts
    │   ├── csrf-protection.ts
    │   └── rate-limiter.ts
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
    ├── .env.example  # Environment variables template
    │   # Template
    │   # Example environment variables
    ├── .env.local  # Local environment variables
    │   # Local development (git-ignored)
    ├── .env.local.example  # Local overrides template
    ├── .env.production  # Production environment
    ├── .env.staging  # Staging environment
    ├── .eslintrc.json  # Root ESLint config
    │   # Web-specific ESLint rules
    ├── .gitignore  # Git ignore patterns
    ├── .nvmrc  # Node.js version (v20.x)
    ├── .prettierignore  # Prettier ignore patterns
    ├── .prettierrc  # Prettier configuration
    ├── error.tsx  # Global error boundary
    ├── globals.css  # Global styles
    │   # - Tailwind base, components, utilities
    │   # - CSS custom properties for theming
    │   # - Dark mode variables
    ├── jest.config.js  # Root Jest configuration
    │   # Web Jest configuration
    ├── layout.tsx  # Root layout
    │   # - HTML lang attribute (i18n)
    │   # - Head meta tags
    │   # - Body layout
    │   # - Font loading
    ├── loading.tsx  # Global loading state
    ├── middleware.ts  # Next.js middleware
    │   # - Security headers
    │   # - CSRF protection
    │   # - Rate limiting (client-side)
    │   # - Auth checks
    ├── next.config.js  # Next.js configuration
    │   # - i18n config
    │   # - Image optimization
    │   # - Bundle analyzer
    │   # - Security headers
    │   # - Rewrites/redirects
    ├── not-found.tsx  # 404 page
    ├── package.json  # Root package (workspace manager)
    │   # Scripts: dev, build, test, lint, type-check
    │   # Web dependencies
    ├── page.tsx  # Root page (redirects to /[locale])
    ├── pnpm-lock.yaml  # Locked dependencies
    ├── pnpm-workspace.yaml  # pnpm workspace configuration
    ├── postcss.config.js  # PostCSS configuration
    ├── README.md  # Root README
    │   # Web app documentation
    ├── tailwind.config.js  # Tailwind configuration
    │   # - Design tokens
    │   # - Custom utilities
    │   # - Plugin configuration
    ├── tsconfig.base.json  # Base TypeScript configuration
    │   # Shared compiler options
    ├── tsconfig.json  # Root TypeScript config
    │   # Web TypeScript config
    │   # Extends tsconfig.base.json
    ├── turbo.json  # Turborepo pipeline configuration
    │   # Build cache, task dependencies
    └── │
```