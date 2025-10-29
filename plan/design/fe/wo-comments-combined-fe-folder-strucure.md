```
fe/
└── fe/  # Root monorepo
    ├── (error-boundaries)/
    │   ├── contracts/
    │   │   └── _layout.tsx  # ErrorBoundary
    │   ├── jobs/
    │   │   └── _layout.tsx  # ErrorBoundary
    │   └── proposals/
    │       └── _layout.tsx  # ErrorBoundary
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
    │   │   ├── bundle-analysis.yml  # Bundle size checks
    │   │   ├── cd-mobile-production.yml  # Mobile production deployment
    │   │   ├── cd-mobile-staging.yml  # Mobile staging deployment
    │   │   ├── cd-mobile.yml  # Mobile deployment
    │   │   ├── cd-web-production.yml  # Web production deployment
    │   │   ├── cd-web-staging.yml  # Web staging deployment
    │   │   ├── cd-web.yml  # Web deployment
    │   │   ├── ci-mobile.yml  # Mobile CI pipeline
    │   │   ├── ci-web.yml  # Web CI pipeline
    │   │   ├── ci.yml  # Continuous Integration
    │   │   ├── dependabot.yml  # Automated dependency updates
    │   │   ├── dependency-review.yml  # Dependency review
    │   │   ├── dependency-update.yml  # Dependabot automation
    │   │   ├── deploy-mobile-production.yml
    │   │   ├── deploy-mobile-staging.yml
    │   │   ├── deploy-web-production.yml
    │   │   ├── deploy-web-staging.yml
    │   │   ├── e2e-tests.yml  # E2E tests
    │   │   ├── lighthouse.yml  # Performance audits
    │   │   ├── lint.yml  # Linting
    │   │   ├── mobile-build-android.yml  # Android build
    │   │   ├── mobile-build-ios.yml  # iOS build
    │   │   ├── performance-tests.yml  # Performance testing
    │   │   ├── release.yml  # Release automation
    │   │   ├── security-scan.yml  # Security scanning
    │   │   ├── security.yml  # Security scanning
    │   │   ├── test.yml  # Test runner
    │   │   ├── type-check.yml  # TypeScript checks
    │   │   ├── visual-regression.yml  # Visual regression tests
    │   │   ├── web-deploy-production.yml  # Web production deployment
    │   │   └── web-deploy-staging.yml  # Web staging deployment
    │   └── CODEOWNERS  # Code ownership
    ├── .husky/  # Git hooks for quality gates
    │   ├── pre-commit  # Runs linting, type checking
    │   └── pre-push  # Runs full test suite before push
    ├── .vscode/  # VS Code workspace configuration
    │   ├── extensions.json  # Recommended extensions list
    │   ├── launch.json  # Debug configurations
    │   └── settings.json  # Workspace settings
    ├── apps/  # Application workspaces
    │   ├── mobile/  # React Native/Expo application
    │   │   ├── (contracts)/
    │   │   │   └── _layout.tsx  # ErrorBoundary (react-error-boundary / Sentry)
    │   │   ├── (jobs)/
    │   │   │   └── _layout.tsx  # ErrorBoundary
    │   │   ├── (proposals)/
    │   │   │   └── _layout.tsx  # ErrorBoundary
    │   │   ├── (public)/  # ENTIRE SECTION - Public/marketing pages
    │   │   │   ├── about/
    │   │   │   │   └── index.tsx  # About page (mobile)
    │   │   │   ├── blog/
    │   │   │   │   ├── [slug]/
    │   │   │   │   │   └── index.tsx  # Blog post detail
    │   │   │   │   └── index.tsx  # Blog list
    │   │   │   ├── case-studies/
    │   │   │   │   ├── [slug]/
    │   │   │   │   │   └── index.tsx  # Case study detail
    │   │   │   │   └── index.tsx  # Case studies list
    │   │   │   ├── contact/
    │   │   │   │   └── index.tsx  # Contact page
    │   │   │   ├── developer/
    │   │   │   │   ├── docs/
    │   │   │   │   │   └── index.tsx  # Developer docs
    │   │   │   │   ├── examples/
    │   │   │   │   │   └── index.tsx  # Code examples
    │   │   │   │   └── index.tsx  # Developer hub
    │   │   │   ├── help/
    │   │   │   │   ├── [slug]/
    │   │   │   │   │   └── index.tsx  # Help article detail
    │   │   │   │   └── index.tsx  # Help center
    │   │   │   ├── legal/
    │   │   │   │   ├── cookies/
    │   │   │   │   │   └── index.tsx  # Cookie policy
    │   │   │   │   ├── privacy/
    │   │   │   │   │   └── index.tsx  # Privacy policy
    │   │   │   │   └── terms/
    │   │   │   │       └── index.tsx  # Terms of service
    │   │   │   ├── pricing/
    │   │   │   │   └── index.tsx  # Pricing page
    │   │   │   └── _layout.tsx  # Public routes layout
    │   │   ├── app/  # Expo Router file-based routing
    │   │   │   ├── (admin)/  # Optional mobile admin area (if enabled later)
    │   │   │   │   ├── analytics/
    │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   └── page.tsx  # Billing analytics  # BE: admin-be/analytics
    │   │   │   │   │   ├── marketplace/
    │   │   │   │   │   │   └── page.tsx  # Marketplace metrics  # BE: admin-be/analytics
    │   │   │   │   │   ├── performance/
    │   │   │   │   │   │   ├── index.tsx  # Performance metrics
    │   │   │   │   │   │   └── page.tsx  # System performance  # BE: admin-be/analytics
    │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   └── page.tsx  # Custom reports  # BE: admin-be/reports
    │   │   │   │   │   ├── revenue/
    │   │   │   │   │   │   └── index.tsx  # Revenue dashboard
    │   │   │   │   │   ├── users/
    │   │   │   │   │   │   └── page.tsx  # User analytics  # BE: admin-be/reports/users
    │   │   │   │   │   ├── index.tsx  # Analytics overview (mobile)
    │   │   │   │   │   ├── performance.tsx  # Performance metrics (mobile)
    │   │   │   │   │   ├── revenue.tsx  # Revenue dashboard (mobile)
    │   │   │   │   │   └── │
    │   │   │   │   ├── audit-logs/
    │   │   │   │   │   ├── [logId]/
    │   │   │   │   │   │   ├── index.tsx  # Detailed audit log
    │   │   │   │   │   │   └── page.tsx  # Detailed audit log  # BE: admin-be/audit
    │   │   │   │   │   ├── index.tsx  # Audit logs list
    │   │   │   │   │   ├── page.tsx  # All audit logs
    │   │   │   │   │   └── │
    │   │   │   │   ├── billing/
    │   │   │   │   │   ├── invoices/
    │   │   │   │   │   │   └── page.tsx  # Invoice management
    │   │   │   │   │   ├── refunds/
    │   │   │   │   │   │   └── page.tsx  # Refund processing
    │   │   │   │   │   └── transactions/
    │   │   │   │   │       └── page.tsx  # Transaction monitoring
    │   │   │   │   ├── compliance/
    │   │   │   │   │   ├── audit/
    │   │   │   │   │   │   └── page.tsx  # Compliance audits
    │   │   │   │   │   ├── financial-reports/
    │   │   │   │   │   │   └── index.tsx  # Financial compliance reports
    │   │   │   │   │   ├── gdpr/
    │   │   │   │   │   │   └── page.tsx  # GDPR management
    │   │   │   │   │   ├── kyc/
    │   │   │   │   │   │   ├── [kycId]/
    │   │   │   │   │   │   │   └── index.tsx  # KYC review detail
    │   │   │   │   │   │   ├── index.tsx  # KYC queue
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   └── page.tsx  # Compliance reports
    │   │   │   │   │   ├── tax-forms/
    │   │   │   │   │   │   └── index.tsx  # Tax compliance
    │   │   │   │   │   └── │
    │   │   │   │   ├── configurations/
    │   │   │   │   │   ├── email-templates/
    │   │   │   │   │   │   └── page.tsx  # Email template management  # BE: communications-be/template
    │   │   │   │   │   ├── feature-flags/
    │   │   │   │   │   │   └── page.tsx  # Feature flag control       # BE: subscriptions-be/feature_toggles
    │   │   │   │   │   ├── rate-limits/
    │   │   │   │   │   │   └── page.tsx  # Rate limiting config
    │   │   │   │   │   └── system-settings/
    │   │   │   │   │       └── page.tsx  # System-wide settings
    │   │   │   │   ├── content-moderation/
    │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   └── page.tsx  # Moderation disputes
    │   │   │   │   │   ├── flags/
    │   │   │   │   │   │   ├── [flagId]/
    │   │   │   │   │   │   │   └── index.tsx  # Flag detail & resolution (mobile)
    │   │   │   │   │   │   ├── index.tsx  # Flagged content queue (mobile)
    │   │   │   │   │   │   └── page.tsx  # Content flags
    │   │   │   │   │   ├── queue/
    │   │   │   │   │   │   └── page.tsx  # Moderation queue
    │   │   │   │   │   └── index.tsx  # Content moderation home (mobile)
    │   │   │   │   ├── disputes/
    │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   ├── evidence/
    │   │   │   │   │   │   │   └── index.tsx  # Dispute evidence review (mobile)
    │   │   │   │   │   │   ├── messages/
    │   │   │   │   │   │   │   └── index.tsx  # Dispute messages thread (mobile)
    │   │   │   │   │   │   ├── index.tsx  # Dispute resolution (mobile)
    │   │   │   │   │   │   ├── page.tsx  # Dispute resolution
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── index.tsx  # Disputes dashboard (mobile)
    │   │   │   │   │   ├── page.tsx  # All disputes
    │   │   │   │   │   └── │
    │   │   │   │   ├── escalations/
    │   │   │   │   │   ├── [escalationId]/
    │   │   │   │   │   │   └── index.tsx  # Escalation detail
    │   │   │   │   │   ├── index.tsx  # Escalations queue
    │   │   │   │   │   └── │
    │   │   │   │   ├── feature-flags/
    │   │   │   │   │   └── [flagId]/
    │   │   │   │   │       └── page.tsx  # Feature flag detail
    │   │   │   │   ├── financial/
    │   │   │   │   │   ├── payouts/
    │   │   │   │   │   │   ├── [payoutId]/
    │   │   │   │   │   │   │   └── index.tsx  # Payout detail (mobile)
    │   │   │   │   │   │   └── index.tsx  # Payout approvals (mobile)
    │   │   │   │   │   ├── refunds/
    │   │   │   │   │   │   ├── [refundId]/
    │   │   │   │   │   │   │   └── index.tsx  # Refund detail (mobile)
    │   │   │   │   │   │   └── index.tsx  # Refund management (mobile)
    │   │   │   │   │   └── transactions/
    │   │   │   │   │       ├── [transactionId]/
    │   │   │   │   │       │   └── index.tsx  # Transaction detail (mobile)
    │   │   │   │   │       └── index.tsx  # Transaction monitoring (mobile)
    │   │   │   │   ├── financial-ops/
    │   │   │   │   │   ├── chargebacks/
    │   │   │   │   │   │   └── page.tsx  # Chargeback management
    │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   └── page.tsx  # Payment disputes
    │   │   │   │   │   ├── escrow/
    │   │   │   │   │   │   └── index.tsx  # Escrow management
    │   │   │   │   │   ├── holds/
    │   │   │   │   │   │   ├── index.tsx  # Fund holds
    │   │   │   │   │   │   └── page.tsx  # Payment holds
    │   │   │   │   │   ├── payouts/
    │   │   │   │   │   │   └── index.tsx  # Payout management
    │   │   │   │   │   ├── reconciliation/
    │   │   │   │   │   │   └── page.tsx  # Financial reconciliation
    │   │   │   │   │   ├── refunds/
    │   │   │   │   │   │   └── index.tsx  # Refund processing
    │   │   │   │   │   ├── withdrawals/
    │   │   │   │   │   │   └── index.tsx  # Withdrawal approvals
    │   │   │   │   │   └── │
    │   │   │   │   ├── jobs/
    │   │   │   │   │   ├── moderation/
    │   │   │   │   │   │   └── page.tsx  # Job moderation queue
    │   │   │   │   │   └── page.tsx  # All jobs (admin view)
    │   │   │   │   ├── kyc/
    │   │   │   │   │   ├── [caseId]/
    │   │   │   │   │   │   └── page.tsx  # KYC case review  # BE: admin-be/kyc_case
    │   │   │   │   │   └── page.tsx  # KYC queue
    │   │   │   │   ├── kyc-verification/
    │   │   │   │   │   ├── documents/
    │   │   │   │   │   │   └── page.tsx  # KYC documents
    │   │   │   │   │   └── pending/
    │   │   │   │   │       └── page.tsx  # Pending verifications
    │   │   │   │   ├── moderation/
    │   │   │   │   │   ├── content/
    │   │   │   │   │   │   ├── [contentId]/
    │   │   │   │   │   │   │   └── index.tsx  # Content review detail
    │   │   │   │   │   │   ├── index.tsx  # Content moderation queue
    │   │   │   │   │   │   ├── page.tsx  # Content moderation
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── profiles/
    │   │   │   │   │   │   └── page.tsx  # Profile reviews
    │   │   │   │   │   ├── queue/
    │   │   │   │   │   │   └── index.tsx  # Moderation queue (mobile)
    │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   ├── [reportId]/
    │   │   │   │   │   │   │   └── index.tsx  # Report detail
    │   │   │   │   │   │   ├── index.tsx  # Reports queue
    │   │   │   │   │   │   ├── page.tsx  # User reports
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── users/
    │   │   │   │   │   │   ├── [userId]/
    │   │   │   │   │   │   │   └── index.tsx  # User moderation detail
    │   │   │   │   │   │   ├── index.tsx  # User moderation queue
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── index.tsx  # Moderation overview (mobile)
    │   │   │   │   │   └── │
    │   │   │   │   ├── notifications/
    │   │   │   │   │   ├── campaigns/
    │   │   │   │   │   │   ├── [campaignId]/
    │   │   │   │   │   │   │   └── index.tsx  # Campaign detail
    │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   └── index.tsx  # Create campaign
    │   │   │   │   │   │   ├── index.tsx  # Campaigns list
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── templates/
    │   │   │   │   │   │   ├── [templateId]/
    │   │   │   │   │   │   │   └── index.tsx  # Template editor
    │   │   │   │   │   │   ├── index.tsx  # Templates list
    │   │   │   │   │   │   └── │
    │   │   │   │   │   └── │
    │   │   │   │   ├── organizations/
    │   │   │   │   │   ├── [orgId]/
    │   │   │   │   │   │   └── page.tsx  # Org detail management
    │   │   │   │   │   └── page.tsx  # All orgs
    │   │   │   │   ├── permissions/
    │   │   │   │   │   ├── roles/
    │   │   │   │   │   │   ├── [roleId]/
    │   │   │   │   │   │   │   └── index.tsx  # Role editor
    │   │   │   │   │   │   ├── index.tsx  # Roles management
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── users/
    │   │   │   │   │   │   ├── [userId]/
    │   │   │   │   │   │   │   └── index.tsx  # User permissions
    │   │   │   │   │   │   ├── index.tsx  # Admin users list
    │   │   │   │   │   │   └── │
    │   │   │   │   │   └── │
    │   │   │   │   ├── platform/
    │   │   │   │   │   ├── announcements/
    │   │   │   │   │   │   ├── [announcementId]/
    │   │   │   │   │   │   │   └── index.tsx  # Announcement editor
    │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   └── index.tsx  # Create announcement
    │   │   │   │   │   │   ├── index.tsx  # Announcements list
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── feature-flags/
    │   │   │   │   │   │   └── index.tsx  # Feature flags management
    │   │   │   │   │   ├── maintenance/
    │   │   │   │   │   │   └── index.tsx  # Maintenance mode
    │   │   │   │   │   └── │
    │   │   │   │   ├── risk-management/
    │   │   │   │   │   ├── alerts/
    │   │   │   │   │   │   └── page.tsx  # Risk alerts
    │   │   │   │   │   ├── fraud-detection/
    │   │   │   │   │   │   └── page.tsx  # Fraud detection
    │   │   │   │   │   └── reports/
    │   │   │   │   │       └── page.tsx  # Risk reports
    │   │   │   │   ├── settings/
    │   │   │   │   │   ├── feature-flags/
    │   │   │   │   │   │   ├── [flagId]/
    │   │   │   │   │   │   │   └── index.tsx  # Feature flag detail (mobile)
    │   │   │   │   │   │   └── index.tsx  # Feature flag management (mobile)
    │   │   │   │   │   └── platform/
    │   │   │   │   │       └── index.tsx  # Platform settings (mobile)
    │   │   │   │   ├── support/
    │   │   │   │   │   ├── [ticketId]/
    │   │   │   │   │   │   └── page.tsx  # Admin ticket detail
    │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   └── page.tsx  # Support analytics
    │   │   │   │   │   ├── escalations/
    │   │   │   │   │   │   └── page.tsx  # Escalated tickets
    │   │   │   │   │   ├── knowledge-base/
    │   │   │   │   │   │   └── page.tsx  # Admin KB
    │   │   │   │   │   ├── tickets/
    │   │   │   │   │   │   ├── [ticketId]/
    │   │   │   │   │   │   │   └── index.tsx  # Admin ticket detail (mobile)
    │   │   │   │   │   │   └── index.tsx  # Admin tickets list (mobile)
    │   │   │   │   │   └── index.tsx  # Support admin overview (mobile)
    │   │   │   │   ├── system/
    │   │   │   │   │   ├── announcements/
    │   │   │   │   │   │   └── page.tsx  # System announcements  # BE: communications-be/announcements
    │   │   │   │   │   ├── feature-flags/
    │   │   │   │   │   │   └── page.tsx  # Feature flags mgmt     # BE: subscriptions-be/feature_toggles
    │   │   │   │   │   └── maintenance/
    │   │   │   │   │       └── page.tsx  # Maintenance mode        # BE: utility/config
    │   │   │   │   ├── system-health/
    │   │   │   │   │   ├── incidents/
    │   │   │   │   │   │   └── page.tsx  # System incidents
    │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   └── page.tsx  # System logs
    │   │   │   │   │   ├── metrics/
    │   │   │   │   │   │   └── page.tsx  # System metrics
    │   │   │   │   │   └── performance/
    │   │   │   │   │       └── page.tsx  # Performance monitoring
    │   │   │   │   ├── users/
    │   │   │   │   │   ├── [userId]/
    │   │   │   │   │   │   ├── actions/
    │   │   │   │   │   │   │   ├── index.tsx  # User actions (mobile)
    │   │   │   │   │   │   │   └── page.tsx  # Admin user actions      # BE: admin-be/users
    │   │   │   │   │   │   ├── activity/
    │   │   │   │   │   │   │   └── index.tsx  # User activity log (mobile)
    │   │   │   │   │   │   ├── impersonate/
    │   │   │   │   │   │   │   └── index.tsx  # User impersonation (mobile)
    │   │   │   │   │   │   ├── index.tsx  # Admin user detail (mobile)
    │   │   │   │   │   │   └── page.tsx  # Admin user detail       # BE: users-be/profile
    │   │   │   │   │   ├── banned/
    │   │   │   │   │   │   └── page.tsx  # Banned users
    │   │   │   │   │   ├── bans/
    │   │   │   │   │   │   └── page.tsx  # Banned users
    │   │   │   │   │   ├── bulk-actions/
    │   │   │   │   │   │   └── page.tsx  # Bulk user actions       # BE: admin-be/users
    │   │   │   │   │   ├── bulk-operations/
    │   │   │   │   │   │   └── page.tsx  # Bulk user operations
    │   │   │   │   │   ├── exports/
    │   │   │   │   │   │   └── page.tsx  # User data exports
    │   │   │   │   │   ├── roles/
    │   │   │   │   │   │   └── page.tsx  # Role management
    │   │   │   │   │   ├── search/
    │   │   │   │   │   │   └── page.tsx  # Advanced user search
    │   │   │   │   │   └── index.tsx  # Admin users list (mobile)
    │   │   │   │   ├── _layout.tsx  # Admin layout (mobile)
    │   │   │   │   ├── index.tsx  # Admin home (mobile)
    │   │   │   │   ├── README.md  # Placeholder for future admin screens on mobile
    │   │   │   │   └── │
    │   │   │   ├── (auth)/  # Auth screens
    │   │   │   │   ├── kyc/  # COMPLETE KYC FLOW
    │   │   │   │   │   ├── _layout.tsx
    │   │   │   │   │   ├── address-proof.tsx  # Step 3: Address verification
    │   │   │   │   │   ├── business-verification.tsx  # Step 4: Business verification (if applicable)
    │   │   │   │   │   ├── identity-verification.tsx  # Step 2: ID upload & selfie
    │   │   │   │   │   ├── index.tsx  # KYC intro & requirements
    │   │   │   │   │   ├── personal-info.tsx  # Step 1: Personal information
    │   │   │   │   │   └── review.tsx  # Step 5: Review & submit
    │   │   │   │   ├── onboarding/
    │   │   │   │   │   └── kyc/
    │   │   │   │   │       └── index.tsx  # KYC onboarding (mobile)
    │   │   │   │   ├── _layout.tsx  # Auth layout
    │   │   │   │   ├── callback.tsx  # OAuth callback
    │   │   │   │   ├── forgot-password.tsx  # Password reset
    │   │   │   │   ├── login.tsx  # Login screen
    │   │   │   │   ├── register.tsx  # Registration screen
    │   │   │   │   ├── reset-password.tsx  # Reset password with code
    │   │   │   │   └── verify-email.tsx  # Email verification
    │   │   │   ├── (billing)/
    │   │   │   │   ├── payout-methods/
    │   │   │   │   │   └── index.tsx  # Payout methods (mobile)
    │   │   │   │   ├── wallet/
    │   │   │   │   │   └── index.tsx  # Wallet (mobile)
    │   │   │   │   └── … (listed above)
    │   │   │   ├── (contracts)/
    │   │   │   │   ├── offers/
    │   │   │   │   │   └── [id]/
    │   │   │   │   │       └── index.tsx  # Offer (mobile - accept/decline)
    │   │   │   │   └── _layout.tsx  # ErrorBoundary
    │   │   │   ├── (dashboard)/
    │   │   │   │   ├── activity-feed/
    │   │   │   │   │   └── index.tsx  # Activity feed
    │   │   │   │   ├── analytics/  # ENTIRE SECTION - User analytics on mobile
    │   │   │   │   │   ├── earnings/
    │   │   │   │   │   │   └── index.tsx  # Earnings analytics (mobile)
    │   │   │   │   │   ├── jobs/
    │   │   │   │   │   │   └── index.tsx  # Job analytics
    │   │   │   │   │   ├── performance/
    │   │   │   │   │   │   └── index.tsx  # Performance metrics (mobile)
    │   │   │   │   │   ├── proposals/
    │   │   │   │   │   │   └── index.tsx  # Proposal analytics
    │   │   │   │   │   ├── _layout.tsx  # Analytics layout
    │   │   │   │   │   └── index.tsx  # Analytics overview (mobile)
    │   │   │   │   ├── authorized-apps/  # ENTIRE SECTION
    │   │   │   │   │   ├── [appId]/
    │   │   │   │   │   │   ├── permissions/
    │   │   │   │   │   │   │   └── index.tsx  # - Granted scopes list
    │   │   │   │   │   │   └── index.tsx  # - App info (name, icon, description)
    │   │   │   │   │   └── index.tsx  # - Connected apps list
    │   │   │   │   ├── billing/
    │   │   │   │   │   └── subscription/
    │   │   │   │   │       ├── manage.tsx  # Manage subscription (mobile)
    │   │   │   │   │       └── plans.tsx  # Subscription plans (mobile)
    │   │   │   │   ├── budgets/  # ENTIRE SECTION - Budget tracking on mobile
    │   │   │   │   │   ├── [budgetId]/
    │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   └── index.tsx  # Edit budget (mobile)
    │   │   │   │   │   │   └── index.tsx  # Budget detail (mobile)
    │   │   │   │   │   ├── create/
    │   │   │   │   │   │   └── index.tsx  # Create budget (mobile)
    │   │   │   │   │   └── index.tsx  # Budgets list (mobile)
    │   │   │   │   ├── contracts/
    │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   └── work-diary/
    │   │   │   │   │   │       └── screenshots/
    │   │   │   │   │   │           └── [entryId].tsx  # Screenshot detail (mobile)
    │   │   │   │   │   ├── [id]/
    │   │   │   │   │   │   ├── amendments/
    │   │   │   │   │   │   │   ├── [amendmentId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Amendment detail (mobile)
    │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   └── index.tsx  # Create amendment (mobile)
    │   │   │   │   │   │   │   └── index.tsx  # Amendments list (mobile)
    │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   │   ├── evidence/
    │   │   │   │   │   │   │   │   │   ├── upload/
    │   │   │   │   │   │   │   │   │   │   └── index.tsx  # Upload evidence (mobile)
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Evidence list (mobile)
    │   │   │   │   │   │   │   │   └── index.tsx  # Dispute detail (mobile)
    │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   └── index.tsx  # Open dispute (mobile)
    │   │   │   │   │   │   │   └── index.tsx  # Disputes list (mobile)
    │   │   │   │   │   │   └── versions/
    │   │   │   │   │   │       ├── [versionId]/
    │   │   │   │   │   │       │   ├── compare/
    │   │   │   │   │   │       │   │   └── index.tsx  # Compare contract versions (mobile)
    │   │   │   │   │   │       │   └── index.tsx  # Contract version detail (mobile)
    │   │   │   │   │   │       └── index.tsx  # Contract versions list (mobile)
    │   │   │   │   │   └── page.tsx  # Contracts list
    │   │   │   │   ├── dashboard/
    │   │   │   │   │   └── page.tsx  # Dashboard home (role-based)
    │   │   │   │   ├── data-export/
    │   │   │   │   │   └── index.tsx  # - All data types selection
    │   │   │   │   ├── deliverables/
    │   │   │   │   │   ├── submit/
    │   │   │   │   │   │   └── page.tsx  # Submit deliverable (upload files, notes)
    │   │   │   │   │   └── page.tsx  # All deliverables (status, history)
    │   │   │   │   ├── feed/
    │   │   │   │   │   └── index.tsx  # Mobile feed (same BE as web)
    │   │   │   │   ├── jobs/
    │   │   │   │   │   ├── archived/
    │   │   │   │   │   │   ├── [jobId]/
    │   │   │   │   │   │   │   └── index.tsx  # Archived job detail (mobile)
    │   │   │   │   │   │   └── index.tsx  # Archived jobs list (mobile)
    │   │   │   │   │   └── drafts/
    │   │   │   │   │       ├── [draftId]/
    │   │   │   │   │       │   ├── edit/
    │   │   │   │   │       │   │   └── index.tsx  # Edit draft job (mobile)
    │   │   │   │   │       │   └── index.tsx  # Draft job detail (mobile)
    │   │   │   │   │       └── index.tsx  # Draft jobs list (mobile)
    │   │   │   │   ├── proposals/
    │   │   │   │   │   ├── drafts/
    │   │   │   │   │   │   ├── [draftId]/
    │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   └── index.tsx  # Edit draft proposal (mobile)
    │   │   │   │   │   │   │   └── index.tsx  # Draft proposal detail (mobile)
    │   │   │   │   │   │   └── index.tsx  # Draft proposals list (mobile)
    │   │   │   │   │   └── templates/
    │   │   │   │   │       ├── [templateId]/
    │   │   │   │   │       │   ├── edit/
    │   │   │   │   │       │   │   └── index.tsx  # Edit proposal template (mobile)
    │   │   │   │   │       │   ├── use/
    │   │   │   │   │       │   │   └── index.tsx  # Use template (mobile)
    │   │   │   │   │       │   └── index.tsx  # Proposal template detail (mobile)
    │   │   │   │   │       ├── create/
    │   │   │   │   │       │   └── index.tsx  # Create proposal template (mobile)
    │   │   │   │   │       └── index.tsx  # Proposal templates list (mobile)
    │   │   │   │   ├── quick-actions/
    │   │   │   │   │   └── index.tsx  # Quick actions screen
    │   │   │   │   ├── settings/
    │   │   │   │   │   ├── authorized-apps/
    │   │   │   │   │   │   ├── [appId]/
    │   │   │   │   │   │   │   ├── permissions/
    │   │   │   │   │   │   │   │   └── index.tsx  # App permissions detail (mobile)
    │   │   │   │   │   │   │   └── index.tsx  # Authorized app detail (mobile)
    │   │   │   │   │   │   └── index.tsx  # Authorized apps list (mobile)
    │   │   │   │   │   ├── developer/
    │   │   │   │   │   │   ├── api-keys/
    │   │   │   │   │   │   │   └── [keyId]/
    │   │   │   │   │   │   │       ├── logs/
    │   │   │   │   │   │   │       │   └── index.tsx  # API key usage logs (mobile)
    │   │   │   │   │   │   │       ├── permissions/
    │   │   │   │   │   │   │       │   └── index.tsx  # API key permissions (mobile)
    │   │   │   │   │   │   │       └── index.tsx  # API key detail (mobile)
    │   │   │   │   │   │   ├── oauth-apps/
    │   │   │   │   │   │   │   ├── [appId]/
    │   │   │   │   │   │   │   │   ├── credentials/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # OAuth credentials (mobile)
    │   │   │   │   │   │   │   │   ├── scopes/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # OAuth app scopes (mobile)
    │   │   │   │   │   │   │   │   └── index.tsx  # OAuth app detail (mobile)
    │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   └── index.tsx  # Create OAuth app (mobile)
    │   │   │   │   │   │   │   └── index.tsx  # OAuth apps list (mobile)
    │   │   │   │   │   │   └── webhooks/
    │   │   │   │   │   │       ├── [webhookId]/
    │   │   │   │   │   │       │   ├── deliveries/
    │   │   │   │   │   │       │   │   ├── [deliveryId]/
    │   │   │   │   │   │       │   │   │   └── index.tsx  # Webhook delivery detail (mobile)
    │   │   │   │   │   │       │   │   └── index.tsx  # Webhook deliveries (mobile)
    │   │   │   │   │   │       │   ├── test/
    │   │   │   │   │   │       │   │   └── index.tsx  # Test webhook (mobile)
    │   │   │   │   │   │       │   └── index.tsx  # Webhook detail (mobile)
    │   │   │   │   │   │       ├── create/
    │   │   │   │   │   │       │   └── index.tsx  # Create webhook (mobile)
    │   │   │   │   │   │       └── index.tsx  # Webhooks list (mobile)
    │   │   │   │   │   ├── devices/
    │   │   │   │   │   │   ├── [deviceId]/
    │   │   │   │   │   │   │   ├── revoke/
    │   │   │   │   │   │   │   │   └── index.tsx  # Revoke device (mobile)
    │   │   │   │   │   │   │   └── index.tsx  # Device detail (mobile)
    │   │   │   │   │   │   └── index.tsx  # Connected devices list (mobile)
    │   │   │   │   │   ├── integrations/
    │   │   │   │   │   │   ├── available/
    │   │   │   │   │   │   │   └── index.tsx  # Available integrations (mobile)
    │   │   │   │   │   │   ├── calendar/
    │   │   │   │   │   │   │   ├── connect/
    │   │   │   │   │   │   │   │   └── index.tsx  # Connect calendar (mobile)
    │   │   │   │   │   │   │   └── index.tsx  # Calendar integration settings (mobile)
    │   │   │   │   │   │   ├── connected/
    │   │   │   │   │   │   │   ├── [integrationId]/
    │   │   │   │   │   │   │   │   ├── disconnect/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Disconnect integration (mobile)
    │   │   │   │   │   │   │   │   └── index.tsx  # Integration detail (mobile)
    │   │   │   │   │   │   │   └── index.tsx  # Connected integrations (mobile)
    │   │   │   │   │   │   └── index.tsx  # Integrations hub (mobile)
    │   │   │   │   │   └── team/
    │   │   │   │   │       ├── categories/
    │   │   │   │   │       │   ├── [categoryId]/
    │   │   │   │   │       │   │   └── index.tsx  # KB category detail (mobile)
    │   │   │   │   │       │   └── index.tsx  # KB categories (mobile, read-only)
    │   │   │   │   │       ├── knowledge-base/
    │   │   │   │   │       │   └── articles/
    │   │   │   │   │       │       └── [articleId]/
    │   │   │   │   │       │           └── index.tsx  # KB article view (mobile, read-only)
    │   │   │   │   │       └── policies/
    │   │   │   │   │           ├── [policyId]/
    │   │   │   │   │           │   ├── attest/
    │   │   │   │   │           │   │   └── index.tsx  # Attest to policy (mobile)
    │   │   │   │   │           │   └── index.tsx  # Policy detail (mobile, read-only)
    │   │   │   │   │           └── index.tsx  # Policies list (mobile, read-only)
    │   │   │   │   ├── support/
    │   │   │   │   │   ├── help-center/
    │   │   │   │   │   │   └── index.tsx  # Help articles (mobile)
    │   │   │   │   │   └── tickets/
    │   │   │   │   │       ├── [ticketId]/
    │   │   │   │   │       │   └── index.tsx  # Ticket detail & reply
    │   │   │   │   │       └── index.tsx  # My tickets list (mobile)
    │   │   │   │   ├── today/
    │   │   │   │   │   └── index.tsx  # Today view
    │   │   │   │   └── widgets/
    │   │   │   │       ├── earnings/
    │   │   │   │       │   └── index.tsx  # Earnings widget
    │   │   │   │       ├── notifications/
    │   │   │   │       │   └── index.tsx  # Notifications widget
    │   │   │   │       └── time-tracker/
    │   │   │   │           └── index.tsx  # Time tracker widget
    │   │   │   ├── (error-boundaries)/  # ENTIRE SECTION - Error boundary wrappers
    │   │   │   │   ├── contracts/
    │   │   │   │   │   └── _layout.tsx  # - Contract ErrorBoundary wrapper
    │   │   │   │   ├── jobs/
    │   │   │   │   │   └── _layout.tsx  # - Jobs ErrorBoundary wrapper
    │   │   │   │   └── proposals/
    │   │   │   │       └── _layout.tsx  # - Proposals ErrorBoundary wrapper
    │   │   │   ├── (inbox)/
    │   │   │   │   ├── messages/
    │   │   │   │   │   ├── [conversationId]/
    │   │   │   │   │   │   └── index.tsx  # Messages (mobile - thread)
    │   │   │   │   │   └── index.tsx  # Messages (mobile - inbox)
    │   │   │   │   └── proposals/
    │   │   │   │       └── index.tsx  # Proposals inbox (mobile)
    │   │   │   ├── (jobs)/
    │   │   │   │   └── _layout.tsx  # ErrorBoundary
    │   │   │   ├── (market)/
    │   │   │   │   └── jobs/
    │   │   │   │       └── index.tsx  # Job discovery (mobile)
    │   │   │   ├── (onboarding)/  # Onboarding flow
    │   │   │   │   ├── _layout.tsx  # Onboarding layout
    │   │   │   │   ├── complete.tsx  # Onboarding complete
    │   │   │   │   ├── preferences.tsx  # Preferences
    │   │   │   │   ├── profile.tsx  # Basic profile setup
    │   │   │   │   ├── skills.tsx  # Skills selection
    │   │   │   │   └── welcome.tsx  # Welcome screen
    │   │   │   ├── (proposals)/
    │   │   │   │   └── _layout.tsx  # Proposals route error boundary
    │   │   │   ├── (public)/  # Optional: if exposing public/marketing in mobile
    │   │   │   │   ├── about/
    │   │   │   │   │   └── index.tsx  # About page (mobile)
    │   │   │   │   ├── api/
    │   │   │   │   │   ├── documentation/
    │   │   │   │   │   │   └── index.tsx  # API docs (static)
    │   │   │   │   │   ├── pricing/
    │   │   │   │   │   │   └── index.tsx  # API pricing (static)
    │   │   │   │   │   ├── sdks/
    │   │   │   │   │   │   └── index.tsx  # SDK docs (static)
    │   │   │   │   │   └── │
    │   │   │   │   ├── blog/
    │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   └── index.tsx  # Blog post detail (mobile)
    │   │   │   │   │   ├── categories/
    │   │   │   │   │   │   └── [category]/
    │   │   │   │   │   │       └── index.tsx  # Blog category (mobile)
    │   │   │   │   │   ├── category/
    │   │   │   │   │   │   └── [category]/
    │   │   │   │   │   │       └── index.tsx  # Blog category
    │   │   │   │   │   ├── index.tsx  # Blog listing (mobile)
    │   │   │   │   │   └── │
    │   │   │   │   ├── careers/
    │   │   │   │   │   ├── [jobId]/
    │   │   │   │   │   │   └── index.tsx  # Career detail (mobile)
    │   │   │   │   │   ├── [jobSlug]/
    │   │   │   │   │   │   ├── apply/
    │   │   │   │   │   │   │   └── index.tsx  # Job application form (mobile)
    │   │   │   │   │   │   ├── index.tsx  # Job posting detail (mobile)
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   └── index.tsx  # Job posting detail (static/CMS)
    │   │   │   │   │   ├── index.tsx  # Careers overview (mobile)
    │   │   │   │   │   └── │
    │   │   │   │   ├── case-studies/
    │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   └── index.tsx  # Case study detail (mobile)
    │   │   │   │   │   ├── index.tsx  # Case studies listing (mobile)
    │   │   │   │   │   └── │
    │   │   │   │   ├── community/
    │   │   │   │   │   ├── events/
    │   │   │   │   │   │   └── index.tsx  # Community events (optional BE)
    │   │   │   │   │   ├── forum/
    │   │   │   │   │   │   └── index.tsx  # Forum (if exists)
    │   │   │   │   │   ├── index.tsx  # Community home (static)
    │   │   │   │   │   └── │
    │   │   │   │   ├── contact/
    │   │   │   │   │   └── index.tsx  # Contact page (mobile)
    │   │   │   │   ├── developer/
    │   │   │   │   │   ├── docs/
    │   │   │   │   │   │   └── index.tsx  # Dev docs hub
    │   │   │   │   │   ├── examples/
    │   │   │   │   │   │   └── index.tsx  # Code examples
    │   │   │   │   │   ├── getting-started/
    │   │   │   │   │   │   └── index.tsx  # Onboarding
    │   │   │   │   │   ├── github/
    │   │   │   │   │   │   └── index.tsx  # GitHub repos
    │   │   │   │   │   ├── plugins/
    │   │   │   │   │   │   ├── shopify/
    │   │   │   │   │   │   │   └── index.tsx  # Shopify plugin docs
    │   │   │   │   │   │   ├── wordpress/
    │   │   │   │   │   │   │   └── index.tsx  # WordPress plugin docs
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── sdks/
    │   │   │   │   │   │   ├── javascript/
    │   │   │   │   │   │   │   └── index.tsx  # JS SDK
    │   │   │   │   │   │   ├── python/
    │   │   │   │   │   │   │   └── index.tsx  # Python SDK
    │   │   │   │   │   │   ├── ruby/
    │   │   │   │   │   │   │   └── index.tsx  # Ruby SDK
    │   │   │   │   │   │   ├── index.tsx  # SDK overview
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── index.tsx  # Developer docs hub
    │   │   │   │   │   └── │
    │   │   │   │   ├── developers/
    │   │   │   │   │   ├── api-reference/
    │   │   │   │   │   │   ├── [endpoint]/
    │   │   │   │   │   │   │   └── index.tsx  # API endpoint detail (static)
    │   │   │   │   │   │   └── index.tsx  # API reference home (static)
    │   │   │   │   │   ├── getting-started/
    │   │   │   │   │   │   └── index.tsx  # Developer onboarding (static)
    │   │   │   │   │   ├── guides/
    │   │   │   │   │   │   └── index.tsx  # Developer guides (static)
    │   │   │   │   │   └── index.tsx  # Developer portal (static)
    │   │   │   │   ├── enterprise/
    │   │   │   │   │   ├── contact/
    │   │   │   │   │   │   └── index.tsx  # Enterprise contact
    │   │   │   │   │   ├── demo/
    │   │   │   │   │   │   └── index.tsx  # Request demo
    │   │   │   │   │   ├── index.tsx  # Enterprise solutions (static)
    │   │   │   │   │   └── │
    │   │   │   │   ├── features/
    │   │   │   │   │   ├── [feature]/
    │   │   │   │   │   │   └── index.tsx  # Feature detail (mobile)
    │   │   │   │   │   ├── index.tsx  # Features overview (mobile)
    │   │   │   │   │   └── │
    │   │   │   │   ├── help/
    │   │   │   │   │   ├── [category]/
    │   │   │   │   │   │   ├── [articleSlug]/
    │   │   │   │   │   │   │   └── index.tsx  # Help article (mobile)
    │   │   │   │   │   │   ├── index.tsx  # Help category (mobile)
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   └── index.tsx  # Help article detail
    │   │   │   │   │   ├── index.tsx  # Help center home (mobile)
    │   │   │   │   │   └── │
    │   │   │   │   ├── how-it-works/
    │   │   │   │   │   └── index.tsx  # How it works page (mobile)
    │   │   │   │   ├── legal/
    │   │   │   │   │   ├── cookies/
    │   │   │   │   │   │   └── index.tsx  # Cookie policy
    │   │   │   │   │   ├── privacy/
    │   │   │   │   │   │   └── index.tsx  # Privacy policy
    │   │   │   │   │   ├── terms/
    │   │   │   │   │   │   └── index.tsx  # Terms of service
    │   │   │   │   │   ├── cookies.tsx  # Cookie policy (mobile)
    │   │   │   │   │   ├── index.tsx  # Legal hub (mobile)
    │   │   │   │   │   ├── privacy.tsx  # Privacy policy (mobile)
    │   │   │   │   │   ├── terms.tsx  # Terms of service (mobile)
    │   │   │   │   │   └── │
    │   │   │   │   ├── pricing/
    │   │   │   │   │   └── index.tsx  # Pricing page (mobile)
    │   │   │   │   ├── status/
    │   │   │   │   │   └── index.tsx  # System status
    │   │   │   │   ├── success-stories/
    │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   └── index.tsx  # Success story detail (mobile)
    │   │   │   │   │   ├── index.tsx  # Success stories listing (mobile)
    │   │   │   │   │   └── │
    │   │   │   │   ├── _layout.tsx  # Public layout (no auth)
    │   │   │   │   ├── index.tsx  # Public home (mobile)
    │   │   │   │   └── │
    │   │   │   ├── (shared)/
    │   │   │   │   └── deliverables/
    │   │   │   │       ├── components/
    │   │   │   │       │   └── DeliverableCard.tsx  # BE: none (typed props), actions wired to mutations
    │   │   │   │       ├── mutations.ts  # BE: contracts-be/deliverable — POST /v1/contracts/{id}/deliverables
    │   │   │   │       └── queries.ts  # BE: contracts-be/deliverable — GET /v1/contracts/{id}/deliverables
    │   │   │   ├── (tabs)/  # Bottom tabs navigation
    │   │   │   │   ├── (authenticated)/
    │   │   │   │   │   ├── admin/  # ⚠️ EXPAND THIS
    │   │   │   │   │   │   ├── financial-ops/  # THIS
    │   │   │   │   │   │   │   ├── chargebacks/
    │   │   │   │   │   │   │   │   ├── [id]/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Chargeback detail
    │   │   │   │   │   │   │   │   └── index.tsx  # Chargeback management
    │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   ├── [id]/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Dispute resolution
    │   │   │   │   │   │   │   │   └── index.tsx  # Active disputes
    │   │   │   │   │   │   │   ├── refunds/
    │   │   │   │   │   │   │   │   ├── [id]/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Process refund
    │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Refund history
    │   │   │   │   │   │   │   │   ├── [id].tsx  # Process refund
    │   │   │   │   │   │   │   │   ├── history.tsx  # Refund history
    │   │   │   │   │   │   │   │   └── index.tsx  # Refund requests
    │   │   │   │   │   │   │   ├── withdrawals/
    │   │   │   │   │   │   │   │   ├── [id]/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Review withdrawal
    │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Bulk approve withdrawals
    │   │   │   │   │   │   │   │   ├── [id].tsx  # Review withdrawal
    │   │   │   │   │   │   │   │   ├── approve.tsx  # Bulk approve
    │   │   │   │   │   │   │   │   └── index.tsx  # Pending withdrawals
    │   │   │   │   │   │   │   ├── _layout.tsx  # Financial ops layout
    │   │   │   │   │   │   │   └── index.tsx  # Financial operations dashboard
    │   │   │   │   │   │   ├── kyc/  # THIS
    │   │   │   │   │   │   │   ├── approved/
    │   │   │   │   │   │   │   │   └── index.tsx  # Approved verifications
    │   │   │   │   │   │   │   ├── flagged/
    │   │   │   │   │   │   │   │   └── index.tsx  # Flagged submissions
    │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   └── index.tsx  # Pending verifications
    │   │   │   │   │   │   │   ├── rejected/
    │   │   │   │   │   │   │   │   └── index.tsx  # Rejected verifications
    │   │   │   │   │   │   │   ├── review/
    │   │   │   │   │   │   │   │   ├── [id]/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Review KYC submission
    │   │   │   │   │   │   │   │   └── [id].tsx  # Review submission
    │   │   │   │   │   │   │   ├── _layout.tsx  # KYC section layout
    │   │   │   │   │   │   │   ├── approved.tsx  # Approved verifications
    │   │   │   │   │   │   │   ├── index.tsx  # KYC dashboard
    │   │   │   │   │   │   │   ├── pending.tsx  # Pending verifications
    │   │   │   │   │   │   │   └── rejected.tsx  # Rejected verifications
    │   │   │   │   │   │   ├── moderation/  # ⚠️ EXPAND THIS
    │   │   │   │   │   │   │   ├── bans/
    │   │   │   │   │   │   │   │   ├── [userId].tsx
    │   │   │   │   │   │   │   │   ├── appeals.tsx  # BE: admin-be/ban — GET /v1/admin/bans
    │   │   │   │   │   │   │   │   └── index.tsx
    │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   ├── [id].tsx
    │   │   │   │   │   │   │   │   ├── escalated.tsx  # BE: admin-be/dispute — GET /v1/admin/disputes
    │   │   │   │   │   │   │   │   └── index.tsx
    │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   ├── [id]/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Review moderation report
    │   │   │   │   │   │   │   │   ├── [id].tsx
    │   │   │   │   │   │   │   │   ├── index.tsx  # Moderation queue
    │   │   │   │   │   │   │   │   └── resolved.tsx  # BE: admin-be/flag — GET /v1/admin/reports
    │   │   │   │   │   │   │   ├── users/
    │   │   │   │   │   │   │   │   ├── [userId]/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # User moderation actions
    │   │   │   │   │   │   │   │   └── index.tsx  # Flagged users list
    │   │   │   │   │   │   │   ├── _layout.tsx  # Moderation layout
    │   │   │   │   │   │   │   └── index.tsx  # Moderation dashboard
    │   │   │   │   │   │   ├── _layout.tsx  # Admin section layout (if not exists)
    │   │   │   │   │   │   └── index.tsx
    │   │   │   │   │   ├── analytics/  # Advanced analytics now on mobile
    │   │   │   │   │   │   ├── clients/
    │   │   │   │   │   │   │   ├── retention/
    │   │   │   │   │   │   │   │   └── index.tsx  # Client retention
    │   │   │   │   │   │   │   └── satisfaction/
    │   │   │   │   │   │   │       └── index.tsx  # Client satisfaction
    │   │   │   │   │   │   ├── cohorts/
    │   │   │   │   │   │   │   └── index.tsx  # Cohort analysis
    │   │   │   │   │   │   ├── conversion/
    │   │   │   │   │   │   │   └── index.tsx  # Conversion funnel
    │   │   │   │   │   │   ├── custom-dashboards/
    │   │   │   │   │   │   │   └── index.tsx  # Custom dashboard builder
    │   │   │   │   │   │   ├── data-exports/
    │   │   │   │   │   │   │   └── index.tsx  # Data export tools
    │   │   │   │   │   │   ├── earnings/
    │   │   │   │   │   │   │   ├── breakdown/
    │   │   │   │   │   │   │   │   └── index.tsx  # Earnings breakdown
    │   │   │   │   │   │   │   ├── forecast/
    │   │   │   │   │   │   │   │   └── index.tsx  # Earnings forecast
    │   │   │   │   │   │   │   └── trends/
    │   │   │   │   │   │   │       └── index.tsx  # Earnings trends
    │   │   │   │   │   │   ├── exports/
    │   │   │   │   │   │   │   └── index.tsx  # Export analytics
    │   │   │   │   │   │   ├── jobs/
    │   │   │   │   │   │   │   └── index.tsx  # Job analytics (client)
    │   │   │   │   │   │   ├── overview/
    │   │   │   │   │   │   │   └── index.tsx  # Analytics overview
    │   │   │   │   │   │   ├── performance/
    │   │   │   │   │   │   │   ├── completion-rate/
    │   │   │   │   │   │   │   │   └── index.tsx  # Project completion rates
    │   │   │   │   │   │   │   ├── response-time/
    │   │   │   │   │   │   │   │   └── index.tsx  # Response time metrics
    │   │   │   │   │   │   │   └── success-rate/
    │   │   │   │   │   │   │       └── index.tsx  # Success rate analysis
    │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   ├── [id].tsx  # View report
    │   │   │   │   │   │   │   ├── create.tsx  # Custom report builder
    │   │   │   │   │   │   │   └── index.tsx  # Reports list
    │   │   │   │   │   │   ├── simple-dashboard/
    │   │   │   │   │   │   │   └── index.tsx  # Basic analytics view
    │   │   │   │   │   │   ├── spending/
    │   │   │   │   │   │   │   ├── by-category/
    │   │   │   │   │   │   │   │   └── index.tsx  # Category breakdown
    │   │   │   │   │   │   │   ├── by-freelancer/
    │   │   │   │   │   │   │   │   └── index.tsx  # Freelancer breakdown
    │   │   │   │   │   │   │   ├── forecast/
    │   │   │   │   │   │   │   │   └── index.tsx  # Spending forecast
    │   │   │   │   │   │   │   └── index.tsx  # Spending analytics (client)
    │   │   │   │   │   │   ├── user-segmentation/
    │   │   │   │   │   │   │   └── index.tsx  # User segments
    │   │   │   │   │   │   ├── _layout.tsx
    │   │   │   │   │   │   ├── contracts.tsx  # Contract metrics
    │   │   │   │   │   │   ├── earnings.tsx  # Earnings analytics
    │   │   │   │   │   │   ├── index.tsx  # Dashboard home
    │   │   │   │   │   │   ├── jobs.tsx  # Job performance
    │   │   │   │   │   │   └── performance.tsx  # Overall performance
    │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   ├── invoices/  # THIS
    │   │   │   │   │   │   │   ├── [id].tsx  # Invoice detail
    │   │   │   │   │   │   │   ├── _layout.tsx
    │   │   │   │   │   │   │   ├── create.tsx  # Create invoice
    │   │   │   │   │   │   │   ├── index.tsx  # Invoice list
    │   │   │   │   │   │   │   └── pay.tsx  # Pay invoice
    │   │   │   │   │   │   ├── payment-methods/  # THIS
    │   │   │   │   │   │   │   ├── [id]/
    │   │   │   │   │   │   │   │   ├── edit.tsx  # Edit payment method
    │   │   │   │   │   │   │   │   └── remove.tsx  # Remove payment method
    │   │   │   │   │   │   │   ├── add.tsx  # Add payment method
    │   │   │   │   │   │   │   └── index.tsx  # Payment methods list
    │   │   │   │   │   │   ├── subscriptions/
    │   │   │   │   │   │   │   └── index.tsx
    │   │   │   │   │   │   └── tax-settings/
    │   │   │   │   │   │       └── index.tsx  # THIS — Tax information & settings
    │   │   │   │   │   ├── camera/
    │   │   │   │   │   │   ├── document-scan/
    │   │   │   │   │   │   │   └── index.tsx  # Document scanner (OCR)
    │   │   │   │   │   │   ├── photo-upload/
    │   │   │   │   │   │   │   └── index.tsx  # Native photo upload
    │   │   │   │   │   │   └── qr-scan/
    │   │   │   │   │   │       └── index.tsx  # QR code scanner
    │   │   │   │   │   ├── contracts/
    │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   ├── amendments/
    │   │   │   │   │   │   │   │   ├── [amendmentId]/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Amendment detail
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Create amendment
    │   │   │   │   │   │   │   │   └── index.tsx  # Amendments list
    │   │   │   │   │   │   │   ├── change-orders/
    │   │   │   │   │   │   │   │   └── index.tsx  # Change orders
    │   │   │   │   │   │   │   └── compliance/
    │   │   │   │   │   │   │       └── index.tsx  # Compliance documents
    │   │   │   │   │   │   ├── [id]/
    │   │   │   │   │   │   │   ├── amendments/
    │   │   │   │   │   │   │   │   ├── [amendmentId]/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Amendment detail
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Create amendment
    │   │   │   │   │   │   │   │   └── index.tsx  # Amendments list
    │   │   │   │   │   │   │   ├── attachments/
    │   │   │   │   │   │   │   │   └── index.tsx  # Contract attachments
    │   │   │   │   │   │   │   ├── audit-trail/
    │   │   │   │   │   │   │   │   └── index.tsx  # Contract audit trail
    │   │   │   │   │   │   │   ├── change-orders/
    │   │   │   │   │   │   │   │   └── index.tsx  # Change orders
    │   │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   │   └── index.tsx  # Compliance check
    │   │   │   │   │   │   │   ├── deliverables/
    │   │   │   │   │   │   │   │   └── index.tsx  # View/upload deliverables
    │   │   │   │   │   │   │   ├── documents/
    │   │   │   │   │   │   │   │   └── index.tsx  # Contract documents
    │   │   │   │   │   │   │   ├── invoicing/
    │   │   │   │   │   │   │   │   └── index.tsx  # Contract invoicing
    │   │   │   │   │   │   │   ├── performance/
    │   │   │   │   │   │   │   │   └── index.tsx  # Performance metrics
    │   │   │   │   │   │   │   ├── renewal/
    │   │   │   │   │   │   │   │   └── index.tsx  # Contract renewal
    │   │   │   │   │   │   │   ├── sow/
    │   │   │   │   │   │   │   │   └── index.tsx  # Statement of Work (SOW)
    │   │   │   │   │   │   │   ├── work-diary/
    │   │   │   │   │   │   │   │   └── index.tsx  # Mobile work diary
    │   │   │   │   │   │   │   ├── milestones.tsx  # View milestones
    │   │   │   │   │   │   │   └── timesheet.tsx  # Log time
    │   │   │   │   │   │   ├── bulk-actions/
    │   │   │   │   │   │   │   └── index.tsx  # Bulk contract actions
    │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   └── index.tsx  # Contract reports
    │   │   │   │   │   │   ├── templates/
    │   │   │   │   │   │   │   ├── [templateId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Contract template
    │   │   │   │   │   │   │   └── index.tsx  # Contract templates
    │   │   │   │   │   │   ├── active.tsx  # Active contracts
    │   │   │   │   │   │   ├── completed.tsx  # Completed contracts
    │   │   │   │   │   │   └── index.tsx  # All contracts
    │   │   │   │   │   ├── enterprise/
    │   │   │   │   │   │   ├── case-studies/
    │   │   │   │   │   │   │   └── index.tsx  # Enterprise case studies
    │   │   │   │   │   │   ├── contact/
    │   │   │   │   │   │   │   └── index.tsx  # Enterprise contact/demo
    │   │   │   │   │   │   ├── pricing/
    │   │   │   │   │   │   │   └── index.tsx  # Enterprise pricing
    │   │   │   │   │   │   └── solutions/
    │   │   │   │   │   │       ├── managed-services/
    │   │   │   │   │   │       │   └── index.tsx  # Managed services
    │   │   │   │   │   │       └── staffing/
    │   │   │   │   │   │           └── index.tsx  # Staffing solutions
    │   │   │   │   │   ├── financial/
    │   │   │   │   │   │   ├── budgets/
    │   │   │   │   │   │   │   ├── [budgetId]/
    │   │   │   │   │   │   │   │   ├── alerts/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Budget alerts
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Edit budget
    │   │   │   │   │   │   │   │   └── index.tsx  # Budget detail
    │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   └── index.tsx  # Create budget
    │   │   │   │   │   │   │   └── index.tsx  # Budgets overview
    │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   └── index.tsx  # Dispute list
    │   │   │   │   │   │   ├── escrow/
    │   │   │   │   │   │   │   ├── [escrowId]/
    │   │   │   │   │   │   │   │   ├── release/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Release escrow
    │   │   │   │   │   │   │   │   └── index.tsx  # Escrow detail
    │   │   │   │   │   │   │   └── index.tsx  # Escrow accounts
    │   │   │   │   │   │   ├── invoices/  # View invoices
    │   │   │   │   │   │   │   ├── bulk-actions/
    │   │   │   │   │   │   │   │   └── index.tsx  # Bulk invoice actions
    │   │   │   │   │   │   │   ├── recurring/
    │   │   │   │   │   │   │   │   └── index.tsx  # Recurring invoices
    │   │   │   │   │   │   │   ├── templates/
    │   │   │   │   │   │   │   │   └── index.tsx  # Invoice templates
    │   │   │   │   │   │   │   └── index.tsx  # List
    │   │   │   │   │   │   ├── invoicing/
    │   │   │   │   │   │   │   ├── bulk-invoice/
    │   │   │   │   │   │   │   │   └── index.tsx  # Bulk invoice actions   # BE: financial-be/invoice
    │   │   │   │   │   │   │   ├── recurring/
    │   │   │   │   │   │   │   │   └── index.tsx  # Recurring invoices     # BE: financial-be/invoice
    │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │       └── index.tsx  # Invoice templates      # BE: financial-be/invoice
    │   │   │   │   │   │   ├── payouts/
    │   │   │   │   │   │   │   ├── [payoutId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Payout detail
    │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   └── index.tsx  # Payout history
    │   │   │   │   │   │   │   ├── request/
    │   │   │   │   │   │   │   │   └── index.tsx  # Request payout
    │   │   │   │   │   │   │   ├── schedule/
    │   │   │   │   │   │   │   │   └── index.tsx  # Payout schedule
    │   │   │   │   │   │   │   └── index.tsx  # Request payout
    │   │   │   │   │   │   ├── reconciliation/
    │   │   │   │   │   │   │   └── index.tsx  # Reconciliation dashboard
    │   │   │   │   │   │   ├── refunds/
    │   │   │   │   │   │   │   └── index.tsx  # Refunds
    │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   ├── cash-flow/
    │   │   │   │   │   │   │   │   └── index.tsx  # Cash flow reports
    │   │   │   │   │   │   │   ├── profit-loss/
    │   │   │   │   │   │   │   │   └── index.tsx  # P&L statements
    │   │   │   │   │   │   │   ├── simple/
    │   │   │   │   │   │   │   │   └── index.tsx  # Basic financial reports
    │   │   │   │   │   │   │   └── tax/
    │   │   │   │   │   │   │       └── index.tsx  # Tax reports
    │   │   │   │   │   │   ├── tax/
    │   │   │   │   │   │   │   ├── documents/
    │   │   │   │   │   │   │   │   └── index.tsx  # Tax documents
    │   │   │   │   │   │   │   ├── forms/
    │   │   │   │   │   │   │   │   └── index.tsx  # Tax forms (W-9, 1099, etc.)
    │   │   │   │   │   │   │   └── settings/
    │   │   │   │   │   │   │       └── index.tsx  # Tax settings
    │   │   │   │   │   │   ├── tax-documents/
    │   │   │   │   │   │   │   ├── [documentId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Tax document detail
    │   │   │   │   │   │   │   └── index.tsx  # Tax documents
    │   │   │   │   │   │   ├── transactions/
    │   │   │   │   │   │   │   ├── [transactionId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Transaction detail
    │   │   │   │   │   │   │   └── index.tsx  # All transactions
    │   │   │   │   │   │   ├── wallet/
    │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   └── index.tsx  # Wallet history
    │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   └── index.tsx  # Wallet settings
    │   │   │   │   │   │   │   └── index.tsx  # Wallet dashboard
    │   │   │   │   │   │   └── transactions.tsx  # Transaction history
    │   │   │   │   │   ├── jobs/
    │   │   │   │   │   │   ├── [jobId]/
    │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   └── index.tsx  # Job analytics
    │   │   │   │   │   │   │   ├── applicants/
    │   │   │   │   │   │   │   │   └── index.tsx  # View applicants (client)
    │   │   │   │   │   │   │   ├── clone/
    │   │   │   │   │   │   │   │   └── index.tsx  # Clone job
    │   │   │   │   │   │   │   ├── invite/
    │   │   │   │   │   │   │   │   └── index.tsx  # Invite candidates
    │   │   │   │   │   │   │   ├── shortlist/
    │   │   │   │   │   │   │   │   └── index.tsx  # Shortlist candidates
    │   │   │   │   │   │   │   ├── apply.tsx  # Quick apply
    │   │   │   │   │   │   │   └── details.tsx  # Job details
    │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   └── index.tsx  # Job analytics
    │   │   │   │   │   │   ├── bulk-actions/
    │   │   │   │   │   │   │   └── index.tsx  # Bulk job actions
    │   │   │   │   │   │   ├── bulk-operations/
    │   │   │   │   │   │   │   └── index.tsx  # Bulk job operations
    │   │   │   │   │   │   ├── templates/
    │   │   │   │   │   │   │   ├── [templateId]/
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Edit job template
    │   │   │   │   │   │   │   │   └── index.tsx  # Template detail
    │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   └── index.tsx  # Create job template
    │   │   │   │   │   │   │   └── index.tsx  # Job templates
    │   │   │   │   │   │   ├── create.tsx  # Simple job creation
    │   │   │   │   │   │   ├── drafts.tsx  # Drafts
    │   │   │   │   │   │   └── index.tsx  # Browse jobs (simplified)
    │   │   │   │   │   ├── learning/
    │   │   │   │   │   │   ├── achievements/
    │   │   │   │   │   │   │   └── index.tsx  # Achievements & badges
    │   │   │   │   │   │   ├── assessments/
    │   │   │   │   │   │   │   ├── [assessmentId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Take assessment
    │   │   │   │   │   │   │   └── index.tsx  # All assessments
    │   │   │   │   │   │   ├── catalog/
    │   │   │   │   │   │   │   ├── [courseId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Course detail
    │   │   │   │   │   │   │   └── index.tsx  # Course catalog
    │   │   │   │   │   │   ├── certifications/
    │   │   │   │   │   │   │   └── index.tsx  # Certifications list
    │   │   │   │   │   │   ├── communities/
    │   │   │   │   │   │   │   └── index.tsx  # Learning communities
    │   │   │   │   │   │   ├── leaderboard/
    │   │   │   │   │   │   │   └── index.tsx  # Leaderboard
    │   │   │   │   │   │   ├── mentorship/
    │   │   │   │   │   │   │   ├── [mentorId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Mentor detail
    │   │   │   │   │   │   │   ├── my-mentees/
    │   │   │   │   │   │   │   │   └── index.tsx  # My mentees
    │   │   │   │   │   │   │   ├── my-mentors/
    │   │   │   │   │   │   │   │   └── index.tsx  # My mentors
    │   │   │   │   │   │   │   └── index.tsx  # Mentorship program
    │   │   │   │   │   │   ├── my-learning/
    │   │   │   │   │   │   │   ├── completed/
    │   │   │   │   │   │   │   │   └── index.tsx  # Completed courses
    │   │   │   │   │   │   │   ├── in-progress/
    │   │   │   │   │   │   │   │   └── index.tsx  # In-progress courses
    │   │   │   │   │   │   │   └── index.tsx  # My learning dashboard
    │   │   │   │   │   │   ├── paths/
    │   │   │   │   │   │   │   ├── [pathId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Learning path detail
    │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   └── index.tsx  # Create learning path
    │   │   │   │   │   │   │   └── index.tsx  # All learning paths
    │   │   │   │   │   │   └── peer-reviews/
    │   │   │   │   │   │       └── index.tsx  # Peer reviews
    │   │   │   │   │   ├── marketplace/
    │   │   │   │   │   │   ├── add-ons/
    │   │   │   │   │   │   │   └── index.tsx  # Add-on marketplace
    │   │   │   │   │   │   ├── agencies/
    │   │   │   │   │   │   │   └── index.tsx  # Agency directory
    │   │   │   │   │   │   ├── freelancers/
    │   │   │   │   │   │   │   └── index.tsx  # Freelancer directory
    │   │   │   │   │   │   ├── integrations/
    │   │   │   │   │   │   │   └── index.tsx  # Integration marketplace
    │   │   │   │   │   │   ├── services/
    │   │   │   │   │   │   │   └── index.tsx  # Service catalog
    │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │       └── index.tsx  # Template marketplace
    │   │   │   │   │   ├── messages/
    │   │   │   │   │   │   └── [conversationId]/
    │   │   │   │   │   │       ├── files/
    │   │   │   │   │   │       │   └── index.tsx  # Conversation files
    │   │   │   │   │   │       ├── participants/
    │   │   │   │   │   │       │   └── index.tsx  # Manage participants
    │   │   │   │   │   │       └── settings/
    │   │   │   │   │   │           └── index.tsx  # Conversation settings
    │   │   │   │   │   ├── notifications/
    │   │   │   │   │   │   └── index.tsx  # Native notifications center
    │   │   │   │   │   ├── offline/
    │   │   │   │   │   │   ├── queue/
    │   │   │   │   │   │   │   └── index.tsx  # Offline queue
    │   │   │   │   │   │   └── sync/
    │   │   │   │   │   │       └── index.tsx  # Sync management
    │   │   │   │   │   ├── organization/
    │   │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   └── index.tsx  # Billing history
    │   │   │   │   │   │   │   ├── payment-methods/
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Add payment method
    │   │   │   │   │   │   │   │   └── index.tsx  # Payment methods
    │   │   │   │   │   │   │   └── subscription/
    │   │   │   │   │   │   │       ├── change/
    │   │   │   │   │   │   │       │   └── index.tsx  # Change subscription
    │   │   │   │   │   │   │       └── index.tsx  # Organization subscription
    │   │   │   │   │   │   ├── roles/
    │   │   │   │   │   │   │   └── index.tsx  # Roles management
    │   │   │   │   │   │   └── team/
    │   │   │   │   │   │       ├── [memberId]/
    │   │   │   │   │   │       │   ├── permissions/
    │   │   │   │   │   │       │   │   └── index.tsx  # Member permissions
    │   │   │   │   │   │       │   └── projects/
    │   │   │   │   │   │       │       └── index.tsx  # Member projects
    │   │   │   │   │   │       ├── bulk-actions/
    │   │   │   │   │   │       │   └── index.tsx  # Bulk team actions
    │   │   │   │   │   │       ├── invite/
    │   │   │   │   │   │       │   ├── bulk/
    │   │   │   │   │   │       │   │   └── index.tsx  # Bulk invite
    │   │   │   │   │   │       │   └── index.tsx  # Invite team member
    │   │   │   │   │   │       └── members/
    │   │   │   │   │   │           └── index.tsx  # View team members
    │   │   │   │   │   ├── portfolio/
    │   │   │   │   │   │   ├── [portfolioId]/
    │   │   │   │   │   │   │   └── index.tsx  # Portfolio item detail
    │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   └── index.tsx  # Portfolio analytics
    │   │   │   │   │   │   ├── collections/
    │   │   │   │   │   │   │   ├── [collectionId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Collection detail
    │   │   │   │   │   │   │   └── index.tsx  # All collections
    │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   └── index.tsx  # Create portfolio item
    │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   └── index.tsx  # Edit portfolio
    │   │   │   │   │   │   └── index.tsx  # Portfolio home
    │   │   │   │   │   ├── profile/
    │   │   │   │   │   │   ├── activity/
    │   │   │   │   │   │   │   └── index.tsx  # Activity history
    │   │   │   │   │   │   ├── availability/
    │   │   │   │   │   │   │   ├── calendar/
    │   │   │   │   │   │   │   │   └── index.tsx  # Availability calendar
    │   │   │   │   │   │   │   └── settings/
    │   │   │   │   │   │   │       └── index.tsx  # Availability settings
    │   │   │   │   │   │   ├── badges/
    │   │   │   │   │   │   │   └── index.tsx  # Earned badges
    │   │   │   │   │   │   ├── bio/
    │   │   │   │   │   │   │   └── index.tsx  # Bio editor
    │   │   │   │   │   │   ├── certifications/
    │   │   │   │   │   │   │   ├── [certId]/
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Edit certification
    │   │   │   │   │   │   │   │   └── verify/
    │   │   │   │   │   │   │   │       └── index.tsx  # Verify certification
    │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   └── index.tsx  # Add certification
    │   │   │   │   │   │   │   └── index.tsx  # Certifications
    │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   ├── bio.tsx  # Quick bio edit
    │   │   │   │   │   │   │   ├── experience.tsx  # Experience
    │   │   │   │   │   │   │   ├── portfolio.tsx  # Portfolio (simple)
    │   │   │   │   │   │   │   └── skills.tsx  # Skills
    │   │   │   │   │   │   ├── education/
    │   │   │   │   │   │   │   └── index.tsx  # Education
    │   │   │   │   │   │   ├── employment/
    │   │   │   │   │   │   │   └── index.tsx  # Employment
    │   │   │   │   │   │   ├── languages/
    │   │   │   │   │   │   │   └── index.tsx  # Language proficiency
    │   │   │   │   │   │   ├── media/
    │   │   │   │   │   │   │   └── index.tsx  # Media gallery
    │   │   │   │   │   │   ├── public-view/
    │   │   │   │   │   │   │   └── index.tsx  # Public profile preview
    │   │   │   │   │   │   ├── publications/
    │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   └── index.tsx  # Add publication
    │   │   │   │   │   │   │   └── index.tsx  # Publications
    │   │   │   │   │   │   ├── skills/
    │   │   │   │   │   │   │   └── index.tsx  # Skills management
    │   │   │   │   │   │   ├── social/
    │   │   │   │   │   │   │   └── index.tsx  # Social links
    │   │   │   │   │   │   ├── stats/
    │   │   │   │   │   │   │   └── index.tsx  # Profile statistics
    │   │   │   │   │   │   ├── testimonials/
    │   │   │   │   │   │   │   └── index.tsx  # Testimonials
    │   │   │   │   │   │   ├── video-intro/
    │   │   │   │   │   │   │   └── index.tsx  # Video introduction
    │   │   │   │   │   │   ├── visibility/
    │   │   │   │   │   │   │   └── index.tsx  # Profile visibility settings
    │   │   │   │   │   │   ├── index.tsx  # Profile view
    │   │   │   │   │   │   ├── reviews.tsx  # Reviews
    │   │   │   │   │   │   └── settings.tsx  # Basic settings
    │   │   │   │   │   ├── projects/
    │   │   │   │   │   │   ├── spaces/
    │   │   │   │   │   │   │   ├── [spaceId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Workspace detail
    │   │   │   │   │   │   │   └── index.tsx  # All workspaces
    │   │   │   │   │   │   └── wiki/
    │   │   │   │   │   │       ├── [pageId]/
    │   │   │   │   │   │       │   └── index.tsx  # Wiki page
    │   │   │   │   │   │       └── index.tsx  # Wiki home
    │   │   │   │   │   ├── proposals/
    │   │   │   │   │   │   ├── [proposalId]/
    │   │   │   │   │   │   │   ├── compare/
    │   │   │   │   │   │   │   │   └── index.tsx  # Compare versions
    │   │   │   │   │   │   │   ├── details.tsx  # Proposal details
    │   │   │   │   │   │   │   ├── edit.tsx  # Quick edit
    │   │   │   │   │   │   │   └── withdraw.tsx  # Withdraw
    │   │   │   │   │   │   ├── bulk-actions/
    │   │   │   │   │   │   │   └── index.tsx  # Bulk proposal actions
    │   │   │   │   │   │   ├── templates/  # Proposal templates (mobile)
    │   │   │   │   │   │   │   ├── [templateId]/
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Edit proposal template
    │   │   │   │   │   │   │   │   └── index.tsx  # Template detail
    │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   └── index.tsx  # Create proposal template
    │   │   │   │   │   │   │   ├── quick/
    │   │   │   │   │   │   │   │   └── index.tsx  # Quick proposal templates
    │   │   │   │   │   │   │   └── index.tsx  # Proposal templates
    │   │   │   │   │   │   ├── active.tsx  # Active proposals
    │   │   │   │   │   │   ├── archived.tsx  # Archived proposals
    │   │   │   │   │   │   └── drafts.tsx  # Drafts
    │   │   │   │   │   ├── search/
    │   │   │   │   │   │   ├── advanced/
    │   │   │   │   │   │   │   └── index.tsx  # Advanced search interface
    │   │   │   │   │   │   ├── alerts/
    │   │   │   │   │   │   │   ├── [alertId]/
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Edit search alert
    │   │   │   │   │   │   │   │   └── index.tsx  # Alert detail
    │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   └── index.tsx  # Create search alert
    │   │   │   │   │   │   │   └── index.tsx  # Search alerts
    │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   └── index.tsx  # Search history
    │   │   │   │   │   │   ├── saved/
    │   │   │   │   │   │   │   ├── [queryId]/
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Edit saved query
    │   │   │   │   │   │   │   │   └── index.tsx  # Run saved query
    │   │   │   │   │   │   │   ├── [savedId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Saved search detail
    │   │   │   │   │   │   │   └── index.tsx  # Saved searches (full mgmt)
    │   │   │   │   │   │   ├── filters.tsx  # Basic filters
    │   │   │   │   │   │   ├── freelancers.tsx  # Search freelancers
    │   │   │   │   │   │   ├── jobs.tsx  # Search jobs
    │   │   │   │   │   │   └── saved.tsx  # Saved searches
    │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   ├── accessibility/
    │   │   │   │   │   │   │   └── index.tsx  # Accessibility settings
    │   │   │   │   │   │   ├── advanced/
    │   │   │   │   │   │   │   ├── api-access/
    │   │   │   │   │   │   │   │   ├── create-token/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Create API token
    │   │   │   │   │   │   │   │   └── index.tsx  # API access management
    │   │   │   │   │   │   │   ├── developer/
    │   │   │   │   │   │   │   │   └── index.tsx  # Developer tools
    │   │   │   │   │   │   │   └── webhooks/
    │   │   │   │   │   │   │       ├── [webhookId]/
    │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │       │   │   └── index.tsx  # Edit webhook
    │   │   │   │   │   │   │       │   ├── logs/
    │   │   │   │   │   │   │       │   │   └── index.tsx  # Webhook logs
    │   │   │   │   │   │   │       │   └── index.tsx  # Webhook detail
    │   │   │   │   │   │   │       ├── create/
    │   │   │   │   │   │   │       │   └── index.tsx  # Create webhook
    │   │   │   │   │   │   │       └── index.tsx  # Webhooks list
    │   │   │   │   │   │   ├── api/
    │   │   │   │   │   │   │   ├── documentation/
    │   │   │   │   │   │   │   │   └── index.tsx  # API documentation
    │   │   │   │   │   │   │   ├── keys/
    │   │   │   │   │   │   │   │   └── index.tsx  # API key management     # BE: users-be/api_keys
    │   │   │   │   │   │   │   └── webhooks/
    │   │   │   │   │   │   │       └── index.tsx  # Webhook configuration  # BE: users-be/integrations
    │   │   │   │   │   │   ├── app-preferences/
    │   │   │   │   │   │   │   ├── data-usage/
    │   │   │   │   │   │   │   │   └── index.tsx  # Data usage controls
    │   │   │   │   │   │   │   ├── notifications/
    │   │   │   │   │   │   │   │   ├── channels/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Notification channels (Android)
    │   │   │   │   │   │   │   │   └── do-not-disturb/
    │   │   │   │   │   │   │   │       └── index.tsx  # DND schedule
    │   │   │   │   │   │   │   └── theme/
    │   │   │   │   │   │   │       └── index.tsx  # Theme settings
    │   │   │   │   │   │   ├── authorized-apps/
    │   │   │   │   │   │   │   ├── [appId]/
    │   │   │   │   │   │   │   │   ├── permissions/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # App permissions detail
    │   │   │   │   │   │   │   │   └── index.tsx  # Authorized app detail
    │   │   │   │   │   │   │   └── index.tsx  # Authorized apps list
    │   │   │   │   │   │   ├── automation/
    │   │   │   │   │   │   │   ├── rules/
    │   │   │   │   │   │   │   │   └── index.tsx  # Automation rules
    │   │   │   │   │   │   │   └── workflows/
    │   │   │   │   │   │   │       └── index.tsx  # Workflow automation
    │   │   │   │   │   │   ├── blocked-users/
    │   │   │   │   │   │   │   └── index.tsx  # Blocked users list
    │   │   │   │   │   │   ├── connected-accounts/
    │   │   │   │   │   │   │   └── index.tsx  # Connected accounts
    │   │   │   │   │   │   ├── data-export/
    │   │   │   │   │   │   │   └── index.tsx  # Data export (comprehensive)
    │   │   │   │   │   │   ├── developer/  # ENTIRE SECTION - Developer tools (mobile)
    │   │   │   │   │   │   │   ├── api-keys/
    │   │   │   │   │   │   │   │   ├── [keyId]/
    │   │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   │   └── index.tsx  # API key usage logs (mobile)
    │   │   │   │   │   │   │   │   │   ├── permissions/
    │   │   │   │   │   │   │   │   │   │   └── index.tsx  # API key permissions (mobile)
    │   │   │   │   │   │   │   │   │   └── index.tsx  # API key detail (mobile)
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Create API key (mobile)
    │   │   │   │   │   │   │   │   └── index.tsx  # API keys list (mobile)
    │   │   │   │   │   │   │   ├── oauth-apps/
    │   │   │   │   │   │   │   │   ├── [appId]/
    │   │   │   │   │   │   │   │   │   ├── authorizations/
    │   │   │   │   │   │   │   │   │   │   └── index.tsx  # OAuth authorizations (mobile)
    │   │   │   │   │   │   │   │   │   ├── credentials/
    │   │   │   │   │   │   │   │   │   │   └── index.tsx  # OAuth credentials (mobile)
    │   │   │   │   │   │   │   │   │   ├── scopes/
    │   │   │   │   │   │   │   │   │   │   └── index.tsx  # OAuth app scopes (mobile)
    │   │   │   │   │   │   │   │   │   └── index.tsx  # OAuth app detail (mobile)
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Create OAuth app (mobile)
    │   │   │   │   │   │   │   │   └── index.tsx  # OAuth apps list (mobile)
    │   │   │   │   │   │   │   ├── sandbox/
    │   │   │   │   │   │   │   │   └── index.tsx  # API sandbox/testing (mobile)
    │   │   │   │   │   │   │   ├── usage/
    │   │   │   │   │   │   │   │   └── index.tsx  # API usage dashboard (mobile)
    │   │   │   │   │   │   │   ├── webhooks/
    │   │   │   │   │   │   │   │   ├── [webhookId]/
    │   │   │   │   │   │   │   │   │   ├── deliveries/
    │   │   │   │   │   │   │   │   │   │   ├── [deliveryId]/
    │   │   │   │   │   │   │   │   │   │   │   └── index.tsx  # Webhook delivery detail (mobile)
    │   │   │   │   │   │   │   │   │   │   └── index.tsx  # Webhook deliveries (mobile)
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── index.tsx  # Edit webhook (mobile)
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Webhook detail (mobile)
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Create webhook (mobile)
    │   │   │   │   │   │   │   │   └── index.tsx  # Webhooks list (mobile)
    │   │   │   │   │   │   │   ├── _layout.tsx  # Developer section layout
    │   │   │   │   │   │   │   └── index.tsx  # Developer settings hub (mobile)
    │   │   │   │   │   │   ├── devices/
    │   │   │   │   │   │   │   ├── [deviceId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Device detail
    │   │   │   │   │   │   │   └── index.tsx  # Active devices list
    │   │   │   │   │   │   ├── integrations/
    │   │   │   │   │   │   │   ├── [integrationId]/
    │   │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Integration settings
    │   │   │   │   │   │   │   │   └── index.tsx  # Integration detail
    │   │   │   │   │   │   │   ├── available/
    │   │   │   │   │   │   │   │   └── index.tsx  # Available integrations
    │   │   │   │   │   │   │   ├── calendar/
    │   │   │   │   │   │   │   │   └── index.tsx  # Calendar integration
    │   │   │   │   │   │   │   ├── connected/
    │   │   │   │   │   │   │   │   └── index.tsx  # Connected integrations
    │   │   │   │   │   │   │   ├── storage/
    │   │   │   │   │   │   │   │   └── index.tsx  # Cloud storage integration
    │   │   │   │   │   │   │   └── index.tsx  # Integrations hub
    │   │   │   │   │   │   └── security-questions/
    │   │   │   │   │   │       └── index.tsx  # Security questions
    │   │   │   │   │   ├── support/
    │   │   │   │   │   │   ├── help-center/  # ⚠️ EXPAND THIS
    │   │   │   │   │   │   │   ├── article/
    │   │   │   │   │   │   │   │   └── [id].tsx  # Article detail
    │   │   │   │   │   │   │   ├── categories/
    │   │   │   │   │   │   │   │   └── [slug].tsx  # — Category articles
    │   │   │   │   │   │   │   ├── index.tsx
    │   │   │   │   │   │   │   └── search.tsx  # — Help search
    │   │   │   │   │   │   ├── knowledge-base/
    │   │   │   │   │   │   │   └── category/
    │   │   │   │   │   │   │       └── [categoryId]/
    │   │   │   │   │   │   │           └── index.tsx  # KB category
    │   │   │   │   │   │   ├── live-chat/
    │   │   │   │   │   │   │   └── index.tsx  # Live chat support
    │   │   │   │   │   │   ├── tickets/  # (basic)
    │   │   │   │   │   │   │   ├── [ticketId]/
    │   │   │   │   │   │   │   │   ├── escalate/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Escalate ticket
    │   │   │   │   │   │   │   │   └── history/
    │   │   │   │   │   │   │   │       └── index.tsx  # Ticket history
    │   │   │   │   │   │   │   ├── [id].tsx
    │   │   │   │   │   │   │   └── index.tsx
    │   │   │   │   │   │   ├── contact.tsx  # THIS — Contact support
    │   │   │   │   │   │   └── faq.tsx  # THIS — FAQ
    │   │   │   │   │   ├── talent/
    │   │   │   │   │   │   ├── discovery/
    │   │   │   │   │   │   │   ├── advanced-filters/
    │   │   │   │   │   │   │   │   └── index.tsx  # Advanced filtering
    │   │   │   │   │   │   │   ├── ai-matching/
    │   │   │   │   │   │   │   │   └── index.tsx  # AI-powered matching    # BE: search-be/query
    │   │   │   │   │   │   │   └── index.tsx  # Talent discovery
    │   │   │   │   │   │   ├── invites/
    │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   └── index.tsx  # Pending invites
    │   │   │   │   │   │   │   └── sent/
    │   │   │   │   │   │   │       └── index.tsx  # Sent invites
    │   │   │   │   │   │   ├── saved/
    │   │   │   │   │   │   │   ├── [listId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Saved list
    │   │   │   │   │   │   │   └── index.tsx  # All saved lists
    │   │   │   │   │   │   ├── shortlists/
    │   │   │   │   │   │   │   ├── [shortlistId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Shortlist detail
    │   │   │   │   │   │   │   └── index.tsx  # All shortlists
    │   │   │   │   │   │   └── sourcing/
    │   │   │   │   │   │       ├── campaigns/
    │   │   │   │   │   │       │   └── index.tsx  # Sourcing campaigns
    │   │   │   │   │   │       └── pipelines/
    │   │   │   │   │   │           └── index.tsx  # Hiring pipelines
    │   │   │   │   │   ├── widgets/  # THIS ENTIRE SECTION
    │   │   │   │   │   │   ├── [id]/
    │   │   │   │   │   │   │   └── edit.tsx  # Edit widget settings
    │   │   │   │   │   │   └── index.tsx  # Widget dashboard & configuration
    │   │   │   │   │   ├── work/
    │   │   │   │   │   │   └── quick-actions/
    │   │   │   │   │   │       └── index.tsx  # Home screen quick actions
    │   │   │   │   │   ├── _layout.tsx  # Bottom tab navigator
    │   │   │   │   │   └── messages.tsx  # Messages tab
    │   │   │   │   ├── account/
    │   │   │   │   │   ├── developer/
    │   │   │   │   │   │   ├── api-keys/
    │   │   │   │   │   │   │   ├── [keyId]/
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Edit API key (mobile)
    │   │   │   │   │   │   │   │   ├── index.tsx  # API key detail (mobile)
    │   │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   └── index.tsx  # Create API key (mobile)
    │   │   │   │   │   │   ├── index.tsx  # API keys list (mobile)
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── oauth-apps/
    │   │   │   │   │   │   ├── [appId]/
    │   │   │   │   │   │   │   ├── credentials/
    │   │   │   │   │   │   │   │   └── index.tsx  # OAuth app credentials (mobile)
    │   │   │   │   │   │   │   ├── scopes/
    │   │   │   │   │   │   │   │   └── index.tsx  # OAuth app scopes (mobile)
    │   │   │   │   │   │   │   ├── index.tsx  # OAuth app detail (mobile)
    │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   └── index.tsx  # Create OAuth app (mobile)
    │   │   │   │   │   │   ├── index.tsx  # OAuth apps list (mobile)
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── sandbox/
    │   │   │   │   │   │   └── index.tsx  # API sandbox/tester (mobile)
    │   │   │   │   │   ├── usage/
    │   │   │   │   │   │   └── index.tsx  # API usage dashboard (mobile)
    │   │   │   │   │   ├── webhooks/
    │   │   │   │   │   │   ├── [webhookId]/
    │   │   │   │   │   │   │   ├── deliveries/
    │   │   │   │   │   │   │   │   ├── [deliveryId]/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Webhook delivery detail (mobile)
    │   │   │   │   │   │   │   │   ├── index.tsx  # Webhook deliveries list (mobile)
    │   │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   └── index.tsx  # Edit webhook (mobile)
    │   │   │   │   │   │   │   ├── events/
    │   │   │   │   │   │   │   │   └── index.tsx  # Webhook events config (mobile)
    │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   └── index.tsx  # Webhook logs (mobile)
    │   │   │   │   │   │   │   ├── test/
    │   │   │   │   │   │   │   │   └── index.tsx  # Test webhook (mobile)
    │   │   │   │   │   │   │   ├── index.tsx  # Webhook detail (mobile)
    │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   └── index.tsx  # Create webhook (mobile)
    │   │   │   │   │   │   ├── index.tsx  # Webhooks list (mobile)
    │   │   │   │   │   │   └── │
    │   │   │   │   │   └── │
    │   │   │   │   ├── admin/
    │   │   │   │   │   ├── kyc/
    │   │   │   │   │   │   ├── [kycId].tsx  # KYC case detail (mobile)
    │   │   │   │   │   │   └── index.tsx  # KYC queue (mobile)
    │   │   │   │   │   ├── moderation/
    │   │   │   │   │   │   ├── [reportId].tsx  # Report detail (mobile)
    │   │   │   │   │   │   └── index.tsx  # Moderation queue (mobile)
    │   │   │   │   │   ├── users/
    │   │   │   │   │   │   ├── [userId].tsx  # User detail (mobile admin)
    │   │   │   │   │   │   └── index.tsx  # Users search (mobile)
    │   │   │   │   │   ├── _layout.tsx  # Admin tab layout
    │   │   │   │   │   └── index.tsx  # Admin dashboard (mobile)
    │   │   │   │   ├── analytics/  # - Analytics section (mobile)
    │   │   │   │   │   ├── contracts/
    │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   └── index.tsx  # Contract analytics detail
    │   │   │   │   │   │   └── index.tsx  # Contracts analytics overview
    │   │   │   │   │   ├── earnings/
    │   │   │   │   │   │   ├── [periodId]/
    │   │   │   │   │   │   │   └── index.tsx  # Period earnings breakdown
    │   │   │   │   │   │   ├── breakdown/
    │   │   │   │   │   │   │   └── index.tsx  # Earnings breakdown by category
    │   │   │   │   │   │   ├── forecast/
    │   │   │   │   │   │   │   └── index.tsx  # Earnings forecast
    │   │   │   │   │   │   └── index.tsx  # Earnings analytics overview
    │   │   │   │   │   ├── jobs/
    │   │   │   │   │   │   ├── [jobId]/
    │   │   │   │   │   │   │   └── index.tsx  # Job analytics detail
    │   │   │   │   │   │   └── index.tsx  # Jobs analytics overview
    │   │   │   │   │   ├── performance/
    │   │   │   │   │   │   ├── quality-score/
    │   │   │   │   │   │   │   └── index.tsx  # Quality score breakdown
    │   │   │   │   │   │   ├── success-rate/
    │   │   │   │   │   │   │   └── index.tsx  # Success rate analytics
    │   │   │   │   │   │   └── index.tsx  # Performance overview
    │   │   │   │   │   ├── profile/
    │   │   │   │   │   │   └── index.tsx  # Profile analytics
    │   │   │   │   │   ├── proposals/
    │   │   │   │   │   │   ├── [proposalId]/
    │   │   │   │   │   │   │   └── index.tsx  # Proposal analytics detail
    │   │   │   │   │   │   ├── conversion/
    │   │   │   │   │   │   │   └── index.tsx  # Proposal conversion analytics
    │   │   │   │   │   │   ├── success-rate/
    │   │   │   │   │   │   │   └── index.tsx  # Proposal success rate
    │   │   │   │   │   │   └── index.tsx  # Proposals analytics overview
    │   │   │   │   │   ├── _layout.tsx  # Analytics stack navigator
    │   │   │   │   │   └── index.tsx  # Analytics dashboard home
    │   │   │   │   ├── browse/
    │   │   │   │   │   ├── freelancers/
    │   │   │   │   │   │   ├── [userId].tsx  # Freelancer profile (mobile)
    │   │   │   │   │   │   ├── filters.tsx  # Freelancer search filters
    │   │   │   │   │   │   └── index.tsx  # Freelancers list (mobile)
    │   │   │   │   │   ├── jobs/
    │   │   │   │   │   │   ├── [jobId].tsx  # Job detail (mobile)
    │   │   │   │   │   │   ├── filters.tsx  # Job search filters
    │   │   │   │   │   │   └── index.tsx  # Jobs list (mobile)
    │   │   │   │   │   ├── _layout.tsx  # Browse tab layout
    │   │   │   │   │   └── categories.tsx  # Browse categories
    │   │   │   │   ├── camera/
    │   │   │   │   │   ├── document-scan/
    │   │   │   │   │   │   └── index.tsx  # (for web parity) - Document scanner
    │   │   │   │   │   ├── photo-upload/
    │   │   │   │   │   │   └── index.tsx  # (for web parity)
    │   │   │   │   │   ├── index.tsx  # Camera/upload hub
    │   │   │   │   │   └── layout.tsx  # Camera section layout
    │   │   │   │   ├── certifications/  # ENTIRE FEATURE
    │   │   │   │   │   ├── [certId]/
    │   │   │   │   │   │   ├── details.tsx  # Certification details (mobile)
    │   │   │   │   │   │   └── verify.tsx  # Verify certification (mobile)
    │   │   │   │   │   ├── add/
    │   │   │   │   │   │   └── index.tsx  # Add certification (mobile)
    │   │   │   │   │   ├── pending-verification/
    │   │   │   │   │   │   └── index.tsx  # Pending verification (mobile)
    │   │   │   │   │   └── index.tsx  # Certifications overview (mobile)
    │   │   │   │   ├── contests/  # ENTIRE FEATURE
    │   │   │   │   │   ├── [contestId]/
    │   │   │   │   │   │   ├── details.tsx  # Contest details (mobile)
    │   │   │   │   │   │   ├── entries.tsx  # Contest entries list (mobile)
    │   │   │   │   │   │   ├── leaderboard.tsx  # Contest leaderboard (mobile)
    │   │   │   │   │   │   └── submit-entry.tsx  # Submit contest entry (mobile)
    │   │   │   │   │   ├── active/
    │   │   │   │   │   │   └── index.tsx  # Active contests list (mobile)
    │   │   │   │   │   ├── browse/
    │   │   │   │   │   │   └── index.tsx  # Browse all contests (mobile)
    │   │   │   │   │   ├── create/
    │   │   │   │   │   │   └── index.tsx  # Create new contest (mobile)
    │   │   │   │   │   └── my-entries/
    │   │   │   │   │       └── index.tsx  # My contest entries (mobile)
    │   │   │   │   ├── contracts/
    │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   ├── disputes/  # - Contract disputes
    │   │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   │   ├── evidence/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Dispute evidence
    │   │   │   │   │   │   │   │   ├── resolution/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Dispute resolution
    │   │   │   │   │   │   │   │   └── index.tsx  # Dispute detail
    │   │   │   │   │   │   │   └── index.tsx  # Disputes list
    │   │   │   │   │   │   ├── chat.tsx  # Contract chat (mobile)
    │   │   │   │   │   │   ├── deliverables.tsx  # Deliverables (mobile)
    │   │   │   │   │   │   ├── details.tsx  # Contract details (mobile)
    │   │   │   │   │   │   ├── disputes.tsx  # Contract disputes (mobile)
    │   │   │   │   │   │   ├── milestones.tsx  # Milestones (mobile)
    │   │   │   │   │   │   ├── workdiary.tsx  # Work diary (mobile)
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── templates/  # - Contract templates
    │   │   │   │   │   │   ├── [templateId]/
    │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   └── index.tsx  # Edit template
    │   │   │   │   │   │   │   ├── preview/
    │   │   │   │   │   │   │   │   └── index.tsx  # Template preview
    │   │   │   │   │   │   │   └── index.tsx  # Template detail
    │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   └── index.tsx  # Create new template
    │   │   │   │   │   │   └── index.tsx  # Templates list
    │   │   │   │   │   ├── _layout.tsx  # Contracts tab layout
    │   │   │   │   │   ├── active.tsx  # Active contracts (mobile)
    │   │   │   │   │   ├── completed.tsx  # Completed contracts (mobile)
    │   │   │   │   │   ├── index.tsx  # All contracts (mobile)
    │   │   │   │   │   └── │
    │   │   │   │   ├── dashboard/
    │   │   │   │   │   ├── _layout.tsx  # Dashboard tab layout
    │   │   │   │   │   ├── analytics.tsx  # User analytics (mobile)
    │   │   │   │   │   ├── earnings.tsx  # Earnings overview (mobile)
    │   │   │   │   │   ├── index.tsx  # Main dashboard (mobile)
    │   │   │   │   │   └── notifications.tsx  # Notifications (mobile)
    │   │   │   │   ├── devices/
    │   │   │   │   │   ├── [deviceId]/
    │   │   │   │   │   │   ├── revoke/
    │   │   │   │   │   │   │   └── index.tsx  # Revoke device session (mobile)
    │   │   │   │   │   │   ├── index.tsx  # Device detail (mobile)
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── index.tsx  # Devices list (mobile)
    │   │   │   │   │   └── │
    │   │   │   │   ├── escrow/  # ENTIRE FEATURE
    │   │   │   │   │   ├── [escrowId]/
    │   │   │   │   │   │   ├── details.tsx  # Escrow details (mobile)
    │   │   │   │   │   │   ├── dispute.tsx  # File escrow dispute (mobile)
    │   │   │   │   │   │   └── release.tsx  # Release escrow funds (mobile)
    │   │   │   │   │   ├── active/
    │   │   │   │   │   │   └── index.tsx  # Active escrow list (mobile)
    │   │   │   │   │   └── history/
    │   │   │   │   │       └── index.tsx  # Escrow history (mobile)
    │   │   │   │   ├── financial/
    │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   ├── [disputeId].tsx  # Dispute detail (mobile)
    │   │   │   │   │   │   └── index.tsx  # Disputes list (mobile)
    │   │   │   │   │   ├── invoices/  # - Invoice management
    │   │   │   │   │   │   ├── [invoiceId]/
    │   │   │   │   │   │   │   ├── download/
    │   │   │   │   │   │   │   │   └── index.tsx  # Download invoice
    │   │   │   │   │   │   │   ├── pay/
    │   │   │   │   │   │   │   │   └── index.tsx  # Pay invoice
    │   │   │   │   │   │   │   └── index.tsx  # Invoice detail
    │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   └── index.tsx  # Create invoice
    │   │   │   │   │   │   ├── [invoiceId].tsx  # Invoice detail (mobile)
    │   │   │   │   │   │   └── index.tsx  # Invoices list (mobile)
    │   │   │   │   │   ├── payouts/
    │   │   │   │   │   │   ├── [payoutId].tsx  # Payout detail (mobile)
    │   │   │   │   │   │   ├── index.tsx  # Payouts list (mobile)
    │   │   │   │   │   │   └── request.tsx  # Request payout (mobile)
    │   │   │   │   │   ├── tax-documents/  # - Tax document management
    │   │   │   │   │   │   ├── [documentId]/
    │   │   │   │   │   │   │   ├── download/
    │   │   │   │   │   │   │   │   └── index.tsx  # Download tax document
    │   │   │   │   │   │   │   └── index.tsx  # Tax document detail
    │   │   │   │   │   │   ├── generate/
    │   │   │   │   │   │   │   └── index.tsx  # Generate tax documents
    │   │   │   │   │   │   └── index.tsx  # Tax documents list
    │   │   │   │   │   ├── _layout.tsx  # Financial tab layout
    │   │   │   │   │   ├── transactions.tsx  # Transaction history (mobile)
    │   │   │   │   │   ├── wallet.tsx  # Wallet (mobile)
    │   │   │   │   │   └── │
    │   │   │   │   ├── groups/  # ENTIRE FEATURE
    │   │   │   │   │   ├── [groupId]/
    │   │   │   │   │   │   ├── events/
    │   │   │   │   │   │   │   └── index.tsx  # Group events (mobile)
    │   │   │   │   │   │   ├── posts/
    │   │   │   │   │   │   │   ├── [postId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Group post detail (mobile)
    │   │   │   │   │   │   │   └── create.tsx  # Create group post (mobile)
    │   │   │   │   │   │   ├── details.tsx  # Group details (mobile)
    │   │   │   │   │   │   └── members.tsx  # Group members (mobile)
    │   │   │   │   │   ├── create/
    │   │   │   │   │   │   └── index.tsx  # Create group (mobile)
    │   │   │   │   │   ├── discover/
    │   │   │   │   │   │   └── index.tsx  # Discover groups (mobile)
    │   │   │   │   │   └── my-groups/
    │   │   │   │   │       └── index.tsx  # My groups (mobile)
    │   │   │   │   ├── jobs/
    │   │   │   │   │   ├── [jobId]/
    │   │   │   │   │   │   ├── apply.tsx  # Apply to job (mobile)
    │   │   │   │   │   │   ├── details.tsx  # Job details (mobile)
    │   │   │   │   │   │   └── proposals.tsx  # Job proposals (client view, mobile)
    │   │   │   │   │   ├── archived/  # ENTIRE SECTION
    │   │   │   │   │   │   ├── [jobId]/
    │   │   │   │   │   │   │   └── index.tsx  # Archived job detail
    │   │   │   │   │   │   └── index.tsx  # Archived jobs list
    │   │   │   │   │   ├── drafts/  # ENTIRE SECTION
    │   │   │   │   │   │   ├── [draftId]/
    │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   └── index.tsx  # Edit draft job
    │   │   │   │   │   │   │   └── index.tsx  # Draft job detail
    │   │   │   │   │   │   └── index.tsx  # Draft jobs list
    │   │   │   │   │   ├── templates/  # - Job templates
    │   │   │   │   │   │   ├── [templateId]/
    │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   └── index.tsx  # Edit job template
    │   │   │   │   │   │   │   ├── use/
    │   │   │   │   │   │   │   │   └── index.tsx  # Use template
    │   │   │   │   │   │   │   └── index.tsx  # Template detail
    │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   └── index.tsx  # Create job template
    │   │   │   │   │   │   └── index.tsx  # Job templates list
    │   │   │   │   │   ├── _layout.tsx  # Jobs tab layout
    │   │   │   │   │   ├── create.tsx  # Create job (mobile)
    │   │   │   │   │   ├── drafts.tsx  # Job drafts (mobile)
    │   │   │   │   │   ├── index.tsx  # Browse jobs (mobile)
    │   │   │   │   │   ├── my-jobs.tsx  # My posted jobs (mobile)
    │   │   │   │   │   ├── saved.tsx  # Saved jobs (mobile)
    │   │   │   │   │   └── │
    │   │   │   │   ├── messages/
    │   │   │   │   │   ├── [conversationId]/
    │   │   │   │   │   │   ├── chat.tsx  # Chat interface (mobile)
    │   │   │   │   │   │   └── details.tsx  # Conversation details (mobile)
    │   │   │   │   │   ├── _layout.tsx  # Messages tab layout
    │   │   │   │   │   ├── archived.tsx  # Archived conversations (mobile)
    │   │   │   │   │   ├── compose.tsx  # New message (mobile)
    │   │   │   │   │   └── index.tsx  # Conversations list (mobile)
    │   │   │   │   ├── milestones/  # ENTIRE FEATURE (as separate tab)
    │   │   │   │   │   ├── [milestoneId]/
    │   │   │   │   │   │   ├── approve.tsx  # Approve milestone (mobile)
    │   │   │   │   │   │   ├── details.tsx  # Milestone details (mobile)
    │   │   │   │   │   │   └── submit.tsx  # Submit milestone deliverable (mobile)
    │   │   │   │   │   ├── completed/
    │   │   │   │   │   │   └── index.tsx  # Completed milestones (mobile)
    │   │   │   │   │   ├── in-review/
    │   │   │   │   │   │   └── index.tsx  # In-review milestones (mobile)
    │   │   │   │   │   └── pending/
    │   │   │   │   │       └── index.tsx  # Pending milestones (mobile)
    │   │   │   │   ├── more/
    │   │   │   │   │   ├── _layout.tsx  # More menu stack
    │   │   │   │   │   ├── about.tsx  # About app
    │   │   │   │   │   ├── account.tsx  # Account settings
    │   │   │   │   │   ├── help.tsx  # Help center
    │   │   │   │   │   └── index.tsx  # More menu home
    │   │   │   │   ├── notifications/
    │   │   │   │   │   ├── [notificationId].tsx  # Notification detail
    │   │   │   │   │   ├── _layout.tsx  # Notifications stack
    │   │   │   │   │   ├── index.tsx  # All notifications
    │   │   │   │   │   └── settings.tsx  # Notification settings
    │   │   │   │   ├── offline/
    │   │   │   │   │   ├── queue/
    │   │   │   │   │   │   └── index.tsx  # Pending actions
    │   │   │   │   │   ├── sync/
    │   │   │   │   │   │   └── index.tsx  # Sync status & controls
    │   │   │   │   │   └── index.tsx  # Offline dashboard
    │   │   │   │   ├── organizations/  # - Organization management
    │   │   │   │   │   ├── [orgId]/
    │   │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   │   ├── invoices/
    │   │   │   │   │   │   │   │   └── index.tsx  # Organization invoices
    │   │   │   │   │   │   │   ├── payment-methods/
    │   │   │   │   │   │   │   │   └── index.tsx  # Org payment methods
    │   │   │   │   │   │   │   └── index.tsx  # Organization billing
    │   │   │   │   │   │   ├── members/
    │   │   │   │   │   │   │   ├── [memberId]/
    │   │   │   │   │   │   │   │   ├── permissions/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Member permissions
    │   │   │   │   │   │   │   │   └── index.tsx  # Member detail
    │   │   │   │   │   │   │   ├── invite/
    │   │   │   │   │   │   │   │   └── index.tsx  # Invite members
    │   │   │   │   │   │   │   └── index.tsx  # Members list
    │   │   │   │   │   │   ├── projects/
    │   │   │   │   │   │   │   ├── [projectId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Project detail
    │   │   │   │   │   │   │   └── index.tsx  # Projects list
    │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   ├── branding/
    │   │   │   │   │   │   │   │   └── index.tsx  # Organization branding
    │   │   │   │   │   │   │   ├── general/
    │   │   │   │   │   │   │   │   └── index.tsx  # General settings
    │   │   │   │   │   │   │   └── index.tsx  # Settings overview
    │   │   │   │   │   │   └── index.tsx  # Organization detail
    │   │   │   │   │   ├── create/
    │   │   │   │   │   │   └── index.tsx  # Create organization
    │   │   │   │   │   └── index.tsx  # Organizations list
    │   │   │   │   ├── profile/
    │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   ├── basic.tsx  # Edit basic info (mobile)
    │   │   │   │   │   │   ├── education.tsx  # Edit education (mobile)
    │   │   │   │   │   │   ├── experience.tsx  # Edit experience (mobile)
    │   │   │   │   │   │   ├── portfolio.tsx  # Edit portfolio (mobile)
    │   │   │   │   │   │   └── skills.tsx  # Edit skills (mobile)
    │   │   │   │   │   ├── portfolio/
    │   │   │   │   │   │   ├── [itemId].tsx  # Portfolio item detail (mobile)
    │   │   │   │   │   │   ├── add.tsx  # Add portfolio item (mobile)
    │   │   │   │   │   │   └── index.tsx  # Portfolio list (mobile)
    │   │   │   │   │   ├── _layout.tsx  # Profile tab layout
    │   │   │   │   │   ├── index.tsx  # View profile (mobile)
    │   │   │   │   │   ├── reviews.tsx  # User reviews (mobile)
    │   │   │   │   │   └── settings.tsx  # Profile settings (mobile)
    │   │   │   │   ├── proposals/
    │   │   │   │   │   ├── [proposalId]/
    │   │   │   │   │   │   ├── details.tsx  # Proposal details (mobile)
    │   │   │   │   │   │   ├── edit.tsx  # Edit proposal (mobile)
    │   │   │   │   │   │   └── withdraw.tsx  # Withdraw proposal (mobile)
    │   │   │   │   │   ├── templates/  # SECTION - Proposal templates
    │   │   │   │   │   │   ├── [templateId]/
    │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   └── index.tsx  # Edit template
    │   │   │   │   │   │   │   ├── use/
    │   │   │   │   │   │   │   │   └── index.tsx  # Use proposal template
    │   │   │   │   │   │   │   └── index.tsx  # Template detail
    │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   └── index.tsx  # Create template
    │   │   │   │   │   │   └── index.tsx  # Templates list
    │   │   │   │   │   ├── _layout.tsx  # Proposals tab layout
    │   │   │   │   │   ├── active.tsx  # Active proposals (mobile)
    │   │   │   │   │   ├── archived.tsx  # Archived proposals (mobile)
    │   │   │   │   │   ├── drafts.tsx  # Proposal drafts (mobile)
    │   │   │   │   │   └── index.tsx  # All proposals (mobile)
    │   │   │   │   ├── search/
    │   │   │   │   │   ├── _layout.tsx  # Search stack navigator
    │   │   │   │   │   ├── filters.tsx  # Advanced filters
    │   │   │   │   │   ├── freelancers.tsx  # Search freelancers (mobile)
    │   │   │   │   │   ├── index.tsx  # Search home
    │   │   │   │   │   ├── jobs.tsx  # Search jobs (mobile)
    │   │   │   │   │   ├── portfolios.tsx  # Search portfolios (mobile)
    │   │   │   │   │   ├── results.tsx  # Search results
    │   │   │   │   │   ├── saved-searches.tsx  # Saved searches (mobile)
    │   │   │   │   │   └── saved.tsx  # Saved searches
    │   │   │   │   ├── settings/
    │   │   │   │   │   ├── account/
    │   │   │   │   │   │   ├── close.tsx  # Close account (mobile)
    │   │   │   │   │   │   ├── email.tsx  # Change email (mobile)
    │   │   │   │   │   │   ├── password.tsx  # Change password (mobile)
    │   │   │   │   │   │   └── phone.tsx  # Change phone (mobile)
    │   │   │   │   │   ├── advanced/
    │   │   │   │   │   │   ├── api-access/
    │   │   │   │   │   │   │   ├── create-token/
    │   │   │   │   │   │   │   │   └── index.tsx  # - Token name input
    │   │   │   │   │   │   │   └── index.tsx  # - Active API keys list
    │   │   │   │   │   │   ├── developer/
    │   │   │   │   │   │   │   ├── api-keys/
    │   │   │   │   │   │   │   │   ├── [keyId]/
    │   │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   │   └── index.tsx  # - Request logs (last 100)
    │   │   │   │   │   │   │   │   │   ├── permissions/
    │   │   │   │   │   │   │   │   │   │   └── index.tsx  # - Granted scopes list
    │   │   │   │   │   │   │   │   │   └── index.tsx  # - Key info (name, created)
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # - Key name input
    │   │   │   │   │   │   │   │   └── index.tsx  # - All API keys list
    │   │   │   │   │   │   │   ├── oauth-apps/
    │   │   │   │   │   │   │   │   ├── [appId]/
    │   │   │   │   │   │   │   │   │   ├── authorizations/
    │   │   │   │   │   │   │   │   │   │   └── index.tsx  # - Users who authorized
    │   │   │   │   │   │   │   │   │   ├── credentials/
    │   │   │   │   │   │   │   │   │   │   └── index.tsx  # - Client ID (copy button)
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── index.tsx  # - Edit app details
    │   │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   │   └── index.tsx  # - Auth events
    │   │   │   │   │   │   │   │   │   ├── scopes/
    │   │   │   │   │   │   │   │   │   │   └── index.tsx  # - Available scopes list
    │   │   │   │   │   │   │   │   │   └── index.tsx  # - App info
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # - App name input
    │   │   │   │   │   │   │   │   └── index.tsx  # - All OAuth apps
    │   │   │   │   │   │   │   ├── sandbox/
    │   │   │   │   │   │   │   │   └── index.tsx  # - Endpoint selector
    │   │   │   │   │   │   │   ├── usage/
    │   │   │   │   │   │   │   │   └── index.tsx  # - Total requests today/week/month
    │   │   │   │   │   │   │   └── index.tsx  # - Developer tools home
    │   │   │   │   │   │   └── webhooks/
    │   │   │   │   │   │       ├── [webhookId]/
    │   │   │   │   │   │       │   ├── deliveries/
    │   │   │   │   │   │       │   │   ├── [deliveryId]/
    │   │   │   │   │   │       │   │   │   └── index.tsx  # - Request/response details
    │   │   │   │   │   │       │   │   └── index.tsx  # - Delivery history
    │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │       │   │   └── index.tsx  # - Edit webhook URL
    │   │   │   │   │   │       │   ├── logs/
    │   │   │   │   │   │       │   │   └── index.tsx  # - Recent deliveries
    │   │   │   │   │   │       │   └── index.tsx  # - Webhook details
    │   │   │   │   │   │       ├── create/
    │   │   │   │   │   │       │   └── index.tsx  # - URL input
    │   │   │   │   │   │       └── index.tsx  # - All webhooks list
    │   │   │   │   │   ├── api/  # - API settings
    │   │   │   │   │   │   ├── documentation/
    │   │   │   │   │   │   │   └── index.tsx  # API documentation
    │   │   │   │   │   │   ├── keys/
    │   │   │   │   │   │   │   ├── [keyId]/
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Edit API key
    │   │   │   │   │   │   │   │   ├── revoke/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Revoke API key
    │   │   │   │   │   │   │   │   └── index.tsx  # API key detail
    │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   └── index.tsx  # Create API key
    │   │   │   │   │   │   │   └── index.tsx  # API keys list
    │   │   │   │   │   │   └── index.tsx  # API settings home
    │   │   │   │   │   ├── authorized-apps/  # ENTIRE SECTION
    │   │   │   │   │   │   ├── [appId]/
    │   │   │   │   │   │   │   ├── permissions/
    │   │   │   │   │   │   │   │   └── index.tsx  # App permissions detail (mobile)
    │   │   │   │   │   │   │   ├── index.tsx  # Authorized app detail (mobile)
    │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   ├── index.tsx  # Authorized apps list (mobile)
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── automation/  # - Automation settings
    │   │   │   │   │   │   ├── rules/  # SECTION
    │   │   │   │   │   │   │   ├── [ruleId]/
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Edit automation rule
    │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Rule execution history
    │   │   │   │   │   │   │   │   ├── index.tsx  # Rule detail
    │   │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   └── index.tsx  # Create automation rule
    │   │   │   │   │   │   │   ├── index.tsx  # Automation rules list
    │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   ├── templates/
    │   │   │   │   │   │   │   ├── [templateId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Automation template detail
    │   │   │   │   │   │   │   └── index.tsx  # Automation templates
    │   │   │   │   │   │   ├── _layout.tsx  # Automation layout
    │   │   │   │   │   │   └── index.tsx  # Automation settings home
    │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   ├── payment-methods.tsx  # Payment methods (mobile)
    │   │   │   │   │   │   └── subscription.tsx  # Subscription (mobile)
    │   │   │   │   │   ├── developer/  # - Developer settings
    │   │   │   │   │   │   ├── api-keys/  # ENTIRE SUBSECTION
    │   │   │   │   │   │   │   ├── [keyId]/
    │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # API key usage logs (mobile)
    │   │   │   │   │   │   │   │   ├── permissions/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # API key permissions (mobile)
    │   │   │   │   │   │   │   │   ├── regenerate/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Regenerate key
    │   │   │   │   │   │   │   │   ├── index.tsx  # API key detail (mobile)
    │   │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   └── index.tsx  # Create API key (mobile)
    │   │   │   │   │   │   │   ├── index.tsx  # API keys list (mobile)
    │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   └── index.tsx  # Create new API key
    │   │   │   │   │   │   ├── documentation/
    │   │   │   │   │   │   │   └── index.tsx  # Embedded API docs
    │   │   │   │   │   │   ├── oauth-apps/  # ENTIRE SUBSECTION
    │   │   │   │   │   │   │   ├── [appId]/
    │   │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # OAuth app analytics
    │   │   │   │   │   │   │   │   ├── authorizations/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # OAuth authorizations (mobile)
    │   │   │   │   │   │   │   │   ├── credentials/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # OAuth app credentials
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Edit OAuth app
    │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # OAuth logs
    │   │   │   │   │   │   │   │   ├── scopes/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # OAuth app scopes (mobile)
    │   │   │   │   │   │   │   │   ├── index.tsx  # OAuth app detail
    │   │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   └── index.tsx  # Create OAuth app
    │   │   │   │   │   │   │   ├── index.tsx  # OAuth apps list
    │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   ├── sandbox/
    │   │   │   │   │   │   │   └── index.tsx  # API sandbox/testing (mobile)
    │   │   │   │   │   │   ├── usage/
    │   │   │   │   │   │   │   └── index.tsx  # API usage dashboard (mobile)
    │   │   │   │   │   │   ├── webhooks/
    │   │   │   │   │   │   │   ├── [webhookId]/
    │   │   │   │   │   │   │   │   ├── deliveries/
    │   │   │   │   │   │   │   │   │   ├── [deliveryId]/
    │   │   │   │   │   │   │   │   │   │   └── index.tsx  # Webhook delivery detail (mobile)
    │   │   │   │   │   │   │   │   │   ├── index.tsx  # Webhook deliveries (mobile)
    │   │   │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Edit webhook
    │   │   │   │   │   │   │   │   ├── events/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Webhook events config (mobile)
    │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Webhook logs
    │   │   │   │   │   │   │   │   ├── test/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Test webhook
    │   │   │   │   │   │   │   │   ├── index.tsx  # Webhook detail
    │   │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   └── index.tsx  # Create webhook
    │   │   │   │   │   │   │   ├── index.tsx  # Webhooks list
    │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   ├── _layout.tsx  # Developer settings layout
    │   │   │   │   │   │   ├── index.tsx  # Developer settings home
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── devices/  # - Connected devices
    │   │   │   │   │   │   ├── [deviceId]/
    │   │   │   │   │   │   │   ├── revoke/
    │   │   │   │   │   │   │   │   └── index.tsx  # Revoke device access
    │   │   │   │   │   │   │   └── index.tsx  # Device detail
    │   │   │   │   │   │   └── index.tsx  # Connected devices list
    │   │   │   │   │   ├── integrations/  # ENTIRE SECTION - Third-party integrations
    │   │   │   │   │   │   ├── [integration]/
    │   │   │   │   │   │   │   └── index.tsx  # Integration detail
    │   │   │   │   │   │   ├── available/
    │   │   │   │   │   │   │   └── index.tsx  # Integrations marketplace
    │   │   │   │   │   │   ├── calendar/  # - Calendar integration
    │   │   │   │   │   │   │   ├── google/
    │   │   │   │   │   │   │   │   ├── authorize/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Authorize Google Calendar
    │   │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Google Calendar settings
    │   │   │   │   │   │   │   │   └── index.tsx  # Google Calendar integration
    │   │   │   │   │   │   │   ├── outlook/
    │   │   │   │   │   │   │   │   ├── authorize/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Authorize Outlook Calendar
    │   │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Outlook Calendar settings
    │   │   │   │   │   │   │   │   └── index.tsx  # Outlook Calendar integration
    │   │   │   │   │   │   │   └── index.tsx  # Calendar integrations
    │   │   │   │   │   │   ├── payment/  # - Payment provider integration
    │   │   │   │   │   │   │   ├── paypal/
    │   │   │   │   │   │   │   │   ├── authorize/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Authorize PayPal
    │   │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # PayPal settings
    │   │   │   │   │   │   │   │   └── index.tsx  # PayPal integration
    │   │   │   │   │   │   │   ├── stripe/
    │   │   │   │   │   │   │   │   ├── authorize/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Authorize Stripe
    │   │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   │   └── index.tsx  # Stripe settings
    │   │   │   │   │   │   │   │   └── index.tsx  # Stripe integration
    │   │   │   │   │   │   │   └── index.tsx  # Payment provider integrations
    │   │   │   │   │   │   ├── slack/
    │   │   │   │   │   │   │   └── index.tsx  # Slack integration
    │   │   │   │   │   │   ├── index.tsx  # Integrations overview
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── labs/  # Experimental features
    │   │   │   │   │   │   └── index.tsx  # Experimental features
    │   │   │   │   │   ├── login-history/  # Login audit trail
    │   │   │   │   │   │   └── index.tsx  # Login history
    │   │   │   │   │   ├── organization/  # - Organization settings (mobile)
    │   │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   │   └── index.tsx  # Organization billing (mobile)
    │   │   │   │   │   │   ├── members/
    │   │   │   │   │   │   │   └── index.tsx  # Organization members (mobile)
    │   │   │   │   │   │   └── index.tsx  # Organization settings (mobile)
    │   │   │   │   │   ├── preferences/
    │   │   │   │   │   │   └── accessibility/
    │   │   │   │   │   │       └── index.tsx  # Accessibility settings
    │   │   │   │   │   ├── privacy-security/
    │   │   │   │   │   │   └── authorized-apps/
    │   │   │   │   │   │       ├── [appId]/
    │   │   │   │   │   │       │   ├── permissions/
    │   │   │   │   │   │       │   │   └── index.tsx  # App permissions detail
    │   │   │   │   │   │       │   ├── index.tsx  # Authorized app detail
    │   │   │   │   │   │       │   └── │
    │   │   │   │   │   │       ├── index.tsx  # Authorized apps list
    │   │   │   │   │   │       └── │
    │   │   │   │   │   ├── security/
    │   │   │   │   │   │   ├── sessions/  # - Active sessions
    │   │   │   │   │   │   │   ├── [sessionId]/
    │   │   │   │   │   │   │   │   └── index.tsx  # Session detail
    │   │   │   │   │   │   │   └── index.tsx  # Active sessions
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── subscription/  # ENTIRE SECTION
    │   │   │   │   │   │   ├── change-plan/
    │   │   │   │   │   │   │   └── index.tsx  # Change subscription plan
    │   │   │   │   │   │   ├── manage-plan/  # full plan management
    │   │   │   │   │   │   │   ├── cancel/
    │   │   │   │   │   │   │   │   └── index.tsx  # Cancel subscription
    │   │   │   │   │   │   │   ├── change/
    │   │   │   │   │   │   │   │   └── index.tsx  # Change subscription plan
    │   │   │   │   │   │   │   ├── pause/
    │   │   │   │   │   │   │   │   └── index.tsx  # Pause subscription
    │   │   │   │   │   │   │   ├── resume/
    │   │   │   │   │   │   │   │   └── index.tsx  # Resume subscription
    │   │   │   │   │   │   │   └── index.tsx  # Manage subscription plan
    │   │   │   │   │   │   └── index.tsx  # Subscription management
    │   │   │   │   │   ├── _layout.tsx  # Settings tab layout
    │   │   │   │   │   ├── index.tsx  # Settings overview (mobile)
    │   │   │   │   │   ├── notifications.tsx  # Notification settings (mobile)
    │   │   │   │   │   ├── privacy.tsx  # Privacy settings (mobile)
    │   │   │   │   │   ├── security.tsx  # Security settings (mobile)
    │   │   │   │   │   └── │
    │   │   │   │   ├── skills-tests/  # ENTIRE FEATURE
    │   │   │   │   │   ├── [testId]/
    │   │   │   │   │   │   ├── results.tsx  # Test results (mobile)
    │   │   │   │   │   │   ├── start.tsx  # Start skills test (mobile)
    │   │   │   │   │   │   └── test.tsx  # Take skills test (mobile)
    │   │   │   │   │   ├── available/
    │   │   │   │   │   │   └── index.tsx  # Available tests (mobile)
    │   │   │   │   │   ├── completed/
    │   │   │   │   │   │   └── index.tsx  # Completed tests (mobile)
    │   │   │   │   │   └── index.tsx  # Skills tests overview (mobile)
    │   │   │   │   ├── talent-cloud/  # ENTIRE FEATURE
    │   │   │   │   │   ├── agencies/
    │   │   │   │   │   │   ├── [agencyId]/
    │   │   │   │   │   │   │   └── details.tsx  # Agency details (mobile)
    │   │   │   │   │   │   └── index.tsx  # Agencies list (mobile)
    │   │   │   │   │   ├── projects/
    │   │   │   │   │   │   ├── [projectId]/
    │   │   │   │   │   │   │   └── details.tsx  # Project details (mobile)
    │   │   │   │   │   │   └── index.tsx  # Projects list (mobile)
    │   │   │   │   │   └── teams/
    │   │   │   │   │       ├── [teamId]/
    │   │   │   │   │       │   ├── details.tsx  # Team details (mobile)
    │   │   │   │   │       │   └── members.tsx  # Team members (mobile)
    │   │   │   │   │       └── index.tsx  # Teams list (mobile)
    │   │   │   │   ├── timesheet/  # ENTIRE FEATURE (as separate tab)
    │   │   │   │   │   ├── [timesheetId]/
    │   │   │   │   │   │   ├── details.tsx  # Timesheet details (mobile)
    │   │   │   │   │   │   └── edit.tsx  # Edit timesheet (mobile)
    │   │   │   │   │   ├── manual-time/
    │   │   │   │   │   │   └── index.tsx  # Manual time entry (mobile)
    │   │   │   │   │   ├── weekly/
    │   │   │   │   │   │   └── index.tsx  # Weekly timesheet (mobile)
    │   │   │   │   │   ├── work-diary/
    │   │   │   │   │   │   ├── [date]/
    │   │   │   │   │   │   │   └── index.tsx  # Work diary for date (mobile)
    │   │   │   │   │   │   └── index.tsx  # Work diary overview (mobile)
    │   │   │   │   │   └── index.tsx  # Timesheet overview (mobile)
    │   │   │   │   ├── _layout.tsx  # Tabs layout
    │   │   │   │   ├── index.tsx  # Home tab / Dashboard
    │   │   │   │   ├── jobs.tsx  # Jobs tab
    │   │   │   │   ├── messages.tsx  # Messages tab
    │   │   │   │   ├── profile.tsx  # Profile tab
    │   │   │   │   ├── proposals.tsx  # Proposals tab
    │   │   │   │   └── │
    │   │   │   ├── (work)/
    │   │   │   │   └── contracts/
    │   │   │   │       └── [id]/
    │   │   │   │           ├── milestones/
    │   │   │   │           │   └── index.tsx  # Milestones (mobile)
    │   │   │   │           └── work-diary/
    │   │   │   │               └── index.tsx  # Work diary & timesheets (mobile)
    │   │   │   ├── accessibility/
    │   │   │   │   ├── screen-reader/
    │   │   │   │   │   └── index.tsx  # Screen reader optimized view
    │   │   │   │   └── voice-commands/
    │   │   │   │       └── index.tsx  # Voice commands
    │   │   │   ├── billing/
    │   │   │   │   └── subscriptions/
    │   │   │   │       └── index.tsx  # Subscriptions (mobile)
    │   │   │   ├── camera/
    │   │   │   │   ├── document-scan/
    │   │   │   │   │   └── index.tsx  # Document scanner
    │   │   │   │   ├── photo-upload/
    │   │   │   │   │   └── index.tsx  # Photo upload
    │   │   │   │   └── qr-scan/
    │   │   │   │       └── index.tsx  # QR code scanner
    │   │   │   ├── contracts/
    │   │   │   │   ├── [id]/
    │   │   │   │   │   ├── amendments/
    │   │   │   │   │   │   ├── [amendmentId]/
    │   │   │   │   │   │   │   └── index.tsx  # Amendment detail (mobile)
    │   │   │   │   │   │   ├── index.tsx  # Amendments list (mobile)
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── deliverables/
    │   │   │   │   │   │   ├── [deliverableId]/
    │   │   │   │   │   │   │   └── index.tsx  # Deliverable detail (mobile)
    │   │   │   │   │   │   ├── index.tsx  # Deliverables list (mobile)
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   └── index.tsx  # Dispute detail (mobile)
    │   │   │   │   │   │   ├── index.tsx  # Disputes list (mobile)
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── invoices/
    │   │   │   │   │   │   ├── [invoiceId]/
    │   │   │   │   │   │   │   └── index.tsx  # Invoice detail (mobile)
    │   │   │   │   │   │   ├── index.tsx  # Invoices list (mobile)
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── milestones/
    │   │   │   │   │   │   ├── [milestoneId]/
    │   │   │   │   │   │   │   └── index.tsx  # Milestone detail (mobile)
    │   │   │   │   │   │   ├── index.tsx  # Milestones list (mobile)
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── time-tracking/
    │   │   │   │   │   │   ├── [entryId]/
    │   │   │   │   │   │   │   └── index.tsx  # Time entry detail (mobile)
    │   │   │   │   │   │   ├── index.tsx  # Time tracking list (mobile)
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── index.tsx  # Contract overview
    │   │   │   │   │   ├── messages.tsx  # Contract messages
    │   │   │   │   │   ├── milestones.tsx  # Milestones
    │   │   │   │   │   ├── timesheet.tsx  # Timesheet
    │   │   │   │   │   ├── work-diary.tsx  # Work diary
    │   │   │   │   │   └── │
    │   │   │   │   └── index.tsx  # Contracts list
    │   │   │   ├── financials/
    │   │   │   │   ├── invoices.tsx  # Invoices
    │   │   │   │   ├── transactions.tsx  # Transaction history
    │   │   │   │   └── wallet.tsx  # Wallet
    │   │   │   ├── jobs/
    │   │   │   │   ├── [id].tsx  # Job detail
    │   │   │   │   ├── post.tsx  # Post job (client)
    │   │   │   │   └── search.tsx  # Job search
    │   │   │   ├── messages/
    │   │   │   │   └── [conversationId].tsx  # Conversation thread
    │   │   │   ├── mobile-settings/
    │   │   │   │   ├── app-settings.tsx  # App-specific settings
    │   │   │   │   ├── biometric.tsx  # Biometric authentication
    │   │   │   │   └── haptics.tsx  # Haptic feedback settings
    │   │   │   ├── notifications/
    │   │   │   │   └── index.tsx  # Notifications list
    │   │   │   ├── offline/
    │   │   │   │   ├── queue/
    │   │   │   │   │   └── index.tsx  # Offline queue
    │   │   │   │   ├── settings/
    │   │   │   │   │   └── index.tsx  # Offline settings
    │   │   │   │   ├── sync/
    │   │   │   │   │   └── index.tsx  # Sync status
    │   │   │   │   ├── queue.tsx  # Offline actions queue
    │   │   │   │   ├── sync.tsx  # Sync status
    │   │   │   │   └── │
    │   │   │   ├── onboarding/
    │   │   │   │   ├── biometric/
    │   │   │   │   │   └── setup.tsx  # Biometric auth setup
    │   │   │   │   ├── intro/
    │   │   │   │   │   └── index.tsx  # App intro screens
    │   │   │   │   ├── permissions/
    │   │   │   │   │   ├── camera.tsx  # Camera permission
    │   │   │   │   │   ├── index.tsx  # All permissions
    │   │   │   │   │   ├── location.tsx  # Location permission
    │   │   │   │   │   └── notifications.tsx  # Notification permission
    │   │   │   │   ├── profile-setup/
    │   │   │   │   │   ├── basic-info.tsx  # Basic information
    │   │   │   │   │   ├── photo.tsx  # Profile photo
    │   │   │   │   │   ├── preferences.tsx  # Initial preferences
    │   │   │   │   │   └── skills.tsx  # Skills selection
    │   │   │   │   ├── tutorial/
    │   │   │   │   │   ├── apply.tsx  # Tutorial: Apply
    │   │   │   │   │   ├── browse-jobs.tsx  # Tutorial: Browse jobs
    │   │   │   │   │   ├── index.tsx  # Interactive tutorial
    │   │   │   │   │   └── time-tracking.tsx  # Tutorial: Time tracking
    │   │   │   │   ├── _layout.tsx  # Onboarding layout
    │   │   │   │   ├── complete.tsx  # Onboarding complete
    │   │   │   │   ├── features.tsx  # Features showcase
    │   │   │   │   ├── notifications-setup.tsx  # Notification preferences
    │   │   │   │   ├── permissions.tsx  # Request permissions
    │   │   │   │   ├── profile-setup.tsx  # Quick profile setup
    │   │   │   │   ├── user-type.tsx  # Select user type
    │   │   │   │   ├── verification.tsx  # Identity verification
    │   │   │   │   └── welcome.tsx  # Welcome screen
    │   │   │   ├── profile/
    │   │   │   │   ├── edit/
    │   │   │   │   │   ├── experience.tsx  # Edit experience
    │   │   │   │   │   ├── index.tsx  # Edit profile
    │   │   │   │   │   ├── portfolio.tsx  # Edit portfolio
    │   │   │   │   │   └── skills.tsx  # Edit skills
    │   │   │   │   └── [userId].tsx  # User profile (public view)
    │   │   │   ├── proposals/
    │   │   │   │   ├── submit/
    │   │   │   │   │   └── [jobId].tsx  # Submit proposal
    │   │   │   │   └── [id].tsx  # Proposal detail
    │   │   │   ├── quick-actions/
    │   │   │   │   ├── quick-apply/
    │   │   │   │   │   └── [jobId].tsx  # Quick apply to job
    │   │   │   │   ├── quick-message/
    │   │   │   │   │   └── [userId].tsx  # Quick message to user
    │   │   │   │   ├── quick-time-entry/
    │   │   │   │   │   └── [contractId].tsx  # Quick time logging
    │   │   │   │   ├── quick-apply.tsx  # Quick proposal submission
    │   │   │   │   ├── quick-invoice.tsx  # Quick invoice creation
    │   │   │   │   ├── quick-message.tsx  # Quick message
    │   │   │   │   └── quick-time-entry.tsx  # Quick time logging
    │   │   │   ├── reviews/
    │   │   │   │   ├── create/
    │   │   │   │   │   └── [contractId].tsx  # Create review
    │   │   │   │   └── index.tsx  # Reviews list
    │   │   │   ├── scanner/
    │   │   │   │   ├── document.tsx  # Document scanner
    │   │   │   │   ├── qr-code.tsx  # QR code scanner
    │   │   │   │   └── │
    │   │   │   ├── search/
    │   │   │   │   ├── alerts/
    │   │   │   │   │   └── index.tsx  # Saved search alerts (mobile)
    │   │   │   │   └── saved/
    │   │   │   │       └── index.tsx  # Saved searches (mobile)
    │   │   │   ├── settings/
    │   │   │   │   ├── advanced/
    │   │   │   │   │   ├── developer/
    │   │   │   │   │   │   └── index.tsx  # Developer options
    │   │   │   │   │   └── experiments/
    │   │   │   │   │       └── index.tsx  # Experimental features
    │   │   │   │   ├── app-preferences/
    │   │   │   │   │   ├── data-usage/
    │   │   │   │   │   │   └── index.tsx  # Data usage settings
    │   │   │   │   │   ├── language/
    │   │   │   │   │   │   └── index.tsx  # Language preferences
    │   │   │   │   │   ├── notifications/
    │   │   │   │   │   │   ├── channels/
    │   │   │   │   │   │   │   └── index.tsx  # Notification channels
    │   │   │   │   │   │   ├── do-not-disturb/
    │   │   │   │   │   │   │   └── index.tsx  # DND settings
    │   │   │   │   │   │   └── index.tsx  # Notification preferences
    │   │   │   │   │   └── theme/
    │   │   │   │   │       └── index.tsx  # Theme settings
    │   │   │   │   ├── app-settings/
    │   │   │   │   │   ├── appearance.tsx  # Appearance settings
    │   │   │   │   │   ├── biometric-auth.tsx  # Biometric authentication
    │   │   │   │   │   ├── haptics.tsx  # Haptic feedback
    │   │   │   │   │   ├── offline-mode.tsx  # Offline settings
    │   │   │   │   │   └── quick-actions.tsx  # Quick action config
    │   │   │   │   ├── biometric/
    │   │   │   │   │   └── index.tsx  # Biometric authentication
    │   │   │   │   ├── mobile-specific/
    │   │   │   │   │   ├── battery-optimization.tsx  # Battery settings
    │   │   │   │   │   ├── cache-management.tsx  # Cache management
    │   │   │   │   │   └── data-usage.tsx  # Data usage settings
    │   │   │   │   ├── notifications/
    │   │   │   │   │   ├── channels.tsx  # Notification channels (mobile)
    │   │   │   │   │   └── digests-quiet-hours.tsx  # Digests & quiet hours (mobile)
    │   │   │   │   ├── privacy/
    │   │   │   │   │   └── gdpr/
    │   │   │   │   │       └── index.tsx  # GDPR (mobile)
    │   │   │   │   ├── security/
    │   │   │   │   │   ├── data-export-delete.tsx  # Export / delete account (mobile)
    │   │   │   │   │   └── devices-sessions.tsx  # Active devices & sessions (mobile)
    │   │   │   │   ├── storage/
    │   │   │   │   │   ├── cache/
    │   │   │   │   │   │   └── index.tsx  # Cache management
    │   │   │   │   │   └── downloads/
    │   │   │   │   │       └── index.tsx  # Downloaded files
    │   │   │   │   ├── about.tsx  # About & support
    │   │   │   │   ├── account.tsx  # Account settings
    │   │   │   │   ├── index.tsx  # Settings menu
    │   │   │   │   ├── notifications.tsx  # Notification settings
    │   │   │   │   ├── privacy.tsx  # Privacy settings
    │   │   │   │   └── security.tsx  # Security settings
    │   │   │   ├── status/
    │   │   │   │   └── index.tsx  # Status & incidents (mobile)
    │   │   │   ├── subscription/
    │   │   │   │   ├── connects.tsx  # Connects management
    │   │   │   │   ├── index.tsx  # Subscription overview
    │   │   │   │   ├── plans.tsx  # Available plans
    │   │   │   │   └── upgrade.tsx  # Upgrade plan
    │   │   │   ├── support/
    │   │   │   │   └── tickets/
    │   │   │   │       └── index.tsx  # Support tickets (mobile)
    │   │   │   ├── widgets/
    │   │   │   │   ├── quick-actions.tsx  # Quick actions widget
    │   │   │   │   └── time-tracker.tsx  # Home screen time tracker widget
    │   │   │   ├── +not-found.tsx  # 404 screen
    │   │   │   ├── _layout.tsx  # Root layout
    │   │   │   ├── index.tsx  # App entry point
    │   │   │   ├── |   └── _layout.tsx  # ErrorBoundary
    │   │   │   └── │
    │   │   ├── assets/  # Mobile assets
    │   │   │   ├── fonts/  # Custom fonts
    │   │   │   ├── icons/  # App icons
    │   │   │   ├── images/  # Images
    │   │   │   └── splash/  # Splash screens
    │   │   ├── offline/
    │   │   │   ├── queue.tsx  # Offline actions queue
    │   │   │   └── sync.tsx  # Sync status
    │   │   ├── scanner/
    │   │   │   ├── document.tsx  # Document scanner
    │   │   │   └── qr-code.tsx  # QR code scanner
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
    │   │   │   │   ├── molecules/
    │   │   │   │   │   └── Navigation/
    │   │   │   │   │       ├── Accordion/
    │   │   │   │   │       │   └── Accordion.native.tsx  # Mobile accordion component
    │   │   │   │   │       ├── Breadcrumb/
    │   │   │   │   │       │   └── Breadcrumb.native.tsx  # Mobile breadcrumb navigation component
    │   │   │   │   │       ├── Drawer/
    │   │   │   │   │       │   └── Drawer.native.tsx  # Mobile drawer component (BottomSheet pattern)
    │   │   │   │   │       └── │
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
    │   │   │   │   ├── QuickActions/
    │   │   │   │   │   └── README.md  # Home screen quick actions (notes)
    │   │   │   │   └── UI/
    │   │   │   │       ├── Avatar.tsx  # Avatar component
    │   │   │   │       ├── Badge.tsx  # Badge component
    │   │   │   │       ├── BottomSheet.tsx  # Bottom sheet modal
    │   │   │   │       ├── Button.tsx  # Button component
    │   │   │   │       ├── Card.tsx  # Card component
    │   │   │   │       ├── Input.tsx  # Input component
    │   │   │   │       └── SearchBar.tsx  # Search bar
    │   │   │   ├── hooks/  # Mobile-specific hooks
    │   │   │   │   ├── use-biometric.ts  # (mobile-specific)
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
    │   │   │   │   ├── analytics.ts  # Mobile analytics
    │   │   │   │   ├── deeplink.ts  # Deep link handling
    │   │   │   │   ├── error-tracking.ts  # Error tracking (Sentry)
    │   │   │   │   ├── keycloak-mobile.ts  # Keycloak mobile config
    │   │   │   │   ├── performance.ts  # Performance utilities
    │   │   │   │   ├── push-notifications.ts  # Push notification setup
    │   │   │   │   ├── storage.ts  # Secure storage wrapper
    │   │   │   │   └── utils.ts  # General utilities
    │   │   │   ├── screens/
    │   │   │   │   └── developer/
    │   │   │   │       └── webhooks/
    │   │   │   │           ├── [webhook_id]/
    │   │   │   │           │   ├── analytics/
    │   │   │   │           │   │   └── index.tsx  # Webhook analytics dashboard (mobile)
    │   │   │   │           │   ├── health/
    │   │   │   │           │   │   └── index.tsx  # Webhook health monitoring (mobile)
    │   │   │   │           │   ├── retry-config/
    │   │   │   │           │   │   └── index.tsx  # Webhook retry configuration (mobile)
    │   │   │   │           │   ├── security/
    │   │   │   │           │   │   └── index.tsx  # Webhook security settings (mobile)
    │   │   │   │           │   └── transform/
    │   │   │   │           │       └── index.tsx  # Webhook payload transformation (mobile)
    │   │   │   │           ├── batch-retry/
    │   │   │   │           │   └── index.tsx  # Webhook batch retry (mobile)
    │   │   │   │           ├── compare/
    │   │   │   │           │   └── index.tsx  # Webhook comparison view (mobile)
    │   │   │   │           ├── docs/
    │   │   │   │           │   └── index.tsx  # Webhook documentation (mobile)
    │   │   │   │           ├── events-catalog/
    │   │   │   │           │   └── index.tsx  # Webhook events catalog (mobile)
    │   │   │   │           ├── import-export/
    │   │   │   │           │   └── index.tsx  # Webhook import/export (mobile)
    │   │   │   │           ├── settings/
    │   │   │   │           │   └── index.tsx  # Webhook global settings (mobile)
    │   │   │   │           └── templates/
    │   │   │   │               └── index.tsx  # Webhook templates (mobile)
    │   │   │   ├── stores/  # Mobile-specific Zustand stores
    │   │   │   │   ├── biometric-store.ts  # Biometric settings
    │   │   │   │   ├── camera-store.ts  # Camera state
    │   │   │   │   └── offline-queue-store.ts  # Offline action queue
    │   │   │   └── types/  # Mobile-specific types
    │   │   │       ├── biometric.ts  # Biometric types
    │   │   │       └── navigation.ts  # Navigation types
    │   │   ├── ui/
    │   │   │   └── src/
    │   │   │       └── components/
    │   │   │           ├── molecules/
    │   │   │           │   ├── Navigation/
    │   │   │           │   │   ├── Accordion/
    │   │   │           │   │   │   └── Accordion.native.tsx  # - React Native collapsible accordion
    │   │   │           │   │   ├── Breadcrumb/
    │   │   │           │   │   │   └── Breadcrumb.native.tsx  # - Mobile breadcrumb navigation
    │   │   │           │   │   └── Drawer/
    │   │   │           │   │       └── Drawer.native.tsx  # - React Native bottom sheet drawer
    │   │   │           │   └── Overlay/
    │   │   │           │       ├── Popover/
    │   │   │           │       │   └── Popover.native.tsx  # - Mobile popover component
    │   │   │           │       └── Tooltip/
    │   │   │           │           └── Tooltip.native.tsx  # - Long-press tooltip
    │   │   │           └── organisms/
    │   │   │               └── DataDisplay/
    │   │   │                   ├── DataGrid/
    │   │   │                   │   └── DataGrid.native.tsx  # - Mobile data grid
    │   │   │                   ├── KanbanBoard/
    │   │   │                   │   └── KanbanBoard.native.tsx  # - Mobile kanban board
    │   │   │                   └── Table/
    │   │   │                       └── Table.native.tsx  # - Mobile responsive table
    │   │   ├── widgets/
    │   │   │   ├── quick-actions.tsx  # Quick actions widget
    │   │   │   └── time-tracker.tsx  # Home screen time tracker widget
    │   │   ├── +not-found.tsx  # 404 screen
    │   │   ├── .env  # Environment variables
    │   │   ├── .eslintrc.json  # ESLint config
    │   │   ├── _layout.tsx  # Root layout
    │   │   ├── app.json  # Expo config
    │   │   ├── babel.config.js  # Babel config
    │   │   ├── eas.json  # EAS Build config
    │   │   ├── global.css  # Global styles (NativeWind)
    │   │   ├── index.js  # Entry point
    │   │   ├── index.tsx  # App entry point
    │   │   ├── metro.config.js  # Metro bundler config
    │   │   ├── package.json  # Mobile dependencies
    │   │   ├── tailwind.config.js  # Tailwind config (NativeWind)
    │   │   ├── tsconfig.json  # TypeScript config
    │   │   └── │
    │   ├── web/  # Next.js web application
    │   │   ├── app/
    │   │   │   └── (authenticated)/
    │   │   │       └── developer/
    │   │   │           └── webhooks/
    │   │   │               ├── [webhook_id]/
    │   │   │               │   ├── analytics/
    │   │   │               │   │   └── page.tsx  # Webhook analytics dashboard
    │   │   │               │   ├── health/
    │   │   │               │   │   └── page.tsx  # Webhook health monitoring
    │   │   │               │   ├── retry-config/
    │   │   │               │   │   └── page.tsx  # Webhook retry configuration
    │   │   │               │   ├── security/
    │   │   │               │   │   └── page.tsx  # Webhook security settings
    │   │   │               │   └── transform/
    │   │   │               │       └── page.tsx  # Webhook payload transformation
    │   │   │               ├── batch-retry/
    │   │   │               │   └── page.tsx  # Webhook batch retry
    │   │   │               ├── compare/
    │   │   │               │   └── page.tsx  # Webhook comparison view
    │   │   │               ├── docs/
    │   │   │               │   └── page.tsx  # Webhook documentation
    │   │   │               ├── events-catalog/
    │   │   │               │   └── page.tsx  # Webhook events catalog
    │   │   │               ├── import-export/
    │   │   │               │   └── page.tsx  # Webhook import/export
    │   │   │               ├── settings/
    │   │   │               │   └── page.tsx  # Webhook global settings
    │   │   │               └── templates/
    │   │   │                   └── page.tsx  # Webhook templates
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
    │   │   │   ├── app/  # Next.js App Router
    │   │   │   │   ├── (admin)/
    │   │   │   │   │   ├── admin/  # WEB ADMIN ROUTES - ALL
    │   │   │   │   │   │   ├── analytics/  # Admin analytics dashboards
    │   │   │   │   │   │   │   ├── contracts/
    │   │   │   │   │   │   │   │   └── page.tsx  # Contracts analytics
    │   │   │   │   │   │   │   ├── financial/
    │   │   │   │   │   │   │   │   └── page.tsx  # Financial analytics
    │   │   │   │   │   │   │   ├── jobs/
    │   │   │   │   │   │   │   │   └── page.tsx  # Jobs analytics
    │   │   │   │   │   │   │   ├── proposals/
    │   │   │   │   │   │   │   │   └── page.tsx  # Proposals analytics
    │   │   │   │   │   │   │   ├── search/
    │   │   │   │   │   │   │   │   └── page.tsx  # Search analytics
    │   │   │   │   │   │   │   ├── users/
    │   │   │   │   │   │   │   │   └── page.tsx  # Users analytics
    │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   ├── audit/  # Audit logs and compliance
    │   │   │   │   │   │   │   ├── events/
    │   │   │   │   │   │   │   │   ├── [eventId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Audit event detail
    │   │   │   │   │   │   │   │   ├── page.tsx  # Audit events list
    │   │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   │   ├── exports/
    │   │   │   │   │   │   │   │   ├── [exportId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Audit export detail
    │   │   │   │   │   │   │   │   ├── page.tsx  # Audit exports list
    │   │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   └── page.tsx  # Compliance reports
    │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   └── page.tsx  # Audit settings
    │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   ├── content/
    │   │   │   │   │   │   │   ├── blog/
    │   │   │   │   │   │   │   │   ├── [postId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit blog post
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Blog post details
    │   │   │   │   │   │   │   │   ├── categories/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Blog categories
    │   │   │   │   │   │   │   │   ├── new/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create blog post
    │   │   │   │   │   │   │   │   └── page.tsx  # Blog posts list
    │   │   │   │   │   │   │   ├── help-center/
    │   │   │   │   │   │   │   │   ├── [articleId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit help article
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Help article details
    │   │   │   │   │   │   │   │   ├── categories/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Help categories
    │   │   │   │   │   │   │   │   ├── new/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create help article
    │   │   │   │   │   │   │   │   └── page.tsx  # Help articles list
    │   │   │   │   │   │   │   └── testimonials/
    │   │   │   │   │   │   │       ├── [testimonialId]/
    │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit testimonial
    │   │   │   │   │   │   │       │   └── page.tsx  # Testimonial details
    │   │   │   │   │   │   │       ├── new/
    │   │   │   │   │   │   │       │   └── page.tsx  # Create testimonial
    │   │   │   │   │   │   │       └── page.tsx  # Testimonials list
    │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   │   ├── evidence/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute evidence
    │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute history
    │   │   │   │   │   │   │   │   ├── messages/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute messages
    │   │   │   │   │   │   │   │   └── page.tsx  # Dispute details
    │   │   │   │   │   │   │   ├── escalated/
    │   │   │   │   │   │   │   │   └── page.tsx  # Escalated disputes
    │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   └── page.tsx  # Pending disputes
    │   │   │   │   │   │   │   └── resolved/
    │   │   │   │   │   │   │       └── page.tsx  # Resolved disputes
    │   │   │   │   │   │   ├── feature-flags/
    │   │   │   │   │   │   │   ├── [flagId]/
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit feature flag
    │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Flag history
    │   │   │   │   │   │   │   │   └── page.tsx  # Flag details
    │   │   │   │   │   │   │   ├── new/
    │   │   │   │   │   │   │   │   └── page.tsx  # Create feature flag
    │   │   │   │   │   │   │   └── page.tsx  # Feature flags list
    │   │   │   │   │   │   ├── financial/
    │   │   │   │   │   │   │   ├── chargebacks/
    │   │   │   │   │   │   │   │   ├── [chargebackId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Chargeback details
    │   │   │   │   │   │   │   │   └── page.tsx  # Chargebacks list
    │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Payment dispute details
    │   │   │   │   │   │   │   │   └── page.tsx  # Payment disputes list
    │   │   │   │   │   │   │   ├── payouts/
    │   │   │   │   │   │   │   │   ├── [payoutId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Payout details
    │   │   │   │   │   │   │   │   ├── batches/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Batch payouts
    │   │   │   │   │   │   │   │   └── page.tsx  # Payouts list
    │   │   │   │   │   │   │   ├── reconciliation/
    │   │   │   │   │   │   │   │   ├── [reportId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Reconciliation report
    │   │   │   │   │   │   │   │   ├── discrepancies/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Discrepancies
    │   │   │   │   │   │   │   │   └── page.tsx  # Reconciliation dashboard
    │   │   │   │   │   │   │   ├── refunds/
    │   │   │   │   │   │   │   │   ├── [refundId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Refund details
    │   │   │   │   │   │   │   │   └── page.tsx  # Refunds list
    │   │   │   │   │   │   │   ├── revenue/
    │   │   │   │   │   │   │   │   └── page.tsx  # Revenue analytics
    │   │   │   │   │   │   │   ├── tax/
    │   │   │   │   │   │   │   │   ├── profiles/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Tax profiles
    │   │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Tax reports
    │   │   │   │   │   │   │   │   └── withholdings/
    │   │   │   │   │   │   │   │       └── page.tsx  # Tax withholdings
    │   │   │   │   │   │   │   └── transactions/
    │   │   │   │   │   │   │       ├── [transactionId]/
    │   │   │   │   │   │   │       │   └── page.tsx  # Transaction details
    │   │   │   │   │   │   │       └── page.tsx  # Transactions list
    │   │   │   │   │   │   ├── moderation/
    │   │   │   │   │   │   │   ├── content/
    │   │   │   │   │   │   │   │   ├── [contentId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Content review
    │   │   │   │   │   │   │   │   ├── appeals/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Content appeals
    │   │   │   │   │   │   │   │   ├── flagged/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Flagged content
    │   │   │   │   │   │   │   │   └── queue/
    │   │   │   │   │   │   │   │       └── page.tsx  # Moderation queue
    │   │   │   │   │   │   │   ├── contracts/
    │   │   │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Contract review
    │   │   │   │   │   │   │   │   └── page.tsx  # Contracts queue
    │   │   │   │   │   │   │   ├── reviews/
    │   │   │   │   │   │   │   │   ├── [reviewId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Review moderation
    │   │   │   │   │   │   │   │   ├── appeals/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Review appeals
    │   │   │   │   │   │   │   │   ├── flagged/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Flagged reviews
    │   │   │   │   │   │   │   │   └── pending/
    │   │   │   │   │   │   │   │       └── page.tsx  # Pending reviews
    │   │   │   │   │   │   │   └── users/
    │   │   │   │   │   │   │       ├── [userId]/
    │   │   │   │   │   │   │       │   └── page.tsx  # User moderation
    │   │   │   │   │   │   │       ├── banned/
    │   │   │   │   │   │   │       │   └── page.tsx  # Banned users
    │   │   │   │   │   │   │       ├── flagged/
    │   │   │   │   │   │   │       │   └── page.tsx  # Flagged users
    │   │   │   │   │   │   │       └── suspended/
    │   │   │   │   │   │   │           └── page.tsx  # Suspended users
    │   │   │   │   │   │   ├── organization/
    │   │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   │   ├── audits/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Compliance audits
    │   │   │   │   │   │   │   │   ├── documents/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Compliance docs
    │   │   │   │   │   │   │   │   └── policies/
    │   │   │   │   │   │   │   │       └── page.tsx  # Policies
    │   │   │   │   │   │   │   ├── hierarchy/
    │   │   │   │   │   │   │   │   ├── chart/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Org chart
    │   │   │   │   │   │   │   │   ├── departments/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Departments
    │   │   │   │   │   │   │   │   └── teams/
    │   │   │   │   │   │   │   │       └── page.tsx  # Teams
    │   │   │   │   │   │   │   ├── integrations/
    │   │   │   │   │   │   │   │   ├── [integrationId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Integration config
    │   │   │   │   │   │   │   │   ├── available/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Available integrations
    │   │   │   │   │   │   │   │   └── installed/
    │   │   │   │   │   │   │   │       └── page.tsx  # Installed integrations
    │   │   │   │   │   │   │   ├── permissions/
    │   │   │   │   │   │   │   │   ├── audit/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Permission audit
    │   │   │   │   │   │   │   │   ├── matrix/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Permission matrix
    │   │   │   │   │   │   │   │   └── roles/
    │   │   │   │   │   │   │   │       └── page.tsx  # Roles
    │   │   │   │   │   │   │   └── procurement/
    │   │   │   │   │   │   │       ├── approvals/
    │   │   │   │   │   │   │       │   └── page.tsx  # Procurement approvals
    │   │   │   │   │   │   │       ├── budgets/
    │   │   │   │   │   │   │       │   └── page.tsx  # Budgets
    │   │   │   │   │   │   │       └── purchase-orders/
    │   │   │   │   │   │   │           └── page.tsx  # Purchase orders
    │   │   │   │   │   │   ├── platform/
    │   │   │   │   │   │   │   ├── api/
    │   │   │   │   │   │   │   │   ├── endpoints/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # API endpoints
    │   │   │   │   │   │   │   │   ├── health/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # API health
    │   │   │   │   │   │   │   │   ├── keys/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # API keys
    │   │   │   │   │   │   │   │   ├── rate-limits/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Rate limits
    │   │   │   │   │   │   │   │   └── webhooks/
    │   │   │   │   │   │   │   │       └── page.tsx  # Webhooks
    │   │   │   │   │   │   │   ├── cache/
    │   │   │   │   │   │   │   │   ├── keys/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Cache keys
    │   │   │   │   │   │   │   │   └── stats/
    │   │   │   │   │   │   │   │       └── page.tsx  # Cache stats
    │   │   │   │   │   │   │   ├── monitoring/
    │   │   │   │   │   │   │   │   ├── alerts/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Alerts
    │   │   │   │   │   │   │   │   ├── dashboards/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Dashboards
    │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Logs
    │   │   │   │   │   │   │   │   ├── metrics/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Metrics
    │   │   │   │   │   │   │   │   └── traces/
    │   │   │   │   │   │   │   │       └── page.tsx  # Traces
    │   │   │   │   │   │   │   ├── queue/
    │   │   │   │   │   │   │   │   ├── dead-letter/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Dead letter queue
    │   │   │   │   │   │   │   │   ├── jobs/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Background jobs
    │   │   │   │   │   │   │   │   └── topics/
    │   │   │   │   │   │   │   │       └── page.tsx  # Queue topics
    │   │   │   │   │   │   │   └── system/
    │   │   │   │   │   │   │       ├── backups/
    │   │   │   │   │   │   │       │   └── page.tsx  # Backups
    │   │   │   │   │   │   │       ├── config/
    │   │   │   │   │   │   │       │   └── page.tsx  # System config
    │   │   │   │   │   │   │       └── maintenance/
    │   │   │   │   │   │   │           └── page.tsx  # Maintenance
    │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   ├── contracts/
    │   │   │   │   │   │   │   │   └── page.tsx  # Contract reports
    │   │   │   │   │   │   │   ├── custom/
    │   │   │   │   │   │   │   │   ├── [reportId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Custom report
    │   │   │   │   │   │   │   │   ├── new/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # New custom report
    │   │   │   │   │   │   │   │   └── page.tsx  # Custom reports
    │   │   │   │   │   │   │   ├── financial/
    │   │   │   │   │   │   │   │   └── page.tsx  # Financial reports
    │   │   │   │   │   │   │   ├── jobs/
    │   │   │   │   │   │   │   │   └── page.tsx  # Job reports
    │   │   │   │   │   │   │   ├── platform/
    │   │   │   │   │   │   │   │   └── page.tsx  # Platform reports
    │   │   │   │   │   │   │   ├── scheduled/
    │   │   │   │   │   │   │   │   └── page.tsx  # Scheduled reports
    │   │   │   │   │   │   │   └── users/
    │   │   │   │   │   │   │       └── page.tsx  # User reports
    │   │   │   │   │   │   ├── search/
    │   │   │   │   │   │   │   ├── categories/
    │   │   │   │   │   │   │   │   ├── [categoryId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Category config
    │   │   │   │   │   │   │   │   ├── new/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # New category
    │   │   │   │   │   │   │   │   └── page.tsx  # Categories
    │   │   │   │   │   │   │   ├── indexing/
    │   │   │   │   │   │   │   │   ├── jobs/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Indexing jobs
    │   │   │   │   │   │   │   │   ├── schemas/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Search schemas
    │   │   │   │   │   │   │   │   └── status/
    │   │   │   │   │   │   │   │       └── page.tsx  # Index status
    │   │   │   │   │   │   │   ├── query-logs/
    │   │   │   │   │   │   │   │   └── page.tsx  # Query logs
    │   │   │   │   │   │   │   └── synonyms/
    │   │   │   │   │   │   │       ├── [synonymId]/
    │   │   │   │   │   │   │       │   └── page.tsx  # Synonym config
    │   │   │   │   │   │   │       ├── new/
    │   │   │   │   │   │   │       │   └── page.tsx  # New synonym
    │   │   │   │   │   │   │       └── page.tsx  # Synonyms
    │   │   │   │   │   │   └── users/
    │   │   │   │   │   │       ├── [userId]/
    │   │   │   │   │   │       │   ├── activity/
    │   │   │   │   │   │       │   │   └── page.tsx  # User activity
    │   │   │   │   │   │       │   ├── contracts/
    │   │   │   │   │   │       │   │   └── page.tsx  # User contracts
    │   │   │   │   │   │       │   ├── financial/
    │   │   │   │   │   │       │   │   └── page.tsx  # User financial
    │   │   │   │   │   │       │   ├── impersonate/
    │   │   │   │   │   │       │   │   └── page.tsx  # Impersonate
    │   │   │   │   │   │       │   ├── security/
    │   │   │   │   │   │       │   │   └── page.tsx  # User security
    │   │   │   │   │   │       │   ├── sessions/
    │   │   │   │   │   │       │   │   └── page.tsx  # User sessions
    │   │   │   │   │   │       │   └── subscriptions/
    │   │   │   │   │   │       │       └── page.tsx  # User subscriptions
    │   │   │   │   │   │       ├── bulk-actions/
    │   │   │   │   │   │       │   └── page.tsx  # Bulk actions
    │   │   │   │   │   │       ├── exports/
    │   │   │   │   │   │       │   └── page.tsx  # User exports
    │   │   │   │   │   │       ├── imports/
    │   │   │   │   │   │       │   └── page.tsx  # User imports
    │   │   │   │   │   │       ├── roles/
    │   │   │   │   │   │       │   ├── [roleId]/
    │   │   │   │   │   │       │   │   └── page.tsx  # Role details
    │   │   │   │   │   │       │   ├── new/
    │   │   │   │   │   │       │   │   └── page.tsx  # New role
    │   │   │   │   │   │       │   └── page.tsx  # Roles list
    │   │   │   │   │   │       └── verification/
    │   │   │   │   │   │           ├── [verificationId]/
    │   │   │   │   │   │           │   └── page.tsx  # Verification review
    │   │   │   │   │   │           ├── identity/
    │   │   │   │   │   │           │   └── page.tsx  # Identity verification
    │   │   │   │   │   │           └── pending/
    │   │   │   │   │   │               └── page.tsx  # Pending verifications
    │   │   │   │   │   ├── currency/
    │   │   │   │   │   │   ├── conversions/
    │   │   │   │   │   │   │   └── page.tsx  # Currency conversions
    │   │   │   │   │   │   ├── hedging/
    │   │   │   │   │   │   │   └── page.tsx  # Currency hedging
    │   │   │   │   │   │   └── rates/
    │   │   │   │   │   │       ├── manual-override/
    │   │   │   │   │   │       │   └── page.tsx  # Manual rate override
    │   │   │   │   │   │       └── page.tsx  # Exchange rates
    │   │   │   │   │   ├── fees/
    │   │   │   │   │   │   ├── overrides/
    │   │   │   │   │   │   │   └── page.tsx  # Fee overrides
    │   │   │   │   │   │   ├── promotions/
    │   │   │   │   │   │   │   ├── [promotionId]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Promotion detail
    │   │   │   │   │   │   │   └── page.tsx  # Fee promotions
    │   │   │   │   │   │   └── structures/
    │   │   │   │   │   │       ├── [structureId]/
    │   │   │   │   │   │       │   └── page.tsx  # Fee structure detail
    │   │   │   │   │   │       └── page.tsx  # Fee structures
    │   │   │   │   │   └── financial/
    │   │   │   │   │       ├── fraud/
    │   │   │   │   │       │   ├── cases/
    │   │   │   │   │       │   │   ├── [caseId]/
    │   │   │   │   │       │   │   │   ├── evidence/
    │   │   │   │   │       │   │   │   │   └── page.tsx  # Case evidence
    │   │   │   │   │       │   │   │   ├── investigation/
    │   │   │   │   │       │   │   │   │   └── page.tsx  # Case investigation
    │   │   │   │   │       │   │   │   └── resolution/
    │   │   │   │   │       │   │   │       └── page.tsx  # Case resolution
    │   │   │   │   │       │   │   └── page.tsx  # Fraud cases
    │   │   │   │   │       │   ├── detection/
    │   │   │   │   │       │   │   ├── alerts/
    │   │   │   │   │       │   │   │   ├── [alertId]/
    │   │   │   │   │       │   │   │   │   └── page.tsx  # Alert investigation
    │   │   │   │   │       │   │   │   └── page.tsx  # Fraud alerts
    │   │   │   │   │       │   │   ├── ml-models/
    │   │   │   │   │       │   │   │   └── page.tsx  # ML fraud models
    │   │   │   │   │       │   │   └── rules/
    │   │   │   │   │       │   │       └── page.tsx  # Fraud detection rules
    │   │   │   │   │       │   ├── patterns/
    │   │   │   │   │       │   └── page.tsx  # Fraud pattern analysis
    │   │   │   │   │       ├── reserves/
    │   │   │   │   │       │   ├── adjustments/
    │   │   │   │   │       │   │   └── page.tsx  # Reserve adjustments
    │   │   │   │   │       │   ├── calculation/
    │   │   │   │   │       │   │   └── page.tsx  # Reserve calculation
    │   │   │   │   │       │   └── releases/
    │   │   │   │   │       │       └── page.tsx  # Reserve releases
    │   │   │   │   │       └── settlement/
    │   │   │   │   │           ├── batches/
    │   │   │   │   │           │   ├── [batchId]/
    │   │   │   │   │           │   │   ├── approve/
    │   │   │   │   │           │   │   │   └── page.tsx  # Approve settlement
    │   │   │   │   │           │   │   ├── review/
    │   │   │   │   │           │   │   │   └── page.tsx  # Review settlement batch
    │   │   │   │   │           │   │   └── page.tsx  # Batch detail
    │   │   │   │   │           │   └── page.tsx  # Settlement batches
    │   │   │   │   │           ├── holds/
    │   │   │   │   │           │   ├── [holdId]/
    │   │   │   │   │           │   │   ├── release/
    │   │   │   │   │           │   │   │   └── page.tsx  # Release hold
    │   │   │   │   │           │   │   └── page.tsx  # Hold detail
    │   │   │   │   │           │   └── page.tsx  # Payment holds
    │   │   │   │   │           └── rules/
    │   │   │   │   │               └── page.tsx  # Settlement rules
    │   │   │   │   ├── (auth)/
    │   │   │   │   ├── (dashboard)/
    │   │   │   │   │   ├── (routes)/
    │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   ├── contracts/
    │   │   │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Verify exists
    │   │   │   │   │   │   │   │   └── page.tsx  # Verify exists
    │   │   │   │   │   │   │   ├── earnings/
    │   │   │   │   │   │   │   │   ├── [periodId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Verify exists
    │   │   │   │   │   │   │   │   └── page.tsx  # Verify exists
    │   │   │   │   │   │   │   ├── jobs/
    │   │   │   │   │   │   │   │   ├── [jobId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Verify exists
    │   │   │   │   │   │   │   │   └── page.tsx  # Verify exists
    │   │   │   │   │   │   │   ├── profile/
    │   │   │   │   │   │   │   │   └── page.tsx  # Verify exists
    │   │   │   │   │   │   │   ├── proposals/
    │   │   │   │   │   │   │   │   ├── [proposalId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Verify exists
    │   │   │   │   │   │   │   │   └── page.tsx  # Verify exists
    │   │   │   │   │   │   │   └── page.tsx  # Analytics dashboard
    │   │   │   │   │   │   └── settings/
    │   │   │   │   │   │       └── (routes)/
    │   │   │   │   │   │           └── advanced/
    │   │   │   │   │   │               └── developer/
    │   │   │   │   │   │                   ├── oauth-apps/
    │   │   │   │   │   │                   │   └── [appId]/
    │   │   │   │   │   │                   │       ├── authorizations/
    │   │   │   │   │   │                   │       │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │                   │       ├── credentials/
    │   │   │   │   │   │                   │       │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │                   │       ├── edit/
    │   │   │   │   │   │                   │       │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │                   │       ├── logs/
    │   │   │   │   │   │                   │       │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │                   │       └── scopes/
    │   │   │   │   │   │                   │           └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │                   └── webhooks/
    │   │   │   │   │   │                       └── [webhookId]/
    │   │   │   │   │   │                           └── deliveries/
    │   │   │   │   │   │                               ├── [deliveryId]/
    │   │   │   │   │   │                               │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │                               └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   ├── [userId]/
    │   │   │   │   │   │   └── settings/
    │   │   │   │   │   │       └── developer/
    │   │   │   │   │   │           ├── api-keys/
    │   │   │   │   │   │           │   └── [keyId]/
    │   │   │   │   │   │           │       ├── logs/
    │   │   │   │   │   │           │       │   └── page.tsx  # API key usage logs
    │   │   │   │   │   │           │       └── permissions/
    │   │   │   │   │   │           │           └── page.tsx  # API key permissions
    │   │   │   │   │   │           └── oauth-apps/
    │   │   │   │   │   │               ├── [appId]/
    │   │   │   │   │   │               │   ├── authorizations/
    │   │   │   │   │   │               │   │   └── page.tsx  # OAuth app authorizations
    │   │   │   │   │   │               │   ├── credentials/
    │   │   │   │   │   │               │   │   └── page.tsx  # OAuth credentials
    │   │   │   │   │   │               │   └── scopes/
    │   │   │   │   │   │               │       └── page.tsx  # OAuth app scopes
    │   │   │   │   │   │               └── new/
    │   │   │   │   │   │                   └── page.tsx  # Create OAuth app
    │   │   │   │   │   ├── camera/  # THIS (Web equivalents)
    │   │   │   │   │   │   ├── document-scan/
    │   │   │   │   │   │   │   └── page.tsx  # Document scanner (web fallback)
    │   │   │   │   │   │   ├── photo-upload/
    │   │   │   │   │   │   │   └── page.tsx  # Photo upload with cropping / bulk upload
    │   │   │   │   │   │   ├── layout.tsx
    │   │   │   │   │   │   └── page.tsx  # Camera/upload hub
    │   │   │   │   │   ├── offline/  # THIS ENTIRE SECTION
    │   │   │   │   │   │   ├── queue/
    │   │   │   │   │   │   │   └── page.tsx  # View pending actions
    │   │   │   │   │   │   ├── sync/
    │   │   │   │   │   │   │   └── page.tsx  # Sync status & management
    │   │   │   │   │   │   ├── layout.tsx
    │   │   │   │   │   │   └── page.tsx  # Offline status page
    │   │   │   │   │   ├── quick-actions/
    │   │   │   │   │   │   └── page.tsx  # THIS — Quick actions (message/proposal/time/invoice)
    │   │   │   │   │   ├── scanner/  # THIS
    │   │   │   │   │   │   ├── document/
    │   │   │   │   │   │   │   └── page.tsx  # Document scanner (web)
    │   │   │   │   │   │   ├── qr-code/
    │   │   │   │   │   │   │   └── page.tsx  # Webcam QR scanner / upload QR / generate codes
    │   │   │   │   │   │   └── layout.tsx
    │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   ├── authorized-apps/  # ⚠️ VERIFY ENTIRE SECTION
    │   │   │   │   │   │   │   ├── [appId]/
    │   │   │   │   │   │   │   │   ├── permissions/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │   │   │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │   │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │   ├── automation/  # ENTIRE SECTION
    │   │   │   │   │   │   │   ├── rules/
    │   │   │   │   │   │   │   │   ├── [ruleId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit automation rule
    │   │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Rule execution history
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Rule detail view
    │   │   │   │   │   │   │   │   ├── new/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create new automation rule
    │   │   │   │   │   │   │   │   └── page.tsx  # Automation rules list
    │   │   │   │   │   │   │   ├── templates/
    │   │   │   │   │   │   │   │   ├── [templateId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Automation template detail
    │   │   │   │   │   │   │   │   └── page.tsx  # Automation templates list
    │   │   │   │   │   │   │   └── page.tsx  # Automation home
    │   │   │   │   │   │   ├── developer/  # ENTIRE SECTION
    │   │   │   │   │   │   │   ├── api-keys/
    │   │   │   │   │   │   │   │   ├── [keyId]/
    │   │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │   │   │   │   ├── permissions/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │   │   │   │   └── page.tsx  # API key detail
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │   │   │   ├── new/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create API key
    │   │   │   │   │   │   │   │   └── page.tsx  # API keys list
    │   │   │   │   │   │   │   ├── oauth-apps/
    │   │   │   │   │   │   │   │   ├── [appId]/
    │   │   │   │   │   │   │   │   │   ├── authorizations/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth app authorizations
    │   │   │   │   │   │   │   │   │   ├── credentials/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth credentials
    │   │   │   │   │   │   │   │   │   ├── scopes/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth app scopes
    │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth app detail
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │   │   │   ├── new/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create OAuth app
    │   │   │   │   │   │   │   │   └── page.tsx  # OAuth apps list
    │   │   │   │   │   │   │   ├── sandbox/
    │   │   │   │   │   │   │   │   ├── environments/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Sandbox environments
    │   │   │   │   │   │   │   │   └── test-data/
    │   │   │   │   │   │   │   │       └── page.tsx  # Test data management
    │   │   │   │   │   │   │   ├── webhooks/
    │   │   │   │   │   │   │   │   ├── [webhookId]/
    │   │   │   │   │   │   │   │   │   ├── deliveries/
    │   │   │   │   │   │   │   │   │   │   ├── [deliveryId]/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Webhook delivery detail
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Webhook deliveries list
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │   │   │   │   ├── events/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Webhook events config
    │   │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Webhook logs
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Webhook detail
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │   │   │   ├── new/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create webhook
    │   │   │   │   │   │   │   │   └── page.tsx  # Webhooks list
    │   │   │   │   │   │   │   └── page.tsx  # Developer home
    │   │   │   │   │   │   ├── devices/  # ⚠️ VERIFY ENTIRE SECTION
    │   │   │   │   │   │   │   ├── [deviceId]/
    │   │   │   │   │   │   │   │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │   │   └── page.tsx  # Connected devices
    │   │   │   │   │   │   ├── integrations/  # ⚠️ PARTIALLY MISSING (expand existing)
    │   │   │   │   │   │   │   ├── [integration]/
    │   │   │   │   │   │   │   │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │   │   ├── available/
    │   │   │   │   │   │   │   │   └── page.tsx  # Available integrations
    │   │   │   │   │   │   │   ├── calendar/
    │   │   │   │   │   │   │   │   └── page.tsx  # Calendar integration
    │   │   │   │   │   │   │   ├── connected/
    │   │   │   │   │   │   │   │   └── page.tsx  # Connected integrations
    │   │   │   │   │   │   │   ├── drive/
    │   │   │   │   │   │   │   │   └── page.tsx  # Drive integration
    │   │   │   │   │   │   │   ├── slack/
    │   │   │   │   │   │   │   │   └── page.tsx  # Slack integration
    │   │   │   │   │   │   │   ├── time-tracking/
    │   │   │   │   │   │   │   │   └── page.tsx  # Time tracking integration
    │   │   │   │   │   │   │   ├── zapier/
    │   │   │   │   │   │   │   │   └── page.tsx  # Zapier integration
    │   │   │   │   │   │   │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │   ├── labs/  # ⚠️ VERIFY
    │   │   │   │   │   │   │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │   ├── login-history/  # ⚠️ VERIFY
    │   │   │   │   │   │   │   └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │   ├── policies/
    │   │   │   │   │   │   │   ├── [policyId]/
    │   │   │   │   │   │   │   │   ├── attestations/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Policy attestations
    │   │   │   │   │   │   │   │   ├── versions/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Policy versions
    │   │   │   │   │   │   │   │   └── page.tsx  # Policy detail
    │   │   │   │   │   │   │   └── page.tsx  # Policies list
    │   │   │   │   │   │   ├── preferences/
    │   │   │   │   │   │   │   └── accessibility/
    │   │   │   │   │   │   │       └── page.tsx  # ⚠️ VERIFY
    │   │   │   │   │   │   ├── team/  # ⚠️ PARTIALLY MISSING (expand existing)
    │   │   │   │   │   │   │   ├── categories/
    │   │   │   │   │   │   │   │   └── page.tsx  # KB categories
    │   │   │   │   │   │   │   ├── knowledge-base/
    │   │   │   │   │   │   │   │   ├── articles/
    │   │   │   │   │   │   │   │   │   └── [articleId]/
    │   │   │   │   │   │   │   │   │       ├── edit/
    │   │   │   │   │   │   │   │   │       │   └── page.tsx  # Edit KB article
    │   │   │   │   │   │   │   │   │       ├── versions/
    │   │   │   │   │   │   │   │   │       │   └── page.tsx  # Article version history
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Article detail
    │   │   │   │   │   │   │   │   └── page.tsx  # Articles list
    │   │   │   │   │   │   │   └── search/
    │   │   │   │   │   │   │       └── page.tsx  # KB search
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── status/
    │   │   │   │   │   │   └── page.tsx  # System status page
    │   │   │   │   │   ├── support/  # ENTIRE SECTION
    │   │   │   │   │   │   └── tickets/
    │   │   │   │   │   │       ├── [ticketId]/
    │   │   │   │   │   │       │   └── page.tsx  # Support ticket detail
    │   │   │   │   │   │       ├── new/
    │   │   │   │   │   │       │   └── page.tsx  # Create support ticket
    │   │   │   │   │   │       └── page.tsx  # Support tickets list
    │   │   │   │   │   └── today/
    │   │   │   │   │       └── page.tsx  # THIS — Today's schedule/tasks/metrics/activity
    │   │   │   │   ├── (legal)/
    │   │   │   │   ├── (marketing)/
    │   │   │   │   ├── [locale]/  # Internationalized routing (en, ar, zh, hi, de, fr, tr, es, ru)
    │   │   │   │   │   ├── (admin)/
    │   │   │   │   │   │   ├── admin/
    │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   ├── experiments/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # A/B testing dashboard
    │   │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Analytics reports
    │   │   │   │   │   │   │   │   └── user-behavior/
    │   │   │   │   │   │   │   │       └── page.tsx  # User behavior analytics
    │   │   │   │   │   │   │   ├── audit-logs/
    │   │   │   │   │   │   │   │   ├── [logId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Audit log details
    │   │   │   │   │   │   │   │   └── page.tsx  # Audit logs viewer
    │   │   │   │   │   │   │   ├── break-glass/
    │   │   │   │   │   │   │   │   ├── active/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Active admin sessions
    │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve break-glass requests
    │   │   │   │   │   │   │   │   └── request/
    │   │   │   │   │   │   │   │       └── page.tsx  # Request break-glass access
    │   │   │   │   │   │   │   ├── business-verification/
    │   │   │   │   │   │   │   │   ├── [caseId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Business verification case detail
    │   │   │   │   │   │   │   │   ├── [verificationId]/
    │   │   │   │   │   │   │   │   │   ├── reverify/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Request reverification
    │   │   │   │   │   │   │   │   │   ├── review/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Review business verification
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Verification details
    │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending verifications
    │   │   │   │   │   │   │   │   └── page.tsx  # Business verification dashboard
    │   │   │   │   │   │   │   ├── communications/
    │   │   │   │   │   │   │   │   ├── broadcasts/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Broadcast management
    │   │   │   │   │   │   │   │   ├── campaigns/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Campaign management
    │   │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │   │       └── page.tsx  # Template management
    │   │   │   │   │   │   │   ├── communications-ops/
    │   │   │   │   │   │   │   │   ├── broadcasts/
    │   │   │   │   │   │   │   │   │   ├── [broadcastId]/
    │   │   │   │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Broadcast analytics
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit broadcast
    │   │   │   │   │   │   │   │   │   │   ├── schedule/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Schedule broadcast
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Broadcast details
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create broadcast
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Broadcasts dashboard
    │   │   │   │   │   │   │   │   ├── campaigns/
    │   │   │   │   │   │   │   │   │   ├── [campaignId]/
    │   │   │   │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Campaign analytics
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit campaign
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Campaign details
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create campaign
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Campaigns dashboard
    │   │   │   │   │   │   │   │   ├── rate-limits/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Communication rate limits
    │   │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │   │       ├── [templateId]/
    │   │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit template
    │   │   │   │   │   │   │   │       │   ├── logs/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Webhook delivery logs
    │   │   │   │   │   │   │   │       │   ├── test/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Test template
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Template details
    │   │   │   │   │   │   │   │       ├── create/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Create template
    │   │   │   │   │   │   │   │       └── page.tsx  # Templates dashboard
    │   │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   │   ├── aml/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # AML monitoring
    │   │   │   │   │   │   │   │   ├── audit-logs/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Audit log viewer
    │   │   │   │   │   │   │   │   ├── data-exports/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # GDPR data exports
    │   │   │   │   │   │   │   │   ├── data-retention/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Data retention policies
    │   │   │   │   │   │   │   │   └── gdpr/
    │   │   │   │   │   │   │   │       └── page.tsx  # GDPR compliance
    │   │   │   │   │   │   │   ├── content-moderation/
    │   │   │   │   │   │   │   │   ├── flagged/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Flagged content review
    │   │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # User reports
    │   │   │   │   │   │   │   │   └── rules/
    │   │   │   │   │   │   │   │       └── page.tsx  # Moderation rules
    │   │   │   │   │   │   │   ├── dashboard/
    │   │   │   │   │   │   │   │   └── page.tsx  # Admin dashboard
    │   │   │   │   │   │   │   ├── dispute-resolution/
    │   │   │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute case detail
    │   │   │   │   │   │   │   │   ├── arbitration/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Arbitration queue
    │   │   │   │   │   │   │   │   ├── mediation/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Mediation queue
    │   │   │   │   │   │   │   │   └── page.tsx  # Disputes dashboard
    │   │   │   │   │   │   │   ├── feature-flags/
    │   │   │   │   │   │   │   │   ├── experiments/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Experiments management
    │   │   │   │   │   │   │   │   └── page.tsx  # Feature flags dashboard
    │   │   │   │   │   │   │   ├── financial/
    │   │   │   │   │   │   │   │   ├── chargebacks/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Chargeback management
    │   │   │   │   │   │   │   │   ├── escrow-monitoring/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Escrow monitoring
    │   │   │   │   │   │   │   │   ├── fraud-detection/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Fraud detection
    │   │   │   │   │   │   │   │   ├── payouts-review/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Payout review queue
    │   │   │   │   │   │   │   │   ├── reconciliation/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Financial reconciliation
    │   │   │   │   │   │   │   │   ├── refunds/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Refund management
    │   │   │   │   │   │   │   │   └── tax-reporting/
    │   │   │   │   │   │   │   │       └── page.tsx  # Tax reporting
    │   │   │   │   │   │   │   ├── financial-ops/
    │   │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   │   │   │   ├── mediate/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Mediate payment dispute
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute details
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Financial disputes
    │   │   │   │   │   │   │   │   ├── payouts/
    │   │   │   │   │   │   │   │   │   ├── [payoutId]/
    │   │   │   │   │   │   │   │   │   │   ├── retry/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Retry failed payout
    │   │   │   │   │   │   │   │   │   │   ├── review/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Review payout
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Payout details
    │   │   │   │   │   │   │   │   │   ├── failed/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Failed payouts
    │   │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending payouts
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Payouts dashboard
    │   │   │   │   │   │   │   │   ├── reconciliation/
    │   │   │   │   │   │   │   │   │   ├── [reconciliationId]/
    │   │   │   │   │   │   │   │   │   │   ├── resolve/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Resolve reconciliation
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reconciliation details
    │   │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending reconciliations
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Reconciliation dashboard
    │   │   │   │   │   │   │   │   └── tax-forms/
    │   │   │   │   │   │   │   │       ├── [formId]/
    │   │   │   │   │   │   │   │       │   ├── review/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Review tax form
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Tax form details
    │   │   │   │   │   │   │   │       ├── generate/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Generate tax forms
    │   │   │   │   │   │   │   │       └── page.tsx  # Tax forms dashboard
    │   │   │   │   │   │   │   ├── goodwill-credits/
    │   │   │   │   │   │   │   │   ├── [creditId]/
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve goodwill credit
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Goodwill credit details
    │   │   │   │   │   │   │   │   ├── issue/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Issue goodwill credit
    │   │   │   │   │   │   │   │   └── page.tsx  # Goodwill credits dashboard
    │   │   │   │   │   │   │   ├── incidents/
    │   │   │   │   │   │   │   │   ├── [incidentId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Incident detail
    │   │   │   │   │   │   │   │   ├── maintenance/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Maintenance windows
    │   │   │   │   │   │   │   │   └── page.tsx  # Incident management
    │   │   │   │   │   │   │   ├── kyc/
    │   │   │   │   │   │   │   │   ├── [caseId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # KYC case detail
    │   │   │   │   │   │   │   │   ├── dashboard/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # KYC dashboard
    │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending cases
    │   │   │   │   │   │   │   │   └── rejected/
    │   │   │   │   │   │   │   │       └── page.tsx  # Rejected cases
    │   │   │   │   │   │   │   ├── kyc-cases/
    │   │   │   │   │   │   │   │   ├── [caseId]/
    │   │   │   │   │   │   │   │   │   ├── documents/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Case documents viewer
    │   │   │   │   │   │   │   │   │   ├── reopen/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reopen KYC case
    │   │   │   │   │   │   │   │   │   ├── review/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Review KYC case
    │   │   │   │   │   │   │   │   │   └── page.tsx  # KYC case details
    │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # KYC queue
    │   │   │   │   │   │   │   │   └── page.tsx  # KYC cases dashboard
    │   │   │   │   │   │   │   ├── moderation/
    │   │   │   │   │   │   │   │   ├── actions/
    │   │   │   │   │   │   │   │   │   ├── [actionId]/
    │   │   │   │   │   │   │   │   │   │   ├── appeal/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Review appeal
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Action details
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Moderation actions
    │   │   │   │   │   │   │   │   ├── patterns/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Abuse patterns detection
    │   │   │   │   │   │   │   │   └── reports/
    │   │   │   │   │   │   │   │       ├── [reportId]/
    │   │   │   │   │   │   │   │       │   ├── review/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Review report
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Report details
    │   │   │   │   │   │   │   │       ├── queue/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Moderation queue
    │   │   │   │   │   │   │   │       └── page.tsx  # Reports dashboard
    │   │   │   │   │   │   │   ├── performance/
    │   │   │   │   │   │   │   │   ├── alerts/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Performance alerts
    │   │   │   │   │   │   │   │   ├── monitoring/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # System monitoring
    │   │   │   │   │   │   │   │   └── page.tsx  # Performance dashboard
    │   │   │   │   │   │   │   ├── search-quality/
    │   │   │   │   │   │   │   │   ├── blacklists/
    │   │   │   │   │   │   │   │   │   ├── [blacklistId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Blacklist entry details
    │   │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Add blacklist entry
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Blacklist management
    │   │   │   │   │   │   │   │   ├── boosts/
    │   │   │   │   │   │   │   │   │   ├── [boostId]/
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit boost rule
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Boost details
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create boost rule
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Boost rules management
    │   │   │   │   │   │   │   │   ├── reindex/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Reindex operations
    │   │   │   │   │   │   │   │   ├── synonyms/
    │   │   │   │   │   │   │   │   │   ├── [synonymId]/
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit synonym
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Synonym details
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create synonym
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Synonyms management
    │   │   │   │   │   │   │   │   ├── tuning/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Search tuning
    │   │   │   │   │   │   │   │   └── page.tsx  # Search quality dashboard
    │   │   │   │   │   │   │   ├── security/
    │   │   │   │   │   │   │   │   ├── sessions/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Break-glass sessions
    │   │   │   │   │   │   │   │   ├── two-person-rule/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Two-person approval
    │   │   │   │   │   │   │   │   └── page.tsx  # Security dashboard
    │   │   │   │   │   │   │   ├── support/
    │   │   │   │   │   │   │   │   ├── canned-responses/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Canned responses
    │   │   │   │   │   │   │   │   └── tickets/
    │   │   │   │   │   │   │   │       ├── [ticketId]/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Support ticket detail
    │   │   │   │   │   │   │   │       └── page.tsx  # Support ticket queue
    │   │   │   │   │   │   │   ├── system/
    │   │   │   │   │   │   │   │   ├── configuration/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # System configuration
    │   │   │   │   │   │   │   │   ├── experiments/
    │   │   │   │   │   │   │   │   │   ├── [experimentId]/
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit experiment
    │   │   │   │   │   │   │   │   │   │   ├── results/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Experiment results
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Experiment details
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create A/B experiment
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Experiments dashboard
    │   │   │   │   │   │   │   │   ├── feature-flags/
    │   │   │   │   │   │   │   │   │   ├── [flagId]/
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit feature flag
    │   │   │   │   │   │   │   │   │   │   ├── rollout/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Manage rollout
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Feature flag details
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create feature flag
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Feature flags dashboard
    │   │   │   │   │   │   │   │   ├── health/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # System health dashboard
    │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # System logs viewer
    │   │   │   │   │   │   │   │   └── metrics/
    │   │   │   │   │   │   │   │       └── page.tsx  # System metrics
    │   │   │   │   │   │   │   ├── system-health/
    │   │   │   │   │   │   │   │   ├── database/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Database health
    │   │   │   │   │   │   │   │   ├── infrastructure/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Infrastructure status
    │   │   │   │   │   │   │   │   └── services/
    │   │   │   │   │   │   │   │       └── page.tsx  # Service health
    │   │   │   │   │   │   │   ├── trust-safety/
    │   │   │   │   │   │   │   │   ├── appeals/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # User appeals
    │   │   │   │   │   │   │   │   ├── bans/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # User bans
    │   │   │   │   │   │   │   │   ├── warnings/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # User warnings
    │   │   │   │   │   │   │   │   └── page.tsx  # Trust & safety dashboard
    │   │   │   │   │   │   │   ├── two-person-rules/
    │   │   │   │   │   │   │   │   ├── [ruleId]/
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve rule change
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Rule details
    │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending approvals
    │   │   │   │   │   │   │   │   └── page.tsx  # Two-person rules dashboard
    │   │   │   │   │   │   │   ├── user-management/
    │   │   │   │   │   │   │   │   ├── [userId]/
    │   │   │   │   │   │   │   │   │   ├── activity/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # User activity
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit user
    │   │   │   │   │   │   │   │   │   └── page.tsx  # User detail
    │   │   │   │   │   │   │   │   ├── impersonate/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # User impersonation
    │   │   │   │   │   │   │   │   └── page.tsx  # User management
    │   │   │   │   │   │   │   ├── webhooks/
    │   │   │   │   │   │   │   │   ├── [webhookId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Webhook detail
    │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Webhook logs
    │   │   │   │   │   │   │   │   └── page.tsx  # Webhook management
    │   │   │   │   │   │   │   └── page.tsx  # Admin dashboard
    │   │   │   │   │   │   ├── admin-session/
    │   │   │   │   │   │   │   └── page.tsx  # JIT break-glass admin sessions
    │   │   │   │   │   │   ├── audit-reporting/
    │   │   │   │   │   │   │   ├── audit-logs/
    │   │   │   │   │   │   │   │   └── page.tsx  # Platform audit trail viewer
    │   │   │   │   │   │   │   ├── bi-exports/
    │   │   │   │   │   │   │   │   └── page.tsx  # BI extracts (CSV/Parquet) request/download
    │   │   │   │   │   │   │   └── page.tsx  # Audit & reporting
    │   │   │   │   │   │   ├── business-verification/
    │   │   │   │   │   │   │   └── page.tsx  # Business Verification cases
    │   │   │   │   │   │   ├── case-mgmt/
    │   │   │   │   │   │   │   └── page.tsx  # Disputes & appeals caseboard (cross-domain)
    │   │   │   │   │   │   ├── change-approval/
    │   │   │   │   │   │   │   └── page.tsx  # Two-person approvals (configs/financial)
    │   │   │   │   │   │   ├── comms/
    │   │   │   │   │   │   │   ├── broadcasts/
    │   │   │   │   │   │   │   │   └── page.tsx  # Broadcasts (email/push/SMS)
    │   │   │   │   │   │   │   ├── campaigns/
    │   │   │   │   │   │   │   │   └── page.tsx  # Lifecycle/marketing campaigns
    │   │   │   │   │   │   │   ├── rate-limits/
    │   │   │   │   │   │   │   │   └── page.tsx  # Throughput & quotas
    │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │       └── page.tsx  # Template library (versioning/AB)
    │   │   │   │   │   │   ├── communications/
    │   │   │   │   │   │   │   ├── broadcasts/
    │   │   │   │   │   │   │   │   └── page.tsx  # Broadcast messages
    │   │   │   │   │   │   │   ├── campaigns/
    │   │   │   │   │   │   │   │   └── page.tsx  # Multi-step campaigns
    │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │       └── page.tsx  # Templates management
    │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   ├── aml-kyc/
    │   │   │   │   │   │   │   │   ├── monitoring/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # AML monitoring dashboard
    │   │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   │   ├── [reportId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # SAR (Suspicious Activity Report) detail
    │   │   │   │   │   │   │   │   │   └── page.tsx  # AML reports list
    │   │   │   │   │   │   │   │   └── risk-assessment/
    │   │   │   │   │   │   │   │       └── page.tsx  # Risk assessment tools
    │   │   │   │   │   │   │   ├── data-retention/
    │   │   │   │   │   │   │   │   ├── audit/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Retention audit log
    │   │   │   │   │   │   │   │   ├── policies/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Retention policies
    │   │   │   │   │   │   │   │   └── schedule/
    │   │   │   │   │   │   │   │       └── page.tsx  # Deletion schedule
    │   │   │   │   │   │   │   ├── document-verification/
    │   │   │   │   │   │   │   │   ├── [documentId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Document review interface
    │   │   │   │   │   │   │   │   ├── automated-checks/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Automated verification rules
    │   │   │   │   │   │   │   │   └── queue/
    │   │   │   │   │   │   │   │       └── page.tsx  # Document verification queue
    │   │   │   │   │   │   │   ├── gdpr/
    │   │   │   │   │   │   │   │   ├── consent-management/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Consent logs & management
    │   │   │   │   │   │   │   │   ├── deletion-requests/
    │   │   │   │   │   │   │   │   │   ├── [requestId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Deletion request detail
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Deletion requests queue
    │   │   │   │   │   │   │   │   ├── export-requests/
    │   │   │   │   │   │   │   │   │   ├── [requestId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Export request detail
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Export requests queue
    │   │   │   │   │   │   │   │   └── reports/
    │   │   │   │   │   │   │   │       └── page.tsx  # GDPR compliance reports
    │   │   │   │   │   │   │   ├── incidents/
    │   │   │   │   │   │   │   │   ├── [incidentId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit incident details
    │   │   │   │   │   │   │   │   │   ├── timeline/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Incident timeline
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Incident detail
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create new incident
    │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Historical incidents
    │   │   │   │   │   │   │   │   └── page.tsx  # Active incidents dashboard
    │   │   │   │   │   │   │   ├── maintenance/
    │   │   │   │   │   │   │   │   ├── [maintenanceId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit maintenance window
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Maintenance detail
    │   │   │   │   │   │   │   │   ├── schedule/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Schedule maintenance
    │   │   │   │   │   │   │   │   └── page.tsx  # Maintenance calendar
    │   │   │   │   │   │   │   └── system-health/
    │   │   │   │   │   │   │       ├── metrics/
    │   │   │   │   │   │   │       │   └── page.tsx  # System metrics dashboard
    │   │   │   │   │   │   │       ├── services/
    │   │   │   │   │   │   │       │   └── page.tsx  # Service health overview
    │   │   │   │   │   │   │       └── page.tsx  # Health dashboard
    │   │   │   │   │   │   ├── experiments/
    │   │   │   │   │   │   │   └── page.tsx  # Experiments dashboard (assignments & metrics)
    │   │   │   │   │   │   ├── feature-flags/
    │   │   │   │   │   │   │   └── page.tsx  # Kill switches & gradual rollout UI
    │   │   │   │   │   │   ├── financial/
    │   │   │   │   │   │   │   └── tax/
    │   │   │   │   │   │   │       └── page.tsx  # Tax ops (admin)
    │   │   │   │   │   │   ├── financial-ops/
    │   │   │   │   │   │   │   ├── chargebacks/
    │   │   │   │   │   │   │   │   └── page.tsx  # Chargeback handling & evidence
    │   │   │   │   │   │   │   ├── payouts-review/
    │   │   │   │   │   │   │   │   └── page.tsx  # Payout approvals & holds
    │   │   │   │   │   │   │   └── reconciliation/
    │   │   │   │   │   │   │       └── page.tsx  # Ledger ↔ gateway reconciliation
    │   │   │   │   │   │   ├── kyc-cases/
    │   │   │   │   │   │   │   └── page.tsx  # KYC case management
    │   │   │   │   │   │   ├── moderation/
    │   │   │   │   │   │   │   ├── [reportId]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Review → warn/suspend/ban → notes & audit
    │   │   │   │   │   │   │   ├── actions/
    │   │   │   │   │   │   │   │   └── page.tsx  # Enforcement actions
    │   │   │   │   │   │   │   ├── appeals/
    │   │   │   │   │   │   │   │   └── page.tsx  # Appeals review
    │   │   │   │   │   │   │   ├── review-queue/
    │   │   │   │   │   │   │   │   └── page.tsx  # Moderation queue
    │   │   │   │   │   │   │   └── page.tsx  # Reports queue & filters
    │   │   │   │   │   │   ├── ops/
    │   │   │   │   │   │   │   ├── admin-session/
    │   │   │   │   │   │   │   │   └── page.tsx  # JIT “break-glass” admin session
    │   │   │   │   │   │   │   ├── change-approval/
    │   │   │   │   │   │   │   │   └── page.tsx  # Two-person change approvals
    │   │   │   │   │   │   │   └── refund-cases/
    │   │   │   │   │   │   │       ├── [caseId]/
    │   │   │   │   │   │   │       │   └── page.tsx  # Case detail → approve/deny → post to ledger
    │   │   │   │   │   │   │       └── page.tsx  # Intake & queue
    │   │   │   │   │   │   ├── platform-config/
    │   │   │   │   │   │   │   ├── integrations/
    │   │   │   │   │   │   │   │   ├── [integrationId]/
    │   │   │   │   │   │   │   │   │   ├── configure/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Configure integration
    │   │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Integration logs
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Integration detail
    │   │   │   │   │   │   │   │   └── page.tsx  # Integrations list
    │   │   │   │   │   │   │   ├── localization/
    │   │   │   │   │   │   │   │   ├── languages/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Language management
    │   │   │   │   │   │   │   │   └── regions/
    │   │   │   │   │   │   │   │       └── page.tsx  # Regional settings
    │   │   │   │   │   │   │   ├── notifications/
    │   │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Notification settings
    │   │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │   │       ├── [templateId]/
    │   │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit notification template
    │   │   │   │   │   │   │   │       │   ├── preview/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Preview template
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Template detail
    │   │   │   │   │   │   │   │       └── page.tsx  # Template library
    │   │   │   │   │   │   │   └── pricing/
    │   │   │   │   │   │   │       └── page.tsx  # Pricing configuration
    │   │   │   │   │   │   ├── refund-cases/
    │   │   │   │   │   │   │   └── page.tsx  # Refund & goodwill credits
    │   │   │   │   │   │   ├── search/
    │   │   │   │   │   │   │   └── explain/
    │   │   │   │   │   │   │       └── page.tsx  # Search explainability (admin)
    │   │   │   │   │   │   ├── search-quality/
    │   │   │   │   │   │   │   ├── blacklists/
    │   │   │   │   │   │   │   │   └── page.tsx  # Blocked entities/terms
    │   │   │   │   │   │   │   ├── boosts/
    │   │   │   │   │   │   │   │   └── page.tsx  # Query boosts & pinning
    │   │   │   │   │   │   │   ├── explainability/
    │   │   │   │   │   │   │   │   └── page.tsx  # Explain results for sample queries
    │   │   │   │   │   │   │   ├── facets/
    │   │   │   │   │   │   │   │   └── page.tsx  # Manage facets & weights
    │   │   │   │   │   │   │   ├── facets-filters/
    │   │   │   │   │   │   │   │   └── page.tsx  # Facets & filters config
    │   │   │   │   │   │   │   ├── hygiene/
    │   │   │   │   │   │   │   │   └── page.tsx  # Dedup/archive/visibility ops
    │   │   │   │   │   │   │   ├── indexing/
    │   │   │   │   │   │   │   │   └── page.tsx  # Indexing & backfills
    │   │   │   │   │   │   │   ├── indices/
    │   │   │   │   │   │   │   │   └── page.tsx  # Rollover/snapshots, lifecycle ops
    │   │   │   │   │   │   │   ├── language/
    │   │   │   │   │   │   │   │   └── page.tsx  # Language analyzers
    │   │   │   │   │   │   │   ├── metrics/
    │   │   │   │   │   │   │   │   └── page.tsx  # Search metrics
    │   │   │   │   │   │   │   ├── performance/
    │   │   │   │   │   │   │   │   └── page.tsx  # Metrics & alerts (slow queries, index health)
    │   │   │   │   │   │   │   ├── query-logs/
    │   │   │   │   │   │   │   │   └── page.tsx  # Query logs & filters; view ES explain
    │   │   │   │   │   │   │   ├── rewrites/
    │   │   │   │   │   │   │   │   └── page.tsx  # Query rewrites & synonyms
    │   │   │   │   │   │   │   ├── speller/
    │   │   │   │   │   │   │   │   └── page.tsx  # Speller configuration
    │   │   │   │   │   │   │   └── synonyms/
    │   │   │   │   │   │   │       └── page.tsx  # Manage synonyms, aliases, taxonomy
    │   │   │   │   │   │   ├── status-incidents/
    │   │   │   │   │   │   │   └── page.tsx  # Incidents & maintenance
    │   │   │   │   │   │   └── storage/
    │   │   │   │   │   │       └── page.tsx  # Storage lifecycle (admin)
    │   │   │   │   │   ├── (auth)/  # Authentication pages (no dashboard layout)
    │   │   │   │   │   │   ├── callback/
    │   │   │   │   │   │   │   └── page.tsx  # OAuth callback handler
    │   │   │   │   │   │   ├── forgot-password/
    │   │   │   │   │   │   │   └── page.tsx  # Password reset request
    │   │   │   │   │   │   ├── login/
    │   │   │   │   │   │   │   └── page.tsx  # Login page
    │   │   │   │   │   │   ├── mfa/
    │   │   │   │   │   │   │   ├── setup/
    │   │   │   │   │   │   │   │   └── page.tsx  # MFA setup
    │   │   │   │   │   │   │   └── verify/
    │   │   │   │   │   │   │       └── page.tsx  # MFA verification
    │   │   │   │   │   │   ├── register/
    │   │   │   │   │   │   │   ├── verification/
    │   │   │   │   │   │   │   │   └── page.tsx  # Email verification callback
    │   │   │   │   │   │   │   └── page.tsx  # Registration page
    │   │   │   │   │   │   ├── reset-password/
    │   │   │   │   │   │   │   └── page.tsx  # Password reset form
    │   │   │   │   │   │   └── layout.tsx  # Auth pages layout
    │   │   │   │   │   ├── (billing)/
    │   │   │   │   │   │   ├── invoices/
    │   │   │   │   │   │   │   └── page.tsx  # Invoices
    │   │   │   │   │   │   ├── payouts/
    │   │   │   │   │   │   │   └── page.tsx  # Payouts
    │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   └── page.tsx  # Billing reports (client)
    │   │   │   │   │   │   ├── tax-forms/
    │   │   │   │   │   │   │   └── page.tsx  # Tax forms (freelancer/client)
    │   │   │   │   │   │   └── wallet/
    │   │   │   │   │   │       └── page.tsx  # Wallet
    │   │   │   │   │   ├── (client)/
    │   │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   │   └── profile/
    │   │   │   │   │   │   │       └── page.tsx  # Billing profile (client)
    │   │   │   │   │   │   ├── jobs/
    │   │   │   │   │   │   │   ├── manage/
    │   │   │   │   │   │   │   │   └── page.tsx  # Job management
    │   │   │   │   │   │   │   └── post/
    │   │   │   │   │   │   │       └── page.tsx  # Post a job
    │   │   │   │   │   │   ├── offers/
    │   │   │   │   │   │   │   └── new/
    │   │   │   │   │   │   │       └── [proposalId]/
    │   │   │   │   │   │   │           └── page.tsx  # Create offer from proposal
    │   │   │   │   │   │   ├── org/
    │   │   │   │   │   │   │   └── page.tsx  # Organization & teams
    │   │   │   │   │   │   ├── shortlists/
    │   │   │   │   │   │   │   └── page.tsx  # Talent shortlists
    │   │   │   │   │   │   └── talent/
    │   │   │   │   │   │       └── invite/
    │   │   │   │   │   │           └── [userId]/
    │   │   │   │   │   │               └── page.tsx  # Invite freelancer to apply
    │   │   │   │   │   ├── (contracts)/
    │   │   │   │   │   │   ├── contracts/
    │   │   │   │   │   │   │   └── [id]/
    │   │   │   │   │   │   │       └── page.tsx  # Contract detail & SOW
    │   │   │   │   │   │   ├── escrow/
    │   │   │   │   │   │   │   └── [contractId]/
    │   │   │   │   │   │   │       ├── fund/
    │   │   │   │   │   │   │       │   └── page.tsx  # Fund escrow (client)
    │   │   │   │   │   │   │       └── release/
    │   │   │   │   │   │   │           └── page.tsx  # Release escrow (client)
    │   │   │   │   │   │   └── workroom/
    │   │   │   │   │   │       └── [id]/
    │   │   │   │   │   │           ├── approvals/
    │   │   │   │   │   │           │   └── page.tsx  # Timesheet approvals (client)
    │   │   │   │   │   │           ├── files/
    │   │   │   │   │   │           │   └── page.tsx  # Workroom files tab
    │   │   │   │   │   │           └── milestones/
    │   │   │   │   │   │               └── page.tsx  # Milestones (web)
    │   │   │   │   │   ├── (dashboard)/  # Main authenticated dashboard
    │   │   │   │   │   │   ├── (admin)/
    │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   ├── performance/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Platform performance metrics
    │   │   │   │   │   │   │   │   ├── retention/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # User retention analysis
    │   │   │   │   │   │   │   │   └── revenue/
    │   │   │   │   │   │   │   │       └── page.tsx  # Revenue dashboard
    │   │   │   │   │   │   │   ├── budgets/
    │   │   │   │   │   │   │   │   ├── [budgetId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit budget
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Budget detail
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create new budget
    │   │   │   │   │   │   │   │   └── page.tsx  # Budgets list
    │   │   │   │   │   │   │   ├── moderation/
    │   │   │   │   │   │   │   │   └── appeals/
    │   │   │   │   │   │   │   │       └── page.tsx  # Appeals queue
    │   │   │   │   │   │   │   ├── search-admin/
    │   │   │   │   │   │   │   │   ├── boosts/
    │   │   │   │   │   │   │   │   │   ├── [boostId]/
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit boost rule
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Boost rule detail
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create boost rule
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Boost rules list
    │   │   │   │   │   │   │   │   ├── facets/
    │   │   │   │   │   │   │   │   │   ├── [facetId]/
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit facet configuration
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Facet detail
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create facet
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Facets list
    │   │   │   │   │   │   │   │   ├── rewrites/
    │   │   │   │   │   │   │   │   │   ├── [rewriteId]/
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit query rewrite
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Rewrite rule detail
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create rewrite rule
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Query rewrites list
    │   │   │   │   │   │   │   │   └── synonyms/
    │   │   │   │   │   │   │   │       ├── [synonymId]/
    │   │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit synonym group
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Synonym group detail
    │   │   │   │   │   │   │   │       ├── create/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Create synonym group
    │   │   │   │   │   │   │   │       ├── import/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Import synonyms (CSV/JSON)
    │   │   │   │   │   │   │   │       └── page.tsx  # Synonyms list
    │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   ├── developer/
    │   │   │   │   │   │   │   │   │   ├── oauth-apps/
    │   │   │   │   │   │   │   │   │   │   └── [appId]/
    │   │   │   │   │   │   │   │   │   │       ├── credentials/
    │   │   │   │   │   │   │   │   │   │       │   └── page.tsx  # OAuth app credentials detail
    │   │   │   │   │   │   │   │   │   │       └── scopes/
    │   │   │   │   │   │   │   │   │   │           └── page.tsx  # OAuth app scopes management
    │   │   │   │   │   │   │   │   │   └── sandbox/
    │   │   │   │   │   │   │   │   │       ├── environments/
    │   │   │   │   │   │   │   │   │       │   ├── [envId]/
    │   │   │   │   │   │   │   │   │       │   │   ├── reset/
    │   │   │   │   │   │   │   │   │       │   │   │   └── page.tsx  # Reset sandbox environment
    │   │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Sandbox environment detail
    │   │   │   │   │   │   │   │   │       │   └── page.tsx  # Sandbox environments list
    │   │   │   │   │   │   │   │   │       └── test-data/
    │   │   │   │   │   │   │   │   │           ├── generate/
    │   │   │   │   │   │   │   │   │           │   └── page.tsx  # Generate test data
    │   │   │   │   │   │   │   │   │           └── page.tsx  # Test data management
    │   │   │   │   │   │   │   │   ├── notifications/
    │   │   │   │   │   │   │   │   │   └── channels/
    │   │   │   │   │   │   │   │   │       ├── email/
    │   │   │   │   │   │   │   │   │       │   └── page.tsx  # Email channel settings
    │   │   │   │   │   │   │   │   │       ├── push/
    │   │   │   │   │   │   │   │   │       │   └── page.tsx  # Push notification settings
    │   │   │   │   │   │   │   │   │       ├── sms/
    │   │   │   │   │   │   │   │   │       │   └── page.tsx  # SMS channel settings
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Notification channels overview
    │   │   │   │   │   │   │   │   └── team/
    │   │   │   │   │   │   │   │       ├── categories/
    │   │   │   │   │   │   │   │       │   ├── [categoryId]/
    │   │   │   │   │   │   │   │       │   │   ├── edit/
    │   │   │   │   │   │   │   │       │   │   │   └── page.tsx  # Edit KB category
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # KB category detail
    │   │   │   │   │   │   │   │       │   └── page.tsx  # KB categories list
    │   │   │   │   │   │   │   │       └── knowledge-base/
    │   │   │   │   │   │   │   │           └── articles/
    │   │   │   │   │   │   │   │               └── [articleId]/
    │   │   │   │   │   │   │   │                   ├── edit/
    │   │   │   │   │   │   │   │                   │   └── page.tsx  # Edit KB article
    │   │   │   │   │   │   │   │                   └── versions/
    │   │   │   │   │   │   │   │                       ├── [versionId]/
    │   │   │   │   │   │   │   │                       │   ├── compare/
    │   │   │   │   │   │   │   │                       │   │   └── page.tsx  # Compare article versions
    │   │   │   │   │   │   │   │                       │   └── page.tsx  # Article version detail
    │   │   │   │   │   │   │   │                       └── page.tsx  # Article version history
    │   │   │   │   │   │   │   └── support/
    │   │   │   │   │   │   │       └── analytics/
    │   │   │   │   │   │   │           ├── performance/
    │   │   │   │   │   │   │           │   └── page.tsx  # Support performance metrics
    │   │   │   │   │   │   │           ├── satisfaction/
    │   │   │   │   │   │   │           │   └── page.tsx  # Customer satisfaction metrics
    │   │   │   │   │   │   │           └── page.tsx  # Support analytics overview
    │   │   │   │   │   │   ├── admin/  # Admin panel (SUPER_ADMIN only)
    │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   ├── platform/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Platform-wide analytics
    │   │   │   │   │   │   │   │   └── page.tsx  # Analytics overview
    │   │   │   │   │   │   │   ├── audit-logs/
    │   │   │   │   │   │   │   │   └── page.tsx  # Audit logs
    │   │   │   │   │   │   │   ├── audits/
    │   │   │   │   │   │   │   │   ├── [auditId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Audit detail
    │   │   │   │   │   │   │   │   ├── exports/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Export audit logs
    │   │   │   │   │   │   │   │   └── page.tsx  # Audit logs
    │   │   │   │   │   │   │   ├── business-verification/
    │   │   │   │   │   │   │   │   ├── [caseId]/
    │   │   │   │   │   │   │   │   │   ├── review/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Review business documents
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Business verification case
    │   │   │   │   │   │   │   │   ├── [verificationId]/
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve business
    │   │   │   │   │   │   │   │   │   ├── documents/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Verification documents
    │   │   │   │   │   │   │   │   │   ├── reject/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reject business
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Business verification detail
    │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending verifications
    │   │   │   │   │   │   │   │   └── page.tsx  # Business verification queue
    │   │   │   │   │   │   │   ├── change-approvals/
    │   │   │   │   │   │   │   │   ├── [approvalId]/
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve change
    │   │   │   │   │   │   │   │   │   ├── reject/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reject change
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Change approval detail
    │   │   │   │   │   │   │   │   ├── [requestId]/
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve/reject change
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Change request detail
    │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending approvals
    │   │   │   │   │   │   │   │   └── page.tsx  # Change approval queue (Two-Person Rule)
    │   │   │   │   │   │   │   ├── communications/
    │   │   │   │   │   │   │   │   ├── broadcasts/
    │   │   │   │   │   │   │   │   │   ├── [broadcastId]/
    │   │   │   │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Broadcast analytics
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Broadcast detail
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create broadcast
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Broadcasts list
    │   │   │   │   │   │   │   │   ├── campaigns/
    │   │   │   │   │   │   │   │   │   ├── [campaignId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Campaign detail
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create campaign
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Campaigns list
    │   │   │   │   │   │   │   │   ├── rate-limits/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Communication rate limits
    │   │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │   │       ├── [templateId]/
    │   │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit template
    │   │   │   │   │   │   │   │       │   ├── preview/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Preview template
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Template detail
    │   │   │   │   │   │   │   │       ├── create/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Create template
    │   │   │   │   │   │   │   │       └── page.tsx  # Templates list
    │   │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   │   ├── aml-kyc/
    │   │   │   │   │   │   │   │   │   ├── monitoring/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # AML monitoring dashboard
    │   │   │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   │   │   ├── [reportId]/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # SAR (Suspicious Activity Report) detail
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # AML reports list
    │   │   │   │   │   │   │   │   │   └── risk-assessment/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Risk assessment tools
    │   │   │   │   │   │   │   │   ├── data-retention/
    │   │   │   │   │   │   │   │   │   ├── audit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Retention audit log
    │   │   │   │   │   │   │   │   │   ├── policies/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Retention policies
    │   │   │   │   │   │   │   │   │   └── schedule/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Deletion schedule
    │   │   │   │   │   │   │   │   ├── document-verification/
    │   │   │   │   │   │   │   │   │   ├── [documentId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Document review interface
    │   │   │   │   │   │   │   │   │   ├── automated-checks/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Automated verification rules
    │   │   │   │   │   │   │   │   │   └── queue/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Document verification queue
    │   │   │   │   │   │   │   │   └── gdpr/
    │   │   │   │   │   │   │   │       ├── consent-management/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Consent logs & management
    │   │   │   │   │   │   │   │       ├── deletion-requests/
    │   │   │   │   │   │   │   │       │   ├── [requestId]/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Deletion request detail
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Deletion requests queue
    │   │   │   │   │   │   │   │       ├── export-requests/
    │   │   │   │   │   │   │   │       │   ├── [requestId]/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Export request detail
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Export requests queue
    │   │   │   │   │   │   │   │       └── reports/
    │   │   │   │   │   │   │   │           └── page.tsx  # GDPR compliance reports
    │   │   │   │   │   │   │   ├── configurations/
    │   │   │   │   │   │   │   │   ├── experiments/
    │   │   │   │   │   │   │   │   │   ├── [experimentId]/
    │   │   │   │   │   │   │   │   │   │   ├── results/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Experiment results
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Experiment detail
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create experiment
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Experiments list
    │   │   │   │   │   │   │   │   ├── feature-flags/
    │   │   │   │   │   │   │   │   │   ├── [flagId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Feature flag detail
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Feature flags list
    │   │   │   │   │   │   │   │   └── page.tsx  # System configurations
    │   │   │   │   │   │   │   ├── content-moderation/
    │   │   │   │   │   │   │   │   ├── [reportId]/
    │   │   │   │   │   │   │   │   │   ├── actions/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Moderation actions
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Report detail
    │   │   │   │   │   │   │   │   ├── appeals/
    │   │   │   │   │   │   │   │   │   ├── [appealId]/
    │   │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve appeal
    │   │   │   │   │   │   │   │   │   │   ├── reject/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reject appeal
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Appeal detail
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Appeals queue
    │   │   │   │   │   │   │   │   ├── queue/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Moderation queue
    │   │   │   │   │   │   │   │   └── rules/
    │   │   │   │   │   │   │   │       └── page.tsx  # Moderation rules
    │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   │   │   ├── assign/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Assign dispute to admin
    │   │   │   │   │   │   │   │   │   ├── escalate/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Escalate dispute
    │   │   │   │   │   │   │   │   │   └── resolve/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Resolve dispute
    │   │   │   │   │   │   │   │   └── page.tsx  # Disputes management
    │   │   │   │   │   │   │   ├── financial-ops/
    │   │   │   │   │   │   │   │   ├── chargebacks/
    │   │   │   │   │   │   │   │   │   ├── [chargebackId]/
    │   │   │   │   │   │   │   │   │   │   ├── dispute/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute chargeback
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Chargeback detail
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Chargebacks list
    │   │   │   │   │   │   │   │   ├── goodwill-credits/
    │   │   │   │   │   │   │   │   │   ├── [creditId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Goodwill credit detail
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve goodwill credit
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create goodwill credit
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Goodwill credits list
    │   │   │   │   │   │   │   │   ├── payouts/
    │   │   │   │   │   │   │   │   │   ├── [payoutId]/
    │   │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve payout
    │   │   │   │   │   │   │   │   │   │   ├── hold/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Hold payout
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Payout detail
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Payouts review queue
    │   │   │   │   │   │   │   │   ├── reconciliation/
    │   │   │   │   │   │   │   │   │   ├── [reconId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reconciliation detail
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Start reconciliation
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Reconciliation reports
    │   │   │   │   │   │   │   │   └── refund-cases/
    │   │   │   │   │   │   │   │       ├── [caseId]/
    │   │   │   │   │   │   │   │       │   ├── approve/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Approve refund
    │   │   │   │   │   │   │   │       │   ├── investigation/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Investigation notes
    │   │   │   │   │   │   │   │       │   ├── reject/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Reject refund
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Refund case detail
    │   │   │   │   │   │   │   │       └── page.tsx  # Refund cases queue
    │   │   │   │   │   │   │   ├── incidents/
    │   │   │   │   │   │   │   │   ├── [incidentId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit incident details
    │   │   │   │   │   │   │   │   │   ├── timeline/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Incident timeline
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Incident detail
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create new incident
    │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Historical incidents
    │   │   │   │   │   │   │   │   └── page.tsx  # Active incidents dashboard
    │   │   │   │   │   │   │   ├── kyc/
    │   │   │   │   │   │   │   │   ├── [caseId]/
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve KYC
    │   │   │   │   │   │   │   │   │   └── reject/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Reject KYC
    │   │   │   │   │   │   │   │   ├── [kycId]/
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve KYC
    │   │   │   │   │   │   │   │   │   ├── documents/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # KYC documents review
    │   │   │   │   │   │   │   │   │   ├── reject/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reject KYC
    │   │   │   │   │   │   │   │   │   ├── reopen/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reopen KYC case
    │   │   │   │   │   │   │   │   │   └── page.tsx  # KYC case detail
    │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending KYC cases
    │   │   │   │   │   │   │   │   ├── rejected/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Rejected KYC cases
    │   │   │   │   │   │   │   │   └── page.tsx  # KYC cases queue
    │   │   │   │   │   │   │   ├── maintenance/
    │   │   │   │   │   │   │   │   ├── [maintenanceId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit maintenance window
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Maintenance detail
    │   │   │   │   │   │   │   │   ├── schedule/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Schedule maintenance
    │   │   │   │   │   │   │   │   └── page.tsx  # Maintenance calendar
    │   │   │   │   │   │   │   ├── moderation/
    │   │   │   │   │   │   │   │   ├── jobs/
    │   │   │   │   │   │   │   │   │   ├── [jobId]/
    │   │   │   │   │   │   │   │   │   │   └── review/
    │   │   │   │   │   │   │   │   │   │       └── page.tsx  # Review flagged job
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Flagged jobs
    │   │   │   │   │   │   │   │   ├── messages/
    │   │   │   │   │   │   │   │   │   └── [messageId]/
    │   │   │   │   │   │   │   │   │       └── review/
    │   │   │   │   │   │   │   │   │           └── page.tsx  # Review flagged message
    │   │   │   │   │   │   │   │   ├── proposals/
    │   │   │   │   │   │   │   │   │   └── [proposalId]/
    │   │   │   │   │   │   │   │   │       └── review/
    │   │   │   │   │   │   │   │   │           └── page.tsx  # Review flagged proposal
    │   │   │   │   │   │   │   │   └── reviews/
    │   │   │   │   │   │   │   │       └── [reviewId]/
    │   │   │   │   │   │   │   │           └── review/
    │   │   │   │   │   │   │   │               └── page.tsx  # Review flagged review
    │   │   │   │   │   │   │   ├── platform-config/
    │   │   │   │   │   │   │   │   ├── integrations/
    │   │   │   │   │   │   │   │   │   ├── [integrationId]/
    │   │   │   │   │   │   │   │   │   │   ├── configure/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Configure integration
    │   │   │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Integration logs
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Integration detail
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Integrations list
    │   │   │   │   │   │   │   │   ├── limits/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Platform limits configuration
    │   │   │   │   │   │   │   │   ├── localization/
    │   │   │   │   │   │   │   │   │   ├── languages/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Language management
    │   │   │   │   │   │   │   │   │   └── regions/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Regional settings
    │   │   │   │   │   │   │   │   ├── notifications/
    │   │   │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Notification settings
    │   │   │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │   │   │       ├── [templateId]/
    │   │   │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit notification template
    │   │   │   │   │   │   │   │   │       │   ├── preview/
    │   │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Preview template
    │   │   │   │   │   │   │   │   │       │   └── page.tsx  # Template detail
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Template library
    │   │   │   │   │   │   │   │   └── pricing/
    │   │   │   │   │   │   │   │       └── page.tsx  # Pricing configuration
    │   │   │   │   │   │   │   ├── refunds/
    │   │   │   │   │   │   │   │   ├── [caseId]/
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve refund
    │   │   │   │   │   │   │   │   │   └── reject/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Reject refund
    │   │   │   │   │   │   │   │   └── page.tsx  # Refund cases
    │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   ├── financial/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Financial reports
    │   │   │   │   │   │   │   │   ├── moderation/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Moderation reports
    │   │   │   │   │   │   │   │   ├── users/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # User reports
    │   │   │   │   │   │   │   │   └── page.tsx  # Admin reports
    │   │   │   │   │   │   │   ├── search-quality/
    │   │   │   │   │   │   │   │   ├── blacklists/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Search blacklists
    │   │   │   │   │   │   │   │   ├── boosts/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Search boosts
    │   │   │   │   │   │   │   │   ├── reindex/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Reindex controls
    │   │   │   │   │   │   │   │   └── synonyms/
    │   │   │   │   │   │   │   │       └── page.tsx  # Search synonyms
    │   │   │   │   │   │   │   ├── sessions/
    │   │   │   │   │   │   │   │   ├── [sessionId]/
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve JIT session
    │   │   │   │   │   │   │   │   │   ├── revoke/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Revoke session
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Session detail
    │   │   │   │   │   │   │   │   ├── request/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Request JIT session
    │   │   │   │   │   │   │   │   └── page.tsx  # Admin sessions list
    │   │   │   │   │   │   │   ├── system/
    │   │   │   │   │   │   │   │   ├── announcements/
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create announcement
    │   │   │   │   │   │   │   │   │   └── page.tsx  # System announcements
    │   │   │   │   │   │   │   │   ├── feature-flags/
    │   │   │   │   │   │   │   │   │   ├── [flagId]/
    │   │   │   │   │   │   │   │   │   │   └── edit/
    │   │   │   │   │   │   │   │   │   │       └── page.tsx  # Edit feature flag
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Feature flags management
    │   │   │   │   │   │   │   │   ├── maintenance/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Maintenance mode
    │   │   │   │   │   │   │   │   └── page.tsx  # System settings
    │   │   │   │   │   │   │   ├── system-health/
    │   │   │   │   │   │   │   │   ├── incidents/
    │   │   │   │   │   │   │   │   │   ├── [incidentId]/
    │   │   │   │   │   │   │   │   │   │   ├── postmortem/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Incident postmortem
    │   │   │   │   │   │   │   │   │   │   ├── resolve/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Resolve incident
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Incident detail
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create incident
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Incidents list
    │   │   │   │   │   │   │   │   ├── maintenance/
    │   │   │   │   │   │   │   │   │   ├── [maintenanceId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Maintenance detail
    │   │   │   │   │   │   │   │   │   ├── schedule/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Schedule maintenance
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Maintenance windows
    │   │   │   │   │   │   │   │   ├── metrics/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # System metrics dashboard
    │   │   │   │   │   │   │   │   ├── services/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Service health overview
    │   │   │   │   │   │   │   │   ├── status/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # System status dashboard
    │   │   │   │   │   │   │   │   └── page.tsx  # Health dashboard
    │   │   │   │   │   │   │   ├── users/
    │   │   │   │   │   │   │   │   ├── [userId]/
    │   │   │   │   │   │   │   │   │   ├── ban/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Ban user
    │   │   │   │   │   │   │   │   │   ├── contracts/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # User contracts
    │   │   │   │   │   │   │   │   │   ├── financials/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # User financial history
    │   │   │   │   │   │   │   │   │   ├── impersonate/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Impersonate user
    │   │   │   │   │   │   │   │   │   ├── suspend/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Suspend user
    │   │   │   │   │   │   │   │   │   ├── unban/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Unban user
    │   │   │   │   │   │   │   │   │   ├── verify/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Verify user (manual)
    │   │   │   │   │   │   │   │   │   ├── warn/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Warn user
    │   │   │   │   │   │   │   │   │   ├── warning/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Issue warning
    │   │   │   │   │   │   │   │   │   └── page.tsx  # User detail
    │   │   │   │   │   │   │   │   ├── banned/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Banned users
    │   │   │   │   │   │   │   │   ├── bulk-actions/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Bulk user actions
    │   │   │   │   │   │   │   │   ├── search/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Search users
    │   │   │   │   │   │   │   │   ├── suspended/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Suspended users
    │   │   │   │   │   │   │   │   └── page.tsx  # Users management
    │   │   │   │   │   │   │   ├── layout.tsx  # Admin layout (RBAC guard)
    │   │   │   │   │   │   │   └── page.tsx  # Admin dashboard
    │   │   │   │   │   │   ├── agency/
    │   │   │   │   │   │   │   ├── overview/
    │   │   │   │   │   │   │   │   └── page.tsx  # Agency dashboard
    │   │   │   │   │   │   │   ├── reporting/
    │   │   │   │   │   │   │   │   ├── clients/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Client reports
    │   │   │   │   │   │   │   │   ├── financial/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Financial reports
    │   │   │   │   │   │   │   │   └── team/
    │   │   │   │   │   │   │   │       └── page.tsx  # Team performance
    │   │   │   │   │   │   │   ├── sub-accounts/
    │   │   │   │   │   │   │   │   ├── [subAccountId]/
    │   │   │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Sub-account settings
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Sub-account detail
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create sub-account
    │   │   │   │   │   │   │   │   └── page.tsx  # All sub-accounts
    │   │   │   │   │   │   │   ├── talent-pool/
    │   │   │   │   │   │   │   │   ├── [poolId]/
    │   │   │   │   │   │   │   │   │   ├── members/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Pool members
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Talent pool detail
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create talent pool
    │   │   │   │   │   │   │   │   └── page.tsx  # All talent pools
    │   │   │   │   │   │   │   └── white-label/
    │   │   │   │   │   │   │       └── page.tsx  # White-label settings
    │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   ├── advanced/
    │   │   │   │   │   │   │   │   └── page.tsx  # Advanced analytics
    │   │   │   │   │   │   │   ├── clients/
    │   │   │   │   │   │   │   │   └── page.tsx  # Client analytics (freelancer)
    │   │   │   │   │   │   │   ├── cohorts/
    │   │   │   │   │   │   │   │   └── page.tsx  # Cohort analysis
    │   │   │   │   │   │   │   ├── conversion/
    │   │   │   │   │   │   │   │   └── page.tsx  # Conversion funnel
    │   │   │   │   │   │   │   ├── custom-reports/
    │   │   │   │   │   │   │   │   ├── [reportId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit custom report
    │   │   │   │   │   │   │   │   │   └── page.tsx  # View custom report
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Report builder
    │   │   │   │   │   │   │   │   ├── new/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create custom report
    │   │   │   │   │   │   │   │   └── page.tsx  # Custom reports list
    │   │   │   │   │   │   │   ├── earnings/
    │   │   │   │   │   │   │   │   ├── forecast/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Earnings forecast
    │   │   │   │   │   │   │   │   └── page.tsx  # Earnings analytics
    │   │   │   │   │   │   │   ├── exports/
    │   │   │   │   │   │   │   │   └── page.tsx  # Data exports
    │   │   │   │   │   │   │   ├── freelancers/
    │   │   │   │   │   │   │   │   └── page.tsx  # Freelancer analytics (client)
    │   │   │   │   │   │   │   ├── funnels/
    │   │   │   │   │   │   │   │   └── page.tsx  # Funnel analysis
    │   │   │   │   │   │   │   ├── market-insights/
    │   │   │   │   │   │   │   │   └── page.tsx  # Market insights
    │   │   │   │   │   │   │   ├── performance/
    │   │   │   │   │   │   │   │   └── page.tsx  # Performance analytics
    │   │   │   │   │   │   │   ├── projects/
    │   │   │   │   │   │   │   │   └── page.tsx  # Project analytics
    │   │   │   │   │   │   │   ├── retention/
    │   │   │   │   │   │   │   │   └── page.tsx  # Retention analysis
    │   │   │   │   │   │   │   ├── revenue/
    │   │   │   │   │   │   │   │   └── page.tsx  # Revenue analytics
    │   │   │   │   │   │   │   ├── spending/
    │   │   │   │   │   │   │   │   ├── forecast/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Spending forecast
    │   │   │   │   │   │   │   │   └── page.tsx  # Spending analytics (client)
    │   │   │   │   │   │   │   ├── user-segmentation/
    │   │   │   │   │   │   │   │   └── page.tsx  # User segments
    │   │   │   │   │   │   │   └── page.tsx  # Analytics dashboard
    │   │   │   │   │   │   ├── availability/
    │   │   │   │   │   │   │   ├── calendar/
    │   │   │   │   │   │   │   │   └── page.tsx  # Availability calendar
    │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   └── page.tsx  # Availability settings
    │   │   │   │   │   │   │   └── page.tsx  # Availability dashboard
    │   │   │   │   │   │   ├── bidding/
    │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   └── page.tsx  # Bidding analytics
    │   │   │   │   │   │   │   ├── auctions/
    │   │   │   │   │   │   │   │   ├── [auctionId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Auction participation
    │   │   │   │   │   │   │   │   └── page.tsx  # Active auctions list
    │   │   │   │   │   │   │   └── strategies/
    │   │   │   │   │   │   │       ├── [strategyId]/
    │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit bid strategy
    │   │   │   │   │   │   │       │   └── page.tsx  # View bid strategy details
    │   │   │   │   │   │   │       ├── new/
    │   │   │   │   │   │   │       │   └── page.tsx  # Create new bid strategy
    │   │   │   │   │   │   │       └── page.tsx  # Bid strategies list
    │   │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   │   ├── budgets/
    │   │   │   │   │   │   │   │   └── page.tsx  # Budgets
    │   │   │   │   │   │   │   ├── subscription/
    │   │   │   │   │   │   │   │   ├── dunning/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Failed invoices & retries
    │   │   │   │   │   │   │   │   ├── manage/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Manage current subscription
    │   │   │   │   │   │   │   │   └── plans/
    │   │   │   │   │   │   │   │       └── page.tsx  # Subscription plans & checkout
    │   │   │   │   │   │   │   └── subscriptions/
    │   │   │   │   │   │   │       ├── history/
    │   │   │   │   │   │   │       │   └── page.tsx  # Subscription history
    │   │   │   │   │   │   │       └── manage/
    │   │   │   │   │   │   │           └── page.tsx  # Manage subscription
    │   │   │   │   │   │   ├── budgets/
    │   │   │   │   │   │   │   └── page.tsx  # Client org budgets & spend controls
    │   │   │   │   │   │   ├── camera/  # ENTIRE SECTION
    │   │   │   │   │   │   │   ├── document-scan/
    │   │   │   │   │   │   │   │   └── page.tsx  # Document scanner (web)
    │   │   │   │   │   │   │   ├── photo-upload/
    │   │   │   │   │   │   │   │   └── page.tsx  # Photo upload with tools
    │   │   │   │   │   │   │   ├── layout.tsx  # Camera/upload layout
    │   │   │   │   │   │   │   └── page.tsx  # Camera/upload hub
    │   │   │   │   │   │   ├── collaboration/
    │   │   │   │   │   │   │   ├── boards/
    │   │   │   │   │   │   │   │   ├── [boardId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Collaboration board
    │   │   │   │   │   │   │   │   └── page.tsx  # All boards
    │   │   │   │   │   │   │   ├── documents/
    │   │   │   │   │   │   │   │   ├── [documentId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Document editor
    │   │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Version history
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Document viewer
    │   │   │   │   │   │   │   │   ├── shared/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Shared documents
    │   │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │   │       └── page.tsx  # Document templates
    │   │   │   │   │   │   │   ├── meetings/
    │   │   │   │   │   │   │   │   ├── [meetingId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Meeting detail
    │   │   │   │   │   │   │   │   ├── recordings/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Meeting recordings
    │   │   │   │   │   │   │   │   └── schedule/
    │   │   │   │   │   │   │   │       └── page.tsx  # Schedule meeting
    │   │   │   │   │   │   │   └── whiteboard/
    │   │   │   │   │   │   │       ├── [whiteboardId]/
    │   │   │   │   │   │   │       │   └── page.tsx  # Whiteboard editor
    │   │   │   │   │   │   │       └── page.tsx  # Whiteboards
    │   │   │   │   │   │   ├── community/
    │   │   │   │   │   │   │   ├── events/
    │   │   │   │   │   │   │   │   ├── [eventId]/
    │   │   │   │   │   │   │   │   │   ├── register/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Event registration
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Event detail
    │   │   │   │   │   │   │   │   ├── my-events/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # My registered events
    │   │   │   │   │   │   │   │   └── upcoming/
    │   │   │   │   │   │   │   │       └── page.tsx  # Upcoming events
    │   │   │   │   │   │   │   ├── forums/
    │   │   │   │   │   │   │   │   ├── [forumId]/
    │   │   │   │   │   │   │   │   │   ├── threads/
    │   │   │   │   │   │   │   │   │   │   └── [threadId]/
    │   │   │   │   │   │   │   │   │   │       ├── reply/
    │   │   │   │   │   │   │   │   │   │       │   └── page.tsx  # Reply to thread
    │   │   │   │   │   │   │   │   │   │       └── page.tsx  # Thread view
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Forum overview
    │   │   │   │   │   │   │   │   └── page.tsx  # All forums
    │   │   │   │   │   │   │   └── groups/
    │   │   │   │   │   │   │       └── [groupId]/
    │   │   │   │   │   │   │           ├── discussions/
    │   │   │   │   │   │   │           │   └── page.tsx  # Group discussions
    │   │   │   │   │   │   │           ├── events/
    │   │   │   │   │   │   │           │   ├── [eventId]/
    │   │   │   │   │   │   │           │   │   └── page.tsx  # Event detail
    │   │   │   │   │   │   │           │   └── page.tsx  # Group events
    │   │   │   │   │   │   │           ├── members/
    │   │   │   │   │   │   │           │   └── page.tsx  # Group members
    │   │   │   │   │   │   │           └── page.tsx  # Group overview
    │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   ├── documents/
    │   │   │   │   │   │   │   │   ├── [documentId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Compliance document details
    │   │   │   │   │   │   │   │   ├── upload/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Upload compliance documents
    │   │   │   │   │   │   │   │   └── page.tsx  # Compliance documents list
    │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   ├── tax-summary/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Annual tax summary
    │   │   │   │   │   │   │   │   └── page.tsx  # Compliance reports
    │   │   │   │   │   │   │   └── tax-profile/
    │   │   │   │   │   │   │       ├── edit/
    │   │   │   │   │   │   │       │   └── page.tsx  # Edit tax profile
    │   │   │   │   │   │   │       └── page.tsx  # Tax profile overview
    │   │   │   │   │   │   ├── connects/
    │   │   │   │   │   │   │   ├── purchase/
    │   │   │   │   │   │   │   │   └── page.tsx  # Purchase connects
    │   │   │   │   │   │   │   ├── usage/
    │   │   │   │   │   │   │   │   └── page.tsx  # Connects usage analytics
    │   │   │   │   │   │   │   └── page.tsx  # Connects dashboard
    │   │   │   │   │   │   ├── contracts/  # Contracts management
    │   │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   │   ├── acceptance/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Accept deliverables/milestones
    │   │   │   │   │   │   │   │   ├── amendments/
    │   │   │   │   │   │   │   │   │   ├── [amendmentId]/
    │   │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve/reject amendment
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Amendment detail
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create amendment
    │   │   │   │   │   │   │   │   │   ├── propose/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Propose amendment
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Contract amendments list
    │   │   │   │   │   │   │   │   ├── approvals/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Approval workflow
    │   │   │   │   │   │   │   │   ├── audit-trail/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Complete contract audit trail
    │   │   │   │   │   │   │   │   ├── budget-tracking/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Budget vs actual
    │   │   │   │   │   │   │   │   ├── change-orders/
    │   │   │   │   │   │   │   │   │   ├── [orderId]/
    │   │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve change order
    │   │   │   │   │   │   │   │   │   │   ├── reject/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reject change order
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Change order details
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create change order
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Change orders list
    │   │   │   │   │   │   │   │   ├── complete/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Complete contract
    │   │   │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Contract compliance tracking
    │   │   │   │   │   │   │   │   ├── compliance-check/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Compliance check
    │   │   │   │   │   │   │   │   ├── deliverables/
    │   │   │   │   │   │   │   │   │   ├── [deliverableId]/
    │   │   │   │   │   │   │   │   │   │   ├── accept/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Accept deliverable
    │   │   │   │   │   │   │   │   │   │   ├── reject/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Reject deliverable
    │   │   │   │   │   │   │   │   │   │   ├── versions/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # File versions (download/preview)
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Deliverable detail
    │   │   │   │   │   │   │   │   │   ├── submit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Submit deliverable
    │   │   │   │   │   │   │   │   │   └── page.tsx  # All deliverables
    │   │   │   │   │   │   │   │   ├── details/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Full contract details
    │   │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   │   │   │   ├── escalate/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Escalate to admin/mediation
    │   │   │   │   │   │   │   │   │   │   ├── respond/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Respond to dispute
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute detail
    │   │   │   │   │   │   │   │   │   ├── open/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Open a dispute
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Disputes list
    │   │   │   │   │   │   │   │   ├── documents/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Contract documents
    │   │   │   │   │   │   │   │   ├── extensions/
    │   │   │   │   │   │   │   │   │   ├── [extensionId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Extension request detail
    │   │   │   │   │   │   │   │   │   ├── request/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Request extension
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Extensions list
    │   │   │   │   │   │   │   │   ├── feedback/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Ongoing feedback
    │   │   │   │   │   │   │   │   ├── gantt/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Gantt chart view
    │   │   │   │   │   │   │   │   ├── invoices/
    │   │   │   │   │   │   │   │   │   ├── [invoiceId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Invoice detail
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create invoice
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Invoices list
    │   │   │   │   │   │   │   │   ├── messages/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Contract-specific messages
    │   │   │   │   │   │   │   │   ├── milestones/
    │   │   │   │   │   │   │   │   │   ├── [milestoneId]/
    │   │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve milestone (client)
    │   │   │   │   │   │   │   │   │   │   ├── dispute/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute milestone
    │   │   │   │   │   │   │   │   │   │   └── submit/
    │   │   │   │   │   │   │   │   │   │       └── page.tsx  # Submit deliverables (freelancer)
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create milestone (if contract allows)
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Milestones list
    │   │   │   │   │   │   │   │   ├── pause/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pause contract
    │   │   │   │   │   │   │   │   ├── payments/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Contract payments
    │   │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Contract reports
    │   │   │   │   │   │   │   │   ├── resume/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Resume paused contract
    │   │   │   │   │   │   │   │   ├── risk-assessment/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Risk assessment
    │   │   │   │   │   │   │   │   ├── signatures/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Signature tracking
    │   │   │   │   │   │   │   │   ├── sow/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit SOW (before signing)
    │   │   │   │   │   │   │   │   │   └── page.tsx  # SOW detail view
    │   │   │   │   │   │   │   │   ├── sow-versions/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # SOW versioning
    │   │   │   │   │   │   │   │   ├── templates/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Contract templates
    │   │   │   │   │   │   │   │   ├── terminate/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Terminate contract
    │   │   │   │   │   │   │   │   ├── timesheet/
    │   │   │   │   │   │   │   │   │   ├── [timesheetId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Timesheet detail
    │   │   │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve timesheet (client)
    │   │   │   │   │   │   │   │   │   ├── submit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Submit timesheet (freelancer)
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Timesheet view (hourly contracts)
    │   │   │   │   │   │   │   │   ├── work-diary/
    │   │   │   │   │   │   │   │   │   ├── [date]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Work diary for specific date
    │   │   │   │   │   │   │   │   │   ├── add-entry/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Add work diary entry (freelancer)
    │   │   │   │   │   │   │   │   │   ├── screenshots/
    │   │   │   │   │   │   │   │   │   │   └── [entryId]/
    │   │   │   │   │   │   │   │   │   │       └── page.tsx  # Screenshot detail (blur/delete if allowed)
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Work diary overview
    │   │   │   │   │   │   │   │   ├── workdiary/
    │   │   │   │   │   │   │   │   │   ├── manual-time/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Add manual time entry
    │   │   │   │   │   │   │   │   │   └── reports/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Work diary reports
    │   │   │   │   │   │   │   │   └── page.tsx  # Contract overview
    │   │   │   │   │   │   │   ├── active/
    │   │   │   │   │   │   │   │   └── page.tsx  # Active contracts only
    │   │   │   │   │   │   │   ├── archived/
    │   │   │   │   │   │   │   │   └── page.tsx  # Archived contracts
    │   │   │   │   │   │   │   ├── benchmarking/
    │   │   │   │   │   │   │   │   └── page.tsx  # Contract performance benchmarking
    │   │   │   │   │   │   │   ├── bulk-actions/
    │   │   │   │   │   │   │   │   └── page.tsx  # Bulk operations
    │   │   │   │   │   │   │   ├── calendar/
    │   │   │   │   │   │   │   │   └── page.tsx  # Contracts calendar view
    │   │   │   │   │   │   │   ├── completed/
    │   │   │   │   │   │   │   │   └── page.tsx  # Completed contracts
    │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   └── [disputeId]/
    │   │   │   │   │   │   │   │       └── page.tsx  # Dispute detail & submit evidence
    │   │   │   │   │   │   │   ├── paused/
    │   │   │   │   │   │   │   │   └── page.tsx  # Paused contracts list
    │   │   │   │   │   │   │   ├── pipeline/
    │   │   │   │   │   │   │   │   └── page.tsx  # Contract pipeline
    │   │   │   │   │   │   │   ├── recurring/
    │   │   │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   │   │   └── renew/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Renew recurring contract
    │   │   │   │   │   │   │   │   └── page.tsx  # Recurring contracts
    │   │   │   │   │   │   │   ├── templates/
    │   │   │   │   │   │   │   │   ├── [templateId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit template
    │   │   │   │   │   │   │   │   │   └── use/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Use template for new contract
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create contract template
    │   │   │   │   │   │   │   │   └── page.tsx  # Contract templates (for recurring work)
    │   │   │   │   │   │   │   ├── error.tsx  # ErrorBoundary
    │   │   │   │   │   │   │   └── page.tsx  # Contracts list
    │   │   │   │   │   │   ├── dashboard/  # Dashboard home (role-based view)
    │   │   │   │   │   │   │   └── page.tsx  # Main dashboard
    │   │   │   │   │   │   ├── deliverables/
    │   │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   │   ├── [deliverableId]/
    │   │   │   │   │   │   │   │   │   ├── review/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Review deliverable (client)
    │   │   │   │   │   │   │   │   │   ├── revisions/
    │   │   │   │   │   │   │   │   │   │   ├── [revisionId]/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Revision detail
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Revision history
    │   │   │   │   │   │   │   │   │   ├── upload/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Upload new version
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Deliverable details
    │   │   │   │   │   │   │   │   ├── new/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Submit new deliverable
    │   │   │   │   │   │   │   │   └── page.tsx  # Contract deliverables list
    │   │   │   │   │   │   │   ├── pending-review/
    │   │   │   │   │   │   │   │   └── page.tsx  # Deliverables pending client review
    │   │   │   │   │   │   │   └── page.tsx  # All deliverables overview
    │   │   │   │   │   │   ├── discover/
    │   │   │   │   │   │   │   ├── advanced-search/
    │   │   │   │   │   │   │   │   └── page.tsx  # Advanced search
    │   │   │   │   │   │   │   ├── ai-matching/
    │   │   │   │   │   │   │   │   └── page.tsx  # AI-powered matching
    │   │   │   │   │   │   │   ├── bulk-actions/
    │   │   │   │   │   │   │   │   └── page.tsx  # Bulk operations
    │   │   │   │   │   │   │   ├── collections/
    │   │   │   │   │   │   │   │   ├── [collectionId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Collection detail
    │   │   │   │   │   │   │   │   └── page.tsx  # Saved collections
    │   │   │   │   │   │   │   ├── filters/
    │   │   │   │   │   │   │   │   ├── save/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Save filter
    │   │   │   │   │   │   │   │   └── page.tsx  # Saved filters
    │   │   │   │   │   │   │   ├── sourcing/
    │   │   │   │   │   │   │   │   ├── campaigns/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Sourcing campaigns
    │   │   │   │   │   │   │   │   ├── pipeline/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Talent pipeline
    │   │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │   │       └── page.tsx  # Outreach templates
    │   │   │   │   │   │   │   └── trending/
    │   │   │   │   │   │   │       └── page.tsx  # Trending talent/jobs
    │   │   │   │   │   │   ├── feed/
    │   │   │   │   │   │   │   └── page.tsx  # Personalized feed (jobs/talent/portfolio)
    │   │   │   │   │   │   ├── financial/
    │   │   │   │   │   │   │   ├── accounting/
    │   │   │   │   │   │   │   │   ├── chart-of-accounts/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Chart of accounts
    │   │   │   │   │   │   │   │   ├── general-ledger/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # General ledger
    │   │   │   │   │   │   │   │   ├── journal-entries/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Journal entries
    │   │   │   │   │   │   │   │   └── reconciliation/
    │   │   │   │   │   │   │   │       └── page.tsx  # Account reconciliation
    │   │   │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Billing history
    │   │   │   │   │   │   │   │   ├── payment-methods/
    │   │   │   │   │   │   │   │   │   ├── [methodId]/
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit payment method
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Payment method detail
    │   │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Add payment method
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Payment methods list
    │   │   │   │   │   │   │   │   └── subscriptions/
    │   │   │   │   │   │   │   │       └── page.tsx  # Subscription billing
    │   │   │   │   │   │   │   ├── budgets/
    │   │   │   │   │   │   │   │   ├── [budgetId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Budget detail
    │   │   │   │   │   │   │   │   ├── forecasting/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Budget forecasting
    │   │   │   │   │   │   │   │   ├── planning/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Budget planning
    │   │   │   │   │   │   │   │   └── variance/
    │   │   │   │   │   │   │   │       └── page.tsx  # Variance analysis
    │   │   │   │   │   │   │   ├── cash-flow/
    │   │   │   │   │   │   │   │   ├── forecasting/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Cash flow forecast
    │   │   │   │   │   │   │   │   └── statements/
    │   │   │   │   │   │   │   │       └── page.tsx  # Cash flow statements
    │   │   │   │   │   │   │   ├── chargebacks/
    │   │   │   │   │   │   │   │   ├── [chargebackId]/
    │   │   │   │   │   │   │   │   │   ├── respond/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Respond to chargeback
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Chargeback details
    │   │   │   │   │   │   │   │   └── page.tsx  # Chargebacks list
    │   │   │   │   │   │   │   ├── cost-centers/
    │   │   │   │   │   │   │   │   ├── [centerId]/
    │   │   │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Cost center analytics
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit cost center
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Cost center details
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create cost center
    │   │   │   │   │   │   │   │   └── page.tsx  # Cost centers list
    │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   │   │   ├── evidence/
    │   │   │   │   │   │   │   │   │   │   ├── submit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Submit evidence
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Evidence list
    │   │   │   │   │   │   │   │   │   ├── messages/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute messages
    │   │   │   │   │   │   │   │   │   ├── resolution/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute resolution
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute detail
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create dispute
    │   │   │   │   │   │   │   │   └── page.tsx  # Disputes list
    │   │   │   │   │   │   │   ├── escrow/
    │   │   │   │   │   │   │   │   ├── [escrowId]/
    │   │   │   │   │   │   │   │   │   ├── fund/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Fund escrow
    │   │   │   │   │   │   │   │   │   ├── release/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Release escrow funds
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Escrow detail
    │   │   │   │   │   │   │   │   └── page.tsx  # Escrow accounts list
    │   │   │   │   │   │   │   ├── expenses/
    │   │   │   │   │   │   │   │   ├── approvals/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Expense approvals
    │   │   │   │   │   │   │   │   ├── categories/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Expense categories
    │   │   │   │   │   │   │   │   ├── policies/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Expense policies
    │   │   │   │   │   │   │   │   └── reimbursements/
    │   │   │   │   │   │   │   │       └── page.tsx  # Reimbursements
    │   │   │   │   │   │   │   ├── forecasting/
    │   │   │   │   │   │   │   │   └── page.tsx  # Revenue/expense forecasting
    │   │   │   │   │   │   │   ├── invoices/
    │   │   │   │   │   │   │   │   ├── [invoiceId]/
    │   │   │   │   │   │   │   │   │   ├── download/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Download invoice PDF
    │   │   │   │   │   │   │   │   │   ├── pay/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Pay invoice
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Invoice detail (already in combined, but completion)
    │   │   │   │   │   │   │   │   └── overdue/
    │   │   │   │   │   │   │   │       └── page.tsx  # Overdue invoices
    │   │   │   │   │   │   │   ├── invoicing/
    │   │   │   │   │   │   │   │   ├── bulk-invoice/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Bulk invoicing
    │   │   │   │   │   │   │   │   ├── recurring/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Recurring invoices
    │   │   │   │   │   │   │   │   ├── reminders/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Payment reminders
    │   │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │   │       └── page.tsx  # Invoice templates
    │   │   │   │   │   │   │   ├── payment-methods/
    │   │   │   │   │   │   │   │   ├── [methodId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit payment method
    │   │   │   │   │   │   │   │   │   ├── remove/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Remove payment method
    │   │   │   │   │   │   │   │   │   └── verify/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Verify payment method
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add payment method
    │   │   │   │   │   │   │   │   ├── bank-accounts/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Bank accounts
    │   │   │   │   │   │   │   │   ├── credit-cards/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Credit cards
    │   │   │   │   │   │   │   │   ├── digital-wallets/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Digital wallets
    │   │   │   │   │   │   │   │   ├── verification/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Payment verification
    │   │   │   │   │   │   │   │   └── page.tsx  # Payment methods list
    │   │   │   │   │   │   │   ├── payout-methods/
    │   │   │   │   │   │   │   │   ├── [methodId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit payout method
    │   │   │   │   │   │   │   │   │   ├── remove/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Remove payout method
    │   │   │   │   │   │   │   │   │   └── verify/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Verify payout method
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add payout method
    │   │   │   │   │   │   │   │   └── page.tsx  # Payout methods list
    │   │   │   │   │   │   │   ├── payouts/
    │   │   │   │   │   │   │   │   ├── [payoutId]/
    │   │   │   │   │   │   │   │   │   ├── details/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Payout transaction detail
    │   │   │   │   │   │   │   │   │   └── receipt/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Payout receipt
    │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending payouts
    │   │   │   │   │   │   │   │   └── schedule/
    │   │   │   │   │   │   │   │       └── page.tsx  # Schedule payout
    │   │   │   │   │   │   │   ├── reconciliation/
    │   │   │   │   │   │   │   │   └── page.tsx  # Financial reconciliation
    │   │   │   │   │   │   │   ├── refunds/
    │   │   │   │   │   │   │   │   ├── [refundId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Refund detail
    │   │   │   │   │   │   │   │   ├── request/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Request refund
    │   │   │   │   │   │   │   │   └── page.tsx  # Refunds list
    │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   ├── earnings/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Earnings reports
    │   │   │   │   │   │   │   │   ├── expenses/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Expense reports
    │   │   │   │   │   │   │   │   ├── tax/
    │   │   │   │   │   │   │   │   │   ├── 1099/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # 1099 tax forms
    │   │   │   │   │   │   │   │   │   ├── vat/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # VAT reports
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Tax reports overview
    │   │   │   │   │   │   │   │   └── page.tsx  # Financial reports overview
    │   │   │   │   │   │   │   ├── tax/
    │   │   │   │   │   │   │   │   ├── 1099-reporting/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # 1099 reporting
    │   │   │   │   │   │   │   │   ├── forms/
    │   │   │   │   │   │   │   │   │   ├── w9/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # W-9 form management
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Tax forms list
    │   │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Tax settings
    │   │   │   │   │   │   │   │   ├── withholding/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Tax withholding
    │   │   │   │   │   │   │   │   └── page.tsx  # Tax overview
    │   │   │   │   │   │   │   ├── treasury/
    │   │   │   │   │   │   │   │   ├── investment/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Treasury investment
    │   │   │   │   │   │   │   │   └── liquidity/
    │   │   │   │   │   │   │   │       └── page.tsx  # Liquidity management
    │   │   │   │   │   │   │   └── wallet/
    │   │   │   │   │   │   │       ├── add-funds/
    │   │   │   │   │   │   │       │   └── page.tsx  # Add funds to wallet
    │   │   │   │   │   │   │       └── withdraw/
    │   │   │   │   │   │   │           └── page.tsx  # Withdraw from wallet
    │   │   │   │   │   │   ├── financials/  # Financial management
    │   │   │   │   │   │   │   ├── escrow/
    │   │   │   │   │   │   │   │   ├── [escrowId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Escrow detail
    │   │   │   │   │   │   │   │   └── page.tsx  # Escrow overview
    │   │   │   │   │   │   │   ├── invoices/
    │   │   │   │   │   │   │   │   ├── [invoiceId]/
    │   │   │   │   │   │   │   │   │   ├── pay/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Pay invoice (client)
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Invoice detail
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create invoice (manual invoicing)
    │   │   │   │   │   │   │   │   └── page.tsx  # Invoices list
    │   │   │   │   │   │   │   ├── payment-methods/
    │   │   │   │   │   │   │   │   ├── [methodId]/
    │   │   │   │   │   │   │   │   │   ├── delete/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Delete payment method
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Payment method detail
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add payment method
    │   │   │   │   │   │   │   │   └── page.tsx  # Payment methods list
    │   │   │   │   │   │   │   ├── payout-methods/
    │   │   │   │   │   │   │   │   ├── [methodId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Payout method detail
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add payout method
    │   │   │   │   │   │   │   │   └── page.tsx  # Payout methods list (freelancer)
    │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   ├── earnings/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Detailed earnings report
    │   │   │   │   │   │   │   │   ├── spending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Detailed spending report
    │   │   │   │   │   │   │   │   └── page.tsx  # Financial reports
    │   │   │   │   │   │   │   ├── tax/
    │   │   │   │   │   │   │   │   ├── forms/
    │   │   │   │   │   │   │   │   │   ├── upload/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Upload tax form
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Tax forms list
    │   │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Tax settings
    │   │   │   │   │   │   │   │   └── page.tsx  # Tax information
    │   │   │   │   │   │   │   ├── transactions/
    │   │   │   │   │   │   │   │   ├── [transactionId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Transaction detail
    │   │   │   │   │   │   │   │   └── page.tsx  # Transaction history
    │   │   │   │   │   │   │   ├── wallet/
    │   │   │   │   │   │   │   │   ├── add-funds/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add funds (client)
    │   │   │   │   │   │   │   │   ├── withdraw/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Withdraw funds (freelancer)
    │   │   │   │   │   │   │   │   └── page.tsx  # Wallet details
    │   │   │   │   │   │   │   └── page.tsx  # Financial overview
    │   │   │   │   │   │   ├── insights/
    │   │   │   │   │   │   │   ├── market-trends/
    │   │   │   │   │   │   │   │   └── page.tsx  # Market trends
    │   │   │   │   │   │   │   ├── predictions/
    │   │   │   │   │   │   │   │   └── page.tsx  # AI predictions
    │   │   │   │   │   │   │   └── recommendations/
    │   │   │   │   │   │   │       └── page.tsx  # AI recommendations
    │   │   │   │   │   │   ├── invitations/
    │   │   │   │   │   │   │   ├── received/
    │   │   │   │   │   │   │   │   ├── [inviteId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Invitation details
    │   │   │   │   │   │   │   │   └── page.tsx  # Received invitations list
    │   │   │   │   │   │   │   ├── sent/
    │   │   │   │   │   │   │   │   ├── [inviteId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Sent invitation tracking
    │   │   │   │   │   │   │   │   └── page.tsx  # Sent invitations list (client)
    │   │   │   │   │   │   │   └── page.tsx  # Invitations overview
    │   │   │   │   │   │   ├── job-alerts/
    │   │   │   │   │   │   │   ├── [alertId]/
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit job alert
    │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Alert history
    │   │   │   │   │   │   │   │   └── page.tsx  # Alert detail
    │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   └── page.tsx  # Create job alert
    │   │   │   │   │   │   │   └── page.tsx  # Job alerts list
    │   │   │   │   │   │   ├── jobs/  # Jobs management
    │   │   │   │   │   │   │   ├── [jobId]/
    │   │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Job analytics (client)
    │   │   │   │   │   │   │   │   ├── bidding/
    │   │   │   │   │   │   │   │   │   ├── place-bid/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Place/update bid (freelancer)
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Active bids on job (client view)
    │   │   │   │   │   │   │   │   ├── close/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Close job
    │   │   │   │   │   │   │   │   ├── duplicate/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Duplicate job form
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit job (client only)
    │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Job edit history
    │   │   │   │   │   │   │   │   ├── invite/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Invite freelancers (client)
    │   │   │   │   │   │   │   │   ├── proposals/
    │   │   │   │   │   │   │   │   │   ├── [proposalId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Proposal detail
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Proposals received (client)
    │   │   │   │   │   │   │   │   ├── repost/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Repost expired job
    │   │   │   │   │   │   │   │   └── page.tsx  # Job detail page
    │   │   │   │   │   │   │   ├── archived/
    │   │   │   │   │   │   │   │   └── page.tsx  # Archived jobs list
    │   │   │   │   │   │   │   ├── batch/
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Bulk job creation
    │   │   │   │   │   │   │   │   └── manage/
    │   │   │   │   │   │   │   │       └── page.tsx  # Batch operations
    │   │   │   │   │   │   │   ├── browse/
    │   │   │   │   │   │   │   │   └── page.tsx  # Job listings with filters
    │   │   │   │   │   │   │   ├── categories/
    │   │   │   │   │   │   │   │   ├── [categoryId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Jobs in category
    │   │   │   │   │   │   │   │   └── page.tsx  # Browse by category
    │   │   │   │   │   │   │   ├── drafts/
    │   │   │   │   │   │   │   │   ├── [draftId]/
    │   │   │   │   │   │   │   │   │   └── edit/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Edit draft
    │   │   │   │   │   │   │   │   └── page.tsx  # Job drafts list
    │   │   │   │   │   │   │   ├── insights/
    │   │   │   │   │   │   │   │   └── page.tsx  # Market insights for job posting
    │   │   │   │   │   │   │   ├── invitations/
    │   │   │   │   │   │   │   │   └── page.tsx  # Job invitations received
    │   │   │   │   │   │   │   ├── my-jobs/
    │   │   │   │   │   │   │   │   ├── active/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Active jobs only
    │   │   │   │   │   │   │   │   ├── closed/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Closed jobs
    │   │   │   │   │   │   │   │   └── page.tsx  # All posted jobs
    │   │   │   │   │   │   │   ├── post/
    │   │   │   │   │   │   │   │   └── page.tsx  # Post a new job (client)
    │   │   │   │   │   │   │   ├── recommendations/
    │   │   │   │   │   │   │   │   └── page.tsx  # Recommended jobs (freelancer)
    │   │   │   │   │   │   │   ├── saved/
    │   │   │   │   │   │   │   │   └── page.tsx  # Saved/bookmarked jobs
    │   │   │   │   │   │   │   ├── saved-searches/
    │   │   │   │   │   │   │   │   └── page.tsx  # Saved job searches
    │   │   │   │   │   │   │   ├── scheduling/
    │   │   │   │   │   │   │   │   └── page.tsx  # Schedule job postings
    │   │   │   │   │   │   │   ├── templates/
    │   │   │   │   │   │   │   │   ├── [templateId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit job template
    │   │   │   │   │   │   │   │   │   └── use/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Use template to create job
    │   │   │   │   │   │   │   │   └── page.tsx  # Job templates list
    │   │   │   │   │   │   │   ├── error.tsx  # ErrorBoundary
    │   │   │   │   │   │   │   └── page.tsx  # Jobs list (role-based)
    │   │   │   │   │   │   ├── learning/
    │   │   │   │   │   │   │   ├── achievements/
    │   │   │   │   │   │   │   │   ├── badges/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Badges earned
    │   │   │   │   │   │   │   │   ├── certifications/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Certifications
    │   │   │   │   │   │   │   │   ├── leaderboard/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Leaderboard
    │   │   │   │   │   │   │   │   └── page.tsx  # Achievements & badges
    │   │   │   │   │   │   │   ├── assessments/
    │   │   │   │   │   │   │   │   ├── [assessmentId]/
    │   │   │   │   │   │   │   │   │   ├── results/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Assessment results
    │   │   │   │   │   │   │   │   │   ├── review/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Review answers
    │   │   │   │   │   │   │   │   │   ├── take/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Take assessment
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Assessment detail
    │   │   │   │   │   │   │   │   └── page.tsx  # Assessments list
    │   │   │   │   │   │   │   ├── certifications/
    │   │   │   │   │   │   │   │   ├── [certId]/
    │   │   │   │   │   │   │   │   │   ├── verify/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Verify certificate
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Certificate detail
    │   │   │   │   │   │   │   │   └── page.tsx  # Manage certifications
    │   │   │   │   │   │   │   ├── courses/
    │   │   │   │   │   │   │   │   ├── [courseId]/
    │   │   │   │   │   │   │   │   │   ├── assessments/
    │   │   │   │   │   │   │   │   │   │   └── [assessmentId]/
    │   │   │   │   │   │   │   │   │   │       ├── results/
    │   │   │   │   │   │   │   │   │   │       │   └── page.tsx  # View results
    │   │   │   │   │   │   │   │   │   │       ├── start/
    │   │   │   │   │   │   │   │   │   │       │   └── page.tsx  # Start assessment
    │   │   │   │   │   │   │   │   │   │       └── submit/
    │   │   │   │   │   │   │   │   │   │           └── page.tsx  # Submit answers
    │   │   │   │   │   │   │   │   │   ├── curriculum/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Course curriculum
    │   │   │   │   │   │   │   │   │   ├── discussions/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Course discussions
    │   │   │   │   │   │   │   │   │   ├── lessons/
    │   │   │   │   │   │   │   │   │   │   └── [lessonId]/
    │   │   │   │   │   │   │   │   │   │       └── page.tsx  # Lesson content
    │   │   │   │   │   │   │   │   │   ├── materials/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Course materials
    │   │   │   │   │   │   │   │   │   ├── progress/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Course progress
    │   │   │   │   │   │   │   │   │   ├── quizzes/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Course quizzes
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Course detail
    │   │   │   │   │   │   │   │   ├── browse/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Browse all courses
    │   │   │   │   │   │   │   │   ├── catalog/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Course catalog
    │   │   │   │   │   │   │   │   ├── my-courses/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Enrolled courses
    │   │   │   │   │   │   │   │   ├── recommended/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Recommended courses
    │   │   │   │   │   │   │   │   └── page.tsx  # Courses catalog
    │   │   │   │   │   │   │   ├── mentorship/
    │   │   │   │   │   │   │   │   ├── [sessionId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Mentorship session details
    │   │   │   │   │   │   │   │   ├── find-mentor/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Find a mentor
    │   │   │   │   │   │   │   │   ├── my-mentees/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Manage mentees
    │   │   │   │   │   │   │   │   └── page.tsx  # Mentorship dashboard
    │   │   │   │   │   │   │   ├── paths/
    │   │   │   │   │   │   │   │   ├── [pathId]/
    │   │   │   │   │   │   │   │   │   ├── progress/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Learning path progress
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Learning path details
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create learning path
    │   │   │   │   │   │   │   │   ├── discover/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Discover learning paths
    │   │   │   │   │   │   │   │   └── page.tsx  # My learning paths
    │   │   │   │   │   │   │   ├── skill-tests/
    │   │   │   │   │   │   │   │   ├── [testId]/
    │   │   │   │   │   │   │   │   │   ├── instructions/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Test instructions
    │   │   │   │   │   │   │   │   │   ├── results/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Test results
    │   │   │   │   │   │   │   │   │   └── take-test/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Take the test
    │   │   │   │   │   │   │   │   └── page.tsx  # Available skill tests
    │   │   │   │   │   │   │   └── skills/
    │   │   │   │   │   │   │       ├── assessments/
    │   │   │   │   │   │   │       │   └── page.tsx  # Skill assessments
    │   │   │   │   │   │   │       ├── gaps/
    │   │   │   │   │   │   │       │   └── page.tsx  # Skill gaps
    │   │   │   │   │   │   │       └── recommendations/
    │   │   │   │   │   │   │           └── page.tsx  # Skill recommendations
    │   │   │   │   │   │   ├── messages/  # Messaging
    │   │   │   │   │   │   │   ├── [conversationId]/
    │   │   │   │   │   │   │   │   ├── archive/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Archive conversation
    │   │   │   │   │   │   │   │   ├── attachments/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Attachment manager
    │   │   │   │   │   │   │   │   ├── info/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Conversation info
    │   │   │   │   │   │   │   │   ├── search/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Search messages
    │   │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Conversation settings
    │   │   │   │   │   │   │   │   └── page.tsx  # Conversation thread
    │   │   │   │   │   │   │   ├── archived/
    │   │   │   │   │   │   │   │   └── page.tsx  # Archived conversations
    │   │   │   │   │   │   │   ├── compose/
    │   │   │   │   │   │   │   │   └── page.tsx  # Compose new message
    │   │   │   │   │   │   │   ├── drafts/
    │   │   │   │   │   │   │   │   └── page.tsx  # Message drafts
    │   │   │   │   │   │   │   ├── new/
    │   │   │   │   │   │   │   │   └── page.tsx  # Start new conversation
    │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   └── notifications/
    │   │   │   │   │   │   │   │       └── page.tsx  # Notification preferences
    │   │   │   │   │   │   │   ├── starred/
    │   │   │   │   │   │   │   │   └── page.tsx  # Starred messages
    │   │   │   │   │   │   │   ├── templates/
    │   │   │   │   │   │   │   │   └── page.tsx  # Message templates
    │   │   │   │   │   │   │   └── page.tsx  # Inbox
    │   │   │   │   │   │   ├── mobile-preview/
    │   │   │   │   │   │   │   └── page.tsx  # Preview mobile experience
    │   │   │   │   │   │   ├── network/
    │   │   │   │   │   │   │   ├── connections/
    │   │   │   │   │   │   │   │   ├── [userId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Connection profile view
    │   │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pending connection requests
    │   │   │   │   │   │   │   │   └── page.tsx  # Connections list
    │   │   │   │   │   │   │   ├── groups/
    │   │   │   │   │   │   │   │   ├── [groupId]/
    │   │   │   │   │   │   │   │   │   ├── members/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Group members
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Group details
    │   │   │   │   │   │   │   │   ├── discover/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Discover groups
    │   │   │   │   │   │   │   │   └── page.tsx  # My groups
    │   │   │   │   │   │   │   ├── recommendations/
    │   │   │   │   │   │   │   │   └── page.tsx  # Connection recommendations
    │   │   │   │   │   │   │   └── referrals/
    │   │   │   │   │   │   │       ├── dashboard/
    │   │   │   │   │   │   │       │   └── page.tsx  # Referral dashboard
    │   │   │   │   │   │   │       └── page.tsx  # Referrals overview
    │   │   │   │   │   │   ├── notifications/  # Notifications center
    │   │   │   │   │   │   │   ├── [notificationId]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Notification detail (redirects to relevant page)
    │   │   │   │   │   │   │   ├── all/
    │   │   │   │   │   │   │   │   └── page.tsx  # All notifications
    │   │   │   │   │   │   │   ├── push-settings/
    │   │   │   │   │   │   │   │   └── page.tsx  # Configure push notifications
    │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   ├── channels/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Notification channels
    │   │   │   │   │   │   │   │   └── page.tsx  # Notification settings
    │   │   │   │   │   │   │   ├── unread/
    │   │   │   │   │   │   │   │   └── page.tsx  # Unread notifications only
    │   │   │   │   │   │   │   └── page.tsx  # Notifications center
    │   │   │   │   │   │   ├── offline/  # ENTIRE SECTION
    │   │   │   │   │   │   │   ├── queue/
    │   │   │   │   │   │   │   │   └── page.tsx  # Pending actions queue
    │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   └── page.tsx  # Offline mode configuration
    │   │   │   │   │   │   │   ├── sync/
    │   │   │   │   │   │   │   │   └── page.tsx  # Sync management
    │   │   │   │   │   │   │   ├── layout.tsx  # Offline mode layout
    │   │   │   │   │   │   │   └── page.tsx  # Offline status overview
    │   │   │   │   │   │   ├── organization/  # Organization management (for clients)
    │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   └── page.tsx  # Organization analytics
    │   │   │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   │   │   ├── contracts/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Billing contracts
    │   │   │   │   │   │   │   │   ├── cost-allocation/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Cost allocation
    │   │   │   │   │   │   │   │   ├── spending-limits/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Spending limits
    │   │   │   │   │   │   │   │   └── usage/
    │   │   │   │   │   │   │   │       └── page.tsx  # Usage reports
    │   │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   │   ├── certifications/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Certifications
    │   │   │   │   │   │   │   │   ├── documents/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Compliance docs
    │   │   │   │   │   │   │   │   └── policies/
    │   │   │   │   │   │   │   │       └── page.tsx  # Company policies
    │   │   │   │   │   │   │   ├── hierarchy/
    │   │   │   │   │   │   │   │   ├── chart/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Org chart
    │   │   │   │   │   │   │   │   ├── departments/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Departments
    │   │   │   │   │   │   │   │   └── teams/
    │   │   │   │   │   │   │   │       └── page.tsx  # Teams structure
    │   │   │   │   │   │   │   ├── integrations/
    │   │   │   │   │   │   │   │   ├── [integrationId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Integration detail
    │   │   │   │   │   │   │   │   ├── available/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Available integrations
    │   │   │   │   │   │   │   │   └── installed/
    │   │   │   │   │   │   │   │       └── page.tsx  # Installed integrations
    │   │   │   │   │   │   │   ├── permissions/
    │   │   │   │   │   │   │   │   ├── audit/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Permission audit
    │   │   │   │   │   │   │   │   ├── matrix/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Permissions matrix
    │   │   │   │   │   │   │   │   └── roles/
    │   │   │   │   │   │   │   │       └── page.tsx  # Role management
    │   │   │   │   │   │   │   ├── procurement/
    │   │   │   │   │   │   │   │   ├── approvals/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Approval workflows
    │   │   │   │   │   │   │   │   ├── budgets/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Budget management
    │   │   │   │   │   │   │   │   ├── purchase-orders/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Purchase orders
    │   │   │   │   │   │   │   │   └── vendors/
    │   │   │   │   │   │   │   │       └── page.tsx  # Vendor management
    │   │   │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Organization billing
    │   │   │   │   │   │   │   │   └── page.tsx  # Organization settings
    │   │   │   │   │   │   │   ├── spending/
    │   │   │   │   │   │   │   │   ├── budgets/
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create budget
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Budget management
    │   │   │   │   │   │   │   │   └── page.tsx  # Spending overview
    │   │   │   │   │   │   │   ├── sso/
    │   │   │   │   │   │   │   │   ├── configuration/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # SSO config
    │   │   │   │   │   │   │   │   ├── mapping/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Attribute mapping
    │   │   │   │   │   │   │   │   └── providers/
    │   │   │   │   │   │   │   │       └── page.tsx  # Identity providers
    │   │   │   │   │   │   │   ├── team/
    │   │   │   │   │   │   │   │   ├── [memberId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit member
    │   │   │   │   │   │   │   │   │   └── remove/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Remove member
    │   │   │   │   │   │   │   │   ├── invite/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Invite team member
    │   │   │   │   │   │   │   │   ├── roles/
    │   │   │   │   │   │   │   │   │   ├── [roleId]/
    │   │   │   │   │   │   │   │   │   │   └── edit/
    │   │   │   │   │   │   │   │   │   │       └── page.tsx  # Edit role
    │   │   │   │   │   │   │   │   │   └── create/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Create custom role
    │   │   │   │   │   │   │   │   └── page.tsx  # Team members list
    │   │   │   │   │   │   │   └── page.tsx  # Organization overview
    │   │   │   │   │   │   ├── orgs/
    │   │   │   │   │   │   │   └── [orgId]/
    │   │   │   │   │   │   │       └── settings/
    │   │   │   │   │   │   │           └── billing/
    │   │   │   │   │   │   │               └── tax-vat/
    │   │   │   │   │   │   │                   └── page.tsx  # Org tax/VAT profile
    │   │   │   │   │   │   ├── portfolio/
    │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   └── page.tsx  # Portfolio analytics
    │   │   │   │   │   │   │   ├── import/
    │   │   │   │   │   │   │   │   └── page.tsx  # Import portfolio
    │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │       └── page.tsx  # Portfolio templates
    │   │   │   │   │   │   ├── profile/  # Current user profile management
    │   │   │   │   │   │   │   ├── availability/
    │   │   │   │   │   │   │   │   └── page.tsx  # Availability management
    │   │   │   │   │   │   │   ├── certifications/
    │   │   │   │   │   │   │   │   ├── [certId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit certification
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Certification detail
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add certification
    │   │   │   │   │   │   │   │   ├── verify/
    │   │   │   │   │   │   │   │   │   ├── [certificationId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Verification request
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Verify certifications
    │   │   │   │   │   │   │   │   └── page.tsx  # Certifications list
    │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   └── page.tsx  # Edit profile form
    │   │   │   │   │   │   │   ├── education/
    │   │   │   │   │   │   │   │   ├── [educationId]/
    │   │   │   │   │   │   │   │   │   └── edit/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Edit education form
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add education form
    │   │   │   │   │   │   │   │   └── page.tsx  # Education list
    │   │   │   │   │   │   │   ├── experience/
    │   │   │   │   │   │   │   │   ├── [experienceId]/
    │   │   │   │   │   │   │   │   │   └── edit/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Edit experience form
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add experience form
    │   │   │   │   │   │   │   │   └── page.tsx  # Work experience list
    │   │   │   │   │   │   │   ├── languages/
    │   │   │   │   │   │   │   │   └── page.tsx  # Language proficiency
    │   │   │   │   │   │   │   ├── portfolio/
    │   │   │   │   │   │   │   │   ├── [portfolioId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit portfolio item
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Portfolio item detail
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add portfolio item
    │   │   │   │   │   │   │   │   ├── reorder/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Reorder portfolio items
    │   │   │   │   │   │   │   │   └── page.tsx  # Portfolio items list
    │   │   │   │   │   │   │   ├── portfolio-items/
    │   │   │   │   │   │   │   │   └── [itemId]/
    │   │   │   │   │   │   │   │       ├── analytics/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Portfolio item analytics
    │   │   │   │   │   │   │   │       └── page.tsx  # Portfolio item detail
    │   │   │   │   │   │   │   ├── references/
    │   │   │   │   │   │   │   │   ├── [referenceId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Reference detail
    │   │   │   │   │   │   │   │   ├── request/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Request reference
    │   │   │   │   │   │   │   │   └── page.tsx  # References list
    │   │   │   │   │   │   │   ├── reviews/
    │   │   │   │   │   │   │   │   └── page.tsx  # Reviews (from profile)
    │   │   │   │   │   │   │   ├── service-catalog/
    │   │   │   │   │   │   │   │   ├── [serviceId]/
    │   │   │   │   │   │   │   │   │   └── edit/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Edit service
    │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Add service
    │   │   │   │   │   │   │   │   └── page.tsx  # Service catalog management (freelancer)
    │   │   │   │   │   │   │   ├── skills/
    │   │   │   │   │   │   │   │   ├── specializations/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Specializations & niche expertise
    │   │   │   │   │   │   │   │   └── page.tsx  # Skills management
    │   │   │   │   │   │   │   ├── social-links/
    │   │   │   │   │   │   │   │   └── page.tsx  # Social media links
    │   │   │   │   │   │   │   ├── verification/
    │   │   │   │   │   │   │   │   ├── identity/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # ID verification
    │   │   │   │   │   │   │   │   ├── phone/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Phone verification
    │   │   │   │   │   │   │   │   └── page.tsx  # Verification status
    │   │   │   │   │   │   │   ├── visibility/
    │   │   │   │   │   │   │   │   └── page.tsx  # Profile visibility settings
    │   │   │   │   │   │   │   └── page.tsx  # Profile overview / public view
    │   │   │   │   │   │   ├── projects/
    │   │   │   │   │   │   │   ├── [projectId]/
    │   │   │   │   │   │   │   │   ├── board/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Kanban board
    │   │   │   │   │   │   │   │   ├── dependencies/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Task dependencies
    │   │   │   │   │   │   │   │   ├── gantt/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Gantt chart
    │   │   │   │   │   │   │   │   ├── kanban/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Kanban view
    │   │   │   │   │   │   │   │   ├── milestones/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Project milestones
    │   │   │   │   │   │   │   │   ├── resources/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Resource allocation
    │   │   │   │   │   │   │   │   ├── risks/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Risk register
    │   │   │   │   │   │   │   │   ├── timeline/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Project timeline
    │   │   │   │   │   │   │   │   └── workload/
    │   │   │   │   │   │   │   │       └── page.tsx  # Workload view
    │   │   │   │   │   │   │   ├── calendar/
    │   │   │   │   │   │   │   │   └── page.tsx  # Project calendar
    │   │   │   │   │   │   │   ├── portfolio/
    │   │   │   │   │   │   │   │   └── page.tsx  # Project portfolio
    │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │       └── page.tsx  # Project templates
    │   │   │   │   │   │   ├── proposals/  # Proposals management
    │   │   │   │   │   │   │   ├── [proposalId]/
    │   │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Proposal performance analytics
    │   │   │   │   │   │   │   │   ├── bidding/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Bid status for proposal
    │   │   │   │   │   │   │   │   ├── collaborators/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Team collaboration on proposal
    │   │   │   │   │   │   │   │   ├── compare/
    │   │   │   │   │   │   │   │   │   ├── [compareWith]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Compare two proposals side-by-side
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Compare with other proposals
    │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit proposal
    │   │   │   │   │   │   │   │   ├── feedback/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Proposal feedback
    │   │   │   │   │   │   │   │   ├── milestones/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit proposal milestones
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Milestones overview
    │   │   │   │   │   │   │   │   ├── negotiation/
    │   │   │   │   │   │   │   │   │   ├── counter-offer/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create counter-offer
    │   │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Negotiation timeline
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Active negotiation dashboard
    │   │   │   │   │   │   │   │   ├── negotiations/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Proposal negotiations history
    │   │   │   │   │   │   │   │   ├── questions/
    │   │   │   │   │   │   │   │   │   ├── [questionId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Answer single question
    │   │   │   │   │   │   │   │   │   └── page.tsx  # All proposal questions
    │   │   │   │   │   │   │   │   ├── revise/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Revise proposal
    │   │   │   │   │   │   │   │   ├── versions/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Proposal revision history
    │   │   │   │   │   │   │   │   ├── withdraw/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Withdraw proposal
    │   │   │   │   │   │   │   │   └── page.tsx  # Proposal detail
    │   │   │   │   │   │   │   ├── accepted/
    │   │   │   │   │   │   │   │   └── page.tsx  # Accepted proposals
    │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   └── page.tsx  # Proposal analytics (freelancer)
    │   │   │   │   │   │   │   ├── archived/
    │   │   │   │   │   │   │   │   └── page.tsx  # Archived proposals
    │   │   │   │   │   │   │   ├── benchmarking/
    │   │   │   │   │   │   │   │   └── page.tsx  # Benchmark against market
    │   │   │   │   │   │   │   ├── declined/  # (already listed above; kept once)
    │   │   │   │   │   │   │   │   └── page.tsx  # Declined proposals
    │   │   │   │   │   │   │   ├── drafts/
    │   │   │   │   │   │   │   │   ├── [draftId]/
    │   │   │   │   │   │   │   │   │   └── edit/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Edit proposal draft
    │   │   │   │   │   │   │   │   └── page.tsx  # Proposal drafts list
    │   │   │   │   │   │   │   ├── insights/
    │   │   │   │   │   │   │   │   ├── pricing-analysis/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Pricing competitiveness
    │   │   │   │   │   │   │   │   ├── response-time/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Response time analytics
    │   │   │   │   │   │   │   │   ├── win-rate/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Win rate analytics
    │   │   │   │   │   │   │   │   └── page.tsx  # Proposal insights
    │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   └── page.tsx  # Pending proposals
    │   │   │   │   │   │   │   ├── portfolio-showcases/
    │   │   │   │   │   │   │   │   ├── [showcaseId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit showcase
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Showcase detail
    │   │   │   │   │   │   │   │   └── page.tsx  # Manage showcases
    │   │   │   │   │   │   │   ├── rejected/
    │   │   │   │   │   │   │   │   └── page.tsx  # Rejected proposals
    │   │   │   │   │   │   │   ├── submit/
    │   │   │   │   │   │   │   │   └── [jobId]/
    │   │   │   │   │   │   │   │       └── page.tsx  # Submit proposal (freelancer)
    │   │   │   │   │   │   │   ├── templates/
    │   │   │   │   │   │   │   │   ├── [templateId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit template
    │   │   │   │   │   │   │   │   │   └── use/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Use template for new proposal
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create template
    │   │   │   │   │   │   │   │   └── page.tsx  # Proposal templates
    │   │   │   │   │   │   │   ├── withdrawn/
    │   │   │   │   │   │   │   │   └── page.tsx  # Withdrawn proposals
    │   │   │   │   │   │   │   ├── error.tsx  # ErrorBoundary
    │   │   │   │   │   │   │   └── page.tsx  # Proposals list
    │   │   │   │   │   │   ├── quick-actions/
    │   │   │   │   │   │   │   └── page.tsx  # - Quick actions (web equivalent)
    │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   ├── earnings/
    │   │   │   │   │   │   │   │   ├── by-client/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Earnings by client
    │   │   │   │   │   │   │   │   ├── by-project/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Earnings by project
    │   │   │   │   │   │   │   │   ├── by-skill/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Earnings by skill
    │   │   │   │   │   │   │   │   └── forecast/
    │   │   │   │   │   │   │   │       └── page.tsx  # Earnings forecast
    │   │   │   │   │   │   │   ├── financial/
    │   │   │   │   │   │   │   │   ├── expenses/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Expense reports
    │   │   │   │   │   │   │   │   ├── invoices/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Invoice reports
    │   │   │   │   │   │   │   │   ├── profit-loss/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # P&L statement
    │   │   │   │   │   │   │   │   └── tax/
    │   │   │   │   │   │   │   │       └── page.tsx  # Tax reports
    │   │   │   │   │   │   │   ├── performance/
    │   │   │   │   │   │   │   │   ├── delivery/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Delivery metrics
    │   │   │   │   │   │   │   │   ├── quality/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Quality metrics
    │   │   │   │   │   │   │   │   └── velocity/
    │   │   │   │   │   │   │   │       └── page.tsx  # Velocity metrics
    │   │   │   │   │   │   │   └── scheduled/
    │   │   │   │   │   │   │       ├── [scheduleId]/
    │   │   │   │   │   │   │       │   └── page.tsx  # Scheduled report detail
    │   │   │   │   │   │   │       └── page.tsx  # Scheduled reports
    │   │   │   │   │   │   ├── reputation/
    │   │   │   │   │   │   │   ├── badges/
    │   │   │   │   │   │   │   │   └── page.tsx  # Achievement badges
    │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   └── page.tsx  # Review disputes
    │   │   │   │   │   │   │   ├── overview/
    │   │   │   │   │   │   │   │   └── page.tsx  # Reputation overview
    │   │   │   │   │   │   │   ├── reviews-given/
    │   │   │   │   │   │   │   │   └── page.tsx  # Reviews given
    │   │   │   │   │   │   │   └── reviews-received/
    │   │   │   │   │   │   │       └── page.tsx  # Reviews received
    │   │   │   │   │   │   ├── reviews/  # Reviews & ratings
    │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   └── page.tsx  # Review analytics
    │   │   │   │   │   │   │   ├── badges/
    │   │   │   │   │   │   │   │   └── page.tsx  # Badges & achievements
    │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create review
    │   │   │   │   │   │   │   │   └── pending/
    │   │   │   │   │   │   │   │       └── page.tsx  # Pending reviews
    │   │   │   │   │   │   │   ├── dispute/
    │   │   │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Review dispute detail
    │   │   │   │   │   │   │   │   ├── submit/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Dispute a review
    │   │   │   │   │   │   │   │   └── page.tsx  # Review disputes list
    │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   ├── [disputeId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Review dispute details
    │   │   │   │   │   │   │   │   └── page.tsx  # Review disputes list
    │   │   │   │   │   │   │   ├── given/
    │   │   │   │   │   │   │   │   ├── [reviewId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit given review
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Given review details
    │   │   │   │   │   │   │   │   └── page.tsx  # Reviews given list
    │   │   │   │   │   │   │   ├── my/
    │   │   │   │   │   │   │   │   └── page.tsx  # My reviews
    │   │   │   │   │   │   │   ├── pending/
    │   │   │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Leave review form
    │   │   │   │   │   │   │   │   └── page.tsx  # Pending reviews to complete
    │   │   │   │   │   │   │   ├── rate/
    │   │   │   │   │   │   │   │   └── [targetId]/
    │   │   │   │   │   │   │   │       └── page.tsx  # Leave a review
    │   │   │   │   │   │   │   ├── received/
    │   │   │   │   │   │   │   │   ├── [reviewId]/
    │   │   │   │   │   │   │   │   │   ├── respond/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Respond to review
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Review detail
    │   │   │   │   │   │   │   │   └── page.tsx  # Reviews received list
    │   │   │   │   │   │   │   ├── stats/
    │   │   │   │   │   │   │   │   └── page.tsx  # Detailed statistics
    │   │   │   │   │   │   │   └── page.tsx  # Reviews overview
    │   │   │   │   │   │   ├── search/  # Advanced search functionality
    │   │   │   │   │   │   │   ├── advanced/
    │   │   │   │   │   │   │   │   └── page.tsx  # Advanced search interface
    │   │   │   │   │   │   │   ├── alerts/
    │   │   │   │   │   │   │   │   └── page.tsx  # Saved-search alerts (channels/schedule)
    │   │   │   │   │   │   │   ├── freelancers/
    │   │   │   │   │   │   │   │   ├── advanced/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Advanced freelancer search
    │   │   │   │   │   │   │   │   └── page.tsx  # Advanced freelancer search (client)
    │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   └── page.tsx  # Search history
    │   │   │   │   │   │   │   ├── jobs/
    │   │   │   │   │   │   │   │   ├── advanced/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Advanced job search
    │   │   │   │   │   │   │   │   └── page.tsx  # Advanced job search
    │   │   │   │   │   │   │   ├── personalization/
    │   │   │   │   │   │   │   │   └── page.tsx  # Pause/reset, prefer/hide entities
    │   │   │   │   │   │   │   ├── portfolio/
    │   │   │   │   │   │   │   │   └── page.tsx  # Search portfolios
    │   │   │   │   │   │   │   ├── portfolios/
    │   │   │   │   │   │   │   │   └── page.tsx  # Search portfolios
    │   │   │   │   │   │   │   ├── recommendations/
    │   │   │   │   │   │   │   │   └── page.tsx  # Personalized recommendations
    │   │   │   │   │   │   │   ├── saved/
    │   │   │   │   │   │   │   │   ├── [searchId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit saved search
    │   │   │   │   │   │   │   │   │   └── results/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # View results from saved search
    │   │   │   │   │   │   │   │   └── page.tsx  # Saved searches list (may be in combined, ensuring here)
    │   │   │   │   │   │   │   ├── saved-searches/
    │   │   │   │   │   │   │   │   ├── [searchId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit saved search
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Execute saved search
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create saved search
    │   │   │   │   │   │   │   │   └── page.tsx  # Saved searches list
    │   │   │   │   │   │   │   └── trending/
    │   │   │   │   │   │   │       └── page.tsx  # Trending searches and jobs
    │   │   │   │   │   │   ├── settings/  # User settings
    │   │   │   │   │   │   │   ├── accessibility/
    │   │   │   │   │   │   │   │   └── page.tsx  # Accessibility preferences
    │   │   │   │   │   │   │   ├── account/
    │   │   │   │   │   │   │   │   ├── close/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Close account
    │   │   │   │   │   │   │   │   ├── data-export/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Export user data (GDPR)
    │   │   │   │   │   │   │   │   ├── deactivate/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Deactivate account
    │   │   │   │   │   │   │   │   ├── delete/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Delete account
    │   │   │   │   │   │   │   │   ├── email/
    │   │   │   │   │   │   │   │   │   └── change/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Change email
    │   │   │   │   │   │   │   │   ├── phone/
    │   │   │   │   │   │   │   │   │   └── change/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Change phone number
    │   │   │   │   │   │   │   │   ├── reactivate/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Reactivate account
    │   │   │   │   │   │   │   │   └── username/
    │   │   │   │   │   │   │   │       └── change/
    │   │   │   │   │   │   │   │           └── page.tsx  # Change username
    │   │   │   │   │   │   │   ├── advanced/
    │   │   │   │   │   │   │   │   └── page.tsx  # Advanced settings
    │   │   │   │   │   │   │   ├── api/
    │   │   │   │   │   │   │   │   ├── documentation/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # API documentation
    │   │   │   │   │   │   │   │   └── page.tsx  # API settings (developer section)
    │   │   │   │   │   │   │   ├── authorized-apps/
    │   │   │   │   │   │   │   │   ├── [appId]/
    │   │   │   │   │   │   │   │   │   └── revoke/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Revoke app access
    │   │   │   │   │   │   │   │   └── page.tsx  # Authorized apps list
    │   │   │   │   │   │   │   ├── blocked-users/
    │   │   │   │   │   │   │   │   └── page.tsx  # Blocked users management
    │   │   │   │   │   │   │   ├── developer/  # Web developer API pages (verify/create)
    │   │   │   │   │   │   │   │   ├── api-keys/
    │   │   │   │   │   │   │   │   │   ├── [keyId]/
    │   │   │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # API key usage logs
    │   │   │   │   │   │   │   │   │   │   ├── permissions/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # API key permissions
    │   │   │   │   │   │   │   │   │   │   ├── regenerate/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Regenerate API key
    │   │   │   │   │   │   │   │   │   │   ├── revoke/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Revoke API key
    │   │   │   │   │   │   │   │   │   │   ├── page.tsx  # API key details
    │   │   │   │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create API key
    │   │   │   │   │   │   │   │   │   ├── page.tsx  # API keys list
    │   │   │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   │   │   ├── oauth-apps/
    │   │   │   │   │   │   │   │   │   ├── [appId]/
    │   │   │   │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth app analytics (web)
    │   │   │   │   │   │   │   │   │   │   ├── authorizations/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth app authorizations
    │   │   │   │   │   │   │   │   │   │   ├── credentials/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth credentials
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit OAuth app
    │   │   │   │   │   │   │   │   │   │   ├── scopes/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth app scopes
    │   │   │   │   │   │   │   │   │   │   ├── page.tsx  # OAuth app detail
    │   │   │   │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create OAuth app
    │   │   │   │   │   │   │   │   │   ├── page.tsx  # OAuth apps list
    │   │   │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   │   │   ├── sandbox/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # API sandbox (web)
    │   │   │   │   │   │   │   │   ├── usage/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # API usage dashboard (web)
    │   │   │   │   │   │   │   │   ├── webhooks/
    │   │   │   │   │   │   │   │   │   ├── [webhookId]/
    │   │   │   │   │   │   │   │   │   │   ├── deliveries/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Webhook deliveries (web)
    │   │   │   │   │   │   │   │   │   │   ├── page.tsx  # Webhook detail (web)
    │   │   │   │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Create webhook (web)
    │   │   │   │   │   │   │   │   │   ├── page.tsx  # Webhooks list (web)
    │   │   │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   │   │   ├── page.tsx  # Developer settings
    │   │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   │   ├── devices/
    │   │   │   │   │   │   │   │   └── page.tsx  # Connected devices
    │   │   │   │   │   │   │   ├── integrations/
    │   │   │   │   │   │   │   │   ├── available/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Available integrations
    │   │   │   │   │   │   │   │   ├── calendar/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Calendar integration
    │   │   │   │   │   │   │   │   ├── connected/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Connected integrations
    │   │   │   │   │   │   │   │   ├── slack/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Slack integration
    │   │   │   │   │   │   │   │   └── webhooks/
    │   │   │   │   │   │   │   │       ├── [webhookId]/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Webhook detail
    │   │   │   │   │   │   │   │       └── create/
    │   │   │   │   │   │   │   │           └── page.tsx  # Create webhook
    │   │   │   │   │   │   │   ├── labs/
    │   │   │   │   │   │   │   │   └── page.tsx  # Experimental features
    │   │   │   │   │   │   │   ├── login-history/
    │   │   │   │   │   │   │   │   └── page.tsx  # Login history
    │   │   │   │   │   │   │   ├── notifications/
    │   │   │   │   │   │   │   │   ├── digest/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Email digest settings
    │   │   │   │   │   │   │   │   ├── email/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Email notification settings
    │   │   │   │   │   │   │   │   └── push/
    │   │   │   │   │   │   │   │       └── page.tsx  # Push notification settings
    │   │   │   │   │   │   │   ├── preferences/
    │   │   │   │   │   │   │   │   ├── accessibility/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Accessibility settings
    │   │   │   │   │   │   │   │   ├── appearance/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Appearance preferences
    │   │   │   │   │   │   │   │   ├── language/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Language settings
    │   │   │   │   │   │   │   │   └── timezone/
    │   │   │   │   │   │   │   │       └── page.tsx  # Timezone settings
    │   │   │   │   │   │   │   ├── privacy/
    │   │   │   │   │   │   │   │   ├── activity/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Activity privacy
    │   │   │   │   │   │   │   │   ├── blocked-users/
    │   │   │   │   │   │   │   │   │   ├── add/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Block user
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Blocked users list
    │   │   │   │   │   │   │   │   ├── data-export/
    │   │   │   │   │   │   │   │   │   ├── request/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Request new data export
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Data export (GDPR)
    │   │   │   │   │   │   │   │   ├── data-sharing/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Data sharing preferences
    │   │   │   │   │   │   │   │   ├── gdpr/
    │   │   │   │   │   │   │   │   │   ├── delete/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # GDPR — Delete account
    │   │   │   │   │   │   │   │   │   └── export/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # GDPR — Data export
    │   │   │   │   │   │   │   │   └── profile-visibility/
    │   │   │   │   │   │   │   │       └── page.tsx  # Profile visibility settings
    │   │   │   │   │   │   │   ├── privacy-security/
    │   │   │   │   │   │   │   │   └── authorized-apps/
    │   │   │   │   │   │   │   │       ├── [appId]/
    │   │   │   │   │   │   │   │       │   ├── permissions/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # App permissions detail (web)
    │   │   │   │   │   │   │   │       │   ├── page.tsx  # Authorized app detail (web)
    │   │   │   │   │   │   │   │       │   └── │
    │   │   │   │   │   │   │   │       ├── page.tsx  # Authorized apps list (web)
    │   │   │   │   │   │   │   │       └── │
    │   │   │   │   │   │   │   ├── security/
    │   │   │   │   │   │   │   │   ├── login-history/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Login history
    │   │   │   │   │   │   │   │   ├── password/
    │   │   │   │   │   │   │   │   │   └── change/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Change password
    │   │   │   │   │   │   │   │   ├── sessions/
    │   │   │   │   │   │   │   │   │   ├── revoke-all/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Revoke all sessions (except current)
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Active sessions
    │   │   │   │   │   │   │   │   └── two-factor/
    │   │   │   │   │   │   │   │       ├── disable/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Disable 2FA
    │   │   │   │   │   │   │   │       └── enable/
    │   │   │   │   │   │   │   │           └── page.tsx  # Enable 2FA
    │   │   │   │   │   │   │   ├── two-factor/
    │   │   │   │   │   │   │   │   ├── backup-codes/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Backup codes
    │   │   │   │   │   │   │   │   ├── disable/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Disable 2FA
    │   │   │   │   │   │   │   │   └── methods/
    │   │   │   │   │   │   │   │       ├── app/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Authenticator app setup
    │   │   │   │   │   │   │   │       └── sms/
    │   │   │   │   │   │   │   │           └── page.tsx  # SMS 2FA setup
    │   │   │   │   │   │   │   ├── web-push/
    │   │   │   │   │   │   │   │   └── page.tsx  # Web Push: subscribe/unsubscribe, device listing
    │   │   │   │   │   │   │   ├── page.tsx  # Settings overview
    │   │   │   │   │   │   │   └── │
    │   │   │   │   │   │   ├── sourcing/
    │   │   │   │   │   │   │   ├── campaigns/
    │   │   │   │   │   │   │   │   ├── [campaignId]/
    │   │   │   │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Campaign analytics
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit sourcing campaign
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Campaign detail
    │   │   │   │   │   │   │   │   ├── create/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create sourcing campaign
    │   │   │   │   │   │   │   │   └── page.tsx  # Campaigns list
    │   │   │   │   │   │   │   ├── invitations/
    │   │   │   │   │   │   │   │   ├── [invitationId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Invitation detail
    │   │   │   │   │   │   │   │   └── page.tsx  # Sent invitations
    │   │   │   │   │   │   │   └── talent-pools/
    │   │   │   │   │   │   │       ├── [poolId]/
    │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit talent pool
    │   │   │   │   │   │   │       │   ├── members/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Pool members
    │   │   │   │   │   │   │       │   └── page.tsx  # Talent pool detail
    │   │   │   │   │   │   │       ├── create/
    │   │   │   │   │   │   │       │   └── page.tsx  # Create talent pool
    │   │   │   │   │   │   │       └── page.tsx  # Talent pools list
    │   │   │   │   │   │   ├── specialized/
    │   │   │   │   │   │   │   ├── plus/
    │   │   │   │   │   │   │   │   ├── subscribe/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Subscribe to Plus
    │   │   │   │   │   │   │   │   └── page.tsx  # Plus membership overview
    │   │   │   │   │   │   │   ├── talent-cloud/
    │   │   │   │   │   │   │   │   ├── projects/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Talent Cloud exclusive projects
    │   │   │   │   │   │   │   │   └── page.tsx  # Talent Cloud info
    │   │   │   │   │   │   │   └── top-rated/
    │   │   │   │   │   │   │       ├── application/
    │   │   │   │   │   │   │       │   └── page.tsx  # Apply for Top Rated
    │   │   │   │   │   │   │       └── page.tsx  # Top Rated program info
    │   │   │   │   │   │   ├── spend-analytics/
    │   │   │   │   │   │   │   ├── by-category/
    │   │   │   │   │   │   │   │   └── page.tsx  # Spending by category
    │   │   │   │   │   │   │   ├── by-department/
    │   │   │   │   │   │   │   │   └── page.tsx  # Spending by department
    │   │   │   │   │   │   │   ├── by-vendor/
    │   │   │   │   │   │   │   │   └── page.tsx  # Spending by vendor
    │   │   │   │   │   │   │   ├── forecasting/
    │   │   │   │   │   │   │   │   └── page.tsx  # Spend forecasting
    │   │   │   │   │   │   │   └── page.tsx  # Spend analytics dashboard
    │   │   │   │   │   │   ├── status/
    │   │   │   │   │   │   │   └── page.tsx  # Public status & incidents
    │   │   │   │   │   │   ├── subscription/  # Subscription management
    │   │   │   │   │   │   │   ├── addons/
    │   │   │   │   │   │   │   │   └── [addonId]/
    │   │   │   │   │   │   │   │       ├── cancel/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Cancel addon
    │   │   │   │   │   │   │   │       └── purchase/
    │   │   │   │   │   │   │   │           └── page.tsx  # Purchase addon
    │   │   │   │   │   │   │   ├── billing-history/
    │   │   │   │   │   │   │   │   ├── [invoiceId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Invoice detail
    │   │   │   │   │   │   │   │   └── page.tsx  # Billing history
    │   │   │   │   │   │   │   ├── cancel/
    │   │   │   │   │   │   │   │   └── page.tsx  # Cancel subscription
    │   │   │   │   │   │   │   ├── connects/
    │   │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Connects usage history
    │   │   │   │   │   │   │   │   ├── purchase/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Purchase connects
    │   │   │   │   │   │   │   │   └── page.tsx  # Connects overview
    │   │   │   │   │   │   │   ├── downgrade/
    │   │   │   │   │   │   │   │   └── page.tsx  # Downgrade plan
    │   │   │   │   │   │   │   ├── plans/
    │   │   │   │   │   │   │   │   ├── compare/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Plan comparison
    │   │   │   │   │   │   │   │   └── page.tsx  # Available plans
    │   │   │   │   │   │   │   ├── reactivate/
    │   │   │   │   │   │   │   │   └── page.tsx  # Reactivate subscription
    │   │   │   │   │   │   │   ├── trial/
    │   │   │   │   │   │   │   │   ├── convert/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Convert trial to paid
    │   │   │   │   │   │   │   │   └── page.tsx  # Trial status
    │   │   │   │   │   │   │   ├── upgrade/
    │   │   │   │   │   │   │   │   ├── confirm/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Confirm upgrade
    │   │   │   │   │   │   │   │   └── page.tsx  # Upgrade plan
    │   │   │   │   │   │   │   ├── usage/
    │   │   │   │   │   │   │   │   └── page.tsx  # Usage statistics
    │   │   │   │   │   │   │   └── page.tsx  # Subscription overview
    │   │   │   │   │   │   ├── subscriptions/
    │   │   │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   │   │   └── page.tsx  # Subscription billing details
    │   │   │   │   │   │   │   ├── cancel/
    │   │   │   │   │   │   │   │   └── page.tsx  # Cancel subscription
    │   │   │   │   │   │   │   ├── change-plan/
    │   │   │   │   │   │   │   │   └── page.tsx  # Change subscription plan
    │   │   │   │   │   │   │   ├── features/
    │   │   │   │   │   │   │   │   └── page.tsx  # Subscription features
    │   │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   │   └── page.tsx  # Subscription history
    │   │   │   │   │   │   │   ├── pause/
    │   │   │   │   │   │   │   │   └── page.tsx  # Pause subscription
    │   │   │   │   │   │   │   ├── reactivate/
    │   │   │   │   │   │   │   │   └── page.tsx  # Reactivate subscription
    │   │   │   │   │   │   │   ├── upgrade/
    │   │   │   │   │   │   │   │   └── page.tsx  # Upgrade subscription
    │   │   │   │   │   │   │   └── usage/
    │   │   │   │   │   │   │       └── page.tsx  # Subscription usage metrics
    │   │   │   │   │   │   ├── support/
    │   │   │   │   │   │   │   ├── help-center/
    │   │   │   │   │   │   │   │   └── page.tsx  # Help center (articles search)
    │   │   │   │   │   │   │   └── tickets/
    │   │   │   │   │   │   │       ├── [ticketId]/
    │   │   │   │   │   │   │       │   └── page.tsx  # Ticket detail, replies, SLA
    │   │   │   │   │   │   │       ├── new/
    │   │   │   │   │   │   │       │   └── page.tsx  # Create ticket (attachments)
    │   │   │   │   │   │   │       └── page.tsx  # My tickets list
    │   │   │   │   │   │   ├── talent/
    │   │   │   │   │   │   │   ├── browse/
    │   │   │   │   │   │   │   │   └── page.tsx  # Browse talent
    │   │   │   │   │   │   │   ├── recommendations/
    │   │   │   │   │   │   │   │   └── page.tsx  # AI-recommended talent for jobs
    │   │   │   │   │   │   │   ├── saved/
    │   │   │   │   │   │   │   │   └── page.tsx  # Saved talent profiles
    │   │   │   │   │   │   │   └── shortlists/
    │   │   │   │   │   │   │       ├── [shortlistId]/
    │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit shortlist
    │   │   │   │   │   │   │       │   └── page.tsx  # Shortlist details
    │   │   │   │   │   │   │       ├── new/
    │   │   │   │   │   │   │       │   └── page.tsx  # Create shortlist
    │   │   │   │   │   │   │       └── page.tsx  # Shortlists overview
    │   │   │   │   │   │   ├── teams/
    │   │   │   │   │   │   │   ├── [teamId]/
    │   │   │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Team compliance dashboard
    │   │   │   │   │   │   │   │   ├── hierarchy/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Team hierarchy management
    │   │   │   │   │   │   │   │   ├── performance/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Team performance metrics
    │   │   │   │   │   │   │   │   └── spending-controls/
    │   │   │   │   │   │   │   │       └── page.tsx  # Team spending controls
    │   │   │   │   │   │   │   └── integrations/
    │   │   │   │   │   │   │       ├── [integrationId]/
    │   │   │   │   │   │   │       │   ├── configure/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Configure integration
    │   │   │   │   │   │   │       │   ├── logs/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Integration logs
    │   │   │   │   │   │   │       │   └── page.tsx  # Integration details
    │   │   │   │   │   │   │       ├── available/
    │   │   │   │   │   │   │       │   └── page.tsx  # Available integrations
    │   │   │   │   │   │   │       └── page.tsx  # Active integrations list
    │   │   │   │   │   │   ├── timesheets/
    │   │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   │   ├── [timesheetId]/
    │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit timesheet
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Timesheet details
    │   │   │   │   │   │   │   │   ├── new/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Create timesheet
    │   │   │   │   │   │   │   │   └── page.tsx  # Contract timesheets list
    │   │   │   │   │   │   │   ├── approve/
    │   │   │   │   │   │   │   │   └── page.tsx  # Timesheets pending approval (client)
    │   │   │   │   │   │   │   └── page.tsx  # All timesheets overview
    │   │   │   │   │   │   ├── vendor-management/
    │   │   │   │   │   │   │   ├── blacklist/
    │   │   │   │   │   │   │   │   ├── [userId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Blacklist entry detail
    │   │   │   │   │   │   │   │   └── page.tsx  # Blacklisted vendors
    │   │   │   │   │   │   │   ├── compliance-docs/
    │   │   │   │   │   │   │   │   ├── [vendorId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Vendor compliance documents
    │   │   │   │   │   │   │   │   └── page.tsx  # Compliance tracking
    │   │   │   │   │   │   │   └── preferred/
    │   │   │   │   │   │   │       ├── [vendorId]/
    │   │   │   │   │   │   │       │   ├── history/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Work history with vendor
    │   │   │   │   │   │   │       │   ├── performance/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Vendor performance metrics
    │   │   │   │   │   │   │       │   └── page.tsx  # Vendor detail
    │   │   │   │   │   │   │       └── page.tsx  # Preferred vendors list
    │   │   │   │   │   │   ├── work-diary/
    │   │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   │   ├── calendar/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Calendar view of work diary
    │   │   │   │   │   │   │   │   ├── screenshots/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Screenshots management
    │   │   │   │   │   │   │   │   └── page.tsx  # Work diary detail
    │   │   │   │   │   │   │   └── page.tsx  # Work diary overview (all contracts)
    │   │   │   │   │   │   └── layout.tsx  # Dashboard layout
    │   │   │   │   │   ├── (delivery)/
    │   │   │   │   │   │   └── deliverables/
    │   │   │   │   │   │       └── [contractId]/
    │   │   │   │   │   │           └── page.tsx  # Deliverables & versions
    │   │   │   │   │   ├── (freelancer)/
    │   │   │   │   │   │   ├── portfolio/
    │   │   │   │   │   │   │   └── manage/
    │   │   │   │   │   │   │       └── page.tsx  # Manage portfolio
    │   │   │   │   │   │   └── proposals/
    │   │   │   │   │   │       └── new/
    │   │   │   │   │   │           └── page.tsx  # New proposal (freelancer)
    │   │   │   │   │   ├── (marketing)/  # Marketing routes group
    │   │   │   │   │   ├── (onboarding)/  # Onboarding flow (post-registration)
    │   │   │   │   │   │   ├── client/  # Client onboarding
    │   │   │   │   │   │   │   ├── billing/
    │   │   │   │   │   │   │   │   └── page.tsx  # Billing setup
    │   │   │   │   │   │   │   ├── company/
    │   │   │   │   │   │   │   │   └── page.tsx  # Company info
    │   │   │   │   │   │   │   ├── preferences/
    │   │   │   │   │   │   │   │   └── page.tsx  # Hiring preferences
    │   │   │   │   │   │   │   ├── team/
    │   │   │   │   │   │   │   │   └── page.tsx  # Team setup (optional)
    │   │   │   │   │   │   │   └── verification/
    │   │   │   │   │   │   │       └── page.tsx  # Business verification
    │   │   │   │   │   │   ├── freelancer/  # Freelancer onboarding
    │   │   │   │   │   │   │   ├── experience/
    │   │   │   │   │   │   │   │   └── page.tsx  # Work experience
    │   │   │   │   │   │   │   ├── portfolio/
    │   │   │   │   │   │   │   │   └── page.tsx  # Portfolio items
    │   │   │   │   │   │   │   ├── preferences/
    │   │   │   │   │   │   │   │   └── page.tsx  # Job preferences
    │   │   │   │   │   │   │   ├── profile/
    │   │   │   │   │   │   │   │   └── page.tsx  # Basic profile
    │   │   │   │   │   │   │   ├── rates/
    │   │   │   │   │   │   │   │   └── page.tsx  # Rate setting
    │   │   │   │   │   │   │   └── skills/
    │   │   │   │   │   │   │       └── page.tsx  # Skills selection
    │   │   │   │   │   │   ├── welcome/
    │   │   │   │   │   │   │   └── page.tsx  # Welcome message
    │   │   │   │   │   │   └── layout.tsx  # Onboarding layout
    │   │   │   │   │   ├── (public)/  # Public pages (no auth required)
    │   │   │   │   │   │   ├── about/
    │   │   │   │   │   │   │   ├── careers/
    │   │   │   │   │   │   │   │   └── page.tsx  # Careers page
    │   │   │   │   │   │   │   ├── company/
    │   │   │   │   │   │   │   │   └── page.tsx  # About company
    │   │   │   │   │   │   │   ├── investors/
    │   │   │   │   │   │   │   │   └── page.tsx  # Investor relations
    │   │   │   │   │   │   │   ├── leadership/
    │   │   │   │   │   │   │   │   └── page.tsx  # Leadership team
    │   │   │   │   │   │   │   ├── press/
    │   │   │   │   │   │   │   │   └── page.tsx  # Press releases
    │   │   │   │   │   │   │   ├── team/
    │   │   │   │   │   │   │   │   └── page.tsx  # Team page
    │   │   │   │   │   │   │   └── page.tsx  # About page
    │   │   │   │   │   │   ├── api/
    │   │   │   │   │   │   │   ├── documentation/
    │   │   │   │   │   │   │   │   └── page.tsx  # API docs
    │   │   │   │   │   │   │   ├── pricing/
    │   │   │   │   │   │   │   │   └── page.tsx  # API pricing
    │   │   │   │   │   │   │   └── sdks/
    │   │   │   │   │   │   │       └── page.tsx  # SDK documentation
    │   │   │   │   │   │   ├── blog/
    │   │   │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Blog post detail
    │   │   │   │   │   │   │   ├── category/
    │   │   │   │   │   │   │   │   └── [category]/
    │   │   │   │   │   │   │   │       └── page.tsx  # Category listing
    │   │   │   │   │   │   │   └── page.tsx  # Blog listing
    │   │   │   │   │   │   ├── careers/
    │   │   │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Individual job posting
    │   │   │   │   │   │   │   └── page.tsx  # Careers overview
    │   │   │   │   │   │   ├── case-studies/
    │   │   │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Case study detail
    │   │   │   │   │   │   │   └── page.tsx  # Case studies list
    │   │   │   │   │   │   ├── community/
    │   │   │   │   │   │   │   ├── events/
    │   │   │   │   │   │   │   │   └── page.tsx  # Community events
    │   │   │   │   │   │   │   ├── forum/
    │   │   │   │   │   │   │   │   └── page.tsx  # Community forum
    │   │   │   │   │   │   │   └── page.tsx  # Community hub
    │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   ├── aml/
    │   │   │   │   │   │   │   │   └── page.tsx  # AML policy
    │   │   │   │   │   │   │   ├── kyc/
    │   │   │   │   │   │   │   │   └── page.tsx  # KYC policy
    │   │   │   │   │   │   │   └── page.tsx  # Compliance overview
    │   │   │   │   │   │   ├── contact/
    │   │   │   │   │   │   │   └── page.tsx  # Contact form
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
    │   │   │   │   │   │   │   ├── changelog/
    │   │   │   │   │   │   │   │   └── page.tsx  # API changelog
    │   │   │   │   │   │   │   ├── community/
    │   │   │   │   │   │   │   │   ├── forum/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Developer forum
    │   │   │   │   │   │   │   │   └── github/
    │   │   │   │   │   │   │   │       └── page.tsx  # GitHub repos
    │   │   │   │   │   │   │   ├── getting-started/
    │   │   │   │   │   │   │   │   └── page.tsx  # Developer onboarding
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
    │   │   │   │   │   │   ├── enterprise/
    │   │   │   │   │   │   │   ├── contact/
    │   │   │   │   │   │   │   │   └── page.tsx  # Enterprise contact
    │   │   │   │   │   │   │   ├── demo/
    │   │   │   │   │   │   │   │   └── page.tsx  # Request demo
    │   │   │   │   │   │   │   └── page.tsx  # Enterprise solutions
    │   │   │   │   │   │   ├── features/
    │   │   │   │   │   │   │   ├── [feature]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Feature detail
    │   │   │   │   │   │   │   └── page.tsx  # Features overview
    │   │   │   │   │   │   ├── help/
    │   │   │   │   │   │   │   ├── [category]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Help category
    │   │   │   │   │   │   │   ├── article/
    │   │   │   │   │   │   │   │   └── [slug]/
    │   │   │   │   │   │   │   │       └── page.tsx  # Help article
    │   │   │   │   │   │   │   └── page.tsx  # Help center home
    │   │   │   │   │   │   ├── how-it-works/
    │   │   │   │   │   │   │   ├── clients/
    │   │   │   │   │   │   │   │   └── page.tsx  # For clients
    │   │   │   │   │   │   │   ├── freelancers/
    │   │   │   │   │   │   │   │   └── page.tsx  # For freelancers
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
    │   │   │   │   │   │   │   ├── cookies/
    │   │   │   │   │   │   │   │   └── page.tsx  # Cookie Policy
    │   │   │   │   │   │   │   ├── dmca/
    │   │   │   │   │   │   │   │   └── page.tsx  # DMCA policy
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
    │   │   │   │   │   │   │   ├── terms/
    │   │   │   │   │   │   │   │   ├── client/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Client Agreement
    │   │   │   │   │   │   │   │   ├── freelancer/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Freelancer Agreement
    │   │   │   │   │   │   │   │   ├── service/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Terms of Service
    │   │   │   │   │   │   │   │   └── page.tsx  # Terms of Service
    │   │   │   │   │   │   │   └── page.tsx  # Legal overview
    │   │   │   │   │   │   ├── partners/
    │   │   │   │   │   │   │   ├── affiliates/
    │   │   │   │   │   │   │   │   └── page.tsx  # Affiliate program
    │   │   │   │   │   │   │   ├── become-partner/
    │   │   │   │   │   │   │   │   └── page.tsx  # Partner application
    │   │   │   │   │   │   │   ├── directory/
    │   │   │   │   │   │   │   │   └── page.tsx  # Partners directory
    │   │   │   │   │   │   │   ├── integrations/
    │   │   │   │   │   │   │   │   └── page.tsx  # Integrations
    │   │   │   │   │   │   │   └── page.tsx  # Partners program
    │   │   │   │   │   │   ├── pricing/
    │   │   │   │   │   │   │   ├── compare/
    │   │   │   │   │   │   │   │   └── page.tsx  # Plan comparison
    │   │   │   │   │   │   │   ├── enterprise/
    │   │   │   │   │   │   │   │   └── page.tsx  # Enterprise pricing
    │   │   │   │   │   │   │   └── page.tsx  # Pricing page
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
    │   │   │   │   │   │   │   ├── downloads/
    │   │   │   │   │   │   │   │   └── page.tsx  # Downloads
    │   │   │   │   │   │   │   ├── guides/
    │   │   │   │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Guide detail
    │   │   │   │   │   │   │   │   └── page.tsx  # All guides
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
    │   │   │   │   │   │   │   │   └── page.tsx  # Tutorials list
    │   │   │   │   │   │   │   ├── webinars/
    │   │   │   │   │   │   │   │   ├── [id]/
    │   │   │   │   │   │   │   │   │   ├── register/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Register for webinar
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Webinar detail
    │   │   │   │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   │   │   │   ├── register/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Webinar registration
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Webinar detail
    │   │   │   │   │   │   │   │   ├── on-demand/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # On-demand recordings
    │   │   │   │   │   │   │   │   ├── upcoming/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Upcoming webinars
    │   │   │   │   │   │   │   │   └── page.tsx  # Webinars list
    │   │   │   │   │   │   │   └── page.tsx  # Resources hub
    │   │   │   │   │   │   ├── security/
    │   │   │   │   │   │   │   └── page.tsx  # Security information
    │   │   │   │   │   │   ├── sitemap/
    │   │   │   │   │   │   │   └── page.tsx  # Sitemap
    │   │   │   │   │   │   ├── solutions/
    │   │   │   │   │   │   │   ├── agencies/
    │   │   │   │   │   │   │   │   └── page.tsx  # Agency solutions
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
    │   │   │   │   │   │   │       └── page.tsx  # Enterprise solutions
    │   │   │   │   │   │   ├── status/
    │   │   │   │   │   │   │   └── page.tsx  # System status page
    │   │   │   │   │   │   ├── success-stories/
    │   │   │   │   │   │   │   ├── [slug]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Success story
    │   │   │   │   │   │   │   └── page.tsx  # Success stories
    │   │   │   │   │   │   ├── trust-safety/
    │   │   │   │   │   │   │   └── page.tsx  # Trust & Safety
    │   │   │   │   │   │   ├── layout.tsx  # Public pages layout
    │   │   │   │   │   │   ├── page.tsx  # Homepage / Landing
    │   │   │   │   │   │   └── sitemap.xml  # Dynamic sitemap
    │   │   │   │   │   ├── (reviews)/
    │   │   │   │   │   │   └── reviews/
    │   │   │   │   │   │       └── new/
    │   │   │   │   │   │           └── [contractId]/
    │   │   │   │   │   │               └── page.tsx  # Write a review
    │   │   │   │   │   ├── (social)/
    │   │   │   │   │   │   └── network/
    │   │   │   │   │   │       └── page.tsx  # Network & groups
    │   │   │   │   │   ├── (support)/
    │   │   │   │   │   │   └── disputes/
    │   │   │   │   │   │       └── [contractId]/
    │   │   │   │   │   │           └── page.tsx  # Dispute center
    │   │   │   │   │   ├── contracts/
    │   │   │   │   │   │   ├── [contractId]/
    │   │   │   │   │   │   │   ├── change-requests/
    │   │   │   │   │   │   │   │   ├── [requestId]/
    │   │   │   │   │   │   │   │   │   ├── approval-chain/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Change request approvals
    │   │   │   │   │   │   │   │   │   ├── impact-analysis/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Change impact analysis
    │   │   │   │   │   │   │   │   │   └── negotiate/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Negotiate change request
    │   │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │   │       └── page.tsx  # Change request templates
    │   │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   │   ├── audits/
    │   │   │   │   │   │   │   │   │   ├── [auditId]/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Audit detail
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Audits list
    │   │   │   │   │   │   │   │   ├── documents/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Compliance documents
    │   │   │   │   │   │   │   │   └── reports/
    │   │   │   │   │   │   │   │       └── page.tsx  # Compliance reports
    │   │   │   │   │   │   │   ├── deliverables/
    │   │   │   │   │   │   │   │   ├── [deliverableId]/
    │   │   │   │   │   │   │   │   │   ├── approvals/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approval workflow
    │   │   │   │   │   │   │   │   │   ├── feedback/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Deliverable feedback
    │   │   │   │   │   │   │   │   │   └── versions/
    │   │   │   │   │   │   │   │   │       ├── [versionId]/
    │   │   │   │   │   │   │   │   │       │   └── page.tsx  # Specific version detail
    │   │   │   │   │   │   │   │   │       └── compare/
    │   │   │   │   │   │   │   │   │           └── page.tsx  # Compare versions
    │   │   │   │   │   │   │   │   └── bulk-operations/
    │   │   │   │   │   │   │   │       └── page.tsx  # Bulk deliverable actions
    │   │   │   │   │   │   │   ├── knowledge-transfer/
    │   │   │   │   │   │   │   │   ├── checklist/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Transfer checklist
    │   │   │   │   │   │   │   │   ├── documentation/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Transfer documentation
    │   │   │   │   │   │   │   │   └── sessions/
    │   │   │   │   │   │   │   │       ├── [sessionId]/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Session detail
    │   │   │   │   │   │   │   │       └── page.tsx  # Sessions list
    │   │   │   │   │   │   │   ├── milestones/
    │   │   │   │   │   │   │   │   ├── [milestoneId]/
    │   │   │   │   │   │   │   │   │   ├── review/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Review milestone
    │   │   │   │   │   │   │   │   │   ├── revisions/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Milestone revision history
    │   │   │   │   │   │   │   │   │   └── submission/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Submit milestone
    │   │   │   │   │   │   │   │   └── reorder/
    │   │   │   │   │   │   │   │       └── page.tsx  # Reorder milestones
    │   │   │   │   │   │   │   ├── quality/
    │   │   │   │   │   │   │   │   ├── metrics/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Quality metrics
    │   │   │   │   │   │   │   │   ├── reviews/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Quality reviews
    │   │   │   │   │   │   │   │   └── standards/
    │   │   │   │   │   │   │   │       └── page.tsx  # Quality standards
    │   │   │   │   │   │   │   ├── risks/
    │   │   │   │   │   │   │   │   ├── monitoring/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Risk monitoring
    │   │   │   │   │   │   │   │   ├── register/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Risk register
    │   │   │   │   │   │   │   │   └── reports/
    │   │   │   │   │   │   │   │       └── page.tsx  # Risk reports
    │   │   │   │   │   │   │   └── work-diary/
    │   │   │   │   │   │   │       ├── activity-levels/
    │   │   │   │   │   │   │       │   └── page.tsx  # Activity level tracking
    │   │   │   │   │   │   │       ├── bulk-entry/
    │   │   │   │   │   │   │       │   └── page.tsx  # Bulk time entry
    │   │   │   │   │   │   │       ├── corrections/
    │   │   │   │   │   │   │       │   └── page.tsx  # Time entry corrections
    │   │   │   │   │   │   │       └── screenshots/
    │   │   │   │   │   │   │           └── page.tsx  # Screenshot gallery
    │   │   │   │   │   │   └── benchmarking/
    │   │   │   │   │   │       ├── costs/
    │   │   │   │   │   │       │   └── page.tsx  # Cost benchmarking
    │   │   │   │   │   │       ├── performance/
    │   │   │   │   │   │       │   └── page.tsx  # Performance benchmarking
    │   │   │   │   │   │       └── quality/
    │   │   │   │   │   │           └── page.tsx  # Quality benchmarking
    │   │   │   │   │   ├── developers/
    │   │   │   │   │   │   ├── api-reference/
    │   │   │   │   │   │   │   ├── [endpoint]/
    │   │   │   │   │   │   │   │   └── page.tsx  # API endpoint reference
    │   │   │   │   │   │   │   └── page.tsx  # API reference home
    │   │   │   │   │   │   ├── docs/
    │   │   │   │   │   │   │   ├── [section]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Documentation section
    │   │   │   │   │   │   │   └── page.tsx  # API documentation home
    │   │   │   │   │   │   ├── sandbox/
    │   │   │   │   │   │   │   └── page.tsx  # API sandbox/playground
    │   │   │   │   │   │   ├── sdks/
    │   │   │   │   │   │   │   └── page.tsx  # SDK downloads and docs
    │   │   │   │   │   │   └── webhooks/
    │   │   │   │   │   │       └── page.tsx  # Webhooks documentation
    │   │   │   │   │   ├── enterprise/
    │   │   │   │   │   │   ├── case-studies/
    │   │   │   │   │   │   │   └── page.tsx  # Enterprise case studies
    │   │   │   │   │   │   ├── contact/
    │   │   │   │   │   │   │   └── page.tsx  # Enterprise contact/demo request
    │   │   │   │   │   │   ├── pricing/
    │   │   │   │   │   │   │   └── page.tsx  # Enterprise pricing
    │   │   │   │   │   │   └── solutions/
    │   │   │   │   │   │       ├── managed-services/
    │   │   │   │   │   │       │   └── page.tsx  # Managed services offering
    │   │   │   │   │   │       ├── staffing/
    │   │   │   │   │   │       │   └── page.tsx  # Enterprise staffing solutions
    │   │   │   │   │   │       └── page.tsx  # Enterprise solutions overview
    │   │   │   │   │   ├── financial/
    │   │   │   │   │   │   ├── analytics/
    │   │   │   │   │   │   │   ├── kpis/
    │   │   │   │   │   │   │   │   ├── custom/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Custom KPI builder
    │   │   │   │   │   │   │   │   └── dashboard/
    │   │   │   │   │   │   │   │       └── page.tsx  # Financial KPI dashboard
    │   │   │   │   │   │   │   ├── margins/
    │   │   │   │   │   │   │   │   └── page.tsx  # Margin analysis
    │   │   │   │   │   │   │   └── profitability/
    │   │   │   │   │   │   │       ├── by-client/
    │   │   │   │   │   │   │       │   └── page.tsx  # Client profitability
    │   │   │   │   │   │   │       ├── by-project/
    │   │   │   │   │   │   │       │   └── page.tsx  # Project profitability
    │   │   │   │   │   │   │       └── by-service/
    │   │   │   │   │   │   │           └── page.tsx  # Service line profitability
    │   │   │   │   │   │   ├── budgets/
    │   │   │   │   │   │   │   ├── [budgetId]/
    │   │   │   │   │   │   │   │   ├── adjustments/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Budget adjustments
    │   │   │   │   │   │   │   │   ├── allocations/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Budget allocations
    │   │   │   │   │   │   │   │   ├── forecasts/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Budget forecasts
    │   │   │   │   │   │   │   │   └── variance/
    │   │   │   │   │   │   │   │       └── page.tsx  # Budget variance
    │   │   │   │   │   │   │   ├── consolidation/
    │   │   │   │   │   │   │   │   └── page.tsx  # Budget consolidation
    │   │   │   │   │   │   │   └── templates/
    │   │   │   │   │   │   │       └── page.tsx  # Budget templates
    │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   ├── aml/
    │   │   │   │   │   │   │   │   ├── monitoring/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # AML monitoring
    │   │   │   │   │   │   │   │   ├── reports/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # AML reports
    │   │   │   │   │   │   │   │   └── page.tsx  # AML dashboard
    │   │   │   │   │   │   │   ├── audits/
    │   │   │   │   │   │   │   │   ├── findings/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Audit findings
    │   │   │   │   │   │   │   │   ├── remediation/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Remediation tracking
    │   │   │   │   │   │   │   │   └── schedule/
    │   │   │   │   │   │   │   │       └── page.tsx  # Audit schedule
    │   │   │   │   │   │   │   ├── kyc/
    │   │   │   │   │   │   │   │   ├── due-diligence/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Enhanced due diligence
    │   │   │   │   │   │   │   │   └── verification/
    │   │   │   │   │   │   │   │       └── page.tsx  # KYC verification
    │   │   │   │   │   │   │   └── sanctions/
    │   │   │   │   │   │   │       ├── alerts/
    │   │   │   │   │   │   │       │   └── page.tsx  # Sanctions alerts
    │   │   │   │   │   │   │       └── screening/
    │   │   │   │   │   │   │           └── page.tsx  # Sanctions screening
    │   │   │   │   │   │   ├── credit-management/
    │   │   │   │   │   │   │   ├── collections/
    │   │   │   │   │   │   │   │   ├── actions/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Collection actions
    │   │   │   │   │   │   │   │   ├── aging/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Aging analysis
    │   │   │   │   │   │   │   │   └── page.tsx  # Collections dashboard
    │   │   │   │   │   │   │   ├── disputes/
    │   │   │   │   │   │   │   │   └── page.tsx  # Credit disputes
    │   │   │   │   │   │   │   ├── limits/
    │   │   │   │   │   │   │   │   ├── [clientId]/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Client credit limit
    │   │   │   │   │   │   │   │   └── page.tsx  # All credit limits
    │   │   │   │   │   │   │   └── scoring/
    │   │   │   │   │   │   │       └── page.tsx  # Credit scoring
    │   │   │   │   │   │   ├── forecasting/
    │   │   │   │   │   │   │   ├── expenses/
    │   │   │   │   │   │   │   │   └── page.tsx  # Expense forecast
    │   │   │   │   │   │   │   ├── profitability/
    │   │   │   │   │   │   │   │   └── page.tsx  # Profitability forecast
    │   │   │   │   │   │   │   ├── revenue/
    │   │   │   │   │   │   │   │   ├── models/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Revenue models
    │   │   │   │   │   │   │   │   ├── scenarios/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Revenue scenarios
    │   │   │   │   │   │   │   │   └── page.tsx  # Revenue forecast
    │   │   │   │   │   │   │   └── validation/
    │   │   │   │   │   │   │       └── page.tsx  # Forecast validation
    │   │   │   │   │   │   ├── reconciliation/
    │   │   │   │   │   │   │   ├── bank/
    │   │   │   │   │   │   │   │   ├── [accountId]/
    │   │   │   │   │   │   │   │   │   ├── auto-match/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Auto-matching
    │   │   │   │   │   │   │   │   │   ├── discrepancies/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Discrepancy resolution
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Account reconciliation
    │   │   │   │   │   │   │   │   └── page.tsx  # Bank reconciliation list
    │   │   │   │   │   │   │   ├── intercompany/
    │   │   │   │   │   │   │   │   └── page.tsx  # Intercompany reconciliation
    │   │   │   │   │   │   │   ├── merchant/
    │   │   │   │   │   │   │   │   └── page.tsx  # Merchant account reconciliation
    │   │   │   │   │   │   │   └── reports/
    │   │   │   │   │   │   │       └── page.tsx  # Reconciliation reports
    │   │   │   │   │   │   ├── risk-management/
    │   │   │   │   │   │   │   ├── exposure/
    │   │   │   │   │   │   │   │   ├── credit/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Credit exposure
    │   │   │   │   │   │   │   │   ├── currency/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Currency exposure
    │   │   │   │   │   │   │   │   └── operational/
    │   │   │   │   │   │   │   │       └── page.tsx  # Operational risk
    │   │   │   │   │   │   │   ├── limits/
    │   │   │   │   │   │   │   │   └── page.tsx  # Risk limits
    │   │   │   │   │   │   │   └── reports/
    │   │   │   │   │   │   │       ├── stress-testing/
    │   │   │   │   │   │   │       │   └── page.tsx  # Stress testing
    │   │   │   │   │   │   │       └── var/
    │   │   │   │   │   │   │           └── page.tsx  # Value at Risk
    │   │   │   │   │   │   └── treasury/
    │   │   │   │   │   │       ├── cash-flow/
    │   │   │   │   │   │       │   ├── analysis/
    │   │   │   │   │   │       │   │   └── page.tsx  # Cash flow analysis
    │   │   │   │   │   │       │   ├── forecast/
    │   │   │   │   │   │       │   │   └── page.tsx  # Cash flow forecasting
    │   │   │   │   │   │       │   └── reports/
    │   │   │   │   │   │       │       └── page.tsx  # Cash flow reports
    │   │   │   │   │   │       ├── dashboard/
    │   │   │   │   │   │       │   └── page.tsx  # Treasury dashboard
    │   │   │   │   │   │       ├── investments/
    │   │   │   │   │   │       │   └── page.tsx  # Short-term investments
    │   │   │   │   │   │       └── liquidity/
    │   │   │   │   │   │           ├── alerts/
    │   │   │   │   │   │           │   └── page.tsx  # Liquidity alerts
    │   │   │   │   │   │           └── management/
    │   │   │   │   │   │               └── page.tsx  # Liquidity management
    │   │   │   │   │   ├── legal/
    │   │   │   │   │   │   ├── accessibility/
    │   │   │   │   │   │   │   └── page.tsx  # Accessibility statement
    │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   ├── ccpa/
    │   │   │   │   │   │   │   │   └── page.tsx  # CCPA compliance
    │   │   │   │   │   │   │   ├── gdpr/
    │   │   │   │   │   │   │   │   └── page.tsx  # GDPR compliance
    │   │   │   │   │   │   │   └── page.tsx  # Compliance overview
    │   │   │   │   │   │   ├── dmca/
    │   │   │   │   │   │   │   └── page.tsx  # DMCA policy
    │   │   │   │   │   │   ├── ip-policy/
    │   │   │   │   │   │   │   └── page.tsx  # Intellectual property policy
    │   │   │   │   │   │   ├── privacy/
    │   │   │   │   │   │   │   ├── cookie-policy/
    │   │   │   │   │   │   │   │   └── page.tsx  # Cookie policy
    │   │   │   │   │   │   │   ├── data-processing/
    │   │   │   │   │   │   │   │   └── page.tsx  # Data processing agreement
    │   │   │   │   │   │   │   └── page.tsx  # Privacy policy
    │   │   │   │   │   │   └── terms/
    │   │   │   │   │   │       ├── client/
    │   │   │   │   │   │       │   └── page.tsx  # Client terms of service
    │   │   │   │   │   │       ├── freelancer/
    │   │   │   │   │   │       │   └── page.tsx  # Freelancer terms of service
    │   │   │   │   │   │       └── page.tsx  # General terms
    │   │   │   │   │   ├── notifications/
    │   │   │   │   │   │   └── page.tsx  # Notifications center (web)
    │   │   │   │   │   ├── pricing/
    │   │   │   │   │   │   └── page.tsx  # Pricing (public)
    │   │   │   │   │   ├── resources/
    │   │   │   │   │   │   ├── blog/
    │   │   │   │   │   │   │   ├── [postId]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Blog post
    │   │   │   │   │   │   │   ├── category/
    │   │   │   │   │   │   │   │   └── [categoryId]/
    │   │   │   │   │   │   │   │       └── page.tsx  # Blog category
    │   │   │   │   │   │   │   └── page.tsx  # Blog home
    │   │   │   │   │   │   ├── case-studies/
    │   │   │   │   │   │   │   ├── [caseStudyId]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Case study detail
    │   │   │   │   │   │   │   └── page.tsx  # Case studies list
    │   │   │   │   │   │   ├── faq/
    │   │   │   │   │   │   │   └── page.tsx  # Frequently asked questions
    │   │   │   │   │   │   ├── guides/
    │   │   │   │   │   │   │   ├── [guideId]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Guide detail
    │   │   │   │   │   │   │   ├── client/
    │   │   │   │   │   │   │   │   └── page.tsx  # Client guides
    │   │   │   │   │   │   │   ├── freelancer/
    │   │   │   │   │   │   │   │   └── page.tsx  # Freelancer guides
    │   │   │   │   │   │   │   └── page.tsx  # All guides
    │   │   │   │   │   │   ├── tutorials/
    │   │   │   │   │   │   │   ├── [tutorialId]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Tutorial detail
    │   │   │   │   │   │   │   └── page.tsx  # Tutorials list
    │   │   │   │   │   │   └── webinars/
    │   │   │   │   │   │       ├── [webinarId]/
    │   │   │   │   │   │       │   └── page.tsx  # Webinar detail & registration
    │   │   │   │   │   │       └── page.tsx  # Upcoming webinars
    │   │   │   │   │   ├── search/
    │   │   │   │   │   │   ├── alerts/
    │   │   │   │   │   │   │   └── page.tsx  # Search alerts (web)
    │   │   │   │   │   │   ├── assist/
    │   │   │   │   │   │   │   └── page.tsx  # Search assist
    │   │   │   │   │   │   ├── feed/
    │   │   │   │   │   │   │   └── page.tsx  # Personalized feed
    │   │   │   │   │   │   ├── portfolios/
    │   │   │   │   │   │   │   └── page.tsx  # Portfolio search
    │   │   │   │   │   │   └── saved/
    │   │   │   │   │   │       └── page.tsx  # Saved searches (web)
    │   │   │   │   │   ├── security/
    │   │   │   │   │   │   ├── bug-bounty/
    │   │   │   │   │   │   │   └── page.tsx  # Bug bounty program
    │   │   │   │   │   │   ├── certifications/
    │   │   │   │   │   │   │   └── page.tsx  # Security certifications (SOC2, ISO, etc.)
    │   │   │   │   │   │   ├── overview/
    │   │   │   │   │   │   │   └── page.tsx  # Security overview
    │   │   │   │   │   │   └── responsible-disclosure/
    │   │   │   │   │   │       └── page.tsx  # Responsible disclosure policy
    │   │   │   │   │   ├── settings/
    │   │   │   │   │   │   ├── account/
    │   │   │   │   │   │   │   └── saved-items/
    │   │   │   │   │   │   │       └── page.tsx  # Saved items
    │   │   │   │   │   │   ├── advanced/
    │   │   │   │   │   │   │   ├── audit-logs/
    │   │   │   │   │   │   │   │   ├── export/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Export audit logs
    │   │   │   │   │   │   │   │   └── page.tsx  # Audit log settings
    │   │   │   │   │   │   │   ├── data-retention/
    │   │   │   │   │   │   │   │   └── page.tsx  # Data retention policies
    │   │   │   │   │   │   │   ├── ip-whitelist/
    │   │   │   │   │   │   │   │   └── page.tsx  # IP whitelist
    │   │   │   │   │   │   │   ├── rate-limiting/
    │   │   │   │   │   │   │   │   └── page.tsx  # Rate limit configuration
    │   │   │   │   │   │   │   └── session-management/
    │   │   │   │   │   │   │       └── page.tsx  # Session settings
    │   │   │   │   │   │   ├── developer/
    │   │   │   │   │   │   │   ├── api-keys/
    │   │   │   │   │   │   │   │   ├── [keyId]/
    │   │   │   │   │   │   │   │   │   ├── logs/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # API key usage logs
    │   │   │   │   │   │   │   │   │   ├── permissions/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Key permissions
    │   │   │   │   │   │   │   │   │   └── rotate/
    │   │   │   │   │   │   │   │   │       └── page.tsx  # Rotate API key
    │   │   │   │   │   │   │   │   └── page.tsx  # API keys management
    │   │   │   │   │   │   │   ├── documentation/
    │   │   │   │   │   │   │   │   └── page.tsx  # API documentation
    │   │   │   │   │   │   │   ├── oauth-apps/
    │   │   │   │   │   │   │   │   ├── [appId]/
    │   │   │   │   │   │   │   │   │   ├── authorizations/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # App authorizations
    │   │   │   │   │   │   │   │   │   ├── credentials/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth credentials
    │   │   │   │   │   │   │   │   │   ├── scopes/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth scopes
    │   │   │   │   │   │   │   │   │   └── page.tsx  # OAuth app detail
    │   │   │   │   │   │   │   │   └── page.tsx  # OAuth apps
    │   │   │   │   │   │   │   ├── sandbox/
    │   │   │   │   │   │   │   │   ├── environments/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Sandbox environments
    │   │   │   │   │   │   │   │   └── test-data/
    │   │   │   │   │   │   │   │       └── page.tsx  # Test data management
    │   │   │   │   │   │   │   └── webhooks/
    │   │   │   │   │   │   │       ├── [webhookId]/
    │   │   │   │   │   │   │       │   ├── deliveries/
    │   │   │   │   │   │   │       │   │   ├── [deliveryId]/
    │   │   │   │   │   │   │       │   │   │   └── page.tsx  # Delivery detail
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Webhook deliveries
    │   │   │   │   │   │   │       │   ├── events/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Webhook event config
    │   │   │   │   │   │   │       │   ├── test/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Test webhook
    │   │   │   │   │   │   │       │   └── page.tsx  # Webhook detail
    │   │   │   │   │   │   │       └── page.tsx  # Webhooks management
    │   │   │   │   │   │   ├── labs/
    │   │   │   │   │   │   │   └── page.tsx  # Experimental features
    │   │   │   │   │   │   ├── notifications/
    │   │   │   │   │   │   │   └── digests-quiet-hours/
    │   │   │   │   │   │   │       └── page.tsx  # Email digests & quiet hours
    │   │   │   │   │   │   └── security/
    │   │   │   │   │   │       ├── data-delete/
    │   │   │   │   │   │       │   └── page.tsx  # Request delete
    │   │   │   │   │   │       ├── data-export/
    │   │   │   │   │   │       │   └── page.tsx  # Export my data
    │   │   │   │   │   │       ├── deactivate/
    │   │   │   │   │   │       │   └── page.tsx  # Deactivate account
    │   │   │   │   │   │       ├── devices/
    │   │   │   │   │   │       │   └── page.tsx  # Devices & sessions
    │   │   │   │   │   │       ├── devices-sessions/
    │   │   │   │   │   │       │   └── page.tsx  # Active devices & sessions
    │   │   │   │   │   │       ├── mfa/
    │   │   │   │   │   │       │   └── page.tsx  # MFA settings
    │   │   │   │   │   │       └── password/
    │   │   │   │   │   │           └── page.tsx  # Change password
    │   │   │   │   │   ├── status/
    │   │   │   │   │   │   ├── current/
    │   │   │   │   │   │   │   └── page.tsx  # Current system status
    │   │   │   │   │   │   ├── history/
    │   │   │   │   │   │   │   └── page.tsx  # Status history
    │   │   │   │   │   │   ├── subscribe/
    │   │   │   │   │   │   │   └── page.tsx  # Subscribe to status updates
    │   │   │   │   │   │   └── page.tsx  # Status page (public)
    │   │   │   │   │   ├── teams/
    │   │   │   │   │   │   ├── [teamId]/
    │   │   │   │   │   │   │   ├── capacity/
    │   │   │   │   │   │   │   │   ├── forecasting/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Capacity forecasting
    │   │   │   │   │   │   │   │   ├── planning/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Capacity planning
    │   │   │   │   │   │   │   │   └── utilization/
    │   │   │   │   │   │   │   │       └── page.tsx  # Utilization tracking
    │   │   │   │   │   │   │   ├── hierarchy/
    │   │   │   │   │   │   │   │   ├── org-chart/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Organization chart
    │   │   │   │   │   │   │   │   ├── reporting-lines/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Reporting structure
    │   │   │   │   │   │   │   │   └── page.tsx  # Team hierarchy overview
    │   │   │   │   │   │   │   ├── knowledge-base/
    │   │   │   │   │   │   │   │   ├── articles/
    │   │   │   │   │   │   │   │   │   ├── [articleId]/
    │   │   │   │   │   │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Edit article
    │   │   │   │   │   │   │   │   │   │   ├── versions/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Article versions
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Article detail
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Articles list
    │   │   │   │   │   │   │   │   ├── categories/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # KB categories
    │   │   │   │   │   │   │   │   └── search/
    │   │   │   │   │   │   │   │       └── page.tsx  # KB search
    │   │   │   │   │   │   │   ├── policies/
    │   │   │   │   │   │   │   │   ├── [policyId]/
    │   │   │   │   │   │   │   │   │   ├── attestations/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Policy attestations
    │   │   │   │   │   │   │   │   │   ├── versions/
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Policy versions
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Policy detail
    │   │   │   │   │   │   │   │   └── page.tsx  # Policies list
    │   │   │   │   │   │   │   ├── procurement/
    │   │   │   │   │   │   │   │   ├── contracts/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Procurement contracts
    │   │   │   │   │   │   │   │   ├── requests/
    │   │   │   │   │   │   │   │   │   ├── [requestId]/
    │   │   │   │   │   │   │   │   │   │   ├── approval/
    │   │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Approve procurement
    │   │   │   │   │   │   │   │   │   │   └── page.tsx  # Request detail
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Procurement requests
    │   │   │   │   │   │   │   │   └── vendors/
    │   │   │   │   │   │   │   │       ├── evaluation/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Vendor evaluation
    │   │   │   │   │   │   │   │       ├── preferred/
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Preferred vendors
    │   │   │   │   │   │   │   │       └── page.tsx  # Vendor management
    │   │   │   │   │   │   │   ├── training/
    │   │   │   │   │   │   │   │   ├── certifications/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Team certifications
    │   │   │   │   │   │   │   │   ├── compliance/
    │   │   │   │   │   │   │   │   │   └── page.tsx  # Training compliance
    │   │   │   │   │   │   │   │   └── programs/
    │   │   │   │   │   │   │   │       ├── [programId]/
    │   │   │   │   │   │   │   │       │   ├── enroll/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Enroll in program
    │   │   │   │   │   │   │   │       │   ├── progress/
    │   │   │   │   │   │   │   │       │   │   └── page.tsx  # Training progress
    │   │   │   │   │   │   │   │       │   └── page.tsx  # Program detail
    │   │   │   │   │   │   │   │       └── page.tsx  # Training programs
    │   │   │   │   │   │   │   └── workflows/
    │   │   │   │   │   │   │       ├── [workflowId]/
    │   │   │   │   │   │   │       │   ├── analytics/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Workflow analytics
    │   │   │   │   │   │   │       │   ├── edit/
    │   │   │   │   │   │   │       │   │   └── page.tsx  # Edit workflow
    │   │   │   │   │   │   │       │   └── page.tsx  # Workflow detail
    │   │   │   │   │   │   │       └── page.tsx  # Workflows list
    │   │   │   │   │   │   └── cross-team/
    │   │   │   │   │   │       ├── benchmarking/
    │   │   │   │   │   │       │   └── page.tsx  # Inter-team benchmarking
    │   │   │   │   │   │       └── collaboration/
    │   │   │   │   │   │           └── page.tsx  # Cross-team collaboration
    │   │   │   │   │   ├── transparency/
    │   │   │   │   │   │   └── page.tsx  # Transparency report
    │   │   │   │   │   ├── error.tsx  # Error boundary
    │   │   │   │   │   ├── layout.tsx  # Root layout for locale
    │   │   │   │   │   └── page.tsx  # Home page
    │   │   │   │   ├── api/
    │   │   │   │   ├── operations/
    │   │   │   │   │   └── search-quality/
    │   │   │   │   │       ├── boosts/
    │   │   │   │   │       │   └── page.tsx  # Search boosts
    │   │   │   │   │       ├── reindexing/
    │   │   │   │   │       │   ├── schedule/
    │   │   │   │   │       │   │   └── page.tsx  # Reindex scheduling
    │   │   │   │   │       │   └── status/
    │   │   │   │   │       │       └── page.tsx  # Reindex status
    │   │   │   │   │       ├── relevance/
    │   │   │   │   │       │   ├── metrics/
    │   │   │   │   │       │   │   └── page.tsx  # Relevance metrics
    │   │   │   │   │       │   ├── testing/
    │   │   │   │   │       │   │   └── page.tsx  # Relevance testing
    │   │   │   │   │       │   └── tuning/
    │   │   │   │   │       │       └── page.tsx  # Relevance tuning
    │   │   │   │   │       └── synonyms/
    │   │   │   │   │           ├── [synonymId]/
    │   │   │   │   │           │   └── page.tsx  # Synonym detail
    │   │   │   │   │           └── page.tsx  # Synonym management
    │   │   │   │   ├── trust-safety/
    │   │   │   │   │   ├── content-moderation/
    │   │   │   │   │   │   ├── appeals/
    │   │   │   │   │   │   │   ├── [appealId]/
    │   │   │   │   │   │   │   │   └── page.tsx  # Appeal review
    │   │   │   │   │   │   │   └── page.tsx  # Appeals queue
    │   │   │   │   │   │   ├── automation/
    │   │   │   │   │   │   │   ├── actions/
    │   │   │   │   │   │   │   │   └── page.tsx  # Automated actions
    │   │   │   │   │   │   │   └── rules/
    │   │   │   │   │   │   │       └── page.tsx  # Auto-moderation rules
    │   │   │   │   │   │   ├── ml-assistance/
    │   │   │   │   │   │   │   ├── accuracy/
    │   │   │   │   │   │   │   │   └── page.tsx  # Model accuracy
    │   │   │   │   │   │   │   ├── predictions/
    │   │   │   │   │   │   │   │   └── page.tsx  # ML predictions
    │   │   │   │   │   │   │   └── training/
    │   │   │   │   │   │   │       └── page.tsx  # Model training
    │   │   │   │   │   │   └── queue/
    │   │   │   │   │   │       ├── categories/
    │   │   │   │   │   │       │   └── [category]/
    │   │   │   │   │   │       │       └── page.tsx  # Category-specific queue
    │   │   │   │   │   │       ├── priority/
    │   │   │   │   │   │       │   └── page.tsx  # Priority queue
    │   │   │   │   │   │       └── page.tsx  # Moderation queue
    │   │   │   │   │   ├── risk-scoring/
    │   │   │   │   │   │   ├── models/
    │   │   │   │   │   │   │   └── page.tsx  # Risk scoring models
    │   │   │   │   │   │   ├── monitoring/
    │   │   │   │   │   │   │   └── page.tsx  # Risk monitoring
    │   │   │   │   │   │   └── thresholds/
    │   │   │   │   │   │       └── page.tsx  # Risk thresholds
    │   │   │   │   │   └── watchlists/
    │   │   │   │   │       ├── custom/
    │   │   │   │   │       │   └── page.tsx  # Custom watchlists
    │   │   │   │   │       ├── global/
    │   │   │   │   │       │   └── page.tsx  # Global watchlists
    │   │   │   │   │       └── monitoring/
    │   │   │   │   │           └── page.tsx  # Watchlist monitoring
    │   │   │   │   └── │
    │   │   │   ├── features/
    │   │   │   │   ├── budget/
    │   │   │   │   │   └── api/
    │   │   │   │   │       └── mutations.ts  # Budget mutations
    │   │   │   │   ├── interviews/
    │   │   │   │   │   └── api/
    │   │   │   │   │       └── mutations.ts  # Interview scheduling
    │   │   │   │   ├── moderation/
    │   │   │   │   │   └── api/
    │   │   │   │   │       └── mutations.ts  # Moderation actions
    │   │   │   │   ├── notifications/
    │   │   │   │   │   └── api/
    │   │   │   │   │       ├── mutations.ts  # Notifications mutations
    │   │   │   │   │       └── queries.ts  # Notifications queries
    │   │   │   │   ├── offers/
    │   │   │   │   │   └── api/
    │   │   │   │   │       └── mutations.ts  # Offer create/respond
    │   │   │   │   ├── reviews/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── mutations.ts  # Review mutations
    │   │   │   │   │   │   └── queries.ts  # Review queries
    │   │   │   │   │   └── components/
    │   │   │   │   │       └── ReviewForm.tsx  # Review form (presentational)
    │   │   │   │   ├── search/
    │   │   │   │   │   └── api/
    │   │   │   │   │       ├── alerts.ts  # Search alerts API
    │   │   │   │   │       └── saved.ts  # Saved queries API
    │   │   │   │   ├── subscriptions/
    │   │   │   │   │   └── api/
    │   │   │   │   │       ├── mutations.ts  # Subscription mutations
    │   │   │   │   │       └── queries.ts  # Subscription queries
    │   │   │   │   └── support/
    │   │   │   │       └── api/
    │   │   │   │           ├── mutations.ts  # Support ticket mutations
    │   │   │   │           └── queries.ts  # Support ticket queries
    │   │   │   ├── hooks/  # FOLDER - Web-specific hooks only
    │   │   │   │   ├── use-form-validation.ts  # ❌ DELETE (duplicate)
    │   │   │   │   ├── use-keycloak.ts  # (web-specific)
    │   │   │   │   └── use-ssr-query.ts  # (web-specific SSR)
    │   │   │   ├── lib/
    │   │   │   │   └── api/  # WEB-SPECIFIC API CLIENTS
    │   │   │   │       ├── audit/
    │   │   │   │       │   ├── audit.ts  # Audit logs client
    │   │   │   │       │   ├── exports.ts  # BI export jobs client
    │   │   │   │       │   └── │
    │   │   │   │       ├── budgets/
    │   │   │   │       │   └── budgets.ts  # Budgets APIs
    │   │   │   │       ├── moderation/
    │   │   │   │       │   └── moderation.ts  # Moderation actions & reports
    │   │   │   │       ├── refunds/
    │   │   │   │       │   └── refunds.ts  # Refund cases & financial refunds
    │   │   │   │       ├── search/
    │   │   │   │       │   ├── personalization.ts  # Personalization API client
    │   │   │   │       │   ├── saved-search.ts  # Saved-search CRUD & alerts
    │   │   │   │       │   └── │
    │   │   │   │       ├── search-admin/
    │   │   │   │       │   ├── hygiene.ts  # Run hygiene jobs
    │   │   │   │       │   ├── query-logs.ts  # Query logs fetcher
    │   │   │   │       │   ├── search-admin.ts  # Synonyms/boosts/facets/rewrites/perf
    │   │   │   │       │   └── │
    │   │   │   │       ├── settings/
    │   │   │   │       │   └── notifications.ts  # Channels, email prefs, quiet hours
    │   │   │   │       ├── status/
    │   │   │   │       │   └── status.ts  # Status & incidents APIs
    │   │   │   │       ├── support/
    │   │   │   │       │   └── tickets.ts  # Tickets CRUD & messages
    │   │   │   │       └── │
    │   │   │   └── │
    │   │   └── middleware.ts  # Next.js middleware
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
    │   │   ├── 004-component-library.md
    │   │   ├── 005-mobile-admin-exclusion.md  # Title: Mobile Admin Portal Exclusion
    │   │   ├── 006-public-pages-mobile-strategy.md  # Title: Public/Marketing Pages Mobile Strategy
    │   │   ├── 007-component-platform-variants.md  # Title: Component Platform Variants Architecture
    │   │   ├── 008-hooks-domain-organization.md  # Title: Hooks Domain Organization
    │   │   └── │
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
    │   │   ├── cross-platform-routing.md  # Cross-Platform Routing Guide
    │   │   ├── deployment.md
    │   │   ├── development-workflow.md
    │   │   ├── development.md
    │   │   ├── error-boundary-patterns.md  # Error Boundary Implementation Patterns
    │   │   ├── getting-started.md
    │   │   ├── mobile-native-components.md  # Mobile Native Component Implementation Guide
    │   │   ├── testing-guide.md
    │   │   ├── testing.md
    │   │   ├── troubleshooting.md
    │   │   └── │
    │   ├── ARCHITECTURE.md  # System architecture
    │   ├── CONTRIBUTING.md  # Contribution guidelines
    │   ├── DEPLOYMENT.md  # Deployment procedures
    │   ├── MICROSERVICES_MAPPING.md  # BE microservices integration
    │   ├── PERFORMANCE.md  # Performance optimization guide
    │   ├── README.md  # Project overview
    │   ├── SETUP.md  # Development setup guide
    │   ├── STATE_MANAGEMENT.md  # TanStack Query + Zustand patterns
    │   ├── TESTING.md  # Testing strategy
    │   └── │
    ├── packages/  # Shared libraries
    │   ├── analytics/  # ⚠️ VERIFY EXISTENCE OR CREATE
    │   │   ├── src/
    │   │   │   ├── hooks/
    │   │   │   │   ├── use-analytics.ts  # Core analytics hook
    │   │   │   │   ├── use-page-view.ts  # Page view tracking
    │   │   │   │   ├── use-track-event.ts  # Custom event tracking
    │   │   │   │   └── │
    │   │   │   ├── providers/
    │   │   │   │   └── AnalyticsProvider.tsx  # Analytics provider
    │   │   │   ├── trackers/
    │   │   │   │   ├── google-analytics.ts  # Google Analytics
    │   │   │   │   ├── mixpanel.ts  # Mixpanel
    │   │   │   │   └── │
    │   │   │   ├── types/
    │   │   │   │   └── events.ts  # Event types
    │   │   │   ├── index.ts  # Package exports
    │   │   │   └── │
    │   │   ├── index.ts  # IF MISSING
    │   │   ├── package.json  # {
    │   │   ├── README.md  # Analytics docs
    │   │   ├── tsconfig.json  # Extends base tsconfig from packages/config/typescript-config
    │   │   ├── use-analytics.ts  # IF MISSING
    │   │   ├── use-earnings.ts  # IF MISSING
    │   │   ├── use-insights.ts  # IF MISSING
    │   │   ├── use-performance.ts  # IF MISSING
    │   │   └── │
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
    │   │   │   └── package.json
    │   │   └── typescript-config/
    │   │       ├── base.json  # Base TS config
    │   │       ├── nextjs.json  # Next.js TS config
    │   │       ├── package.json
    │   │       └── react-native.json  # React Native TS config
    │   ├── hooks/  # ALL shared hooks belong HERE
    │   │   ├── admin/  # ⚠️ CREATE IF MISSING - Admin hooks
    │   │   │   ├── index.ts  # export * from './use-audit-logs';
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-analytics.ts  # Hook: useAnalytics(scope, dateRange)
    │   │   │   ├── use-audit-logs.ts  # Hook: useAuditLogs(filters)
    │   │   │   ├── use-moderation.ts  # Hook: useModeration(contentType, contentId)
    │   │   │   ├── use-reports.ts  # Hook: useReports(reportType, params)
    │   │   │   ├── use-system-config.ts  # Hook: useSystemConfig()
    │   │   │   ├── use-system-health.ts  # Hook: useSystemHealth()
    │   │   │   ├── use-user-management.ts  # Hook: useUserManagement()
    │   │   │   ├── useAdminSession.ts  # JIT admin session hook
    │   │   │   ├── useAdminUsers.ts
    │   │   │   ├── useDisputes.ts
    │   │   │   ├── useKYCCases.ts
    │   │   │   ├── useKycCases.ts  # KYC cases management
    │   │   │   ├── useModeration.ts  # Moderation actions
    │   │   │   └── useSystemHealth.ts  # System health monitoring
    │   │   ├── analytics/  # ⚠️ VERIFY OR CREATE ENTIRE DOMAIN - Analytics hooks
    │   │   │   ├── index.ts  # Analytics hooks barrel
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-analytics-dashboard.ts
    │   │   │   ├── use-analytics-events.ts
    │   │   │   ├── use-analytics-report.ts
    │   │   │   ├── use-analytics.ts  # Analytics overview
    │   │   │   ├── use-earnings.ts  # Earnings
    │   │   │   ├── use-insights.ts  # Insights
    │   │   │   ├── use-performance-metrics.ts
    │   │   │   ├── use-performance.ts  # Performance KPIs
    │   │   │   ├── useAnalytics.ts  # Track events
    │   │   │   ├── useContractAnalytics.ts  # Contract analytics
    │   │   │   ├── useConversionTracking.ts  # Track conversions
    │   │   │   ├── useJobAnalytics.ts  # Job analytics
    │   │   │   ├── usePageView.ts  # Track page views
    │   │   │   ├── useProfileAnalytics.ts  # Profile analytics
    │   │   │   └── useRevenueAnalytics.ts  # Revenue analytics
    │   │   ├── api/  # - Generic API hooks
    │   │   │   ├── use-api.ts  # (shared)
    │   │   │   ├── use-mutation.ts  # (shared)
    │   │   │   └── use-query.ts  # (shared)
    │   │   ├── auctions/
    │   │   │   ├── useActiveAuctions.ts  # Active auctions list
    │   │   │   ├── useAuction.ts  # Single auction
    │   │   │   ├── useAuctionBid.ts  # Place bid
    │   │   │   └── useAuctionHistory.ts  # Bid history
    │   │   ├── auth/  # - Authentication hooks
    │   │   │   ├── use-auth.ts  # (shared)
    │   │   │   ├── use-permissions.ts  # (shared)
    │   │   │   ├── use-session.ts  # (shared)
    │   │   │   ├── useAuth.ts  # Main auth hook
    │   │   │   ├── useKeycloak.ts  # Keycloak integration
    │   │   │   ├── usePermissions.ts  # RBAC permissions
    │   │   │   └── useSession.ts  # Session management
    │   │   ├── bidding/
    │   │   │   ├── useBidAnalytics.ts  # Bid analytics
    │   │   │   ├── useBidHistory.ts  # Bid history
    │   │   │   ├── useBidStrategies.ts  # List strategies
    │   │   │   ├── useBidStrategy.ts  # Bid strategy management
    │   │   │   └── usePlaceBid.ts  # Place bid
    │   │   ├── certifications/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-add-certification.ts  # Add certification mutation
    │   │   │   ├── use-certification.ts  # Single certification hook
    │   │   │   ├── use-certifications.ts  # Certifications list hook
    │   │   │   └── use-verify-certification.ts  # Verify certification mutation
    │   │   ├── chat/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-conversation.ts
    │   │   │   ├── use-conversations.ts
    │   │   │   ├── use-message.ts
    │   │   │   ├── use-messages.ts
    │   │   │   └── use-typing-indicator.ts
    │   │   ├── collaboration/
    │   │   │   ├── useCollaboration.ts  # Collaboration session
    │   │   │   ├── useCursors.ts  # Cursor tracking
    │   │   │   ├── usePresence.ts  # User presence
    │   │   │   └── useSharedState.ts  # Shared state sync
    │   │   ├── compliance/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-compliance-documents.ts
    │   │   │   ├── use-tax-profile.ts
    │   │   │   ├── use-tax-reports.ts
    │   │   │   ├── useComplianceDocuments.ts  # Document management
    │   │   │   ├── useComplianceProfile.ts  # Compliance profile
    │   │   │   ├── useTaxProfile.ts  # Tax profile management
    │   │   │   └── useTaxReports.ts  # Tax reports
    │   │   ├── connects/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-connect-packages.ts
    │   │   │   ├── use-connect-refund.ts
    │   │   │   ├── use-connects.ts
    │   │   │   ├── use-purchase-connects.ts
    │   │   │   ├── useConnectPackages.ts  # Available packages
    │   │   │   ├── useConnectRefund.ts  # Request refund
    │   │   │   ├── useConnects.ts  # Connects balance and history
    │   │   │   └── usePurchaseConnects.ts  # Purchase connects
    │   │   ├── contests/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-contest-entries.ts  # Contest entries hook
    │   │   │   ├── use-contest-leaderboard.ts  # Leaderboard hook
    │   │   │   ├── use-contest.ts  # Single contest hook
    │   │   │   ├── use-contests.ts  # Contests list hook
    │   │   │   └── use-submit-entry.ts  # Submit entry mutation
    │   │   ├── contracts/  # THIS (domain hooks)
    │   │   │   ├── index.ts  # Export all contract hooks
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-amendments.ts  # Contract amendment operations
    │   │   │   ├── use-contract-actions.ts  # Contract actions
    │   │   │   ├── use-contract-amendments.ts  # Contract amendments
    │   │   │   ├── use-contract.ts  # Single contract operations
    │   │   │   ├── use-contracts.ts  # List contracts
    │   │   │   ├── use-deliverables.ts  # Deliverable management
    │   │   │   ├── use-dispute.ts  # Contract dispute
    │   │   │   ├── use-disputes.ts  # Dispute management
    │   │   │   ├── use-escrow.ts  # Escrow operations
    │   │   │   ├── use-invoices.ts  # Contract invoices
    │   │   │   ├── use-milestones.ts  # Milestone operations
    │   │   │   ├── use-payment-schedule.ts  # Payment schedule operations
    │   │   │   ├── use-time-tracking.ts  # Time tracking operations
    │   │   │   ├── use-timesheet.ts  # Timesheet operations
    │   │   │   ├── use-work-diary.ts  # Work diary operations
    │   │   │   ├── useContract.ts
    │   │   │   ├── useContracts.ts
    │   │   │   ├── useDeliverables.ts
    │   │   │   ├── useDisputes.ts
    │   │   │   ├── useMilestones.ts
    │   │   │   ├── useTimesheets.ts
    │   │   │   ├── useWorkDiary.ts
    │   │   │   └── │
    │   │   ├── deliverables/
    │   │   │   ├── useDeliverable.ts  # Single deliverable
    │   │   │   ├── useDeliverableRevisions.ts  # Revision management
    │   │   │   ├── useDeliverables.ts  # Deliverables list
    │   │   │   ├── useReviewDeliverable.ts  # Review deliverable (client)
    │   │   │   └── useUploadDeliverable.ts  # Upload deliverable
    │   │   ├── developer/
    │   │   │   ├── use-api-keys.ts  # API keys management
    │   │   │   ├── use-api-usage.ts  # API usage tracking
    │   │   │   ├── use-oauth-apps.ts  # OAuth apps management
    │   │   │   ├── use-webhooks.ts  # Webhooks management
    │   │   │   └── │
    │   │   ├── disputes/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-dispute-actions.ts
    │   │   │   ├── use-dispute-evidence.ts  # Dispute evidence hook
    │   │   │   ├── use-dispute-mediation.ts  # Dispute mediation hook
    │   │   │   ├── use-dispute-resolution.ts
    │   │   │   ├── use-dispute.ts  # Single dispute hook
    │   │   │   ├── use-disputes.ts  # Disputes list hook
    │   │   │   ├── use-file-dispute.ts  # File dispute mutation
    │   │   │   ├── useDispute.ts  # Single dispute
    │   │   │   ├── useDisputeEvidence.ts  # Evidence management
    │   │   │   ├── useDisputeResolution.ts  # Resolution actions
    │   │   │   └── useDisputes.ts  # Disputes list
    │   │   ├── events/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-create-event.ts  # Create event mutation
    │   │   │   ├── use-event-attendees.ts  # Event attendees hook
    │   │   │   ├── use-event-registration.ts  # Event registration mutation
    │   │   │   ├── use-event.ts  # Single event hook
    │   │   │   ├── use-events.ts  # Events list hook
    │   │   │   ├── usePresence.ts  # User presence tracking
    │   │   │   ├── useRealTimeUpdates.ts  # Real-time data sync
    │   │   │   ├── useTypingIndicator.ts  # Typing indicators
    │   │   │   └── useWebSocket.ts  # WebSocket connection
    │   │   ├── experiments/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-experiment-variant.ts
    │   │   │   ├── use-experiment.ts
    │   │   │   ├── use-experiments.ts
    │   │   │   ├── useExperiment.ts  # Experiment hook
    │   │   │   ├── useExperimentTracking.ts  # Track experiment events
    │   │   │   ├── useFeatureVariant.ts  # Feature variant hook
    │   │   │   └── useVariant.ts  # Variant hook
    │   │   ├── feature-flags/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-feature-flag.ts
    │   │   │   ├── use-feature-flags.ts
    │   │   │   ├── useFeatureFlag.ts  # Check single flag
    │   │   │   ├── useFeatureFlags.ts  # Get all flags
    │   │   │   └── useFeatureFlagVariant.ts  # A/B test variant
    │   │   ├── financial/  # THIS
    │   │   │   ├── index.ts  # Export all financial hooks
    │   │   │   ├── package.json  # Package config for @packages/hooks/financial
    │   │   │   ├── tsconfig.json  # TypeScript config extending base config
    │   │   │   ├── use-balance.ts  # Balance queries
    │   │   │   ├── use-billing-history.ts  # Billing history
    │   │   │   ├── use-bonus.ts  # Bonus
    │   │   │   ├── use-connect-packages.ts  # Connect packages
    │   │   │   ├── use-connects.ts  # Connects
    │   │   │   ├── use-earnings.ts
    │   │   │   ├── use-escrow.ts  # Escrow operations
    │   │   │   ├── use-invoice.ts
    │   │   │   ├── use-invoices.ts  # Invoice operations
    │   │   │   ├── use-payment-method-actions.ts  # Payment method actions
    │   │   │   ├── use-payment-method.ts  # Single payment method
    │   │   │   ├── use-payment-methods.ts  # Payment methods
    │   │   │   ├── use-payment.ts  # Payment
    │   │   │   ├── use-payout.ts  # Payout
    │   │   │   ├── use-payouts.ts  # Freelancer payouts
    │   │   │   ├── use-purchase-connects.ts  # Purchase connects
    │   │   │   ├── use-refund.ts  # Refund
    │   │   │   ├── use-refunds.ts  # Refund management
    │   │   │   ├── use-subscription.ts  # Subscription management
    │   │   │   ├── use-transaction.ts  # Transaction
    │   │   │   ├── use-transactions.ts  # Transaction history
    │   │   │   ├── use-wallet.ts  # Wallet operations
    │   │   │   ├── use-withdraw.ts  # Withdrawal operations
    │   │   │   ├── useEscrow.ts
    │   │   │   ├── useInvoices.ts
    │   │   │   ├── usePaymentMethods.ts
    │   │   │   ├── usePayoutMethods.ts
    │   │   │   ├── useTransactions.ts
    │   │   │   ├── useWallet.ts
    │   │   │   └── │
    │   │   ├── flags/
    │   │   │   ├── useFeatureFlag.ts  # Single flag hook
    │   │   │   ├── useFeatureFlags.ts  # Multiple flags
    │   │   │   └── useFeatureFlagVariant.ts  # A/B variant
    │   │   ├── forms/  # - Form validation hooks
    │   │   │   ├── use-form-state.ts  # (shared)
    │   │   │   ├── use-form.ts  # (shared)
    │   │   │   └── use-validation.ts  # (shared)
    │   │   ├── gamification/
    │   │   │   ├── useAchievements.ts  # Achievements
    │   │   │   ├── useBadges.ts  # Badges
    │   │   │   ├── useLeaderboard.ts  # Leaderboard
    │   │   │   └── usePoints.ts  # User points
    │   │   ├── geolocation/
    │   │   │   ├── useCountry.ts  # Country detection
    │   │   │   ├── useCountryDetection.ts  # Detect country
    │   │   │   ├── useDistanceCalculation.ts  # Calculate distance
    │   │   │   ├── useGeolocation.ts  # Get current location
    │   │   │   └── useTimezone.ts  # Detect timezone
    │   │   ├── groups/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-create-group.ts  # Create group mutation
    │   │   │   ├── use-group-members.ts  # Group members hook
    │   │   │   ├── use-group-posts.ts  # Group posts hook
    │   │   │   ├── use-group.ts  # Single group hook
    │   │   │   ├── use-groups.ts  # Groups list hook
    │   │   │   └── use-join-group.ts  # Join group mutation
    │   │   ├── i18n/
    │   │   │   ├── useCurrencyFormat.ts  # Currency formatting
    │   │   │   ├── useDateFormat.ts  # Date formatting
    │   │   │   ├── useNumberFormat.ts  # Number formatting
    │   │   │   ├── useRTL.ts  # RTL detection
    │   │   │   └── useTranslation.ts  # (already exists, enhanced)
    │   │   ├── interviews/
    │   │   │   ├── useInterview.ts  # Single interview
    │   │   │   ├── useInterviewFeedback.ts  # Interview feedback
    │   │   │   ├── useInterviews.ts  # Interviews list
    │   │   │   └── useScheduleInterview.ts  # Schedule interview
    │   │   ├── invitations/
    │   │   │   ├── useAcceptInvite.ts  # Accept invite (freelancer)
    │   │   │   ├── useDeclineInvite.ts  # Decline invite (freelancer)
    │   │   │   ├── useInvitationAnalytics.ts  # Invitation metrics
    │   │   │   ├── useInvitations.ts  # Invitations management
    │   │   │   └── useSendInvitation.ts  # Send invitation (client)
    │   │   ├── jobs/  # THIS
    │   │   │   ├── index.ts  # Export all job hooks
    │   │   │   ├── package.json  # Jobs hooks package
    │   │   │   ├── tsconfig.json  # TypeScript config
    │   │   │   ├── use-job-actions.ts  # Job actions
    │   │   │   ├── use-job-analytics.ts  # Job analytics
    │   │   │   ├── use-job-applicants.ts  # Applicants
    │   │   │   ├── use-job-application-actions.ts  # Application actions
    │   │   │   ├── use-job-application.ts  # Job application operations
    │   │   │   ├── use-job-applications.ts  # User applications
    │   │   │   ├── use-job-filters.ts  # Job search filters
    │   │   │   ├── use-job-invitations.ts  # send/accept/reject; BE: jobs-be/invitation
    │   │   │   ├── use-job-invites.ts
    │   │   │   ├── use-job-recommendations.ts  # Job recommendations
    │   │   │   ├── use-job-saved.ts  # Saved jobs
    │   │   │   ├── use-job-search.ts  # Job search functionality
    │   │   │   ├── use-job-stats.ts  # Job stats
    │   │   │   ├── use-job-templates.ts  # Job template operations
    │   │   │   ├── use-job.ts  # Single job operations
    │   │   │   ├── use-jobs.ts  # List/filter jobs
    │   │   │   ├── use-save-job.ts  # Save/unsave job
    │   │   │   ├── use-saved-jobs.ts  # Saved jobs list
    │   │   │   ├── use-similar-jobs.ts  # Similar jobs
    │   │   │   ├── use-trending-jobs.ts  # Trending jobs
    │   │   │   ├── useCreateJob.ts  # Create job mutation
    │   │   │   ├── useDeleteJob.ts  # Delete job mutation
    │   │   │   ├── useJob.ts  # Single job query
    │   │   │   ├── useJobFilters.ts  # Job filters state
    │   │   │   ├── useJobs.ts  # List jobs query
    │   │   │   ├── useSaveJob.ts  # Save/bookmark job
    │   │   │   ├── useUpdateJob.ts  # Update job mutation
    │   │   │   └── │
    │   │   ├── learning/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-achievement.ts
    │   │   │   ├── use-achievements.ts
    │   │   │   ├── use-certificate.ts  # Certificate hook
    │   │   │   ├── use-course.ts  # Single course hook
    │   │   │   ├── use-courses.ts  # Courses list hook
    │   │   │   ├── use-learning-path.ts  # Learning path hook
    │   │   │   ├── use-learning-paths.ts
    │   │   │   ├── use-learning-progress.ts
    │   │   │   ├── use-lesson-progress.ts  # Lesson progress hook
    │   │   │   ├── use-lesson.ts  # Single lesson hook
    │   │   │   ├── useAchievements.ts  # Achievements/badges
    │   │   │   ├── useLearningPath.ts  # Single learning path
    │   │   │   ├── useLearningPaths.ts  # Learning paths list
    │   │   │   ├── useLearningProgress.ts  # Track progress
    │   │   │   └── useMentorship.ts  # Mentorship management
    │   │   ├── live-updates/
    │   │   │   ├── useLiveDocument.ts  # Live document hook
    │   │   │   └── useLiveQuery.ts  # Live query hook
    │   │   ├── messages/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Export all message hooks
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-conversation.ts  # Single conversation operations
    │   │   │   ├── use-conversations.ts  # List conversations
    │   │   │   ├── use-message-actions.ts
    │   │   │   ├── use-message.ts  # Message operations
    │   │   │   ├── use-messages.ts  # Messages in conversation
    │   │   │   ├── use-real-time.ts  # Real-time messaging (WebSocket)
    │   │   │   ├── use-realtime-messages.ts
    │   │   │   ├── use-send-message.ts  # Send message with optimistic UI
    │   │   │   ├── use-typing-indicator.ts  # Typing indicator
    │   │   │   ├── useConversation.ts
    │   │   │   ├── useConversations.ts
    │   │   │   ├── useMessages.ts
    │   │   │   ├── useRealtimeMessages.ts  # WebSocket
    │   │   │   └── useSendMessage.ts
    │   │   ├── messaging/  # THIS
    │   │   │   ├── index.ts  # Export all messaging hooks
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-conversation.ts  # Single conversation ops
    │   │   │   ├── use-conversations.ts  # Conversations list
    │   │   │   ├── use-message-send.ts  # Send message operations
    │   │   │   ├── use-messages.ts  # Messages for conversation
    │   │   │   ├── use-real-time.ts  # Real-time messaging (WS)
    │   │   │   ├── use-send-message.ts  # Send message with optimistic UI
    │   │   │   └── use-typing-indicator.ts  # Typing indicators
    │   │   ├── moderation/
    │   │   │   ├── useBlockUser.ts  # Block/unblock users
    │   │   │   ├── useContentModeration.ts  # Check content
    │   │   │   ├── useContentStatus.ts  # Content status check
    │   │   │   ├── useContentValidation.ts  # Validation hook
    │   │   │   ├── useModerationActions.ts  # Moderation actions
    │   │   │   ├── useModerationStatus.ts  # Status hook
    │   │   │   ├── useReportContent.ts  # Report content
    │   │   │   └── useReporting.ts  # Report content
    │   │   ├── negotiations/
    │   │   │   ├── useNegotiation.ts  # Single negotiation
    │   │   │   ├── useNegotiationHistory.ts  # Negotiation history
    │   │   │   └── useNegotiationOffer.ts  # Make/accept/reject offer
    │   │   ├── networking/
    │   │   │   ├── useConnectionRequest.ts  # Send/accept/reject
    │   │   │   ├── useConnections.ts  # Connections management
    │   │   │   ├── useGroups.ts  # Groups management
    │   │   │   ├── useNetworkRecommendations.ts  # Connection recommendations
    │   │   │   └── useReferrals.ts  # Referral management
    │   │   ├── notifications/  # THIS
    │   │   │   ├── index.ts  # Export all notification hooks
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-mark-as-read.ts
    │   │   │   ├── use-mark-read.ts  # Mark notification as read
    │   │   │   ├── use-notification-actions.ts
    │   │   │   ├── use-notification-preferences.ts  # Notification preferences
    │   │   │   ├── use-notifications.ts  # Notifications list
    │   │   │   ├── use-realtime-notifications.ts
    │   │   │   ├── use-unread-count.ts
    │   │   │   ├── useMarkAsRead.ts
    │   │   │   ├── useNotifications.ts
    │   │   │   ├── useRealtimeNotifications.ts  # WebSocket
    │   │   │   └── useUnreadCount.ts
    │   │   ├── offline/
    │   │   │   ├── useNetworkStatus.ts  # Network status
    │   │   │   ├── useOfflineQueue.ts  # Offline action queue
    │   │   │   ├── useOfflineStorage.ts  # Local storage
    │   │   │   ├── useOfflineSync.ts  # Data synchronization
    │   │   │   ├── useOnlineStatus.ts  # Online status hook
    │   │   │   └── useSyncStatus.ts  # Sync status hook
    │   │   ├── organizations/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-budgets.ts
    │   │   │   ├── use-organization.ts
    │   │   │   ├── use-team-members.ts
    │   │   │   ├── use-vendors.ts
    │   │   │   ├── useBudgets.ts  # Budget management
    │   │   │   ├── useOrganization.ts  # Organization details
    │   │   │   ├── useTeamMembers.ts  # Team member management
    │   │   │   └── useVendors.ts  # Vendor management
    │   │   ├── packages/
    │   │   │   ├── useClickOutside.ts
    │   │   │   ├── useCopyToClipboard.ts
    │   │   │   ├── useDebounce.ts
    │   │   │   ├── useIntersectionObserver.ts
    │   │   │   ├── useLocalStorage.ts
    │   │   │   ├── useMediaQuery.ts
    │   │   │   └── useToggle.ts
    │   │   ├── payments/  # ENTIRE FOLDER - Payments domain hooks
    │   │   │   ├── index.ts  # Export all payment hooks
    │   │   │   ├── use-escrow.ts  # Escrow
    │   │   │   ├── use-invoices.ts  # Invoices
    │   │   │   ├── use-payment-methods.ts  # Payment methods CRUD
    │   │   │   ├── use-payment.ts  # Single payment
    │   │   │   ├── use-payments.ts  # Payments list
    │   │   │   ├── use-transactions.ts  # Transaction history
    │   │   │   └── use-wallet.ts  # Wallet operations
    │   │   ├── performance/
    │   │   │   ├── useAnalyticsTracking.ts  # Analytics events
    │   │   │   ├── useErrorTracking.ts  # Error tracking
    │   │   │   ├── usePerformanceMetrics.ts  # Web vitals tracking
    │   │   │   ├── usePerformanceMonitor.ts  # Monitor performance
    │   │   │   └── useWebVitals.ts  # Web Vitals
    │   │   ├── presence/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-last-seen.ts
    │   │   │   ├── use-online-status.ts
    │   │   │   ├── use-presence-subscription.ts
    │   │   │   ├── useLastSeen.ts  # Last seen time
    │   │   │   ├── useOnlineStatus.ts  # Online status
    │   │   │   ├── useOnlineUsers.ts  # Online users hook
    │   │   │   ├── usePresence.ts  # Presence hook
    │   │   │   └── usePresenceSubscription.ts  # Subscribe to presence
    │   │   ├── profile/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-education.ts
    │   │   │   ├── use-experience.ts
    │   │   │   ├── use-portfolio.ts
    │   │   │   ├── use-profile.ts
    │   │   │   ├── use-skills.ts
    │   │   │   ├── use-update-profile.ts
    │   │   │   ├── useEducation.ts
    │   │   │   ├── useExperience.ts
    │   │   │   ├── usePortfolio.ts
    │   │   │   ├── useProfile.ts
    │   │   │   ├── useServiceCatalog.ts
    │   │   │   ├── useSkills.ts
    │   │   │   └── useUpdateProfile.ts
    │   │   ├── proposals/  # THIS
    │   │   │   ├── index.ts  # Export all proposal hooks
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-bid-recommendations.ts  # Bid recommendations
    │   │   │   ├── use-bid.ts  # Bid/pricing operations
    │   │   │   ├── use-bidding.ts
    │   │   │   ├── use-interview.ts  # Interview scheduling
    │   │   │   ├── use-milestones.ts  # Proposal milestones
    │   │   │   ├── use-negotiation.ts  # Proposal negotiation
    │   │   │   ├── use-proposal-actions.ts  # Proposal actions
    │   │   │   ├── use-proposal-analytics.ts  # Proposal analytics
    │   │   │   ├── use-proposal-feedback.ts  # Proposal feedback
    │   │   │   ├── use-proposal-submission.ts  # Proposal submission flow
    │   │   │   ├── use-proposal-template-actions.ts  # Template actions
    │   │   │   ├── use-proposal-template.ts  # Single template
    │   │   │   ├── use-proposal-templates.ts  # Proposal templates
    │   │   │   ├── use-proposal-versions.ts  # Proposal versions
    │   │   │   ├── use-proposal.ts  # Single proposal operations
    │   │   │   ├── use-proposals.ts  # List proposals
    │   │   │   ├── use-rate-cards.ts  # Rate cards
    │   │   │   ├── use-revision.ts  # Proposal revisions
    │   │   │   ├── use-send-invitation.ts  # Send job invitation
    │   │   │   ├── use-submission.ts  # submit/validate/draft; BE: proposals-be/proposal
    │   │   │   ├── use-submit-proposal.ts
    │   │   │   ├── use-withdraw-proposal.ts
    │   │   │   ├── useBidding.ts  # Bidding hooks
    │   │   │   ├── useProposal.ts
    │   │   │   ├── useProposals.ts
    │   │   │   ├── useSubmitProposal.ts
    │   │   │   ├── useWithdrawProposal.ts
    │   │   │   └── │
    │   │   ├── realtime/
    │   │   │   ├── usePresence.ts  # User presence (online/offline)
    │   │   │   ├── useRealtimeAuction.ts  # Real-time auction updates
    │   │   │   ├── useRealtimeMessages.ts  # Real-time messages
    │   │   │   ├── useRealtimeNotifications.ts  # Real-time notifications
    │   │   │   ├── useTypingIndicator.ts  # Typing indicator
    │   │   │   └── useWebSocket.ts  # WebSocket connection
    │   │   ├── referrals/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-referral-code.ts  # Referral code hook
    │   │   │   ├── use-referral-earnings.ts  # Referral earnings hook
    │   │   │   ├── use-referral-stats.ts  # Referral stats hook
    │   │   │   ├── use-send-referral.ts  # Send referral mutation
    │   │   │   ├── useReferralCode.ts  # User referral code
    │   │   │   ├── useReferralProgram.ts  # Referral program details
    │   │   │   ├── useReferralStats.ts  # Referral statistics
    │   │   │   └── useRewards.ts  # Rewards management
    │   │   ├── reviews/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-dispute-review.ts  # Dispute review
    │   │   │   ├── use-review-actions.ts  # Review actions
    │   │   │   ├── use-review-appeal.ts  # Review appeal
    │   │   │   ├── use-review-draft.ts  # Review draft
    │   │   │   ├── use-review-flag.ts  # Review flag
    │   │   │   ├── use-review-helpful.ts  # Review helpful
    │   │   │   ├── use-review-moderation.ts  # Review moderation
    │   │   │   ├── use-review-reminders.ts  # Review reminders
    │   │   │   ├── use-review-reputation.ts  # Reputation
    │   │   │   ├── use-review-response.ts  # Review response
    │   │   │   ├── use-review-stats.ts  # Review stats
    │   │   │   ├── use-review-templates.ts  # Review templates
    │   │   │   ├── use-review.ts  # Single review
    │   │   │   ├── use-reviews.ts  # Reviews list
    │   │   │   ├── use-submit-review.ts  # Submit review
    │   │   │   ├── use-verify-review.ts  # Verify review eligibility
    │   │   │   ├── useBadges.ts
    │   │   │   ├── useCreateReview.ts
    │   │   │   ├── useReviews.ts
    │   │   │   ├── useReviewStats.ts
    │   │   │   └── │
    │   │   ├── search/  # ⚠️ VERIFY - Search hooks
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-autocomplete.ts  # Search autocomplete
    │   │   │   ├── use-facets.ts  # Search facets
    │   │   │   ├── use-filters.ts  # Search filters state
    │   │   │   ├── use-personalization.ts  # Search personalization
    │   │   │   ├── use-recent-searches.ts  # Recent searches
    │   │   │   ├── use-saved-search-actions.ts  # Saved search actions
    │   │   │   ├── use-saved-search.ts  # Single saved search
    │   │   │   ├── use-saved-searches.ts  # Hook: useSavedSearches()
    │   │   │   ├── use-search-alerts.ts  # Hook: useSearchAlerts()
    │   │   │   ├── use-search-contracts.ts  # Search contracts
    │   │   │   ├── use-search-facets.ts  # Search facets
    │   │   │   ├── use-search-filters.ts  # Search filters
    │   │   │   ├── use-search-freelancers.ts  # Search freelancers
    │   │   │   ├── use-search-history.ts  # Search history
    │   │   │   ├── use-search-jobs.ts  # Search jobs (alias)
    │   │   │   ├── use-search-params.ts  # Search params
    │   │   │   ├── use-search-personalization.ts  # Hook: useSearchPersonalization()
    │   │   │   ├── use-search-results.ts  # Search results
    │   │   │   ├── use-search-suggestions.ts  # Search suggestions
    │   │   │   ├── use-search.ts  # Core search
    │   │   │   ├── use-trending-searches.ts  # Trending searches
    │   │   │   ├── useFreelancerSearch.ts
    │   │   │   ├── useJobSearch.ts
    │   │   │   ├── useRecommendations.ts  # Recommendations
    │   │   │   ├── useSavedSearches.ts  # Saved searches
    │   │   │   ├── useSearch.ts  # Search execution
    │   │   │   ├── useSearchHistory.ts  # Search history
    │   │   │   ├── useSearchSuggestions.ts  # Auto-complete suggestions
    │   │   │   ├── useTrending.ts  # Trending items
    │   │   │   └── │
    │   │   ├── service-catalog/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-book-service.ts  # Book service mutation
    │   │   │   ├── use-create-service.ts  # Create service mutation
    │   │   │   ├── use-service-reviews.ts  # Service reviews hook
    │   │   │   ├── use-service.ts  # Single service hook
    │   │   │   └── use-services.ts  # Services list hook
    │   │   ├── settings/
    │   │   │   └── use-authorized-apps.ts  # Authorized apps management
    │   │   ├── shortlists/
    │   │   │   ├── useAddToShortlist.ts  # Add candidate
    │   │   │   ├── useRemoveFromShortlist.ts  # Remove candidate
    │   │   │   ├── useShortlist.ts  # Single shortlist
    │   │   │   └── useShortlists.ts  # Shortlists management
    │   │   ├── skills-tests/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-start-test.ts  # Start test mutation
    │   │   │   ├── use-submit-test.ts  # Submit test mutation
    │   │   │   ├── use-test-results.ts  # Test results hook
    │   │   │   ├── use-test.ts  # Single test hook
    │   │   │   └── use-tests.ts  # Tests list hook
    │   │   ├── storage/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-file-download.ts
    │   │   │   ├── use-file-upload.ts  # File upload
    │   │   │   ├── use-files.ts  # File management
    │   │   │   ├── use-presigned-url.ts
    │   │   │   ├── use-storage.ts  # Storage
    │   │   │   ├── use-upload.ts
    │   │   │   ├── useFileDownload.ts
    │   │   │   ├── usePresignedUrl.ts
    │   │   │   ├── useUpload.ts
    │   │   │   └── │
    │   │   ├── subscriptions/  # ⚠️ VERIFY - Subscription hooks
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-addon.ts  # Subscription addon
    │   │   │   ├── use-entitlements.ts
    │   │   │   ├── use-plan-comparison.ts  # Plan comparison
    │   │   │   ├── use-plan.ts  # Subscription plan
    │   │   │   ├── use-plans.ts
    │   │   │   ├── use-promo.ts  # Promo code
    │   │   │   ├── use-subscription-actions.ts  # Hook: useSubscriptionActions()
    │   │   │   ├── use-subscription-billing.ts  # Subscription billing
    │   │   │   ├── use-subscription-features.ts  # Feature entitlements
    │   │   │   ├── use-subscription-invoices.ts  # Hook: useSubscriptionInvoices()
    │   │   │   ├── use-subscription-usage.ts  # Subscription usage
    │   │   │   ├── use-subscription.ts  # Hook: useSubscription()
    │   │   │   ├── use-trial.ts  # Trial management
    │   │   │   ├── use-upgrade.ts
    │   │   │   ├── use-usage.ts
    │   │   │   ├── useConnects.ts
    │   │   │   ├── useEntitlements.ts  # Feature entitlements
    │   │   │   ├── usePlans.ts  # Subscription plans
    │   │   │   ├── useSubscription.ts  # Current subscription
    │   │   │   ├── useUpgrade.ts
    │   │   │   ├── useUsage.ts  # Usage metrics
    │   │   │   └── │
    │   │   ├── support/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts  # export * from './use-support-ticket';
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-support-categories.ts  # Support categories
    │   │   │   ├── use-support-ticket.ts  # Hook: useSupportTicket(ticketId)
    │   │   │   ├── use-support-tickets.ts  # Hook: useSupportTickets(filters)
    │   │   │   ├── use-ticket-actions.ts  # Hook: useTicketActions(ticketId)
    │   │   │   ├── use-ticket-messages.ts  # Ticket messages
    │   │   │   ├── use-ticket.ts  # Ticket
    │   │   │   ├── use-tickets.ts  # Support tickets list
    │   │   │   └── │
    │   │   ├── talent-cloud/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-agencies.ts  # Agencies list hook
    │   │   │   ├── use-agency-projects.ts  # Agency projects hook
    │   │   │   ├── use-agency.ts  # Single agency hook
    │   │   │   ├── use-team.ts  # Single team hook
    │   │   │   └── use-teams.ts  # Teams list hook
    │   │   ├── tax/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-tax-documents.ts
    │   │   │   ├── use-tax-profile.ts
    │   │   │   ├── use-tax-reports.ts
    │   │   │   ├── use-tax-withholding.ts
    │   │   │   ├── useTaxForms.ts  # Tax forms management
    │   │   │   ├── useTaxReports.ts  # Tax reports
    │   │   │   └── useTaxSettings.ts  # Tax settings
    │   │   ├── timesheets/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-log-time.ts  # Log time mutation
    │   │   │   ├── use-manual-time.ts  # Manual time entry mutation
    │   │   │   ├── use-timesheet.ts  # Single timesheet hook
    │   │   │   ├── use-timesheets.ts  # Timesheets list hook
    │   │   │   └── use-work-diary.ts  # Work diary hook
    │   │   ├── typing-indicators/
    │   │   │   └── useTypingIndicator.ts  # Typing hook
    │   │   ├── ui/
    │   │   ├── users/  # 6 HOOKS - Users domain
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-availability.ts  # User availability
    │   │   │   ├── use-certifications.ts  # User certifications
    │   │   │   ├── use-education.ts  # Education history
    │   │   │   ├── use-experience.ts  # Work experience
    │   │   │   ├── use-organization.ts  # Organization
    │   │   │   ├── use-portfolio.ts  # Portfolio items
    │   │   │   ├── use-profile-actions.ts  # Profile actions
    │   │   │   ├── use-profile.ts  # User profile
    │   │   │   ├── use-skills.ts  # User skills
    │   │   │   ├── use-team.ts  # Team
    │   │   │   ├── use-update-profile.ts  # Update profile
    │   │   │   ├── use-user-badge.ts  # User badge
    │   │   │   ├── use-verification.ts  # Verification
    │   │   │   ├── use-wallet-actions.ts  # Wallet actions
    │   │   │   └── │
    │   │   ├── video/
    │   │   │   ├── useRecording.ts  # Call recording
    │   │   │   ├── useScreenShare.ts  # Screen sharing
    │   │   │   ├── useVideoCall.ts  # Video call management
    │   │   │   └── useVideoDevices.ts  # Device selection
    │   │   ├── webhooks/  # ENTIRE DOMAIN
    │   │   │   ├── index.ts
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-webhook-logs.ts
    │   │   │   ├── use-webhook-test.ts
    │   │   │   ├── use-webhook.ts
    │   │   │   ├── use-webhooks.ts  # ... CONTINUE WITH OTHER DOMAINS ...
    │   │   │   ├── useWebhook.ts  # Single webhook
    │   │   │   ├── useWebhookLogs.ts  # Webhook delivery logs
    │   │   │   ├── useWebhooks.ts  # List webhooks
    │   │   │   └── useWebhookTest.ts  # Test webhook
    │   │   ├── work-tracking/
    │   │   │   ├── useApproveTimesheet.ts  # Approve timesheet (client)
    │   │   │   ├── useTimesheet.ts  # Timesheet management
    │   │   │   ├── useTimeTracking.ts  # Real-time time tracking
    │   │   │   └── useWorkDiary.ts  # Work diary entries
    │   │   └── │
    │   ├── shared/  # Business logic, hooks, utilities
    │   │   ├── components/
    │   │   │   ├── Auth/
    │   │   │   │   └── BiometricButton/
    │   │   │   │       ├── BiometricButton.native.tsx  # Biometric auth (native)
    │   │   │   │       ├── BiometricButton.web.tsx  # Web Authn fallback
    │   │   │   │       └── │
    │   │   │   ├── FileUpload/
    │   │   │   │   ├── DocumentScanner/
    │   │   │   │   │   └── DocumentScanner.native.tsx  # Document scanner
    │   │   │   │   ├── ImageCropper/
    │   │   │   │   │   └── ImageCropper.native.tsx  # Image cropper
    │   │   │   │   └── │
    │   │   │   ├── Offline/
    │   │   │   │   ├── OfflineIndicator/
    │   │   │   │   │   └── OfflineIndicator.native.tsx  # Offline banner
    │   │   │   │   ├── SyncStatus/
    │   │   │   │   │   └── SyncStatus.native.tsx  # Sync progress
    │   │   │   │   └── │
    │   │   │   ├── QRCode/
    │   │   │   │   ├── QRCodeGenerator/
    │   │   │   │   │   └── QRCodeGenerator.native.tsx  # QR generator
    │   │   │   │   ├── QRCodeScanner/
    │   │   │   │   │   └── QRCodeScanner.native.tsx  # QR scanner
    │   │   │   │   └── │
    │   │   │   └── │
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
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── communications/
    │   │   │   │   │   ├── client.ts  # Communications API client
    │   │   │   │   │   ├── conversations.ts  # /v1/conversations, /v1/messages
    │   │   │   │   │   └── notifications.ts  # /v1/notifications
    │   │   │   │   ├── compliance/
    │   │   │   │   │   ├── compliance-client.ts  # Compliance API client
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── contracts/
    │   │   │   │   │   ├── client.ts  # Contracts API client
    │   │   │   │   │   ├── contracts.ts  # /v1/contracts
    │   │   │   │   │   ├── deliverables.ts  # /v1/contracts/{id}/deliverables
    │   │   │   │   │   ├── diary.ts  # /v1/contracts/{id}/work-diary|timesheets
    │   │   │   │   │   ├── disputes.ts  # /v1/contracts/{id}/disputes
    │   │   │   │   │   ├── escrow.ts  # /v1/escrow/{contractId}/fund|release
    │   │   │   │   │   ├── milestones.ts  # /v1/contracts/{id}/milestones
    │   │   │   │   │   └── offers.ts  # /v1/offers (proposed)
    │   │   │   │   ├── experiments/
    │   │   │   │   │   ├── experiments-client.ts  # Experiments API client
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── financial/
    │   │   │   │   │   ├── client.ts  # Financial API client
    │   │   │   │   │   ├── invoices.ts  # Invoices endpoints
    │   │   │   │   │   ├── payouts.ts  # Payouts endpoints
    │   │   │   │   │   ├── refunds.ts  # Refunds endpoints
    │   │   │   │   │   ├── subscription.ts  # Subscription endpoints
    │   │   │   │   │   └── wallet.ts  # Wallet endpoints
    │   │   │   │   ├── flags/
    │   │   │   │   │   ├── flags-client.ts  # Feature flags API client
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── incidents/
    │   │   │   │   │   ├── incidents-client.ts  # Incidents API client
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── jobs/
    │   │   │   │   │   ├── client.ts  # Jobs API client
    │   │   │   │   │   ├── invites.ts  # /v1/invites
    │   │   │   │   │   └── jobs.ts  # /v1/jobs + subroutes
    │   │   │   │   ├── learning/
    │   │   │   │   │   ├── learning-client.ts  # Learning API client
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── moderation/
    │   │   │   │   │   ├── moderation-client.ts  # Moderation API client
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── presence/
    │   │   │   │   │   ├── presence-client.ts  # Presence API client
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── proposals/
    │   │   │   │   │   ├── client.ts  # Proposals API client
    │   │   │   │   │   ├── interviews.ts  # Interviews scheduling
    │   │   │   │   │   └── proposals.ts  # Proposals CRUD
    │   │   │   │   ├── search/
    │   │   │   │   │   ├── alerts.ts  # /v1/search/alerts (proposed)
    │   │   │   │   │   ├── client.ts  # Search API client
    │   │   │   │   │   ├── feed.ts  # /v1/feed|trending|similarity
    │   │   │   │   │   ├── portfolios.ts  # /v1/search/portfolios
    │   │   │   │   │   ├── query.ts  # /v1/search
    │   │   │   │   │   └── saved.ts  # /v1/search/saved (proposed)
    │   │   │   │   ├── sourcing/
    │   │   │   │   │   ├── sourcing-client.ts  # Sourcing API client
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── storage/
    │   │   │   │   │   ├── assets.ts  # /v1/storage/uploads|commit|assets
    │   │   │   │   │   └── client.ts  # Storage API client
    │   │   │   │   ├── users/
    │   │   │   │   │   ├── client.ts  # Users API client
    │   │   │   │   │   ├── org.ts  # /v1/orgs, /v1/invites
    │   │   │   │   │   ├── profile.ts  # /v1/profile, /v1/users/{handle}
    │   │   │   │   │   ├── saved.ts  # /v1/saved-items
    │   │   │   │   │   └── security.ts  # /v1/sessions, /v1/mfa
    │   │   │   │   └── webhooks/
    │   │   │   │       ├── types.ts
    │   │   │   │       └── webhooks-client.ts  # Webhooks API client
    │   │   │   ├── compliance/
    │   │   │   │   ├── compliance-client.ts  # Compliance API client
    │   │   │   │   └── types.ts
    │   │   │   ├── experiments/
    │   │   │   │   ├── ab-testing/
    │   │   │   │   │   ├── bucketing.ts  # User bucketing
    │   │   │   │   │   ├── experiment-engine.ts  # A/B test engine
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
    │   │   │   │   ├── experiments-client.ts  # Experiments API client
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
    │   │   │   │   │   │   ├── analytics-api.ts  # Admin analytics API
    │   │   │   │   │   │   ├── kyc-api.ts  # KYC management API
    │   │   │   │   │   │   └── moderation-api.ts  # Moderation API
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── admin-mutations.ts  # Admin mutations
    │   │   │   │   │   │   └── admin-queries.ts  # BE: admin-be
    │   │   │   │   │   ├── admin-api.ts
    │   │   │   │   │   └── types.ts  # Admin types
    │   │   │   │   ├── analytics/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── analytics-api.ts  # Analytics API
    │   │   │   │   │   ├── events/
    │   │   │   │   │   │   ├── contract-events.ts  # Contract events
    │   │   │   │   │   │   ├── job-events.ts  # Job events
    │   │   │   │   │   │   ├── payment-events.ts  # Payment events
    │   │   │   │   │   │   ├── proposal-events.ts  # Proposal events
    │   │   │   │   │   │   ├── system-events.ts  # System events
    │   │   │   │   │   │   └── user-events.ts  # User events
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
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── auctions-mutations.ts  # Auction mutations
    │   │   │   │   │   │   └── auctions-queries.ts  # Auction queries
    │   │   │   │   │   ├── store/
    │   │   │   │   │   │   └── auction-store.ts  # Real-time auction state (Zustand)
    │   │   │   │   │   └── types.ts  # Auction types
    │   │   │   │   ├── auth/  # Authentication
    │   │   │   │   │   ├── stores/
    │   │   │   │   │   │   └── auth-store.ts  # Zustand auth store
    │   │   │   │   │   ├── utils/
    │   │   │   │   │   │   ├── rbac.ts  # RBAC utilities
    │   │   │   │   │   │   ├── session.ts  # Session utilities
    │   │   │   │   │   │   └── token.ts  # Token utilities
    │   │   │   │   │   └── types.ts  # Auth types
    │   │   │   │   ├── bidding/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── bid-api.ts  # Bid placement API
    │   │   │   │   │   │   └── bid-strategy-api.ts  # Bid strategy API
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
    │   │   │   │   │   ├── providers/
    │   │   │   │   │   │   └── CollaborationProvider.tsx  # Collab context
    │   │   │   │   │   └── types.ts  # Collaboration types
    │   │   │   │   ├── compliance/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── compliance-api.ts  # Compliance API
    │   │   │   │   │   │   └── tax-profile-api.ts  # Tax profile API
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── compliance-mutations.ts  # Compliance mutations
    │   │   │   │   │   │   └── compliance-queries.ts  # Compliance queries
    │   │   │   │   │   └── types.ts  # Compliance types
    │   │   │   │   ├── connects/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── connects-api.ts  # Connects API client
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── connects-mutations.ts  # Connect mutations
    │   │   │   │   │   │   └── connects-queries.ts  # Connect queries
    │   │   │   │   │   └── types.ts  # Connect types
    │   │   │   │   ├── contracts/  # Contracts feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── contracts-api.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── contracts-mutations.ts
    │   │   │   │   │   │   └── contracts-queries.ts  # BE: contracts-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── deliverables/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── deliverables-api.ts  # Deliverables API client
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── deliverables-mutations.ts  # Deliverable mutations
    │   │   │   │   │   │   └── deliverables-queries.ts  # Deliverable queries
    │   │   │   │   │   └── types.ts  # Deliverable types
    │   │   │   │   ├── disputes/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── disputes-api.ts  # Disputes API client
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── disputes-mutations.ts  # Dispute mutations
    │   │   │   │   │   │   └── disputes-queries.ts  # Dispute queries
    │   │   │   │   │   └── types.ts  # Dispute types
    │   │   │   │   ├── events/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── events-api.ts  # Events API client
    │   │   │   │   │   ├── providers/
    │   │   │   │   │   │   └── WebSocketProvider.tsx  # WebSocket context
    │   │   │   │   │   └── types.ts  # Event types
    │   │   │   │   ├── experiments/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── experiments-api.ts  # Experiments API
    │   │   │   │   │   ├── providers/
    │   │   │   │   │   │   ├── ExperimentProvider.tsx  # Experiment context
    │   │   │   │   │   │   └── ExperimentsProvider.tsx  # Experiments context
    │   │   │   │   │   ├── utils/
    │   │   │   │   │   │   ├── experiment-context.ts  # Experiment context
    │   │   │   │   │   │   ├── experiment-storage.ts  # Store assignments
    │   │   │   │   │   │   ├── tracking.ts  # Track events
    │   │   │   │   │   │   └── variant-assignment.ts  # Assign variant
    │   │   │   │   │   └── types.ts  # Experiment types
    │   │   │   │   ├── feature-flags/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── flags-api.ts  # Feature flags API
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   └── flags-queries.ts  # Flag queries
    │   │   │   │   │   ├── store/
    │   │   │   │   │   │   └── flags-store.ts  # Flags state (Zustand)
    │   │   │   │   │   └── types.ts  # Flag types
    │   │   │   │   ├── financial/  # Financial feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── financial-api.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── financial-mutations.ts
    │   │   │   │   │   │   └── financial-queries.ts  # BE: financial-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── flags/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── flags-api.ts  # Feature flags API
    │   │   │   │   │   ├── providers/
    │   │   │   │   │   │   └── FeatureFlagsProvider.tsx  # Context provider
    │   │   │   │   │   ├── types.ts  # Flag types
    │   │   │   │   │   └── utils.ts  # Flag utilities
    │   │   │   │   ├── gamification/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── achievements-api.ts  # Achievements API
    │   │   │   │   │   │   ├── badges-api.ts  # Badges API
    │   │   │   │   │   │   ├── gamification-api.ts  # Gamification API
    │   │   │   │   │   │   └── leaderboards-api.ts  # Leaderboards API
    │   │   │   │   │   ├── components/
    │   │   │   │   │   │   ├── AchievementToast.tsx  # Achievement notification
    │   │   │   │   │   │   ├── BadgeCollection.tsx  # Badge collection
    │   │   │   │   │   │   ├── LeaderboardWidget.tsx  # Leaderboard widget
    │   │   │   │   │   │   └── PointsDisplay.tsx  # Points display
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── gamification-mutations.ts  # Gamification mutations
    │   │   │   │   │   │   └── gamification-queries.ts  # Gamification queries
    │   │   │   │   │   └── types.ts  # Gamification types
    │   │   │   │   ├── geolocation/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── geolocation-api.ts  # Geolocation API
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
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── interviews-mutations.ts  # Interview mutations
    │   │   │   │   │   │   └── interviews-queries.ts  # Interview queries
    │   │   │   │   │   └── types.ts  # Interview types
    │   │   │   │   ├── invitations/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── job-invitations-api.ts  # Job invitations API (client)
    │   │   │   │   │   │   └── proposal-invites-api.ts  # Proposal invites API (freelancer)
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── invitations-mutations.ts  # Invitation mutations
    │   │   │   │   │   │   └── invitations-queries.ts  # Invitation queries
    │   │   │   │   │   └── types.ts  # Invitation types
    │   │   │   │   ├── jobs/  # Jobs feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── jobs-api.ts  # Jobs API client
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── jobs-mutations.ts  # TanStack Query mutations
    │   │   │   │   │   │   └── jobs-queries.ts  # TanStack Query queries
    │   │   │   │   │   ├── utils/
    │   │   │   │   │   │   ├── job-helpers.ts  # Job utility functions
    │   │   │   │   │   │   └── job-validation.ts  # Job validation (Zod)
    │   │   │   │   │   └── types.ts  # Jobs types
    │   │   │   │   ├── learning/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── learning-paths-api.ts  # Learning paths API
    │   │   │   │   │   │   └── mentorship-api.ts  # Mentorship API
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── learning-mutations.ts  # Learning mutations
    │   │   │   │   │   │   └── learning-queries.ts  # Learning queries
    │   │   │   │   │   └── types.ts  # Learning types
    │   │   │   │   ├── messages/  # Messaging feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── messages-api.ts
    │   │   │   │   │   │   └── websocket.ts  # WebSocket client
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── messages-mutations.ts
    │   │   │   │   │   │   └── messages-queries.ts  # BE: communications-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── moderation/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── moderation-api.ts  # Moderation API
    │   │   │   │   │   ├── utils/
    │   │   │   │   │   │   ├── content-validator.ts  # Content validation
    │   │   │   │   │   │   └── profanity-filter.ts  # Profanity filtering (client-side)
    │   │   │   │   │   ├── types.ts  # Moderation types
    │   │   │   │   │   └── utils.ts  # Content validation
    │   │   │   │   ├── negotiations/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── negotiations-api.ts  # Negotiations API client
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── negotiations-mutations.ts  # Negotiation mutations
    │   │   │   │   │   │   └── negotiations-queries.ts  # Negotiation queries
    │   │   │   │   │   └── types.ts  # Negotiation types
    │   │   │   │   ├── networking/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── connections-api.ts  # Connections API
    │   │   │   │   │   │   ├── groups-api.ts  # Groups API
    │   │   │   │   │   │   └── referrals-api.ts  # Referrals API
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── networking-mutations.ts  # Networking mutations
    │   │   │   │   │   │   └── networking-queries.ts  # Networking queries
    │   │   │   │   │   └── types.ts  # Networking types
    │   │   │   │   ├── notifications/  # Notifications feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── notifications-api.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── notifications-mutations.ts
    │   │   │   │   │   │   └── notifications-queries.ts  # BE: communications-be
    │   │   │   │   │   ├── store.ts  # Notifications store (Zustand)
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── offline/
    │   │   │   │   │   ├── store/
    │   │   │   │   │   │   ├── offline-store.ts  # Offline state management
    │   │   │   │   │   │   └── sync-store.ts  # Sync state management
    │   │   │   │   │   └── types.ts  # Offline types
    │   │   │   │   ├── organizations/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── budgets-api.ts  # Budget management API
    │   │   │   │   │   │   ├── organizations-api.ts  # Organizations API
    │   │   │   │   │   │   └── vendors-api.ts  # Vendor management API
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── organizations-mutations.ts  # Org mutations
    │   │   │   │   │   │   └── organizations-queries.ts  # Org queries
    │   │   │   │   │   └── types.ts  # Organization types
    │   │   │   │   ├── performance/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── performance-api.ts  # Performance API
    │   │   │   │   │   ├── utils/
    │   │   │   │   │   │   ├── error-reporter.ts  # Error reporter
    │   │   │   │   │   │   ├── performance-observer.ts  # Performance observer
    │   │   │   │   │   │   └── web-vitals-reporter.ts  # Web Vitals reporter
    │   │   │   │   │   └── types.ts  # Performance types
    │   │   │   │   ├── presence/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── presence-api.ts  # Presence API
    │   │   │   │   │   └── types.ts  # Presence types
    │   │   │   │   ├── profile/  # Profile feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── profile-api.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── profile-mutations.ts
    │   │   │   │   │   │   └── profile-queries.ts  # BE: users-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── proposals/  # Proposals feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── proposals-api.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── proposals-mutations.ts
    │   │   │   │   │   │   └── proposals-queries.ts  # BE: proposals-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── realtime/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── websocket-client.ts  # WebSocket client
    │   │   │   │   │   ├── providers/
    │   │   │   │   │   │   └── WebSocketProvider.tsx  # WebSocket context
    │   │   │   │   │   ├── store/
    │   │   │   │   │   │   └── realtime-store.ts  # Real-time state (Zustand)
    │   │   │   │   │   ├── websocket/
    │   │   │   │   │   │   ├── client.ts  # WebSocket client
    │   │   │   │   │   │   ├── heartbeat.ts  # Connection health
    │   │   │   │   │   │   └── reconnection.ts  # Reconnection logic
    │   │   │   │   │   ├── types.ts  # Real-time types
    │   │   │   │   │   └── utils.ts  # Connection management
    │   │   │   │   ├── referrals/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── referrals-api.ts  # Referrals API (enhanced)
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── referrals-mutations.ts  # Referral mutations
    │   │   │   │   │   │   └── referrals-queries.ts  # Referral queries
    │   │   │   │   │   └── types.ts  # Referral types
    │   │   │   │   ├── reviews/  # Reviews feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── reviews-api.ts
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
    │   │   │   │   │   │   ├── saved-searches-api.ts  # Saved searches API
    │   │   │   │   │   │   ├── search-api.ts  # Search API (already may exist, ensuring completeness)
    │   │   │   │   │   │   └── trending-api.ts  # Trending API
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── search-mutations.ts  # Search mutations
    │   │   │   │   │   │   └── search-queries.ts  # BE: search-be
    │   │   │   │   │   ├── store/
    │   │   │   │   │   │   └── search-store.ts  # Search UI state (filters, etc.)
    │   │   │   │   │   └── types.ts  # Search types
    │   │   │   │   ├── shortlists/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── shortlists-api.ts  # Shortlists API
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── shortlists-mutations.ts  # Shortlist mutations
    │   │   │   │   │   │   └── shortlists-queries.ts  # Shortlist queries
    │   │   │   │   │   └── types.ts  # Shortlist types
    │   │   │   │   ├── storage/  # File storage feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── storage-api.ts  # BE: storage-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── subscriptions/  # Subscriptions feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── subscriptions-api.ts  # Subscriptions API
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── subscriptions-mutations.ts  # Subscription mutations
    │   │   │   │   │   │   └── subscriptions-queries.ts  # BE: subscriptions-be
    │   │   │   │   │   ├── query-keys.ts  # ['subscription','plans'|'me']
    │   │   │   │   │   ├── store.ts  # Local UI state for plans/checkout
    │   │   │   │   │   └── types.ts  # Subscription types
    │   │   │   │   ├── tax/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── tax-api.ts  # Tax API client
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── tax-mutations.ts  # Tax mutations
    │   │   │   │   │   │   └── tax-queries.ts  # Tax queries
    │   │   │   │   │   └── types.ts  # Tax types
    │   │   │   │   ├── video/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── video-api.ts  # Video call API
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
    │   │   │   │   │   └── types.ts  # Video types
    │   │   │   │   ├── webhooks/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── webhooks-api.ts  # Webhooks API client
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
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── webhooks-mutations.ts  # Webhook mutations
    │   │   │   │   │   │   └── webhooks-queries.ts  # Webhook queries
    │   │   │   │   │   └── types.ts  # Webhook types
    │   │   │   │   └── work-tracking/
    │   │   │   │       ├── api/
    │   │   │   │       │   ├── timesheet-api.ts  # Timesheet API
    │   │   │   │       │   └── work-diary-api.ts  # Work diary API
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
    │   │   │   │   └── types.ts
    │   │   │   ├── gamification/
    │   │   │   │   ├── achievements/
    │   │   │   │   │   ├── achievement-engine.ts  # Achievement system
    │   │   │   │   │   ├── achievement-notifier.ts  # Achievement notifications
    │   │   │   │   │   ├── achievement-tracker.ts  # Progress tracking
    │   │   │   │   │   └── types.ts  # Achievement types
    │   │   │   │   ├── badges/
    │   │   │   │   │   ├── badge-criteria.ts  # Badge criteria
    │   │   │   │   │   ├── badge-display.ts  # Badge rendering
    │   │   │   │   │   └── badge-system.ts  # Badge management
    │   │   │   │   ├── leaderboards/
    │   │   │   │   │   ├── leaderboard-engine.ts  # Leaderboard system
    │   │   │   │   │   ├── ranking-algorithms.ts  # Ranking logic
    │   │   │   │   │   └── time-windows.ts  # Time-based rankings
    │   │   │   │   └── points/
    │   │   │   │       ├── earning-rules.ts  # Point earning rules
    │   │   │   │       ├── points-system.ts  # Points management
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
    │   │   │   │   ├── ip-detection/
    │   │   │   │   │   ├── cache.ts  # Location cache
    │   │   │   │   │   └── ip-resolver.ts  # IP-based location
    │   │   │   │   └── timezone/
    │   │   │   │       ├── dst-handler.ts  # Daylight saving handling
    │   │   │   │       ├── timezone-converter.ts  # Timezone conversion
    │   │   │   │       └── timezone-detector.ts  # Timezone detection
    │   │   │   ├── i18n/  # Internationalization
    │   │   │   │   ├── locales/  # Locale files
    │   │   │   │   │   └── en/
    │   │   │   │   │       ├── admin.json
    │   │   │   │   │       ├── auth.json
    │   │   │   │   │       ├── common.json
    │   │   │   │   │       ├── contracts.json
    │   │   │   │   │       ├── errors.json
    │   │   │   │   │       ├── financial.json
    │   │   │   │   │       ├── jobs.json
    │   │   │   │   │       ├── messages.json
    │   │   │   │   │       ├── profile.json
    │   │   │   │   │       ├── proposals.json
    │   │   │   │   │       ├── reviews.json
    │   │   │   │   │       ├── settings.json
    │   │   │   │   │       └── subscription.json
    │   │   │   │   └── config.ts  # i18n configuration
    │   │   │   ├── incidents/
    │   │   │   │   ├── incidents-client.ts  # Incidents API client
    │   │   │   │   └── types.ts
    │   │   │   ├── learning/
    │   │   │   │   ├── learning-client.ts  # Learning API client
    │   │   │   │   └── types.ts
    │   │   │   ├── lib/  # Shared utilities
    │   │   │   │   ├── api/
    │   │   │   │   │   ├── client.ts  # Axios/Fetch client setup
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
    │   │   │   │   ├── reporting/
    │   │   │   │   │   ├── evidence-collection.ts  # Evidence gathering
    │   │   │   │   │   ├── report-submission.ts  # Report submission
    │   │   │   │   │   └── report-types.ts  # Report categories
    │   │   │   │   ├── moderation-client.ts  # Moderation API client
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
    │   │   │   │   │   ├── profiling.ts  # Performance profiling
    │   │   │   │   │   ├── web-vitals-reporter.ts  # Web Vitals reporting
    │   │   │   │   │   └── web-vitals.ts  # Web Vitals monitoring
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
    │   │   │   │   └── types.ts
    │   │   │   ├── realtime/
    │   │   │   │   ├── live-updates/
    │   │   │   │   │   ├── event-stream.ts  # Server-sent events
    │   │   │   │   │   └── polling-fallback.ts  # Polling fallback
    │   │   │   │   ├── presence/
    │   │   │   │   │   ├── presence-tracker.ts  # User presence tracking
    │   │   │   │   │   └── types.ts  # Presence types
    │   │   │   │   ├── typing-indicators/
    │   │   │   │   │   └── typing-manager.ts  # Typing indicator management
    │   │   │   │   └── websocket/
    │   │   │   │       ├── connection-manager.ts  # WebSocket connection management
    │   │   │   │       ├── error-recovery.ts  # Error handling & recovery
    │   │   │   │       ├── message-handler.ts  # Message routing
    │   │   │   │       └── subscription-manager.ts  # Subscription management
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
    │   │   │   │   ├── encryption.ts  # Client-side encryption
    │   │   │   │   ├── permissions.ts  # Permission checks
    │   │   │   │   ├── sanitization.ts  # Input sanitization
    │   │   │   │   └── validation.ts  # Input validation
    │   │   │   ├── sourcing/
    │   │   │   │   ├── sourcing-client.ts  # Sourcing API client
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
    │   │   ├── types/
    │   │   │   └── src/
    │   │   │       ├── api/  # API response types
    │   │   │       │   └── microservices/
    │   │   │       │       ├── admin-be.ts  # Admin backend API response types
    │   │   │       │       ├── communications-be.ts  # Communications backend API response types
    │   │   │       │       ├── contracts-be.ts  # Contracts backend API response types
    │   │   │       │       ├── financial-be.ts  # Financial backend API response types
    │   │   │       │       ├── jobs-be.ts  # Jobs backend API response types
    │   │   │       │       ├── proposals-be.ts  # Proposals backend API response types
    │   │   │       │       ├── reviews-be.ts  # Reviews backend API response types
    │   │   │       │       ├── search-be.ts  # Search backend API response types
    │   │   │       │       ├── storage-be.ts  # Storage backend API response types
    │   │   │       │       ├── subscriptions-be.ts  # Subscriptions backend API response types
    │   │   │       │       └── users-be.ts  # Users backend API response types
    │   │   │       ├── domains/  # Domain-specific types
    │   │   │       │   ├── admin/  # Admin domain types
    │   │   │       │   │   ├── admin-activity-log.ts  # Activity logging
    │   │   │       │   │   ├── admin-role.ts  # Admin role types
    │   │   │       │   │   ├── admin-session.ts  # Session management
    │   │   │       │   │   ├── admin-user.ts  # Admin user types
    │   │   │       │   │   ├── announcement.ts  # Platform announcements
    │   │   │       │   │   ├── audit-trail.ts  # Audit logging
    │   │   │       │   │   ├── content-action.ts  # Actions on content (remove, hide)
    │   │   │       │   │   ├── experiment.ts  # A/B experiments
    │   │   │       │   │   ├── feature-flag.ts  # Feature flags
    │   │   │       │   │   ├── index.ts  # Barrel export
    │   │   │       │   │   ├── maintenance.ts  # Maintenance mode
    │   │   │       │   │   ├── metrics.ts  # System metrics
    │   │   │       │   │   ├── moderation-queue.ts  # Moderation queue items
    │   │   │       │   │   ├── policy-doc.ts  # Policy documents
    │   │   │       │   │   ├── rate-limit.ts  # Rate limiting
    │   │   │       │   │   ├── support-ticket.ts  # Support tickets
    │   │   │       │   │   ├── system-health.ts  # System health monitoring
    │   │   │       │   │   ├── system-settings.ts  # System configuration
    │   │   │       │   │   ├── throttle-policy.ts  # Throttle policies
    │   │   │       │   │   └── user-action.ts  # Actions on users (suspend, ban)
    │   │   │       │   ├── communications/  # Communications domain
    │   │   │       │   │   ├── attachment.ts  # Message attachments
    │   │   │       │   │   ├── audit.ts  # Audit logs
    │   │   │       │   │   ├── collaboration.ts  # Real-time collaboration
    │   │   │       │   │   ├── compliance.ts  # Data retention/GDPR
    │   │   │       │   │   ├── conversation.ts  # Conversation types
    │   │   │       │   │   ├── digest.ts  # Email digests
    │   │   │       │   │   ├── email-notification.ts  # Email notifications
    │   │   │       │   │   ├── index.ts  # Barrel export
    │   │   │       │   │   ├── mention.ts  # @mentions
    │   │   │       │   │   ├── message.ts  # Message types
    │   │   │       │   │   ├── notification-preference.ts  # Notification settings
    │   │   │       │   │   ├── notification.ts  # Notification types
    │   │   │       │   │   ├── participant.ts  # Conversation participants
    │   │   │       │   │   ├── presence.ts  # User presence/online status
    │   │   │       │   │   ├── push-notification.ts  # Push notifications
    │   │   │       │   │   ├── rate-limit.ts  # Rate limiting
    │   │   │       │   │   ├── reaction.ts  # Message reactions
    │   │   │       │   │   ├── read-receipt.ts  # Read receipts
    │   │   │       │   │   ├── search.ts  # Message search
    │   │   │       │   │   ├── sms-notification.ts  # SMS notifications
    │   │   │       │   │   ├── thread.ts  # Message threads
    │   │   │       │   │   ├── typing-indicator.ts  # Typing indicators
    │   │   │       │   │   ├── webhook.ts  # Webhook subscriptions
    │   │   │       │   │   └── websocket.ts  # WebSocket connection
    │   │   │       │   ├── contracts/  # but needs additions
    │   │   │       │   │   ├── activity-level.ts  # Activity tracking
    │   │   │       │   │   ├── amendment.ts  # Contract amendments
    │   │   │       │   │   ├── approval.ts  # Approval workflow
    │   │   │       │   │   ├── bonus.ts  # Bonus clauses
    │   │   │       │   │   ├── budget.ts  # Budget tracking
    │   │   │       │   │   ├── clause.ts  # Contract clauses
    │   │   │       │   │   ├── compliance.ts  # Compliance tracking
    │   │   │       │   │   ├── contract.ts  # Contract entity types
    │   │   │       │   │   ├── dispute.ts  # Contract dispute types
    │   │   │       │   │   ├── expense.ts  # Expense tracking
    │   │   │       │   │   ├── index.ts  # Barrel export for contracts domain
    │   │   │       │   │   ├── insurance.ts  # Insurance coverage
    │   │   │       │   │   ├── liability.ts  # Liability clauses
    │   │   │       │   │   ├── milestone.ts  # Milestone entity types
    │   │   │       │   │   ├── penalty.ts  # Penalty clauses
    │   │   │       │   │   ├── renewal.ts  # Auto-renewal
    │   │   │       │   │   ├── screenshot.ts  # Screenshot monitoring
    │   │   │       │   │   ├── signature.ts  # Contract signatures
    │   │   │       │   │   ├── sla.ts  # Service level agreements
    │   │   │       │   │   ├── template.ts  # Contract templates
    │   │   │       │   │   ├── termination.ts  # Termination conditions
    │   │   │       │   │   ├── timesheet.ts  # Timesheet entity types
    │   │   │       │   │   └── work-diary.ts  # Work diary entries
    │   │   │       │   ├── financial/  # but needs additions
    │   │   │       │   │   ├── analytics.ts  # Financial analytics
    │   │   │       │   │   ├── balance-snapshot.ts  # Balance snapshots
    │   │   │       │   │   ├── bank-account.ts  # Bank accounts
    │   │   │       │   │   ├── bank-verification.ts  # Bank verification
    │   │   │       │   │   ├── bonus.ts  # Bonus payments
    │   │   │       │   │   ├── chargeback.ts  # Chargebacks
    │   │   │       │   │   ├── connects.ts  # Connects purchases
    │   │   │       │   │   ├── currency-preference.ts  # Currency preferences
    │   │   │       │   │   ├── dispute.ts  # Payment disputes
    │   │   │       │   │   ├── escrow.ts  # Escrow account types
    │   │   │       │   │   ├── expense.ts  # Expense reimbursements
    │   │   │       │   │   ├── fee-version.ts  # Fee versioning
    │   │   │       │   │   ├── fee.ts  # Fee structure types
    │   │   │       │   │   ├── forex.ts  # Currency exchange
    │   │   │       │   │   ├── fraud-alert.ts  # Fraud alerts
    │   │   │       │   │   ├── gateway.ts  # Payment gateway configs
    │   │   │       │   │   ├── index.ts  # Barrel export for financial domain
    │   │   │       │   │   ├── insurance.ts  # Payment insurance
    │   │   │       │   │   ├── international-payment.ts  # International payments
    │   │   │       │   │   ├── invoice.ts  # Invoices
    │   │   │       │   │   ├── ledger.ts  # Immutable ledger
    │   │   │       │   │   ├── payment-method.ts  # Payment method types
    │   │   │       │   │   ├── payment-schedule.ts  # Payment schedules
    │   │   │       │   │   ├── payment.ts  # Payment entity types
    │   │   │       │   │   ├── payout.ts  # Payout entity types
    │   │   │       │   │   ├── payroll.ts  # Payroll processing
    │   │   │       │   │   ├── promo.ts  # Promotional credits
    │   │   │       │   │   ├── reconciliation.ts  # Reconciliation
    │   │   │       │   │   ├── refund.ts  # Refund entity types
    │   │   │       │   │   ├── reminder.ts  # Payment reminders
    │   │   │       │   │   ├── risk.ts  # Risk assessment
    │   │   │       │   │   ├── subscription-billing.ts  # Subscription billing
    │   │   │       │   │   ├── tax-form.ts  # Tax forms (1099, W9)
    │   │   │       │   │   ├── tax.ts  # Tax calculation types
    │   │   │       │   │   ├── transaction.ts  # Transaction entity types
    │   │   │       │   │   ├── wallet.ts  # Digital wallet types
    │   │   │       │   │   └── withdrawal-limit.ts  # Withdrawal limits
    │   │   │       │   ├── jobs/  # but needs additions
    │   │   │       │   │   ├── analytics.ts  # Job analytics
    │   │   │       │   │   ├── archive.ts  # Archived jobs
    │   │   │       │   │   ├── boost.ts  # Job boosting
    │   │   │       │   │   ├── bulk-ops.ts  # Bulk operations
    │   │   │       │   │   ├── category.ts  # Job categories
    │   │   │       │   │   ├── custom-field.ts  # Custom fields
    │   │   │       │   │   ├── esg.ts  # ESG attributes
    │   │   │       │   │   ├── fraud-signal.ts  # Fraud detection
    │   │   │       │   │   ├── health-checkpoint.ts  # Health checks
    │   │   │       │   │   ├── index.ts  # Barrel export for jobs domain
    │   │   │       │   │   ├── invite.ts  # Freelancer invites
    │   │   │       │   │   ├── job-draft.ts  # Job drafts
    │   │   │       │   │   ├── job-post.ts  # Job posting types
    │   │   │       │   │   ├── job.ts  # Job entity types
    │   │   │       │   │   ├── localization.ts  # Multi-language
    │   │   │       │   │   ├── payment-schedule.ts  # Payment terms
    │   │   │       │   │   ├── preference.ts  # Client preferences
    │   │   │       │   │   ├── question.ts  # Screening questions
    │   │   │       │   │   ├── requirement.ts  # Job requirements
    │   │   │       │   │   ├── sharing.ts  # Share links
    │   │   │       │   │   ├── skill.ts  # Required skills
    │   │   │       │   │   ├── subcategory.ts  # Job subcategories
    │   │   │       │   │   ├── template.ts  # Job templates
    │   │   │       │   │   ├── upsell.ts  # Upsell suggestions
    │   │   │       │   │   ├── visibility.ts  # Job visibility settings
    │   │   │       │   │   └── webhook.ts  # Webhooks
    │   │   │       │   ├── proposals/  # but needs additions
    │   │   │       │   │   ├── ai-assist.ts  # AI suggestions
    │   │   │       │   │   ├── analytics.ts  # Proposal analytics types
    │   │   │       │   │   ├── archive.ts  # Archive tracking
    │   │   │       │   │   ├── attachment.ts  # Proposal attachments
    │   │   │       │   │   ├── auction.ts  # Auction mechanics
    │   │   │       │   │   ├── bid-strategy.ts  # Auto-bidding
    │   │   │       │   │   ├── bid.ts  # Bidding system
    │   │   │       │   │   ├── boost.ts  # Proposal boosting
    │   │   │       │   │   ├── collaboration.ts  # Team proposals
    │   │   │       │   │   ├── compliance.ts  # Compliance checks
    │   │   │       │   │   ├── connect.ts  # Connects usage
    │   │   │       │   │   ├── context.ts  # Context enrichment
    │   │   │       │   │   ├── conversation.ts  # Proposal conversations
    │   │   │       │   │   ├── cover-letter.ts  # Cover letters
    │   │   │       │   │   ├── engagement.ts  # Engagement tracking
    │   │   │       │   │   ├── expiration.ts  # Expiration tracking
    │   │   │       │   │   ├── feedback.ts  # Client feedback
    │   │   │       │   │   ├── flag.ts  # Flagging system
    │   │   │       │   │   ├── index.ts  # Barrel export for proposals domain
    │   │   │       │   │   ├── interview.ts  # Interview scheduling
    │   │   │       │   │   ├── invitation.ts  # Direct invitations
    │   │   │       │   │   ├── milestone-proposal.ts  # Milestone-based proposal types
    │   │   │       │   │   ├── milestone.ts  # Proposal milestones
    │   │   │       │   │   ├── negotiation.ts  # Negotiations
    │   │   │       │   │   ├── performance.ts  # Analytics
    │   │   │       │   │   ├── pipeline.ts  # Pipeline stages
    │   │   │       │   │   ├── portfolio-item.ts  # Portfolio item reference types
    │   │   │       │   │   ├── portfolio.ts  # Portfolio showcase
    │   │   │       │   │   ├── proposal.ts  # Proposal entity types
    │   │   │       │   │   ├── question-answer.ts  # Q&A responses
    │   │   │       │   │   ├── questions-answers.ts  # Q&A types
    │   │   │       │   │   ├── rate-card.ts  # Rate cards
    │   │   │       │   │   ├── recommendation.ts  # AI recommendations
    │   │   │       │   │   ├── recycling.ts  # Proposal reuse
    │   │   │       │   │   ├── reference.ts  # References
    │   │   │       │   │   ├── revision.ts  # Revision history
    │   │   │       │   │   ├── risk-assessment.ts  # Risk assessment
    │   │   │       │   │   ├── shortlist.ts  # Shortlisting
    │   │   │       │   │   ├── similarity.ts  # Deduplication
    │   │   │       │   │   ├── skill-match.ts  # Skill matching
    │   │   │       │   │   ├── spam-detection.ts  # Spam detection
    │   │   │       │   │   ├── template.ts  # Proposal templates
    │   │   │       │   │   ├── urgency.ts  # Urgency tracking
    │   │   │       │   │   ├── video-intro.ts  # Video introductions
    │   │   │       │   │   ├── withdrawal.ts  # Withdrawals
    │   │   │       │   │   └── withdrawn.ts  # Withdrawn proposal types
    │   │   │       │   ├── reviews/  # but needs additions
    │   │   │       │   │   ├── appeal.ts  # Review appeals
    │   │   │       │   │   ├── audit-trail.ts  # Audit trail
    │   │   │       │   │   ├── badge.ts  # Achievement badges
    │   │   │       │   │   ├── compliance.ts  # Compliance tracking
    │   │   │       │   │   ├── criteria.ts  # Rating criteria types
    │   │   │       │   │   ├── double-blind.ts  # Double-blind windows
    │   │   │       │   │   ├── draft.ts  # Review draft types
    │   │   │       │   │   ├── eligibility.ts  # Review eligibility
    │   │   │       │   │   ├── evidence.ts  # Appeal evidence
    │   │   │       │   │   ├── featured.ts  # Featured reviews
    │   │   │       │   │   ├── feedback.ts  # Private feedback types
    │   │   │       │   │   ├── flag.ts  # Review flagging
    │   │   │       │   │   ├── helpful-vote.ts  # Helpful votes
    │   │   │       │   │   ├── index.ts  # Barrel export for reviews domain
    │   │   │       │   │   ├── moderation.ts  # Review moderation
    │   │   │       │   │   ├── private-feedback.ts  # Private feedback
    │   │   │       │   │   ├── rating-criteria.ts  # Rating dimensions
    │   │   │       │   │   ├── rating.ts  # Rating entity types
    │   │   │       │   │   ├── redaction.ts  # PII redaction
    │   │   │       │   │   ├── reminder.ts  # Review reminders
    │   │   │       │   │   ├── reputation.ts  # Reputation scores
    │   │   │       │   │   ├── response.ts  # Review responses
    │   │   │       │   │   ├── review-draft.ts  # Review drafts
    │   │   │       │   │   ├── review.ts  # Review entity types
    │   │   │       │   │   ├── stats.ts  # Review statistics
    │   │   │       │   │   └── user-badge.ts  # User badges
    │   │   │       │   ├── search/  # but needs additions
    │   │   │       │   │   ├── ab-test.ts  # A/B testing
    │   │   │       │   │   ├── aggregation.ts  # Search aggregation types
    │   │   │       │   │   ├── alert.ts  # Search alerts
    │   │   │       │   │   ├── analytics.ts  # Search analytics
    │   │   │       │   │   ├── auto-complete.ts  # Auto-complete types
    │   │   │       │   │   ├── boost.ts  # Search boost types
    │   │   │       │   │   ├── facet.ts  # Search facet types
    │   │   │       │   │   ├── filter.ts  # Search filter types
    │   │   │       │   │   ├── highlight.ts  # Search highlight types
    │   │   │       │   │   ├── index.ts  # Barrel export for search domain
    │   │   │       │   │   ├── personalization.ts  # Personalized results
    │   │   │       │   │   ├── query.ts  # Search query types
    │   │   │       │   │   ├── ranking.ts  # Result ranking
    │   │   │       │   │   ├── recommendation.ts  # Search-based recommendation types
    │   │   │       │   │   ├── result.ts  # Search result types
    │   │   │       │   │   ├── saved-search.ts  # Saved searches
    │   │   │       │   │   ├── search-query.ts
    │   │   │       │   │   ├── search-result.ts
    │   │   │       │   │   ├── sort.ts  # Sort option types
    │   │   │       │   │   ├── spell-check.ts  # Spell check types
    │   │   │       │   │   ├── stopword.ts  # Stopword types
    │   │   │       │   │   ├── suggestion.ts  # Auto-suggestions
    │   │   │       │   │   └── synonym.ts  # Search synonyms
    │   │   │       │   ├── storage/  # but needs additions
    │   │   │       │   │   ├── access.ts  # Access control types
    │   │   │       │   │   ├── artifact.ts  # Processing artifact types
    │   │   │       │   │   ├── audit.ts  # Storage audit log types
    │   │   │       │   │   ├── cdn.ts  # CDN configuration
    │   │   │       │   │   ├── compliance.ts  # Compliance tracking
    │   │   │       │   │   ├── download.ts  # Download tracking
    │   │   │       │   │   ├── extraction.ts  # Content extraction types
    │   │   │       │   │   ├── file.ts  # File entity types
    │   │   │       │   │   ├── folder.ts  # Folders/directories
    │   │   │       │   │   ├── index.ts  # Barrel export for storage domain
    │   │   │       │   │   ├── lock.ts  # File lock types
    │   │   │       │   │   ├── media-processing.ts  # Media processing types
    │   │   │       │   │   ├── metadata.ts  # File metadata
    │   │   │       │   │   ├── permission.ts  # Access permissions
    │   │   │       │   │   ├── policy.ts  # Storage policy types
    │   │   │       │   │   ├── quarantine.ts  # Quarantined file types
    │   │   │       │   │   ├── quota.ts  # Storage quotas
    │   │   │       │   │   ├── scan.ts  # Virus scanning
    │   │   │       │   │   ├── share-link.ts  # Share link types
    │   │   │       │   │   ├── sharing.ts  # File sharing
    │   │   │       │   │   ├── thumbnail.ts  # Thumbnail types
    │   │   │       │   │   ├── upload-session.ts  # Resumable upload session types
    │   │   │       │   │   ├── upload.ts  # Upload sessions
    │   │   │       │   │   └── version.ts  # File versioning
    │   │   │       │   ├── subscriptions/  # but needs additions
    │   │   │       │   │   ├── addon.ts  # Addons
    │   │   │       │   │   ├── allowance.ts  # Allowances
    │   │   │       │   │   ├── analytics.ts  # Subscription analytics types
    │   │   │       │   │   ├── billing-cycle.ts  # Billing cycle types
    │   │   │       │   │   ├── billing-history.ts  # Billing history
    │   │   │       │   │   ├── billing-profile.ts  # Billing profiles
    │   │   │       │   │   ├── cancellation.ts  # Cancellation types
    │   │   │       │   │   ├── connect-balance.ts  # Connects system
    │   │   │       │   │   ├── credit-note.ts  # Credit notes
    │   │   │       │   │   ├── discount.ts  # Discount/coupon types
    │   │   │       │   │   ├── dunning.ts  # Dunning management
    │   │   │       │   │   ├── entitlement.ts  # Feature entitlements
    │   │   │       │   │   ├── feature-toggle.ts  # Feature toggles
    │   │   │       │   │   ├── feature.ts  # Feature entitlement types
    │   │   │       │   │   ├── index.ts  # Barrel export for subscriptions domain
    │   │   │       │   │   ├── invoice.ts  # Subscription invoices
    │   │   │       │   │   ├── payment-attempt.ts  # Payment attempts
    │   │   │       │   │   ├── payment.ts  # Subscription payment types
    │   │   │       │   │   ├── plan-pricing.ts  # Pricing tiers
    │   │   │       │   │   ├── plan-version.ts  # Plan versioning
    │   │   │       │   │   ├── plan.ts  # Subscription plan types
    │   │   │       │   │   ├── promotion.ts  # Promotions
    │   │   │       │   │   ├── proration.ts  # Proration types
    │   │   │       │   │   ├── renewal.ts  # Renewal types
    │   │   │       │   │   ├── seat.ts  # Seat billing
    │   │   │       │   │   ├── subscription.ts  # Subscription entity types
    │   │   │       │   │   ├── tax-class.ts  # Tax classes
    │   │   │       │   │   ├── tier.ts  # Plan tier types
    │   │   │       │   │   ├── trial.ts  # Free trials
    │   │   │       │   │   ├── upgrade.ts  # Upgrade/downgrade types
    │   │   │       │   │   └── usage.ts  # Usage tracking
    │   │   │       │   ├── users/  # but needs major additions
    │   │   │       │   │   ├── achievement.ts  # Achievements
    │   │   │       │   │   ├── activity.ts  # Activity feed
    │   │   │       │   │   ├── api-key.ts  # API keys
    │   │   │       │   │   ├── audit-log.ts  # User audit logs
    │   │   │       │   │   ├── automation.ts  # Automation rules
    │   │   │       │   │   ├── availability.ts  # Availability calendar
    │   │   │       │   │   ├── background-check.ts  # Background checks
    │   │   │       │   │   ├── badge.ts  # Platform badges
    │   │   │       │   │   ├── block.ts  # Blocked users
    │   │   │       │   │   ├── capability.ts  # Skills & specializations
    │   │   │       │   │   ├── certification.ts  # Certifications
    │   │   │       │   │   ├── client-profile.ts  # Client-specific profile types
    │   │   │       │   │   ├── client.ts  # Client-specific
    │   │   │       │   │   ├── compliance.ts  # Compliance records
    │   │   │       │   │   ├── earning.ts  # Earnings tracking
    │   │   │       │   │   ├── education.ts  # Education
    │   │   │       │   │   ├── experience.ts  # Work experience
    │   │   │       │   │   ├── follow.ts  # Follow/connection types
    │   │   │       │   │   ├── freelancer-profile.ts  # Freelancer-specific profile types
    │   │   │       │   │   ├── freelancer.ts  # Freelancer-specific
    │   │   │       │   │   ├── identity-verification.ts  # KYC/KYB
    │   │   │       │   │   ├── index.ts  # Barrel export for users domain
    │   │   │       │   │   ├── language.ts  # Language proficiency
    │   │   │       │   │   ├── metric.ts  # User metrics
    │   │   │       │   │   ├── notification-preference.ts  # Notification preference types
    │   │   │       │   │   ├── notification-setting.ts  # Notification prefs
    │   │   │       │   │   ├── oauth-app.ts  # OAuth applications
    │   │   │       │   │   ├── organization.ts  # Organization
    │   │   │       │   │   ├── portfolio-collection.ts  # Portfolio collections
    │   │   │       │   │   ├── portfolio-item.ts  # Portfolio item types
    │   │   │       │   │   ├── portfolio.ts  # Portfolio items
    │   │   │       │   │   ├── preference.ts  # User preferences
    │   │   │       │   │   ├── privacy-setting.ts  # Privacy settings
    │   │   │       │   │   ├── profile.ts  # User profile types
    │   │   │       │   │   ├── rate.ts  # Rate settings
    │   │   │       │   │   ├── referral.ts  # Referral program
    │   │   │       │   │   ├── report.ts  # User reports
    │   │   │       │   │   ├── reputation.ts  # Reputation system
    │   │   │       │   │   ├── saved-freelancer.ts  # Saved freelancers
    │   │   │       │   │   ├── saved-job.ts  # Saved jobs
    │   │   │       │   │   ├── saved-search.ts  # Saved searches
    │   │   │       │   │   ├── security-setting.ts  # Security settings
    │   │   │       │   │   ├── service-catalog.ts  # Service offerings
    │   │   │       │   │   ├── session.ts  # User sessions
    │   │   │       │   │   ├── setting.ts  # User setting types
    │   │   │       │   │   ├── skill-taxonomy.ts  # Skill taxonomy
    │   │   │       │   │   ├── skill.ts  # Individual skills
    │   │   │       │   │   ├── social-connection.ts  # Social connections
    │   │   │       │   │   ├── specialization.ts  # Specializations
    │   │   │       │   │   ├── spending.ts  # Spending tracking
    │   │   │       │   │   ├── team.ts  # Team management
    │   │   │       │   │   ├── trust-score.ts  # Trust scoring
    │   │   │       │   │   ├── user.ts  # User entity types
    │   │   │       │   │   ├── verification.ts  # Identity verification types
    │   │   │       │   │   └── webhook.ts  # User webhooks
    │   │   │       │   └── │
    │   │   │       ├── entities/  # Domain entities (DTOs)
    │   │   │       │   ├── admin/  # ENTIRE SECTION
    │   │   │       │   │   ├── admin-action.ts
    │   │   │       │   │   ├── admin-user.ts
    │   │   │       │   │   ├── announcement.ts
    │   │   │       │   │   ├── audit-entry.ts
    │   │   │       │   │   ├── experiment.ts
    │   │   │       │   │   ├── feature-flag.ts
    │   │   │       │   │   ├── moderation-item.ts
    │   │   │       │   │   ├── policy-document.ts
    │   │   │       │   │   ├── support-ticket.ts
    │   │   │       │   │   └── system-metric.ts
    │   │   │       │   ├── communications/  # ENTIRE SECTION
    │   │   │       │   │   ├── collaboration-session.ts
    │   │   │       │   │   ├── conversation.ts
    │   │   │       │   │   ├── email.ts
    │   │   │       │   │   ├── message.ts
    │   │   │       │   │   ├── notification.ts
    │   │   │       │   │   ├── push-notification.ts
    │   │   │       │   │   └── webhook-delivery.ts
    │   │   │       │   ├── contracts/
    │   │   │       │   ├── financial/
    │   │   │       │   ├── proposals/  # ENTIRE SECTION
    │   │   │       │   │   ├── bid-dto.ts
    │   │   │       │   │   ├── interview-dto.ts
    │   │   │       │   │   ├── negotiation-dto.ts
    │   │   │       │   │   ├── proposal-dto.ts
    │   │   │       │   │   ├── shortlist-dto.ts
    │   │   │       │   │   └── template-dto.ts
    │   │   │       │   ├── reviews/  # ENTIRE SECTION
    │   │   │       │   │   ├── appeal-dto.ts
    │   │   │       │   │   ├── badge-dto.ts
    │   │   │       │   │   ├── moderation-dto.ts
    │   │   │       │   │   ├── rating-dto.ts
    │   │   │       │   │   ├── reputation-dto.ts
    │   │   │       │   │   └── review-dto.ts
    │   │   │       │   ├── search/
    │   │   │       │   ├── storage/  # ENTIRE SECTION
    │   │   │       │   │   ├── file-dto.ts
    │   │   │       │   │   ├── folder-dto.ts
    │   │   │       │   │   ├── quota-dto.ts
    │   │   │       │   │   ├── share-link-dto.ts
    │   │   │       │   │   └── upload-session-dto.ts
    │   │   │       │   ├── subscriptions/  # ENTIRE SECTION
    │   │   │       │   │   ├── addon-dto.ts
    │   │   │       │   │   ├── entitlement-dto.ts
    │   │   │       │   │   ├── invoice-dto.ts
    │   │   │       │   │   ├── plan-dto.ts
    │   │   │       │   │   ├── subscription-dto.ts
    │   │   │       │   │   └── usage-dto.ts
    │   │   │       │   ├── users/  # ENTIRE SECTION
    │   │   │       │   │   ├── availability-dto.ts
    │   │   │       │   │   ├── capability-dto.ts
    │   │   │       │   │   ├── certification-dto.ts
    │   │   │       │   │   ├── education-dto.ts
    │   │   │       │   │   ├── experience-dto.ts
    │   │   │       │   │   ├── metric-dto.ts
    │   │   │       │   │   ├── organization-dto.ts
    │   │   │       │   │   ├── portfolio-dto.ts
    │   │   │       │   │   ├── profile-dto.ts
    │   │   │       │   │   ├── skill-dto.ts
    │   │   │       │   │   ├── team-dto.ts
    │   │   │       │   │   ├── trust-score-dto.ts
    │   │   │       │   │   ├── user-dto.ts
    │   │   │       │   │   └── verification-dto.ts
    │   │   │       │   └── │
    │   │   │       ├── enums/  # Enums
    │   │   │       │   ├── admin/  # ❌ CREATE
    │   │   │       │   │   ├── admin-permission.enum.ts  # Admin permission enumeration
    │   │   │       │   │   ├── admin-role.enum.ts  # Admin role enumeration
    │   │   │       │   │   ├── content-action-type.enum.ts  # Content action type enumeration
    │   │   │       │   │   ├── moderation-status.enum.ts  # Moderation status enumeration
    │   │   │       │   │   ├── system-health-status.enum.ts  # System health status enumeration
    │   │   │       │   │   ├── ticket-priority.enum.ts  # Ticket priority enumeration
    │   │   │       │   │   └── user-action-type.enum.ts  # User action type enumeration
    │   │   │       │   ├── communications/  # ❌ CREATE
    │   │   │       │   │   ├── conversation-type.enum.ts  # Conversation type enumeration
    │   │   │       │   │   ├── message-type.enum.ts  # Message type enumeration
    │   │   │       │   │   ├── notification-channel.enum.ts  # Notification channel enumeration
    │   │   │       │   │   ├── notification-priority.enum.ts  # Notification priority enumeration
    │   │   │       │   │   ├── notification-type.enum.ts  # Notification type enumeration
    │   │   │       │   │   ├── presence-status.enum.ts  # Presence status enumeration
    │   │   │       │   │   └── webhook-event.enum.ts  # Webhook event enumeration
    │   │   │       │   ├── contracts/  # but needs additions
    │   │   │       │   │   ├── amendment-status.enum.ts  # Amendment status enumeration
    │   │   │       │   │   ├── contract-status.enum.ts  # Contract status enumeration
    │   │   │       │   │   ├── contract-type.enum.ts  # Contract type enumeration
    │   │   │       │   │   ├── dispute-status.enum.ts  # Dispute status enumeration
    │   │   │       │   │   ├── milestone-status.enum.ts  # Milestone status enumeration
    │   │   │       │   │   ├── signature-status.enum.ts  # Signature status enumeration
    │   │   │       │   │   └── timesheet-status.enum.ts  # Timesheet status enumeration
    │   │   │       │   ├── financial/  # but needs additions
    │   │   │       │   │   ├── chargeback-status.enum.ts  # Chargeback status enumeration
    │   │   │       │   │   ├── dispute-status.enum.ts  # Payment dispute status enumeration
    │   │   │       │   │   ├── escrow-status.enum.ts  # Escrow status enumeration
    │   │   │       │   │   ├── payment-method-type.enum.ts  # Payment method type enumeration
    │   │   │       │   │   ├── payment-status.enum.ts  # Payment status enumeration
    │   │   │       │   │   ├── payout-status.enum.ts  # Payout status enumeration
    │   │   │       │   │   ├── refund-reason.enum.ts  # Refund reason enumeration
    │   │   │       │   │   ├── risk-level.enum.ts  # Risk level enumeration
    │   │   │       │   │   ├── tax-form-type.enum.ts  # Tax form type enumeration
    │   │   │       │   │   └── transaction-type.enum.ts  # Transaction type enumeration
    │   │   │       │   ├── jobs/  # but needs additions
    │   │   │       │   │   ├── boost-type.enum.ts  # Job boost type enumeration
    │   │   │       │   │   ├── duration.enum.ts  # Job duration enumeration
    │   │   │       │   │   ├── experience-level.enum.ts  # Experience level enumeration
    │   │   │       │   │   ├── job-status.enum.ts  # Job status enumeration
    │   │   │       │   │   ├── job-type.enum.ts  # Job type enumeration
    │   │   │       │   │   └── visibility.enum.ts  # Job visibility enumeration
    │   │   │       │   ├── proposals/  # but needs additions
    │   │   │       │   │   ├── bid-status.enum.ts  # Bid status enumeration
    │   │   │       │   │   ├── boost-type.enum.ts  # Proposal boost type enumeration
    │   │   │       │   │   ├── interview-status.enum.ts  # Interview status enumeration
    │   │   │       │   │   ├── negotiation-status.enum.ts  # Negotiation status enumeration
    │   │   │       │   │   ├── proposal-status.enum.ts  # Proposal status enumeration
    │   │   │       │   │   └── shortlist-status.enum.ts  # Shortlist status enumeration
    │   │   │       │   ├── reviews/  # but needs additions
    │   │   │       │   │   ├── appeal-status.enum.ts  # Appeal status enumeration
    │   │   │       │   │   ├── badge-type.enum.ts  # Badge type enumeration
    │   │   │       │   │   ├── moderation-status.enum.ts  # Review moderation status enumeration
    │   │   │       │   │   ├── rating-category.enum.ts  # Rating category enumeration
    │   │   │       │   │   ├── review-rating.enum.ts  # Review rating enumeration
    │   │   │       │   │   └── review-status.enum.ts  # Review status enumeration
    │   │   │       │   ├── storage/  # ❌ CREATE
    │   │   │       │   │   ├── access-level.enum.ts  # Access level enumeration
    │   │   │       │   │   ├── file-status.enum.ts  # File status enumeration
    │   │   │       │   │   ├── mime-type.enum.ts  # MIME type enumeration
    │   │   │       │   │   ├── upload-status.enum.ts  # Upload status enumeration
    │   │   │       │   │   └── virus-scan-status.enum.ts  # Virus scan status enumeration
    │   │   │       │   ├── subscriptions/  # ❌ CREATE
    │   │   │       │   │   ├── addon-type.enum.ts  # Add-on type enumeration
    │   │   │       │   │   ├── billing-cycle.enum.ts  # Billing cycle enumeration
    │   │   │       │   │   ├── dunning-status.enum.ts  # Dunning status enumeration
    │   │   │       │   │   ├── payment-attempt-status.enum.ts  # Payment attempt status enumeration
    │   │   │       │   │   ├── plan-type.enum.ts  # Plan type enumeration
    │   │   │       │   │   └── subscription-status.enum.ts  # Subscription status enumeration
    │   │   │       │   ├── users/  # but needs additions
    │   │   │       │   │   ├── account-status.enum.ts  # Account status enumeration
    │   │   │       │   │   ├── availability-status.enum.ts  # Availability status enumeration
    │   │   │       │   │   ├── badge-type.enum.ts  # Badge type enumeration
    │   │   │       │   │   ├── language-proficiency.enum.ts  # Language proficiency enumeration
    │   │   │       │   │   ├── referral-status.enum.ts  # Referral status enumeration
    │   │   │       │   │   ├── skill-proficiency.enum.ts  # Skill proficiency enumeration
    │   │   │       │   │   ├── trust-level.enum.ts  # Trust level enumeration
    │   │   │       │   │   ├── user-type.enum.ts  # User type enumeration
    │   │   │       │   │   └── verification-status.enum.ts  # Verification status enumeration
    │   │   │       │   └── │
    │   │   │       ├── models/  # Data transfer objects (DTOs)
    │   │   │       │   ├── admin/  # Admin DTOs
    │   │   │       │   │   ├── admin-role.ts  # Admin role DTO
    │   │   │       │   │   ├── admin-user.ts  # Admin user DTO
    │   │   │       │   │   ├── announcement.ts  # Announcement DTO
    │   │   │       │   │   ├── audit-entry.ts  # Audit entry DTO
    │   │   │       │   │   ├── experiment.ts  # Experiment DTO
    │   │   │       │   │   ├── feature-flag.ts  # Feature flag DTO
    │   │   │       │   │   ├── moderation-item.ts  # Moderation item DTO
    │   │   │       │   │   ├── policy-document.ts  # Policy document DTO
    │   │   │       │   │   ├── support-ticket.ts  # Support ticket DTO
    │   │   │       │   │   ├── system-metric.ts  # System metric DTO
    │   │   │       │   │   └── system-setting.ts  # System setting DTO
    │   │   │       │   ├── communications/  # Communications DTOs
    │   │   │       │   │   ├── collaboration-session.ts  # Collaboration session DTO
    │   │   │       │   │   ├── conversation.ts  # Conversation DTO
    │   │   │       │   │   ├── email.ts  # Email notification DTO
    │   │   │       │   │   ├── message.ts  # Message DTO
    │   │   │       │   │   ├── notification.ts  # Notification DTO
    │   │   │       │   │   ├── push-notification.ts  # Push notification DTO
    │   │   │       │   │   └── webhook-delivery.ts  # Webhook delivery DTO
    │   │   │       │   ├── proposals/  # Proposals DTOs
    │   │   │       │   │   ├── bid-dto.ts  # Bid DTO
    │   │   │       │   │   ├── interview-dto.ts  # Interview DTO
    │   │   │       │   │   ├── negotiation-dto.ts  # Negotiation DTO
    │   │   │       │   │   ├── proposal-dto.ts  # Proposal DTO
    │   │   │       │   │   ├── shortlist-dto.ts  # Shortlist DTO
    │   │   │       │   │   └── template-dto.ts  # Proposal template DTO
    │   │   │       │   ├── reviews/  # Reviews DTOs
    │   │   │       │   │   ├── appeal-dto.ts  # Appeal DTO
    │   │   │       │   │   ├── badge-dto.ts  # Badge DTO
    │   │   │       │   │   ├── moderation-dto.ts  # Moderation DTO
    │   │   │       │   │   ├── rating-dto.ts  # Rating DTO
    │   │   │       │   │   ├── reputation-dto.ts  # Reputation DTO
    │   │   │       │   │   └── review-dto.ts  # Review DTO
    │   │   │       │   ├── storage/  # Storage DTOs
    │   │   │       │   │   ├── file-dto.ts  # File DTO
    │   │   │       │   │   ├── folder-dto.ts  # Folder DTO
    │   │   │       │   │   ├── quota-dto.ts  # Quota DTO
    │   │   │       │   │   ├── share-link-dto.ts  # Share link DTO
    │   │   │       │   │   └── upload-session-dto.ts  # Upload session DTO
    │   │   │       │   ├── subscriptions/  # Subscriptions DTOs
    │   │   │       │   │   ├── addon-dto.ts  # Add-on DTO
    │   │   │       │   │   ├── entitlement-dto.ts  # Entitlement DTO
    │   │   │       │   │   ├── invoice-dto.ts  # Invoice DTO
    │   │   │       │   │   ├── plan-dto.ts  # Plan DTO
    │   │   │       │   │   ├── subscription-dto.ts  # Subscription DTO
    │   │   │       │   │   └── usage-dto.ts  # Usage DTO
    │   │   │       │   ├── users/  # Users DTOs
    │   │   │       │   │   ├── availability-dto.ts  # Availability DTO
    │   │   │       │   │   ├── capability-dto.ts  # Capability DTO
    │   │   │       │   │   ├── certification-dto.ts  # Certification DTO
    │   │   │       │   │   ├── education-dto.ts  # Education DTO
    │   │   │       │   │   ├── experience-dto.ts  # Experience DTO
    │   │   │       │   │   ├── metric-dto.ts  # Metric DTO
    │   │   │       │   │   ├── organization-dto.ts  # Organization DTO
    │   │   │       │   │   ├── portfolio-dto.ts  # Portfolio DTO
    │   │   │       │   │   ├── profile-dto.ts  # Profile DTO
    │   │   │       │   │   ├── skill-dto.ts  # Skill DTO
    │   │   │       │   │   ├── team-dto.ts  # Team DTO
    │   │   │       │   │   ├── trust-score-dto.ts  # Trust score DTO
    │   │   │       │   │   ├── user-dto.ts  # User DTO
    │   │   │       │   │   └── verification-dto.ts  # Verification DTO
    │   │   │       │   └── │
    │   │   │       └── │
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
    │   │   ├── src/
    │   │   │   ├── api/  # API response types
    │   │   │   │   ├── performance/
    │   │   │   │   │   └── index.ts  # Performance types
    │   │   │   │   ├── realtime/
    │   │   │   │   │   └── index.ts  # Real-time types
    │   │   │   │   ├── webhooks/
    │   │   │   │   │   └── index.ts  # Webhook types
    │   │   │   │   ├── admin.ts  # Admin API types
    │   │   │   │   ├── common.ts  # Common types (pagination, filters, etc.)
    │   │   │   │   ├── contracts.ts  # - Contracts API types
    │   │   │   │   ├── financial.ts  # - Financial API types
    │   │   │   │   ├── jobs.ts  # - Jobs API types
    │   │   │   │   ├── messages.ts  # - Messages API types
    │   │   │   │   ├── notifications.ts  # - Notifications API types
    │   │   │   │   ├── proposals.ts  # - Proposals API types
    │   │   │   │   ├── review.ts  # - Review API types
    │   │   │   │   ├── search.ts  # - Search API types
    │   │   │   │   ├── storage.ts  # - Storage API types
    │   │   │   │   ├── subscriptions.ts  # - Subscriptions API types
    │   │   │   │   └── users.ts  # - Users API types
    │   │   │   ├── config/
    │   │   │   │   └── index.ts  # Configuration types
    │   │   │   ├── domains/
    │   │   │   │   ├── achievements/
    │   │   │   │   │   └── index.ts  # Achievement types
    │   │   │   │   ├── activity/
    │   │   │   │   │   └── index.ts  # Activity feed types
    │   │   │   │   ├── compliance/
    │   │   │   │   │   └── index.ts  # Compliance types
    │   │   │   │   ├── experiments/
    │   │   │   │   │   └── index.ts  # Experiment types
    │   │   │   │   ├── flags/
    │   │   │   │   │   └── index.ts  # Feature flag types
    │   │   │   │   ├── incidents/
    │   │   │   │   │   └── index.ts  # Incident types
    │   │   │   │   ├── learning/
    │   │   │   │   │   └── index.ts  # Learning types
    │   │   │   │   ├── moderation/
    │   │   │   │   │   └── index.ts  # Moderation types
    │   │   │   │   ├── presence/
    │   │   │   │   │   └── index.ts  # Presence types
    │   │   │   │   └── sourcing/
    │   │   │   │       └── index.ts  # Sourcing types
    │   │   │   ├── entities/  # Domain entities
    │   │   │   │   ├── admin/
    │   │   │   │   │   ├── audit.ts  # Audit entity
    │   │   │   │   │   ├── flag.ts  # Content flag entity
    │   │   │   │   │   ├── incident.ts  # System incident entity
    │   │   │   │   │   ├── moderation-case.ts  # Moderation case entity
    │   │   │   │   │   └── report.ts  # Report entity
    │   │   │   │   ├── communications/
    │   │   │   │   │   ├── attachment.ts  # Message attachment entity
    │   │   │   │   │   ├── channel.ts  # Communication channel entity
    │   │   │   │   │   ├── template.ts  # Message template entity
    │   │   │   │   │   └── thread.ts  # Message thread entity
    │   │   │   │   ├── contracts/
    │   │   │   │   │   ├── amendment.ts  # Contract amendment entity
    │   │   │   │   │   ├── change-order.ts  # Change order entity
    │   │   │   │   │   ├── compliance.ts  # Contract compliance entity
    │   │   │   │   │   ├── deliverable.ts  # Deliverable entity
    │   │   │   │   │   ├── dispute.ts  # Dispute entity types
    │   │   │   │   │   ├── milestone.ts  # Milestone entity types
    │   │   │   │   │   ├── offer.ts  # Offer entity types
    │   │   │   │   │   ├── terms.ts  # Contract terms entity
    │   │   │   │   │   └── timesheet.ts  # Timesheet entity types
    │   │   │   │   ├── financial/
    │   │   │   │   │   ├── credit-note.ts  # Credit note entity
    │   │   │   │   │   ├── currency.ts  # Currency entity
    │   │   │   │   │   ├── escrow.ts  # Escrow DTOs
    │   │   │   │   │   ├── fee.ts  # Fee entity
    │   │   │   │   │   ├── payment-method.ts  # Payment method entity
    │   │   │   │   │   ├── payout.ts  # Payout entity
    │   │   │   │   │   ├── refund.ts  # Refund entity
    │   │   │   │   │   ├── report.ts  # Billing report row DTOs
    │   │   │   │   │   ├── reserve.ts  # Reserve entity
    │   │   │   │   │   ├── settlement.ts  # Settlement entity
    │   │   │   │   │   ├── tax.ts  # Tax DTOs
    │   │   │   │   │   └── wallet.ts  # Wallet entity
    │   │   │   │   ├── jobs/
    │   │   │   │   │   ├── application.ts  # Job application entity
    │   │   │   │   │   ├── invite.ts  # Job invite entity
    │   │   │   │   │   ├── requirement.ts  # Job requirement entity
    │   │   │   │   │   └── template.ts  # Job template entity
    │   │   │   │   ├── organizations/
    │   │   │   │   │   ├── budget.ts  # Budget entity
    │   │   │   │   │   ├── department.ts  # Department entity
    │   │   │   │   │   ├── organization.ts  # Organization entity
    │   │   │   │   │   ├── policy.ts  # Policy entity
    │   │   │   │   │   ├── procurement.ts  # Procurement entity
    │   │   │   │   │   └── team.ts  # Team entity
    │   │   │   │   ├── proposals/
    │   │   │   │   │   ├── bid.ts  # Bid entity
    │   │   │   │   │   ├── quote.ts  # Quote entity
    │   │   │   │   │   ├── rate-card.ts  # Rate card entity
    │   │   │   │   │   └── template.ts  # Proposal template entity
    │   │   │   │   ├── search/
    │   │   │   │   │   ├── alert.ts  # Search alert entity
    │   │   │   │   │   ├── facet.ts  # Search facet entity
    │   │   │   │   │   ├── filter.ts  # Search filter entity
    │   │   │   │   │   └── saved_query.ts  # Saved query/alert DTOs
    │   │   │   │   ├── storage/
    │   │   │   │   │   ├── access.ts  # File access entity
    │   │   │   │   │   ├── file.ts  # File entity
    │   │   │   │   │   ├── folder.ts  # Folder entity
    │   │   │   │   │   └── scan.ts  # File scan entity
    │   │   │   │   ├── subscriptions/
    │   │   │   │   │   ├── connects.ts  # Connects entity
    │   │   │   │   │   ├── entitlement.ts  # Entitlement entity
    │   │   │   │   │   ├── plan.ts  # Subscription plan entity
    │   │   │   │   │   └── usage.ts  # Usage entity
    │   │   │   │   ├── users/
    │   │   │   │   │   ├── availability.ts  # Availability entity
    │   │   │   │   │   ├── certification.ts  # Certification entity
    │   │   │   │   │   ├── education.ts  # Education entity
    │   │   │   │   │   ├── experience.ts  # Experience entity
    │   │   │   │   │   ├── portfolio.ts  # Portfolio entity
    │   │   │   │   │   ├── preference.ts  # User preference entity
    │   │   │   │   │   ├── profile.ts  # User profile entity
    │   │   │   │   │   ├── skill.ts  # Skill entity
    │   │   │   │   │   └── verification.ts  # Verification entity
    │   │   │   │   ├── contract.ts  # - Contract entity
    │   │   │   │   ├── invoice.ts  # - Invoice entity
    │   │   │   │   ├── job.ts  # - Job entity
    │   │   │   │   ├── message.ts  # - Message entity
    │   │   │   │   ├── notification.ts  # Notification DTOs
    │   │   │   │   ├── proposal.ts  # - Proposal entity
    │   │   │   │   ├── review.ts  # - Review entity
    │   │   │   │   ├── subscription.ts  # - Subscription entity
    │   │   │   │   ├── transaction.ts  # - Transaction entity
    │   │   │   │   ├── user.ts  # - User entity
    │   │   │   │   └── │
    │   │   │   ├── enums/  # Enums
    │   │   │   │   ├── contract-status.ts
    │   │   │   │   ├── country.ts  # - Country codes (ISO 3166)
    │   │   │   │   ├── currency.ts  # - USD, EUR, GBP, etc (ISO 4217)
    │   │   │   │   ├── dispute-status.ts  # - Open, InProgress, Resolved, Escalated
    │   │   │   │   ├── file-type.ts  # - Document, Image, Video, Audio, Archive
    │   │   │   │   ├── job-status.ts
    │   │   │   │   ├── language.ts  # - Language codes (ISO 639)
    │   │   │   │   ├── milestone-status.ts  # - Pending, InProgress, Completed, Approved
    │   │   │   │   ├── moderation-status.ts  # - Pending, Approved, Rejected, Flagged
    │   │   │   │   ├── notification-type.ts  # - System, Email, Push, SMS, InApp
    │   │   │   │   ├── payment-status.ts
    │   │   │   │   ├── permission.ts  # - RBAC permissions
    │   │   │   │   ├── proposal-status.ts
    │   │   │   │   ├── review-rating.ts
    │   │   │   │   ├── role.ts  # - Admin, User, Moderator, etc
    │   │   │   │   ├── skill-level.ts  # - Beginner, Intermediate, Advanced, Expert
    │   │   │   │   ├── subscription-status.ts  # - Active, PastDue, Cancelled, Paused, Trial
    │   │   │   │   ├── timezone.ts  # - IANA timezone identifiers
    │   │   │   │   ├── user-type.ts
    │   │   │   │   └── verification-status.ts  # - Pending, Verified, Rejected, Expired
    │   │   │   ├── errors/
    │   │   │   │   └── index.ts  # Error types
    │   │   │   ├── forms/
    │   │   │   │   └── index.ts  # Form types
    │   │   │   ├── routing/
    │   │   │   │   └── index.ts  # Routing types
    │   │   │   ├── shared/
    │   │   │   │   ├── address/
    │   │   │   │   │   └── index.ts  # Address types
    │   │   │   │   ├── analytics/
    │   │   │   │   │   └── index.ts  # Analytics types
    │   │   │   │   ├── api-response/
    │   │   │   │   │   └── index.ts  # API response types
    │   │   │   │   ├── contact/
    │   │   │   │   │   └── index.ts  # Contact types
    │   │   │   │   ├── datetime/
    │   │   │   │   │   └── index.ts  # Date/time types
    │   │   │   │   ├── filters/
    │   │   │   │   │   └── index.ts  # Filter types
    │   │   │   │   ├── geolocation/
    │   │   │   │   │   └── index.ts  # Geolocation types
    │   │   │   │   ├── media/
    │   │   │   │   │   └── index.ts  # Media types
    │   │   │   │   ├── money/
    │   │   │   │   │   └── index.ts  # Money/currency types
    │   │   │   │   ├── pagination/
    │   │   │   │   │   └── index.ts  # Pagination types
    │   │   │   │   ├── sort/
    │   │   │   │   │   └── index.ts  # Sort types
    │   │   │   │   └── validation/
    │   │   │   │       └── index.ts  # Validation types
    │   │   │   ├── state/
    │   │   │   │   └── index.ts  # State management types
    │   │   │   ├── audit.ts  # Audit Log/Export DTOs
    │   │   │   ├── index.ts  # Export all types
    │   │   │   ├── moderation.ts  # ModerationCase/Appeal types
    │   │   │   ├── search-admin.ts  # Taxonomy/Facet/Rewrite/Speller DTOs
    │   │   │   ├── subscription.ts  # Plan/Entitlement types
    │   │   │   ├── support.ts  # Ticket/Message types
    │   │   │   └── │
    │   │   ├── package.json
    │   │   ├── README.md
    │   │   ├── tsconfig.json
    │   │   └── │
    │   ├── ui/  # Cross-platform component library
    │   │   ├── src/
    │   │   │   ├── a11y/
    │   │   │   │   ├── Announcer/
    │   │   │   │   │   ├── LiveAnnouncer.native.tsx
    │   │   │   │   │   ├── LiveAnnouncer.tsx  # Live region announcements
    │   │   │   │   │   ├── LiveAnnouncer.types.ts
    │   │   │   │   │   └── LiveAnnouncer.web.tsx
    │   │   │   │   ├── FocusTrap/
    │   │   │   │   │   ├── FocusTrap.tsx
    │   │   │   │   │   ├── FocusTrap.types.ts
    │   │   │   │   │   └── FocusTrap.web.tsx
    │   │   │   │   ├── SkipLink/
    │   │   │   │   │   ├── SkipLink.tsx
    │   │   │   │   │   ├── SkipLink.types.ts
    │   │   │   │   │   └── SkipLink.web.tsx
    │   │   │   │   └── VisuallyHidden/
    │   │   │   │       ├── VisuallyHidden.native.tsx
    │   │   │   │       ├── VisuallyHidden.tsx
    │   │   │   │       ├── VisuallyHidden.types.ts
    │   │   │   │       └── VisuallyHidden.web.tsx
    │   │   │   ├── accessibility/
    │   │   │   │   ├── FocusTrap/
    │   │   │   │   │   ├── FocusTrap.native.tsx
    │   │   │   │   │   ├── FocusTrap.tsx
    │   │   │   │   │   └── FocusTrap.web.tsx
    │   │   │   │   ├── ScreenReaderAnnouncer/
    │   │   │   │   │   ├── ScreenReaderAnnouncer.native.tsx
    │   │   │   │   │   ├── ScreenReaderAnnouncer.tsx
    │   │   │   │   │   └── ScreenReaderAnnouncer.web.tsx
    │   │   │   │   └── SkipLinks/
    │   │   │   │       ├── SkipLinks.tsx
    │   │   │   │       └── SkipLinks.web.tsx
    │   │   │   ├── ai/
    │   │   │   │   ├── AIAssistant/
    │   │   │   │   │   ├── AIAssistant.native.tsx
    │   │   │   │   │   ├── AIAssistant.tsx
    │   │   │   │   │   └── AIAssistant.web.tsx
    │   │   │   │   ├── AutoComplete/
    │   │   │   │   │   ├── AIAutoComplete.native.tsx
    │   │   │   │   │   ├── AIAutoComplete.tsx
    │   │   │   │   │   └── AIAutoComplete.web.tsx
    │   │   │   │   └── SmartSuggestions/
    │   │   │   │       ├── SmartSuggestions.native.tsx
    │   │   │   │       ├── SmartSuggestions.tsx
    │   │   │   │       └── SmartSuggestions.web.tsx
    │   │   │   ├── auction/
    │   │   │   │   ├── AuctionTimer.native.tsx
    │   │   │   │   ├── AuctionTimer.tsx  # Countdown timer
    │   │   │   │   ├── AuctionTimer.web.tsx
    │   │   │   │   ├── BidHistoryChart.native.tsx
    │   │   │   │   ├── BidHistoryChart.tsx  # Bid history visualization
    │   │   │   │   ├── BidHistoryChart.web.tsx
    │   │   │   │   ├── LiveBidFeed.native.tsx
    │   │   │   │   ├── LiveBidFeed.tsx  # Real-time bid feed
    │   │   │   │   └── LiveBidFeed.web.tsx
    │   │   │   ├── charts/
    │   │   │   │   ├── EarningsChart.native.tsx
    │   │   │   │   ├── EarningsChart.tsx  # Earnings visualization
    │   │   │   │   ├── EarningsChart.web.tsx
    │   │   │   │   ├── PerformanceChart.native.tsx
    │   │   │   │   ├── PerformanceChart.tsx  # Performance metrics
    │   │   │   │   ├── PerformanceChart.web.tsx
    │   │   │   │   ├── TrendChart.native.tsx
    │   │   │   │   ├── TrendChart.tsx  # Trend visualization
    │   │   │   │   └── TrendChart.web.tsx
    │   │   │   ├── collaboration/
    │   │   │   │   ├── CollaborationPanel.native.tsx
    │   │   │   │   ├── CollaborationPanel.tsx  # Team collaboration
    │   │   │   │   ├── CollaborationPanel.web.tsx
    │   │   │   │   ├── GroupCard.native.tsx
    │   │   │   │   ├── GroupCard.tsx  # User group card
    │   │   │   │   ├── GroupCard.web.tsx
    │   │   │   │   ├── MentorCard.native.tsx
    │   │   │   │   ├── MentorCard.tsx  # Mentor profile card
    │   │   │   │   └── MentorCard.web.tsx
    │   │   │   ├── compliance/
    │   │   │   │   ├── DocumentUploader.native.tsx
    │   │   │   │   ├── DocumentUploader.tsx  # Compliance doc uploader
    │   │   │   │   ├── DocumentUploader.web.tsx
    │   │   │   │   ├── VerificationStatus.native.tsx
    │   │   │   │   ├── VerificationStatus.tsx  # Verification status badge
    │   │   │   │   └── VerificationStatus.web.tsx
    │   │   │   ├── components/
    │   │   │   │   ├── Accessibility/
    │   │   │   │   │   ├── FocusTrap/
    │   │   │   │   │   │   ├── FocusTrap.tsx
    │   │   │   │   │   │   ├── FocusTrap.types.ts
    │   │   │   │   │   │   └── FocusTrap.web.tsx
    │   │   │   │   │   ├── KeyboardShortcuts/
    │   │   │   │   │   │   ├── ShortcutDialog/
    │   │   │   │   │   │   │   └── ShortcutDialog.tsx
    │   │   │   │   │   │   ├── KeyboardShortcuts.tsx
    │   │   │   │   │   │   └── KeyboardShortcuts.types.ts
    │   │   │   │   │   ├── LiveRegion/
    │   │   │   │   │   │   ├── LiveRegion.native.tsx
    │   │   │   │   │   │   ├── LiveRegion.tsx
    │   │   │   │   │   │   ├── LiveRegion.types.ts
    │   │   │   │   │   │   └── LiveRegion.web.tsx
    │   │   │   │   │   └── ScreenReaderOnly/
    │   │   │   │   │       ├── ScreenReaderOnly.tsx
    │   │   │   │   │       └── ScreenReaderOnly.types.ts
    │   │   │   │   ├── Accordion/
    │   │   │   │   ├── AI/
    │   │   │   │   │   ├── AIAssistant/
    │   │   │   │   │   │   ├── chat/
    │   │   │   │   │   │   │   ├── ChatBubble.tsx
    │   │   │   │   │   │   │   └── ChatInput.tsx
    │   │   │   │   │   │   ├── ChatInterface/
    │   │   │   │   │   │   │   └── ChatInterface.tsx
    │   │   │   │   │   │   ├── AIAssistant.native.tsx
    │   │   │   │   │   │   ├── AIAssistant.tsx  # AI chat assistant
    │   │   │   │   │   │   ├── AIAssistant.types.ts
    │   │   │   │   │   │   └── AIAssistant.web.tsx
    │   │   │   │   │   ├── AutoComplete/
    │   │   │   │   │   │   ├── AIAutoComplete.native.tsx
    │   │   │   │   │   │   ├── AIAutoComplete.tsx  # AI-powered autocomplete
    │   │   │   │   │   │   ├── AIAutoComplete.types.ts
    │   │   │   │   │   │   ├── AIAutoComplete.web.tsx
    │   │   │   │   │   │   ├── SmartAutoComplete.native.tsx
    │   │   │   │   │   │   ├── SmartAutoComplete.tsx
    │   │   │   │   │   │   ├── SmartAutoComplete.types.ts
    │   │   │   │   │   │   └── SmartAutoComplete.web.tsx
    │   │   │   │   │   ├── ContentGeneration/
    │   │   │   │   │   │   ├── templates/
    │   │   │   │   │   │   │   └── GenerationTemplates.tsx
    │   │   │   │   │   │   ├── ContentGeneration.types.ts
    │   │   │   │   │   │   └── ContentGenerator.tsx  # AI content generation
    │   │   │   │   │   └── SmartSuggestions/
    │   │   │   │   │       ├── SuggestionCard/
    │   │   │   │   │       │   └── SuggestionCard.tsx
    │   │   │   │   │       ├── SmartSuggestions.native.tsx
    │   │   │   │   │       ├── SmartSuggestions.tsx  # AI suggestions
    │   │   │   │   │       ├── SmartSuggestions.types.ts
    │   │   │   │   │       └── SmartSuggestions.web.tsx
    │   │   │   │   ├── Alert/
    │   │   │   │   ├── auction/
    │   │   │   │   │   ├── AuctionTimer.native.tsx
    │   │   │   │   │   ├── AuctionTimer.tsx  # Countdown timer
    │   │   │   │   │   ├── AuctionTimer.web.tsx
    │   │   │   │   │   ├── BidHistoryChart.native.tsx
    │   │   │   │   │   ├── BidHistoryChart.tsx  # Bid history visualization
    │   │   │   │   │   ├── BidHistoryChart.web.tsx
    │   │   │   │   │   ├── LiveBidFeed.native.tsx
    │   │   │   │   │   ├── LiveBidFeed.tsx  # Real-time bid feed
    │   │   │   │   │   └── LiveBidFeed.web.tsx
    │   │   │   │   ├── Avatar/
    │   │   │   │   ├── Badge/
    │   │   │   │   ├── Breadcrumb/
    │   │   │   │   ├── Button/
    │   │   │   │   │   ├── Button.native.tsx  # Native-specific overrides
    │   │   │   │   │   ├── Button.stories.tsx  # Storybook stories
    │   │   │   │   │   ├── Button.test.tsx  # Component tests
    │   │   │   │   │   ├── Button.tsx  # Base button component
    │   │   │   │   │   └── Button.web.tsx  # Web-specific overrides
    │   │   │   │   ├── Calendar/
    │   │   │   │   │   ├── DatePicker/
    │   │   │   │   │   │   ├── DatePicker.native.tsx
    │   │   │   │   │   │   ├── DatePicker.tsx
    │   │   │   │   │   │   └── DatePicker.web.tsx
    │   │   │   │   │   ├── DateRangePicker/
    │   │   │   │   │   │   ├── DateRangePicker.native.tsx
    │   │   │   │   │   │   ├── DateRangePicker.tsx
    │   │   │   │   │   │   └── DateRangePicker.web.tsx
    │   │   │   │   │   ├── TimePicker/
    │   │   │   │   │   │   ├── TimePicker.native.tsx
    │   │   │   │   │   │   ├── TimePicker.tsx
    │   │   │   │   │   │   └── TimePicker.web.tsx
    │   │   │   │   │   ├── Calendar.native.tsx
    │   │   │   │   │   ├── Calendar.tsx  # Full calendar
    │   │   │   │   │   ├── Calendar.types.ts
    │   │   │   │   │   └── Calendar.web.tsx
    │   │   │   │   ├── Card/
    │   │   │   │   ├── charts/
    │   │   │   │   │   ├── EarningsChart.native.tsx
    │   │   │   │   │   ├── EarningsChart.tsx  # Earnings visualization
    │   │   │   │   │   ├── EarningsChart.web.tsx
    │   │   │   │   │   ├── PerformanceChart.native.tsx
    │   │   │   │   │   ├── PerformanceChart.tsx  # Performance metrics
    │   │   │   │   │   ├── PerformanceChart.web.tsx
    │   │   │   │   │   ├── TrendChart.native.tsx
    │   │   │   │   │   ├── TrendChart.tsx  # Trend visualization
    │   │   │   │   │   └── TrendChart.web.tsx
    │   │   │   │   ├── Charts/
    │   │   │   │   │   ├── FunnelChart/
    │   │   │   │   │   │   ├── FunnelChart.tsx
    │   │   │   │   │   │   └── FunnelChart.types.ts
    │   │   │   │   │   ├── GanttChart/
    │   │   │   │   │   │   ├── Task/
    │   │   │   │   │   │   │   └── Task.tsx
    │   │   │   │   │   │   ├── GanttChart.tsx
    │   │   │   │   │   │   └── GanttChart.types.ts
    │   │   │   │   │   ├── HeatMap/
    │   │   │   │   │   │   ├── HeatMap.tsx
    │   │   │   │   │   │   ├── HeatMap.types.ts
    │   │   │   │   │   │   └── HeatMap.web.tsx  # D3/Recharts
    │   │   │   │   │   └── OrgChart/
    │   │   │   │   │       ├── Node/
    │   │   │   │   │       │   └── OrgNode.tsx
    │   │   │   │   │       ├── OrgChart.tsx
    │   │   │   │   │       ├── OrgChart.types.ts
    │   │   │   │   │       └── OrgChart.web.tsx
    │   │   │   │   ├── Checkbox/
    │   │   │   │   ├── CodeEditor/
    │   │   │   │   │   ├── CodeEditor.native.tsx
    │   │   │   │   │   ├── CodeEditor.tsx
    │   │   │   │   │   ├── CodeEditor.types.ts
    │   │   │   │   │   └── CodeEditor.web.tsx  # Web (e.g., Monaco/CodeMirror)
    │   │   │   │   ├── collaboration/
    │   │   │   │   │   ├── CollaborationPanel.native.tsx
    │   │   │   │   │   ├── CollaborationPanel.tsx  # Team collaboration
    │   │   │   │   │   ├── CollaborationPanel.web.tsx
    │   │   │   │   │   ├── GroupCard.native.tsx
    │   │   │   │   │   ├── GroupCard.tsx  # User group card
    │   │   │   │   │   ├── GroupCard.web.tsx
    │   │   │   │   │   ├── MentorCard.native.tsx
    │   │   │   │   │   ├── MentorCard.tsx  # Mentor profile card
    │   │   │   │   │   └── MentorCard.web.tsx
    │   │   │   │   ├── compliance/
    │   │   │   │   │   ├── DocumentUploader.native.tsx
    │   │   │   │   │   ├── DocumentUploader.tsx  # Compliance doc uploader
    │   │   │   │   │   ├── DocumentUploader.web.tsx
    │   │   │   │   │   ├── VerificationStatus.native.tsx
    │   │   │   │   │   ├── VerificationStatus.tsx  # Verification status badge
    │   │   │   │   │   └── VerificationStatus.web.tsx
    │   │   │   │   ├── contests/  # ENTIRE SECTION
    │   │   │   │   │   ├── ContestCard.native.tsx  # Contest card (native)
    │   │   │   │   │   ├── ContestCard.tsx  # Contest card (base)
    │   │   │   │   │   ├── ContestCard.web.tsx  # Contest card (web)
    │   │   │   │   │   ├── EntryCard.native.tsx  # Entry card (native)
    │   │   │   │   │   ├── EntryCard.tsx  # Entry card (base)
    │   │   │   │   │   ├── EntryCard.web.tsx  # Entry card (web)
    │   │   │   │   │   ├── Leaderboard.native.tsx  # Leaderboard (native)
    │   │   │   │   │   ├── Leaderboard.tsx  # Leaderboard (base)
    │   │   │   │   │   └── Leaderboard.web.tsx  # Leaderboard (web)
    │   │   │   │   ├── DataDisplay/
    │   │   │   │   │   ├── Kanban/
    │   │   │   │   │   │   ├── KanbanBoard.native.tsx  # Touch gestures
    │   │   │   │   │   │   ├── KanbanBoard.tsx
    │   │   │   │   │   │   ├── KanbanBoard.types.ts
    │   │   │   │   │   │   └── KanbanBoard.web.tsx  # Drag & drop
    │   │   │   │   │   ├── List/
    │   │   │   │   │   │   ├── VirtualList/
    │   │   │   │   │   │   │   ├── VirtualList.native.tsx  # FlashList
    │   │   │   │   │   │   │   ├── VirtualList.tsx
    │   │   │   │   │   │   │   └── VirtualList.web.tsx
    │   │   │   │   │   │   ├── List.native.tsx
    │   │   │   │   │   │   ├── List.tsx
    │   │   │   │   │   │   ├── List.types.ts
    │   │   │   │   │   │   └── List.web.tsx
    │   │   │   │   │   ├── Table/
    │   │   │   │   │   │   ├── DataTable/
    │   │   │   │   │   │   │   ├── DataTable.tsx  # With sorting/filtering
    │   │   │   │   │   │   │   └── DataTable.types.ts
    │   │   │   │   │   │   ├── VirtualTable/
    │   │   │   │   │   │   │   ├── VirtualTable.tsx  # Virtualized table
    │   │   │   │   │   │   │   └── VirtualTable.types.ts
    │   │   │   │   │   │   ├── Table.tsx  # Base table
    │   │   │   │   │   │   ├── Table.types.ts
    │   │   │   │   │   │   └── Table.web.tsx  # Full featured table
    │   │   │   │   │   └── Timeline/
    │   │   │   │   │       ├── Timeline.native.tsx
    │   │   │   │   │       ├── Timeline.tsx
    │   │   │   │   │       ├── Timeline.types.ts
    │   │   │   │   │       └── Timeline.web.tsx
    │   │   │   │   ├── DataTable/
    │   │   │   │   ├── DatePicker/
    │   │   │   │   ├── Dropdown/
    │   │   │   │   ├── Editor/
    │   │   │   │   │   ├── CodeEditor/
    │   │   │   │   │   │   ├── CodeEditor.native.tsx
    │   │   │   │   │   │   ├── CodeEditor.tsx
    │   │   │   │   │   │   └── CodeEditor.web.tsx  # Web (e.g., Monaco/CodeMirror)
    │   │   │   │   │   ├── MarkdownEditor/
    │   │   │   │   │   │   ├── MarkdownEditor.native.tsx
    │   │   │   │   │   │   ├── MarkdownEditor.tsx
    │   │   │   │   │   │   └── MarkdownEditor.web.tsx
    │   │   │   │   │   └── RichTextEditor/
    │   │   │   │   │       ├── RichTextEditor.native.tsx  # Native implementation
    │   │   │   │   │       ├── RichTextEditor.tsx
    │   │   │   │   │       └── RichTextEditor.web.tsx  # Web (e.g., TipTap/Slate)
    │   │   │   │   ├── escrow/  # ENTIRE SECTION
    │   │   │   │   │   ├── EscrowStatus.native.tsx  # Escrow status (native)
    │   │   │   │   │   ├── EscrowStatus.tsx  # Escrow status (base)
    │   │   │   │   │   ├── EscrowStatus.web.tsx  # Escrow status (web)
    │   │   │   │   │   ├── ReleaseButton.native.tsx  # Release button (native)
    │   │   │   │   │   ├── ReleaseButton.tsx  # Release button (base)
    │   │   │   │   │   └── ReleaseButton.web.tsx  # Release button (web)
    │   │   │   │   ├── events/  # ENTIRE SECTION
    │   │   │   │   │   ├── EventCalendar.native.tsx  # Event calendar (native)
    │   │   │   │   │   ├── EventCalendar.tsx  # Event calendar (base)
    │   │   │   │   │   ├── EventCalendar.web.tsx  # Event calendar (web)
    │   │   │   │   │   ├── EventCard.native.tsx  # Event card (native)
    │   │   │   │   │   ├── EventCard.tsx  # Event card (base)
    │   │   │   │   │   └── EventCard.web.tsx  # Event card (web)
    │   │   │   │   ├── Feedback/
    │   │   │   │   │   ├── Alert/
    │   │   │   │   │   │   ├── Alert.native.tsx
    │   │   │   │   │   │   ├── Alert.tsx
    │   │   │   │   │   │   ├── Alert.types.ts
    │   │   │   │   │   │   └── Alert.web.tsx
    │   │   │   │   │   ├── EmptyState/
    │   │   │   │   │   │   ├── EmptyState.native.tsx
    │   │   │   │   │   │   ├── EmptyState.tsx
    │   │   │   │   │   │   ├── EmptyState.types.ts
    │   │   │   │   │   │   └── EmptyState.web.tsx
    │   │   │   │   │   ├── Notification/
    │   │   │   │   │   │   ├── Notification.native.tsx
    │   │   │   │   │   │   ├── Notification.tsx
    │   │   │   │   │   │   ├── Notification.types.ts
    │   │   │   │   │   │   └── Notification.web.tsx  # Toast notifications
    │   │   │   │   │   ├── Progress/
    │   │   │   │   │   │   ├── ProgressBar/
    │   │   │   │   │   │   │   ├── ProgressBar.native.tsx
    │   │   │   │   │   │   │   ├── ProgressBar.tsx
    │   │   │   │   │   │   │   └── ProgressBar.web.tsx
    │   │   │   │   │   │   ├── ProgressCircle/
    │   │   │   │   │   │   │   ├── ProgressCircle.native.tsx
    │   │   │   │   │   │   │   ├── ProgressCircle.tsx
    │   │   │   │   │   │   │   └── ProgressCircle.web.tsx
    │   │   │   │   │   │   └── Progress.types.ts
    │   │   │   │   │   └── Skeleton/
    │   │   │   │   │       ├── Skeleton.native.tsx
    │   │   │   │   │       ├── Skeleton.tsx
    │   │   │   │   │       ├── Skeleton.types.ts
    │   │   │   │   │       └── Skeleton.web.tsx
    │   │   │   │   ├── FileUpload/
    │   │   │   │   │   ├── ImageUpload/
    │   │   │   │   │   │   ├── ImageCropper.tsx  # Image cropping
    │   │   │   │   │   │   ├── ImageUpload.native.tsx  # Camera/gallery
    │   │   │   │   │   │   ├── ImageUpload.tsx
    │   │   │   │   │   │   └── ImageUpload.web.tsx
    │   │   │   │   │   ├── MultiFileUpload/
    │   │   │   │   │   │   ├── MultiFileUpload.native.tsx
    │   │   │   │   │   │   ├── MultiFileUpload.tsx
    │   │   │   │   │   │   └── MultiFileUpload.web.tsx
    │   │   │   │   │   ├── FileUpload.native.tsx  # Document picker
    │   │   │   │   │   ├── FileUpload.tsx  # Base file upload
    │   │   │   │   │   ├── FileUpload.types.ts
    │   │   │   │   │   └── FileUpload.web.tsx  # Drag & drop support
    │   │   │   │   ├── FileUploader/
    │   │   │   │   │   ├── FileUploader.native.tsx  # Uploader (native)
    │   │   │   │   │   ├── FileUploader.tsx  # Uploader (shared)
    │   │   │   │   │   └── FileUploader.web.tsx  # Uploader (web)
    │   │   │   │   ├── Form/
    │   │   │   │   │   ├── CodeEditor/
    │   │   │   │   │   │   ├── LanguageSelector/
    │   │   │   │   │   │   │   └── LanguageSelector.tsx
    │   │   │   │   │   │   ├── syntax-highlighting/
    │   │   │   │   │   │   │   ├── languages.ts
    │   │   │   │   │   │   │   └── themes.ts
    │   │   │   │   │   │   ├── CodeEditor.native.tsx  # Native code editor
    │   │   │   │   │   │   ├── CodeEditor.tsx  # Base code editor
    │   │   │   │   │   │   ├── CodeEditor.types.ts
    │   │   │   │   │   │   └── CodeEditor.web.tsx  # Monaco/CodeMirror
    │   │   │   │   │   ├── DateRangePicker/
    │   │   │   │   │   │   ├── presets/
    │   │   │   │   │   │   │   └── DatePresets.tsx
    │   │   │   │   │   │   ├── DateRangePicker.native.tsx
    │   │   │   │   │   │   ├── DateRangePicker.tsx
    │   │   │   │   │   │   ├── DateRangePicker.types.ts
    │   │   │   │   │   │   └── DateRangePicker.web.tsx
    │   │   │   │   │   ├── MarkdownEditor/
    │   │   │   │   │   │   ├── Preview/
    │   │   │   │   │   │   │   └── MarkdownPreview.tsx
    │   │   │   │   │   │   ├── preview/
    │   │   │   │   │   │   │   └── MarkdownPreview.tsx
    │   │   │   │   │   │   ├── MarkdownEditor.native.tsx
    │   │   │   │   │   │   ├── MarkdownEditor.tsx
    │   │   │   │   │   │   ├── MarkdownEditor.types.ts
    │   │   │   │   │   │   └── MarkdownEditor.web.tsx
    │   │   │   │   │   ├── RichTextEditor/
    │   │   │   │   │   │   ├── Toolbar/
    │   │   │   │   │   │   │   ├── Toolbar.tsx
    │   │   │   │   │   │   │   └── Toolbar.types.ts
    │   │   │   │   │   │   ├── toolbar/
    │   │   │   │   │   │   │   ├── FormatButtons.tsx
    │   │   │   │   │   │   │   ├── InsertButtons.tsx
    │   │   │   │   │   │   │   └── Toolbar.tsx
    │   │   │   │   │   │   ├── RichTextEditor.native.tsx  # Limited rich text
    │   │   │   │   │   │   ├── RichTextEditor.tsx  # Base editor
    │   │   │   │   │   │   ├── RichTextEditor.types.ts
    │   │   │   │   │   │   └── RichTextEditor.web.tsx  # Web (TipTap/Slate)
    │   │   │   │   │   └── SignaturePad/
    │   │   │   │   │       ├── SignaturePad.native.tsx  # React Native Canvas
    │   │   │   │   │       ├── SignaturePad.tsx
    │   │   │   │   │       ├── SignaturePad.types.ts
    │   │   │   │   │       └── SignaturePad.web.tsx  # Canvas-based
    │   │   │   │   ├── Forms/
    │   │   │   │   │   ├── Checkbox/
    │   │   │   │   │   │   ├── Checkbox.native.tsx
    │   │   │   │   │   │   ├── Checkbox.tsx
    │   │   │   │   │   │   ├── Checkbox.types.ts
    │   │   │   │   │   │   └── Checkbox.web.tsx
    │   │   │   │   │   ├── FormField/
    │   │   │   │   │   │   ├── FormError.tsx
    │   │   │   │   │   │   ├── FormField.native.tsx
    │   │   │   │   │   │   ├── FormField.tsx  # Form field wrapper
    │   │   │   │   │   │   ├── FormField.types.ts
    │   │   │   │   │   │   ├── FormField.web.tsx
    │   │   │   │   │   │   ├── FormHelperText.tsx
    │   │   │   │   │   │   └── FormLabel.tsx
    │   │   │   │   │   ├── Radio/
    │   │   │   │   │   │   ├── Radio.types.ts
    │   │   │   │   │   │   ├── RadioGroup.native.tsx
    │   │   │   │   │   │   ├── RadioGroup.tsx
    │   │   │   │   │   │   └── RadioGroup.web.tsx
    │   │   │   │   │   ├── Select/
    │   │   │   │   │   │   ├── MultiSelect/
    │   │   │   │   │   │   │   ├── MultiSelect.native.tsx
    │   │   │   │   │   │   │   ├── MultiSelect.tsx
    │   │   │   │   │   │   │   └── MultiSelect.web.tsx
    │   │   │   │   │   │   ├── Select.native.tsx
    │   │   │   │   │   │   ├── Select.tsx
    │   │   │   │   │   │   ├── Select.types.ts
    │   │   │   │   │   │   └── Select.web.tsx
    │   │   │   │   │   ├── Slider/
    │   │   │   │   │   │   ├── RangeSlider/
    │   │   │   │   │   │   │   ├── RangeSlider.native.tsx
    │   │   │   │   │   │   │   ├── RangeSlider.tsx
    │   │   │   │   │   │   │   └── RangeSlider.web.tsx
    │   │   │   │   │   │   ├── Slider.native.tsx
    │   │   │   │   │   │   ├── Slider.tsx
    │   │   │   │   │   │   ├── Slider.types.ts
    │   │   │   │   │   │   └── Slider.web.tsx
    │   │   │   │   │   └── Switch/
    │   │   │   │   │       ├── Switch.native.tsx
    │   │   │   │   │       ├── Switch.tsx
    │   │   │   │   │       ├── Switch.types.ts
    │   │   │   │   │       └── Switch.web.tsx
    │   │   │   │   ├── groups/  # ENTIRE SECTION
    │   │   │   │   │   ├── GroupCard.native.tsx  # Group card (native)
    │   │   │   │   │   ├── GroupCard.tsx  # Group card (base)
    │   │   │   │   │   ├── GroupCard.web.tsx  # Group card (web)
    │   │   │   │   │   ├── GroupPost.native.tsx  # Group post (native)
    │   │   │   │   │   ├── GroupPost.tsx  # Group post (base)
    │   │   │   │   │   └── GroupPost.web.tsx  # Group post (web)
    │   │   │   │   ├── Input/
    │   │   │   │   │   ├── Input.native.tsx
    │   │   │   │   │   ├── Input.test.tsx
    │   │   │   │   │   ├── Input.tsx
    │   │   │   │   │   └── Input.web.tsx
    │   │   │   │   ├── learning/  # ENTIRE SECTION
    │   │   │   │   │   ├── AchievementBadge.native.tsx
    │   │   │   │   │   ├── AchievementBadge.tsx  # Achievement badge
    │   │   │   │   │   ├── AchievementBadge.web.tsx
    │   │   │   │   │   ├── CourseCard.native.tsx  # Course card (native)
    │   │   │   │   │   ├── CourseCard.tsx  # Course card (base)
    │   │   │   │   │   ├── CourseCard.web.tsx  # Course card (web)
    │   │   │   │   │   ├── LearningPathCard.native.tsx
    │   │   │   │   │   ├── LearningPathCard.tsx  # Learning path card
    │   │   │   │   │   ├── LearningPathCard.web.tsx
    │   │   │   │   │   ├── LessonPlayer.native.tsx  # Lesson player (native)
    │   │   │   │   │   ├── LessonPlayer.tsx  # Lesson player (base)
    │   │   │   │   │   ├── LessonPlayer.web.tsx  # Lesson player (web)
    │   │   │   │   │   ├── ProgressTracker.native.tsx  # Progress tracker (native)
    │   │   │   │   │   ├── ProgressTracker.tsx  # Progress visualization
    │   │   │   │   │   └── ProgressTracker.web.tsx  # Progress tracker (web)
    │   │   │   │   ├── LoadingSpinner/
    │   │   │   │   │   ├── LoadingSpinner.native.tsx  # Spinner (native)
    │   │   │   │   │   ├── LoadingSpinner.tsx  # Spinner (shared)
    │   │   │   │   │   └── LoadingSpinner.web.tsx  # Spinner (web)
    │   │   │   │   ├── Modal/
    │   │   │   │   │   ├── Modal.native.tsx  # Modal (native)
    │   │   │   │   │   ├── Modal.tsx  # Modal (shared)
    │   │   │   │   │   └── Modal.web.tsx  # Modal (web)
    │   │   │   │   ├── molecules/
    │   │   │   │   │   ├── Navigation/
    │   │   │   │   │   │   ├── Accordion/
    │   │   │   │   │   │   │   ├── Accordion.native.tsx  # React Native collapsible accordion component
    │   │   │   │   │   │   │   ├── Accordion.tsx  # - Base component
    │   │   │   │   │   │   │   ├── Accordion.types.ts  # - Shared types
    │   │   │   │   │   │   │   └── Accordion.web.tsx  # - Web implementation
    │   │   │   │   │   │   ├── Breadcrumb/
    │   │   │   │   │   │   │   ├── Breadcrumb.native.tsx  # Mobile breadcrumb navigation component
    │   │   │   │   │   │   │   ├── Breadcrumb.tsx  # - Base component
    │   │   │   │   │   │   │   ├── Breadcrumb.types.ts  # - Shared types
    │   │   │   │   │   │   │   └── Breadcrumb.web.tsx  # - Web implementation
    │   │   │   │   │   │   ├── Drawer/
    │   │   │   │   │   │   │   ├── Drawer.native.tsx  # React Native bottom sheet drawer component
    │   │   │   │   │   │   │   ├── Drawer.tsx  # - Base component
    │   │   │   │   │   │   │   ├── Drawer.types.ts  # - Shared types
    │   │   │   │   │   │   │   └── Drawer.web.tsx  # - Web implementation
    │   │   │   │   │   │   ├── Tabs/
    │   │   │   │   │   │   │   └── Tabs.native.tsx  # 🔧 ENHANCE - Mobile tabs
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── Overlay/
    │   │   │   │   │   │   ├── Popover/
    │   │   │   │   │   │   │   ├── Popover.native.tsx  # Mobile popover component
    │   │   │   │   │   │   │   ├── Popover.tsx  # - Base component
    │   │   │   │   │   │   │   ├── Popover.types.ts  # - Shared types
    │   │   │   │   │   │   │   └── Popover.web.tsx  # - Web implementation
    │   │   │   │   │   │   ├── Tooltip/
    │   │   │   │   │   │   │   ├── Tooltip.native.tsx  # Mobile tooltip component
    │   │   │   │   │   │   │   ├── Tooltip.tsx  # - Base component
    │   │   │   │   │   │   │   ├── Tooltip.types.ts  # - Shared types
    │   │   │   │   │   │   │   └── Tooltip.web.tsx  # - Web implementation
    │   │   │   │   │   │   └── │
    │   │   │   │   │   └── │
    │   │   │   │   ├── Navigation/
    │   │   │   │   │   ├── Breadcrumb/
    │   │   │   │   │   │   ├── Breadcrumb.tsx
    │   │   │   │   │   │   ├── Breadcrumb.types.ts
    │   │   │   │   │   │   └── Breadcrumb.web.tsx
    │   │   │   │   │   ├── Pagination/
    │   │   │   │   │   │   ├── Pagination.native.tsx
    │   │   │   │   │   │   ├── Pagination.tsx
    │   │   │   │   │   │   ├── Pagination.types.ts
    │   │   │   │   │   │   └── Pagination.web.tsx
    │   │   │   │   │   ├── Stepper/
    │   │   │   │   │   │   ├── Stepper.native.tsx
    │   │   │   │   │   │   ├── Stepper.tsx
    │   │   │   │   │   │   ├── Stepper.types.ts
    │   │   │   │   │   │   └── Stepper.web.tsx
    │   │   │   │   │   └── Tabs/
    │   │   │   │   │       ├── Tabs.native.tsx
    │   │   │   │   │       ├── Tabs.tsx
    │   │   │   │   │       ├── Tabs.types.ts
    │   │   │   │   │       └── Tabs.web.tsx
    │   │   │   │   ├── organisms/
    │   │   │   │   │   ├── Charts/
    │   │   │   │   │   │   ├── FunnelChart/
    │   │   │   │   │   │   │   └── FunnelChart.native.tsx  # Mobile funnel chart
    │   │   │   │   │   │   ├── HeatMap/
    │   │   │   │   │   │   │   └── HeatMap.native.tsx  # Mobile heat map
    │   │   │   │   │   │   ├── OrgChart/
    │   │   │   │   │   │   │   └── OrgChart.native.tsx  # Mobile org chart
    │   │   │   │   │   │   └── │
    │   │   │   │   │   ├── DataDisplay/
    │   │   │   │   │   │   ├── DataGrid/
    │   │   │   │   │   │   │   ├── DataGrid.native.tsx  # Mobile data grid component
    │   │   │   │   │   │   │   ├── DataGrid.tsx  # - Base component
    │   │   │   │   │   │   │   ├── DataGrid.types.ts  # - Shared types
    │   │   │   │   │   │   │   └── DataGrid.web.tsx  # - Web implementation
    │   │   │   │   │   │   ├── KanbanBoard/
    │   │   │   │   │   │   │   ├── KanbanBoard.native.tsx  # Mobile kanban board component
    │   │   │   │   │   │   │   ├── KanbanBoard.tsx  # - Base component
    │   │   │   │   │   │   │   ├── KanbanBoard.types.ts  # - Shared types
    │   │   │   │   │   │   │   └── KanbanBoard.web.tsx  # - Web implementation
    │   │   │   │   │   │   ├── Popover/
    │   │   │   │   │   │   │   └── Popover.native.tsx  # Positioned overlay; arrow; fade+scale; portal
    │   │   │   │   │   │   ├── Table/
    │   │   │   │   │   │   │   ├── Table.native.tsx  # Mobile responsive table component
    │   │   │   │   │   │   │   ├── Table.tsx  # - Base component
    │   │   │   │   │   │   │   ├── Table.types.ts  # - Shared types
    │   │   │   │   │   │   │   └── Table.web.tsx  # - Web implementation
    │   │   │   │   │   │   ├── Tooltip/
    │   │   │   │   │   │   │   └── Tooltip.native.tsx  # Long press; auto-dismiss; a11y; simple text
    │   │   │   │   │   │   └── │
    │   │   │   │   │   └── │
    │   │   │   │   ├── Overlay/
    │   │   │   │   │   ├── Popover/
    │   │   │   │   │   │   └── Popover.native.tsx  # Mobile popover component
    │   │   │   │   │   ├── Tooltip/
    │   │   │   │   │   │   └── Tooltip.native.tsx  # Mobile tooltip component
    │   │   │   │   │   └── │
    │   │   │   │   ├── Pagination/
    │   │   │   │   ├── Popover/
    │   │   │   │   ├── Progress/
    │   │   │   │   ├── Radio/
    │   │   │   │   ├── Rating/
    │   │   │   │   ├── referrals/  # ENTIRE SECTION
    │   │   │   │   │   ├── ReferralCard.native.tsx  # Referral card (native)
    │   │   │   │   │   ├── ReferralCard.tsx  # Referral card (base)
    │   │   │   │   │   ├── ReferralCard.web.tsx  # Referral card (web)
    │   │   │   │   │   ├── ReferralStats.native.tsx  # Referral stats (native)
    │   │   │   │   │   ├── ReferralStats.tsx  # Referral stats (base)
    │   │   │   │   │   └── ReferralStats.web.tsx  # Referral stats (web)
    │   │   │   │   ├── Select/
    │   │   │   │   ├── Skeleton/
    │   │   │   │   ├── skills-tests/  # ENTIRE SECTION
    │   │   │   │   │   ├── TestCard.native.tsx  # Test card (native)
    │   │   │   │   │   ├── TestCard.tsx  # Test card (base)
    │   │   │   │   │   ├── TestCard.web.tsx  # Test card (web)
    │   │   │   │   │   ├── TestQuestion.native.tsx  # Test question (native)
    │   │   │   │   │   ├── TestQuestion.tsx  # Test question (base)
    │   │   │   │   │   ├── TestQuestion.web.tsx  # Test question (web)
    │   │   │   │   │   ├── TestResults.native.tsx  # Test results (native)
    │   │   │   │   │   ├── TestResults.tsx  # Test results (base)
    │   │   │   │   │   └── TestResults.web.tsx  # Test results (web)
    │   │   │   │   ├── Slider/
    │   │   │   │   ├── Stepper/
    │   │   │   │   ├── Switch/
    │   │   │   │   ├── Tabs/
    │   │   │   │   ├── Textarea/
    │   │   │   │   ├── Timeline/
    │   │   │   │   ├── TimePicker/
    │   │   │   │   ├── Toast/
    │   │   │   │   │   ├── Toast.native.tsx  # Toasts (native)
    │   │   │   │   │   ├── Toast.tsx  # Toasts (shared)
    │   │   │   │   │   └── Toast.web.tsx  # Toasts (web)
    │   │   │   │   ├── Tooltip/
    │   │   │   │   ├── tracking/
    │   │   │   │   │   ├── TimesheetTable.native.tsx
    │   │   │   │   │   ├── TimesheetTable.tsx  # Timesheet grid
    │   │   │   │   │   ├── TimesheetTable.web.tsx
    │   │   │   │   │   ├── TimeTracker.native.tsx
    │   │   │   │   │   ├── TimeTracker.tsx  # Time tracking widget
    │   │   │   │   │   ├── TimeTracker.web.tsx
    │   │   │   │   │   ├── WorkDiaryEntry.native.tsx
    │   │   │   │   │   ├── WorkDiaryEntry.tsx  # Work diary card
    │   │   │   │   │   └── WorkDiaryEntry.web.tsx
    │   │   │   │   ├── video/
    │   │   │   │   │   ├── VideoPlayer.native.tsx
    │   │   │   │   │   ├── VideoPlayer.tsx  # Video player
    │   │   │   │   │   ├── VideoPlayer.web.tsx
    │   │   │   │   │   ├── VideoUploader.native.tsx
    │   │   │   │   │   ├── VideoUploader.tsx  # Video upload
    │   │   │   │   │   └── VideoUploader.web.tsx
    │   │   │   │   ├── Visualization/
    │   │   │   │   │   ├── GanttChart/
    │   │   │   │   │   │   ├── timeline/
    │   │   │   │   │   │   │   ├── Timeline.tsx
    │   │   │   │   │   │   │   └── TimelineItem.tsx
    │   │   │   │   │   │   ├── GanttChart.native.tsx
    │   │   │   │   │   │   ├── GanttChart.tsx
    │   │   │   │   │   │   ├── GanttChart.types.ts
    │   │   │   │   │   │   └── GanttChart.web.tsx
    │   │   │   │   │   ├── HeatMap/
    │   │   │   │   │   │   ├── HeatMap.native.tsx
    │   │   │   │   │   │   ├── HeatMap.tsx
    │   │   │   │   │   │   ├── HeatMap.types.ts
    │   │   │   │   │   │   └── HeatMap.web.tsx
    │   │   │   │   │   ├── KanbanBoard/
    │   │   │   │   │   │   ├── card/
    │   │   │   │   │   │   │   └── KanbanCard.tsx
    │   │   │   │   │   │   ├── column/
    │   │   │   │   │   │   │   └── KanbanColumn.tsx
    │   │   │   │   │   │   ├── KanbanBoard.native.tsx
    │   │   │   │   │   │   ├── KanbanBoard.tsx
    │   │   │   │   │   │   ├── KanbanBoard.types.ts
    │   │   │   │   │   │   └── KanbanBoard.web.tsx
    │   │   │   │   │   ├── OrgChart/
    │   │   │   │   │   │   ├── node/
    │   │   │   │   │   │   │   └── OrgNode.tsx
    │   │   │   │   │   │   ├── OrgChart.native.tsx
    │   │   │   │   │   │   ├── OrgChart.tsx
    │   │   │   │   │   │   ├── OrgChart.types.ts
    │   │   │   │   │   │   └── OrgChart.web.tsx
    │   │   │   │   │   └── TreeView/
    │   │   │   │   │       ├── node/
    │   │   │   │   │       │   └── TreeNode.tsx
    │   │   │   │   │       ├── TreeView.native.tsx
    │   │   │   │   │       ├── TreeView.tsx
    │   │   │   │   │       ├── TreeView.types.ts
    │   │   │   │   │       └── TreeView.web.tsx
    │   │   │   │   └── │
    │   │   │   ├── forms/  # Form components
    │   │   │   │   ├── CodeEditor/
    │   │   │   │   │   ├── CodeEditor.native.tsx
    │   │   │   │   │   ├── CodeEditor.tsx
    │   │   │   │   │   └── CodeEditor.web.tsx
    │   │   │   │   ├── DateRangePicker/
    │   │   │   │   │   ├── DateRangePicker.native.tsx
    │   │   │   │   │   ├── DateRangePicker.tsx
    │   │   │   │   │   └── DateRangePicker.web.tsx
    │   │   │   │   ├── FormError/
    │   │   │   │   ├── FormField/
    │   │   │   │   ├── FormGroup/
    │   │   │   │   ├── FormHelper/
    │   │   │   │   ├── FormLabel/
    │   │   │   │   ├── MarkdownEditor/
    │   │   │   │   │   ├── MarkdownEditor.native.tsx
    │   │   │   │   │   ├── MarkdownEditor.tsx
    │   │   │   │   │   └── MarkdownEditor.web.tsx
    │   │   │   │   ├── RichTextEditor/
    │   │   │   │   │   ├── RichTextEditor.native.tsx
    │   │   │   │   │   ├── RichTextEditor.tsx
    │   │   │   │   │   └── RichTextEditor.web.tsx
    │   │   │   │   └── SignatureInput/
    │   │   │   │       ├── SignatureInput.native.tsx
    │   │   │   │       ├── SignatureInput.tsx
    │   │   │   │       └── SignatureInput.web.tsx
    │   │   │   ├── icons/  # Icon components
    │   │   │   │   └── index.ts  # Export all icons
    │   │   │   ├── layout/  # Layout components
    │   │   │   │   ├── Container/
    │   │   │   │   ├── Divider/
    │   │   │   │   ├── Grid/
    │   │   │   │   ├── Spacer/
    │   │   │   │   └── Stack/
    │   │   │   ├── learning/
    │   │   │   │   ├── AchievementBadge.native.tsx
    │   │   │   │   ├── AchievementBadge.tsx  # Achievement badge
    │   │   │   │   ├── AchievementBadge.web.tsx
    │   │   │   │   ├── LearningPathCard.native.tsx
    │   │   │   │   ├── LearningPathCard.tsx  # Learning path card
    │   │   │   │   ├── LearningPathCard.web.tsx
    │   │   │   │   ├── ProgressTracker.native.tsx
    │   │   │   │   ├── ProgressTracker.tsx  # Progress visualization
    │   │   │   │   └── ProgressTracker.web.tsx
    │   │   │   ├── molecules/
    │   │   │   │   ├── Navigation/
    │   │   │   │   │   ├── Accordion/
    │   │   │   │   │   │   ├── Accordion.native.tsx  # - Mobile needs accordion
    │   │   │   │   │   │   ├── Accordion.tsx  # Base component
    │   │   │   │   │   │   ├── Accordion.types.ts  # Shared types
    │   │   │   │   │   │   └── Accordion.web.tsx  # Web implementation
    │   │   │   │   │   ├── Breadcrumb/
    │   │   │   │   │   │   ├── Breadcrumb.native.tsx#  - Mobile needs breadcrumb
    │   │   │   │   │   │   ├── Breadcrumb.tsx  # Base component
    │   │   │   │   │   │   ├── Breadcrumb.types.ts  # Shared types
    │   │   │   │   │   │   └── Breadcrumb.web.tsx  # Web implementation
    │   │   │   │   │   ├── Drawer/
    │   │   │   │   │   │   ├── Drawer.native.tsx  # - Mobile needs drawer (BottomSheet)
    │   │   │   │   │   │   ├── Drawer.tsx  # Base component
    │   │   │   │   │   │   ├── Drawer.types.ts  # Shared types
    │   │   │   │   │   │   └── Drawer.web.tsx  # Web implementation
    │   │   │   │   │   └── Tabs/
    │   │   │   │   │       ├── Tabs.native.tsx  # ⚠️ EXISTS but may need enhancement
    │   │   │   │   │       ├── Tabs.tsx  # Base component
    │   │   │   │   │       ├── Tabs.types.ts  # Shared types
    │   │   │   │   │       └── Tabs.web.tsx  # Web implementation
    │   │   │   │   └── Overlay/
    │   │   │   │       ├── Popover/
    │   │   │   │       │   ├── Popover.native.tsx  # - Mobile needs popover
    │   │   │   │       │   ├── Popover.tsx  # Base component
    │   │   │   │       │   ├── Popover.types.ts  # Shared types
    │   │   │   │       │   └── Popover.web.tsx  # Web implementation
    │   │   │   │       └── Tooltip/
    │   │   │   │           ├── Tooltip.native.tsx  # - Mobile needs tooltip (long press)
    │   │   │   │           ├── Tooltip.tsx  # Base component
    │   │   │   │           ├── Tooltip.types.ts  # Shared types
    │   │   │   │           └── Tooltip.web.tsx  # Web implementation
    │   │   │   ├── organisms/
    │   │   │   │   └── DataDisplay/
    │   │   │   │       ├── DataGrid/
    │   │   │   │       │   ├── DataGrid.native.tsx  # - Mobile needs data grid
    │   │   │   │       │   ├── DataGrid.tsx  # Base component
    │   │   │   │       │   ├── DataGrid.types.ts  # Shared types
    │   │   │   │       │   └── DataGrid.web.tsx  # Web implementation
    │   │   │   │       ├── KanbanBoard/
    │   │   │   │       │   ├── KanbanBoard.native.tsx  # - Mobile needs kanban
    │   │   │   │       │   ├── KanbanBoard.tsx  # Base component
    │   │   │   │       │   ├── KanbanBoard.types.ts  # Shared types
    │   │   │   │       │   └── KanbanBoard.web.tsx  # Web implementation
    │   │   │   │       └── Table/
    │   │   │   │           ├── Table.native.tsx  # - Mobile needs table (FlashList)
    │   │   │   │           ├── Table.tsx  # Base component
    │   │   │   │           ├── Table.types.ts  # Shared types
    │   │   │   │           └── Table.web.tsx  # Web implementation
    │   │   │   ├── tracking/
    │   │   │   │   ├── TimesheetTable.native.tsx
    │   │   │   │   ├── TimesheetTable.tsx  # Timesheet grid
    │   │   │   │   ├── TimesheetTable.web.tsx
    │   │   │   │   ├── TimeTracker.native.tsx
    │   │   │   │   ├── TimeTracker.tsx  # Time tracking widget
    │   │   │   │   ├── TimeTracker.web.tsx
    │   │   │   │   ├── WorkDiaryEntry.native.tsx
    │   │   │   │   ├── WorkDiaryEntry.tsx  # Work diary card
    │   │   │   │   └── WorkDiaryEntry.web.tsx
    │   │   │   ├── video/
    │   │   │   │   ├── VideoPlayer.native.tsx
    │   │   │   │   ├── VideoPlayer.tsx  # Video player
    │   │   │   │   ├── VideoPlayer.web.tsx
    │   │   │   │   ├── VideoUploader.native.tsx  # Video upload
    │   │   │   │   ├── VideoUploader.tsx  # Video upload
    │   │   │   │   └── VideoUploader.web.tsx
    │   │   │   └── visualization/
    │   │   │       ├── Gantt/
    │   │   │       │   ├── GanttChart.native.tsx
    │   │   │       │   ├── GanttChart.tsx
    │   │   │       │   └── GanttChart.web.tsx
    │   │   │       ├── Heatmap/
    │   │   │       │   ├── Heatmap.native.tsx
    │   │   │       │   ├── Heatmap.tsx
    │   │   │       │   └── Heatmap.web.tsx
    │   │   │       ├── Kanban/
    │   │   │       │   ├── KanbanBoard.native.tsx
    │   │   │       │   ├── KanbanBoard.tsx
    │   │   │       │   └── KanbanBoard.web.tsx
    │   │   │       └── OrgChart/
    │   │   │           ├── OrganizationChart.native.tsx
    │   │   │           ├── OrganizationChart.tsx
    │   │   │           └── OrganizationChart.web.tsx
    │   │   ├── package.json
    │   │   ├── README.md
    │   │   └── tsconfig.json
    │   └── │
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
    │   ├── migration/  # MIGRATION SCRIPTS
    │   │   ├── migrate-hooks.sh  # Hook migration script
    │   │   ├── validate-structure.sh  # Structure validation script
    │   │   └── │
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
    ├── .env.example  # Environment variables template
    ├── .env.local  # Local environment variables
    ├── .env.local.example  # Local overrides template
    ├── .env.production  # Production environment
    ├── .env.staging  # Staging environment
    ├── .eslintrc.json  # Root ESLint config
    ├── .gitignore  # Git ignore patterns
    ├── .nvmrc  # Node.js version (v20.x)
    ├── .prettierignore  # Prettier ignore patterns
    ├── .prettierrc  # Prettier configuration
    ├── error.tsx  # Global error boundary
    ├── globals.css  # Global styles
    ├── jest.config.js  # Root Jest configuration
    ├── layout.tsx  # Root layout
    ├── loading.tsx  # Global loading state
    ├── middleware.ts  # Next.js middleware
    ├── next.config.js  # Next.js configuration
    ├── not-found.tsx  # 404 page
    ├── package.json  # Root package (workspace manager)
    ├── page.tsx  # Root page (redirects to /[locale])
    ├── pnpm-lock.yaml  # Locked dependencies
    ├── pnpm-workspace.yaml  # pnpm workspace configuration
    ├── postcss.config.js  # PostCSS configuration
    ├── README.md  # Root README
    ├── tailwind.config.js  # Tailwind configuration
    ├── tsconfig.base.json  # Base TypeScript configuration
    ├── tsconfig.json  # Root TypeScript config
    ├── turbo.json  # Turborepo pipeline configuration
    └── │
```