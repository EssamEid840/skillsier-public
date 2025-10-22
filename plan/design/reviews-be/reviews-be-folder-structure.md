## **📦9️⃣ reviews-be (UPDATED - COVERS RATINGS)**

```
apps/be/reviews-be/
│
├── cmd/
│   # =============================
│   # 🚀 APP ENTRYPOINTS
│   # =============================
│   ├── api/
│   │   └── main.go                               # 📝 Gin + Dapr + Postgres bootstrap; loadConfig(); platform-shared/logging
│   └── worker/                                    # 🆕 background workers (cron-ish)
│       └── main.go                                # 🆕 open/close double-blind windows, send reminders, rebuild stats/reputation
│
├── internal/
│   # =============================
│   # 🔧 CONFIGURATION (LOAD FIRST)
│   # =============================
│   ├── config/
│   │   ├── schema.go                              # Typed Config: App, Server, Postgres, Kafka, Redis, Windows, Limits
│   │   ├── loader.go                              # Viper loader: flags → env → file → defaults
│   │   └── docs/CONFIGURATION.md                  # ENV vars, defaults, examples, sample commands
│   │
│   # =============================
│   # 🏛️ DOMAIN LAYER (DDD)
│   # =============================
│   ├── domain/
│   │   # =========================
│   │   # ⭐ CORE REVIEW PRIMITIVES
│   │   # =========================
│   │   ├── review/
│   │   │   ├── entity.go                           # Aggregate: ids, body, ratings ver, status(draft|published|withdrawn), created/edited
│   │   │   ├── rating.go                           # Per-dimension values + normalized score
│   │   │   ├── enums.go                            # ReviewType(client→freelancer|freelancer→client), Status, Category
│   │   │   ├── response.go                         # Single public response from subject (when policy allows)
│   │   │   ├── helpful.go                          # Helpful votes with actor dedupe & 24h change window
│   │   │   ├── errors.go                           # 🆕 ReviewNotFound, WindowClosed, AlreadySubmitted, NotAuthor
│   │   │   ├── repository.go                       # CRUD + list by contract/user + helpful ops (interface)
│   │   │   └── events.go                           # 🆕 review.submitted/edited/published/withdrawn/helpful.v1
│   │   ├── eligibility/                            # 🆕 who may review whom (and when)
│   │   │   ├── entity.go                           # EligibilityResult{allowed,bool,reason,not_before,not_after}
│   │   │   ├── policy.go                           # Contract/payment state checks; daily limits; cooldowns
│   │   │   ├── errors.go                           # NotEligible, WindowNotOpen, RateLimited
│   │   │   ├── repository.go                       # Persist last decision per (author,contract) for auditability
│   │   │   └── events.go                           # 🆕 review.eligibility.policy.updated.v1
│   │   ├── review_draft/                           # 🆕 save-before-submit
│   │   │   ├── entity.go                           # Draft{author,contract_id,body,rating_map,updated_at}
│   │   │   ├── errors.go                           # DraftNotFound, TooLarge
│   │   │   └── repository.go                       # Upsert/Get/Delete
│   │
│   │   # =========================
│   │   # 🧮 RATING & CRITERIA
│   │   # =========================
│   │   ├── rating/
│   │   │   ├── entity.go                           # Criteria set (id, category, version, dims with weights & scales)
│   │   │   ├── criteria.go                         # Versioned criteria definitions per job category
│   │   │   ├── aggregation.go                      # Rollups per subject (mean, weighted mean, Wilson)
│   │   │   ├── versioning.go                       # 🆕 freeze criteria version on each review
│   │   │   ├── errors.go                           # CriteriaNotFound, InvalidWeight, VersionMismatch
│   │   │   ├── repository.go                       # Store criteria versions & aggregates
│   │   │   └── events.go                           # 🆕 rating.criteria.updated / rating.aggregated.v1
│   │
│   │   # =========================
│   │   # 🏅 BADGES & REPUTATION
│   │   # =========================
│   │   ├── badge/
│   │   │   ├── entity.go                           # Badge meta (id, name, levels, description, active)
│   │   │   ├── criteria.go                         # Rule expressions referencing aggregates/reputation
│   │   │   ├── types.go                            # Built-ins: Top Rated, Rising Talent, …
│   │   │   ├── level.go                            # Thresholds per level (Bronze/Silver/Gold)
│   │   │   ├── errors.go                           # InvalidRule, BadgeNotFound
│   │   │   ├── repository.go                       # Badge & criteria CRUD
│   │   │   └── events.go                           # 🆕 badge.created/criteria.updated/level.updated/archived.v1
│   │   ├── user_badge/
│   │   │   ├── entity.go                           # Assignment{user_id,badge_id,level,awarded_at,expires_at}
│   │   │   ├── achievement_date.go                  # Derived dates (renewals)
│   │   │   ├── errors.go                            # AlreadyHasBadge, NotEligible
│   │   │   ├── repository.go                        # Assign/Revoke/ListByUser
│   │   │   └── events.go                            # 🆕 badge.awarded/revoked.v1
│   │   ├── reputation/
│   │   │   ├── entity.go                            # Score components + total
│   │   │   ├── score_calculator.go                  # Bayesian prior + decay + recency boost
│   │   │   ├── history.go                           # Time series for trends
│   │   │   ├── decay.go                              # Outlier dampening & rolling windows
│   │   │   ├── top_rated_rules.go                    # Eligibility rules for "Top Rated"
│   │   │   ├── errors.go                             
│   │   │   ├── repository.go                        # Upsert score/history
│   │   │   └── events.go                             # 🆕 reputation.updated / top_rated.eligibility.changed.v1
│   │
│   │   # =========================
│   │   # 🔒 DOUBLE-BLIND & WINDOWS
│   │   # =========================
│   │   ├── double_blind/
│   │   │   ├── entity.go                            # Window{contract_id, open_at, close_at, submitted_client, submitted_freelancer}
│   │   │   ├── window.go                            # Publish when both sides submit or on timeout
│   │   │   ├── policy.go                            # Durations, grace periods, extension rules
│   │   │   ├── errors.go                            # WindowNotOpen, PublishBlocked
│   │   │   ├── repository.go                        # Window persistence
│   │   │   └── events.go                            # 🆕 double_blind.submitted/published/timeout.v1
│   │   ├── review_reminder/
│   │   │   ├── entity.go                            # Reminder{contract_id,user_id,scheduled_for,status}
│   │   │   ├── errors.go                            # ReminderNotFound, AlreadySent
│   │   │   ├── repository.go                        # Schedule/MarkSent/Failures
│   │   │   └── events.go                            # 🆕 review.reminder.scheduled/sent/failed.v1
│   │
│   │   # =========================
│   │   # 🗣️ PRIVATE FEEDBACK
│   │   # =========================
│   │   ├── feedback/
│   │   │   ├── entity.go                            # Private feedback (hidden from public)
│   │   │   ├── category.go                          # Confidential categories
│   │   │   ├── nps.go                                # 🆕 NPS-style numeric score
│   │   │   ├── errors.go
│   │   │   ├── repository.go
│   │   │   └── events.go                             # 🆕 private_feedback.submitted/updated.v1
│   │
│   │   # =========================
│   │   # 🛡️ SAFETY & GOVERNANCE
│   │   # =========================
│   │   ├── flag/
│   │   │   ├── entity.go                            # Flag{review_id,reason,reporter, state}
│   │   │   ├── reason.go                            # Reason enum + custom text
│   │   │   ├── errors.go
│   │   │   ├── repository.go                        # Open/Resolve/Dismiss
│   │   │   └── events.go                             # 🆕 review.flag.submitted/resolved/dismissed.v1
│   │   ├── moderation/
│   │   │   ├── entity.go                            # Moderation state machine
│   │   │   ├── auto_flag.go                         # Heuristics: bursts, swaps, duplicate IPs
│   │   │   ├── signals.go                           # Evidence signals kept for decisioning
│   │   │   ├── errors.go
│   │   │   ├── repository.go                        # Set state; store decisions
│   │   │   └── events.go                             # 🆕 review.abuse.auto_flagged/state_changed.v1
│   │   ├── appeal/                                  # 🆕 structured challenges
│   │   │   ├── entity.go                            # Appeal{review_id, raised_by, reason, state, outcome}
│   │   │   ├── policy.go                            # Time windows + required evidence
│   │   │   ├── repository.go
│   │   │   └── events.go                             # 🆕 review.appeal.opened/decided.v1
│   │   ├── evidence/                                # 🆕 file attachments for moderation/appeals
│   │   │   ├── entity.go                            # Evidence{file_id(storage-be), type, notes}
│   │   │   └── repository.go
│   │   ├── redaction/                               # 🆕 PII masking
│   │   │   ├── policy.go                            # Allowlist/denylist + length caps
│   │   │   ├── masker.go                            # Apply redaction to review/response text
│   │   │   └── events.go                             # 🆕 review.redacted.v1
│   │   ├── compliance/                              # 🆕 DSAR/erasure/export
│   │   │   ├── entity.go                            # ErasureTask, ExportTask, Status
│   │   │   ├── erasure.go                           # Hard/soft delete policy + tombstones
│   │   │   └── events.go                             # 🆕 compliance.erasure.requested/completed.v1
│   │   ├── audit/                                   # 🆕 immutable trail
│   │   │   └── trail.go                             # who/what/when for edits, moderation, publishes
│   │
│   │   # =========================
│   │   # 📊 STATS & EXPOSURE
│   │   # =========================
│   │   ├── stats/
│   │   │   ├── entity.go                            # Counts, histograms, recency
│   │   │   ├── aggregates.go                        # Recalc logic & snapshotting
│   │   │   ├── errors.go
│   │   │   ├── repository.go
│   │   │   └── events.go                             # 🆕 review.stats.updated/recalculated.v1
│   │   ├── featured/                                # 🆕 editorial (non-paid) highlights
│   │   │   ├── entity.go                            # FeaturedReview{review_id, boost, expires_at, reason}
│   │   │   ├── policy.go                            # Guardrails (no pay-to-boost ever)
│   │   │   └── repository.go
│   │
│   │   └── outbox/
│   │       ├── entity.go                            # ❌ REMOVED (use platform-shared/outbox/entity.go)
│   │       └── repository.go                        # ❌ REMOVED (use platform-shared/outbox/repository.go)
│   │
│   # =============================
│   # 📋 APPLICATION LAYER (CQRS)
│   # =============================
│   ├── application/
│   │   # =========================
│   │   # 📡 EVENT CONSUMERS (INBOX)
│   │   # =========================
│   │   ├── eventhandler/
│   │   │   ├── contract_handler.go                  # Consumes: contract.ended → open double-blind window for both parties
│   │   │   ├── payment_handler.go                   # Consumes: payment.captured → qualify/open window (hourly/fixed)
│   │   │   ├── admin_moderation_handler.go          # Consumes: admin.moderation.actioned → remove/restore; recalc aggregates
│   │   │   ├── admin_flags_handler.go               # Consumes: admin.feature_flag/threshold/experiment.updated → refresh toggles
│   │   │   ├── search_projection_handler.go         # 🆕 On publish/withdraw → emit aggregates to search-be
│   │   │   └── compliance_handler.go                # 🆕 Consumes: user.erasure.requested → cascade DSAR into reviews
│   │   #
│   │   # =========================
│   │   # 🧠 USE CASES BY DOMAIN
│   │   # =========================
│   │   # ---- CORE REVIEWS ----
│   │   ├── review/
│   │   │   ├── service.go                           # Orchestrates create/edit/publish/withdraw & helpful votes
│   │   │   ├── commands.go                          # CreateReview, EditReview, PublishReview, WithdrawReview, VoteHelpful
│   │   │   ├── queries.go                           # GetReview, ListByUser, ListByContract, SearchReviews
│   │   │   ├── validators.go                        # Body length, PII mask, eligibility gate, double-blind state
│   │   │   ├── mapper.go                            # Entity ↔ DTO (includes rating version binding)
│   │   │   ├── dto.go                                # ReviewDTO, NewReviewDTO, ResponseDTO, HelpfulDTO
│   │   │   └── business_rules.go                    # Cross-cutting: one-per-contract-per-side, response window policy
│   │   ├── review_draft/                            # 🆕
│   │   │   ├── service.go                           # Save/Clear/Load draft (redaction applied on read)
│   │   │   ├── commands.go                          # SaveDraft, ClearDraft
│   │   │   ├── queries.go                           # GetDraft
│   │   │   ├── validators.go                        # Max size, allowed placeholders
│   │   │   ├── mapper.go                            # Draft ↔ DTO
│   │   │   └── dto.go                                # ReviewDraftDTO
│   │   ├── eligibility/                             # 🆕
│   │   │   ├── service.go                           # CheckEligibility(author, contract)
│   │   │   ├── queries.go                           # GetEligibility
│   │   │   ├── validators.go                        # Rate limit keying, window bounds
│   │   │   └── dto.go                                # EligibilityDTO
│   │   # ---- RATING/CRITERIA ----
│   │   ├── rating/
│   │   │   ├── service.go                           # Criteria CRUD + aggregate pipeline
│   │   │   ├── aggregator.go                        # Rebuild aggregates for user/contract
│   │   │   ├── calculator.go                        # Weighted means & Wilson intervals
│   │   │   ├── commands.go                          # UpsertCriteria, RecalculateForContract
│   │   │   ├── queries.go                           # GetCriteria, GetAggregates
│   │   │   ├── validators.go                        # Weight sum=1, value in scale bounds
│   │   │   ├── mapper.go                            # Criteria/aggregate ↔ DTO
│   │   │   └── dto.go                                # CriteriaDTO, AggregateDTO
│   │   # ---- BADGES & USER BADGES ----
│   │   ├── badge/
│   │   │   ├── service.go                           # Evaluate rules; emit award/revoke
│   │   │   ├── achievement_checker.go               # Check badge eligibility from aggregates/reputation
│   │   │   ├── commands.go                          # AwardBadge, RevokeBadge, UpsertBadge
│   │   │   ├── queries.go                           # GetBadge, ListBadges
│   │   │   ├── validators.go                        # Rule safety, level thresholds
│   │   │   ├── mapper.go                            # Badge ↔ DTO
│   │   │   └── dto.go                                # BadgeDTO, AwardDTO
│   │   ├── user_badge/                              # 🆕
│   │   │   ├── service.go                           # Read/refresh assignments
│   │   │   ├── queries.go                           # ListUserBadges
│   │   │   ├── mapper.go                            # Assignment ↔ DTO
│   │   │   └── dto.go                                # UserBadgeDTO
│   │   # ---- REPUTATION ----
│   │   ├── reputation/
│   │   │   ├── service.go                           # Compute + persist score
│   │   │   ├── calculator.go                        # Totals + components
│   │   │   ├── updater.go                           # Apply decay & outlier handling
│   │   │   ├── commands.go                          # RecomputeUserReputation, ApplyDecay
│   │   │   ├── queries.go                           # GetReputation, GetHistory
│   │   │   ├── validators.go                        # Decay params bounds
│   │   │   ├── mapper.go                            # Reputation ↔ DTO
│   │   │   └── dto.go                                # ReputationDTO, HistoryPointDTO
│   │   # ---- DOUBLE-BLIND & REMINDERS ----
│   │   ├── double_blind/
│   │   │   ├── service.go                           # OpenWindow, SubmitSide, PublishIfReady
│   │   │   ├── commands.go                          # OpenWindow, SubmitClient, SubmitFreelancer, ForcePublish
│   │   │   ├── queries.go                           # GetWindow, GetPublishState
│   │   │   ├── validators.go                        # Time & state transitions
│   │   │   └── dto.go                                # WindowDTO, PublishStateDTO
│   │   ├── review_reminder/
│   │   │   ├── service.go                           # Plan & send nudges
│   │   │   ├── commands.go                          # ScheduleReminder, SendReminder
│   │   │   └── dto.go                                # ReminderDTO
│   │   # ---- PRIVATE FEEDBACK ----
│   │   ├── feedback/
│   │   │   ├── service.go
│   │   │   ├── commands.go                          # SubmitPrivateFeedback
│   │   │   ├── queries.go                           # GetPrivateFeedback
│   │   │   ├── validators.go                        # NPS bounds; forbid PII
│   │   │   ├── mapper.go                            # Feedback ↔ DTO
│   │   │   └── dto.go                                # PrivateFeedbackDTO
│   │   # ---- SAFETY/GOVERNANCE ----
│   │   ├── flag/
│   │   │   ├── service.go                           # File/resolve/dismiss
│   │   │   ├── commands.go                          # FlagReview, ResolveFlag, DismissFlag
│   │   │   ├── queries.go                           # GetFlags, GetFlag
│   │   │   ├── validators.go                        # Reason whitelist; reviewer role
│   │   │   ├── mapper.go                            # Flag ↔ DTO
│   │   │   └── dto.go                                # FlagDTO
│   │   ├── moderation/
│   │   │   ├── service.go                           # AutoFlag & SetState
│   │   │   ├── commands.go                          # AutoFlagReview, SetModerationState
│   │   │   ├── queries.go                           # GetModerationState, ListFlags
│   │   │   ├── validators.go                        # Threshold/transition checks
│   │   │   ├── mapper.go                            # Moderation ↔ DTO
│   │   │   └── dto.go                                # ModerationDTO
│   │   ├── appeal/
│   │   │   ├── service.go                           # Open/decide
│   │   │   ├── commands.go                          # OpenAppeal, DecideAppeal
│   │   │   ├── queries.go                           # GetAppeal, ListAppeals
│   │   │   ├── validators.go                        # Window & evidence checks
│   │   │   ├── mapper.go                            # Appeal ↔ DTO
│   │   │   └── dto.go                                # AppealDTO
│   │   ├── evidence/
│   │   │   ├── service.go                           # Link to storage-be
│   │   │   ├── commands.go                          # AttachEvidence, RemoveEvidence
│   │   │   ├── queries.go                           # ListEvidenceForReview
│   │   │   ├── validators.go                        # File kind/size; ownership
│   │   │   ├── mapper.go                            # Evidence ↔ DTO
│   │   │   └── dto.go                                # EvidenceDTO
│   │   ├── redaction/
│   │   │   ├── service.go                           # ApplyMask(review/response)
│   │   │   ├── validators.go                        # Allowlist/denylist compile
│   │   │   └── dto.go                                # RedactedTextDTO
│   │   ├── compliance/
│   │   │   ├── service.go                           # DSAR: erasure/export orchestration
│   │   │   ├── commands.go                          # RequestErasure, ExportData
│   │   │   ├── queries.go                           # GetTaskStatus
│   │   │   ├── validators.go                        # Subject ownership check
│   │   │   └── dto.go                                # ErasureTaskDTO, ExportTaskDTO
│   │   # ---- STATS & EXPOSURE ----
│   │   ├── stats/
│   │   │   ├── service.go                           # Rebuild/serve aggregates
│   │   │   ├── commands.go                          # RebuildStats
│   │   │   ├── queries.go                           # GetStats (by user/time window)
│   │   │   ├── mapper.go                            # Stats ↔ DTO
│   │   │   └── dto.go                                # StatsDTO
│   │   ├── featured/
│   │   │   ├── service.go                           # Editorial pin/unpin (policy-guarded)
│   │   │   ├── commands.go                          # FeatureReview, UnfeatureReview
│   │   │   ├── queries.go                           # ListFeatured
│   │   │   ├── validators.go                        # No paid bias; expiry in future
│   │   │   ├── mapper.go                            # Featured ↔ DTO
│   │   │   └── dto.go                                # FeaturedReviewDTO
│   │   # ---- AUDIT READS ----
│   │   ├── audit/                                   # 🆕 read-only views
│   │   │   ├── service.go                           # Fetch audit trail entries
│   │   │   └── queries.go                           # GetAuditTrailForReview
│   │
│   # =============================
│   # 🔌 INFRASTRUCTURE LAYER
│   # =============================
│   ├── infrastructure/
│   │   # =========================
│   │   # 🗄️ PERSISTENCE (POSTGRES)
│   │   # =========================
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       # 🧱 COMMON (DB BOOTSTRAP)
│   │   │       ├── connection.go                    # DSN & pooling (env from config)
│   │   │       ├── transaction.go                   # WithTransaction helpers
│   │   │       ├── migrations.go                    # 📝 Auto-migrate with versioning
│   │   │       ├── version.go                       # SchemaVersion table
│   │   │       ├── safety.go                        # Env & disk checks before migration
│   │   │       # ⭐ CORE & WINDOWS
│   │   │       ├── review_repository.go
│   │   │       ├── review_draft_repository.go       # 🆕
│   │   │       ├── eligibility_repository.go        # 🆕
│   │   │       ├── double_blind_repository.go
│   │   │       ├── review_reminder_repository.go
│   │   │       # 🧮 RATING / REPUTATION / BADGES
│   │   │       ├── rating_repository.go
│   │   │       ├── badge_repository.go
│   │   │       ├── user_badge_repository.go
│   │   │       ├── reputation_repository.go
│   │   │       # 🗣️ PRIVATE FEEDBACK
│   │   │       ├── feedback_repository.go
│   │   │       # 🛡️ SAFETY & GOVERNANCE
│   │   │       ├── flag_repository.go
│   │   │       ├── moderation_repository.go
│   │   │       ├── appeal_repository.go            # 🆕
│   │   │       ├── evidence_repository.go          # 🆕
│   │   │       ├── redaction_repository.go         # 🆕 policies
│   │   │       ├── compliance_repository.go        # 🆕 erasure/export tasks
│   │   │       ├── audit_trail_repository.go       # 🆕 immutable audit log
│   │   │       # 📊 STATS & EXPOSURE
│   │   │       ├── stats_repository.go
│   │   │       └── featured_repository.go          # 🆕
│   │   # =========================
│   │   # ⚡ CACHE (REDIS)
│   │   # =========================
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go                   # Pooling, retry, timeouts
│   │   │       ├── review_cache.go                 # Review by ID
│   │   │       ├── reputation_cache.go             # Reputation scores
│   │   │       ├── stats_cache.go                  # Aggregates by user
│   │   │       ├── double_blind_cache.go           # Window state
│   │   │       ├── eligibility_cache.go            # 🆕 Eligibility result (short TTL)
│   │   │       └── badge_cache.go                  # Badge lookups
│   │   # =========================
│   │   # 📨 MESSAGING (KAFKA)
│   │   # =========================
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go                     # 📝 platform-shared/inbox (dedupe, ack)
│   │   │       ├── producer.go                     # 📝 platform-shared/outbox (exactly-once-ish)
│   │   │       ├── topics.go                       # 📝 contracts/events: review.*, badge.*, reputation.*, compliance.*
│   │   │       └── scram.go                        # SASL/SCRAM-256 auth
│   │   # =========================
│   │   # 🔗 EXTERNAL CLIENTS
│   │   # =========================
│   │   └── storage/
│   │       └── client.go                           # Minimal HTTP client to storage-be for evidence files
│   │
│   # =============================
│   # 🌐 HTTP INTERFACE (v1)
│   # =============================
│   ├── interfaces/
│   │   └── http/
│   │       └── v1/
│   │           # =========================
│   │           # 🧭 HANDLERS
│   │           # =========================
│   │           ├── handlers/
│   │           # ---- CORE REVIEWS ----
│   │           │   ├── review_handler.go            # POST/PUT/DELETE/GET reviews + helpful vote
│   │           │   ├── draft_handler.go             # 🆕 save/clear/get drafts
│   │           │   ├── eligibility_handler.go       # 🆕 GET eligibility for contract
│   │           # ---- RATING/REPUTATION/BADGES ----
│   │           │   ├── rating_handler.go            # CRUD criteria; get aggregates
│   │           │   ├── reputation_handler.go        # GET score/history
│   │           │   ├── badge_handler.go             # CRUD badges; list
│   │           │   ├── user_badge_handler.go        # 🆕 list badges for user
│   │           # ---- DOUBLE-BLIND & REMINDERS ----
│   │           │   ├── double_blind_handler.go      # open window; submit; publish-if-ready
│   │           │   ├── reminder_handler.go          # schedule/send reminders
│   │           # ---- PRIVATE FEEDBACK ----
│   │           │   ├── feedback_handler.go          # submit/get private feedback
│   │           # ---- SAFETY & GOVERNANCE ----
│   │           │   ├── flag_handler.go              # flag/resolve/dismiss
│   │           │   ├── moderation_handler.go        # set state; list flags
│   │           │   ├── appeal_handler.go            # open/decide appeals
│   │           │   ├── evidence_handler.go          # attach/remove evidence files
│   │           │   ├── redaction_handler.go         # preview/apply redaction
│   │           │   ├── compliance_handler.go        # DSAR erasure/export requests
│   │           # ---- STATS & EXPOSURE ----
│   │           │   ├── stats_handler.go             # get aggregates
│   │           │   ├── featured_handler.go          # editorial features
│   │           # ---- MISC ----
│   │           │   └── health_handler.go            # /health, /ready, /live
│   │           # =========================
│   │           # 🧰 MIDDLEWARE
│   │           # =========================
│   │           ├── middleware/
│   │           │   ├── auth.go                      # JWT verify (pkg/auth)
│   │           │   ├── rbac.go                      # role checks
│   │           │   ├── cors.go                      # platform-shared/ginx
│   │           │   ├── rate_limit.go                # token bucket
│   │           │   ├── logging.go                   # structured logs
│   │           │   ├── error_handler.go             # error → HTTP
│   │           │   └── request_id.go                # X-Request-ID
│   │           # =========================
│   │           # 📨 RESPONSES
│   │           # =========================
│   │           ├── responses/
│   │           │   ├── success.go                   # platform-shared/httpx/response
│   │           │   ├── error.go                     # platform-shared/httpx/errors
│   │           │   └── pagination.go                # platform-shared/httpx/pagination
│   │           # =========================
│   │           # 🗺️ ROUTES (SECTIONED)
│   │           # =========================
│   │           ├── routes/
│   │           │   # ---- CORE REVIEWS ----
│   │           │   ├── review_routes.go             # /reviews/*
│   │           │   ├── draft_routes.go              # 🆕 /reviews/:id/draft/*
│   │           │   ├── eligibility_routes.go        # 🆕 /eligibility/*
│   │           │   # ---- RATING/REPUTATION/BADGES ----
│   │           │   ├── rating_routes.go             # /ratings/*
│   │           │   ├── reputation_routes.go         # /reputation/*
│   │           │   ├── badge_routes.go              # /badges/*
│   │           │   ├── user_badge_routes.go         # 🆕 /user-badges/*
│   │           │   # ---- DOUBLE-BLIND & REMINDERS ----
│   │           │   ├── double_blind_routes.go       # /double-blind/*
│   │           │   ├── reminder_routes.go           # /reminders/*
│   │           │   # ---- PRIVATE FEEDBACK ----
│   │           │   ├── feedback_routes.go           # /feedback/*
│   │           │   # ---- SAFETY & GOVERNANCE ----
│   │           │   ├── flag_routes.go               # /flags/*
│   │           │   ├── moderation_routes.go         # /moderation/*
│   │           │   ├── appeal_routes.go             # /appeals/*
│   │           │   ├── evidence_routes.go           # /evidence/*
│   │           │   ├── redaction_routes.go          # /redactions/*
│   │           │   ├── compliance_routes.go         # /compliance/*
│   │           │   # ---- STATS & EXPOSURE ----
│   │           │   ├── stats_routes.go              # /stats/*
│   │           │   └── featured_routes.go           # /featured/*
│   │           └── router.go                        # Gin engine + common middleware wiring
│   │
│   └── (end internal)
│
├── config/
│   ├── default.yaml                               # Defaults (ports, DSN templates, kafka/redis hosts)
│   ├── dev.yaml                                   # Dev overrides
│   └── prod.yaml                                  # Prod overrides (timeouts, quotas)
│
├── dapr/
│   ├── local/
│   │   ├── pubsub.yaml                            # Kafka pub/sub component
│   │   └── statestore.yaml                        # State store component
│   └── k8s/
│       ├── pubsub.yaml                            # Scopes: ["reviews-be"]
│       ├── statestore.yaml                        # Scopes
│       └── secrets.yaml                           # Secret store
│
├── pkg/
│   # =============================
│   # 🧰 LOCAL UTILITIES
│   # =============================
│   ├── errors/
│   │   ├── errors.go                              # Service-specific errors
│   │   └── codes.go                               # Error codes (e.g., REVIEW_NOT_FOUND)
│   ├── utils/
│   │   ├── validator.go                           # Reusable validations
│   │   ├── sanitizer.go                           # HTML sanitization; guard XSS
│   │   └── sentiment_analyzer.go                  # Optional sentiment signal for moderation
│   └── constants/
│       ├── badges.go                              # Badge ids/types/levels
│       └── README.md                              # Events/topics come from contracts/events
│
├── seeds/
│   ├── badges.sql                                 # Seed default badges & levels
│   ├── rating_criteria.sql                        # 🆕 Seed criteria & versions per category
│   └── top_rated_rules.sql                        # 🆕 Seed reputation thresholds
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
│   ├── seed-criteria.sh                           # 🆕 load criteria versions
│   └── recalculate-reputation.sh                  # Rebuild reputation for all users
│
├── tests/
│   # =============================
│   # ✅ TEST SUITES
│   # =============================
│   ├── unit/                                      # domain & services (eligibility, double_blind, rating, reputation, moderation)
│   ├── integration/                               # handlers + repositories (Postgres/Redis)
│   └── e2e/                                       # contract end → blind window → publish → search-be emit
│
├── docs/
│   ├── README.md                                  # Overview
│   ├── api.md                                     # HTTP API
│   ├── events.md                                  # 📝 review.*, badge.*, reputation.*, compliance.*
│   ├── rating-system.md                           # Criteria, weights, versioning
│   ├── reputation-algorithm.md                    # Bayesian + decay details
│   ├── badge-system.md                            # Badges, levels, rules
│   ├── double-blind.md                            # Windows & publish policy
│   ├── moderation.md                               # Abuse controls & flows
│   ├── compliance.md                               # DSAR/erasure/export
│   ├── MIGRATIONS.md                               # History of schema changes
│   ├── SCHEMA.md                                   # Database schema snapshot
│   └── RUNBOOK.md                                  # Ops (DLQ, replayer, rebuilds)
│
├── .github/
│   └── workflows/
│       ├── ci.yml                                 # Lint, tests, docker build
│       └── cd.yml                                 # Deploy to K8s (staging/prod)
│
├── go.mod                                         # 📝 imports pkg/auth, platform-shared, contracts/events
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md
