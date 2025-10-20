


## **1 - CORE PROPOSAL DOMAIN**

### 1.1 proposal/

#### User Stories
- As a **freelancer**, I want to **create a proposal** with title, description, pricing, and timeline so that clients can evaluate my offer.
- As a **freelancer**, I want to **save proposals as drafts** so that I can complete them later before submitting.
- As a **freelancer**, I want to **submit proposals** so that they become visible to clients.
- As a **freelancer**, I want to **update proposal details** before submission so that information stays accurate.
- As a **freelancer**, I want to **withdraw proposals** so that I can exit opportunities that no longer fit.
- As a **system**, I want to **validate proposal status transitions** (Draft → Submitted → Shortlisted → Accepted/Rejected) so that workflows are consistent.
- As a **system**, I want to **enforce job requirements** (connects cost, screening questions) so that platform rules are followed.
- As a **client**, I want to **view proposals** with all details so that I can make informed hiring decisions.
- As a **freelancer**, I want to **track proposal views** so that I know client interest levels.

#### Flow
1. **CreateProposalCommand**(job_id, freelancer_id, initial_data, is_draft) → ValidateJobIsOpen() | CheckFreelancerEligibility() | ValidateInitialData() | AntiPiiLint(description) → CreateProposal(status=Draft) → **Outbox:** proposal.created.v1
2. **UpdateProposalCommand**(proposal_id, updates) → GuardEditableWindow() | ValidatePatch() | AntiPiiLint() → UpdateProposal() → **Outbox:** proposal.updated.v1
3. **SubmitProposalCommand**(proposal_id) → ValidateCompleteness() | ValidateScreeningAnswers() | ReserveConnects(subscriptions-be) | MarkSubmitted() → **Outbox:** proposal.submitted.v1
4. **WithdrawProposalCommand**(proposal_id, reason) → GuardWithdrawState() | MarkWithdrawn() | MaybeRefundConnects() → **Outbox:** proposal.withdrawn.v1
5. **AcceptProposalCommand**(proposal_id, client_id, acceptance_terms) → ValidateClientOwnership(jobs-be) | MarkAccepted() | TriggerContractCreation(contracts-be) → **Outbox:** proposal.accepted.v1
6. **RejectProposalCommand**(proposal_id, client_id, reason) → ValidateClientOwnership() | MarkRejected() → **Outbox:** proposal.rejected.v1
7. **ShortlistProposalCommand**(proposal_id, client_id) → ValidateClientOwnership() | MarkShortlisted() → **Outbox:** proposal.shortlisted.v1
8. **GetProposalQuery**(proposal_id) → AuthorizeAccess() | Fetch() → ProposalDTO
9. **ListProposalsQuery**(filters, pagination) → ApplyFilters() | Fetch() → ProposalListDTO
10. **GetProposalsByJobQuery**(job_id, client_id) → AuthorizeClient() | Fetch() → ProposalListDTO

#### Projections
- proposal_read
- proposal_status_read
- proposal_by_job_read
- proposal_by_freelancer_read

#### Events Published
- proposal.created.v1
- proposal.updated.v1
- proposal.submitted.v1
- proposal.withdrawn.v1
- proposal.accepted.v1
- proposal.rejected.v1
- proposal.shortlisted.v1
- proposal.viewed.v1

#### Events Consumed
- job.published.v1 (from jobs-be - to validate job exists)
- job.closed.v1 (from jobs-be - to prevent new submissions)
- user.verified.v1 (from users-be - to check freelancer eligibility)
- connects.debited.v1 (from subscriptions-be - to confirm connects payment)

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (create/update/submit/withdraw), CLIENT (accept/reject/shortlist/view), PUBLIC (view published if job allows)
- **SLO:** P95 < 300ms (create/update/submit), P95 < 200ms (read), P95 < 150ms (accept/reject)

---

### 1.2 cover_letter/

#### User Stories
- As a **freelancer**, I want to **write a compelling cover letter** so that I differentiate myself from competitors.
- As a **freelancer**, I want to **update my cover letter** before submission so that messaging stays relevant.
- As a **system**, I want to **enforce word count limits** so that quality is maintained.
- As a **system**, I want to **analyze tone and readability** so that freelancers receive quality feedback.
- As a **system**, I want to **perform PII and profanity checks** so that content is safe.

#### Flow
1. **CreateCoverLetterCommand**(proposal_id, content) → ValidateWordCount() | AnalyzeTone() | AntiPiiLint() | AntiProfanity() → StoreCoverLetter() → **Outbox:** cover_letter.created.v1
2. **UpdateCoverLetterCommand**(proposal_id, content) → GuardEditableWindow() | ValidateWordCount() | AnalyzeTone() | AntiPiiLint() → UpdateCoverLetter() → **Outbox:** cover_letter.updated.v1
3. **GetCoverLetterQuery**(proposal_id) → AuthorizeAccess() | Fetch() → CoverLetterDTO

#### Projections
- cover_letter_read
- cover_letter_quality_read

#### Events Published
- cover_letter.created.v1
- cover_letter.updated.v1

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (create/update), CLIENT/FREELANCER (view)
- **SLO:** P95 < 180ms (create/update), Min 50 words, Max 5000 words

---

### 1.3 attachment/

#### User Stories
- As a **freelancer**, I want to **attach relevant work samples** so that I demonstrate capabilities.
- As a **freelancer**, I want to **remove attachments** so that I can update my portfolio presentation.
- As a **system**, I want to **validate file types and sizes** so that uploads are safe.
- As a **system**, I want to **store file references only** (actual files in storage-be) so that concerns are separated.
- As a **system**, I want to **limit attachment count** so that abuse is prevented.

#### Flow
1. **AddProposalAttachmentCommand**(proposal_id, file_url, file_name, file_type, file_size) → ValidateFileType() | ValidateFileSize() | ValidateAttachmentLimit(max=10) | StoreReference() → **Outbox:** proposal_attachment.added.v1
2. **RemoveProposalAttachmentCommand**(attachment_id) → AuthorizeOwner() | DeleteReference() → **Outbox:** proposal_attachment.removed.v1
3. **ListProposalAttachmentsQuery**(proposal_id) → Fetch() → AttachmentListDTO

#### Projections
- proposal_attachments_read

#### Events Published
- proposal_attachment.added.v1
- proposal_attachment.removed.v1

#### Events Consumed
- storage.file.uploaded (to validate file exists)
- storage.file.deleted (to clean up references)

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (add/remove), CLIENT/FREELANCER (view)
- **SLO:** P95 < 200ms (add), P95 < 120ms (list), Max 10 attachments per proposal, Max 25MB per file

---

### 1.4 question_answer/

#### User Stories
- As a **freelancer**, I want to **answer screening questions** so that I meet job requirements.
- As a **freelancer**, I want to **update answers** before submission so that responses are accurate.
- As a **system**, I want to **validate required questions** so that compliance is enforced.
- As a **system**, I want to **support multiple answer types** (text, file upload, multiple choice) so that flexibility exists.
- As a **client**, I want to **view screening answers** so that I can assess candidate fit.

#### Flow
1. **AnswerQuestionCommand**(proposal_id, question_id, answer, attachment_ids[]) → ValidateQuestionExists(jobs-be) | ValidateAnswerType() | StoreAnswer() → **Outbox:** question.answered.v1
2. **UpdateAnswerCommand**(answer_id, new_answer) → GuardEditableWindow() | ValidateAnswerType() | UpdateAnswer() → **Outbox:** answer.updated.v1
3. **ValidateScreeningAnswersQuery**(proposal_id) → FetchRequiredQuestions(jobs-be) | CheckCompleteness() → ValidationResultDTO
4. **GetProposalAnswersQuery**(proposal_id) → AuthorizeAccess() | Fetch() → AnswerListDTO

#### Projections
- proposal_answers_read

#### Events Published
- question.answered.v1
- answer.updated.v1

#### Events Consumed
- job.screening.configured.v1 (from jobs-be - to know required questions)

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (answer/update), CLIENT (view)
- **SLO:** P95 < 180ms (answer/update)

---

### 1.5 milestone/

#### User Stories
- As a **freelancer**, I want to **define milestones** (description, amount, due date) so that scope and billing are clear.
- As a **client**, I want to **validate that sum(milestones) matches fixed-price total** so that expectations align.
- As a **system**, I want to **ensure milestones are compatible with contracts/payments** so that downstream flows are seamless.
- As a **freelancer**, I want to **reorder or remove milestones** before submission so that my plan is coherent.
- As a **freelancer**, I want to **specify deliverables per milestone** so that output expectations are defined.

#### Flow
1. **CreateMilestoneCommand**(proposal_id, description, amount, due_date, deliverables[]) → ValidatePricingModel(fixed_or_hybrid) | ValidateAmount() | AntiPiiLint(description) → CreateMilestone(seq) | ValidateSumAgainstPricing() → **Outbox:** proposal.milestone.created.v1
2. **UpdateMilestoneCommand**(milestone_id, updates) → GuardEditableWindow() | ValidatePatch() | RecalcTotals() → **Outbox:** proposal.milestone.updated.v1
3. **ReorderMilestonesCommand**(proposal_id, new_order[]) → ValidatePermutation() | ApplyOrder() → **Outbox:** proposal.milestone.reordered.v1
4. **RemoveMilestoneCommand**(milestone_id) → GuardEditableWindow() | Delete() | RecalcTotals() → **Outbox:** proposal.milestone.removed.v1
5. **GetProposalMilestonesQuery**(proposal_id) → Fetch() → MilestoneListDTO

#### Projections
- proposal_milestones_read
- proposal_pricing_read
- milestone_audit_read

#### Events Published
- proposal.milestone.created.v1
- proposal.milestone.updated.v1
- proposal.milestone.reordered.v1
- proposal.milestone.removed.v1

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (create/update/reorder/remove), CLIENT/FREELANCER (view)
- **SLO:** P95 < 200ms (create/update/remove), Max 20 milestones per proposal

---

## **2 - BIDDING SYSTEM DOMAIN**

### 2.1 bid/

#### User Stories
- As a **freelancer**, I want to **place a bid** (amount) so that I can compete dynamically on price.
- As a **freelancer**, I want to **update my bid** so that I can respond to market conditions.
- As a **freelancer**, I want to **withdraw a bid** so that I can exit bidding competition.
- As a **client**, I want to **see transparent ranking** and distribution stats so that I quickly gauge competitiveness.
- As a **system**, I want to **maintain full bid history** with currency normalization so that analytics and anomaly detection are possible.
- As a **freelancer**, I want **hard floors/ceilings per job** so that I don't breach client constraints.
- As a **freelancer**, I want to **see my bid rank** among all proposals so that I know my competitive position.

#### Flow
1. **PlaceBidCommand**(proposal_id, amount, currency) → ValidateJobIsOpen() | NormalizeCurrency(amount, currency, job_currency) | ValidateMinMax(job_rules) | ValidateSingleActiveBid(proposal) → UpsertBid(state=Active) | RecomputeRank(job_id) → **Outbox:** bid.placed.v1 + (if rank changed: bid.rank.changed.v1)
2. **UpdateBidCommand**(bid_id, amount, currency) → ValidateOwnership() | NormalizeCurrency() | ValidateMinMax() → UpdateBid() | RecomputeRank() → **Outbox:** bid.updated.v1 + (bid.rank.changed.v1 if applicable)
3. **WithdrawBidCommand**(bid_id, reason) → GuardWithdrawState() | MarkRetracted() | RecomputeRank() → **Outbox:** bid.retracted.v1 + (bid.rank.changed.v1 if applicable)
4. **SyncBidFromInviteCommand**(invite_id) → MapInviteToProposal() | SeedOrUpdateBid() | RecomputeRank() → **Outbox:** bid.synced_from_invite.v1
5. **GetBidQuery**(proposal_id) → AuthorizeOwner() | Fetch() → BidDTO
6. **GetBidRankQuery**(proposal_id) → ComputeRank() → BidRankDTO
7. **GetJobBidDistributionQuery**(job_id, client_id) → AuthorizeClient() | AggregateStats() → BidDistributionDTO

#### Projections
- bid_read
- bid_history_read
- bid_rank_read

#### Events Published
- bid.placed.v1
- bid.updated.v1
- bid.retracted.v1
- bid.rank.changed.v1
- bid.synced_from_invite.v1

#### Events Consumed
- job.invitation.sent.v1 (from jobs-be - to sync target rate)

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (place/update/withdraw), CLIENT (view distribution)
- **SLO:** P95 < 160ms (place/update/withdraw), Rank recompute atomic in-tx, Idempotent by (proposal_id, amount_norm, ts_bucket)

---

### 2.2 bid_strategy/

#### User Stories
- As a **freelancer**, I want to **create bid strategies** (fixed, step-down, undercut, auto-rebalance) so that my bid adapts to competition.
- As a **system**, I want **guardrails** (min floor, max daily changes, cool-downs) so that pricing remains stable.
- As a **freelancer**, I want a **preview/simulation** so that I understand strategy impact before enabling.
- As a **freelancer**, I want to **activate/deactivate strategies** so that I control automation.
- As a **system**, I want to **auto-adjust bids** based on active strategies so that competition is dynamic.

#### Flow
1. **CreateBidStrategyCommand**(proposal_id, strategy_type, params) → ValidateParamsByType() | GuardCompatibilityWithJob() → CreateStrategy(state=Draft) → **Outbox:** bid.strategy.created.v1
2. **SimulateBidStrategyCommand**(strategy_id, horizon_days) → LoadMarketSignals(job_id) | RunSimulation() → **Outbox:** bid.strategy.simulated.v1
3. **ActivateBidStrategyCommand**(strategy_id) → GuardNoActiveStrategy() | Activate() → **Outbox:** bid.strategy.activated.v1
4. **AutoAdjustBidsJob**() → SelectActiveStrategies() | ComputeNextAmount() | ApplyIfWithinGuardrails() | RecomputeRank() → **Outbox:** bid.auto.adjusted.v1 + (bid.rank.changed.v1 if applicable) (system-triggered)
5. **DeactivateBidStrategyCommand**(strategy_id) → Deactivate() → **Outbox:** bid.strategy.deactivated.v1
6. **GetBidStrategyQuery**(strategy_id) → AuthorizeOwner() | Fetch() → BidStrategyDTO

#### Projections
- bid_strategy_read
- bid_strategy_state_read
- bid_strategy_simulations_read

#### Events Published
- bid.strategy.created.v1
- bid.strategy.simulated.v1
- bid.strategy.activated.v1
- bid.auto.adjusted.v1
- bid.strategy.deactivated.v1

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (create/simulate/activate/deactivate), SYSTEM (auto-adjust)
- **SLO:** P95 < 150ms (control ops), Auto-job batch 10k < 60s, Idempotent by (strategy_id, tick)

---

### 2.3 bid_notification/

#### User Stories
- As a **freelancer**, I want **outbid alerts** with threshold options so that I react only when it matters.
- As a **system**, I want **reliable queueing**, exponential backoff, and per-channel fallbacks so that notifications are dependable.
- As a **freelancer**, I want **quiet hours** so that I'm not spammed.
- As a **freelancer**, I want to **configure alert sensitivity** (immediate, daily digest, off) so that I control frequency.

#### Flow
1. **TriggerOutbidAlertCommand**(proposal_id, outbidder_rank_change) → CheckThreshold() | CheckQuietHours() | QueueNotification(communications-be) → **Outbox:** bid.outbid.alert.triggered.v1
2. **ConfigureAlertPreferencesCommand**(freelancer_id, sensitivity, quiet_hours) → UpdatePreferences() → **Outbox:** bid.alert.preferences.updated.v1
3. **SendBatchAlertsJob**() → AggregateAlerts() | SendDigest(communications-be) → **Outbox:** bid.alerts.batch.sent.v1 (system-triggered)

#### Projections
- bid_alert_preferences_read
- bid_alert_history_read

#### Events Published
- bid.outbid.alert.triggered.v1
- bid.alert.preferences.updated.v1
- bid.alerts.batch.sent.v1

#### Events Consumed
- bid.rank.changed.v1 (to trigger outbid alerts)

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (configure preferences), SYSTEM (trigger/send alerts)
- **SLO:** P95 < 100ms (trigger), Async (notification delivery with DLQ)

---

### 2.4 auction/

#### User Stories
- As a **client**, I want to **open auction windows** (duration, slots, reserve price) so that competitive bidding drives optimal pricing.
- As a **freelancer**, I want to **place auction bids** with incremental bidding so that I can compete fairly.
- As a **freelancer**, I want to **cancel auction bids** before close so that I can adjust strategy.
- As a **system**, I want to **track top slots** in real-time so that rankings are live.
- As a **system**, I want to **end auctions** and assign winning slots so that results are definitive.
- As a **system**, I want to **hold funds/connects** during bidding so that commitment is guaranteed.

#### Flow
1. **OpenAuctionWindowCommand**(job_id, slots, reserve_price, min_increment, ttl) → ValidateJob() | CreateAuction() → **Outbox:** auction.opened.v1
2. **PlaceAuctionBidCommand**(job_id, proposal_id, amount) → GuardWindowOpen() | ValidateIncrement() | HoldFunds(connects_or_currency) | UpsertAuctionBid() | RecalcTopSlots() → **Outbox:** auction.bid.placed.v1 + (auction.top.changed.v1 if applicable)
3. **CancelAuctionBidCommand**(auction_bid_id) → GuardCancelable() | ReleaseHold() | MarkCanceled() | RecalcTopSlots() → **Outbox:** auction.bid.canceled.v1 + (auction.top.changed.v1 if applicable)
4. **EndAuctionCommand**(job_id) → GuardEnd() | CloseAndAssignSlots() | CaptureWinningHolds() | ReleaseLosingHolds() → **Outbox:** auction.ended.v1, auction.top.assigned.v1
5. **GetAuctionStatusQuery**(job_id) → Fetch() → AuctionStatusDTO
6. **GetAuctionLeaderboardQuery**(job_id) → FetchTopBids() → LeaderboardDTO

#### Projections
- auction_status_read
- auction_bids_read
- auction_slot_assignment_read

#### Events Published
- auction.opened.v1
- auction.bid.placed.v1
- auction.bid.canceled.v1
- auction.top.changed.v1
- auction.ended.v1
- auction.top.assigned.v1

#### RBAC/SLO
- **RBAC:** CLIENT (open/end), FREELANCER/OWNER (place/cancel bid), SYSTEM (end/assign), PUBLIC (view status/leaderboard)
- **SLO:** P95 < 170ms (place bid), P95 < 2s (reconciliation job), Idempotent by (job_id, proposal_id, amount, ts_bucket)

---

### 2.5 bid_anomaly_detection/

#### User Stories
- As a **trust analyst**, I want **anomalies flagged** (suspicious undercut, collusion patterns) so that marketplace quality is preserved.
- As a **system**, I want **triage states** (Open, UnderReview, Closed) and evidence so that reviews are trackable.
- As a **system**, I want to **auto-close benign anomalies** so that analysts focus on real threats.

#### Flow
1. **DetectBidAnomalyCommand**(job_id) → FetchBidDistribution() | ComputeZScores()/IQR() | RunCollusionHeuristics() → UpsertAnomalies() → **Outbox:** bid.anomaly.detected.v1 (system-triggered)
2. **UpdateAnomalyReviewCommand**(anomaly_id, state, notes) → ValidateStateTransition() | Update() → **Outbox:** bid.anomaly.review.updated.v1
3. **AutoCloseBenignAnomaliesJob**() → ApplyAutoRules() | Close() → **Outbox:** bid.anomaly.auto_closed.v1 (system-triggered)
4. **GetBidAnomaliesQuery**(job_id) → AuthorizeTrust() | Fetch() → AnomalyListDTO

#### Projections
- bid_anomaly_read
- bid_anomaly_audit_read

#### Events Published
- bid.anomaly.detected.v1
- bid.anomaly.review.updated.v1
- bid.anomaly.auto_closed.v1

#### RBAC/SLO
- **RBAC:** SYSTEM (detect), TRUST/ADMIN (review), ANALYST (view)
- **SLO:** Detect batch 50k < 90s, Idempotent by (job_id, proposal_id, feature_hash, day)

---

## **3 - CONNECTS & BOOST DOMAIN**

### 3.1 connect/

#### User Stories
- As a **freelancer**, I want **connects reserved/debited on submission** so that platform economics are enforced.
- As a **system**, I want **tiered costs by job popularity** and time-of-day so that pricing is dynamic and fair.
- As a **freelancer**, I want to **view balance, aging buckets, and ledger** so that I can plan submissions.
- As a **freelancer**, I want to **purchase connects** so that I can submit proposals.
- As a **system**, I want to **track connect transactions** so that accounting is accurate.

#### Flow
1. **ReserveConnectsOnSubmitCommand**(job_id, freelancer_id) → GetTierCost(job_id) | CheckBalance() | CreateReservation(hold) → **Outbox:** connects.reserved.v1
2. **FinalizeSubmissionCommand**(proposal_id) → ConsumeReservation() | DebitLedger() → **Outbox:** connects.debited.v1
3. **CancelSubmissionCommand**(proposal_id) → ReleaseReservation() → **Outbox:** connects.released.v1
4. **CreditConnectsCommand**(user_id, amount, reason) → CreditLedger() → **Outbox:** connects.credited.v1
5. **GetConnectBalanceQuery**(user_id) → Fetch() → ConnectBalanceDTO
6. **GetConnectTransactionsQuery**(user_id, filters) → Fetch() → TransactionListDTO

#### Projections
- connect_balance_read
- connect_transactions_read
- connect_tier_read

#### Events Published
- connects.reserved.v1
- connects.debited.v1
- connects.released.v1
- connects.credited.v1
- connects.tier.changed.v1

#### Events Consumed
- subscription.connects.purchased.v1 (from subscriptions-be - to credit connects)

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (view balance/transactions), SYSTEM (reserve/debit/release/credit)
- **SLO:** P95 < 120ms (reserve), Strong idempotency on (freelancer_id, job_id), Ledger ACID

---

### 3.2 connect_refund/

#### User Stories
- As a **freelancer**, I want **automatic and manual refunds** (job closed early, spam job) so that I'm protected.
- As a **system**, I want **eligibility policies and SLA-based processing** so that fairness and throughput are balanced.
- As a **support agent**, I want to **process refund requests** so that disputes are resolved.
- As a **freelancer**, I want to **track refund status** so that I know outcomes.

#### Flow
1. **RequestConnectRefundCommand**(proposal_id, reason) → CheckEligibility(policy_matrix) | OpenCase(state=Pending) → **Outbox:** connect.refund.requested.v1
2. **AutoEvaluateRefundsJob**() → ApplyPolicy() | ApproveOrEscalate() → **Outbox:** connect.refund.auto_processed.v1 (system-triggered)
3. **ProcessConnectRefundCommand**(case_id, decision, agent_id) → ApplyDecision() | CreditLedgerIfApproved() → **Outbox:** connect.refund.processed.v1 or connect.refund.denied.v1
4. **GetRefundStatusQuery**(case_id) → AuthorizeOwner() | Fetch() → RefundStatusDTO

#### Projections
- connect_refund_read
- connect_refund_policy_read

#### Events Published
- connect.refund.requested.v1
- connect.refund.auto_processed.v1
- connect.refund.processed.v1
- connect.refund.denied.v1

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (request/view), SUPPORT/TRUST (process), SYSTEM (auto-evaluate)
- **SLO:** Auto-job P95 case < 2s, Idempotent by (proposal_id, policy_snapshot_hash)

---

### 3.3 boost/

#### User Stories
- As a **freelancer**, I want to **purchase proposal boosts** (levels) so that my proposal appears higher in rankings.
- As a **freelancer**, I want to **upgrade boost levels** so that I can increase visibility further.
- As a **system**, I want **expiry handling** so that boosts don't persist indefinitely.
- As a **system**, I want **conflict resolution with auctions** so that boosts don't double-count.
- As a **freelancer**, I want to **preview boost impact** so that I understand value before purchasing.

#### Flow
1. **PurchaseBoostCommand**(proposal_id, boost_level, term) → ChargeWallet(financial-be) | ActivateBoost(level, expires_at) → **Outbox:** proposal.boost.purchased.v1
2. **UpgradeBoostCommand**(proposal_id, new_level) → GuardUpgradePath() | ProrateCharge() | ApplyUpgrade() → **Outbox:** proposal.boost.upgraded.v1
3. **ExpireBoostsJob**() → FindExpired() | MarkExpired() | RecomputeRanking() → **Outbox:** proposal.boost.expired.v1, proposal.ranking.updated.v1 (system-triggered)
4. **PreviewBoostImpactQuery**(proposal_id, boost_level) → SimulateRanking() → BoostImpactDTO
5. **GetProposalBoostQuery**(proposal_id) → Fetch() → BoostDTO



#### Projections

*   proposal\_boost\_read
    
*   boost\_ledger\_read
    
*   boost\_impact\_simulations\_read
    

#### Events Published

*   proposal.boost.purchased.v1
    
*   proposal.boost.upgraded.v1
    
*   proposal.boost.expired.v1
    
*   proposal.ranking.updated.v1
    

#### Events Consumed

*   payment.captured.v1 (from financial-be - to confirm boost payment)
    

#### RBAC/SLO

*   **RBAC:** FREELANCER/OWNER (purchase/upgrade/preview/view), SYSTEM (expire)
    
*   **SLO:** P95 < 180ms (purchase/upgrade), Idempotent by (proposal\_id, level, purchase\_ts\_bucket)


## **4 - TEMPLATES & RATE CARDS DOMAIN**

### 4.1 template/

#### User Stories
- As a **freelancer**, I want to **create reusable proposal templates** with categories and placeholders so that I draft faster.
- As a **system**, I want **placeholder validation** so that personalization is correct.
- As a **freelancer**, I want **version history and archiving** so that I can maintain quality over time.
- As a **freelancer**, I want to **localize templates** for different languages so that I can work globally.
- As a **freelancer**, I want to **categorize templates** by job type so that selection is easier.
- As a **freelancer**, I want to **clone templates** so that I can create variations quickly.

#### Flow
1. **CreateTemplateCommand**(user_id, title, content, categories[], placeholders[]) → ValidatePlaceholders(content) | AntiPiiLint(content) → CreateTemplate(state=Active, version=1) → **Outbox:** proposal.template.created.v1
2. **UpdateTemplateCommand**(template_id, updates) → GuardOwnership() | ValidatePlaceholders() | AppendVersion() → **Outbox:** proposal.template.updated.v1
3. **ArchiveTemplateCommand**(template_id) → Archive() → **Outbox:** proposal.template.archived.v1
4. **LocalizeTemplateCommand**(template_id, locale, translated_content) → TranslateContent() | PersistVariant() → **Outbox:** proposal.template.localized.v1
5. **CloneTemplateCommand**(template_id, new_title) → CopyTemplate() | CreateNew() → **Outbox:** proposal.template.cloned.v1
6. **GetTemplateQuery**(template_id) → AuthorizeOwner() | Fetch() → TemplateDTO
7. **ListTemplatesQuery**(user_id, filters) → AuthorizeOwner() | Fetch() → TemplateListDTO
8. **GetTemplateVersionsQuery**(template_id) → AuthorizeOwner() | Fetch() → VersionListDTO

#### Projections
- template_read
- template_versions_read
- template_category_read
- template_localization_read

#### Events Published
- proposal.template.created.v1
- proposal.template.updated.v1
- proposal.template.archived.v1
- proposal.template.localized.v1
- proposal.template.cloned.v1

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (create/update/archive/localize/clone/view)
- **SLO:** P95 < 160ms (create/update), Idempotent by (user_id, title_hash, content_hash)

---

### 4.2 rate_card/

#### User Stories
- As a **freelancer**, I want to **create rate cards** with package tiers (Starter/Standard/Premium) so that pricing is consistent.
- As a **client**, I want **clear inclusions/exclusions** so that expectations are set.
- As a **system**, I want **currency normalization and taxes/fees modeling** so that downstream quotes are accurate.
- As a **freelancer**, I want to **version rate cards** so that I can track pricing changes over time.
- As a **freelancer**, I want to **archive old rate cards** so that my offerings stay current.

#### Flow
1. **CreateRateCardCommand**(user_id, currency, packages[]) → ValidatePackageSchema() | NormalizeCurrency() → CreateRateCard(state=Active, version=1) → **Outbox:** proposal.rate_card.created.v1
2. **UpdateRateCardCommand**(card_id, updates) → GuardOwnership() | Validate() | AppendVersion() → **Outbox:** proposal.rate_card.updated.v1
3. **ArchiveRateCardCommand**(card_id) → Archive() → **Outbox:** proposal.rate_card.archived.v1
4. **GetRateCardQuery**(card_id) → AuthorizeOwner() | Fetch() → RateCardDTO
5. **ListRateCardsQuery**(user_id, filters) → AuthorizeOwner() | Fetch() → RateCardListDTO
6. **GetRateCardVersionsQuery**(card_id) → AuthorizeOwner() | Fetch() → VersionListDTO

#### Projections
- rate_card_read
- rate_card_versions_read

#### Events Published
- proposal.rate_card.created.v1
- proposal.rate_card.updated.v1
- proposal.rate_card.archived.v1

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (create/update/archive/view)
- **SLO:** P95 < 170ms, Idempotent by (user_id, currency, packages_hash)

---

## **5 - ANALYTICS & TRACKING DOMAIN (CONSOLIDATED)**

### 5.1 performance/ (CONSOLIDATED: analytics + engagement + response_tracker + conversion + metrics + ranking + health)

#### User Stories
- As a **freelancer**, I want to **track proposal views** and CTR so that I can iterate on quality.
- As a **system**, I want **view beacons with bot filtering** so that stats are trustworthy.
- As a **freelancer**, I want to **know if a client viewed/responded** so that I time my follow-ups appropriately.
- As a **freelancer**, I want to **see conversion rates** (view-to-shortlist, shortlist-to-hire) so that I measure effectiveness.
- As a **freelancer**, I want **proposal health scores** so that I know what needs improvement.
- As a **freelancer**, I want **competitive ranking insights** so that I understand my position.
- As a **freelancer**, I want **benchmark comparisons** so that I measure against peers.
- As a **system**, I want to **track micro-interactions** (expand, scroll depth, portfolio clicks) so that interest signals are captured.

#### Flow
1. **TrackProposalViewCommand**(proposal_id, session_id, ua, ip_hash) → DedupWithin(30m) | BotFilter(ua, ip) | RecordView() → **Outbox:** proposal.viewed.v1
2. **TrackClientViewCommand**(proposal_id, viewer_context) → ValidateViewer(client_team) | Record() → **Outbox:** proposal.viewed.by_client.v1
3. **TrackClientResponseCommand**(proposal_id, response_type) → Record() → **Outbox:** client.responded.v1
4. **TrackEngagementCommand**(proposal_id, event_type, metadata) → Validate() | Dedup() | Record() → **Outbox:** proposal.engagement.interaction.tracked.v1
5. **CalculateInterestCommand**(proposal_id) → ComputeScore(weighted_micro_events) → **Outbox:** proposal.engagement.interest.scored.v1
6. **RecomputeAnalyticsCommand**(proposal_id|job_id|user_id) → AggregateViewsClicksResponses() → **Outbox:** proposal.analytics.updated.v1 (system-triggered)
7. **CalculateHealthScoreCommand**(proposal_id) → AnalyzeCompleteness() | AnalyzeQuality() | ComputeScore() → **Outbox:** proposal.health.scored.v1
8. **CalculateConversionMetricsCommand**(freelancer_id, time_period) → AggregateConversions() → **Outbox:** proposal.conversion.calculated.v1
9. **GetProposalAnalyticsQuery**(proposal_id) → AuthorizeOwner() | Fetch() → AnalyticsDTO
10. **GetProposalHealthQuery**(proposal_id) → AuthorizeOwner() | Fetch() → HealthScoreDTO
11. **GetProposalRankingQuery**(proposal_id) → ComputeRank() → RankingDTO
12. **GetBenchmarksQuery**(freelancer_id, category) → AggregatePeerData() → BenchmarksDTO
13. **GetEngagementInsightsQuery**(proposal_id) → AuthorizeOwner() | Fetch() → EngagementInsightsDTO

#### Projections
- proposal_analytics_read
- response_stats_read
- analytics_audit_read
- response_tracker_read
- client_response_latency_read
- proposal_engagement_read
- interest_score_read
- proposal_health_read
- conversion_metrics_read
- proposal_ranking_read
- benchmarks_read

#### Events Published
- proposal.viewed.v1
- proposal.viewed.by_client.v1
- client.responded.v1
- proposal.analytics.updated.v1
- proposal.engagement.interaction.tracked.v1
- proposal.engagement.interest.scored.v1
- proposal.health.scored.v1
- proposal.conversion.calculated.v1
- proposal.ranking.updated.v1

#### RBAC/SLO
- **RBAC:** Track: PUBLIC (signed URL), SYSTEM (recompute), FREELANCER/OWNER (view analytics/health/ranking), CLIENT (track views/responses)
- **SLO:** Track P95 < 80ms (view), P95 < 120ms (engagement), Recompute P95 < 200ms

---

## **6 - SIMILARITY & DEDUPLICATION DOMAIN (CONSOLIDATED)**

### 6.1 similarity/ (CONSOLIDATED: similarity + duplicate_check)

#### User Stories
- As a **system**, I want to **create fingerprints** (text, structural, semantic) so that proposals can be compared.
- As a **system**, I want to **cluster similar proposals** so that duplicates are minimized.
- As a **system**, I want to **detect duplicate proposals** so that spam is prevented.
- As a **freelancer**, I want a **differentiation score** so that I know how to stand out from peers.
- As a **system**, I want to **flag near-duplicates** so that quality is maintained.

#### Flow
1. **CreateFingerprintCommand**(proposal_id) → HashTokens() | GenerateEmbedding() | Persist() → **Outbox:** proposal.fingerprint.created.v1 (system-triggered on submission)
2. **ClusterProposalsJob**() → FormClusters() → **Outbox:** proposal.cluster.formed.v1 (system-triggered batch)
3. **DetectDuplicatesCommand**(proposal_id) → CompareFingerprintsWithinJob() | FlagIfNearDuplicate() → **Outbox:** proposal.duplicate.detected.v1 (system-triggered)
4. **ScoreDifferentiationCommand**(proposal_id) → CompareWithCluster() | ComputeScore() → **Outbox:** proposal.differentiation.scored.v1
5. **GetSimilarProposalsQuery**(proposal_id) → FetchClusterMembers() → SimilarProposalsDTO
6. **GetDuplicateCheckQuery**(proposal_id) → CheckDuplicates() → DuplicateStatusDTO
7. **GetDifferentiationScoreQuery**(proposal_id) → Fetch() → DifferentiationScoreDTO

#### Projections
- proposal_similarity_read
- proposal_clusters_read
- duplicate_detection_read
- differentiation_scores_read

#### Events Published
- proposal.fingerprint.created.v1
- proposal.cluster.formed.v1
- proposal.duplicate.detected.v1
- proposal.differentiation.scored.v1

#### Events Consumed
- proposal.submitted.v1 (to trigger fingerprinting)

#### RBAC/SLO
- **RBAC:** SYSTEM (fingerprint/cluster/detect), FREELANCER/OWNER (view similar/differentiation), MODERATOR (view duplicates)
- **SLO:** Fingerprint P95 < 200ms, Batch clustering 20k < 60s

---

## **7 - PORTFOLIO DOMAIN (CONSOLIDATED)**

### 7.1 portfolio/ (CONSOLIDATED: portfolio_link + portfolio_selector)

#### User Stories
- As a **freelancer**, I want to **link portfolio items** to proposals so that I demonstrate relevant work.
- As a **freelancer**, I want **auto-selected portfolio items** relevant to the job so that proof is targeted.
- As a **system**, I want **explainable relevance scoring** so that users trust the selection.
- As a **freelancer**, I want to **manually override auto-selection** so that I maintain control.
- As a **freelancer**, I want to **reorder portfolio items** so that best work appears first.
- As a **freelancer**, I want to **remove portfolio links** so that I can update my presentation.

#### Flow
1. **LinkPortfolioItemCommand**(proposal_id, portfolio_item_id, relevance_note) → ValidateItemExists(users-be) | CreateLink() → **Outbox:** portfolio.item.linked.v1
2. **AutoSelectPortfolioCommand**(proposal_id) → ComputeRelevance(job_skills, portfolio_tags) | PickTopK() → AttachLinks() → **Outbox:** proposal.portfolio.auto_selected.v1 (system-triggered)
3. **OverrideSelectionCommand**(proposal_id, item_ids[]) → ApplyOverride() → **Outbox:** proposal.portfolio.overridden.v1
4. **ReorderPortfolioItemsCommand**(proposal_id, new_order[]) → ValidatePermutation() | ApplyOrder() → **Outbox:** portfolio.items.reordered.v1
5. **UnlinkPortfolioItemCommand**(link_id) → AuthorizeOwner() | DeleteLink() → **Outbox:** portfolio.item.unlinked.v1
6. **GetProposalPortfolioQuery**(proposal_id) → Fetch() → PortfolioLinksDTO

#### Projections
- proposal_portfolio_read
- portfolio_selection_read

#### Events Published
- portfolio.item.linked.v1
- portfolio.item.unlinked.v1
- proposal.portfolio.auto_selected.v1
- proposal.portfolio.overridden.v1
- portfolio.items.reordered.v1

#### Events Consumed
- proposal.submitted.v1 (to trigger auto-selection)

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (link/override/reorder/unlink), SYSTEM (auto-select), CLIENT/FREELANCER (view)
- **SLO:** P95 < 250ms (auto-select), P95 < 150ms (manual operations)

---

## **8 - ENGAGEMENT DOMAIN (CONSOLIDATED)**

### 8.1 engagement/ (CONSOLIDATED: engagement + follow_up)

#### User Stories
- As a **freelancer**, I want to **schedule follow-ups** with clients so that I stay top-of-mind.
- As a **system**, I want **quiet hours and frequency caps** so that clients aren't spammed.
- As a **freelancer**, I want to **bump proposals** to refresh visibility so that older proposals get attention.
- As a **freelancer**, I want **gentle nudges** to clients so that I encourage responses without being pushy.
- As a **system**, I want to **trigger high-interest follow-ups** when engagement signals are strong so that timing is optimal.
- As a **freelancer**, I want to **track follow-up history** so that I don't repeat outreach.

#### Flow
1. **ScheduleFollowUpCommand**(proposal_id, when, template_id) → ValidateWindow() | ApplyQuietHours() | PersistSchedule() → **Outbox:** proposal.follow_up.scheduled.v1
2. **SendFollowUpTask**(schedule_id) → RenderTemplate() | Send(communications-be) | Record() → **Outbox:** proposal.follow_up.sent.v1 (system-triggered worker)
3. **BumpProposalCommand**(proposal_id) → GuardBumpRules() | UpdateVisibility() → **Outbox:** proposal.bumped.v1
4. **NudgeClientCommand**(proposal_id, nudge_type) → ValidateNudgePolicy() | Send(communications-be) → **Outbox:** client.nudged.v1
5. **TriggerHighInterestFollowUpCommand**(proposal_id) → GuardCoolDown() | ScheduleFollowUp() → **Outbox:** proposal.engagement.high_interest.detected.v1 (system-triggered)
6. **GetFollowUpHistoryQuery**(proposal_id) → AuthorizeOwner() | Fetch() → FollowUpHistoryDTO
7. **GetFollowUpScheduleQuery**(freelancer_id) → AuthorizeOwner() | Fetch() → ScheduleListDTO

#### Projections
- proposal_follow_up_read
- proposal_follow_up_schedule_read
- follow_up_history_read

#### Events Published
- proposal.follow_up.scheduled.v1
- proposal.follow_up.sent.v1
- proposal.bumped.v1
- client.nudged.v1
- proposal.engagement.high_interest.detected.v1

#### Events Consumed
- proposal.engagement.interest.scored.v1 (to trigger high-interest follow-ups)

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (schedule/bump/nudge/view), SYSTEM (send/trigger)
- **SLO:** Schedule P95 < 140ms, Send P95 < 2s, Idempotent by (proposal_id, schedule_slot)

---

## **9 - MODERATION & COMPLIANCE DOMAIN**

### 9.1 spam_detection/

#### User Stories
- As a **system**, I want to **automatically detect spam proposals** (low-quality text, suspicious links, mass submission) so that platform quality is preserved.
- As a **system**, I want to **score spam likelihood** so that prioritization is possible.
- As a **moderator**, I want to **review flagged spam** so that false positives are handled.
- As a **system**, I want to **clear spam flags** when proposals are edited to be compliant so that freelancers can self-correct.

#### Flow
1. **ScanForSpamCommand**(proposal_id) → AnalyzeText() | CheckLinks() | DetectPatterns() | ComputeScore() → **Outbox:** spam.detected.v1 (system-triggered)
2. **FlagAsSpamCommand**(proposal_id, spam_score, reasons[]) → MarkSpam() | NotifyModerators() → **Outbox:** spam.flagged.v1 (system-triggered)
3. **ReviewSpamCommand**(proposal_id, moderator_id, decision) → ApplyDecision() → **Outbox:** spam.reviewed.v1
4. **ClearSpamFlagCommand**(proposal_id) → RemoveFlag() → **Outbox:** spam.cleared.v1
5. **GetSpamStatusQuery**(proposal_id) → AuthorizeModerator() | Fetch() → SpamStatusDTO

#### Projections
- spam_detection_read

#### Events Published
- spam.detected.v1
- spam.flagged.v1
- spam.reviewed.v1
- spam.cleared.v1

#### Events Consumed
- proposal.submitted.v1 (to trigger spam scan)
- proposal.updated.v1 (to re-scan after edits)

#### RBAC/SLO
- **RBAC:** SYSTEM (scan/flag), MODERATOR (review/clear), FREELANCER/OWNER (view status)
- **SLO:** Scan P95 < 180ms

---

### 9.2 flag/

#### User Stories
- As a **user**, I want to **flag proposals** for inappropriate content so that platform safety is maintained.
- As a **moderator**, I want to **review flags** with reason codes so that action is informed.
- As a **moderator**, I want to **resolve or dismiss flags** so that cases are closed.
- As a **system**, I want to **prevent self-flagging** so that abuse is prevented.
- As a **freelancer**, I want to **see if my proposal is flagged** so that I can respond.

#### Flow
1. **FlagProposalCommand**(proposal_id, flagger_id, reason, details) → ValidateNotSelfFlag() | ValidateReason() | CreateFlag() → **Outbox:** proposal.flagged.v1
2. **ReviewFlagCommand**(flag_id, moderator_id, decision, notes) → ValidateStateTransition() | ApplyDecision() → **Outbox:** flag.reviewed.v1
3. **ResolveFlagCommand**(flag_id, moderator_id, action_taken) → Resolve() → **Outbox:** flag.resolved.v1
4. **DismissFlagCommand**(flag_id, moderator_id, reason) → Dismiss() → **Outbox:** flag.dismissed.v1
5. **GetProposalFlagsQuery**(proposal_id) → AuthorizeModerator() | Fetch() → FlagListDTO
6. **GetMyFlaggedProposalsQuery**(freelancer_id) → AuthorizeOwner() | Fetch() → FlaggedProposalsDTO

#### Projections
- proposal_flags_read
- flag_review_history_read

#### Events Published
- proposal.flagged.v1
- flag.reviewed.v1
- flag.resolved.v1
- flag.dismissed.v1

#### RBAC/SLO
- **RBAC:** ANY_USER (flag), MODERATOR (review/resolve/dismiss), FREELANCER/OWNER (view own)
- **SLO:** P95 < 150ms (flag), P95 < 200ms (review/resolve/dismiss)

---

### 9.3 compliance/

#### User Stories
- As a **system**, I want to **run compliance checks** (content policy, TOS, legal review, data privacy) so that regulatory requirements are met.
- As a **compliance officer**, I want to **review compliance violations** so that enforcement is consistent.
- As a **system**, I want to **delegate heavy compliance checks to compliance-be** so that proposals-be stays performant.
- As a **freelancer**, I want to **see compliance status** so that I know if my proposal meets requirements.

#### Flow
1. **RequestComplianceCheckCommand**(proposal_id, check_types[]) → DelegateToComplianceBe() | CreateCheckRecord() → **Outbox:** compliance.check.requested.v1
2. **ProcessComplianceResultCommand**(proposal_id, check_results[]) → StoreResults() | FlagViolations() → **Outbox:** compliance.check.completed.v1 (triggered by compliance-be callback)
3. **FlagViolationCommand**(proposal_id, violation_type, evidence) → MarkViolation() | NotifyCompliance() → **Outbox:** violation.detected.v1
4. **GetComplianceStatusQuery**(proposal_id) → Fetch() → ComplianceStatusDTO

#### Projections
- proposal_compliance_read
- compliance_violations_read

#### Events Published
- compliance.check.requested.v1
- compliance.check.completed.v1
- violation.detected.v1

#### Events Consumed
- proposal.submitted.v1 (to trigger compliance checks)
- compliance.result.ready.v1 (from compliance-be)

#### RBAC/SLO
- **RBAC:** SYSTEM (request/process checks), COMPLIANCE_OFFICER (view violations), FREELANCER/OWNER (view status)
- **SLO:** Request P95 < 120ms (async processing), Callback processing P95 < 180ms

---

## **10 - CLIENT INTERACTION DOMAIN**

### 10.1 interview/

#### User Stories
- As a **client**, I want to **request interviews** with freelancers so that I can assess fit before hiring.
- As a **freelancer**, I want to **accept or decline interview requests** so that I control my time.
- As a **client**, I want to **schedule interviews** with calendar integration so that coordination is easy.
- As a **client**, I want to **track interview status** (requested, accepted, scheduled, completed) so that workflow is clear.
- As a **freelancer**, I want to **reschedule interviews** so that conflicts can be managed.

#### Flow
1. **RequestInterviewCommand**(proposal_id, client_id, interview_type, proposed_times[]) → ValidateProposalState() | CreateRequest() → **Outbox:** interview.requested.v1
2. **AcceptInterviewCommand**(interview_id, freelancer_id, selected_time) → ValidateOwner() | Accept() | ScheduleInterview() → **Outbox:** interview.accepted.v1
3. **DeclineInterviewCommand**(interview_id, freelancer_id, reason) → ValidateOwner() | Decline() → **Outbox:** interview.declined.v1
4. **RescheduleInterviewCommand**(interview_id, new_time, requester_id) → ValidateParticipant() | Reschedule() → **Outbox:** interview.rescheduled.v1
5. **CompleteInterviewCommand**(interview_id, client_id, notes) → MarkComplete() → **Outbox:** interview.completed.v1
6. **GetInterviewQuery**(interview_id) → AuthorizeParticipant() | Fetch() → InterviewDTO
7. **ListProposalInterviewsQuery**(proposal_id) → AuthorizeParticipant() | Fetch() → InterviewListDTO

#### Projections
- interview_read
- interview_schedule_read

#### Events Published
- interview.requested.v1
- interview.accepted.v1
- interview.declined.v1
- interview.rescheduled.v1
- interview.completed.v1

#### RBAC/SLO
- **RBAC:** CLIENT (request/complete), FREELANCER/OWNER (accept/decline/reschedule), PARTICIPANTS (view)
- **SLO:** P95 < 170ms

---

### 10.2 feedback/

#### User Stories
- As a **client**, I want to **provide proposal feedback** so that freelancers can improve.
- As a **freelancer**, I want to **receive feedback** on proposals so that I can iterate.
- As a **system**, I want to **aggregate feedback** so that insights are surfaced.
- As a **freelancer**, I want to **respond to feedback** so that I can clarify or thank clients.

#### Flow
1. **SubmitProposalFeedbackCommand**(proposal_id, client_id, rating, feedback_text, categories[]) → ValidateClient() | StoreFeedback() → **Outbox:** proposal.feedback.submitted.v1
2. **RespondToFeedbackCommand**(feedback_id, freelancer_id, response_text) → ValidateOwner() | StoreResponse() → **Outbox:** feedback.response.submitted.v1
3. **GetProposalFeedbackQuery**(proposal_id) → AuthorizeParticipants() | Fetch() → FeedbackDTO
4. **GetMyFeedbackQuery**(freelancer_id) → AuthorizeOwner() | Fetch() → FeedbackListDTO

#### Projections
- proposal_feedback_read
- feedback_aggregates_read

#### Events Published
- proposal.feedback.submitted.v1
- feedback.response.submitted.v1

#### RBAC/SLO
- **RBAC:** CLIENT (submit feedback), FREELANCER/OWNER (respond/view), PARTICIPANTS (view)
- **SLO:** P95 < 160ms

---

### 10.3 shortlist/

#### User Stories
- As a **client**, I want to **add proposals to shortlist** so that I can organize top candidates.
- As a **client**, I want to **remove proposals from shortlist** so that I can narrow down choices.
- As a **client**, I want to **reorder shortlist** so that priority is clear.
- As a **client**, I want to **add notes to shortlisted proposals** so that evaluation criteria are documented.
- As a **system**, I want to **limit shortlist size** so that focus is maintained.

#### Flow
1. **AddToShortlistCommand**(job_id, proposal_id, client_id, notes) → ValidateShortlistLimit() | Add() → **Outbox:** proposal.shortlisted.v1
2. **RemoveFromShortlistCommand**(job_id, proposal_id, client_id) → Remove() → **Outbox:** proposal.removed_from_shortlist.v1
3. **ReorderShortlistCommand**(job_id, client_id, new_order[]) → ValidatePermutation() | ApplyOrder() → **Outbox:** shortlist.reordered.v1
4. **UpdateShortlistNotesCommand**(job_id, proposal_id, client_id, notes) → UpdateNotes() → **Outbox:** shortlist.notes.updated.v1
5. **GetShortlistQuery**(job_id, client_id) → AuthorizeClient() | Fetch() → ShortlistDTO

#### Projections
- job_shortlist_read

#### Events Published
- proposal.shortlisted.v1
- proposal.removed_from_shortlist.v1
- shortlist.reordered.v1
- shortlist.notes.updated.v1

#### RBAC/SLO
- **RBAC:** CLIENT/HIRING_TEAM (add/remove/reorder/update notes/view)
- **SLO:** P95 < 150ms, Max 15 proposals per shortlist

---

### 10.4 conversation/

#### User Stories
- As a **client/freelancer**, I want **pre-hire messaging** so that scope is clarified before commitment.
- As a **system**, I want **sentiment analysis** and response-time metrics so that engagement health is visible.
- As a **system**, I want **link/PII moderation** and rate limits so that safety is preserved.
- As a **client/freelancer**, I want to **see message read status** so that I know if communication was received.
- As a **system**, I want to **archive conversations** when proposals are closed so that history is preserved.

#### Flow
1. **SendMessageCommand**(proposal_id, sender_id, message_text) → RateLimit(user) | AntiPiiLint() | ModerateLinks() | AppendMessage() → **Outbox:** proposal.conversation.message.sent.v1
2. **MarkMessageReadCommand**(message_id, reader_id) → UpdateReadStatus() → **Outbox:** message.read.v1
3. **AnalyzeSentimentJob**() → ScoreMessageBatch() → **Outbox:** proposal.conversation.sentiment.scored.v1 (system-triggered)
4. **TrackResponseTimeCommand**(proposal_id, actor_id) → UpdateSLA() → **Outbox:** proposal.conversation.response_time.tracked.v1
5. **GetConversationQuery**(proposal_id) → AuthorizeParticipants() | Fetch() → ConversationDTO
6. **GetSentimentAnalysisQuery**(proposal_id) → AuthorizeParticipants() | Fetch() → SentimentDTO

#### Projections
- proposal_conversation_read
- proposal_sentiment_read
- proposal_response_time_read

#### Events Published
- proposal.conversation.message.sent.v1
- message.read.v1
- proposal.conversation.sentiment.scored.v1
- proposal.conversation.response_time.tracked.v1

#### RBAC/SLO
- **RBAC:** PARTICIPANTS_ONLY (send/read/view), SYSTEM (analyze/track)
- **SLO:** Send P95 < 130ms, 20 messages/hr per party, Idempotent by (thread_id, content_hash, ts_bucket)

---

## **11 - WORKFLOW & COLLABORATION DOMAIN**

**11.1 negotiation/ (COMPLETE)**
--------------------------------

#### User Stories

*   As a **client/freelancer**, I want to **negotiate terms** (price, timeline, scope) before accepting so that agreements are mutual.
    
*   As a **system**, I want to **track negotiation rounds** so that history is preserved.
    
*   As a **client/freelancer**, I want to **make counter-offers** so that terms can be adjusted iteratively.
    
*   As a **system**, I want to **enforce negotiation limits** (max 10 rounds) so that process doesn't drag indefinitely.
    
*   As a **client/freelancer**, I want to **accept negotiated terms** so that final agreement is clear.
    
*   As a **client/freelancer**, I want to **view negotiation history** so that I can track progress.
    
*   As a **system**, I want to **notify both parties** of counter-offers so that responsiveness is encouraged.
    

#### Flow

1.  **InitiateNegotiationCommand**(proposal\_id, initiator\_id, proposed\_changes) → ValidateProposalState() | CreateNegotiationThread() → **Outbox:** negotiation.initiated.v1
    
2.  **MakeCounterOfferCommand**(negotiation\_id, party\_id, counter\_terms) → ValidateRoundLimit() | AppendCounterOffer() | Notify(communications-be) → **Outbox:** negotiation.counter\_offer.made.v1
    
3.  **AcceptNegotiatedTermsCommand**(negotiation\_id, party\_id) → ValidateBothPartiesAgree() | FinalizeTerms() | UpdateProposal() → **Outbox:** negotiation.terms.accepted.v1
    
4.  **RejectNegotiationCommand**(negotiation\_id, party\_id, reason) → CloseNegotiation() → **Outbox:** negotiation.rejected.v1
    
5.  **GetNegotiationThreadQuery**(negotiation\_id) → AuthorizeParticipants() | Fetch() → NegotiationThreadDTO
    
6.  **ListProposalNegotiationsQuery**(proposal\_id) → AuthorizeParticipants() | Fetch() → NegotiationListDTO
    
7.  **GetNegotiationHistoryQuery**(negotiation\_id) → AuthorizeParticipants() | FetchFullHistory() → NegotiationHistoryDTO
    

#### Projections

*   negotiation\_read
    
*   negotiation\_history\_read
    
*   negotiation\_rounds\_read
    

#### Events Published

*   negotiation.initiated.v1
    
*   negotiation.counter\_offer.made.v1
    
*   negotiation.terms.accepted.v1
    
*   negotiation.rejected.v1
    

#### Events Consumed

*   proposal.submitted.v1 (to enable negotiation)
    

#### RBAC/SLO

*   **RBAC:** PARTICIPANTS (client/freelancer - initiate/counter/accept/reject/view)
    
*   **SLO:** P95 < 180ms, Max 10 negotiation rounds per proposal, Idempotent by (negotiation\_id, party\_id, terms\_hash)
---

### 11.2 invite/

#### User Stories
- As a **client**, I want to **send job invitations** to specific freelancers so that I can target talent.
- As a **freelancer**, I want to **receive invitations** with job details so that I can quickly assess fit.
- As a **freelancer**, I want to **accept or decline invitations** so that I control my workload.
- As a **system**, I want to **track invitation status** (sent, accepted, declined, expired) so that outcomes are clear.
- As a **system**, I want to **auto-expire invitations** after a time period so that old invitations are cleaned up.
- As a **system**, I want to **sync invitation data to proposals** when accepted so that proposal creation is seeded.

#### Flow
1. **SendJobInvitationCommand**(job_id, freelancer_id, client_id, message, target_rate) → ValidateJob() | CreateInvitation() | Notify(communications-be) → **Outbox:** job_invitation.sent.v1
2. **AcceptInvitationCommand**(invitation_id, freelancer_id) → ValidateOwner() | MarkAccepted() | SeedProposal() → **Outbox:** invitation.accepted.v1
3. **DeclineInvitationCommand**(invitation_id, freelancer_id, reason) → ValidateOwner() | MarkDeclined() → **Outbox:** invitation.declined.v1
4. **ExpireInvitationsJob**() → FindExpired() | MarkExpired() → **Outbox:** invitation.expired.v1 (system-triggered)
5. **GetInvitationQuery**(invitation_id) → AuthorizeOwner() | Fetch() → InvitationDTO
6. **GetMyInvitationsQuery**(freelancer_id) → AuthorizeOwner() | Fetch() → InvitationListDTO

#### Projections
- invitation_read
- invitation_by_freelancer_read
- invitation_by_job_read

#### Events Published
- job_invitation.sent.v1
- invitation.accepted.v1
- invitation.declined.v1
- invitation.expired.v1

#### Events Consumed
- job.invitation.sent.v1 (from jobs-be - to create invitation record)

#### RBAC/SLO
- **RBAC:** CLIENT (send), FREELANCER/OWNER (accept/decline/view), SYSTEM (expire)
- **SLO:** P95 < 150ms

---

### 11.3 revision/

#### User Stories
- As a **freelancer**, I want to **create proposal revisions** so that I can improve proposals iteratively.
- As a **system**, I want to **track revision history** so that changes are auditable.
- As a **freelancer**, I want to **compare revisions** so that I can see what changed.
- As a **freelancer**, I want to **restore previous versions** so that I can revert changes.
- As a **system**, I want to **limit revision count** so that storage is managed.

#### Flow
1. **CreateRevisionCommand**(proposal_id, changes, revision_notes) → ValidateRevisionLimit(max=50) | CreateSnapshot() | IncrementVersion() → **Outbox:** proposal.revision.created.v1
2. **RestoreRevisionCommand**(proposal_id, revision_id) → ValidateRevisionExists() | RestoreSnapshot() → **Outbox:** proposal.revision.restored.v1
3. **GetRevisionHistoryQuery**(proposal_id) → AuthorizeOwner() | Fetch() → RevisionHistoryDTO
4. **CompareRevisionsQuery**(revision_id_1, revision_id_2) → AuthorizeOwner() | GenerateDiff() → RevisionDiffDTO

#### Projections
- proposal_revisions_read

#### Events Published
- proposal.revision.created.v1
- proposal.revision.restored.v1

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (create/restore/view/compare)
- **SLO:** P95 < 190ms, Max 50 revisions per proposal

---

### 11.4 collaboration/ (CONSOLIDATED: team + collaboration_network)

#### User Stories
- As a **team lead**, I want to **create team proposals** so that multiple freelancers can collaborate on one job.
- As a **team lead**, I want **role-based permissions** (edit/comment/view) so that access is controlled.
- As a **team member**, I want **internal-only notes** separate from client view so that we coordinate privately.
- As a **system**, I want **activity analytics per member** so that contribution tracking is possible.
- As a **team lead**, I want **handoff/ownership change flows** so that continuity is maintained.
- As a **team member**, I want to **invite collaborators** to proposals so that expertise can be added.
- As a **system**, I want to **split revenue/connects** based on team configuration so that economics are fair.

#### Flow
1. **CreateTeamProposalCommand**(job_id, lead_id, team_members[], roles[]) → Validate() | CreateProposal(team_mode) → **Outbox:** team.proposal.created.v1
2. **AddTeamMemberCommand**(proposal_id, user_id, role, added_by) → ValidateLead() | AddMember() → **Outbox:** team.member.added.v1
3. **ChangeTeamRoleCommand**(proposal_id, user_id, new_role, changed_by) → ValidateLead() | UpdateRole() → **Outbox:** team.role.changed.v1
4. **RemoveTeamMemberCommand**(proposal_id, user_id, removed_by) → ValidateLead() | RemoveMember() | RecalcRevenueShare() → **Outbox:** team.member.removed.v1
5. **TransferOwnershipCommand**(proposal_id, new_lead_id, current_lead_id) → ValidateOwner() | TransferLead() → **Outbox:** team.ownership.transferred.v1
6. **AddInternalNoteCommand**(proposal_id, author_id, note_text) → ValidateTeamMember() | StoreNote() → **Outbox:** team.note.added.v1
7. **ConfigureRevenueSplitCommand**(proposal_id, split_config[]) → ValidateTotal100() | Configure() → **Outbox:** team.revenue.split.configured.v1
8. **GetTeamActivityQuery**(proposal_id) → AuthorizeLead() | AggregateActivity() → TeamActivityDTO
9. **GetTeamProposalQuery**(proposal_id) → AuthorizeTeamMember() | Fetch() → TeamProposalDTO

#### Projections
- team_proposals_read
- team_members_read
- team_activity_read
- team_notes_read
- revenue_split_read

#### Events Published
- team.proposal.created.v1
- team.member.added.v1
- team.role.changed.v1
- team.member.removed.v1
- team.ownership.transferred.v1
- team.note.added.v1
- team.revenue.split.configured.v1

#### RBAC/SLO
- **RBAC:** TEAM_LEAD (all operations), TEAM_MEMBER (view/add notes), SYSTEM (track activity)
- **SLO:** P95 < 200ms

---

## **12 - LIFECYCLE MANAGEMENT DOMAIN**

### 12.1 expiration/

#### User Stories
- As a **freelancer**, I want to **set proposal expiration dates** so that time-limited offers are clear.
- As a **freelancer**, I want to **extend expiration dates** so that I can keep proposals active longer.
- As a **system**, I want to **notify of upcoming expiry** so that freelancers can take action.
- As a **system**, I want to **auto-expire proposals** so that old proposals are cleaned up.
- As a **client**, I want to **see expiration dates** so that I know deadlines for accepting.

#### Flow
1. **SetProposalExpiryCommand**(proposal_id, expires_at) → ValidateFutureDate() | SetExpiry() → **Outbox:** proposal.expiry.set.v1
2. **ExtendProposalExpiryCommand**(proposal_id, new_expires_at) → ValidateExtension() | Extend() → **Outbox:** proposal.expiry.extended.v1
3. **NotifyUpcomingExpiryJob**() → FindExpiringSoon() | SendNotifications(communications-be) → **Outbox:** proposal.expiry.notified.v1 (system-triggered)
4. **ExpireProposalsJob**() → FindExpired() | MarkExpired() → **Outbox:** proposal.expired.v1 (system-triggered)
5. **GetProposalExpiryQuery**(proposal_id) → Fetch() → ExpiryDTO

#### Projections
- proposal_expiry_read

#### Events Published
- proposal.expiry.set.v1
- proposal.expiry.extended.v1
- proposal.expiry.notified.v1
- proposal.expired.v1

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (set/extend), SYSTEM (notify/expire), CLIENT/FREELANCER (view)
- **SLO:** P95 < 150ms

---

### 12.2 withdrawal/

#### User Stories
- As a **freelancer**, I want **soft-withdraw with a courteous message template** so that I maintain goodwill.
- As a **system**, I want **reversible withdraw within a window** (e.g., 24 hours) so that freelancers can re-enter if circumstances change.
- As a **product team**, I want **analytics on withdraw reasons** so that we can improve matching.
- As a **freelancer**, I want to **auto-notify client when withdrawing** so that communication is clear.
- As a **system**, I want to **refund connects** based on withdrawal timing so that fairness is maintained.

#### Flow
1. **WithdrawProposalCommand**(proposal_id, reason, message_to_client) → GuardWithdrawState() | MarkWithdrawn() | MaybeRefundConnects() | NotifyClient(communications-be) → **Outbox:** proposal.withdrawn.v1
2. **UpdateWithdrawalReasonCommand**(proposal_id, new_reason) → AuthorizeOwner() | UpdateReason() → **Outbox:** withdrawal.reason.updated.v1
3. **ReactivateWithdrawnProposalCommand**(proposal_id) → ValidateWithdrawWindow(24h) | Reactivate() → **Outbox:** proposal.reactivated.v1
4. **GetWithdrawalAnalyticsQuery**(filters) → AuthorizeProduct() | AggregateReasons() → WithdrawalAnalyticsDTO

#### Projections
- proposal_withdrawal_read
- withdrawal_analytics_read

#### Events Published
- proposal.withdrawn.v1
- withdrawal.reason.updated.v1
- proposal.reactivated.v1

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (withdraw/update reason/reactivate), PRODUCT_TEAM (view analytics)
- **SLO:** P95 < 140ms

---

### 12.3 archive/

#### User Stories
- As a **freelancer**, I want **manual archiving** of old proposals so that my workspace stays clean.
- As a **freelancer**, I want to **search and filter across archived proposals** so that I can still learn from history.
- As a **system**, I want **auto-archive rules** (e.g., stale 90 days post-close) so that maintenance is automated.
- As a **compliance team**, I want **retention windows and export** so that deletion follows policy.
- As a **freelancer**, I want to **restore archived proposals** so that I can reference them again.

#### Flow
1. **ArchiveProposalCommand**(proposal_id, reason) → AuthorizeOwner() | Archive() → **Outbox:** proposal.archived.v1
2. **RestoreProposalCommand**(proposal_id) → AuthorizeOwner() | Restore() → **Outbox:** proposal.restored.v1
3. **AutoArchiveStaleProposalsJob**() → FindStaleProposals(90d_closed) | Archive() → **Outbox:** proposal.auto_archived.v1 (system-triggered)
4. **ExportArchivedProposalsCommand**(freelancer_id, date_range) → AuthorizeOwner() | GenerateExport() → ExportFileDTO
5. **SearchArchivedProposalsQuery**(freelancer_id, search_criteria) → AuthorizeOwner() | Search() → ArchivedProposalListDTO

#### Projections
- proposal_archive_read

#### Events Published
- proposal.archived.v1
- proposal.restored.v1
- proposal.auto_archived.v1

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (archive/restore/search/export), SYSTEM (auto-archive)
- **SLO:** P95 < 140ms (manual operations), Export async

---

### 12.4 pipeline/

#### User Stories
- As a **freelancer**, I want a **Kanban-style pipeline** (Draft → Submitted → Viewed → Interview → Offer → Hired/Declined) with **per-stage WIP limits** so that I can prioritize effort.
- As a **freelancer**, I want **bulk stage moves** (e.g., multi-select to Archive/Decline) so that I can tidy my funnel quickly.
- As a **system**, I want **auto-moves on signals** (client viewed, interview scheduled, offer sent) so that pipeline stays accurate.
- As a **manager** (team account), I want **aggregate pipeline and per-member views** so that I can coach performance.
- As a **freelancer**, I want **stage-level analytics** (conversion rates, time-in-stage) so that I can optimize process.

#### Flow
1. **MoveProposalToStageCommand**(proposal_id, target_stage, actor_id) → ValidateStageTransition() | CheckWIPLimit() | Move() | TrackStageTime() → **Outbox:** proposal.pipeline.moved.v1
2. **BulkMoveProposalsCommand**(proposal_ids[], target_stage, actor_id) → ValidateBatch() | MoveBatch() → **Outbox:** Multiple proposal.pipeline.moved.v1 + proposal.pipeline.bulk_moved.summary.v1
3. **AutoMovePipelineCommand**(proposal_id, signal_type) → DetermineTargetStage() | Move() → **Outbox:** proposal.pipeline.auto_moved.v1 (system-triggered)
4. **GetPipelineViewQuery**(freelancer_id) → AuthorizeOwner() | FetchAllStages() → PipelineViewDTO
5. **GetTeamPipelineQuery**(team_id) → AuthorizeManager() | AggregateTeamPipeline() → TeamPipelineDTO
6. **GetPipelineAnalyticsQuery**(freelancer_id, time_period) → AuthorizeOwner() | ComputeMetrics() → PipelineAnalyticsDTO

#### Projections
- proposal_pipeline_read
- pipeline_stage_times_read
- pipeline_analytics_read

#### Events Published
- proposal.pipeline.moved.v1
- proposal.pipeline.bulk_moved.summary.v1
- proposal.pipeline.auto_moved.v1

#### Events Consumed
- proposal.viewed.by_client.v1 (to auto-move to Viewed)
- interview.scheduled.v1 (to auto-move to Interview)
- proposal.accepted.v1 (to auto-move to Hired)

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (move/bulk move/view), MANAGER (view team), SYSTEM (auto-move)
- **SLO:** P95 < 150ms (single move), P95 < 500ms (bulk move ≤50)

---

### 12.5 recycling/

#### User Stories
- As a **freelancer**, I want to **recycle or adapt past proposals** to similar jobs so that I can submit high-quality drafts faster.
- As a **system**, I want **similarity matching and PII redaction** so that reuse is relevant and safe.
- As a **freelancer**, I want a **change log of adaptations** so that I can track what was altered.
- As a **manager**, I want **freshness and policy gates** so that outdated content isn't reused without review.
- As a **freelancer**, I want to **preview recycled proposal** before finalizing so that quality is ensured.

#### Flow
1. **RecycleProposalCommand**(original_proposal_id, new_job_id, keep_sections[]) → FindSimilarJobSections() | RedactPII() | CreateDraftVersion() → **Outbox:** proposal.recycled.v1
2. **AdaptRecycledProposalCommand**(proposal_id, deltas) → ApplyDeltas() | AppendVersion() → **Outbox:** proposal.adapted.v1
3. **PreviewRecycledProposalQuery**(original_proposal_id, new_job_id) → GeneratePreview() → RecycledProposalPreviewDTO
4. **GetRecyclingHistoryQuery**(proposal_id) → AuthorizeOwner() | Fetch() → RecyclingHistoryDTO

#### Projections
- proposal_recycling_read
- proposal_versions_read

#### Events Published
- proposal.recycled.v1
- proposal.adapted.v1
- proposal.version.created.v1

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (recycle/adapt/preview/view), TEAM_LEAD (policy override)
- **SLO:** Recycle P95 < 220ms, Idempotent by (original_proposal_id, new_job_id, keep_sections_hash)

---

## **13 - RECOMMENDATION & VISIBILITY DOMAIN**

### 13.1 recommendation/

#### User Stories
- As a **freelancer**, I want **recommended jobs** that my proposal would rank well on so that I invest time wisely.
- As a **system**, I want **click/convert feedback loops** so that models improve continuously.
- As a **freelancer**, I want **recommendation explanations** so that I understand why jobs are suggested.
- As a **system**, I want to **personalize recommendations** based on success history so that suggestions are relevant.

#### Flow
1. **GenerateJobRecommendationsCommand**(freelancer_id) → AnalyzeProfile() | ScoreJobs() | RankByFit() → **Outbox:** recommendations.generated.v1 (system-triggered)
2. **TrackRecommendationClickCommand**(recommendation_id, freelancer_id) → RecordClick() | UpdateModel() → **Outbox:** recommendation.clicked.v1
3. **TrackRecommendationConversionCommand**(recommendation_id, proposal_id) → RecordConversion() | UpdateModel() → **Outbox:** recommendation.converted.v1
4. **GetJobRecommendationsQuery**(freelancer_id) → AuthorizeOwner() | Fetch() → RecommendationListDTO

#### Projections
- job_recommendations_read
- recommendation_performance_read

#### Events Published
- recommendations.generated.v1
- recommendation.clicked.v1
- recommendation.converted.v1

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (view), SYSTEM (generate/track)
- **SLO:** Generate P95 < 400ms, Track P95 < 100ms

---

### 13.2 context/

#### User Stories
- As a **freelancer**, I want a **job/client/market context pack** so that I can tailor proposals effectively.
- As a **system**, I want **periodic refreshes on job-signal changes** so that context stays current.
- As a **freelancer**, I want **competitive insights** (average bid, proposal count) so that strategy is informed.
- As a **freelancer**, I want **client history insights** (hiring patterns, budget range) so that targeting is optimized.

#### Flow
1. **GenerateContextCommand**(proposal_id) → AnalyzeJob() | AnalyzeClient() | AnalyzeMarket() | Store() → **Outbox:** proposal.context.generated.v1
2. **RefreshContextJob**() → UpdateContextForActiveProposals() → **Outbox:** proposal.context.updated.v1 (system-triggered)
3. **GetProposalContextQuery**(proposal_id) → AuthorizeOwner() | Fetch() → ProposalContextDTO

#### Projections
- proposal_context_read

#### Events Published
- proposal.context.generated.v1
- proposal.context.updated.v1

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (view), SYSTEM (generate/refresh)
- **SLO:** Generate P95 < 500ms

---

### 13.3 urgency/

#### User Stories
- As a **freelancer**, I want an **urgency score and recommended action** so that I act at the right time.
- As a **system**, I want **surge notifications** when competition spikes so that freelancers respond quickly.
- As a **freelancer**, I want **deadline reminders** so that I don't miss submission windows.

#### Flow
1. **CalculateUrgencyCommand**(proposal_id) → AnalyzeDeadline() | AnalyzeCompetition() | ComputeScore() → **Outbox:** proposal.urgency.set.v1
2. **TriggerUrgencyAlertJob**(job_id) → DetectSurge() | NotifyRelevantFreelancers(communications-be) → **Outbox:** proposal.urgency.alert.triggered.v1 (system-triggered)
3. **GetUrgencyScoreQuery**(proposal_id) → AuthorizeOwner() | Fetch() → UrgencyScoreDTO

#### Projections
- proposal_urgency_read

#### Events Published
- proposal.urgency.set.v1
- proposal.urgency.alert.triggered.v1

#### RBAC/SLO
- **RBAC:** FREELANCER/OWNER (view), SYSTEM (calculate/trigger alerts)
- **SLO:** Calculate P95 < 150ms

---

## **14 - RISK & COMPLIANCE DOMAIN**

### 14.1 risk_assessment/

#### User Stories
- As a **trust analyst**, I want **risk scores** (unrealistic timeline, too-cheap bids, spam language) so that I can intervene early.
- As a **system**, I want **mitigation flows** (warnings, gating, require-verify) so that platform risk is reduced.
- As a **freelancer**, I want **remediation guidance** so that I can fix flagged issues.
- As a **system**, I want to **track risk score changes over time** so that improvement is measurable.

#### Flow
1. **AssessProposalRiskCommand**(proposal_id) → AnalyzeTimeline() | AnalyzePricing() | AnalyzeContent() | ComputeScore() → **Outbox:** proposal.risk.assessed.v1 (system-triggered)
2. **ApplyMitigationCommand**(proposal_id, mitigation_type) → ApplyGating() | SendWarning(communications-be) → **Outbox:** proposal.risk.mitigation.applied.v1
3. **RemediateRiskCommand**(proposal_id, freelancer_id, remediation_actions[]) → ValidateActions() | UpdateProposal() | Reassess() → **Outbox:** proposal.risk.remediated.v1
4. **GetRiskAssessmentQuery**(proposal_id) → AuthorizeTrust() | Fetch() → RiskAssessmentDTO
5. **GetRemediationGuidanceQuery**(proposal_id) → AuthorizeOwner() | GenerateGuidance() → RemediationGuidanceDTO

#### Projections
- proposal_risk_read
- risk_mitigation_read

#### Events Published
- proposal.risk.assessed.v1
- proposal.risk.mitigation.applied.v1
- proposal.risk.remediated.v1

#### Events Consumed
- proposal.submitted.v1 (to trigger risk assessment)

#### RBAC/SLO
- **RBAC:** SYSTEM (assess/apply mitigation), TRUST_ANALYST (view/review), FREELANCER/OWNER (remediate/view guidance)
- **SLO:** Assess P95 < 250ms

---

## **GLOBAL CONVENTIONS & CROSS-CUTTING CONCERNS**

### Event Envelope (All Domains)
```json
{
  "event_id": "uuid",
  "event_ts": "ISO8601",
  "aggregate_id": "proposal_id",
  "partition_key": "proposal_id",
  "correlation_id": "uuid",
  "causation_id": "uuid",
  "actor": {
    "id": "user_id",
    "role": "FREELANCER|CLIENT|SYSTEM|..."
  },
  "user_context": {
    "ip": "xxx.xxx.xxx.xxx",
    "user_agent": "...",
    "device_type": "web|mobile|api"
  },
  "data_zone": "EU|US",
  "schema_ref": "schema_version",
  "compliance_context": {
    "pii_flags": [],
    "anonymized": true|false
  },
  "data": { ... }
}
```

### Write-Path Defaults (All Domains)
- **Idempotency:** HTTP Header `Idempotency-Key` or envelope-level key; safe retries return 200 with no duplicate events
- **Transactions:** Database transaction + outbox pattern with (aggregate_id, event_type, idempotency_key) deduplication
- **Retries/DLQ:** For external service calls (jobs-be, subscriptions-be, financial-be, communications-be, contracts-be, users-be, storage-be, compliance-be); exponential backoff; poison messages to DLQ
- **Projections:** `_read` materialized views per domain; metric `event_to_projector_lag_ms` tracked
- **Security:** RBAC enforced on all commands/queries; field-level encryption for sensitive data; secrets never logged
- **Performance:** Typical write P95 ≤ 300ms, read P95 ≤ 250ms (unless specified otherwise)
- **Rate Limiting:** Per-endpoint rate limits enforced (e.g., 60 req/min/freelancer for proposal submissions)

### PII Handling (All Domains)
- **NO raw PII in events:** Emit hashes, storage_ids, references only
- **Examples:** No plaintext emails, phone numbers, file contents, file names
- **Compliance:** PII flags in envelope; anonymization tracked; data residency respected

### Integration Patterns
- **Async Event-Driven:** Primary integration via Kafka events
- **Sync Queries:** REST APIs for read operations with caching
- **Command Validation:** External service validation (jobs-be, subscriptions-be, financial-be) before persistence
- **Circuit Breakers:** Implemented for all external service calls
- **Saga Pattern:** Used for multi-service workflows (e.g., proposal acceptance → contract creation)

---

## **SUMMARY**

This document covers **all 14 feature domains** in proposals-be with complete user stories following the pattern:
- **As a [role], I want [capability] so that [benefit]**
- **Flow:** Commands → Queries with validation and integration points
- **Projections:** Read models for query optimization
- **Events:** Published and consumed with full envelope format
- **RBAC/SLO:** Role-based access control and service level objectives

All features align with:
- ✅ Event-driven architecture (Kafka + Outbox Pattern)
- ✅ CQRS with read projections
- ✅ Clean Architecture (DDD)
- ✅ Non-PII compliance
- ✅ Idempotent write-path
- ✅ Cross-service integration
- ✅ Large-scale Upwork-like platform requirements

**Domains Covered:**
1. Core Proposal (proposal, cover_letter, attachment, question_answer, milestone)
2. Bidding System (bid, bid_strategy, bid_notification, auction, bid_anomaly_detection)
3. Connects & Boost (connect, connect_refund, boost)
4. Templates & Rate Cards (template, rate_card)
5. Analytics & Tracking (performance - consolidated)
6. Similarity & Deduplication (similarity - consolidated)
7. Portfolio (portfolio - consolidated)
8. Engagement (engagement - consolidated)
9. Moderation & Compliance (spam_detection, flag, compliance)
10. Client Interaction (interview, feedback, shortlist, conversation)
11. Workflow & Collaboration (negotiation, invite, revision, collaboration - consolidated)
12. Lifecycle Management (expiration, withdrawal, archive, pipeline, recycling)
13. Recommendation & Visibility (recommendation, context, urgency)
14. Risk & Compliance (risk_assessment)#### Projections
- proposal_boost_read
- boost_ledger_read


