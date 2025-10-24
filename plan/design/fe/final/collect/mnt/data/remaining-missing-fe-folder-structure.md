fe/
├── .github/
│   ├── actions/
│   │   ├── cache-dependencies/
│   │   │   └── action.yml
│   │   ├── deploy-preview/
│   │   │   └── action.yml
│   │   └── setup-node/
│   │       └── action.yml
│   └── workflows/
│       ├── accessibility.yml  # A11y checks
│       ├── cd-mobile.yml  # Mobile deployment
│       │   # - EAS build
│       │   # - Submit to stores
│       ├── cd-web.yml  # Web deployment
│       │   # - Build Next.js production
│       │   # - Deploy to K8s
│       ├── ci.yml  # Continuous integration
│       ├── dependency-review.yml  # Dependency checks
│       ├── lighthouse.yml  # Performance audits
│       ├── release.yml  # Release automation
│       └── security.yml  # Security scanning
├── .husky/
│   ├── pre-commit  # Runs linting, type checking
│   │   # Validates no console.log in prod code
│   └── pre-push
├── .vscode/
│   ├── extensions.json
│   ├── launch.json
│   └── settings.json
├── apps/
│   ├── mobile/
│   │   ├── app/
│   │   │   ├── (auth)/
│   │   │   │   ├── _layout.tsx
│   │   │   │   ├── callback.tsx
│   │   │   │   ├── login.tsx
│   │   │   │   └── register.tsx
│   │   │   ├── +not-found.tsx
│   │   │   ├── _layout.tsx
│   │   │   ├── index.tsx
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
│   │   │   ├── onboarding/
│   │   │   │   ├── _layout.tsx  # Onboarding layout
│   │   │   │   ├── complete.tsx  # Onboarding complete
│   │   │   │   ├── features.tsx  # Features showcase
│   │   │   │   ├── permissions.tsx  # Request permissions
│   │   │   │   │   # - Notifications
│   │   │   │   │   # - Camera
│   │   │   │   │   # - Location
│   │   │   │   └── profile-setup.tsx  # Quick profile setup
│   │   │   │       # BE: users-be/profile
│   │   │   │       # POST /v1/users/me/profile
│   │   │   └── quick-actions/
│   │   │       ├── quick-apply.tsx  # Quick proposal submission
│   │   │       │   # - Minimal form
│   │   │       │   # - Draft save
│   │   │       │   # BE: proposals-be/proposal
│   │   │       │   # POST /v1/proposals/quick
│   │   │       ├── quick-message.tsx  # Quick message
│   │   │       │   # - Contact selection
│   │   │       │   # - Quick send
│   │   │       │   # BE: communications-be/message
│   │   │       │   # POST /v1/messages
│   │   │       └── quick-time-entry.tsx  # Quick time logging
│   │   │           # - Current contract
│   │   │           # - Duration input
│   │   │           # BE: contracts-be/workdiary
│   │   │           # POST /v1/contracts/{contract_id}/workdiary/quick
│   │   └── src/
│   └── web/
│       ├── e2e/
│       │   ├── auth/
│       │   │   ├── login.spec.ts
│       │   │   ├── password-reset.spec.ts
│       │   │   └── register.spec.ts
│       │   ├── contracts/
│       │   │   ├── contract-creation.spec.ts
│       │   │   └── contract-execution.spec.ts
│       │   ├── payments/
│       │   │   ├── escrow.spec.ts
│       │   │   └── payment-methods.spec.ts
│       │   ├── proposals/
│       │   │   ├── proposal-review.spec.ts
│       │   │   └── proposal-submission.spec.ts
│       │   ├── jobs/
│       │   │   ├── job-application.spec.ts
│       │   │   ├── job-posting.spec.ts
│       │   │   └── job-search.spec.ts
│       │   └── fixtures/
│       │       ├── contracts.json
│       │       ├── jobs.json
│       │       └── users.json
│       ├── middleware.ts  # Next.js middleware for security headers
│       ├── security/
│       │   ├── cors.ts  # CORS configuration
│       │   ├── csp.ts  # Content Security Policy
│       │   └── headers.ts  # Security headers config
│       └── src/
│           └── app/
│               └── [locale]/
│                   ├── (admin)/
│                   │   └── admin/
│                   │       ├── audit-logs/
│                   │       │   ├── [logId]/
│                   │       │   │   └── page.tsx  # Audit log details
│                   │       │   │       # BE: utility-be/audit
│                   │       │   │       # GET /v1/admin/audit-logs/{log_id}
│                   │       │   └── page.tsx  # Audit logs viewer
│                   │       │       # - System-wide audit trail
│                   │       │       # - Filter by entity/action
│                   │       │       # - Export capabilities
│                   │       │       # BE: utility-be/audit
│                   │       │       # GET /v1/admin/audit-logs
│                   │       ├── business-verification/
│                   │       │   ├── [verificationId]/
│                   │       │   │   ├── review/
│                   │       │   │   │   └── page.tsx  # Review business verification
│                   │       │   │   │       # - Company evidence review
│                   │       │   │   │       # - Decision making
│                   │       │   │   │       # BE: admin-be/business_verification
│                   │       │   │   │       # POST /v1/admin/business-verification/{verification_id}/review
│                   │       │   │   ├── reverify/
│                   │       │   │   │   └── page.tsx  # Request reverification
│                   │       │   │   │       # BE: admin-be/business_verification
│                   │       │   │   │       # POST /v1/admin/business-verification/{verification_id}/reverify
│                   │       │   │   └── page.tsx  # Verification details
│                   │       │   │       # BE: admin-be/business_verification
│                   │       │   │       # GET /v1/admin/business-verification/{verification_id}
│                   │       │   ├── pending/
│                   │       │   │   └── page.tsx  # Pending verifications
│                   │       │   │       # BE: admin-be/business_verification
│                   │       │   │       # GET /v1/admin/business-verification?status=pending
│                   │       │   └── page.tsx  # Business verification dashboard
│                   │       │       # BE: admin-be/business_verification
│                   │       │       # GET /v1/admin/business-verification
│                   │       ├── communications-ops/
│                   │       │   ├── broadcasts/
│                   │       │   │   ├── [broadcastId]/
│                   │       │   │   │   ├── analytics/
│                   │       │   │   │   │   └── page.tsx  # Broadcast analytics
│                   │       │   │   │   │       # BE: communications-be/broadcast
│                   │       │   │   │   │       # GET /v1/admin/broadcasts/{broadcast_id}/analytics
│                   │       │   │   │   ├── edit/
│                   │       │   │   │   │   └── page.tsx  # Edit broadcast
│                   │       │   │   │   │       # BE: communications-be/broadcast
│                   │       │   │   │   │       # PUT /v1/admin/broadcasts/{broadcast_id}
│                   │       │   │   │   ├── schedule/
│                   │       │   │   │   │   └── page.tsx  # Schedule broadcast
│                   │       │   │   │   │       # BE: communications-be/broadcast
│                   │       │   │   │   │       # POST /v1/admin/broadcasts/{broadcast_id}/schedule
│                   │       │   │   │   └── page.tsx  # Broadcast details
│                   │       │   │   │       # BE: communications-be/broadcast
│                   │       │   │   │       # GET /v1/admin/broadcasts/{broadcast_id}
│                   │       │   │   ├── create/
│                   │       │   │   │   └── page.tsx  # Create broadcast
│                   │       │   │   │       # - Target audience
│                   │       │   │   │       # - Message content
│                   │       │   │   │       # - Delivery channels
│                   │       │   │   │       # BE: communications-be/broadcast
│                   │       │   │   │       # POST /v1/admin/broadcasts
│                   │       │   │   └── page.tsx  # Broadcasts dashboard
│                   │       │   │       # BE: communications-be/broadcast
│                   │       │   │       # GET /v1/admin/broadcasts
│                   │       │   ├── campaigns/
│                   │       │   │   ├── [campaignId]/
│                   │       │   │   │   ├── analytics/
│                   │       │   │   │   │   └── page.tsx  # Campaign analytics
│                   │       │   │   │   │       # BE: communications-be/campaign
│                   │       │   │   │   │       # GET /v1/admin/campaigns/{campaign_id}/analytics
│                   │       │   │   │   ├── edit/
│                   │       │   │   │   │   └── page.tsx  # Edit campaign
│                   │       │   │   │   │       # BE: communications-be/campaign
│                   │       │   │   │   │       # PUT /v1/admin/campaigns/{campaign_id}
│                   │       │   │   │   └── page.tsx  # Campaign details
│                   │       │   │   │       # BE: communications-be/campaign
│                   │       │   │   │       # GET /v1/admin/campaigns/{campaign_id}
│                   │       │   │   ├── create/
│                   │       │   │   │   └── page.tsx  # Create campaign
│                   │       │   │   │       # BE: communications-be/campaign
│                   │       │   │   │       # POST /v1/admin/campaigns
│                   │       │   │   └── page.tsx  # Campaigns dashboard
│                   │       │   │       # BE: communications-be/campaign
│                   │       │   │       # GET /v1/admin/campaigns
│                   │       │   ├── rate-limits/
│                   │       │   │   └── page.tsx  # Communication rate limits
│                   │       │   │       # - Configure limits
│                   │       │   │       # - Monitor usage
│                   │       │   │       # BE: communications-be/rate_limit
│                   │       │   │       # GET /v1/admin/rate-limits
│                   │       │   │       # PUT /v1/admin/rate-limits
│                   │       │   └── templates/
│                   │       │       ├── [templateId]/
│                   │       │       │   ├── edit/
│                   │       │       │   │   └── page.tsx  # Edit template
│                   │       │       │   │       # BE: communications-be/template
│                   │       │       │   │       # PUT /v1/admin/templates/{template_id}
│                   │       │       │   ├── logs/
│                   │       │       │   │   └── page.tsx  # Webhook delivery logs
│                   │       │       │   │       # BE: communications-be/template
│                   │       │       │   │       # GET /v1/developer/webhooks/{webhook_id}/logs
│                   │       │       │   ├── test/
│                   │       │       │   │   └── page.tsx  # Test template
│                   │       │       │   │       # BE: communications-be/template
│                   │       │       │   │       # POST /v1/admin/templates/{template_id}/test
│                   │       │       │   └── page.tsx  # Template details
│                   │       │       │       # BE: communications-be/template
│                   │       │       │       # GET /v1/admin/templates/{template_id}
│                   │       │       ├── create/
│                   │       │       │   └── page.tsx  # Create template
│                   │       │       │       # BE: communications-be/template
│                   │       │       │       # POST /v1/admin/templates
│                   │       │       └── page.tsx  # Templates dashboard
│                   │       │           # BE: communications-be/template
│                   │       │           # GET /v1/admin/templates
│                   │       ├── dashboard/
│                   │       │   └── page.tsx  # Admin dashboard
│                   │       │       # - Key metrics
│                   │       │       # - Pending actions
│                   │       │       # - System alerts
│                   │       │       # BE: admin-be/dashboard
│                   │       │       # GET /v1/admin/dashboard
│                   │       ├── financial-ops/
│                   │       │   ├── disputes/
│                   │       │   │   ├── [disputeId]/
│                   │       │   │   │   ├── mediate/
│                   │       │   │   │   │   └── page.tsx  # Mediate payment dispute
│                   │       │   │   │   │       # BE: admin-be/dispute_resolution
│                   │       │   │   │   │       # POST /v1/admin/financial-disputes/{dispute_id}/mediate
│                   │       │   │   │   └── page.tsx  # Dispute details
│                   │       │   │   │       # BE: admin-be/dispute_resolution
│                   │       │   │   │       # GET /v1/admin/financial-disputes/{dispute_id}
│                   │       │   │   └── page.tsx  # Financial disputes
│                   │       │   │       # BE: admin-be/dispute_resolution
│                   │       │   │       # GET /v1/admin/financial-disputes
│                   │       │   ├── payouts/
│                   │       │   │   ├── [payoutId]/
│                   │       │   │   │   ├── review/
│                   │       │   │   │   │   └── page.tsx  # Review payout
│                   │       │   │   │   │       # BE: financial-be/payout
│                   │       │   │   │   │       # POST /v1/admin/payouts/{payout_id}/review
│                   │       │   │   │   ├── retry/
│                   │       │   │   │   │   └── page.tsx  # Retry failed payout
│                   │       │   │   │   │       # BE: financial-be/payout
│                   │       │   │   │   │       # POST /v1/admin/payouts/{payout_id}/retry
│                   │       │   │   │   └── page.tsx  # Payout details
│                   │       │   │   │       # BE: financial-be/payout
│                   │       │   │   │       # GET /v1/admin/payouts/{payout_id}
│                   │       │   │   ├── failed/
│                   │       │   │   │   └── page.tsx  # Failed payouts
│                   │       │   │   │       # BE: financial-be/payout
│                   │       │   │   │       # GET /v1/admin/payouts?status=failed
│                   │       │   │   ├── pending/
│                   │       │   │   │   └── page.tsx  # Pending payouts
│                   │       │   │   │       # BE: financial-be/payout
│                   │       │   │   │       # GET /v1/admin/payouts?status=pending
│                   │       │   │   └── page.tsx  # Payouts dashboard
│                   │       │   │       # BE: financial-be/payout
│                   │       │   │       # GET /v1/admin/payouts
│                   │       │   ├── reconciliation/
│                   │       │   │   ├── [reconciliationId]/
│                   │       │   │   │   ├── resolve/
│                   │       │   │   │   │   └── page.tsx  # Resolve reconciliation
│                   │       │   │   │   │       # BE: financial-be/reconciliation
│                   │       │   │   │   │       # POST /v1/admin/reconciliation/{reconciliation_id}/resolve
│                   │       │   │   │   └── page.tsx  # Reconciliation details
│                   │       │   │   │       # BE: financial-be/reconciliation
│                   │       │   │   │       # GET /v1/admin/reconciliation/{reconciliation_id}
│                   │       │   │   ├── pending/
│                   │       │   │   │   └── page.tsx  # Pending reconciliations
│                   │       │   │   │       # BE: financial-be/reconciliation
│                   │       │   │   │       # GET /v1/admin/reconciliation?status=pending
│                   │       │   │   └── page.tsx  # Reconciliation dashboard
│                   │       │   │       # BE: financial-be/reconciliation
│                   │       │   │       # GET /v1/admin/reconciliation
│                   │       │   └── tax-forms/
│                   │       │       ├── [formId]/
│                   │       │       │   ├── review/
│                   │       │       │   │   └── page.tsx  # Review tax form
│                   │       │       │   │       # BE: financial-be/tax
│                   │       │       │   │       # POST /v1/admin/tax-forms/{form_id}/review
│                   │       │       │   └── page.tsx  # Tax form details
│                   │       │       │       # BE: financial-be/tax
│                   │       │       │       # GET /v1/admin/tax-forms/{form_id}
│                   │       │       ├── generate/
│                   │       │       │   └── page.tsx  # Generate tax forms
│                   │       │       │       # - Bulk 1099 generation
│                   │       │       │       # - Tax year selection
│                   │       │       │       # BE: financial-be/tax
│                   │       │       │       # POST /v1/admin/tax-forms/generate
│                   │       │       └── page.tsx  # Tax forms dashboard
│                   │       │           # BE: financial-be/tax
│                   │       │           # GET /v1/admin/tax-forms
│                   │       ├── goodwill-credits/
│                   │       │   ├── [creditId]/
│                   │       │   │   ├── approve/
│                   │       │   │   │   └── page.tsx  # Approve goodwill credit
│                   │       │   │   │       # BE: admin-be/goodwill_credit
│                   │       │   │   │       # POST /v1/admin/goodwill-credits/{credit_id}/approve
│                   │       │   │   └── page.tsx  # Goodwill credit details
│                   │       │   │       # BE: admin-be/goodwill_credit
│                   │       │   │       # GET /v1/admin/goodwill-credits/{credit_id}
│                   │       │   ├── issue/
│                   │       │   │   └── page.tsx  # Issue goodwill credit
│                   │       │   │       # - User selection
│                   │       │   │       # - Amount and reason
│                   │       │   │       # BE: admin-be/goodwill_credit
│                   │       │   │       # POST /v1/admin/goodwill-credits
│                   │       │   └── page.tsx  # Goodwill credits dashboard
│                   │       │       # BE: admin-be/goodwill_credit
│                   │       │       # GET /v1/admin/goodwill-credits
│                   │       ├── kyc-cases/
│                   │       │   ├── [caseId]/
│                   │       │   │   ├── documents/
│                   │       │   │   │   └── page.tsx  # Case documents viewer
│                   │       │   │   │       # BE: admin-be/kyc_case, storage-be/asset
│                   │       │   │   │       # GET /v1/admin/kyc-cases/{case_id}/documents
│                   │       │   │   ├── reopen/
│                   │       │   │   │   └── page.tsx  # Reopen KYC case
│                   │       │   │   │       # BE: admin-be/kyc_case
│                   │       │   │   │       # POST /v1/admin/kyc-cases/{case_id}/reopen
│                   │       │   │   ├── review/
│                   │       │   │   │   └── page.tsx  # Review KYC case
│                   │       │   │   │       # - Document verification
│                   │       │   │   │       # - Approve/reject/escalate
│                   │       │   │   │       # BE: admin-be/kyc_case
│                   │       │   │   │       # POST /v1/admin/kyc-cases/{case_id}/review
│                   │       │   │   └── page.tsx  # KYC case details
│                   │       │   │       # BE: admin-be/kyc_case
│                   │       │   │       # GET /v1/admin/kyc-cases/{case_id}
│                   │       │   ├── pending/
│                   │       │   │   └── page.tsx  # KYC queue
│                   │       │   │       # - Prioritization
│                   │       │   │       # - Assignment
│                   │       │   │       # BE: admin-be/kyc_case
│                   │       │   │       # GET /v1/admin/kyc-cases/queue
│                   │       │   └── page.tsx  # KYC cases dashboard
│                   │       │       # BE: admin-be/kyc_case
│                   │       │       # GET /v1/admin/kyc-cases
│                   │       ├── moderation/
│                   │       │   ├── actions/
│                   │       │   │   ├── [actionId]/
│                   │       │   │   │   ├── appeal/
│                   │       │   │   │   │   └── page.tsx  # Review appeal
│                   │       │   │   │   │       # BE: admin-be/moderation
│                   │       │   │   │   │       # POST /v1/admin/moderation/actions/{action_id}/appeal
│                   │       │   │   │   └── page.tsx  # Action details
│                   │       │   │   │       # BE: admin-be/moderation
│                   │       │   │   │       # GET /v1/admin/moderation/actions/{action_id}
│                   │       │   │   └── page.tsx  # Moderation actions
│                   │       │   │       # - Warnings
│                   │       │   │       # - Suspensions
│                   │       │   │       # - Bans
│                   │       │   │       # BE: admin-be/moderation
│                   │       │   │       # GET /v1/admin/moderation/actions
│                   │       │   ├── patterns/
│                   │       │   │   └── page.tsx  # Abuse patterns detection
│                   │       │   │       # - Pattern analysis
│                   │       │   │       # - Risk scoring
│                   │       │   │       # BE: admin-be/moderation
│                   │       │   │       # GET /v1/admin/moderation/patterns
│                   │       │   └── reports/
│                   │       │       ├── [reportId]/
│                   │       │       │   ├── review/
│                   │       │       │   │   └── page.tsx  # Review report
│                   │       │       │   │       # - Content review
│                   │       │       │   │       # - Take action
│                   │       │       │   │       # BE: admin-be/moderation
│                   │       │       │   │       # POST /v1/admin/moderation/reports/{report_id}/review
│                   │       │       │   └── page.tsx  # Report details
│                   │       │       │       # BE: admin-be/moderation
│                   │       │       │       # GET /v1/admin/moderation/reports/{report_id}
│                   │       │       ├── queue/
│                   │       │       │   └── page.tsx  # Moderation queue
│                   │       │       │       # BE: admin-be/moderation
│                   │       │       │       # GET /v1/admin/moderation/reports/queue
│                   │       │       └── page.tsx  # Reports dashboard
│                   │       │           # BE: admin-be/moderation
│                   │       │           # GET /v1/admin/moderation/reports
│                   │       ├── search-quality/
│                   │       │   ├── blacklists/
│                   │       │   │   ├── [blacklistId]/
│                   │       │   │   │   └── page.tsx  # Blacklist entry details
│                   │       │   │   │       # BE: search-be/admin
│                   │       │   │   │       # GET /v1/admin/search/blacklists/{blacklist_id}
│                   │       │   │   ├── add/
│                   │       │   │   │   └── page.tsx  # Add blacklist entry
│                   │       │   │   │       # BE: search-be/admin
│                   │       │   │   │       # POST /v1/admin/search/blacklists
│                   │       │   │   └── page.tsx  # Blacklist management
│                   │       │   │       # BE: search-be/admin
│                   │       │   │       # GET /v1/admin/search/blacklists
│                   │       │   ├── boosts/
│                   │       │   │   ├── [boostId]/
│                   │       │   │   │   ├── edit/
│                   │       │   │   │   │   └── page.tsx  # Edit boost rule
│                   │       │   │   │   │       # BE: search-be/admin
│                   │       │   │   │   │       # PUT /v1/admin/search/boosts/{boost_id}
│                   │       │   │   │   └── page.tsx  # Boost details
│                   │       │   │   │       # BE: search-be/admin
│                   │       │   │   │       # GET /v1/admin/search/boosts/{boost_id}
│                   │       │   │   ├── create/
│                   │       │   │   │   └── page.tsx  # Create boost rule
│                   │       │   │   │       # BE: search-be/admin
│                   │       │   │   │       # POST /v1/admin/search/boosts
│                   │       │   │   └── page.tsx  # Boost rules management
│                   │       │   │       # BE: search-be/admin
│                   │       │   │       # GET /v1/admin/search/boosts
│                   │       │   ├── reindex/
│                   │       │   │   └── page.tsx  # Reindex operations
│                   │       │   │       # - Full reindex
│                   │       │   │       # - Selective reindex
│                   │       │   │       # - Progress monitoring
│                   │       │   │       # BE: search-be/admin
│                   │       │   │       # POST /v1/admin/search/reindex
│                   │       │   └── synonyms/
│                   │       │       ├── [synonymId]/
│                   │       │       │   ├── edit/
│                   │       │       │   │   └── page.tsx  # Edit synonym
│                   │       │       │   │       # BE: search-be/admin
│                   │       │       │   │       # PUT /v1/admin/search/synonyms/{synonym_id}
│                   │       │       │   └── page.tsx  # Synonym details
│                   │       │       │       # BE: search-be/admin
│                   │       │       │       # GET /v1/admin/search/synonyms/{synonym_id}
│                   │       │       ├── create/
│                   │       │       │   └── page.tsx  # Create synonym
│                   │       │       │       # BE: search-be/admin
│                   │       │       │       # POST /v1/admin/search/synonyms
│                   │       │       └── page.tsx  # Synonyms management
│                   │       │           # BE: search-be/admin
│                   │       │           # GET /v1/admin/search/synonyms
│                   │       ├── system/
│                   │       │   ├── configuration/
│                   │       │   │   └── page.tsx  # System configuration
│                   │       │   │       # - Global settings
│                   │       │   │       # - Environment variables
│                   │       │   │       # BE: utility-be/config
│                   │       │   │       # GET /v1/admin/system/config
│                   │       │   │       # PUT /v1/admin/system/config
│                   │       │   ├── experiments/
│                   │       │   │   ├── [experimentId]/
│                   │       │   │   │   ├── edit/
│                   │       │   │   │   │   └── page.tsx  # Edit experiment
│                   │       │   │   │   │       # BE: utility-be/experiments
│                   │       │   │   │   │       # PUT /v1/admin/experiments/{experiment_id}
│                   │       │   │   │   ├── results/
│                   │       │   │   │   │   └── page.tsx  # Experiment results
│                   │       │   │   │   │       # BE: utility-be/experiments
│                   │       │   │   │   │       # GET /v1/admin/experiments/{experiment_id}/results
│                   │       │   │   │   └── page.tsx  # Experiment details
│                   │       │   │   │       # BE: utility-be/experiments
│                   │       │   │   │       # GET /v1/admin/experiments/{experiment_id}
│                   │       │   │   ├── create/
│                   │       │   │   │   └── page.tsx  # Create A/B experiment
│                   │       │   │   │       # BE: utility-be/experiments
│                   │       │   │   │       # POST /v1/admin/experiments
│                   │       │   │   └── page.tsx  # Experiments dashboard
│                   │       │   │       # BE: utility-be/experiments
│                   │       │   │       # GET /v1/admin/experiments
│                   │       │   ├── feature-flags/
│                   │       │   │   ├── [flagId]/
│                   │       │   │   │   ├── edit/
│                   │       │   │   │   │   └── page.tsx  # Edit feature flag
│                   │       │   │   │   │       # BE: utility-be/feature_flags
│                   │       │   │   │   │       # PUT /v1/admin/feature-flags/{flag_id}
│                   │       │   │   │   ├── rollout/
│                   │       │   │   │   │   └── page.tsx  # Manage rollout
│                   │       │   │   │   │       # BE: utility-be/feature_flags
│                   │       │   │   │   │       # POST /v1/admin/feature-flags/{flag_id}/rollout
│                   │       │   │   │   └── page.tsx  # Feature flag details
│                   │       │   │   │       # BE: utility-be/feature_flags
│                   │       │   │   │       # GET /v1/admin/feature-flags/{flag_id}
│                   │       │   │   ├── create/
│                   │       │   │   │   └── page.tsx  # Create feature flag
│                   │       │   │   │       # BE: utility-be/feature_flags
│                   │       │   │   │       # POST /v1/admin/feature-flags
│                   │       │   │   └── page.tsx  # Feature flags dashboard
│                   │       │   │       # BE: utility-be/feature_flags
│                   │       │   │       # GET /v1/admin/feature-flags
│                   │       │   ├── health/
│                   │       │   │   └── page.tsx  # System health dashboard
│                   │       │   │       # - Service status
│                   │       │   │       # - Resource usage
│                   │       │   │       # - Error rates
│                   │       │   │       # BE: utility-be/health
│                   │       │   │       # GET /v1/admin/system/health
│                   │       │   ├── logs/
│                   │       │   │   └── page.tsx  # System logs viewer
│                   │       │   │       # - Real-time logs
│                   │       │   │       # - Search and filter
│                   │       │   │       # BE: utility-be/logs
│                   │       │   │       # GET /v1/admin/system/logs
│                   │       │   └── metrics/
│                   │       │       └── page.tsx  # System metrics
│                   │       │           # - Performance metrics
│                   │       │           # - Custom dashboards
│                   │       │           # BE: utility-be/metrics
│                   │       │           # GET /v1/admin/system/metrics
│                   │       ├── two-person-rules/
│                   │       │   ├── [ruleId]/
│                   │       │   │   ├── approve/
│                   │       │   │   │   └── page.tsx  # Approve rule change
│                   │       │   │   │       # BE: admin-be/change_approval
│                   │       │   │   │       # POST /v1/admin/two-person-rules/{rule_id}/approve
│                   │       │   │   └── page.tsx  # Rule details
│                   │       │   │       # BE: admin-be/change_approval
│                   │       │   │       # GET /v1/admin/two-person-rules/{rule_id}
│                   │       │   ├── pending/
│                   │       │   │   └── page.tsx  # Pending approvals
│                   │       │   │       # BE: admin-be/change_approval
│                   │       │   │       # GET /v1/admin/two-person-rules?status=pending
│                   │       │   └── page.tsx  # Two-person rules dashboard
│                   │       │       # BE: admin-be/change_approval
│                   │       │       # GET /v1/admin/two-person-rules
│                   │       └── break-glass/
│                   │           ├── active/
│                   │           │   └── page.tsx  # Active admin sessions
│                   │           │       # - Time-boxed access monitoring
│                   │           │       # - Force termination
│                   │           │       # BE: admin-be/admin_session
│                   │           │       # GET /v1/admin/break-glass/active
│                   │           ├── approve/
│                   │           │   └── page.tsx  # Approve break-glass requests
│                   │           │       # - Two-person rule
│                   │           │       # - Request review
│                   │           │       # BE: admin-be/admin_session
│                   │           │       # POST /v1/admin/break-glass/approve/{request_id}
│                   │           └── request/
│                   │               └── page.tsx  # Request break-glass access
│                   │                   # - Justification
│                   │                   # - Duration request
│                   │                   # BE: admin-be/admin_session
│                   │                   # POST /v1/admin/break-glass/request
│                   └── (dashboard)/
│                       ├── contracts/
│                       │   ├── [contractId]/
│                       │   │   ├── audit-trail/
│                       │   │   │   └── page.tsx  # Complete contract audit trail
│                       │   │   │       # - All changes
│                       │   │   │       # - Approval history
│                       │   │   │       # - Access logs
│                       │   │   │       # BE: contracts-be/contract, utility-be/audit
│                       │   │   │       # GET /v1/contracts/{contract_id}/audit-trail
│                       │   │   ├── change-orders/
│                       │   │   │   ├── [orderId]/
│                       │   │   │   │   ├── approve/
│                       │   │   │   │   │   └── page.tsx  # Approve change order
│                       │   │   │   │   │       # BE: contracts-be/change_order
│                       │   │   │   │   │       # POST /v1/contracts/{contract_id}/change-orders/{order_id}/approve
│                       │   │   │   │   ├── reject/
│                       │   │   │   │   │   └── page.tsx  # Reject change order
│                       │   │   │   │   │       # BE: contracts-be/change_order
│                       │   │   │   │   │       # POST /v1/contracts/{contract_id}/change-orders/{order_id}/reject
│                       │   │   │   │   └── page.tsx  # Change order details
│                       │   │   │   │       # BE: contracts-be/change_order
│                       │   │   │   │       # GET /v1/contracts/{contract_id}/change-orders/{order_id}
│                       │   │   │   ├── create/
│                       │   │   │   │   └── page.tsx  # Create change order
│                       │   │   │   │       # - Scope modifications
│                       │   │   │   │       # - Budget adjustments
│                       │   │   │   │       # - Timeline changes
│                       │   │   │   │       # BE: contracts-be/change_order
│                       │   │   │   │       # POST /v1/contracts/{contract_id}/change-orders
│                       │   │   │   └── page.tsx  # Change orders list
│                       │   │   │       # BE: contracts-be/change_order
│                       │   │   │       # GET /v1/contracts/{contract_id}/change-orders
│                       │   │   └── compliance/
│                       │   │       └── page.tsx  # Contract compliance tracking
│                       │   │           # - KPIs monitoring
│                       │   │           # - SLA compliance
│                       │   │           # - Penalty tracking
│                       │   │           # BE: contracts-be/compliance
│                       │   │           # GET /v1/contracts/{contract_id}/compliance
│                       │   ├── benchmarking/
│                       │   │   └── page.tsx  # Contract performance benchmarking
│                       │   │       # - Industry comparisons
│                       │   │       # - Performance metrics
│                       │   │       # - Best practices
│                       │   │       # BE: contracts-be/analytics
│                       │   │       # GET /v1/contracts/benchmarking
│                       │   └── calendar/
│                       │       └── page.tsx  # Contracts calendar view
│                       │           # - Milestone timeline
│                       │           # - Payment schedules
│                       │           # - Deliverable deadlines
│                       │           # BE: contracts-be/contract, contracts-be/milestone
│                       │           # GET /v1/contracts/calendar
│                       ├── financial/
│                       │   ├── chargebacks/
│                       │   │   ├── [chargebackId]/
│                       │   │   │   ├── respond/
│                       │   │   │   │   └── page.tsx  # Respond to chargeback
│                       │   │   │   │       # - Upload evidence
│                       │   │   │   │       # - Add defense statement
│                       │   │   │   │       # BE: financial-be/chargeback
│                       │   │   │   │       # POST /v1/financial/chargebacks/{chargeback_id}/respond
│                       │   │   │   └── page.tsx  # Chargeback details
│                       │   │   │       # BE: financial-be/chargeback
│                       │   │   │       # GET /v1/financial/chargebacks/{chargeback_id}
│                       │   │   └── page.tsx  # Chargebacks list
│                       │   │       # - Active disputes
│                       │   │       # - Resolution status
│                       │   │       # BE: financial-be/chargeback
│                       │   │       # GET /v1/financial/chargebacks
│                       │   ├── forecasting/
│                       │   │   └── page.tsx  # Revenue/expense forecasting
│                       │   │       # - Projected earnings
│                       │   │       # - Cash flow predictions
│                       │   │       # - Scenario planning
│                       │   │       # BE: financial-be/analytics
│                       │   │       # GET /v1/financial/forecasts
│                       │   ├── payout-methods/
│                       │   │   ├── [methodId]/
│                       │   │   │   ├── edit/
│                       │   │   │   │   └── page.tsx  # Edit payout method
│                       │   │   │   │       # BE: financial-be/payout_method
│                       │   │   │   │       # PUT /v1/financial/payout-methods/{method_id}
│                       │   │   │   ├── remove/
│                       │   │   │   │   └── page.tsx  # Remove payout method
│                       │   │   │   │       # BE: financial-be/payout_method
│                       │   │   │   │       # DELETE /v1/financial/payout-methods/{method_id}
│                       │   │   │   └── verify/
│                       │   │   │       └── page.tsx  # Verify payout method
│                       │   │   │           # BE: financial-be/payout_method
│                       │   │   │           # POST /v1/financial/payout-methods/{method_id}/verify
│                       │   │   ├── add/
│                       │   │   │   └── page.tsx  # Add payout method
│                       │   │   │       # - Bank account
│                       │   │   │       # - PayPal
│                       │   │   │       # - Wire transfer
│                       │   │   │       # BE: financial-be/payout_method
│                       │   │   │       # POST /v1/financial/payout-methods
│                       │   │   └── page.tsx  # Payout methods list
│                       │   │       # BE: financial-be/payout_method
│                       │   │       # GET /v1/financial/payout-methods
│                       │   ├── payment-methods/
│                       │   │   ├── [methodId]/
│                       │   │   │   ├── edit/
│                       │   │   │   │   └── page.tsx  # Edit payment method
│                       │   │   │   │       # BE: financial-be/payment_method
│                       │   │   │   │       # PUT /v1/financial/payment-methods/{method_id}
│                       │   │   │   ├── remove/
│                       │   │   │   │   └── page.tsx  # Remove payment method
│                       │   │   │   │       # BE: financial-be/payment_method
│                       │   │   │   │       # DELETE /v1/financial/payment-methods/{method_id}
│                       │   │   │   └── verify/
│                       │   │   │       └── page.tsx  # Verify payment method
│                       │   │   │           # - Micro-deposit verification
│                       │   │   │           # - Card verification
│                       │   │   │           # BE: financial-be/payment_method
│                       │   │   │           # POST /v1/financial/payment-methods/{method_id}/verify
│                       │   │   ├── add/
│                       │   │   │   └── page.tsx  # Add payment method
│                       │   │   │       # - Credit/debit card
│                       │   │   │       # - Bank account (ACH)
│                       │   │   │       # - Digital wallet
│                       │   │   │       # BE: financial-be/payment_method
│                       │   │   │       # POST /v1/financial/payment-methods
│                       │   │   └── page.tsx  # Payment methods list
│                       │   │       # BE: financial-be/payment_method
│                       │   │       # GET /v1/financial/payment-methods
│                       │   ├── reconciliation/
│                       │   │   └── page.tsx  # Financial reconciliation
│                       │   │       # - Match transactions
│                       │   │       # - Resolve discrepancies
│                       │   │       # - Bank statement upload
│                       │   │       # BE: financial-be/reconciliation
│                       │   │       # GET /v1/financial/reconciliation
│                       │   │       # POST /v1/financial/reconciliation/upload
│                       │   └── cost-centers/
│                       │       ├── [centerId]/
│                       │       │   ├── analytics/
│                       │       │   │   └── page.tsx  # Cost center analytics
│                       │       │   │       # BE: financial-be/cost_center, financial-be/analytics
│                       │       │   │       # GET /v1/financial/cost-centers/{center_id}/analytics
│                       │       │   ├── edit/
│                       │       │   │   └── page.tsx  # Edit cost center
│                       │       │   │       # BE: financial-be/cost_center
│                       │       │   │       # PUT /v1/financial/cost-centers/{center_id}
│                       │       │   └── page.tsx  # Cost center details
│                       │       │       # BE: financial-be/cost_center
│                       │       │       # GET /v1/financial/cost-centers/{center_id}
│                       │       ├── create/
│                       │       │   └── page.tsx  # Create cost center
│                       │       │       # BE: financial-be/cost_center
│                       │       │       # POST /v1/financial/cost-centers
│                       │       └── page.tsx  # Cost centers list
│                       │           # - Department/project budgets
│                       │           # - Spend tracking
│                       │           # BE: financial-be/cost_center
│                       │           # GET /v1/financial/cost-centers
│                       ├── jobs/
│                       │   ├── batch/
│                       │   │   ├── create/
│                       │   │   │   └── page.tsx  # Bulk job creation
│                       │   │   │       # - CSV upload
│                       │   │   │       # - Template selection
│                       │   │   │       # - Preview and publish
│                       │   │   │       # BE: jobs-be/job
│                       │   │   │       # POST /v1/jobs/batch
│                       │   │   └── manage/
│                       │   │       └── page.tsx  # Batch operations
│                       │   │           # - Bulk edit
│                       │   │           # - Bulk status change
│                       │   │           # - Bulk archive
│                       │   │           # BE: jobs-be/job
│                       │   │           # PATCH /v1/jobs/batch
│                       │   ├── insights/
│                       │   │   └── page.tsx  # Market insights for job posting
│                       │   │       # - Recommended rates
│                       │   │       # - Competition analysis
│                       │   │       # - Time-to-hire estimates
│                       │   │       # - Skill demand trends
│                       │   │       # BE: jobs-be/analytics, search-be/trending
│                       │   │       # GET /v1/jobs/market-insights
│                       │   │       # GET /v1/trending/skills
│                       │   └── scheduling/
│                       │       └── page.tsx  # Schedule job postings
│                       │           # - Future publish dates
│                       │           # - Auto-repost settings
│                       │           # - Expiration reminders
│                       │           # BE: jobs-be/job
│                       │           # POST /v1/jobs/{job_id}/schedule
│                       ├── proposals/
│                       │   ├── [proposalId]/
│                       │   │   ├── compare/
│                       │   │   │   └── page.tsx  # Compare with other proposals
│                       │   │   │       # - Side-by-side comparison
│                       │   │   │       # - Highlight differences
│                       │   │   │       # BE: proposals-be/proposal
│                       │   │   │       # GET /v1/proposals/compare?ids=...
│                       │   │   ├── feedback/
│                       │   │   │   └── page.tsx  # Proposal feedback
│                       │   │   │       # - Client feedback
│                       │   │   │       # - Improvement suggestions
│                       │   │   │       # BE: proposals-be/proposal
│                       │   │   │       # GET /v1/proposals/{proposal_id}/feedback
│                       │   │   └── negotiations/
│                       │   │       └── page.tsx  # Proposal negotiations history
│                       │   │           # - Negotiation timeline
│                       │   │           # - Counter-offers
│                       │   │           # - Terms evolution
│                       │   │           # BE: proposals-be/negotiation
│                       │   │           # GET /v1/proposals/{proposal_id}/negotiations
│                       │   ├── archived/
│                       │   │   └── page.tsx  # Archived proposals
│                       │   │       # - Old proposals
│                       │   │       # - Historical reference
│                       │   │       # BE: proposals-be/proposal
│                       │   │       # GET /v1/proposals?status=archived
│                       │   ├── benchmarking/
│                       │   │   └── page.tsx  # Benchmark against market
│                       │   │       # - Rate comparisons
│                       │   │       # - Proposal quality metrics
│                       │   │       # - Success factors
│                       │   │       # BE: proposals-be/analytics, search-be/trending
│                       │   │       # GET /v1/proposals/benchmarking
│                       │   └── insights/
│                       │       └── page.tsx  # Proposal insights
│                       │           # - Win rate analysis
│                       │           # - Optimal pricing insights
│                       │           # - Response time impact
│                       │           # - Proposal quality score
│                       │           # BE: proposals-be/analytics
│                       │           # GET /v1/proposals/insights
│                       ├── search/
│                       │   ├── freelancers/
│                       │   │   └── page.tsx  # Advanced freelancer search (client)
│                       │   │       # - Search by skills
│                       │   └── jobs/
│                       │       └── page.tsx  # Advanced job search
│                       │           # - Full-text search
│                       │           # - Faceted filters
│                       │           # - Autocomplete suggestions
│                       │           # - Search history
│                       │           # - Save search
│                       │           # BE: search-be/query
│                       │           # POST /v1/search/jobs
│                       │           # Body: { query, filters: {...}, sort, page }
│                       │           # BE: search-be/suggestions
│                       │           # GET /v1/suggestions?q={query}
│                       ├── settings/
│                       │   ├── advanced/
│                       │   │   └── page.tsx  # Advanced settings
│                       │   │       # - Debug mode
│                       │   │       # - Experimental APIs
│                       │   │       # - Performance monitoring opt-in
│                       │   │       # BE: users-be/preferences
│                       │   │       # GET /v1/settings/advanced
│                       │   │       # PUT /v1/settings/advanced
│                       │   ├── authorized-apps/
│                       │   │   ├── [appId]/
│                       │   │   │   └── revoke/
│                       │   │   │       └── page.tsx  # Revoke app access
│                       │   │   │           # BE: users-be/oauth_token
│                       │   │   │           # DELETE /v1/settings/authorized-apps/{app_id}
│                       │   │   └── page.tsx  # Authorized apps list
│                       │   │       # - Third-party app permissions
│                       │   │       # - Last used dates
│                       │   │       # BE: users-be/oauth_token
│                       │   │       # GET /v1/settings/authorized-apps
│                       │   ├── developer/
│                       │   │   ├── api-keys/
│                       │   │   │   ├── [keyId]/
│                       │   │   │   │   ├── regenerate/
│                       │   │   │   │   │   └── page.tsx  # Regenerate API key
│                       │   │   │   │   │       # BE: users-be/api_key
│                       │   │   │   │   │       # POST /v1/developer/api-keys/{key_id}/regenerate
│                       │   │   │   │   ├── revoke/
│                       │   │   │   │   │   └── page.tsx  # Revoke API key
│                       │   │   │   │   │       # BE: users-be/api_key
│                       │   │   │   │   │       # DELETE /v1/developer/api-keys/{key_id}
│                       │   │   │   │   └── page.tsx  # API key details
│                       │   │   │   │       # BE: users-be/api_key
│                       │   │   │   │       # GET /v1/developer/api-keys/{key_id}
│                       │   │   │   ├── create/
│                       │   │   │   │   └── page.tsx  # Create API key
│                       │   │   │   │       # - Name and permissions
│                       │   │   │   │       # - Expiration settings
│                       │   │   │   │       # BE: users-be/api_key
│                       │   │   │   │       # POST /v1/developer/api-keys
│                       │   │   │   └── page.tsx  # API keys list
│                       │   │   │       # BE: users-be/api_key
│                       │   │   │       # GET /v1/developer/api-keys
│                       │   │   ├── oauth-apps/
│                       │   │   │   ├── [appId]/
│                       │   │   │   │   ├── credentials/
│                       │   │   │   │   │   └── page.tsx  # OAuth credentials
│                       │   │   │   │   │       # - Client ID/Secret
│                       │   │   │   │   │       # - Regenerate secret
│                       │   │   │   │   │       # BE: users-be/oauth_app
│                       │   │   │   │   │       # POST /v1/developer/oauth-apps/{app_id}/regenerate-secret
│                       │   │   │   │   ├── edit/
│                       │   │   │   │   │   └── page.tsx  # Edit OAuth app
│                       │   │   │   │   │       # BE: users-be/oauth_app
│                       │   │   │   │   │       # PUT /v1/developer/oauth-apps/{app_id}
│                       │   │   │   │   └── page.tsx  # OAuth app details
│                       │   │   │   │       # BE: users-be/oauth_app
│                       │   │   │   │       # GET /v1/developer/oauth-apps/{app_id}
│                       │   │   │   ├── create/
│                       │   │   │   │   └── page.tsx  # Create OAuth app
│                       │   │   │   │       # BE: users-be/oauth_app
│                       │   │   │   │       # POST /v1/developer/oauth-apps
│                       │   │   │   └── page.tsx  # OAuth apps list
│                       │   │   │       # BE: users-be/oauth_app
│                       │   │   │       # GET /v1/developer/oauth-apps
│                       │   │   └── page.tsx  # Developer settings hub
│                       │   │       # BE: None (navigation)
│                       │   ├── labs/
│                       │   │   └── page.tsx  # Experimental features
│                       │   │       # - Beta feature toggles
│                       │   │       # - Early access programs
│                       │   │       # BE: utility-be/feature_flags
│                       │   │       # GET /v1/labs/features
│                       │   │       # PUT /v1/labs/features/{feature_id}/toggle
│                       │   └── authorized-apps/
│                       │       └── page.tsx  # Authorized apps list
│                       │           # - Third-party app permissions
│                       │           # - Last used dates
│                       │           # BE: users-be/oauth_token
│                       │           # GET /v1/settings/authorized-apps
│                       └── teams/
│                           ├── [teamId]/
│                           │   ├── compliance/
│                           │   │   └── page.tsx  # Team compliance dashboard
│                           │   │       # - Document status
│                           │   │       # - Training completion
│                           │   │       # - Policy acknowledgments
│                           │   │       # BE: users-be/team, admin-be/business_verification
│                           │   │       # GET /v1/teams/{team_id}/compliance
│                           │   ├── hierarchy/
│                           │   │   └── page.tsx  # Team hierarchy management
│                           │   │       # - Organizational chart
│                           │   │       # - Reporting structure
│                           │   │       # - Role relationships
│                           │   │       # BE: users-be/team
│                           │   │       # GET /v1/teams/{team_id}/hierarchy
│                           │   ├── performance/
│                           │   │   └── page.tsx  # Team performance metrics
│                           │   │       # - Productivity stats
│                           │   │       # - Quality metrics
│                           │   │       # - Member contributions
│                           │   │       # BE: users-be/team, contracts-be/analytics
│                           │   │       # GET /v1/teams/{team_id}/performance
│                           │   └── spending-controls/
│                           │       └── page.tsx  # Team spending controls
│                           │           # - Approval workflows
│                           │           # - Spending limits
│                           │           # - Auto-approval rules
│                           │           # BE: users-be/team, financial-be/budget
│                           │           # GET /v1/teams/{team_id}/spending-controls
│                           │           # PUT /v1/teams/{team_id}/spending-controls
│                           └── integrations/
│                               ├── [integrationId]/
│                               │   ├── configure/
│                               │   │   └── page.tsx  # Configure integration
│                               │   │       # BE: users-be/integration
│                               │   │       # PUT /v1/teams/integrations/{integration_id}
│                               │   ├── logs/
│                               │   │   └── page.tsx  # Integration logs
│                               │   │       # BE: users-be/integration, utility-be/audit
│                               │   │       # GET /v1/teams/integrations/{integration_id}/logs
│                               │   └── page.tsx  # Integration details
│                               │       # BE: users-be/integration
│                               │       # GET /v1/teams/integrations/{integration_id}
│                               ├── available/
│                               │   └── page.tsx  # Available integrations
│                               │       # - Slack, JIRA, etc.
│                               │       # - Feature descriptions
│                               │       # BE: users-be/integration
│                               │       # GET /v1/teams/integrations/available
│                               └── page.tsx  # Active integrations list
│                                   # BE: users-be/integration
│                                   # GET /v1/teams/integrations
├── config/
│   ├── .env.development
│   ├── .env.production
│   ├── .env.staging
│   └── environments/
│       ├── development.ts
│       ├── production.ts
│       └── staging.ts
├── docs/
│   ├── api/
│   │   ├── authentication.md  # API authentication
│   │   ├── error-handling.md  # Error handling
│   │   ├── introduction.md  # API integration guide
│   │   ├── rate-limiting.md  # Rate limiting
│   │   └── microservices/
│   │       ├── admin-be.md
│   │       ├── communications-be.md
│   │       ├── contracts-be.md
│   │       ├── financial-be.md
│   │       ├── jobs-be.md
│   │       ├── proposals-be.md
│   │       ├── reviews-be.md
│   │       ├── search-be.md
│   │       ├── storage-be.md
│   │       ├── subscriptions-be.md
│   │       └── users-be.md
│   ├── architecture/
│   │   ├── authentication.md  # Auth flow
│   │   ├── data-fetching.md  # Data fetching patterns
│   │   ├── frontend-architecture.md  # FE architecture details
│   │   ├── overview.md  # System architecture
│   │   ├── performance.md  # Performance optimization
│   │   ├── routing.md  # Routing strategy
│   │   └── state-management.md  # State management patterns
│   ├── components/
│   │   ├── accessibility.md
│   │   ├── component-library.md
│   │   ├── design-tokens.md
│   │   └── theming.md
│   └── guides/
│       ├── contributing.md
│       ├── deployment.md
│       ├── development.md
│       ├── getting-started.md
│       ├── testing.md
│       └── troubleshooting.md
├── packages/
│   ├── shared/
│   │   └── src/
│   │       ├── analytics/
│   │       │   ├── events/
│   │       │   │   ├── auth-events.ts  # Auth analytics
│   │       │   │   ├── contract-events.ts  # Contract analytics
│   │       │   │   ├── job-events.ts  # Job analytics
│   │       │   │   ├── payment-events.ts  # Payment analytics
│   │       │   │   ├── proposal-events.ts  # Proposal analytics
│   │       │   │   └── user-events.ts  # User behavior
│   │       │   ├── trackers/
│   │       │   │   ├── interaction-tracker.ts
│   │       │   │   ├── page-view-tracker.ts
│   │       │   │   └── conversion-tracker.ts
│   │       │   └── utils/
│   │       │       ├── anonymize.ts  # PII anonymization
│   │       │       └── event-builder.ts
│   │       ├── features/
│   │       │   ├── events/
│   │       │   │   ├── api/
│   │       │   │   │   └── events-api.ts  # Events API client
│   │       │   │   │   # BE: communications-be/events (if exists)
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── usePresence.ts  # User presence tracking
│   │       │   │   │   ├── useRealTimeUpdates.ts  # Real-time data sync
│   │       │   │   │   ├── useTypingIndicator.ts  # Typing indicators
│   │       │   │   │   └── useWebSocket.ts  # WebSocket connection
│   │       │   │   ├── providers/
│   │       │   │   │   └── WebSocketProvider.tsx  # WebSocket context
│   │       │   │   └── types.ts  # Event types
│   │       │   ├── experiments/
│   │       │   │   ├── api/
│   │       │   │   │   └── experiments-api.ts  # Experiments API
│   │       │   │   │   # BE: utility-be/experiments
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useExperiment.ts  # A/B test variant
│   │       │   │   │   ├── useExperimentTracking.ts  # Track experiment events
│   │       │   │   │   └── useFeatureVariant.ts  # Feature variant
│   │       │   │   ├── utils/
│   │       │   │   │   ├── experiment-context.ts  # Experiment context
│   │       │   │   │   └── variant-assignment.ts  # Variant logic
│   │       │   │   └── types.ts  # Experiment types
│   │       │   ├── geolocation/
│   │       │   │   ├── api/
│   │       │   │   │   └── geolocation-api.ts  # Geolocation API
│   │       │   │   │   # BE: utility-be/geolocation (if exists)
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useCountry.ts  # Country detection
│   │       │   │   │   ├── useGeolocation.ts  # Device location
│   │       │   │   │   └── useTimezone.ts  # User timezone
│   │       │   │   ├── utils/
│   │       │   │   │   ├── distance-calculator.ts  # Distance calculations
│   │       │   │   │   └── geocoding.ts  # Geocoding utilities
│   │       │   │   └── types.ts  # Geolocation types
│   │       │   ├── moderation/
│   │       │   │   ├── api/
│   │       │   │   │   └── moderation-api.ts  # Moderation API
│   │       │   │   │   # BE: admin-be/moderation
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useContentStatus.ts  # Content status check
│   │       │   │   │   ├── useModerationActions.ts  # Moderation actions
│   │       │   │   │   └── useReportContent.ts  # Report content
│   │       │   │   ├── utils/
│   │       │   │   │   ├── content-validator.ts  # Content validation
│   │       │   │   │   └── profanity-filter.ts  # Profanity filtering (client-side)
│   │       │   │   └── types.ts  # Moderation types
│   │       │   ├── offline/
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useNetworkStatus.ts  # Network status
│   │       │   │   │   ├── useOfflineQueue.ts  # Offline action queue
│   │       │   │   │   ├── useOfflineStorage.ts  # Local storage
│   │       │   │   │   └── useOfflineSync.ts  # Data synchronization
│   │       │   │   ├── store/
│   │       │   │   │   ├── offline-store.ts  # Offline state management
│   │       │   │   │   └── sync-store.ts  # Sync state management
│   │       │   │   └── types.ts  # Offline types
│   │       │   ├── gamification/
│   │       │   │   ├── api/
│   │       │   │   │   ├── achievements-api.ts  # Achievements API
│   │       │   │   │   │   # BE: users-be/achievement
│   │       │   │   │   ├── badges-api.ts  # Badges API
│   │       │   │   │   │   # BE: users-be/badge
│   │       │   │   │   └── leaderboards-api.ts  # Leaderboards API
│   │       │   │   │       # BE: users-be/leaderboard
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useAchievements.ts  # User achievements
│   │       │   │   │   ├── useBadges.ts  # User badges
│   │       │   │   │   ├── useLeaderboard.ts  # Leaderboard data
│   │       │   │   │   └── usePoints.ts  # Points/reputation
│   │       │   │   ├── queries/
│   │       │   │   │   ├── gamification-mutations.ts  # Gamification mutations
│   │       │   │   │   └── gamification-queries.ts  # Gamification queries
│   │       │   │   └── types.ts  # Gamification types
│   │       │   └── referrals/
│   │       │       ├── api/
│   │       │       │   └── referrals-api.ts  # Referrals API (enhanced)
│   │       │       │       # BE: users-be/referral
│   │       │       ├── hooks/
│   │       │       │   ├── useReferralCode.ts  # User referral code
│   │       │       │   ├── useReferralProgram.ts  # Referral program details
│   │       │       │   ├── useReferralStats.ts  # Referral statistics
│   │       │       │   └── useRewards.ts  # Rewards management
│   │       │       ├── queries/
│   │       │       │   ├── referrals-mutations.ts  # Referral mutations
│   │       │       │   └── referrals-queries.ts  # Referral queries
│   │       │       └── types.ts  # Referral types
│   │       ├── monitoring/
│   │       │   ├── analytics.ts  # Analytics setup
│   │       │   ├── logger.ts  # Logging configuration
│   │       │   ├── sentry.ts  # Error tracking setup
│   │       │   └── web-vitals.ts  # Performance monitoring
│   │       ├── performance/
│   │       │   ├── hooks/
│   │       │   │   ├── useAnalyticsTracking.ts  # Analytics events
│   │       │   │   ├── useErrorTracking.ts  # Error monitoring
│   │       │   │   └── usePerformanceMetrics.ts  # Web vitals tracking
│   │       │   ├── utils/
│   │       │   │   ├── error-reporter.ts  # Error reporting
│   │       │   │   ├── metrics-collector.ts  # Metrics collection
│   │       │   │   └── trace-headers.ts  # Distributed tracing
│   │       │   └── types.ts  # Performance types
│   │       ├── security/
│   │       │   ├── csrf.ts  # CSRF protection
│   │       │   ├── encryption.ts  # Client-side encryption
│   │       │   ├── permissions.ts  # Permission checks
│   │       │   ├── sanitization.ts  # Input sanitization
│   │       │   └── validation.ts  # Input validation
│   │       └── testing/
│   │           ├── mock-api/
│   │           │   ├── handlers.ts  # MSW handlers
│   │           │   └── server.ts  # MSW server setup
│   │           ├── mock-data/
│   │           │   ├── contracts.ts  # Mock contract data
│   │           │   ├── jobs.ts  # Mock job data
│   │           │   ├── messages.ts  # Mock message data
│   │           │   ├── proposals.ts  # Mock proposal data
│   │           │   └── users.ts  # Mock user data
│   │           ├── providers/
│   │           │   └── TestProviders.tsx  # Test provider wrapper
│   │           └── test-utils.ts  # Common test utilities
│   └── ui/
│       └── src/
│           ├── accessibility/
│           │   ├── FocusTrap/
│           │   │   ├── FocusTrap.native.tsx
│           │   │   ├── FocusTrap.tsx
│           │   │   └── FocusTrap.web.tsx
│           │   ├── ScreenReaderAnnouncer/
│           │   │   ├── ScreenReaderAnnouncer.native.tsx
│           │   │   ├── ScreenReaderAnnouncer.tsx
│           │   │   └── ScreenReaderAnnouncer.web.tsx
│           │   └── SkipLinks/
│           │       ├── SkipLinks.tsx
│           │       └── SkipLinks.web.tsx
│           ├── ai/
│           │   ├── AIAssistant/
│           │   │   ├── AIAssistant.native.tsx
│           │   │   ├── AIAssistant.tsx
│           │   │   └── AIAssistant.web.tsx
│           │   ├── AutoComplete/
│           │   │   ├── AIAutoComplete.native.tsx
│           │   │   ├── AIAutoComplete.tsx
│           │   │   └── AIAutoComplete.web.tsx
│           │   └── SmartSuggestions/
│           │       ├── SmartSuggestions.native.tsx
│           │       ├── SmartSuggestions.tsx
│           │       └── SmartSuggestions.web.tsx
│           ├── forms/
│           │   ├── CodeEditor/
│           │   │   ├── CodeEditor.native.tsx
│           │   │   ├── CodeEditor.tsx
│           │   │   └── CodeEditor.web.tsx
│           │   ├── DateRangePicker/
│           │   │   ├── DateRangePicker.native.tsx
│           │   │   ├── DateRangePicker.tsx
│           │   │   └── DateRangePicker.web.tsx
│           │   ├── MarkdownEditor/
│           │   │   ├── MarkdownEditor.native.tsx
│           │   │   ├── MarkdownEditor.tsx
│           │   │   └── MarkdownEditor.web.tsx
│           │   ├── RichTextEditor/
│           │   │   ├── RichTextEditor.native.tsx
│           │   │   ├── RichTextEditor.tsx
│           │   │   └── RichTextEditor.web.tsx
│           │   └── SignatureInput/
│           │       ├── SignatureInput.native.tsx
│           │       ├── SignatureInput.tsx
│           │       └── SignatureInput.web.tsx
│           └── visualization/
│               ├── Gantt/
│               │   ├── GanttChart.native.tsx
│               │   ├── GanttChart.tsx
│               │   └── GanttChart.web.tsx
│               ├── Heatmap/
│               │   ├── Heatmap.native.tsx
│               │   ├── Heatmap.tsx
│               │   └── Heatmap.web.tsx
│               ├── Kanban/
│               │   ├── KanbanBoard.native.tsx
│               │   ├── KanbanBoard.tsx
│               │   └── KanbanBoard.web.tsx
│               └── OrgChart/
│                   ├── OrganizationChart.native.tsx
│                   ├── OrganizationChart.tsx
│                   └── OrganizationChart.web.tsx
└── .env.example  # Environment variables template
