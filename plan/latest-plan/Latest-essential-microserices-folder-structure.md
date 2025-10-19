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
│   └── api/
│       └── main.go                           # Entry point - initializes Gin, Dapr, connects to Postgres
│
├── internal/
│   ├── domain/                               # 🏛️ Domain Layer
│   │   ├── contract/
│   │   │   ├── entity.go                     # Contract aggregate root (ID, JobID, FreelancerID, ClientID, Terms, Amount, Status)
│   │   │   ├── terms.go                      # Contract terms (payment terms, deliverables, timeline)
│   │   │   ├── enums.go                      # Type (Fixed, Hourly), Status (Active, Paused, Completed, Terminated), PaymentType
│   │   │   ├── value_objects.go              # ContractAmount, ContractDuration
│   │   │   ├── errors.go                     # Domain errors (ContractNotFound, InvalidTerms)
│   │   │   ├── repository.go                 # ContractRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: ContractCreated/Activated/Paused/Resumed/Completed/Terminated
│   │   │
│   │   ├── milestone/
│   │   │   ├── entity.go                     # Contract milestones (ContractID, Description, Amount, DueDate, Status)
│   │   │   ├── delivery.go                   # Milestone deliverables (files, descriptions)
│   │   │   ├── status.go                     # Milestone status tracking (InProgress, Submitted, UnderReview, Approved, Rejected)
│   │   │   ├── acceptance_policy.go          # 🆕 Acceptance windows, grace periods, per-milestone rules
│   │   │   ├── auto_release.go               # 🆕 Auto-release rules (timed release when not reviewed)
│   │   │   ├── errors.go                     # Milestone errors
│   │   │   ├── repository.go                 # MilestoneRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: MilestoneCreated/Submitted/Approved/Rejected/Released/AutoReleased
│   │   │
│   │   ├── deliverable/
│   │   │   ├── entity.go                     # Contract deliverables (ContractID, Description, FileURLs, SubmittedAt)
│   │   │   ├── submission.go                 # Deliverable submissions (version tracking)
│   │   │   ├── revision.go                   # Revision requests (feedback, changes needed)
│   │   │   ├── acceptance_link.go            # 🆕 Link deliverables to milestone acceptance state
│   │   │   ├── errors.go                     # Deliverable errors
│   │   │   ├── repository.go                 # DeliverableRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: DeliverableSubmitted/Accepted/RevisionRequested/RevisionSubmitted
│   │   │
│   │   ├── timesheet/
│   │   │   ├── entity.go                     # Time entries (ContractID, WeekStartDate, TotalHours, Status)
│   │   │   ├── week.go                       # Weekly timesheets (aggregated by week)
│   │   │   ├── time_entry.go                 # Individual time entry (Date, Hours, Description, Manual/Tracked)
│   │   │   ├── screenshot.go                 # Work diary screenshots (TimeEntryID, URL, Timestamp)
│   │   │   ├── weekly_caps.go                # 🆕 Weekly approvals & hour caps (rate * cap guardrails)
│   │   │   ├── errors.go                     # Timesheet errors
│   │   │   ├── repository.go                 # TimesheetRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: TimesheetSubmitted/Approved/Rejected/Locked/Unlocked
│   │   │
│   │   ├── workdiary/
│   │   │   ├── entity.go                     # Work diary entries (ContractID, Date, Hours, Screenshots, ActivityLevel)
│   │   │   ├── activity_level.go             # Activity tracking (keyboard/mouse activity percentage)
│   │   │   ├── policy.go                     # 🆕 Screenshot policy flags (blur/off-hours/manual edits) & compliance
│   │   │   ├── metrics.go                    # 🆕 Aggregated activity metrics & anomaly flags
│   │   │   ├── errors.go                     # Work diary errors
│   │   │   └── repository.go                 # WorkDiaryRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: ScreenshotCaptured/ActivityRecorded/PolicyViolated
│   │   │
│   │   ├── template/
│   │   │   ├── entity.go                     # Contract templates (Title, Terms, Clauses)
│   │   │   ├── clause.go                     # Contract clauses (standard terms, custom clauses)
│   │   │   ├── errors.go                     # Template errors
│   │   │   ├── repository.go                 # TemplateRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: TemplateCreated/Updated/Archived
│   │   │
│   │   ├── amendment/
│   │   │   ├── entity.go                     # Contract amendments (ContractID, Type, OldValue, NewValue, Status)
│   │   │   ├── change_request.go             # Change requests (scope, budget, timeline changes)
│   │   │   ├── rate_change.go                # 🆕 Rate changes with EffectiveDate
│   │   │   ├── scope_addendum.go             # 🆕 Scope addenda (additional items linked to SOW)
│   │   │   ├── errors.go                     # Amendment errors
│   │   │   ├── repository.go                 # AmendmentRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: AmendmentRequested/Approved/Rejected/Applied
│   │   │
│   │   ├── pause/
│   │   │   ├── entity.go                     # Contract pause/resume (ContractID, Reason, PausedAt, ResumedAt)
│   │   │   ├── policy.go                     # 🆕 Pause policy (max duration, billing behavior during pause)
│   │   │   ├── errors.go                     # Pause errors
│   │   │   ├── repository.go                 # PauseRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: ContractPausePlaced/Extended/Resumed
│   │   │
│   │   ├── termination/
│   │   │   ├── entity.go                     # Contract termination (ContractID, Reason, TerminatedBy, TerminatedAt)
│   │   │   ├── reason.go                     # Termination reasons (Completed, Breach, Mutual, Unilateral)
│   │   │   ├── errors.go                     # Termination errors
│   │   │   ├── repository.go                 # TerminationRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: TerminationRequested/Confirmed/RolledBack
│   │   │
│   │   ├── dispute/
│   │   │   ├── entity.go                     # Contract disputes (ContractID, Reason, FiledBy, Status, Resolution)
│   │   │   ├── evidence.go                   # Evidence submission (files, descriptions, timestamps)
│   │   │   ├── resolution.go                 # Dispute resolution (outcome, awarded amount, penalties)
│   │   │   ├── status.go                     # Dispute status (Open, UnderReview, Mediation, Resolved, Escalated)
│   │   │   ├── anchors.go                    # 🆕 Anchor contract states (under-review, on-hold) during disputes
│   │   │   ├── errors.go                     # Dispute errors
│   │   │   ├── repository.go                 # DisputeRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: DisputeOpened/EvidenceSubmitted/ResolutionUpdated/Resolved/Escalated
│   │   │
│   │   ├── sow/                              # 🆕 NEW DOMAIN: Statement of Work (SOW) versioning
│   │   │   ├── entity.go                     # SOW aggregate (ContractID, CurrentVersionID, Active)
│   │   │   ├── version.go                    # Versioned scope items (Version, Scope, Budget, CreatedBy, CreatedAt, ApprovedBy, ApprovedAt)
│   │   │   ├── changelog.go                  # Changelog entries (who/when/what changed; approvals)
│   │   │   ├── enums.go                      # SOW status (Draft, PendingApproval, Approved, Rejected)
│   │   │   ├── errors.go                     # SOW domain errors (SOWNotFound, VersionConflict, NotApprover)
│   │   │   ├── repository.go                 # SOWRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: SOWVersionCreated/SubmittedForApproval/Approved/Rejected/Activated
│   │   │
│   │   └── financial_hold/                   # 🆕 NEW DOMAIN: Financial holds aligned to disputes
│   │       ├── entity.go                     # FinancialHold (ContractID, Amount, Reason, Status, PlacedBy, PlacedAt, ReleasedAt)
│   │       ├── reason.go                     # Hold reasons (Dispute, Compliance, Chargeback, Manual)
│   │       ├── status.go                     # Hold status (Active, Released)
│   │       ├── errors.go                     # Financial hold errors
│   │       ├── repository.go                 # FinancialHoldRepository interface
│   │       └── events.go                     # 🆕 Domain events: FinancialHoldPlaced/Adjusted/Released
│   │
│   ├── application/                          # 📋 Application Layer (complete with commands/queries/service/validators)
│   │   ├── eventhandler/
│   │   │   ├── proposal_handler.go            # Consumes: proposal.accepted (create Contract);
│   │   │   │                                     #           invite.accepted | negotiation.accepted (pre-contract flows that finalize into ContractCreated).
│   │   │   │
│   │   │   ├── payment_handler.go             # Consumes: payment.processed (advance contract/milestone state; emit ContractPaid/PaymentProcessed);
│   │   │   │                                     #           payment.failed (record failure → retry/dunning hooks).
│   │   │   │
│   │   │   ├── escrow_handler.go              # Consumes: escrow.released (release path confirms payout → close/approve relevant milestone);
│   │   │   │                                     #           escrow.funded | escrow.refund.processed (sync escrow→contract financial status).
│   │   │   │
│   │   │   ├── dispute_handler.go             # Consumes: dispute.opened (anchor contract under-review/on-hold);
│   │   │   │                                     #           dispute.resolved | dispute.closed (resume/complete, may trigger ContractEnded).
│   │   │   │
│   │   │   ├── admin_handler.go               # Consumes: admin.feature_flag.updated | admin.config.updated (runtime flags/config used by contracts-be);
│   │   │   │                                     #           admin.moderation.actioned (approve/reject/remove/hide → reflect visibility/state if contract content is impacted);
│   │   │   │                                     #           admin.case.user_report.* (link case refs affecting contract participants/visibility).
│   │   │   │
│   │   │   ├── financial_risk_handler.go      # Consumes: financial.risk.alert.emitted (risk signal → place internal hold, may emit contract.financial_hold.placed);
│   │   │   │                                     #           financial.chargeback.created | financial.chargeback.updated (mirror chargeback state; may emit hold placed/released).
│   │   │   │
│   │   │   ├── contract_status_handler.go     # Consumes: contract.financial_hold.placed | contract.financial_hold.released (idempotently sync local read models, cascade to policy checks).
│   │   │   │
│   │   │   ├── message_handler.go             # Consumes: message.notification_delivered (optional: contract-related notification hygiene, SLA timers).
│   │   │   │
│   │   │   └── system_handler.go              # Consumes: admin.threshold.updated | admin.experiment.updated (feature gates/thresholds that influence auto-close, review windows).
│   │   ├
│   │   ├── contract/
│   │   │   ├── service.go                    # Contract logic (Create, Update, Pause, Resume, End)
│   │   │   ├── commands.go                   # CreateContract, UpdateContract, PauseContract, EndContract
│   │   │   ├── queries.go                    # GetContract, ListContracts, FilterContracts
│   │   │   ├── dto.go                        # ContractDTO, CreateContractDTO, UpdateContractDTO
│   │   │   ├── mapper.go                     # Entity-DTO mapping
│   │   │   ├── validators.go                 # Contract validation (terms, amounts, dates)
│   │   │   └── business_rules.go             # Business rules (payment release conditions)
│   │   │
│   │   ├── milestone/
│   │   │   ├── service.go                    # Milestone logic (Create, Complete, Approve, Request)
│   │   │   ├── commands.go                   # CreateMilestone, CompleteMilestone, ApproveMilestone, RequestChanges
│   │   │   ├── queries.go                    # GetMilestone, ListMilestones, GetPendingApprovals
│   │   │   ├── dto.go                        # MilestoneDTO
│   │   │   ├── mapper.go                     # Milestone mappers
│   │   │   ├── validators.go                 # Validate grace periods & auto-release
│   │   │   └── policy_evaluator.go           # 🆕 Evaluate acceptance policy per milestone
│   │   │
│   │   ├── deliverable/
│   │   │   ├── service.go                    # Submit, Review, Request revisions
│   │   │   ├── commands.go                   # SubmitDeliverable, ApproveDeliverable, RequestDeliverableChanges
│   │   │   ├── queries.go                    # GetDeliverable, ListDeliverables
│   │   │   ├── dto.go                        # DeliverableDTO
│   │   │   ├── mapper.go                     # Deliverable mappers
│   │   │   └── validators.go                 # Acceptance linkage & revision limits
│   │   │
│   │   ├── timesheet/
│   │   │   ├── service.go                    # Timesheet logic (Log time, Submit week, Approve)
│   │   │   ├── commands.go                   # LogTime, SubmitWeek, ApproveTimesheet, DisputeHours
│   │   │   ├── queries.go                    # GetTimesheets, GetWeeklySummary, Calculate earnings
│   │   │   ├── calculator.go                 # Calculate hours, earnings based on hourly rate
│   │   │   ├── dto.go                        # TimesheetDTO, TimeEntryDTO
│   │   │   ├── mapper.go                     # Timesheet mappers
│   │   │   └── validators.go                 # Max hours, dates, weekly caps, policy flags
│   │   │
│   │   ├── workdiary/
│   │   │   ├── service.go                    # Track activity, Upload screenshots
│   │   │   ├── commands.go                   # AddDiaryEntry, UploadScreenshot, FlagPolicyViolation
│   │   │   ├── queries.go                    # GetDiaryDay, ListDiaryRange, ActivityMetrics
│   │   │   ├── dto.go                        # WorkDiaryDTO
│   │   │   ├── mapper.go                     # Work diary mappers
│   │   │   └── validators.go                 # Screenshot policy & activity thresholds
│   │   │
│   │   ├── template/
│   │   │   ├── service.go                    # Create, Update, Use template
│   │   │   ├── commands.go                   # CreateTemplate, UpdateTemplate
│   │   │   ├── queries.go                    # GetTemplate, ListTemplates
│   │   │   ├── dto.go                        # TemplateDTO
│   │   │   ├── mapper.go                     # Template mappers
│   │   │   └── validators.go                 # Clause & term checks
│   │   │
│   │   ├── amendment/
│   │   │   ├── service.go                    # Request, Approve, Reject
│   │   │   ├── commands.go                   # RequestAmendment, ApproveAmendment, RejectAmendment
│   │   │   ├── queries.go                    # GetAmendment, ListAmendments
│   │   │   ├── dto.go                        # AmendmentDTO
│   │   │   ├── mapper.go                     # Amendment mappers
│   │   │   └── validators.go                 # Effective dates, overlap, scope/rate checks
│   │   │
│   │   ├── termination/
│   │   │   ├── service.go                    # Terminate contract, Calculate final payment
│   │   │   ├── commands.go                   # TerminateContract
│   │   │   ├── queries.go                    # GetTermination
│   │   │   ├── dto.go                        # TerminationDTO
│   │   │   ├── mapper.go                     # Termination mappers
│   │   │   └── validators.go                 # Validate reasons & payouts
│   │   │
│   │   ├── dispute/
│   │   │   ├── service.go                    # Open, Submit evidence, Resolve
│   │   │   ├── commands.go                   # OpenDispute, SubmitEvidence, ResolveDispute, EscalateDispute
│   │   │   ├── queries.go                    # GetDisputes, GetDisputeHistory
│   │   │   ├── dto.go                        # DisputeDTO, EvidenceDTO, ResolutionDTO
│   │   │   ├── mapper.go                     # Dispute mappers
│   │   │   └── validators.go                 # Allowed transitions, anchor states
│   │   │
│   │   ├── sow/                              # SOW versioning application
│   │   │   ├── service.go                    # Versioning, propose changes, approve/reject
│   │   │   ├── commands.go                   # ProposeSOWVersion, ApproveSOW, RejectSOW
│   │   │   ├── queries.go                    # GetSOW, ListSOWVersions, GetChangelog
│   │   │   ├── dto.go                        # SOWDTO, SOWVersionDTO, ChangelogEntryDTO
│   │   │   ├── mapper.go                     # Map SOW entities ⇄ DTOs
│   │   │   └── validators.go                 # Version monotonicity, approver roles
│   │   │
│   │   └── financial_hold/                   # 🆕 Application for financial holds
│   │       ├── service.go                    # Place/release holds; compute net available for payout
│   │       ├── commands.go                   # PlaceHold, ReleaseHold
│   │       ├── queries.go                    # GetHolds, GetActiveHold, GetHeldAmount
│   │       ├── dto.go                        # FinancialHoldDTO
│   │       ├── mapper.go                     # Map financial hold entities ⇄ DTOs
│   │       └── validators.go                 # Hold reasons, amounts >= 0, dispute linkage
│   │
│   ├── infrastructure/                       # 🔧 Infrastructure Layer
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection
│   │   │       ├── transaction.go            # Transaction helpers
│   │   │       ├── migrations.go             # 📝 UPDATED: Auto-migration with version tracking
│   │   │       ├── version.go                # 🆕 Schema version tracking
│   │   │       ├── safety.go                 # 🆕 Pre-migration safety checks
│   │   │       ├── contract_repository.go    # ContractRepository implementation
│   │   │       ├── milestone_repository.go   # MilestoneRepository implementation
│   │   │       ├── deliverable_repository.go # DeliverableRepository implementation
│   │   │       ├── timesheet_repository.go   # TimesheetRepository implementation
│   │   │       ├── workdiary_repository.go   # WorkDiaryRepository implementation
│   │   │       ├── template_repository.go    # TemplateRepository implementation
│   │   │       ├── amendment_repository.go   # AmendmentRepository implementation
│   │   │       ├── pause_repository.go       # PauseRepository implementation
│   │   │       ├── termination_repository.go # TerminationRepository implementation
│   │   │       ├── dispute_repository.go     # DisputeRepository implementation
│   │   │       ├── sow_repository.go         # 🆕 SOWRepository implementation
│   │   │       └── financial_hold_repository.go # 🆕 FinancialHoldRepository implementation
│   │   │
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go             # Redis connection
│   │   │       ├── contract_cache.go         # Contract caching
│   │   │       └── hold_cache.go             # 🆕 Financial hold cache (by contract)
│   │   │
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go               # 📝 UPDATED: Uses platform-shared/inbox
│   │   │       ├── producer.go               # 📝 UPDATED: Uses platform-shared/outbox
│   │   │       ├── topics.go                 # 📝 UPDATED: From contracts/events (contract.created, milestone.completed)
│   │   │       └── scram.go                  # SCRAM authentication
│   │   │
│   │   └── storage/
│   │       └── client.go                     # Storage service client (upload deliverables, screenshots)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── contract_handler.go       # Contract handlers (POST /contracts, GET /contracts/:id)
│   │       │   ├── sow_handler.go            # SOW endpoints (POST/GET SOW versions, approvals)
│   │       │   ├── financial_hold_handler.go # 🆕 Financial holds (POST holds, POST release)
│   │       │   ├── milestone_handler.go      # Milestone handlers (POST /contracts/:id/milestones)
│   │       │   ├── deliverable_handler.go    # Deliverable handlers (POST /contracts/:id/deliverables)
│   │       │   ├── timesheet_handler.go      # Timesheet handlers (POST /contracts/:id/timesheet)
│   │       │   ├── workdiary_handler.go      # Work diary handlers (POST /contracts/:id/workdiary)
│   │       │   ├── template_handler.go       # Template handlers (GET /templates)
│   │       │   ├── amendment_handler.go      # Amendment handlers (POST /contracts/:id/amendments)
│   │       │   ├── termination_handler.go    # Termination handlers (POST /contracts/:id/terminate)
│   │       │   ├── dispute_handler.go        # Dispute handlers (POST /contracts/:id/disputes)
│   │       │   └── health_handler.go         # Health check
│   │       │
│   │       ├── routes/                        # ✅ Route registrars for each handler group
│   │       │   ├── contract_routes.go         # /contracts, /contracts/:id
│   │       │   ├── sow_routes.go              # /contracts/:id/sow, /contracts/:id/sow/versions
│   │       │   ├── financial_hold_routes.go   # 🆕 /contracts/:id/holds, /contracts/:id/holds/:holdId/release
│   │       │   ├── milestone_routes.go        # /contracts/:id/milestones
│   │       │   ├── deliverable_routes.go      # /contracts/:id/deliverables
│   │       │   ├── timesheet_routes.go        # /contracts/:id/timesheet
│   │       │   ├── workdiary_routes.go        # /contracts/:id/workdiary
│   │       │   ├── template_routes.go         # /templates
│   │       │   ├── amendment_routes.go        # /contracts/:id/amendments
│   │       │   ├── termination_routes.go      # /contracts/:id/terminate
│   │       │   └── dispute_routes.go          # /contracts/:id/disputes
│   │       │
│   │       ├── middleware/                    # 📝 UPDATED: Uses platform-shared
│   │       │   ├── auth.go                    # 📝 UPDATED: Uses pkg/auth
│   │       │   ├── rbac.go                    # 📝 UPDATED: Uses pkg/auth
│   │       │   ├── cors.go                    # 📝 UPDATED: Uses platform-shared/ginx
│   │       │   ├── rate_limit.go              # Rate limiting
│   │       │   └── idempotency.go             # 🆕 Uses platform-shared/idempotency
│   │       │
│   │       ├── responses/
│   │       │   └── README.md                  # 📝 Points to platform-shared/httpx
│   │       │
│   │       └── router.go                      # 📝 UPDATED: Uses platform-shared/ginx
│   │
│   └── config/                               # 🆕 MEDIUM
│       ├── schema.go                         # 🆕 Typed Config
│       ├── loader.go                         # 🆕 Viper loader
│       └── docs/
│           └── CONFIGURATION.md              # 🆕 Config docs
│
├── config/                                   # 🆕 MEDIUM
│   ├── default.yaml
│   ├── dev.yaml
│   └── prod.yaml
│
├── dapr/                                     # 🆕 MEDIUM
│   ├── local/
│   │   ├── pubsub.yaml
│   │   └── statestore.yaml
│   └── k8s/
│       ├── pubsub.yaml                       # Scopes: ["contracts-be"]
│       ├── statestore.yaml
│       └── secrets.yaml
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   └── codes.go                          # CONTRACT_NOT_FOUND, MILESTONE_NOT_COMPLETED
│   ├── logger/                               # ❌ REMOVED
│   │   └── README.md
│   ├── utils/
│   │   ├── validator.go
│   │   ├── time_calculator.go                # Time calculation utilities
│   │   └── date_utils.go                     # Date utilities
│   └── constants/
│       ├── events.go                         # ❌ REMOVED
│       └── topics.go                         # ❌ REMOVED
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   └── seed-data.sh
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── API.md
│   ├── EVENTS.md                             # 🆕 Events (contract.created, milestone.completed, timesheet.submitted, dispute.opened)
│   ├── ARCHITECTURE.md
│   ├── MIGRATIONS.md                         # 🆕
│   ├── SCHEMA.md                             # 🆕
│   ├── RUNBOOK.md                            # 🆕
│   ├── contract-lifecycle.md                 # Contract lifecycle documentation
│   ├── sow-versioning.md                     # 🆕 SOW versioning & approvals
│   └── financial-holds.md                    # 🆕 Holds policy & dispute anchors
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod                                    # 📝 UPDATED: Imports pkg/auth, platform-shared, contracts/events
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md

```

---

## 📦 **5️⃣ financial-be (Payments and Wallet Management Service)**

```
apps/be/financial-be/
│
├── cmd/
│   └── api/
│       └── main.go                           # Entry point - initializes Gin, Dapr, connects to Postgres
│
├── internal/
│   ├── domain/                               # 🏛️ Domain Layer
│   │   ├── wallet/
│   │   │   ├── entity.go                     # User wallet (UserID, Balance, Currency, Status, CreatedAt)
│   │   │   ├── balance.go                    # Balance tracking (Available, Pending, Reserved)
│   │   │   ├── currency.go                   # Multi-currency support (USD, EUR, GBP, etc.)
│   │   │   ├── errors.go                     # Wallet errors (InsufficientFunds, WalletNotFound)
│   │   │   ├── repository.go                 # WalletRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: WalletCreated, WalletUpdated, WalletClosed, FundsReserved, FundsReleased
│   │   │
│   │   ├── transaction/
│   │   │   ├── entity.go                     # Financial transactions (ID, WalletID, Amount, Type, Status, Reference)
│   │   │   ├── enums.go                      # Type (Deposit, Withdrawal, Transfer, Payment, Refund), Status (Pending, Completed, Failed), Category
│   │   │   ├── ledger.go                     # Double-entry ledger (debit/credit accounting)
│   │   │   ├── errors.go                     # Transaction errors (TransactionFailed, DuplicateTransaction)
│   │   │   ├── repository.go                 # TransactionRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: TransactionInitiated, TransactionPosted, TransactionFailed, TransactionReversed
│   │   │
│   │   ├── ledger_journal/                   # 🆕 Wallet ledger (double-entry) - immutable journal & audit
│   │   │   ├── entity.go                     # JournalEntry (ID, DebitAccount, CreditAccount, Amount, Currency, EffectiveAt, Hash, PrevHash)
│   │   │   ├── transfer.go                   # Transfers between accounts (builds 2-sided entries)
│   │   │   ├── adjustment.go                 # Manual adjustments with approvals
│   │   │   ├── audit.go                      # Audit trail (entry proofs, tamper-evidence)
│   │   │   ├── errors.go                     # Ledger errors (Imbalance, ImmutableViolation)
│   │   │   ├── repository.go                 # LedgerJournalRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: JournalEntryRecorded, JournalAdjustmentApproved, JournalAdjustmentRejected
│   │   │
│   │   ├── payment/
│   │   │   ├── entity.go                     # Payment records (ID, Amount, Currency, PayerID, PayeeID, Method, Status)
│   │   │   ├── payment_method.go             # Payment methods (CreditCard, PayPal, BankTransfer, Wallet)
│   │   │   ├── method_profile.go             # 🆕 Payout preferences per method (limits, cutoffs, fees)
│   │   │   ├── onboarding.go                 # 🆕 Onboarding KYC/KYB statuses per method/provider
│   │   │   ├── gateway.go                    # Payment gateway abstraction (Stripe, PayPal interface)
│   │   │   ├── errors.go                     # Payment errors (PaymentFailed, InvalidCard)
│   │   │   ├── repository.go                 # PaymentRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: PaymentAuthorized, PaymentCaptured, PaymentFailed, PaymentRefundInitiated
│   │   │
│   │   ├── escrow/
│   │   │   ├── entity.go                     # Escrow accounts (ID, ContractID, Amount, Status, HeldAt, ReleasedAt)
│   │   │   ├── hold.go                       # Fund holds (amount reserved until milestone completion)
│   │   │   ├── release.go                    # Fund release conditions (milestone approved, dispute resolved)
│   │   │   ├── pro_rata.go                   # 🆕 Pro-rata partial releases by milestone & refunds
│   │   │   ├── errors.go                     # Escrow errors (EscrowNotFound, InsufficientEscrow)
│   │   │   ├── repository.go                 # EscrowRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: EscrowFunded, EscrowPartiallyReleased, EscrowReleased, EscrowRefunded
│   │   │
│   │   ├── payout/
│   │   │   ├── entity.go                     # Payout requests (ID, UserID, Amount, Method, Status, RequestedAt, ProcessedAt)
│   │   │   ├── method.go                     # Payout methods (BankTransfer, PayPal, Payoneer, Wire)
│   │   │   ├── schedule.go                   # Payout scheduling (instant, daily, weekly, monthly)
│   │   │   ├── errors.go                     # Payout errors (PayoutFailed, BelowMinimum)
│   │   │   ├── repository.go                 # PayoutRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: PayoutRequested, PayoutScheduled, PayoutProcessed, PayoutFailed
│   │   │
│   │   ├── invoice/
│   │   │   ├── entity.go                     # Invoice generation (ID, ContractID, Number, Amount, DueDate, Status, PaidAt)
│   │   │   ├── line_item.go                  # Invoice line items (description, quantity, price, total)
│   │   │   ├── tax.go                        # Tax calculations (VAT, sales tax by jurisdiction)
│   │   │   ├── errors.go                     # Invoice errors (InvoiceNotFound, AlreadyPaid)
│   │   │   ├── repository.go                 # InvoiceRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: InvoiceIssued, InvoiceUpdated, InvoicePaid, InvoiceOverdue, InvoiceCanceled
│   │   │
│   │   ├── fee/
│   │   │   ├── entity.go                     # Platform fees (ID, TransactionID, Amount, Type, Rate)
│   │   │   ├── calculator.go                 # Fee calculation rules (percentage, tiered rates)
│   │   │   ├── tier.go                       # Fee tiers (based on subscription, volume, user type)
│   │   │   ├── v2/                           # 🆕 Fee engine v2 (tiered/volume, country exceptions, coupons)
│   │   │   │   ├── rules.go                  # Rules model (volume discounts, locale overrides, coupons)
│   │   │   │   ├── coupon.go                 # Coupon/promo code entity & redemption logic
│   │   │   │   ├── country_exceptions.go     # Country/regional exceptions
│   │   │   │   ├── errors.go                 # Fee v2 errors (InvalidCoupon, IneligibleTier)
│   │   │   │   └── repository.go             # FeeRulesRepository interface
│   │   │   ├── errors.go                     # Fee errors
│   │   │   ├── repository.go                 # FeeRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: FeeCalculated, FeeAdjusted, CouponApplied, CouponRevoked
│   │   │
│   │   ├── refund/
│   │   │   ├── entity.go                     # Refund processing (ID, PaymentID, Amount, Reason, Status, ProcessedAt)
│   │   │   ├── policy.go                     # Refund policies (full, partial, time limits)
│   │   │   ├── pro_rata.go                   # 🆕 Pro-rata refunds aligned to milestones
│   │   │   ├── errors.go                     # Refund errors (RefundNotAllowed, RefundExpired)
│   │   │   ├── repository.go                 # RefundRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: RefundRequested, RefundApproved, RefundDeclined, RefundProcessed
│   │   │
│   │   ├── dispute_payment/
│   │   │   ├── entity.go                     # Payment disputes (ID, PaymentID, Reason, FiledBy, Status, Resolution)
│   │   │   ├── chargeback.go                 # Chargeback handling (from credit card companies)
│   │   │   ├── errors.go                     # Dispute errors
│   │   │   ├── repository.go                 # DisputePaymentRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: PaymentDisputeOpened, ChargebackRecorded, PaymentDisputeResolved
│   │   │
│   │   ├── tax/
│   │   │   ├── entity.go                     # Tax records (ID, UserID, Year, Type, Amount, Status)
│   │   │   ├── vat_gst.go                    # 🆕 VAT/GST per locale & reverse-charge logic
│   │   │   ├── reverse_charge.go             # 🆕 Reverse-charge mechanics & validations
│   │   │   ├── forms_1099k.go                # 🆕 1099/K forms generation & thresholds
│   │   │   ├── profile_link.go               # 🆕 Link to Users-BE tax profile (TIN/VAT ID)
│   │   │   ├── form.go                       # Tax forms (W9, 1099, VAT returns)
│   │   │   ├── errors.go                     # Tax errors
│   │   │   ├── repository.go                 # TaxRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: TaxRecordCreated, TaxFormGenerated, TaxProfileLinked, VATReverseChargeApplied
│   │   │
│   │   ├── fx/                               # 🆕 Multi-currency rigor (quotes vs settlement, FX rates)
│   │   │   ├── rate.go                       # FXRate (Base, Quote, Rate, EffectiveFrom, Provider, Precision)
│   │   │   ├── quote_settlement.go           # Quote vs. settlement currency handling
│   │   │   ├── rounding.go                   # Rounding rules per currency/amount type
│   │   │   ├── errors.go                     # FX errors (RateNotFound, StaleRate)
│   │   │   ├── repository.go                 # FXRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: FXRateUpdated, FXQuoteCreated, FXSettlementCalculated
│   │   │
│   │   └── risk/                             # 🆕 Risk/holds (reserves, rolling reserves, negative balances)
│   │       ├── reserve.go                    # Reserve balances & rolling reserve schedules
│   │       ├── hold_workflow.go              # Holds/auto-release based on risk signals & time
│   │       ├── chargeback_workflow.go        # Chargeback workflows & recovery
│   │       ├── negative_balance.go           # Negative balance handling & collections
│   │       ├── errors.go                     # Risk errors
│   │       ├── repository.go                 # RiskRepository interface
│   │       └── events.go                     # 🆕 Domain events: ReserveSet, ReserveUpdated, RiskHoldPlaced, RiskHoldReleased, NegativeBalanceCreated
│   │
│   ├── application/                          # 📋 Application Layer (now with commands/queries/service/validators everywhere)
│   │   ├── eventhandler/
│   │   │   ├── billing_handler.go             # Consumes (from subscriptions-be → financial-be):
│   │   │   │                                     #   billing.invoice.exported — ingest exported invoices for charging/AR.
│   │   │   │
│   │   │   ├── admin_risk_handler.go          # Consumes (from admin-be → financial-be; Risk dashboards):
│   │   │   │                                     #   risk.hold.placed
│   │   │   │                                     #   risk.hold.released
│   │   │   │                                     #   risk.reserve.set
│   │   │   │                                     #   risk.reserve.updated
│   │   │   │                                     #   risk.chargeback.recorded
│   │   │   │                                     #   risk.chargeback.updated
│   │   │   │                                     #   risk.velocity.alert
│   │   │   │
│   │   │   ├── admin_flags_handler.go         # Consumes (from admin-be; Config & flags):
│   │   │   │                                     #   admin.feature_flag.updated
│   │   │   │                                     #   admin.threshold.updated
│   │   │   │                                     #   admin.experiment.updated
│   │   │   │
│   │   │   └── seats_billing_handler.go       # Consumes (from subscriptions-be → financial-be; Billing & seats):
│   │   │
│   │   ├── wallet/
│   │   │   ├── service.go                    # Wallet logic (Create, Deposit, Withdraw, Transfer, GetBalance)
│   │   │   ├── commands.go                   # DepositCommand, WithdrawCommand, TransferCommand
│   │   │   ├── queries.go                    # GetBalance, GetHistory, GetTransactions
│   │   │   ├── validators.go                 # 🆕 Validate currency, amounts, sufficient balance, status
│   │   │   ├── dto.go                        # WalletDTO, TransactionDTO
│   │   │   └── mapper.go                     # Wallet mappers
│   │   │
│   │   ├── transaction/
│   │   │   ├── service.go                    # Transaction logic (Create, Reverse, Reconcile)
│   │   │   ├── commands.go                   # CreateTransaction, ReverseTransaction, ReconcileTransactions
│   │   │   ├── queries.go                    # GetTransaction, ListTransactions, GetLedger
│   │   │   ├── validators.go                 # 🆕 Immutable journal checks, double-entry balance
│   │   │   ├── dto.go                        # TransactionDTO, LedgerDTO
│   │   │   ├── mapper.go                     # Transaction mappers
│   │   │   └── reconciliation.go             # Reconciliation logic (match payments with bank statements)
│   │   │
│   │   ├── ledger_journal/                   # 🆕 Ledger journal application surface
│   │   │   ├── service.go                    # AppendEntry, Transfer, Adjust, AuditTrail
│   │   │   ├── commands.go                   # AppendJournalEntry, TransferFunds, CreateAdjustment
│   │   │   ├── queries.go                    # GetEntry, ListEntries, GetAuditTrail
│   │   │   ├── validators.go                 # Entry hash chain, debits=credits, approvals
│   │   │   ├── dto.go                        # JournalEntryDTO, TransferDTO, AdjustmentDTO
│   │   │   └── mapper.go                     # Journal mappers
│   │   │
│   │   ├── payment/
│   │   │   ├── service.go                    # Payment logic (Process, Capture, Void, Refund)
│   │   │   ├── commands.go                   # ProcessPayment, CapturePayment, VoidPayment, RefundPayment
│   │   │   ├── queries.go                    # GetPayment, ListPayments
│   │   │   ├── validators.go                 # 🆕 Method eligibility, onboarding status, limits
│   │   │   ├── dto.go                        # PaymentDTO, ProcessPaymentDTO
│   │   │   ├── mapper.go                     # Payment mappers
│   │   │   ├── stripe_processor.go           # Stripe payment processor implementation
│   │   │   ├── paypal_processor.go           # PayPal payment processor implementation
│   │   │   └── processor_factory.go          # Payment processor factory (select processor by method)
│   │   │
│   │   ├── escrow/
│   │   │   ├── service.go                    # Escrow logic (Hold, Release, Refund)
│   │   │   ├── commands.go                   # HoldEscrow, ReleaseEscrow, RefundEscrow, 🆕 PartialReleaseEscrow
│   │   │   ├── queries.go                    # 🆕 GetEscrow, ListEscrows, GetEscrowHistory
│   │   │   ├── validators.go                 # 🆕 Hold coverage, release conditions, pro-rata math
│   │   │   ├── dto.go                        # EscrowDTO, HoldEscrowDTO
│   │   │   ├── mapper.go                     # Escrow mappers
│   │   │   └── pro_rata_release_manager.go   # 🆕 Dispute-driven conditional releases
│   │   │
│   │   ├── payout/
│   │   │   ├── service.go                    # Payout logic (Request, Process, Cancel, GetHistory)
│   │   │   ├── commands.go                   # RequestPayout, ProcessPayout, CancelPayout
│   │   │   ├── queries.go                    # GetPayouts, GetPayoutHistory
│   │   │   ├── validators.go                 # 🆕 Method limits, compliance holds, min amounts
│   │   │   ├── dto.go                        # PayoutDTO, RequestPayoutDTO
│   │   │   ├── mapper.go                     # Payout mappers
│   │   │   └── batch_processor.go            # Batch payout processor (process multiple payouts together)
│   │   │
│   │   ├── invoice/
│   │   │   ├── service.go                    # Invoice logic (Generate, Send, MarkPaid, Cancel)
│   │   │   ├── commands.go                   # GenerateInvoice, SendInvoice, MarkInvoicePaid
│   │   │   ├── queries.go                    # GetInvoice, ListInvoices
│   │   │   ├── validators.go                 # 🆕 Line totals, tax rounding, currency consistency
│   │   │   ├── dto.go                        # InvoiceDTO, GenerateInvoiceDTO
│   │   │   ├── mapper.go                     # Invoice mappers
│   │   │   ├── generator.go                  # Invoice PDF generation
│   │   │   └── tax_calculator.go             # Tax calculation (VAT, sales tax)
│   │   │
│   │   ├── fee/
│   │   │   ├── service.go                    # Fee logic (Calculate, Apply, Waive)
│   │   │   ├── calculator.go                 # Fee calculation (percentage, tiered, flat)
│   │   │   ├── validators.go                 # 🆕 Fee caps, non-negative, eligibility
│   │   │   ├── dto.go                        # FeeDTO
│   │   │   └── rules_engine.go               # Fee rules engine (apply rules based on user type, volume)
│   │   │
│   │   ├── fee_v2/                           # 🆕 App layer for Fee engine v2
│   │   │   ├── service.go                    # Calculate with tiers, volume, country exceptions, coupons
│   │   │   ├── commands.go                   # ApplyCoupon, RevokeCoupon, UpdateFeeRules
│   │   │   ├── queries.go                    # GetEffectiveFee, GetUserFeeTier, GetCoupon
│   │   │   ├── validators.go                 # Coupon validity, stacking rules, locale overrides
│   │   │   ├── dto.go                        # FeeV2DTO, CouponDTO, FeeRuleDTO
│   │   │   └── mapper.go                     # Fee v2 mappers
│   │   │
│   │   ├── refund/
│   │   │   ├── service.go                    # Refund logic (Process, Cancel, GetHistory)
│   │   │   ├── commands.go                   # ProcessRefund, CancelRefund, 🆕 ProcessPartialRefund
│   │   │   ├── queries.go                    # 🆕 GetRefund, ListRefunds
│   │   │   ├── validators.go                 # 🆕 Pro-rata vs eligibility, time windows
│   │   │   ├── dto.go                        # RefundDTO
│   │   │   └── mapper.go                     # Refund mappers
│   │   │
│   │   ├── tax/
│   │   │   ├── service.go                    # Tax logic (Calculate, Generate forms, File)
│   │   │   ├── form_generator.go             # Generate tax forms (W9, 1099, 1099-K)
│   │   │   ├── commands.go                   # 🆕 SyncTaxProfileFromUsersBE, Generate1099K
│   │   │   ├── queries.go                    # 🆕 GetTaxProfile, GetVATRate, GetReverseChargeEligibility
│   │   │   ├── validators.go                 # 🆕 VAT ID formats, threshold checks
│   │   │   ├── dto.go                        # TaxDTO
│   │   │   └── mapper.go                     # Tax mappers
│   │   │
│   │   ├── fx/
│   │   │   ├── service.go                    # Quote→settlement conversion, fetch/apply FX rates
│   │   │   ├── commands.go                   # 🆕 UpsertFXRate, SetRoundingRule
│   │   │   ├── queries.go                    # 🆕 GetFXRate, ConvertAmount, GetEffectiveRateAt
│   │   │   ├── validators.go                 # 🆕 Effective timestamps, precision, currency pairs
│   │   │   ├── dto.go                        # FXRateDTO, ConversionRequestDTO
│   │   │   └── mapper.go                     # FX mappers
│   │   │
│   │   └── risk/
│   │       ├── service.go                    # Holds, reserves, chargebacks, negative balances
│   │       ├── commands.go                   # 🆕 CreateReserve, ReleaseReserve, PlaceHold, RemoveHold
│   │       ├── queries.go                    # 🆕 GetReserves, GetHolds, GetRiskScore
│   │       ├── validators.go                 # 🆕 Rolling reserve schedules, chargeback states
│   │       ├── dto.go                        # RiskDTO, ReserveDTO, HoldDTO
│   │       └── mapper.go                     # Risk mappers
│   │
│   ├── infrastructure/                       # 🔧 Infrastructure Layer
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection
│   │   │       ├── transaction.go            # Transaction helpers
│   │   │       ├── migrations.go             # 📝 UPDATED: Auto-migration with version tracking
│   │   │       ├── version.go                # 🆕 Schema version tracking
│   │   │       ├── safety.go                 # 🆕 Pre-migration safety checks
│   │   │       ├── wallet_repository.go      # WalletRepository implementation
│   │   │       ├── transaction_repository.go # TransactionRepository implementation
│   │   │       ├── ledger_journal_repository.go # 🆕 LedgerJournalRepository implementation
│   │   │       ├── payment_repository.go     # PaymentRepository implementation
│   │   │       ├── escrow_repository.go      # EscrowRepository implementation
│   │   │       ├── payout_repository.go      # PayoutRepository implementation
│   │   │       ├── invoice_repository.go     # InvoiceRepository implementation
│   │   │       ├── fee_repository.go         # FeeRepository implementation
│   │   │       ├── fee_rules_repository.go   # 🆕 FeeRulesRepository implementation (v2)
│   │   │       ├── refund_repository.go      # RefundRepository implementation
│   │   │       ├── dispute_payment_repository.go # DisputePaymentRepository implementation
│   │   │       ├── tax_repository.go         # TaxRepository implementation
│   │   │       ├── fx_repository.go          # 🆕 FXRepository implementation
│   │   │       ├── risk_repository.go        # 🆕 RiskRepository implementation
│   │   │       └── payment_schedule_repository.go # PaymentScheduleRepository implementation
│   │   │
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go             # Redis connection
│   │   │       ├── wallet_cache.go           # Wallet balance caching
│   │   │       ├── rate_cache.go             # Exchange rate caching
│   │   │       └── risk_cache.go             # 🆕 Holds/reserve snapshots
│   │   │
│   │   ├── clients/
│   │   │   ├── users/                        # 🆕 Users-BE client (tax profiles / KYC)
│   │   │   │   └── client.go                 # Fetch tax profiles & KYC statuses
│   │   │   └── coupons/                      # 🆕 Coupon/Promo validation backend (if separate)
│   │   │       └── client.go                 # Validate & redeem promo codes
│   │   │
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go               # 📝 UPDATED: Uses platform-shared/inbox
│   │   │       ├── producer.go               # 📝 UPDATED: Uses platform-shared/outbox
│   │   │       ├── topics.go                 # 📝 UPDATED: From contracts/events (payment.processed, escrow.released, payout.processed)
│   │   │       └── scram.go                  # SCRAM authentication
│   │   │
│   │   ├── payment_gateway/
│   │   │   ├── stripe/
│   │   │   │   ├── client.go                 # Stripe API client
│   │   │   │   ├── webhook_handler.go        # Stripe webhook handler (payment.succeeded, charge.failed)
│   │   │   │   └── mapper.go                 # Stripe event mapper
│   │   │   ├── paypal/
│   │   │   │   ├── client.go                 # PayPal API client
│   │   │   │   ├── webhook_handler.go        # PayPal webhook handler
│   │   │   │   └── mapper.go                 # PayPal event mapper
│   │   │   └── factory.go                    # Payment gateway factory
│   │   │
│   │   └── pdf/
│   │       └── generator.go                  # PDF invoice generation (using library like wkhtmltopdf)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── wallet_handler.go         # Wallet handlers (GET /wallets/:id, POST /wallets/:id/deposit)
│   │       │   ├── transaction_handler.go    # Transaction handlers (GET /transactions)
│   │       │   ├── ledger_journal_handler.go # 🆕 Journal handlers (POST /ledger/entry, GET /ledger/audit)
│   │       │   ├── payment_handler.go        # Payment handlers (POST /payments, GET /payments/:id)
│   │       │   ├── escrow_handler.go         # Escrow handlers (POST /escrow/hold, POST /escrow/release)
│   │       │   ├── payout_handler.go         # Payout handlers (POST /payouts, GET /payouts/:id)
│   │       │   ├── invoice_handler.go        # Invoice handlers (POST /invoices, GET /invoices/:id)
│   │       │   ├── fee_handler.go            # Fee handlers (GET /fees/calculate)
│   │       │   ├── fee_v2_handler.go         # 🆕 Fee v2 endpoints (GET /fees/v2/effective, POST /fees/v2/coupon)
│   │       │   ├── refund_handler.go         # Refund handlers (POST /refunds)
│   │       │   ├── tax_handler.go            # Tax handlers (GET /tax/forms)
│   │       │   ├── fx_handler.go             # 🆕 FX handlers (GET /fx/rate, POST /fx/rate)
│   │       │   ├── risk_handler.go           # 🆕 Risk/holds handlers (POST /risk/reserve, GET /risk/holds)
│   │       │   ├── bank_account_handler.go   # Bank account handlers (POST /bank-accounts)
│   │       │   ├── webhook_handler.go        # Webhook handlers (POST /webhooks/stripe, POST /webhooks/paypal)
│   │       │   └── health_handler.go         # Health check
│   │       │
│   │       ├── routes/                       # 🆕 Route registrars for each handler group
│   │       │   ├── wallet_routes.go          # /wallets/*
│   │       │   ├── transaction_routes.go     # /transactions/*
│   │       │   ├── ledger_journal_routes.go  # /ledger/*
│   │       │   ├── payment_routes.go         # /payments/*
│   │       │   ├── escrow_routes.go          # /escrow/*
│   │       │   ├── payout_routes.go          # /payouts/*
│   │       │   ├── invoice_routes.go         # /invoices/*
│   │       │   ├── fee_routes.go             # /fees/*
│   │       │   ├── fee_v2_routes.go          # /fees/v2/*
│   │       │   ├── refund_routes.go          # /refunds/*
│   │       │   ├── tax_routes.go             # /tax/*
│   │       │   ├── fx_routes.go              # /fx/*
│   │       │   └── risk_routes.go            # /risk/*
│   │       │
│   │       ├── middleware/                   # 📝 UPDATED: Uses platform-shared
│   │       │   ├── auth.go                   # 📝 UPDATED: Uses pkg/auth
│   │       │   ├── rbac.go                   # 📝 UPDATED: Uses pkg/auth
│   │       │   ├── cors.go                   # 📝 UPDATED: Uses platform-shared/ginx
│   │       │   ├── rate_limit.go             # Rate limiting
│   │       │   └── idempotency.go            # 🆕 Uses platform-shared/idempotency (critical for payments!)
│   │       │
│   │       ├── responses/
│   │       │   └── README.md                 # 📝 Points to platform-shared/httpx
│   │       │
│   │       └── router.go                     # 📝 UPDATED: Uses platform-shared/ginx
│   │
│   └── config/                               # 🆕 MEDIUM
│       ├── schema.go                         # 🆕 Typed Config (includes Stripe/PayPal keys)
│       ├── loader.go                         # 🆕 Viper loader
│       └── docs/
│           └── CONFIGURATION.md              # 🆕 Config docs
│
├── config/                                   # 🆕 MEDIUM
│   ├── default.yaml
│   ├── dev.yaml
│   └── prod.yaml
│
├── dapr/                                     # 🆕 MEDIUM
│   ├── local/
│   │   ├── pubsub.yaml
│   │   └── statestore.yaml
│   └── k8s/
│       ├── pubsub.yaml                       # Scopes: ["financial-be"]
│       ├── statestore.yaml
│       └── secrets.yaml
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   ├── codes.go                          # INSUFFICIENT_FUNDS, PAYMENT_FAILED
│   │   └── payment_errors.go                 # Payment-specific error types
│   ├── logger/                               # ❌ REMOVED
│   │   └── README.md
│   ├── utils/
│   │   ├── validator.go
│   │   ├── currency.go                       # Currency conversion utilities
│   │   ├── decimal.go                        # Decimal math for money (avoid floating point errors)
│   │   └── encryption.go                     # Encryption utilities (for sensitive data like card numbers)
│   └── constants/
│       ├── events.go                         # ❌ REMOVED
│       ├── topics.go                         # ❌ REMOVED
│       ├── currencies.go                     # Currency constants (USD, EUR, GBP)
│       └── payment_methods.go                # Payment method constants
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml                      # Contains Stripe/PayPal API keys
│       ├── hpa.yaml
│       ├── pdb.yaml
│       ├── networkpolicy.yaml                # Extra security for financial service
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   ├── seed-data.sh
│   └── reconciliation.sh                     # Daily reconciliation script
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── API.md
│   ├── EVENTS.md                             # 🆕 Events (payment.processed, escrow.held, escrow.released, payout.processed)
│   ├── ARCHITECTURE.md
│   ├── MIGRATIONS.md                         # 🆕
│   ├── SCHEMA.md                             # 🆕 Domain schemas (journal, fees v2, fx)
│   ├── RUNBOOK.md                            # 🆕
│   ├── payment-flows.md                      # Payment flow documentation
│   ├── escrow-system.md                      # Escrow system documentation
│   ├── fee-structure.md                      # Platform fee structure
│   ├── fee-v2.md                             # 🆕 Fee engine v2 rules & examples
│   ├── fx-and-rounding.md                    # 🆕 FX, effective timestamps, rounding
│   ├── risk-holds.md                         # 🆕 Reserves, holds, chargebacks
│   └── compliance.md                         # PCI DSS & tax compliance notes
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod                                    # 📝 UPDATED: Imports pkg/auth, platform-shared, contracts/events
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md

```

---
# Skillsier Platform - Remaining Microservices Folder Structures
## Services 6-11 with Full Comments (Old + New) ✨

---

## 📦 **6️⃣ communications-be (Real-time Messaging and Notifications Service)**

```
apps/be/communications-be/
│
├── cmd/
│   └── api/
│       └── main.go                           # 📝 UPDATED: Application entry point - initializes Gin, Dapr, connects to Postgres (now uses loadConfig from internal/config, platform-shared/logging)
│
├── internal/
│   ├── domain/                               # 🛍️ Domain Layer - Business logic & entities
│   │   ├── conversation/
│   │   │   ├── entity.go                     # Chat conversations (ID, Participants, Type, Status, CreatedAt, UpdatedAt)
│   │   │   ├── participant.go                # Conversation participants (ConversationID, UserID, Role, JoinedAt, LastReadAt)
│   │   │   ├── settings.go                   # Conversation settings (mute, pin, archive flags per user)
│   │   │   ├── typing_indicator.go           # Typing indicators (ConversationID, UserID, TypingAt, ExpiresAt)
│   │   │   ├── errors.go                     # Domain errors (ConversationNotFound, ParticipantNotFound, UnauthorizedAccess)
│   │   │   ├── repository.go                 # ConversationRepository interface (Create, FindByID, AddParticipant, RemoveParticipant)
│   │   │   └── events.go                     # 🆕 Domain events: ConversationCreated, ParticipantAdded, ParticipantRemoved, ConversationMuted, ConversationArchived
│   │   │
│   │   ├── message/
│   │   │   ├── entity.go                     # Chat messages (ID, ConversationID, SenderID, Content, Type, SentAt, EditedAt, DeletedAt)
│   │   │   ├── attachment.go                 # Message attachments (MessageID, FileURL, FileType, FileName, FileSize, ThumbnailURL)
│   │   │   ├── read_receipt.go               # Read receipts (MessageID, UserID, ReadAt)
│   │   │   ├── reaction.go                   # Message reactions (MessageID, UserID, Emoji, ReactedAt)
│   │   │   ├── mention.go                    # User mentions (@username) (MessageID, MentionedUserID, Position)
│   │   │   ├── errors.go                     # Message errors (MessageNotFound, InvalidContent, MessageTooLong)
│   │   │   └── repository.go                 # MessageRepository interface (Create, Update, Delete, FindByConversation, MarkAsRead)
│   │   │   └── events.go                     # 🆕 Domain events: MessageSent, MessageEdited, MessageDeleted, MessageRead, AttachmentAdded, ReactionAdded
│   │   │
│   │   ├── notification/
│   │   │   ├── entity.go                     # All notifications (ID, UserID, Type, Title, Body, Data JSON, Priority, ReadAt, CreatedAt)
│   │   │   ├── enums.go                      # Type (JobAlert, ProposalUpdate, BidAlert, ContractUpdate, Payment, Review, Message, System), Priority (Low, Medium, High, Urgent), Category (Transactional, Marketing, System)
│   │   │   ├── preferences.go                # User notification preferences (Email, SMS, Push, InApp enabled/disabled per notification type)
│   │   │   ├── settings.go                   # Notification settings per type (instant delivery, daily digest, weekly summary, muted)
│   │   │   ├── quiet_hours.go                # 🆕 Quiet hours per user (start, end, timezone)
│   │   │   ├── errors.go                     # Notification errors (NotificationNotFound, InvalidType, PreferencesNotSet)
│   │   │   └── repository.go                 # NotificationRepository interface (Create, MarkAsRead, Delete, FindByUser, GetUnreadCount)
│   │   │   └── events.go                     # 🆕 Domain events: NotificationCreated, NotificationUpdated, NotificationRead, NotificationDeleted
│   │   │
│   │   ├── in_app_notification/
│   │   │   ├── entity.go                     # 🆕 In-app real-time notifications (ID, NotificationID, UserID, DisplayedAt, DismissedAt, ClickedAt)
│   │   │   ├── badge_count.go                # 🆕 Unread notification badges (UserID, Category, UnreadCount, LastUpdatedAt)
│   │   │   ├── group.go                      # 🆕 Notification grouping (GroupID, NotificationIDs, GroupType, Summary)
│   │   │   ├── action.go                     # 🆕 Notification actions/CTAs (NotificationID, ActionType, ActionURL, Label)
│   │   │   ├── errors.go                     # 🆕 In-app notification errors
│   │   │   └── repository.go                 # 🆕 InAppNotificationRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: InAppNotificationDisplayed, InAppNotificationDismissed, InAppNotificationClicked, BadgeCountUpdated
│   │   │
│   │   ├── notification_template/
│   │   │   ├── entity.go                     # Notification templates (ID, Name, Type, Subject, Body, Variables, IsActive)
│   │   │   ├── variable.go                   # Template variables ({{user_name}}, {{job_title}}, {{amount}}, {{date}})
│   │   │   ├── localization.go               # Multi-language templates (TemplateID, Language, TranslatedSubject, TranslatedBody)
│   │   │   ├── errors.go                     # Template errors (TemplateNotFound, InvalidVariable, MissingTranslation)
│   │   │   └── repository.go                 # TemplateRepository interface (FindByType, Render, GetTranslation)
│   │   │   └── events.go                     # 🆕 Domain events: TemplateCreated, TemplateUpdated, TemplateLocalized, TemplateDeactivated
│   │   │
│   │   ├── email/
│   │   │   ├── entity.go                     # Email records (ID, RecipientID, RecipientEmail, Subject, Body, TemplateID, Status, SentAt, DeliveredAt)
│   │   │   ├── template.go                   # Email templates (HTML/Text templates with CSS inlining support)
│   │   │   ├── batch.go                      # Batch email sending (BatchID, Recipients, Subject, Body, Status, ScheduledFor)
│   │   │   ├── errors.go                     # Email errors (InvalidEmail, SendFailed, TemplateRenderFailed)
│   │   │   └── repository.go                 # EmailRepository interface (Create, FindByRecipient, GetDeliveryStatus)
│   │   │   └── events.go                     # 🆕 Domain events: EmailQueued, EmailSent, EmailDelivered, EmailBounced, EmailFailed
│   │   │
│   │   ├── notification_queue/
│   │   │   ├── entity.go                     # 🆕 Queued notifications (ID, NotificationID, Priority, ScheduledFor, Status, Retries, CreatedAt)
│   │   │   ├── priority_queue.go             # 🆕 Priority-based queue (urgent notifications processed first)
│   │   │   ├── errors.go                     # 🆕 Queue errors
│   │   │   └── repository.go                 # 🆕 NotificationQueueRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: NotificationEnqueued, NotificationDequeued, NotificationRetryScheduled, NotificationDeadLettered
│   │   │
│   │   ├── delivery_log/
│   │   │   ├── entity.go                     # 🆕 Delivery tracking (ID, NotificationID, Channel, Status, DeliveredAt, FailureReason, RetryCount)
│   │   │   ├── status.go                     # 🆕 Delivery status (Pending, Sent, Delivered, Failed, Bounced)
│   │   │   ├── errors.go                     # 🆕 Delivery errors
│   │   │   └── repository.go                 # 🆕 DeliveryLogRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: DeliveryLogged, DeliveryFailed, DeliveryBounced
│   │   │
│   │   ├── unsubscribe/
│   │   │   ├── entity.go                     # Unsubscribe management (UserID, NotificationType, UnsubscribedAt, Reason)
│   │   │   ├── errors.go                     # Unsubscribe errors
│   │   │   └── repository.go                 # UnsubscribeRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: Unsubscribed, Resubscribed
│   │   │
│   │   ├── online_status/
│   │   │   ├── entity.go                     # User online/offline status (UserID, Status, LastSeenAt, DeviceType)
│   │   │   ├── errors.go                     # Status errors
│   │   │   └── repository.go                 # OnlineStatusRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: UserOnline, UserOffline, LastSeenUpdated
│   │   │
│   │   ├── message_flag/
│   │   │   ├── entity.go                     # Flagged messages (ID, MessageID, ReporterID, Reason, Status, ReviewedAt, ReviewedBy)
│   │   │   ├── reason.go                     # Flag reasons enum (Spam, Harassment, Inappropriate, Scam, Violence)
│   │   │   ├── status.go                     # Flag status (Pending, UnderReview, Resolved, Dismissed)
│   │   │   ├── errors.go                     # Flag errors
│   │   │   └── repository.go                 # MessageFlagRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: MessageFlagSubmitted, MessageFlagResolved, MessageFlagDismissed
│   │   │
│   │   ├── system_message/                   # 🆕 System messages domain (events → system feed)
│   │   │   ├── entity.go                     # SystemMessage (ID, ConversationID, Type, PayloadJSON, CreatedAt)
│   │   │   ├── types.go                      # Types (MilestoneFunded, MilestoneReleased, Approval, DisputeOpened, DisputeClosed)
│   │   │   ├── errors.go                     # System message errors
│   │   │   └── repository.go                 # SystemMessageRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: SystemMessageCreated, SystemMessageBroadcasted
│   │   │
│   │   ├── call/                             # 🆕 Voice/Video & scheduling hooks
│   │   │   ├── entity.go                     # Call (ID, ConversationID, OrganizerID, Link, Provider, StartsAt, EndsAt, Status, Timezone)
│   │   │   ├── provider.go                   # Providers (Jitsi/Zoom/Meet) abstraction
│   │   │   ├── errors.go                     # Call errors (InvalidTimeWindow, ProviderError)
│   │   │   └── repository.go                 # CallRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: CallScheduled, CallStarted, CallEnded, CallCanceled
│   │   │
│   │   └── calendar_invite/                  # 🆕 Calendar invites for scheduling
│   │       ├── entity.go                     # CalendarInvite (CallID, UserID, ICal, SentAt, Status)
│   │       ├── errors.go                     # Invite errors
│   │       └── repository.go                 # CalendarInviteRepository interface
│   │       └── events.go                     # 🆕 Domain events: CalendarInviteSent, CalendarInviteAccepted, CalendarInviteDeclined, CalendarInviteBounced
│   │
│   ├── application/                          # 📋 Application Layer - Use cases & orchestration
│   │   ├── eventhandler/
│   │   │   ├── user_handler.go                # Consumes: user.created → send welcome (emails/in-app). 
│   │   │   │                                     # Ref: Event Consumption Matrix lists communications-be as consumer of user.created. 
│   │   │   │
│   │   │   ├── job_handler.go                 # Consumes: job.posted → notify matching freelancers.
│   │   │   │                                     # Ref: Matrix shows communications-be consumes job.posted. 
│   │   │   │
│   │   │   ├── proposal_handler.go            # Consumes: proposal.submitted → notify client.
│   │   │   │                                     # Ref: Matrix shows communications-be consumes proposal.submitted. 
│   │   │   │
│   │   │   ├── contract_handler.go            # Consumes: contract.created → notify both parties.
│   │   │   │                                     # Ref: Matrix shows communications-be consumes contract.created. 
│   │   │   │
│   │   │   ├── payment_handler.go             # Consumes: payment.processed → send receipt.
│   │   │   │                                     # Ref: Matrix shows communications-be consumes payment.processed. 
│   │   │   │
│   │   │   ├── review_handler.go              # Consumes: review.submitted → notify reviewee;
│   │   │   │                                     #           review.double_blind.window.opened|closed → send reminders (from reviews-be).
│   │   │   │                                     # Refs: Matrix shows communications-be consumes review.submitted; 
│   │   │   │                                     #       reviews-be publishes double_blind window events to communications-be. 
│   │   │   │
│   │   │   ├── admin_case_handler.go          # Consumes: admin.case.user_report.created|triaged|actioned|dismissed → notify subjects as needed; 
│   │   │   │                                     #           admin.ticket.note.added (if notify).
│   │   │   │                                     # Ref: Added-events doc routes these to communications-be. 
│   │   │   │
│   │   │   ├── delivery_logger_handler.go     # Emits audit on delivery completion (comm.delivery.logged) for admin-be audit streams.
│   │   │   │                                     # Ref: Added-events doc notes audit trail feed from communications-be. 
│   │   │   │
│   │   │   └── partitioning_notes.go          # (comments only) Partition keys used by handlers:
│   │   │                                         #   user_id (user events), job_id (job events), contract_id (contract),
│   │   │                                         #   transaction_id/contract_id (payments), conversation_id (message topics).
│   │   │                                         # Ref: Partition Key Strategy. 
│   │   ├── conversation/
│   │   │   ├── service.go                    # Conversation business logic (Create, Archive, Mute, Delete conversations)
│   │   │   ├── commands.go                   # CreateConversation, ArchiveConversation, MuteConversation, DeleteConversation
│   │   │   ├── queries.go                    # GetConversation, ListConversations, SearchConversations
│   │   │   ├── dto.go                        # ConversationDTO, CreateConversationDTO, ConversationListDTO
│   │   │   ├── mapper.go                     # Entity-DTO mapping
│   │   │   └── validators.go                 # Input validation
│   │   │
│   │   ├── message/
│   │   │   ├── service.go                    # Message business logic (Send, Edit, Delete, React to messages)
│   │   │   ├── commands.go                   # SendMessage, EditMessage, DeleteMessage, ReactToMessage, MarkAsRead
│   │   │   ├── queries.go                    # GetMessages, SearchMessages, GetUnreadCount
│   │   │   ├── dto.go                        # MessageDTO, SendMessageDTO, MessageListDTO
│   │   │   ├── mapper.go                     # Message mappers
│   │   │   ├── validators.go                 # Message validation (content length, attachment size)
│   │   │   └── realtime_service.go           # 🆕 WebSocket handling (broadcast messages, typing indicators, presence)
│   │   │
│   │   ├── notification/
│   │   │   ├── service.go                    # Notification business logic (Send, MarkAsRead, Delete, ClearAll)
│   │   │   ├── commands.go                   # SendNotification, MarkAsRead, DeleteNotification, ClearAllNotifications
│   │   │   ├── queries.go                    # GetNotifications, GetUnreadCount, GetNotificationHistory
│   │   │   ├── dto.go                        # NotificationDTO, SendNotificationDTO, NotificationListDTO
│   │   │   ├── mapper.go                     # Notification mappers
│   │   │   ├── validators.go                 # Notification validation
│   │   │   ├── orchestrator.go               # 🆕 Multi-channel orchestration (send to email, push, in-app simultaneously)
│   │   │   ├── preferences_service.go        # 🆕 Manage user preferences (update notification settings per type)
│   │   │   └── aggregator.go                 # 🆕 Aggregate similar notifications (group by type, time window)
│   │   │
│   │   ├── notification_preferences/         # 🆕 Explicit app module for per-channel prefs / digests / quiet hours
│   │   │   ├── service.go                    # UpdatePreferences, GetPreferences, SetQuietHours, SetDigestSchedule
│   │   │   ├── commands.go                   # UpdatePreferences, SetQuietHours, SetDigestSchedule
│   │   │   ├── queries.go                    # GetPreferences, GetEffectivePreferences
│   │   │   ├── validators.go                 # Channel validation, time windows, timezone safety
│   │   │   ├── dto.go                        # NotificationPreferencesDTO, QuietHoursDTO, DigestScheduleDTO
│   │   │   └── mapper.go                     # Preferences mappers
│   │   │
│   │   ├── in_app_notification/
│   │   │   ├── service.go                    # 🆕 In-app notification logic
│   │   │   ├── commands.go                   # 🆕 PushInAppNotification, DismissInAppNotification, ClickInAppAction
│   │   │   ├── queries.go                    # 🆕 GetInAppNotifications, GetBadgeCount
│   │   │   ├── validators.go                 # 🆕 Action validation, throttling
│   │   │   ├── real_time_sender.go           # 🆕 Real-time notification delivery via WebSocket/SSE
│   │   │   ├── badge_manager.go              # 🆕 Badge count management (calculate, update, reset counts)
│   │   │   ├── grouping_engine.go            # 🆕 Group similar notifications
│   │   │   ├── dto.go                        # 🆕 InAppNotificationDTO
│   │   │   └── mapper.go                     # 🆕 In-app notification mappers
│   │   │
│   │   ├── email/
│   │   │   ├── service.go                    # Email business logic (Send, SendBatch, GetDeliveryStatus)
│   │   │   ├── commands.go                   # 🆕 SendEmail, SendEmailBatch
│   │   │   ├── queries.go                    # 🆕 GetEmailStatus, ListEmailsForUser
│   │   │   ├── validators.go                 # 🆕 Email address format, batch sizes, template vars
│   │   │   ├── sender.go                     # Email sending logic (via WildDuck SMTP)
│   │   │   ├── template_renderer.go          # Render email templates (inject variables, compile HTML)
│   │   │   ├── batch_sender.go               # Batch email sending (queue, rate limit, send in batches)
│   │   │   ├── dto.go                        # EmailDTO, SendEmailDTO, BatchEmailDTO
│   │   │   ├── mapper.go                     # Email mappers
│   │   │   └── wildduck_client.go            # 🆕 WildDuck SMTP integration (connect, authenticate, send)
│   │   │
│   │   ├── template/
│   │   │   ├── service.go                    # Template business logic (Create, Update, Render templates)
│   │   │   ├── commands.go                   # 🆕 CreateTemplate, UpdateTemplate
│   │   │   ├── queries.go                    # 🆕 GetTemplate, ListTemplates
│   │   │   ├── validators.go                 # 🆕 Variable whitelisting, translation presence
│   │   │   ├── renderer.go                   # Template rendering engine (replace variables, compile)
│   │   │   ├── variable_injector.go          # Inject dynamic variables into templates
│   │   │   ├── dto.go                        # TemplateDTO, RenderTemplateDTO
│   │   │   └── mapper.go                     # Template mappers
│   │   │
│   │   ├── online_status/
│   │   │   ├── service.go                    # Online status logic (Update status, GetStatus, GetOnlineUsers)
│   │   │   ├── commands.go                   # 🆕 SetOnline, SetAway, SetBusy, SetOffline
│   │   │   ├── queries.go                    # 🆕 GetUserStatus, GetOnlineUsers
│   │   │   ├── validators.go                 # 🆕 TTL bounds, allowed transitions
│   │   │   ├── tracker.go                    # Track online/offline status (heartbeat mechanism)
│   │   │   ├── presence_manager.go           # Manage user presence (online, away, busy, offline)
│   │   │   └── dto.go                        # OnlineStatusDTO
│   │   │
│   │   ├── flag/
│   │   │   ├── service.go                    # Flag business logic (FlagMessage, UnflagMessage, ReviewFlag)
│   │   │   ├── commands.go                   # FlagMessage, UnflagMessage, ResolveFlag
│   │   │   ├── queries.go                    # 🆕 GetFlags, GetFlag
│   │   │   ├── validators.go                 # 🆕 Reason whitelist, reviewer role checks
│   │   │   ├── dto.go                        # FlagDTO, FlagMessageDTO
│   │   │   └── mapper.go                     # Flag mappers
│   │   │
│   │   ├── system_message/                   # 🆕 System messages app layer
│   │   │   ├── service.go                    # PublishSystemMessage (milestones funded/released, approvals, disputes)
│   │   │   ├── commands.go                   # CreateSystemMessage
│   │   │   ├── queries.go                    # GetSystemMessagesForConversation
│   │   │   ├── validators.go                 # Allowed types, payload schema
│   │   │   ├── dto.go                        # SystemMessageDTO
│   │   │   └── mapper.go                     # System message mappers
│   │   │
│   │   ├── call/                             # 🆕 Voice/Video & scheduling hooks
│   │   │   ├── service.go                    # CreateCallLink, CancelCall, RescheduleCall
│   │   │   ├── commands.go                   # CreateCall, CancelCall, RescheduleCall
│   │   │   ├── queries.go                    # GetCall, ListCallsForConversation
│   │   │   ├── validators.go                 # Timezone-safe windows, overlaps, provider limits
│   │   │   ├── dto.go                        # CallDTO
│   │   │   └── mapper.go                     # Call mappers
│   │   │
│   │   ├── calendar_invite/                  # 🆕 Calendar invites surface
│   │   │   ├── service.go                    # GenerateICal, SendInvites, UpdateInviteStatus
│   │   │   ├── commands.go                   # SendCalendarInvite
│   │   │   ├── queries.go                    # GetInvitesForCall
│   │   │   ├── validators.go                 # Email, timezone, iCal bounds
│   │   │   ├── dto.go                        # CalendarInviteDTO
│   │   │   └── mapper.go                     # Invite mappers
│   │   
│   │
│   ├── infrastructure/                       # 🔧 Infrastructure Layer - External concerns
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection setup (DSN from config, connection pooling)
│   │   │       ├── transaction.go            # Transaction helpers (Begin, Commit, Rollback, WithTransaction wrapper)
│   │   │       ├── migrations.go             # 📝 UPDATED: Auto-migration logic (now with version tracking, GORM AutoMigrate for all tables)
│   │   │       ├── version.go                # 🆕 Schema version tracking (SchemaVersion table, RecordMigration function)
│   │   │       ├── safety.go                 # 🆕 Pre-migration safety checks (environment validation, disk space check)
│   │   │       ├── conversation_repository.go # ConversationRepository implementation with GORM
│   │   │       ├── message_repository.go     # MessageRepository implementation
│   │   │       ├── notification_repository.go # NotificationRepository implementation
│   │   │       ├── in_app_notification_repository.go # 🆕 InAppNotificationRepository implementation
│   │   │       ├── template_repository.go    # TemplateRepository implementation
│   │   │       ├── email_repository.go       # EmailRepository implementation
│   │   │       ├── notification_queue_repository.go # 🆕 NotificationQueueRepository implementation
│   │   │       ├── delivery_log_repository.go # 🆕 DeliveryLogRepository implementation
│   │   │       ├── unsubscribe_repository.go # UnsubscribeRepository implementation
│   │   │       ├── online_status_repository.go # OnlineStatusRepository implementation
│   │   │       ├── system_message_repository.go # 🆕 SystemMessageRepository implementation
│   │   │       ├── call_repository.go        # 🆕 CallRepository implementation
│   │   │       └── calendar_invite_repository.go # 🆕 CalendarInviteRepository implementation
│   │   │
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go             # Redis connection setup (connection pooling, retry logic)
│   │   │       ├── conversation_cache.go     # Conversation caching (Get, Set, Invalidate with TTL)
│   │   │       ├── online_status_cache.go    # 🆕 Cache online status (fast reads for presence indicator)
│   │   │       ├── typing_indicator_cache.go # 🆕 Cache typing indicators (ephemeral data with short TTL)
│   │   │       ├── notification_cache.go     # 🆕 Cache unread counts (fast badge updates)
│   │   │       └── presence_cache.go         # 🆕 User presence caching
│   │   │
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go               # 📝 UPDATED: Kafka consumer (now uses platform-shared/inbox for message deduplication)
│   │   │       ├── producer.go               # 📝 UPDATED: Kafka producer (now uses platform-shared/outbox for reliable publishing)
│   │   │       ├── topics.go                 # 📝 UPDATED: Topic constants (now imported from contracts/events - message.sent, notification.delivered)
│   │   │       └── scram.go                  # SCRAM authentication for Kafka (SASL/SCRAM-SHA-256)
│   │   │
│   │   ├── realtime/
│   │   │   ├── websocket/
│   │   │   │   ├── hub.go                    # 🆕 WebSocket hub - connection manager (manages all active WS connections, broadcast to users/rooms)
│   │   │   │   ├── client.go                 # 🆕 WebSocket client (represents a single connection, read/write goroutines)
│   │   │   │   ├── handler.go                # 🆕 WebSocket message handler (handle incoming WS messages, route to services)
│   │   │   │   ├── broadcaster.go            # 🆕 Broadcast to multiple clients (send to all, send to room, send to user)
│   │   │   │   └── room.go                   # 🆕 WebSocket rooms/channels (group connections by conversation ID)
│   │   │   └── sse/
│   │   │       ├── handler.go                # 🆕 Server-Sent Events handler (alternative to WebSocket for one-way real-time updates)
│   │   │       └── stream.go                 # 🆕 SSE stream management (maintain open connections, send events)
│   │   │
│   │   ├── email/
│   │   │   ├── wildduck/
│   │   │   │   ├── client.go                 # 🆕 WildDuck SMTP/API client (connect to self-hosted WildDuck)
│   │   │   │   ├── smtp_sender.go            # 🆕 SMTP sending via WildDuck
│   │   │   │   ├── api_client.go             # 🆕 WildDuck REST API client (manage mailboxes, filters)
│   │   │   │   └── config.go                 # 🆕 WildDuck configuration (SMTP host, port, credentials)
│   │   │   └── smtp/
│   │   │       ├── client.go                 # Generic SMTP fallback (if not using WildDuck)
│   │   │       └── config.go                 # SMTP configuration
│   │   │
│   │   ├── storage/
│   │   │   └── client.go                     # Storage service client (upload message attachments via HTTP API)
│   │   │
│   │   └── outbox/                           # ❌ REMOVED: Use platform-shared/outbox instead
│   │       ├── processor.go                  # ❌ Use platform-shared/outbox/forwarder.go
│   │       └── scheduler.go                  # ❌ Use platform-shared/outbox/scheduler.go
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── conversation_handler.go   # Conversation HTTP handlers (GET, POST, PUT, DELETE /conversations, /conversations/:id)
│   │       │   ├── message_handler.go        # Message HTTP handlers (POST /conversations/:id/messages, GET /messages/:id, PUT, DELETE)
│   │       │   ├── notification_handler.go   # Notification HTTP handlers (GET /notifications, PUT /notifications/:id/read, DELETE /notifications/:id)
│   │       │   ├── in_app_notification_handler.go # 🆕 In-app notification handlers (GET /notifications/in-app, GET /notifications/badge-count)
│   │       │   ├── email_handler.go          # Email HTTP handlers (POST /emails, GET /emails/:id/status)
│   │       │   ├── template_handler.go       # Template HTTP handlers (GET /templates, POST /templates, GET /templates/:id)
│   │       │   ├── preferences_handler.go    # Preferences HTTP handlers (GET /preferences, PUT /preferences)
│   │       │   ├── websocket_handler.go      # 🆕 WebSocket upgrade handler (HTTP → WebSocket upgrade)
│   │       │   ├── sse_handler.go            # 🆕 SSE handler (establish SSE connections)
│   │       │   ├── online_status_handler.go  # Online status HTTP handlers (GET /users/:id/status, PUT /status/online)
│   │       │   ├── unsubscribe_handler.go    # Unsubscribe HTTP handlers (POST /unsubscribe, GET /unsubscribe/preferences)
│   │       │   ├── system_message_handler.go # 🆕 System message feed endpoints
│   │       │   ├── call_handler.go           # 🆕 Call links & scheduling endpoints
│   │       │   ├── calendar_invite_handler.go# 🆕 Calendar invite endpoints
│   │       │   ├── flag_handler.go           # Flag HTTP handlers (POST /messages/:id/flag, GET /flags)
│   │       │   └── health_handler.go         # Health check endpoints (/health, /ready, /live)
│   │       │
│   │       ├── routes/                       # 🆕 Route registrars mapped 1:1 to handlers
│   │       │   ├── conversation_routes.go    # /conversations, /conversations/:id
│   │       │   ├── message_routes.go         # /conversations/:id/messages
│   │       │   ├── notification_routes.go    # /notifications/*
│   │       │   ├── in_app_notification_routes.go # /notifications/in-app/*
│   │       │   ├── email_routes.go           # /emails/*
│   │       │   ├── template_routes.go        # /templates/*
│   │       │   ├── preferences_routes.go     # /preferences/*
│   │       │   ├── websocket_routes.go       # /ws (upgrade)
│   │       │   ├── sse_routes.go             # /sse/*
│   │       │   ├── online_status_routes.go   # /status/*
│   │       │   ├── unsubscribe_routes.go     # /unsubscribe/*
│   │       │   ├── system_message_routes.go  # /system-messages/*
│   │       │   ├── call_routes.go            # /calls/*
│   │       │   └── calendar_invite_routes.go # /calendar-invites/*
│   │       │
│   │       ├── middleware/                   # 📝 UPDATED: Middleware (now uses platform-shared middleware)
│   │       │   ├── auth.go                   # 📝 UPDATED: Authentication middleware (uses pkg/auth for JWT verification)
│   │       │   ├── rbac.go                   # 📝 UPDATED: RBAC middleware (uses pkg/auth authorizer for role checks)
│   │       │   ├── cors.go                   # 📝 UPDATED: CORS middleware (uses platform-shared/ginx CORS)
│   │       │   ├── rate_limit.go             # Rate limiting middleware (token bucket algorithm)
│   │       │   └── idempotency.go            # 🆕 Idempotency middleware (uses platform-shared/idempotency)
│   │       │
│   │       ├── responses/                    # 📝 UPDATED: Response helpers
│   │       │   └── README.md                 # 📝 Points to platform-shared/httpx (use shared response wrappers)
│   │       │
│   │       └── router.go                     # 📝 UPDATED: HTTP router setup (uses Gin, applies platform-shared/ginx middleware)
│   │
│   └── config/                               # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                         # 🆕 Typed Config struct (App, Server, Postgres, Kafka, Redis, Auth, WildDuck, WebSocket)
│       ├── loader.go                         # 🆕 Config loader using Viper (precedence: flags → env → file → defaults)
│       └── docs/
│           └── CONFIGURATION.md              # 🆕 Configuration documentation (all ENV vars, defaults, examples)
│
├── config/
│   ├── default.yaml                          # 🆕 Default configuration
│   ├── dev.yaml                              # 🆕 Development overrides
│   └── prod.yaml                             # 🆕 Production overrides
│
├── dapr/
│   ├── local/
│   │   ├── pubsub.yaml                       # Kafka pub/sub component
│   │   └── statestore.yaml                   # State store component
│   └── k8s/
│       ├── pubsub.yaml                       # Kafka with scopes: ["communications-be"]
│       ├── statestore.yaml                   # State store with scopes
│       └── secrets.yaml                      # Dapr secret store
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                         # Service-specific errors
│   │   └── codes.go                          # Error codes (MESSAGE_NOT_FOUND, CONVERSATION_NOT_FOUND)
│   ├── logger/                               # ❌ REMOVED: Use platform-shared/logging
│   │   └── README.md                         # 📝 Points to platform-shared/logging
│   ├── utils/
│   │   ├── validator.go                      # Local validation utilities
│   │   ├── template_engine.go                # Template rendering utilities
│   │   ├── sanitizer.go                      # Sanitize message content (prevent XSS)
│   │   └── html_to_text.go                   # Convert HTML to plain text (for email fallback)
│   └── constants/
│       ├── events.go                         # ❌ REMOVED: Use contracts/events
│       ├── topics.go                         # ❌ REMOVED: Use contracts/events
│       ├── notification_types.go             # Notification type constants
│       └── websocket_events.go               # WebSocket event types
│
├── templates/
│   ├── email/                                # Email HTML templates
│   │   ├── base.html                         # Base email template (header, footer, styles)
│   │   ├── welcome.html                      # Welcome email
│   │   ├── job_alert.html                    # Job alert email
│   │   ├── new_proposal.html                 # New proposal notification
│   │   ├── proposal_accepted.html            # Proposal accepted notification
│   │   ├── bid_received.html                 # New bid notification
│   │   ├── outbid_alert.html                 # Outbid alert
│   │   ├── contract_created.html             # Contract created notification
│   │   ├── milestone_completed.html          # Milestone completed notification
│   │   ├── payment_received.html             # Payment received notification
│   │   ├── payment_sent.html                 # Payment sent notification
│   │   ├── review_request.html               # Review request
│   │   ├── review_received.html              # Review received notification
│   │   ├── new_message.html                  # New message notification
│   │   ├── subscription_expiring.html        # Subscription expiring reminder
│   │   ├── password_reset.html               # Password reset email
│   │   ├── verify_email.html                 # Email verification
│   │   └── weekly_summary.html               # Weekly activity summary
│   └── notification/                         # In-app notification templates (JSON)
│       ├── job_posted.json                   # Job posted notification template
│       ├── proposal_received.json            # Proposal received template
│       ├── contract_created.json             # Contract created template
│       ├── milestone_completed.json          # Milestone completed template
│       ├── payment_received.json             # Payment received template
│       └── review_received.json              # Review received template
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   ├── seed-templates.sh
│   └── seed-data.sh
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── API.md
│   ├── EVENTS.md                             # 🆕 Events (published: message.sent, notification.delivered, consumed: all domain events)
│   ├── ARCHITECTURE.md
│   ├── MIGRATIONS.md                         # 🆕 Migration history
│   ├── SCHEMA.md                             # 🆕 Database schema
│   ├── RUNBOOK.md                            # 🆕 Operational procedures
│   ├── websocket-protocol.md                 # 🆕 WebSocket protocol documentation
│   ├── notification-system.md                # 🆕 Notification system overview
│   ├── in-app-notifications.md               # 🆕 In-app notifications guide
│   ├── email-templates.md
│   └── wildduck-integration.md               # 🆕 WildDuck integration guide
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod                                    # 📝 UPDATED: Imports pkg/auth, platform-shared, contracts/events
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md

```
---

## **📦 7️⃣ storage-be (UPDATED)**
```
apps/be/storage-be/
│
├── cmd/
│   └── api/
│       └── main.go                           # 📝 UPDATED: Application entry point - initializes Gin, Dapr, connects to Postgres (now uses loadConfig from internal/config, platform-shared/logging)
│
├── internal/
│   ├── domain/
│   │   ├── file/
│   │   │   ├── entity.go                     # File metadata
│   │   │   ├── enums.go                      # FileType, Status, Visibility
│   │   │   ├── metadata.go                   # File metadata
│   │   │   ├── errors.go                     # 🆕 File domain errors (FileNotFound, InvalidMimeType, FileTooLarge)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: FileCreated, FileUpdated, FileMoved, FileDeleted (hard)
│   │   │
│   │   ├── folder/
│   │   │   ├── entity.go                     # Folder structure
│   │   │   ├── errors.go                     # 🆕 Folder domain errors (FolderNotFound, NameConflict)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: FolderCreated, FolderRenamed, FolderMoved, FolderDeleted
│   │   │
│   │   ├── upload/
│   │   │   ├── entity.go                     # Upload sessions
│   │   │   ├── chunk.go                      # Chunked uploads
│   │   │   ├── resumable.go                  # Resumable upload support
│   │   │   ├── errors.go                     # 🆕 Upload domain errors (UploadSessionNotFound, ChunkOutOfOrder)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: UploadStarted, UploadChunkAppended, UploadResumed, UploadCompleted, UploadAborted
│   │   │
│   │   ├── media/
│   │   │   ├── entity.go                     # Media processing records
│   │   │   ├── thumbnail.go                  # Thumbnail generation
│   │   │   ├── variant.go                    # Image variants
│   │   │   ├── errors.go                     # 🆕 Media domain errors (ProcessingFailed, UnsupportedFormat)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: MediaProcessingStarted, MediaProcessingSucceeded, MediaProcessingFailed, ThumbnailGenerated
│   │   │
│   │   ├── access_control/
│   │   │   ├── entity.go                     # File access permissions
│   │   │   ├── errors.go                     # 🆕 Access control errors (PermissionDenied, ACLNotFound)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: AccessGranted, AccessRevoked, ACLUpdated
│   │   │
│   │   ├── version/
│   │   │   ├── entity.go                     # File versioning
│   │   │   ├── errors.go                     # 🆕 Versioning errors (VersionNotFound, ImmutableVersion)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: FileVersionCreated, FileVersionPromoted, FileVersionRestored
│   │   │
│   │   ├── share/
│   │   │   ├── entity.go                     # File sharing
│   │   │   ├── link.go                       # Share links
│   │   │   ├── errors.go                     # 🆕 Share errors (ShareNotFound, ShareExpired)
│   │   │   ├── repository.go                 # ShareRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: ShareCreated, ShareRevoked, ShareExpired
│   │   │
│   │   ├── file_flag/
│   │   │   ├── entity.go                     # Flagged files
│   │   │   ├── reason.go                     # Flag reasons
│   │   │   ├── status.go                     # Flag status
│   │   │   ├── errors.go                     # 🆕 Flag errors (FlagNotFound, InvalidFlagReason)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: FileFlagSubmitted, FileFlagResolved, FileFlagDismissed
│   │   │
│   │   ├── policy/                           # 🆕 File policies: size/type, virus scan, DLP
│   │   │   ├── entity.go                     # Policy (ID, Name, MaxSizeMB, AllowedMIMEs, BlockedMIMEs, EnforceVirusScan, EnforceDLP)
│   │   │   ├── dlp_pattern.go                # DLP patterns for PII/PCI (regexes, detectors)
│   │   │   ├── result.go                     # Policy evaluation results (violations, reasons)
│   │   │   ├── errors.go                     # Policy errors
│   │   │   ├── repository.go                 # PolicyRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: FilePolicyCreated, FilePolicyUpdated, FilePolicyViolationDetected
│   │   │
│   │   ├── lifecycle/                        # 🆕 Lifecycle rules: retention, legal hold, soft-delete
│   │   │   ├── entity.go                     # LifecycleRule (ContractState, RetentionDays, LegalHold, RestoreWindowDays)
│   │   │   ├── soft_delete.go                # Soft-delete markers (DeletedAt, RestoreBy)
│   │   │   ├── legal_hold.go                 # Legal hold records (PlacedBy, Reason, ExpiresAt)
│   │   │   ├── errors.go                     # Lifecycle errors
│   │   │   ├── repository.go                 # LifecycleRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: FileSoftDeleted, FileRestored, LegalHoldPlaced, LegalHoldRemoved
│   │   │
│   │   ├── linking/                          # 🆕 Versioning & linking: signed URLs, download audit logs
│   │   │   ├── signed_url.go                 # Signed URL (FileID, Action, ExpiresAt, Signature)
│   │   │   ├── audit_log.go                  # Download audit log (FileID, UserID, IP, UserAgent, At)
│   │   │   ├── errors.go                     # Linking errors
│   │   │   ├── repository.go                 # LinkingRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: SignedURLCreated, SignedURLRevoked, FileDownloadLogged
│   │   │
│   │   └── outbox/
│   │       ├── entity.go                     # ❌ REMOVED (use platform-shared/outbox/entity.go)
│   │       └── repository.go                 # ❌ REMOVED (use platform-shared/outbox/repository.go)
│   │
│   ├── application/
│   │   ├── eventhandler/
│   │   │   ├── user_handler.go                # Consumes: user.updated
│   │   │   │                                     # Purpose: refresh ownership/ACL context (e.g., org/role changes) used by access_control/policy.
│   │   │   │
│   │   │   ├── contract_handler.go            # Consumes: contract.state.changed
│   │   │   │                                     # Purpose: enforce lifecycle/retention & legal holds by contract state
│   │   │   │                                     #           (place/remove holds; adjust restore windows/retention markers).
│   │   │   │
│   │   │   ├── admin_policy_handler.go        # Consumes: admin.policy.updated
│   │   │   │                                     # Purpose: update storage policy cache (size/type limits, virus-scan & DLP patterns/rules).
│   │   │   │
│   │   │   └── admin_moderation_handler.go    # Consumes: admin.moderation.actioned (approve/reject/remove/hide)
│   │   │                                         # Purpose: reflect moderation outcomes in storage:
│   │   │                                         #           - quarantine or restore files,
│   │   │                                         #           - revoke or reissue signed URLs,
│   │   │                                         #           - toggle visibility flags for indexed artifacts.
│   │   ├── file/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Upload, Delete, Move, Copy
│   │   │   ├── queries.go                    # Get, List, Search
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── upload/
│   │   │   ├── service.go
│   │   │   ├── chunked_upload.go
│   │   │   ├── resumable.go
│   │   │   ├── dto.go
│   │   │   ├── commands.go                   # 🆕 StartUpload, AppendChunk, CompleteUpload, AbortUpload
│   │   │   ├── queries.go                    # 🆕 GetUploadSession, GetUploadProgress
│   │   │   └── validators.go                 # 🆕 Validate chunk order/size, checksum
│   │   ├── media/
│   │   │   ├── service.go
│   │   │   ├── image_processor.go
│   │   │   ├── video_processor.go
│   │   │   ├── thumbnail_generator.go
│   │   │   ├── dto.go
│   │   │   ├── commands.go                   # 🆕 ProcessImage, ProcessVideo, GenerateThumbnail
│   │   │   ├── queries.go                    # 🆕 GetMediaJob, ListMediaJobs
│   │   │   └── validators.go                 # 🆕 Validate dimensions/codec/bitrate
│   │   ├── folder/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                   # 🆕 CreateFolder, RenameFolder, MoveFolder, DeleteFolder
│   │   │   ├── queries.go                    # 🆕 GetFolder, ListFolderContents, SearchFolders
│   │   │   └── validators.go                 # 🆕 Validate names, cycles, path depth
│   │   ├── share/
│   │   │   ├── service.go
│   │   │   ├── link_generator.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                   # 🆕 CreateShare, RevokeShare, UpdateShare
│   │   │   ├── queries.go                    # 🆕 GetShare, ListSharesByFile
│   │   │   └── validators.go                 # 🆕 Validate expiry/scopes/access
│   │   ├── version/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   │   ├── commands.go                   # 🆕 CreateVersion, RestoreVersion, DeleteVersion
│   │   │   ├── queries.go                    # 🆕 GetVersion, ListVersions
│   │   │   └── validators.go                 # 🆕 Validate immutable rules
│   │   ├── flag/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Flag, Unflag file
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go                 # 🆕 Validate flag reason/state transitions
│   │   ├── policy/                           # 🆕 File policy application layer
│   │   │   ├── service.go                    # EvaluatePolicy, UpdatePolicy, GetPolicy
│   │   │   ├── commands.go                   # SetPolicy, EnableDLP, DisableDLP
│   │   │   ├── queries.go                    # GetPolicy, ListPolicies
│   │   │   ├── validators.go                 # Validate size/type lists, regexes
│   │   │   ├── dto.go                        # PolicyDTO, DLPResultDTO
│   │   │   └── mapper.go                     # Policy mappers
│   │   ├── lifecycle/                        # 🆕 Lifecycle application layer
│   │   │   ├── service.go                    # ApplyRules, SoftDelete, Restore, PlaceLegalHold, RemoveLegalHold
│   │   │   ├── commands.go                   # DefineRule, UpdateRule, DeleteRule, SoftDeleteFile, RestoreFile, PlaceLegalHold
│   │   │   ├── queries.go                    # GetRules, GetFileLifecycle, GetLegalHolds
│   │   │   ├── validators.go                 # Validate retention windows, restore periods
│   │   │   ├── dto.go                        # LifecycleRuleDTO, LegalHoldDTO
│   │   │   └── mapper.go                     # Lifecycle mappers
│   │   ├── linking/                          # 🆕 Signed URLs & audit logs
│   │   │   ├── service.go                    # CreateSignedURL, RevokeSignedURL, LogDownload, ListAuditLogs
│   │   │   ├── commands.go                   # CreateSignedURL, RevokeSignedURL
│   │   │   ├── queries.go                    # GetSignedURL, ListSignedURLs, GetAuditLogs
│   │   │   ├── validators.go                 # Validate expiry, actions, scopes
│   │   │   ├── dto.go                        # SignedURLDTO, AuditLogDTO
│   │   │   └── mapper.go                     # Linking mappers
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection setup (DSN from config, connection pooling)
│   │   │       ├── transaction.go            # Transaction helpers (Begin, Commit, Rollback, WithTransaction wrapper)
│   │   │       ├── migrations.go             # 📝 UPDATED: Auto-migrate logic (now with version tracking, GORM AutoMigrate for all tables)
│   │   │       ├── version.go                # 🆕 Schema version tracking (SchemaVersion table, RecordMigration function)
│   │   │       ├── safety.go                 # 🆕 Pre-migration safety checks (environment validation, disk space check)
│   │   │       ├── file_repository.go        # FileRepository implementation with GORM
│   │   │       ├── folder_repository.go
│   │   │       ├── upload_repository.go
│   │   │       ├── media_repository.go
│   │   │       ├── access_control_repository.go
│   │   │       ├── version_repository.go
│   │   │       ├── share_repository.go
│   │   │       ├── file_flag_repository.go
│   │   │       ├── policy_repository.go      # 🆕 PolicyRepository implementation
│   │   │       ├── lifecycle_repository.go   # 🆕 LifecycleRepository implementation
│   │   │       └── linking_repository.go     # 🆕 LinkingRepository implementation (signed URLs, audit logs)
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go             # Redis connection setup (connection pooling, retry logic)
│   │   │       └── file_cache.go             # File metadata caching (Get, Set, Invalidate with TTL)
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go               # 📝 UPDATED: Kafka consumer (now uses platform-shared/inbox for message deduplication)
│   │   │       ├── producer.go               # 📝 UPDATED: Kafka producer (now uses platform-shared/outbox for reliable publishing)
│   │   │       ├── topics.go                 # 📝 UPDATED: Topic constants (now imported from contracts/events - file.uploaded, file.deleted, media.processed)
│   │   │       └── scram.go                  # SCRAM authentication for Kafka (SASL/SCRAM-SHA-256)
│   │   ├── object_storage/
│   │   │   ├── local/
│   │   │   │   ├── storage.go                # Local file system storage
│   │   │   │   └── config.go                 # Local storage config
│   │   │   ├── minio/
│   │   │   │   ├── client.go                 # Self-hosted MinIO client (upload, download, presigned URLs)
│   │   │   │   └── config.go                 # MinIO configuration (endpoint, credentials, bucket names)
│   │   │   ├── signer.go                     # 🆕 Signed URL signer abstraction (create/revoke)
│   │   │   └── provider.go                   # Storage provider abstraction
│   │   ├── media_processing/
│   │   │   ├── image/
│   │   │   │   ├── resizer.go                # Image resizing logic
│   │   │   │   ├── optimizer.go              # Image optimization (compression)
│   │   │   │   └── watermark.go              # Watermarking images
│   │   │   ├── video/
│   │   │   │   ├── transcoder.go             # Video transcoding
│   │   │   │   └── thumbnail.go              # Video thumbnail generation
│   │   │   └── processor.go                  # General media processor
│   │   ├── virus_scan/
│   │   │   └── clamav.go                     # ClamAV integration for virus scanning
│   │   ├── dlp/                              # 🆕 DLP engine adapters
│   │   │   ├── regex_engine.go               # Regex-based detectors (PII/PCI patterns)
│   │   │   └── provider.go                   # Pluggable DLP providers (custom/third-party)
│   │   └── outbox/
│   │       ├── processor.go                  # ❌ REMOVED (use platform-shared/outbox/forwarder.go)
│   │       └── scheduler.go                  # ❌ REMOVED (use platform-shared/outbox/scheduler.go)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── file_handler.go           # File HTTP handlers (GET, POST, DELETE /files)
│   │       │   ├── upload_handler.go         # Upload handlers (POST /upload, chunked support)
│   │       │   ├── download_handler.go       # Download handlers (GET /download/:id)
│   │       │   ├── folder_handler.go         # Folder handlers
│   │       │   ├── share_handler.go          # Share handlers
│   │       │   ├── media_handler.go          # Media processing handlers
│   │       │   ├── flag_handler.go           # Flag handlers
│   │       │   ├── policy_handler.go         # 🆕 Policy endpoints (GET/PUT policies, evaluate)
│   │       │   ├── lifecycle_handler.go      # 🆕 Lifecycle endpoints (soft-delete/restore/legal hold)
│   │       │   ├── linking_handler.go        # 🆕 Signed URL + audit log endpoints
│   │       │   └── health_handler.go         # Health check endpoints (/health, /ready, /live)
│   │       ├── middleware/
│   │       │   ├── auth.go                   # 📝 UPDATED: Authentication middleware (uses pkg/auth for JWT verification)
│   │       │   ├── rbac.go                   # 📝 UPDATED: RBAC middleware (uses pkg/auth authorizer for role checks)
│   │       │   ├── cors.go                   # 📝 UPDATED: CORS middleware (uses platform-shared/ginx CORS)
│   │       │   ├── rate_limit.go             # Rate limiting middleware (token bucket algorithm)
│   │       │   ├── logging.go                # 📝 UPDATED: Logging middleware (uses platform-shared/ginx/logging)
│   │       │   ├── error_handler.go          # Error handling middleware
│   │       │   ├── request_id.go             # 📝 UPDATED: Request ID middleware (uses platform-shared/ginx/requestid)
│   │       │   └── file_size_limit.go        # File size limit middleware
│   │       ├── responses/
│   │       │   ├── success.go                # 📝 UPDATED: Success response wrappers (use platform-shared/httpx/response.go)
│   │       │   ├── error.go                  # 📝 UPDATED: Error responses (use platform-shared/httpx/errors.go)
│   │       │   └── pagination.go             # 📝 UPDATED: Pagination (use platform-shared/httpx/pagination.go)
│   │       ├── routes/                       # 🆕 Route registrars mapped 1:1 to handlers
│   │       │   ├── file_routes.go            # /files/*
│   │       │   ├── upload_routes.go          # /upload/*
│   │       │   ├── download_routes.go        # /download/*
│   │       │   ├── folder_routes.go          # /folders/*
│   │       │   ├── share_routes.go           # /shares/*
│   │       │   ├── media_routes.go           # /media/*
│   │       │   ├── flag_routes.go            # /flags/*
│   │       │   ├── policy_routes.go          # 🆕 /policies/*
│   │       │   ├── lifecycle_routes.go       # 🆕 /lifecycle/*
│   │       │   └── linking_routes.go         # 🆕 /links/* (signed URLs, audit logs)
│   │       └── router.go                     # 📝 UPDATED: HTTP router setup (uses Gin, applies platform-shared/ginx middleware)
│   │
│   └── config/                               # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                         # 🆕 Typed Config struct (App, Server, Postgres, Kafka, Redis, MinIO, Storage)
│       ├── loader.go                         # 🆕 Config loader using Viper (precedence: flags → env → file → defaults)
│       └── docs/
│           └── CONFIGURATION.md              # 🆕 Configuration documentation (all ENV vars, defaults, examples)
│
├── config/
│   ├── default.yaml
│   ├── dev.yaml
│   └── prod.yaml
│
├── dapr/
│   ├── local/
│   │   ├── pubsub.yaml
│   │   └── statestore.yaml
│   └── k8s/
│       ├── pubsub.yaml                       # Kafka with scopes: ["storage-be"]
│       ├── statestore.yaml                   # State store with scopes
│       └── secrets.yaml                      # Dapr secret store
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                         # Service-specific errors
│   │   └── codes.go                          # Error codes
│   ├── logger/
│   │   └── logger.go                         # ❌ REMOVED: Use platform-shared/logging
│   ├── utils/
│   │   ├── validator.go                      # Local validation utilities
│   │   ├── file_utils.go                     # File utilities (path manipulation, extension extraction)
│   │   ├── mime_detector.go                  # MIME type detection
│   │   └── hash.go                           # File hash calculation (MD5, SHA256)
│   └── constants/
│       ├── events.go                         # ❌ REMOVED: Use contracts/events
│       ├── topics.go                         # ❌ REMOVED: Use contracts/events
│       └── mime_types.go                     # Supported MIME types
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       ├── pvc.yaml
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   ├── seed-data.sh
│   └── cleanup-orphaned.sh
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── api.md
│   ├── events.md                             # 📝 UPDATED: Events (published: file.uploaded, file.deleted, consumed: user.created, etc.)
│   ├── upload-flow.md
│   ├── media-processing.md
│   ├── MIGRATIONS.md
│   ├── SCHEMA.md
│   └── RUNBOOK.md
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md

```
---

## **📦8️⃣ search-be (UPDATED WITH ENHANCED RECOMMENDATIONS)**

```
apps/be/search-be/
│
├── cmd/
│   └── api/
│       └── main.go                           # 📝 UPDATED: Application entry point - initializes Gin, Dapr, connects to Elasticsearch (now uses loadConfig from internal/config, platform-shared/logging)
│
├── internal/
│   ├── domain/
│   │   ├── search_index/
│   │   │   ├── entity.go                     # Search index metadata
│   │   │   ├── job_index.go                  # Job search document
│   │   │   ├── user_index.go                 # User/Freelancer document
│   │   │   ├── errors.go                     # 🆕 Index errors (IndexNotFound, DocumentConflict)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: search.document.indexed / search.document.reindexed / search.index.visibility.changed / search.index.archived
│   │   │
│   │   ├── search_query/
│   │   │   ├── entity.go                     # Search query logs
│   │   │   ├── filters.go                    # Search filters
│   │   │   ├── errors.go                     # 🆕 Query errors (InvalidFilter, BadSort)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: search.query.logged / search.query.alert.triggered
│   │   │
│   │   ├── recommendation/
│   │   │   ├── entity.go                     # Recommendation records
│   │   │   ├── score.go                      # Scoring algorithm
│   │   │   ├── reason.go                     # Recommendation reasons
│   │   │   ├── feedback.go                   # User feedback on recommendations
│   │   │   ├── errors.go                     # 🆕 Recommendation errors (ModelUnavailable, ScoringFailed)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: search.recommendation.generated / search.recommendation.feedback.recorded
│   │   │
│   │   ├── recommendation_model/
│   │   │   ├── entity.go                     # ML model metadata
│   │   │   ├── feature.go                    # Feature vectors
│   │   │   ├── training_data.go              # Training data
│   │   │   ├── errors.go                     # 🆕 Model errors (VersionNotFound, FeatureMismatch)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: search.model.version.registered / search.model.deployed / search.model.rolled_back / search.model.training_data.updated
│   │   │
│   │   ├── user_preference/
│   │   │   ├── entity.go                     # User preferences for recommendations
│   │   │   ├── implicit_signals.go           # Implicit user signals (clicks, views)
│   │   │   ├── explicit_preferences.go       # Explicit preferences
│   │   │   ├── errors.go                     # 🆕 Preference errors (PreferenceNotFound)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: search.personalization.updated / search.preference.updated
│   │   │
│   │   ├── matching/
│   │   │   ├── entity.go                     # Job-Freelancer matches
│   │   │   ├── criteria.go                   # Matching criteria
│   │   │   ├── score_breakdown.go            # Detailed match scoring
│   │   │   ├── errors.go                     # 🆕 Matching errors (CriteriaInvalid)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: search.match.calculated / search.match.accepted / search.match.dismissed
│   │   │
│   │   ├── feed/
│   │   │   ├── entity.go                     # User feeds
│   │   │   ├── item.go                        # Feed items
│   │   │   ├── personalization.go             # Personalization data
│   │   │   ├── errors.go                      # 🆕 Feed errors (FeedNotFound)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: search.feed.item.added / search.feed.item.removed / search.feed.updated
│   │   │
│   │   ├── trending/
│   │   │   ├── entity.go                      # Trending jobs/skills
│   │   │   ├── calculator.go                  # Calculate trending items
│   │   │   ├── errors.go                      # 🆕 Trending errors
│   │   │   ├── repository.go
│   │   │   └── events.go                      # 🆕 Domain events: search.trending.calculated / search.trending.updated
│   │   │
│   │   ├── saved_search/
│   │   │   ├── entity.go                      # Saved searches
│   │   │   ├── alert.go                        # Search alerts
│   │   │   ├── errors.go                       # 🆕 Saved search errors
│   │   │   ├── repository.go
│   │   │   └── events.go                       # 🆕 Domain events: search.saved_search.created / search.saved_search.updated / search.saved_search.deleted / search.saved_search.alert.sent
│   │   │
│   │   ├── similarity/
│   │   │   ├── entity.go                      # Similar jobs/users
│   │   │   ├── vector.go                      # Similarity vectors
│   │   │   ├── errors.go                      # 🆕 Similarity errors
│   │   │   ├── repository.go
│   │   │   └── events.go                      # 🆕 Domain events: search.similarity.computed / search.similarity.model.updated
│   │   │
│   │   ├── taxonomy/                          # 🆕 Taxonomy & synonyms
│   │   │   ├── entity.go                       # Normalized skills/categories (SkillID, Name, CategoryID, Aliases)
│   │   │   ├── category.go                     # Categories (ID, Name, ParentID, Path)
│   │   │   ├── synonym.go                      # Synonyms & alias mapping (SkillID, Alias, ToleranceScore)
│   │   │   ├── typos.go                        # Typo tolerance rules (edit distance thresholds)
│   │   │   ├── errors.go                        # Taxonomy errors (SkillNotFound, AliasConflict)
│   │   │   ├── repository.go                    # TaxonomyRepository interface
│   │   │   └── events.go                        # 🆕 Domain events: search.taxonomy.updated / search.taxonomy.synonym.changed / search.taxonomy.typo_rules.updated
│   │   │
│   │   ├── ltr/                                # 🆕 Learning-to-rank hooks
│   │   │   ├── entity.go                       # LTR signals (docID, conversions, responseSpeed, jobSuccess, intent)
│   │   │   ├── signal_source.go                # Signal sources (events → features)
│   │   │   ├── feature_store.go                # Persisted LTR features (effectiveAt timestamps)
│   │   │   ├── errors.go                       # LTR errors (SignalOutOfRange, FeatureStale)
│   │   │   ├── repository.go                   # LTRRepository interface
│   │   │   └── events.go                       # 🆕 Domain events: search.ltr.signal.recorded / search.ltr.features.updated
│   │   │
│   │   ├── facets/                             # 🆕 Facets & filters
│   │   │   ├── entity.go                       # Facet definitions (price bands, availability, tz overlap, languages, badges)
│   │   │   ├── banding.go                      # Price/rate band logic
│   │   │   ├── tz_overlap.go                   # Timezone overlap calculator
│   │   │   ├── errors.go                       # Facet errors (InvalidBand, UnknownBadge)
│   │   │   ├── repository.go                   # FacetRepository interface
│   │   │   └── events.go                       # 🆕 Domain events: search.facets.definition.updated / search.facets.banding.updated / search.facets.tz_overlap.updated
│   │   │
│   │   ├── personalization/                    # 🆕 Personalization signals
│   │   │   ├── entity.go                       # Profiles (recent activity, saved searches, preferred rates)
│   │   │   ├── cold_start.go                   # Cold-start defaults
│   │   │   ├── errors.go                       # Personalization errors
│   │   │   ├── repository.go                   # PersonalizationRepository interface
│   │   │   └── events.go                       # 🆕 Domain events: search.personalization.updated / search.personalization.cold_start.applied
│   │   │
│   │   ├── hygiene/                            # 🆕 Index hygiene
│   │   │   ├── entity.go                       # Hygiene tasks (incremental updates, dedup, archival, visibility)
│   │   │   ├── incremental.go                  # Incremental update markers (version, changed fields)
│   │   │   ├── dedupe.go                       # Dedupe fingerprints
│   │   │   ├── visibility.go                   # Visibility states
│   │   │   ├── errors.go                       # Hygiene errors
│   │   │   ├── repository.go                   # HygieneRepository interface
│   │   │   └── events.go                       # 🆕 Domain events: search.index.hygiene.deduplicated / search.index.hygiene.archived / search.index.hygiene.visibility.changed
│   │   │
│   │   └── outbox/
│   │       ├── entity.go                       # ❌ REMOVED (use platform-shared/outbox/entity.go)
│   │       └── repository.go                    # ❌ REMOVED (use platform-shared/outbox/repository.go)
│   │
│   ├── application/
│   │   ├── eventhandler/
│   │   │   ├── job_handler.go                 # Consumes: job.posted | job.updated | job.closed
│   │   │   │                                     # Purpose: create/update/remove job documents in the search index; refresh facets.
│   │   │   │
│   │   │   ├── user_handler.go                # Consumes: user.updated
│   │   │   │                                     # Purpose: refresh user/freelancer search documents (profile fields, visibility, badges).
│   │   │   │
│   │   │   ├── review_handler.go              # Consumes: review.double_blind.published | review.published
│   │   │   │                                     # Purpose: update public rating aggregates on indexed user/job docs; recalc ranking features.
│   │   │   │
│   │   │   ├── entitlement_handler.go         # Consumes: subscription.feature.changed
│   │   │   │                                     # Purpose: adjust facet/visibility rules and gated search features tied to plan changes.
│   │   │   │
│   │   │   ├── admin_content_handler.go       # Consumes: admin.content.actioned (approve/reject/remove/hide)
│   │   │   │                                     # Purpose: hide/unhide or restore affected indexed documents promptly.
│   │   │   │
│   │   │   ├── admin_flags_handler.go         # Consumes: admin.feature_flag.updated | admin.threshold.updated | admin.experiment.updated
│   │   │   │                                     # Purpose: refresh runtime toggles that control search facets, boosts, filters, and experiments.
│   │   │   │
│   │   │   ├── storage_lifecycle_handler.go   # Consumes: file.lifecycle.soft_deleted | file.lifecycle.restored
│   │   │   │                                     # Purpose: sync file-backed assets’ visibility in indexed documents (thumbnails, attachments).
│   │   │   │
│   │   │   └── _partitioning_notes.go         # (comments only) Partition keys used by handlers:
│   │   │                                         #   job.* → key by job_id
│   │   │                                         #   user.* → key by user_id
│   │   │                                         #   review.* → key by subject_user_id (or contract_id when provided)
│   │   │                                         #   admin.* → key by affected_resource_id
│   │   │                                         #   file.lifecycle.* → key by file_id
│   │   ├── search/
│   │   │   ├── service.go
│   │   │   ├── job_search.go
│   │   │   ├── freelancer_search.go
│   │   │   ├── query_builder.go
│   │   │   ├── facet_builder.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                    # 🆕 ExecuteSearch (jobs/freelancers), SaveSearch
│   │   │   ├── queries.go                     # 🆕 GetSearchResults, GetSearchSuggestions
│   │   │   └── validators.go                  # 🆕 Validate filters/sorts/facets
│   │   ├── indexing/
│   │   │   ├── service.go
│   │   │   ├── job_indexer.go
│   │   │   ├── user_indexer.go
│   │   │   ├── bulk_indexer.go
│   │   │   ├── dto.go
│   │   │   ├── commands.go                    # 🆕 IndexJob, IndexUser, ReindexAll
│   │   │   ├── queries.go                     # 🆕 GetIndexStatus, GetDocByID
│   │   │   └── validators.go                  # 🆕 Validate index payloads
│   │   ├── recommendation/
│   │   │   ├── service.go
│   │   │   ├── job_recommender.go             # Recommend jobs to freelancers
│   │   │   ├── freelancer_recommender.go      # Recommend freelancers to clients
│   │   │   ├── collaborative_filtering.go     # Collaborative filtering algorithm
│   │   │   ├── content_based.go               # Content-based filtering
│   │   │   ├── hybrid_recommender.go          # Hybrid recommendation approach
│   │   │   ├── scoring_engine.go              # Calculate recommendation scores
│   │   │   ├── personalization.go             # Personalization logic
│   │   │   ├── diversity_optimizer.go         # Ensure diverse recommendations
│   │   │   ├── cold_start_handler.go          # Handle new users/jobs
│   │   │   ├── dto.go
│   │   │   ├── ml_model.go                    # ML model integration
│   │   │   ├── commands.go                    # 🆕 GenerateRecommendations, RecordFeedback
│   │   │   ├── queries.go                     # 🆕 GetRecommendations, GetReasons
│   │   │   └── validators.go                  # 🆕 Validate rec limits, diversity params
│   │   ├── matching/
│   │   │   ├── service.go
│   │   │   ├── matcher.go                     # Job-Freelancer matching
│   │   │   ├── criteria_evaluator.go          # Evaluate match criteria
│   │   │   ├── skill_matcher.go               # Match based on skills
│   │   │   ├── experience_matcher.go          # Match based on experience
│   │   │   ├── rate_matcher.go                # Match based on rates
│   │   │   ├── availability_matcher.go        # Match based on availability
│   │   │   ├── score_calculator.go            # Calculate match score
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                    # 🆕 CreateMatchRun
│   │   │   ├── queries.go                     # 🆕 GetMatchesForJob, GetMatchesForUser
│   │   │   └── validators.go                  # 🆕 Validate criteria completeness
│   │   ├── feed/
│   │   │   ├── service.go
│   │   │   ├── generator.go                   # Generate personalized feed
│   │   │   ├── ranking.go                     # Rank feed items
│   │   │   ├── freshness_scorer.go            # Score by freshness
│   │   │   ├── relevance_scorer.go            # Score by relevance
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                    # 🆕 GenerateFeed
│   │   │   ├── queries.go                     # 🆕 GetFeed
│   │   │   └── validators.go                  # 🆕 Validate feed window/size
│   │   ├── trending/
│   │   │   ├── service.go
│   │   │   ├── calculator.go                  # Calculate trending items
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                    # 🆕 RecomputeTrending
│   │   │   ├── queries.go                     # 🆕 GetTrending
│   │   │   └── validators.go                  # 🆕 Validate window/minSupport
│   │   ├── similarity/
│   │   │   ├── service.go
│   │   │   ├── job_similarity.go              # Find similar jobs
│   │   │   ├── user_similarity.go             # Find similar users
│   │   │   ├── vector_calculator.go           # Calculate similarity vectors
│   │   │   ├── dto.go
│   │   │   ├── commands.go                    # 🆕 RebuildSimilarityVectors
│   │   │   ├── queries.go                     # 🆕 GetSimilarJobs, GetSimilarUsers
│   │   │   └── validators.go                  # 🆕 Validate k/threshold
│   │   ├── suggestion/
│   │   │   ├── service.go                     # Autocomplete suggestions
│   │   │   ├── dto.go
│   │   │   ├── cache_warmer.go
│   │   │   ├── commands.go                    # 🆕 WarmSuggestionCache
│   │   │   ├── queries.go                     # 🆕 GetSuggestions
│   │   │   └── validators.go                  # 🆕 Validate prefix/lang
│   │   ├── taxonomy/                          # 🆕 Taxonomy application layer
│   │   │   ├── service.go                      # Manage skills/categories, synonyms, typo rules
│   │   │   ├── commands.go                     # UpsertSkill, UpsertCategory, AddAlias, RemoveAlias
│   │   │   ├── queries.go                      # GetSkill, ListSkills, GetCategoryTree
│   │   │   ├── validators.go                   # Validate alias conflicts, edit distance
│   │   │   ├── dto.go                           # TaxonomyDTO, CategoryDTO, AliasDTO
│   │   │   └── mapper.go                        # Taxonomy mappers
│   │   ├── ltr/                                # 🆕 LTR application layer
│   │   │   ├── service.go                      # Ingest signals, build feature vectors, rerank
│   │   │   ├── commands.go                     # IngestSignalBatch, SnapshotFeatures
│   │   │   ├── queries.go                      # GetFeaturesForDoc, PreviewRerank
│   │   │   ├── validators.go                   # Validate signal bounds, effectiveAt
│   │   │   ├── dto.go                           # SignalDTO, FeatureDTO, RerankPreviewDTO
│   │   │   └── mapper.go                        # LTR mappers
│   │   ├── facets/                             # 🆕 Facets application layer
│   │   │   ├── service.go                      # Build facets, apply filters, banding
│   │   │   ├── commands.go                     # DefineFacet, UpdateFacetBands
│   │   │   ├── queries.go                      # GetFacetsForQuery
│   │   │   ├── validators.go                   # Validate bands, tz overlap rules
│   │   │   ├── dto.go                           # FacetDTO, BandDTO
│   │   │   └── mapper.go                        # Facet mappers
│   │   ├── personalization/                    # 🆕 Personalization application layer
│   │   │   ├── service.go                      # Update profile, compute defaults
│   │   │   ├── commands.go                     # UpdatePersonalizationProfile
│   │   │   ├── queries.go                      # GetPersonalizationProfile
│   │   │   ├── validators.go                   # Validate rate prefs, recent activity bounds
│   │   │   ├── dto.go                           # PersonalizationDTO
│   │   │   └── mapper.go                        # Personalization mappers
│   │   ├── hygiene/                            # 🆕 Index hygiene application layer
│   │   │   ├── service.go                      # Incremental updates, dedupe, archive/visibility
│   │   │   ├── commands.go                     # RunIncrementalUpdate, RunDedup, ArchiveDoc, SetVisibility
│   │   │   ├── queries.go                      # GetHygieneStatus, GetDocHistory
│   │   │   ├── validators.go                   # Validate dedupe fingerprint size, state transitions
│   │   │   ├── dto.go                           # HygieneTaskDTO, VisibilityDTO
│   │   │   └── mapper.go                        # Hygiene mappers
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go              # PostgreSQL connection setup (DSN from config, connection pooling)
│   │   │       ├── transaction.go             # Transaction helpers (Begin, Commit, Rollback, WithTransaction wrapper)
│   │   │       ├── migrations.go              # 📝 UPDATED: Auto-migrate logic (now with version tracking, GORM AutoMigrate for all tables)
│   │   │       ├── version.go                 # 🆕 Schema version tracking (SchemaVersion table, RecordMigration function)
│   │   │       ├── safety.go                  # 🆕 Pre-migration safety checks (environment validation, disk space check)
│   │   │       ├── search_query_repository.go
│   │   │       ├── recommendation_repository.go
│   │   │       ├── recommendation_model_repository.go
│   │   │       ├── user_preference_repository.go
│   │   │       ├── matching_repository.go
│   │   │       ├── feed_repository.go
│   │   │       ├── trending_repository.go
│   │   │       ├── saved_search_repository.go
│   │   │       ├── similarity_repository.go
│   │   │       ├── taxonomy_repository.go     # 🆕 TaxonomyRepository implementation
│   │   │       ├── ltr_repository.go          # 🆕 LTRRepository implementation
│   │   │       ├── facets_repository.go       # 🆕 FacetRepository implementation
│   │   │       ├── personalization_repository.go # 🆕 PersonalizationRepository implementation
│   │   │       └── hygiene_repository.go      # 🆕 HygieneRepository implementation
│   │   ├── elasticsearch/
│   │   │   ├── client.go                      # Elasticsearch client
│   │   │   ├── index_manager.go               # Index management
│   │   │   ├── job_mapper.go                  # Job mapping to ES document
│   │   │   ├── user_mapper.go                 # User mapping to ES document
│   │   │   └── config.go                      # ES configuration (hosts, auth)
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go              # Redis connection setup (connection pooling, retry logic)
│   │   │       ├── search_cache.go            # Cache search results (Get, Set, Invalidate with TTL)
│   │   │       ├── suggestion_cache.go        # Cache autocomplete suggestions
│   │   │       ├── feed_cache.go              # Cache user feeds
│   │   │       ├── recommendation_cache.go    # Cache recommendations
│   │   │       ├── taxonomy_cache.go          # 🆕 Cache skills/categories & aliases
│   │   │       └── ltr_cache.go               # 🆕 Cache LTR features snapshot
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go                # 📝 UPDATED: Kafka consumer (now uses platform-shared/inbox for message deduplication)
│   │   │       ├── producer.go                # 📝 UPDATED: Kafka producer (now uses platform-shared/outbox for reliable publishing)
│   │   │       ├── topics.go                  # 📝 UPDATED: Topic constants (now imported from contracts/events - job.indexed, user.indexed, recommendation.generated)
│   │   │       └── scram.go                   # SCRAM authentication for Kafka (SASL/SCRAM-SHA-256)
│   │   ├── ml/
│   │   │   ├── model_loader.go                # Load ML models
│   │   │   ├── predictor.go                   # Make predictions
│   │   │   ├── feature_extractor.go           # Extract features
│   │   │   ├── trainer.go                     # Train models
│   │   │   └── evaluator.go                   # Evaluate model performance
│   │   └── outbox/
│   │       ├── processor.go                   # ❌ REMOVED (use platform-shared/outbox/forwarder.go)
│   │       └── scheduler.go                   # ❌ REMOVED (use platform-shared/outbox/scheduler.go)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── search_handler.go          # Search HTTP handlers (POST /search/jobs, /search/freelancers)
│   │       │   ├── recommendation_handler.go  # Recommendation handlers (GET /recommendations/jobs, /recommendations/freelancers)
│   │       │   ├── matching_handler.go        # Matching handlers
│   │       │   ├── feed_handler.go            # Feed handlers
│   │       │   ├── trending_handler.go        # Trending handlers
│   │       │   ├── similarity_handler.go      # Similarity handlers
│   │       │   ├── suggestion_handler.go      # Autocomplete handlers
│   │       │   ├── indexing_handler.go        # Admin endpoint for indexing
│   │       │   ├── taxonomy_handler.go        # 🆕 Taxonomy handlers (skills/categories, aliases)
│   │       │   ├── ltr_handler.go             # 🆕 LTR handlers (ingest signals, preview rerank)
│   │       │   ├── facets_handler.go          # 🆕 Facet handlers (definitions, bands)
│   │       │   ├── personalization_handler.go # 🆕 Personalization handlers (get/update profile)
│   │       │   ├── hygiene_handler.go         # 🆕 Hygiene handlers (incrementals/dedup/visibility)
│   │       │   └── health_handler.go          # Health check endpoints (/health, /ready, /live)
│   │       ├── routes/                         # 🆕 Route registrars mapped 1:1 to handlers
│   │       │   ├── search_routes.go            # /search/*
│   │       │   ├── recommendation_routes.go    # /recommendations/*
│   │       │   ├── matching_routes.go          # /matching/*
│   │       │   ├── feed_routes.go              # /feed/*
│   │       │   ├── trending_routes.go          # /trending/*
│   │       │   ├── similarity_routes.go        # /similarity/*
│   │       │   ├── suggestion_routes.go        # /suggestions/*
│   │       │   ├── indexing_routes.go          # /indexing/*
│   │       │   ├── taxonomy_routes.go          # 🆕 /taxonomy/*
│   │       │   ├── ltr_routes.go               # 🆕 /ltr/*
│   │       │   ├── facets_routes.go            # 🆕 /facets/*
│   │       │   ├── personalization_routes.go   # 🆕 /personalization/*
│   │       │   └── hygiene_routes.go           # 🆕 /hygiene/*
│   │       ├── middleware/
│   │       │   ├── auth.go                     # 📝 UPDATED: Authentication middleware (uses pkg/auth for JWT verification)
│   │       │   ├── rbac.go                     # 📝 UPDATED: RBAC middleware (uses pkg/auth authorizer for role checks)
│   │       │   ├── cors.go                     # 📝 UPDATED: CORS middleware (uses platform-shared/ginx CORS)
│   │       │   ├── rate_limit.go               # Rate limiting middleware (token bucket algorithm)
│   │       │   ├── logging.go                  # 📝 UPDATED: Logging middleware (structured logs per request, uses platform-shared/ginx/logging)
│   │       │   ├── error_handler.go            # Error handling middleware
│   │       │   └── request_id.go               # 📝 UPDATED: Request ID middleware (X-Request-ID header, uses platform-shared/ginx/requestid)
│   │       ├── responses/
│   │       │   ├── success.go                  # 📝 UPDATED: Success response wrappers (use platform-shared/httpx/response.go)
│   │       │   ├── error.go                    # 📝 UPDATED: Error responses (use platform-shared/httpx/errors.go)
│   │       │   └── pagination.go               # 📝 UPDATED: Pagination (use platform-shared/httpx/pagination.go)
│   │       └── router.go                       # 📝 UPDATED: HTTP router setup (uses Gin, applies platform-shared/ginx middleware)
│   │
│   └── config/                                 # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                           # 🆕 Typed Config struct (App, Server, Postgres, Kafka, Redis, Elasticsearch, ML)
│       ├── loader.go                           # 🆕 Config loader using Viper (precedence: flags → env → file → defaults)
│       └── docs/
│           └── CONFIGURATION.md                # 🆕 Configuration documentation (all ENV vars, defaults, examples)
│
├── config/                                     # 🆕 MEDIUM - Configuration files
│   ├── default.yaml                            # 🆕 Default configuration
│   ├── dev.yaml                                # 🆕 Development overrides
│   └── prod.yaml                               # 🆕 Production overrides
│
├── dapr/                                       # 🆕 MEDIUM - Dapr components split by environment
│   ├── local/
│   │   ├── pubsub.yaml                         # Kafka pub/sub component
│   │   └── statestore.yaml                     # State store component
│   └── k8s/
│       ├── pubsub.yaml                         # Kafka with scopes: ["search-be"]
│       ├── statestore.yaml                     # State store with scopes
│       └── secrets.yaml                        # Dapr secret store
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                           # Service-specific errors
│   │   └── codes.go                            # Error codes
│   ├── logger/
│   │   └── logger.go                           # ❌ REMOVED: Use platform-shared/logging
│   ├── utils/
│   │   ├── validator.go                        # Local validation utilities
│   │   ├── text_analyzer.go                    # Text analysis utilities
│   │   ├── normalizer.go                       # Data normalization
│   │   └── vector_math.go                      # Vector math for similarity
│   └── constants/
│       ├── events.go                           # ❌ REMOVED: Use contracts/events
│       ├── topics.go                           # ❌ REMOVED: Use contracts/events
│       └── indices.go                          # Index names constants
│
├── elasticsearch/
│   ├── mappings/
│   │   ├── jobs.json                           # Job index mapping
│   │   └── users.json                          # User index mapping
│   └── analyzers/
│       └── custom_analyzers.json               # Custom analyzers
│
├── ml_models/
│   ├── job_recommendation/
│   │   ├── model.pkl                           # ML model for job recommendations
│   │   ├── features.json                       # Features list
│   │   └── metadata.json                       # Model metadata
│   ├── freelancer_recommendation/
│   │   ├── model.pkl                           # ML model for freelancer recommendations
│   │   ├── features.json                       # Features list
│   │   └── metadata.json                       # Model metadata
│   └── matching/
│       ├── model.pkl                           # ML model for matching
│       ├── features.json                       # Features list
│       └── metadata.json                       # Model metadata
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                     # Kubernetes Deployment
│       ├── service.yaml                        # Kubernetes Service
│       ├── configmap.yaml                      # ConfigMap
│       ├── secrets.yaml                        # Secrets
│       ├── hpa.yaml                            # HPA
│       ├── pdb.yaml                            # PDB
│       └── servicemonitor.yaml                 # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                          # Setup local environment
│   ├── get-secrets.sh                          # Fetch secrets
│   ├── seed-data.sh                            # Seed test data
│   ├── create-indices.sh                       # Create ES indices
│   ├── reindex-all.sh                          # Reindex all data
│   └── train-models.sh                         # Train ML models
│
├── tests/
│   ├── unit/                                   # Unit tests
│   ├── integration/                            # Integration tests
│   └── e2e/                                    # End-to-end tests
│
├── docs/
│   ├── README.md                               # Service overview
│   ├── api.md                                  # API documentation
│   ├── events.md                               # 📝 UPDATED: Events (published: job.indexed, user.indexed, recommendation.generated, consumed: job.posted, user.updated, etc.)
│   ├── search-algorithms.md                    # Search algorithms
│   ├── recommendation-engine.md                # Recommendation engine details
│   ├── recommendation-types.md                 # Recommendation types
│   ├── matching-algorithm.md                   # Matching algorithm
│   ├── ml-models.md                            # ML models documentation
│   ├── elasticsearch-setup.md                  # Elasticsearch setup
│   ├── MIGRATIONS.md                           # 🆕 Migration history
│   ├── SCHEMA.md                               # 🆕 Database schema
│   └── RUNBOOK.md                              # 🆕 Operational procedures
│
├── .github/
│   └── workflows/
│       ├── ci.yml                              # CI workflow
│       └── cd.yml                              # CD workflow
│
├── go.mod                                      # 📝 UPDATED: Imports pkg/auth, platform-shared, contracts/events
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md

```

---

## **📦9️⃣ reviews-be (UPDATED - COVERS RATINGS)**

```
apps/be/reviews-be/
│
├── cmd/
│   └── api/
│       └── main.go                           # 📝 UPDATED: Application entry point - initializes Gin, Dapr, connects to Postgres (now uses loadConfig from internal/config, platform-shared/logging)
│
├── internal/
│   ├── domain/
│   │   ├── review/
│   │   │   ├── entity.go                     # Review aggregate root
│   │   │   ├── rating.go                     # Rating breakdown (overall, quality, communication, etc.)
│   │   │   ├── enums.go                      # ReviewType, Status, ReviewCategory
│   │   │   ├── response.go                   # Review responses
│   │   │   ├── helpful.go                    # Helpful votes
│   │   │   ├── errors.go                     # 🆕 Domain errors (ReviewNotFound, WindowClosed, AlreadySubmitted)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: ReviewSubmitted, ReviewEdited, ReviewPublished, ReviewWithdrawn
│   │   ├── rating/
│   │   │   ├── entity.go                     # Rating system
│   │   │   ├── criteria.go                   # Rating criteria
│   │   │   ├── aggregation.go                # Aggregate ratings
│   │   │   ├── errors.go                     # 🆕 Rating errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: RatingCriteriaUpdated, RatingAggregated
│   │   ├── badge/
│   │   │   ├── entity.go                     # Achievement badges
│   │   │   ├── criteria.go                   # Badge criteria
│   │   │   ├── types.go                      # Badge types (Top Rated, Rising Talent, etc.)
│   │   │   ├── level.go                      # Badge levels
│   │   │   ├── errors.go                     # 🆕 Badge errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: BadgeCreated, BadgeCriteriaUpdated, BadgeLevelUpdated, BadgeArchived
│   │   ├── user_badge/
│   │   │   ├── entity.go                     # User badge assignments
│   │   │   ├── achievement_date.go           # When badge was earned
│   │   │   ├── errors.go                     # 🆕 UserBadge errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: BadgeAwardedToUser, BadgeRevokedFromUser
│   │   ├── reputation/
│   │   │   ├── entity.go                     # Reputation scores
│   │   │   ├── score_calculator.go           # Score calculation
│   │   │   ├── history.go                    # Reputation history
│   │   │   ├── decay.go                      # 🆕 Rolling averages, decay, outlier dampening
│   │   │   ├── top_rated_rules.go            # 🆕 “Top-rated” eligibility rules
│   │   │   ├── errors.go                     # 🆕 Reputation errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: ReputationUpdated, TopRatedEligibilityChanged
│   │   ├── feedback/
│   │   │   ├── entity.go                     # Private feedback
│   │   │   ├── category.go                   # Feedback categories
│   │   │   ├── nps.go                        # 🆕 Confidential NPS-style scores
│   │   │   ├── errors.go                     # 🆕 Feedback errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: PrivateFeedbackSubmitted, PrivateFeedbackUpdated
│   │   ├── flag/
│   │   │   ├── entity.go                     # Flagged reviews
│   │   │   ├── reason.go                     # Flag reasons
│   │   │   ├── errors.go                     # 🆕 Flag errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: ReviewFlagSubmitted, ReviewFlagResolved, ReviewFlagDismissed
│   │   ├── stats/
│   │   │   ├── entity.go                     # Review statistics
│   │   │   ├── aggregates.go                 # Aggregated stats
│   │   │   ├── errors.go                     # 🆕 Stats errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: ReviewStatsUpdated, ReviewStatsRecalculated
│   │   ├── review_reminder/
│   │   │   ├── entity.go                     # Review reminders
│   │   │   ├── errors.go                     # 🆕 Reminder errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: ReviewReminderScheduled, ReviewReminderSent, ReviewReminderFailed
│   │   ├── double_blind/                     # 🆕 Double-blind reviews domain
│   │   │   ├── entity.go                     # Windows (contractID, openAt, closeAt, submitted{client,freelancer})
│   │   │   ├── window.go                     # Publish after both submit or timeout
│   │   │   ├── policy.go                     # Rules (durations, grace periods)
│   │   │   ├── errors.go                     # Double-blind errors (WindowNotOpen, PublishBlocked)
│   │   │   ├── repository.go                 # DoubleBlindRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: DoubleBlindSubmitted, DoubleBlindPublished, DoubleBlindTimeout
│   │   ├── weighting/                        # 🆕 Dimensions & weights
│   │   │   ├── entity.go                     # Weight profiles per category (quality, comms, timeliness)
│   │   │   ├── profile.go                    # Category-specific weighting resolution
│   │   │   ├── errors.go                     # Weighting errors (ProfileNotFound)
│   │   │   ├── repository.go                 # WeightingRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: WeightProfileCreated, WeightProfileUpdated
│   │   ├── public_response/                  # 🆕 Public response rights
│   │   │   ├── entity.go                     # Right-of-reply records (reviewID, authorID, body, postedAt)
│   │   │   ├── policy.go                     # Response windows/length limits
│   │   │   ├── errors.go                     # Public response errors
│   │   │   ├── repository.go                 # PublicResponseRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: PublicResponsePosted, PublicResponseEdited, PublicResponseRemoved
│   │   ├── moderation/                       # 🆕 Abuse controls
│   │   │   ├── entity.go                     # Moderation states (Pending, UnderReview, Removed, Restored)
│   │   │   ├── auto_flag.go                  # Auto-flags: swapping, bursts, suspicious patterns
│   │   │   ├── signals.go                    # Detection signals
│   │   │   ├── errors.go                     # Moderation errors
│   │   │   ├── repository.go                 # ModerationRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: ReviewAbuseAutoFlagged, ReviewModerationStateChanged
│   │   └── outbox/
│   │       ├── entity.go                     # ❌ REMOVED (use platform-shared/outbox/entity.go)
│   │       └── repository.go                 # ❌ REMOVED (use platform-shared/outbox/repository.go)
│   │
│   ├── application/
│   │   ├── eventhandler/
│   │   │   ├── contract_handler.go            # Consumes: contract.ended
│   │   │   │                                     # Purpose: opens the double-blind review window for both parties at contract end.
│   │   │   │
│   │   │   ├── payment_handler.go             # Consumes: payment.captured
│   │   │   │                                     # Purpose: starts/qualifies the double-blind window when payment is captured (hourly/fixed).
│   │   │   │
│   │   │   ├── admin_moderation_handler.go    # Consumes: admin.moderation.actioned (approve/reject/remove/hide)
│   │   │   │                                     # Purpose: reflect moderation decisions on reviews (remove/restore visibility, recalc aggregates if needed).
│   │   │   │
│   │   │   ├── badge_handler.go               # Consumes: badge.criteria.met  # (optional downstream hook)
│   │   │   │                                     # Purpose: optional: feed reputation/badge eligibility workflows that affect review surfaces.
│   │   │   │
│   │   │   └── admin_flags_handler.go         # Consumes: admin.feature_flag.updated | admin.threshold.updated | admin.experiment.updated
│   │   │                                         # Purpose: refresh runtime toggles/config used by review windows, moderation thresholds,
│   │   │                                         #           decay/outlier dampening, and publishing policies.
│   │   ├── review/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Create, Update, Delete, Respond
│   │   │   ├── queries.go                    # Get, List, Filter
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── business_rules.go
│   │   ├── rating/
│   │   │   ├── service.go
│   │   │   ├── aggregator.go                 # Aggregate ratings
│   │   │   ├── calculator.go                 # Calculate average ratings
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                   # 🆕 UpsertCriteria, RecalculateForContract
│   │   │   ├── queries.go                    # 🆕 GetCriteria, GetAggregates
│   │   │   └── validators.go                 # 🆕 Validate criteria weights & bounds
│   │   ├── badge/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Award, Revoke badges
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── achievement_checker.go        # Check badge criteria
│   │   ├── reputation/
│   │   │   ├── service.go
│   │   │   ├── calculator.go                 # Calculate reputation
│   │   │   ├── updater.go                    # Update reputation
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                   # 🆕 RecomputeUserReputation, ApplyDecay
│   │   │   ├── queries.go                    # 🆕 GetReputation, GetHistory
│   │   │   └── validators.go                 # 🆕 Validate decay/outlier params
│   │   ├── feedback/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                   # 🆕 SubmitPrivateFeedback
│   │   │   ├── queries.go                    # 🆕 GetPrivateFeedback
│   │   │   └── validators.go                 # 🆕 Validate NPS bounds & confidentiality
│   │   ├── stats/
│   │   │   ├── service.go
│   │   │   ├── aggregator.go                 # Aggregate stats
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                   # 🆕 RebuildStats
│   │   │   ├── queries.go                    # 🆕 GetStats
│   │   │   └── validators.go                 # 🆕 Validate windows & min samples
│   │   ├── double_blind/                     # 🆕 App layer for double-blind reviews
│   │   │   ├── service.go                    # OpenWindow, SubmitSide, PublishIfReady
│   │   │   ├── commands.go                   # OpenWindow, SubmitClient, SubmitFreelancer, ForcePublish
│   │   │   ├── queries.go                    # GetWindow, GetPublishState
│   │   │   └── validators.go                 # Validate window times & state transitions
│   │   ├── weighting/                        # 🆕 App layer for dimensions & weights
│   │   │   ├── service.go                    # ResolveWeightProfile, UpsertProfile
│   │   │   ├── commands.go                   # UpsertWeightProfile
│   │   │   ├── queries.go                    # GetWeightProfile, ListProfiles
│   │   │   └── validators.go                 # Validate weight sums & ranges
│   │   ├── public_response/                  # 🆕 App layer for right-of-reply
│   │   │   ├── service.go                    # PostResponse, UpdateResponse
│   │   │   ├── commands.go                   # CreateResponse, EditResponse
│   │   │   ├── queries.go                    # GetResponse
│   │   │   └── validators.go                 # Validate length & window
│   │   ├── moderation/                       # 🆕 Abuse controls
│   │   │   ├── service.go                    # AutoFlag, ReviewModeration, ChangeState
│   │   │   ├── commands.go                   # AutoFlagReview, SetModerationState
│   │   │   ├── queries.go                    # GetModerationState, ListFlags
│   │   │   └── validators.go                 # Validate reasons & thresholds
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection setup (DSN from config, connection pooling)
│   │   │       ├── transaction.go            # Transaction helpers (Begin, Commit, Rollback, WithTransaction wrapper)
│   │   │       ├── migrations.go             # 📝 UPDATED: Auto-migrate logic (now with version tracking, GORM AutoMigrate for all tables)
│   │   │       ├── version.go                # 🆕 Schema version tracking (SchemaVersion table, RecordMigration function)
│   │   │       ├── safety.go                 # 🆕 Pre-migration safety checks (environment validation, disk space check)
│   │   │       ├── review_repository.go
│   │   │       ├── rating_repository.go
│   │   │       ├── badge_repository.go
│   │   │       ├── user_badge_repository.go
│   │   │       ├── reputation_repository.go
│   │   │       ├── feedback_repository.go
│   │   │       ├── flag_repository.go
│   │   │       ├── stats_repository.go
│   │   │       ├── review_reminder_repository.go
│   │   │       ├── double_blind_repository.go # 🆕 DoubleBlindRepository implementation
│   │   │       ├── weighting_repository.go    # 🆕 WeightingRepository implementation
│   │   │       ├── public_response_repository.go # 🆕 PublicResponseRepository implementation
│   │   │       └── moderation_repository.go   # 🆕 ModerationRepository implementation
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go             # Redis connection setup (connection pooling, retry logic)
│   │   │       ├── review_cache.go           # Cache reviews (Get, Set, Invalidate with TTL)
│   │   │       ├── badge_cache.go            # Cache badges
│   │   │       ├── reputation_cache.go       # Cache reputation scores
│   │   │       ├── stats_cache.go            # Cache stats
│   │   │       ├── double_blind_cache.go     # 🆕 Cache windows/publish state
│   │   │       └── weighting_cache.go        # 🆕 Cache weight profiles
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go               # 📝 UPDATED: Kafka consumer (now uses platform-shared/inbox for message deduplication)
│   │   │       ├── producer.go               # 📝 UPDATED: Kafka producer (now uses platform-shared/outbox for reliable publishing)
│   │   │       ├── topics.go                 # 📝 UPDATED: Topic constants (imported from contracts/events - review.submitted, review.published, badge.awarded)
│   │   │       └── scram.go                  # SCRAM authentication for Kafka (SASL/SCRAM-SHA-256)
│   │   └── outbox/
│   │       ├── processor.go                  # ❌ REMOVED (use platform-shared/outbox/forwarder.go)
│   │       └── scheduler.go                  # ❌ REMOVED (use platform-shared/outbox/scheduler.go)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── review_handler.go         # Review HTTP handlers (POST /reviews, GET /reviews/:id)
│   │       │   ├── rating_handler.go         # Rating handlers
│   │       │   ├── badge_handler.go          # Badge handlers
│   │       │   ├── reputation_handler.go     # Reputation handlers
│   │       │   ├── feedback_handler.go       # Feedback handlers
│   │       │   ├── stats_handler.go          # Stats handlers
│   │       │   ├── flag_handler.go           # Flag handlers
│   │       │   ├── double_blind_handler.go   # 🆕 Double-blind endpoints
│   │       │   ├── weighting_handler.go      # 🆕 Weight profiles endpoints
│   │       │   ├── public_response_handler.go# 🆕 Right-of-reply endpoints
│   │       │   └── moderation_handler.go     # 🆕 Moderation endpoints
│   │       ├── routes/                       # 🆕 Route registrars mapped 1:1 to handlers
│   │       │   ├── review_routes.go          # /reviews/*
│   │       │   ├── rating_routes.go          # /ratings/*
│   │       │   ├── badge_routes.go           # /badges/*
│   │       │   ├── reputation_routes.go      # /reputation/*
│   │       │   ├── feedback_routes.go        # /feedback/*
│   │       │   ├── stats_routes.go           # /stats/*
│   │       │   ├── flag_routes.go            # /flags/*
│   │       │   ├── double_blind_routes.go    # 🆕 /double-blind/*
│   │       │   ├── weighting_routes.go       # 🆕 /weights/*
│   │       │   ├── public_response_routes.go # 🆕 /responses/*
│   │       │   └── moderation_routes.go      # 🆕 /moderation/*
│   │       ├── middleware/
│   │       │   ├── auth.go                   # 📝 UPDATED: Authentication middleware (uses pkg/auth for JWT verification)
│   │       │   ├── rbac.go                   # 📝 UPDATED: RBAC middleware (uses pkg/auth authorizer for role checks)
│   │       │   ├── cors.go                   # 📝 UPDATED: CORS middleware (uses platform-shared/ginx CORS)
│   │       │   ├── rate_limit.go             # Rate limiting middleware (token bucket algorithm)
│   │       │   ├── logging.go                # 📝 UPDATED: Logging middleware (structured logs per request, uses platform-shared/ginx/logging)
│   │       │   ├── error_handler.go          # Error handling middleware
│   │       │   └── request_id.go             # 📝 UPDATED: Request ID middleware (X-Request-ID header, uses platform-shared/ginx/requestid)
│   │       ├── responses/
│   │       │   ├── success.go                # 📝 UPDATED: Success response wrappers (use platform-shared/httpx/response.go)
│   │       │   ├── error.go                  # 📝 UPDATED: Error responses (use platform-shared/httpx/errors.go)
│   │       │   └── pagination.go             # 📝 UPDATED: Pagination (use platform-shared/httpx/pagination.go)
│   │       └── router.go                     # 📝 UPDATED: HTTP router setup (uses Gin, applies platform-shared/ginx middleware)
│   │
│   └── config/                               # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                         # 🆕 Typed Config struct (App, Server, Postgres, Kafka, Redis)
│       ├── loader.go                         # 🆕 Config loader using Viper (precedence: flags → env → file → defaults)
│       └── docs/
│           └── CONFIGURATION.md              # 🆕 Configuration documentation (all ENV vars, defaults, examples)
│
├── config/
│   ├── default.yaml                          # 🆕 Default configuration
│   ├── dev.yaml                              # 🆕 Development overrides
│   └── prod.yaml                             # 🆕 Production overrides
│
├── dapr/
│   ├── local/
│   │   ├── pubsub.yaml                       # Kafka pub/sub component
│   │   └── statestore.yaml                   # State store component
│   └── k8s/
│       ├── pubsub.yaml                       # Kafka with scopes: ["reviews-be"]
│       ├── statestore.yaml                   # State store with scopes
│       └── secrets.yaml                      # Dapr secret store
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                         # Service-specific errors
│   │   └── codes.go                          # Error codes
│   ├── logger/
│   │   └── logger.go                         # ❌ REMOVED: Use platform-shared/logging
│   ├── utils/
│   │   ├── validator.go                      # Local validation utilities
│   │   ├── sanitizer.go                      # Sanitize review content (prevent XSS)
│   │   └── sentiment_analyzer.go             # Sentiment analysis for reviews
│   └── constants/
│       ├── events.go                         # ❌ REMOVED: Use contracts/events
│       ├── topics.go                         # ❌ REMOVED: Use contracts/events
│       └── badges.go                         # Badge constants
│
├── seeds/
│   └── badges.sql                            # Seed initial badges
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   ├── seed-badges.sh
│   └── recalculate-reputation.sh
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── api.md
│   ├── events.md                             # 📝 UPDATED: (review.submitted, review.published, badge.awarded, reputation.updated)
│   ├── badge-system.md
│   ├── reputation-algorithm.md
│   ├── rating-system.md
│   ├── double-blind.md                       # 🆕 Double-blind policy & flows
│   ├── moderation.md                          # 🆕 Abuse controls & moderation states
│   ├── MIGRATIONS.md                          # 🆕
│   ├── SCHEMA.md                              # 🆕
│   └── RUNBOOK.md                             # 🆕
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md

```


---

## **📦🔟 subscriptions-be (UPDATED)**

```
apps/be/subscriptions-be/
│
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── domain/
│   │   ├── plan/
│   │   │   ├── entity.go                     # Subscription plans
│   │   │   ├── features.go                   # Plan features
│   │   │   ├── pricing.go                    # Pricing tiers
│   │   │   ├── limits.go                     # Plan limits
│   │   │   ├── errors.go                     # 🆕 Domain errors (PlanNotFound, InvalidFeature, LimitExceeded)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: plan.created / plan.updated / plan.archived
│   │   ├── subscription/
│   │   │   ├── entity.go                     # User subscriptions
│   │   │   ├── billing_cycle.go              # Billing cycle management
│   │   │   ├── enums.go                      # Status, Type
│   │   │   ├── auto_renewal.go               # Auto-renewal logic
│   │   │   ├── errors.go                     # 🆕 Subscription errors (AlreadyActive, PastDue, Canceled)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: subscription.created / subscription.changed / subscription.paused / subscription.canceled / subscription.renewed
│   │   ├── connect/
│   │   │   ├── entity.go                     # Connects/Credits
│   │   │   ├── package.go                    # Connect packages
│   │   │   ├── transaction.go                # Connect transactions
│   │   │   ├── balance.go                    # Balance tracking
│   │   │   ├── expiry.go                     # 🆕 Expiry & rollover rules (effective/expiry windows)
│   │   │   ├── grant.go                      # 🆕 Promotional grants (campaign/ref IDs)
│   │   │   ├── errors.go                     # 🆕 Connect errors (InsufficientBalance, Expired, RolloverNotAllowed)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: connects.purchased / connects.used / connects.expired / connects.granted
│   │   ├── usage/
│   │   │   ├── entity.go                     # Usage tracking
│   │   │   ├── quota.go                      # Usage quotas
│   │   │   ├── limit.go                      # Usage limits
│   │   │   ├── meter.go                      # 🆕 Per-feature meters (messages_to_non_hires, boosts, invites)
│   │   │   ├── errors.go                     # 🆕 Usage errors (LimitReached, CounterNotFound)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: usage.meter.incremented / usage.limit.reached
│   │   ├── entitlement/                      # 🆕 Plans & entitlements / feature flags
│   │   │   ├── entity.go                     # Entitlement (User/OrgID, PlanID, FeatureKey, Allowed, Scope)
│   │   │   ├── rules.go                      # Rule evaluation (plan→features, addons, promos)
│   │   │   ├── errors.go                     # 🆕 Entitlement errors (FeatureDenied, NotInPlan)
│   │   │   ├── repository.go                 # EntitlementRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: entitlement.feature.enabled / entitlement.feature.disabled
│   │   ├── seat_billing/                     # 🆕 Seat-based billing (clients/agencies)
│   │   │   ├── entity.go                     # Seat allocation (SubscriptionID, Seats, Assigned, OveragePolicy)
│   │   │   ├── overage.go                    # Overage calculations (per-seat over cap)
│   │   │   ├── proration.go                  # Proration rules for mid-cycle changes
│   │   │   ├── invoice_export.go             # Invoice export model (→ Financial-BE)
│   │   │   ├── errors.go                     # 🆕 Seat billing errors
│   │   │   ├── repository.go                 # SeatBillingRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: billing.proration.applied / billing.invoice.exported / seat.overage.incurred / seat.allocated
│   │   ├── addon/
│   │   │   ├── entity.go                     # Plan add-ons
│   │   │   ├── errors.go                     # 🆕 Addon errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: addon.added / addon.removed / addon.updated
│   │   ├── promotion/
│   │   │   ├── entity.go                     # Promotional codes
│   │   │   ├── discount.go                   # Discount rules
│   │   │   ├── usage_limit.go                # Usage limits for promos
│   │   │   ├── errors.go                     # 🆕 Promotion errors (InvalidCode, Exhausted, Ineligible)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: promo.created / promo.applied / promo.revoked / promo.exhausted
│   │   ├── trial/
│   │   │   ├── entity.go                     # Free trials
│   │   │   ├── eligibility.go                # Trial eligibility
│   │   │   ├── errors.go                     # 🆕 Trial errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: trial.started / trial.ended / trial.eligibility.updated
│   │   ├── billing_history/
│   │   │   ├── entity.go                     # Billing history
│   │   │   ├── errors.go                     # 🆕 Billing history errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: billing.invoice.generated / billing.payment.applied / billing.credit.issued
│   │   ├── feature_toggle/
│   │   │   ├── entity.go                     # Feature toggles per plan
│   │   │   ├── errors.go                     # 🆕 Feature toggle errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.feature_flag.updated (plan scope) / feature.toggle.enabled / feature.toggle.disabled
│   │   └── outbox/
│   │       ├── entity.go                     # ❌ REMOVED (use platform-shared/outbox/entity.go)
│   │       └── repository.go                 # ❌ REMOVED (use platform-shared/outbox/repository.go)
│   │
│   ├── application/
│   │   ├── eventhandler/
│   │   │   ├── financial_handler.go           # Consumes: payment.processed | payment.failed    (from financial-be)
│   │   │   │                                     # Purpose: post-payment workflows (activate/renew/pause as needed),
│   │   │   │                                     #          dunning/reactivation on failures.
│   │   │   │
│   │   │   ├── proposal_handler.go            # Consumes: proposal.submitted            (from proposals-be)
│   │   │   │                                     # Purpose: deduct connects / enforce submit limits based on plan/entitlements.
│   │   │   │
│   │   │   ├── job_handler.go                 # Consumes: job.posted                    (from jobs-be)
│   │   │   │                                     # Purpose: check posting limits, apply plan gating, consume connects if required.
│   │   │   │
│   │   │   └── admin_flags_handler.go         # Consumes: admin.feature_flag.updated |
│   │   │                                         #           admin.threshold.updated |
│   │   │                                         #           admin.experiment.updated          (from admin-be)
│   │   │                                         # Purpose: refresh runtime flags/thresholds/experiments that gate
│   │   │                                         #          features like boosts, extra invites, posting quotas.
│   │   ├── plan/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Create, Update plans
│   │   │   ├── queries.go                    # Get, List plans
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go                 # 🆕 Validate features/limits/pricing
│   │   ├── subscription/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Subscribe, Cancel, Change, Pause
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── lifecycle_manager.go          # Manage subscription lifecycle
│   │   │   ├── renewal_manager.go            # Handle renewals
│   │   │   └── validators.go                 # 🆕 Validate status transitions, proration inputs
│   │   ├── connect/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Purchase, Use, Transfer, Refund
│   │   │   ├── queries.go                    # Get balance, History
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── calculator.go                 # Calculate connect costs
│   │   │   └── validators.go                 # 🆕 Validate expiries/rollovers/amounts
│   │   ├── usage/
│   │   │   ├── service.go
│   │   │   ├── tracker.go                    # Track usage
│   │   │   ├── quota_checker.go              # Check quotas
│   │   │   ├── limiter.go                    # Enforce limits
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go                 # 🆕 Validate counters & limits
│   │   ├── entitlement/                      # 🆕 Plans & entitlements
│   │   │   ├── service.go                    # Resolve entitlements, check feature access
│   │   │   ├── commands.go                   # GrantEntitlement, RevokeEntitlement (admin/system)
│   │   │   ├── queries.go                    # GetEntitlements, CheckFeature
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go                 # Validate feature scopes and subjects
│   │   ├── seat_billing/                     # 🆕 Seats + overages + proration + export
│   │   │   ├── service.go                    # Allocate seats, compute overages, proration, export invoice
│   │   │   ├── commands.go                   # AssignSeat, ReleaseSeat, SetSeatCap, RecalculateOverages
│   │   │   ├── queries.go                    # GetSeatSummary, GetOverages
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go                 # Validate seat counts and policies
│   │   ├── addon/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                   # 🆕 AddAddon, RemoveAddon
│   │   │   ├── queries.go                    # 🆕 GetAddons, GetAddon
│   │   │   └── validators.go                 # 🆕 Validate addon compatibility
│   │   ├── promotion/
│   │   │   ├── service.go
│   │   │   ├── validator.go                  # Validate promo codes (kept as-is)
│   │   │   ├── validators.go                 # 🆕 Aliases to keep naming consistency
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── trial/
│   │   │   ├── service.go
│   │   │   ├── eligibility_checker.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                   # 🆕 StartTrial, EndTrial
│   │   │   ├── queries.go                    # 🆕 GetTrial, IsEligible
│   │   │   └── validators.go                 # 🆕 Validate trial eligibility/rules
│   │   ├── billing/
│   │   │   ├── service.go
│   │   │   ├── invoice_generator.go          # Generate invoices
│   │   │   ├── payment_processor.go          # Process payments
│   │   │   ├── exporter.go                   # 🆕 Export invoices to Financial-BE
│   │   │   ├── dto.go
│   │   │   ├── commands.go                   # 🆕 GenerateInvoice, ExportInvoice, CapturePayment
│   │   │   ├── queries.go                    # 🆕 GetInvoice, ListInvoices
│   │   │   └── validators.go                 # 🆕 Validate seat counts, proration, exports
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go
│   │   │       ├── transaction.go
│   │   │       ├── migrations.go
│   │   │       ├── version.go
│   │   │       ├── safety.go
│   │   │       ├── plan_repository.go
│   │   │       ├── subscription_repository.go
│   │   │       ├── connect_repository.go
│   │   │       ├── usage_repository.go
│   │   │       ├── entitlement_repository.go      # 🆕 Implementation
│   │   │       ├── seat_billing_repository.go     # 🆕 Implementation
│   │   │       ├── addon_repository.go
│   │   │       ├── promotion_repository.go
│   │   │       ├── trial_repository.go
│   │   │       ├── billing_history_repository.go
│   │   │       ├── feature_toggle_repository.go
│   │   │       └── outbox_repository.go           # ❌ REMOVED (use platform-shared/outbox/postgres/repository.go)
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go
│   │   │       ├── subscription_cache.go
│   │   │       ├── plan_cache.go
│   │   │       ├── connect_cache.go
│   │   │       ├── entitlement_cache.go           # 🆕 Cache entitlements
│   │   │       └── feature_toggle_cache.go
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go
│   │   │       ├── producer.go
│   │   │       ├── topics.go
│   │   │       └── scram.go
│   │   ├── scheduler/
│   │   │   ├── cron.go
│   │   │   └── jobs.go
│   │   ├── payment/
│   │   │   └── client.go                         # Financial-BE client (unchanged)
│   │   ├── billing_export/                        # 🆕 Explicit Financial-BE export client
│   │   │   ├── client.go                          # REST/gRPC client for invoice export
│   │   │   └── mapper.go                          # Map local invoices/overages → Financial-BE DTO
│   │   └── outbox/
│   │       ├── processor.go                       # ❌ REMOVED
│   │       └── scheduler.go                       # ❌ REMOVED
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── plan_handler.go
│   │       │   ├── subscription_handler.go
│   │       │   ├── connect_handler.go
│   │       │   ├── usage_handler.go
│   │       │   ├── entitlement_handler.go        # 🆕 Entitlement endpoints
│   │       │   ├── seat_billing_handler.go       # 🆕 Seat mgmt & overages
│   │       │   ├── addon_handler.go
│   │       │   ├── promotion_handler.go
│   │       │   ├── trial_handler.go
│   │       │   ├── billing_handler.go
│   │       │   └── health_handler.go
│   │       ├── middleware/
│   │       │   ├── auth.go
│   │       │   ├── rbac.go
│   │       │   ├── cors.go
│   │       │   ├── rate_limit.go
│   │       │   ├── logging.go
│   │       │   ├── error_handler.go
│   │       │   ├── request_id.go
│   │       │   └── feature_gate.go
│   │       ├── responses/
│   │       │   ├── success.go
│   │       │   ├── error.go
│   │       │   └── pagination.go
│   │       ├── routes/                           # Route registrars mapped 1:1 to handlers
│   │       │   ├── plan_routes.go                # /plans/*
│   │       │   ├── subscription_routes.go        # /subscriptions/*
│   │       │   ├── connect_routes.go             # /connects/*
│   │       │   ├── usage_routes.go               # /usage/*
│   │       │   ├── entitlement_routes.go         # 🆕 /entitlements/*
│   │       │   ├── seat_billing_routes.go        # 🆕 /seats/*
│   │       │   ├── addon_routes.go               # /addons/*
│   │       │   ├── promotion_routes.go           # /promotions/*
│   │       │   ├── trial_routes.go               # /trials/*
│   │       │   └── billing_routes.go             # /billing/*
│   │       └── router.go
│   │
│   └── config/
│       ├── schema.go
│       ├── loader.go
│       └── docs/
│           └── CONFIGURATION.md
│
├── config/
│   ├── default.yaml
│   ├── dev.yaml
│   └── prod.yaml
│
├── dapr/
│   ├── local/
│   │   ├── pubsub.yaml
│   │   └── statestore.yaml
│   └── k8s/
│       ├── pubsub.yaml
│       ├── statestore.yaml
│       └── secrets.yaml
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   └── codes.go
│   ├── logger/
│   │   └── logger.go
│   ├── utils/
│   │   ├── validator.go
│   │   ├── billing_calculator.go
│   │   └── proration.go
│   └── constants/
│       ├── events.go
│       ├── topics.go
│       ├── plans.go
│       └── features.go
│
├── seeds/
│   ├── plans.sql
│   └── connect_packages.sql
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       ├── cronjob-renewal.yaml
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   ├── seed-plans.sh
│   ├── seed-data.sh
│   └── process-renewals.sh
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── api.md
│   ├── events.md
│   ├── subscription-plans.md
│   ├── connects-system.md
│   ├── billing-logic.md
│   ├── feature-toggles.md
│   ├── MIGRATIONS.md
│   ├── SCHEMA.md
│   └── RUNBOOK.md
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md

```


---

## **📦1️⃣1️⃣ admin-be (NEW - COMPREHENSIVE**

```
apps/be/admin-be/
│
├── cmd/
│   └── api/
│       └── main.go                           # 📝 UPDATED: Application entry point - initializes Gin, Dapr, connects to Postgres (now uses loadConfig from internal/config, platform-shared/logging)
│
├── internal/
│   ├── domain/
│   │   ├── admin_user/
│   │   │   ├── entity.go                     # Admin user accounts
│   │   │   ├── role.go                       # Admin roles (SuperAdmin, Moderator, Support)
│   │   │   ├── permission.go                 # Granular permissions
│   │   │   ├── activity_log.go               # Admin activity tracking
│   │   │   ├── errors.go                     # 🆕 Domain errors (AdminUserNotFound, PermissionDenied)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.user.created / admin.user.updated / admin.user.role.changed / admin.user.permissions.updated / admin.user.login.logged
│   │   ├── support_ticket/
│   │   │   ├── entity.go                     # Support tickets
│   │   │   ├── priority.go                   # Priority levels (Low, Medium, High, Urgent)
│   │   │   ├── status.go                     # Status (Open, InProgress, Resolved, Closed)
│   │   │   ├── category.go                   # Ticket categories
│   │   │   ├── assignment.go                 # Agent assignment
│   │   │   ├── errors.go                     # 🆕 Ticket errors (TicketNotFound, InvalidStatus)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.case.opened / admin.case.updated / admin.case.assigned / admin.case.closed
│   │   ├── ticket_message/
│   │   │   ├── entity.go                     # Ticket conversation messages
│   │   │   ├── attachment.go                 # Message attachments
│   │   │   ├── errors.go                     # 🆕 Message errors (MessageNotFound, AttachmentTooLarge)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: ticket.note.added / ticket.note.edited / ticket.note.deleted / ticket.attachment.added
│   │   ├── support_agent/
│   │   │   ├── entity.go                     # Support agent profiles
│   │   │   ├── availability.go               # Agent availability/online status
│   │   │   ├── stats.go                      # Agent performance stats
│   │   │   ├── errors.go                     # 🆕 Agent errors (AgentNotFound, AgentUnavailable)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: support.agent.status.changed / support.agent.assigned / support.agent.workload.updated
│   │   ├── canned_response/
│   │   │   ├── entity.go                     # Predefined responses
│   │   │   ├── category.go                   # Response categories
│   │   │   ├── errors.go                     # 🆕 Canned response errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: canned_response.created / canned_response.updated / canned_response.archived
│   │   ├── knowledge_base/
│   │   │   ├── entity.go                     # KB articles
│   │   │   ├── category.go                   # Article categories
│   │   │   ├── tag.go                        # Article tags
│   │   │   ├── version.go                    # Article versioning
│   │   │   ├── errors.go                     # 🆕 KB errors (ArticleNotFound, VersionConflict)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: kb.article.created / kb.article.updated / kb.article.versioned / kb.article.published / kb.article.unpublished
│   │   ├── faq/
│   │   │   ├── entity.go                     # Frequently asked questions
│   │   │   ├── category.go                   # FAQ categories
│   │   │   ├── errors.go                     # 🆕 FAQ errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: faq.created / faq.updated / faq.published / faq.unpublished
│   │   ├── moderation_queue/
│   │   │   ├── entity.go                     # Moderation queue items
│   │   │   ├── content_type.go               # Job, User, Review, Message, etc.
│   │   │   ├── flag_reason.go                # Reasons for flagging
│   │   │   ├── action.go                     # Moderation actions taken
│   │   │   ├── errors.go                     # 🆕 Moderation errors (QueueItemNotFound)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.moderation.enqueued / admin.moderation.assigned / admin.moderation.state.changed
│   │   ├── user_action/
│   │   │   ├── entity.go                     # Admin actions on users
│   │   │   ├── action_type.go                # Suspend, Ban, Verify, Warn, etc.
│   │   │   ├── reason.go                     # Action reasons
│   │   │   ├── errors.go                     # 🆕 User action errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.user_action.applied (suspend/ban/verify/warn) / admin.user_action.reversed
│   │   ├── content_action/
│   │   │   ├── entity.go                     # Admin actions on content
│   │   │   ├── action_type.go                # Remove, Hide, Approve, Reject
│   │   │   ├── errors.go                     # 🆕 Content action errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.content.actioned (approve/reject/remove/hide)
│   │   ├── dispute_resolution/
│   │   │   ├── entity.go                     # Dispute cases
│   │   │   ├── evidence.go                   # Evidence submitted
│   │   │   ├── decision.go                   # Admin decision
│   │   │   ├── errors.go                     # 🆕 Dispute errors (DisputeNotFound)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.case.opened (dispute) / admin.case.updated / admin.case.closed / admin.dispute.decision.made
│   │   ├── system_config/
│   │   │   ├── entity.go                     # System configuration
│   │   │   ├── feature_flag.go               # Feature flags
│   │   │   ├── maintenance.go                # Maintenance mode settings
│   │   │   ├── errors.go                     # 🆕 Config errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.feature_flag.updated / admin.threshold.updated / admin.experiment.updated / admin.maintenance.window.set
│   │   ├── announcement/
│   │   │   ├── entity.go                     # Platform announcements
│   │   │   ├── target.go                     # Target audience (All, Freelancers, Clients)
│   │   │   ├── schedule.go                   # Scheduled announcements
│   │   │   ├── errors.go                     # 🆕 Announcement errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.announcement.created / admin.announcement.scheduled / admin.announcement.published / admin.announcement.canceled
│   │   ├── report/
│   │   │   ├── entity.go                     # Generated reports
│   │   │   ├── report_type.go                # Report types (Users, Revenue, Activity)
│   │   │   ├── schedule.go                   # Scheduled reports
│   │   │   ├── errors.go                     # 🆕 Report errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.report.requested / admin.report.generated / admin.report.delivery.failed
│   │   ├── analytics_dashboard/
│   │   │   ├── entity.go                     # Dashboard configurations
│   │   │   ├── widget.go                     # Dashboard widgets
│   │   │   ├── metric.go                     # Custom metrics
│   │   │   ├── errors.go                     # 🆕 Analytics dashboard errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.dashboard.created / admin.dashboard.updated / admin.dashboard.widget.added / admin.dashboard.widget.removed
│   │   ├── audit_log/
│   │   │   ├── entity.go                     # Complete audit trail
│   │   │   ├── action.go                     # Actions performed
│   │   │   ├── resource.go                   # Resources affected
│   │   │   ├── errors.go                     # 🆕 Audit log errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.audit.logged (staff action appended)
│   │   ├── notification_blast/
│   │   │   ├── entity.go                     # Bulk notifications/emails
│   │   │   ├── audience.go                   # Target audience
│   │   │   ├── schedule.go                   # Scheduled blasts
│   │   │   ├── errors.go                     # 🆕 Blast errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.blast.scheduled / admin.blast.sent / admin.blast.failed / admin.blast.canceled
│   │   ├── platform_stats/
│   │   │   ├── entity.go                     # Platform statistics
│   │   │   ├── realtime.go                   # Real-time stats
│   │   │   ├── errors.go                     # 🆕 Platform stats errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.platform.stats.snapshot.created / admin.platform.stats.realtime.updated
│   │   └── outbox/
│   │       ├── entity.go                     # ❌ REMOVED (use platform-shared/outbox/entity.go)
│   │       └── repository.go                 # ❌ REMOVED (use platform-shared/outbox/repository.go)
│   │
│   │   # --------------------------- NEW DOMAINS (ADD-ONLY) ---------------------------
│   │   ├── fraud_review/                     # 🆕 Case mgmt: fraud reviews
│   │   │   ├── entity.go
│   │   │   ├── status.go                     # Pending, Investigating, Cleared, Confirmed
│   │   │   ├── severity.go                   # Low, Medium, High, Critical
│   │   │   ├── reason.go                     # SuspiciousActivity, StolenCard, etc.
│   │   │   ├── notes.go                      # Investigator notes
│   │   │   ├── sla.go                        # SLA definitions for review steps
│   │   │   ├── errors.go                     # 🆕 FraudReviewNotFound, InvalidTransition
│   │   │   └── repository.go
│   │   ├── user_report/                      # 🆕 Case mgmt: user reports
│   │   │   ├── entity.go                     # Reports filed by users about users/content
│   │   │   ├── category.go                   # Spam, Harassment, IPViolation, etc.
│   │   │   ├── status.go                     # Open, Triaged, Actioned, Dismissed
│   │   │   ├── notes.go                      # Moderator notes
│   │   │   ├── sla.go                        # SLA targets for first-response/resolve
│   │   │   ├── errors.go                     # 🆕 UserReportNotFound, InvalidCategory
│   │   │   └── repository.go
│   │   ├── risk_management/                  # 🆕 Risk dashboards
│   │   │   ├── hold.go                       # Payment/account holds
│   │   │   ├── reserve.go                    # Rolling/reserve configs
│   │   │   ├── chargeback.go                 # Chargeback cases & states
│   │   │   ├── velocity_alert.go             # Velocity rule hits
│   │   │   ├── country_rate_anomaly.go       # Country/rate anomaly entity
│   │   │   ├── errors.go                     # 🆕 RiskItemNotFound, RuleConflict
│   │   │   └── repository.go
│   │   ├── ticket_note/                      # 🆕 Case mgmt: internal ticket notes
│   │   │   ├── entity.go
│   │   │   ├── errors.go                     # 🆕 TicketNoteNotFound
│   │   │   └── repository.go
│   │   └── data_export/                      # 🆕 Audit & export (GDPR/CCPA)
│   │       ├── entity.go                     # Export job (type, requester, scope)
│   │       ├── request.go                    # Data subject request metadata
│   │       ├── status.go                     # Requested, Approved, Generating, Delivered
│   │       ├── format.go                     # ZIP, CSV, JSONL, Parquet
│   │       ├── errors.go                     # 🆕 ExportNotFound, AccessDenied
│   │       └── repository.go
│   │   # ---------------------------------------------------------------------------
│   │
│   ├── application/
│   │   ├── eventhandler/
│   │   │   ├── storage_handler.go             # Consumes (from storage-be):
│   │   │   │                                     #   file.policy.updated             — update moderation/DLP config views;
│   │   │   │                                     #   file.policy.violation.detected    — open/route moderation case; risk signal;
│   │   │   │                                     #   file.lifecycle.soft_deleted       — reflect visibility change in dashboards;
│   │   │   │                                     #   file.lifecycle.restored           — restore visibility state;
│   │   │   │                                     #   file.lifecycle.legal_hold.placed  — show/propagate legal hold in case mgmt;
│   │   │   │                                     #   file.lifecycle.legal_hold.removed — clear legal hold indicators;
│   │   │   │                                     #   file.link.signed_url.created      — append to audit trail;
│   │   │   │                                     #   file.link.signed_url.revoked      — append to audit trail / revoke confirmations;
│   │   │   │                                     #   file.download.logged              — append to audit (SIEM/data-lake feed).
│   │   │   │
│   │   │   ├── search_handler.go              # Consumes (from search-be):
│   │   │   │                                     #   search.taxonomy.updated           — taxonomy quality oversight dashboards;
│   │   │   │                                     #   search.ltr.signal.recorded        — LTR quality monitoring;
│   │   │   │                                     #   search.personalization.updated    — privacy/compliance review surfaces;
│   │   │   │                                     #   search.index.hygiene.deduplicated — hygiene queue visibility;
│   │   │   │                                     #   search.index.hygiene.archived     — archive tracking;
│   │   │   │                                     #   search.index.hygiene.visibility.changed — moderation/visibility change intake;
│   │   │   │                                     #   search.document.indexed           — index ops audit;
│   │   │   │                                     #   search.document.reindexed         — reindex ops audit.
│   │   │   │
│   │   │   ├── reviews_handler.go             # Consumes (from reviews-be):
│   │   │   │                                     #   review.double_blind.submitted     — case intake; SLA timers;
│   │   │   │                                     #   review.double_blind.published     — publish audit / analytics;
│   │   │   │                                     #   review.double_blind.timeout       — timeout audit, follow-ups;
│   │   │   │                                     #   review.response.posted            — right-of-reply audit/visibility;
│   │   │   │                                     #   review.feedback.private_submitted — research/analytics sink;
│   │   │   │                                     #   reputation.updated                — eligibility/quality dashboards;
│   │   │   │                                     #   review.abuse.auto_flagged         — auto-moderation queue intake;
│   │   │   │                                     #   review.moderation.state.changed   — reflect action history & state.
│   │   │   │
│   │   │   ├── subscriptions_handler.go       # Consumes (from subscriptions-be):
│   │   │   │                                     #   subscription.created              — customer lifecycle dashboards;
│   │   │   │                                     #   subscription.changed              — same;
│   │   │   │                                     #   subscription.paused               — alerts & eligibility views;
│   │   │   │                                     #   subscription.canceled             — churn dashboards;
│   │   │   │                                     #   subscription.renewed              — renewal audit;
│   │   │   │                                     #   entitlement.feature.enabled       — propagate plan feature flags to runtime registry;
│   │   │   │                                     #   entitlement.feature.disabled      — revoke flags in registry;
│   │   │   │                                     #   connects.purchased                — credits oversight;
│   │   │   │                                     #   connects.used                     — consumption monitoring;
│   │   │   │                                     #   connects.expired                  — expiry monitoring;
│   │   │   │                                     #   connects.granted                  — promo/grant audit;
│   │   │   │                                     #   billing.invoice.generated         — billing pipeline visibility;
│   │   │   │                                     #   billing.invoice.exported          — export audit to financial-be;
│   │   │   │                                     #   billing.proration.applied         — proration audit;
│   │   │   │                                     #   usage.meter.incremented           — usage alerts dashboards;
│   │   │   │                                     #   usage.limit.reached               — limit alerts dashboards.
│   │   │   │
│   │   │   ├── contracts_handler.go           # Consumes (from contracts-be):
│   │   │   │                                     #   contract.financial_hold.placed    — open/annotate risk case;
│   │   │   │                                     #   contract.financial_hold.released  — close/annotate case;
│   │   │   │                                     #   contract.state.changed            — visibility/policy reactions (read-only intake).
│   │   │   │
│   │   │   ├── financial_handler.go           # Consumes (from financial-be & payments rails):
│   │   │   │                                     #   payment.processed | payment.failed  — dunning/collections dashboards;
│   │   │   │                                     #   payment.chargeback.created        — chargeback case intake;
│   │   │   │                                     #   payment.chargeback.updated        — case status updates;
│   │   │   │                                     #   risk.hold.placed | risk.hold.released — risk dashboard streams;
│   │   │   │                                     #   risk.reserve.set | risk.reserve.updated — reserve monitoring;
│   │   │   │                                     #   risk.chargeback.recorded | updated  — chargeback metrics;
│   │   │   │                                     #   risk.velocity.alert               — anomaly alerts.
│   │   │   │
│   │   │   ├── users_handler.go               # Consumes (from users-be / identity flows):
│   │   │   │                                     #   user.updated                      — reflect flags/verifications in case views;
│   │   │   │                                     #   contract.dispute.opened (user-side) — cross-link disputes to user case.
│   │   │   │
│   │   │   ├── reports_exports_handler.go     # Consumes (cross-service signals to drive admin exports/notifications):
│   │   │   │                                     #   admin.data_export.requested       — (when originated by other surfaces) intake/export flow;
│   │   │   │                                     #   admin.data_export.approved        — approval audit;
│   │   │   │                                     #   admin.data_export.generated       — delivery orchestration audit;
│   │   │   │                                     #   admin.data_export.delivered       — completion audit;
│   │   │   │                                     #   admin.data_export.revoked         — revoke audit.
│   │   │   │
│   │   │   └── communications_handler.go      # Consumes (from communications-be):
│   │   │                                         #   comm.delivery.logged   
│   │   ├── admin_user/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Create, Update, Deactivate
│   │   │   ├── queries.go                    # Get, List admins
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── permission_manager.go         # Manage permissions
│   │   ├── support_ticket/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Create, Assign, Resolve, Close
│   │   │   ├── queries.go                    # Get, List, Search tickets
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   ├── assignment_engine.go          # Auto-assign tickets
│   │   │   ├── escalation_manager.go         # Escalate urgent tickets
│   │   │   └── sla_tracker.go                # Track SLA compliance
│   │   ├── ticket_message/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go                 # 🆕 Input validation for messages/attachments
│   │   ├── support_agent/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── stats_calculator.go           # Calculate agent stats
│   │   ├── canned_response/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── knowledge_base/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Create, Update, Publish
│   │   │   ├── queries.go                    # Search, Get articles
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── search_service.go             # KB search
│   │   ├── faq/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── moderation/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Approve, Reject, Remove
│   │   │   ├── queries.go                    # Get queue, Filter
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   ├── queue_manager.go              # Manage moderation queue
│   │   │   ├── auto_moderator.go             # Automatic moderation rules
│   │   │   └── content_scanner.go            # Scan for violations
│   │   ├── user_management/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Suspend, Ban, Verify, Warn
│   │   │   ├── queries.go                    # Search users, Get details
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── action_validator.go           # Validate admin actions
│   │   ├── content_management/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Remove, Hide, Feature
│   │   │   ├── queries.go                    # List content, Search
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── dispute_resolution/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Review, Decide, Close
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── decision_engine.go            # Help make decisions
│   │   ├── system_config/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Update configs
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── announcement/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Create, Schedule, Send
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── report/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   ├── generator.go                  # Generate reports
│   │   │   ├── scheduler.go                  # Schedule reports
│   │   │   └── exporters/
│   │   │       ├── pdf_exporter.go           # PDF exporter
│   │   │       ├── csv_exporter.go           # CSV exporter
│   │   │       └── excel_exporter.go         # Excel exporter
│   │   ├── analytics/
│   │   │   ├── service.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── aggregator.go                 # Aggregate analytics
│   │   │   ├── metrics_calculator.go         # Calculate KPIs
│   │   │   ├── dashboard_builder.go          # Build dashboards
│   │   │   └── validators.go                 # 🆕 Validate analytics query params
│   │   ├── notification_blast/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Create, Schedule, Send
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── audience_selector.go          # Select target users
│   │   ├── audit/
│   │   │   ├── service.go
│   │   │   ├── queries.go                    # Query audit logs
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go                 # 🆕 Validate audit queries/filters
│   │
│   │   # --------------------------- NEW APPLICATION LAYERS (ADD-ONLY) ----------------
│   │   ├── fraud_review/                     # 🆕
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # OpenReview, Assign, AddNote, SetStatus, LinkEvidence
│   │   │   ├── queries.go                    # Get, List, Search by user/payment/risk-signal
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── user_report/                      # 🆕
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Create, Triage, Reassign, AddNote, Resolve, Dismiss
│   │   │   ├── queries.go                    # Get, List, Filter by category/status/SLA
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── risk/                             # 🆕 dashboards & actions
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # PlaceHold, ReleaseHold, SetReserve, RecordChargeback
│   │   │   ├── queries.go                    # Dashboards: holds/reserves/chargebacks/alerts/anomalies
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── ticket_note/                      # 🆕
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # AddNote, EditNote, DeleteNote
│   │   │   ├── queries.go                    # ListNotesForTicket
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   └── data_export/                      # 🆕
│   │       ├── service.go
│   │       ├── commands.go                   # RequestExport, Approve, Generate, Deliver, Revoke
│   │       ├── queries.go                    # GetExport, ListExports, GetRequest
│   │       ├── dto.go
│   │       ├── mapper.go
│   │       └── validators.go
│   │   # ---------------------------------------------------------------------------
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection setup (DSN from config, connection pooling)
│   │   │       ├── transaction.go            # Transaction helpers (Begin, Commit, Rollback, WithTransaction wrapper)
│   │   │       ├── migrations.go             # 📝 UPDATED: Auto-migrate logic (now with version tracking, GORM AutoMigrate for all tables)
│   │   │       ├── version.go                # 🆕 Schema version tracking (SchemaVersion table, RecordMigration function)
│   │   │       ├── safety.go                 # 🆕 Pre-migration safety checks (environment validation, disk space check)
│   │   │       ├── admin_user_repository.go
│   │   │       ├── support_ticket_repository.go
│   │   │       ├── ticket_message_repository.go
│   │   │       ├── support_agent_repository.go
│   │   │       ├── canned_response_repository.go
│   │   │       ├── knowledge_base_repository.go
│   │   │       ├── faq_repository.go
│   │   │       ├── moderation_queue_repository.go
│   │   │       ├── user_action_repository.go
│   │   │       ├── content_action_repository.go
│   │   │       ├── dispute_resolution_repository.go
│   │   │       ├── system_config_repository.go
│   │   │       ├── announcement_repository.go
│   │   │       ├── report_repository.go
│   │   │       ├── analytics_dashboard_repository.go
│   │   │       ├── audit_log_repository.go
│   │   │       ├── notification_blast_repository.go
│   │   │       ├── platform_stats_repository.go
│   │   │       └── outbox_repository.go      # ❌ REMOVED (use platform-shared/outbox/postgres/repository.go)
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go             # Redis connection setup (connection pooling, retry logic)
│   │   │       ├── admin_cache.go            # Cache admin data (Get, Set, Invalidate with TTL)
│   │   │       ├── ticket_cache.go           # Cache tickets
│   │   │       ├── stats_cache.go            # Cache stats
│   │   │       └── config_cache.go           # Cache configs
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go               # 📝 UPDATED: Kafka consumer (now uses platform-shared/inbox for message deduplication)
│   │   │       ├── producer.go               # 📝 UPDATED: Kafka producer (now uses platform-shared/outbox for reliable publishing)
│   │   │       ├── topics.go                 # 📝 UPDATED: Topic constants (now imported from contracts/events - admin.user.suspended, admin.dispute.resolved)
│   │   │       └── scram.go                  # SCRAM authentication for Kafka (SASL/SCRAM-SHA-256)
│   │   ├── external_services/
│   │   │   ├── users_client.go               # Users service client
│   │   │   ├── jobs_client.go                # Jobs service client
│   │   │   ├── proposals_client.go           # Proposals service client
│   │   │   ├── contracts_client.go           # Contracts service client
│   │   │   ├── financial_client.go           # Financial service client
│   │   │   ├── reviews_client.go             # Reviews service client
│   │   │   ├── communications_client.go      # Communications service client
│   │   │   ├── search_client.go              # Search service client
│   │   │   ├── storage_client.go             # Storage service client
│   │   │   └── subscriptions_client.go       # Subscriptions service client
│   │   ├── keycloak/
│   │   │   ├── admin_client.go               # Keycloak admin operations
│   │   │   └── user_manager.go               # Manage users via Keycloak
│   │   ├── reporting/
│   │   │   ├── pdf_generator.go              # PDF generator
│   │   │   ├── csv_generator.go              # CSV generator
│   │   │   └── excel_generator.go            # Excel generator
│   │   └── outbox/
│   │       ├── processor.go                  # ❌ REMOVED (use platform-shared/outbox/forwarder.go)
│   │       └── scheduler.go                  # ❌ REMOVED (use platform-shared/outbox/scheduler.go)
│   │
│   │   # --------------------------- NEW POSTGRES REPOSITORIES (ADD-ONLY) -------------
│   │   └── persistence/
│   │       └── postgres/
│   │           ├── fraud_review_repository.go              # 🆕
│   │           ├── user_report_repository.go               # 🆕
│   │           ├── risk_hold_repository.go                 # 🆕
│   │           ├── risk_reserve_repository.go              # 🆕
│   │           ├── risk_chargeback_repository.go           # 🆕
│   │           ├── risk_velocity_alert_repository.go       # 🆕
│   │           ├── risk_country_rate_anomaly_repository.go # 🆕
│   │           ├── ticket_note_repository.go               # 🆕
│   │           └── data_export_repository.go               # 🆕
│   │   # ---------------------------------------------------------------------------
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── admin_user_handler.go
│   │       │   ├── support_ticket_handler.go
│   │       │   ├── ticket_message_handler.go
│   │       │   ├── support_agent_handler.go
│   │       │   ├── canned_response_handler.go
│   │       │   ├── knowledge_base_handler.go
│   │       │   ├── faq_handler.go
│   │       │   ├── moderation_handler.go
│   │       │   ├── user_management_handler.go
│   │       │   ├── content_management_handler.go
│   │       │   ├── dispute_resolution_handler.go
│   │       │   ├── system_config_handler.go
│   │       │   ├── announcement_handler.go
│   │       │   ├── report_handler.go
│   │       │   ├── analytics_handler.go
│   │       │   ├── notification_blast_handler.go
│   │       │   ├── audit_log_handler.go
│   │       │   ├── dashboard_handler.go
│   │       │   └── health_handler.go          # Health check endpoints (/health, /ready, /live)
│   │       ├── middleware/
│   │       │   ├── auth.go                    # 📝 UPDATED: Authentication middleware (uses pkg/auth for JWT verification)
│   │       │   ├── admin_auth.go              # 📝 UPDATED: Admin-specific auth (uses pkg/auth)
│   │       │   ├── permission_check.go        # 📝 UPDATED: Check admin permissions (uses pkg/auth)
│   │       │   ├── audit_logger.go            # Auto-log all admin actions (platform-shared/logging)
│   │       │   ├── cors.go                    # 📝 UPDATED: CORS (platform-shared/ginx)
│   │       │   ├── rate_limit.go              # Rate limiting
│   │       │   ├── logging.go                 # 📝 UPDATED: Structured logs (platform-shared/ginx/logging)
│   │       │   ├── error_handler.go           # Error handling
│   │       │   └── request_id.go              # 📝 UPDATED: X-Request-ID (platform-shared/ginx/requestid)
│   │       ├── responses/
│   │       │   ├── success.go                 # 📝 UPDATED: use platform-shared/httpx/response.go
│   │       │   ├── error.go                   # 📝 UPDATED: use platform-shared/httpx/errors.go
│   │       │   └── pagination.go              # 📝 UPDATED: use platform-shared/httpx/pagination.go
│   │       ├── routes/                        # 🆕 Route registrars mapped 1:1 to handlers
│   │       │   ├── admin_user_routes.go           # /admin/users/*
│   │       │   ├── support_ticket_routes.go       # /admin/tickets/*
│   │       │   ├── ticket_message_routes.go       # /admin/tickets/:id/messages/*
│   │       │   ├── support_agent_routes.go        # /admin/agents/*
│   │       │   ├── canned_response_routes.go      # /admin/canned-responses/*
│   │       │   ├── knowledge_base_routes.go       # /admin/kb/*
│   │       │   ├── faq_routes.go                  # /admin/faqs/*
│   │       │   ├── moderation_routes.go           # /admin/moderation/*
│   │       │   ├── user_management_routes.go      # /admin/users/actions/*
│   │       │   ├── content_management_routes.go   # /admin/content/*
│   │       │   ├── dispute_resolution_routes.go   # /admin/disputes/*
│   │       │   ├── system_config_routes.go        # /admin/config/*
│   │       │   ├── announcement_routes.go         # /admin/announcements/*
│   │       │   ├── report_routes.go               # /admin/reports/*
│   │       │   ├── analytics_routes.go            # /admin/analytics/*
│   │       │   ├── notification_blast_routes.go   # /admin/blasts/*
│   │       │   ├── audit_log_routes.go            # /admin/audit/*
│   │       │   ├── dashboard_routes.go            # /admin/dashboard/*
│   │       │   # --------------------------- NEW ROUTES (ADD-ONLY) -------------------
│   │       │   ├── fraud_review_routes.go         # 🆕 /admin/fraud-reviews/*
│   │       │   ├── user_report_routes.go          # 🆕 /admin/user-reports/*
│   │       │   ├── risk_routes.go                 # 🆕 /admin/risk/*
│   │       │   ├── data_export_routes.go          # 🆕 /admin/exports/*
│   │       │   └── ticket_note_routes.go          # 🆕 /admin/tickets/:id/notes/*
│   │       └── router.go                          # 📝 UPDATED: HTTP router setup (uses Gin, applies platform-shared/ginx middleware)
│   │
│   └── config/                                   # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                             # 🆕 Typed Config struct (App, Server, Postgres, Kafka, Redis, Keycloak)
│       ├── loader.go                             # 🆕 Config loader using Viper (precedence: flags → env → file → defaults)
│       └── docs/
│           └── CONFIGURATION.md                  # 🆕 Configuration documentation (all ENV vars, defaults, examples)
│
├── config/                                       # 🆕 MEDIUM - Configuration files
│   ├── default.yaml                              # 🆕 Default configuration
│   ├── dev.yaml                                  # 🆕 Development overrides
│   └── prod.yaml                                 # 🆕 Production overrides
│
├── dapr/                                         # 🆕 MEDIUM - Dapr components split by environment
│   ├── local/                                    # 🆕 For dapr run
│   │   ├── pubsub.yaml                           # Kafka pub/sub component
│   │   └── statestore.yaml                       # State store component
│   └── k8s/                                      # 🆕 For kubectl apply
│       ├── pubsub.yaml                           # Kafka with scopes: ["admin-be"]
│       ├── statestore.yaml                       # State store with scopes
│       └── secrets.yaml                          # Dapr secret store
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                             # Service-specific errors
│   │   └── codes.go                              # Error codes
│   ├── logger/
│   │   └── logger.go                             # ❌ REMOVED: Use platform-shared/logging
│   ├── utils/
│   │   ├── validator.go                          # Local validation utilities
│   │   ├── sanitizer.go                          # Sanitize input
│   │   ├── permission_checker.go                 # Permission checking utilities
│   │   └── report_formatter.go                   # Report formatting
│   └── constants/
│       ├── events.go                             # ❌ REMOVED: Use contracts/events
│       ├── topics.go                             # ❌ REMOVED: Use contracts/events
│       ├── permissions.go                        # Permissions constants
│       └── moderation_actions.go                 # Moderation actions constants
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                       # Kubernetes Deployment
│       ├── service.yaml                          # Kubernetes Service
│       ├── configmap.yaml                        # ConfigMap
│       ├── secrets.yaml                          # Secrets
│       ├── hpa.yaml                              # HPA
│       ├── pdb.yaml                              # PDB
│       ├── rbac.yaml                             # Admin RBAC policies
│       └── servicemonitor.yaml                   # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                            # Setup local environment
│   ├── get-secrets.sh                            # Fetch secrets
│   ├── seed-admin-users.sh                       # Seed admin users
│   ├── seed-canned-responses.sh                  # Seed canned responses
│   └── seed-data.sh                              # Seed test data
│
├── tests/
│   ├── unit/
│   │   ├── domain/
│   │   │   ├── admin_user_test.go               # Admin user tests
│   │   │   ├── support_ticket_test.go           # Support ticket tests
│   │   │   └── moderation_queue_test.go         # Moderation queue tests
│   │   ├── application/
│   │   │   ├── admin_user_service_test.go       # Admin user service tests
│   │   │   ├── support_ticket_service_test.go   # Support ticket service tests
│   │   │   └── moderation_service_test.go       # Moderation service tests
│   │   └── infrastructure/
│   │       ├── postgres_repository_test.go      # Postgres repository tests
│   │       └── kafka_producer_test.go           # Kafka producer tests
│   ├── integration/
│   │   ├── handlers/
│   │   │   ├── admin_user_handler_test.go      # Admin user handler tests
│   │   │   ├── support_ticket_handler_test.go  # Support ticket handler tests
│   │   │   └── moderation_handler_test.go      # Moderation handler tests
│   │   └── repositories/
│   │       ├── admin_user_repository_test.go   # Admin user repository tests
│   │       └── support_ticket_repository_test.go # Support ticket repository tests
│   └── e2e/
│       └── scenarios/
│           ├── ticket_workflow_test.go         # Ticket workflow tests
│           ├── moderation_workflow_test.go     # Moderation workflow tests
│           └── dispute_resolution_test.go      # Dispute resolution tests
│
├── docs/
│   ├── README.md                               # Service overview
│   ├── api.md                                  # API documentation
│   ├── events.md                               # 📝 UPDATED: Events (published: admin.user.suspended, admin.dispute.resolved, consumed: user.flagged, contract.dispute.opened, etc.)
│   ├── admin-roles.md                          # Admin roles
│   ├── permissions.md                          # Permissions
│   ├── moderation-guide.md                     # Moderation guide
│   ├── support-workflows.md                    # Support workflows
│   ├── reporting.md                            # Reporting
│   ├── MIGRATIONS.md                           # 🆕 Migration history
│   ├── SCHEMA.md                               # 🆕 Database schema
│   └── RUNBOOK.md                              # 🆕 Operational procedures
│
├── .github/
│   └── workflows/
│       ├── ci.yml                              # CI workflow
│       └── cd.yml                              # CD workflow
│
├── go.mod                                      # 📝 UPDATED: Imports pkg/auth, platform-shared, contracts/events
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md

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