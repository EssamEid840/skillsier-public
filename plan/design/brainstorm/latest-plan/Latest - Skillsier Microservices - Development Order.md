Skillsier Microservices - Development Order
===========================================

Development Strategy
--------------------

Complete the shared modules and infrastructure first, as they are dependencies for all microservices. Then, complete each microservice **fully** (code + tests + deployment) before moving to the next. Each service should be production-ready with all features implemented, incorporating the shared modules (pkg/auth, platform-shared, contracts/events), config standardization, Dapr components, enhanced migrations, and K8s overlays.

Phase 0: Shared Infrastructure & Modules
----------------------------------------

### 0️⃣ Shared Modules (BEFORE ALL SERVICES)

**Why First:**

*   All services import these modules
    
*   Establishes common patterns (auth, outbox, events, logging, etc.)
    
*   Prevents duplication across services
    
*   Allows consistent implementation from the start
    

**What to Complete:**

*   ✅ **pkg/auth**: Keycloak integration, JWT verifier, RBAC middleware
    
*   ✅ **platform-shared**: Logging, tracing, metrics, httpx helpers, ginx middleware, outbox/inbox patterns, idempotency
    
*   ✅ **contracts/events**: Protobuf event schemas for all events, EVENTS.md catalog
    
*   ✅ Root monorepo structure: deployments/k8s (base + overlays), docs/decisions, scripts, Makefile
    

**Testing Milestones:**

*   Test auth flows (JWT validation, roles)
    
*   Test outbox pattern (reliable publishing)
    
*   Generate protobuf code and validate schemas
    
*   Test shared middleware (request ID, logging, CORS)
    

**Deliverable:**Shared modules ready for import by all services

Phase 1: Foundation (Core Identity & Discovery)
-----------------------------------------------

### 1️⃣ users-be (FIRST SERVICE)

**Why First:**

*   Every other service depends on user data
    
*   Integrates with Keycloak via pkg/auth
    
*   No dependencies on other services
    
*   Foundation for the entire platform
    

**What to Complete:**

*   ✅ User registration & authentication (Keycloak integration via pkg/auth)
    
*   ✅ Freelancer profiles (skills, experience, education, portfolio)
    
*   ✅ Client profiles (company info)
    
*   ✅ User settings & preferences
    
*   ✅ Admin actions (suspend, ban, warn) - prepare for admin-be later
    
*   ✅ Integrate platform-shared (outbox, logging, etc.)
    
*   ✅ Event publishing via contracts/events
    
*   ✅ Config standardization (internal/config/schema.go)
    
*   ✅ Dapr components (local/k8s)
    
*   ✅ Enhanced migrations with version tracking
    
*   ✅ K8s overlays for local/prod
    

**Testing Milestones:**

*   Create user accounts via Keycloak
    
*   Build complete freelancer profiles
    
*   Build complete client profiles
    
*   Update profiles and settings
    
*   Test user suspension/ban flows
    
*   Test event publishing and outbox
    

**Deliverable:**Fully functional user management system with profiles, using shared modules

### 2️⃣ jobs-be (SECOND)

**Why Second:**

*   Core platform feature (clients post jobs)
    
*   Only depends on users-be
    
*   Enables testing of job posting workflow
    
*   Establishes categories and skills taxonomy for entire platform
    

**What to Complete:**

*   ✅ Job posting (fixed-price, hourly)
    
*   ✅ Job categories & subcategories hierarchy
    
*   ✅ Skills taxonomy (global skills database)
    
*   ✅ Screening questions
    
*   ✅ Job invitations to freelancers
    
*   ✅ Job visibility controls
    
*   ✅ Saved searches
    
*   ✅ Job flags (prepare for admin-be)
    
*   ✅ Integrate platform-shared and pkg/auth
    
*   ✅ Event publishing via contracts/events
    
*   ✅ Config standardization
    
*   ✅ Dapr components
    
*   ✅ Enhanced migrations
    
*   ✅ K8s overlays
    

**Testing Milestones:**

*   Clients can post jobs
    
*   Browse jobs by category
    
*   Filter jobs by skills
    
*   Invite freelancers to jobs
    
*   Save job searches
    
*   View job analytics
    

**Deliverable:**Complete job posting and discovery system

### 3️⃣ proposals-be (THIRD)

**Why Third:**

*   Completes the connection loop (freelancers respond to jobs)
    
*   Depends on: users-be, jobs-be
    
*   Includes complex bidding system
    
*   Introduces connects system
    

**What to Complete:**

*   ✅ Proposal submission (cover letter, milestones)
    
*   ✅ **Complete bidding system** (bid amounts, bid history, auto-bidding)
    
*   ✅ **Bid strategies** (aggressive, conservative, auto)
    
*   ✅ Outbid alerts
    
*   ✅ Proposal templates
    
*   ✅ **Connects system** (deduct connects on submission)
    
*   ✅ Proposal boosts
    
*   ✅ Answer screening questions
    
*   ✅ Proposal flags (prepare for admin-be)
    
*   ✅ Integrate shared modules
    
*   ✅ Event publishing via contracts/events
    
*   ✅ Config, Dapr, migrations, K8s updates
    

**Testing Milestones:**

*   Freelancers submit proposals
    
*   Test bidding workflows (place bid, update bid)
    
*   Test auto-bidding strategies
    
*   Receive outbid alerts
    
*   Use proposal templates
    
*   Track connects usage
    
*   Boost proposals
    

**Deliverable:**Complete proposal and bidding system (Upwork-like)

Phase 2: Core Transactions (Work & Money)
-----------------------------------------

### 4️⃣ contracts-be (FOURTH)

**Why Fourth:**

*   Manages the actual work
    
*   Depends on: users-be, proposals-be, jobs-be
    
*   Critical for platform operations
    
*   Before financial-be because you need contracts to have payments
    

**What to Complete:**

*   ✅ Contract creation from accepted proposals
    
*   ✅ Milestone-based contracts
    
*   ✅ Hourly contracts with time tracking
    
*   ✅ Timesheet management (weekly submission)
    
*   ✅ Work diary (activity tracking)
    
*   ✅ Deliverable submission and approval
    
*   ✅ Contract amendments
    
*   ✅ Contract pause/resume
    
*   ✅ Contract termination
    
*   ✅ **Dispute system** (evidence, status tracking)
    
*   ✅ Integrate shared modules
    
*   ✅ Event publishing via contracts/events
    
*   ✅ Config, Dapr, migrations, K8s updates
    

**Testing Milestones:**

*   Create contracts from proposals
    
*   Submit and approve milestones
    
*   Log time for hourly contracts
    
*   Submit weekly timesheets
    
*   Upload and review deliverables
    
*   Pause and resume contracts
    
*   Handle contract disputes
    
*   Terminate contracts
    

**Deliverable:**Complete contract lifecycle management with dispute handling

### 5️⃣ financial-be (FIFTH)

**Why Fifth:**

*   Handles all money transactions
    
*   Depends on: users-be, contracts-be
    
*   Critical and complex (needs careful implementation)
    
*   Enables end-to-end payment flow
    

**What to Complete:**

*   ✅ Wallet management
    
*   ✅ Payment processing (Stripe, PayPal integration)
    
*   ✅ **Escrow system** (hold and release funds)
    
*   ✅ Transaction ledger
    
*   ✅ Payout processing (freelancer withdrawals)
    
*   ✅ Invoice generation (PDF)
    
*   ✅ Platform fee calculations
    
*   ✅ Refund processing
    
*   ✅ Bank account management
    
*   ✅ Tax form generation (W9, 1099)
    
*   ✅ Payment disputes
    
*   ✅ Integrate shared modules
    
*   ✅ Event publishing via contracts/events
    
*   ✅ Config, Dapr, migrations, K8s updates
    

**Testing Milestones:**

*   Fund wallet
    
*   Pay for contracts (money goes to escrow)
    
*   Release escrow on milestone completion
    
*   Calculate platform fees
    
*   Process payouts to freelancers
    
*   Generate invoices
    
*   Handle refunds
    
*   Process payment disputes
    
*   Generate tax forms
    

**Deliverable:**Complete financial system with escrow and multi-gateway support

Phase 3: User Experience Enhancement
------------------------------------

### 6️⃣ communications-be (SIXTH)

**Why Sixth:**

*   Enhances user experience significantly
    
*   Depends on: ALL previous services (sends notifications for all events)
    
*   Real-time features
    
*   Can now send notifications for all previous services' events
    

**What to Complete:**

*   ✅ Real-time messaging (WebSocket)
    
*   ✅ Conversations management
    
*   ✅ **In-app notifications system** (40+ notification types)
    
*   ✅ Email notifications (WildDuck integration)
    
*   ✅ Notification preferences
    
*   ✅ Online status / presence
    
*   ✅ Typing indicators
    
*   ✅ Read receipts
    
*   ✅ Notification grouping
    
*   ✅ Badge counts
    
*   ✅ Message flags (prepare for admin-be)
    
*   ✅ Integrate shared modules
    
*   ✅ Event publishing via contracts/events
    
*   ✅ Config, Dapr, migrations, K8s updates
    

**Testing Milestones:**

*   Send and receive real-time messages
    
*   Test all notification types from previous services
    
*   Configure notification preferences
    
*   Test online/offline status
    
*   Test typing indicators
    
*   Group similar notifications
    
*   Test email delivery via WildDuck
    

**Deliverable:**Complete real-time communication and notification system

### 7️⃣ storage-be (SEVENTH)

**Why Seventh:**

*   File management needed by multiple services
    
*   Depends on: users-be (portfolios), contracts-be (deliverables), proposals-be (attachments)
    
*   Can now handle all file upload needs
    

**What to Complete:**

*   ✅ File upload (chunked, resumable)
    
*   ✅ MinIO integration (self-hosted object storage)
    
*   ✅ Image processing (resize, optimize, thumbnails)
    
*   ✅ Video processing (thumbnails, transcoding)
    
*   ✅ File versioning
    
*   ✅ Access control
    
*   ✅ File sharing (share links)
    
*   ✅ Virus scanning (ClamAV)
    
*   ✅ File flags (prepare for admin-be)
    
*   ✅ Integrate shared modules
    
*   ✅ Event publishing via contracts/events
    
*   ✅ Config, Dapr, migrations, K8s updates
    

**Testing Milestones:**

*   Upload portfolio images
    
*   Upload contract deliverables
    
*   Upload proposal attachments
    
*   Process and optimize images
    
*   Generate video thumbnails
    
*   Create share links
    
*   Test file versioning
    
*   Scan for viruses
    

**Deliverable:**Complete file storage and media processing system

### 8️⃣ search-be (EIGHTH)

**Why Eighth:**

*   Improves job and freelancer discovery
    
*   Depends on: jobs-be, users-be
    
*   Requires indexed data from previous services
    
*   ML recommendations need historical data
    

**What to Complete:**

*   ✅ Elasticsearch integration
    
*   ✅ Job search with filters
    
*   ✅ Freelancer search with filters
    
*   ✅ Autocomplete suggestions
    
*   ✅ **Job recommendations** (ML-powered)
    
*   ✅ **Freelancer recommendations** (ML-powered)
    
*   ✅ Collaborative filtering
    
*   ✅ Content-based filtering
    
*   ✅ Job-freelancer matching
    
*   ✅ Personalized feeds
    
*   ✅ Trending jobs/skills
    
*   ✅ Similar jobs/users
    
*   ✅ Integrate shared modules
    
*   ✅ Event publishing via contracts/events
    
*   ✅ Config, Dapr, migrations, K8s updates
    

**Testing Milestones:**

*   Search jobs by keywords and filters
    
*   Search freelancers by skills
    
*   Test autocomplete
    
*   Get job recommendations (as freelancer)
    
*   Get freelancer recommendations (as client)
    
*   View personalized feed
    
*   Test matching algorithm
    
*   View trending items
    

**Deliverable:**Complete search and recommendation engine with ML

### 9️⃣ reviews-be (NINTH)

**Why Ninth:**

*   Builds reputation system
    
*   Depends on: contracts-be (can only review after contract completion)
    
*   Impacts search rankings from search-be
    
*   Should come after work has been done in previous services
    

**What to Complete:**

*   ✅ Review submission
    
*   ✅ Multi-criteria ratings
    
*   ✅ Review responses
    
*   ✅ **Badge system** (Top Rated, Rising Talent, etc.)
    
*   ✅ **Reputation calculation**
    
*   ✅ Review statistics
    
*   ✅ Private feedback
    
*   ✅ Helpful votes
    
*   ✅ Review flags (prepare for admin-be)
    
*   ✅ Integrate shared modules
    
*   ✅ Event publishing via contracts/events
    
*   ✅ Config, Dapr, migrations, K8s updates
    

**Testing Milestones:**

*   Submit reviews after contract completion
    
*   Rate on multiple criteria
    
*   Respond to reviews
    
*   Earn badges based on criteria
    
*   Calculate reputation scores
    
*   View review statistics
    
*   Submit private feedback
    
*   Vote on helpful reviews
    

**Deliverable:**Complete reputation and badge system

Phase 4: Monetization & Administration
--------------------------------------

### 🔟 subscriptions-be (TENTH)

**Why Tenth:**

*   Platform monetization
    
*   Depends on: users-be, proposals-be (connects usage)
    
*   Should come after core features work
    
*   Enables business model validation
    

**What to Complete:**

*   ✅ Subscription plans (Free, Basic, Plus, Enterprise)
    
*   ✅ Plan features and limits
    
*   ✅ Subscription billing
    
*   ✅ **Connects packages** (purchase connects)
    
*   ✅ Connects balance tracking
    
*   ✅ Usage quota enforcement
    
*   ✅ Promotional codes
    
*   ✅ Free trials
    
*   ✅ Feature toggles per plan
    
*   ✅ Auto-renewal
    
*   ✅ Subscription analytics
    
*   ✅ Integrate shared modules
    
*   ✅ Event publishing via contracts/events
    
*   ✅ Config, Dapr, migrations, K8s updates
    

**Testing Milestones:**

*   Subscribe to different plans
    
*   Purchase connect packages
    
*   Track connects usage (from proposals-be)
    
*   Enforce job posting limits
    
*   Apply promotional codes
    
*   Start free trials
    
*   Test auto-renewal
    
*   Cancel and change plans
    

**Deliverable:**Complete subscription and monetization system

### 1️⃣1️⃣ admin-be (ELEVENTH - LAST)

**Why Last:**

*   Needs ALL other services to be running
    
*   Admin operations span all services
    
*   Support system handles issues from all services
    
*   Moderation handles content from all services
    

**What to Complete:**

*   ✅ Admin user management (roles, permissions)
    
*   ✅ **Support ticket system** (SLA tracking, auto-assignment)
    
*   ✅ Ticket conversations
    
*   ✅ Support agent management
    
*   ✅ Canned responses
    
*   ✅ **Knowledge base** (articles, search)
    
*   ✅ FAQs
    
*   ✅ **Moderation queue** (centralized for all services)
    
*   ✅ User management (suspend, ban, verify, warn)
    
*   ✅ Content management (remove, hide, approve)
    
*   ✅ **Dispute resolution** (review disputes from contracts-be)
    
*   ✅ System configuration
    
*   ✅ Platform announcements
    
*   ✅ **Reports & analytics** (PDF, CSV, Excel)
    
*   ✅ Custom dashboards
    
*   ✅ Notification blasts
    
*   ✅ **Complete audit logging**
    
*   ✅ Integration with all 10 services
    
*   ✅ Integrate shared modules
    
*   ✅ Event publishing via contracts/events
    
*   ✅ Config, Dapr, migrations, K8s updates
    

**Testing Milestones:**

*   Create admin users with different roles
    
*   Create and assign support tickets
    
*   Use canned responses
    
*   Create KB articles
    
*   Review moderation queue (flags from all services)
    
*   Suspend/ban users
    
*   Remove flagged content
    
*   Resolve payment disputes
    
*   Generate platform reports
    
*   Send platform announcements
    
*   Review audit logs
    
*   Test permissions for different admin roles
    

**Deliverable:**Complete admin panel, support system, and moderation platform

Development Order Summary
-------------------------

```
Phase 0: Shared Infrastructure & Modules
└── 0. Shared Modules    (pkg/auth, platform-shared, contracts/events)

Phase 1: Foundation (Core Identity & Discovery)
├── 1. users-be          (Foundation)
├── 2. jobs-be           (Job Discovery)  
└── 3. proposals-be      (Bidding System)

Phase 2: Core Transactions (Work & Money)
├── 4. contracts-be      (Work Management)
└── 5. financial-be      (Payments & Escrow)

Phase 3: User Experience Enhancement
├── 6. communications-be (Real-time & Notifications)
├── 7. storage-be        (File Management)
├── 8. search-be         (Discovery & Recommendations)
└── 9. reviews-be        (Reputation System)

Phase 4: Monetization & Administration
├── 10. subscriptions-be (Monetization)
└── 11. admin-be         (Admin & Support)
```

Completion Criteria Per Service
-------------------------------

Before moving to the next service, ensure:

✅ **Code Complete:**

*   All domains implemented
    
*   All application services with commands/queries
    
*   All repositories
    
*   All event handlers (using contracts/events)
    
*   All HTTP handlers
    

✅ **Infrastructure Complete:**

*   PostgreSQL enhanced auto-migrations (with version tracking/safety)
    
*   Redis caching implemented
    
*   Kafka event publishing/subscribing (via outbox/inbox)
    
*   All external service clients (where needed)
    
*   Dapr components configured (local/k8s)
    

✅ **Shared Integration:**

*   Import and use pkg/auth for authentication
    
*   Use platform-shared for logging, tracing, metrics, middleware, outbox, etc.
    
*   Publish/subscribe events via contracts/events schemas
    

✅ **Testing Complete:**

*   Unit tests for domain logic
    
*   Integration tests for repositories
    
*   Integration tests for HTTP handlers
    
*   E2E tests for critical workflows
    

✅ **Deployment Ready:**

*   Dockerfile working
    
*   Kubernetes manifests with Kustomize overlays (base, local, prod)
    
*   ConfigMaps and Secrets defined
    
*   Health checks implemented
    

✅ **Documentation Complete:**

*   API documentation
    
*   Event documentation
    
*   README with setup instructions
    
*   MIGRATIONS.md with migration history
    

✅ **Can Demo:**

*   End-to-end user workflows work
    
*   Can show business value
    
*   No critical bugs
    

Testing Strategy Per Phase
--------------------------

### After Phase 0 (Shared Modules):

**You Can Test:**

*   Auth flows across mock services
    
*   Outbox/inbox reliability
    
*   Event schema validation
    
*   Shared middleware in isolation
    

### After Phase 1 (users, jobs, proposals):

**You Can Test:**

*   Complete freelancer onboarding
    
*   Complete client onboarding
    
*   Job posting flow
    
*   Job discovery and filtering
    
*   Proposal submission with bidding
    
*   Connects usage
    

### After Phase 2 (contracts, financial):

**You Can Test:**

*   End-to-end transaction flow
    
*   Contract creation from proposal
    
*   Milestone-based payments
    
*   Hourly time tracking
    
*   Escrow hold and release
    
*   Platform fee calculation
    
*   Payouts to freelancers
    

### After Phase 3 (communications, storage, search, reviews):

**You Can Test:**

*   Real-time messaging between users
    
*   Notifications for all previous workflows
    
*   File uploads for portfolios and deliverables
    
*   Advanced job search and recommendations
    
*   Review and reputation system
    
*   Complete user journey with reputation building
    

### After Phase 4 (subscriptions, admin):

**You Can Test:**

*   Subscription-based access
    
*   Platform monetization
    
*   Support ticket workflows
    
*   Content moderation
    
*   User management by admins
    
*   Platform analytics and reporting
    

Estimated Timeline Considerations
---------------------------------

**Per Shared Module (Rough Estimates):**

*   pkg/auth: 1-2 weeks
    
*   platform-shared: 2-3 weeks
    
*   contracts/events: 1-2 weeks
    

**Per Service (Rough Estimates):**

*   Simple services (reviews-be): 2-3 weeks
    
*   Medium services (jobs-be, proposals-be, contracts-be): 3-4 weeks
    
*   Complex services (financial-be, communications-be, search-be): 4-6 weeks
    
*   Very complex (admin-be): 6-8 weeks
    

**Total Estimated:** 45-55 weeks (about 1 year) for complete development, including shared modules

Key Dependencies to Remember
----------------------------

1.  **Shared modules** are needed by everyone
    
2.  **users-be** is needed by everyone
    
3.  **jobs-be** is needed by proposals, search, admin
    
4.  **proposals-be** is needed by contracts, subscriptions, admin
    
5.  **contracts-be** is needed by financial, reviews, admin
    
6.  **financial-be** is needed by subscriptions, admin
    
7.  **All services** are needed by communications-be (for notifications)
    
8.  **All services** are needed by admin-be (for moderation and management)
    

Recommendation
--------------

Start with **Phase 0: Shared Modules** immediately to establish the foundation. Once complete, move to **users-be** fully with all tests and deployment, then proceed in order.

Each completed phase is a major milestone and shows tangible progress toward the complete platform! 🚀