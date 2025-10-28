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

