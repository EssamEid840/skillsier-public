# users-be User Stories - Refactored
## Skillsier Platform - User Management Service

> **Last Updated:** 2025-10-19  
> **Service:** users-be  
> **Architecture:** Clean Architecture + DDD + Event-Driven + CQRS

---

## Table of Contents
1. [Core User Management](#1---core-user-management)
2. [Profile Management](#2---profile-management)
3. [Skills & Experience](#3---skills--experience)
4. [Freelancer & Client Profiles](#4---freelancer--client-profiles)
5. [Verification & Trust](#5---verification--trust)
6. [Account Health & Reputation](#6---account-health--reputation)
7. [Organizations & Teams](#7---organizations--teams)
8. [Security & Authentication](#8---security--authentication)
9. [Privacy & Compliance](#9---privacy--compliance)
10. [Saved Items & Collections](#10---saved-items--collections)
11. [Blocking & Moderation](#11---blocking--moderation)
12. [Suspension & Bans](#12---suspension--bans)
13. [Warnings & Appeals](#13---warnings--appeals)
14. [Risk & Fraud Prevention](#14---risk--fraud-prevention)
15. [Settings & Preferences](#15---settings--preferences)

---

## EVENT CONVENTIONS

### Event Format
- **Pattern**: `aggregate.resource.action.past_tense.v1`
- **Example**: `user.email.verified.v1`, `profile.updated.v1`

### Event Envelope
All events include:
- event_id, event_ts, aggregate_id
- partition_key=user_id
- correlation_id, causation_id
- actor{id, role}
- user_context{ip, user_agent}
- data_zone (EU|US)
- schema_ref
- compliance_context{pii_flags}

### Event Guidelines
- **Batch Operations**: Per-entity events + one `*.summary.v1`
- **PII Protection**: Emit hashes/storage_ids only (no raw PII)
- **Idempotency**: Header `Idempotency-Key` or envelope-based
- **Transactions**: DB tx + outbox with (aggregate_id, event_type, idempotency_key) dedupe
- **Retries/DLQ**: For external calls (Keycloak, storage-be, KYC providers)

---

## 1 - CORE USER MANAGEMENT

### 1.1 user/

#### User Stories
- As a **prospective user**, I want to **register** (freelancer/client/org) so that I can access the platform.
- As a **prospective user**, I want **default profile/settings bootstrapped** so that onboarding is seamless.
- As a **system**, I want **case-insensitive uniqueness** of KeycloakID/username/email so that duplicates are prevented.
- As a **prospective user**, I want to **apply a referral code at registration** so that I get startup benefits.
- As a **user**, I want to **create a guest account** so that I can browse before committing to full registration.
- As a **system**, I want **account type validation** (freelancer/client/org) to prevent role conflicts.
- As a **user**, I want **password setup and recovery** integrated with Keycloak so that login is smooth.

#### Flow
1. **RegisterUserCommand**(keycloak_id, email, username, account_type, referral_code?) → ValidateUniqueness() | ValidateReferral() | BootstrapProfile() | CreateUser() → **Outbox:** user.registered.v1
2. **CreateGuestAccountCommand**(email) → CreateTempUser() → **Outbox:** user.guest.created.v1
3. **UpgradeGuestToFullCommand**(user_id, password) → UpgradeAccount() → **Outbox:** user.upgraded.v1
4. **UpdateUserCommand**(user_id, updates) → Validate() | Update() → **Outbox:** user.updated.v1
5. **DeactivateAccountCommand**(user_id, reason) → Deactivate() → **Outbox:** user.deactivated.v1
6. **ReactivateAccountCommand**(user_id) → Reactivate() → **Outbox:** user.reactivated.v1
7. **DeleteAccountCommand**(user_id) → SoftDelete() | QueueDataDeletion() → **Outbox:** user.deleted.v1
8. **GetUserQuery**(user_id) → Fetch() → UserDTO
9. **SearchUsersQuery**(filters, pagination) → Search() → UserListDTO
10. **GetUserStatisticsQuery**() → Aggregate() → UserStatsDTO (total, active, by_type, growth)

#### Projections
- user_read
- user_search_read
- user_stats_read
- user_retention_cohorts_read

#### Events Published
- user.registered.v1
- user.guest.created.v1
- user.upgraded.v1
- user.updated.v1
- user.deactivated.v1
- user.reactivated.v1
- user.deleted.v1

#### Events Consumed
- keycloak.user.created
- keycloak.user.updated
- keycloak.user.deleted

#### RBAC/SLO
- **RBAC:** OWNER (update own), ADMIN (manage all)
- **SLO:** P95 < 200ms, P99 < 500ms

---

### 1.2 username/

#### User Stories
- As a **user**, I want to **change my username** so that I can update my brand.
- As a **system**, I want **username cooldown** (90 days) so that abuse is prevented.
- As a **system**, I want **reserved usernames** (admin, support, etc.) so that impersonation is prevented.
- As a **user**, I want **username suggestions** when my preferred name is taken.

#### Flow
1. **ChangeUsernameCommand**(user_id, new_username) → ValidateAvailable() | CheckCooldown() | CheckReserved() | Update() → **Outbox:** user.username.changed.v1
2. **CheckUsernameAvailabilityQuery**(username) → Check() → AvailabilityDTO
3. **SuggestUsernamesQuery**(base_username) → GenerateSuggestions() → UsernameListDTO
4. **GetUsernameHistoryQuery**(user_id) → FetchHistory() → UsernameHistoryDTO

#### Projections
- username_history_read
- username_availability_read

#### Events Published
- user.username.changed.v1

#### Username Rules
- 3-20 characters
- Alphanumeric + underscores/hyphens
- Case-insensitive uniqueness
- 90-day cooldown between changes
- Reserved list: admin, support, mod, moderator, staff, skillsier, etc.

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 150ms

---

### 1.3 user_alias/

#### User Stories
- As a **user**, I want **multiple aliases** so that I can maintain different professional identities.
- As a **system**, I want **alias limits** (max 5) so that abuse is controlled.
- As a **user**, I want to **switch between aliases** so that I can segment my work.

#### Flow
1. **CreateAliasCommand**(user_id, alias_name, purpose) → ValidateLimit() | ValidateUnique() | Create() → **Outbox:** user.alias.created.v1
2. **DeleteAliasCommand**(alias_id) → Delete() → **Outbox:** user.alias.deleted.v1
3. **SetPrimaryAliasCommand**(user_id, alias_id) → SetPrimary() → **Outbox:** user.alias.primary.set.v1
4. **ListAliasesQuery**(user_id) → Fetch() → AliasListDTO

#### Projections
- alias_read

#### Events Published
- user.alias.created.v1
- user.alias.deleted.v1
- user.alias.primary.set.v1

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 130ms

---

### 1.4 session/

#### User Stories
- As a **user**, I want to **view active sessions** so that I can monitor access.
- As a **user**, I want to **revoke sessions** so that I can secure my account.
- As a **system**, I want **auto-expiry on inactivity** so that security is maintained.
- As a **system**, I want **MFA challenge on anomalous logins** so that security is layered.
- As a **security admin**, I want to **force global sign-out** so that compromised accounts are secured.

#### Flow
1. **RecordLoginCommand**(user_id, ip, user_agent, device_fingerprint) → ValidateAnomaly() | TriggerMFAIfNeeded() | RecordSession() → **Outbox:** user.login.recorded.v1
2. **RevokeSessionCommand**(session_id) → Revoke() → **Outbox:** user.session.revoked.v1
3. **RevokeAllSessionsCommand**(user_id) → RevokeAll() → **Outbox:** user.sessions.all.revoked.v1
4. **ExpireInactiveSessionsJob**() → FindInactive() | Expire() → **Outbox:** user.session.expired.v1
5. **GetActiveSessionsQuery**(user_id) → Fetch() → SessionListDTO
6. **DetectAnomalousLoginCommand**(user_id, context) → AnalyzeSignals() | ChallengeMFA() → **Outbox:** user.login.anomaly.detected.v1

#### Projections
- session_read
- login_history_read
- anomalous_login_read

#### Events Published
- user.login.recorded.v1
- user.session.revoked.v1
- user.sessions.all.revoked.v1
- user.session.expired.v1
- user.login.anomaly.detected.v1

#### RBAC/SLO
- **RBAC:** OWNER (own sessions), ADMIN (force signout)
- **SLO:** P95 < 140ms

---

## 2 - PROFILE MANAGEMENT

### 2.1 profile/

#### User Stories
- As a **user**, I want to **update my profile** (bio, location, timezone) so that clients/freelancers know about me.
- As a **user**, I want **profile completion tracking** so that I know what's missing.
- As a **user**, I want **profile picture upload** so that my profile is personalized.
- As a **system**, I want **profile completeness scoring** so that quality is incentivized.

#### Flow
1. **UpdateProfileCommand**(user_id, updates) → Validate() | Update() | RecalculateCompleteness() → **Outbox:** profile.updated.v1
2. **UploadProfilePictureCommand**(user_id, file) → ValidateImage() | UploadToStorage() | UpdateURL() → **Outbox:** profile.picture.updated.v1
3. **GetProfileQuery**(user_id) → Fetch() → ProfileDTO
4. **GetProfileCompletenessQuery**(user_id) → CalculateScore() → CompletenessDTO (score, missing_fields)

#### Projections
- profile_read
- profile_completeness_read

#### Events Published
- profile.updated.v1
- profile.picture.updated.v1
- profile.completeness.changed.v1

#### Events Consumed
- storage.file.uploaded

#### Profile Fields
- Bio (max 2000 chars)
- Location (city, country, timezone)
- ProfilePictureURL
- PreferredLanguage
- PreferredCurrency
- CompletionPercentage (calculated)

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 180ms

---

### 2.2 profile_visibility/

#### User Stories
- As a **user**, I want **granular visibility controls** (public/private/connections-only) so that I control my presence.
- As a **user**, I want **stealth mode** so that I can browse without appearing online.
- As a **user**, I want **search visibility toggle** so that I can control discoverability.

#### Flow
1. **UpdateVisibilityCommand**(user_id, visibility_settings) → Update() → **Outbox:** profile.visibility.changed.v1
2. **EnableStealthModeCommand**(user_id) → EnableStealth() → **Outbox:** profile.stealth.enabled.v1
3. **DisableStealthModeCommand**(user_id) → DisableStealth() → **Outbox:** profile.stealth.disabled.v1
4. **SetSearchVisibilityCommand**(user_id, visible) → Update() → **Outbox:** profile.search_visibility.changed.v1
5. **GetVisibilitySettingsQuery**(user_id) → Fetch() → VisibilityDTO

#### Projections
- visibility_settings_read

#### Events Published
- profile.visibility.changed.v1
- profile.stealth.enabled.v1
- profile.stealth.disabled.v1
- profile.search_visibility.changed.v1

#### Visibility Levels
- Public (anyone can view)
- ConnectionsOnly (only connections can view)
- Private (only user can view)
- Stealth (appear offline)

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 120ms

---

### 2.3 availability/

#### User Stories
- As a **freelancer**, I want to **set my availability status** (available/busy/away) so that clients know when I'm free.
- As a **freelancer**, I want **recurring availability schedules** so that I don't have to update daily.
- As a **freelancer**, I want **vacation mode** so that I can pause work without losing visibility.
- As a **freelancer**, I want **calendar sync** (Google/Outlook) so that availability is accurate.

#### Flow
1. **UpdateAvailabilityStatusCommand**(user_id, status) → Update() → **Outbox:** availability.status.updated.v1
2. **SetRecurringScheduleCommand**(user_id, schedule) → Validate() | Persist() → **Outbox:** availability.schedule.set.v1
3. **EnableVacationModeCommand**(user_id, start_date, end_date, auto_message) → Enable() → **Outbox:** availability.vacation.enabled.v1
4. **DisableVacationModeCommand**(user_id) → Disable() → **Outbox:** availability.vacation.disabled.v1
5. **SyncCalendarCommand**(user_id, calendar_provider) → FetchCalendar() | UpdateAvailability() → **Outbox:** availability.calendar.synced.v1
6. **GetAvailabilityQuery**(user_id) → Fetch() → AvailabilityDTO

#### Projections
- availability_read
- availability_schedule_read

#### Events Published
- availability.status.updated.v1
- availability.schedule.set.v1
- availability.vacation.enabled.v1
- availability.vacation.disabled.v1
- availability.calendar.synced.v1

#### Availability Statuses
- Available, Busy, Away, DoNotDisturb, Vacation

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 160ms

---

### 2.4 workload_capacity/

#### User Stories
- As a **freelancer**, I want **capacity tracking** so that I don't overcommit.
- As a **system**, I want **overload prevention** so that quality is maintained.
- As a **freelancer**, I want to **see available hours** so that I can plan new work.

#### Flow
1. **UpdateCapacityCommand**(user_id, max_capacity, available_hours) → Update() → **Outbox:** capacity.updated.v1
2. **TrackCommitmentCommand**(user_id, contract_id, hours) → UpdateLoad() → **Outbox:** capacity.commitment.tracked.v1
3. **CheckCapacityAvailabilityQuery**(user_id, required_hours) → Calculate() → CapacityAvailabilityDTO
4. **GetWorkloadQuery**(user_id) → Fetch() → WorkloadDTO (current_load, max_capacity, available_hours)

#### Projections
- workload_read
- capacity_analytics_read

#### Events Published
- capacity.updated.v1
- capacity.commitment.tracked.v1
- capacity.full.v1
- capacity.available.v1

#### Events Consumed
- contract.started
- contract.completed

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 140ms

---

## 3 - SKILLS & EXPERIENCE

### 3.1 skill/

#### User Stories
- As a **user**, I want to **add skills with proficiency levels** so that my expertise is clear.
- As a **user**, I want **skill endorsements** so that credibility is established.
- As a **system**, I want **skill normalization** so that search is consistent.
- As a **system**, I want **skill trending analytics** so that market insights are provided.

#### Flow
1. **AddSkillCommand**(user_id, skill_name, proficiency, years_experience) → NormalizeSkill() | Add() → **Outbox:** skill.added.v1
2. **UpdateSkillCommand**(skill_id, updates) → Update() → **Outbox:** skill.updated.v1
3. **RemoveSkillCommand**(skill_id) → Remove() → **Outbox:** skill.removed.v1
4. **EndorseSkillCommand**(skill_id, endorser_id) → ValidateEndorser() | AddEndorsement() → **Outbox:** skill.endorsed.v1
5. **GetSkillsQuery**(user_id) → Fetch() → SkillListDTO
6. **GetSkillTrendsQuery**() → Aggregate() → SkillTrendsDTO (trending, demand, growth)

#### Projections
- skill_read
- skill_endorsement_read
- skill_trends_read

#### Events Published
- skill.added.v1
- skill.updated.v1
- skill.removed.v1
- skill.endorsed.v1

#### Proficiency Levels
- Beginner, Intermediate, Advanced, Expert

#### RBAC/SLO
- **RBAC:** OWNER (manage own), ANY_USER (endorse)
- **SLO:** P95 < 150ms

---

### 3.2 experience/

#### User Stories
- As a **freelancer**, I want to **add work experience** so that my history is visible.
- As a **freelancer**, I want to **mark current position** so that my status is clear.
- As a **system**, I want **date validation** so that data quality is maintained.

#### Flow
1. **AddExperienceCommand**(user_id, company, title, description, start_date, end_date?, is_current) → ValidateDates() | Add() → **Outbox:** experience.added.v1
2. **UpdateExperienceCommand**(experience_id, updates) → Validate() | Update() → **Outbox:** experience.updated.v1
3. **RemoveExperienceCommand**(experience_id) → Remove() → **Outbox:** experience.removed.v1
4. **GetExperienceQuery**(user_id) → Fetch() → ExperienceListDTO

#### Projections
- experience_read

#### Events Published
- experience.added.v1
- experience.updated.v1
- experience.removed.v1

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 140ms

---

### 3.3 education/

#### User Stories
- As a **user**, I want to **add education** so that my qualifications are visible.
- As a **system**, I want **degree verification** so that trust is maintained.

#### Flow
1. **AddEducationCommand**(user_id, school, degree, field, graduation_year, description) → ValidateYear() | Add() → **Outbox:** education.added.v1
2. **UpdateEducationCommand**(education_id, updates) → Update() → **Outbox:** education.updated.v1
3. **RemoveEducationCommand**(education_id) → Remove() → **Outbox:** education.removed.v1
4. **GetEducationQuery**(user_id) → Fetch() → EducationListDTO

#### Projections
- education_read

#### Events Published
- education.added.v1
- education.updated.v1
- education.removed.v1

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 130ms

---

### 3.4 certification/

#### User Stories
- As a **freelancer**, I want to **add certifications** so that my credentials are visible.
- As a **system**, I want **expiry tracking** so that outdated certs are flagged.
- As a **system**, I want **certification verification** via third-party APIs.

#### Flow
1. **AddCertificationCommand**(user_id, name, issuing_org, issue_date, expiry_date?, credential_id?, url?) → Add() | ScheduleVerification() → **Outbox:** certification.added.v1
2. **UpdateCertificationCommand**(certification_id, updates) → Update() → **Outbox:** certification.updated.v1
3. **RemoveCertificationCommand**(certification_id) → Remove() → **Outbox:** certification.removed.v1
4. **VerifyCertificationCommand**(certification_id) → VerifyWithIssuer() | UpdateStatus() → **Outbox:** certification.verified.v1
5. **ExpireCertificationsJob**() → FindExpired() | MarkExpired() → **Outbox:** certification.expired.v1
6. **GetCertificationsQuery**(user_id) → Fetch() → CertificationListDTO

#### Projections
- certification_read
- certification_verification_read

#### Events Published
- certification.added.v1
- certification.updated.v1
- certification.removed.v1
- certification.verified.v1
- certification.expired.v1

#### Verification Statuses
- Pending, Verified, Rejected, Expired

#### RBAC/SLO
- **RBAC:** OWNER (manage), SYSTEM (verify)
- **SLO:** P95 < 180ms

---

### 3.5 portfolio/

#### User Stories
- As a **freelancer**, I want to **create portfolio items** so that I can showcase my work.
- As a **freelancer**, I want **multiple images per item** so that I can show comprehensive examples.
- As a **system**, I want **portfolio visibility controls** so that users control what's shown.

#### Flow
1. **CreatePortfolioItemCommand**(user_id, title, description, url?, project_date?, tags[]) → Create() → **Outbox:** portfolio.item.created.v1
2. **UpdatePortfolioItemCommand**(item_id, updates) → Update() → **Outbox:** portfolio.item.updated.v1
3. **DeletePortfolioItemCommand**(item_id) → Delete() → **Outbox:** portfolio.item.deleted.v1
4. **AddPortfolioImageCommand**(item_id, file) → UploadToStorage() | AddImage() → **Outbox:** portfolio.image.added.v1
5. **RemovePortfolioImageCommand**(image_id) → Remove() → **Outbox:** portfolio.image.removed.v1
6. **SetPortfolioVisibilityCommand**(item_id, visibility) → Update() → **Outbox:** portfolio.visibility.changed.v1
7. **GetPortfolioQuery**(user_id) → Fetch() → PortfolioListDTO

#### Projections
- portfolio_read

#### Events Published
- portfolio.item.created.v1
- portfolio.item.updated.v1
- portfolio.item.deleted.v1
- portfolio.image.added.v1
- portfolio.image.removed.v1
- portfolio.visibility.changed.v1

#### Events Consumed
- storage.file.uploaded

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 200ms

---

## 4 - FREELANCER & CLIENT PROFILES

### 4.1 freelancer/

#### User Stories
- As a **freelancer**, I want a **specialized profile** (hourly_rate, job_success_rate, total_earnings) so that I attract clients.
- As a **freelancer**, I want **rate cards** so that pricing is consistent.
- As a **system**, I want **job success tracking** so that performance is visible.

#### Flow
1. **CreateFreelancerProfileCommand**(user_id, hourly_rate, title, overview, category) → Create() → **Outbox:** freelancer.profile.created.v1
2. **UpdateFreelancerProfileCommand**(user_id, updates) → Update() → **Outbox:** freelancer.profile.updated.v1
3. **UpdateJobSuccessRateCommand**(user_id, success_rate) → Update() → **Outbox:** freelancer.success_rate.updated.v1
4. **GetFreelancerProfileQuery**(user_id) → Fetch() → FreelancerProfileDTO
5. **GetFreelancerStatsQuery**(user_id) → Aggregate() → FreelancerStatsDTO (earnings, jobs_completed, success_rate)

#### Projections
- freelancer_profile_read
- freelancer_stats_read

#### Events Published
- freelancer.profile.created.v1
- freelancer.profile.updated.v1
- freelancer.success_rate.updated.v1

#### Events Consumed
- contract.completed
- review.submitted

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 170ms

---

### 4.2 client/

#### User Stories
- As a **client**, I want a **specialized profile** (company_size, industry, total_spent, hire_rate) so that I attract quality freelancers.
- As a **client**, I want **payment verification** so that trust is established.
- As a **system**, I want **spending forecasts** so that budgeting is accurate.

#### Flow
1. **CreateClientProfileCommand**(user_id, company_name, industry, company_size, website?) → Create() → **Outbox:** client.profile.created.v1
2. **UpdateClientProfileCommand**(user_id, updates) → Update() → **Outbox:** client.profile.updated.v1
3. **VerifyPaymentMethodCommand**(user_id) → VerifyWithFinancialBE() | UpdateStatus() → **Outbox:** client.payment.verified.v1
4. **GetClientProfileQuery**(user_id) → Fetch() → ClientProfileDTO
5. **GetSpendingForecastQuery**(user_id) → CalculateForecast() → SpendingForecastDTO

#### Projections
- client_profile_read
- client_stats_read
- client_budget_forecast_read

#### Events Published
- client.profile.created.v1
- client.profile.updated.v1
- client.payment.verified.v1

#### Events Consumed
- financial.payment.processed
- contract.created

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 180ms

---

## 5 - VERIFICATION & TRUST

### 5.1 verification/ (KYC/KYB)

#### User Stories
- As a **user**, I want to **submit KYC documents** so that I can unlock trusted features.
- As a **compliance reviewer**, I want to **approve/reject verifications** so that trust is enforced.
- As a **system**, I want **encrypted storage** via storage-be so that privacy is protected.
- As a **system**, I want **sanctions/OFAC/PEP screening** so that risky users are flagged.
- As a **system**, I want **VAT/GST validation** for international users.

#### Flow
1. **SubmitKYCCommand**(user_id, document_type, file) → UploadEncrypted() | CreateVerificationRequest() → **Outbox:** kyc.submitted.v1
2. **ReviewKYCCommand**(verification_id, decision, reviewer_id, notes) → UpdateStatus() | NotifyUser() → **Outbox:** kyc.reviewed.v1
3. **RunSanctionsCheckCommand**(user_id) → CheckOFAC() | CheckPEP() | UpdateRiskProfile() → **Outbox:** sanctions.check.completed.v1
4. **ValidateVATCommand**(user_id, vat_number, country) → ValidateWithTaxAuthority() → **Outbox:** vat.validated.v1
5. **GetVerificationStatusQuery**(user_id) → Fetch() → VerificationStatusDTO

#### Projections
- verification_read
- sanctions_screening_read

#### Events Published
- kyc.submitted.v1
- kyc.reviewed.v1
- kyc.approved.v1
- kyc.rejected.v1
- sanctions.check.completed.v1
- vat.validated.v1

#### Verification Statuses
- Pending, UnderReview, Approved, Rejected, Expired

#### Document Types
- GovernmentID, Passport, DriversLicense, Utility Bill, Business Registration

#### RBAC/SLO
- **RBAC:** OWNER (submit), COMPLIANCE_REVIEWER (review)
- **SLO:** P95 < 300ms (async processing)

---

### 5.2 badge/

#### User Stories
- As a **user**, I want **achievement badges** so that my accomplishments are visible.
- As a **system**, I want **auto-award badges** based on milestones.
- As a **user**, I want to **display badges on profile** so that trust is increased.

#### Flow
1. **AwardBadgeCommand**(user_id, badge_type, awarded_by) → Award() → **Outbox:** badge.awarded.v1
2. **RevokeBadgeCommand**(badge_id, reason) → Revoke() → **Outbox:** badge.revoked.v1
3. **CheckBadgeEligibilityJob**() → FindEligible() | AutoAward() → **Outbox:** badge.auto.awarded.v1
4. **GetBadgesQuery**(user_id) → Fetch() → BadgeListDTO

#### Projections
- badge_read
- badge_achievements_read

#### Events Published
- badge.awarded.v1
- badge.revoked.v1
- badge.auto.awarded.v1

#### Badge Types
- TopRated, RisingTalent, Verified, Responsive, HighEarner, ExpertVetted

#### RBAC/SLO
- **RBAC:** SYSTEM (award), ADMIN (revoke)
- **SLO:** P95 < 150ms

---

### 5.3 endorsement/

#### User Stories
- As a **user**, I want to **endorse colleagues** so that I can vouch for their work.
- As a **system**, I want **endorsement limits** (10/client/year) so that abuse is prevented.
- As a **user**, I want to **see who endorsed me** so that credibility is transparent.

#### Flow
1. **CreateEndorsementCommand**(endorser_id, endorsee_id, skill_id, comment) → ValidateLimit() | ValidateRelationship() | Create() → **Outbox:** endorsement.created.v1
2. **RevokeEndorsementCommand**(endorsement_id) → Revoke() → **Outbox:** endorsement.revoked.v1
3. **GetEndorsementsQuery**(user_id) → Fetch() → EndorsementListDTO
4. **CheckEndorsementEligibilityQuery**(endorser_id, endorsee_id) → CheckLimits() → EligibilityDTO

#### Projections
- endorsement_read

#### Events Published
- endorsement.created.v1
- endorsement.revoked.v1

#### RBAC/SLO
- **RBAC:** ANY_USER (endorse), OWNER (revoke own)
- **SLO:** P95 < 160ms

---

## 6 - ACCOUNT HEALTH & REPUTATION

### 6.1 account_health/

#### User Stories
- As a **user**, I want to **see my account health score** so that I know where I stand.
- As a **system**, I want **health calculation from multiple factors** (profile completeness, activity, responsiveness, quality).
- As a **user**, I want **actionable recommendations** to improve my health score.

#### Flow
1. **CalculateHealthScoreCommand**(user_id) → FetchMetrics() | ComputeFactors() | CalculateScore() | GenerateRecommendations() → **Outbox:** health.score.updated.v1
2. **DetectHealthIssuesCommand**(user_id) → AnalyzePatterns() | FlagIssues() → **Outbox:** health.issue.detected.v1
3. **GetHealthScoreQuery**(user_id) → Fetch() → HealthScoreDTO (score, factors, issues, recommendations)
4. **RecalculateAllHealthScoresJob**() → BatchRecalculate() → **Outbox:** health.scores.recalculated.v1

#### Projections
- account_health_read
- health_trends_read

#### Events Published
- health.score.updated.v1
- health.issue.detected.v1
- health.improved.v1

#### Health Factors
- ProfileCompleteness (30%)
- ActivityLevel (20%)
- ResponseTime (20%)
- QualityWork (20%)
- ClientSatisfaction (10%)

#### Score Ranges
- Excellent: 90-100
- Good: 75-89
- Average: 60-74
- Poor: 40-59
- Critical: 0-39

#### RBAC/SLO
- **RBAC:** OWNER (view own), SYSTEM (calculate)
- **SLO:** P95 < 250ms

---

### 6.2 reputation/

#### User Stories
- As a **user**, I want **reputation tracking** so that my standing is visible.
- As a **system**, I want **reputation decay** so that old behavior doesn't dominate forever.
- As a **user**, I want to **rebuild reputation** after issues.

#### Flow
1. **UpdateReputationCommand**(user_id, event_type, impact) → CalculateImpact() | Update() → **Outbox:** reputation.updated.v1
2. **DecayReputationJob**() → FindStale() | ApplyDecay() → **Outbox:** reputation.decayed.v1
3. **GetReputationQuery**(user_id) → Fetch() → ReputationDTO (score, rank, history)

#### Projections
- reputation_read
- reputation_history_read

#### Events Published
- reputation.updated.v1
- reputation.decayed.v1

#### Events Consumed
- contract.completed
- review.submitted
- dispute.resolved
- warning.issued

#### RBAC/SLO
- **RBAC:** OWNER (view own), SYSTEM (update)
- **SLO:** P95 < 180ms

---

## 7 - ORGANIZATIONS & TEAMS

### 7.1 org/

#### User Stories
- As an **org admin**, I want to **create an organization** so that my team can collaborate.
- As an **org admin**, I want to **manage members and roles** so that permissions are controlled.
- As an **org admin**, I want **seat management** so that capacity is tracked.
- As an **org admin**, I want **shared billing integration** so that costs are centralized.

#### Flow
1. **CreateOrgCommand**(name, industry, size, founder_id) → Create() | AssignOwner() → **Outbox:** org.created.v1
2. **UpdateOrgCommand**(org_id, updates) → Update() → **Outbox:** org.updated.v1
3. **AddMemberCommand**(org_id, user_id, role) → ValidateSeats() | Add() → **Outbox:** org.member.added.v1
4. **RemoveMemberCommand**(org_id, user_id) → Remove() → **Outbox:** org.member.removed.v1
5. **UpdateSeatsCommand**(org_id, new_limit) → UpdateLimit() → **Outbox:** org.seats.updated.v1
6. **LinkBillingCommand**(org_id, billing_profile_id) → Link() → **Outbox:** org.billing.linked.v1
7. **GetOrgQuery**(org_id) → Fetch() → OrgDTO
8. **ListMembersQuery**(org_id) → Fetch() → MemberListDTO

#### Projections
- org_read
- org_members_read
- org_seats_usage_read

#### Events Published
- org.created.v1
- org.updated.v1
- org.member.added.v1
- org.member.removed.v1
- org.seats.updated.v1
- org.billing.linked.v1

#### Member Roles
- Owner, Admin, Member, Viewer

#### RBAC/SLO
- **RBAC:** ORG_OWNER (full control), ORG_ADMIN (manage members)
- **SLO:** P95 < 200ms

---

### 7.2 org_talent_pool/

#### User Stories
- As a **recruiter**, I want to **create talent pools** so that hiring pipelines are organized.
- As a **recruiter**, I want **AI-curated talent recommendations** so that I find vetted freelancers quickly.
- As a **recruiter**, I want to **add users to pools** so that I can segment talent.

#### Flow
1. **CreateTalentPoolCommand**(org_id, name, description, criteria) → Create() → **Outbox:** talent_pool.created.v1
2. **AddToTalentPoolCommand**(pool_id, user_id) → Add() → **Outbox:** talent_pool.member.added.v1
3. **RemoveFromTalentPoolCommand**(pool_id, user_id) → Remove() → **Outbox:** talent_pool.member.removed.v1
4. **GetAIRecommendationsQuery**(pool_id) → AnalyzeCriteria() | ScoreMatches() → RecommendationListDTO
5. **ListTalentPoolsQuery**(org_id) → Fetch() → TalentPoolListDTO

#### Projections
- talent_pool_read
- talent_pool_recommendations_read

#### Events Published
- talent_pool.created.v1
- talent_pool.member.added.v1
- talent_pool.member.removed.v1

#### RBAC/SLO
- **RBAC:** ORG_ADMIN
- **SLO:** P95 < 220ms

---

## 8 - SECURITY & AUTHENTICATION

### 8.1 two_factor/

#### User Stories
- As a **user**, I want to **enable 2FA** so that my account is secure.
- As a **user**, I want **multiple 2FA methods** (TOTP, SMS, Email) so that I have options.
- As a **user**, I want **backup codes** so that I can recover access.

#### Flow
1. **EnableTwoFactorCommand**(user_id, method) → GenerateSecret() | SendSetupInstructions() → **Outbox:** two_factor.enabled.v1
2. **DisableTwoFactorCommand**(user_id, confirmation_code) → Verify() | Disable() → **Outbox:** two_factor.disabled.v1
3. **GenerateBackupCodesCommand**(user_id) → GenerateCodes() | Encrypt() | Store() → **Outbox:** backup_codes.generated.v1
4. **VerifyTwoFactorCodeCommand**(user_id, code) → Verify() → VerificationResultDTO

#### Projections
- two_factor_settings_read

#### Events Published
- two_factor.enabled.v1
- two_factor.disabled.v1
- two_factor.method.changed.v1
- backup_codes.generated.v1

#### 2FA Methods
- TOTP (Authenticator app)
- SMS
- Email
- Hardware key (future)

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 180ms

---

### 8.2 device_management/

#### User Stories
- As a **user**, I want to **see registered devices** so that I can monitor access.
- As a **user**, I want to **revoke device access** so that I can secure my account.
- As a **system**, I want **device fingerprinting** so that anomalies are detected.

#### Flow
1. **RegisterDeviceCommand**(user_id, device_fingerprint, device_name) → Register() → **Outbox:** device.registered.v1
2. **RevokeDeviceCommand**(device_id) → Revoke() | InvalidateSessions() → **Outbox:** device.revoked.v1
3. **GetDevicesQuery**(user_id) → Fetch() → DeviceListDTO
4. **DetectNewDeviceCommand**(user_id, device_fingerprint) → CheckKnown() | AlertIfNew() → **Outbox:** device.new.detected.v1

#### Projections
- device_read

#### Events Published
- device.registered.v1
- device.revoked.v1
- device.new.detected.v1

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 140ms

---

### 8.3 account_recovery/

#### User Stories
- As a **user**, I want **multiple recovery methods** so that I can regain access.
- As a **system**, I want **recovery rate limiting** so that brute force is prevented.
- As a **user**, I want **recovery keys** so that I have a last resort.

#### Flow
1. **InitiateRecoveryCommand**(email_or_username) → FindUser() | SendRecoveryCode() → **Outbox:** recovery.initiated.v1
2. **VerifyRecoveryCodeCommand**(code, user_id) → ValidateCode() | RateLimitCheck() → **Outbox:** recovery.code.verified.v1
3. **CompleteRecoveryCommand**(user_id, new_password) → ResetPassword() | InvalidateOldSessions() → **Outbox:** recovery.completed.v1
4. **GenerateRecoveryKeysCommand**(user_id) → GenerateKeys() | Encrypt() → **Outbox:** recovery_keys.generated.v1

#### Projections
- recovery_attempts_read

#### Events Published
- recovery.initiated.v1
- recovery.code.verified.v1
- recovery.completed.v1
- recovery_keys.generated.v1

#### Recovery Methods
- Email, SMS, SecurityQuestions, RecoveryCodes

#### RBAC/SLO
- **RBAC:** PUBLIC (initiate), OWNER (complete)
- **SLO:** P95 < 200ms

---

## 9 - PRIVACY & COMPLIANCE

### 9.1 privacy/

#### User Stories
- As a **user**, I want to **control data sharing** so that my privacy is protected.
- As a **user**, I want to **export my data** so that I comply with GDPR.
- As a **user**, I want to **request data deletion** so that I exercise my right to be forgotten.

#### Flow
1. **UpdatePrivacySettingsCommand**(user_id, settings) → Update() → **Outbox:** privacy.settings.updated.v1
2. **RequestDataExportCommand**(user_id) → QueueExport() → **Outbox:** data.export.requested.v1
3. **ProcessDataExportJob**(export_id) → GatherData() | Encrypt() | GenerateDownloadLink() → **Outbox:** data.export.ready.v1
4. **RequestDataDeletionCommand**(user_id) → QueueDeletion() | NotifyDependentServices() → **Outbox:** data.deletion.requested.v1
5. **GetPrivacySettingsQuery**(user_id) → Fetch() → PrivacySettingsDTO

#### Projections
- privacy_settings_read
- data_export_requests_read

#### Events Published
- privacy.settings.updated.v1
- data.export.requested.v1
- data.export.ready.v1
- data.deletion.requested.v1
- data.deletion.completed.v1

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 180ms (async for export/deletion)

---

### 9.2 consent/

#### User Stories
- As a **user**, I want **granular consent controls** (marketing, analytics, AI features) so that I control data usage.
- As a **compliance officer**, I want **consent audit trails** so that compliance is provable.
- As a **system**, I want **consent versioning** so that changes are tracked.

#### Flow
1. **UpdateConsentCommand**(user_id, consent_type, granted) → RecordConsent() | UpdateSettings() → **Outbox:** consent.updated.v1
2. **GetConsentHistoryQuery**(user_id) → FetchAuditTrail() → ConsentHistoryDTO
3. **RequestConsentCommand**(user_id, consent_type, reason) → SendRequest() → **Outbox:** consent.requested.v1

#### Projections
- consent_read
- consent_audit_read

#### Events Published
- consent.updated.v1
- consent.requested.v1
- consent.withdrawn.v1

#### Consent Types
- Marketing, Analytics, AIFeatures, DataSharing, Cookies

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 120ms

---

### 9.3 data_residency/

#### User Stories
- As a **user**, I want to **choose data residency** (EU/US) so that I comply with local laws.
- As a **system**, I want **automatic migration** when preference changes.

#### Flow
1. **SetDataResidencyCommand**(user_id, region) → ValidateRegion() | QueueMigration() → **Outbox:** data_residency.updated.v1
2. **MigrateDataJob**(user_id, source_region, target_region) → CopyData() | VerifyIntegrity() | UpdatePointers() → **Outbox:** data_residency.migration.completed.v1
3. **GetDataResidencyQuery**(user_id) → Fetch() → DataResidencyDTO

#### Projections
- data_residency_read

#### Events Published
- data_residency.updated.v1
- data_residency.migration.completed.v1

#### Regions
- EU, US, APAC

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 160ms

---

## 10 - SAVED ITEMS & COLLECTIONS

### 10.1 saved_item/

#### User Stories
- As a **user**, I want to **save jobs or freelancers** so that I can revisit later.
- As a **user**, I want to **add notes** to saved items so that context is remembered.
- As a **system**, I want **duplicate prevention** so that lists are clean.

#### Flow
1. **SaveItemCommand**(user_id, item_type, item_id, notes?) → ValidateDuplicate() | Save() → **Outbox:** item.saved.v1
2. **UnsaveItemCommand**(saved_item_id) → Remove() → **Outbox:** item.unsaved.v1
3. **UpdateSavedItemNoteCommand**(saved_item_id, notes) → Update() → **Outbox:** item.note.updated.v1
4. **ListSavedItemsQuery**(user_id, item_type?) → Fetch() → SavedItemListDTO
5. **SearchSavedItemsQuery**(user_id, query) → Search() → SavedItemListDTO

#### Projections
- saved_items_read

#### Events Published
- item.saved.v1
- item.unsaved.v1
- item.note.updated.v1

#### Item Types
- Job, Freelancer, Agency

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 130ms

---

### 10.2 collection/

#### User Stories
- As a **user**, I want to **organize saved items in collections** so that I can segment my interests.
- As an **org member**, I want to **share collections** with my team.
- As a **user**, I want to **move items between collections** so that organization is flexible.

#### Flow
1. **CreateCollectionCommand**(user_id, name, description, is_shared?) → Create() → **Outbox:** collection.created.v1
2. **MoveItemToCollectionCommand**(saved_item_id, collection_id) → Move() → **Outbox:** item.moved.v1
3. **ShareCollectionCommand**(collection_id, user_ids[], org_id?, permissions) → Share() → **Outbox:** collection.shared.v1
4. **DeleteCollectionCommand**(collection_id) → Delete() → **Outbox:** collection.deleted.v1
5. **ListCollectionsQuery**(user_id) → Fetch() → CollectionListDTO

#### Projections
- collection_read

#### Events Published
- collection.created.v1
- collection.updated.v1
- collection.deleted.v1
- item.moved.v1
- collection.shared.v1

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 150ms

---

### 10.3 saved_item_reminder/

#### User Stories
- As a **user**, I want **reminders for saved items** so that I don't forget to follow up.
- As a **system**, I want **scheduled reminder delivery** so that users are notified on time.

#### Flow
1. **SetReminderCommand**(saved_item_id, remind_at, message?) → Schedule() → **Outbox:** reminder.scheduled.v1
2. **SendRemindersJob**() → FindDue() | SendNotifications() → **Outbox:** reminder.sent.v1
3. **CancelReminderCommand**(reminder_id) → Cancel() → **Outbox:** reminder.cancelled.v1

#### Projections
- reminders_read

#### Events Published
- reminder.scheduled.v1
- reminder.sent.v1
- reminder.cancelled.v1

#### Events Consumed (to send notification)
- reminder.sent.v1 → communications-be

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 140ms

---

## 11 - BLOCKING & MODERATION

### 11.1 block/

#### User Stories
- As a **user**, I want to **block other users** so that I control interactions.
- As a **user**, I want **scoped blocks** (messaging/invites/full) so that restrictions fit my needs.
- As a **user**, I want **time-bound blocks** so that I can cool down temporarily.
- As a **system**, I want **self-block prevention** so that edge cases are handled.

#### Flow
1. **BlockUserCommand**(blocker_id, blocked_id, reason, scope?, duration?) → ValidateNotSelf() | Block() → **Outbox:** user.blocked.v1
2. **UnblockUserCommand**(blocker_id, blocked_id) → Unblock() → **Outbox:** user.unblocked.v1
3. **ExtendBlockCommand**(block_id, new_duration) → Extend() → **Outbox:** block.extended.v1
4. **ExpireBlocksJob**() → FindExpired() | AutoUnblock() → **Outbox:** block.expired.v1
5. **GetBlockedUsersQuery**(user_id) → Fetch() → BlockedUserListDTO
6. **IsBlockedQuery**(user_id1, user_id2) → Check() → BlockStatusDTO

#### Projections
- block_read

#### Events Published
- user.blocked.v1
- user.unblocked.v1
- block.extended.v1
- block.expired.v1

#### Block Scopes
- Full (all interactions blocked)
- Messaging (only messages blocked)
- Invites (only job invites blocked)

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 110ms

---

### 11.2 flag/

#### User Stories
- As a **user**, I want to **flag inappropriate profiles** so that moderation can review.
- As a **moderator**, I want to **review flags** so that cases are handled efficiently.
- As a **system**, I want **flag aggregation** so that repeat offenders are identified.

#### Flow
1. **FlagUserCommand**(flagger_id, flagged_user_id, reason, description) → ValidateNotOwn() | CreateFlag() → **Outbox:** user.flagged.v1
2. **ReviewFlagCommand**(flag_id, reviewer_id, action, notes) → TakeAction() | NotifyReporter() → **Outbox:** flag.reviewed.v1
3. **ResolveFlagCommand**(flag_id, resolution) → Close() → **Outbox:** flag.resolved.v1
4. **GetFlagsQuery**(user_id) → Fetch() → FlagListDTO
5. **GetPendingFlagsQuery**(filters) → FetchPending() → FlagListDTO

#### Projections
- flag_read
- flag_history_read

#### Events Published
- user.flagged.v1
- flag.reviewed.v1
- flag.resolved.v1

#### Flag Reasons
- Spam, Harassment, FakeProfile, Inappropriate, Scam, Other

#### RBAC/SLO
- **RBAC:** ANY_USER (flag), MODERATOR (review)
- **SLO:** P95 < 170ms

---

## 12 - SUSPENSION & BANS

### 12.1 suspension/

#### User Stories
- As a **moderator**, I want to **suspend users** so that policy is enforced.
- As a **moderator**, I want to **schedule suspensions** so that grace periods are honored.
- As a **system**, I want **auto-suspend based on warnings+risk** so that enforcement is consistent.
- As a **user**, I want to **appeal suspensions** so that fairness is preserved.

#### Flow
1. **SuspendUserCommand**(user_id, reason, duration, moderator_id) → Suspend() | NotifyUser() → **Outbox:** user.suspended.v1
2. **ReleaseSuspensionCommand**(user_id, moderator_id) → Release() → **Outbox:** user.suspension.released.v1
3. **ExtendSuspensionCommand**(suspension_id, new_duration, reason) → Extend() → **Outbox:** user.suspension.extended.v1
4. **ScheduleSuspensionCommand**(user_id, starts_at, duration, reason) → Schedule() → **Outbox:** user.suspension.scheduled.v1
5. **AutoSuspendJob**() → FindEligible() | ApplyMatrix() | Suspend() → **Outbox:** user.suspension.auto.placed.v1
6. **AppealSuspensionCommand**(suspension_id, appeal_text, evidence?) → CreateAppeal() → **Outbox:** user.suspension.appealed.v1
7. **GetSuspensionHistoryQuery**(user_id) → FetchHistory() → SuspensionHistoryDTO

#### Projections
- suspension_read
- suspension_history_read
- suspension_appeals_read

#### Events Published
- user.suspended.v1
- user.suspension.released.v1
- user.suspension.extended.v1
- user.suspension.scheduled.v1
- user.suspension.auto.placed.v1
- user.suspension.appealed.v1

#### RBAC/SLO
- **RBAC:** MODERATOR (suspend/release), OWNER (appeal)
- **SLO:** P95 < 180ms

---

### 12.2 ban/

#### User Stories
- As a **trust & safety admin**, I want to **ban users permanently** so that severe violations are handled.
- As an **admin**, I want **ban evasion detection** so that repeat offenders are caught.
- As a **moderator**, I want **shadow-ban** so that spam is dampened during investigations.

#### Flow
1. **BanUserCommand**(user_id, reason, is_permanent, admin_id) → Ban() | NotifyUser() | InvalidateAllSessions() → **Outbox:** user.banned.v1
2. **ReleaseBanCommand**(user_id, admin_id, reason) → Release() → **Outbox:** user.ban.released.v1
3. **EnableShadowBanCommand**(user_id, moderator_id) → EnableShadowBan() → **Outbox:** user.shadowban.enabled.v1
4. **DisableShadowBanCommand**(user_id) → DisableShadowBan() → **Outbox:** user.shadowban.disabled.v1
5. **DetectBanEvasionCommand**(user_id) → AnalyzeIP() | AnalyzeDevice() | CheckPatterns() → **Outbox:** ban.evasion.detected.v1
6. **GetBanHistoryQuery**(user_id) → FetchHistory() → BanHistoryDTO

#### Projections
- ban_read
- ban_history_read
- ban_evasion_read

#### Events Published
- user.banned.v1
- user.ban.released.v1
- user.shadowban.enabled.v1
- user.shadowban.disabled.v1
- ban.evasion.detected.v1

#### RBAC/SLO
- **RBAC:** TRUST_SAFETY_ADMIN
- **SLO:** P95 < 200ms

---

## 13 - WARNINGS & APPEALS

### 13.1 warning/

#### User Stories
- As a **moderator**, I want to **issue warnings** so that users can correct behavior.
- As a **user**, I want to **acknowledge warnings** so that my record is accurate.
- As a **system**, I want **warning escalation** on repeated violations.
- As a **system**, I want **warning decay** over time so that good behavior is rewarded.

#### Flow
1. **IssueWarningCommand**(user_id, reason, severity, moderator_id, ack_deadline?) → Issue() | NotifyUser() → **Outbox:** user.warning.issued.v1
2. **AcknowledgeWarningCommand**(warning_id, user_id) → Acknowledge() → **Outbox:** user.warning.acknowledged.v1
3. **EscalateWarningsCommand**(user_id) → CheckThreshold() | ApplyPenalty() → **Outbox:** user.warning.escalated.v1
4. **DecayWarningsJob**() → FindEligible() | ReduceSeverity() → **Outbox:** user.warning.decayed.v1
5. **GetWarningsQuery**(user_id) → Fetch() → WarningListDTO

#### Projections
- warning_read
- warning_history_read

#### Events Published
- user.warning.issued.v1
- user.warning.acknowledged.v1
- user.warning.escalated.v1
- user.warning.decayed.v1

#### Warning Severities
- Low, Medium, High, Critical

#### Escalation Matrix
- 3 Low = 1 Medium
- 2 Medium = 1 High
- 2 High = Suspension
- 1 Critical = Immediate Suspension

#### RBAC/SLO
- **RBAC:** MODERATOR (issue), OWNER (acknowledge)
- **SLO:** P95 < 160ms

---

### 13.2 appeal/

#### User Stories
- As a **user**, I want to **appeal moderation actions** so that mistakes are corrected.
- As a **moderator**, I want to **review appeals** with evidence.
- As a **system**, I want **appeal SLAs** so that reviews are timely.

#### Flow
1. **CreateAppealCommand**(user_id, action_type, action_id, appeal_text, evidence?) → CreateAppeal() → **Outbox:** appeal.created.v1
2. **ReviewAppealCommand**(appeal_id, reviewer_id, decision, notes) → Review() | TakeAction() → **Outbox:** appeal.reviewed.v1
3. **ResolveAppealCommand**(appeal_id, resolution) → Resolve() | NotifyUser() → **Outbox:** appeal.resolved.v1
4. **GetAppealsQuery**(user_id) → Fetch() → AppealListDTO
5. **GetPendingAppealsQuery**() → FetchPending() → AppealListDTO

#### Projections
- appeal_read
- appeal_history_read

#### Events Published
- appeal.created.v1
- appeal.reviewed.v1
- appeal.resolved.v1

#### Appeal Types
- Suspension, Ban, Warning, AccountDeletion

#### Appeal Decisions
- Approved (action reversed)
- PartiallyApproved (action modified)
- Denied (action upheld)

#### RBAC/SLO
- **RBAC:** OWNER (create), MODERATOR (review)
- **SLO:** P95 < 200ms, Review within 48h

---

## 14 - RISK & FRAUD PREVENTION

### 14.1 risk/

#### User Stories
- As a **risk engine**, I want to **record risk signals** so that patterns are detected.
- As a **system**, I want **risk scoring** so that users are categorized.
- As a **moderator**, I want to **place/release holds** so that risks are mitigated.
- As a **analyst**, I want **explainable risk scores** so that decisions are transparent.

#### Flow
1. **RecordRiskSignalCommand**(user_id, signal_type, severity, metadata) → Record() → **Outbox:** risk.signal.recorded.v1
2. **ComputeRiskScoreCommand**(user_id) → AggregateSignals() | WeightFactors() | CalculateScore() → **Outbox:** risk.score.updated.v1
3. **PlaceHoldCommand**(user_id, hold_type, reason, moderator_id) → PlaceHold() → **Outbox:** risk.hold.placed.v1
4. **ReleaseHoldCommand**(hold_id, moderator_id) → ReleaseHold() → **Outbox:** risk.hold.released.v1
5. **GetRiskScoreQuery**(user_id) → FetchScore() → RiskScoreDTO (score, factors, explanation)
6. **GetRiskExplanationQuery**(user_id) → BuildExplanation() → RiskExplanationDTO (components, weights, recommendations)
7. **CorrelateSignalsCommand**(user_id) → FindPatterns() | DetectFraudRings() → **Outbox:** risk.signals.correlated.v1
8. **SendRiskAlertCommand**(user_id, alert_type) → SendToModerators() → **Outbox:** risk.alert.sent.v1

#### Projections
- risk_read
- risk_signals_read
- risk_holds_read
- risk_correlation_read

#### Events Published
- risk.signal.recorded.v1
- risk.score.updated.v1
- risk.hold.placed.v1
- risk.hold.released.v1
- risk.signals.correlated.v1
- risk.alert.sent.v1

#### Risk Signal Types
- IPGeoMismatch, SuspiciousLogin, HighDisputeRate, Chargeback, FakeDocuments, BotBehavior

#### Risk Levels
- Low: 0-30
- Medium: 31-60
- High: 61-80
- Critical: 81-100

#### Hold Types
- Withdrawal, Payout, NewContracts, ProfileChanges

#### RBAC/SLO
- **RBAC:** SYSTEM (record signals), MODERATOR (holds), ANALYST (view)
- **SLO:** P95 < 250ms

---

## 15 - SETTINGS & PREFERENCES

### 15.1 notification_preferences/

#### User Stories
- As a **user**, I want to **control notification channels** (email/sms/push/in-app) so that I'm not overwhelmed.
- As a **user**, I want **notification frequency settings** so that I control timing.
- As a **user**, I want **quiet hours** so that I'm not disturbed at night.

#### Flow
1. **UpdateNotificationPreferencesCommand**(user_id, preferences) → Validate() | Update() → **Outbox:** notification.preferences.updated.v1
2. **SetQuietHoursCommand**(user_id, start_time, end_time, timezone) → Set() → **Outbox:** notification.quiet_hours.set.v1
3. **GetNotificationPreferencesQuery**(user_id) → Fetch() → NotificationPreferencesDTO

#### Projections
- notification_preferences_read

#### Events Published
- notification.preferences.updated.v1
- notification.quiet_hours.set.v1

#### Notification Channels
- Email, SMS, Push, InApp

#### Frequency Options
- RealTime, Hourly, Daily, Weekly

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 120ms

---

### 15.2 language_preferences/

#### User Stories
- As a **user**, I want to **set my preferred language** so that the platform is localized.
- As a **system**, I want **language fallback** so that unsupported languages default gracefully.

#### Flow
1. **SetLanguageCommand**(user_id, language_code) → ValidateLanguage() | Update() → **Outbox:** language.preference.updated.v1
2. **GetLanguagePreferenceQuery**(user_id) → Fetch() → LanguagePreferenceDTO

#### Projections
- language_preferences_read

#### Events Published
- language.preference.updated.v1

#### Supported Languages
- en, es, fr, de, pt, ar, zh, ja, ru, hi

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 100ms

---

### 15.3 accessibility_settings/

#### User Stories
- As a **user**, I want **high-contrast mode** so that I can read better.
- As a **user**, I want **screen reader preferences** so that the platform is accessible.
- As a **user**, I want **font size controls** so that text is comfortable.

#### Flow
1. **UpdateAccessibilitySettingsCommand**(user_id, settings) → Update() → **Outbox:** accessibility.settings.updated.v1
2. **GetAccessibilitySettingsQuery**(user_id) → Fetch() → AccessibilitySettingsDTO

#### Projections
- accessibility_settings_read

#### Events Published
- accessibility.settings.updated.v1

#### Accessibility Options
- HighContrast, ReducedMotion, LargeFonts, ScreenReaderOptimized, KeyboardNavigation

#### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** P95 < 110ms

---

## EXTERNALIZED CAPABILITIES

The following capabilities have been **externalized to dedicated services** and are accessed via client interfaces:

### 🔐 Authentication (via Keycloak / pkg/auth)
- **JWT Verification**: Token validation, signature verification
- **RBAC**: Role-based access control, permission management
- **SSO**: Single sign-on, federated identity
- **User Federation**: LDAP, Active Directory integration

### 📁 Storage (via storage-be)
- **File Upload**: Profile pictures, portfolio images, documents
- **File Management**: Versioning, access control, expiry
- **CDN Integration**: Fast content delivery

### 💰 Financial (via financial-be)
- **Payment Verification**: Client payment method validation
- **Spending Tracking**: Total spent, forecasts, budgets
- **Earnings Tracking**: Freelancer earnings, payouts

### 📊 Analytics (via search-be)
- **User Metrics**: Profile views, search appearances
- **Engagement Tracking**: Activity patterns, retention
- **Performance Analytics**: Conversion rates, success metrics

### 📨 Communications (via communications-be)
- **Notifications**: Email, SMS, push, in-app
- **Messaging**: Direct messages, chat history
- **Email Delivery**: Transactional emails, marketing

### 📝 Contracts (via contracts-be)
- **Contract Tracking**: Active contracts, workload
- **Milestone Management**: Milestone completion tracking
- **Performance Impact**: Contract success on reputation

### ⭐ Reviews (via reviews-be)
- **Review Aggregation**: Overall ratings, feedback
- **Reputation Impact**: Review influence on reputation
- **Badge Eligibility**: TopRated, HighQuality badges

### 🔍 Search (via search-be)
- **Profile Indexing**: User profile search
- **Skill Matching**: Job-skill matching
- **Recommendations**: AI-powered suggestions

### 🎫 Subscriptions (via subscriptions-be)
- **Premium Features**: Tier management, feature access
- **Connects Management**: Connect balances, purchases
- **Subscription Status**: Active/expired subscriptions

### 🛡️ Compliance (via compliance-be)
- **KYC Processing**: Document verification, status
- **Sanctions Screening**: OFAC, PEP checks
- **Audit Trails**: Compliance event logging

---

## EVENT FLOWS

### Events Published by users-be
```
# Core User
- user.registered.v1
- user.guest.created.v1
- user.upgraded.v1
- user.updated.v1
- user.deactivated.v1
- user.reactivated.v1
- user.deleted.v1
- user.username.changed.v1

# Profile
- profile.updated.v1
- profile.picture.updated.v1
- profile.completeness.changed.v1
- profile.visibility.changed.v1
- profile.stealth.enabled.v1

# Skills & Experience
- skill.added.v1
- skill.updated.v1
- skill.endorsed.v1
- experience.added.v1
- education.added.v1
- certification.verified.v1
- portfolio.item.created.v1

# Freelancer & Client
- freelancer.profile.created.v1
- freelancer.success_rate.updated.v1
- client.profile.created.v1
- client.payment.verified.v1

# Verification
- kyc.submitted.v1
- kyc.approved.v1
- badge.awarded.v1
- endorsement.created.v1

# Account Health
- health.score.updated.v1
- health.issue.detected.v1
- reputation.updated.v1

# Organizations
- org.created.v1
- org.member.added.v1
- org.seats.updated.v1
- talent_pool.created.v1

# Security
- user.login.recorded.v1
- user.session.revoked.v1
- two_factor.enabled.v1
- device.registered.v1
- recovery.completed.v1

# Privacy
- privacy.settings.updated.v1
- data.export.requested.v1
- data.deletion.requested.v1
- consent.updated.v1

# Saved Items
- item.saved.v1
- collection.created.v1
- reminder.scheduled.v1

# Moderation
- user.blocked.v1
- user.flagged.v1
- user.suspended.v1
- user.banned.v1
- user.warning.issued.v1
- appeal.created.v1

# Risk
- risk.signal.recorded.v1
- risk.score.updated.v1
- risk.hold.placed.v1
- risk.alert.sent.v1

# Settings
- notification.preferences.updated.v1
- language.preference.updated.v1
- accessibility.settings.updated.v1
```

### Events Consumed by users-be
```
# From Keycloak
- keycloak.user.created
- keycloak.user.updated
- keycloak.user.deleted

# From contracts-be
- contract.created
- contract.started
- contract.completed

# From reviews-be
- review.submitted
- rating.given

# From financial-be
- payment.processed
- payment.failed
- payout.completed

# From jobs-be
- job.posted
- job.closed

# From proposals-be
- proposal.submitted
- proposal.accepted

# From storage-be
- file.uploaded
- file.deleted

# From compliance-be
- compliance.check.completed
- sanctions.check.completed

# From subscriptions-be
- subscription.created
- subscription.expired
- connects.purchased
```

---

## ARCHITECTURE PATTERNS

### Domain-Driven Design (DDD)
- **Aggregates**: User, Profile, Org, Security
- **Value Objects**: Money, Location, Schedule, Score
- **Domain Events**: All state changes emit events
- **Repositories**: Interface-based, implementation in infrastructure

### CQRS (Command Query Responsibility Segregation)
- **Commands**: Write operations with business logic
- **Queries**: Read operations from optimized projections
- **Projections**: Denormalized read models for performance

### Event Sourcing (Partial)
- **Audit Trails**: Full event history for compliance
- **Replay Capability**: Rebuild state from events
- **Temporal Queries**: State at any point in time

### Outbox Pattern
- **Reliable Publishing**: DB transaction + outbox table
- **At-Least-Once Delivery**: Outbox processor with retries
- **Idempotency**: Event deduplication in consumers

### Clean Architecture Layers
1. **Domain**: Pure business logic, no dependencies
2. **Application**: Use cases, orchestration, DTOs
3. **Infrastructure**: DB, Kafka, Redis, external services
4. **Interface**: HTTP handlers, middleware, routes

---

## DATA CONSISTENCY

### Within Service
- **Strong Consistency**: ACID transactions for writes
- **Immediate Consistency**: Within same aggregate

### Across Services
- **Eventual Consistency**: Via event propagation
- **Compensation**: Sagas for distributed transactions
- **Idempotency**: All event handlers are idempotent

### Caching Strategy
- **Redis Cache**: Hot data with 15-minute TTL
- **Cache Invalidation**: Event-driven, on updates
- **Read-Through**: Populate cache on miss
- **Write-Through**: Update cache on write

---

## SECURITY CONSIDERATIONS

### Authentication
- **Keycloak Integration**: Via pkg/auth
- **JWT Tokens**: Stateless authentication
- **Token Refresh**: Automatic token renewal
- **MFA**: Multi-factor authentication support

### Authorization
- **RBAC**: Role-based access control
- **Resource Ownership**: Users own their data
- **Admin Roles**: MODERATOR, TRUST_SAFETY_ADMIN
- **Org Roles**: ORG_OWNER, ORG_ADMIN, ORG_MEMBER

### Data Protection
- **PII Encryption**: Field-level encryption at rest
- **TLS**: All communication encrypted in transit
- **Key Rotation**: Regular encryption key rotation
- **Audit Logging**: All access logged with actor

### Input Validation
- **Multi-Layer**: API, Application, Domain
- **Sanitization**: XSS prevention, SQL injection protection
- **Rate Limiting**: Per-user, per-endpoint limits
- **CAPTCHA**: Bot protection on sensitive endpoints

---

## SCALABILITY CONSIDERATIONS

### Horizontal Scaling
- **Stateless Services**: No session state in memory
- **Load Balancing**: Round-robin across instances
- **Database Connection Pooling**: Efficient resource usage

### Database Optimization
- **Indexes**: Strategic indexes on common queries
- **Partitioning**: Large tables partitioned by date/region
- **Read Replicas**: For analytics and reporting
- **Query Optimization**: N+1 query prevention

### Caching
- **Redis Cluster**: Distributed caching
- **Cache Warming**: Pre-populate hot data
- **Cache Stampede Prevention**: Singleflight pattern

### Background Jobs
- **Leader Election**: Single instance for cron jobs
- **Job Queues**: Async processing for heavy tasks
- **Retry Logic**: Exponential backoff with max attempts

---

## OBSERVABILITY

### Logging
- **Structured Logs**: JSON format via platform-shared/logging
- **Log Levels**: DEBUG, INFO, WARN, ERROR
- **Correlation IDs**: Request tracing across services
- **PII Redaction**: Sensitive data masked in logs

### Metrics
- **Prometheus**: Metrics collection via platform-shared/metrics
- **Business Metrics**: User growth, engagement, health scores
- **Technical Metrics**: Latency, error rates, throughput
- **Custom Metrics**: Domain-specific KPIs

### Tracing
- **OpenTelemetry**: Distributed tracing via platform-shared/tracing
- **Span Attributes**: Enrich traces with context
- **Sampling**: Smart sampling for performance

### Alerting
- **SLO-Based**: Alerts on SLO violations
- **Error Rate**: Threshold-based error alerts
- **Latency**: P95, P99 latency alerts
- **Business Events**: Critical events (bans, risk alerts)

---

## TESTING STRATEGY

### Unit Tests
- **Domain Logic**: Business rules, validations
- **Value Objects**: Immutability, equality
- **Services**: Mocked dependencies
- **Coverage Target**: >80%

### Integration Tests
- **Repository Tests**: Real database interactions
- **Event Handlers**: End-to-end event processing
- **API Tests**: HTTP handler integration
- **Test Containers**: Docker-based test infrastructure

### E2E Tests
- **User Journeys**: Complete flows (registration → verification → profile completion)
- **Cross-Service**: Multi-service scenarios
- **Performance Tests**: Load testing critical paths

### Contract Tests
- **Event Schemas**: Validate protobuf contracts
- **API Contracts**: OpenAPI validation
- **Backward Compatibility**: Schema evolution tests

---

## SLO SUMMARY

| Domain | P95 Latency | P99 Latency | Availability |
|--------|-------------|-------------|--------------|
| Core User Management | < 200ms | < 500ms | 99.95% |
| Profile Management | < 180ms | < 400ms | 99.9% |
| Skills & Experience | < 150ms | < 350ms | 99.9% |
| Verification (KYC) | < 300ms | < 800ms | 99.5% |
| Account Health | < 250ms | < 600ms | 99.5% |
| Organizations | < 200ms | < 500ms | 99.9% |
| Security | < 180ms | < 450ms | 99.95% |
| Privacy | < 180ms | < 400ms | 99.9% |
| Saved Items | < 130ms | < 300ms | 99.9% |
| Blocking | < 110ms | < 250ms | 99.9% |
| Suspension & Bans | < 180ms | < 450ms | 99.95% |
| Risk & Fraud | < 250ms | < 600ms | 99.5% |
| Settings | < 120ms | < 280ms | 99.9% |

---

## MIGRATION & DEPLOYMENT

### Database Migrations
- **Auto-Migration**: GORM AutoMigrate on startup
- **Version Tracking**: SchemaVersion table
- **Safety Checks**: Pre-migration validation
- **Rollback Support**: Manual rollback procedures

### Zero-Downtime Deployment
- **Blue-Green**: Parallel deployments
- **Rolling Updates**: Gradual rollout
- **Health Checks**: Kubernetes readiness/liveness
- **Graceful Shutdown**: Finish in-flight requests

### Feature Flags
- **Toggle New Features**: Enable/disable without deploy
- **Canary Releases**: Gradual feature rollout
- **A/B Testing**: Experimental feature testing

---

## GLOSSARY

- **Aggregate**: Domain object cluster with transactional consistency
- **Projection**: Denormalized read model built from events
- **Outbox**: Pattern for reliable event publishing
- **Idempotency**: Safe to retry operations without side effects
- **CQRS**: Command Query Responsibility Segregation
- **DDD**: Domain-Driven Design
- **KYC**: Know Your Customer (identity verification)
- **RBAC**: Role-Based Access Control
- **2FA**: Two-Factor Authentication
- **PII**: Personally Identifiable Information
- **GDPR**: General Data Protection Regulation
- **SLO**: Service Level Objective
- **P95/P99**: 95th/99th percentile latency

---

## IMPLEMENTATION PRIORITIES

### Phase 1: Core (MVP)
1. User registration & authentication (1.1)
2. Profile management (2.1)
3. Skills & experience (3.1-3.3)
4. Freelancer/client profiles (4.1-4.2)
5. Basic security (8.1-8.2)

### Phase 2: Trust & Safety
1. Verification (KYC) (5.1)
2. Account health (6.1)
3. Blocking (11.1)
4. Warnings (13.1)
5. Risk scoring (14.1)

### Phase 3: Enhanced Features
1. Organizations (7.1-7.2)
2. Badges & endorsements (5.2-5.3)
3. Portfolio (3.5)
4. Saved items (10.1-10.3)
5. Advanced security (8.3, two-factor)

### Phase 4: Compliance & Advanced
1. Privacy & GDPR (9.1-9.3)
2. Suspension & bans (12.1-12.2)
3. Appeals (13.2)
4. Advanced risk (14.1 full)
5. Settings & preferences (15.1-15.3)

---

**End of users-be User Stories v2.0**