# Skillsier Frontend Folder Structure
## Complete Monorepo Architecture with Backend Mappings

> **CRITICAL**: This document contains ONLY the folder structure, filenames, and inline backend API mappings.  
> **NO CODE IMPLEMENTATIONS** are included per the strict output policy.

---

## Table of Contents
1. [Root Structure](#root-structure)
2. [Apps - Web Application](#apps---web-application)
3. [Apps - Mobile Application](#apps---mobile-application)
4. [Packages - Shared Libraries](#packages---shared-libraries)
5. [Infrastructure & Configuration](#infrastructure--configuration)

---

## Root Structure

```
skillsier-fe/
├── .husky/                                   # Git hooks for quality gates
│   ├── pre-commit                           # Runs linting, type checking
│   └── pre-push                             # Runs tests before push
│
├── .vscode/                                  # VS Code workspace configuration
│   ├── extensions.json                      # Recommended extensions
│   ├── launch.json                          # Debug configurations
│   └── settings.json                        # Workspace settings (format on save, etc.)
│
├── apps/                                     # Application workspaces
│   ├── web/                                 # Next.js web app (see detailed structure below)
│   └── mobile/                              # React Native/Expo app (see detailed structure below)
│
├── packages/                                 # Shared libraries
│   ├── ui/                                  # Cross-platform component library
│   ├── shared/                              # Shared business logic, hooks, utilities
│   ├── types/                               # TypeScript type definitions
│   └── config/                              # Shared configurations (ESLint, TS, Tailwind)
│
├── deploy/                                   # Deployment configurations
│   ├── k8s/                                 # Kubernetes manifests
│   │   ├── web-deployment.yaml
│   │   ├── web-service.yaml
│   │   ├── web-ingress.yaml
│   │   └── configmap.yaml
│   └── docker/
│       ├── web.Dockerfile
│       └── mobile-build.Dockerfile
│
├── docs/                                     # Documentation
│   ├── ARCHITECTURE.md
│   ├── SETUP.md
│   ├── PERFORMANCE.md
│   └── CONTRIBUTING.md
│
├── .env.example                              # Environment variables template
├── .gitignore
├── .nvmrc                                    # Node version specification
├── package.json                              # Root package (workspace manager)
├── pnpm-workspace.yaml                       # pnpm workspace configuration
├── pnpm-lock.yaml                            # Locked dependencies
├── turbo.json                                # Turborepo pipeline configuration
├── tsconfig.base.json                        # Base TypeScript configuration
└── README.md
```

---

## Apps - Web Application

### Main Structure

```
apps/web/
├── public/                                   # Static assets
│   ├── images/
│   │   ├── logo.svg
│   │   ├── hero-bg.webp
│   │   └── placeholder-avatar.png
│   ├── fonts/
│   │   ├── inter-var.woff2
│   │   └── arabic-font.woff2
│   ├── locales/                             # Public locale files (for static routing)
│   │   ├── en/
│   │   ├── ar/
│   │   └── ...
│   ├── robots.txt
│   ├── sitemap.xml
│   └── manifest.json                        # PWA manifest
│
├── src/
│   ├── app/                                 # Next.js 15 App Router (see detailed structure below)
│   ├── components/                          # Shared web-specific components
│   ├── features/                            # Feature modules (see detailed structure below)
│   ├── lib/                                 # Web-specific utilities
│   ├── hooks/                               # Web-specific hooks
│   ├── styles/                              # Global styles
│   └── middleware.ts                        # Next.js middleware (auth, i18n routing)
│                                            # BE: Keycloak token validation via JWT
│
├── .env.local
├── .eslintrc.json
├── next.config.js                           # Next.js configuration (i18n, image optimization, etc.)
├── postcss.config.js
├── tailwind.config.js
├── tsconfig.json
└── package.json
```

### App Router Structure (`apps/web/src/app`)

```
src/app/
├── [locale]/                                # Internationalized routing
│   │
│   ├── (auth)/                             # Auth route group (no dashboard layout)
│   │   ├── login/
│   │   │   └── page.tsx                    # Login page
│   │   │                                   # BE: Keycloak OAuth2 flow → users-be/auth
│   │   │                                   # POST /v1/auth/login → JWT tokens
│   │   │
│   │   ├── register/
│   │   │   ├── page.tsx                    # Registration page
│   │   │   │                               # BE: users-be/user — POST /v1/users/register
│   │   │   │                               # Keycloak user creation + profile initialization
│   │   │   └── verification/
│   │   │       └── page.tsx                # Email verification callback
│   │   │                                   # BE: users-be/user — POST /v1/users/verify-email
│   │   │
│   │   ├── forgot-password/
│   │   │   └── page.tsx                    # Password reset request
│   │   │                                   # BE: users-be/security/recovery — POST /v1/auth/forgot-password
│   │   │
│   │   ├── reset-password/
│   │   │   └── page.tsx                    # Password reset form (token-based)
│   │   │                                   # BE: users-be/security/recovery — POST /v1/auth/reset-password
│   │   │
│   │   ├── callback/
│   │   │   └── page.tsx                    # OAuth callback handler (Keycloak, Google, GitHub)
│   │   │                                   # BE: Keycloak token exchange
│   │   │
│   │   └── layout.tsx                      # Auth pages layout (no header/footer)
│   │
│   ├── (landing)/                          # Public landing pages (no auth required)
│   │   ├── page.tsx                        # Homepage
│   │   ├── about/
│   │   │   └── page.tsx                    # About page
│   │   ├── how-it-works/
│   │   │   ├── freelancers/
│   │   │   │   └── page.tsx
│   │   │   └── clients/
│   │   │       └── page.tsx
│   │   ├── pricing/
│   │   │   └── page.tsx                    # Pricing page
│   │   │                                   # BE: subscriptions-be/plans — GET /v1/plans?public=true
│   │   │
│   │   ├── trust-safety/
│   │   │   └── page.tsx
│   │   ├── contact/
│   │   │   └── page.tsx                    # Contact form
│   │   │                                   # BE: communications-be/messages — POST /v1/contact
│   │   │
│   │   └── layout.tsx                      # Landing layout (public header/footer)
│   │
│   ├── (dashboard)/                        # Authenticated dashboard (all roles)
│   │   │
│   │   ├── dashboard/                      # Main dashboard (role-based view)
│   │   │   └── page.tsx                    # Dashboard home (redirects based on role)
│   │   │                                   # BE: users-be/user — GET /v1/users/me
│   │   │                                   # BE: analytics-be — GET /v1/analytics/dashboard
│   │   │
│   │   ├── profile/                        # User profile management
│   │   │   ├── page.tsx                    # Profile overview
│   │   │   │                               # BE: users-be/profile — GET /v1/users/{id}/profile
│   │   │   │
│   │   │   ├── edit/
│   │   │   │   └── page.tsx                # Edit profile form
│   │   │   │                               # BE: users-be/profile — PATCH /v1/users/{id}/profile
│   │   │   │
│   │   │   ├── skills/
│   │   │   │   └── page.tsx                # Skills management
│   │   │   │                               # BE: users-be/capabilities — GET /v1/users/{id}/skills
│   │   │   │                               # BE: users-be/capabilities — POST /v1/users/{id}/skills
│   │   │   │                               # BE: users-be/capabilities — PUT /v1/users/{id}/skills/{skill_id}
│   │   │   │                               # BE: users-be/capabilities — DELETE /v1/users/{id}/skills/{skill_id}
│   │   │   │
│   │   │   ├── experience/
│   │   │   │   └── page.tsx                # Work experience management
│   │   │   │                               # BE: users-be/experience — GET /v1/users/{id}/experience
│   │   │   │                               # BE: users-be/experience — POST /v1/users/{id}/experience
│   │   │   │                               # BE: users-be/experience — PUT /v1/users/{id}/experience/{exp_id}
│   │   │   │                               # BE: users-be/experience — DELETE /v1/users/{id}/experience/{exp_id}
│   │   │   │
│   │   │   ├── education/
│   │   │   │   └── page.tsx                # Education management
│   │   │   │                               # BE: users-be/education — GET /v1/users/{id}/education
│   │   │   │                               # BE: users-be/education — POST /v1/users/{id}/education
│   │   │   │
│   │   │   ├── portfolio/
│   │   │   │   ├── page.tsx                # Portfolio items list
│   │   │   │   │                           # BE: users-be/portfolio — GET /v1/users/{id}/portfolio
│   │   │   │   │
│   │   │   │   ├── new/
│   │   │   │   │   └── page.tsx            # Create portfolio item
│   │   │   │   │                           # BE: users-be/portfolio — POST /v1/users/{id}/portfolio
│   │   │   │   │                           # BE: storage-be/uploads — POST /v1/storage/upload (for media)
│   │   │   │   │
│   │   │   │   └── [itemId]/
│   │   │   │       ├── page.tsx            # Portfolio item details
│   │   │   │       │                       # BE: users-be/portfolio — GET /v1/portfolio/{item_id}
│   │   │   │       │
│   │   │   │       └── edit/
│   │   │   │           └── page.tsx        # Edit portfolio item
│   │   │   │                               # BE: users-be/portfolio — PUT /v1/portfolio/{item_id}
│   │   │   │
│   │   │   ├── services/
│   │   │   │   └── page.tsx                # Service catalog management
│   │   │   │                               # BE: users-be/service_catalog — GET /v1/users/{id}/services
│   │   │   │                               # BE: users-be/service_catalog — POST /v1/users/{id}/services
│   │   │   │
│   │   │   ├── availability/
│   │   │   │   └── page.tsx                # Availability & calendar management
│   │   │   │                               # BE: users-be/availability — GET /v1/users/{id}/availability
│   │   │   │                               # BE: users-be/availability — PUT /v1/users/{id}/availability
│   │   │   │
│   │   │   └── settings/
│   │   │       ├── page.tsx                # Profile settings
│   │   │       │                           # BE: users-be/profile — GET /v1/users/{id}/settings
│   │   │       │
│   │   │       ├── privacy/
│   │   │       │   └── page.tsx            # Privacy settings
│   │   │       │                           # BE: users-be/privacy — GET /v1/users/{id}/privacy
│   │   │       │                           # BE: users-be/privacy — PATCH /v1/users/{id}/privacy
│   │   │       │
│   │   │       ├── visibility/
│   │   │       │   └── page.tsx            # Profile visibility settings
│   │   │       │                           # BE: users-be/profile_visibility — GET /v1/users/{id}/visibility
│   │   │       │                           # BE: users-be/profile_visibility — PATCH /v1/users/{id}/visibility
│   │   │       │
│   │   │       └── notifications/
│   │   │           └── page.tsx            # Notification preferences
│   │   │                                   # BE: users-be/email_preferences — GET /v1/users/{id}/notifications
│   │   │                                   # BE: users-be/email_preferences — PATCH /v1/users/{id}/notifications
│   │   │
│   │   ├── jobs/                           # Job browsing & management
│   │   │   ├── page.tsx                    # Job search/browse page
│   │   │   │                               # BE: search-be/search — POST /v1/search/jobs
│   │   │   │                               # Query params: skills[], location, budget_min, budget_max, job_type
│   │   │   │                               # Pagination: cursor-based, staleTime=30s
│   │   │   │
│   │   │   ├── saved/
│   │   │   │   └── page.tsx                # Saved jobs
│   │   │   │                               # BE: search-be/saved_searches — GET /v1/search/saved
│   │   │   │
│   │   │   ├── recommended/
│   │   │   │   └── page.tsx                # Recommended jobs (ML-powered)
│   │   │   │                               # BE: search-be/recommendations — GET /v1/recommendations/jobs
│   │   │   │
│   │   │   ├── [jobId]/
│   │   │   │   ├── page.tsx                # Job details page
│   │   │   │   │                           # BE: jobs-be/job — GET /v1/jobs/{job_id}
│   │   │   │   │                           # BE: jobs-be/attachments — GET /v1/jobs/{job_id}/attachments
│   │   │   │   │                           # BE: jobs-be/screening — GET /v1/jobs/{job_id}/screening
│   │   │   │   │
│   │   │   │   └── apply/
│   │   │   │       └── page.tsx            # Job application form (submit proposal)
│   │   │   │                               # BE: proposals-be/proposal — POST /v1/proposals
│   │   │   │                               # Includes: cover_letter, proposed_rate, milestones[], question_answers[]
│   │   │   │                               # Invalidates: ['proposals', 'list']
│   │   │   │
│   │   │   └── new/                        # Post a job (client only)
│   │   │       ├── page.tsx                # Job posting wizard step 1
│   │   │       │                           # BE: jobs-be/job — POST /v1/jobs (creates draft)
│   │   │       │
│   │   │       ├── details/
│   │   │       │   └── page.tsx            # Job details step
│   │   │       │                           # BE: jobs-be/job — PATCH /v1/jobs/{job_id}
│   │   │       │
│   │   │       ├── budget/
│   │   │       │   └── page.tsx            # Budget & payment setup
│   │   │       │                           # BE: jobs-be/pricing — PATCH /v1/jobs/{job_id}/pricing
│   │   │       │
│   │   │       ├── requirements/
│   │   │       │   └── page.tsx            # Requirements & screening questions
│   │   │       │                           # BE: jobs-be/requirements — POST /v1/jobs/{job_id}/requirements
│   │   │       │                           # BE: jobs-be/screening — POST /v1/jobs/{job_id}/screening/questions
│   │   │       │
│   │   │       ├── review/
│   │   │       │   └── page.tsx            # Review & publish
│   │   │       │                           # BE: jobs-be/job — POST /v1/jobs/{job_id}/publish
│   │   │       │                           # Triggers: job.published.v1 event → search-be indexing
│   │   │       │
│   │   │       └── success/
│   │   │           └── page.tsx            # Job posted successfully
│   │   │
│   │   ├── proposals/                      # Proposal management
│   │   │   ├── page.tsx                    # All proposals (list/table view)
│   │   │   │                               # BE: proposals-be/proposal — GET /v1/proposals?freelancer_id={user_id}
│   │   │   │                               # Filters: status[], job_type[], date_range
│   │   │   │                               # Pagination: cursor-based, keepPreviousData=true
│   │   │   │
│   │   │   ├── [proposalId]/
│   │   │   │   ├── page.tsx                # Proposal details
│   │   │   │   │                           # BE: proposals-be/proposal — GET /v1/proposals/{proposal_id}
│   │   │   │   │                           # BE: proposals-be/conversation — GET /v1/proposals/{proposal_id}/conversation
│   │   │   │   │
│   │   │   │   ├── edit/
│   │   │   │   │   └── page.tsx            # Edit proposal (if still in DRAFT/PENDING)
│   │   │   │   │                           # BE: proposals-be/proposal — PUT /v1/proposals/{proposal_id}
│   │   │   │   │
│   │   │   │   └── withdraw/
│   │   │   │       └── page.tsx            # Withdraw proposal confirmation
│   │   │   │                               # BE: proposals-be/proposal — POST /v1/proposals/{proposal_id}/withdraw
│   │   │   │                               # Invalidates: ['proposals', 'detail', proposal_id]
│   │   │   │
│   │   │   ├── received/                   # Client view: proposals received on jobs
│   │   │   │   └── page.tsx                # All proposals received
│   │   │   │                               # BE: proposals-be/proposal — GET /v1/proposals?client_id={user_id}
│   │   │   │
│   │   │   └── templates/
│   │   │       └── page.tsx                # Proposal templates (reusable)
│   │   │                                   # BE: proposals-be/template — GET /v1/proposals/templates
│   │   │                                   # BE: proposals-be/template — POST /v1/proposals/templates
│   │   │
│   │   ├── contracts/                      # Contract management
│   │   │   ├── page.tsx                    # All contracts (list view)
│   │   │   │                               # BE: contracts-be/contract — GET /v1/contracts?user_id={user_id}
│   │   │   │                               # Filters: status[], contract_type[], role (freelancer/client)
│   │   │   │
│   │   │   ├── active/
│   │   │   │   └── page.tsx                # Active contracts only
│   │   │   │                               # BE: contracts-be/contract — GET /v1/contracts?status=ACTIVE
│   │   │   │
│   │   │   ├── [contractId]/
│   │   │   │   ├── page.tsx                # Contract overview
│   │   │   │   │                           # BE: contracts-be/contract — GET /v1/contracts/{contract_id}
│   │   │   │   │                           # BE: contracts-be/milestones — GET /v1/contracts/{contract_id}/milestones
│   │   │   │   │                           # BE: financial-be/escrow — GET /v1/escrow?contract_id={contract_id}
│   │   │   │   │
│   │   │   │   ├── milestones/
│   │   │   │   │   ├── page.tsx            # Milestones list
│   │   │   │   │   │                       # BE: contracts-be/milestones — GET /v1/contracts/{contract_id}/milestones
│   │   │   │   │   │
│   │   │   │   │   ├── [milestoneId]/
│   │   │   │   │   │   ├── page.tsx        # Milestone details
│   │   │   │   │   │   │                   # BE: contracts-be/milestones — GET /v1/milestones/{milestone_id}
│   │   │   │   │   │   │                   # BE: contracts-be/deliverables — GET /v1/milestones/{milestone_id}/deliverables
│   │   │   │   │   │   │
│   │   │   │   │   │   ├── submit/
│   │   │   │   │   │   │   └── page.tsx    # Submit milestone (freelancer)
│   │   │   │   │   │   │                   # BE: contracts-be/milestones — POST /v1/milestones/{milestone_id}/submit
│   │   │   │   │   │   │                   # BE: storage-be/uploads — POST /v1/storage/upload (deliverables)
│   │   │   │   │   │   │                   # Triggers: milestone.submitted.v1 → communications-be notification
│   │   │   │   │   │   │
│   │   │   │   │   │   └── review/
│   │   │   │   │   │       └── page.tsx    # Review & approve/reject (client)
│   │   │   │   │   │                       # BE: contracts-be/milestones — POST /v1/milestones/{milestone_id}/approve
│   │   │   │   │   │                       # BE: contracts-be/milestones — POST /v1/milestones/{milestone_id}/reject
│   │   │   │   │   │                       # Triggers: milestone.approved.v1 → financial-be escrow release
│   │   │   │   │   │
│   │   │   │   │   └── new/
│   │   │   │   │       └── page.tsx        # Create milestone (client)
│   │   │   │   │                           # BE: contracts-be/milestones — POST /v1/milestones
│   │   │   │   │
│   │   │   │   ├── timesheet/              # Hourly contracts only
│   │   │   │   │   └── page.tsx            # Timesheet management
│   │   │   │   │                           # BE: contracts-be/timesheet — GET /v1/contracts/{contract_id}/timesheets
│   │   │   │   │                           # BE: contracts-be/timesheet — POST /v1/timesheets (log hours)
│   │   │   │   │
│   │   │   │   ├── work-diary/             # Work diary (screenshot tracking)
│   │   │   │   │   └── page.tsx            # Work diary view
│   │   │   │   │                           # BE: contracts-be/workdiary — GET /v1/contracts/{contract_id}/work-diary
│   │   │   │   │
│   │   │   │   ├── messages/
│   │   │   │   │   └── page.tsx            # Contract-specific messaging
│   │   │   │   │                           # BE: communications-be/conversations — GET /v1/conversations?context=contract&context_id={contract_id}
│   │   │   │   │                           # BE: communications-be/messages — POST /v1/messages
│   │   │   │   │                           # WebSocket: wss://api/v1/ws/conversations/{conversation_id}
│   │   │   │   │
│   │   │   │   ├── disputes/
│   │   │   │   │   ├── page.tsx            # Disputes list
│   │   │   │   │   │                       # BE: contracts-be/dispute — GET /v1/contracts/{contract_id}/disputes
│   │   │   │   │   │
│   │   │   │   │   ├── new/
│   │   │   │   │   │   └── page.tsx        # Create dispute
│   │   │   │   │   │                       # BE: contracts-be/dispute — POST /v1/disputes
│   │   │   │   │   │                       # BE: storage-be/uploads — POST /v1/storage/upload (evidence)
│   │   │   │   │   │
│   │   │   │   │   └── [disputeId]/
│   │   │   │   │       └── page.tsx        # Dispute details & resolution
│   │   │   │   │                           # BE: contracts-be/dispute — GET /v1/disputes/{dispute_id}
│   │   │   │   │                           # BE: admin-be/case_mgmt — GET /v1/admin/disputes/{dispute_id} (admin view)
│   │   │   │   │
│   │   │   │   ├── amendments/
│   │   │   │   │   └── page.tsx            # Contract amendments (SOW changes)
│   │   │   │   │                           # BE: contracts-be/sow — GET /v1/contracts/{contract_id}/amendments
│   │   │   │   │                           # BE: contracts-be/sow — POST /v1/contracts/{contract_id}/amendments
│   │   │   │   │
│   │   │   │   └── end/
│   │   │   │       └── page.tsx            # End contract flow
│   │   │   │                               # BE: contracts-be/contract — POST /v1/contracts/{contract_id}/end
│   │   │   │                               # BE: reviews-be/review — POST /v1/reviews (after contract end)
│   │   │   │
│   │   │   └── templates/
│   │   │       └── page.tsx                # Contract templates (SOW templates)
│   │   │                                   # BE: contracts-be/sow — GET /v1/contracts/templates
│   │   │
│   │   ├── payments/                       # Financial management
│   │   │   ├── page.tsx                    # Payment dashboard
│   │   │   │                               # BE: financial-be/wallet — GET /v1/wallets/{user_id}
│   │   │   │                               # BE: financial-be/transactions — GET /v1/transactions?user_id={user_id}
│   │   │   │
│   │   │   ├── wallet/
│   │   │   │   ├── page.tsx                # Wallet overview
│   │   │   │   │                           # BE: financial-be/wallet — GET /v1/wallets/{user_id}
│   │   │   │   │                           # Shows: available_balance, pending_balance, reserved_balance
│   │   │   │   │
│   │   │   │   ├── deposit/
│   │   │   │   │   └── page.tsx            # Deposit funds
│   │   │   │   │                           # BE: financial-be/wallet — POST /v1/wallets/{wallet_id}/deposit
│   │   │   │   │                           # BE: financial-be/payment — POST /v1/payments (process payment)
│   │   │   │   │
│   │   │   │   └── withdraw/
│   │   │   │       └── page.tsx            # Withdraw funds
│   │   │   │                               # BE: financial-be/payout — POST /v1/payouts/request
│   │   │   │                               # BE: financial-be/payment_method — GET /v1/payment-methods (verify payout method)
│   │   │   │
│   │   │   ├── methods/
│   │   │   │   ├── page.tsx                # Payment methods list
│   │   │   │   │                           # BE: financial-be/payment_method — GET /v1/payment-methods?user_id={user_id}
│   │   │   │   │
│   │   │   │   ├── add/
│   │   │   │   │   └── page.tsx            # Add payment method
│   │   │   │   │                           # BE: financial-be/payment_method — POST /v1/payment-methods
│   │   │   │   │                           # Gateway: Stripe/PayPal tokenization
│   │   │   │   │
│   │   │   │   └── [methodId]/
│   │   │   │       ├── page.tsx            # Payment method details
│   │   │   │       └── verify/
│   │   │   │           └── page.tsx        # Verify payment method
│   │   │   │                               # BE: financial-be/payment_method — POST /v1/payment-methods/{method_id}/verify
│   │   │   │
│   │   │   ├── transactions/
│   │   │   │   └── page.tsx                # Transaction history
│   │   │   │                               # BE: financial-be/transactions — GET /v1/transactions?user_id={user_id}
│   │   │   │                               # Filters: type[], status[], date_range
│   │   │   │                               # Pagination: cursor-based
│   │   │   │
│   │   │   ├── invoices/
│   │   │   │   ├── page.tsx                # Invoices list
│   │   │   │   │                           # BE: financial-be/invoice — GET /v1/invoices?user_id={user_id}
│   │   │   │   │
│   │   │   │   └── [invoiceId]/
│   │   │   │       └── page.tsx            # Invoice details & download
│   │   │   │                               # BE: financial-be/invoice — GET /v1/invoices/{invoice_id}
│   │   │   │                               # BE: financial-be/invoice — GET /v1/invoices/{invoice_id}/pdf (download)
│   │   │   │
│   │   │   ├── refunds/
│   │   │   │   └── page.tsx                # Refunds history
│   │   │   │                               # BE: financial-be/refund — GET /v1/refunds?user_id={user_id}
│   │   │   │
│   │   │   └── tax/
│   │   │       ├── page.tsx                # Tax documents & settings
│   │   │       │                           # BE: financial-be/tax — GET /v1/tax/profile?user_id={user_id}
│   │   │       │                           # BE: financial-be/tax — GET /v1/tax/documents?user_id={user_id}
│   │   │       │
│   │   │       └── setup/
│   │   │           └── page.tsx            # Tax profile setup (W9, VAT, etc.)
│   │   │                                   # BE: financial-be/tax — POST /v1/tax/profile
│   │   │
│   │   ├── messages/                       # Messaging center
│   │   │   ├── page.tsx                    # All conversations
│   │   │   │                               # BE: communications-be/conversations — GET /v1/conversations?user_id={user_id}
│   │   │   │                               # WebSocket: wss://api/v1/ws/user/{user_id} (real-time updates)
│   │   │   │
│   │   │   ├── [conversationId]/
│   │   │   │   └── page.tsx                # Conversation thread
│   │   │   │                               # BE: communications-be/messages — GET /v1/conversations/{conv_id}/messages
│   │   │   │                               # BE: communications-be/messages — POST /v1/messages
│   │   │   │                               # BE: communications-be/messages — POST /v1/messages/{msg_id}/read
│   │   │   │                               # WebSocket: wss://api/v1/ws/conversations/{conv_id}
│   │   │   │                               # Real-time: typing indicators, read receipts, new messages
│   │   │   │
│   │   │   └── compose/
│   │   │       └── page.tsx                # Start new conversation
│   │   │                                   # BE: communications-be/conversations — POST /v1/conversations
│   │   │
│   │   ├── notifications/
│   │   │   └── page.tsx                    # Notifications center
│   │   │                                   # BE: communications-be/notifications — GET /v1/notifications?user_id={user_id}
│   │   │                                   # BE: communications-be/notifications — POST /v1/notifications/{notif_id}/read
│   │   │                                   # BE: communications-be/notifications — POST /v1/notifications/read-all
│   │   │                                   # WebSocket: wss://api/v1/ws/notifications/{user_id}
│   │   │
│   │   ├── subscriptions/                  # Subscription management
│   │   │   ├── page.tsx                    # Current subscription
│   │   │   │                               # BE: subscriptions-be/subscriptions — GET /v1/subscriptions?user_id={user_id}
│   │   │   │                               # BE: subscriptions-be/entitlements — GET /v1/entitlements?user_id={user_id}
│   │   │   │
│   │   │   ├── plans/
│   │   │   │   └── page.tsx                # Available plans
│   │   │   │                               # BE: subscriptions-be/plans — GET /v1/plans
│   │   │   │
│   │   │   ├── upgrade/
│   │   │   │   └── page.tsx                # Upgrade subscription
│   │   │   │                               # BE: subscriptions-be/subscriptions — POST /v1/subscriptions/{sub_id}/upgrade
│   │   │   │                               # BE: financial-be/payment — POST /v1/payments (process upgrade payment)
│   │   │   │
│   │   │   ├── addons/
│   │   │   │   └── page.tsx                # Manage addons
│   │   │   │                               # BE: subscriptions-be/addons — GET /v1/addons
│   │   │   │                               # BE: subscriptions-be/subscriptions — POST /v1/subscriptions/{sub_id}/addons
│   │   │   │
│   │   │   └── connects/                   # Connects management (proposal credits)
│   │   │       ├── page.tsx                # Connects balance
│   │   │       │                           # BE: subscriptions-be/connects — GET /v1/connects/balance?user_id={user_id}
│   │   │       │
│   │   │       └── purchase/
│   │   │           └── page.tsx            # Purchase connects
│   │   │                                   # BE: subscriptions-be/connects — POST /v1/connects/purchase
│   │   │                                   # BE: financial-be/payment — POST /v1/payments
│   │   │
│   │   ├── reviews/                        # Reviews management
│   │   │   ├── page.tsx                    # All reviews (given & received)
│   │   │   │                               # BE: reviews-be/reviews — GET /v1/reviews?user_id={user_id}
│   │   │   │
│   │   │   ├── received/
│   │   │   │   └── page.tsx                # Reviews received
│   │   │   │                               # BE: reviews-be/reviews — GET /v1/reviews?reviewee_id={user_id}
│   │   │   │
│   │   │   ├── given/
│   │   │   │   └── page.tsx                # Reviews given
│   │   │   │                               # BE: reviews-be/reviews — GET /v1/reviews?reviewer_id={user_id}
│   │   │   │
│   │   │   └── write/
│   │   │       └── page.tsx                # Write review (after contract completion)
│   │   │                                   # BE: reviews-be/reviews — POST /v1/reviews
│   │   │                                   # Triggers: review.submitted.v1 → reputation update
│   │   │
│   │   ├── analytics/                      # Analytics & insights (client only)
│   │   │   ├── page.tsx                    # Analytics dashboard
│   │   │   │                               # BE: analytics-be — GET /v1/analytics/overview?user_id={user_id}
│   │   │   │
│   │   │   ├── jobs/
│   │   │   │   └── page.tsx                # Job performance analytics
│   │   │   │                               # BE: analytics-be — GET /v1/analytics/jobs?client_id={user_id}
│   │   │   │
│   │   │   ├── spending/
│   │   │   │   └── page.tsx                # Spending analytics
│   │   │   │                               # BE: analytics-be — GET /v1/analytics/spending?client_id={user_id}
│   │   │   │
│   │   │   └── reports/
│   │   │       └── page.tsx                # Custom reports
│   │   │                                   # BE: analytics-be — POST /v1/analytics/reports/generate
│   │   │
│   │   ├── organization/                   # Organization management (client only)
│   │   │   ├── page.tsx                    # Organization overview
│   │   │   │                               # BE: users-be/org — GET /v1/organizations/{org_id}
│   │   │   │
│   │   │   ├── team/
│   │   │   │   ├── page.tsx                # Team members
│   │   │   │   │                           # BE: users-be/team — GET /v1/organizations/{org_id}/team
│   │   │   │   │
│   │   │   │   ├── invite/
│   │   │   │   │   └── page.tsx            # Invite team member
│   │   │   │   │                           # BE: users-be/team — POST /v1/organizations/{org_id}/team/invite
│   │   │   │   │
│   │   │   │   └── [memberId]/
│   │   │   │       └── page.tsx            # Member details & permissions
│   │   │   │                               # BE: users-be/team — GET /v1/organizations/{org_id}/team/{member_id}
│   │   │   │                               # BE: users-be/role — PUT /v1/organizations/{org_id}/team/{member_id}/role
│   │   │   │
│   │   │   ├── billing/
│   │   │   │   └── page.tsx                # Organization billing
│   │   │   │                               # BE: financial-be/billing_profile — GET /v1/billing/profile?org_id={org_id}
│   │   │   │                               # BE: financial-be/invoice — GET /v1/invoices?org_id={org_id}
│   │   │   │
│   │   │   └── settings/
│   │   │       └── page.tsx                # Organization settings
│   │   │                                   # BE: users-be/org — PATCH /v1/organizations/{org_id}
│   │   │
│   │   ├── settings/                       # Account settings
│   │   │   ├── page.tsx                    # Settings overview
│   │   │   │
│   │   │   ├── account/
│   │   │   │   └── page.tsx                # Account settings
│   │   │   │                               # BE: users-be/user — GET /v1/users/{user_id}
│   │   │   │                               # BE: users-be/user — PATCH /v1/users/{user_id}
│   │   │   │
│   │   │   ├── security/
│   │   │   │   ├── page.tsx                # Security settings
│   │   │   │   │                           # BE: users-be/security — GET /v1/users/{user_id}/security
│   │   │   │   │
│   │   │   │   ├── password/
│   │   │   │   │   └── page.tsx            # Change password
│   │   │   │   │                           # BE: users-be/security — POST /v1/auth/change-password
│   │   │   │   │
│   │   │   │   ├── two-factor/
│   │   │   │   │   ├── page.tsx            # 2FA settings
│   │   │   │   │   │                       # BE: users-be/security/mfa — GET /v1/users/{user_id}/2fa
│   │   │   │   │   │
│   │   │   │   │   └── setup/
│   │   │   │   │       └── page.tsx        # Setup 2FA
│   │   │   │   │                           # BE: users-be/security/mfa — POST /v1/users/{user_id}/2fa/enable
│   │   │   │   │
│   │   │   │   └── sessions/
│   │   │   │       └── page.tsx            # Active sessions
│   │   │   │                               # BE: users-be/security/session — GET /v1/users/{user_id}/sessions
│   │   │   │                               # BE: users-be/security/session — DELETE /v1/users/{user_id}/sessions/{session_id}
│   │   │   │
│   │   │   ├── privacy/
│   │   │   │   └── page.tsx                # Privacy & data settings
│   │   │   │                               # BE: users-be/privacy — GET /v1/users/{user_id}/privacy
│   │   │   │                               # BE: users-be/privacy — PATCH /v1/users/{user_id}/privacy
│   │   │   │                               # BE: users-be/data — POST /v1/users/{user_id}/data-export (GDPR)
│   │   │   │                               # BE: users-be/data — POST /v1/users/{user_id}/data-deletion (GDPR)
│   │   │   │
│   │   │   └── close-account/
│   │   │       └── page.tsx                # Close account
│   │   │                                   # BE: users-be/account — POST /v1/users/{user_id}/deactivate
│   │   │
│   │   └── layout.tsx                      # Dashboard layout (header, sidebar, notifications)
│   │
│   ├── (admin)/                            # Admin panel (ADMIN role only)
│   │   ├── admin/
│   │   │   ├── page.tsx                    # Admin dashboard
│   │   │   │                               # BE: admin-be/dashboard — GET /v1/admin/dashboard
│   │   │   │
│   │   │   ├── users/
│   │   │   │   ├── page.tsx                # User management
│   │   │   │   │                           # BE: admin-be/user_mgmt — GET /v1/admin/users
│   │   │   │   │
│   │   │   │   └── [userId]/
│   │   │   │       ├── page.tsx            # User details & actions
│   │   │   │       │                       # BE: admin-be/user_mgmt — GET /v1/admin/users/{user_id}
│   │   │   │       │                       # BE: admin-be/user_mgmt — POST /v1/admin/users/{user_id}/suspend
│   │   │   │       │                       # BE: admin-be/user_mgmt — POST /v1/admin/users/{user_id}/ban
│   │   │   │       │
│   │   │   │       └── audit/
│   │   │   │           └── page.tsx        # User audit trail
│   │   │   │                               # BE: admin-be/audit — GET /v1/admin/audit?user_id={user_id}
│   │   │   │
│   │   │   ├── kyc/
│   │   │   │   ├── page.tsx                # KYC cases queue
│   │   │   │   │                           # BE: admin-be/kyc_case — GET /v1/admin/kyc/cases
│   │   │   │   │
│   │   │   │   └── [caseId]/
│   │   │   │       └── page.tsx            # KYC case review
│   │   │   │                               # BE: admin-be/kyc_case — GET /v1/admin/kyc/cases/{case_id}
│   │   │   │                               # BE: admin-be/kyc_case — POST /v1/admin/kyc/cases/{case_id}/approve
│   │   │   │                               # BE: admin-be/kyc_case — POST /v1/admin/kyc/cases/{case_id}/reject
│   │   │   │
│   │   │   ├── disputes/
│   │   │   │   ├── page.tsx                # Dispute cases
│   │   │   │   │                           # BE: admin-be/case_mgmt — GET /v1/admin/disputes
│   │   │   │   │
│   │   │   │   └── [disputeId]/
│   │   │   │       └── page.tsx            # Dispute resolution
│   │   │   │                               # BE: admin-be/case_mgmt — GET /v1/admin/disputes/{dispute_id}
│   │   │   │                               # BE: admin-be/case_mgmt — POST /v1/admin/disputes/{dispute_id}/resolve
│   │   │   │
│   │   │   ├── refunds/
│   │   │   │   └── page.tsx                # Refund cases
│   │   │   │                               # BE: admin-be/refund_case — GET /v1/admin/refunds
│   │   │   │                               # BE: admin-be/refund_case — POST /v1/admin/refunds/{case_id}/approve
│   │   │   │
│   │   │   ├── moderation/
│   │   │   │   └── page.tsx                # Content moderation queue
│   │   │   │                               # BE: admin-be/moderation — GET /v1/admin/moderation/queue
│   │   │   │                               # BE: admin-be/moderation — POST /v1/admin/moderation/{item_id}/action
│   │   │   │
│   │   │   └── analytics/
│   │   │       └── page.tsx                # Platform analytics
│   │   │                                   # BE: admin-be/analytics — GET /v1/admin/analytics
│   │   │
│   │   └── layout.tsx                      # Admin layout
│   │
│   ├── globals.css                         # Global Tailwind styles
│   ├── layout.tsx                          # Root layout (providers, fonts, metadata)
│   ├── page.tsx                            # Root redirect (to [locale])
│   └── providers.tsx                       # Client providers (TanStack Query, Keycloak, Zustand)
│                                           # TanStack Query: default staleTime=30s, gcTime=10min
│                                           # Query keys: ['domain', 'action', ...params]
│
├── landing/                                # Landing page sections (imported by (landing)/page.tsx)
│   ├── Hero.tsx                            # Hero section
│   ├── Features.tsx                        # Features showcase
│   ├── Stats.tsx                           # Platform statistics
│   ├── Testimonials.tsx                    # User testimonials
│   └── CTA.tsx                             # Call-to-action section
│
└── layout/                                 # Layout components
    ├── Header.tsx                          # Public header
    ├── DashboardHeader.tsx                 # Dashboard header (user menu, notifications)
    ├── Footer.tsx                          # Public footer
    └── LanguageSwitcher.tsx                # Language switcher component
```

### Components Structure (`apps/web/src/components`)

```
src/components/
├── ui/                                     # Base UI components (shadcn/ui style)
│   ├── Button/
│   │   ├── Button.tsx
│   │   └── Button.stories.tsx
│   ├── Input/
│   │   ├── Input.tsx
│   │   └── Input.stories.tsx
│   ├── Card/
│   ├── Modal/
│   ├── Dropdown/
│   ├── Tabs/
│   ├── Badge/
│   ├── Avatar/
│   ├── Progress/
│   ├── Skeleton/
│   ├── Toast/
│   ├── Dialog/
│   ├── Select/
│   ├── Checkbox/
│   ├── Radio/
│   ├── Switch/
│   ├── Textarea/
│   ├── Alert/
│   ├── Tooltip/
│   └── Pagination/
│
├── forms/                                  # Form components
│   ├── FormField/
│   │   └── FormField.tsx                   # Generic form field wrapper
│   ├── FormError/
│   │   └── FormError.tsx                   # Error message display
│   ├── FormSection/
│   │   └── FormSection.tsx                 # Form section grouping
│   └── FileUpload/
│       └── FileUpload.tsx                  # File upload component
│                                           # BE: storage-be/uploads — POST /v1/storage/upload
│                                           # Signed URL flow, chunk uploads for large files
│
├── layouts/                                # Layout components
│   ├── DashboardLayout/
│   │   ├── DashboardLayout.tsx
│   │   ├── Sidebar.tsx
│   │   └── MobileMenu.tsx
│   └── AuthLayout/
│       └── AuthLayout.tsx
│
├── navigation/                             # Navigation components
│   ├── Breadcrumbs/
│   │   └── Breadcrumbs.tsx
│   ├── NavItem/
│   │   └── NavItem.tsx
│   └── Pagination/
│       └── Pagination.tsx
│
└── common/                                 # Common/shared components
    ├── ErrorBoundary/
    │   └── ErrorBoundary.tsx               # Error boundary wrapper
    ├── Loading/
    │   ├── LoadingSpinner.tsx
    │   ├── LoadingSkeleton.tsx
    │   └── PageLoader.tsx
    ├── EmptyState/
    │   └── EmptyState.tsx                  # Empty state placeholder
    ├── ErrorState/
    │   └── ErrorState.tsx                  # Error state display
    ├── ConfirmDialog/
    │   └── ConfirmDialog.tsx               # Confirmation dialog
    ├── SearchBar/
    │   └── SearchBar.tsx                   # Search input with autocomplete
    │                                       # BE: search-be/autocomplete — GET /v1/search/autocomplete
    │
    ├── NotificationBell/
    │   └── NotificationBell.tsx            # Notification bell icon with badge
    │                                       # BE: communications-be/notifications — GET /v1/notifications/unread-count
    │                                       # WebSocket: real-time notification updates
    │
    ├── UserMenu/
    │   └── UserMenu.tsx                    # User dropdown menu
    │                                       # Includes: profile, settings, logout
    │
    └── LanguageSelector/
        └── LanguageSelector.tsx            # Language selection dropdown
```

### Features Structure (`apps/web/src/features`)

```
src/features/
├── auth/                                   # Authentication feature
│   ├── components/
│   │   ├── LoginForm.tsx                   # Login form
│   │   │                                   # BE: Keycloak OAuth2 flow
│   │   ├── RegisterForm.tsx                # Registration form
│   │   │                                   # BE: users-be/user — POST /v1/users/register
│   │   ├── ForgotPasswordForm.tsx          # Forgot password form
│   │   │                                   # BE: users-be/security/recovery — POST /v1/auth/forgot-password
│   │   ├── ResetPasswordForm.tsx           # Reset password form
│   │   │                                   # BE: users-be/security/recovery — POST /v1/auth/reset-password
│   │   └── SocialButtons.tsx               # Social login buttons (Google, GitHub)
│   │                                       # BE: Keycloak social identity providers
│   │
│   ├── hooks/
│   │   ├── useAuth.ts                      # Auth context hook
│   │   ├── useLogin.ts                     # Login mutation hook
│   │   ├── useRegister.ts                  # Register mutation hook
│   │   └── useLogout.ts                    # Logout hook
│   │
│   └── api/
│       ├── auth.queries.ts                 # Auth queries (TanStack Query)
│       │                                   # BE: users-be/auth — GET /v1/auth/me
│       │                                   # Query key: ['auth', 'me']
│       │
│       └── auth.mutations.ts               # Auth mutations (TanStack Query)
│                                           # BE: users-be/auth — POST /v1/auth/login
│                                           # BE: users-be/auth — POST /v1/auth/logout
│                                           # BE: users-be/auth — POST /v1/auth/refresh
│                                           # Invalidates: ['auth', 'me']
│
├── profile/                                # Profile management feature
│   ├── components/
│   │   ├── ProfileCard.tsx                 # Profile overview card
│   │   ├── ProfileForm.tsx                 # Edit profile form
│   │   ├── ProfileCompleteness.tsx         # Profile completeness indicator
│   │   ├── SkillsSection.tsx               # Skills display/edit
│   │   ├── SkillForm.tsx                   # Add/edit skill form
│   │   ├── ExperienceList.tsx              # Experience timeline
│   │   ├── ExperienceForm.tsx              # Add/edit experience
│   │   ├── EducationList.tsx               # Education list
│   │   ├── EducationForm.tsx               # Add/edit education
│   │   ├── PortfolioGrid.tsx               # Portfolio items grid
│   │   ├── PortfolioCard.tsx               # Portfolio item card
│   │   ├── PortfolioForm.tsx               # Add/edit portfolio item
│   │   ├── ServiceCatalogList.tsx          # Services list
│   │   ├── ServiceForm.tsx                 # Add/edit service
│   │   ├── AvailabilityCalendar.tsx        # Availability calendar
│   │   └── AvailabilityForm.tsx            # Update availability
│   │
│   ├── hooks/
│   │   ├── useProfile.ts                   # Profile query hook
│   │   ├── useUpdateProfile.ts             # Update profile mutation
│   │   ├── useSkills.ts                    # Skills queries
│   │   ├── useExperience.ts                # Experience queries
│   │   ├── useEducation.ts                 # Education queries
│   │   ├── usePortfolio.ts                 # Portfolio queries
│   │   └── useAvailability.ts              # Availability queries
│   │
│   └── api/
│       ├── profile.queries.ts              # Profile queries
│       │                                   # BE: users-be/profile — GET /v1/users/{id}/profile
│       │                                   # Query key: ['profile', user_id]
│       │                                   # staleTime: 60s
│       │
│       ├── profile.mutations.ts            # Profile mutations
│       │                                   # BE: users-be/profile — PATCH /v1/users/{id}/profile
│       │                                   # Invalidates: ['profile', user_id]
│       │
│       ├── skills.queries.ts               # Skills queries
│       │                                   # BE: users-be/capabilities — GET /v1/users/{id}/skills
│       │                                   # Query key: ['skills', user_id]
│       │
│       ├── skills.mutations.ts             # Skills mutations
│       │                                   # BE: users-be/capabilities — POST/PUT/DELETE /v1/users/{id}/skills
│       │                                   # Invalidates: ['skills', user_id], ['profile', user_id]
│       │
│       ├── portfolio.queries.ts            # Portfolio queries
│       │                                   # BE: users-be/portfolio — GET /v1/users/{id}/portfolio
│       │                                   # Query key: ['portfolio', user_id]
│       │
│       └── portfolio.mutations.ts          # Portfolio mutations
│                                           # BE: users-be/portfolio — POST/PUT/DELETE /v1/portfolio/{item_id}
│                                           # Invalidates: ['portfolio', user_id]
│
├── jobs/                                   # Job browsing & posting feature
│   ├── components/
│   │   ├── JobCard.tsx                     # Job listing card
│   │   ├── JobList.tsx                     # Job list view
│   │   ├── JobGrid.tsx                     # Job grid view
│   │   ├── JobDetails.tsx                  # Job details display
│   │   ├── JobFilters.tsx                  # Job search filters
│   │   ├── JobSearchBar.tsx                # Job search input
│   │   ├── JobForm.tsx                     # Job posting form (multi-step)
│   │   ├── JobDetailsStep.tsx              # Job details form step
│   │   ├── JobBudgetStep.tsx               # Budget & pricing step
│   │   ├── JobRequirementsStep.tsx         # Requirements step
│   │   ├── JobReviewStep.tsx               # Review & publish step
│   │   ├── SavedJobsList.tsx               # Saved jobs list
│   │   └── RecommendedJobs.tsx             # Recommended jobs widget
│   │
│   ├── hooks/
│   │   ├── useJobs.ts                      # Jobs search query
│   │   ├── useJob.ts                       # Single job query
│   │   ├── useCreateJob.ts                 # Create job mutation
│   │   ├── useUpdateJob.ts                 # Update job mutation
│   │   ├── usePublishJob.ts                # Publish job mutation
│   │   ├── useSavedJobs.ts                 # Saved jobs query
│   │   └── useJobRecommendations.ts        # Recommended jobs query
│   │
│   └── api/
│       ├── jobs.queries.ts                 # Job queries
│       │                                   # BE: search-be/search — POST /v1/search/jobs
│       │                                   # Query key: ['jobs', 'list', filters]
│       │                                   # Pagination: cursor-based, keepPreviousData: true
│       │                                   # staleTime: 30s
│       │                                   #
│       │                                   # BE: jobs-be/job — GET /v1/jobs/{job_id}
│       │                                   # Query key: ['jobs', 'detail', job_id]
│       │                                   # staleTime: 60s
│       │
│       ├── jobs.mutations.ts               # Job mutations
│       │                                   # BE: jobs-be/job — POST /v1/jobs
│       │                                   # BE: jobs-be/job — PATCH /v1/jobs/{job_id}
│       │                                   # BE: jobs-be/job — POST /v1/jobs/{job_id}/publish
│       │                                   # Invalidates: ['jobs', 'list'], ['jobs', 'detail', job_id]
│       │
│       └── saved-jobs.queries.ts           # Saved jobs queries
│                                           # BE: search-be/saved_searches — GET /v1/search/saved
│                                           # Query key: ['saved-jobs', user_id]
│
├── proposals/                              # Proposal management feature
│   ├── components/
│   │   ├── ProposalCard.tsx                # Proposal card
│   │   ├── ProposalList.tsx                # Proposals list
│   │   ├── ProposalDetails.tsx             # Proposal details view
│   │   ├── ProposalForm.tsx                # Submit proposal form
│   │   ├── MilestoneForm.tsx               # Define milestones
│   │   ├── CoverLetterInput.tsx            # Cover letter editor
│   │   ├── QuestionAnswers.tsx             # Screening questions answers
│   │   ├── ProposalStats.tsx               # Proposal statistics
│   │   └── BidStrategy.tsx                 # Bidding strategy selector
│   │
│   ├── hooks/
│   │   ├── useProposals.ts                 # Proposals list query
│   │   ├── useProposal.ts                  # Single proposal query
│   │   ├── useSubmitProposal.ts            # Submit proposal mutation
│   │   ├── useUpdateProposal.ts            # Update proposal mutation
│   │   └── useWithdrawProposal.ts          # Withdraw proposal mutation
│   │
│   └── api/
│       ├── proposals.queries.ts            # Proposal queries
│       │                                   # BE: proposals-be/proposal — GET /v1/proposals?freelancer_id={id}
│       │                                   # Query key: ['proposals', 'list', filters]
│       │                                   # Pagination: cursor-based
│       │                                   #
│       │                                   # BE: proposals-be/proposal — GET /v1/proposals/{proposal_id}
│       │                                   # Query key: ['proposals', 'detail', proposal_id]
│       │
│       └── proposals.mutations.ts          # Proposal mutations
│                                           # BE: proposals-be/proposal — POST /v1/proposals
│                                           # BE: proposals-be/proposal — PUT /v1/proposals/{proposal_id}
│                                           # BE: proposals-be/proposal — POST /v1/proposals/{proposal_id}/withdraw
│                                           # Invalidates: ['proposals', 'list'], ['proposals', 'detail', proposal_id]
│                                           # Optimistic updates: proposal status changes
│
├── contracts/                              # Contract management feature
│   ├── components/
│   │   ├── ContractCard.tsx                # Contract card
│   │   ├── ContractList.tsx                # Contracts list
│   │   ├── ContractDetails.tsx             # Contract details view
│   │   ├── ContractTimeline.tsx            # Contract timeline
│   │   ├── MilestoneCard.tsx               # Milestone card
│   │   ├── MilestoneList.tsx               # Milestones list
│   │   ├── MilestoneForm.tsx               # Create milestone form
│   │   ├── MilestoneSubmission.tsx         # Submit milestone form
│   │   ├── MilestoneReview.tsx             # Review milestone (approve/reject)
│   │   ├── DeliverableUpload.tsx           # Upload deliverables
│   │   ├── TimesheetForm.tsx               # Log hours form
│   │   ├── TimesheetList.tsx               # Timesheet entries list
│   │   ├── WorkDiaryView.tsx               # Work diary view
│   │   ├── DisputeForm.tsx                 # Create dispute form
│   │   ├── DisputeDetails.tsx              # Dispute details view
│   │   └── ContractMessaging.tsx           # Contract-specific chat
│   │
│   ├── hooks/
│   │   ├── useContracts.ts                 # Contracts list query
│   │   ├── useContract.ts                  # Single contract query
│   │   ├── useMilestones.ts                # Milestones query
│   │   ├── useCreateMilestone.ts           # Create milestone mutation
│   │   ├── useSubmitMilestone.ts           # Submit milestone mutation
│   │   ├── useApproveMilestone.ts          # Approve milestone mutation
│   │   ├── useRejectMilestone.ts           # Reject milestone mutation
│   │   ├── useTimesheets.ts                # Timesheets query
│   │   ├── useLogHours.ts                  # Log hours mutation
│   │   └── useEndContract.ts               # End contract mutation
│   │
│   └── api/
│       ├── contracts.queries.ts            # Contract queries
│       │                                   # BE: contracts-be/contract — GET /v1/contracts?user_id={id}
│       │                                   # Query key: ['contracts', 'list', filters]
│       │                                   #
│       │                                   # BE: contracts-be/contract — GET /v1/contracts/{contract_id}
│       │                                   # Query key: ['contracts', 'detail', contract_id]
│       │
│       ├── contracts.mutations.ts          # Contract mutations
│       │                                   # BE: contracts-be/contract — POST /v1/contracts/{contract_id}/end
│       │                                   # Invalidates: ['contracts', 'detail', contract_id], ['contracts', 'list']
│       │
│       ├── milestones.queries.ts           # Milestone queries
│       │                                   # BE: contracts-be/milestones — GET /v1/contracts/{contract_id}/milestones
│       │                                   # Query key: ['milestones', 'list', contract_id]
│       │
│       ├── milestones.mutations.ts         # Milestone mutations
│       │                                   # BE: contracts-be/milestones — POST /v1/milestones
│       │                                   # BE: contracts-be/milestones — POST /v1/milestones/{id}/submit
│       │                                   # BE: contracts-be/milestones — POST /v1/milestones/{id}/approve
│       │                                   # BE: contracts-be/milestones — POST /v1/milestones/{id}/reject
│       │                                   # Invalidates: ['milestones', 'list', contract_id], ['contracts', 'detail', contract_id]
│       │                                   # Optimistic updates: milestone status changes
│       │
│       └── timesheets.queries.ts           # Timesheet queries
│                                           # BE: contracts-be/timesheet — GET /v1/contracts/{contract_id}/timesheets
│                                           # Query key: ['timesheets', contract_id]
│
├── payments/                               # Payment & financial feature
│   ├── components/
│   │   ├── WalletCard.tsx                  # Wallet balance card
│   │   ├── TransactionList.tsx             # Transaction history
│   │   ├── TransactionItem.tsx             # Transaction item
│   │   ├── PaymentMethodCard.tsx           # Payment method card
│   │   ├── PaymentMethodForm.tsx           # Add payment method form
│   │   ├── DepositForm.tsx                 # Deposit funds form
│   │   ├── WithdrawForm.tsx                # Withdraw funds form
│   │   ├── InvoiceCard.tsx                 # Invoice card
│   │   ├── InvoiceDetails.tsx              # Invoice details view
│   │   ├── RefundCard.tsx                  # Refund card
│   │   └── TaxDocuments.tsx                # Tax documents list
│   │
│   ├── hooks/
│   │   ├── useWallet.ts                    # Wallet query
│   │   ├── useTransactions.ts              # Transactions query
│   │   ├── usePaymentMethods.ts            # Payment methods query
│   │   ├── useAddPaymentMethod.ts          # Add payment method mutation
│   │   ├── useDeposit.ts                   # Deposit mutation
│   │   ├── useWithdraw.ts                  # Withdraw mutation
│   │   ├── useInvoices.ts                  # Invoices query
│   │   └── useTaxProfile.ts                # Tax profile query
│   │
│   └── api/
│       ├── wallet.queries.ts               # Wallet queries
│       │                                   # BE: financial-be/wallet — GET /v1/wallets/{user_id}
│       │                                   # Query key: ['wallet', user_id]
│       │                                   # staleTime: 10s (frequently updated)
│       │
│       ├── wallet.mutations.ts             # Wallet mutations
│       │                                   # BE: financial-be/wallet — POST /v1/wallets/{wallet_id}/deposit
│       │                                   # BE: financial-be/payout — POST /v1/payouts/request
│       │                                   # Invalidates: ['wallet', user_id], ['transactions', user_id]
│       │                                   # Optimistic updates: balance changes
│       │
│       ├── transactions.queries.ts         # Transaction queries
│       │                                   # BE: financial-be/transactions — GET /v1/transactions?user_id={id}
│       │                                   # Query key: ['transactions', user_id, filters]
│       │                                   # Pagination: cursor-based
│       │
│       ├── payment-methods.queries.ts      # Payment method queries
│       │                                   # BE: financial-be/payment_method — GET /v1/payment-methods?user_id={id}
│       │                                   # Query key: ['payment-methods', user_id]
│       │
│       ├── payment-methods.mutations.ts    # Payment method mutations
│       │                                   # BE: financial-be/payment_method — POST /v1/payment-methods
│       │                                   # BE: financial-be/payment_method — DELETE /v1/payment-methods/{id}
│       │                                   # Invalidates: ['payment-methods', user_id]
│       │
│       └── invoices.queries.ts             # Invoice queries
│                                           # BE: financial-be/invoice — GET /v1/invoices?user_id={id}
│                                           # Query key: ['invoices', user_id, filters]
│
├── messaging/                              # Messaging feature
│   ├── components/
│   │   ├── ConversationList.tsx            # Conversations list
│   │   ├── ConversationItem.tsx            # Conversation list item
│   │   ├── MessageThread.tsx               # Message thread view
│   │   ├── MessageBubble.tsx               # Message bubble
│   │   ├── MessageInput.tsx                # Message input composer
│   │   ├── TypingIndicator.tsx             # Typing indicator
│   │   ├── ReadReceipts.tsx                # Read receipts
│   │   └── AttachmentPreview.tsx           # Attachment preview
│   │
│   ├── hooks/
│   │   ├── useConversations.ts             # Conversations query
│   │   ├── useMessages.ts                  # Messages query (with WebSocket)
│   │   ├── useSendMessage.ts               # Send message mutation
│   │   ├── useMarkRead.ts                  # Mark as read mutation
│   │   └── useTypingIndicator.ts           # Typing indicator hook
│   │
│   └── api/
│       ├── conversations.queries.ts        # Conversation queries
│       │                                   # BE: communications-be/conversations — GET /v1/conversations?user_id={id}
│       │                                   # Query key: ['conversations', user_id]
│       │                                   # WebSocket: wss://api/v1/ws/user/{user_id}
│       │
│       ├── messages.queries.ts             # Message queries
│       │                                   # BE: communications-be/messages — GET /v1/conversations/{conv_id}/messages
│       │                                   # Query key: ['messages', conversation_id]
│       │                                   # WebSocket: wss://api/v1/ws/conversations/{conv_id}
│       │                                   # Real-time updates via WebSocket
│       │
│       └── messages.mutations.ts           # Message mutations
│                                           # BE: communications-be/messages — POST /v1/messages
│                                           # BE: communications-be/messages — POST /v1/messages/{msg_id}/read
│                                           # Invalidates: ['conversations', user_id], ['messages', conversation_id]
│                                           # Optimistic updates: new message appears immediately
│
├── notifications/                          # Notifications feature
│   ├── components/
│   │   ├── NotificationList.tsx            # Notifications list
│   │   ├── NotificationItem.tsx            # Notification item
│   │   ├── NotificationBadge.tsx           # Unread count badge
│   │   └── NotificationSettings.tsx        # Notification preferences
│   │
│   ├── hooks/
│   │   ├── useNotifications.ts             # Notifications query
│   │   ├── useUnreadCount.ts               # Unread count query
│   │   ├── useMarkNotificationRead.ts      # Mark as read mutation
│   │   └── useMarkAllRead.ts               # Mark all as read mutation
│   │
│   └── api/
│       ├── notifications.queries.ts        # Notification queries
│       │                                   # BE: communications-be/notifications — GET /v1/notifications?user_id={id}
│       │                                   # Query key: ['notifications', user_id]
│       │                                   # WebSocket: wss://api/v1/ws/notifications/{user_id}
│       │                                   #
│       │                                   # BE: communications-be/notifications — GET /v1/notifications/unread-count
│       │                                   # Query key: ['notifications', 'unread-count', user_id]
│       │                                   # staleTime: 5s (very fresh)
│       │
│       └── notifications.mutations.ts      # Notification mutations
│                                           # BE: communications-be/notifications — POST /v1/notifications/{id}/read
│                                           # BE: communications-be/notifications — POST /v1/notifications/read-all
│                                           # Invalidates: ['notifications', user_id], ['notifications', 'unread-count']
│                                           # Optimistic updates: mark as read immediately
│
├── subscriptions/                          # Subscription management feature
│   ├── components/
│   │   ├── PlanCard.tsx                    # Subscription plan card
│   │   ├── PlanComparison.tsx              # Plans comparison table
│   │   ├── CurrentSubscription.tsx         # Current subscription overview
│   │   ├── UpgradeFlow.tsx                 # Upgrade flow wizard
│   │   ├── AddonCard.tsx                   # Addon card
│   │   ├── ConnectsBalance.tsx             # Connects balance widget
│   │   └── ConnectsPurchase.tsx            # Purchase connects form
│   │
│   ├── hooks/
│   │   ├── usePlans.ts                     # Plans query
│   │   ├── useSubscription.ts              # User subscription query
│   │   ├── useEntitlements.ts              # User entitlements query
│   │   ├── useUpgrade.ts                   # Upgrade subscription mutation
│   │   ├── useAddons.ts                    # Addons query
│   │   ├── useConnectsBalance.ts           # Connects balance query
│   │   └── usePurchaseConnects.ts          # Purchase connects mutation
│   │
│   └── api/
│       ├── plans.queries.ts                # Plans queries
│       │                                   # BE: subscriptions-be/plans — GET /v1/plans
│       │                                   # Query key: ['plans']
│       │                                   # staleTime: 5min (rarely changes)
│       │
│       ├── subscription.queries.ts         # Subscription queries
│       │                                   # BE: subscriptions-be/subscriptions — GET /v1/subscriptions?user_id={id}
│       │                                   # Query key: ['subscription', user_id]
│       │                                   #
│       │                                   # BE: subscriptions-be/entitlements — GET /v1/entitlements?user_id={id}
│       │                                   # Query key: ['entitlements', user_id]
│       │
│       ├── subscription.mutations.ts       # Subscription mutations
│       │                                   # BE: subscriptions-be/subscriptions — POST /v1/subscriptions/{id}/upgrade
│       │                                   # BE: subscriptions-be/subscriptions — POST /v1/subscriptions/{id}/cancel
│       │                                   # Invalidates: ['subscription', user_id], ['entitlements', user_id]
│       │
│       └── connects.queries.ts             # Connects queries
│                                           # BE: subscriptions-be/connects — GET /v1/connects/balance?user_id={id}
│                                           # Query key: ['connects', 'balance', user_id]
│                                           #
│                                           # BE: subscriptions-be/connects — POST /v1/connects/purchase
│                                           # Invalidates: ['connects', 'balance', user_id]
│
├── reviews/                                # Reviews feature
│   ├── components/
│   │   ├── ReviewCard.tsx                  # Review card
│   │   ├── ReviewList.tsx                  # Reviews list
│   │   ├── ReviewForm.tsx                  # Write review form
│   │   ├── RatingStars.tsx                 # Star rating component
│   │   ├── ReviewStats.tsx                 # Review statistics
│   │   └── ReviewFilters.tsx               # Review filters
│   │
│   ├── hooks/
│   │   ├── useReviews.ts                   # Reviews query
│   │   ├── useSubmitReview.ts              # Submit review mutation
│   │   └── useReviewStats.ts               # Review statistics query
│   │
│   └── api/
│       ├── reviews.queries.ts              # Review queries
│       │                                   # BE: reviews-be/reviews — GET /v1/reviews?user_id={id}
│       │                                   # Query key: ['reviews', 'list', user_id, filters]
│       │                                   #
│       │                                   # BE: reviews-be/reviews — GET /v1/reviews/stats?user_id={id}
│       │                                   # Query key: ['reviews', 'stats', user_id]
│       │
│       └── reviews.mutations.ts            # Review mutations
│                                           # BE: reviews-be/reviews — POST /v1/reviews
│                                           # Invalidates: ['reviews', 'list'], ['reviews', 'stats']
│                                           # Triggers: review.submitted.v1 → reputation update
│
├── analytics/                              # Analytics feature (client only)
│   ├── components/
│   │   ├── AnalyticsDashboard.tsx          # Main analytics dashboard
│   │   ├── MetricCard.tsx                  # Metric card widget
│   │   ├── ChartWidget.tsx                 # Chart widget
│   │   ├── JobPerformance.tsx              # Job performance analytics
│   │   ├── SpendingChart.tsx               # Spending analytics
│   │   └── CustomReport.tsx                # Custom report builder
│   │
│   ├── hooks/
│   │   ├── useAnalytics.ts                 # Analytics overview query
│   │   ├── useJobAnalytics.ts              # Job analytics query
│   │   ├── useSpendingAnalytics.ts         # Spending analytics query
│   │   └── useGenerateReport.ts            # Generate report mutation
│   │
│   └── api/
│       └── analytics.queries.ts            # Analytics queries
│                                           # BE: analytics-be — GET /v1/analytics/overview?user_id={id}
│                                           # BE: analytics-be — GET /v1/analytics/jobs?client_id={id}
│                                           # BE: analytics-be — GET /v1/analytics/spending?client_id={id}
│                                           # BE: analytics-be — POST /v1/analytics/reports/generate
│                                           # Query keys: ['analytics', 'overview'], ['analytics', 'jobs'], etc.
│
├── admin/                                  # Admin feature (admin only)
│   ├── components/
│   │   ├── AdminDashboard.tsx              # Admin dashboard
│   │   ├── UserManagement.tsx              # User management table
│   │   ├── KYCQueue.tsx                    # KYC cases queue
│   │   ├── KYCCaseCard.tsx                 # KYC case review card
│   │   ├── DisputeQueue.tsx                # Dispute cases queue
│   │   ├── DisputeResolution.tsx           # Dispute resolution interface
│   │   ├── RefundQueue.tsx                 # Refund cases queue
│   │   ├── ModerationQueue.tsx             # Content moderation queue
│   │   └── PlatformAnalytics.tsx           # Platform analytics
│   │
│   ├── hooks/
│   │   ├── useAdminDashboard.ts            # Admin dashboard query
│   │   ├── useUserManagement.ts            # User management queries
│   │   ├── useKYCCases.ts                  # KYC cases query
│   │   ├── useApproveKYC.ts                # Approve KYC mutation
│   │   ├── useDisputeCases.ts              # Dispute cases query
│   │   ├── useResolveDispute.ts            # Resolve dispute mutation
│   │   └── useRefundCases.ts               # Refund cases query
│   │
│   └── api/
│       ├── admin.queries.ts                # Admin queries
│       │                                   # BE: admin-be/dashboard — GET /v1/admin/dashboard
│       │                                   # BE: admin-be/user_mgmt — GET /v1/admin/users
│       │                                   # BE: admin-be/kyc_case — GET /v1/admin/kyc/cases
│       │                                   # BE: admin-be/case_mgmt — GET /v1/admin/disputes
│       │                                   # BE: admin-be/refund_case — GET /v1/admin/refunds
│       │                                   # Query keys: ['admin', 'dashboard'], ['admin', 'kyc-cases'], etc.
│       │
│       └── admin.mutations.ts              # Admin mutations
│                                           # BE: admin-be/user_mgmt — POST /v1/admin/users/{id}/suspend
│                                           # BE: admin-be/user_mgmt — POST /v1/admin/users/{id}/ban
│                                           # BE: admin-be/kyc_case — POST /v1/admin/kyc/cases/{id}/approve
│                                           # BE: admin-be/case_mgmt — POST /v1/admin/disputes/{id}/resolve
│                                           # Invalidates: respective query keys
│
└── search/                                 # Search feature
    ├── components/
    │   ├── SearchInput.tsx                 # Search input with autocomplete
    │   ├── SearchFilters.tsx               # Search filters panel
    │   ├── SearchResults.tsx               # Search results list
    │   ├── SearchSuggestions.tsx           # Search suggestions
    │   └── SavedSearches.tsx               # Saved searches list
    │
    ├── hooks/
    │   ├── useSearch.ts                    # Search query
    │   ├── useAutocomplete.ts              # Autocomplete query
    │   ├── useSavedSearches.ts             # Saved searches query
    │   └── useSaveSearch.ts                # Save search mutation
    │
    └── api/
        ├── search.queries.ts               # Search queries
        │                                   # BE: search-be/search — POST /v1/search/jobs
        │                                   # BE: search-be/search — POST /v1/search/freelancers
        │                                   # BE: search-be/autocomplete — GET /v1/search/autocomplete
        │                                   # Query keys: ['search', 'jobs', query], ['search', 'freelancers', query]
        │                                   # staleTime: 30s
        │
        └── saved-searches.queries.ts       # Saved searches queries
                                            # BE: search-be/saved_searches — GET /v1/search/saved
                                            # BE: search-be/saved_searches — POST /v1/search/saved
                                            # Query key: ['saved-searches', user_id]
```

### Hooks Structure (`apps/web/src/hooks`)

```
src/hooks/
├── useDebounce.ts                          # Debounce hook
├── useMediaQuery.ts                        # Media query hook (responsive)
├── useLocalStorage.ts                      # Local storage hook
├── useSessionStorage.ts                    # Session storage hook
├── usePagination.ts                        # Pagination logic hook
├── useInfiniteScroll.ts                    # Infinite scroll hook
├── useWebSocket.ts                         # WebSocket connection hook
│                                           # BE: wss://api/v1/ws/* (various WebSocket endpoints)
│                                           # Auto-reconnect, heartbeat, message queue
│
├── useUpload.ts                            # File upload hook
│                                           # BE: storage-be/uploads — POST /v1/storage/upload
│                                           # Handles: signed URLs, chunked uploads, progress tracking
│
├── useClipboard.ts                         # Clipboard operations hook
├── useKeyboard.ts                          # Keyboard shortcuts hook
├── useOnClickOutside.ts                    # Click outside detection hook
├── useToggle.ts                            # Toggle state hook
├── useInterval.ts                          # Interval hook
├── useTimeout.ts                           # Timeout hook
└── useFocus.ts                             # Focus management hook
```

### Lib Structure (`apps/web/src/lib`)

```
src/lib/
├── keycloak.ts                             # Keycloak client configuration
│                                           # BE: Keycloak server (users-be integrated)
│                                           # OAuth2 flows, token management
│
├── api/
│   ├── client.ts                           # API client (axios/fetch wrapper)
│   │                                       # Base URL, auth interceptors, error handling
│   │                                       # Auto-retry with exponential backoff
│   │
│   ├── websocket.ts                        # WebSocket client
│   │                                       # BE: wss://api/v1/ws/*
│   │                                       # Connection management, reconnection logic
│   │
│   └── endpoints.ts                        # API endpoint constants
│                                           # Centralized endpoint definitions
│
├── tanstack-query/
│   ├── client.ts                           # TanStack Query client configuration
│   │                                       # Default options: staleTime, gcTime, retry
│   │                                       # Query client with persister
│   │
│   └── keys.ts                             # Query key factory functions
│                                           # Standardized query key generation
│
├── zustand/
│   ├── store.ts                            # Root Zustand store
│   ├── slices/
│   │   ├── authSlice.ts                    # Auth state slice
│   │   ├── uiSlice.ts                      # UI state slice (theme, sidebar, etc.)
│   │   ├── notificationSlice.ts            # Notification state slice
│   │   └── filterSlice.ts                  # Filter state slice (search filters)
│   │
│   └── middleware/
│       ├── persist.ts                      # Persist middleware config
│       └── devtools.ts                     # Devtools middleware config
│
├── validation/
│   ├── schemas.ts                          # Zod validation schemas
│   │                                       # Reusable validation schemas for forms
│   │
│   └── rules.ts                            # Custom validation rules
│
├── utils/
│   ├── date.ts                             # Date formatting utilities
│   ├── currency.ts                         # Currency formatting utilities
│   ├── string.ts                           # String manipulation utilities
│   ├── array.ts                            # Array utilities
│   ├── object.ts                           # Object utilities
│   ├── number.ts                           # Number utilities
│   ├── file.ts                             # File utilities (size formatting, validation)
│   ├── url.ts                              # URL utilities
│   └── error.ts                            # Error handling utilities
│
├── i18n/
│   ├── config.ts                           # i18n configuration
│   ├── locales.ts                          # Locale definitions
│   └── formatters.ts                       # Locale-specific formatters
│
├── constants/
│   ├── routes.ts                           # Route constants
│   ├── config.ts                           # App configuration constants
│   ├── api.ts                              # API-related constants
│   └── app.ts                              # General app constants
│
└── types/
    ├── api.ts                              # API response types
    ├── query.ts                            # Query/filter types
    └── common.ts                           # Common types
```

---

## Apps - Mobile Application

```
apps/mobile/
├── app/                                    # Expo Router file-based routing
│   ├── (auth)/                            # Auth screens group
│   │   ├── _layout.tsx                    # Auth layout
│   │   ├── login.tsx                      # Login screen
│   │   │                                  # BE: Keycloak mobile OAuth2 flow
│   │   ├── register.tsx                   # Register screen
│   │   │                                  # BE: users-be/user — POST /v1/users/register
│   │   └── callback.tsx                   # OAuth callback handler
│   │
│   ├── (tabs)/                            # Main tab navigation
│   │   ├── _layout.tsx                    # Tab layout
│   │   ├── index.tsx                      # Home/Dashboard tab
│   │   ├── jobs.tsx                       # Jobs tab
│   │   ├── messages.tsx                   # Messages tab
│   │   ├── notifications.tsx              # Notifications tab
│   │   └── profile.tsx                    # Profile tab
│   │
│   ├── +not-found.tsx                     # 404 screen
│   ├── _layout.tsx                        # Root layout
│   └── index.tsx                          # App entry (redirects to tabs or auth)
│
├── src/
│   ├── components/                        # Mobile-specific components
│   │   ├── Auth/
│   │   │   ├── LoginForm.tsx
│   │   │   ├── RegisterForm.tsx
│   │   │   └── SocialButtons.tsx
│   │   │
│   │   ├── Common/
│   │   │   ├── ErrorBoundary.tsx
│   │   │   ├── Loading.tsx
│   │   │   └── OptimizedFlashList.tsx     # Optimized FlatList wrapper
│   │   │
│   │   └── landing/
│   │       ├── FeaturesMobile.tsx
│   │       └── HeroMobile.tsx
│   │
│   ├── hooks/
│   │   └── useHighFPSAnimation.ts         # 60fps animation hook
│   │
│   ├── lib/
│   │   ├── keycloak-mobile.ts             # Keycloak mobile client
│   │   │                                  # BE: Keycloak OAuth2 PKCE flow
│   │   ├── performance.ts                 # Performance optimization utilities
│   │   ├── utils.ts                       # General utilities
│   │   └── i18n/
│   │       └── index.ts                   # i18n setup
│   │
│   └── features/                          # Feature modules (similar to web)
│       ├── auth/                          # Auth feature
│       ├── profile/                       # Profile feature
│       ├── jobs/                          # Jobs feature
│       ├── proposals/                     # Proposals feature
│       ├── contracts/                     # Contracts feature
│       ├── payments/                      # Payments feature
│       ├── messaging/                     # Messaging feature
│       └── notifications/                 # Notifications feature
│
├── assets/                                # Mobile assets
│   ├── images/
│   ├── fonts/
│   └── icons/
│
├── .env
├── .eslintrc.json
├── app.json                               # Expo configuration
├── babel.config.js
├── global.css                             # NativeWind global styles
├── index.js                               # App entry point
├── metro.config.js                        # Metro bundler configuration
├── package.json
├── tailwind.config.js                     # NativeWind Tailwind config
└── tsconfig.json
```

---

## Packages - Shared Libraries

### UI Package (`packages/ui`)

```
packages/ui/
├── src/
│   ├── components/                        # Cross-platform components
│   │   ├── Button/
│   │   │   ├── Button.web.tsx            # Web-specific implementation
│   │   │   ├── Button.native.tsx         # Native-specific implementation
│   │   │   ├── Button.tsx                # Shared logic/types
│   │   │   └── Button.stories.tsx
│   │   │
│   │   ├── Input/
│   │   │   ├── Input.web.tsx
│   │   │   ├── Input.native.tsx
│   │   │   └── Input.tsx
│   │   │
│   │   ├── Card/
│   │   ├── Modal/
│   │   ├── Dropdown/
│   │   ├── Badge/
│   │   ├── Avatar/
│   │   ├── Skeleton/
│   │   └── ...                           # All UI components
│   │
│   ├── theme/                            # Design tokens
│   │   ├── colors.ts
│   │   ├── spacing.ts
│   │   ├── typography.ts
│   │   ├── shadows.ts
│   │   └── breakpoints.ts
│   │
│   └── index.ts                          # Package exports
│
├── package.json
└── tsconfig.json
```

### Shared Package (`packages/shared`)

```
packages/shared/
├── src/
│   ├── features/
│   │   ├── auth/                         # Shared auth logic
│   │   │   ├── hooks/
│   │   │   │   ├── useAuth.ts           # Platform-agnostic auth hook
│   │   │   │   └── useKeycloak.ts       # Keycloak integration hook
│   │   │   │
│   │   │   ├── stores/
│   │   │   │   └── authStore.ts         # Zustand auth store
│   │   │   │
│   │   │   └── utils/
│   │   │       ├── tokens.ts            # JWT token utilities
│   │   │       └── permissions.ts       # Permission checking utilities
│   │   │
│   │   ├── i18n/                        # Internationalization
│   │   │   ├── locales/                # JSON locale files
│   │   │   │   ├── en.json
│   │   │   │   ├── ar.json
│   │   │   │   ├── zh.json
│   │   │   │   ├── hi.json
│   │   │   │   ├── de.json
│   │   │   │   ├── fr.json
│   │   │   │   ├── tr.json
│   │   │   │   ├── es.json
│   │   │   │   └── ru.json
│   │   │   │
│   │   │   ├── config.ts               # i18n configuration
│   │   │   └── utils.ts                # i18n utilities
│   │   │
│   │   └── api/                         # Shared API utilities
│   │       ├── client.ts               # API client factory
│   │       ├── websocket.ts            # WebSocket client
│   │       └── interceptors.ts         # Request/response interceptors
│   │
│   ├── hooks/                           # Shared hooks
│   │   ├── useDebounce.ts
│   │   ├── useMediaQuery.ts
│   │   ├── useLocalStorage.ts
│   │   └── ...                         # All reusable hooks
│   │
│   ├── utils/                           # Shared utilities
│   │   ├── date.ts
│   │   ├── currency.ts
│   │   ├── string.ts
│   │   ├── validation.ts
│   │   └── ...                         # All utility functions
│   │
│   ├── constants/                       # Shared constants
│   │   ├── api.ts                      # API endpoints
│   │   ├── config.ts                   # Configuration constants
│   │   └── enums.ts                    # Shared enums
│   │
│   └── index.ts                         # Package exports
│
├── package.json
└── tsconfig.json
```

### Types Package (`packages/types`)

```
packages/types/
├── src/
│   ├── api/                             # API types
│   │   ├── response.ts                 # Standard API response types
│   │   ├── pagination.ts               # Pagination types
│   │   └── error.ts                    # Error types
│   │
│   ├── entities/                        # Domain entity types (aligned with BE)
│   │   ├── user.ts                     # User types
│   │   │                               # Corresponds to: users-be/user domain
│   │   │
│   │   ├── profile.ts                  # Profile types
│   │   │                               # Corresponds to: users-be/profile domain
│   │   │
│   │   ├── job.ts                      # Job types
│   │   │                               # Corresponds to: jobs-be/job domain
│   │   │
│   │   ├── proposal.ts                 # Proposal types
│   │   │                               # Corresponds to: proposals-be/proposal domain
│   │   │
│   │   ├── contract.ts                 # Contract types
│   │   │                               # Corresponds to: contracts-be/contract domain
│   │   │
│   │   ├── milestone.ts                # Milestone types
│   │   │                               # Corresponds to: contracts-be/milestones domain
│   │   │
│   │   ├── payment.ts                  # Payment types
│   │   │                               # Corresponds to: financial-be/payment domain
│   │   │
│   │   ├── wallet.ts                   # Wallet types
│   │   │                               # Corresponds to: financial-be/wallet domain
│   │   │
│   │   ├── transaction.ts              # Transaction types
│   │   │                               # Corresponds to: financial-be/transactions domain
│   │   │
│   │   ├── message.ts                  # Message types
│   │   │                               # Corresponds to: communications-be/messages domain
│   │   │
│   │   ├── notification.ts             # Notification types
│   │   │                               # Corresponds to: communications-be/notifications domain
│   │   │
│   │   ├── subscription.ts             # Subscription types
│   │   │                               # Corresponds to: subscriptions-be/subscriptions domain
│   │   │
│   │   └── review.ts                   # Review types
│   │                                   # Corresponds to: reviews-be/reviews domain
│   │
│   ├── forms/                           # Form types
│   │   ├── login.ts
│   │   ├── register.ts
│   │   ├── profile.ts
│   │   ├── job.ts
│   │   ├── proposal.ts
│   │   └── ...                         # All form types
│   │
│   ├── filters/                         # Filter types
│   │   ├── job-filters.ts
│   │   ├── proposal-filters.ts
│   │   ├── contract-filters.ts
│   │   └── ...                         # All filter types
│   │
│   └── index.ts                         # Package exports
│
├── package.json
└── tsconfig.json
```

### Config Package (`packages/config`)

```
packages/config/
├── eslint/
│   ├── next.js                          # Next.js ESLint config
│   ├── react-native.js                  # React Native ESLint config
│   └── base.js                          # Base ESLint config
│
├── typescript/
│   ├── nextjs.json                      # Next.js TS config
│   ├── react-native.json                # React Native TS config
│   └── base.json                        # Base TS config
│
├── tailwind/
│   ├── web.js                           # Web Tailwind config
│   └── native.js                        # NativeWind config
│
├── package.json
└── README.md
```

---

## Infrastructure & Configuration

### Deployment (`deploy/`)

```
deploy/
├── k8s/                                    # Kubernetes manifests
│   ├── web-deployment.yaml                # Web app deployment
│   ├── web-service.yaml                   # Web app service (LoadBalancer)
│   ├── web-ingress.yaml                   # Ingress (HTTPS, domain routing)
│   ├── web-hpa.yaml                       # Horizontal Pod Autoscaler
│   ├── configmap.yaml                     # Environment variables
│   ├── secrets.yaml                       # Secrets (Keycloak, API keys)
│   └── namespace.yaml                     # Namespace definition
│
└── docker/
    ├── web.Dockerfile                     # Web app Dockerfile
    │                                      # Multi-stage build: build -> runtime
    │                                      # Node 20 Alpine, optimized layer caching
    │
    └── mobile-build.Dockerfile            # Mobile build Dockerfile (for CI)
```

### Root Configuration Files

```
skillsier-fe/
├── .env.example                           # Environment variables template
│                                          # NEXT_PUBLIC_API_BASE_URL=https://api.skillsier.com
│                                          # NEXT_PUBLIC_WS_URL=wss://api.skillsier.com/ws
│                                          # NEXT_PUBLIC_KEYCLOAK_URL=https://auth.skillsier.com
│                                          # NEXT_PUBLIC_KEYCLOAK_REALM=skillsier
│                                          # NEXT_PUBLIC_KEYCLOAK_CLIENT_ID=skillsier-web
│                                          # NEXT_PUBLIC_STORAGE_URL=https://storage.skillsier.com
│
├── .eslintrc.json                         # Root ESLint configuration
├── .prettierrc                            # Prettier configuration
├── .gitignore
├── .nvmrc                                 # Node version (v20.x.x)
│
├── package.json                           # Root package
│                                          # Scripts: dev, build, lint, test, clean, format
│                                          # Workspaces: apps/*, packages/*
│
├── pnpm-workspace.yaml                    # pnpm workspace definition
│                                          # packages:
│                                          #   - 'apps/*'
│                                          #   - 'packages/*'
│
├── turbo.json                             # Turborepo pipeline
│                                          # Pipelines: build, dev, lint, test, clean
│                                          # Cache: outputs, inputs, dependencies
│
├── tsconfig.base.json                     # Base TypeScript configuration
│                                          # Path aliases: @skillsier/* packages
│                                          # Strict mode, ES2022 target
│
└── README.md                              # Project documentation
```

---

## State Management & Data Flow Annotations

### TanStack Query Configuration

```
Query Client Default Options:
├── queries:
│   ├── staleTime: 30s (lists), 60s (details)
│   ├── gcTime: 10min
│   ├── retry: 3 (with exponential backoff)
│   ├── refetchOnWindowFocus: true
│   └── refetchOnReconnect: true
│
├── mutations:
│   ├── retry: 1
│   └── onError: global error handler
│
└── Query Key Patterns:
    ├── ['domain', 'action', ...params]
    ├── Example: ['jobs', 'list', { filters }]
    ├── Example: ['jobs', 'detail', job_id]
    ├── Example: ['proposals', 'list', { user_id, status }]
    └── Example: ['wallet', user_id]
```

### Zustand Store Structure

```
Root Store:
├── authSlice:                             # Server state: NEVER duplicated here
│   ├── isAuthenticated: boolean
│   ├── user: User | null                  # Minimal user info (from JWT)
│   └── permissions: string[]
│
├── uiSlice:                               # Client/UI state only
│   ├── theme: 'light' | 'dark'           # Persisted
│   ├── locale: string                     # Persisted
│   ├── sidebarOpen: boolean
│   ├── activeModal: string | null
│   └── toasts: Toast[]
│
├── notificationSlice:                     # Client state only
│   ├── unreadCount: number               # Real-time from WebSocket
│   └── latestNotifications: Notification[] # Last 5 for bell dropdown
│
└── filterSlice:                           # UI state for search/filters
    ├── jobFilters: JobFilters
    ├── proposalFilters: ProposalFilters
    └── contractFilters: ContractFilters
```

### Invalidation Rules

```
Mutation → Query Invalidation Mapping:

# Auth
POST /auth/login → Invalidate: ['auth', 'me']
POST /auth/logout → Invalidate: ['auth', 'me'], clear all queries

# Jobs
POST /jobs → Invalidate: ['jobs', 'list']
PATCH /jobs/{id} → Invalidate: ['jobs', 'detail', id], ['jobs', 'list']
POST /jobs/{id}/publish → Invalidate: ['jobs', 'detail', id], ['jobs', 'list']

# Proposals
POST /proposals → Invalidate: ['proposals', 'list'], ['jobs', 'detail', job_id]
PUT /proposals/{id} → Invalidate: ['proposals', 'detail', id], ['proposals', 'list']
POST /proposals/{id}/withdraw → Invalidate: ['proposals', 'detail', id], ['proposals', 'list']

# Contracts
POST /contracts/{id}/end → Invalidate: ['contracts', 'detail', id], ['contracts', 'list']

# Milestones
POST /milestones → Invalidate: ['milestones', 'list', contract_id], ['contracts', 'detail', contract_id]
POST /milestones/{id}/submit → Invalidate: ['milestones', 'list', contract_id], ['milestones', 'detail', id]
POST /milestones/{id}/approve → Invalidate: ['milestones', 'list', contract_id], ['milestones', 'detail', id], ['contracts', 'detail', contract_id]

# Payments
POST /wallets/{id}/deposit → Invalidate: ['wallet', user_id], ['transactions', user_id]
POST /payouts/request → Invalidate: ['wallet', user_id], ['transactions', user_id]
POST /payment-methods → Invalidate: ['payment-methods', user_id]

# Messaging
POST /messages → Invalidate: ['conversations', user_id], ['messages', conversation_id]
POST /messages/{id}/read → Invalidate: ['conversations', user_id], ['notifications', 'unread-count']

# Notifications
POST /notifications/{id}/read → Invalidate: ['notifications', user_id], ['notifications', 'unread-count']
POST /notifications/read-all → Invalidate: ['notifications', user_id], ['notifications', 'unread-count']

# Reviews
POST /reviews → Invalidate: ['reviews', 'list'], ['reviews', 'stats'], ['profile', reviewee_id]
```

### WebSocket Real-Time Updates

```
WebSocket Connections:
├── wss://api/v1/ws/user/{user_id}
│   ├── Purpose: User-level real-time updates
│   ├── Events:
│   │   ├── notification.new → Update ['notifications', 'unread-count'], toast notification
│   │   ├── message.new → Update ['conversations', user_id], ['messages', conversation_id]
│   │   └── contract.updated → Invalidate ['contracts', 'detail', contract_id]
│   │
│   └── Auto-reconnect: exponential backoff (1s, 2s, 4s, 8s, max 30s)
│
├── wss://api/v1/ws/conversations/{conversation_id}
│   ├── Purpose: Conversation-level real-time updates
│   ├── Events:
│   │   ├── message.sent → Append to ['messages', conversation_id]
│   │   ├── message.read → Update read receipts
│   │   └── typing → Show typing indicator
│   │
│   └── Optimistic updates: message appears immediately before server confirmation
│
└── wss://api/v1/ws/notifications/{user_id}
    ├── Purpose: Notifications stream
    ├── Events:
    │   └── notification.new → Update ['notifications', user_id], unread count
    │
    └── Fallback: polling every 30s if WebSocket fails
```

---

## Backend API Endpoint Summary

### Core Endpoints Reference

```
# Authentication (Keycloak + users-be/auth)
POST   /v1/auth/login                      # Login
POST   /v1/auth/logout                     # Logout
POST   /v1/auth/refresh                    # Refresh token
POST   /v1/auth/forgot-password            # Request password reset
POST   /v1/auth/reset-password             # Reset password
GET    /v1/auth/me                         # Current user

# Users (users-be/user)
POST   /v1/users/register                  # Register user
POST   /v1/users/verify-email              # Verify email
GET    /v1/users/{id}                      # Get user
PATCH  /v1/users/{id}                      # Update user
POST   /v1/users/{id}/deactivate           # Deactivate account

# Profile (users-be/profile)
GET    /v1/users/{id}/profile              # Get profile
PATCH  /v1/users/{id}/profile              # Update profile
GET    /v1/users/{id}/settings             # Get settings

# Skills (users-be/capabilities)
GET    /v1/users/{id}/skills               # List skills
POST   /v1/users/{id}/skills               # Add skill
PUT    /v1/users/{id}/skills/{skill_id}    # Update skill
DELETE /v1/users/{id}/skills/{skill_id}    # Delete skill

# Portfolio (users-be/portfolio)
GET    /v1/users/{id}/portfolio            # List portfolio items
POST   /v1/users/{id}/portfolio            # Create portfolio item
GET    /v1/portfolio/{item_id}             # Get portfolio item
PUT    /v1/portfolio/{item_id}             # Update portfolio item
DELETE /v1/portfolio/{item_id}             # Delete portfolio item

# Jobs (jobs-be/job)
GET    /v1/jobs                            # List jobs (with filters)
GET    /v1/jobs/{job_id}                   # Get job details
POST   /v1/jobs                            # Create job (draft)
PATCH  /v1/jobs/{job_id}                   # Update job
POST   /v1/jobs/{job_id}/publish           # Publish job
DELETE /v1/jobs/{job_id}                   # Delete job

# Search (search-be/search)
POST   /v1/search/jobs                     # Search jobs
POST   /v1/search/freelancers              # Search freelancers
GET    /v1/search/autocomplete             # Autocomplete suggestions
GET    /v1/search/saved                    # Saved searches
POST   /v1/search/saved                    # Save search

# Recommendations (search-be/recommendations)
GET    /v1/recommendations/jobs            # Recommended jobs

# Proposals (proposals-be/proposal)
GET    /v1/proposals                       # List proposals
GET    /v1/proposals/{proposal_id}         # Get proposal
POST   /v1/proposals                       # Submit proposal
PUT    /v1/proposals/{proposal_id}         # Update proposal
POST   /v1/proposals/{proposal_id}/withdraw # Withdraw proposal

# Contracts (contracts-be/contract)
GET    /v1/contracts                       # List contracts
GET    /v1/contracts/{contract_id}         # Get contract
POST   /v1/contracts/{contract_id}/end     # End contract

# Milestones (contracts-be/milestones)
GET    /v1/contracts/{contract_id}/milestones # List milestones
POST   /v1/milestones                      # Create milestone
GET    /v1/milestones/{milestone_id}       # Get milestone
POST   /v1/milestones/{milestone_id}/submit # Submit milestone
POST   /v1/milestones/{milestone_id}/approve # Approve milestone
POST   /v1/milestones/{milestone_id}/reject # Reject milestone

# Wallet (financial-be/wallet)
GET    /v1/wallets/{user_id}               # Get wallet
POST   /v1/wallets/{wallet_id}/deposit     # Deposit funds

# Transactions (financial-be/transactions)
GET    /v1/transactions                    # List transactions

# Payment Methods (financial-be/payment_method)
GET    /v1/payment-methods                 # List payment methods
POST   /v1/payment-methods                 # Add payment method
DELETE /v1/payment-methods/{method_id}     # Delete payment method

# Payouts (financial-be/payout)
POST   /v1/payouts/request                 # Request payout

# Invoices (financial-be/invoice)
GET    /v1/invoices                        # List invoices
GET    /v1/invoices/{invoice_id}           # Get invoice
GET    /v1/invoices/{invoice_id}/pdf       # Download invoice PDF

# Messaging (communications-be/conversations, messages)
GET    /v1/conversations                   # List conversations
GET    /v1/conversations/{conv_id}/messages # List messages
POST   /v1/messages                        # Send message
POST   /v1/messages/{msg_id}/read          # Mark message as read

# Notifications (communications-be/notifications)
GET    /v1/notifications                   # List notifications
GET    /v1/notifications/unread-count      # Get unread count
POST   /v1/notifications/{notif_id}/read   # Mark as read
POST   /v1/notifications/read-all          # Mark all as read

# Subscriptions (subscriptions-be/plans, subscriptions)
GET    /v1/plans                           # List plans
GET    /v1/subscriptions                   # Get user subscription
POST   /v1/subscriptions/{sub_id}/upgrade  # Upgrade subscription

# Entitlements (subscriptions-be/entitlements)
GET    /v1/entitlements                    # Get user entitlements

# Connects (subscriptions-be/connects)
GET    /v1/connects/balance                # Get connects balance
POST   /v1/connects/purchase               # Purchase connects

# Reviews (reviews-be/reviews)
GET    /v1/reviews                         # List reviews
POST   /v1/reviews                         # Submit review
GET    /v1/reviews/stats                   # Review statistics

# Storage (storage-be/uploads)
POST   /v1/storage/upload                  # Upload file (signed URL flow)

# Analytics (analytics-be)
GET    /v1/analytics/overview              # Analytics overview
GET    /v1/analytics/jobs                  # Job analytics
GET    /v1/analytics/spending              # Spending analytics

# Admin (admin-be/*)
GET    /v1/admin/dashboard                 # Admin dashboard
GET    /v1/admin/users                     # User management
POST   /v1/admin/users/{id}/suspend        # Suspend user
POST   /v1/admin/users/{id}/ban            # Ban user
GET    /v1/admin/kyc/cases                 # KYC cases
POST   /v1/admin/kyc/cases/{id}/approve    # Approve KYC
GET    /v1/admin/disputes                  # Dispute cases
POST   /v1/admin/disputes/{id}/resolve     # Resolve dispute
```

---

## Technology Stack Summary

```
Core Technologies:
├── Monorepo: pnpm workspaces + Turborepo
├── Web: Next.js 15 (App Router), React 19
├── Mobile: Expo (React Native), Expo Router
├── Language: TypeScript 5.x
├── Styling: Tailwind CSS (web), NativeWind (mobile)
├── Auth: Keycloak (OAuth2, JWT)
├── State Management:
│   ├── Server State: TanStack Query (React Query v5)
│   └── Client State: Zustand
├── i18n: next-intl (web), i18n-js (mobile)
├── Real-time: WebSocket (native WebSocket API)
└── Testing: Jest, React Testing Library

Build & Deploy:
├── Bundler: Turbopack (Next.js), Metro (React Native)
├── Package Manager: pnpm
├── CI/CD: GitHub Actions
├── Deployment: Kubernetes (K8s)
└── Container: Docker (multi-stage builds)

Code Quality:
├── Linting: ESLint
├── Formatting: Prettier
├── Type Checking: TypeScript strict mode
├── Git Hooks: Husky
└── Pre-commit: lint-staged
```

---

## Performance Optimization Strategy

```
Web Performance:
├── SSR/ISR: Static pages cached at edge, dynamic SSR for personalized content
├── Code Splitting: Route-level automatic, manual for large modals/components
├── Image Optimization: Next.js Image component, WebP/AVIF, lazy loading
├── Font Optimization: next/font with font subsetting
├── Bundle Analysis: @next/bundle-analyzer in CI
├── Prefetching: Automatic link prefetching, manual for critical paths
├── Caching:
│   ├── TanStack Query: staleTime/gcTime per domain
│   ├── HTTP Cache: Cache-Control headers from BE
│   └── CDN: Cloudflare/CloudFront for static assets
├── Compression: Brotli/Gzip at CDN
└── Monitoring: Web Vitals tracking (LCP, FID, CLS, TTFB)

Mobile Performance:
├── FlatList → FlashList: 60fps scrolling for large lists
├── Image Optimization: expo-image with caching, blurhash placeholders
├── Bundle Size: Metro tree-shaking, code splitting
├── Offline Support: React Query persistence for offline-first
├── Navigation: Expo Router file-based, lazy loading
└── Animations: Reanimated 3 for 60fps native animations
```

---

## Accessibility (a11y) Standards

```
WCAG 2.2 AA Compliance:
├── Semantic HTML: Proper heading hierarchy, landmarks, lists
├── Keyboard Navigation: All interactive elements accessible via keyboard
├── Focus Management: Visible focus indicators, focus trapping in modals
├── ARIA: aria-label, aria-describedby, aria-live regions for dynamic content
├── Color Contrast: 4.5:1 for text, 3:1 for large text/UI components
├── Screen Readers: Tested with NVDA (Windows), VoiceOver (Mac/iOS), TalkBack (Android)
├── Motion: prefers-reduced-motion support
└── Forms: Associated labels, error messages linked via aria-describedby
```

---

## Security Measures

```
Security Implementations:
├── Authentication:
│   ├── Keycloak OAuth2 + PKCE (mobile)
│   ├── JWT with HttpOnly cookies (web)
│   ├── Token refresh flow
│   └── Session management
│
├── Authorization:
│   ├── Role-Based Access Control (RBAC)
│   ├── Permission checks at component level
│   └── Route guards for protected pages
│
├── Data Protection:
│   ├── PII redaction in client-side logs
│   ├── Sensitive data not persisted in localStorage
│   ├── Encrypted storage for tokens (mobile Keychain/KeyStore)
│   └── HTTPS-only (enforced via CSP)
│
├── Input Validation:
│   ├── Zod schemas for all forms
│   ├── Server-side validation (BE validates all inputs)
│   └── XSS prevention via React escaping + DOMPurify for rich text
│
├── CSRF Protection:
│   ├── SameSite cookies
│   └── CSRF tokens for mutations
│
└── Headers:
    ├── Content-Security-Policy (CSP)
    ├── X-Frame-Options: DENY
    ├── X-Content-Type-Options: nosniff
    └── Referrer-Policy: strict-origin-when-cross-origin
```

---

## End of Document

**Total Structure Coverage:**
- ✅ Complete folder hierarchy for web & mobile
- ✅ All feature modules with backend API mappings
- ✅ Shared packages (ui, shared, types, config)
- ✅ State management (TanStack Query + Zustand)
- ✅ Real-time (WebSocket) integration
- ✅ Internationalization (i18n) setup
- ✅ Authentication (Keycloak) integration
- ✅ File upload (storage-be) flows
- ✅ Query key patterns & invalidation rules
- ✅ Performance optimization strategies
- ✅ Accessibility (a11y) standards
- ✅ Security measures
- ✅ Deployment configuration (K8s, Docker)

**NO CODE IMPLEMENTATIONS INCLUDED** — Structure & mappings only, per requirements.
