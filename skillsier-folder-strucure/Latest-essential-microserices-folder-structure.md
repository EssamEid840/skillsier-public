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
│   └── api/
│       └── main.go                           # 📝 UPDATED: Application entry point - initializes Gin, Dapr, connects to Postgres (now uses loadConfig from internal/config, platform-shared/logging)
│
├── internal/
│   ├── domain/                               # 🏛️ Domain Layer - Business logic & entities
│   │   ├── user/
│   │   │   ├── entity.go                     # User aggregate root (ID, Username, Email, FirstName, LastName, UserType, Status)
│   │   │   ├── value_objects.go             # Email, Phone, Address value objects (validation, formatting)
│   │   │   ├── enums.go                     # UserType (Freelancer, Client), AccountStatus (Active, Suspended, Banned), VerificationStatus
│   │   │   ├── errors.go                    # Domain-specific errors (UserNotFound, EmailTaken, InvalidEmail)
│   │   │   └── repository.go                # User repository interface (Create, Update, FindByID, FindByEmail, List)
│   │   │
│   │   ├── profile/
│   │   │   ├── entity.go                     # Extended profile info (Bio, Location, ProfilePictureURL, CompletionPercentage)
│   │   │   ├── preferences.go               # User preferences and settings (Language, Timezone, Currency)
│   │   │   ├── availability.go              # Work availability, timezone, working hours
│   │   │   ├── errors.go                    # Profile-specific errors (ProfileNotFound, InvalidBio)
│   │   │   └── repository.go                # Profile repository interface
│   │   │
│   │   ├── skill/
│   │   │   ├── entity.go                     # User skills (UserID, SkillID, Proficiency, YearsOfExperience)
│   │   │   ├── proficiency.go               # Skill level enum (Beginner, Intermediate, Advanced, Expert)
│   │   │   ├── errors.go                    # Skill-specific errors (SkillNotFound, DuplicateSkill)
│   │   │   └── repository.go                # Skill repository interface
│   │   │
│   │   ├── experience/
│   │   │   ├── entity.go                     # Work experience (Company, Title, Description, StartDate, EndDate, IsCurrent)
│   │   │   ├── errors.go                    # Experience-specific errors (InvalidDateRange)
│   │   │   └── repository.go                # Experience repository interface
│   │   │
│   │   ├── education/
│   │   │   ├── entity.go                     # Educational background (School, Degree, Field, GraduationYear, Description)
│   │   │   ├── errors.go                    # Education-specific errors (InvalidYear)
│   │   │   └── repository.go                # Education repository interface
│   │   │
│   │   ├── certification/
│   │   │   ├── entity.go                     # Certifications (Name, IssuingOrganization, IssueDate, ExpiryDate, CredentialID, URL)
│   │   │   ├── verification.go              # Verification status (Pending, Verified, Rejected, Expired)
│   │   │   ├── errors.go                    # Certification-specific errors (CertificationNotFound, AlreadyVerified)
│   │   │   └── repository.go                # Certification repository interface
│   │   │
│   │   ├── portfolio/
│   │   │   ├── entity.go                     # Portfolio items (UserID, Title, Description, URL, ThumbnailURL, DisplayOrder)
│   │   │   ├── item.go                      # Portfolio item details (Images, Videos, Links)
│   │   │   ├── media.go                     # Associated media (Images, Videos, Documents)
│   │   │   ├── errors.go                    # Portfolio-specific errors (PortfolioNotFound, MaxItemsExceeded)
│   │   │   └── repository.go                # Portfolio repository interface
│   │   │
│   │   ├── language/
│   │   │   ├── entity.go                     # Language proficiency (UserID, Language, ProficiencyLevel)
│   │   │   ├── errors.go                    # Language-specific errors (LanguageNotFound, DuplicateLanguage)
│   │   │   └── repository.go                # Language repository interface
│   │   │
│   │   ├── freelancer/
│   │   │   ├── entity.go                     # Freelancer-specific data (UserID, Title, Overview, VideoIntroURL)
│   │   │   ├── profile.go                   # Freelancer profile details (Tagline, ServiceOffering)
│   │   │   ├── rates.go                     # Hourly/fixed rates (HourlyRate, MinimumBudget, Currency)
│   │   │   ├── stats.go                     # Job stats, earnings (TotalJobs, TotalEarnings, SuccessRate, ResponseTime)
│   │   │   ├── errors.go                    # Freelancer-specific errors (FreelancerNotFound, InvalidRate)
│   │   │   └── repository.go                # Freelancer repository interface
│   │   │
│   │   ├── client/
│   │   │   ├── entity.go                     # Client-specific data (UserID, CompanyName, CompanySize, Industry)
│   │   │   ├── profile.go                   # Client profile details (About, Website, Location)
│   │   │   ├── company.go                   # Company info (Name, Size, Industry, Founded, Employees)
│   │   │   ├── stats.go                     # Hiring stats, spending (TotalHires, TotalSpent, ActiveContracts, AverageRating)
│   │   │   ├── errors.go                    # Client-specific errors (ClientNotFound, InvalidCompanySize)
│   │   │   └── repository.go                # Client repository interface
│   │   │
│   │   ├── verification/
│   │   │   ├── entity.go                     # Identity verification (UserID, Type, Status, SubmittedAt, VerifiedAt, RejectionReason)
│   │   │   ├── document.go                  # Verification documents (ID, Passport, ProofOfAddress, Selfie)
│   │   │   ├── errors.go                    # Verification-specific errors (VerificationNotFound, DocumentMissing)
│   │   │   └── repository.go                # Verification repository interface
│   │   │
│   │   ├── settings/
│   │   │   ├── entity.go                     # User settings (UserID, Settings JSON, Theme, Language, Timezone)
│   │   │   ├── notification_prefs.go        # Notification preferences (Email, SMS, Push, InApp for different event types)
│   │   │   ├── privacy_settings.go          # Privacy settings (ProfileVisibility, ShowEmail, ShowPhone, SearchableProfile)
│   │   │   ├── errors.go                    # Settings-specific errors (SettingsNotFound, InvalidPreference)
│   │   │   └── repository.go                # Settings repository interface
│   │   │
│   │   ├── saved_items/
│   │   │   ├── entity.go                     # Saved jobs, freelancers (UserID, ItemType, ItemID, SavedAt, Notes)
│   │   │   ├── errors.go                    # Saved items-specific errors (ItemNotFound, DuplicateSave)
│   │   │   └── repository.go                # Saved items repository interface
│   │   │
│   │   ├── blocked_users/
│   │   │   ├── entity.go                     # Blocked user relationships (BlockerID, BlockedID, BlockedAt, Reason)
│   │   │   ├── errors.go                    # Blocking-specific errors (AlreadyBlocked, CannotBlockSelf)
│   │   │   └── repository.go                # Blocked users repository interface
│   │   │
│   │   ├── user_suspension/
│   │   │   ├── entity.go                     # Track user suspensions (UserID, Reason, StartDate, EndDate, SuspendedBy, IsActive)
│   │   │   ├── reason.go                    # Suspension reasons enum (TOSViolation, PaymentIssue, QualityIssues, AbusiveeBehavior)
│   │   │   ├── duration.go                  # Suspension duration (Days, Weeks, Months, Permanent)
│   │   │   ├── errors.go                    # Suspension-specific errors (AlreadySuspended, InvalidDuration)
│   │   │   └── repository.go                # Suspension repository interface
│   │   │
│   │   ├── user_ban/
│   │   │   ├── entity.go                     # Track user bans (UserID, Reason, BannedAt, BannedBy, IsPermanent, ExpiresAt)
│   │   │   ├── reason.go                    # Ban reasons enum (Fraud, SevereAbuse, MultipleViolations, SecurityThreat)
│   │   │   ├── permanent.go                 # Permanent vs temporary flag
│   │   │   ├── errors.go                    # Ban-specific errors (AlreadyBanned, CannotUnban)
│   │   │   └── repository.go                # Ban repository interface
│   │   │
│   │   └── user_warning/
│   │       ├── entity.go                     # Track user warnings (UserID, Reason, IssuedAt, IssuedBy, AcknowledgedAt)
│   │       ├── reason.go                    # Warning reasons enum (LateDelivery, PoorQuality, UnresponsiveCommunication, MinorViolation)
│   │       ├── severity.go                  # Warning severity level (Low, Medium, High, Critical)
│   │       ├── errors.go                    # Warning-specific errors (WarningNotFound, TooManyWarnings)
│   │       └── repository.go                # Warning repository interface
│   │
│   ├── application/                          # 📋 Application Layer - Use cases & orchestration
│   │   ├── user/
│   │   │   ├── service.go                    # User business logic (Create, Update, Delete, Verify, Search)
│   │   │   ├── commands.go                  # Command handlers (CreateUser, UpdateUser, DeleteUser, VerifyEmail)
│   │   │   ├── queries.go                   # Query handlers (GetUser, ListUsers, SearchUsers, GetUserStats)
│   │   │   ├── dto.go                       # Data transfer objects (UserDTO, CreateUserDTO, UpdateUserDTO, UserListDTO)
│   │   │   ├── mapper.go                    # Entity-DTO mapping (ToDTO, FromDTO, ToEntity, BulkToDTO)
│   │   │   └── validators.go                # Input validation (ValidateEmail, ValidateUsername, ValidatePassword, ValidatePhone)
│   │   │
│   │   ├── profile/
│   │   │   ├── service.go                    # Profile business logic (Complete, Update, CalculateCompleteness)
│   │   │   ├── dto.go                       # Profile DTOs (ProfileDTO, UpdateProfileDTO)
│   │   │   ├── mapper.go                    # Profile mappers
│   │   │   └── validators.go                # Profile validators (ValidateBio, ValidateLocation)
│   │   │
│   │   ├── skill/
│   │   │   ├── service.go                    # Skill business logic (Add, Remove, Update, Reorder)
│   │   │   ├── dto.go                       # Skill DTOs (SkillDTO, AddSkillDTO)
│   │   │   └── mapper.go                    # Skill mappers
│   │   │
│   │   ├── experience/
│   │   │   ├── service.go                    # Experience business logic (Add, Update, Delete, Verify)
│   │   │   ├── dto.go                       # Experience DTOs (ExperienceDTO, CreateExperienceDTO)
│   │   │   └── mapper.go                    # Experience mappers
│   │   │
│   │   ├── education/
│   │   │   ├── service.go                    # Education business logic (Add, Update, Delete)
│   │   │   ├── dto.go                       # Education DTOs (EducationDTO, CreateEducationDTO)
│   │   │   └── mapper.go                    # Education mappers
│   │   │
│   │   ├── certification/
│   │   │   ├── service.go                    # Certification business logic (Add, Update, Delete, RequestVerification)
│   │   │   ├── dto.go                       # Certification DTOs (CertificationDTO, CreateCertificationDTO)
│   │   │   └── mapper.go                    # Certification mappers
│   │   │
│   │   ├── portfolio/
│   │   │   ├── service.go                    # Portfolio business logic (Add, Update, Delete, Reorder)
│   │   │   ├── dto.go                       # Portfolio DTOs (PortfolioItemDTO, CreatePortfolioItemDTO)
│   │   │   └── mapper.go                    # Portfolio mappers
│   │   │
│   │   ├── language/
│   │   │   ├── service.go                    # Language business logic (Add, Update, Delete)
│   │   │   ├── dto.go                       # Language DTOs (LanguageDTO, AddLanguageDTO)
│   │   │   └── mapper.go                    # Language mappers
│   │   │
│   │   ├── freelancer/
│   │   │   ├── service.go                    # Freelancer business logic (CompleteProfile, UpdateRates, UpdateStats)
│   │   │   ├── commands.go                  # Freelancer commands (UpdateRates, UpdateAvailability)
│   │   │   ├── queries.go                   # Freelancer queries (GetStats, GetEarnings, GetSuccessRate)
│   │   │   ├── dto.go                       # Freelancer DTOs (FreelancerDTO, UpdateRatesDTO, FreelancerStatsDTO)
│   │   │   ├── mapper.go                    # Freelancer mappers
│   │   │   └── stats_calculator.go          # Calculate freelancer statistics (TotalEarnings, SuccessRate, ResponseTime)
│   │   │
│   │   ├── client/
│   │   │   ├── service.go                    # Client business logic (CompleteProfile, UpdateCompany, UpdateStats)
│   │   │   ├── commands.go                  # Client commands (UpdateCompany, UpdateIndustry)
│   │   │   ├── queries.go                   # Client queries (GetHiringStats, GetSpendingHistory)
│   │   │   ├── dto.go                       # Client DTOs (ClientDTO, UpdateCompanyDTO, ClientStatsDTO)
│   │   │   ├── mapper.go                    # Client mappers
│   │   │   └── stats_calculator.go          # Calculate client statistics (TotalSpent, HireRate, AverageProjectValue)
│   │   │
│   │   ├── verification/
│   │   │   ├── service.go                    # Verification business logic (Submit, Approve, Reject, RequestAdditionalDocs)
│   │   │   ├── dto.go                       # Verification DTOs (VerificationDTO, SubmitVerificationDTO)
│   │   │   └── mapper.go                    # Verification mappers
│   │   │
│   │   ├── settings/
│   │   │   ├── service.go                    # Settings business logic (Update, GetPreferences, UpdatePrivacy)
│   │   │   ├── dto.go                       # Settings DTOs (SettingsDTO, UpdateSettingsDTO, NotificationPrefsDTO)
│   │   │   └── mapper.go                    # Settings mappers
│   │   │
│   │   ├── saved_items/
│   │   │   ├── service.go                    # Saved items business logic (Save, Unsave, List, Search)
│   │   │   ├── dto.go                       # Saved items DTOs (SavedItemDTO, SaveItemDTO)
│   │   │   └── mapper.go                    # Saved items mappers
│   │   │
│   │   ├── suspension/
│   │   │   ├── service.go                    # Suspension business logic (Suspend, Unsuspend, Extend, GetHistory)
│   │   │   ├── commands.go                  # Suspend, Unsuspend, ExtendSuspension commands
│   │   │   ├── dto.go                       # Suspension DTOs (SuspensionDTO, SuspendUserDTO)
│   │   │   └── mapper.go                    # Suspension mappers
│   │   │
│   │   ├── ban/
│   │   │   ├── service.go                    # Ban business logic (Ban, Unban, GetBanHistory)
│   │   │   ├── commands.go                  # Ban, Unban commands
│   │   │   ├── dto.go                       # Ban DTOs (BanDTO, BanUserDTO)
│   │   │   └── mapper.go                    # Ban mappers
│   │   │
│   │   ├── warning/
│   │   │   ├── service.go                    # Warning business logic (IssueWarning, AcknowledgeWarning, GetWarnings)
│   │   │   ├── commands.go                  # IssueWarning, AcknowledgeWarning commands
│   │   │   ├── dto.go                       # Warning DTOs (WarningDTO, IssueWarningDTO)
│   │   │   └── mapper.go                    # Warning mappers
│   │   │
│   │   └── eventhandler/                    # 📝 UPDATED: Event handlers (now uses contracts/events for event types)
│   │       ├── keycloak_handler.go          # Handle Keycloak user registration events (sync user data)
│   │       ├── contract_handler.go          # Handle contract completion events (update freelancer stats, earnings)
│   │       ├── review_handler.go            # Handle review submission events (update user ratings, reputation)
│   │       ├── transaction_handler.go       # Handle payment events (update earnings, balance)
│   │       ├── badge_handler.go             # Handle badge award events (update profile badges, achievements)
│   │       └── admin_handler.go             # Handle admin actions on users (suspension, ban, verification)
│   │
│   ├── infrastructure/                       # 🔧 Infrastructure Layer - External concerns
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection setup (DSN from config, connection pooling)
│   │   │       ├── transaction.go           # Transaction helpers (Begin, Commit, Rollback, WithTransaction wrapper)
│   │   │       ├── migrations.go            # 📝 UPDATED: Auto-migration logic (now with version tracking, GORM AutoMigrate for all tables)
│   │   │       ├── version.go               # 🆕 Schema version tracking (SchemaVersion table, RecordMigration function)
│   │   │       ├── safety.go                # 🆕 Pre-migration safety checks (environment validation, disk space check, backup verification)
│   │   │       ├── user_repository.go       # User repository implementation (CRUD operations using GORM)
│   │   │       ├── profile_repository.go    # Profile repository implementation
│   │   │       ├── skill_repository.go      # Skill repository implementation
│   │   │       ├── experience_repository.go # Experience repository implementation
│   │   │       ├── education_repository.go  # Education repository implementation
│   │   │       ├── certification_repository.go # Certification repository implementation
│   │   │       ├── portfolio_repository.go  # Portfolio repository implementation
│   │   │       ├── language_repository.go   # Language repository implementation
│   │   │       ├── freelancer_repository.go # Freelancer repository implementation
│   │   │       ├── client_repository.go     # Client repository implementation
│   │   │       ├── verification_repository.go # Verification repository implementation
│   │   │       ├── settings_repository.go   # Settings repository implementation
│   │   │       ├── saved_items_repository.go # Saved items repository implementation
│   │   │       ├── blocked_users_repository.go # Blocked users repository implementation
│   │   │       ├── user_suspension_repository.go # Suspension repository implementation
│   │   │       ├── user_ban_repository.go   # Ban repository implementation
│   │   │       └── user_warning_repository.go # Warning repository implementation
│   │   │
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go            # Redis connection setup (connection pooling, retry logic)
│   │   │       ├── user_cache.go            # User caching logic (Get, Set, Invalidate with TTL, cache-aside pattern)
│   │   │       └── profile_cache.go         # Profile caching logic (similar to user cache)
│   │   │
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go              # 📝 UPDATED: Kafka consumer (now uses platform-shared/inbox for message deduplication)
│   │   │       ├── producer.go              # 📝 UPDATED: Kafka producer (now uses platform-shared/outbox for reliable publishing)
│   │   │       ├── topics.go                # 📝 UPDATED: Topic constants (now imported from contracts/events - user.created, user.updated, etc.)
│   │   │       └── scram.go                 # SCRAM authentication for Kafka (SASL/SCRAM-SHA-256)
│   │   │
│   │   ├── storage/
│   │   │   ├── client.go                    # Storage service client (upload profile pics, portfolios, documents via HTTP API)
│   │   │   └── local.go                     # Local file storage fallback (for development/testing)
│   │   │
│   │   └── keycloak/
│   │       ├── client.go                    # 📝 UPDATED: Keycloak API client (now uses pkg/auth for token verification)
│   │       └── sync.go                      # 📝 UPDATED: Sync users with Keycloak (create, update, delete users via Keycloak Admin API, uses pkg/auth)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── user_handler.go          # User HTTP handlers (GET, POST, PUT, DELETE /users, /users/:id, /users/search)
│   │       │   ├── profile_handler.go       # Profile HTTP handlers (GET, PUT /users/:id/profile, /users/:id/profile/completion)
│   │       │   ├── skill_handler.go         # Skill HTTP handlers (GET, POST, DELETE /users/:id/skills)
│   │       │   ├── experience_handler.go    # Experience HTTP handlers (GET, POST, PUT, DELETE /users/:id/experience)
│   │       │   ├── education_handler.go     # Education HTTP handlers (GET, POST, PUT, DELETE /users/:id/education)
│   │       │   ├── certification_handler.go # Certification HTTP handlers (GET, POST, PUT, DELETE /users/:id/certifications)
│   │       │   ├── portfolio_handler.go     # Portfolio HTTP handlers (GET, POST, PUT, DELETE /users/:id/portfolio)
│   │       │   ├── language_handler.go      # Language HTTP handlers (GET, POST, DELETE /users/:id/languages)
│   │       │   ├── freelancer_handler.go    # Freelancer HTTP handlers (GET, PUT /users/:id/freelancer, /users/:id/freelancer/stats)
│   │       │   ├── client_handler.go        # Client HTTP handlers (GET, PUT /users/:id/client, /users/:id/client/stats)
│   │       │   ├── verification_handler.go  # Verification HTTP handlers (POST, GET /users/:id/verification)
│   │       │   ├── settings_handler.go      # Settings HTTP handlers (GET, PUT /users/:id/settings)
│   │       │   ├── saved_items_handler.go   # Saved items HTTP handlers (GET, POST, DELETE /users/:id/saved)
│   │       │   ├── suspension_handler.go    # Suspension HTTP handlers (POST, DELETE /admin/users/:id/suspend) - admin only
│   │       │   ├── ban_handler.go           # Ban HTTP handlers (POST, DELETE /admin/users/:id/ban) - admin only
│   │       │   ├── warning_handler.go       # Warning HTTP handlers (POST /admin/users/:id/warn) - admin only
│   │       │   └── health_handler.go        # Health check endpoints (/health, /ready, /live)
│   │       │
│   │       ├── middleware/                  # 📝 UPDATED: Middleware (now uses platform-shared middleware)
│   │       │   ├── auth.go                  # 📝 UPDATED: Authentication middleware (uses pkg/auth for JWT verification)
│   │       │   ├── rbac.go                  # 📝 UPDATED: RBAC middleware (uses pkg/auth authorizer for role/permission checks)
│   │       │   ├── cors.go                  # 📝 UPDATED: CORS middleware (uses platform-shared/ginx CORS)
│   │       │   ├── rate_limit.go            # Rate limiting middleware (token bucket, per-user limits)
│   │       │   └── idempotency.go           # 🆕 Idempotency middleware (uses platform-shared/idempotency for duplicate request prevention)
│   │       │
│   │       ├── responses/                   # 📝 UPDATED: Response helpers
│   │       │   └── README.md                # 📝 Points to platform-shared/httpx (use shared response wrappers: Success, Error, Paginated)
│   │       │
│   │       └── router.go                    # 📝 UPDATED: HTTP router setup (uses Gin, applies platform-shared/ginx middleware: requestid, logging, recover, otel, cors, timeout)
│   │
│   └── config/                              # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                        # 🆕 Typed Config struct (App name/version, Server port/host, Postgres DSN/pool, Kafka brokers/topics, Redis addr/password, Auth issuer/audience, Storage endpoint, Keycloak realm/client)
│       ├── loader.go                        # 🆕 Config loader using Viper (precedence: CLI flags → environment variables → config file → defaults)
│       └── docs/
│           └── CONFIGURATION.md             # 🆕 Configuration documentation (all ENV vars, default values, precedence rules, examples)
│
├── config/                                  # 🆕 MEDIUM - Configuration files
│   ├── default.yaml                         # 🆕 Default configuration (dev-friendly defaults: debug logging, local URLs, small pool sizes)
│   ├── dev.yaml                             # 🆕 Development overrides (debug mode enabled, local service URLs, no SSL)
│   └── prod.yaml                            # 🆕 Production overrides (info logging, production URLs, SSL enabled, larger pools, strict validation)
│
├── dapr/                                    # 🆕 MEDIUM - Dapr components split by environment
│   ├── local/                               # 🆕 For `dapr run` (local development with Dapr CLI)
│   │   ├── pubsub.yaml                      # Kafka pub/sub component (localhost:9092, no auth, no scopes needed for local)
│   │   └── statestore.yaml                  # State store component (Redis or Postgres for local dev)
│   └── k8s/                                 # 🆕 For `kubectl apply` (production Kubernetes deployment)
│       ├── pubsub.yaml                      # Kafka pub/sub with proper auth (SCRAM-SHA-256) and scopes: ["users-be"]
│       ├── statestore.yaml                  # Production state store with scopes: ["users-be"]
│       └── secrets.yaml                     # Dapr secret store component (Kubernetes secrets or external vault)
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                        # Service-specific errors (domain errors, application errors)
│   │   └── codes.go                         # Error codes (USER_NOT_FOUND, EMAIL_TAKEN, INVALID_EMAIL, UNAUTHORIZED)
│   ├── logger/                              # ❌ REMOVED: Use platform-shared/logging instead
│   │   └── README.md                        # 📝 Points to platform-shared/logging (import from skillsier.dev/platform-shared/logging)
│   ├── utils/
│   │   ├── validator.go                     # Local validation utilities (if any service-specific validation beyond platform-shared)
│   │   ├── slug.go                          # Generate URL-friendly slugs (from username, company name)
│   │   └── sanitizer.go                     # Sanitize user input (HTML, SQL injection prevention)
│   └── constants/
│       ├── events.go                        # ❌ REMOVED: Use contracts/events instead (import from skillsier.dev/contracts/events/user/v1)
│       └── topics.go                        # ❌ REMOVED: Use contracts/events instead
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                  # Kubernetes Deployment manifest (replicas, resources, health checks)
│       ├── service.yaml                     # Kubernetes Service manifest (ClusterIP, ports)
│       ├── configmap.yaml                   # Configuration as ConfigMap (non-sensitive config from config/prod.yaml)
│       ├── secrets.yaml                     # Secrets (DB password, Kafka credentials, Keycloak client secret)
│       ├── hpa.yaml                         # Horizontal Pod Autoscaler (target CPU 70%, min 2, max 10 replicas)
│       ├── pdb.yaml                         # Pod Disruption Budget (maxUnavailable: 1 for safe rolling updates)
│       └── servicemonitor.yaml              # Prometheus ServiceMonitor (scrape /metrics endpoint)
│
├── scripts/
│   ├── setup-local.sh                       # Setup local development environment (Docker Compose for Postgres, Redis, Kafka)
│   ├── get-secrets.sh                       # Fetch secrets from Kubernetes (kubectl get secret)
│   └── seed-data.sh                         # Seed initial data (admin user, test users, sample profiles)
│
├── tests/
│   ├── unit/
│   │   ├── domain/
│   │   │   ├── user_test.go                 # User entity tests (validation, business rules)
│   │   │   ├── profile_test.go              # Profile entity tests
│   │   │   └── freelancer_test.go           # Freelancer entity tests
│   │   ├── application/
│   │   │   ├── user_service_test.go         # User service tests (mocked repositories)
│   │   │   ├── profile_service_test.go      # Profile service tests
│   │   │   └── freelancer_service_test.go   # Freelancer service tests
│   │   └── infrastructure/
│   │       ├── user_repository_test.go      # User repository tests (mocked DB)
│   │       └── kafka_producer_test.go       # Kafka producer tests
│   ├── integration/
│   │   ├── handlers/
│   │   │   ├── user_handler_test.go         # User handler integration tests (with test DB)
│   │   │   └── profile_handler_test.go      # Profile handler integration tests
│   │   └── repositories/
│   │       └── user_repository_test.go      # User repository integration tests (with real PostgreSQL testcontainer)
│   └── e2e/
│       └── scenarios/
│           ├── user_registration_test.go    # End-to-end user registration flow (Keycloak + DB + Kafka)
│           └── profile_completion_test.go   # End-to-end profile completion flow
│
├── docs/
│   ├── README.md                            # Service overview (purpose, responsibilities, dependencies)
│   ├── API.md                               # API documentation (endpoints, request/response examples, authentication)
│   ├── EVENTS.md                            # 🆕 Service-specific events (published: user.created, user.updated, consumed: contract.completed, review.submitted)
│   ├── ARCHITECTURE.md                      # Service architecture details (layers, patterns, data flow)
│   ├── MIGRATIONS.md                        # 🆕 Migration history (versions applied, changes made, rollback procedures, applied dates)
│   ├── SCHEMA.md                            # 🆕 Current database schema documentation (tables, columns, indexes, relationships)
│   └── RUNBOOK.md                           # 🆕 Operational procedures (how to restart, debug logs, rollback deployment, common issues)
│
├── .github/
│   └── workflows/
│       ├── ci.yml                           # CI workflow (go test, golangci-lint, docker build)
│       └── cd.yml                           # CD workflow (deploy to staging on merge to dev, deploy to prod on tag)
│
├── go.mod                                   # 📝 UPDATED: Dependencies (now imports pkg/auth, platform-shared, contracts/events)
├── go.sum                                   # Dependency checksums
├── .env.example                             # Example environment variables (DB_DSN, KAFKA_BROKERS, REDIS_ADDR, KEYCLOAK_REALM)
├── Makefile                                 # Build targets (run, test, lint, docker-build, k8s-deploy, migrate, seed)
├── Dockerfile                               # Multi-stage Docker build (builder + runtime, minimal alpine image)
├── .dockerignore                            # Docker ignore patterns (tests, docs, .git)
├── .gitignore                               # Git ignore patterns (vendor, bin, .env)
└── README.md                                # Service README (quick start, development setup, deployment)
```

---

---

## 📦 **2️⃣ jobs-be (Job Posting and Management Service)**

```
apps/be/jobs-be/
│
├── cmd/
│   └── api/
│       └── main.go                           # Entry point - initializes Gin, Dapr, connects to Postgres (uses loadConfig, platform-shared)
│
├── internal/
│   ├── domain/                               # 🏛️ Domain Layer - Business logic & entities
│   │   ├── job/
│   │   │   ├── entity.go                     # Job aggregate root (ID, ClientID, Title, Description, Budget, Skills, Status, etc.)
│   │   │   ├── value_objects.go             # JobBudget, JobSkill, JobCategory, JobDuration value objects
│   │   │   ├── enums.go                     # JobType (Fixed, Hourly), Status (Draft, Open, InProgress, Closed), ExperienceLevel, BudgetType
│   │   │   ├── visibility.go                # Public, Private, Invite-only visibility settings
│   │   │   ├── errors.go                    # Domain errors (ErrJobNotFound, ErrInvalidBudget, ErrJobAlreadyClosed)
│   │   │   └── repository.go                # JobRepository interface (Create, Update, FindByID, FindByClient, Search)
│   │   │
│   │   ├── category/
│   │   │   ├── entity.go                     # Category entity (ID, Name, Slug, ParentID, Level) - hierarchical categories
│   │   │   ├── subcategory.go               # Nested categories support (tree structure)
│   │   │   ├── errors.go                    # Category errors
│   │   │   └── repository.go                # CategoryRepository interface
│   │   │
│   │   ├── skill/
│   │   │   ├── entity.go                     # Skill taxonomy (ID, Name, CategoryID, Popularity)
│   │   │   ├── category.go                  # Skill categories (grouping related skills)
│   │   │   ├── errors.go                    # Skill errors
│   │   │   └── repository.go                # SkillRepository interface
│   │   │
│   │   ├── job_skill/
│   │   │   ├── entity.go                     # Skills required for job (JobID, SkillID, IsRequired)
│   │   │   ├── requirement_level.go         # Required vs Preferred skill levels
│   │   │   ├── errors.go                    # Job skill errors
│   │   │   └── repository.go                # JobSkillRepository interface
│   │   │
│   │   ├── job_question/
│   │   │   ├── entity.go                     # Screening questions (JobID, Question, QuestionType, IsRequired)
│   │   │   ├── question_type.go             # Text, MultiChoice, File upload question types
│   │   │   ├── errors.go                    # Question errors
│   │   │   └── repository.go                # JobQuestionRepository interface
│   │   │
│   │   ├── job_attachment/
│   │   │   ├── entity.go                     # Job attachments (JobID, FileURL, FileType, FileName)
│   │   │   ├── errors.go                    # Attachment errors
│   │   │   └── repository.go                # JobAttachmentRepository interface
│   │   │
│   │   ├── job_invitation/
│   │   │   ├── entity.go                     # Invitations to freelancers (JobID, FreelancerID, Message, Status)
│   │   │   ├── status.go                    # Invitation status (Pending, Accepted, Declined, Expired)
│   │   │   ├── errors.go                    # Invitation errors
│   │   │   └── repository.go                # JobInvitationRepository interface
│   │   │
│   │   ├── job_view/
│   │   │   ├── entity.go                     # Job view tracking (JobID, UserID, ViewedAt, Duration)
│   │   │   ├── errors.go                    # View tracking errors
│   │   │   └── repository.go                # JobViewRepository interface
│   │   │
│   │   ├── saved_search/
│   │   │   ├── entity.go                     # Saved search filters (UserID, SearchQuery, Filters JSON, AlertEnabled)
│   │   │   ├── errors.go                    # Saved search errors
│   │   │   └── repository.go                # SavedSearchRepository interface
│   │   │
│   │   └── job_flag/
│   │       ├── entity.go                     # Flagged jobs (JobID, ReporterID, Reason, Status, ReviewedAt)
│   │       ├── reason.go                    # Flag reasons (Spam, Inappropriate, Fraud, Misleading)
│   │       ├── status.go                    # Flag status (Pending, Reviewed, Resolved, Dismissed)
│   │       ├── errors.go                    # Flag errors
│   │       └── repository.go                # JobFlagRepository interface
│   │
│   ├── application/                          # 📋 Application Layer - Use cases & orchestration
│   │   ├── job/
│   │   │   ├── service.go                    # Job business logic (Create, Update, Close, Delete, Repost)
│   │   │   ├── commands.go                  # CreateJobCommand, UpdateJobCommand, PublishJobCommand, CloseJobCommand (validates client exists via users-be)
│   │   │   ├── queries.go                   # GetJobQuery, ListJobsQuery, SearchJobsQuery with filters (category, budget, skills)
│   │   │   ├── dto.go                       # JobDTO, CreateJobDTO, UpdateJobDTO, JobListDTO, JobSearchDTO
│   │   │   ├── mapper.go                    # Entity-DTO mapping
│   │   │   ├── validators.go                # Input validation (budget, skills, dates)
│   │   │   └── business_rules.go            # Business rules (can close job only with accepted proposal)
│   │   │
│   │   ├── category/
│   │   │   ├── service.go                    # Category business logic (List, GetTree, Search)
│   │   │   ├── dto.go                       # CategoryDTO, CategoryTreeDTO
│   │   │   ├── mapper.go                    # Category mappers
│   │   │   └── tree_builder.go              # Build hierarchical category tree
│   │   │
│   │   ├── skill/
│   │   │   ├── service.go                    # Skill business logic (List, Search, GetPopular)
│   │   │   ├── dto.go                       # SkillDTO, SkillListDTO
│   │   │   ├── mapper.go                    # Skill mappers
│   │   │   └── suggestion_service.go        # Suggest skills based on job description
│   │   │
│   │   ├── invitation/
│   │   │   ├── service.go                    # Invitation business logic (Send, Accept, Decline)
│   │   │   ├── dto.go                       # InvitationDTO, SendInvitationDTO
│   │   │   └── mapper.go                    # Invitation mappers
│   │   │
│   │   ├── search/
│   │   │   ├── service.go                    # Search business logic (Search, Filter, Sort)
│   │   │   ├── filter_builder.go            # Build complex search filters
│   │   │   └── dto.go                       # SearchDTO, FilterDTO
│   │   │
│   │   ├── flag/
│   │   │   ├── service.go                    # Flag business logic (Flag, Review, Resolve)
│   │   │   ├── commands.go                  # FlagJobCommand, UnflagJobCommand
│   │   │   ├── dto.go                       # FlagDTO, FlagJobDTO
│   │   │   └── mapper.go                    # Flag mappers
│   │   │
│   │   └── eventhandler/                    # 📝 UPDATED: Event handlers (uses contracts/events)
│   │       ├── user_handler.go              # Handle user events (client verification → allow job posting)
│   │       ├── proposal_handler.go          # Handle proposal events (proposal accepted → update job status)
│   │       ├── contract_handler.go          # Handle contract events (contract started → mark job filled)
│   │       ├── subscription_handler.go      # Handle subscription events (limit job postings per plan)
│   │       └── admin_handler.go             # Handle admin actions on jobs (remove, feature, hide)
│   │
│   ├── infrastructure/                       # 🔧 Infrastructure Layer
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection setup
│   │   │       ├── transaction.go           # Transaction helpers
│   │   │       ├── migrations.go            # 📝 UPDATED: Auto-migration with version tracking
│   │   │       ├── version.go               # 🆕 Schema version tracking
│   │   │       ├── safety.go                # 🆕 Pre-migration safety checks
│   │   │       ├── job_repository.go        # JobRepository implementation with GORM
│   │   │       ├── category_repository.go   # CategoryRepository implementation
│   │   │       ├── skill_repository.go      # SkillRepository implementation
│   │   │       ├── job_skill_repository.go  # JobSkillRepository implementation
│   │   │       ├── job_question_repository.go # JobQuestionRepository implementation
│   │   │       ├── job_attachment_repository.go # JobAttachmentRepository implementation
│   │   │       ├── job_invitation_repository.go # JobInvitationRepository implementation
│   │   │       ├── job_view_repository.go   # JobViewRepository implementation
│   │   │       ├── saved_search_repository.go # SavedSearchRepository implementation
│   │   │       └── job_flag_repository.go   # JobFlagRepository implementation
│   │   │
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go            # Redis connection
│   │   │       ├── job_cache.go             # Job caching (Get, Set, Invalidate)
│   │   │       ├── category_cache.go        # Category tree caching
│   │   │       └── skill_cache.go           # Popular skills caching
│   │   │
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go              # 📝 UPDATED: Uses platform-shared/inbox
│   │   │       ├── producer.go              # 📝 UPDATED: Uses platform-shared/outbox
│   │   │       ├── topics.go                # 📝 UPDATED: Constants from contracts/events (job.posted, job.closed)
│   │   │       └── scram.go                 # SCRAM authentication
│   │   │
│   │   ├── search/
│   │   │   └── client.go                    # Search service client (call search-be for indexing)
│   │   │
│   │   └── storage/
│   │       └── client.go                    # Storage service client (upload job attachments)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── job_handler.go           # Job HTTP handlers (POST /jobs, GET /jobs/:id, PUT /jobs/:id, DELETE /jobs/:id)
│   │       │   ├── category_handler.go      # Category handlers (GET /categories, GET /categories/tree)
│   │       │   ├── skill_handler.go         # Skill handlers (GET /skills, GET /skills/popular)
│   │       │   ├── invitation_handler.go    # Invitation handlers (POST /jobs/:id/invitations)
│   │       │   ├── search_handler.go        # Search handlers (GET /jobs/search)
│   │       │   ├── saved_search_handler.go  # Saved search handlers (POST /saved-searches)
│   │       │   ├── flag_handler.go          # Flag handlers (POST /jobs/:id/flag)
│   │       │   └── health_handler.go        # Health check endpoints
│   │       │
│   │       ├── middleware/                  # 📝 UPDATED: Uses platform-shared middleware
│   │       │   ├── auth.go                  # 📝 UPDATED: Uses pkg/auth
│   │       │   ├── rbac.go                  # 📝 UPDATED: Uses pkg/auth authorizer
│   │       │   ├── cors.go                  # 📝 UPDATED: Uses platform-shared/ginx
│   │       │   ├── rate_limit.go            # Rate limiting
│   │       │   └── idempotency.go           # 🆕 Uses platform-shared/idempotency
│   │       │
│   │       ├── responses/                   # 📝 UPDATED: Response helpers
│   │       │   └── README.md                # 📝 Points to platform-shared/httpx
│   │       │
│   │       └── router.go                    # 📝 UPDATED: Uses platform-shared/ginx middleware
│   │
│   └── config/                              # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                        # 🆕 Typed Config struct
│       ├── loader.go                        # 🆕 Config loader using Viper
│       └── docs/
│           └── CONFIGURATION.md             # 🆕 Configuration documentation
│
├── config/                                  # 🆕 MEDIUM - Configuration files
│   ├── default.yaml                         # 🆕 Default configuration
│   ├── dev.yaml                             # 🆕 Development overrides
│   └── prod.yaml                            # 🆕 Production overrides
│
├── dapr/                                    # 🆕 MEDIUM - Dapr components
│   ├── local/                               # 🆕 For dapr run
│   │   ├── pubsub.yaml                      # Kafka pub/sub component
│   │   └── statestore.yaml                  # State store component
│   └── k8s/                                 # 🆕 For kubectl apply
│       ├── pubsub.yaml                      # Kafka with scopes: ["jobs-be"]
│       ├── statestore.yaml                  # State store with scopes
│       └── secrets.yaml                     # Dapr secret store
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                        # Service-specific errors
│   │   └── codes.go                         # Error codes (JOB_NOT_FOUND, INVALID_BUDGET)
│   ├── logger/                              # ❌ REMOVED: Use platform-shared/logging
│   │   └── README.md                        # 📝 Points to platform-shared/logging
│   ├── utils/
│   │   ├── validator.go                     # Local validation utilities
│   │   ├── slug.go                          # Generate slugs
│   │   └── sanitizer.go                     # Sanitize input
│   └── constants/
│       ├── events.go                        # ❌ REMOVED: Use contracts/events
│       └── topics.go                        # ❌ REMOVED: Use contracts/events
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                  # Kubernetes Deployment
│       ├── service.yaml                     # Kubernetes Service
│       ├── configmap.yaml                   # ConfigMap
│       ├── secrets.yaml                     # Secrets
│       ├── hpa.yaml                         # HPA
│       ├── pdb.yaml                         # PDB
│       └── servicemonitor.yaml              # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                       # Setup local environment
│   ├── get-secrets.sh                       # Fetch secrets
│   ├── seed-categories.sh                   # Seed job categories
│   └── seed-skills.sh                       # Seed skills taxonomy
│
├── tests/
│   ├── unit/                                # Unit tests
│   ├── integration/                         # Integration tests
│   └── e2e/                                 # End-to-end tests
│
├── docs/
│   ├── README.md                            # Service overview
│   ├── API.md                               # API documentation
│   ├── EVENTS.md                            # 🆕 Events (published: job.posted, job.closed, consumed: user.verified, proposal.accepted)
│   ├── ARCHITECTURE.md                      # Architecture details
│   ├── MIGRATIONS.md                        # 🆕 Migration history
│   ├── SCHEMA.md                            # 🆕 Database schema
│   ├── RUNBOOK.md                           # 🆕 Operational procedures
│   ├── categories.md                        # Job categories documentation
│   └── skills-taxonomy.md                   # Skills taxonomy documentation
│
├── .github/
│   └── workflows/
│       ├── ci.yml                           # CI workflow
│       └── cd.yml                           # CD workflow
│
├── go.mod                                   # 📝 UPDATED: Imports pkg/auth, platform-shared, contracts/events
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md
```

---

## 📦 **3️⃣ proposals-be (Freelancer Proposal Submission Service)**

```
apps/be/proposals-be/
│
├── cmd/
│   └── api/
│       └── main.go                           # Entry point - initializes Gin, Dapr, connects to Postgres
│
├── internal/
│   ├── domain/                               # 🏛️ Domain Layer
│   │   ├── proposal/
│   │   │   ├── entity.go                     # Proposal aggregate root (ID, JobID, FreelancerID, CoverLetter, Amount, Duration)
│   │   │   ├── value_objects.go             # ProposalAmount, ProposalDuration
│   │   │   ├── enums.go                     # Status (Draft, Submitted, Accepted, Rejected, Withdrawn), Type (Hourly, Fixed)
│   │   │   ├── errors.go                    # Domain errors (ProposalNotFound, InvalidAmount, AlreadySubmitted)
│   │   │   └── repository.go                # ProposalRepository interface
│   │   │
│   │   ├── bid/
│   │   │   ├── entity.go                     # Bidding system (ProposalID, Amount, Rank, PlacedAt)
│   │   │   ├── bid_amount.go                # Bid amount management (validation, currency conversion)
│   │   │   ├── bid_history.go               # Bid history tracking (all bids for a job)
│   │   │   ├── auto_bid.go                  # Auto-bidding logic (automatic bid adjustments)
│   │   │   ├── bid_rank.go                  # Bid ranking (1st, 2nd, 3rd lowest bid)
│   │   │   ├── errors.go                    # Bid errors
│   │   │   └── repository.go                # BidRepository interface
│   │   │
│   │   ├── bid_strategy/
│   │   │   ├── entity.go                     # Bidding strategies (Aggressive, Conservative, Auto)
│   │   │   ├── aggressive.go                # Aggressive bidding strategy (undercut by X%)
│   │   │   ├── conservative.go              # Conservative bidding (competitive but safe)
│   │   │   ├── auto_strategy.go             # Auto-bid strategy (ML-based optimal bidding)
│   │   │   ├── errors.go                    # Strategy errors
│   │   │   └── repository.go                # BidStrategyRepository interface
│   │   │
│   │   ├── bid_notification/
│   │   │   ├── entity.go                     # Bid notifications (UserID, Message, Type, SentAt)
│   │   │   ├── outbid_alert.go              # Outbid alerts (notify when outbid)
│   │   │   ├── errors.go                    # Notification errors
│   │   │   └── repository.go                # BidNotificationRepository interface
│   │   │
│   │   ├── milestone/
│   │   │   ├── entity.go                     # Proposal milestones (ProposalID, Description, Amount, DueDate)
│   │   │   ├── milestone_group.go           # Group milestones (for fixed-price projects)
│   │   │   ├── errors.go                    # Milestone errors
│   │   │   └── repository.go                # MilestoneRepository interface
│   │   │
│   │   ├── cover_letter/
│   │   │   ├── entity.go                     # Cover letter content (ProposalID, Content, Version)
│   │   │   ├── errors.go                    # Cover letter errors
│   │   │   └── repository.go                # CoverLetterRepository interface
│   │   │
│   │   ├── proposal_attachment/
│   │   │   ├── entity.go                     # Proposal attachments (ProposalID, FileURL, FileType)
│   │   │   ├── errors.go                    # Attachment errors
│   │   │   └── repository.go                # AttachmentRepository interface
│   │   │
│   │   ├── proposal_question_answer/
│   │   │   ├── entity.go                     # Answers to job questions (ProposalID, QuestionID, Answer)
│   │   │   ├── errors.go                    # Answer errors
│   │   │   └── repository.go                # AnswerRepository interface
│   │   │
│   │   ├── template/
│   │   │   ├── entity.go                     # Proposal templates (FreelancerID, Title, Content)
│   │   │   ├── template_category.go         # Template categories (by job type, industry)
│   │   │   ├── errors.go                    # Template errors
│   │   │   └── repository.go                # TemplateRepository interface
│   │   │
│   │   ├── connect/
│   │   │   ├── entity.go                     # Connect usage tracking (FreelancerID, JobID, ConnectsUsed, UsedAt)
│   │   │   ├── transaction.go               # Connect transactions (purchase, use, refund)
│   │   │   ├── errors.go                    # Connect errors (InsufficientConnects)
│   │   │   └── repository.go                # ConnectRepository interface
│   │   │
│   │   ├── boost/
│   │   │   ├── entity.go                     # Boosted proposals (ProposalID, BoostLevel, ExpiresAt)
│   │   │   ├── boost_level.go               # Boost levels (Basic, Premium, Elite)
│   │   │   ├── errors.go                    # Boost errors
│   │   │   └── repository.go                # BoostRepository interface
│   │   │
│   │   ├── proposal_analytics/
│   │   │   ├── entity.go                     # Proposal performance analytics (Views, ClickThroughRate, ResponseRate)
│   │   │   ├── view_tracking.go             # Track proposal views
│   │   │   ├── errors.go                    # Analytics errors
│   │   │   └── repository.go                # AnalyticsRepository interface
│   │   │
│   │   └── proposal_flag/
│   │       ├── entity.go                     # Flagged proposals (ProposalID, ReporterID, Reason, Status)
│   │       ├── reason.go                    # Flag reasons (Spam, Plagiarism, Inappropriate)
│   │       ├── status.go                    # Flag status
│   │       ├── errors.go                    # Flag errors
│   │       └── repository.go                # FlagRepository interface
│   │
│   ├── application/                          # 📋 Application Layer
│   │   ├── proposal/
│   │   │   ├── service.go                    # Proposal business logic (Submit, Withdraw, Update)
│   │   │   ├── commands.go                  # SubmitProposalCommand, WithdrawProposalCommand, UpdateProposalCommand
│   │   │   ├── queries.go                   # GetProposal, ListProposals, FilterProposals
│   │   │   ├── dto.go                       # ProposalDTO, CreateProposalDTO, UpdateProposalDTO
│   │   │   ├── mapper.go                    # Entity-DTO mapping
│   │   │   ├── validators.go                # Validation (amount, duration, cover letter)
│   │   │   └── business_rules.go            # Business rules (cannot submit duplicate proposals)
│   │   │
│   │   ├── bid/
│   │   │   ├── service.go                    # Bid business logic
│   │   │   ├── commands.go                  # PlaceBidCommand, UpdateBidCommand, WithdrawBidCommand
│   │   │   ├── queries.go                   # GetBidStatus, GetBidHistory, GetBidRank
│   │   │   ├── dto.go                       # BidDTO, PlaceBidDTO
│   │   │   ├── mapper.go                    # Bid mappers
│   │   │   ├── bid_calculator.go            # Calculate optimal bid amounts
│   │   │   ├── bid_validator.go             # Validate bid amounts (min/max constraints)
│   │   │   ├── auto_bid_manager.go          # Auto-bidding management
│   │   │   └── ranking_engine.go            # Rank bids (1st, 2nd, 3rd)
│   │   │
│   │   ├── bid_strategy/
│   │   │   ├── service.go                    # Bid strategy logic
│   │   │   ├── strategy_executor.go         # Execute bidding strategies
│   │   │   ├── dto.go                       # StrategyDTO
│   │   │   └── mapper.go                    # Strategy mappers
│   │   │
│   │   ├── milestone/
│   │   │   ├── service.go                    # Milestone logic (Create, Update, Delete)
│   │   │   ├── dto.go                       # MilestoneDTO
│   │   │   └── mapper.go                    # Milestone mappers
│   │   │
│   │   ├── template/
│   │   │   ├── service.go                    # Template logic (Create, Update, Use)
│   │   │   ├── dto.go                       # TemplateDTO
│   │   │   └── mapper.go                    # Template mappers
│   │   │
│   │   ├── connect/
│   │   │   ├── service.go                    # Connect logic (Use, Purchase, Refund)
│   │   │   ├── calculator.go                # Calculate connect cost per job
│   │   │   └── dto.go                       # ConnectDTO
│   │   │
│   │   ├── boost/
│   │   │   ├── service.go                    # Boost logic (Boost proposal, Calculate cost)
│   │   │   └── dto.go                       # BoostDTO
│   │   │
│   │   ├── analytics/
│   │   │   ├── service.go                    # Analytics logic (Track views, Calculate metrics)
│   │   │   ├── performance_tracker.go       # Track proposal performance
│   │   │   └── dto.go                       # AnalyticsDTO
│   │   │
│   │   ├── flag/
│   │   │   ├── service.go                    # Flag logic
│   │   │   ├── commands.go                  # FlagProposal, UnflagProposal
│   │   │   ├── dto.go                       # FlagDTO
│   │   │   └── mapper.go                    # Flag mappers
│   │   │
│   │   └── eventhandler/                    # 📝 UPDATED: Uses contracts/events
│   │       ├── job_handler.go               # Handle job events (job closed → mark proposals as rejected)
│   │       ├── user_handler.go              # Handle user events (freelancer verified → allow proposals)
│   │       ├── contract_handler.go          # Handle contract events (contract created → mark proposal as accepted)
│   │       ├── subscription_handler.go      # Handle subscription events (connects purchased, limit reached)
│   │       ├── payment_handler.go           # Handle payment events (payment successful → release connects)
│   │       └── admin_handler.go             # Handle admin actions on proposals (remove, flag)
│   │
│   ├── infrastructure/                       # 🔧 Infrastructure Layer
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection
│   │   │       ├── transaction.go           # Transaction helpers
│   │   │       ├── migrations.go            # 📝 UPDATED: Auto-migration with version tracking
│   │   │       ├── version.go               # 🆕 Schema version tracking
│   │   │       ├── safety.go                # 🆕 Pre-migration safety checks
│   │   │       ├── proposal_repository.go   # ProposalRepository implementation
│   │   │       ├── bid_repository.go        # BidRepository implementation
│   │   │       ├── bid_strategy_repository.go # BidStrategyRepository implementation
│   │   │       ├── bid_notification_repository.go # NotificationRepository implementation
│   │   │       ├── milestone_repository.go  # MilestoneRepository implementation
│   │   │       ├── cover_letter_repository.go # CoverLetterRepository implementation
│   │   │       ├── attachment_repository.go # AttachmentRepository implementation
│   │   │       ├── question_answer_repository.go # AnswerRepository implementation
│   │   │       ├── template_repository.go   # TemplateRepository implementation
│   │   │       ├── connect_repository.go    # ConnectRepository implementation
│   │   │       ├── boost_repository.go      # BoostRepository implementation
│   │   │       ├── analytics_repository.go  # AnalyticsRepository implementation
│   │   │       └── proposal_flag_repository.go # FlagRepository implementation
│   │   │
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go            # Redis connection
│   │   │       ├── proposal_cache.go        # Proposal caching
│   │   │       └── bid_cache.go             # Bid ranking cache
│   │   │
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go              # 📝 UPDATED: Uses platform-shared/inbox
│   │   │       ├── producer.go              # 📝 UPDATED: Uses platform-shared/outbox
│   │   │       ├── topics.go                # 📝 UPDATED: From contracts/events (proposal.submitted, bid.placed, outbid.alert)
│   │   │       └── scram.go                 # SCRAM authentication
│   │   │
│   │   └── storage/
│   │       └── client.go                    # Storage service client (upload attachments)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── proposal_handler.go      # Proposal handlers (POST /proposals, GET /proposals/:id)
│   │       │   ├── bid_handler.go           # Bid handlers (POST /proposals/:id/bid)
│   │       │   ├── bid_strategy_handler.go  # Strategy handlers (POST /bid-strategies)
│   │       │   ├── milestone_handler.go     # Milestone handlers (POST /proposals/:id/milestones)
│   │       │   ├── template_handler.go      # Template handlers (GET /templates, POST /templates)
│   │       │   ├── connect_handler.go       # Connect handlers (GET /connects/balance)
│   │       │   ├── boost_handler.go         # Boost handlers (POST /proposals/:id/boost)
│   │       │   ├── analytics_handler.go     # Analytics handlers (GET /proposals/:id/analytics)
│   │       │   ├── flag_handler.go          # Flag handlers (POST /proposals/:id/flag)
│   │       │   └── health_handler.go        # Health check
│   │       │
│   │       ├── middleware/                  # 📝 UPDATED: Uses platform-shared
│   │       │   ├── auth.go                  # 📝 UPDATED: Uses pkg/auth
│   │       │   ├── rbac.go                  # 📝 UPDATED: Uses pkg/auth
│   │       │   ├── cors.go                  # 📝 UPDATED: Uses platform-shared/ginx
│   │       │   ├── rate_limit.go            # Rate limiting
│   │       │   └── idempotency.go           # 🆕 Uses platform-shared/idempotency
│   │       │
│   │       ├── responses/
│   │       │   └── README.md                # 📝 Points to platform-shared/httpx
│   │       │
│   │       └── router.go                    # 📝 UPDATED: Uses platform-shared/ginx
│   │
│   └── config/                              # 🆕 MEDIUM
│       ├── schema.go                        # 🆕 Typed Config
│       ├── loader.go                        # 🆕 Viper loader
│       └── docs/
│           └── CONFIGURATION.md             # 🆕 Config docs
│
├── config/                                  # 🆕 MEDIUM
│   ├── default.yaml
│   ├── dev.yaml
│   └── prod.yaml
│
├── dapr/                                    # 🆕 MEDIUM
│   ├── local/
│   │   ├── pubsub.yaml
│   │   └── statestore.yaml
│   └── k8s/
│       ├── pubsub.yaml                      # Scopes: ["proposals-be"]
│       ├── statestore.yaml
│       └── secrets.yaml
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   └── codes.go                         # PROPOSAL_NOT_FOUND, INSUFFICIENT_CONNECTS
│   ├── logger/                              # ❌ REMOVED
│   │   └── README.md
│   ├── utils/
│   │   ├── validator.go
│   │   ├── sanitizer.go
│   │   └── bid_calculator.go                # Bid calculation utilities
│   └── constants/
│       ├── events.go                        # ❌ REMOVED
│       ├── topics.go                        # ❌ REMOVED
│       └── bid_types.go                     # Bid type constants
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
│   ├── EVENTS.md                            # 🆕 Events (proposal.submitted, bid.placed, outbid.alert, connect.used)
│   ├── ARCHITECTURE.md
│   ├── MIGRATIONS.md                        # 🆕
│   ├── SCHEMA.md                            # 🆕
│   ├── RUNBOOK.md                           # 🆕
│   ├── bidding-system.md                    # Bidding system documentation
│   ├── connects-system.md                   # Connects/credits documentation
│   └── auto-bidding.md                      # Auto-bidding strategy documentation
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod                                   # 📝 UPDATED: Imports pkg/auth, platform-shared, contracts/events
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md
```

---

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
│   │   │   ├── terms.go                     # Contract terms (payment terms, deliverables, timeline)
│   │   │   ├── enums.go                     # Type (Fixed, Hourly), Status (Active, Paused, Completed, Terminated), PaymentType
│   │   │   ├── value_objects.go             # ContractAmount, ContractDuration
│   │   │   ├── errors.go                    # Domain errors (ContractNotFound, InvalidTerms)
│   │   │   └── repository.go                # ContractRepository interface
│   │   │
│   │   ├── milestone/
│   │   │   ├── entity.go                     # Contract milestones (ContractID, Description, Amount, DueDate, Status)
│   │   │   ├── delivery.go                  # Milestone deliverables (files, descriptions)
│   │   │   ├── status.go                    # Milestone status tracking (InProgress, Submitted, UnderReview, Approved, Rejected)
│   │   │   ├── errors.go                    # Milestone errors
│   │   │   └── repository.go                # MilestoneRepository interface
│   │   │
│   │   ├── timesheet/
│   │   │   ├── entity.go                     # Time entries (ContractID, WeekStartDate, TotalHours, Status)
│   │   │   ├── week.go                      # Weekly timesheets (aggregated by week)
│   │   │   ├── time_entry.go                # Individual time entry (Date, Hours, Description, Manual/Tracked)
│   │   │   ├── screenshot.go                # Work diary screenshots (TimeEntryID, URL, Timestamp)
│   │   │   ├── errors.go                    # Timesheet errors
│   │   │   └── repository.go                # TimesheetRepository interface
│   │   │
│   │   ├── workdiary/
│   │   │   ├── entity.go                     # Work diary entries (ContractID, Date, Hours, Screenshots, ActivityLevel)
│   │   │   ├── activity_level.go            # Activity tracking (keyboard/mouse activity percentage)
│   │   │   ├── errors.go                    # Work diary errors
│   │   │   └── repository.go                # WorkDiaryRepository interface
│   │   │
│   │   ├── template/
│   │   │   ├── entity.go                     # Contract templates (Title, Terms, Clauses)
│   │   │   ├── clause.go                    # Contract clauses (standard terms, custom clauses)
│   │   │   ├── errors.go                    # Template errors
│   │   │   └── repository.go                # TemplateRepository interface
│   │   │
│   │   ├── amendment/
│   │   │   ├── entity.go                     # Contract amendments (ContractID, Type, OldValue, NewValue, Status)
│   │   │   ├── change_request.go            # Change requests (scope, budget, timeline changes)
│   │   │   ├── errors.go                    # Amendment errors
│   │   │   └── repository.go                # AmendmentRepository interface
│   │   │
│   │   ├── deliverable/
│   │   │   ├── entity.go                     # Contract deliverables (ContractID, Description, FileURLs, SubmittedAt)
│   │   │   ├── submission.go                # Deliverable submissions (version tracking)
│   │   │   ├── revision.go                  # Revision requests (feedback, changes needed)
│   │   │   ├── errors.go                    # Deliverable errors
│   │   │   └── repository.go                # DeliverableRepository interface
│   │   │
│   │   ├── pause/
│   │   │   ├── entity.go                     # Contract pause/resume (ContractID, Reason, PausedAt, ResumedAt)
│   │   │   ├── errors.go                    # Pause errors
│   │   │   └── repository.go                # PauseRepository interface
│   │   │
│   │   ├── termination/
│   │   │   ├── entity.go                     # Contract termination (ContractID, Reason, TerminatedBy, TerminatedAt)
│   │   │   ├── reason.go                    # Termination reasons (Completed, Breach, Mutual, Unilateral)
│   │   │   ├── errors.go                    # Termination errors
│   │   │   └── repository.go                # TerminationRepository interface
│   │   │
│   │   └── dispute/
│   │       ├── entity.go                     # Contract disputes (ContractID, Reason, FiledBy, Status, Resolution)
│   │       ├── evidence.go                  # Evidence submission (files, descriptions, timestamps)
│   │       ├── resolution.go                # Dispute resolution (outcome, awarded amount, penalties)
│   │       ├── status.go                    # Dispute status (Open, UnderReview, Mediation, Resolved, Escalated)
│   │       ├── errors.go                    # Dispute errors
│   │       └── repository.go                # DisputeRepository interface
│   │
│   ├── application/                          # 📋 Application Layer
│   │   ├── contract/
│   │   │   ├── service.go                    # Contract logic (Create, Update, Pause, Resume, End)
│   │   │   ├── commands.go                  # CreateContract, UpdateContract, PauseContract, EndContract
│   │   │   ├── queries.go                   # GetContract, ListContracts, FilterContracts
│   │   │   ├── dto.go                       # ContractDTO, CreateContractDTO, UpdateContractDTO
│   │   │   ├── mapper.go                    # Entity-DTO mapping
│   │   │   ├── validators.go                # Contract validation (terms, amounts, dates)
│   │   │   └── business_rules.go            # Business rules (payment release conditions)
│   │   │
│   │   ├── milestone/
│   │   │   ├── service.go                    # Milestone logic (Create, Complete, Approve, Request)
│   │   │   ├── commands.go                  # CreateMilestone, CompleteMilestone, ApproveMilestone, RequestChanges
│   │   │   ├── dto.go                       # MilestoneDTO
│   │   │   └── mapper.go                    # Milestone mappers
│   │   │
│   │   ├── timesheet/
│   │   │   ├── service.go                    # Timesheet logic (Log time, Submit week, Approve)
│   │   │   ├── commands.go                  # LogTime, SubmitWeek, ApproveTimesheet, DisputeHours
│   │   │   ├── queries.go                   # GetTimesheets, GetWeeklySummary, Calculate earnings
│   │   │   ├── calculator.go                # Calculate hours, earnings based on hourly rate
│   │   │   ├── dto.go                       # TimesheetDTO, TimeEntryDTO
│   │   │   ├── mapper.go                    # Timesheet mappers
│   │   │   └── validators.go                # Timesheet validation (max hours, dates)
│   │   │
│   │   ├── workdiary/
│   │   │   ├── service.go                    # Work diary logic (Track activity, Upload screenshots)
│   │   │   ├── dto.go                       # WorkDiaryDTO
│   │   │   └── mapper.go                    # Work diary mappers
│   │   │
│   │   ├── template/
│   │   │   ├── service.go                    # Template logic (Create, Update, Use template)
│   │   │   ├── dto.go                       # TemplateDTO
│   │   │   └── mapper.go                    # Template mappers
│   │   │
│   │   ├── amendment/
│   │   │   ├── service.go                    # Amendment logic (Request, Approve, Reject)
│   │   │   ├── dto.go                       # AmendmentDTO
│   │   │   └── mapper.go                    # Amendment mappers
│   │   │
│   │   ├── deliverable/
│   │   │   ├── service.go                    # Deliverable logic (Submit, Review, Request revisions)
│   │   │   ├── dto.go                       # DeliverableDTO
│   │   │   └── mapper.go                    # Deliverable mappers
│   │   │
│   │   ├── termination/
│   │   │   ├── service.go                    # Termination logic (Terminate contract, Calculate final payment)
│   │   │   ├── dto.go                       # TerminationDTO
│   │   │   └── mapper.go                    # Termination mappers
│   │   │
│   │   ├── dispute/
│   │   │   ├── service.go                    # Dispute logic (Open, Submit evidence, Resolve)
│   │   │   ├── commands.go                  # OpenDispute, SubmitEvidence, ResolveDispute, EscalateDispute
│   │   │   ├── queries.go                   # GetDisputes, GetDisputeHistory
│   │   │   ├── dto.go                       # DisputeDTO, EvidenceDTO, ResolutionDTO
│   │   │   └── mapper.go                    # Dispute mappers
│   │   │
│   │   └── eventhandler/                    # 📝 UPDATED: Uses contracts/events
│   │       ├── proposal_handler.go          # Handle proposal events (proposal accepted → create contract)
│   │       ├── user_handler.go              # Handle user events (user suspended → pause contracts)
│   │       ├── payment_handler.go           # Handle payment events (payment received → release milestone)
│   │       ├── escrow_handler.go            # Handle escrow events (funds released → mark milestone paid)
│   │       ├── dispute_handler.go           # Handle dispute events from admin-be
│   │       └── admin_handler.go             # Handle admin dispute resolutions
│   │
│   ├── infrastructure/                       # 🔧 Infrastructure Layer
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection
│   │   │       ├── transaction.go           # Transaction helpers
│   │   │       ├── migrations.go            # 📝 UPDATED: Auto-migration with version tracking
│   │   │       ├── version.go               # 🆕 Schema version tracking
│   │   │       ├── safety.go                # 🆕 Pre-migration safety checks
│   │   │       ├── contract_repository.go   # ContractRepository implementation
│   │   │       ├── milestone_repository.go  # MilestoneRepository implementation
│   │   │       ├── timesheet_repository.go  # TimesheetRepository implementation
│   │   │       ├── workdiary_repository.go  # WorkDiaryRepository implementation
│   │   │       ├── template_repository.go   # TemplateRepository implementation
│   │   │       ├── amendment_repository.go  # AmendmentRepository implementation
│   │   │       ├── deliverable_repository.go # DeliverableRepository implementation
│   │   │       ├── pause_repository.go      # PauseRepository implementation
│   │   │       ├── termination_repository.go # TerminationRepository implementation
│   │   │       └── dispute_repository.go    # DisputeRepository implementation
│   │   │
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go            # Redis connection
│   │   │       └── contract_cache.go        # Contract caching
│   │   │
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go              # 📝 UPDATED: Uses platform-shared/inbox
│   │   │       ├── producer.go              # 📝 UPDATED: Uses platform-shared/outbox
│   │   │       ├── topics.go                # 📝 UPDATED: From contracts/events (contract.created, milestone.completed)
│   │   │       └── scram.go                 # SCRAM authentication
│   │   │
│   │   └── storage/
│   │       └── client.go                    # Storage service client (upload deliverables, screenshots)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── contract_handler.go      # Contract handlers (POST /contracts, GET /contracts/:id)
│   │       │   ├── milestone_handler.go     # Milestone handlers (POST /contracts/:id/milestones)
│   │       │   ├── timesheet_handler.go     # Timesheet handlers (POST /contracts/:id/timesheet)
│   │       │   ├── workdiary_handler.go     # Work diary handlers (POST /contracts/:id/workdiary)
│   │       │   ├── template_handler.go      # Template handlers (GET /templates)
│   │       │   ├── amendment_handler.go     # Amendment handlers (POST /contracts/:id/amendments)
│   │       │   ├── deliverable_handler.go   # Deliverable handlers (POST /contracts/:id/deliverables)
│   │       │   ├── termination_handler.go   # Termination handlers (POST /contracts/:id/terminate)
│   │       │   ├── dispute_handler.go       # Dispute handlers (POST /contracts/:id/disputes)
│   │       │   └── health_handler.go        # Health check
│   │       │
│   │       ├── middleware/                  # 📝 UPDATED: Uses platform-shared
│   │       │   ├── auth.go                  # 📝 UPDATED: Uses pkg/auth
│   │       │   ├── rbac.go                  # 📝 UPDATED: Uses pkg/auth
│   │       │   ├── cors.go                  # 📝 UPDATED: Uses platform-shared/ginx
│   │       │   ├── rate_limit.go            # Rate limiting
│   │       │   └── idempotency.go           # 🆕 Uses platform-shared/idempotency
│   │       │
│   │       ├── responses/
│   │       │   └── README.md                # 📝 Points to platform-shared/httpx
│   │       │
│   │       └── router.go                    # 📝 UPDATED: Uses platform-shared/ginx
│   │
│   └── config/                              # 🆕 MEDIUM
│       ├── schema.go                        # 🆕 Typed Config
│       ├── loader.go                        # 🆕 Viper loader
│       └── docs/
│           └── CONFIGURATION.md             # 🆕 Config docs
│
├── config/                                  # 🆕 MEDIUM
│   ├── default.yaml
│   ├── dev.yaml
│   └── prod.yaml
│
├── dapr/                                    # 🆕 MEDIUM
│   ├── local/
│   │   ├── pubsub.yaml
│   │   └── statestore.yaml
│   └── k8s/
│       ├── pubsub.yaml                      # Scopes: ["contracts-be"]
│       ├── statestore.yaml
│       └── secrets.yaml
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   └── codes.go                         # CONTRACT_NOT_FOUND, MILESTONE_NOT_COMPLETED
│   ├── logger/                              # ❌ REMOVED
│   │   └── README.md
│   ├── utils/
│   │   ├── validator.go
│   │   ├── time_calculator.go               # Time calculation utilities
│   │   └── date_utils.go                    # Date utilities
│   └── constants/
│       ├── events.go                        # ❌ REMOVED
│       └── topics.go                        # ❌ REMOVED
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
│   ├── EVENTS.md                            # 🆕 Events (contract.created, milestone.completed, timesheet.submitted, dispute.opened)
│   ├── ARCHITECTURE.md
│   ├── MIGRATIONS.md                        # 🆕
│   ├── SCHEMA.md                            # 🆕
│   ├── RUNBOOK.md                           # 🆕
│   ├── contract-lifecycle.md                # Contract lifecycle documentation
│   └── timesheet-system.md                  # Timesheet tracking system
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod                                   # 📝 UPDATED: Imports pkg/auth, platform-shared, contracts/events
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
│   │   │   ├── balance.go                   # Balance tracking (Available, Pending, Reserved)
│   │   │   ├── currency.go                  # Multi-currency support (USD, EUR, GBP, etc.)
│   │   │   ├── errors.go                    # Wallet errors (InsufficientFunds, WalletNotFound)
│   │   │   └── repository.go                # WalletRepository interface
│   │   │
│   │   ├── transaction/
│   │   │   ├── entity.go                     # Financial transactions (ID, WalletID, Amount, Type, Status, Reference)
│   │   │   ├── enums.go                     # Type (Deposit, Withdrawal, Transfer, Payment, Refund), Status (Pending, Completed, Failed), Category
│   │   │   ├── ledger.go                    # Double-entry ledger (debit/credit accounting)
│   │   │   ├── errors.go                    # Transaction errors (TransactionFailed, DuplicateTransaction)
│   │   │   └── repository.go                # TransactionRepository interface
│   │   │
│   │   ├── payment/
│   │   │   ├── entity.go                     # Payment records (ID, Amount, Currency, PayerID, PayeeID, Method, Status)
│   │   │   ├── payment_method.go            # Payment methods (CreditCard, PayPal, BankTransfer, Wallet)
│   │   │   ├── gateway.go                   # Payment gateway abstraction (Stripe, PayPal interface)
│   │   │   ├── errors.go                    # Payment errors (PaymentFailed, InvalidCard)
│   │   │   └── repository.go                # PaymentRepository interface
│   │   │
│   │   ├── escrow/
│   │   │   ├── entity.go                     # Escrow accounts (ID, ContractID, Amount, Status, HeldAt, ReleasedAt)
│   │   │   ├── hold.go                      # Fund holds (amount reserved until milestone completion)
│   │   │   ├── release.go                   # Fund release conditions (milestone approved, dispute resolved)
│   │   │   ├── errors.go                    # Escrow errors (EscrowNotFound, InsufficientEscrow)
│   │   │   └── repository.go                # EscrowRepository interface
│   │   │
│   │   ├── payout/
│   │   │   ├── entity.go                     # Payout requests (ID, UserID, Amount, Method, Status, RequestedAt, ProcessedAt)
│   │   │   ├── method.go                    # Payout methods (BankTransfer, PayPal, Payoneer, Wire)
│   │   │   ├── schedule.go                  # Payout scheduling (instant, daily, weekly, monthly)
│   │   │   ├── errors.go                    # Payout errors (PayoutFailed, BelowMinimum)
│   │   │   └── repository.go                # PayoutRepository interface
│   │   │
│   │   ├── invoice/
│   │   │   ├── entity.go                     # Invoice generation (ID, ContractID, Number, Amount, DueDate, Status, PaidAt)
│   │   │   ├── line_item.go                 # Invoice line items (description, quantity, price, total)
│   │   │   ├── tax.go                       # Tax calculations (VAT, sales tax by jurisdiction)
│   │   │   ├── errors.go                    # Invoice errors (InvoiceNotFound, AlreadyPaid)
│   │   │   └── repository.go                # InvoiceRepository interface
│   │   │
│   │   ├── fee/
│   │   │   ├── entity.go                     # Platform fees (ID, TransactionID, Amount, Type, Rate)
│   │   │   ├── calculator.go                # Fee calculation rules (percentage, tiered rates)
│   │   │   ├── tier.go                      # Fee tiers (based on subscription, volume, user type)
│   │   │   ├── errors.go                    # Fee errors
│   │   │   └── repository.go                # FeeRepository interface
│   │   │
│   │   ├── refund/
│   │   │   ├── entity.go                     # Refund processing (ID, PaymentID, Amount, Reason, Status, ProcessedAt)
│   │   │   ├── policy.go                    # Refund policies (full, partial, time limits)
│   │   │   ├── errors.go                    # Refund errors (RefundNotAllowed, RefundExpired)
│   │   │   └── repository.go                # RefundRepository interface
│   │   │
│   │   ├── dispute_payment/
│   │   │   ├── entity.go                     # Payment disputes (ID, PaymentID, Reason, FiledBy, Status, Resolution)
│   │   │   ├── chargeback.go                # Chargeback handling (from credit card companies)
│   │   │   ├── errors.go                    # Dispute errors
│   │   │   └── repository.go                # DisputePaymentRepository interface
│   │   │
│   │   ├── tax/
│   │   │   ├── entity.go                     # Tax records (ID, UserID, Year, Type, Amount, Status)
│   │   │   ├── form.go                      # Tax forms (W9, 1099, VAT returns)
│   │   │   ├── errors.go                    # Tax errors
│   │   │   └── repository.go                # TaxRepository interface
│   │   │
│   │   ├── bank_account/
│   │   │   ├── entity.go                     # Bank account info (ID, UserID, AccountNumber, RoutingNumber, BankName, Status)
│   │   │   ├── verification.go              # Micro-deposit verification (verify account ownership)
│   │   │   ├── errors.go                    # Bank account errors (InvalidAccount, VerificationFailed)
│   │   │   └── repository.go                # BankAccountRepository interface
│   │   │
│   │   └── payment_schedule/
│   │       ├── entity.go                     # Recurring payment schedules (ID, ContractID, Amount, Frequency, NextPaymentDate)
│   │       ├── errors.go                    # Schedule errors
│   │       └── repository.go                # PaymentScheduleRepository interface
│   │
│   ├── application/                          # 📋 Application Layer
│   │   ├── wallet/
│   │   │   ├── service.go                    # Wallet logic (Create, Deposit, Withdraw, Transfer, GetBalance)
│   │   │   ├── commands.go                  # DepositCommand, WithdrawCommand, TransferCommand
│   │   │   ├── queries.go                   # GetBalance, GetHistory, GetTransactions
│   │   │   ├── dto.go                       # WalletDTO, TransactionDTO
│   │   │   └── mapper.go                    # Wallet mappers
│   │   │
│   │   ├── transaction/
│   │   │   ├── service.go                    # Transaction logic (Create, Reverse, Reconcile)
│   │   │   ├── commands.go                  # CreateTransaction, ReverseTransaction, ReconcileTransactions
│   │   │   ├── queries.go                   # GetTransaction, ListTransactions, GetLedger
│   │   │   ├── dto.go                       # TransactionDTO, LedgerDTO
│   │   │   ├── mapper.go                    # Transaction mappers
│   │   │   └── reconciliation.go            # Reconciliation logic (match payments with bank statements)
│   │   │
│   │   ├── payment/
│   │   │   ├── service.go                    # Payment logic (Process, Capture, Void, Refund)
│   │   │   ├── commands.go                  # ProcessPayment, CapturePayment, VoidPayment, RefundPayment
│   │   │   ├── queries.go                   # GetPayment, ListPayments
│   │   │   ├── dto.go                       # PaymentDTO, ProcessPaymentDTO
│   │   │   ├── mapper.go                    # Payment mappers
│   │   │   ├── stripe_processor.go          # Stripe payment processor implementation
│   │   │   ├── paypal_processor.go          # PayPal payment processor implementation
│   │   │   └── processor_factory.go         # Payment processor factory (select processor by method)
│   │   │
│   │   ├── escrow/
│   │   │   ├── service.go                    # Escrow logic (Hold, Release, Refund)
│   │   │   ├── commands.go                  # HoldEscrow, ReleaseEscrow, RefundEscrow
│   │   │   ├── dto.go                       # EscrowDTO, HoldEscrowDTO
│   │   │   ├── mapper.go                    # Escrow mappers
│   │   │   └── release_manager.go           # Escrow release manager (handle release conditions)
│   │   │
│   │   ├── payout/
│   │   │   ├── service.go                    # Payout logic (Request, Process, Cancel, GetHistory)
│   │   │   ├── commands.go                  # RequestPayout, ProcessPayout, CancelPayout
│   │   │   ├── queries.go                   # GetPayouts, GetPayoutHistory
│   │   │   ├── dto.go                       # PayoutDTO, RequestPayoutDTO
│   │   │   ├── mapper.go                    # Payout mappers
│   │   │   └── batch_processor.go           # Batch payout processor (process multiple payouts together)
│   │   │
│   │   ├── invoice/
│   │   │   ├── service.go                    # Invoice logic (Generate, Send, MarkPaid, Cancel)
│   │   │   ├── commands.go                  # GenerateInvoice, SendInvoice, MarkInvoicePaid
│   │   │   ├── queries.go                   # GetInvoice, ListInvoices
│   │   │   ├── dto.go                       # InvoiceDTO, GenerateInvoiceDTO
│   │   │   ├── mapper.go                    # Invoice mappers
│   │   │   ├── generator.go                 # Invoice PDF generation
│   │   │   └── tax_calculator.go            # Tax calculation (VAT, sales tax)
│   │   │
│   │   ├── fee/
│   │   │   ├── service.go                    # Fee logic (Calculate, Apply, Waive)
│   │   │   ├── calculator.go                # Fee calculation (percentage, tiered, flat)
│   │   │   ├── dto.go                       # FeeDTO
│   │   │   └── rules_engine.go              # Fee rules engine (apply rules based on user type, volume)
│   │   │
│   │   ├── refund/
│   │   │   ├── service.go                    # Refund logic (Process, Cancel, GetHistory)
│   │   │   ├── commands.go                  # ProcessRefund, CancelRefund
│   │   │   ├── dto.go                       # RefundDTO
│   │   │   └── mapper.go                    # Refund mappers
│   │   │
│   │   ├── tax/
│   │   │   ├── service.go                    # Tax logic (Calculate, Generate forms, File)
│   │   │   ├── form_generator.go            # Generate tax forms (W9, 1099)
│   │   │   ├── dto.go                       # TaxDTO
│   │   │   └── mapper.go                    # Tax mappers
│   │   │
│   │   ├── bank_account/
│   │   │   ├── service.go                    # Bank account logic (Add, Verify, Remove)
│   │   │   ├── commands.go                  # AddBankAccount, VerifyBankAccount, RemoveBankAccount
│   │   │   ├── dto.go                       # BankAccountDTO
│   │   │   └── mapper.go                    # Bank account mappers
│   │   │
│   │   └── eventhandler/                    # 📝 UPDATED: Uses contracts/events
│   │       ├── contract_handler.go          # Handle contract events (contract created → hold escrow)
│   │       ├── milestone_handler.go         # Handle milestone events (milestone approved → release escrow)
│   │       ├── dispute_handler.go           # Handle dispute events (dispute resolved → process payment)
│   │       ├── user_handler.go              # Handle user events (user created → create wallet)
│   │       ├── subscription_handler.go      # Handle subscription events (subscription created → process payment)
│   │       └── admin_handler.go             # Handle admin actions on payments (refund, chargeback)
│   │
│   ├── infrastructure/                       # 🔧 Infrastructure Layer
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection
│   │   │       ├── transaction.go           # Transaction helpers
│   │   │       ├── migrations.go            # 📝 UPDATED: Auto-migration with version tracking
│   │   │       ├── version.go               # 🆕 Schema version tracking
│   │   │       ├── safety.go                # 🆕 Pre-migration safety checks
│   │   │       ├── wallet_repository.go     # WalletRepository implementation
│   │   │       ├── transaction_repository.go # TransactionRepository implementation
│   │   │       ├── payment_repository.go    # PaymentRepository implementation
│   │   │       ├── escrow_repository.go     # EscrowRepository implementation
│   │   │       ├── payout_repository.go     # PayoutRepository implementation
│   │   │       ├── invoice_repository.go    # InvoiceRepository implementation
│   │   │       ├── fee_repository.go        # FeeRepository implementation
│   │   │       ├── refund_repository.go     # RefundRepository implementation
│   │   │       ├── dispute_payment_repository.go # DisputePaymentRepository implementation
│   │   │       ├── tax_repository.go        # TaxRepository implementation
│   │   │       ├── bank_account_repository.go # BankAccountRepository implementation
│   │   │       └── payment_schedule_repository.go # PaymentScheduleRepository implementation
│   │   │
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go            # Redis connection
│   │   │       ├── wallet_cache.go          # Wallet balance caching
│   │   │       └── rate_cache.go            # Exchange rate caching
│   │   │
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go              # 📝 UPDATED: Uses platform-shared/inbox
│   │   │       ├── producer.go              # 📝 UPDATED: Uses platform-shared/outbox
│   │   │       ├── topics.go                # 📝 UPDATED: From contracts/events (payment.processed, escrow.released, payout.processed)
│   │   │       └── scram.go                 # SCRAM authentication
│   │   │
│   │   ├── payment_gateway/
│   │   │   ├── stripe/
│   │   │   │   ├── client.go                # Stripe API client
│   │   │   │   ├── webhook_handler.go       # Stripe webhook handler (payment.succeeded, charge.failed)
│   │   │   │   └── mapper.go                # Stripe event mapper
│   │   │   ├── paypal/
│   │   │   │   ├── client.go                # PayPal API client
│   │   │   │   ├── webhook_handler.go       # PayPal webhook handler
│   │   │   │   └── mapper.go                # PayPal event mapper
│   │   │   └── factory.go                   # Payment gateway factory
│   │   │
│   │   └── pdf/
│   │       └── generator.go                 # PDF invoice generation (using library like wkhtmltopdf)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── wallet_handler.go        # Wallet handlers (GET /wallets/:id, POST /wallets/:id/deposit)
│   │       │   ├── transaction_handler.go   # Transaction handlers (GET /transactions)
│   │       │   ├── payment_handler.go       # Payment handlers (POST /payments, GET /payments/:id)
│   │       │   ├── escrow_handler.go        # Escrow handlers (POST /escrow/hold, POST /escrow/release)
│   │       │   ├── payout_handler.go        # Payout handlers (POST /payouts, GET /payouts/:id)
│   │       │   ├── invoice_handler.go       # Invoice handlers (POST /invoices, GET /invoices/:id)
│   │       │   ├── fee_handler.go           # Fee handlers (GET /fees/calculate)
│   │       │   ├── refund_handler.go        # Refund handlers (POST /refunds)
│   │       │   ├── tax_handler.go           # Tax handlers (GET /tax/forms)
│   │       │   ├── bank_account_handler.go  # Bank account handlers (POST /bank-accounts)
│   │       │   ├── webhook_handler.go       # Webhook handlers (POST /webhooks/stripe, POST /webhooks/paypal)
│   │       │   └── health_handler.go        # Health check
│   │       │
│   │       ├── middleware/                  # 📝 UPDATED: Uses platform-shared
│   │       │   ├── auth.go                  # 📝 UPDATED: Uses pkg/auth
│   │       │   ├── rbac.go                  # 📝 UPDATED: Uses pkg/auth
│   │       │   ├── cors.go                  # 📝 UPDATED: Uses platform-shared/ginx
│   │       │   ├── rate_limit.go            # Rate limiting
│   │       │   └── idempotency.go           # 🆕 Uses platform-shared/idempotency (critical for payments!)
│   │       │
│   │       ├── responses/
│   │       │   └── README.md                # 📝 Points to platform-shared/httpx
│   │       │
│   │       └── router.go                    # 📝 UPDATED: Uses platform-shared/ginx
│   │
│   └── config/                              # 🆕 MEDIUM
│       ├── schema.go                        # 🆕 Typed Config (includes Stripe/PayPal keys)
│       ├── loader.go                        # 🆕 Viper loader
│       └── docs/
│           └── CONFIGURATION.md             # 🆕 Config docs
│
├── config/                                  # 🆕 MEDIUM
│   ├── default.yaml
│   ├── dev.yaml
│   └── prod.yaml
│
├── dapr/                                    # 🆕 MEDIUM
│   ├── local/
│   │   ├── pubsub.yaml
│   │   └── statestore.yaml
│   └── k8s/
│       ├── pubsub.yaml                      # Scopes: ["financial-be"]
│       ├── statestore.yaml
│       └── secrets.yaml
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   ├── codes.go                         # INSUFFICIENT_FUNDS, PAYMENT_FAILED
│   │   └── payment_errors.go                # Payment-specific error types
│   ├── logger/                              # ❌ REMOVED
│   │   └── README.md
│   ├── utils/
│   │   ├── validator.go
│   │   ├── currency.go                      # Currency conversion utilities
│   │   ├── decimal.go                       # Decimal math for money (avoid floating point errors)
│   │   └── encryption.go                    # Encryption utilities (for sensitive data like card numbers)
│   └── constants/
│       ├── events.go                        # ❌ REMOVED
│       ├── topics.go                        # ❌ REMOVED
│       ├── currencies.go                    # Currency constants (USD, EUR, GBP)
│       └── payment_methods.go               # Payment method constants
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml                     # Contains Stripe/PayPal API keys
│       ├── hpa.yaml
│       ├── pdb.yaml
│       ├── networkpolicy.yaml               # Extra security for financial service
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   ├── seed-data.sh
│   └── reconciliation.sh                    # Daily reconciliation script
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── API.md
│   ├── EVENTS.md                            # 🆕 Events (payment.processed, escrow.held, escrow.released, payout.processed)
│   ├── ARCHITECTURE.md
│   ├── MIGRATIONS.md                        # 🆕
│   ├── SCHEMA.md                            # 🆕
│   ├── RUNBOOK.md                           # 🆕
│   ├── payment-flows.md                     # Payment flow documentation
│   ├── escrow-system.md                     # Escrow system documentation
│   ├── fee-structure.md                     # Platform fee structure
│   └── compliance.md                        # PCI DSS compliance notes
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod                                   # 📝 UPDATED: Imports pkg/auth, platform-shared, contracts/events
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
│   │   │   └── repository.go                 # ConversationRepository interface (Create, FindByID, AddParticipant, RemoveParticipant)
│   │   │
│   │   ├── message/
│   │   │   ├── entity.go                     # Chat messages (ID, ConversationID, SenderID, Content, Type, SentAt, EditedAt, DeletedAt)
│   │   │   ├── attachment.go                 # Message attachments (MessageID, FileURL, FileType, FileName, FileSize, ThumbnailURL)
│   │   │   ├── read_receipt.go               # Read receipts (MessageID, UserID, ReadAt)
│   │   │   ├── reaction.go                   # Message reactions (MessageID, UserID, Emoji, ReactedAt)
│   │   │   ├── mention.go                    # User mentions (@username) (MessageID, MentionedUserID, Position)
│   │   │   ├── errors.go                     # Message errors (MessageNotFound, InvalidContent, MessageTooLong)
│   │   │   └── repository.go                 # MessageRepository interface (Create, Update, Delete, FindByConversation, MarkAsRead)
│   │   │
│   │   ├── notification/
│   │   │   ├── entity.go                     # All notifications (ID, UserID, Type, Title, Body, Data JSON, Priority, ReadAt, CreatedAt)
│   │   │   ├── enums.go                      # Type (JobAlert, ProposalUpdate, BidAlert, ContractUpdate, Payment, Review, Message, System), Priority (Low, Medium, High, Urgent), Category (Transactional, Marketing, System)
│   │   │   ├── preferences.go                # User notification preferences (Email, SMS, Push, InApp enabled/disabled per notification type)
│   │   │   ├── settings.go                   # Notification settings per type (instant delivery, daily digest, weekly summary, muted)
│   │   │   ├── errors.go                     # Notification errors (NotificationNotFound, InvalidType, PreferencesNotSet)
│   │   │   └── repository.go                 # NotificationRepository interface (Create, MarkAsRead, Delete, FindByUser, GetUnreadCount)
│   │   │
│   │   ├── in_app_notification/
│   │   │   ├── entity.go                     # 🆕 In-app real-time notifications (ID, NotificationID, UserID, DisplayedAt, DismissedAt, ClickedAt)
│   │   │   ├── badge_count.go                # 🆕 Unread notification badges (UserID, Category, UnreadCount, LastUpdatedAt)
│   │   │   ├── group.go                      # 🆕 Notification grouping (GroupID, NotificationIDs, GroupType, Summary)
│   │   │   ├── action.go                     # 🆕 Notification actions/CTAs (NotificationID, ActionType, ActionURL, Label)
│   │   │   ├── errors.go                     # 🆕 In-app notification errors
│   │   │   └── repository.go                 # 🆕 InAppNotificationRepository interface
│   │   │
│   │   ├── notification_template/
│   │   │   ├── entity.go                     # Notification templates (ID, Name, Type, Subject, Body, Variables, IsActive)
│   │   │   ├── variable.go                   # Template variables ({{user_name}}, {{job_title}}, {{amount}}, {{date}})
│   │   │   ├── localization.go               # Multi-language templates (TemplateID, Language, TranslatedSubject, TranslatedBody)
│   │   │   ├── errors.go                     # Template errors (TemplateNotFound, InvalidVariable, MissingTranslation)
│   │   │   └── repository.go                 # TemplateRepository interface (FindByType, Render, GetTranslation)
│   │   │
│   │   ├── email/
│   │   │   ├── entity.go                     # Email records (ID, RecipientID, RecipientEmail, Subject, Body, TemplateID, Status, SentAt, DeliveredAt)
│   │   │   ├── template.go                   # Email templates (HTML/Text templates with CSS inlining support)
│   │   │   ├── batch.go                      # Batch email sending (BatchID, Recipients, Subject, Body, Status, ScheduledFor)
│   │   │   ├── errors.go                     # Email errors (InvalidEmail, SendFailed, TemplateRenderFailed)
│   │   │   └── repository.go                 # EmailRepository interface (Create, FindByRecipient, GetDeliveryStatus)
│   │   │
│   │   ├── notification_queue/
│   │   │   ├── entity.go                     # 🆕 Queued notifications (ID, NotificationID, Priority, ScheduledFor, Status, Retries, CreatedAt)
│   │   │   ├── priority_queue.go             # 🆕 Priority-based queue (urgent notifications processed first)
│   │   │   ├── errors.go                     # 🆕 Queue errors
│   │   │   └── repository.go                 # 🆕 NotificationQueueRepository interface
│   │   │
│   │   ├── delivery_log/
│   │   │   ├── entity.go                     # 🆕 Delivery tracking (ID, NotificationID, Channel, Status, DeliveredAt, FailureReason, RetryCount)
│   │   │   ├── status.go                     # 🆕 Delivery status (Pending, Sent, Delivered, Failed, Bounced)
│   │   │   ├── errors.go                     # 🆕 Delivery errors
│   │   │   └── repository.go                 # 🆕 DeliveryLogRepository interface
│   │   │
│   │   ├── unsubscribe/
│   │   │   ├── entity.go                     # Unsubscribe management (UserID, NotificationType, UnsubscribedAt, Reason)
│   │   │   ├── errors.go                     # Unsubscribe errors
│   │   │   └── repository.go                 # UnsubscribeRepository interface
│   │   │
│   │   ├── online_status/
│   │   │   ├── entity.go                     # User online/offline status (UserID, Status, LastSeenAt, DeviceType)
│   │   │   ├── errors.go                     # Status errors
│   │   │   └── repository.go                 # OnlineStatusRepository interface
│   │   │
│   │   ├── message_flag/
│   │   │   ├── entity.go                     # Flagged messages (ID, MessageID, ReporterID, Reason, Status, ReviewedAt, ReviewedBy)
│   │   │   ├── reason.go                     # Flag reasons enum (Spam, Harassment, Inappropriate, Scam, Violence)
│   │   │   ├── status.go                     # Flag status (Pending, UnderReview, Resolved, Dismissed)
│   │   │   ├── errors.go                     # Flag errors
│   │   │   └── repository.go                 # MessageFlagRepository interface
│   │   │
│   │   └── outbox/                           # ❌ REMOVED from individual service (moved to platform-shared/outbox)
│   │       ├── entity.go                     # ❌ Use platform-shared/outbox/entity.go
│   │       └── repository.go                 # ❌ Use platform-shared/outbox/repository.go
│   │
│   ├── application/                          # 📋 Application Layer - Use cases & orchestration
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
│   │   ├── in_app_notification/
│   │   │   ├── service.go                    # 🆕 In-app notification logic
│   │   │   ├── real_time_sender.go           # 🆕 Real-time notification delivery via WebSocket/SSE
│   │   │   ├── badge_manager.go              # 🆕 Badge count management (calculate, update, reset counts)
│   │   │   ├── grouping_engine.go            # 🆕 Group similar notifications (5 new proposals → "You have 5 new proposals")
│   │   │   ├── dto.go                        # 🆕 InAppNotificationDTO
│   │   │   └── mapper.go                     # 🆕 In-app notification mappers
│   │   │
│   │   ├── email/
│   │   │   ├── service.go                    # Email business logic (Send, SendBatch, GetDeliveryStatus)
│   │   │   ├── sender.go                     # Email sending logic (via WildDuck SMTP)
│   │   │   ├── template_renderer.go          # Render email templates (inject variables, compile HTML)
│   │   │   ├── batch_sender.go               # Batch email sending (queue, rate limit, send in batches)
│   │   │   ├── dto.go                        # EmailDTO, SendEmailDTO, BatchEmailDTO
│   │   │   ├── mapper.go                     # Email mappers
│   │   │   └── wildduck_client.go            # 🆕 WildDuck SMTP integration (connect, authenticate, send)
│   │   │
│   │   ├── template/
│   │   │   ├── service.go                    # Template business logic (Create, Update, Render templates)
│   │   │   ├── renderer.go                   # Template rendering engine (replace variables, compile)
│   │   │   ├── variable_injector.go          # Inject dynamic variables into templates
│   │   │   ├── dto.go                        # TemplateDTO, RenderTemplateDTO
│   │   │   └── mapper.go                     # Template mappers
│   │   │
│   │   ├── online_status/
│   │   │   ├── service.go                    # Online status logic (Update status, GetStatus, GetOnlineUsers)
│   │   │   ├── tracker.go                    # Track online/offline status (heartbeat mechanism)
│   │   │   ├── presence_manager.go           # Manage user presence (online, away, busy, offline)
│   │   │   └── dto.go                        # OnlineStatusDTO
│   │   │
│   │   ├── flag/
│   │   │   ├── service.go                    # Flag business logic (FlagMessage, UnflagMessage, ReviewFlag)
│   │   │   ├── commands.go                   # FlagMessage, UnflagMessage, ResolveFlag
│   │   │   ├── dto.go                        # FlagDTO, FlagMessageDTO
│   │   │   └── mapper.go                     # Flag mappers
│   │   │
│   │   └── eventhandler/                     # 📝 UPDATED: Event handlers (now uses contracts/events for event types)
│   │       ├── user_handler.go               # Handle user events (user.created → send welcome notification, user.verified → send congrats)
│   │       ├── job_handler.go                # Handle job events (job.posted → notify matching freelancers, job.closed → notify applicants)
│   │       ├── proposal_handler.go           # Handle proposal events (proposal.submitted → notify client, proposal.accepted → notify freelancer)
│   │       ├── bid_handler.go                # Handle bid events (bid.placed → notify job owner, outbid.alert → notify freelancer)
│   │       ├── contract_handler.go           # Handle contract events (contract.created → notify both parties, milestone.completed → notify client)
│   │       ├── payment_handler.go            # Handle payment events (payment.processed → notify both parties, payout.processed → notify freelancer)
│   │       ├── review_handler.go             # Handle review events (review.submitted → notify reviewed user)
│   │       ├── message_handler.go            # Handle new message notifications (message.sent → notify recipient if offline)
│   │       ├── subscription_handler.go       # Handle subscription events (subscription.expiring → remind user)
│   │       ├── system_handler.go             # Handle system events (maintenance scheduled → notify all users)
│   │       └── admin_handler.go              # Handle admin actions on messages (message.removed → notify sender)
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
│   │   │       └── message_flag_repository.go # MessageFlagRepository implementation
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
│   │       │   ├── flag_handler.go           # Flag HTTP handlers (POST /messages/:id/flag, GET /flags)
│   │       │   └── health_handler.go         # Health check endpoints (/health, /ready, /live)
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
├── config/                                   # 🆕 MEDIUM - Configuration files
│   ├── default.yaml                          # 🆕 Default configuration
│   ├── dev.yaml                              # 🆕 Development overrides
│   └── prod.yaml                             # 🆕 Production overrides
│
├── dapr/                                     # 🆕 MEDIUM - Dapr components split by environment
│   ├── local/                                # 🆕 For dapr run
│   │   ├── pubsub.yaml                       # Kafka pub/sub component
│   │   └── statestore.yaml                   # State store component
│   └── k8s/                                  # 🆕 For kubectl apply
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
│       ├── deployment.yaml                   # Kubernetes Deployment
│       ├── service.yaml                      # Kubernetes Service
│       ├── configmap.yaml                    # ConfigMap
│       ├── secrets.yaml                      # Secrets
│       ├── hpa.yaml                          # HPA
│       ├── pdb.yaml                          # PDB
│       └── servicemonitor.yaml               # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                        # Setup local environment
│   ├── get-secrets.sh                        # Fetch secrets
│   ├── seed-templates.sh                     # Seed notification templates
│   └── seed-data.sh                          # Seed test data
│
├── tests/
│   ├── unit/                                 # Unit tests
│   ├── integration/                          # Integration tests
│   └── e2e/                                  # End-to-end tests
│
├── docs/
│   ├── README.md                             # Service overview
│   ├── API.md                                # API documentation
│   ├── EVENTS.md                             # 🆕 Events (published: message.sent, notification.delivered, consumed: all domain events)
│   ├── ARCHITECTURE.md                       # Architecture details
│   ├── MIGRATIONS.md                         # 🆕 Migration history
│   ├── SCHEMA.md                             # 🆕 Database schema
│   ├── RUNBOOK.md                            # 🆕 Operational procedures
│   ├── websocket-protocol.md                 # 🆕 WebSocket protocol documentation
│   ├── notification-system.md                # 🆕 Notification system overview
│   ├── in-app-notifications.md               # 🆕 In-app notifications guide
│   ├── email-templates.md                    # Email template guide
│   └── wildduck-integration.md               # 🆕 WildDuck integration guide
│
├── .github/
│   └── workflows/
│       ├── ci.yml                            # CI workflow
│       └── cd.yml                            # CD workflow
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
│   │   │   ├── entity.go              # File metadata
│   │   │   ├── enums.go               # FileType, Status, Visibility
│   │   │   ├── metadata.go            # File metadata
│   │   │   └── repository.go
│   │   ├── folder/
│   │   │   ├── entity.go              # Folder structure
│   │   │   └── repository.go
│   │   ├── upload/
│   │   │   ├── entity.go              # Upload sessions
│   │   │   ├── chunk.go               # Chunked uploads
│   │   │   ├── resumable.go           # Resumable upload support
│   │   │   └── repository.go
│   │   ├── media/
│   │   │   ├── entity.go              # Media processing records
│   │   │   ├── thumbnail.go           # Thumbnail generation
│   │   │   ├── variant.go             # Image variants
│   │   │   └── repository.go
│   │   ├── access_control/
│   │   │   ├── entity.go              # File access permissions
│   │   │   └── repository.go
│   │   ├── version/
│   │   │   ├── entity.go              # File versioning
│   │   │   └── repository.go
│   │   ├── share/
│   │   │   ├── entity.go              # File sharing
│   │   │   ├── link.go                # Share links
│   │   │   └── repository.go
│   │   ├── file_flag/
│   │   │   ├── entity.go              # Flagged files
│   │   │   ├── reason.go              # Flag reasons
│   │   │   ├── status.go              # Flag status
│   │   │   └── repository.go
│   │   └── outbox/
│   │       ├── entity.go              # ❌ REMOVED (use platform-shared/outbox/entity.go)
│   │       └── repository.go          # ❌ REMOVED (use platform-shared/outbox/repository.go)
│   │
│   ├── application/
│   │   ├── file/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Upload, Delete, Move, Copy
│   │   │   ├── queries.go             # Get, List, Search
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── upload/
│   │   │   ├── service.go
│   │   │   ├── chunked_upload.go
│   │   │   ├── resumable.go
│   │   │   └── dto.go
│   │   ├── media/
│   │   │   ├── service.go
│   │   │   ├── image_processor.go
│   │   │   ├── video_processor.go
│   │   │   ├── thumbnail_generator.go
│   │   │   └── dto.go
│   │   ├── folder/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── share/
│   │   │   ├── service.go
│   │   │   ├── link_generator.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── version/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── flag/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Flag, Unflag file
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   └── eventhandler/
│   │       ├── user_handler.go        # 📝 UPDATED: Handle user events (now uses contracts/events for event types)
│   │       ├── contract_handler.go    # 📝 UPDATED: Handle contract events (now uses contracts/events)
│   │       ├── portfolio_handler.go   # 📝 UPDATED: Handle portfolio events (now uses contracts/events)
│   │       └── admin_handler.go       # 📝 UPDATED: Handle admin actions on files (now uses contracts/events)
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go      # PostgreSQL connection setup (DSN from config, connection pooling)
│   │   │       ├── transaction.go     # Transaction helpers (Begin, Commit, Rollback, WithTransaction wrapper)
│   │   │       ├── migrations.go      # 📝 UPDATED: Auto-migrate logic (now with version tracking, GORM AutoMigrate for all tables)
│   │   │       ├── version.go         # 🆕 Schema version tracking (SchemaVersion table, RecordMigration function)
│   │   │       ├── safety.go          # 🆕 Pre-migration safety checks (environment validation, disk space check)
│   │   │       ├── file_repository.go # FileRepository implementation with GORM
│   │   │       ├── folder_repository.go
│   │   │       ├── upload_repository.go
│   │   │       ├── media_repository.go
│   │   │       ├── access_control_repository.go
│   │   │       ├── version_repository.go
│   │   │       ├── share_repository.go
│   │   │       ├── file_flag_repository.go
│   │   │       └── outbox_repository.go # ❌ REMOVED (use platform-shared/outbox/postgres/repository.go)
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go      # Redis connection setup (connection pooling, retry logic)
│   │   │       └── file_cache.go      # File metadata caching (Get, Set, Invalidate with TTL)
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go        # 📝 UPDATED: Kafka consumer (now uses platform-shared/inbox for message deduplication)
│   │   │       ├── producer.go        # 📝 UPDATED: Kafka producer (now uses platform-shared/outbox for reliable publishing)
│   │   │       ├── topics.go          # 📝 UPDATED: Topic constants (now imported from contracts/events - file.uploaded, file.deleted, media.processed)
│   │   │       └── scram.go           # SCRAM authentication for Kafka (SASL/SCRAM-SHA-256)
│   │   ├── object_storage/
│   │   │   ├── local/
│   │   │   │   ├── storage.go         # Local file system storage
│   │   │   │   └── config.go          # Local storage config
│   │   │   ├── minio/
│   │   │   │   ├── client.go          # Self-hosted MinIO client (upload, download, presigned URLs)
│   │   │   │   └── config.go          # MinIO configuration (endpoint, credentials, bucket names)
│   │   │   └── provider.go            # Storage provider abstraction
│   │   ├── media_processing/
│   │   │   ├── image/
│   │   │   │   ├── resizer.go         # Image resizing logic
│   │   │   │   ├── optimizer.go       # Image optimization (compression)
│   │   │   │   └── watermark.go       # Watermarking images
│   │   │   ├── video/
│   │   │   │   ├── transcoder.go      # Video transcoding
│   │   │   │   └── thumbnail.go       # Video thumbnail generation
│   │   │   └── processor.go           # General media processor
│   │   ├── virus_scan/
│   │   │   └── clamav.go              # ClamAV integration for virus scanning
│   │   └── outbox/
│   │       ├── processor.go           # ❌ REMOVED (use platform-shared/outbox/forwarder.go)
│   │       └── scheduler.go           # ❌ REMOVED (use platform-shared/outbox/scheduler.go)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── file_handler.go    # File HTTP handlers (GET, POST, DELETE /files)
│   │       │   ├── upload_handler.go  # Upload handlers (POST /upload, chunked support)
│   │       │   ├── download_handler.go # Download handlers (GET /download/:id)
│   │       │   ├── folder_handler.go  # Folder handlers
│   │       │   ├── share_handler.go   # Share handlers
│   │       │   ├── media_handler.go   # Media processing handlers
│   │       │   ├── flag_handler.go    # Flag handlers
│   │       │   └── health_handler.go  # Health check endpoints (/health, /ready, /live)
│   │       ├── middleware/
│   │       │   ├── auth.go            # 📝 UPDATED: Authentication middleware (uses pkg/auth for JWT verification)
│   │       │   ├── rbac.go            # 📝 UPDATED: RBAC middleware (uses pkg/auth authorizer for role checks)
│   │       │   ├── cors.go            # 📝 UPDATED: CORS middleware (uses platform-shared/ginx CORS)
│   │       │   ├── rate_limit.go      # Rate limiting middleware (token bucket algorithm)
│   │       │   ├── logging.go         # 📝 UPDATED: Logging middleware (uses platform-shared/ginx/logging)
│   │       │   ├── error_handler.go   # Error handling middleware
│   │       │   ├── request_id.go      # 📝 UPDATED: Request ID middleware (uses platform-shared/ginx/requestid)
│   │       │   └── file_size_limit.go # File size limit middleware
│   │       ├── responses/
│   │       │   ├── success.go         # 📝 UPDATED: Success response wrappers (use platform-shared/httpx/response.go)
│   │       │   ├── error.go           # 📝 UPDATED: Error responses (use platform-shared/httpx/errors.go)
│   │       │   └── pagination.go      # 📝 UPDATED: Pagination (use platform-shared/httpx/pagination.go)
│   │       └── router.go              # 📝 UPDATED: HTTP router setup (uses Gin, applies platform-shared/ginx middleware)
│   │
│   └── config/                               # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                         # 🆕 Typed Config struct (App, Server, Postgres, Kafka, Redis, MinIO, Storage)
│       ├── loader.go                         # 🆕 Config loader using Viper (precedence: flags → env → file → defaults)
│       └── docs/
│           └── CONFIGURATION.md              # 🆕 Configuration documentation (all ENV vars, defaults, examples)
│
├── config/                                   # 🆕 MEDIUM - Configuration files
│   ├── default.yaml                          # 🆕 Default configuration
│   ├── dev.yaml                              # 🆕 Development overrides
│   └── prod.yaml                             # 🆕 Production overrides
│
├── dapr/                                     # 🆕 MEDIUM - Dapr components split by environment
│   ├── local/                                # 🆕 For dapr run
│   │   ├── pubsub.yaml                       # Kafka pub/sub component
│   │   └── statestore.yaml                   # State store component
│   └── k8s/                                  # 🆕 For kubectl apply
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
│       ├── deployment.yaml                   # Kubernetes Deployment
│       ├── service.yaml                      # Kubernetes Service
│       ├── configmap.yaml                    # ConfigMap
│       ├── secrets.yaml                      # Secrets
│       ├── hpa.yaml                          # HPA
│       ├── pdb.yaml                          # PDB
│       ├── pvc.yaml                          # Persistent Volume Claim
│       └── servicemonitor.yaml               # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                        # Setup local environment
│   ├── get-secrets.sh                        # Fetch secrets
│   ├── seed-data.sh                          # Seed test data
│   └── cleanup-orphaned.sh                   # Cleanup orphaned files
│
├── tests/
│   ├── unit/                                 # Unit tests
│   ├── integration/                          # Integration tests
│   └── e2e/                                  # End-to-end tests
│
├── docs/
│   ├── README.md                             # Service overview
│   ├── api.md                                # API documentation
│   ├── events.md                             # 📝 UPDATED: Events (published: file.uploaded, file.deleted, consumed: user.created, etc.)
│   ├── upload-flow.md                        # Upload flow documentation
│   ├── media-processing.md                   # Media processing guide
│   ├── MIGRATIONS.md                         # 🆕 Migration history
│   ├── SCHEMA.md                             # 🆕 Database schema
│   └── RUNBOOK.md                            # 🆕 Operational procedures
│
├── .github/
│   └── workflows/
│       ├── ci.yml                            # CI workflow
│       └── cd.yml                            # CD workflow
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
│   │   │   ├── entity.go              # Search index metadata
│   │   │   ├── job_index.go           # Job search document
│   │   │   ├── user_index.go          # User/Freelancer document
│   │   │   └── repository.go
│   │   ├── search_query/
│   │   │   ├── entity.go              # Search query logs
│   │   │   ├── filters.go             # Search filters
│   │   │   └── repository.go
│   │   ├── recommendation/
│   │   │   ├── entity.go              # Recommendation records
│   │   │   ├── score.go               # Scoring algorithm
│   │   │   ├── reason.go              # Recommendation reasons
│   │   │   ├── feedback.go            # User feedback on recommendations
│   │   │   └── repository.go
│   │   ├── recommendation_model/
│   │   │   ├── entity.go              # ML model metadata
│   │   │   ├── feature.go             # Feature vectors
│   │   │   ├── training_data.go       # Training data
│   │   │   └── repository.go
│   │   ├── user_preference/
│   │   │   ├── entity.go              # User preferences for recommendations
│   │   │   ├── implicit_signals.go    # Implicit user signals (clicks, views)
│   │   │   ├── explicit_preferences.go # Explicit preferences
│   │   │   └── repository.go
│   │   ├── matching/
│   │   │   ├── entity.go              # Job-Freelancer matches
│   │   │   ├── criteria.go            # Matching criteria
│   │   │   ├── score_breakdown.go     # Detailed match scoring
│   │   │   └── repository.go
│   │   ├── feed/
│   │   │   ├── entity.go              # User feeds
│   │   │   ├── item.go                # Feed items
│   │   │   ├── personalization.go     # Personalization data
│   │   │   └── repository.go
│   │   ├── trending/
│   │   │   ├── entity.go              # Trending jobs/skills
│   │   │   ├── calculator.go          # Calculate trending items
│   │   │   └── repository.go
│   │   ├── saved_search/
│   │   │   ├── entity.go              # Saved searches
│   │   │   ├── alert.go               # Search alerts
│   │   │   └── repository.go
│   │   ├── similarity/
│   │   │   ├── entity.go              # Similar jobs/users
│   │   │   ├── vector.go              # Similarity vectors
│   │   │   └── repository.go
│   │   └── outbox/
│   │       ├── entity.go              # ❌ REMOVED (use platform-shared/outbox/entity.go)
│   │       └── repository.go          # ❌ REMOVED (use platform-shared/outbox/repository.go)
│   │
│   ├── application/
│   │   ├── search/
│   │   │   ├── service.go
│   │   │   ├── job_search.go
│   │   │   ├── freelancer_search.go
│   │   │   ├── query_builder.go
│   │   │   ├── facet_builder.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── indexing/
│   │   │   ├── service.go
│   │   │   ├── job_indexer.go
│   │   │   ├── user_indexer.go
│   │   │   ├── bulk_indexer.go
│   │   │   └── dto.go
│   │   ├── recommendation/
│   │   │   ├── service.go
│   │   │   ├── job_recommender.go     # Recommend jobs to freelancers
│   │   │   ├── freelancer_recommender.go # Recommend freelancers to clients
│   │   │   ├── collaborative_filtering.go # Collaborative filtering algorithm
│   │   │   ├── content_based.go       # Content-based filtering
│   │   │   ├── hybrid_recommender.go  # Hybrid recommendation approach
│   │   │   ├── scoring_engine.go      # Calculate recommendation scores
│   │   │   ├── personalization.go     # Personalization logic
│   │   │   ├── diversity_optimizer.go # Ensure diverse recommendations
│   │   │   ├── cold_start_handler.go  # Handle new users/jobs
│   │   │   ├── dto.go
│   │   │   └── ml_model.go            # ML model integration
│   │   ├── matching/
│   │   │   ├── service.go
│   │   │   ├── matcher.go             # Job-Freelancer matching
│   │   │   ├── criteria_evaluator.go  # Evaluate match criteria
│   │   │   ├── skill_matcher.go       # Match based on skills
│   │   │   ├── experience_matcher.go  # Match based on experience
│   │   │   ├── rate_matcher.go        # Match based on rates
│   │   │   ├── availability_matcher.go # Match based on availability
│   │   │   ├── score_calculator.go    # Calculate match score
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── feed/
│   │   │   ├── service.go
│   │   │   ├── generator.go           # Generate personalized feed
│   │   │   ├── ranking.go             # Rank feed items
│   │   │   ├── freshness_scorer.go    # Score by freshness
│   │   │   ├── relevance_scorer.go    # Score by relevance
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── trending/
│   │   │   ├── service.go
│   │   │   ├── calculator.go          # Calculate trending items
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── similarity/
│   │   │   ├── service.go
│   │   │   ├── job_similarity.go      # Find similar jobs
│   │   │   ├── user_similarity.go     # Find similar users
│   │   │   ├── vector_calculator.go   # Calculate similarity vectors
│   │   │   └── dto.go
│   │   ├── suggestion/
│   │   │   ├── service.go             # Autocomplete suggestions
│   │   │   ├── dto.go
│   │   │   └── cache_warmer.go
│   │   └── eventhandler/
│   │       ├── job_handler.go         # 📝 UPDATED: Handle job events (now uses contracts/events for event types)
│   │       ├── user_handler.go        # 📝 UPDATED: Handle user events (now uses contracts/events)
│   │       ├── proposal_handler.go    # 📝 UPDATED: Handle proposal events (now uses contracts/events)
│   │       ├── contract_handler.go    # 📝 UPDATED: Handle contract events (now uses contracts/events)
│   │       ├── review_handler.go      # 📝 UPDATED: Handle review events (now uses contracts/events)
│   │       ├── skill_handler.go       # 📝 UPDATED: Handle skill events (now uses contracts/events)
│   │       ├── interaction_handler.go # 📝 UPDATED: Track user interactions (now uses contracts/events)
│   │       └── admin_handler.go       # 📝 UPDATED: Handle content removal from index (now uses contracts/events)
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go      # PostgreSQL connection setup (DSN from config, connection pooling)
│   │   │       ├── transaction.go     # Transaction helpers (Begin, Commit, Rollback, WithTransaction wrapper)
│   │   │       ├── migrations.go      # 📝 UPDATED: Auto-migrate logic (now with version tracking, GORM AutoMigrate for all tables)
│   │   │       ├── version.go         # 🆕 Schema version tracking (SchemaVersion table, RecordMigration function)
│   │   │       ├── safety.go          # 🆕 Pre-migration safety checks (environment validation, disk space check)
│   │   │       ├── search_query_repository.go
│   │   │       ├── recommendation_repository.go
│   │   │       ├── recommendation_model_repository.go
│   │   │       ├── user_preference_repository.go
│   │   │       ├── matching_repository.go
│   │   │       ├── feed_repository.go
│   │   │       ├── trending_repository.go
│   │   │       ├── saved_search_repository.go
│   │   │       ├── similarity_repository.go
│   │   │       └── outbox_repository.go # ❌ REMOVED (use platform-shared/outbox/postgres/repository.go)
│   │   ├── elasticsearch/
│   │   │   ├── client.go              # Elasticsearch client
│   │   │   ├── index_manager.go       # Index management
│   │   │   ├── job_mapper.go          # Job mapping to ES document
│   │   │   ├── user_mapper.go         # User mapping to ES document
│   │   │   └── config.go              # ES configuration (hosts, auth)
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go      # Redis connection setup (connection pooling, retry logic)
│   │   │       ├── search_cache.go    # Cache search results (Get, Set, Invalidate with TTL)
│   │   │       ├── suggestion_cache.go # Cache autocomplete suggestions
│   │   │       ├── feed_cache.go      # Cache user feeds
│   │   │       ├── recommendation_cache.go # Cache recommendations
│   │   │       └── trending_cache.go  # Cache trending items
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go        # 📝 UPDATED: Kafka consumer (now uses platform-shared/inbox for message deduplication)
│   │   │       ├── producer.go        # 📝 UPDATED: Kafka producer (now uses platform-shared/outbox for reliable publishing)
│   │   │       ├── topics.go          # 📝 UPDATED: Topic constants (now imported from contracts/events - job.indexed, user.indexed, recommendation.generated)
│   │   │       └── scram.go           # SCRAM authentication for Kafka (SASL/SCRAM-SHA-256)
│   │   ├── ml/
│   │   │   ├── model_loader.go        # Load ML models
│   │   │   ├── predictor.go           # Make predictions
│   │   │   ├── feature_extractor.go   # Extract features
│   │   │   ├── trainer.go             # Train models
│   │   │   └── evaluator.go           # Evaluate model performance
│   │   └── outbox/
│   │       ├── processor.go           # ❌ REMOVED (use platform-shared/outbox/forwarder.go)
│   │       └── scheduler.go           # ❌ REMOVED (use platform-shared/outbox/scheduler.go)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── search_handler.go  # Search HTTP handlers (POST /search/jobs, /search/freelancers)
│   │       │   ├── recommendation_handler.go # Recommendation handlers (GET /recommendations/jobs, /recommendations/freelancers)
│   │       │   ├── matching_handler.go # Matching handlers
│   │       │   ├── feed_handler.go    # Feed handlers
│   │       │   ├── trending_handler.go # Trending handlers
│   │       │   ├── similarity_handler.go # Similarity handlers
│   │       │   ├── suggestion_handler.go # Autocomplete handlers
│   │       │   ├── indexing_handler.go # Admin endpoint for indexing
│   │       │   └── health_handler.go  # Health check endpoints (/health, /ready, /live)
│   │       ├── middleware/
│   │       │   ├── auth.go            # 📝 UPDATED: Authentication middleware (uses pkg/auth for JWT verification)
│   │       │   ├── rbac.go            # 📝 UPDATED: RBAC middleware (uses pkg/auth authorizer for role checks)
│   │       │   ├── cors.go            # 📝 UPDATED: CORS middleware (uses platform-shared/ginx CORS)
│   │       │   ├── rate_limit.go      # Rate limiting middleware (token bucket algorithm)
│   │       │   ├── logging.go         # 📝 UPDATED: Logging middleware (structured logs per request, uses platform-shared/ginx/logging)
│   │       │   ├── error_handler.go   # Error handling middleware
│   │       │   └── request_id.go      # 📝 UPDATED: Request ID middleware (X-Request-ID header, uses platform-shared/ginx/requestid)
│   │       ├── responses/
│   │       │   ├── success.go         # 📝 UPDATED: Success response wrappers (use platform-shared/httpx/response.go)
│   │       │   ├── error.go           # 📝 UPDATED: Error responses (use platform-shared/httpx/errors.go)
│   │       │   └── pagination.go      # 📝 UPDATED: Pagination (use platform-shared/httpx/pagination.go)
│   │       └── router.go              # 📝 UPDATED: HTTP router setup (uses Gin, applies platform-shared/ginx middleware)
│   │
│   └── config/                               # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                         # 🆕 Typed Config struct (App, Server, Postgres, Kafka, Redis, Elasticsearch, ML)
│       ├── loader.go                         # 🆕 Config loader using Viper (precedence: flags → env → file → defaults)
│       └── docs/
│           └── CONFIGURATION.md              # 🆕 Configuration documentation (all ENV vars, defaults, examples)
│
├── config/                                   # 🆕 MEDIUM - Configuration files
│   ├── default.yaml                          # 🆕 Default configuration
│   ├── dev.yaml                              # 🆕 Development overrides
│   └── prod.yaml                             # 🆕 Production overrides
│
├── dapr/                                     # 🆕 MEDIUM - Dapr components split by environment
│   ├── local/                                # 🆕 For dapr run
│   │   ├── pubsub.yaml                       # Kafka pub/sub component
│   │   └── statestore.yaml                   # State store component
│   └── k8s/                                  # 🆕 For kubectl apply
│       ├── pubsub.yaml                       # Kafka with scopes: ["search-be"]
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
│   │   ├── text_analyzer.go                  # Text analysis utilities
│   │   ├── normalizer.go                     # Data normalization
│   │   └── vector_math.go                    # Vector math for similarity
│   └── constants/
│       ├── events.go                         # ❌ REMOVED: Use contracts/events
│       ├── topics.go                         # ❌ REMOVED: Use contracts/events
│       └── indices.go                        # Index names constants
│
├── elasticsearch/
│   ├── mappings/
│   │   ├── jobs.json                         # Job index mapping
│   │   └── users.json                        # User index mapping
│   └── analyzers/
│       └── custom_analyzers.json             # Custom analyzers
│
├── ml_models/
│   ├── job_recommendation/
│   │   ├── model.pkl                         # ML model for job recommendations
│   │   ├── features.json                     # Features list
│   │   └── metadata.json                     # Model metadata
│   ├── freelancer_recommendation/
│   │   ├── model.pkl                         # ML model for freelancer recommendations
│   │   ├── features.json                     # Features list
│   │   └── metadata.json                     # Model metadata
│   └── matching/
│       ├── model.pkl                         # ML model for matching
│       ├── features.json                     # Features list
│       └── metadata.json                     # Model metadata
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                   # Kubernetes Deployment
│       ├── service.yaml                      # Kubernetes Service
│       ├── configmap.yaml                    # ConfigMap
│       ├── secrets.yaml                      # Secrets
│       ├── hpa.yaml                          # HPA
│       ├── pdb.yaml                          # PDB
│       └── servicemonitor.yaml               # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                        # Setup local environment
│   ├── get-secrets.sh                        # Fetch secrets
│   ├── seed-data.sh                          # Seed test data
│   ├── create-indices.sh                     # Create ES indices
│   ├── reindex-all.sh                        # Reindex all data
│   └── train-models.sh                       # Train ML models
│
├── tests/
│   ├── unit/                                 # Unit tests
│   ├── integration/                          # Integration tests
│   └── e2e/                                  # End-to-end tests
│
├── docs/
│   ├── README.md                             # Service overview
│   ├── api.md                                # API documentation
│   ├── events.md                             # 📝 UPDATED: Events (published: job.indexed, user.indexed, recommendation.generated, consumed: job.posted, user.updated, etc.)
│   ├── search-algorithms.md                  # Search algorithms
│   ├── recommendation-engine.md              # Recommendation engine details
│   ├── recommendation-types.md               # Recommendation types
│   ├── matching-algorithm.md                 # Matching algorithm
│   ├── ml-models.md                          # ML models documentation
│   ├── elasticsearch-setup.md                # Elasticsearch setup
│   ├── MIGRATIONS.md                         # 🆕 Migration history
│   ├── SCHEMA.md                             # 🆕 Database schema
│   └── RUNBOOK.md                            # 🆕 Operational procedures
│
├── .github/
│   └── workflows/
│       ├── ci.yml                            # CI workflow
│       └── cd.yml                            # CD workflow
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
│   │   │   ├── entity.go              # Review aggregate root
│   │   │   ├── rating.go              # Rating breakdown (overall, quality, communication, etc.)
│   │   │   ├── enums.go               # ReviewType, Status, ReviewCategory
│   │   │   ├── response.go            # Review responses
│   │   │   ├── helpful.go             # Helpful votes
│   │   │   └── repository.go
│   │   ├── rating/
│   │   │   ├── entity.go              # Rating system
│   │   │   ├── criteria.go            # Rating criteria
│   │   │   ├── aggregation.go         # Aggregate ratings
│   │   │   └── repository.go
│   │   ├── badge/
│   │   │   ├── entity.go              # Achievement badges
│   │   │   ├── criteria.go            # Badge criteria
│   │   │   ├── types.go               # Badge types (Top Rated, Rising Talent, etc.)
│   │   │   ├── level.go               # Badge levels
│   │   │   └── repository.go
│   │   ├── user_badge/
│   │   │   ├── entity.go              # User badge assignments
│   │   │   ├── achievement_date.go    # When badge was earned
│   │   │   └── repository.go
│   │   ├── reputation/
│   │   │   ├── entity.go              # Reputation scores
│   │   │   ├── score_calculator.go    # Score calculation
│   │   │   ├── history.go             # Reputation history
│   │   │   └── repository.go
│   │   ├── feedback/
│   │   │   ├── entity.go              # Private feedback
│   │   │   ├── category.go            # Feedback categories
│   │   │   └── repository.go
│   │   ├── flag/
│   │   │   ├── entity.go              # Flagged reviews
│   │   │   ├── reason.go              # Flag reasons
│   │   │   └── repository.go
│   │   ├── stats/
│   │   │   ├── entity.go              # Review statistics
│   │   │   ├── aggregates.go          # Aggregated stats
│   │   │   └── repository.go
│   │   ├── review_reminder/
│   │   │   ├── entity.go              # Review reminders
│   │   │   └── repository.go
│   │   └── outbox/
│   │       ├── entity.go              # ❌ REMOVED (use platform-shared/outbox/entity.go)
│   │       └── repository.go          # ❌ REMOVED (use platform-shared/outbox/repository.go)
│   │
│   ├── application/
│   │   ├── review/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Update, Delete, Respond
│   │   │   ├── queries.go             # Get, List, Filter
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── business_rules.go
│   │   ├── rating/
│   │   │   ├── service.go
│   │   │   ├── aggregator.go          # Aggregate ratings
│   │   │   ├── calculator.go          # Calculate average ratings
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── badge/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Award, Revoke badges
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── achievement_checker.go # Check badge criteria
│   │   ├── reputation/
│   │   │   ├── service.go
│   │   │   ├── calculator.go          # Calculate reputation
│   │   │   ├── updater.go             # Update reputation
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── feedback/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── stats/
│   │   │   ├── service.go
│   │   │   ├── aggregator.go          # Aggregate stats
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   └── eventhandler/
│   │       ├── contract_handler.go    # 📝 UPDATED: Enable review after contract (now uses contracts/events)
│   │       ├── user_handler.go        # 📝 UPDATED: Handle user events (now uses contracts/events)
│   │       ├── job_handler.go         # 📝 UPDATED: Update job stats (now uses contracts/events)
│   │       ├── proposal_handler.go    # 📝 UPDATED: Update proposal stats (now uses contracts/events)
│   │       ├── payment_handler.go     # 📝 UPDATED: Ensure payment before review (now uses contracts/events)
│   │       └── admin_handler.go       # 📝 UPDATED: Handle admin actions on reviews (now uses contracts/events)
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go      # PostgreSQL connection setup (DSN from config, connection pooling)
│   │   │       ├── transaction.go     # Transaction helpers (Begin, Commit, Rollback, WithTransaction wrapper)
│   │   │       ├── migrations.go      # 📝 UPDATED: Auto-migrate logic (now with version tracking, GORM AutoMigrate for all tables)
│   │   │       ├── version.go         # 🆕 Schema version tracking (SchemaVersion table, RecordMigration function)
│   │   │       ├── safety.go          # 🆕 Pre-migration safety checks (environment validation, disk space check)
│   │   │       ├── review_repository.go
│   │   │       ├── rating_repository.go
│   │   │       ├── badge_repository.go
│   │   │       ├── user_badge_repository.go
│   │   │       ├── reputation_repository.go
│   │   │       ├── feedback_repository.go
│   │   │       ├── flag_repository.go
│   │   │       ├── stats_repository.go
│   │   │       ├── review_reminder_repository.go
│   │   │       └── outbox_repository.go # ❌ REMOVED (use platform-shared/outbox/postgres/repository.go)
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go      # Redis connection setup (connection pooling, retry logic)
│   │   │       ├── review_cache.go    # Cache reviews (Get, Set, Invalidate with TTL)
│   │   │       ├── badge_cache.go     # Cache badges
│   │   │       ├── reputation_cache.go # Cache reputation scores
│   │   │       └── stats_cache.go     # Cache stats
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go        # 📝 UPDATED: Kafka consumer (now uses platform-shared/inbox for message deduplication)
│   │   │       ├── producer.go        # 📝 UPDATED: Kafka producer (now uses platform-shared/outbox for reliable publishing)
│   │   │       ├── topics.go          # 📝 UPDATED: Topic constants (now imported from contracts/events - review.submitted, review.responded, badge.awarded)
│   │   │       └── scram.go           # SCRAM authentication for Kafka (SASL/SCRAM-SHA-256)
│   │   └── outbox/
│   │       ├── processor.go           # ❌ REMOVED (use platform-shared/outbox/forwarder.go)
│   │       └── scheduler.go           # ❌ REMOVED (use platform-shared/outbox/scheduler.go)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── review_handler.go  # Review HTTP handlers (POST /reviews, GET /reviews/:id)
│   │       │   ├── rating_handler.go  # Rating handlers
│   │       │   ├── badge_handler.go   # Badge handlers
│   │       │   ├── reputation_handler.go # Reputation handlers
│   │       │   ├── feedback_handler.go # Feedback handlers
│   │       │   ├── stats_handler.go   # Stats handlers
│   │       │   ├── flag_handler.go    # Flag handlers
│   │       │   └── health_handler.go  # Health check endpoints (/health, /ready, /live)
│   │       ├── middleware/
│   │       │   ├── auth.go            # 📝 UPDATED: Authentication middleware (uses pkg/auth for JWT verification)
│   │       │   ├── rbac.go            # 📝 UPDATED: RBAC middleware (uses pkg/auth authorizer for role checks)
│   │       │   ├── cors.go            # 📝 UPDATED: CORS middleware (uses platform-shared/ginx CORS)
│   │       │   ├── rate_limit.go      # Rate limiting middleware (token bucket algorithm)
│   │       │   ├── logging.go         # 📝 UPDATED: Logging middleware (structured logs per request, uses platform-shared/ginx/logging)
│   │       │   ├── error_handler.go   # Error handling middleware
│   │       │   └── request_id.go      # 📝 UPDATED: Request ID middleware (X-Request-ID header, uses platform-shared/ginx/requestid)
│   │       ├── responses/
│   │       │   ├── success.go         # 📝 UPDATED: Success response wrappers (use platform-shared/httpx/response.go)
│   │       │   ├── error.go           # 📝 UPDATED: Error responses (use platform-shared/httpx/errors.go)
│   │       │   └── pagination.go      # 📝 UPDATED: Pagination (use platform-shared/httpx/pagination.go)
│   │       └── router.go              # 📝 UPDATED: HTTP router setup (uses Gin, applies platform-shared/ginx middleware)
│   │
│   └── config/                               # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                         # 🆕 Typed Config struct (App, Server, Postgres, Kafka, Redis)
│       ├── loader.go                         # 🆕 Config loader using Viper (precedence: flags → env → file → defaults)
│       └── docs/
│           └── CONFIGURATION.md              # 🆕 Configuration documentation (all ENV vars, defaults, examples)
│
├── config/                                   # 🆕 MEDIUM - Configuration files
│   ├── default.yaml                          # 🆕 Default configuration
│   ├── dev.yaml                              # 🆕 Development overrides
│   └── prod.yaml                             # 🆕 Production overrides
│
├── dapr/                                     # 🆕 MEDIUM - Dapr components split by environment
│   ├── local/                                # 🆕 For dapr run
│   │   ├── pubsub.yaml                       # Kafka pub/sub component
│   │   └── statestore.yaml                   # State store component
│   └── k8s/                                  # 🆕 For kubectl apply
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
│       ├── deployment.yaml                   # Kubernetes Deployment
│       ├── service.yaml                      # Kubernetes Service
│       ├── configmap.yaml                    # ConfigMap
│       ├── secrets.yaml                      # Secrets
│       ├── hpa.yaml                          # HPA
│       ├── pdb.yaml                          # PDB
│       └── servicemonitor.yaml               # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                        # Setup local environment
│   ├── get-secrets.sh                        # Fetch secrets
│   ├── seed-badges.sh                        # Seed badges
│   └── recalculate-reputation.sh             # Recalculate reputation
│
├── tests/
│   ├── unit/                                 # Unit tests
│   ├── integration/                          # Integration tests
│   └── e2e/                                  # End-to-end tests
│
├── docs/
│   ├── README.md                             # Service overview
│   ├── api.md                                # API documentation
│   ├── events.md                             # 📝 UPDATED: Events (published: review.submitted, badge.awarded, reputation.updated, consumed: contract.ended, etc.)
│   ├── badge-system.md                       # Badge system
│   ├── reputation-algorithm.md               # Reputation algorithm
│   ├── rating-system.md                      # Rating system
│   ├── review-guidelines.md                  # Review guidelines
│   ├── MIGRATIONS.md                         # 🆕 Migration history
│   ├── SCHEMA.md                             # 🆕 Database schema
│   └── RUNBOOK.md                            # 🆕 Operational procedures
│
├── .github/
│   └── workflows/
│       ├── ci.yml                            # CI workflow
│       └── cd.yml                            # CD workflow
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

## **📦🔟 subscriptions-be (UPDATED)**

```
apps/be/subscriptions-be/
│
├── cmd/
│   └── api/
│       └── main.go                           # 📝 UPDATED: Application entry point - initializes Gin, Dapr, connects to Postgres (now uses loadConfig from internal/config, platform-shared/logging)
│
├── internal/
│   ├── domain/
│   │   ├── plan/
│   │   │   ├── entity.go              # Subscription plans
│   │   │   ├── features.go            # Plan features
│   │   │   ├── pricing.go             # Pricing tiers
│   │   │   ├── limits.go              # Plan limits
│   │   │   └── repository.go
│   │   ├── subscription/
│   │   │   ├── entity.go              # User subscriptions
│   │   │   ├── billing_cycle.go       # Billing cycle management
│   │   │   ├── enums.go               # Status, Type
│   │   │   ├── auto_renewal.go        # Auto-renewal logic
│   │   │   └── repository.go
│   │   ├── connect/
│   │   │   ├── entity.go              # Connects/Credits
│   │   │   ├── package.go             # Connect packages
│   │   │   ├── transaction.go         # Connect transactions
│   │   │   ├── balance.go             # Balance tracking
│   │   │   └── repository.go
│   │   ├── usage/
│   │   │   ├── entity.go              # Usage tracking
│   │   │   ├── quota.go               # Usage quotas
│   │   │   ├── limit.go               # Usage limits
│   │   │   └── repository.go
│   │   ├── addon/
│   │   │   ├── entity.go              # Plan add-ons
│   │   │   └── repository.go
│   │   ├── promotion/
│   │   │   ├── entity.go              # Promotional codes
│   │   │   ├── discount.go            # Discount rules
│   │   │   ├── usage_limit.go         # Usage limits for promos
│   │   │   └── repository.go
│   │   ├── trial/
│   │   │   ├── entity.go              # Free trials
│   │   │   ├── eligibility.go         # Trial eligibility
│   │   │   └── repository.go
│   │   ├── billing_history/
│   │   │   ├── entity.go              # Billing history
│   │   │   └── repository.go
│   │   ├── feature_toggle/
│   │   │   ├── entity.go              # Feature toggles per plan
│   │   │   └── repository.go
│   │   └── outbox/
│   │       ├── entity.go              # ❌ REMOVED (use platform-shared/outbox/entity.go)
│   │       └── repository.go          # ❌ REMOVED (use platform-shared/outbox/repository.go)
│   │
│   ├── application/
│   │   ├── plan/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Update plans
│   │   │   ├── queries.go             # Get, List plans
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── subscription/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Subscribe, Cancel, Change, Pause
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── lifecycle_manager.go   # Manage subscription lifecycle
│   │   │   └── renewal_manager.go     # Handle renewals
│   │   ├── connect/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Purchase, Use, Transfer, Refund
│   │   │   ├── queries.go             # Get balance, History
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── calculator.go          # Calculate connect costs
│   │   ├── usage/
│   │   │   ├── service.go
│   │   │   ├── tracker.go             # Track usage
│   │   │   ├── quota_checker.go       # Check quotas
│   │   │   ├── limiter.go             # Enforce limits
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── addon/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── promotion/
│   │   │   ├── service.go
│   │   │   ├── validator.go           # Validate promo codes
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── trial/
│   │   │   ├── service.go
│   │   │   ├── eligibility_checker.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── billing/
│   │   │   ├── service.go
│   │   │   ├── invoice_generator.go   # Generate invoices
│   │   │   ├── payment_processor.go   # Process payments
│   │   │   └── dto.go
│   │   ├── feature_toggle/
│   │   │   ├── service.go
│   │   │   ├── checker.go             # Check feature access
│   │   │   └── dto.go
│   │   └── eventhandler/
│   │       ├── user_handler.go        # 📝 UPDATED: Handle new user events (now uses contracts/events)
│   │       ├── payment_handler.go     # 📝 UPDATED: Handle payment events (now uses contracts/events)
│   │       ├── proposal_handler.go    # 📝 UPDATED: Deduct connects on proposal (now uses contracts/events)
│   │       ├── job_handler.go         # 📝 UPDATED: Check posting limits (now uses contracts/events)
│   │       └── admin_handler.go       # 📝 UPDATED: Handle admin subscription actions (now uses contracts/events)
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go      # PostgreSQL connection setup (DSN from config, connection pooling)
│   │   │       ├── transaction.go     # Transaction helpers (Begin, Commit, Rollback, WithTransaction wrapper)
│   │   │       ├── migrations.go      # 📝 UPDATED: Auto-migrate logic (now with version tracking, GORM AutoMigrate for all tables)
│   │   │       ├── version.go         # 🆕 Schema version tracking (SchemaVersion table, RecordMigration function)
│   │   │       ├── safety.go          # 🆕 Pre-migration safety checks (environment validation, disk space check)
│   │   │       ├── plan_repository.go
│   │   │       ├── subscription_repository.go
│   │   │       ├── connect_repository.go
│   │   │       ├── usage_repository.go
│   │   │       ├── addon_repository.go
│   │   │       ├── promotion_repository.go
│   │   │       ├── trial_repository.go
│   │   │       ├── billing_history_repository.go
│   │   │       ├── feature_toggle_repository.go
│   │   │       └── outbox_repository.go # ❌ REMOVED (use platform-shared/outbox/postgres/repository.go)
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go      # Redis connection setup (connection pooling, retry logic)
│   │   │       ├── subscription_cache.go # Cache user subscriptions (Get, Set, Invalidate with TTL)
│   │   │       ├── plan_cache.go      # Cache plans
│   │   │       ├── connect_cache.go   # Cache connect balances
│   │   │       └── feature_toggle_cache.go # Cache feature toggles
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go        # 📝 UPDATED: Kafka consumer (now uses platform-shared/inbox for message deduplication)
│   │   │       ├── producer.go        # 📝 UPDATED: Kafka producer (now uses platform-shared/outbox for reliable publishing)
│   │   │       ├── topics.go          # 📝 UPDATED: Topic constants (now imported from contracts/events - subscription.created, subscription.renewed, usage.limit.reached)
│   │   │       └── scram.go           # SCRAM authentication for Kafka (SASL/SCRAM-SHA-256)
│   │   ├── scheduler/
│   │   │   ├── cron.go                # Cron jobs for renewals
│   │   │   └── jobs.go                # Scheduled jobs
│   │   ├── payment/
│   │   │   └── client.go              # Financial service client for payments
│   │   └── outbox/
│   │       ├── processor.go           # ❌ REMOVED (use platform-shared/outbox/forwarder.go)
│   │       └── scheduler.go           # ❌ REMOVED (use platform-shared/outbox/scheduler.go)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── plan_handler.go    # Plan HTTP handlers (GET /plans)
│   │       │   ├── subscription_handler.go # Subscription handlers (POST /subscribe)
│   │       │   ├── connect_handler.go # Connect handlers (POST /connects/purchase)
│   │       │   ├── usage_handler.go   # Usage handlers
│   │       │   ├── addon_handler.go   # Addon handlers
│   │       │   ├── promotion_handler.go # Promotion handlers
│   │       │   ├── trial_handler.go   # Trial handlers
│   │       │   ├── billing_handler.go # Billing handlers
│   │       │   ├── feature_toggle_handler.go # Feature toggle handlers
│   │       │   └── health_handler.go  # Health check endpoints (/health, /ready, /live)
│   │       ├── middleware/
│   │       │   ├── auth.go            # 📝 UPDATED: Authentication middleware (uses pkg/auth for JWT verification)
│   │       │   ├── rbac.go            # 📝 UPDATED: RBAC middleware (uses pkg/auth authorizer for role checks)
│   │       │   ├── cors.go            # 📝 UPDATED: CORS middleware (uses platform-shared/ginx CORS)
│   │       │   ├── rate_limit.go      # Rate limiting middleware (token bucket algorithm)
│   │       │   ├── logging.go         # 📝 UPDATED: Logging middleware (structured logs per request, uses platform-shared/ginx/logging)
│   │       │   ├── error_handler.go   # Error handling middleware
│   │       │   ├── request_id.go      # 📝 UPDATED: Request ID middleware (X-Request-ID header, uses platform-shared/ginx/requestid)
│   │       │   └── feature_gate.go    # Feature access control middleware
│   │       ├── responses/
│   │       │   ├── success.go         # 📝 UPDATED: Success response wrappers (use platform-shared/httpx/response.go)
│   │       │   ├── error.go           # 📝 UPDATED: Error responses (use platform-shared/httpx/errors.go)
│   │       │   └── pagination.go      # 📝 UPDATED: Pagination (use platform-shared/httpx/pagination.go)
│   │       └── router.go              # 📝 UPDATED: HTTP router setup (uses Gin, applies platform-shared/ginx middleware)
│   │
│   └── config/                               # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                         # 🆕 Typed Config struct (App, Server, Postgres, Kafka, Redis)
│       ├── loader.go                         # 🆕 Config loader using Viper (precedence: flags → env → file → defaults)
│       └── docs/
│           └── CONFIGURATION.md              # 🆕 Configuration documentation (all ENV vars, defaults, examples)
│
├── config/                                   # 🆕 MEDIUM - Configuration files
│   ├── default.yaml                          # 🆕 Default configuration
│   ├── dev.yaml                              # 🆕 Development overrides
│   └── prod.yaml                             # 🆕 Production overrides
│
├── dapr/                                     # 🆕 MEDIUM - Dapr components split by environment
│   ├── local/                                # 🆕 For dapr run
│   │   ├── pubsub.yaml                       # Kafka pub/sub component
│   │   └── statestore.yaml                   # State store component
│   └── k8s/                                  # 🆕 For kubectl apply
│       ├── pubsub.yaml                       # Kafka with scopes: ["subscriptions-be"]
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
│   │   ├── billing_calculator.go             # Billing calculation utilities
│   │   └── proration.go                      # Proration logic
│   └── constants/
│       ├── events.go                         # ❌ REMOVED: Use contracts/events
│       ├── topics.go                         # ❌ REMOVED: Use contracts/events
│       ├── plans.go                          # Plan constants
│       └── features.go                       # Feature constants
│
├── seeds/
│   ├── plans.sql                             # Seed subscription plans
│   └── connect_packages.sql                  # Seed connect packages
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                   # Kubernetes Deployment
│       ├── service.yaml                      # Kubernetes Service
│       ├── configmap.yaml                    # ConfigMap
│       ├── secrets.yaml                      # Secrets
│       ├── hpa.yaml                          # HPA
│       ├── pdb.yaml                          # PDB
│       ├── cronjob-renewal.yaml              # Renewal cron job
│       └── servicemonitor.yaml               # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                        # Setup local environment
│   ├── get-secrets.sh                        # Fetch secrets
│   ├── seed-plans.sh                         # Seed plans
│   ├── seed-data.sh                          # Seed test data
│   └── process-renewals.sh                   # Process renewals
│
├── tests/
│   ├── unit/                                 # Unit tests
│   ├── integration/                          # Integration tests
│   └── e2e/                                  # End-to-end tests
│
├── docs/
│   ├── README.md                             # Service overview
│   ├── api.md                                # API documentation
│   ├── events.md                             # 📝 UPDATED: Events (published: subscription.created, connects.purchased, usage.limit.reached, consumed: payment.processed, proposal.submitted, etc.)
│   ├── subscription-plans.md                 # Subscription plans details
│   ├── connects-system.md                    # Connects system
│   ├── billing-logic.md                      # Billing logic
│   ├── feature-toggles.md                    # Feature toggles
│   ├── MIGRATIONS.md                         # 🆕 Migration history
│   ├── SCHEMA.md                             # 🆕 Database schema
│   └── RUNBOOK.md                            # 🆕 Operational procedures
│
├── .github/
│   └── workflows/
│       ├── ci.yml                            # CI workflow
│       └── cd.yml                            # CD workflow
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
│   │   │   ├── entity.go              # Admin user accounts
│   │   │   ├── role.go                # Admin roles (SuperAdmin, Moderator, Support)
│   │   │   ├── permission.go          # Granular permissions
│   │   │   ├── activity_log.go        # Admin activity tracking
│   │   │   └── repository.go
│   │   ├── support_ticket/
│   │   │   ├── entity.go              # Support tickets
│   │   │   ├── priority.go            # Priority levels (Low, Medium, High, Urgent)
│   │   │   ├── status.go              # Status (Open, InProgress, Resolved, Closed)
│   │   │   ├── category.go            # Ticket categories
│   │   │   ├── assignment.go          # Agent assignment
│   │   │   └── repository.go
│   │   ├── ticket_message/
│   │   │   ├── entity.go              # Ticket conversation messages
│   │   │   ├── attachment.go          # Message attachments
│   │   │   └── repository.go
│   │   ├── support_agent/
│   │   │   ├── entity.go              # Support agent profiles
│   │   │   ├── availability.go        # Agent availability/online status
│   │   │   ├── stats.go               # Agent performance stats
│   │   │   └── repository.go
│   │   ├── canned_response/
│   │   │   ├── entity.go              # Predefined responses
│   │   │   ├── category.go            # Response categories
│   │   │   └── repository.go
│   │   ├── knowledge_base/
│   │   │   ├── entity.go              # KB articles
│   │   │   ├── category.go            # Article categories
│   │   │   ├── tag.go                 # Article tags
│   │   │   ├── version.go             # Article versioning
│   │   │   └── repository.go
│   │   ├── faq/
│   │   │   ├── entity.go              # Frequently asked questions
│   │   │   ├── category.go            # FAQ categories
│   │   │   └── repository.go
│   │   ├── moderation_queue/
│   │   │   ├── entity.go              # Moderation queue items
│   │   │   ├── content_type.go        # Job, User, Review, Message, etc.
│   │   │   ├── flag_reason.go         # Reasons for flagging
│   │   │   ├── action.go              # Moderation actions taken
│   │   │   └── repository.go
│   │   ├── user_action/
│   │   │   ├── entity.go              # Admin actions on users
│   │   │   ├── action_type.go         # Suspend, Ban, Verify, Warn, etc.
│   │   │   ├── reason.go              # Action reasons
│   │   │   └── repository.go
│   │   ├── content_action/
│   │   │   ├── entity.go              # Admin actions on content
│   │   │   ├── action_type.go         # Remove, Hide, Approve, Reject
│   │   │   └── repository.go
│   │   ├── dispute_resolution/
│   │   │   ├── entity.go              # Dispute cases
│   │   │   ├── evidence.go            # Evidence submitted
│   │   │   ├── decision.go            # Admin decision
│   │   │   └── repository.go
│   │   ├── system_config/
│   │   │   ├── entity.go              # System configuration
│   │   │   ├── feature_flag.go        # Feature flags
│   │   │   ├── maintenance.go         # Maintenance mode settings
│   │   │   └── repository.go
│   │   ├── announcement/
│   │   │   ├── entity.go              # Platform announcements
│   │   │   ├── target.go              # Target audience (All, Freelancers, Clients)
│   │   │   ├── schedule.go            # Scheduled announcements
│   │   │   └── repository.go
│   │   ├── report/
│   │   │   ├── entity.go              # Generated reports
│   │   │   ├── report_type.go         # Report types (Users, Revenue, Activity)
│   │   │   ├── schedule.go            # Scheduled reports
│   │   │   └── repository.go
│   │   ├── analytics_dashboard/
│   │   │   ├── entity.go              # Dashboard configurations
│   │   │   ├── widget.go              # Dashboard widgets
│   │   │   ├── metric.go              # Custom metrics
│   │   │   └── repository.go
│   │   ├── audit_log/
│   │   │   ├── entity.go              # Complete audit trail
│   │   │   ├── action.go              # Actions performed
│   │   │   ├── resource.go            # Resources affected
│   │   │   └── repository.go
│   │   ├── notification_blast/
│   │   │   ├── entity.go              # Bulk notifications/emails
│   │   │   ├── audience.go            # Target audience
│   │   │   ├── schedule.go            # Scheduled blasts
│   │   │   └── repository.go
│   │   ├── platform_stats/
│   │   │   ├── entity.go              # Platform statistics
│   │   │   ├── realtime.go            # Real-time stats
│   │   │   └── repository.go
│   │   └── outbox/
│   │       ├── entity.go              # ❌ REMOVED (use platform-shared/outbox/entity.go)
│   │       └── repository.go          # ❌ REMOVED (use platform-shared/outbox/repository.go)
│   │
│   ├── application/
│   │   ├── admin_user/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Update, Deactivate
│   │   │   ├── queries.go             # Get, List admins
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── permission_manager.go  # Manage permissions
│   │   ├── support_ticket/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Assign, Resolve, Close
│   │   │   ├── queries.go             # Get, List, Search tickets
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   ├── assignment_engine.go   # Auto-assign tickets
│   │   │   ├── escalation_manager.go  # Escalate urgent tickets
│   │   │   └── sla_tracker.go         # Track SLA compliance
│   │   ├── ticket_message/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── support_agent/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── stats_calculator.go    # Calculate agent stats
│   │   ├── canned_response/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── knowledge_base/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Update, Publish
│   │   │   ├── queries.go             # Search, Get articles
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── search_service.go      # KB search
│   │   ├── faq/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── moderation/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Approve, Reject, Remove
│   │   │   ├── queries.go             # Get queue, Filter
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   ├── queue_manager.go       # Manage moderation queue
│   │   │   ├── auto_moderator.go      # Automatic moderation rules
│   │   │   └── content_scanner.go     # Scan for violations
│   │   ├── user_management/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Suspend, Ban, Verify, Warn
│   │   │   ├── queries.go             # Search users, Get details
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── action_validator.go    # Validate admin actions
│   │   ├── content_management/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Remove, Hide, Feature
│   │   │   ├── queries.go             # List content, Search
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── dispute_resolution/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Review, Decide, Close
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── decision_engine.go     # Help make decisions
│   │   ├── system_config/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Update configs
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── announcement/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Schedule, Send
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
│   │   │   ├── generator.go           # Generate reports
│   │   │   ├── scheduler.go           # Schedule reports
│   │   │   └── exporters/
│   │   │       ├── pdf_exporter.go    # PDF exporter
│   │   │       ├── csv_exporter.go    # CSV exporter
│   │   │       └── excel_exporter.go  # Excel exporter
│   │   ├── analytics/
│   │   │   ├── service.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── aggregator.go          # Aggregate analytics
│   │   │   ├── metrics_calculator.go  # Calculate KPIs
│   │   │   └── dashboard_builder.go   # Build dashboards
│   │   ├── notification_blast/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Schedule, Send
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── audience_selector.go   # Select target users
│   │   ├── audit/
│   │   │   ├── service.go
│   │   │   ├── queries.go             # Query audit logs
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── logger.go              # Log admin actions
│   │   └── eventhandler/
│   │       ├── user_handler.go        # 📝 UPDATED: Handle user events (now uses contracts/events)
│   │       ├── job_handler.go         # 📝 UPDATED: Handle job events (now uses contracts/events)
│   │       ├── proposal_handler.go    # 📝 UPDATED: Handle proposal events (now uses contracts/events)
│   │       ├── contract_handler.go    # 📝 UPDATED: Handle contract events (now uses contracts/events)
│   │       ├── payment_handler.go     # 📝 UPDATED: Handle payment events (now uses contracts/events)
│   │       ├── review_handler.go      # 📝 UPDATED: Handle review flags (now uses contracts/events)
│   │       ├── message_handler.go     # 📝 UPDATED: Handle message flags (now uses contracts/events)
│   │       ├── file_handler.go        # 📝 UPDATED: Handle file flags (now uses contracts/events)
│   │       └── system_handler.go      # 📝 UPDATED: Handle system events (now uses contracts/events)
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go      # PostgreSQL connection setup (DSN from config, connection pooling)
│   │   │       ├── transaction.go     # Transaction helpers (Begin, Commit, Rollback, WithTransaction wrapper)
│   │   │       ├── migrations.go      # 📝 UPDATED: Auto-migrate logic (now with version tracking, GORM AutoMigrate for all tables)
│   │   │       ├── version.go         # 🆕 Schema version tracking (SchemaVersion table, RecordMigration function)
│   │   │       ├── safety.go          # 🆕 Pre-migration safety checks (environment validation, disk space check)
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
│   │   │       └── outbox_repository.go # ❌ REMOVED (use platform-shared/outbox/postgres/repository.go)
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go      # Redis connection setup (connection pooling, retry logic)
│   │   │       ├── admin_cache.go     # Cache admin data (Get, Set, Invalidate with TTL)
│   │   │       ├── ticket_cache.go    # Cache tickets
│   │   │       ├── stats_cache.go     # Cache stats
│   │   │       └── config_cache.go    # Cache configs
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go        # 📝 UPDATED: Kafka consumer (now uses platform-shared/inbox for message deduplication)
│   │   │       ├── producer.go        # 📝 UPDATED: Kafka producer (now uses platform-shared/outbox for reliable publishing)
│   │   │       ├── topics.go          # 📝 UPDATED: Topic constants (now imported from contracts/events - admin.user.suspended, admin.dispute.resolved)
│   │   │       └── scram.go           # SCRAM authentication for Kafka (SASL/SCRAM-SHA-256)
│   │   ├── external_services/
│   │   │   ├── users_client.go        # Users service client
│   │   │   ├── jobs_client.go         # Jobs service client
│   │   │   ├── proposals_client.go    # Proposals service client
│   │   │   ├── contracts_client.go    # Contracts service client
│   │   │   ├── financial_client.go    # Financial service client
│   │   │   ├── reviews_client.go      # Reviews service client
│   │   │   ├── communications_client.go # Communications service client
│   │   │   ├── search_client.go       # Search service client
│   │   │   ├── storage_client.go      # Storage service client
│   │   │   └── subscriptions_client.go # Subscriptions service client
│   │   ├── keycloak/
│   │   │   ├── admin_client.go        # Keycloak admin operations
│   │   │   └── user_manager.go        # Manage users via Keycloak
│   │   ├── reporting/
│   │   │   ├── pdf_generator.go       # PDF generator
│   │   │   ├── csv_generator.go       # CSV generator
│   │   │   └── excel_generator.go     # Excel generator
│   │   └── outbox/
│   │       ├── processor.go           # ❌ REMOVED (use platform-shared/outbox/forwarder.go)
│   │       └── scheduler.go           # ❌ REMOVED (use platform-shared/outbox/scheduler.go)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── admin_user_handler.go # Admin user handlers
│   │       │   ├── support_ticket_handler.go # Ticket handlers
│   │       │   ├── ticket_message_handler.go # Message handlers
│   │       │   ├── support_agent_handler.go # Agent handlers
│   │       │   ├── canned_response_handler.go # Canned response handlers
│   │       │   ├── knowledge_base_handler.go # KB handlers
│   │       │   ├── faq_handler.go     # FAQ handlers
│   │       │   ├── moderation_handler.go # Moderation handlers
│   │       │   ├── user_management_handler.go # User management handlers
│   │       │   ├── content_management_handler.go # Content management handlers
│   │       │   ├── dispute_resolution_handler.go # Dispute handlers
│   │       │   ├── system_config_handler.go # Config handlers
│   │       │   ├── announcement_handler.go # Announcement handlers
│   │       │   ├── report_handler.go  # Report handlers
│   │       │   ├── analytics_handler.go # Analytics handlers
│   │       │   ├── notification_blast_handler.go # Blast handlers
│   │       │   ├── audit_log_handler.go # Audit handlers
│   │       │   ├── dashboard_handler.go # Dashboard handlers
│   │       │   └── health_handler.go  # Health check endpoints (/health, /ready, /live)
│   │       ├── middleware/
│   │       │   ├── auth.go            # 📝 UPDATED: Authentication middleware (uses pkg/auth for JWT verification)
│   │       │   ├── admin_auth.go      # 📝 UPDATED: Admin-specific auth (uses pkg/auth)
│   │       │   ├── permission_check.go # 📝 UPDATED: Check admin permissions (uses pkg/auth)
│   │       │   ├── audit_logger.go    # Auto-log all admin actions # 📝 UPDATED: Uses platform-shared/logging
│   │       │   ├── cors.go            # 📝 UPDATED: CORS middleware (uses platform-shared/ginx CORS)
│   │       │   ├── rate_limit.go      # Rate limiting middleware (token bucket algorithm)
│   │       │   ├── logging.go         # 📝 UPDATED: Logging middleware (structured logs per request, uses platform-shared/ginx/logging)
│   │       │   ├── error_handler.go   # Error handling middleware
│   │       │   └── request_id.go      # 📝 UPDATED: Request ID middleware (X-Request-ID header, uses platform-shared/ginx/requestid)
│   │       ├── responses/
│   │       │   ├── success.go         # 📝 UPDATED: Success response wrappers (use platform-shared/httpx/response.go)
│   │       │   ├── error.go           # 📝 UPDATED: Error responses (use platform-shared/httpx/errors.go)
│   │       │   └── pagination.go      # 📝 UPDATED: Pagination (use platform-shared/httpx/pagination.go)
│   │       └── router.go              # 📝 UPDATED: HTTP router setup (uses Gin, applies platform-shared/ginx middleware)
│   │
│   └── config/                               # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                         # 🆕 Typed Config struct (App, Server, Postgres, Kafka, Redis, Keycloak)
│       ├── loader.go                         # 🆕 Config loader using Viper (precedence: flags → env → file → defaults)
│       └── docs/
│           └── CONFIGURATION.md              # 🆕 Configuration documentation (all ENV vars, defaults, examples)
│
├── config/                                   # 🆕 MEDIUM - Configuration files
│   ├── default.yaml                          # 🆕 Default configuration
│   ├── dev.yaml                              # 🆕 Development overrides
│   └── prod.yaml                             # 🆕 Production overrides
│
├── dapr/                                     # 🆕 MEDIUM - Dapr components split by environment
│   ├── local/                                # 🆕 For dapr run
│   │   ├── pubsub.yaml                       # Kafka pub/sub component
│   │   └── statestore.yaml                   # State store component
│   └── k8s/                                  # 🆕 For kubectl apply
│       ├── pubsub.yaml                       # Kafka with scopes: ["admin-be"]
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
│   │   ├── sanitizer.go                      # Sanitize input
│   │   ├── permission_checker.go             # Permission checking utilities
│   │   └── report_formatter.go               # Report formatting
│   └── constants/
│       ├── events.go                         # ❌ REMOVED: Use contracts/events
│       ├── topics.go                         # ❌ REMOVED: Use contracts/events
│       ├── permissions.go                    # Permissions constants
│       └── moderation_actions.go             # Moderation actions constants
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                   # Kubernetes Deployment
│       ├── service.yaml                      # Kubernetes Service
│       ├── configmap.yaml                    # ConfigMap
│       ├── secrets.yaml                      # Secrets
│       ├── hpa.yaml                          # HPA
│       ├── pdb.yaml                          # PDB
│       ├── rbac.yaml                         # Admin RBAC policies
│       └── servicemonitor.yaml               # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                        # Setup local environment
│   ├── get-secrets.sh                        # Fetch secrets
│   ├── seed-admin-users.sh                   # Seed admin users
│   ├── seed-canned-responses.sh              # Seed canned responses
│   └── seed-data.sh                          # Seed test data
│
├── tests/
│   ├── unit/
│   │   ├── domain/
│   │   │   ├── admin_user_test.go     # Admin user tests
│   │   │   ├── support_ticket_test.go # Support ticket tests
│   │   │   └── moderation_queue_test.go # Moderation queue tests
│   │   ├── application/
│   │   │   ├── admin_user_service_test.go # Admin user service tests
│   │   │   ├── support_ticket_service_test.go # Support ticket service tests
│   │   │   └── moderation_service_test.go # Moderation service tests
│   │   └── infrastructure/
│   │       ├── postgres_repository_test.go # Postgres repository tests
│   │       └── kafka_producer_test.go # Kafka producer tests
│   ├── integration/
│   │   ├── handlers/
│   │   │   ├── admin_user_handler_test.go # Admin user handler tests
│   │   │   ├── support_ticket_handler_test.go # Support ticket handler tests
│   │   │   └── moderation_handler_test.go # Moderation handler tests
│   │   └── repositories/
│   │       ├── admin_user_repository_test.go # Admin user repository tests
│   │       └── support_ticket_repository_test.go # Support ticket repository tests
│   └── e2e/
│       └── scenarios/
│           ├── ticket_workflow_test.go # Ticket workflow tests
│           ├── moderation_workflow_test.go # Moderation workflow tests
│           └── dispute_resolution_test.go # Dispute resolution tests
│
├── docs/
│   ├── README.md                             # Service overview
│   ├── api.md                                # API documentation
│   ├── events.md                             # 📝 UPDATED: Events (published: admin.user.suspended, admin.dispute.resolved, consumed: user.flagged, contract.dispute.opened, etc.)
│   ├── admin-roles.md                        # Admin roles
│   ├── permissions.md                        # Permissions
│   ├── moderation-guide.md                   # Moderation guide
│   ├── support-workflows.md                  # Support workflows
│   ├── reporting.md                          # Reporting
│   ├── MIGRATIONS.md                         # 🆕 Migration history
│   ├── SCHEMA.md                             # 🆕 Database schema
│   └── RUNBOOK.md                            # 🆕 Operational procedures
│
├── .github/
│   └── workflows/
│       ├── ci.yml                            # CI workflow
│       └── cd.yml                            # CD workflow
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