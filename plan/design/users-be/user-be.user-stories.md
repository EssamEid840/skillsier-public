# Users-be User Stories


Event conventions (applies to all)
==================================

*   **Format:** aggregate.resource.action.past\_tense.v1 (e.g., user.email.verified.v1).
    
*   **Envelope includes:** event\_id, event\_ts, aggregate\_id, partition\_key=user\_id, correlation\_id, causation\_id, actor{id,role}, user\_context{ip,ua}, data\_zone(EU|US), schema\_ref, compliance\_context{pii\_flags}.
    
*   **Batch ops:** Per-entity events + one \*.summary.v1.
    
*   **PII:** Emit hashes/storage\_ids only (no raw PII).
    

Write-path defaults
===================

*   **Idempotency:** Header Idempotency-Key (or envelope).
    
*   **Transactions:** DB tx + outbox with (aggregate\_id,event\_type,idempotency\_key) dedupe.
    
*   **Retries/DLQ:** For external calls (Keycloak, storage-be, KYC, providers).
    
*   **Projections:** \_read views; metric event\_to\_projector\_lag\_ms.
    
*   **Security/Perf:** RBAC on commands/queries; SLO/SLA and rate limits noted where important.
    

## 1) User (Core Account)

1.1 Registration & Bootstrap
----------------------------

### Stories

*   As a prospective user, I want to register (freelancer/client/org) so that I can access the platform while avoiding duplicate identities.
    
*   As a prospective user, I want default profile/settings bootstrapped so that onboarding is seamless.
    
*   As a system, I want case-insensitive uniqueness of KeycloakID/username/email so that global duplicates are prevented.
    
*   As a prospective org user, I want initial seats and roles created so that my team can start immediately.
    
*   **(Added)** As a prospective user, I want to apply a referral code at registration so that I get startup connects/benefits.
    
*   **(Added)** As a system, I want referral validation and attribution so that fraud is prevented.
    
*   **(Added)** As a user, I want password setup at signup and recovery bootstrap so that login is smooth (Keycloak-integrated).
    
*   **(Added)** As a prospective user, I want to select my account type (freelancer/client/org) during registration so that my role is clear.
    
*   **(Added)** As a system, I want to validate account type to prevent role conflicts.

*   **(New)** As a prospective user, I want to create a guest account so that I can browse jobs/profiles before committing to full registration.
    
    

### Flow

*   CreateUserCommand(referral\_code?, password\_setup\_token?, account\_type) →ValidateReferralCodeQuery | ValidateAccountTypeQuery | GetUserByKeycloakIDQuery | GetUserByEmailQuery | GetUserByUsernameQuery →CreateUser() (tx: user + profile + settings \[+ org seats/roles\], attribute referrer, grant bonus connects, initialize password via Keycloak; enforce single role)
    
*   **(New)** CreateGuestUserCommand → GetGuestUserStatusQuery → CreateGuestUser() (limited access, expires in 7 days)

### Projections

*   users\_read, profiles\_read, settings\_read, org\_seats\_read, referrals\_read, connects\_read
*   **(New)** guest\_users\_read

### Events

*   user.account.created.v1, user.verification.required.v1, user.referral.applied.v1, **user.account\_type.assigned.v1**
*   **(New)** user.guest.created.v1

### RBAC/SLO

*   Public; P95 < 700ms; event-to-index ≤ 1s.
    

1.2 Identity Verification (Email / Phone / Identity / Social / Mobile)
----------------------------------------------------------------------

### Stories

*   As a user, I want to verify email/phone/identity so that my account is trusted and fully active.
    
*   As an enterprise admin, I want batch verification so that many accounts can be verified efficiently.
    
*   As a system, I want automatic “fully verified” activation when all checks pass so that permissions unlock.
    
*   As a compliance officer, I want audit logs of verifications so that reviews are traceable.
    
*   **(Added)** As a user, I want social provider verification (Google/LinkedIn OAuth) so that signup is faster.
    
*   **(Added)** As a mobile user, I want SMS-based mobile verification so that I can verify on phone.
    
*   **(Added)** As a system, I want to enforce multi-factor verification for high-risk accounts so that trust is enhanced.
    

### Flow

*   VerifyEmailCommand | VerifyPhoneCommand | VerifyIdentityCommand | VerifyEmailBatchCommand | VerifySocialIdentityCommand | VerifyMobileNumberCommand | EnforceMultiFactorVerificationCommand →GetUnverifiedUsersQuery | GetVerifiedUsersQuery | GetSocialProvidersQuery | GetVerificationRequirementsQuery →VerifyEmail() | VerifyPhone() | VerifyIdentity() | VerifySocialIdentity() | VerifyMobileNumber() | EnforceMultiFactorVerification()
    

### Projections

*   users\_read.verification\_flags, verification\_audit\_read
    

### Events

*   user.email.verified.v1, user.phone.verified.v1, user.identity.verified.v1, user.social.verified.v1, user.mobile.verified.v1, user.fully\_verified.attained.v1, user.email.batch\_verified.summary.v1, **user.verification.multifactor.required.v1**
    

### RBAC/SLO

*   Self; P95 single < 5s; batch 10k < 60s; idempotent tokens.
    

1.3 Account Updates & Partial Update Policy + Password Policy
-------------------------------------------------------------

### Stories

*   As a user, I want partial updates so that I can change only the fields I need.
    
*   As a user, I want re-verification if critical fields change so that trust is maintained.
    
*   As a system, I want rate-limiting on updates so that abuse is prevented.
    
*   **(Added)** As a user, I want to update my password with strength validation so that security is maintained.
    

### Flow

*   UpdateUserCommand → GetUserByIDQuery | ListUsersQuery → UpdateUser() (detect critical change → reverification trigger)
    
*   **(Added)** UpdatePasswordCommand(old,new) → UpdatePassword() (Keycloak)
    

### Projections

*   users\_read, users\_search\_read
    

### Events

*   user.account.updated.v1, user.reverification.triggered.v1, user.password.updated.v1
    

### RBAC/SLO

*   Self; P95 < 200ms; 30 req/min/user.
    

1.4 Search, Listing & Filtering
-------------------------------

### Stories

*   As an admin, I want to search users with full-text and filters so that I can manage at scale.
    
*   As an admin, I want minimum query length and timing metrics so that performance is monitored.
    
*   As an admin, I want country/status/type filters with pagination so that large data is handleable.
    

### Flow

*   SearchUsersQuery | ListUsersQuery | GetUsersByCountryQuery | GetUsersByStatusQuery → SearchUsers() | ListUsers()
    

### Projections

*   users\_search\_read, users\_read — **No events**
    

1.5 Stats, Growth & Retention Analytics
---------------------------------------

### Stories

*   As a platform admin, I want growth trends and user counts so that I can plan capacity.
    
*   As an analyst, I want cohort analysis so that retention is measured.
    
*   As a marketer, I want referral stats so that campaigns are optimized.
    
*   **(Added)** As an analyst, I want retention cohorts (e.g., 30-day active by signup month) so that health is tracked.
    
*   **(Added)** As a marketer, I want churn prediction analytics so that I can target retention campaigns.

*   **(New)** As an analyst, I want engagement metrics (e.g., time on platform, feature usage) so that I can optimize UX.
    

### Flow

*   (Updated) GetUserStatisticsQuery | GetUserGrowthStatsQuery | CountUsersByTypeQuery | CountUsersByStatusQuery | GetUsersByReferrerQuery | GetReferralCountQuery | GetUserRetentionCohortsQuery | GetChurnPredictionQuery | GetEngagementMetricsQuery → GetUserStatistics() | GetUserGrowthStats() | GetReferralCount() | GetUserRetentionCohorts() | GetChurnPrediction() | GetEngagementMetrics()
    

### Projections

*   users\_analytics\_read, referrals\_read, **churn\_predictions\_read** — **No events**
*   **(New)** engagement\_metrics\_read

1.6 Session Hygiene, Login Recording & MFA-on-Anomaly
-----------------------------------------------------

### Stories

*   As a security system, I want to record logins so that risk signals are tracked.
    
*   As a user, I want to view active sessions so that I can monitor access.
    
*   As a security admin, I want to force global sign-out so that compromised accounts are secured.
    
*   As a system, I want auto-expiry on inactivity so that risks are minimized.
    
*   **(Added)** As a system, I want MFA challenge on anomalous logins so that security is layered.
    

### Flow

*   RecordLoginCommand | ForceLogoutAllSessionsCommand | EnforceMFAOnLoginCommand →GetRecentlyActiveUsersQuery | ListActiveSessionsQuery →RecordLogin() | ForceLogoutAllSessions() | EnforceMFAOnLogin() (Keycloak actions)
    

### Projections

*   sessions\_read, users\_read.last\_login
    

### Events

*   user.login.recorded.v1, user.sessions.revoked.v1, user.mfa.challenged.v1
    

1.7 Deactivation & Soft Delete (+ Restore)
------------------------------------------

### Stories

*   As a user, I want to deactivate or soft-delete so that I can pause or leave safely.
    
*   As an admin, I want batch deletions with reports so that cleanup is efficient.
    
*   As a system, I want retention before hard-delete so that compliance is met.
    
*   As a user, I want the ability to restore a soft-deleted account so that I can return later.
    

### Flow

*   DeactivateUserCommand | DeleteUserCommand | DeleteUserBatchCommand | RestoreSoftDeletedUserCommand →DeactivateUser() | DeleteUser() | RestoreSoftDeletedUser()
    

### Projections

*   users\_read.status, gdpr\_deletion\_read
    

### Events

*   user.account.deactivated.v1, user.account.deleted.v1, user.account.batch\_deleted.summary.v1, user.account.restored.v1
    

1.8 Account Linking & Aliases
-----------------------------

### Stories

*   As a user, I want to link/unlink secondary emails/phones/socials so that recovery is easier.
    
*   As a user, I want to promote a verified alias to primary so that my preferred contact is used.
    
*   As a system, I want alias limits so that abuse is prevented.
    
*   As a user, I want deliverability checks and OTP proof so that ownership is verified.
    

### Flow

*   LinkIdentityCommand | UnlinkIdentityCommand | PromotePrimaryIdentityCommand →ListLinkedIdentitiesQuery | GetPrimaryIdentityQuery →LinkIdentity() | UnlinkIdentity() | PromotePrimaryIdentity()
    

### Projections

*   user\_identities\_read
    

### Events

*   user.identity.linked.v1, user.identity.unlinked.v1, user.identity.primary.changed.v1
    

### Policy

*   Max aliases = 5/user.
    

1.9 Delegated Access
--------------------

### Stories

*   As an org owner, I want to grant scoped, expiring delegation so that assistants can help safely.
    
*   As a delegate, I want to accept delegation explicitly so that consent is auditable.
    
*   As a system, I want auto-expiry and revocation logs so that risk is low.
    

### Flow

*   GrantDelegatedAccessCommand | RevokeDelegatedAccessCommand | AcceptDelegationCommand →ListDelegationsQuery | GetDelegationStatusQuery →GrantDelegatedAccess() | RevokeDelegatedAccess() | AcceptDelegation()
    

### Projections

*   delegations\_read
    

### Events

*   user.delegation.granted.v1, user.delegation.accepted.v1, user.delegation.revoked.v1
    

1.10 Security Freeze / Account Lock
-----------------------------------

### Stories

*   As a user, I want to lock/unlock my account so that activity is frozen after suspicious events.
    
*   As a system, I want auto-lock on anomalies so that proactive defense is applied.
    
*   As an admin, I want lock history so that investigations are complete.
    

### Flow

*   LockAccountCommand | UnlockAccountCommand →GetAccountLockStateQuery | GetLockHistoryQuery →LockAccount() | UnlockAccount()
    

### Projections

*   account\_locks\_read
    

### Events

*   user.account.locked.v1, user.account.unlocked.v1
    

1.11 Username Lifecycle
-----------------------

### Stories

*   As a user, I want to request a username change with a 90-day cooldown so that spam is prevented.
    
*   As a system, I want reserved and profanity filters so that names are acceptable.
    
*   As an admin, I want to approve premium names so that brand safety is preserved.
    

### Flow

*   RequestUsernameChangeCommand →GetUsernameHistoryQuery | GetUsernameAvailabilityQuery →RequestUsernameChange()
    

### Projections

*   username\_history\_read
    

### Events

*   user.username.changed.v1, user.username.request.approved.v1
    

1.12 Admin Impersonation with Audit
-----------------------------------

### Stories

*   As a support lead, I want impersonation with banners so that I can troubleshoot.
    
*   As a system, I want time-limited, fully logged sessions so that compliance is ensured.
    
*   As a user, I want notifications when impersonation occurs so that transparency is maintained.
    

### Flow

*   StartImpersonationCommand | StopImpersonationCommand →GetImpersonationSessionsQuery →StartImpersonation() | StopImpersonation()
    

### Projections

*   impersonation\_read
    

### Events

*   user.impersonation.started.v1, user.impersonation.stopped.v1
    

1.13 Referral Program Management
--------------------------------

### Stories

*   As a user, I want referral links and tracking so that I can earn incentives.
    
*   As a platform admin, I want referral fraud detection so that abuse is reduced.
    
*   As a system, I want payouts triggered by milestones so that incentives are automated.
    

### Flow

*   GenerateReferralLinkCommand | TrackReferralSignupCommand →GetReferralsQuery | GetReferralStatsQuery →GenerateReferralLink() | TrackReferralSignup()
    

### Projections

*   referrals\_read
    

### Events

*   user.referral.link.generated.v1, user.referral.signup.tracked.v1
    

1.14 Sharding & Migration
-------------------------

### Stories

*   As a platform engineer, I want regional/hash sharding so that we scale to millions.
    
*   As a system, I want zero-downtime migrations so that operations are uninterrupted.
    

### Flow

*   MigrateUserToShardCommand →GetUserShardQuery | ListShardMigrationJobsQuery →MigrateUserToShard()
    
1.15 Engagement Metrics (New Subsection)
----------------------------------------

**Stories**

*   **(New)** As a support agent, I want real-time user engagement tracking (e.g., live session activity) so that I can assist users proactively.
    

**Flow**

*   **(New)** GetRealTimeEngagementQuery → GetRealTimeEngagement()
    

**Projections**

*   **(New)** realtime\_engagement\_read

### Projections

*   shard\_migrations\_read
    

### Events

*   user.shard.migrated.v1
    

2) Profile
==========

2.1 Profile Completion & Update (+ Video Intro & Transcription)
---------------------------------------------------------------

### Stories

*   As a freelancer, I want to complete/update my profile so that clients can evaluate me.
    
*   As a user, I want real-time completeness % so that I know what to improve.
    
*   As a user, I want optional AI suggestions so that I can improve quickly.
    
*   **(Added)** As a freelancer, I want to upload a video intro so that I can personalize my profile.
    
*   **(Added)** As a system, I want video transcription for searchability so that AI/search can index it.
    
*   **(Added)** As a compliance officer, I want profile versioning so that all changes are auditable.

*   **(New)** As a freelancer, I want badges for profile completion milestones so that I’m motivated to improve my profile.
    
*   **(New)** As a freelancer, I want my profile dynamically personalized for clients so that relevant skills are highlighted.

### Flow

*   (Updated) UpdateProfileCommand | CalculateCompletenessCommand | UploadProfileVideoCommand | ProcessVideoTranscriptionCommand | LogProfileVersionCommand | AwardProfileCompletionBadgeCommand | PersonalizeProfileDisplayCommand → GetProfileByUserIDQuery | ListProfilesByCompletenessQuery | GetProfileVersionHistoryQuery | GetProfileCompletionMilestonesQuery | GetPersonalizedProfileQuery → UpdateProfile() | CalculateCompleteness() | UploadProfileVideo() | ProcessVideoTranscription() | LogProfileVersion() | AwardProfileCompletionBadge() | PersonalizeProfileDisplay()
    

### Projections

*   profiles\_read, profile\_completeness\_read, profile\_media\_read, profile\_transcripts\_read

*   **(New)** personalized\_profiles\_read

### Events

*   profile.updated.v1, profile.completeness.updated.v1, profile.video.uploaded.v1, profile.video.transcribed.v1, **profile.version.logged.v1**

*   **(New)** profile.completion.badge.awarded.v1, profile.personalization.applied.v1
    

2.2 Preferences & Visibility
----------------------------

### Stories

*   As a user, I want to set language/timezone so that my experience matches my locale.
    
*   As a user, I want visibility scopes (public/invited/private) with previews so that I understand exposure.
    

### Flow

*   UpdatePreferencesCommand | SetProfileVisibilityCommand →GetPublicProfileQuery | GetVisibilityPreviewQuery →UpdatePreferences() | SetProfileVisibility()
    

### Projections

*   profiles\_read.visibility
    

### Events

*   profile.preferences.updated.v1, profile.visibility.updated.v1
    

2.3 Locale Variants
-------------------

### Stories

*   As a user, I want multi-language profile content so that I can target different audiences.
    
*   As a system, I want auto-translation suggestions so that creating locales is easier.
    
*   As a client, I want to view in my preferred locale so that it’s accessible.
    

### Flow

*   AddProfileLocaleVariantCommand | RemoveProfileLocaleVariantCommand | SuggestTranslationCommand →GetProfileLocalesQuery | GetProfileInLocaleQuery →AddProfileLocaleVariant() | RemoveProfileLocaleVariant() | SuggestTranslation()
    

### Events

*   profile.locale.added.v1, profile.locale.removed.v1, profile.translation.suggested.v1
    

2.4 Snapshots & Restore
-----------------------

### Stories

*   As an admin, I want to snapshot/restore profiles so that I can revert bad or accidental changes.
    
*   As a user, I want self-restore so that I can fix mistakes.
    
*   As a system, I want auto-snapshots on major updates.
    

### Flow

*   CreateProfileSnapshotCommand | RestoreProfileFromSnapshotCommand →ListProfileSnapshotsQuery | GetProfileSnapshotQuery →CreateProfileSnapshot() | RestoreProfileFromSnapshot()
    

### Events

*   profile.snapshot.created.v1, profile.snapshot.restored.v1
    

2.5 AI Profile Optimization + AI Training Opt-Out
-------------------------------------------------

### Stories

*   As a user, I want AI optimization suggestions so that my visibility improves.
    
*   As a user, I want to control data usage and consent so that privacy is respected.
    
*   **(Added)** As a user, I want to opt-out of AI training so that my data is not used for model training.
    

### Flow

*   OptimizeProfileWithAICommand | ToggleAIOptOutCommand →GetProfileOptimizationSuggestionsQuery | GetAIOptOutStatusQuery →OptimizeProfileWithAI() | ToggleAIOptOut()
    

### Events

*   profile.ai.optimized.v1, profile.ai.optout.updated.v1
    

3) Skill
========

3.1 Manage Skills + Skill Tests
-------------------------------

### Stories

*   As a freelancer, I want to add/update/remove/reorder skills so that my expertise is clear.
    
*   As a freelancer, I want proficiency validated and duplicates prevented so that data is clean.
    
*   **(Added)** As a freelancer, I want to take skills tests so that my proficiency is certified.

*   **(New)** As a freelancer, I want to set skill-based pricing tiers (e.g., beginner/intermediate/expert) so that I can offer flexible pricing.

### Flow

*   (Updated) AddSkillCommand | UpdateSkillCommand | RemoveSkillCommand | ReorderSkillsCommand | TakeSkillTestCommand | SetSkillPricingTierCommand → GetSkillsByUserQuery | ListSkillsQuery | SearchSkillsQuery | GetSkillTestResultsQuery | GetSkillPricingTiersQuery → AddSkill() | UpdateSkill() | RemoveSkill() | ReorderSkills() | TakeSkillTest() | SetSkillPricingTier()
    

### Projections

*   skills\_read, skill\_tests\_read

*   **(New)** skill\_pricing\_read

### Events

*   skill.entry.added.v1, skill.entry.updated.v1, skill.entry.removed.v1, skill.entries.reordered.v1, skill.test.completed.v1

*   **(New)** skill.pricing.tier.set.v1
    

3.2 Normalize to Taxonomy
-------------------------

### Stories

*   As a platform admin, I want skills normalized to a taxonomy so that search is consistent.
    
*   As a system, I want batch normalization so that legacy data is fixed efficiently.
    

### Flow

*   NormalizeSkillCommand | NormalizeSkillsBatchCommand →GetSkillsByTaxonomyVersionQuery →NormalizeSkill() | NormalizeSkillsBatch()
    

### Events

*   skill.taxonomy.normalized.v1, skill.taxonomy.batch\_normalized.summary.v1
    

3.3 Endorsements (+ Abuse Limits)
---------------------------------

### Stories

*   As a client, I want to endorse a skill so that credible signals appear.
    
*   As a freelancer, I want moderation of endorsements so that spam is prevented.
    
*   As a system, I want endorsements to influence ranking so that relevance improves.
    
*   **(Added Policy)** Max endorsements per client per year = 10 (prevent gaming).
    

### Flow

*   EndorseSkillCommand | RevokeSkillEndorsementCommand →GetSkillEndorsementsQuery | ListEndorsedSkillsQuery →EndorseSkill() | RevokeSkillEndorsement() (enforce limit)
    

### Events

*   skill.endorsement.added.v1, skill.endorsement.revoked.v1
    

3.4 Taxonomy Sync
-----------------

### Stories

*   As a system, I want to sync skills to latest taxonomy so that indexing stays current.
    
*   As a user, I want notifications on deprecated skills.
    

### Flow

*   SyncSkillTaxonomyCommand →GetDeprecatedSkillsQuery | GetSkillsByTaxonomyVersionQuery →SyncSkillTaxonomy()
    

### Events

*   skill.taxonomy.synced.v1, skill.deprecated.notified.v1
    

3.5 Skill Proficiency Proof
---------------------------

### Stories

*   As a freelancer, I want to attach test scores as proof so that claims are substantiated.
    
*   As a client, I want verified proofs visible.
    

### Flow

*   AttachSkillProofCommand →GetSkillProofsQuery →AttachSkillProof()
    

### Events

*   skill.proof.attached.v1
    

3.6 Skill-Based Job Recommendations **(Added)**
-----------------------------------------------

### Stories

*   As a freelancer, I want skill-based job recommendations so that I find relevant opportunities.
    

### Flow

*   GetSkillBasedJobRecommendationsQuery → GetSkillBasedJobRecommendations()
    

### Projections

*   job\_recommendations\_read
    

4) Experience
=============

4.1 Manage Experiences (+ Client References)
--------------------------------------------

### Stories

*   As a freelancer, I want CRUD with valid ranges so that my timeline is credible.
    
*   As a system, I want overlap detection and sorting so that display is correct.
    
*   **(Added)** As a freelancer, I want to add client references/testimonials to experiences so that entries are verifiable.
    

### Flow

*   AddExperienceCommand | UpdateExperienceCommand | RemoveExperienceCommand | FlagCurrentExperienceCommand | ValidateDateRangesCommand | AddExperienceReferenceCommand →GetExperiencesByUserQuery | ListExperiencesQuery | SearchExperiencesQuery | GetExperienceReferencesQuery →AddExperience() | UpdateExperience() | RemoveExperience() | FlagCurrentExperience() | ValidateDateRanges() | AddExperienceReference()
    

### Events

*   experience.entry.added.v1, experience.entry.updated.v1, experience.entry.removed.v1, experience.entry.flagged\_current.v1, experience.reference.added.v1
    

4.2 Experience Verification
---------------------------

### Stories

*   As a user, I want employment verification so that clients trust my background.
    
*   As an admin, I want third-party checks so that verification is robust.
    

### Flow

*   VerifyExperienceCommand | RejectExperienceVerificationCommand →GetExperienceVerificationsQuery →VerifyExperience() | RejectExperienceVerification()
    

### Events

*   experience.verification.approved.v1, experience.verification.rejected.v1
    

4.3 Experience Gaps Analysis
----------------------------

### Stories

*   As a freelancer, I want gap analysis so that I can add clarifications.
    
*   As an admin, I want gap reports to flag suspicious timelines.
    

### Flow

*   AnalyzeExperienceGapsCommand →GetExperienceGapsQuery →AnalyzeExperienceGaps()
    

### Events

*   experience.gaps.analyzed.v1
    

5) Education
============

5.1 Manage Educations
---------------------

### Stories

*   As a freelancer, I want CRUD with year validation so that my qualifications are verifiable.
    
*   As a system, I want search and sort by year so that reviews are easy.
    

### Flow

*   AddEducationCommand | UpdateEducationCommand | RemoveEducationCommand →GetEducationsByUserQuery | ListEducationsQuery | SearchEducationsQuery →AddEducation() | UpdateEducation() | RemoveEducation()
    

### Events

*   education.entry.added.v1, education.entry.updated.v1, education.entry.removed.v1
    

5.2 Credential Registry Link
----------------------------

### Stories

*   As a user, I want to attach registry links so that third parties can validate degrees.
    
*   As a system, I want to verify links so that dead URLs are avoided.
    

### Flow

*   AttachCredentialRegistryLinkCommand →GetEducationLinksQuery →AttachCredentialRegistryLink()
    

### Events

*   education.credential.link.attached.v1
    

5.3 GPA/Transcript Attachment
-----------------------------

### Stories

*   As a freelancer, I want to attach transcripts securely so that details are shareable.
    
*   As a client, I want transcript metadata so that I can verify claims.
    

### Flow

*   AttachEducationTranscriptCommand →GetEducationAttachmentsQuery →AttachEducationTranscript()
    

### Events

*   education.transcript.attached.v1
    

6) Certification
================

6.1 Certification Lifecycle (+ Provider Import)
-----------------------------------------------

### Stories

*   As a freelancer, I want certifications tracked and verified so that clients see credible credentials.
    
*   As a user, I want renewal reminders so that I stay current.
    
*   As a system, I want expiries handled automatically so that statuses are accurate.
    
*   **(Added)** As a freelancer, I want auto-import from external providers (e.g., Coursera) so that updates are easy.
    

### Flow

*   AddCertificationCommand | VerifyCertificationCommand | RejectCertificationCommand | ExpireCertificationCommand | SendCertificationRenewalReminderCommand | ImportCertificationFromProviderCommand →ListCertificationsQuery | GetExpiredCertificationsQuery | GetUpcomingExpiriesQuery →AddCertification() | VerifyCertification() | RejectCertification() | ExpireCertification() | SendCertificationRenewalReminder() | ImportCertificationFromProvider()
    

### Events

*   certification.entry.added.v1, certification.verification.approved.v1, certification.verification.rejected.v1, certification.expired.v1, certification.renewal.reminder.sent.v1, certification.imported.v1
    

6.2 Proctoring & Provider Webhooks
----------------------------------

### Stories

*   As a system, I want to store proctoring metadata so that integrity is provable.
    
*   As a system, I want provider webhooks processed (HMAC) with retry/DLQ so that status is consistent.
    

### Flow

*   AttachProctoringMetadataCommand | ProcessCertificationWebhookCommand →GetCertificationProviderStatusQuery →AttachProctoringMetadata() | ProcessCertificationWebhook()
    

### Events

*   certification.proctoring.metadata.attached.v1, certification.provider.webhook.processed.v1
    

7) Portfolio
============

7.1 Portfolio Items & Media (+ Share Links)
-------------------------------------------

### Stories

*   As a freelancer, I want to add/update/remove/reorder items so that my work is showcased.
    
*   As a system, I want async media processing so that UX stays snappy.
    
*   **(Added)** As a freelancer, I want shareable public portfolio links so that I can promote outside the platform.
    
*   **(Added)** As a freelancer, I want to categorize portfolio items by type/industry so that clients can filter easily.
    

### Flow

*   AddPortfolioItemCommand | UpdatePortfolioItemCommand | RemovePortfolioItemCommand | ReorderPortfolioItemsCommand | ProcessMediaCommand | GeneratePortfolioShareLinkCommand | AddPortfolioItemCategoryCommand →GetPortfolioItemsByUserQuery | ListPortfolioItemsQuery | SearchPortfolioItemsQuery | GetPortfolioShareLinksQuery | GetPortfolioCategoriesQuery →AddPortfolioItem() | UpdatePortfolioItem() | RemovePortfolioItem() | ReorderPortfolioItems() | ProcessMedia() | GeneratePortfolioShareLink() | AddPortfolioItemCategory()
    

### Events

*   portfolio.item.added.v1, portfolio.item.updated.v1, portfolio.item.removed.v1, portfolio.items.reordered.v1, portfolio.media.processed.v1, portfolio.share.link.generated.v1, **portfolio.item.category.added.v1**
    

7.2 Safe-Media Pipeline
-----------------------

### Stories

*   As a moderation system, I want AV/DLP/policy scans so that unsafe media is blocked.
    
*   As a user, I want scan results and appeals so that false positives can be resolved.
    

### Flow

*   ScanPortfolioMediaCommand →GetMediaScanResultsQuery →ScanPortfolioMedia()
    

### Events

*   portfolio.media.scanned.v1, portfolio.media.blocked.v1
    

7.3 External Showcases
----------------------

### Stories

*   As a freelancer, I want to link GitHub/Behance/Dribbble so that my work stays in sync.
    
*   As a system, I want scheduled refreshes with change detection so that updates are automatic.
    

### Flow

*   AttachExternalShowcaseCommand | RefreshExternalShowcaseCommand →GetExternalShowcasesQuery →AttachExternalShowcase() | RefreshExternalShowcase()
    

### Events

*   portfolio.external.attached.v1, portfolio.external.refreshed.v1
    

7.4 Portfolio Analytics
-----------------------

### Stories

*   As a freelancer, I want views/engagement stats so that I can optimize.
    
*   As an analyst, I want time-series analytics so that growth can be measured.
    

### Flow

*   GetPortfolioAnalyticsQuery → GetPortfolioAnalytics()
    

### Projections

*   portfolio\_analytics\_read
    

8) Language
===========

8.1 Language Proficiency
------------------------

### Stories

*   As a freelancer, I want to manage languages with proficiency so that clients can filter multilingual talent.
    
*   As a system, I want CEFR-style enums and dedupe so that quality is high.
    

### Flow

*   AddLanguageCommand | UpdateLanguageCommand | RemoveLanguageCommand →GetLanguagesByUserQuery | ListLanguagesQuery | SearchLanguagesQuery →AddLanguage() | UpdateLanguage() | RemoveLanguage()
    

### Events

*   language.entry.added.v1, language.entry.updated.v1, language.entry.removed.v1
    

8.2 CEFR & Certificates
-----------------------

### Stories

*   As a user, I want CEFR mapping and certificate attachments so that my level is standardized.
    
*   As a system, I want link verification so that broken links are avoided.
    

### Flow

*   MapLanguageToCEFRCommand | AttachLanguageCertificateCommand →GetLanguageCertificationsQuery →MapLanguageToCEFR() | AttachLanguageCertificate()
    

### Events

*   language.cefr.mapped.v1, language.certificate.attached.v1
    

8.3 Language Fluency Tests
--------------------------

### Stories

*   As a freelancer, I want integrated fluency tests so that proficiency is measured.
    
*   As a client, I want to see verified fluency so that selection is easier.
    

### Flow

*   CompleteLanguageFluencyTestCommand →GetFluencyTestResultsQuery →CompleteLanguageFluencyTest()
    

### Events

*   language.fluency.tested.v1
    

9) Freelancer
=============

9.1 Profile, Rates, Stats, Earnings, Connects (+ JSS)
-----------------------------------------------------

### Stories

*   As a freelancer, I want to manage my profile, rates, stats, earnings, and connects so that I can present and price my services.
    
*   As a freelancer, I want connects deductions to fail on insufficient balance so that I don’t overspend.
    
*   As a system, I want rate history recorded so that trends are visible.
    
*   **(Added)** As a freelancer, I want Job Success Score (JSS) computed from history so that my ranking improves.
    
*   **(New)** As a freelancer, I want to apply for a vetted elite tier through a screening pipeline so that I can access premium opportunities.
    
### Flow

*   (Updated) UpdateFreelancerProfileCommand | UpdateRatesCommand | UpdateStatsCommand | UpdateEarningsCommand | AddConnectsCommand | DeductConnectsCommand | ComputeJSSCommand | ApplyForVettingCommand → GetFreelancerByUserQuery | GetFreelancerStatsQuery | GetConnectsBalanceQuery | GetJSSQuery | GetVettingStatusQuery → UpdateFreelancerProfile() | UpdateRates() | UpdateStats() | UpdateEarnings() | AddConnects() | DeductConnects() | ComputeJSS() | ApplyForVetting()
    

### Projections

*   freelancers\_read, freelancer\_stats\_read, connects\_read, freelancer\_jss\_read

*   **(New)** vetting\_status\_read

### Events

*   freelancer.profile.updated.v1, freelancer.rates.updated.v1, freelancer.stats.updated.v1, freelancer.earnings.updated.v1, freelancer.connects.added.v1, freelancer.connects.deducted.v1, freelancer.jss.updated.v1

*   **(New)** freelancer.vetting.applied.v1, freelancer.vetting.approved.v1

9.2 Open-to-Work & Calendar
---------------------------

### Stories

*   As a freelancer, I want to toggle Open-to-Work and sync availability so that clients know when I’m available.
    
*   As a system, I want Google/Outlook sync so that calendars stay aligned.
    

### Flow

*   SetOpenToWorkCommand | SyncAvailabilityCalendarCommand →GetOpenToWorkStateQuery | GetAvailabilityCalendarQuery →SetOpenToWork() | SyncAvailabilityCalendar()
    

### Events

*   freelancer.open\_to\_work.updated.v1, freelancer.calendar.synced.v1
    

9.3 Rate Bands & Auto Top-Up
----------------------------

### Stories

*   As a freelancer, I want rate bands (hour/day/retainer) so that pricing is flexible.
    
*   As a freelancer, I want connects auto top-up so that proposals don’t get blocked.
    
*   As a system, I want auto-top-up via financial-be when balance dips below threshold.
    

### Flow

*   UpdateRateBandsCommand | ConfigureConnectsAutoTopUpCommand →GetRateBandsQuery | GetAutoTopUpConfigQuery →UpdateRateBands() | ConfigureConnectsAutoTopUp()
    

### Events

*   freelancer.rate\_bands.updated.v1, freelancer.connects.autotopup.configured.v1
    

9.4 Freelancer Tiering (+ Freelancer Plus)
------------------------------------------

### Stories

*   As a system, I want to assign tiers (basic/pro/elite) from stats so that benefits unlock.
    
*   As a freelancer, I want to know my tier so that I understand eligibility.
    
*   **(Added)** As a freelancer, I want to upgrade to Plus (subscription) for extra connects and premium benefits (e.g., reduced fees/direct contracts) so that I can compete better.
    

### Flow

*   AssignFreelancerTierCommand | UpgradeToPlusCommand (subscriptions-be) →GetFreelancerTierQuery | GetTierBenefitsQuery →AssignFreelancerTier() | UpgradeToPlus()
    

### Events

*   freelancer.tier.assigned.v1, freelancer.tier.plus.upgraded.v1
    

9.5 Proposal Analytics **(Added)**
----------------------------------

### Stories

*   As a freelancer, I want proposal analytics (views, acceptance rates) so that I can optimize submissions.

*   **(New)** As a freelancer, I want to see client-provided rejection reasons for my proposals so that I can improve.

### Flow

*   (Updated) GetProposalAnalyticsQuery | RecordProposalRejectionReasonCommand → GetProposalAnalytics() | RecordProposalRejectionReason() → GetProposalRejectionReasonsQuery

**Events**

*   **(New)** proposal.rejection.reason.recorded.v1

### Projections

*   proposal\_analytics\_read
    

9.6 Direct Contracts for Plus/Elite **(Added)**
-----------------------------------------------

### Stories

*   As a Plus/elite freelancer, I want no-fee direct contracts so that I keep more earnings.
    

### Flow

*   CreateDirectContractCommand (integrate contracts-be, check tier) →GetDirectContractEligibilityQuery →CreateDirectContract()
    

### Events

*   freelancer.contract.direct.created.v1

9.7 Real-Time Collaboration Tools (New Subsection)
--------------------------------------------------

**Stories**

*   **(New)** As a freelancer, I want real-time collaboration tools (e.g., live document editing, chat) so that I can work with clients efficiently.
    

**Flow**

*   **(New)** InitiateCollaborationSessionCommand → GetCollaborationSessionQuery → InitiateCollaborationSession()
    

**Events**

*   **(New)** freelancer.collaboration.session.started.v1
    

**Projections**

*   **(New)** collaboration\_sessions\_read

9.9 Subscription Bundles for Freelancers (New Subsection)
----------------------------------------------------------

**Stories**

*   **(New)** As a freelancer, I want bundled subscription plans (e.g., Plus + analytics) so that I get premium features at a discount.
    

**Flow**

*   **(New)** SubscribeToBundleCommand → GetBundleSubscriptionStatusQuery → SubscribeToBundle() (integrates with subscriptions-be)
    

**Events**

*   **(New)** freelancer.subscription.bundle.subscribed.v1
    

**Projections**

*   **(New)** subscription\_bundles\_read

10) Client
==========

10.1 Client Profile, Spending & Stats (+ Payment Methods)
---------------------------------------------------------

### Stories

*   As a client, I want to manage company profile so that freelancers understand us.
    
*   As a client, I want spending/hiring stats so that I can plan.
    
*   As a client, I want budget caps so that spending is controlled.
    
*   **(Added)** As a client, I want to add/verify payment methods so that I receive a verified payment badge.

*   **(New)** As a client, I want a dashboard with job performance analytics (e.g., hire rates, response times) so that I can optimize postings.

### Flow

*   UpdateClientProfileCommand | UpdateCompanyDetailsCommand | UpdateClientStatsCommand | UpdateSpendingCommand | AddPaymentMethodCommand →GetClientByUserQuery | GetClientStatsQuery | VerifyPaymentMethodQuery →UpdateClientProfile() | UpdateCompanyDetails() | UpdateClientStats() | UpdateSpending() | AddPaymentMethod() | GetClientJobAnalytics()
    

### Projections

*   clients\_read, client\_stats\_read, client\_payment\_methods\_read

*   **(New)** client\_job\_analytics\_read

### Events

*   client.profile.updated.v1, client.company.updated.v1, client.stats.updated.v1, client.spending.updated.v1, client.payment.method.added.v1
    

10.2 Verified Payment & Hire Visibility
---------------------------------------

### Stories

*   As a client, I want a verified payment badge so that freelancers trust us.
    
*   As an org admin, I want org-wide hire visibility so that teams coordinate invites.
    

### Flow

*   SetVerifiedPaymentStatusCommand | EnableOrgHireVisibilityCommand →GetPaymentVerificationStatusQuery →SetVerifiedPaymentStatus() | EnableOrgHireVisibility()
    

### Events

*   client.payment.verified.v1 (and client.payment.unverified.v1), client.org.visibility.enabled.v1
    

10.3 Budget Forecasting
-----------------------

### Stories

*   As a client, I want spending forecasts from active contracts so that budgeting is accurate.
    
*   As an analyst, I want forecast reports so that finance can plan.
    
*   **(Added)** As a client, I want budget overrun alerts so that I can control spending.
    

### Flow

*   GetSpendingForecastQuery | SendBudgetOverrunAlertCommand →GetSpendingForecast() | SendBudgetOverrunAlert()
    

### Projections

*   client\_budget\_forecast\_read
    

### Events

*   **client.budget.overrun.alerted.v1**
    

10.4 Client Subscription Tiers **(Added)**
------------------------------------------

### Stories

*   As a client, I want to subscribe to a premium tier for priority job posting and higher budgets.
    

### Flow

*   UpgradeClientToPremiumCommand (subscriptions-be) →GetClientTierQuery →UpgradeClientToPremium()
    

### Events

*   client.tier.upgraded.v1
    
10.5 Client-Freelancer Matchmaking Feedback (New Subsection)
------------------------------------------------------------

**Stories**

*   **(New)** As a client, I want to provide feedback on AI matchmaking results so that recommendations improve.
    

**Flow**

*   **(New)** SubmitMatchmakingFeedbackCommand → GetMatchmakingFeedbackQuery → SubmitMatchmakingFeedback()
    

**Events**

*   **(New)** client.matchmaking.feedback.submitted.v1
    

**Projections**

*   **(New)** matchmaking\_feedback\_read

10.6 Enterprise SLA Guarantees (New Subsection)
-----------------------------------------------

**Stories**

*   **(New)** As an enterprise client, I want SLA guarantees (e.g., response times, match quality) so that my needs are prioritized.
    

**Flow**

*   **(New)** ConfigureEnterpriseSLACommand → GetSLAStatusQuery → ConfigureEnterpriseSLA()
    

**Events**

*   **(New)** client.sla.configured.v1
    

**Projections**

*   **(New)** sla\_configs\_read

11) Verification (KYC/KYB)
==========================

11.1 Submit & Review (+ Sanctions/OFAC/PEP)
-------------------------------------------

### Stories

*   As a user, I want to submit KYC docs so that I can unlock trusted features.
    
*   As a compliance reviewer, I want to approve/reject (batch) so that trust is enforced.
    
*   As a system, I want encrypted storage via storage-be so that privacy is protected.
    
*   **(Added)** As a compliance system, I want sanctions/OFAC/PEP screening so that risky users are flagged.

*   **(New)** As a system, I want to validate VAT/GST for international users so that cross-border tax compliance is ensured.

### Flow

*   SubmitVerificationCommand | ApproveVerificationCommand | RejectVerificationCommand | BatchApproveVerificationsCommand | ListVerificationsQuery | ScreenForSanctionsCommand | ValidateInternationalTaxCommand → SubmitVerification() | ApproveVerification() | RejectVerification() | BatchApproveVerifications() | ScreenForSanctions() | ValidateInternationalTax() → GetScreeningResultsQuery | GetInternationalTaxStatusQuery
    

### Projections

*   kyc\_read, kyb\_read, verification\_audit\_read, sanctions\_screening\_read
    

### Events

*   verification.submitted.v1, verification.approved.v1, verification.rejected.v1, verification.batch\_approved.summary.v1, verification.sanctions.flagged.v1

*   **(New)** verification.tax.international.validated.v1


### RBAC

*   role:COMPLIANCE\_REVIEWER
    

11.2 Trigger Re-Verification
----------------------------

### Stories

*   As a system, I want re-verification on expiry/risk so that statuses stay current.
    
*   As a compliance reviewer, I want a queue of candidates so that workloads are clear.
    

### Flow

*   TriggerReVerificationCommand →GetReverificationCandidatesQuery →TriggerReVerification()
    

### Events

*   verification.reverify.triggered.v1
    

11.3 Risk-Tiered KYC
--------------------

### Stories

*   As a risk engine, I want to set KYC tier so that checks match risk.
    
*   As a system, I want automatic tier upgrades on high-value actions.
    

### Flow

*   SetKYCTierCommand →GetKYCTierQuery →SetKYCTier()
    

### Events

*   verification.tier.updated.v1
    

11.4 Liveness Check
-------------------

### Stories

*   As a user, I want a liveness check so that spoofing is prevented.
    
*   As a system, I want integration with external liveness APIs.
    

### Flow

*   RecordLivenessCheckCommand →GetLivenessResultsQuery →RecordLivenessCheck()
    

### Events

*   verification.liveness.recorded.v1
    

11.5 KYB for Orgs
-----------------

### Stories

*   As an org admin, I want KYB submission so that enterprise features unlock.
    
*   As a compliance reviewer, I want KYB status tracking.
    

### Flow

*   SubmitKYBCommand →GetKYBStatusQuery →SubmitKYB()
    

### Events

*   verification.kyb.submitted.v1
    

11.6 Automated Tax Withholding Verification **(Added)**
-------------------------------------------------------

### Stories

*   As a system, I want automated tax withholding verification (e.g., TIN matching) so that tax compliance is ensured.
    

### Flow

*   VerifyTaxWithholdingCommand →GetTaxWithholdingStatusQuery →VerifyTaxWithholding()
    

### Events

*   verification.tax.withholding.verified.v1
    
11.7 Compliance Certifications (New Subsection)
-----------------------------------------------

**Stories**

*   **(New)** As an org admin, I want a SOC 2/ISO 27001 compliance badge so that clients trust my organization.
    

**Flow**

*   **(New)** AssignComplianceBadgeCommand → GetComplianceBadgeStatusQuery → AssignComplianceBadge()
    

**Events**

*   **(New)** verification.compliance.badge.assigned.v1
    

**Projections**

*   **(New)** compliance\_badges\_read

12) Settings
============

12.1 Personalization & Privacy
------------------------------

### Stories

*   As a user, I want to set theme/language/timezone so that my UI is comfortable.
    
*   As a user, I want privacy controls so that I manage my exposure.
    

### Flow

*   UpdateSettingsCommand | UpdatePrivacySettingsCommand →GetSettingsByUserQuery | ListSettingsQuery →UpdateSettings() | UpdatePrivacySettings()
    

### Events

*   settings.updated.v1
    

12.2 Notification Preferences
-----------------------------

### Stories

*   As a user, I want notification preferences so that I choose how I’m alerted.
    

### Flow

*   UpdateNotificationPrefsCommand →GetNotificationPrefsQuery →UpdateNotificationPrefs()
    

### Events

*   settings.notification\_prefs.updated.v1
    

12.3 Channel Matrix per Event
-----------------------------

### Stories

*   As a user, I want per-event channel routing (email/SMS/push/in-app/webhook) so that control is granular.
    
*   As a developer, I want webhook channel selection so that integrations work.
    

### Flow

*   UpdateChannelMatrixCommand →GetChannelMatrixQuery →UpdateChannelMatrix()
    

### Events

*   settings.channels.updated.v1
    

12.4 Data Residency Preference
------------------------------

### Stories

*   As a user, I want to pick a data residency (EU/US) so that compliance is met.
    
*   As a system, I want migration on change.
    

### Flow

*   SetDataResidencyPreferenceCommand →GetDataResidencyQuery →SetDataResidencyPreference()
    

### Events

*   settings.data\_residency.updated.v1
    

12.5 Accessibility Settings
---------------------------

### Stories

*   As a user, I want high-contrast/screen-reader prefs so that the platform is accessible.
    

### Flow

*   UpdateAccessibilitySettingsCommand →GetAccessibilitySettingsQuery →UpdateAccessibilitySettings()
    

### Events

*   settings.accessibility.updated.v1
    

13) Saved Items
===============

13.1 Save / Unsave / Notes
--------------------------

### Stories

*   As a user, I want to save/unsave jobs or freelancers so that I can revisit later.
    
*   As a user, I want to add notes so that context is remembered.
    
*   As a system, I want duplicate prevention so that lists are clean.
    

### Flow

*   SaveItemCommand | UnsaveItemCommand | UpdateSavedItemNoteCommand →GetSavedItemsByUserQuery | ListSavedItemsQuery | SearchSavedItemsQuery →SaveItem() | UnsaveItem() | UpdateSavedItemNote()
    

### Events

*   saved.item.saved.v1, saved.item.unsaved.v1, saved.item.note.updated.v1
    

13.2 Collections & Sharing
--------------------------

### Stories

*   As an org member, I want collections and sharing/permissions so that my team collaborates.
    
*   As a user, I want to move items between collections.
    

### Flow

*   CreateSavedCollectionCommand | MoveSavedItemToCollectionCommand | ShareSavedCollectionCommand →GetSavedCollectionsQuery →CreateSavedCollection() | MoveSavedItemToCollection() | ShareSavedCollection()
    

### Events

*   saved.collection.created.v1, saved.item.moved.v1, saved.collection.shared.v1
    

13.3 Expiry & Reminders
-----------------------

### Stories

*   As a user, I want expiries/reminders so that my list stays fresh.
    
*   As a system, I want scheduled reminders.
    

### Flow

*   SetSavedItemExpiryCommand | SendSavedItemReminderCommand →GetExpiringSavedItemsQuery →SetSavedItemExpiry() | SendSavedItemReminder()
    

### Events

*   saved.item.expiry.set.v1, saved.item.reminder.sent.v1
    

14) Blocked Users
=================

14.1 Block / Unblock
--------------------

### Stories

*   As a user, I want to block/unblock with a reason so that I control interactions.
    
*   As a system, I want self-block prevention.
    

### Flow

*   BlockUserCommand | UnblockUserCommand →GetBlockedUsersByUserQuery | IsBlockedQuery →BlockUser() | UnblockUser()
    

### Events

*   user.blocked.v1, user.unblocked.v1
    

14.2 Scoped & Time-Bound
------------------------

### Stories

*   As a user, I want scoped (messaging/invites/full) and time-bound blocks so that restrictions fit my needs.
    
*   As a system, I want auto-unblock on expiry.
    

### Flow

*   BlockUserScopedCommand | ExtendBlockDurationCommand →GetBlockScopesQuery →BlockUserScoped() | ExtendBlockDuration()
    

### Events

*   user.block.scoped.v1, user.block.duration.extended.v1
    

14.3 Block Appeals
------------------

### Stories

*   As a blocked user, I want to appeal a block so that mistakes get resolved.
    

### Flow

*   AppealBlockCommand →GetBlockAppealsQuery →AppealBlock()
    

### Events

*   user.block.appealed.v1
    

15) User Suspension
===================

15.1 Place / Release / Extend
-----------------------------

### Stories

*   As a moderator, I want to place/release/extend suspensions so that policy is enforced.
    
*   As a system, I want history logged.
    

### Flow

*   SuspendUserCommand | ReleaseSuspensionCommand | ExtendSuspensionCommand | AddSuspensionHistoryCommand →GetSuspensionsByUserQuery | ListActiveSuspensionsQuery | GetSuspensionHistoryQuery →SuspendUser() | ReleaseSuspension() | ExtendSuspension() | AddSuspensionHistory()
    

### Events

*   user.suspension.placed.v1, user.suspension.released.v1, user.suspension.extended.v1, user.suspension.history.added.v1
    

15.2 Schedule / Cancel
----------------------

### Stories

*   As a moderator, I want to schedule/cancel future suspensions so that grace periods are honored.
    

### Flow

*   ScheduleSuspensionCommand | CancelScheduledSuspensionCommand →GetScheduledSuspensionsQuery →ScheduleSuspension() | CancelScheduledSuspension()
    

### Events

*   user.suspension.scheduled.v1, user.suspension.canceled.v1
    

15.3 Auto-Suspend Matrix
------------------------

### Stories

*   As a system, I want auto-suspend based on warnings+risk with grace so that enforcement is consistent.
    

### Flow

*   AutoSuspendUserCommand →GetAutoSuspendCandidatesQuery →AutoSuspendUser()
    

### Events

*   user.suspension.auto.placed.v1
    

15.4 Suspension Appeals
-----------------------

### Stories

*   As a user, I want to appeal suspensions with evidence so that fairness is preserved.
    

### Flow

*   AppealSuspensionCommand →GetSuspensionAppealsQuery →AppealSuspension()
    

### Events

*   user.suspension.appealed.v1
    

16) User Ban
============

16.1 Ban / Release / History
----------------------------

### Stories

*   As a trust & safety admin, I want to ban/release users and log history so that severe cases are handled.
    
*   As a system, I want expiry for temporary bans.
    

### Flow

*   BanUserCommand | ReleaseBanCommand | AddBanHistoryCommand →GetBansByUserQuery | ListActiveBansQuery | GetBanHistoryQuery →BanUser() | ReleaseBan() | AddBanHistory()
    

### Events

*   user.ban.placed.v1, user.ban.released.v1, user.ban.history.added.v1
    

16.2 Shadow-Ban
---------------

### Stories

*   As a moderator, I want to enable/disable shadow-ban so that spam is dampened during investigations.
    

### Flow

*   EnableShadowBanCommand | DisableShadowBanCommand →GetShadowBanStatusQuery →EnableShadowBan() | DisableShadowBan()
    

### Events

*   user.shadowban.enabled.v1, user.shadowban.disabled.v1
    

16.3 Ban Evasion Detection
--------------------------

### Stories

*   As a system, I want to detect ban evasion via IP/device so that repeat offenders are caught.
    

### Flow

*   DetectBanEvasionCommand →GetEvasionLinksQuery →DetectBanEvasion()
    

### Events

*   user.ban.evasion.detected.v1
    

17) User Warning
================

17.1 Issue / Acknowledge / Escalate
-----------------------------------

### Stories

*   As a moderator, I want to issue warnings so that users can correct behavior.
    
*   As a user, I want to acknowledge warnings so that my record is accurate.
    
*   As a system, I want escalation on repeated warnings so that enforcement is predictable.
    

### Flow

*   IssueWarningCommand | AcknowledgeWarningCommand | EscalateWarningsCommand →GetWarningsByUserQuery | ListUnacknowledgedWarningsQuery | GetWarningHistoryQuery →IssueWarning() | AcknowledgeWarning() | EscalateWarnings()
    

### Events

*   user.warning.issued.v1, user.warning.acknowledged.v1, user.warning.escalated.v1
    

17.2 Acknowledgement SLA & Appeal Link
--------------------------------------

### Stories

*   As a system, I want an acknowledgement deadline so that users respond in time.
    
*   As a moderator, I want to link an appeal ticket so that due process is auditable.
    

### Flow

*   SetWarningAcknowledgementDeadlineCommand | LinkWarningAppealTicketCommand →GetWarningSLAsQuery →SetWarningAcknowledgementDeadline() | LinkWarningAppealTicket()
    

### Events

*   user.warning.deadline.set.v1, user.warning.appeal.linked.v1
    

17.3 Warning Decay
------------------

### Stories

*   As a system, I want warnings to decay over time so that good behavior is rewarded.
    

### Flow

*   DecayWarningsCommand →GetDecayingWarningsQuery →DecayWarnings()
    

### Events

*   user.warning.decayed.v1
    

18) Org (Agencies/Companies)
============================

18.1 Org Lifecycle & Membership (+ Shared Billing)
--------------------------------------------------

### Stories

*   As an org admin, I want to create/update an org so that my team can collaborate.
    
*   As an org admin, I want to manage members/roles/invites/seats so that permissions and capacity are controlled.
    
*   As a member, I want to accept invites so that I can join.
    
*   **(Added)** As an org admin, I want shared billing integration so that seats and hires are charged centrally.
    

### Flow

*   CreateOrgCommand | UpdateOrgCommand | AddMemberCommand | RemoveMemberCommand | InviteMemberCommand | ManageSeatsCommand | LinkOrgBillingCommand →GetOrgByIDQuery | GetOrgsByUserQuery | ListMembersQuery | GetSeatUsageQuery | GetOrgBillingStatusQuery →CreateOrg() | UpdateOrg() | AddMember() | RemoveMember() | InviteMember() | ManageSeats() | LinkOrgBilling()
    

### Events

*   org.account.created.v1, org.account.updated.v1, org.member.added.v1, org.member.removed.v1, org.invite.sent.v1, org.seats.updated.v1, org.billing.linked.v1
    

18.2 Seat Policy Templates
--------------------------

### Stories

*   As an org admin, I want seat policy templates so that permissions are consistent.
    

### Flow

*   ApplySeatPolicyTemplateCommand →GetSeatPoliciesQuery →ApplySeatPolicyTemplate()
    

### Events

*   org.policy.applied.v1
    

18.3 Talent Pools
-----------------

### Stories

*   As a recruiter, I want to create talent pools and add users so that hiring pipelines are organized.

*   **(New)** As an org admin, I want AI-curated talent pool recommendations so that I can quickly find vetted freelancers.

### Flow

*   CreateOrgTalentPoolCommand | AddUserToTalentPoolCommand →ListTalentPoolsQuery | GetTalentPoolMembersQuery →CreateOrgTalentPool() | AddUserToTalentPool()
    
*   **(New)** GetAITalentPoolRecommendationsQuery → GetAITalentPoolRecommendations() (integrates with 9.1 JSS, vetting)

### Events

*   org.talent\_pool.created.v1, org.talent\_pool.member.added.v1
    
**Projections**

*   **(New)** talent\_pool\_recommendations\_read

18.4 Org Hierarchy
------------------

### Stories

*   As an enterprise admin, I want parent/child org hierarchies so that conglomerates are modeled.
    

### Flow

*   CreateOrgHierarchyCommand | LinkChildOrgCommand →GetOrgHierarchyQuery →CreateOrgHierarchy() | LinkChildOrg()
    

### Events

*   org.hierarchy.created.v1, org.child.linked.v1
    

19) Security Center
===================

19.1 Two-Factor, Devices & Sessions
-----------------------------------

### Stories

*   As a user, I want to enable/disable 2FA so that my account is secure.
    
*   As a user, I want to register devices and revoke sessions so that I control access.
    
*   As a system, I want recovery keys so that lockouts are avoided.
    

### Flow

*   Enable2FACommand | Disable2FACommand | RegisterDeviceCommand | RevokeSessionCommand | GenerateRecoveryKeysCommand →GetSecuritySettingsQuery | ListDevicesQuery | ListActiveSessionsQuery →Enable2FA() | Disable2FA() | RegisterDevice() | RevokeSession() | GenerateRecoveryKeys()
    

### Events

*   security.2fa.enabled.v1, security.2fa.disabled.v1, security.device.registered.v1, security.session.revoked.v1, security.recovery.keys.generated.v1
    

19.2 Device Fingerprint & Anomalies
-----------------------------------

### Stories

*   As a system, I want device fingerprints so that risky device reuse is tracked.
    
*   As a system, I want impossible-travel detection so that anomalies are flagged.
    

### Flow

*   RecordDeviceFingerprintCommand | RecordLoginAnomalyCommand →GetDeviceFingerprintsQuery | GetAnomalyLogsQuery →RecordDeviceFingerprint() | RecordLoginAnomaly()
    

### Events

*   security.device.fingerprint.recorded.v1, security.anomaly.detected.v1
    

19.3 Passkeys
-------------

### Stories

*   As a user, I want passwordless passkeys so that login is secure and easy.
    

### Flow

*   RegisterPasskeyCommand | AuthenticateWithPasskeyCommand →GetPasskeysQuery →RegisterPasskey() | AuthenticateWithPasskey()
    

### Events

*   security.passkey.registered.v1, security.passkey.authenticated.v1
    

20) Compliance
==============

20.1 Tax, Residency & Artifacts (+ Tax Form Validation)
-------------------------------------------------------

### Stories

*   As a freelancer, I want to submit tax profiles and residency so that I’m compliant.
    
*   As a user, I want to upload compliance artifacts so that audits pass.
    
*   As a system, I want GDPR/CCPA-friendly storage so that regulations are obeyed.
    
*   **(Added)** As a freelancer, I want tax form validation (W-9/W-8BEN) so that submissions are error-free.
    
*   **(New)** As a compliance officer, I want configurable audit log retention policies so that regulatory requirements are met.
    

### Flow

*   UpdateTaxProfileCommand | UpdateResidencyCommand | SubmitComplianceArtifactCommand | ValidateTaxFormCommand | SetAuditLogRetentionPolicyCommand → GetComplianceByUserQuery | GetTaxProfileQuery | ListComplianceArtifactsQuery | GetAuditLogRetentionPolicyQuery → UpdateTaxProfile() | UpdateResidency() | SubmitComplianceArtifact() | ValidateTaxForm() | SetAuditLogRetentionPolicy()
    

### Events

*   compliance.tax.updated.v1, compliance.residency.updated.v1, compliance.artifact.submitted.v1, compliance.taxform.validated.v1
    
*   **(New)** compliance.audit.retention.set.v1


20.2 GDPR/CCPA Export
---------------------

### Stories

*   As a user, I want a full data export so that I can receive my data.
    
*   As a compliance officer, I want export status so that I can track completion.
    

### Flow

*   BulkExportDataCommand →GetAuditLogsQuery | GetExportStatusQuery →ExportData()
    

### Events

*   compliance.export.requested.v1, compliance.export.completed.v1
    

20.3 Consent Management & Receipts
----------------------------------

### Stories

*   As a user, I want to grant/revoke consents so that data usage is under my control.
    
*   As a user, I want signed consent receipts so that changes are auditable.
    

### Flow

*   UpdateConsentCommand | GenerateConsentReceiptCommand →GetConsentsQuery →UpdateConsent() | GenerateConsentReceipt()
    

### Events

*   compliance.consent.updated.v1, compliance.receipt.generated.v1
    

20.4 GDPR Deletion Orchestration
--------------------------------

### Stories

*   As a user, I want to request account deletion so that my personal data is erased.
    
*   As a system, I want to orchestrate deletion across services so that consistency is achieved.
    

### Flow

*   StartGDPRDeletionCommand | CompleteGDPRDeletionCommand →GetDeletionStatusQuery →StartGDPRDeletion() | CompleteGDPRDeletion()
    

### Events

*   compliance.deletion.started.v1, compliance.deletion.completed.v1
    

20.5 Audit Trail Export
-----------------------

### Stories

*   As a compliance officer, I want audit trail exports (JSON/CSV) so that regulators can review.
    

### Flow

*   ExportAuditTrailCommand →GetAuditTrailQuery →ExportAuditTrail()
    

### Events

*   compliance.audit.exported.v1
    
20.6 Data Breach Notification Workflow (New Subsection)
-------------------------------------------------------

**Stories**

*   **(New)** As a compliance officer, I want an automated data breach notification system so that users are informed per GDPR/CCPA.
    

**Flow**

*   **(New)** TriggerBreachNotificationCommand → GetBreachNotificationStatusQuery → TriggerBreachNotification()
    

**Events**

*   **(New)** compliance.breach.notification.sent.v1
    

**Projections**

*   **(New)** breach\_notifications\_read

21) Risk Signals
================

21.1 Signals, Scoring & Holds
-----------------------------

### Stories

*   As a risk engine, I want to ingest signals and compute scores so that platform risk is managed.
    
*   As a moderator, I want to place/release risk holds so that funds or features are protected.
    

### Flow

*   RecordRiskSignalCommand | ComputeRiskScoreCommand | PlaceHoldCommand | ReleaseHoldCommand →GetRiskSignalsByUserQuery | GetRiskScoreQuery | ListHighRiskUsersQuery →RecordRiskSignal() | ComputeRiskScore() | PlaceHold() | ReleaseHold()
    

### Events

*   risk.signal.recorded.v1, risk.score.updated.v1, risk.hold.placed.v1, risk.hold.released.v1
    

21.2 Explainable Score
----------------------

### Stories

*   As a reviewer, I want explainable risk components so that I understand decisions.
    

### Flow

*   ComputeExplainableRiskScoreCommand →GetRiskExplanationQuery →ComputeExplainableRiskScore()
    

### Events

*   risk.score.explainable.v1
    

21.3 Risk Rate-Limit Policy
---------------------------

### Stories

*   As a system, I want stricter write rate-limits at high risk so that abuse is mitigated.
    

### Flow

*   ApplyRiskRateLimitPolicyCommand →GetRiskPoliciesQuery →ApplyRiskRateLimitPolicy()
    

### Events

*   risk.policy.applied.v1
    

21.4 Signal Correlation
-----------------------

### Stories

*   As a risk analyst, I want correlated signals (IP+device+payment) so that fraud rings are detected.
    

### Flow

*   CorrelateRiskSignalsCommand →GetCorrelatedSignalsQuery →CorrelateRiskSignals()
    

### Events

*   risk.signals.correlated.v1
    

21.5 Real-Time Risk Alerts **(Added)**
--------------------------------------

### Stories

*   As a moderator, I want real-time risk alerts so that I can act on high-risk actions immediately.
    

### Flow

*   SendRiskAlertCommand →GetRealTimeRiskAlertsQuery →SendRiskAlert()
    

### Events

*   risk.alert.sent.v1
    

22) Profile Depth
=================

22.1 Rates, Availability, Badges, Normalized Skills
---------------------------------------------------

### Stories

*   As a freelancer, I want rate history and availability schedules so that my market presence is clear.
    
*   As a system, I want normalized skills so that search is consistent.
    
*   As a client, I want badges to reflect achievements.
    

### Flow

*   UpdateRateHistoryCommand | UpdateAvailabilityScheduleCommand | AssignBadgeCommand | RemoveBadgeCommand | NormalizeSkillsCommand →GetRateHistoryQuery | GetAvailabilityScheduleQuery | ListBadgesByUserQuery | GetNormalizedSkillsQuery →UpdateRateHistory() | UpdateAvailabilitySchedule() | AssignBadge() | RemoveBadge() | NormalizeSkills()
    

### Events

*   profile.rate\_history.updated.v1, profile.availability.updated.v1, profile.badge.assigned.v1, profile.badge.removed.v1, profile.skills.normalized.v1
    

22.2 Milestone Badges
---------------------

### Stories

*   As a system, I want to assign/revoke milestone badges via a versioned engine so that recognition is consistent.
    

### Flow

*   AssignMilestoneBadgeCommand | RevokeMilestoneBadgeCommand →GetMilestoneCriteriaQuery →AssignMilestoneBadge() | RevokeMilestoneBadge()
    

### Events

*   profile.badge.milestone.assigned.v1, profile.badge.milestone.revoked.v1
    

22.3 Availability Exceptions
----------------------------

### Stories

*   As a freelancer, I want holiday/leave exceptions so that my schedule reflects reality.
    

### Flow

*   AddAvailabilityExceptionCommand | RemoveAvailabilityExceptionCommand →GetAvailabilityExceptionsQuery →AddAvailabilityException() | RemoveAvailabilityException()
    

### Events

*   profile.availability.exception.added.v1, profile.availability.exception.removed.v1
    

22.4 Depth Scoring
------------------

### Stories

*   As a system, I want a depth score (content/verifications/history) so that search ranking is improved.
    

### Flow

*   ComputeProfileDepthScoreCommand →GetProfileDepthScoreQuery →ComputeProfileDepthScore()
    

### Events

*   profile.depth.score.updated.v1
    

23) Cross-Cutting (All Domains)
===============================

23.1 Webhooks (HMAC) & Personal Access Tokens (+ Event Filters)
---------------------------------------------------------------

### Stories

*   As a developer, I want to register secure webhooks so that I can integrate.
    
*   As a developer, I want PATs so that I can access APIs programmatically.
    
*   **(Added)** As a developer, I want event filters/subscriptions so that I only receive relevant events.
    
*   **(Added)** As a developer, I want webhook delivery status logs so that I can debug failures.
    

### Flow

*   RegisterWebhookEndpointCommand | DisableWebhookEndpointCommand | CreatePersonalAccessTokenCommand | RevokePersonalAccessTokenCommand | UpdateWebhookFiltersCommand →ListWebhooksQuery | ListPATsQuery | GetWebhookDeliveryStatusQuery →RegisterWebhookEndpoint() | DisableWebhookEndpoint() | CreatePersonalAccessToken() | RevokePersonalAccessToken() | UpdateWebhookFilters() | GetWebhookDeliveryStatus()
    

### Projections

*   **webhook\_delivery\_read**
    

### Events

*   webhook.endpoint.registered.v1, webhook.endpoint.disabled.v1, pat.created.v1, pat.revoked.v1, webhook.filters.updated.v1
    

23.2 PII Encryption & Key Rotation
----------------------------------

### Stories

*   As a compliance officer, I want field-level encryption with rotation so that PII is protected.
    

### Flow

*   RotatePIIEncryptionKeysCommand →GetEncryptionStatusQuery →RotatePIIEncryptionKeys()
    

### Events

*   pii.keys.rotated.v1
    

23.3 Idempotency & Outbox v2 (Observability for Rate Limits)
------------------------------------------------------------

### Stories

*   As a platform engineer, I want idempotent writes and deduped outbox so that events are exactly-once.
    

### Flow

*   Infra: IdempotencyCheck() middleware; PublishOutbox() with dedupe.
    

### Metrics

*   idempotency\_hits\_total, outbox\_publish\_failures\_total, cmd\_latency\_ms{command=\*}
    
*   **(Added)** Global rate-limit metrics (requests per endpoint).

**(New: Specific Metrics)**

*   requests\_per\_endpoint\_total{endpoint=\*,method=\*}
    
*   rate\_limit\_exceeded\_total{user=\*}

23.4 Global Audit & Logging
---------------------------

### Stories

*   As an auditor, I want cross-domain searchable audit logs so that every change is traceable.
    

### Flow

*   AuditChangeCommand →GetAuditLogsQuery | SearchAuditLogsQuery →AuditChange()
    

### Events

*   audit.change.recorded.v1
    

23.5 Event Retry & DLQ
----------------------

### Stories

*   As a reliability engineer, I want retries with exponential backoff and DLQ so that transient failures heal and stuck events are visible.
    

### Flow

*   RetryFailedEventCommand | MoveToDLQCommand →GetFailedEventsQuery →RetryFailedEvent() | MoveToDLQ()
    

### Events

*   event.retry.attempted.v1, event.dlq.moved.v1
    

23.6 Sharded Analytics Aggregation
----------------------------------

### Stories

*   As a data analyst, I want cross-shard aggregation so that global stats are accurate.
    

### Flow

*   AggregateCrossShardStatsCommand →GetGlobalStatsQuery →AggregateCrossShardStats()
    

### Events

*   analytics.aggregated.v1
    

23.7 API Usage Analytics **(NEW)**
----------------------------------

### Stories

*   As a developer, I want API usage stats (per key/user, rate-limit status) so that I can monitor quotas.
    
*   As a platform engineer, I want anomaly detection on API usage so that abuse is prevented.
    
*   **(Added)** As a developer, I want API rate-limit notifications so that I can manage usage.
    
*   **(New)** As a developer, I want API versioning so that my integrations remain stable across updates.
    
### Flow

*   (Updated) GetAPIUsageQuery | GetAPIQuotaQuery | NotifyAPIRateLimitCommand | GetAPIVersionedEndpointQuery → GetAPIUsage() | GetAPIQuota() | NotifyAPIRateLimit() | GetAPIVersionedEndpoint()
    

### Projections

*   api\_usage\_read, api\_keys\_read — **No events** (read-only for usage/quotas)

*   **(New)** api\_version\_read

### Events

*   **api.rate\_limit.notified.v1**


23.8 AI Governance **(NEW)**
----------------------------

### Stories

*   As a privacy officer, I want explicit opt-in/out controls for AI features so that consent is respected.
    
*   As an ethics reviewer, I want bias/fairness audit logs for AI-assisted rankings so that outcomes are explainable.
    

### Flow

*   ToggleAIOptInCommand (feature\_scope) | RecordAIFairnessAuditCommand →GetAIConsentStateQuery | GetAIFairnessReportsQuery →ToggleAIOptIn() | RecordAIFairnessAudit()
    

### Events

*   ai.feature.optin.updated.v1, ai.fairness.audit.recorded.v1
    

23.9 Multi-Tenancy **(Added)**
------------------------------

### Stories

*   As an enterprise admin, I want tenant-isolated data so that my org’s data is private.
    

### Flow

*   ConfigureTenantIsolationCommand →GetTenantConfigQuery →ConfigureTenantIsolation()
    

### Events

*   tenant.isolation.configured.v1
    
23.10 Reliability Engineering (New Subsection)
----------------------------------------------

**Stories**

*   **(New)** As a platform engineer, I want circuit breakers for cross-service calls so that failures are isolated.
    

**Flow**

*   **(New)** ConfigureCircuitBreakerCommand → GetCircuitBreakerStatusQuery → ConfigureCircuitBreaker()
    

**Events**

*   **(New)** reliability.circuit\_breaker.configured.v1
    
23.11 Feature Experimentation (New Subsection)
----------------------------------------------

**Stories**

*   **(New)** As a product manager, I want A/B test success metrics (e.g., statistical significance) so that I can validate feature impact.
    

**Flow**

*   **(New)** GetExperimentSuccessMetricsQuery → GetExperimentSuccessMetrics()
    

**Projections**

*   **(New)** experiment\_success\_metrics\_read
    

23.12 Observability Engineering (New Subsection)
------------------------------------------------

**Stories**

*   **(New)** As a platform engineer, I want error budget tracking for SLOs so that reliability is measurable.
    

**Flow**

*   **(New)** TrackErrorBudgetCommand → GetErrorBudgetStatusQuery → TrackErrorBudget()
    

**Events**

*   **(New)** reliability.error\_budget.updated.v1
    

**Projections**

*   **(New)** error\_budget\_read
    

23.17 User Experience Feedback (New Subsection)
-----------------------------------------------

**Stories**

*   **(New)** As a product manager, I want in-platform user feedback surveys (e.g., NPS) so that I can measure satisfaction.
    

**Flow**

*   **(New)** SendUserSurveyCommand → GetSurveyResponsesQuery → SendUserSurvey()
    

**Events**

*   **(New)** user.survey.sent.v1, user.survey.responded.v1
    

**Projections**

*   **(New)** survey\_responses\_read
    

24) Platform Enhancements (AI/Search Acceleration)
==================================================

24.1 Assistive Features
-----------------------

Stories
-------

*   As a user, I want AI-assisted search and matching so that I find jobs/talent faster.
    
*   As a client, I want intent-based job posts prefilled by AI so that posting is faster.
    
*   As a freelancer, I want AI proposal suggestions so that I can iterate quickly.
    
*   **(Added)** As a user, I want AI-assisted dispute resolution suggestions so that conflicts resolve faster.
    
*   **(New)** As a user, I want AI-assisted contract negotiation suggestions so that agreements are reached faster.

Flow (read/assist)
------------------

*   GetAISearchRecommendationsQuery | GetAIMatchCandidatesQuery | GetAIProposalSuggestionsQuery | GetAIDisputeResolutionSuggestionsQuery →GetAISearchRecommendations() | GetAIMatchCandidates() | GetAIProposalSuggestions() | GetAIDisputeResolutionSuggestions()



Projections
-----------

*   ai\_search\_read, ai\_match\_read

*   **(New)** ai\_contract\_read

Events
------

*   — (assistive, read-side); governed by **23.8** consent.
    
*   **(Added)** ai.dispute.suggestion.generated.v1 (subject to 23.8 consent)

*   **(New)** ai.contract.suggestion.generated.v1

24.2 Scalable AI Model Management (New Subsection)
--------------------------------------------------

**Stories**

*   **(New)** As a platform engineer, I want AI model versioning and rollback so that AI updates are stable.
    

**Flow**

*   **(New)** DeployAIModelVersionCommand | RollbackAIModelCommand → GetAIModelVersionQuery → DeployAIModelVersion() | RollbackAIModel()
    

**Events**

*   **(New)** ai.model.version.deployed.v1, ai.model.version.rolledback.v1
    

**Projections**

*   **(New)** ai\_model\_versions\_read
    

24.3 AI-Driven Onboarding Guidance (New Subsection)
---------------------------------------------------

**Stories**

*   **(New)** As a new user, I want an AI-powered onboarding assistant so that setup is personalized and faster.
    

**Flow**

*   **(New)** GetAIOnboardingRecommendationsQuery → GetAIOnboardingRecommendations() (subject to 23.8 AI consent)
    

**Events**

*   **(New)** ai.onboarding.recommendation.generated.v1
    

**Projections**

*   **(New)** ai\_onboarding\_read

24.4 Cross-Domain Search Optimization **(Added)**
--------------------------------------------

### Stories

*   As a system, I want optimized search indexing so that queries are sub-100ms.
    

### Flow

*   OptimizeSearchIndexCommand →GetSearchIndexStatusQuery →OptimizeSearchIndex()
    

### Events

*   search.index.optimized.v1
    

Acceptance Criteria Patterns (apply across endpoints)
=====================================================

*   **Latency:** Write P95 ≤ 300ms (unless noted), Read P95 ≤ 250ms; batch ops documented (e.g., 10k/60s).
    
*   **Idempotency:** Safe retries with Idempotency-Key; duplicate requests return 200 + no duplicate events.
    
*   **Audit:** AuditChange() called for write commands with actor attribution.
    
*   **Security:** Field-level PII encryption; secrets never logged; RBAC enforced; MFA required on high-risk flows.
    
*   **Observability:** cmd\_latency\_ms, projector\_lag\_ms, event\_delivery\_success\_rate, error\_budget\_burn\_rate.
    
*   **Search:** Min query length = 2; pagination required; timeouts with partial results allowed.
    
*   **Policies:** Endorsement limit (10/client/year), username cooldown (90d), alias cap (5), delegation TTL ≤ 30d.
    

Folder & Code Shape (aligned with Latest-essential-microserices-folder-structure.md)
====================================================================================

```
apps/be/users-be/
  internal/
    domain/<context>/
      commands.go
      service.go            # tx + outbox + retries/DLQ + compensations
      queries.go
      projector.go          # read-model updaters
      policy.go             # limits/cooldowns/endorsement caps
    interfaces/http/routes/<context>.go
    interfaces/subscribers/...
    application/
      outbox/               # topic mappers, dedupe keys
      orchestrators/        # gdpr, kyc_tier, taxonomy_sync, shard_migration
      policies/             # rate-limits, reserved names, abuse limits
    infrastructure/
      encryption/           # PII crypto + rotation
      webhooks/             # HMAC, filters, retries/DLQ
      tokens/               # PAT issuance/validation
      ai_governance/        # consent, audits

```