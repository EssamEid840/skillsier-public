# proposals-be User Stories - Refactored
## Skillsier Platform - Proposal Management Service

> **Last Updated:** 2025-10-19  
> **Service:** proposals-be  
> **Architecture:** Clean Architecture + DDD + Event-Driven + CQRS

---

## Table of Contents
1. [Core Proposal Management](#1---core-proposal-management)
2. [Lifecycle Management](#2---lifecycle-management)
3. [Bidding & Auctions](#3---bidding--auctions)
4. [Connects & Credits](#4---connects--credits)
5. [Workflow & Collaboration](#5---workflow--collaboration)
6. [Client Interaction](#6---client-interaction)
7. [Performance & Analytics](#7---performance--analytics)
8. [Similarity & Deduplication](#8---similarity--deduplication)
9. [Portfolio Management](#9---portfolio-management)
10. [Engagement & Follow-ups](#10---engagement--follow-ups)
11. [Moderation & Compliance](#11---moderation--compliance)
12. [Templates & Rate Cards](#12---templates--rate-cards)

---

## 1 - CORE PROPOSAL MANAGEMENT

### 1.1 proposal/

#### User Stories
- As a **freelancer**, I want to **create draft proposals** so that I can prepare multiple applications before submission.
- As a **freelancer**, I want to **submit proposals with required fields** so that clients can evaluate my offer completely.
- As a **freelancer**, I want to **update draft proposals** so that I can refine my application before submission.
- As a **freelancer**, I want to **view my proposal status** so that I know where I stand in the hiring process.
- As a **system**, I want **status transitions to be validated** so that proposals follow proper workflow.

#### Flow
1. **CreateProposalCommand**(job_id, freelancer_id, draft_data) → Validate() | Persist() → **Outbox:** proposal.created.v1
2. **UpdateProposalCommand**(proposal_id, updates) → AuthorizeOwner() | Validate() | Apply() → **Outbox:** proposal.updated.v1
3. **SubmitProposalCommand**(proposal_id) → ValidateComplete() | UseConnects() | Submit() → **Outbox:** proposal.submitted.v1
4. **GetProposalQuery**(proposal_id) → AuthorizeAccess() | Fetch() → ProposalDTO
5. **ListProposalsQuery**(filters) → ApplyFilters() | Paginate() → ProposalListDTO

#### Projections
- proposal_read
- proposal_stats_read

#### Events Published
- proposal.created.v1
- proposal.updated.v1
- proposal.submitted.v1
- proposal.accepted.v1
- proposal.rejected.v1

#### RBAC/SLO
- **RBAC:** OWNER (create/update), CLIENT (view/accept/reject), SYSTEM (status updates)
- **SLO:** P95 < 180ms

---

### 1.2 cover_letter/

#### User Stories
- As a **freelancer**, I want to **write a compelling cover letter** so that I can personalize my application.
- As a **freelancer**, I want **AI-assisted cover letter suggestions** so that I can improve quality.
- As a **system**, I want **tone and quality analysis** so that I can provide feedback.

#### Flow
1. **CreateCoverLetterCommand**(proposal_id, content) → ValidateLength() | AnalyzeTone() | Persist() → **Outbox:** cover_letter.created.v1
2. **UpdateCoverLetterCommand**(proposal_id, content) → Validate() | Update() → **Outbox:** cover_letter.updated.v1
3. **AnalyzeCoverLetterQuery**(proposal_id) → RunAnalysis() → AnalysisDTO (word_count, tone, quality_score)

#### Projections
- cover_letter_read
- cover_letter_analysis_read

#### Events Published
- cover_letter.created.v1
- cover_letter.updated.v1

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 220ms

---

### 1.3 attachment/

#### User Stories
- As a **freelancer**, I want to **attach portfolio samples** so that I can showcase relevant work.
- As a **freelancer**, I want to **manage multiple attachments** so that I can provide comprehensive evidence.
- As a **system**, I want **file type and size validation** so that uploads are safe and manageable.

#### Flow
1. **UploadAttachmentCommand**(proposal_id, file) → ValidateFile() | VirusScan() | Upload() | CreateRecord() → **Outbox:** attachment.added.v1
2. **DeleteAttachmentCommand**(attachment_id) → AuthorizeOwner() | Delete() → **Outbox:** attachment.removed.v1
3. **ListAttachmentsQuery**(proposal_id) → Fetch() → AttachmentListDTO

#### Projections
- attachment_read

#### Events Published
- attachment.added.v1
- attachment.removed.v1

#### Events Consumed
- storage.file.uploaded
- storage.file.deleted

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 300ms (upload), P95 < 100ms (list)

---

### 1.4 question_answer/

#### User Stories
- As a **client**, I want to **ask screening questions** so that I can filter candidates efficiently.
- As a **freelancer**, I want to **answer job questions** so that I can demonstrate my understanding.
- As a **system**, I want **required question validation** so that proposals are complete.

#### Flow
1. **AnswerQuestionCommand**(proposal_id, question_id, answer) → ValidateAnswer() | Persist() → **Outbox:** question.answered.v1
2. **UpdateAnswerCommand**(answer_id, new_answer) → Validate() | Update() → **Outbox:** answer.updated.v1
3. **GetAnswersQuery**(proposal_id) → Fetch() → QuestionAnswerListDTO

#### Projections
- question_answer_read

#### Events Published
- question.answered.v1
- answer.updated.v1

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 150ms

---

### 1.5 milestone/

#### User Stories
- As a **freelancer**, I want to **propose milestone-based projects** so that payment is tied to deliverables.
- As a **client**, I want to **see milestone breakdown** so that I can understand the delivery plan.
- As a **system**, I want **milestone amounts to sum to total** so that proposals are financially accurate.

#### Flow
1. **CreateMilestoneCommand**(proposal_id, description, amount, due_date) → ValidateAmountSplit() | Persist() → **Outbox:** milestone.created.v1
2. **UpdateMilestoneCommand**(milestone_id, updates) → Validate() | Update() → **Outbox:** milestone.updated.v1
3. **GetMilestonesQuery**(proposal_id) → Fetch() → MilestoneListDTO

#### Projections
- milestone_read

#### Events Published
- milestone.created.v1
- milestone.updated.v1
- milestone.completed.v1

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 160ms

---

## 2 - LIFECYCLE MANAGEMENT

### 2.1 expiration/

#### User Stories
- As a **freelancer**, I want **automatic proposal expiration** so that stale proposals don't remain active.
- As a **freelancer**, I want to **extend expiration** so that I can keep proposals active longer.
- As a **system**, I want **expiration notifications** so that freelancers can take action before expiry.

#### Flow
1. **SetExpirationCommand**(proposal_id, expires_at) → Persist() → **Outbox:** expiration.set.v1
2. **ExtendExpirationCommand**(proposal_id, new_date) → ValidateExtension() | Update() → **Outbox:** expiration.extended.v1
3. **ExpirationSweepJob**() → FindExpiring() | NotifyFreelancers() | MarkExpired() → **Outbox:** proposal.expiring.v1, proposal.expired.v1
4. **GetExpirationQuery**(proposal_id) → Fetch() → ExpirationDTO

#### Projections
- expiration_read

#### Events Published
- expiration.set.v1
- expiration.extended.v1
- proposal.expiring.v1 (24h before)
- proposal.expired.v1

#### RBAC/SLO
- **RBAC:** OWNER (extend), SYSTEM (sweep)
- **SLO:** P95 < 120ms, Sweep runs every hour

---

### 2.2 withdrawal/

#### User Stories
- As a **freelancer**, I want to **withdraw proposals** so that I can remove applications I'm no longer interested in.
- As a **freelancer**, I want to **specify withdrawal reason** so that the system can learn from my decisions.
- As a **system**, I want **connect refund eligibility** so that withdrawals are handled fairly.

#### Flow
1. **WithdrawProposalCommand**(proposal_id, reason) → ValidateCanWithdraw() | CheckRefundEligibility() | Withdraw() → **Outbox:** proposal.withdrawn.v1
2. **UpdateWithdrawalReasonCommand**(proposal_id, reason) → Update() → **Outbox:** withdrawal.reason.updated.v1
3. **GetWithdrawalHistoryQuery**(freelancer_id) → Fetch() → WithdrawalListDTO

#### Projections
- withdrawal_read
- withdrawal_stats_read

#### Events Published
- proposal.withdrawn.v1
- withdrawal.reason.updated.v1

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 150ms

---

### 2.3 archive/

#### User Stories
- As a **freelancer**, I want to **archive old proposals** so that my active proposals list stays clean.
- As a **freelancer**, I want to **restore archived proposals** so that I can reference past work.
- As a **system**, I want **automatic archival** so that old proposals don't clutter the database.

#### Flow
1. **ArchiveProposalCommand**(proposal_id) → ValidateArchivable() | Archive() → **Outbox:** proposal.archived.v1
2. **RestoreProposalCommand**(proposal_id) → ValidateRestorable() | Restore() → **Outbox:** proposal.restored.v1
3. **PurgeExpiredArchivesJob**() → FindExpired() | HardDelete() → **Outbox:** proposal.purged.v1
4. **ListArchivedProposalsQuery**(freelancer_id) → Fetch() → ArchivedProposalListDTO

#### Projections
- archive_read

#### Events Published
- proposal.archived.v1
- proposal.restored.v1
- proposal.purged.v1

#### RBAC/SLO
- **RBAC:** OWNER (archive/restore), SYSTEM (purge)
- **SLO:** P95 < 140ms

---

### 2.4 pipeline/

#### User Stories
- As a **freelancer**, I want to **track proposals through pipeline stages** so that I can manage my applications.
- As a **freelancer**, I want **pipeline analytics** so that I can see conversion rates at each stage.
- As a **system**, I want **valid stage transitions** so that the pipeline maintains integrity.

#### Flow
1. **MoveToStageCommand**(proposal_id, new_stage) → ValidateTransition() | Move() → **Outbox:** pipeline.stage.moved.v1
2. **GetPipelineQuery**(proposal_id) → Fetch() → PipelineDTO
3. **GetPipelineAnalyticsQuery**(freelancer_id) → Aggregate() → PipelineAnalyticsDTO (conversion rates, avg time per stage)
4. **ListProposalsByStageQuery**(stage, filters) → Fetch() → ProposalListDTO

#### Projections
- pipeline_read
- pipeline_analytics_read

#### Events Published
- pipeline.stage.moved.v1
- stage.changed.v1

#### Pipeline Stages
- Drafting → Submitted → UnderReview → Shortlisted → Interviewing → Negotiating → Won/Lost

#### RBAC/SLO
- **RBAC:** OWNER (view), SYSTEM (move stages)
- **SLO:** P95 < 160ms

---

### 2.5 recycling/

#### User Stories
- As a **freelancer**, I want to **reuse successful proposals** so that I can apply faster to similar jobs.
- As a **freelancer**, I want **version history** so that I can track changes over time.
- As a **system**, I want **modification tracking** so that recycling is auditable.

#### Flow
1. **RecycleProposalCommand**(original_proposal_id, new_job_id, modifications) → ValidateRecyclable() | CopyWithModifications() | CreateNew() → **Outbox:** proposal.recycled.v1
2. **GetVersionHistoryQuery**(proposal_id) → FetchAllVersions() → VersionHistoryDTO
3. **GetRecyclableProposalsQuery**(freelancer_id) → FindSuccessful() → ProposalListDTO

#### Projections
- recycling_read
- version_history_read

#### Events Published
- proposal.recycled.v1
- version.created.v1

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 200ms

---

## 3 - BIDDING & AUCTIONS

### 3.1 bid/

#### User Stories
- As a **freelancer**, I want to **place competitive bids** so that I can compete on price.
- As a **freelancer**, I want to **see my bid rank** so that I know where I stand.
- As a **freelancer**, I want **outbid notifications** so that I can adjust my bid if needed.
- As a **system**, I want **bid validation** so that bids follow marketplace rules.

#### Flow
1. **PlaceBidCommand**(proposal_id, amount, bid_type, currency) → ValidateBid() | CalculateRank() | Persist() → **Outbox:** bid.placed.v1
2. **UpdateBidCommand**(bid_id, new_amount) → ValidateUpdate() | RecalculateRank() | Update() → **Outbox:** bid.updated.v1
3. **WithdrawBidCommand**(bid_id) → Withdraw() → **Outbox:** bid.withdrawn.v1
4. **GetBidRankQuery**(bid_id) → CalculateCurrentRank() → BidRankDTO
5. **GetCompetitiveBidsQuery**(job_id) → FetchCompetition() → CompetitiveBidsDTO

#### Projections
- bid_read
- bid_ranking_read

#### Events Published
- bid.placed.v1
- bid.updated.v1
- bid.withdrawn.v1
- bid.accepted.v1
- bid.outbid.v1

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 180ms

---

### 3.2 bid_strategy/

#### User Stories
- As a **freelancer**, I want **auto-bidding strategies** so that I can bid competitively without manual monitoring.
- As a **freelancer**, I want **strategy templates** (competitive, premium, value-based) so that I can choose a bidding approach.
- As a **system**, I want **strategy optimization** so that bid success rates improve over time.

#### Flow
1. **CreateBidStrategyCommand**(freelancer_id, strategy_type, rules, default_markup) → Validate() | Persist() → **Outbox:** bid_strategy.created.v1
2. **UpdateBidStrategyCommand**(strategy_id, updates) → Validate() | Update() → **Outbox:** bid_strategy.updated.v1
3. **ApplyStrategyToProposalCommand**(proposal_id, strategy_id) → LoadStrategy() | CalculateBid() | Apply() → **Outbox:** bid_strategy.applied.v1
4. **GetRecommendedBidQuery**(job_id, strategy_id) → AnalyzeCompetition() | ApplyRules() → RecommendedBidDTO

#### Projections
- bid_strategy_read
- strategy_performance_read

#### Events Published
- bid_strategy.created.v1
- bid_strategy.updated.v1
- bid_strategy.applied.v1

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 240ms

---

### 3.3 auction/

#### User Stories
- As a **client**, I want **auction-style bidding** so that I can get the best price.
- As a **freelancer**, I want to **participate in live auctions** so that I can compete in real-time.
- As a **system**, I want **auction lifecycle management** so that auctions run fairly.

#### Flow
1. **StartAuctionCommand**(job_id, start_time, end_time, minimum_bid) → Schedule() → **Outbox:** auction.started.v1
2. **PlaceAuctionBidCommand**(auction_id, proposal_id, amount) → ValidateActive() | CheckMinimum() | UpdateHighBid() → **Outbox:** auction.new_high_bid.v1
3. **EndAuctionCommand**(auction_id) → DeclareWinner() | Close() → **Outbox:** auction.ended.v1
4. **GetAuctionStatusQuery**(auction_id) → FetchStatus() → AuctionStatusDTO

#### Projections
- auction_read
- auction_history_read

#### Events Published
- auction.started.v1
- auction.new_high_bid.v1
- auction.ended.v1
- auction.cancelled.v1

#### RBAC/SLO
- **RBAC:** CLIENT (start/end), OWNER (bid)
- **SLO:** P95 < 160ms

---

### 3.4 bid_anomaly/

#### User Stories
- As a **trust & safety team**, I want **anomaly detection** so that suspicious bidding patterns are flagged.
- As a **system**, I want **automated anomaly scoring** so that reviews are prioritized.
- As an **analyst**, I want **anomaly reports** so that I can investigate patterns.

#### Flow
1. **DetectAnomaliesJob**() → AnalyzeBids() | ScoreAnomalies() | Flag() → **Outbox:** bid_anomaly.detected.v1
2. **ReviewAnomalyCommand**(anomaly_id, action) → Investigate() | TakeAction() → **Outbox:** bid_anomaly.reviewed.v1
3. **ConfirmAnomalyCommand**(anomaly_id) → ConfirmSuspicious() | ApplyPenalty() → **Outbox:** bid_anomaly.confirmed.v1
4. **GetAnomalyReportQuery**(filters) → Aggregate() → AnomalyReportDTO

#### Projections
- bid_anomaly_read
- anomaly_patterns_read

#### Events Published
- bid_anomaly.detected.v1
- bid_anomaly.reviewed.v1
- bid_anomaly.confirmed.v1

#### Anomaly Types
- UnusuallyLow, UnusuallyHigh, RapidBidding, SuspiciousPattern

#### RBAC/SLO
- **RBAC:** SYSTEM (detect), TRUST_SAFETY (review)
- **SLO:** P95 < 300ms

---

## 4 - CONNECTS & CREDITS

### 4.1 connect/

#### User Stories
- As a **freelancer**, I want to **use connects to apply** so that I can access premium jobs.
- As a **freelancer**, I want to **see connect balance** so that I know when to purchase more.
- As a **system**, I want **tiered pricing** so that popular jobs cost more connects.

#### Flow
1. **UseConnectsCommand**(proposal_id, job_id) → CalculateCost() | ValidateBalance() | Deduct() → **Outbox:** connects.used.v1
2. **GetConnectBalanceQuery**(freelancer_id) → FetchBalance() → ConnectBalanceDTO
3. **GetConnectUsageHistoryQuery**(freelancer_id) → FetchHistory() → ConnectUsageListDTO
4. **CalculateConnectCostQuery**(job_id) → DeterminePopularity() | ApplyTier() → ConnectCostDTO

#### Projections
- connect_balance_read
- connect_usage_read

#### Events Published
- connects.used.v1
- connects.refunded.v1
- connects.expired.v1

#### Events Consumed
- subscriptions.connects.purchased
- subscriptions.connects.granted

#### Connect Tiers
- Tier 1: 1 connect (low popularity)
- Tier 2: 2 connects (medium popularity)
- Tier 3: 4 connects (high popularity)
- Tier 4: 6 connects (very high popularity)

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 100ms

---

### 4.2 connect_refund/

#### User Stories
- As a **freelancer**, I want **connect refunds** so that I'm not charged for invalid job posts.
- As a **system**, I want **refund eligibility rules** so that refunds are fair.
- As a **support team**, I want **refund approval workflow** so that requests are handled consistently.

#### Flow
1. **RequestRefundCommand**(proposal_id, reason) → ValidateEligibility() | CreateRequest() → **Outbox:** refund.requested.v1
2. **ApproveRefundCommand**(refund_id) → ProcessRefund() | RestoreConnects() → **Outbox:** refund.approved.v1
3. **DenyRefundCommand**(refund_id, reason) → Deny() → **Outbox:** refund.denied.v1
4. **GetRefundStatusQuery**(refund_id) → FetchStatus() → RefundStatusDTO

#### Projections
- connect_refund_read
- refund_stats_read

#### Events Published
- refund.requested.v1
- refund.approved.v1
- refund.denied.v1

#### Refund Reasons
- JobCancelled, ClientUnresponsive, TechnicalIssue, MisleadingJobPost

#### RBAC/SLO
- **RBAC:** OWNER (request), SUPPORT (approve/deny)
- **SLO:** P95 < 160ms

---

## 5 - WORKFLOW & COLLABORATION

### 5.1 negotiation/

#### User Stories
- As a **client**, I want to **negotiate terms** so that we can agree on scope and pricing.
- As a **freelancer**, I want to **propose counter-offers** so that I can reach better terms.
- As both parties, I want **negotiation history** so that discussions are transparent.

#### Flow
1. **OpenNegotiationCommand**(proposal_id) → CreateThread() → **Outbox:** negotiation.opened.v1
2. **ProposeCounterCommand**(negotiation_id, terms, amount) → Validate() | Persist() → **Outbox:** negotiation.countered.v1
3. **AcceptNegotiationCommand**(negotiation_id, counter_id) → LockTerms() | Accept() → **Outbox:** negotiation.accepted.v1
4. **DeclineNegotiationCommand**(negotiation_id, reason) → Decline() → **Outbox:** negotiation.declined.v1
5. **GetNegotiationHistoryQuery**(negotiation_id) → FetchThread() → NegotiationHistoryDTO

#### Projections
- negotiation_read
- negotiation_thread_read

#### Events Published
- negotiation.opened.v1
- negotiation.countered.v1
- negotiation.accepted.v1
- negotiation.declined.v1
- negotiation.expired.v1

#### RBAC/SLO
- **RBAC:** OWNER, CLIENT
- **SLO:** P95 < 180ms

---

### 5.2 invite/

#### User Stories
- As a **client**, I want to **invite freelancers** so that I can recruit talent proactively.
- As a **freelancer**, I want to **accept/decline invites** so that I can manage my opportunities.
- As a **system**, I want **invite expiration** so that old invites don't clutter the system.

#### Flow
1. **SendInviteCommand**(job_id, freelancer_id, message) → Validate() | CreateInvite() → **Outbox:** invite.sent.v1
2. **AcceptInviteCommand**(invite_id) → PrefillProposal() | Accept() → **Outbox:** invite.accepted.v1
3. **DeclineInviteCommand**(invite_id, reason) → Decline() → **Outbox:** invite.declined.v1
4. **ExpireInviteJob**() → FindExpired() | MarkExpired() → **Outbox:** invite.expired.v1
5. **GetInviteStatsQuery**(freelancer_id) → Aggregate() → InviteStatsDTO

#### Projections
- invite_read
- invite_stats_read

#### Events Published
- invite.sent.v1
- invite.accepted.v1
- invite.declined.v1
- invite.expired.v1

#### Decline Reasons
- NotInterested, Busy, RateTooLow, SkillMismatch, Other

#### RBAC/SLO
- **RBAC:** CLIENT (send), OWNER (accept/decline)
- **SLO:** P95 < 140ms

---

### 5.3 revision/

#### User Stories
- As a **freelancer**, I want **revision tracking** so that I can see how my proposal evolved.
- As a **client**, I want to **approve/reject revisions** so that proposals meet requirements.
- As a **system**, I want **immutable history** so that changes are auditable.

#### Flow
1. **CreateRevisionCommand**(proposal_id, changes) → CreateSnapshot() | Persist() → **Outbox:** revision.created.v1
2. **ApproveRevisionCommand**(revision_id) → Approve() | ApplyChanges() → **Outbox:** revision.approved.v1
3. **RejectRevisionCommand**(revision_id, reason) → Reject() → **Outbox:** revision.rejected.v1
4. **RevertToRevisionCommand**(proposal_id, revision_id) → Revert() → **Outbox:** proposal.reverted.v1
5. **GetRevisionHistoryQuery**(proposal_id) → FetchAll() → RevisionHistoryDTO

#### Projections
- revision_read
- revision_history_read

#### Events Published
- revision.created.v1
- revision.approved.v1
- revision.rejected.v1

#### RBAC/SLO
- **RBAC:** OWNER (create/revert), CLIENT (approve/reject)
- **SLO:** P95 < 160ms

---

### 5.4 collaboration/

#### User Stories
- As a **lead freelancer**, I want to **form teams** so that we can tackle larger projects together.
- As a **team member**, I want **clear roles and revenue splits** so that collaboration is fair.
- As a **system**, I want **permission enforcement** so that only authorized members can modify proposals.

#### Flow
1. **FormTeamCommand**(proposal_id, members[], roles[], revenue_splits[]) → ValidateRoles() | ValidateSplits() | Persist() → **Outbox:** team.formed.v1
2. **AddTeamMemberCommand**(proposal_id, member_id, role, revenue_share) → AuthorizeLead() | ValidateSplit() | Add() → **Outbox:** team.member.added.v1
3. **RemoveTeamMemberCommand**(proposal_id, member_id) → AuthorizeLead() | Redistribute() | Remove() → **Outbox:** team.member.removed.v1
4. **UpdateRevenueSplitCommand**(proposal_id, splits[]) → Validate100Percent() | Update() → **Outbox:** revenue.split.updated.v1
5. **GetTeamQuery**(proposal_id) → Fetch() → TeamDTO

#### Projections
- collaboration_read
- revenue_split_read

#### Events Published
- team.formed.v1
- team.member.added.v1
- team.member.removed.v1
- revenue.split.updated.v1

#### Roles
- Lead, Contributor, Consultant, Subcontractor

#### RBAC/SLO
- **RBAC:** LEAD (manage team), MEMBER (view)
- **SLO:** P95 < 200ms

---

## 6 - CLIENT INTERACTION

### 6.1 interview/

#### User Stories
- As a **client**, I want to **request interviews** so that I can evaluate candidates.
- As a **freelancer**, I want to **schedule interviews** so that I can showcase my skills.
- As a **system**, I want **calendar integration** so that scheduling is seamless.

#### Flow
1. **RequestInterviewCommand**(proposal_id, preferred_times[]) → CreateRequest() → **Outbox:** interview.requested.v1
2. **ScheduleInterviewCommand**(interview_id, scheduled_at) → ValidateTime() | Schedule() → **Outbox:** interview.scheduled.v1
3. **CompleteInterviewCommand**(interview_id, notes) → MarkComplete() → **Outbox:** interview.completed.v1
4. **CancelInterviewCommand**(interview_id, reason) → Cancel() → **Outbox:** interview.cancelled.v1
5. **GetInterviewAvailabilityQuery**(freelancer_id) → FetchCalendar() → AvailabilityDTO

#### Projections
- interview_read
- interview_schedule_read

#### Events Published
- interview.requested.v1
- interview.scheduled.v1
- interview.completed.v1
- interview.cancelled.v1

#### Events Consumed
- calendar.availability.updated

#### RBAC/SLO
- **RBAC:** CLIENT (request), OWNER (schedule/cancel)
- **SLO:** P95 < 180ms

---

### 6.2 feedback/

#### User Stories
- As a **client**, I want to **provide feedback** so that freelancers can improve.
- As a **freelancer**, I want to **see feedback** so that I can learn from client perspectives.
- As a **system**, I want **feedback aggregation** so that quality patterns emerge.

#### Flow
1. **GiveFeedbackCommand**(proposal_id, rating, comments) → ValidateRating() | Persist() → **Outbox:** feedback.given.v1
2. **UpdateFeedbackCommand**(feedback_id, updates) → Update() → **Outbox:** feedback.updated.v1
3. **RespondToFeedbackCommand**(feedback_id, response) → Respond() → **Outbox:** feedback.responded.v1
4. **GetFeedbackStatsQuery**(freelancer_id) → Aggregate() → FeedbackStatsDTO

#### Projections
- feedback_read
- feedback_stats_read

#### Events Published
- feedback.given.v1
- feedback.updated.v1
- feedback.responded.v1

#### RBAC/SLO
- **RBAC:** CLIENT (give), OWNER (respond/view)
- **SLO:** P95 < 150ms

---

### 6.3 shortlist/

#### User Stories
- As a **client**, I want to **shortlist proposals** so that I can narrow down candidates.
- As a **client**, I want to **rank shortlisted candidates** so that I can track preferences.
- As a **system**, I want **shortlist size limits** so that clients make timely decisions.

#### Flow
1. **AddToShortlistCommand**(job_id, proposal_id, rank) → ValidateShortlistSize() | Add() → **Outbox:** proposal.shortlisted.v1
2. **RemoveFromShortlistCommand**(job_id, proposal_id) → Remove() → **Outbox:** proposal.removed_from_shortlist.v1
3. **UpdateShortlistStatusCommand**(shortlist_id, status) → UpdateStatus() → **Outbox:** shortlist.status.updated.v1
4. **GetShortlistQuery**(job_id) → FetchSorted() → ShortlistDTO

#### Projections
- shortlist_read

#### Events Published
- proposal.shortlisted.v1
- proposal.removed_from_shortlist.v1
- shortlist.status.updated.v1

#### Shortlist Statuses
- Active, Interviewing, Selected, Rejected

#### RBAC/SLO
- **RBAC:** CLIENT
- **SLO:** P95 < 130ms

---

### 6.4 conversation/

#### User Stories
- As a **client/freelancer**, I want to **message about proposals** so that we can discuss details.
- As a **system**, I want **sentiment analysis** so that communication quality is tracked.
- As a **system**, I want **conversation threading** so that discussions stay organized.

#### Flow
1. **SendMessageCommand**(proposal_id, sender_id, content) → ValidateLength() | Persist() | AnalyzeSentiment() → **Outbox:** message.sent.v1
2. **MarkAsReadCommand**(message_id) → MarkRead() → **Outbox:** message.read.v1
3. **GetConversationQuery**(proposal_id) → FetchThread() → ConversationDTO
4. **GetSentimentAnalysisQuery**(proposal_id) → AggregateSentiment() → SentimentDTO

#### Projections
- conversation_read
- sentiment_analysis_read

#### Events Published
- message.sent.v1
- message.read.v1
- sentiment.analyzed.v1

#### Events Consumed
- communications.message.sent
- communications.message.read

#### RBAC/SLO
- **RBAC:** OWNER, CLIENT
- **SLO:** P95 < 160ms

---

## 7 - PERFORMANCE & ANALYTICS (CONSOLIDATED)

### 7.1 performance/ (Consolidates: analytics, engagement, response_tracking, conversion, metrics, ranking, health)

#### User Stories
- As a **freelancer**, I want to **track proposal views** so that I know client interest levels.
- As a **freelancer**, I want to **see engagement metrics** so that I can optimize my proposals.
- As a **freelancer**, I want **conversion funnel analysis** so that I understand where I'm losing opportunities.
- As a **freelancer**, I want **health scores** so that I know which proposals need improvement.
- As a **freelancer**, I want **ranking visibility** so that I understand my competitive position.
- As a **system**, I want **performance benchmarks** so that freelancers can compare against averages.

#### Flow
1. **TrackProposalViewCommand**(proposal_id, viewer_context) → ValidateViewer() | Record() | UpdateCache() → **Outbox:** proposal.viewed.v1
2. **RecordEngagementEventCommand**(proposal_id, event_type, metadata) → Record() | UpdateScore() → **Outbox:** engagement.recorded.v1
3. **RecordConversionCommand**(proposal_id, conversion_type) → UpdateFunnel() → **Outbox:** conversion.recorded.v1
4. **RecalculateHealthScoreCommand**(proposal_id) → FetchMetrics() | ComputeHealth() | PersistScore() → **Outbox:** health.score.calculated.v1
5. **UpdateRankingsJob**() → RecalculateAllRankings() → **Outbox:** ranking.updated.v1
6. **GetAnalyticsQuery**(proposal_id) → FetchMetrics() → AnalyticsDTO (views, click_rate, response_rate)
7. **GetEngagementMetricsQuery**(proposal_id) → FetchEngagement() → EngagementDTO (interest_signals, time_spent, interaction_count)
8. **GetConversionFunnelQuery**(proposal_id) → BuildFunnel() → ConversionFunnelDTO (view → message → interview → shortlist → hire)
9. **GetHealthScoreQuery**(proposal_id) → FetchHealth() → HealthDTO (score, issues, recommendations)
10. **GetRankingsQuery**(proposal_id) → FetchRank() → RankingDTO (position, visibility_score, factors)
11. **GetBenchmarksQuery**(category) → AggregateIndustry() → BenchmarkDTO (avg_views, avg_conversion_rate)

#### Projections
- proposal_analytics_read (views, impressions, CTR, sources)
- proposal_engagement_read (signals, scores, time_spent)
- conversion_funnel_read (touchpoints, attribution, rates)
- proposal_health_read (scores, warnings, recommendations)
- proposal_ranking_read (positions, visibility_scores, factors)
- performance_benchmarks_read (industry averages, percentiles)

#### Events Published
- proposal.viewed.v1
- proposal.viewed.by_client.v1
- engagement.recorded.v1
- conversion.recorded.v1
- health.score.calculated.v1
- ranking.updated.v1

#### Events Consumed
- client.responded.v1
- interview.scheduled.v1
- proposal.shortlisted.v1

#### Health Score Factors
- Completeness (all sections filled)
- Quality (grammar, length, relevance)
- Competitiveness (bid position, response time)
- Engagement (client views, interactions)

#### Ranking Factors
- Relevance score
- Profile strength
- Response time
- Client engagement
- Historical success rate
- Boost status

#### RBAC/SLO
- **RBAC:** OWNER (view own), SYSTEM (update)
- **SLO:** P95 < 200ms (queries), P95 < 300ms (calculations)

---

## 8 - SIMILARITY & DEDUPLICATION (CONSOLIDATED)

### 8.1 similarity/ (Consolidates: proposal_similarity, duplicate_check)

#### User Stories
- As a **system**, I want to **detect duplicate proposals** so that spam is prevented.
- As a **system**, I want to **find similar proposals** so that I can provide recommendations.
- As a **freelancer**, I want to **see differentiation scores** so that I can make my proposal unique.
- As a **trust team**, I want **clustering** so that patterns of abuse are visible.

#### Flow
1. **CreateFingerprintCommand**(proposal_id) → ExtractText() | ComputeHash() | PersistFingerprint() → **Outbox:** fingerprint.created.v1
2. **DetectDuplicatesCommand**(proposal_id) → CompareFingerprints() | ScoreSimilarity() | FlagIfDuplicate() → **Outbox:** duplicate.detected.v1
3. **ClusterProposalsJob**() → GroupSimilar() | ComputeCentroids() → **Outbox:** proposals.clustered.v1
4. **CalculateDifferentiationCommand**(proposal_id) → CompareToCluster() | ScoreUniqueness() → **Outbox:** differentiation.scored.v1
5. **GetSimilarProposalsQuery**(proposal_id, threshold) → FindSimilar() → SimilarProposalListDTO
6. **GetDuplicatesQuery**(proposal_id) → FetchMatches() → DuplicateListDTO
7. **GetDifferentiationScoreQuery**(proposal_id) → FetchScore() → DifferentiationDTO (score, unique_elements, suggestions)

#### Projections
- fingerprint_read (text_hash, structural_hash, semantic_hash)
- similarity_cluster_read (cluster_id, members, centroid)
- duplicate_detection_read (matches, scores, match_types)
- differentiation_read (scores, unique_elements, competitive_advantages)

#### Events Published
- fingerprint.created.v1
- duplicate.detected.v1
- proposals.clustered.v1
- differentiation.scored.v1

#### Match Types
- Exact (100% identical)
- Near (90-99% similar)
- Partial (70-89% similar)

#### RBAC/SLO
- **RBAC:** SYSTEM (detect), OWNER (view scores)
- **SLO:** P95 < 350ms (fingerprinting), P95 < 500ms (clustering)

---

## 9 - PORTFOLIO MANAGEMENT (CONSOLIDATED)

### 9.1 portfolio/ (Consolidates: portfolio_link, proposal_portfolio_selector)

#### User Stories
- As a **freelancer**, I want to **link portfolio items** so that I can showcase relevant work.
- As a **system**, I want **auto-selection** so that the most relevant portfolio items are attached.
- As a **freelancer**, I want to **reorder portfolio items** so that my best work appears first.
- As a **system**, I want **relevance scoring** so that portfolio selection is intelligent.

#### Flow
1. **LinkPortfolioItemCommand**(proposal_id, portfolio_item_id, relevance) → ValidateOwnership() | Link() → **Outbox:** portfolio.item.linked.v1
2. **UnlinkPortfolioItemCommand**(proposal_id, portfolio_item_id) → Unlink() → **Outbox:** portfolio.item.unlinked.v1
3. **AutoAttachPortfolioCommand**(proposal_id, max_items) → AnalyzeJob() | ScoreRelevance() | SelectTop() | Attach() → **Outbox:** portfolio.auto.selected.v1
4. **ReorderPortfolioCommand**(proposal_id, display_order[]) → ValidateOrder() | Update() → **Outbox:** portfolio.reordered.v1
5. **GetLinkedPortfolioQuery**(proposal_id) → FetchSorted() → PortfolioListDTO
6. **GetRelevanceScoresQuery**(proposal_id, portfolio_items[]) → ComputeScores() → RelevanceScoreListDTO

#### Projections
- portfolio_link_read (links, relevance_scores, display_order)
- portfolio_selection_read (selection_criteria, auto_selected_items)

#### Events Published
- portfolio.item.linked.v1
- portfolio.item.unlinked.v1
- portfolio.auto.selected.v1
- portfolio.reordered.v1

#### Selection Criteria
- Skill match
- Industry relevance
- Recency
- Client ratings
- Project size similarity

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 180ms (manual), P95 < 400ms (auto-select)

---

## 10 - ENGAGEMENT & FOLLOW-UPS (CONSOLIDATED)

### 10.1 engagement/ (Consolidates: proposal_engagement, proposal_follow_up)

#### User Stories
- As a **system**, I want to **track interest signals** so that engagement is measurable.
- As a **freelancer**, I want **automated follow-ups** so that I can stay top-of-mind.
- As a **freelancer**, I want to **bump proposals** so that they get renewed visibility.
- As a **system**, I want **follow-up limits** so that freelancers don't spam clients.

#### Flow
1. **RecordInterestSignalCommand**(proposal_id, signal_type, metadata) → Record() | UpdateEngagementScore() → **Outbox:** interest.detected.v1
2. **ScheduleFollowUpCommand**(proposal_id, trigger_conditions, delay) → ValidateLimit() | Schedule() → **Outbox:** follow_up.scheduled.v1
3. **SendFollowUpCommand**(follow_up_id) → GenerateMessage() | Send() | RecordSent() → **Outbox:** follow_up.sent.v1
4. **BumpProposalCommand**(proposal_id, bump_type) → ValidateBumpAllowed() | IncrementVisibility() → **Outbox:** proposal.bumped.v1
5. **GetInterestLevelQuery**(proposal_id) → ComputeInterest() → InterestLevelDTO (signals, score, trend)
6. **GetFollowUpScheduleQuery**(freelancer_id) → FetchUpcoming() → FollowUpScheduleDTO
7. **GetBumpHistoryQuery**(proposal_id) → FetchHistory() → BumpHistoryDTO

#### Projections
- interest_tracking_read (signals, scores, timestamps)
- follow_up_schedule_read (schedules, triggers, sent_history)
- bump_history_read (bump_count, last_bumped, response_after_bump)

#### Events Published
- interest.detected.v1
- follow_up.scheduled.v1
- follow_up.sent.v1
- proposal.bumped.v1

#### Interest Signal Types
- View, Save, Message, Share, DownloadAttachment

#### Bump Types
- Gentle, Standard, Urgent

#### Follow-up Triggers
- ClientViewed (send after 48h if no response)
- NoResponseAfter72h
- JobDeadlineApproaching
- CompetitorActivity

#### RBAC/SLO
- **RBAC:** OWNER (schedule/bump), SYSTEM (send)
- **SLO:** P95 < 160ms

---

## 11 - MODERATION & COMPLIANCE

### 11.1 spam_detection/

#### User Stories
- As a **system**, I want to **detect spam proposals** so that platform quality is maintained.
- As a **trust team**, I want **spam indicators** so that reviews are efficient.
- As a **freelancer**, I want **spam clearing** so that false positives are resolved.

#### Flow
1. **ScanProposalForSpamCommand**(proposal_id) → AnalyzeContent() | CheckLinks() | DetectPatterns() | ScoreSpam() → **Outbox:** spam.detected.v1
2. **FlagAsSpamCommand**(proposal_id, reasons[]) → Hide() | NotifyFreelancer() → **Outbox:** spam.flagged.v1
3. **ClearSpamFlagsCommand**(proposal_id) → Review() | Restore() → **Outbox:** spam.cleared.v1
4. **GetSpamStatusQuery**(proposal_id) → FetchStatus() → SpamStatusDTO (score, reasons, reviewed)

#### Projections
- spam_detection_read (scores, indicators, detection_timestamps)

#### Events Published
- spam.detected.v1
- spam.flagged.v1
- spam.cleared.v1

#### Spam Indicators
- LowQualityText (generic, template-like)
- SuspiciousLinks (external sites, tracking)
- MassSubmission (too many too fast)
- BotBehavior (timing patterns, identical content)

#### RBAC/SLO
- **RBAC:** SYSTEM (detect), TRUST_SAFETY (flag/clear)
- **SLO:** P95 < 300ms

---

### 11.2 flag/

#### User Stories
- As a **user**, I want to **flag inappropriate proposals** so that moderation can review.
- As a **moderator**, I want **flag management** so that cases are resolved efficiently.
- As a **system**, I want **flag aggregation** so that repeat offenders are identified.

#### Flow
1. **FlagProposalCommand**(proposal_id, reason, description) → ValidateCanFlag() | CreateFlag() → **Outbox:** proposal.flagged.v1
2. **UnflagProposalCommand**(flag_id) → Withdraw() → **Outbox:** proposal.unflagged.v1
3. **ReviewFlagCommand**(flag_id, action, notes) → TakeAction() | NotifyReporter() → **Outbox:** flag.reviewed.v1
4. **ResolveFlagCommand**(flag_id, resolution) → Close() → **Outbox:** flag.resolved.v1
5. **GetFlagsQuery**(proposal_id) → Fetch() → FlagListDTO

#### Projections
- flag_read (flags, reporters, statuses)
- flag_history_read (resolutions, moderator_actions)

#### Events Published
- proposal.flagged.v1
- proposal.unflagged.v1
- flag.reviewed.v1
- flag.resolved.v1

#### Flag Reasons
- Spam, Offensive, Plagiarism, Scam, MisleadingInformation, Other

#### RBAC/SLO
- **RBAC:** ANY_USER (flag), MODERATOR (review/resolve)
- **SLO:** P95 < 170ms

---

### 11.3 compliance/

#### User Stories
- As a **system**, I want **automated compliance checks** so that policy violations are caught early.
- As a **compliance team**, I want **violation reports** so that enforcement is consistent.
- As a **freelancer**, I want **clear violation explanations** so that I can correct issues.

#### Flow
1. **RequestComplianceCheckCommand**(proposal_id, check_types[]) → DelegateToComplianceBE() → **Outbox:** compliance.check.requested.v1
2. **ResolveViolationCommand**(proposal_id, violation_id, action) → TakeAction() | NotifyFreelancer() → **Outbox:** compliance.violation.resolved.v1
3. **GetComplianceStatusQuery**(proposal_id) → FetchChecks() → ComplianceStatusDTO (checks_passed, violations)
4. **GetViolationsQuery**(proposal_id) → FetchDetails() → ViolationListDTO

#### Projections
- compliance_read (check_results, violations, statuses)

#### Events Published
- compliance.check.requested.v1
- compliance.violation.resolved.v1

#### Events Consumed
- compliance.check.completed (from compliance-be)
- compliance.violation.detected (from compliance-be)

#### Check Types
- ContentPolicy, TermsOfService, LegalReview, DataPrivacy, IntellectualProperty

#### RBAC/SLO
- **RBAC:** SYSTEM (request), COMPLIANCE_TEAM (resolve)
- **SLO:** P95 < 250ms (async checks take longer)

---

## 12 - TEMPLATES & RATE CARDS

### 12.1 template/

#### User Stories
- As a **freelancer**, I want to **save proposal templates** so that I can reuse successful formats.
- As a **freelancer**, I want to **categorize templates** so that I can find them easily.
- As a **system**, I want **usage tracking** so that popular templates are highlighted.

#### Flow
1. **CreateTemplateCommand**(user_id, name, content, category) → Validate() | Persist() → **Outbox:** template.created.v1
2. **UpdateTemplateCommand**(template_id, updates) → Validate() | Update() → **Outbox:** template.updated.v1
3. **DeleteTemplateCommand**(template_id) → Delete() → **Outbox:** template.deleted.v1
4. **UseTemplateCommand**(template_id, proposal_id) → ApplyToProposal() | IncrementUsageCount() → **Outbox:** template.used.v1
5. **ListTemplatesQuery**(user_id, category) → Fetch() → TemplateListDTO
6. **GetTemplateUsageStatsQuery**(template_id) → AggregateUsage() → TemplateStatsDTO

#### Projections
- template_read (templates, metadata, categories)
- template_usage_read (usage_count, success_rate, last_used)

#### Events Published
- template.created.v1
- template.updated.v1
- template.deleted.v1
- template.used.v1

#### Template Categories
- WebDevelopment, Design, Writing, Marketing, DataScience, Other

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 130ms

---

### 12.2 rate_card/

#### User Stories
- As a **freelancer**, I want to **create rate cards** so that I can standardize pricing.
- As a **freelancer**, I want **package tiers** (starter, standard, premium) so that clients have options.
- As a **freelancer**, I want to **apply rate cards to proposals** so that pricing is consistent.

#### Flow
1. **CreateRateCardCommand**(freelancer_id, packages[], default_rates) → Validate() | Persist() → **Outbox:** rate_card.created.v1
2. **UpdateRateCardCommand**(rate_card_id, updates) → Validate() | Update() → **Outbox:** rate_card.updated.v1
3. **AddPackageCommand**(rate_card_id, package) → ValidatePackage() | Add() → **Outbox:** package.added.v1
4. **UpdatePackageCommand**(package_id, updates) → Update() → **Outbox:** package.updated.v1
5. **ApplyToProposalCommand**(proposal_id, package_id) → LoadPackage() | ApplyPricing() → **Outbox:** rate_card.applied.v1
6. **GetRateCardQuery**(freelancer_id) → Fetch() → RateCardDTO
7. **GetDefaultRatesQuery**(freelancer_id, job_type) → FetchDefaults() → DefaultRatesDTO

#### Projections
- rate_card_read (packages, rates, tiers)
- package_usage_read (package_id, usage_count, conversion_rate)

#### Events Published
- rate_card.created.v1
- rate_card.updated.v1
- package.added.v1
- package.updated.v1
- rate_card.applied.v1

#### Package Tiers
- Starter (basic deliverables, lower price)
- Standard (full deliverables, mid price)
- Premium (extended deliverables, higher price)

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 150ms

---

## EXTERNALIZED CAPABILITIES

The following capabilities have been **externalized to dedicated services** and are accessed via client interfaces:

### 🧠 Intelligence Services (via intelligence-be)
- **Strategy & Optimization**: Bid strategies, pricing optimization, timing recommendations
- **Predictions**: Success probability, win likelihood, competition analysis
- **Recommendations**: Personalized suggestions, best practices, content improvements
- **Keyword Optimization**: SEO suggestions, keyword extraction, relevance scoring
- **Personalization**: Client-style mirroring, variant generation, tone matching
- **A/B Testing**: Experiment management, variant testing, winner selection

### 🛡️ Risk & Assurance Services (via risk-be)
- **Risk Assessment**: Dispute prediction, red flag detection, risk scoring
- **Insurance & Escrow**: Payment protection, guarantee management
- **Dispute Prediction**: Early warning systems, mitigation playbooks

### 📄 Contracts Services (via contracts-be)
- **Terms Management**: NDA generation, IP rights, payment terms
- **Pricing**: Rate negotiation, pricing structures
- **Approval Workflows**: Multi-stage approvals, redlining, version control

### 📋 Procurement Services (via procurement-be)
- **RFP Management**: Request for proposals, quote management
- **Evaluation**: Rubrics, scoring, security questionnaires

### 🏪 Marketplace Services (via marketplace-be)
- **Visibility**: Boost campaigns, featured listings, premium placements
- **Budget Management**: Visibility spending, ROI tracking
- **Fees & Subscriptions**: Service fees, subscription management
- **Marketplace Listings**: Package marketplace, productized services

### 👥 Client Services (via clients-be / crm-be)
- **Client Insights**: Hire history, spend patterns, preferences
- **CRM**: Relationship management, client profiles

### 📅 Calendar Services (via calendar-be)
- **Deadline Tracking**: Milestone deadlines, proposal expiration
- **Sync**: External calendar integration

### ⏱️ Time Tracking Services (via time-tracking-be)
- **Estimates**: Time estimation, accuracy tracking
- **Tracking**: Actual hours, variance analysis

### 🌐 Localization Services (via localization-be)
- **Translations**: Multi-language proposal support
- **Locale Formatting**: Currency, dates, numbers

### 🔗 Integration Services (via integrations-be)
- **External Systems**: Third-party tool connections
- **Data Sync**: Bi-directional synchronization

### 🔐 Secure Sharing Services (via sharing-be)
- **Secure Links**: Time-limited, password-protected sharing
- **Access Logs**: View tracking, audit trails

### ✅ Compliance Services (via compliance-be)
- **Compliance Checks**: Policy validation, legal review
- **Verification**: Identity verification, credential checks

---

## EVENT FLOWS

### Events Published by proposals-be
```
# Core Proposal
- proposal.created.v1
- proposal.updated.v1
- proposal.submitted.v1
- proposal.accepted.v1
- proposal.rejected.v1
- proposal.withdrawn.v1
- proposal.archived.v1
- proposal.restored.v1

# Lifecycle
- expiration.set.v1
- expiration.extended.v1
- proposal.expiring.v1
- proposal.expired.v1
- pipeline.stage.moved.v1
- proposal.recycled.v1

# Bidding
- bid.placed.v1
- bid.updated.v1
- bid.withdrawn.v1
- bid.accepted.v1
- bid.outbid.v1
- auction.started.v1
- auction.ended.v1

# Collaboration
- team.formed.v1
- team.member.added.v1
- team.member.removed.v1
- revenue.split.updated.v1
- negotiation.opened.v1
- negotiation.accepted.v1

# Client Interaction
- interview.requested.v1
- interview.scheduled.v1
- proposal.shortlisted.v1
- feedback.given.v1

# Performance
- proposal.viewed.v1
- engagement.recorded.v1
- conversion.recorded.v1
- health.score.calculated.v1
- ranking.updated.v1

# Connects
- connects.used.v1
- connects.refunded.v1
- refund.requested.v1
```

### Events Consumed by proposals-be
```
# From jobs-be
- job.posted
- job.updated
- job.closed
- job.cancelled

# From contracts-be
- contract.created
- contract.started
- milestone.completed

# From users-be
- user.suspended
- user.banned
- user.profile.updated

# From reviews-be
- review.submitted

# From communications-be
- message.sent
- message.read

# From financial-be
- payment.processed
- payment.failed

# From subscriptions-be
- connects.purchased
- connects.granted
- subscription.expired

# From compliance-be
- compliance.check.completed
- compliance.violation.detected

# From intelligence-be
- prediction.completed
- recommendation.generated

# From storage-be
- file.uploaded
- file.deleted
```

---

## IMPLEMENTATION NOTES

### Architecture Patterns
- **Clean Architecture**: Domain → Application → Infrastructure → Interface layers
- **DDD**: Aggregates, value objects, domain events, repositories
- **CQRS**: Command/query separation with read models
- **Event Sourcing**: Partial (for audit trails, performance projections)
- **Outbox Pattern**: Reliable event publishing via platform-shared
- **Inbox Pattern**: Idempotent event consumption via platform-shared

### Data Consistency
- **Eventual Consistency**: Between services via Kafka events
- **Strong Consistency**: Within service boundaries via transactions
- **Read Models**: Denormalized projections for performance queries
- **Caching**: Redis for hot data (15-minute TTL, invalidation on updates)

### Scalability Considerations
- **Horizontal Scaling**: Stateless API servers
- **Sharding**: By freelancer_id for large datasets
- **Read Replicas**: For analytics and reporting queries
- **Background Jobs**: Leader election for singleton tasks (expiration sweeps, health calculations)
- **Rate Limiting**: Per-user, per-endpoint limits

### Security
- **Authentication**: Keycloak JWT via pkg/auth
- **Authorization**: RBAC with owner/client/system roles
- **PII Protection**: Field-level encryption for sensitive data
- **Audit Logging**: All mutations logged with actor context
- **Input Validation**: Multi-layer (API, application, domain)

### Observability
- **Logging**: Structured logs via platform-shared/logging
- **Metrics**: Prometheus metrics via platform-shared/metrics
- **Tracing**: OpenTelemetry via platform-shared/tracing
- **Alerts**: P95 latency, error rates, queue depths

### Testing Strategy
- **Unit Tests**: Domain logic, value objects, business rules
- **Integration Tests**: Repository implementations, event handlers
- **E2E Tests**: Full proposal lifecycle scenarios
- **Property Tests**: Performance scoring, similarity algorithms
- **Chaos Tests**: Event delivery, duplicate handling, reordering

---

## SLO SUMMARY

| Domain | P95 Latency | Availability |
|--------|-------------|--------------|
| Core Proposal | < 180ms | 99.9% |
| Lifecycle | < 160ms | 99.9% |
| Bidding | < 200ms | 99.95% |
| Connects | < 100ms | 99.95% |
| Collaboration | < 200ms | 99.9% |
| Client Interaction | < 180ms | 99.9% |
| Performance | < 300ms | 99.5% |
| Similarity | < 500ms | 99% |
| Portfolio | < 400ms | 99% |
| Engagement | < 160ms | 99.9% |
| Moderation | < 300ms | 99% |
| Templates | < 130ms | 99.9% |

---

## GLOSSARY

- **Proposal**: A freelancer's application to a job posting
- **Bid**: The pricing component of a proposal
- **Connect**: Platform currency used to submit proposals
- **Pipeline**: The stages a proposal moves through (draft → submitted → shortlisted → hired)
- **Health Score**: Automated assessment of proposal quality and competitiveness
- **Fingerprint**: Hash-based signature for duplicate detection
- **Bump**: Action to increase proposal visibility
- **Rate Card**: Predefined pricing packages offered by a freelancer
- **Shortlist**: Client's curated list of top candidates
- **Negotiation Thread**: Back-and-forth discussion on terms and pricing

---

**End of proposals-be User Stories v2.0**