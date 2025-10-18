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
│   └── api/
│       └── main.go                           # 📝 UPDATED: Application entry point - initializes Gin, Dapr, connects to Postgres (now uses loadConfig from internal/config, platform-shared/logging)
│
├── internal/
│   ├── domain/                               # 🏛️ Domain Layer - Business logic & entities
│   │   ├── user/
│   │   │   ├── entity.go                     # User aggregate root (ID, Username, Email, FirstName, LastName, UserType, Status)
│   │   │   ├── value_objects.go              # Email, Phone, Address value objects (validation, formatting)
│   │   │   ├── enums.go                      # UserType (Freelancer, Client), AccountStatus (Active, Suspended, Banned), VerificationStatus
│   │   │   ├── errors.go                     # Domain-specific errors (UserNotFound, EmailTaken, InvalidEmail)
│   │   │   ├── repository.go                 # User repository interface (Create, Update, FindByID, FindByEmail, List)
│   │   │   ├── list_filter.go                # ListFilter defines filtering options for list queries
│   │   │   ├── statistics.go                 # UserStatistics contains comprehensive user statistics
│   │   │   └── events.go                     # User events: UserCreated/Updated/Verified/Suspended/Banned/WarningIssued (v1 topics)
│   │   │
│   │   ├── profile/
│   │   │   ├── entity.go                     # Extended profile info (Bio, Location, ProfilePictureURL, CompletionPercentage)
│   │   │   ├── preferences.go                # User preferences and settings (Language, Timezone, Currency)
│   │   │   ├── availability.go               # Work availability, timezone, working hours
│   │   │   ├── errors.go                     # Profile-specific errors (ProfileNotFound, InvalidBio)
│   │   │   ├── repository.go                 # Profile repository interface
│   │   │   └── events.go                     # ProfileUpdated, PreferencesUpdated, AvailabilityUpdated
│   │   │
│   │   ├── skill/
│   │   │   ├── entity.go                     # User skills (UserID, SkillID, Proficiency, YearsOfExperience)
│   │   │   ├── proficiency.go                # Skill level enum (Beginner, Intermediate, Advanced, Expert)
│   │   │   ├── errors.go                     # Skill-specific errors (SkillNotFound, DuplicateSkill)
│   │   │   ├── repository.go                 # Skill repository interface
│   │   │   └── events.go                     # SkillAdded/Updated/Removed
│   │   │
│   │   ├── experience/
│   │   │   ├── entity.go                     # Work experience (Company, Title, Description, StartDate, EndDate, IsCurrent)
│   │   │   ├── errors.go                     # Experience-specific errors (InvalidDateRange)
│   │   │   ├── repository.go                 # Experience repository interface
│   │   │   └── events.go                     # ExperienceAdded/Updated/Removed
│   │   │
│   │   ├── education/
│   │   │   ├── entity.go                     # Educational background (School, Degree, Field, GraduationYear, Description)
│   │   │   ├── errors.go                     # Education-specific errors (InvalidYear)
│   │   │   ├── repository.go                 # Education repository interface
│   │   │   └── events.go                     # EducationAdded/Updated/Removed
│   │   │
│   │   ├── certification/
│   │   │   ├── entity.go                     # Certifications (Name, IssuingOrganization, IssueDate, ExpiryDate, CredentialID, URL)
│   │   │   ├── verification.go               # Verification status (Pending, Verified, Rejected, Expired)
│   │   │   ├── errors.go                     # Certification-specific errors (CertificationNotFound, AlreadyVerified)
│   │   │   ├── repository.go                 # Certification repository interface
│   │   │   └── events.go                     # CertificationAdded/Verified/Rejected/Expired
│   │   │
│   │   ├── portfolio/
│   │   │   ├── entity.go                     # Portfolio items (UserID, Title, Description, URL, ThumbnailURL, DisplayOrder)
│   │   │   ├── item.go                       # Portfolio item details (Images, Videos, Links)
│   │   │   ├── media.go                      # Associated media (Images, Videos, Documents)
│   │   │   ├── errors.go                     # Portfolio-specific errors (PortfolioNotFound, MaxItemsExceeded)
│   │   │   ├── repository.go                 # Portfolio repository interface
│   │   │   └── events.go                     # PortfolioItemAdded/Updated/Removed, MediaProcessed
│   │   │
│   │   ├── language/
│   │   │   ├── entity.go                     # Language proficiency (UserID, Language, ProficiencyLevel)
│   │   │   ├── errors.go                     # Language-specific errors (LanguageNotFound, DuplicateLanguage)
│   │   │   ├── repository.go                 # Language repository interface
│   │   │   └── events.go                     # LanguageAdded/Updated/Removed
│   │   │
│   │   ├── freelancer/
│   │   │   ├── entity.go                     # Freelancer-specific data (UserID, Title, Overview, VideoIntroURL)
│   │   │   ├── profile.go                    # Freelancer profile details (Tagline, ServiceOffering)
│   │   │   ├── rates.go                      # Hourly/fixed rates (HourlyRate, MinimumBudget, Currency)
│   │   │   ├── stats.go                      # Job stats, earnings (TotalJobs, TotalEarnings, SuccessRate, ResponseTime)
│   │   │   ├── errors.go                     # Freelancer-specific errors (FreelancerNotFound, InvalidRate)
│   │   │   ├── repository.go                 # Freelancer repository interface
│   │   │   └── events.go                     # FreelancerProfileUpdated, RatesUpdated, StatsUpdated
│   │   │
│   │   ├── client/
│   │   │   ├── entity.go                     # Client-specific data (UserID, CompanyName, CompanySize, Industry)
│   │   │   ├── profile.go                    # Client profile details (About, Website, Location)
│   │   │   ├── company.go                    # Company info (Name, Size, Industry, Founded, Employees)
│   │   │   ├── stats.go                      # Hiring stats, spending (TotalHires, TotalSpent, ActiveContracts, AverageRating)
│   │   │   ├── errors.go                     # Client-specific errors (ClientNotFound, InvalidCompanySize)
│   │   │   ├── repository.go                 # Client repository interface
│   │   │   └── events.go                     # ClientProfileUpdated, CompanyUpdated, ClientStatsUpdated
│   │   │
│   │   ├── verification/
│   │   │   ├── entity.go                     # Identity verification (UserID, Type, Status, SubmittedAt, VerifiedAt, RejectionReason)
│   │   │   ├── document.go                   # Verification documents (ID, Passport, ProofOfAddress, Selfie)
│   │   │   ├── errors.go                     # Verification-specific errors (VerificationNotFound, DocumentMissing)
│   │   │   ├── repository.go                 # Verification repository interface
│   │   │   └── events.go                     # VerificationSubmitted/Approved/Rejected
│   │   │
│   │   ├── settings/
│   │   │   ├── entity.go                     # User settings (UserID, Settings JSON, Theme, Language, Timezone)
│   │   │   ├── notification_prefs.go         # Notification preferences (Email, SMS, Push, InApp for different event types)
│   │   │   ├── privacy_settings.go           # Privacy settings (ProfileVisibility, ShowEmail, ShowPhone, SearchableProfile)
│   │   │   ├── errors.go                     # Settings-specific errors (SettingsNotFound, InvalidPreference)
│   │   │   ├── repository.go                 # Settings repository interface
│   │   │   └── events.go                     # SettingsUpdated, NotificationPrefsUpdated, PrivacySettingsUpdated
│   │   │
│   │   ├── saved_items/
│   │   │   ├── entity.go                     # Saved jobs, freelancers (UserID, ItemType, ItemID, SavedAt, Notes)
│   │   │   ├── errors.go                     # Saved items-specific errors (ItemNotFound, DuplicateSave)
│   │   │   ├── repository.go                 # Saved items repository interface
│   │   │   └── events.go                     # ItemSaved/Unsaved
│   │   │
│   │   ├── blocked_users/
│   │   │   ├── entity.go                     # Blocked user relationships (BlockerID, BlockedID, BlockedAt, Reason)
│   │   │   ├── errors.go                     # Blocking-specific errors (AlreadyBlocked, CannotBlockSelf)
│   │   │   ├── repository.go                 # Blocked users repository interface
│   │   │   └── events.go                     # UserBlocked/Unblocked
│   │   │
│   │   ├── user_suspension/
│   │   │   ├── entity.go                     # Track user suspensions (UserID, Reason, StartDate, EndDate, SuspendedBy, IsActive)
│   │   │   ├── reason.go                     # Suspension reasons enum (TOSViolation, PaymentIssue, QualityIssues, AbusiveeBehavior)
│   │   │   ├── duration.go                   # Suspension duration (Days, Weeks, Months, Permanent)
│   │   │   ├── errors.go                     # Suspension-specific errors (AlreadySuspended, InvalidDuration)
│   │   │   ├── repository.go                 # Suspension repository interface
│   │   │   └── events.go                     # UserSuspensionPlaced/Released
│   │   │
│   │   ├── user_ban/
│   │   │   ├── entity.go                     # Track user bans (UserID, Reason, BannedAt, BannedBy, IsPermanent, ExpiresAt)
│   │   │   ├── reason.go                     # Ban reasons enum (Fraud, SevereAbuse, MultipleViolations, SecurityThreat)
│   │   │   ├── permanent.go                  # Permanent vs temporary flag
│   │   │   ├── errors.go                     # Ban-specific errors (AlreadyBanned, CannotUnban)
│   │   │   ├── repository.go                 # Ban repository interface
│   │   │   └── events.go                     # UserBanPlaced/Released
│   │   │
│   │   ├── user_warning/
│   │   │   ├── entity.go                     # Track user warnings (UserID, Reason, IssuedAt, IssuedBy, AcknowledgedAt)
│   │   │   ├── reason.go                     # Warning reasons enum (LateDelivery, PoorQuality, UnresponsiveCommunication, MinorViolation)
│   │   │   ├── severity.go                   # Warning severity level (Low, Medium, High, Critical)
│   │   │   ├── errors.go                     # Warning-specific errors (WarningNotFound, TooManyWarnings)
│   │   │   ├── repository.go                 # Warning repository interface
│   │   │   └── events.go                     # UserWarningIssued/Acknowledged
│   │   │
│   │   ├── org/                              # 🆕 Org/agency accounts: teams, roles, seats, shared billing profiles
│   │   │   ├── entity.go                     # Org aggregate (OrgID, Name, OwnerID, BillingProfileID, SeatLimit)
│   │   │   ├── member.go                     # Membership (OrgID, UserID, Role[owner/admin/member], JoinedAt)
│   │   │   ├── seat.go                       # Seats usage (Total, Used, PendingInvites)
│   │   │   ├── errors.go                     # Org-specific errors (OrgNotFound, MemberExists, SeatLimitExceeded)
│   │   │   ├── repository.go                 # Org repository interface
│   │   │   └── events.go                     # OrgCreated/Updated, OrgMemberAdded/Removed, SeatUpdated
│   │   │
│   │   ├── security_center/                  # 🆕 2FA, device/session management, login alerts, recovery keys
│   │   │   ├── entity.go                     # Security settings (2FA enabled, RecoveryKeys hash)
│   │   │   ├── device.go                     # Registered devices (DeviceID, Fingerprint, LastSeen, Revoked)
│   │   │   ├── session.go                    # Sessions (SessionID, IP, UA, CreatedAt, ExpiresAt, Revoked)
│   │   │   ├── errors.go                     # Security-specific errors (DeviceNotFound, SessionNotFound)
│   │   │   ├── repository.go                 # Security repository interface
│   │   │   └── events.go                     # TwoFAEnabled/Disabled, DeviceRegistered/Revoked, SessionRevoked
│   │   │
│   │   ├── compliance/                       # 🆕 Tax profiles (W-forms/VAT), residency, country-specific fields
│   │   │   ├── tax_profile.go                # TaxProfile (UserID, Country, VAT/GST, TIN, W-Form refs)
│   │   │   ├── residency.go                  # Residency data (Country, Since, Proof docs)
│   │   │   ├── artifacts.go                  # Stored tax artifacts (W-8/W-9/VAT docs)
│   │   │   ├── errors.go                     # Compliance errors (TaxProfileNotFound, InvalidCountryFields)
│   │   │   ├── repository.go                 # Compliance repository interface
│   │   │   └── events.go                     # TaxProfileUpdated, ResidencyUpdated, ComplianceArtifactAdded
│   │   │
│   │   ├── risk_signals/                     # 🆕 Trust & safety signals
│   │   │   ├── signal.go                     # Signals (type: ip_geo_mismatch, disputes, chargebacks; severity; occurredAt)
│   │   │   ├── score.go                      # Risk score (UserID, score, updatedAt)
│   │   │   ├── hold.go                       # Account holds (type, reason, actor, until)
│   │   │   ├── errors.go                     # Risk errors (HoldExists, HoldNotFound)
│   │   │   ├── repository.go                 # Risk repository interface
│   │   │   └── events.go                     # RiskSignalRecorded, RiskScoreUpdated, RiskHoldPlaced/Released
│   │   │
│   │   └── profile_depth/                    # 🆕 Availability schedule, hourly rate history, standardized skills taxonomy, achievements/badges
│   │       ├── rate_history.go               # Rate entries (amount, currency, effectiveAt)
│   │       ├── availability_schedule.go      # Repeating availability slots (weekday, start, end, tz)
│   │       ├── taxonomy.go                   # Normalized skills taxonomy mapping
│   │       ├── badges.go                     # Achievements/badges (slug, awardedAt, reason)
│   │       ├── errors.go                     # Depth errors (DuplicateBadge, InvalidSlot)
│   │       ├── repository.go                 # Profile depth repository interface
│   │       └── events.go                     # RateHistoryUpdated, AvailabilityScheduleUpdated, BadgeAwarded
│   │   │
│   │   └── profile_depth/                    # 🆕 Availability schedule, hourly rate history, standardized skills taxonomy, achievements/badges
│   │       ├── rate_history.go               # Rate entries (amount, currency, effectiveAt)
│   │       ├── availability_schedule.go      # Repeating availability slots (weekday, start, end, tz)
│   │       ├── taxonomy.go                   # Normalized skills taxonomy mapping
│   │       ├── badges.go                     # Achievements/badges (slug, awardedAt, reason)
│   │       ├── errors.go                     # Depth errors (DuplicateBadge, InvalidSlot)
│   │       ├── repository.go                 # Profile depth repository interface
│   │       └── events.go                     # 🆕 Domain events: RateHistoryUpdated, AvailabilityUpdated, BadgeAwarded, BadgeRevoked

│   │
│   ├── application/                          # 📋 Application Layer - Use cases & orchestration
│   │   │
│   │   └── eventhandler/
│   │       ├── contract_handler.go              # Consumes: contract.created → update user stats;
│   │       │                                     # End-to-end flow confirmation: ContractCreated → users-be.
│   │       │
│   │       ├── escrow_handler.go                # Consumes (end-to-end flow confirmation): EscrowReleased → users-be.
│   │       │
│   │       ├── payment_handler.go               # Consumes: payment.processed → update earnings/spend;
│   │       │                                     # End-to-end flow confirmation: PaymentProcessed → users-be.
│   │       │
│   │       ├── review_handler.go                # Consumes: review.submitted → update ratings;
│   │       │                                     # End-to-end flow confirmation: ReviewSubmitted → users-be.
│   │       │
│   │       ├── message_handler.go               # Consumes: message.notification_delivered (NotificationDelivered) → users-be.
│   │       │
│   │       ├── admin_handler.go                 # Consumes: admin.user_suspended → enforce suspension;
│   │       │                                     #           admin.case.user_report.* → users-be;
│   │       │                                     #           admin.config.updated | admin.feature_flag.updated → users-be;
│   │       │                                     #           admin.data_export.* → notify subject via users-be.
│   │       │
│   │       ├── financial_risk_handler.go        # Consumes: financial.risk.alert.emitted → users-be (risk/status);
│   │       │                                     #           financial.chargeback.created | financial.chargeback.updated → users-be.
│   │       │
│   │       └── contract_status_handler.go       # Consumes: contract.financial_hold.placed | contract.financial_hold.released → users-be status effects.
│   │   ├
│   │   ├── user/
│   │   │   ├── service.go                    # User business logic (Create, Update, Delete, Verify, Search)
│   │   │   ├── commands.go                   # Command handlers (CreateUser, UpdateUser, DeleteUser, VerifyEmail)
│   │   │   ├── queries.go                    # Query handlers (GetUser, ListUsers, SearchUsers, GetUserStats)
│   │   │   ├── dto.go                        # Data transfer objects (UserDTO, CreateUserDTO, UpdateUserDTO, UserListDTO)
│   │   │   ├── mapper.go                     # Entity-DTO mapping (ToDTO, FromDTO, ToEntity, BulkToDTO)
│   │   │   └── validators.go                 # Input validation (ValidateEmail, ValidateUsername, ValidatePassword, ValidatePhone)
│   │   │
│   │   ├── profile/
│   │   │   ├── service.go                    # Profile business logic (Complete, Update, CalculateCompleteness)
│   │   │   ├── commands.go                   # UpdateProfile, UpdatePreferences, UpdateAvailability
│   │   │   ├── queries.go                    # GetProfile, GetAvailability, GetProfileCompletion
│   │   │   ├── dto.go                        # Profile DTOs (ProfileDTO, UpdateProfileDTO)
│   │   │   ├── mapper.go                     # Profile mappers
│   │   │   └── validators.go                 # Profile validators (ValidateBio, ValidateLocation)
│   │   │
│   │   ├── skill/
│   │   │   ├── service.go                    # Skill business logic (Add, Remove, Update, Reorder)
│   │   │   ├── commands.go                   # AddSkill, UpdateSkill, RemoveSkill, ReorderSkills
│   │   │   ├── queries.go                    # ListSkills, GetSkill
│   │   │   ├── dto.go                        # Skill DTOs (SkillDTO, AddSkillDTO)
│   │   │   ├── mapper.go                     # Skill mappers
│   │   │   └── validators.go                 # Validate skill name, proficiency level, duplicates
│   │   │
│   │   ├── experience/
│   │   │   ├── service.go                    # Experience business logic (Add, Update, Delete, Verify)
│   │   │   ├── commands.go                   # AddExperience, UpdateExperience, DeleteExperience
│   │   │   ├── queries.go                    # GetExperience, ListExperience
│   │   │   ├── dto.go                        # Experience DTOs (ExperienceDTO, CreateExperienceDTO)
│   │   │   ├── mapper.go                     # Experience mappers
│   │   │   └── validators.go                 # Validate date ranges, title/company presence
│   │   │
│   │   ├── education/
│   │   │   ├── service.go                    # Education business logic (Add, Update, Delete)
│   │   │   ├── commands.go                   # AddEducation, UpdateEducation, DeleteEducation
│   │   │   ├── queries.go                    # GetEducation, ListEducation
│   │   │   ├── dto.go                        # Education DTOs (EducationDTO, CreateEducationDTO)
│   │   │   ├── mapper.go                     # Education mappers
│   │   │   └── validators.go                 # Validate school/degree, graduation year range
│   │   │
│   │   ├── certification/
│   │   │   ├── service.go                    # Certification business logic (Add, Update, Delete, RequestVerification)
│   │   │   ├── commands.go                   # AddCertification, UpdateCertification, RemoveCertification, RequestVerification
│   │   │   ├── queries.go                    # GetCertification, ListCertifications
│   │   │   ├── dto.go                        # Certification DTOs (CertificationDTO, CreateCertificationDTO)
│   │   │   ├── mapper.go                     # Certification mappers
│   │   │   └── validators.go                 # Validate issuer, dates, credential URL/ID
│   │   │
│   │   ├── portfolio/
│   │   │   ├── service.go                    # Portfolio business logic (Add, Update, Delete, Reorder)
│   │   │   ├── commands.go                   # AddPortfolioItem, UpdatePortfolioItem, DeletePortfolioItem, ReorderPortfolio
│   │   │   ├── queries.go                    # GetPortfolioItem, ListPortfolio
│   │   │   ├── dto.go                        # Portfolio DTOs (PortfolioItemDTO, CreatePortfolioItemDTO)
│   │   │   ├── mapper.go                     # Portfolio mappers
│   │   │   └── validators.go                 # Validate URLs, media types, display order
│   │   │
│   │   ├── language/
│   │   │   ├── service.go                    # Language business logic (Add, Update, Delete)
│   │   │   ├── commands.go                   # AddLanguage, UpdateLanguage, RemoveLanguage
│   │   │   ├── queries.go                    # ListLanguages
│   │   │   ├── dto.go                        # Language DTOs (LanguageDTO, AddLanguageDTO)
│   │   │   ├── mapper.go                     # Language mappers
│   │   │   └── validators.go                 # Validate language code, proficiency enum
│   │   │
│   │   ├── freelancer/
│   │   │   ├── service.go                    # Freelancer business logic (CompleteProfile, UpdateRates, UpdateStats)
│   │   │   ├── commands.go                   # Freelancer commands (UpdateRates, UpdateAvailability)
│   │   │   ├── queries.go                    # Freelancer queries (GetStats, GetEarnings, GetSuccessRate)
│   │   │   ├── dto.go                        # Freelancer DTOs (FreelancerDTO, UpdateRatesDTO, FreelancerStatsDTO)
│   │   │   ├── mapper.go                     # Freelancer mappers
│   │   │   └── validators.go                 # Validate rate currency, min/max bounds, availability windows
│   │   │
│   │   ├── client/
│   │   │   ├── service.go                    # Client business logic (CompleteProfile, UpdateCompany, UpdateStats)
│   │   │   ├── commands.go                   # Client commands (UpdateCompany, UpdateIndustry)
│   │   │   ├── queries.go                    # Client queries (GetHiringStats, GetSpendingHistory)
│   │   │   ├── dto.go                        # Client DTOs (ClientDTO, UpdateCompanyDTO, ClientStatsDTO)
│   │   │   ├── mapper.go                     # Client mappers
│   │   │   └── validators.go                 # Validate company size, industry slug/NAICS, website URL
│   │   │
│   │   ├── verification/
│   │   │   ├── service.go                    # Verification business logic (Submit, Approve, Reject, RequestAdditionalDocs)
│   │   │   ├── commands.go                   # SubmitKYC, SubmitKYB, ApproveVerification, RejectVerification, TriggerReverification
│   │   │   ├── queries.go                    # GetVerification, ListVerifications, GetVerificationDocs
│   │   │   ├── dto.go                        # Verification DTOs (VerificationDTO, SubmitVerificationDTO)
│   │   │   ├── mapper.go                     # Verification mappers
│   │   │   └── validators.go                 # Validate document presence/types, reason codes, state transitions
│   │   │
│   │   ├── settings/
│   │   │   ├── service.go                    # Settings business logic (Update, GetPreferences, UpdatePrivacy)
│   │   │   ├── commands.go                   # UpdateSettings, UpdateNotificationPrefs, UpdatePrivacy
│   │   │   ├── queries.go                    # GetSettings, GetNotificationPrefs
│   │   │   ├── dto.go                        # Settings DTOs (SettingsDTO, UpdateSettingsDTO, NotificationPrefsDTO)
│   │   │   ├── mapper.go                     # Settings mappers
│   │   │   └── validators.go                 # Validate notification toggles, privacy flags, timezone/currency
│   │   │
│   │   ├── saved_items/
│   │   │   ├── service.go                    # Saved items business logic (Save, Unsave, List, Search)
│   │   │   ├── commands.go                   # SaveItem, UnsaveItem
│   │   │   ├── queries.go                    # ListSavedItems, SearchSaved
│   │   │   ├── dto.go                        # Saved items DTOs (SavedItemDTO, SaveItemDTO)
│   │   │   ├── mapper.go                     # Saved items mappers
│   │   │   └── validators.go                 # Validate item type/id combinations, dedupe constraints
│   │   │
│   │   ├── suspension/
│   │   │   ├── service.go                    # Suspension business logic (Suspend, Unsuspend, Extend, GetHistory)
│   │   │   ├── commands.go                   # SuspendUser, UnsuspendUser, ExtendSuspension
│   │   │   ├── queries.go                    # GetSuspensionHistory, GetActiveSuspension
│   │   │   ├── dto.go                        # Suspension DTOs (SuspensionDTO, SuspendUserDTO)
│   │   │   ├── mapper.go                     # Suspension mappers
│   │   │   └── validators.go                 # Validate duration ranges, reason enums, actor permissions
│   │   │
│   │   ├── ban/
│   │   │   ├── service.go                    # Ban business logic (Ban, Unban, GetBanHistory)
│   │   │   ├── commands.go                   # BanUser, UnbanUser
│   │   │   ├── queries.go                    # GetBanHistory, GetActiveBan
│   │   │   ├── dto.go                        # Ban DTOs (BanDTO, BanUserDTO)
│   │   │   ├── mapper.go                     # Ban mappers
│   │   │   └── validators.go                 # Validate permanence flags, expiry dates, reason enums
│   │   │
│   │   ├── warning/
│   │   │   ├── service.go                    # Warning business logic (IssueWarning, AcknowledgeWarning, GetWarnings)
│   │   │   ├── commands.go                   # IssueWarning, AcknowledgeWarning
│   │   │   ├── queries.go                    # ListWarnings
│   │   │   ├── dto.go                        # Warning DTOs (WarningDTO, IssueWarningDTO)
│   │   │   ├── mapper.go                     # Warning mappers
│   │   │   └── validators.go                 # Validate severity, reason, acknowledgement timestamp
│   │   │
│   │   ├── org/
│   │   │   ├── service.go                    # Org business logic (CreateOrg, InviteMember, AssignRole, ManageSeats)
│   │   │   ├── commands.go                   # CreateOrg, UpdateOrg, InviteMember, RemoveMember, AssignRole, SetSeatCount, LinkBillingProfile
│   │   │   ├── queries.go                    # GetOrg, ListOrgsForUser, ListMembers, GetRoles, GetSeatUsage
│   │   │   ├── dto.go                        # OrgDTO, OrgMemberDTO, SeatUsageDTO, BillingProfileLinkDTO
│   │   │   ├── mapper.go                     # Org mappers
│   │   │   └── validators.go                 # Org validators
│   │   │
│   │   ├── security_center/
│   │   │   ├── service.go                    # Enable2FA, Disable2FA, RegisterDevice, RevokeSession, SetRecoveryKeys
│   │   │   ├── commands.go                   # Enable2FA, Rotate2FA, RegisterDevice, RevokeDevice, RevokeSession, SetLoginAlerts, SetRecoveryKeys
│   │   │   ├── queries.go                    # Get2FAStatus, ListDevices, ListSessions, GetSecuritySettings
│   │   │   ├── dto.go                        # Security DTOs
│   │   │   ├── mapper.go                     # Security mappers
│   │   │   └── validators.go                 # Security validators
│   │   │
│   │   ├── compliance/
│   │   │   ├── service.go                    # UpsertTaxProfile, LinkUserTaxProfile, ValidateCountryFields
│   │   │   ├── commands.go                   # CreateOrUpdateTaxProfile, SetResidency, AttachWForm, AttachVAT
│   │   │   ├── queries.go                    # GetTaxProfile, GetResidency, ListComplianceArtifacts
│   │   │   ├── dto.go                        # TaxProfileDTO (W-forms/VAT), ResidencyDTO
│   │   │   ├── mapper.go                     # Compliance mappers
│   │   │   └── validators.go                 # Compliance validators
│   │   │
│   │   ├── risk_signals/
│   │   │   ├── service.go                    # ComputeRiskScore, RecordSignal, SetAccountHold
│   │   │   ├── commands.go                   # RecordIpGeoMismatch, RecordDisputeCount, RecordChargeback, ApplyHold, ReleaseHold
│   │   │   ├── queries.go                    # GetRiskScore, ListSignals, GetAccountState
│   │   │   ├── dto.go                        # RiskSignalDTO, RiskScoreDTO, AccountHoldDTO
│   │   │   ├── mapper.go                     # Risk mappers
│   │   │   └── validators.go                 # Risk validators
│   │   │
│   │   └── profile_depth/
│   │       ├── service.go                    # AppendRateHistory, AwardBadge, NormalizeSkills, UpsertAvailability
│   │       ├── commands.go                   # AddHourlyRateEntry, AwardAchievement, NormalizeSkillSet, SetAvailabilitySchedule
│   │       ├── queries.go                    # GetRateHistory, ListBadges, GetNormalizedSkills, GetAvailabilitySchedule
│   │       ├── dto.go                        # RateHistoryDTO, BadgeDTO, NormalizedSkillDTO, AvailabilitySlotDTO
│   │       ├── mapper.go                     # Profile depth mappers
│   │       └── validators.go                 # Profile depth validators
│   │
|   |
│   ├── infrastructure/                       # 🔧 Infrastructure Layer - External concerns
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection setup (DSN from config, connection pooling)
│   │   │       ├── transaction.go            # Transaction helpers (Begin, Commit, Rollback, WithTransaction wrapper)
│   │   │       ├── migrations.go             # 📝 UPDATED: Auto-migration logic (now with version tracking, GORM AutoMigrate for all tables)
│   │   │       ├── version.go                # 🆕 Schema version tracking (SchemaVersion table, RecordMigration function)
│   │   │       ├── safety.go                 # 🆕 Pre-migration safety checks (environment validation, disk space check, backup verification)
│   │   │       ├── user_repository.go        # User repository implementation (CRUD operations using GORM)
│   │   │       ├── profile_repository.go     # Profile repository implementation
│   │   │       ├── skill_repository.go       # Skill repository implementation
│   │   │       ├── experience_repository.go  # Experience repository implementation
│   │   │       ├── education_repository.go   # Education repository implementation
│   │   │       ├── certification_repository.go # Certification repository implementation
│   │   │       ├── portfolio_repository.go   # Portfolio repository implementation
│   │   │       ├── language_repository.go    # Language repository implementation
│   │   │       ├── freelancer_repository.go  # Freelancer repository implementation
│   │   │       ├── client_repository.go      # Client repository implementation
│   │   │       ├── verification_repository.go # Verification repository implementation
│   │   │       ├── settings_repository.go    # Settings repository implementation
│   │   │       ├── saved_items_repository.go # Saved items repository implementation
│   │   │       ├── blocked_users_repository.go # Blocked users repository implementation
│   │   │       ├── user_suspension_repository.go # Suspension repository implementation
│   │   │       ├── user_ban_repository.go    # Ban repository implementation
│   │   │       ├── user_warning_repository.go # Warning repository implementation
│   │   │       ├── org_repository.go         # 🆕 Orgs/teams/seats repository implementation
│   │   │       ├── security_repository.go    # 🆕 Devices/sessions/2FA/recovery keys repository
│   │   │       ├── compliance_repository.go  # 🆕 Tax/residency artifacts repository
│   │   │       ├── risk_signals_repository.go # 🆕 Risk signals & account holds repository
│   │   │       └── profile_depth_repository.go # 🆕 Rate history, badges, normalized skills, availability repository
│   │   │
│   │   │──  http/
│   │   │       └─ idempotency_adapter.go             # 🆕 NEW — bind platform-shared idempotency to Gin handlers
│   │   │──  messaging/
│   │   │       ├─ inbox/
│   │   │       │  └─ dapr_subscriptions.yaml         # 🆕 NEW — service-scoped Dapr subscriptions
│   │   │       └─ outbox/
│   │   │          └─ publisher.go                    # 🆕 NEW — thin wrapper over platform-shared outbox
│   │   │
│   │   │
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go             # Redis connection setup (connection pooling, retry logic)
│   │   │       ├── user_cache.go             # User caching logic (Get, Set, Invalidate with TTL, cache-aside pattern)
│   │   │       ├── profile_cache.go          # Profile caching logic (similar to user cache)
│   │   │       ├── org_cache.go              # 🆕 Cached org membership/roles
│   │   │       └── security_cache.go         # 🆕 Session/device cache
│   │   │
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go               # 📝 UPDATED: Kafka consumer (now uses platform-shared/inbox for message deduplication)
│   │   │       ├── producer.go               # 📝 UPDATED: Kafka producer (now uses platform-shared/outbox for reliable publishing)
│   │   │       ├── topics.go                 # 📝 UPDATED: Topic constants (now imported from contracts/events - user.created, user.updated, etc.)
│   │   │       └── scram.go                  # SCRAM authentication for Kafka (SASL/SCRAM-SHA-256)
│   │   │
│   │   ├── storage/
│   │   │   ├── client.go                     # Storage service client (upload profile pics, portfolios, documents via HTTP API)
│   │   │   └── local.go                      # Local file storage fallback (for development/testing)
│   │   │
│   │   └── keycloak/
│   │       ├── client.go                     # 📝 UPDATED: Keycloak API client (now uses pkg/auth for token verification)
│   │       └── sync.go                       # 📝 UPDATED: Sync users with Keycloak (create, update, delete users via Keycloak Admin API, uses pkg/auth)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── user_handler.go           # User HTTP handlers (GET, POST, PUT, DELETE /users, /users/:id, /users/search)
│   │       │   ├── profile_handler.go        # Profile HTTP handlers (GET, PUT /users/:id/profile, /users/:id/profile/completion)
│   │       │   ├── skill_handler.go          # Skill HTTP handlers (GET, POST, DELETE /users/:id/skills)
│   │       │   ├── experience_handler.go     # Experience HTTP handlers (GET, POST, PUT, DELETE /users/:id/experience)
│   │       │   ├── education_handler.go      # Education HTTP handlers (GET, POST, PUT, DELETE /users/:id/education)
│   │       │   ├── certification_handler.go  # Certification HTTP handlers (GET, POST, PUT, DELETE /users/:id/certifications)
│   │       │   ├── portfolio_handler.go      # Portfolio HTTP handlers (GET, POST, PUT, DELETE /users/:id/portfolio)
│   │       │   ├── language_handler.go       # Language HTTP handlers (GET, POST, DELETE /users/:id/languages)
│   │       │   ├── freelancer_handler.go     # Freelancer HTTP handlers (GET, PUT /users/:id/freelancer, /users/:id/freelancer/stats)
│   │       │   ├── client_handler.go         # Client HTTP handlers (GET, PUT /users/:id/client, /users/:id/client/stats)
│   │       │   ├── verification_handler.go   # Verification HTTP handlers (POST, GET /users/:id/verification)
│   │       │   ├── settings_handler.go       # Settings HTTP handlers (GET, PUT /users/:id/settings)
│   │       │   ├── saved_items_handler.go    # Saved items HTTP handlers (GET, POST, DELETE /users/:id/saved)
│   │       │   ├── suspension_handler.go     # Suspension HTTP handlers (POST, DELETE /admin/users/:id/suspend) - admin only
│   │       │   ├── ban_handler.go            # Ban HTTP handlers (POST, DELETE /admin/users/:id/ban) - admin only
│   │       │   ├── warning_handler.go        # Warning HTTP handlers (POST /admin/users/:id/warn) - admin only
│   │       │   ├── org_handler.go            # 🆕 Org/agency APIs (teams, roles, seats, billing links)
│   │       │   ├── security_center_handler.go# 🆕 2FA, sessions/devices, alerts, recovery keys
│   │       │   ├── compliance_handler.go     # 🆕 Tax/residency profiles
│   │       │   ├── risk_signals_handler.go   # 🆕 Risk score & account holds
│   │       │   └── profile_depth_handler.go  # 🆕 Rate history, badges, normalized skills, availability
│   │       │   └── health_handler.go               # 🆕 NEW — /healthz/live, /healthz/ready
│   │       │
│   │       ├── middleware/
│   │       │   ├─ requestid.go                    # 🆕 NEW — wraps platform-shared request id
│   │       │   ├─ logging.go                      # 🆕 NEW — wraps platform-shared logging
│   │       │   ├─ recovery.go                     # 🆕 NEW — wraps platform-shared recovery
│   │       │   ├─ tracing.go                      # 🆕 NEW — wraps platform-shared otel middleware
│   │       │   ├─ auth.go                         # 🆕 NEW — wraps pkg/auth (Keycloak)
│   │       │   └─ rbac.go                         # 🆕 NEW — role/permission gate
│   │       ├── routes/
│   │       │   ├── user_routes.go            # /users, /users/:id, /users/search
│   │       │   ├── profile_routes.go         # /users/:id/profile, /users/:id/profile/completion
│   │       │   ├── skill_routes.go           # /users/:id/skills
│   │       │   ├── experience_routes.go      # /users/:id/experience
│   │       │   ├── education_routes.go       # /users/:id/education
│   │       │   ├── certification_routes.go   # /users/:id/certifications
│   │       │   ├── portfolio_routes.go       # /users/:id/portfolio
│   │       │   ├── language_routes.go        # /users/:id/languages
│   │       │   ├── freelancer_routes.go      # /users/:id/freelancer, /users/:id/freelancer/stats
│   │       │   ├── client_routes.go          # /users/:id/client, /users/:id/client/stats
│   │       │   ├── verification_routes.go    # /users/:id/verification
│   │       │   ├── settings_routes.go        # /users/:id/settings
│   │       │   ├── saved_items_routes.go     # /users/:id/saved
│   │       │   ├── suspension_routes.go      # /admin/users/:id/suspend
│   │       │   ├── ban_routes.go             # /admin/users/:id/ban
│   │       │   ├── warning_routes.go         # /admin/users/:id/warn
│   │       │   ├── org_routes.go             # 🆕 /orgs, /orgs/:id, /orgs/:id/members, /orgs/:id/seats
│   │       │   ├── security_center_routes.go # 🆕 /security/2fa, /security/devices, /security/sessions
│   │       │   ├── compliance_routes.go      # 🆕 /compliance/tax-profile, /compliance/residency
│   │       │   ├── risk_signals_routes.go    # 🆕 /risk/signals, /risk/score, /risk/holds
│   │       │   └── profile_depth_routes.go   # 🆕 /users/:id/profile-depth (rates, badges, availability)
│   │       │
│   │       └── router.go                     # 📝 UPDATED: HTTP router setup (uses Gin, applies platform-shared/ginx middleware: requestid, logging, recover, otel, cors, timeout)
│   │
│   └── config/                               # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                         # 🆕 Typed Config struct (App name/version, Server port/host, Postgres DSN/pool, Kafka brokers/topics, Redis addr/password, Auth issuer/audience, Storage endpoint, Keycloak realm/client)
│       ├── loader.go                         # 🆕 Config loader using Viper (precedence: CLI flags → environment variables → config file → defaults)
│       └── docs/
│           └── CONFIGURATION.md              # 🆕 Configuration documentation (all ENV vars, default values, precedence rules, examples)
│
├── config/                                   # 🆕 MEDIUM - Configuration files
│   ├── default.yaml                          # 🆕 Default configuration (dev-friendly defaults: debug logging, local URLs, small pool sizes)
│   ├── dev.yaml                              # 🆕 Development overrides (debug mode enabled, local service URLs, no SSL)
│   └── prod.yaml                             # 🆕 Production overrides (info logging, production URLs, SSL enabled, larger pools, strict validation)
│
├── dapr/                                     # 🆕 MEDIUM - Dapr components split by environment
│   ├── local/
│   │   ├── pubsub.yaml                       # Kafka pub/sub component (localhost:9092, no auth, no scopes needed for local)
│   │   └── statestore.yaml                   # State store component (Redis or Postgres for local dev)
│   └── k8s/
│       ├── pubsub.yaml                       # Kafka pub/sub with proper auth (SCRAM-SHA-256) and scopes: ["users-be"]
│       ├── statestore.yaml                   # Production state store with scopes: ["users-be"]
│       └── secrets.yaml                      # Dapr secret store component (Kubernetes secrets or external vault)
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                         # Service-specific errors (domain errors, application errors)
│   │   └── codes.go                          # Error codes (USER_NOT_FOUND, EMAIL_TAKEN, INVALID_EMAIL, UNAUTHORIZED)
│   ├── logger/                               # ❌ REMOVED: Use platform-shared/logging instead
│   │   └── README.md
│   ├── utils/
│   │   ├── validator.go                      # Local validation utilities (if any service-specific validation beyond platform-shared)
│   │   ├── slug.go                           # Generate URL-friendly slugs (from username, company name)
│   │   └── sanitizer.go                      # Sanitize user input (HTML, SQL injection prevention)
│   │   └─ etag.go                               # 🆕 NEW — weak ETag helper for GET /users
│   └── constants/
│       ├── events.go                         # ❌ REMOVED: Use contracts/events instead (import from skillsier.dev/contracts/events/user/v1)
│       └── topics.go                         # ❌ REMOVED: Use contracts/events instead
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                   # Kubernetes Deployment manifest (replicas, resources, health checks)
│       ├── service.yaml                      # Kubernetes Service manifest (ClusterIP, ports)
│       ├── configmap.yaml                    # Configuration as ConfigMap (non-sensitive config from config/prod.yaml)
│       ├── secrets.yaml                      # Secrets (DB password, Kafka credentials, Keycloak client secret)
│       ├── hpa.yaml                          # Horizontal Pod Autoscaler (target CPU 70%, min 2, max 10 replicas)
│       ├── pdb.yaml                          # Pod Disruption Budget (maxUnavailable: 1 for safe rolling updates)
│       └── servicemonitor.yaml               # Prometheus ServiceMonitor (scrape /metrics endpoint)
│
├── scripts/
│   ├── setup-local.sh                        # Setup local development environment (Docker Compose for Postgres, Redis, Kafka)
│   ├── get-secrets.sh                        # Fetch secrets from Kubernetes (kubectl get secret)
│   └── seed-data.sh                          # Seed initial data (admin user, test users, sample profiles)
│
├── tests/
│   ├── unit/
│   │   ├── domain/
│   │   │   ├── user_test.go                  # User entity tests (validation, business rules)
│   │   │   ├── profile_test.go               # Profile entity tests
│   │   │   └── freelancer_test.go            # Freelancer entity tests
│   │   │   └─ user_validators_test.go               # 🆕 NEW — VO/validator tests
│   │   ├── application/
│   │   │   ├── user_service_test.go          # User service tests (mocked repositories)
│   │   │   ├── profile_service_test.go       # Profile service tests
│   │   │   └── freelancer_service_test.go    # Freelancer service tests
│   │   └── infrastructure/
│   │       ├── user_repository_test.go       # User repository tests (mocked DB)
│   │       └── kafka_producer_test.go        # Kafka producer tests
│   ├── integration/
│   │   ├── handlers/
│   │   │   ├── user_handler_test.go          # User handler integration tests (with test DB)
│   │   │   └── profile_handler_test.go       # Profile handler integration tests
│   │   └── repositories/
│   │       └── user_repository_test.go       # User repository integration tests (with real PostgreSQL testcontainer)
│   └── e2e/
│       └── scenarios/
│           ├── user_registration_test.go     # End-to-end user registration flow (Keycloak + DB + Kafka)
│           └── profile_completion_test.go    # End-to-end profile completion flow
│
├── docs/
│   ├── README.md                             # Service overview (purpose, responsibilities, dependencies)
│   ├── API.md                                # API documentation (endpoints, request/response examples, authentication)
│   ├── EVENTS.md                             # Service-specific events (published: user.created, user.updated; consumed: contract.completed, review.submitted)
│   ├── ARCHITECTURE.md                       # Service architecture details (layers, patterns, data flow)
│   ├── MIGRATIONS.md                         # Migration history (versions applied, changes made, rollback procedures, applied dates)
│   ├── SCHEMA.md                             # Current database schema documentation (tables, columns, indexes, relationships)
│   └── RUNBOOK.md                            # Operational procedures (how to restart, debug logs, rollback deployment, common issues)
│
├── .github/
│   └── workflows/
│       ├── ci.yml                            # CI workflow (go test, golangci-lint, docker build)
│       └── cd.yml                            # CD workflow (deploy to staging on merge to dev, deploy to prod on tag)
│
├── go.mod                                    # 📝 UPDATED: Dependencies (may import pkg/auth, platform-shared, contracts/events)
├── go.sum
├── .env.example                              # Example environment variables (DB_DSN, KAFKA_BROKERS, REDIS_ADDR, KEYCLOAK_REALM)
├── Makefile                                  # Build targets (run, test, lint, docker-build, k8s-deploy, migrate, seed)
├── Dockerfile                                # Multi-stage Docker build (builder + runtime, minimal alpine image)
├── .dockerignore                             # Docker ignore patterns (tests, docs, .git)
├── .gitignore                                # Git ignore patterns (vendor, bin, .env)
└── README.md                                 # Service README (quick start, development setup, deployment)

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
│   │   │   ├── value_objects.go              # JobBudget, JobSkill, JobCategory, JobDuration value objects
│   │   │   ├── enums.go                      # JobType (Fixed, Hourly, Retainer), Status (Draft, Open, InProgress, Closed), ExperienceLevel, BudgetType
│   │   │   ├── visibility.go                 # Public, Invite-only, Private-link visibility settings
│   │   │   ├── lifecycle.go                  # Lifecycle states (scheduled, auto-close rules, extensions/renewals)
│   │   │   ├── errors.go                     # Domain errors (ErrJobNotFound, ErrInvalidBudget, ErrJobAlreadyClosed)
│   │   │   ├── repository.go                 # JobRepository interface (Create, Update, FindByID, FindByClient, Search)
│   │   │   └── events.go                     # 🆕 Domain events: JobCreated, JobUpdated, JobPublished, JobClosed, JobArchived
│   │   │
│   │   ├── category/
│   │   │   ├── entity.go                     # Category entity (ID, Name, Slug, ParentID, Level) - hierarchical categories
│   │   │   ├── subcategory.go                # Nested categories support (tree structure)
│   │   │   ├── errors.go                     # Category errors
│   │   │   └── repository.go                 # CategoryRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: CategoryCreated, CategoryUpdated, CategoryReparented
│   │   │
│   │   ├── skill/
│   │   │   ├── entity.go                     # Skill taxonomy (ID, Name, CategoryID, Popularity)
│   │   │   ├── category.go                   # Skill categories (grouping related skills)
│   │   │   ├── errors.go                     # Skill errors
│   │   │   └── repository.go                 # SkillRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: SkillCreated, SkillUpdated, SkillDeprecated
│   │   │
│   │   ├── job_skill/
│   │   │   ├── entity.go                     # Skills required for job (JobID, SkillID, IsRequired)
│   │   │   ├── requirement_level.go          # Required vs Preferred skill levels
│   │   │   ├── errors.go                     # Job skill errors
│   │   │   └── repository.go                 # JobSkillRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: JobSkillAdded, JobSkillUpdated, JobSkillRemoved
│   │   │
│   │   ├── job_question/
│   │   │   ├── entity.go                     # Screening questions (JobID, Question, QuestionType, IsRequired)
│   │   │   ├── question_type.go              # Text, MultiChoice, File upload question types
│   │   │   ├── errors.go                     # Question errors
│   │   │   └── repository.go                 # JobQuestionRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: JobQuestionAdded, JobQuestionUpdated, JobQuestionRemoved
│   │   │
│   │   ├── job_attachment/
│   │   │   ├── entity.go                     # Job attachments (JobID, FileURL, FileType, FileName)
│   │   │   ├── errors.go                     # Attachment errors
│   │   │   └── repository.go                 # JobAttachmentRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: JobAttachmentAdded, JobAttachmentRemoved
│   │   │
│   │   ├── job_invitation/
│   │   │   ├── entity.go                     # Invitations to freelancers (JobID, FreelancerID, Message, Status)
│   │   │   ├── status.go                     # Invitation status (Pending, Accepted, Declined, Expired)
│   │   │   ├── errors.go                     # Invitation errors
│   │   │   └── repository.go                 # JobInvitationRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: JobInvitationSent, JobInvitationAccepted, JobInvitationDeclined, JobInvitationExpired
│   │   │
│   │   ├── job_view/
│   │   │   ├── entity.go                     # Job view tracking (JobID, UserID, ViewedAt, Duration)
│   │   │   ├── errors.go                     # View tracking errors
│   │   │   └── repository.go                 # JobViewRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: JobViewed, JobViewSessionUpdated
│   │   │
│   │   ├── saved_search/
│   │   │   ├── entity.go                     # Saved search filters (UserID, SearchQuery, Filters JSON, AlertEnabled)
│   │   │   ├── errors.go                     # Saved search errors
│   │   │   └── repository.go                 # SavedSearchRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: SavedSearchCreated, SavedSearchUpdated, SavedSearchDeleted
│   │   │
│   │   ├── job_flag/
│   │   │   ├── entity.go                     # Flagged jobs (JobID, ReporterID, Reason, Status, ReviewedAt)
│   │   │   ├── reason.go                     # Flag reasons (Spam, Inappropriate, Fraud, Misleading)
│   │   │   ├── status.go                     # Flag status (Pending, Reviewed, Resolved, Dismissed)
│   │   │   ├── errors.go                     # Flag errors
│   │   │   └── repository.go                 # JobFlagRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: JobFlagSubmitted, JobFlagResolved, JobFlagDismissed
│   │   │
│   │   ├── job_template/                     # 🆕 Job types & templates (fixed, hourly, retainer; reusable templates)
│   │   │   ├── entity.go                     # Template (TemplateID, Title, Type, DefaultBudget, DefaultScope, Skills, Attachments)
│   │   │   ├── type.go                       # JobType definitions (Fixed, Hourly, Retainer) + validation
│   │   │   ├── errors.go                     # Template errors (TemplateNotFound, InvalidTemplateType)
│   │   │   └── repository.go                 # JobTemplateRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: JobTemplateCreated, JobTemplateUpdated, JobTemplateArchived
│   │   │
│   │   ├── screening_compliance/             # 🆕 Screening & compliance (questions, tests, NDA, export-control)
│   │   │   ├── entity.go                     # Screening config (JobID, RequiredQuestions, SkillsTestLinks, NDARequired, ExportControlFlag)
│   │   │   ├── question.go                   # Question bank refs (id, type, isRequired)
│   │   │   ├── errors.go                     # Screening errors
│   │   │   └── repository.go                 # ScreeningRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: ScreeningConfigured, ScreeningRequirementUpdated, ScreeningCleared
│   │   │
│   │   ├── sourcing_modes/                   # 🆕 Sourcing modes (public, invite-only, private link; shortlists & talent pools)
│   │   │   ├── entity.go                     # Sourcing (JobID, Mode, PrivateLink, Pools[], Shortlists[])
│   │   │   ├── mode.go                       # Mode enum + constraints
│   │   │   ├── errors.go                     # Sourcing errors
│   │   │   └── repository.go                 # SourcingRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: SourcingModeSet, TalentPoolAttached, ShortlistUpdated
│   │   │
│   │   ├── budget_controls/                  # 🆕 Budget controls (min/max ranges, currency normalization, hourly caps)
│   │   │   ├── entity.go                     # BudgetControl (JobID, MinBudget, MaxBudget, Currency, RateCapHourly)
│   │   │   ├── fx_rule.go                    # FX rules (quote vs settlement currency, rounding)
│   │   │   ├── errors.go                     # Budget control errors
│   │   │   └── repository.go                 # BudgetControlRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: BudgetControlSet, BudgetControlUpdated
│   │   │
│   │   └── visibility_lifecycle/             # 🆕 Visibility & lifecycle (schedules, auto-close, extensions, draft/published)
│   │   │   ├── entity.go                     # VisibilityLifecycle (JobID, ScheduledAt, AutoCloseAt, RenewalPolicy, Draft, Publishe│At)
│   │   │   ├── policy.go                     # Rules engine structs for lifecycle transitions
│   │   │   ├── errors.go                     # Visibility/lifecycle errors
│   │   │   └── repository.go                 # VisibilityLifecycleRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: JobScheduled, JobPublished, JobAutoClosed, JobExtended
│   │   ├── template_versions/                # NEW — versioned job templates (semver, deprecations)
│   │   │   ├── entity.go                     # TemplateVersion aggregate (TemplateID, Version, Changelog, Deprecated)
│   │   │   ├── errors.go                     # ErrVersionConflict, ErrTemplateDeprecated, ErrInvalidSemVer
│   │   │   ├── repository.go                 # CreateVersion, GetLatest, ListByTemplate, Deprecate
│   │   │   └── events.go                     # JobTemplateVersionCreated, JobTemplateVersionDeprecated
│   │   ├── eligibility_rules/                # NEW — hard apply gates (geo allow/deny, KYC, min tier)
│   │   │   ├── entity.go                     # Eligibility{JobID, GeoRules, KYCRequired, MinTier, AgencyAllowed}
│   │   │   ├── enums.go                      # GateType, RuleEffect, Decision
│   │   │   ├── errors.go                     # ErrApplicantIneligible, ErrRuleConflict
│   │   │   ├── repository.go                 # SetRules, GetRules, EvaluateApplicant
│   │   │   └── events.go                     # JobEligibilityRulesSet, JobApplicantBlocked
│   │   ├── hiring_team/                      # NEW — collaborators/roles per job
│   │   │   ├── entity.go                     # HiringTeam{JobID, Members[UserID, Role], Notes}
│   │   │   ├── enums.go                      # Role: OWNER, REVIEWER, INTERVIEWER
│   │   │   ├── errors.go                     # ErrNotTeamMember, ErrRoleExists, ErrRoleChangeForbidden
│   │   │   ├── repository.go                 # AddMember, RemoveMember, ChangeRole, GetTeam
│   │   │   └── events.go                     # JobHiringTeamMemberAdded, JobHiringTeamMemberRemoved, JobHiringTeamRoleChanged
│   │   ├── pipeline/                         # NEW — per-job funnel stages & SLAs
│   │   │   ├── entity.go                     # Pipeline{JobID, Stages, WIPLimits, SLATimers, Timestamps}
│   │   │   ├── enums.go                      # Stage: NEW, SHORTLISTED, INTERVIEW, OFFER, HIRED, REJECTED
│   │   │   ├── errors.go                     # ErrInvalidStageTransition, ErrWipExceeded
│   │   │   ├── repository.go                 # Advance, Revert, SetWIP, Get
│   │   │   └── events.go                     # JobPipelineStageAdvanced, JobPipelineSLAExceeded
│   │   ├── requirements_matrix/              # NEW — weighted must vs nice-to-have
│   │   │   ├── entity.go                     # Requirements{JobID, Items[Key, MustHave, Weight], Version}
│   │   │   ├── errors.go                     # ErrInvalidWeighting, ErrDuplicateKey
│   │   │   ├── repository.go                 # SetMatrix, GetMatrix, ScoreProfile
│   │   │   └── events.go                     # JobRequirementsMatrixSet, JobRequirementWeightUpdated
│   │   ├── job_analytics/                    # NEW — aggregated KPIs (views → hire)
│   │   │   ├── entity.go                     # Analytics{ViewCount, ProposalCount, InviteCount, ShortlistCount, InterviewCount, OfferCount, HireCount, AvgResponseTime, EngagementScore}
│   │   │   ├── errors.go                     # ErrAnalyticsNotAvailable, ErrStatsStale
│   │   │   ├── repository.go                 # UpdateStats, GetStats, RecomputeScore
│   │   │   └── events.go                     # JobAnalyticsUpdated, JobLowEngagementAlert
│   │   ├── ab_experiments/                   # NEW — listing A/B (titles, descriptions, Q packs)
│   │   │   ├── entity.go                     # Experiment{JobID, Arms[], Allocation, Guardrails, AssignmentSalt}
│   │   │   ├── errors.go                     # ErrExperimentActive, ErrInvalidAllocation
│   │   │   ├── repository.go                 # Start, Stop, AssignArm, Get
│   │   │   └── events.go                     # JobExperimentStarted, JobExperimentStopped
│   │   ├── moderation_lifecycle/             # NEW — quarantine/limited visibility states
│   │   │   ├── entity.go                     # ModerationState{JobID, State, Reasons[], Actions[], Until}
│   │   │   ├── enums.go                      # State: CLEAN, LIMITED, QUARANTINED
│   │   │   ├── errors.go                     # ErrModerationStateLocked, ErrInvalidModerationTransition
│   │   │   ├── repository.go                 # ApplyState, Lift, Get
│   │   │   └── events.go                     # JobModerationStateApplied, JobModerationStateLifted
│   │   ├── syndication/                      # NEW — external boards/feeds integration
│   │   │   ├── entity.go                     # Syndication{JobID, Partner, Status, ExternalPostID, LastError, Retries}
│   │   │   ├── enums.go                      # Status: PENDING, POSTED, FAILED, TAKEDOWN
│   │   │   ├── errors.go                     # ErrPartnerConfigMissing, ErrPartnerQuotaExceeded
│   │   │   ├── repository.go                 # QueuePost, MarkPosted, MarkFailed, Takedown
│   │   │   └── events.go                     # JobSyndicated, JobSyndicationFailed, JobSyndicationTakedown
│   │   ├── drafts/                           # NEW — autosave & multi-draft
│   │   │   ├── entity.go                     # Draft{DraftID, JobID?, ClientID, Snapshot, LastEditorID, UpdatedAt}
│   │   │   ├── errors.go                     # ErrDraftNotFound, ErrStaleDraftVersion
│   │   │   ├── repository.go                 # Create, Update, Get, List, RestoreToJob
│   │   │   └── events.go                     # JobDraftSaved, JobDraftRestored
│   │   ├── geo_requirements/                 # NEW — onsite/remote/hybrid, TZ overlap, radius
│   │   │   ├── entity.go                     # Geo{JobID, Mode, Regions[], RadiusKm, TZOverlapHours}
│   │   │   ├── enums.go                      # Mode: ONSITE, REMOTE, HYBRID
│   │   │   ├── errors.go                     # ErrInvalidRegion, ErrInvalidRadius
│   │   │   ├── repository.go                 # SetGeo, GetGeo, ValidateCandidate
│   │   │   └── events.go                     # JobGeoRequirementsSet, JobGeoRequirementsUpdated
│   │   ├── duplicate_detection/              # NEW — near-duplicate jobs detection
│   │   │   ├── entity.go                     # DuplicateKey{JobID, SimHash, ClusterID, FirstSeenAt}
│   │   │   ├── errors.go                     # ErrDuplicateDetected, ErrHashUnavailable
│   │   │   ├── repository.go                 # UpsertHash, FindNearDuplicates
│   │   │   └── events.go                     # JobDuplicateDetected, JobDuplicateMerged
│   │   ├── interview_hooks/                  # NEW — interview policy flags/hooks
│   │   │   ├── entity.go                     # InterviewPolicy{JobID, PreferredSlotsRef, RequirePanel, Notes}
│   │   │   ├── errors.go                     # ErrInvalidSlotRef, ErrPolicyConflict
│   │   │   ├── repository.go                 # SetPolicy, GetPolicy
│   │   │   └── events.go                     # JobInterviewPolicySet
│   │   ├── legal_controls/                   # NEW — export control, NDA pinning, legal hold
│   │   │   ├── entity.go                     # Legal{JobID, ExportControlFlag, NDATemplateID, NDAVersion, LegalHold}
│   │   │   ├── errors.go                     # ErrNDAVersionMismatch, ErrHoldAlreadyPlaced
│   │   │   ├── repository.go                 # SetLegal, GetLegal, PlaceHold, RemoveHold
│   │   │   └── events.go                     # JobLegalControlsSet, JobLegalHoldPlaced, JobLegalHoldRemoved
│   │   ├── campaign_tags/                    # NEW — internal labels for campaigns/VIPs
│   │   │   ├── entity.go                     # CampaignTags{JobID, Tags[], AddedBy, AddedAt}
│   │   │   ├── errors.go                     # ErrTagLimitExceeded, ErrTagInvalid
│   │   │   ├── repository.go                 # AddTag, RemoveTag, ListTags
│   │   │   └── events.go                     # JobCampaignTagAdded, JobCampaignTagRemoved
│   │   ├── retention_rules/                  # NEW — archival/anonymization schedules
│   │   │   ├── entity.go                     # Retention{JobID, ArchiveAt, PurgeAt, ExemptReason}
│   │   │   ├── errors.go                     # ErrRetentionConflict, ErrPurgeBlocked
│   │   │   ├── repository.go                 # SetRetention, GetRetention
│   │   │   └── events.go                     # JobRetentionSet, JobArchived, JobPurged
│   │   ├── job_promotion/                    # NEW — paid boosts/featured jobs
│   │   │   ├── entity.go                     # Promotion{JobID, Status, BadgeType, StartAt, EndAt, FeeAmount, RenewalCount}
│   │   │   ├── enums.go                      # Status: PENDING, ACTIVE, SUSPENDED, EXPIRED
│   │   │   ├── errors.go                     # ErrPromotionNotEligible, ErrMaxRenewalsExceeded, ErrPromoSuspended
│   │   │   ├── repository.go                 # ActivatePromotion, RenewPromotion, Suspend, Expire, GetByJob
│   │   │   └── events.go                     # JobPromotionActivated, JobPromotionExpired, JobPromotionRenewed
│   │   ├── job_preference/                   # NEW — advisory matching preferences (non-blocking)
│   │   │   ├── entity.go                     # Preference{JobID, FreelancerType, PreferredLocations, TimeZones, MinSuccessScore, FluencyLevel, MinPlatformEarnings, GuidanceLevel, ToolProvision}
│   │   │   ├── enums.go                      # FreelancerType, FluencyLevel, GuidanceLevel, ToolProvision
│   │   │   ├── errors.go                     # ErrInvalidPreference, ErrPreferenceConflict
│   │   │   ├── repository.go                 # SetPreferences, Get, Update, Remove
│   │   │   └── events.go                     # JobPreferencesSet, JobPreferencesUpdated, JobPreferencesRemoved
│   │   └── job_hiring_option/                # NEW — multi-hire & repost config
│   │       ├── entity.go                     # HiringOption{JobID, MultiHireEnabled, MaxHires, OpenSlots, RepostAllowed, RepostReason, DuplicateCheckHash}
│   │       ├── errors.go                     # ErrMultiHireLimitExceeded, ErrJobDuplicateDetected
│   │       ├── repository.go                 # EnableMultiHire, ReserveSlot, ReleaseSlot, RepostJob, CheckDuplicate
│   │       └── events.go                     # JobMultiHireEnabled, JobReposted, JobDuplicatePrevented
│   │   ├── ai_assist/                        # 🆕 AI Assist (skills/category suggestions + description optimization)
│   │   │   ├── entity.go                     # SuggestionSet, OptimizationFeedback aggregates
│   │   │   ├── repository.go                 # Store accepted suggestions / optimization diffs
│   │   │   ├── errors.go                     # Domain errors (ErrNoSuggestions, ErrOptimizationConflict)
│   │   │   └── events.go                     # job.ai.suggestions.accepted.v1, job.description.optimized.v1
│   │   │
│   │   ├── job_inclusivity/                  # 🆕 Inclusivity / Accessibility flags
│   │   │   ├── entity.go                     # Flags (FlexibleHours, NoVideoRequired, ScreenReaderFriendly, etc.)
│   │   │   ├── repository.go                 # SetFlags, GetFlags
│   │   │   ├── errors.go                     # ErrInvalidInclusivityFlag
│   │   │   └── events.go                     # job.inclusivity.flags.set.v1
│   │   │
│   │   ├── job_payment/                      # 🆕 Payment schedule & terms
│   │   │   ├── entity.go                     # Terms (Milestones[], HourlyTerms, NetTerms, UpfrontPercent)
│   │   │   ├── repository.go                 # SetSchedule, UpdateTerms, GetByJob
│   │   │   ├── errors.go                     # ErrInvalidMilestone, ErrTermsConflict
│   │   │   └── events.go                     # job.payment.schedule.set.v1, job.payment.terms.updated.v1
│   │   │
│   │   ├── contract_transition/              # 🆕 Job → contracts-be transition mapping
│   │   │   ├── entity.go                     # TransitionRequest, Mapping (JobID→ContractSeed)
│   │   │   ├── repository.go                 # Queue, MarkSucceeded, MarkFailed, GetByJob
│   │   │   ├── errors.go                     # ErrTransitionNotReady, ErrMappingInvalid
│   │   │   └── events.go                     # job.contract.transitioned.v1
│   │   │
│   │   ├── fraud_detection/                  # 🆕 Fraud detection / risk auto-flags
│   │   │   ├── entity.go                     # RiskSignal, AutoFlag (reason, score, rule)
│   │   │   ├── repository.go                 # UpsertSignals, GetSignals, AutoFlagJob
│   │   │   ├── errors.go                     # ErrRiskRuleViolation
│   │   │   └── events.go                     # job.fraud.flagged.v1
│   │   │
│   │   ├── job_security/                     # 🆕 High-value posting checks (MFA)
│   │   │   ├── entity.go                     # HighValuePostCheck, MFAStatus
│   │   │   ├── repository.go                 # RequireMFA, VerifyMFAForJob
│   │   │   ├── errors.go                     # ErrMFANotVerified
│   │   │   └── events.go                     # job.security.mfa.verified.v1
│   │   │
│   │   ├── job_feedback/                     # 🆕 Post-closure feedback (to reviews-be)
│   │   │   ├── entity.go                     # Feedback{JobID, Role, Ratings, Comments}
│   │   │   ├── repository.go                 # Submit, GetByJob
│   │   │   ├── errors.go                     # ErrFeedbackDuplicate, ErrInvalidRating
│   │   │   └── events.go                     # job.feedback.submitted.v1
│   │   │
│   │   ├── job_esg/                          # 🆕 ESG / Sustainability
│   │   │   ├── entity.go                     # ESGFlags, CarbonEstimate snapshot
│   │   │   ├── repository.go                 # SetFlags, SetEstimate, Get
│   │   │   ├── errors.go                     # ErrInvalidESGFlag
│   │   │   └── events.go                     # job.esg.flags.set.v1, job.esg.estimate.calculated.v1
│   │   │
│   │   ├── job_sharing/                      # 🆕 Job sharing & referral incentives
│   │   │   ├── entity.go                     # TrackedShareLink{UTM}, Incentive
│   │   │   ├── repository.go                 # GenerateLink, SetIncentive, GetByJob
│   │   │   ├── errors.go                     # ErrShareQuotaExceeded
│   │   │   └── events.go                     # job.sharing.link.generated.v1, job.referral.incentive.set.v1
│   │   │
│   │   ├── job_tax/                          # 🆕 Tax requirements & reports
│   │   │   ├── entity.go                     # RequiredForms, ReportRequest
│   │   │   ├── repository.go                 # SetRequirements, RequestReport, Get
│   │   │   ├── errors.go                     # ErrFormMissing, ErrReportGenerationFailed
│   │   │   └── events.go                     # job.tax.requirements.set.v1, job.tax.report.generated.v1
│   │   │
│   │   ├── job_health/                       # 🆕 Post-hire health checkpoints
│   │   │   ├── entity.go                     # CheckpointSchedule, Nudge
│   │   │   ├── repository.go                 # Schedule, RecordNudge, GetByJob
│   │   │   ├── errors.go                     # ErrCheckpointInvalid
│   │   │   └── events.go                     # job.health.checkpoint.scheduled.v1
│   │   │
│   │   ├── job_upsell/                       # 🆕 Advisory upsell (read/advise; persist accepted)
│   │   │   ├── entity.go                     # UpsellSuggestion{type, payload, accepted}
│   │   │   ├── repository.go                 # List, RecordAcceptance
│   │   │   ├── errors.go                     # ErrUpsellInvalid
│   │   │   └── events.go                     # job.upsell.suggestion.accepted.v1
│   │   │
│   │   ├── job_previews/                     # 🆕 VR/AR previews (separate from attachments)
│   │   │   ├── entity.go                     # VRPreview, ARSpec
│   │   │   ├── repository.go                 # AttachPreview, GetByJob
│   │   │   ├── errors.go                     # ErrPreviewFormatUnsupported
│   │   │   └── events.go                     # job.preview.vr.attached.v1
│   │   │
│   │   ├── bulk_ops/                         # 🆕 Bulk operations (multi-job changes/imports)
│   │   │   ├── entity.go                     # BulkUpdateBatch, ImportBatch
│   │   │   ├── repository.go                 # CreateBatch, Append, MarkDone, Get
│   │   │   ├── errors.go                     # ErrBatchConflict, ErrImportParse
│   │   │   └── events.go                     # job.bulk.updated.v1, job.bulk.imported.v1
│   │   │
│   │   ├── webhooks/                         # 🆕 Webhook subscriptions (job-scoped)
│   │   │   ├── entity.go                     # WebhookSubscription, DeliveryPolicy
│   │   │   ├── repository.go                 # Subscribe, Unsubscribe, List
│   │   │   ├── errors.go                     # ErrInvalidEndpoint, ErrHMACRequired
│   │   │   └── events.go                     # job.webhook.subscribed.v1
│   │   │
│   │   ├── job_archive/                      # 🆕 User-managed archive/reactivate
│   │   │   ├── entity.go                     # ArchiveRecord{JobID, Reason, Actor, At}
│   │   │   ├── repository.go                 # Archive, Reactivate, GetHistory
│   │   │   ├── errors.go                     # ErrAlreadyArchived, ErrNotArchived
│   │   │   └── events.go                     # job.archived.v1, job.reactivated.v1
│   │   │
│   │   ├── job_tools/                        # 🆕 Job tools (time tracking prefs)
│   │   │   ├── entity.go                     # TimeTrackingPrefs{required, provider, rounding}
│   │   │   ├── repository.go                 # SetPrefs, GetByJob
│   │   │   ├── errors.go                     # ErrInvalidTimeTrackingProvider
│   │   │   └── events.go                     # job.tools.time_tracking.set.v1
│   │   │
│   │   ├── job_custom_fields/                # 🆕 Custom fields (schema + values)
│   │   │   ├── entity.go                     # CustomFieldSchema, FieldValue
│   │   │   ├── repository.go                 # AddField, RemoveField, SetValue, GetValues
│   │   │   ├── errors.go                     # ErrSchemaConflict, ErrValueInvalid
│   │   │   └── events.go                     # job.custom.field.added.v1
│   │   │
│   │   ├── data_residency/                   # 🆕 Data residency policy (if not under legal_controls)
│   │   │   ├── entity.go                     # ResidencyPolicy{Region: EU|US}
│   │   │   ├── repository.go                 # SetPolicy, GetPolicy
│   │   │   ├── errors.go                     # ErrResidencyConflict
│   │   │   └── events.go                     # job.compliance.residency.set.v1
│   │   │
│   │   └── job_localization/                 # 🆕 Multi-language job content
│   │       ├── entity.go                     # LocalizedDescription{Locale, Title, Summary, Body}, PrimaryLocale
│   │       ├── repository.go                 # UpsertLocale, RemoveLocale, GetLocales
│   │       ├── errors.go                     # ErrLocaleInvalid, ErrPrimaryLocaleMissing
│   │       └── events.go                     # job.languages.set.v1
│   │
│   ├── application/                          # 📋 Application Layer - Use cases & orchestration
│   │   ├── eventhandler/
│   │   │   ├── _routing_notes.go              # Partition keys: job events → key by job_id; user events → user_id.
│   │   │   │                                     # (See Partition Key Strategy.) 
│   │   │   │
│   │   │   ├── user_handler.go                # Consumes: user.created → track clients/owners for jobs domain.
│   │   │   │                                     # (Event Consumption Matrix: user.created → jobs-be.) 
│   │   │   │
│   │   │   ├── proposal_handler.go            # Consumes: proposal.submitted → update job proposal counters / readiness for acceptance.
│   │   │   │                                     # (Event Consumption Matrix: proposal.submitted → jobs-be.) 
│   │   │   │
│   │   │   ├── admin_handler.go               # Consumes: admin.feature_flag.updated / admin.config.updated → refresh jobs-be flags & config;
│   │   │   │                                     #           admin.moderation.action.applied → hide/remove job listing;
│   │   │   │                                     #           admin.user_suspended → enforce suspension constraints on job actions.
│   │   │   │                                     # (Added events → jobs-be consumer for flags/config; moderation action; All services suspend.) 
│   │   │   │
│   │   │   ├── subscription_handler.go        # Consumes: subscriptions.entitlement.updated → gate boosts/invites/posting features;
│   │   │   │                                     #           usage.counter.incremented / usage.limit.reached → enforce posting/boosts rate limits.
│   │   │   │                                     # (Subscriptions → jobs-be for entitlements & usage.) 
│   │   │   │
│   │   │   ├── financial_handler.go           # Consumes: financial.invoice.overdue → restrict posting/boosts; 
│   │   │   │                                     #           financial.fee.schedule.updated → update in-app budget/fee displays & validations.
│   │   │   │                                     # (Financial → jobs-be consumers.) 
│   │   │   │
│   │   │   └── search_handler.go              # Consumes: search.taxonomy.synonym.updated → refresh skills/categories & job facets mapping.
│   │   │    
│   │   ├── job/
│   │   │   ├── service.go                    # Job business logic (Create, Update, Publish, Close, Repost)
│   │   │   ├── commands.go                   # CreateJob, UpdateJob, PublishJob, CloseJob (validates client exists via users-be)
│   │   │   ├── queries.go                    # GetJob, ListJobs, SearchJobs with filters (category, budget, skills)
│   │   │   ├── dto.go                        # JobDTO, CreateJobDTO, UpdateJobDTO, JobListDTO, JobSearchDTO
│   │   │   ├── mapper.go                     # Entity-DTO mapping
│   │   │   ├── validators.go                 # Input validation (title, type∈{fixed,hourly,retainer}, budget ranges, dates)
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
│   │   │   ├── service.go                    # Skill business logic (List, Search, GetPopular)
│   │   │   ├── commands.go                   # UpsertSkill, RetireSkill
│   │   │   ├── queries.go                    # ListSkills, PopularSkills
│   │   │   ├── dto.go                        # SkillDTO, SkillListDTO
│   │   │   ├── mapper.go                     # Skill mappers
│   │   │   └── validators.go                 # Validate names, categories, duplicates
│   │   │
│   │   ├── invitation/
│   │   │   ├── service.go                    # Invitation business logic (Send, Accept, Decline)
│   │   │   ├── commands.go                   # SendInvitation, AcceptInvitation, DeclineInvitation
│   │   │   ├── queries.go                    # ListInvitationsForJob, GetInvitation
│   │   │   ├── dto.go                        # InvitationDTO, SendInvitationDTO
│   │   │   ├── mapper.go                     # Invitation mappers
│   │   │   └── validators.go                 # Validate recipient, message length, job state
│   │   │
│   │   ├── search/
│   │   │   ├── service.go                    # Search business logic (Search, Filter, Sort)
│   │   │   ├── commands.go                   # IndexJob, RemoveFromIndex
│   │   │   ├── queries.go                    # SearchJobs
│   │   │   ├── filter_builder.go             # Build complex search filters
│   │   │   ├── dto.go                        # SearchDTO, FilterDTO
│   │   │   └── validators.go                 # Validate filters, pagination, sort keys
│   │   │
│   │   ├── flag/
│   │   │   ├── service.go                    # Flag business logic (Flag, Review, Resolve)
│   │   │   ├── commands.go                   # FlagJob, UnflagJob, ResolveFlag
│   │   │   ├── queries.go                    # ListFlagsForJob
│   │   │   ├── dto.go                        # FlagDTO, FlagJobDTO
│   │   │   ├── mapper.go                     # Flag mappers
│   │   │   └── validators.go                 # Validate reasons/status transitions
│   │   │
│   │   ├── job_template/                     # 🆕 Job types & templates
│   │   │   ├── service.go                    # Create template, update defaults, clone into new job
│   │   │   ├── commands.go                   # CreateTemplate, UpdateTemplate, ArchiveTemplate
│   │   │   ├── queries.go                    # GetTemplate, ListTemplates, ListByType
│   │   │   ├── dto.go                        # JobTemplateDTO
│   │   │   ├── mapper.go                     # Template mappers
│   │   │   └── validators.go                 # Validate type, required defaults, allowed skills/attachments
│   │   │
│   │   ├── screening_compliance/             # 🆕 Screening & compliance
│   │   │   ├── service.go                    # Attach screening pack, toggle NDA, set export-control flag
│   │   │   ├── commands.go                   # SetQuestions, SetSkillsTests, ToggleNDA, SetExportControl
│   │   │   ├── queries.go                    # GetScreeningPack
│   │   │   ├── dto.go                        # ScreeningPackDTO, QuestionDTO
│   │   │   ├── mapper.go                     # Screening mappers
│   │   │   └── validators.go                 # Validate question types, NDA toggles, export-control flags
│   │   │
│   │   ├── sourcing_modes/                   # 🆕 Sourcing modes
│   │   │   ├── service.go                    # Switch mode, manage shortlists & talent pools
│   │   │   ├── commands.go                   # SetMode, AddToShortlist, RemoveFromShortlist, AddPool, RemovePool
│   │   │   ├── queries.go                    # GetSourcing, ListShortlists, ListPools
│   │   │   ├── dto.go                        # SourcingDTO, ShortlistDTO, PoolDTO
│   │   │   ├── mapper.go                     # Sourcing mappers
│   │   │   └── validators.go                 # Validate mode transitions, link formats, list sizes
│   │   │
│   │   ├── budget_controls/                  # 🆕 Budget controls
│   │   │   ├── service.go                    # Set ranges, normalize currency, enforce hourly caps
│   │   │   ├── commands.go                   # SetBudgetRange, SetCurrency, SetHourlyRateCap
│   │   │   ├── queries.go                    # GetBudgetControls, GetFxRules
│   │   │   ├── dto.go                        # BudgetControlsDTO, FxRuleDTO
│   │   │   ├── mapper.go                     # Budget mappers
│   │   │   └── validators.go                 # Validate min/max, ISO currency, non-negative rate caps
│   │   │
│   │   ├── visibility_lifecycle/             # 🆕 Visibility & lifecycle
│   │   │   ├── service.go                    # Schedule, publish/unpublish, auto-close, extend/renew
│   │   │   ├── commands.go                   # SchedulePosting, Publish, Unpublish, SetAutoClose, ExtendPosting
│   │   │   ├── queries.go                    # GetVisibilityLifecycle
│   │   │   ├── dto.go                        # VisibilityLifecycleDTO
│   │   │   ├── mapper.go                     # Visibility/Lifecycle mappers
│   │   │   ├── validators.go                 # Validate schedule times, state transitions, auto-close rules
│   │   ├── template_versions/                # NEW — version ops & queries
│   │   │   ├── service.go                    # Orchestrates version create/deprecate; emits events
│   │   │   ├── commands.go                   # CreateTemplateVersion, DeprecateTemplateVersion
│   │   │   ├── queries.go                    # GetLatestVersion, ListTemplateVersions
│   │   │   ├── dto.go                        # SemVer strings, changelog text
│   │   │   ├── mapper.go                     # Entity↔DTO conversion
│   │   │   └── validators.go                 # SemVer rules; non-empty changelog on major bump
│   │   ├── eligibility_rules/                # NEW — hard-gate evaluation
│   │   │   ├── service.go                    # Set/evaluate rules; return decision+reasons
│   │   │   ├── commands.go                   # SetRules, EvaluateApplicant
│   │   │   ├── queries.go                    # GetRules, GetRuleHistory
│   │   │   ├── dto.go                        # Geo/KYC/tier/agency payloads
│   │   │   ├── mapper.go                     # ISO country/tz normalization
│   │   │   └── validators.go                 # Conflicts & ranges
│   │   ├── hiring_team/                      # NEW — collaborators & roles
│   │   │   ├── service.go                    # Enforce OWNER invariants; audit actor ids
│   │   │   ├── commands.go                   # AddMember, ChangeRole, RemoveMember
│   │   │   ├── queries.go                    # GetTeam
│   │   │   ├── dto.go                        # MemberDTO, RoleDTO
│   │   │   ├── mapper.go                     # UserID↔Member mapping
│   │   │   └── validators.go                 # No duplicate roles; can’t remove last OWNER
│   │   ├── pipeline/                         # NEW — stage transitions & SLAs
│   │   │   ├── service.go                    # Guard rails; start/stop timers; emit transitions
│   │   │   ├── commands.go                   # AdvanceStage, RevertStage, SetWIP, SetSLA
│   │   │   ├── queries.go                    # GetPipeline, ListStageDurations
│   │   │   ├── dto.go                        # StageDTO, SLADTO
│   │   │   ├── mapper.go                     # Stage↔DTO; timestamps hydration
│   │   │   └── validators.go                 # Valid transitions; WIP/SLA ranges
│   │   ├── requirements_matrix/              # NEW — write & score matrix
│   │   │   ├── service.go                    # Score preview with reasons
│   │   │   ├── commands.go                   # SetMatrix, UpdateWeight
│   │   │   ├── queries.go                    # GetMatrix, ScoreCandidatePreview
│   │   │   ├── dto.go                        # RequirementDTO, ScoreDTO
│   │   │   ├── mapper.go                     # Flattened arrays ↔ entity items
│   │   │   └── validators.go                 # Weight sum ≤1; unique keys
│   │   ├── job_analytics/                    # NEW — analytics rollup
│   │   │   ├── service.go                    # Merge counters; recompute engagement score
│   │   │   ├── commands.go                   # UpdateStats, RecomputeScore
│   │   │   ├── queries.go                    # GetStats
│   │   │   ├── dto.go                        # AnalyticsDTO
│   │   │   ├── mapper.go                     # Safe zeros; monotonic merges
│   │   │   └── validators.go                 # Non-negative counters; window checks
│   │   ├── ab_experiments/                   # NEW — listing A/B orchestration
│   │   │   ├── service.go                    # Start/stop; guardrails; assignment hash
│   │   │   ├── commands.go                   # StartExperiment, StopExperiment
│   │   │   ├── queries.go                    # GetExperiment, ListExperiments
│   │   │   ├── dto.go                        # ExperimentDTO, ArmDTO
│   │   │   ├── mapper.go                     # Percent↔decimal
│   │   │   └── validators.go                 # Sum=100; unique arms
│   │   ├── moderation_lifecycle/             # NEW — apply/lift moderation
│   │   │   ├── service.go                    # Reconcile with admin cases
│   │   │   ├── commands.go                   # ApplyModerationState, LiftModerationState
│   │   │   ├── queries.go                    # GetModerationState, GetHistory
│   │   │   ├── dto.go                        # ModerationStateDTO
│   │   │   ├── mapper.go                     # Reason/action normalization
│   │   │   └── validators.go                 # Transition rules; reason codes
│   │   ├── syndication/                      # NEW — partner posting flow
│   │   │   ├── service.go                    # Queue post/takedown; handle callbacks
│   │   │   ├── commands.go                   # QueueSyndication, MarkFailed, Takedown
│   │   │   ├── queries.go                    # GetStatus, ListPartners
│   │   │   ├── dto.go                        # SyndicationDTO
│   │   │   ├── mapper.go                     # Partner ids mapping
│   │   │   └── validators.go                 # Retry caps; config presence
│   │   ├── drafts/                           # NEW — autosave & restore
│   │   │   ├── service.go                    # Conflict detection; restore into job
│   │   │   ├── commands.go                   # SaveDraft, RestoreToJob
│   │   │   ├── queries.go                    # GetDraft, ListDrafts
│   │   │   ├── dto.go                        # DraftDTO
│   │   │   ├── mapper.go                     # Snapshot merge
│   │   │   └── validators.go                 # Size/version checks
│   │   ├── geo_requirements/                 # NEW — geo/TZ rules
│   │   │   ├── service.go                    # TZ overlap calc; region normalization
│   │   │   ├── commands.go                   # SetGeo
│   │   │   ├── queries.go                    # GetGeo
│   │   │   ├── dto.go                        # GeoDTO
│   │   │   ├── mapper.go                     # Mode/regions mapping
│   │   │   └── validators.go                 # Radius/TZ ranges
│   │   ├── duplicate_detection/              # NEW — near-dupe ops
│   │   │   ├── service.go                    # Upsert simhash; search near-dupes
│   │   │   ├── commands.go                   # UpsertDuplicateKey
│   │   │   ├── queries.go                    # FindNearDuplicates
│   │   │   ├── dto.go                        # DuplicateMatchDTO (score, ids)
│   │   │   ├── mapper.go                     # Hash/score formatting
│   │   │   └── validators.go                 # Hash len; threshold
│   │   ├── interview_hooks/                  # NEW — policy flags
│   │   │   ├── service.go                    # Link to comms-be templates; no scheduling here
│   │   │   ├── commands.go                   # SetInterviewPolicy
│   │   │   ├── queries.go                    # GetInterviewPolicy
│   │   │   ├── dto.go                        # InterviewPolicyDTO
│   │   │   ├── mapper.go                     # External ref hydration
│   │   │   └── validators.go                 # SlotRef shape; invariants
│   │   ├── legal_controls/                   # NEW — NDA/export/holds
│   │   │   ├── service.go                    # Set flags; place/remove legal hold with audit
│   │   │   ├── commands.go                   # SetLegal, PlaceLegalHold, RemoveLegalHold
│   │   │   ├── queries.go                    # GetLegal
│   │   │   ├── dto.go                        # LegalDTO
│   │   │   ├── mapper.go                     # Template version pinning
│   │   │   └── validators.go                 # NDA version exists; hold blocks purge
│   │   ├── campaign_tags/                    # NEW — internal labels
│   │   │   ├── service.go                    # Add/remove; enforce limits; emit signal
│   │   │   ├── commands.go                   # AddTag, RemoveTag
│   │   │   ├── queries.go                    # ListTags
│   │   │   ├── dto.go                        # CampaignTagsDTO
│   │   │   ├── mapper.go                     # Normalize (lower, slug)
│   │   │   └── validators.go                 # Count≤limit; charset whitelist
│   │   ├── retention_rules/                  # NEW — archive/purge plans
│   │   │   ├── service.go                    # Compute deadlines; respect holds
│   │   │   ├── commands.go                   # SetRetention
│   │   │   ├── queries.go                    # GetRetention
│   │   │   ├── dto.go                        # RetentionDTO
│   │   │   ├── mapper.go                     # RFC3339 parsing
│   │   │   └── validators.go                 # Archive<Purge; future times
│   │   ├── job_promotion/                    # NEW — promotions lifecycle
│   │   │   ├── service.go                    # Activate/Renew/Suspend/Expire (after payment capture)
│   │   │   ├── commands.go                   # ActivatePromotion, RenewPromotion, SuspendPromotion, ExpirePromotion
│   │   │   ├── queries.go                    # GetPromotionByJob
│   │   │   ├── dto.go                        # PromotionDTO
│   │   │   ├── mapper.go                     # Money types; badge enum
│   │   │   └── validators.go                 # Windows; caps; transitions
│   │   ├── job_preference/                   # NEW — advisory prefs
│   │   │   ├── service.go                    # Set/update/remove; publish soft signals
│   │   │   ├── commands.go                   # SetPreferences, UpdatePreferences, RemovePreferences
│   │   │   ├── queries.go                    # GetPreferences
│   │   │   ├── dto.go                        # PreferenceDTO
│   │   │   ├── mapper.go                     # ISO country/tz codes
│   │   │   └── validators.go                 # Enum bounds; conflicts
│   │   └── job_hiring_option/                # NEW — multi-hire & repost
│   │       ├── service.go                    # Maintain slots; safe repost with dedupe
│   │       ├── commands.go                   # EnableMultiHire, ReserveSlot, ReleaseSlot, RepostJob
│   │       ├── queries.go                    # GetHiringOption
│   │       ├── dto.go                        # HiringOptionDTO
│   │       ├── mapper.go                     # Slot math; dup hash
│   │       └── validators.go                 # MaxHires≥OpenSlots; cooldowns
│   │
│   │   ├── ai_assist/
│   │   │   ├── service.go                    # Run suggestions, accept selections, apply optimization
│   │   │   ├── commands.go                   # AcceptSuggestions, OptimizeDescription
│   │   │   ├── queries.go                    # GetSuggestionsForJob
│   │   │   ├── dto.go                        # SuggestionDTO, OptimizationFeedbackDTO
│   │   │   ├── mapper.go                     # Entity↔DTO mappers
│   │   │   └── validators.go                 # Validate suggestion payloads / diffs
│   │   │
│   │   ├── job_inclusivity/
│   │   │   ├── service.go                    # Set/Get inclusivity flags
│   │   │   ├── commands.go                   # SetInclusivityFlags
│   │   │   ├── queries.go                    # GetInclusivityFlags
│   │   │   ├── dto.go                        # InclusivityFlagsDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate flags
│   │   │
│   │   ├── job_payment/
│   │   │   ├── service.go                    # Configure schedule & terms
│   │   │   ├── commands.go                   # SetPaymentSchedule, UpdatePaymentTerms
│   │   │   ├── queries.go                    # GetPaymentConfig
│   │   │   ├── dto.go                        # MilestoneDTO, TermsDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate milestones/terms
│   │   │
│   │   ├── contract_transition/
│   │   │   ├── service.go                    # Request/track contract transition
│   │   │   ├── commands.go                   # StartTransition, MarkTransitionSucceeded, MarkTransitionFailed
│   │   │   ├── queries.go                    # GetTransitionByJob
│   │   │   ├── dto.go                        # TransitionRequestDTO, MappingDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate readiness/mapping
│   │   │
│   │   ├── fraud_detection/
│   │   │   ├── service.go                    # Upsert risk signals, auto-flag
│   │   │   ├── commands.go                   # RecordRiskSignal, AutoFlagJob
│   │   │   ├── queries.go                    # GetRiskSignals
│   │   │   ├── dto.go                        # RiskSignalDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate rule inputs
│   │   │
│   │   ├── job_security/
│   │   │   ├── service.go                    # Enforce MFA for high-value posts
│   │   │   ├── commands.go                   # RequireJobMFA, VerifyJobMFA
│   │   │   ├── queries.go                    # GetMFAStatus
│   │   │   ├── dto.go                        # MFAStatusDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate challenge proofs
│   │   │
│   │   ├── job_feedback/
│   │   │   ├── service.go                    # Submit feedback to reviews-be
│   │   │   ├── commands.go                   # SubmitJobFeedback
│   │   │   ├── queries.go                    # GetJobFeedback
│   │   │   ├── dto.go                        # FeedbackDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate rating scales
│   │   │
│   │   ├── job_esg/
│   │   │   ├── service.go                    # Set ESG flags / carbon estimate
│   │   │   ├── commands.go                   # SetESGFlags, SetCarbonEstimate
│   │   │   ├── queries.go                    # GetESG
│   │   │   ├── dto.go                        # ESGFlagsDTO, CarbonEstimateDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate flags/units
│   │   │
│   │   ├── job_sharing/
│   │   │   ├── service.go                    # Generate tracked links, set incentives
│   │   │   ├── commands.go                   # GenerateShareLink, SetReferralIncentive
│   │   │   ├── queries.go                    # ListShareLinks
│   │   │   ├── dto.go                        # ShareLinkDTO, IncentiveDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate UTM/incentives
│   │   │
│   │   ├── job_tax/
│   │   │   ├── service.go                    # Set tax reqs & request reports
│   │   │   ├── commands.go                   # SetTaxRequirements, RequestTaxReport
│   │   │   ├── queries.go                    # GetTaxSettings
│   │   │   ├── dto.go                        # TaxRequirementsDTO, ReportRequestDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate forms/regions
│   │   │
│   │   ├── job_health/
│   │   │   ├── service.go                    # Schedule checkpoints, send nudges
│   │   │   ├── commands.go                   # ScheduleCheckpoint, SendNudge
│   │   │   ├── queries.go                    # GetHealthPlan
│   │   │   ├── dto.go                        # CheckpointDTO, NudgeDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate cadence/SLA
│   │   │
│   │   ├── job_upsell/
│   │   │   ├── service.go                    # Fetch upsell suggestions, record acceptance
│   │   │   ├── commands.go                   # AcceptUpsellSuggestion
│   │   │   ├── queries.go                    # ListUpsellSuggestions
│   │   │   ├── dto.go                        # UpsellSuggestionDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate suggestion types
│   │   │
│   │   ├── job_previews/
│   │   │   ├── service.go                    # Attach VR/AR previews
│   │   │   ├── commands.go                   # AttachVRPreview, AttachARSpec
│   │   │   ├── queries.go                    # GetPreviews
│   │   │   ├── dto.go                        # VRPreviewDTO, ARSpecDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate formats/links
│   │   │
│   │   ├── bulk_ops/
│   │   │   ├── service.go                    # Orchestrate bulk updates & CSV imports
│   │   │   ├── commands.go                   # StartBulkUpdate, AppendBulkUpdate, StartImport
│   │   │   ├── queries.go                    # GetBatchStatus
│   │   │   ├── dto.go                        # BulkUpdateBatchDTO, ImportBatchDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate batch sizes / CSV schema
│   │   │
│   │   ├── webhooks/
│   │   │   ├── service.go                    # Manage job-scoped webhook subscriptions
│   │   │   ├── commands.go                   # SubscribeWebhook, UnsubscribeWebhook
│   │   │   ├── queries.go                    # ListWebhooks
│   │   │   ├── dto.go                        # WebhookSubscriptionDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate endpoint/HMAC
│   │   │
│   │   ├── job_archive/
│   │   │   ├── service.go                    # Archive / Reactivate by user action
│   │   │   ├── commands.go                   # ArchiveJob, ReactivateJob
│   │   │   ├── queries.go                    # GetArchiveHistory
│   │   │   ├── dto.go                        # ArchiveRecordDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate state transitions
│   │   │
│   │   ├── job_tools/
│   │   │   ├── service.go                    # Set time tracking prefs
│   │   │   ├── commands.go                   # SetTimeTrackingPrefs
│   │   │   ├── queries.go                    # GetTimeTrackingPrefs
│   │   │   ├── dto.go                        # TimeTrackingPrefsDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate rounding/providers
│   │   │
│   │   ├── job_custom_fields/
│   │   │   ├── service.go                    # Manage schema & values
│   │   │   ├── commands.go                   # AddCustomField, RemoveCustomField, SetFieldValue
│   │   │   ├── queries.go                    # ListCustomFields, GetFieldValues
│   │   │   ├── dto.go                        # CustomFieldSchemaDTO, FieldValueDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate schema/value types
│   │   │
│   │   ├── data_residency/
│   │   │   ├── service.go                    # Set/Get data residency policy
│   │   │   ├── commands.go                   # SetResidencyPolicy
│   │   │   ├── queries.go                    # GetResidencyPolicy
│   │   │   ├── dto.go                        # ResidencyPolicyDTO
│   │   │   ├── mapper.go                     # Mappers
│   │   │   └── validators.go                 # Validate region compatibility
│   │   │
│   │   └── job_localization/
│   │       ├── service.go                    # Upsert/remove locale variants
│   │       ├── commands.go                   # UpsertJobLocale, RemoveJobLocale, SetPrimaryLocale
│   │       ├── queries.go                    # GetJobLocales
│   │       ├── dto.go                        # LocalizedDescriptionDTO
│   │       ├── mapper.go                     # Mappers
│   │       └── validators.go                 # Validate locale codes / primary locale
│   ├── infrastructure/                       # 🔧 Infrastructure Layer
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection setup
│   │   │       ├── transaction.go            # Transaction helpers
│   │   │       ├── migrations.go             # 📝 UPDATED: Auto-migration with version tracking
│   │   │       ├── version.go                # 🆕 Schema version tracking
│   │   │       ├── safety.go                 # 🆕 Pre-migration safety checks
│   │   │       ├── job_repository.go         # JobRepository implementation with GORM
│   │   │       ├── category_repository.go    # CategoryRepository implementation
│   │   │       ├── skill_repository.go       # SkillRepository implementation
│   │   │       ├── job_skill_repository.go   # JobSkillRepository implementation
│   │   │       ├── job_question_repository.go# JobQuestionRepository implementation
│   │   │       ├── job_attachment_repository.go # JobAttachmentRepository implementation
│   │   │       ├── job_invitation_repository.go # JobInvitationRepository implementation
│   │   │       ├── job_view_repository.go    # JobViewRepository implementation
│   │   │       ├── saved_search_repository.go# SavedSearchRepository implementation
│   │   │       ├── job_flag_repository.go    # JobFlagRepository implementation
│   │   │       ├── job_template_repository.go# 🆕 Templates repository implementation
│   │   │       ├── screening_repository.go   # 🆕 Screening & compliance repository implementation
│   │   │       ├── sourcing_repository.go    # 🆕 Sourcing modes repository implementation
│   │   │       ├── budget_controls_repository.go # 🆕 Budget controls repository implementation
│   │   │       ├── visibility_lifecycle_repository.go # 🆕 Visibility/lifecycle repository implementation
│   │   │       ├── template_versions_repository.go    # NEW — unique (template_id, version)
│   │   │       ├── eligibility_rules_repository.go    # NEW — JSONB policies; idx job_id
│   │   │       ├── hiring_team_repository.go          # NEW — team_members table; uniq (job_id, user_id)
│   │   │       ├── pipeline_repository.go             # NEW — stage timeline; SLA durations
│   │   │       ├── requirements_matrix_repository.go  # NEW — weights; check 0..1
│   │   │       ├── job_analytics_repository.go        # NEW — counters upsert by job_id
│   │   │       ├── ab_experiments_repository.go       # NEW — experiments/arms tables
│   │   │       ├── moderation_lifecycle_repository.go # NEW — append-only history
│   │   │       ├── syndication_repository.go          # NEW — partner status + retries
│   │   │       ├── drafts_repository.go               # NEW — JSONB drafts; last_editor_id
│   │   │       ├── geo_requirements_repository.go     # NEW — geo/TZ; region indexes
│   │   │       ├── duplicate_detection_repository.go  # NEW — simhash; nearest-neighbor idx
│   │   │       ├── interview_hooks_repository.go      # NEW — policy refs to comms artifacts
│   │   │       ├── legal_controls_repository.go       # NEW — legal holds immutable
│   │   │       ├── campaign_tags_repository.go        # NEW — case-insensitive tags
│   │   │       ├── retention_rules_repository.go      # NEW — archive/purge indexes
│   │   │       ├── job_promotion_repository.go        # NEW — promo state + billing corr id
│   │   │       ├── job_preference_repository.go       # NEW — ISO arrays; GIN
│   │   │       └── job_hiring_option_repository.go    # NEW — slot counters; dup hash
│   │   │       ├── ai_assist_repository.go           # NEW — AI Assist repo (suggestions/optimizations)
│   │   │       ├── job_inclusivity_repository.go     # NEW — Inclusivity flags repo
│   │   │       ├── job_payment_repository.go         # NEW — Payment schedule/terms repo
│   │   │       ├── contract_transition_repository.go # NEW — Contract transition mapping repo
│   │   │       ├── fraud_detection_repository.go     # NEW — Risk signals & auto-flags repo
│   │   │       ├── job_security_repository.go        # NEW — High-value MFA status repo
│   │   │       ├── job_feedback_repository.go        # NEW — Post-closure feedback repo
│   │   │       ├── job_esg_repository.go             # NEW — ESG flags & carbon estimate repo
│   │   │       ├── job_sharing_repository.go         # NEW — Tracked links & referral incentives repo
│   │   │       ├── job_tax_repository.go             # NEW — Tax requirements & reports repo
│   │   │       ├── job_health_repository.go          # NEW — Health checkpoint schedules repo
│   │   │       ├── job_upsell_repository.go          # NEW — Upsell suggestions repo
│   │   │       ├── job_previews_repository.go        # NEW — VR/AR previews repo
│   │   │       ├── bulk_ops_repository.go            # NEW — Bulk update/import batches repo
│   │   │       ├── webhooks_repository.go            # NEW — Webhook subscriptions repo
│   │   │       ├── job_archive_repository.go         # NEW — User-managed archive/reactivate repo
│   │   │       ├── job_tools_repository.go           # NEW — Time tracking prefs repo
│   │   │       ├── job_custom_fields_repository.go   # NEW — Custom field schema/values repo
│   │   │       ├── data_residency_repository.go      # NEW — Data residency policy repo
│   │   │       └── job_localization_repository.go    # NEW — Localized descriptions repo
│   │   │
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go             # Redis connection
│   │   │       ├── job_cache.go              # Job caching (Get, Set, Invalidate)
│   │   │       ├── category_cache.go         # Category tree caching
│   │   │       ├── skill_cache.go            # Popular skills caching
│   │   │       ├── templates_cache.go        # 🆕 Templates caching
│   │   │       └── sourcing_cache.go         # 🆕 Sourcing mode & lists cache
│   │   │       ├── template_versions_cache.go         # NEW — latest version cache
│   │   │       ├── eligibility_rules_cache.go         # NEW — policy by job_id w/ short TTL
│   │   │       ├── hiring_team_cache.go               # NEW — roster cache; bust on changes
│   │   │       ├── pipeline_cache.go                  # NEW — current stage & SLAs
│   │   │       ├── requirements_matrix_cache.go       # NEW — matrix + ETag
│   │   │       ├── job_analytics_cache.go             # NEW — counters & score TTL
│   │   │       ├── moderation_lifecycle_cache.go      # NEW — moderation state
│   │   │       ├── drafts_cache.go                    # NEW — user-scoped snapshots
│   │   │       ├── geo_requirements_cache.go          # NEW — geo/TZ for filtering
│   │   │       ├── job_promotion_cache.go             # NEW — promo status until expiry
│   │   │       └── job_preference_cache.go            # NEW — soft prefs for prefilter
│   │   |
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go               # 📝 UPDATED: Uses platform-shared/inbox
│   │   │       ├── producer.go               # 📝 UPDATED: Uses platform-shared/outbox
│   │   │       ├── scram.go                  # SCRAM authentication
│   │   │       ├── topics_jobs_extensions.go          # NEW — helpers for new events (contracts/events)
│   │   │       └── producers_jobs_extensions.go       # NEW — thin producers via platform-shared/outbox
│   │   │
│   │   ├── search/
│   │   │   └── client.go                     # Search service client (call search-be for indexing)
│   │   │
│   │   └── storage/
│   │       └── client.go                     # Storage service client (upload job attachments)
│   │       └── features_signals.go                    # NEW — emit search signals (promotion/prefs/analytics)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── job_handler.go            # Job HTTP handlers (POST /jobs, GET /jobs/:id, PUT /jobs/:id, DELETE /jobs/:id)
│   │       │   ├── category_handler.go       # Category handlers (GET /categories, GET /categories/tree)
│   │       │   ├── skill_handler.go          # Skill handlers (GET /skills, GET /skills/popular)
│   │       │   ├── invitation_handler.go     # Invitation handlers (POST /jobs/:id/invitations)
│   │       │   ├── search_handler.go         # Search handlers (GET /jobs/search)
│   │       │   ├── saved_search_handler.go   # Saved search handlers (POST /saved-searches)
│   │       │   ├── flag_handler.go           # Flag handlers (POST /jobs/:id/flag)
│   │       │   ├── template_handler.go       # 🆕 Templates (CRUD, clone to new job)
│   │       │   ├── screening_handler.go      # 🆕 Screening & compliance (questions, tests, NDA, export-control)
│   │       │   ├── sourcing_handler.go       # 🆕 Sourcing modes & talent pools
│   │       │   ├── budget_handler.go         # 🆕 Budget controls (ranges, currency, hourly caps)
│   │       │   └── visibility_handler.go     # 🆕 Visibility & lifecycle (schedules, publish, auto-close, extend)
│   │       │   ├── template_versions_handler.go       # NEW — version CRUD; emits version events
│   │       │   ├── eligibility_rules_handler.go       # NEW — set/evaluate rules; returns decision+reasons
│   │       │   ├── hiring_team_handler.go             # NEW — manage roster/roles; owner-only ops
│   │       │   ├── pipeline_handler.go                # NEW — advance/revert; expose SLA breaches
│   │       │   ├── requirements_matrix_handler.go     # NEW — CRUD matrix; score preview
│   │       │   ├── job_analytics_handler.go           # NEW — get snapshot; admin recompute
│   │       │   ├── ab_experiments_handler.go          # NEW — start/stop experiments
│   │       │   ├── moderation_lifecycle_handler.go    # NEW — apply/lift moderation (admin)
│   │       │   ├── syndication_handler.go             # NEW — queue/takedown; partner status
│   │       │   ├── drafts_handler.go                  # NEW — save/restore drafts
│   │       │   ├── geo_requirements_handler.go        # NEW — set/get geo rules; candidate check
│   │       │   ├── duplicate_detection_handler.go     # NEW — find near-dupes
│   │       │   ├── interview_hooks_handler.go         # NEW — set/get interview policy flags
│   │       │   ├── legal_controls_handler.go          # NEW — set legal; place/remove holds
│   │       │   ├── campaign_tags_handler.go           # NEW — add/remove/list tags
│   │       │   ├── retention_rules_handler.go         # NEW — set/get retention; show deadlines
│   │       │   ├── job_promotion_handler.go           # NEW — activate/renew/suspend/expire promotions
│   │       │   └── job_preference_handler.go          # NEW — set/update/get advisory preferences
│   │       │   ├── ai_assist_handler.go              # AI Assist — accept suggestions & optimize descriptions
│   │       │   ├── job_inclusivity_handler.go        # Inclusivity/Accessibility — set/get flags
│   │       │   ├── job_payment_handler.go            # Payment schedule & terms — milestones/hourly/net terms
│   │       │   ├── contract_transition_handler.go    # Contract transition — request/track mapping to contracts-be
│   │       │   ├── fraud_detection_handler.go        # Fraud/Risk — record signals & auto-flag jobs
│   │       │   ├── job_security_handler.go           # Job security — require/verify MFA for high-value posts
│   │       │   ├── job_feedback_handler.go           # Job feedback — submit post-closure feedback
│   │       │   ├── job_esg_handler.go                # ESG/Sustainability — set flags & carbon estimate
│   │       │   ├── job_sharing_handler.go            # Sharing/Referrals — generate tracked links & set incentives
│   │       │   ├── job_tax_handler.go                # Tax — set requirements & request reports
│   │       │   ├── job_health_handler.go             # Job health — schedule checkpoints & send nudges
│   │       │   ├── job_upsell_handler.go             # Upsell — list suggestions & record acceptance
│   │       │   ├── job_previews_handler.go           # Previews (VR/AR) — attach/query previews
│   │       │   ├── bulk_ops_handler.go               # Bulk ops — start update/import batches & check status
│   │       │   ├── webhooks_handler.go               # Webhooks — subscribe/unsubscribe/list job-scoped hooks
│   │       │   ├── job_archive_handler.go            # Archive — archive/reactivate job (user-managed)
│   │       │   ├── job_tools_handler.go              # Job tools — set/get time tracking preferences
│   │       │   ├── job_custom_fields_handler.go      # Custom fields — manage schema & values
│   │       │   ├── data_residency_handler.go         # Data residency — set/get residency policy
│   │       │   └── job_localization_handler.go       # Localization — upsert/remove locales & list variants
│   │       │
│   │       ├── routes/
│   │       │   ├── job_routes.go             # /jobs, /jobs/:id
│   │       │   ├── category_routes.go        # /categories, /categories/tree
│   │       │   ├── skill_routes.go           # /skills, /skills/popular
│   │       │   ├── invitation_routes.go      # /jobs/:id/invitations
│   │       │   ├── search_routes.go          # /jobs/search
│   │       │   ├── saved_search_routes.go    # /saved-searches
│   │       │   ├── flag_routes.go            # /jobs/:id/flag
│   │       │   ├── template_routes.go        # /templates, /templates/:id, /templates/:id/clone
│   │       │   ├── screening_routes.go       # /jobs/:id/screening
│   │       │   ├── sourcing_routes.go        # /jobs/:id/sourcing
│   │       │   ├── budget_routes.go          # /jobs/:id/budget-controls
│   │       │   └── visibility_routes.go      # /jobs/:id/visibility
│   │       |   ├── template_versions_routes.go        # NEW — /templates/:id/versions
│   │       |   ├── eligibility_rules_routes.go        # NEW — /jobs/:id/eligibility
│   │       |   ├── hiring_team_routes.go              # NEW — /jobs/:id/team
│   │       |   ├── pipeline_routes.go                 # NEW — /jobs/:id/pipeline
│   │       |   ├── requirements_matrix_routes.go      # NEW — /jobs/:id/requirements
│   │       |   ├── job_analytics_routes.go            # NEW — /jobs/:id/analytics
│   │       |   ├── ab_experiments_routes.go           # NEW — /jobs/:id/experiments
│   │       |   ├── moderation_lifecycle_routes.go     # NEW — /jobs/:id/moderation
│   │       |   ├── syndication_routes.go              # NEW — /jobs/:id/syndication
│   │       |   ├── drafts_routes.go                   # NEW — /jobs/:id/drafts
│   │       |   ├── geo_requirements_routes.go         # NEW — /jobs/:id/geo
│   │       |   ├── duplicate_detection_routes.go      # NEW — /jobs/:id/duplicates
│   │       |   ├── interview_hooks_routes.go          # NEW — /jobs/:id/interview-policy
│   │       |   ├── legal_controls_routes.go           # NEW — /jobs/:id/legal
│   │       |   ├── campaign_tags_routes.go            # NEW — /jobs/:id/tags
│   │       |   ├── retention_rules_routes.go          # NEW — /jobs/:id/retention
│   │       |   ├── job_promotion_routes.go            # NEW — /jobs/:id/promotion
│   │       |   └── job_preference_routes.go           # NEW — /jobs/:id/preferences
│   │       │   ├── ai_assist_routes.go               # /jobs/:id/ai/accept-suggestions, /jobs/:id/ai/optimize
│   │       │   ├── job_inclusivity_routes.go         # /jobs/:id/inclusivity
│   │       │   ├── job_payment_routes.go             # /jobs/:id/payment
│   │       │   ├── contract_transition_routes.go     # /jobs/:id/contract/transition, /jobs/:id/contract/transition/status
│   │       │   ├── fraud_detection_routes.go         # /jobs/:id/fraud/signal, /jobs/:id/fraud/autoflag
│   │       │   ├── job_security_routes.go            # /jobs/:id/security/require-mfa, /jobs/:id/security/verify
│   │       │   ├── job_feedback_routes.go            # /jobs/:id/feedback
│   │       │   ├── job_esg_routes.go                 # /jobs/:id/esg, /jobs/:id/esg/estimate
│   │       │   ├── job_sharing_routes.go             # /jobs/:id/sharing/link, /jobs/:id/sharing/incentive
│   │       │   ├── job_tax_routes.go                 # /jobs/:id/tax, /jobs/:id/tax/report
│   │       │   ├── job_health_routes.go              # /jobs/:id/health/checkpoint, /jobs/:id/health/nudge
│   │       │   ├── job_upsell_routes.go              # /jobs/:id/upsell, /jobs/:id/upsell/accept
│   │       │   ├── job_previews_routes.go            # /jobs/:id/previews/vr, /jobs/:id/previews/ar
│   │       │   ├── bulk_ops_routes.go                # /jobs/bulk/update, /jobs/bulk/import, /jobs/bulk/:batchId
│   │       │   ├── webhooks_routes.go                # /jobs/:id/webhooks (POST/DELETE/GET)
│   │       │   ├── job_archive_routes.go             # /jobs/:id/archive, /jobs/:id/reactivate
│   │       │   ├── job_tools_routes.go               # /jobs/:id/tools/time-tracking
│   │       │   ├── job_custom_fields_routes.go       # /jobs/:id/custom-fields (schema+values)
│   │       │   ├── data_residency_routes.go          # /jobs/:id/residency
│   │       │   └── job_localization_routes.go        # /jobs/:id/locales (PUT/DELETE/GET)
│   │       │
│   │       ├── middleware/                   # 📝 UPDATED: Uses platform-shared middleware
│   │       │   ├── auth.go                   # 📝 UPDATED: Uses pkg/auth
│   │       │   ├── rbac.go                   # 📝 UPDATED: Uses pkg/auth authorizer
│   │       │   ├── cors.go                   # 📝 UPDATED: Uses platform-shared/ginx
│   │       │   ├── rate_limit.go             # Rate limiting
│   │       │   └── idempotency.go            # 🆕 Uses platform-shared/idempotency
│   │       │
│   │       ├── responses/                    # 📝 UPDATED: Response helpers
│   │       │   └── README.md                 # 📝 Points to platform-shared/httpx
│   │       │
│   │       └── router.go                     # 📝 UPDATED: Uses platform-shared/ginx middleware
│   │
│   └── config/                               # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                         # 🆕 Typed Config struct
│       ├── loader.go                         # 🆕 Config loader using Viper
│       └── docs/
│           └── CONFIGURATION.md              # 🆕 Configuration documentation
│
├── config/                                   # 🆕 MEDIUM - Configuration files
│   ├── default.yaml                          # 🆕 Default configuration
│   ├── dev.yaml                              # 🆕 Development overrides
│   └── prod.yaml                             # 🆕 Production overrides
│
├── dapr/                                     # 🆕 MEDIUM - Dapr components
│   ├── local/                                # 🆕 For dapr run
│   │   ├── pubsub.yaml                       # Kafka pub/sub component
│   │   └── statestore.yaml                   # State store component
│   └── k8s/                                  # 🆕 For kubectl apply
│       ├── pubsub.yaml                       # Kafka with scopes: ["jobs-be"]
│       ├── statestore.yaml                   # State store with scopes
│       └── secrets.yaml                      # Dapr secret store
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                         # Service-specific errors
│   │   └── codes.go                          # Error codes (JOB_NOT_FOUND, INVALID_BUDGET)
│   ├── logger/                               # ❌ REMOVED: Use platform-shared/logging
│   │   └── README.md                         # 📝 Points to platform-shared/logging
│   ├── utils/
│   │   ├── validator.go                      # Local validation utilities (shared helpers, regexes, ISO currency codes, etc.)
│   │   ├── slug.go                           # Generate slugs
│   │   └── sanitizer.go                      # Sanitize input
│   └── constants/
│       ├── events.go                         # ❌ REMOVED: Use contracts/events
│       └── topics.go                         # ❌ REMOVED: Use contracts/events
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
│   ├── seed-categories.sh                    # Seed job categories
│   └── seed-skills.sh                        # Seed skills taxonomy
│
├── tests/
│   ├── unit/                                 # Unit tests
│   ├── integration/                          # Integration tests
│   └── e2e/                                  # End-to-end tests
│
├── docs/
│   ├── README.md                             # Service overview
│   ├── API.md                                # API documentation
│   ├── EVENTS.md                             # 🆕 Events (published: job.posted, job.closed, consumed: user.verified, proposal.accepted)
│   ├── ARCHITECTURE.md                       # Architecture details
│   ├── MIGRATIONS.md                         # 🆕 Migration history
│   ├── SCHEMA.md                             # 🆕 Database schema
│   ├── RUNBOOK.md                            # 🆕 Operational procedures
│   ├── categories.md                         # Job categories documentation
│   └── skills-taxonomy.md                    # Skills taxonomy documentation
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

## 📦 **3️⃣ proposals-be (Freelancer Proposal Submission Service)**

```
apps/be/proposals-be/
│
├── cmd/
│   └── api/
│       └── main.go                           # Entry point - initializes Gin, Dapr, connects to Postgres
│
├── internal/
│   ├── domain/                               # 🏛️ Domain Layer (Core Business Logic - Load Third)
│   │   │
│   │   # ========================= CORE PROPOSAL LIFECYCLE =========================
│   │   ├── proposal/
│   │   │   ├── entity.go                     # Proposal aggregate root (ID, JobID, FreelancerID, CoverLetter, Amount, Duration)
│   │   │   ├── value_objects.go              # ProposalAmount, ProposalDuration
│   │   │   ├── enums.go                      # Status (Draft, Submitted, Accepted, Rejected, Withdrawn), Type (Hourly, Fixed)
│   │   │   ├── errors.go                     # Domain errors (ProposalNotFound, InvalidAmount, AlreadySubmitted)
│   │   │   ├── repository.go                 # ProposalRepository interface
│   │   │   ├── schema.go                     # 🆕 Structured schema (cover letter, Q&A, milestones, delivery estimates)
│   │   │   └── events.go                     # 🆕 ProposalDrafted, ProposalSubmitted, ProposalEdited, ProposalWithdrawn, ProposalAccepted, ProposalRejected
│   │   │
│   │   ├── cover_letter/
│   │   │   ├── entity.go                     # Cover letter content (ProposalID, Content, Version)
│   │   │   ├── errors.go                     # Cover letter errors
│   │   │   ├── repository.go                 # CoverLetterRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: CoverLetterAdded, CoverLetterEdited, CoverLetterVersionReverted
│   │   │
│   │   ├── proposal_attachment/
│   │   │   ├── entity.go                     # Proposal attachments (ProposalID, FileURL, FileType)
│   │   │   ├── errors.go                     # Attachment errors
│   │   │   ├── repository.go                 # AttachmentRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: ProposalAttachmentAdded, ProposalAttachmentRemoved
│   │   │
│   │   ├── proposal_question_answer/
│   │   │   ├── entity.go                     # Answers to job questions (ProposalID, QuestionID, Answer)
│   │   │   ├── errors.go                     # Answer errors
│   │   │   ├── repository.go                 # AnswerRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: ProposalAnswerSubmitted, ProposalAnswerUpdated, ProposalAnswerRemoved
│   │   │
│   │   ├── milestone/
│   │   │   ├── entity.go                     # Proposal milestones (ProposalID, Description, Amount, DueDate)
│   │   │   ├── milestone_group.go            # Group milestones (for fixed-price projects)
│   │   │   ├── errors.go                     # Milestone errors
│   │   │   ├── repository.go                 # MilestoneRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: MilestoneCreated, MilestoneUpdated, MilestoneReordered, MilestoneRemoved
│   │   │
│   │   # ========================= BIDDING SYSTEM =========================
│   │   ├── bid/
│   │   │   ├── entity.go                     # Bidding system (ProposalID, Amount, Rank, PlacedAt)
│   │   │   ├── bid_amount.go                 # Bid amount management (validation, currency conversion)
│   │   │   ├── bid_history.go                # Bid history tracking (all bids for a job)
│   │   │   ├── auto_bid.go                   # Auto-bidding logic (automatic bid adjustments)
│   │   │   ├── bid_rank.go                   # Bid ranking (1st, 2nd, 3rd lowest bid)
│   │   │   ├── errors.go                     # Bid errors
│   │   │   ├── repository.go                 # BidRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: BidPlaced, BidUpdated, BidAutoAdjusted, BidRetracted, BidRankChanged
│   │   │
│   │   ├── bid_strategy/
│   │   │   ├── entity.go                     # Bidding strategies (Aggressive, Conservative, Auto)
│   │   │   ├── aggressive.go                 # Aggressive bidding strategy (undercut by X%)
│   │   │   ├── conservative.go               # Conservative bidding (competitive but safe)
│   │   │   ├── auto_strategy.go              # Auto-bid strategy (ML-based optimal bidding)
│   │   │   ├── errors.go                     # Strategy errors
│   │   │   ├── repository.go                 # BidStrategyRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: BidStrategyCreated, BidStrategyUpdated, BidStrategyActivated, BidStrategyDeactivated
│   │   │
│   │   ├── bid_notification/
│   │   │   ├── entity.go                     # Bid notifications (UserID, Message, Type, SentAt)
│   │   │   ├── outbid_alert.go               # Outbid alerts (notify when outbid)
│   │   │   ├── errors.go                     # Notification errors
│   │   │   ├── repository.go                 # BidNotificationRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: OutbidAlertTriggered, BidNotificationQueued, BidNotificationSent, BidNotificationFailed
│   │   │
│   │   ├── auction/                          # 🆕 Auction mechanics for boosted proposals
│   │   │   ├── entity.go                     # Auction (JobID, Bids[], CurrentTop, AuctionEnd)
│   │   │   ├── errors.go                     # Domain errors (ErrBidTooLow, ErrAuctionClosed)
│   │   │   ├── repository.go                 # AuctionRepository interface (PlaceBid, GetStatus, EndAuction)
│   │   │   └── events.go                     # Domain events: AuctionBidPlaced, AuctionEnded, TopPositionChanged
│   │   │
│   │   ├── bid_anomaly_detection/            # 🆕 Detect anomalous bidding behavior
│   │   │   ├── entity.go                     # Anomaly(score, reasons), ReviewState
│   │   │   ├── enums.go                      # ReviewState: Open, UnderReview, Closed
│   │   │   ├── errors.go                     # ErrThresholdNotMet, ErrReviewStateInvalid
│   │   │   ├── repository.go                 # Detect, UpdateReview, ListByJob
│   │   │   └── events.go                     # bid.anomaly.detected, bid.anomaly.review.updated
│   │   │
│   │   # ========================= CONNECTS & BOOST =========================
│   │   ├── connect/
│   │   │   ├── entity.go                     # Connect usage tracking (FreelancerID, JobID, ConnectsUsed, UsedAt)
│   │   │   ├── transaction.go                # Connect transactions (purchase, use, refund)
│   │   │   ├── errors.go                     # Connect errors (InsufficientConnects)
│   │   │   ├── repository.go                 # ConnectRepository interface
│   │   │   ├── tiers.go                      # 🆕 Credit cost tiers by job popularity (tier rules)
│   │   │   └── events.go                     # 🆕 Domain events: ConnectsDebited, ConnectsCredited, ConnectsRefunded, ConnectTierChanged
│   │   │
│   │   ├── connect_refund/                   # 🆕 Connects refund logic
│   │   │   ├── entity.go                     # ConnectRefund (ProposalID, ConnectsRefunded, Reason, RefundedAt)
│   │   │   ├── reason.go                     # Reason: JobClosed, InvalidJob, Spam
│   │   │   ├── errors.go                     # Domain errors (ErrNotEligibleForRefund)
│   │   │   ├── repository.go                 # ConnectRefundRepository (RequestRefund, ProcessRefund, GetByProposalID)
│   │   │   └── events.go                     # Events: RefundRequested, RefundProcessed, RefundDenied
│   │   │
│   │   ├── boost/
│   │   │   ├── entity.go                     # Boosted proposals (ProposalID, BoostLevel, ExpiresAt)
│   │   │   ├── boost_level.go                # Boost levels (Basic, Premium, Elite)
│   │   │   ├── errors.go                     # Boost errors
│   │   │   ├── repository.go                 # BoostRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: ProposalBoostPurchased, ProposalBoostUpgraded, ProposalBoostExpired
│   │   │
│   │   # ========================= TEMPLATES & RATE CARDS =========================
│   │   ├── template/
│   │   │   ├── entity.go                     # Proposal templates (FreelancerID, Title, Content)
│   │   │   ├── template_category.go          # Template categories (by job type, industry)
│   │   │   ├── errors.go                     # Template errors
│   │   │   ├── repository.go                 # TemplateRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: ProposalTemplateCreated, ProposalTemplateUpdated, ProposalTemplateArchived
│   │   │
│   │   ├── rate_card/                        # 🆕 Rate cards (saved packages: starter / standard / premium)
│   │   │   ├── entity.go                     # RateCard (FreelancerID, PackageName, Price, ScopeSummary, DeliveryEstimate, Inclusions)
│   │   │   ├── package.go                    # Package tiers (Starter, Standard, Premium)
│   │   │   ├── errors.go                     # Rate card errors
│   │   │   ├── repository.go                 # RateCardRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: RateCardCreated, RateCardUpdated, RateCardArchived
│   │   │
│   │   # ========================= ANALYTICS & TRACKING =========================
│   │   ├── proposal_analytics/
│   │   │   ├── entity.go                     # Proposal performance analytics (Views, ClickThroughRate, ResponseRate)
│   │   │   ├── view_tracking.go              # Track proposal views
│   │   │   ├── errors.go                     # Analytics errors
│   │   │   ├── repository.go                 # AnalyticsRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: ProposalViewed, ProposalEngagementUpdated, ProposalResponseRateUpdated
│   │   │
│   │   ├── response_tracker/                 # 🆕 Track client responses to proposals
│   │   │   ├── entity.go                     # ResponseTracker (ProposalID, ViewedAt, RespondedAt, ResponseType, Notes)
│   │   │   ├── errors.go                     # Domain errors (ErrNoResponse)
│   │   │   ├── repository.go                 # ResponseTrackerRepository interface (TrackView, TrackResponse, GetStats)
│   │   │   └── events.go                     # Domain events: ProposalViewedByClient, ClientResponded
│   │   │
│   │   ├── insight/                          # 🆕 Proposal insights (competing bids, stats, engagement)
│   │   │   ├── entity.go                     # Insight aggregate (ProposalID, CompetingBidsCount, AvgBidAmount, ClientEngagementRate, FreelancerStats)
│   │   │   ├── errors.go                     # Domain errors (ErrInsightNotAvailable, ErrCalculationFailed)
│   │   │   ├── repository.go                 # InsightRepository interface (Upsert, GetByProposalID, Recalc)
│   │   │   └── events.go                     # Domain events: InsightUpdated, InsightAccessed
│   │   │
│   │   ├── ranking/                          # 🆕 Visibility scoring
│   │   │   ├── entity.go                     # Ranking (ProposalID, Score, Position, Factors)
│   │   │   ├── factors.go                    # Factors (Relevance, Boost, CTR, ResponseTime, ProfileStrength)
│   │   │   ├── errors.go                     # Domain errors (ErrRankComputeFailed)
│   │   │   ├── repository.go                 # RankingRepository (Upsert, GetByProposalID, Recompute)
│   │   │   └── events.go                     # Events: ProposalRanked, RankingUpdated
│   │   │
│   │   # ========================= COLLABORATION & WORKFLOW =========================
│   │   ├── negotiation/                      # 🆕 Negotiation thread (counter-offers before contract creation)
│   │   │   ├── entity.go                     # Negotiation (ProposalID, CounterRate, CounterScope, CounterMilestones, Notes, State)
│   │   │   ├── state.go                      # States (Open, Countered, Accepted, Declined, Expired)
│   │   │   ├── errors.go                     # Negotiation errors (InvalidState, NotAllowed, StaleCounter)
│   │   │   ├── repository.go                 # NegotiationRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: NegotiationOpened, NegotiationCountered, NegotiationAccepted, NegotiationDeclined, NegotiationExpired
│   │   │
│   │   ├── invite/                           # 🆕 Invite flow (client invites; pre-filled proposals; decline with reason)
│   │   │   ├── entity.go                     # Invite (JobID, ClientID, FreelancerID, PrefilledProposalRef, Message, Status, DeclineReason)
│   │   │   ├── status.go                     # InviteStatus (Pending, Accepted, Declined, Expired)
│   │   │   ├── errors.go                     # Invite errors
│   │   │   ├── repository.go                 # InviteRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: InviteSent, InviteAccepted, InviteDeclined, InviteExpired
│   │   │
│   │   ├── revision/                         # 🆕 Post-submit versioning
│   │   │   ├── entity.go                     # Revision entity (ProposalID, Version, Changes, RevisedAt, Reason, State)
│   │   │   ├── errors.go                     # Domain errors (ErrRevisionNotAllowed, ErrAlreadyFinal)
│   │   │   ├── repository.go                 # RevisionRepository (Create, GetHistory, GetLatest, Approve, Reject)
│   │   │   └── events.go                     # Domain events: ProposalRevised, RevisionApproved, RevisionRejected
│   │   │
│   │   ├── collaboration/                    # 🆕 Collaborative editing (multi-user)
│   │   │   ├── entity.go                     # Collaboration (ProposalID, Collaborators[], Permissions, LastEdit)
│   │   │   ├── errors.go                     # Domain errors (ErrAccessDenied, ErrCollaboratorLimitExceeded)
│   │   │   ├── repository.go                 # CollaborationRepository (AddCollaborator, RemoveCollaborator, GetCollaborators)
│   │   │   └── events.go                     # Events: CollaboratorAdded, CollaboratorRemoved, ProposalCollaborated
│   │   │
│   │   ├── approval_workflow/                # 🆕 Multi-step approvals for proposals
│   │   │   ├── entity.go                     # ApprovalFlow, Step(role, approver_ids, sla), State
│   │   │   ├── enums.go                      # Roles, StepResult, SLAOutcome
│   │   │   ├── errors.go                     # ErrApproverRequired, ErrStepMismatch, ErrSLAExceeded, ErrInvalidTransition
│   │   │   ├── repository.go                 # Start, Advance, Reject, GetByProposalID
│   │   │   └── events.go                     # approval.requested, approval.granted, approval.rejected, approval.sla.warning
│   │   │
│   │   ├── document_redlining/               # 🆕 Redline threads for shared docs
│   │   │   ├── entity.go                     # RedlineThread, DiffRefs, Resolution
│   │   │   ├── errors.go                     # ErrThreadLocked, ErrDiffNotFound
│   │   │   ├── repository.go                 # StartThread, AppendDiff, ResolveThread, GetThread
│   │   │   └── events.go                     # redline.thread.started, redline.thread.resolved
│   │   │
│   │   # ========================= CLIENT INTERACTION =========================
│   │   ├── interview_request/                # 🆕 Interviews on proposals
│   │   │   ├── entity.go                     # InterviewRequest (ProposalID, RequesterID, ScheduleOptions, Notes, Status)
│   │   │   ├── status.go                     # Status: Pending, Scheduled, Completed, Declined
│   │   │   ├── errors.go                     # Domain errors (ErrSlotUnavailable, ErrNotAuthorized)
│   │   │   ├── repository.go                 # InterviewRequestRepository (Create, UpdateStatus, GetByProposalID, List)
│   │   │   └── events.go                     # Events: InterviewRequested, InterviewScheduled, InterviewCompleted, InterviewDeclined
│   │   │
│   │   ├── feedback/                         # 🆕 Client feedback pre-contract
│   │   │   ├── entity.go                     # Feedback (ProposalID, ClientID, Rating, Comments, GivenAt)
│   │   │   ├── errors.go                     # Domain errors (ErrDuplicateFeedback, ErrInvalidRating)
│   │   │   ├── repository.go                 # FeedbackRepository (Create, GetByProposalID, Respond)
│   │   │   └── events.go                     # Events: ProposalFeedbackGiven, FeedbackResponded
│   │   │
│   │   ├── shortlist/                        # 🆕 Client shortlisting of proposals
│   │   │   ├── entity.go                     # Shortlist (JobID, ProposalID, ClientID, Notes, ShortlistedAt, Status)
│   │   │   ├── status.go                     # Status: Shortlisted, Interviewed, Hired, Rejected
│   │   │   ├── errors.go                     # Domain errors (ErrAlreadyShortlisted, ErrProposalNotFound)
│   │   │   ├── repository.go                 # ShortlistRepository (Add, Remove, GetByJobID, UpdateStatus)
│   │   │   └── events.go                     # Events: ProposalShortlisted, ShortlistUpdated, ShortlistRemoved
│   │   │
│   │   ├── client_insight/                   # 🆕 Client history insights for freelancers
│   │   │   ├── entity.go                     # ClientInsight (ClientID, HireRate, AvgBudget, FeedbackScore, TotalSpend, RecentJobs[])
│   │   │   ├── errors.go                     # Domain errors (ErrInsightUnavailable)
│   │   │   ├── repository.go                 # ClientInsightRepository (GetByClientID, Refresh)
│   │   │   └── events.go                     # Events: ClientInsightAccessed, ClientInsightUpdated
│   │   │
│   │   ├── client_preference/                # 🆕 Infer and store client preferences for future proposals
│   │   │   ├── entity.go                     # ClientPreference (ClientID, PreferredStyles[], BudgetRange, ResponseTime, Keywords[])
│   │   │   ├── errors.go                     # Domain errors (ErrPreferenceNotFound)
│   │   │   ├── repository.go                 # ClientPreferenceRepository interface (Infer, Update, GetByClientID)
│   │   │   └── events.go                     # Domain events: PreferenceInferred, PreferenceUpdated
│   │   │
│   │   # ========================= COMPLIANCE & LEGAL =========================
│   │   ├── compliance/                       # 🆕 Legal/policy checks (pre-contract)
│   │   │   ├── entity.go                     # ComplianceCheck (ProposalID, ChecksPerformed[], Results, Status, Score)
│   │   │   ├── check.go                      # Check defs (Plagiarism, Spam, Legal, PII, ToS)
│   │   │   ├── errors.go                     # Domain errors (ErrCheckFailed, ErrEvidenceMissing)
│   │   │   ├── repository.go                 # ComplianceRepository (RecordCheck, GetLatest, Resolve)
│   │   │   └── events.go                     # Events: ComplianceChecked, CompliancePassed, ComplianceFailed, ComplianceResolved
│   │   │
│   │   ├── proposal_flag/
│   │   │   ├── entity.go                     # Flagged proposals (ProposalID, ReporterID, Reason, Status)
│   │   │   ├── reason.go                     # Flag reasons (Spam, Plagiarism, Inappropriate)
│   │   │   ├── status.go                     # Flag status
│   │   │   ├── errors.go                     # Flag errors
│   │   │   ├── repository.go                 # FlagRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: ProposalFlagSubmitted, ProposalFlagResolved, ProposalFlagDismissed
│   │   │
│   │   ├── spam_detection/                   # 🆕 Anti-spam measures
│   │   │   ├── entity.go                     # SpamDetection (ProposalID, Score, Flags[], Status)
│   │   │   ├── errors.go                     # Domain errors (ErrSpamDetected)
│   │   │   ├── repository.go                 # SpamDetectionRepository (Scan, Flag, ClearFlags)
│   │   │   └── events.go                     # Events: SpamScanned, SpamFlagged, SpamCleared
│   │   │
│   │   ├── duplicate_check/                  # 🆕 Detect duplicate proposals
│   │   │   ├── entity.go                     # DuplicateCheck (ProposalID, SimilarProposals[], SimilarityScore)
│   │   │   ├── errors.go                     # Domain errors (ErrDuplicateDetected)
│   │   │   ├── repository.go                 # DuplicateCheckRepository (CheckDuplicates, MarkAsUnique)
│   │   │   └── events.go                     # Events: DuplicateDetected, DuplicateResolved
│   │   │
│   │   ├── security_questionnaire/           # 🆕 Security questionnaires attached to proposals
│   │   │   ├── entity.go                     # Questionnaire(type, status, files[])
│   │   │   ├── enums.go                      # QuestionnaireType, QuestionnaireStatus
│   │   │   ├── errors.go                     # ErrQuestionnaireRejected, ErrUnsupportedType
│   │   │   ├── repository.go                 # Issue, Complete, Reject, GetByProposalID
│   │   │   └── events.go                     # questionnaire.issued, questionnaire.completed, questionnaire.rejected
│   │   │
│   │   ├── legal_hold_retention/             # 🆕 Legal holds & retention policies
│   │   │   ├── entity.go                     # Hold(policy_id, scope), RetentionPolicy
│   │   │   ├── errors.go                     # ErrHoldConflict, ErrRetentionOverlap
│   │   │   ├── repository.go                 # ApplyHold, LiftHold, SetRetention, GetStatus
│   │   │   └── events.go                     # legal.hold.applied, legal.hold.lifted, retention.policy.set
│   │   │
│   │   ├── data_residency/                   # 🆕 Data residency constraints on proposal data
│   │   │   ├── entity.go                     # DataZone, ResidencyRule, Violation
│   │   │   ├── errors.go                     # ErrZoneUnknown, ErrRuleViolation
│   │   │   ├── repository.go                 # SetRule, ListRules, RecordViolation
│   │   │   └── events.go                     # residency.rule.set, residency.violation.detected
│   │   │
│   │   ├── anonymization/                    # 🆕 PII redaction & anonymized variants
│   │   │   ├── entity.go                     # RedactionRules, AnonymizedVariant
│   │   │   ├── errors.go                     # ErrRedactionInvalid, ErrVariantNotFound
│   │   │   ├── repository.go                 # ApplyRules, CreateVariant, GetVariant
│   │   │   └── events.go                     # anonymization.rules.applied, anonymization.variant.created
│   │   │
│   │   # ========================= CONTRACT & TERMS =========================
│   │   ├── terms/                            # 🆕 Custom terms in proposals
│   │   │   ├── entity.go                     # Terms (ProposalID, CustomClauses[], PaymentTerms, NDA, IPAssignment, Version)
│   │   │   ├── errors.go                     # Domain errors (ErrInvalidTerms, ErrTermsNotAccepted)
│   │   │   ├── repository.go                 # TermsRepository (Create, Update, GetByProposalID, Accept)
│   │   │   └── events.go                     # Events: TermsAdded, TermsUpdated, TermsAccepted, TermsRejected
│   │   │
│   │   ├── nda/                              # 🆕 NDA integration
│   │   │   ├── entity.go                     # NDA (ProposalID, TemplateID, SignedBy[], SignedAt, Status)
│   │   │   ├── status.go                     # Status: Pending, Signed, Rejected
│   │   │   ├── errors.go                     # Domain errors (ErrNDANotSigned, ErrTemplateNotFound)
│   │   │   ├── repository.go                 # NDARepository (ProposeNDA, SignNDA, GetByProposalID)
│   │   │   └── events.go                     # Events: NDAProposed, NDASigned, NDARejected
│   │   │
│   │   ├── ip_rights/                        # 🆕 Intellectual property rights
│   │   │   ├── entity.go                     # IPRights (ProposalID, AssignmentType, Clauses, AgreedBy[], AgreedAt)
│   │   │   ├── errors.go                     # Domain errors (ErrIPNotAgreed)
│   │   │   ├── repository.go                 # IPRightsRepository (SetRights, Agree, GetByProposalID)
│   │   │   └── events.go                     # Events: IPRightsSet, IPRightsAgreed
│   │   │
│   │   ├── payment_terms/                    # 🆕 Detailed payment terms
│   │   │   ├── entity.go                     # PaymentTerms (ProposalID, Schedule, Installments[], Fees, Currency, TermsText)
│   │   │   ├── errors.go                     # Domain errors (ErrInvalidSchedule, ErrFeeMismatch)
│   │   │   ├── repository.go                 # PaymentTermsRepository (Create, Update, GetByProposalID)
│   │   │   └── events.go                     # Events: PaymentTermsSet, PaymentTermsUpdated
│   │   │
│   │   ├── pricing_model/                    # 🆕 Advanced pricing (tiered, dynamic)
│   │   │   ├── entity.go                     # PricingModel (ProposalID, Type, Tiers[], DynamicFactors, Currency)
│   │   │   ├── type.go                       # Pricing types (Tiered, Subscription, PerformanceBased)
│   │   │   ├── errors.go                     # Domain errors (ErrInvalidPricing, ErrTierOverlap)
│   │   │   ├── repository.go                 # PricingModelRepository (Create, Update, GetByProposalID)
│   │   │   └── events.go                     # Events: PricingModelUpdated, TierAdded, DynamicPriceAdjusted
│   │   │
│   │   # ========================= ENTERPRISE & PROCUREMENT =========================
│   │   ├── rfp/                              # 🆕 RFP import & compliance mapping
│   │   │   ├── entity.go                     # RFPImport, ComplianceMatrix, SectionMap
│   │   │   ├── errors.go                     # ErrInvalidRFPFormat, ErrSectionMapConflict
│   │   │   ├── repository.go                 # Import, MapSections, GetCompliance
│   │   │   └── events.go                     # rfp.imported, rfp.compliance.completed
│   │   │
│   │   ├── quote/                            # 🆕 Quote generation & lifecycle
│   │   │   ├── entity.go                     # Quote, QuoteLineItem, Taxes, Discounts, FxLock, ValidUntil
│   │   │   ├── enums.go                      # QuoteStatus: Draft, Issued, Repriced, Expired
│   │   │   ├── errors.go                     # ErrInvalidLineItem, ErrFxLockExpired
│   │   │   ├── repository.go                 # Issue, Reprice, GetByProposalID, Expire
│   │   │   └── events.go                     # quote.issued, quote.repriced, quote.expired
│   │   │
│   │   ├── procurement/                      # 🆕 Enterprise procurement links
│   │   │   ├── entity.go                     # PORef, CostCenter, BudgetCheck(status, evidence)
│   │   │   ├── errors.go                     # ErrPOInvalid, ErrBudgetCheckFailed
│   │   │   ├── repository.go                 # AttachPO, RunBudgetCheck, GetStatus
│   │   │   └── events.go                     # procurement.po.attached, procurement.budget.check.requested/satisfied/failed
│   │   │
│   │   ├── evaluation_rubric/                # 🆕 Rubrics for proposal evaluation
│   │   │   ├── entity.go                     # Rubric(criteria[], weights), RubricScore
│   │   │   ├── errors.go                     # ErrWeightNormalization, ErrCriterionMissing
│   │   │   ├── repository.go                 # CreateRubric, ScoreProposal, GetRubric, GetScore
│   │   │   └── events.go                     # rubric.created, rubric.scored
│   │   │
│   │   # ========================= AI & OPTIMIZATION =========================
│   │   ├── ai_assist/                        # 🆕 AI assistance (generation/suggestions)
│   │   │   ├── entity.go                     # AIAssist aggregate (ProposalID, GeneratedContent, Score, Suggestions[], ModelRef)
│   │   │   ├── generator.go                  # Provider abstraction (+ guardrails hooks)
│   │   │   ├── errors.go                     # Domain errors (ErrGenerationFailed, ErrGuardrailViolated)
│   │   │   ├── repository.go                 # AIAssistRepository interface (SaveDraft, GetLatest, LogUsage)
│   │   │   └── events.go                     # Domain events: AISuggestionGenerated, AIProposalCreated, AIAssistUsed
│   │   │
│   │   ├── keyword_optimization/             # 🆕 Keyword scanning and optimization for proposals
│   │   │   ├── entity.go                     # KeywordOpt (ProposalID, SuggestedKeywords[], MatchScore, OptimizedContent)
│   │   │   ├── errors.go                     # Domain errors (ErrLowMatchScore)
│   │   │   ├── repository.go                 # KeywordOptRepository interface (Optimize, Scan, GetSuggestions)
│   │   │   └── events.go                     # Domain events: KeywordsOptimized, OptimizationApplied
│   │   │
│   │   ├── success_predictor/                # 🆕 Predict proposal success based on historical data
│   │   │   ├── entity.go                     # SuccessPredictor (ProposalID, Probability, Factors[], Recommendations)
│   │   │   ├── errors.go                     # Domain errors (ErrInsufficientData)
│   │   │   ├── repository.go                 # SuccessPredictorRepository interface (Predict, UpdateModel, GetPrediction)
│   │   │   └── events.go                     # Domain events: SuccessPredicted, PredictionUpdated
│   │   │
│   │   ├── best_practices/                   # 🆕 Embed best practices and tips for proposals
│   │   │   ├── entity.go                     # BestPractice (Category, Tips[], ComplianceRules)
│   │   │   ├── errors.go                     # Domain errors (ErrPracticeNotFound)
│   │   │   ├── repository.go                 # BestPracticeRepository interface (GetTips, ApplyToProposal)
│   │   │   └── events.go                     # Domain events: BestPracticeApplied, TipsViewed
│   │   │
│   │   ├── grammar_check/                    # 🆕 Grammar and spell checking for proposals
│   │   │   ├── entity.go                     # GrammarCheck (ProposalID, Issues[], Suggestions, Score, CheckedAt)
│   │   │   ├── errors.go                     # Domain errors (ErrCheckFailed, ErrHighErrorRate)
│   │   │   ├── repository.go                 # GrammarCheckRepository interface (Check, GetIssues, ApplySuggestions)
│   │   │   └── events.go                     # Domain events: GrammarChecked, IssuesResolved
│   │   │
│   │   ├── personalization/                  # 🆕 Personalization engine for proposals
│   │   │   ├── entity.go                     # Personalization (ProposalID, ClientPreferences, CustomizedContent, Score)
│   │   │   ├── errors.go                     # Domain errors (ErrPersonalizationFailed)
│   │   │   ├── repository.go                 # PersonalizationRepository interface (Personalize, GetCustomized, UpdateScore)
│   │   │   └── events.go                     # Domain events: ProposalPersonalized, PersonalizationApplied
│   │   │
│   │   ├── ab_testing/                       # 🆕 A/B testing for proposal templates/strategies
│   │   │   ├── entity.go                     # ABTest (TestID, Variants[], Metrics, StartDate, EndDate, Winner)
│   │   │   ├── errors.go                     # Domain errors (ErrTestInProgress, ErrNoWinner)
│   │   │   ├── repository.go                 # ABTestRepository interface (CreateTest, RunTest, GetResults, DeclareWinner)
│   │   │   └── events.go                     # Domain events: ABTestStarted, ABTestEnded, VariantPerformed
│   │   │
│   │   # ========================= PROFILE & PORTFOLIO =========================
│   │   ├── portfolio_link/                   # 🆕 Link portfolio items to proposals
│   │   │   ├── entity.go                     # PortfolioLink (ProposalID, PortfolioItemID, RelevanceNote, DisplayOrder)
│   │   │   ├── errors.go                     # Domain errors (ErrPortfolioNotFound, ErrDuplicateLink)
│   │   │   ├── repository.go                 # PortfolioLinkRepository (Link, Unlink, ListByProposalID)
│   │   │   └── events.go                     # Events: PortfolioLinked, PortfolioUnlinked
│   │   │
│   │   ├── skill_match/                      # 🆕 Match proposal skills to job requirements
│   │   │   ├── entity.go                     # SkillMatch (ProposalID, JobSkills[], MatchedSkills[], Score, Gaps[])
│   │   │   ├── errors.go                     # Domain errors (ErrNoSkillsMatched)
│   │   │   ├── repository.go                 # SkillMatchRepository (ComputeMatch, GetMatchByProposalID)
│   │   │   └── events.go                     # Events: SkillsMatched, MatchScoreUpdated
│   │   │
│   │   ├── video_introduction/               # 🆕 Video intro attachments
│   │   │   ├── entity.go                     # VideoIntroduction (ProposalID, VideoURL, Transcript, Duration, UploadedAt)
│   │   │   ├── errors.go                     # Domain errors (ErrVideoTooLong, ErrInvalidFormat)
│   │   │   ├── repository.go                 # VideoIntroductionRepository (Upload, GetByProposalID, Transcribe)
│   │   │   └── events.go                     # Events: VideoUploaded, VideoTranscribed
│   │   │
│   │   ├── reference/                        # 🆕 References/testimonials
│   │   │   ├── entity.go                     # Reference (ProposalID, RefereeID, Testimonial, Verified, AddedAt)
│   │   │   ├── errors.go                     # Domain errors (ErrReferenceNotVerified)
│   │   │   ├── repository.go                 # ReferenceRepository (Add, Verify, ListByProposalID)
│   │   │   └── events.go                     # Events: ReferenceAdded, ReferenceVerified
│   │   │
│   │   ├── membership_perk/                  # 🆕 Perks based on membership level
│   │   │   ├── entity.go                     # MembershipPerk (FreelancerID, PerkType, AppliedToProposalID, UsageCount)
│   │   │   ├── type.go                       # Perk types (ExtraConnects, PriorityReview, AdvancedAnalytics)
│   │   │   ├── errors.go                     # Domain errors (ErrPerkLimitExceeded)
│   │   │   ├── repository.go                 # MembershipPerkRepository (ApplyPerk, CheckEligibility, GetUsage)
│   │   │   └── events.go                     # Events: PerkApplied, PerkExhausted
│   │   │
│   │   # ========================= LIFECYCLE MANAGEMENT =========================
│   │   ├── expiration/                       # 🆕 Proposal expiration handling
│   │   │   ├── entity.go                     # Expiration (ProposalID, ExpiresAt, Reason, AutoRenew, NotifiedAt)
│   │   │   ├── errors.go                     # Domain errors (ErrAlreadyExpired, ErrInvalidExpiryDate)
│   │   │   ├── repository.go                 # ExpirationRepository (SetExpiry, Extend, CheckExpired, Notify)
│   │   │   └── events.go                     # Events: ProposalExpired, ExpiryExtended, ExpiryNotified
│   │   │
│   │   ├── withdrawal/                       # 🆕 Withdrawal with reasons
│   │   │   ├── entity.go                     # Withdrawal (ProposalID, FreelancerID, Reason, WithdrawnAt, Notes)
│   │   │   ├── reason.go                     # Reason: BetterOpportunity, ClientUnresponsive, ScopeChange, PersonalReasons
│   │   │   ├── errors.go                     # Domain errors (ErrAlreadyWithdrawn, ErrInvalidReason)
│   │   │   ├── repository.go                 # WithdrawalRepository (Withdraw, GetByProposalID)
│   │   │   └── events.go                     # Events: ProposalWithdrawn, WithdrawalReasonUpdated
│   │   │
│   │   ├── archive/                          # 🆕 Archiving inactive proposals
│   │   │   ├── entity.go                     # Archive (ProposalID, ArchivedAt, Reason, RestorableUntil)
│   │   │   ├── errors.go                     # Domain errors (ErrAlreadyArchived, ErrNotRestorable)
│   │   │   ├── repository.go                 # ArchiveRepository (Archive, Restore, ListArchived)
│   │   │   └── events.go                     # Events: ProposalArchived, ProposalRestored
│   │   │
│   │   # ========================= TEAM & AUDIT =========================
│   │   ├── team/                             # 🆕 Team/agency proposals
│   │   │   ├── entity.go                     # TeamProposal (ProposalID, TeamID, Members[], Roles[], LeadFreelancerID)
│   │   │   ├── errors.go                     # Domain errors (ErrTeamNotFound, ErrMemberNotAvailable)
│   │   │   ├── repository.go                 # TeamRepository (CreateTeamProposal, AddMember, GetTeamByProposalID)
│   │   │   └── events.go                     # Events: TeamProposalCreated, TeamMemberAdded, TeamMemberRemoved
│   │   │
│   │   ├── audit/                            # 🆕 Proposal audit trails
│   │   │   ├── entity.go                     # AuditLog (ProposalID, Action, ActorID, Timestamp, Details)
│   │   │   ├── errors.go                     # Domain errors (ErrLogFailed)
│   │   │   ├── repository.go                 # AuditRepository (LogAction, GetLogsByProposalID)
│   │   │   └── events.go                     # Events: ProposalAudited
│   │   │
│   │   # ========================= INTEGRATIONS & SECURITY =========================
│   │   ├── integration/                      # 🆕 External integrations (calendars, tools)
│   │   │   ├── entity.go                     # Integration (ProposalID, Type, Config, Status)
│   │   │   ├── type.go                       # Integration types (Calendar, PaymentGateway, CRM)
│   │   │   ├── errors.go                     # Domain errors (ErrIntegrationFailed, ErrInvalidConfig)
│   │   │   ├── repository.go                 # IntegrationRepository (Connect, Disconnect, GetByProposalID)
│   │   │   └── events.go                     # Events: IntegrationConnected, IntegrationDisconnected, IntegrationSynced
│   │   │
│   │   ├── secure_share/                     # 🆕 Secure shares for proposal artifacts
│   │   │   ├── entity.go                     # ShareLink(acl, ttl, watermark), AccessLog
│   │   │   ├── errors.go                     # ErrShareExpired, ErrAccessDenied
│   │   │   ├── repository.go                 # CreateLink, RevokeLink, LogAccess, GetLink
│   │   │   └── events.go                     # share.link.created, share.link.accessed, share.link.revoked
│   │   │
│   │   ├── localization/                     # 🆕 Multi-language proposals
│   │   │   ├── entity.go                     # Localization (ProposalID, Language, TranslatedContent, OriginalLanguage)
│   │   │   ├── errors.go                     # Domain errors (ErrTranslationFailed, ErrUnsupportedLanguage)
│   │   │   ├── repository.go                 # LocalizationRepository (Translate, GetTranslation, ListLanguages)
│   │   │   └── events.go                     # Events: ProposalTranslated, TranslationUpdated
│   │   │
│   │   └── dispute_prediction/               # 🆕 Early dispute risk signal
│   │       ├── entity.go                     # DisputePrediction (ProposalID, RiskScore, Factors[], Recommendations)
│   │       ├── errors.go                     # Domain errors (ErrHighRisk)
│   │       ├── repository.go                 # DisputePredictionRepository (Predict, UpdateScore, GetByProposalID)
│   │       └── events.go                     # Events: DisputeRiskPredicted, RiskMitigated
│   │
│   ├── application/                          # 📋 Application Layer (Use Cases - Load Fourth)
│   │   │
│   │   # ========================= EVENT HANDLERS (Inbound Events) =========================
│   │   ├── eventhandler/
│   │   │   ├── subscription_handler.go       # Consumes: subscription.created | subscription.changed | subscription.paused |
│   │   │   │                                 #           subscription.canceled | subscription.renewed
│   │   │   │                                 # Purpose: keep the current plan on the proposer (plan badge/limits context).
│   │   │   │
│   │   │   ├── entitlement_handler.go        # Consumes: entitlement.feature.enabled | entitlement.feature.disabled
│   │   │   │                                 # Purpose: flip feature flags/capabilities that gate proposal actions (posting limits,
│   │   │   │                                 #           boosts eligibility, extra invites), as per subscriptions-be → proposals-be.
│   │   │   │
│   │   │   ├── connects_handler.go           # Consumes: connects.purchased | connects.used | connects.expired | connects.granted
│   │   │   │                                 # Purpose: mirror connects balance/aging locally to enforce submit/boost flows.
│   │   │   │
│   │   │   ├── usage_handler.go              # Consumes: usage.meter.incremented | usage.limit.reached
│   │   │   │                                 # Purpose: enforce per-feature metering (e.g., boosts, extra invites, messages_to_non_hires).
│   │   │   │
│   │   │   └── admin_flags_handler.go        # Consumes: admin.feature_flag.updated | admin.config.updated | admin.experiment.updated
│   │   │                                     # Purpose: refresh runtime toggles/config that affect proposal flows and UI gating.
│   │   │
│   │   # ========================= CORE PROPOSAL SERVICES =========================
│   │   ├── proposal/
│   │   │   ├── service.go                    # Proposal business logic (Submit, Withdraw, Update)
│   │   │   ├── commands.go                   # SubmitProposalCommand, WithdrawProposalCommand, UpdateProposalCommand
│   │   │   ├── queries.go                    # GetProposal, ListProposals, FilterProposals
│   │   │   ├── dto.go                        # ProposalDTO, CreateProposalDTO, UpdateProposalDTO
│   │   │   ├── mapper.go                     # Entity-DTO mapping
│   │   │   ├── validators.go                 # Validation (amount, duration, cover letter; 🆕 structured schema)
│   │   │   └── business_rules.go             # Business rules (cannot submit duplicate proposals)
│   │   │
│   │   ├── milestone/
│   │   │   ├── service.go                    # Milestone logic (Create, Update, Delete)
│   │   │   ├── dto.go                        # MilestoneDTO
│   │   │   └── mapper.go                     # Milestone mappers
│   │   │
│   │   # ========================= BIDDING SYSTEM SERVICES =========================
│   │   ├── bid/
│   │   │   ├── service.go                    # Bid business logic
│   │   │   ├── commands.go                   # PlaceBidCommand, UpdateBidCommand, WithdrawBidCommand
│   │   │   ├── queries.go                    # GetBidStatus, GetBidHistory, GetBidRank
│   │   │   ├── dto.go                        # BidDTO, PlaceBidDTO
│   │   │   ├── mapper.go                     # Bid mappers
│   │   │   ├── bid_calculator.go             # Calculate optimal bid amounts
│   │   │   ├── bid_validator.go              # Validate bid amounts (min/max constraints)
│   │   │   ├── auto_bid_manager.go           # Auto-bidding management
│   │   │   └── ranking_engine.go             # Rank bids (1st, 2nd, 3rd)
│   │   │
│   │   ├── bid_strategy/
│   │   │   ├── service.go                    # Bid strategy logic
│   │   │   ├── strategy_executor.go          # Execute bidding strategies
│   │   │   ├── dto.go                        # StrategyDTO
│   │   │   └── mapper.go                     # Strategy mappers
│   │   │
│   │   ├── auction/
│   │   │   ├── service.go                    # Place bid; compute top; end auction safely
│   │   │   ├── commands.go                   # PlaceBid, EndAuction
│   │   │   ├── queries.go                    # GetAuctionStatus
│   │   │   ├── dto.go                        # AuctionDTO, BidDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # Min increment, auction window
│   │   │
│   │   ├── bid_anomaly_detection/
│   │   │   ├── service.go                    # Detect & triage anomalies
│   │   │   ├── commands.go                   # DetectBidAnomaly, UpdateAnomalyReview
│   │   │   ├── queries.go                    # GetAnomaliesForJob
│   │   │   ├── dto.go                        # AnomalyDTO, ReviewStateDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # Thresholds & reviewer gates
│   │   │
│   │   # ========================= CONNECTS & BOOST SERVICES =========================
│   │   ├── connect/
│   │   │   ├── service.go                    # Connect logic (Use, Purchase, Refund)
│   │   │   ├── calculator.go                 # Calculate connect cost per job
│   │   │   ├── dto.go                        # ConnectDTO
│   │   │   └── validators.go                 # Validate balance, tier inputs
│   │   │
│   │   ├── connect_refund/
│   │   │   ├── service.go                    # Request/process/deny
│   │   │   ├── commands.go                   # RequestConnectRefund, ProcessConnectRefund, DenyConnectRefund
│   │   │   ├── queries.go                    # GetConnectRefundStatus
│   │   │   ├── dto.go                        # ConnectRefundDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Eligibility/causation
│   │   │
│   │   ├── boost/
│   │   │   ├── service.go                    # Boost logic (Boost proposal, Calculate cost)
│   │   │   └── dto.go                        # BoostDTO
│   │   │
│   │   # ========================= TEMPLATES & RATE CARDS SERVICES =========================
│   │   ├── template/
│   │   │   ├── service.go                    # Template logic (Create, Update, Use)
│   │   │   ├── dto.go                        # TemplateDTO
│   │   │   └── mapper.go                     # Template mappers
│   │   │
│   │   ├── rate_card/
│   │   │   ├── service.go                    # Create/Update/Delete cards; pick best-fit card for job
│   │   │   ├── commands.go                   # CreateRateCard, UpdateRateCard, DeleteRateCard
│   │   │   ├── queries.go                    # GetRateCard, ListRateCardsByFreelancer, SuggestRateCardForJob
│   │   │   ├── dto.go                        # RateCardDTO
│   │   │   ├── mapper.go                     # Rate card mappers
│   │   │   └── validators.go                 # Tier names, price>0, delivery estimate present
│   │   │
│   │   # ========================= ANALYTICS & TRACKING SERVICES =========================
│   │   ├── analytics/
│   │   │   ├── service.go                    # Analytics logic (Track views, Calculate metrics)
│   │   │   ├── performance_tracker.go        # Track proposal performance
│   │   │   └── dto.go                        # AnalyticsDTO
│   │   │
│   │   ├── response_tracker/
│   │   │   ├── service.go                    # Track views & responses
│   │   │   ├── commands.go                   # TrackProposalView, TrackClientResponse
│   │   │   ├── queries.go                    # GetResponseStats
│   │   │   ├── dto.go                        # ResponseStatsDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # Idempotency, timestamps
│   │   │
│   │   ├── insight/
│   │   │   ├── service.go                    # Upsert insights; calculate aggregates
│   │   │   ├── commands.go                   # RecalculateInsight
│   │   │   ├── queries.go                    # GetProposalInsight
│   │   │   ├── dto.go                        # InsightDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Window/access validation
│   │   │
│   │   ├── ranking/
│   │   │   ├── service.go                    # Compute/recompute
│   │   │   ├── commands.go                   # RecomputeRanking
│   │   │   ├── queries.go                    # GetRankingForProposal
│   │   │   ├── dto.go                        # RankingDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Factor guards
│   │   │
│   │   # ========================= COLLABORATION & WORKFLOW SERVICES =========================
│   │   ├── negotiation/
│   │   │   ├── service.go                    # ProposeCounter, AcceptCounter, DeclineCounter, CloseThread
│   │   │   ├── commands.go                   # ProposeCounter, AcceptCounter, DeclineCounter
│   │   │   ├── queries.go                    # GetNegotiation, ListNegotiationsForProposal
│   │   │   ├── dto.go                        # NegotiationDTO, CounterOfferDTO
│   │   │   ├── mapper.go                     # Negotiation mappers
│   │   │   └── validators.go                 # Validate states, amounts, milestones
│   │   │
│   │   ├── invite/
│   │   │   ├── service.go                    # SendInvite, AcceptInvite, DeclineInvite (reason), PrefillProposal
│   │   │   ├── commands.go                   # CreateInvite, AcceptInvite, DeclineInvite, ResendInvite
│   │   │   ├── queries.go                    # GetInvite, ListInvitesForJob, ListInvitesForFreelancer
│   │   │   ├── dto.go                        # InviteDTO
│   │   │   ├── mapper.go                     # Invite mappers
│   │   │   └── validators.go                 # Validate receiver, expiry, decline reasons
│   │   │
│   │   ├── revision/
│   │   │   ├── service.go                    # Create/approve/reject
│   │   │   ├── commands.go                   # CreateRevision, ApproveRevision, RejectRevision
│   │   │   ├── queries.go                    # GetRevisionHistory, GetLatestRevision
│   │   │   ├── dto.go                        # RevisionDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # State transitions
│   │   │
│   │   ├── collaboration/
│   │   │   ├── service.go                    # Add/remove collaborators
│   │   │   ├── commands.go                   # AddCollaborator, RemoveCollaborator
│   │   │   ├── queries.go                    # GetCollaborators
│   │   │   ├── dto.go                        # CollaborationDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Access control
│   │   │
│   │   ├── approval_workflow/
│   │   │   ├── service.go                    # Start/advance/reject approvals; SLA checks
│   │   │   ├── commands.go                   # RequestApproval, GrantApproval, RejectApproval
│   │   │   ├── queries.go                    # GetApprovalState
│   │   │   ├── dto.go                        # ApprovalFlowDTO, StepDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # Role/approver/SLA validation
│   │   │
│   │   ├── document_redlining/
│   │   │   ├── service.go                    # Start/resolution of redline threads
│   │   │   ├── commands.go                   # StartRedlineThread, ResolveRedlineThread
│   │   │   ├── queries.go                    # GetRedlineThread
│   │   │   ├── dto.go                        # RedlineThreadDTO, DiffRefDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # Diff validation & permissions
│   │   │
│   │   # ========================= CLIENT INTERACTION SERVICES =========================
│   │   ├── interview_request/
│   │   │   ├── service.go                    # Request/schedule/complete
│   │   │   ├── commands.go                   # RequestInterview, UpdateInterviewStatus
│   │   │   ├── queries.go                    # ListInterviewRequests
│   │   │   ├── dto.go                        # InterviewRequestDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Availability/ownership
│   │   │
│   │   ├── feedback/
│   │   │   ├── service.go                    # Give/respond
│   │   │   ├── commands.go                   # GiveProposalFeedback, RespondToFeedback
│   │   │   ├── queries.go                    # GetProposalFeedback
│   │   │   ├── dto.go                        # FeedbackDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Bounds/duplicates
│   │   │
│   │   ├── shortlist/
│   │   │   ├── service.go                    # Add/remove/update
│   │   │   ├── commands.go                   # ShortlistProposal, UpdateShortlistStatus, RemoveFromShortlist
│   │   │   ├── queries.go                    # GetShortlistForJob
│   │   │   ├── dto.go                        # ShortlistDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Status transitions
│   │   │
│   │   ├── client_insight/
│   │   │   ├── service.go                    # Fetch/refresh
│   │   │   ├── commands.go                   # RefreshClientInsight
│   │   │   ├── queries.go                    # GetClientInsight
│   │   │   ├── dto.go                        # ClientInsightDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Access/PII
│   │   │
│   │   ├── client_preference/
│   │   │   ├── service.go                    # Infer/update preferences
│   │   │   ├── commands.go                   # InferClientPreference, UpdateClientPreference
│   │   │   ├── queries.go                    # GetClientPreference
│   │   │   ├── dto.go                        # ClientPreferenceDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # Inference signals, bounds
│   │   │
│   │   # ========================= COMPLIANCE & LEGAL SERVICES =========================
│   │   ├── compliance/
│   │   │   ├── service.go                    # Run/resolve checks
│   │   │   ├── commands.go                   # RecordComplianceCheck, ResolveComplianceFinding
│   │   │   ├── queries.go                    # GetLatestComplianceCheck
│   │   │   ├── dto.go                        # ComplianceDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Evidence/score
│   │   │
│   │   ├── flag/
│   │   │   ├── service.go                    # Flag logic
│   │   │   ├── commands.go                   # FlagProposal, UnflagProposal
│   │   │   ├── dto.go                        # FlagDTO
│   │   │   └── mapper.go                     # Flag mappers
│   │   │
│   │   ├── spam_detection/
│   │   │   ├── service.go                    # Scan/flag/clear
│   │   │   ├── commands.go                   # ScanProposalSpam, FlagSpam, ClearSpamFlags
│   │   │   ├── queries.go                    # GetSpamStatus
│   │   │   ├── dto.go                        # SpamDetectionDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Scores/reasons
│   │   │
│   │   ├── duplicate_check/
│   │   │   ├── service.go                    # Check/mark unique
│   │   │   ├── commands.go                   # CheckDuplicateProposals, MarkAsUnique
│   │   │   ├── queries.go                    # GetDuplicateSummary
│   │   │   ├── dto.go                        # DuplicateCheckDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Thresholds
│   │   │
│   │   ├── security_questionnaire/
│   │   │   ├── service.go                    # Issue/collect questionnaires
│   │   │   ├── commands.go                   # IssueQuestionnaire, CompleteQuestionnaire, RejectQuestionnaire
│   │   │   ├── queries.go                    # GetQuestionnaire
│   │   │   ├── dto.go                        # QuestionnaireDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # Type/status/file checks
│   │   │
│   │   ├── legal_hold_retention/
│   │   │   ├── service.go                    # Apply/lift holds; set retention
│   │   │   ├── commands.go                   # ApplyLegalHold, LiftLegalHold, SetRetentionPolicy
│   │   │   ├── queries.go                    # GetLegalHoldStatus
│   │   │   ├── dto.go                        # LegalHoldDTO, RetentionPolicyDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # Scope/policy conflicts
│   │   │
│   │   ├── data_residency/
│   │   │   ├── service.go                    # Enforce residency rules
│   │   │   ├── commands.go                   # SetResidencyRule
│   │   │   ├── queries.go                    # GetResidencyRules, GetViolations
│   │   │   ├── dto.go                        # ResidencyRuleDTO, ViolationDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # Zone/rule validation
│   │   │
│   │   ├── anonymization/
│   │   │   ├── service.go                    # Apply redaction; create anonymized variant
│   │   │   ├── commands.go                   # ApplyRedactionRules, CreateAnonymizedVariant
│   │   │   ├── queries.go                    # GetAnonymizedVariant
│   │   │   ├── dto.go                        # RedactionRulesDTO, AnonymizedVariantDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # PII policy checks
│   │   │
│   │   # ========================= CONTRACT & TERMS SERVICES =========================
│   │   ├── terms/
│   │   │   ├── service.go                    # Add/update/accept
│   │   │   ├── commands.go                   # AddTerms, UpdateTerms, AcceptTerms
│   │   │   ├── queries.go                    # GetTerms
│   │   │   ├── dto.go                        # TermsDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Clause checks
│   │   │
│   │   ├── nda/
│   │   │   ├── service.go                    # Propose/sign NDA
│   │   │   ├── commands.go                   # ProposeNDA, SignNDA
│   │   │   ├── queries.go                    # GetNDAForProposal
│   │   │   ├── dto.go                        # NDADTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Template/signature
│   │   │
│   │   ├── ip_rights/
│   │   │   ├── service.go                    # Set/agree rights
│   │   │   ├── commands.go                   # SetIPRights, AgreeIPRights
│   │   │   ├── queries.go                    # GetIPRights
│   │   │   ├── dto.go                        # IPRightsDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Assignment policy
│   │   │
│   │   ├── payment_terms/
│   │   │   ├── service.go                    # Set/update terms
│   │   │   ├── commands.go                   # SetPaymentTerms, UpdatePaymentTerms
│   │   │   ├── queries.go                    # GetPaymentTerms
│   │   │   ├── dto.go                        # PaymentTermsDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Schedule/fees
│   │   │
│   │   ├── pricing_model/
│   │   │   ├── service.go                    # Create/update model
│   │   │   ├── commands.go                   # SetPricingModel, AddTier
│   │   │   ├── queries.go                    # GetPricingModel
│   │   │   ├── dto.go                        # PricingModelDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Tier overlap/currency
│   │   │
│   │   # ========================= ENTERPRISE & PROCUREMENT SERVICES =========================
│   │   ├── rfp/
│   │   │   ├── service.go                    # Import RFP; build compliance matrix
│   │   │   ├── commands.go                   # ImportRFP, MapRFPSections
│   │   │   ├── queries.go                    # GetRFPCompliance
│   │   │   ├── dto.go                        # RFPImportDTO, ComplianceDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # File format & mapping guards
│   │   │
│   │   ├── quote/
│   │   │   ├── service.go                    # Issue/reprice/expire quotes
│   │   │   ├── commands.go                   # IssueQuote, RepriceQuote, ExpireQuote
│   │   │   ├── queries.go                    # GetQuoteForProposal
│   │   │   ├── dto.go                        # QuoteDTO, LineItemDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # Taxes/discounts/fx rules
│   │   │
│   │   ├── procurement/
│   │   │   ├── service.go                    # Attach PO; run budget checks
│   │   │   ├── commands.go                   # AttachPORef, RunBudgetCheck
│   │   │   ├── queries.go                    # GetProcurementStatus
│   │   │   ├── dto.go                        # ProcurementDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # Cost center & evidence checks
│   │   │
│   │   ├── evaluation_rubric/
│   │   │   ├── service.go                    # Create rubric; score proposals
│   │   │   ├── commands.go                   # CreateRubric, ScoreProposal
│   │   │   ├── queries.go                    # GetRubric, GetRubricScore
│   │   │   ├── dto.go                        # RubricDTO, RubricScoreDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # Weight normalization, criteria
│   │   │
│   │   ├── secure_share/
│   │   │   ├── service.go                    # Create/revoke secure links; log access
│   │   │   ├── commands.go                   # CreateShareLink, RevokeShareLink, LogShareAccess
│   │   │   ├── queries.go                    # GetShareLink, GetAccessLog
│   │   │   ├── dto.go                        # ShareLinkDTO, AccessLogDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # TTL/watermark/ACL guards
│   │   │
│   │   # ========================= AI & OPTIMIZATION SERVICES =========================
│   │   ├── ai_assist/
│   │   │   ├── service.go                    # Generate content; store drafts
│   │   │   ├── commands.go                   # GenerateProposalWithAI, SaveAIDraft
│   │   │   ├── queries.go                    # GetLatestAIDraft
│   │   │   ├── dto.go                        # AIDraftDTO, SuggestionDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Prompt/safety checks
│   │   │
│   │   ├── keyword_optimization/
│   │   │   ├── service.go                    # Scan & optimize proposal keywords
│   │   │   ├── commands.go                   # OptimizeKeywords, ApplyOptimization
│   │   │   ├── queries.go                    # GetKeywordSuggestions
│   │   │   ├── dto.go                        # KeywordOptDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # Keyword limits, score thresholds
│   │   │
│   │   ├── success_predictor/
│   │   │   ├── service.go                    # Predict & refresh models
│   │   │   ├── commands.go                   # PredictSuccess, UpdatePredictionModel
│   │   │   ├── queries.go                    # GetSuccessPrediction
│   │   │   ├── dto.go                        # SuccessPredictionDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # Data sufficiency, feature gates
│   │   │
│   │   ├── best_practices/
│   │   │   ├── service.go                    # Fetch tips; apply to proposal
│   │   │   ├── commands.go                   # ApplyBestPractice
│   │   │   ├── queries.go                    # GetBestPracticeTips
│   │   │   ├── dto.go                        # BestPracticeDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # Category/rule checks
│   │   │
│   │   ├── grammar_check/
│   │   │   ├── service.go                    # Run grammar/spell check; apply fixes
│   │   │   ├── commands.go                   # CheckGrammar, ApplyGrammarSuggestions
│   │   │   ├── queries.go                    # GetGrammarIssues
│   │   │   ├── dto.go                        # GrammarCheckDTO, IssueDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # Max issues, text size
│   │   │
│   │   ├── personalization/
│   │   │   ├── service.go                    # Personalize content for client
│   │   │   ├── commands.go                   # PersonalizeProposal
│   │   │   ├── queries.go                    # GetPersonalizedVariant
│   │   │   ├── dto.go                        # PersonalizationDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # PII & policy constraints
│   │   │
│   │   ├── ab_testing/
│   │   │   ├── service.go                    # Start/end test; compute winner
│   │   │   ├── commands.go                   # StartABTest, EndABTest, DeclareVariantWinner
│   │   │   ├── queries.go                    # GetABTestResults
│   │   │   ├── dto.go                        # ABTestDTO, VariantDTO
│   │   │   ├── mapper.go                     # Entity ↔ DTO
│   │   │   └── validators.go                 # Sample size, overlap rules
│   │   │
│   │   # ========================= PROFILE & PORTFOLIO SERVICES =========================
│   │   ├── portfolio_link/
│   │   │   ├── service.go                    # Link/unlink/list
│   │   │   ├── commands.go                   # LinkPortfolioItem, UnlinkPortfolioItem
│   │   │   ├── queries.go                    # ListPortfolioLinks
│   │   │   ├── dto.go                        # PortfolioLinkDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Dedup/order
│   │   │
│   │   ├── skill_match/
│   │   │   ├── service.go                    # Compute & fetch
│   │   │   ├── commands.go                   # ComputeSkillMatch
│   │   │   ├── queries.go                    # GetSkillMatch
│   │   │   ├── dto.go                        # SkillMatchDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Inputs/scoring
│   │   │
│   │   ├── video_introduction/
│   │   │   ├── service.go                    # Upload/transcribe
│   │   │   ├── commands.go                   # UploadVideoIntro, TranscribeVideoIntro
│   │   │   ├── queries.go                    # GetVideoIntro
│   │   │   ├── dto.go                        # VideoIntroDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Duration/format
│   │   │
│   │   ├── reference/
│   │   │   ├── service.go                    # Add/verify/list
│   │   │   ├── commands.go                   # AddReference, VerifyReference
│   │   │   ├── queries.go                    # ListReferences
│   │   │   ├── dto.go                        # ReferenceDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Verification
│   │   │
│   │   ├── membership_perk/
│   │   │   ├── service.go                    # Apply & track
│   │   │   ├── commands.go                   # ApplyPerkToProposal
│   │   │   ├── queries.go                    # GetPerkUsage
│   │   │   ├── dto.go                        # MembershipPerkDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Tier/limits
│   │   │
│   │   # ========================= LIFECYCLE MANAGEMENT SERVICES =========================
│   │   ├── expiration/
│   │   │   ├── service.go                    # Set/extend/notify
│   │   │   ├── commands.go                   # SetProposalExpiry, ExtendProposalExpiry, NotifyExpiry
│   │   │   ├── queries.go                    # GetExpiry
│   │   │   ├── dto.go                        # ExpiryDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Dates/renewals
│   │   │
│   │   ├── withdrawal/
│   │   │   ├── service.go                    # Withdraw/update reason
│   │   │   ├── commands.go                   # WithdrawProposal, UpdateWithdrawalReason
│   │   │   ├── queries.go                    # GetWithdrawal
│   │   │   ├── dto.go                        # WithdrawalDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Reason policy
│   │   │
│   │   ├── archive/
│   │   │   ├── service.go                    # Archive & restore proposals
│   │   │   ├── commands.go                   # ArchiveProposal, RestoreProposal
│   │   │   ├── queries.go                    # ListArchivedProposals
│   │   │   ├── dto.go                        # ArchiveDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Restore window
│   │   │
│   │   # ========================= TEAM & AUDIT SERVICES =========================
│   │   ├── team/
│   │   │   ├── service.go                    # Create/manage members
│   │   │   ├── commands.go                   # CreateTeamProposal, AddTeamMember, RemoveTeamMember
│   │   │   ├── queries.go                    # GetTeamForProposal
│   │   │   ├── dto.go                        # TeamProposalDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Capacity/roles
│   │   │
│   │   ├── audit/
│   │   │   ├── service.go                    # Append audit events
│   │   │   ├── commands.go                   # LogProposalAction
│   │   │   ├── queries.go                    # GetProposalAuditLog
│   │   │   ├── dto.go                        # AuditLogDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Actor/action checks
│   │   │
│   │   # ========================= INTEGRATIONS & LOCALIZATION SERVICES =========================
│   │   ├── integration/
│   │   │   ├── service.go                    # Connect/disconnect/sync
│   │   │   ├── commands.go                   # ConnectIntegration, DisconnectIntegration
│   │   │   ├── queries.go                    # GetIntegrations
│   │   │   ├── dto.go                        # IntegrationDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Config schema
│   │   │
│   │   ├── localization/
│   │   │   ├── service.go                    # Translate/update
│   │   │   ├── commands.go                   # TranslateProposal, UpdateTranslation
│   │   │   ├── queries.go                    # ListProposalLanguages
│   │   │   ├── dto.go                        # LocalizationDTO
│   │   │   ├── mapper.go                     # Entity↔DTO
│   │   │   └── validators.go                 # Language/size
│   │   │
│   │   └── dispute_prediction/
│   │       ├── service.go                    # Predict/update risk
│   │       ├── commands.go                   # PredictDisputeRisk, UpdateRiskScore
│   │       ├── queries.go                    # GetDisputeRisk
│   │       ├── dto.go                        # DisputeRiskDTO
│   │       ├── mapper.go                     # Entity↔DTO
│   │       └── validators.go                 # Model gates
│   │
│   ├── infrastructure/                       # 🔧 Infrastructure Layer
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection setup
│   │   │       ├── transaction.go            # Transaction helpers
│   │   │       ├── migrations.go             # 📝 UPDATED: Auto-migration with version tracking
│   │   │       ├── version.go                # 🆕 Schema version tracking
│   │   │       ├── safety.go                 # 🆕 Pre-migration safety checks
│   │   │       │
│   │   │       # ------------------------- CORE PROPOSAL LIFECYCLE REPOSITORIES -------------------------
│   │   │       ├── proposal_repository.go    # ProposalRepository implementation
│   │   │       ├── cover_letter_repository.go# CoverLetterRepository implementation
│   │   │       ├── attachment_repository.go  # AttachmentRepository implementation
│   │   │       ├── question_answer_repository.go # AnswerRepository implementation
│   │   │       ├── milestone_repository.go   # MilestoneRepository implementation
│   │   │       │
│   │   │       # ------------------------- BIDDING SYSTEM REPOSITORIES -------------------------
│   │   │       ├── bid_repository.go         # BidRepository implementation
│   │   │       ├── bid_strategy_repository.go# BidStrategyRepository implementation
│   │   │       ├── bid_notification_repository.go # NotificationRepository implementation
│   │   │       ├── auction_repository.go     # AuctionRepository implementation (bids/top/close)
│   │   │       ├── bid_anomaly_detection_repository.go # Anomaly snapshots & review states
│   │   │       │
│   │   │       # ------------------------- CONNECTS & BOOST REPOSITORIES -------------------------
│   │   │       ├── connect_repository.go     # ConnectRepository implementation
│   │   │       ├── connect_refund_repository.go # Connect refunds state
│   │   │       ├── boost_repository.go       # BoostRepository implementation
│   │   │       │
│   │   │       # ------------------------- TEMPLATE & RATE CARD REPOSITORIES -------------------------
│   │   │       ├── template_repository.go    # TemplateRepository implementation
│   │   │       ├── rate_card_repository.go   # RateCardRepository implementation
│   │   │       │
│   │   │       # ------------------------- ANALYTICS & TRACKING REPOSITORIES -------------------------
│   │   │       ├── analytics_repository.go   # AnalyticsRepository implementation
│   │   │       ├── response_tracker_repository.go # Views/responses stats
│   │   │       ├── insight_repository.go     # InsightRepository implementation (UPSERT, aggregates)
│   │   │       ├── ranking_repository.go     # Ranking scores & factors
│   │   │       │
│   │   │       # ------------------------- COLLABORATION & WORKFLOW REPOSITORIES -------------------------
│   │   │       ├── negotiation_repository.go # 🆕 Negotiation repository implementation
│   │   │       ├── invite_repository.go      # 🆕 Invite repository implementation
│   │   │       ├── revision_repository.go    # Proposal revision history
│   │   │       ├── collaboration_repository.go # Collaborators/permissions
│   │   │       ├── approval_workflow_repository.go # Store approval flows & steps; state machine
│   │   │       ├── document_redlining_repository.go # Redline threads & diffs
│   │   │       │
│   │   │       # ------------------------- CLIENT INTERACTION REPOSITORIES -------------------------
│   │   │       ├── interview_repository.go   # InterviewRequest storage
│   │   │       ├── feedback_repository.go    # Feedback CRUD
│   │   │       ├── shortlist_repository.go   # Shortlist entries
│   │   │       ├── client_insight_repository.go # Client insights cache
│   │   │       ├── client_preference_repository.go # Inferred client prefs
│   │   │       │
│   │   │       # ------------------------- COMPLIANCE & LEGAL REPOSITORIES -------------------------
│   │   │       ├── compliance_repository.go  # Compliance checks/results
│   │   │       ├── proposal_flag_repository.go # FlagRepository implementation
│   │   │       ├── spam_detection_repository.go # Spam scores/flags
│   │   │       ├── duplicate_check_repository.go # Similarity clusters
│   │   │       ├── security_questionnaire_repository.go # Questionnaires & statuses
│   │   │       ├── legal_hold_retention_repository.go # Holds & retention policies
│   │   │       ├── data_residency_repository.go # Residency rules & violations
│   │   │       ├── anonymization_repository.go # Redaction rules & variants
│   │   │       │
│   │   │       # ------------------------- CONTRACT & TERMS REPOSITORIES -------------------------
│   │   │       ├── terms_repository.go       # Terms & acceptance
│   │   │       ├── nda_repository.go         # NDA proposals/signatures
│   │   │       ├── ip_rights_repository.go   # IP rights agreements
│   │   │       ├── payment_terms_repository.go # Payment terms
│   │   │       ├── pricing_model_repository.go # Pricing models/tiers
│   │   │       │
│   │   │       # ------------------------- ENTERPRISE & PROCUREMENT REPOSITORIES -------------------------
│   │   │       ├── rfp_repository.go         # RFP imports, section maps, compliance matrix
│   │   │       ├── quote_repository.go       # Quotes, line items, fx locks
│   │   │       ├── procurement_repository.go # PO refs & budget checks
│   │   │       ├── evaluation_rubric_repository.go # Rubrics & scores
│   │   │       │
│   │   │       # ------------------------- AI & OPTIMIZATION REPOSITORIES -------------------------
│   │   │       ├── ai_assist_repository.go   # AIAssist drafts & usage logs
│   │   │       ├── keyword_optimization_repository.go # KeywordOpt storage (suggestions/diffs)
│   │   │       ├── success_predictor_repository.go # Predictions & model versions
│   │   │       ├── grammar_check_repository.go # Grammar issues & fixes
│   │   │       ├── personalization_repository.go # Personalized variants
│   │   │       ├── best_practices_repository.go # Tips/compliance rules
│   │   │       ├── ab_testing_repository.go  # A/B tests & metrics
│   │   │       │
│   │   │       # ------------------------- PROFILE & PORTFOLIO REPOSITORIES -------------------------
│   │   │       ├── portfolio_link_repository.go # Linked portfolio items
│   │   │       ├── skill_match_repository.go # Match scores & gaps
│   │   │       ├── video_introduction_repository.go # Video intro & transcript refs
│   │   │       ├── reference_repository.go   # References & verification
│   │   │       ├── membership_perk_repository.go # Perk usage/eligibility
│   │   │       │
│   │   │       # ------------------------- LIFECYCLE MANAGEMENT REPOSITORIES -------------------------
│   │   │       ├── expiration_repository.go  # Expiration & notifications
│   │   │       ├── withdrawal_repository.go  # Withdrawals
│   │   │       ├── archive_repository.go     # Archive/restore state
│   │   │       │
│   │   │       # ------------------------- TEAM & AUDIT REPOSITORIES -------------------------
│   │   │       ├── team_repository.go        # Team proposal & members
│   │   │       ├── audit_repository.go       # Audit log append/read
│   │   │       │
│   │   │       # ------------------------- INTEGRATIONS & SECURITY REPOSITORIES -------------------------
│   │   │       ├── integration_repository.go # External integrations
│   │   │       ├── secure_share_repository.go # Share links & access logs
│   │   │       ├── localization_repository.go # Translations
│   │   │       └── dispute_prediction_repository.go # Risk predictions
│   │   │
│   │   │
│   │   ├── clients/
│   │   │   └── subscriptions/                # 🆕 Subscriptions-BE integration (connects/credits)
│   │   │       └── client.go                 # HTTP/Dapr client to consume “connects/credits”; fetch tier pricing by job popularity
│   │   │
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go             # Redis connection
│   │   │       ├── proposal_cache.go         # Proposal caching
│   │   │       |── bid_cache.go              # Bid ranking cache
│   │   │       ├── insight_cache.go          # Cache proposal insights (Get/Set/Invalidate)
│   │   │       ├── ranking_cache.go          # Cache ranking snapshots
│   │   │       |── client_insight_cache.go   # Cache client insight lookups
│   │   │       ├── auction_cache.go                  # Cache auction status/current top (short TTL)
│   │   │       ├── keyword_opt_cache.go              # Cache keyword suggestions (bust on apply)
│   │   │       ├── success_prediction_cache.go       # Cache last prediction per proposal
│   │   │       ├── best_practices_cache.go           # Cache tips by category
│   │   │       ├── grammar_check_cache.go            # Cache latest issues snapshot
│   │   │       ├── personalization_cache.go          # Cache personalized variants
│   │   │       ├── ab_testing_cache.go               # Cache active tests & assignments
│   │   │       └── response_stats_cache.go           # Cache response stats aggregates
│   │   │       ├── topics_auction_and_opt.go         # 🆕 proposal.auction.*, proposal.keywords.*
│   │   │       |── producers_auction_and_opt.go      # Thin producers for new events (outbox)
│   │   │       ├── approval_workflow_cache.go          # Cache current approval step/SLA
│   │   │       ├── rfp_cache.go                        # Cached compliance summaries
│   │   │       ├── quote_cache.go                      # Cache active quote by proposal
│   │   │       ├── rubric_cache.go                     # Cache rubric configs
│   │   │       ├── residency_rule_cache.go             # Cache residency rules by zone
│   │   │       └── anomaly_cache.go                    # Cache recent anomalies per job
│   │   │
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go               # 📝 UPDATED: Uses platform-shared/inbox
│   │   │       ├── producer.go               # 📝 UPDATED: Uses platform-shared/outbox
│   │   │       ├── topics.go                 # 📝 UPDATED: From contracts/events (proposal.submitted, bid.placed, outbid.alert)
│   │   │       └── scram.go                  # SCRAM authentication
│   │   │
│   │   |── storage/
│   │   |    └── client.go                     # Storage service client (upload attachments)
│   │   └── ai/
│   │       └── provider_client.go            # AI provider client (rate limits, retries, safety tags)
│   │       ├── grammar_provider_client.go            # Provider client for grammar/spell check
│   │       └── personalization_engine_client.go      # Personalization model client
│   │
│   ├── interfaces/
│       └── http/
│           │
│           # ========================= CORE PROPOSAL HANDLERS =========================
│           ├── handlers/
│           │   ├── proposal_handler.go       # Proposal handlers (POST /proposals, GET /proposals/:id)
│           │   ├── milestone_handler.go      # Milestone handlers (POST /proposals/:id/milestones)
│           │   │
│           │   # ========================= BIDDING SYSTEM HANDLERS =========================
│           │   ├── bid_handler.go            # Bid handlers (POST /proposals/:id/bid)
│           │   ├── bid_strategy_handler.go   # Strategy handlers (POST /bid-strategies)
│           │   ├── auction_handler.go        # POST /jobs/:id/auction/bid, GET /jobs/:id/auction/status, POST /jobs/:id/auction/end
│           │   ├── bid_anomaly_detection_handler.go # POST detect/update review; GET anomalies
│           │   │
│           │   # ========================= CONNECTS & BOOST HANDLERS =========================
│           │   ├── connect_handler.go        # Connect handlers (GET /connects/balance)
│           │   ├── connect_refund_handler.go # POST request/process/deny, GET status
│           │   ├── boost_handler.go          # Boost handlers (POST /proposals/:id/boost)
│           │   │
│           │   # ========================= TEMPLATES & RATE CARDS HANDLERS =========================
│           │   ├── template_handler.go       # Template handlers (GET /templates, POST /templates)
│           │   ├── rate_card_handler.go      # Rate card handlers
│           │   │
│           │   # ========================= ANALYTICS & TRACKING HANDLERS =========================
│           │   ├── analytics_handler.go      # Analytics handlers (GET /proposals/:id/analytics)
│           │   ├── response_tracker_handler.go # POST /proposals/:id/track/view, POST /proposals/:id/track/response, GET /proposals/:id/track/stats
│           │   ├── insight_handler.go        # GET /proposals/:id/insight
│           │   ├── ranking_handler.go        # POST recompute, GET rank
│           │   │
│           │   # ========================= COLLABORATION & WORKFLOW HANDLERS =========================
│           │   ├── negotiation_handler.go    # 🆕 Negotiation endpoints (/proposals/:id/negotiations/*)
│           │   ├── invite_handler.go         # 🆕 Invite endpoints (/invites/*)
│           │   ├── revision_handler.go       # POST /proposals/:id/revisions, PATCH /proposals/:id/revisions/:ver
│           │   ├── collaboration_handler.go  # POST add/remove collaborator
│           │   ├── approval_workflow_handler.go # POST request/grant/reject; GET state
│           │   ├── document_redlining_handler.go # POST start/resolve; GET thread
│           │   │
│           │   # ========================= CLIENT INTERACTION HANDLERS =========================
│           │   ├── interview_handler.go      # POST /proposals/:id/interviews, PATCH status
│           │   ├── feedback_handler.go       # POST feedback, POST respond
│           │   ├── shortlist_handler.go      # POST shortlist, PATCH status, DELETE
│           │   ├── client_insight_handler.go # GET /clients/:id/insight
│           │   ├── client_preference_handler.go # POST /clients/:id/preferences/infer, PATCH /clients/:id/preferences, GET /clients/:id/preferences
│           │   │
│           │   # ========================= COMPLIANCE & LEGAL HANDLERS =========================
│           │   ├── compliance_handler.go     # POST checks, GET latest
│           │   ├── flag_handler.go           # Flag handlers (POST /proposals/:id/flag)
│           │   ├── spam_handler.go           # POST scan, POST flag/clear, GET status
│           │   ├── duplicate_handler.go      # POST check, POST mark-unique
│           │   ├── security_questionnaire_handler.go # POST issue/complete/reject; GET questionnaire
│           │   ├── legal_hold_retention_handler.go # POST apply/lift; POST set policy; GET status
│           │   ├── data_residency_handler.go # POST set rule; GET rules/violations
│           │   ├── anonymization_handler.go  # POST apply rules / create variant; GET variant
│           │   │
│           │   # ========================= CONTRACT & TERMS HANDLERS =========================
│           │   ├── terms_handler.go          # POST add/update, POST accept
│           │   ├── nda_handler.go            # POST propose/sign, GET
│           │   ├── ip_rights_handler.go      # POST set/agree, GET
│           │   ├── payment_terms_handler.go  # POST set/update, GET
│           │   ├── pricing_handler.go        # POST set model, POST add tier, GET
│           │   │
│           │   # ========================= ENTERPRISE & PROCUREMENT HANDLERS =========================
│           │   ├── rfp_handler.go            # POST import/map; GET compliance
│           │   ├── quote_handler.go          # POST issue/reprice/expire; GET
│           │   ├── procurement_handler.go    # POST attach PO / run budget check; GET status
│           │   ├── evaluation_rubric_handler.go # POST create rubric / score; GET rubric/score
│           │   ├── secure_share_handler.go   # POST create/revoke; GET link/access log
│           │   │
│           │   # ========================= AI & OPTIMIZATION HANDLERS =========================
│           │   ├── ai_assist_handler.go      # POST /proposals/:id/ai/generate, GET /proposals/:id/ai/draft
│           │   ├── keyword_optimization_handler.go # POST /proposals/:id/keywords/optimize, GET /proposals/:id/keywords/suggestions
│           │   ├── success_predictor_handler.go # POST /proposals/:id/success/predict, GET /proposals/:id/success
│           │   ├── best_practices_handler.go # GET /proposals/:id/best-practices, POST /proposals/:id/best-practices/apply
│           │   ├── grammar_check_handler.go  # POST /proposals/:id/grammar/check, POST /proposals/:id/grammar/apply, GET /proposals/:id/grammar/issues
│           │   ├── personalization_handler.go # POST /proposals/:id/personalize, GET /proposals/:id/personalized
│           │   ├── ab_testing_handler.go     # POST /ab-tests, POST /ab-tests/:id/end, POST /ab-tests/:id/declare-winner, GET /ab-tests/:id
│           │   │
│           │   # ========================= PROFILE & PORTFOLIO HANDLERS =========================
│           │   ├── portfolio_link_handler.go # POST link, DELETE unlink, GET list
│           │   ├── skill_match_handler.go    # POST compute, GET match
│           │   ├── video_intro_handler.go    # POST upload/transcribe, GET
│           │   ├── reference_handler.go      # POST add/verify, GET list
│           │   ├── membership_perk_handler.go # POST apply perk, GET usage
│           │   │
│           │   # ========================= LIFECYCLE MANAGEMENT HANDLERS =========================
│           │   ├── expiration_handler.go     # POST set/extend, POST notify
│           │   ├── withdrawal_handler.go     # POST withdraw, PATCH reason
│           │   ├── archive_handler.go        # POST archive, POST restore
│           │   │
│           │   # ========================= TEAM & AUDIT HANDLERS =========================
│           │   ├── team_handler.go           # POST team, POST add/remove member
│           │   ├── audit_handler.go          # GET audit log
│           │   │
│           │   # ========================= INTEGRATIONS & LOCALIZATION HANDLERS =========================
│           │   ├── integration_handler.go    # POST connect/disconnect, GET
│           │   ├── localization_handler.go   # POST translate/update, GET languages
│           │   └── dispute_prediction_handler.go # POST predict/update, GET
│           │
│           # ========================= ROUTES (1:1 with Handlers) =========================
│           └── routes/                       # 🆕 Route registrars mapped 1:1 to handlers
│               │
│               # ========================= CORE PROPOSAL ROUTES =========================
│               ├── proposal_routes.go        # /proposals, /proposals/:id
│               ├── milestone_routes.go       # /proposals/:id/milestones
│               │
│               # ========================= BIDDING SYSTEM ROUTES =========================
│               ├── bid_routes.go             # /proposals/:id/bid
│               ├── bid_strategy_routes.go    # /bid-strategies
│               ├── auction_routes.go         # /jobs/:id/auction/*
│               ├── bid_anomaly_detection_routes.go # /jobs/:id/bid-anomalies/*
│               │
│               # ========================= CONNECTS & BOOST ROUTES =========================
│               ├── connect_routes.go         # /connects/*
│               ├── connect_refund_routes.go  # /proposals/:id/connect-refund
│               ├── boost_routes.go           # /proposals/:id/boost
│               │
│               # ========================= TEMPLATES & RATE CARDS ROUTES =========================
│               ├── template_routes.go        # /templates
│               ├── rate_card_routes.go       # /rate-cards
│               │
│               # ========================= ANALYTICS & TRACKING ROUTES =========================
│               ├── analytics_routes.go       # /proposals/:id/analytics
│               ├── response_tracker_routes.go # /proposals/:id/track/*
│               ├── insight_routes.go         # /proposals/:id/insight
│               ├── ranking_routes.go         # /proposals/:id/ranking
│               │
│               # ========================= COLLABORATION & WORKFLOW ROUTES =========================
│               ├── negotiation_routes.go     # 🆕 /proposals/:id/negotiations (post/accept/decline/list)
│               ├── invite_routes.go          # 🆕 /invites (create/accept/decline/resend)
│               ├── revision_routes.go        # /proposals/:id/revisions
│               ├── collaboration_routes.go   # /proposals/:id/collaboration
│               ├── approval_workflow_routes.go # /proposals/:id/approvals/*
│               ├── document_redlining_routes.go # /proposals/:id/redlines/*
│               │
│               # ========================= CLIENT INTERACTION ROUTES =========================
│               ├── interview_routes.go       # /proposals/:id/interviews
│               ├── feedback_routes.go        # /proposals/:id/feedback
│               ├── shortlist_routes.go       # /jobs/:id/shortlist
│               ├── client_insight_routes.go  # /clients/:id/insight
│               ├── client_preference_routes.go # /clients/:id/preferences
│               │
│               # ========================= COMPLIANCE & LEGAL ROUTES =========================
│               ├── compliance_routes.go      # /proposals/:id/compliance
│               ├── flag_routes.go            # /proposals/:id/flag
│               ├── spam_routes.go            # /proposals/:id/spam
│               ├── duplicate_routes.go       # /proposals/:id/duplicates
│               ├── security_questionnaire_routes.go # /proposals/:id/security-questionnaire/*
│               ├── legal_hold_retention_routes.go # /proposals/:id/legal/*
│               ├── data_residency_routes.go  # /proposals/:id/residency/*
│               ├── anonymization_routes.go   # /proposals/:id/anonymization/*
│               │
│               # ========================= CONTRACT & TERMS ROUTES =========================
│               ├── terms_routes.go           # /proposals/:id/terms
│               ├── nda_routes.go             # /proposals/:id/nda
│               ├── ip_rights_routes.go       # /proposals/:id/ip-rights
│               ├── payment_terms_routes.go   # /proposals/:id/payment-terms
│               ├── pricing_routes.go         # /proposals/:id/pricing
│               │
│               # ========================= ENTERPRISE & PROCUREMENT ROUTES =========================
│               ├── rfp_routes.go             # /proposals/:id/rfp/*
│               ├── quote_routes.go           # /proposals/:id/quote/*
│               ├── procurement_routes.go     # /proposals/:id/procurement/*
│               ├── evaluation_rubric_routes.go # /proposals/:id/rubric/*
│               ├── secure_share_routes.go    # /proposals/:id/share/*
│               │
│               # ========================= AI & OPTIMIZATION ROUTES =========================
│               ├── ai_assist_routes.go       # /proposals/:id/ai/*
│               ├── keyword_optimization_routes.go # /proposals/:id/keywords/*
│               ├── success_predictor_routes.go # /proposals/:id/success
│               ├── best_practices_routes.go  # /proposals/:id/best-practices
│               ├── grammar_check_routes.go   # /proposals/:id/grammar/*
│               ├── personalization_routes.go # /proposals/:id/personalize
│               ├── ab_testing_routes.go      # /ab-tests/*
│               │
│               # ========================= PROFILE & PORTFOLIO ROUTES =========================
│               ├── portfolio_link_routes.go  # /proposals/:id/portfolio-links
│               ├── skill_match_routes.go     # /proposals/:id/skill-match
│               ├── video_intro_routes.go     # /proposals/:id/video-intro
│               ├── reference_routes.go       # /proposals/:id/references
│               ├── membership_perk_routes.go # /proposals/:id/membership-perk
│               │
│               # ========================= LIFECYCLE MANAGEMENT ROUTES =========================
│               ├── expiration_routes.go      # /proposals/:id/expiration
│               ├── withdrawal_routes.go      # /proposals/:id/withdrawal
│               ├── archive_routes.go         # /proposals/:id/archive
│               │
│               # ========================= TEAM & AUDIT ROUTES =========================
│               ├── team_routes.go            # /proposals/:id/team
│               ├── audit_routes.go           # /proposals/:id/audit
│               │
│               # ========================= INTEGRATIONS & LOCALIZATION ROUTES =========================
│               ├── integration_routes.go     # /proposals/:id/integrations
│               ├── localization_routes.go    # /proposals/:id/localization
│               │── dispute_prediction_routes.go # /proposals/:id/dispute-risk│   │       │
│               │
│               ├── middleware/                   # 📝 UPDATED: Uses platform-shared
│               │   ├── auth.go                   # 📝 UPDATED: Uses pkg/auth
│               │   ├── rbac.go                   # 📝 UPDATED: Uses pkg/auth
│               │   ├── cors.go                   # 📝 UPDATED: Uses platform-shared/ginx
│               │   ├── rate_limit.go             # Rate limiting
│               │   └── idempotency.go            # 🆕 Uses platform-shared/idempotency
│               │
│               ├── responses/
│               │   └── README.md                 # 📝 Points to platform-shared/httpx
│               │
│               ├── router.go                     # 📝 UPDATED: Uses platform-shared/ginx
│               │
│               ├── config/                               # 🆕 MEDIUM - Configuration Layer (Load First)
│               │   ├── schema.go                         # 🆕 Typed Config
│               │   ├── loader.go                         # 🆕 Viper loader
│               │   └── docs/
│               │       └── CONFIGURATION.md              # 🆕 Config docs
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
│       ├── pubsub.yaml                       # Scopes: ["proposals-be"]
│       ├── statestore.yaml
│       └── secrets.yaml
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   └── codes.go                          # PROPOSAL_NOT_FOUND, INSUFFICIENT_CONNECTS
│   ├── logger/                               # ❌ REMOVED
│   │   └── README.md
│   ├── utils/
│   │   ├── validator.go
│   │   ├── sanitizer.go
│   │   └── bid_calculator.go                 # Bid calculation utilities
│   └── constants/
│       ├── events.go                         # ❌ REMOVED
│       ├── topics.go                         # ❌ REMOVED
│       └── bid_types.go                      # Bid type constants
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
│   ├── EVENTS.md                             # 🆕 Events (proposal.submitted, bid.placed, outbid.alert, connect.used)
│   ├── ARCHITECTURE.md
│   ├── MIGRATIONS.md                         # 🆕
│   ├── SCHEMA.md                             # 🆕 Structured proposal schema reference
│   ├── RUNBOOK.md                            # 🆕
│   ├── bidding-system.md                     # Bidding system documentation
│   ├── connects-system.md                    # Connects/credits documentation
│   ├── credit-tiers.md                       # 🆕 Credit cost tiers & popularity mapping
│   ├── negotiations.md                       # 🆕 Negotiation thread lifecycle
│   ├── invites.md                            # 🆕 Invite flow & decline reasons
│   └── rate-cards.md                         # 🆕 Package tiers (starter/standard/premium)
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