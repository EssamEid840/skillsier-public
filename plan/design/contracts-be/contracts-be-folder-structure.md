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
