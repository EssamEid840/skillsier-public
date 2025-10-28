# 📦 **contracts-be - Contract Management Service - Complete User Stories**

---

## **Global Conventions**

### Event Envelope (All Events)
```json
{
  "event_id": "uuid-v7",
  "event_type": "contract.created.v1",
  "event_version": "v1",
  "aggregate_id": "contract-uuid",
  "aggregate_type": "Contract",
  "occurred_at": "ISO8601",
  "causation_id": "parent-event-id",
  "correlation_id": "trace-id",
  "metadata": {
    "user_id": "...",
    "service": "contracts-be",
    "idempotency_key": "..."
  },
  "payload": { ... }
}
```

### Idempotent Write-Path
- All commands accept `idempotency_key` parameter
- Store `(idempotency_key, aggregate_id, event_version)` in `idempotency_log` table
- TTL: 7 days
- Return existing result if duplicate detected within TTL window

### Non-PII Event Payloads
- Events NEVER contain PII in payload (no emails, phones, addresses, SSNs)
- Events reference IDs only: `contract_id`, `user_id`, `milestone_id`, `sow_id`
- Consumers fetch PII via authenticated API calls if needed
- Example: `contract.created.v1` contains `contract_id`, `client_id`, `freelancer_id` but NOT names or emails

### Folder Structure Alignment
- All domain entities map to `internal/domain/{context}/`
- All repositories map to `internal/infrastructure/persistence/postgres/{context}_repository.go`
- All services map to `internal/application/{context}/service.go`
- All handlers map to `internal/interfaces/http/v1/handlers/{context}_handler.go`
- All routes map to `internal/interfaces/http/v1/routes/{context}_routes.go`

### Events Catalog Integration
- All events published are registered in `contracts/events/contracts/` catalog
- Event schemas versioned with semantic versioning (v1, v2, etc.)
- Breaking changes require new event version
- Consumers subscribe via Dapr pub/sub with scopes: `["contracts-be"]`

### Caching Strategy
- Cache keys follow pattern: `contracts:{contract_id}:{context}:{version}`
- TTLs defined in `internal/infrastructure/cache/redis/keys.go`
- Invalidation rules map events to cache keys in `invalidation_rules.go`
- Singleflight prevents cache stampedes for hot keys

### Observability
- All commands/queries emit spans with OpenTelemetry
- Metrics tracked: P95 latency, error rate, event publish lag, contract_value_total, milestone_release_latency, dispute_resolution_time, sla_breach_rate, budget_overrun_count
- Structured logging with correlation_id for tracing
- Health checks: `/healthz/live` (liveness), `/healthz/ready` (readiness)

### Security
- All endpoints require JWT authentication via Keycloak
- RBAC enforced at service layer (CLIENT, FREELANCER, ADMIN, SYSTEM, PUBLIC)
- PII encrypted at rest using KMS envelope encryption
- Sensitive fields redacted in logs via PII redactor

### Data Retention & Erasure
- Contract data retention: 10 years (legal/compliance)
- Event logs retention: 90 days (projections can replay)
- GDPR/CCPA erasure: cascading deletion with `contract.erased.v1` event
- Financial records: 7 years minimum (legal requirement)

---

## **1 - CORE CONTRACT DOMAIN**

### 1.1 contract/

#### User Stories
- As a **client**, I want to **create a contract from an accepted proposal** so that work can begin formally.
- As a **client**, I want to **create a direct contract** (without proposal) so that I can quickly hire trusted freelancers.
- As a **freelancer**, I want to **view contract details** so that I understand my obligations.
- As a **client**, I want to **activate a contract** after terms are finalized so that work officially starts.
- As a **client/freelancer**, I want to **pause a contract** temporarily so that work can be suspended.
- As a **client/freelancer**, I want to **resume a paused contract** so that work can continue.
- As a **client/freelancer**, I want to **complete a contract** when all milestones are done so that final payment is triggered.
- As a **client/freelancer**, I want to **terminate a contract** early so that engagement can end with proper handling.
- As a **client**, I want to **update contract metadata** (title, notes) so that organization is maintained.
- As a **system**, I want to **validate contract state transitions** so that invalid operations are prevented.
- As a **system**, I want to **integrate with financial-be** for escrow/payments so that funds are managed properly.
- As a **system**, I want to **track contract value metrics** so that business reporting is accurate.

#### Flow
1. **CreateContractFromProposalCommand**(proposal_id, job_id, client_id, freelancer_id, terms) → ValidateProposal() | ValidateParticipants() | CreateEscrow(financial-be) | Persist() → **Outbox:** contract.created.v1
2. **CreateDirectContractCommand**(client_id, freelancer_id, job_id, terms, contract_type) → ValidateParticipants() | ValidateTerms() | CreateEscrow(financial-be) | Persist() → **Outbox:** contract.direct.created.v1
3. **ActivateContractCommand**(contract_id, activated_by, start_date) → ValidateState(Draft→Active) | FundEscrow(financial-be) | Activate() → **Outbox:** contract.activated.v1
4. **PauseContractCommand**(contract_id, reason, paused_by) → ValidateState(Active→Paused) | HoldPayments(financial-be) | Pause() → **Outbox:** contract.paused.v1
5. **ResumeContractCommand**(contract_id, resumed_by) → ValidateState(Paused→Active) | ResumePayments(financial-be) | Resume() → **Outbox:** contract.resumed.v1
6. **CompleteContractCommand**(contract_id, completed_by, completion_notes) → ValidateAllMilestones() | ReleaseFinalPayment(financial-be) | Complete() → **Outbox:** contract.completed.v1
7. **TerminateContractCommand**(contract_id, reason, termination_type, refund_policy, terminated_by) → ValidateReason() | HandleRefunds(financial-be) | Terminate() → **Outbox:** contract.terminated.v1
8. **UpdateContractCommand**(contract_id, updates, updated_by) → AuthorizeParty() | ValidateUpdates() | Apply() → **Outbox:** contract.updated.v1
9. **GetContractQuery**(contract_id) → AuthorizeAccess() | Fetch() | Enrich() → ContractDTO
10. **ListContractsQuery**(filters, pagination) → ApplyFilters() | Paginate() → ContractListDTO
11. **GetContractStatsQuery**(contract_id) → Aggregate() → ContractStatsDTO
12. **GetContractTimelineQuery**(contract_id) → FetchEvents() → TimelineDTO

#### Projections
- contract_read
- contract_stats_read
- contract_timeline_read
- contract_financials_read
- contract_parties_read

#### Events Published
- contract.created.v1
- contract.direct.created.v1
- contract.activated.v1
- contract.paused.v1
- contract.resumed.v1
- contract.completed.v1
- contract.terminated.v1
- contract.updated.v1

#### Events Consumed
- proposal.accepted.v1 (trigger contract creation)
- payment.processed.v1 (update financial status)
- escrow.funded.v1 (enable activation)
- user.suspended.v1 (auto-pause contracts)

#### RBAC/SLO
- **RBAC:** CLIENT (create/activate/pause/terminate), FREELANCER (pause/complete), BOTH (resume/update), PUBLIC (view own contracts), ADMIN (view all/force terminate)
- **SLO:** P95 < 300ms (create), P95 < 200ms (lifecycle operations), P95 < 150ms (read)

---

### 1.2 contract_type/

#### User Stories
- As a **system**, I want to **support multiple contract types** (Hourly, Fixed-Price, Milestone-Based, Retainer, Subscription) so that diverse engagement models are covered.
- As a **client**, I want to **specify contract type during creation** so that payment terms are clear.
- As a **system**, I want to **enforce type-specific rules** so that contract behavior matches type.
- As a **system**, I want to **validate rate structures per type** so that pricing is appropriate.

#### Flow
1. **SetContractTypeCommand**(contract_id, contract_type, type_config) → ValidateType() | ValidateConfig() | Set() → **Outbox:** contract.type.set.v1
2. **ValidateContractTypeConfigQuery**(contract_type, config) → Validate() → ValidationResultDTO
3. **GetContractTypeRulesQuery**(contract_type) → Fetch() → ContractTypeRulesDTO

#### Projections
- contract_types_read
- contract_type_rules_read

#### Events Published
- contract.type.set.v1

#### RBAC/SLO
- **RBAC:** CLIENT/FREELANCER (set during creation), PUBLIC (view rules)
- **SLO:** P95 < 180ms

---

## **2 - STATEMENT OF WORK (SOW) DOMAIN**

### 2.1 sow/

#### User Stories
- As a **client**, I want to **create a detailed SOW** with scope, deliverables, timeline, and acceptance criteria so that expectations are documented.
- As a **client**, I want to **version SOWs** when changes occur so that history is tracked.
- As a **client/freelancer**, I want to **approve SOW** bilaterally so that both parties agree.
- As a **client/freelancer**, I want to **reject SOW** with reasons so that rework is clear.
- As a **client**, I want to **view SOW diff** between versions so that changes are transparent.
- As a **system**, I want to **enforce SOW approval** before contract activation so that terms are finalized.
- As a **system**, I want to **support SOW templates** so that creation is faster.

#### Flow
1. **CreateSOWCommand**(contract_id, title, scope, deliverables, timeline, acceptance_criteria, created_by) → ValidateContract() | ValidateContent() | Persist() → **Outbox:** sow.created.v1
2. **UpdateSOWCommand**(sow_id, updates, updated_by) → AuthorizeParty() | CreateVersion() | GenerateDiff() | Persist() → **Outbox:** sow.updated.v1
3. **ApproveSOWCommand**(sow_id, approved_by, signature) → ValidatePendingApproval() | Approve() | CheckBilateralApproval() → **Outbox:** sow.approved.v1
4. **RejectSOWCommand**(sow_id, reason, rejected_by) → ValidatePendingApproval() | Reject() → **Outbox:** sow.rejected.v1
5. **GetSOWQuery**(sow_id) → AuthorizeAccess() | Fetch() → SOWDTO
6. **ListSOWVersionsQuery**(sow_id) → Fetch() → SOWVersionListDTO
7. **GetSOWDiffQuery**(sow_id, version_from, version_to) → GenerateDiff() → SOWDiffDTO
8. **GetLatestSOWQuery**(contract_id) → Fetch() → SOWDTO

#### Projections
- sow_read
- sow_versions_read
- sow_approvals_read

#### Events Published
- sow.created.v1
- sow.updated.v1
- sow.approved.v1
- sow.rejected.v1
- sow.version.created.v1

#### RBAC/SLO
- **RBAC:** CLIENT/FREELANCER (create/update/approve/reject), PUBLIC (view own SOWs)
- **SLO:** P95 < 250ms (create/update), P95 < 180ms (approve/reject), P95 < 120ms (read)

---

## **3 - FINANCIAL HOLDS DOMAIN (RISK MANAGEMENT)**

### 3.1 financial_hold/

#### User Stories
- As a **risk system**, I want to **place financial holds** on contracts due to fraud signals so that losses are prevented.
- As a **compliance officer**, I want to **extend holds** when investigations need more time.
- As a **admin**, I want to **release holds** when issues are resolved so that work can resume.
- As a **system**, I want to **track hold reasons** so that transparency is maintained.
- As a **client/freelancer**, I want to **be notified of holds** so that I'm aware of status.
- As a **system**, I want to **prevent milestone releases** during holds so that funds are protected.

#### Flow
1. **PlaceFinancialHoldCommand**(contract_id, hold_type, reason, severity, placed_by, duration) → ValidateContract() | NotifyParties(communications-be) | PlaceHold() | FreezePayments(financial-be) → **Outbox:** financial_hold.placed.v1
2. **ExtendFinancialHoldCommand**(hold_id, additional_duration, reason, extended_by) → AuthorizeAdmin() | Extend() | NotifyParties() → **Outbox:** financial_hold.extended.v1
3. **ReleaseFinancialHoldCommand**(hold_id, resolution_notes, released_by) → AuthorizeAdmin() | Release() | UnfreezePayments(financial-be) | NotifyParties() → **Outbox:** financial_hold.released.v1
4. **GetActiveHoldsQuery**(contract_id) → Fetch() → FinancialHoldListDTO
5. **GetHoldHistoryQuery**(contract_id) → Fetch() → HoldHistoryDTO

#### Projections
- financial_holds_read
- hold_history_read
- hold_impact_read

#### Events Published
- financial_hold.placed.v1
- financial_hold.extended.v1
- financial_hold.released.v1

#### Events Consumed
- risk.alert.triggered.v1 (auto-place holds)
- dispute.opened.v1 (auto-place holds)
- chargeback.created.v1 (auto-place holds)

#### RBAC/SLO
- **RBAC:** RISK_SYSTEM (place), ADMIN (extend/release), CLIENT/FREELANCER (view own)
- **SLO:** P95 < 200ms (place/release), P95 < 150ms (extend), P95 < 100ms (read)

---

## **4 - MILESTONES DOMAIN**

### 4.1 milestone/

#### User Stories
- As a **client**, I want to **create milestones** with amount, due date, and description so that project phases are clear.
- As a **freelancer**, I want to **submit milestone for approval** when work is complete so that payment is requested.
- As a **client**, I want to **approve milestone** after reviewing so that freelancer gets paid.
- As a **client**, I want to **reject milestone** with feedback so that revisions are requested.
- As a **client**, I want to **release payment** for approved milestone so that freelancer is compensated.
- As a **freelancer**, I want to **dispute milestone rejection** so that disagreements are handled.
- As a **system**, I want to **auto-release milestones** after N days if no action so that payments aren't indefinitely delayed.
- As a **system**, I want to **track milestone completion rate** so that project health is monitored.
- As a **client**, I want to **reorder milestones** so that sequencing is flexible.

#### Flow
1. **CreateMilestoneCommand**(contract_id, title, description, amount, due_date, order, created_by) → ValidateContract() | ValidateBudget() | CreateEscrowHold(financial-be) | Persist() → **Outbox:** milestone.created.v1
2. **SubmitMilestoneCommand**(milestone_id, deliverable_ids[], submission_notes, submitted_by) → ValidateState(InProgress→Submitted) | LinkDeliverables() | NotifyClient(communications-be) | StartAutoReleaseTimer() → **Outbox:** milestone.submitted.v1
3. **ApproveMilestoneCommand**(milestone_id, approval_notes, approved_by) → ValidateState(Submitted→Approved) | StopAutoReleaseTimer() → **Outbox:** milestone.approved.v1
4. **RejectMilestoneCommand**(milestone_id, rejection_reason, feedback, rejected_by) → ValidateState(Submitted→Rejected) | StopAutoReleaseTimer() | NotifyFreelancer(communications-be) → **Outbox:** milestone.rejected.v1
5. **ReleaseMilestonePaymentCommand**(milestone_id, released_by) → ValidateState(Approved→Released) | ReleaseEscrow(financial-be) | ProcessPayout() → **Outbox:** milestone.payment.released.v1
6. **DisputeMilestoneCommand**(milestone_id, dispute_reason, evidence_urls[], disputed_by) → ValidateRejection() | CreateDispute(admin-be) | PlaceHold() → **Outbox:** milestone.disputed.v1
7. **UpdateMilestoneCommand**(milestone_id, updates, updated_by) → AuthorizeClient() | ValidateState(Draft/InProgress) | Apply() → **Outbox:** milestone.updated.v1
8. **ReorderMilestonesCommand**(contract_id, milestone_order[], reordered_by) → ValidateContract() | Reorder() → **Outbox:** milestones.reordered.v1
9. **GetMilestoneQuery**(milestone_id) → AuthorizeAccess() | Fetch() → MilestoneDTO
10. **ListMilestonesQuery**(contract_id, filters) → ApplyFilters() | Paginate() → MilestoneListDTO
11. **GetMilestoneStatsQuery**(contract_id) → Aggregate() → MilestoneStatsDTO

#### Projections
- milestone_read
- milestone_payments_read
- milestone_timeline_read
- milestone_stats_read

#### Events Published
- milestone.created.v1
- milestone.updated.v1
- milestone.submitted.v1
- milestone.approved.v1
- milestone.rejected.v1
- milestone.payment.released.v1
- milestone.disputed.v1
- milestone.auto_released.v1
- milestones.reordered.v1

#### Events Consumed
- escrow.funded.v1 (enable milestone creation)
- dispute.resolved.v1 (handle disputed milestones)
- timer.expired (auto-release trigger)

#### RBAC/SLO
- **RBAC:** CLIENT (create/update/approve/reject/release/reorder), FREELANCER (submit/dispute), PUBLIC (view own), ADMIN (override)
- **SLO:** P95 < 250ms (create), P95 < 200ms (submit/approve/reject), P95 < 300ms (release), P95 < 150ms (read)

---

## **5 - DELIVERABLES DOMAIN**

### 5.1 deliverable/

#### User Stories
- As a **freelancer**, I want to **submit deliverables** (files, links, descriptions) for milestone so that work is provided.
- As a **client**, I want to **review deliverables** so that quality is assessed.
- As a **client**, I want to **accept deliverables** when satisfied so that milestone approval follows.
- As a **client**, I want to **request revisions** with specific feedback so that improvements are clear.
- As a **freelancer**, I want to **resubmit revised deliverables** so that feedback is addressed.
- As a **system**, I want to **link deliverables to milestones** so that tracking is accurate.
- As a **system**, I want to **integrate with storage-be** for file management so that documents are properly stored.
- As a **system**, I want to **track revision count** so that excessive rework is flagged.

#### Flow
1. **SubmitDeliverableCommand**(milestone_id, title, description, file_urls[], external_links[], submitted_by) → ValidateMilestone() | UploadFiles(storage-be) | Persist() | NotifyClient(communications-be) → **Outbox:** deliverable.submitted.v1
2. **ReviewDeliverableCommand**(deliverable_id, review_status, review_notes, reviewed_by) → ValidateSubmission() | Review() → **Outbox:** deliverable.reviewed.v1
3. **AcceptDeliverableCommand**(deliverable_id, acceptance_notes, accepted_by) → ValidateReview() | Accept() → **Outbox:** deliverable.accepted.v1
4. **RequestRevisionCommand**(deliverable_id, revision_feedback, areas_for_improvement[], requested_by) → ValidateReview() | IncrementRevisionCount() | NotifyFreelancer(communications-be) → **Outbox:** deliverable.revision.requested.v1
5. **ResubmitDeliverableCommand**(deliverable_id, updates, new_files[], resubmitted_by) → ValidateRevisionRequest() | UpdateFiles(storage-be) | CreateVersion() → **Outbox:** deliverable.resubmitted.v1
6. **GetDeliverableQuery**(deliverable_id) → AuthorizeAccess() | Fetch() → DeliverableDTO
7. **ListDeliverablesQuery**(milestone_id) → Fetch() → DeliverableListDTO
8. **GetDeliverableVersionsQuery**(deliverable_id) → Fetch() → DeliverableVersionListDTO

#### Projections
- deliverable_read
- deliverable_versions_read
- deliverable_reviews_read

#### Events Published
- deliverable.submitted.v1
- deliverable.reviewed.v1
- deliverable.accepted.v1
- deliverable.revision.requested.v1
- deliverable.resubmitted.v1
- deliverable.version.created.v1

#### Events Consumed
- milestone.approved.v1 (auto-accept linked deliverables)
- storage.file.uploaded.v1 (track file uploads)

#### RBAC/SLO
- **RBAC:** FREELANCER (submit/resubmit), CLIENT (review/accept/request revision), PUBLIC (view own)
- **SLO:** P95 < 250ms (submit), P95 < 180ms (review/accept/request), P95 < 150ms (read)

---

## **6 - TIME TRACKING DOMAIN**

### 6.1 timesheet/

#### User Stories
- As a **freelancer**, I want to **add time entries** (date, hours, description) so that work is logged.
- As a **freelancer**, I want to **submit weekly timesheets** for approval so that payment is requested.
- As a **client**, I want to **approve timesheets** after review so that freelancer gets paid.
- As a **client**, I want to **reject timesheets** with feedback so that corrections are made.
- As a **freelancer**, I want to **dispute timesheet rejection** so that disagreements are resolved.
- As a **system**, I want to **calculate total hours and cost** so that billing is accurate.
- As a **system**, I want to **validate weekly hour caps** so that overwork is flagged.
- As a **system**, I want to **integrate with work diary** so that time is verified.

#### Flow
1. **AddTimeEntryCommand**(contract_id, date, hours, description, task_type, added_by) → ValidateContract(Hourly) | ValidateHours() | Persist() → **Outbox:** time_entry.added.v1
2. **UpdateTimeEntryCommand**(entry_id, updates, updated_by) → AuthorizeOwner() | ValidateNotSubmitted() | Apply() → **Outbox:** time_entry.updated.v1
3. **DeleteTimeEntryCommand**(entry_id, deleted_by) → AuthorizeOwner() | ValidateNotSubmitted() | Delete() → **Outbox:** time_entry.deleted.v1
4. **SubmitTimesheetCommand**(contract_id, week_ending_date, entry_ids[], submitted_by) → ValidateWeek() | CalculateTotals() | NotifyClient(communications-be) → **Outbox:** timesheet.submitted.v1
5. **ApproveTimesheetCommand**(timesheet_id, approval_notes, approved_by) → ValidateSubmission() | Approve() | TriggerPayment(financial-be) → **Outbox:** timesheet.approved.v1
6. **RejectTimesheetCommand**(timesheet_id, rejection_reason, rejected_by) → ValidateSubmission() | Reject() | NotifyFreelancer(communications-be) → **Outbox:** timesheet.rejected.v1
7. **DisputeTimesheetCommand**(timesheet_id, dispute_reason, evidence[], disputed_by) → ValidateRejection() | CreateDispute(admin-be) → **Outbox:** timesheet.disputed.v1
8. **GetTimesheetQuery**(timesheet_id) → AuthorizeAccess() | Fetch() → TimesheetDTO
9. **ListTimesheetsQuery**(contract_id, filters) → ApplyFilters() → TimesheetListDTO
10. **GetTimeEntriesQuery**(contract_id, date_range) → Fetch() → TimeEntryListDTO
11. **GetWeeklySummaryQuery**(contract_id, week_ending_date) → Aggregate() → WeeklySummaryDTO

#### Projections
- timesheet_read
- time_entries_read
- timesheet_approvals_read
- weekly_summaries_read

#### Events Published
- time_entry.added.v1
- time_entry.updated.v1
- time_entry.deleted.v1
- timesheet.submitted.v1
- timesheet.approved.v1
- timesheet.rejected.v1
- timesheet.disputed.v1

#### Events Consumed
- work_diary.entry.recorded.v1 (cross-validate hours)
- contract.paused.v1 (freeze time entry)

#### RBAC/SLO
- **RBAC:** FREELANCER (add/update/delete/submit/dispute), CLIENT (approve/reject), PUBLIC (view own)
- **SLO:** P95 < 180ms (add/update), P95 < 220ms (submit), P95 < 200ms (approve/reject), P95 < 120ms (read)

---

### 6.2 work_diary/

#### User Stories
- As a **freelancer**, I want to **record work diary entries** (activity level, screenshots, app usage) so that work is verified.
- As a **system**, I want to **capture periodic screenshots** so that activity is documented.
- As a **freelancer**, I want to **blur screenshots** for privacy so that sensitive info is protected.
- As a **freelancer**, I want to **delete inappropriate screenshots** within grace period so that control is maintained.
- As a **client**, I want to **view work diary** so that activity is monitored.
- As a **system**, I want to **apply retention policies** (30/60/90 days) so that storage is managed.
- As a **system**, I want to **track activity levels** (idle/low/medium/high) so that productivity is measured.

#### Flow
1. **RecordWorkDiaryEntryCommand**(contract_id, timestamp, activity_level, screenshot_url, apps_used[], recorded_by) → ValidateContract(Hourly) | UploadScreenshot(storage-be) | Persist() → **Outbox:** work_diary.entry.recorded.v1
2. **BlurScreenshotCommand**(entry_id, blur_regions[], blurred_by) → AuthorizeOwner() | ApplyBlur(storage-be) → **Outbox:** work_diary.screenshot.blurred.v1
3. **DeleteScreenshotCommand**(entry_id, reason, deleted_by) → AuthorizeOwner() | ValidateGracePeriod() | Delete(storage-be) → **Outbox:** work_diary.screenshot.deleted.v1
4. **GetWorkDiaryQuery**(contract_id, date_range) → AuthorizeAccess() | Fetch() → WorkDiaryDTO
5. **GetActivitySummaryQuery**(contract_id, date) → Aggregate() → ActivitySummaryDTO
6. **ApplyRetentionPolicyCommand**(retention_days) → FindExpiredEntries() | DeleteScreenshots(storage-be) → **Outbox:** work_diary.retention.applied.v1

#### Projections
- work_diary_read
- activity_summary_read
- screenshot_metadata_read

#### Events Published
- work_diary.entry.recorded.v1
- work_diary.screenshot.blurred.v1
- work_diary.screenshot.deleted.v1
- work_diary.retention.applied.v1

#### RBAC/SLO
- **RBAC:** FREELANCER (record/blur/delete), CLIENT (view), SYSTEM (retention)
- **SLO:** P95 < 200ms (record), P95 < 180ms (blur/delete), P95 < 150ms (read)

---

## **7 - TEMPLATES DOMAIN**

### 7.1 template/

#### User Stories
- As a **client**, I want to **create contract templates** with predefined terms so that hiring is faster.
- As a **client**, I want to **publish templates** so that they're available for use.
- As a **client**, I want to **use template** to create contracts quickly so that setup is streamlined.
- As a **client**, I want to **clone templates** so that similar contracts are easy to create.
- As a **system**, I want to **provide standard templates** (NDA, hourly, fixed-price) so that common scenarios are covered.
- As a **client**, I want to **search templates by category** so that finding is easy.
- As a **client**, I want to **version templates** so that updates don't break active contracts.

#### Flow
1. **CreateTemplateCommand**(created_by, name, description, category, template_config) → ValidateConfig() | Persist() → **Outbox:** template.created.v1
2. **UpdateTemplateCommand**(template_id, updates, updated_by) → AuthorizeOwner() | CreateVersion() | Apply() → **Outbox:** template.updated.v1
3. **PublishTemplateCommand**(template_id, published_by) → ValidateComplete() | Publish() → **Outbox:** template.published.v1
4. **UnpublishTemplateCommand**(template_id, unpublished_by) → AuthorizeOwner() | Unpublish() → **Outbox:** template.unpublished.v1
5. **UseTemplateCommand**(template_id, contract_params, used_by) → ValidatePublished() | CloneConfig() | CreateContract() → **Outbox:** template.used.v1
6. **CloneTemplateCommand**(template_id, new_name, cloned_by) → AuthorizeAccess() | Clone() → **Outbox:** template.cloned.v1
7. **GetTemplateQuery**(template_id) → AuthorizeAccess() | Fetch() → TemplateDTO
8. **SearchTemplatesQuery**(filters, pagination) → ApplyFilters() | Search() → TemplateListDTO
9. **ListTemplateVersionsQuery**(template_id) → Fetch() → TemplateVersionListDTO

#### Projections
- template_read
- template_versions_read
- template_usage_read

#### Events Published
- template.created.v1
- template.updated.v1
- template.published.v1
- template.unpublished.v1
- template.used.v1
- template.cloned.v1

#### RBAC/SLO
- **RBAC:** CLIENT (create/update/publish/use/clone), PUBLIC (search published), ADMIN (manage standard templates)
- **SLO:** P95 < 200ms (create/update), P95 < 150ms (use/clone), P95 < 120ms (search)

---

## **8 - CONTRACT CHANGES DOMAIN**

### 8.1 amendment/

#### User Stories
- As a **client/freelancer**, I want to **propose contract amendments** (rate, scope, timeline changes) so that terms can evolve.
- As a **client/freelancer**, I want to **approve amendments** bilaterally so that both parties agree.
- As a **client/freelancer**, I want to **reject amendments** with reasons so that disagreements are clear.
- As a **system**, I want to **apply approved amendments** so that contract is updated.
- As a **system**, I want to **track amendment history** so that changes are auditable.
- As a **system**, I want to **validate financial impact** before approval so that budget is maintained.

#### Flow
1. **ProposeAmendmentCommand**(contract_id, amendment_type, proposed_changes, justification, proposed_by) → ValidateContract(Active) | ValidateChanges() | Persist() | NotifyCounterparty(communications-be) → **Outbox:** amendment.proposed.v1
2. **ApproveAmendmentCommand**(amendment_id, approval_notes, approved_by) → ValidateCounterparty() | Approve() | CheckBilateralApproval() → **Outbox:** amendment.approved.v1
3. **RejectAmendmentCommand**(amendment_id, rejection_reason, rejected_by) → ValidateCounterparty() | Reject() → **Outbox:** amendment.rejected.v1
4. **ApplyAmendmentCommand**(amendment_id, applied_by) → ValidateBilateralApproval() | ApplyChanges() | UpdateFinancials(financial-be) | CreateContractVersion() → **Outbox:** amendment.applied.v1
5. **WithdrawAmendmentCommand**(amendment_id, withdrawn_by) → AuthorizeProposer() | Withdraw() → **Outbox:** amendment.withdrawn.v1
6. **GetAmendmentQuery**(amendment_id) → AuthorizeAccess() | Fetch() → AmendmentDTO
7. **ListAmendmentsQuery**(contract_id, filters) → ApplyFilters() → AmendmentListDTO
8. **GetAmendmentImpactQuery**(amendment_id) → CalculateImpact() → AmendmentImpactDTO

#### Projections
- amendment_read
- amendment_approvals_read
- amendment_history_read

#### Events Published
- amendment.proposed.v1
- amendment.approved.v1
- amendment.rejected.v1
- amendment.applied.v1
- amendment.withdrawn.v1

#### Events Consumed
- financial.escrow.adjusted.v1 (track financial changes)

#### RBAC/SLO
- **RBAC:** CLIENT/FREELANCER (propose/approve/reject/withdraw), SYSTEM (apply), PUBLIC (view own)
- **SLO:** P95 < 250ms (propose), P95 < 200ms (approve/reject), P95 < 300ms (apply), P95 < 150ms (read)

---

### 8.2 extension/

#### User Stories
- As a **client/freelancer**, I want to **request contract extensions** so that more time is available.
- As a **client/freelancer**, I want to **approve extensions** so that timeline is updated.
- As a **system**, I want to **validate extension against budget** so that funds are sufficient.
- As a **system**, I want to **notify about expiring contracts** so that extensions can be requested proactively.

#### Flow
1. **RequestExtensionCommand**(contract_id, new_end_date, justification, requested_by) → ValidateContract() | CalculateAdditionalCost() | Persist() | NotifyCounterparty(communications-be) → **Outbox:** extension.requested.v1
2. **ApproveExtensionCommand**(extension_id, approved_by) → ValidateCounterparty() | CheckBilateralApproval() | ApplyExtension() | AdjustEscrow(financial-be) → **Outbox:** extension.approved.v1
3. **RejectExtensionCommand**(extension_id, reason, rejected_by) → ValidateCounterparty() | Reject() → **Outbox:** extension.rejected.v1
4. **GetExtensionQuery**(extension_id) → AuthorizeAccess() | Fetch() → ExtensionDTO
5. **GetExpiringContractsQuery**(days_until_expiry) → Filter() → ContractListDTO

#### Projections
- extension_read
- expiring_contracts_read

#### Events Published
- extension.requested.v1
- extension.approved.v1
- extension.rejected.v1
- contract.expiring.soon.v1

#### RBAC/SLO
- **RBAC:** CLIENT/FREELANCER (request/approve/reject), PUBLIC (view own)
- **SLO:** P95 < 220ms (request), P95 < 250ms (approve), P95 < 180ms (reject)

---

## **9 - DISPUTE DOMAIN**

### 9.1 dispute/

#### User Stories
- As a **client/freelancer**, I want to **open disputes** for milestone/timesheet disagreements so that resolution process starts.
- As a **client/freelancer**, I want to **provide evidence** (files, screenshots, messages) so that case is supported.
- As a **mediator**, I want to **review dispute details** so that fair resolution is possible.
- As a **mediator**, I want to **request additional information** from parties so that context is complete.
- As a **admin**, I want to **resolve disputes** with binding decisions so that conflicts end.
- As a **system**, I want to **escalate unresolved disputes** after N days so that resolution isn't delayed.
- As a **system**, I want to **place financial holds** during disputes so that funds are protected.

#### Flow
1. **OpenDisputeCommand**(contract_id, dispute_type, subject_id, reason, evidence_urls[], opened_by) → ValidateContract() | PlaceFinancialHold() | CreateDisputeCase(admin-be) | Persist() | NotifyCounterparty(communications-be) → **Outbox:** dispute.opened.v1
2. **AddEvidenceCommand**(dispute_id, evidence_type, evidence_urls[], description, added_by) → AuthorizeParty() | Upload(storage-be) | Attach() → **Outbox:** dispute.evidence.added.v1
3. **RespondToDisputeCommand**(dispute_id, response, counter_evidence[], responded_by) → AuthorizeCounterparty() | Respond() → **Outbox:** dispute.responded.v1
4. **RequestInfoCommand**(dispute_id, info_requested, requested_by) → AuthorizeMediator() | Request() | NotifyParty(communications-be) → **Outbox:** dispute.info.requested.v1
5. **ResolveDisputeCommand**(dispute_id, resolution_type, decision, refund_amount, resolved_by) → AuthorizeAdmin() | ApplyResolution() | ReleaseHold() | ProcessRefunds(financial-be) | Resolve() → **Outbox:** dispute.resolved.v1
6. **EscalateDisputeCommand**(dispute_id, escalation_reason, escalated_by) → ValidateDuration() | Escalate() | NotifyAdmins(communications-be) → **Outbox:** dispute.escalated.v1
7. **GetDisputeQuery**(dispute_id) → AuthorizeAccess() | Fetch() → DisputeDTO
8. **ListDisputesQuery**(contract_id, filters) → ApplyFilters() → DisputeListDTO
9. **GetDisputeTimelineQuery**(dispute_id) → FetchEvents() → DisputeTimelineDTO

#### Projections
- dispute_read
- dispute_evidence_read
- dispute_timeline_read
- dispute_stats_read

#### Events Published
- dispute.opened.v1
- dispute.evidence.added.v1
- dispute.responded.v1
- dispute.info.requested.v1
- dispute.resolved.v1
- dispute.escalated.v1
- dispute.closed.v1

#### Events Consumed
- milestone.disputed.v1 (auto-create dispute)
- timesheet.disputed.v1 (auto-create dispute)

#### RBAC/SLO
- **RBAC:** CLIENT/FREELANCER (open/add evidence/respond), MEDIATOR (request info), ADMIN (resolve/escalate), PUBLIC (view own)
- **SLO:** P95 < 300ms (open), P95 < 200ms (add evidence/respond), P95 < 250ms (resolve), P95 < 150ms (read)

---

## **10 - BUDGET TRACKING DOMAIN**

### 10.1 budget/

#### User Stories
- As a **client**, I want to **set contract budget** (total amount, warning threshold) so that spending is controlled.
- As a **system**, I want to **track budget consumption** in real-time so that overruns are detected.
- As a **client**, I want to **receive alerts** when budget reaches threshold so that action can be taken.
- As a **client**, I want to **increase budget** when more work is needed so that contract continues.
- As a **system**, I want to **prevent milestone creation** when budget exhausted so that overspending is blocked.
- As a **system**, I want to **track budget vs actual** so that variance is visible.

#### Flow
1. **SetBudgetCommand**(contract_id, total_budget, warning_threshold_pct, currency, set_by) → ValidateContract() | ReserveEscrow(financial-be) | Persist() → **Outbox:** budget.set.v1
2. **UpdateBudgetCommand**(contract_id, new_total_budget, reason, updated_by) → AuthorizeClient() | ValidateIncrease() | AdjustEscrow(financial-be) | Update() → **Outbox:** budget.updated.v1
3. **TrackBudgetConsumptionCommand**(contract_id, amount_consumed, source_type, source_id) → UpdateConsumption() | CheckThresholds() | TriggerAlerts(communications-be) → **Outbox:** budget.consumption.tracked.v1
4. **GetBudgetStatusQuery**(contract_id) → Calculate() → BudgetStatusDTO
5. **GetBudgetHistoryQuery**(contract_id) → Fetch() → BudgetHistoryDTO
6. **GetBudgetVarianceQuery**(contract_id) → CalculateVariance() → BudgetVarianceDTO

#### Projections
- budget_read
- budget_consumption_read
- budget_alerts_read

#### Events Published
- budget.set.v1
- budget.updated.v1
- budget.consumption.tracked.v1
- budget.threshold.reached.v1
- budget.exhausted.v1

#### Events Consumed
- milestone.created.v1 (reserve budget)
- milestone.payment.released.v1 (consume budget)
- timesheet.approved.v1 (consume budget)

#### RBAC/SLO
- **RBAC:** CLIENT (set/update), SYSTEM (track), PUBLIC (view own)
- **SLO:** P95 < 200ms (set/update), P95 < 150ms (track), P95 < 100ms (read)

---

## **11 - REMINDERS & NOTIFICATIONS DOMAIN**

### 11.1 reminder/

#### User Stories
- As a **system**, I want to **send deadline reminders** (milestone due, contract expiring) so that parties are notified.
- As a **freelancer**, I want to **receive timesheet submission reminders** weekly so that I don't forget.
- As a **client**, I want to **receive approval reminders** for pending timesheets/milestones so that delays are minimized.
- As a **system**, I want to **escalate overdue approvals** to admins so that bottlenecks are cleared.
- As a **user**, I want to **customize reminder preferences** so that notifications suit my workflow.

#### Flow
1. **ScheduleReminderCommand**(contract_id, reminder_type, trigger_condition, recipient_ids[], scheduled_by) → ValidateContract() | Schedule() → **Outbox:** reminder.scheduled.v1
2. **SendReminderCommand**(reminder_id) → ValidateTrigger() | Send(communications-be) | LogDelivery() → **Outbox:** reminder.sent.v1
3. **CancelReminderCommand**(reminder_id, cancelled_by) → Cancel() → **Outbox:** reminder.cancelled.v1
4. **UpdateReminderPreferencesCommand**(user_id, preferences) → AuthorizeOwner() | Update() → **Outbox:** reminder.preferences.updated.v1
5. **GetPendingRemindersQuery**(contract_id) → Fetch() → ReminderListDTO
6. **GetReminderHistoryQuery**(contract_id) → Fetch() → ReminderHistoryDTO

#### Projections
- reminder_read
- reminder_history_read
- reminder_preferences_read

#### Events Published
- reminder.scheduled.v1
- reminder.sent.v1
- reminder.cancelled.v1
- reminder.preferences.updated.v1

#### Events Consumed
- milestone.submitted.v1 (schedule approval reminder)
- timesheet.submitted.v1 (schedule approval reminder)
- contract.created.v1 (schedule deadline reminders)

#### RBAC/SLO
- **RBAC:** SYSTEM (schedule/send), USER (update preferences/cancel), PUBLIC (view own)
- **SLO:** P95 < 180ms (schedule), P95 < 150ms (send), P95 < 100ms (preferences)

---

## **12 - CONTRACT HISTORY & AUDIT DOMAIN**

### 12.1 audit/

#### User Stories
- As a **compliance officer**, I want to **view complete audit trail** of contract changes so that compliance is verified.
- As a **client/freelancer**, I want to **view contract history** so that evolution is tracked.
- As a **system**, I want to **log all contract actions** with actor, timestamp, and changes so that accountability is maintained.
- As a **admin**, I want to **export audit logs** for legal purposes so that records are preserved.
- As a **system**, I want to **detect suspicious patterns** in audit logs so that fraud is identified.

#### Flow
1. **LogAuditEventCommand**(contract_id, action_type, actor_id, changes_json, ip_address, user_agent) → Persist() | DetectAnomalies() → **Outbox:** audit.event.logged.v1
2. **GetAuditTrailQuery**(contract_id, filters, pagination) → AuthorizeAccess() | Fetch() → AuditTrailDTO
3. **ExportAuditLogsCommand**(contract_id, export_format, exported_by) → AuthorizeAdmin() | Export() | Upload(storage-be) → **Outbox:** audit.logs.exported.v1
4. **DetectAuditAnomaliesQuery**(contract_id, time_range) → AnalyzePatterns() → AnomalyReportDTO
5. **GetContractVersionsQuery**(contract_id) → Fetch() → ContractVersionListDTO

#### Projections
- audit_trail_read
- contract_versions_read
- audit_anomalies_read

#### Events Published
- audit.event.logged.v1
- audit.logs.exported.v1
- audit.anomaly.detected.v1

#### RBAC/SLO
- **RBAC:** SYSTEM (log), CLIENT/FREELANCER (view own trail), ADMIN (export/view all), COMPLIANCE (anomalies)
- **SLO:** P95 < 100ms (log), P95 < 200ms (query), P95 < 500ms (export)

---

## **13 - DIRECT CONTRACTS & INVITATIONS DOMAIN**

### 13.1 direct_contract/

#### User Stories
- As a **client**, I want to **invite freelancers directly** without proposal process so that trusted talent is hired quickly.
- As a **client**, I want to **send contract invitation** with pre-filled terms so that negotiation is minimal.
- As a **freelancer**, I want to **accept invitation** if terms are agreeable so that contract is created.
- As a **freelancer**, I want to **decline invitation** with optional reason so that interest is communicated.
- As a **client**, I want to **revoke pending invitations** so that cancelled opportunities don't linger.
- As a **system**, I want to **track invitation expiry** (7 days default) so that stale invites are closed.
- As a **system**, I want to **generate secure invitation tokens** so that only intended recipients can accept.

#### Flow
1. **InviteFreelancerCommand**(client_id, freelancer_id, job_id, contract_terms, invitation_message, expires_at, invited_by) → ValidateFreelancer() | GenerateToken() | Persist() | SendInvitation(communications-be) → **Outbox:** invitation.sent.v1
2. **AcceptInvitationCommand**(invitation_id, invitation_token, accepted_by) → ValidateToken() | ValidateExpiry() | CreateDirectContract() | Activate() → **Outbox:** invitation.accepted.v1
3. **DeclineInvitationCommand**(invitation_id, decline_reason, declined_by) → ValidateToken() | Decline() → **Outbox:** invitation.declined.v1
4. **RevokeInvitationCommand**(invitation_id, reason, revoked_by) → AuthorizeClient() | Revoke() | NotifyFreelancer(communications-be) → **Outbox:** invitation.revoked.v1
5. **ResendInvitationCommand**(invitation_id, resent_by) → ValidatePending() | ExtendExpiry() | Send(communications-be) → **Outbox:** invitation.resent.v1
6. **GetInvitationQuery**(invitation_id) → AuthorizeAccess() | Fetch() → InvitationDTO
7. **ListInvitationsQuery**(filters, pagination) → ApplyFilters() → InvitationListDTO
8. **GetPendingInvitationsQuery**(freelancer_id) → Fetch() → InvitationListDTO

#### Projections
- invitation_read
- invitation_status_read
- direct_contract_read

#### Events Published
- invitation.sent.v1
- invitation.accepted.v1
- invitation.declined.v1
- invitation.revoked.v1
- invitation.expired.v1
- invitation.resent.v1

#### Events Consumed
- user.suspended.v1 (auto-revoke invitations)

#### RBAC/SLO
- **RBAC:** CLIENT (invite/revoke/resend), FREELANCER (accept/decline/view own), PUBLIC (view own)
- **SLO:** P95 < 250ms (invite), P95 < 300ms (accept - includes contract creation), P95 < 180ms (decline/revoke), P95 < 120ms (read)

---

## **14 - RATE CARDS & PRICING DOMAIN**

### 14.1 rate_card/

#### User Stories
- As a **freelancer**, I want to **create rate cards** (hourly rates by skill level) so that pricing is standardized.
- As a **freelancer**, I want to **apply rate card** to contract so that rates are pre-filled.
- As a **freelancer**, I want to **update rate cards** so that pricing evolves.
- As a **system**, I want to **support tiered rates** (junior/mid/senior) so that skill levels are reflected.
- As a **client**, I want to **view freelancer rate cards** so that cost expectations are clear.

#### Flow
1. **CreateRateCardCommand**(freelancer_id, name, rate_tiers[], currency, created_by) → ValidateRates() | Persist() → **Outbox:** rate_card.created.v1
2. **UpdateRateCardCommand**(rate_card_id, updates, updated_by) → AuthorizeOwner() | ValidateRates() | Apply() → **Outbox:** rate_card.updated.v1
3. **ApplyRateCardCommand**(contract_id, rate_card_id, tier_selection, applied_by) → ValidateContract() | ApplyRates() → **Outbox:** rate_card.applied.v1
4. **GetRateCardQuery**(rate_card_id) → Fetch() → RateCardDTO
5. **ListRateCardsQuery**(freelancer_id) → Fetch() → RateCardListDTO

#### Projections
- rate_card_read
- rate_card_usage_read

#### Events Published
- rate_card.created.v1
- rate_card.updated.v1
- rate_card.applied.v1
- rate_card.deleted.v1

#### RBAC/SLO
- **RBAC:** FREELANCER (create/update/delete), CLIENT (view), PUBLIC (view published)
- **SLO:** P95 < 180ms (create/update), P95 < 150ms (apply), P95 < 100ms (read)

---

## **15 - CONTRACT ANALYTICS DOMAIN**

### 15.1 analytics/

#### User Stories
- As a **client**, I want to **view contract performance metrics** (completion rate, avg time-to-complete, budget variance) so that hiring quality is assessed.
- As a **freelancer**, I want to **view earnings analytics** (total earned, avg project value, payment timeline) so that business health is tracked.
- As a **platform admin**, I want to **aggregate contract statistics** (active contracts, total value, dispute rate) so that platform health is monitored.
- As a **system**, I want to **generate contract reports** (monthly summaries, tax documents) so that record-keeping is automated.
- As a **system**, I want to **track SLA metrics** (milestone approval time, payment release time) so that service quality is measured.

#### Flow
1. **UpdateContractMetricsCommand**(contract_id, metrics_update) → Aggregate() | Update() → **Outbox:** contract.metrics.updated.v1
2. **GenerateContractReportCommand**(contract_id, report_type, date_range, generated_by) → AggregateData() | GenerateReport() | Upload(storage-be) → **Outbox:** contract.report.generated.v1
3. **GetContractPerformanceQuery**(contract_id) → Aggregate() → ContractPerformanceDTO
4. **GetEarningsAnalyticsQuery**(freelancer_id, date_range) → Aggregate() → EarningsAnalyticsDTO
5. **GetPlatformStatsQuery**(date_range, filters) → AuthorizeAdmin() | Aggregate() → PlatformStatsDTO
6. **GetSLAMetricsQuery**(contract_id) → Calculate() → SLAMetricsDTO

#### Projections
- contract_analytics_read
- earnings_analytics_read
- platform_stats_read
- sla_metrics_read

#### Events Published
- contract.metrics.updated.v1
- contract.report.generated.v1
- analytics.aggregation.completed.v1

#### Events Consumed
- contract.completed.v1 (update metrics)
- milestone.payment.released.v1 (update earnings)
- dispute.resolved.v1 (update dispute rate)

#### RBAC/SLO
- **RBAC:** CLIENT/FREELANCER (view own analytics), ADMIN (platform stats), SYSTEM (update metrics)
- **SLO:** P95 < 300ms (generate report), P95 < 200ms (query analytics), P95 < 150ms (metrics update)

---

## **16 - COMPLIANCE & VERIFICATION DOMAIN**

### 16.1 compliance/

#### User Stories
- As a **compliance officer**, I want to **verify contract compliance** with labor laws so that legal risk is minimized.
- As a **system**, I want to **check sanctions lists** for parties so that prohibited entities are blocked.
- As a **system**, I want to **validate tax requirements** (W9/W8, 1099 thresholds) so that reporting is complete.
- As a **admin**, I want to **flag non-compliant contracts** so that remediation is triggered.
- As a **system**, I want to **generate 1099 forms** annually for US freelancers so that tax reporting is automated.
- As a **system**, I want to **enforce geographic restrictions** (e.g., OFAC) so that compliance is maintained.

#### Flow
1. **CheckComplianceCommand**(contract_id, compliance_checks[]) → RunChecks() | FlagIssues() | Persist() → **Outbox:** compliance.check.completed.v1
2. **VerifySanctionsCommand**(user_id, contract_id) → CheckSanctionsLists() | ValidateGeography() → **Outbox:** sanctions.check.completed.v1
3. **ValidateTaxRequirementsCommand**(contract_id, tax_year) → CheckThresholds() | ValidateForms() | FlagMissing() → **Outbox:** tax.requirements.validated.v1
4. **FlagComplianceIssueCommand**(contract_id, issue_type, severity, details, flagged_by) → Flag() | NotifyParties(communications-be) | EscalateIfCritical() → **Outbox:** compliance.issue.flagged.v1
5. **Generate1099Command**(freelancer_id, tax_year) → AggregateEarnings() | GenerateForm() | Upload(storage-be) | Send(communications-be) → **Outbox:** tax.1099.generated.v1
6. **GetComplianceStatusQuery**(contract_id) → Fetch() → ComplianceStatusDTO
7. **ListNonCompliantContractsQuery**(filters) → AuthorizeAdmin() | Filter() → ContractListDTO

#### Projections
- compliance_status_read
- sanctions_check_read
- tax_requirements_read
- compliance_issues_read

#### Events Published
- compliance.check.completed.v1
- sanctions.check.completed.v1
- tax.requirements.validated.v1
- compliance.issue.flagged.v1
- tax.1099.generated.v1

#### Events Consumed
- contract.created.v1 (trigger compliance checks)
- payment.threshold.reached.v1 (trigger tax validation)

#### RBAC/SLO
- **RBAC:** SYSTEM (check/validate/generate), COMPLIANCE (flag/view all), CLIENT/FREELANCER (view own status)
- **SLO:** P95 < 500ms (compliance checks), P95 < 300ms (sanctions check), P95 < 200ms (tax validation)

---

## **17 - FEEDBACK & REVIEWS DOMAIN**

### 17.1 contract_feedback/

#### User Stories
- As a **client**, I want to **leave feedback** on freelancer performance so that future clients benefit.
- As a **freelancer**, I want to **leave feedback** on client behavior so that reputation is tracked.
- As a **system**, I want to **trigger review prompts** after contract completion so that feedback is collected.
- As a **system**, I want to **link feedback to reviews-be** so that public ratings are updated.
- As a **user**, I want to **view feedback received** so that improvement areas are identified.
- As a **system**, I want to **validate feedback authenticity** (verified contracts only) so that trust is maintained.

#### Flow
1. **SubmitFeedbackCommand**(contract_id, rating, review_text, feedback_type, submitted_by) → ValidateContractCompleted() | Persist() | TriggerReview(reviews-be) → **Outbox:** feedback.submitted.v1
2. **RespondToFeedbackCommand**(feedback_id, response_text, responded_by) → AuthorizeRecipient() | Respond() → **Outbox:** feedback.responded.v1
3. **FlagFeedbackCommand**(feedback_id, reason, flagged_by) → Flag() | NotifyModerators(admin-be) → **Outbox:** feedback.flagged.v1
4. **GetFeedbackQuery**(feedback_id) → AuthorizeAccess() | Fetch() → FeedbackDTO
5. **ListFeedbackQuery**(user_id, filters) → Fetch() → FeedbackListDTO
6. **GetFeedbackStatsQuery**(user_id) → Aggregate() → FeedbackStatsDTO

#### Projections
- contract_feedback_read
- feedback_stats_read

#### Events Published
- feedback.submitted.v1
- feedback.responded.v1
- feedback.flagged.v1

#### Events Consumed
- contract.completed.v1 (trigger feedback prompt)
- review.submitted.v1 (link feedback)

#### RBAC/SLO
- **RBAC:** CLIENT/FREELANCER (submit/respond on own contracts), PUBLIC (view public feedback), MODERATOR (flag)
- **SLO:** P95 < 220ms (submit), P95 < 180ms (respond), P95 < 150ms (read)

---

## **18 - CONTRACT SEARCH & DISCOVERY DOMAIN**

### 18.1 search/

#### User Stories
- As a **client**, I want to **search my contracts** by status, freelancer, date range so that finding is easy.
- As a **freelancer**, I want to **filter contracts** by client, earnings, type so that tracking is simple.
- As a **admin**, I want to **advanced search** across all contracts so that investigations are possible.
- As a **system**, I want to **index contracts in search-be** so that discovery is fast.
- As a **system**, I want to **emit search signals** (hygiene, activity) so that search ranking is informed.

#### Flow
1. **IndexContractCommand**(contract_id) → EnrichData() | Send(search-be) → **Outbox:** contract.indexed.v1
2. **RemoveContractFromIndexCommand**(contract_id) → Send(search-be) → **Outbox:** contract.index.removed.v1
3. **SearchContractsQuery**(search_params, filters, pagination) → Query(search-be) | Fetch() → ContractSearchResultsDTO
4. **EmitSearchSignalsCommand**(contract_id, signals[]) → Send(search-be) → **Outbox:** contract.search.signals.emitted.v1

#### Projections
- — (remote in search-be)

#### Events Published
- contract.indexed.v1
- contract.index.removed.v1
- contract.search.signals.emitted.v1

#### Events Consumed
- contract.created.v1 (index)
- contract.updated.v1 (reindex)
- contract.completed.v1 (emit signals)

#### RBAC/SLO
- **RBAC:** CLIENT/FREELANCER (search own), ADMIN (search all), SYSTEM (index/signal)
- **SLO:** P95 < 100ms (index), P95 < 200ms (search), async (signals)

---

## **19 - INSURANCE & GUARANTEES DOMAIN**

### 19.1 insurance/

#### User Stories
- As a **client**, I want to **purchase contract insurance** so that I'm protected from non-delivery.
- As a **freelancer**, I want to **purchase liability insurance** so that I'm protected from disputes.
- As a **system**, I want to **integrate with insurance-be** for policy management so that coverage is tracked.
- As a **system**, I want to **validate insurance requirements** per contract value so that adequate coverage is ensured.
- As a **user**, I want to **file insurance claims** when issues arise so that losses are covered.

#### Flow
1. **PurchaseInsuranceCommand**(contract_id, insurance_type, coverage_amount, purchased_by) → ValidateContract() | CreatePolicy(insurance-be) | Persist() → **Outbox:** insurance.purchased.v1
2. **FileClaimCommand**(insurance_policy_id, claim_type, claim_amount, evidence[], filed_by) → ValidatePolicy() | SubmitClaim(insurance-be) → **Outbox:** insurance.claim.filed.v1
3. **GetInsuranceStatusQuery**(contract_id) → Fetch() → InsuranceStatusDTO
4. **ListPoliciesQuery**(user_id) → Fetch() → InsurancePolicyListDTO

#### Projections
- insurance_policies_read
- insurance_claims_read

#### Events Published
- insurance.purchased.v1
- insurance.claim.filed.v1
- insurance.claim.approved.v1

#### Events Consumed
- insurance.claim.approved.v1 (from insurance-be)
- contract.terminated.v1 (trigger insurance review)

#### RBAC/SLO
- **RBAC:** CLIENT/FREELANCER (purchase/file claim on own contracts), SYSTEM (validate), PUBLIC (view own)
- **SLO:** P95 < 300ms (purchase), P95 < 250ms (file claim), P95 < 150ms (read)

---

## **20 - ESCROW MANAGEMENT DOMAIN**

### 20.1 escrow/

#### User Stories
- As a **client**, I want to **fund escrow** when contract activates so that freelancer has payment assurance.
- As a **system**, I want to **hold funds in escrow** until milestone completion so that both parties are protected.
- As a **system**, I want to **release escrow** upon approval so that freelancer receives payment.
- As a **system**, I want to **refund escrow** on termination so that client recovers unearned funds.
- As a **system**, I want to **integrate with financial-be** for all escrow operations so that financial accuracy is maintained.
- As a **admin**, I want to **manually release escrow** in dispute resolution so that decisions can be enforced.

#### Flow
1. **CreateEscrowCommand**(contract_id, total_amount, currency, created_by) → CreateAccount(financial-be) | Persist() → **Outbox:** escrow.created.v1
2. **FundEscrowCommand**(escrow_id, amount, payment_method_id, funded_by) → ProcessPayment(financial-be) | Fund() → **Outbox:** escrow.funded.v1
3. **HoldEscrowCommand**(escrow_id, milestone_id, hold_amount) → PlaceHold(financial-be) | Reserve() → **Outbox:** escrow.hold.placed.v1
4. **ReleaseEscrowCommand**(escrow_id, milestone_id, release_amount, released_by) → ReleaseHold(financial-be) | ProcessPayout() → **Outbox:** escrow.released.v1
5. **RefundEscrowCommand**(escrow_id, refund_amount, reason, refunded_by) → ProcessRefund(financial-be) | Refund() → **Outbox:** escrow.refunded.v1
6. **GetEscrowStatusQuery**(contract_id) → Fetch(financial-be) → EscrowStatusDTO
7. **GetEscrowHistoryQuery**(escrow_id) → Fetch() → EscrowHistoryDTO

#### Projections
- escrow_status_read
- escrow_history_read

#### Events Published
- escrow.created.v1
- escrow.funded.v1
- escrow.hold.placed.v1
- escrow.released.v1
- escrow.refunded.v1

#### Events Consumed
- milestone.approved.v1 (trigger release)
- contract.terminated.v1 (trigger refund)
- dispute.resolved.v1 (trigger manual release/refund)

#### RBAC/SLO
- **RBAC:** CLIENT (fund), SYSTEM (hold/release/refund), ADMIN (manual operations), PUBLIC (view own)
- **SLO:** P95 < 350ms (fund/release/refund - includes financial-be calls), P95 < 200ms (hold), P95 < 150ms (read)

---

## **21 - NOTIFICATIONS & ALERTS DOMAIN**

### 21.1 notification/

#### User Stories
- As a **user**, I want to **receive contract notifications** (created, activated, paused, completed) so that I stay informed.
- As a **client**, I want to **receive milestone submission alerts** so that approvals aren't delayed.
- As a **freelancer**, I want to **receive payment alerts** so that I know when I'm paid.
- As a **system**, I want to **integrate with communications-be** for all notifications so that delivery is centralized.
- As a **user**, I want to **customize notification preferences** (email, in-app, SMS) so that I control frequency.
- As a **system**, I want to **batch notifications** to avoid spam so that user experience is good.

#### Flow
1. **SendNotificationCommand**(contract_id, notification_type, recipient_ids[], payload, priority) → Enrich() | Send(communications-be) | LogDelivery() → **Outbox:** notification.sent.v1
2. **BatchNotificationsCommand**(notifications[], batch_window) → Group() | SendBatch(communications-be) → **Outbox:** notifications.batched.v1
3. **UpdateNotificationPreferencesCommand**(user_id, preferences) → AuthorizeOwner() | Update() → **Outbox:** notification.preferences.updated.v1
4. **GetNotificationHistoryQuery**(contract_id, filters) → Fetch() → NotificationHistoryDTO

#### Projections
- notification_history_read
- notification_preferences_read

#### Events Published
- notification.sent.v1
- notifications.batched.v1
- notification.preferences.updated.v1
- notification.delivery.failed.v1

#### Events Consumed
- ALL contract events (trigger notifications)

#### RBAC/SLO
- **RBAC:** SYSTEM (send), USER (update preferences/view own)
- **SLO:** P95 < 180ms (send), P95 < 150ms (batch), async delivery

---

## **22 - CONTRACT TERMINATION DOMAIN**

### 22.1 termination/

#### User Stories
- As a **client/freelancer**, I want to **initiate termination** with reason so that contract ends formally.
- As a **system**, I want to **calculate final settlement** (hours worked, milestones completed, refunds) so that payouts are fair.
- As a **system**, I want to **process termination refunds** based on policy so that funds are distributed correctly.
- As a **client/freelancer**, I want to **negotiate termination terms** so that mutual agreement is reached.
- As a **system**, I want to **enforce termination policies** (notice period, kill fees) so that rules are consistent.
- As a **admin**, I want to **force terminate** contracts in violation so that platform integrity is maintained.

#### Flow
1. **InitiateTerminationCommand**(contract_id, termination_reason, termination_type, notice_period, initiated_by) → ValidateContract() | CalculateSettlement() | Persist() | NotifyCounterparty(communications-be) → **Outbox:** termination.initiated.v1
2. **NegotiateTerminationCommand**(termination_id, counter_proposal, negotiated_by) → ValidateParty() | Persist() | NotifyCounterparty() → **Outbox:** termination.negotiated.v1
3. **ApproveTerminationCommand**(termination_id, approved_by) → CheckBilateralApproval() | ProcessFinalPayment(financial-be) | TerminateContract() → **Outbox:** termination.approved.v1
4. **RejectTerminationCommand**(termination_id, rejection_reason, rejected_by) → ValidateCounterparty() | Reject() → **Outbox:** termination.rejected.v1
5. **ForceTerminateCommand**(contract_id, admin_reason, forced_by) → AuthorizeAdmin() | CalculateSettlement() | ProcessRefunds(financial-be) | Terminate() → **Outbox:** contract.force.terminated.v1
6. **GetTerminationDetailsQuery**(termination_id) → AuthorizeAccess() | Fetch() → TerminationDTO
7. **GetSettlementCalculationQuery**(termination_id) → Calculate() → SettlementCalculationDTO

#### Projections
- termination_read
- termination_settlement_read
- termination_negotiations_read

#### Events Published
- termination.initiated.v1
- termination.negotiated.v1
- termination.approved.v1
- termination.rejected.v1
- contract.force.terminated.v1
- termination.settlement.processed.v1

#### RBAC/SLO
- **RBAC:** CLIENT/FREELANCER (initiate/negotiate/approve/reject), ADMIN (force terminate), PUBLIC (view own)
- **SLO:** P95 < 300ms (initiate), P95 < 250ms (negotiate), P95 < 400ms (approve - includes payment), P95 < 180ms (reject)

---

## **23 - COLLABORATION TOOLS DOMAIN**

### 23.1 shared_workspace/

#### User Stories
- As a **client/freelancer**, I want to **access shared workspace** for contract so that collaboration is centralized.
- As a **client/freelancer**, I want to **share documents** in workspace so that materials are accessible.
- As a **client/freelancer**, I want to **comment on documents** so that feedback is tracked.
- As a **system**, I want to **integrate with storage-be** for document management so that files are properly stored.
- As a **system**, I want to **track workspace activity** so that engagement is visible.

#### Flow
1. **CreateWorkspaceCommand**(contract_id, created_by) → ValidateContract() | CreateWorkspace() | SetPermissions() → **Outbox:** workspace.created.v1
2. **ShareDocumentCommand**(workspace_id, document_url, description, shared_by) → ValidateAccess() | Upload(storage-be) | Persist() | NotifyParties(communications-be) → **Outbox:** workspace.document.shared.v1
3. **AddCommentCommand**(document_id, comment_text, commented_by) → ValidateAccess() | Persist() | NotifyOwner(communications-be) → **Outbox:** workspace.comment.added.v1
4. **GetWorkspaceQuery**(contract_id) → AuthorizeAccess() | Fetch() → WorkspaceDTO
5. **GetWorkspaceActivityQuery**(workspace_id) → Fetch() → WorkspaceActivityDTO

#### Projections
- workspace_read
- workspace_documents_read
- workspace_activity_read

#### Events Published
- workspace.created.v1
- workspace.document.shared.v1
- workspace.comment.added.v1
- workspace.document.deleted.v1

#### RBAC/SLO
- **RBAC:** CLIENT/FREELANCER (all operations on own contracts), PUBLIC (view own)
- **SLO:** P95 < 200ms (create/share/comment), P95 < 150ms (read)

---

## **24 - RECURRING CONTRACTS DOMAIN**

### 24.1 recurring/

#### User Stories
- As a **client**, I want to **create recurring contracts** (monthly retainer) so that ongoing work is managed.
- As a **system**, I want to **auto-renew contracts** based on schedule so that manual work is reduced.
- As a **client/freelancer**, I want to **cancel recurring contract** so that future periods are stopped.
- As a **system**, I want to **send renewal reminders** before auto-renewal so that parties can opt out.
- As a **system**, I want to **handle failed renewals** (payment issues) so that contracts are paused.

#### Flow
1. **CreateRecurringContractCommand**(client_id, freelancer_id, recurrence_config, contract_terms, created_by) → ValidateConfig() | CreateContract() | ScheduleRenewal() → **Outbox:** recurring.contract.created.v1
2. **RenewContractCommand**(contract_id) → ValidateActive() | ProcessPayment(financial-be) | CreateNewPeriod() | ScheduleNextRenewal() → **Outbox:** contract.renewed.v1
3. **CancelRecurringCommand**(contract_id, cancellation_reason, cancelled_by) → AuthorizeParty() | CancelFutureRenewals() | CompleteCurrent() → **Outbox:** recurring.cancelled.v1
4. **HandleRenewalFailureCommand**(contract_id, failure_reason) → PauseContract() | NotifyParties(communications-be) | ScheduleRetry() → **Outbox:** renewal.failed.v1
5. **GetRecurringConfigQuery**(contract_id) → Fetch() → RecurringConfigDTO
6. **GetRenewalHistoryQuery**(contract_id) → Fetch() → RenewalHistoryDTO

#### Projections
- recurring_contracts_read
- renewal_schedule_read
- renewal_history_read

#### Events Published
- recurring.contract.created.v1
- contract.renewed.v1
- recurring.cancelled.v1
- renewal.reminder.sent.v1
- renewal.failed.v1

#### Events Consumed
- payment.failed.v1 (trigger renewal failure)
- timer.renewal.due (trigger renewal)

#### RBAC/SLO
- **RBAC:** CLIENT (create/cancel), SYSTEM (renew/handle failure), PUBLIC (view own)
- **SLO:** P95 < 300ms (create), P95 < 400ms (renew - includes payment), P95 < 200ms (cancel)

---

## **25 - CONTRACT SIGNING & E-SIGNATURE DOMAIN**

### 25.1 signature/

#### User Stories
- As a **client/freelancer**, I want to **sign contracts electronically** so that legal validity is established.
- As a **system**, I want to **generate signature documents** (PDF with terms) so that signing is formal.
- As a **system**, I want to **track signature status** (pending, signed, declined) so that completion is monitored.
- As a **system**, I want to **store signed documents** immutably so that legal records are preserved.
- As a **client/freelancer**, I want to **view signature audit trail** so that signing history is transparent.
- As a **system**, I want to **enforce bilateral signing** before activation so that agreement is mutual.

#### Flow
1. **GenerateSignatureDocumentCommand**(contract_id, document_type, generated_by) → RenderPDF() | Upload(storage-be) | Persist() → **Outbox:** signature.document.generated.v1
2. **RequestSignatureCommand**(contract_id, signer_ids[], request_message, requested_by) → ValidateContract() | Send(communications-be) | StartTimer() → **Outbox:** signature.requested.v1
3. **SignDocumentCommand**(signature_request_id, signature_data, ip_address, signed_by) → ValidateRequest() | CaptureSignature() | Timestamp() | CheckBilateralSigning() → **Outbox:** document.signed.v1
4. **DeclineSignatureCommand**(signature_request_id, decline_reason, declined_by) → ValidateRequest() | Decline() | NotifyCounterparty(communications-be) → **Outbox:** signature.declined.v1
5. **GetSignatureStatusQuery**(contract_id) → Fetch() → SignatureStatusDTO
6. **GetSignatureAuditTrailQuery**(contract_id) → Fetch() → SignatureAuditTrailDTO
7. **DownloadSignedDocumentQuery**(contract_id) → AuthorizeAccess() | GeneratePresignedURL(storage-be) → SignedDocumentURLDTO

#### Projections
- signature_requests_read
- signature_status_read
- signature_audit_trail_read
- signed_documents_read

#### Events Published
- signature.document.generated.v1
- signature.requested.v1
- document.signed.v1
- signature.declined.v1
- contract.fully.signed.v1

#### RBAC/SLO
- **RBAC:** CLIENT/FREELANCER (sign/decline on own contracts), SYSTEM (generate/request), PUBLIC (view/download own)
- **SLO:** P95 < 300ms (generate), P95 < 200ms (request/sign/decline), P95 < 150ms (read)

---

## **26 - INBOX: EVENT CONSUMERS**

### 26.1 Proposal Events → Contract Creation

#### User Stories
- As a **system**, I want to **consume proposal.accepted events** so that contract creation is triggered automatically.
- As a **system**, I want to **consume proposal.updated events** so that pre-contract terms stay in sync.

#### Flow
- Consume: `proposal.accepted.v1` → Trigger `CreateContractFromProposalCommand`
- Consume: `proposal.updated.v1` → Update pre-contract draft if exists

#### Projections
- contract_draft_read

#### Events Consumed
- proposal.accepted.v1
- proposal.updated.v1
- proposal.rejected.v1

#### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 200ms (processing lag)

---

### 26.2 Payment Events → Financial Tracking

#### User Stories
- As a **system**, I want to **consume payment.processed events** so that contract financials are updated.
- As a **system**, I want to **consume payment.failed events** so that payment issues are handled.

#### Flow
- Consume: `payment.processed.v1` → Update contract financial status
- Consume: `payment.failed.v1` → Notify parties, retry logic
- Consume: `refund.processed.v1` → Update financial records

#### Projections
- contract_financials_read
- payment_history_read

#### Events Consumed
- payment.processed.v1
- payment.failed.v1
- refund.processed.v1
- payout.completed.v1

#### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 150ms

---

### 26.3 User Events → Contract Management

#### User Stories
- As a **system**, I want to **consume user.suspended events** so that contracts are auto-paused.
- As a **system**, I want to **consume user.banned events** so that contracts are terminated.

#### Flow
- Consume: `user.suspended.v1` → Auto-pause active contracts
- Consume: `user.banned.v1` → Auto-terminate contracts with refunds
- Consume: `user.reactivated.v1` → Resume paused contracts

#### Projections
- contract_status_read

#### Events Consumed
- user.suspended.v1
- user.banned.v1
- user.reactivated.v1

#### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 180ms

---

### 26.4 Job Events → Contract Context

#### User Stories
- As a **system**, I want to **consume job.closed events** so that related contracts are marked complete.
- As a **system**, I want to **consume job.updated events** so that contract context stays fresh.

#### Flow
- Consume: `job.closed.v1` → Mark contracts as complete if applicable
- Consume: `job.cancelled.v1` → Handle contract cancellation

#### Projections
- contract_job_context_read

#### Events Consumed
- job.closed.v1
- job.cancelled.v1
- job.updated.v1

#### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 150ms

---

### 26.5 Admin Events → Platform Operations

#### User Stories
- As a **system**, I want to **consume admin.config.updated events** so that contract rules are refreshed.
- As a **system**, I want to **consume admin.feature_flag.updated events** so that feature availability is current.

#### Flow
- Consume: `admin.config.updated.v1` → Refresh config cache
- Consume: `admin.feature_flag.updated.v1` → Refresh feature flags
- Consume: `admin.moderation.action.applied.v1` → Handle moderation actions

#### Projections
- service_config_read
- feature_flags_read

#### Events Consumed
- admin.config.updated.v1
- admin.feature_flag.updated.v1
- admin.moderation.action.applied.v1

#### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 120ms

---

### 26.6 Financial-be Events → Escrow Management

#### User Stories
- As a **system**, I want to **consume escrow.funded events** so that contract activation is enabled.
- As a **system**, I want to **consume escrow.released events** so that milestone completion is tracked.

#### Flow
- Consume: `escrow.funded.v1` → Enable contract activation
- Consume: `escrow.released.v1` → Mark milestone as paid
- Consume: `chargeback.created.v1` → Place financial hold

#### Projections
- escrow_status_read

#### Events Consumed
- escrow.funded.v1
- escrow.released.v1
- chargeback.created.v1
- chargeback.resolved.v1

#### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 150ms

---

### 26.7 Dispute Events → Dispute Handling

#### User Stories
- As a **system**, I want to **consume dispute.opened events** so that contracts are paused.
- As a **system**, I want to **consume dispute.resolved events** so that holds are released.

#### Flow
- Consume: `dispute.opened.v1` → Pause contract, place hold
- Consume: `dispute.resolved.v1` → Resume contract, release hold
- Consume: `dispute.escalated.v1` → Notify admins

#### Projections
- contract_dispute_status_read

#### Events Consumed
- dispute.opened.v1
- dispute.resolved.v1
- dispute.escalated.v1

#### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 150ms

---

### 26.8 Risk Events → Risk Management

#### User Stories
- As a **system**, I want to **consume risk.alert events** so that financial holds are placed.
- As a **system**, I want to **consume fraud.detected events** so that contracts are flagged.

#### Flow
- Consume: `risk.alert.triggered.v1` → Place financial hold
- Consume: `fraud.detected.v1` → Flag contract, notify compliance

#### Projections
- contract_risk_status_read

#### Events Consumed
- risk.alert.triggered.v1
- fraud.detected.v1
- risk.score.updated.v1

#### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 150ms

---

## **END OF CONTRACTS-BE USER STORIES**

**Total Domains Covered:** 26  
**Total Sections:** 50+  
**Total User Stories:** 400+  
**Total Flows:** 300+  
**Total Events:** 250+  

All stories follow the pattern: **Stories → Flow → Projections → Events → RBAC/SLO**  
All flows include: **idempotent write-path, event envelope, non-PII payloads**  
All components align with: **folder structure, events catalog, platform conventions**

### Summary of Coverage

✅ **Core Contract Domain** - Contract lifecycle, types, CRUD  
✅ **Statement of Work** - SOW creation, versioning, approval  
✅ **Financial Holds** - Risk management, fraud prevention  
✅ **Milestones** - Creation, submission, approval, payment release  
✅ **Deliverables** - Submission, review, revisions  
✅ **Time Tracking** - Timesheets, work diary, hourly tracking  
✅ **Templates** - Contract templates, standard terms  
✅ **Contract Changes** - Amendments, extensions  
✅ **Disputes** - Dispute management, resolution, escalation  
✅ **Budget Tracking** - Budget management, alerts, variance  
✅ **Reminders** - Deadline reminders, approval reminders  
✅ **Audit** - Audit trail, compliance, history  
✅ **Direct Contracts** - Invitations, direct hiring  
✅ **Rate Cards** - Pricing templates, tiered rates  
✅ **Analytics** - Performance metrics, reports, SLA tracking  
✅ **Compliance** - Legal compliance, tax requirements, sanctions  
✅ **Feedback** - Contract feedback, reviews integration  
✅ **Search** - Contract discovery, indexing  
✅ **Insurance** - Contract insurance, claims  
✅ **Escrow** - Escrow management, funds protection  
✅ **Notifications** - Alert system, preferences  
✅ **Termination** - Contract termination, settlement  
✅ **Collaboration** - Shared workspace, document sharing  
✅ **Recurring** - Subscription contracts, auto-renewal  
✅ **E-Signature** - Digital signing, document management  
✅ **Event Consumers** - Integration with all platform services