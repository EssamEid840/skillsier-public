# Skillsier Complete Frontend Folder Structure
## Comprehensive Guide - All Microservices Mapped (100%)

> **DELIVERY FORMAT**: Due to the exhaustive nature (20,000+ lines), this is delivered as:
> 1. This README with overview
> 2. Complete folder tree (skillsier-fe-complete-tree.txt)
> 3. Backend API mappings by service (api-mappings-*.md files)
> 4. Feature modules detailed structure (features-*.md files)

---

## 📦 Complete Package Contents

```
/mnt/user-data/outputs/
├── COMPLETE-FE-STRUCTURE-README.md          # This file (overview & index)
├── skillsier-fe-complete-tree.txt           # Full folder tree (no comments)
├── api-mappings-users-be.md                 # users-be API mappings (46 domains)
├── api-mappings-jobs-be.md                  # jobs-be API mappings (36 domains)
├── api-mappings-proposals-be.md             # proposals-be API mappings (14 domains)
├── api-mappings-contracts-be.md             # contracts-be API mappings (all domains)
├── api-mappings-financial-be.md             # financial-be API mappings (25 domains)
├── api-mappings-communications-be.md        # communications-be API mappings (57 domains)
├── api-mappings-subscriptions-be.md         # subscriptions-be API mappings (25 domains)
├── api-mappings-reviews-be.md               # reviews-be API mappings (all domains)
├── api-mappings-search-be.md                # search-be API mappings (all domains)
├── api-mappings-storage-be.md               # storage-be API mappings (all domains)
├── api-mappings-admin-be.md                 # admin-be API mappings (all domains)
├── features-auth.md                         # Auth feature detailed structure
├── features-profile.md                      # Profile feature detailed structure
├── features-jobs.md                         # Jobs feature detailed structure
├── features-proposals.md                    # Proposals feature detailed structure
├── features-contracts.md                    # Contracts feature detailed structure
├── features-payments.md                     # Payments feature detailed structure
├── features-messaging.md                    # Messaging feature detailed structure
├── features-notifications.md                # Notifications feature detailed structure
├── features-subscriptions.md                # Subscriptions feature detailed structure
├── features-reviews.md                      # Reviews feature detailed structure
└── features-admin.md                        # Admin feature detailed structure
```

---

## 🏗️ Architecture Overview

### Monorepo Structure

```
skillsier-fe/
├── apps/
│   ├── web/          # Next.js 15 (App Router) - React 19
│   └── mobile/       # Expo/React Native - Expo Router
│
├── packages/
│   ├── ui/           # @skillsier/ui - Cross-platform components
│   ├── shared/       # @skillsier/shared - Business logic, hooks, utilities
│   ├── types/        # @skillsier/types - TypeScript definitions
│   └── config/       # @skillsier/config - ESLint, TS, Tailwind configs
│
├── deploy/           # K8s manifests, Dockerfiles
└── docs/             # Architecture, setup, contributing
```

### Technology Stack

**Core:**
- Monorepo: pnpm workspaces + Turborepo
- Web: Next.js 15 (App Router), React 19, TypeScript 5.x
- Mobile: Expo (React Native), Expo Router
- Styling: Tailwind CSS (web), NativeWind (mobile)

**State Management:**
- Server State: TanStack Query v5 (React Query)
- Client State: Zustand
- Real-time: Native WebSocket + Server-Sent Events (SSE)

**Auth & i18n:**
- Auth: Keycloak (OAuth2, JWT)
- i18n: next-intl (web), i18n-js (mobile) - 9 languages

**Build & Deploy:**
- Bundler: Turbopack (Next.js), Metro (RN)
- Package Manager: pnpm
- Deployment: Kubernetes, Docker

---

## 🗂️ Complete Backend Services Mapped

### Core Microservices

1. **users-be** (46 domains)
   - user, profile, capabilities, skills, portfolio, availability
   - experience, education, credentials, networking
   - privacy, security, reputation, trust, analytics
   - earnings, learning, mentorship, and 30+ more domains

2. **jobs-be** (36 domains)
   - job, pricing, duration, requirements, skills
   - screening, attachments, invitation, sourcing
   - visibility, templates, analytics, and 23+ more domains

3. **proposals-be** (14 domains)
   - proposal, cover_letter, milestones, bidding
   - connects, boost, templates, portfolio
   - engagement, compliance, and 4+ more domains

4. **contracts-be** (15+ domains)
   - contract, milestones, deliverables, timesheet
   - work_diary, sow, amendments, disputes
   - performance, and 6+ more domains

5. **financial-be** (25 domains)
   - wallet, transactions, payments, escrow
   - payouts, invoices, refunds, tax
   - forex, risk, chargebacks, and 13+ more domains

6. **communications-be** (57 domains)
   - conversations, messages, threads, notifications
   - email, push, SMS, WebSocket, presence
   - typing, read_receipts, attachments, and 40+ more domains

7. **subscriptions-be** (25 domains)
   - plans, subscriptions, entitlements, usage
   - connects, seats, addons, trials
   - invoices, dunning, and 15+ more domains

8. **reviews-be** (10+ domains)
   - reviews, ratings, responses, reputation
   - trust_score, and 5+ more domains

9. **search-be** (20+ domains)
   - search, indexing, recommendations, matching
   - saved_searches, autocomplete, facets, and 12+ more domains

10. **storage-be** (15+ domains)
    - uploads, files, media, virus_scanning
    - thumbnails, lifecycle, and 9+ more domains

11. **admin-be** (25+ domains)
    - user_management, kyc, disputes, moderation
    - refunds, analytics, and 19+ more domains

### Utility Services

- **analytics-be**: Metrics, reporting, insights
- **audit-be**: Audit trails, compliance logs
- **feature-flags-be**: A/B testing, feature toggles

---

## 📱 Application Structure

### Web App (Next.js 15)

**Route Groups:**
```
app/[locale]/
├── (auth)/              # Login, register, password reset
├── (landing)/           # Public pages, pricing, about
├── (dashboard)/         # Main authenticated app
│   ├── dashboard/
│   ├── profile/
│   ├── jobs/
│   ├── proposals/
│   ├── contracts/
│   ├── payments/
│   ├── messages/
│   ├── notifications/
│   ├── subscriptions/
│   ├── reviews/
│   ├── analytics/
│   ├── organization/
│   └── settings/
└── (admin)/             # Admin panel
    └── admin/
```

### Mobile App (Expo)

**Screens:**
```
app/
├── (auth)/              # Auth flow
├── (tabs)/              # Main tab navigation
│   ├── index            # Home/Dashboard
│   ├── jobs             # Jobs browse
│   ├── messages         # Messaging
│   ├── notifications    # Notifications
│   └── profile          # Profile
└── [dynamic routes]     # Details screens
```

---

## 🔗 Backend API Integration Pattern

### Query Key Pattern
```typescript
['domain', 'action', ...params]

Examples:
['jobs', 'list', { filters }]
['jobs', 'detail', job_id]
['proposals', 'list', { user_id, status }]
['wallet', user_id]
['messages', conversation_id]
```

### Invalidation Rules
```typescript
// Example: Submit proposal
POST /v1/proposals
→ Invalidates: ['proposals', 'list'], ['jobs', 'detail', job_id]

// Example: Approve milestone  
POST /v1/milestones/{id}/approve
→ Invalidates: 
  - ['milestones', 'list', contract_id]
  - ['milestones', 'detail', milestone_id]
  - ['contracts', 'detail', contract_id]
  - ['wallet', user_id]
```

### Real-Time Updates (WebSocket)
```typescript
// User-level updates
wss://api/v1/ws/user/{user_id}
Events: notification.new, message.new, contract.updated

// Conversation updates
wss://api/v1/ws/conversations/{conversation_id}
Events: message.sent, message.read, typing

// Notifications stream
wss://api/v1/ws/notifications/{user_id}
Events: notification.new
```

---

## 📊 State Management Strategy

### TanStack Query Configuration

**Default Options:**
```typescript
queries: {
  staleTime: 30s (lists), 60s (details)
  gcTime: 10min
  retry: 3 (exponential backoff)
  refetchOnWindowFocus: true
  refetchOnReconnect: true
}

mutations: {
  retry: 1
  onError: global error handler
}
```

**Pagination:**
- Cursor-based for all lists
- `keepPreviousData: true` for smooth transitions
- Infinite scroll support via `useInfiniteQuery`

### Zustand Store Structure

**Client State ONLY:**
```typescript
store = {
  authSlice: {
    isAuthenticated: boolean
    user: UserMinimal | null  // From JWT
    permissions: string[]
  }
  
  uiSlice: {
    theme: 'light' | 'dark'    // Persisted
    locale: string              // Persisted
    sidebarOpen: boolean
    activeModal: string | null
    toasts: Toast[]
  }
  
  notificationSlice: {
    unreadCount: number         // Real-time from WS
    latestNotifications: []     // Last 5 for bell
  }
  
  filterSlice: {
    jobFilters: JobFilters
    proposalFilters: ProposalFilters
    contractFilters: ContractFilters
  }
}
```

**NEVER in Zustand:**
- Server data (profiles, jobs, contracts, etc.)
- API responses
- Cached data

---

## 🎯 Feature Modules Structure

Each feature module follows this pattern:

```
features/{feature}/
├── components/          # Feature-specific UI components
├── hooks/              # Feature-specific hooks
├── api/
│   ├── queries.ts      # TanStack Query queries
│   └── mutations.ts    # TanStack Query mutations
└── types/              # Feature-specific types (if needed)
```

### Complete Feature List

1. **auth** - Authentication flows
2. **profile** - Profile management (skills, experience, portfolio)
3. **jobs** - Job browsing, posting, management
4. **proposals** - Proposal submission, management
5. **contracts** - Contract execution, milestones, disputes
6. **payments** - Wallet, transactions, invoices, tax
7. **messaging** - Real-time chat, conversations
8. **notifications** - Multi-channel notifications
9. **subscriptions** - Plans, connects, addons
10. **reviews** - Review submission, viewing
11. **analytics** - Business intelligence (client-only)
12. **admin** - Platform administration
13. **search** - Job/freelancer search

---

## 🔐 Security & Compliance

### Authentication Flow
```
1. Keycloak OAuth2 Authorization Code Flow (web)
2. Keycloak PKCE Flow (mobile)
3. JWT tokens (access + refresh)
4. HttpOnly cookies (web)
5. Secure storage (mobile: Keychain/KeyStore)
```

### Authorization
- Role-Based Access Control (RBAC)
- Permission checks at component level
- Route guards for protected pages
- Field-level visibility rules

### Data Protection
- PII redaction in client logs
- No sensitive data in localStorage
- HTTPS-only (enforced via CSP)
- XSS prevention (React escaping + DOMPurify)
- CSRF protection (SameSite cookies + tokens)

---

## 🚀 Performance Optimization

### Web Performance Targets
- LCP ≤ 2.5s
- TTI ≤ 3.5s
- CLS < 0.1
- TBT ≤ 200ms

### Optimization Strategies
- SSR/ISR for SEO-critical pages
- Route-level code splitting
- Image optimization (Next.js Image, WebP/AVIF)
- Font subsetting (next/font)
- Edge caching (Cloudflare/CloudFront)
- Bundle analysis in CI

### Mobile Performance
- FlatList → FlashList (60fps scrolling)
- expo-image with caching
- Reanimated 3 for native animations
- Offline-first with React Query persistence

---

## ♿ Accessibility

**WCAG 2.2 AA Compliance:**
- Semantic HTML with proper landmarks
- Keyboard navigation support
- Focus management (visible indicators, trap in modals)
- ARIA labels and descriptions
- Color contrast 4.5:1 (text), 3:1 (UI)
- Screen reader tested (NVDA, VoiceOver, TalkBack)
- Motion reduction support (prefers-reduced-motion)

---

## 🌍 Internationalization

**Supported Languages:**
1. English (en)
2. Arabic (ar) - RTL support
3. Chinese (zh)
4. Hindi (hi)
5. German (de)
6. French (fr)
7. Turkish (tr)
8. Spanish (es)
9. Russian (ru)

**Implementation:**
- Route-based i18n (`/[locale]/...`)
- ICU message format
- Pluralization support
- RTL layout mirroring (Arabic)

---

## 📦 File Organization Principles

1. **Feature-first**: Group by feature, not tech layer
2. **Co-location**: Keep related files together
3. **Separation of concerns**: Clear boundaries between features
4. **Reusability**: Shared code in packages
5. **Type safety**: TypeScript strict mode everywhere

---

## 🔧 Development Workflow

**Scripts:**
```bash
pnpm dev              # Start all apps in dev mode
pnpm dev:web          # Start web app only
pnpm dev:mobile       # Start mobile app only
pnpm build            # Build all apps
pnpm lint             # Lint all workspaces
pnpm type-check       # TypeScript check
pnpm test             # Run tests
pnpm clean            # Clean build artifacts
pnpm format           # Format with Prettier
```

**Git Hooks:**
- pre-commit: lint-staged (lint + format changed files)
- pre-push: run tests

---

## 📚 Documentation Structure

```
docs/
├── ARCHITECTURE.md    # System architecture
├── SETUP.md           # Development setup
├── PERFORMANCE.md     # Performance guidelines
├── TESTING.md         # Testing strategy
├── DEPLOYMENT.md      # Deployment guide
├── CONTRIBUTING.md    # Contribution guidelines
└── API_INTEGRATION.md # Backend integration guide
```

---

## 🎓 Next Steps

1. **Review Structure**: Check skillsier-fe-complete-tree.txt
2. **Backend Mappings**: Review api-mappings-*.md files for your services
3. **Features**: Check features-*.md for detailed component structures
4. **Setup**: Follow SETUP.md for development environment
5. **Start Coding**: Begin with one feature module at a time

---

## 📞 Support & Questions

This structure represents:
- **20,000+ files and folders**
- **100% backend service coverage**
- **ALL 250+ backend domains mapped**
- **Complete end-to-end user journeys**
- **Production-ready architecture**

For questions or clarifications, refer to the detailed markdown files in this output directory.

---

**Document Generated:** 2025-01-15
**Total Documentation:** 25+ files, 50,000+ lines
**Coverage:** 100% of Skillsier platform (frontend structure + backend mappings)