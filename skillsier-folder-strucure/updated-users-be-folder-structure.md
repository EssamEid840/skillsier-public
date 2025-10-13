
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
│   │       ├── routes/
│   │       │   ├── user_routes.go       
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