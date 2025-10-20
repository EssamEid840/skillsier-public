# 📦 **jobs-be - Job Management Service - Complete User Stories**

---

## **1 - CORE JOB DOMAIN**

### 1.1 job/

#### User Stories
- As a **client**, I want to **create a job posting** (title, description, type, budget, category, skills) so that freelancers can find and apply to my project.
- As a **client**, I want to **save jobs as drafts** so that I can complete them later before publishing.
- As a **client**, I want to **update job details** so that information stays current.
- As a **client**, I want to **publish jobs** so that they become visible to freelancers.
- As a **client**, I want to **close jobs** when positions are filled so that I stop receiving applications.
- As a **client**, I want to **archive old jobs** so that my job list stays organized.
- As a **system**, I want to **validate job types** (Fixed, Hourly, Retainer) so that budget structures are correct.
- As a **system**, I want to **track job lifecycle states** (Draft, Open, InProgress, Closed, Archived) so that status is clear.
- As a **system**, I want to **prevent duplicate job postings** within 14 days so that spam is reduced.
- As a **system**, I want to **perform profanity and PII checks** on job descriptions so that content is safe.
- As a **client**, I want to **schedule job publication** so that jobs go live at optimal times.
- As a **client**, I want to **set auto-close dates** so that jobs automatically close after a period.
- As a **client**, I want to **extend job expiration** so that I can keep positions open longer.

#### Flow
1. **CreateJobCommand**(client_id, title, description, job_type, budget, category_id, skills, is_draft) → ValidateFields() | CheckDuplicates() | LintContent() | Persist() → **Outbox:** job.created.v1
2. **UpdateJobCommand**(job_id, updates) → AuthorizeOwner() | ValidateChanges() | Apply() → **Outbox:** job.updated.v1
3. **PublishJobCommand**(job_id, scheduled_at?) → ValidateComplete() | Publish() → **Outbox:** job.published.v1
4. **CloseJobCommand**(job_id, reason) → ValidateState() | Close() → **Outbox:** job.closed.v1
5. **ArchiveJobCommand**(job_id) → ValidateState() | Archive() → **Outbox:** job.archived.v1
6. **ScheduleJobPublicationCommand**(job_id, scheduled_at) → Validate() | Schedule() → **Outbox:** job.scheduled.v1
7. **SetAutoCloseCommand**(job_id, auto_close_at) → Validate() | SetAutoClose() → **Outbox:** job.auto_close_set.v1
8. **ExtendJobCommand**(job_id, new_expiry) → ValidateExtension() | Extend() → **Outbox:** job.extended.v1
9. **GetJobQuery**(job_id) → AuthorizeAccess() | Fetch() → JobDTO
10. **ListJobsQuery**(filters, pagination) → ApplyFilters() | Fetch() → JobListDTO
11. **SearchJobsQuery**(query, filters) → Search() → JobSearchResultsDTO

#### Projections
- job_read
- job_lifecycle_read
- job_search_index

#### Events Published
- job.created.v1
- job.updated.v1
- job.published.v1
- job.closed.v1
- job.archived.v1
- job.scheduled.v1
- job.auto_close_set.v1
- job.extended.v1
- job.auto_closed.v1 (system-triggered)

#### RBAC/SLO
- **RBAC:** OWNER (create/update/publish/close/archive), EDITOR (update), PUBLIC (view published), SYSTEM (auto-close)
- **SLO:** P95 < 200ms (create/update), P95 < 150ms (read), P95 < 100ms (auto-close check)

---

### 1.2 category/

#### User Stories
- As a **client**, I want to **select from hierarchical job categories** so that jobs are properly classified.
- As a **system**, I want to **maintain a category taxonomy** (parent-child relationships) so that browsing is intuitive.
- As an **admin**, I want to **create and update categories** so that the taxonomy evolves with market needs.
- As an **admin**, I want to **reparent categories** so that the hierarchy can be reorganized.
- As a **freelancer**, I want to **browse jobs by category** so that I find relevant opportunities.

#### Flow
1. **CreateCategoryCommand**(name, slug, parent_id, level) → ValidateSlug() | CheckHierarchy() | Persist() → **Outbox:** category.created.v1
2. **UpdateCategoryCommand**(category_id, updates) → AuthorizeAdmin() | Validate() | Update() → **Outbox:** category.updated.v1
3. **ReparentCategoryCommand**(category_id, new_parent_id) → ValidateHierarchy() | Reparent() → **Outbox:** category.reparented.v1
4. **GetCategoryTreeQuery**() → BuildTree() → CategoryTreeDTO
5. **GetCategoryQuery**(category_id) → Fetch() → CategoryDTO
6. **ListCategoriesQuery**(parent_id?) → Filter() | Fetch() → CategoryListDTO

#### Projections
- category_read
- category_tree_read

#### Events Published
- category.created.v1
- category.updated.v1
- category.reparented.v1

#### RBAC/SLO
- **RBAC:** ADMIN (create/update/reparent), PUBLIC (view)
- **SLO:** P95 < 120ms

---

### 1.3 skill/

#### User Stories
- As a **system**, I want to **maintain a global skills taxonomy** so that skills are standardized across the platform.
- As an **admin**, I want to **create and update skills** so that new technologies are represented.
- As an **admin**, I want to **categorize skills** (Programming, Design, Marketing, etc.) so that they're organized.
- As an **admin**, I want to **track skill popularity** so that trending skills are visible.
- As an **admin**, I want to **deprecate obsolete skills** so that the taxonomy stays current.
- As a **client**, I want to **search for skills** when creating jobs so that I find the right requirements.

#### Flow
1. **CreateSkillCommand**(name, category_id, popularity) → ValidateUnique() | Persist() → **Outbox:** skill.created.v1
2. **UpdateSkillCommand**(skill_id, updates) → AuthorizeAdmin() | Update() → **Outbox:** skill.updated.v1
3. **DeprecateSkillCommand**(skill_id, reason) → Validate() | Deprecate() → **Outbox:** skill.deprecated.v1
4. **GetSkillQuery**(skill_id) → Fetch() → SkillDTO
5. **SearchSkillsQuery**(query) → Search() → SkillListDTO
6. **GetPopularSkillsQuery**(category_id?, limit) → OrderByPopularity() → SkillListDTO
7. **GetSkillsByCategoryQuery**(category_id) → Filter() → SkillListDTO

#### Projections
- skill_read
- skill_popularity_read

#### Events Published
- skill.created.v1
- skill.updated.v1
- skill.deprecated.v1

#### RBAC/SLO
- **RBAC:** ADMIN (create/update/deprecate), PUBLIC (view/search)
- **SLO:** P95 < 100ms

---

### 1.4 job_skill/

#### User Stories
- As a **client**, I want to **specify required skills** for my job so that qualified freelancers apply.
- As a **client**, I want to **mark skills as required vs preferred** so that applicants know priorities.
- As a **client**, I want to **add or remove skills from a job** so that requirements stay accurate.
- As a **system**, I want to **link jobs to the global skill taxonomy** so that matching is effective.

#### Flow
1. **AddJobSkillCommand**(job_id, skill_id, is_required) → ValidateSkillExists() | Add() → **Outbox:** job_skill.added.v1
2. **UpdateJobSkillCommand**(job_skill_id, is_required) → AuthorizeOwner() | Update() → **Outbox:** job_skill.updated.v1
3. **RemoveJobSkillCommand**(job_skill_id) → AuthorizeOwner() | Remove() → **Outbox:** job_skill.removed.v1
4. **ListJobSkillsQuery**(job_id) → Fetch() → JobSkillListDTO

#### Projections
- job_skill_read

#### Events Published
- job_skill.added.v1
- job_skill.updated.v1
- job_skill.removed.v1

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (add/update/remove), PUBLIC (view)
- **SLO:** P95 < 130ms

---

## **2 - SCREENING DOMAIN (CONSOLIDATED)**

### 2.1 screening/

#### User Stories
- As a **client**, I want to **add screening questions** so that I filter applicants.
- As a **client**, I want to **support multiple question types** (text, multiple choice, file upload) so that I gather different information.
- As a **client**, I want to **mark questions as required** so that applicants must answer them.
- As a **client**, I want to **require skills tests** so that I verify capabilities.
- As a **client**, I want to **require NDA acceptance** so that confidentiality is ensured.
- As a **client**, I want to **set export control flags** so that compliance requirements are clear.
- As a **client**, I want to **configure compliance policies** so that regulatory requirements are met.
- As an **applicant**, I want to **see screening requirements before applying** so that I know what's expected.

#### Flow
1. **AddScreeningQuestionCommand**(job_id, question_text, question_type, is_required) → Validate() | Add() → **Outbox:** screening_question.added.v1
2. **UpdateScreeningQuestionCommand**(question_id, updates) → AuthorizeOwner() | Update() → **Outbox:** screening_question.updated.v1
3. **RemoveScreeningQuestionCommand**(question_id) → AuthorizeOwner() | Remove() → **Outbox:** screening_question.removed.v1
4. **ToggleNDARequiredCommand**(job_id, required) → Update() → **Outbox:** nda.required.v1
5. **SetExportControlCommand**(job_id, export_control_flag) → Validate() | Update() → **Outbox:** export_control.set.v1
6. **SetCompliancePolicyCommand**(job_id, policy_type, policy_data) → Validate() | Store() → **Outbox:** compliance_policy.set.v1
7. **GetScreeningConfigQuery**(job_id) → Fetch() → ScreeningConfigDTO
8. **ListScreeningQuestionsQuery**(job_id) → Fetch() → ScreeningQuestionListDTO

#### Projections
- screening_read
- compliance_config_read

#### Events Published
- screening_question.added.v1
- screening_question.updated.v1
- screening_question.removed.v1
- nda.required.v1
- export_control.set.v1
- compliance_policy.set.v1
- screening.configured.v1

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (configure), PUBLIC (view requirements)
- **SLO:** P95 < 180ms

---

## **3 - ATTACHMENTS DOMAIN (CONSOLIDATED)**

### 3.1 attachments/

#### User Stories
- As a **client**, I want to **attach documents to jobs** (PDFs, images, videos) so that applicants have complete information.
- As a **client**, I want to **add VR previews** so that remote work environments can be visualized.
- As a **client**, I want to **add AR specifications** so that physical space requirements are clear.
- As a **client**, I want to **remove attachments** so that outdated information is deleted.
- As a **system**, I want to **store file references only** (actual files in storage-be) so that concerns are separated.
- As a **system**, I want to **validate file types and sizes** so that uploads are safe.

#### Flow
1. **AddJobAttachmentCommand**(job_id, file_url, file_type, file_name, attachment_type) → ValidateFile() | StoreReference() → **Outbox:** job_attachment.added.v1
2. **RemoveJobAttachmentCommand**(attachment_id) → AuthorizeOwner() | DeleteReference() → **Outbox:** job_attachment.removed.v1
3. **AddVRPreviewCommand**(job_id, vr_url, vr_metadata) → Validate() | Store() → **Outbox:** job_vr_preview.added.v1
4. **AddARSpecCommand**(job_id, ar_spec_url, ar_metadata) → Validate() | Store() → **Outbox:** job_ar_spec.added.v1
5. **ListJobAttachmentsQuery**(job_id) → Fetch() → JobAttachmentListDTO

#### Projections
- job_attachments_read

#### Events Published
- job_attachment.added.v1
- job_attachment.removed.v1
- job_vr_preview.added.v1
- job_ar_spec.added.v1

#### Events Consumed
- storage.file.uploaded (to validate file exists)
- storage.file.deleted (to clean up references)

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (add/remove), PUBLIC (view)
- **SLO:** P95 < 200ms (add), P95 < 120ms (list)

---

## **4 - INVITATION & SOURCING DOMAINS**

### 4.1 invitation/

#### User Stories
- As a **client**, I want to **invite specific freelancers to apply** so that I can target talent.
- As a **client**, I want to **add a personalized message** to invitations so that they're compelling.
- As a **client**, I want to **set invitation expiry dates** so that urgency is created.
- As a **freelancer**, I want to **receive and accept/decline invitations** so that I control my applications.
- As a **system**, I want to **track invitation status** (Pending, Accepted, Declined, Expired) so that outcomes are clear.
- As a **system**, I want to **auto-expire invitations** so that old invitations are cleaned up.

#### Flow
1. **SendJobInvitationCommand**(job_id, freelancer_id, message, expires_at) → ValidateFreelancer() | Send() → **Outbox:** job_invitation.sent.v1
2. **AcceptInvitationCommand**(invitation_id, freelancer_id) → ValidateOwner() | Accept() → **Outbox:** job_invitation.accepted.v1
3. **DeclineInvitationCommand**(invitation_id, freelancer_id, reason?) → ValidateOwner() | Decline() → **Outbox:** job_invitation.declined.v1
4. **ExpireInvitationCommand**(invitation_id) → Validate() | Expire() → **Outbox:** job_invitation.expired.v1 (system-triggered)
5. **ListJobInvitationsQuery**(job_id) → AuthorizeOwner() | Fetch() → InvitationListDTO
6. **GetMyInvitationsQuery**(freelancer_id) → Fetch() → InvitationListDTO

#### Projections
- job_invitation_read

#### Events Published
- job_invitation.sent.v1
- job_invitation.accepted.v1
- job_invitation.declined.v1
- job_invitation.expired.v1

#### Events Consumed
- user.verified.v1 (to validate freelancer eligibility)

#### RBAC/SLO
- **RBAC:** OWNER (send/view invitations), FREELANCER (accept/decline/view own), SYSTEM (expire)
- **SLO:** P95 < 170ms

---

### 4.2 sourcing/

#### User Stories
- As a **client**, I want to **set job visibility mode** (Public, InviteOnly, PrivateLink) so that I control who sees the job.
- As a **client**, I want to **generate private links** for invite-only jobs so that I can share selectively.
- As a **client**, I want to **attach talent pools** so that I source from specific groups.
- As a **client**, I want to **manage shortlists** so that I track preferred candidates.
- As a **system**, I want to **enforce sourcing constraints** so that visibility rules are followed.

#### Flow
1. **SetSourcingModeCommand**(job_id, mode, constraints) → Validate() | Update() → **Outbox:** sourcing_mode.set.v1
2. **GeneratePrivateLinkCommand**(job_id) → AuthorizeOwner() | GenerateLink() → **Outbox:** private_link.generated.v1
3. **AttachTalentPoolCommand**(job_id, talent_pool_id) → ValidatePool() | Attach() → **Outbox:** talent_pool.attached.v1
4. **UpdateShortlistCommand**(job_id, freelancer_ids) → AuthorizeOwner() | Update() → **Outbox:** shortlist.updated.v1
5. **GetSourcingConfigQuery**(job_id) → AuthorizeOwner() | Fetch() → SourcingConfigDTO

#### Projections
- sourcing_read

#### Events Published
- sourcing_mode.set.v1
- private_link.generated.v1
- talent_pool.attached.v1
- shortlist.updated.v1

#### RBAC/SLO
- **RBAC:** OWNER (configure), SYSTEM (enforce)
- **SLO:** P95 < 150ms

---

## **5 - BUDGET & VISIBILITY DOMAINS**

### 5.1 budget_controls/

#### User Stories
- As a **client**, I want to **set budget ranges** (min/max) so that expectations are clear.
- As a **client**, I want to **specify currency** so that international work is supported.
- As a **client**, I want to **set hourly rate caps** so that costs are controlled.
- As a **system**, I want to **apply FX rules** (quote vs settlement currency, rounding) so that conversions are accurate.
- As a **freelancer**, I want to **see budget information** so that I know if jobs are suitable.

#### Flow
1. **SetBudgetControlsCommand**(job_id, min_budget, max_budget, currency, rate_cap_hourly?) → Validate() | Store() → **Outbox:** budget_controls.set.v1
2. **UpdateBudgetControlsCommand**(job_id, updates) → AuthorizeOwner() | Update() → **Outbox:** budget_controls.updated.v1
3. **SetFXRulesCommand**(job_id, quote_currency, settlement_currency, rounding_mode) → Validate() | Store() → **Outbox:** fx_rules.set.v1
4. **GetBudgetControlsQuery**(job_id) → Fetch() → BudgetControlsDTO

#### Projections
- budget_controls_read
- fx_rules_read

#### Events Published
- budget_controls.set.v1
- budget_controls.updated.v1
- fx_rules.set.v1

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (set/update), PUBLIC (view)
- **SLO:** P95 < 150ms

---

### 5.2 visibility_lifecycle/

#### User Stories
- As a **client**, I want to **schedule job publication** so that jobs go live at specific times.
- As a **client**, I want to **set auto-close dates** so that jobs automatically close.
- As a **client**, I want to **extend job duration** so that I can keep positions open longer.
- As a **client**, I want to **save jobs as drafts** so that I can publish later.
- As a **system**, I want to **enforce lifecycle transitions** (Draft → Published → Closed → Archived) so that states are valid.
- As a **system**, I want to **auto-close expired jobs** so that old jobs don't stay open.

#### Flow
1. **ScheduleJobPublicationCommand**(job_id, scheduled_at) → Validate() | Schedule() → **Outbox:** job.scheduled.v1
2. **SetAutoCloseCommand**(job_id, auto_close_at) → Validate() | Store() → **Outbox:** job.auto_close_set.v1
3. **ExtendJobCommand**(job_id, new_expiry, renewal_policy?) → ValidateExtension() | Extend() → **Outbox:** job.extended.v1
4. **AutoCloseJobCommand**(job_id) → ValidateExpired() | Close() → **Outbox:** job.auto_closed.v1 (system-triggered)
5. **GetVisibilityLifecycleQuery**(job_id) → Fetch() → VisibilityLifecycleDTO

#### Projections
- visibility_lifecycle_read

#### Events Published
- job.scheduled.v1
- job.auto_close_set.v1
- job.extended.v1
- job.auto_closed.v1
- job.published.v1 (when schedule triggers)

#### RBAC/SLO
- **RBAC:** OWNER (schedule/extend), SYSTEM (auto-close), PUBLIC (view published status)
- **SLO:** P95 < 180ms (commands), P95 < 50ms (auto-close check)

---

## **6 - TEMPLATE DOMAIN (CONSOLIDATED)**

### 6.1 template/

#### User Stories
- As a **client**, I want to **create job templates** so that I can reuse common job structures.
- As a **client**, I want to **update templates** so that they stay current.
- As a **client**, I want to **archive unused templates** so that my template library stays organized.
- As a **client**, I want to **clone templates to new jobs** so that posting is faster.
- As an **org admin**, I want to **share templates with my organization** so that team members use consistent formats.
- As an **org admin**, I want to **approve shared templates** so that quality is maintained.
- As a **system**, I want to **version templates** (semantic versioning) so that changes are tracked.
- As a **system**, I want to **deprecate old template versions** so that users migrate to new versions.

#### Flow
1. **CreateJobTemplateCommand**(title, type, default_budget, default_scope, skills, attachments) → Validate() | Persist() → **Outbox:** job_template.created.v1
2. **UpdateJobTemplateCommand**(template_id, updates) → AuthorizeOwner() | Update() → **Outbox:** job_template.updated.v1
3. **ArchiveJobTemplateCommand**(template_id) → AuthorizeOwner() | Archive() → **Outbox:** job_template.archived.v1
4. **CloneTemplateToJobCommand**(template_id, client_id) → Validate() | Clone() | CreateJob() → **Outbox:** job_template.cloned_to_job.v1
5. **ShareTemplateToOrgCommand**(template_id, org_id) → AuthorizeOrgAdmin() | Share() → **Outbox:** job_template.org_shared.v1
6. **ApproveTemplateCommand**(template_id, approved_by) → AuthorizeOrgAdmin() | Approve() → **Outbox:** job_template.approved.v1
7. **CreateTemplateVersionCommand**(template_id, version, changelog) → ValidateSemVer() | CreateVersion() → **Outbox:** job_template.version.created.v1
8. **DeprecateTemplateVersionCommand**(template_version_id, reason) → Deprecate() → **Outbox:** job_template.version.deprecated.v1
9. **GetJobTemplateQuery**(template_id) → AuthorizeAccess() | Fetch() → JobTemplateDTO
10. **ListJobTemplatesQuery**(owner_id?, org_id?) → Filter() | Fetch() → JobTemplateListDTO
11. **ListTemplateVersionsQuery**(template_id) → Fetch() → TemplateVersionListDTO
12. **GetLatestTemplateVersionQuery**(template_id) → Fetch() → TemplateVersionDTO

#### Projections
- job_template_read
- template_versions_read
- org_templates_read

#### Events Published
- job_template.created.v1
- job_template.updated.v1
- job_template.archived.v1
- job_template.cloned_to_job.v1
- job_template.org_shared.v1
- job_template.approved.v1
- job_template.version.created.v1
- job_template.version.deprecated.v1

#### RBAC/SLO
- **RBAC:** OWNER (create/update/archive/clone), ORG_ADMIN (share/approve), SYSTEM (version/deprecate)
- **SLO:** P95 < 200ms (template operations), P95 < 180ms (versioning)

---

## **7 - ELIGIBILITY & REQUIREMENTS DOMAINS**

### 7.1 eligibility_rules/

#### User Stories
- As a **client**, I want to **set geographic restrictions** (allow/deny lists) so that I target specific regions.
- As a **client**, I want to **require KYC verification** so that I hire verified freelancers.
- As a **client**, I want to **set minimum trust tier** so that I ensure quality applicants.
- As a **client**, I want to **allow or disallow agencies** so that I control applicant types.
- As a **client**, I want to **set timezone overlap requirements** so that collaboration is feasible.
- As a **client**, I want to **set radius requirements** for location-based work so that proximity is ensured.
- As a **system**, I want to **evaluate applicant eligibility** so that ineligible applicants are blocked.
- As a **freelancer**, I want to **see eligibility requirements** so that I know if I qualify.

#### Flow
1. **SetJobEligibilityRulesCommand**(job_id, geo_rules, kyc_required, min_tier, agency_allowed, tz_overlap, radius) → Validate() | Store() → **Outbox:** job_eligibility.rules.set.v1
2. **UpdateEligibilityRulesCommand**(job_id, updates) → AuthorizeOwner() | Update() → **Outbox:** job_eligibility.rules.updated.v1
3. **EvaluateApplicantCommand**(job_id, freelancer_id) → FetchRules() | FetchProfile() | Evaluate() → Decision + Reasons
4. **GetJobEligibilityRulesQuery**(job_id) → Fetch() → EligibilityRulesDTO

#### Projections
- job_eligibility_read
- eligibility_decisions_read

#### Events Published
- job_eligibility.rules.set.v1
- job_eligibility.rules.updated.v1
- job_eligibility.applicant.blocked.v1 (when ineligible)
- job_eligibility.applicant.allowed.v1 (when eligible)

#### Events Consumed
- user.verified.v1 (to check KYC status)
- user.trust_level.updated.v1 (to check trust tier)

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (set/update rules), SYSTEM (evaluate), PUBLIC (view rules)
- **SLO:** P95 < 150ms (set rules), P95 < 50ms (evaluate - cached)

---

### 7.2 requirements_matrix/

#### User Stories
- As a **client**, I want to **create a weighted requirements matrix** so that candidates are scored objectively.
- As a **client**, I want to **mark requirements as must-have vs nice-to-have** so that priorities are clear.
- As a **client**, I want to **assign weights to requirements** so that important criteria count more.
- As a **reviewer**, I want to **preview candidate scores** with reasons so that I can assess fit.
- As a **system**, I want to **score candidates against the matrix** so that ranking is automated.

#### Flow
1. **SetRequirementsMatrixCommand**(job_id, requirements[key, must_have, weight], version) → ValidateWeighting() | Store() → **Outbox:** job_requirements.matrix.set.v1
2. **UpdateRequirementWeightCommand**(requirement_id, weight) → AuthorizeOwner() | Update() → **Outbox:** job_requirements.weight.updated.v1
3. **ScoreCandidatePreviewQuery**(job_id, freelancer_id) → FetchMatrix() | FetchProfile() | ComputeScore() → CandidateScoreDTO
4. **GetRequirementsMatrixQuery**(job_id) → Fetch() → RequirementsMatrixDTO

#### Projections
- job_requirements_matrix_read

#### Events Published
- job_requirements.matrix.set.v1
- job_requirements.weight.updated.v1

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (set/update matrix), REVIEWER (score preview)
- **SLO:** P95 < 200ms (set matrix), P95 < 100ms (score preview)

---


## **8 - HIRING TEAM DOMAIN (CONTINUED)**

### 8.1 hiring_team/

#### User Stories
- As a **job owner**, I want to **add team members** (OWNER, REVIEWER, INTERVIEWER) so that hiring is collaborative.
- As a **job owner**, I want to **change member roles** so that responsibilities can be adjusted.
- As a **job owner**, I want to **remove team members** so that access is controlled.
- As a **system**, I want to **prevent removal of the last OWNER** so that jobs aren't orphaned.
- As a **team member**, I want to **view my role permissions** so that I know what I can do.
- As a **job owner**, I want to **view team activity logs** so that I can track changes and actions.

#### Flow
1. **AddHiringTeamMemberCommand**(job_id, user_id, role, added_by) → ValidateRole() | Add() → **Outbox:** job_hiring_team.member.added.v1
2. **ChangeHiringTeamRoleCommand**(job_id, user_id, new_role, changed_by) → AuthorizeOwner() | PreventLastOwnerChange() | Update() → **Outbox:** job_hiring_team.role.changed.v1
3. **RemoveHiringTeamMemberCommand**(job_id, user_id, removed_by) → AuthorizeOwner() | PreventLastOwnerRemoval() | Remove() → **Outbox:** job_hiring_team.member.removed.v1
4. **GetHiringTeamQuery**(job_id) → AuthorizeTeamMember() | Fetch() → HiringTeamDTO
5. **GetTeamActivityLogQuery**(job_id) → AuthorizeOwner() | Fetch() → TeamActivityLogDTO

#### Projections
- hiring_team_read
- team_activity_log_read

#### Events Published
- job_hiring_team.member.added.v1
- job_hiring_team.role.changed.v1
- job_hiring_team.member.removed.v1

#### RBAC/SLO
- **RBAC:** OWNER (add/change/remove members), TEAM_MEMBER (view team)
- **SLO:** P95 < 150ms (add/change/remove), P95 < 100ms (view)

---

## **9 - ANALYTICS DOMAIN**

### 9.1 analytics/

#### User Stories
- As a **client**, I want to **view job performance metrics** (views, applications, engagement) so that I can assess posting effectiveness.
- As a **client**, I want to **track application sources** so that I know where candidates are coming from.
- As a **client**, I want to **see time-to-hire metrics** so that I can optimize my hiring process.
- As a **client**, I want to **compare job performance** against similar jobs so that I can benchmark.
- As a **system**, I want to **aggregate analytics data** from events so that reporting is accurate.
- As an **analyst**, I want to **export analytics data** so that I can perform deep analysis.

#### Flow
1. **TrackJobViewCommand**(job_id, viewer_id, source) → Validate() | Increment() → **Outbox:** job.analytics.view.tracked.v1
2. **TrackJobApplicationCommand**(job_id, applicant_id, source) → Validate() | Record() → **Outbox:** job.analytics.application.tracked.v1
3. **GetJobAnalyticsQuery**(job_id) → AuthorizeOwner() | Aggregate() → JobAnalyticsDTO
4. **GetJobPerformanceQuery**(job_id) → AuthorizeOwner() | Calculate() → JobPerformanceDTO
5. **CompareJobsQuery**(job_ids[], benchmark_criteria) → AuthorizeOwner() | Compare() → JobComparisonDTO
6. **ExportAnalyticsCommand**(job_id, format) → AuthorizeOwner() | Generate() → ExportFileDTO

#### Projections
- job_analytics_read
- job_performance_read
- job_engagement_read

#### Events Published
- job.analytics.view.tracked.v1
- job.analytics.application.tracked.v1
- job.analytics.source.recorded.v1
- job.analytics.exported.v1

#### Events Consumed
- proposal.submitted.v1 (from proposals-be)
- job.published.v1
- job.closed.v1

#### RBAC/SLO
- **RBAC:** OWNER/REVIEWER (view analytics), OWNER (export)
- **SLO:** P95 < 200ms (queries), P95 < 500ms (export)

---

## **10 - MODERATION DOMAIN (CONSOLIDATED)**

### 10.1 moderation/

#### User Stories
- As a **user**, I want to **flag inappropriate jobs** so that the platform stays safe.
- As a **moderator**, I want to **review flagged jobs** so that I can take action.
- As a **moderator**, I want to **apply moderation states** (CLEAN, LIMITED, QUARANTINED, BANNED) so that violating content is managed.
- As a **moderator**, I want to **dismiss false flags** so that valid jobs aren't penalized.
- As a **system**, I want to **auto-flag suspicious jobs** using ML so that moderation is proactive.
- As a **client**, I want to **appeal moderation decisions** so that I can contest mistakes.
- As a **moderator**, I want to **view flag history** so that I can see patterns.

#### Flow
1. **FlagJobCommand**(job_id, flagger_id, reason, details) → ValidateReason() | Create() → **Outbox:** job.flag.submitted.v1
2. **ResolveJobFlagCommand**(flag_id, moderator_id, resolution, action) → AuthorizeModerator() | Resolve() → **Outbox:** job.flag.resolved.v1
3. **DismissJobFlagCommand**(flag_id, moderator_id, reason) → AuthorizeModerator() | Dismiss() → **Outbox:** job.flag.dismissed.v1
4. **ApplyModerationStateCommand**(job_id, moderator_id, state, reason) → AuthorizeModerator() | ApplyState() → **Outbox:** job.moderation.state.applied.v1
5. **LiftModerationStateCommand**(job_id, moderator_id, reason) → AuthorizeModerator() | Lift() → **Outbox:** job.moderation.state.lifted.v1
6. **AppealModerationCommand**(job_id, client_id, appeal_reason) → ValidateAppeal() | Submit() → **Outbox:** job.moderation.appeal.submitted.v1
7. **AutoFlagJobCommand**(job_id, ml_model_id, confidence_score, signals) → ValidateThreshold() | Flag() → **Outbox:** job.auto_flagged.v1 (system-triggered)
8. **GetJobFlagsQuery**(job_id) → AuthorizeModerator() | Fetch() → JobFlagsDTO
9. **GetModerationHistoryQuery**(job_id) → AuthorizeModerator() | Fetch() → ModerationHistoryDTO

#### Projections
- job_flags_read
- job_moderation_read
- moderation_appeals_read

#### Events Published
- job.flag.submitted.v1
- job.flag.resolved.v1
- job.flag.dismissed.v1
- job.moderation.state.applied.v1
- job.moderation.state.lifted.v1
- job.moderation.appeal.submitted.v1
- job.auto_flagged.v1

#### Events Consumed
- user.banned.v1 (to auto-flag jobs by banned users)
- fraud.detected.v1 (from admin-be)

#### RBAC/SLO
- **RBAC:** ANY_USER (flag), MODERATOR (resolve/dismiss/apply state/lift state), OWNER (appeal)
- **SLO:** P95 < 150ms (flag), P95 < 200ms (moderation actions)

---

## **11 - A/B EXPERIMENTS DOMAIN**

### 11.1 ab_experiments/

#### User Stories
- As a **product manager**, I want to **create A/B tests** for job features so that I can optimize conversions.
- As a **product manager**, I want to **assign jobs to experiment variants** so that I can test different approaches.
- As a **product manager**, I want to **track experiment metrics** so that I can measure impact.
- As a **system**, I want to **randomize variant assignments** so that tests are unbiased.
- As a **product manager**, I want to **conclude experiments** and roll out winners so that improvements are deployed.

#### Flow
1. **CreateExperimentCommand**(name, description, variants[], metrics[], start_date, end_date) → ValidateConfig() | Create() → **Outbox:** experiment.created.v1
2. **AssignJobToVariantCommand**(job_id, experiment_id) → GetActiveExperiments() | Randomize() | Assign() → **Outbox:** job.experiment.assigned.v1
3. **TrackExperimentMetricCommand**(job_id, experiment_id, variant_id, metric_name, value) → Record() → **Outbox:** experiment.metric.tracked.v1
4. **ConcludeExperimentCommand**(experiment_id, winner_variant_id, notes) → AuthorizePM() | Conclude() → **Outbox:** experiment.concluded.v1
5. **GetExperimentResultsQuery**(experiment_id) → AuthorizePM() | Aggregate() → ExperimentResultsDTO
6. **ListActiveExperimentsQuery**() → Fetch() → ExperimentListDTO

#### Projections
- ab_experiments_read
- experiment_metrics_read

#### Events Published
- experiment.created.v1
- job.experiment.assigned.v1
- experiment.metric.tracked.v1
- experiment.concluded.v1

#### RBAC/SLO
- **RBAC:** PRODUCT_MANAGER (create/conclude), SYSTEM (assign/track), PM (view results)
- **SLO:** P95 < 100ms (assign), P95 < 200ms (track), P95 < 300ms (results)

---

## **12 - SYNDICATION & DRAFTS DOMAINS (CONSOLIDATED)**

### 12.1 syndication/

#### User Stories
- As a **client**, I want to **syndicate jobs to partner boards** so that I reach more candidates.
- As a **client**, I want to **select which boards to syndicate to** so that I control distribution.
- As a **system**, I want to **retry failed syndications** with backoff so that temporary failures are handled.
- As a **client**, I want to **track syndication status** so that I know where jobs are posted.
- As a **client**, I want to **take down syndicated jobs** so that I can remove jobs from external boards.
- As a **moderator**, I want to **force takedown of violating jobs** so that compliance is maintained.

#### Flow
1. **QueueJobSyndicationCommand**(job_id, board_ids[], priority) → ValidateBoards() | Queue() → **Outbox:** job.syndication.queued.v1
2. **PostToExternalBoardCommand**(job_id, board_id) → CallExternalAPI() | Record() → **Outbox:** job.syndication.posted.v1 (system-triggered)
3. **MarkSyndicationFailedCommand**(job_id, board_id, error, retry_count) → UpdateStatus() | ScheduleRetry() → **Outbox:** job.syndication.failed.v1 (system-triggered)
4. **TakedownSyndicatedJobCommand**(job_id, board_ids[]) → AuthorizeOwner() | CallTakedownAPI() → **Outbox:** job.syndication.takedown.v1
5. **GetSyndicationStatusQuery**(job_id) → AuthorizeOwner() | Fetch() → SyndicationStatusDTO
6. **ListSyndicationPartnersQuery**() → Fetch() → BoardListDTO

#### Projections
- job_syndication_read
- syndication_partners_read

#### Events Published
- job.syndication.queued.v1
- job.syndication.posted.v1
- job.syndication.failed.v1
- job.syndication.takedown.v1

#### Events Consumed
- job.published.v1 (to trigger auto-syndication)
- job.closed.v1 (to trigger takedown)

#### RBAC/SLO
- **RBAC:** OWNER (queue/takedown), SYSTEM (post/retry), MODERATOR (force takedown)
- **SLO:** P95 < 200ms (queue), P95 < 150ms (status query), Async (external API calls with DLQ)

---

### 12.2 drafts/

#### User Stories
- As a **client**, I want to **autosave job drafts** so that I don't lose work.
- As a **client**, I want to **manage multiple drafts** so that I can work on several jobs at once.
- As a **client**, I want to **restore drafts to jobs** so that I can resume work.
- As a **client**, I want to **see last editor info** so that I know who made changes (for team scenarios).
- As a **system**, I want to **detect draft conflicts** so that simultaneous edits don't overwrite each other.
- As a **client**, I want to **delete old drafts** so that my draft list stays clean.

#### Flow
1. **SaveJobDraftCommand**(job_id, draft_data, editor_id) → ValidateDraft() | CheckConflicts() | Save() → **Outbox:** job.draft.saved.v1
2. **RestoreDraftToJobCommand**(draft_id, job_id) → AuthorizeOwner() | Restore() | PublishOrUpdate() → **Outbox:** job.draft.restored.v1
3. **DeleteDraftCommand**(draft_id) → AuthorizeOwner() | Delete() → **Outbox:** job.draft.deleted.v1
4. **GetDraftQuery**(draft_id) → AuthorizeOwner() | Fetch() → DraftDTO
5. **ListDraftsQuery**(client_id, filters) → AuthorizeOwner() | Fetch() → DraftListDTO

#### Projections
- job_drafts_read
- draft_conflicts_read

#### Events Published
- job.draft.saved.v1
- job.draft.restored.v1
- job.draft.deleted.v1
- job.draft.conflict.detected.v1

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (save/restore/delete/view)
- **SLO:** P95 < 120ms (autosave), P95 < 180ms (restore)

---

## **13 - DUPLICATE DETECTION DOMAIN**

### 13.1 duplicate_detection/

#### User Stories
- As a **client**, I want to **set work mode** (onsite/remote/hybrid) so that location expectations are clear.
- As a **client**, I want to **specify allowed regions** so that I target specific geographies.
- As a **client**, I want to **set radius requirements** (in km) for onsite work so that proximity is defined.
- As a **client**, I want to **require timezone overlap** so that collaboration is feasible.
- As a **freelancer**, I want to **see geo requirements** so that I know if jobs are suitable for my location.
- As a **system**, I want to **check candidate geo eligibility** (advisory) so that matching is informed.

#### Flow
1. **SetJobGeoRulesCommand**(job_id, work_mode, allowed_regions[], denied_regions[], radius_km, tz_overlap_hours) → ValidateRules() | Store() → **Outbox:** job.geo.rules.set.v1
2. **UpdateJobGeoRulesCommand**(job_id, updates) → AuthorizeOwner() | Update() → **Outbox:** job.geo.rules.updated.v1
3. **CheckCandidateGeoEligibilityQuery**(job_id, freelancer_id) → FetchRules() | FetchFreelancerLocation() | Evaluate() → EligibilityCheckDTO
4. **GetJobGeoRulesQuery**(job_id) → Fetch() → GeoRulesDTO

#### Projections
- job_geo_read
- geo_eligibility_checks_read

#### Events Published
- job.geo.rules.set.v1
- job.geo.rules.updated.v1

#### Events Consumed
- user.location.updated.v1 (from users-be)

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (set/update rules), PUBLIC (view rules), SYSTEM/TEAM (check eligibility)
- **SLO:** P95 < 200ms (set/update), P95 < 80ms (eligibility check - cached)

---

## **15 - DUPLICATE DETECTION DOMAIN**

### 15.1 duplicate_detection/

#### User Stories
- As a **system**, I want to **detect near-duplicate jobs** using simhash so that spam is prevented.
- As a **system**, I want to **cluster similar jobs** so that duplicates can be identified.
- As a **moderator**, I want to **merge duplicate jobs** so that candidates aren't confused.
- As a **moderator**, I want to **prevent duplicate posting** so that platform quality is maintained.
- As a **client**, I want to **see warnings about similar jobs** so that I can avoid duplication.

#### Flow
1. **UpsertJobDuplicateKeyCommand**(job_id) → ComputeSimhash(title + description) | Store() → **Outbox:** job.duplicate.key.computed.v1 (system-triggered)
2. **FindNearDuplicateJobsQuery**(job_id, threshold) → CompareSimhashes() | ClusterJobs() → DuplicateListDTO
3. **MergeDuplicateJobsCommand**(source_job_id, target_job_id, moderator_id) → AuthorizeModerator() | Validate() | Merge() → **Outbox:** job.duplicate.merged.v1
4. **PreventDuplicatePostCommand**(job_id) → CheckWindow(14days) | Block() → **Outbox:** job.duplicate.prevented.v1 (system-triggered)

#### Projections
- job_duplicates_read
- duplicate_clusters_read

#### Events Published
- job.duplicate.key.computed.v1
- job.duplicate.detected.v1
- job.duplicate.merged.v1
- job.duplicate.prevented.v1

#### Events Consumed
- job.core.created.v1
- job.core.updated.v1

#### RBAC/SLO
- **RBAC:** SYSTEM (compute/detect/prevent), MODERATOR (merge), PUBLIC (view warnings)
- **SLO:** P95 < 150ms (detection), P95 < 100ms (simhash computation - cached)

---

## **14 - LEGAL & COMPLIANCE DOMAIN (CONSOLIDATED)**

### 14.1 legal_controls/

#### User Stories
- As a **client**, I want to **set interview policy** (require panel, preferred slots) so that scheduling is standardized.
- As a **client**, I want to **link calendar slots** so that interview scheduling is automated.
- As a **client**, I want to **set up instant video interviews** (pre-recorded Q&A) so that screening is faster.
- As a **freelancer**, I want to **see interview requirements** so that I can prepare.
- As a **freelancer**, I want to **clear instructions and privacy terms** for video interviews so that I can record confidently.
- As a **system**, I want to **store only non-PII references** to media (storage-be IDs) so that content remains safe.

#### Flow
1. **SetJobInterviewPolicyCommand**(job_id, require_panel, preferred_slots_ref, calendar_link) → ValidatePolicy() | Store() → **Outbox:** job.interview.policy.set.v1
2. **LinkCalendarCommand**(job_id, calendar_provider, calendar_id) → ValidateCalendar() | Link() → **Outbox:** job.interview.calendar.linked.v1
3. **SetVideoInterviewCommand**(job_id, questions[], time_limits, retry_policy) → ValidateQuestions() | UploadQuestionPrompts(communications-be) → **Outbox:** job.video.interview.set.v1
4. **UpdateVideoInterviewCommand**(job_id, updates) → AuthorizeOwner() | ValidatePatch() | Update() → **Outbox:** job.video.interview.updated.v1
5. **DisableVideoInterviewCommand**(job_id) → AuthorizeOwner() | Disable() → **Outbox:** job.video.interview.disabled.v1
6. **GetJobInterviewPolicyQuery**(job_id) → Fetch() → InterviewPolicyDTO
7. **GetVideoInterviewSetupQuery**(job_id) → Fetch() → VideoInterviewSetupDTO

#### Projections
- job_interview_policy_read
- job_calendar_read
- job_video_read

#### Events Published
- job.interview.policy.set.v1
- job.interview.calendar.linked.v1
- job.video.interview.set.v1
- job.video.interview.updated.v1
- job.video.interview.disabled.v1

#### Events Consumed
- storage.file.uploaded (to validate media exists)

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (set/update policy/video), PUBLIC (view requirements)
- **SLO:** P95 < 200ms (set policy), P95 < 250ms (video setup), Async media moderation (DLQ on failures), Rate limit: ≤10 questions per pack

---

## **17 - LEGAL CONTROLS DOMAIN (CONSOLIDATED)**

### 17.1 legal_controls/

#### User Stories
- As **legal counsel**, I want to **set NDA requirements** with template/version pinning so that confidentiality is enforced.
- As **legal counsel**, I want to **flag export control requirements** so that compliance is clear.
- As **legal counsel**, I want to **place legal holds on jobs** so that data is preserved for litigation.
- As **legal counsel**, I want to **remove legal holds** when litigation is resolved so that normal operations resume.
- As **compliance officer**, I want to **set data residency requirements** (EU/US) so that regional regulations are met.
- As a **system**, I want to **block purging of jobs under legal hold** so that evidence is preserved.
- As **legal counsel**, I want to **audit legal control changes** so that compliance is traceable.

#### Flow
1. **SetJobLegalControlsCommand**(job_id, nda_template_id, nda_version, export_control_flag) → ValidateTemplate() | Store() → **Outbox:** job.legal.controls.set.v1
2. **PlaceJobLegalHoldCommand**(job_id, legal_counsel_id, reason, case_id) → AuthorizeLegal() | PlaceHold() → **Outbox:** job.legal.hold.placed.v1
3. **RemoveJobLegalHoldCommand**(job_id, legal_counsel_id, resolution) → AuthorizeLegal() | RemoveHold() → **Outbox:** job.legal.hold.removed.v1
4. **SetDataResidencyCommand**(job_id, residency_policy) → ValidatePolicy() | Store() → **Outbox:** job.compliance.residency.set.v1
5. **GetJobLegalQuery**(job_id) → AuthorizeLegal() | Fetch() → LegalControlsDTO

#### Projections
- job_legal_read
- job_compliance_read
- legal_holds_read

#### Events Published
- job.legal.controls.set.v1
- job.legal.hold.placed.v1
- job.legal.hold.removed.v1
- job.compliance.residency.set.v1

#### Events Consumed
- job.retention.purged.v1 (to validate no legal hold exists)

#### RBAC/SLO
- **RBAC:** LEGAL_COUNSEL (set/place hold/remove hold), COMPLIANCE_OFFICER (set residency), ADMIN (view)
- **SLO:** P95 < 200ms

---

## **15 - CAMPAIGN & RETENTION DOMAINS (CONSOLIDATED)**

### 15.1 campaign_tags/

#### User Stories
- As a **client**, I want to **add internal campaign tags** so that I can organize jobs by marketing initiatives.
- As a **client**, I want to **tag VIP jobs** so that they receive special handling.
- As a **system**, I want to **normalize tags** (lowercase, trim) so that duplicates are avoided.
- As a **system**, I want to **limit tags per job** so that abuse is prevented.
- As a **client**, I want to **remove tags** so that I can update campaigns.

#### Flow
1. **AddJobCampaignTagCommand**(job_id, tag, client_id) → NormalizeTag() | ValidateLimit() | Add() → **Outbox:** job.campaign.tag.added.v1
2. **RemoveJobCampaignTagCommand**(job_id, tag, client_id) → AuthorizeOwner() | Remove() → **Outbox:** job.campaign.tag.removed.v1
3. **ListJobCampaignTagsQuery**(job_id) → AuthorizeOwner() | Fetch() → CampaignTagsDTO

#### Projections
- job_campaign_tags_read

#### Events Published
- job.campaign.tag.added.v1
- job.campaign.tag.removed.v1

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (add/remove/view tags)
- **SLO:** P95 < 120ms

---

### 15.2 retention_rules/

#### User Stories
- As **legal/admin**, I want to **set archive schedules** so that old jobs are automatically archived.
- As **legal/admin**, I want to **set purge schedules** so that data retention policies are met.
- As **legal/admin**, I want to **block purging under legal hold** so that evidence is preserved.
- As a **compliance officer**, I want to **anonymize job data** after retention periods so that GDPR is satisfied.
- As a **system**, I want to **execute retention policies automatically** so that compliance is maintained.

#### Flow
1. **SetJobRetentionRulesCommand**(job_id, archive_after_days, purge_after_days, anonymize_after_days) → ValidateRules() | Store() → **Outbox:** job.retention.set.v1
2. **ArchiveJobByRetentionCommand**(job_id) → CheckLegalHold() | Archive() → **Outbox:** job.retention.archived.v1 (system-triggered)
3. **PurgeJobByRetentionCommand**(job_id) → CheckLegalHold() | Purge() → **Outbox:** job.retention.purged.v1 (system-triggered)
4. **AnonymizeJobDataCommand**(job_id) → Anonymize() → **Outbox:** job.retention.anonymized.v1 (system-triggered)
5. **GetJobRetentionQuery**(job_id) → AuthorizeLegal() | Fetch() → RetentionRulesDTO

#### Projections
- job_retention_read

#### Events Published
- job.retention.set.v1
- job.retention.archived.v1
- job.retention.purged.v1
- job.retention.anonymized.v1

#### Events Consumed
- job.legal.hold.placed.v1 (to block purge)
- job.legal.hold.removed.v1 (to resume purge eligibility)

#### RBAC/SLO
- **RBAC:** LEGAL/ADMIN (set rules), SYSTEM (execute policies)
- **SLO:** P95 < 200ms (set rules), P95 < 50ms (retention check - async batch)

---

## **16 - PROMOTION & JOB PREFERENCES DOMAINS (CONSOLIDATED)**

### 16.1 promotion/

#### User Stories
- As a **client**, I want to **activate job promotions** (featured/boosted) after payment so that visibility increases.
- As a **client**, I want to **renew promotions** so that jobs stay featured.
- As a **client**, I want to **suspend promotions** so that I can pause spending.
- As a **system**, I want to **expire promotions** automatically so that billing is accurate.
- As a **client**, I want to **preview promotion fees** so that costs are transparent.
- As a **client**, I want to **apply budget boosts** so that my job gets priority placement.
- As a **client**, I want to **AI-driven visibility boosts** so that high-potential jobs get more exposure.
- As a **system**, I want to **bound AI boosts by plan/entitlement** so that costs stay predictable.
- As an **analyst**, I want to **track boost attribution** (lift vs control) so that ROI is measurable.

#### Flow
1. **ActivateJobPromotionCommand**(job_id, promotion_type, duration, payment_token) → CheckPayment(financial-be) | Activate() → **Outbox:** job.promotion.activated.v1
2. **RenewJobPromotionCommand**(job_id, promotion_id) → CheckPayment(financial-be) | Renew() → **Outbox:** job.promotion.renewed.v1
3. **SuspendJobPromotionCommand**(job_id, promotion_id) → AuthorizeOwner() | Suspend() → **Outbox:** job.promotion.suspended.v1
4. **ExpireJobPromotionCommand**(job_id, promotion_id) → Expire() → **Outbox:** job.promotion.expired.v1 (system-triggered)
5. **ApplyBudgetBoostCommand**(job_id, boost_level) → CheckFees(financial-be) | ApplyBoost() → **Outbox:** job.budget.boost.applied.v1
6. **ApplyAIVisibilityBoostCommand**(job_id) → AnalyzeEngagement(job_analytics_read) | CheckEntitlements(subscriptions-be) | CheckFinancialHolds(financial-be) | SetBoostLevel(search-be signal) → **Outbox:** job.ai.visibility.boosted.v1 (system-triggered)
7. **RecomputeBoostsBatchCommand**() → IterateEligibleJobs() | AdjustBoostLevels() → Multiple boost events (system-triggered)
8. **GetJobPromotionQuery**(job_id) → AuthorizeOwner() | Fetch() → PromotionDTO
9. **PreviewPromotionFeesQuery**(job_id, promotion_type, duration) → Calculate(financial-be) → FeePreviewDTO

#### Projections

*   job\_promotions\_read
    
*   job\_boosts\_read
    
*   job\_visibility\_read
    
*   boost\_attribution\_read
    

#### Events Published

*   job.promotion.activated.v1
    
*   job.promotion.renewed.v1
    
*   job.promotion.suspended.v1
    
*   job.promotion.expired.v1
    
*   job.budget.boost.applied.v1
    
*   job.ai.visibility.boosted.v1
    
*   job.promotion.attribution.tracked.v1
    

#### Events Consumed

*   payment.captured.v1 (from financial-be - to confirm payment before activation)
    
*   payment.failed.v1 (from financial-be - to handle payment failures)
    
*   subscription.usage.checked.v1 (from subscriptions-be - to validate entitlements for AI boosts)
    
*   job.analytics.view.tracked.v1 (to analyze engagement for AI boost decisions)
    

#### RBAC/SLO

*   **RBAC:** OWNER (activate/renew/suspend/preview fees/apply budget boost), SYSTEM (expire/AI boost/recompute batch), ANALYST (view attribution/track ROI)
    
*   **SLO:**
    
    *   P95 < 180ms (activation/renewal)
        
    *   P95 < 150ms (preview fees/suspend)
        
    *   P95 < 300ms (AI boost computation with analytics aggregation)
        
    *   Async (batch recompute - runs every 6 hours)
        
    *   **AI boost constraints:** Capped by plan tier; cooldown 6h between changes per job; max boost multiplier bounded by subscription
        
    *   **Rate limits:** 10 promotion activations per client per hour

---


### 16.2 job_preference/

#### User Stories
- As a **client**, I want to **set advisory matching preferences** (freelancer type, preferred locations/TZs, min success score, fluency level) so that I attract suitable candidates.
- As a **client**, I want to **specify minimum platform earnings** so that experienced freelancers apply.
- As a **client**, I want to **indicate guidance level** (hands-on vs autonomous) so that work style is clear.
- As a **client**, I want to **specify tool provision** (client-provided vs freelancer-owned) so that resource expectations are set.
- As a **system**, I want to **treat preferences as soft signals** (non-blocking) so that matching is flexible.
- As a **client**, I want to **update preferences** so that requirements stay current.
- As a **client**, I want to **remove preferences** so that I can reset to defaults.

#### Flow
1. **SetJobPreferencesCommand**(job_id, freelancer_type, preferred_locations[], preferred_tzs[], min_success_score, fluency_level, min_earnings, guidance_level, tool_provision) → ValidatePreferences() | Store() → **Outbox:** job.preferences.set.v1
2. **UpdateJobPreferencesCommand**(job_id, updates) → AuthorizeOwner() | ValidatePatch() | Update() → **Outbox:** job.preferences.updated.v1
3. **RemoveJobPreferencesCommand**(job_id) → AuthorizeOwner() | Remove() → **Outbox:** job.preferences.removed.v1
4. **GetJobPreferencesQuery**(job_id) → Fetch() → JobPreferencesDTO

#### Projections
- job_preferences_read

#### Events Published
- job.preferences.set.v1
- job.preferences.updated.v1
- job.preferences.removed.v1

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (set/update/remove), PUBLIC (view)
- **SLO:** P95 < 150ms

---

## **17 - MULTI-HIRE & REPOST DOMAIN**

### 17.1 hiring_option/

#### User Stories
- As a **client**, I want to **enable multi-hire** with a maximum hire count so that I can fill multiple positions from one job.
- As a **client**, I want to **set hiring count** (plan-bounded) so that slots are defined.
- As a **client**, I want to **reserve open slots** as candidates are hired so that availability is tracked.
- As a **client**, I want to **release slots** when hires fall through so that positions become available again.
- As a **client**, I want to **repost jobs safely** with cooldowns so that spam is prevented.
- As a **system**, I want to **check for duplicates on repost** so that the same job isn't posted multiple times.
- As a **system**, I want to **enforce repost cooldowns** so that abuse is prevented.

#### Flow
1. **EnableMultiHireCommand**(job_id, max_hires) → ValidateMaxHires(subscriptions-be) | Enable() → **Outbox:** job.multihire.enabled.v1
2. **SetHiringCountCommand**(job_id, hiring_count) → ValidateAgainstPlan(subscriptions-be) | Set() → **Outbox:** job.hiring.count.set.v1
3. **ReserveOpenSlotCommand**(job_id, candidate_id) → CheckAvailableSlots() | Reserve() → **Outbox:** job.slot.reserved.v1
4. **ReleaseOpenSlotCommand**(job_id, slot_id, reason) → AuthorizeOwner() | Release() → **Outbox:** job.slot.released.v1
5. **RepostJobCommand**(job_id, reason) → CheckCooldown() | CheckDuplicate() | Repost() → **Outbox:** job.reposted.v1
6. **GetHiringOptionsQuery**(job_id) → AuthorizeOwner() | Fetch() → HiringOptionsDTO

#### Projections
- job_hiring_options_read
- repost_history_read

#### Events Published
- job.multihire.enabled.v1
- job.hiring.count.set.v1
- job.slot.reserved.v1
- job.slot.released.v1
- job.reposted.v1
- job.repost.cooldown.blocked.v1

#### Events Consumed
- contract.created.v1 (to reserve slots)
- contract.ended.v1 (to release slots)

#### RBAC/SLO
- **RBAC:** OWNER (enable/set count/reserve/release/repost), PUBLIC (view)
- **SLO:** P95 < 120ms

---

## **18 - AI ASSIST DOMAIN**

### 18.1 ai_assist/

#### User Stories
- As a **client**, I want to **receive AI suggestions** for skills, categories, and description optimization so that job postings are improved.
- As a **client**, I want to **accept or reject suggestions** so that I maintain control.
- As a **client**, I want to **optimize job descriptions** using AI so that clarity and appeal are enhanced.
- As a **client**, I want to **see optimization feedback** so that I understand changes.
- As a **system**, I want to **respect guardrails** (RBAC, moderation, budget caps, policy) so that the AI never violates platform rules.
- As a **compliance officer**, I want to **audit all AI actions** so that decisions are traceable.
- As a **client**, I want to **approve or auto-apply safe refinements** so that listings improve without constant manual edits.
- As a **reviewer (hiring team)**, I want to **see digest of suggested changes** with reasons and projected impact so that I can make informed approvals.

#### Flow
1. **RequestAISuggestionsCommand**(job_id) → AnalyzeJob() | GenerateSuggestions(AI engine) | Store() → **Outbox:** job.ai.suggestions.generated.v1
2. **AcceptAISuggestionsCommand**(job_id, suggestion_ids[], client_id) → ValidateSuggestions() | ApplySuggestions() → **Outbox:** job.ai.suggestions.accepted.v1
3. **OptimizeJobDescriptionCommand**(job_id, optimization_type) → ValidateJobState() | OptimizeDescription(AI engine) | GenerateDiff() → **Outbox:** job.description.optimized.v1
4. **GetAISuggestionsQuery**(job_id) → AuthorizeOwner() | Fetch() → AISuggestionsDTO
5. **GetOptimizationFeedbackQuery**(job_id) → AuthorizeOwner() | Fetch() → OptimizationFeedbackDTO
6. **AutoApplySafeRefinementsCommand**(job_id) → AnalyzePerformance() | GenerateRefinements() | CheckSafetyRules() | ApplyOrQueue() → **Outbox:** job.ai.refinement.applied.v1 (system-triggered)

#### Projections
- job_ai_suggestions_read
- job_optimization_feedback_read
- ai_audit_log_read

#### Events Published
- job.ai.suggestions.generated.v1
- job.ai.suggestions.accepted.v1
- job.description.optimized.v1
- job.ai.refinement.applied.v1
- job.ai.action.audited.v1

#### Events Consumed
- job.analytics.view.tracked.v1 (to analyze performance)
- proposal.submitted.v1 (to assess proposal quality)

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (request/accept/optimize), REVIEWER (view digests), COMPLIANCE_OFFICER (audit), SYSTEM (auto-apply)
- **SLO:** P95 < 300ms (suggestions), P95 < 400ms (optimization), Rate limit: 3 optimizations per job/day

---

## **19 - INCLUSIVITY DOMAIN**

### 19.1 inclusivity/

#### User Stories
- As a **client**, I want to **set accessibility flags** (flexible hours, no video required, screen reader friendly) so that diverse freelancers can apply.
- As a **client**, I want to **enable asynchronous work flag** so that timezone constraints are relaxed.
- As a **client**, I want to **set neurodiversity-friendly flag** so that I signal inclusive practices.
- As a **freelancer**, I want to **see inclusivity flags** so that I know if jobs accommodate my needs.
- As a **system**, I want to **emit inclusivity signals to search-be** so that matching considers accessibility.

#### Flow
1. **SetAccessibilityFlagsCommand**(job_id, flexible_hours, no_video_required, screen_reader_friendly, asynchronous_work, neurodiversity_friendly) → ValidateFlags() | Store() | EmitSignal(search-be) → **Outbox:** job.inclusivity.flags.set.v1
2. **UpdateAccessibilityFlagsCommand**(job_id, updates) → AuthorizeOwner() | Update() | EmitSignal(search-be) → **Outbox:** job.inclusivity.flags.updated.v1
3. **GetInclusivityFlagsQuery**(job_id) → Fetch() → InclusivityFlagsDTO

#### Projections
- job_inclusivity_read

#### Events Published
- job.inclusivity.flags.set.v1
- job.inclusivity.flags.updated.v1

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (set/update), PUBLIC (view)
- **SLO:** P95 < 150ms

---

## **20 - CONTRACT TRANSITION DOMAIN**

### 20.1 contract_transition/

#### User Stories
- As a **client**, I want to **job details to auto-populate contracts upon hire** so that contract creation is seamless.
- As a **system**, I want to **map job data to contract seed data** so that transition is accurate.
- As a **system**, I want to **queue transition requests** so that contracts-be can process them.
- As a **client**, I want to **track transition status** so that I know when contracts are ready.
- As a **system**, I want to **handle transition failures** with retries so that reliability is high.

#### Flow
1. **TransitionToContractCommand**(job_id, freelancer_id, proposal_id) → MapJobToContractSeed() | QueueTransition() → **Outbox:** job.contract.transition.queued.v1
2. **MarkTransitionSucceededCommand**(transition_id, contract_id) → UpdateStatus() → **Outbox:** job.contract.transitioned.v1 (system-triggered)
3. **MarkTransitionFailedCommand**(transition_id, error, retry_count) → UpdateStatus() | ScheduleRetry() → **Outbox:** job.contract.transition.failed.v1 (system-triggered)
4. **GetTransitionStatusQuery**(job_id) → AuthorizeOwner() | Fetch() → TransitionStatusDTO

#### Projections
- job_transition_read

#### Events Published
- job.contract.transition.queued.v1
- job.contract.transitioned.v1
- job.contract.transition.failed.v1

#### Events Consumed
- contract.created.v1 (from contracts-be to mark succeeded)
- proposal.accepted.v1 (from proposals-be to trigger transition)

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (initiate transition), SYSTEM (process queue)
- **SLO:** P95 < 250ms (queue transition), Async (contract creation with DLQ)

---

## **21 - FRAUD DETECTION DOMAIN**

### 21.1 fraud_detection/

#### User Stories
- As a **system**, I want to **detect fraud signals automatically** (suspicious patterns, rapid posting, fake budgets) so that risk is identified.
- As a **system**, I want to **score risk signals** so that prioritization is possible.
- As a **system**, I want to **auto-flag high-risk jobs** so that moderators can review.
- As a **moderator**, I want to **view risk signals** so that I can investigate.
- As a **system**, I want to **integrate with admin-be fraud detection** so that patterns are correlated across the platform.

#### Flow
1. **DetectFraudSignalsCommand**(job_id) → AnalyzeJob() | ComputeRiskScore() | ApplyRules() → **Outbox:** job.fraud.signal.detected.v1 (system-triggered)
2. **AutoFlagJobCommand**(job_id, signal_type, score, rule) → ValidateThreshold() | FlagJob() → **Outbox:** job.fraud.flagged.v1 (system-triggered)
3. **GetRiskSignalsQuery**(job_id) → AuthorizeModerator() | Fetch() → RiskSignalsDTO
4. **ListHighRiskJobsQuery**(min_score, filters) → AuthorizeModerator() | Fetch() → HighRiskJobsDTO

#### Projections
- job_fraud_signals_read
- high_risk_jobs_read

#### Events Published
- job.fraud.signal.detected.v1
- job.fraud.flagged.v1

#### Events Consumed
- job.core.created.v1 (to analyze new jobs)
- user.flagged.v1 (from admin-be to correlate)

#### RBAC/SLO
- **RBAC:** SYSTEM (detect/auto-flag), MODERATOR (view signals/list high-risk)
- **SLO:** P95 < 150ms (detection), Async (fraud analysis)

---

## **22 - ESG DOMAIN**

### 22.1 esg/

#### User Stories
- As a **client**, I want to **set ESG flags** (remote-first, local hire, sustainable tools, diversity commitment) so that values are communicated.
- As a **client**, I want to **calculate carbon estimate** for the job so that environmental impact is transparent.
- As a **freelancer**, I want to **see ESG attributes** so that I can choose aligned opportunities.
- As a **system**, I want to **emit ESG signals to search-be** so that filtering is possible.

#### Flow
1. **SetESGFlagsCommand**(job_id, remote_first, local_hire, sustainable_tools, diversity_commitment) → ValidateFlags() | Store() | EmitSignal(search-be) → **Outbox:** job.esg.flags.set.v1
2. **CalculateCarbonEstimateCommand**(job_id) → AnalyzeJobType() | EstimateCarbon() | Store() → **Outbox:** job.esg.estimate.calculated.v1
3. **GetESGAttributesQuery**(job_id) → Fetch() → ESGAttributesDTO

#### Projections
- job_esg_read

#### Events Published
- job.esg.flags.set.v1
- job.esg.estimate.calculated.v1

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (set flags/calculate), PUBLIC (view)
- **SLO:** P95 < 180ms

---

## **23 - SHARING DOMAIN**

### 23.1 sharing/

#### User Stories
- As a **client**, I want to **generate tracked sharing links** with UTM parameters so that referral sources are tracked.
- As a **client**, I want to **set referral incentives** so that sharing is encouraged.
- As a **client**, I want to **track link performance** so that I know which channels are effective.
- As a **system**, I want to **enforce sharing quotas** so that spam is prevented.

#### Flow
1. **GenerateShareLinkCommand**(job_id, utm_params, incentive_type) → ValidateQuota() | GenerateUniqueLink() | Store() → **Outbox:** job.sharing.link.generated.v1
2. **SetReferralIncentiveCommand**(job_id, incentive_type, incentive_value) → ValidateIncentive() | Store() → **Outbox:** job.referral.incentive.set.v1
3. **TrackShareClickCommand**(link_id, visitor_id) → Record() → **Outbox:** job.sharing.link.clicked.v1
4. **GetShareLinksQuery**(job_id) → AuthorizeOwner() | Fetch() → ShareLinksDTO
5. **GetSharePerformanceQuery**(job_id) → AuthorizeOwner() | Aggregate() → SharePerformanceDTO

#### Projections
- job_sharing_read
- share_performance_read

#### Events Published
- job.sharing.link.generated.v1
- job.referral.incentive.set.v1
- job.sharing.link.clicked.v1

#### RBAC/SLO
- **RBAC:** OWNER (generate/set incentive/view), PUBLIC (click tracking)
- **SLO:** P95 < 150ms (generate), P95 < 80ms (track click)

---

## **24 - BULK OPERATIONS DOMAIN**

### 24.1 bulk_ops/

#### User Stories
- As an **agency client**, I want to **bulk-update jobs** (close/extend/tag) (≤100 jobs) so that management is efficient.
- As an **enterprise client**, I want to **bulk-import jobs from CSV** (≤500 jobs) so that onboarding is fast.
- As a **system**, I want to **validate batch sizes** so that performance is maintained.
- As a **system**, I want to **emit per-job events plus summary** so that tracking is comprehensive.
- As a **client**, I want to **track bulk operation status** so that I know when operations complete.

#### Flow
1. **BulkUpdateJobsCommand**(job_ids[], operation, params) → ValidateBatchSize(≤100) | ValidatePermissions() | ApplyPerJob() | GenerateSummary() → **Outbox:** Multiple per-job events + job.bulk.updated.summary.v1
2. **BulkImportJobsCommand**(file_url, client_id) → FetchFile(storage-be) | ParseCSV() | ValidateEach() | CreateJobs() | GenerateSummary() → **Outbox:** Multiple job.core.created.v1 + job.bulk.imported.summary.v1
3. **GetBulkOperationStatusQuery**(batch_id) → AuthorizeOwner() | Fetch() → BulkOperationStatusDTO
4. **ListBulkOperationsQuery**(client_id, filters) → AuthorizeOwner() | Fetch() → BulkOperationListDTO

#### Projections
- job_bulk_ops_read
- job_imports_read

#### Events Published
- job.bulk.updated.summary.v1 (plus per-job events)
- job.bulk.imported.summary.v1 (plus per-job events)

#### Events Consumed
- storage.file.uploaded (to fetch CSV)

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (bulk update - ≤100 jobs), ENTERPRISE_API (bulk import - ≤500 jobs)
- **SLO:** P95 < 600ms (bulk operations), Async (large imports with progress tracking)

---

## **25 - WEBHOOKS DOMAIN**

### 25.1 webhooks/

#### User Stories
- As a **developer**, I want to **subscribe webhooks to job events** so that my system receives real-time updates.
- As a **developer**, I want to **configure delivery policies** (retries, timeouts) so that reliability is high.
- As a **developer**, I want to **secure webhooks with HMAC** so that authenticity is verified.
- As an **admin**, I want to **view delivery logs** so that I can debug issues.
- As an **admin**, I want to **retry failed deliveries** so that no events are lost.
- As a **developer**, I want to **unsubscribe webhooks** so that I can stop receiving events.

#### Flow
1. **SubscribeWebhookCommand**(job_id, endpoint_url, events[], secret, delivery_policy) → ValidateEndpoint() | ValidateHMAC() | Subscribe() → **Outbox:** job.webhook.subscribed.v1
2. **UnsubscribeWebhookCommand**(webhook_id) → AuthorizeDeveloper() | Unsubscribe() → **Outbox:** job.webhook.unsubscribed.v1
3. **DeliverWebhookCommand**(webhook_id, event) → SignPayload(HMAC) | Post() | LogDelivery() → **Outbox:** webhook.delivered.v1 or webhook.delivery.failed.v1 (system-triggered)
4. **RetryWebhookDeliveryCommand**(delivery_id) → AuthorizeAdmin() | Retry() → **Outbox:** webhook.delivery.retried.v1
5. **ListWebhooksQuery**(job_id) → AuthorizeOwner() | Fetch() → WebhookListDTO
6. **GetWebhookDeliveryLogsQuery**(webhook_id) → AuthorizeOwner() | Fetch() → DeliveryLogsDTO

#### Projections
- job_webhooks_read
- webhook_delivery_logs_read

#### Events Published
- job.webhook.subscribed.v1
- job.webhook.unsubscribed.v1
- webhook.delivered.v1
- webhook.delivery.failed.v1
- webhook.delivery.retried.v1

#### RBAC/SLO
- **RBAC:** DEVELOPER/OWNER (subscribe/unsubscribe/view), ADMIN (retry/view logs)
- **SLO:** P95 < 150ms (subscribe), P95 < 100ms (unsubscribe), Async (webhook delivery with DLQ and exponential backoff)

---

## **26 - ARCHIVE DOMAIN**

### 26.1 archive/

#### User Stories
- As a **client**, I want to **manually archive jobs** so that my active job list stays clean.
- As a **client**, I want to **specify archive reasons** so that organization is clear.
- As a **client**, I want to **reactivate archived jobs** so that I can reopen positions.
- As a **client**, I want to **view archive history** so that I can track job lifecycle.
- As a **system**, I want to **prevent archiving of jobs with active contracts** so that operations aren't disrupted.
- As a **system**, I want to **distinguish user-managed archive from retention policy archive** so that intent is clear.

#### Flow
1. **ArchiveJobCommand**(job_id, reason, actor_id) → ValidateNoActiveContracts() | Archive() → **Outbox:** job.archived.v1
2. **ReactivateJobCommand**(job_id, actor_id) → AuthorizeOwner() | ValidateEligibility() | Reactivate() → **Outbox:** job.reactivated.v1
3. **GetArchiveHistoryQuery**(job_id) → AuthorizeOwner() | Fetch() → ArchiveHistoryDTO
4. **ListArchivedJobsQuery**(client_id, filters) → AuthorizeOwner() | Fetch() → ArchivedJobListDTO

#### Projections
- job_archive_read
- archive_history_read

#### Events Published
- job.archived.v1
- job.reactivated.v1

#### Events Consumed
- contract.created.v1 (to prevent archive)
- contract.ended.v1 (to allow archive)

#### RBAC/SLO
- **RBAC:** OWNER (archive/reactivate/view history)
- **SLO:** P95 < 180ms (archive/reactivate), P95 < 150ms (list)

---

## **27 - CUSTOM FIELDS DOMAIN**

### 27.1 custom_fields/

#### User Stories
- As a **client**, I want to **define custom field schemas** for my jobs so that I can capture organization-specific data.
- As a **client**, I want to **add custom fields** (text, number, date, dropdown) so that flexible data collection is possible.
- As a **client**, I want to **set field values** so that job-specific data is stored.
- As a **client**, I want to **validate field values against schema** so that data integrity is maintained.
- As a **client**, I want to **remove custom fields** so that schemas can evolve.
- As a **system**, I want to **limit custom fields per job** so that complexity is managed.

#### Flow
1. **AddCustomFieldCommand**(job_id, field_name, field_type, field_config, is_required) → ValidateFieldLimit() | ValidateSchema() | Add() → **Outbox:** job.custom_field.added.v1
2. **RemoveCustomFieldCommand**(job_id, field_id) → AuthorizeOwner() | Remove() → **Outbox:** job.custom_field.removed.v1
3. **SetCustomFieldValueCommand**(job_id, field_id, value) → ValidateAgainstSchema() | Set() → **Outbox:** job.custom_field.value.set.v1
4. **UpdateCustomFieldSchemaCommand**(job_id, field_id, schema_updates) → AuthorizeOwner() | ValidateMigration() | Update() → **Outbox:** job.custom_field.schema.updated.v1
5. **GetCustomFieldsQuery**(job_id) → AuthorizeOwner() | Fetch() → CustomFieldsDTO
6. **GetCustomFieldValuesQuery**(job_id) → Fetch() → CustomFieldValuesDTO

#### Projections
- job_custom_fields_read
- custom_field_values_read

#### Events Published
- job.custom_field.added.v1
- job.custom_field.removed.v1
- job.custom_field.value.set.v1
- job.custom_field.schema.updated.v1

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (add/remove/set value/update schema), PUBLIC (view values)
- **SLO:** P95 < 180ms

---

## **28 - LOCALIZATION DOMAIN**

### 28.1 localization/

#### User Stories
- As a **client**, I want to **add multi-language job content** (title, description) so that I can reach international talent.
- As a **client**, I want to **set primary locale** so that default language is clear.
- As a **client**, I want to **update localized content** so that translations stay current.
- As a **client**, I want to **remove locales** so that I can stop supporting certain languages.
- As a **freelancer**, I want to **view jobs in my preferred language** so that understanding is easier.
- As a **system**, I want to **validate locale codes** (ISO 639-1) so that standards are followed.

#### Flow
1. **UpsertJobLocaleCommand**(job_id, locale_code, title, summary, body) → ValidateLocaleCode() | ValidatePrimaryLocale() | Upsert() → **Outbox:** job.locale.upserted.v1
2. **SetPrimaryLocaleCommand**(job_id, locale_code) → ValidateLocaleExists() | SetPrimary() → **Outbox:** job.locale.primary.set.v1
3. **RemoveJobLocaleCommand**(job_id, locale_code) → AuthorizeOwner() | PreventRemovePrimary() | Remove() → **Outbox:** job.locale.removed.v1
4. **GetJobLocalesQuery**(job_id) → Fetch() → JobLocalesDTO
5. **GetJobInLocaleQuery**(job_id, locale_code) → FetchLocale() → LocalizedJobDTO

#### Projections
- job_localization_read

#### Events Published
- job.locale.upserted.v1
- job.locale.primary.set.v1
- job.locale.removed.v1

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (upsert/set primary/remove), PUBLIC (view)
- **SLO:** P95 < 200ms

---

## **29 - PAYMENT SCHEDULE DOMAIN**

### 29.1 payment_schedule/

#### User Stories
- As a **client**, I want to **set payment schedules** (milestones/hourly) so that compensation structure is clear.
- As a **client**, I want to **define payment terms** (net-30, upfront percentage) so that expectations are set.
- As a **client**, I want to **validate payment terms with financial-be** so that feasibility is confirmed.
- As a **system**, I want to **emit payment config to proposals-be** so that proposals include correct terms.
- As a **freelancer**, I want to **see payment terms** before applying so that I can assess suitability.

#### Flow
1. **SetPaymentScheduleCommand**(job_id, schedule_type, payment_terms, upfront_percentage, net_days) → ValidateTerms(financial-be) | Store() | EmitToProposals() → **Outbox:** job.payment.schedule.set.v1
2. **UpdatePaymentTermsCommand**(job_id, term_updates) → AuthorizeOwner() | ValidateTerms(financial-be) | Update() → **Outbox:** job.payment.terms.updated.v1
3. **GetPaymentScheduleQuery**(job_id) → Fetch() → PaymentScheduleDTO

#### Projections
- job_payment_read

#### Events Published
- job.payment.schedule.set.v1
- job.payment.terms.updated.v1

#### Events Consumed
- financial.validation.completed.v1 (from financial-be)

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (set/update), PUBLIC (view)
- **SLO:** P95 < 200ms

---

## **30 - SECURITY DOMAIN**

### 30.1 security/

#### User Stories
- As a **compliance officer**, I want to **require MFA for high-value job postings** (above budget threshold) so that account takeovers are mitigated.
- As a **system**, I want to **verify MFA before allowing job publication** so that security is enforced.
- As a **client**, I want to **understand MFA requirements** before posting so that preparation is possible.
- As a **security admin**, I want to **configure MFA thresholds** so that policy can be adjusted.

#### Flow
1. **VerifyHighValuePostCommand**(job_id, client_id, mfa_token) → CheckBudgetThreshold() | VerifyMFA(pkg/auth) | AllowOrBlock() → **Outbox:** job.security.mfa.verified.v1 or job.security.mfa.blocked.v1
2. **ConfigureMFAThresholdCommand**(threshold_amount, currency) → AuthorizeSecurityAdmin() | UpdateConfig() → **Outbox:** job.security.threshold.updated.v1
3. **GetMFARequirementQuery**(job_id) → CheckBudget() | DetermineRequirement() → MFARequirementDTO

#### Projections
- job_security_read

#### Events Published
- job.security.mfa.verified.v1
- job.security.mfa.blocked.v1
- job.security.threshold.updated.v1

#### RBAC/SLO
- **RBAC:** OWNER (verify MFA), SECURITY_ADMIN (configure threshold), PUBLIC (view requirements)
- **SLO:** P95 < 200ms

---

## **31 - FEEDBACK DOMAIN**

### 31.1 feedback/

#### User Stories
- As a **client**, I want to **submit feedback on the job posting process** after closure so that platform improvement is enabled.
- As a **freelancer**, I want to **submit feedback on job quality** so that issues are reported.
- As a **system**, I want to **validate post-closure feedback timing** so that feedback is relevant.
- As a **system**, I want to **emit feedback to reviews-be** for aggregation so that insights are consolidated.
- As a **product manager**, I want to **analyze job feedback trends** so that improvements can be prioritized.

#### Flow
1. **SubmitJobFeedbackCommand**(job_id, submitter_id, submitter_role, feedback_text, rating, categories[]) → ValidatePostClosure() | ValidateFeedback() | Store() | EmitToReviews() → **Outbox:** job.feedback.submitted.v1
2. **GetJobFeedbackQuery**(job_id) → AuthorizeOwnerOrModerator() | Fetch() → JobFeedbackDTO
3. **GetFeedbackTrendsQuery**(filters) → AuthorizeProductManager() | Aggregate() → FeedbackTrendsDTO

#### Projections
- job_feedback_read
- feedback_trends_read

#### Events Published
- job.feedback.submitted.v1

#### Events Consumed
- job.closed.v1 (to allow feedback submission)

#### RBAC/SLO
- **RBAC:** OWNER/FREELANCER (submit), OWNER/MODERATOR (view), PRODUCT_MANAGER (trends)
- **SLO:** P95 < 150ms

---

## **32 - TAX DOMAIN**

### 32.1 tax/

#### User Stories
- As a **compliance officer**, I want to **set tax requirements** (W-9/1099) for jobs so that regulatory compliance is met.
- As a **client**, I want to **validate tax forms with financial-be** so that requirements are accurate.
- As a **client**, I want to **generate automated tax reports** so that filing is simplified.
- As a **system**, I want to **aggregate data from contracts-be and financial-be** for tax reporting so that reports are comprehensive.
- As a **client**, I want to **track tax compliance status** so that I know if requirements are met.

#### Flow
1. **SetTaxRequirementsCommand**(job_id, required_forms[], jurisdiction) → ValidateForms(financial-be) | Store() → **Outbox:** job.tax.requirements.set.v1
2. **UpdateTaxRequirementsCommand**(job_id, updates) → AuthorizeComplianceOfficer() | ValidateForms(financial-be) | Update() → **Outbox:** job.tax.requirements.updated.v1
3. **GenerateTaxReportQuery**(job_id, tax_year) → AggregateData(contracts-be, financial-be) | GenerateReport() → TaxReportDTO
4. **GetTaxComplianceStatusQuery**(job_id) → AuthorizeOwner() | CheckCompliance() → TaxComplianceDTO
5. **GetTaxRequirementsQuery**(job_id) → Fetch() → TaxRequirementsDTO

#### Projections
- job_tax_read
- job_tax_reports_read

#### Events Published
- job.tax.requirements.set.v1
- job.tax.requirements.updated.v1
- job.tax.report.generated.v1

#### Events Consumed
- contract.created.v1 (for tax reporting)
- payment.processed.v1 (from financial-be for tax reporting)

#### RBAC/SLO
- **RBAC:** COMPLIANCE_OFFICER/OWNER (set/update requirements), OWNER/ADMIN (generate reports/view compliance)
- **SLO:** P95 < 200ms (set/update), P95 < 500ms (generate report - aggregation)

---

## **33 - HEALTH CHECKPOINTS DOMAIN**

### 33.1 health/

#### User Stories
- As a **client**, I want to **schedule automated progress checkpoints** (milestone reminders) so that project health is monitored.
- As a **client**, I want to **configure checkpoint frequency** so that cadence matches project needs.
- As a **system**, I want to **enqueue checkpoint notifications to communications-be** so that reminders are sent.
- As a **client**, I want to **track checkpoint history** so that I can see past reminders.
- As a **system**, I want to **trigger checkpoints based on job lifecycle events** so that timing is appropriate.

#### Flow
1. **ScheduleHealthCheckpointCommand**(job_id, checkpoint_type, frequency, trigger_conditions) → ValidateSchedule() | EnqueueNotification(communications-be) | Store() → **Outbox:** job.health.checkpoint.scheduled.v1
2. **TriggerCheckpointCommand**(job_id, checkpoint_id) → SendNotification(communications-be) | RecordTrigger() → **Outbox:** job.health.checkpoint.triggered.v1 (system-triggered)
3. **UpdateCheckpointFrequencyCommand**(checkpoint_id, new_frequency) → AuthorizeOwner() | Update() → **Outbox:** job.health.checkpoint.updated.v1
4. **CancelCheckpointCommand**(checkpoint_id) → AuthorizeOwner() | Cancel() → **Outbox:** job.health.checkpoint.cancelled.v1
5. **GetCheckpointHistoryQuery**(job_id) → AuthorizeOwner() | Fetch() → CheckpointHistoryDTO

#### Projections
- job_health_read
- checkpoint_history_read

#### Events Published
- job.health.checkpoint.scheduled.v1
- job.health.checkpoint.triggered.v1
- job.health.checkpoint.updated.v1
- job.health.checkpoint.cancelled.v1

#### Events Consumed
- contract.milestone.completed.v1 (to trigger checkpoints)
- job.published.v1 (to start checkpoint scheduling)

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (schedule/update/cancel), SYSTEM (trigger)
- **SLO:** P95 < 200ms

---

## **34 - UPSELL DOMAIN**

### 34.1 upsell/

#### User Stories
- As a **client**, I want to **receive recommendations to convert successful short jobs to long-term contracts** so that retention is increased.
- As a **system**, I want to **analyze job performance** (reviews, completion rate) to identify upsell candidates so that suggestions are relevant.
- As a **client**, I want to **view upsell suggestions with reasoning** so that I can make informed decisions.
- As a **client**, I want to **accept or dismiss suggestions** so that I maintain control.
- As a **system**, I want to **track upsell conversion rates** so that recommendation quality is measured.

#### Flow
1. **AnalyzeUpsellOpportunityQuery**(job_id) → AnalyzePerformance(reviews-be, analytics-be) | EvaluateCriteria() | GenerateRecommendation() → UpsellRecommendationDTO
2. **AcceptUpsellSuggestionCommand**(job_id, suggestion_id, client_id) → CreateLongTermJobDraft() | NotifyClient() → **Outbox:** job.upsell.accepted.v1
3. **DismissUpsellSuggestionCommand**(job_id, suggestion_id, reason) → AuthorizeOwner() | Dismiss() → **Outbox:** job.upsell.dismissed.v1
4. **GetUpsellSuggestionsQuery**(client_id, filters) → AuthorizeOwner() | Fetch() → UpsellSuggestionsDTO
5. **TrackUpsellConversionCommand**(suggestion_id, outcome) → RecordConversion() → **Outbox:** job.upsell.conversion.tracked.v1

#### Projections
- job_upsell_read
- upsell_conversion_read

#### Events Published
- job.upsell.suggestion.generated.v1
- job.upsell.accepted.v1
- job.upsell.dismissed.v1
- job.upsell.conversion.tracked.v1

#### Events Consumed
- contract.ended.v1 (to analyze for upsell)
- review.submitted.v1 (to evaluate performance)

#### RBAC/SLO
- **RBAC:** OWNER (view/accept/dismiss), SYSTEM (analyze/track)
- **SLO:** P95 < 250ms (analysis), P95 < 150ms (accept/dismiss)

---

## **35 - VR/AR PREVIEWS DOMAIN**

### 35.1 previews/

#### User Stories
- As a **creative client**, I want to **attach VR previews** to jobs so that immersive work environments can be showcased.
- As a **creative client**, I want to **attach AR specifications** so that spatial requirements are clear.
- As a **system**, I want to **validate media formats with storage-be** so that compatibility is ensured.
- As a **freelancer**, I want to **view VR/AR previews** so that I can assess work context.
- As a **system**, I want to **store only references to media** (storage-be IDs) so that separation of concerns is maintained.

#### Flow
1. **AttachVRPreviewCommand**(job_id, vr_file_url, vr_metadata) → ValidateFormat(storage-be) | ValidateMediaType() | StoreReference() → **Outbox:** job.preview.vr.attached.v1
2. **AttachARSpecCommand**(job_id, ar_file_url, ar_metadata) → ValidateFormat(storage-be) | ValidateMediaType() | StoreReference() → **Outbox:** job.preview.ar.attached.v1
3. **RemovePreviewCommand**(preview_id, preview_type) → AuthorizeOwner() | DeleteReference() → **Outbox:** job.preview.removed.v1
4. **GetJobPreviewsQuery**(job_id) → Fetch() → JobPreviewsDTO

#### Projections
- job_previews_read

#### Events Published
- job.preview.vr.attached.v1
- job.preview.ar.attached.v1
- job.preview.removed.v1

#### Events Consumed
- storage.file.uploaded (to validate media exists)
- storage.file.deleted (to clean up references)

#### RBAC/SLO
- **RBAC:** OWNER/EDITOR (attach/remove), PUBLIC (view)
- **SLO:** P95 < 200ms

---

## **36 - SAVED SEARCHES DOMAIN**

### 36.1 saved_searches/

#### User Stories
- As a **freelancer**, I want to **save job searches** with filters so that I can quickly find relevant opportunities.
- As a **freelancer**, I want to **set up search alerts** so that I'm notified of new matching jobs.
- As a **freelancer**, I want to **update saved searches** so that criteria stay current.
- As a **freelancer**, I want to **delete saved searches** so that my list stays relevant.
- As a **system**, I want to **validate minimum query length** so that searches are meaningful.
- As a **system**, I want to **trigger alerts when matching jobs are posted** so that notifications are timely.

#### Flow
1. **SaveJobSearchCommand**(freelancer_id, search_query, filters, alert_enabled, alert_frequency) → ValidateQueryLength() | Store() → **Outbox:** job.search.saved.v1
2. **UpdateSavedSearchCommand**(search_id, updates) → AuthorizeOwner() | Update() → **Outbox:** job.search.updated.v1
3. **DeleteSavedSearchCommand**(search_id) → AuthorizeOwner() | Delete() → **Outbox:** job.search.deleted.v1
4. **TriggerSearchAlertCommand**(search_id, matching_job_ids[]) → ValidateAlert() | SendNotification(communications-be) → **Outbox:** job.search.alert.sent.v1 (system-triggered)
5. **GetSavedSearchesQuery**(freelancer_id) → AuthorizeOwner() | Fetch() → SavedSearchesDTO
6. **GetSearchAlertHistoryQuery**(search_id) → AuthorizeOwner() | Fetch() → AlertHistoryDTO

#### Projections
- saved_job_searches_read
- search_alert_history_read

#### Events Published
- job.search.saved.v1
- job.search.updated.v1
- job.search.deleted.v1
- job.search.alert.sent.v1

#### Events Consumed
- job.published.v1 (to trigger alerts for matching searches)

#### RBAC/SLO
- **RBAC:** SELF (save/update/delete/view), SYSTEM (trigger alerts)
- **SLO:** P95 < 120ms (save/update/delete), P95 < 100ms (trigger alert), Min query length = 2 characters

---

## **GLOBAL CONVENTIONS & CROSS-CUTTING CONCERNS**

### Event Envelope (All Domains)
```json
{
  "event_id": "uuid",
  "event_ts": "ISO8601",
  "aggregate_id": "job_id",
  "partition_key": "job_id",
  "correlation_id": "uuid",
  "causation_id": "uuid",
  "actor": {
    "id": "user_id",
    "role": "OWNER|EDITOR|REVIEWER|..."
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
- **Retries/DLQ:** For external service calls (search-be, storage-be, financial-be, subscriptions-be, communications-be, contracts-be, proposals-be, admin-be); exponential backoff; poison messages to DLQ
- **Projections:** `_read` materialized views per domain; metric `event_to_projector_lag_ms` tracked
- **Security:** RBAC enforced on all commands/queries; field-level encryption for sensitive data; secrets never logged
- **Performance:** Typical write P95 ≤ 300ms, read P95 ≤ 250ms (unless specified otherwise)
- **Rate Limiting:** Per-endpoint rate limits enforced (e.g., 60 req/min/client for job creation)

### PII Handling (All Domains)
- **NO raw PII in events:** Emit hashes, storage_ids, references only
- **Examples:** No plaintext emails, phone numbers, file contents, file names
- **Compliance:** PII flags in envelope; anonymization tracked; data residency respected

### Integration Patterns
- **Async Event-Driven:** Primary integration via Kafka events
- **Sync Queries:** REST APIs for read operations with caching
- **Command Validation:** External service validation (financial-be, subscriptions-be) before persistence
- **Circuit Breakers:** Implemented for all external service calls
- **Saga Pattern:** Used for multi-service workflows (e.g., job → contract transition)

---

## **SUMMARY**

This document covers **all 36 feature domains** in jobs-be with complete user stories following the pattern:
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
- ✅ Large-scale Upwork-like platform requirements# 📦 **jobs-be - Job Management Service - Complete User Stories (Continuation)**

---
