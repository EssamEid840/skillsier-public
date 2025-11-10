Skillsier — Complete Frontend Journeys & Flows (Web + Mobile)
=============================================================

**Version:** 2.0 (Enhanced & Complete)**Last Updated:** 2025-11-06

This document provides comprehensive, production-ready user journeys for the Skillsier freelancing platform. It aligns with the frontend route-group conventions from combined-fe-folder-strucure.md and integrates with all 11 backend microservices.

Document Structure
------------------

For each journey:

*   **ID**: Unique journey identifier (for EPIC/Story tracking)
*   **Persona**: Target user role(s)
*   **Preconditions & Triggers**: Entry points and prerequisites
*   **Primary Screens**: Web and Mobile routes using (group)/route naming
*   **System Touchpoints**: Backend microservices involved
*   **Flow Steps**: Happy path numbered sequence
*   **Branches & Edge Cases**: Alternative paths, error handling, limits
*   **Notifications**: In-app, push, email rules
*   **Analytics**: Suggested event tracking
*   **Sources**: Reference documents from project knowledge
    

Table of Contents
-----------------


    

A) Onboarding & First-Run
-------------------------

### ONB-1 — Sign Up → Verify → Choose Roles

**Persona:** Any (New)**Preconditions & Triggers:** None; arrives via marketing/referral/invite
**Primary Screens:**
*   **Web:** /app/(auth)/sign-up, /app/(auth)/verify, /app/(onboarding)
*   **Mobile:** /(auth)/sign-up, /(auth)/verify, /(tabs)/(authenticated)/(onboarding)
**System Touchpoints:** users-be (accounts/sessions), communications-be (OTP/email/SMS), admin-be (policy gates)
**Flow Steps:**

1.  Choose sign-up method (email/password or SSO: Google, LinkedIn, GitHub)
    
2.  Submit credentials → users-be creates account stub; password strength validated
    
3.  Send verification (OTP via email/SMS or magic link)
    
4.  User verifies → account activated
    
5.  Choose role(s): Client, Freelancer (multi-select allowed)
    
6.  Minimal profile seed (name, country, timezone, locale, currency)
    
7.  Role-aware onboarding checklist created
    
8.  Redirect to role-specific onboarding flow (ONB-2 for Client, ONB-3 for Freelancer)
    

**Branches & Edge Cases:**
*   Existing account → redirect to sign-in with "Already have an account?" link
*   Rate-limit signup attempts (5 per hour per IP)
*   Blocked/disposable email domains → show error
*   Weak password → show strength indicator and requirements
*   SSO failure → fallback to email/password with clear error message
*   Verification expiry (10 minutes) → resend option
*   Multi-role selection → show combined checklist
*   Country-specific requirements (e.g., EU GDPR consent, age verification)
**Notifications:**
*   Welcome email with "Complete your profile" CTA
*   In-app nudges for profile completion milestones
**Analytics:**
*   auth.signup\_initiated
*   auth.signup\_method\_selected (email/sso)
*   auth.verify\_sent
*   auth.verify\_ok
*   role.selected (client/freelancer/both)
*   onboarding.started
**Sources:** users-be.user-stories.md, communications-be-folder-structure.md, combined-fe-folder-strucure.md

### ONB-2 — Client Fast-Start

**Persona:** Client (New)**Preconditions & Triggers:** Completed ONB-1 with Client role selected
**Primary Screens:**
*   **Web:** /app/(settings)/organization, /app/(billing)/payment-methods, /app/(jobs)/new
*   **Mobile:** /(tabs)/(authenticated)/(settings)/organization, /(tabs)/(authenticated)/(billing)/payment-methods, /(tabs)/(authenticated)/(jobs)/new
**System Touchpoints:** users-be, financial-be, jobs-be, search-be, subscriptions-be
**Flow Steps:**

1.  Company/org setup wizard:
    
    *   Organization name, size, industry
    *   VAT/Tax ID (required for EU/business accounts)
    *   Billing contact and address
    *   Logo upload (optional)
        
2.  Add payment method:
    
    *   Choose type: Credit Card, ACH/Bank Transfer, PayPal, or Crypto Wallet
    *   Tokenize via financial-be → Stripe/payment processor
    *   Set as default payment method
    *   Verify via micro-transaction (for bank accounts)
        
3.  Plan selection (if not free tier):
    
    *   View plans comparison
    *   Select plan → subscriptions-be
    *   Allocate connects budget
        
4.  Create first job (guided):
    
    *   Job title and description
    *   Select category and skills (autocomplete)
    *   Budget/rate and hiring type (hourly/fixed)
    *   Visibility (Public/Invite-only/Private link)
    *   Screening questions (optional, templates available)
    *   Attachments (optional)
    *   Preview → Publish
        
5.  Job indexed by search-be
    
6.  Receive talent recommendations and invite suggestions
    

**Branches & Edge Cases:**
*   Skip organization setup → minimal profile, can complete later
*   Payment method failure → show specific error (card declined, insufficient funds)
*   Job preview warnings (missing skills, budget too low/high for market)
*   Compliance flags (restricted categories, sanctioned countries) → hold for review
*   Invite-only job → skip indexing
*   Draft auto-save every 30 seconds
*   Template usage → pre-fill from saved templates
**Notifications:**
*   Org setup completion confirmation
*   Payment method added/verified
*   Job published with performance tracking link
*   Talent suggestions and invite opportunities
**Analytics:**
*   org.created
*   pm.added (payment\_method\_type)
*   plan.selected (plan\_tier)
*   job.draft\_created
*   job.published
*   onboarding.client\_completed
**Sources:** users-be.user-stories.md, financial-be.user-stories.md, jobs-be.user-stories.md, combined-fe-folder-strucure.md

### ONB-3 — Freelancer Fast-Start

**Persona:** Freelancer (New)**Preconditions & Triggers:** Completed ONB-1 with Freelancer role selected
**Primary Screens:**
*   **Web:** /app/(profile)/edit, /app/(billing)/payouts, /app/(market)/alerts, /app/(profile)/portfolio
*   **Mobile:** /(tabs)/(authenticated)/(profile)/edit, /(tabs)/(authenticated)/(billing)/payouts, /(tabs)/(authenticated)/(market)/alerts, /(tabs)/(authenticated)/(profile)/portfolio
**System Touchpoints:** users-be, admin-be (KYC), financial-be (payout), communications-be (alerts), search-be
**Flow Steps:**

1.  Complete professional profile (wizard):
    
    *   Professional headline (80 chars)
    *   Professional summary (5000 chars)
    *   Skills with proficiency levels (max 15 primary skills)
    *   Hourly rate or project budget range
    *   Languages with proficiency (native/fluent/conversational/basic)
    *   Education and certifications
    *   Work experience (title, company, dates, description)
        
2.  KYC/ID verification (if required by country/threshold):
    
    *   Document upload: ID/Passport + selfie
    *   Address proof (utility bill, bank statement < 3 months old)
    *   Submit to admin-be for review
    *   Real-time status tracking
        
3.  Add payout method:
    
    *   Choose type: Bank Transfer, PayPal, Payoneer, Crypto
    *   Enter account details → tokenized
    *   Verify via micro-deposit (for bank accounts)
    *   Set default payout method
        
4.  Portfolio setup:
    
    *   Upload 3-10 work samples
    *   Add project descriptions, skills used, outcomes
    *   Link to external portfolio/GitHub
    *   Video introduction (optional, 60 seconds max)
        
5.  Create saved searches & job alerts:
    
    *   Set search criteria (skills, budget, job type, location)
    *   Alert frequency (real-time, daily digest, weekly)
    *   Alert channels (email, push, in-app)
        
6.  Profile optimization score shown (aim for 100%)
    
7.  Suggested skills tests to increase visibility
    

**Branches & Edge Cases:**
*   Skip steps → save as draft, continue later (profile completeness score shown)
*   KYC rejection → clear reason provided, retry with corrected docs
*   Unsupported payout corridor → show alternatives, link to FAQ
*   Locale-specific tax fields (W-9 for US, W-8BEN for non-US, VAT for EU)
*   Portfolio item too large → compress or link externally
*   Skills mismatch → suggestions based on experience/education
*   Rate too high/low for market → show market insights
*   Profile visibility settings → choose who can see full profile
**Notifications:**
*   Profile milestones reached (25%, 50%, 75%, 100% complete)
*   KYC approved/rejected with next steps
*   Payout method added/verified
*   Job alert confirmations with sample matches
*   Skills test recommendations
**Analytics:**
*   profile.created
*   profile.milestone\_reached (completion\_percent)
*   kyc.submitted
*   kyc.approved|rejected
*   payout.added (method\_type)
*   portfolio.item\_added
*   alert.created
*   onboarding.freelancer\_completed
**Sources:** users-be.user-stories.md, admin-be.user-stories.md, financial-be.user-stories.md, combined-fe-folder-strucure.md

### ONB-4 — Post-Signup Profile Completion

**Persona:** Both (New users with incomplete profiles)**Preconditions & Triggers:** Signed up but profile < 70% complete; returns to platform
**Primary Screens:**
*   **Web:** /app/(onboarding)/checklist, /app/(profile)/edit
*   **Mobile:** /(tabs)/(authenticated)/(onboarding)/checklist, /(tabs)/(authenticated)/(profile)/edit
**System Touchpoints:** users-be, subscriptions-be, communications-be
**Flow Steps:**

1.  Dashboard shows profile completion widget with progress bar
    
2.  Checklist of remaining items:
    
    *   Profile photo
    *   Professional summary
    *   Skills (minimum 5)
    *   Portfolio items (freelancers)
    *   Payment method (clients)
    *   Payout method (freelancers)
    *   Email verification
    *   Phone verification (optional but recommended)
        
3.  Click item → navigate to relevant screen with contextual help
    
4.  Complete item → check mark, progress updates
    
5.  Reach 100% → unlock benefits:
    
    *   Higher search ranking
    *   Profile badge
    *   Bonus connects (clients)
    *   Featured status (freelancers, temporary)
        
6.  Celebration modal on completion
    

**Branches & Edge Cases:**
*   Dismiss widget → reappear after 7 days or on login
*   Skip specific items → show impact on profile strength
*   Partial completion → save progress automatically
*   Return user → resume from last completed item
*   Mobile vs web → sync progress across devices
*   Profile optimization suggestions beyond checklist
**Notifications:**
*   Daily reminder emails (max 3, then stop to avoid spam)
*   Push notification on mobile (if enabled)
*   In-app banner on key pages
**Analytics:**
*   onboarding.checklist\_viewed
*   onboarding.item\_completed (item\_type)
*   onboarding.profile\_completed
*   onboarding.dismissed
**Sources:** users-be.user-stories.md, combined-fe-folder-strucure.md

### ONB-5 — Guided Tutorials & Tooltips

**Persona:** Both (New users, first-time feature users)**Preconditions & Triggers:** First time accessing specific features
**Primary Screens:**
*   **Web:** Overlays on various screens with spotlight/tooltips
*   **Mobile:** Same + native bottom sheets
**System Touchpoints:** users-be (tutorial completion tracking), communications-be
**Flow Steps:**

1.  Detect first-time access to major feature areas:
    
    *   Job posting (clients)
    *   Proposal submission (freelancers)
    *   Inbox/messaging
    *   Contract creation
    *   Workroom
    *   Billing/payments
        
2.  Show contextual tutorial:
    
    *   Dimmed background with spotlight on relevant UI elements
    *   Step-by-step tooltips (3-5 steps max)
    *   "Next" / "Skip Tour" buttons
    *   Progress dots indicator
        
3.  User completes or skips tour
    
4.  Tutorial completion tracked per feature
    
5.  Option to replay tutorials from help menu
    

**Branches & Edge Cases:**
*   Skip tour → can access from "?" help icon
*   Mobile: use native bottom sheets for better UX
*   Tutorial shown max once unless user resets
*   Admin can force-show tutorial on major feature changes
*   A/B test tutorial effectiveness
*   Video tutorials for complex features (e.g., contract setup)
**Notifications:**
*   None (in-context only)
**Analytics:**
*   tutorial.started (feature\_area)
*   tutorial.step\_viewed (step\_number)
*   tutorial.completed
*   tutorial.skipped (step\_number)
*   tutorial.replayed
**Sources:** combined-fe-folder-strucure.md, users-be.user-stories.md

B) Client (Hiring Lifecycle)
----------------------------

### CL-1 — Job - Draft → Publish → Manage

**Persona:** Client**Preconditions & Triggers:** User has payment method and sufficient connects (if required by plan)
**Primary Screens:**
*   **Web:** /app/(jobs)/new, /app/(jobs)/\[jobId\]/edit, /app/(jobs)/\[jobId\], /app/(jobs)/\[jobId\]/analytics
*   **Mobile:** /(tabs)/(authenticated)/(client)/jobs/new, /(tabs)/(authenticated)/(client)/jobs/\[jobId\]/edit, /(tabs)/(authenticated)/(client)/jobs/\[jobId\], /(tabs)/(authenticated)/(client)/jobs/\[jobId\]/analytics
**System Touchpoints:** jobs-be, search-be, communications-be, subscriptions-be, storage-be
**Flow Steps:**

1.  Create new job → navigate to /app/(jobs)/new
    
2.  Fill job details (auto-save every 30 seconds):
    
    *   Job title (required, 5-100 chars)
    *   Category and subcategory (hierarchical selection)
    *   Skills required (autocomplete, max 20, mark primary vs nice-to-have)
    *   Scope: One-time project, ongoing, contract-to-hire
    *   Experience level: Entry, Intermediate, Expert
    *   Job description (rich text editor, 100-5000 chars)
    *   Budget: Hourly rate range OR fixed price
    *   Duration estimate (hours/weeks)
    *   Time zone preference (optional)
    *   Attachments (briefs, designs, etc., max 10 files, 25MB each)
        
3.  Configure job settings:
    
    *   Visibility: Public / Invite-only / Private link
    *   Application deadline (optional)
    *   Max proposals to receive (optional, default 50)
    *   Screening questions (add custom, max 5)
    *   Require portfolio samples
    *   Require cover letter
        
4.  Job promotion (optional):
    
    *   Featured listing (costs extra connects)
    *   Urgent badge (highlighted in search)
    *   Target specific freelancer segments
        
5.  Preview job → check formatting, detect issues (missing skills, unclear budget)
    
6.  Publish:
    
    *   Deduct connects (if applicable)
    *   Job goes live immediately or scheduled
    *   Indexed by search-be (if public)
    *   Notification sent to matching freelancers with active alerts
        
7.  Post-publish management:
    
    *   View proposals as they arrive
    *   Edit job (minor edits allowed: description, deadline, budget increase)
    *   Major edits (skills, category) require republishing
    *   Close job early
    *   Archive job (removes from search)
    *   Repost job (creates new job ID, preserves template)
    *   Clone job for similar positions
        
8.  Job analytics:
    
    *   Views count
    *   Proposal submissions
    *   Click-through rate (CTR) from search
    *   Conversion funnel
    *   Top skills in applicants
    *   Average bid amount
**Branches & Edge Cases:**
*   Draft expires after 90 days → notification before deletion
*   Budget change after proposals received → notify all applicants
*   Clone job → pre-fill from existing job, review before publishing
*   Repost with improvements → suggest based on performance data
*   Restricted categories (gambling, adult content, etc.) → approval required
*   Country restrictions → verify compliance
*   Insufficient connects → prompt to purchase or upgrade plan
*   Duplicate job detection → warn before publishing
*   Template creation → save successful job as template for future use
*   Syndication to external job boards (optional, costs extra)
*   A/B test job variants to optimize response rate
*   Talent cloud integration → also post to agency marketplace
*   Screening questions library → select from popular questions
*   Auto-close after deadline
*   Extend deadline with notification to existing applicants
*   Maximum active jobs per plan tier
**Notifications:**
*   Job published confirmation with link and preview
*   Proposals received (configurable: real-time, hourly digest, daily)
*   Performance milestone (e.g., 50 views, 10 proposals)
*   Job closing reminder (if deadline approaching)
*   Low proposal count → suggestions to improve job posting
**Analytics:**
*   job.draft\_created
*   job.draft\_update
*   job.preview\_viewed
*   job.published (visibility, budget\_type, experience\_level, featured)
*   job.edited
*   job.closed (reason: hired/cancelled/expired)
*   job.archived
*   job.reposted
*   job.cloned
*   job.analytics\_viewed
*   job.promotion\_added
**Sources:** jobs-be.user-stories.md, jobs-be-folder-structure.md, jobs-be.database-design.md, combined-fe-folder-strucure.md

### CL-2 — Talent Sourcing & Invites

**Persona:** Client**Preconditions & Triggers:** Has published job or planning to hire
**Primary Screens:**
*   **Web:** /app/(talent), /app/(talent)/\[freelancerId\], /app/(jobs)/\[jobId\]/invite
*   **Mobile:** /(tabs)/(authenticated)/(talent), /(tabs)/(authenticated)/(talent)/\[freelancerId\], /(tabs)/(authenticated)/(client)/jobs/\[jobId\]/invite
**System Touchpoints:** search-be, users-be, communications-be, jobs-be, subscriptions-be
**Flow Steps:**

1.  Navigate to Talent Search → /app/(talent)
    
2.  Apply filters:
    
    *   Skills (AND/OR logic, tag-based)
    *   Hourly rate range
    *   Success rate (job success score)
    *   Earned amount (total lifetime earnings)
    *   Availability (available now, within 1 week, etc.)
    *   Location/timezone (single or multiple)
    *   Languages spoken
    *   Badges and certifications
    *   Top Rated, Rising Talent, etc.
        
3.  Browse search results (paginated, 20 per page):
    
    *   Freelancer card shows: name, photo, headline, rate, location, success rate, badges
    *   Quick actions: Save, Invite, View profile
        
4.  Click freelancer → view full profile:
    
    *   Professional summary
    *   Skills with endorsements
    *   Portfolio items with thumbnails/previews
    *   Work history with feedback (most recent first)
    *   Reviews and ratings (overall + category breakdown)
    *   Earnings visibility (if public)
    *   Availability calendar
    *   Response rate and time
    *   Video introduction (if available)
    *   Recommendations
        
5.  Send personalized invite:
    
    *   Select job to invite for (or write custom project description)
    *   Add personal message (recommended, 100-1000 chars)
    *   Include budget/rate information
    *   Set expiry (7 days default)
        
6.  Bulk invites:
    
    *   Select multiple freelancers (max 20 per job)
    *   Send batch invites with same job/message
    *   Rate limit: max 50 invites per week per job
        
7.  Track invites:
    
    *   View sent invites list
    *   See who viewed, accepted, declined
    *   Resend if no response after 3 days
    *   Revoke invite if position filled
**Branches & Edge Cases:**
*   Rate-limit invites: 50 per week per job (prevents spam)
*   Dedupe against existing conversations (don't invite if already in conversation)
*   Freelancer has invites disabled → show "Not accepting invites" badge
*   Freelancer blocking client → invite fails silently (privacy protection)
*   Saved searches → create and manage saved talent searches
*   Talent collections → save freelancers to custom lists for future projects
*   Smart recommendations → AI-suggested freelancers based on job requirements
*   Previously hired freelancers → quick re-hire button
*   Freelancer unavailable → show next available date
*   Premium plan feature → see freelancer contact history
**Notifications:**
*   Freelancer receives invite (in-app, push, email based on preferences)
*   Client receives notification when invite is viewed/accepted/declined
*   Reminder to follow up on pending invites
**Analytics:**
*   talent.search (filters\_applied)
*   talent.profile\_viewed (from\_search, from\_recommendation)
*   invite.sent (personalized, bulk)
*   invite.viewed
*   invite.accepted
*   invite.declined (reason if provided)
*   invite.expired
*   invite.revoked
*   talent.saved (to\_list)
*   talent.recommended\_clicked
**Sources:** search-be.user-stories.md, users-be.user-stories.md, communications-be-folder-structure.md, combined-fe-folder-strucure.md

### CL-3 — Proposal Intake & Shortlisting

**Persona:** Client**Preconditions & Triggers:** Job published and receiving proposals
**Primary Screens:**
*   **Web:** /app/(jobs)/\[jobId\]/proposals, /app/(proposals)/\[proposalId\]
*   **Mobile:** /(tabs)/(authenticated)/(client)/jobs/\[jobId\]/proposals, /(tabs)/(authenticated)/(proposals)/\[proposalId\]
**System Touchpoints:** proposals-be, communications-be, search-be, users-be, jobs-be
**Flow Steps:**

1.  Navigate to job proposals list → /app/(jobs)/\[jobId\]/proposals
    
2.  View proposals overview:
    
    *   Total proposals received
    *   Unread count (badge notification)
    *   Shortlisted count
    *   Archived count
    *   Filter options at top
        
3.  Apply filters and sorting:
    
    *   Status: All / Unread / Shortlisted / Archived / Declined
    *   Price: Lowest to highest / Highest to lowest
    *   Relevance score (AI match score)
    *   Date submitted (newest/oldest first)
    *   Freelancer badges (Top Rated, Rising Talent)
    *   Response time to invitation
        
4.  Browse proposal cards (list or grid view):
    
    *   Freelancer photo, name, headline
    *   Bid amount (hourly rate or fixed price)
    *   Proposal excerpt (first 100 chars)
    *   Match score indicator (0-100%)
    *   Quick stats: success rate, earnings, location
    *   Quick actions: Message, Shortlist, Decline, Archive
        
5.  Click proposal → view full details:
    
    *   Cover letter (formatted text)
    *   Answers to screening questions
    *   Attached portfolio samples relevant to job
    *   Proposed timeline/milestones (for fixed-price)
    *   Bid breakdown (if itemized)
    *   Similar past work with outcomes
    *   Freelancer full profile (side panel or tab)
        
6.  Review freelancer profile context:
    
    *   Reviews from past clients
    *   Work history with this client (if any)
    *   Success rate on similar projects
    *   Portfolio items
    *   Skills match percentage
    *   Availability
        
7.  Take actions:
    
    *   **Message:** Open conversation thread (navigate to inbox)
    *   **Shortlist:** Add to shortlist (max 20 per job)
    *   **Decline:** Remove from active list, freelancer notified
    *   **Archive:** Hide from view but keep for records
    *   **Add internal notes:** Private notes visible only to hiring team
    *   **Add tags:** Custom tags for organization (e.g., "second round", "backup option")
        
8.  Batch actions (select multiple proposals):
    
    *   Decline all selected
    *   Archive all selected
    *   Export to CSV for review
        
9.  Shortlist management:
    
    *   View all shortlisted proposals side-by-side
    *   Compare proposals (side-by-side table view)
    *   Schedule interviews
    *   Send bulk messages to shortlisted candidates
        
10.  Decision tracking:
    
    *   Mark proposal as "Interviewing"
    *   Mark as "Offer Extended"
    *   Mark as "Hired" (converts to contract)
**Branches & Edge Cases:**
*   No proposals after 7 days → suggestions to improve job posting
*   Too many proposals (>50) → filtering becomes critical, consider narrowing job scope
*   Proposal withdrawal by freelancer → marked as "Withdrawn", not counted
*   Freelancer updates proposal → notification to client, "Updated" badge
*   Connects refund rules → if job closed without hire, partial refund logic
*   Auto-decline after 30 days if no response
*   Proposal expiry (freelancer can set expiry date)
*   Price negotiation → counter-offer feature
*   Red flags in proposals → AI flagging (plagiarism, unrealistic promises, etc.)
*   Duplicate applications (same freelancer to same job) → merge or flag
*   Team collaboration → share proposals with team members, add comments
*   Hiring team permissions → limit who can decline/shortlist
*   Proposal quality score → AI-generated score based on relevance, professionalism, past success
*   Recommendation engine → "You might also like these proposals"
**Notifications:**
*   New proposal received (real-time, digest, or daily based on settings)
*   Shortlisted proposal updated by freelancer
*   Proposal about to expire (freelancer set deadline)
*   Interview scheduled confirmation
**Analytics:**
*   proposal.list\_viewed (job\_id, filters\_applied)
*   proposal.viewed (proposal\_id, time\_spent)
*   proposal.shortlisted
*   proposal.declined (reason if provided)
*   proposal.archived
*   proposal.messaged\_from
*   proposal.compared (proposal\_ids)
*   proposal.batch\_action (action\_type, count)
*   proposal.internal\_note\_added
*   proposal.tag\_added
**Sources:** proposals-be.user-stories.md, proposals-be-folder-structure.md, proposals-be.database-design.md, combined-fe-folder-strucure.md

### CL-4 — Interviews - Schedule → Call → Decision)

**Persona:** Client**Preconditions & Triggers:** Shortlisted candidates ready for interview
**Primary Screens:**
*   **Web:** /app/(proposals)/\[proposalId\]/interview, /app/(inbox)/\[conversationId\]/call, /app/(jobs)/\[jobId\]/interviews
*   **Mobile:** /(tabs)/(authenticated)/(proposals)/\[proposalId\]/interview, /(tabs)/(authenticated)/(messages)/\[conversationId\]/call, /(tabs)/(authenticated)/(client)/jobs/\[jobId\]/interviews + native calendar integration
**System Touchpoints:** proposals-be, communications-be, contracts-be, users-be
**Flow Steps:**

1.  From shortlisted proposal → click "Schedule Interview"
    
2.  Interview scheduling form:
    
    *   Select date and time (client's local time shown)
    *   Duration (15/30/45/60 minutes, default 30)
    *   Type: Video call (in-platform), Phone, or In-person (rare)
    *   Add agenda/questions (optional, private notes)
    *   Include team members (CC additional interviewers)
    *   Send calendar invite (ICS file attached)
        
3.  Freelancer receives interview request:
    
    *   View proposed times
    *   Accept, decline, or propose alternate times
    *   Add to personal calendar
        
4.  Confirmation once accepted:
    
    *   Both parties receive confirmation email with details
    *   Calendar event created
    *   In-app notification 15 minutes before interview
    *   Join call link active 5 minutes before start
        
5.  Conduct interview (in-platform video call):
    
    *   Navigate to call screen (same UI as regular calls)
    *   Video/audio controls
    *   Screen sharing option
    *   Chat sidebar for notes/links
    *   Record call option (with consent notification)
        
6.  During interview:
    
    *   Take private notes (not visible to freelancer)
    *   Share screen to review portfolio or project brief
    *   Co-browsing (if enabled)
        
7.  Post-interview:
    
    *   Interview marked as "Completed"
    *   Prompt to add feedback/notes
    *   Rate interview quality (internal only)
    *   Next step actions: "Make Offer", "Decline", "Schedule Follow-up"
        
8.  Interview management dashboard:
    
    *   View all upcoming/past interviews for a job
    *   Track interview statuses
    *   Interviewer assignments
    *   Interview notes and feedback
    *   Comparison matrix across candidates
**Branches & Edge Cases:**
*   Rescheduling → client or freelancer can request reschedule, must be accepted
*   No-show tracking → mark as no-show, automatic notification
*   Time zone confusion → always show both parties' local times
*   Technical issues during call → fallback to phone number, email support
*   Interview recording → requires explicit consent, stored securely, auto-delete after 30 days
*   Multiple interviewers → panel interview mode with shared notes
*   Interview templates → save common questions/agendas
*   Follow-up interview → schedule second round easily
*   Decline after interview → send polite rejection message with optional feedback
*   Group interviews (multiple candidates) → not recommended but supported
*   Interview preparation materials → share docs/links with candidate beforehand
*   Calendar integration → sync with Google Calendar, Outlook
**Notifications:**
*   Interview request sent
*   Interview confirmed by freelancer
*   Interview rescheduled
*   Interview reminder (15 min before)
*   Interview starting now (push notification on mobile)
*   Interview completed (prompt for feedback)
*   Candidate declined interview (with reason if provided)
**Analytics:**
*   interview.scheduled (proposal\_id, duration, type)
*   interview.confirmed
*   interview.rescheduled
*   interview.started (on\_time, late)
*   interview.completed (duration\_actual)
*   interview.no\_show (party)
*   interview.declined (reason)
*   interview.feedback\_added
**Sources:** proposals-be.user-stories.md, communications-be-folder-structure.md, combined-fe-folder-strucure.md

### CL-5 — Hire: Contract Generation & Funding

**Persona:** Client**Preconditions & Triggers:** Decision made to hire a freelancer from proposals or invites
**Primary Screens:**
*   **Web:** /app/(contracts)/new, /app/(contracts)/\[contractId\], /app/(billing)/escrow
*   **Mobile:** /(tabs)/(authenticated)/(contracts)/new, /(tabs)/(authenticated)/(contracts)/\[contractId\], /(tabs)/(authenticated)/(billing)/escrow
**System Touchpoints:** contracts-be, financial-be, proposals-be, communications-be, subscriptions-be, storage-be
**Flow Steps:**

1.  From proposal or freelancer profile → click "Hire" or "Make Offer"
    
2.  Contract creation wizard:
    
    *   **Contract type:** Hourly / Fixed-price / Milestones / Retainer
    *   **Job title/description:** Pre-filled from job posting, editable
    *   **Start date:** Immediate or scheduled
    *   **End date / Duration:** Optional for ongoing work
        
3.  **For Hourly contracts:**
    
    *   Set hourly rate (pre-filled from proposal)
    *   Weekly limit (max hours per week, e.g., 40)
    *   Manual time or use tracker
    *   Timesheet approval process: Auto-approve or manual review
    *   Billing cycle: Weekly or bi-weekly
        
4.  **For Fixed-price contracts:**
    
    *   Total project amount (pre-filled from proposal)
    *   Payment structure:
        *   Single payment on completion
        *   Milestone-based payments (add milestones with amounts and deadlines)
    *   Each milestone: Title, description, amount, due date
    *   Deliverables checklist per milestone
        
5.  **For Retainer contracts:**
    
    *   Monthly/quarterly retainer amount
    *   Included hours per period
    *   Overage rate (for hours beyond included)
    *   Renewal terms (auto-renew or manual)
        
6.  **Contract terms:**
    
    *   Statement of Work (SOW): Detailed scope, rich text editor
    *   Attachments: Upload project briefs, designs (from storage-be)
    *   Payment terms: Net 7, Net 14, Net 30 days
    *   Intellectual property (IP) ownership: Client / Freelancer / Shared
    *   Confidentiality: NDA required (check box, template attached)
    *   Termination clause: Notice period (3/7/14/30 days)
    *   Dispute resolution: Platform mediation / Arbitration / Court
        
7.  **Escrow funding (for fixed-price and milestones):**
    
    *   Calculate total escrow amount (sum of all milestones or full amount)
    *   Add safety deposit (optional, 10-20% extra)
    *   Choose funding source:
        *   Credit card (immediate, fees apply)
        *   Bank account (ACH, 3-5 business days)
        *   Wallet balance (if sufficient)
    *   Review fees (platform fee, payment processing fee)
    *   Authorize and fund escrow
    *   Escrow held by platform, released per milestone completion
        
8.  Review contract summary:
    
    *   All terms displayed clearly
    *   Estimated total cost and payment schedule
    *   Freelancer earnings breakdown
    *   Platform fees transparency
        
9.  Send contract to freelancer:
    
    *   Freelancer receives notification (in-app, push, email)
    *   Contract appears in their "Offers" section
    *   Freelancer can:
        *   Accept contract as-is
        *   Request modifications (negotiate terms)
        *   Decline offer with reason
            
10.  Contract acceptance:
    
    *   E-signature by both parties (DocuSign-style, stored in storage-be)
    *   Contract becomes active immediately
    *   Workroom created automatically (/app/(workroom)/\[contractId\])
    *   Both parties receive confirmation and next steps
**Branches & Edge Cases:**
*   Negotiation round → freelancer requests changes, client can accept/counter/decline
*   Escrow funding failure → clear error message, retry or change payment method
*   Insufficient funds → prompt to add funds or change payment method
*   Contract templates → save commonly used contract structures
*   Multiple contracts with same freelancer → show history, suggest similar terms
*   Currency conversion → show amounts in both parties' currencies
*   VAT/GST handling → calculate and display taxes if applicable
*   Compliance checks → flag contracts exceeding thresholds for additional verification
*   Long-term contract → set review milestones every 30/60/90 days
*   Auto-renewal terms → clearly disclosed, cancelable
*   Subcontracting clause → allow/prohibit freelancer from subcontracting
*   Non-compete clause → optional, duration-limited
*   Contract amendments → post-signature changes require mutual agreement
*   Escrow refund policy → if contract terminated early, refund process
*   Payment holds → admin can place holds for disputes or compliance issues
*   Multi-currency support → escrow in USD but pay in local currency
*   Bonus payments → add bonus clause for exceptional work
*   Performance incentives → tie bonuses to KPIs (e.g., on-time delivery, quality score)
**Notifications:**
*   Contract offer sent (client confirmation)
*   Contract received (freelancer notification with 48-hour response reminder)
*   Contract accepted (both parties)
*   Contract declined (both parties, with reason)
*   Negotiation request (modifications proposed)
*   Escrow funded successfully
*   Contract signed and active
*   Workroom ready notification
**Analytics:**
*   contract.offer\_created (contract\_type, amount, duration)
*   contract.offer\_sent
*   contract.offer\_viewed (by\_freelancer)
*   contract.negotiation\_requested (by\_freelancer)
*   contract.offer\_accepted (by\_freelancer)
*   contract.offer\_declined (reason)
*   contract.escrow\_funded (amount, method)
*   contract.signed (both\_parties)
*   contract.activated
**Sources:** contracts-be.user-stories.md, contracts-be-folder-structure.md, financial-be.user-stories.md, combined-fe-folder-strucure.md

### CL-6 — Workroom: Collaboration & Files

**Persona:** Client**Preconditions & Triggers:** Active contract exists
**Primary Screens:**
*   **Web:** /app/(workroom)/\[contractId\], /app/(workroom)/\[contractId\]/files, /app/(workroom)/\[contractId\]/activity
*   **Mobile:** /(tabs)/(authenticated)/(workroom)/\[contractId\], /(tabs)/(authenticated)/(workroom)/\[contractId\]/files, /(tabs)/(authenticated)/(workroom)/\[contractId\]/activity
**System Touchpoints:** contracts-be, storage-be, communications-be, users-be
**Flow Steps:**

1.  Navigate to workroom → /app/(workroom)/\[contractId\]
    
2.  Workroom dashboard shows:
    
    *   Contract summary card (type, rate/budget, status, progress)
    *   Recent activity feed (last 20 actions)
    *   Quick actions: Upload file, Send message, Add milestone, Submit timesheet
    *   Navigation tabs: Overview, Files, Milestones (if applicable), Time (if hourly), Activity, Settings
        
3.  **Files management:**
    
    *   Navigate to Files tab
    *   Organized folder structure:
        *   Project Brief (auto-created)
        *   Deliverables (freelancer uploads)
        *   Feedback & Revisions (client uploads)
        *   Assets & Resources (shared resources)
        *   Custom folders (both parties can create)
    *   Upload files:
        *   Drag and drop (up to 10 files at once)
        *   Max 100MB per file, common formats supported
        *   Virus scan automatic before confirmation
        *   Thumbnail generation for images/videos
    *   File actions:
        *   Preview (in-browser for common formats)
        *   Download (generates signed URL, 24-hour validity)
        *   Share link (public or private with password)
        *   Add comments/annotations
        *   Request approval/feedback
        *   Version history (track changes)
        *   Rename, move to folder, delete
    *   File notifications:
        *   Notify when file uploaded (optional per folder)
        *   Notify when feedback requested
        *   Notify on file approval/rejection
            
4.  **Activity feed:**
    
    *   Chronological log of all workroom actions:
        *   File uploads/downloads
        *   Milestone submissions/approvals
        *   Timesheet submissions/approvals
        *   Contract amendments
        *   Messages sent
        *   Payments released
        *   Status changes
    *   Filter by: File activity, Payment activity, Time activity, All
    *   Export activity log to CSV/PDF
        
5.  **Collaboration tools:**
    
    *   In-workroom chat (separate from main inbox):
        *   Contextual to specific contract
        *   File attachments inline
        *   @ mentions
        *   Thread replies
    *   Shared notes/documents:
        *   Collaborative rich text editor (basic Google Docs-like)
        *   Markdown support
        *   Auto-save every 10 seconds
        *   Version history
    *   Task lists/checklists:
        *   Add tasks related to contract
        *   Assign to freelancer or self
        *   Set due dates
        *   Mark complete
    *   Calendar integration:
        *   Milestone deadlines
        *   Review meetings
        *   Payment schedules
            
6.  **Workroom settings:**
    
    *   Notification preferences for this contract
    *   File organization rules
    *   Activity feed filters
    *   Permissions (if team contract)
**Branches & Edge Cases:**
*   Large file uploads → chunked upload with progress bar, resume on failure
*   Storage quota → show remaining space, prompt to upgrade if full
*   Malware detected → file quarantined, both parties notified
*   Deleted files → move to trash, 30-day retention before permanent deletion
*   Version conflicts → if both parties edit simultaneously, last save wins (with conflict notification)
*   External file sharing → generate public link with expiry and password protection
*   File approval workflow → client can set required approval before accepting deliverables
*   OCR and text extraction → automatic for PDFs, searchable content
*   Mobile upload → camera/gallery integration, compress images to save data
*   Offline mode (mobile) → queue uploads, sync when reconnected
*   Access revocation → when contract ends, restrict file access (view-only for 90 days, then archive)
*   Backup and export → download entire workroom as ZIP
*   Integrations → sync with Google Drive, Dropbox, OneDrive (optional)
*   Watermarking → optional for sensitive files (e.g., designs)
*   File linking → reference files in messages/milestones by @filename
**Notifications:**
*   File uploaded by other party (configurable: real-time, digest)
*   Feedback requested on deliverable
*   File approved/rejected
*   Storage quota warning (80%, 95% full)
*   Virus detected in upload
**Analytics:**
*   workroom.viewed (contract\_id)
*   file.uploaded (file\_type, size, source)
*   file.downloaded
*   file.previewed
*   file.commented
*   file.approved|rejected
*   activity.viewed
*   activity.exported
*   workroom.settings\_updated
**Sources:** contracts-be.user-stories.md, storage-be.user-stories.md, storage-be.database-design.md, combined-fe-folder-strucure.md

### CL-7 — Timesheets (Hourly)

**Persona:** Client**Preconditions & Triggers:** Active hourly contract
**Primary Screens:**
*   **Web:** /app/(workroom)/\[contractId\]/time, /app/(billing)/timesheets
*   **Mobile:** /(tabs)/(authenticated)/(workroom)/\[contractId\]/time, /(tabs)/(authenticated)/(billing)/timesheets + time tracker widget
**System Touchpoints:** contracts-be, financial-be, communications-be
**Flow Steps:**

1.  Navigate to Time tab in workroom → /app/(workroom)/\[contractId\]/time
    
2.  View timesheet options:
    
    *   **Manual time entry** (freelancer logs hours)
    *   **Time tracker** (automatic tracking via desktop app or browser extension)
    *   **Hybrid** (tracker with manual adjustments)
        
3.  Weekly timesheet view:
    
    *   Current week shown by default (Mon-Sun)
    *   Navigate to previous weeks
    *   Each day shows:
        *   Hours logged (with memos)
        *   Time tracker screenshots (if enabled, 6 per hour, blurred by default)
        *   Work diary summary (apps used, activity level)
    *   Total hours for the week
    *   Billable amount calculation
    *   Status: Pending review / Approved / Disputed
        
4.  Review timesheet entries:
    
    *   Click day → see detailed breakdown
    *   View memos for each time entry (what was worked on)
    *   View tracker data:
        *   Screenshots (click to see full-size, user can delete sensitive screenshots)
        *   Activity levels (keyboard/mouse activity heatmap)
        *   Apps used (non-intrusive, grouped categories)
    *   Manual time entries (for offline work)
        
5.  Timesheet approval workflow:
    
    *   Automatic approval (if enabled in contract):
        *   Time approved automatically every Monday morning
        *   Payment processed immediately
    *   Manual approval:
        *   Client receives notification: "Timesheet ready for review"
        *   Deadline: 5 business days to review
        *   Actions:
            *   **Approve all:** Approve entire week
            *   **Approve selected:** Select specific hours to approve
            *   **Dispute hours:** Flag specific entries with reason
            *   **Request edits:** Ask freelancer to correct entries
                
6.  Dispute resolution:
    
    *   If hours disputed:
        *   Freelancer notified with reason
        *   Freelancer can provide clarification/evidence
        *   Client can reconsider or escalate to platform support
        *   Disputed hours held until resolved
            
7.  Payment processing:
    
    *   Approved hours billed immediately
    *   Payment debited from client's payment method or wallet
    *   Held in escrow for 5-day dispute window
    *   Released to freelancer after window
        
8.  Weekly limit monitoring:
    
    *   Contract has max hours per week (e.g., 40 hours)
    *   System shows hours remaining
    *   Notify client if limit approached (80%, 100%)
    *   Freelancer can request limit increase
    *   Overage handling:
        *   Block tracking if hard limit
        *   Allow with approval for soft limit
            
9.  Timesheet reports:
    
    *   Export timesheets to CSV/PDF
    *   Date range selection
    *   Summary: total hours, total cost, average hourly work
    *   Breakdown by day/week/month
    *   Charts: hours over time, activity levels, cost trends
**Branches & Edge Cases:**
*   No time logged for a week → reminder notification to freelancer and client
*   Screenshot privacy → freelancer can delete sensitive screenshots before submission
*   Activity level too low → flagged for review, but not auto-rejected
*   Timezone issues → all times displayed in client's timezone with original timezone noted
*   Manual adjustments → client can manually adjust hours with reason (requires freelancer agreement)
*   Tracker offline → manual entry required with explanation
*   Contract pause → stop billing, resume when contract reactivated
*   Weekly limit exceeded → system blocks tracking or requires approval
*   Payment method failure → retry logic, notify client, freeze contract if unresolved
*   Bonus hours → client can add bonus hours beyond logged time
*   Holidays → freelancer can mark days as non-working
*   Sick leave → freelancer can request paid/unpaid leave
*   Timesheet amendments → allow edits within 7 days of approval (both parties consent)
*   Audit trail → full history of timesheet edits, approvals, disputes
**Notifications:**
*   Timesheet submitted (to client)
*   Timesheet approved (to freelancer)
*   Timesheet disputed (to freelancer, with reason)
*   Hours approaching weekly limit (to freelancer, 80% and 100%)
*   Payment processed (to both parties)
*   Timesheet review deadline approaching (to client, 3 days before)
*   Auto-approval reminder (if enabled)
**Analytics:**
*   timesheet.submitted (contract\_id, hours, week)
*   timesheet.reviewed (time\_to\_review)
*   timesheet.approved (hours, amount)
*   timesheet.disputed (hours, reason)
*   timesheet.amended
*   timesheet.exported
*   weekly\_limit.approached (percentage)
*   tracker.screenshot\_deleted (count)
**Sources:** contracts-be.user-stories.md, financial-be.user-stories.md, combined-fe-folder-strucure.md

### CL-8 — Billing: Invoices, Payments, Refunds

**Persona:** Client**Preconditions & Triggers:** Active contracts with invoices generated
**Primary Screens:**
*   **Web:** /app/(billing)/invoices, /app/(billing)/payments, /app/(billing)/payment-methods
*   **Mobile:** /(tabs)/(authenticated)/(billing)/invoices, /(tabs)/(authenticated)/(billing)/payments, /(tabs)/(authenticated)/(billing)/payment-methods
**System Touchpoints:** financial-be, contracts-be, communications-be, subscriptions-be, storage-be
**Flow Steps:**

1.  **Invoices dashboard** → /app/(billing)/invoices
    
    *   View all invoices (past and upcoming)
    *   Filters: Status (Paid, Pending, Overdue, Draft), Date range, Contract
    *   List shows:
        *   Invoice number (unique ID)
        *   Date issued
        *   Due date
        *   Amount
        *   Status badge
        *   Quick actions: View, Download PDF, Pay now
            
2.  **View invoice details:**
    
    *   Invoice header: Client name, address, VAT ID
    *   Freelancer details: Name, address, tax ID
    *   Line items:
        *   Description (e.g., "Hourly work - Week of Jan 1")
        *   Quantity (hours or units)
        *   Rate
        *   Subtotal
    *   Subtotal, platform fees, taxes (VAT/GST), total
    *   Payment terms (Net 7, Net 14, Net 30)
    *   Payment status and history
    *   Download invoice as PDF
        
3.  **Payment processing:**
    
    *   For pending invoices → "Pay Now" button
    *   Select payment method:
        *   Primary payment method (pre-selected)
        *   Alternative payment methods
        *   Add new payment method
    *   Review payment summary:
        *   Amount breakdown (work amount + fees + taxes)
        *   Payment method
        *   Estimated processing time
    *   Confirm payment
    *   Payment processed:
        *   Immediate for card payments
        *   3-5 business days for ACH
    *   Receipt generated and emailed
        
4.  **Automatic payments (if enabled):**
    
    *   Invoices auto-paid on due date
    *   From primary payment method
    *   Notification sent after successful payment
    *   Retry logic if payment fails (3 attempts over 7 days)
    *   Contract paused if payment still fails
        
5.  **Payment methods management:**
    
    *   Navigate to /app/(billing)/payment-methods
    *   View all saved methods:
        *   Credit/debit cards (last 4 digits, expiry, brand)
        *   Bank accounts (last 4 digits, bank name)
        *   PayPal (email)
        *   Crypto wallets (address)
    *   Add new method:
        *   Card: enter details, tokenize via Stripe
        *   Bank: enter routing/account, verify via micro-deposits
        *   PayPal: OAuth link
        *   Crypto: generate wallet address
    *   Set default payment method
    *   Edit method (update expiry, billing address)
    *   Remove method (if not default and no pending payments)
        
6.  **Refunds:**
    
    *   Client requests refund (rare, usually via dispute)
    *   Navigate to specific invoice → "Request Refund" (if eligible)
    *   Select reason: Duplicate charge, Service not delivered, Overcharged, Other
    *   Provide details (required for Other)
    *   Submit request
    *   Admin reviews refund request (1-3 business days)
    *   If approved:
        *   Refund processed to original payment method
        *   5-10 business days for credit to appear
        *   Notification sent
    *   If denied:
        *   Reason provided
        *   Option to appeal or open dispute
            
7.  **Payment history:**
    
    *   View all payments made
    *   Filters: Date range, Amount range, Status, Contract
    *   Export to CSV/PDF for accounting
    *   Breakdown by month/quarter/year
    *   Charts: spending over time, by project/freelancer
        
8.  **Billing preferences:**
    
    *   Currency preference (display vs. payment)
    *   Invoice delivery: Email, In-app only
    *   Auto-pay settings: Enable/disable, grace period
    *   Payment reminders: Timing (3 days before, 1 day before, on due date)
    *   Receipt preferences: Email, download, both
**Branches & Edge Cases:**
*   Payment method expires → notification 30 days before, reminder to update
*   Payment failure → immediate notification, retry logic, suggest alternative method
*   Chargeback by client → admin investigates, may pause account pending resolution
*   Currency conversion → show amounts in both currencies (payment currency + display currency)
*   VAT/GST calculation → automatic based on client location and applicable thresholds
*   Reverse charge mechanism (EU B2B) → VAT exemption if client provides valid VAT ID
*   Invoice amendments → allow corrections within 24 hours of issuance
*   Duplicate invoices → system detects and flags potential duplicates
*   Overpayment → refunded automatically or credited to wallet
*   Partial payments → allow for large invoices, track payment installments
*   Payment plans → for large projects, split into installments
*   Late payment fees → configurable in contract terms, auto-added if enabled
*   Invoice disputes → client can flag invoice issues, pause payment until resolved
*   Multi-contract invoices → consolidate multiple contracts into one invoice
*   Tax documentation → download tax summaries (1099, VAT reports)
*   Wallet funding → alternative to direct payment, pre-fund wallet for faster transactions
**Notifications:**
*   Invoice issued (email with PDF attachment)
*   Invoice due soon (3 days, 1 day before)
*   Invoice overdue (on due date, then every 3 days)
*   Payment processed successfully
*   Payment failed (with retry information)
*   Refund approved
*   Payment method expiring soon
*   Auto-pay enabled/disabled confirmation
**Analytics:**
*   invoice.viewed (invoice\_id)
*   invoice.downloaded
*   payment.initiated (amount, method)
*   payment.completed (amount, method, duration)
*   payment.failed (reason, method)
*   payment\_method.added (type)
*   payment\_method.updated
*   payment\_method.removed
*   autopay.toggled (enabled/disabled)
*   refund.requested (amount, reason)
*   refund.approved|denied
**Sources:** financial-be.user-stories.md, financial-be-folder-structure.md, contracts-be.user-stories.md, combined-fe-folder-strucure.md

### CL-9 — Disputes & Resolution

**Persona:** Client**Preconditions & Triggers:** Contract dispute arises (quality issues, payment disputes, scope disagreements)
**Primary Screens:**
*   **Web:** /app/(contracts)/\[contractId\]/dispute, /app/(billing)/disputes
*   **Mobile:** /(tabs)/(authenticated)/(contracts)/\[contractId\]/dispute, /(tabs)/(authenticated)/(billing)/disputes (view-only, resolution requires desktop)
**System Touchpoints:** contracts-be, financial-be, admin-be, communications-be, storage-be
**Flow Steps:**

1.  **Open a dispute:**
    
    *   From contract screen → "Open Dispute" button
    *   Dispute categories:
        *   Work quality not as expected
        *   Deliverable not received / incomplete
        *   Payment issue (amount, timing)
        *   Scope disagreement
        *   Freelancer unresponsive
        *   Contract breach
        *   Other (describe)
    *   Select category and provide detailed description (500-5000 chars)
    *   Attach evidence:
        *   Screenshots, documents, correspondence
        *   Reference specific files/milestones/timesheets
        *   Upload from device or link from workroom files
    *   Set priority: Low, Medium, High, Urgent
    *   Submit dispute
        
2.  **Dispute created:**
    
    *   Dispute ID generated (e.g., DSP-2025-123456)
    *   Status: Open
    *   Both parties notified
    *   Contract may be paused (payments on hold, work paused)
    *   Assigned to mediation queue
        
3.  **Mediation period (7 days):**
    
    *   Both parties can communicate via dispute thread:
        *   Private messages visible only to parties and mediator
        *   Add evidence anytime
        *   Propose solutions
        *   Request specific actions
    *   Platform encourages resolution without escalation:
        *   Suggested compromise options
        *   AI-powered settlement recommendations based on similar cases
    *   Either party can propose settlement:
        *   Partial refund
        *   Revised deliverables
        *   Extended deadline
        *   Revised payment terms
    *   Other party can accept, counter, or reject
        
4.  **Admin mediation (if not resolved):**
    
    *   After 7 days, admin mediator assigned
    *   Mediator reviews:
        *   Dispute description and evidence
        *   Contract terms
        *   Communication history
        *   Work completed vs. expected
        *   Payment history
    *   Mediator may:
        *   Request additional information from either party
        *   Schedule call/video conference
        *   Review files with expert (e.g., code review, design review)
    *   Mediator proposes resolution:
        *   Full refund to client
        *   Partial refund
        *   Payment released to freelancer
        *   Work revision required
        *   Contract termination
        *   No action (dispute dismissed)
            
5.  **Resolution:**
    
    *   Both parties notified of mediator's decision
    *   3-day window to appeal decision
    *   If no appeal:
        *   Decision implemented automatically
        *   Payments adjusted accordingly
        *   Contract status updated
        *   Feedback/review rules updated (partial review allowed)
    *   If appeal:
        *   Escalated to senior admin
        *   Extended review (5-7 business days)
        *   Final decision binding
            
6.  **Post-resolution:**
    
    *   Dispute marked as Closed
    *   Outcome recorded:
        *   Resolved (mutual agreement)
        *   Decided (admin decision)
        *   Withdrawn (client withdrew dispute)
    *   Impact on users:
        *   Negative outcome may affect reputation (hidden from public, used for internal risk scoring)
        *   Frivolous disputes flagged
        *   Repeat offenders warned
            
7.  **Dispute history:**
    
    *   View all disputes → /app/(billing)/disputes
    *   Filter by: Status, Category, Date, Outcome
    *   See trends (if multiple disputes)
    *   Export dispute logs
**Branches & Edge Cases:**
*   Urgent disputes → expedited review (e.g., payment holds for large amounts)
*   Multiple disputes on same contract → consolidated into single case
*   Frivolous or bad-faith disputes → warning, potential account suspension
*   Dispute withdrawal → client can withdraw anytime before resolution, fees may apply
*   Evidence tampering → if detected, automatic decision against offending party
*   Escalation to legal arbitration → for high-value contracts, optional binding arbitration
*   Partial release → mediator can release partial payment while dispute continues
*   Mutual agreement during mediation → dispute closed immediately
*   Time zone for calls/meetings → scheduled in mutually convenient time
*   Language barriers → translation services available for international disputes
*   IP disputes → special handling, may involve legal team
*   Chargeback vs. dispute → if client files chargeback with bank, platform dispute process paused
*   Dispute after contract completion → allowed within 30 days of final payment
*   Reputation impact → disputes resolved in client's favor don't count against freelancer if deemed legitimate client concern
**Notifications:**
*   Dispute opened (to both parties)
*   Evidence added by other party
*   Settlement proposed
*   Mediator assigned
*   Mediator request for information
*   Mediation call scheduled
*   Resolution proposed (3-day appeal window starts)
*   Resolution finalized
*   Dispute closed
**Analytics:**
*   dispute.opened (contract\_id, category, priority)
*   dispute.evidence\_added (file\_count, file\_types)
*   dispute.settlement\_proposed (by\_party, terms)
*   dispute.settlement\_accepted|rejected
*   dispute.mediator\_assigned
*   dispute.resolution\_proposed (outcome)
*   dispute.appealed
*   dispute.closed (outcome, resolution\_time)
*   dispute.withdrawn
**Sources:** contracts-be.user-stories.md, admin-be.user-stories.md, financial-be.user-stories.md, combined-fe-folder-strucure.md

### CL-10 — Close, Review & Rehire

**Persona:** Client**Preconditions & Triggers:** Contract work completed or ready to terminate
**Primary Screens:**
*   **Web:** /app/(contracts)/\[contractId\]/close, /app/(reviews)/new, /app/(talent)/\[freelancerId\]
*   **Mobile:** /(tabs)/(authenticated)/(contracts)/\[contractId\]/close, /(tabs)/(authenticated)/(reviews)/new, /(tabs)/(authenticated)/(talent)/\[freelancerId\]
**System Touchpoints:** contracts-be, reviews-be, communications-be, financial-be, users-be
**Flow Steps:**

1.  **Close contract:**
    
    *   Navigate to contract → "Close Contract" button
    *   Close reasons:
        *   Work completed successfully
        *   Project cancelled
        *   Mutual agreement to end
        *   Freelancer unavailable
        *   Client no longer needs services
        *   Other (describe)
    *   Confirm outstanding items:
        *   All milestones approved or cancelled
        *   All timesheets approved or disputed
        *   All payments processed or pending
        *   No open disputes
    *   System checks:
        *   Any pending deliverables
        *   Unreleased funds in escrow
        *   Open feedback requests
    *   Review final costs:
        *   Total amount paid
        *   Hours worked or milestones completed
        *   Platform fees
        *   Any refunds/adjustments
    *   Confirm closure
        
2.  **Post-closure process:**
    
    *   Contract status changed to "Closed"
    *   Final invoice generated (if any pending amounts)
    *   Escrow released (after 5-day window)
    *   Workroom access: Read-only for both parties (90 days)
    *   Files archived but accessible for download
    *   Export option: Download all contract data (work logs, files, messages)
        
3.  **Review prompt (double-blind system):**
    
    *   After closure, both parties prompted to leave reviews
    *   Review window: 14 days (extendable once by 14 days)
    *   Reviews not visible until both parties submit or window closes
    *   Navigate to review form → /app/(reviews)/new?contractId=\[contractId\]
        
4.  **Leave review:**
    
    *   Overall rating: 1-5 stars (required)
    *   Category ratings (each 1-5 stars):
        *   Quality of work / Results delivered
        *   Communication responsiveness
        *   Professionalism
        *   Budget/Timeline adherence
        *   Would hire again (Yes/No, prominently displayed)
    *   Written feedback (100-5000 chars, required minimum 100):
        *   What went well
        *   Areas for improvement
        *   Specifics about deliverables
    *   Private feedback (optional, visible only to freelancer and platform):
        *   Constructive criticism
        *   Sensitive feedback not suitable for public
    *   Make review public or private (client's choice)
    *   Attachments (optional): Screenshots of delivered work (with consent)
    *   Submit review
        
5.  **Review publication:**
    
    *   If both reviews submitted within window:
        *   Both reviews published simultaneously
        *   No editing after publication
        *   Both parties notified
    *   If only one review submitted:
        *   Published after 14-day window expires
        *   Other party can still submit late review (marked as "Late")
    *   If neither submits:
        *   No reviews published
        *   Contract marked as "No feedback"
            
6.  **Review response (optional):**
    
    *   Freelancer can respond to client's review (once, within 30 days)
    *   Response character limit: 1000 chars
    *   Should be professional and constructive
    *   Displayed below client's review
    *   Cannot change or delete response after posting
        
7.  **Rehire freelancer:**
    
    *   From freelancer profile → "Hire Again" button
    *   Quick contract creation:
        *   Pre-filled with previous terms
        *   Same rate (or updated)
        *   Same contract type (or change)
        *   Reference previous contract
    *   Streamlined approval (no job posting needed)
    *   Freelancer receives direct offer
    *   Can accept/decline/negotiate
        
8.  **Save to favorites:**
    
    *   Add freelancer to favorites list
    *   Create custom lists (e.g., "Designers", "Emergency contacts")
    *   Quick access for future hiring
    *   Get notified of freelancer's availability
**Branches & Edge Cases:**
*   Close with disputed hours/milestones → must resolve first or accept admin decision
*   Early termination → may require mutual consent or per contract terms (notice period)
*   Final payment pending → contract can close but marked "Awaiting final payment"
*   Partial work completion → prorate payment, adjust review accordingly
*   Review editing → not allowed after submission, but can flag for removal if violates policies
*   Review disputes → flagging system for inappropriate reviews (admin reviews)
*   Anonymous reviews → not allowed, transparency enforced
*   Fake reviews → AI detection, manual review if flagged
*   Review incentives → prohibited, disclosure required if any compensation exchanged
*   Negative review appeal → freelancer can request review if factually incorrect or abusive
*   Late review submission → marked as late but still counts
*   Review without contract → not possible, requires closed contract
*   Mutual NDA → review may be limited or require approval before publication
*   Long-term contract → periodic reviews at milestones, final review at close
*   Auto-close → contracts auto-close after 90 days of inactivity with warnings
*   Tax reporting → closed contracts trigger tax document generation (e.g., 1099)
**Notifications:**
*   Contract closed confirmation (to both parties)
*   Review reminder (3 days, 7 days, last day of window)
*   Review submitted by you
*   Review received from other party
*   Both reviews published (viewable now)
*   Review responded to (if freelancer responds to your review)
*   Contract data export ready
**Analytics:**
*   contract.close\_initiated
*   contract.closed (reason, total\_amount, duration\_days)
*   review.prompt\_shown
*   review.started (party)
*   review.submitted (party, overall\_rating, would\_hire\_again)
*   review.published (time\_to\_publish)
*   review.response\_added
*   review.flagged (reason)
*   freelancer.saved\_to\_favorites
*   rehire.initiated (contract\_id\_reference)
**Sources:** contracts-be.user-stories.md, reviews-be.user-stories.md, reviews-be-folder-structure.md, combined-fe-folder-strucure.md

### CL-11 — Job Templates & Cloning

**Persona:** Client**Preconditions & Triggers:** Posted multiple similar jobs, wants to streamline process
**Primary Screens:**
*   **Web:** /app/(jobs)/templates, /app/(jobs)/\[jobId\]/save-as-template, /app/(jobs)/new?templateId=\[id\]
*   **Mobile:** /(tabs)/(authenticated)/(client)/jobs/templates, /(tabs)/(authenticated)/(client)/jobs/\[jobId\]/save-as-template, /(tabs)/(authenticated)/(client)/jobs/new?templateId=\[id\]
**System Touchpoints:** jobs-be, storage-be, users-be
**Flow Steps:**

1.  **Create template from existing job:**
    
    *   Navigate to successful completed job
    *   Click "Save as Template"
    *   Template creation form:
        *   Template name (e.g., "Web Developer - React Projects")
        *   Description (for internal use)
        *   What to save:
            *   Job description (required)
            *   Skills and categories (required)
            *   Screening questions
            *   Attachments (optional)
            *   Budget range guidance
            *   Visibility settings
        *   Privacy: Private (only me) / Shared (team members)
    *   Save template
        
2.  **Browse templates:**
    
    *   Navigate to /app/(jobs)/templates
    *   View all saved templates (personal + shared)
    *   Filter by: Category, Date created, Usage count
    *   Sort by: Most used, Recently used, A-Z
    *   Search templates
        
3.  **Use template:**
    
    *   From templates page → click "Use Template"
    *   OR from new job page → "Start from Template" dropdown
    *   Template loaded into job draft:
        *   All fields pre-filled
        *   Editable before publishing
        *   Update job-specific details (budget, timeline)
    *   Review and adjust
    *   Publish job
        
4.  **Clone existing job:**
    
    *   From any published/closed job → "Clone Job"
    *   Creates new draft with copied data
    *   Automatically appends "Copy" to title
    *   Review and modify
    *   Publish as new job (new job ID, fresh posting)
        
5.  **Template management:**
    
    *   Edit template:
        *   Update description, skills, questions
        *   Version history (optional)
    *   Delete template (with confirmation)
    *   Share template with team
    *   Track template usage stats:
        *   Times used
        *   Success rate (hire from templated jobs)
        *   Average proposals received
        *   Average time-to-hire
            
6.  **Smart suggestions:**
    
    *   System suggests templates when creating new job:
        *   Based on job title, description, category
        *   "You've posted similar jobs before. Use a template?"
    *   Suggest improvements to templates:
        *   Based on performance of jobs posted using template
        *   "Jobs from this template get 40% more proposals when budget is X−X-X−Y"
**Branches & Edge Cases:**
*   Template with outdated skills → system suggests refreshing skill tags
*   Shared template modified → notify team members using it
*   Template versioning → keep history of changes, rollback option
*   Template library → curated templates from successful clients (opt-in to share)
*   Template variables → use placeholders like {{project\_name}}, {{budget}} for quick customization
*   Import/export templates → JSON/CSV format for bulk management
*   Template categories → organize by project type, department, frequency
**Notifications:**
*   Template created
*   Template used by team member (if shared)
*   Template updated (if shared)
*   Template performance insights (monthly digest)
**Analytics:**
*   job\_template.created (from\_job\_id)
*   job\_template.used
*   job\_template.edited
*   job\_template.deleted
*   job\_template.shared
*   job.cloned (source\_job\_id)
*   template\_suggestion.shown
*   template\_suggestion.accepted
**Sources:** jobs-be.user-stories.md, jobs-be-folder-structure.md, combined-fe-folder-strucure.md

### CL-12 — Bulk Operations & Team Collaboration

**Persona:** Client (Team lead, hiring manager)**Preconditions & Triggers:** Managing multiple jobs or proposals, working with team
**Primary Screens:**
*   **Web:** /app/(jobs), /app/(jobs)/\[jobId\]/proposals, /app/(settings)/teams, /app/(jobs)/\[jobId\]/team
*   **Mobile:** limited (view-only for most operations, actions require web)
**System Touchpoints:** jobs-be, proposals-be, users-be, communications-be, admin-be
**Flow Steps:**

1.  **Bulk job actions:**
    
    *   Navigate to jobs dashboard → /app/(jobs)
    *   Select multiple jobs (checkboxes)
    *   Bulk actions available:
        *   Close all selected
        *   Archive all selected
        *   Export selected to CSV
        *   Change visibility (Public / Invite-only)
        *   Add team member access
        *   Apply tag/label
    *   Confirm action → processed in background for large batches
    *   Progress notification
        
2.  **Bulk proposal actions:**
    
    *   From job proposals list → select multiple proposals
    *   Bulk actions:
        *   Shortlist all
        *   Decline all (with optional bulk message)
        *   Archive all
        *   Export to CSV/PDF for offline review
        *   Send bulk message (same message to all selected)
        *   Apply tag (e.g., "Second Round")
    *   Confirmation modal with preview
    *   Execute action
        
3.  **Team collaboration on jobs:**
    
    *   Add team members to job:
        *   Navigate to /app/(jobs)/\[jobId\]/team
        *   Invite team members by email
        *   Assign roles:
            *   Owner: Full control
            *   Editor: Can edit job, review proposals, message candidates
            *   Reviewer: Can view proposals, add comments (no direct actions)
            *   Viewer: Read-only access
        *   Send invitations
    *   Team members receive invitation:
        *   Accept invitation
        *   Access job from their jobs list
        *   See "Team Job" badge
            
4.  **Collaborative proposal review:**
    
    *   Internal notes on proposals:
        *   Any team member can add private notes
        *   Notes visible only to team, not freelancer
        *   @ mention team members in notes for discussion
        *   Thread replies on notes
    *   Voting/rating system:
        *   Each team member can rate proposal (1-5 stars)
        *   Average team score shown
        *   Individual votes visible to team
    *   Approval workflow:
        *   Owner can require X approvals before moving to interview/hire
        *   Pending approvals shown
        *   Approve/reject with reason
            
5.  **Team dashboard:**
    
    *   View all jobs where you're a team member
    *   Filter by your role (Owner / Editor / Reviewer)
    *   See pending actions (approvals needed, proposals to review)
    *   Team activity feed
        
6.  **Communication management:**
    
    *   Shared inbox for team jobs:
        *   All messages visible to team members
        *   Any editor can respond
        *   Tag conversations for routing
    *   Internal team chat per job:
        *   Discuss candidates privately
        *   Strategy discussions
        *   Share notes and feedback
            
7.  **Permissions and access control:**
    
    *   Job owner can:
        *   Add/remove team members
        *   Change team member roles
        *   View all actions by team members
        *   Set approval requirements
    *   Audit trail:
        *   Track who made what changes
        *   View history of team member actions
        *   Export audit log
**Branches & Edge Cases:**
*   Bulk action failure → partial success, retry for failed items
*   Team member removal → revoke access immediately, notification sent
*   Conflicting team opinions → owner has final say, escalation process
*   Team size limits → based on subscription plan
*   External collaborators → temporary access for contractors/consultants
*   Team templates → save team structure for reuse across jobs
*   Permissions inheritance → new jobs inherit team from organization settings
*   Approval deadlines → auto-escalate if no response within X days
*   Slack/MS Teams integration → notify team in external tools
*   Mobile team actions → view and comment only, no approvals on mobile
**Notifications:**
*   Invited to team job
*   Team member added/removed
*   Proposal needs your review
*   Approval requested
*   Team decision made
*   Mentioned in internal note
*   Bulk action completed
**Analytics:**
*   bulk\_action.initiated (action\_type, item\_count)
*   bulk\_action.completed (success\_count, failure\_count)
*   team.member\_added (role)
*   team.member\_removed
*   proposal.note\_added (by\_team\_member)
*   proposal.vote\_added (score)
*   approval.requested
*   approval.responded (approved/rejected)
**Sources:** jobs-be.user-stories.md, proposals-be.user-stories.md, users-be.user-stories.md, combined-fe-folder-strucure.md

C) Freelancer (Winning Work & Delivery)
---------------------------------------

### FR-1 — Profile, Portfolio, Availability

**Persona:** Freelancer**Preconditions & Triggers:** Completed basic onboarding, wants to optimize profile
**Primary Screens:**
*   **Web:** /app/(profile)/edit, /app/(profile)/portfolio, /app/(profile)/availability, /app/(profile)/\[freelancerId\] (public view)
*   **Mobile:** /(tabs)/(authenticated)/(profile)/edit, /(tabs)/(authenticated)/(profile)/portfolio, /(tabs)/(authenticated)/(profile)/availability, /(tabs)/(authenticated)/(profile)/\[freelancerId\] (public view)
**System Touchpoints:** users-be, storage-be, search-be, reviews-be
**Flow Steps:**

1.  **Profile editing:**
    
    *   Navigate to /app/(profile)/edit
    *   Tabbed interface:
        *   **Overview:** Headline, summary, photo
        *   **Skills:** Add/remove skills, set proficiency
        *   **Experience:** Work history entries
        *   **Education:** Degrees, certifications
        *   **Portfolio:** Work samples
        *   **Availability:** Work schedule, vacation
    *   **Overview tab:**
        *   Profile photo: Upload, crop, preview (square, min 400x400px)
        *   Professional headline (80 chars): SEO-optimized, keyword-rich
        *   Professional summary (5000 chars): Rich text editor, markdown support
        *   Hourly rate: Set rate, adjust by project type
        *   Job preferences: Full-time, Part-time, Contract
        *   Languages: Add languages with proficiency levels
        *   Location: City, country, timezone (auto-detected, editable)
        *   Willing to relocate or travel (checkboxes)
    *   **Skills tab:**
        *   Add skills: Autocomplete from taxonomy
        *   Primary skills (max 15): Most important skills
        *   Secondary skills: Nice-to-have skills
        *   Proficiency levels: Beginner, Intermediate, Advanced, Expert
        *   Years of experience per skill
        *   Endorsements count (from clients/colleagues)
        *   Skill tests: Take tests to verify skills, badge displayed
    *   **Experience tab:**
        *   Add work experience entries:
            *   Job title (required)
            *   Company name
            *   Location (city, country or "Remote")
            *   Start date, end date (or "Current")
            *   Description (rich text, 500-2000 chars)
            *   Skills used (link to skill tags)
            *   Achievements/outcomes (quantify when possible)
        *   Reorder entries (drag and drop)
        *   Mark as featured
    *   **Education tab:**
        *   Add education entries:
            *   Institution name
            *   Degree type (Bachelor's, Master's, PhD, Certificate, etc.)
            *   Field of study
            *   Graduation year (or expected)
            *   Description (optional)
        *   Certifications:
            *   Certification name
            *   Issuing organization
            *   Issue date, expiry date (if applicable)
            *   Credential ID (link for verification)
            *   Certificate file upload
    *   **Portfolio tab:** (see detailed flow below)
    *   **Availability tab:** (see detailed flow below)
        
2.  **Portfolio management:**
    
    *   Navigate to /app/(profile)/portfolio
    *   Add portfolio item:
        *   Project title (required, 50-100 chars)
        *   Description (500-3000 chars, describe problem, solution, impact)
        *   Skills used (tag skills from profile)
        *   Project type: Client work, Personal project, Open source, Side project
        *   Role: Solo / Team (specify role if team)
        *   Completion date
        *   Project URL (if live)
        *   Media:
            *   Images (max 10 per project, 10MB each)
            *   Video (YouTube/Vimeo link or upload, max 5 min)
            *   Files (PDFs, design files, code samples, max 25MB)
        *   Case study (optional):
            *   Challenge (what was the problem)
            *   Solution (how you solved it)
            *   Results (quantified outcomes, e.g., "40% increase in conversions")
        *   Client testimonial (optional, can request from past clients)
        *   Visibility: Public / Clients only / Private
    *   Manage portfolio:
        *   Reorder items (featured first, drag-and-drop)
        *   Edit items
        *   Delete items
        *   Download individual items
        *   Share portfolio externally (generate public link)
    *   Portfolio analytics:
        *   Views per item
        *   Clicks on project URLs
        *   Most viewed items
            
3.  **Availability settings:**
    
    *   Navigate to /app/(profile)/availability
    *   Current status:
        *   Available now (actively looking)
        *   Available soon (within 1-4 weeks)
        *   Not available (off the market)
        *   Vacation mode (temporary, auto-expiry)
    *   Weekly availability:
        *   Hours available per week (0-40+ hours)
        *   Flexible / Fixed schedule
        *   Preferred working hours (timezone-aware)
    *   Work preferences:
        *   Project duration: Short-term (<3 months), Long-term (3-12 months), Ongoing
        *   Contract type: Hourly, Fixed-price, Retainer (select multiple)
        *   Minimum project budget
        *   Workload capacity: How many active projects (1-5+)
    *   Vacation/unavailability calendar:
        *   Mark dates as unavailable
        *   Recurring unavailability (e.g., weekends, holidays)
        *   Vacation mode: Pause all alerts and hide from search during dates
    *   Response time commitment:
        *   Set expected response time (e.g., within 24 hours, same day)
        *   Auto-reply message when unavailable
            
4.  **Profile optimization:**
    
    *   Profile strength indicator (0-100%):
        *   Photo: 10%
        *   Headline: 5%
        *   Summary: 15%
        *   Skills (min 5): 15%
        *   Portfolio (min 3 items): 20%
        *   Experience (min 2 entries): 10%
        *   Education: 5%
        *   Reviews (min 1): 10%
        *   Availability set: 5%
        *   Response rate: 5%
    *   Suggestions for improvement:
        *   "Add 2 more portfolio items to reach 100%"
        *   "Take a skill test to boost profile visibility by 20%"
        *   "Add a video introduction to stand out"
    *   SEO score:
        *   Keyword density in headline/summary
        *   Skills match high-demand jobs
        *   Suggestions to improve discoverability
            
5.  **Public profile preview:**
    
    *   View profile as clients see it
    *   Share profile link externally
    *   QR code for profile (useful for networking)
**Branches & Edge Cases:**
*   Profile photo rejection → if inappropriate, flagged by AI or reported
*   Plagiarism detection → AI scans summary/portfolio descriptions for copied content
*   Skill overload → warn if too many skills (>50), dilutes focus
*   Portfolio NDA → option to obscure client names or sensitive details
*   Video introduction → optional 60-second video, shows on profile
*   Profile verification badges → email, phone, payment method, ID verification
*   Profile visibility → public, limited (clients only), private (off search)
*   Profile export → download profile as PDF resume
*   Profile URL customization → custom vanity URL (e.g., skillsier.com/john-designer)
*   Languages → add language certifications (TOEFL, IELTS scores)
*   Portfolio collaboration → if team project, invite teammates to confirm collaboration
*   Dynamic availability → auto-update based on current workload
*   Work samples copyright → ensure you have rights to display work
**Notifications:**
*   Profile milestone reached (50%, 75%, 100% complete)
*   Profile viewed by client (if premium feature)
*   Portfolio item featured in search
*   Skill test available for your skills
*   Profile optimization tips (weekly digest)
**Analytics:**
*   profile.edit\_started
*   profile.section\_updated (section\_name)
*   profile.photo\_uploaded
*   profile.skill\_added
*   profile.experience\_added
*   portfolio.item\_added (type, media\_count)
*   portfolio.item\_viewed (by\_client)
*   availability.updated (status)
*   profile.optimization\_score (score)
*   profile.preview\_clicked
*   profile.shared (method)
**Sources:** users-be.user-stories.md, user-be-folder-structure.md, storage-be.user-stories.md, combined-fe-folder-strucure.md

### FR-2 — Job Discovery & Alerts

**Persona:** Freelancer**Preconditions & Triggers:** Active profile, looking for work
**Primary Screens:**
*   **Web:** /app/(market)/jobs, /app/(market)/saved, /app/(market)/alerts, /app/(market)/recommended
*   **Mobile:** /(tabs)/(authenticated)/(freelancer)/market/jobs, /(tabs)/(authenticated)/(freelancer)/market/saved, /(tabs)/(authenticated)/(freelancer)/market/alerts, /(tabs)/(authenticated)/(freelancer)/market/recommended + push notifications for alerts
**System Touchpoints:** search-be, jobs-be, users-be, communications-be, subscriptions-be
**Flow Steps:**

1.  **Job search:**
    
    *   Navigate to /app/(market)/jobs (main job marketplace)
    *   Search bar with autocomplete:
        *   Search by keywords (job title, skills, description)
        *   Search by job ID
    *   Filter panel (left sidebar on web, modal on mobile):
        *   **Category:** Web Development, Design, Writing, Marketing, etc.
        *   **Skills:** Multi-select with AND/OR logic
        *   **Job type:** Fixed-price / Hourly / Contract-to-hire
        *   **Experience level:** Entry / Intermediate / Expert
        *   **Budget range:** Min-Max sliders (adjusts to hourly or fixed based on type)
        *   **Duration:** < 1 month / 1-3 months / 3-6 months / 6+ months / Ongoing
        *   **Location:** Client location (optional, for timezone matching)
        *   **Remote:** Remote only / Hybrid / On-site
        *   **Posted:** Last 24 hours / Last 7 days / Last 30 days
        *   **Proposals received:** Less than 5 / 5-20 / 20-50 / 50+
        *   **Job features:** Verified payment / Top client / Previous hire from you
    *   Sorting options:
        *   Relevance (default, AI-based match score)
        *   Newest first
        *   Budget (high to low / low to high)
        *   Fewest proposals (best odds)
        *   Most proposals (indicates high demand)
    *   Applied filters shown as chips (removable)
        
2.  **Job search results:**
    
    *   List view (default) or Grid view toggle
    *   Each job card shows:
        *   Job title (bold, clickable)
        *   Budget (hourly range or fixed amount)
        *   Posted time (e.g., "2 hours ago")
        *   Client info: Name, location, timezone, rating, spend, hire rate
        *   Job excerpt (first 150 chars of description)
        *   Skills required (top 5 skills as tags)
        *   Proposals received count
        *   Badges: Verified payment, Payment verified, Great client, Previous client
        *   Quick actions: Save job, Submit proposal (opens proposal modal)
    *   Pagination or infinite scroll
    *   Match score indicator (0-100%, colored badge):
        - Skills match: Your skills vs required skills (weighted by proficiency)
        - Budget fit: Your hourly rate vs client budget
        - Experience level: Your experience vs job requirements
        - Past performance: Success rate in similar jobs
        - Location/timezone: Overlap with client timezone
        - Response time: Your typical response speed matches client expectations
        - Availability: Your current capacity vs job duration

3. **Job details view:**
   - Click job card → Navigate to /app/(market)/jobs/[jobId]
   - Full job description with rich formatting
   - Client profile sidebar:
     - Client name, location, company size
     - Member since date
     - Total spend on platform
     - Jobs posted count
     - Open jobs count
     - Hire rate (% of jobs that resulted in hire)
     - Average rating from freelancers
     - Payment verified badge
     - Total hours hired
   - Skills required (all skills listed, not just top 5)
   - Job preferences:
     - Project duration
     - Time commitment (part-time/full-time)
     - Timezone overlap requirements
   - Budget details:
     - Hourly: Rate range (e.g., $50-$75/hr)
     - Fixed: Total budget (e.g., $5,000)
     - Weekly limit (for hourly jobs)
   - Screening questions preview (if any)
   - Proposals received count and distribution chart
   - Average proposal amount
   - Job activity timeline:
     - Posted date
     - Last viewed by client
     - Invitations sent count
     - Shortlisted candidates count
   - Similar jobs (based on skills/category)
   - Quick actions:
     - Submit proposal (prominent CTA)
     - Save job
     - Share job (copy link, email)
     - Report job (flag inappropriate)
   - Breadcrumbs: Market > Jobs > [Category] > [Job Title]

4. **Save job:**
   - Click "Save job" icon/button
   - Job saved to /app/(market)/saved
   - Confirmation toast: "Job saved. View in Saved Jobs"
   - Saved jobs are synced across devices
   - Can organize saved jobs into folders/collections (optional feature)
   - Auto-expire saved jobs after 90 days if job is closed

5. **Job recommendations:**
   - Navigate to /app/(market)/recommended
   - AI-powered job recommendations based on:
     - Profile skills and experience
     - Past proposals and accepted contracts
     - Search history and saved jobs
     - Jobs you've viewed/saved
     - Similar freelancers' activity
   - Each recommendation shows:
     - Match score (0-100%)
     - Match reasons: "Matches your React & Node.js skills"
     - "Client has hired 15 freelancers with your skills"
     - "Budget aligns with your typical rate"
   - Feedback buttons:
     - "Good match" (thumbs up)
     - "Not relevant" (thumbs down) → prompts for reason
   - Refresh recommendations daily
   - Can dismiss recommendations (won't show again)

6. **Job alerts setup:**
   - Navigate to /app/(market)/alerts
   - Create new alert:
     - Alert name (e.g., "React Senior Jobs")
     - Search query + filters (same as job search)
     - Alert frequency:
       - Real-time (push notification as posted)
       - Daily digest (8 AM local time)
       - Weekly digest (Monday 8 AM)
     - Delivery channels:
       - Email (toggleable)
       - Push notifications (toggleable)
       - In-app notifications (always on)
   - Alert preview: "You'll receive ~5-10 jobs per week"
   - Manage existing alerts:
     - Edit alert (change filters, frequency)
     - Pause alert (temporary disable)
     - Delete alert
   - Alert performance stats:
     - Total jobs received
     - Jobs you viewed from alert
     - Jobs you saved from alert
     - Jobs you applied to from alert
     - Click-through rate

7. **Receiving job alerts:**
   - **Real-time alerts:**
     - Push notification: "New job matches your alert: [Job Title]"
     - Tap notification → opens job details in app
     - In-app notification bell icon shows red dot
   - **Daily/Weekly digests:**
     - Email with job summary cards
     - Click job → opens in web app
     - In-app: Notification with count "5 new jobs from your alerts"
   - Alert notification includes:
     - Job title
     - Budget
     - Client rating
     - Match score
     - Posted time

8. **Saved searches:**
   - From job search results, click "Save this search"
   - Name the search (e.g., "Remote React Jobs")
   - Option to enable alerts on saved search
   - Access saved searches from /app/(market)/saved-searches
   - Quick access: Run saved search with one click
   - Edit saved search filters
   - Duplicate saved search to create variations

9. **Trending & featured:**
   - /app/(market)/trending shows:
     - Trending jobs (high proposal activity)
     - Trending skills (most demanded this week)
     - Featured jobs (promoted by clients)
     - Rising clients (new clients with good budgets)
   - "What's hot" section on main /app/(market)/jobs page
   - Trending analytics:
     - Skill demand change (+20% this week)
     - Average budget trends
     - Popular job categories

**Branches & Edge Cases:**

- **No results:** Show "No jobs match your filters. Try broadening your search" + suggested filters to remove
- **Saved job limit:** Max 500 saved jobs, warn at 450, prevent at 500 → "Remove old jobs to save new ones"
- **Alert limit:** Max 20 active alerts, warn at 15
- **Duplicate alerts:** System detects if creating alert with same filters → "You already have a similar alert"
- **Expired jobs:** Saved jobs that are closed show "This job is no longer accepting proposals"
- **Job removed:** If saved job is deleted by client → notification "A saved job was removed by the client"
- **Search syntax:** Support boolean operators (AND, OR, NOT) in advanced search
- **Saved search conflicts:** If filters become invalid (e.g., category removed), alert user to update
- **Zero-proposals jobs:** Highlight "Be the first to apply" for jobs with 0 proposals
- **High-competition warnings:** "This job has 50+ proposals. Consider other opportunities"
- **Budget mismatch:** If job budget < your min rate, show warning
- **Client reputation:** Flag low-rated clients with warning tooltip
- **Timezone mismatch:** Warn if client requires timezone overlap you can't meet
- **Required skills missing:** Highlight skills you don't have on profile
- **Premium job badges:** Featured/promoted jobs have special badge
- **Job urgency:** "Client needs to hire within 3 days" urgency indicator
- **Proposal deadline:** Show countdown timer if job has application deadline
- **Invite-only jobs:** Some jobs are invite-only, show "Not accepting public proposals"
- **Agency-only jobs:** Jobs restricted to agencies, show appropriate badge
- **Verification required:** Some jobs require verified payment method to apply

**Notifications:**

- New job alert (real-time/digest)
- Saved job closing soon (48 hours before close)
- Saved job opened again (if reposted)
- Job invitation received from client
- Recommended job matches your profile (daily)
- Trending skill alert: "React is trending this week, 200+ jobs posted"
- Search alert paused due to no activity (after 90 days)
- Search alert performance: "Your alert generated 5 proposals last month"

**Analytics:**

- market.jobs_searched (query, filters, result_count)
- market.job_viewed (job_id, from_source, match_score)
- market.job_saved (job_id, from_source)
- market.job_unsaved (job_id)
- market.recommendation_viewed (job_id, match_score, rank)
- market.recommendation_feedback (job_id, feedback_type, reason)
- market.recommendation_applied (job_id, match_score)
- market.alert_created (alert_name, filters, frequency)
- market.alert_edited (alert_id, changes)
- market.alert_deleted (alert_id, reason)
- market.alert_received (alert_id, job_count)
- market.alert_clicked (alert_id, job_id)
- market.saved_search_created (search_name, filters)
- market.saved_search_executed (search_id, result_count)
- market.trending_viewed (trending_type)
- market.filter_applied (filter_name, filter_value)
- market.filter_removed (filter_name)
- market.sort_changed (sort_option)
- market.search_refined (refine_action)
- market.zero_results_search (query, filters)

**Sources:** jobs-be.user-stories.md, search-be.user-stories.md, search-be-folder-structure.md, combined-fe-folder-strucure.md

---

### FR-3 — Proposal Compose & Connects

**Persona:** Freelancer

**Preconditions & Triggers:** Found a job, ready to apply

**Primary Screens:**
- **Web:** /app/(proposals)/new/[jobId], /app/(proposals)/drafts, /app/(proposals), /app/(proposals)/[proposalId]
- **Mobile:** Same routes with mobile-optimized editor

**System Touchpoints:** proposals-be, jobs-be, users-be, financial-be, communications-be, storage-be

**Flow Steps:**

1. **Initiate proposal:**
   - From job details (/app/(market)/jobs/[jobId]), click "Submit Proposal"
   - OR from job search results, click "Quick Apply" → opens proposal modal
   - System checks:
     - Connects balance (deduct 2-6 connects based on job tier)
     - If insufficient connects → prompt to purchase
     - Profile completeness (min 60% required)
     - If job has screening questions → must answer first
   - Navigate to /app/(proposals)/new/[jobId]

2. **Proposal editor:**
   - **Cover letter section:**
     - Rich text editor with formatting (bold, italic, lists, links)
     - Character count (min 100, max 5000 characters)
     - AI writing assistant (optional, premium feature):
       - "Improve this proposal" → suggestions
       - "Check grammar" → corrections
       - "Make more concise" → shortened version
     - Template library:
       - "Use from past proposals" → load previous cover letters
       - Built-in templates: Introduction, Skills showcase, Questions, Closing
       - Save as template for future use
   - **Bid amount:**
     - Hourly rate: Enter your rate (pre-filled from profile)
     - Fixed price: Enter your bid amount
     - Show suggested bid range based on:
       - Your profile rate
       - Market average for similar jobs
       - Client's budget
     - Bid too low warning: "This is 40% below market rate"
     - Bid too high warning: "This exceeds client's budget by 50%"
   - **Project duration:**
     - Estimate: < 1 week / 1-4 weeks / 1-3 months / 3-6 months / 6+ months
     - For hourly: hours per week commitment
     - Milestone suggestions (for fixed-price):
       - Suggest milestones with amounts
       - "Milestone 1: Design mockups - $500"
       - "Milestone 2: Development - $2000"
   - **Availability:**
     - Start date: When can you start? (date picker)
     - Current workload: X hours/week available
     - Auto-fill from profile availability settings
   - **Attachments:**
     - Drag & drop files (max 10 files, 25 MB each)
     - Supported: PDF, DOC, PPT, images, videos
     - Portfolio items: Quick attach from your portfolio
     - Past work samples relevant to job
   - **Screening questions** (if required by job):
     - Display all screening questions from job
     - Text answers (short/long), multiple choice, file upload
     - Required questions marked with *
     - Character limits enforced

3. **Proposal AI analysis (optional feature):**
   - Click "Analyze proposal" button
   - AI scans your proposal and provides:
     - Match score vs job requirements (0-100%)
     - Missing keywords from job description
     - Tone analysis: Professional / Casual / Overly formal
     - Suggestions:
       - "Mention your experience with React mentioned in the job"
       - "Add specific metrics: 'Increased performance by 40%'"
       - "Your bid is 20% higher than average, consider explaining value"
     - Estimated win probability based on proposal quality

4. **Connects usage:**
   - Display connects required (varies by job tier):
     - Entry-level job: 2 connects
     - Intermediate: 4 connects
     - Expert: 6 connects
     - Featured/Premium jobs: 10+ connects
   - Current balance displayed: "You have 50 connects remaining"
   - Option to boost proposal (premium):
     - Spend extra connects to feature proposal (appears higher)
     - Costs 5-10 additional connects
   - Purchase connects:
     - If insufficient, click "Buy Connects" button
     - Opens /app/(billing)/connects/purchase
     - Packages: 10 connects ($1), 40 connects ($3), 80 connects ($5)
     - After purchase, return to proposal editor

5. **Save as draft:**
   - Click "Save Draft" → saves to /app/(proposals)/drafts
   - Auto-save every 30 seconds
   - Draft indicator: "Last saved 2 minutes ago"
   - Drafts expire after 90 days
   - Can resume drafts later

6. **Submit proposal:**
   - Click "Submit Proposal" button
   - Final validation:
     - Cover letter min 100 chars
     - Bid amount entered
     - All required screening questions answered
     - Sufficient connects balance
   - Deduct connects from balance
   - Create proposal in proposals-be
   - Send notification to client (communications-be)
   - Redirect to /app/(proposals)/[proposalId] (submitted state)
   - Success message: "Proposal submitted! The client will be notified."
   - Cannot edit proposal after submission (by default, some jobs allow edits within 24h)

7. **Proposal sent view:**
   - /app/(proposals)/[proposalId] shows:
     - Proposal status: "Submitted" with timestamp
     - Full proposal details (read-only)
     - Client activity:
       - "Client viewed your proposal 2 hours ago"
       - "Client is actively hiring (viewed 15+ proposals)"
     - Similar proposals count: "12 other freelancers applied"
     - Actions:
       - Message client (if allowed by job settings)
       - Withdraw proposal
       - Report job/client
   - Proposal ranking (if visible):
     - "Your proposal is ranked #3 of 12" (based on client's sorting)

8. **Manage proposals:**
   - /app/(proposals) shows all proposals with tabs:
     - Active (pending client review)
     - Interviewing (client expressed interest)
     - Archived (withdrawn/declined/expired)
   - Each proposal card shows:
     - Job title
     - Bid amount
     - Submitted date
     - Proposal status
     - Client activity indicator
     - Quick actions: View, Message, Withdraw
   - Filters: Status, Date range, Job category, Bid amount range
   - Sort: Newest, Oldest, Most active (client engagement)

9. **Withdraw proposal:**
   - From /app/(proposals)/[proposalId], click "Withdraw Proposal"
   - Confirmation modal: "Are you sure? You won't get your connects back"
   - Select reason (optional): No longer interested / Found other work / Budget too low
   - Confirm withdrawal
   - Connects are NOT refunded
   - Proposal moved to Archived tab
   - Client is notified of withdrawal

10. **Proposal expiration:**
    - Proposals auto-expire after 30 days if no client response
    - Notification: "Your proposal for [Job Title] has expired"
    - Expired proposals moved to Archived
    - Can resubmit if job still open (costs connects again)

**Branches & Edge Cases:**

- **Insufficient connects:** Prompt to purchase connects before allowing submission
- **Purchase connects flow:** Inline purchase modal, complete payment, return to proposal editor
- **Job closed during drafting:** Alert "This job is no longer accepting proposals"
- **Rate below minimum:** If your bid < your profile minimum rate → warning
- **Duplicate proposal:** Cannot apply to same job twice (system prevents)
- **Screening question errors:** Highlight unanswered required questions
- **Attachment upload errors:** Retry mechanism, show error for unsupported formats
- **Proposal character limits:** Hard stop at 5000 characters, show warning at 4800
- **Auto-save failures:** "Could not save draft. Check your connection." with retry button
- **Connects purchase timeout:** If payment processing delayed, allow draft save, complete purchase later
- **Job invitation + proposal:** If invited, proposal costs 0 connects
- **Agency proposals:** Agency members can submit on behalf of team, select team member for project
- **Premium proposal boost:** Option to pay for higher placement in client's proposal list
- **Portfolio auto-attach:** Suggest relevant portfolio items based on job skills
- **Video proposals:** Some jobs accept video pitches (60-90 seconds)
- **Template reuse:** Load successful past proposals as templates
- **Bid calculator:** Help estimate fixed-price bid based on hourly rate × estimated hours
- **Proposal analytics:** Track view rate, response rate, win rate by proposal type
- **Connect bonus:** Occasionally earn bonus connects for profile completion milestones
- **Connect refund:** If job is removed by client within 48h, connects refunded
- **Proposal edits:** Some clients allow proposal edits within 24h of submission
- **Urgent proposals:** Jobs with urgent deadlines show countdown timer
- **Bulk proposals:** Premium feature: apply to multiple similar jobs with one proposal
- **Proposal feedback:** After rejection, client can provide optional feedback
- **Blacklisted clients:** System warns if client has low rating or payment issues

**Notifications:**

- Connects purchased successfully
- Proposal submitted confirmation
- Client viewed your proposal
- Client shortlisted your proposal
- Client sent you a message about proposal
- Client hired someone else for job
- Proposal expired (30 days, no response)
- Job closed (no hire made)
- Connects balance low (< 10 connects remaining)
- Proposal draft reminder (unfinished draft after 7 days)
- Proposal boost expiring soon (if used boost feature)

**Analytics:**

- proposals.draft_started (job_id, from_source)
- proposals.draft_saved (job_id, completion_percentage)
- proposals.ai_analysis_used (job_id)
- proposals.ai_suggestion_accepted (suggestion_type)
- proposals.template_used (template_id)
- proposals.attachment_added (file_type, file_size)
- proposals.screening_question_answered (question_id, answer_length)
- proposals.bid_amount_entered (job_id, bid_amount, market_comparison)
- proposals.connects_purchase_initiated (package_size)
- proposals.connects_purchased (package_size, payment_method)
- proposals.proposal_submitted (job_id, bid_amount, connects_spent, proposal_length)
- proposals.proposal_boosted (job_id, boost_type, connects_spent)
- proposals.proposal_viewed (proposal_id)
- proposals.proposal_withdrawn (proposal_id, reason)
- proposals.proposal_expired (proposal_id, days_since_submitted)
- proposals.client_activity_viewed (proposal_id)

**Sources:** proposals-be.user-stories.md, proposals-be-folder-structure.md, jobs-be.user-stories.md, financial-be.user-stories.md, combined-fe-folder-strucure.md

---

### FR-4 — Messaging & Interviewing

**Persona:** Freelancer (also applies to Client)

**Preconditions & Triggers:** Client interested in proposal, wants to interview

**Primary Screens:**
- **Web:** /app/(inbox), /app/(inbox)/[conversationId], /app/(inbox)/archived, /app/(inbox)/unread
- **Mobile:** Same routes with mobile-optimized messaging UI

**System Touchpoints:** communications-be, users-be, proposals-be, jobs-be, contracts-be, storage-be

**Flow Steps:**

1. **Client initiates message:**
   - Client views freelancer proposal, clicks "Message Freelancer"
   - Creates new conversation in communications-be
   - Freelancer receives notification:
     - Push: "New message about [Job Title]"
     - Email: "A client wants to discuss your proposal"
     - In-app: Red dot on inbox icon

2. **Access inbox:**
   - Navigate to /app/(inbox)
   - Left sidebar shows conversation list:
     - Sorted by most recent activity
     - Each conversation shows:
       - Participant name & avatar
       - Last message preview (50 chars)
       - Timestamp (e.g., "2 hours ago", "Yesterday")
       - Unread indicator (bold, blue dot)
       - Job title context (if related to job/proposal)
       - Attachment indicator icon
   - Tabs:
     - All messages
     - Unread (only unread conversations)
     - Archived (archived conversations)
   - Search box: Search messages by participant, content, job title

3. **Open conversation:**
   - Click conversation from list → opens in /app/(inbox)/[conversationId]
   - Conversation view (right pane or full-screen on mobile):
     - **Conversation header:**
       - Participant name, avatar, online status (green dot if online)
       - Job title (if conversation is linked to job/proposal)
       - Quick actions:
         - Video call button (if available)
         - Phone call button (if available)
         - Archive conversation
         - Mark as unread
         - Report/Block user
         - Conversation settings (mute notifications)
     - **Message thread:**
       - Chronological messages (oldest at top, newest at bottom)
       - Each message shows:
         - Sender avatar & name
         - Message text (supports markdown: **bold**, *italic*, `code`, links)
         - Timestamp (hover for full date/time)
         - Read receipts (double checkmark if read by recipient)
         - Edit/Delete icons (own messages only, within 5 min)
         - React to message: 👍 ❤️ 😂 🎉 (emoji reactions)
       - System messages:
         - "Proposal submitted" (with link to proposal)
         - "Contract created" (with link to contract)
         - "Payment released" (with amount)
       - Message status indicators:
         - Sending (spinner)
         - Sent (single checkmark)
         - Delivered (double checkmark)
         - Read (double checkmark, blue)
         - Failed (red exclamation, retry button)

4. **Compose message:**
   - **Message input field** (bottom of conversation):
     - Rich text editor (optional toolbar):
       - Bold, Italic, Code, Link
       - Bullet list, Numbered list
       - @ mentions (if multi-party conversation)
     - Character counter (max 10,000 characters)
     - Emoji picker button
     - Attach files button:
       - Click to open file picker
       - Drag & drop files directly into message field
       - Supported: documents, images, videos (max 25 MB per file, 10 files per message)
       - Show upload progress bar
       - Thumbnail preview for images
     - Voice message button (mobile):
       - Hold to record, release to send
       - Max 2 minutes
       - Waveform preview
   - Type message, press Enter to send (Shift+Enter for new line)
   - OR click "Send" button

5. **Send message:**
   - Message sent to communications-be
   - Optimistic UI: message appears immediately with "sending" status
   - On success: status → "sent"
   - On failure: show error, allow retry
   - Recipient receives push notification (if online push enabled)
   - Email notification sent (if offline > 1 hour)
   - Message read receipts update when recipient reads

6. **File attachments:**
   - Click "Attach files" button OR drag & drop into message field
   - Select files from device (documents, images, videos)
   - Upload to storage-be (virus scan, signed URLs)
   - Show upload progress (percentage, cancel option)
   - Thumbnail preview for images/videos
   - File metadata: name, size, type
   - Recipient can download/view files:
     - Click to download
     - Images/videos open in lightbox viewer
     - Documents open in browser (if PDF) or download

7. **Voice messages (mobile):**
   - Tap & hold microphone icon to record
   - Release to send, slide to cancel
   - Max duration: 2 minutes
   - Waveform visualization during recording
   - Play back before sending (optional)
   - Sent as audio file attachment
   - Recipient sees play button, waveform, duration
   - Auto-transcription (optional, premium feature)

8. **Video/Audio calls (optional feature):**
   - Click video/audio call button in conversation header
   - Sends call invitation message to recipient
   - Recipient receives push notification: "Incoming call from [Name]"
   - Accept/Decline buttons
   - If accepted → opens video/audio call screen:
     - WebRTC-based call (uses communications-be signaling)
     - Screen sharing option
     - Mute/Unmute, Video on/off
     - Chat sidebar during call
     - Record call option (with consent)
     - End call button
   - Call duration tracked, shown in conversation timeline
   - Call recording saved to conversation (if recorded)

9. **Scheduling interviews:**
   - **Option 1: Inline scheduling**
     - Type "/schedule" in message to trigger scheduling assistant
     - Propose times: "I'm available Tuesday 2-4 PM or Wednesday 10 AM-12 PM"
     - Recipient can click suggested times to accept
     - Calendar integration: creates calendar event automatically
   - **Option 2: Calendar link**
     - Share calendar availability link (via /app/(settings)/calendar/availability)
     - Recipient clicks link, books time slot
     - Both parties receive calendar invite
   - **Option 3: Third-party integration**
     - Connect Calendly, Cal.com, or Google Calendar
     - Share scheduling link in message
     - Booking syncs back to Skillsier

10. **Interview best practices (in-app tips):**
    - First-time interview tooltip shows tips:
      - "Ask about project scope and timeline"
      - "Clarify budget and payment terms"
      - "Discuss communication preferences"
      - "Request portfolio review or skills demonstration"
    - Sample interview questions provided (for both parties)
    - Post-interview checklist: "Did you discuss...?"

11. **Archive conversation:**
    - Click "Archive" in conversation header
    - Conversation moved to /app/(inbox)/archived
    - No longer appears in main inbox
    - Still searchable, can unarchive anytime
    - Useful for completed projects or inactive chats

12. **Search messages:**
    - Global search: /app/(inbox) search bar → search across all conversations
    - In-conversation search: Find specific messages within conversation
    - Search by: keyword, participant name, job title, date range, attachment type
    - Highlights matching text
    - Jump to message in thread

13. **Mute/Unmute notifications:**
    - Click conversation settings → "Mute notifications"
    - Options: 1 hour / 8 hours / 1 day / Until I turn back on
    - Muted conversations show mute icon
    - Still receive messages, but no push/email notifications

14. **Block/Report user:**
    - Click conversation settings → "Report user"
    - Select reason: Spam / Harassment / Inappropriate / Other
    - Optionally block user (cannot message you anymore)
    - Submit report to admin-be for review
    - Blocked users see "Message could not be delivered" error

15. **Message reactions:**
    - Hover over message → click emoji icon
    - Select emoji reaction (👍 ❤️ 😂 😮 😢 🎉)
    - Reaction appears below message with count
    - Multiple users can react with same emoji (count increments)
    - Click reaction to remove your reaction

16. **Edit/Delete messages:**
    - Own messages can be edited/deleted within 5 minutes of sending
    - Click "..." menu on message → Edit / Delete
    - Edit: inline editor, save changes → shows "(edited)" indicator
    - Delete: confirmation modal → "Delete for me" or "Delete for everyone"
    - Deleted messages show "This message was deleted"

17. **Threads/Replies (optional feature):**
    - Hover over message → click "Reply in thread"
    - Opens thread sidebar
    - Threaded replies keep conversations organized
    - Thread indicator on parent message: "3 replies"
    - Threads are private to original conversation participants

18. **Typing indicators:**
    - When recipient is typing, show "Typing..." indicator at bottom of message list
    - Shows participant name if multi-party: "John is typing..."
    - Disappears after 5 seconds of inactivity

19. **Message delivery & read receipts:**
    - Sent: Single grey checkmark (message delivered to server)
    - Delivered: Double grey checkmark (message delivered to recipient's device)
    - Read: Double blue checkmark (recipient opened conversation and saw message)
    - Can disable read receipts in settings (but won't see others' receipts either)

20. **Offline behavior (mobile):**
    - Messages queue locally if offline
    - Show "Sending..." status
    - Auto-send when connection restored
    - Download recent messages for offline reading
    - Indicate offline mode with banner: "You're offline. Messages will send when reconnected."

**Branches & Edge Cases:**

- **First message icebreaker:** Suggest conversation starters: "Hi [Name], I reviewed your proposal..."
- **Spam detection:** If user sends identical messages to multiple people → flag for review
- **Message limits:** Max 50 messages/hour to prevent spam
- **File upload failures:** Retry mechanism, show error if virus detected or file too large
- **Voice message permissions:** Request microphone permission on first use (mobile)
- **Video call bandwidth:** Warn if low bandwidth detected, suggest audio-only
- **Screen share permissions:** Request screen capture permission (desktop)
- **Calendar sync errors:** Show error if calendar integration fails, offer manual scheduling
- **Timezone confusion:** Show participant's timezone in conversation header: "John (EST)"
- **Auto-translate:** Optional feature: auto-translate messages to your preferred language
- **Message retention:** Messages deleted after 5 years (configurable in settings)
- **Group conversations:** Support multi-party conversations (e.g., client + multiple freelancers)
- **@ mentions:** In group chats, mention specific participants to notify them
- **Message pinning:** Pin important messages to top of conversation
- **Conversation labels:** Tag conversations with labels (e.g., "Urgent", "Follow-up", "Interview")
- **Smart replies:** AI-suggested quick replies: "Thanks!", "Sounds good", "Let me check"
- **Message formatting:** Support markdown, code blocks, quotes
- **Link previews:** Auto-generate preview cards for shared URLs (title, image, description)
- **Contract context:** Show contract details sidebar if conversation is linked to active contract
- **Proposal context:** Show proposal summary sidebar if conversation is about proposal
- **Unread counter:** Show unread message count on inbox icon (max 99+)
- **Desktop notifications:** Browser push notifications when app is in background (opt-in)
- **Email notifications:** Send email digest if unread messages after 1 hour offline
- **Read receipt privacy:** Option to disable read receipts for your messages
- **Message search filters:** Filter search by date, participant, has attachments, job related
- **Export conversation:** Download conversation history as PDF or text file
- **Conversation analytics:** Track response time, message count, engagement

**Notifications:**

- New message received (push, email if offline > 1h)
- Mentioned in message (@ mention)
- Message read by recipient (optional)
- Voice message received
- File shared in conversation
- Video/audio call incoming
- Call missed
- Scheduled interview reminder (15 min before)
- Conversation archived
- User typing (live indicator, not notification)
- Message failed to send (retry prompt)
- Large attachment upload completed
- Conversation muted/unmuted

**Analytics:**

- inbox.conversation_opened (conversation_id, from_source)
- inbox.message_sent (conversation_id, message_length, has_attachment, has_emoji)
- inbox.message_edited (message_id, time_since_sent)
- inbox.message_deleted (message_id, time_since_sent, delete_type)
- inbox.message_reacted (message_id, emoji)
- inbox.attachment_uploaded (file_type, file_size, upload_duration)
- inbox.attachment_downloaded (file_id)
- inbox.voice_message_recorded (duration)
- inbox.voice_message_played (message_id)
- inbox.call_initiated (conversation_id, call_type)
- inbox.call_answered (call_id, response_time)
- inbox.call_declined (call_id)
- inbox.call_ended (call_id, duration)
- inbox.screen_share_started (call_id)
- inbox.conversation_archived (conversation_id)
- inbox.conversation_muted (conversation_id, duration)
- inbox.user_blocked (user_id, reason)
- inbox.user_reported (user_id, reason, report_category)
- inbox.search_performed (query, result_count)
- inbox.typing_indicator_shown (conversation_id)
- inbox.read_receipt_sent (message_id)
- inbox.interview_scheduled (conversation_id, method)
- inbox.offline_messages_queued (count)
- inbox.offline_messages_sent (count, delay)

**Sources:** communications-be.user-stories.md, communications-be-folder-structure.md, users-be.user-stories.md, contracts-be.user-stories.md, storage-be.user-stories.md, combined-fe-folder-strucure.md

---

### FR-5 — Accept Contract & Start Work

**Persona:** Freelancer

**Preconditions & Triggers:** Client sends offer/contract after interview

**Primary Screens:**
- **Web:** /app/(contracts)/pending, /app/(contracts)/[contractId], /app/(contracts)/[contractId]/accept
- **Mobile:** Same routes

**System Touchpoints:** contracts-be, users-be, proposals-be, jobs-be, financial-be, communications-be

**Flow Steps:**

1. **Receive contract offer:**
   - Client creates contract offer from proposal
   - Freelancer receives notification:
     - Push: "You received a contract offer for [Job Title]"
     - Email: "A client wants to hire you"
     - In-app: Red dot on contracts icon
   - Navigate to /app/(contracts)/pending to view offer

2. **Review contract offer:**
   - /app/(contracts)/pending lists all pending offers
   - Click offer → navigate to /app/(contracts)/[contractId]
   - Contract details page shows:
     - **Job information:**
       - Job title
       - Job description summary (link to full job)
       - Category & skills required
     - **Contract terms:**
       - Contract type: Hourly / Fixed-price
       - **For Hourly:**
         - Hourly rate (your bid or negotiated rate)
         - Weekly limit (hours per week, e.g., 40 hours/week)
         - Manual time (freelancer logs time) or Time Tracker (auto-tracking)
       - **For Fixed-price:**
         - Total budget
         - Milestones breakdown (if defined):
           - Milestone 1: [Description] - $[Amount] - Due [Date]
           - Milestone 2: [Description] - $[Amount] - Due [Date]
         - Payment schedule: Per milestone / Upon completion
       - Start date: When work begins
       - End date: Estimated completion (optional)
       - Weekly hours commitment (if hourly)
     - **Client information:**
       - Client name, company
       - Client rating & total spend
       - Payment method verified badge
       - Total contracts completed
       - Hire rate
     - **Contract clauses:**
       - Confidentiality agreement (NDA) if required
       - IP ownership: Work-for-hire (client owns) / Freelancer retains
       - Termination policy: Notice period, refund terms
       - Dispute resolution process
       - Platform fees: Skillsier's cut (e.g., 10% for first $500, 5% after)
     - **Payment protection:**
       - For Fixed-price: Funds escrowed by client
       - Escrow status: "Client has deposited $[Amount] in escrow"
       - For Hourly: Payment method on file, auto-charge weekly
   - **Required documents to review:**
     - Terms of Service (link)
     - User Agreement (link)
     - Custom contract addendum (if client added)

3. **Contract actions:**
   - **Accept offer:**
     - Click "Accept & Start Work" button
     - Confirmation modal: "By accepting, you agree to the terms and conditions"
     - Checkboxes:
       - ✅ I have read and agree to the contract terms
       - ✅ I confirm my availability to start on [Start Date]
       - ✅ I understand the payment schedule and platform fees
     - Click "Confirm & Accept"
     - Contract status → Active
     - Client is notified
     - Redirect to contract workroom: /app/(workroom)/[contractId]
   - **Decline offer:**
     - Click "Decline Offer" button
     - Modal: "Why are you declining? (optional)"
       - Reasons: Budget too low / Timeline doesn't work / Found another opportunity / Other
     - Click "Confirm Decline"
     - Contract status → Declined
     - Client is notified
     - Proposal returns to "Not hired" status
   - **Request changes (negotiate):**
     - Click "Request Changes" button
     - Opens inline editor to send message to client:
       - Suggest different rate/budget
       - Request milestone adjustments
       - Propose different start date
       - Ask for clarifications
     - Send message → client receives negotiation request
     - Contract status → Negotiating
     - Client can accept changes, counter-offer, or withdraw offer
   - **Ask question:**
     - Click "Message Client" button
     - Opens conversation in inbox
     - Ask clarifying questions before accepting
     - Does not change contract status

4. **Contract expiration:**
   - Contract offers expire after 7 days if not accepted
   - Reminder notifications sent:
     - 3 days before expiration: "Contract offer expires in 3 days"
     - 1 day before: "Contract offer expires tomorrow"
   - After expiration:
     - Contract auto-declines
     - Freelancer can request client to resend offer

5. **Start work (after acceptance):**
   - Contract status → Active
   - Access workroom: /app/(workroom)/[contractId]
   - Workroom shows:
     - Milestone tracker (fixed-price) or timesheet (hourly)
     - File sharing area
     - Communication thread
     - Contract actions: Request milestone release, Log time, End contract
   - Enable Time Tracker app (if hourly with auto-tracking):
     - Download Time Tracker desktop app
     - Install & authenticate
     - Start tracking time for this contract
   - For fixed-price: Begin work on first milestone
   - For hourly: Start logging hours immediately

6. **Onboarding checklist (first contract):**
   - If first contract, show onboarding checklist:
     - ✅ Contract accepted
     - ☐ Complete W-9 / W-8BEN tax form (if not done)
     - ☐ Add payout method (if not done)
     - ☐ Set up Time Tracker (if hourly)
     - ☐ Upload first deliverable
     - ☐ Log first work entry
   - Progress bar shows completion percentage
   - Tips & help articles linked for each step

7. **Tax & compliance forms:**
   - If US-based freelancer and first contract: prompt to complete W-9
   - If international: complete W-8BEN form
   - Navigate to /app/(billing)/tax-center
   - Fill out tax information form:
     - Legal name
     - Tax ID / EIN / SSN
     - Business type (individual / LLC / Corporation)
     - Country of residence
     - Tax treaty benefits (if applicable)
   - Electronic signature
   - Submit → stored in financial-be
   - Client won't be charged until tax forms completed

8. **Set up payout method (if not already done):**
   - Navigate to /app/(billing)/payout-methods
   - Add bank account or payment method:
     - **Bank account (direct deposit):**
       - Country, Currency
       - Account holder name
       - Account number, Routing number (US) / IBAN / Sort code
       - Bank name
     - **PayPal / Payoneer / Wise:**
       - Connect account via OAuth
       - Confirm email linked to account
     - **Crypto (if supported):**
       - Wallet address
       - Currency (BTC / ETH / USDC)
   - Verify payout method:
     - For bank: micro-deposits verification (2 small deposits, confirm amounts)
     - For PayPal/etc: instant OAuth verification
   - Set as default payout method
   - Can add multiple methods, set one as default

9. **Mobile: Enable push notifications:**
   - Prompt to enable push notifications for contract updates:
     - "Stay updated on milestones, payments, and messages"
     - Allow / Don't Allow
   - If allowed: notifications enabled for:
     - Milestone approved
     - Payment released
     - Client sent message
     - Weekly timesheet reminder
     - Contract ending soon

10. **Contract status tracking:**
    - View all contracts at /app/(contracts)
    - Tabs:
      - Active (ongoing contracts)
      - Pending (offers waiting for your action)
      - Ended (completed or terminated contracts)
    - Each contract card shows:
      - Job title, client name
      - Contract type, rate/budget
      - Start date, duration
      - Earnings to date
      - Status indicator
      - Quick actions: View workroom, Message client, End contract

**Branches & Edge Cases:**

- **Multiple offers:** If multiple clients send offers simultaneously → all appear in Pending tab, can accept multiple
- **Conflict with current contracts:** System warns if accepting would exceed your stated availability
- **No payout method:** Cannot accept contract until payout method added
- **No tax forms:** US-based freelancers must complete W-9 before contract activation
- **Client withdraws offer:** If client cancels offer before you accept → notification "Contract offer was withdrawn"
- **Contract modifications:** Client can modify contract terms before you accept → you're notified of changes
- **Escrow insufficient:** For fixed-price, client must fully fund escrow before you can start
- **KYC required:** Some high-value contracts require ID verification before acceptance
- **Agency contracts:** If you're part of agency, agency admin must approve contract
- **Non-compete clauses:** Some contracts have non-compete or exclusivity terms → clearly highlighted
- **Automatic time tracking:** If required, must install Time Tracker app before starting hourly work
- **Contract insurance:** Option to purchase contract insurance (protects against non-payment, scope changes)
- **Contract templates:** For repeat clients, can accept using saved contract templates
- **Multi-currency:** Contract may be in different currency → conversion rates shown
- **Rate negotiation history:** View history of rate negotiations for this job
- **Contract comparison:** Compare multiple offers side-by-side
- **Background check:** Some contracts require background check clearance
- **Freelancer blocklist:** System checks if client has blocked you (shouldn't happen, but safeguard)
- **Contract capacity:** Cannot accept if you've reached max active contracts (e.g., 10 active contracts limit)
- **Skill mismatch warning:** If contract requires skills not on your profile → warning shown
- **Contract review period:** 24-hour review period before contract becomes active (can cancel within)
- **Auto-decline:** If no action after 7 days, contract auto-declines
- **Milestone pre-approval:** Some clients require milestone deliverables pre-approved before starting
- **Work samples required:** Contract may require submitting work samples before acceptance
- **Contract amendments:** After acceptance, changes require mutual agreement via amendment flow

**Notifications:**

- Contract offer received
- Contract offer reminder (3 days, 1 day before expiry)
- Contract offer expiring soon (1 hour before)
- Contract offer expired
- Contract offer withdrawn by client
- Contract terms updated (if modified before acceptance)
- Contract accepted (confirmation)
- Contract declined (confirmation)
- Negotiation request sent
- Client responded to negotiation
- Work can begin (all prerequisites met)
- First contract milestone/timesheet due
- Tax forms required reminder
- Payout method required reminder
- KYC verification required

**Analytics:**

- contracts.offer_received (contract_id, contract_type, rate_budget)
- contracts.offer_viewed (contract_id, view_duration)
- contracts.offer_accepted (contract_id, time_to_accept, contract_type, rate_budget)
- contracts.offer_declined (contract_id, reason)
- contracts.offer_expired (contract_id, no_action)
- contracts.negotiation_requested (contract_id, requested_changes)
- contracts.negotiation_accepted (contract_id)
- contracts.negotiation_declined (contract_id)
- contracts.client_messaged (contract_id, from_source)
- contracts.workroom_accessed (contract_id, first_access)
- contracts.tax_form_started (form_type)
- contracts.tax_form_completed (form_type, completion_time)
- contracts.payout_method_added (method_type, is_first)
- contracts.payout_method_verified (method_type, verification_time)
- contracts.time_tracker_downloaded (contract_id)
- contracts.time_tracker_authenticated (contract_id)
- contracts.onboarding_checklist_viewed (is_first_contract)
- contracts.onboarding_step_completed (step_name)
- contracts.first_work_started (contract_id, time_since_acceptance)
- contracts.multiple_offers_received (count)
- contracts.contract_comparison_used (contract_ids)

**Sources:** contracts-be.user-stories.md, contracts-be-folder-structure.md, financial-be.user-stories.md, users-be.user-stories.md, proposals-be.user-stories.md, combined-fe-folder-strucure.md

---

### FR-6 — Delivery (Fixed) / Timesheets (Hourly)

**Persona:** Freelancer

**Preconditions & Triggers:** Active contract, work in progress

**Primary Screens:**
- **Web:** /app/(workroom)/[contractId], /app/(contracts)/[contractId]/milestones, /app/(contracts)/[contractId]/timesheets
- **Mobile:** Same routes with mobile-optimized time entry and file uploads

**System Touchpoints:** contracts-be, financial-be, storage-be, communications-be, users-be

**Flow Steps:**

#### **For Fixed-Price Contracts:**

1. **View milestones:**
   - Navigate to /app/(contracts)/[contractId]/milestones
   - Milestone list shows:
     - Milestone number & description
     - Amount (e.g., "$500")
     - Due date
     - Status: Pending / In Progress / Submitted / In Review / Approved / Paid
     - Progress indicator
   - Click milestone → expanded view with deliverable requirements

2. **Work on milestone:**
   - Start working on deliverable
   - Upload work-in-progress files to workroom:
     - Navigate to /app/(workroom)/[contractId]/files
     - Drag & drop files or click "Upload"
     - Organize in folders (e.g., "Milestone 1 - Drafts")
     - Add version notes
   - Communicate progress with client via workroom chat
   - Update milestone status manually: "In Progress" (optional)

3. **Submit milestone for review:**
   - When milestone complete, go to /app/(contracts)/[contractId]/milestones
   - Click milestone → "Submit for Review" button
   - **Submission modal:**
     - Attach final deliverables:
       - Upload files (documents, images, videos, code, etc.)
       - Max 10 files, 100 MB each
       - Link to external files (GitHub, Figma, Google Drive)
     - Add submission notes:
       - Describe what was delivered
       - Highlight key features or decisions
       - Note any deviations from original scope
       - Mention next steps or dependencies
     - Character limit: 2000 characters
   - Click "Submit for Review"
   - Milestone status → Submitted
   - Client receives notification: "Milestone [N] submitted for review"
   - Escrow funds for this milestone now pending release

4. **Client reviews submission:**
   - Client views submission in their contract view
   - Client options:
     - **Approve:** Releases milestone payment
     - **Request changes:** Sends feedback, milestone status → In Progress
     - **Reject:** Disputes milestone (escalates to dispute process)

5. **Handle feedback (if changes requested):**
   - Receive notification: "Client requested changes on Milestone [N]"
   - View client feedback in /app/(contracts)/[contractId]/milestones/[milestoneId]
   - Client's comments show in milestone thread
   - Make requested changes
   - Upload revised deliverables
   - Resubmit milestone: "Submit Revised Version"
   - Milestone status → Submitted (again)
   - Version history tracked (Revision 1, 2, 3...)

6. **Milestone approved & payment released:**
   - Client approves milestone
   - Notification: "Milestone [N] approved! $[Amount] released to your account"
   - Milestone status → Approved
   - Funds moved from escrow to your Skillsier balance
   - Funds available for withdrawal after 5-day security hold
   - Move to next milestone (if any)

7. **Final milestone & contract completion:**
   - After all milestones approved:
     - Contract status → Completed
     - Navigate to /app/(contracts)/[contractId]/review
     - Prompt to leave review for client
   - Contract closed, all funds released
   - View final earnings in /app/(billing)/earnings

#### **For Hourly Contracts:**

8. **Timesheet week view:**
   - Navigate to /app/(contracts)/[contractId]/timesheets
   - Weekly view shows:
     - Current week (Mon-Sun)
     - Days of week with hours logged per day
     - Total hours for week
     - Weekly limit (if set, e.g., "35 / 40 hours")
     - Status: Draft / Submitted / Approved
   - Previous weeks listed below (expandable)

9. **Manual time logging:**
   - **Add time entry for a day:**
     - Click "+ Add Time" button
     - Select date (from current or previous days)
     - Enter hours worked (e.g., 5.5 hours)
     - Add description of work done:
       - "Implemented user authentication module"
       - "Fixed bugs in checkout flow"
       - "Client meeting and requirements gathering"
     - Character limit: 500 characters
     - Attach work files (optional): screenshots, code snippets
   - **Edit time entry:**
     - Click existing entry → inline edit
     - Modify hours or description
     - Cannot edit after timesheet submitted (unless client rejects)
   - **Delete time entry:**
     - Click "Delete" icon
     - Confirm deletion
     - Cannot delete after timesheet submitted

10. **Automatic Time Tracker (if enabled):**
    - Download and install Time Tracker desktop app (if using auto-tracking)
    - Authenticate app with contract
    - Start timer when working on contract:
      - Click "Start" button
      - App tracks active time (keyboard/mouse activity)
      - Screenshots taken at intervals (e.g., every 10 min) for client transparency
      - Idle time detection: pauses after 5 min inactivity
    - Stop timer when done working
    - Tracked time automatically syncs to /app/(contracts)/[contractId]/timesheets
    - Review tracked hours before submitting timesheet:
      - View screenshots and activity levels
       - Edit/remove entries if needed (e.g., delete accidental tracking)

11. **Submit weekly timesheet:**
    - At end of week (Sunday midnight or custom cutoff), timesheet ready for submission
    - Navigate to /app/(contracts)/[contractId]/timesheets
    - Review total hours for week
    - Click "Submit Timesheet" button
    - **Submission modal:**
      - Summary: "Week of [Date Range]: [Total Hours] hours"
      - Breakdown by day
      - Total amount: "[Total Hours] × $[Rate] = $[Amount]"
      - Add weekly summary notes (optional):
        - "Completed features X, Y, Z"
        - "Blocked by: need client feedback on design"
      - Platform fee breakdown shown: "Your earnings: $[Amount] (after [X]% fee)"
    - Click "Submit for Approval"
    - Timesheet status → Submitted
    - Client receives notification: "Timesheet for week of [Date] submitted"
    - Cannot edit timesheet after submission

12. **Client reviews timesheet:**
    - Client views timesheet in their contract view
    - Client sees:
      - Hours breakdown by day
      - Work descriptions
      - Screenshots (if Time Tracker used)
      - Activity levels
    - Client options:
      - **Approve:** Releases payment for week
      - **Dispute hours:** Questions specific entries, timesheet status → In Review
      - **Reject:** Disputes entire timesheet (escalates to dispute)

13. **Handle disputed hours:**
    - If client disputes, receive notification: "Client disputed [X] hours on your timesheet"
    - View disputed entries with client's comments
    - Options:
      - **Accept dispute:** Remove disputed hours, resubmit
      - **Provide clarification:** Reply to client's comments, explain work done
      - **Escalate to Skillsier support:** If cannot resolve, request admin mediation
    - Dispute resolution tracked in workroom thread

14. **Timesheet approved & payment:**
    - Client approves timesheet
    - Notification: "Timesheet approved! $[Amount] added to your account"
    - Timesheet status → Approved
    - Funds added to your Skillsier balance
    - Payment available for withdrawal after 5-day security hold
    - Next week's timesheet starts fresh

15. **Weekly limit enforcement:**
    - If contract has weekly limit (e.g., 40 hours/week):
      - Time Tracker auto-stops at limit
      - Manual entries show warning: "You've reached your weekly limit of 40 hours"
      - Cannot log more hours unless client approves overage
   - To exceed limit:
     - Request overage approval from client
     - Client receives request, can approve extra hours
     - If approved, limit lifted for that week
     - Extra hours billed at same rate (or overtime rate if specified)

16. **Missed timesheet submission:**
    - Reminder notifications:
      - Friday: "Submit your timesheet by Sunday for timely payment"
      - Sunday 6 PM: "Timesheet due in 6 hours"
    - If not submitted by deadline:
      - Auto-reminder to client: "Freelancer hasn't submitted timesheet yet"
      - Can submit late (but may delay payment processing)
      - Late submission flagged in contract history

17. **Workroom activity:**
    - All deliverables, time logs, and communications happen in workroom
    - Navigate to /app/(workroom)/[contractId]
    - **Tabs:**
      - **Overview:** Contract summary, milestones/timesheet status, earnings summary
      - **Files:** Shared files and deliverables (organized by milestone or date)
      - **Chat:** Real-time messaging with client (separate from inbox)
      - **Activity:** Timeline of all contract events (time logged, milestones submitted, payments released)
    - **Workroom features:**
      - Pin important messages
      - File version control
      - @mention client in messages
      - Quick actions: Submit milestone, Log time, Request payment

**Branches & Edge Cases:**

- **Milestone scope creep:** If client requests work beyond milestone scope → negotiate additional payment or new milestone
- **Milestone rejection:** If client rejects milestone → dispute process (see Disputes journey)
- **Late milestone delivery:** If past due date → client can extend deadline or close contract
- **Partial milestone payment:** Some contracts allow partial payment for partial delivery
- **Milestone dependencies:** Some milestones require prior milestone completion
- **Overtime hours:** Hourly contracts may have overtime rate (e.g., 1.5× after 40 hours/week)
- **Bonus payments:** Client can add bonus payments for exceptional work (outside regular milestones/timesheets)
- **Time rounding:** Time entries rounded to nearest 15-minute increment (configurable)
- **Weekend work:** Can log hours on weekends (no restrictions unless contract specifies)
- **Holiday pay:** No automatic holiday pay (only hours worked are paid)
- **Timesheet resubmission:** If client rejects, must fix and resubmit timesheet
- **Screenshot privacy:** Can disable screenshots in Time Tracker settings (but client may require them)
- **Idle time:** Idle time (>5 min inactivity) auto-paused and not billed
- **Multiple contracts:** Time Tracker can track multiple contracts simultaneously (switch between them)
- **Offline time tracking:** Time Tracker works offline, syncs when connection restored
- **Time entry backdating:** Can log time for past days (up to 7 days back), older requires explanation
- **Timezone differences:** Timesheets use your local timezone, but show client's timezone for reference
- **Weekly cutoff:** Timesheet week ends Sunday 11:59 PM your timezone (or custom cutoff)
- **Payment delays:** If client doesn't approve within 5 days, auto-approve timesheet (configurable)
- **Disputed milestones:** Funds remain in escrow during dispute resolution
- **Delivery format requirements:** Some milestones specify required file formats (e.g., PSD, MP4, PDF)
- **Version control:** Maintain version history for all deliverables (auto-versioning)
- **Work samples:** Can mark deliverables as "OK to use in portfolio" (with client consent)
- **Collaborative deliverables:** If team project, credit multiple freelancers per milestone
- **Milestone amendments:** Client can modify milestone requirements mid-contract (requires your approval)
- **Early completion bonus:** Some contracts offer bonus for completing milestones ahead of schedule
- **Quality assurance:** Some contracts have QA review period before milestone approval
- **Auto-screenshots opt-out:** Can opt out of screenshots (but may affect client trust)
- **Manual time limits:** Cannot log >24 hours in single day (anti-fraud measure)
- **Time tracking audit:** Random audits for high-value contracts (ensure accuracy)

**Notifications:**

- Milestone submitted for review
- Client reviewed milestone (approved / changes requested)
- Milestone payment released
- Milestone disputed
- Timesheet submitted
- Timesheet approved
- Timesheet disputed
- Weekly timesheet reminder (Friday, Sunday)
- Weekly limit reached (approaching/at limit)
- Payment available for withdrawal (5-day hold ended)
- Late timesheet submission warning
- Client messaged you in workroom
- File uploaded to workroom
- Overtime hours approval needed
- Bonus payment added
- Contract milestone approaching due date
- Time Tracker not running (if required)
- Screenshot upload failed (retry)

**Analytics:**

- milestones.milestone_viewed (milestone_id, contract_id)
- milestones.work_started (milestone_id, time_to_start)
- milestones.file_uploaded (milestone_id, file_type, file_size)
- milestones.milestone_submitted (milestone_id, files_count, notes_length, on_time)
- milestones.milestone_resubmitted (milestone_id, revision_number)
- milestones.milestone_approved (milestone_id, time_to_approval, revision_count)
- milestones.feedback_received (milestone_id, feedback_length)
- milestones.changes_made (milestone_id, change_type)
- timesheets.time_entry_added (contract_id, hours, entry_method, has_description)
- timesheets.time_entry_edited (entry_id, hours_before, hours_after)
- timesheets.time_entry_deleted (entry_id, reason)
- timesheets.time_tracker_started (contract_id)
- timesheets.time_tracker_stopped (contract_id, duration, idle_time)
- timesheets.timesheet_viewed (timesheet_id, week_of)
- timesheets.timesheet_submitted (timesheet_id, total_hours, within_limit, on_time)
- timesheets.timesheet_approved (timesheet_id, time_to_approval)
- timesheets.timesheet_disputed (timesheet_id, disputed_hours, reason)
- timesheets.dispute_resolved (timesheet_id, resolution, hours_adjusted)
- timesheets.weekly_limit_reached (contract_id, week_of)
- timesheets.overage_requested (contract_id, additional_hours)
- timesheets.overage_approved (contract_id, additional_hours)
- workroom.accessed (contract_id, tab_name)
- workroom.file_shared (contract_id, file_type, file_size)
- workroom.message_sent (contract_id, message_length)
- workroom.activity_viewed (contract_id)
- payment.milestone_released (milestone_id, amount)
- payment.timesheet_released (timesheet_id, amount)
- payment.bonus_added (contract_id, bonus_amount, reason)

**Sources:** contracts-be.user-stories.md, contracts-be-folder-structure.md, financial-be.user-stories.md, storage-be.user-stories.md, communications-be.user-stories.md, combined-fe-folder-strucure.md

---

### FR-7 — Get Paid & Withdraw

**Persona:** Freelancer

**Preconditions & Triggers:**
- Freelancer has completed work and earned money
- Contract milestones approved or timesheets approved
- Payment released from escrow to freelancer's wallet
- Freelancer has verified payout method on file
- Freelancer meets minimum withdrawal threshold

**Primary Screens:**
- **Web:**
  - `/app/(billing)/wallet` — Wallet overview with balance breakdown
  - `/app/(billing)/wallet/withdraw` — Initiate withdrawal
  - `/app/(billing)/payout-methods` — Manage payout methods
  - `/app/(billing)/payout-methods/add` — Add bank account, PayPal, etc.
  - `/app/(billing)/payouts` — Payout history
  - `/app/(billing)/payouts/[payoutId]` — Track specific payout status
  - `/app/(billing)/payouts/[payoutId]/receipt` — Download receipt
  - `/app/(billing)/payouts/schedule` — Set up automatic withdrawals
  - `/app/(billing)/payouts/pending` — View in-progress payouts
  - `/app/(billing)/transactions` — Complete transaction history

- **Mobile:**
  - `/(billing)/wallet`
  - `/(billing)/wallet/withdraw`
  - `/(billing)/payout-methods`
  - `/(billing)/payout-methods/add`
  - `/(billing)/payouts`
  - `/(billing)/payouts/[payoutId]`
  - `/(billing)/payouts/[payoutId]/receipt`

**System Touchpoints:** financial-be (wallet, payout, payout_method, transaction, withdrawal_limit, tax, bank_verification), users-be (KYC verification level), communications-be (notifications), storage-be (tax documents, receipts)

**Flow Steps:**

1. **View available balance:**
   - Navigate to `/app/(billing)/wallet`
   - See balance breakdown:
     - **Available for withdrawal:** Funds released from escrow and past security hold period
     - **Pending:** Funds in escrow or security hold (5-day hold for new clients)
     - **Reserved:** Funds allocated to active contracts but not yet released
   - View recent earnings summary
   - See upcoming payouts

2. **Set up payout method (first time):**
   - Navigate to `/app/(billing)/payout-methods`
   - Click "Add payout method"
   - Choose method type:
     - **Bank account (ACH):** Most common, 3-5 business days, low/no fees
     - **PayPal:** Faster, higher fees (~2%)
     - **Wire transfer:** International, 1-2 days, higher fees
     - **Wise (TransferWise):** International, competitive rates
     - **Crypto wallet:** Instant, volatile (if enabled)
   - For **bank account:**
     - Enter bank name, routing number, account number
     - Account type (checking/savings)
     - Account holder name (must match profile)
     - Upload void check or bank statement (optional but recommended)
   - For **PayPal:**
     - Enter PayPal email
     - Verify email via link
   - Submit for verification
   - **Bank verification process:**
     - System sends 2 micro-deposits (< $1 each) to bank account
     - Takes 1-3 business days to appear
     - Return to verify page, enter exact amounts
     - Maximum 3 attempts allowed
     - Once verified, method is activated
     - Alternative: Instant verification via Plaid (if available)

3. **Request withdrawal:**
   - Navigate to `/app/(billing)/wallet/withdraw`
   - Enter withdrawal amount
   - System validates:
     - **Minimum threshold:** Usually $10-100 depending on region
     - **Available balance:** Must have sufficient funds
     - **Daily/monthly limits:** Based on KYC verification level
     - **Account status:** Not frozen or under review
   - Choose payout method from verified methods
   - Review:
     - Withdrawal amount
     - Payout fee (if any)
     - Net amount to receive
     - Estimated arrival date
     - Tax withholding (if applicable)
   - Confirm withdrawal
   - Enter 2FA code if required for large withdrawals

4. **Withdrawal processing:**
   - Withdrawal queued for processing
   - Status progression:
     - **Pending:** In queue, awaiting batch processing
     - **Processing:** Batch submitted to payment gateway
     - **In transit:** Funds sent to bank/PayPal
     - **Completed:** Funds arrived (confirmed)
     - **Failed:** Issue occurred (retry or contact support)
   - Processing timelines:
     - **ACH (US):** 3-5 business days
     - **Wire:** 1-2 business days
     - **PayPal:** 1-3 business days
     - **International:** 5-10 business days
   - Batch processing schedule:
     - Daily batches for standard withdrawals
     - Instant payouts available for premium members (higher fee)

5. **Track payout status:**
   - Navigate to `/app/(billing)/payouts`
   - View all payouts with status
   - Click on specific payout for details:
     - Current status and timeline
     - Amount, fees, net amount
     - Payout method used
     - Expected arrival date
     - Reference number for bank tracking
   - Download receipt for tax records
   - Cancel pending payout (if not yet processed)

6. **Set up automatic withdrawals (optional):**
   - Navigate to `/app/(billing)/payouts/schedule`
   - Configure auto-payout settings:
     - **Frequency:** Weekly, bi-weekly, monthly
     - **Minimum threshold:** Only withdraw when balance reaches X
     - **Day of week/month:** Preferred withdrawal day
     - **Default payout method:** Which account to use
   - Enable/disable automatic withdrawals
   - System automatically processes withdrawals based on schedule

7. **View transaction history:**
   - Navigate to `/app/(billing)/transactions`
   - Complete ledger of all financial activity:
     - Earnings (milestone payments, timesheet payments)
     - Platform fees deducted
     - Withdrawals processed
     - Refunds received
     - Connect purchases
     - Subscription charges
   - Filter by:
     - Date range
     - Transaction type
     - Contract
     - Amount range
   - Search by description
   - Export to CSV for accounting

8. **Tax considerations:**
   - System tracks annual earnings
   - Navigate to `/app/(billing)/tax`
   - View year-to-date earnings
   - Download tax forms:
     - **1099-NEC** (US freelancers, if earnings > $600)
     - **1099-K** (if payment volume threshold met)
   - Update tax information:
     - W-9 form (US)
     - W-8BEN (international)
     - VAT information (EU)
   - Some earnings subject to withholding (e.g., international contractors)

**Branches & Edge Cases:**

- **Insufficient balance:** Cannot withdraw more than available balance (excluding pending/reserved)
- **Below minimum threshold:** Cannot withdraw less than minimum (e.g., $10)
- **Verification level limits:**
  - **Unverified:** $500/day, $2,000/month
  - **Email verified:** $1,000/day, $5,000/month
  - **ID verified:** $5,000/day, $20,000/month
  - **Full KYC:** $50,000/day, $200,000/month
- **High-value withdrawals:** Withdrawals over certain amount (e.g., $10,000) require additional verification
- **First withdrawal delay:** First withdrawal to new payout method may have 5-7 day security hold
- **Bank verification failed:** After 3 failed attempts, must contact support or use different method
- **Payout failed:** If bank rejects transfer (wrong account, closed account), funds returned to wallet
- **Payout method expired:** Bank account closed or PayPal email changed → update method
- **Frozen account:** If account under review or compliance issue, withdrawals blocked
- **Chargeback protection:** Recent earnings may have extended hold if client has history of disputes
- **Currency conversion:** If bank account currency differs from earnings, conversion applied (with margin)
- **Holiday delays:** Weekends and bank holidays extend processing time
- **Instant payout option:** Premium members can request instant payout (2% fee, funds within 30 minutes)
- **Partial withdrawals:** Can withdraw partial balance, leave rest for future withdrawals
- **Withdrawal fees:**
  - ACH (US): Free or $0.25
  - Wire (domestic): $15-30
  - Wire (international): $30-50
  - PayPal: 2% of withdrawal amount
  - Crypto: Network gas fees
- **Tax withholding:**
  - US: Optional backup withholding (24%) if missing/invalid TIN
  - International: Withholding based on tax treaty (0-30%)
- **Multiple currency wallets:** If earning in multiple currencies, can withdraw from each separately
- **Negative balance:** If dispute results in negative balance, must add funds before withdrawing
- **Payout method limits:** Some methods have maximum transaction size (e.g., ACH: $25,000)
- **Anti-fraud measures:**
  - First withdrawal to new method requires extra verification
  - Large sudden withdrawals may trigger manual review
  - IP address and device fingerprinting
- **Payout scheduling conflicts:** If scheduled day is weekend/holiday, processes next business day
- **Batch timing:** Withdrawals requested after cutoff time (e.g., 3 PM ET) process next day
- **Receipt generation:** Automatic receipt for every withdrawal (for tax records)
- **Payout cancellation window:** Can cancel withdrawal within 1 hour of request (before batch processing)
- **Failed verification retry:** Can retry micro-deposit verification after 24 hours
- **Multiple payout methods:** Can have multiple methods on file, choose per withdrawal
- **Default method:** Set preferred method for one-click withdrawals
- **Payout method verification badges:** Visual indicator for verified, unverified, expired methods
- **Security holds for new clients:** Payments from first-time clients held for 5 days before available
- **Dispute clawback:** If client wins dispute after payment released, amount deducted from balance
- **Overdraft protection:** Cannot go negative (except in dispute scenarios)

**Notifications:**

- **Earnings available:** "Your $X payment is now available for withdrawal"
- **Withdrawal requested:** "Withdrawal of $X requested successfully"
- **Withdrawal processing:** "Your $X withdrawal is being processed"
- **Withdrawal completed:** "Your $X withdrawal has been sent to [bank/PayPal]"
- **Withdrawal failed:** "Your withdrawal failed - [reason]"
- **Micro-deposits sent:** "Check your bank account for verification deposits"
- **Bank verification required:** "Please verify your bank account to enable withdrawals"
- **Approaching withdrawal limit:** "You're approaching your daily withdrawal limit"
- **Withdrawal limit exceeded:** "Withdrawal exceeds your current limit - increase KYC level"
- **New payout method added:** "New payout method added successfully"
- **Payout method verified:** "Your [bank/PayPal] is now verified"
- **Payout method verification failed:** "Bank verification failed - please retry"
- **Scheduled withdrawal processed:** "Your scheduled $X withdrawal has been processed"
- **Low balance alert:** "Your available balance is below $X" (if alerts enabled)
- **Tax form available:** "Your 1099 for [year] is ready for download"
- **Security hold ended:** "Payment from [client] is now available for withdrawal"

**Analytics:**

- payouts.wallet\_viewed (available\_balance, pending\_balance, reserved\_balance)
- payouts.withdrawal\_initiated (amount, payout\_method\_type, verification\_level)
- payouts.withdrawal\_validated (passed, failure\_reason)
- payouts.withdrawal\_confirmed (amount, fee, net\_amount, payout\_method\_id, estimated\_days)
- payouts.withdrawal\_cancelled (reason, time\_since\_request)
- payouts.payout\_method\_add\_started (method\_type)
- payouts.payout\_method\_added (method\_type, verification\_required)
- payouts.bank\_verification\_requested (method\_id)
- payouts.micro\_deposits\_sent (method\_id)
- payouts.bank\_verification\_attempted (method\_id, attempt\_number, success)
- payouts.bank\_verification\_failed (method\_id, reason)
- payouts.bank\_verification\_completed (method\_id, verification\_method)
- payouts.instant\_verification\_used (method\_id, provider)
- payouts.payout\_status\_viewed (payout\_id, current\_status)
- payouts.payout\_receipt\_downloaded (payout\_id, amount)
- payouts.payout\_cancelled\_by\_user (payout\_id, time\_until\_processing)
- payouts.auto\_withdrawal\_configured (frequency, threshold, method\_id)
- payouts.auto\_withdrawal\_processed (amount, scheduled\_date)
- payouts.transaction\_history\_viewed (date\_range, filters\_applied)
- payouts.transaction\_history\_exported (format, date\_range, transaction\_count)
- payouts.tax\_form\_viewed (form\_type, tax\_year)
- payouts.tax\_form\_downloaded (form\_type, tax\_year)
- payouts.tax\_info\_updated (fields\_changed)
- payouts.withdrawal\_limit\_reached (limit\_type, current\_amount, attempted\_amount)
- payouts.high\_value\_withdrawal\_requested (amount, additional\_verification\_required)
- payouts.payout\_failed (payout\_id, failure\_reason, gateway\_code)
- payouts.payout\_completed (payout\_id, amount, processing\_days)
- payouts.default\_method\_changed (old\_method\_type, new\_method\_type)
- payouts.balance\_alert\_triggered (balance, threshold)
- payouts.currency\_conversion\_applied (from\_currency, to\_currency, amount, rate, margin)
- payouts.instant\_payout\_requested (amount, fee)

**Sources:** financial-be.user-stories.md, financial-be-folder-structure.md, financial-be-database-design.md, combined-fe-folder-strucure.md, users-be.user-stories.md (KYC levels)

---

### FR-8 — Reviews & Reputation

**Persona:** Freelancer (receiving reviews, building reputation)

**Preconditions & Triggers:**
- Contract completed or closed
- Client submits review (triggers freelancer review window)
- Double-blind review period (both parties review simultaneously)
- Review period expires (typically 14 days)

**Primary Screens:**
- **Web:**
  - `/app/(profile)/reviews` — All reviews received
  - `/app/(reviews)/leave/[contractId]` — Leave review for client
  - `/app/(reviews)/[reviewId]` — View specific review detail
  - `/app/(reviews)/[reviewId]/respond` — Respond to review (owner response)
  - `/app/(reviews)/[reviewId]/flag` — Flag inappropriate review
  - `/app/(profile)/stats` — Reputation stats and badges
  - `/app/(reviews)/private-feedback` — View private feedback from clients
  - `/app/(reviews)/drafts` — Draft reviews not yet submitted

- **Mobile:**
  - `/(profile)/reviews`
  - `/(reviews)/leave/[contractId]`
  - `/(reviews)/[reviewId]`
  - `/(reviews)/[reviewId]/respond`
  - `/(reviews)/[reviewId]/flag`
  - `/(profile)/stats`

**System Touchpoints:** reviews-be (reviews, review_responses, review_flags, private_feedback, reputation), contracts-be (contract completion status), users-be (profile, reputation score), communications-be (notifications), admin-be (moderation)

**Flow Steps:**

1. **Receive review notification:**
   - Contract is completed or closed
   - Client submits review of your work
   - You receive notification: "Time to review [Client Name] - both reviews released when complete"
   - Navigate to review page: `/app/(reviews)/leave/[contractId]`

2. **Leave review for client:**
   - **Rating dimensions:**
     - **Overall:** 1-5 stars
     - **Communication:** How responsive and clear was the client?
     - **Payment:** Did they pay on time and as agreed?
     - **Clarity:** Were project requirements clear and well-defined?
     - **Professionalism:** Was the working relationship professional?
     - **Likelihood to work again:** Would you accept another project from this client?
   - **Written feedback:**
     - Public review (visible to everyone, 200-2000 characters)
     - Private feedback (only visible to client, helps them improve)
   - **Additional fields:**
     - Skills used (tags)
     - Project difficulty (easy, moderate, challenging)
     - Would recommend client (yes/no)
   - Save as draft (can return later)
   - Submit review

3. **Double-blind review period:**
   - Neither party sees the other's review until both submit
   - Or until 14-day period expires (whichever comes first)
   - Ensures honest, unbiased feedback
   - Reminders sent:
     - Day 3: "Don't forget to review [Client]"
     - Day 7: "7 days left to complete your review"
     - Day 13: "Last day to leave your review"
   - Can edit review before period ends

4. **Review reveal:**
   - Once both reviews submitted (or period expires):
     - Both reviews become visible
     - Posted to profiles publicly
     - Reputation scores updated
   - If one party doesn't review:
     - Other review still publishes
     - Non-reviewer sees reminder: "You didn't review [Name] - review opportunity closed"

5. **View reviews on profile:**
   - Navigate to `/app/(profile)/reviews`
   - See all reviews received from clients
   - **Filters:**
     - Star rating (5, 4, 3, 2, 1)
     - Date range
     - Contract type (hourly, fixed)
     - Skills/categories
   - **Sort options:**
     - Most recent
     - Highest rated
     - Lowest rated
     - Most helpful (based on votes)
   - **Review display:**
     - Client name (or "Anonymous Client" if privacy requested)
     - Star ratings (overall and by dimension)
     - Written review text
     - Contract title/category
     - Date posted
     - "Helpful" votes count
     - Your response (if provided)

6. **Respond to reviews (owner response):**
   - For any review received, can add a public response
   - Navigate to `/app/(reviews)/[reviewId]/respond`
   - Write response (up to 1000 characters)
   - **Use cases:**
     - Thank client for positive feedback
     - Address concerns in negative reviews professionally
     - Provide context or clarification
     - Showcase professionalism
   - Response appears below original review
   - Can edit response within 24 hours
   - Only one response allowed per review

7. **Flag inappropriate reviews:**
   - If review violates policies:
     - Contains personal information (phone, email, address)
     - Profanity or abusive language
     - Off-topic (doesn't relate to work)
     - Discriminatory or harassing
     - Spam or promotional content
     - Factually inaccurate (wrong person, wrong project)
   - Navigate to `/app/(reviews)/[reviewId]/flag`
   - Select reason for flagging
   - Provide explanation (optional but helpful)
   - Submit flag for admin review
   - Flagged reviews may be hidden pending investigation
   - If valid, review removed; if invalid, stays with note

8. **Reputation score & badges:**
   - Navigate to `/app/(profile)/stats`
   - View your reputation metrics:
     - **Overall rating:** Average of all reviews (weighted)
     - **Total reviews:** Number of reviews received
     - **Response rate:** % of messages replied within 24 hours
     - **Job success score:** % of jobs completed successfully (Upwork-style)
     - **Rehire rate:** % of clients who hired you multiple times
     - **On-time delivery:** % of milestones delivered on/before deadline
     - **Earnings (all-time):** Total lifetime earnings on platform
   - **Badges earned:**
     - **Top Rated:** Consistent 4.8+ rating, high success score
     - **Rising Talent:** New freelancer with strong early performance
     - **100% Job Success:** Perfect success score for 3+ months
     - **Repeat Client Champion:** 5+ repeat clients
     - **Fast Responder:** <1 hour avg response time for 3+ months
     - **Expert Verified:** Passed skills tests at expert level
     - **Long-term:** Active for 2+ years with consistent work
   - Badges displayed prominently on profile
   - Help attract better clients

9. **Private feedback:**
   - Navigate to `/app/(reviews)/private-feedback`
   - View private feedback from clients (only you see this)
   - Helps identify areas for improvement
   - Not factored into public reputation
   - Can filter by date, contract, rating

10. **Review drafts:**
    - Navigate to `/app/(reviews)/drafts`
    - View all draft reviews not yet submitted
    - Continue editing and submit
    - Drafts auto-expire with review period

**Branches & Edge Cases:**

- **Late reviews:** If review period expires, can still see client's review but cannot leave your own
- **No review received:** If client doesn't review, your review still publishes (one-sided)
- **Editing reviews:** Can edit review within 24 hours of submission (before reveal)
- **Deleting reviews:** Cannot delete reviews once published (platform policy)
- **Review appeals:** If you believe review is unfair/violates policies, can appeal to support
- **Review removal:** Only admins can remove reviews (after investigation)
- **Anonymous clients:** Some clients request anonymity (name hidden, just "Anonymous Client")
- **Star rating weight:**
  - More recent reviews weighted more heavily
  - High-value contracts weighted more heavily
  - Reviews from verified clients weighted more heavily
- **Job success score calculation:**
  - Factors: project completion rate, client satisfaction, rehires, dispute resolution
  - Updated monthly
  - Private contracts not counted (to prevent gaming)
- **Badge requirements change:** Platform may update badge criteria over time
- **Reputation impact of disputes:** If dispute resolved against you, hurts job success score
- **Review manipulation:** Gaming reviews (fake reviews, review swaps) results in account suspension
- **Helpful votes:** Other users can vote reviews as "helpful" (for social proof)
- **Response editing:** Can edit owner response within 24 hours, after that locked
- **Flagged review investigation:** Takes 3-5 business days for admin review
- **Review removal compensation:** If review removed as invalid, reputation score recalculated
- **Private feedback accumulation:** Private feedback trends shown over time (e.g., "clients consistently mention...")
- **Reputation trend chart:** Graph showing rating trend over time
- **Category-specific ratings:** If you work in multiple categories, separate ratings per category
- **Client review credibility:** Reviews from established clients weighted more than new clients
- **Zero review impact:** New freelancers with no reviews can still compete (Rising Talent badge helps)
- **Review quota:** Unlimited reviews, all counted (no cherry-picking)
- **Language filter:** Profanity auto-flagged for admin review
- **Review length minimum:** Public review must be at least 200 characters (encourages detail)
- **Emoji/special characters:** Allowed in reviews (but abuse flagged)
- **Review notification settings:** Can customize when you're notified about reviews
- **Bulk review export:** Can export all reviews to PDF for portfolio/resume

**Notifications:**

- **Review reminder:** "Don't forget to review [Client Name] - [X days] left"
- **Review received:** "[Client] left you a review - view now"
- **Review revealed:** "Your review exchange with [Client] is now public"
- **Response to your review:** "[Freelancer] responded to your review"
- **Helpful vote:** "Your review of [Client] received a helpful vote"
- **Review flagged:** "Your flag on [Name]'s review is under investigation"
- **Badge earned:** "Congratulations! You've earned the [Badge Name] badge"
- **Reputation milestone:** "You've reached 100 reviews with a 4.9 average!"
- **Job success score update:** "Your job success score increased to 98%"
- **Private feedback received:** "[Client] left you private feedback"
- **Review appeal update:** "Update on your review appeal - [status]"
- **Flagged review removed:** "Review by [Name] has been removed after investigation"

**Analytics:**

- reviews.review\_page\_viewed (contract\_id, time\_remaining\_days)
- reviews.review\_started (contract\_id)
- reviews.rating\_given (dimension, star\_rating)
- reviews.review\_text\_entered (character\_count, includes\_private\_feedback)
- reviews.review\_saved\_as\_draft (contract\_id)
- reviews.review\_submitted (contract\_id, overall\_rating, has\_private\_feedback, time\_to\_complete)
- reviews.review\_edited (review\_id, fields\_changed)
- reviews.double\_blind\_period\_expired (contract\_id, both\_reviewed)
- reviews.review\_revealed (review\_id, days\_until\_reveal)
- reviews.reviews\_page\_viewed (filters\_applied, sort\_by)
- reviews.review\_detail\_viewed (review\_id)
- reviews.owner\_response\_started (review\_id)
- reviews.owner\_response\_submitted (review\_id, response\_length)
- reviews.owner\_response\_edited (review\_id)
- reviews.review\_flagged (review\_id, flag\_reason)
- reviews.flag\_appeal\_submitted (review\_id, appeal\_reason)
- reviews.reputation\_stats\_viewed
- reviews.badge\_earned (badge\_type)
- reviews.private\_feedback\_viewed (count, date\_range)
- reviews.draft\_review\_resumed (contract\_id, days\_since\_draft)
- reviews.helpful\_vote\_given (review\_id, voter\_role)
- reviews.reputation\_trend\_viewed (time\_range)
- reviews.review\_export\_requested (format, count)

**Sources:** reviews-be.user-stories.md, reviews-be-folder-structure.md, reviews-be-database-design.md, combined-fe-folder-strucure.md, contracts-be.user-stories.md

---

### FR-9 — Skills Tests & Certifications

**Persona:** Freelancer (demonstrating expertise)

**Preconditions & Triggers:**
- Freelancer wants to stand out in search results
- Wants to validate skills publicly
- Client requires certain certifications for job
- Platform offers free skills assessments

**Primary Screens:**
- **Web:**
  - `/app/(tools)/skills-tests` — Browse available tests
  - `/app/(tools)/skills-tests/[testId]` — Test details and start
  - `/app/(tools)/skills-tests/[testId]/take` — Take the test
  - `/app/(tools)/skills-tests/[testId]/results` — View results
  - `/app/(profile)/certifications` — Manage earned certificates
  - `/app/(tools)/skills-tests/retake` — Retake failed tests
  - `/app/(tools)/learning-paths` — Recommended learning paths based on profile

- **Mobile:**
  - `/(tools)/skills-tests`
  - `/(tools)/skills-tests/[testId]`
  - `/(tools)/skills-tests/[testId]/take`
  - `/(tools)/skills-tests/[testId]/results`
  - `/(profile)/certifications`

**System Touchpoints:** users-be (profile, skills, certifications), search-be (skills index, freelancer search ranking), admin-be (test content management), storage-be (certificate PDFs)

**Flow Steps:**

1. **Browse skills tests:**
   - Navigate to `/app/(tools)/skills-tests`
   - Browse tests by category:
     - **Development:** JavaScript, Python, React, Node.js, SQL, etc.
     - **Design:** Photoshop, Figma, UI/UX, Illustration, etc.
     - **Writing:** Content Writing, Copywriting, SEO Writing, Technical Writing
     - **Marketing:** Social Media, SEO, PPC, Email Marketing
     - **Business:** Project Management, Excel, Data Analysis, etc.
   - View test details:
     - Difficulty level (Beginner, Intermediate, Advanced, Expert)
     - Duration (typically 20-40 minutes)
     - Number of questions
     - Pass rate
     - Skills covered
     - Freelancers who passed
   - See if certificate displayed on your profile

2. **Start a skills test:**
   - Click "Take Test" on test page
   - Read instructions and guidelines:
     - Time limit (cannot pause once started)
     - No external help or resources allowed
     - Random question order (anti-cheating)
     - Must score 70%+ to pass (vary by test)
     - Can retake once every 30 days if failed
   - Confirm you're ready
   - Begin test

3. **Take the test:**
   - Answer multiple-choice or coding questions
   - **Question types:**
     - **Multiple choice:** Select best answer from 4-5 options
     - **Code challenges:** Write working code in online IDE
     - **Fill-in-blank:** Complete code snippet
     - **True/False:** Statement evaluation
     - **Scenario-based:** Real-world problem solving
   - Timer displayed at top (warns when 5 min remaining)
   - Can skip and return to questions
   - Progress bar shows completion
   - Submit test when finished (or auto-submit when time expires)

4. **View results:**
   - Receive instant results
   - **Pass:**
     - Overall score (e.g., 85%)
     - Percentile rank (e.g., "Top 15% of test-takers")
     - Certificate earned badge
     - Certificate auto-added to profile
     - Can download PDF certificate
     - Certificate includes:
       - Your name
       - Test name and date
       - Score and percentile
       - Unique certificate ID
       - QR code for verification
   - **Fail:**
     - Overall score
     - Areas of weakness identified
     - Recommended study resources
     - "Retake available in 30 days"
     - Private result (not shown on profile)

5. **Display certificates on profile:**
   - Certificates appear in "Skills & Certifications" section
   - Navigate to `/app/(profile)/certifications`
   - Manage display:
     - Reorder certificates (drag-and-drop)
     - Pin top 3 certificates (featured)
     - Hide certain certificates (if you want)
     - Show/hide scores (optional)
   - Certificates boost search ranking:
     - Appear higher for jobs requiring that skill
     - "Verified [Skill]" badge in search results
     - Client can filter for certified freelancers

6. **Retake failed tests:**
   - If failed, can retake after 30-day cooling period
   - Navigate to `/app/(tools)/skills-tests/retake`
   - See eligible retakes
   - Prepare using recommended resources
   - Retake test (same process as first attempt)
   - New result overwrites old (whether better or worse)
   - Unlimited retakes allowed (with 30-day spacing)

7. **Recommended learning paths:**
   - Navigate to `/app/(tools)/learning-paths`
   - Based on your profile and tests taken, see personalized recommendations:
     - "Complete these tests to become a Verified [Role]"
     - "Top tests for [your category]"
     - "Trending certifications in [industry]"
   - Learning path tracks progress:
     - Which tests completed
     - Which tests in progress
     - Which tests recommended next
     - Estimated time to complete path

**Branches & Edge Cases:**

- **Test integrity:** Tests monitored for suspicious activity (too fast, perfect scores, pattern matching)
- **Cheating detection:** Flagged tests reviewed by admins, may result in disqualification
- **Certificate expiration:** Some certificates expire after 1-2 years (must retake to renew)
- **Test updates:** Tests updated periodically to reflect current best practices
- **Proctoring:** Some advanced tests require webcam proctoring (optional, coming soon)
- **Test preparation:** Free study guides and practice questions available
- **Language options:** Tests available in multiple languages
- **Accessibility:** Accommodations available for disabilities (extra time, screen reader support)
- **Mobile testing:** Most tests can be taken on mobile (except coding tests)
- **Partial credit:** Some questions give partial credit for partially correct answers
- **Negative marking:** No penalty for wrong answers (guess if unsure)
- **Certificate sharing:** Can share certificate link on LinkedIn, resume, portfolio
- **Employer verification:** Certificates include unique ID for employers to verify authenticity
- **Top performers leaderboard:** Opt-in to appear on public leaderboard for each test
- **Test bundle discounts:** Premium tests available (paid) for advanced certifications
- **Industry-recognized certs:** Some tests partnered with industry orgs (Google, AWS, etc.)
- **Test analytics:** See which questions you missed after passing (for learning)
- **Group testing:** Agencies can assign tests to team members (track compliance)
- **Test reminders:** Get reminder to renew expiring certificates
- **Failed attempt history:** Can see history of attempts and score progression (private)
- **Test recommendations:** System recommends tests based on job types you apply to
- **Certificate badges:** Different badge tiers (bronze, silver, gold) based on score/percentile
- **Offline testing:** Cannot take tests offline (requires internet connection)
- **Test abandonment:** If you close tab mid-test, counted as attempt (must retake after 30 days)
- **Adaptive testing:** Some tests adjust difficulty based on your answers (coming soon)

**Notifications:**

- **Test passed:** "Congratulations! You passed the [Test Name] with [Score]%"
- **Certificate earned:** "New certificate added to your profile - [Test Name]"
- **Test failed:** "You scored [Score]% on [Test] - retake available [Date]"
- **Retake available:** "You can now retake the [Test Name]"
- **Certificate expiring soon:** "[Certificate] expires in 30 days - renew now"
- **Certificate expired:** "Your [Certificate] has expired - retake to renew"
- **New test available:** "New [Category] test available - [Test Name]"
- **Test recommendation:** "Based on your profile, consider taking [Test Name]"
- **Leaderboard achievement:** "You're in the top 10% for [Test Name]"

**Analytics:**

- skills\_tests.tests\_list\_viewed (category, filters\_applied)
- skills\_tests.test\_detail\_viewed (test\_id, difficulty\_level)
- skills\_tests.test\_started (test\_id, user\_skill\_level)
- skills\_tests.test\_question\_answered (test\_id, question\_id, correct, time\_spent)
- skills\_tests.test\_question\_skipped (test\_id, question\_id)
- skills\_tests.test\_question\_reviewed (test\_id, question\_id)
- skills\_tests.test\_abandoned (test\_id, questions\_completed, time\_elapsed)
- skills\_tests.test\_submitted (test\_id, questions\_answered, time\_taken)
- skills\_tests.test\_passed (test\_id, score, percentile, time\_taken)
- skills\_tests.test\_failed (test\_id, score, weak\_areas)
- skills\_tests.certificate\_downloaded (test\_id, format)
- skills\_tests.certificate\_shared (test\_id, platform)
- skills\_tests.certificate\_pinned (test\_id)
- skills\_tests.certificate\_hidden (test\_id)
- skills\_tests.retake\_initiated (test\_id, previous\_score, days\_since\_last\_attempt)
- skills\_tests.learning\_path\_viewed (path\_id, tests\_completed)
- skills\_tests.test\_recommended (test\_id, recommendation\_reason)
- skills\_tests.study\_resources\_accessed (test\_id, resource\_type)
- skills\_tests.leaderboard\_viewed (test\_id)

**Sources:** users-be.user-stories.md, search-be.user-stories.md, combined-fe-folder-strucure.md (inferred from (tools) and (profile) routes)

---

### FR-10 — Talent Cloud & Agencies

**Persona:** Freelancer (joining talent pools, creating/joining agencies)

**Preconditions & Triggers:**
- Freelancer invited to join a talent cloud/pool
- Freelancer wants to form an agency with other freelancers
- Client creates a talent cloud for vetted freelancers
- Freelancer wants to collaborate on larger projects

**Primary Screens:**
- **Web:**
  - `/app/(talent)/clouds` — Browse and join talent clouds
  - `/app/(talent)/clouds/[cloudId]` — View talent cloud details
  - `/app/(talent)/clouds/[cloudId]/join` — Request to join cloud
  - `/app/(talent)/my-clouds` — Talent clouds you're part of
  - `/app/(agencies)` — Agency dashboard (if agency owner)
  - `/app/(agencies)/create` — Create new agency
  - `/app/(agencies)/[agencyId]` — Agency profile and management
  - `/app/(agencies)/[agencyId]/team` — Manage agency team members
  - `/app/(agencies)/[agencyId]/jobs` — Jobs assigned to agency
  - `/app/(agencies)/[agencyId]/settings` — Agency settings and billing

- **Mobile:**
  - `/(talent)/clouds`
  - `/(talent)/clouds/[cloudId]`
  - `/(talent)/my-clouds`
  - `/(agencies)`
  - `/(agencies)/[agencyId]`

**System Touchpoints:** users-be (talent_clouds, agencies, team_members), jobs-be (cloud-specific jobs), contracts-be (team contracts), financial-be (split payments), communications-be (team messaging)

**Flow Steps:**

1. **Discover talent clouds:**
   - Navigate to `/app/(talent)/clouds`
   - Browse available talent clouds:
     - **Private clouds:** Invitation-only, created by specific clients
     - **Public clouds:** Open to qualified freelancers (application required)
     - **Premium clouds:** Vetted, high-performing freelancers only
   - View cloud details:
     - Cloud name and description
     - Client/organization running the cloud
     - Number of members
     - Average earnings of members
     - Skills required
     - Benefits (priority access to jobs, higher rates, dedicated support)
   - See if you're eligible to apply

2. **Join a talent cloud:**
   - Click "Request to Join" on cloud page
   - Fill application:
     - Why you want to join
     - Relevant experience
     - Portfolio samples
     - Availability
   - Submit application
   - **Application review:**
     - Reviewed by cloud manager (client or admin)
     - May involve skills tests or interview
     - Takes 3-7 days typically
   - **If accepted:**
     - Gain access to cloud-specific jobs
     - Jobs from this client/org appear first in your feed
     - Often higher rates than public jobs
     - Dedicated cloud community chat
     - Priority support
   - **If rejected:**
     - Receive feedback on why (skill gaps, availability mismatch, etc.)
     - Can reapply after improving profile/skills

3. **Manage cloud memberships:**
   - Navigate to `/app/(talent)/my-clouds`
   - View all talent clouds you're part of
   - For each cloud:
     - See active jobs available to you
     - Cloud-specific announcements
     - Performance metrics within cloud
     - Cloud-specific benefits
   - Leave cloud if desired (can rejoin later with new application)

4. **Create an agency:**
   - Navigate to `/app/(agencies)/create`
   - **Agency setup:**
     - Agency name
     - Description and services offered
     - Logo and cover image
     - Primary skills/specializations
     - Location (if relevant)
     - Agency website (optional)
   - **Agency model:**
     - **Cooperative:** Equal ownership, split profits evenly
     - **Owner-led:** You own agency, hire team members
     - **Hybrid:** Custom profit-sharing arrangement
   - **Initial team:**
     - Invite existing Skillsier freelancers to join
     - Send invites via email or Skillsier username
     - Invited members must accept invitation
   - Submit for approval (agencies reviewed to prevent spam)

5. **Manage agency team:**
   - Navigate to `/app/(agencies)/[agencyId]/team`
   - **Add members:**
     - Invite by email or Skillsier username
     - Specify role (Admin, Member, Contractor)
     - Set permissions (can bid on jobs, can message clients, can manage finances)
   - **Member profiles:**
     - View each member's skills and portfolio
     - See individual reputation and reviews
     - Set availability status
   - **Remove members:**
     - Can remove underperforming or inactive members
     - Must handle any active contracts first
   - **Profit sharing:**
     - Set split percentages for each member
     - Can be equal or performance-based
     - Updated per project or agency-wide default

6. **Apply to jobs as agency:**
   - Browse jobs as normal
   - When applying, select "Apply as [Agency Name]"
   - **Agency proposal:**
     - Showcase full team capabilities
     - Assign team members to specific roles
     - Provide combined portfolio (all members' work)
     - Highlight agency experience and past successful projects
   - Agency profile and collective reviews visible to client
   - Higher chance of landing large/complex projects

7. **Agency job management:**
   - When hired as agency:
     - Contract created with agency (not individual)
     - Navigate to `/app/(agencies)/[agencyId]/jobs`
     - View all active agency contracts
     - Assign tasks to specific members
     - Track progress of each member
     - Submit deliverables on behalf of agency
   - **Agency workroom:**
     - Shared workspace for team collaboration
     - Client communicates with agency (not individuals)
     - Internal team chat separate from client chat
     - File sharing and version control

8. **Agency billing and payouts:**
   - Navigate to `/app/(agencies)/[agencyId]/settings`
   - **Payment flow:**
     - Client pays agency as whole
     - Funds go to agency wallet
     - Agency owner distributes to members based on agreed splits
   - **Split payment options:**
     - **Automatic:** System splits payment based on percentages (released when milestone approved)
     - **Manual:** Agency owner reviews and splits payment manually
     - **Hybrid:** Some split auto, some manual (e.g., bonuses)
   - **Agency fees:**
     - Platform takes standard fee from total agency payment
     - Agency can add markup on top of member rates (agency margin)
   - **Member earnings:**
     - Each member sees their portion of earnings
     - Can withdraw as normal from individual wallet

9. **Agency profile and reputation:**
   - Agency has its own profile page (like freelancer profile)
   - **Agency profile includes:**
     - Team member showcase
     - Combined portfolio (best work from all members)
     - Agency reviews (separate from individual reviews)
     - Total projects completed
     - Agency success score
     - Specializations
   - Clients can hire agency specifically or hire individuals within agency

**Branches & Edge Cases:**

- **Cloud invitation:** Can be directly invited by client to join cloud (skip application)
- **Cloud-specific rates:** Some clouds guarantee minimum rates or rate increases
- **Cloud exclusivity:** Some clouds require exclusivity (can't work with competing clients)
- **Cloud removal:** Cloud managers can remove members for poor performance
- **Cloud dissolution:** If client closes cloud, members notified and can apply to other clouds
- **Agency name conflicts:** Agency name must be unique on platform
- **Agency verification:** Verified agencies (badge) after completing 10+ projects successfully
- **Agency minimum size:** Agencies must have at least 2 active members
- **Member departure:** If member leaves agency mid-project, must replace or reduce scope
- **Agency disputes:** If member disagrees with profit split, can escalate to support
- **Agency profile ownership:** Owner controls profile, but members can add their individual work
- **Agency messaging:** Client can message entire agency or specific members
- **Multi-agency membership:** Freelancers can join multiple non-competing agencies
- **Agency contracts vs individual:** Can choose to apply as individual even if in agency
- **Cloud job priority:** Cloud-specific jobs may bypass public job board entirely
- **Agency hierarchy:** Can set multiple admin roles with different permissions
- **Agency white-label:** Premium agencies can have custom domain/branding (future)
- **Agency analytics:** Track team performance, utilization rates, revenue trends
- **Cloud recommendations:** System recommends clouds based on your profile/skills
- **Agency invitation acceptance:** Invited members have 7 days to accept before invite expires
- **Agency settings:** Can make agency profile public or private (invite-only)
- **Cloud job allocation:** Cloud managers decide how to allocate jobs among members
- **Agency reviews:** Clients review agency as unit, individual members not separately reviewed

**Notifications:**

- **Cloud invite received:** "[Client] invited you to join [Cloud Name]"
- **Cloud application accepted:** "You've been accepted to [Cloud Name]!"
- **Cloud application rejected:** "Your application to [Cloud Name] was not accepted"
- **New cloud job available:** "New job available in [Cloud Name] - [Job Title]"
- **Agency invite received:** "[Name] invited you to join [Agency Name]"
- **Agency invite accepted:** "[Name] accepted your agency invitation"
- **Agency member joined:** "[Name] joined your agency [Agency Name]"
- **Agency member left:** "[Name] left your agency [Agency Name]"
- **Agency job awarded:** "Your agency won the job - [Job Title]"
- **Agency payment received:** "$X payment received for [Agency Project]"
- **Agency payment split:** "$X deposited to your wallet from [Agency Project]"
- **Cloud membership expiring:** "Your [Cloud Name] membership expires in 30 days"

**Analytics:**

- talent\_clouds.clouds\_list\_viewed (filter\_by, sort\_by)
- talent\_clouds.cloud\_detail\_viewed (cloud\_id, cloud\_type)
- talent\_clouds.cloud\_application\_started (cloud\_id)
- talent\_clouds.cloud\_application\_submitted (cloud\_id, portfolio\_items, experience\_years)
- talent\_clouds.cloud\_application\_accepted (cloud\_id, days\_to\_review)
- talent\_clouds.cloud\_application\_rejected (cloud\_id, rejection\_reason)
- talent\_clouds.cloud\_joined (cloud\_id, invitation\_or\_application)
- talent\_clouds.cloud\_left (cloud\_id, membership\_duration, jobs\_completed)
- talent\_clouds.cloud\_job\_viewed (job\_id, cloud\_id)
- talent\_clouds.my\_clouds\_viewed (active\_clouds\_count)
- agencies.agency\_creation\_started
- agencies.agency\_created (agency\_name, model\_type, initial\_team\_size)
- agencies.team\_invite\_sent (agency\_id, invitee\_role)
- agencies.team\_invite\_accepted (agency\_id, inviter\_response\_time)
- agencies.team\_member\_added (agency\_id, member\_role, permissions)
- agencies.team\_member\_removed (agency\_id, removal\_reason)
- agencies.agency\_profile\_viewed (agency\_id, viewer\_role)
- agencies.agency\_proposal\_submitted (agency\_id, job\_id, team\_members\_assigned)
- agencies.agency\_hired (agency\_id, contract\_id, contract\_value)
- agencies.agency\_job\_assigned (agency\_id, job\_id, assigned\_to)
- agencies.agency\_payment\_received (agency\_id, amount, split\_method)
- agencies.member\_payment\_distributed (agency\_id, member\_id, amount, split\_percentage)
- agencies.agency\_settings\_updated (agency\_id, settings\_changed)

**Sources:** users-be.user-stories.md (inferred talent clouds and agencies), jobs-be.user-stories.md, contracts-be.user-stories.md, financial-be.user-stories.md, combined-fe-folder-strucure.md

---

### FR-11 — Learning Paths & Achievements

**Persona:** Freelancer (skill development, gamification)

**Preconditions & Triggers:**
- Freelancer wants to upskill
- Platform offers learning recommendations
- Freelancer completes milestones or achieves goals
- Gamification for engagement and retention

**Primary Screens:**
- **Web:**
  - `/app/(tools)/learning` — Learning dashboard
  - `/app/(tools)/learning/paths` — Browse learning paths
  - `/app/(tools)/learning/paths/[pathId]` — Learning path detail
  - `/app/(tools)/learning/paths/[pathId]/enroll` — Enroll in learning path
  - `/app/(tools)/learning/courses/[courseId]` — Individual course/tutorial
  - `/app/(tools)/learning/achievements` — View all achievements earned
  - `/app/(profile)/achievements` — Display achievements on profile
  - `/app/(tools)/learning/progress` — Track learning progress

- **Mobile:**
  - `/(tools)/learning`
  - `/(tools)/learning/paths`
  - `/(tools)/learning/paths/[pathId]`
  - `/(tools)/learning/courses/[courseId]`
  - `/(tools)/learning/achievements`
  - `/(profile)/achievements`

**System Touchpoints:** users-be (learning_progress, achievements, profile), storage-be (course content), subscriptions-be (premium course access)

**Flow Steps:**

1. **Learning dashboard:**
   - Navigate to `/app/(tools)/learning`
   - **Overview includes:**
     - **Your active learning paths:** Paths you're currently enrolled in
     - **Progress summary:** % completion of each path
     - **Recommended for you:** Personalized path recommendations
     - **Trending paths:** Popular paths among similar freelancers
     - **Recently completed:** Courses/paths finished recently
     - **Achievement highlights:** Recent badges earned

2. **Browse learning paths:**
   - Navigate to `/app/(tools)/learning/paths`
   - **Filter by:**
     - Category (Development, Design, Writing, Marketing, etc.)
     - Skill level (Beginner, Intermediate, Advanced)
     - Duration (Hours to complete)
     - Free vs. Premium
   - **Learning path examples:**
     - **"Become a React Developer"** (20 hours, 5 courses, 3 projects)
     - **"Master UI/UX Design"** (30 hours, 8 courses, 2 certifications)
     - **"SEO Expert Path"** (15 hours, 4 courses, 1 capstone)
     - **"Full Stack JavaScript"** (50 hours, 12 courses, 5 projects)
     - **"Technical Writer Pro"** (10 hours, 3 courses, writing samples)

3. **View learning path details:**
   - Click on a learning path
   - Navigate to `/app/(tools)/learning/paths/[pathId]`
   - **Path details:**
     - Description and learning outcomes
     - Prerequisites (if any)
     - Estimated time to complete
     - Number of courses included
     - Projects/assignments
     - Skills tests required
     - Certificate upon completion
     - Completion rate (% of enrollees who finish)
   - **Curriculum breakdown:**
     - List of courses in order
     - For each course:
       - Title and description
       - Duration (video/reading time)
       - Quiz/assignment (if any)
       - Lock status (must complete previous courses first)
   - **Instructors:** Who created the content
   - **Reviews:** Ratings from other learners
   - **Related jobs:** Jobs that require these skills

4. **Enroll in learning path:**
   - Click "Enroll" button
   - **Free paths:** Instant enrollment
   - **Premium paths:** Subscribe to Skillsier Pro or pay one-time fee
   - Path added to "Your Learning" dashboard
   - First course unlocked automatically

5. **Take a course:**
   - Navigate to `/app/(tools)/learning/courses/[courseId]`
   - **Course content types:**
     - **Video lessons:** Watch instructor-led videos
     - **Reading materials:** Articles, guides, documentation
     - **Interactive coding:** Code directly in browser (for dev courses)
     - **Design exercises:** Follow along in Figma/Photoshop
     - **Quizzes:** Test understanding after each module
   - **Course features:**
     - Progress tracking (% complete)
     - Bookmarks/notes (mark important sections)
     - Playback speed control (video)
     - Closed captions/transcripts
     - Download resources (PDFs, templates, starter code)
   - **Completion:**
     - Must complete all modules
     - Pass quiz with 70%+ (if applicable)
     - Submit project/assignment (if required)
     - Get feedback on submission
     - Mark course as complete

6. **Track learning progress:**
   - Navigate to `/app/(tools)/learning/progress`
   - **Metrics:**
     - **Total hours learned:** Cumulative learning time
     - **Courses completed:** Count of finished courses
     - **Paths in progress:** Active enrollments
     - **Paths completed:** Fully finished paths with certificates
     - **Streak:** Consecutive days with learning activity
   - **Calendar view:** See learning activity over time (like GitHub contributions)
   - **Upcoming milestones:** Next achievement or certificate

7. **Earn achievements:**
   - Navigate to `/app/(tools)/learning/achievements`
   - **Achievement types:**
     - **Course completion:** Finish any course
     - **Path completion:** Complete full learning path
     - **Skill mastery:** Pass expert-level skills test
     - **Learning streak:** Learn for 7, 30, 100 consecutive days
     - **Speedrun:** Complete path in record time
     - **Perfect quiz:** 100% on all quizzes in a path
     - **Referral teacher:** Refer 3 friends who complete a course
     - **Community helper:** Top contributor in course discussions
     - **Early adopter:** Take brand new course in first week
     - **Polymath:** Complete paths in 3+ different categories
   - **Achievement display:**
     - Badge icon and name
     - Description of how earned
     - Date achieved
     - Rarity (% of freelancers who have it)
   - **Locked achievements:** See future achievements you haven't earned yet (motivational)

8. **Display achievements on profile:**
   - Navigate to `/app/(profile)/achievements`
   - **Manage displayed achievements:**
     - Select up to 6 achievements to feature on profile
     - Reorder by drag-and-drop
     - Choose layout (grid, carousel)
   - Achievements visible to clients when viewing your profile
   - Helps demonstrate continuous learning and dedication

9. **Share progress:**
   - Share course completion, path completion, or achievement on:
     - Your Skillsier profile timeline
     - LinkedIn (with deep link to Skillsier)
     - Twitter
     - Portfolio website
   - Social proof for attracting clients

**Branches & Edge Cases:**

- **Partially completed paths:** Can pause and resume anytime
- **Course expiration:** Premium courses remain accessible for 1 year after purchase (or while subscribed)
- **Path updates:** Paths updated periodically with new content (enrolled users get free updates)
- **Certificate validity:** Learning path certificates don't expire (unlike skills test certs)
- **Course difficulty mismatch:** Can skip beginner courses if you already know the basics
- **Custom paths:** Can create custom learning path by combining individual courses (future)
- **Group learning:** Agencies can assign paths to team members (bulk enrollment)
- **Progress syncing:** Progress syncs across web and mobile
- **Offline courses:** Can download course videos for offline viewing (mobile app)
- **Course discussions:** Can ask questions and discuss with other learners (community forum)
- **Instructor feedback:** Premium paths may include direct feedback from instructors
- **Course prerequisites:** Some courses require completing prior courses first
- **Path branching:** Some paths offer electives (choose 2 of 4 courses to complete)
- **Learning reminders:** Opt-in to daily/weekly reminders to continue learning
- **Achievement notifications:** Push notifications when you earn a new badge
- **Leaderboards:** Opt-in to appear on path completion leaderboards
- **Certificate download:** Download PDF certificate for completed paths
- **Course ratings:** Rate courses and leave feedback after completion
- **Learning analytics:** See which topics you're strong/weak in (personalized insights)
- **Recommended next course:** AI suggests next course based on your progress and goals

**Notifications:**

- **Course completed:** "You completed [Course Name]! [X]% through [Path Name]"
- **Path completed:** "Congratulations! You completed [Path Name] - certificate available"
- **Achievement earned:** "New achievement unlocked: [Achievement Name]"
- **Streak milestone:** "You're on a [X]-day learning streak! Keep it up"
- **New course available:** "New course added to [Path Name] - [Course Title]"
- **Learning reminder:** "Continue your learning journey - [Path Name]"
- **Quiz failed:** "You scored [X]% on [Quiz]. Review and retake?"
- **Assignment feedback:** "Instructor left feedback on your [Assignment]"
- **Course expiring soon:** "Your access to [Course] expires in [X] days"
- **Path recommendation:** "Based on your profile, try [Path Name]"

**Analytics:**

- learning.dashboard\_viewed (active\_paths, completed\_paths)
- learning.paths\_browsed (filters\_applied, category)
- learning.path\_viewed (path\_id, duration, skill\_level, completion\_rate)
- learning.path\_enrolled (path\_id, free\_or\_premium, reason)
- learning.course\_started (course\_id, path\_id, prerequisite\_completed)
- learning.video\_watched (course\_id, video\_duration, completion\_percentage, playback\_speed)
- learning.reading\_completed (course\_id, reading\_duration)
- learning.interactive\_exercise\_completed (course\_id, exercise\_type, time\_taken, correct)
- learning.quiz\_taken (course\_id, score, passing, attempt\_number)
- learning.assignment\_submitted (course\_id, assignment\_type)
- learning.course\_completed (course\_id, time\_to\_complete, quiz\_score)
- learning.path\_completed (path\_id, total\_days, courses\_completed)
- learning.progress\_viewed (total\_hours, courses\_count, paths\_count)
- learning.achievement\_earned (achievement\_id, rarity)
- learning.achievement\_shared (achievement\_id, platform)
- learning.achievements\_viewed (achievements\_count, locked\_count)
- learning.certificate\_downloaded (path\_id)
- learning.course\_bookmarked (course\_id, timestamp)
- learning.course\_note\_added (course\_id, note\_length)
- learning.course\_rated (course\_id, rating, review\_text\_length)
- learning.streak\_continued (streak\_days)
- learning.streak\_broken (streak\_days, last\_activity\_date)
- learning.course\_skipped (course\_id, reason)
- learning.path\_abandoned (path\_id, completion\_percentage, reason)

**Sources:** Inferred from combined-fe-folder-strucure.md ((tools) and (profile) routes) and general platform best practices for learning/gamification features.

---

### FR-12 — Professional Network & Referrals

**Persona:** Freelancer (building connections, getting referrals)

**Preconditions & Triggers:**
- Freelancer wants to grow professional network
- Wants to refer other freelancers to clients (earn referral bonuses)
- Wants to be referred by others for jobs
- Building reputation through connections

**Primary Screens:**
- **Web:**
  - `/app/(profile)/network` — Your professional network
  - `/app/(profile)/network/find-connections` — Discover other freelancers
  - `/app/(profile)/network/[userId]` — View connection's profile
  - `/app/(profile)/network/refer` — Refer a freelancer to a client/job
  - `/app/(profile)/network/requests` — Pending connection requests
  - `/app/(profile)/referrals` — Referral dashboard (track earnings)
  - `/app/(profile)/referrals/invite` — Invite new freelancers to platform

- **Mobile:**
  - `/(profile)/network`
  - `/(profile)/network/find-connections`
  - `/(profile)/network/[userId]`
  - `/(profile)/network/refer`
  - `/(profile)/referrals`

**System Touchpoints:** users-be (connections, referrals), jobs-be (referral matching), financial-be (referral bonuses), communications-be (connection messaging)

**Flow Steps:**

1. **View your network:**
   - Navigate to `/app/(profile)/network`
   - See all your connections (freelancers you're connected with)
   - **Network stats:**
     - Total connections
     - Mutual connections with other freelancers
     - Network reach (1st, 2nd, 3rd degree connections)
     - Referrals made and received
   - **Connection list:**
     - Profile picture and name
     - Headline/specialization
     - Connection strength (how often you interact)
     - Mutual connections count
   - **Recent activity:** See recent activity from your connections (completed projects, earned badges, etc.)

2. **Find new connections:**
   - Navigate to `/app/(profile)/network/find-connections`
   - **Discovery methods:**
     - **People you may know:** Suggested based on mutual connections, similar skills, shared clients
     - **By skills:** Search for freelancers by specific skills
     - **By location:** Find freelancers in your area (for local collaboration)
     - **By category:** Browse by industry (developers, designers, writers, etc.)
     - **Import contacts:** Import from email/LinkedIn to find existing connections
   - **Connection suggestions:**
     - See why you're being recommended to connect (mutual clients, shared skills, same agency, etc.)
     - View profile snippet (skills, rating, bio)

3. **Send connection request:**
   - Click "Connect" on a freelancer's profile
   - **Add personal message (optional but recommended):**
     - "Hi [Name], I noticed we both work in [category]. Would love to connect and potentially collaborate!"
     - "We have [X] mutual connections. Great work on your portfolio!"
   - Send request
   - Request appears in recipient's "Requests" inbox

4. **Manage connection requests:**
   - Navigate to `/app/(profile)/network/requests`
   - **Received requests:**
     - See who wants to connect with you
     - View their profile, mutual connections, skills
     - **Accept:** They become a connection
     - **Ignore:** Request is dismissed (no notification sent)
   - **Sent requests:**
     - See pending requests you've sent
     - Can withdraw request if changed mind

5. **Refer a freelancer to a job:**
   - You see a job that's perfect for a connection in your network
   - Navigate to the job page
   - Click "Refer a Freelancer"
   - **Select freelancer from your network:**
     - System suggests connections whose skills match the job
     - Or search your network manually
   - **Add referral note (visible to client):**
     - "I've worked with [Name] before and they're excellent at [skill]"
     - "Highly recommend [Name] for this project - their [skill] is top-notch"
   - **Notify your connection:**
     - System sends them notification about the referral
     - They can choose to apply or ignore
   - **Referral tracking:**
     - If they apply and get hired, you earn referral bonus
     - Referral bonus (typically 5-10% of first project value, up to $500)

6. **Refer new freelancers to platform:**
   - Navigate to `/app/(profile)/referrals/invite`
   - **Invite methods:**
     - **Email invite:** Enter email addresses, system sends invite
     - **Referral link:** Copy your unique referral link to share anywhere
     - **Social sharing:** Share on LinkedIn, Twitter, Facebook
   - **Referral incentives:**
     - When they sign up using your link and earn their first $100, you get $25 bonus (example)
     - They might get a bonus too (e.g., $10 in connects)
   - **Track referral status:**
     - See who you've invited
     - Who signed up
     - Who completed first project
     - Your earnings from referrals

7. **View referral dashboard:**
   - Navigate to `/app/(profile)/referrals`
   - **Referral stats:**
     - **Freelancers referred:** How many you brought to platform
     - **Jobs referred:** How many job referrals you made
     - **Successful referrals:** How many resulted in hires
     - **Total earnings from referrals:** Lifetime bonus earnings
     - **Pending bonuses:** Referrals in progress (waiting for first project completion)
   - **Referral leaderboard:** Opt-in to appear on monthly top referrers leaderboard

8. **Network value:**
   - **Collaboration opportunities:** Can team up with connections on large projects
   - **Skill sharing:** Learn from connections, share knowledge
   - **Job opportunities:** Connections can refer you to their clients
   - **Social proof:** Clients see your connections (strong network = credibility)
   - **Agency building:** Can invite connections to join your agency
   - **Support system:** Get advice and help from fellow freelancers

**Branches & Edge Cases:**

- **Connection privacy:** Can set network visibility (public, connections only, private)
- **Connection limits:** No hard limit on connections, but quality > quantity
- **Connection removal:** Can remove connections anytime (no notification sent)
- **Blocking:** Can block freelancers (prevents any future connection)
- **Referral fraud:** Fake referrals (self-referrals, spam) result in account suspension
- **Referral bonus caps:** Maximum $X per month in referral bonuses (anti-abuse)
- **Referral timing:** Must refer before the freelancer applies (no retroactive referrals)
- **Multiple referrals:** If multiple people refer same freelancer, first referrer gets bonus
- **Referral attribution:** Referral tracked via cookies/links for 90 days
- **Connection messaging:** Can message connections directly (separate from job messaging)
- **Mutual connections:** Easier to connect if you have mutual connections
- **Connection endorsements:** Can endorse connections for specific skills (like LinkedIn)
- **Referral disclosure:** Client knows if freelancer was referred (transparency)
- **Network recommendations:** System recommends connections who might help you grow
- **Connection strength:** More interaction = stronger connection (affects discovery algorithm)
- **Professional network vs social:** This is professional networking, not social media
- **Connection requests spam:** Limit on how many requests you can send per day (anti-spam)
- **Second-degree connections:** Can see 2nd degree connections (friends of friends)
- **Network growth goals:** Set goals for network size and track progress
- **Connection insights:** See who's viewing your profile from your network

**Notifications:**

- **Connection request received:** "[Name] wants to connect with you"
- **Connection accepted:** "[Name] accepted your connection request"
- **Referral received:** "[Name] referred you to a job - [Job Title]"
- **Referral hired:** "[Name] you referred was hired! Bonus: $X pending"
- **Referral bonus earned:** "$X referral bonus added to your wallet"
- **New freelancer signed up:** "[Name] joined Skillsier using your link"
- **Mutual connection:** "You and [Name] both connected with [Mutual Name]"
- **Connection milestone:** "You've reached 100 connections!"
- **Referral opportunity:** "This job might be perfect for [Connection Name] - refer them?"
- **Network activity:** "[Connection] just completed a project in [Category]"
- **Connection message:** "[Connection Name] sent you a message"

**Analytics:**

- network.network\_page\_viewed (connections\_count, network\_reach)
- network.find\_connections\_viewed (discovery\_method)
- network.connection\_profile\_viewed (user\_id, mutual\_connections\_count)
- network.connection\_request\_sent (user\_id, has\_personal\_message)
- network.connection\_request\_accepted (user\_id, time\_to\_accept)
- network.connection\_request\_ignored (user\_id, reason)
- network.connection\_removed (user\_id, connection\_duration)
- network.user\_blocked (user\_id, reason)
- network.referral\_started (job\_id, referred\_user\_id)
- network.referral\_note\_added (referral\_id, note\_length)
- network.referral\_sent (referral\_id, job\_match\_score)
- network.referral\_accepted (referral\_id, time\_to\_accept)
- network.referral\_applied (referral\_id, proposal\_submitted)
- network.referral\_hired (referral\_id, contract\_value, bonus\_amount)
- network.referral\_bonus\_earned (referral\_id, bonus\_amount)
- network.invite\_sent (method, invitee\_count)
- network.referral\_link\_shared (platform, method)
- network.new\_signup\_from\_referral (referrer\_id, referee\_id)
- network.referral\_dashboard\_viewed (pending\_bonuses, total\_earnings)
- network.connection\_messaged (connection\_id, message\_length)
- network.connection\_endorsed (connection\_id, skill)
- network.network\_goal\_set (target\_connections, deadline)
- network.leaderboard\_viewed (period)

**Sources:** Inferred from combined-fe-folder-strucure.md ((profile) routes) and general best practices for professional networking and referral systems.

---

## D) Communications (Both Roles)

### COM-1 — Inbox: Conversations & Unread

**ID:** COM-1  
**Persona:** Client, Freelancer  
**Preconditions:** User is logged in, has messaging access  
**Triggers:** User navigates to messages inbox, receives new message, searches for conversation

**Primary Screens:**
- **Web:** `/app/(messages)/messages` — Conversation list inbox
- **Web:** `/app/(messages)/messages/[conversationId]` — Conversation detail
- **Mobile:** `/app/(messages)/messages` — Messages inbox
- **Mobile:** `/app/(messages)/messages/[conversationId]` — Conversation thread

**System Touchpoints:**
- **communications-be/conversations** — GET /v1/conversations (list all conversations)
- **communications-be/conversations** — GET /v1/conversations/{conversation_id} (conversation detail)
- **communications-be/message** — GET /v1/conversations/{conversation_id}/messages (messages in conversation)
- **communications-be/conversations** — PUT /v1/conversations/{conversation_id}/read (mark as read)
- **WebSocket** — Real-time message delivery and presence

**Flow Steps:**

1. **User navigates to messages inbox:**
   - Navigate to `/app/(messages)/messages`
   - System loads conversations list
   - **Conversations displayed with:**
     - **Participant info:** Name, avatar, role (client/freelancer)
     - **Last message preview:** Text snippet (truncated), timestamp (relative time)
     - **Unread indicator:** Badge with unread count (if messages unread)
     - **Related context:** Job title, contract reference (if applicable)
     - **Conversation status:** Active, archived, pinned
   - **List sorted by:** Last message timestamp (most recent first)
   - **Pinned conversations:** Appear at top

2. **Unread count management:**
   - **Total unread badge:** Shows on inbox icon in navigation
   - **Per-conversation unread:** Badge next to conversation in list
   - **System counts:** Unread messages per conversation (you haven't read)
   - **Real-time updates:** Unread count updates instantly when new message arrives via WebSocket

3. **Filter conversations:**
   - **All:** Show all conversations
   - **Unread:** Show only conversations with unread messages
   - **Archived:** Show archived conversations (default hidden)
   - **Starred/Pinned:** Show pinned conversations only
   - **By project/job:** Filter by specific job or contract

4. **Search conversations:**
   - Search bar at top of inbox
   - **Search by:** Participant name, message content, job title
   - **Results:** Matching conversations with highlighted keywords
   - System uses communications-be/conversations search endpoint

5. **Select conversation:**
   - Click on conversation in list
   - Navigate to `/app/(messages)/messages/[conversationId]`
   - System loads full conversation thread
   - **Conversation detail shows:**
     - Full message history (paginated, scroll to load more)
     - Participant profile info in header
     - Related job/contract info in sidebar (if applicable)
     - Message composer at bottom

6. **Mark conversation as read:**
   - When user opens conversation, system automatically marks all messages as read
   - Unread badge disappears from conversation list
   - Total unread count decrements
   - Read receipt sent to sender (if enabled)

7. **Conversation actions:**
   - **Pin conversation:** Keeps conversation at top of inbox
   - **Archive conversation:** Removes from main inbox (move to archived)
   - **Mute conversation:** Stop notifications for this conversation
   - **Mark as unread:** Manually mark conversation as unread (even if read)
   - **Delete conversation:** Soft delete (can recover from trash for 30 days)

**Branches & Edge Cases:**

- **Empty inbox:** Show empty state with CTA ("Start a conversation" or "Browse jobs")
- **New user:** Tutorial tooltip explaining inbox features
- **No internet (mobile):** Show cached conversations with "Offline" indicator
- **Deleted participant:** Show conversation as "Deleted User" but preserve messages
- **Spam/blocked:** Filtered conversations don't appear in inbox
- **Archive management:** Archived conversations can be restored to inbox
- **Mark all as read:** Bulk action to mark all conversations as read
- **Conversation pagination:** Load older conversations as user scrolls (infinite scroll)
- **Message preview truncation:** Long messages truncated with "..." in preview
- **Typing indicator:** Show "..." when other participant is typing
- **Presence indicator:** Show online/offline status of participant
- **Conversation grouping:** Group conversations by project/job (optional)
- **Starred conversations:** Can star important conversations for quick access
- **Notification settings:** Per-conversation notification preferences
- **Read receipts:** Optional "seen by" timestamp (privacy setting)
- **Message count:** Total message count per conversation
- **Last activity:** Show "Active 5 minutes ago" or "Last seen yesterday"
- **Conversation export:** Export conversation as PDF/TXT
- **Conversation search within:** Search messages within a specific conversation

**Notifications:**

- **New message:** "[Name] sent you a message" (push, in-app, email)
- **Message read:** "[Name] read your message" (optional, privacy setting)
- **New conversation:** "[Name] started a conversation with you about [Job]"
- **Typing indicator:** Real-time "..." indicator (WebSocket, no notification)
- **Mention:** "[Name] mentioned you in [Job] conversation" (if mentioned via @)

**Analytics:**

- messages.inbox\_viewed (conversation\_count, unread\_count)
- messages.conversation\_opened (conversation\_id, source)
- messages.conversation\_searched (query, results\_count)
- messages.conversation\_filtered (filter\_type)
- messages.conversation\_marked\_read (conversation\_id, unread\_count\_before)
- messages.conversation\_pinned (conversation\_id)
- messages.conversation\_archived (conversation\_id)
- messages.conversation\_muted (conversation\_id)
- messages.conversation\_deleted (conversation\_id, soft\_delete)
- messages.mark\_all\_read\_clicked
- messages.empty\_state\_viewed
- messages.typing\_indicator\_seen

**Sources:** Inferred from combined-fe-folder-structure.md ((messages) routes), communications-be user stories and database design.

---

### COM-2 — Compose/Send/React/Edit/Delete

**ID:** COM-2  
**Persona:** Client, Freelancer  
**Preconditions:** User has an active conversation or can start new one  
**Triggers:** User types and sends message, reacts to message, edits or deletes own message

**Primary Screens:**
- **Web:** `/app/(messages)/messages/[conversationId]` — Message composer within conversation
- **Mobile:** `/app/(messages)/messages/[conversationId]` — Message thread with input

**System Touchpoints:**
- **communications-be/message** — POST /v1/conversations/{conversation_id}/messages (send message)
- **communications-be/message** — PUT /v1/messages/{message_id} (edit message)
- **communications-be/message** — DELETE /v1/messages/{message_id} (delete message)
- **communications-be/message** — POST /v1/messages/{message_id}/reactions (add reaction)
- **communications-be/message** — DELETE /v1/messages/{message_id}/reactions/{reaction_id} (remove reaction)
- **storage-be/uploads** — POST /v1/storage/upload (upload attachments)

**Flow Steps:**

1. **Compose new message:**
   - User types message in message composer (text input at bottom of conversation)
   - **Composer features:**
     - **Text input:** Multi-line text field, auto-expands as user types
     - **Formatting toolbar:** Bold, italic, links (optional, rich text)
     - **Emoji picker:** Click emoji icon to insert emojis
     - **Attachment button:** Add files, images, documents
     - **Mention:** Type "@" to mention participant (autocomplete)
   - **Character limit:** Max 5000 characters per message (system enforces)
   - **Typing indicator:** System sends typing event to other participant (WebSocket)

2. **Send message:**
   - Click "Send" button or press Enter (Shift+Enter for new line)
   - System validates message (not empty, within character limit)
   - **POST /v1/conversations/{conversation_id}/messages:**
     ```json
     {
       "content": "Hi! I reviewed your proposal and I'm impressed...",
       "attachments": ["file_id_1", "file_id_2"],
       "mentions": ["user_id_x"],
       "reply_to_message_id": "uuid" // if replying
     }
     ```
   - Message appears in conversation thread immediately (optimistic UI)
   - **Message status indicators:**
     - **Sending:** Clock icon (message being sent)
     - **Sent:** Single checkmark (message delivered to server)
     - **Delivered:** Double checkmark (message delivered to recipient)
     - **Read:** Blue double checkmark (recipient read message)
     - **Failed:** Red exclamation (retry option)

3. **Message sent confirmation:**
   - Message appears in conversation with timestamp
   - Message ID assigned by backend
   - Other participant receives message via WebSocket (real-time)
   - Notification sent to recipient (push, in-app, email based on preferences)
   - Composer clears, ready for next message

4. **Upload attachments:**
   - Click attachment button
   - **File picker opens:**
     - Select files from device (images, documents, videos)
     - Max file size: 100MB per file (configurable)
     - Max files per message: 10 files
   - **Upload process:**
     - System uploads file to storage-be (POST /v1/storage/upload)
     - Progress bar shows upload progress
     - File appears as preview in composer (thumbnail for images)
   - **File types supported:**
     - Images: JPG, PNG, GIF, WEBP
     - Documents: PDF, DOCX, XLSX, TXT
     - Videos: MP4, MOV
     - Archives: ZIP (if enabled)
   - **Virus scan:** Storage-be scans file for viruses before accepting
   - Once uploaded, file_id attached to message

5. **React to message:**
   - User hovers over any message in thread
   - **Reaction button appears:** Emoji icon
   - Click to open reaction picker
   - **Reaction options:** 👍 ❤️ 😂 😮 😢 🎉 (common emojis)
   - Select emoji to react
   - **System adds reaction:**
     - POST /v1/messages/{message_id}/reactions
     - Reaction appears below message with user's avatar
     - If multiple users react with same emoji, count increments
   - **Remove reaction:** Click same emoji again to unreact

6. **Edit message:**
   - User can edit their own messages within 15 minutes of sending
   - Hover over own message, click "Edit" (three dots menu)
   - Message content becomes editable in place
   - Make changes, click "Save" or press Enter
   - **PUT /v1/messages/{message_id}:**
     ```json
     {
       "content": "Hi! I reviewed your proposal and I'm very impressed..."
     }
     ```
   - Message updates in real-time for all participants
   - **"(edited)" label:** Appears next to timestamp
   - **Edit history:** Clicking "(edited)" shows edit history (audit trail)

7. **Delete message:**
   - User can delete their own messages
   - Hover over own message, click "Delete" (three dots menu)
   - **Confirmation dialog:** "Are you sure you want to delete this message?"
   - **Delete options:**
     - **Delete for me:** Message removed from your view only
     - **Delete for everyone:** Message removed for all participants (if within 1 hour)
   - **DELETE /v1/messages/{message_id}** (with soft_delete flag)
   - Message replaced with "[Message deleted]" placeholder
   - Attachments remain in storage but are inaccessible

**Branches & Edge Cases:**

- **Empty message:** Send button disabled until user types something
- **Attachment upload failed:** Show error, allow retry
- **Virus detected:** File rejected, user notified
- **Message send failed:** Show error, "Retry" button to resend
- **Edit time limit:** Can't edit after 15 minutes (system enforces)
- **Delete time limit (everyone):** Can't delete for everyone after 1 hour
- **Deleted message:** Shows "[Message deleted]" with timestamp
- **Reply to message:** Can reply to specific message (creates thread)
- **Quote message:** Can quote part of message when replying
- **Draft messages:** Auto-save draft as user types (recover on page reload)
- **Offline send (mobile):** Message queued, sent when online
- **Attachment preview:** Image attachments show preview, documents show file icon
- **Link preview:** URLs in messages auto-generate preview cards
- **Message length:** Very long messages truncated with "Show more" button
- **Multiple reactions:** Users can add multiple different reactions
- **Reaction limit:** Max 50 reactions per message (anti-spam)
- **Mention notification:** Mentioned user gets notification
- **Message formatting:** Support basic markdown (if rich text enabled)
- **Paste images:** Can paste images directly into composer (Ctrl+V)
- **Voice messages (mobile):** Record and send voice notes (if enabled)
- **Read receipts:** Optional "seen" timestamp below message
- **Message delivery:** Failed delivery shows error icon
- **Network issues:** Show "Connecting..." if WebSocket disconnected
- **Edit indication:** Show who edited and when (hover over "(edited)")

**Notifications:**

- **New message:** "[Name] sent you a message: [Preview]"
- **Mention:** "[Name] mentioned you in a message"
- **Reaction:** "[Name] reacted ❤️ to your message"
- **Message edited:** No notification (to avoid spam)
- **Message deleted:** No notification

**Analytics:**

- messages.message\_composed (character\_count, has\_attachments, has\_mentions)
- messages.message\_sent (conversation\_id, message\_length, attachments\_count)
- messages.message\_send\_failed (error\_type)
- messages.attachment\_uploaded (file\_type, file\_size)
- messages.attachment\_upload\_failed (error\_type)
- messages.reaction\_added (message\_id, reaction\_emoji)
- messages.reaction\_removed (message\_id, reaction\_emoji)
- messages.message\_edited (message\_id, edit\_count, time\_since\_send)
- messages.message\_deleted (message\_id, delete\_type, time\_since\_send)
- messages.typing\_indicator\_sent
- messages.emoji\_picker\_opened
- messages.mention\_added (mentioned\_user\_id)
- messages.draft\_auto\_saved
- messages.message\_retry\_clicked

**Sources:** Inferred from combined-fe-folder-structure.md ((messages) routes), communications-be user stories and database design.

---

### COM-3 — Threads/Replies, Mentions, Pins

**ID:** COM-3  
**Persona:** Client, Freelancer  
**Preconditions:** User is in an active conversation  
**Triggers:** User replies to a specific message, mentions another user, pins a message

**Primary Screens:**
- **Web:** `/app/(messages)/messages/[conversationId]` — Conversation with threading
- **Mobile:** `/app/(messages)/messages/[conversationId]` — Message thread with replies

**System Touchpoints:**
- **communications-be/message** — POST /v1/messages/{message_id}/replies (reply to message)
- **communications-be/mention** — POST /v1/mentions (create mention)
- **communications-be/message** — PUT /v1/messages/{message_id}/pin (pin message)
- **communications-be/conversations** — GET /v1/conversations/{conversation_id}/pinned (get pinned messages)

**Flow Steps:**

1. **Reply to message (Create thread):**
   - User hovers over any message in conversation
   - **"Reply" button appears** (or swipe right on mobile)
   - Click "Reply"
   - **Thread indicator:** Original message gets highlighted, shown as "parent" message
   - Composer shows: "Replying to [Name]: [Message preview]" (with X to cancel)
   - User types reply message
   - Send reply:
     - POST /v1/messages/{parent_message_id}/replies
     - Reply linked to parent message via reply_to_message_id
   - **Thread visualization:**
     - Reply appears indented below parent message
     - Thread line connects reply to parent
     - Parent message shows "X replies" count

2. **View thread:**
   - Click on parent message or "X replies" count
   - **Thread view opens:**
     - Shows parent message at top
     - All replies below (chronologically)
     - Reply composer at bottom
   - **Thread can be expanded in place or opened in side panel** (web)
   - **Mobile:** Thread opens in new screen with back button

3. **Nested threading:**
   - **Max thread depth:** 2 levels (parent → reply → sub-reply)
   - After 2 levels, all further replies are flat at level 2
   - **Thread collapse:** Can collapse/expand thread by clicking on parent message

4. **Mention user (@mention):**
   - While composing message, type "@"
   - **Autocomplete dropdown appears:**
     - Lists conversation participants
     - Filter by name as user types
   - Select user to mention
   - User's name inserted as "@[Name]" (clickable mention)
   - **POST /v1/mentions:**
     ```json
     {
       "message_id": "uuid",
       "mentioned_user_id": "uuid",
       "mention_text": "@John"
     }
     ```
   - Mentioned user receives notification
   - Message text highlights mention (e.g., blue color)

5. **View mentions:**
   - **Mentions inbox:** Separate view for all messages where user was mentioned
   - Navigate to `/app/(messages)/messages?filter=mentions`
   - Lists all messages containing @mentions of current user
   - Click mention to jump to conversation at that message

6. **Pin message:**
   - User hovers over important message
   - Click "Pin" (pushpin icon)
   - **Confirmation:** "Pin this message?"
   - **PUT /v1/messages/{message_id}/pin:**
     ```json
     {
       "pinned": true
     }
     ```
   - **Pinned message:**
     - Appears at top of conversation (above all messages)
     - Shows "Pinned by [Name]" label
     - Pinned icon next to message
   - **Max pinned messages:** 3 per conversation (anti-clutter)

7. **View pinned messages:**
   - **Pinned section:** At top of conversation thread
   - Shows all pinned messages in carousel (if multiple)
   - Click pin icon in header to see all pinned messages in modal
   - **GET /v1/conversations/{conversation_id}/pinned**

8. **Unpin message:**
   - Click "Unpin" on pinned message
   - Message moves back to original position in thread
   - PUT /v1/messages/{message_id}/pin (pinned: false)

**Branches & Edge Cases:**

- **Reply notification:** Original message author gets notification
- **Thread depth limit:** After 2 levels, replies become flat
- **Thread indicators:** Visual lines showing thread structure
- **Jump to parent:** Can click to jump from reply to parent message
- **Thread count:** Parent message shows "X replies" badge
- **Empty thread:** If all replies deleted, thread collapses
- **Mention yourself:** Can't mention yourself (system prevents)
- **Mention multiple:** Can mention multiple users in one message
- **Mention from reply:** Can mention in threaded replies
- **Mention privacy:** Only participants in conversation can be mentioned
- **Pinned message access:** All conversation participants see pinned messages
- **Pin permissions:** Both parties can pin/unpin (no ownership restriction)
- **Pin limit reached:** Show error "Max 3 pinned messages" if trying to pin 4th
- **Pinned message deleted:** Unpin automatically if message deleted
- **Thread notification:** Receive notification for replies to your messages
- **Thread search:** Search within thread
- **Quote vs reply:** Reply creates thread, quote copies text into new message
- **Thread in mobile:** Threads visualized differently (less space)
- **Mention in edited message:** Can add/remove mentions when editing
- **Mention formatting:** Mentions styled differently (blue, bold)
- **Thread collapse:** Auto-collapse old threads to save space
- **Jump to latest reply:** Button to jump to most recent reply in thread

**Notifications:**

- **Reply received:** "[Name] replied to your message: [Preview]"
- **Mention:** "[Name] mentioned you: [Preview]"
- **Message pinned:** "[Name] pinned a message" (in-app only, no push)
- **Thread activity:** "New reply in [Job] conversation" (if many replies)

**Analytics:**

- messages.reply\_button\_clicked (message\_id)
- messages.reply\_sent (parent\_message\_id, thread\_depth)
- messages.thread\_viewed (thread\_id, reply\_count)
- messages.thread\_collapsed (thread\_id)
- messages.thread\_expanded (thread\_id)
- messages.mention\_autocomplete\_shown
- messages.mention\_selected (mentioned\_user\_id)
- messages.mention\_clicked (message\_id, mentioned\_user\_id)
- messages.mentions\_inbox\_viewed (mention\_count)
- messages.message\_pinned (message\_id, conversation\_id)
- messages.message\_unpinned (message\_id)
- messages.pinned\_messages\_viewed (pinned\_count)
- messages.jump\_to\_parent\_clicked
- messages.jump\_to\_latest\_reply\_clicked

**Sources:** Inferred from combined-fe-folder-structure.md ((messages) routes), communications-be user stories for threads, mentions, and message management.

---

### COM-4 — Calls & Screen Share

**ID:** COM-4  
**Persona:** Client, Freelancer  
**Preconditions:** User is in conversation, has camera/mic permissions (mobile), WebRTC supported  
**Triggers:** User initiates voice/video call, joins ongoing call, shares screen

**Primary Screens:**
- **Web:** `/app/(messages)/messages/[conversationId]/call` — In-call interface
- **Mobile:** `/app/(messages)/messages/[conversationId]/call` — Call screen

**System Touchpoints:**
- **communications-be/call** — POST /v1/conversations/{conversation_id}/calls (start call)
- **communications-be/call** — PUT /v1/calls/{call_id}/join (join call)
- **communications-be/call** — PUT /v1/calls/{call_id}/end (end call)
- **WebRTC signaling server** — P2P connection establishment
- **WebSocket** — Call events (incoming call, call ended, etc.)

**Flow Steps:**

1. **Start call:**
   - User clicks "Call" button in conversation header
   - **Call type selection:**
     - **Voice call:** Audio only
     - **Video call:** Audio + video
   - **System checks:**
     - Camera/mic permissions (browser prompt)
     - Network connectivity
     - Participant availability (online status)
   - **POST /v1/conversations/{conversation_id}/calls:**
     ```json
     {
       "call_type": "VIDEO", // or "AUDIO"
       "participants": ["user_id_1", "user_id_2"]
     }
     ```
   - **Calling state:** System shows "Calling [Name]..." with ringing animation

2. **Receive call (incoming):**
   - Other participant receives call notification via WebSocket
   - **Call notification:**
     - In-app banner: "[Name] is calling you..." (video/audio icon)
     - Browser notification: "Incoming call from [Name]"
     - Mobile: Full-screen incoming call UI with ringtone
   - **Call actions:**
     - **Answer:** Accept call (audio or video)
     - **Decline:** Reject call (sends busy signal)
     - **Ignore:** Let it ring (timeout after 30 seconds)

3. **Answer call:**
   - Participant clicks "Answer"
   - **PUT /v1/calls/{call_id}/join**
   - System establishes WebRTC connection
   - **Signaling:** Exchange ICE candidates, SDP offer/answer
   - **Connection established:** Both parties see/hear each other
   - Navigate to `/app/(messages)/messages/[conversationId]/call`

4. **In-call interface:**
   - **Video layout:**
     - **Large video:** Remote participant (person you're calling)
     - **Small PiP video:** Your own video (self-view, draggable)
     - **Mobile:** Full screen video, swipe for controls
   - **Audio call:** Shows avatars instead of video
   - **Call info banner:** Duration timer, participant name, connection quality
   - **Call controls (toolbar at bottom):**
     - **Mute/Unmute mic:** Toggle audio
     - **Camera on/off:** Toggle video
     - **Screen share:** Share your screen
     - **Chat:** Open text chat sidebar (without leaving call)
     - **Settings:** Audio/video device settings
     - **End call:** Hang up (red button)
   - **Connection quality indicator:** Green/Yellow/Red dot (based on RTT, packet loss)

5. **Screen share:**
   - Click "Share Screen" button during call
   - **Screen picker dialog:** Select screen, window, or browser tab to share
   - System requests screen capture permission (browser/OS)
   - Once approved, remote participant sees your screen
   - **Screen share indicator:** Banner showing "You're sharing your screen"
   - **Stop sharing:** Click "Stop Sharing" to end screen share
   - **Mobile:** Screen share less common (limited by OS restrictions)

6. **Call controls:**
   - **Mute/unmute:** Click mic icon (or spacebar hotkey)
   - **Video on/off:** Click camera icon
   - **Switch camera (mobile):** Toggle front/back camera
   - **Speaker/earpiece:** Toggle speaker mode (mobile)
   - **Add participant:** Invite another user to call (3-way call)
   - **Record call:** Record audio/video (if enabled, with consent)

7. **End call:**
   - Click "End Call" (red button)
   - **PUT /v1/calls/{call_id}/end**
   - System terminates WebRTC connection
   - **Call summary:**
     - Duration: "Call lasted 15m 23s"
     - Participants: Who was on call
     - Recording (if recorded): Link to download
   - Call summary saved in conversation thread as system message

**Branches & Edge Cases:**

- **No answer:** If recipient doesn't answer within 30s, call auto-cancels
- **Busy:** If recipient is already on another call, show "User is busy"
- **Rejected call:** Show "Call declined"
- **Network issues:** Show "Reconnecting..." if connection drops
- **Call quality:** Auto-adjust video quality based on bandwidth
- **Mic/camera permissions denied:** Show error, guide user to grant permissions
- **Screen share not supported:** Disable button on unsupported browsers
- **Call history:** Past calls visible in conversation as system messages
- **Missed call:** Show notification "Missed call from [Name]"
- **Callback:** Click missed call notification to call back
- **Background call (mobile):** Minimize to PiP, continue call while using other apps
- **Call in progress indicator:** Show banner when call active
- **Multiple calls:** Can't start new call while already in one
- **Group calls:** Support 3+ participants (up to 10)
- **Call recording consent:** Both parties must consent before recording
- **Recording notification:** Show banner "Call is being recorded"
- **Call encryption:** WebRTC uses DTLS-SRTP (end-to-end encrypted)
- **Echo cancellation:** Automatic echo/noise suppression
- **Call dropped:** Auto-reconnect if connection lost briefly
- **Bluetooth devices:** Support Bluetooth headsets (mobile)
- **Speaker phone:** Toggle between earpiece and speaker
- **Call feedback:** After call, option to rate call quality

**Notifications:**

- **Incoming call:** "[Name] is calling you" (push, in-app, ringtone)
- **Missed call:** "Missed call from [Name]" (push, in-app)
- **Call ended:** "Call ended (15m 23s)" (in-app only)
- **Screen share started:** "[Name] is sharing their screen" (in-app)
- **Recording started:** "Call recording started" (in-app, with consent)

**Analytics:**

- calls.call\_button\_clicked (call\_type)
- calls.call\_initiated (call\_type, conversation\_id)
- calls.call\_answered (call\_id, time\_to\_answer)
- calls.call\_declined (call\_id, reason)
- calls.call\_missed (call\_id)
- calls.call\_connected (call\_id, connection\_time)
- calls.call\_ended (call\_id, duration, end\_reason)
- calls.mic\_toggled (call\_id, muted)
- calls.camera\_toggled (call\_id, enabled)
- calls.screen\_share\_started (call\_id)
- calls.screen\_share\_ended (call\_id, duration)
- calls.call\_quality\_poor (call\_id, metrics)
- calls.call\_reconnect\_attempted (call\_id)
- calls.call\_recording\_started (call\_id, consent\_given)
- calls.participant\_added (call\_id, participant\_count)
- calls.call\_feedback\_submitted (call\_id, rating)

**Sources:** Inferred from combined-fe-folder-structure.md ((messages) routes), communications-be user stories for real-time communication and WebRTC integration.

---

### COM-5 — Message Bookmarks & Drafts

**ID:** COM-5  
**Persona:** Client, Freelancer  
**Preconditions:** User has active conversations  
**Triggers:** User bookmarks message for later reference, system auto-saves draft

**Primary Screens:**
- **Web:** `/app/(messages)/messages/[conversationId]` — Conversation with bookmark option
- **Web:** `/app/(messages)/messages/bookmarks` — Bookmarked messages list
- **Mobile:** `/app/(messages)/messages/[conversationId]` — Message actions with bookmark
- **Mobile:** `/app/(messages)/messages/drafts` — Draft messages

**System Touchpoints:**
- **communications-be/message** — POST /v1/messages/{message_id}/bookmark (bookmark message)
- **communications-be/message** — GET /v1/users/me/bookmarks (get bookmarks)
- **communications-be/message** — POST /v1/messages/drafts (save draft)
- **communications-be/message** — GET /v1/messages/drafts (get drafts)

**Flow Steps:**

1. **Bookmark message:**
   - User hovers over message (or long-press on mobile)
   - **"Bookmark" option appears** (star icon)
   - Click to bookmark
   - **POST /v1/messages/{message_id}/bookmark:**
     ```json
     {
       "message_id": "uuid",
       "user_id": "uuid",
       "note": "Important contract detail" // optional
     }
     ```
   - Star icon becomes filled (golden)
   - Success toast: "Message bookmarked"

2. **View bookmarks:**
   - Navigate to `/app/(messages)/messages/bookmarks`
   - **Bookmarked messages list:**
     - Grouped by conversation
     - Shows message preview, sender, timestamp
     - Optional note attached to bookmark
     - Link to jump to original message in conversation
   - **Sort options:** By date, by conversation, by note
   - **Search bookmarks:** Search within bookmarked messages

3. **Remove bookmark:**
   - Click bookmark star icon again (on message or in bookmarks list)
   - Star becomes unfilled
   - Message removed from bookmarks list
   - DELETE /v1/messages/{message_id}/bookmark

4. **Add note to bookmark:**
   - When bookmarking, option to add note: "Why is this important?"
   - Note helps remember context later
   - Note visible only to user who bookmarked (private)

5. **Auto-save draft:**
   - User starts typing message in composer
   - **Auto-save triggers:**
     - After 2 seconds of inactivity
     - When navigating away from conversation
     - When closing browser tab
   - **POST /v1/messages/drafts:**
     ```json
     {
       "conversation_id": "uuid",
       "content": "Hey, I wanted to discuss...",
       "attachments": ["file_id_1"],
       "reply_to_message_id": "uuid" // if replying
     }
     ```
   - Draft saved in backend (persists across devices)

6. **View drafts:**
   - Navigate to `/app/(messages)/messages/drafts`
   - **Drafts list:**
     - Shows conversation name
     - Draft preview (truncated text)
     - Timestamp ("Draft saved 5 minutes ago")
     - "Continue" button to resume writing
   - **GET /v1/messages/drafts**

7. **Resume draft:**
   - Click on draft in list
   - System navigates to conversation
   - Composer auto-fills with draft content
   - User can edit and send or discard

8. **Discard draft:**
   - Click "X" or "Discard" on draft
   - **Confirmation:** "Are you sure you want to delete this draft?"
   - DELETE /v1/messages/drafts/{draft_id}
   - Draft removed from list

**Branches & Edge Cases:**

- **Bookmark limit:** Max 500 bookmarks per user (system enforces)
- **Bookmark overflow:** If limit reached, show warning "Remove old bookmarks"
- **Bookmarked message deleted:** Bookmark remains but shows "[Message deleted]"
- **Export bookmarks:** Can export all bookmarks as PDF/TXT
- **Bookmark search:** Search by keyword within bookmarks
- **Bookmark categories:** Tag bookmarks (e.g., "Important", "To-do", "Reference")
- **Draft expiry:** Drafts auto-delete after 30 days (configurable)
- **Multiple drafts:** One draft per conversation (latest overwrites)
- **Draft sync:** Drafts sync across devices (web + mobile)
- **Draft conflict:** If typing on two devices, latest draft wins
- **Draft notification:** Show "Draft saved" indicator in composer
- **Draft on send:** Draft auto-deleted when message sent
- **Draft recovery:** Can recover drafts deleted within 24h (undo)
- **Bookmark sharing:** Can't share bookmarks with others (personal)
- **Bookmarks in sidebar:** Quick access to bookmarks from sidebar
- **Draft preview:** Hover to see full draft content
- **Empty drafts:** Empty drafts (no text) not saved
- **Attachment in draft:** Attachments in drafts remain until sent or discarded

**Notifications:**

- **No notifications** for bookmarks or drafts (user-initiated actions)
- **Draft reminder (optional):** "You have 3 unsent drafts" (weekly email)

**Analytics:**

- messages.message\_bookmarked (message\_id, conversation\_id, has\_note)
- messages.bookmark\_note\_added (bookmark\_id, note\_length)
- messages.bookmarks\_viewed (bookmark\_count)
- messages.bookmark\_removed (bookmark\_id, days\_since\_bookmarked)
- messages.bookmark\_searched (query)
- messages.bookmarks\_exported (format)
- messages.draft\_auto\_saved (conversation\_id, character\_count)
- messages.draft\_viewed (draft\_id)
- messages.draft\_resumed (draft\_id, days\_since\_saved)
- messages.draft\_discarded (draft\_id)
- messages.drafts\_page\_viewed (draft\_count)

**Sources:** Inferred from combined-fe-folder-structure.md ((messages) routes), general messaging best practices for bookmarks and draft management.

---

### COM-6 — Read Receipts & Delivery Status

**ID:** COM-6  
**Persona:** Client, Freelancer  
**Preconditions:** User sends messages in conversations  
**Triggers:** Message sent, message delivered, message read

**Primary Screens:**
- **Web:** `/app/(messages)/messages/[conversationId]` — Message with delivery indicators
- **Mobile:** `/app/(messages)/messages/[conversationId]` — Message thread with read receipts

**System Touchpoints:**
- **communications-be/read_receipt** — POST /v1/messages/{message_id}/read (mark message as read)
- **communications-be/read_receipt** — GET /v1/messages/{message_id}/receipts (get read receipts)
- **communications-be/delivery** — GET /v1/messages/{message_id}/delivery-status (delivery status)
- **communications-be/read_state** — POST /v1/conversations/{conversation_id}/mark-read (mark conversation as read up to sequence)

**Flow Steps:**

1. **Message sent status:**
   - User sends message
   - Message shows **"Sending..."** indicator (gray clock icon)
   - **POST /v1/messages** succeeds
   - Status changes to **"Sent"** (single gray checkmark)
   - System queues delivery to recipient devices

2. **Message delivered status:**
   - Recipient's device acknowledges delivery
   - **delivery.acknowledged.v1** event received
   - Sender sees **"Delivered"** (double gray checkmarks)
   - **GET /v1/messages/{message_id}/delivery-status:**
     ```json
     {
       "message_id": "uuid",
       "status": "DELIVERED",
       "sent_at": "2025-11-06T10:00:00Z",
       "delivered_at": "2025-11-06T10:00:02Z",
       "device_count": 2,
       "devices_delivered": 2
     }
     ```

3. **Message read status:**
   - Recipient opens conversation and views message
   - System calls **POST /v1/conversations/{conversation_id}/mark-read:**
     ```json
     {
       "last_read_message_seq": 45,
       "user_id": "recipient_uuid"
     }
     ```
   - Read receipt created: **POST /v1/messages/{message_id}/read**
   - Sender sees **"Read"** (double blue checkmarks)
   - Timestamp appears: "Read 2 minutes ago"

4. **View detailed read receipts:**
   - In group conversations, hover over message checkmarks
   - **Tooltip shows:**
     - "Read by: John (2m ago), Sarah (5m ago)"
     - "Delivered to: Mike, Lisa"
     - "Not delivered: Alex (offline)"
   - **GET /v1/messages/{message_id}/receipts:**
     ```json
     {
       "message_id": "uuid",
       "receipts": [
         {
           "user_id": "uuid",
           "user_name": "John Doe",
           "status": "READ",
           "read_at": "2025-11-06T10:02:00Z"
         },
         {
           "user_id": "uuid",
           "user_name": "Sarah Smith",
           "status": "DELIVERED",
           "delivered_at": "2025-11-06T10:00:05Z"
         }
       ]
     }
     ```

5. **Group conversation read indicators:**
   - Message shows **participant count badge** (e.g., "2/5 read")
   - Click to expand: "Read by John, Sarah" / "Unread by Mike, Lisa, Alex"
   - Real-time updates as more participants read

6. **Read receipt privacy settings:**
   - User can disable read receipts in settings
   - Navigate to `/app/settings/privacy`
   - Toggle **"Send read receipts"** off
   - **PUT /v1/users/me/settings/privacy:**
     ```json
     {
       "send_read_receipts": false
     }
     ```
   - User's read status no longer visible to others
   - User still sees others' read receipts

7. **Delivery retry for failed messages:**
   - Message fails to deliver (network error, offline recipient)
   - Status shows **"Not delivered"** (red exclamation icon)
   - System retries with exponential backoff (5 attempts)
   - User can manually retry: Click "Retry" button
   - **POST /v1/messages/{message_id}/retry**

8. **Read pointer synchronization:**
   - System maintains **read pointer** (last read sequence number)
   - When user scrolls and views messages, pointer advances
   - **POST /v1/conversations/{conversation_id}/mark-read** called
   - Unread count decreases in inbox
   - Read pointer syncs across devices

**Branches & Edge Cases:**

- **Offline recipient:** Delivery pending until recipient comes online
- **Multiple devices:** Delivered when acknowledged by any device
- **Read receipts disabled:** Shows "Delivered" max, never "Read"
- **Deleted message:** Read receipts retained but message shows "[Deleted]"
- **Group read threshold:** In groups >10, show "Read by 8 people" instead of names
- **Read receipt race:** If multiple participants read simultaneously, deduplicate events
- **Delivery failure permanent:** After 5 retries, mark "Failed to deliver"
- **Read receipt compaction:** System compacts old receipts derived from read pointers
- **Sender blocked recipient:** No delivery status shown (privacy)
- **Read in notification:** Marking as read from notification updates read pointer
- **Partial delivery:** Some devices delivered, some failed (show mixed status)
- **Read receipt delay:** Network delay may cause late read receipt arrival
- **Bulk mark as read:** Mark entire conversation as read (all messages up to latest)
- **Read receipt revocation:** Can't unmark as read once marked
- **Delivery status for attachments:** Separate status for attachment downloads
- **Read in thread:** Thread replies also trigger read receipts
- **Read receipt aggregation:** Daily job compacts redundant receipts
- **Read status in search:** Search results show read/unread indicator
- **Read receipt notification:** Optional notification when message is read
- **Anonymous read:** In public channels, read receipts not shown

**Notifications:**

- **Push (optional):** "John read your message" (if enabled in settings)
- **In-app:** Real-time update of checkmark status (no explicit notification)

**Analytics:**

- messages.message\_sent (message\_id, conversation\_id)
- messages.message\_delivered (message\_id, delivery\_time\_ms)
- messages.message\_read (message\_id, time\_to\_read\_seconds)
- messages.read\_receipt\_viewed (message\_id, receipt\_count)
- messages.delivery\_status\_checked (message\_id)
- messages.delivery\_retry\_clicked (message\_id, attempt\_count)
- messages.read\_receipts\_disabled (user\_id)
- messages.read\_receipts\_enabled (user\_id)
- messages.bulk\_mark\_read (conversation\_id, message\_count)
- messages.read\_pointer\_advanced (conversation\_id, sequence\_advance)
- messages.delivery\_failed\_permanent (message\_id, reason)
- messages.group\_read\_status\_expanded (message\_id, participant\_count)

**Sources:** communications-be.user-stories.md (Section 10: Delivery Status, Section 1.2: Participant State), communications-be.database-design.md (read_receipts table, delivery status), combined-fe-folder-structure.md ((messages) routes, ReadReceipt.tsx).

---

### COM-7 — Message Search & Filtering

**ID:** COM-7  
**Persona:** Client, Freelancer  
**Preconditions:** User has message history  
**Triggers:** User searches for specific messages or filters conversations

**Primary Screens:**
- **Web:** `/app/(messages)/messages` — Messages inbox with search bar
- **Web:** `/app/(messages)/messages/search` — Search results page
- **Mobile:** `/app/(messages)/messages` — Search overlay

**System Touchpoints:**
- **communications-be/message** — GET /v1/messages/search (full-text search)
- **communications-be/conversation** — GET /v1/conversations/search (conversation search)
- **communications-be/message** — GET /v1/messages (with filters: date range, participant, attachment type, etc.)

**Flow Steps:**

1. **Open search interface:**
   - User clicks **search icon** in messages header
   - Search bar expands (web) or search overlay appears (mobile)
   - Placeholder text: "Search messages, files, or people..."
   - Recent searches shown below (if any)

2. **Execute basic text search:**
   - User types query: "project deadline"
   - System shows **live suggestions** as user types (after 3 characters)
   - Press Enter or click "Search"
   - **GET /v1/messages/search?q=project+deadline&limit=50:**
     ```json
     {
       "query": "project deadline",
       "total_count": 127,
       "results": [
         {
           "message_id": "uuid",
           "conversation_id": "uuid",
           "conversation_name": "Project Alpha",
           "sender_name": "John Doe",
           "body_snippet": "...the project deadline is next Friday...",
           "sent_at": "2025-10-15T10:30:00Z",
           "matched_terms": ["project", "deadline"],
           "has_attachments": false
         }
       ],
       "facets": {
         "conversations": [...],
         "senders": [...],
         "date_ranges": [...]
       }
     }
     ```

3. **View search results:**
   - **Results displayed:**
     - Message snippet with matched terms highlighted
     - Conversation name and sender
     - Timestamp
     - Attachment icons if present
   - **Grouped by conversation** (optional toggle)
   - Click result to **jump to message** in original conversation
   - Message highlighted with yellow background (fades after 2 seconds)

4. **Apply filters:**
   - **Filter panel** on left (web) or bottom sheet (mobile)
   - **Available filters:**
     - **From:** Select participant(s)
     - **In conversation:** Select conversation(s)
     - **Date range:** Today, Last week, Last month, Custom range
     - **Has attachments:** Images, Documents, Links, Files
     - **Message type:** Text, System messages, Calls
     - **Status:** Unread, Starred, Bookmarked
   - Select filter: "From: John Doe", "Date: Last week", "Has: Documents"
   - **GET /v1/messages/search?q=project+deadline&from=john_uuid&start_date=2025-11-01&has_attachments=document**
   - Results update instantly

5. **Advanced search syntax:**
   - User can use **operators:**
     - **Exact phrase:** "project deadline" (with quotes)
     - **Exclude:** deadline -postponed (exclude term)
     - **OR:** deadline OR milestone (either term)
     - **From:** from:john@email.com
     - **In:** in:"Project Alpha"
     - **Has:** has:attachment, has:link, has:image
     - **Before/After:** before:2025-11-01, after:2025-10-01
   - Example: `"project deadline" from:john has:document before:2025-11-15`
   - Advanced syntax parsed and applied as filters

6. **Save search:**
   - After executing search with filters
   - Click **"Save this search"** button
   - Enter name: "Project deadline discussions"
   - **POST /v1/messages/searches/saved:**
     ```json
     {
       "name": "Project deadline discussions",
       "query": "project deadline",
       "filters": {
         "from": "john_uuid",
         "has_attachments": "document"
       }
     }
     ```
   - Saved search added to sidebar
   - Click saved search to re-run instantly

7. **Search within conversation:**
   - User is in conversation `/app/(messages)/messages/[conversationId]`
   - Opens search: "Search in this conversation"
   - **GET /v1/conversations/{conversation_id}/messages/search?q=budget**
   - Results limited to current conversation
   - Navigation arrows: "Previous match" / "Next match"
   - Match count shown: "3 of 12 matches"

8. **Search attachments & files:**
   - Filter: "Has: Documents"
   - **Search file names:** "invoice.pdf"
   - **Search file content (OCR):** Text extracted from images/PDFs (if available)
   - Results show file preview thumbnail
   - Click to download or preview file

**Branches & Edge Cases:**

- **No results:** Show "No messages found. Try different keywords."
- **Too many results:** Paginated results (50 per page)
- **Search timeout:** If search takes >5s, show progress spinner
- **Invalid syntax:** Parse error → show "Invalid search syntax" with hint
- **Empty query:** Can't search without query (button disabled)
- **Recent searches:** Show last 5 searches, click to re-run
- **Clear recent searches:** Option to clear search history
- **Search deleted messages:** Excluded from results by default (admin can include)
- **Search archived conversations:** Option to include archived
- **Encrypted messages:** Search only works on unencrypted body (metadata searchable)
- **Multilingual search:** Supports UTF-8, emoji, special characters
- **Fuzzy matching:** Typo tolerance (optional setting)
- **Search ranking:** Results ranked by relevance (TF-IDF or ML)
- **Search highlighting:** Multiple terms highlighted in different colors
- **Search export:** Export search results as CSV or JSON
- **Search limits:** Rate limit: 100 searches/hour per user
- **Search cache:** Recent searches cached for 5 minutes (performance)
- **Search suggestions:** Autocomplete from message index
- **Search shortcuts:** Keyboard shortcuts (Ctrl+F, Cmd+F)
- **Search mobile optimization:** Minimal UI, touch-friendly

**Notifications:**

- **No notifications** for search actions (user-initiated)
- **Optional:** "New messages matching saved search" (weekly digest)

**Analytics:**

- messages.search\_executed (query, filter\_count, result\_count)
- messages.search\_result\_clicked (message\_id, result\_position)
- messages.search\_filter\_applied (filter\_type, filter\_value)
- messages.search\_saved (search\_name)
- messages.saved\_search\_executed (search\_id)
- messages.search\_within\_conversation (conversation\_id, query)
- messages.search\_no\_results (query)
- messages.search\_timeout (query, duration\_ms)
- messages.advanced\_search\_used (syntax\_operators)
- messages.search\_attachment\_filtered (attachment\_type)
- messages.search\_cleared (query)
- messages.recent\_search\_clicked (query)
- messages.search\_exported (format, result\_count)

**Sources:** communications-be.user-stories.md (Section 8: Search & Filtering), combined-fe-folder-structure.md ((messages) routes, search functionality).

---

### COM-8 — Conversation Archiving & Retention

**ID:** COM-8  
**Persona:** Client, Freelancer  
**Preconditions:** User has active or completed conversations  
**Triggers:** User wants to clean up inbox, conversation completed, retention policy triggers

**Primary Screens:**
- **Web:** `/app/(messages)/messages/[conversationId]` — Conversation with archive option
- **Web:** `/app/(messages)/messages/archived` — Archived conversations list
- **Web:** `/app/settings/messages/retention` — Retention policy settings
- **Mobile:** `/app/(messages)/messages/archived` — Archived messages

**System Touchpoints:**
- **communications-be/conversation** — POST /v1/conversations/{conversation_id}/archive (archive conversation)
- **communications-be/conversation** — POST /v1/conversations/{conversation_id}/unarchive (restore conversation)
- **communications-be/conversation** — GET /v1/conversations/archived (list archived)
- **communications-be/retention** — POST /v1/conversations/{conversation_id}/retention-policy (set retention)
- **communications-be/export** — POST /v1/conversations/{conversation_id}/export (export before deletion)

**Flow Steps:**

1. **Archive conversation:**
   - User opens conversation actions menu (three dots)
   - Select **"Archive conversation"**
   - **Confirmation modal:** "Archive this conversation? You can restore it anytime."
   - User confirms
   - **POST /v1/conversations/{conversation_id}/archive:**
     ```json
     {
       "conversation_id": "uuid",
       "archived_by": "user_uuid",
       "archived_at": "2025-11-06T10:00:00Z",
       "reason": "PROJECT_COMPLETED" // optional
     }
     ```
   - Conversation removed from main inbox
   - Success toast: "Conversation archived"

2. **View archived conversations:**
   - Navigate to `/app/(messages)/messages/archived`
   - **GET /v1/conversations/archived:**
     ```json
     {
       "conversations": [
         {
           "conversation_id": "uuid",
           "name": "Project Alpha Discussion",
           "archived_at": "2025-11-01T15:30:00Z",
           "message_count": 145,
           "last_message_at": "2025-10-28T09:20:00Z",
           "participants": [...]
         }
       ],
       "total_count": 23
     }
     ```
   - **List shows:**
     - Conversation name
     - Archived date
     - Last message date
     - Participant count
     - "Restore" button

3. **Restore archived conversation:**
   - Click **"Restore"** on archived conversation
   - **POST /v1/conversations/{conversation_id}/unarchive**
   - Conversation moves back to main inbox
   - All messages and history intact
   - Success toast: "Conversation restored to inbox"

4. **Auto-archive rules:**
   - Navigate to `/app/settings/messages/retention`
   - **Create auto-archive rule:**
     - **Trigger:** "After X days of inactivity"
     - **Apply to:** All conversations, or specific types (e.g., "Job proposals")
     - **Exclude:** Starred, Important
   - Example: "Archive conversations after 90 days of inactivity"
   - **POST /v1/users/me/settings/auto-archive:**
     ```json
     {
       "enabled": true,
       "inactivity_days": 90,
       "exclude_starred": true,
       "exclude_important": true,
       "conversation_types": ["JOB_PROPOSAL", "GENERAL"]
     }
     ```
   - System runs nightly job to apply rules

5. **Retention policy (auto-delete):**
   - **Admin or user** can set retention policy
   - **Navigate to:** `/app/settings/messages/retention`
   - **Set policy:** "Delete conversations after X days"
   - **Warning:** "Deleted conversations cannot be recovered"
   - **POST /v1/conversations/{conversation_id}/retention-policy:**
     ```json
     {
       "retention_days": 365,
       "apply_to_archived_only": true,
       "export_before_deletion": true
     }
     ```
   - System schedules deletion after retention period
   - User receives **email warning** 7 days before deletion

6. **Export before deletion:**
   - Before retention policy deletes conversation
   - System auto-exports to user's storage
   - **POST /v1/conversations/{conversation_id}/export:**
     ```json
     {
       "format": "JSON",
       "include_attachments": true,
       "export_reason": "RETENTION_POLICY"
     }
     ```
   - Export stored in `/app/settings/data-export`
   - User notified: "Conversation exported before deletion"

7. **Bulk archive:**
   - User selects multiple conversations (checkbox)
   - Click **"Archive selected"** (bulk action)
   - Confirmation: "Archive 5 conversations?"
   - **POST /v1/conversations/bulk-archive:**
     ```json
     {
       "conversation_ids": ["uuid1", "uuid2", "uuid3", "uuid4", "uuid5"]
     }
     ```
   - All conversations archived simultaneously
   - Success: "5 conversations archived"

8. **Search within archived:**
   - In archived messages view, search bar available
   - **GET /v1/conversations/archived/search?q=budget**
   - Search limited to archived conversations
   - Results show archived date and restore option

**Branches & Edge Cases:**

- **Archive limit:** No limit on archived conversations (unlimited storage)
- **Archive notification:** Archived conversations don't send notifications
- **Unread count:** Archived conversations excluded from unread count
- **Archive syncs:** Archiving syncs across all devices
- **Archive undo:** Can immediately undo archive within 10 seconds (toast with "Undo")
- **Archive filter:** Filter archived by date range, participants, type
- **Archive search:** Full-text search within archived conversations
- **Archive export all:** Export all archived conversations at once
- **Retention override:** Admin can override user retention policies (compliance)
- **Retention warning:** Users warned 30, 7, and 1 day before deletion
- **Retention exemption:** "Legal hold" flag prevents deletion (compliance)
- **Deleted conversation recovery:** Soft delete for 30 days (admin recovery only)
- **Archive performance:** Archiving instant (<100ms), bulk operations <5s for 100 conversations
- **Archive statistics:** Show total archived count, storage used
- **Archive mobile optimization:** Swipe to archive on mobile
- **Auto-archive notification:** Optional weekly summary of auto-archived conversations
- **Archive conflict:** If conversation archived on device A, syncs to device B instantly
- **Restore multiple:** Bulk restore multiple archived conversations
- **Archive empty state:** "No archived conversations" with helpful tips
- **Retention compliance:** Audit log for all retention actions (GDPR, HIPAA)

**Notifications:**

- **Email (retention warning):** "Conversation will be deleted in 7 days" (sent 30d, 7d, 1d before)
- **In-app (archive):** Toast notification "Conversation archived" (with Undo)
- **Email (auto-archive):** Weekly digest of auto-archived conversations (optional)

**Analytics:**

- messages.conversation\_archived (conversation\_id, archive\_reason)
- messages.conversation\_unarchived (conversation\_id, days\_archived)
- messages.archived\_conversations\_viewed (archived\_count)
- messages.auto\_archive\_rule\_created (inactivity\_days, exclude\_flags)
- messages.auto\_archive\_executed (conversation\_count)
- messages.retention\_policy\_set (retention\_days, conversation\_id)
- messages.retention\_warning\_sent (conversation\_id, days\_until\_deletion)
- messages.conversation\_exported\_before\_deletion (conversation\_id, export\_format)
- messages.conversation\_deleted\_by\_retention (conversation\_id, retention\_days)
- messages.bulk\_archive\_executed (conversation\_count)
- messages.archived\_conversation\_searched (query, result\_count)
- messages.archive\_undo\_clicked (conversation\_id)
- messages.bulk\_restore\_executed (conversation\_count)

**Sources:** communications-be.user-stories.md (Section 9: Archiving & Export), combined-fe-folder-structure.md ((messages) routes, archived folder), general best practices for message retention.

---