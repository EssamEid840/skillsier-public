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


