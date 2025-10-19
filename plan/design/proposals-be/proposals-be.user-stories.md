Global Conventions (applies to all sections)
============================================

**Event envelope (Kafka / outbox)**

Plain textANTLR4BashCC#CSSCoffeeScriptCMakeDartDjangoDockerEJSErlangGitGoGraphQLGroovyHTMLJavaJavaScriptJSONJSXKotlinLaTeXLessLuaMakefileMarkdownMATLABMarkupObjective-CPerlPHPPowerShell.propertiesProtocol BuffersPythonRRubySass (Sass)Sass (Scss)SchemeSQLShellSwiftSVGTSXTypeScriptWebAssemblyYAMLXML`   event_id (ULID), event_type, version, occurred_at (UTC), actor_id, tenant_id,  correlation_id, causation_id, partition_key, schema_ref (hash),  nonpii_payload (domain DTO; no raw files or secrets)   `

**Idempotent write-path**

*   All write handlers accept Idempotency-Key (UUIDv4). Server returns the original success payload on safe retries (TTL 24h).
    
*   Natural keys prevent duplicates (e.g., {freelancer\_id, job\_id} for proposal; {proposal\_id, file\_id} for attachments).
    
*   Outbox is _exactly-once_ per event\_id. Inbox performs dedupe on event\_id + causation\_id.
    

**Non-PII rules**

*   Never store raw emails/phones/IM handles inside free-text; lint + redact at write.
    
*   Files are **references only** to storage-be; all media undergo AV/DLP scan asynchronously.
    
*   Audit logs scrub sensitive text; indices keep only hashed/fingerprinted fields where needed.
    

**Platform alignment**

*   Follows provided folder structure (domain/application/interfaces layers).
    
*   Topics are grouped by feature (e.g., proposal.lifecycle.\*, proposal.auction.\*, proposal.intelligence.\*).
    
*   Projections are suffixed \_read and kept query-optimized; commands/queries in application layer.
    

1 - CORE PROPOSAL LIFECYCLE
==========================

1.1 proposal/
-------------

### Stories

*   As a **freelancer**, I want to create a proposal (job, pricing model, duration, availability) so that I can apply cleanly.
    
*   As a **freelancer**, I want submission guardrails (job open, not already applied, connects available) so that I don’t waste time.
    
*   As a **client**, I want to see structured proposals (pricing, milestones, Q&A) so that I can compare apples-to-apples.
    
*   As a **system**, I want dedupe within 24h for {freelancer\_id, job\_id} so that spam/accidental duplicates are prevented.
    
*   As a **system**, I want auto-expiry on job close so that stale proposals don’t linger.
    
*   As a **freelancer**, I want to withdraw a proposal with a reason so that I keep my pipeline clean.
    

### Flow

*   CreateProposalCommand(job\_id, pricing, duration, availability, template\_id?, invite\_id?)→ ValidateJobIsOpen() | CheckDuplicateWindow() | AntiPiiLint() | SeedFromTemplateOrInvite()→ CreateProposal() (tx: proposal + initial version + lifecycle = Draft)→ Outbox: proposal.core.created.v1
    
*   SubmitProposalCommand(proposal\_id)→ GuardSubmit(job\_is\_open, not\_duplicate, connects\_balance) | ConsumeConnects()→ MarkSubmitted() → Outbox: proposal.lifecycle.submitted.v1
    
*   WithdrawProposalCommand(proposal\_id, reason)→ GuardWithdraw() → MarkWithdrawn() → MaybeRefundConnects()→ Outbox: proposal.lifecycle.withdrawn.v1
    
*   ExpireProposalsForJobCommand(job\_id) (system)→ MarkExpiredBatch() → Outbox (per proposal): proposal.lifecycle.expired.v1
    

### Projections

*   proposals\_read (core fields + status)
    
*   proposals\_search\_read (normalized text sans PII)
    
*   proposal\_versions\_read (lightweight diffs)
    
*   proposal\_lifecycle\_read (stage history)
    

### Events

*   proposal.core.created.v1, proposal.lifecycle.submitted.v1, proposal.lifecycle.withdrawn.v1, proposal.lifecycle.expired.v1, proposal.core.duplicate.prevented.v1
    

### RBAC/SLO

*   Create/Submit/Withdraw: **OWNER**; Expire: **SYSTEM**.
    
*   P95 create < 250ms, submit < 300ms (includes connects check); event-to-index ≤ 1s.
    

1.2 cover\_letter/
------------------

### Stories

*   As a **freelancer**, I want to add/edit my cover letter with tone/length guardrails so that it’s effective.
    
*   As a **client**, I want version history so that I can see material changes.
    
*   As a **system**, I want to block PII and dangerous links in the letter so that policy is enforced.
    

### Flow

*   AddCoverLetterCommand(proposal\_id, content) → AntiPiiLint() | CheckLength() → SaveNewVersion()→ Outbox: proposal.cover\_letter.added.v1
    
*   EditCoverLetterCommand(proposal\_id, content) → GuardEditPolicy() | AntiPiiLint() → AppendVersion()→ Outbox: proposal.cover\_letter.edited.v1
    
*   RevertCoverLetterVersionCommand(proposal\_id, version) → GuardRevert() → AppendVersionFrom(version)→ Outbox: proposal.cover\_letter.version.reverted.v1
    

### Projections

*   cover\_letter\_read (current + version pointers)
    
*   cover\_letter\_versions\_read
    

### Events

*   proposal.cover\_letter.added.v1, proposal.cover\_letter.edited.v1, proposal.cover\_letter.version.reverted.v1
    

### RBAC/SLO

*   **OWNER**; P95 < 200ms; text ≤ 10k chars; idempotent by (proposal\_id, content\_hash).
    

1.3 proposal\_attachment/
-------------------------

### Stories

*   As a **freelancer**, I want to attach work samples (PDF, images, links) so that clients can evaluate my proof of work.
    
*   As a **client**, I want previews and file-type validation so that I don’t download unsafe or irrelevant files.
    
*   As a **system**, I want storage references only and asynchronous AV/DLP scanning so that compliance and security are enforced.
    
*   As a **freelancer**, I want to remove or replace attachments so that my proposal stays current.
    

### Flow

*   AddProposalAttachmentCommand(proposal\_id, storage\_file\_id, file\_name, file\_type, byte\_size?)→ ValidateStorageRef() | ValidateFileTypeWhitelist() | ValidateSizeLimit() | AntiPiiLint(filename)→ LinkAttachment() (tx: proposal\_attachment row + state=PendingScan)→ EnqueueMediaScan(storage\_file\_id)→ Outbox: proposal.attachment.added.v1
    
*   ReplaceProposalAttachmentCommand(proposal\_id, old\_file\_id, new\_storage\_file\_id, ...)→ GuardOwnership() | UnlinkAttachment(old) | LinkAttachment(new) | EnqueueMediaScan(new)→ Outbox: proposal.attachment.replaced.v1
    
*   RemoveProposalAttachmentCommand(proposal\_id, storage\_file\_id)→ GuardOwnership() | UnlinkAttachment()→ Outbox: proposal.attachment.removed.v1
    
*   MediaScanResultEvent(storage\_file\_id, verdict, engine\_meta) (from storage-be)→ UpdateAttachmentScanStatus(verdict) (Clean|Flagged|Quarantined)→ If Flagged/Quarantined then SoftUnlinkFromProposal()→ Outbox: proposal.attachment.scan.updated.v1
    

### Projections

*   proposal\_attachments\_read (per attachment: type, size, scan\_status, preview\_meta)
    
*   media\_scan\_read (file\_id → verdict timeline, engine metadata)
    
*   proposal\_attachment\_audit\_read
    

### Events

*   proposal.attachment.added.v1, proposal.attachment.removed.v1, proposal.attachment.replaced.v1, proposal.attachment.scan.updated.v1
    

### RBAC/SLO

*   **OWNER**; max 20 attachments/proposal; file ≤ 50 MB (configurable); add/remove/replace P95 < 180 ms; scan async P95 < 5 min; idempotent by (proposal\_id, storage\_file\_id).
    

1.4 proposal\_question\_answer/
-------------------------------

### Stories

*   As a **freelancer**, I want to answer screening questions (MCQ, boolean, short/long text, file) so that my application is complete.
    
*   As a **client**, I want normalized answers with types so that automated rubric scoring is possible.
    
*   As a **system**, I want redaction of free-text answers to remove emails/phones/handles so that policy is enforced.
    
*   As a **freelancer**, I want to update or remove an answer before submission so that I can correct mistakes.
    

### Flow

*   SubmitProposalAnswersCommand(proposal\_id, answers\[\])→ ValidateQuestionSet(job\_id) | ValidateTypes() | AntiPiiLint(textual\_answers) | ValidateAttachmentRefs()→ UpsertAnswers() (tx: insert-or-update per {proposal\_id, question\_id})→ Outbox: proposal.screening.answers.submitted.v1
    
*   UpdateProposalAnswerCommand(proposal\_id, question\_id, answer)→ GuardEditableWindow() | ValidateType() | AntiPiiLint()→ Update()→ Outbox: proposal.screening.answer.updated.v1
    
*   RemoveProposalAnswerCommand(proposal\_id, question\_id)→ GuardEditableWindow() | Delete()→ Outbox: proposal.screening.answer.removed.v1
    
*   BulkImportAnswersCommand(proposal\_id, source\_template\_id)→ LoadTemplate() | MapToQuestions() | UpsertAnswers()→ Outbox: proposal.screening.answers.imported.v1
    

### Projections

*   proposal\_screening\_read (typed answers + redaction flags)
    
*   proposal\_screening\_index\_read (denormalized for rubric queries; no PII)
    
*   proposal\_screening\_audit\_read
    

### Events

*   proposal.screening.answers.submitted.v1, proposal.screening.answer.updated.v1, proposal.screening.answer.removed.v1, proposal.screening.answers.imported.v1
    

### RBAC/SLO

*   **OWNER**; per-proposal ≤ 25 questions; submit/update/remove P95 < 220 ms; idempotent by (proposal\_id, question\_id, answer\_hash).
    

1.5 milestone/
--------------

### Stories

*   As a **freelancer**, I want to define milestones (desc, amount, due) so that scope and billing are clear.
    
*   As a **client**, I want validation that sum(milestones) matches fixed-price total so that expectations align.
    
*   As a **system**, I want milestones compatible with contracts/payments so that downstream flows are seamless.
    
*   As a **freelancer**, I want to reorder or remove milestones before submission so that my plan is coherent.
    

### Flow

*   CreateMilestoneCommand(proposal\_id, description, amount, due?)→ ValidatePricingModel(fixed\_or\_hybrid) | ValidateAmount() | AntiPiiLint(description)→ CreateMilestone(seq) | ValidateSumAgainstPricing()→ Outbox: proposal.milestone.created.v1
    
*   UpdateMilestoneCommand(milestone\_id, patch)→ GuardEditableWindow() | ValidatePatch() | RecalcTotals()→ Outbox: proposal.milestone.updated.v1
    
*   ReorderMilestonesCommand(proposal\_id, new\_order\[\])→ ValidatePermutation() | ApplyOrder()→ Outbox: proposal.milestone.reordered.v1
    
*   RemoveMilestoneCommand(milestone\_id)→ GuardEditableWindow() | Delete() | RecalcTotals()→ Outbox: proposal.milestone.removed.v1
    

### Projections

*   proposal\_milestones\_read (ordered list with totals)
    
*   proposal\_pricing\_read (pricing model, totals, currency)
    
*   milestone\_audit\_read
    

### Events

*   proposal.milestone.created.v1, proposal.milestone.updated.v1, proposal.milestone.reordered.v1, proposal.milestone.removed.v1
    

### RBAC/SLO

*   **OWNER**; max 20 milestones; create/update/remove P95 < 200 ms; idempotent by (proposal\_id, description\_hash, amount, due) when creating.
    

2 - BIDDING SYSTEM
=================

2.1 bid/
--------

### Stories

*   As a **freelancer**, I want to place/update/withdraw a bid so that I can compete dynamically.
    
*   As a **client**, I want transparent ranking and distribution stats so that I quickly gauge competitiveness.
    
*   As a **system**, I want a full bid history with currency normalization so that analytics and anomaly detection are possible.
    
*   As a **freelancer**, I want hard floors/ceilings per job so that I don’t breach client constraints.
    

### Flow

*   PlaceBidCommand(proposal\_id, amount, currency)→ ValidateJobIsOpen() | NormalizeCurrency(amount,currency,job\_currency) | ValidateMinMax(job\_rules) | ValidateSingleActiveBid(proposal)→ UpsertBid(state=Active) | RecomputeRank(job\_id)→ Outbox: bid.placed.v1 (+ if rank changed: bid.rank.changed.v1)
    
*   UpdateBidCommand(bid\_id, amount, currency)→ ValidateOwnership() | NormalizeCurrency() | ValidateMinMax()→ UpdateBid() | RecomputeRank()→ Outbox: bid.updated.v1 (+ bid.rank.changed.v1?)
    
*   WithdrawBidCommand(bid\_id, reason?)→ GuardWithdrawState() | MarkRetracted() | RecomputeRank()→ Outbox: bid.retracted.v1 (+ bid.rank.changed.v1?)
    
*   SyncBidFromInviteCommand(invite\_id) (if client sets target rate)→ MapInviteToProposal() | SeedOrUpdateBid() | RecomputeRank()→ Outbox: bid.synced\_from\_invite.v1
    

### Projections

*   bid\_read (current active amount, normalized)
    
*   bid\_history\_read (timeline, deltas)
    
*   bid\_rank\_read (rank, spread, percentiles)
    

### Events

*   bid.placed.v1, bid.updated.v1, bid.retracted.v1, bid.rank.changed.v1, bid.synced\_from\_invite.v1
    

### RBAC/SLO

*   **OWNER (proposal)**; place/update/withdraw P95 < 160 ms; rank recompute atomic in-tx; idempotent by (proposal\_id, amount\_norm, ts\_bucket).
    

2.2 bid\_strategy/
------------------

### Stories

*   As a **freelancer**, I want strategies (fixed, step-down, undercut, auto-rebalance) so that my bid adapts to competition.
    
*   As a **system**, I want guardrails (min floor, max daily changes, cool-downs) so that pricing remains stable.
    
*   As a **freelancer**, I want a preview/simulation so that I understand strategy impact before enabling.
    

### Flow

*   CreateBidStrategyCommand(proposal\_id, type, params)→ ValidateParamsByType() | GuardCompatibilityWithJob()→ CreateStrategy(state=Draft)→ Outbox: bid.strategy.created.v1
    
*   SimulateBidStrategyCommand(strategy\_id, horizon\_days)→ LoadMarketSignals(job\_id) | RunSimulation()→ Outbox: bid.strategy.simulated.v1
    
*   ActivateBidStrategyCommand(strategy\_id)→ GuardNoActiveStrategy() | Activate()→ Outbox: bid.strategy.activated.v1
    
*   AutoAdjustBidsJob() (scheduler)→ SelectActiveStrategies() | ComputeNextAmount() | ApplyIfWithinGuardrails() | RecomputeRank()→ Outbox: bid.auto.adjusted.v1 (+ bid.rank.changed.v1?)
    
*   DeactivateBidStrategyCommand(strategy\_id)→ Deactivate()→ Outbox: bid.strategy.deactivated.v1
    

### Projections

*   bid\_strategy\_read (config & state)
    
*   bid\_strategy\_state\_read (last\_run, changes\_today, guardrail\_hits)
    
*   bid\_strategy\_simulations\_read
    

### Events

*   bid.strategy.created.v1, bid.strategy.simulated.v1, bid.strategy.activated.v1, bid.auto.adjusted.v1, bid.strategy.deactivated.v1
    

### RBAC/SLO

*   **OWNER**; control ops P95 < 150 ms; auto-job batch 10k < 60 s; idempotent by (strategy\_id, tick).
    

2.3 bid\_notification/
----------------------

### Stories

*   As a **freelancer**, I want outbid alerts with threshold options so that I react only when it matters.
    
*   As a **system**, I want reliable queueing, exponential backoff, and per-channel fallbacks so that notifications are dependable.
    
*   As a **freelancer**, I want quiet hours so that I’m not spammed.
    

### Flow

*   OnRankChangedEvent(job\_id, proposal\_id, old\_rank, new\_rank)→ DetectOutbid(thresholds, quiet\_hours) | EnqueueNotification(channel\_pref)→ Outbox: bid.notification.queued.v1
    
*   Worker SendNotificationTask(notification\_id)→ Send(channel) | RecordDelivery() | RetryOnFailure()→ Outbox: bid.notification.sent.v1 | bid.notification.failed.v1
    
*   UpdateNotificationPrefsCommand(user\_id, prefs)→ ValidatePrefs() | Save()→ Outbox: bid.notification.prefs.updated.v1
    

### Projections

*   bid\_notification\_read (status, attempts, channel)
    
*   bid\_notification\_prefs\_read
    

### Events

*   bid.outbid.alert.triggered.v1, bid.notification.queued.v1, bid.notification.sent.v1, bid.notification.failed.v1, bid.notification.prefs.updated.v1
    

### RBAC/SLO

*   Trigger: **SYSTEM**; prefs: **OWNER**; P99 send < 5 s; at-least-once with dedupe on message key; idempotent by (user\_id, proposal\_id, rank\_change\_bucket).
    

2.4 auction/ (boosted top slots)
--------------------------------

### Stories

*   As a **freelancer**, I want to bid for limited top visibility slots so that my proposal is seen first.
    
*   As a **client**, I want fair windows, reserve price, and min increments so that the auction is credible.
    
*   As a **system**, I want automatic slot assignment and refunds on failed bids so that balances stay accurate.
    

### Flow

*   OpenAuctionWindowCommand(job\_id, slots, reserve, min\_increment, ttl)→ ValidateJob() | CreateAuction()→ Outbox: auction.opened.v1
    
*   PlaceAuctionBidCommand(job\_id, proposal\_id, amount)→ GuardWindowOpen() | ValidateIncrement() | HoldFunds(connects\_or\_currency) | UpsertAuctionBid() | RecalcTopSlots()→ Outbox: auction.bid.placed.v1 (+ auction.top.changed.v1?)
    
*   CancelAuctionBidCommand(auction\_bid\_id)→ GuardCancelable() | ReleaseHold() | MarkCanceled() | RecalcTopSlots()→ Outbox: auction.bid.canceled.v1 (+ auction.top.changed.v1?)
    
*   EndAuctionCommand(job\_id)→ GuardEnd() | CloseAndAssignSlots() | CaptureWinningHolds() | ReleaseLosingHolds()→ Outbox: auction.ended.v1, auction.top.assigned.v1
    

### Projections

*   auction\_status\_read (open/closed, config, countdown)
    
*   auction\_bids\_read (leaderboard)
    
*   auction\_slot\_assignment\_read
    

### Events

*   auction.opened.v1, auction.bid.placed.v1, auction.bid.canceled.v1, auction.top.changed.v1, auction.ended.v1, auction.top.assigned.v1
    

### RBAC/SLO

*   Open/End: **SYSTEM**; Place/Cancel: **OWNER**; P95 place < 170 ms; reconciliation job P95 < 2 s; idempotent by (job\_id, proposal\_id, amount, ts\_bucket).
    

2.5 bid\_anomaly\_detection/
----------------------------

### Stories

*   As a **trust analyst**, I want anomalies flagged (suspicious undercut, collusion patterns) so that marketplace quality is preserved.
    
*   As a **system**, I want triage states (Open, UnderReview, Closed) and evidence so that reviews are trackable.
    

### Flow

*   DetectBidAnomalyCommand(job\_id)→ FetchBidDistribution() | ComputeZScores()/IQR() | RunCollusionHeuristics()→ UpsertAnomalies()→ Outbox: bid.anomaly.detected.v1
    
*   UpdateAnomalyReviewCommand(anomaly\_id, state, notes)→ ValidateStateTransition() | Update()→ Outbox: bid.anomaly.review.updated.v1
    
*   AutoCloseBenignAnomaliesJob()→ ApplyAutoRules() | Close()→ Outbox: bid.anomaly.auto\_closed.v1
    

### Projections

*   bid\_anomaly\_read (score, features, state)
    
*   bid\_anomaly\_audit\_read
    

### Events

*   bid.anomaly.detected.v1, bid.anomaly.review.updated.v1, bid.anomaly.auto\_closed.v1
    

### RBAC/SLO

*   Detect: **SYSTEM**; Review: **TRUST/ADMIN**; detect batch 50k < 90 s; idempotent by (job\_id, proposal\_id, feature\_hash, day).
    

3 - CONNECTS & BOOST
===================

3.1 connect/
------------

### Stories

*   As a **freelancer**, I want connects reserved/debited on submission so that platform economics are enforced.
    
*   As a **system**, I want tiered costs by job popularity and time-of-day so that pricing is dynamic and fair.
    
*   As a **freelancer**, I want balance, aging buckets, and ledger so that I can plan submissions.
    

### Flow

*   ReserveConnectsOnSubmitCommand(job\_id, freelancer\_id)→ GetTierCost(job\_id) | CheckBalance() | CreateReservation(hold)→ Outbox: connects.reserved.v1
    
*   FinalizeSubmissionCommand(proposal\_id)→ ConsumeReservation() | DebitLedger()→ Outbox: connects.debited.v1
    
*   CancelSubmissionCommand(proposal\_id)→ ReleaseReservation()→ Outbox: connects.released.v1
    
*   CreditConnectsCommand(user\_id, amount, reason)→ CreditLedger()→ Outbox: connects.credited.v1
    

### Projections

*   connect\_balance\_read (available, holds, aging)
    
*   connect\_transactions\_read (ledger)
    
*   connect\_tier\_read (job-tier, dynamic rules)
    

### Events

*   connects.reserved.v1, connects.debited.v1, connects.released.v1, connects.credited.v1, connects.tier.changed.v1
    

### RBAC/SLO

*   **OWNER**; reserve P95 < 120 ms; strong idempotency on (freelancer\_id, job\_id); ledger ACID.
    

3.2 connect\_refund/
--------------------

### Stories

*   As a **freelancer**, I want automatic and manual refunds (job closed early, spam job) so that I’m protected.
    
*   As a **system**, I want eligibility policies and SLA-based processing so that fairness and throughput are balanced.
    

### Flow

*   RequestConnectRefundCommand(proposal\_id, reason)→ CheckEligibility(policy\_matrix) | OpenCase(state=Pending)→ Outbox: connect.refund.requested.v1
    
*   AutoEvaluateRefundsJob()→ ApplyPolicy() | ApproveOrEscalate()→ Outbox: connect.refund.auto\_processed.v1
    
*   ProcessConnectRefundCommand(case\_id, decision)→ ApplyDecision() | CreditLedgerIfApproved()→ Outbox: connect.refund.processed.v1 | connect.refund.denied.v1
    

### Projections

*   connect\_refund\_read (case, decision, auditor)
    
*   connect\_refund\_policy\_read
    

### Events

*   connect.refund.requested.v1, connect.refund.auto\_processed.v1, connect.refund.processed.v1, connect.refund.denied.v1
    

### RBAC/SLO

*   Request: **OWNER**; Process: **SUPPORT/TRUST**; auto-job P95 case < 2 s; idempotent by (proposal\_id, policy\_snapshot\_hash).
    

3.3 boost/
----------

### Stories

*   As a **freelancer**, I want to purchase/upgrade boosts (levels) so that my proposal appears higher.
    
*   As a **system**, I want expiry handling and conflict resolution with auctions so that boosts don’t persist or double-count.
    

### Flow

*   PurchaseBoostCommand(proposal\_id, level, term)→ ChargeWallet() | ActivateBoost(level, expires\_at)→ Outbox: proposal.boost.purchased.v1
    
*   UpgradeBoostCommand(proposal\_id, new\_level)→ GuardUpgradePath() | ProrateCharge() | ApplyUpgrade()→ Outbox: proposal.boost.upgraded.v1
    
*   ExpireBoostsJob()→ FindExpired() | MarkExpired() | RecomputeRanking()→ Outbox: proposal.boost.expired.v1, proposal.ranking.updated.v1
    

### Projections

*   proposal\_boost\_read (level, expires\_at)
    
*   boost\_ledger\_read
    

### Events

*   proposal.boost.purchased.v1, proposal.boost.upgraded.v1, proposal.boost.expired.v1
    

### RBAC/SLO

*   **OWNER**; purchase/upgrade P95 < 180 ms; idempotent by (proposal\_id, level, purchase\_ts\_bucket).
    

4 - TEMPLATES & RATE CARDS
=========================

4.1 template/
-------------

### Stories

*   As a **freelancer**, I want reusable templates with categories and placeholders so that I draft faster.
    
*   As a **system**, I want placeholder validation and localization so that personalization is correct.
    
*   As a **freelancer**, I want version history and archiving so that I can maintain quality over time.
    

### Flow

*   CreateTemplateCommand(user\_id, title, content, categories\[\], placeholders\[\])→ ValidatePlaceholders(content) | AntiPiiLint(content)→ CreateTemplate(state=Active, version=1)→ Outbox: proposal.template.created.v1
    
*   UpdateTemplateCommand(template\_id, patch)→ GuardOwnership() | ValidatePlaceholders() | AppendVersion()→ Outbox: proposal.template.updated.v1
    
*   ArchiveTemplateCommand(template\_id)→ Archive()→ Outbox: proposal.template.archived.v1
    
*   LocalizeTemplateCommand(template\_id, locale)→ TranslateContent() | PersistVariant()→ Outbox: proposal.template.localized.v1
    

### Projections

*   template\_read (current version)
    
*   template\_versions\_read
    
*   template\_category\_read
    
*   template\_localization\_read
    

### Events

*   proposal.template.created.v1, proposal.template.updated.v1, proposal.template.archived.v1, proposal.template.localized.v1
    

### RBAC/SLO

*   **OWNER**; create/update P95 < 160 ms; idempotent by (user\_id, title\_hash, content\_hash).
    

4.2 rate\_card/
---------------

### Stories

*   As a **freelancer**, I want package tiers (Starter/Standard/Premium) with scope/delivery SLAs so that pricing is consistent.
    
*   As a **client**, I want clear inclusions/exclusions so that expectations are set.
    
*   As a **system**, I want currency normalization and taxes/fees modeling so that downstream quotes are accurate.
    

### Flow

*   CreateRateCardCommand(user\_id, currency, packages\[\])→ ValidatePackageSchema() | NormalizeCurrency()→ CreateRateCard(state=Active, version=1)→ Outbox: proposal.rate\_card.created.v1
    
*   UpdateRateCardCommand(card\_id, patch)→ GuardOwnership() | Validate() | AppendVersion()→ Outbox: proposal.rate\_card.updated.v1
    
*   ArchiveRateCardCommand(card\_id)→ Archive()→ Outbox: proposal.rate\_card.archived.v1
    

### Projections

*   rate\_card\_read (active version)
    
*   rate\_card\_versions\_read
    

### Events

*   proposal.rate\_card.created.v1, proposal.rate\_card.updated.v1, proposal.rate\_card.archived.v1
    

### RBAC/SLO

*   **OWNER**; P95 < 170 ms; idempotent by (user\_id, currency, packages\_hash).
    

5 - ANALYTICS & TRACKING
=======================

5.1 proposal\_analytics/
------------------------

### Stories

*   As a **freelancer**, I want views/CTR/response rate so that I can iterate.
    
*   As a **system**, I want view beacons with bot filtering so that stats are trustworthy.
    

### Flow

*   TrackProposalViewCommand(proposal\_id, session\_id, ua, ip\_hash)→ DedupWithin(30m) | BotFilter(ua,ip) | RecordView()→ Outbox: proposal.viewed.v1
    
*   RecomputeAnalyticsCommand(proposal\_id|job\_id|user\_id)→ AggregateViewsClicksResponses()→ Outbox: proposal.analytics.updated.v1
    

### Projections

*   proposal\_analytics\_read (views, unique\_views, CTR)
    
*   response\_stats\_read (client response, time-to-first-view)
    
*   analytics\_audit\_read
    

### Events

*   proposal.viewed.v1, proposal.analytics.updated.v1
    

### RBAC/SLO

*   Track: **PUBLIC** signed URL; recompute: **SYSTEM**; view P95 < 80 ms.
    

5.2 response\_tracker/
----------------------

### Stories

*   As a **freelancer**, I want to know if a client viewed/responded so that I time my follow-ups.
    
*   As a **client**, I want passive tracking (no extra clicks) so that workflow is smooth.
    

### Flow

*   TrackClientViewCommand(proposal\_id, viewer\_context)→ ValidateViewer(client\_team) | Record()→ Outbox: proposal.viewed.by\_client.v1
    
*   TrackClientResponseCommand(proposal\_id, response\_type) (message, request\_interview, reject)→ Record()→ Outbox: client.responded.v1
    

### Projections

*   response\_tracker\_read (per-proposal client interactions)
    
*   client\_response\_latency\_read
    

### Events

*   proposal.viewed.by\_client.v1, client.responded.v1
    

### RBAC/SLO

*   **SYSTEM/CLIENT**; P95 < 120 ms; idempotent by (proposal\_id, type, day).
    

5.3 insight/
------------

### Stories

*   As a **freelancer**, I want insights (competing bids count, average bid, engagement) so that I adapt strategy.
    
*   As a **system**, I want upsert logic to recompute on new signals.
    

### Flow

*   UpsertInsightCommand(proposal\_id)→ CollectSignals(bids, views, shortlist) | ComputeKPIs() | Upsert()→ Outbox: proposal.insight.updated.v1
    
*   AccessInsightCommand(proposal\_id)→ AuthorizeOwner() | LogAccess()→ Outbox: proposal.insight.accessed.v1
    

### Projections

*   proposal\_insight\_read (KPIs + trend)
    

### Events

*   proposal.insight.updated.v1, proposal.insight.accessed.v1
    

### RBAC/SLO

*   **OWNER**; compute P95 < 300 ms; cache 15 min.
    

5.4 ranking/
------------

### Stories

*   As a **system**, I want a visibility score (relevance, boost, CTR, response time, profile strength) so that listings are fair.
    
*   As a **freelancer**, I want to see my rank/score so that I can decide on edits/boosts.
    

### Flow

*   RecomputeRankingCommand(proposal\_id)→ FetchFactors() | WeightedScore() | PersistScore()→ Outbox: proposal.ranked.v1
    
*   BulkRecomputeRankingJob(job\_id)→ RecomputeAll()→ Outbox: proposal.ranking.updated.v1 (per proposal)
    

### Projections

*   proposal\_ranking\_read (score, factors)
    

### Events

*   proposal.ranked.v1, proposal.ranking.updated.v1
    

### RBAC/SLO

*   Recompute: **SYSTEM**; read: **OWNER**; P95 < 200 ms.
    

6 - PROPOSAL INTELLIGENCE & STRATEGY
===================================

6.1 proposal\_strategy/
-----------------------

### Stories

*   As a **freelancer**, I want model-backed strategies (pricing/timing/differentiation) so that win odds improve.
    
*   As a **system**, I want tactic application + outcome tracking so that learning loops exist.
    

### Flow

*   GenerateStrategyCommand(proposal\_id, type)→ GuardModelAccess() | CallStrategyModel() | Redact()→ Store(strategy, version)→ Outbox: proposal.strategy.generated.v1
    
*   ApplyStrategyCommand(strategy\_id)→ ApplyTactics(content/bid/boost) | TrackAppliedDelta()→ Outbox: proposal.strategy.applied.v1
    
*   TrackOutcomeCommand(strategy\_id, outcome)→ UpdateOutcome()→ Outbox: proposal.strategy.outcome.tracked.v1
    

### Projections

*   proposal\_strategy\_read (recommendations, diffs)
    
*   proposal\_strategy\_outcome\_read
    

### Events

*   proposal.strategy.generated.v1, proposal.strategy.applied.v1, proposal.strategy.outcome.tracked.v1
    

### RBAC/SLO

*   **OWNER**; gen P95 < 2.5 s (async permitted).
    

6.2 proposal\_intelligence/
---------------------------

### Stories

*   As a **freelancer**, I want competitor count, pricing bands, and client hiring patterns so that I calibrate bids.
    
*   As a **system**, I want win-probability updates as signals change.
    

### Flow

*   GenerateIntelligenceCommand(job\_id, freelancer\_id)→ AggregateMarketSignals() | ComputeWinProbability() | Upsert()→ Outbox: proposal.intelligence.generated.v1, proposal.win\_probability.calculated.v1
    
*   RefreshIntelligenceJob()→ RecomputeOnNewSignals()→ Outbox: proposal.intelligence.refreshed.v1
    

### Projections

*   proposal\_intelligence\_read
    

### Events

*   proposal.intelligence.generated.v1, proposal.intelligence.refreshed.v1, proposal.win\_probability.calculated.v1
    

### RBAC/SLO

*   **OWNER**; compute P95 < 600 ms; cache 30 min.
    

6.3 proposal\_optimization/
---------------------------

### Stories

*   As a **freelancer**, I want real-time suggestions (keywords, tone, readability, CTA) so that content quality improves.
    
*   As a **system**, I want apply/diff flows with measurable uplift so that ROI is visible.
    

### Flow

*   OptimizeProposalCommand(proposal\_id, content\_delta)→ Analyze(content) | SuggestEdits() | ScoreBeforeAfter()→ Outbox: proposal.optimization.suggested.v1
    
*   ApplyOptimizationCommand(proposal\_id, suggestion\_id)→ Apply() | Re-score()→ Outbox: proposal.optimization.applied.v1, proposal.score.improved.v1
    

### Projections

*   proposal\_optimization\_read (suggestions, scores)
    

### Events

*   proposal.optimization.suggested.v1, proposal.optimization.applied.v1, proposal.score.improved.v1
    

### RBAC/SLO

*   **OWNER**; suggest P95 < 800 ms; idempotent by (proposal\_id, suggestion\_hash).
    

6.4 proposal\_portfolio\_selector/
----------------------------------

### Stories

*   As a **freelancer**, I want auto-selected portfolio items relevant to the job so that proof is targeted.
    
*   As a **system**, I want explainable relevance so that the user trusts the selection.
    

### Flow

*   SelectPortfolioCommand(proposal\_id)→ ComputeRelevance(job\_skills, portfolio\_tags) | PickTopK()→ AttachLinks()→ Outbox: proposal.portfolio.selected.v1
    
*   OverrideSelectionCommand(proposal\_id, item\_ids\[\])→ ApplyOverride()→ Outbox: proposal.portfolio.overridden.v1
    

### Projections

*   portfolio\_selection\_read (items + relevance reasons)
    

### Events

*   proposal.portfolio.selected.v1, proposal.portfolio.overridden.v1, proposal.portfolio.auto\_attached.v1
    

### RBAC/SLO

*   **OWNER**; P95 < 250 ms.
    

7 - CLIENT INTERACTION & ENGAGEMENT
==================================

7.1 proposal\_conversation/
---------------------------

### Stories

*   As a **client/freelancer**, I want pre-hire messaging so that scope is clear.
    
*   As a **system**, I want sentiment/response-time metrics so that engagement health is visible.
    
*   As a **platform**, I want link/PII moderation and rate limits so that safety is preserved.
    

### Flow

*   SendMessageCommand(proposal\_id, message)→ RateLimit(user) | AntiPiiLint() | ModerateLinks()→ AppendMessage()→ Outbox: proposal.conversation.message.sent.v1
    
*   AnalyzeSentimentJob()→ Score(message\_batch)→ Outbox: proposal.conversation.sentiment.scored.v1
    
*   TrackResponseTimeCommand(proposal\_id, actor)→ UpdateSLA()→ Outbox: proposal.conversation.response\_time.tracked.v1
    

### Projections

*   proposal\_conversation\_read (thread, participants)
    
*   proposal\_sentiment\_read (scores over time)
    
*   proposal\_response\_time\_read (SLA counters)
    

### Events

*   proposal.conversation.message.sent.v1, proposal.conversation.sentiment.scored.v1, proposal.conversation.response\_time.tracked.v1
    

### RBAC/SLO

*   Participants only; send P95 < 130 ms; 20 msgs/hr per party; idempotent by (thread\_id, content\_hash, ts\_bucket).
    

7.2 proposal\_engagement/
-------------------------

### Stories

*   As a **freelancer**, I want to see micro-interactions (expand, scroll depth, portfolio clicks) so that I infer interest.
    
*   As a **system**, I want an interest score and follow-up triggers so that nudges are timely.
    

### Flow

*   TrackEngagementCommand(proposal\_id, event\_type, meta)→ Validate() | Dedup() | Record()→ Outbox: proposal.engagement.interaction.tracked.v1
    
*   CalculateInterestCommand(proposal\_id)→ ComputeScore() (weighted micro-events)→ Outbox: proposal.engagement.interest.scored.v1
    
*   TriggerHighInterestFollowUpCommand(proposal\_id) (policy-driven)→ GuardCoolDown() | ScheduleFollowUp()→ Outbox: proposal.engagement.high\_interest.detected.v1
    

### Projections

*   proposal\_engagement\_read (events)
    
*   interest\_score\_read (score, reasons)
    

### Events

*   proposal.engagement.interaction.tracked.v1, proposal.engagement.interest.scored.v1, proposal.engagement.high\_interest.detected.v1
    

### RBAC/SLO

*   Track: **SYSTEM**; reads: **OWNER/CLIENT**; track P95 < 120 ms.
    

7.3 proposal\_follow\_up/
-------------------------

### Stories

*   As a **freelancer**, I want scheduled follow-ups and gentle nudges so that I stay top-of-mind.
    
*   As a **system**, I want quiet hours and frequency caps so that clients aren’t spammed.
    

### Flow

*   ScheduleFollowUpCommand(proposal\_id, when, template\_id)→ ValidateWindow() | ApplyQuietHours() | PersistSchedule()→ Outbox: proposal.follow\_up.scheduled.v1
    
*   Worker SendFollowUpTask(schedule\_id)→ RenderTemplate() | Send() | Record()→ Outbox: proposal.follow\_up.sent.v1
    
*   BumpProposalCommand(proposal\_id)→ GuardBumpRules() | UpdateVisibility()→ Outbox: proposal.bumped.v1
    

### Projections

*   proposal\_follow\_up\_read (history)
    
*   proposal\_follow\_up\_schedule\_read
    

### Events

*   proposal.follow\_up.scheduled.v1, proposal.follow\_up.sent.v1, proposal.bumped.v1, client.nudged.v1
    

### RBAC/SLO

*   **OWNER**; schedule P95 < 140 ms; send P95 < 2 s; idempotent by (proposal\_id, schedule\_slot).

8 - PROPOSAL LIFECYCLE MANAGEMENT
================================

8.1 proposal\_pipeline/
-----------------------

### Stories

*   As a **freelancer**, I want a **Kanban-style pipeline** (Draft → Submitted → Viewed → Interview → Offer → Hired/Declined) with **per-stage WIP limits** so that I can prioritize my effort effectively.
    
*   As a **freelancer**, I want **bulk stage moves** (e.g., multi-select to Archive/Decline) so that I can tidy my funnel quickly.
    
*   As a **system**, I want **auto-moves on signals** (client viewed, interview scheduled, offer sent) so that the pipeline stays accurate without manual updates.
    
*   As a **manager** (team account), I want **aggregate pipeline and per-member views** so that I can coach performance and balance workload.
    
*   As a **system**, I want **SLA nudges** (e.g., “no client response for 7 days”) so that follow-ups are triggered automatically before opportunities go cold.
    

### Flow

1.  GetPipelineQuery(user\_id, filters) → ResolveOwnerAndTeamScopes() | LoadStagesAndCounts() | ComputeWIPOverages() → Return(lanes, counts, WIP flags)
    
2.  MoveStageCommand(proposal\_id, stage, reason?) → ValidateTransitionMatrix() | CheckTerminalLock() | ApplyStage() | AppendLifecycleNote(reason?) → **Outbox:** proposal.stage.moved.v1
    
3.  BulkMoveStageCommand(proposal\_ids\[\], stage, reason?) → ValidateBatch() | ApplyStageTransactional()(partial-fail tolerant) → **Outbox** (per move): proposal.stage.moved.v1
    
4.  AutoStageTransitionJob() → ConsumeSignals(view, reply, interview, offer) | ApplyRules() | If terminal → MarkTerminal() → **Outbox:** proposal.stage.auto\_moved.v1, proposal.stage.completed.v1
    
5.  PipelineWipSlaCheckJob() → DetectOverWIP() | DetectStaleInStage() | EmitNudges() → **Outbox:** proposal.pipeline.sla.nudged.v1
    

### Projections

*   proposal\_pipeline\_read (proposal\_id, stage, since\_at, reason?, owner\_id, team\_id, wip\_flags) \[idx: owner\_id+stage, team\_id+stage, since\_at DESC\]
    
*   pipeline\_analytics\_read (owner\_id/team\_id, stage\_counts, avg\_time\_in\_stage, conversion\_rates, stale\_count)
    
*   pipeline\_sla\_read (proposal\_id, last\_activity\_at, nudge\_due\_at, nudge\_state)
    

### Events

proposal.stage.moved.v1, proposal.stage.auto\_moved.v1, proposal.stage.completed.v1, proposal.pipeline.conversion.calculated.v1, proposal.pipeline.sla.nudged.v1

### RBAC/SLO

*   **RBAC:** **OWNER** for single; **TEAM\_LEAD** for bulk/team; **SYSTEM** for auto.
    
*   **SLO:** **P95 move < 180 ms**; bulk 100 items **< 1.2 s**; analytics refresh **≤ 5 s**; event→index **≤ 1 s**.
    
*   **Idempotency:** (proposal\_id, target\_stage, reason\_hash, Idempotency-Key).
    

8.2 proposal\_recycling/
------------------------

### Stories

*   As a **freelancer**, I want to **recycle or adapt past proposals** to similar jobs so that I can submit high-quality drafts faster.
    
*   As a **system**, I want **similarity matching and PII redaction** so that reuse is relevant and safe.
    
*   As a **freelancer**, I want a **change log of adaptations** so that I can track what was altered and why.
    
*   As a **manager**, I want **freshness and policy gates** so that outdated content isn’t reused without review.
    

### Flow

1.  RecycleProposalCommand(original\_proposal\_id, new\_job\_id, keep\_sections\[\]) → FindSimilarJobSections() | RedactPII() | CreateDraftVersion() → **Outbox:** proposal.recycled.v1
    
2.  AdaptRecycledProposalCommand(proposal\_id, deltas) → ApplyDeltas() | AppendVersion() → **Outbox:** proposal.adapted.v1
    

### Projections

*   proposal\_recycling\_read
    
*   proposal\_versions\_read
    

### Events

proposal.recycled.v1, proposal.adapted.v1, proposal.version.created.v1

### RBAC/SLO

*   **RBAC:** **OWNER**; policy override **TEAM\_LEAD**.
    
*   **SLO:** recycle **P95 < 220 ms**.
    
*   **Idempotency:** (original\_proposal\_id, new\_job\_id, keep\_sections\_hash).
    

8.3 proposal\_calendar/
-----------------------

### Stories

*   As a **freelancer**, I want **proposal deadlines and reminder schedules** so that I never miss submission windows.
    
*   As a **freelancer**, I want **availability checks against my calendar** so that I commit to realistic start dates.
    
*   As a **system**, I want **optional external calendar sync** so that reminders surface where users work.
    

### Flow

1.  SetProposalDeadlineCommand(proposal\_id, deadline\_at) → ValidateFuture() → Save() → **Outbox:** proposal.deadline.set.v1
    
2.  SetReminderCommand(proposal\_id, when) → GuardWithinWindow() | Schedule() → **Outbox:** proposal.reminder.scheduled.v1
    
3.  Worker → SendReminder() → **Outbox:** proposal.reminder.sent.v1
    
4.  SyncCalendarIntegrationCommand(user\_id, provider) → Connect() → **Outbox:** proposal.calendar.synced.v1
    

### Projections

*   proposal\_calendar\_read (deadlines, reminders)
    
*   proposal\_calendar\_integration\_read
    

### Events

proposal.deadline.set.v1, proposal.reminder.scheduled.v1, proposal.reminder.sent.v1, proposal.calendar.synced.v1

### RBAC/SLO

*   **RBAC:** **OWNER**;
    
*   **SLO:** set **P95 < 160 ms**; sync **async**.
    
*   **Idempotency:** (proposal\_id, deadline\_at) and (proposal\_id, reminder\_at).
    

9 - PROPOSAL RECOMMENDATION & VISIBILITY
=======================================

9.1 proposal\_recommendation/
-----------------------------

### Stories

*   As a **freelancer**, I want **recommended jobs** that my proposal would rank well on so that I invest my time wisely.
    
*   As a **system**, I want **click/convert feedback loops** so that models improve continuously.
    
*   As a **freelancer**, I want **visible “why” factors** so that I trust the recommendations.
    

### Flow

1.  GenerateRecommendationsCommand(user\_id, limit, filters) → FetchProfileSignals() | ComputeMatchScore() | Rank() → **Outbox:** proposal.recommended.v1
    
2.  TrackRecommendationClickCommand(reco\_id, job\_id) → RecordClick() → **Outbox:** proposal.recommendation.clicked.v1
    
3.  TrackRecommendationConversionCommand(reco\_id, proposal\_id) → RecordConversion() → **Outbox:** proposal.recommendation.converted.v1
    

### Projections

*   proposal\_recommendation\_read (list, reasons)
    
*   proposal\_recommendation\_kpi\_read (CTR, CVR)
    

### Events

proposal.recommended.v1, proposal.recommendation.clicked.v1, proposal.recommendation.converted.v1

### RBAC/SLO

*   **RBAC:** **OWNER**;
    
*   **SLO:** generate **P95 < 400 ms** (cached); idempotent by (user\_id, filter\_hash, day).
    

9.2 visibility\_budget/
-----------------------

### Stories

*   As a **freelancer**, I want to **set a CPM budget for impressions** so that I can control daily spend.
    
*   As a **system**, I want to **track spend and throttle when budgets exhaust** so that billing remains accurate.
    
*   As a **freelancer**, I want to **pause/resume a campaign** so that I can react to workload changes.
    

### Flow

1.  SetVisibilityBudgetCommand(proposal\_id, daily\_budget, cpm, start\_at, end\_at?) → ValidateLimits() | CreateCampaign() → **Outbox:** proposal.visibility.budget.set.v1
    
2.  RecordImpressionCommand(proposal\_id, viewer\_session) → DedupImpressions() | Bill(cpm) | ThrottleIfExhausted() → **Outbox:** proposal.visibility.impression.recorded.v1
    
3.  PauseVisibilityCampaignCommand(campaign\_id) → Pause() → **Outbox:** proposal.visibility.campaign.paused.v1
    

### Projections

*   visibility\_budget\_read (spend, remaining)
    
*   visibility\_spend\_read
    

### Events

proposal.visibility.budget.set.v1, proposal.visibility.impression.recorded.v1, proposal.visibility.campaign.paused.v1, proposal.visibility.budget.exhausted.v1

### RBAC/SLO

*   **RBAC:** **OWNER**;
    
*   **SLO:** record **P95 < 50 ms**; idempotent by (proposal\_id, session\_id, view\_seq).
    

9.3 proposal\_metrics/
----------------------

### Stories

*   As a **freelancer**, I want **view/response/hire rates and peer benchmarks** so that I can see where I stand.
    
*   As a **system**, I want **periodic recompute and caching** so that dashboards stay fresh without heavy load.
    

### Flow

1.  CalculateMetricsCommand(proposal\_id) → Aggregate() → **Outbox:** proposal.metrics.calculated.v1
    
2.  UpdateBenchmarkCommand(category, region) → RecomputePeerStats() → **Outbox:** proposal.benchmark.updated.v1
    

### Projections

*   proposal\_metrics\_read
    
*   proposal\_benchmark\_read
    

### Events

proposal.metrics.calculated.v1, proposal.benchmark.updated.v1

### RBAC/SLO

*   **RBAC:** **OWNER/SYSTEM**;
    
*   **SLO:** compute **P95 < 280 ms**.
    

9.4 proposal\_health/
---------------------

### Stories

*   As a **freelancer**, I want a **health score** (completeness, competitiveness, quality) with **actionable tips** so that I know what to fix first.
    
*   As a **system**, I want **tip application tracked** so that we can measure uplift.
    

### Flow

1.  CalculateHealthCommand(proposal\_id) → Score() → **Outbox:** proposal.health.scored.v1
    
2.  ApplyHealthTipCommand(proposal\_id, tip\_id) → ApplyChange() → **Outbox:** proposal.health.tip.applied.v1
    

### Projections

*   proposal\_health\_read (score, tips)
    

### Events

proposal.health.scored.v1, proposal.health.tip.applied.v1

### RBAC/SLO

*   **RBAC:** **OWNER**;
    
*   **SLO:** **P95 < 220 ms**.
    

9.5 conversion\_tracking/
-------------------------

### Stories

*   As a **system**, I want **funnel tracking and attribution** so that we know what drives outcomes.
    
*   As a **product analyst**, I want **drop-off analysis** so that I can recommend fixes.
    

### Flow

1.  TrackStageEntryCommand(proposal\_id, stage) → Record() → **Outbox:** proposal.stage.entered.v1
    
2.  RecordConversionCommand(proposal\_id, conversion\_type) → Upsert() → **Outbox:** proposal.converted.v1
    
3.  AnalyzeDropOffJob(job\_id) → ComputeDropOff() → **Outbox:** proposal.funnel.dropped.v1
    

### Projections

*   proposal\_conversion\_funnel\_read
    
*   proposal\_attribution\_read
    

### Events

proposal.stage.entered.v1, proposal.converted.v1, proposal.funnel.dropped.v1, proposal.attribution.tracked.v1

### RBAC/SLO

*   **RBAC:** **SYSTEM**;
    
*   **SLO:** track **P95 < 100 ms**.
    

9.6 proposal\_similarity/
-------------------------

### Stories

*   As a **system**, I want **fingerprinting and clustering** so that duplicates are minimized and differentiation is encouraged.
    
*   As a **freelancer**, I want a **differentiation score** so that I can stand out from peers.
    

### Flow

1.  CreateFingerprintCommand(proposal\_id) → HashTokens()/Embedding() | Persist() → **Outbox:** proposal.fingerprint.created.v1
    
2.  ClusterProposalsJob() → FormClusters() → **Outbox:** proposal.cluster.formed.v1
    
3.  ScoreDifferentiationCommand(proposal\_id) → ComputeScore() → **Outbox:** proposal.differentiation.scored.v1
    

### Projections

*   proposal\_similarity\_read
    
*   proposal\_clusters\_read
    

### Events

proposal.fingerprint.created.v1, proposal.cluster.formed.v1, proposal.differentiation.scored.v1

### RBAC/SLO

*   **RBAC:** **SYSTEM**;
    
*   **SLO:** batch 20k **< 60 s**.
    

9.7 proposal\_context/
----------------------

### Stories

*   As a **freelancer**, I want a **job/client/market context pack** so that I can tailor my proposal effectively.
    
*   As a **system**, I want **periodic refreshes on job-signal changes** so that context stays current.
    

### Flow

1.  GenerateContextCommand(proposal\_id) → AnalyzeJob() | AnalyzeClient() | AnalyzeMarket() | Store() → **Outbox:** proposal.context.generated.v1
    
2.  RefreshContextJob() → Update() → **Outbox:** proposal.context.updated.v1
    

### Projections

*   proposal\_context\_read (highlights, risks, suggested angles)
    

### Events

proposal.context.generated.v1, proposal.context.updated.v1

### RBAC/SLO

*   **RBAC:** **OWNER**;
    
*   **SLO:** gen **P95 < 500 ms**.
    

9.8 proposal\_urgency/
----------------------

### Stories

*   As a **freelancer**, I want an **urgency score and recommended action** so that I act at the right time.
    
*   As a **system**, I want **surge notifications** when competition spikes so that freelancers can respond quickly.
    

### Flow

1.  CalculateUrgencyCommand(proposal\_id) → Score() → **Outbox:** proposal.urgency.set.v1
    
2.  TriggerUrgencyAlertJob(job\_id) → DetectSurge() → **Outbox:** proposal.urgency.alert.triggered.v1
    

### Projections

*   proposal\_urgency\_read
    

### Events

proposal.urgency.set.v1, proposal.urgency.alert.triggered.v1

### RBAC/SLO

*   **RBAC:** **OWNER/SYSTEM**;
    
*   **SLO:** **P95 < 150 ms**.
    

10 - PROPOSAL COMPLIANCE & RISK
==============================

10.1 proposal\_risk\_assessment/
--------------------------------

### Stories

*   As a **trust analyst**, I want **risk scores** (unrealistic timeline, too-cheap bids, spam language) so that I can intervene early.
    
*   As a **system**, I want **mitigation flows** (warnings, gating, require-verify) so that platform risk is reduced.
    
*   As a **freelancer**, I want **remediation guidance** so that I can fix flagged issues.
    

### Flow

1.  AssessRiskCommand(proposal\_id) → ExtractFeatures() | ScoreModel() | Persist() → **Outbox:** proposal.risk.detected.v1
    
2.  MitigateRiskCommand(proposal\_id, action) → ApplyMitigation() → **Outbox:** proposal.risk.mitigated.v1
    

### Projections

*   proposal\_risk\_read (score, reasons, actions)
    

### Events

proposal.risk.detected.v1, proposal.risk.mitigated.v1, proposal.red\_flagged.v1

### RBAC/SLO

*   **RBAC:** **TRUST/ADMIN**;
    
*   **SLO:** assess **P95 < 300 ms**.
    

10.2 proposal\_escrow\_requirement/
-----------------------------------

### Stories

*   As a **client**, I want to **set escrow requirements per milestone** so that payment assurance exists.
    
*   As a **freelancer**, I want to **view funded status** so that I can decide next steps confidently.
    

### Flow

1.  SetEscrowRequirementCommand(proposal\_id, rules) → Validate() | Persist() → **Outbox:** proposal.escrow.set.v1
    
2.  UpdateEscrowFundingEvent(milestone\_id, status) → SyncStatus() → **Outbox:** proposal.escrow.funded.v1 | proposal.escrow.released.v1
    

### Projections

*   proposal\_escrow\_read (requirements + status)
    

### Events

proposal.escrow.set.v1, proposal.escrow.funded.v1, proposal.escrow.released.v1

### RBAC/SLO

*   **RBAC:** **OWNER/CLIENT**;
    
*   **SLO:** **P95 < 220 ms**.
    

10.3 proposal\_insurance/
-------------------------

### Stories

*   As a **freelancer**, I want **delivery/liability coverage options** so that I de-risk the engagement.
    
*   As a **system**, I want **premium calculation and claims workflow** so that coverage is end-to-end usable.
    

### Flow

1.  PurchaseInsuranceCommand(proposal\_id, coverage, term) → CalculatePremium() | Charge() | BindPolicy() → **Outbox:** proposal.insurance.purchased.v1
    
2.  FileInsuranceClaimCommand(policy\_id, reason) → OpenClaim() → **Outbox:** proposal.claim.filed.v1
    
3.  ProcessClaimCommand(claim\_id, decision) → Apply() → **Outbox:** proposal.claim.paid.v1 | proposal.claim.denied.v1
    

### Projections

*   proposal\_insurance\_read (policy)
    
*   proposal\_claims\_read
    

### Events

proposal.insurance.purchased.v1, proposal.claim.filed.v1, proposal.claim.paid.v1, proposal.claim.denied.v1

### RBAC/SLO

*   **RBAC:** **OWNER**;
    
*   **SLO:** **P95 < 240 ms**.
    

11 - PROPOSAL MONETIZATION
=========================

11.1 proposal\_subscription/
----------------------------

### Stories

*   As a **freelancer**, I want **plan tiers** (Free/Basic/Plus/Enterprise) with **monthly connects and rollover** so that I can scale activity.
    
*   As a **system**, I want **proration and billing alignment** so that charges remain fair.
    

### Flow

1.  UpgradeSubscriptionCommand(user\_id, plan) → ChargeOrProrate() | AllocateConnects() → **Outbox:** subscription.upgraded.v1, connects.allocated.v1
    
2.  RolloverConnectsJob() → ApplyRolloverCaps() → **Outbox:** connects.rolled\_over.v1
    

### Projections

*   proposal\_subscription\_read
    
*   connect\_allocation\_read
    

### Events

subscription.upgraded.v1, connects.allocated.v1, connects.rolled\_over.v1

### RBAC/SLO

*   **RBAC:** **OWNER**;
    
*   **SLO:** **P95 < 200 ms**.
    

11.2 proposal\_marketplace\_fee/
--------------------------------

### Stories

*   As a **system**, I want **dynamic/tiered marketplace fees and volume discounts** so that economics are predictable.
    
*   As a **freelancer**, I want **fee transparency** so that I can price bids appropriately.
    

### Flow

1.  CalculateFeeCommand(proposal\_id) → LoadRules() | Compute() → **Outbox:** proposal.fee.calculated.v1
    
2.  ApplyDiscountCommand(proposal\_id, reason) → Apply() → **Outbox:** proposal.fee.discount.applied.v1
    
3.  ChargeFeeCommand(proposal\_id) → Charge() → **Outbox:** proposal.fee.charged.v1
    

### Projections

*   proposal\_fee\_read
    

### Events

proposal.fee.calculated.v1, proposal.fee.discount.applied.v1, proposal.fee.charged.v1

### RBAC/SLO

*   **RBAC:** **SYSTEM/FINANCE**;
    
*   **SLO:** **P95 < 180 ms**.
    

11.3 proposal\_premium\_features/
---------------------------------

### Stories

*   As a **freelancer**, I want **add-ons** (featured, urgent review, guaranteed response) so that I can boost visibility when needed.
    
*   As a **system**, I want **ROI tracking** so that value of add-ons is measurable.
    

### Flow

1.  PurchasePremiumFeatureCommand(proposal\_id, feature, term) → Charge() | Activate() → **Outbox:** premium.feature.purchased.v1
    
2.  ExpirePremiumFeaturesJob() → Expire() → **Outbox:** premium.feature.expired.v1
    
3.  CalculatePremiumFeatureROICommand(proposal\_id) → ComputeLift() → **Outbox:** premium.feature.roi.calculated.v1
    

### Projections

*   premium\_features\_read
    
*   premium\_feature\_roi\_read
    

### Events

premium.feature.purchased.v1, premium.feature.expired.v1, premium.feature.roi.calculated.v1, premium.feature.activated.v1

### RBAC/SLO

*   **RBAC:** **OWNER**;
    
*   **SLO:** **P95 < 200 ms**.
    

11.4 proposal\_guarantee/
-------------------------

### Stories

*   As a **platform**, I want to **offer guarantees** (money-back/quality/delivery) so that client trust increases.
    
*   As a **freelancer**, I want **transparent claim handling** so that expectations are clear.
    

### Flow

1.  OfferGuaranteeCommand(proposal\_id, type, terms) → Persist() → **Outbox:** proposal.guarantee.offered.v1
    
2.  FileGuaranteeClaimCommand(guarantee\_id, reason) → OpenCase() → **Outbox:** proposal.guarantee.claim.filed.v1
    
3.  HonorGuaranteeCommand(case\_id) → IssueRefund() → **Outbox:** proposal.guarantee.honored.v1
    

### Projections

*   proposal\_guarantee\_read
    

### Events

proposal.guarantee.offered.v1, proposal.guarantee.claim.filed.v1, proposal.guarantee.honored.v1

### RBAC/SLO

*   **RBAC:** **OWNER/ADMIN**;
    
*   **SLO:** **P95 < 220 ms**.
    

11.5 proposal\_marketplace/
---------------------------

### Stories

*   As a **freelancer**, I want to **list pre-packaged proposal services** so that clients can buy quickly.
    
*   As a **freelancer**, I want **featuring options** so that my listings get more visibility.
    
*   As a **system**, I want **sale recording** so that fulfillment and analytics are consistent.
    

### Flow

1.  ListInMarketplaceCommand(proposal\_package) → Validate() | Publish() → **Outbox:** proposal.listed.v1
    
2.  FeatureListingCommand(listing\_id, term) → Charge() | Feature() → **Outbox:** proposal.featured.v1
    
3.  RecordPackageSaleEvent(listing\_id, order\_id) → Record() → **Outbox:** proposal.package.sold.v1
    

### Projections

*   proposal\_marketplace\_read
    

### Events

proposal.listed.v1, proposal.package.sold.v1, proposal.featured.v1, proposal.featured.expired.v1, proposal.listing.optimized.v1

### RBAC/SLO

*   **RBAC:** **OWNER**;
    
*   **SLO:** **P95 < 200 ms**.
    

12 - PROPOSAL COMMUNITY & SOCIAL
===============================

12.1 proposal\_mentorship/
--------------------------

### Stories

*   As a **freelancer**, I want **mentor reviews with actionable feedback** so that my proposal quality improves.
    
*   As a **system**, I want **fair matching and tracked completion** so that the program scales.
    

### Flow

1.  RequestMentorReviewCommand(proposal\_id) → MatchMentor() | OpenReview() → **Outbox:** proposal.mentorship.requested.v1
    
2.  SubmitMentorFeedbackCommand(review\_id, feedback) → Save() → **Outbox:** proposal.mentorship.feedback.received.v1
    
3.  CloseMentorshipCommand(review\_id) → Complete() → **Outbox:** proposal.mentorship.completed.v1
    

### Projections

*   proposal\_mentorship\_read
    

### Events

proposal.mentorship.requested.v1, proposal.mentorship.feedback.received.v1, proposal.mentorship.completed.v1

### RBAC/SLO

*   **RBAC:** **OWNER**;
    
*   **SLO:** **P95 < 240 ms**.
    

12.2 proposal\_showcase/
------------------------

### Stories

*   As a **freelancer**, I want to **showcase anonymized proposals** so that I build reputation without exposing PII.
    
*   As a **community member**, I want to **vote and learn** so that best practices surface.
    

### Flow

1.  ShowcaseProposalCommand(proposal\_id) → Anonymize() | Publish() → **Outbox:** proposal.showcased.v1
    
2.  VoteOnShowcaseCommand(showcase\_id, vote) → RateLimit() | Record() → **Outbox:** proposal.liked.v1
    
3.  FeatureShowcaseCommand(showcase\_id) → Feature() → **Outbox:** proposal.featured.v1
    

### Projections

*   proposal\_showcase\_read, proposal\_learning\_library\_read
    

### Events

proposal.showcased.v1, proposal.liked.v1, proposal.featured.v1, proposal.showcase.view.recorded.v1

### RBAC/SLO

*   **RBAC:** **OWNER/COMMUNITY**;
    
*   **SLO:** **P95 < 200 ms**.
    

12.3 proposal\_collaboration\_network/
--------------------------------------

### Stories

*   As a **freelancer**, I want to **form teams, assign roles, and set revenue splits** so that we can pursue larger proposals together.
    
*   As a **lead**, I want **add/remove permissions** so that collaboration remains safe.
    

### Flow

1.  FormTeamCommand(proposal\_id, members\[\]) → ValidateRoles() | Persist() → **Outbox:** team.formed.v1
    
2.  AddCollaboratorCommand(proposal\_id, member) → AuthorizeLead() | Add() → **Outbox:** team.member.added.v1
    
3.  RemoveCollaboratorCommand(proposal\_id, member\_id) → Remove() → **Outbox:** team.member.removed.v1
    
4.  SetRevenueSplitCommand(proposal\_id, splits\[\]) → Validate100Percent() | Save() → **Outbox:** revenue.split.set.v1
    

### Projections

*   proposal\_team\_read, proposal\_revenue\_share\_read
    

### Events

team.formed.v1, team.member.added.v1, team.member.removed.v1, revenue.split.set.v1, revenue.distributed.v1

### RBAC/SLO

*   **RBAC:** **OWNER/TEAM LEAD**;
    
*   **SLO:** **P95 < 220 ms**.
    

12.4 proposal\_social\_proof/
-----------------------------

### Stories

*   As a **freelancer**, I want **badges, endorsements, and success stories** so that clients trust my proposal more.
    
*   As a **system**, I want **verification of endorsers** so that social proof remains credible.
    

### Flow

1.  EarnBadgeCommand(user\_id, badge\_type) → Validate() | Award() → **Outbox:** proposal.badge.earned.v1
    
2.  AddEndorsementCommand(proposal\_id, endorsement) → VerifyEndorser?() | Persist() → **Outbox:** proposal.endorsed.v1
    
3.  AttachSuccessStoryCommand(proposal\_id, story\_id) → Link() → **Outbox:** proposal.success\_story.added.v1
    

### Projections

*   proposal\_social\_proof\_read
    

### Events

proposal.badge.earned.v1, proposal.endorsed.v1, proposal.success\_story.added.v1, proposal.social\_proof.updated.v1

### RBAC/SLO

*   **RBAC:** **OWNER/SYSTEM**;
    
*   **SLO:** **P95 < 180 ms**.
    

12.5 proposal\_experiment/
--------------------------

### Stories

*   As a **freelancer**, I want **A/B tests** (cover letter, bid, boost) with **significance** so that I can choose winning variants.
    
*   As a **system**, I want **clean experiment lifecycles** so that results are trustworthy.
    

### Flow

1.  StartExperimentCommand(proposal\_id, variants\[\]) → ValidateVariantCount() | Randomize() → **Outbox:** abtest.started.v1
    
2.  RecordVariantOutcomeCommand(experiment\_id, variant\_id, metric) → Record() → **Outbox:** abtest.variant.tested.v1
    
3.  DeclareWinnerCommand(experiment\_id) → ComputeSig() | SetWinner() → **Outbox:** abtest.winner.declared.v1
    
4.  EndExperimentCommand(experiment\_id) → Close() → **Outbox:** abtest.ended.v1
    

### Projections

*   proposal\_experiment\_read
    

### Events

abtest.started.v1, abtest.variant.tested.v1, abtest.winner.declared.v1, abtest.ended.v1

### RBAC/SLO

*   **RBAC:** **OWNER**;
    
*   **SLO:** **P95 < 250 ms**.
    

12.6 proposal\_coaching/
------------------------

### Stories

*   As a **freelancer**, I want **AI coaching with concrete examples** so that I can improve quickly.
    
*   As a **system**, I want **tracked improvements** so that coaching value is measurable.
    

### Flow

1.  RequestCoachingCommand(proposal\_id, goal) → GenerateAdvice() → **Outbox:** proposal.coaching.requested.v1
    
2.  ApplyCoachingSuggestionCommand(proposal\_id, suggestion\_id) → Apply() → **Outbox:** proposal.coaching.suggestion.applied.v1
    
3.  CompleteCoachingCommand(proposal\_id) → Summarize() → **Outbox:** proposal.coaching.completed.v1
    

### Projections

*   proposal\_coaching\_read
    

### Events

proposal.coaching.requested.v1, proposal.coaching.suggestion.applied.v1, proposal.coaching.completed.v1

### RBAC/SLO

*   **RBAC:** **OWNER**;
    
*   **SLO:** **P95 < 240 ms**.
    

12.7 proposal\_gamification/
----------------------------

### Stories

*   As a **freelancer**, I want **points, levels, achievements, streaks, and leaderboards** so that I stay motivated.
    
*   As a **system**, I want **fair anti-abuse rules** so that competition remains healthy.
    

### Flow

1.  EarnPointsCommand(user\_id, action, points) → Record() → **Outbox:** points.earned.v1
    
2.  LevelUpJob() → CheckThresholds() → **Outbox:** level.up.v1
    
3.  UnlockAchievementCommand(user\_id, badge) → Unlock() → **Outbox:** achievement.unlocked.v1
    
4.  TrackStreakJob() → Compute() → **Outbox:** streak.updated.v1 | streak.broken.v1
    
5.  UpdateLeaderboardJob() → Rank() → **Outbox:** leaderboard.updated.v1
    

### Projections

*   gamification\_read, leaderboard\_read
    

### Events

points.earned.v1, level.up.v1, achievement.unlocked.v1, streak.updated.v1, streak.broken.v1, leaderboard.updated.v1

### RBAC/SLO

*   **RBAC:** **OWNER**;
    
*   **SLO:** **P95 < 200 ms**.
    

12.8 proposal\_learning/
------------------------

### Stories

*   As a **system**, I want to **learn from outcomes and feedback** so that future proposals improve.
    
*   As a **freelancer**, I want **suggested improvements from learned patterns** so that I can iterate faster.
    

### Flow

1.  DetectPatternJob() → MineSuccessfulSignals() → **Outbox:** proposal.pattern.detected.v1
    
2.  ProcessFeedbackCommand(proposal\_id, feedback) → Normalize() | Persist() → **Outbox:** proposal.lesson.learned.v1
    
3.  SuggestImprovementCommand(proposal\_id) → Recommend() → **Outbox:** proposal.improvement.suggested.v1
    

### Projections

*   proposal\_learning\_read
    

### Events

proposal.pattern.detected.v1, proposal.lesson.learned.v1, proposal.improvement.suggested.v1

### RBAC/SLO

*   **RBAC:** **OWNER/SYSTEM**;
    
*   **SLO:** **P95 < 260 ms**.
    

12.9 proposal\_time\_integration/
---------------------------------

### Stories

*   As a **freelancer**, I want **time estimates and actual tracked hours** so that I can report accuracy and improve planning.
    
*   As a **system**, I want **estimate-vs-actual accuracy** so that we can calibrate future estimates.
    

### Flow

1.  EstimateTimeCommand(proposal\_id, estimate\_hours) → Save() → **Outbox:** proposal.time.estimated.v1
    
2.  TrackTimeCommand(proposal\_id, hours) → Record() → **Outbox:** proposal.time.tracked.v1
    
3.  UpdateEstimateAccuracyJob() → Compare() → **Outbox:** proposal.estimate.accuracy.updated.v1
    

### Projections

*   proposal\_time\_read
    

### Events

proposal.time.estimated.v1, proposal.time.tracked.v1, proposal.estimate.accuracy.updated.v1

### RBAC/SLO

*   **RBAC:** **OWNER**;
    
*   **SLO:** **P95 < 180 ms**.
    

13 - COLLABORATION & WORKFLOW
============================

13.1 negotiation/
-----------------

### Stories

*   As a **client/freelancer**, I want **counter-offers** (rate/scope/milestones) with **expiry** so that we converge quickly.
    
*   As a **system**, I want **immutable threads and states** so that outcomes are auditable.
    

### Flow

1.  OpenNegotiationCommand(proposal\_id) → CreateThread() → **Outbox:** negotiation.opened.v1
    
2.  ProposeCounterCommand(thread\_id, terms) → Validate() | Append() → **Outbox:** negotiation.countered.v1
    
3.  AcceptCounterCommand(counter\_id) → LockTerms() → **Outbox:** negotiation.accepted.v1
    
4.  DeclineCounterCommand(counter\_id, reason) → Record() → **Outbox:** negotiation.declined.v1
    
5.  ExpireNegotiationJob() → Expire() → **Outbox:** negotiation.expired.v1
    

### Projections

*   proposal\_negotiation\_read
    

### Events

negotiation.opened.v1, negotiation.countered.v1, negotiation.accepted.v1, negotiation.declined.v1, negotiation.expired.v1

### RBAC/SLO

*   **RBAC:** **OWNER/CLIENT**;
    
*   **SLO:** **P95 < 180 ms**.
    

13.2 invite/
------------

### Stories

*   As a **client**, I want to **invite a freelancer with a prefilled draft** so that they can accept quickly.
    
*   As a **freelancer**, I want to **accept/decline with reason** so that expectations are clear.
    

### Flow

1.  CreateInviteCommand(job\_id, freelancer\_id, message?) → PreventDuplicateOpenInvite() → **Outbox:** invite.sent.v1
    
2.  AcceptInviteCommand(invite\_id) → CreateDraftProposalFromInvite() → **Outbox:** invite.accepted.v1, proposal.core.created.v1
    
3.  DeclineInviteCommand(invite\_id, reason) → Record() → **Outbox:** invite.declined.v1
    
4.  ExpireInviteJob() → Expire() → **Outbox:** invite.expired.v1
    

### Projections

*   proposal\_invite\_read
    

### Events

invite.sent.v1, invite.accepted.v1, invite.declined.v1, invite.expired.v1

### RBAC/SLO

*   **RBAC:** **CLIENT→FREELANCER**;
    
*   **SLO:** **P95 < 160 ms**.
    

13.3 revision/
--------------

### Stories

*   As a **freelancer**, I want **post-submit revisions with client approval** so that I can improve without breaking trust.
    
*   As a **client**, I want **approve/reject with reasons** so that changes remain controlled.
    

### Flow

1.  CreateRevisionCommand(proposal\_id, changes) → AppendRevision() → **Outbox:** proposal.revised.v1
    
2.  ApproveRevisionCommand(revision\_id) → Merge() → **Outbox:** proposal.revision.approved.v1
    
3.  RejectRevisionCommand(revision\_id, reason) → Record() → **Outbox:** proposal.revision.rejected.v1
    

### Projections

*   proposal\_revision\_read
    

### Events

proposal.revised.v1, proposal.revision.approved.v1, proposal.revision.rejected.v1

### RBAC/SLO

*   **RBAC:** **OWNER/CLIENT**;
    
*   **SLO:** **P95 < 200 ms**.
    

13.4 collaboration/
-------------------

### Stories

*   As a **lead freelancer**, I want to **add/remove collaborators with permissions** so that we can co-author safely.
    
*   As a **collaborator**, I want **clear roles** so that I know what I can edit.
    

### Flow

1.  AddCollaboratorCommand(proposal\_id, member\_id, role) → ValidateRole() | Add() → **Outbox:** proposal.collaborator.added.v1
    
2.  RemoveCollaboratorCommand(proposal\_id, member\_id) → Remove() → **Outbox:** proposal.collaborator.removed.v1
    

### Projections

*   proposal\_collaboration\_read
    

### Events

proposal.collaborator.added.v1, proposal.collaborator.removed.v1

### RBAC/SLO

*   **RBAC:** **OWNER**;
    
*   **SLO:** **P95 < 160 ms**.
    

13.5 approval\_workflow/
------------------------

### Stories

*   As a **team/enterprise**, I want **multi-step approvals with SLA warnings** so that submissions meet internal policy before sending.
    
*   As an **approver**, I want **role-based routing** so that requests reach the right person.
    

### Flow

1.  RequestApprovalCommand(proposal\_id, approver\_chain\[\]) → CreateFlow() → **Outbox:** approval.requested.v1
    
2.  GrantApprovalCommand(flow\_id, step\_id) → Advance() → **Outbox:** approval.granted.v1
    
3.  RejectApprovalCommand(flow\_id, step\_id, reason) → Stop() → **Outbox:** approval.rejected.v1
    
4.  SlaCheckJob() → WarnOverdue() → **Outbox:** approval.sla.warning.v1
    

### Projections

*   approval\_flow\_read
    

### Events

approval.requested.v1, approval.granted.v1, approval.rejected.v1, approval.sla.warning.v1

### RBAC/SLO

*   **RBAC:** **OWNER/APPROVER**;
    
*   **SLO:** **P95 < 220 ms**.
    

13.6 document\_redlining/
-------------------------

### Stories

*   As **client/freelancer**, I want **redline threads on shared docs with diffs** so that we can resolve comments precisely.
    
*   As a **system**, I want **resolvable threads** so that we keep a clean audit trail.
    

### Flow

1.  StartRedlineThreadCommand(proposal\_id, doc\_ref, start\_range) → OpenThread() → **Outbox:** redline.thread.started.v1
    
2.  AppendRedlineDiffCommand(thread\_id, diff) → Append() → **Outbox:** redline.diff.appended.v1
    
3.  ResolveRedlineThreadCommand(thread\_id) → Resolve() → **Outbox:** redline.thread.resolved.v1
    

### Projections

*   redline\_thread\_read
    

### Events

redline.thread.started.v1, redline.diff.appended.v1, redline.thread.resolved.v1

### RBAC/SLO

*   **RBAC:** **OWNER/CLIENT**;
    
*   **SLO:** **P95 < 200 ms**.
    

14 - CLIENT INTERACTION
======================

14.1 interview\_request/
------------------------

### Stories

*   As a **client**, I want to **request/schedule/complete interviews** so that we can evaluate candidates efficiently.
    
*   As a **freelancer**, I want **status tracking** so that I know where I stand.
    

### Flow

1.  RequestInterviewCommand(proposal\_id, slots\[\]) → Persist() → **Outbox:** interview.requested.v1
    
2.  ScheduleInterviewCommand(request\_id, final\_slot) → Confirm() → **Outbox:** interview.scheduled.v1
    
3.  CompleteInterviewCommand(interview\_id, notes?) → MarkCompleted() → **Outbox:** interview.completed.v1
    
4.  DeclineInterviewCommand(request\_id, reason) → Record() → **Outbox:** interview.declined.v1
    

### Projections

*   interview\_request\_read
    

### Events

interview.requested.v1, interview.scheduled.v1, interview.completed.v1, interview.declined.v1

### RBAC/SLO

*   **RBAC:** **CLIENT/OWNER**;
    
*   **SLO:** **P95 < 180 ms**.
    

14.2 feedback/
--------------

### Stories

*   As a **client**, I want to **give pre-contract feedback** so that freelancers can improve proposals.
    
*   As a **freelancer**, I want to **respond** so that I can clarify and iterate.
    

### Flow

1.  GiveProposalFeedbackCommand(proposal\_id, rating, comments) → AntiPiiLint() | Save() → **Outbox:** proposal.feedback.given.v1
    
2.  RespondToFeedbackCommand(feedback\_id, response) → AntiPiiLint() | Save() → **Outbox:** proposal.feedback.responded.v1
    

### Projections

*   proposal\_feedback\_read
    

### Events

proposal.feedback.given.v1, proposal.feedback.responded.v1

### RBAC/SLO

*   **RBAC:** **CLIENT/OWNER**;
    
*   **SLO:** **P95 < 150 ms**.
    

14.3 shortlist/
---------------

### Stories

*   As a **client**, I want to **add/update/remove shortlist status** so that I can manage candidates efficiently.
    

### Flow

1.  ShortlistProposalCommand(proposal\_id) → Add() → **Outbox:** proposal.shortlisted.v1
    
2.  UpdateShortlistStatusCommand(proposal\_id, status) → Validate() | Update() → **Outbox:** shortlist.updated.v1
    
3.  RemoveFromShortlistCommand(proposal\_id) → Remove() → **Outbox:** shortlist.removed.v1
    

### Projections

*   shortlist\_read
    

### Events

proposal.shortlisted.v1, shortlist.updated.v1, shortlist.removed.v1

### RBAC/SLO

*   **RBAC:** **CLIENT/TEAM**;
    
*   **SLO:** **P95 < 160 ms**.
    

14.4 client\_insight/
---------------------

### Stories

*   As a **freelancer**, I want **client hire rate/spend/feedback history** so that I can calibrate my effort.
    
*   As a **system**, I want **access logging** so that client data usage remains compliant.
    

### Flow

1.  GetClientInsightQuery(client\_id) → Assemble() | LogAccess() → **Outbox:** client.insight.accessed.v1
    
2.  RefreshClientInsightJob(client\_id) → Recompute() → **Outbox:** client.insight.updated.v1
    

### Projections

*   client\_insight\_read
    

### Events

client.insight.accessed.v1, client.insight.updated.v1

### RBAC/SLO

*   **RBAC:** **OWNER**;
    
*   **SLO:** **P95 < 220 ms**.
    

14.5 client\_preference/
------------------------

### Stories

*   As a **system**, I want **inferred client preferences** so that recommendations personalize.
    
*   As a **client**, I want to **update preferences** so that future proposals match my style.
    

### Flow

1.  InferClientPreferenceJob(client\_id) → AggregateSignals() | Persist() → **Outbox:** client.preference.inferred.v1
    
2.  UpdateClientPreferenceCommand(client\_id, patch) → Apply() → **Outbox:** client.preference.updated.v1
    

### Projections

*   client\_preference\_read
    

### Events

client.preference.inferred.v1, client.preference.updated.v1

### RBAC/SLO

*   **RBAC:** **SYSTEM/CLIENT**;
    
*   **SLO:** **P95 < 200 ms**.
    

15 - COMPLIANCE & LEGAL
======================

15.1 compliance/
----------------

### Stories

*   As a **system**, I want **plagiarism/spam/legal/ToS checks** so that platform integrity holds.
    
*   As a **trust analyst**, I want **remediation flows** so that issues are resolved quickly.
    

### Flow

1.  RecordComplianceCheckCommand(proposal\_id, check\_types\[\]) → RunChecks() | PersistFindings() → **Outbox:** compliance.checked.v1
    
2.  ResolveComplianceFindingCommand(finding\_id, action) → Resolve() → **Outbox:** compliance.resolved.v1
    

### Projections

*   compliance\_read, compliance\_audit\_read
    

### Events

compliance.checked.v1, compliance.passed.v1, compliance.failed.v1, compliance.resolved.v1

### RBAC/SLO

*   **RBAC:** **TRUST/ADMIN**;
    
*   **SLO:** **P95 < 300 ms**.
    

15.2 proposal\_flag/
--------------------

### Stories

*   As a **user**, I want to **flag proposals** so that moderators can review potential issues promptly.
    
*   As a **moderator**, I want **resolve/dismiss actions** so that cases close cleanly.
    

### Flow

1.  FlagProposalCommand(proposal\_id, reason) → OpenCase() → **Outbox:** proposal.flag.submitted.v1
    
2.  ResolveFlagCommand(flag\_id, action) → Resolve() → **Outbox:** proposal.flag.resolved.v1 | proposal.flag.dismissed.v1
    

### Projections

*   proposal\_flag\_read
    

### Events

proposal.flag.submitted.v1, proposal.flag.resolved.v1


16) CONTRACT & TERMS
=====================

16.1 terms/
-----------

### Stories

*   As a **freelancer**, I want a **clause library** with searchable templates and **jurisdiction tags** so that I can assemble terms quickly and correctly.
    
*   As a **client**, I want **inline redlining and comment threads tied to clauses** so that we resolve disagreements in context.
    
*   As a **system**, I want **automatic conflict detection** (e.g., payment terms vs. IP assignment) so that incompatible clauses are flagged before signature.
    
*   As an **approver**, I want **version diffs and compare-by-clause** so that I can review what changed efficiently.
    
*   As **legal**, I want **mandatory policy gates** (e.g., minimum liability caps) so that risky terms are blocked.
    
*   As **either party**, I want **e-sign handoff and acceptance order** (counter-sign required) so that execution is unambiguous.
    
*   As a **freelancer**, I want **expiry/renewal reminders** on term bundles so that I can renegotiate on time.
    
*   As a **system**, I want **PDF/DOCX export with watermark and cryptographic hash** so that executed terms are portable and verifiable.
    

### Flow

1.  AddTermsCommand(proposal\_id, terms\_doc\_ref) → Persist() → **Outbox:** terms.added.v1
    
2.  UpdateTermsCommand(terms\_id, patch) → Version() → **Outbox:** terms.updated.v1
    
3.  AcceptTermsCommand(terms\_id, actor) → RecordAcceptance() → **Outbox:** terms.accepted.v1
    
4.  RejectTermsCommand(terms\_id, actor, reason) → Record() → **Outbox:** terms.rejected.v1
    

### Projections

*   terms\_read
    

### Events

terms.added.v1, terms.updated.v1, terms.accepted.v1, terms.rejected.v1

### RBAC/SLO

*   **RBAC:** **OWNER/CLIENT**
    
*   **SLO:** **P95 < 180 ms**.
    

16.2 nda/
---------

### Stories

*   As a **client**, I want **NDA templates** (mutual/unilateral) with **regional variants** so that I can comply with local law.
    
*   As a **freelancer**, I want **gatekeeping** (hide files/links until NDA signed) so that sensitive info isn’t exposed prematurely.
    
*   As a **system**, I want **identity checks** (verified profile/KYC-lite) before signature so that signers are credible.
    
*   As a **freelancer**, I want **NDA renewal and auto-extend options** so that long evaluations remain covered.
    
*   As **legal**, I want **NDA exceptions** (portfolio carve-outs) so that I can safely reuse non-sensitive material later.
    
*   As a **client**, I want **reminders and escalation** if an NDA sits unsigned so that timelines aren’t blocked.
    

### Flow

1.  ProposeNDACommand(proposal\_id, template) → Create() → **Outbox:** nda.proposed.v1
    
2.  SignNDACommand(nda\_id, actor) → Sign() → **Outbox:** nda.signed.v1
    
3.  RejectNDACommand(nda\_id, reason) → Record() → **Outbox:** nda.rejected.v1
    

### Projections

*   nda\_read
    

### Events

nda.proposed.v1, nda.signed.v1, nda.rejected.v1

### RBAC/SLO

*   **RBAC:** **OWNER/CLIENT**
    
*   **SLO:** **P95 < 200 ms**.
    

16.3 ip\_rights/
----------------

### Stories

*   As a **client**, I want **work-for-hire vs. license-only toggles** so that ownership is crystal clear.
    
*   As a **freelancer**, I want **third-party/open-source disclosure** with **license compatibility checks** so that I avoid downstream disputes.
    
*   As **legal**, I want **moral-rights waiver and attribution options** so that jurisdictions with extra rights are covered.
    
*   As **either party**, I want **field-of-use/territory/term constraints** so that usage limits are explicit.
    
*   As a **system**, I want **conflict alerts** if IP terms contradict the NDA or contract so that issues are surfaced early.
    

### Flow

1.  SetIPRightsCommand(proposal\_id, terms) → Persist() → **Outbox:** ip.rights.set.v1
    
2.  AgreeIPRightsCommand(ip\_rights\_id, actor) → Record() → **Outbox:** ip.rights.agreed.v1
    

### Projections

*   ip\_rights\_read
    

### Events

ip.rights.set.v1, ip.rights.agreed.v1

### RBAC/SLO

*   **RBAC:** **OWNER/CLIENT**
    
*   **SLO:** **P95 < 180 ms**.
    

16.4 payment\_terms/
--------------------

### Stories

*   As a **client**, I want **milestone schedules** with **dependencies and approvals** so that release timing is controlled.
    
*   As a **freelancer**, I want **deposits/retainers** and **“start work when funded” gates** so that cash risk is reduced.
    
*   As **finance**, I want **late-fee and net-terms settings** so that overdue behavior is consistent.
    
*   As a **system**, I want **FX lock windows and auto-reprice warnings** so that cross-currency deals are predictable.
    
*   As **either party**, I want **tax handling** (VAT/GST) and **who-bears-it flags** so that invoices are correct.
    

### Flow

1.  SetPaymentTermsCommand(proposal\_id, schedule\[\]) → ValidateSum() → **Outbox:** payment.terms.set.v1
    
2.  UpdatePaymentTermsCommand(terms\_id, patch) → Validate() → **Outbox:** payment.terms.updated.v1
    

### Projections

*   payment\_terms\_read
    

### Events

payment.terms.set.v1, payment.terms.updated.v1

### RBAC/SLO

*   **RBAC:** **OWNER/CLIENT**
    
*   **SLO:** **P95 < 180 ms**.
    

16.5 pricing\_model/
--------------------

### Stories

*   As a **freelancer**, I want **hourly/T&M with caps and alerts** so that budgets don’t spiral.
    
*   As a **freelancer**, I want **tiered packages and add-ons** so that clients can choose value levels.
    
*   As a **client**, I want **performance/bonus clauses** with measurable KPIs so that incentives align.
    
*   As a **system**, I want **promotional discounts and expiries** so that limited offers are enforced.
    
*   As **finance**, I want **auto-invoice schedule hooks per model** so that billing matches pricing rules.
    

### Flow

1.  SetPricingModelCommand(proposal\_id, model) → Validate() → **Outbox:** pricing.model.updated.v1
    
2.  AddPricingTierCommand(model\_id, tier) → Validate() → **Outbox:** pricing.tier.added.v1
    
3.  AutoAdjustDynamicPricingJob() → Reprice() → **Outbox:** pricing.dynamic.adjusted.v1
    

### Projections

*   pricing\_model\_read
    

### Events

pricing.model.updated.v1, pricing.tier.added.v1, pricing.dynamic.adjusted.v1

### RBAC/SLO

*   **RBAC:** **OWNER**
    
*   **SLO:** **P95 < 170 ms**.
    

17 - ENTERPRISE & PROCUREMENT
============================

17.1 rfp/
---------

### Stories

*   As **enterprise**, I want **multi-user section assignments and progress tracking** so that teams can author in parallel.
    
*   As a **proposal lead**, I want a **requirements compliance matrix with gap flags** so that nothing is missed.
    
*   As a **system**, I want **Q&A tracking** (vendor questions → client answers) so that clarifications are centralized.
    
*   As a **content manager**, I want **auto-suggestions from a knowledge base** so that boilerplate is consistent.
    
*   As **procurement**, I want **submission packaging** (format, portal export) so that we meet issuer specs.
    

### Flow

1.  ImportRFPCommand(file\_ref) → Parse() | Sectionize() → **Outbox:** rfp.imported.v1
    
2.  MapRFPSectionsCommand(rfp\_id, mapping) → Persist() → **Outbox:** rfp.mapped.v1
    
3.  GenerateComplianceMatrixCommand(rfp\_id) → Compute() → **Outbox:** rfp.compliance.completed.v1
    

### Projections

*   rfp\_import\_read, rfp\_compliance\_read
    

### Events

rfp.imported.v1, rfp.mapped.v1, rfp.compliance.completed.v1

### RBAC/SLO

*   **RBAC:** **OWNER/ENTERPRISE**
    
*   **SLO:** import **P95 < 2 s**.
    

17.2 quote/
-----------

### Stories

*   As a **client**, I want **accept/decline/counter on quotes** so that we can iterate formally.
    
*   As **finance**, I want **tax rules per locale and per-item discounts** so that totals are accurate.
    
*   As a **freelancer**, I want **multi-currency presentation** with a **billing currency lock** so that I can sell globally.
    
*   As a **system**, I want **signed-PDF generation and audit trails** so that quotes become hard records.
    
*   As **sales**, I want **quote-expiry reminders and nudges** so that deals don’t stall.
    

### Flow

1.  IssueQuoteCommand(proposal\_id, lines\[\], taxes\[\], fx\_lock?) → ComputeTotals() → **Outbox:** quote.issued.v1
    
2.  RepriceQuoteCommand(quote\_id, new\_rates) → Recompute() → **Outbox:** quote.repriced.v1
    
3.  ExpireQuoteCommand(quote\_id) → Expire() → **Outbox:** quote.expired.v1
    

### Projections

*   quote\_read
    

### Events

quote.issued.v1, quote.repriced.v1, quote.expired.v1

### RBAC/SLO

*   **RBAC:** **OWNER/CLIENT**
    
*   **SLO:** **P95 < 220 ms**.
    

17.3 procurement/
-----------------

### Stories

*   As a **client**, I want **vendor onboarding** (W-9/IBAN/company info) so that payments can be issued.
    
*   As **compliance**, I want **SOC2/ISO doc storage with expiry reminders** so that attestations stay fresh.
    
*   As a **system**, I want **approval chains and budget owners** so that spend control is enforced.
    
*   As **finance**, I want **PO matching and receipt evidence** so that payments reconcile cleanly.
    
*   As a **client**, I want **spend-cap alerts and freeze rules** so that overruns are prevented.
    

### Flow

1.  AttachPORefCommand(proposal\_id, po\_ref) → Persist() → **Outbox:** procurement.po.attached.v1
    
2.  RunBudgetCheckCommand(proposal\_id) → CallERP() → **Outbox:** procurement.budget.check.requested.v1, procurement.budget.check.satisfied.v1 | procurement.budget.check.failed.v1
    

### Projections

*   procurement\_read
    

### Events

procurement.po.attached.v1, procurement.budget.check.requested.v1, procurement.budget.check.satisfied.v1, procurement.budget.check.failed.v1

### RBAC/SLO

*   **RBAC:** **CLIENT/ENTERPRISE**
    
*   **SLO:** **P95 < 240 ms**.
    

17.4 evaluation\_rubric/
------------------------

### Stories

*   As a **client**, I want **multiple reviewers with blind mode** so that scoring is unbiased.
    
*   As a **system**, I want **normalization and outlier detection** so that inconsistent scorers don’t skew results.
    
*   As a **reviewer**, I want **per-criterion comments and evidence links** so that decisions are defensible.
    
*   As **procurement**, I want **calibration examples and weights per criterion** so that teams score consistently.
    
*   As a **client**, I want **exportable scorecards** so that I can share results offline.
    

### Flow

1.  CreateRubricCommand(job\_id, criteria\[\]) → ValidateWeights100() → **Outbox:** rubric.created.v1
    
2.  ScoreProposalCommand(rubric\_id, proposal\_id, scores\[\]) → ValidateRanges() → **Outbox:** rubric.scored.v1
    

### Projections

*   rubric\_read, rubric\_score\_read
    

### Events

rubric.created.v1, rubric.scored.v1

### RBAC/SLO

*   **RBAC:** **CLIENT**
    
*   **SLO:** **P95 < 200 ms**.
    

18 - AI & OPTIMIZATION
=====================

18.1 ai\_assist/
----------------

### Stories

*   As a **freelancer**, I want **tone/voice presets** (formal, concise, friendly) so that drafts match the client culture.
    
*   As a **system**, I want **retrieval from my past wins/profile** so that generated content is accurate and on-brand.
    
*   As **trust**, I want **hallucination guardrails and source citations** so that risky claims are minimized.
    
*   As a **freelancer**, I want **accept/reject granular blocks with diffs** so that I stay in control.
    
*   As **privacy**, I want **opt-out from training logs** so that sensitive data isn’t retained.
    

### Flow

1.  GenerateProposalWithAICommand(job\_id, profile\_context) → GuardModelUse() | CallModel() | Redact() | SaveDraft() → **Outbox:** ai.proposal.created.v1
    
2.  SaveAIDraftCommand(proposal\_id, draft\_id) → Persist() → **Outbox:** ai.suggestion.generated.v1, ai.assist.used.v1
    

### Projections

*   ai\_assist\_read, ai\_usage\_log\_read
    

### Events

ai.suggestion.generated.v1, ai.proposal.created.v1, ai.assist.used.v1

### RBAC/SLO

*   **RBAC:** **OWNER**
    
*   **SLO:** gen **P95 < 2.5 s**.
    

18.2 keyword\_optimization/
---------------------------

### Stories

*   As a **freelancer**, I want **job-specific keywords and synonyms per locale** so that I mirror client language.
    
*   As a **system**, I want **competitor/trend signals** (market language) so that suggestions reflect demand.
    
*   As a **freelancer**, I want **instant preview diffs and density meters** so that I avoid keyword stuffing.
    
*   As **product**, I want **uplift measurement on views/replies** so that optimization impact is proven.
    

### Flow

1.  OptimizeKeywordsCommand(proposal\_id) → Scan() | Suggest() → **Outbox:** keywords.optimized.v1
    
2.  ApplyKeywordOptimizationCommand(proposal\_id, suggestion\_id) → Apply() → **Outbox:** keyword.optimization.applied.v1
    

### Projections

*   keyword\_opt\_read
    

### Events

keywords.optimized.v1, keyword.optimization.applied.v1

### RBAC/SLO

*   **RBAC:** **OWNER**
    
*   **SLO:** **P95 < 300 ms**.
    

18.3 success\_predictor/
------------------------

### Stories

*   As a **freelancer**, I want **feature explanations** (why this score) so that I know what to improve.
    
*   As a **user**, I want **“what-if” simulators** (change bid/skills) so that I can test scenarios.
    
*   As a **system**, I want **cold-start fallbacks and confidence bands** so that low-data cases behave sensibly.
    
*   As **risk**, I want **bias monitoring and fairness reports** so that models remain compliant.
    

### Flow

1.  PredictSuccessCommand(proposal\_id) → LoadModel() | Score() → **Outbox:** success.predicted.v1
    
2.  UpdatePredictionModelCommand(version, checksum) → Validate() | Activate() → **Outbox:** prediction.updated.v1
    

### Projections

*   success\_prediction\_read
    

### Events

success.predicted.v1, prediction.updated.v1

### RBAC/SLO

*   **RBAC:** **OWNER/SYSTEM**
    
*   **SLO:** **P95 < 450 ms**.
    

18.4 best\_practices/
---------------------

### Stories

*   As a **freelancer**, I want **industry-specific checklists** so that I don’t miss essentials.
    
*   As a **system**, I want **one-click apply with preview and rollback** so that changes are safe.
    
*   As **product**, I want **cohort A/B of tips** so that we learn which practices actually work.
    

### Flow

1.  FetchBestPracticesCommand(category) → Return() → **Outbox:** best.practice.viewed.v1
    
2.  ApplyBestPracticeCommand(proposal\_id, tip\_id) → Apply() → **Outbox:** best.practice.applied.v1
    

### Projections

*   best\_practices\_read
    

### Events

best.practice.viewed.v1, best.practice.applied.v1

### RBAC/SLO

*   **RBAC:** **OWNER**
    
*   **SLO:** **P95 < 150 ms**.
    

18.5 grammar\_check/
--------------------

### Stories

*   As a **freelancer**, I want **locale-aware spelling and style guides** so that language fits the client region.
    
*   As a **system**, I want **tone (formal/casual) and inclusive-language checks** so that writing is professional and respectful.
    
*   As a **freelancer**, I want **batch-apply with per-suggestion control** so that I can fix many issues quickly.
    

### Flow

1.  CheckGrammarCommand(proposal\_id) → Analyze() → **Outbox:** grammar.checked.v1
    
2.  ApplyGrammarSuggestionsCommand(proposal\_id, suggestion\_ids\[\]) → Apply() → **Outbox:** grammar.issues.resolved.v1
    

### Projections

*   grammar\_check\_read
    

### Events

grammar.checked.v1, grammar.issues.resolved.v1

### RBAC/SLO

*   **RBAC:** **OWNER**
    
*   **SLO:** **P95 < 300 ms**.
    

18.6 personalization/
---------------------

### Stories

*   As a **freelancer**, I want **client-style mirroring from public signals** so that my proposal speaks their language.
    
*   As **privacy**, I want **strict PII redaction on generated variants** so that confidential data isn’t leaked.
    
*   As a **system**, I want **multi-variant generation** (short/long) so that I can pick what fits the job.
    

### Flow

1.  PersonalizeProposalCommand(proposal\_id, client\_context) → GenerateVariant() | Redact() → **Outbox:** proposal.personalized.v1
    
2.  ApplyPersonalizationCommand(variant\_id) → Apply() → **Outbox:** personalization.applied.v1
    

### Projections

*   personalization\_read
    

### Events

proposal.personalized.v1, personalization.applied.v1

### RBAC/SLO

*   **RBAC:** **OWNER**
    
*   **SLO:** **P95 < 350 ms**.
    

18.7 ab\_testing/
-----------------

### Stories

*   As a **freelancer**, I want **automatic traffic allocation (bandit)** so that better variants get more exposure.
    
*   As a **system**, I want **sample-ratio-mismatch and power checks** so that tests are trustworthy.
    
*   As a **user**, I want **guardrails for minimum sample size and max duration** so that I don’t declare winners too early.
    
*   As **product**, I want **per-metric selection** (view→response→hire) so that tests optimize the right outcome.
    

### Flow

1.  StartABTestCommand(proposal\_id, variants) → Start() → **Outbox:** abtest.started.v1
    
2.  DeclareVariantWinnerCommand(test\_id) → Compute() → **Outbox:** abtest.winner.declared.v1
    
3.  EndABTestCommand(test\_id) → Close() → **Outbox:** abtest.ended.v1
    

### Projections

*   ab\_test\_read
    

### Events

abtest.started.v1, abtest.winner.declared.v1, abtest.ended.v1

### RBAC/SLO

*   **RBAC:** **OWNER**
    
*   **SLO:** **P95 < 220 ms**.
    

19 - PROFILE & PORTFOLIO
========================

19.1 portfolio\_link/
---------------------

### Stories

*   As a **freelancer**, I want **relevance scoring to the job** so that suggested items bubble to the top.
    
*   As a **system**, I want **broken-link/file checks** so that clients never hit dead ends.
    
*   As a **freelancer**, I want **per-item analytics** (views, dwell, assists) so that I learn what convinces clients.
    
*   As **privacy**, I want **visibility controls** (public/share-link/under-NDA) so that sensitive work is protected.
    

### Flow

1.  LinkPortfolioItemCommand(proposal\_id, item\_id, note?) → Add() → **Outbox:** portfolio.linked.v1
    
2.  UnlinkPortfolioItemCommand(proposal\_id, item\_id) → Remove() → **Outbox:** portfolio.unlinked.v1
    
3.  ReorderPortfolioItemsCommand(proposal\_id, order\[\]) → Apply() → **Outbox:** portfolio.reordered.v1
    

### Projections

*   portfolio\_link\_read
    

### Events

portfolio.linked.v1, portfolio.unlinked.v1, portfolio.reordered.v1

### RBAC/SLO

*   **RBAC:** **OWNER**
    
*   **SLO:** **P95 < 140 ms**.
    

19.2 skill\_match/
------------------

### Stories

*   As a **freelancer**, I want **missing-skills suggestions with learning links** so that I can close gaps.
    
*   As a **system**, I want **collaborator recommendations to cover gaps** so that team bids are stronger.
    
*   As a **user**, I want **history of skill-match over time** so that I can see progress.
    

### Flow

1.  ComputeSkillMatchCommand(proposal\_id) → ExtractSkills() | CompareToJob() → **Outbox:** skills.matched.v1
    
2.  UpdateSkillMatchOnProfileChangeEvent(user\_id) → RecomputeAffected() → **Outbox:** skill.match.score.updated.v1
    

### Projections

*   skill\_match\_read
    

### Events

skills.matched.v1, skill.match.score.updated.v1

### RBAC/SLO

*   **RBAC:** **OWNER**
    
*   **SLO:** **P95 < 260 ms**.
    

19.3 video\_introduction/
-------------------------

### Stories

*   As a **freelancer**, I want **auto-captions and transcript editing** so that accessibility is solid.
    
*   As a **system**, I want **background noise and PII detection** so that videos are safe and clear.
    
*   As a **freelancer**, I want **multiple takes with performance hints** so that I can iterate to a strong intro.
    
*   As a **client**, I want **quick preview and playback speed** so that I can evaluate faster.
    

### Flow

1.  UploadVideoIntroCommand(proposal\_id, file\_id) → ValidateRef() | StartTranscode() | EnqueueAVScan() → **Outbox:** video.uploaded.v1
    
2.  TranscribeVideoIntroEvent(file\_id, text) → PersistTranscript() → **Outbox:** video.transcribed.v1
    

### Projections

*   video\_intro\_read (duration, transcript, scan\_status)
    

### Events

video.uploaded.v1, video.transcribed.v1

### RBAC/SLO

*   **RBAC:** **OWNER**
    
*   **SLO:** upload **P95 < 200 ms**; scans **async**.
    

19.4 reference/
---------------

### Stories

*   As a **freelancer**, I want **request-reference links with tokenized verification** so that referees can respond easily.
    
*   As a **system**, I want **freshness and relationship checks** so that testimonials are credible.
    
*   As a **freelancer**, I want **private vs public display controls** so that I tailor visibility per proposal.
    
*   As **trust**, I want **conflict-of-interest flags** so that biased references are transparent.
    

### Flow

1.  AddReferenceCommand(proposal\_id, ref\_contact\_hash, text) → AntiPiiLint() | Persist() → **Outbox:** reference.added.v1
    
2.  VerifyReferenceCommand(reference\_id, token) → MarkVerified() → **Outbox:** reference.verified.v1
    

### Projections

*   reference\_read
    

### Events

reference.added.v1, reference.verified.v1

### RBAC/SLO

*   **RBAC:** **OWNER**
    
*   **SLO:** **P95 < 180 ms**.
    

19.5 membership\_perk/
----------------------

### Stories

*   As a **member**, I want **auto-apply perks based on rules** (e.g., urgent job) so that I don’t miss benefits.
    
*   As **product**, I want **perk ROI tracking per proposal** so that value is measurable.
    
*   As a **team lead**, I want **perk sharing quotas across members** so that benefits suit team workflows.
    
*   As a **user**, I want a **perk ledger and remaining balance** so that usage is clear.
    

### Flow

1.  ApplyPerkToProposalCommand(proposal\_id, perk\_type) → ValidateEligibility() | Apply() → **Outbox:** perk.applied.v1
    
2.  ExhaustPerkEvent(perk\_grant\_id) → MarkExhausted() → **Outbox:** perk.exhausted.v1
    

### Projections

*   membership\_perk\_read
    

### Events

perk.applied.v1, perk.exhausted.v1

### RBAC/SLO

*   **RBAC:** **OWNER**
    
*   **SLO:** **P95 < 160 ms**.
    

20 - LIFECYCLE MANAGEMENT
========================

20.1 expiration/
----------------

### Stories

*   As a **freelancer**, I want **category-based default expiries** so that I don’t overthink settings.
    
*   As a **system**, I want **snooze and escalation paths** so that important proposals get timely attention.
    
*   As a **user**, I want **bulk extend/shorten** so that I can adjust many proposals at once.
    
*   As a **system**, I want **conditional expiry pauses during active negotiations** so that we don’t auto-expire mid-deal.
    

### Flow

1.  SetProposalExpiryCommand(proposal\_id, expires\_at) → ValidateFuture() | Save() → **Outbox:** proposal.expiry.set.v1
    
2.  ExtendProposalExpiryCommand(proposal\_id, new\_expires\_at) → Validate() → **Outbox:** proposal.expiry.extended.v1
    
3.  NotifyUpcomingExpiryJob() → Send() → **Outbox:** proposal.expiry.notified.v1
    
4.  ExpireProposalsJob() → MarkExpired() → **Outbox:** proposal.expired.v1
    

### Projections

*   proposal\_expiry\_read
    

### Events

proposal.expiry.set.v1, proposal.expiry.extended.v1, proposal.expiry.notified.v1, proposal.expired.v1

### RBAC/SLO

*   **RBAC:** **OWNER/SYSTEM**
    
*   **SLO:** set/extend **P95 < 150 ms**.
    

20.2 withdrawal/
----------------

### Stories

*   As a **freelancer**, I want **soft-withdraw with a courteous message template** so that I maintain goodwill.
    
*   As a **system**, I want **reversible withdraw within a window** so that I can re-enter if circumstances change.
    
*   As **product**, I want **analytics on withdraw reasons** so that we can improve matching.
    
*   As a **freelancer**, I want **auto-notify collaborators and adjust pipeline** so that the team stays aligned.
    

### Flow

1.  WithdrawProposalCommand(proposal\_id, reason) → GuardWithdraw() | MarkWithdrawn() | MaybeRefundConnects() → **Outbox:** proposal.withdrawn.v1
    
2.  UpdateWithdrawalReasonCommand(proposal\_id, reason) → Update() → **Outbox:** withdrawal.reason.updated.v1
    

### Projections

*   proposal\_withdrawal\_read
    

### Events

proposal.withdrawn.v1, withdrawal.reason.updated.v1

### RBAC/SLO

*   **RBAC:** **OWNER**
    
*   **SLO:** **P95 < 140 ms**.
    

20.3 archive/
-------------

### Stories

*   As a **freelancer**, I want **search and filters across archived proposals** so that I can still learn from history.
    
*   As a **system**, I want **auto-archive rules** (e.g., stale 90 days) so that the workspace stays clean.
    
*   As **compliance**, I want **export and retention windows** so that deletion follows policy.
    

### Flow

1.  ArchiveProposalCommand(proposal\_id) → Archive() → **Outbox:** proposal.archived.v1
    
2.  RestoreProposalCommand(proposal\_id) → Restore() → **Outbox:** proposal.restored.v1
    

### Projections

*   proposal\_archive\_read
    

### Events

proposal.archived.v1, proposal.restored.v1

### RBAC/SLO

*   **RBAC:** **OWNER**
    
*   **SLO:** **P95 < 140 ms**.
    

21 - TEAM & AUDIT
================

21.1 team/
----------

### Stories

*   As a **team lead**, I want **role-based permissions** (edit/comment/view) so that access is safe.
    
*   As **collaborators**, we want **internal-only notes separate from client view** so that we coordinate privately.
    
*   As a **system**, I want **activity analytics per member** so that coaching is data-driven.
    
*   As a **team**, I want **handoff/ownership change flows** so that continuity is maintained.
    

### Flow

1.  CreateTeamProposalCommand(job\_id, lead\_id, team\_members\[\]) → Validate() | CreateProposal(team\_mode) → **Outbox:** team.proposal.created.v1
    
2.  AddTeamMemberCommand(proposal\_id, user\_id, role) → Add() → **Outbox:** team.member.added.v1
    
3.  RemoveTeamMemberCommand(proposal\_id, user\_id) → Remove() → **Outbox:** team.member.removed.v1
    

### Projections

*   team\_proposal\_read
    

### Events

team.proposal.created.v1, team.member.added.v1, team.member.removed.v1

### RBAC/SLO

*   **RBAC:** **OWNER/TEAM LEAD**
    
*   **SLO:** **P95 < 200 ms**.
    

21.2 audit/
-----------

### Stories

*   As **compliance**, I want **tamper-evident hash chains** so that audit logs can be verified.
    
*   As **legal**, I want **exportable audit bundles** (JSON/PDF) so that I can respond to requests.
    
*   As **privacy**, I want **PII minimization in logs with role-gated access** so that audits don’t leak data.
    
*   As a **system**, I want **DSAR-friendly search** (actor/action/date) so that data requests are fast.
    

### Flow

1.  LogProposalActionCommand(actor\_id, proposal\_id, action, meta) → AppendWORM() → **Outbox:** proposal.audited.v1
    

### Projections

*   proposal\_audit\_read
    

### Events

proposal.audited.v1

### RBAC/SLO

*   **RBAC:** **SYSTEM**
    
*   **SLO:** append **P95 < 60 ms**.
    

22 - INTEGRATIONS & SECURITY
============================

22.1 integration/
-----------------

### Stories

*   As a **user**, I want **granular scopes** (read jobs, write contacts) so that integrations follow least privilege.
    
*   As a **system**, I want **token health checks and auto-reconnect** so that syncs are reliable.
    
*   As **ops**, I want **field mapping to/from CRM** so that data lands in the right places.
    
*   As a **user**, I want **conflict-resolution policies** (ours wins/theirs wins/merge) so that syncs are predictable.
    

### Flow

1.  ConnectIntegrationCommand(user\_id, provider) → OAuth() | PersistToken() → **Outbox:** integration.connected.v1
    
2.  DisconnectIntegrationCommand(integration\_id) → Revoke() → **Outbox:** integration.disconnected.v1
    
3.  SyncIntegrationCommand(integration\_id, scope) → FetchPush() → **Outbox:** integration.synced.v1
    

### Projections

*   integration\_read
    

### Events

integration.connected.v1, integration.disconnected.v1, integration.synced.v1

### RBAC/SLO

*   **RBAC:** **OWNER**
    
*   **SLO:** **P95 < 220 ms**.
    

22.2 secure\_share/
-------------------

### Stories

*   As a **freelancer**, I want **domain allow-lists and OTP links** so that only intended viewers get access.
    
*   As a **system**, I want **view-only mode with watermark** (email/time/IP) so that leaks are deterred.
    
*   As **security**, I want **device/browser fingerprint logging** so that suspicious access is traceable.
    
*   As a **user**, I want **per-link analytics** (views, unique viewers, geo) so that I know engagement and risk.
    

### Flow

1.  CreateShareLinkCommand(proposal\_id, acl, ttl, watermark?) → GenerateToken() | Persist() → **Outbox:** share.link.created.v1
    
2.  RevokeShareLinkCommand(share\_id) → Revoke() → **Outbox:** share.link.revoked.v1
    
3.  RecordShareAccessEvent(share\_id, viewer\_ctx) → Log() → **Outbox:** share.link.accessed.v1
    

### Projections

*   share\_link\_read, share\_access\_log\_read
    

### Events

share.link.created.v1, share.link.revoked.v1, share.link.accessed.v1

### RBAC/SLO

*   **RBAC:** **OWNER**
    
*   **SLO:** **P95 < 140 ms**.
    

22.3 localization/
------------------

### Stories

*   As a **freelancer**, I want **in-context translation editing** so that I can fine-tune phrasing.
    
*   As a **system**, I want **translation memory and glossary** so that repeated phrases stay consistent.
    
*   As a **user**, I want **locale-specific formatting** (dates, numbers, currency) so that content feels native.
    
*   As **product**, I want **reviewer workflows with native speakers** so that quality is high.
    

### Flow

1.  TranslateProposalCommand(proposal\_id, locale) → MachineTranslate() | Redact() → **Outbox:** proposal.translated.v1
    
2.  UpdateTranslationCommand(translation\_id, patch) → Apply() → **Outbox:** translation.updated.v1
    

### Projections

*   localization\_read
    

### Events

proposal.translated.v1, translation.updated.v1

### RBAC/SLO

*   **RBAC:** **OWNER**
    
*   **SLO:** **P95 < 260 ms**.
    

22.4 dispute\_prediction/
-------------------------

### Stories

*   As a **system**, I want **explainable risk factors and confidence** so that actions are transparent.
    
*   As **ops**, I want **playbook suggestions** (escrow, clearer milestones) so that risk is mitigated early.
    
*   As **privacy**, I want **opt-out and sensitive-feature controls** so that prediction respects user rights.
    
*   As **product**, I want **continuous learning from outcomes and false positives** so that the model improves.
    

### Flow

1.  PredictDisputeRiskCommand(proposal\_id) → Score() → **Outbox:** dispute.risk.predicted.v1
    
2.  MitigateDisputeRiskCommand(proposal\_id, recommendation) → Apply() → **Outbox:** dispute.risk.mitigated.v1
    

### Projections

*   dispute\_risk\_read
    

### Events

dispute.risk.predicted.v1, dispute.risk.mitigated.v1

### RBAC/SLO

*   **RBAC:** **SYSTEM**
    
*   **SLO:** **P95 < 300 ms**.