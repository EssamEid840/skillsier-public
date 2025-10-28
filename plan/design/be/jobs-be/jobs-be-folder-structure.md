
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
