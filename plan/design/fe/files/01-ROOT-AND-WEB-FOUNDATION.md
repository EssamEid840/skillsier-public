# Skillsier Frontend - Complete Folder Structure
## Part 1: Root Structure & Web App Foundation

> **CRITICAL**: This document contains ONLY the folder structure, filenames, and inline backend API mappings.  
> **NO CODE IMPLEMENTATIONS** are included per the strict output policy.

---

## Root Structure

```
skillsier-fe/
│
├── .husky/                                   # Git hooks for quality gates
│   ├── pre-commit                           # Runs linting, type checking
│   │                                        # Validates no console.log in prod code
│   │                                        # Ensures all tests pass
│   └── pre-push                             # Runs full test suite before push
│                                            # Bundle size check
│
├── .vscode/                                  # VS Code workspace configuration
│   ├── extensions.json                      # Recommended extensions list
│   │                                        # - ESLint, Prettier, Tailwind IntelliSense
│   │                                        # - i18n Ally, Error Lens
│   ├── launch.json                          # Debug configurations
│   │                                        # - Next.js debug config
│   │                                        # - Expo debug config
│   │                                        # - Jest test debug config
│   └── settings.json                        # Workspace settings
│                                            # - Format on save enabled
│                                            # - Auto import organization
│                                            # - Tailwind CSS IntelliSense config
│
├── .github/                                  # GitHub workflows
│   └── workflows/
│       ├── ci.yml                           # Continuous Integration
│       │                                    # - Lint, type-check, test
│       │                                    # - Build verification
│       │                                    # - Bundle size check
│       ├── cd-web.yml                       # Web deployment
│       │                                    # - Build Next.js production
│       │                                    # - Deploy to K8s
│       ├── cd-mobile.yml                    # Mobile deployment
│       │                                    # - EAS build
│       │                                    # - Submit to stores
│       └── dependabot.yml                   # Automated dependency updates
│
├── apps/                                     # Application workspaces
│   ├── web/                                 # Next.js web application
│   └── mobile/                              # React Native/Expo application
│
├── packages/                                 # Shared libraries
│   ├── ui/                                  # Cross-platform component library
│   ├── shared/                              # Business logic, hooks, utilities
│   ├── types/                               # TypeScript type definitions
│   └── config/                              # Shared configurations
│
├── deploy/                                   # Deployment configurations
│   ├── k8s/                                 # Kubernetes manifests
│   │   ├── web/
│   │   │   ├── deployment.yaml              # Web app deployment
│   │   │   ├── service.yaml                 # ClusterIP service
│   │   │   ├── ingress.yaml                 # NGINX ingress
│   │   │   ├── hpa.yaml                     # Horizontal Pod Autoscaler
│   │   │   └── configmap.yaml               # Environment config
│   │   ├── mobile-api/                      # Mobile API gateway (if separate)
│   │   └── monitoring/
│   │       ├── prometheus.yaml              # Metrics scraping
│   │       └── grafana-dashboard.yaml       # Dashboards
│   │
│   └── docker/
│       ├── web.Dockerfile                   # Multi-stage Next.js build
│       ├── web.dockerignore
│       ├── mobile-build.Dockerfile          # EAS build container
│       └── nginx.conf                       # NGINX config for static serving
│
├── docs/                                     # Documentation
│   ├── README.md                            # Project overview
│   ├── ARCHITECTURE.md                      # System architecture
│   ├── SETUP.md                             # Development setup guide
│   ├── PERFORMANCE.md                       # Performance optimization guide
│   ├── CONTRIBUTING.md                      # Contribution guidelines
│   ├── MICROSERVICES_MAPPING.md             # BE microservices integration
│   ├── STATE_MANAGEMENT.md                  # TanStack Query + Zustand patterns
│   ├── TESTING.md                           # Testing strategy
│   └── DEPLOYMENT.md                        # Deployment procedures
│
├── scripts/                                  # Automation scripts
│   ├── setup.sh                             # Initial project setup
│   ├── dev.sh                               # Start all dev servers
│   ├── build-all.sh                         # Build all apps
│   ├── test-all.sh                          # Run all tests
│   ├── clean.sh                             # Clean build artifacts
│   ├── generate-types.sh                    # Generate types from OpenAPI
│   ├── check-bundle-size.sh                 # Bundle size analysis
│   └── db-seed.sh                           # Seed local dev data
│
├── .env.example                              # Environment variables template
├── .env.local.example                        # Local overrides template
├── .gitignore                               # Git ignore patterns
├── .nvmrc                                    # Node.js version (v20.x)
├── .prettierrc                              # Prettier configuration
├── .prettierignore                          # Prettier ignore patterns
├── .eslintrc.json                           # Root ESLint config
├── package.json                              # Root package (workspace manager)
│                                            # Scripts: dev, build, test, lint, type-check
├── pnpm-workspace.yaml                       # pnpm workspace configuration
├── pnpm-lock.yaml                            # Locked dependencies
├── turbo.json                                # Turborepo pipeline configuration
│                                            # Build cache, task dependencies
├── tsconfig.base.json                        # Base TypeScript configuration
│                                            # Shared compiler options
├── tsconfig.json                             # Root TypeScript config
├── jest.config.js                           # Root Jest configuration
└── README.md                                # Root README
```

---

## Apps/Web - Main Structure

```
apps/web/
│
├── public/                                   # Static assets
│   ├── images/
│   │   ├── logo.svg                         # Skillsier logo
│   │   ├── logo-dark.svg                    # Dark mode logo
│   │   ├── favicon.ico
│   │   ├── hero-bg.webp                     # Hero background
│   │   ├── hero-bg-mobile.webp              # Mobile hero
│   │   ├── placeholder-avatar.png           # Default avatar
│   │   ├── placeholder-company.png          # Default company logo
│   │   ├── payment-providers/               # Payment method icons
│   │   │   ├── visa.svg
│   │   │   ├── mastercard.svg
│   │   │   ├── paypal.svg
│   │   │   └── stripe.svg
│   │   └── social/                          # Social login icons
│   │       ├── google.svg
│   │       ├── github.svg
│   │       └── linkedin.svg
│   │
│   ├── fonts/                               # Web fonts
│   │   ├── inter-var.woff2                  # Variable font
│   │   ├── inter-var-latin.woff2            # Latin subset
│   │   └── noto-sans-arabic.woff2           # Arabic font
│   │
│   ├── locales/                             # Public locale files (deprecated - moved to packages/shared)
│   │
│   ├── animations/                          # Lottie/animation files
│   │   ├── loading.json
│   │   ├── success.json
│   │   ├── error.json
│   │   └── empty-state.json
│   │
│   ├── robots.txt                           # Search engine directives
│   ├── sitemap.xml                          # Static sitemap
│   ├── sitemap-dynamic.xml                  # Dynamic routes sitemap
│   ├── manifest.json                        # PWA manifest
│   ├── sw.js                                # Service worker (if using PWA)
│   └── browserconfig.xml                    # Windows tile config
│
├── src/
│   ├── app/                                 # Next.js 15 App Router
│   ├── features/                            # Feature modules (detailed below)
│   ├── components/                          # Shared web-specific components
│   ├── lib/                                 # Web-specific utilities
│   ├── hooks/                               # Web-specific hooks
│   ├── styles/                              # Global styles
│   ├── types/                               # Web-specific types
│   └── middleware.ts                        # Next.js middleware
│                                            # - Auth validation (Keycloak JWT)
│                                            # - i18n routing
│                                            # - Rate limiting (client-side)
│                                            # - Request tracing
│                                            # BE: Keycloak token validation
│
├── .env.local                               # Local environment variables
├── .env.development                         # Development environment
├── .env.production                          # Production environment
├── .eslintrc.json                           # Web-specific ESLint rules
├── next.config.js                           # Next.js configuration
│                                            # - i18n config
│                                            # - Image optimization
│                                            # - Bundle analyzer
│                                            # - Security headers
│                                            # - Rewrites/redirects
├── postcss.config.js                        # PostCSS configuration
├── tailwind.config.js                       # Tailwind configuration
│                                            # - Design tokens
│                                            # - Custom utilities
│                                            # - Plugin configuration
├── tsconfig.json                            # Web TypeScript config
│                                            # Extends tsconfig.base.json
├── jest.config.js                           # Web Jest configuration
├── package.json                             # Web dependencies
└── README.md                                # Web app documentation
```

---

## App Router Structure (apps/web/src/app)

```
src/app/
│
├── [locale]/                                # Internationalized routing (en, ar, zh, hi, de, fr, tr, es, ru)
│   │
│   ├── (public)/                           # Public pages (no auth required)
│   │   │
│   │   ├── page.tsx                        # Homepage / Landing
│   │   │                                   # - Hero section
│   │   │                                   # - Features overview
│   │   │                                   # - Stats showcase
│   │   │                                   # - Testimonials
│   │   │                                   # - CTA sections
│   │   │                                   # BE: None (static/cached content)
│   │   │
│   │   ├── about/
│   │   │   └── page.tsx                    # About page
│   │   │                                   # - Company story
│   │   │                                   # - Team showcase
│   │   │                                   # - Mission & values
│   │   │                                   # BE: None (static)
│   │   │
│   │   ├── how-it-works/
│   │   │   ├── page.tsx                    # How it works overview
│   │   │   ├── freelancers/
│   │   │   │   └── page.tsx                # For freelancers
│   │   │   │                               # - Getting started
│   │   │   │                               # - Finding jobs
│   │   │   │                               # - Earning money
│   │   │   │                               # BE: None (static)
│   │   │   └── clients/
│   │   │       └── page.tsx                # For clients
│   │   │                                   # - Posting jobs
│   │   │                                   # - Hiring process
│   │   │                                   # - Payment & escrow
│   │   │                                   # BE: None (static)
│   │   │
│   │   ├── pricing/
│   │   │   └── page.tsx                    # Pricing page
│   │   │                                   # - Plan comparison
│   │   │                                   # - Feature matrix
│   │   │                                   # - FAQ
│   │   │                                   # BE: subscriptions-be/plans
│   │   │                                   # GET /v1/plans?public=true
│   │   │                                   # Returns: plans with public pricing
│   │   │
│   │   ├── trust-safety/
│   │   │   └── page.tsx                    # Trust & Safety
│   │   │                                   # - Security measures
│   │   │                                   # - Payment protection
│   │   │                                   # - Dispute resolution
│   │   │                                   # BE: None (static)
│   │   │
│   │   ├── contact/
│   │   │   └── page.tsx                    # Contact form
│   │   │                                   # - Support inquiries
│   │   │                                   # - Partnership requests
│   │   │                                   # - Media contacts
│   │   │                                   # BE: communications-be/messages
│   │   │                                   # POST /v1/contact
│   │   │                                   # Body: { name, email, subject, message, type }
│   │   │
│   │   ├── careers/
│   │   │   ├── page.tsx                    # Careers overview
│   │   │   └── [slug]/
│   │   │       └── page.tsx                # Individual job posting
│   │   │                                   # BE: None (static or CMS)
│   │   │
│   │   ├── blog/
│   │   │   ├── page.tsx                    # Blog listing
│   │   │   ├── [slug]/
│   │   │   │   └── page.tsx                # Blog post detail
│   │   │   └── category/
│   │   │       └── [category]/
│   │   │           └── page.tsx            # Category listing
│   │   │                                   # BE: CMS or separate service
│   │   │
│   │   ├── help/
│   │   │   ├── page.tsx                    # Help center home
│   │   │   ├── [category]/
│   │   │   │   └── page.tsx                # Help category
│   │   │   └── article/
│   │   │       └── [slug]/
│   │   │           └── page.tsx            # Help article
│   │   │                                   # BE: communications-be/knowledge-base
│   │   │                                   # GET /v1/kb/articles
│   │   │                                   # GET /v1/kb/articles/{slug}
│   │   │
│   │   ├── legal/
│   │   │   ├── terms/
│   │   │   │   └── page.tsx                # Terms of Service
│   │   │   ├── privacy/
│   │   │   │   └── page.tsx                # Privacy Policy
│   │   │   ├── cookies/
│   │   │   │   └── page.tsx                # Cookie Policy
│   │   │   └── compliance/
│   │   │       └── page.tsx                # Compliance information
│   │   │                                   # BE: None (static/versioned)
│   │   │
│   │   └── layout.tsx                      # Public pages layout
│                                            # - Public header (no auth UI)
│                                            # - Footer
│                                            # - Language switcher
│
│   ├── (auth)/                             # Authentication pages (no dashboard layout)
│   │   │
│   │   ├── login/
│   │   │   └── page.tsx                    # Login page
│   │   │                                   # - Email/password form
│   │   │                                   # - Social login buttons (Google, GitHub, LinkedIn)
│   │   │                                   # - Remember me option
│   │   │                                   # - Forgot password link
│   │   │                                   # - Register link
│   │   │                                   # BE: Keycloak OAuth2 flow
│   │   │                                   # POST /v1/auth/login (users-be)
│   │   │                                   # Returns: JWT access_token, refresh_token
│   │   │                                   # Social: OAuth2 redirect to Keycloak
│   │   │
│   │   ├── register/
│   │   │   ├── page.tsx                    # Registration page
│   │   │   │                               # - User type selection (Freelancer/Client)
│   │   │   │                               # - Email, password, name fields
│   │   │   │                               # - Terms acceptance
│   │   │   │                               # - Social registration
│   │   │   │                               # BE: users-be/user
│   │   │   │                               # POST /v1/users/register
│   │   │   │                               # Body: { email, password, first_name, last_name, user_type, terms_accepted }
│   │   │   │                               # Returns: user_id, verification_email_sent
│   │   │   │                               # Publishes: UserCreated event
│   │   │   │
│   │   │   └── verification/
│   │   │       └── page.tsx                # Email verification callback
│   │   │                                   # - Verify token from email link
│   │   │                                   # - Success/error messaging
│   │   │                                   # - Auto-redirect to onboarding
│   │   │                                   # BE: users-be/user
│   │   │                                   # POST /v1/users/verify-email
│   │   │                                   # Body: { token }
│   │   │                                   # Returns: verification_status
│   │   │                                   # Publishes: UserVerified event
│   │   │
│   │   ├── forgot-password/
│   │   │   └── page.tsx                    # Password reset request
│   │   │                                   # - Email input
│   │   │                                   # - Send reset link
│   │   │                                   # BE: users-be/security/recovery
│   │   │                                   # POST /v1/auth/forgot-password
│   │   │                                   # Body: { email }
│   │   │                                   # Returns: reset_email_sent
│   │   │
│   │   ├── reset-password/
│   │   │   └── page.tsx                    # Password reset form
│   │   │                                   # - New password input
│   │   │                                   # - Confirm password
│   │   │                                   # - Token validation
│   │   │                                   # BE: users-be/security/recovery
│   │   │                                   # POST /v1/auth/reset-password
│   │   │                                   # Body: { token, new_password }
│   │   │                                   # Returns: password_reset_success
│   │   │
│   │   ├── callback/
│   │   │   └── page.tsx                    # OAuth callback handler
│   │   │                                   # - Keycloak callback
│   │   │                                   # - Google OAuth callback
│   │   │                                   # - GitHub OAuth callback
│   │   │                                   # - LinkedIn OAuth callback
│   │   │                                   # - Token exchange
│   │   │                                   # BE: Keycloak token exchange
│   │   │                                   # POST /oauth2/token (Keycloak)
│   │   │                                   # Returns: access_token, refresh_token
│   │   │
│   │   ├── mfa/
│   │   │   ├── setup/
│   │   │   │   └── page.tsx                # MFA setup
│   │   │   │                               # - QR code display
│   │   │   │                               # - Backup codes
│   │   │   │                               # - Verify setup
│   │   │   │                               # BE: users-be/security/mfa
│   │   │   │                               # POST /v1/auth/mfa/setup
│   │   │   │                               # GET /v1/auth/mfa/qrcode
│   │   │   │                               # POST /v1/auth/mfa/verify-setup
│   │   │   │
│   │   │   └── verify/
│   │   │       └── page.tsx                # MFA verification
│   │   │                                   # - OTP input
│   │   │                                   # - Backup code option
│   │   │                                   # - Trust device option
│   │   │                                   # BE: users-be/security/mfa
│   │   │                                   # POST /v1/auth/mfa/verify
│   │   │                                   # Body: { code, trust_device }
│   │   │
│   │   └── layout.tsx                      # Auth pages layout
│                                            # - Minimal layout (logo + form)
│                                            # - No header/footer
│                                            # - Language switcher
│
│   ├── (onboarding)/                       # Onboarding flow (post-registration)
│   │   │
│   │   ├── welcome/
│   │   │   └── page.tsx                    # Welcome message
│   │   │                                   # - User type confirmation
│   │   │                                   # - Next steps preview
│   │   │                                   # BE: users-be/user
│   │   │                                   # GET /v1/users/me
│   │   │
│   │   ├── freelancer/                     # Freelancer onboarding
│   │   │   ├── profile/
│   │   │   │   └── page.tsx                # Basic profile
│   │   │   │                               # - Professional title
│   │   │   │                               # - Bio
│   │   │   │                               # - Location
│   │   │   │                               # - Profile photo
│   │   │   │                               # BE: users-be/profile
│   │   │   │                               # PATCH /v1/users/{id}/profile
│   │   │   │
│   │   │   ├── skills/
│   │   │   │   └── page.tsx                # Skills selection
│   │   │   │                               # - Skill search & autocomplete
│   │   │   │                               # - Skill level (Beginner/Intermediate/Expert)
│   │   │   │                               # - Primary skills (minimum 3)
│   │   │   │                               # BE: users-be/capabilities
│   │   │   │                               # POST /v1/users/{id}/skills
│   │   │   │                               # Body: { skills: [{ skill_id, level }] }
│   │   │   │
│   │   │   ├── experience/
│   │   │   │   └── page.tsx                # Work experience
│   │   │   │                               # - Add previous positions
│   │   │   │                               # - Company, role, dates, description
│   │   │   │                               # BE: users-be/experience
│   │   │   │                               # POST /v1/users/{id}/experience
│   │   │   │
│   │   │   ├── portfolio/
│   │   │   │   └── page.tsx                # Portfolio items
│   │   │   │                               # - Upload work samples
│   │   │   │                               # - Project descriptions
│   │   │   │                               # - Links to live work
│   │   │   │                               # BE: users-be/portfolio
│   │   │   │                               # POST /v1/users/{id}/portfolio
│   │   │   │                               # BE: storage-be/uploads
│   │   │   │                               # POST /v1/storage/upload (signed URL)
│   │   │   │
│   │   │   ├── rates/
│   │   │   │   └── page.tsx                # Rate setting
│   │   │   │                               # - Hourly rate
│   │   │   │                               # - Preferred project budget range
│   │   │   │                               # - Currency
│   │   │   │                               # BE: users-be/freelancer
│   │   │   │                               # PATCH /v1/users/{id}/rates
│   │   │   │
│   │   │   └── preferences/
│   │   │       └── page.tsx                # Job preferences
│   │   │                                   # - Job categories
│   │   │                                   # - Work availability
│   │   │                                   # - Preferred job types
│   │   │                                   # - Notification settings
│   │   │                                   # BE: users-be/preferences
│   │   │                                   # POST /v1/users/{id}/preferences
│   │   │                                   # Publishes: FreelancerProfileCompleted event
│   │   │
│   │   ├── client/                         # Client onboarding
│   │   │   ├── company/
│   │   │   │   └── page.tsx                # Company info
│   │   │   │                               # - Company name
│   │   │   │                               # - Industry
│   │   │   │                               # - Company size
│   │   │   │                               # - Website
│   │   │   │                               # - Logo upload
│   │   │   │                               # BE: users-be/organization
│   │   │   │                               # POST /v1/organizations
│   │   │   │                               # BE: storage-be/uploads
│   │   │   │
│   │   │   ├── billing/
│   │   │   │   └── page.tsx                # Billing setup
│   │   │   │                               # - Billing address
│   │   │   │                               # - Tax information
│   │   │   │                               # - VAT/GST number
│   │   │   │                               # - Payment method
│   │   │   │                               # BE: financial-be/billing_profile
│   │   │   │                               # POST /v1/billing-profiles
│   │   │   │                               # BE: financial-be/payment_method
│   │   │   │                               # POST /v1/payment-methods
│   │   │   │
│   │   │   ├── verification/
│   │   │   │   └── page.tsx                # Business verification
│   │   │   │                               # - Upload business documents
│   │   │   │                               # - Company registration
│   │   │   │                               # - Tax documents
│   │   │   │                               # BE: admin-be/business_verification
│   │   │   │                               # POST /v1/admin/business-verification
│   │   │   │                               # BE: storage-be/uploads
│   │   │   │
│   │   │   ├── team/
│   │   │   │   └── page.tsx                # Team setup (optional)
│   │   │   │                               # - Invite team members
│   │   │   │                               # - Assign roles
│   │   │   │                               # BE: users-be/team
│   │   │   │                               # POST /v1/organizations/{id}/members
│   │   │   │                               # BE: communications-be
│   │   │   │                               # POST /v1/invitations/team
│   │   │   │
│   │   │   └── preferences/
│   │   │       └── page.tsx                # Hiring preferences
│   │   │                                   # - Typical project types
│   │   │                                   # - Budget ranges
│   │   │                                   # - Notification settings
│   │   │                                   # BE: users-be/preferences
│   │   │                                   # POST /v1/users/{id}/preferences
│   │   │                                   # Publishes: ClientProfileCompleted event
│   │   │
│   │   └── layout.tsx                      # Onboarding layout
│                                            # - Progress indicator
│                                            # - Skip options
│                                            # - Help sidebar
│
│   └── (dashboard)/                        # Main authenticated dashboard
│       │
│       ├── dashboard/                      # Dashboard home (role-based view)
│       │   └── page.tsx                    # Main dashboard
│       │                                   # Freelancer view:
│       │                                   # - Active proposals
│       │                                   # - Active contracts
│       │                                   # - Earnings overview
│       │                                   # - Job recommendations
│       │                                   # - Profile completion score
│       │                                   # Client view:
│       │                                   # - Active jobs
│       │                                   # - Active contracts
│       │                                   # - Spending overview
│       │                                   # - Recent proposals
│       │                                   # - Talent recommendations
│       │                                   # BE: users-be/user
│       │                                   # GET /v1/users/me
│       │                                   # BE: Multiple services for dashboard data
│       │                                   # GET /v1/analytics/dashboard (analytics service)
│       │                                   # GET /v1/jobs/my-jobs (jobs-be)
│       │                                   # GET /v1/proposals/my-proposals (proposals-be)
│       │                                   # GET /v1/contracts/active (contracts-be)
│
│       ├── profile/                        # Current user profile management
│       │   # (Detailed structure in next part)
│
│       ├── jobs/                           # Jobs management
│       │   # (Detailed structure in next part)
│
│       ├── proposals/                      # Proposals management
│       │   # (Detailed structure in next part)
│
│       ├── contracts/                      # Contracts management
│       │   # (Detailed structure in next part)
│
│       ├── messages/                       # Messaging
│       │   # (Detailed structure in next part)
│
│       ├── financials/                     # Financial management
│       │   # (Detailed structure in next part)
│
│       ├── reviews/                        # Reviews & ratings
│       │   # (Detailed structure in next part)
│
│       ├── settings/                       # User settings
│       │   # (Detailed structure in next part)
│
│       ├── subscription/                   # Subscription management
│       │   # (Detailed structure in next part)
│
│       ├── admin/                          # Admin panel (SUPER_ADMIN only)
│       │   # (Detailed structure in next part)
│
│       └── layout.tsx                      # Dashboard layout
│                                            # - Authenticated header with user menu
│                                            # - Sidebar navigation (responsive)
│                                            # - Notification bell with unread count
│                                            # - Messages indicator
│                                            # - Breadcrumbs
│                                            # - Footer
│                                            # BE: communications-be
│                                            # GET /v1/notifications/unread-count
│                                            # WebSocket: ws://communications-be/v1/realtime
│
├── api/                                    # API routes (if needed for BE proxy)
│   ├── auth/
│   │   └── [...nextauth]/
│   │       └── route.ts                    # NextAuth.js API routes (if using)
│   ├── webhooks/
│   │   ├── stripe/
│   │   │   └── route.ts                    # Stripe webhook handler
│   │   └── keycloak/
│   │       └── route.ts                    # Keycloak event webhook
│   └── upload/
│       └── presign/
│           └── route.ts                    # Generate signed upload URLs
│                                            # BE: storage-be
│                                            # POST /v1/storage/presign
│
├── globals.css                             # Global styles
│                                            # - Tailwind base, components, utilities
│                                            # - CSS custom properties for theming
│                                            # - Dark mode variables
│
├── layout.tsx                              # Root layout
│                                            # - HTML lang attribute (i18n)
│                                            # - Head meta tags
│                                            # - Body layout
│                                            # - Font loading
│
├── page.tsx                                # Root page (redirects to /[locale])
│
├── error.tsx                               # Global error boundary
├── not-found.tsx                           # 404 page
└── loading.tsx                             # Global loading state
```

---

## State Management Strategy

### TanStack Query (Server State)

**Configuration** (apps/web/src/lib/react-query.ts):
```typescript
// Query client configuration
// - staleTime defaults by domain
// - gcTime defaults
// - retry with exponential backoff + jitter
// - error handling
// BE: All microservices
```

**Query Key Patterns**:
```typescript
// Users
['users', 'me']                              // Current user
['users', 'detail', userId]                  // User profile
['users', 'skills', userId]                  // User skills
['users', 'portfolio', userId]               // User portfolio

// Jobs
['jobs', 'list', filters]                    // Job list (paginated)
['jobs', 'detail', jobId]                    // Job detail
['jobs', 'my-jobs', filters]                 // Current user's jobs
['jobs', 'recommendations']                  // Job recommendations

// Proposals
['proposals', 'list', filters]               // Proposal list
['proposals', 'detail', proposalId]          // Proposal detail
['proposals', 'my-proposals', filters]       // Current user's proposals
['proposals', 'job', jobId]                  // Proposals for a job

// Contracts
['contracts', 'list', filters]               // Contract list
['contracts', 'detail', contractId]          // Contract detail
['contracts', 'active']                      // Active contracts
['contracts', 'milestones', contractId]      // Contract milestones

// Messages
['conversations', 'list']                    // Conversation list
['conversations', 'detail', conversationId]  // Conversation detail
['messages', conversationId]                 // Messages in conversation

// Notifications
['notifications', 'list']                    // Notification list
['notifications', 'unread-count']            // Unread count

// Financial
['wallet', 'balance']                        // Wallet balance
['transactions', 'list', filters]            // Transaction history
['invoices', 'list', filters]                // Invoice list

// Subscriptions
['subscription', 'current']                  // Current subscription
['plans', 'list']                            // Available plans
['connects', 'balance']                      // Connects balance

// Reviews
['reviews', 'list', userId]                  // User reviews
['reviews', 'stats', userId]                 // Review statistics
```

### Zustand (Client State)

**Store Structure**:
```typescript
// Auth store (packages/shared/src/stores/auth-store.ts)
// - user, isAuthenticated, login, logout
// - Token management
// - Session persistence

// UI store (apps/web/src/stores/ui-store.ts)
// - Theme (light/dark/system)
// - Sidebar collapsed state
// - Modal stack
// - Toast notifications

// Search store (apps/web/src/stores/search-store.ts)
// - Recent searches
// - Search filters
// - Search history

// Draft store (apps/web/src/stores/draft-store.ts)
// - Job drafts (autosave)
// - Proposal drafts (autosave)
// - Message drafts
```

---

**End of Part 1**

**Continue to Part 2 for:**
- Profile Management (Freelancer & Client)
- Jobs Module (Post, Browse, Manage)
- Proposals & Bidding
- Search & Discovery
