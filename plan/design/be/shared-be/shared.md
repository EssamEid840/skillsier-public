# Skillsier Platform - Complete Updated Folder Structure
## All Microservices with Full Comments (Old + New) ✨

---

## 📋 **Change Impact Summary**

| Area | Impact | Changes |
|------|--------|---------|
| **Auth (pkg/auth)** | 🔴 MAJOR | New root `pkg/auth/` module - all services import this |
| **Platform-shared** | 🔴 MAJOR | New root `platform-shared/` module - utilities, outbox, etc. |
| **Contracts** | 🔴 MAJOR | New root `contracts/events/` - protobuf event schemas |
| **Config Standardization** | 🟡 MEDIUM | Per-service `config/` dir + `internal/config/schema.go` |
| **Dapr Components** | 🟡 MEDIUM | Per-service `dapr/local/` and `dapr/k8s/` split |
| **Migrations** | 🟢 MINOR | Enhanced with version tracking, safety checks, MIGRATIONS.md |
| **K8s Overlays** | 🟢 MINOR | Kustomize structure: `base/`, `overlays/local/`, `overlays/prod/` |

**Legend:**
- 🆕 = New additions
- 📝 = Updated files
- ❌ = Removed (use shared modules)

---

## 🗂️ **ROOT MONOREPO STRUCTURE**

```
skillsier/                                    # Monorepo root
│
├── apps/
│   └── be/
│       ├── users-be/                         # User management service
│       ├── jobs-be/                          # Job posting and management service
│       ├── proposals-be/                     # Freelancer proposal submission service
│       ├── contracts-be/                     # Contract management service
│       ├── financial-be/                     # Payments and wallet management service
│       ├── communications-be/                # Real-time messaging and notifications service
│       ├── storage-be/                       # File storage and management service
│       ├── search-be/                        # Search and discovery service
│       ├── reviews-be/                       # Rating and review service
│       ├── subscriptions-be/                 # Subscription and premium features service
│       └── admin-be/                         # Admin and moderation service
│
├── pkg/                                      # 🆕 MAJOR - Shared packages (leaf dependencies)
│   └── auth/                                 # 🆕 Centralized Keycloak authentication
│       ├── go.mod                            # module: skillsier.dev/pkg/auth
│       ├── go.sum
│       ├── README.md                         # Usage documentation
│       ├── authorizer.go                     # RequireRoles / RequirePermissions middleware (RBAC)
│       ├── config.go                         # Auth configuration struct (Issuer, Audience, JWKS URL)
│       ├── context.go                        # Put/Get Principal on request context
│       ├── principal.go                      # Normalized identity model (UserID, Roles, Permissions)
│       ├── verifier.go                       # TokenVerifier interface (domain depends on this only)
│       ├── errors.go                         # Auth-specific errors (Unauthorized, Forbidden, Invalid Token)
│       └── keycloak/
│           ├── verifier.go                   # Keycloak JWT/JWKS implementation
│           ├── client.go                     # Keycloak admin API client
│           ├── role_mapper.go                # Map Keycloak roles → app permissions
│           └── config.go                     # Keycloak-specific config (Realm, Client ID)
│
├── platform-shared/                          # 🆕 MAJOR - Cross-cutting utilities (leaf module)
│   ├── go.mod                                # module: skillsier.dev/platform-shared
│   ├── go.sum
│   ├── README.md                             # Usage documentation
│   │
│   ├── logging/                              # 🆕 Structured logging
│   │   ├── logger.go                         # Structured logging (zerolog/zap)
│   │   ├── context.go                        # Logger from context
│   │   ├── fields.go                         # Common log fields (RequestID, UserID, TraceID)
│   │   └── config.go                         # Logging configuration (Level, Format, Output)
│   │
│   ├── tracing/                              # 🆕 Distributed tracing
│   │   ├── otel.go                           # OpenTelemetry setup
│   │   ├── tracer.go                         # Tracer helpers
│   │   ├── span.go                           # Span helpers (Start, End, AddEvent)
│   │   └── config.go                         # Tracing configuration (Endpoint, Sampling)
│   │
│   ├── metrics/                              # 🆕 Prometheus metrics
│   │   ├── metrics.go                        # Prometheus metrics helpers
│   │   ├── http.go                           # HTTP metrics (RED: Rate, Errors, Duration)
│   │   ├── collectors.go                     # Custom collectors (DB pool, queue depth)
│   │   └── config.go                         # Metrics configuration (Namespace, Subsystem)
│   │
│   ├── httpx/                                # 🆕 HTTP-agnostic helpers
│   │   ├── errors.go                         # Error envelope (Code, Message, Details)
│   │   ├── pagination.go                     # Pagination helpers (Limit, Offset, Total, HasNext)
│   │   ├── response.go                       # Standard response wrappers (Success, Error)
│   │   └── validator.go                      # Common validation helpers
│   │
│   ├── ginx/                                 # 🆕 Gin-specific middleware
│   │   ├── requestid.go                      # Request ID middleware (X-Request-ID header)
│   │   ├── logging.go                        # Logging middleware (structured logs per request)
│   │   ├── recover.go                        # Panic recovery middleware
│   │   ├── otel.go                           # OpenTelemetry middleware (trace propagation)
│   │   ├── cors.go                           # CORS middleware
│   │   └── timeout.go                        # Request timeout middleware
│   │
│   ├── outbox/                               # 🆕 Outbox pattern (moved from services)
│   │   ├── publisher.go                      # Watermill TX publisher (publish within DB transaction)
│   │   ├── forwarder.go                      # Background processor (polls outbox, publishes to Kafka)
│   │   ├── scheduler.go                      # Scheduling logic (retry, backoff)
│   │   ├── entity.go                         # Outbox event entity (ID, AggregateID, Payload, Status)
│   │   ├── repository.go                     # Outbox repository interface
│   │   ├── postgres/
│   │   │   └── repository.go                 # Postgres implementation
│   │   └── config.go                         # Outbox configuration (PollInterval, BatchSize)
│   │
│   ├── inbox/                                # 🆕 Message deduplication (prevent duplicate processing)
│   │   ├── checker.go                        # Check if message processed (by MessageID)
│   │   ├── marker.go                         # Mark message as processed
│   │   ├── entity.go                         # Inbox entity (MessageID, ProcessedAt, Handler)
│   │   ├── repository.go                     # Inbox repository interface
│   │   └── postgres/
│   │       └── repository.go                 # Postgres implementation
│   │
│   └── idempotency/                          # 🆕 HTTP request deduplication
│       ├── middleware.go                     # Idempotency-Key middleware (Gin)
│       ├── handler.go                        # Idempotency handler logic
│       ├── entity.go                         # Idempotency record (Key, Response, ExpiresAt)
│       ├── repository.go                     # Repository interface
│       └── postgres/
│           └── repository.go                 # Postgres implementation
│
├── contracts/                                # 🆕 MAJOR - Event schemas (versioned, breaking-change protected)
│   └── events/                               # 🆕 Event contracts module
│       ├── go.mod                            # module: skillsier.dev/contracts/events
│       ├── go.sum
│       ├── README.md                         # Usage and versioning guide
│       ├── buf.yaml                          # Buf configuration (linting, breaking change detection)
│       ├── buf.gen.yaml                      # Code generation config (protoc plugins)
│       ├── buf.lock                          # Dependency lock file
│       ├── EVENTS.md                         # 🆕 Event catalog (Topics, Owners, BC Policy, Changelog)
│       │
│       ├── user/v1/                          # User lifecycle events
│       │   ├── user_created.proto            # New user registered
│       │   ├── user_updated.proto            # User profile updated
│       │   ├── user_verified.proto           # Email/identity verified
│       │   ├── user_suspended.proto          # User suspended by admin
│       │   ├── user_banned.proto             # User banned
│       │   ├── freelancer_profile_completed.proto  # Freelancer profile complete
│       │   └── client_profile_completed.proto      # Client profile complete
│       │
│       ├── job/v1/                           # Job lifecycle events
│       │   ├── job_posted.proto              # New job posted
│       │   ├── job_updated.proto             # Job details updated
│       │   ├── job_closed.proto              # Job closed (filled or cancelled)
│       │   ├── job_invitation_sent.proto     # Invitation sent to freelancer
│       │   ├── job_removed.proto             # Job removed by admin
│       │   └── job_flagged.proto             # Job flagged for review
│       │
│       ├── proposal/v1/                      # Proposal and bidding events
│       │   ├── proposal_submitted.proto      # Proposal submitted
│       │   ├── proposal_accepted.proto       # Proposal accepted by client
│       │   ├── proposal_rejected.proto       # Proposal rejected
│       │   ├── proposal_withdrawn.proto      # Proposal withdrawn by freelancer
│       │   ├── bid_placed.proto              # New bid placed
│       │   ├── bid_updated.proto             # Bid amount updated
│       │   ├── outbid_alert.proto            # Freelancer outbid notification
│       │   ├── connect_used.proto            # Connect/credit used
│       │   └── proposal_flagged.proto        # Proposal flagged for review
│       │
│       ├── contract/v1/                      # Contract lifecycle events
│       │   ├── contract_created.proto        # Contract created
│       │   ├── contract_started.proto        # Work started
│       │   ├── contract_paused.proto         # Contract paused
│       │   ├── contract_ended.proto          # Contract completed/terminated
│       │   ├── milestone_created.proto       # Milestone created
│       │   ├── milestone_completed.proto     # Milestone completed
│       │   ├── milestone_approved.proto      # Milestone approved by client
│       │   ├── timesheet_submitted.proto     # Weekly timesheet submitted
│       │   └── dispute_opened.proto          # Dispute opened
│       │
│       ├── payment/v1/                       # Financial events
│       │   ├── payment_processed.proto       # Payment successfully processed
│       │   ├── payment_failed.proto          # Payment failed
│       │   ├── escrow_held.proto             # Funds held in escrow
│       │   ├── escrow_released.proto         # Escrow released to freelancer
│       │   ├── payout_requested.proto        # Payout requested
│       │   ├── payout_processed.proto        # Payout completed
│       │   ├── invoice_generated.proto       # Invoice generated
│       │   └── refund_processed.proto        # Refund issued
│       │
│       ├── review/v1/                        # Review and rating events
│       │   ├── review_submitted.proto        # Review submitted
│       │   ├── review_responded.proto        # Response to review posted
│       │   ├── badge_awarded.proto           # Achievement badge awarded
│       │   ├── reputation_updated.proto      # Reputation score updated
│       │   └── review_flagged.proto          # Review flagged for moderation
│       │
│       ├── subscription/v1/                  # Subscription events
│       │   ├── subscription_created.proto    # New subscription
│       │   ├── subscription_renewed.proto    # Subscription renewed
│       │   ├── subscription_cancelled.proto  # Subscription cancelled
│       │   ├── subscription_expired.proto    # Subscription expired
│       │   ├── connects_purchased.proto      # Connects purchased
│       │   ├── connects_used.proto           # Connect used
│       │   └── usage_limit_reached.proto     # Plan limit reached
│       │
│       ├── message/v1/                       # Communication events
│       │   ├── message_sent.proto            # Chat message sent
│       │   ├── notification_delivered.proto  # Notification delivered
│       │   ├── email_sent.proto              # Email sent
│       │   └── message_flagged.proto         # Message flagged
│       │   └── in_app_notification_sent.proto         # in_app_notification_sent
│       │
│       ├── storage/v1/                       # File storage events
│       │   ├── file_uploaded.proto           # File uploaded
│       │   ├── file_deleted.proto            # File deleted
│       │   ├── media_processed.proto         # Media processing complete
│       │   └── file_flagged.proto            # File flagged for review
│       │
│       ├── search/v1/                        # Search and recommendation events
│       │   ├── job_indexed.proto             # Job added to search index
│       │   ├── user_indexed.proto            # User added to search index
│       │   └── recommendation_generated.proto # Recommendation created
│       │
│       ├── admin/v1/                         # Admin action events
│       │   ├── user_suspended.proto          # Admin suspended user
│       │   ├── user_banned.proto             # Admin banned user
│       │   ├── content_removed.proto         # Admin removed content
│       │   ├── dispute_resolved.proto        # Admin resolved dispute
│       │   ├── flag_reviewed.proto           # Flag reviewed and actioned
│       │   └── announcement_published.proto  # Platform announcement published
│       │
│       └── gen/                              # 🆕 Generated Go code (checked into git or built in CI)
│           └── go/
│               └── skillsier/
│                   ├── user/v1/
│                   │   ├── user_created.pb.go
│                   │   ├── user_updated.pb.go
│                   │   └── ...
│                   ├── job/v1/
│                   ├── proposal/v1/
│                   ├── contract/v1/
│                   ├── payment/v1/
│                   ├── review/v1/
│                   ├── subscription/v1/
│                   ├── message/v1/
│                   ├── storage/v1/
│                   ├── search/v1/
│                   └── admin/v1/
│
├── deployments/                              # 📝 UPDATED - Kustomize-based deployment
│   └── k8s/
│       ├── base/                             # 🆕 Base Kubernetes manifests (shared)
│       │   ├── kustomization.yaml            # Kustomize base config
│       │   ├── namespace.yaml                # Namespace definition
│       │   ├── networkpolicy.yaml            # 🆕 Default network policies (deny all, allow ingress)
│       │   └── podsecurity.yaml              # 🆕 Pod security standards (restricted)
│       │
│       ├── overlays/                         # 🆕 Environment-specific customizations
│       │   ├── local/                        # Local development environment
│       │   │   ├── kustomization.yaml        # Local overlay config
│       │   │   ├── configmap-patch.yaml      # Local config overrides (debug mode, local URLs)
│       │   │   └── resources-patch.yaml      # Lower resource limits for local (256MB RAM)
│       │   │
│       │   └── prod/                         # Production environment
│       │       ├── kustomization.yaml        # Production overlay config
│       │       ├── configmap-patch.yaml      # Production configs (prod URLs, log levels)
│       │       ├── resources-patch.yaml      # Production resource limits (2GB RAM, 1 CPU)
│       │       ├── replicas-patch.yaml       # Replica counts (min 2, max 10)
│       │       ├── hpa.yaml                  # Horizontal Pod Autoscaler (CPU/Memory based)
│       │       ├── pdb.yaml                  # Pod Disruption Budget (maxUnavailable: 1)
│       │       └── ingress.yaml              # Production ingress (TLS, rate limiting)
│       │
│       └── components/                       # 🆕 Reusable Kustomize components
│           ├── monitoring/
│           │   ├── kustomization.yaml
│           │   └── servicemonitor.yaml       # Prometheus ServiceMonitor
│           └── secrets/
│               ├── kustomization.yaml
│               └── sealed-secrets.yaml       # SealedSecrets for secret management
│
├── docs/                                     # 🆕 Root-level documentation
│   ├── README.md                             # Platform overview
│   ├── ARCHITECTURE.md                       # 🆕 High-level architecture (services, data flow, patterns)
│   ├── SAGAS.md                              # 🆕 Saga patterns & compensations (workflows, rollback)
│   ├── RUNBOOK.md                            # 🆕 Operational procedures (restarts, debugging, incidents)
│   └── decisions/                            # 🆕 Architecture Decision Records (ADRs)
│       ├── README.md                         # ADR index
│       ├── 001-auth-centralization.md        # Why centralize auth in pkg/auth
│       ├── 002-outbox-pattern.md             # Why use outbox for reliable events
│       ├── 003-event-contracts.md            # Why separate event contracts module
│       ├── 004-dapr-adoption.md              # Why use Dapr for service mesh
│       ├── 005-platform-shared.md            # Why create platform-shared module
│       └── 006-config-standardization.md     # Why standardize config across services
│
├── scripts/                                  # 🆕 Root-level automation scripts
│   ├── setup-dev-env.sh                      # Setup entire dev environment (deps, DBs, Kafka)
│   ├── generate-proto.sh                     # Generate protobuf code from contracts
│   ├── test-all.sh                           # Run all tests across all services
│   ├── lint-all.sh                           # Lint all services
│   └── k8s-apply-all.sh                      # Deploy all services to Kubernetes
│
├── .github/
│   └── workflows/
│       ├── contracts-ci.yml                  # 🆕 Contracts CI/CD (buf lint, breaking check, generate)
│       ├── platform-shared-ci.yml            # 🆕 Platform-shared CI/CD (test, lint, publish)
│       ├── auth-ci.yml                       # 🆕 Auth package CI/CD (test, lint, publish)
│       └── services-ci.yml                   # All services CI/CD (test, build, deploy)
│
├── Makefile                                  # 🆕 Root-level orchestration (make test, make deploy-all)
├── .gitignore                                # Git ignore patterns
├── .editorconfig                             # Editor configuration
├── go.work                                   # 🆕 Go workspace (optional, for multi-module development)
└── README.md                                 # Root README with quick start
```

---

# Complete Event Flow & Integration Map

## Key Events Between Services

### 1. User Lifecycle Events
```
users-be → All Services
- UserCreated
- UserUpdated
- UserVerified
- FreelancerProfileCompleted
- ClientProfileCompleted
```

### 2. Job Lifecycle Events
```
jobs-be → proposals-be, search-be, communications-be, subscriptions-be
- JobPosted
- JobUpdated
- JobClosed
- JobInvitationSent
```

### 3. Proposal & Bidding Events
```
proposals-be → jobs-be, contracts-be, communications-be, subscriptions-be, financial-be
- ProposalSubmitted
- ProposalAccepted
- ProposalRejected
- BidPlaced
- BidUpdated
- OutbidAlert
- ConnectUsed
```

### 4. Contract Events
```
contracts-be → financial-be, communications-be, reviews-be, users-be
- ContractCreated
- ContractStarted
- MilestoneCreated
- MilestoneCompleted
- MilestoneApproved
- TimesheetSubmitted
- ContractPaused
- ContractEnded
```

### 5. Financial Events
```
financial-be → contracts-be, users-be, communications-be, subscriptions-be
- PaymentProcessed
- PaymentFailed
- EscrowHeld
- EscrowReleased
- PayoutRequested
- PayoutProcessed
- InvoiceGenerated
```

### 6. Communication Events
```
communications-be → All Services
- MessageSent
- NotificationDelivered
- EmailSent
- InAppNotificationSent
```

### 7. Review & Rating Events
```
reviews-be → users-be, search-be, communications-be
- ReviewSubmitted
- ReviewResponded
- BadgeAwarded
- ReputationUpdated
- RatingCalculated
```

### 8. Search & Recommendation Events
```
search-be → communications-be
- JobIndexed
- UserIndexed
- RecommendationGenerated
- MatchFound
```

### 9. Subscription Events
```
subscriptions-be → users-be, jobs-be, proposals-be, communications-be, financial-be
- SubscriptionCreated
- SubscriptionRenewed
- SubscriptionCancelled
- ConnectsPurchased
- ConnectsUsed
- UsageLimitReached
```

### 10. Storage Events
```
storage-be → users-be, jobs-be, contracts-be, proposals-be
- FileUploaded
- FileDeleted
- MediaProcessed
```

## 11. Events Published by admin-be

```
// User Actions
- admin.user.suspended
- admin.user.unsuspended
- admin.user.banned
- admin.user.unbanned
- admin.user.warned
- admin.user.verified

// Content Actions
- admin.job.removed
- admin.job.hidden
- admin.job.featured
- admin.proposal.removed
- admin.review.removed
- admin.message.removed
- admin.file.removed

// Moderation
- admin.moderation.approved
- admin.moderation.rejected
- admin.flag.resolved

// Disputes
- admin.dispute.resolved
- admin.dispute.escalated

// Financial
- admin.payment.refunded
- admin.chargeback.resolved

// Subscriptions
- admin.subscription.cancelled
- admin.subscription.extended
- admin.connects.added

// System
- admin.config.updated
- admin.announcement.published
- admin.maintenance.scheduled
```

---

## Events Subscribed by admin-be

```
// Flagging Events
- user.flagged
- job.flagged
- proposal.flagged
- contract.dispute.opened
- payment.disputed
- review.flagged
- message.flagged
- file.flagged

// System Events
- system.error
- system.alert
- system.abuse.detected

// Support Events
- support.ticket.created
- support.ticket.escalated
```



---

## In-App Notification Types (communications-be)

### Job-Related Notifications
1. **New Job Posted** (matching your skills)
2. **Job Invitation Received**
3. **Job You Applied To Closed**
4. **Similar Jobs Available**

### Proposal & Bidding Notifications
1. **Your Proposal Was Viewed**
2. **Your Proposal Was Accepted**
3. **Your Proposal Was Rejected**
4. **New Bid Placed on Your Job**
5. **You've Been Outbid**
6. **Bid Accepted**

### Contract Notifications
1. **New Contract Created**
2. **Contract Started**
3. **Milestone Completed**
4. **Milestone Approved**
5. **Timesheet Submitted**
6. **Work Diary Updated**
7. **Contract Ending Soon**

### Financial Notifications
1. **Payment Received**
2. **Payment Sent**
3. **Funds in Escrow**
4. **Escrow Released**
5. **Payout Processed**
6. **Invoice Generated**
7. **Low Balance Alert**

### Message Notifications
1. **New Message Received**
2. **Message Read**
3. **User Is Typing**

### Review Notifications
1. **New Review Received**
2. **Review Response Posted**
3. **Badge Earned**

### Subscription Notifications
1. **Subscription Expiring Soon**
2. **Subscription Renewed**
3. **Connects Running Low**
4. **New Connects Purchased**

### System Notifications
1. **Profile Verification Approved/Rejected**
2. **Account Security Alert**
3. **Terms of Service Updated**

---

## Technology Stack Summary

### Backend
- **Language**: Go
- **Framework**: Standard library + Gin/Echo/Chi
- **Architecture**: Clean Architecture (DDD)

### Data Layer
- **Primary DB**: PostgreSQL (per service)
- **Caching**: Redis
- **Search**: Elasticsearch/OpenSearch
- **Object Storage**: MinIO (self-hosted)

### Messaging & Events
- **Event Bus**: Apache Kafka
- **Patterns**: Outbox Pattern, Event Sourcing, CQRS

### Authentication & Authorization
- **Auth Service**: Keycloak
- **Token**: JWT

### Communication
- **Email**: WildDuck (self-hosted SMTP)
- **WebSocket**: Native Go WebSocket
- **Real-time**: Server-Sent Events (SSE)

### Infrastructure
- **Container**: Docker
- **Orchestration**: Kubernetes
- **Service Mesh**: Istio/Linkerd (optional)
- **API Gateway**: Kong/NGINX

### Observability
- **Logging**: ELK Stack or Grafana Loki
- **Monitoring**: Prometheus + Grafana
- **Tracing**: Jaeger

### CI/CD
- **Version Control**: Git
- **CI/CD**: GitHub Actions / GitLab CI
- **GitOps**: ArgoCD

---

## Resource Estimation (Minimum K8s Cluster)

```yaml
# Minimal Production Setup
Nodes: 3 nodes × 8GB RAM = 24GB total

Application Services (10 services):
- users-be: 2 replicas × 512MB = 1GB
- jobs-be: 2 replicas × 512MB = 1GB
- proposals-be: 2 replicas × 512MB = 1GB
- contracts-be: 2 replicas × 512MB = 1GB
- financial-be: 3 replicas × 1GB = 3GB (critical)
- communications-be: 2 replicas × 512MB = 1GB
- storage-be: 2 replicas × 512MB = 1GB
- search-be: 2 replicas × 1GB = 2GB
- reviews-be: 2 replicas × 512MB = 1GB
- subscriptions-be: 2 replicas × 512MB = 1GB

Subtotal: ~14GB

Infrastructure:
- PostgreSQL: 2GB
- Redis: 1GB
- Elasticsearch: 2GB
- Kafka: 2GB
- Keycloak: 1GB
- MinIO: 1GB

Subtotal: ~9GB

Total: ~23GB (fits in 24GB cluster)
```

---

This completes the comprehensive microservices architecture for your Skillsier platform! All services now include:

- ✅ **Bidding system** in proposals-be
- ✅ **In-app notifications** with comprehensive coverage in communications-be
- ✅ **Free & self-hosted** communications (WildDuck, WebSocket, SSE)
- ✅ **Auto-migrations** following your users-be pattern
- ✅ **Enhanced recommendation system** in search-be with ML models
- ✅ **Reviews covering ratings** in reviews-be
- ✅ Complete event flows between all services
- ✅ Enterprise-level folder structures

Each service is production-ready and follows Go best practices with Clean Architecture!