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

## 📦 **1️⃣ users-be (User Management Service)**

```

apps/be/users-be/
│
├── cmd/
│   ├── api/
│   │   └── main.go                           # 📝 API entrypoint - initializes Gin, Dapr, Postgres (uses platform-shared/logging, internal/config)
│   │                                         # ensure outbox dispatcher & leader election wired here (see infra/coordination)
│   │
│   └── worker/
│       └── main.go                           # 🧵 Worker entrypoint - bootstraps DI, registers cron tasks, runs background jobs
│
├── internal/
│   │
│   ├── config/                               # 🔧 Configuration (Load First)
│   │   ├── feature_flags.go                🆕  # central toggles (enable_pro_network, enable_learning_path, etc.)
│   │   ├── schema.go                         # Typed Config struct (App, Server, Postgres, Kafka, Redis, Auth, Storage, Keycloak), ♻️  # group flags under Config.FeatureFlags
│   │   ├── loader.go                         # Config loader using Viper (CLI → ENV → file → defaults)
│   │   └── docs/
│   │       └── CONFIGURATION.md              # Configuration documentation
│   │
│   ├── ioc/                                  # 🧩 Dependency Injection Container
│   │   ├── container.go                      # DI graph: constructs DB/Redis/Kafka clients, repositories, services, handlers, schedulers
│   │   └── wiring.go                         # Env-driven wiring & feature flags: selects implementations (local vs cloud), ♻️  # wire outbox, coordination, observers based on env
│   │
│   ├── infrastructure/                       # 🔧 Infrastructure Layer (Load Second)
│   │   │
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection setup
│   │   │       ├── transaction.go            # Transaction helpers , ♻️  # ensure outbox write is in same DB tx
│   │   │       ├── migrations.go             # Auto-migration logic with version tracking
│   │   │       ├── version.go                # Schema version tracking
│   │   │       ├── safety.go                 # Pre-migration safety checks
│   │   │       ├── outbox_store.go         🆕  # tx-coupled outbox table access
│   │   │       │
│   │   │       # ========================= CORE USER REPOSITORIES =========================
│   │   │       ├── user_repository.go        # User CRUD (Create, Update, FindByID, FindByEmail, List, Search)
│   │   │       ├── profile_repository.go     # Profile repository
│   │   │       │
│   │   │       # ========================= SKILLS & CAPABILITIES REPOSITORIES =========================
│   │   │       ├── capabilities_repository.go # 🔄 CONSOLIDATED: skills + specializations + taxonomy
│   │   │       │                             # Stores: skills, proficiency, specializations, niche expertise, taxonomy mapping
│   │   │       ├── service_catalog_repository.go # Service catalog (references capability IDs)
│   │   │       │
│   │   │       # ========================= EXPERIENCE & EDUCATION REPOSITORIES =========================
│   │   │       ├── experience_repository.go  # Work experience repository
│   │   │       ├── education_repository.go   # Education repository
│   │   │       ├── language_repository.go    # Language proficiency repository
│   │   │       │
│   │   │       # ========================= CREDENTIALS REPOSITORIES (CONSOLIDATED) =========================
│   │   │       ├── credentials_repository.go # 🔄 CONSOLIDATED: external_certifications + platform_certifications
│   │   │       │                             # Stores: external certs (AWS, Google) + platform certs (Upwork badges)
│   │   │       │
│   │   │       # ========================= PORTFOLIO & SHOWCASE REPOSITORIES =========================
│   │   │       ├── portfolio_repository.go   # Portfolio items repository
│   │   │       │
│   │   │       # ========================= USER TYPE SPECIFIC REPOSITORIES =========================
│   │   │       ├── freelancer_repository.go  # Freelancer-specific data repository
│   │   │       ├── client_repository.go      # Client-specific data repository
│   │   │       │
│   │   │       # ========================= IDENTITY VERIFICATION REPOSITORIES (CONSOLIDATED) =========================
│   │   │       ├── identity_verification_repository.go # 🔄 CONSOLIDATED: KYC/KYB ONLY (was verification)
│   │   │       │                             # Stores: ID verification, passport, address proof, selfie
│   │   │       │
│   │   │       # ========================= TRUST REPOSITORIES =========================
│   │   │       ├── trust_repository.go       # 🔄 RENAMED from trust_verification
│   │   │       │                             # Stores: trust level, trust badges (VerifiedPayment, IDVerified)
│   │   │       │
│   │   │       # ========================= BADGING REPOSITORIES (CONSOLIDATED) =========================
│   │   │       ├── badging_repository.go     # 🔄 CONSOLIDATED: all badge issuance (achievement, certification, trust, platform)
│   │   │       │                             # Single source of truth for ALL badges
│   │   │       │
│   │   │       # ========================= SETTINGS & PREFERENCES REPOSITORIES =========================
│   │   │       ├── settings_repository.go    # User settings repository
│   │   │       ├── privacy_repository.go     # 🔄 SPLIT from settings - privacy settings ONLY (ShowEmail, ShowPhone, ShareActivity)
│   │   │       ├── saved_items_repository.go # Saved items repository
│   │   │       ├── blocked_users_repository.go # Blocked users repository
│   │   │       │
│   │   │       # ========================= MODERATION REPOSITORIES (CONSOLIDATED) =========================
│   │   │       ├── moderation_repository.go  # 🔄 CONSOLIDATED: suspension + ban + warning
│   │   │       │                             # Single repository for all moderation actions
│   │   │       │
│   │   │       # ========================= SCORING REPOSITORIES (CONSOLIDATED) =========================
│   │   │       ├── user_metrics_repository.go # 🔄 SINGLE SOURCE: raw metrics (response time, completion rate, etc.)
│   │   │       │                             # All scoring contexts read from here
│   │   │       ├── reputation_repository.go  # Market-facing reputation (computed from user_metrics)
│   │   │       ├── quality_repository.go     # Work quality score (computed from user_metrics)
│   │   │       ├── account_health_repository.go # Operational health (computed from user_metrics)
│   │   │       ├── risk_repository.go        # Fraud/safety signals (computed from user_metrics)
│   │   │       │
│   │   │       # ========================= ORGANIZATION & TEAM REPOSITORIES =========================
│   │   │       ├── org_repository.go         # 🔄 SINGLE SOURCE for company/org data
│   │   │       │                             # Client references OrgID (no duplicate company fields)
│   │   │       │
│   │   │       # ========================= SECURITY REPOSITORIES (CONSOLIDATED) =========================
│   │   │       ├── security_repository.go    # 🔄 CONSOLIDATED: 2FA + devices + sessions + recovery
│   │   │       │                             # Single repository for all security features
│   │   │       │
│   │   │       # ========================= PROFILE ENHANCEMENT REPOSITORIES (CONSOLIDATED) =========================
│   │   │       ├── profile_depth_repository.go # 🔄 CONSOLIDATED: rate history + normalized skills taxonomy
│   │   │       │                             # NO badges (moved to badging), NO availability (moved to availability)
│   │   │       ├── profile_completeness_repository.go # Completeness score ONLY (NO view counts, NO badges)
│   │   │       ├── profile_analytics_repository.go # Views/search ONLY (NO completeness, NO badges)
│   │   │       ├── profile_optimization_repository.go # AI optimization suggestions
│   │   │       │
│   │   │       # ========================= PROFILE VISIBILITY REPOSITORIES (CONSOLIDATED) =========================
│   │   │       ├── profile_visibility_repository.go # 🔄 Search/discoverability ONLY
│   │   │       │                             # NO privacy flags (moved to privacy_repository)
│   │   │       │
│   │   │       # ========================= AVAILABILITY REPOSITORIES (CONSOLIDATED) =========================
│   │   │       ├── availability_repository.go # 🔄 SINGLE SOURCE: status + calendar + recurring rules + vacation
│   │   │       │                             # Profile/profile_depth reference this, NO duplicate fields
│   │   │       │
│   │   │       # ========================= NETWORKING & CONNECTIONS REPOSITORIES =========================
│   │   │       ├── professional_network_repository.go # Professional networking
│   │   │       ├── referrals_repository.go   # Referral system
│   │   │       ├── user_groups_repository.go # Community groups
│   │   │       │
│   │   │       # ========================= FINANCIAL PROFILE REPOSITORIES =========================
│   │   │       ├── payment_methods_repository.go # Payment methods management
│   │   │       ├── financial_profile_repository.go # Financial profile and preferences
│   │   │       ├── earning_goals_repository.go # Earning goals tracking
│   │   │       │
│   │   │       # ========================= PROFESSIONAL DEVELOPMENT REPOSITORIES =========================
│   │   │       ├── learning_path_repository.go # Learning paths
│   │   │       ├── mentorship_repository.go  # Mentorship program
│   │   │       ├── achievements_repository.go # Gamified achievements (emits events → badging owns issuance)
│   │   │       │
│   │   │       # ========================= COMPLIANCE REPOSITORIES =========================
│   │   │       ├── compliance_repository.go  # Tax/residency artifacts
│   │   │       │
│   │   │       # ========================= COMMUNICATION PREFERENCES REPOSITORIES =========================
│   │   │       ├── communication_channels_repository.go # Multi-channel communication
│   │   │       └── email_preferences_repository.go # Email preferences
│   │   │
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go             # Redis connection setup
│   │   │       ├── keys.go                   # 🆕 Canonical cache keys & TTLs (no magic strings)
│   │   │       │
│   │   │       # ========================= CORE CACHING =========================
│   │   │       ├── user_cache.go             # User caching
│   │   │       ├── profile_cache.go          # Profile caching
│   │   │       ├── capabilities_cache.go     # 🔄 Capabilities caching
│   │   │       ├── org_cache.go              # Org membership/roles caching
│   │   │       │
│   │   │       # ========================= SCORING CACHING (CONSOLIDATED) =========================
│   │   │       ├── user_metrics_cache.go     # 🔄 Raw metrics cache (single source)
│   │   │       ├── reputation_cache.go       # Reputation score cache
│   │   │       ├── quality_cache.go          # Quality score cache
│   │   │       ├── account_health_cache.go   # Health score cache
│   │   │       ├── risk_cache.go             # Risk score cache
│   │   │       │
│   │   │       # ========================= OTHER CACHING =========================
│   │   │       ├── security_cache.go         # Session/device cache
│   │   │       ├── availability_cache.go     # 🔄 Availability cache
│   │   │       ├── badging_cache.go          # 🔄 Badge cache
│   │   │       ├── profile_analytics_cache.go # Profile analytics cache
│   │   │       └── network_cache.go          # Professional network cache
│   │   │       # =========================  =========================
│   │   │       ├── singleflight.go         🆕  # stampede protection for hot keys
│   │   │       └── invalidation_rules.go   🆕  # map events → keys to drop (doc alongside)
│   │   │
│   │   │
│   │   ├── messaging/
│   │   │   │── topics.go                   ♻️  # thin re-export from contracts/events with per-context keys - # ✅ Single source (re-export contracts/events)
│   │   │   ├── outbox/                     🆕  # exactly-once-ish publisher
│   │   │   │   ├── dispatcher.go           🆕  # reads outbox → Kafka (with retries/DLQ)
│   │   │   │   └── metrics.go              🆕  # publish lag, failures, retries
│   │   │   │
│   │   │   └── kafka/
│   │   │       ├── consumer.go               # 📝 Kafka consumer (uses platform-shared/inbox)
│   │   │       ├── producer.go               # 📝 Kafka producer (uses platform-shared/outbox)
│   │   │       │
│   │   │       # ========================= TOPIC DEFINITIONS =========================
│   │   │       ├── topics_user.go            # User lifecycle topics
│   │   │       ├── topics_scoring.go         # 🔄 Scoring topics (reputation, quality, health, risk)
│   │   │       ├── topics_badging.go         # 🔄 Badging topics
│   │   │       ├── topics_moderation.go      # 🔄 Moderation topics
│   │   │       │
│   │   │       # ========================= PRODUCERS =========================
│   │   │       ├── producers_user.go         # User event producers
│   │   │       ├── producers_scoring.go      # 🔄 Scoring event producers
│   │   │       ├── producers_badging.go      # 🔄 Badging event producers
│   │   │       ├── producers_moderation.go   # 🔄 Moderation event producers
│   │   │       │
│   │   │       └── scram.go                  # SCRAM authentication for Kafka
│   │   │
│   │   ├── http/
│   │   │   ├── idempotency_adapter.go        # 🆕 Bind platform-shared idempotency to Gin
│   │   │   └── etag.go                      🆕  # middleware-level ETag (pairs with utils/etag)
│   │   │
│   │   ├── messaging_adapters/
│   │   │   ├── inbox/
│   │   │   │   └── dapr_subscriptions.yaml   # 🆕 Service-scoped Dapr subscriptions
│   │   │   └── outbox/
│   │   │       └── publisher.go              # 🆕 Thin wrapper over platform-shared outbox
│   │   │
│   │   ├── storage/
│   │   │   ├── client.go                     # Storage service client (upload profile pics, portfolios, documents)
│   │   │   └── local.go                      # Local file storage fallback
│   │   │
│   │   ├── keycloak/
│   │   │   ├── client.go                     # 📝 Keycloak API client (uses pkg/auth)
│   │   │   └── sync.go                       # 📝 Sync users with Keycloak
│   │   │
│   │   ├── scheduler/                        # 🆕 Cron/Task Scheduler
│   │   │   ├── cron.go                       # ⏱️ Cron registry: schedules, idempotency guards, safe shutdown,  ♻️ # wrap each task with distlock + jitter
│   │   │   └── tasks/
│   │   │       ├── metrics_aggregation_task.go # 🆕 Aggregate raw metrics for scoring
│   │   │       ├── reputation_recalc_task.go # 🆕 Recalculate reputation scores
│   │   │       ├── quality_recalc_task.go    # 🆕 Recalculate quality scores
│   │   │       ├── health_recalc_task.go     # 🆕 Recalculate health scores
│   │   │       ├── risk_recalc_task.go       # 🆕 Recalculate risk scores
│   │   │       ├── badge_audit_task.go       # 🆕 Audit badge eligibility
│   │   │       └── profile_completeness_task.go # 🆕 Recalculate completeness scores
│   │   │       └── task_guard.go           🆕  # idempotency tokens + last-run watermark
│   │   │
│   │   ├── projections/                      # 🆕 CQRS Projections
│   │   │   ├── user_metrics/
│   │   │   │   ├── projector.go              # 🧮 Event projector: builds user_metrics read model
│   │   │   │   └── readmodel_repository.go   # 📖 Read-model repository for fast queries
│   │   │   │
│   │   │   ├── reputation/
│   │   │   │   ├── projector.go              # Consumes events → builds reputation read model
│   │   │   │   └── readmodel_repository.go
│   │   │   │
│   │   │   ├── quality/
│   │   │   │   ├── projector.go              # Consumes events → builds quality read model
│   │   │   │   └── readmodel_repository.go
│   │   │   │
│   │   │   └── profile_analytics/
│   │   │       ├── projector.go              # Consumes ProfileViewed events → builds analytics
│   │   │       └── readmodel_repository.go
│   │   ├── coordination/                   🆕  # distributed coordination primitives
│   │   │   ├── leader_election.go          🆕  # single-active cron/worker
│   │   │   └── distlock.go                 🆕  # Redis/PG advisory locks for cron tasks
│   │   │
│   │   │
│   │   ├── security/                         # 🆕 Security Utilities
│   │   │   ├── kms.go                        # 🔐 KMS wrapper: envelope-encrypt sensitive fields
│   │   │   └── pii_redactor.go               # 🧽 PII redactor: masks emails/phones in logs
│   │   │
│   │   └── ai/
│   │       ├── profile_optimizer_client.go   # AI client for profile optimization
│   │       └── skill_matcher_client.go       # AI client for skill matching
│   │
│   ├── domain/                               # 🏛️ Domain Layer - Business logic & entities (Load Third)
│   │   │
│   │   # ========================= CORE USER DOMAIN =========================
│   │   ├── user/
│   │   │   ├── entity.go                     # User aggregate root (ID, Username, Email, FirstName, LastName, UserType, Status)
│   │   │   ├── value_objects.go              # Email, Phone, Address value objects
│   │   │   ├── enums.go                      # UserType (Freelancer, Client), AccountStatus (Active, Suspended, Banned)
│   │   │   ├── errors.go                     # Domain-specific errors (UserNotFound, EmailTaken, InvalidEmail)
│   │   │   ├── repository.go                 # User repository interface
│   │   │   ├── list_filter.go                # ListFilter for queries
│   │   │   ├── statistics.go                 # UserStatistics
│   │   │   └── events.go                     # UserCreated, UserUpdated, UserVerified, UserSuspended, UserBanned
│   │   │
│   │   ├── profile/
│   │   │   ├── entity.go                     # Extended profile (Bio, Location, ProfilePictureURL)
│   │   │   │                                 # References AvailabilityID (NO availability fields here)
│   │   │   ├── preferences.go                # User preferences (Language, Timezone, Currency)
│   │   │   ├── errors.go                     # Profile-specific errors
│   │   │   ├── repository.go                 # Profile repository interface
│   │   │   └── events.go                     # ProfileUpdated, PreferencesUpdated
│   │   │
│   │   # ========================= CAPABILITIES DOMAIN (CONSOLIDATED) =========================
│   │   ├── capabilities/                     # 🔄 CONSOLIDATED: skills + specializations + taxonomy
│   │   │   ├── entity.go                     # Capabilities aggregate (UserID, Skills[], Specializations[])
│   │   │   │
│   │   │   ├── skills/
│   │   │   │   ├── skill.go                  # Skill (ID, Name, Proficiency, YearsOfExperience)
│   │   │   │   ├── proficiency.go            # Skill level enum (Beginner, Intermediate, Advanced, Expert)
│   │   │   │   └── taxonomy.go               # Standardized taxonomy (React → WebDev → Engineering)
│   │   │   │
│   │   │   ├── specializations/
│   │   │   │   ├── specialization.go         # Specialization (Skills[], Industry, Verified)
│   │   │   │   └── niche.go                  # Niche expertise (e.g., "React + TypeScript for FinTech")
│   │   │   │
│   │   │   ├── errors.go                     # Capabilities-specific errors
│   │   │   ├── repository.go                 # Capabilities repository interface
│   │   │   └── events.go                     # SkillAdded, SkillUpdated, SpecializationAdded, SpecializationVerified
│   │   │
│   │   ├── service_catalog/
│   │   │   ├── entity.go                     # ServiceCatalog (UserID, Services[], Packages[])
│   │   │   ├── service.go                    # Service (Name, Description, CapabilityIDs[], Price, Duration)
│   │   │   │                                 # References capabilities (NO duplicate skill data)
│   │   │   ├── service_packages.go           # Package services (Basic, Standard, Premium)
│   │   │   ├── errors.go                     # Service catalog errors
│   │   │   ├── repository.go                 # Service catalog repository interface
│   │   │   └── events.go                     # ServiceAdded, PackageCreated
│   │   │
│   │   # ========================= EXPERIENCE & EDUCATION DOMAIN =========================
│   │   ├── experience/
│   │   │   ├── entity.go                     # Work experience (Company, Title, Description, StartDate, EndDate, IsCurrent)
│   │   │   ├── errors.go                     # Experience-specific errors
│   │   │   ├── repository.go                 # Experience repository interface
│   │   │   └── events.go                     # ExperienceAdded, ExperienceUpdated, ExperienceRemoved
│   │   │
│   │   ├── education/
│   │   │   ├── entity.go                     # Educational background (School, Degree, Field, GraduationYear, Description)
│   │   │   ├── errors.go                     # Education-specific errors
│   │   │   ├── repository.go                 # Education repository interface
│   │   │   └── events.go                     # EducationAdded, EducationUpdated, EducationRemoved
│   │   │
│   │   ├── language/
│   │   │   ├── entity.go                     # Language proficiency (UserID, Language, ProficiencyLevel)
│   │   │   ├── errors.go                     # Language-specific errors
│   │   │   ├── repository.go                 # Language repository interface
│   │   │   └── events.go                     # LanguageAdded, LanguageUpdated, LanguageRemoved
│   │   │
│   │   # ========================= CREDENTIALS DOMAIN (CONSOLIDATED) =========================
│   │   ├── credentials/                      # 🔄 CONSOLIDATED: external + platform certifications
│   │   │   ├── entity.go                     # Credentials aggregate (UserID, ExternalCerts[], PlatformCerts[])
│   │   │   │
│   │   │   ├── external_certifications/
│   │   │   │   ├── certification.go          # External certifications (AWS, Google, Microsoft, etc.)
│   │   │   │   ├── verification.go           # Verification status (Pending, Verified, Rejected, Expired)
│   │   │   │   └── document.go               # Certification documents/credentials
│   │   │   │
│   │   │   ├── platform_certifications/
│   │   │   │   ├── certification.go          # Platform-issued certifications (Upwork badges)
│   │   │   │   ├── exam.go                   # Certification exams
│   │   │   │   └── recertification.go        # Periodic recertification
│   │   │   │
│   │   │   ├── errors.go                     # Credentials-specific errors
│   │   │   ├── repository.go                 # Credentials repository interface
│   │   │   └── events.go                     # ExternalCertificationAdded, ExternalCertificationVerified, PlatformCertificationEarned
│   │   │
│   │   # ========================= PORTFOLIO & SHOWCASE DOMAIN =========================
│   │   ├── portfolio/
│   │   │   ├── entity.go                     # Portfolio items (UserID, Title, Description, URL, ThumbnailURL, DisplayOrder)
│   │   │   ├── item.go                       # Portfolio item details
│   │   │   ├── media.go                      # Associated media (Images, Videos, Documents)
│   │   │   ├── errors.go                     # Portfolio-specific errors
│   │   │   ├── repository.go                 # Portfolio repository interface
│   │   │   └── events.go                     # PortfolioItemAdded, PortfolioItemUpdated, PortfolioItemRemoved
│   │   │
│   │   # ========================= USER TYPE SPECIFIC DOMAINS =========================
│   │   ├── freelancer/
│   │   │   ├── entity.go                     # Freelancer-specific data (UserID, Title, Overview, VideoIntroURL)
│   │   │   ├── profile.go                    # Freelancer profile details
│   │   │   ├── rates.go                      # Hourly/fixed rates (HourlyRate, MinimumBudget, Currency)
│   │   │   ├── stats.go                      # Job stats, earnings (TotalJobs, TotalEarnings, SuccessRate)
│   │   │   ├── errors.go                     # Freelancer-specific errors
│   │   │   ├── repository.go                 # Freelancer repository interface
│   │   │   └── events.go                     # FreelancerProfileUpdated, RatesUpdated, StatsUpdated
│   │   │
│   │   ├── client/
│   │   │   ├── entity.go                     # Client-specific data (UserID, OrgID)
│   │   │   │                                 # 🔄 References OrgID (NO duplicate company fields)
│   │   │   ├── stats.go                      # Hiring stats (TotalHires, TotalSpent, ActiveContracts)
│   │   │   ├── errors.go                     # Client-specific errors
│   │   │   ├── repository.go                 # Client repository interface
│   │   │   └── events.go                     # ClientProfileUpdated, ClientStatsUpdated
│   │   │
│   │   # ========================= IDENTITY VERIFICATION DOMAIN (CONSOLIDATED) =========================
│   │   ├── identity_verification/            # 🔄 RENAMED from verification - KYC/KYB ONLY
│   │   │   ├── entity.go                     # IdentityVerification (UserID, Type, Status, SubmittedAt, VerifiedAt)
│   │   │   ├── document.go                   # Verification documents (ID, Passport, ProofOfAddress, Selfie)
│   │   │   ├── verification_type.go          # Type: KYC (individual), KYB (business)
│   │   │   ├── errors.go                     # Verification-specific errors
│   │   │   ├── repository.go                 # IdentityVerification repository interface
│   │   │   └── events.go                     # IdentityVerificationSubmitted, IdentityVerificationApproved, IdentityVerificationRejected
│   │   │
│   │   # ========================= TRUST DOMAIN (CONSOLIDATED) =========================
│   │   ├── trust/                            # 🔄 RENAMED from trust_verification
│   │   │   ├── entity.go                     # Trust (UserID, TrustLevel, TrustBadges[], LastUpdated)
│   │   │   ├── trust_level.go                # TrustLevel enum (Unverified, Basic, Enhanced, Premium)
│   │   │   ├── badge_refs.go                 # 🆕 reference types from badging (remove duplicate trust_badge.go here) - # ✅ References badging/types/trust_badge.go
│   │   │   ├── errors.go                     # Trust-specific errors
│   │   │   ├── repository.go                 # Trust repository interface
│   │   │   └── events.go                     # TrustLevelUpgraded, TrustBadgeEarned, TrustBadgeRevoked
│   │   │
│   │   # ========================= BADGING DOMAIN (CONSOLIDATED) =========================
│   │   ├── badging/                          # 🔄 CONSOLIDATED: ALL badge issuance (single source of truth)
│   │   │   ├── entity.go                     # Badge aggregate (ID, Type, UserID, Slug, AwardedAt, RevokedAt)
│   │   │   ├── badge_type.go                 # BadgeType enum (Achievement, Certification, Trust, Platform)
│   │   │   │
│   │   │   ├── achievement_badge.go          # Achievement badges (FirstJob, TopRated, QuickResponder, etc.)
│   │   │   ├── certification_badge.go        # Certification badges (from credentials domain)
│   │   │   ├── trust_badge.go                # Trust badges (from trust domain)
│   │   │   └── types/
│   │   │   │    └── trust_badge.go          🔄  # canonical trust badge type lives here
│   │   │   │
│   │   │   ├── platform_badge.go             # Platform badges (RisingTalent, TopRated, ExpertVetted)
│   │   │   │
│   │   │   ├── errors.go                     # Badging-specific errors
│   │   │   ├── repository.go                 # Badging repository interface
│   │   │   └── events.go                     # BadgeAwarded, BadgeRevoked
│   │   │   │
│   │   │   # NOTE: Other contexts emit events, badging consumes and issues badges
│   │   │   # Example: achievements/ emits AchievementUnlocked → badging/ awards achievement badge
│   │   │
│   │   # ========================= SETTINGS & PREFERENCES DOMAINS =========================
│   │   ├── settings/
│   │   │   ├── entity.go                     # User settings (UserID, Settings JSON, Theme, Language, Timezone)
│   │   │   ├── notification_prefs.go         # Notification preferences (Email, SMS, Push, InApp)
│   │   │   ├── errors.go                     # Settings-specific errors
│   │   │   ├── repository.go                 # Settings repository interface
│   │   │   └── events.go                     # SettingsUpdated, NotificationPrefsUpdated
│   │   │
│   │   ├── privacy/                          # 🔄 SPLIT from settings - privacy ONLY
│   │   │   ├── entity.go                     # PrivacySettings (UserID, ShowEmail, ShowPhone, ShareActivity, AllowDirectContact)
│   │   │   ├── errors.go                     # Privacy-specific errors
│   │   │   ├── repository.go                 # Privacy repository interface
│   │   │   └── events.go                     # PrivacySettingsUpdated
│   │   │
│   │   ├── saved_items/
│   │   │   ├── entity.go                     # Saved jobs, freelancers (UserID, ItemType, ItemID, SavedAt, Notes)
│   │   │   ├── errors.go                     # Saved items-specific errors
│   │   │   ├── repository.go                 # Saved items repository interface
│   │   │   └── events.go                     # ItemSaved, ItemUnsaved
│   │   │
│   │   ├── blocked_users/
│   │   │   ├── entity.go                     # Blocked user relationships (BlockerID, BlockedID, BlockedAt, Reason)
│   │   │   ├── errors.go                     # Blocking-specific errors
│   │   │   ├── repository.go                 # Blocked users repository interface
│   │   │   └── events.go                     # UserBlocked, UserUnblocked
│   │   │
│   │   # ========================= MODERATION DOMAIN (CONSOLIDATED) =========================
│   │   ├── moderation/                       # 🔄 CONSOLIDATED: suspension + ban + warning
│   │   │   ├── entity.go                     # Moderation aggregate (UserID, Actions[], ActiveStatus)
│   │   │   │
│   │   │   ├── suspension/
│   │   │   │   ├── suspension.go             # Suspension (Reason, StartDate, EndDate, SuspendedBy, IsActive)
│   │   │   │   ├── reason.go                 # Suspension reasons enum (TOSViolation, PaymentIssue, QualityIssues, AbusiveBehavior)
│   │   │   │   └── duration.go               # Suspension duration (Days, Weeks, Months, Permanent)
│   │   │   │
│   │   │   ├── ban/
│   │   │   │   ├── ban.go                    # Ban (Reason, BannedAt, BannedBy, IsPermanent, ExpiresAt)
│   │   │   │   ├── reason.go                 # Ban reasons enum (Fraud, SevereAbuse, MultipleViolations, SecurityThreat)
│   │   │   │   └── permanent.go              # Permanent vs temporary flag
│   │   │   │
│   │   │   ├── warning/
│   │   │   │   ├── warning.go                # Warning (Reason, IssuedAt, IssuedBy, AcknowledgedAt)
│   │   │   │   ├── reason.go                 # Warning reasons enum (LateDelivery, PoorQuality, UnresponsiveCommunication)
│   │   │   │   └── severity.go               # Warning severity level (Low, Medium, High, Critical)
│   │   │   │
│   │   │   ├── shared/
│   │   │   │   ├── moderation_reason.go      # 🔄 Shared reason enums (no duplication)
│   │   │   │   └── moderation_actor.go       # Who took action (AdminID, System)
│   │   │   │
│   │   │   ├── errors.go                     # Moderation-specific errors
│   │   │   ├── repository.go                 # Moderation repository interface (single repo for all actions)
│   │   │   └── events.go                     # SuspensionPlaced, SuspensionReleased, BanPlaced, BanReleased, WarningIssued, WarningAcknowledged
│   │   │
│   │   # ========================= SCORING DOMAINS (CONSOLIDATED) =========================
│   │   ├── user_metrics/                     # 🔄 SINGLE SOURCE OF TRUTH for raw metrics
│   │   │   ├── entity.go                     # UserMetrics (UserID, ResponseTime, CompletionRate, ClientSatisfaction, etc.)
│   │   │   │                                 # ALL scoring contexts read from here
│   │   │   ├── errors.go                     # Metrics-specific errors
│   │   │   ├── repository.go                 # UserMetrics repository interface
│   │   │   └── events.go                     # MetricsUpdated, MetricRecorded
│   │   │
│   │   ├── reputation/                       # Market-facing reputation (computed from user_metrics)
│   │   │   ├── entity.go                     # Reputation (UserID, ReputationScore, Components[], History[])
│   │   │   │                                 # Uses pkg/scoring to compute from user_metrics
│   │   │   ├── reputation_components.go      # Components: Reviews, Completion, Response, Quality
│   │   │   ├── errors.go                     # Reputation-specific errors
│   │   │   ├── repository.go                 # Reputation repository interface
│   │   │   └── events.go                     # ReputationUpdated, ReputationScoreChanged
│   │   │
│   │   ├── quality/                          # Work quality (computed from user_metrics)
│   │   │   ├── entity.go                     # QualityScore (UserID, Score, Factors[], TrendAnalysis)
│   │   │   │                                 # Uses pkg/scoring to compute from user_metrics
│   │   │   ├── scoring_factors.go            # Factors: CompletionRate, ResponseTime, ClientSatisfaction, WorkQuality
│   │   │   ├── errors.go                     # Quality-specific errors
│   │   │   ├── repository.go                 # Quality repository interface
│   │   │   └── events.go                     # QualityScoreUpdated, QualityImproved, QualityDeclined
│   │   │
│   │   ├── account_health/                   # Operational health (computed from user_metrics)
│   │   │   ├── entity.go                     # AccountHealth (UserID, HealthScore, Issues[], Recommendations[])
│   │   │   │                                 # Uses pkg/scoring to compute from user_metrics
│   │   │   ├── health_factors.go             # Factors: ProfileComplete, Active, Responsive, QualityWork
│   │   │   ├── errors.go                     # Health-specific errors
│   │   │   ├── repository.go                 # AccountHealth repository interface
│   │   │   └── events.go                     # HealthScoreUpdated, IssueDetected, HealthImproved
│   │   │
│   │   ├── risk/                             # Fraud/safety (computed from user_metrics + signals)
│   │   │   ├── entity.go                     # Risk (UserID, RiskScore, Signals[], Holds[])
│   │   │   │                                 # Uses pkg/scoring to compute from user_metrics + signals
│   │   │   ├── signal.go                     # Signals (type: ip_geo_mismatch, disputes, chargebacks; severity; occurredAt)
│   │   │   ├── hold.go                       # Account holds (type, reason, actor, until)
│   │   │   ├── errors.go                     # Risk-specific errors
│   │   │   ├── repository.go                 # Risk repository interface
│   │   │   └── events.go                     # RiskSignalRecorded, RiskScoreUpdated, RiskHoldPlaced, RiskHoldReleased
│   │   │
│   │   # ========================= ORGANIZATION & TEAM DOMAIN =========================
│   │   ├── org/                              # 🔄 SINGLE SOURCE for company/org data
│   │   │   ├── entity.go                     # Org aggregate (OrgID, Name, Industry, Size, Founded, Employees, BillingProfileID, SeatLimit)
│   │   │   │                                 # Client references this (NO duplicate company fields)
│   │   │   ├── member.go                     # Membership (OrgID, UserID, Role[owner/admin/member], JoinedAt)
│   │   │   ├── seat.go                       # Seats usage (Total, Used, PendingInvites)
│   │   │   ├── errors.go                     # Org-specific errors
│   │   │   ├── repository.go                 # Org repository interface
│   │   │   └── events.go                     # OrgCreated, OrgUpdated, OrgMemberAdded, OrgMemberRemoved, SeatUpdated
│   │   │
│   │   # ========================= SECURITY DOMAIN (CONSOLIDATED) =========================
│   │   ├── security/                         # 🔄 CONSOLIDATED: 2FA + devices + sessions + recovery
│   │   │   ├── entity.go                     # Security aggregate (UserID, TwoFAEnabled, RecoveryKeys, Devices[], Sessions[])
│   │   │   │
│   │   │   ├── two_factor/
│   │   │   │   ├── two_factor.go             # 2FA settings (Enabled, Method, Secret, BackupCodes)
│   │   │   │   └── method.go                 # Method enum (TOTP, SMS, Email)
│   │   │   │
│   │   │   ├── devices/
│   │   │   │   ├── device.go                 # Registered devices (DeviceID, Fingerprint, LastSeen, Revoked)
│   │   │   │   └── fingerprint.go            # Device fingerprinting
│   │   │   │
│   │   │   ├── sessions/
│   │   │   │   ├── session.go                # Sessions (SessionID, IP, UserAgent, CreatedAt, ExpiresAt, Revoked)
│   │   │   │   └── session_validator.go      # Session validation logic
│   │   │   │
│   │   │   ├── recovery/                     # 🔄 MOVED from account_recovery
│   │   │   │   ├── recovery.go               # AccountRecovery (UserID, RecoveryMethod, Attempts[], Status)
│   │   │   │   ├── methods.go                # Methods: Email, SMS, SecurityQuestions, RecoveryCodes
│   │   │   │   ├── recovery_keys.go          # Recovery key management
│   │   │   │   ├── process.go                # Step-by-step recovery process
│   │   │   │   ├── rate_limiter.go           # Attempt limits
│   │   │   │   └── security_questions.go     # Security questions
│   │   │   │
│   │   │   ├── errors.go                     # Security-specific errors
│   │   │   ├── repository.go                 # Security repository interface (single repo for all security features)
│   │   │   └── events.go                     # TwoFAEnabled, TwoFADisabled, DeviceRegistered, DeviceRevoked, SessionRevoked, RecoveryInitiated, RecoveryCompleted
│   │   │
│   │   # ========================= PROFILE ENHANCEMENT DOMAINS (CONSOLIDATED) =========================
│   │   ├── profile_depth/                    # 🔄 CONSOLIDATED: rate history + taxonomy mapping ONLY
│   │   │   ├── entity.go                     # ProfileDepth (UserID, RateHistory[], TaxonomyMapping)
│   │   │   │                                 # NO badges (moved to badging)
│   │   │   │                                 # NO availability (moved to availability)
│   │   │   ├── rate_history.go               # Rate entries (amount, currency, effectiveAt)
│   │   │   ├── taxonomy.go                   # Normalized skills taxonomy mapping
│   │   │   ├── errors.go                     # Depth-specific errors
│   │   │   ├── repository.go                 # ProfileDepth repository interface
│   │   │   └── events.go                     # RateHistoryUpdated, TaxonomyUpdated
│   │   │
│   │   ├── profile_completeness/             # Completeness score ONLY
│   │   │   ├── entity.go                     # ProfileCompleteness (UserID, Score, MissingSections[], Recommendations[])
│   │   │   │                                 # NO view counts (moved to profile_analytics)
│   │   │   │                                 # NO badges (moved to badging)
│   │   │   ├── section_weights.go            # Different section weights (Skills=20%, Experience=15%, etc.)
│   │   │   ├── completeness_calculator.go    # Calculate profile completeness percentage
│   │   │   ├── recommendations.go            # Recommendations to improve profile
│   │   │   ├── errors.go                     # Completeness-specific errors
│   │   │   ├── repository.go                 # ProfileCompleteness repository interface
│   │   │   └── events.go                     # CompletenessUpdated, SectionCompleted, MilestoneReached
│   │   │
│   │   ├── profile_analytics/                # Views/search ONLY
│   │   │   ├── entity.go                     # ProfileAnalytics (UserID, Views[], ViewSources[], SearchAppearances)
│   │   │   │                                 # NO completeness (moved to profile_completeness)
│   │   │   │                                 # NO badges (moved to badging)
│   │   │   ├── view_tracking.go              # Track profile views (who, when, from where)
│   │   │   ├── search_analytics.go           # Track search appearances and rankings
│   │   │   ├── engagement_metrics.go         # Click-through rates, interest signals
│   │   │   ├── errors.go                     # Analytics-specific errors
│   │   │   ├── repository.go                 # ProfileAnalytics repository interface
│   │   │   └── events.go                     # ProfileViewed, SearchAppearance, EngagementRecorded
│   │   │
│   │   ├── profile_optimization/             # AI-powered profile optimization
│   │   │   ├── entity.go                     # ProfileOptimization (UserID, Suggestions[], Score, AppliedSuggestions[])
│   │   │   ├── optimization_engine.go        # AI suggestions for profile improvement
│   │   │   ├── keyword_optimization.go       # SEO keywords for better discoverability
│   │   │   ├── headline_suggestions.go       # Suggested headlines based on skills
│   │   │   ├── errors.go                     # Optimization-specific errors
│   │   │   ├── repository.go                 # ProfileOptimization repository interface
│   │   │   └── events.go                     # SuggestionGenerated, SuggestionApplied, OptimizationCompleted
│   │   │
│   │   # ========================= PROFILE VISIBILITY DOMAIN (CONSOLIDATED) =========================
│   │   ├── profile_visibility/               # 🔄 Search/discoverability ONLY
│   │   │   ├── entity.go                     # ProfileVisibility (UserID, VisibilityLevel, SearchableCategories[], HiddenFrom[])
│   │   │   │                                 # NO privacy flags (moved to privacy domain)
│   │   │   ├── visibility_level.go           # Levels: Public, LimitedPublic, Private, AnonymousMode
│   │   │   ├── search_preferences.go         # Control how profile appears in search
│   │   │   ├── stealth_mode.go               # Stealth/anonymous browsing
│   │   │   ├── errors.go                     # Visibility-specific errors
│   │   │   ├── repository.go                 # ProfileVisibility repository interface
│   │   │   └── events.go                     # VisibilityChanged, StealthModeEnabled
│   │   │
│   │   # ========================= AVAILABILITY DOMAIN (CONSOLIDATED) =========================
│   │   ├── availability/                     # 🔄 SINGLE SOURCE: status + calendar + rules + vacation
│   │   │   ├── entity.go                     # Availability aggregate (UserID, Status, Calendar, RecurringRules[], VacationMode)
│   │   │   │                                 # Profile/profile_depth reference this (NO duplicate fields)
│   │   │   ├── availability_status.go        # Status: Available, Busy, Away, DoNotDisturb
│   │   │   ├── recurring_schedule.go         # Weekly/monthly recurring availability slots
│   │   │   ├── vacation_mode.go              # Vacation/out-of-office mode
│   │   │   ├── calendar_sync.go              # Sync with Google Calendar, Outlook
│   │   │   ├── errors.go                     # Availability-specific errors
│   │   │   ├── repository.go                 # Availability repository interface
│   │   │   └── events.go                     # AvailabilityUpdated, VacationModeEnabled, CalendarSynced
│   │   │
│   │   ├── workload_capacity/                # Workload and capacity tracking
│   │   │   ├── entity.go                     # WorkloadCapacity (UserID, CurrentLoad, MaxCapacity, AvailableHours)
│   │   │   │                                 # References AvailabilityID
│   │   │   ├── capacity_calculator.go        # Calculate available capacity
│   │   │   ├── commitment_tracker.go         # Track current commitments
│   │   │   ├── overload_prevention.go        # Prevent overcommitment
│   │   │   ├── errors.go                     # Capacity-specific errors
│   │   │   ├── repository.go                 # WorkloadCapacity repository interface
│   │   │   └── events.go                     # CapacityUpdated, CapacityFull, CapacityAvailable
│   │   │
│   │   # ========================= NETWORKING & CONNECTIONS DOMAINS =========================
│   │   ├── professional_network/             # Professional networking (LinkedIn-style)
│   │   │   ├── entity.go                     # Network (UserID, Connections[], ConnectionRequests[])
│   │   │   ├── connection.go                 # Connection (UserID1, UserID2, ConnectedAt, Relationship)
│   │   │   ├── connection_request.go         # Pending connection requests
│   │   │   ├── relationship_type.go          # Types: Colleague, Client, Peer, Mentor
│   │   │   ├── network_analytics.go          # Network size, growth, strength
│   │   │   ├── errors.go                     # Network-specific errors
│   │   │   ├── repository.go                 # Network repository interface
│   │   │   └── events.go                     # ConnectionRequested, ConnectionAccepted, ConnectionRemoved
│   │   │
│   │   ├── referrals/                        # User referral system
│   │   │   ├── entity.go                     # Referral (ReferrerID, ReferredID, Status, Reward)
│   │   │   ├── referral_code.go              # Generate and manage referral codes
│   │   │   ├── referral_rewards.go           # Reward calculation and distribution
│   │   │   ├── referral_tracking.go          # Track referral conversions
│   │   │   ├── errors.go                     # Referral-specific errors
│   │   │   ├── repository.go                 # Referral repository interface
│   │   │   └── events.go                     # ReferralSent, ReferralConverted, RewardEarned
│   │   │
│   │   ├── user_groups/                      # Community groups (by skill, location, industry)
│   │   │   ├── entity.go                     # UserGroup (GroupID, Name, Members[], Moderators[])
│   │   │   ├── group_membership.go           # Manage group memberships
│   │   │   ├── group_categories.go           # Categories: Skills, Location, Industry, Interest
│   │   │   ├── group_activity.go             # Track group activity and engagement
│   │   │   ├── errors.go                     # Group-specific errors
│   │   │   ├── repository.go                 # UserGroup repository interface
│   │   │   └── events.go                     # GroupJoined, GroupLeft, GroupCreated
│   │   │
│   │   # ========================= FINANCIAL PROFILE DOMAINS =========================
│   │   ├── payment_methods/                  # Payment methods management
│   │   │   ├── entity.go                     # PaymentMethod (UserID, Type, Details, IsDefault, Verified)
│   │   │   ├── method_types.go               # Types: BankAccount, Card, PayPal, Crypto, Wise
│   │   │   ├── verification.go               # Verify payment methods
│   │   │   ├── withdrawal_preferences.go     # Preferred withdrawal methods
│   │   │   ├── errors.go                     # Payment method errors
│   │   │   ├── repository.go                 # PaymentMethod repository interface
│   │   │   └── events.go                     # PaymentMethodAdded, PaymentMethodVerified, DefaultMethodChanged
│   │   │
│   │   ├── financial_profile/                # Financial profile and preferences
│   │   │   ├── entity.go                     # FinancialProfile (UserID, PreferredCurrency, PaymentTerms, InvoiceSettings)
│   │   │   ├── currency_preferences.go       # Preferred currencies for work
│   │   │   ├── invoice_settings.go           # Invoice customization (logo, footer, terms)
│   │   │   ├── payment_terms.go              # Default payment terms (NET 30, upfront, milestones)
│   │   │   ├── errors.go                     # Financial profile errors
│   │   │   ├── repository.go                 # FinancialProfile repository interface
│   │   │   └── events.go                     # CurrencyPreferenceUpdated, InvoiceSettingsUpdated
│   │   │
│   │   ├── earning_goals/                    # Earning goals and tracking
│   │   │   ├── entity.go                     # EarningGoal (UserID, TargetAmount, Period, Progress)
│   │   │   ├── goal_types.go                 # Types: Monthly, Quarterly, Annual, Project-based
│   │   │   ├── progress_tracking.go          # Track progress toward goals
│   │   │   ├── achievement_notifications.go  # Notify when goals are met
│   │   │   ├── errors.go                     # Goal-specific errors
│   │   │   ├── repository.go                 # EarningGoal repository interface
│   │   │   └── events.go                     # GoalSet, GoalAchieved, ProgressUpdated
│   │   │
│   │   # ========================= PROFESSIONAL DEVELOPMENT DOMAINS =========================
│   │   ├── learning_path/                    # Personalized learning paths
│   │   │   ├── entity.go                     # LearningPath (UserID, Skills[], Courses[], Progress)
│   │   │   ├── skill_gaps.go                 # Identify skill gaps
│   │   │   ├── course_recommendations.go     # Recommend courses to fill gaps
│   │   │   ├── progress_tracking.go          # Track learning progress
│   │   │   ├── errors.go                     # Learning path errors
│   │   │   ├── repository.go                 # LearningPath repository interface
│   │   │   └── events.go                     # PathCreated, CourseCompleted, SkillAcquired
│   │   │
│   │   ├── mentorship/                       # Mentorship program
│   │   │   ├── entity.go                     # Mentorship (MentorID, MenteeID, Status, Goals[])
│   │   │   ├── mentor_profile.go             # Mentor availability and expertise
│   │   │   ├── mentorship_sessions.go        # Schedule and track sessions
│   │   │   ├── mentor_matching.go            # Match mentors with mentees
│   │   │   ├── errors.go                     # Mentorship errors
│   │   │   ├── repository.go                 # Mentorship repository interface
│   │   │   └── events.go                     # MentorshipRequested, SessionScheduled, MentorshipCompleted
│   │   │
│   │   ├── achievements/                     # Gamified achievements system (emits events for badging)
│   │   │   ├── entity.go                     # Achievement (UserID, AchievementType, UnlockedAt, Progress)
│   │   │   │                                 # Emits events → badging domain issues badges
│   │   │   ├── achievement_types.go          # Types: FirstJob, 10Jobs, TopRated, QuickResponder
│   │   │   ├── achievement_tiers.go          # Tiers: Bronze, Silver, Gold, Platinum
│   │   │   ├── progress_tracking.go          # Track progress toward achievements
│   │   │   ├── errors.go                     # Achievement errors
│   │   │   ├── repository.go                 # Achievement repository interface
│   │   │   └── events.go                     # AchievementUnlocked, TierReached, ProgressUpdated
│   │   │
│   │   # ========================= COMPLIANCE DOMAIN =========================
│   │   ├── compliance/
│   │   │   ├── tax_profile.go                # TaxProfile (UserID, Country, VAT/GST, TIN, W-Form refs)
│   │   │   ├── residency.go                  # Residency data (Country, Since, Proof docs)
│   │   │   ├── artifacts.go                  # Stored tax artifacts (W-8/W-9/VAT docs)
│   │   │   ├── errors.go                     # Compliance errors
│   │   │   ├── repository.go                 # Compliance repository interface
│   │   │   └── events.go                     # TaxProfileUpdated, ResidencyUpdated, ComplianceArtifactAdded
│   │   │
│   │   # ========================= COMMUNICATION PREFERENCES DOMAINS =========================
│   │   ├── communication_channels/           # Multi-channel communication preferences
│   │   │   ├── entity.go                     # CommunicationChannel (UserID, Channels[], Preferences)
│   │   │   ├── channel_types.go              # Types: Email, SMS, Push, InApp, WhatsApp, Slack
│   │   │   ├── notification_routing.go       # Route notifications to preferred channels
│   │   │   ├── quiet_hours.go                # Define quiet hours (no notifications)
│   │   │   ├── errors.go                     # Communication channel errors
│   │   │   ├── repository.go                 # CommunicationChannel repository interface
│   │   │   └── events.go                     # ChannelAdded, PreferencesUpdated, QuietHoursSet
│   │   │
│   │   └── email_preferences/                # Detailed email preferences
│   │       ├── entity.go                     # EmailPreferences (UserID, Frequencies[], Categories[], Digest)
│   │       ├── frequency_settings.go         # Frequency: RealTime, Daily, Weekly, Never
│   │       ├── category_settings.go          # Categories: JobAlerts, Messages, Updates, Marketing
│   │       ├── digest_settings.go            # Daily/weekly digest preferences
│   │       ├── errors.go                     # Email preference errors
│   │       ├── repository.go                 # EmailPreferences repository interface
│   │       └── events.go                     # PreferencesUpdated, DigestEnabled, CategoryMuted
│   │
│   ├── application/                          # 📋 Application Layer - Use cases & orchestration (Load Fourth)
│   │   │
│   │   # ========================= EVENT HANDLERS (Inbound Events) =========================
│   │   ├── eventhandler/
│   │   │   ├── contract_handler.go           # Consumes: contract.created → update user stats
│   │   │   ├── escrow_handler.go             # Consumes: escrow.released → update earnings
│   │   │   ├── payment_handler.go            # Consumes: payment.processed → update earnings/spend
│   │   │   ├── review_handler.go             # Consumes: review.submitted → update ratings/metrics
│   │   │   ├── message_handler.go            # Consumes: message.notification_delivered → update metrics
│   │   │   ├── admin_handler.go              # Consumes: admin.user_suspended → enforce suspension
│   │   │   ├── financial_risk_handler.go     # Consumes: financial.risk.alert.emitted → update risk score
│   │   │   ├── contract_status_handler.go    # Consumes: contract.financial_hold.placed/released → update status
│   │   │   │
│   │   │   # ========================= BADGING EVENT HANDLERS (Consumes from other contexts) =========================
│   │   │   ├── achievement_badging_handler.go # 🆕 Consumes: AchievementUnlocked → Issue achievement badge
│   │   │   ├── credential_badging_handler.go  # 🆕 Consumes: ExternalCertificationVerified/PlatformCertificationEarned → Issue cert badge
│   │   │   ├── trust_badging_handler.go       # 🆕 Consumes: TrustBadgeEarned → Issue trust badge
│   │   │   └── _common/
│   │   │       └── idempotency.go          🆕  # store last-seen event version per aggregate
│   │   ├── authz/                          🆕  # service-level permission checks (defense-in-depth)
│   │   │   ├── policies.go                  🆕 # map roles → actions (admin, owner, same-user)
│   │   │   └── guards.go                    🆕 # helpers used by services
│   │   │ 
│   │   │
│   │   # ========================= CORE USER SERVICES =========================
│   │   ├── user/
│   │   │   ├── service.go                    # User business logic (Create, Update, Delete, Verify, Search)
│   │   │   ├── commands.go                   # CreateUser, UpdateUser, DeleteUser, VerifyEmail
│   │   │   ├── queries.go                    # GetUser, ListUsers, SearchUsers, GetUserStats
│   │   │   ├── dto.go                        # UserDTO, CreateUserDTO, UpdateUserDTO, UserListDTO, UserStatsDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Input validation (email, username, password, phone)
│   │   │
│   │   ├── profile/
│   │   │   ├── service.go                    # Profile business logic (Complete, Update, CalculateCompleteness)
│   │   │   ├── commands.go                   # UpdateProfile, UpdatePreferences
│   │   │   ├── queries.go                    # GetProfile, GetProfileCompletion
│   │   │   │                                 # References AvailabilityService for availability data
│   │   │   ├── dto.go                        # ProfileDTO, UpdateProfileDTO
│   │   │   ├── mapper.go                     # Profile mappers
│   │   │   └── validators.go                 # Profile validators (bio, location)
│   │   │
│   │   # ========================= CAPABILITIES SERVICES (CONSOLIDATED) =========================
│   │   ├── capabilities/                     # 🔄 CONSOLIDATED service
│   │   │   ├── service.go                    # Capabilities business logic (Add/Update/Remove Skills, Add/Verify Specializations)
│   │   │   ├── commands.go                   # AddSkill, UpdateSkill, RemoveSkill, AddSpecialization, VerifySpecialization
│   │   │   ├── queries.go                    # ListSkills, GetSkillTaxonomy, ListSpecializations, GetNicheExpertise
│   │   │   ├── dto.go                        # CapabilitiesDTO, SkillDTO, SpecializationDTO, TaxonomyDTO
│   │   │   ├── mapper.go                     # Capabilities mappers
│   │   │   └── validators.go                 # Skill/specialization validation
│   │   │
│   │   ├── service_catalog/
│   │   │   ├── service.go                    # Service catalog business logic (Create, Update, Reference Capabilities)
│   │   │   ├── commands.go                   # CreateService, UpdateService, CreateServicePackage
│   │   │   ├── queries.go                    # GetServiceCatalog, GetServices, GetPackages
│   │   │   │                                 # Joins with CapabilitiesService to fetch capability details
│   │   │   ├── dto.go                        # ServiceCatalogDTO, ServiceDTO, ServicePackageDTO
│   │   │   ├── mapper.go                     # Service catalog mappers
│   │   │   └── validators.go                 # Service validation (pricing, capability references)
│   │   │
│   │   # ========================= EXPERIENCE & EDUCATION SERVICES =========================
│   │   ├── experience/
│   │   │   ├── service.go                    # Experience business logic (Add, Update, Delete, Verify)
│   │   │   ├── commands.go                   # AddExperience, UpdateExperience, DeleteExperience
│   │   │   ├── queries.go                    # GetExperience, ListExperience
│   │   │   ├── dto.go                        # ExperienceDTO, CreateExperienceDTO
│   │   │   ├── mapper.go                     # Experience mappers
│   │   │   └── validators.go                 # Date range validation, title/company presence
│   │   │
│   │   ├── education/
│   │   │   ├── service.go                    # Education business logic (Add, Update, Delete)
│   │   │   ├── commands.go                   # AddEducation, UpdateEducation, DeleteEducation
│   │   │   ├── queries.go                    # GetEducation, ListEducation
│   │   │   ├── dto.go                        # EducationDTO, CreateEducationDTO
│   │   │   ├── mapper.go                     # Education mappers
│   │   │   └── validators.go                 # School/degree validation, graduation year range
│   │   │
│   │   ├── language/
│   │   │   ├── service.go                    # Language business logic (Add, Update, Delete)
│   │   │   ├── commands.go                   # AddLanguage, UpdateLanguage, RemoveLanguage
│   │   │   ├── queries.go                    # ListLanguages
│   │   │   ├── dto.go                        # LanguageDTO, AddLanguageDTO
│   │   │   ├── mapper.go                     # Language mappers
│   │   │   └── validators.go                 # Language code, proficiency enum validation
│   │   │
│   │   # ========================= CREDENTIALS SERVICES (CONSOLIDATED) =========================
│   │   ├── credentials/                      # 🔄 CONSOLIDATED service
│   │   │   ├── service.go                    # Credentials business logic (Add/Verify External, Earn/Renew Platform)
│   │   │   ├── commands.go                   # AddExternalCertification, VerifyExternalCertification, EarnPlatformCertification, RenewPlatformCertification
│   │   │   ├── queries.go                    # GetCredentials, ListExternalCertifications, ListPlatformCertifications
│   │   │   ├── dto.go                        # CredentialsDTO, ExternalCertificationDTO, PlatformCertificationDTO
│   │   │   ├── mapper.go                     # Credentials mappers
│   │   │   └── validators.go                 # Certification validation (issuer, dates, credentials)
│   │   │
│   │   # ========================= PORTFOLIO & SHOWCASE SERVICES =========================
│   │   ├── portfolio/
│   │   │   ├── service.go                    # Portfolio business logic (Add, Update, Delete, Reorder)
│   │   │   ├── commands.go                   # AddPortfolioItem, UpdatePortfolioItem, DeletePortfolioItem, ReorderPortfolio
│   │   │   ├── queries.go                    # GetPortfolioItem, ListPortfolio
│   │   │   ├── dto.go                        # PortfolioItemDTO, CreatePortfolioItemDTO
│   │   │   ├── mapper.go                     # Portfolio mappers
│   │   │   └── validators.go                 # URL validation, media types, display order
│   │   │
│   │   # ========================= USER TYPE SPECIFIC SERVICES =========================
│   │   ├── freelancer/
│   │   │   ├── service.go                    # Freelancer business logic (CompleteProfile, UpdateRates, UpdateStats)
│   │   │   ├── commands.go                   # UpdateRates, UpdateAvailability
│   │   │   ├── queries.go                    # GetStats, GetEarnings, GetSuccessRate
│   │   │   ├── dto.go                        # FreelancerDTO, UpdateRatesDTO, FreelancerStatsDTO
│   │   │   ├── mapper.go                     # Freelancer mappers
│   │   │   └── validators.go                 # Rate currency, min/max bounds validation
│   │   │
│   │   ├── client/
│   │   │   ├── service.go                    # Client business logic (CompleteProfile, UpdateStats)
│   │   │   │                                 # References OrgService for company data (NO duplicate company logic)
│   │   │   ├── commands.go                   # UpdateClient, LinkToOrg
│   │   │   ├── queries.go                    # GetHiringStats, GetSpendingHistory
│   │   │   ├── dto.go                        # ClientDTO, ClientStatsDTO
│   │   │   ├── mapper.go                     # Client mappers
│   │   │   └── validators.go                 # Client validation
│   │   │
│   │   # ========================= IDENTITY VERIFICATION SERVICES (CONSOLIDATED) =========================
│   │   ├── identity_verification/            # 🔄 RENAMED service - KYC/KYB ONLY
│   │   │   ├── service.go                    # Identity verification business logic (Submit, Approve, Reject)
│   │   │   ├── commands.go                   # SubmitKYC, SubmitKYB, ApproveVerification, RejectVerification
│   │   │   ├── queries.go                    # GetVerification, ListVerifications, GetVerificationDocs
│   │   │   ├── dto.go                        # IdentityVerificationDTO, SubmitVerificationDTO
│   │   │   ├── mapper.go                     # Identity verification mappers
│   │   │   └── validators.go                 # Document validation, state transitions
│   │   │
│   │   # ========================= TRUST SERVICES (CONSOLIDATED) =========================
│   │   ├── trust/                            # 🔄 RENAMED service
│   │   │   ├── service.go                    # Trust business logic (CalculateTrustLevel, AwardTrustBadge)
│   │   │   ├── commands.go                   # UpdateTrustLevel, AwardTrustBadge, RevokeTrustBadge
│   │   │   ├── queries.go                    # GetTrustLevel, ListTrustBadges
│   │   │   ├── dto.go                        # TrustDTO, TrustLevelDTO, TrustBadgeDTO
│   │   │   ├── mapper.go                     # Trust mappers
│   │   │   └── validators.go                 # Trust level validation
│   │   │
│   │   # ========================= BADGING SERVICES (CONSOLIDATED) =========================
│   │   ├── badging/                          # 🔄 CONSOLIDATED service - ALL badge issuance
│   │   │   ├── service.go                    # Badging business logic (AwardBadge, RevokeBadge, CheckEligibility)
│   │   │   │                                 # Central service that issues ALL badges
│   │   │   ├── commands.go                   # AwardBadge, RevokeBadge
│   │   │   ├── queries.go                    # GetBadges, GetBadgesByType, CheckBadgeEligibility
│   │   │   ├── dto.go                        # BadgeDTO, BadgeEligibilityDTO
│   │   │   ├── mapper.go                     # Badging mappers
│   │   │   └── validators.go                 # Badge validation (eligibility criteria)
│   │   │
│   │   # ========================= SETTINGS & PREFERENCES SERVICES =========================
│   │   ├── settings/
│   │   │   ├── service.go                    # Settings business logic (Update, GetPreferences)
│   │   │   ├── commands.go                   # UpdateSettings, UpdateNotificationPrefs
│   │   │   ├── queries.go                    # GetSettings, GetNotificationPrefs
│   │   │   ├── dto.go                        # SettingsDTO, UpdateSettingsDTO, NotificationPrefsDTO
│   │   │   ├── mapper.go                     # Settings mappers
│   │   │   └── validators.go                 # Notification toggles, timezone/currency validation
│   │   │
│   │   ├── privacy/                          # 🔄 SPLIT service - privacy ONLY
│   │   │   ├── service.go                    # Privacy business logic (UpdatePrivacy)
│   │   │   ├── commands.go                   # UpdatePrivacySettings
│   │   │   ├── queries.go                    # GetPrivacySettings
│   │   │   ├── dto.go                        # PrivacySettingsDTO
│   │   │   ├── mapper.go                     # Privacy mappers
│   │   │   └── validators.go                 # Privacy flags validation
│   │   │
│   │   ├── saved_items/
│   │   │   ├── service.go                    # Saved items business logic (Save, Unsave, List, Search)
│   │   │   ├── commands.go                   # SaveItem, UnsaveItem
│   │   │   ├── queries.go                    # ListSavedItems, SearchSaved
│   │   │   ├── dto.go                        # SavedItemDTO, SaveItemDTO
│   │   │   ├── mapper.go                     # Saved items mappers
│   │   │   └── validators.go                 # Item type/id validation, dedupe constraints
│   │   │
│   │   # ========================= MODERATION SERVICES (CONSOLIDATED) =========================
│   │   ├── moderation/                       # 🔄 CONSOLIDATED service
│   │   │   ├── service.go                    # Moderation business logic (Suspend, Ban, Warn, Unsuspend, Unban)
│   │   │   │                                 # Single service for all moderation actions
│   │   │   ├── commands.go                   # SuspendUser, UnsuspendUser, BanUser, UnbanUser, IssueWarning, AcknowledgeWarning
│   │   │   ├── queries.go                    # GetModerationHistory, GetActiveSuspension, GetActiveBan, ListWarnings
│   │   │   ├── dto.go                        # ModerationDTO, SuspensionDTO, BanDTO, WarningDTO
│   │   │   ├── mapper.go                     # Moderation mappers
│   │   │   └── validators.go                 # Duration validation, reason enums, actor permissions
│   │   │
│   │   # ========================= SCORING SERVICES (CONSOLIDATED) =========================
│   │   ├── user_metrics/                     # 🔄 Service for raw metrics (single source)
│   │   │   ├── service.go                    # UserMetrics business logic (RecordMetric, UpdateMetrics, AggregateMetrics)
│   │   │   ├── commands.go                   # RecordMetric, UpdateMetric
│   │   │   ├── queries.go                    # GetMetrics, GetMetricHistory
│   │   │   ├── dto.go                        # UserMetricsDTO, MetricDTO
│   │   │   ├── mapper.go                     # Metrics mappers
│   │   │   └── validators.go                 # Metrics validation
│   │   │
│   │   ├── reputation/                       # Market-facing reputation service
│   │   │   ├── service.go                    # Reputation business logic (CalculateReputation, UpdateReputation)
│   │   │   │                                 # Uses pkg/scoring to compute from user_metrics
│   │   │   ├── commands.go                   # RecalculateReputation, RecordReputationEvent
│   │   │   ├── queries.go                    # GetReputationScore, GetReputationComponents, GetReputationHistory
│   │   │   ├── dto.go                        # ReputationDTO, ReputationComponentDTO
│   │   │   ├── mapper.go                     # Reputation mappers
│   │   │   └── validators.go                 # Reputation validation
│   │   │
│   │   ├── quality/                          # Work quality service
│   │   │   ├── service.go                    # Quality business logic (CalculateQualityScore, AnalyzeTrends)
│   │   │   │                                 # Uses pkg/scoring to compute from user_metrics
│   │   │   ├── commands.go                   # RecalculateQualityScore, RecordQualityMetric
│   │   │   ├── queries.go                    # GetQualityScore, GetScoringFactors, GetScoreTrend
│   │   │   ├── dto.go                        # QualityScoreDTO, ScoringFactorDTO, TrendAnalysisDTO
│   │   │   ├── mapper.go                     # Quality mappers
│   │   │   └── validators.go                 # Quality validation
│   │   │
│   │   ├── account_health/                   # Operational health service
│   │   │   ├── service.go                    # AccountHealth business logic (CalculateHealth, DetectIssues, GenerateRecommendations)
│   │   │   │                                 # Uses pkg/scoring to compute from user_metrics
│   │   │   ├── commands.go                   # RecalculateHealthScore, RecordHealthIssue
│   │   │   ├── queries.go                    # GetAccountHealth, GetHealthIssues, GetHealthRecommendations
│   │   │   ├── dto.go                        # AccountHealthDTO, HealthIssueDTO, HealthRecommendationDTO
│   │   │   ├── mapper.go                     # Account health mappers
│   │   │   └── validators.go                 # Health validation
│   │   │
│   │   ├── risk/                             # Fraud/safety service
│   │   │   ├── service.go                    # Risk business logic (ComputeRiskScore, RecordSignal, SetAccountHold)
│   │   │   │                                 # Uses pkg/scoring to compute from user_metrics + signals
│   │   │   ├── commands.go                   # RecordIpGeoMismatch, RecordDisputeCount, RecordChargeback, ApplyHold, ReleaseHold
│   │   │   ├── queries.go                    # GetRiskScore, ListSignals, GetAccountState
│   │   │   ├── dto.go                        # RiskDTO, RiskSignalDTO, RiskScoreDTO, AccountHoldDTO
│   │   │   ├── mapper.go                     # Risk mappers
│   │   │   └── validators.go                 # Risk validation
│   │   │
│   │   # ========================= ORGANIZATION & TEAM SERVICES =========================
│   │   ├── org/                              # 🔄 Service for company/org data (single source)
│   │   │   ├── service.go                    # Org business logic (CreateOrg, UpdateOrg, InviteMember, ManageSeats)
│   │   │   │                                 # Client service references this for company data
│   │   │   ├── commands.go                   # CreateOrg, UpdateOrg, InviteMember, RemoveMember, AssignRole, SetSeatCount
│   │   │   ├── queries.go                    # GetOrg, ListOrgsForUser, ListMembers, GetRoles, GetSeatUsage
│   │   │   ├── dto.go                        # OrgDTO, OrgMemberDTO, SeatUsageDTO
│   │   │   ├── mapper.go                     # Org mappers
│   │   │   └── validators.go                 # Org validation (seat limits, member capacity)
│   │   │
│   │   # ========================= SECURITY SERVICES (CONSOLIDATED) =========================
│   │   ├── security/                         # 🔄 CONSOLIDATED service
│   │   │   ├── service.go                    # Security business logic (Enable2FA, RegisterDevice, RevokeSession, InitiateRecovery)
│   │   │   │                                 # Single service for all security features
│   │   │   ├── commands.go                   # Enable2FA, Disable2FA, RegisterDevice, RevokeDevice, RevokeSession, StartRecoveryProcess, CompleteRecovery
│   │   │   ├── queries.go                    # Get2FAStatus, ListDevices, ListSessions, GetRecoveryMethods, GetSecuritySettings
│   │   │   ├── dto.go                        # SecurityDTO, TwoFADTO, DeviceDTO, SessionDTO, RecoveryDTO
│   │   │   ├── mapper.go                     # Security mappers
│   │   │   └── validators.go                 # Security validation (2FA methods, device fingerprints, recovery attempts)
│   │   │
│   │   # ========================= PROFILE ENHANCEMENT SERVICES (CONSOLIDATED) =========================
│   │   ├── profile_depth/                    # 🔄 Service for rate history + taxonomy ONLY
│   │   │   ├── service.go                    # ProfileDepth business logic (AppendRateHistory, NormalizeSkills)
│   │   │   │                                 # NO badge logic (moved to badging)
│   │   │   │                                 # NO availability logic (moved to availability)
│   │   │   ├── commands.go                   # AddHourlyRateEntry, NormalizeSkillSet
│   │   │   ├── queries.go                    # GetRateHistory, GetNormalizedSkills
│   │   │   ├── dto.go                        # RateHistoryDTO, NormalizedSkillDTO
│   │   │   ├── mapper.go                     # Profile depth mappers
│   │   │   └── validators.go                 # Rate validation
│   │   │
│   │   ├── profile_completeness/             # Completeness service ONLY
│   │   │   ├── service.go                    # ProfileCompleteness business logic (CalculateCompleteness, IdentifyMissingSections)
│   │   │   │                                 # NO view counts (moved to profile_analytics)
│   │   │   ├── commands.go                   # RecalculateCompleteness, MarkSectionComplete
│   │   │   ├── queries.go                    # GetCompletenessScore, GetMissingSections, GetRecommendations
│   │   │   ├── dto.go                        # ProfileCompletenessDTO, SectionWeightDTO, RecommendationDTO
│   │   │   ├── mapper.go                     # Completeness mappers
│   │   │   └── validators.go                 # Completeness validation
│   │   │
│   │   ├── profile_analytics/                # Views/search service ONLY
│   │   │   ├── service.go                    # ProfileAnalytics business logic (TrackProfileView, RecordSearchAppearance)
│   │   │   │                                 # NO completeness (moved to profile_completeness)
│   │   │   ├── commands.go                   # RecordView, RecordSearchImpression
│   │   │   ├── queries.go                    # GetProfileViews, GetSearchAnalytics, GetEngagementMetrics
│   │   │   ├── dto.go                        # ProfileAnalyticsDTO, ViewTrackingDTO, SearchAnalyticsDTO
│   │   │   ├── mapper.go                     # Analytics mappers
│   │   │   └── validators.go                 # Analytics validation
│   │   │
│   │   ├── profile_optimization/             # AI optimization service
│   │   │   ├── service.go                    # ProfileOptimization business logic (GenerateOptimizations, ApplySuggestion)
│   │   │   ├── commands.go                   # GenerateAISuggestions, ApplyOptimization
│   │   │   ├── queries.go                    # GetOptimizations, GetKeywordSuggestions
│   │   │   ├── dto.go                        # ProfileOptimizationDTO, OptimizationSuggestionDTO
│   │   │   ├── mapper.go                     # Optimization mappers
│   │   │   └── validators.go                 # Optimization validation
│   │   │
│   │   # ========================= PROFILE VISIBILITY SERVICES (CONSOLIDATED) =========================
│   │   ├── profile_visibility/               # 🔄 Search/discoverability service ONLY
│   │   │   ├── service.go                    # ProfileVisibility business logic (UpdateVisibility, EnableStealthMode)
│   │   │   │                                 # References PrivacyService for privacy flags (NO duplicate flags)
│   │   │   ├── commands.go                   # ChangeVisibilityLevel, UpdateSearchableCategories, ToggleStealthMode
│   │   │   ├── queries.go                    # GetVisibilitySettings, GetSearchPreferences, GetStealthStatus
│   │   │   ├── dto.go                        # ProfileVisibilityDTO, VisibilityLevelDTO, SearchPreferencesDTO
│   │   │   ├── mapper.go                     # Visibility mappers
│   │   │   └── validators.go                 # Visibility validation
│   │   │
│   │   # ========================= AVAILABILITY SERVICES (CONSOLIDATED) =========================
│   │   ├── availability/                     # 🔄 SINGLE SOURCE service
│   │   │   ├── service.go                    # Availability business logic (UpdateAvailability, SetRecurringSchedule, EnableVacationMode)
│   │   │   │                                 # Profile/profile_depth reference this service (NO duplicate availability logic)
│   │   │   ├── commands.go                   # SetAvailabilityStatus, CreateRecurringSchedule, ToggleVacationMode, SyncExternalCalendar
│   │   │   ├── queries.go                    # GetAvailability, GetRecurringSchedule, GetVacationStatus
│   │   │   ├── dto.go                        # AvailabilityDTO, RecurringScheduleDTO, VacationModeDTO
│   │   │   ├── mapper.go                     # Availability mappers
│   │   │   └── validators.go                 # Slot conflicts, timezone validation
│   │   │
│   │   ├── workload_capacity/
│   │   │   ├── service.go                    # WorkloadCapacity business logic (CalculateCapacity, TrackCommitments)
│   │   │   │                                 # References AvailabilityService
│   │   │   ├── commands.go                   # UpdateCurrentLoad, AddCommitment, RemoveCommitment, SetMaxCapacity
│   │   │   ├── queries.go                    # GetCapacity, GetCurrentLoad, GetAvailableHours
│   │   │   ├── dto.go                        # WorkloadCapacityDTO, CommitmentDTO
│   │   │   ├── mapper.go                     # Workload mappers
│   │   │   └── validators.go                 # Capacity limits, commitment conflicts
│   │   │
│   │   # ========================= NETWORKING & CONNECTIONS SERVICES =========================
│   │   ├── professional_network/
│   │   │   ├── service.go                    # Network business logic (SendConnectionRequest, AcceptConnection, AnalyzeNetwork)
│   │   │   ├── commands.go                   # CreateConnectionRequest, AcceptRequest, DeclineRequest, RemoveConnection
│   │   │   ├── queries.go                    # GetConnections, GetConnectionRequests, GetNetworkAnalytics
│   │   │   ├── dto.go                        # NetworkDTO, ConnectionDTO, ConnectionRequestDTO
│   │   │   ├── mapper.go                     # Network mappers
│   │   │   └── validators.go                 # Prevent duplicate requests, relationship validation
│   │   │
│   │   ├── referrals/
│   │   │   ├── service.go                    # Referral business logic (GenerateReferralCode, TrackReferral, CalculateReward)
│   │   │   ├── commands.go                   # CreateReferralCode, RecordReferralClick, MarkReferralConverted, IssueReward
│   │   │   ├── queries.go                    # GetReferralCode, GetReferralStats, GetReferralRewards
│   │   │   ├── dto.go                        # ReferralDTO, ReferralCodeDTO, ReferralRewardDTO
│   │   │   ├── mapper.go                     # Referral mappers
│   │   │   └── validators.go                 # Code uniqueness, reward eligibility
│   │   │
│   │   ├── user_groups/
│   │   │   ├── service.go                    # UserGroups business logic (CreateGroup, JoinGroup, ModerateGroup)
│   │   │   ├── commands.go                   # CreateUserGroup, AddMember, RemoveMember, AssignModerator
│   │   │   ├── queries.go                    # GetGroup, ListUserGroups, GetGroupMembers
│   │   │   ├── dto.go                        # UserGroupDTO, GroupMembershipDTO
│   │   │   ├── mapper.go                     # User group mappers
│   │   │   └── validators.go                 # Member limits, category validation
│   │   │
│   │   # ========================= FINANCIAL PROFILE SERVICES =========================
│   │   ├── payment_methods/
│   │   │   ├── service.go                    # PaymentMethods business logic (AddMethod, VerifyMethod, SetDefault)
│   │   │   ├── commands.go                   # CreatePaymentMethod, VerifyMethod, SetAsDefault, DeletePaymentMethod
│   │   │   ├── queries.go                    # GetPaymentMethods, GetDefaultMethod
│   │   │   ├── dto.go                        # PaymentMethodDTO, MethodVerificationDTO
│   │   │   ├── mapper.go                     # Payment method mappers
│   │   │   └── validators.go                 # Method type validation, account verification
│   │   │
│   │   ├── financial_profile/
│   │   │   ├── service.go                    # FinancialProfile business logic (UpdateProfile, SetCurrencyPreferences)
│   │   │   ├── commands.go                   # UpdateCurrencyPreferences, UpdateInvoiceSettings, SetDefaultPaymentTerms
│   │   │   ├── queries.go                    # GetFinancialProfile, GetCurrencyPreferences, GetInvoiceSettings
│   │   │   ├── dto.go                        # FinancialProfileDTO, CurrencyPreferencesDTO, InvoiceSettingsDTO
│   │   │   ├── mapper.go                     # Financial profile mappers
│   │   │   └── validators.go                 # Currency codes, invoice customization limits
│   │   │
│   │   ├── earning_goals/
│   │   │   ├── service.go                    # EarningGoals business logic (SetGoal, TrackProgress, NotifyAchievement)
│   │   │   ├── commands.go                   # CreateEarningGoal, UpdateGoalProgress, MarkGoalAchieved
│   │   │   ├── queries.go                    # GetEarningGoals, GetGoalProgress, GetAchievements
│   │   │   ├── dto.go                        # EarningGoalDTO, GoalProgressDTO
│   │   │   ├── mapper.go                     # Earning goal mappers
│   │   │   └── validators.go                 # Amount validation, period validation
│   │   │
│   │   # ========================= PROFESSIONAL DEVELOPMENT SERVICES =========================
│   │   ├── learning_path/
│   │   │   ├── service.go                    # LearningPath business logic (CreatePath, IdentifySkillGaps, RecommendCourses)
│   │   │   ├── commands.go                   # GenerateLearningPath, EnrollInCourse, CompleteCourse
│   │   │   ├── queries.go                    # GetLearningPath, GetSkillGaps, GetCourseRecommendations
│   │   │   ├── dto.go                        # LearningPathDTO, SkillGapDTO, CourseRecommendationDTO
│   │   │   ├── mapper.go                     # Learning path mappers
│   │   │   └── validators.go                 # Skill prerequisite validation
│   │   │
│   │   ├── mentorship/
│   │   │   ├── service.go                    # Mentorship business logic (RequestMentorship, MatchMentor, ScheduleSession)
│   │   │   ├── commands.go                   # CreateMentorshipRequest, AcceptMentorship, ScheduleMentorshipSession
│   │   │   ├── queries.go                    # GetMentorship, GetMentorProfile, ListAvailableMentors
│   │   │   ├── dto.go                        # MentorshipDTO, MentorProfileDTO, MentorshipSessionDTO
│   │   │   ├── mapper.go                     # Mentorship mappers
│   │   │   └── validators.go                 # Mentor eligibility, session scheduling
│   │   │
│   │   ├── achievements/                     # Achievement tracking service (emits events for badging)
│   │   │   ├── service.go                    # Achievements business logic (TrackProgress, UnlockAchievement)
│   │   │   │                                 # Emits AchievementUnlocked event → BadgingService issues badge
│   │   │   ├── commands.go                   # RecordAchievementProgress, UnlockAchievement
│   │   │   ├── queries.go                    # GetAchievements, GetAchievementProgress
│   │   │   ├── dto.go                        # AchievementDTO, AchievementProgressDTO
│   │   │   ├── mapper.go                     # Achievement mappers
│   │   │   └── validators.go                 # Unlock criteria validation
│   │   │
│   │   # ========================= COMPLIANCE SERVICES =========================
│   │   ├── compliance/
│   │   │   ├── service.go                    # Compliance business logic (UpsertTaxProfile, ValidateCountryFields)
│   │   │   ├── commands.go                   # CreateOrUpdateTaxProfile, SetResidency, AttachWForm, AttachVAT
│   │   │   ├── queries.go                    # GetTaxProfile, GetResidency, ListComplianceArtifacts
│   │   │   ├── dto.go                        # TaxProfileDTO, ResidencyDTO
│   │   │   ├── mapper.go                     # Compliance mappers
│   │   │   └── validators.go                 # Compliance validators
│   │   │
│   │   # ========================= COMMUNICATION PREFERENCES SERVICES =========================
│   │   ├── communication_channels/
│   │   │   ├── service.go                    # Communication business logic (ManageChannels, RouteNotifications, SetQuietHours)
│   │   │   ├── commands.go                   # AddChannel, RemoveChannel, SetChannelPreferences, ConfigureQuietHours
│   │   │   ├── queries.go                    # GetChannels, GetChannelPreferences, GetQuietHours
│   │   │   ├── dto.go                        # CommunicationChannelDTO, ChannelPreferencesDTO, QuietHoursDTO
│   │   │   ├── mapper.go                     # Communication channel mappers
│   │   │   └── validators.go                 # Channel type validation, quiet hours validation
│   │   │
│   │   └── email_preferences/
│   │       ├── service.go                    # EmailPreferences business logic (UpdatePreferences, SetFrequency, ConfigureDigest)
│   │       ├── commands.go                   # UpdateEmailFrequency, UpdateCategoryPreferences, EnableDigest
│   │       ├── queries.go                    # GetEmailPreferences, GetFrequencySettings, GetDigestSettings
│   │       ├── dto.go                        # EmailPreferencesDTO, FrequencySettingsDTO, DigestSettingsDTO
│   │       ├── mapper.go                     # Email preferences mappers
│   │       └── validators.go                 # Frequency validation, category validation
│   │
│   └── interfaces/                           # 🌐 Interface Layer (API/HTTP - Load Fifth)
│       └── http/
│           │
│           ├── middleware/
│           │   ├── etag.go                           # 🆕 Thin wrapper → uses infra/http/etag
│           │   ├── requestid.go              # 🆕 Wraps platform-shared request id
│           │   ├── logging.go                # 🆕 Wraps platform-shared logging
│           │   ├── recovery.go               # 🆕 Wraps platform-shared recovery
│           │   ├── tracing.go                # 🆕 Wraps platform-shared otel middleware
│           │   ├── auth.go                   # 🆕 Wraps pkg/auth (Keycloak)
│           │   ├── rbac.go                   # 🆕 Role/permission gate
│           │   └── cors.go                   # 🆕 CORS middleware
│           │
│           ├── presenters/
│           │   └── errors.go                 # 🎯 Maps domain/service errors → HTTP status & problem+json
│           │
│           ├── handlers/
│           │   └── health_handler.go                 # 🆕 Unversioned: /healthz/live, /healthz/ready
│           ├── routes/
│           │   └── health_routes.go                  # 🆕 Mount unversioned /healthz/*
│           │
│           ├── v1/                           # 🆕 Versioned API surface
│           │   │
│           │   ├── openapi/                  # 🆕 OpenAPI Contract
│           │   │   ├── openapi.yaml          # 📜 OpenAPI contract v1: request/response schemas
│           │   │   └── generator.go          # 🧭 Serves /v1/swagger and /v1/openapi.json (dev only)
│           │   │
│           │   ├── handlers/
│           │   │   │
│           │   │   # ========================= CORE USER HANDLERS =========================
│           │   │   ├── user_handler.go           # GET, POST, PUT, DELETE /users, /users/:id, /users/search
│           │   │   │
│           │   │   # ========================= PROFILE HANDLERS =========================
│           │   │   ├── profile_handler.go        # GET, PUT /users/:id/profile
│           │   │   │
│           │   │   # ========================= CAPABILITIES HANDLERS (CONSOLIDATED) =========================
│           │   │   ├── capabilities_handler.go   # 🔄 CONSOLIDATED: skills + specializations
│           │   │   │                             # GET, POST, DELETE /users/:id/capabilities/skills
│           │   │   │                             # GET, POST /users/:id/capabilities/specializations
│           │   │   ├── service_catalog_handler.go # GET, POST, PUT /users/:id/service-catalog
│           │   │   │
│           │   │   # ========================= EXPERIENCE & EDUCATION HANDLERS =========================
│           │   │   ├── experience_handler.go     # GET, POST, PUT, DELETE /users/:id/experience
│           │   │   ├── education_handler.go      # GET, POST, PUT, DELETE /users/:id/education
│           │   │   ├── language_handler.go       # GET, POST, DELETE /users/:id/languages
│           │   │   │
│           │   │   # ========================= CREDENTIALS HANDLERS (CONSOLIDATED) =========================
│           │   │   ├── credentials_handler.go    # 🔄 CONSOLIDATED: external + platform certifications
│           │   │   │                             # GET, POST /users/:id/credentials/external
│           │   │   │                             # GET, POST /users/:id/credentials/platform
│           │   │   │
│           │   │   # ========================= PORTFOLIO HANDLERS =========================
│           │   │   ├── portfolio_handler.go      # GET, POST, PUT, DELETE /users/:id/portfolio
│           │   │   │
│           │   │   # ========================= USER TYPE SPECIFIC HANDLERS =========================
│           │   │   ├── freelancer_handler.go     # GET, PUT /users/:id/freelancer, /users/:id/freelancer/stats
│           │   │   ├── client_handler.go         # GET, PUT /users/:id/client, /users/:id/client/stats
│           │   │   │
│           │   │   # ========================= IDENTITY VERIFICATION HANDLERS (CONSOLIDATED) =========================
│           │   │   ├── identity_verification_handler.go # 🔄 RENAMED: POST, GET /users/:id/identity-verification
│           │   │   │
│           │   │   # ========================= TRUST HANDLERS (CONSOLIDATED) =========================
│           │   │   ├── trust_handler.go          # 🔄 RENAMED: GET /users/:id/trust-level, GET /users/:id/trust-badges
│           │   │   │
│           │   │   # ========================= BADGING HANDLERS (CONSOLIDATED) =========================
│           │   │   ├── badging_handler.go        # 🔄 CONSOLIDATED: ALL badges
│           │   │   │                             # GET /users/:id/badges, GET /users/:id/badges/:type
│           │   │   │
│           │   │   # ========================= SETTINGS & PREFERENCES HANDLERS =========================
│           │   │   ├── settings_handler.go       # GET, PUT /users/:id/settings
│           │   │   ├── privacy_handler.go        # 🔄 SPLIT: GET, PUT /users/:id/privacy
│           │   │   ├── saved_items_handler.go    # GET, POST, DELETE /users/:id/saved
│           │   │   │
│           │   │   # ========================= MODERATION HANDLERS (CONSOLIDATED) =========================
│           │   │   ├── moderation_handler.go     # 🔄 CONSOLIDATED: suspension + ban + warning
│           │   │   │                             # POST /admin/users/:id/moderation/suspend
│           │   │   │                             # POST /admin/users/:id/moderation/ban
│           │   │   │                             # POST /admin/users/:id/moderation/warn
│           │   │   │                             # GET /admin/users/:id/moderation/history
│           │   │   │
│           │   │   # ========================= SCORING HANDLERS (CONSOLIDATED) =========================
│           │   │   ├── user_metrics_handler.go   # 🔄 GET /users/:id/metrics (raw metrics - internal only)
│           │   │   ├── reputation_handler.go     # GET /users/:id/reputation
│           │   │   ├── quality_handler.go        # GET /users/:id/quality-score
│           │   │   ├── account_health_handler.go # GET /users/:id/account-health
│           │   │   ├── risk_handler.go           # GET /admin/users/:id/risk (admin only)
│           │   │   │
│           │   │   # ========================= ORGANIZATION HANDLERS =========================
│           │   │   ├── org_handler.go            # GET, POST, PUT /orgs, /orgs/:id, /orgs/:id/members
│           │   │   │
│           │   │   # ========================= SECURITY HANDLERS (CONSOLIDATED) =========================
│           │   │   ├── security_handler.go       # 🔄 CONSOLIDATED: 2FA + devices + sessions + recovery
│           │   │   │                             # POST /users/:id/security/2fa/enable
│           │   │   │                             # GET, POST, DELETE /users/:id/security/devices
│           │   │   │                             # GET, DELETE /users/:id/security/sessions
│           │   │   │                             # POST /users/:id/security/recovery/initiate
│           │   │   │
│           │   │   # ========================= PROFILE ENHANCEMENT HANDLERS (CONSOLIDATED) =========================
│           │   │   ├── profile_depth_handler.go  # 🔄 GET /users/:id/profile-depth/rate-history
│           │   │   │                             # GET /users/:id/profile-depth/taxonomy
│           │   │   ├── profile_completeness_handler.go # GET /users/:id/profile/completeness
│           │   │   ├── profile_analytics_handler.go # GET /users/:id/profile/analytics
│           │   │   ├── profile_optimization_handler.go # POST /users/:id/profile/optimize, GET /users/:id/profile/suggestions
│           │   │   │
│           │   │   # ========================= PROFILE VISIBILITY HANDLERS (CONSOLIDATED) =========================
│           │   │   ├── profile_visibility_handler.go # 🔄 PUT /users/:id/profile/visibility
│           │   │   │
│           │   │   # ========================= AVAILABILITY HANDLERS (CONSOLIDATED) =========================
│           │   │   ├── availability_handler.go   # 🔄 GET, PUT /users/:id/availability
│           │   │   │                             # POST /users/:id/availability/vacation-mode
│           │   │   ├── workload_capacity_handler.go # GET, PUT /users/:id/capacity
│           │   │   │
│           │   │   # ========================= NETWORKING HANDLERS =========================
│           │   │   ├── professional_network_handler.go # POST /users/:id/connections/request, GET /users/:id/connections
│           │   │   ├── referrals_handler.go      # GET /users/:id/referral-code, POST /referrals/track
│           │   │   ├── user_groups_handler.go    # POST /groups, POST /groups/:id/join, GET /users/:id/groups
│           │   │   │
│           │   │   # ========================= FINANCIAL PROFILE HANDLERS =========================
│           │   │   ├── payment_methods_handler.go # POST, PUT, DELETE /users/:id/payment-methods
│           │   │   ├── financial_profile_handler.go # GET, PUT /users/:id/financial-profile
│           │   │   ├── earning_goals_handler.go  # POST, GET /users/:id/earning-goals
│           │   │   │
│           │   │   # ========================= PROFESSIONAL DEVELOPMENT HANDLERS =========================
│           │   │   ├── learning_path_handler.go  # POST, GET /users/:id/learning-path
│           │   │   ├── mentorship_handler.go     # POST, GET /users/:id/mentorship
│           │   │   ├── achievements_handler.go   # GET /users/:id/achievements
│           │   │   │
│           │   │   # ========================= COMPLIANCE HANDLERS =========================
│           │   │   ├── compliance_handler.go     # GET, PUT /users/:id/compliance/tax-profile
│           │   │   │
│           │   │   # ========================= COMMUNICATION PREFERENCES HANDLERS =========================
│           │   │   ├── communication_channels_handler.go # GET, PUT /users/:id/communication/channels
│           │   │   └── email_preferences_handler.go # GET, PUT /users/:id/email-preferences
│           │   │
│           │   └── routes/
│           │       │
│           │       # ========================= CORE USER ROUTES =========================
│           │       ├── user_routes.go            # /users, /users/:id, /users/search
│           │       │
│           │       # ========================= PROFILE ROUTES =========================
│           │       ├── profile_routes.go         # /users/:id/profile
│           │       │
│           │       # ========================= CAPABILITIES ROUTES (CONSOLIDATED) =========================
│           │       ├── capabilities_routes.go    # 🔄 /users/:id/capabilities/*
│           │       ├── service_catalog_routes.go # /users/:id/service-catalog
│           │       │
│           │       # ========================= EXPERIENCE & EDUCATION ROUTES =========================
│           │       ├── experience_routes.go      # /users/:id/experience
│           │       ├── education_routes.go       # /users/:id/education
│           │       ├── language_routes.go        # /users/:id/languages
│           │       │
│           │       # ========================= CREDENTIALS ROUTES (CONSOLIDATED) =========================
│           │       ├── credentials_routes.go     # 🔄 /users/:id/credentials/*
│           │       │
│           │       # ========================= PORTFOLIO ROUTES =========================
│           │       ├── portfolio_routes.go       # /users/:id/portfolio
│           │       │
│           │       # ========================= USER TYPE SPECIFIC ROUTES =========================
│           │       ├── freelancer_routes.go      # /users/:id/freelancer
│           │       ├── client_routes.go          # /users/:id/client
│           │       │
│           │       # ========================= IDENTITY VERIFICATION ROUTES (CONSOLIDATED) =========================
│           │       ├── identity_verification_routes.go # 🔄 /users/:id/identity-verification
│           │       │
│           │       # ========================= TRUST ROUTES (CONSOLIDATED) =========================
│           │       ├── trust_routes.go           # 🔄 /users/:id/trust-level, /users/:id/trust-badges
│           │       │
│           │       # ========================= BADGING ROUTES (CONSOLIDATED) =========================
│           │       ├── badging_routes.go         # 🔄 /users/:id/badges
│           │       │
│           │       # ========================= SETTINGS & PREFERENCES ROUTES =========================
│           │       ├── settings_routes.go        # /users/:id/settings
│           │       ├── privacy_routes.go         # 🔄 /users/:id/privacy
│           │       ├── saved_items_routes.go     # /users/:id/saved
│           │       │
│           │       # ========================= MODERATION ROUTES (CONSOLIDATED) =========================
│           │       ├── moderation_routes.go      # 🔄 /admin/users/:id/moderation/*
│           │       │
│           │       # ========================= SCORING ROUTES (CONSOLIDATED) =========================
│           │       ├── user_metrics_routes.go    # 🔄 /users/:id/metrics (internal only)
│           │       ├── reputation_routes.go      # /users/:id/reputation
│           │       ├── quality_routes.go         # /users/:id/quality-score
│           │       ├── account_health_routes.go  # /users/:id/account-health
│           │       ├── risk_routes.go            # /admin/users/:id/risk
│           │       │
│           │       # ========================= ORGANIZATION ROUTES =========================
│           │       ├── org_routes.go             # /orgs, /orgs/:id, /orgs/:id/members
│           │       │
│           │       # ========================= SECURITY ROUTES (CONSOLIDATED) =========================
│           │       ├── security_routes.go        # 🔄 /users/:id/security/*
│           │       │
│           │       # ========================= PROFILE ENHANCEMENT ROUTES (CONSOLIDATED) =========================
│           │       ├── profile_depth_routes.go   # 🔄 /users/:id/profile-depth/*
│           │       ├── profile_completeness_routes.go # /users/:id/profile/completeness
│           │       ├── profile_analytics_routes.go # /users/:id/profile/analytics
│           │       ├── profile_optimization_routes.go # /users/:id/profile/optimize
│           │       │
│           │       # ========================= PROFILE VISIBILITY ROUTES (CONSOLIDATED) =========================
│           │       ├── profile_visibility_routes.go # 🔄 /users/:id/profile/visibility
│           │       │
│           │       # ========================= AVAILABILITY ROUTES (CONSOLIDATED) =========================
│           │       ├── availability_routes.go    # 🔄 /users/:id/availability
│           │       ├── workload_capacity_routes.go # /users/:id/capacity
│           │       │
│           │       # ========================= NETWORKING ROUTES =========================
│           │       ├── professional_network_routes.go # /users/:id/connections
│           │       ├── referrals_routes.go       # /users/:id/referral-code, /referrals
│           │       ├── user_groups_routes.go     # /groups, /users/:id/groups
│           │       │
│           │       # ========================= FINANCIAL PROFILE ROUTES =========================
│           │       ├── payment_methods_routes.go # /users/:id/payment-methods
│           │       ├── financial_profile_routes.go # /users/:id/financial-profile
│           │       ├── earning_goals_routes.go   # /users/:id/earning-goals
│           │       │
│           │       # ========================= PROFESSIONAL DEVELOPMENT ROUTES =========================
│           │       ├── learning_path_routes.go   # /users/:id/learning-path
│           │       ├── mentorship_routes.go      # /users/:id/mentorship
│           │       ├── achievements_routes.go    # /users/:id/achievements
│           │       │
│           │       # ========================= COMPLIANCE ROUTES =========================
│           │       ├── compliance_routes.go      # /users/:id/compliance
│           │       │
│           │       # ========================= COMMUNICATION PREFERENCES ROUTES =========================
│           │       ├── communication_channels_routes.go # /users/:id/communication
│           │       └── email_preferences_routes.go # /users/:id/email-preferences
│           │
│           └── router.go                     # 📝 HTTP router setup (Gin with platform-shared middleware; mounts /v1 group)
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                         # Service-specific errors
│   │   └── codes.go                          # Error codes (USER_NOT_FOUND, EMAIL_TAKEN, VETTING_NOT_COMPLETE, etc.)
│   │
│   ├── logger/                               # ❌ REMOVED
│   │   └── README.md                         # Points to platform-shared/logging
│   │
│   ├── utils/
│   │   ├── validator.go                      # Local validation utilities
│   │   ├── slug.go                           # Generate URL-friendly slugs
│   │   ├── sanitizer.go                      # Sanitize user input
│   │   ├── etag.go                           # 🆕 Weak ETag helper for GET /users
│   │   ├── text_analyzer.go                  # Text analysis utilities
│   │   └── matching_algorithm.go             # Mentor matching, skill gap analysis
│   │
│   ├── scoring/                              # 🆕 Shared scoring kernel
│   │   ├── calculator.go                     # 🆕 Shared scoring engine (used by reputation/quality/health/risk)
│   │   ├── weights.go                        # 🆕 Configurable weights per scoring context
│   │   └── metrics_aggregator.go             # 🆕 Aggregate raw metrics from user_metrics
│   │
│   └── constants/
│       ├── events.go                         # ❌ REMOVED - Use contracts/events
│       ├── topics.go                         # ❌ REMOVED - Use contracts/events
│       ├── badge_types.go                    # 🔄 Badge type constants (Achievement, Certification, Trust, Platform)
│       ├── moderation_reasons.go             # 🔄 Moderation reason constants (shared across suspension/ban/warning)
│       ├── trust_levels.go                   # 🔄 Trust level constants
│       └── verification_types.go             # 🔄 Verification type constants (KYC, KYB)
│
├── config/                                   # 🆕 Configuration files
│   ├── default.yaml                          # Default configuration
│   ├── dev.yaml                              # Development overrides
│   └── prod.yaml                             # Production overrides
│
├── dapr/                                     # 🆕 Dapr components
│   ├── local/
│   │   ├── pubsub.yaml                       # Local Kafka pub/sub
│   │   └── statestore.yaml                   # Local state store
│   └── k8s/
│       ├── pubsub.yaml                       # Production Kafka with scopes: ["users-be"]
│       ├── statestore.yaml                   # Production state store with scopes: ["users-be"]
│       └── secrets.yaml                      # Secrets component
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                   # Kubernetes Deployment
│       ├── service.yaml                      # Kubernetes Service
│       ├── configmap.yaml                    # ConfigMap
│       ├── secrets.yaml                      # Secrets
│       ├── hpa.yaml                          # Horizontal Pod Autoscaler
│       ├── pdb.yaml                          # Pod Disruption Budget
│       └── servicemonitor.yaml               # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                        # Local dev setup
│   ├── get-secrets.sh                        # Fetch secrets
│   ├── seed-data.sh                          # Seed initial data
│   ├── migrate.sh                           🆕  # SQL migrations (dev/prod flags)
│   ├── openapi-diff.sh                      🆕  # fail CI on breaking API changes
│   ├── schema-diff.sh                       🆕  # DB schema guardrail (pg_dump + mig verify)
│   ├── generate-sdks.sh                     🆕  # produce /sdk clients from OpenAPI
│   └── sbom-sign.sh                         🆕  # build SBOM + cosign sign
│
├── tests/
│   ├── reliability/                         🆕
│   │   ├── projections_replay_test.go       🆕  # rebuild read models from event logs
│   │   └── outbox_dispatcher_test.go        🆕  # at-least-once + idempotency assertions
│   ├── property/                            🆕
│   │   └── scoring_property_test.go         🆕  # property/fuzz tests for scoring kernel
│   ├── unit/
│   │   ├── domain/
│   │   │   # ========================= CORE DOMAIN TESTS =========================
│   │   │   ├── user_test.go                  # User entity tests
│   │   │   ├── user_validators_test.go       # VO/validator tests
│   │   │   ├── profile_test.go               # Profile entity tests
│   │   │   #
│   │   │   # ========================= CONSOLIDATED DOMAIN TESTS =========================
│   │   │   ├── capabilities_test.go          # 🔄 Capabilities tests (skills + specializations)
│   │   │   ├── credentials_test.go           # 🔄 Credentials tests (external + platform)
│   │   │   ├── badging_test.go               # 🔄 Badging tests (all badge types)
│   │   │   ├── moderation_test.go            # 🔄 Moderation tests (suspension + ban + warning)
│   │   │   ├── security_test.go              # 🔄 Security tests (2FA + devices + sessions + recovery)
│   │   │   ├── availability_test.go          # 🔄 Availability tests (single source)
│   │   │   #
│   │   │   # ========================= SCORING DOMAIN TESTS =========================
│   │   │   ├── user_metrics_test.go          # 🔄 User metrics tests (raw metrics)
│   │   │   ├── reputation_test.go            # 🔄 Reputation tests
│   │   │   ├── quality_test.go               # 🔄 Quality tests
│   │   │   ├── account_health_test.go        # 🔄 Account health tests
│   │   │   ├── risk_test.go                  # 🔄 Risk tests
│   │   │   #
│   │   │   # ========================= PROFILE ENHANCEMENT TESTS =========================
│   │   │   ├── profile_depth_test.go         # 🔄 Profile depth tests (rate history + taxonomy)
│   │   │   ├── profile_completeness_test.go  # 🔄 Completeness tests
│   │   │   ├── profile_analytics_test.go     # 🔄 Analytics tests
│   │   │   ├── profile_visibility_test.go    # 🔄 Visibility tests
│   │   │   #
│   │   │   # ========================= OTHER DOMAIN TESTS =========================
│   │   │   ├── org_test.go                   # 🔄 Org tests (single source for company)
│   │   │   ├── trust_test.go                 # 🔄 Trust tests
│   │   │   └── identity_verification_test.go # 🔄 Identity verification tests
│   │   ├── application/
│   │   │   # ========================= CORE SERVICE TESTS =========================
│   │   │   ├── user_service_test.go          # User service tests
│   │   │   ├── profile_service_test.go       # Profile service tests
│   │   │   #
│   │   │   # ========================= CONSOLIDATED SERVICE TESTS =========================
│   │   │   ├── capabilities_service_test.go  # 🔄 Capabilities service tests
│   │   │   ├── credentials_service_test.go   # 🔄 Credentials service tests
│   │   │   ├── badging_service_test.go       # 🔄 Badging service tests
│   │   │   ├── moderation_service_test.go    # 🔄 Moderation service tests
│   │   │   ├── security_service_test.go      # 🔄 Security service tests
│   │   │   ├── availability_service_test.go  # 🔄 Availability service tests
│   │   │   #
│   │   │   # ========================= SCORING SERVICE TESTS =========================
│   │   │   ├── user_metrics_service_test.go  # 🔄 User metrics service tests
│   │   │   ├── reputation_service_test.go    # 🔄 Reputation service tests
│   │   │   ├── quality_service_test.go       # 🔄 Quality service tests
│   │   │   ├── account_health_service_test.go # 🔄 Account health service tests
│   │   │   └── risk_service_test.go          # 🔄 Risk service tests
│   │   └── infrastructure/
│   │       ├── user_repository_test.go       # User repository tests
│   │       ├── capabilities_repository_test.go # 🔄 Capabilities repository tests
│   │       ├── badging_repository_test.go    # 🔄 Badging repository tests
│   │       ├── moderation_repository_test.go # 🔄 Moderation repository tests
│   │       ├── security_repository_test.go   # 🔄 Security repository tests
│   │       ├── user_metrics_repository_test.go # 🔄 User metrics repository tests
│   │       ├── scheduler_tasks_test.go       # 🆕 Scheduler task tests
│   │       └── projections_test.go           # 🆕 Projection tests
│   ├── integration/
│   │   ├── handlers/
│   │   │   ├── user_handler_test.go          # User handler integration tests
│   │   │   ├── capabilities_handler_test.go  # 🔄 Capabilities handler tests
│   │   │   ├── badging_handler_test.go       # 🔄 Badging handler tests
│   │   │   ├── moderation_handler_test.go    # 🔄 Moderation handler tests
│   │   │   ├── security_handler_test.go      # 🔄 Security handler tests
│   │   │   └── reputation_handler_test.go    # 🔄 Reputation handler tests
│   │   ├── repositories/
│   │   │   ├── user_repository_integration_test.go # User repository integration tests
│   │   │   ├── capabilities_repository_integration_test.go # 🔄 Capabilities repository tests
│   │   │   └── user_metrics_repository_integration_test.go # 🔄 User metrics repository tests
│   │   └── worker_tasks_integration_test.go  # 🆕 Worker tasks integration tests
│   └── e2e/
|       ├── README.md                                      # 🆕 How to run: `go test -tags=e2e ./tests/e2e/...`
│       ├── chaos_event_delivery_test.go      🆕  # simulate duplicates, reordering, delays
│       └── scenarios/
│           ├── user_registration_test.go     # E2E user registration flow
│           ├── profile_completion_test.go    # E2E profile completion flow
│           ├── vetting_flow_test.go          # 🔄 E2E vetting flow
│           ├── badging_flow_test.go          # 🔄 E2E badging flow
│           ├── moderation_flow_test.go       # 🔄 E2E moderation flow
│           └── scoring_flow_test.go          # 🔄 E2E scoring flow
|
├── docs/
│   ├── README.md                             # Service overview
│   ├── API.md                                # API documentation
│   ├── ARCHITECTURE.md                       # Service architecture
│   ├── MIGRATIONS.md                         # Migration history
│   ├── SCHEMA.md                             # Database schema
│   ├── RUNBOOK.md                            # Operations runbook
│   │
│   # ========================= PLATFORM & OPERATIONS DOCS =========================
│   ├── API_VERSIONING.md                     🆕  # HTTP contract versioning & deprecation policy
│   ├── CACHING.md                            🆕  # keys, TTLs, SWR, invalidation events
│   ├── DATA_RETENTION.md                     🛡️🆕 # per-domain retention windows
│   ├── ERASURE.md                            🛡️🆕 # GDPR/CCPA erasure hooks & playbooks
│   ├── SLOS.md                               📏🆕 # target P99s, projection lag, queue delay
│   ├── OUTBOX.md                             🆕  # exactly-once-ish design, retries, DLQ
│   ├── RELEASE_CHECKLIST.md                  🆕  # preflight (openapi diff, schema diff, migrations plan)
│   └── NAMING.md                             ♻️  # package snake_case; HTTP kebab-case conventions
│
│   # ========================= CORE FEATURE DOCUMENTATION =========================
│   ├── EVENTS.md                             # Events (published & consumed)
│   ├── EVENT_VERSIONING.md                   # 🆕 Event versioning guide
│   ├── user-types.md                         # User type documentation
│   ├── profile-management.md                 # Profile management
│   │
│   # ========================= CONSOLIDATED FEATURE DOCUMENTATION =========================
│   ├── capabilities.md                       # 🔄 CONSOLIDATED: skills + specializations + taxonomy
│   ├── credentials.md                        # 🔄 CONSOLIDATED: external + platform certifications
│   ├── badging.md                            # 🔄 CONSOLIDATED: ALL badge issuance
│   ├── moderation.md                         # 🔄 CONSOLIDATED: suspension + ban + warning
│   ├── security.md                           # 🔄 CONSOLIDATED: 2FA + devices + sessions + recovery
│   ├── availability.md                       # 🔄 CONSOLIDATED: single source for availability
│   │
│   # ========================= SCORING DOCUMENTATION =========================
│   ├── scoring-architecture.md               # 🆕 Shared scoring kernel architecture
│   ├── user-metrics.md                       # 🆕 Raw metrics (single source)
│   ├── reputation.md                         # Reputation scoring
│   ├── quality.md                            # Quality scoring
│   ├── account-health.md                     # Health scoring
│   ├── risk.md                               # Risk scoring
│   │
│   # ========================= PROFILE ENHANCEMENT DOCUMENTATION =========================
│   ├── profile-depth.md                      # 🔄 Rate history + taxonomy
│   ├── profile-completeness.md               # Completeness tracking
│   ├── profile-analytics.md                  # Views/search analytics
│   ├── profile-visibility.md                 # Visibility settings
│   │
│   # ========================= OTHER FEATURE DOCUMENTATION =========================
│   ├── identity-verification.md              # 🔄 KYC/KYB documentation
│   ├── trust.md                              # 🔄 Trust levels & badges
│   ├── organizations.md                      # 🔄 Company/org management (single source)
│   ├── professional-network.md               # Professional networking
│   ├── referral-program.md                   # Referral system
│   ├── user-groups.md                        # Community groups
│   ├── payment-methods.md                    # Payment methods
│   ├── financial-profile.md                  # Financial profile
│   ├── earning-goals.md                      # Earning goals
│   ├── learning-paths.md                     # Learning paths
│   ├── mentorship-program.md                 # Mentorship
│   ├── achievements.md                       # Achievements
│   ├── compliance.md                         # Tax/residency compliance
│   ├── communication-channels.md             # Communication preferences
│   └── migration-guide.md                    # 🆕 Migration guide from old to new structure
│
├── .github/
│   └── workflows/
│       ├── ci.yml                            # CI workflow
│       ├── cd.yml                            # CD workflow
│       ├── contract-ci.yml                  🆕  # openapi-diff + event schema checks
│       ├── security.yml                     🆕  # golangci-lint, govulncheck, trivy, cosign verify
│       └── load-tests.yml                   🆕  # k6/gatling smoke against /healthz & hot paths
│
│
├── go.mod                                    # 📝 Dependencies (imports pkg/auth, platform-shared, contracts/events)
├── go.sum
├── .env.example                              # Example environment variables
├── Makefile                                  # Build targets
├── Dockerfile                                # Multi-stage Docker build
├── .dockerignore                             # Docker ignore patterns
├── .gitignore                                # Git ignore patterns
└── README.md                                 # Service README



```

---

## 📦 **2️⃣ jobs-be (Job Posting and Management Service)**
```

apps/be/jobs-be/
│
├── cmd/
│   ├── api/
│   │   └── main.go                           # 📝 API entrypoint - initializes Gin, Dapr, Postgres (uses platform-shared/logging, internal/config)
│   │                                         # ♻️ Serve versioned routes (/v1/*) & enable ETag middleware when flag is on
│   │
│   └── worker/
│       └── main.go                           # 🧵 Worker entrypoint - bootstraps DI, registers cron tasks, runs background jobs
│                                             # ♻️ Wire leader election + 🆕 outbox dispatcher (see internal/infrastructure/{coordination,messaging/outbox})
│
├── internal/
│   │
│   ├── config/                               # 🔧 Configuration (Load First)
│   │   ├── feature_flags.go                  🆕  # central toggles (enable_job_etag, enable_ai_assist, enable_search_signals, etc.)
│   │   ├── schema.go                         ♻️  # group flags under Config.FeatureFlags
│   │   ├── loader.go                         # Config loader using Viper (CLI → ENV → file → defaults)
│   │   └── docs/
│   │       └── CONFIGURATION.md              # Configuration documentation
│   │
│   ├── ioc/                                  # 🧩 Dependency Injection Container
│   │   ├── container.go                      # DI graph: constructs DB/Redis/Kafka clients, repositories, services, handlers, schedulers
│   │   └── wiring.go                         ♻️  # wire outbox, coordination, observers based on env & FeatureFlags
│   │
│   ├── platform/                             # 🆕 Consolidated external/platform adapters
│   │   ├── clients/                          # MOVED FROM internal/external + internal/storage + internal/search
│   │   │   ├── users_client.go               # 📝 Users service client (validate client exists, get user details)
│   │   │   ├── proposals_client.go           # 📝 Proposals service client (check proposal count, acceptance status)
│   │   │   ├── contracts_client.go           # 📝 Contracts service client (transition job → contract seed)
│   │   │   │                                 # contracts-be owns: payment schedules, milestones, hourly terms, invoices
│   │   │   ├── financial_client.go           # 📝 Financial service client (check overdue invoices, get fee schedules)
│   │   │   │                                 # financial-be owns: wallets, escrow, tax forms, actual payment processing
│   │   │   ├── reviews_client.go             # 📝 Reviews service client (submit feedback for aggregation)
│   │   │   │                                 # reviews-be owns: ratings, reputation, feedback aggregation, badges
│   │   │   ├── storage_client.go             # 📝 Storage service client (upload job attachments/previews → storage-be)  # RENAMED FROM internal/storage/client.go
│   │   │   │                                 # Jobs-be only stores references, actual files in storage-be
│   │   │   ├── search_client.go              # 📝 Search service client (emit signals to search-be)                       # RENAMED FROM internal/search/client.go
│   │   │   │                                 # search-be owns: saved searches, detailed analytics, taxonomy synonyms
│   │   │   └── auth_client.go                # 📝 Auth service client (MFA verification via Keycloak/pkg/auth)
│   │   │                                     # Keycloak owns: MFA state, 2FA, authentication
│   │   └── http/
│   │       ├── idempotency_adapter.go        # 🆕 Bind platform-shared idempotency to Gin (MOVED from internal/http/idempotency_adapter.go)
│   │       └── etag.go                             🆕  # middleware-level ETag (pairs with pkg/utils/etag)
│   │
│   ├── infrastructure/                       # 🔧 Infrastructure Layer (Load Second)
│   │   │
│   │   ├── coordination/                     🆕  # distributed coordination primitives
│   │   │   ├── leader_election.go            🆕  # single-active cron/worker
│   │   │   └── distlock.go                   🆕  # Redis/PG advisory locks for cron tasks
│   │   │
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection setup
│   │   │       ├── transaction.go            # Transaction helpers
│   │   │       ├── outbox_store.go           🆕  # tx-coupled outbox table access
│   │   │       ├── migrations.go             # Auto-migration logic with version tracking
│   │   │       ├── version.go                # Schema version tracking
│   │   │       ├── safety.go                 # Pre-migration safety checks
│   │   │       ├── migrations/               🆕  # SQL-first, forward-only migrations (prod)
│   │   │       │   ├── jobs/                 🆕  # per-schema migration folders
│   │   │       │   ├── security/             🆕
│   │   │       │   ├── scoring/              🆕
│   │   │       │   └── growth/               🆕  # promo/tags/sharing (future split)
│   │   │       │
│   │   │       # ========================= CORE JOB REPOSITORIES =========================
│   │   │       ├── job_repository.go         # Job CRUD (Create, Update, FindByID, FindByClient, Search)
│   │   │       ├── category_repository.go    # Category repository (hierarchical tree)
│   │   │       ├── skill_repository.go       # Skill taxonomy repository (SINGLE SOURCE)
│   │   │       ├── job_skill_repository.go   # Job-skill relationships
│   │   │       │
│   │   │       # ========================= SCREENING REPOSITORY (CONSOLIDATED) =========================
│   │   │       ├── screening_repository.go   # 🔄 CONSOLIDATED: job_question + screening_compliance
│   │   │       │                             # Stores: screening questions, skill tests, NDA flags, export control, compliance policies
│   │   │       │
│   │   │       # ========================= ATTACHMENTS REPOSITORY (CONSOLIDATED) =========================
│   │   │       ├── attachments_repository.go # 🔄 CONSOLIDATED: job_attachment + job_previews
│   │   │       │                             # Stores: file references (type: Document, Image, VR, AR) - actual files in storage-be
│   │   │       │
│   │   │       # ========================= INVITATION & SOURCING REPOSITORIES =========================
│   │   │       ├── invitation_repository.go  # Job invitations repository
│   │   │       ├── sourcing_repository.go    # Sourcing modes (public, invite-only, private link, shortlists, talent pools)
│   │   │       │
│   │   │       # ========================= BUDGET & VISIBILITY REPOSITORIES =========================
│   │   │       ├── budget_controls_repository.go # Budget controls (min/max ranges, currency, hourly caps)
│   │   │       ├── visibility_lifecycle_repository.go # Visibility/lifecycle (schedules, auto-close, extensions, draft/published)
│   │   │       │
│   │   │       # ========================= TEMPLATE REPOSITORY (CONSOLIDATED) =========================
│   │   │       ├── template_repository.go    # 🔄 CONSOLIDATED: job_template + template_versions
│   │   │       │                             # Single repository for templates WITH versioning (semver, changelog, deprecation)
│   │   │       │
│   │   │       # ========================= ELIGIBILITY & REQUIREMENTS REPOSITORIES =========================
│   │   │       ├── eligibility_rules_repository.go # Hard apply gates (geo allow/deny, KYC, min tier, agency allowed)    # UPDATED: includes geo/TZ/radius merged from geo_requirements
│   │   │       ├── requirements_matrix_repository.go # Weighted must-have vs nice-to-have requirements
│   │   │       │
│   │   │       # ========================= HIRING TEAM REPOSITORY =========================
│   │   │       ├── hiring_team_repository.go      # Collaborators/roles per job (OWNER, REVIEWER, INTERVIEWER)
│   │   │       │                                  # (Pipeline moved to proposals-be)
│   │   │       │
│   │   │       # ========================= ANALYTICS READ MODEL =========================
│   │   │       │                                  # Source of truth for lightweight analytics lives under internal/projections
│   │   │       │                                  # Detailed behavior analytics → search-be (via events)
│   │   │       │
│   │   │       # ========================= MODERATION REPOSITORY (CONSOLIDATED) =========================
│   │   │       ├── moderation_repository.go  # 🔄 CONSOLIDATED: job_flag + moderation_lifecycle
│   │   │       │                             # Single repository for: flags (reports) + moderation state (CLEAN, LIMITED, QUARANTINED)
│   │   │       │
│   │   │       # ========================= A/B TESTING & EXPERIMENTS REPOSITORIES =========================
│   │   │       ├── ab_experiments_repository.go   # Listing A/B tests (titles, descriptions, question packs)
│   │   │       │
│   │   │       # ========================= SYNDICATION & DRAFTS REPOSITORIES =========================
│   │   │       ├── syndication_repository.go      # External boards/feeds integration
│   │   │       ├── drafts_repository.go           # Autosave & multi-draft support
│   │   │       │
│   │   │       # ========================= DUPLICATE DETECTION REPOSITORY =========================
│   │   │       ├── duplicate_detection_repository.go # Near-duplicate jobs detection (simhash, clustering)
│   │   │       │
│   │   │       # ========================= LEGAL & COMPLIANCE REPOSITORIES (CONSOLIDATED) =========================
│   │   │       ├── legal_controls_repository.go   # 🔄 CONSOLIDATED: legal_controls + data_residency
│   │   │       │                             # Stores: export control, NDA pinning, legal hold, data residency policy (EU/US)
│   │   │       │
│   │   │       # ========================= CAMPAIGN & RETENTION REPOSITORIES =========================
│   │   │       ├── campaign_tags_repository.go    # Internal labels for campaigns/VIPs
│   │   │       ├── retention_rules_repository.go  # Archival/anonymization schedules
│   │   │       │
│   │   │       # ========================= PROMOTION & JOB PREFERENCES REPOSITORIES =========================
│   │   │       ├── promotion_repository.go        # Paid boosts/featured jobs
│   │   │       ├── job_preference_repository.go   # 🔄 RENAMED FROM preference_repository.go — advisory matching preferences (non-blocking soft signals)
│   │   │       │                                  # NO payment terms (→ contracts-be), NO tax (→ financial-be)
│   │   │       │
│   │   │       # ========================= MULTI-HIRE & REPOST REPOSITORY =========================
│   │   │       ├── hiring_option_repository.go    # Multi-hire & repost config
│   │   │       │
│   │   │       # ========================= AI ASSIST REPOSITORY =========================
│   │   │       ├── ai_assist_repository.go        # AI suggestions (skills, categories, description optimization)
│   │   │       │
│   │   │       # ========================= INCLUSIVITY REPOSITORY =========================
│   │   │       ├── inclusivity_repository.go      # Inclusivity/accessibility flags (FlexibleHours, NoVideoRequired, ScreenReaderFriendly)
│   │   │       │
│   │   │       # ========================= CONTRACT TRANSITION REPOSITORY =========================
│   │   │       ├── contract_transition_repository.go # Job → contracts-be transition mapping
│   │   │       │
│   │   │       # ========================= FRAUD DETECTION REPOSITORY =========================
│   │   │       ├── fraud_detection_repository.go  # Fraud detection / risk auto-flags
│   │   │       │
│   │   │       # ========================= ESG REPOSITORY =========================
│   │   │       ├── esg_repository.go              # ESG / sustainability flags & carbon estimate
│   │   │       │
│   │   │       # ========================= SHARING REPOSITORY =========================
│   │   │       ├── sharing_repository.go          # Job sharing & referral incentives (tracked links with UTM)
│   │   │       │
│   │   │       # ========================= BULK OPERATIONS REPOSITORY =========================
│   │   │       ├── bulk_ops_repository.go         # Bulk operations (multi-job changes/imports)
│   │   │       │
│   │   │       # ========================= WEBHOOKS REPOSITORY =========================
│   │   │       ├── webhooks_repository.go         # Webhook subscriptions (job-scoped)
│   │   │       │
│   │   │       # ========================= ARCHIVE REPOSITORY =========================
│   │   │       ├── archive_repository.go          # User-managed archive/reactivate
│   │   │       │
│   │   │       # ========================= CUSTOM FIELDS REPOSITORY =========================
│   │   │       ├── custom_fields_repository.go    # Custom fields (schema + values)
│   │   │       │
│   │   │       # ========================= LOCALIZATION REPOSITORY =========================
│   │   │       └── localization_repository.go     # Multi-language job content
│   │   │
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go             # Redis connection setup
│   │   │       ├── singleflight.go           🆕  # stampede protection for hot keys
│   │   │       ├── invalidation_rules.go     🆕  # map events → keys to drop (doc alongside)
│   │   │       ├── keys.go                   # 🆕 Canonical cache keys & TTLs (no magic strings)
│   │   │       │
│   │   │       # ========================= CORE CACHING =========================
│   │   │       ├── job_cache.go              # Job caching
│   │   │       ├── category_cache.go         # Category tree caching
│   │   │       ├── skill_cache.go            # Popular skills caching
│   │   │       │
│   │   │       # ========================= CONSOLIDATED CACHING =========================
│   │   │       ├── template_cache.go         # 🔄 Templates WITH versions cache
│   │   │       ├── screening_cache.go        # 🔄 Screening (questions + compliance) cache
│   │   │       ├── attachments_cache.go      # 🔄 Attachments (documents + previews) cache
│   │   │       ├── moderation_cache.go       # 🔄 Moderation (flags + state) cache
│   │   │       ├── legal_controls_cache.go   # 🔄 Legal controls (export + NDA + residency) cache
│   │   │       │
│   │   │       # ========================= OTHER CACHING =========================
│   │   │       ├── sourcing_cache.go         # Sourcing mode & lists cache
│   │   │       ├── eligibility_rules_cache.go # Policy by job_id (short TTL) — includes geo/TZ/radius keys  # UPDATED (absorbed former geo cache)
│   │   │       ├── hiring_team_cache.go      # Roster cache (bust on changes)
│   │   │       ├── requirements_matrix_cache.go # Matrix + ETag
│   │   │       ├── job_analytics_cache.go    # Lightweight counters cache
│   │   │       ├── promotion_cache.go        # Promo status until expiry
│   │   │       └── job_preference_cache.go   # 🔄 RENAMED FROM preference_cache.go — Soft prefs for prefilter
│   │   │
│   │   ├── scheduler/
│   │   │   ├── cron.go                       ♻️ # wrap each task with distlock + jitter
│   │   │   └── tasks/
│   │   │       └── task_guard.go             🆕  # idempotency tokens + last-run watermark
│   │   │
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go               # 📝 Kafka consumer (uses platform-shared/inbox)
│   │   │       ├── producer.go               # 📝 Kafka producer (uses platform-shared/outbox)
│   │   │       ├── topics.go                 # 📝 Topic constants (imported from contracts/events)
│   │   │       │
│   │   │       # ========================= TOPIC DEFINITIONS =========================
│   │   │       ├── topics_job.go             # Job lifecycle topics
│   │   │       ├── topics_moderation.go      # 🔄 Moderation topics
│   │   │       ├── topics_analytics.go       # 🔄 Analytics signal topics (to search-be)
│   │   │       │
│   │   │       # ========================= PRODUCERS =========================
│   │   │       ├── producers_job.go          # Job event producers
│   │   │       ├── producers_moderation.go   # 🔄 Moderation event producers
│   │   │       ├── producers_analytics.go    # 🔄 Analytics signal producers (to search-be)
│   │   │       │
│   │   │       └── scram.go                  # SCRAM authentication for Kafka
│   │   │
│   │   ├── messaging_adapters/
│   │   │   ├── inbox/
│   │   │   │   └── dapr_subscriptions.yaml   # 🆕 Service-scoped Dapr subscriptions
│   │   │   └── outbox/
│   │   │       ├── dispatcher.go             🆕  # reads outbox → Kafka (with retries/DLQ)
│   │   │       ├── metrics.go                🆕  # publish lag, failures, retries
│   │   │       └── publisher.go              # 🆕 Thin wrapper over platform-shared outbox
│   │   │
│   │   ├── observability/                    🆕  # consolidated telemetry utilities
│   │   │   ├── metrics.go                    🆕  # RED/USE metrics helpers
│   │   │   ├── tracing.go                    🆕  # common attrs (hashed user_id, org_id, event_id)
│   │   │   └── slo_monitor.go                🆕  # projection lag, queue delay, API P99s
│   │   │
│   │   ├── security/                         # 🆕 Security Utilities
│   │   │   ├── kms.go                        # 🔐 KMS wrapper: envelope-encrypt sensitive fields (NDA content, export control data)
│   │   │   └── pii_redactor.go               # 🧽 PII redactor: masks emails/phones in logs
│   │   │
│   │   └── ai/
│   │       ├── ai_assist_client.go           # AI client for job optimization (skills/categories suggestions, description optimization)
│   │       └── duplicate_detector_client.go  # AI client for near-duplicate detection (simhash, clustering)
│   │
│   ├── projections/                          # 🆕 CQRS Projections
│   │   ├── job_analytics/
│   │   │   ├── projector.go                  # 🧮 Event projector: builds lightweight analytics read model
│   │   │   │                                 # Consumes: JobViewed, ProposalSubmitted, InvitationSent → increments counters
│   │   │   │                                 # Emits detailed events to search-be for full analytics
│   │   │   └── readmodel_repository.go       # 📖 Read-model repository for fast counter queries (SOURCE OF TRUTH)
│   │   │
│   │   └── job_search/
│   │       ├── projector.go                  # Consumes job events → emits search signals to search-be
│   │       │                                 # search-be owns full-text search, facets, saved searches
│   │       └── signal_emitter.go             # Signal emitter to search-be
│   │
│   ├── domain/                               # 🏛️ Domain Layer - Business logic & entities (Load Third)
│   │   │
│   │   # ========================= CORE JOB DOMAIN =========================
│   │   ├── job/
│   │   │   ├── entity.go                     # Job aggregate root (ID, ClientID, Title, Description, Budget, Skills, Status, Type)
│   │   │   ├── value_objects.go              # JobBudget, JobSkill, JobCategory, JobDuration value objects
│   │   │   ├── enums.go                      # JobType (Fixed, Hourly, Retainer), Status (Draft, Open, InProgress, Closed), ExperienceLevel
│   │   │   ├── lifecycle.go                  # Lifecycle states (scheduled, auto-close rules, extensions/renewals)
│   │   │   ├── errors.go                     # Domain errors (ErrJobNotFound, ErrInvalidBudget, ErrJobAlreadyClosed)
│   │   │   ├── repository.go                 # JobRepository interface
│   │   │   └── events.go                     # JobCreated, JobUpdated, JobPublished, JobClosed, JobArchived
│   │   │
│   │   ├── category/
│   │   │   ├── entity.go                     # Category entity (ID, Name, Slug, ParentID, Level) - hierarchical categories
│   │   │   ├── subcategory.go                # Nested categories support (tree structure)
│   │   │   ├── errors.go                     # Category errors
│   │   │   ├── repository.go                 # CategoryRepository interface
│   │   │   └── events.go                     # CategoryCreated, CategoryUpdated, CategoryReparented
│   │   │
│   │   ├── skill/
│   │   │   ├── entity.go                     # Skill taxonomy (ID, Name, CategoryID, Popularity) - SINGLE SOURCE
│   │   │   ├── category.go                   # Skill categories (grouping related skills)
│   │   │   ├── errors.go                     # Skill errors
│   │   │   ├── repository.go                 # SkillRepository interface
│   │   │   └── events.go                     # SkillCreated, SkillUpdated, SkillDeprecated
│   │   │
│   │   ├── job_skill/
│   │   │   ├── entity.go                     # Skills required for job (JobID, SkillID, IsRequired)
│   │   │   ├── requirement_level.go          # Required vs Preferred skill levels
│   │   │   ├── errors.go                     # Job skill errors
│   │   │   ├── repository.go                 # JobSkillRepository interface
│   │   │   └── events.go                     # JobSkillAdded, JobSkillUpdated, JobSkillRemoved
│   │   │
│   │   # ========================= SCREENING DOMAIN (CONSOLIDATED) =========================
│   │   ├── screening/                        # 🔄 CONSOLIDATED: job_question + screening_compliance
│   │   │   ├── entity.go                     # Screening aggregate (JobID, Questions[], SkillsTests[], NDARequired, ExportControlFlag, CompliancePolicies)
│   │   │   ├── questions/
│   │   │   │   ├── question.go               # Screening questions (Question, QuestionType, IsRequired)
│   │   │   │   └── question_type.go          # Text, MultiChoice, File upload question types
│   │   │   ├── compliance/
│   │   │   │   ├── policy.go                 # Compliance policies (NDA, export control, security questionnaires)
│   │   │   │   ├── nda.go                    # NDA requirements
│   │   │   │   └── export_control.go         # Export control flags
│   │   │   ├── errors.go                     # Screening errors
│   │   │   ├── repository.go                 # Screening repository interface
│   │   │   └── events.go                     # ScreeningConfigured, ScreeningQuestionAdded, ScreeningQuestionRemoved, NDARequired, ExportControlSet
│   │   │
│   │   # ========================= ATTACHMENTS DOMAIN (CONSOLIDATED) =========================
│   │   ├── attachments/                      # 🔄 CONSOLIDATED: job_attachment + job_previews
│   │   │   ├── entity.go                     # Attachment aggregate (JobID, FileURL, FileType, FileName, AttachmentType)
│   │   │   │                                 # Stores REFERENCES ONLY - actual files in storage-be
│   │   │   ├── attachment_type.go            # Type: Document, Image, Video, VR, AR
│   │   │   ├── vr_preview.go                 # VR preview metadata
│   │   │   ├── ar_spec.go                    # AR specification metadata
│   │   │   ├── errors.go                     # Attachment errors
│   │   │   ├── repository.go                 # Attachments repository interface
│   │   │   └── events.go                     # JobAttachmentAdded, JobAttachmentRemoved, JobVRPreviewAdded, JobARSpecAdded
│   │   │
│   │   # ========================= INVITATION & SOURCING DOMAINS =========================
│   │   ├── invitation/
│   │   │   ├── entity.go                     # Invitations to freelancers (JobID, FreelancerID, Message, Status, ExpiresAt)
│   │   │   ├── status.go                     # Invitation status (Pending, Accepted, Declined, Expired)
│   │   │   ├── errors.go                     # Invitation errors
│   │   │   ├── repository.go                 # JobInvitationRepository interface
│   │   │   └── events.go                     # JobInvitationSent, JobInvitationAccepted, JobInvitationDeclined, JobInvitationExpired
│   │   │
│   │   ├── sourcing/
│   │   │   ├── entity.go                     # Sourcing (JobID, Mode, PrivateLink, TalentPools[], Shortlists[])
│   │   │   ├── mode.go                       # Mode enum (Public, InviteOnly, PrivateLink) + constraints
│   │   │   ├── talent_pool.go                # Talent pool references
│   │   │   ├── shortlist.go                  # Shortlist management
│   │   │   ├── errors.go                     # Sourcing errors
│   │   │   ├── repository.go                 # SourcingRepository interface
│   │   │   └── events.go                     # SourcingModeSet, TalentPoolAttached, ShortlistUpdated
│   │   │
│   │   # ========================= BUDGET & VISIBILITY DOMAINS =========================
│   │   ├── budget_controls/
│   │   │   ├── entity.go                     # BudgetControl (JobID, MinBudget, MaxBudget, Currency, RateCapHourly)
│   │   │   ├── fx_rule.go                    # FX rules (quote vs settlement currency, rounding)
│   │   │   ├── errors.go                     # Budget control errors
│   │   │   ├── repository.go                 # BudgetControlRepository interface
│   │   │   └── events.go                     # BudgetControlSet, BudgetControlUpdated
│   │   │
│   │   ├── visibility_lifecycle/
│   │   │   ├── entity.go                     # VisibilityLifecycle (JobID, ScheduledAt, AutoCloseAt, RenewalPolicy, Draft, PublishedAt)
│   │   │   ├── policy.go                     # Rules engine structs for lifecycle transitions
│   │   │   ├── errors.go                     # Visibility/lifecycle errors
│   │   │   ├── repository.go                 # VisibilityLifecycleRepository interface
│   │   │   └── events.go                     # JobScheduled, JobPublished, JobAutoClosed, JobExtended
│   │   │
│   │   # ========================= TEMPLATE DOMAIN (CONSOLIDATED) =========================
│   │   ├── template/                         # 🔄 CONSOLIDATED: job_template + template_versions
│   │   │   ├── entity.go                     # Template aggregate (TemplateID, Title, Type, DefaultBudget, DefaultScope, Skills, Attachments, Versions[])
│   │   │   │                                 # Single aggregate WITH versioning
│   │   │   ├── type.go                       # JobType definitions (Fixed, Hourly, Retainer) + validation
│   │   │   ├── versions/
│   │   │   │   ├── version.go                # TemplateVersion (TemplateID, Version, Changelog, Deprecated, CreatedAt)
│   │   │   │   ├── semver.go                 # Semantic versioning logic
│   │   │   │   └── deprecation.go            # Deprecation rules
│   │   │   ├── errors.go                     # Template errors (TemplateNotFound, InvalidTemplateType, VersionConflict, TemplateDeprecated)
│   │   │   ├── repository.go                 # TemplateRepository interface (includes version operations)
│   │   │   └── events.go                     # JobTemplateCreated, JobTemplateUpdated, JobTemplateArchived, JobTemplateVersionCreated, JobTemplateVersionDeprecated
│   │   │
│   │   # ========================= ELIGIBILITY & REQUIREMENTS DOMAINS =========================
│   │   ├── eligibility_rules/
│   │   │   ├── entity.go                     # Eligibility (JobID, GeoRules, KYCRequired, MinTier, AgencyAllowed)
│   │   │   ├── enums.go                      # GateType, RuleEffect, Decision
│   │   │   ├── geo_rules.go                  # Geographic allow/deny lists + TZ overlap + radius (merged)
│   │   │   ├── errors.go                     # ErrApplicantIneligible, ErrRuleConflict
│   │   │   ├── repository.go                 # EligibilityRulesRepository interface (SetRules, GetRules, EvaluateApplicant)
│   │   │   └── events.go                     # JobEligibilityRulesSet, JobApplicantBlocked
│   │   │
│   │   ├── requirements_matrix/
│   │   │   ├── entity.go                     # Requirements (JobID, Items[Key, MustHave, Weight], Version)
│   │   │   ├── item.go                       # Requirement item (must-have vs nice-to-have)
│   │   │   ├── errors.go                     # ErrInvalidWeighting, ErrDuplicateKey
│   │   │   ├── repository.go                 # RequirementsMatrixRepository interface (SetMatrix, GetMatrix, ScoreProfile)
│   │   │   └── events.go                     # JobRequirementsMatrixSet, JobRequirementWeightUpdated
│   │   │
│   │   # ========================= HIRING TEAM DOMAIN =========================
│   │   ├── hiring_team/
│   │   │   ├── entity.go                     # HiringTeam (JobID, Members[UserID, Role], Notes)
│   │   │   ├── enums.go                      # Role: OWNER, REVIEWER, INTERVIEWER
│   │   │   ├── member.go                     # Team member details
│   │   │   ├── errors.go                     # ErrNotTeamMember, ErrRoleExists, ErrRoleChangeForbidden
│   │   │   ├── repository.go                 # HiringTeamRepository interface (AddMember, RemoveMember, ChangeRole, GetTeam)
│   │   │   └── events.go                     # JobHiringTeamMemberAdded, JobHiringTeamMemberRemoved, JobHiringTeamRoleChanged
│   │   │
│   │   # ========================= ANALYTICS DOMAIN (CONSOLIDATED) =========================
│   │   ├── analytics/                        # 🔄 CONSOLIDATED: job_view + job_analytics
│   │   │   ├── entity.go                     # Analytics (JobID, ViewCount, ProposalCount, InviteCount, ShortlistCount, InterviewCount, OfferCount, HireCount)
│   │   │   │                                 # LIGHTWEIGHT counters ONLY
│   │   │   │                                 # Detailed analytics (view tracking, engagement, response time) → search-be
│   │   │   ├── counters.go                   # Counter management (increment, reset)
│   │   │   ├── health_score.go               # Job health score calculation
│   │   │   ├── errors.go                     # ErrAnalyticsNotAvailable, ErrStatsStale
│   │   │   ├── repository.go                 # AnalyticsRepository interface (UpdateStats, GetStats, RecomputeScore)
│   │   │   └── events.go                     # JobAnalyticsUpdated, JobLowEngagementAlert
│   │   │                                     # ALSO emits detailed events to search-be: JobViewed, JobEngagementRecorded
│   │   │
│   │   # ========================= MODERATION DOMAIN (CONSOLIDATED) =========================
│   │   ├── moderation/                       # 🔄 CONSOLIDATED: job_flag + moderation_lifecycle
│   │   │   ├── entity.go                     # Moderation aggregate (JobID, Flags[], ModerationState, StateHistory[], Actions[])
│   │   │   │                                 # Single aggregate for flags + state machine
│   │   │   ├── flags/
│   │   │   │   ├── flag.go                   # Flag (JobID, ReporterID, Reason, Status, ReviewedAt)
│   │   │   │   ├── reason.go                 # Flag reasons (Spam, Inappropriate, Fraud, Misleading)
│   │   │   │   └── status.go                 # Flag status (Pending, Reviewed, Resolved, Dismissed)
│   │   │   ├── state/
│   │   │   │   ├── state.go                  # ModerationState (State, Reasons[], Actions[], Until)
│   │   │   │   ├── enums.go                  # State: CLEAN, LIMITED, QUARANTINED
│   │   │   │   └── transition.go             # State transition rules
│   │   │   ├── errors.go                     # ErrModerationStateLocked, ErrInvalidModerationTransition, ErrFlagNotFound
│   │   │   ├── repository.go                 # ModerationRepository interface (single repo for flags + state)
│   │   │   └── events.go                     # JobFlagSubmitted, JobFlagResolved, JobFlagDismissed, JobModerationStateApplied, JobModerationStateLifted
│   │   │
│   │   # ========================= A/B TESTING DOMAIN =========================
│   │   ├── ab_experiments/
│   │   │   ├── entity.go                     # Experiment (JobID, Arms[], Allocation, Guardrails, AssignmentSalt)
│   │   │   ├── arm.go                        # Experiment arm (variant configuration)
│   │   │   ├── allocation.go                 # Traffic allocation rules
│   │   │   ├── errors.go                     # ErrExperimentActive, ErrInvalidAllocation
│   │   │   ├── repository.go                 # ABExperimentsRepository interface (Start, Stop, AssignArm, Get)
│   │   │   └── events.go                     # JobExperimentStarted, JobExperimentStopped
│   │   │
│   │   # ========================= SYNDICATION & DRAFTS DOMAINS =========================
│   │   ├── syndication/
│   │   │   ├── entity.go                     # Syndication (JobID, Partner, Status, ExternalPostID, LastError, Retries)
│   │   │   ├── enums.go                      # Status: PENDING, POSTED, FAILED, TAKEDOWN
│   │   │   ├── partner.go                    # Partner configuration
│   │   │   ├── errors.go                     # ErrPartnerConfigMissing, ErrPartnerQuotaExceeded
│   │   │   ├── repository.go                 # SyndicationRepository interface (QueuePost, MarkPosted, MarkFailed, Takedown)
│   │   │   └── events.go                     # JobSyndicated, JobSyndicationFailed, JobSyndicationTakedown
│   │   │
│   │   ├── drafts/
│   │   │   ├── entity.go                     # Draft (DraftID, JobID?, ClientID, Snapshot, LastEditorID, UpdatedAt)
│   │   │   ├── snapshot.go                   # Draft snapshot management
│   │   │   ├── errors.go                     # ErrDraftNotFound, ErrStaleDraftVersion
│   │   │   ├── repository.go                 # DraftsRepository interface (Create, Update, Get, List, RestoreToJob)
│   │   │   └── events.go                     # JobDraftSaved, JobDraftRestored
│   │   │
│   │   # ========================= DUPLICATE DETECTION DOMAIN =========================
│   │   ├── duplicate_detection/
│   │   │   ├── entity.go                     # DuplicateKey (JobID, SimHash, ClusterID, FirstSeenAt)
│   │   │   ├── simhash.go                    # SimHash calculation logic
│   │   │   ├── clustering.go                 # Clustering algorithm
│   │   │   ├── errors.go                     # ErrDuplicateDetected, ErrHashUnavailable
│   │   │   ├── repository.go                 # DuplicateDetectionRepository interface (UpsertHash, FindNearDuplicates)
│   │   │   └── events.go                     # JobDuplicateDetected, JobDuplicateMerged
│   │   │
│   │   # ========================= LEGAL & COMPLIANCE DOMAIN (CONSOLIDATED) =========================
│   │   ├── legal_controls/                   # 🔄 CONSOLIDATED: legal_controls + data_residency
│   │   │   ├── entity.go                     # Legal (JobID, ExportControlFlag, NDATemplateID, NDAVersion, LegalHold, DataResidencyPolicy)
│   │   │   ├── export_control.go             # Export control flags
│   │   │   ├── nda.go                        # NDA pinning (template ID, version)
│   │   │   ├── legal_hold.go                 # Legal hold (immutable once placed)
│   │   │   ├── data_residency.go             # Data residency policy (EU, US, etc.)
│   │   │   ├── errors.go                     # ErrNDAVersionMismatch, ErrHoldAlreadyPlaced, ErrResidencyConflict
│   │   │   ├── repository.go                 # LegalControlsRepository interface (SetLegal, GetLegal, PlaceHold, RemoveHold, SetResidency)
│   │   │   └── events.go                     # JobLegalControlsSet, JobLegalHoldPlaced, JobLegalHoldRemoved, JobDataResidencySet
│   │   │
│   │   # ========================= CAMPAIGN & RETENTION DOMAINS =========================
│   │   ├── campaign_tags/
│   │   │   ├── entity.go                     # CampaignTags (JobID, Tags[], AddedBy, AddedAt)
│   │   │   ├── tag.go                        # Tag management (normalization, validation)
│   │   │   ├── errors.go                     # ErrTagLimitExceeded, ErrTagInvalid
│   │   │   ├── repository.go                 # CampaignTagsRepository interface (AddTag, RemoveTag, ListTags)
│   │   │   └── events.go                     # JobCampaignTagAdded, JobCampaignTagRemoved
│   │   │
│   │   ├── retention_rules/
│   │   │   ├── entity.go                     # Retention (JobID, ArchiveAt, PurgeAt, ExemptReason)
│   │   │   ├── policy.go                     # Retention policy rules
│   │   │   ├── errors.go                     # ErrRetentionConflict, ErrPurgeBlocked
│   │   │   ├── repository.go                 # RetentionRulesRepository interface (SetRetention, GetRetention)
│   │   │   └── events.go                     # JobRetentionSet, JobArchived, JobPurged
│   │   │
│   │   # ========================= PROMOTION & JOB PREFERENCES DOMAINS =========================
│   │   ├── promotion/
│   │   │   ├── entity.go                     # Promotion (JobID, Status, BadgeType, StartAt, EndAt, FeeAmount, RenewalCount)
│   │   │   ├── enums.go                      # Status: PENDING, ACTIVE, SUSPENDED, EXPIRED
│   │   │   ├── badge_type.go                 # Badge types (Featured, Urgent, TopJob)
│   │   │   ├── errors.go                     # ErrPromotionNotEligible, ErrMaxRenewalsExceeded, ErrPromoSuspended
│   │   │   ├── repository.go                 # PromotionRepository interface (ActivatePromotion, RenewPromotion, Suspend, Expire, GetByJob)
│   │   │   └── events.go                     # JobPromotionActivated, JobPromotionExpired, JobPromotionRenewed
│   │   │
│   │   ├── job_preference/                   # 🔄 RENAMED FROM preference/ — Advisory preferences ONLY (non-blocking soft signals)
│   │   │   ├── entity.go                     # Preference (JobID, FreelancerType, PreferredLocations, TimeZones, MinSuccessScore, FluencyLevel, MinPlatformEarnings, GuidanceLevel, ToolProvision)
│   │   │   ├── enums.go                      # FreelancerType, FluencyLevel, GuidanceLevel, ToolProvision
│   │   │   ├── errors.go                     # ErrInvalidPreference, ErrPreferenceConflict
│   │   │   ├── repository.go                 # PreferenceRepository interface (SetPreferences, Get, Update, Remove)
│   │   │   └── events.go                     # JobPreferencesSet, JobPreferencesUpdated, JobPreferencesRemoved
│   │   │
│   │   # ========================= MULTI-HIRE & REPOST DOMAIN =========================
│   │   ├── hiring_option/
│   │   │   ├── entity.go                     # HiringOption (JobID, MultiHireEnabled, MaxHires, OpenSlots, RepostAllowed, RepostReason, DuplicateCheckHash)
│   │   │   ├── multi_hire.go                 # Multi-hire configuration
│   │   │   ├── repost.go                     # Repost configuration & deduplication
│   │   │   ├── errors.go                     # ErrMultiHireLimitExceeded, ErrJobDuplicateDetected
│   │   │   ├── repository.go                 # HiringOptionRepository interface (EnableMultiHire, ReserveSlot, ReleaseSlot, RepostJob, CheckDuplicate)
│   │   │   └── events.go                     # JobMultiHireEnabled, JobReposted, JobDuplicatePrevented
│   │   │
│   │   # ========================= AI ASSIST DOMAIN =========================
│   │   ├── ai_assist/
│   │   │   ├── entity.go                     # SuggestionSet (JobID, SkillSuggestions[], CategorySuggestions[], OptimizationFeedback)
│   │   │   ├── suggestion.go                 # Suggestion details (skill, category, description)
│   │   │   ├── optimization.go               # Optimization feedback
│   │   │   ├── errors.go                     # ErrNoSuggestions, ErrOptimizationConflict
│   │   │   ├── repository.go                 # AIAssistRepository interface (Store accepted suggestions, optimization diffs)
│   │   │   └── events.go                     # JobAISuggestionsAccepted, JobDescriptionOptimized
│   │   │
│   │   # ========================= INCLUSIVITY DOMAIN =========================
│   │   ├── inclusivity/
│   │   │   ├── entity.go                     # InclusivityFlags (JobID, FlexibleHours, NoVideoRequired, ScreenReaderFriendly, AsynchronousWork, NeurodiversityFriendly)
│   │   │   ├── flags.go                      # Flag definitions and descriptions
│   │   │   ├── errors.go                     # ErrInvalidInclusivityFlag
│   │   │   ├── repository.go                 # InclusivityRepository interface (SetFlags, GetFlags)
│   │   │   └── events.go                     # JobInclusivityFlagsSet
│   │   │
│   │   # ========================= CONTRACT TRANSITION DOMAIN =========================
│   │   ├── contract_transition/              # 🔄 Job → contracts-be transition mapping
│   │   │   ├── entity.go                     # TransitionRequest (JobID, ContractSeedData, Status, CreatedAt, CompletedAt)
│   │   │   ├── mapping.go                    # Mapping (JobID → ContractSeed)
│   │   │   ├── errors.go                     # ErrTransitionNotReady, ErrMappingInvalid
│   │   │   ├── repository.go                 # ContractTransitionRepository interface (Queue, MarkSucceeded, MarkFailed, GetByJob)
│   │   │   └── events.go                     # JobContractTransitioned
│   │   │
│   │   # ========================= FRAUD DETECTION DOMAIN =========================
│   │   ├── fraud_detection/
│   │   │   ├── entity.go                     # RiskSignal (JobID, SignalType, Score, Rule, DetectedAt), AutoFlag
│   │   │   ├── signal.go                     # Signal types and scoring
│   │   │   ├── auto_flag.go                  # Auto-flag rules
│   │   │   ├── errors.go                     # ErrRiskRuleViolation
│   │   │   ├── repository.go                 # FraudDetectionRepository interface (UpsertSignals, GetSignals, AutoFlagJob)
│   │   │   └── events.go                     # JobFraudFlagged
│   │   │
│   │   # ========================= ESG DOMAIN =========================
│   │   ├── esg/
│   │   │   ├── entity.go                     # ESG (JobID, ESGFlags, CarbonEstimate)
│   │   │   ├── flags.go                      # ESG flags (RemoteFirst, LocalHire, SustainableTools, DiversityCommitment)
│   │   │   ├── carbon_estimate.go            # Carbon estimate snapshot
│   │   │   ├── errors.go                     # ErrInvalidESGFlag
│   │   │   ├── repository.go                 # ESGRepository interface (SetFlags, SetEstimate, Get)
│   │   │   └── events.go                     # JobESGFlagsSet, JobESGEstimateCalculated
│   │   │
│   │   # ========================= SHARING DOMAIN =========================
│   │   ├── sharing/
│   │   │   ├── entity.go                     # TrackedShareLink (JobID, LinkID, UTMParams, Incentive, CreatedAt)
│   │   │   ├── tracked_link.go               # Tracked link generation with UTM
│   │   │   ├── incentive.go                  # Referral incentive configuration
│   │   │   ├── errors.go                     # ErrShareQuotaExceeded
│   │   │   ├── repository.go                 # SharingRepository interface (GenerateLink, SetIncentive, GetByJob)
│   │   │   └── events.go                     # JobSharingLinkGenerated, JobReferralIncentiveSet
│   │   │
│   │   # ========================= BULK OPERATIONS DOMAIN =========================
│   │   ├── bulk_ops/
│   │   │   ├── entity.go                     # BulkUpdateBatch (BatchID, JobIDs[], Operations[], Status), ImportBatch
│   │   │   ├── batch.go                      # Batch operation management
│   │   │   ├── import.go                     # CSV import handling
│   │   │   ├── errors.go                     # ErrBatchConflict, ErrImportParse
│   │   │   ├── repository.go                 # BulkOpsRepository interface (CreateBatch, Append, MarkDone, Get)
│   │   │   └── events.go                     # JobBulkUpdated, JobBulkImported
│   │   │
│   │   # ========================= WEBHOOKS DOMAIN =========================
│   │   ├── webhooks/
│   │   │   ├── entity.go                     # WebhookSubscription (JobID, EndpointURL, Events[], DeliveryPolicy, Secret)
│   │   │   ├── delivery_policy.go            # Delivery policy (retries, timeouts)
│   │   │   ├── errors.go                     # ErrInvalidEndpoint, ErrHMACRequired
│   │   │   ├── repository.go                 # WebhooksRepository interface (Subscribe, Unsubscribe, List)
│   │   │   └── events.go                     # JobWebhookSubscribed
│   │   │
│   │   # ========================= ARCHIVE DOMAIN =========================
│   │   ├── archive/
│   │   │   ├── entity.go                     # ArchiveRecord (JobID, Reason, Actor, ArchivedAt, ReactivatedAt)
│   │   │   ├── reason.go                     # Archive reasons
│   │   │   ├── errors.go                     # ErrAlreadyArchived, ErrNotArchived
│   │   │   ├── repository.go                 # ArchiveRepository interface (Archive, Reactivate, GetHistory)
│   │   │   └── events.go                     # JobArchived, JobReactivated
│   │   │
│   │   # ========================= CUSTOM FIELDS DOMAIN =========================
│   │   ├── custom_fields/
│   │   │   ├── entity.go                     # CustomFieldSchema (JobID, FieldDefinitions[]), FieldValue
│   │   │   ├── schema.go                     # Field schema definitions
│   │   │   ├── value.go                      # Field value management
│   │   │   ├── errors.go                     # ErrSchemaConflict, ErrValueInvalid
│   │   │   ├── repository.go                 # CustomFieldsRepository interface (AddField, RemoveField, SetValue, GetValues)
│   │   │   └── events.go                     # JobCustomFieldAdded
│   │   │
│   │   # ========================= LOCALIZATION DOMAIN =========================
│   │   └── localization/
│   │       ├── entity.go                     # LocalizedDescription (JobID, Locale, Title, Summary, Body, PrimaryLocale)
│   │       ├── locale.go                     # Locale management
│   │       ├── errors.go                     # ErrLocaleInvalid, ErrPrimaryLocaleMissing
│   │       ├── repository.go                 # LocalizationRepository interface (UpsertLocale, RemoveLocale, GetLocales)
│   │       └── events.go                     # JobLanguagesSet
│   │
│   ├── application/                          # 📋 Application Layer - Use cases & orchestration (Load Fourth)
│   │   │
│   │   ├── authz/                            🆕  # service-level permission checks (defense-in-depth)
│   │   │   ├── policies.go                   🆕  # map roles → actions (admin, owner, same-user)
│   │   │   └── guards.go                     🆕  # helpers used by services
│   │   │
│   │   # ========================= EVENT HANDLERS (Inbound Events) =========================
│   │   ├── eventhandler/
│   │   │   ├── _routing_notes.go             # Partition keys: job events → key by job_id; user events → user_id
│   │   │   ├── user_handler.go               # Consumes: user.created → track clients for jobs domain
│   │   │   ├── proposal_handler.go           # Consumes: proposal.submitted → update job proposal counters
│   │   │   ├── admin_handler.go              # Consumes: admin.feature_flag.updated, admin.moderation.action.applied, admin.user_suspended
│   │   │   ├── subscription_handler.go       # Consumes: subscriptions.entitlement.updated, usage.counter.incremented, usage.limit.reached
│   │   │   ├── financial_handler.go          # Consumes: financial.invoice.overdue, financial.fee.schedule.updated
│   │   │   └── search_handler.go             # Consumes: search.taxonomy.synonym.updated → refresh skills/categories mapping
│   │   │
│   │   # ========================= CORE JOB SERVICES =========================
│   │   ├── job/
│   │   │   ├── service.go                    # Job business logic (Create, Update, Publish, Close, Repost)
│   │   │   ├── commands.go                   # CreateJob, UpdateJob, PublishJob, CloseJob (validates client exists via users-be)
│   │   │   ├── queries.go                    # GetJob, ListJobs, SearchJobs with filters (category, budget, skills)
│   │   │   ├── dto.go                        # JobDTO, CreateJobDTO, UpdateJobDTO, JobListDTO, JobSearchDTO
│   │   │   ├── mapper.go                     # Entity-DTO mapping
│   │   │   ├── validators.go                 # Input validation (title, type, budget ranges, dates)
│   │   │   └── business_rules.go             # Business rules (e.g., can close job only with accepted proposal)
│   │   │
│   │   ├── category/
│   │   │   ├── service.go                    # Category business logic (List, GetTree, Search)
│   │   │   ├── commands.go                   # UpsertCategory, ReorderCategories
│   │   │   ├── queries.go                    # ListCategories, GetCategoryTree
│   │   │   ├── dto.go                        # CategoryDTO, CategoryTreeDTO
│   │   │   ├── mapper.go                     # Category mappers
│   │   │   └── validators.go                 # Validate names/slugs, parent-child cycles
│   │   │
│   │   ├── skill/
│   │   │   ├── service.go                    # Skill business logic (List, Search, GetPopular) - SINGLE SOURCE
│   │   │   ├── commands.go                   # UpsertSkill, RetireSkill
│   │   │   ├── queries.go                    # ListSkills, PopularSkills
│   │   │   ├── dto.go                        # SkillDTO, SkillListDTO
│   │   │   ├── mapper.go                     # Skill mappers
│   │   │   └── validators.go                 # Validate names, categories, duplicates
│   │   │
│   │   # ========================= SCREENING SERVICE (CONSOLIDATED) =========================
│   │   ├── screening/                        # 🔄 CONSOLIDATED service
│   │   │   ├── service.go                    # Screening business logic (SetQuestions, SetSkillsTests, ToggleNDA, SetExportControl, SetCompliancePolicies)
│   │   │   ├── commands.go                   # SetScreeningQuestions, SetSkillsTests, ToggleNDA, SetExportControl, SetCompliancePolicies
│   │   │   ├── queries.go                    # GetScreeningPack, GetQuestions, GetCompliancePolicies
│   │   │   ├── dto.go                        # ScreeningPackDTO, QuestionDTO, CompliancePolicyDTO
│   │   │   ├── mapper.go                     # Screening mappers
│   │   │   └── validators.go                 # Validate question types, NDA toggles, export-control flags
│   │   │
│   │   # ========================= ATTACHMENTS SERVICE (CONSOLIDATED) =========================
│   │   ├── attachments/                      # 🔄 CONSOLIDATED service
│   │   │   ├── service.go                    # Attachments business logic (AddAttachment, RemoveAttachment, AddVRPreview, AddARSpec)
│   │   │   │                                 # Coordinates with storage-be for actual file uploads
│   │   │   ├── commands.go                   # AddAttachment, RemoveAttachment, AddVRPreview, AddARSpec
│   │   │   ├── queries.go                    # GetAttachments, GetPreviews
│   │   │   ├── dto.go                        # AttachmentDTO, VRPreviewDTO, ARSpecDTO
│   │   │   ├── mapper.go                     # Attachments mappers
│   │   │   └── validators.go                 # Validate file types, URLs, preview formats
│   │   │
│   │   # ========================= INVITATION & SOURCING SERVICES =========================
│   │   ├── invitation/
│   │   │   ├── service.go                    # Invitation business logic (Send, Accept, Decline)
│   │   │   ├── commands.go                   # SendInvitation, AcceptInvitation, DeclineInvitation
│   │   │   ├── queries.go                    # ListInvitationsForJob, GetInvitation
│   │   │   ├── dto.go                        # InvitationDTO, SendInvitationDTO
│   │   │   ├── mapper.go                     # Invitation mappers
│   │   │   └── validators.go                 # Validate recipient, message length, job state
│   │   │
│   │   ├── sourcing/
│   │   │   ├── service.go                    # Sourcing business logic (SetMode, ManageShortlists, ManageTalentPools)
│   │   │   ├── commands.go                   # SetMode, AddToShortlist, RemoveFromShortlist, AddPool, RemovePool
│   │   │   ├── queries.go                    # GetSourcing, ListShortlists, ListPools
│   │   │   ├── dto.go                        # SourcingDTO, ShortlistDTO, PoolDTO
│   │   │   ├── mapper.go                     # Sourcing mappers
│   │   │   └── validators.go                 # Validate mode transitions, link formats, list sizes
│   │   │
│   │   # ========================= BUDGET & VISIBILITY SERVICES =========================
│   │   ├── budget_controls/
│   │   │   ├── service.go                    # Budget business logic (SetRanges, NormalizeCurrency, EnforceHourlyCaps)
│   │   │   ├── commands.go                   # SetBudgetRange, SetCurrency, SetHourlyRateCap
│   │   │   ├── queries.go                    # GetBudgetControls, GetFxRules
│   │   │   ├── dto.go                        # BudgetControlsDTO, FxRuleDTO
│   │   │   ├── mapper.go                     # Budget mappers
│   │   │   └── validators.go                 # Validate min/max, ISO currency, non-negative rate caps
│   │   │
│   │   ├── visibility_lifecycle/
│   │   │   ├── service.go                    # Visibility business logic (Schedule, Publish, Unpublish, AutoClose, Extend, Renew)
│   │   │   ├── commands.go                   # SchedulePosting, Publish, Unpublish, SetAutoClose, ExtendPosting
│   │   │   ├── queries.go                    # GetVisibilityLifecycle
│   │   │   ├── dto.go                        # VisibilityLifecycleDTO
│   │   │   ├── mapper.go                     # Visibility/Lifecycle mappers
│   │   │   └── validators.go                 # Validate schedule times, state transitions, auto-close rules
│   │   │
│   │   # ========================= TEMPLATE SERVICE (CONSOLIDATED) =========================
│   │   ├── template/
│   │   │   ├── service.go                    # Template business logic (CreateTemplate, UpdateTemplate, ArchiveTemplate, CreateVersion, DeprecateVersion, CloneToJob)
│   │   │   ├── commands.go                   # CreateTemplate, UpdateTemplate, ArchiveTemplate, CreateTemplateVersion, DeprecateTemplateVersion
│   │   │   ├── queries.go                    # GetTemplate, ListTemplates, GetLatestVersion, ListTemplateVersions, ListByType
│   │   │   ├── dto.go                        # JobTemplateDTO, TemplateVersionDTO, SemVerDTO
│   │   │   ├── mapper.go                     # Template mappers (includes version mapping)
│   │   │   └── validators.go                 # Validate type, required defaults, semver rules, changelog
│   │   │
│   │   # ========================= ELIGIBILITY & REQUIREMENTS SERVICES =========================
│   │   ├── eligibility_rules/
│   │   │   ├── service.go                    # Eligibility business logic (SetRules, EvaluateApplicant)
│   │   │   ├── commands.go                   # SetRules, EvaluateApplicant
│   │   │   ├── queries.go                    # GetRules, GetRuleHistory
│   │   │   ├── dto.go                        # EligibilityRulesDTO, EvaluationResultDTO (includes Geo/TZ/Radius fields)  # UPDATED to absorb geo
│   │   │   ├── mapper.go                     # ISO country/tz normalization
│   │   │   └── validators.go                 # Conflicts & ranges
│   │   │
│   │   ├── requirements_matrix/
│   │   │   ├── service.go                    # Requirements business logic (SetMatrix, UpdateWeight, ScoreCandidatePreview)
│   │   │   ├── commands.go                   # SetMatrix, UpdateWeight
│   │   │   ├── queries.go                    # GetMatrix, ScoreCandidatePreview
│   │   │   ├── dto.go                        # RequirementDTO, ScoreDTO
│   │   │   ├── mapper.go                     # Flattened arrays ↔ entity items
│   │   │   └── validators.go                 # Weight sum ≤1; unique keys
│   │   │
│   │   # ========================= HIRING TEAM SERVICES =========================
│   │   ├── hiring_team/
│   │   │   ├── service.go                    # HiringTeam business logic (AddMember, ChangeRole, RemoveMember) - enforce OWNER invariants
│   │   │   ├── commands.go                   # AddMember, ChangeRole, RemoveMember
│   │   │   ├── queries.go                    # GetTeam
│   │   │   ├── dto.go                        # MemberDTO, RoleDTO
│   │   │   ├── mapper.go                     # UserID ↔ Member mapping
│   │   │   └── validators.go                 # No duplicate roles; can't remove last OWNER
│   │   │
│   │   # ========================= ANALYTICS SERVICE (CONSOLIDATED) =========================
│   │   ├── analytics/
│   │   │   ├── service.go                    # Analytics business logic (UpdateStats, RecomputeScore, EmitDetailedEvents)
│   │   │   ├── commands.go                   # UpdateStats, RecomputeScore, RecordView, RecordEngagement
│   │   │   ├── queries.go                    # GetStats, GetHealthScore
│   │   │   ├── dto.go                        # AnalyticsDTO, HealthScoreDTO
│   │   │   ├── mapper.go                     # Safe zeros; monotonic merges
│   │   │   └── validators.go                 # Non-negative counters; window checks
│   │   │
│   │   # ========================= MODERATION SERVICE (CONSOLIDATED) =========================
│   │   ├── moderation/
│   │   │   ├── service.go                    # Moderation business logic (FlagJob, UnflagJob, ResolveFlag, ApplyModerationState, LiftModerationState)
│   │   │   ├── commands.go                   # FlagJob, UnflagJob, ResolveFlag, ApplyModerationState, LiftModerationState
│   │   │   ├── queries.go                    # ListFlagsForJob, GetModerationState, GetHistory
│   │   │   ├── dto.go                        # FlagDTO, ModerationStateDTO
│   │   │   ├── mapper.go                     # Reason/action normalization
│   │   │   └── validators.go                 # Validate reasons/status transitions, state transition rules
│   │   │
│   │   # ========================= A/B TESTING SERVICE =========================
│   │   ├── ab_experiments/
│   │   │   ├── service.go                    # ABExperiments business logic (StartExperiment, StopExperiment, AssignArm)
│   │   │   ├── commands.go                   # StartExperiment, StopExperiment
│   │   │   ├── queries.go                    # GetExperiment, ListExperiments, AssignArm
│   │   │   ├── dto.go                        # ExperimentDTO, ArmDTO
│   │   │   ├── mapper.go                     # Percent ↔ decimal
│   │   │   └── validators.go                 # Sum=100; unique arms
│   │   │
│   │   # ========================= SYNDICATION & DRAFTS SERVICES =========================
│   │   ├── syndication/
│   │   │   ├── service.go                    # Syndication business logic (QueuePost, MarkPosted, MarkFailed, Takedown, HandleCallbacks)
│   │   │   ├── commands.go                   # QueueSyndication, MarkFailed, Takedown
│   │   │   ├── queries.go                    # GetStatus, ListPartners
│   │   │   ├── dto.go                        # SyndicationDTO
│   │   │   ├── mapper.go                     # Partner ids mapping
│   │   │   └── validators.go                 # Retry caps; config presence
│   │   │
│   │   ├── drafts/
│   │   │   ├── service.go                    # Drafts business logic (SaveDraft, RestoreToJob) - conflict detection
│   │   │   ├── commands.go                   # SaveDraft, RestoreToJob
│   │   │   ├── queries.go                    # GetDraft, ListDrafts
│   │   │   ├── dto.go                        # DraftDTO
│   │   │   ├── mapper.go                     # Snapshot merge
│   │   │   └── validators.go                 # Size/version checks
│   │   │
│   │   # ========================= DUPLICATE DETECTION SERVICE =========================
│   │   ├── duplicate_detection/
│   │   │   ├── service.go                    # DuplicateDetection business logic (UpsertSimhash, FindNearDuplicates)
│   │   │   ├── commands.go                   # UpsertDuplicateKey
│   │   │   ├── queries.go                    # FindNearDuplicates
│   │   │   ├── dto.go                        # DuplicateMatchDTO (score, ids)
│   │   │   ├── mapper.go                     # Hash/score formatting
│   │   │   └── validators.go                 # Hash len; threshold
│   │   │
│   │   # ========================= LEGAL & COMPLIANCE SERVICE (CONSOLIDATED) =========================
│   │   ├── legal_controls/
│   │   │   ├── service.go                    # LegalControls business logic (SetLegal, PlaceLegalHold, RemoveLegalHold, SetResidency)
│   │   │   ├── commands.go                   # SetLegal, PlaceLegalHold, RemoveLegalHold, SetResidency
│   │   │   ├── queries.go                    # GetLegal, GetResidencyPolicy
│   │   │   ├── dto.go                        # LegalDTO, ResidencyPolicyDTO
│   │   │   ├── mapper.go                     # Template version pinning
│   │   │   └── validators.go                 # NDA version exists; hold blocks purge; region compatibility
│   │   │
│   │   # ========================= CAMPAIGN & RETENTION SERVICES =========================
│   │   ├── campaign_tags/
│   │   │   ├── service.go                    # CampaignTags business logic (AddTag, RemoveTag)
│   │   │   ├── commands.go                   # AddTag, RemoveTag
│   │   │   ├── queries.go                    # ListTags
│   │   │   ├── dto.go                        # CampaignTagsDTO
│   │   │   ├── mapper.go                     # Normalize (lower, slug)
│   │   │   └── validators.go                 # Count ≤ limit; charset whitelist
│   │   │
│   │   ├── retention_rules/
│   │   │   ├── service.go                    # RetentionRules business logic (SetRetention, ComputeDeadlines)
│   │   │   ├── commands.go                   # SetRetention
│   │   │   ├── queries.go                    # GetRetention
│   │   │   ├── dto.go                        # RetentionDTO
│   │   │   ├── mapper.go                     # RFC3339 parsing
│   │   │   └── validators.go                 # Archive < Purge; future times
│   │   │
│   │   # ========================= PROMOTION & JOB PREFERENCES SERVICES =========================
│   │   ├── promotion/
│   │   │   ├── service.go                    # Promotion business logic (ActivatePromotion, RenewPromotion, SuspendPromotion, ExpirePromotion)
│   │   │   ├── commands.go                   # ActivatePromotion, RenewPromotion, SuspendPromotion, ExpirePromotion
│   │   │   ├── queries.go                    # GetPromotionByJob
│   │   │   ├── dto.go                        # PromotionDTO
│   │   │   ├── mapper.go                     # Money types; badge enum
│   │   │   └── validators.go                 # Windows; caps; transitions
│   │   │
│   │   ├── job_preference/                   # 🔄 RENAMED FROM preference/ — Advisory preferences service
│   │   │   ├── service.go                    # Preference business logic (SetPreferences, UpdatePreferences, RemovePreferences)
│   │   │   ├── commands.go                   # SetPreferences, UpdatePreferences, RemovePreferences
│   │   │   ├── queries.go                    # GetPreferences
│   │   │   ├── dto.go                        # PreferenceDTO
│   │   │   ├── mapper.go                     # ISO country/tz codes
│   │   │   └── validators.go                 # Enum bounds; conflicts
│   │   │
│   │   # ========================= MULTI-HIRE & REPOST SERVICE =========================
│   │   ├── hiring_option/
│   │   │   ├── service.go                    # HiringOption business logic (EnableMultiHire, ReserveSlot, ReleaseSlot, RepostJob)
│   │   │   ├── commands.go                   # EnableMultiHire, ReserveSlot, ReleaseSlot, RepostJob
│   │   │   ├── queries.go                    # GetHiringOption
│   │   │   ├── dto.go                        # HiringOptionDTO
│   │   │   ├── mapper.go                     # Slot math; dup hash
│   │   │   └── validators.go                 # MaxHires ≥ OpenSlots; cooldowns
│   │   │
│   │   # ========================= AI ASSIST SERVICE =========================
│   │   ├── ai_assist/
│   │   │   ├── service.go                    # AIAssist business logic (RunSuggestions, AcceptSelections, OptimizeDescription)
│   │   │   ├── commands.go                   # AcceptSuggestions, OptimizeDescription
│   │   │   ├── queries.go                    # GetSuggestionsForJob
│   │   │   ├── dto.go                        # SuggestionDTO, OptimizationFeedbackDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mappers
│   │   │   └── validators.go                 # Validate suggestion payloads / diffs
│   │   │
│   │   # ========================= INCLUSIVITY SERVICE =========================
│   │   ├── inclusivity/
│   │   │   ├── service.go                    # Inclusivity business logic (SetFlags, GetFlags)
│   │   │   ├── commands.go                   # SetInclusivityFlags
│   │   │   ├── queries.go                    # GetInclusivityFlags
│   │   │   ├── dto.go                        # InclusivityFlagsDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate flags
│   │   │
│   │   # ========================= CONTRACT TRANSITION SERVICE =========================
│   │   ├── contract_transition/
│   │   │   ├── service.go                    # ContractTransition business logic (StartTransition, MarkTransitionSucceeded, MarkTransitionFailed)
│   │   │   ├── commands.go                   # StartTransition, MarkTransitionSucceeded, MarkTransitionFailed
│   │   │   ├── queries.go                    # GetTransitionByJob
│   │   │   ├── dto.go                        # TransitionRequestDTO, MappingDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate readiness/mapping
│   │   │
│   │   # ========================= FRAUD DETECTION SERVICE =========================
│   │   ├── fraud_detection/
│   │   │   ├── service.go                    # FraudDetection business logic (RecordRiskSignal, AutoFlagJob)
│   │   │   ├── commands.go                   # RecordRiskSignal, AutoFlagJob
│   │   │   ├── queries.go                    # GetRiskSignals
│   │   │   ├── dto.go                        # RiskSignalDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate rule inputs
│   │   │
│   │   # ========================= ESG SERVICE =========================
│   │   ├── esg/
│   │   │   ├── service.go                    # ESG business logic (SetESGFlags, SetCarbonEstimate)
│   │   │   ├── commands.go                   # SetESGFlags, SetCarbonEstimate
│   │   │   ├── queries.go                    # GetESG
│   │   │   ├── dto.go                        # ESGFlagsDTO, CarbonEstimateDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate flags/units
│   │   │
│   │   # ========================= SHARING SERVICE =========================
│   │   ├── sharing/
│   │   │   ├── service.go                    # Sharing business logic (GenerateTrackedLink, SetReferralIncentive)
│   │   │   ├── commands.go                   # GenerateShareLink, SetReferralIncentive
│   │   │   ├── queries.go                    # ListShareLinks
│   │   │   ├── dto.go                        # ShareLinkDTO, IncentiveDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate UTM/incentives
│   │   │
│   │   # ========================= BULK OPERATIONS SERVICE =========================
│   │   ├── bulk_ops/
│   │   │   ├── service.go                    # BulkOps business logic (StartBulkUpdate, AppendBulkUpdate, StartImport)
│   │   │   ├── commands.go                   # StartBulkUpdate, AppendBulkUpdate, StartImport
│   │   │   ├── queries.go                    # GetBatchStatus
│   │   │   ├── dto.go                        # BulkUpdateBatchDTO, ImportBatchDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate batch sizes / CSV schema
│   │   │
│   │   # ========================= WEBHOOKS SERVICE =========================
│   │   ├── webhooks/
│   │   │   ├── service.go                    # Webhooks business logic (SubscribeWebhook, UnsubscribeWebhook)
│   │   │   ├── commands.go                   # SubscribeWebhook, UnsubscribeWebhook
│   │   │   ├── queries.go                    # ListWebhooks
│   │   │   ├── dto.go                        # WebhookSubscriptionDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate endpoint/HMAC
│   │   │
│   │   # ========================= ARCHIVE SERVICE =========================
│   │   ├── archive/
│   │   │   ├── service.go                    # Archive business logic (ArchiveJob, ReactivateJob)
│   │   │   ├── commands.go                   # ArchiveJob, ReactivateJob
│   │   │   ├── queries.go                    # GetArchiveHistory
│   │   │   ├── dto.go                        # ArchiveRecordDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate state transitions
│   │   │
│   │   # ========================= CUSTOM FIELDS SERVICE =========================
│   │   ├── custom_fields/
│   │   │   ├── service.go                    # CustomFields business logic (AddCustomField, RemoveCustomField, SetFieldValue)
│   │   │   ├── commands.go                   # AddCustomField, RemoveCustomField, SetFieldValue
│   │   │   ├── queries.go                    # ListCustomFields, GetFieldValues
│   │   │   ├── dto.go                        # CustomFieldSchemaDTO, FieldValueDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate schema/value types
│   │   │
│   │   # ========================= LOCALIZATION SERVICE =========================
│   │   └── localization/
│   │       ├── service.go                    # Localization business logic (UpsertJobLocale, RemoveJobLocale, SetPrimaryLocale)
│   │       ├── commands.go                   # UpsertJobLocale, RemoveJobLocale, SetPrimaryLocale
│   │       ├── queries.go                    # GetJobLocales
│   │       ├── dto.go                        # LocalizedDescriptionDTO
│   │       ├── mapper.go                     # Mappers
│   │       └── validators.go                 # Validate locale codes / primary locale
│   │
│   └── interfaces/                           # 🌐 Interface Layer (API/HTTP - Load Fifth)
│       └── http/
│           │
│           ├── middleware/
│           │   ├── requestid.go              # 🆕 Wraps platform-shared request id
│           │   ├── logging.go                # 🆕 Wraps platform-shared logging
│           │   ├── recovery.go               # 🆕 Wraps platform-shared recovery
│           │   ├── tracing.go                # 🆕 Wraps platform-shared otel middleware
│           │   ├── auth.go                   # 🆕 Wraps pkg/auth (Keycloak)
│           │   │                             # MFA verification delegated to Keycloak/pkg/auth
│           │   ├── rbac.go                   ♻️  # ensure aligns with application/authz
│           │   └── cors.go                   # 🆕 CORS middleware
│           │
│           ├── v1/                           🛑🚦 # NEW: versioned API surface
│           │   ├── openapi/
│           │   │   ├── openapi.yaml          # 📜 OpenAPI contract v1: request/response schemas; served as /v1/openapi
│           │   │   └── generator.go          # 🧭 Serves /swagger and /openapi.json (dev only toggle)
│           │   │
│           │   ├── presenters/
│           │   │   └── errors.go             # 🎯 Maps domain/service errors → HTTP status & problem+json
│           │   │
│           │   ├── handlers/                 🔄  # move all handlers here
│           │   │   # ========================= CORE JOB HANDLERS =========================
│           │   │   ├── job_handler.go        # POST, GET, PUT, DELETE /jobs, /jobs/:id
│           │   │   ├── health_handler.go     # 🆕 GET /healthz/live, GET /healthz/ready
│           │   │   # ========================= CATEGORY & SKILL HANDLERS =========================
│           │   │   ├── category_handler.go   # GET /categories, GET /categories/tree
│           │   │   ├── skill_handler.go      # GET /skills, GET /skills/popular
│           │   │   # ========================= CONSOLIDATED HANDLERS =========================
│           │   │   ├── screening_handler.go  # 🔄 CONSOLIDATED: questions + compliance
│           │   │   ├── attachments_handler.go# 🔄 CONSOLIDATED: attachments + previews (VR/AR)
│           │   │   ├── template_handler.go   # 🔄 CONSOLIDATED: templates + versions
│           │   │   ├── moderation_handler.go # 🔄 CONSOLIDATED: flags + state
│           │   │   ├── legal_controls_handler.go # 🔄 CONSOLIDATED: export + NDA + holds + residency
│           │   │   ├── analytics_handler.go  # 🔄 CONSOLIDATED: lightweight counters
│           │   │   # ========================= INVITATION & SOURCING HANDLERS =========================
│           │   │   ├── invitation_handler.go # POST, GET /jobs/:id/invitations
│           │   │   ├── sourcing_handler.go   # POST, GET /jobs/:id/sourcing
│           │   │   # ========================= BUDGET & VISIBILITY HANDLERS =========================
│           │   │   ├── budget_handler.go     # POST, GET /jobs/:id/budget-controls
│           │   │   ├── visibility_handler.go # POST, GET /jobs/:id/visibility
│           │   │   # ========================= ELIGIBILITY & REQUIREMENTS HANDLERS =========================
│           │   │   ├── eligibility_rules_handler.go   # POST, GET /jobs/:id/eligibility (includes Geo/TZ/Radius merged)
│           │   │   ├── requirements_matrix_handler.go # POST, GET /jobs/:id/requirements
│           │   │   # ========================= HIRING TEAM HANDLERS =========================
│           │   │   ├── hiring_team_handler.go# POST, DELETE, PUT /jobs/:id/team
│           │   │   # ========================= A/B TESTING HANDLER =========================
│           │   │   ├── ab_experiments_handler.go # POST, GET, DELETE /jobs/:id/experiments
│           │   │   # ========================= SYNDICATION & DRAFTS HANDLERS =========================
│           │   │   ├── syndication_handler.go# POST, GET /jobs/:id/syndication
│           │   │   ├── drafts_handler.go     # POST, GET /jobs/:id/drafts
│           │   │   # ========================= DUPLICATE DETECTION HANDLER =========================
│           │   │   ├── duplicate_detection_handler.go # GET /jobs/:id/duplicates
│           │   │   # ========================= CAMPAIGN & RETENTION HANDLERS =========================
│           │   │   ├── campaign_tags_handler.go      # POST, DELETE, GET /jobs/:id/tags
│           │   │   ├── retention_rules_handler.go    # POST, GET /jobs/:id/retention
│           │   │   # ========================= PROMOTION & JOB PREFERENCES HANDLERS =========================
│           │   │   ├── promotion_handler.go  # POST, PUT, GET /jobs/:id/promotion
│           │   │   ├── job_preference_handler.go # 🔄 RENAMED FROM preference_handler.go — POST, PUT, GET, DELETE /jobs/:id/preferences
│           │   │   # ========================= MULTI-HIRE & REPOST HANDLER =========================
│           │   │   ├── hiring_option_handler.go     # POST, GET /jobs/:id/hiring-option
│           │   │   # ========================= AI ASSIST HANDLER =========================
│           │   │   ├── ai_assist_handler.go        # POST, GET /jobs/:id/ai/suggestions, POST /jobs/:id/ai/optimize
│           │   │   # ========================= INCLUSIVITY HANDLER =========================
│           │   │   ├── inclusivity_handler.go      # POST, GET /jobs/:id/inclusivity
│           │   │   # ========================= CONTRACT TRANSITION HANDLER =========================
│           │   │   ├── contract_transition_handler.go # POST, GET /jobs/:id/contract/transition
│           │   │   # ========================= FRAUD DETECTION HANDLER =========================
│           │   │   ├── fraud_detection_handler.go   # POST, GET /jobs/:id/fraud
│           │   │   # ========================= ESG HANDLER =========================
│           │   │   ├── esg_handler.go              # POST, GET /jobs/:id/esg
│           │   │   # ========================= SHARING HANDLER =========================
│           │   │   ├── sharing_handler.go          # POST, GET /jobs/:id/sharing
│           │   │   # ========================= BULK OPERATIONS HANDLER =========================
│           │   │   ├── bulk_ops_handler.go         # POST, GET /jobs/bulk
│           │   │   # ========================= WEBHOOKS HANDLER =========================
│           │   │   ├── webhooks_handler.go         # POST, DELETE, GET /jobs/:id/webhooks
│           │   │   # ========================= ARCHIVE HANDLER =========================
│           │   │   ├── archive_handler.go          # POST /jobs/:id/archive, POST /jobs/:id/reactivate, GET /jobs/:id/archive/history
│           │   │   # ========================= CUSTOM FIELDS HANDLER =========================
│           │   │   ├── custom_fields_handler.go    # POST, DELETE, PUT, GET /jobs/:id/custom-fields
│           │   │   # ========================= LOCALIZATION HANDLER =========================
│           │   │   └── localization_handler.go     # PUT, DELETE, GET /jobs/:id/locales
│           │   │
│           │   └── routes/                  🔄  # move all routes here
│           │       # ========================= CORE JOB ROUTES =========================
│           │       ├── job_routes.go             # /jobs, /jobs/:id
│           │       # ========================= CATEGORY & SKILL ROUTES =========================
│           │       ├── category_routes.go        # /categories
│           │       ├── skill_routes.go           # /skills
│           │       # ========================= CONSOLIDATED ROUTES =========================
│           │       ├── screening_routes.go       # 🔄 /jobs/:id/screening
│           │       ├── attachments_routes.go     # 🔄 /jobs/:id/attachments
│           │       ├── template_routes.go        # 🔄 /templates
│           │       ├── moderation_routes.go      # 🔄 /jobs/:id/flag, /jobs/:id/moderation
│           │       ├── legal_controls_routes.go  # 🔄 /jobs/:id/legal
│           │       ├── analytics_routes.go       # 🔄 /jobs/:id/analytics
│           │       # ========================= INVITATION & SOURCING ROUTES =========================
│           │       ├── invitation_routes.go      # /jobs/:id/invitations
│           │       ├── sourcing_routes.go        # /jobs/:id/sourcing
│           │       # ========================= BUDGET & VISIBILITY ROUTES =========================
│           │       ├── budget_routes.go          # /jobs/:id/budget-controls
│           │       ├── visibility_routes.go      # /jobs/:id/visibility
│           │       # ========================= ELIGIBILITY & REQUIREMENTS ROUTES =========================
│           │       ├── eligibility_rules_routes.go   # /jobs/:id/eligibility (includes Geo/TZ/Radius merged)
│           │       ├── requirements_matrix_routes.go # /jobs/:id/requirements
│           │       # ========================= HIRING TEAM ROUTES =========================
│           │       ├── hiring_team_routes.go     # /jobs/:id/team
│           │       # ========================= A/B TESTING ROUTES =========================
│           │       ├── ab_experiments_routes.go  # /jobs/:id/experiments
│           │       # ========================= SYNDICATION & DRAFTS ROUTES =========================
│           │       ├── syndication_routes.go     # /jobs/:id/syndication
│           │       ├── drafts_routes.go          # /jobs/:id/drafts
│           │       # ========================= DUPLICATE DETECTION ROUTES =========================
│           │       ├── duplicate_detection_routes.go # /jobs/:id/duplicates
│           │       # ========================= CAMPAIGN & RETENTION ROUTES =========================
│           │       ├── campaign_tags_routes.go   # /jobs/:id/tags
│           │       ├── retention_rules_routes.go # /jobs/:id/retention
│           │       # ========================= PROMOTION & JOB PREFERENCES ROUTES =========================
│           │       ├── promotion_routes.go       # /jobs/:id/promotion
│           │       ├── job_preference_routes.go  # 🔄 RENAMED FROM preference_routes.go — /jobs/:id/preferences
│           │       # ========================= MULTI-HIRE & REPOST ROUTES =========================
│           │       ├── hiring_option_routes.go   # /jobs/:id/hiring-option
│           │       # ========================= AI ASSIST ROUTES =========================
│           │       ├── ai_assist_routes.go       # /jobs/:id/ai
│           │       # ========================= INCLUSIVITY ROUTES =========================
│           │       ├── inclusivity_routes.go     # /jobs/:id/inclusivity
│           │       # ========================= CONTRACT TRANSITION ROUTES =========================
│           │       ├── contract_transition_routes.go # /jobs/:id/contract/transition
│           │       # ========================= FRAUD DETECTION ROUTES =========================
│           │       ├── fraud_detection_routes.go # /jobs/:id/fraud
│           │       # ========================= ESG ROUTES =========================
│           │       ├── esg_routes.go             # /jobs/:id/esg
│           │       # ========================= SHARING ROUTES =========================
│           │       ├── sharing_routes.go         # /jobs/:id/sharing
│           │       # ========================= BULK OPERATIONS ROUTES =========================
│           │       ├── bulk_ops_routes.go        # /jobs/bulk
│           │       # ========================= WEBHOOKS ROUTES =========================
│           │       ├── webhooks_routes.go        # /jobs/:id/webhooks
│           │       # ========================= ARCHIVE ROUTES =========================
│           │       ├── archive_routes.go         # /jobs/:id/archive
│           │       # ========================= CUSTOM FIELDS ROUTES =========================
│           │       ├── custom_fields_routes.go   # /jobs/:id/custom-fields
│           │       # ========================= LOCALIZATION ROUTES =========================
│           │       └── localization_routes.go    # /jobs/:id/locales
│           │
│           ├── etag.go                      🆕  # middleware-level ETag (pairs with pkg/utils/etag)
│           └── router.go                     # 📝 HTTP router setup (Gin with platform-shared middleware)
│
├── db/                                      🆕  # developer-friendly entrypoint for SQL
│   └── migrations/                          🔄  # symlink or mirror of internal/.../migrations (optional)
│
├── config/                                   # 🆕 Configuration files
│   ├── default.yaml                          # Default configuration
│   ├── dev.yaml                              # Development overrides
│   └── prod.yaml                             # Production overrides
│
├── dapr/                                     # 🆕 Dapr components
│   ├── local/
│   │   ├── pubsub.yaml                       # Local Kafka pub/sub
│   │   └── statestore.yaml                   # Local state store
│   └── k8s/
│       ├── pubsub.yaml                       # Production Kafka with scopes: ["jobs-be"]
│       ├── statestore.yaml                   # Production state store with scopes: ["jobs-be"]
│       └── secrets.yaml                      # Secrets component
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                         # Service-specific errors
│   │   └── codes.go                          # Error codes (JOB_NOT_FOUND, INVALID_BUDGET, JOB_ALREADY_CLOSED)
│   │
│   ├── logger/                               # ❌ REMOVED
│   │   └── README.md                         # Points to platform-shared/logging
│   │
│   ├── utils/
│   │   ├── validator.go                      # Local validation utilities
│   │   ├── slug.go                           # Generate URL-friendly slugs
│   │   ├── sanitizer.go                      # Sanitize user input
│   │   └── etag.go                                 🆕  # canonical ETag helpers (hashing, weak/strong)
│   │
│   └── constants/
│       # events.go / topics.go removed previously (use contracts/events)
│       # 🔻 De-dup cleanup applied (source of truth is domain enums)
│       # job_types.go — REMOVED
│       # job_status.go — REMOVED
│       # visibility_modes.go — REMOVED
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                   # Kubernetes Deployment
│       ├── service.yaml                      # Kubernetes Service
│       ├── configmap.yaml                    # ConfigMap
│       ├── secrets.yaml                      # Secrets
│       ├── hpa.yaml                          # Horizontal Pod Autoscaler
│       ├── pdb.yaml                          # Pod Disruption Budget
│       └── servicemonitor.yaml               # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                        # Local dev setup
│   ├── get-secrets.sh                        # Fetch secrets
│   ├── seed-categories.sh                    # Seed job categories
│   ├── seed-skills.sh                        # Seed skills taxonomy
│   ├── migrate.sh                            🆕  # SQL migrations (dev/prod flags)
│   ├── openapi-diff.sh                       🆕  # fail CI on breaking API changes
│   ├── schema-diff.sh                        🆕  # DB schema guardrail (pg_dump + mig verify)
│   ├── generate-sdks.sh                      🆕  # produce /sdk clients from OpenAPI
│   └── sbom-sign.sh                          🆕  # build SBOM + cosign sign
│
├── tests/
│   ├── reliability/                          🆕
│   │   ├── projections_replay_test.go        🆕  # rebuild read models from event logs
│   │   └── outbox_dispatcher_test.go         🆕  # at-least-once + idempotency assertions
│   ├── property/                             🆕
│   │   └── scoring_property_test.go          🆕  # property/fuzz tests for scoring kernel
│   ├── e2e/
│   │   └── chaos_event_delivery_test.go      🆕  # simulate duplicates, reordering, delays
│   ├── unit/
│   │   ├── domain/
│   │   │   # ========================= CORE DOMAIN TESTS =========================
│   │   │   ├── job_test.go                   # Job entity tests
│   │   │   ├── category_test.go              # Category entity tests
│   │   │   ├── skill_test.go                 # Skill taxonomy tests
│   │   │   # ========================= CONSOLIDATED DOMAIN TESTS =========================
│   │   │   ├── screening_test.go             # 🔄 Screening tests (questions + compliance)
│   │   │   ├── attachments_test.go           # 🔄 Attachments tests (documents + previews)
│   │   │   ├── template_test.go              # 🔄 Template tests (templates + versions)
│   │   │   ├── moderation_test.go            # 🔄 Moderation tests (flags + state)
│   │   │   ├── legal_controls_test.go        # 🔄 Legal controls tests (export + NDA + residency)
│   │   │   ├── analytics_test.go             # 🔄 Analytics tests (lightweight counters)
│   │   │   # ========================= OTHER DOMAIN TESTS =========================
│   │   │   ├── sourcing_test.go              # Sourcing tests
│   │   │   ├── eligibility_rules_test.go     # Eligibility rules tests (includes geo/TZ/radius)  # UPDATED
│   │   │   └── duplicate_detection_test.go   # Duplicate detection tests
│   │   ├── application/
│   │   │   # ========================= CORE SERVICE TESTS =========================
│   │   │   ├── job_service_test.go           # Job service tests
│   │   │   ├── category_service_test.go      # Category service tests
│   │   │   ├── skill_service_test.go         # Skill service tests
│   │   │   # ========================= CONSOLIDATED SERVICE TESTS =========================
│   │   │   ├── screening_service_test.go     # 🔄 Screening service tests
│   │   │   ├── attachments_service_test.go   # 🔄 Attachments service tests
│   │   │   ├── template_service_test.go      # 🔄 Template service tests
│   │   │   ├── moderation_service_test.go    # 🔄 Moderation service tests
│   │   │   ├── legal_controls_service_test.go# 🔄 Legal controls service tests
│   │   │   └── analytics_service_test.go     # 🔄 Analytics service tests
│   │   ├── infrastructure/
│   │   │   ├── job_repository_test.go        # Job repository tests
│   │   │   ├── screening_repository_test.go  # 🔄 Screening repository tests
│   │   │   ├── attachments_repository_test.go# 🔄 Attachments repository tests
│   │   │   ├── template_repository_test.go   # 🔄 Template repository tests
│   │   │   ├── moderation_repository_test.go # 🔄 Moderation repository tests
│   │   │   ├── scheduler_tasks_test.go       # 🆕 Scheduler task tests
│   │   │   └── projections_test.go           # 🆕 Projection tests
│   │   ├── integration/
│   │   │   ├── handlers/
│   │   │   │   ├── job_handler_test.go           # Job handler integration tests
│   │   │   │   ├── screening_handler_test.go     # 🔄 Screening handler tests
│   │   │   │   ├── attachments_handler_test.go   # 🔄 Attachments handler tests
│   │   │   │   ├── template_handler_test.go      # 🔄 Template handler tests
│   │   │   │   └── moderation_handler_test.go    # 🔄 Moderation handler tests
│   │   │   ├── repositories/
│   │   │   │   ├── job_repository_integration_test.go        # Job repository integration tests
│   │   │   │   └── screening_repository_integration_test.go  # 🔄 Screening repository tests
│   │   │   └── worker_tasks_integration_test.go  # 🆕 Worker tasks integration tests
│   │
├── docs/
│   ├── README.md                             # Service overview
│   ├── API.md                                # API documentation
│   ├── ARCHITECTURE.md                       # Service architecture
│   ├── MIGRATIONS.md                         # Migration history
│   ├── SCHEMA.md                             # Database schema
│   ├── RUNBOOK.md                            # Operations runbook
│   # ========================= CORE FEATURE DOCUMENTATION =========================
│   ├── EVENTS.md                             # 🆕 Events (published & consumed)
│   ├── EVENT_VERSIONING.md                   # 🆕 Event versioning guide
│   ├── job-lifecycle.md                      # Job lifecycle documentation
│   ├── categories.md                         # Job categories documentation
│   ├── skills-taxonomy.md                    # Skills taxonomy documentation (SINGLE SOURCE)
│   # ========================= CONSOLIDATED FEATURE DOCUMENTATION =========================
│   ├── screening.md                          # 🔄 CONSOLIDATED: questions + compliance
│   ├── attachments.md                        # 🔄 CONSOLIDATED: documents + previews (VR/AR)
│   ├── templates.md                          # 🔄 CONSOLIDATED: templates + versions
│   ├── moderation.md                          # 🔄 CONSOLIDATED: flags + state machine
│   ├── legal-controls.md                     # 🔄 CONSOLIDATED: export + NDA + holds + residency
│   ├── analytics.md                          # 🔄 CONSOLIDATED: lightweight counters (detailed → search-be)
│   # ========================= OTHER FEATURE DOCUMENTATION =========================
│   ├── sourcing-modes.md                     # Sourcing modes documentation
│   ├── eligibility-rules.md                  # Eligibility rules documentation (now includes Geo/TZ/Radius)  # UPDATED
│   ├── hiring-team.md                        # Hiring team documentation
│   ├── duplicate-detection.md                # Duplicate detection documentation
│   ├── promotions.md                         # Promotions documentation
│   ├── contract-transition.md                # Contract transition documentation
│   # ========================= EXTERNAL SERVICE BOUNDARIES DOCUMENTATION =========================
│   ├── EXTERNAL_SERVICES.md                  # 🆕 Docs for all externalized services: search-be, reviews-be, contracts-be, financial-be, storage-be, users-be, proposals-be, Keycloak/pkg/auth
│   ├── API_VERSIONING.md                     🆕  # HTTP contract versioning & deprecation policy
│   ├── CACHING.md                            🆕  # keys, TTLs, SWR, invalidation events
│   ├── DATA_RETENTION.md                     🛡️🆕 # per-domain retention windows
│   ├── ERASURE.md                            🛡️🆕 # GDPR/CCPA erasure hooks & playbooks
│   ├── SLOS.md                               📏🆕 # target P99s, projection lag, queue delay
│   ├── OUTBOX.md                             🆕  # exactly-once-ish design, retries, DLQ
│   ├── RELEASE_CHECKLIST.md                  🆕  # preflight (openapi diff, schema diff, migrations plan)
│   ├── NAMING.md                             ♻️  # package snake_case; HTTP kebab-case conventions
│   └── migration-guide.md                    # 🆕 Guide for migrating from old structure to new lean structure
│
├── .github/
│   └── workflows/
│       ├── contract-ci.yml                   🆕  # openapi-diff + event schema checks
│       ├── security.yml                      🆕  # golangci-lint, govulncheck, trivy, cosign verify
│       ├── load-tests.yml                    🆕  # k6/gatling smoke against /healthz & hot paths
│       ├── ci.yml                            # CI workflow
│       └── cd.yml                            # CD workflow
│
├── sdk/                                     🆕  # generated clients (optional but handy)
│   ├── go/                                  🆕
│   └── ts/                                  🆕
│
├── go.mod                                    # 📝 Dependencies (imports pkg/auth, platform-shared, contracts/events)
├── go.sum
├── .env.example                              # Example environment variables
├── .golangci.yml                             🆕  # linter config (ci parity)
├── CODEOWNERS                                🆕  # explicit ownership per context (ease future splits)
├── Makefile                                  # Build targets
├── Dockerfile                                # Multi-stage Docker build
├── .dockerignore                             # Docker ignore patterns
├── .gitignore                                # Git ignore patterns
└── README.md                                 # Service README


```
---
---

## 📦 **2️⃣ proposals-be (Proposal Management Service) - REFACTORED**

apps/be/proposals-be/
│
├── cmd/
│   │── api/
│   │   └── main.go                           # 📝 Application entry point - initializes Gin, Dapr, Postgres (uses platform-shared/logging, internal/config)
│   └── worker/
│       └── main.go                 # runs cron/Dapr bindings for: expirations, follow-ups, health calc, pipeline sweeps, ♻️  # wire leader election + outbox dispatcher (in addition to cron)
│
├── internal/
│   │
│   ├── ioc/
│   │   ├── container.go                 # 🧩 Dependency graph: constructs DB/Redis/Kafka clients, repositories, services, handlers, schedulers
│   │   └── wiring.go                    # 🔌 Env-driven wiring & feature flags: selects implementations (e.g., local vs. cloud), toggles modules
│   │
│   ├── config/                               # 🔧 Configuration (Load First)
│   │   ├── feature_flags.go                        🆕  # e.g., enable_bidding, enable_auctions, enable_intelligence_facade
│   │   ├── schema.go                         # Typed Config struct (App, Server, Postgres, Kafka, Redis, Auth, Storage)
│   │   ├── loader.go                         # Config loader using Viper (CLI → ENV → file → defaults)
│   │   └── docs/
│   │       └── CONFIGURATION.md              # Configuration documentation
│   │
│   ├── infrastructure/                       # 🔧 Infrastructure Layer (Load Second)
│   │   │
│   │   ├── coordination/                           🆕  # distributed coordination primitives
│   │   │   ├── leader_election.go                  🆕  # single-active worker/cron instance
│   │   │   └── distlock.go                         🆕  # Redis/PG advisory locks for cron/tasks
│   │   │
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection setup
│   │   │       ├── transaction.go            # Transaction helpers
│   │   │       ├── migrations.go             # Auto-migration logic with version tracking
│   │   │       ├── version.go                # Schema version tracking
│   │   │       ├── safety.go                 # Pre-migration safety checks
│   │   │       ├── outbox_store.go                 🆕  # tx-coupled outbox table access
│   │   │       ├── transaction.go                  ♻️  # ensure outbox write is in same DB tx
│   │   │       │
│   │   │       │
│   │   │       # ========================= CORE PROPOSAL REPOSITORIES =========================
│   │   │       ├── proposal_repository.go    # Proposal CRUD (Create, Update, FindByID, List, Search)
│   │   │       ├── cover_letter_repository.go # Cover letter repository
│   │   │       ├── attachment_repository.go  # Proposal attachments repository
│   │   │       ├── question_answer_repository.go # Q&A repository
│   │   │       ├── milestone_repository.go   # Milestones repository
│   │   │       │
│   │   │       # ========================= LIFECYCLE REPOSITORIES =========================
│   │   │       ├── expiration_repository.go  # Expiration tracking repository
│   │   │       ├── withdrawal_repository.go  # Withdrawal repository
│   │   │       ├── archive_repository.go     # Archive repository
│   │   │       ├── pipeline_repository.go    # Pipeline stages repository
│   │   │       ├── recycling_repository.go   # Proposal recycling repository
│   │   │       │
│   │   │       # ========================= BIDDING REPOSITORIES =========================
│   │   │       ├── bid_repository.go         # Bid repository (amount, type, status)
│   │   │       ├── bid_strategy_repository.go # Bid strategy repository
│   │   │       ├── auction_repository.go     # Auction repository (if core mechanic)
│   │   │       ├── bid_anomaly_repository.go # Bid anomaly detection repository
│   │   │       │
│   │   │       # ========================= CONNECTS & CREDITS REPOSITORIES =========================
│   │   │       ├── connect_repository.go     # Connects/credits repository
│   │   │       ├── connect_refund_repository.go # Connect refund repository
│   │   │       │
│   │   │       # ========================= WORKFLOW & COLLABORATION REPOSITORIES =========================
│   │   │       ├── negotiation_repository.go # Negotiation threads repository
│   │   │       ├── invite_repository.go      # Invitation repository
│   │   │       ├── revision_repository.go    # Revision history repository
│   │   │       ├── collaboration_repository.go # Team collaboration repository (MERGED: team + collaboration_network)
│   │   │       │
│   │   │       # ========================= CLIENT INTERACTION REPOSITORIES =========================
│   │   │       ├── interview_repository.go   # Interview request repository
│   │   │       ├── feedback_repository.go    # Proposal feedback repository
│   │   │       ├── shortlist_repository.go   # Shortlist repository
│   │   │       ├── conversation_repository.go # Pre-hire messaging repository
│   │   │       │
│   │   │       # ========================= PERFORMANCE & ANALYTICS REPOSITORIES (CONSOLIDATED) =========================
│   │   │       ├── performance_repository.go # CONSOLIDATED: analytics + engagement + response_tracker + conversion + metrics + ranking + health
│   │   │       │                             # Stores: views, interactions, conversions, rankings, health scores, benchmarks
│   │   │       │
│   │   │       # ========================= SIMILARITY & DEDUPLICATION REPOSITORIES (CONSOLIDATED) =========================
│   │   │       ├── similarity_repository.go  # CONSOLIDATED: similarity + duplicate_check
│   │   │       │                             # Stores: fingerprints, clusters, duplicate markers, differentiation scores
│   │   │       │
│   │   │       # ========================= PORTFOLIO REPOSITORIES (CONSOLIDATED) =========================
│   │   │       ├── portfolio_repository.go   # CONSOLIDATED: portfolio_link + portfolio_selector
│   │   │       │                             # Stores: linked items, selection logic, display order
│   │   │       │
│   │   │       # ========================= ENGAGEMENT REPOSITORIES (CONSOLIDATED) =========================
│   │   │       ├── engagement_repository.go  # CONSOLIDATED: engagement + follow_up
│   │   │       │                             # Stores: interest signals, follow-up schedules, bumps, nudges
│   │   │       │
│   │   │       # ========================= MODERATION & COMPLIANCE REPOSITORIES =========================
│   │   │       ├── spam_detection_repository.go # Spam detection repository
│   │   │       ├── flag_repository.go        # Flagging repository
│   │   │       ├── compliance_repository.go  # Compliance checks repository (delegates to compliance-be)
│   │   │       │
│   │   │       # ========================= TEMPLATES & RATE CARDS REPOSITORIES =========================
│   │   │       ├── template_repository.go    # Proposal templates repository
│   │   │       └── rate_card_repository.go   # Rate card repository
│   │   │
│   │   ├── cache/
│   │   │   │── keys.go                  # 🗝️ Canonical cache keys & TTLs: avoids stringly-typed keys; defines per-domain TTL (proposal, performance, engagement)
│   │   │   │
│   │   │   └── redis/
│   │   │       ├── singleflight.go                 🆕  # stampede protection for hot keys
│   │   │       ├── invalidation_rules.go           🆕  # map events → keys to drop (proposal, performance, engagement)
│   │   │       ├── connection.go             # Redis connection setup
│   │   │       ├── proposal_cache.go         # Proposal caching (Get, Set, Invalidate, TTL)
│   │   │       ├── bid_cache.go              # Bid caching
│   │   │       ├── performance_cache.go      # Performance metrics caching
│   │   │       ├── similarity_cache.go       # Similarity/fingerprint caching
│   │   │       └── engagement_cache.go       # Engagement tracking caching
│   │   │
│   │   ├── messaging/
│   │   │   ├── topics.go                           ♻️  # thin re-export from contracts/events with context keys
│   │   │   │
│   │   │   ├── outbox/                             🆕  # exactly-once-ish publisher
│   │   │   │   ├── dispatcher.go                   🆕  # reads outbox → Kafka (retries/DLQ)
│   │   │   │   └── metrics.go                      🆕  # publish lag, failures, retries
│   │   │   └── kafka/
│   │   │       ├── consumer.go               # 📝 Kafka consumer (uses platform-shared/inbox)
│   │   │       ├── producer.go               # 📝 Kafka producer (uses platform-shared/outbox)
│   │   │       ├── topics.go                 # 📝 Topic constants (imported from contracts/events)
│   │   │       │
│   │   │       # ========================= TOPIC DEFINITIONS =========================
│   │   │       ├── topics_proposal.go        # Proposal lifecycle topics
│   │   │       ├── topics_bidding.go         # Bidding topics
│   │   │       ├── topics_collaboration.go   # Collaboration topics
│   │   │       ├── topics_performance.go     # Performance/analytics topics
│   │   │       │
│   │   │       # ========================= PRODUCERS =========================
│   │   │       ├── producers_proposal.go     # Proposal event producers
│   │   │       ├── producers_bidding.go      # Bidding event producers
│   │   │       ├── producers_collaboration.go # Collaboration event producers
│   │   │       ├── producers_performance.go  # Performance event producers
│   │   │       │
│   │   │       └── scram.go                  # SCRAM authentication for Kafka
│   │   │
│   │   ├── scheduler/
│   │   │   ├── cron.go                  # ⏱️ Cron registry: defines schedules, idempotency guards, and safe shutdown hooks (uses tasks/*)
│   │   │   └── tasks/
│   │   │       ├── expiration_task.go   # 📅 Sweeps proposals nearing/after expiry; emits notifications; marks expired (idempotent)
│   │   │       ├── followup_task.go     # 🔔 Sends scheduled follow-ups/bumps from engagement; respects frequency/quiet hours; deduped
│   │   │       ├── performance_recalc_task.go    # 📊 Recomputes performance/health/rankings periodically; writes to read-model; cheap if no changes
│   │   │       └── task_guard.go                   🆕  # idempotency tokens + last-run watermark
│   │   │
│   │   ├── projections/
│   │   │   └── performance/
│   │   │       ├── projector.go         # 🧮 Event projector: consumes ProposalViewed/Engagement/Conversion etc. → builds denormalized read model (CQRS)
│   │   │       └── readmodel_repository.go
│   │   │                                   # 📖 Read-model repository: fast query store for analytics/health/ranking (Postgres table or Redis hash)
│   │   │
│   │   ├── security/
│   │   │   ├── kms.go                   # 🔐 KMS wrapper: envelope-encrypt sensitive fields at rest; integrates with cloud KMS or local cipher
│   │   │   └── pii_redactor.go          # 🧽 PII redactor: masks emails/phones/links in logs & traces; safe logging helpers
│   │   │
│   │   │
│   │   ├── http/
│   │   │   ├── idempotency_adapter.go        # 🆕 Bind platform-shared idempotency to Gin
│   │   │   └── etag.go                             🆕  # middleware-level ETag (pairs with pkg/utils/etag)
│   │   │
│   │   ├── messaging_adapters/
│   │   │   ├── inbox/
│   │   │   │   └── dapr_subscriptions.yaml   # 🆕 Service-scoped Dapr subscriptions
│   │   │   └── outbox/
│   │   │       └── publisher.go              # 🆕 Thin wrapper over platform-shared outbox
│   │   │
│   │   ├── storage/
│   │   │   ├── client.go                     # Storage service client (upload attachments, cover letters)
│   │   │   └── local.go                      # Local file storage fallback
│   │   │
│   │   └── external_services/                # 🆕 External service clients (for externalized domains)
│   │       ├── intelligence_client.go        # 🆕 ML/Intelligence service client (strategy, optimization, predictions, recommendations, keywords, personalization, A/B testing)
│   │       ├── risk_client.go                # 🆕 Risk/Assurance service client (risk assessment, insurance, escrow, guarantee, dispute prediction)
│   │       ├── contracts_client.go           # 🆕 Contracts service client (terms, NDA, IP rights, payment terms, pricing, redlining, approval workflows)
│   │       ├── procurement_client.go         # 🆕 Procurement service client (RFP, quote, procurement, evaluation rubric, security questionnaire)
│   │       ├── marketplace_client.go         # 🆕 Marketplace service client (boost, visibility budget, premium features, fees, subscriptions, marketplace listings)
│   │       ├── clients_crm_client.go         # 🆕 Clients/CRM service client (client insights, client preferences)
│   │       ├── calendar_client.go            # 🆕 Calendar service client (deadline tracking, sync)
│   │       ├── time_tracking_client.go       # 🆕 Time tracking service client (time estimates, tracking)
│   │       ├── localization_client.go        # 🆕 Localization service client (translations)
│   │       ├── integrations_client.go        # 🆕 Integrations service client (external integrations)
│   │       ├── sharing_client.go             # 🆕 Secure sharing service client (secure links, access logs)
│   │       └── compliance_client.go          # 🆕 Compliance service client (compliance checks, verification)
│   │
│   ├── domain/                               # 🏛️ Domain Layer - Business logic & entities (Load Third)
│   │   │
│   │   # ========================= CORE PROPOSAL DOMAIN =========================
│   │   ├── proposal/
│   │   │   ├── entity.go                     # Proposal aggregate root (ID, JobID, FreelancerID, Title, Description, Status, Amount, Currency, DeliveryTime)
│   │   │   ├── value_objects.go              # Money, Duration, ProposalStatus value objects
│   │   │   ├── enums.go                      # ProposalStatus (Draft, Submitted, Shortlisted, Accepted, Rejected, Withdrawn, Archived)
│   │   │   ├── errors.go                     # Domain errors (ProposalNotFound, InvalidAmount, ProposalAlreadySubmitted)
│   │   │   ├── repository.go                 # Proposal repository interface
│   │   │   ├── list_filter.go                # ListFilter for queries (status, date range, amount range, job type)
│   │   │   └── events.go                     # ProposalCreated, ProposalSubmitted, ProposalAccepted, ProposalRejected, ProposalWithdrawn
│   │   │
│   │   ├── cover_letter/
│   │   │   ├── entity.go                     # CoverLetter (ProposalID, Content, WordCount, Tone)
│   │   │   ├── errors.go                     # CoverLetterTooShort, CoverLetterTooLong
│   │   │   ├── repository.go                 # CoverLetter repository interface
│   │   │   └── events.go                     # CoverLetterCreated, CoverLetterUpdated
│   │   │
│   │   ├── attachment/
│   │   │   ├── entity.go                     # ProposalAttachment (ProposalID, FileURL, FileName, FileSize, FileType)
│   │   │   ├── errors.go                     # AttachmentTooLarge, InvalidFileType
│   │   │   ├── repository.go                 # Attachment repository interface
│   │   │   └── events.go                     # AttachmentAdded, AttachmentRemoved
│   │   │
│   │   ├── question_answer/
│   │   │   ├── entity.go                     # ProposalQuestionAnswer (ProposalID, QuestionID, Answer)
│   │   │   ├── errors.go                     # QuestionNotAnswered, InvalidAnswer
│   │   │   ├── repository.go                 # Q&A repository interface
│   │   │   └── events.go                     # QuestionAnswered, AnswerUpdated
│   │   │
│   │   ├── milestone/
│   │   │   ├── entity.go                     # Milestone (ProposalID, Description, Amount, DueDate, Status)
│   │   │   ├── milestone_status.go           # MilestoneStatus (Pending, InProgress, Completed, Paid)
│   │   │   ├── errors.go                     # MilestoneNotFound, InvalidMilestoneAmount
│   │   │   ├── repository.go                 # Milestone repository interface
│   │   │   └── events.go                     # MilestoneCreated, MilestoneCompleted, MilestonePaid
│   │   │
│   │   # ========================= LIFECYCLE DOMAIN =========================
│   │   ├── expiration/
│   │   │   ├── entity.go                     # ProposalExpiration (ProposalID, ExpiresAt, ExtendedUntil, NotificationSent)
│   │   │   ├── errors.go                     # ProposalExpired, CannotExtend
│   │   │   ├── repository.go                 # Expiration repository interface
│   │   │   └── events.go                     # ProposalExpiring, ProposalExpired, ExpirationExtended
│   │   │
│   │   ├── withdrawal/
│   │   │   ├── entity.go                     # ProposalWithdrawal (ProposalID, Reason, WithdrawnAt, WithdrawnBy)
│   │   │   ├── withdrawal_reason.go          # WithdrawalReason enum (NotInterested, BetterOffer, ClientUnresponsive, ChangedMind)
│   │   │   ├── errors.go                     # CannotWithdraw, AlreadyWithdrawn
│   │   │   ├── repository.go                 # Withdrawal repository interface
│   │   │   └── events.go                     # ProposalWithdrawn, WithdrawalReasonUpdated
│   │   │
│   │   ├── archive/
│   │   │   ├── entity.go                     # ProposalArchive (ProposalID, ArchivedAt, RestorableUntil)
│   │   │   ├── errors.go                     # CannotArchive, RestorePeriodExpired
│   │   │   ├── repository.go                 # Archive repository interface
│   │   │   └── events.go                     # ProposalArchived, ProposalRestored
│   │   │
│   │   ├── pipeline/
│   │   │   ├── entity.go                     # ProposalPipeline (ProposalID, Stage, MovedAt, MovedBy)
│   │   │   ├── pipeline_stage.go             # PipelineStage enum (Drafting, Submitted, UnderReview, Shortlisted, Interviewing, Negotiating, Won, Lost)
│   │   │   ├── stage_transitions.go          # Valid stage transition rules
│   │   │   ├── errors.go                     # InvalidStageTransition, StageNotFound
│   │   │   ├── repository.go                 # Pipeline repository interface
│   │   │   └── events.go                     # PipelineStageMoved, StageChanged
│   │   │
│   │   ├── recycling/
│   │   │   ├── entity.go                     # ProposalRecycling (OriginalProposalID, NewProposalID, RecycledAt, Modifications)
│   │   │   ├── errors.go                     # CannotRecycle, TooManyVersions
│   │   │   ├── repository.go                 # Recycling repository interface
│   │   │   └── events.go                     # ProposalRecycled, VersionCreated
│   │   │
│   │   # ========================= BIDDING DOMAIN =========================
│   │   ├── bid/
│   │   │   ├── entity.go                     # Bid (ProposalID, BidType, Amount, Currency, ValidUntil, Rank)
│   │   │   ├── bid_type.go                   # BidType enum (Hourly, Fixed, Milestone, Retainer)
│   │   │   ├── bid_status.go                 # BidStatus enum (Active, Outbid, Withdrawn, Accepted)
│   │   │   ├── errors.go                     # BidTooLow, BidTooHigh, Outbid, InvalidBidType
│   │   │   ├── repository.go                 # Bid repository interface
│   │   │   └── events.go                     # BidPlaced, BidUpdated, Outbid, BidAccepted, BidRejected
│   │   │
│   │   ├── bid_strategy/
│   │   │   ├── entity.go                     # BidStrategy (FreelancerID, StrategyType, Rules, DefaultMarkup, MinAcceptableRate)
│   │   │   ├── strategy_type.go              # StrategyType enum (Competitive, Premium, ValueBased, Flexible)
│   │   │   ├── pricing_rule.go               # Pricing rules (job size, urgency, competition)
│   │   │   ├── errors.go                     # StrategyNotFound, InvalidRule
│   │   │   ├── repository.go                 # BidStrategy repository interface
│   │   │   └── events.go                     # StrategyCreated, StrategyUpdated, StrategyApplied
│   │   │
│   │   ├── auction/
│   │   │   ├── entity.go                     # Auction (JobID, StartTime, EndTime, CurrentHighBid, BidCount, Status)
│   │   │   ├── auction_status.go             # AuctionStatus enum (Scheduled, Active, Ended, Cancelled)
│   │   │   ├── errors.go                     # AuctionNotActive, AuctionEnded, BidBelowMinimum
│   │   │   ├── repository.go                 # Auction repository interface
│   │   │   └── events.go                     # AuctionStarted, NewHighBid, AuctionEnded, AuctionCancelled
│   │   │
│   │   ├── bid_anomaly/
│   │   │   ├── entity.go                     # BidAnomaly (BidID, AnomalyType, Severity, DetectedAt, ReviewedAt)
│   │   │   ├── anomaly_type.go               # AnomalyType enum (UnusuallyLow, UnusuallyHigh, RapidBidding, SuspiciousPattern)
│   │   │   ├── errors.go                     # AnomalyNotFound, AlreadyReviewed
│   │   │   ├── repository.go                 # BidAnomaly repository interface
│   │   │   └── events.go                     # AnomalyDetected, AnomalyReviewed, AnomalyConfirmed
│   │   │
│   │   # ========================= CONNECTS & CREDITS DOMAIN =========================
│   │   ├── connect/
│   │   │   ├── entity.go                     # ConnectUsage (ProposalID, FreelancerID, ConnectsUsed, UsedAt, RefundStatus)
│   │   │   ├── connect_tier.go               # ConnectTier (JobPopularity → connects cost: 1, 2, 4, 6)
│   │   │   ├── errors.go                     # InsufficientConnects, ConnectAlreadyUsed
│   │   │   ├── repository.go                 # Connect repository interface
│   │   │   └── events.go                     # ConnectUsed, ConnectRefunded, ConnectExpired
│   │   │
│   │   ├── connect_refund/
│   │   │   ├── entity.go                     # ConnectRefund (ProposalID, RequestedAt, ApprovedAt, Reason, Status)
│   │   │   ├── refund_reason.go              # RefundReason enum (JobCancelled, ClientUnresponsive, TechnicalIssue)
│   │   │   ├── errors.go                     # RefundNotEligible, RefundAlreadyProcessed
│   │   │   ├── repository.go                 # ConnectRefund repository interface
│   │   │   └── events.go                     # RefundRequested, RefundApproved, RefundDenied
│   │   │
│   │   # ========================= WORKFLOW & COLLABORATION DOMAIN =========================
│   │   ├── negotiation/
│   │   │   ├── entity.go                     # Negotiation (ProposalID, Thread, CurrentOffer, CounterOffers[], Status)
│   │   │   ├── counter_offer.go              # CounterOffer (Amount, Terms, OfferedBy, OfferedAt)
│   │   │   ├── negotiation_status.go         # NegotiationStatus enum (Active, Accepted, Declined, Expired)
│   │   │   ├── errors.go                     # NegotiationNotActive, InvalidCounterOffer
│   │   │   ├── repository.go                 # Negotiation repository interface
│   │   │   └── events.go                     # NegotiationStarted, CounterOfferProposed, NegotiationAccepted, NegotiationDeclined
│   │   │
│   │   ├── invite/
│   │   │   ├── entity.go                     # ProposalInvite (JobID, FreelancerID, InvitedBy, InvitedAt, ExpiresAt, Status)
│   │   │   ├── invite_status.go              # InviteStatus enum (Pending, Accepted, Declined, Expired)
│   │   │   ├── decline_reason.go             # DeclineReason enum (NotInterested, Busy, RateTooLow, SkillMismatch)
│   │   │   ├── errors.go                     # InviteExpired, InviteAlreadyActioned
│   │   │   ├── repository.go                 # Invite repository interface
│   │   │   └── events.go                     # InviteSent, InviteAccepted, InviteDeclined, InviteExpired
│   │   │
│   │   ├── revision/
│   │   │   ├── entity.go                     # ProposalRevision (ProposalID, Version, Changes, CreatedAt, ApprovedAt)
│   │   │   ├── revision_status.go            # RevisionStatus enum (Pending, Approved, Rejected)
│   │   │   ├── errors.go                     # RevisionNotFound, CannotRevert
│   │   │   ├── repository.go                 # Revision repository interface
│   │   │   └── events.go                     # RevisionCreated, RevisionApproved, RevisionRejected
│   │   │
│   │   ├── collaboration/                    # 🔄 CONSOLIDATED: team + proposal_collaboration_network
│   │   │   ├── entity.go                     # Collaboration (ProposalID, TeamMembers[], Roles[], RevenueShares[])
│   │   │   ├── team_member.go                # TeamMember (UserID, Role, RevenueSharePercentage, JoinedAt)
│   │   │   ├── role.go                       # Role enum (Lead, Contributor, Consultant, Subcontractor)
│   │   │   ├── revenue_split.go              # RevenueSplit (MemberID, Percentage, Amount, Status)
│   │   │   ├── errors.go                     # TeamFull, InvalidRevenueSplit, RoleTaken
│   │   │   ├── repository.go                 # Collaboration repository interface
│   │   │   └── events.go                     # TeamFormed, MemberAdded, MemberRemoved, RevenueSplitUpdated
│   │   │
│   │   # ========================= CLIENT INTERACTION DOMAIN =========================
│   │   ├── interview/
│   │   │   ├── entity.go                     # InterviewRequest (ProposalID, RequestedBy, ScheduledAt, Status, Notes)
│   │   │   ├── interview_status.go           # InterviewStatus enum (Requested, Scheduled, Completed, Cancelled)
│   │   │   ├── errors.go                     # InterviewNotFound, CannotSchedule
│   │   │   ├── repository.go                 # Interview repository interface
│   │   │   └── events.go                     # InterviewRequested, InterviewScheduled, InterviewCompleted, InterviewCancelled
│   │   │
│   │   ├── feedback/
│   │   │   ├── entity.go                     # ProposalFeedback (ProposalID, GivenBy, Rating, Comments, CreatedAt)
│   │   │   ├── errors.go                     # FeedbackAlreadyGiven, InvalidRating
│   │   │   ├── repository.go                 # Feedback repository interface
│   │   │   └── events.go                     # FeedbackGiven, FeedbackUpdated
│   │   │
│   │   ├── shortlist/
│   │   │   ├── entity.go                     # Shortlist (JobID, ProposalID, AddedBy, AddedAt, Rank, Status)
│   │   │   ├── shortlist_status.go           # ShortlistStatus enum (Active, Interviewing, Selected, Rejected)
│   │   │   ├── errors.go                     # AlreadyShortlisted, ShortlistFull
│   │   │   ├── repository.go                 # Shortlist repository interface
│   │   │   └── events.go                     # ProposalShortlisted, ShortlistStatusUpdated, RemovedFromShortlist
│   │   │
│   │   ├── conversation/
│   │   │   ├── entity.go                     # ProposalConversation (ProposalID, Messages[], Participants[], SentimentScore)
│   │   │   ├── message.go                    # Message (SenderID, Content, SentAt, ReadAt)
│   │   │   ├── sentiment.go                  # Sentiment analysis (Positive, Neutral, Negative, Score)
│   │   │   ├── errors.go                     # ConversationNotFound, MessageTooLong
│   │   │   ├── repository.go                 # Conversation repository interface
│   │   │   └── events.go                     # MessageSent, MessageRead, SentimentAnalyzed
│   │   │
│   │   # ========================= PERFORMANCE & ANALYTICS DOMAIN (CONSOLIDATED) =========================
│   │   ├── performance/                      # 🔄 CONSOLIDATED: analytics + engagement + response_tracker + conversion_tracking + proposal_metrics + ranking + proposal_health
│   │   │   ├── entity.go                     # ProposalPerformance (ProposalID, Metrics, Analytics, HealthScore, Rankings)
│   │   │   │
│   │   │   # --- Analytics (was proposal_analytics) ---
│   │   │   ├── analytics.go                  # ViewCount, ViewSources, ClientViews, Impressions, CTR
│   │   │   │
│   │   │   # --- Engagement (was proposal_engagement) ---
│   │   │   ├── engagement.go                 # InterestSignals, EngagementScore, TimeSpent, InteractionCount
│   │   │   │
│   │   │   # --- Response Tracking (was response_tracker) ---
│   │   │   ├── response_tracking.go          # ResponseTime, ResponseRate, ViewedAt, RespondedAt
│   │   │   │
│   │   │   # --- Conversion (was conversion_tracking) ---
│   │   │   ├── conversion.go                 # ConversionFunnel, ConversionRate, AttributionData, Touchpoints
│   │   │   │
│   │   │   # --- Metrics & Benchmarking (was proposal_metrics) ---
│   │   │   ├── metrics.go                    # PerformanceMetrics, Benchmarks, Percentile, ComparisonData
│   │   │   │
│   │   │   # --- Ranking (was ranking) ---
│   │   │   ├── ranking.go                    # SearchRank, VisibilityScore, RankingFactors, Position
│   │   │   │
│   │   │   # --- Health Score (was proposal_health) ---
│   │   │   ├── health.go                     # HealthScore, IssuesDetected, RecommendedActions, Warnings
│   │   │   │
│   │   │   ├── errors.go                     # PerformanceDataNotFound, InvalidMetric
│   │   │   ├── repository.go                 # Performance repository interface (stores all metrics)
│   │   │   └── events.go                     # ProposalViewed, EngagementRecorded, ConversionTracked, RankingUpdated, HealthScoreCalculated
│   │   │
│   │   # ========================= SIMILARITY & DEDUPLICATION DOMAIN (CONSOLIDATED) =========================
│   │   ├── similarity/                       # 🔄 CONSOLIDATED: proposal_similarity + duplicate_check
│   │   │   ├── entity.go                     # ProposalSimilarity (ProposalID, Fingerprint, Clusters, Duplicates, DifferentiationScore)
│   │   │   │
│   │   │   # --- Fingerprinting ---
│   │   │   ├── fingerprint.go                # TextFingerprint, StructuralFingerprint, SemanticHash
│   │   │   │
│   │   │   # --- Clustering ---
│   │   │   ├── clustering.go                 # ClusterID, ClusterMembers, ClusterCentroid, SimilarityThreshold
│   │   │   │
│   │   │   # --- Duplicate Detection ---
│   │   │   ├── duplicate_detection.go        # DuplicateProposalID, DuplicateScore, MatchType (Exact, Near, Partial)
│   │   │   │
│   │   │   # --- Differentiation ---
│   │   │   ├── differentiation.go            # DifferentiationScore, UniqueElements, CompetitiveAdvantage
│   │   │   │
│   │   │   ├── errors.go                     # FingerprintNotFound, DuplicateDetected, ClusteringFailed
│   │   │   ├── repository.go                 # Similarity repository interface (stores fingerprints, clusters, duplicates)
│   │   │   └── events.go                     # FingerprintCreated, DuplicateDetected, ProposalClustered, DifferentiationScored
│   │   │
│   │   # ========================= PORTFOLIO DOMAIN (CONSOLIDATED) =========================
│   │   ├── portfolio/                        # 🔄 CONSOLIDATED: portfolio_link + proposal_portfolio_selector
│   │   │   ├── entity.go                     # ProposalPortfolio (ProposalID, LinkedItems[], SelectionCriteria, DisplayOrder)
│   │   │   │
│   │   │   # --- Portfolio Linking ---
│   │   │   ├── portfolio_link.go             # PortfolioItemID, ProposalID, Relevance, LinkedAt
│   │   │   │
│   │   │   # --- Auto-Selection Logic ---
│   │   │   ├── portfolio_selector.go         # SelectionAlgorithm, RelevanceScore, AutoSelectCriteria, MaxItems
│   │   │   │
│   │   │   ├── errors.go                     # PortfolioItemNotFound, MaxItemsExceeded, InvalidSelection
│   │   │   ├── repository.go                 # Portfolio repository interface (stores links and selection logic)
│   │   │   └── events.go                     # PortfolioItemLinked, PortfolioItemUnlinked, PortfolioAutoSelected
│   │   │
│   │   # ========================= ENGAGEMENT DOMAIN (CONSOLIDATED) =========================
│   │   ├── engagement/                       # 🔄 CONSOLIDATED: proposal_engagement + proposal_follow_up
│   │   │   ├── entity.go                     # ProposalEngagement (ProposalID, InterestSignals[], FollowUps[], Bumps[], Nudges[])
│   │   │   │
│   │   │   # --- Interest Tracking ---
│   │   │   ├── interest_tracking.go          # InterestLevel, SignalType (View, Save, Message, Share), DetectedAt
│   │   │   │
│   │   │   # --- Follow-up Management ---
│   │   │   ├── follow_up.go                  # FollowUpSchedule, ReminderType, TriggerConditions, SentAt
│   │   │   │
│   │   │   # --- Bumps & Nudges ---
│   │   │   ├── bump.go                       # BumpType (Gentle, Standard, Urgent), BumpedAt, ResponseReceived
│   │   │   │
│   │   │   ├── errors.go                     # TooManyFollowUps, BumpNotAllowed, EngagementNotTracked
│   │   │   ├── repository.go                 # Engagement repository interface (stores interest signals and follow-ups)
│   │   │   └── events.go                     # InterestDetected, FollowUpScheduled, FollowUpSent, ProposalBumped
│   │   │
│   │   # ========================= MODERATION & COMPLIANCE DOMAIN =========================
│   │   ├── spam_detection/
│   │   │   ├── entity.go                     # SpamDetection (ProposalID, SpamScore, Reasons[], DetectedAt, ReviewedAt)
│   │   │   ├── spam_indicators.go            # SpamIndicators (LowQualityText, SuspiciousLinks, MassSubmission, BotBehavior)
│   │   │   ├── errors.go                     # SpamDetected, AlreadyFlagged
│   │   │   ├── repository.go                 # SpamDetection repository interface
│   │   │   └── events.go                     # SpamDetected, SpamFlagged, SpamCleared
│   │   │
│   │   ├── flag/
│   │   │   ├── entity.go                     # ProposalFlag (ProposalID, FlaggedBy, Reason, FlaggedAt, Status)
│   │   │   ├── flag_reason.go                # FlagReason enum (Spam, Offensive, Plagiarism, Scam, Other)
│   │   │   ├── errors.go                     # AlreadyFlagged, CannotFlagOwn
│   │   │   ├── repository.go                 # Flag repository interface
│   │   │   └── events.go                     # ProposalFlagged, FlagReviewed, FlagResolved
│   │   │
│   │   ├── compliance/
│   │   │   ├── entity.go                     # ProposalCompliance (ProposalID, Checks[], Status, CompletedAt)
│   │   │   ├── compliance_check.go           # CheckType (ContentPolicy, TOS, LegalReview, DataPrivacy), Result, Evidence
│   │   │   ├── errors.go                     # ComplianceCheckFailed, ViolationDetected
│   │   │   ├── repository.go                 # Compliance repository interface (delegates heavy checks to compliance-be)
│   │   │   └── events.go                     # ComplianceCheckRequested, ComplianceCheckCompleted, ViolationDetected
│   │   │
│   │   # ========================= TEMPLATES & RATE CARDS DOMAIN =========================
│   │   ├── template/
│   │   │   ├── entity.go                     # ProposalTemplate (UserID, Name, Content, Category, UsageCount)
│   │   │   ├── template_category.go          # TemplateCategory enum (WebDevelopment, Design, Writing, Marketing)
│   │   │   ├── errors.go                     # TemplateNotFound, TemplateNameTaken
│   │   │   ├── repository.go                 # Template repository interface
│   │   │   └── events.go                     # TemplateCreated, TemplateUpdated, TemplateUsed
│   │   │
│   │   └── rate_card/
│   │       ├── entity.go                     # RateCard (FreelancerID, Packages[], Tiers[], DefaultRates)
│   │       ├── package.go                    # Package (Name, Description, Price, Deliverables, DeliveryTime)
│   │       ├── package_tier.go               # PackageTier enum (Starter, Standard, Premium)
│   │       ├── errors.go                     # RateCardNotFound, InvalidPackage
│   │       ├── repository.go                 # RateCard repository interface
│   │       └── events.go                     # RateCardCreated, RateCardUpdated, PackageAdded
│   │
│   ├── application/                          # 📋 Application Layer - Use cases & orchestration (Load Fourth)
│   │   │
│   │   │── entitlements/
│   │   │       └── checker.go               # 🔐 Feature/plan checks: central helper used by handlers/services to AllowIf(...) against marketplace/subscriptions (cached)
│   │   │
│   │   ├── authz/                                  🆕  # service-level permission checks
│   │   │   ├── policies.go                         🆕  # map roles → actions (owner, collaborator, admin)
│   │   │   └── guards.go  
│   │   │
│   │   # ========================= EVENT HANDLERS (Inbound Events) =========================
│   │   ├── eventhandler/
│   │   │   ├── job_handler.go                # Consumes: job.posted → enable proposals; job.closed → disable proposals
│   │   │   ├── contract_handler.go           # Consumes: contract.created → mark proposal as accepted
│   │   │   ├── user_handler.go               # Consumes: user.suspended → hide proposals; user.banned → archive proposals
│   │   │   ├── review_handler.go             # Consumes: review.submitted → update proposal performance
│   │   │   ├── message_handler.go            # Consumes: message.sent → update engagement tracking
│   │   │   ├── payment_handler.go            # Consumes: payment.processed → update milestone status
│   │   │   │
│   │   │   # ========================= EXTERNAL SERVICE EVENT HANDLERS =========================
│   │   │   ├── intelligence_handler.go       # 🆕 Consumes: intelligence.* events from ML service
│   │   │   ├── risk_handler.go               # 🆕 Consumes: risk.* events from risk service
│   │   │   ├── contracts_handler.go          # 🆕 Consumes: contracts.* events from contracts service
│   │   │   ├── procurement_handler.go        # 🆕 Consumes: procurement.* events from procurement service
│   │   │   ├── marketplace_handler.go        # 🆕 Consumes: marketplace.* events (boost, visibility, subscriptions)
│   │   │   ├── compliance_handler.go         # 🆕 Consumes: compliance.check.completed from compliance service
│   │   │   └── _common/
│   │   │       └── idempotency.go                  🆕  # store last-seen event version per aggregate
│   │   │
│   │   │
│   │   # ========================= CORE PROPOSAL SERVICES =========================
│   │   ├── proposal/
│   │   │   ├── service.go                    # Proposal business logic (Create, Update, Submit, Withdraw, Archive)
│   │   │   ├── commands.go                   # CreateProposal, UpdateProposal, SubmitProposal, WithdrawProposal, ArchiveProposal
│   │   │   ├── queries.go                    # GetProposal, ListProposals, SearchProposals, GetProposalStats
│   │   │   ├── dto.go                        # ProposalDTO, CreateProposalDTO, UpdateProposalDTO, ProposalListDTO, ProposalStatsDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Input validation (amount, delivery time, status transitions)
│   │   │
│   │   ├── cover_letter/
│   │   │   ├── service.go                    # CoverLetter business logic (Create, Update, Analyze)
│   │   │   ├── commands.go                   # CreateCoverLetter, UpdateCoverLetter
│   │   │   ├── queries.go                    # GetCoverLetter, AnalyzeCoverLetter
│   │   │   ├── dto.go                        # CoverLetterDTO, CreateCoverLetterDTO, AnalysisDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Content validation (length, tone, quality)
│   │   │
│   │   ├── attachment/
│   │   │   ├── service.go                    # Attachment business logic (Upload, Delete, List)
│   │   │   ├── commands.go                   # UploadAttachment, DeleteAttachment
│   │   │   ├── queries.go                    # ListAttachments, GetAttachment
│   │   │   ├── dto.go                        # AttachmentDTO, UploadAttachmentDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # File validation (size, type, virus scan)
│   │   │
│   │   ├── question_answer/
│   │   │   ├── service.go                    # Q&A business logic (Answer, Update)
│   │   │   ├── commands.go                   # AnswerQuestion, UpdateAnswer
│   │   │   ├── queries.go                    # GetAnswers, ListQuestions
│   │   │   ├── dto.go                        # QuestionAnswerDTO, AnswerQuestionDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Answer validation (completeness, relevance)
│   │   │
│   │   ├── milestone/
│   │   │   ├── service.go                    # Milestone business logic (Create, Update, Complete)
│   │   │   ├── commands.go                   # CreateMilestone, UpdateMilestone, CompleteMilestone
│   │   │   ├── queries.go                    # GetMilestones, GetMilestoneStatus
│   │   │   ├── dto.go                        # MilestoneDTO, CreateMilestoneDTO, MilestoneStatusDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Milestone validation (amount splits, dates)
│   │   │
│   │   # ========================= LIFECYCLE SERVICES =========================
│   │   ├── expiration/
│   │   │   ├── service.go                    # Expiration business logic (Set, Extend, Notify, Check)
│   │   │   ├── commands.go                   # SetExpiration, ExtendExpiration, NotifyExpiration
│   │   │   ├── queries.go                    # GetExpiration, GetExpiringProposals
│   │   │   ├── dto.go                        # ExpirationDTO, ExtendExpirationDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Expiration date validation
│   │   │
│   │   ├── withdrawal/
│   │   │   ├── service.go                    # Withdrawal business logic (Withdraw, UpdateReason)
│   │   │   ├── commands.go                   # WithdrawProposal, UpdateWithdrawalReason
│   │   │   ├── queries.go                    # GetWithdrawal, GetWithdrawalHistory
│   │   │   ├── dto.go                        # WithdrawalDTO, WithdrawProposalDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Withdrawal validation (eligibility, reason)
│   │   │
│   │   ├── archive/
│   │   │   ├── service.go                    # Archive business logic (Archive, Restore, PurgeExpired)
│   │   │   ├── commands.go                   # ArchiveProposal, RestoreProposal, PurgeExpiredArchives
│   │   │   ├── queries.go                    # ListArchivedProposals, GetArchiveStatus
│   │   │   ├── dto.go                        # ArchiveDTO, ArchivedProposalListDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Archive validation (restore window)
│   │   │
│   │   ├── pipeline/
│   │   │   ├── service.go                    # Pipeline business logic (MoveStage, GetPipeline, AnalyzePipeline)
│   │   │   ├── commands.go                   # MoveToStage, UpdatePipeline
│   │   │   ├── queries.go                    # GetPipeline, GetPipelineAnalytics, ListProposalsByStage
│   │   │   ├── dto.go                        # PipelineDTO, PipelineStageDTO, PipelineAnalyticsDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Stage transition validation
│   │   │
│   │   ├── recycling/
│   │   │   ├── service.go                    # Recycling business logic (RecycleProposal, GetVersions)
│   │   │   ├── commands.go                   # RecycleProposal, CreateVersion
│   │   │   ├── queries.go                    # GetVersionHistory, GetRecyclableProposals
│   │   │   ├── dto.go                        # RecyclingDTO, VersionHistoryDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Recycling validation (version limits)
│   │   │
│   │   # ========================= BIDDING SERVICES =========================
│   │   ├── bid/
│   │   │   ├── service.go                    # Bid business logic (PlaceBid, UpdateBid, WithdrawBid, CalculateRank)
│   │   │   ├── commands.go                   # PlaceBid, UpdateBid, WithdrawBid, AcceptBid
│   │   │   ├── queries.go                    # GetBid, ListBids, GetBidRank, GetCompetitiveBids
│   │   │   ├── dto.go                        # BidDTO, PlaceBidDTO, BidRankDTO, CompetitiveBidsDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Bid validation (amount ranges, type consistency)
│   │   │
│   │   ├── bid_strategy/
│   │   │   ├── service.go                    # BidStrategy business logic (CreateStrategy, ApplyStrategy, OptimizeStrategy)
│   │   │   ├── commands.go                   # CreateBidStrategy, UpdateBidStrategy, ApplyStrategyToProposal
│   │   │   ├── queries.go                    # GetBidStrategy, GetRecommendedBid, AnalyzeStrategy
│   │   │   ├── dto.go                        # BidStrategyDTO, CreateBidStrategyDTO, RecommendedBidDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Strategy validation (rules, markup ranges)
│   │   │
│   │   ├── auction/
│   │   │   ├── service.go                    # Auction business logic (StartAuction, PlaceBid, EndAuction)
│   │   │   ├── commands.go                   # StartAuction, EndAuction, CancelAuction
│   │   │   ├── queries.go                    # GetAuction, GetAuctionStatus, GetHighestBid
│   │   │   ├── dto.go                        # AuctionDTO, AuctionStatusDTO, AuctionBidDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Auction validation (timing, minimum bid)
│   │   │
│   │   ├── bid_anomaly/
│   │   │   ├── service.go                    # BidAnomaly business logic (DetectAnomalies, ReviewAnomaly)
│   │   │   ├── commands.go                   # DetectAnomalies, ReviewAnomaly, ConfirmAnomaly
│   │   │   ├── queries.go                    # GetAnomalies, GetAnomalyReport
│   │   │   ├── dto.go                        # BidAnomalyDTO, AnomalyReportDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Anomaly detection thresholds
│   │   │
│   │   # ========================= CONNECTS & CREDITS SERVICES =========================
│   │   ├── connect/
│   │   │   ├── service.go                    # Connect business logic (UseConnects, GetBalance, CalculateCost)
│   │   │   ├── commands.go                   # UseConnects, RefundConnects
│   │   │   ├── queries.go                    # GetConnectBalance, GetConnectUsageHistory, CalculateConnectCost
│   │   │   ├── dto.go                        # ConnectDTO, ConnectBalanceDTO, ConnectUsageDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Connect validation (sufficient balance)
│   │   │
│   │   ├── connect_refund/
│   │   │   ├── service.go                    # ConnectRefund business logic (RequestRefund, ProcessRefund, DenyRefund)
│   │   │   ├── commands.go                   # RequestRefund, ApproveRefund, DenyRefund
│   │   │   ├── queries.go                    # GetRefundStatus, ListRefundRequests
│   │   │   ├── dto.go                        # ConnectRefundDTO, RefundRequestDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Refund eligibility validation
│   │   │
│   │   # ========================= WORKFLOW & COLLABORATION SERVICES =========================
│   │   ├── negotiation/
│   │   │   ├── service.go                    # Negotiation business logic (StartNegotiation, ProposeCounter, AcceptNegotiation)
│   │   │   ├── commands.go                   # StartNegotiation, ProposeCounterOffer, AcceptNegotiation, DeclineNegotiation
│   │   │   ├── queries.go                    # GetNegotiation, ListNegotiations, GetNegotiationHistory
│   │   │   ├── dto.go                        # NegotiationDTO, CounterOfferDTO, NegotiationHistoryDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Negotiation validation (offer amounts, terms)
│   │   │
│   │   ├── invite/
│   │   │   ├── service.go                    # Invite business logic (SendInvite, AcceptInvite, DeclineInvite, PrefillProposal)
│   │   │   ├── commands.go                   # SendInvite, AcceptInvite, DeclineInvite, ResendInvite
│   │   │   ├── queries.go                    # GetInvite, ListInvites, GetInviteStats
│   │   │   ├── dto.go                        # InviteDTO, SendInviteDTO, InviteStatsDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Invite validation (expiry, recipient)
│   │   │
│   │   ├── revision/
│   │   │   ├── service.go                    # Revision business logic (CreateRevision, ApproveRevision, RejectRevision)
│   │   │   ├── commands.go                   # CreateRevision, ApproveRevision, RejectRevision, RevertToRevision
│   │   │   ├── queries.go                    # GetRevisionHistory, GetLatestRevision, CompareRevisions
│   │   │   ├── dto.go                        # RevisionDTO, RevisionHistoryDTO, RevisionComparisonDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Revision validation (state transitions)
│   │   │
│   │   ├── collaboration/
│   │   │   ├── service.go                    # Collaboration business logic (FormTeam, AddMember, SetRevenueSplit)
│   │   │   ├── commands.go                   # FormTeam, AddTeamMember, RemoveTeamMember, UpdateRevenueSplit
│   │   │   ├── queries.go                    # GetTeam, GetTeamMembers, GetRevenueSplits
│   │   │   ├── dto.go                        # CollaborationDTO, TeamMemberDTO, RevenueSplitDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Team validation (member limits, revenue split totals)
│   │   │
│   │   # ========================= CLIENT INTERACTION SERVICES =========================
│   │   ├── interview/
│   │   │   ├── service.go                    # Interview business logic (RequestInterview, ScheduleInterview, CompleteInterview)
│   │   │   ├── commands.go                   # RequestInterview, ScheduleInterview, UpdateInterviewStatus, CancelInterview
│   │   │   ├── queries.go                    # GetInterview, ListInterviews, GetInterviewAvailability
│   │   │   ├── dto.go                        # InterviewDTO, ScheduleInterviewDTO, InterviewAvailabilityDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Interview validation (scheduling conflicts)
│   │   │
│   │   ├── feedback/
│   │   │   ├── service.go                    # Feedback business logic (GiveFeedback, RespondToFeedback)
│   │   │   ├── commands.go                   # GiveFeedback, UpdateFeedback, RespondToFeedback
│   │   │   ├── queries.go                    # GetFeedback, ListFeedback, GetFeedbackStats
│   │   │   ├── dto.go                        # FeedbackDTO, GiveFeedbackDTO, FeedbackStatsDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Feedback validation (rating bounds, duplicates)
│   │   │
│   │   ├── shortlist/
│   │   │   ├── service.go                    # Shortlist business logic (AddToShortlist, RemoveFromShortlist, UpdateStatus)
│   │   │   ├── commands.go                   # AddToShortlist, RemoveFromShortlist, UpdateShortlistStatus
│   │   │   ├── queries.go                    # GetShortlist, ListShortlistedProposals
│   │   │   ├── dto.go                        # ShortlistDTO, ShortlistEntryDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Shortlist validation (size limits, status transitions)
│   │   │
│   │   ├── conversation/
│   │   │   ├── service.go                    # Conversation business logic (SendMessage, AnalyzeSentiment)
│   │   │   ├── commands.go                   # SendMessage, MarkAsRead
│   │   │   ├── queries.go                    # GetConversation, GetMessages, GetSentimentAnalysis
│   │   │   ├── dto.go                        # ConversationDTO, MessageDTO, SentimentDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Message validation (length, content policy)
│   │   │
│   │   # ========================= PERFORMANCE & ANALYTICS SERVICES (CONSOLIDATED) =========================
│   │   ├── performance/
│   │   │   ├── service.go                    # Performance business logic (TrackView, RecordEngagement, CalculateHealth, UpdateRankings)
│   │   │   │
│   │   │   # --- Commands ---
│   │   │   ├── commands.go                   # TrackProposalView, RecordEngagementEvent, RecordConversion, RecalculateHealthScore, UpdateRankings
│   │   │   │
│   │   │   # --- Queries ---
│   │   │   ├── queries.go                    # GetAnalytics, GetEngagementMetrics, GetConversionFunnel, GetHealthScore, GetRankings, GetBenchmarks
│   │   │   │
│   │   │   # --- DTOs (all consolidated metrics) ---
│   │   │   ├── dto.go                        # PerformanceDTO, AnalyticsDTO, EngagementDTO, ConversionDTO, HealthDTO, RankingDTO, BenchmarkDTO
│   │   │   │
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Metrics validation
│   │   │
│   │   # ========================= SIMILARITY & DEDUPLICATION SERVICES (CONSOLIDATED) =========================
│   │   ├── similarity/
│   │   │   ├── service.go                    # Similarity business logic (CreateFingerprint, DetectDuplicates, ClusterProposals, ScoreDifferentiation)
│   │   │   ├── commands.go                   # CreateFingerprint, ClusterProposals, MarkAsUnique, CalculateDifferentiation
│   │   │   ├── queries.go                    # GetSimilarProposals, GetDuplicates, GetClusters, GetDifferentiationScore
│   │   │   ├── dto.go                        # SimilarityDTO, FingerprintDTO, DuplicateDTO, ClusterDTO, DifferentiationDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Similarity thresholds validation
│   │   │
│   │   # ========================= PORTFOLIO SERVICES (CONSOLIDATED) =========================
│   │   ├── portfolio/
│   │   │   ├── service.go                    # Portfolio business logic (LinkPortfolioItem, AutoSelectItems, ReorderItems)
│   │   │   ├── commands.go                   # LinkPortfolioItem, UnlinkPortfolioItem, AutoAttachPortfolio, ReorderPortfolio
│   │   │   ├── queries.go                    # GetLinkedPortfolio, GetAutoSelectionCriteria, GetRelevanceScores
│   │   │   ├── dto.go                        # PortfolioDTO, PortfolioLinkDTO, AutoSelectionDTO, RelevanceScoreDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Portfolio validation (item limits, relevance)
│   │   │
│   │   # ========================= ENGAGEMENT SERVICES (CONSOLIDATED) =========================
│   │   ├── engagement/
│   │   │   ├── service.go                    # Engagement business logic (TrackInterest, ScheduleFollowUp, SendBump)
│   │   │   ├── commands.go                   # RecordInterestSignal, ScheduleFollowUp, SendFollowUp, BumpProposal
│   │   │   ├── queries.go                    # GetInterestLevel, GetFollowUpSchedule, GetBumpHistory
│   │   │   ├── dto.go                        # EngagementDTO, InterestDTO, FollowUpDTO, BumpDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Engagement validation (follow-up limits, bump frequency)
│   │   │
│   │   # ========================= MODERATION & COMPLIANCE SERVICES =========================
│   │   ├── spam_detection/
│   │   │   ├── service.go                    # SpamDetection business logic (ScanProposal, FlagSpam, ClearSpam)
│   │   │   ├── commands.go                   # ScanProposalForSpam, FlagAsSpam, ClearSpamFlags
│   │   │   ├── queries.go                    # GetSpamStatus, GetSpamScore, GetSpamReasons
│   │   │   ├── dto.go                        # SpamDetectionDTO, SpamScoreDTO, SpamIndicatorsDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Spam score thresholds
│   │   │
│   │   ├── flag/
│   │   │   ├── service.go                    # Flag business logic (FlagProposal, ReviewFlag, ResolveFlag)
│   │   │   ├── commands.go                   # FlagProposal, UnflagProposal, ReviewFlag, ResolveFlag
│   │   │   ├── queries.go                    # GetFlags, GetFlagHistory
│   │   │   ├── dto.go                        # FlagDTO, FlagReviewDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Flag reason validation
│   │   │
│   │   ├── compliance/
│   │   │   ├── service.go                    # Compliance business logic (RunChecks, ResolveViolation) - delegates to compliance-be
│   │   │   ├── commands.go                   # RequestComplianceCheck, ResolveViolation
│   │   │   ├── queries.go                    # GetComplianceStatus, GetViolations
│   │   │   ├── dto.go                        # ComplianceDTO, ViolationDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Compliance validation
│   │   │
│   │   # ========================= TEMPLATES & RATE CARDS SERVICES =========================
│   │   ├── template/
│   │   │   ├── service.go                    # Template business logic (CreateTemplate, UseTemplate, UpdateTemplate)
│   │   │   ├── commands.go                   # CreateTemplate, UpdateTemplate, DeleteTemplate, UseTemplate
│   │   │   ├── queries.go                    # GetTemplate, ListTemplates, GetTemplateUsageStats
│   │   │   ├── dto.go                        # TemplateDTO, CreateTemplateDTO, TemplateUsageStatsDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Template validation (name, content)
│   │   │
│   │   ├── rate_card/
│   │   │   ├── service.go                    # RateCard business logic (CreateRateCard, AddPackage, ApplyToProposal)
│   │   │   ├── commands.go                   # CreateRateCard, UpdateRateCard, AddPackage, UpdatePackage
│   │   │   ├── queries.go                    # GetRateCard, ListPackages, GetDefaultRates
│   │   │   ├── dto.go                        # RateCardDTO, PackageDTO, DefaultRatesDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO mapping
│   │   │   └── validators.go                 # Rate card validation (pricing, packages)
│   │   │
│   │   # ========================= 🆕 INTELLIGENCE FAÇADE (delegates to intelligence-be) =========================
│   │   └── intelligence_facade/              # 🆕 Thin façade for ML/Intelligence service
│   │       ├── service.go                    # Intelligence façade (GetStrategy, GetOptimization, GetRecommendations, GetPredictions)
│   │       ├── queries.go                    # GetProposalStrategy, GetOptimizationSuggestions, GetRecommendations, GetSuccessPrediction, GetKeywordSuggestions
│   │       ├── dto.go                        # StrategyDTO, OptimizationDTO, RecommendationDTO, PredictionDTO, KeywordOptimizationDTO
│   │       └── validators.go                 # Request validation for intelligence service
│   │
│   └── interfaces/                           # 🌐 Interface Layer (API/HTTP - Load Fifth)
│       └── http/
│           │
│           ├── etag.go                       🆕  # middleware-level ETag (pairs with pkg/utils/etag)
│           │
│           ├── presenters/
│           │   └── errors.go                 # 🎯 Maps domain/service errors → HTTP status & problem+json shape; stable error codes for clients
│           │
│           ├── middleware/
│           │   ├── requestid.go              # 🆕 Wraps platform-shared request id
│           │   ├── logging.go                # 🆕 Wraps platform-shared logging
│           │   ├── recovery.go               # 🆕 Wraps platform-shared recovery
│           │   ├── tracing.go                # 🆕 Wraps platform-shared otel middleware
│           │   ├── auth.go                   # 🆕 Wraps pkg/auth (Keycloak)
│           │   ├── rbac.go                   # 🆕 Role/permission gate (aligns with application/authz)
│           │   └── cors.go                   # 🆕 CORS middleware
│           │
│           ├── v1/                           # 🆕 Versioned API surface (/v1/*)
│           │   │
│           │   ├── openapi/                  # 🆕 OpenAPI Contract (v1)
│           │   │   ├── openapi.yaml          # 📜 OpenAPI v1: request/response schemas; source of truth for clients/tests
│           │   │   └── generator.go          # 🧭 Serves /v1/swagger and /v1/openapi.json (dev-only toggle); optional request validation in debug
│           │   │
│           │   ├── handlers/
│           │   │   │
│           │   │   # ========================= CORE PROPOSAL HANDLERS =========================
│           │   │   ├── proposal_handler.go       # POST /v1/proposals, GET /v1/proposals/:id, PUT /v1/proposals/:id, DELETE /v1/proposals/:id
│           │   │   ├── cover_letter_handler.go   # POST /v1/proposals/:id/cover-letter, PUT /v1/proposals/:id/cover-letter
│           │   │   ├── attachment_handler.go     # POST /v1/proposals/:id/attachments, DELETE /v1/proposals/:id/attachments/:attachmentId
│           │   │   ├── question_answer_handler.go# POST /v1/proposals/:id/answers, PUT /v1/proposals/:id/answers/:answerId
│           │   │   ├── milestone_handler.go      # POST /v1/proposals/:id/milestones, PUT /v1/proposals/:id/milestones/:milestoneId
│           │   │   ├── health_handler.go         # 🆕 GET /v1/healthz/live, GET /v1/healthz/ready
│           │   │   │
│           │   │   # ========================= LIFECYCLE HANDLERS =========================
│           │   │   ├── expiration_handler.go     # POST /v1/proposals/:id/expiration, PUT /v1/proposals/:id/expiration/extend
│           │   │   ├── withdrawal_handler.go     # POST /v1/proposals/:id/withdraw, PUT /v1/proposals/:id/withdrawal/reason
│           │   │   ├── archive_handler.go        # POST /v1/proposals/:id/archive, POST /v1/proposals/:id/restore
│           │   │   ├── pipeline_handler.go       # GET /v1/proposals/pipeline, POST /v1/proposals/:id/pipeline/move, GET /v1/proposals/pipeline/analytics
│           │   │   ├── recycling_handler.go      # POST /v1/proposals/:id/recycle, GET /v1/proposals/:id/versions
│           │   │   │
│           │   │   # ========================= BIDDING HANDLERS =========================
│           │   │   ├── bid_handler.go            # POST /v1/proposals/:id/bid, PUT /v1/proposals/:id/bid, GET /v1/proposals/:id/bid
│           │   │   ├── bid_strategy_handler.go   # POST /v1/bid-strategies, PUT /v1/bid-strategies/:id, GET /v1/bid-strategies/:id/recommend
│           │   │   ├── auction_handler.go        # POST /v1/jobs/:jobId/auction/bid, GET /v1/jobs/:jobId/auction/status, POST /v1/jobs/:jobId/auction/end
│           │   │   ├── bid_anomaly_handler.go    # POST /v1/bids/:bidId/anomaly/detect, POST /v1/bids/:bidId/anomaly/review, GET /v1/bids/:bidId/anomalies
│           │   │   │
│           │   │   # ========================= CONNECTS & CREDITS HANDLERS =========================
│           │   │   ├── connect_handler.go        # GET /v1/connects/balance, POST /v1/proposals/:id/connects/use
│           │   │   ├── connect_refund_handler.go # POST /v1/proposals/:id/connects/refund, GET /v1/proposals/:id/connects/refund/status
│           │   │   │
│           │   │   # ========================= WORKFLOW & COLLABORATION HANDLERS =========================
│           │   │   ├── negotiation_handler.go    # POST /v1/proposals/:id/negotiations, POST /v1/proposals/:id/negotiations/:negotiationId/counter, POST /v1/proposals/:id/negotiations/:negotiationId/accept
│           │   │   ├── invite_handler.go         # POST /v1/invites, POST /v1/invites/:inviteId/accept, POST /v1/invites/:inviteId/decline
│           │   │   ├── revision_handler.go       # POST /v1/proposals/:id/revisions, POST /v1/proposals/:id/revisions/:revisionId/approve, GET /v1/proposals/:id/revisions
│           │   │   ├── collaboration_handler.go  # POST /v1/proposals/:id/team, POST /v1/proposals/:id/team/members, PUT /v1/proposals/:id/team/revenue-split
│           │   │   │
│           │   │   # ========================= CLIENT INTERACTION HANDLERS =========================
│           │   │   ├── interview_handler.go      # POST /v1/proposals/:id/interviews, PUT /v1/proposals/:id/interviews/:interviewId/schedule, POST /v1/proposals/:id/interviews/:interviewId/complete
│           │   │   ├── feedback_handler.go       # POST /v1/proposals/:id/feedback, PUT /v1/proposals/:id/feedback, GET /v1/proposals/:id/feedback
│           │   │   ├── shortlist_handler.go      # POST /v1/jobs/:jobId/shortlist, DELETE /v1/jobs/:jobId/shortlist/:proposalId, PUT /v1/jobs/:jobId/shortlist/:proposalId/status
│           │   │   ├── conversation_handler.go   # POST /v1/proposals/:id/conversation/messages, GET /v1/proposals/:id/conversation, GET /v1/proposals/:id/conversation/sentiment
│           │   │   │
│           │   │   # ========================= PERFORMANCE & ANALYTICS HANDLERS (CONSOLIDATED) =========================
│           │   │   ├── performance_handler.go    # 🔄 CONSOLIDATED: analytics + engagement + response_tracker + conversion + metrics + ranking + health
│           │   │   │                             # GET /v1/proposals/:id/analytics, GET /v1/proposals/:id/engagement, GET /v1/proposals/:id/conversions
│           │   │   │                             # GET /v1/proposals/:id/health, GET /v1/proposals/:id/ranking, GET /v1/proposals/:id/benchmarks
│           │   │   │                             # POST /v1/proposals/:id/track/view, POST /v1/proposals/:id/track/engagement
│           │   │   │
│           │   │   # ========================= SIMILARITY & DEDUPLICATION HANDLERS (CONSOLIDATED) =========================
│           │   │   ├── similarity_handler.go     # 🔄 CONSOLIDATED: similarity + duplicate_check
│           │   │   │                             # POST /v1/proposals/:id/similarity/check, GET /v1/proposals/:id/similar, GET /v1/proposals/:id/duplicates
│           │   │   │                             # GET /v1/proposals/:id/differentiation, POST /v1/proposals/:id/clusters
│           │   │   │
│           │   │   # ========================= PORTFOLIO HANDLERS (CONSOLIDATED) =========================
│           │   │   ├── portfolio_handler.go      # 🔄 CONSOLIDATED: portfolio_link + portfolio_selector
│           │   │   │                             # POST /v1/proposals/:id/portfolio/link, POST /v1/proposals/:id/portfolio/auto-select
│           │   │   │                             # GET /v1/proposals/:id/portfolio, DELETE /v1/proposals/:id/portfolio/:itemId
│           │   │   │
│           │   │   # ========================= ENGAGEMENT HANDLERS (CONSOLIDATED) =========================
│           │   │   ├── engagement_handler.go     # 🔄 CONSOLIDATED: engagement + follow_up
│           │   │   │                             # POST /v1/proposals/:id/engagement/track, GET /v1/proposals/:id/engagement/interest
│           │   │   │                             # POST /v1/proposals/:id/follow-up/schedule, POST /v1/proposals/:id/bump
│           │   │   │
│           │   │   # ========================= MODERATION & COMPLIANCE HANDLERS =========================
│           │   │   ├── spam_handler.go           # POST /v1/proposals/:id/spam/scan, POST /v1/proposals/:id/spam/flag, GET /v1/proposals/:id/spam/status
│           │   │   ├── flag_handler.go           # POST /v1/proposals/:id/flag, POST /v1/proposals/:id/flag/review, GET /v1/proposals/:id/flags
│           │   │   ├── compliance_handler.go     # POST /v1/proposals/:id/compliance/check, GET /v1/proposals/:id/compliance/status
│           │   │
│           │   ├── routes/
│           │   │   │
│           │   │   # ========================= CORE PROPOSAL ROUTES =========================
│           │   │   ├── proposal_routes.go        # /v1/proposals, /v1/proposals/:id
│           │   │   ├── cover_letter_routes.go    # /v1/proposals/:id/cover-letter
│           │   │   ├── attachment_routes.go      # /v1/proposals/:id/attachments
│           │   │   ├── question_answer_routes.go # /v1/proposals/:id/answers
│           │   │   ├── milestone_routes.go       # /v1/proposals/:id/milestones
│           │   │   │
│           │   │   # ========================= LIFECYCLE ROUTES =========================
│           │   │   ├── expiration_routes.go      # /v1/proposals/:id/expiration
│           │   │   ├── withdrawal_routes.go      # /v1/proposals/:id/withdraw
│           │   │   ├── archive_routes.go         # /v1/proposals/:id/archive
│           │   │   ├── pipeline_routes.go        # /v1/proposals/pipeline
│           │   │   ├── recycling_routes.go       # /v1/proposals/:id/recycle
│           │   │   │
│           │   │   # ========================= BIDDING ROUTES =========================
│           │   │   ├── bid_routes.go             # /v1/proposals/:id/bid
│           │   │   ├── bid_strategy_routes.go    # /v1/bid-strategies
│           │   │   ├── auction_routes.go         # /v1/jobs/:jobId/auction
│           │   │   ├── bid_anomaly_routes.go     # /v1/bids/:bidId/anomaly
│           │   │   │
│           │   │   # ========================= CONNECTS & CREDITS ROUTES =========================
│           │   │   ├── connect_routes.go         # /v1/connects
│           │   │   ├── connect_refund_routes.go  # /v1/proposals/:id/connects/refund
│           │   │   │
│           │   │   # ========================= WORKFLOW & COLLABORATION ROUTES =========================
│           │   │   ├── negotiation_routes.go     # /v1/proposals/:id/negotiations
│           │   │   ├── invite_routes.go          # /v1/invites
│           │   │   ├── revision_routes.go        # /v1/proposals/:id/revisions
│           │   │   ├── collaboration_routes.go   # /v1/proposals/:id/team
│           │   │   │
│           │   │   # ========================= CLIENT INTERACTION ROUTES =========================
│           │   │   ├── interview_routes.go       # /v1/proposals/:id/interviews
│           │   │   ├── feedback_routes.go        # /v1/proposals/:id/feedback
│           │   │   ├── shortlist_routes.go       # /v1/jobs/:jobId/shortlist
│           │   │   ├── conversation_routes.go    # /v1/proposals/:id/conversation
│           │   │   │
│           │   │   # ========================= PERFORMANCE & ANALYTICS ROUTES (CONSOLIDATED) =========================
│           │   │   ├── performance_routes.go     # 🔄 /v1/proposals/:id/analytics, /v1/proposals/:id/engagement, /v1/proposals/:id/health, /v1/proposals/:id/ranking
│           │   │   │
│           │   │   # ========================= SIMILARITY & DEDUPLICATION ROUTES (CONSOLIDATED) =========================
│           │   │   ├── similarity_routes.go      # 🔄 /v1/proposals/:id/similarity, /v1/proposals/:id/duplicates
│           │   │   │
│           │   │   # ========================= PORTFOLIO ROUTES (CONSOLIDATED) =========================
│           │   │   ├── portfolio_routes.go       # 🔄 /v1/proposals/:id/portfolio
│           │   │   │
│           │   │   # ========================= ENGAGEMENT ROUTES (CONSOLIDATED) =========================
│           │   │   ├── engagement_routes.go      # 🔄 /v1/proposals/:id/engagement, /v1/proposals/:id/follow-up, /v1/proposals/:id/bump
│           │   │   │
│           │   │   # ========================= MODERATION & COMPLIANCE ROUTES =========================
│           │   │   ├── spam_routes.go            # /v1/proposals/:id/spam
│           │   │   ├── flag_routes.go            # /v1/proposals/:id/flag
│           │   │   ├── compliance_routes.go      # /v1/proposals/:id/compliance
│           │   │   │
│           │   │   # ========================= TEMPLATES & RATE CARDS ROUTES =========================
│           │   │   ├── template_routes.go        # /v1/templates
│           │   │   └── rate_card_routes.go       # /v1/rate-cards
│           │
│           └── router.go                     # 📝 HTTP router setup (mounts /v1 with platform-shared middleware + ETag)
│
├── config/                                   # 🆕 Configuration files
│   ├── default.yaml                          # Default configuration
│   ├── dev.yaml                              # Development overrides
│   └── prod.yaml                             # Production overrides
│
├── dapr/                                     # 🆕 Dapr components
│   ├── local/
│   │   ├── pubsub.yaml                       # Local Kafka pub/sub
│   │   └── statestore.yaml                   # Local state store
│   └── k8s/
│       ├── pubsub.yaml                       # Production Kafka with scopes: ["proposals-be"]
│       ├── statestore.yaml                   # Production state store with scopes: ["proposals-be"]
│       └── secrets.yaml                      # Secrets component
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                         # Service-specific errors
│   │   └── codes.go                          # Error codes (PROPOSAL_NOT_FOUND, INSUFFICIENT_CONNECTS, OUTBID, etc.)
│   │
│   ├── logger/                               # ❌ REMOVED
│   │   └── README.md                         # Points to platform-shared/logging
│   │
│   ├── utils/
│   │   ├── validator.go                      # Local validation utilities
│   │   ├── sanitizer.go                      # Sanitize user input
│   │   ├── bid_calculator.go                 # Bid calculation utilities
│   │   ├── text_analyzer.go                  # Text analysis utilities
│   │   ├── score_calculator.go               # Generic scoring utilities
│   │   └── etag.go                                 🆕  # canonical ETag helpers (hashing, weak/strong)
│   │
│   └── constants/
│       ├── events.go                         # ❌ REMOVED - Use contracts/events
│       ├── topics.go                         # ❌ REMOVED - Use contracts/events
│       ├── bid_types.go                      # Bid type constants
│       ├── proposal_status.go                # Proposal status constants
│       ├── pipeline_stages.go                # Pipeline stage constants
│       └── refund_reasons.go                 # Refund reason constants
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                   # Kubernetes Deployment
│       ├── service.yaml                      # Kubernetes Service
│       ├── configmap.yaml                    # ConfigMap
│       ├── secrets.yaml                      # Secrets
│       ├── hpa.yaml                          # Horizontal Pod Autoscaler
│       ├── pdb.yaml                          # Pod Disruption Budget
│       └── servicemonitor.yaml               # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                        # Local dev setup
│   ├── get-secrets.sh                        # Fetch secrets
│   └── seed-data.sh                          # Seed initial data
│   ├── migrate.sh                                  🆕  # SQL migrations (dev/prod flags)
│   ├── openapi-diff.sh                             🆕  # fail CI on breaking API changes
│   ├── schema-diff.sh                              🆕  # DB schema guardrail (pg_dump + mig verify)
│   ├── generate-sdks.sh                            🆕  # produce /sdk clients from OpenAPI
│   └── sbom-sign.sh 
│
├── tests/
│   ├── reliability/                                🆕
│   │   ├── projections_replay_test.go              🆕  # rebuild performance read models from events
│   │   └── outbox_dispatcher_test.go               🆕  # at-least-once + idempotency assertions
│   ├── property/                                   🆕
│   │   └── performance_property_test.go            🆕  # fuzz tests for scoring/health kernel
|   |
│   ├── unit/
│   │   ├── domain/
│   │   │   ├── proposal_test.go              # Proposal entity tests
│   │   │   ├── bid_test.go                   # Bid entity tests
│   │   │   ├── performance_test.go           # 🔄 CONSOLIDATED performance tests
│   │   │   ├── similarity_test.go            # 🔄 CONSOLIDATED similarity tests
│   │   │   ├── portfolio_test.go             # 🔄 CONSOLIDATED portfolio tests
│   │   │   ├── engagement_test.go            # 🔄 CONSOLIDATED engagement tests
│   │   │   └── collaboration_test.go         # 🔄 CONSOLIDATED collaboration tests
│   │   │
│   │   ├── application/
│   │   │   ├── proposal_service_test.go      # Proposal service tests
│   │   │   ├── bid_service_test.go           # Bid service tests
│   │   │   ├── performance_service_test.go   # 🔄 CONSOLIDATED performance service tests
│   │   │   ├── similarity_service_test.go    # 🔄 CONSOLIDATED similarity service tests
│   │   │   └── engagement_service_test.go    # 🔄 CONSOLIDATED engagement service tests
│   │   │
│   │   └── infrastructure/
│   │       ├── proposal_repository_test.go   # Proposal repository tests
│   │       ├── performance_repository_test.go # 🔄 CONSOLIDATED performance repository tests
│   │       └── kafka_producer_test.go        # Kafka producer tests
│   │       ├── scheduler_expiration_task_test.go
│   │       │                               # ✅ Verifies idempotency, respects now(), and no double notifications on retries
│   │       └── projections_performance_projector_test.go
│   │                                       # ✅ Event replay builds the same read model; partial updates are additive and ordered
│   │
│   │
│   ├── integration/
│   │   ├── handlers/
│   │   │   ├── proposal_handler_test.go      # Proposal handler integration tests
│   │   │   ├── bid_handler_test.go           # Bid handler integration tests
│   │   │   └── performance_handler_test.go   # 🔄 CONSOLIDATED performance handler tests
│   │   │
│   │   │─ worker_tasks_integration_test.go
│   │   │                                            # 🚦 Spins minimal DI container, runs cron once against test DB/Redis; asserts side-effects & metric
│   │   └── repositories/
│   │       ├── proposal_repository_integration_test.go # Proposal repository integration tests
│   │       └── performance_repository_integration_test.go # 🔄 CONSOLIDATED performance repository tests
│   │
│   └── e2e/
│       └── scenarios/
│           ├── proposal_submission_test.go   # E2E proposal submission flow
│           ├── bidding_flow_test.go          # E2E bidding flow
│           ├── negotiation_flow_test.go      # E2E negotiation flow
│           ├── collaboration_flow_test.go    # 🔄 E2E team collaboration flow
│           ├── performance_tracking_test.go  # 🔄 E2E performance tracking flow
│           └── chaos_event_delivery_test.go            🆕  # simulate duplicates, reordering, delays
│
│
├── docs/
│   ├── README.md                             # Service overview
│   ├── API.md                                # API documentation
│   ├── ARCHITECTURE.md                       # Service architecture
│   ├── MIGRATIONS.md                         # Migration history
│   ├── SCHEMA.md                             # Database schema
│   ├── RUNBOOK.md                            # Operations runbook
│   │
│   # ========================= PLATFORM & OPERATIONS DOCS =========================
│   ├── API_VERSIONING.md                           🆕  # HTTP contract versioning & deprecation policy
│   ├── CACHING.md                                  🆕  # keys, TTLs, SWR, invalidation events
│   ├── DATA_RETENTION.md                           🛡️🆕 # per-domain retention windows (proposals, messages, analytics)
│   ├── ERASURE.md                                  🛡️🆕 # GDPR/CCPA erasure hooks & playbooks
│   ├── SLOS.md                                     📏🆕 # target P99s, projection lag, queue delay
│   ├── OUTBOX.md                                   🆕  # design, retries, DLQ
│   ├── RELEASE_CHECKLIST.md                        🆕  # openapi diff, schema diff, migrations plan
│   └── NAMING.md                                   ♻️  # package snake_case; HTTP kebab-case
│   │
│   # ========================= CORE FEATURE DOCUMENTATION =========================
│   ├── EVENTS.md                             # Events (published & consumed)
│   ├── proposal-lifecycle.md                 # Proposal lifecycle documentation
│   ├── bidding-system.md                     # Bidding system documentation
│   ├── connects-system.md                    # Connects/credits documentation
│   ├── collaboration.md                      # 🔄 CONSOLIDATED team collaboration documentation
│   │── EVENT_VERSIONING.md              # 🧾 Guide: version public events (e.g., ProposalSubmitted.v1); change rules, deprecation policy, and migration tips
│   │
│   # ========================= CONSOLIDATED FEATURE DOCUMENTATION =========================
│   ├── performance-analytics.md              # 🔄 CONSOLIDATED: analytics + engagement + conversion + health + ranking
│   ├── similarity-deduplication.md           # 🔄 CONSOLIDATED: similarity + duplicate detection
│   ├── portfolio-management.md               # 🔄 CONSOLIDATED: portfolio linking + auto-selection
│   ├── engagement-followup.md                # 🔄 CONSOLIDATED: engagement tracking + follow-ups
│   │
│   # ========================= 🆕 EXTERNALIZED SERVICES DOCUMENTATION =========================
│   ├── EXTERNAL_SERVICES.md                  # 🆕 Documentation for all externalized services
│   │   │
│   │   # Documents integration with:
│   │   # - intelligence-be (strategy, optimization, predictions, recommendations, keywords, personalization, A/B testing)
│   │   # - risk-be (risk assessment, insurance, escrow, guarantee, dispute prediction)
│   │   # - contracts-be (terms, NDA, IP rights, payment terms, pricing, redlining, approval workflows)
│   │   # - procurement-be (RFP, quote, procurement, evaluation rubric, security questionnaire)
│   │   # - marketplace-be (boost, visibility budget, premium features, fees, subscriptions, marketplace listings)
│   │   # - clients-be / crm-be (client insights, client preferences)
│   │   # - calendar-be (deadline tracking, sync)
│   │   # - time-tracking-be (time estimates, tracking)
│   │   # - localization-be (translations)
│   │   # - integrations-be (external integrations)
│   │   # - sharing-be (secure links, access logs)
│   │   # - compliance-be (compliance checks, verification)
│   │
│   └── migration-guide.md                    # 🆕 Guide for migrating from old bloated structure to new lean structure
│
├── .github/
│   └── workflows/
│       ├── ci.yml                            # CI workflow
│       ├── cd.yml                            # CD workflow
│       ├── contract-ci.yml                         🆕  # openapi-diff + event schema checks
│       ├── security.yml                            🆕  # golangci-lint, govulncheck, trivy, cosign verify
│       └── load-tests.yml                          🆕  # k6/gatling smoke against /healthz & hot paths
│
│
├── go.mod                                    # 📝 Dependencies (imports pkg/auth, platform-shared, contracts/events)
├── go.sum
├── .env.example                              # Example environment variables
├── Makefile                                  # Build targets
├── Dockerfile                                # Multi-stage Docker build
├── .dockerignore                             # Docker ignore patterns
├── .gitignore                                # Git ignore patterns
└── README.md                                 # Service README


```
---

---

```

## 📦 **4️⃣ contracts-be (Contract Management Service)**

```
apps/be/contracts-be/
│
├── cmd/
│   ├── api/
│   │   └── main.go  # 📝 API entrypoint; init Gin, Dapr, Postgres (platform-shared/logging, internal/config); serve /v1/*; ETag behind flag
│   │
│   └── worker/  # 🆕 background jobs / cron / outbox dispatcher
│       └── main.go  # 🆕 boot DI; run cron/bindings; outbox dispatcher; leader election; runs: milestone auto-release, SLA monitors, contract expiry, dispute escalation, inbox consumer, timesheet auto-submit
│
├── internal/
│   │
│   # ========================================================================================
│   # 🔧 CONFIGURATION LAYER (Load First)
│   # ========================================================================================
│   ├── config/
│   │   ├── schema.go  # Config struct: Server(Port, Host, Read/Write timeouts); DB(Host, Port, Name, User, Password, Max/MinConns); Redis; Kafka(Brokers, Topics, CGs); Storage(Endpoint, Bucket, Keys); FeatureFlags; Observability(Tracing, MetricsPort, LogLevel)
│   │   ├── loader.go  # Load config from env/yaml/flags
│   │   ├── validator.go  # Validate configuration
│   │   └── feature_flags.go  # 🆕 toggles: enable_contract_drafts, enable_etag, enable_auto_release, enable_sla_tracking, enable_agency_contracts, enable_direct_contracts, enable_work_diary, enable_e_signatures, enable_auto_renewal, enable_compliance_checks
│   │
│   # ========================================================================================
│   # 🧩 DEPENDENCY INJECTION (Load Second)
│   # ========================================================================================
│   ├── ioc/
│   │   ├── container.go  # 🆕 DI container: Config, DB, Redis, Kafka(prod/cons), repositories, services, handlers, external clients, observers, scheduler
│   │   └── wiring.go  # 🆕 Wire order: config→infra→repos→services→handlers; feature-flagged wiring; observers toggle; outbox dispatcher; scheduled tasks
│   │
│   # ========================================================================================
│   # 🏗️ INFRASTRUCTURE LAYER (Load Second)
│   # ========================================================================================
│   ├── infrastructure/
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # 🔄 COORDINATION - Distributed Primitives
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── coordination/
│   │   │   ├── leader_election.go  # 🆕 single-active worker via Redis/etcd; leader election; failover; health checks
│   │   │   └── distlock.go  # 🆕 distributed locks: Redis SETNX+expiry; PG advisory fallback; timeouts; deadlock prevention
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # 💾 PERSISTENCE - Database Layer
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go  # PGX pool; max/min conns; health checks; reconnection; query timeouts
│   │   │       ├── transaction.go  # Tx helpers: Begin/Commit/Rollback; savepoints; ctx propagation; auto-rollback on panic
│   │   │       ├── migrations.go  # 📝 auto-migrations (dev only); skip in prod; integrates with version.go & safety.go
│   │   │       ├── version.go  # 🆕 schema version tracking; schema_versions table; compare current vs target; optional rollback
│   │   │       ├── safety.go  # 🆕 pre-migration checks: destructive ops guard; backup hints; long-run warnings; concurrent index creation
│   │   │       ├── outbox_store.go  # 🆕 tx-coupled outbox: insert events; mark published; cleanup; DLQ for failures
│   │   │
│   │   │       # ===== CORE CONTRACT REPOSITORIES =====
│   │   │       ├── contract_repository.go  # Contract CRUD/queries: by ID/freelancer/client/status/type; active/expiring; search+filters
│   │   │       ├── sow_repository.go  # 🆕 SOW: create/update; by contract; version history/latest; approve/reject version
│   │   │       ├── financial_hold_repository.go  # 🆕 Holds: place/release/extend; list by contract/active/expiring; total held
│   │   │
│   │   │       # ===== MILESTONE & DELIVERABLE REPOSITORIES =====
│   │   │       ├── milestone_repository.go  # Milestones: CRUD; pending/overdue; totals/stats
│   │   │       ├── deliverable_repository.go  # Deliverables: CRUD; by milestone/status; pending review; rejected with revisions
│   │   │
│   │   │       # ===== TIME TRACKING REPOSITORIES =====
│   │   │       ├── timesheet_repository.go  # Timesheets: CRUD; by contract+week/status; pending approval/disputed; totals
│   │   │       ├── workdiary_repository.go  # Work diary: entries by date range; activity summary; screenshot URLs; delete screenshot
│   │   │
│   │   │       # ===== TEMPLATE REPOSITORIES =====
│   │   │       ├── template_repository.go  # Templates: CRUD; public/by owner/category; search; popular (usage)
│   │   │
│   │   │       # ===== CONTRACT CHANGE REPOSITORIES =====
│   │   │       ├── amendment_repository.go  # Amendments: CRUD; by contract; pending; history/latest
│   │   │       ├── pause_repository.go  # Pause/resume: record; history; total paused days; current pause
│   │   │       ├── termination_repository.go  # Terminations: CRUD; by contract; pending settlement; stats by reason/type
│   │   │
│   │   │       # ===== DISPUTE REPOSITORIES =====
│   │   │       ├── dispute_repository.go  # Disputes: CRUD; by contract/status; open/escalated; stats/resolution time
│   │   │
│   │   │       # ===== DOCUMENT & SIGNATURE REPOSITORIES =====
│   │   │       ├── attachment_repository.go  # 🆕 Attachments: CRUD+versioning; upload/delete/get; by contract/type; version history/latest
│   │   │       ├── signature_repository.go  # 🆕 E-signatures: request/record; by contract/signer; status/pending; verification
│   │   │
│   │   │       # ===== FINANCIAL REPOSITORIES =====
│   │   │       ├── budget_repository.go  # 🆕 Budget: create/update/by contract; record spending; remaining; trends/forecast; thresholds/over-budget
│   │   │       ├── invoice_repository.go  # 🆕 Invoices: generate/update/by contract/status; mark paid; overdue; stats
│   │   │
│   │   │       # ===== COMPLIANCE & LEGAL REPOSITORIES =====
│   │   │       ├── sla_repository.go  # 🆕 SLA: define/update/by contract; record metrics; detect breaches; lists/report; penalties/rewards
│   │   │       ├── agency_repository.go  # 🆕 Agency: by contract; members/roles; billing split; member payments
│   │   │       ├── compliance_repository.go  # 🆕 Compliance: add/verify; audits/history; list/status; expiring items
│   │   │       ├── ip_rights_repository.go  # 🆕 IP rights: assign/grant; update terms; license history
│   │   │       ├── nda_repository.go  # 🆕 NDA: sign/breach/enforce; active/expired tracking
│   │   │
│   │   │       # ===== REPORTING & COLLABORATION REPOSITORIES =====
│   │   │       ├── performance_repository.go  # 🆕 KPIs: define/record; scores/trends; benchmarks/report
│   │   │       ├── report_repository.go  # 🆕 Reports: generate/get/list/schedule; analytics data
│   │   │       ├── feedback_repository.go  # 🆕 Feedback: submit/get/list; summaries/averages
│   │   │       └── workroom_repository.go  # 🆕 Workroom: tasks/notes/files; list/query
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # 📨 MESSAGING - Event-Driven Communication
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── messaging/
│   │   │   ├── outbox/
│   │   │   │   ├── dispatcher.go  # 🆕 poll outbox; publish to Kafka; mark published; retries (exponential); DLQ; cleanup
│   │   │   │   └── metrics.go  # 🆕 publish lag; failed attempts; retry distribution; DLQ size
│   │   │   ├── topics.go  # ♻️ Kafka topic constants (re-export contracts/events): lifecycle, milestones, financial, timesheet, disputes, compliance, collaboration
│   │   │   ├── producer.go  # Kafka producer wrapper: partition by contract_id; compression (lz4); async/sync
│   │   │   └── consumer.go  # Kafka consumer wrapper: subscribe; CG mgmt; offset strategies; retries/error handling
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # 🗄️ CACHING - Performance
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── client.go  # Redis client wrapper: pooling; reconnection; timeouts
│   │   │       ├── keys.go  # Key patterns: contract:{id}, contract:{id}:milestones/status/budget; user:{id}:contracts; freelancer:{id}:active_contracts
│   │   │       ├── ttl.go  # TTLs: short 5m status/budget; medium 1h details; long 24h templates/static
│   │   │       ├── singleflight.go  # 🆕 stampede protection: dedupe concurrent requests; share result
│   │   │       └── invalidation_rules.go  # 🆕 event→keys mapping (status change, milestone release, dispute opened, budget updated, amendment applied)
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # ⏰ SCHEDULER - Cron & Background Jobs
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── scheduler/
│   │   │   ├── cron.go  # ♻️ cron runner: wrap with distlock; jitter; errors/logging; health
│   │   │   └── tasks/
│   │   │       ├── task_guard.go  # 🆕 idempotency tokens; last-run watermark; timeouts; graceful shutdown
│   │   │       ├── milestone_auto_release.go  # hourly; auto-release approved > grace(3d) if no disputes; notify
│   │   │       ├── sla_breach_monitor.go  # every 15m; check response/deadline SLAs; record breaches; penalties; alerts
│   │   │       ├── contract_expiry.go  # daily midnight; expiring in 7d; renewal reminders; mark expired; auto-renew if enabled
│   │   │       ├── dispute_escalation.go  # twice daily; stale disputes > 7d; escalate; notify
│   │   │       ├── timesheet_reminder.go  # Fridays 17:00; hourly contracts without weekly timesheet; reminders
│   │   │       └── budget_alert.go  # every 4h; thresholds 50/75/90%; alerts; overrun prediction
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # 📊 OBSERVABILITY - Metrics & Tracing
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── observability/
│   │   │   ├── metrics.go  # 🆕 Prometheus: RED (api reqs, errors, duration P50/95/99); USE (db pool, cache hit, outbox lag); business (contract_value_total, milestone_release_latency, dispute_resolution_time, sla_breach_rate, budget_overrun_count)
│   │   │   ├── tracing.go  # 🆕 Jaeger/OTel helpers: spans; attrs (hashed contract_id/user_id, event_id, correlation_id); errors; baggage
│   │   │   └── slo_monitor.go  # 🆕 SLOs: API P99<500ms(99.9%), projection lag<1s(99%), outbox lag<5s(99.5%), creation success>99.9%, milestone release<24h(99%); alert on breach
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # 📦 STORAGE - MinIO/S3
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── storage/
│   │   │   ├── client.go  # MinIO/S3 client: upload/download; presigned URLs; delete; list by prefix
│   │   │   ├── buckets.go  # Buckets: contracts-documents; contracts-attachments; work-diary-screenshots; dispute-evidence
│   │   │   └── policies.go  # Access: parties-only; admin-read; time-limited links; immutable after signing
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # 🔌 PLATFORM CLIENTS - External Adapters
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   └── platform/
│   │       └── clients/
│   │           ├── users_client.go  # 🆕 verify user status; profiles; stats update; circuit breaker; retries
│   │           ├── financial_client.go  # 🆕 escrow create/release; holds place/release; verify balance; tx history
│   │           ├── communications_client.go  # 🆕 notifications: contract/milestone/dispute/reminder/budget; async fire-and-forget
│   │           ├── jobs_client.go  # job details; job status check; mark filled; update contract count
│   │           ├── proposals_client.go  # proposal details; mark accepted; deduct connects
│   │           ├── reviews_client.go  # trigger review on completion; can leave review; link to contract
│   │           └── storage_client.go  # storage-be API alternative: upload/get URL/delete; virus scan integration
│   │
│   # ========================================================================================
│   # 🏛️ DOMAIN LAYER - Business Entities & Rules (Load Third)
│   # ========================================================================================
│   ├── domain/
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # 📋 CORE CONTRACT DOMAIN
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── contract/
│   │   │   ├── entity.go  # Aggregate: ID(ULID), JobID, ProposalID?, FreelancerID, ClientID, Type(Fixed/Hourly/Milestone/Retainer), Status(Draft/Active/Paused/Disputed/Completed/Terminated/Expired), Terms(VO), Amount/Currency, Start/End, Created/Updated/Version; methods Activate/Pause/Resume/Complete/Terminate; guards CanBeAmended/Terminated/Paused; IsExpired/Active
│   │   │   ├── terms.go  # Terms VO: PaymentSchedule(upfront/milestone/weekly/monthly), Deliverables[], Scope, Confidentiality, IPRights, Termination, DisputeResolution
│   │   │   ├── enums.go  # Enums: ContractType, ContractStatus, PaymentSchedule
│   │   │   ├── value_objects.go  # VOs: ContractAmount, ContractDuration, ContractRate
│   │   │   ├── errors.go  # ErrContractNotFound, ErrInvalidTerms, ErrAlreadyActive, ErrUnauthorized, ErrCannotPause/Terminate, ErrExpired
│   │   │   ├── repository.go  # Interface: Create/Update/Delete/GetByID; List by freelancer/client/status/type; active/expiring; Search(ListFilter)
│   │   │   ├── list_filter.go  # Filters: status/type; amount min/max; start/end ranges; freelancer/client; pagination/sort
│   │   │   └── events.go  # ContractCreated/Activated/Paused/Resumed/Completed/Terminated/Expired/Renewed/Disputed/StatusChanged
│   │   │
│   │   ├── sow/
│   │   │   ├── entity.go  # SOW aggregate: ID, ContractID, Scope, Objectives, Deliverables, Timeline, AcceptanceCriteria, CurrentVersion, Versions[]; methods CreateVersion/Approve/Reject
│   │   │   ├── version.go  # Version: Number, CreatedAt/By, Changes diff, ApprovalStatus(Pending/Approved/Rejected), ApprovedBy/At
│   │   │   ├── scope.go  # Scope: WBS, inclusions/exclusions, assumptions
│   │   │   ├── errors.go  # ErrSOWNotFound, ErrSOWVersionNotFound
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # SOWCreated/Updated/Approved/VersionCreated
│   │   │
│   │   ├── financial_hold/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, Type(Risk/Dispute/Compliance/Chargeback/Manual), Amount/Currency, Reason, PlacedAt/By, ExpiresAt, ReleasedAt/By, Status(Active/Released/Expired); methods CanRelease/Release/Extend
│   │   │   ├── hold_type.go  # Types & rules
│   │   │   ├── release_rules.go  # Time/Event/Approval-based auto-release rules
│   │   │   ├── errors.go  # ErrHoldNotFound, ErrAlreadyActive, ErrInsufficientFunds
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # FinancialHoldPlaced/Released/Expired
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # 🎯 MILESTONES & DELIVERABLES
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── milestone/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, Title/Desc, Amount/Currency, Due/Submitted/Approved/Released, Status(Pending/InProgress/UnderReview/Approved/Rejected/Released/Disputed), Order, Deliverables[]; methods Submit/Approve/Reject/Release/Dispute; CanAutoRelease
│   │   │   ├── status.go  # Transitions
│   │   │   ├── approval.go  # ClientApproval, feedback, revision cycle
│   │   │   ├── auto_release.go  # GracePeriod(3d); checks no disputes; release date calc
│   │   │   ├── errors.go  # ErrMilestoneNotFound, ErrAlreadyReleased, ErrNotApproved
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # Created/Submitted/Approved/Rejected/Released/Disputed
│   │   │
│   │   ├── deliverable/
│   │   │   ├── entity.go  # Entity: ID, MilestoneID, Title/Desc, FileURLs[], SubmittedAt/ReviewedAt, Status(Pending/Submitted/UnderReview/Accepted/Rejected/Revision), ClientFeedback; methods Submit/Accept/Reject/RequestRevision
│   │   │   ├── status.go  # Enum
│   │   │   ├── review.go  # ReviewFeedback/RevisionRequest/AcceptanceCriteria
│   │   │   ├── errors.go  # ErrDeliverableNotFound, ErrNoFilesAttached
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # Submitted/Accepted/Rejected/RevisionRequested
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # ⏱️ TIME TRACKING
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── timesheet/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, WeekStart/End, Entries[], TotalHours/Amount, Status(Draft/Submitted/Approved/Rejected/Disputed), Submitted/Approved/RejectedAt, ApprovalNotes; methods AddEntry/Submit/Approve/Reject/Dispute; calculators
│   │   │   ├── entry.go  # TimeEntry VO: Date, Hours/Minutes, Description, Task?, Billable; validations (≤24h/day, no dup same day)
│   │   │   ├── approval.go  # ApprovalDecision/Notes; disputed hours
│   │   │   ├── dispute.go  # Reasons; disputed entries; resolution outcomes
│   │   │   ├── errors.go  # ErrTimesheetNotFound, ErrAlreadySubmitted, ErrInvalidHours
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # Created/Submitted/Approved/Rejected/Disputed
│   │   │
│   │   ├── workdiary/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, Date, Entries[], TotalTrackedHours, ScreenshotURLs[], PrivacySettings; methods RecordEntry/AddManualTime/DeleteScreenshot
│   │   │   ├── entry.go  # 10-min slots: timestamp, duration, activity%, screenshotURL?, keystrokes/mouse, app title
│   │   │   ├── screenshot.go  # URL, capturedAt, blurred, deletedAt, retention(30d)
│   │   │   ├── privacy.go  # ScreenshotFrequency, BlurScreenshots, RedactAppTitles, ExcludedApps[]
│   │   │   ├── manual_time.go  # Offline entries; reason; requires approval
│   │   │   ├── errors.go  # ErrWorkDiaryNotFound, ErrScreenshotNotFound
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # EntryRecorded/ScreenshotCaptured/ManualTimeAdded/PrivacyUpdated
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # 📄 TEMPLATES
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── template/
│   │   │   ├── entity.go  # Aggregate: ID, Name/Desc, Type, Terms, Clauses[], Category, IsPublic, OwnerID, UsageCount; methods Customize/Publish/Unpublish/Clone
│   │   │   ├── clause.go  # Payment/Termination/IP/Confidentiality/Dispute/Warranty clauses
│   │   │   ├── category.go  # Enums (Development/Design/Writing/Marketing/Other)
│   │   │   ├── customization.go  # Customizable/Required fields; defaults
│   │   │   ├── errors.go  # ErrTemplateNotFound, ErrTemplateNotPublic
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # Created/Updated/Published/Used
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # ✏️ AMENDMENTS & CHANGES
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── amendment/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, Type(Scope/Budget/Timeline/Terms/Rate), Changes map, Reason, RequestedBy, Status(Pending/Approved/Rejected/Applied), Timestamps, ApprovalsByParty; methods Approve/Reject/Apply; IsBilaterallyApproved
│   │   │   ├── amendment_type.go  # Enum
│   │   │   ├── approval.go  # Bilateral workflow; timeout(30d); reminders(7d)
│   │   │   ├── version_control.go  # Previous/New version; diff; rollback support
│   │   │   ├── errors.go  # ErrAmendmentNotFound/Rejected/NotBilaterallyApproved
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # Requested/Approved/Rejected/Applied
│   │   │
│   │   ├── termination/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, Type(Mutual/ForCause/ForConvenience/Breach/NonPerformance), Reason, InitiatedBy, Dates, NoticePeriod, Status, FinalPayment/Refund; methods Approve/Complete/Dispute; settlement calculator
│   │   │   ├── termination_type.go  # Enum
│   │   │   ├── notice.go  # Notice calculation; earliest termination; waiver
│   │   │   ├── settlement.go  # Final payment/refund/penalties/severance
│   │   │   ├── errors.go  # ErrTerminationNotAllowed, ErrInvalidType
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # Initiated/Approved/Completed/Disputed
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # ⚖️ DISPUTES
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── dispute/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, Type(Quality/Payment/Scope/Timeline/Communication/Breach), Description, RaisedBy, Evidence[], Status(Open/UnderReview/Escalated/Resolved/Closed), Resolution, ResolvedAt/By; methods AddEvidence/Escalate/Resolve/Close
│   │   │   ├── dispute_type.go  # Enum
│   │   │   ├── evidence.go  # Evidence: type, file info, description, submitted by/at
│   │   │   ├── resolution.go  # Resolution types (Mediation/Arbitration/Refund/ReWork), outcomes, financial adjustments
│   │   │   ├── escalation.go  # Levels: Mediation(7d)/Admin(14d)/Arbitration(30d); auto-escalate
│   │   │   ├── errors.go  # ErrDisputeNotFound, ErrAlreadyClosed
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # Opened/EvidenceSubmitted/Escalated/Resolved/Closed
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # 📎 DOCUMENTS & SIGNATURES
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── attachment/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, Type(SOW/NDA/SignedContract/Supporting/Amendment), FileName/URL/Size/Mime, UploadedBy/At, Version, IsSigned; methods CreateVersion/GetLatest/Delete
│   │   │   ├── attachment_type.go  # Enum
│   │   │   ├── version.go  # Versioning: number, createdAt/By, changes, isCurrent
│   │   │   ├── errors.go  # ErrAttachmentNotFound, ErrFileTooLarge
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # Uploaded/Deleted/VersionCreated
│   │   │
│   │   ├── signature/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, SignerID/Role, SignatureData, IP/UserAgent, SignedAt, Status(Pending/Signed/Declined/Expired), VerificationHash; methods Sign/Verify/Decline
│   │   │   ├── signer.go  # Name/Email/Role; signing order; required flag
│   │   │   ├── signing_flow.go  # Sequential signing; notifications; expiration(30d); AllPartiesSigned
│   │   │   ├── verification.go  # Hash gen/verify; audit trail
│   │   │   ├── errors.go  # ErrSignerNotAuthorized, ErrExpired
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # Requested/Signed/Verified/AllPartiesSigned/Declined
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # 💰 FINANCIAL
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── budget/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, TotalBudget/Currency, Spent/Remaining, Thresholds[], Alerts[], BurnRate; methods RecordSpending/CheckThresholds/Forecast
│   │   │   ├── threshold.go  # % thresholds; reached flag/at; notification sent
│   │   │   ├── alert.go  # Type(ThresholdReached/BudgetExceeded/ForecastOverrun); message; sentAt/to
│   │   │   ├── forecast.go  # BurnRate; exhaustion date; overrun; recommendations
│   │   │   ├── errors.go  # ErrBudgetExceeded, ErrInvalidBudget
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # BudgetUpdated/ThresholdReached/BudgetExceeded/BudgetAdjusted
│   │   │
│   │   ├── invoice/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, InvoiceNumber, LineItems[], Subtotal/Tax/Total/Currency, Due/Generated/PaidAt, Status(Draft/Sent/Paid/Overdue/Cancelled), PaymentMethod; methods Generate/Send/MarkPaid/Cancel
│   │   │   ├── line_item.go  # Type(Milestone/Hours/Expense/Fee); desc; qty/price/amount; reference
│   │   │   ├── tax.go  # Jurisdiction calc; rates; exemptions; TaxID
│   │   │   ├── status.go  # Enum & transitions
│   │   │   ├── errors.go  # ErrAlreadyPaid, ErrNotFound
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # Generated/Sent/Paid/Overdue/Cancelled
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # ⚖️ COMPLIANCE & LEGAL
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── sla/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, Metrics[], Thresholds[], Breaches[], Penalties[], Rewards[]; methods RecordMetric/DetectBreach/ApplyPenalty/GrantReward
│   │   │   ├── metric.go  # Type(Response/Delivery/Availability/Quality), target, actual, measuredAt
│   │   │   ├── breach.go  # MetricType, threshold, actual, detected/resolvedAt, severity
│   │   │   ├── penalty.go  # Type(Refund/Credit/Discount); amount/currency; reason; appliedAt
│   │   │   ├── reward.go  # Type(Bonus/Badge/PrioritySupport); amount?; reason; grantedAt
│   │   │   ├── errors.go  # ErrSLANotDefined, ErrInvalidMetric
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # BreachDetected/PenaltyApplied/RewardEarned/SLAReportGenerated
│   │   │
│   │   ├── agency/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, AgencyID, Members[], Roles[], SplitRules; methods Add/Remove/Update/CalculateSplit
│   │   │   ├── member.go  # Member details: role/permissions; bill rate; split%; joined/left
│   │   │   ├── role.go  # Role definitions & permissions
│   │   │   ├── billing.go  # Consolidated invoicing; split payments; rules; member payment calc
│   │   │   ├── errors.go  # ErrAgencyNotFound, ErrMemberNotAuthorized
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # MemberAdded/Removed/RoleChanged/BillingUpdated
│   │   │
│   │   ├── compliance/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, Requirements[], Status, Audits[], ExpiresAt; methods AddRequirement/Verify/RunAudit
│   │   │   ├── requirement.go  # Type(NDA/Background/Certification/Insurance), desc, required, status, verifiedAt/By, documentURL
│   │   │   ├── audit.go  # Type(Periodic/Triggered/Random), date/by, findings[], status, remediation plan
│   │   │   ├── jurisdiction.go  # GDPR/CCPA/LaborLaws; GetApplicableLaws(location)
│   │   │   ├── errors.go  # ErrComplianceNotMet, ErrAuditFailed
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # RequirementAdded/ComplianceMet/ComplianceFailed/AuditCompleted
│   │   │
│   │   ├── performance/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, KPIs[], Scores[], Trends[], Benchmarks[]; methods DefineKPI/Record/Calculate/Compare
│   │   │   ├── kpi.go  # KPI types; targets; weights; frequency
│   │   │   ├── score.go  # Overall/Per-KPI scores; calculatedAt; grade
│   │   │   ├── trend.go  # Period; historical scores; direction; predicted score
│   │   │   ├── benchmark.go  # Industry/category; averages; top percentile; comparisons
│   │   │   ├── errors.go  # ErrKPINotDefined, ErrInvalidScore
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # PerformanceUpdated/KPIBreached/TrendAlerted/BenchmarkCompared
│   │   │
│   │   ├── negotiation/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID(draft), Offers[], CounterOffers[], Status, History[], ExpiresAt; methods MakeOffer/Counter/Accept/Reject
│   │   │   ├── offer.go  # Terms/amount/currency/timeline; proposedBy/At; expires; notes
│   │   │   ├── counter_offer.go  # Link original; changes; rationale; proposedBy/At
│   │   │   ├── workflow.go  # MaxRounds; response timeout(7d) auto-reject; reminders
│   │   │   ├── errors.go  # ErrNotFound, ErrInvalidOffer, ErrExpired
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # OfferMade/CounterOffered/Accepted/Rejected/Expired
│   │   │
│   │   ├── renewal/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, RenewalType(Auto/Manual/Extension), NewTerms, RenewalDate, AutoRenew, Notifications[], Status; methods Request/Approve/Reject/Extend
│   │   │   ├── renewal_type.go  # Enum
│   │   │   ├── auto_renewal.go  # Conditions; opt-out(30d); notifications(60/30/7d); terms adjustment
│   │   │   ├── extension.go  # Duration; additional budget; reason; bilateral approval
│   │   │   ├── errors.go  # ErrNotEligible, ErrExtensionDenied
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # RenewalRequested/Approved/AutoRenewed/ContractExtended
│   │   │
│   │   ├── ip_rights/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, AssignmentType(FullTransfer/License/WorkForHire), RightsGranted[], Exclusions[], TransferDate, License; methods AssignIP/GrantLicense/UpdateTerms
│   │   │   ├── assignment.go  # Types description
│   │   │   ├── license.go  # License: type, duration, territory, usage, sublicensing
│   │   │   ├── protection.go  # Confidentiality/NonCompete/Attribution
│   │   │   ├── errors.go  # ErrIPConflict, ErrAssignmentInvalid
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # IPAssigned/IPLicensed/IPProtectionApplied
│   │   │
│   │   ├── nda/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, Parties[], ConfidentialInfo, Duration, Breaches[], Penalties, SignedAt, ExpiresAt; methods Sign/ReportBreach/Enforce
│   │   │   ├── terms.go  # Definitions/Obligations/Exceptions/ReturnOfMaterials
│   │   │   ├── breach.go  # Type; description; reportedBy/At; evidence[]; penalties/remedies
│   │   │   ├── errors.go  # ErrNDANotSigned, ErrBreachDetected
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # NDASigned/NDABreached/NDAEnforced
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # 📊 REPORTING & COLLABORATION
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── report/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, Type(Progress/Performance/Financial/Compliance/TimeTracking), GeneratedAt/By, Data map, Format(PDF/Excel/JSON), FileURL; methods Generate/Schedule/Export
│   │   │   ├── report_type.go  # Enum
│   │   │   ├── analytics.go  # Efficiency/Risks/Predictions/Trends
│   │   │   ├── errors.go  # ErrReportNotGenerated, ErrInvalidType
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # ReportGenerated/AnalyticsUpdated/ReportScheduled
│   │   │
│   │   ├── feedback/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, Feedbacks[], AverageScore, LastFeedbackAt; methods Submit/Respond
│   │   │   ├── mid_contract.go  # Frequency/monthly; due date; mandatory; reminders
│   │   │   ├── score.go  # Quality/Communication/Timeliness/Professionalism; comments; submittedBy/At
│   │   │   ├── errors.go  # ErrNotDue, ErrAlreadySubmitted
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # FeedbackSubmitted/FeedbackResponded
│   │   │
│   │   ├── workroom/
│   │   │   ├── entity.go  # Aggregate: ID, ContractID, Tasks[], Notes[], Files[], Messages[]; methods CreateTask/Complete/AddNote/ShareFile
│   │   │   ├── task.go  # Title/Desc; assignedTo/createdBy; due/completed; status; priority
│   │   │   ├── note.go  # Markdown content; created/updated; type; tags
│   │   │   ├── shared_file.go  # FileName/URL/Size; uploadedBy/At; type; access permissions
│   │   │   ├── errors.go  # ErrTaskNotFound, ErrUnauthorized
│   │   │   ├── repository.go  # Interface
│   │   │   └── events.go  # TaskCreated/TaskCompleted/NoteAdded/FileShared
│   │   │
│   │   └── direct_contract/
│   │       ├── entity.go  # Aggregate: ID, ClientID, FreelancerID, InviteToken, Terms, Status(Pending/Accepted/Rejected/Expired), InvitedAt/ExpiresAt, AcceptedAt/RejectedAt; methods Accept/Reject/Expire
│   │       ├── invitation.go  # Token gen; send invitation; validate; expiry(30d)
│   │       ├── acceptance.go  # Review/Negotiate/Accept&Activate/RejectWithReason
│   │       ├── errors.go  # ErrInviteExpired, ErrAlreadyAccepted, ErrInvalidToken
│   │       ├── repository.go  # Interface
│   │       └── events.go  # Invited/Accepted/Rejected/Activated/Expired
│   │
│   # ========================================================================================
│   # 📋 APPLICATION LAYER - Use Cases & Policies
│   # ========================================================================================
│   ├── application/
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # 🔐 AUTHORIZATION
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── authz/
│   │   │   ├── policies.go  # 🆕 RBAC policies: canCreate/View/Update/Terminate/Approve/Release/Resolve/AccessDiary/Sign/ViewFinancials...
│   │   │   └── guards.go  # 🆕 Enforcement helpers: ensure participant/client/freelancer/admin; bilateral approval; active/no-disputes
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # 📬 EVENT HANDLERS (Consumers)
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   ├── eventhandler/
│   │   │   ├── _common/idempotency.go  # 🆕 store last-seen event version; event_id+causation_id; TTL 7d
│   │   │   ├── proposal_handler.go  # proposal.accepted→CreateFromProposal; invite.accepted→CreateDirect; negotiation.completed→FinalizeTerms
│   │   │   ├── payment_handler.go  # payment.processed→Record/UpdateBudget; payment.failed→Notify/Retry; refund→UpdateFinancials
│   │   │   ├── escrow_handler.go  # escrow.released→CompleteMilestone; funded→ActivateContract; refund.processed→TerminationRefund
│   │   │   ├── dispute_handler.go  # dispute.opened→PauseContract/Hold; resolved→Resume/ReleaseHold; closed→UpdateStatus
│   │   │   ├── admin_handler.go  # feature_flag.updated→Refresh; config.updated→Reload; moderation.actioned→HandleAction
│   │   │   ├── financial_risk_handler.go  # risk.alert→PlaceHold; chargeback.created→PlaceHold; chargeback.resolved→Release/Update
│   │   │   ├── contract_status_handler.go  # internal hold placed/released→update reads/notify/trigger payments
│   │   │   ├── message_handler.go  # notification_delivered→UpdateSLATimers; message.read→Comm metrics
│   │   │   ├── user_handler.go  # user.suspended→Freeze; banned→Terminate; reactivated→Resume
│   │   │   └── job_handler.go  # job.closed→MarkContractsComplete; job.cancelled→HandleCancellation
│   │   │
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   # 🎯 DOMAIN SERVICES (Use Cases)
│   │   # ──────────────────────────────────────────────────────────────────────────────────
│   │   #
│   │   # ===== CORE CONTRACT SERVICES =====
│   │   ├── contract/
│   │   │   ├── service.go       # Contract lifecycle management (Activate, Pause/Resume, Complete, Terminate), outbox emit, read-model updates
│   │   │   ├── commands.go      # CreateContract, UpdateContract, PauseContract, ResumeContract, CompleteContract, TerminateContract
│   │   │   ├── queries.go       # GetContract, ListContracts, FilterContracts, GetContractStats (pagination/sorting)
│   │   │   ├── validators.go    # ValidateTerms, ValidateAmount, ValidateDates, ValidateParticipants, status transition guards
│   │   │   └── dto.go           # ContractDTO, CreateContractRequest, UpdateContractRequest, mappers
│   │   #
│   │   # ===== STATEMENT OF WORK (SOW) =====
│   │   ├── sow/
│   │   │   ├── service.go       # SOW creation/versioning/approval workflow; link to contract; enforce approval rules
│   │   │   ├── commands.go      # CreateSOW, UpdateSOW, ApproveSOW, RejectSOW, CreateSOWVersion (diffs & metadata)
│   │   │   ├── queries.go       # GetSOW, ListSOWVersions, GetLatestSOW (history access)
│   │   │   ├── validators.go    # ValidateSOWCompleteness, ValidateScope, version constraints & ACL checks
│   │   │   └── dto.go           # SOWDTO, CreateSOWRequest, SOWVersionDTO
│   │   #
│   │   # ===== FINANCIAL HOLDS (RISK MGMT) =====
│   │   ├── financial_hold/
│   │   │   ├── service.go       # Place/Release/Extend/Expire holds; escrow & disputes hooks; DLQ-safe updates
│   │   │   ├── commands.go      # PlaceHold, ReleaseHold, ExtendHold, ExpireHold
│   │   │   ├── queries.go       # GetHold, ListHolds, GetActiveHolds, GetHoldsExpiringSoon
│   │   │   ├── validators.go    # ValidateHoldAmount, ValidateHoldReason, authz checks
│   │   │   └── dto.go           # FinancialHoldDTO, PlaceHoldRequest, ReleaseHoldRequest
│   │   #
│   │   # ===== MILESTONES =====
│   │   ├── milestone/
│   │   │   ├── service.go       # Milestone CRUD/submission/approval/rejection/release/dispute; auto-release integration
│   │   │   ├── commands.go      # CreateMilestone, SubmitMilestone, ApproveMilestone, RejectMilestone, ReleaseMilestone, DisputeMilestone
│   │   │   ├── queries.go       # GetMilestone, ListMilestones, GetPendingMilestones, GetOverdueMilestones, GetMilestoneStats
│   │   │   ├── validators.go    # ValidateMilestoneAmount, ValidateDueDate, ValidateDeliverables, ordering rules
│   │   │   └── dto.go           # MilestoneDTO, CreateMilestoneRequest, SubmitMilestoneRequest
│   │   #
│   │   # ===== DELIVERABLES =====
│   │   ├── deliverable/
│   │   │   ├── service.go       # Deliverable submit/review/accept/reject/revision loop; file links; audit
│   │   │   ├── commands.go      # SubmitDeliverable, ReviewDeliverable, AcceptDeliverable, RejectDeliverable, RequestRevision
│   │   │   ├── queries.go       # GetDeliverable, ListDeliverables, GetPendingReview
│   │   │   ├── validators.go    # ValidateDeliverableFiles, ValidateDescription, state checks
│   │   │   └── dto.go           # DeliverableDTO, SubmitDeliverableRequest, ReviewDeliverableRequest
│   │   #
│   │   # ===== TIME TRACKING =====
│   │   ├── timesheet/
│   │   │   ├── service.go       # Weekly timesheets: add entries, submit, approve/reject/dispute; totals calc
│   │   │   ├── commands.go      # CreateTimesheet, AddTimeEntry, SubmitTimesheet, ApproveTimesheet, RejectTimesheet, DisputeTimesheet
│   │   │   ├── queries.go       # GetTimesheet, ListTimesheets, GetTimesheetSummary, GetPendingApproval, GetDisputedTimesheets
│   │   │   ├── validators.go    # ValidateTimeEntries, ValidateHours (24h/day cap, duplicates), ValidateWeek
│   │   │   └── dto.go           # TimesheetDTO, TimeEntryDTO, SubmitTimesheetRequest
│   │   ├── workdiary/
│   │   │   ├── service.go       # Work diary tracking & screenshots; manual time; privacy controls; retention rules
│   │   │   ├── commands.go      # RecordWorkDiaryEntry, AddManualTime, DeleteScreenshot, UpdatePrivacySettings
│   │   │   ├── queries.go       # GetWorkDiary, ListWorkDiaryEntries, GetActivitySummary, GetScreenshots
│   │   │   ├── validators.go    # ValidateTrackingData (intervals, activity), ValidatePrivacySettings
│   │   │   └── dto.go           # WorkDiaryDTO, WorkDiaryEntryDTO, PrivacySettingsDTO
│   │   #
│   │   # ===== TEMPLATES =====
│   │   ├── template/
│   │   │   ├── service.go       # Template CRUD/customization/publish; usage counts
│   │   │   ├── commands.go      # CreateTemplate, UpdateTemplate, PublishTemplate, UnpublishTemplate, UseTemplate, CloneTemplate
│   │   │   ├── queries.go       # GetTemplate, ListTemplates, SearchTemplates, GetPublicTemplates, GetPopularTemplates
│   │   │   ├── validators.go    # ValidateTemplateCompleteness, ValidateClauses, required fields
│   │   │   └── dto.go           # TemplateDTO, CreateTemplateRequest, UseTemplateRequest
│   │   #
│   │   # ===== CONTRACT CHANGES =====
│   │   ├── amendment/
│   │   │   ├── service.go       # Propose/approve/reject/apply; bilateral approval; version control
│   │   │   ├── commands.go      # ProposeAmendment, ApproveAmendment, RejectAmendment, ApplyAmendment, WithdrawAmendment
│   │   │   ├── queries.go       # GetAmendment, ListAmendments, GetPendingAmendments, GetAmendmentHistory
│   │   │   ├── validators.go    # ValidateAmendmentChanges, ValidateBilateralApproval, timeouts
│   │   │   └── dto.go           # AmendmentDTO, ProposeAmendmentRequest, ApproveAmendmentRequest
│   │   ├── termination/
│   │   │   ├── service.go       # Initiate/approve/complete/dispute terminations; settlement calc
│   │   │   ├── commands.go      # InitiateTermination, ApproveTermination, CompleteTermination, DisputeTermination, CalculateSettlement
│   │   │   ├── queries.go       # GetTermination, ListTerminations, GetTerminationsByType
│   │   │   ├── validators.go    # ValidateTerminationReason, ValidateNoticePeriod, ValidateSettlement
│   │   │   └── dto.go           # TerminationDTO, InitiateTerminationRequest, SettlementDTO
│   │   #
│   │   # ===== DISPUTES =====
│   │   ├── dispute/
│   │   │   ├── service.go       # Open/escalate/resolve/close disputes; evidence handling; SLAs
│   │   │   ├── commands.go      # OpenDispute, SubmitEvidence, EscalateDispute, ResolveDispute, CloseDispute
│   │   │   ├── queries.go       # GetDispute, ListDisputes, GetOpenDisputes, GetEscalatedDisputes, GetDisputeStats
│   │   │   ├── validators.go    # ValidateDisputeReason, ValidateEvidence, state guards
│   │   │   └── dto.go           # DisputeDTO, OpenDisputeRequest, SubmitEvidenceRequest
│   │   #
│   │   # ===== DOCUMENTS & SIGNATURES =====
│   │   ├── attachment/
│   │   │   ├── service.go       # Upload/version/delete attachments; storage links; audit/versioning
│   │   │   ├── commands.go      # UploadAttachment, DeleteAttachment, CreateAttachmentVersion
│   │   │   ├── queries.go       # GetAttachment, ListAttachments, GetAttachmentVersions
│   │   │   ├── validators.go    # ValidateFileType, ValidateFileSize, ValidateAttachmentType
│   │   │   └── dto.go           # AttachmentDTO, UploadAttachmentRequest
│   │   ├── signature/
│   │   │   ├── service.go       # Request/sign/verify/decline; sequential flow; audit trail; integrity hash
│   │   │   ├── commands.go      # RequestSignature, SignContract, DeclineSignature, VerifySignature, ResendSignatureRequest
│   │   │   ├── queries.go       # GetSignature, GetSigningStatus, ListSignatures, GetPendingSignatures
│   │   │   ├── validators.go    # ValidateSigner, ValidateSigningOrder, ValidateSignatureData
│   │   │   └── dto.go           # SignatureDTO, RequestSignatureRequest, SignContractRequest
│   │   #
│   │   # ===== FINANCIAL =====
│   │   ├── budget/
│   │   │   ├── service.go       # Track/update budget; alerts & forecasting; thresholds logic
│   │   │   ├── commands.go      # UpdateBudget, SetThreshold, RecordSpending, AdjustBudget
│   │   │   ├── queries.go       # GetBudget, GetBudgetForecast, GetSpendingTrend, GetOverBudgetContracts
│   │   │   ├── validators.go    # ValidateBudgetAmount, ValidateThresholds
│   │   │   └── dto.go           # BudgetDTO, SetThresholdRequest, BudgetForecastDTO
│   │   ├── invoice/
│   │   │   ├── service.go       # Generate/send/mark paid/cancel invoices; payment tracking
│   │   │   ├── commands.go      # GenerateInvoice, SendInvoice, MarkInvoicePaid, CancelInvoice, RegenerateInvoice
│   │   │   ├── queries.go       # GetInvoice, ListInvoices, GetOverdueInvoices, GetInvoiceStats
│   │   │   ├── validators.go    # ValidateLineItems, ValidateTaxCalculation
│   │   │   └── dto.go           # InvoiceDTO, GenerateInvoiceRequest, InvoiceLineItemDTO
│   │   #
│   │   # ===== COMPLIANCE & LEGAL =====
│   │   ├── sla/
│   │   │   ├── service.go       # Define/record metrics; detect breaches; penalties/rewards; reports
│   │   │   ├── commands.go      # DefineSLA, RecordMetric, DetectBreach, ApplyPenalty, GrantReward, GenerateSLAReport
│   │   │   ├── queries.go       # GetSLA, GetSLAReport, ListBreaches, GetSLACompliance
│   │   │   ├── validators.go    # ValidateSLAMetrics, ValidateThresholds
│   │   │   └── dto.go           # SLADTO, DefineSLARequest, SLAMetricDTO, SLABreachDTO
│   │   ├── agency/
│   │   │   ├── service.go       # Team contracts (multi-freelancer); roles/permissions; payment split
│   │   │   ├── commands.go      # AddAgencyMember, RemoveAgencyMember, UpdateMemberRole, UpdateBillingSplit, CalculatePayments
│   │   │   ├── queries.go       # GetAgencyContract, ListAgencyMembers, GetBillingSplit
│   │   │   ├── validators.go    # ValidateMemberPermissions, ValidateSplitRules
│   │   │   └── dto.go           # AgencyContractDTO, AgencyMemberDTO, BillingSplitDTO
│   │   ├── compliance/
│   │   │   ├── service.go       # Track requirements, verify, run audits, expire; jurisdiction rules
│   │   │   ├── commands.go      # AddRequirement, VerifyCompliance, RunAudit, UpdateCompliance, ExpireCompliance
│   │   │   ├── queries.go       # GetCompliance, ListRequirements, GetAuditReport, GetExpiringCompliance
│   │   │   ├── validators.go    # ValidateRequirement, ValidateAuditCriteria
│   │   │   └── dto.go           # ComplianceDTO, ComplianceRequirementDTO, AuditDTO
│   │   #
│   │   # ===== PERFORMANCE =====
│   │   ├── performance/
│   │   │   ├── service.go       # Define KPIs, record metrics, calc scores/trends; benchmarking & reports
│   │   │   ├── commands.go      # DefineKPI, RecordPerformance, CalculateScore, GeneratePerformanceReport, UpdateBenchmark
│   │   │   ├── queries.go       # GetPerformance, GetKPIScores, GetTrends, CompareToBenchmark, GetPerformanceReport
│   │   │   ├── validators.go    # ValidateKPIDefinition, ValidateWeights, ValidateScores
│   │   │   └── dto.go           # PerformanceDTO, KPIDTO, PerformanceScoreDTO, TrendDTO
│   │   #
│   │   # ===== NEGOTIATION =====
│   │   ├── negotiation/
│   │   │   ├── service.go       # Offers/counter-offers; accept/reject; rounds & timeouts; audit history
│   │   │   ├── commands.go      # MakeOffer, MakeCounterOffer, AcceptOffer, RejectOffer, WithdrawOffer, ExpireNegotiation
│   │   │   ├── queries.go       # GetNegotiation, ListOffers, GetActiveNegotiation, GetNegotiationHistory
│   │   │   ├── validators.go    # ValidateOffer, ValidateOfferTerms, ValidateCounterOffer
│   │   │   └── dto.go           # NegotiationDTO, OfferDTO, CounterOfferDTO
│   │   #
│   │   # ===== RENEWAL & EXTENSION =====
│   │   ├── renewal/
│   │   │   ├── service.go       # Request/approve/reject renewals; extensions; auto-renew rules & notifications
│   │   │   ├── commands.go      # RequestRenewal, ApproveRenewal, RejectRenewal, ExtendContract, EnableAutoRenewal, DisableAutoRenewal
│   │   │   ├── queries.go       # GetRenewal, ListRenewals, GetRenewalEligibility, GetExpiringContracts
│   │   │   ├── validators.go    # ValidateRenewalTerms, ValidateExtension
│   │   │   └── dto.go           # RenewalDTO, RequestRenewalRequest, ExtensionDTO
│   │   #
│   │   # ===== IP RIGHTS =====
│   │   ├── ip_rights/
│   │   │   ├── service.go       # Assign IP, license grants, term updates; protect & audit
│   │   │   ├── commands.go      # AssignIP, GrantLicense, UpdateIPTerms, TransferRights
│   │   │   ├── queries.go       # GetIPRights, GetLicenseTerms, ListIPAssignments
│   │   │   ├── validators.go    # ValidateIPTerms, ValidateLicense, ValidateAssignment
│   │   │   └── dto.go           # IPRightsDTO, AssignIPRequest, LicenseDTO
│   │   #
│   │   # ===== NDA =====
│   │   ├── nda/
│   │   │   ├── service.go       # Sign/expire NDAs; breach reporting & enforcement; penalties
│   │   │   ├── commands.go      # SignNDA, ReportBreach, EnforceNDA, ExpireNDA
│   │   │   ├── queries.go       # GetNDA, ListNDAs, ListBreaches, GetActiveNDAs
│   │   │   ├── validators.go    # ValidateNDATerms, ValidateBreach
│   │   │   └── dto.go           # NDADTO, SignNDARequest, NDABreachDTO
│   │   #
│   │   # ===== REPORTING & COLLABORATION =====
│   │   ├── report/
│   │   │   ├── service.go       # Generate/schedule/export reports; analytics aggregation
│   │   │   ├── commands.go      # GenerateReport, ScheduleReport, ExportReport, CancelScheduledReport
│   │   │   ├── queries.go       # GetReport, ListReports, GetAnalytics, GetScheduledReports
│   │   │   ├── validators.go    # ValidateReportType, ValidateReportParameters
│   │   │   └── dto.go           # ReportDTO, GenerateReportRequest, AnalyticsDTO
│   │   ├── feedback/
│   │   │   ├── service.go       # Mid-contract feedback; reminders; average scores
│   │   │   ├── commands.go      # SubmitFeedback, RespondToFeedback, RequestFeedback
│   │   │   ├── queries.go       # GetFeedback, ListFeedbacks, GetFeedbackSummary, GetDueFeedbacks, GetAverageScores
│   │   │   ├── validators.go    # ValidateFeedbackScores, ValidateFeedbackTiming
│   │   │   └── dto.go           # FeedbackDTO, SubmitFeedbackRequest, FeedbackScoreDTO
│   │   ├── workroom/
│   │   │   ├── service.go       # Collaborative tasks/notes/files/messages; permissions
│   │   │   ├── commands.go      # CreateTask, UpdateTask, CompleteTask, DeleteTask, AddNote, UpdateNote, DeleteNote, ShareFile, DeleteFile
│   │   │   ├── queries.go       # GetTask, ListTasks, GetNote, ListNotes, GetFile, ListFiles, GetWorkroomActivity
│   │   │   ├── validators.go    # ValidateTask, ValidateFileType, ValidateAccess
│   │   │   └── dto.go           # WorkroomTaskDTO, WorkroomNoteDTO, WorkroomFileDTO
│   │   #
│   │   # ===== DIRECT CONTRACT =====
│   │   └── direct_contract/
│   │       ├── service.go       # Invite/accept/reject; activate direct contracts; token flow
│   │       ├── commands.go      # InviteFreelancer, AcceptInvite, RejectInvite, ActivateDirectContract, CancelInvite, ResendInvite
│   │       ├── queries.go       # GetInvite, ListInvites, GetInviteStatus, GetPendingInvites
│   │       ├── validators.go    # ValidateInviteTerms, ValidateToken, ValidateFreelancer
│   │       └── dto.go           # DirectContractDTO, InviteFreelancerRequest, AcceptInviteRequest
│   │
│   # ========================================================================================
│   # 🌐 INTERFACES LAYER - HTTP (v1), OpenAPI, Middleware
│   # ========================================================================================
│   └── interfaces/
│       └── http/
│           ├── v1/                                   # 🚦 Versioned API surface (mounted at /v1)
│           │   │
│           │   # ──────────────────────────────────────────────────────────────────────────
│           │   # 🎯 HANDLERS (Use-Case Endpoints)
│           │   # ──────────────────────────────────────────────────────────────────────────
│           │   ├── handlers/
│           │   │   # ===== CORE CONTRACT HANDLERS =====
│           │   │   ├── contract_handler.go           # Contract CRUD & lifecycle; authz; RFC7807 mapping; metrics/tracing
│           │   │
│           │   │   # ===== STATEMENT OF WORK (SOW) =====
│           │   │   ├── sow_handler.go                # SOW create/version/approve/reject; version rules; Problem Details
│           │   │
│           │   │   # ===== FINANCIAL HOLDS (RISK MGMT) =====
│           │   │   ├── financial_hold_handler.go     # Place/extend/release holds; permissions; notifications
│           │   │
│           │   │   # ===== MILESTONES =====
│           │   │   ├── milestone_handler.go          # Create/submit/approve/reject/release/dispute; auto-release integration
│           │   │
│           │   │   # ===== DELIVERABLES =====
│           │   │   ├── deliverable_handler.go        # Submit/review/accept/reject/revision; file links; auditing
│           │   │
│           │   │   # ===== TIME TRACKING =====
│           │   │   ├── timesheet_handler.go          # Add entries/submit/approve/reject/dispute; week checks; totals
│           │   │   ├── workdiary_handler.go          # Record entries/screenshots; privacy settings; retention-safe deletes
│           │   │
│           │   │   # ===== TEMPLATES =====
│           │   │   ├── template_handler.go           # CRUD/publish/use/clone; search & pagination; cache hints
│           │   │
│           │   │   # ===== CONTRACT CHANGES =====
│           │   │   ├── amendment_handler.go          # Propose/approve/reject/apply; bilateral approval; timeouts
│           │   │   ├── termination_handler.go        # Initiate/approve/complete/dispute; settlement calc; idempotency keys
│           │   │
│           │   │   # ===== DISPUTES =====
│           │   │   ├── dispute_handler.go            # Open/escalate/resolve/close; evidence uploads; SLA timers
│           │   │
│           │   │   # ===== DOCUMENTS & SIGNATURES =====
│           │   │   ├── attachment_handler.go         # Upload/delete/version; presigned URLs; AV scan results
│           │   │   ├── signature_handler.go          # Request/sign/decline/verify; sequential flow; integrity hash
│           │   │
│           │   │   # ===== FINANCIAL =====
│           │   │   ├── budget_handler.go             # Update thresholds/forecast; burn-rate output; alert triggers
│           │   │   ├── invoice_handler.go            # Generate/send/mark-paid/cancel; pay links; state transitions
│           │   │
│           │   │   # ===== COMPLIANCE & LEGAL =====
│           │   │   ├── sla_handler.go                # Define/record/detect-breach/report; penalties/rewards
│           │   │   ├── agency_handler.go             # Team membership/roles/billing split; permissions matrix
│           │   │   ├── compliance_handler.go         # Requirements verify/audits/run/expire; jurisdiction rules
│           │   │
│           │   │   # ===== PERFORMANCE =====
│           │   │   ├── performance_handler.go        # KPIs define/record/scores/trends; benchmark comparisons
│           │   │
│           │   │   # ===== NEGOTIATION =====
│           │   │   ├── negotiation_handler.go        # Offers/counter-offers/accept/reject/expire; round limits
│           │   │
│           │   │   # ===== RENEWAL & EXTENSION =====
│           │   │   ├── renewal_handler.go            # Request/approve/reject/extend; auto-renew opt-in/out notices
│           │   │
│           │   │   # ===== IP RIGHTS =====
│           │   │   ├── ip_rights_handler.go          # Assign/licensing/updates; protection flags; auditability
│           │   │
│           │   │   # ===== NDA =====
│           │   │   ├── nda_handler.go                # Sign/expire/report breach/enforce; penalties mapping
│           │   │
│           │   │   # ===== REPORTING & COLLABORATION =====
│           │   │   ├── report_handler.go             # Generate/schedule/export; async job handoff; list pagination
│           │   │   ├── feedback_handler.go           # Submit/respond/request; averages & due filters
│           │   │   ├── workroom_handler.go           # Tasks/notes/files/messages CRUD; access checks
│           │   │
│           │   │   # ===== DIRECT CONTRACT =====
│           │   │   └── direct_contract_handler.go    # Invite/accept/reject/activate; token validation; expiry handling
│           │   │
│           │   # ──────────────────────────────────────────────────────────────────────────
│           │   # 🛣️ ROUTES (HTTP Surface)
│           │   # ──────────────────────────────────────────────────────────────────────────
│           │   ├── routes/
│           │   │   # ===== CORE CONTRACT ROUTES =====
│           │   │   ├── contract_routes.go            # /v1/contracts (POST,GET); /v1/contracts/:id (GET,PATCH,DELETE); filters,pagination,ETag
│           │   │
│           │   │   # ===== STATEMENT OF WORK (SOW) =====
│           │   │   ├── sow_routes.go                 # /v1/contracts/:id/sow (POST,GET); /v1/contracts/:id/sow/versions (POST,GET); approve/reject
│           │   │
│           │   │   # ===== FINANCIAL HOLDS (RISK MGMT) =====
│           │   │   ├── financial_hold_routes.go      # /v1/contracts/:id/holds (POST,GET); /v1/holds/:holdId/release (POST); extend/expire
│           │   │
│           │   │   # ===== MILESTONES =====
│           │   │   ├── milestone_routes.go           # /v1/contracts/:id/milestones (POST,GET); /v1/milestones/:id/{submit|approve|reject|release} (POST)
│           │   │
│           │   │   # ===== DELIVERABLES =====
│           │   │   ├── deliverable_routes.go         # /v1/milestones/:id/deliverables (POST,GET); review/accept/reject/revision (POST)
│           │   │
│           │   │   # ===== TIME TRACKING =====
│           │   │   ├── timesheet_routes.go           # /v1/contracts/:id/timesheets (POST,GET); /v1/timesheets/:id/{approve|reject|dispute} (POST)
│           │   │   ├── workdiary_routes.go           # /v1/contracts/:id/workdiary (GET); /v1/workdiary/privacy (PATCH); screenshots CRUD
│           │   │
│           │   │   # ===== TEMPLATES =====
│           │   │   ├── template_routes.go            # /v1/templates (POST,GET); /v1/templates/:id (GET,PATCH,DELETE); /:id/{use|publish|unpublish} (POST)
│           │   │
│           │   │   # ===== CONTRACT CHANGES =====
│           │   │   ├── amendment_routes.go           # /v1/contracts/:id/amendments (POST,GET); /v1/amendments/:id/{approve|reject|apply} (POST)
│           │   │   ├── termination_routes.go         # /v1/contracts/:id/terminate (POST); /v1/terminations/:id/settlement (GET); approve/complete/dispute (POST)
│           │   │
│           │   │   # ===== DISPUTES =====
│           │   │   ├── dispute_routes.go             # /v1/contracts/:id/disputes (POST,GET); /v1/disputes/:id/evidence (POST,GET); {escalate|resolve|close} (POST)
│           │   │
│           │   │   # ===== DOCUMENTS & SIGNATURES =====
│           │   │   ├── attachment_routes.go          # /v1/contracts/:id/attachments (POST,GET); /v1/attachments/:id/versions (POST,GET); delete (DELETE)
│           │   │   ├── signature_routes.go           # /v1/contracts/:id/signatures (POST,GET); /v1/signatures/:id/{sign|decline|verify} (POST)
│           │   │
│           │   │   # ===== FINANCIAL =====
│           │   │   ├── budget_routes.go              # /v1/contracts/:id/budget (GET,PATCH); /v1/budget/forecast (GET); thresholds (PATCH)
│           │   │   ├── invoice_routes.go             # /v1/contracts/:id/invoices (POST,GET); /v1/invoices/:id/pay (POST); mark-paid|cancel (POST)
│           │   │
│           │   │   # ===== COMPLIANCE & LEGAL =====
│           │   │   ├── sla_routes.go                 # /v1/contracts/:id/sla (POST,GET); metrics/record (POST); /v1/sla/breaches (GET)
│           │   │   ├── agency_routes.go              # /v1/contracts/:id/agency (GET,PATCH); /v1/agency/members (POST,DELETE,PATCH); billing-split (PATCH)
│           │   │   ├── compliance_routes.go          # /v1/contracts/:id/compliance (GET,PATCH); /v1/compliance/audits (POST,GET); verify/expire (POST)
│           │   │
│           │   │   # ===== PERFORMANCE =====
│           │   │   ├── performance_routes.go         # /v1/contracts/:id/performance (GET); /v1/performance/kpis (POST,GET,PATCH); scores/trends (GET)
│           │   │
│           │   │   # ===== NEGOTIATION =====
│           │   │   ├── negotiation_routes.go         # /v1/contracts/:id/negotiation (GET); /v1/negotiation/offers (POST,GET); accept|reject|withdraw (POST)
│           │   │
│           │   │   # ===== RENEWAL & EXTENSION =====
│           │   │   ├── renewal_routes.go             # /v1/contracts/:id/renewal (POST,GET); /v1/renewal/extend (POST); auto-renew enable/disable (POST)
│           │   │
│           │   │   # ===== IP RIGHTS =====
│           │   │   ├── ip_rights_routes.go           # /v1/contracts/:id/ip-rights (GET,PATCH); /v1/ip-rights/assign (POST); license ops (POST)
│           │   │
│           │   │   # ===== NDA =====
│           │   │   ├── nda_routes.go                 # /v1/contracts/:id/nda (GET,POST); /v1/nda/sign (POST); breach/enforce (POST)
│           │   │
│           │   │   # ===== REPORTING & COLLABORATION =====
│           │   │   ├── report_routes.go              # /v1/contracts/:id/reports (POST,GET); /v1/reports/analytics (GET); export/schedule (POST)
│           │   │   ├── feedback_routes.go            # /v1/contracts/:id/feedback (POST,GET); /v1/feedback/submit (POST); respond (POST)
│           │   │   ├── workroom_routes.go            # /v1/contracts/:id/workroom (GET); /v1/workroom/{tasks|notes|files} (CRUD); activity (GET)
│           │   │
│           │   │   # ===== DIRECT CONTRACT =====
│           │   │   └── direct_contract_routes.go     # /v1/direct-contracts (POST,GET); /v1/direct-contracts/:token/accept (POST); reject|resend (POST)
│           │   │
│           │   # ===== OPENAPI =====
│           │   └── openapi/
│           │       ├── openapi.yaml  # OpenAPI 3.0: endpoints, schemas, auth, errors
│           │       └── generator.go  # Generate /swagger + /openapi.json (dev); optional request validation in debug
│           │
│           # ===== PRESENTERS & HTTP UTILITIES =====
│           ├── presenters/
│           │   └── errors.go  # 🆕 RFC7807 Problem+JSON: type/title/status/detail/instance/extensions
│           ├── etag.go  # 🆕 ETag middleware for cacheable GET: generate hash; If-None-Match; 304; feature-flagged
│           └── middleware/
│               ├── auth.go  # JWT via pkg/auth: verify token; extract claims; inject user ctx
│               ├── rbac.go  # ♻️ role/permission checks; align with app/authz policies; deny unauthorized
│               ├── cors.go  # CORS: origins, methods, preflight
│               ├── rate_limit.go  # per-user/endpoint limits; Redis counters; 429 on exceed
│               ├── idempotency.go  # Idempotency-Key(UUIDv4); store req/resp in Redis (24h); dedupe
│               ├── requestid.go  # 🆕 X-Request-ID (ULID); add header; inject to logs
│               ├── logging.go  # 🆕 structured req/resp logs; duration; correlation id
│               ├── recovery.go  # 🆕 panic recovery; stack trace; 500; keep server alive
│               └── tracing.go  # 🆕 distributed tracing: spans; inject context; attrs(method/path/status)
│
# ========================================================================================
# 🗂️ DATABASE MIGRATIONS
# ========================================================================================
├── db/
│   └── migrations/  # 🔄 symlink/mirror of internal migrations; make migrate-{up,down}
│       ├── contracts/  # core contract schema
│       ├── milestones/  # milestones & deliverables
│       ├── financial/  # holds, budgets, invoices
│       ├── compliance/  # SLAs, NDAs, IP rights
│       ├── amendments/  # amendments, terminations
│       └── collaboration/  # workroom, templates, agency
│
# ========================================================================================
# 📚 DOCUMENTATION
# ========================================================================================
├── docs/
│   ├── README.md  # overview & quick start
│   ├── API.md  # high-level API
│   ├── API_VERSIONING.md  # 🆕 versioning & deprecation policy (semver; breaking vs non-breaking; 6mo deprecation; Sunset headers)
│   ├── EVENTS.md  # event schemas & topics (lifecycle/milestones/financial/timesheet/disputes/compliance/collaboration)
│   ├── ARCHITECTURE.md  # layers; DDD; event-driven; CQRS
│   ├── MIGRATIONS.md  # 🆕 forward-only; version tracking; safety checks; zero-downtime; concurrent indexes
│   ├── SCHEMA.md  # 🆕 ERD; table schemas (33+); indexes; constraints; retention
│   ├── RUNBOOK.md  # 🆕 deploy/rollback/incidents/troubleshooting/health checks
│   ├── CACHING.md  # 🆕 keys; TTLs; warming; SWR; event invalidation; singleflight
│   ├── DATA_RETENTION.md  # 🛡️🆕 contracts 7y; screenshots 30d; dispute evidence 3y; audit logs 2y; soft-deletes 90d; cleanup jobs
│   ├── ERASURE.md  # 🛡️🆕 GDPR/CCPA erasure: hooks; anonymization; legal holds; playbooks; audit trail
│   ├── SLOS.md  # 📏🆕 SLO targets (latency, success, lags); alerting
│   ├── OUTBOX.md  # 🆕 exactly-once-ish; tx writes; retries; DLQ; cleanup; metrics
│   ├── RELEASE_CHECKLIST.md  # 🆕 openapi-diff; schema-diff; migration plan; flags; env vars; rollback; dashboards
│   ├── NAMING.md  # ♻️ packages snake_case; endpoints kebab-case; tables snake_case; protobuf PascalCase; events dot.case
│   ├── OPENAPI.md  # 🆕 regenerate spec; SDK gen (Go/TS); publish; breaking change detection
│   ├── contract-lifecycle.md  # flows & state diagrams
│   ├── milestone-system.md  # milestone/deliverable flows; auto-release(3d); payouts
│   ├── timesheet-workdiary.md  # weekly timesheets; approvals; work diary; screenshots; manual time; disputes
│   ├── dispute-resolution.md  # open/evidence/escalate/resolve; holds
│   ├── amendment-termination.md  # bilateral approvals; versioning; types; notice; settlements
│   ├── sla-performance.md  # metrics; breaches; penalties/rewards; KPIs; trends; benchmarks
│   ├── agency-contracts.md  # teams; permissions; splits; consolidated invoicing
│   ├── direct-contracts.md  # invites; token acceptance; pre-negotiation; bypass proposals
│   └── compliance-legal.md  # NDAs; IP rights; compliance reqs; audits; jurisdictions
│
# ========================================================================================
# 🧪 TESTS
# ========================================================================================
├── tests/
│   ├── unit/
│   │   ├── domain/  # 33 aggregates
│   │   │   ├── contract_test.go
│   │   │   ├── milestone_test.go
│   │   │   ├── timesheet_test.go
│   │   │   └── ... (30 more)
│   │   ├── application/  # 33 services
│   │   │   ├── contract_service_test.go
│   │   │   ├── milestone_service_test.go
│   │   │   └── ... (31 more)
│   │   └── infrastructure/
│   │       ├── cache_test.go
│   │       ├── outbox_test.go
│   │       └── clients_test.go
│   ├── integration/
│   │   ├── repository/  # real DB
│   │   │   ├── contract_repository_test.go
│   │   │   ├── milestone_repository_test.go
│   │   │   └── ... (31 more)
│   │   ├── eventhandler/
│   │   │   ├── proposal_handler_test.go
│   │   │   ├── payment_handler_test.go
│   │   │   └── ... (8 more)
│   │   └── external_clients/
│   │       ├── users_client_test.go
│   │       ├── financial_client_test.go
│   │       └── ... (5 more)
│   ├── reliability/
│   │   ├── projections_replay_test.go  # 🆕 replay all contract events; verify read models; ordering; idempotency
│   │   └── outbox_dispatcher_test.go  # 🆕 at-least-once; idempotency; retries; DLQ
│   ├── property/
│   │   ├── contract_property_test.go  # 🆕 invariants: amount>0; start<end; valid transitions
│   │   ├── milestone_property_test.go  # 🆕 sum≤contract; ordered due dates; no release before approval
│   │   └── budget_property_test.go  # 🆕 spent+remaining=total; spent≥0; ascending thresholds
│   └── e2e/
│       ├── contract_lifecycle_test.go  # create→activate→work→complete; dispute path; amendment path
│       ├── dispute_resolution_test.go  # open→evidence→escalate→resolve
│       └── chaos_event_delivery_test.go  # 🆕 dup/out-of-order/delayed events; resilience
│
# ========================================================================================
# 🔧 SCRIPTS
# ========================================================================================
├── scripts/
│   ├── migrate.sh  # 🆕 migrate up/down/version; prod safety checks
│   ├── openapi-diff.sh  # 🆕 compare spec vs main; detect breakage; CI fail; changelog
│   ├── schema-diff.sh  # 🆕 pg_dump; compare; drift detection; CI guard
│   ├── generate-sdks.sh  # 🆕 gen Go SDK (sdk/go) & TS SDK (sdk/ts); publish
│   └── sbom-sign.sh  # 🆕 SBOM build + cosign sign/verify
│
# ========================================================================================
# 🚀 CI/CD
# ========================================================================================
├── .github/
│   └── workflows/
│       ├── ci.yml  # tests (unit/integration), linters, build, push image
│       ├── cd.yml  # deploy dev/staging/prod; smoke; rollback on fail
│       ├── contract-ci.yml  # 🆕 openapi-diff; event schema checks (buf breaking); fail on breaking
│       ├── security.yml  # 🆕 golangci-lint(security); govulncheck; trivy; cosign verify
│       └── load-tests.yml  # 🆕 k6/Gatling; /healthz & hot paths; P99 alerts
│
# ========================================================================================
# 📦 SDK GENERATION
# ========================================================================================
├── sdk/
│   ├── go/
│   │   ├── client.go  # main client
│   │   ├── contracts.go  # contract ops
│   │   ├── milestones.go  # milestone ops
│   │   └── ... (31 more)
│   └── ts/
│       ├── client.ts  # main client
│       ├── contracts.ts  # contract ops
│       ├── milestones.ts  # milestone ops
│       └── ... (31 more)
│
# ========================================================================================
# 🔧 CONFIG / ROOT FILES
# ========================================================================================
├── .golangci.yml  # 🆕 linters: enable security; no unused; complexity limits; CI parity
├── CODEOWNERS  # 🆕 ownership: domain & infra dirs; docs
├── go.mod  # deps: pkg/auth, platform-shared(telemetry/outbox), contracts/events (protobuf)
├── go.sum  # checksums
├── .env.example  # DATABASE_URL, REDIS_URL, KAFKA_BROKERS, MINIO_ENDPOINT, KEYCLOAK_URL, FEATURE_FLAGS_*
├── Makefile  # build/test/lint/fmt/migrate/openapi/sdks/docker/deploy
├── Dockerfile  # multi-stage; minimal runtime; non-root; healthcheck
├── .dockerignore  # docker ignores
├── .gitignore  # git ignores
└── README.md  # overview; quick start; architecture; dev & deploy guides; contributing


```

---
## 📦 **financial-be (Financial Management Service) - REFACTORED**

```

apps/be/financial-be/
│
├── cmd/
│   ├── api/
│   │   └── main.go                           # 📝 API entrypoint - initializes Gin, Dapr, Postgres (uses platform-shared/logging, internal/config)
│   │                                         # ♻️ Serve versioned routes (/v1/*) & enable ETag middleware when flag is on
│   │                                         # NO background jobs here - delegated to cmd/worker
│   │
│   └── worker/
│       └── main.go                           # 🆕 Worker entrypoint - background jobs, inbox/outbox, leader election
│                                             # Runs: payout batches, payment schedules, reminders, reconciliation, outbox dispatcher
│                                             # Uses: infrastructure/coordination for leader election & distributed locks
│
├── internal/
│   │
│   # ──────────────────────────────────────────────────────────────────────────────────
│   # 🔧 CONFIGURATION (Load First)
│   # ──────────────────────────────────────────────────────────────────────────────────
│   ├── config/
│   │   ├── schema.go                         # ♻️ Typed Config struct (App, Server, Postgres, Kafka, Redis, Auth, Stripe/PayPal keys, FX providers, risk thresholds)
│   │   │                                     # Group flags under Config.FeatureFlags
│   │   ├── feature_flags.go                  # 🆕 Central toggles (enable_ledger_audit, enable_etag, enable_fx_auto_update, enable_risk_holds, etc.)
│   │   ├── loader.go                         # Config loader using Viper (CLI → ENV → file → defaults)
│   │   └── docs/
│   │       └── CONFIGURATION.md              # Configuration documentation
│   │
│   # ──────────────────────────────────────────────────────────────────────────────────
│   # 🧩 DEPENDENCY INJECTION CONTAINER
│   # ──────────────────────────────────────────────────────────────────────────────────
│   ├── ioc/
│   │   ├── container.go                      # 🆕 DI graph: constructs DB/Redis/Kafka clients, repositories, services, handlers, schedulers
│   │   │                                     # Wire outbox, coordination, observers based on env
│   │   └── wiring.go                         # 🆕 Env-driven wiring & feature flags: selects implementations (local vs cloud)
│   │                                         # Wire outbox, coordination, observers based on env
│   │
│   # ──────────────────────────────────────────────────────────────────────────────────
│   # 🔧 INFRASTRUCTURE LAYER (Load Second)
│   # ──────────────────────────────────────────────────────────────────────────────────
│   ├── infrastructure/
│   │   │
│   │   # ===== DISTRIBUTED COORDINATION =====
│   │   ├── coordination/
│   │   │   ├── leader_election.go            # 🆕 Single-active cron/worker (ensures only one worker processes jobs)
│   │   │   │                                 # Used by: payout batches, payment schedules, reminders, reconciliation
│   │   │   └── distlock.go                   # 🆕 Redis/PG advisory locks for cron tasks (prevents duplicate execution)
│   │   │
│   │   # ===== PERSISTENCE LAYER =====
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection (pooling, tracing, retries)
│   │   │       ├── transaction.go            # ♻️ Transaction helpers (unit-of-work, retry on serialization)
│   │   │       │                             # Ensure outbox write is in same DB tx
│   │   │       ├── migrations.go             # 📝 Auto-migration logic with version tracking
│   │   │       ├── version.go                # Schema version tracking
│   │   │       ├── safety.go                 # Pre-migration safety checks
│   │   │       │
│   │   │       ├── migrations/               # 🆕 SQL-first, forward-only migrations (prod)
│   │   │       │   ├── wallet/               # 🆕 Per-schema migration folders
│   │   │       │   ├── payment/              # 🆕
│   │   │       │   ├── escrow/               # 🆕
│   │   │       │   ├── tax/                  # 🆕
│   │   │       │   └── risk/                 # 🆕
│   │   │       │
│   │   │       # ===== CORE WALLET & LEDGER REPOSITORIES =====
│   │   │       ├── wallet_repository.go      # WalletRepository implementation (balances, reserves, CRUD)
│   │   │       ├── transaction_repository.go # TransactionRepository implementation (journal linkage, persist, lookup, list)
│   │   │       ├── ledger_journal_repository.go # LedgerJournalRepository (append-only journal, audit trail, hash chain verification)
│   │   │       │
│   │   │       # ===== PAYMENT & ESCROW REPOSITORIES =====
│   │   │       ├── payment_repository.go     # PaymentRepository implementation (provider refs, states, search)
│   │   │       ├── escrow_repository.go      # EscrowRepository implementation (holds, releases, refunds, pro-rata)
│   │   │       ├── payout_repository.go      # PayoutRepository implementation (queue, batch, process, cancel)
│   │   │       │
│   │   │       # ===== INVOICE & FEE REPOSITORIES =====
│   │   │       ├── invoice_repository.go     # InvoiceRepository implementation (lines, taxes, generation, send, mark paid)
│   │   │       ├── fee_repository.go         # FeeRepository implementation (fee rows, audits, calculations)
│   │   │       ├── fee_rules_repository.go   # FeeRulesRepository implementation (v2 rulesets: tiers, coupons, locale exceptions)
│   │   │       │
│   │   │       # ===== REFUND & DISPUTE REPOSITORIES =====
│   │   │       ├── refund_repository.go      # RefundRepository implementation (process, cancel, partial refunds, states, audit)
│   │   │       ├── dispute_payment_repository.go # DisputePaymentRepository implementation (chargebacks, representment, resolution)
│   │   │       │
│   │   │       # ===== TAX & FX REPOSITORIES =====
│   │   │       ├── tax_repository.go         # TaxRepository implementation (records, forms, filing, VAT/GST, 1099-K)
│   │   │       ├── fx_repository.go          # FXRepository implementation (time-based rates, quote/settlement, rounding rules)
│   │   │       │
│   │   │       # ===== RISK MANAGEMENT REPOSITORIES =====
│   │   │       ├── risk_repository.go        # RiskRepository implementation (holds, reserves, chargeback workflows, negative balances)
│   │   │       │
│   │   │       # ===== PROTECTION & FEE UPDATE REPOSITORIES =====
│   │   │       ├── protection_plan_repository.go  # ProtectionPlanRepository implementation (plans, claims, eligibility, payouts)
│   │   │       ├── fee_update_repository.go       # FeeUpdateRepository implementation (fee versions, migrations, impact calculations)
│   │   │       │
│   │   │       # ===== INTERNATIONAL PAYMENTS REPOSITORY =====
│   │   │       ├── international_payment_repository.go # InternationalPaymentRepository (routing, compliance checks, local rails, FX adjustments)
│   │   │       │
│   │   │       # ===== BONUS & EXPENSE REPOSITORIES =====
│   │   │       ├── bonus_repository.go            # BonusRepository implementation (award, pay, reverse, conditions)
│   │   │       ├── expense_repository.go          # ExpenseRepository implementation (submit, approve, reimburse, receipts)
│   │   │       │
│   │   │       # ===== PAYMENT SCHEDULE & REMINDER REPOSITORIES =====
│   │   │       ├── payment_schedule_repository.go # PaymentScheduleRepository (schedules, frequency, due windows, automation)
│   │   │       ├── reminder_repository.go         # ReminderRepository implementation (triggers, escalation, dunning, templates)
│   │   │       │
│   │   │       # ===== INSURANCE & TAX FORM REPOSITORIES =====
│   │   │       ├── insurance_repository.go        # InsuranceRepository implementation (policies, claims, coverage, providers)
│   │   │       ├── tax_form_repository.go         # TaxFormRepository implementation (W9, 1099, VAT returns, validation, reporting)
│   │   │       │
│   │   │       # ===== PAYROLL & CURRENCY REPOSITORIES =====
│   │   │       ├── payroll_repository.go          # PayrollRepository implementation (process, withholding, pay periods, reporting)
│   │   │       ├── currency_repository.go         # CurrencyRepository implementation (preferences, rate locks, conversions)
│   │   │       │
│   │   │       # ===== BANK ACCOUNT REPOSITORY =====
│   │   │       ├── bank_account_repository.go     # 🆕 BankAccountRepository implementation (CRUD, verification, default payout method)
│   │   │       │
│   │   │       # ===== OUTBOX REPOSITORY =====
│   │   │       └── outbox_store.go           # 🆕 Outbox repository (tx-coupled outbox table access for exactly-once publishing)
│   │   │
│   │   # ===== CACHING LAYER =====
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go             # Redis connection (pooling, timeouts)
│   │   │       ├── keys.go                   # 🆕 Canonical cache keys & TTLs (no magic strings)
│   │   │       ├── singleflight.go           # 🆕 Stampede protection for hot keys (wallet balances, FX rates)
│   │   │       ├── invalidation_rules.go     # 🆕 Map events → keys to drop (documented alongside)
│   │   │       │
│   │   │       # ===== DOMAIN-SPECIFIC CACHES =====
│   │   │       ├── wallet_cache.go           # Wallet balance caching (hot path reads)
│   │   │       ├── rate_cache.go             # Exchange rate caching (FX provider TTL)
│   │   │       ├── risk_cache.go             # Holds/reserve snapshots for quick checks
│   │   │       └── schedule_cache.go         # NextDue caches for payment schedules (cron efficiency)
│   │   │
│   │   # ===== EXTERNAL SERVICE CLIENTS =====
│   │   ├── clients/
│   │   │   ├── users/                        # Users-BE client (tax profiles, KYC/KYB statuses)
│   │   │   │   └── client.go                 # Fetch tax profiles & KYC statuses; circuit breaker/retry
│   │   │   │
│   │   │   ├── coupons/                      # Coupon/Promo validation backend (if separate service)
│   │   │   │   └── client.go                 # Validate & redeem promo codes; idempotent redemption
│   │   │   │
│   │   │   └── contracts/                    # 🆕 Contracts-BE client (contract status, milestones)
│   │   │       └── client.go                 # Query contract status for escrow releases & payment schedules
│   │   │
│   │   # ===== MESSAGING LAYER =====
│   │   ├── messaging/
│   │   │   ├── kafka/
│   │   │   │   ├── consumer.go               # Kafka consumer (uses platform-shared/inbox, DLQ, retries)
│   │   │   │   ├── producer.go               # Kafka producer (uses platform-shared/outbox, exactly-once semantics)
│   │   │   │   ├── topics.go                 # ♻️ Topics: payment.*, escrow.*, payout.*, fee.*, tax.*, fx.*, risk.*, bonus.*, expense.*, schedules.*, reminders.*
│   │   │   │   │                             # Thin re-export from contracts/events with per-context keys (financial.*)
│   │   │   │   └── scram.go                  # SCRAM authentication (SASL/SCRAM-SHA-512)
│   │   │   │
│   │   │   ├── outbox/                       # 🆕 Exactly-once-ish publisher
│   │   │   │   ├── dispatcher.go             # 🆕 Reads outbox → Kafka (with retries/DLQ, leader-elected)
│   │   │   │   └── metrics.go                # 🆕 Publish lag, failures, retries (Prometheus metrics)
│   │   │   │
│   │   │   └── bootstrap.go                  # 🆕 Consumer/producer wiring (initialize subscriptions, configure handlers)
│   │   │
│   │   # ===== PAYMENT GATEWAY INTEGRATIONS =====
│   │   ├── payment_gateway/
│   │   │   ├── stripe/
│   │   │   │   ├── client.go                 # Stripe API client (auth, retries, idempotency keys)
│   │   │   │   ├── webhook_handler.go        # Stripe webhook handler (payment_intent.succeeded, charge.failed)
│   │   │   │   │                             # 🆕 Includes signature verification middleware
│   │   │   │   └── mapper.go                 # Stripe → internal event mapper (normalize payloads)
│   │   │   │
│   │   │   ├── paypal/
│   │   │   │   ├── client.go                 # PayPal API client (orders/captures/refunds)
│   │   │   │   ├── webhook_handler.go        # PayPal webhook handler (capture completed/denied)
│   │   │   │   │                             # 🆕 Includes signature verification middleware
│   │   │   │   └── mapper.go                 # PayPal → internal event mapper
│   │   │   │
│   │   │   └── factory.go                    # Payment gateway factory (select provider by method/region, fallback/retry logic)
│   │   │
│   │   # ===== PDF GENERATION =====
│   │   ├── pdf/
│   │   │   └── generator.go                  # PDF invoice generation (wkhtmltopdf/Chrome headless wrapper)
│   │   │
│   │   # ===== SCHEDULER (BACKGROUND JOBS) =====
│   │   ├── scheduler/
│   │   │   ├── cron.go                       # ♻️ Cron registry: schedules, idempotency guards, safe shutdown
│   │   │   │                                 # Wrap each task with distlock + jitter
│   │   │   └── tasks/
│   │   │       ├── task_guard.go             # 🆕 Idempotency tokens + last-run watermark
│   │   │       ├── payout_batch_task.go      # 🆕 Batch payout processing (grouped by method/currency, runs in worker)
│   │   │       ├── payment_schedule_task.go  # 🆕 Run due payment schedules (auto-payment triggers)
│   │   │       ├── reminder_task.go          # 🆕 Send reminders (invoice due, tax form due, escalation)
│   │   │       ├── reconciliation_task.go    # 🆕 Daily reconciliation (bank vs journal, generates reports)
│   │   │       ├── fx_rate_update_task.go    # 🆕 Update FX rates from providers (periodic refresh)
│   │   │       └── risk_review_task.go       # 🆕 Review holds/reserves for auto-release (SLA windows)
│   │   │
│   │   # ===== OBSERVABILITY =====
│   │   ├── observability/                    # 🆕 Consolidated telemetry utilities
│   │   │   ├── metrics.go                    # 🆕 RED/USE metrics helpers (payment latency, escrow holds, payout success rate)
│   │   │   ├── tracing.go                    # 🆕 Common attrs (hashed user_id, org_id, event_id, payment_id)
│   │   │   └── slo_monitor.go                # 🆕 SLO monitoring (payment P99, payout batch lag, reconciliation delay)
│   │   │
│   │   # ===== HTTP UTILITIES =====
│   │   └── http/
│   │       ├── idempotency_adapter.go        # Bind platform-shared idempotency to Gin (critical for payments!)
│   │       └── etag.go                       # 🆕 Middleware-level ETag (pairs with utils/etag, optional feature flag)
│   │
│   # ──────────────────────────────────────────────────────────────────────────────────
│   # 🏛️ DOMAIN LAYER (Business Logic & Entities - Load Third)
│   # ──────────────────────────────────────────────────────────────────────────────────
│   ├── domain/
│   │   │
│   │   # ===== CORE WALLET & BALANCE =====
│   │   ├── wallet/
│   │   │   ├── entity.go                     # User wallet (UserID, Balance, Currency, Status, CreatedAt)
│   │   │   ├── balance.go                    # Balance tracking (Available, Pending, Reserved) with atomic adjustments
│   │   │   ├── currency.go                   # Multi-currency support (USD, EUR, GBP, etc.) and default currency rules
│   │   │   ├── errors.go                     # Wallet errors (InsufficientFunds, WalletNotFound, WalletClosed)
│   │   │   ├── repository.go                 # WalletRepository interface (CRUD, Reserve, Release, GetBalance)
│   │   │   └── events.go                     # Domain events: WalletCreated, WalletUpdated, WalletClosed, FundsReserved, FundsReleased
│   │   │
│   │   # ===== TRANSACTION & LEDGER =====
│   │   ├── transaction/
│   │   │   ├── entity.go                     # Financial transactions (ID, WalletID, Amount, Type, Status, Reference, CreatedAt)
│   │   │   ├── enums.go                      # Type (Deposit, Withdrawal, Transfer, Payment, Refund), Status (Pending, Completed, Failed), Category
│   │   │   ├── ledger.go                     # Double-entry ledger (debit/credit accounting) + balancing checks
│   │   │   ├── errors.go                     # Transaction errors (TransactionFailed, DuplicateTransaction, ImbalancedEntry)
│   │   │   ├── repository.go                 # TransactionRepository interface (persist, lookup, list, reconcile)
│   │   │   └── events.go                     # Domain events: TransactionInitiated, TransactionPosted, TransactionFailed, TransactionReversed
│   │   │
│   │   ├── ledger_journal/                   # Wallet ledger (double-entry) - immutable journal & audit
│   │   │   ├── entity.go                     # JournalEntry (ID, DebitAccount, CreditAccount, Amount, Currency, EffectiveAt, Hash, PrevHash)
│   │   │   ├── transfer.go                   # Transfers between accounts (builds 2-sided entries; validates debits=credits)
│   │   │   ├── adjustment.go                 # Manual adjustments with maker/checker approvals & notes
│   │   │   ├── audit.go                      # Audit trail (entry proofs, tamper-evidence via hash chain)
│   │   │   ├── errors.go                     # Ledger errors (Imbalance, ImmutableViolation, HashMismatch)
│   │   │   ├── repository.go                 # LedgerJournalRepository interface (append, list, audit)
│   │   │   └── events.go                     # Domain events: JournalEntryRecorded, JournalAdjustmentApproved, JournalAdjustmentRejected
│   │   │
│   │   # ===== PAYMENT PROCESSING =====
│   │   ├── payment/
│   │   │   ├── entity.go                     # Payment records (ID, Amount, Currency, PayerID, PayeeID, Method, Status, ProviderRef)
│   │   │   ├── payment_method.go             # Payment methods (CreditCard, PayPal, BankTransfer, Wallet)
│   │   │   ├── method_profile.go             # Payout preferences per method (limits, cutoffs, fees, destination IDs)
│   │   │   ├── onboarding.go                 # Onboarding KYC/KYB statuses per method/provider (required docs, states)
│   │   │   ├── gateway.go                    # Payment gateway abstraction (Stripe, PayPal interface) & retries
│   │   │   ├── errors.go                     # Payment errors (PaymentFailed, InvalidCard, GatewayTimeout, InsufficientFunds)
│   │   │   ├── repository.go                 # PaymentRepository interface (persist, update, search, reconcile)
│   │   │   └── events.go                     # Domain events: PaymentAuthorized, PaymentCaptured, PaymentFailed, PaymentRefundInitiated
│   │   │
│   │   # ===== ESCROW MANAGEMENT =====
│   │   ├── escrow/
│   │   │   ├── entity.go                     # Escrow accounts (ID, ContractID, Amount, Status, HeldAt, ReleasedAt)
│   │   │   ├── hold.go                       # Fund holds (amount reserved until milestone completion or dispute close)
│   │   │   ├── release.go                    # Fund release conditions (milestone approved, dispute resolved, partial releases)
│   │   │   ├── pro_rata.go                   # Pro-rata partial releases by milestone & refunds (distribution calc)
│   │   │   ├── errors.go                     # Escrow errors (EscrowNotFound, InsufficientEscrow, EscrowLocked)
│   │   │   ├── repository.go                 # EscrowRepository interface (create, hold, release, refund, partial release)
│   │   │   └── events.go                     # Domain events: EscrowFunded, EscrowPartiallyReleased, EscrowReleased, EscrowRefunded
│   │   │
│   │   # ===== PAYOUT PROCESSING =====
│   │   ├── payout/
│   │   │   ├── entity.go                     # Payout requests (ID, UserID, Amount, Method, Status, RequestedAt, ProcessedAt)
│   │   │   ├── method.go                     # Payout methods (BankTransfer, PayPal, Payoneer, Wire) & metadata
│   │   │   ├── schedule.go                   # Payout scheduling (instant, daily, weekly, monthly) with cut-off logic
│   │   │   ├── errors.go                     # Payout errors (PayoutFailed, BelowMinimum, MethodUnavailable, InsufficientBalance)
│   │   │   ├── repository.go                 # PayoutRepository interface (queue, batch, update, cancel, process)
│   │   │   └── events.go                     # Domain events: PayoutRequested, PayoutScheduled, PayoutProcessed, PayoutFailed
│   │   │
│   │   # ===== INVOICE MANAGEMENT =====
│   │   ├── invoice/
│   │   │   ├── entity.go                     # Invoice generation (ID, ContractID, Number, Amount, DueDate, Status, PaidAt)
│   │   │   ├── line_item.go                  # Invoice line items (description, quantity, price, total, tax)
│   │   │   ├── tax.go                        # Tax calculations (VAT, sales tax by jurisdiction & reverse charge flags)
│   │   │   ├── errors.go                     # Invoice errors (InvoiceNotFound, AlreadyPaid, InvalidTotals, InvalidTax)
│   │   │   ├── repository.go                 # InvoiceRepository interface (create, send, update, lookup, mark paid, cancel)
│   │   │   └── events.go                     # Domain events: InvoiceIssued, InvoiceUpdated, InvoicePaid, InvoiceOverdue, InvoiceCanceled
│   │   │
│   │   # ===== PLATFORM FEES =====
│   │   ├── fee/
│   │   │   ├── entity.go                     # Platform fees (ID, TransactionID, Amount, Type, Rate, Currency)
│   │   │   ├── calculator.go                 # Fee calculation rules (percentage, tiered rates, min/max caps)
│   │   │   ├── tier.go                       # Fee tiers (based on subscription, volume, user type)
│   │   │   │
│   │   │   ├── v2/                           # Fee engine v2 (advanced rules)
│   │   │   │   ├── rules.go                  # Rules model (volume discounts, locale overrides, coupons, experiments)
│   │   │   │   ├── coupon.go                 # Coupon/promo code entity & redemption logic (single-use, stacking)
│   │   │   │   ├── country_exceptions.go     # Country/regional exceptions (local regulation overrides)
│   │   │   │   ├── errors.go                 # Fee v2 errors (InvalidCoupon, IneligibleTier, ExpiredRule)
│   │   │   │   └── repository.go             # FeeRulesRepository interface (CRUD for rule sets)
│   │   │   │
│   │   │   ├── errors.go                     # Fee errors (FeeConfigMissing, NegativeFee, InvalidTier)
│   │   │   ├── repository.go                 # FeeRepository interface (persist & audit fee applications)
│   │   │   └── events.go                     # Domain events: FeeCalculated, FeeAdjusted, CouponApplied, CouponRevoked
│   │   │
│   │   # ===== REFUND PROCESSING =====
│   │   ├── refund/
│   │   │   ├── entity.go                     # Refund processing (ID, PaymentID, Amount, Reason, Status, ProcessedAt)
│   │   │   ├── policy.go                     # Refund policies (full, partial, time limits, dispute linkage)
│   │   │   ├── pro_rata.go                   # Pro-rata refunds aligned to milestones (safe allocation)
│   │   │   ├── errors.go                     # Refund errors (RefundNotAllowed, RefundExpired, OverRefund, InvalidAmount)
│   │   │   ├── repository.go                 # RefundRepository interface (queue, process, audit, cancel)
│   │   │   └── events.go                     # Domain events: RefundRequested, RefundApproved, RefundDeclined, RefundProcessed
│   │   │
│   │   # ===== PAYMENT DISPUTES =====
│   │   ├── dispute_payment/
│   │   │   ├── entity.go                     # Payment disputes (ID, PaymentID, Reason, FiledBy, Status, Resolution)
│   │   │   ├── chargeback.go                 # Chargeback handling (card network flows, representment)
│   │   │   ├── errors.go                     # Dispute errors (DisputeNotFound, ChargebackInvalidState, AlreadyResolved)
│   │   │   ├── repository.go                 # DisputePaymentRepository interface (open, update, resolve, representment)
│   │   │   └── events.go                     # Domain events: PaymentDisputeOpened, ChargebackRecorded, PaymentDisputeResolved
│   │   │
│   │   # ===== TAX MANAGEMENT =====
│   │   ├── tax/
│   │   │   ├── entity.go                     # Tax records (ID, UserID, Year, Type, Amount, Status, Jurisdiction)
│   │   │   ├── vat_gst.go                    # VAT/GST per locale & reverse-charge logic (EU/UK/CA, etc.)
│   │   │   ├── reverse_charge.go             # Reverse-charge mechanics & validations (B2B cross-border)
│   │   │   ├── forms_1099k.go                # 1099-K forms generation & thresholds (US)
│   │   │   ├── profile_link.go               # Link to Users-BE tax profile (TIN/VAT ID; sync contract-level data)
│   │   │   ├── form.go                       # Tax forms (W9, 1099, VAT returns; states & versions)
│   │   │   ├── errors.go                     # Tax errors (TaxProfileMissing, InvalidVATID, FormInvalid, ThresholdNotMet)
│   │   │   ├── repository.go                 # TaxRepository interface (create, update, file, generate forms)
│   │   │   └── events.go                     # Domain events: TaxRecordCreated, TaxFormGenerated, TaxProfileLinked, VATReverseChargeApplied
│   │   │
│   │   # ===== FOREIGN EXCHANGE (FX) =====
│   │   ├── fx/
│   │   │   ├── rate.go                       # FXRate (Base, Quote, Rate, EffectiveFrom, Provider, Precision, Source)
│   │   │   ├── quote_settlement.go           # Quote vs. settlement currency handling (timing, spreads)
│   │   │   ├── rounding.go                   # Rounding rules per currency/amount type (bankers, up/down)
│   │   │   ├── errors.go                     # FX errors (RateNotFound, StaleRate, PrecisionExceeded, InvalidPair)
│   │   │   ├── repository.go                 # FXRepository interface (upsert rate, query time-bounded, get effective rate)
│   │   │   └── events.go                     # Domain events: FXRateUpdated, FXQuoteCreated, FXSettlementCalculated
│   │   │
│   │   # ===== RISK MANAGEMENT =====
│   │   ├── risk/
│   │   │   ├── reserve.go                    # Reserve balances & rolling reserve schedules (percent/time windows)
│   │   │   ├── hold_workflow.go              # Holds/auto-release based on risk signals & time (SLA windows)
│   │   │   ├── chargeback_workflow.go        # Chargeback workflows & recovery (fees, evidence, outcomes)
│   │   │   ├── negative_balance.go           # Negative balance handling & collections (repayment plans)
│   │   │   ├── errors.go                     # Risk errors (ReserveConflict, HoldNotFound, NegativeBalanceLimit)
│   │   │   ├── repository.go                 # RiskRepository interface (holds, reserves, risk metrics, negative balance tracking)
│   │   │   └── events.go                     # Domain events: ReserveSet, ReserveUpdated, RiskHoldPlaced, RiskHoldReleased, NegativeBalanceCreated
│   │   │
│   │   # ===== PAYMENT PROTECTION PLANS =====
│   │   ├── protection_plan/
│   │   │   ├── entity.go                     # ProtectionPlan aggregate (ContractID, Type, CoverageAmount, Premium, Status, Period)
│   │   │   ├── type.go                       # Plan types (HourlyProtection, FixedPriceProtection) & defaults
│   │   │   ├── claim.go                      # Protection claims (filing, review, payout; integrates with risk & escrow)
│   │   │   ├── eligibility.go                # Eligibility checks (contract type, verification, history)
│   │   │   ├── errors.go                     # Protection errors (NotEligible, ClaimDenied, CoverageExceeded, AlreadyClaimed)
│   │   │   ├── repository.go                 # ProtectionPlanRepository interface (plan lifecycle, claims, eligibility checks)
│   │   │   └── events.go                     # Domain events: ProtectionApplied, ClaimFiled, ClaimApproved, ClaimPaid
│   │   │
│   │   # ===== FEE UPDATES & MIGRATIONS =====
│   │   ├── fee_update/
│   │   │   ├── entity.go                     # FeeUpdate aggregate (Version, EffectiveDate, Rules[], Impact, Status)
│   │   │   ├── version.go                    # Fee versions (e.g., flat 10% for freelancers, 3–5% for clients)
│   │   │   ├── impact.go                     # Impact calculations (per segment; notifications, proration)
│   │   │   ├── migration.go                  # Fee migration logic (from old to new rule sets; rollback)
│   │   │   ├── errors.go                     # FeeUpdate errors (VersionConflict, InvalidRate, MigrationFailed, ActiveVersionExists)
│   │   │   ├── repository.go                 # FeeUpdateRepository interface (publish, activate, archive, rollback)
│   │   │   └── events.go                     # Domain events: FeeUpdated, ImpactCalculated, Migrated, RolledBack
│   │   │
│   │   # ===== INTERNATIONAL PAYMENTS =====
│   │   ├── international_payment/
│   │   │   ├── entity.go                     # InternationalPayment (TransactionID, LocalCurrency, ComplianceChecks, Routing)
│   │   │   ├── compliance.go                 # OFAC/AML/Sanctions checks; links user/tax profile for KYC/KYB
│   │   │   ├── local_method.go               # Local payout methods (SEPA, ACH, local banks) + routing preferences
│   │   │   ├── fee_adjustment.go             # Cross-border fee adjustments (FX spreads, local rails fees)
│   │   │   ├── errors.go                     # InternationalPayment errors (ComplianceFailed, MethodUnavailable, RoutingError)
│   │   │   ├── repository.go                 # InternationalPaymentRepository interface (route, track status, compliance checks)
│   │   │   └── events.go                     # Domain events: InternationalPaymentInitiated, InternationalPaymentCompliant, InternationalPaymentProcessed
│   │   │
│   │   # ===== BONUS PAYMENTS =====
│   │   ├── bonus/
│   │   │   ├── entity.go                     # Bonus aggregate (ContractID, Amount, Reason, AwardedBy, AwardedAt, Status)
│   │   │   ├── type.go                       # Bonus types (Performance, Completion, Referral) & default flows
│   │   │   ├── condition.go                  # Bonus conditions (milestone met, on-time delivery, KPI triggers)
│   │   │   ├── errors.go                     # Bonus errors (BonusNotEligible, AlreadyAwarded, InvalidReason, AmountExceedsLimit)
│   │   │   ├── repository.go                 # BonusRepository interface (award, reverse, list, get by contract)
│   │   │   └── events.go                     # Domain events: BonusRequested, BonusAwarded, BonusRejected, BonusPaid
│   │   │
│   │   # ===== EXPENSE REIMBURSEMENTS =====
│   │   ├── expense/
│   │   │   ├── entity.go                     # Expense aggregate (ContractID, ItemID, Amount, Description, SubmittedBy, Status, SubmittedAt)
│   │   │   ├── receipt.go                    # Receipt attachments and verification (OCR hooks)
│   │   │   ├── approval.go                   # Approval workflow (client approval required; levels)
│   │   │   ├── policy.go                     # Reimbursement policies (caps, eligible categories, per diem)
│   │   │   ├── errors.go                     # Expense errors (ExpenseNotApproved, OverCap, InvalidReceipt, MissingDocumentation)
│   │   │   ├── repository.go                 # ExpenseRepository interface (submit, approve, reimburse, reject)
│   │   │   └── events.go                     # Domain events: ExpenseSubmitted, ExpenseApproved, ExpenseRejected, ExpenseReimbursed
│   │   │
│   │   # ===== PAYMENT SCHEDULES =====
│   │   ├── payment_schedule/
│   │   │   ├── entity.go                     # PaymentSchedule aggregate (ContractID, Frequency, Amount, NextDue, Status)
│   │   │   ├── frequency.go                  # Payment frequency (Weekly, BiWeekly, Monthly) & cut-off rules
│   │   │   ├── adjustment.go                 # Schedule adjustments (prorating, skips, deferrals)
│   │   │   ├── automation.go                 # Auto-payment triggers (cron, Dapr bindings)
│   │   │   ├── errors.go                     # Schedule errors (ScheduleConflict, PaymentMissed, DisabledSchedule, InsufficientFunds)
│   │   │   ├── repository.go                 # PaymentScheduleRepository interface (create, update, list, due, process)
│   │   │   └── events.go                     # Domain events: ScheduleCreated, ScheduleUpdated, PaymentDue, PaymentProcessed
│   │   │
│   │   # ===== AUTOMATED REMINDERS =====
│   │   ├── reminder/
│   │   │   ├── entity.go                     # Reminder aggregate (ContractID, Type, Trigger, SentAt, NextAt, Attempts)
│   │   │   ├── type.go                       # Reminder types (InvoicePay, PayoutDue, TaxFormDue, ScheduleRun)
│   │   │   ├── template.go                   # Reminder templates (customizable messages; locale-aware)
│   │   │   ├── escalation.go                 # Escalation paths (repeated reminders, penalties, dunning)
│   │   │   ├── errors.go                     # Reminder errors (ReminderNotSent, ChannelUnavailable, ThrottleLimitReached)
│   │   │   ├── repository.go                 # ReminderRepository interface (queue, mark sent, escalate, list)
│   │   │   └── events.go                     # Domain events: ReminderTriggered, ReminderSent, ReminderEscalated
│   │   │
│   │   # ===== INSURANCE & PROTECTION =====
│   │   ├── insurance/
│   │   │   ├── entity.go                     # Insurance aggregate (ContractID, PolicyID, Coverage, Premium, Status)
│   │   │   ├── coverage.go                   # Coverage details (payment protection, liability, limits)
│   │   │   ├── claim.go                      # Insurance claims (filing, status, payout; external provider codes)
│   │   │   ├── provider.go                   # Integration with insurance providers (quote/bind/claim)
│   │   │   ├── errors.go                     # Insurance errors (PolicyNotActive, ClaimDenied, ProviderUnavailable, CoverageExceeded)
│   │   │   ├── repository.go                 # InsuranceRepository interface (create policy, claim lifecycle, provider integration)
│   │   │   └── events.go                     # Domain events: InsuranceApplied, ClaimFiled, ClaimApproved, ClaimPaid
│   │   │
│   │   # ===== TAX FORMS =====
│   │   ├── tax_form/
│   │   │   ├── entity.go                     # TaxForm aggregate (ContractID, FormType, SubmittedBy, Status, SubmittedAt)
│   │   │   ├── type.go                       # Form types (W9, 1099, VAT ID, 1099-K support docs)
│   │   │   ├── validation.go                 # Form validation and verification (format & identity checks)
│   │   │   ├── reporting.go                  # Tax reporting integration (export to tax authorities/providers)
│   │   │   ├── errors.go                     # Tax form errors (FormInvalid, NotSubmitted, VerificationFailed, IncompleteData)
│   │   │   ├── repository.go                 # TaxFormRepository interface (submit, verify, fetch, list, report)
│   │   │   └── events.go                     # Domain events: TaxFormSubmitted, TaxFormVerified, TaxReportGenerated
│   │   │
│   │   # ===== PAYROLL INTEGRATION =====
│   │   ├── payroll/
│   │   │   ├── entity.go                     # Payroll aggregate (ContractID, WorkerStatus, PayPeriods[], Taxes, Deductions)
│   │   │   ├── status.go                     # Worker classification (Freelancer, Employee) + validation hooks
│   │   │   ├── pay_period.go                 # Pay periods (bi-weekly, monthly; with deductions & employer taxes)
│   │   │   ├── tax.go                        # Tax withholding and reporting (per jurisdiction)
│   │   │   ├── errors.go                     # Payroll errors (ClassificationMismatch, PeriodClosed, InsufficientDeductions)
│   │   │   ├── repository.go                 # PayrollRepository interface (process, post, report, withhold)
│   │   │   └── events.go                     # Domain events: PayrollProcessed, TaxWithheld, PayPeriodClosed
│   │   │
│   │   # ===== CURRENCY MANAGEMENT =====
│   │   ├── currency/
│   │   │   ├── entity.go                     # Currency aggregate (ContractID, PreferredCurrency, ExchangeRates, LockedAt)
│   │   │   ├── conversion.go                 # Real-time conversion logic (preferred currency, rounding)
│   │   │   ├── rate_lock.go                  # Rate locking for payments (validity windows, expirations)
│   │   │   ├── errors.go                     # Currency errors (ConversionFailed, RateNotAvailable, LockExpired, InvalidCurrency)
│   │   │   ├── repository.go                 # CurrencyRepository interface (locks, preferences, histories, conversions)
│   │   │   └── events.go                     # Domain events: CurrencyChanged, RateLocked, Converted
│   │   │
│   │   # ===== BANK ACCOUNT MANAGEMENT =====
│   │   └── bank_account/                     # 🆕 Bank account domain (for exposed endpoints)
│   │       ├── entity.go                     # BankAccount aggregate (UserID, AccountNumber, RoutingNumber, BankName, IsDefault, Status, VerifiedAt)
│   │       ├── verifier.go                   # Bank account verification (micro-deposits, Plaid integration)
│   │       ├── errors.go                     # BankAccount errors (AccountNotFound, VerificationFailed, InvalidRouting, AlreadyExists)
│   │       ├── repository.go                 # BankAccountRepository interface (CRUD, verify, set default, list by user)
│   │       └── events.go                     # Domain events: BankAccountAdded, BankAccountVerified, BankAccountSetDefault, BankAccountRemoved
│   │
│   # ──────────────────────────────────────────────────────────────────────────────────
│   # 📋 APPLICATION LAYER (Use Cases & Orchestration - Load Fourth)
│   # ──────────────────────────────────────────────────────────────────────────────────
│   ├── application/
│   │   │
│   │   # ===== EVENT HANDLERS (Inbound Events) =====
│   │   ├── eventhandler/
│   │   │   ├── _common/
│   │   │   │   └── idempotency.go            # 🆕 Store last-seen event version per aggregate (prevent duplicate processing)
│   │   │   │
│   │   │   ├── billing_handler.go            # Consumes: subscriptions-be → billing.invoice.exported (ingest AR for charging)
│   │   │   ├── admin_risk_handler.go         # Consumes: risk.* (hold/reserve/chargeback/velocity alerts → apply effects)
│   │   │   ├── admin_flags_handler.go        # Consumes: admin.feature_flag/threshold/experiment.updated (runtime behaviors)
│   │   │   ├── seats_billing_handler.go      # Consumes: seat usage/billing streams (seat changes → invoice lines)
│   │   │   ├── contract_handler.go           # 🆕 Consumes: contract.* (contract status changes → escrow releases, payment schedules)
│   │   │   └── dispute_handler.go            # 🆕 Consumes: dispute.* (dispute events → escrow holds, refund triggers)
│   │   │
│   │   # ===== AUTHORIZATION LAYER =====
│   │   ├── authz/                            # 🆕 Service-level permission checks (defense-in-depth)
│   │   │   ├── policies.go                   # 🆕 Map roles → actions (admin, wallet_owner, contract_party, finance_admin)
│   │   │   └── guards.go                     # 🆕 Helpers used by services (CanAccessWallet, CanInitiateRefund, CanViewInvoice)
│   │   │
│   │   # ===== CORE WALLET & LEDGER SERVICES =====
│   │   ├── wallet/
│   │   │   ├── service.go                    # Wallet logic (Create, Deposit, Withdraw, Transfer, GetBalance, Reserve, Release)
│   │   │   ├── commands.go                   # DepositCommand, WithdrawCommand, TransferCommand, CreateWallet, ReserveFunds, ReleaseFunds
│   │   │   ├── queries.go                    # GetBalance, GetHistory, GetTransactions (filter/pagination)
│   │   │   ├── validators.go                 # Validate currency codes, positive amounts, sufficient balance, wallet status
│   │   │   ├── dto.go                        # WalletDTO, TransactionDTO, BalanceDTO (in/out mapping contracts)
│   │   │   └── mapper.go                     # Wallet ⇄ DTO mapping helpers
│   │   │
│   │   ├── transaction/
│   │   │   ├── service.go                    # Transaction logic (Create, Reverse, Reconcile with journal)
│   │   │   ├── commands.go                   # CreateTransaction, ReverseTransaction, ReconcileTransactions
│   │   │   ├── queries.go                    # GetTransaction, ListTransactions, GetLedger (with filters)
│   │   │   ├── validators.go                 # Immutable journal checks, debits=credits, idempotency keys
│   │   │   ├── dto.go                        # TransactionDTO, LedgerDTO, ReconciliationReportDTO
│   │   │   ├── mapper.go                     # Transaction ⇄ DTO mappers
│   │   │   └── reconciliation.go             # Reconciliation logic (match payments with bank statements)
│   │   │
│   │   ├── ledger_journal/
│   │   │   ├── service.go                    # AppendEntry, Transfer, Adjust, AuditTrail orchestration
│   │   │   ├── commands.go                   # AppendJournalEntry, TransferFunds, CreateAdjustment (maker/checker)
│   │   │   ├── queries.go                    # GetEntry, ListEntries, GetAuditTrail (hash continuity)
│   │   │   ├── validators.go                 # Entry hash chain, debits=credits, approval gates
│   │   │   ├── dto.go                        # JournalEntryDTO, TransferDTO, AdjustmentDTO, AuditTrailDTO
│   │   │   └── mapper.go                     # Journal ⇄ DTO mappers
│   │   │
│   │   # ===== PAYMENT PROCESSING SERVICES =====
│   │   ├── payment/
│   │   │   ├── service.go                    # Process, Capture, Void, Refund (routes to provider processors)
│   │   │   ├── commands.go                   # ProcessPayment, CapturePayment, VoidPayment, RefundPayment
│   │   │   ├── queries.go                    # GetPayment, ListPayments (status/provider filters)
│   │   │   ├── validators.go                 # Method eligibility, onboarding status, amount/limits, risk holds
│   │   │   ├── dto.go                        # PaymentDTO, ProcessPaymentDTO, PaymentResultDTO
│   │   │   ├── mapper.go                     # Payment ⇄ DTO mappers
│   │   │   ├── stripe_processor.go           # Stripe payment processor implementation (auth/capture/refund)
│   │   │   ├── paypal_processor.go           # PayPal payment processor implementation (orders/captures/refunds)
│   │   │   └── processor_factory.go          # Payment processor factory (select by method/provider; fallback/retry)
│   │   │
│   │   # ===== ESCROW MANAGEMENT SERVICES =====
│   │   ├── escrow/
│   │   │   ├── service.go                    # Hold, Release, Refund flows (interacts with contracts-be events)
│   │   │   ├── commands.go                   # HoldEscrow, ReleaseEscrow, RefundEscrow, PartialReleaseEscrow
│   │   │   ├── queries.go                    # GetEscrow, ListEscrows, GetEscrowHistory
│   │   │   ├── validators.go                 # Coverage checks, release conditions, pro-rata math, dispute gates
│   │   │   ├── dto.go                        # EscrowDTO, HoldEscrowDTO, ReleaseEscrowDTO, PartialReleaseDTO
│   │   │   ├── mapper.go                     # Escrow ⇄ DTO mappers
│   │   │   └── pro_rata_release_manager.go   # Dispute-driven conditional releases computation engine
│   │   │
│   │   # ===== PAYOUT PROCESSING SERVICES =====
│   │   ├── payout/
│   │   │   ├── service.go                    # Request, Schedule/Batch, Process, Cancel, GetHistory
│   │   │   ├── commands.go                   # RequestPayout, ProcessPayout, CancelPayout, SchedulePayouts
│   │   │   ├── queries.go                    # GetPayouts, GetPayoutHistory (time/status filters)
│   │   │   ├── validators.go                 # Method limits, compliance holds, minimum amounts, KYC checks
│   │   │   ├── dto.go                        # PayoutDTO, RequestPayoutDTO, PayoutBatchDTO
│   │   │   ├── mapper.go                     # Payout ⇄ DTO mappers
│   │   │   └── batch_processor.go            # Batch payout processor (grouped by method/currency, runs in worker)
│   │   │
│   │   # ===== INVOICE MANAGEMENT SERVICES =====
│   │   ├── invoice/
│   │   │   ├── service.go                    # Generate, Send, MarkPaid, Cancel; PDF rendering
│   │   │   ├── commands.go                   # GenerateInvoice, SendInvoice, MarkInvoicePaid, CancelInvoice
│   │   │   ├── queries.go                    # GetInvoice, ListInvoices (filters: contract/user/status/date)
│   │   │   ├── validators.go                 # Line totals, tax rounding, currency consistency, unique number
│   │   │   ├── dto.go                        # InvoiceDTO, LineItemDTO, GenerateInvoiceDTO
│   │   │   ├── mapper.go                     # Invoice ⇄ DTO mappers
│   │   │   ├── generator.go                  # Invoice PDF generation orchestrator
│   │   │   └── tax_calculator.go             # Tax calculation (VAT, sales tax; reverse charge)
│   │   │
│   │   # ===== FEE CALCULATION SERVICES =====
│   │   ├── fee/
│   │   │   ├── service.go                    # Calculate, Apply, Waive fees to transactions/payouts
│   │   │   ├── calculator.go                 # Base calculator (percentage, tiered, flat; min/max caps)
│   │   │   ├── validators.go                 # Fee caps, non-negative, eligibility & stacking rules
│   │   │   ├── dto.go                        # FeeDTO, FeeRequestDTO, FeeCalculationResultDTO
│   │   │   └── rules_engine.go               # Rules engine (user type, volume, geography)
│   │   │
│   │   ├── fee_v2/
│   │   │   ├── service.go                    # Tiered/volume, locale exceptions, coupons, experiments
│   │   │   ├── commands.go                   # ApplyCoupon, RevokeCoupon, UpdateFeeRules
│   │   │   ├── queries.go                    # GetEffectiveFee, GetUserFeeTier, GetCoupon
│   │   │   ├── validators.go                 # Coupon validity, stacking, locale overrides, experiment flags
│   │   │   ├── dto.go                        # FeeV2DTO, CouponDTO, FeeRuleDTO
│   │   │   └── mapper.go                     # Fee v2 ⇄ DTO mappers
│   │   │
│   │   # ===== REFUND PROCESSING SERVICES =====
│   │   ├── refund/
│   │   │   ├── service.go                    # Process & Cancel refunds (full/partial) with audit trail
│   │   │   ├── commands.go                   # ProcessRefund, CancelRefund, ProcessPartialRefund
│   │   │   ├── queries.go                    # GetRefund, ListRefunds (by payment/user/contract)
│   │   │   ├── validators.go                 # Pro-rata vs eligibility windows, idempotency, limits
│   │   │   ├── dto.go                        # RefundDTO, RefundRequestDTO, PartialRefundDTO
│   │   │   └── mapper.go                     # Refund ⇄ DTO mappers
│   │   │
│   │   # ===== TAX SERVICES =====
│   │   ├── tax/
│   │   │   ├── service.go                    # Calculate, Generate forms, File (or export)
│   │   │   ├── form_generator.go             # Generate tax forms (W9, 1099, 1099-K, VAT returns)
│   │   │   ├── commands.go                   # SyncTaxProfileFromUsersBE, Generate1099K, FileVATReturn
│   │   │   ├── queries.go                    # GetTaxProfile, GetVATRate, GetReverseChargeEligibility
│   │   │   ├── validators.go                 # VAT ID formats per locale, threshold checks, 1099-K thresholds
│   │   │   ├── dto.go                        # TaxDTO, TaxFormDTO, VATRateDTO
│   │   │   └── mapper.go                     # Tax ⇄ DTO mappers
│   │   │
│   │   # ===== FOREIGN EXCHANGE SERVICES =====
│   │   ├── fx/
│   │   │   ├── service.go                    # Quote→settlement conversion, fetch/apply FX rates, rounding
│   │   │   ├── commands.go                   # UpsertFXRate, SetRoundingRule
│   │   │   ├── queries.go                    # GetFXRate, ConvertAmount, GetEffectiveRateAt
│   │   │   ├── validators.go                 # Effective timestamps, precision, allowed currency pairs
│   │   │   ├── dto.go                        # FXRateDTO, ConversionRequestDTO, ConversionResultDTO
│   │   │   └── mapper.go                     # FX ⇄ DTO mappers
│   │   │
│   │   # ===== RISK MANAGEMENT SERVICES =====
│   │   ├── risk/
│   │   │   ├── service.go                    # Holds, reserves, chargebacks, negative balances orchestration
│   │   │   ├── commands.go                   # CreateReserve, ReleaseReserve, PlaceHold, RemoveHold, RecordChargeback
│   │   │   ├── queries.go                    # GetReserves, GetHolds, GetRiskScore (if available)
│   │   │   ├── validators.go                 # Rolling reserve schedules, chargeback states, policy gates
│   │   │   ├── dto.go                        # RiskDTO, ReserveDTO, HoldDTO, ChargebackDTO
│   │   │   └── mapper.go                     # Risk ⇄ DTO mappers
│   │   │
│   │   # ===== PROTECTION PLAN SERVICES =====
│   │   ├── protection_plan/
│   │   │   ├── service.go                    # Apply plan, file/approve claims, compute payouts
│   │   │   ├── commands.go                   # ApplyProtection, FileClaim, ApproveClaim, PayClaim
│   │   │   ├── queries.go                    # GetPlan, GetClaims, GetEligibility
│   │   │   ├── validators.go                 # Eligibility checks & coverage caps
│   │   │   ├── dto.go                        # ProtectionPlanDTO, ClaimDTO, EligibilityDTO
│   │   │   └── mapper.go                     # Protection plan ⇄ DTO mappers
│   │   │
│   │   # ===== FEE UPDATE & MIGRATION SERVICES =====
│   │   ├── fee_update/
│   │   │   ├── service.go                    # Manage fee versions & migrations; compute impact
│   │   │   ├── commands.go                   # PublishFeeVersion, RunFeeMigration, ActivateFeeVersion, RollbackFeeVersion
│   │   │   ├── queries.go                    # GetActiveFeeVersion, GetImpact, ListFeeVersions
│   │   │   ├── validators.go                 # Version monotonicity, rate bounds, window checks
│   │   │   ├── dto.go                        # FeeUpdateDTO, FeeImpactDTO, FeeVersionDTO
│   │   │   └── mapper.go                     # Fee update ⇄ DTO mappers
│   │   │
│   │   # ===== INTERNATIONAL PAYMENT SERVICES =====
│   │   ├── international_payment/
│   │   │   ├── service.go                    # Run compliance, convert FX, route to local rails
│   │   │   ├── commands.go                   # InitiateInternationalPayment, MarkCompliant, RoutePayment
│   │   │   ├── queries.go                    # GetInternationalPayment, ListInternationalPayments
│   │   │   ├── validators.go                 # OFAC/AML gates, supported corridors
│   │   │   ├── dto.go                        # InternationalPaymentDTO, ComplianceDTO
│   │   │   └── mapper.go                     # Intl payment ⇄ DTO mappers
│   │   │
│   │   # ===== BONUS SERVICES =====
│   │   ├── bonus/
│   │   │   ├── service.go                    # Award & pay bonuses, enforce conditions
│   │   │   ├── commands.go                   # AwardBonus, RejectBonus, PayBonus
│   │   │   ├── queries.go                    # GetBonuses (by contract/user), GetBonusPolicy
│   │   │   ├── validators.go                 # Condition checks, caps, eligibility
│   │   │   ├── dto.go                        # BonusDTO, AwardBonusDTO
│   │   │   └── mapper.go                     # Bonus ⇄ DTO mappers
│   │   │
│   │   # ===== EXPENSE SERVICES =====
│   │   ├── expense/
│   │   │   ├── service.go                    # Submit/approve/reimburse expenses
│   │   │   ├── commands.go                   # SubmitExpense, ApproveExpense, RejectExpense, ReimburseExpense
│   │   │   ├── queries.go                    # GetExpenses, GetExpenseByID
│   │   │   ├── validators.go                 # Policy caps, categories, receipt validation
│   │   │   ├── dto.go                        # ExpenseDTO, ReceiptDTO
│   │   │   └── mapper.go                     # Expense ⇄ DTO mappers
│   │   │
│   │   # ===== PAYMENT SCHEDULE SERVICES =====
│   │   ├── payment_schedule/
│   │   │   ├── service.go                    # Create/update schedules, run due payments
│   │   │   ├── commands.go                   # CreateSchedule, UpdateSchedule, RunDuePayments
│   │   │   ├── queries.go                    # GetSchedules, GetNextDue
│   │   │   ├── validators.go                 # Conflicts, proration, disabled checks
│   │   │   ├── dto.go                        # PaymentScheduleDTO, FrequencyDTO
│   │   │   └── mapper.go                     # Schedule ⇄ DTO mappers
│   │   │
│   │   # ===== REMINDER SERVICES =====
│   │   ├── reminder/
│   │   │   ├── service.go                    # Trigger & escalate reminders (dunning support)
│   │   │   ├── commands.go                   # TriggerReminder, EscalateReminder
│   │   │   ├── queries.go                    # GetReminders, GetReminderTemplates
│   │   │   ├── validators.go                 # Channel availability, throttle
│   │   │   ├── dto.go                        # ReminderDTO, ReminderTemplateDTO
│   │   │   └── mapper.go                     # Reminder ⇄ DTO mappers
│   │   │
│   │   # ===== INSURANCE SERVICES =====
│   │   ├── insurance/
│   │   │   ├── service.go                    # Apply policy, file/approve claims, coordinate payouts
│   │   │   ├── commands.go                   # ApplyInsurance, FileClaim, ApproveClaim, PayInsuranceClaim
│   │   │   ├── queries.go                    # GetInsurance, GetClaims, GetCoverage
│   │   │   ├── validators.go                 # Policy eligibility, coverage limits
│   │   │   ├── dto.go                        # InsuranceDTO, InsuranceClaimDTO
│   │   │   └── mapper.go                     # Insurance ⇄ DTO mappers
│   │   │
│   │   # ===== TAX FORM SERVICES =====
│   │   ├── tax_form/
│   │   │   ├── service.go                    # Submit/verify forms; link to tax repo
│   │   │   ├── commands.go                   # SubmitTaxForm, VerifyTaxForm
│   │   │   ├── queries.go                    # GetTaxForms, GetTaxFormByID
│   │   │   ├── validators.go                 # Form completeness, identity checks
│   │   │   ├── dto.go                        # TaxFormDTO
│   │   │   └── mapper.go                     # Tax form ⇄ DTO mappers
│   │   │
│   │   # ===== PAYROLL SERVICES =====
│   │   ├── payroll/
│   │   │   ├── service.go                    # Process payroll; apply withholdings & deductions
│   │   │   ├── commands.go                   # ProcessPayroll, WithholdTax
│   │   │   ├── queries.go                    # GetPayroll, GetPayPeriods
│   │   │   ├── validators.go                 # Classification rules, closed periods
│   │   │   ├── dto.go                        # PayrollDTO, PayPeriodDTO
│   │   │   └── mapper.go                     # Payroll ⇄ DTO mappers
│   │   │
│   │   # ===== CURRENCY SERVICES =====
│   │   ├── currency/
│   │   │   ├── service.go                    # Manage preferred currency & conversions
│   │   │   ├── commands.go                   # LockRate, ConvertNow, SetPreferredCurrency
│   │   │   ├── queries.go                    # GetPreferredCurrency, GetRateLock
│   │   │   ├── validators.go                 # Lock windows, precision, supported pairs
│   │   │   ├── dto.go                        # CurrencyPreferenceDTO, RateLockDTO
│   │   │   └── mapper.go                     # Currency ⇄ DTO mappers
│   │   │
│   │   # ===== BANK ACCOUNT SERVICES =====
│   │   └── bank_account/                     # 🆕 Bank account app module
│   │       ├── service.go                    # CRUD bank accounts, verify, set default payout method
│   │       ├── commands.go                   # AddBankAccount, VerifyBankAccount, SetDefaultBankAccount, RemoveBankAccount
│   │       ├── queries.go                    # GetBankAccount, ListBankAccounts, GetDefaultBankAccount
│   │       ├── validators.go                 # Validate routing number, account number, bank details
│   │       ├── dto.go                        # BankAccountDTO, AddBankAccountDTO, VerificationDTO
│   │       └── mapper.go                     # BankAccount ⇄ DTO mappers
│   │
│   # ──────────────────────────────────────────────────────────────────────────────────
│   # 🌐 INTERFACES LAYER (HTTP/API - Load Fifth)
│   # ──────────────────────────────────────────────────────────────────────────────────
│   └── interfaces/
│       └── http/
│           │
│           # ===== API VERSIONING =====
│           ├── v1/                           # 🆕 Versioned API surface (all handlers/routes under /v1)
│           │   │
│           │   # ===== HANDLERS =====
│           │   ├── handlers/
│           │   │   │
│           │   │   # ===== CORE WALLET & LEDGER HANDLERS =====
│           │   │   ├── wallet_handler.go         # GET /v1/wallets/:id, POST /v1/wallets/:id/deposit, /withdraw, /transfer
│           │   │   ├── transaction_handler.go    # GET /v1/transactions, GET /v1/transactions/:id
│           │   │   ├── ledger_journal_handler.go # POST /v1/ledger/entry, POST /v1/ledger/adjustment, GET /v1/ledger/audit
│           │   │   │
│           │   │   # ===== PAYMENT & ESCROW HANDLERS =====
│           │   │   ├── payment_handler.go        # POST /v1/payments, POST /v1/payments/:id/capture, /void, /refund, GET /v1/payments/:id
│           │   │   ├── escrow_handler.go         # POST /v1/escrow/hold, POST /v1/escrow/release, POST /v1/escrow/partial-release
│           │   │   ├── payout_handler.go         # POST /v1/payouts, GET /v1/payouts/:id, POST /v1/payouts/:id/cancel
│           │   │   │
│           │   │   # ===== INVOICE & FEE HANDLERS =====
│           │   │   ├── invoice_handler.go        # POST /v1/invoices, GET /v1/invoices/:id, POST /v1/invoices/:id/send, /pay
│           │   │   ├── fee_handler.go            # GET /v1/fees/calculate
│           │   │   ├── fee_v2_handler.go         # GET /v1/fees/v2/effective, POST /v1/fees/v2/coupon
│           │   │   │
│           │   │   # ===== REFUND & TAX HANDLERS =====
│           │   │   ├── refund_handler.go         # POST /v1/refunds, POST /v1/refunds/:id/cancel
│           │   │   ├── tax_handler.go            # GET /v1/tax/forms, POST /v1/tax/1099k
│           │   │   │
│           │   │   # ===== FX & RISK HANDLERS =====
│           │   │   ├── fx_handler.go             # GET /v1/fx/rate?base=&quote=, POST /v1/fx/rate
│           │   │   ├── risk_handler.go           # POST /v1/risk/reserve, POST /v1/risk/hold, GET /v1/risk/holds
│           │   │   │
│           │   │   # ===== PROTECTION & FEE UPDATE HANDLERS =====
│           │   │   ├── protection_plan_handler.go    # POST /v1/protection/apply, POST /v1/protection/claims, GET /v1/protection/:contractId
│           │   │   ├── fee_update_handler.go         # POST /v1/fees/updates, POST /v1/fees/activate, GET /v1/fees/active
│           │   │   │
│           │   │   # ===== INTERNATIONAL PAYMENT HANDLERS =====
│           │   │   ├── international_payment_handler.go # POST /v1/intl/payments, GET /v1/intl/payments/:id
│           │   │   │
│           │   │   # ===== BONUS & EXPENSE HANDLERS =====
│           │   │   ├── bonus_handler.go              # POST /v1/bonus/award, GET /v1/bonus?contractId=
│           │   │   ├── expense_handler.go            # POST /v1/expenses, PUT /v1/expenses/:id/approve, PUT /v1/expenses/:id/reject
│           │   │   │
│           │   │   # ===== SCHEDULE & REMINDER HANDLERS =====
│           │   │   ├── payment_schedule_handler.go   # POST /v1/schedules, PUT /v1/schedules/:id, POST /v1/schedules/run-due
│           │   │   ├── reminder_handler.go           # POST /v1/reminders/trigger, POST /v1/reminders/escalate, GET /v1/reminders
│           │   │   │
│           │   │   # ===== INSURANCE & TAX FORM HANDLERS =====
│           │   │   ├── insurance_handler.go          # POST /v1/insurance/apply, POST /v1/insurance/claims, GET /v1/insurance/:contractId
│           │   │   ├── tax_form_handler.go           # POST /v1/tax/forms, GET /v1/tax/forms?contractId=
│           │   │   │
│           │   │   # ===== PAYROLL & BANK ACCOUNT HANDLERS =====
│           │   │   ├── payroll_handler.go            # POST /v1/payroll/process, GET /v1/payroll?contractId=
│           │   │   ├── bank_account_handler.go       # 🆕 POST /v1/bank-accounts, GET /v1/bank-accounts/:id, DELETE /v1/bank-accounts/:id
│           │   │   │
│           │   │   # ===== WEBHOOK HANDLERS =====
│           │   │   ├── webhook_handler.go        # POST /v1/webhooks/stripe, POST /v1/webhooks/paypal
│           │   │   │                             # 🆕 Includes signature verification middleware
│           │   │   │
│           │   │   # ===== HEALTH HANDLERS =====
│           │   │   └── health_handler.go         # 🆕 GET /v1/healthz/live, GET /v1/healthz/ready
│           │   │
│           │   # ===== ROUTES =====
│           │   ├── routes/
│           │   │   │
│           │   │   # ===== CORE ROUTES =====
│           │   │   ├── wallet_routes.go          # /v1/wallets/*
│           │   │   ├── transaction_routes.go     # /v1/transactions/*
│           │   │   ├── ledger_journal_routes.go  # /v1/ledger/*
│           │   │   │
│           │   │   # ===== PAYMENT & ESCROW ROUTES =====
│           │   │   ├── payment_routes.go         # /v1/payments/*
│           │   │   ├── escrow_routes.go          # /v1/escrow/*
│           │   │   ├── payout_routes.go          # /v1/payouts/*
│           │   │   │
│           │   │   # ===== INVOICE & FEE ROUTES =====
│           │   │   ├── invoice_routes.go         # /v1/invoices/*
│           │   │   ├── fee_routes.go             # /v1/fees/*
│           │   │   ├── fee_v2_routes.go          # /v1/fees/v2/*
│           │   │   │
│           │   │   # ===== REFUND & TAX ROUTES =====
│           │   │   ├── refund_routes.go          # /v1/refunds/*
│           │   │   ├── tax_routes.go             # /v1/tax/*
│           │   │   │
│           │   │   # ===== FX & RISK ROUTES =====
│           │   │   ├── fx_routes.go              # /v1/fx/*
│           │   │   ├── risk_routes.go            # /v1/risk/*
│           │   │   │
│           │   │   # ===== PROTECTION & FEE UPDATE ROUTES =====
│           │   │   ├── protection_plan_routes.go     # /v1/protection/*
│           │   │   ├── fee_update_routes.go          # /v1/fees/updates/*
│           │   │   │
│           │   │   # ===== INTERNATIONAL PAYMENT ROUTES =====
│           │   │   ├── international_payment_routes.go # /v1/intl/*
│           │   │   │
│           │   │   # ===== BONUS & EXPENSE ROUTES =====
│           │   │   ├── bonus_routes.go               # /v1/bonus/*
│           │   │   ├── expense_routes.go             # /v1/expenses/*
│           │   │   │
│           │   │   # ===== SCHEDULE & REMINDER ROUTES =====
│           │   │   ├── payment_schedule_routes.go    # /v1/schedules/*
│           │   │   ├── reminder_routes.go            # /v1/reminders/*
│           │   │   │
│           │   │   # ===== INSURANCE & TAX FORM ROUTES =====
│           │   │   ├── insurance_routes.go           # /v1/insurance/*
│           │   │   ├── tax_form_routes.go            # /v1/tax/forms/*
│           │   │   │
│           │   │   # ===== PAYROLL & BANK ACCOUNT ROUTES =====
│           │   │   ├── payroll_routes.go             # /v1/payroll/*
│           │   │   ├── bank_account_routes.go        # 🆕 /v1/bank-accounts/*
│           │   │   │
│           │   │   # ===== WEBHOOK ROUTES =====
│           │   │   ├── webhook_routes.go         # 🆕 /v1/webhooks/*
│           │   │   │
│           │   │   # ===== HEALTH ROUTES =====
│           │   │   └── health_routes.go          # 🆕 /v1/healthz/*
│           │   │
│           │   # ===== OPENAPI SPEC =====
│           │   └── openapi/
│           │       ├── openapi.yaml              # 🆕 OpenAPI 3.0 spec (served as /v1/openapi)
│           │       └── generator.go              # 🆕 Serves /v1/swagger and /v1/openapi.json (dev only)
│           │
│           # ===== MIDDLEWARE =====
│           ├── middleware/
│           │   ├── requestid.go              # 🆕 Wraps platform-shared request id
│           │   ├── logging.go                # 🆕 Wraps platform-shared logging
│           │   ├── recovery.go               # 🆕 Wraps platform-shared recovery
│           │   ├── tracing.go                # 🆕 Wraps platform-shared otel middleware
│           │   ├── auth.go                   # Uses pkg/auth (Keycloak/OIDC token verification)
│           │   ├── rbac.go                   # ♻️ Uses pkg/auth (role checks & resource scoping)
│           │   │                             # Ensure aligns with application/authz
│           │   ├── cors.go                   # Uses platform-shared/ginx (CORS config)
│           │   ├── rate_limit.go             # Rate limiting (per-IP/per-user, burst)
│           │   ├── idempotency.go            # Uses platform-shared/idempotency (CRITICAL for payments!)
│           │   └── webhook_signature.go      # 🆕 Webhook signature verification (Stripe/PayPal)
│           │
│           # ===== ERROR PRESENTERS =====
│           ├── presenters/
│           │   └── errors.go                 # 🆕 Maps domain/service errors → HTTP status & problem+json
│           │
│           # ===== ROUTER =====
│           └── router.go                     # ♻️ Uses platform-shared/ginx; mounts v1 route registrars & middleware
│
├── db/                                       # 🆕 Developer-friendly entrypoint for SQL
│   └── migrations/                           # 🆕 Symlink or mirror of internal/.../migrations (optional)
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                         # Base error helpers (wrap, with code/context)
│   │   ├── codes.go                          # Error codes: INSUFFICIENT_FUNDS, PAYMENT_FAILED, ESCROW_LOCK, etc.
│   │   └── payment_errors.go                 # Payment-specific error types (GatewayDecline, 3DSRequired, RiskHold)
│   │
│   ├── logger/                               # ❌ REMOVED
│   │   └── README.md                         # Placeholder: use platform-shared logging
│   │
│   ├── utils/
│   │   ├── validator.go                      # Common validation helpers (UUIDs, enums, ranges)
│   │   ├── currency.go                       # Currency conversion helpers (symbols, ISO codes)
│   │   ├── decimal.go                        # Decimal math for money (fixed-point ops, rounding)
│   │   ├── encryption.go                     # Encryption utilities (field-level encryption, PCI scope)
│   │   └── etag.go                           # 🆕 ETag generation helpers (pairs with infrastructure/http/etag)
│   │
│   └── constants/
│       ├── events.go                         # ❌ REMOVED (use platform-shared constants)
│       ├── topics.go                         # ❌ REMOVED (use platform-shared topics)
│       ├── currencies.go                     # Currency constants (USD, EUR, GBP, supported list)
│       └── payment_methods.go                # Payment method constants (Stripe, PayPal, BankTransfer, Wallet)
│
├── config/                                   # Runtime config profiles
│   ├── default.yaml                          # Base config (safe defaults)
│   ├── dev.yaml                              # Development overrides (local gateways, debug flags)
│   └── prod.yaml                             # Production overrides (timeouts, pools, feature flags)
│
├── dapr/                                     # Dapr components
│   ├── local/
│   │   ├── pubsub.yaml                       # Local pub/sub (scoped to financial-be)
│   │   └── statestore.yaml                   # Local state store (idempotency keys, outbox)
│   └── k8s/
│       ├── pubsub.yaml                       # Scopes: ["financial-be"] (topics for payments/escrow/payouts/etc.)
│       ├── statestore.yaml                   # State store for production (HA, TTLs)
│       └── secrets.yaml                      # References to K8s secrets (API keys, tokens)
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                   # Deployment spec (resources, env, probes)
│       ├── service.yaml                      # Service spec (ports, selectors)
│       ├── configmap.yaml                    # App config mounted as files/env
│       ├── secrets.yaml                      # Contains Stripe/PayPal API keys, DB creds (sealed/encrypted)
│       ├── hpa.yaml                          # Horizontal Pod Autoscaler (CPU/RAM/requests-per-sec)
│       ├── pdb.yaml                          # PodDisruptionBudget (availability during maintenance)
│       ├── networkpolicy.yaml                # Extra security (ingress/egress restrictions)
│       └── servicemonitor.yaml               # Prometheus ServiceMonitor (metrics scraping)
│
├── scripts/
│   ├── setup-local.sh                        # Local dev bootstrap (migrations, seeds, Dapr components)
│   ├── migrate.sh                            # 🆕 SQL migrations (dev/prod flags)
│   ├── openapi-diff.sh                       # 🆕 Fail CI on breaking API changes
│   ├── schema-diff.sh                        # 🆕 DB schema guardrail (pg_dump + mig verify)
│   ├── generate-sdks.sh                      # 🆕 Produce /sdk clients from OpenAPI
│   ├── get-secrets.sh                        # Script to fetch secrets to local env (vault/secret manager)
│   ├── seed-data.sh                          # Seed demo data (wallets, rates, fee rules)
│   ├── reconciliation.sh                     # Daily reconciliation script (bank vs. journal; reports)
│   └── sbom-sign.sh                          # 🆕 Build SBOM + cosign sign
│
├── tests/
│   ├── unit/                                 # Unit tests (domain/application)
│   ├── integration/                          # Integration tests (DB/cache/gateways)
│   ├── e2e/                                  # End-to-end tests (HTTP flows, idempotency, race conditions)
│   │   └── chaos_event_delivery_test.go      # 🆕 Simulate duplicates, reordering, delays
│   │
│   ├── reliability/                          # 🆕 Reliability tests
│   │   ├── projections_replay_test.go        # 🆕 Rebuild read models from event logs
│   │   └── outbox_dispatcher_test.go         # 🆕 At-least-once + idempotency assertions
│   │
│   └── property/                             # 🆕 Property-based tests
│       └── fee_calculation_property_test.go  # 🆕 Property/fuzz tests for fee calculation kernel
│
├── docs/
│   ├── README.md                             # Overview & getting started
│   ├── API.md                                # HTTP API surface (endpoints, payloads, examples)
│   ├── API_VERSIONING.md                     # 🆕 HTTP contract versioning & deprecation policy
│   ├── OPENAPI.md                            # 🆕 How to regenerate SDKs from OpenAPI spec
│   ├── EVENTS.md                             # Events (payment.processed, escrow.held/released, payout.processed, fee.updated, intl.payment.*)
│   ├── ARCHITECTURE.md                       # Service architecture (layers, dependencies, flows)
│   ├── MIGRATIONS.md                         # DB migrations & versioning strategy
│   ├── SCHEMA.md                             # Domain schemas (journal, fees v2, fx, protection/insurance/schedules)
│   ├── RUNBOOK.md                            # Ops runbook (alerts, dashboards, SLOs, incident steps)
│   ├── SLOS.md                               # 🆕 Target P99s, projection lag, queue delay
│   ├── OUTBOX.md                             # 🆕 Exactly-once-ish design, retries, DLQ
│   ├── CACHING.md                            # 🆕 Keys, TTLs, SWR, invalidation events
│   ├── DATA_RETENTION.md                     # 🆕 Per-domain retention windows
│   ├── ERASURE.md                            # 🆕 GDPR/CCPA erasure hooks & playbooks
│   ├── RELEASE_CHECKLIST.md                  # 🆕 Preflight (openapi diff, schema diff, migrations plan)
│   ├── NAMING.md                             # 🆕 Package snake_case; HTTP kebab-case conventions
│   ├── payment-flows.md                      # Payment flow documentation (auth/capture/refund)
│   ├── escrow-system.md                      # Escrow system documentation (holds/releases/disputes)
│   ├── fee-structure.md                      # Platform fee structure (old engine)
│   ├── fee-v2.md                             # Fee engine v2 rules & examples (tiers, coupons, locales)
│   ├── fx-and-rounding.md                    # FX, effective timestamps, rounding rules
│   ├── risk-holds.md                         # Reserves, holds, chargebacks (policies & flows)
│   ├── intl-payments.md                      # International payments flows & compliance checks
│   ├── protection-plans.md                   # Protection plans & claims flows (eligibility, payouts)
│   ├── schedules-and-reminders.md            # PaymentSchedules & Reminders (cron windows, dunning)
│   └── payroll-and-taxforms.md               # Payroll + tax_form integration notes (jurisdictions)
│
├── sdk/                                      # 🆕 Generated clients (optional but handy)
│   ├── go/                                   # 🆕 Generated Go client from OpenAPI
│   └── ts/                                   # 🆕 Generated TypeScript client from OpenAPI
│
├── .github/
│   └── workflows/
│       ├── ci.yml                            # CI pipeline (lint, test, build)
│       ├── cd.yml                            # CD pipeline (image build/push, deploy)
│       ├── contract-ci.yml                   # 🆕 OpenAPI-diff + event schema checks
│       ├── security.yml                      # 🆕 golangci-lint, govulncheck, trivy, cosign verify
│       └── load-tests.yml                    # 🆕 k6/gatling smoke against /healthz & hot paths
│
├── go.mod                                    # Imports pkg/auth, platform-shared libs, contracts/events schemas
├── go.sum                                    # Module dependency checksums
├── .env.example                              # Example environment variables (non-secret)
├── .golangci.yml                             # 🆕 Linter config (CI parity)
├── CODEOWNERS                                # 🆕 Explicit ownership per context (ease future splits)
├── Makefile                                  # Common dev tasks (run, test, lint, migrate)
├── Dockerfile                                # Container build (multi-stage; minimal runtime)
├── .dockerignore                             # Docker build context exclusions
├── .gitignore                                # Git ignore rules
└── README.md                                 # High-level service description & quickstart


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

✅ **Bidding system** in proposals-be
✅ **In-app notifications** with comprehensive coverage in communications-be
✅ **Free & self-hosted** communications (WildDuck, WebSocket, SSE)
✅ **Auto-migrations** following your users-be pattern
✅ **Enhanced recommendation system** in search-be with ML models
✅ **Reviews covering ratings** in reviews-be
✅ Complete event flows between all services
✅ Enterprise-level folder structures

Each service is production-ready and follows Go best practices with Clean Architecture!