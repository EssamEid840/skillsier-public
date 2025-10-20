# 📦 **users-be - User Management Service - Complete User Stories**

---

## **1 - CORE USER DOMAIN**

### 1.1 user/

#### User Stories
- As a **new user**, I want to **register with email and password** so that I can create an account on the platform.
- As a **user**, I want to **verify my email address** so that I can activate my account and access platform features.
- As a **user**, I want to **update my profile information** (first name, last name, phone) so that I keep my account details current.
- As a **user**, I want to **search for other users by username or email** so that I can find colleagues or connections.
- As an **admin**, I want to **suspend or ban problematic users** so that I can maintain platform integrity.
- As a **system**, I want to **validate email uniqueness** so that duplicate accounts are prevented.
- As a **system**, I want to **track user account status** (Active, Suspended, Banned) so that access control is enforced.

#### Flow
1. **CreateUserCommand**(email, password, user_type, first_name, last_name) → ValidateEmail() | HashPassword() | Persist() → **Outbox:** user.created.v1
2. **VerifyEmailCommand**(user_id, verification_token) → ValidateToken() | ActivateAccount() → **Outbox:** user.verified.v1
3. **UpdateUserCommand**(user_id, updates) → AuthorizeOwner() | ValidateFields() | Apply() → **Outbox:** user.updated.v1
4. **SearchUsersQuery**(query, filters) → ApplyFilters() | Paginate() → UserListDTO
5. **GetUserQuery**(user_id) → AuthorizeAccess() | Fetch() → UserDTO
6. **GetUserStatisticsQuery**(user_id) → Aggregate() → UserStatisticsDTO

#### Projections
- user_read
- user_stats_read
- user_search_index

#### Events Published
- user.created.v1
- user.updated.v1
- user.verified.v1
- user.suspended.v1
- user.banned.v1
- user.deleted.v1

#### RBAC/SLO
- **RBAC:** OWNER (update), ADMIN (suspend/ban), PUBLIC (search/view)
- **SLO:** P95 < 150ms (read), P95 < 200ms (write)

---

### 1.2 profile/

#### User Stories
- As a **user**, I want to **complete my extended profile** (bio, location, profile picture) so that I can present myself professionally.
- As a **user**, I want to **set my preferences** (language, timezone, currency) so that the platform adapts to my needs.
- As a **user**, I want to **upload a profile picture** so that others can recognize me.
- As a **freelancer**, I want to **see my profile completeness score** so that I know what sections to improve.
- As a **system**, I want to **reference availability data from availability service** so that profile displays current availability status.

#### Flow
1. **UpdateProfileCommand**(user_id, bio, location, profile_picture_url) → AuthorizeOwner() | Validate() | Persist() → **Outbox:** profile.updated.v1
2. **UpdatePreferencesCommand**(user_id, language, timezone, currency) → Validate() | Apply() → **Outbox:** preferences.updated.v1
3. **GetProfileQuery**(user_id) → Fetch() | EnrichWithAvailability() → ProfileDTO
4. **GetProfileCompletionQuery**(user_id) → Calculate() → ProfileCompletenessDTO

#### Projections
- profile_read
- profile_completeness_read

#### Events Published
- profile.updated.v1
- preferences.updated.v1

#### Events Consumed
- availability.updated.v1 (to enrich profile display)

#### RBAC/SLO
- **RBAC:** OWNER (update), PUBLIC (view)
- **SLO:** P95 < 180ms

---

## **2 - CAPABILITIES DOMAIN (CONSOLIDATED)**

### 2.1 capabilities/

#### User Stories
- As a **freelancer**, I want to **add skills with proficiency levels** so that clients know my expertise.
- As a **freelancer**, I want to **specify years of experience per skill** so that I demonstrate depth of knowledge.
- As a **freelancer**, I want to **add specializations** (e.g., "React + TypeScript for FinTech") so that I stand out in niche markets.
- As a **system**, I want to **map skills to standardized taxonomy** (React → WebDev → Engineering) so that search and matching work effectively.
- As a **freelancer**, I want to **verify my specializations** so that I gain credibility.
- As a **system**, I want to **track capability updates** so that profile freshness is maintained.

#### Flow
1. **AddSkillCommand**(user_id, skill_name, proficiency, years_exp) → ValidateSkill() | MapToTaxonomy() | Persist() → **Outbox:** skill.added.v1
2. **UpdateSkillCommand**(user_id, skill_id, proficiency, years_exp) → AuthorizeOwner() | Update() → **Outbox:** skill.updated.v1
3. **RemoveSkillCommand**(user_id, skill_id) → AuthorizeOwner() | Delete() → **Outbox:** skill.removed.v1
4. **AddSpecializationCommand**(user_id, specialization_data) → Validate() | Persist() → **Outbox:** specialization.added.v1
5. **VerifySpecializationCommand**(user_id, specialization_id) → ValidateEvidence() | Approve() → **Outbox:** specialization.verified.v1
6. **ListSkillsQuery**(user_id) → Fetch() → SkillListDTO
7. **GetSkillTaxonomyQuery**(skill_name) → Fetch() → TaxonomyDTO

#### Projections
- capabilities_read
- skill_taxonomy_read
- specializations_read

#### Events Published
- skill.added.v1
- skill.updated.v1
- skill.removed.v1
- specialization.added.v1
- specialization.verified.v1

#### RBAC/SLO
- **RBAC:** OWNER (add/update/remove), ADMIN (verify specializations), PUBLIC (view)
- **SLO:** P95 < 160ms

---

### 2.2 service_catalog/

#### User Stories
- As a **freelancer**, I want to **create service offerings** so that clients can purchase predefined packages.
- As a **freelancer**, I want to **reference my capabilities in services** so that clients see what skills apply.
- As a **freelancer**, I want to **create service packages** (Basic, Standard, Premium) so that I offer tiered pricing.
- As a **client**, I want to **browse a freelancer's service catalog** so that I can quickly purchase services.
- As a **system**, I want to **join service data with capabilities** so that service descriptions are enriched.

#### Flow
1. **CreateServiceCommand**(user_id, name, description, capability_ids, price, duration) → ValidateCapabilityRefs() | Persist() → **Outbox:** service.created.v1
2. **UpdateServiceCommand**(service_id, updates) → AuthorizeOwner() | Update() → **Outbox:** service.updated.v1
3. **CreateServicePackageCommand**(user_id, package_data) → Validate() | Persist() → **Outbox:** package.created.v1
4. **GetServiceCatalogQuery**(user_id) → FetchWithCapabilities() → ServiceCatalogDTO

#### Projections
- service_catalog_read

#### Events Published
- service.created.v1
- service.updated.v1
- package.created.v1

#### RBAC/SLO
- **RBAC:** OWNER (create/update), PUBLIC (view)
- **SLO:** P95 < 190ms

---

## **3 - EXPERIENCE & EDUCATION DOMAIN**

### 3.1 experience/

#### User Stories
- As a **freelancer**, I want to **add work experience entries** so that clients see my professional background.
- As a **freelancer**, I want to **mark current positions** so that my profile shows ongoing work.
- As a **freelancer**, I want to **update or remove experience entries** so that my profile stays accurate.
- As a **system**, I want to **validate date ranges** so that experience timelines are logical.

#### Flow
1. **AddExperienceCommand**(user_id, company, title, description, start_date, end_date, is_current) → ValidateDates() | Persist() → **Outbox:** experience.added.v1
2. **UpdateExperienceCommand**(experience_id, updates) → AuthorizeOwner() | Validate() | Update() → **Outbox:** experience.updated.v1
3. **DeleteExperienceCommand**(experience_id) → AuthorizeOwner() | Delete() → **Outbox:** experience.removed.v1
4. **ListExperienceQuery**(user_id) → Fetch() → ExperienceListDTO

#### Projections
- experience_read

#### Events Published
- experience.added.v1
- experience.updated.v1
- experience.removed.v1

#### RBAC/SLO
- **RBAC:** OWNER (add/update/delete), PUBLIC (view)
- **SLO:** P95 < 140ms

---

### 3.2 education/

#### User Stories
- As a **freelancer**, I want to **add educational background** so that clients see my qualifications.
- As a **freelancer**, I want to **specify degree, field of study, and graduation year** so that education is properly documented.
- As a **system**, I want to **validate graduation year ranges** so that data integrity is maintained.

#### Flow
1. **AddEducationCommand**(user_id, school, degree, field, graduation_year, description) → Validate() | Persist() → **Outbox:** education.added.v1
2. **UpdateEducationCommand**(education_id, updates) → AuthorizeOwner() | Update() → **Outbox:** education.updated.v1
3. **DeleteEducationCommand**(education_id) → AuthorizeOwner() | Delete() → **Outbox:** education.removed.v1
4. **ListEducationQuery**(user_id) → Fetch() → EducationListDTO

#### Projections
- education_read

#### Events Published
- education.added.v1
- education.updated.v1
- education.removed.v1

#### RBAC/SLO
- **RBAC:** OWNER (add/update/delete), PUBLIC (view)
- **SLO:** P95 < 140ms

---

### 3.3 language/

#### User Stories
- As a **freelancer**, I want to **add language proficiencies** so that clients know what languages I can work in.
- As a **freelancer**, I want to **specify proficiency levels** (Beginner, Intermediate, Advanced, Expert) so that my abilities are clear.
- As a **client**, I want to **filter freelancers by language** so that I find suitable candidates.

#### Flow
1. **AddLanguageCommand**(user_id, language_code, proficiency_level) → ValidateLanguageCode() | Persist() → **Outbox:** language.added.v1
2. **UpdateLanguageCommand**(language_id, proficiency_level) → AuthorizeOwner() | Update() → **Outbox:** language.updated.v1
3. **RemoveLanguageCommand**(language_id) → AuthorizeOwner() | Delete() → **Outbox:** language.removed.v1
4. **ListLanguagesQuery**(user_id) → Fetch() → LanguageListDTO

#### Projections
- language_read

#### Events Published
- language.added.v1
- language.updated.v1
- language.removed.v1

#### RBAC/SLO
- **RBAC:** OWNER (add/update/remove), PUBLIC (view)
- **SLO:** P95 < 130ms

---

## **4 - CREDENTIALS DOMAIN (CONSOLIDATED)**

### 4.1 credentials/

#### User Stories
- As a **freelancer**, I want to **add external certifications** (AWS, Google, Microsoft) so that I demonstrate verified expertise.
- As a **freelancer**, I want to **submit verification documents** for external certifications so that they are validated.
- As a **system**, I want to **verify external certifications** so that only legitimate credentials are displayed.
- As a **freelancer**, I want to **earn platform certifications** (Upwork badges) so that I gain trust on the platform.
- As a **system**, I want to **issue platform certifications based on exams** so that skill verification is standardized.
- As a **system**, I want to **track certification expiry and recertification** so that credentials remain current.

#### Flow
1. **AddExternalCertificationCommand**(user_id, issuer, name, credential_url, issued_date, expiry_date) → Validate() | Persist() → **Outbox:** external_certification.added.v1
2. **VerifyExternalCertificationCommand**(cert_id, verification_status, verified_by) → ValidateEvidence() | Update() → **Outbox:** external_certification.verified.v1
3. **EarnPlatformCertificationCommand**(user_id, certification_type, exam_result) → ValidateExamPass() | IssueCert() → **Outbox:** platform_certification.earned.v1
4. **RenewPlatformCertificationCommand**(cert_id, renewal_exam_result) → ValidateRenewal() | Update() → **Outbox:** platform_certification.renewed.v1
5. **GetCredentialsQuery**(user_id) → Fetch() → CredentialsDTO

#### Projections
- credentials_read
- external_certifications_read
- platform_certifications_read

#### Events Published
- external_certification.added.v1
- external_certification.verified.v1
- external_certification.rejected.v1
- external_certification.expired.v1
- platform_certification.earned.v1
- platform_certification.renewed.v1

#### RBAC/SLO
- **RBAC:** OWNER (add external), ADMIN (verify external), SYSTEM (issue platform), PUBLIC (view verified)
- **SLO:** P95 < 170ms

---

## **5 - PORTFOLIO & SHOWCASE DOMAIN**

### 5.1 portfolio/

#### User Stories
- As a **freelancer**, I want to **add portfolio items** so that I can showcase my work.
- As a **freelancer**, I want to **upload images, videos, and documents** so that clients see tangible examples.
- As a **freelancer**, I want to **reorder portfolio items** so that my best work appears first.
- As a **client**, I want to **browse a freelancer's portfolio** so that I can assess work quality.
- As a **system**, I want to **validate media types and file sizes** so that uploads are safe.

#### Flow
1. **AddPortfolioItemCommand**(user_id, title, description, url, thumbnail_url, media_files) → ValidateFiles() | Upload() | Persist() → **Outbox:** portfolio_item.added.v1
2. **UpdatePortfolioItemCommand**(item_id, updates) → AuthorizeOwner() | Update() → **Outbox:** portfolio_item.updated.v1
3. **DeletePortfolioItemCommand**(item_id) → AuthorizeOwner() | Delete() → **Outbox:** portfolio_item.removed.v1
4. **ReorderPortfolioCommand**(user_id, item_order) → AuthorizeOwner() | UpdateOrder() → **Outbox:** portfolio.reordered.v1
5. **ListPortfolioQuery**(user_id) → Fetch() → PortfolioListDTO

#### Projections
- portfolio_read

#### Events Published
- portfolio_item.added.v1
- portfolio_item.updated.v1
- portfolio_item.removed.v1
- portfolio.reordered.v1

#### Events Consumed
- storage.file.uploaded
- storage.file.deleted

#### RBAC/SLO
- **RBAC:** OWNER (add/update/delete/reorder), PUBLIC (view)
- **SLO:** P95 < 250ms (upload), P95 < 140ms (list)

---

## **6 - USER TYPE SPECIFIC DOMAINS**

### 6.1 freelancer/

#### User Stories
- As a **freelancer**, I want to **set my professional title and overview** so that clients understand my expertise.
- As a **freelancer**, I want to **set my hourly rate and minimum budget** so that clients know my pricing.
- As a **freelancer**, I want to **record a video introduction** so that I can personalize my profile.
- As a **system**, I want to **track freelancer statistics** (total jobs, earnings, success rate) so that reputation is measurable.
- As a **freelancer**, I want to **view my job statistics** so that I can track my performance.

#### Flow
1. **UpdateFreelancerProfileCommand**(user_id, title, overview, video_intro_url) → AuthorizeOwner() | Validate() | Persist() → **Outbox:** freelancer_profile.updated.v1
2. **UpdateRatesCommand**(user_id, hourly_rate, minimum_budget, currency) → AuthorizeOwner() | ValidateRates() | Update() → **Outbox:** rates.updated.v1
3. **UpdateStatsCommand**(user_id, stats_updates) → ValidateSource() | Update() → **Outbox:** freelancer_stats.updated.v1
4. **GetFreelancerStatsQuery**(user_id) → Fetch() → FreelancerStatsDTO

#### Projections
- freelancer_read
- freelancer_stats_read

#### Events Published
- freelancer_profile.updated.v1
- rates.updated.v1
- freelancer_stats.updated.v1

#### Events Consumed
- contract.completed.v1 (to update stats)
- payment.received.v1 (to update earnings)

#### RBAC/SLO
- **RBAC:** OWNER (update), PUBLIC (view)
- **SLO:** P95 < 160ms

---

### 6.2 client/

#### User Stories
- As a **client**, I want to **link my profile to an organization** so that I can hire on behalf of my company.
- As a **client**, I want to **view my hiring statistics** (total hires, total spent, active contracts) so that I track my activity.
- As a **system**, I want to **reference organization data from org service** so that there's no duplication of company information.
- As a **system**, I want to **update client statistics based on contracts** so that metrics stay current.

#### Flow
1. **UpdateClientProfileCommand**(user_id, client_data) → AuthorizeOwner() | Validate() | Persist() → **Outbox:** client_profile.updated.v1
2. **LinkToOrgCommand**(user_id, org_id) → ValidateOrgExists() | Link() → **Outbox:** client.linked_to_org.v1
3. **GetClientStatsQuery**(user_id) → Fetch() → ClientStatsDTO

#### Projections
- client_read
- client_stats_read

#### Events Published
- client_profile.updated.v1
- client.linked_to_org.v1
- client_stats.updated.v1

#### Events Consumed
- contract.created.v1 (to update hiring stats)
- payment.sent.v1 (to update spending)
- org.updated.v1 (to refresh org reference)

#### RBAC/SLO
- **RBAC:** OWNER (update/link), PUBLIC (view)
- **SLO:** P95 < 150ms

---

## **7 - IDENTITY VERIFICATION DOMAIN (CONSOLIDATED)**

### 7.1 identity_verification/

#### User Stories
- As a **freelancer**, I want to **submit KYC documents** (ID, passport, proof of address, selfie) so that I can verify my identity.
- As a **client organization**, I want to **submit KYB documents** so that I can verify my business.
- As a **system**, I want to **validate submitted documents** so that only legitimate users are verified.
- As an **admin**, I want to **approve or reject identity verification requests** so that verification quality is maintained.
- As a **user**, I want to **track my verification status** so that I know when my account is verified.

#### Flow
1. **SuspendUserCommand**(user_id, reason, duration, suspended_by) → ValidateReason() | ApplySuspension() → **Outbox:** user.suspended.v1
2. **UnsuspendUserCommand**(user_id, unsuspended_by) → ValidateActive() | Release() → **Outbox:** user.unsuspended.v1
3. **BanUserCommand**(user_id, reason, is_permanent, expires_at, banned_by) → ValidateReason() | ApplyBan() → **Outbox:** user.banned.v1
4. **UnbanUserCommand**(user_id, unbanned_by) → ValidateActive() | Release() → **Outbox:** user.unbanned.v1
5. **IssueWarningCommand**(user_id, reason, severity, issued_by) → ValidateReason() | IssueWarning() → **Outbox:** warning.issued.v1
6. **AcknowledgeWarningCommand**(warning_id, user_id) → ValidateOwner() | Acknowledge() → **Outbox:** warning.acknowledged.v1
7. **GetModerationHistoryQuery**(user_id) → Fetch() → ModerationHistoryDTO

#### Projections
- moderation_read
- moderation_history_read

#### Events Published
- user.suspended.v1
- user.unsuspended.v1
- user.banned.v1
- user.unbanned.v1
- warning.issued.v1
- warning.acknowledged.v1

#### RBAC/SLO
- **RBAC:** ADMIN (suspend/ban/warn/release), OWNER (acknowledge warning), ADMIN (view history)
- **SLO:** P95 < 180ms

---

## **8 - TRUST DOMAIN (CONSOLIDATED)**

### 8.1 trust/

#### User Stories
- As a **user**, I want to **have a trust level calculated** (Unverified, Basic, Enhanced, Premium) so that clients can gauge my reliability.
- As a **system**, I want to **award trust badges** (VerifiedPayment, IDVerified) so that trust signals are visible.
- As a **system**, I want to **revoke trust badges** if conditions change so that trust remains accurate.
- As a **user**, I want to **see my trust level and badges** so that I understand my standing.
- As a **client**, I want to **filter freelancers by trust level** so that I find reliable candidates.

#### Flow
1. **CalculateTrustLevelCommand**(user_id) → AggregateSignals() | ComputeLevel() | Update() → **Outbox:** trust_level.updated.v1
2. **AwardTrustBadgeCommand**(user_id, badge_type) → ValidateEligibility() | Award() → **Outbox:** trust_badge.awarded.v1
3. **RevokeTrustBadgeCommand**(user_id, badge_type, reason) → Validate() | Revoke() → **Outbox:** trust_badge.revoked.v1
4. **GetTrustLevelQuery**(user_id) → Fetch() → TrustLevelDTO
5. **ListTrustBadgesQuery**(user_id) → Fetch() → TrustBadgeListDTO

#### Projections
- trust_read
- trust_badge_read

#### Events Published
- trust_level.updated.v1
- trust_badge.awarded.v1
- trust_badge.revoked.v1

#### Events Consumed
- identity_verification.approved.v1 (to award IDVerified)
- payment.verified.v1 (to award VerifiedPayment)

#### RBAC/SLO
- **RBAC:** SYSTEM (calculate/award/revoke), PUBLIC (view)
- **SLO:** P95 < 180ms

---

## **9 - BADGING DOMAIN (CONSOLIDATED)**

### 9.1 badging/

#### User Stories
- As a **freelancer**, I want to **earn achievement badges** (FirstJob, TopRated, QuickResponder) so that I showcase my accomplishments.
- As a **freelancer**, I want to **earn certification badges** when I complete platform exams so that my skills are verified.
- As a **system**, I want to **issue trust badges** based on verification events so that trust is signaled.
- As a **system**, I want to **issue platform badges** (RisingTalent, TopRated, ExpertVetted) so that top performers are recognized.
- As a **system**, I want to **revoke badges** if criteria are no longer met so that badges remain meaningful.
- As a **user**, I want to **view all my badges** so that I see my achievements.

#### Flow
1. **AwardBadgeCommand**(user_id, badge_type, badge_slug, metadata) → ValidateEligibility() | Issue() → **Outbox:** badge.awarded.v1
2. **RevokeBadgeCommand**(badge_id, reason) → Validate() | Revoke() → **Outbox:** badge.revoked.v1
3. **CheckBadgeEligibilityQuery**(user_id, badge_type) → EvaluateCriteria() → EligibilityDTO
4. **GetBadgesQuery**(user_id) → Fetch() → BadgeListDTO
5. **GetBadgesByTypeQuery**(user_id, badge_type) → Filter() → BadgeListDTO

#### Projections
- badge_read

#### Events Published
- badge.awarded.v1
- badge.revoked.v1

#### Events Consumed
- achievement.unlocked.v1 (to issue achievement badge)
- platform_certification.earned.v1 (to issue certification badge)
- external_certification.verified.v1 (to issue certification badge)
- trust_badge.awarded.v1 (to issue trust badge)

#### RBAC/SLO
- **RBAC:** SYSTEM (award/revoke), PUBLIC (view)
- **SLO:** P95 < 160ms

---

## **10 - SETTINGS & PREFERENCES DOMAINS**

### 10.1 settings/

#### User Stories
- As a **user**, I want to **configure my notification preferences** (email, SMS, push, in-app) so that I control how I'm contacted.
- As a **user**, I want to **set my theme preference** (light/dark mode) so that the UI matches my preference.
- As a **user**, I want to **configure language, timezone, and currency** so that the platform is localized.
- As a **system**, I want to **store settings as JSON** so that new settings can be added without schema changes.

#### Flow
1. **UpdateSettingsCommand**(user_id, settings_json) → Validate() | Merge() | Persist() → **Outbox:** settings.updated.v1
2. **UpdateNotificationPrefsCommand**(user_id, email, sms, push, in_app) → Validate() | Update() → **Outbox:** notification_prefs.updated.v1
3. **GetSettingsQuery**(user_id) → Fetch() → SettingsDTO
4. **GetNotificationPrefsQuery**(user_id) → Fetch() → NotificationPrefsDTO

#### Projections
- settings_read

#### Events Published
- settings.updated.v1
- notification_prefs.updated.v1

#### RBAC/SLO
- **RBAC:** OWNER (update), OWNER (view)
- **SLO:** P95 < 140ms

---

### 10.2 privacy/

#### User Stories
- As a **user**, I want to **control email visibility** (show/hide email) so that I manage my privacy.
- As a **user**, I want to **control phone visibility** (show/hide phone) so that I control contact methods.
- As a **user**, I want to **control activity sharing** so that I decide what others see.
- As a **user**, I want to **enable/disable direct contact** so that I control who can message me.
- As a **system**, I want to **separate privacy settings from general settings** so that privacy controls are explicit.

#### Flow
1. **UpdatePrivacySettingsCommand**(user_id, show_email, show_phone, share_activity, allow_direct_contact) → Validate() | Update() → **Outbox:** privacy_settings.updated.v1
2. **GetPrivacySettingsQuery**(user_id) → Fetch() → PrivacySettingsDTO

#### Projections
- privacy_read

#### Events Published
- privacy_settings.updated.v1

#### RBAC/SLO
- **RBAC:** OWNER (update/view)
- **SLO:** P95 < 130ms

---

### 10.3 saved_items/

#### User Stories
- As a **user**, I want to **save jobs for later** so that I can apply when ready.
- As a **client**, I want to **save freelancers** so that I can contact them for future projects.
- As a **user**, I want to **add notes to saved items** so that I remember why I saved them.
- As a **user**, I want to **unsave items** so that my saved list stays relevant.
- As a **user**, I want to **search my saved items** so that I can find them easily.

#### Flow
1. **SaveItemCommand**(user_id, item_type, item_id, notes) → Validate() | CheckDuplicate() | Persist() → **Outbox:** item.saved.v1
2. **UnsaveItemCommand**(user_id, item_id) → AuthorizeOwner() | Delete() → **Outbox:** item.unsaved.v1
3. **ListSavedItemsQuery**(user_id, item_type) → Filter() | Fetch() → SavedItemListDTO
4. **SearchSavedQuery**(user_id, query) → Search() → SavedItemListDTO

#### Projections
- saved_items_read

#### Events Published
- item.saved.v1
- item.unsaved.v1

#### RBAC/SLO
- **RBAC:** OWNER (save/unsave/view)
- **SLO:** P95 < 140ms

---

### 10.4 blocked_users/

#### User Stories
- As a **user**, I want to **block other users** so that I don't see their content or receive messages.
- As a **user**, I want to **provide a reason for blocking** so that context is recorded.
- As a **user**, I want to **unblock users** so that I can restore communication if needed.
- As a **system**, I want to **enforce blocking in messaging and search** so that blocked users can't interact.

#### Flow
1. **BlockUserCommand**(blocker_id, blocked_id, reason) → Validate() | Persist() → **Outbox:** user.blocked.v1
2. **UnblockUserCommand**(blocker_id, blocked_id) → Validate() | Delete() → **Outbox:** user.unblocked.v1
3. **ListBlockedUsersQuery**(blocker_id) → Fetch() → BlockedUserListDTO

#### Projections
- blocked_users_read

#### Events Published
- user.blocked.v1
- user.unblocked.v1

#### RBAC/SLO
- **RBAC:** OWNER (block/unblock/view)
- **SLO:** P95 < 130ms

---
11 - Moderation (Consolidated)
-----------------------------

### 11.1 Moderation Aggregate

#### Stories

*   As a **system**, I want a **single moderation aggregate per user** so that all actions are tracked coherently.
    
*   As an **admin**, I want to **view a user’s active moderation state** (clean/limited/warned/suspended/banned) so that I can act quickly.
    
*   As a **system**, I want **chronological action history** so that repeat offenses are visible.
    

#### Flow

*   **GetModerationAggregateQuery(user\_id)** → AuthorizeAdmin() | FetchAggregate() → ModerationAggregateDTO
    
*   **GetActiveStatusQuery(user\_id)** → CheckSuspension() | CheckBan() | CheckUnackWarnings() → ActiveStatusDTO
    
*   **GetActionCountsQuery(user\_id)** → CountByType() → ModerationActionCountsDTO
    

#### Projections

*   moderation\_aggregate\_read
    

#### Events

*   moderation.aggregate.created.v1, moderation.aggregate.updated.v1
    

#### RBAC/SLO

*   **RBAC:** ADMIN view; SYSTEM update
    
*   **SLO:** reads **P95 < 100 ms**
    

### 11.2 Suspension

#### Stories

*   As an **admin**, I want to **suspend users for a bounded duration** so that proportional enforcement is possible.
    
*   As a **system**, I want to **record who suspended the user and why** so that accountability exists.
    
*   As a **suspended user**, I want to **see reason and expiry** so that I understand the decision.
    

#### Flow

*   **SuspendUserCommand(user\_id, reason, start\_at, end\_at, suspended\_by, notes?)** → ValidateReason() | ValidateDates() | CreateSuspension() | UpdateUserStatus() | NotifyUser() → **Outbox:** suspension.placed.v1
    
*   **UnsuspendUserCommand(user\_id, by, reason?)** → ValidateActiveSuspension() | DeactivateSuspension() | UpdateUserStatus() | NotifyUser() → **Outbox:** suspension.released.v1
    
*   **ExtendSuspensionCommand(suspension\_id, new\_end\_at, by, reason?)** → ValidateActive() | UpdateEnd() | NotifyUser() → **Outbox:** suspension.extended.v1
    
*   **AutoExpireSuspensionsJob()** → FindExpiredActive() | Deactivate() → **Outbox:** suspension.expired.v1
    
*   **GetActiveSuspensionQuery(user\_id)** → FetchActive() → SuspensionDTO
    
*   **GetSuspensionHistoryQuery(user\_id)** → ListByUser() → SuspensionListDTO
    

#### Projections

*   suspension\_read, active\_suspensions\_read, suspension\_history\_read
    

#### Events

*   suspension.placed.v1, suspension.released.v1, suspension.extended.v1, suspension.expired.v1
    

#### Events Consumed

*   user.severe\_violation.detected.v1, payment.fraud.confirmed.v1
    

#### RBAC/SLO

*   **RBAC:** ADMIN (place/extend/release), OWNER (view own), SYSTEM (auto-expire)
    
*   **SLO:** writes **P95 < 160 ms**, reads **P95 < 100 ms**
    

##### 11.2.1 Suspension Reasons (Enum)

*   **TOSViolation**, **PaymentIssue**, **QualityIssues**, **AbusiveBehavior**, **SpamReported**, **FakeProfile**, **MultipleAccounts**, **ContractAbandonment**, **ResponseFailure**, **AdminRequest**
    

**Flow:**

*   **GetSuspensionReasonsQuery()** → ListReasons()
    
*   **ValidateSuspensionReasonCommand(reason)** → CheckEnum() → is\_valid
    

**RBAC/SLO:** PUBLIC view, ADMIN use; **P95 < 50 ms**

##### 11.2.2 Suspension Duration

*   **Types:** Days (1–30), Weeks (1–12), Months (1–12), **Permanent**
    
*   **Escalation:** 1st=3–7d, 2nd=2–4w, 3rd=1–3m, 4th+=≥6m or permanent
    

**Flow:**

*   **CalculateEndDateCommand(duration\_type, value, start\_at)** → ComputeEnd()
    
*   **GetRecommendedDurationQuery(reason, offense\_count)** → ApplyEscalationRules()
    
*   **IsPermanentQuery(suspension\_id)** → CheckFlag()
    

**RBAC/SLO:** ADMIN set, SYSTEM calculate; **P95 < 50 ms**

### 11.3 Ban

#### Stories

*   As an **admin**, I want to **ban malicious users** (temporary or permanent) so that platform safety is protected.
    
*   As a **system**, I want to **revoke sessions/devices immediately** on ban so that access is cut off.
    

#### Flow

*   **BanUserCommand(user\_id, reason, is\_permanent, expires\_at?, banned\_by, notes?)** → ValidateReason() | ValidateSeverity() | CreateBan() | RevokeAllSessions() | RevokeAllDevices() | UpdateUserStatus() | NotifyUser() → **Outbox:** ban.placed.v1
    
*   **UnbanUserCommand(user\_id, by, reason?)** → ValidateActiveBan() | DeactivateBan() | UpdateUserStatus() | NotifyUser() → **Outbox:** ban.released.v1
    
*   **ConvertToPermanentBanCommand(ban\_id, by, reason?)** → ValidateTempBan() | MarkPermanent() | NotifyUser() → **Outbox:** ban.converted\_to\_permanent.v1
    
*   **AutoExpireTempBansJob()** → FindExpired() | Deactivate() → **Outbox:** ban.expired.v1
    
*   **GetActiveBanQuery(user\_id)** → FetchActive() → BanDTO
    
*   **GetBanHistoryQuery(user\_id)** → ListByUser() → BanListDTO
    

#### Projections

*   ban\_read, active\_bans\_read, ban\_history\_read
    

#### Events

*   ban.placed.v1, ban.released.v1, ban.converted\_to\_permanent.v1, ban.expired.v1
    

#### Events Consumed

*   fraud.detection.confirmed.v1, security.severe\_breach.detected.v1, moderation.repeated\_violations.threshold\_met.v1
    

#### RBAC/SLO

*   **RBAC:** ADMIN (ban/unban/convert), OWNER (view own), SYSTEM (auto-expire)
    
*   **SLO:** writes **P95 < 200 ms**, reads **P95 < 100 ms**
    

##### 11.3.1 Ban Reasons (Enum)

*   **Fraud**, **SevereAbuse**, **MultipleViolations**, **SecurityThreat**, **CriminalActivity**, **ChargebackAbuse**, **ImitationAccount**, **PlatformManipulation**, **DataBreach**, **CoordinatedAttack**
    

**Flow:**

*   **GetBanReasonsQuery()** → ListReasons()
    
*   **ValidateBanReasonCommand(reason)** → CheckEnum() | CheckSeverity() → is\_valid
    
*   **RequiresEvidenceQuery(reason)** → CheckPolicy() → boolean
    

**RBAC/SLO:** ADMIN use; SUPER\_ADMIN approve severe; **P95 < 50 ms**

##### 11.3.2 Permanent Ban Flag

**Flow:**

*   **MarkAsPermanentCommand(ban\_id)** → ValidateBan() | SetPermanent() | RemoveExpiry() → **Outbox:** ban.marked\_permanent.v1
    
*   **IsPermanentQuery(ban\_id)** → CheckFlag()
    
*   **OverridePermanentBanCommand(ban\_id, by, justification)** → AuthorizeSuperAdmin() | Approve() | Release() → **Outbox:** permanent\_ban.overridden.v1
    

**RBAC/SLO:** ADMIN mark; SUPER\_ADMIN override; **P95 < 100 ms**

### 11.4 Warning

#### Stories

*   As an **admin**, I want to **issue warnings** for minor violations so that users are nudged before penalties.
    
*   As a **user**, I want to **acknowledge warnings** so that receipt is confirmed.
    
*   As a **system**, I want **auto-escalation** for repeats or missed acknowledgements.
    

#### Flow

*   **IssueWarningCommand(user\_id, reason, severity, issued\_by, message?)** → ValidateReason() | CreateWarning() | NotifyUser() | SetAckDeadline() → **Outbox:** warning.issued.v1
    
*   **AcknowledgeWarningCommand(warning\_id, user\_id, at)** → ValidateOwner() | ValidatePending() | MarkAcknowledged() → **Outbox:** warning.acknowledged.v1
    
*   **EscalateWarningCommand(warning\_id, new\_severity, by, reason?)** → ValidateUnacknowledged() | UpdateSeverity() | NotifyUser() → **Outbox:** warning.escalated.v1
    
*   **DismissWarningCommand(warning\_id, by, reason?)** → AuthorizeAdmin() | Dismiss() → **Outbox:** warning.dismissed.v1
    
*   **AutoExpireUnackedWarningsJob()** → FindPastDeadline() | MarkExpired() | MaybeEscalate() → **Outbox:** warning.expired\_unacknowledged.v1
    
*   **GetActiveWarningsQuery(user\_id)** → ListUnacknowledged()
    
*   **GetWarningHistoryQuery(user\_id)** → ListByUser()
    

#### Projections

*   warning\_read, active\_warnings\_read, warning\_history\_read
    

#### Events

*   warning.issued.v1, warning.acknowledged.v1, warning.escalated.v1, warning.dismissed.v1, warning.expired\_unacknowledged.v1
    

#### Events Consumed

*   contract.quality\_issue.detected.v1, communication.unresponsive.threshold\_met.v1
    

#### RBAC/SLO

*   **RBAC:** ADMIN (issue/escalate/dismiss), OWNER (acknowledge/view)
    
*   **SLO:** writes **P95 < 140 ms**
    

##### 11.4.1 Warning Reasons (Enum)

*   **LateDelivery**, **PoorQuality**, **UnresponsiveCommunication**, **MissedDeadlines**, **InappropriateLanguage**, **MinorToSViolation**, **InaccurateProposal**, **PaymentDispute**, **ProfileMisrepresentation**, **IncompleteWork**
    

**Flow:**

*   **GetWarningReasonsQuery()** → ListReasons()
    
*   **ValidateWarningReasonCommand(reason)** → CheckEnum() → is\_valid
    
*   **GetHelpResourceQuery(reason)** → MapToGuidance() → ResourceDTO
    

**RBAC/SLO:** PUBLIC view, ADMIN use; **P95 < 50 ms**

##### 11.4.2 Warning Severity

*   **Levels:** Low, Medium, High, Critical
    
*   **Escalation:** 3×Low→Medium; 2×Medium→High; 2×High→Critical→Suspension; Critical+new violation→Ban
    

**Flow:**

*   **SetSeverityCommand(warning\_id, level)** → Validate() | Update() → **Outbox:** warning.severity\_updated.v1
    
*   **AutoEscalateSeverityCommand(user\_id)** → CheckRepeatedWarnings() | IncreaseSeverity() → **Outbox:** warning.auto\_escalated.v1
    
*   **ConvertToSuspensionCommand(warning\_id, by)** → ValidateCritical() | CreateSuspension() | CloseWarning() → **Outbox:** warning.converted\_to\_suspension.v1
    

**RBAC/SLO:** ADMIN set; SYSTEM auto-escalate/convert; **P95 < 100 ms**

### 11.5 Shared Moderation Components

#### 11.5.1 Shared Reason Model

*   **Goal:** common structures for reasons; DRY validation and severity mapping.
    

**Flow:**

*   **ValidateAnyReasonCommand(reason, action\_type)** → CheckEnum() | ValidateForAction()
    
*   **GetAllReasonsQuery()** → FetchAllGrouped()
    
*   **MapReasonToSeverityQuery(reason)** → PolicyMap()
    

**RBAC/SLO:** SYSTEM; **P95 < 50 ms**

#### 11.5.2 Moderation Actor

*   **Actor Types:** admin, system, automated, super\_admin, appeal\_reviewer
    
*   **Fields:** actor\_type, actor\_id, timestamp, reason
    

**Flow:**

*   **RecordActorCommand(action\_id, actor\_type, actor\_id)** → Validate() | Record() → **Outbox:** actor.recorded.v1
    
*   **GetActionActorQuery(action\_id)** → Fetch()
    
*   **GetActorHistoryQuery(actor\_id)** → ListActions()
    

**RBAC/SLO:** SYSTEM record; ADMIN view; **P95 < 50 ms**

### 11.6 Repository, Errors, and Events

#### 11.6.1 Repository Interface (consolidated)

*   **Suspension:** Create, GetActive, Release, ListHistory
    
*   **Ban:** Create, GetActive, Release, ListHistory
    
*   **Warning:** Create, Acknowledge, ListActive, ListHistory
    
*   **Aggregate:** GetAggregate, GetStats
    

**RBAC/SLO:** SYSTEM/ADMIN; writes **P95 < 120 ms**, reads **P95 < 80 ms**

#### 11.6.2 Domain Errors

*   **Suspension:** ErrSuspensionNotFound, ErrSuspensionAlreadyActive, ErrSuspensionExpired, ErrInvalidSuspensionDuration, ErrCannotUnsuspendBannedUser
    
*   **Ban:** ErrBanNotFound, ErrBanAlreadyActive, ErrCannotBanSuspendedUser, ErrPermanentBanCannotExpire, ErrBanRequiresSuperAdmin
    
*   **Warning:** ErrWarningNotFound, ErrWarningAlreadyAcknowledged, ErrWarningExpired, ErrCannotEscalateDismissedWarning, ErrMaxWarningsReached
    
*   **General:** ErrModerationActionNotFound, ErrInvalidModerationReason, ErrUnauthorizedModerationAction, ErrConflictingModerationActions
    

#### 11.6.3 Event Payloads (envelope-only PII rules)

*   **SuspensionPlacedEvent:** user\_id, suspension\_id, reason, start\_at, end\_at, suspended\_by
    
*   **BanPlacedEvent:** user\_id, ban\_id, reason, is\_permanent, expires\_at?, banned\_by
    
*   **WarningIssuedEvent:** user\_id, warning\_id, reason, severity, issued\_by, deadline\_at
    

**RBAC/SLO:** SYSTEM publish; **P95 < 50 ms**

### Operational Notes (applies to §11)

*   **Idempotency:** all commands accept Idempotency-Key; safe retries return original success; dedupe by (aggregate\_id,event\_type,idempotency\_key).
    
*   **Non-PII Events:** only IDs and codes; no raw emails/phones; fetch details via API if needed.
    
*   **Caching:** active bans/suspensions (TTL 5m); unacknowledged warning counts (TTL 1m); invalidate on write.
    
*   **Automation:** cron to expire time-boxed actions; auto-escalation from warnings → suspension → ban.
    
*   **Notifications:** email + in-app with appeal links on every action.


## **12 - SCORING DOMAINS (CONSOLIDATED)**

### 12.1 user_metrics/

#### User Stories
- As a **system**, I want to **store raw user metrics** (response time, completion rate, client satisfaction) so that all scoring contexts have a single source of truth.
- As a **system**, I want to **record metric updates** as events occur so that metrics stay current.
- As a **system**, I want to **aggregate metrics periodically** so that derived scores are computed efficiently.
- As a **system**, I want to **provide metric history** so that trends can be analyzed.

#### Flow
1. **RecordMetricCommand**(user_id, metric_name, metric_value, timestamp) → Validate() | Persist() → **Outbox:** metric.recorded.v1
2. **UpdateMetricsCommand**(user_id, metrics_batch) → ValidateBatch() | BulkUpdate() → **Outbox:** metrics.updated.v1
3. **AggregateMetricsCommand**(user_id, aggregation_period) → Compute() | Store() → **Outbox:** metrics.aggregated.v1
4. **GetMetricsQuery**(user_id) → AuthorizeInternal() | Fetch() → UserMetricsDTO
5. **GetMetricHistoryQuery**(user_id, metric_name, time_range) → Fetch() → MetricHistoryDTO

#### Projections
- user_metrics_read
- metric_history_read

#### Events Published
- metric.recorded.v1
- metrics.updated.v1
- metrics.aggregated.v1

#### Events Consumed
- contract.completed.v1 (to update completion rate)
- message.sent.v1 (to update response time)
- review.submitted.v1 (to update satisfaction)

#### RBAC/SLO
- **RBAC:** SYSTEM (record/update), INTERNAL (view raw metrics)
- **SLO:** P95 < 120ms

---

### 12.2 reputation/

#### User Stories
- As a **freelancer**, I want to **have a reputation score calculated** so that clients can assess my reliability.
- As a **system**, I want to **compute reputation from user_metrics** (reviews, completion, response, quality) so that scoring is standardized.
- As a **freelancer**, I want to **view my reputation components** so that I understand what affects my score.
- As a **system**, I want to **recalculate reputation periodically** so that scores reflect recent activity.
- As a **client**, I want to **filter freelancers by reputation** so that I find reliable talent.

#### Flow
1. **RecalculateReputationCommand**(user_id) → FetchMetrics() | ComputeScore() | Update() → **Outbox:** reputation.updated.v1
2. **RecordReputationEventCommand**(user_id, event_type, impact) → Apply() | Trigger Recalc() → **Outbox:** reputation.score_changed.v1
3. **GetReputationScoreQuery**(user_id) → Fetch() → ReputationDTO
4. **GetReputationComponentsQuery**(user_id) → Fetch() → ReputationComponentsDTO
5. **GetReputationHistoryQuery**(user_id, time_range) → Fetch() → ReputationHistoryDTO

#### Projections
- reputation_read
- reputation_history_read

#### Events Published
- reputation.updated.v1
- reputation.score_changed.v1

#### Events Consumed
- metrics.updated.v1 (trigger recalculation)
- review.submitted.v1 (trigger recalculation)

#### RBAC/SLO
- **RBAC:** SYSTEM (recalculate), PUBLIC (view), OWNER (view components/history)
- **SLO:** P95 < 150ms

---

### 12.3 quality/

#### User Stories
- As a **freelancer**, I want to **have a quality score calculated** so that clients see my work quality.
- As a **system**, I want to **compute quality from user_metrics** (completion rate, response time, satisfaction, work quality) so that quality is measurable.
- As a **freelancer**, I want to **view my quality scoring factors** so that I know what to improve.
- As a **system**, I want to **analyze quality trends** (improving, declining, stable) so that issues are detected early.

#### Flow
1. **RecalculateQualityScoreCommand**(user_id) → FetchMetrics() | ComputeScore() | Update() → **Outbox:** quality_score.updated.v1
2. **RecordQualityMetricCommand**(user_id, metric_type, value) → Validate() | Apply() → **Outbox:** quality_metric.recorded.v1
3. **GetQualityScoreQuery**(user_id) → Fetch() → QualityScoreDTO
4. **GetScoringFactorsQuery**(user_id) → Fetch() → ScoringFactorsDTO
5. **GetScoreTrendQuery**(user_id, time_range) → AnalyzeTrend() → TrendAnalysisDTO

#### Projections
- quality_read
- quality_trend_read

#### Events Published
- quality_score.updated.v1
- quality_metric.recorded.v1
- quality.improved.v1
- quality.declined.v1

#### Events Consumed
- metrics.updated.v1 (trigger recalculation)

#### RBAC/SLO
- **RBAC:** SYSTEM (recalculate), PUBLIC (view score), OWNER (view factors/trends)
- **SLO:** P95 < 150ms

---

### 12.4 account_health/

#### User Stories
- As a **freelancer**, I want to **have an account health score** so that I know my overall account status.
- As a **system**, I want to **compute health from user_metrics** (profile completeness, activity, responsiveness, quality) so that health is measurable.
- As a **freelancer**, I want to **see detected health issues** so that I can address them.
- As a **system**, I want to **generate health recommendations** so that users can improve their accounts.

#### Flow
1. **RecalculateHealthScoreCommand**(user_id) → FetchMetrics() | ComputeScore() | DetectIssues() | Update() → **Outbox:** health_score.updated.v1
2. **RecordHealthIssueCommand**(user_id, issue_type, severity) → Validate() | Record() → **Outbox:** health_issue.detected.v1
3. **GetAccountHealthQuery**(user_id) → Fetch() → AccountHealthDTO
4. **GetHealthIssuesQuery**(user_id) → Fetch() → HealthIssueListDTO
5. **GetHealthRecommendationsQuery**(user_id) → Generate() → RecommendationListDTO

#### Projections
- account_health_read
- health_issues_read

#### Events Published
- health_score.updated.v1
- health_issue.detected.v1
- health.improved.v1

#### Events Consumed
- metrics.updated.v1 (trigger recalculation)
- profile_completeness.updated.v1 (trigger recalculation)

#### RBAC/SLO
- **RBAC:** SYSTEM (recalculate), OWNER (view)
- **SLO:** P95 < 160ms

---

### 12.5 risk/

#### User Stories
- As a **system**, I want to **compute risk scores from user_metrics and signals** so that fraud/safety risks are detected.
- As a **system**, I want to **record risk signals** (ip_geo_mismatch, disputes, chargebacks) so that patterns are tracked.
- As an **admin**, I want to **place account holds** so that risky accounts are restricted.
- As an **admin**, I want to **release account holds** so that false positives can be corrected.
- As a **system**, I want to **update risk scores when signals change** so that risk assessment is current.

#### Flow
1. **RecordRiskSignalCommand**(user_id, signal_type, severity, occurred_at, metadata) → Validate() | Record() | TriggerRecalc() → **Outbox:** risk_signal.recorded.v1
2. **RecalculateRiskScoreCommand**(user_id) → FetchMetrics() | FetchSignals() | ComputeScore() | Update() → **Outbox:** risk_score.updated.v1
3. **ApplyAccountHoldCommand**(user_id, hold_type, reason, actor, until) → Validate() | Apply() → **Outbox:** risk_hold.placed.v1
4. **ReleaseAccountHoldCommand**(user_id, released_by, reason) → Validate() | Release() → **Outbox:** risk_hold.released.v1
5. **GetRiskScoreQuery**(user_id) → AuthorizeAdmin() | Fetch() → RiskScoreDTO
6. **ListRiskSignalsQuery**(user_id) → AuthorizeAdmin() | Fetch() → RiskSignalListDTO
7. **GetAccountStateQuery**(user_id) → AuthorizeAdmin() | Fetch() → AccountStateDTO

#### Projections
- risk_read
- risk_signals_read
- account_holds_read

#### Events Published
- risk_signal.recorded.v1
- risk_score.updated.v1
- risk_hold.placed.v1
- risk_hold.released.v1

#### Events Consumed
- payment.chargeback.v1 (record signal)
- dispute.created.v1 (record signal)
- login.ip_geo_mismatch.v1 (record signal)
- financial.risk.alert.emitted.v1 (record signal)

#### RBAC/SLO
- **RBAC:** SYSTEM (record/recalculate), ADMIN (view/hold/release)
- **SLO:** P95 < 170ms

---

## **13 - ORGANIZATION & TEAM DOMAIN**

### 13.1 org/

#### User Stories
- As a **client**, I want to **create an organization** so that multiple team members can hire on behalf of the company.
- As an **org owner**, I want to **invite team members** so that they can collaborate on hiring.
- As an **org admin**, I want to **assign roles** (owner, admin, member) so that permissions are managed.
- As an **org owner**, I want to **set seat limits** so that team size is controlled.
- As a **system**, I want to **track seat usage** so that limits are enforced.
- As a **client service**, I want to **reference org data** so that there's no duplication of company information.

#### Flow
1. **CreateOrgCommand**(name, industry, size, founded, employees, created_by) → Validate() | Persist() | AssignOwner() → **Outbox:** org.created.v1
2. **UpdateOrgCommand**(org_id, updates) → AuthorizeAdmin() | Validate() | Update() → **Outbox:** org.updated.v1
3. **InviteMemberCommand**(org_id, user_id, role, invited_by) → CheckSeats() | ValidateRole() | Invite() → **Outbox:** org.member_invited.v1
4. **RemoveMemberCommand**(org_id, user_id, removed_by) → AuthorizeAdmin() | Remove() → **Outbox:** org.member_removed.v1
5. **AssignRoleCommand**(org_id, user_id, role, assigned_by) → AuthorizeAdmin() | Update() → **Outbox:** org.role_assigned.v1
6. **SetSeatCountCommand**(org_id, seat_limit) → AuthorizeOwner() | ValidateUsage() | Update() → **Outbox:** org.seats_updated.v1
7. **GetOrgQuery**(org_id) → AuthorizeMember() | Fetch() → OrgDTO
8. **ListOrgsForUserQuery**(user_id) → Fetch() → OrgListDTO
9. **ListMembersQuery**(org_id) → AuthorizeMember() | Fetch() → MemberListDTO
10. **GetSeatUsageQuery**(org_id) → AuthorizeAdmin() | Calculate() → SeatUsageDTO

#### Projections
- org_read
- org_members_read
- org_seats_read

#### Events Published
- org.created.v1
- org.updated.v1
- org.member_invited.v1
- org.member_added.v1
- org.member_removed.v1
- org.role_assigned.v1
- org.seats_updated.v1

#### RBAC/SLO
- **RBAC:** OWNER (create/update/seats), ADMIN (invite/remove/assign roles), MEMBER (view)
- **SLO:** P95 < 170ms

---

## **14 - SECURITY DOMAIN (CONSOLIDATED)**

### 14.1 security/

#### User Stories
- As a **user**, I want to **enable 2FA** (TOTP, SMS, Email) so that my account is more secure.
- As a **user**, I want to **generate backup codes** so that I can recover access if I lose my 2FA device.
- As a **user**, I want to **register trusted devices** so that I don't need 2FA on familiar devices.
- As a **user**, I want to **view active sessions** so that I can see where I'm logged in.
- As a **user**, I want to **revoke sessions** so that I can log out remotely.
- As a **user**, I want to **revoke devices** so that lost devices can't access my account.
- As a **user**, I want to **initiate account recovery** if I lose access so that I can regain control.
- As a **system**, I want to **rate limit recovery attempts** so that brute force attacks are prevented.

#### Flow
1. **Enable2FACommand**(user_id, method, secret) → GenerateSecret() | StoreSecret() | GenerateBackupCodes() → **Outbox:** two_fa.enabled.v1
2. **Disable2FACommand**(user_id, verification_code) → ValidateCode() | Disable() → **Outbox:** two_fa.disabled.v1
3. **RegisterDeviceCommand**(user_id, device_fingerprint, device_info) → ValidateFingerprint() | Register() → **Outbox:** device.registered.v1
4. **RevokeDeviceCommand**(user_id, device_id) → AuthorizeOwner() | Revoke() → **Outbox:** device.revoked.v1
5. **RevokeSessionCommand**(user_id, session_id) → AuthorizeOwner() | Revoke() → **Outbox:** session.revoked.v1
6. **InitiateRecoveryCommand**(email, recovery_method) → ValidateAccount() | RateLimitCheck() | SendToken() → **Outbox:** recovery.initiated.v1
7. **CompleteRecoveryCommand**(recovery_token, new_password) → ValidateToken() | ResetPassword() → **Outbox:** recovery.completed.v1
8. **Get2FAStatusQuery**(user_id) → AuthorizeOwner() | Fetch() → TwoFAStatusDTO
9. **ListDevicesQuery**(user_id) → AuthorizeOwner() | Fetch() → DeviceListDTO
10. **ListSessionsQuery**(user_id) → AuthorizeOwner() | Fetch() → SessionListDTO
11. **GetSecuritySettingsQuery**(user_id) → AuthorizeOwner() | Fetch() → SecuritySettingsDTO

#### Projections
- security_read
- devices_read
- sessions_read

#### Events Published
- two_fa.enabled.v1
- two_fa.disabled.v1
- device.registered.v1
- device.revoked.v1
- session.revoked.v1
- recovery.initiated.v1
- recovery.completed.v1

#### RBAC/SLO
- **RBAC:** OWNER (all security operations)
- **SLO:** P95 < 160ms (non-recovery), P95 < 300ms (recovery)

---

## **15 - PROFILE ENHANCEMENT DOMAINS (CONSOLIDATED)**

### 15.1 profile_depth/

#### User Stories
- As a **system**, I want to **track rate history** so that rate changes over time are visible.
- As a **freelancer**, I want to **view my rate history** so that I can track pricing evolution.
- As a **system**, I want to **map skills to normalized taxonomy** so that skills are standardized across the platform.
- As a **system**, I want to **store taxonomy mapping** so that search and matching work effectively.

#### Flow
1. **AddHourlyRateEntryCommand**(user_id, amount, currency, effective_at) → Validate() | Append() → **Outbox:** rate_history.updated.v1
2. **NormalizeSkillSetCommand**(user_id, skills) → MapToTaxonomy() | Update() → **Outbox:** taxonomy.updated.v1
3. **GetRateHistoryQuery**(user_id) → Fetch() → RateHistoryDTO
4. **GetNormalizedSkillsQuery**(user_id) → Fetch() → NormalizedSkillsDTO

#### Projections
- profile_depth_read
- rate_history_read
- taxonomy_mapping_read

#### Events Published
- rate_history.updated.v1
- taxonomy.updated.v1

#### RBAC/SLO
- **RBAC:** SYSTEM (normalize), OWNER (add rate), PUBLIC (view)
- **SLO:** P95 < 140ms

---

### 15.2 profile_completeness/

#### User Stories
- As a **freelancer**, I want to **see my profile completeness score** so that I know how complete my profile is.
- As a **freelancer**, I want to **see missing sections** so that I know what to add.
- As a **system**, I want to **calculate completeness based on section weights** so that important sections are prioritized.
- As a **system**, I want to **provide recommendations** so that users know how to improve.
- As a **freelancer**, I want to **reach completeness milestones** so that I'm encouraged to complete my profile.

#### Flow
1. **RecalculateCompletenessCommand**(user_id) → FetchProfileSections() | CalculateScore() | IdentifyMissing() | GenerateRecommendations() → **Outbox:** completeness.updated.v1
2. **MarkSectionCompleteCommand**(user_id, section_name) → Validate() | Update() | TriggerRecalc() → **Outbox:** section.completed.v1
3. **GetCompletenessScoreQuery**(user_id) → Fetch() → ProfileCompletenessDTO
4. **GetMissingSectionsQuery**(user_id) → Fetch() → MissingSectionsDTO
5. **GetRecommendationsQuery**(user_id) → Fetch() → RecommendationListDTO

#### Projections
- profile_completeness_read

#### Events Published
- completeness.updated.v1
- section.completed.v1
- milestone.reached.v1

#### Events Consumed
- profile.updated.v1 (trigger recalculation)
- skill.added.v1 (trigger recalculation)
- experience.added.v1 (trigger recalculation)

#### RBAC/SLO
- **RBAC:** SYSTEM (recalculate), OWNER (view)
- **SLO:** P95 < 150ms

---

### 15.3 profile_analytics/

#### User Stories
- As a **freelancer**, I want to **track profile views** so that I know how much exposure I'm getting.
- As a **freelancer**, I want to **see view sources** (search, direct link, referral) so that I understand where traffic comes from.
- As a **freelancer**, I want to **track search appearances** so that I know how visible I am.
- As a **system**, I want to **record engagement metrics** (click-through rates, interest signals) so that profile performance is measurable.

#### Flow
1. **RecordViewCommand**(profile_id, viewer_id, source, viewed_at) → Validate() | Record() → **Outbox:** profile.viewed.v1
2. **RecordSearchImpressionCommand**(profile_id, search_query, rank, appeared_at) → Validate() | Record() → **Outbox:** search.appearance.v1
3. **RecordEngagementCommand**(profile_id, engagement_type, metadata) → Validate() | Record() → **Outbox:** engagement.recorded.v1
4. **GetProfileViewsQuery**(user_id, time_range) → AuthorizeOwner() | Fetch() → ProfileViewsDTO
5. **GetSearchAnalyticsQuery**(user_id, time_range) → AuthorizeOwner() | Fetch() → SearchAnalyticsDTO
6. **GetEngagementMetricsQuery**(user_id, time_range) → AuthorizeOwner() | Fetch() → EngagementMetricsDTO

#### Projections
- profile_analytics_read
- profile_views_read
- search_analytics_read

#### Events Published
- profile.viewed.v1
- search.appearance.v1
- engagement.recorded.v1

#### RBAC/SLO
- **RBAC:** SYSTEM (record), OWNER (view)
- **SLO:** P95 < 140ms

---

### 15.4 profile_optimization/

#### User Stories
- As a **freelancer**, I want to **receive AI-powered profile suggestions** so that I can improve my profile quality.
- As a **freelancer**, I want to **get keyword optimization recommendations** so that I appear in more searches.
- As a **freelancer**, I want to **receive headline suggestions** so that my profile is more attractive.
- As a **system**, I want to **track applied suggestions** so that effectiveness can be measured.

#### Flow
1. **GenerateAISuggestionsCommand**(user_id) → AnalyzeProfile() | GenerateSuggestions() | Store() → **Outbox:** suggestions.generated.v1
2. **ApplyOptimizationCommand**(user_id, suggestion_id) → FetchSuggestion() | ApplyToProfile() | MarkApplied() → **Outbox:** optimization.applied.v1
3. **GetOptimizationsQuery**(user_id) → AuthorizeOwner() | Fetch() → ProfileOptimizationDTO
4. **GetKeywordSuggestionsQuery**(user_id) → Generate() → KeywordSuggestionsDTO

#### Projections
- profile_optimization_read

#### Events Published
- suggestions.generated.v1
- optimization.applied.v1
- optimization.completed.v1

#### RBAC/SLO
- **RBAC:** SYSTEM (generate), OWNER (apply/view)
- **SLO:** P95 < 250ms (generate), P95 < 180ms (apply)

---

## **16 - PROFILE VISIBILITY DOMAIN (CONSOLIDATED)**

### 16.1 profile_visibility/

#### User Stories
- As a **freelancer**, I want to **set my visibility level** (Public, LimitedPublic, Private, AnonymousMode) so that I control who sees my profile.
- As a **freelancer**, I want to **control searchable categories** so that I appear in relevant searches only.
- As a **freelancer**, I want to **hide my profile from specific users** so that I manage visibility granularly.
- As a **freelancer**, I want to **enable stealth mode** so that I can browse profiles anonymously.
- As a **system**, I want to **enforce visibility rules in search** so that privacy preferences are respected.

#### Flow
1. **ChangeVisibilityLevelCommand**(user_id, visibility_level) → Validate() | Update() → **Outbox:** visibility.changed.v1
2. **UpdateSearchableCategoriesCommand**(user_id, categories) → Validate() | Update() → **Outbox:** searchable_categories.updated.v1
3. **ToggleStealthModeCommand**(user_id, enabled) → Update() → **Outbox:** stealth_mode.toggled.v1
4. **GetVisibilitySettingsQuery**(user_id) → AuthorizeOwner() | Fetch() → ProfileVisibilityDTO
5. **GetSearchPreferencesQuery**(user_id) → AuthorizeOwner() | Fetch() → SearchPreferencesDTO

#### Projections
- profile_visibility_read

#### Events Published
- visibility.changed.v1
- searchable_categories.updated.v1
- stealth_mode.toggled.v1

#### RBAC/SLO
- **RBAC:** OWNER (update/view)
- **SLO:** P95 < 140ms

---

## **17 - AVAILABILITY DOMAIN (CONSOLIDATED)**

### 17.1 availability/

#### User Stories
- As a **freelancer**, I want to **set my availability status** (Available, Busy, Away, DoNotDisturb) so that clients know when I can work.
- As a **freelancer**, I want to **create recurring availability schedules** (weekly/monthly) so that my availability is predictable.
- As a **freelancer**, I want to **enable vacation mode** so that I can take time off without losing visibility.
- As a **freelancer**, I want to **sync with external calendars** (Google Calendar, Outlook) so that availability is automatic.
- As a **system**, I want to **provide single source of availability** so that profile and other services reference consistent data.

#### Flow
1. **SetAvailabilityStatusCommand**(user_id, status) → Validate() | Update() → **Outbox:** availability.updated.v1
2. **CreateRecurringScheduleCommand**(user_id, schedule_rules) → Validate() | Store() → **Outbox:** recurring_schedule.created.v1
3. **ToggleVacationModeCommand**(user_id, enabled, start_date, end_date) → Validate() | Update() → **Outbox:** vacation_mode.toggled.v1
4. **SyncExternalCalendarCommand**(user_id, calendar_provider, auth_token) → AuthorizeProvider() | Sync() → **Outbox:** calendar.synced.v1
5. **GetAvailabilityQuery**(user_id) → Fetch() → AvailabilityDTO
6. **GetRecurringScheduleQuery**(user_id) → Fetch() → RecurringScheduleDTO

#### Projections
- availability_read

#### Events Published
- availability.updated.v1
- recurring_schedule.created.v1
- recurring_schedule.updated.v1
- vacation_mode.toggled.v1
- calendar.synced.v1

#### RBAC/SLO
- **RBAC:** OWNER (update), PUBLIC (view status)
- **SLO:** P95 < 150ms

---

### 17.2 workload_capacity/

#### User Stories
- As a **freelancer**, I want to **track my current workload** so that I don't overcommit.
- As a **freelancer**, I want to **set maximum capacity** so that the platform prevents overbooking.
- As a **freelancer**, I want to **see available hours** so that I know how much work I can take.
- As a **system**, I want to **calculate capacity automatically** so that availability is accurate.
- As a **system**, I want to **prevent overcommitment** so that freelancers maintain quality.

#### Flow
1. **UpdateCurrentLoadCommand**(user_id, current_hours, commitments) → Calculate() | Update() → **Outbox:** workload.updated.v1
2. **AddCommitmentCommand**(user_id, commitment_hours, start_date, end_date) → ValidateCapacity() | Add() → **Outbox:** commitment.added.v1
3. **RemoveCommitmentCommand**(user_id, commitment_id) → Remove() | RecalculateCapacity() → **Outbox:** commitment.removed.v1
4. **SetMaxCapacityCommand**(user_id, max_hours_per_week) → Validate() | Update() → **Outbox:** max_capacity.updated.v1
5. **GetCapacityQuery**(user_id) → Calculate() → WorkloadCapacityDTO
6. **GetAvailableHoursQuery**(user_id, time_range) → Calculate() → AvailableHoursDTO

#### Projections
- workload_capacity_read

#### Events Published
- workload.updated.v1
- commitment.added.v1
- commitment.removed.v1
- max_capacity.updated.v1
- capacity.full.v1
- capacity.available.v1

#### Events Consumed
- contract.created.v1 (add commitment)
- contract.completed.v1 (remove commitment)

#### RBAC/SLO
- **RBAC:** OWNER (update), SYSTEM (auto-calculate), PUBLIC (view available hours)
- **SLO:** P95 < 160ms

---

## **18 - NETWORKING & CONNECTIONS DOMAINS**

### 18.1 professional_network/

#### User Stories
- As a **user**, I want to **send connection requests** so that I can build my professional network.
- As a **user**, I want to **accept or decline connection requests** so that I control my network.
- As a **user**, I want to **specify relationship types** (Colleague, Client, Peer, Mentor) so that connections are categorized.
- As a **user**, I want to **remove connections** so that I can manage my network.
- As a **user**, I want to **view network analytics** (size, growth, strength) so that I understand my network.

#### Flow
1. **CreateConnectionRequestCommand**(requester_id, target_id, relationship_type, message) → ValidateNoDuplicate() | Send() → **Outbox:** connection.requested.v1
2. **AcceptRequestCommand**(request_id, acceptor_id) → Validate() | CreateConnection() → **Outbox:** connection.accepted.v1
3. **DeclineRequestCommand**(request_id, decliner_id) → Validate() | Decline() → **Outbox:** connection.declined.v1
4. **RemoveConnectionCommand**(user_id, connection_id) → Validate() | Remove() → **Outbox:** connection.removed.v1
5. **GetConnectionsQuery**(user_id) → Fetch() → ConnectionListDTO
6. **GetConnectionRequestsQuery**(user_id) → Fetch() → ConnectionRequestListDTO
7. **GetNetworkAnalyticsQuery**(user_id) → Calculate() → NetworkAnalyticsDTO

#### Projections
- professional_network_read
- connection_requests_read
- network_analytics_read

#### Events Published
- connection.requested.v1
- connection.accepted.v1
- connection.declined.v1
- connection.removed.v1

#### RBAC/SLO
- **RBAC:** OWNER (send/accept/decline/remove), PUBLIC (view connections)
- **SLO:** P95 < 160ms

---

### 18.2 referrals/

#### User Stories
- As a **user**, I want to **generate a referral code** so that I can invite others to the platform.
- As a **user**, I want to **track referral clicks** so that I know who's using my code.
- As a **system**, I want to **mark referrals as converted** when referred users sign up so that rewards can be issued.
- As a **user**, I want to **receive referral rewards** so that I'm incentivized to invite others.
- As a **user**, I want to **view referral statistics** so that I can track my referrals.

#### Flow
1. **CreateReferralCodeCommand**(user_id) → GenerateUniqueCode() | Persist() → **Outbox:** referral_code.created.v1
2. **RecordReferralClickCommand**(referral_code, clicked_at, metadata) → Validate() | Track() → **Outbox:** referral.clicked.v1
3. **MarkReferralConvertedCommand**(referral_code, referred_user_id) → Validate() | Convert() | CalculateReward() → **Outbox:** referral.converted.v1
4. **IssueRewardCommand**(user_id, reward_amount, reward_type) → ValidateEligibility() | Issue() → **Outbox:** reward.earned.v1
5. **GetReferralCodeQuery**(user_id) → Fetch() → ReferralCodeDTO
6. **GetReferralStatsQuery**(user_id) → Fetch() → ReferralStatsDTO
7. **GetReferralRewardsQuery**(user_id) → Fetch() → ReferralRewardListDTO

#### Projections
- referrals_read
- referral_stats_read

#### Events Published
- referral_code.created.v1
- referral.clicked.v1
- referral.converted.v1
- reward.earned.v1

#### Events Consumed
- user.created.v1 (to mark conversion)

#### RBAC/SLO
- **RBAC:** OWNER (create code/view stats), SYSTEM (track/convert/reward)
- **SLO:** P95 < 150ms

---

### 18.3 user_groups/

#### User Stories
- As a **user**, I want to **create community groups** (by skill, location, industry) so that I can connect with like-minded people.
- As a **user**, I want to **join groups** so that I can participate in communities.
- As a **user**, I want to **leave groups** so that I can manage my memberships.
- As a **group owner**, I want to **assign moderators** so that groups are well-managed.
- As a **system**, I want to **track group activity** so that engagement is measurable.
- As a **system**, I want to **enforce member limits** so that groups remain manageable.

#### Flow
1. **CreateUserGroupCommand**(creator_id, name, category, description, member_limit) → Validate() | Create() | AssignOwner() → **Outbox:** group.created.v1
2. **AddMemberCommand**(group_id, user_id) → CheckMemberLimit() | Add() → **Outbox:** group.member_added.v1
3. **RemoveMemberCommand**(group_id, user_id, removed_by) → AuthorizeModerator() | Remove() → **Outbox:** group.member_removed.v1
4. **AssignModeratorCommand**(group_id, user_id, assigned_by) → AuthorizeOwner() | Assign() → **Outbox:** group.moderator_assigned.v1
5. **GetGroupQuery**(group_id) → Fetch() → UserGroupDTO
6. **ListUserGroupsQuery**(user_id) → Fetch() → UserGroupListDTO
7. **GetGroupMembersQuery**(group_id) → Fetch() → GroupMemberListDTO

#### Projections
- user_groups_read
- group_members_read
- group_activity_read

#### Events Published
- group.created.v1
- group.member_added.v1
- group.member_removed.v1
- group.moderator_assigned.v1

#### RBAC/SLO
- **RBAC:** PUBLIC (create/join), OWNER (assign moderators), MODERATOR (remove members), PUBLIC (view)
- **SLO:** P95 < 170ms

---

## **19 - FINANCIAL PROFILE DOMAINS**

### 19.1 payment_methods/

#### User Stories
- As a **user**, I want to **add payment methods** (bank account, card, PayPal, crypto, Wise) so that I can receive payments.
- As a **user**, I want to **verify payment methods** so that they're activated.
- As a **user**, I want to **set a default payment method** so that withdrawals are automatic.
- As a **user**, I want to **delete payment methods** so that I can remove outdated methods.
- As a **system**, I want to **validate payment method details** so that transactions succeed.

#### Flow
1. **CreatePaymentMethodCommand**(user_id, method_type, details, is_default) → ValidateDetails() | Encrypt() | Persist() → **Outbox:** payment_method.added.v1
2. **VerifyMethodCommand**(method_id, verification_code) → ValidateCode() | Activate() → **Outbox:** payment_method.verified.v1
3. **SetAsDefaultCommand**(user_id, method_id) → Validate() | UpdateDefault() → **Outbox:** default_method.changed.v1
4. **DeletePaymentMethodCommand**(user_id, method_id) → ValidateNotDefault() | Delete() → **Outbox:** payment_method.deleted.v1
5. **GetPaymentMethodsQuery**(user_id) → AuthorizeOwner() | Fetch() → PaymentMethodListDTO
6. **GetDefaultMethodQuery**(user_id) → AuthorizeOwner() | Fetch() → PaymentMethodDTO

#### Projections
- payment_methods_read

#### Events Published
- payment_method.added.v1
- payment_method.verified.v1
- default_method.changed.v1
- payment_method.deleted.v1

#### RBAC/SLO
- **RBAC:** OWNER (add/verify/set default/delete/view)
- **SLO:** P95 < 180ms

---

### 19.2 financial_profile/

#### User Stories
- As a **freelancer**, I want to **set my preferred currency** so that I'm paid in my currency.
- As a **freelancer**, I want to **customize invoice settings** (logo, footer, terms) so that invoices are branded.
- As a **freelancer**, I want to **set default payment terms** (NET 30, upfront, milestones) so that clients know expectations.
- As a **system**, I want to **validate currency codes** so that only supported currencies are used.

#### Flow
1. **UpdateCurrencyPreferencesCommand**(user_id, preferred_currencies) → ValidateCurrencyCodes() | Update() → **Outbox:** currency_preference.updated.v1
2. **UpdateInvoiceSettingsCommand**(user_id, logo_url, footer_text, custom_terms) → Validate() | Update() → **Outbox:** invoice_settings.updated.v1
3. **SetDefaultPaymentTermsCommand**(user_id, payment_terms) → Validate() | Update() → **Outbox:** payment_terms.updated.v1
4. **GetFinancialProfileQuery**(user_id) → AuthorizeOwner() | Fetch() → FinancialProfileDTO
5. **GetCurrencyPreferencesQuery**(user_id) → Fetch() → CurrencyPreferencesDTO
6. **GetInvoiceSettingsQuery**(user_id) → AuthorizeOwner() | Fetch() → InvoiceSettingsDTO

#### Projections
- financial_profile_read

#### Events Published
- currency_preference.updated.v1
- invoice_settings.updated.v1
- payment_terms.updated.v1

#### RBAC/SLO
- **RBAC:** OWNER (update/view)
- **SLO:** P95 < 150ms

---

### 19.3 earning_goals/

#### User Stories
- As a **freelancer**, I want to **set earning goals** (monthly, quarterly, annual) so that I can track my progress.
- As a **freelancer**, I want to **track progress toward goals** so that I know how I'm doing.
- As a **system**, I want to **notify users when goals are achieved** so that accomplishments are celebrated.
- As a **freelancer**, I want to **view goal achievements** so that I can see my history.

#### Flow
1. **CreateEarningGoalCommand**(user_id, target_amount, period, goal_type) → Validate() | Persist() → **Outbox:** goal.created.v1
2. **UpdateGoalProgressCommand**(goal_id, current_amount) → Calculate() | Update() → **Outbox:** goal_progress.updated.v1
3. **MarkGoalAchievedCommand**(goal_id) → Validate() | MarkAchieved() → **Outbox:** goal.achieved.v1
4. **GetEarningGoalsQuery**(user_id) → Fetch() → EarningGoalListDTO
5. **GetGoalProgressQuery**(goal_id) → Calculate() → GoalProgressDTO
6. **GetAchievementsQuery**(user_id) → Fetch() → GoalAchievementListDTO

#### Projections
- earning_goals_read
- goal_progress_read

#### Events Published
- goal.created.v1
- goal_progress.updated.v1
- goal.achieved.v1

#### Events Consumed
- payment.received.v1 (to update progress)

#### RBAC/SLO
- **RBAC:** OWNER (create/view), SYSTEM (update progress)
- **SLO:** P95 < 150ms

---

## **20 - PROFESSIONAL DEVELOPMENT DOMAINS**

### 20.1 learning_path/

#### User Stories
- As a **freelancer**, I want to **receive personalized learning paths** so that I can develop relevant skills.
- As a **freelancer**, I want to **see my skill gaps** so that I know what to learn.
- As a **system**, I want to **recommend courses** based on skill gaps so that learning is targeted.
- As a **freelancer**, I want to **enroll in courses** so that I can start learning.
- As a **freelancer**, I want to **track course completion** so that I see my progress.

#### Flow
1. **GenerateLearningPathCommand**(user_id) → AnalyzeSkills() | IdentifyGaps() | RecommendCourses() | Create() → **Outbox:** learning_path.created.v1
2. **EnrollInCourseCommand**(user_id, course_id) → Validate() | Enroll() → **Outbox:** course.enrolled.v1
3. **CompleteCourseCommand**(user_id, course_id, completion_data) → Validate() | MarkComplete() → **Outbox:** course.completed.v1
4. **GetLearningPathQuery**(user_id) → Fetch() → LearningPathDTO
5. **GetSkillGapsQuery**(user_id) → Analyze() → SkillGapListDTO
6. **GetCourseRecommendationsQuery**(user_id) → Generate() → CourseRecommendationListDTO

#### Projections
- learning_path_read
- skill_gaps_read

#### Events Published
- learning_path.created.v1
- course.enrolled.v1
- course.completed.v1
- skill.acquired.v1

#### RBAC/SLO
- **RBAC:** SYSTEM (generate), OWNER (enroll/view)
- **SLO:** P95 < 200ms

---

### 20.2 mentorship/

#### User Stories
- As a **freelancer**, I want to **request mentorship** so that I can learn from experienced professionals.
- As a **mentor**, I want to **set my availability and expertise** so that mentees can find me.
- As a **system**, I want to **match mentors with mentees** based on skills and goals so that pairings are effective.
- As a **user**, I want to **schedule mentorship sessions** so that meetings are organized.
- As a **system**, I want to **track mentorship progress** so that outcomes are measurable.

#### Flow
1. **CreateMentorshipRequestCommand**(mentee_id, desired_skills, goals) → Validate() | Create() → **Outbox:** mentorship.requested.v1
2. **AcceptMentorshipCommand**(mentor_id, request_id) → Validate() | CreatePairing() → **Outbox:** mentorship.accepted.v1
3. **ScheduleMentorshipSessionCommand**(mentorship_id, scheduled_at, duration) → ValidateAvailability() | Schedule() → **Outbox:** session.scheduled.v1
4. **CompleteMentorshipCommand**(mentorship_id, completion_notes) → Validate() | Complete() → **Outbox:** mentorship.completed.v1
5. **GetMentorshipQuery**(mentorship_id) → Fetch() → MentorshipDTO
6. **GetMentorProfileQuery**(mentor_id) → Fetch() → MentorProfileDTO
7. **ListAvailableMentorsQuery**(skills) → Filter() | Fetch() → MentorListDTO

#### Projections
- mentorship_read
- mentor_profile_read

#### Events Published
- mentorship.requested.v1
- mentorship.accepted.v1
- session.scheduled.v1
- mentorship.completed.v1

#### RBAC/SLO
- **RBAC:** OWNER (request), MENTOR (accept/schedule), PUBLIC (view mentors)
- **SLO:** P95 < 170ms

---

### 20.3 achievements/

#### User Stories
- As a **freelancer**, I want to **unlock achievements** (FirstJob, 10Jobs, TopRated, QuickResponder) so that my progress is recognized.
- As a **system**, I want to **track achievement progress** so that users see how close they are.
- As a **system**, I want to **have achievement tiers** (Bronze, Silver, Gold, Platinum) so that progression is gamified.
- As a **system**, I want to **emit achievement events** so that badging service can issue badges.
- As a **freelancer**, I want to **view my achievements** so that I see my accomplishments.

#### Flow
1. **RecordAchievementProgressCommand**(user_id, achievement_type, progress_increment) → Update() | CheckUnlock() → **Outbox:** achievement_progress.updated.v1
2. **UnlockAchievementCommand**(user_id, achievement_type, tier) → Validate() | Unlock() → **Outbox:** achievement.unlocked.v1
3. **GetAchievementsQuery**(user_id) → Fetch() → AchievementListDTO
4. **GetAchievementProgressQuery**(user_id, achievement_type) → Fetch() → AchievementProgressDTO

#### Projections
- achievements_read
- achievement_progress_read

#### Events Published
- achievement_progress.updated.v1
- achievement.unlocked.v1
- tier.reached.v1

#### Events Consumed
- contract.completed.v1 (to update progress)
- review.submitted.v1 (to update progress)

#### RBAC/SLO
- **RBAC:** SYSTEM (record/unlock), OWNER (view)
- **SLO:** P95 < 150ms

---

## **21 - COMPLIANCE DOMAIN**

### 21.1 compliance/

#### User Stories
- As a **freelancer**, I want to **submit my tax profile** (country, VAT/GST, TIN) so that I comply with tax regulations.
- As a **freelancer**, I want to **update my residency information** so that my location is current.
- As a **freelancer**, I want to **attach tax documents** (W-8/W-9/VAT docs) so that I can prove compliance.
- As a **system**, I want to **validate country-specific fields** so that compliance data is accurate.
- As an **admin**, I want to **review compliance artifacts** so that I can verify legitimacy.

#### Flow
1. **CreateOrUpdateTaxProfileCommand**(user_id, country, vat_gst, tin, w_form_refs) → ValidateCountryFields() | Persist() → **Outbox:** tax_profile.updated.v1
2. **SetResidencyCommand**(user_id, country, since, proof_docs) → Validate() | Update() → **Outbox:** residency.updated.v1
3. **AttachWFormCommand**(user_id, form_type, document_url) → ValidateForm() | Attach() → **Outbox:** w_form.attached.v1
4. **AttachVATCommand**(user_id, vat_number, document_url) → ValidateVAT() | Attach() → **Outbox:** vat.attached.v1
5. **GetTaxProfileQuery**(user_id) → AuthorizeOwner() | Fetch() → TaxProfileDTO
6. **GetResidencyQuery**(user_id) → AuthorizeOwner() | Fetch() → ResidencyDTO
7. **ListComplianceArtifactsQuery**(user_id) → AuthorizeOwner() | Fetch() → ComplianceArtifactListDTO

#### Projections
- compliance_read
- tax_profile_read
- residency_read

#### Events Published
- tax_profile.updated.v1
- residency.updated.v1
- w_form.attached.v1
- vat.attached.v1
- compliance_artifact.added.v1

#### RBAC/SLO
- **RBAC:** OWNER (submit/update/view), ADMIN (review artifacts)
- **SLO:** P95 < 170ms

---

## **22 - COMMUNICATION PREFERENCES DOMAINS**

### 22.1 communication_channels/

#### User Stories
- As a **user**, I want to **add communication channels** (Email, SMS, Push, InApp, WhatsApp, Slack) so that I can be reached.
- As a **user**, I want to **set channel preferences** so that I control where notifications go.
- As a **user**, I want to **define quiet hours** so that I'm not disturbed during specific times.
- As a **system**, I want to **route notifications to preferred channels** so that communication is effective.

#### Flow
1. **AddChannelCommand**(user_id, channel_type, channel_details) → Validate() | Add() → **Outbox:** channel.added.v1
2. **RemoveChannelCommand**(user_id, channel_id) → Validate() | Remove() → **Outbox:** channel.removed.v1
3. **SetChannelPreferencesCommand**(user_id, preferences) → Validate() | Update() → **Outbox:** channel_preferences.updated.v1
4. **ConfigureQuietHoursCommand**(user_id, start_time, end_time, timezone) → Validate() | Configure() → **Outbox:** quiet_hours.set.v1
5. **GetChannelsQuery**(user_id) → AuthorizeOwner() | Fetch() → CommunicationChannelListDTO
6. **GetChannelPreferencesQuery**(user_id) → AuthorizeOwner() | Fetch() → ChannelPreferencesDTO
7. **GetQuietHoursQuery**(user_id) → AuthorizeOwner() | Fetch() → QuietHoursDTO

#### Projections
- communication_channels_read

#### Events Published
- channel.added.v1
- channel.removed.v1
- channel_preferences.updated.v1
- quiet_hours.set.v1

#### RBAC/SLO
- **RBAC:** OWNER (add/remove/configure/view)
- **SLO:** P95 < 150ms

---

### 22.2 email_preferences/

#### User Stories
- As a **user**, I want to **set email frequency** (RealTime, Daily, Weekly, Never) so that I control email volume.
- As a **user**, I want to **configure email categories** (JobAlerts, Messages, Updates, Marketing) so that I receive relevant emails.
- As a **user**, I want to **enable email digests** so that I get batched updates instead of individual emails.
- As a **user**, I want to **mute specific categories** so that I don't receive certain types of emails.

#### Flow
1. **UpdateEmailFrequencyCommand**(user_id, category, frequency) → Validate() | Update() → **Outbox:** email_frequency.updated.v1
2. **UpdateCategoryPreferencesCommand**(user_id, category_prefs) → Validate() | Update() → **Outbox:** category_preferences.updated.v1
3. **EnableDigestCommand**(user_id, digest_type, schedule) → Validate() | Enable() → **Outbox:** digest.enabled.v1
4. **MuteCategoryCommand**(user_id, category) → Validate() | Mute() → **Outbox:** category.muted.v1
5. **GetEmailPreferencesQuery**(user_id) → AuthorizeOwner() | Fetch() → EmailPreferencesDTO
6. **GetFrequencySettingsQuery**(user_id) → AuthorizeOwner() | Fetch() → FrequencySettingsDTO
7. **GetDigestSettingsQuery**(user_id) → AuthorizeOwner() | Fetch() → DigestSettingsDTO

#### Projections
- email_preferences_read

#### Events Published
- email_frequency.updated.v1
- category_preferences.updated.v1
- digest.enabled.v1
- category.muted.v1

#### RBAC/SLO
- **RBAC:** OWNER (update/view)
- **SLO:** P95 < 140ms

---

## **GLOBAL CONVENTIONS & PLATFORM ALIGNMENT**

### Event Envelope Structure
All events published from users-be follow the standard envelope:
```json
{
  "event_id": "uuid",
  "event_type": "user.created.v1",
  "event_version": "1.0",
  "aggregate_id": "user_id",
  "aggregate_type": "user",
  "timestamp": "ISO8601",
  "correlation_id": "trace_id",
  "causation_id": "parent_event_id",
  "actor": {
    "user_id": "actor_user_id",
    "type": "user|system|admin"
  },
  "metadata": {
    "source_service": "users-be",
    "idempotency_key": "unique_key"
  },
  "payload": { }
}
```

### Idempotent Write-Path
- All commands use **idempotency keys** to prevent duplicate processing
- Idempotency keys stored in outbox table with TTL
- Responses cached for duplicate requests within TTL window
- Commands include: `CreateUserCommand`, `UpdateProfileCommand`, `AddSkillCommand`, etc.

### Non-PII Event Payloads
- Events NEVER contain PII in payload (no emails, phones, addresses, SSNs)
- Events reference IDs only: `user_id`, `profile_id`, `certification_id`
- Consumers fetch PII via authenticated API calls if needed
- Example: `user.created.v1` contains `user_id` but NOT `email`

### Folder Structure Alignment
- All domain entities map to `internal/domain/{context}/`
- All repositories map to `internal/infrastructure/persistence/postgres/{context}_repository.go`
- All services map to `internal/application/{context}/service.go`
- All handlers map to `internal/interfaces/http/v1/handlers/{context}_handler.go`
- All routes map to `internal/interfaces/http/v1/routes/{context}_routes.go`

### Events Catalog Integration
- All events published are registered in `contracts/events/users/` catalog
- Event schemas versioned with semantic versioning (v1, v2, etc.)
- Breaking changes require new event version
- Consumers subscribe via Dapr pub/sub with scopes: `["users-be"]`

### Caching Strategy
- Cache keys follow pattern: `users:{user_id}:{context}:{version}`
- TTLs defined in `internal/infrastructure/cache/redis/keys.go`
- Invalidation rules map events to cache keys in `invalidation_rules.go`
- Singleflight prevents cache stampedes for hot keys

### Observability
- All commands/queries emit spans with OpenTelemetry
- Metrics tracked: P95 latency, error rate, event publish lag
- Structured logging with correlation_id for tracing
- Health checks: `/healthz/live` (liveness), `/healthz/ready` (readiness)

### Security
- All endpoints require JWT authentication via Keycloak
- RBAC enforced at service layer (OWNER, ADMIN, SYSTEM, PUBLIC)
- PII encrypted at rest using KMS envelope encryption
- Sensitive fields redacted in logs via PII redactor

### Data Retention & Erasure
- User data retention: 7 years (compliance)
- Event logs retention: 90 days (projections can replay)
- GDPR/CCPA erasure: `DELETE /users/:id/erase` triggers cascading deletion
- Erasure emits `user.erased.v1` event for downstream cleanup

---

## **END OF USERS-BE USER STORIES**

**Total Domains Covered:** 22  
**Total Sections:** 42  
**Total User Stories:** 350+  
**Total Flows:** 250+  
**Total Events:** 200+  

All stories follow the pattern: **Stories → Flow → Projections → Events → RBAC/SLO**  
All flows include: **idempotent write-path, event envelope, non-PII payloads**  
All components align with: **folder structure, events catalog, platform conventions**ubmitKYCCommand**(user_id, id_document, passport, address_proof, selfie) → ValidateDocuments() | Persist() → **Outbox:** kyc.submitted.v1
2. **SubmitKYBCommand**(org_id, business_docs) → ValidateBusinessDocs() | Persist() → **Outbox:** kyb.submitted.v1
3. **ApproveVerificationCommand**(verification_id, approved_by) → Validate() | Approve() → **Outbox:** identity_verification.approved.v1
4. **RejectVerificationCommand**(verification_id, reason, rejected_by) → Validate() | Reject() → **Outbox:** identity_verification.rejected.v1
5. **GetVerificationStatusQuery**(user_id) → Fetch() → VerificationStatusDTO

#### Projections
- identity_verification_read

#### Events Published
- kyc.submitted.v1
- kyb.submitted.v1
- identity_verification.approved.v1
- identity_verification.rejected.v1

#### RBAC/SLO
- **RBAC:** OWNER (submit), ADMIN (approve/reject), OWNER (view status)
- **SLO:** P95 < 200ms

---
