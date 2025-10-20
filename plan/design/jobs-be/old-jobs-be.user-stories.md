Jobs-be User Stories
====================

Event conventions (applies to all) — jobs-be
--------------------------------------------

*   **Format:** aggregate.resource.action.past\_tense.v1 (e.g., job.lifecycle.published.v1)
    
*   **Envelope includes:** event\_id, event\_ts, aggregate\_id, partition\_key=job\_id, correlation\_id, causation\_id, actor{ id, role }, user\_context{ ip, ua }, data\_zone(EU|US), schema\_ref, compliance\_context{ pii\_flags }
    
*   **Batch ops:** Per-entity events + one \*.summary.v1
    
*   **PII:** Emit hashes/storage\_ids/refs only (no raw PII—e.g., no plaintext emails, phone numbers, or file contents/names)
    

Write-path defaults — jobs-be
-----------------------------

*   **Idempotency:** Header Idempotency-Key (or envelope key). Safe retries must return 200 with no duplicate events.
    
*   **Transactions:** DB tx + outbox with (aggregate\_id, event\_type, idempotency\_key) dedupe.
    
*   **Retries/DLQ:** For external calls (search-be, storage-be, financial-be, subscriptions-be, communications-be, contracts-be, proposals-be, admin-be). Exponential backoff; poison messages to DLQ.
    
*   **Projections:** \_read views per domain; metric event\_to\_projector\_lag\_ms tracked.
    
*   **Security/Perf:** RBAC enforced on commands/queries; SLO/SLA and rate limits as specified per endpoint/feature; field-level encryption for sensitive data; secrets never logged; typical **write P95 ≤ 300 ms**, **read P95 ≤ 250 ms** (unless noted).
    

## 1) Job (Core Posting)

### 1.1 Creation & Bootstrap

#### Stories

*   As a **client**, I want to create a job with title/description/type/budget/category/skills so that candidates see a complete brief.
    
*   As a **client**, I want sane defaults (currency, visibility=public, invite quota, auto-close=30d) so that posting is quick.
    
*   As a **client**, I want “start from template” so that repeated postings are fast.
    
*   As a **system**, I want case-insensitive dedupe on (client\_id,title,normalized\_desc\_hash,window=14d) so that repost spam is reduced.
    
*   As a **system**, I want profanity/PII linting on title/desc so that content is safe.
    
*   As an **org admin**, I want plan quotas enforced so that usage stays within limits.
    
*   As a **client on mobile**, I want a simplified/voice-friendly creation flow so that I can post on-the-go.
    
*   As an **enterprise client**, I want to post via SSO/API so that HR systems sync automatically.
    

### Flow

CreateJobCommand(template\_id?) → ValidateJobInput() | LintTitleDesc() | CheckDuplicateWindow() | CheckPlanQuota() | PreloadTemplateDefaults() → CreateJob()UpdateJobCommand(partial) → GetJobByIDQuery → ValidatePatch() | ReLintIfChanged() → UpdateJob()MobileCreateJobCommand → DetectDevice(user\_context) → SimplifiedValidate() → CreateJob()ExternalPostJobCommand(SSO/API) → Auth(Keycloak) → CreateJob()

### Projections

jobs\_read, jobs\_validation\_read, jobs\_lint\_read, jobs\_quota\_read

### Events

job.core.created.v1, job.core.updated.v1, job.core.duplicate.previewed.v1, job.core.validation.failed.v1, job.mobile.created.v1, job.integration.posted.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 300 ms**; duplicate check ≤ 100 ms (cached simhash); 60 req/min/client.

1.2 Lifecycle (Draft/Publish/Unpublish/Close/Extend/Schedule/Cancel)
--------------------------------------------------------------------

### Stories

*   As a **client**, I want to publish/unpublish/close/extend/schedule so that I control visibility.
    
*   As a **system**, I want guardrails (e.g., cannot close if an offer is pending) so that hiring isn’t broken.
    
*   As a **client**, I want to cancel a job with optional connect refunds so that fairness is maintained.
    

### Flow

PublishJob | UnpublishJob | CloseJob | ExtendJob | ScheduleJob | CancelJob → GetJobLifecycleQuery → GuardRails(offers, moderation, finance) → ApplyLifecycle() (+ refunds via proposals-be/financial-be)

### Projections

jobs\_lifecycle\_read, jobs\_audit\_read, job\_cancellations\_read

### Events

job.lifecycle.published.v1, job.lifecycle.unpublished.v1, job.lifecycle.closed.v1, job.lifecycle.extended.v1, job.lifecycle.scheduled.v1, job.lifecycle.guardrail.blocked.v1, job.cancelled.v1

### RBAC/SLO

OWNER/EDITOR (cancel: OWNER); **P95 < 200 ms**; guard rails < 20 ms (cached).

1.3 Retrieval & Listing
-----------------------

### Stories

*   As a **client**, I want to list my jobs by status/date so I can manage them.
    
*   As a **freelancer**, I want public job details with resolved categories/skills so I can assess fit.
    
*   As a **moderator**, I want moderation state included so I can act quickly.
    

### Flow

ListJobsByClientQuery | ListJobsByStatusQuery | GetJobByIDQuery → ListJobsByClient() | ListJobsByStatus() | GetJobByID()

### Projections

jobs\_read (indexed joins for category/skills/moderation)

### Events

— (read-only)

### RBAC/SLO

OWNER/TEAM for private; Public for public listings; **P95 < 150 ms**.

1.4 Metadata: Duration & Experience Level
-----------------------------------------

### Stories

*   As a **client**, I want to set expected duration (short-term/long-term) and required experience level (Entry/Intermediate/Expert) so that freelancers self-assess fit and search rankings improve.
    

### Flow

SetJobDurationCommand | SetExperienceLevelCommand → ValidateEnums() → Persist() → Emit ranking signal (search-be)

### Projections

job\_metadata\_read

### Events

job.metadata.duration.set.v1, job.requirements.experience\_level.set.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 150 ms**.

1.5 Multi-Language Descriptions
-------------------------------

### Stories

*   As a **global client**, I want multi-language descriptions (with primary language) so that international freelancers apply.
    

### Flow

SetJobLanguagesCommand → ValidateLocales() → Persist translations

### Projections

job\_languages\_read

### Events

job.languages.set.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 200 ms**.

2) Category (Taxonomy)
======================

\=======================

2.1 Assign Category/Subcategory
-------------------------------

### Stories

*   As a **client**, I want to assign/replace a job’s category/subcategory so it’s discoverable.
    
*   As a **system**, I want to validate parent-child cycles so taxonomy integrity holds.
    

### Flow

AssignJobCategoryCommand → GetCategoryTreeQuery → ValidateParentChild() → AssignJobCategory()

### Projections

job\_categories\_read

### Events

job.category.assigned.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 200 ms**.

2.2 Taxonomy Sync
-----------------

### Stories

*   As a **system**, I want to sync categories from search-be so we use canonical taxonomy.
    
*   As an **admin**, I want an orphan remap summary so I can review.
    

### Flow

SyncCategoriesFromTaxonomyCommand → GetCategoryDiffQuery → SyncCategories() | RemapOrphans()

### Projections

categories\_read, job\_categories\_read

### Events

job.taxonomy.synced.v1, job.category.remapped.summary.v1

### RBAC/SLO

System; batch 50k < 10 m; DLQ on conflicts.

3) Skill (Dictionary for Jobs)
==============================

\================================

3.1 Attach/Remove Skills
------------------------

### Stories

*   As a **client**, I want to attach/remove skills to a job so requirements are explicit.
    
*   As a **system**, I want to block deprecated skills so we avoid dead terms.
    

### Flow

AttachSkillsToJobCommand | RemoveJobSkillCommand → ListPopularSkillsQuery → ValidateNotDeprecated() → Attach() | Remove()

### Projections

job\_skills\_read

### Events

job.skill.added.v1, job.skill.removed.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 200 ms**.

4) Job Skill (Requirement Levels)
=================================

\===================================

4.1 Must vs Preferred + Level
-----------------------------

### Stories

*   As a **client**, I want to mark skills as must-have/preferred (with optional level) so matching is accurate.
    

### Flow

SetJobSkillRequirementCommand | UpdateJobSkillRequirementCommand → GetJobSkillsQuery → SetRequirement()

### Projections

job\_skill\_requirements\_read

### Events

job.skill.requirement.set.v1, job.skill.requirement.updated.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 200 ms**.

5) Job Questions (Screening)
============================

\==============================

5.1 CRUD (+ Conditional)
------------------------

### Stories

*   As a **client**, I want to add/update/remove screening questions (text/multi-choice/file upload).
    
*   As a **client**, I want conditional logic (if/then) so screening adapts to answers.
    

### Flow

AddJobQuestionCommand | UpdateJobQuestionCommand | RemoveJobQuestionCommand → GetJobScreeningPackQuery → UpsertQuestion() | RemoveQuestion()

### Projections

job\_screening\_read

### Events

job.screening.question.added.v1, job.screening.question.updated.v1, job.screening.question.removed.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 200 ms**.

6) Job Attachments
==================

\====================

6.1 Add/Remove + Safe Media
---------------------------

### Stories

*   As a **client**, I want to link/remove attachments via storage-be.
    
*   As a **system**, I want AV/DLP scans and PII redaction suggestions so content stays safe.
    

### Flow

AddJobAttachmentCommand | RemoveJobAttachmentCommand → ListJobAttachmentsQuery → storage-be upload/unlink → Scan → Record results

### Projections

job\_attachments\_read, attachment\_scan\_read

### Events

job.attachment.added.v1, job.attachment.removed.v1, job.attachment.scanned.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 200 ms**.

7) Job Invitations
==================

\====================

7.1 Send/Respond/Expire
-----------------------

### Stories

*   As a **client**, I want to send invitations to freelancers.
    
*   As a **freelancer**, I want to accept/decline.
    
*   As a **system**, I want auto-expire on deadline and invite rate limits per job/day.
    

### Flow

SendJobInvitationCommand | AcceptJobInvitationCommand | DeclineJobInvitationCommand | ExpireJobInvitationCommand → ListInvitationsForJobQuery → Send() | Respond() | Expire()

### Projections

job\_invitations\_read

### Events

job.invitation.sent.v1, job.invitation.accepted.v1, job.invitation.declined.v1, job.invitation.expired.v1, job.invitation.ratelimit.hit.v1

### RBAC/SLO

OWNER/EDITOR to send; invitee to respond; **P95 < 200 ms**.

8) Job View (Engagement)
========================

\==========================

8.1 Record Views & Sessions
---------------------------

### Stories

*   As a **system**, I want to record each view and session duration, de-botting via UA/IP reputation.
    

### Flow

RecordJobViewCommand | UpdateJobViewSessionCommand → ListJobViewsQuery → RecordView() | UpdateSession()

### Projections

job\_views\_read, job\_view\_quality\_read

### Events

job.view.recorded.v1, job.view.session.updated.v1, job.view.filtered.bot.v1

### RBAC/SLO

Public write-behind; **P95 < 80 ms**.

9) Saved Search (Freelancer Alerts)
===================================

\=====================================

9.1 Save/Update/Delete + Alerts
-------------------------------

### Stories

*   As a **freelancer**, I want to save searches and toggle alerts (email/push/in-app).
    

### Flow

CreateSavedSearchCommand | UpdateSavedSearchCommand | DeleteSavedSearchCommand → ListSavedSearchesQuery → Save() | Update() | Delete()Alerts via scheduler → job.search.alert.sent.v1.

### Projections

saved\_job\_searches\_read

### Events

job.search.saved.v1, job.search.updated.v1, job.search.deleted.v1, job.search.alert.sent.v1

### RBAC/SLO

SELF; **P95 < 120 ms**; min query length=2.

10) Job Flag (Abuse)
====================

\======================

10.1 Submit/Resolve/Dismiss
---------------------------

### Stories

*   As **any user**, I want to flag spam/fraud.
    
*   As a **moderator**, I want to resolve/dismiss with reason codes.
    

### Flow

FlagJobCommand | ResolveJobFlagCommand | DismissJobFlagCommand → ListJobFlagsQuery → Flag() | Resolve() | Dismiss()

### Projections

job\_flags\_read

### Events

job.flag.submitted.v1, job.flag.resolved.v1, job.flag.dismissed.v1

### RBAC/SLO

Any user to flag; Moderator resolves; **P95 < 150 ms**.

11) Job Templates
=================

\===================

11.1 CRUD & Clone (+ Org Approvals)
-----------------------------------

### Stories

*   As a **client**, I want to create/update/archive templates and clone to new job.
    
*   As an **org admin**, I want org-shared templates with approval workflows.
    

### Flow

CreateJobTemplateCommand | UpdateJobTemplateCommand | ArchiveJobTemplateCommand | CloneTemplateToJobCommandShareTemplateToOrgCommand | ApproveTemplateCommand → ValidateOrgAccess() → Persist(approval\_status)

### Projections

job\_templates\_read, org\_templates\_read

### Events

job.template.created.v1, job.template.updated.v1, job.template.archived.v1, job.template.cloned\_to\_job.v1, job.template.org\_shared.v1, job.template.approved.v1

### RBAC/SLO

OWNER/EDITOR; org ops: ORG\_ADMIN; **P95 < 200 ms**.

12) Template Versions
=====================

\=======================

12.1 SemVer & Deprecations
--------------------------

### Stories

*   As a **client**, I want versioned templates with changelogs and deprecations.
    

### Flow

CreateJobTemplateVersionCommand | DeprecateJobTemplateVersionCommand → ListTemplateVersionsQuery | GetLatestTemplateVersionQuery

### Projections

job\_template\_versions\_read

### Events

job.template.version.created.v1, job.template.version.deprecated.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 180 ms**.

13) Screening & Compliance
==========================

\============================

13.1 NDA / Export / Pack
------------------------

### Stories

*   As a **client**, I want NDA, export-control flags, and screening packs.
    

### Flow

ToggleNDACommand | SetExportControlCommand | SetJobScreeningPackCommand → GetJobScreeningPackQuery

### Projections

job\_screening\_read

### Events

job.compliance.nda.required.v1, job.compliance.export\_control.flagged.v1, job.screening.configured.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 250 ms**.

13.2 External Tests
-------------------

### Stories

*   As a **client**, I want to link third-party skills tests and record completion.
    

### Flow

AttachExternalTestCommand | RecordExternalTestResultCommand → GetExternalTestsQuery

### Projections

job\_external\_tests\_read

### Events

job.screening.external\_test.attached.v1, job.screening.external\_test.completed.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 200 ms**.

14) Sourcing Modes
==================

\====================

14.1 Public / Invite-only / Private Link
----------------------------------------

### Stories

*   As a **client**, I want to switch sourcing mode and generate private links, enforcing constraints per mode.
    

### Flow

SetSourcingModeCommand → GetSourcingConfigQuery

### Projections

job\_sourcing\_read

### Events

job.sourcing.mode.set.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 200 ms**.

14.2 Shortlists & Talent Pools
------------------------------

### Stories

*   As a **reviewer**, I want add/remove shortlist entries; attach/detach org talent pools.
    

### Flow

AddToShortlistCommand | RemoveFromShortlistCommand | AttachTalentPoolCommand | DetachTalentPoolCommand → ListShortlistsQuery

### Projections

job\_shortlists\_read

### Events

job.sourcing.shortlist.updated.v1, job.sourcing.pool.attached.v1, job.sourcing.pool.detached.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 200 ms**.

15) Budget Controls
===================

\=====================

15.1 Ranges/Currency/Hourly Caps
--------------------------------

### Stories

*   As a **client**, I want min/max budget, currency, and hourly caps.
    

### Flow

SetBudgetRangeCommand | SetJobCurrencyCommand | SetHourlyCapCommand → GetBudgetControlsQuery

### Projections

job\_budget\_controls\_read

### Events

job.budget.controls.set.v1, job.budget.controls.updated.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 200 ms**.

15.2 FX Rules
-------------

### Stories

*   As a **finance admin**, I want quote vs settlement currency and rounding mode.
    

### Flow

SetFxRuleCommand → GetFxRulesQuery

### Projections

fx\_rules\_read

### Events

job.budget.fx.rule.set.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 150 ms**.

16) Visibility Lifecycle
========================

\==========================

16.1 Schedule & Auto-close
--------------------------

### Stories

*   As a **client**, I want to schedule publish, set auto-close date, and extend.
    

### Flow

SchedulePostingCommand | SetAutoCloseCommand | ExtendPostingCommand → GetVisibilityLifecycleQuery

### Projections

jobs\_lifecycle\_read

### Events

job.visibility.scheduled.v1, job.visibility.autoclosed.v1, job.visibility.extended.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 200 ms**.

17) Eligibility Rules
=====================

\=======================

17.1 Hard Gates & Decisions
---------------------------

### Stories

*   As a **client**, I want geo allow/deny, KYC required, min trust tier, agency allowed; return decision + reasons.
    

### Flow

SetJobEligibilityRulesCommand | EvaluateApplicantCommand → GetJobEligibilityRulesQuery → SetRules() | Evaluate()

### Projections

job\_eligibility\_read, job\_eligibility\_decisions\_read

### Events

job.eligibility.rules.set.v1, job.eligibility.applicant.blocked.v1, job.eligibility.applicant.allowed.v1

### RBAC/SLO

OWNER/EDITOR; eval **≤ 50 ms** (cached).

18) Hiring Team
===============

\=================

18.1 Members & Roles
--------------------

### Stories

*   As a **job owner**, I want to add/change/remove members (OWNER/REVIEWER/INTERVIEWER) but cannot remove the last OWNER.
    

### Flow

AddHiringTeamMemberCommand | ChangeHiringTeamRoleCommand | RemoveHiringTeamMemberCommand → GetHiringTeamQuery

### Projections

job\_hiring\_team\_read

### Events

job.hiring\_team.member.added.v1, job.hiring\_team.member.removed.v1, job.hiring\_team.role.changed.v1

### RBAC/SLO

OWNER manages; **P95 < 200 ms**.

19) Pipeline
============

\==============

19.1 Stages, SLAs, WIP
----------------------

### Stories

*   As a **reviewer**, I want NEW→SHORTLISTED→INTERVIEW→OFFER→HIRED/REJECTED with SLA timers and WIP limits.
    

### Flow

AdvancePipelineStageCommand | RevertPipelineStageCommand | SetPipelineSLACommand | SetPipelineWIPLimitCommand → GetPipelineQuery

### Projections

job\_pipeline\_read, job\_pipeline\_sla\_read

### Events

job.pipeline.stage.advanced.v1, job.pipeline.stage.reverted.v1, job.pipeline.sla.exceeded.v1, job.pipeline.wip.updated.v1

### RBAC/SLO

TEAM; **P95 < 150 ms**.

20) Requirements Matrix
=======================

\=========================

20.1 Weighted Must/Nice + Scoring
---------------------------------

### Stories

*   As a **client**, I want weighted requirements; as a **reviewer**, I want a score preview with reasons.
    

### Flow

SetRequirementsMatrixCommand | UpdateRequirementWeightCommand | ScoreCandidatePreviewQuery → GetRequirementsMatrixQuery

### Projections

job\_requirements\_matrix\_read

### Events

job.requirements.matrix.set.v1, job.requirements.weight.updated.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 200 ms**.

21) Job Analytics
=================

\===================

21.1 KPIs & Alerts + Post-Job Reports & Funnel/Suggestions
----------------------------------------------------------

### Stories

*   As an **analyst**, I want KPIs (views→proposals→interviews→offers→hires) and an engagement score.
    
*   As a **system**, I want low-engagement alerts.
    
*   As a **client**, I want post-job analytics (time-to-hire, applicant quality, cost per hire).
    
*   As a **client**, I want suggested edits when engagement is low.
    

### Flow

UpdateJobStatsCommand | RecomputeJobEngagementScoreCommand → GetJobStatsQueryGeneratePostJobReportQuery → Aggregate(proposals-be, contracts-be)VisualizeFunnelQuery → job\_analytics\_readSuggestJobIterationsQuery → Analyze low-engagement

### Projections

job\_analytics\_read, job\_post\_analytics\_read

### Events

job.analytics.updated.v1, job.analytics.low\_engagement.alert.v1, job.analytics.post\_report.generated.v1

### RBAC/SLO

OWNER/ANALYST; **P95 120–300 ms**.

22) A/B Experiments
===================

\=====================

22.1 Start/Stop + Guardrails
----------------------------

### Stories

*   As a **client**, I want A/B tests on titles/descriptions/Q-packs with allocations summing to 100.
    

### Flow

StartJobExperimentCommand | StopJobExperimentCommand → GetJobExperimentQuery | ListJobExperimentsQuery

### Projections

job\_experiments\_read

### Events

job.experiment.started.v1, job.experiment.stopped.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 150 ms**.

23) Moderation Lifecycle
========================

\==========================

23.1 Apply/Lift States
----------------------

### Stories

*   As **T&S**, I want CLEAN/LIMITED/QUARANTINED with reasons; append-only history.
    

### Flow

ApplyModerationStateCommand | LiftModerationStateCommand → GetModerationStateQuery

### Projections

job\_moderation\_read

### Events

job.moderation.state.applied.v1, job.moderation.state.lifted.v1

### RBAC/SLO

Admin/T&S; **P95 < 200 ms**.

24) Syndication
===============

\=================

24.1 Queue/Post/Fail/Takedown
-----------------------------

### Stories

*   As a **client**, I want partner-board syndication with retries/backoff and takedown.
    

### Flow

QueueJobSyndicationCommand | MarkSyndicationFailedCommand | TakedownSyndicatedJobCommand → GetSyndicationStatusQuery | ListSyndicationPartnersQuery

### Projections

job\_syndication\_read

### Events

job.syndication.queued.v1, job.syndication.posted.v1, job.syndication.failed.v1, job.syndication.takedown.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 200 ms**.

25) Drafts
==========

\============

25.1 Autosave & Restore
-----------------------

### Stories

*   As a **client**, I want autosave snapshots; restore to job; last-editor tracking with conflict detection.
    

### Flow

SaveJobDraftCommand | RestoreDraftToJobCommand → GetDraftQuery | ListDraftsQuery

### Projections

job\_drafts\_read

### Events

job.draft.saved.v1, job.draft.restored.v1, job.draft.conflict.detected.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 120 ms**.

26) Geo Requirements
====================

\======================

26.1 Regions/Radius/TZ Overlap + Mode
-------------------------------------

### Stories

*   As a **client**, I want onsite/remote/hybrid, allowed regions, radius km, TZ overlap hours.
    

### Flow

SetJobGeoRulesCommand → GetJobGeoRulesQuery

### Projections

job\_geo\_read

### Events

job.geo.rules.set.v1, job.geo.rules.updated.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 200 ms**.

26.2 Candidate Check
--------------------

### Stories

*   As a **recruiter**, I want to evaluate a candidate’s location/TZ vs rules (advisory).
    

### Flow

CheckCandidateGeoEligibilityQuery → CheckCandidateGeoEligibility()

### Projections

job\_geo\_checks\_read

### Events

— (read)

### RBAC/SLO

TEAM; **P95 < 80 ms**.

27) Duplicate Detection
=======================

\=========================

27.1 Simhash & Clusters
-----------------------

### Stories

*   As a **system**, I want near-dupes clustered; as an **owner/moderator**, I want repost prevention/merge.
    

### Flow

UpsertJobDuplicateKeyCommand | FindNearDuplicateJobsQuery → UpsertDuplicateKey() | FindNearDuplicates()

### Projections

job\_duplicates\_read

### Events

job.duplicate.detected.v1, job.duplicate.merged.v1, job.duplicate.prevented.v1

### RBAC/SLO

System/OWNER; **P95 < 150 ms**.

28) Interview Hooks
===================

\=====================

28) Interview Hooks
===================

28.1 Policy Flags + Calendar Slots
----------------------------------

### Stories

*   As a **client**, I want interview policy (require panel, preferred slots ref) and linked calendar slots for automated scheduling.
    

### Flow

SetJobInterviewPolicyCommand | LinkCalendarCommand → GetJobInterviewPolicyQuery

### Projections

job\_interview\_policy\_read, job\_calendar\_read

### Events

job.interview.policy.set.v1, job.interview.calendar.linked.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 120–200 ms**.


28.2 Instant Video Interviews (pre-recorded Q&A)
------------------------------------------------

### Stories

*   As a client, I want to set up instant video interviews (pre-recorded questions, time limits) so that screening is faster.
    
*   As a freelancer, I want clear instructions and privacy terms so that I can record confidently.
    
*   As a system, I want to store only non-PII references to media (storage-be ids) so that content remains safe.
    

### Flow

SetVideoInterviewCommand(job\_id, questions\[\], time\_limits, retry\_policy) → ValidateQuestions() → UploadQuestionPrompts(communications-be) → PersistSetup()

UpdateVideoInterviewCommand(job\_id, partial) → ValidatePatch() → Persist()

DisableVideoInterviewCommand(job\_id) → Persist()

### Projections

*   job\_video\_read (setup state, question refs, limits)
    

### Events

*   job.video.interview.set.v1
    
*   job.video.interview.updated.v1
    
*   job.video.interview.disabled.v1
    

### RBAC/SLO

*   OWNER/EDITOR; P95 < 250 ms; media moderation async (DLQ on failures).
    
*   Rate limit: ≤ 10 questions per pack.

29) Legal Controls
==================

\====================

29.1 NDA/Export/Holds + Data Residency
--------------------------------------

### Stories

*   As **legal**, I want NDA template/version pinning and export-control flags.
    
*   As **legal/admin**, I want legal holds with audit and purge blocking.
    
*   As a **compliance officer**, I want job data residency (EU/US).
    

### Flow

SetJobLegalControlsCommand | PlaceJobLegalHoldCommand | RemoveJobLegalHoldCommand | SetDataResidencyCommand → GetJobLegalQuery

### Projections

job\_legal\_read, job\_compliance\_read

### Events

job.legal.controls.set.v1, job.legal.hold.placed.v1, job.legal.hold.removed.v1, job.compliance.residency.set.v1

### RBAC/SLO

Legal/Admin; **P95 < 150–200 ms**.

30) Campaign Tags
=================

\===================

30.1 Add/Remove/Limit
---------------------

### Stories

*   As a **client**, I want internal campaign tags with normalization and limits.
    

### Flow

AddJobCampaignTagCommand | RemoveJobCampaignTagCommand → ListJobCampaignTagsQuery

### Projections

job\_campaign\_tags\_read

### Events

job.campaign.tag.added.v1, job.campaign.tag.removed.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 120 ms**.

31) Retention Rules (Policy)
============================

\==============================

31.1 Archive/Purge Schedules
----------------------------

### Stories

*   As **legal/admin**, I want archive/purge windows; block purge under legal hold.
    

### Flow

SetJobRetentionRulesCommand → GetJobRetentionQuery

### Projections

job\_retention\_read

### Events

job.retention.set.v1, job.retention.archived.v1, job.retention.purged.v1

### RBAC/SLO

Legal/Admin; **P95 < 200 ms**.

32) Job Promotion (Boosts/Featured)
===================================

\=====================================

32.1 Activate/Renew/Suspend/Expire (+ Budget Boosts)
----------------------------------------------------

### Stories

*   As a **client**, I want promotions after payment capture, with renewal/suspend/expire and fee previews.
    

### Flow

ActivateJobPromotionCommand | RenewJobPromotionCommand | SuspendJobPromotionCommand | ExpireJobPromotionCommand → GetJobPromotionQueryApplyBudgetBoostCommand → CheckFees(financial-be) → Persist boost level

### Projections

job\_promotions\_read, job\_boosts\_read

### Events

job.promotion.activated.v1, job.promotion.renewed.v1, job.promotion.suspended.v1, job.promotion.expired.v1, job.budget.boost.applied.v1

### RBAC/SLO

OWNER; **P95 < 150–180 ms**.


32.2 AI Visibility Boosts (dynamic, performance-aware)
------------------------------------------------------

### Stories

*   As a client, I want AI-driven boosts to increase exposure for high-potential jobs so that I reach more qualified talent.
    
*   As a system, I want boosts bounded by plan/entitlement and budget so that costs stay predictable.
    
*   As an analyst, I want attribution (lift vs control) so that ROI is measurable.
    

### Flow

ApplyAIVisibilityBoostCommand(job\_id) → AnalyzeEngagement(job\_analytics\_read) | CheckEntitlements(subscriptions-be) | CheckFinancialHolds(financial-be) → SetBoostLevel(search-be signal) → PersistBoost()

RecomputeBoostsBatchCommand() (SYSTEM) → IterateEligibleJobs() → AdjustBoostLevels()

### Projections

*   job\_visibility\_read (current boost level, source=manual|AI, expiry)
    

### Events

*   job.ai.visibility.boosted.v1
    
*   job.promotion.activated.v1 / job.promotion.renewed.v1 (existing, when paid tiers apply)
    

### RBAC/SLO

*   SYSTEM/OWNER; P95 < 300 ms; AI boosts capped by plan; cooldown 6h between changes.

33) Job Preference (Advisory Matching)
======================================

\========================================

33.1 Set/Update/Remove
----------------------

### Stories

*   As a **client**, I want advisory preferences (freelancer type, preferred locations/TZs, min success/earnings, fluency, guidance, tool provision).
    

### Flow

SetJobPreferencesCommand | UpdateJobPreferencesCommand | RemoveJobPreferencesCommand → GetJobPreferencesQuery

### Projections

job\_preferences\_read

### Events

job.preferences.set.v1, job.preferences.updated.v1, job.preferences.removed.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 150 ms**.

34) Job Hiring Option (Multi-Hire & Repost)
===========================================

\=============================================

34.1 Slots & Hiring Count & Safe Repost
---------------------------------------

### Stories

*   As a **client**, I want multi-hire with open slots and an explicit hiring count (plan-bounded).
    
*   As a **system**, I want repost cooldowns and duplicate checks.
    

### Flow

EnableMultiHireCommand | SetHiringCountCommand | ReserveOpenSlotCommand | ReleaseOpenSlotCommand | RepostJobCommand → GetHiringOptionsQuery

### Projections

job\_hiring\_options\_read

### Events

job.multihire.enabled.v1, job.hiring.count.set.v1, job.slot.reserved.v1, job.slot.released.v1, job.reposted.v1, job.repost.cooldown.blocked.v1

### RBAC/SLO

OWNER; **P95 < 120 ms**.

35) Search (Indexing Integration)
=================================

\===================================

35.1 Index/Remove & Signals
---------------------------

### Stories

*   As a **system**, I want index/remove jobs; emit hygiene/promotion/engagement signals to search-be.
    

### Flow

IndexJobCommand | RemoveJobFromIndexCommand → (search-be)EmitSearchSignalsTask → hygiene/promo/engagement

### Projections

— (remote)

### Events

job.search.indexed.v1, job.search.removed.v1, job.search.hygiene.signal.v1

### RBAC/SLO

System; async; **P95 < 100 ms**.

36) AI Assist
=============

\===============

36) AI Assist
=============

36.1 AI Agent for Job Management (Autonomous refinements)
---------------------------------------------------------

### Stories

*   As a client, I want an AI agent to continuously monitor my job’s performance (views, proposal quality) so that refinements are suggested in real time.
    
*   As a client, I want to approve or auto-apply safe refinements so that my listing improves without constant manual edits.
    
*   As a reviewer (hiring team), I want a digest of suggested changes with reasons and projected impact so that I can make informed approvals.
    
*   As a system, I want to respect guardrails (RBAC, moderation, budget caps, policy) so that the agent never violates platform rules.
    
*   As a compliance officer, I want all agent actions and prompts/audits recorded so that decisions are traceable.
    

### Flow

ActivateAIAgentCommand(job\_id, policy, auto\_apply\_flags) → ValidateJobOwnership() | FetchEligibility(job\_moderation/budget/prefs) → ConfigureAgentRules() → PersistAgentState()

RunAIAgentTickCommand(job\_id) (periodic) → PullSignals(job\_analytics\_read, search-be ML) → GenerateRefinementSuggestions() → If auto\_apply: ValidatePatch() → UpdateJob()/PublishPatch() → RecordSuggestion()

ReviewAIAgentSuggestionCommand(job\_id, suggestion\_id, decision=approve|reject|edit) → ApplyOrDismissSuggestion()

### Projections

*   job\_ai\_agent\_read (agent state, rules, suggestion history, auto-apply flags)
    
*   job\_analytics\_read (as input signal, existing)
    

### Events

*   job.ai.agent.activated.v1
    
*   job.ai.refinement.suggested.v1
    
*   job.ai.refinement.approved.v1 / job.ai.refinement.applied.v1 / job.ai.refinement.rejected.v1
    

### RBAC/SLO

*   OWNER/EDITOR; agent ticks run as SYSTEM with scoped permissions.
    
*   P95 < 350 ms per suggestion generation (ML latency); tick batching allowed.
    
*   Rate limits: max 1 suggestion set per job/15 min; auto-apply behind feature flag.
    

36.2 Automatic Job Post Optimization (on create/update)
-------------------------------------------------------

### Stories

*   As a client, I want AI to rewrite the title/description for clarity and attractiveness so that better talent applies.
    
*   As a client, I want a side-by-side diff and a one-click “Apply” so that I control final copy.
    
*   As a system, I want to enforce policy (no PII/profanity, no policy-blocked claims) before offering the rewrite.
    

### Flow

AutoOptimizeJobCommand(job\_id, trigger=create|update) → RunNLP(search-be ML) → LintContent(PII/profanity) → ProduceDiff() → If auto\_apply enabled: ValidatePatch() → UpdateJob()

AcceptOptimizationCommand(job\_id, optimization\_id) → ApplyDiff()

### Projections

*   job\_optimizations\_read (diffs, scores, approvals)
    

### Events

*   job.auto.optimized.v1
    
*   job.description.optimized.v1 (if manual accept)
    
*   job.ai.suggestions.accepted.v1 (accept subset)
    

### RBAC/SLO

*   OWNER; P95 < 400 ms; guardrails (moderation & compliance) must pass.
    
*   Rate limit: 3 optimizations per job/day.

37) Inclusivity
===============

\=================

37.1 Accessibility Flags
------------------------

### Stories

*   As a **client**, I want accessibility flags (e.g., flexible\_hours, no\_video\_required) so diverse freelancers apply.
    

### Flow

SetAccessibilityFlagsCommand → Persist → Emit soft signal (search-be)

### Projections

job\_inclusivity\_read

### Events

job.inclusivity.flags.set.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 150 ms**.

38) Job Payment
===============

\=================

38.1 Payment Schedule & Terms
-----------------------------

### Stories

*   As a **client**, I want payment schedules (milestones/hourly) and terms (net-30, upfront %) so compensation is clear.
    

### Flow

SetPaymentScheduleCommand → ValidateTerms(financial-be) → Persist → Emit to proposals-be

### Projections

job\_payment\_read

### Events

job.payment.schedule.set.v1, job.payment.terms.updated.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 200 ms**.

39) Contract Transition
=======================

\=========================

39.1 Seamless Job→Contract
--------------------------

### Stories

*   As a **client**, I want job details to auto-populate contracts upon hire.
    

### Flow

TransitionToContractCommand → MapJobData → Trigger contracts-be

### Projections

job\_transition\_read

### Events

job.contract.transitioned.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 250 ms**.

40) Fraud Detection
===================

\=====================

40.1 Auto-Flags (Risk)
----------------------

### Stories

*   As a **moderator**, I want auto-flags for fraud signals so reviews prioritize risk.
    

### Flow

FlagForFraudCommand(system) → AnalyzeSignals(admin-be) → Persist flag

### Projections

job\_fraud\_read

### Events

job.fraud.flagged.v1

### RBAC/SLO

SYSTEM/MODERATOR; **P95 < 150 ms**.

41) Job Security
================

\==================

41.1 MFA for High-Value Jobs
----------------------------

### Stories

*   As a **compliance officer**, I want MFA required above budget thresholds to mitigate account takeovers.
    

### Flow

VerifyHighValuePostCommand → CheckMFA(pkg/auth) → Allow/Block

### Projections

job\_security\_read

### Events

job.security.mfa.verified.v1

### RBAC/SLO

OWNER; **P95 < 200 ms**.

42) Job Feedback
================

\==================

42.1 Post-Closure Feedback
--------------------------

### Stories

*   As a **client/freelancer**, I want to submit feedback on the job process.
    

### Flow

SubmitJobFeedbackCommand → ValidatePostClosure → Persist → Emit to reviews-be

### Projections

job\_feedback\_read

### Events

job.feedback.submitted.v1

### RBAC/SLO

OWNER/FREELANCER; **P95 < 150 ms**.

43) Job ESG
===========

\=============

43.1 ESG Flags & Carbon Estimate
--------------------------------

### Stories

*   As a **client**, I want ESG attributes (sustainable practices, carbon-neutral).
    
*   As a **client**, I want an estimated carbon footprint.
    

### Flow

SetESGFlagsCommand → ValidateAttributes → Persist → Emit ranking boost (search-be)EstimateCarbonFootprintQuery → Calculate(model/API) → Return

### Projections

job\_esg\_read

### Events

job.esg.flags.set.v1, job.esg.estimate.calculated.v1 (optional)

### RBAC/SLO

OWNER/EDITOR; **P95 < 150–300 ms**.


43.3 Dynamic ESG Scoring (auto-score & updates)
-----------------------------------------------

### Stories

*   As a client, I want an ESG score (e.g., remote/hybrid, async practices) to show eco-friendly posture so that ESG-minded talent is attracted.
    
*   As a system, I want to recompute the ESG score whenever sourcing mode, geo rules, or on-site flags change so that the rating stays fresh.
    
*   As a compliance officer, I want the scoring formula versioned so that audits can reproduce results.
    

### Flow

UpdateESGScoreCommand(job\_id) → GatherSignals(job\_esg.flags, geo\_requirements, sourcing\_modes) → CalculateScore(versioned\_formula) → PersistScore()

OnJobSignalChange (event-driven): visibility/sourcing/geo updated → Enqueue UpdateESGScoreCommand

### Projections

*   job\_esg\_read (flags, score, formula\_version, updated\_at)
    

### Events

*   job.esg.score.updated.v1
    
*   (existing) job.esg.flags.set.v1, job.esg.estimate.calculated.v1
    

### RBAC/SLO

*   SYSTEM; P95 < 200 ms; idempotent recompute by (job\_id, change\_seq).

44) Job Sharing
===============

\=================

44.1 Social Links & Referral Incentives
---------------------------------------

### Stories

*   As a **client**, I want tracked share links (X/LinkedIn).
    
*   As a **client**, I want referral bonuses for shares that lead to hires.
    

### Flow

GenerateShareLinkCommand → CreateTrackedURL → PersistSetReferralIncentiveCommand → ValidateBonus(financial-be) → Persist

### Projections

job\_sharing\_read, job\_referrals\_read

### Events

job.sharing.link.generated.v1, job.referral.incentive.set.v1

### RBAC/SLO

OWNER/EDITOR (incentive: OWNER); **P95 < 180–200 ms**.

45) Job Tax
===========

\=============

45.1 Tax Requirements & Reports
-------------------------------

### Stories

*   As a **compliance officer**, I want tax forms (e.g., W-9/1099) required for certain jobs.
    
*   As a **client**, I want automated tax reports.
    

### Flow

SetTaxRequirementsCommand → ValidateForms(financial-be) → PersistGenerateTaxReportQuery → Aggregate(contracts-be, financial-be)

### Projections

job\_tax\_read, job\_tax\_reports\_read

### Events

job.tax.requirements.set.v1, job.tax.report.generated.v1

### RBAC/SLO

OWNER/ADMIN; **P95 < 150–250 ms**.

46) Job Health
==============

\================

46.1 Health Checkpoints (Post-Hire)
-----------------------------------

### Stories

*   As a **client**, I want automated progress checkpoints (milestone reminders).
    

### Flow

ScheduleHealthCheckpointCommand → Enqueue(comms-be) → Persist schedule

### Projections

job\_health\_read

### Events

job.health.checkpoint.scheduled.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 200 ms**.

47) Job Upsell
==============

\================

47.1 Upsell to Long-Term
------------------------

### Stories

*   As a **client**, I want recommendations to convert successful short jobs to long-term contracts.
    

### Flow

SuggestUpsellQuery → AnalyzePerformance(reviews/analytics) → Recommend

### Projections

— (read)

### Events

— (read)

### RBAC/SLO

OWNER; **P95 < 250 ms**.

48) Job Previews
================

\==================

48.1 VR/AR Attachments
----------------------

### Stories

*   As a **creative client**, I want to attach VR/AR previews.
    

### Flow

AttachVRPreviewCommand → ValidateFormat(storage-be) → Persist link

### Projections

job\_previews\_read

### Events

job.preview.vr.attached.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 200 ms**.

49) Bulk Operations
===================

\=====================

49.1 Bulk Updates & CSV Import
------------------------------

### Stories

*   As an **agency client**, I want to bulk-update jobs (close/extend/tag) (≤100).
    
*   As an **enterprise client**, I want to bulk-import jobs from CSV (≤500).
    

### Flow

BulkUpdateJobsCommand(job\_ids\[\]) → ValidateBatchSize → Apply per job → per-job events + job.bulk.updated.summaryBulkImportJobsCommand(file) → storage-be fetch → ParseCSV → ValidateEach → CreateJobs → per-job events + job.bulk.imported.summary

### Projections

job\_bulk\_read, job\_imports\_read

### Events

job.bulk.updated.v1, job.bulk.imported.v1 (+ per-job)

### RBAC/SLO

OWNER/EDITOR; ENTERPRISE\_API; **P95 < 500–600 ms**.

50) Webhooks
============

\==============

50.1 Subscriptions & Delivery Logs
----------------------------------

### Stories

*   As a **developer**, I want to subscribe webhooks to job events.
    
*   As an **admin**, I want delivery logs and retries.
    

### Flow

SubscribeWebhookCommand(url, events\[\]) → ValidateURL → PersistGetWebhookLogsQuery → Fetch delivery status/attempts/latency

### Projections

job\_webhooks\_read, webhook\_logs\_read

### Events

job.webhook.subscribed.v1 (+ system emits payloads via outbox)

### RBAC/SLO

OWNER/ADMIN; **P95 < 150–200 ms**.

51) Job Archives (User-managed)
===============================

\=================================

51.1 Archive & Reactivate
-------------------------

### Stories

*   As a **client**, I want to archive completed jobs and reactivate for quick repost.
    

### Flow

ArchiveJobCommand → CheckClosure → MoveToArchive → Update search-beReactivateJobCommand → CloneFromArchive → ApplyUpdates → Publish

### Projections

job\_archives\_read

### Events

job.archived.v1, job.reactivated.v1

### RBAC/SLO

OWNER; **P95 < 200–250 ms**.

52) Job Tools
=============

\===============

52.1 Time Tracking Preferences
------------------------------

### Stories

*   As a **client**, I want preferred time-tracking tools specified pre-hire.
    

### Flow

SetTimeTrackingPrefsCommand → ValidateTools → Persist → Link to contracts-be

### Projections

job\_tools\_read

### Events

job.tools.time\_tracking.set.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 150 ms**.

53) Job Custom Fields
=====================

\=======================

53.1 Custom Fields (Extensibility)
----------------------------------

### Stories

*   As a **client**, I want custom key-value fields for niche requirements.
    

### Flow

AddCustomFieldCommand → ValidateSchema → Persist

### Projections

job\_custom\_read

### Events

job.custom.field.added.v1

### RBAC/SLO

OWNER/EDITOR; **P95 < 180 ms**.

54) Inbox: User Signals
=======================

\=========================

54.1 User/Account Consumers
---------------------------

### Stories

*   As a **system**, I want user/account signals (created/updated/verified) to sync ownership, seats/roles, and cached permissions.
    

### Flow

Consume: user.account.created.v1, user.account.updated.v1, user.fully\_verified.attained.v1 → update owner summaries & hiring team.

### Projections

job\_owner\_summaries\_read, hiring\_team\_read

### Events

— (consumer)

### RBAC/SLO

System; projector **P95 < 150 ms**.

55) Inbox: Proposals
====================

\======================

55.1 Proposal Events → Pipeline/Analytics
-----------------------------------------

### Stories

*   As a **system**, I want proposal submitted/accepted/rejected to update counters and pipeline readiness; lifecycle guards.
    

### Flow

Consume: proposal.submitted.v1, proposal.accepted.v1, proposal.rejected.v1 → update job\_analytics\_read & job\_pipeline\_read.

### Projections

job\_analytics\_read, job\_pipeline\_read

### Events

— (consumer)

### RBAC/SLO

System; **P95 < 150 ms**.

56) Inbox: Subscriptions
========================

\==========================

56.1 Entitlements & Usage
-------------------------

### Stories

*   As **subscriptions**, I want entitlements to gate posting/boosts/invites; usage limits enforced.
    

### Flow

Consume: subscriptions.entitlement.updated.v1, usage.limit.reached.v1 → refresh caches; block gated commands.

### Projections

entitlements\_read, usage\_limits\_read

### Events

— (consumer)

### RBAC/SLO

System; cache refresh ≤ 500 ms; denial check < 10 ms.

57) Inbox: Financial
====================

\======================

57.1 Holds & Fees
-----------------

### Stories

*   As **finance**, I want overdue invoices to block promotions/posting; fee schedule updates to refresh validations.
    

### Flow

Consume: financial.invoice.overdue.v1, financial.fee.schedule.updated.v1 → apply holds; refresh fee validations.

### Projections

financial\_holds\_read, fee\_schedules\_read

### Events

— (consumer)

### RBAC/SLO

System; **P95 < 150 ms**.

58) Inbox: Admin/Moderation/Config
==================================

\====================================

58.1 Feature Flags, Config, Moderation, Suspensions
---------------------------------------------------

### Stories

*   As an **admin**, I want feature flags/config applied quickly; as **T&S**, I want global moderation and suspensions enforced.
    

### Flow

Consume: admin.feature\_flag.updated.v1, admin.config.updated.v1, admin.user\_suspended.v1, admin.moderation.action.applied.v1 → refresh and enforce.

### Projections

feature\_flags\_read, service\_config\_read, suspension\_flags\_read

### Events

— (consumer)

### RBAC/SLO

System; apply **P95 < 120 ms**; enforcement < 10 ms.

59) Inbox: Search Synonyms
==========================

\============================

59.1 Synonym Updates
--------------------

### Stories

*   As a **system**, I want taxonomy synonym updates to refresh mappings.
    

### Flow

Consume: search.taxonomy.synonym.updated.v1 → refresh caches; re-hydrate facets.

### Projections

taxonomy\_synonyms\_read, categories\_read, skills\_read

### Events

— (consumer)

### RBAC/SLO

System; **P95 < 120 ms**.

60) Inbox: Storage Media
========================

\==========================

60.1 Media Processed Callbacks
------------------------------

### Stories

*   As a **system**, I want storage callbacks to update scan status and quarantine unsafe files.
    

### Flow

Consume: storage.media.processed.v1 → update attachment\_scan\_read; enforce quarantine policy.

### Projections

attachment\_scan\_read, attachment\_policy\_read

### Events

— (consumer)

### RBAC/SLO

System; **P95 < 120 ms**.