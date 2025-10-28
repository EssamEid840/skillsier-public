Platform Overview
-----------------

**Skillsier** is an enterprise-grade freelancing marketplace platform (similar to Upwork) built with a microservices architecture. The platform connects freelancers with clients for various projects and services.

### Core Business Model

*   **Freelancers** create profiles, browse jobs, submit proposals/bids, work on contracts
    
*   **Clients** post jobs, review proposals, hire freelancers, manage contracts
    
*   **Platform** facilitates connections, handles payments via escrow, manages disputes, provides support
    

Architecture Principles
-----------------------

### 1\. Microservices Architecture

*   **11 independent microservices** - each with its own database
    
*   **Clean Architecture** (DDD) - Domain, Application, Infrastructure, Interface layers
    
*   **Event-Driven Communication** - Kafka for async messaging, with protobuf event schemas in contracts/events
    
*   **Polyglot Persistence** - PostgreSQL per service, Redis for caching, Elasticsearch for search
    

### 2\. Technology Stack

**Backend:**

*   Language: **Go (Golang)**
    
*   Framework: Standard library + Gin/Echo/Chi
    
*   Database: **PostgreSQL** (per service with enhanced auto-migrations including version tracking and safety checks)
    
*   Cache: **Redis**
    
*   Search: **Elasticsearch/OpenSearch**
    
*   Message Queue: **Apache Kafka**
    
*   Object Storage: **MinIO** (self-hosted)
    

**Shared Modules:**

*   **pkg/auth**: Centralized Keycloak authentication (JWT verification, RBAC middleware)
    
*   **platform-shared**: Cross-cutting utilities (logging, tracing, metrics, outbox/inbox patterns, HTTP helpers, idempotency)
    
*   **contracts/events**: Protobuf-defined event schemas for all inter-service events
    

**Authentication & Authorization:**

*   **Keycloak** - Handles all authentication and authorization (integrated via pkg/auth)
    
*   JWT tokens for API authentication
    

**Communications:**

*   **WildDuck** - Self-hosted SMTP email server (free, no paid services)
    
*   **WebSocket** - Real-time messaging and notifications
    
*   **Server-Sent Events (SSE)** - Live updates
    

**Infrastructure:**

*   **Kubernetes (K8s)** - Container orchestration with Kustomize (base manifests + overlays for local/prod)
    
*   **Docker** - Containerization
    
*   **Kong/NGINX** - API Gateway
    
*   **Dapr** - For pub/sub (Kafka) and state management, with per-service components (local/k8s split)
    
*   Optional: **Istio/Linkerd** - Service Mesh
    

**Observability:**

*   Logging: **ELK Stack** or **Grafana Loki** (using platform-shared/logging)
    
*   Monitoring: **Prometheus + Grafana** (using platform-shared/metrics)
    
*   Tracing: **Jaeger** (using platform-shared/tracing)
    

**CI/CD:**

*   **GitHub Actions** / GitLab CI
    
*   **ArgoCD** - GitOps for K8s deployments
    

11 Microservices Architecture
-----------------------------

### 1\. users-be

**Responsibility:** User profiles, authentication integration, freelancer/client management

**Key Domains:**

*   User accounts (integrated with Keycloak via pkg/auth)
    
*   Freelancer profiles (skills, experience, education, certifications, portfolio, languages, rates, stats)
    
*   Client profiles (company info, hiring stats, spending)
    
*   User verification
    
*   Settings & preferences
    
*   Saved items (jobs, freelancers)
    
*   User suspensions, bans, warnings (admin actions)
    

**Key Features:**

*   Complete freelancer profile management
    
*   Client company profiles
    
*   Portfolio management with media
    
*   Multi-language proficiency
    
*   Availability & timezone handling
    
*   Admin action tracking (suspensions, bans, warnings)
    

**Events Published (via contracts/events):**

*   UserCreated, UserUpdated, UserVerified
    
*   FreelancerProfileCompleted, ClientProfileCompleted
    
*   UserSuspended, UserBanned, UserWarned
    

### 2\. jobs-be

**Responsibility:** Job postings, categories, skills taxonomy, job lifecycle

**Key Domains:**

*   Job postings (fixed-price, hourly)
    
*   Job categories & subcategories
    
*   Skills taxonomy (global skill database)
    
*   Job skills (required, preferred)
    
*   Screening questions
    
*   Job attachments
    
*   Job invitations to freelancers
    
*   Job views tracking
    
*   Saved searches
    
*   Job flags (admin moderation)
    

**Key Features:**

*   Hierarchical job categories
    
*   Skills requirement matching
    
*   Custom screening questions
    
*   Private/public/invite-only jobs
    
*   Job view analytics
    
*   Search filter persistence
    
*   Admin moderation (flag, remove, hide, feature)
    

**Events Published (via contracts/events):**

*   JobPosted, JobUpdated, JobClosed
    
*   JobInvitationSent
    
*   JobFlagged, JobRemoved, JobHidden, JobFeatured
    

### 3\. proposals-be

**Responsibility:** Proposals/bids from freelancers, bidding system, connects management

**Key Domains:**

*   Proposals (cover letter, milestones, attachments, question answers)
    
*   **Bidding system** (competitive bidding, bid amounts, bid history)
    
*   **Bid strategies** (aggressive, conservative, auto-bid)
    
*   Bid notifications (outbid alerts)
    
*   Proposal templates
    
*   Connects (credits used to submit proposals)
    
*   Proposal boosts (increase visibility)
    
*   Proposal analytics (view tracking)
    
*   Proposal flags (admin moderation)
    

**Key Features:**

*   **Full bidding system** like Upwork
    
*   Auto-bidding with strategies
    
*   Outbid alerts and notifications
    
*   Connects system (pay-to-apply)
    
*   Proposal boosting
    
*   Template library
    
*   Performance analytics
    
*   Admin moderation
    

**Events Published (via contracts/events):**

*   ProposalSubmitted, ProposalAccepted, ProposalRejected
    
*   BidPlaced, BidUpdated, OutbidAlert
    
*   ConnectUsed
    
*   ProposalFlagged, ProposalRemoved
    

### 4\. contracts-be

**Responsibility:** Contract management, milestones, time tracking, work diary

**Key Domains:**

*   Contracts (fixed-price, hourly, milestone-based)
    
*   Contract milestones (deliverables, approvals)
    
*   Timesheets (weekly time tracking)
    
*   Time entries (individual logged hours)
    
*   Work diary (activity tracking, screenshots)
    
*   Contract templates
    
*   Contract amendments (change requests)
    
*   Deliverables (submissions, revisions)
    
*   Contract pause/resume
    
*   Contract termination
    
*   **Contract disputes** (evidence, resolution)
    

**Key Features:**

*   Milestone-based payments
    
*   Hourly time tracking with work diary
    
*   Contract modification system
    
*   Deliverable management with revisions
    
*   Dispute resolution system
    
*   Admin intervention for disputes
    

**Events Published (via contracts/events):**

*   ContractCreated, ContractStarted, ContractEnded
    
*   MilestoneCompleted, MilestoneApproved
    
*   TimesheetSubmitted
    
*   DisputeOpened, DisputeResolved
    

### 5\. financial-be

**Responsibility:** Payments, escrow, wallets, payouts, invoices, platform fees

**Key Domains:**

*   Wallets (user balance management)
    
*   Transactions (double-entry ledger)
    
*   Payments (Stripe, PayPal integration)
    
*   Escrow (hold funds until work completion)
    
*   Payouts (freelancer withdrawals)
    
*   Invoices (PDF generation, tax calculations)
    
*   Platform fees (tiered fee structure)
    
*   Refunds
    
*   Payment disputes (chargebacks)
    
*   Tax management (W9, 1099 forms)
    
*   Bank accounts (verification)
    
*   Payment schedules (recurring)
    

**Key Features:**

*   Multi-currency support
    
*   Escrow system for secure payments
    
*   Automated invoicing
    
*   Tiered platform fees
    
*   Batch payout processing
    
*   Tax form generation
    
*   Payment gateway abstraction (Stripe, PayPal)
    
*   Admin payment dispute resolution
    

**Events Published (via contracts/events):**

*   PaymentProcessed, PaymentFailed
    
*   EscrowHeld, EscrowReleased
    
*   PayoutProcessed, InvoiceGenerated
    
*   PaymentDisputeOpened, RefundProcessedByAdmin
    

### 6\. communications-be

**Responsibility:** Real-time messaging, notifications (in-app, email), online status

**Key Domains:**

*   Conversations (chat threads)
    
*   Messages (attachments, reactions, read receipts, mentions)
    
*   **In-app notifications** (real-time, badge counts, grouping, actions)
    
*   Notification preferences (per type, per channel)
    
*   Email (WildDuck integration, templates)
    
*   Notification queue (priority-based)
    
*   Delivery logs
    
*   Unsubscribe management
    
*   Online status (presence)
    
*   Message flags (admin moderation)
    

**Key Features:**

*   **Real-time WebSocket** for instant messaging
    
*   **Server-Sent Events (SSE)** for live updates
    
*   **40+ notification types** covering all platform events
    
*   Notification grouping and aggregation
    
*   Badge count management
    
*   Typing indicators
    
*   Online/offline presence
    
*   **WildDuck email server** (self-hosted, free)
    
*   Email template system
    
*   Multi-channel orchestration
    
*   Admin message moderation
    

**Events Published (via contracts/events):**

*   MessageSent, NotificationDelivered
    
*   EmailSent, InAppNotificationSent
    
*   MessageFlagged, MessageRemoved
    

**Notification Types Covered:**

*   Job-related (new job, invitation, job closed)
    
*   Proposal/Bidding (viewed, accepted, rejected, outbid alerts)
    
*   Contract (created, milestone completed, timesheet submitted)
    
*   Financial (payment received, payment sent, escrow events)
    
*   Reviews (new review, badge earned)
    
*   Messages (new message, typing indicator)
    
*   Subscriptions (expiring, connects low)
    
*   System (verification, security alerts)
    

### 7\. storage-be

**Responsibility:** File uploads, media processing, file management

**Key Domains:**

*   Files (metadata, versioning)
    
*   Folders (hierarchical structure)
    
*   Uploads (chunked, resumable)
    
*   Media (image/video processing, thumbnails, variants)
    
*   Access control (permissions)
    
*   File sharing (share links)
    
*   File versioning
    
*   File flags (admin moderation)
    

**Key Features:**

*   Chunked & resumable uploads
    
*   Image optimization & thumbnails
    
*   Video transcoding
    
*   File versioning system
    
*   Share link generation
    
*   **MinIO** integration (self-hosted object storage)
    
*   ClamAV virus scanning
    
*   Admin file moderation
    

**Events Published (via contracts/events):**

*   FileUploaded, FileDeleted
    
*   MediaProcessed
    
*   FileFlagged, FileRemoved
    

### 8\. search-be

**Responsibility:** Search indexing, recommendations, matching, personalized feeds

**Key Domains:**

*   Search indices (jobs, users/freelancers)
    
*   Search queries (logs, filters)
    
*   **Recommendations** (job recommendations, freelancer recommendations)
    
*   **Recommendation models** (ML models, features, training data)
    
*   User preferences (implicit signals, explicit preferences)
    
*   Matching (job-freelancer matching with scoring)
    
*   Personalized feeds
    
*   Trending (jobs, skills)
    
*   Saved searches (alerts)
    
*   Similarity (similar jobs, similar users)
    

**Key Features:**

*   **Elasticsearch** integration for full-text search
    
*   **ML-powered recommendations**:
    
    *   Collaborative filtering
        
    *   Content-based filtering
        
    *   Hybrid recommendation engine
        
*   User behavior tracking
    
*   Cold-start handling for new users
    
*   Diversity optimization in recommendations
    
*   Real-time trending calculation
    
*   Faceted search
    
*   Autocomplete suggestions
    
*   Admin content removal from index
    

**Events Published (via contracts/events):**

*   JobIndexed, UserIndexed
    
*   RecommendationGenerated, MatchFound
    

### 9\. reviews-be

**Responsibility:** Reviews, ratings, badges, reputation system

**Key Domains:**

*   Reviews (ratings breakdown, responses, helpful votes)
    
*   Ratings (criteria-based, aggregation)
    
*   Badges (Top Rated, Rising Talent, achievement badges)
    
*   User badges (assignments, achievement dates)
    
*   Reputation (scores, history, calculation)
    
*   Private feedback
    
*   Review flags (admin moderation)
    
*   Review statistics
    
*   Review reminders
    

**Key Features:**

*   Multi-criteria rating system
    
*   Review responses
    
*   Badge achievement system
    
*   Reputation calculation algorithm
    
*   Review reminder system
    
*   Helpful vote system
    
*   Admin review moderation
    

**Events Published (via contracts/events):**

*   ReviewSubmitted, ReviewResponded
    
*   BadgeAwarded, ReputationUpdated
    
*   ReviewFlagged, ReviewRemoved
    

### 10\. subscriptions-be

**Responsibility:** Membership plans, connects/credits, billing, usage tracking

**Key Domains:**

*   Subscription plans (features, pricing, limits)
    
*   User subscriptions (billing cycles, auto-renewal)
    
*   Connects (credits for proposals)
    
*   Connect packages
    
*   Connect transactions
    
*   Usage tracking (quotas, limits)
    
*   Add-ons
    
*   Promotional codes (discounts, usage limits)
    
*   Free trials (eligibility)
    
*   Billing history
    
*   Feature toggles (per plan)
    

**Key Features:**

*   Tiered subscription plans
    
*   Connects system (Upwork-style credits)
    
*   Usage quota enforcement
    
*   Promotional code system
    
*   Free trial management
    
*   Auto-renewal handling
    
*   Feature gating per plan
    
*   Admin subscription overrides
    

**Events Published (via contracts/events):**

*   SubscriptionCreated, SubscriptionRenewed, SubscriptionCancelled
    
*   ConnectsPurchased, ConnectsUsed
    
*   SubscriptionCancelledByAdmin, ConnectsAddedByAdmin
    

### 11\. admin-be

**Responsibility:** Admin operations, support ticketing, moderation, dispute resolution, platform management

**Key Domains:**

*   Admin users (roles: SuperAdmin, Moderator, SupportAgent, FinancialAdmin, ReportingAdmin)
    
*   Admin permissions (granular permissions)
    
*   **Support tickets** (categories, priorities, SLA tracking)
    
*   Ticket messages (conversations)
    
*   Support agents (availability, performance stats)
    
*   Canned responses
    
*   **Knowledge base** (articles, versioning, search)
    
*   FAQs
    
*   **Moderation queue** (content flags from all services)
    
*   User actions (suspend, ban, verify, warn)
    
*   Content actions (remove, hide, approve, feature)
    
*   **Dispute resolution** (evidence, decisions)
    
*   System configuration (feature flags, maintenance mode)
    
*   Announcements (platform-wide, scheduled)
    
*   Reports (PDF, CSV, Excel exports)
    
*   Analytics dashboards
    
*   Audit logs (complete audit trail)
    
*   Notification blasts (bulk emails)
    
*   Platform statistics
    

**Key Features:**

*   Complete support ticket system with SLA tracking
    
*   Auto-assignment engine for tickets
    
*   Ticket escalation management
    
*   Knowledge base with full-text search
    
*   Centralized moderation queue for all flagged content
    
*   Auto-moderation rules
    
*   User management (suspend, ban, warn)
    
*   Content management across all services
    
*   Dispute resolution workflow
    
*   Comprehensive reporting system
    
*   Custom analytics dashboards
    
*   Platform announcement system
    
*   Complete audit trail for compliance
    
*   Integration with all 10 other services
    

**Events Published (via contracts/events):**

*   admin.user.suspended, admin.user.banned, admin.user.warned
    
*   admin.job.removed, admin.proposal.removed, admin.review.removed
    
*   admin.moderation.approved, admin.flag.resolved
    
*   admin.dispute.resolved
    
*   admin.config.updated, admin.announcement.published
    

**Events Subscribed:**

*   All flagging events from other services
    
*   System errors and alerts
    

Key Architectural Patterns
--------------------------

### 1\. Communication Patterns

**Synchronous (REST/HTTP):**

*   Client → API Gateway → Microservices
    
*   Admin-be → Other services (for direct operations)
    

**Asynchronous (Kafka Events):**

*   All inter-service communication
    
*   Event-driven architecture
    
*   Eventual consistency
    
*   Events defined in protobuf schemas (contracts/events)
    

**Outbox Pattern:**

*   Every service uses outbox pattern for reliable event publishing (via platform-shared/outbox)
    
*   Ensures no lost events even if Kafka is down
    

**Inbox Pattern:**

*   Message deduplication to prevent duplicate processing (via platform-shared/inbox)
    

**Idempotency:**

*   HTTP request deduplication (via platform-shared/idempotency)
    

### 2\. Data Patterns

**Database Per Service:**

*   Each microservice has its own PostgreSQL database
    
*   Schema isolation
    
*   Independent scaling
    

**Auto-Migrations:**

*   Using GORM/migrate auto-migration (not SQL files)
    
*   Enhanced with version tracking, safety checks, and MIGRATIONS.md documentation
    
*   Migrations run on service startup
    

**Caching Strategy:**

*   Redis for frequently accessed data
    
*   Cache-aside pattern
    
*   TTL-based invalidation
    

**Config Standardization:**

*   Per-service config/ dir with schema.go (typed config struct)
    
*   Viper loader (flags > env > file > defaults)
    
*   Environment-specific files (default.yaml, dev.yaml, prod.yaml)
    

### 3\. Security Patterns

**Authentication:**

*   Keycloak handles all authentication (centralized via pkg/auth)
    
*   JWT tokens validated at API Gateway
    
*   Admin-be has special admin role checks
    

**Authorization:**

*   Role-Based Access Control (RBAC) via pkg/auth
    
*   Permission checks in middleware
    
*   Admin permissions are granular (50+ permissions)
    

**Data Protection:**

*   Sensitive financial data encrypted
    
*   PII handling compliance
    
*   Audit logging for all admin actions
    

Inter-Service Event Flow
------------------------

### Example: Complete Job-to-Payment Flow

1.  **Client posts job** (jobs-be)
    
    *   Event: JobPosted → search-be, communications-be, subscriptions-be
        
2.  **Freelancer submits proposal** (proposals-be)
    
    *   Event: ProposalSubmitted → jobs-be, communications-be, subscriptions-be
        
    *   Event: ConnectUsed → subscriptions-be
        
3.  **Client accepts proposal** (proposals-be)
    
    *   Event: ProposalAccepted → contracts-be, communications-be
        
4.  **Contract created** (contracts-be)
    
    *   Event: ContractCreated → financial-be, communications-be, users-be
        
5.  **Client funds escrow** (financial-be)
    
    *   Event: EscrowHeld → contracts-be, communications-be
        
6.  **Freelancer completes milestone** (contracts-be)
    
    *   Event: MilestoneCompleted → financial-be, communications-be
        
7.  **Client approves milestone** (contracts-be)
    
    *   Event: MilestoneApproved → financial-be, communications-be
        
8.  **Escrow releases payment** (financial-be)
    
    *   Event: EscrowReleased → contracts-be, users-be, communications-be
        
    *   Event: PaymentProcessed → users-be, communications-be
        
9.  **Contract completed** (contracts-be)
    
    *   Event: ContractEnded → reviews-be, communications-be
        
10.  **Review submitted** (reviews-be)
    
    *   Event: ReviewSubmitted → users-be, search-be, communications-be
        

Resource Requirements
---------------------

### Minimum Production Cluster

yaml

```
Kubernetes Cluster:
  Nodes: 3 nodes × 10GB RAM = 30GB total

Application Services (11 services):
  - users-be: 2 replicas × 512MB = 1GB
  - jobs-be: 2 replicas × 512MB = 1GB
  - proposals-be: 2 replicas × 512MB = 1GB
  - contracts-be: 2 replicas × 512MB = 1GB
  - financial-be: 3 replicas × 1GB = 3GB (critical service)
  - communications-be: 2 replicas × 512MB = 1GB
  - storage-be: 2 replicas × 512MB = 1GB
  - search-be: 2 replicas × 1GB = 2GB
  - reviews-be: 2 replicas × 512MB = 1GB
  - subscriptions-be: 2 replicas × 512MB = 1GB
  - admin-be: 2 replicas × 1GB = 2GB
  Subtotal: ~16GB

Infrastructure Services:
  - PostgreSQL: 2GB (multiple instances)
  - Redis: 1GB
  - Elasticsearch: 3GB
  - Kafka: 2GB
  - Keycloak: 1GB
  - MinIO: 1GB
  - WildDuck: 1GB
  Subtotal: ~11GB

Total: ~27GB (fits in 30GB cluster)
```

Key Design Decisions
--------------------

### 1\. Why Separate Microservices?

*   **users-be** vs **jobs-be**: Different scaling needs, different access patterns
    
*   **proposals-be** separate: Bidding system is complex enough to warrant its own service
    
*   **financial-be** separate: Critical service, needs special security, separate scaling
    
*   **communications-be** separate: Real-time requirements, WebSocket handling
    
*   **admin-be** separate: Different authentication requirements, admin-only access
    

### 2\. Why Event-Driven?

*   Loose coupling between services
    
*   Async processing for better performance
    
*   Easy to add new services without modifying existing ones
    
*   Natural audit trail
    
*   Protobuf schemas (contracts/events) for type-safe events
    

### 3\. Why Keycloak?

*   Production-ready authentication solution
    
*   SSO support
    
*   Admin UI for user management
    
*   Role and permission management
    
*   Open-source, self-hosted
    
*   Centralized in pkg/auth for all services
    

### 4\. Why Self-Hosted Tools?

*   **WildDuck** for email: Free, no monthly costs
    
*   **MinIO** for storage: S3-compatible, self-hosted
    
*   **Elasticsearch**: Self-hosted instead of cloud services
    
*   Cost savings at scale
    
*   Data sovereignty
    
*   No vendor lock-in
    

### 5\. Why Go?

*   Excellent performance
    
*   Great for microservices
    
*   Strong concurrency support
    
*   Small memory footprint
    
*   Fast compilation
    
*   Great tooling and libraries
    

### 6\. Why Shared Modules?

*   **pkg/auth**: Centralize auth logic, avoid duplication
    
*   **platform-shared**: Common utilities (outbox, logging, etc.) for consistency
    
*   **contracts/events**: Versioned event schemas to prevent breaking changes
    

### 7\. Why Dapr?

*   Simplifies pub/sub and state management
    
*   Environment-agnostic components (local vs k8s)
    
*   Integrates with Kafka
    

### 8\. Why Kustomize for K8s?

*   Base manifests + overlays for env-specific configs
    
*   Better than raw YAML for maintainability
    

Common Folder Structure Pattern
-------------------------------

All 11 microservices follow this structure (updated with shared modules and config standardization):

text

```
<service-name>-be/
├── cmd/api/main.go
├── internal/
│   ├── domain/              # Entities, value objects, repositories
│   ├── application/         # Business logic, commands, queries, DTOs
│   ├── infrastructure/      # DB, Kafka, Redis, external clients
│   ├── interfaces/http/     # HTTP handlers, middleware, router
│   └── config/              # Standardized config (schema.go, loader.go)
├── config/                  # Env-specific YAML files (default, dev, prod)
├── dapr/                    # Dapr components (local/, k8s/)
├── pkg/                     # Service-specific utilities (imports shared modules)
├── deployments/k8s/         # Kubernetes manifests
├── scripts/                 # Setup and utility scripts
├── tests/                   # Unit, integration, e2e tests
├── docs/                    # Documentation (incl. MIGRATIONS.md)
├── go.mod, Makefile, Dockerfile, etc.
```

*   Imports: pkg/auth, platform-shared, contracts/events
    
*   Migrations: Enhanced with version tracking and safety checks
    

Critical Features Summary
-------------------------

### Bidding System (proposals-be)

*   Like Upwork's bidding
    
*   Auto-bidding strategies
    
*   Outbid alerts
    
*   Bid history tracking
    

### Connects System (proposals-be + subscriptions-be)

*   Pay-to-apply model
    
*   Connect packages
    
*   Usage tracking
    

### Escrow System (financial-be)

*   Secure payments
    
*   Milestone-based releases
    
*   Dispute protection
    

### Real-time Communications (communications-be)

*   WebSocket for instant messaging
    
*   SSE for live notifications
    
*   40+ notification types
    
*   Presence/online status
    

### Recommendation Engine (search-be)

*   ML-powered job matching
    
*   Collaborative filtering
    
*   Content-based filtering
    
*   Cold-start handling
    

### Admin Panel (admin-be)

*   Support ticket system
    
*   Content moderation
    
*   User management
    
*   Dispute resolution
    
*   Analytics & reporting
    

Future Extensibility
--------------------

The architecture supports:

*   Adding new microservices easily
    
*   Horizontal scaling per service
    
*   Multiple deployment environments (via Kustomize overlays)
    
*   A/B testing with feature flags
    
*   Multi-region deployment
    
*   Mobile app integration (same APIs)
    

Important Notes
---------------

1.  **No SQL migration files**: All services use auto-migrations via GORM (enhanced with version tracking/safety)
    
2.  **No paid external services**: Everything is self-hosted (WildDuck, MinIO, etc.)
    
3.  **Event-driven**: All inter-service communication via Kafka events (defined in protobuf schemas)
    
4.  **Keycloak integration**: All authentication goes through Keycloak (via pkg/auth)
    
5.  **Admin service**: Has clients for all other 10 services to perform admin operations
    
6.  **Outbox pattern**: Used in all services for reliable event publishing (via platform-shared)
    
7.  **Clean Architecture**: All services follow DDD with clean architecture layers
    
8.  **Dapr**: Used for pub/sub and state, with env-specific components
    

This is the complete Skillsier platform architecture with 11 production-ready microservices.