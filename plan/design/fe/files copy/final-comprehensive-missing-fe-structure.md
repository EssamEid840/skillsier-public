# Final Comprehensive Missing Frontend Folder Structure
## Completing All Requirements from fe-folder-structure-prompt.md

> **Note**: This document contains ONLY the folder structure elements that are:
> 1. Required by `fe-folder-structure-prompt.md`
> 2. NOT present in `combined-folder-structure.md`
> 3. NOT present in any of the missing-fe-folder-structure documents

---

## I. Critical Missing Dashboard Routes

### 1. Learning & Professional Development (Freelancer)

```
apps/web/src/app/[locale]/(dashboard)/
│
├── learning/
│   ├── courses/
│   │   ├── [courseId]/
│   │   │   ├── lessons/
│   │   │   │   └── [lessonId]/
│   │   │   │       └── page.tsx  # Lesson view
│   │   │   │           # - Video/content player
│   │   │   │           # - Notes taking
│   │   │   │           # - Progress tracking
│   │   │   │           # BE: utility-be/learning OR external LMS
│   │   │   │           # GET /v1/courses/{course_id}/lessons/{lesson_id}
│   │   │   ├── assessments/
│   │   │   │   └── [assessmentId]/
│   │   │   │       ├── start/
│   │   │   │       │   └── page.tsx  # Start assessment
│   │   │   │       │       # BE: utility-be/learning
│   │   │   │       │       # POST /v1/assessments/{assessment_id}/start
│   │   │   │       ├── submit/
│   │   │   │       │   └── page.tsx  # Submit answers
│   │   │   │       │       # BE: utility-be/learning
│   │   │   │       │       # POST /v1/assessments/{assessment_id}/submit
│   │   │   │       └── results/
│   │   │   │           └── page.tsx  # View results
│   │   │   │               # BE: utility-be/learning
│   │   │   │               # GET /v1/assessments/{assessment_id}/results
│   │   │   └── page.tsx  # Course overview
│   │   │       # - Syllabus
│   │   │       # - Progress
│   │   │       # - Enroll button
│   │   │       # BE: utility-be/learning
│   │   │       # GET /v1/courses/{course_id}
│   │   │       # POST /v1/courses/{course_id}/enroll
│   │   ├── browse/
│   │   │   └── page.tsx  # Browse all courses
│   │   │       # - Filter by skill
│   │   │       # - Search
│   │   │       # - Recommendations
│   │   │       # BE: utility-be/learning
│   │   │       # GET /v1/courses
│   │   └── my-courses/
│   │       └── page.tsx  # Enrolled courses
│   │           # - In progress
│   │           # - Completed
│   │           # - Certificates
│   │           # BE: utility-be/learning
│   │           # GET /v1/users/me/courses
│   │
│   ├── certifications/
│   │   ├── [certId]/
│   │   │   ├── verify/
│   │   │   │   └── page.tsx  # Verify certificate
│   │   │   │       # BE: utility-be/learning
│   │   │   │       # GET /v1/certifications/{cert_id}/verify
│   │   │   └── page.tsx  # Certificate detail
│   │   │       # - Download PDF
│   │   │       # - Share link
│   │   │       # - Add to profile
│   │   │       # BE: utility-be/learning
│   │   │       # GET /v1/certifications/{cert_id}
│   │   └── page.tsx  # All certifications
│   │       # - Earned certificates
│   │       # - Available certifications
│   │       # BE: utility-be/learning
│   │       # GET /v1/users/me/certifications
│   │
│   └── skill-tests/
│       ├── [testId]/
│       │   ├── instructions/
│       │   │   └── page.tsx  # Test instructions
│       │   │       # BE: users-be/capabilities
│       │   │       # GET /v1/skill-tests/{test_id}
│       │   ├── take-test/
│       │   │   └── page.tsx  # Take the test
│       │   │       # - Timed test interface
│       │   │       # - Submit answers
│       │   │       # BE: users-be/capabilities
│       │   │       # POST /v1/skill-tests/{test_id}/submit
│       │   └── results/
│       │       └── page.tsx  # Test results
│       │           # - Score
│       │           # - Percentile
│       │           # - Badge earned
│       │           # BE: users-be/capabilities
│       │           # GET /v1/skill-tests/{test_id}/results
│       └── page.tsx  # Available skill tests
│           # - Browse by skill
│           # - Test history
│           # - Top scores
│           # BE: users-be/capabilities
│           # GET /v1/skill-tests
```

### 2. Community & Networking

```
apps/web/src/app/[locale]/(dashboard)/
│
├── community/
│   ├── forums/
│   │   ├── [forumId]/
│   │   │   ├── threads/
│   │   │   │   └── [threadId]/
│   │   │   │       ├── page.tsx  # Thread view
│   │   │   │       │   # - Posts
│   │   │   │       │   # - Reply
│   │   │   │       │   # - Voting
│   │   │   │       │   # BE: communications-be/forums (if exists) OR community service
│   │   │   │       │   # GET /v1/forums/{forum_id}/threads/{thread_id}
│   │   │   │       └── reply/
│   │   │   │           └── page.tsx  # Reply to thread
│   │   │   │               # BE: communications-be/forums
│   │   │   │               # POST /v1/threads/{thread_id}/replies
│   │   │   └── page.tsx  # Forum overview
│   │   │       # - Thread list
│   │   │       # - Create thread
│   │   │       # BE: communications-be/forums
│   │   │       # GET /v1/forums/{forum_id}/threads
│   │   └── page.tsx  # All forums
│   │       # - Categories
│   │       # - Popular threads
│   │       # BE: communications-be/forums
│   │       # GET /v1/forums
│   │
│   ├── groups/
│   │   ├── [groupId]/
│   │   │   ├── members/
│   │   │   │   └── page.tsx  # Group members
│   │   │   │       # - Member list
│   │   │   │       # - Invite
│   │   │   │       # - Roles
│   │   │   │       # BE: users-be/groups (if exists) OR community service
│   │   │   │       # GET /v1/groups/{group_id}/members
│   │   │   ├── events/
│   │   │   │   ├── [eventId]/
│   │   │   │   │   └── page.tsx  # Event detail
│   │   │   │   │       # - RSVP
│   │   │   │   │       # - Calendar add
│   │   │   │   │       # BE: communications-be/events OR community service
│   │   │   │   │       # GET /v1/events/{event_id}
│   │   │   │   └── page.tsx  # Group events
│   │   │   │       # - Upcoming
│   │   │   │       # - Past
│   │   │   │       # BE: communications-be/events
│   │   │   │       # GET /v1/groups/{group_id}/events
│   │   │   ├── discussions/
│   │   │   │   └── page.tsx  # Group discussions
│   │   │   │       # - Posts feed
│   │   │   │       # - Comment
│   │   │   │       # BE: communications-be/discussions
│   │   │   │       # GET /v1/groups/{group_id}/discussions
│   │   │   └── page.tsx  # Group overview
│   │   │       # - About
│   │   │       # - Join/leave
│   │   │       # - Activity feed
│   │   │       # BE: users-be/groups
│   │   │       # GET /v1/groups/{group_id}
│   │   ├── discover/
│   │   │   └── page.tsx  # Discover groups
│   │   │       # - Recommended
│   │   │       # - By interest
│   │   │       # BE: users-be/groups
│   │   │       # GET /v1/groups/discover
│   │   └── my-groups/
│   │       └── page.tsx  # My groups
│   │           # - Joined groups
│   │           # - Manage groups
│   │           # BE: users-be/groups
│   │           # GET /v1/users/me/groups
│   │
│   └── events/
│       ├── [eventId]/
│       │   ├── register/
│       │   │   └── page.tsx  # Event registration
│       │   │       # - RSVP form
│       │   │       # - Add to calendar
│       │   │       # BE: communications-be/events
│       │   │       # POST /v1/events/{event_id}/register
│       │   └── page.tsx  # Event detail
│       │       # - Info
│       │       # - Attendees
│       │       # - Check-in (QR code)
│       │       # BE: communications-be/events
│       │       # GET /v1/events/{event_id}
│       ├── upcoming/
│       │   └── page.tsx  # Upcoming events
│       │       # BE: communications-be/events
│       │       # GET /v1/events?status=upcoming
│       └── my-events/
│           └── page.tsx  # My registered events
│               # - Attending
│               # - Past attended
│               # BE: communications-be/events
│               # GET /v1/users/me/events
```

### 3. Enterprise/Agency Routes (Client-specific)

```
apps/web/src/app/[locale]/(dashboard)/
│
├── agency/
│   ├── overview/
│   │   └── page.tsx  # Agency dashboard
│   │       # - Revenue overview
│   │       # - Active projects
│   │       # - Team utilization
│   │       # - Client list
│   │       # BE: users-be/org, financial-be/analytics
│   │       # GET /v1/agencies/me/overview
│   │
│   ├── sub-accounts/
│   │   ├── [subAccountId]/
│   │   │   ├── settings/
│   │   │   │   └── page.tsx  # Sub-account settings
│   │   │   │       # BE: users-be/org
│   │   │   │       # PUT /v1/agencies/sub-accounts/{sub_account_id}
│   │   │   └── page.tsx  # Sub-account detail
│   │   │       # - Jobs posted
│   │   │       # - Contracts
│   │   │       # - Spending
│   │   │       # BE: users-be/org
│   │   │       # GET /v1/agencies/sub-accounts/{sub_account_id}
│   │   ├── create/
│   │   │   └── page.tsx  # Create sub-account
│   │   │       # BE: users-be/org
│   │   │       # POST /v1/agencies/sub-accounts
│   │   └── page.tsx  # All sub-accounts
│   │       # - List all
│   │       # - Manage
│   │       # BE: users-be/org
│   │       # GET /v1/agencies/sub-accounts
│   │
│   ├── talent-pool/
│   │   ├── [poolId]/
│   │   │   ├── members/
│   │   │   │   └── page.tsx  # Pool members
│   │   │   │       # - Freelancer list
│   │   │   │       # - Add/remove
│   │   │   │       # BE: users-be/org, search-be/talent-pool
│   │   │   │       # GET /v1/talent-pools/{pool_id}/members
│   │   │   └── page.tsx  # Talent pool detail
│   │   │       # BE: search-be/talent-pool
│   │   │       # GET /v1/talent-pools/{pool_id}
│   │   ├── create/
│   │   │   └── page.tsx  # Create talent pool
│   │   │       # BE: search-be/talent-pool
│   │   │       # POST /v1/talent-pools
│   │   └── page.tsx  # All talent pools
│   │       # BE: search-be/talent-pool
│   │       # GET /v1/talent-pools
│   │
│   ├── white-label/
│   │   └── page.tsx  # White-label settings
│   │       # - Branding
│   │       # - Custom domain
│   │       # - Logo/colors
│   │       # BE: users-be/org
│   │       # PUT /v1/agencies/white-label
│   │
│   └── reporting/
│       ├── clients/
│       │   └── page.tsx  # Client reports
│       │       # - Per-client spending
│       │       # - Project status
│       │       # BE: financial-be/analytics
│       │       # GET /v1/agencies/reports/clients
│       ├── team/
│       │   └── page.tsx  # Team performance
│       │       # - Utilization rates
│       │       # - Project assignments
│       │       # BE: users-be/org, jobs-be/analytics
│       │       # GET /v1/agencies/reports/team
│       └── financial/
│           └── page.tsx  # Financial reports
│               # - Revenue by period
│               # - Margins
│               # - Forecasts
│               # BE: financial-be/analytics
│               # GET /v1/agencies/reports/financial
```

### 4. Freelancer Specializations

```
apps/web/src/app/[locale]/(dashboard)/
│
├── specialized/
│   ├── top-rated/
│   │   ├── application/
│   │   │   └── page.tsx  # Apply for Top Rated
│   │   │       # - Eligibility check
│   │   │       # - Submit application
│   │   │       # BE: users-be/badges OR reviews-be/reputation
│   │   │       # POST /v1/users/me/top-rated/apply
│   │   └── page.tsx  # Top Rated program info
│   │       # - Benefits
│   │       # - Requirements
│   │       # - Application status
│   │       # BE: users-be/badges
│   │       # GET /v1/programs/top-rated
│   │
│   ├── plus/
│   │   ├── subscribe/
│   │   │   └── page.tsx  # Subscribe to Plus
│   │   │       # - Plans
│   │   │       # - Payment
│   │   │       # BE: subscriptions-be/subscription (financial-be)
│   │   │       # POST /v1/subscriptions/plus
│   │   └── page.tsx  # Plus membership overview
│   │       # - Benefits
│   │       # - Pricing
│   │       # - Comparison
│   │       # BE: subscriptions-be/plans
│   │       # GET /v1/subscriptions/plans/plus
│   │
│   └── talent-cloud/
│       ├── projects/
│       │   └── page.tsx  # Talent Cloud exclusive projects
│       │       # - High-value projects
│       │       # - Direct invites
│       │       # BE: jobs-be/job (with exclusive flag)
│       │       # GET /v1/jobs/talent-cloud
│       └── page.tsx  # Talent Cloud info
│           # - Qualification
│           # - Apply
│           # - Status
│           # BE: users-be/badges
│           # GET /v1/programs/talent-cloud
```

---

## II. Missing Public/Marketing Routes

### 1. Solution Pages (Marketing)

```
apps/web/src/app/[locale]/(public)/
│
├── solutions/
│   ├── by-role/
│   │   ├── developers/
│   │   │   └── page.tsx  # Developer hiring solutions
│   │   ├── designers/
│   │   │   └── page.tsx  # Designer hiring solutions
│   │   ├── writers/
│   │   │   └── page.tsx  # Writer hiring solutions
│   │   └── page.tsx  # Browse by role
│   ├── by-industry/
│   │   ├── tech/
│   │   │   └── page.tsx  # Tech industry solutions
│   │   ├── marketing/
│   │   │   └── page.tsx  # Marketing solutions
│   │   ├── finance/
│   │   │   └── page.tsx  # Finance solutions
│   │   └── page.tsx  # Browse by industry
│   ├── enterprise/
│   │   ├── contact-sales/
│   │   │   └── page.tsx  # Enterprise contact form
│   │   │       # BE: communications-be/contact
│   │   │       # POST /v1/contact/enterprise
│   │   └── page.tsx  # Enterprise solutions
│   │       # - MSA contracts
│   │       # - Dedicated support
│   │       # - Volume pricing
│   └── agencies/
│       └── page.tsx  # Agency solutions
│           # - White-label
│           # - Multi-client
│           # - Revenue share
```

### 2. Legal & Compliance Pages

```
apps/web/src/app/[locale]/(public)/
│
├── legal/
│   ├── terms/
│   │   ├── service/
│   │   │   └── page.tsx  # Terms of Service
│   │   ├── freelancer/
│   │   │   └── page.tsx  # Freelancer Agreement
│   │   ├── client/
│   │   │   └── page.tsx  # Client Agreement
│   │   └── page.tsx  # Legal terms overview
│   ├── privacy/
│   │   ├── policy/
│   │   │   └── page.tsx  # Privacy Policy
│   │   ├── cookies/
│   │   │   └── page.tsx  # Cookie Policy
│   │   ├── gdpr/
│   │   │   └── page.tsx  # GDPR information
│   │   └── ccpa/
│   │       └── page.tsx  # CCPA information
│   ├── compliance/
│   │   ├── kyc/
│   │   │   └── page.tsx  # KYC policy
│   │   ├── aml/
│   │   │   └── page.tsx  # AML policy
│   │   └── page.tsx  # Compliance overview
│   ├── dmca/
│   │   └── page.tsx  # DMCA policy
│   ├── accessibility/
│   │   └── page.tsx  # Accessibility statement
│   └── licenses/
│       └── page.tsx  # Open source licenses
```

### 3. Resource Center

```
apps/web/src/app/[locale]/(public)/
│
├── resources/
│   ├── guides/
│   │   ├── [slug]/
│   │   │   └── page.tsx  # Guide detail
│   │   └── page.tsx  # All guides
│   │       # - Getting started
│   │       # - Best practices
│   │       # - How-tos
│   ├── case-studies/
│   │   ├── [slug]/
│   │   │   └── page.tsx  # Case study detail
│   │   └── page.tsx  # All case studies
│   ├── whitepapers/
│   │   ├── [slug]/
│   │   │   └── page.tsx  # Whitepaper detail
│   │   │       # - Download gated content
│   │   │       # BE: communications-be/leads
│   │   │       # POST /v1/leads/download
│   │   └── page.tsx  # All whitepapers
│   ├── webinars/
│   │   ├── [slug]/
│   │   │   ├── register/
│   │   │   │   └── page.tsx  # Webinar registration
│   │   │   │       # BE: communications-be/events
│   │   │   │       # POST /v1/webinars/{webinar_id}/register
│   │   │   └── page.tsx  # Webinar detail
│   │   ├── upcoming/
│   │   │   └── page.tsx  # Upcoming webinars
│   │   └── on-demand/
│   │       └── page.tsx  # On-demand recordings
│   ├── tools/
│   │   ├── rate-calculator/
│   │   │   └── page.tsx  # Hourly rate calculator
│   │   ├── roi-calculator/
│   │   │   └── page.tsx  # ROI calculator
│   │   └── budget-estimator/
│   │       └── page.tsx  # Project budget estimator
│   └── api/
│       ├── docs/
│       │   └── page.tsx  # API documentation
│       ├── reference/
│       │   └── page.tsx  # API reference
│       └── changelog/
│           └── page.tsx  # API changelog
```

### 4. Developer Portal (Public)

```
apps/web/src/app/[locale]/(public)/
│
├── developers/
│   ├── api/
│   │   ├── getting-started/
│   │   │   └── page.tsx  # API getting started
│   │   ├── authentication/
│   │   │   └── page.tsx  # OAuth & API keys
│   │   ├── endpoints/
│   │   │   ├── [category]/
│   │   │   │   └── page.tsx  # Endpoint category
│   │   │   └── page.tsx  # All endpoints
│   │   ├── webhooks/
│   │   │   └── page.tsx  # Webhook documentation
│   │   ├── rate-limits/
│   │   │   └── page.tsx  # Rate limiting info
│   │   └── page.tsx  # API overview
│   ├── sdks/
│   │   ├── javascript/
│   │   │   └── page.tsx  # JavaScript SDK
│   │   ├── python/
│   │   │   └── page.tsx  # Python SDK
│   │   ├── ruby/
│   │   │   └── page.tsx  # Ruby SDK
│   │   └── page.tsx  # All SDKs
│   ├── plugins/
│   │   ├── wordpress/
│   │   │   └── page.tsx  # WordPress plugin
│   │   ├── shopify/
│   │   │   └── page.tsx  # Shopify app
│   │   └── page.tsx  # All plugins/integrations
│   ├── sample-code/
│   │   └── page.tsx  # Code examples
│   └── community/
│       ├── forum/
│       │   └── page.tsx  # Developer forum
│       └── github/
│           └── page.tsx  # GitHub repos
```

---

## III. Missing Mobile App Routes

### 1. Mobile-Specific Quick Actions

```
apps/mobile/app/
│
├── quick-actions/
│   ├── quick-apply/
│   │   └── [jobId].tsx  # Quick apply to job
│   │       # - Pre-filled proposal
│   │       # - One-tap apply
│   │       # - Connect deduction
│   │       # BE: proposals-be/proposal
│   │       # POST /v1/proposals/quick-apply
│   │
│   ├── quick-message/
│   │   └── [userId].tsx  # Quick message to user
│   │       # - Template messages
│   │       # - Voice-to-text
│   │       # BE: communications-be/conversation
│   │       # POST /v1/conversations/{conversation_id}/messages/quick
│   │
│   └── quick-time-entry/
│       └── [contractId].tsx  # Quick time logging
│           # - Timer widget
│           # - Quick notes
│           # - One-tap submit
│           # BE: contracts-be/work_diary
│           # POST /v1/work-diary/quick-entry
```

### 2. Mobile Onboarding

```
apps/mobile/app/
│
├── onboarding/
│   ├── intro/
│   │   └── index.tsx  # App intro screens
│   │       # - Swipeable intro
│   │       # - Feature highlights
│   │
│   ├── permissions/
│   │   ├── camera.tsx  # Camera permission
│   │   ├── location.tsx  # Location permission
│   │   ├── notifications.tsx  # Notification permission
│   │   └── index.tsx  # All permissions
│   │
│   ├── biometric/
│   │   └── setup.tsx  # Biometric auth setup
│   │       # - Face ID
│   │       # - Touch ID
│   │       # BE: users-be/auth
│   │       # POST /v1/auth/biometric/enable
│   │
│   └── tutorial/
│       ├── browse-jobs.tsx  # Tutorial: Browse jobs
│       ├── apply.tsx  # Tutorial: Apply
│       ├── time-tracking.tsx  # Tutorial: Time tracking
│       └── index.tsx  # Interactive tutorial
```

### 3. Mobile Settings

```
apps/mobile/app/
│
├── mobile-settings/
│   ├── app-settings/
│   │   ├── theme.tsx  # Theme selection
│   │   ├── language.tsx  # Language selection
│   │   ├── notifications.tsx  # Notification preferences
│   │   ├── data-usage.tsx  # Data usage settings
│   │   ├── cache.tsx  # Clear cache
│   │   └── index.tsx  # App settings overview
│   │
│   ├── security/
│   │   ├── biometric.tsx  # Biometric settings
│   │   ├── auto-lock.tsx  # Auto-lock timer
│   │   ├── pin.tsx  # PIN setup
│   │   └── index.tsx  # Security settings
│   │
│   ├── offline/
│   │   ├── sync-frequency.tsx  # Sync settings
│   │   ├── offline-mode.tsx  # Offline capabilities
│   │   └── storage.tsx  # Storage management
│   │
│   └── advanced/
│       ├── developer-mode.tsx  # Developer options
│       ├── logs.tsx  # Error logs
│       └── diagnostics.tsx  # Diagnostics tools
```

---

## IV. Missing Shared Features (packages/shared/src/)

### 1. Webhooks Management

```
packages/shared/src/
│
├── features/
│   ├── webhooks/
│   │   ├── api/
│   │   │   └── webhooks-api.ts  # Webhooks API client
│   │   │       # BE: utility-be/webhooks OR admin-be
│   │   │       # GET /v1/webhooks
│   │   │       # POST /v1/webhooks
│   │   │       # PUT /v1/webhooks/{webhook_id}
│   │   │       # DELETE /v1/webhooks/{webhook_id}
│   │   │       # POST /v1/webhooks/{webhook_id}/test
│   │   ├── hooks/
│   │   │   ├── useWebhooks.ts  # List webhooks
│   │   │   ├── useWebhook.ts  # Single webhook
│   │   │   ├── useWebhookLogs.ts  # Webhook delivery logs
│   │   │   └── useWebhookTest.ts  # Test webhook
│   │   ├── queries/
│   │   │   ├── webhooks-mutations.ts  # Webhook mutations
│   │   │   └── webhooks-queries.ts  # Webhook queries
│   │   ├── components/
│   │   │   ├── WebhookForm/
│   │   │   │   ├── WebhookForm.tsx
│   │   │   │   ├── WebhookForm.web.tsx
│   │   │   │   └── WebhookForm.types.ts
│   │   │   ├── WebhookLogs/
│   │   │   │   ├── WebhookLogs.tsx
│   │   │   │   └── WebhookLogs.types.ts
│   │   │   └── EventSelector/
│   │   │       ├── EventSelector.tsx
│   │   │       └── EventSelector.types.ts
│   │   └── types.ts  # Webhook types
```

### 2. Feature Experimentation (A/B Testing)

```
packages/shared/src/
│
├── features/
│   ├── experiments/
│   │   ├── api/
│   │   │   └── experiments-api.ts  # Experiments API
│   │   │       # BE: utility-be/experiments
│   │   │       # GET /v1/experiments/active
│   │   │       # POST /v1/experiments/{experiment_id}/track
│   │   ├── hooks/
│   │   │   ├── useExperiment.ts  # Get experiment variant
│   │   │   ├── useExperimentTracking.ts  # Track experiment events
│   │   │   └── useFeatureVariant.ts  # Feature variant
│   │   ├── providers/
│   │   │   └── ExperimentProvider.tsx  # Experiment context
│   │   ├── utils/
│   │   │   ├── variant-assignment.ts  # Assign variant
│   │   │   ├── experiment-storage.ts  # Store assignments
│   │   │   └── tracking.ts  # Track events
│   │   └── types.ts  # Experiment types
```

### 3. Real-time Collaboration

```
packages/shared/src/
│
├── features/
│   ├── collaboration/
│   │   ├── api/
│   │   │   └── collaboration-api.ts  # Collaboration API
│   │   │       # BE: communications-be/realtime
│   │   ├── hooks/
│   │   │   ├── useCollaboration.ts  # Collaboration session
│   │   │   ├── usePresence.ts  # User presence
│   │   │   ├── useCursors.ts  # Cursor tracking
│   │   │   └── useSharedState.ts  # Shared state sync
│   │   ├── providers/
│   │   │   └── CollaborationProvider.tsx  # Collab context
│   │   ├── components/
│   │   │   ├── PresenceIndicator/
│   │   │   │   ├── PresenceIndicator.tsx
│   │   │   │   └── PresenceIndicator.types.ts
│   │   │   ├── CollaboratorCursor/
│   │   │   │   ├── CollaboratorCursor.tsx
│   │   │   │   └── CollaboratorCursor.types.ts
│   │   │   └── ActiveUsers/
│   │   │       ├── ActiveUsers.tsx
│   │   │       └── ActiveUsers.types.ts
│   │   └── types.ts  # Collaboration types
```

### 4. Voice & Video Integration

```
packages/shared/src/
│
├── features/
│   ├── video/
│   │   ├── api/
│   │   │   └── video-api.ts  # Video call API
│   │   │       # BE: communications-be/video
│   │   │       # POST /v1/video/rooms
│   │   │       # GET /v1/video/rooms/{room_id}/token
│   │   ├── hooks/
│   │   │   ├── useVideoCall.ts  # Video call management
│   │   │   ├── useScreenShare.ts  # Screen sharing
│   │   │   ├── useRecording.ts  # Call recording
│   │   │   └── useVideoDevices.ts  # Device selection
│   │   ├── components/
│   │   │   ├── VideoRoom/
│   │   │   │   ├── VideoRoom.tsx
│   │   │   │   ├── VideoRoom.web.tsx
│   │   │   │   ├── VideoRoom.native.tsx
│   │   │   │   └── VideoRoom.types.ts
│   │   │   ├── VideoControls/
│   │   │   │   ├── VideoControls.tsx
│   │   │   │   └── VideoControls.types.ts
│   │   │   └── ParticipantGrid/
│   │   │       ├── ParticipantGrid.tsx
│   │   │       └── ParticipantGrid.types.ts
│   │   └── types.ts  # Video types
```

### 5. Localization Tools

```
packages/shared/src/
│
├── features/
│   ├── i18n/
│   │   ├── tools/
│   │   │   ├── translation-manager.ts  # Manage translations
│   │   │   ├── missing-keys-detector.ts  # Detect missing keys
│   │   │   ├── pluralization-helper.ts  # Pluralization
│   │   │   └── currency-formatter.ts  # Currency formatting
│   │   ├── hooks/
│   │   │   ├── useTranslation.ts  # (already exists, enhanced)
│   │   │   ├── useDateFormat.ts  # Date formatting
│   │   │   ├── useNumberFormat.ts  # Number formatting
│   │   │   ├── useCurrencyFormat.ts  # Currency formatting
│   │   │   └── useRTL.ts  # RTL detection
│   │   ├── components/
│   │   │   ├── TranslationProvider/
│   │   │   │   └── TranslationProvider.tsx
│   │   │   ├── LocaleSwitcher/
│   │   │   │   ├── LocaleSwitcher.tsx
│   │   │   │   ├── LocaleSwitcher.web.tsx
│   │   │   │   └── LocaleSwitcher.native.tsx
│   │   │   └── FormattedMessage/
│   │   │       └── FormattedMessage.tsx
│   │   └── utils/
│   │       ├── locale-detection.ts  # Auto-detect locale
│   │       ├── fallback-loader.ts  # Fallback translations
│   │       └── interpolation.ts  # Message interpolation
```

---

## V. Missing UI Components (packages/ui/src/)

### 1. Advanced Input Components

```
packages/ui/src/
│
├── components/
│   ├── Form/
│   │   ├── RichTextEditor/
│   │   │   ├── RichTextEditor.tsx
│   │   │   ├── RichTextEditor.web.tsx  # Quill/TipTap
│   │   │   ├── RichTextEditor.native.tsx  # Limited rich text
│   │   │   ├── Toolbar/
│   │   │   │   ├── Toolbar.tsx
│   │   │   │   └── Toolbar.types.ts
│   │   │   └── RichTextEditor.types.ts
│   │   ├── CodeEditor/
│   │   │   ├── CodeEditor.tsx
│   │   │   ├── CodeEditor.web.tsx  # Monaco/CodeMirror
│   │   │   ├── LanguageSelector/
│   │   │   │   └── LanguageSelector.tsx
│   │   │   └── CodeEditor.types.ts
│   │   ├── MarkdownEditor/
│   │   │   ├── MarkdownEditor.tsx
│   │   │   ├── MarkdownEditor.web.tsx
│   │   │   ├── Preview/
│   │   │   │   └── MarkdownPreview.tsx
│   │   │   └── MarkdownEditor.types.ts
│   │   ├── SignaturePad/
│   │   │   ├── SignaturePad.tsx
│   │   │   ├── SignaturePad.web.tsx  # Canvas-based
│   │   │   ├── SignaturePad.native.tsx  # React Native Canvas
│   │   │   └── SignaturePad.types.ts
│   │   └── DateRangePicker/
│   │       ├── DateRangePicker.tsx
│   │       ├── DateRangePicker.web.tsx
│   │       ├── DateRangePicker.native.tsx
│   │       └── DateRangePicker.types.ts
```

### 2. Data Visualization Components

```
packages/ui/src/
│
├── components/
│   ├── Charts/
│   │   ├── HeatMap/
│   │   │   ├── HeatMap.tsx
│   │   │   ├── HeatMap.web.tsx  # D3/Recharts
│   │   │   └── HeatMap.types.ts
│   │   ├── GanttChart/
│   │   │   ├── GanttChart.tsx
│   │   │   ├── GanttChart.web.tsx
│   │   │   ├── Task/
│   │   │   │   └── Task.tsx
│   │   │   └── GanttChart.types.ts
│   │   ├── FunnelChart/
│   │   │   ├── FunnelChart.tsx
│   │   │   └── FunnelChart.types.ts
│   │   └── OrgChart/
│   │       ├── OrgChart.tsx
│   │       ├── OrgChart.web.tsx
│   │       ├── Node/
│   │       │   └── OrgNode.tsx
│   │       └── OrgChart.types.ts
```

### 3. AI/Smart Components

```
packages/ui/src/
│
├── components/
│   ├── AI/
│   │   ├── AIAssistant/
│   │   │   ├── AIAssistant.tsx
│   │   │   ├── AIAssistant.web.tsx
│   │   │   ├── AIAssistant.native.tsx
│   │   │   ├── ChatInterface/
│   │   │   │   └── ChatInterface.tsx
│   │   │   └── AIAssistant.types.ts
│   │   ├── SmartSuggestions/
│   │   │   ├── SmartSuggestions.tsx
│   │   │   ├── SuggestionCard/
│   │   │   │   └── SuggestionCard.tsx
│   │   │   └── SmartSuggestions.types.ts
│   │   ├── AutoComplete/
│   │   │   ├── SmartAutoComplete.tsx
│   │   │   ├── SmartAutoComplete.web.tsx
│   │   │   ├── SmartAutoComplete.native.tsx
│   │   │   └── SmartAutoComplete.types.ts
│   │   └── ContentGenerator/
│   │       ├── ContentGenerator.tsx
│   │       ├── TemplateSelector/
│   │       │   └── TemplateSelector.tsx
│   │       └── ContentGenerator.types.ts
```

### 4. Accessibility Components

```
packages/ui/src/
│
├── components/
│   ├── Accessibility/
│   │   ├── SkipLinks/
│   │   │   ├── SkipLinks.tsx
│   │   │   ├── SkipLinks.web.tsx
│   │   │   └── SkipLinks.types.ts
│   │   ├── ScreenReaderOnly/
│   │   │   ├── ScreenReaderOnly.tsx
│   │   │   └── ScreenReaderOnly.types.ts
│   │   ├── FocusTrap/
│   │   │   ├── FocusTrap.tsx
│   │   │   ├── FocusTrap.web.tsx
│   │   │   └── FocusTrap.types.ts
│   │   ├── LiveRegion/
│   │   │   ├── LiveRegion.tsx
│   │   │   ├── LiveRegion.web.tsx
│   │   │   ├── LiveRegion.native.tsx
│   │   │   └── LiveRegion.types.ts
│   │   └── KeyboardShortcuts/
│   │       ├── KeyboardShortcuts.tsx
│   │       ├── ShortcutDialog/
│   │       │   └── ShortcutDialog.tsx
│   │       └── KeyboardShortcuts.types.ts
```

---

## VI. Testing Infrastructure

### 1. Test Utilities

```
packages/shared/src/
│
├── testing/
│   ├── utils/
│   │   ├── render.tsx  # Custom render with providers
│   │   ├── mock-providers.tsx  # Mock context providers
│   │   ├── test-data-generators.ts  # Generate test data
│   │   ├── mock-api.ts  # Mock API responses
│   │   └── test-helpers.ts  # Helper functions
│   ├── mocks/
│   │   ├── auth.ts  # Auth mocks
│   │   ├── users.ts  # User data mocks
│   │   ├── jobs.ts  # Job data mocks
│   │   ├── proposals.ts  # Proposal mocks
│   │   ├── contracts.ts  # Contract mocks
│   │   └── financial.ts  # Financial mocks
│   ├── fixtures/
│   │   ├── user.ts  # User fixtures
│   │   ├── job.ts  # Job fixtures
│   │   ├── proposal.ts  # Proposal fixtures
│   │   └── contract.ts  # Contract fixtures
│   └── setup/
│       ├── jest.setup.ts  # Jest configuration
│       ├── msw.setup.ts  # MSW setup
│       └── testing-library.setup.ts  # Testing Library setup
```

### 2. E2E Test Structure

```
apps/web/
│
├── e2e/
│   ├── tests/
│   │   ├── auth/
│   │   │   ├── login.spec.ts
│   │   │   ├── register.spec.ts
│   │   │   ├── password-reset.spec.ts
│   │   │   └── oauth.spec.ts
│   │   ├── jobs/
│   │   │   ├── post-job.spec.ts
│   │   │   ├── search-jobs.spec.ts
│   │   │   ├── apply-to-job.spec.ts
│   │   │   └── manage-jobs.spec.ts
│   │   ├── proposals/
│   │   │   ├── create-proposal.spec.ts
│   │   │   ├── edit-proposal.spec.ts
│   │   │   └── withdraw-proposal.spec.ts
│   │   ├── contracts/
│   │   │   ├── create-contract.spec.ts
│   │   │   ├── time-tracking.spec.ts
│   │   │   ├── submit-deliverable.spec.ts
│   │   │   └── dispute.spec.ts
│   │   ├── payments/
│   │   │   ├── add-payment-method.spec.ts
│   │   │   ├── escrow-funding.spec.ts
│   │   │   ├── release-payment.spec.ts
│   │   │   └── withdraw.spec.ts
│   │   └── messaging/
│   │       ├── send-message.spec.ts
│   │       ├── real-time-chat.spec.ts
│   │       └── notifications.spec.ts
│   ├── fixtures/
│   │   ├── test-users.json
│   │   ├── test-jobs.json
│   │   └── test-data.ts
│   ├── helpers/
│   │   ├── auth.ts
│   │   ├── navigation.ts
│   │   └── assertions.ts
│   └── playwright.config.ts
```

---

## VII. Documentation Structure

### 1. Architecture Documentation

```
docs/
│
├── architecture/
│   ├── overview.md  # System overview
│   ├── frontend/
│   │   ├── structure.md  # Folder structure
│   │   ├── routing.md  # Routing strategy
│   │   ├── state-management.md  # State management
│   │   └── authentication.md  # Auth flow
│   ├── backend-integration/
│   │   ├── api-patterns.md  # API patterns
│   │   ├── error-handling.md  # Error handling
│   │   ├── caching.md  # Caching strategy
│   │   └── real-time.md  # WebSocket/SSE
│   ├── mobile/
│   │   ├── architecture.md  # Mobile architecture
│   │   ├── offline-support.md  # Offline capabilities
│   │   └── performance.md  # Performance optimization
│   └── diagrams/
│       ├── system-architecture.mmd
│       ├── auth-flow.mmd
│       └── data-flow.mmd
```

### 2. API Documentation

```
docs/
│
├── api/
│   ├── getting-started.md
│   ├── authentication.md
│   ├── endpoints/
│   │   ├── users.md
│   │   ├── jobs.md
│   │   ├── proposals.md
│   │   ├── contracts.md
│   │   ├── payments.md
│   │   ├── messages.md
│   │   └── notifications.md
│   ├── webhooks.md
│   ├── rate-limiting.md
│   └── errors.md
```

### 3. Development Guides

```
docs/
│
├── guides/
│   ├── setup/
│   │   ├── local-development.md
│   │   ├── environment-variables.md
│   │   └── troubleshooting.md
│   ├── development/
│   │   ├── coding-standards.md
│   │   ├── git-workflow.md
│   │   ├── testing.md
│   │   └── debugging.md
│   ├── deployment/
│   │   ├── web.md
│   │   ├── mobile.md
│   │   └── ci-cd.md
│   └── contributing/
│       ├── getting-started.md
│       ├── pull-requests.md
│       └── code-review.md
```

### 4. Component Documentation

```
docs/
│
├── components/
│   ├── design-system.md
│   ├── atoms/
│   │   ├── button.md
│   │   ├── input.md
│   │   └── ...
│   ├── molecules/
│   │   ├── card.md
│   │   ├── form-field.md
│   │   └── ...
│   └── organisms/
│       ├── header.md
│       ├── sidebar.md
│       └── ...
```

---

## VIII. CI/CD Configuration

### 1. GitHub Actions

```
.github/
│
├── workflows/
│   ├── ci.yml  # Main CI pipeline
│   ├── deploy-web-staging.yml
│   ├── deploy-web-production.yml
│   ├── deploy-mobile-staging.yml
│   ├── deploy-mobile-production.yml
│   ├── test.yml  # Test runner
│   ├── lint.yml  # Linting
│   ├── type-check.yml  # TypeScript checks
│   ├── bundle-analysis.yml  # Bundle size checks
│   ├── lighthouse.yml  # Performance checks
│   ├── security-scan.yml  # Security scanning
│   └── dependency-update.yml  # Dependabot automation
│
├── actions/
│   ├── setup-pnpm/
│   │   └── action.yml
│   ├── cache-dependencies/
│   │   └── action.yml
│   └── deploy/
│       └── action.yml
│
└── CODEOWNERS  # Code ownership
```

### 2. Environment Configuration

```
fe/
│
├── .env.example  # Example environment variables
├── .env.local  # Local development (git-ignored)
├── .env.development  # Development environment
├── .env.staging  # Staging environment
├── .env.production  # Production environment
│
└── config/
    ├── environments/
    │   ├── development.ts
    │   ├── staging.ts
    │   └── production.ts
    └── index.ts  # Config loader
```

---

## IX. Monitoring & Observability

### 1. Error Tracking

```
packages/shared/src/
│
├── monitoring/
│   ├── error-tracking/
│   │   ├── sentry.ts  # Sentry configuration
│   │   ├── error-boundary.tsx  # Error boundary HOC
│   │   ├── error-reporter.ts  # Custom error reporter
│   │   └── error-filters.ts  # Filter errors
│   ├── performance/
│   │   ├── web-vitals.ts  # Core Web Vitals
│   │   ├── performance-observer.ts  # Performance monitoring
│   │   └── metrics.ts  # Custom metrics
│   └── analytics/
│       ├── google-analytics.ts  # GA4
│       ├── mixpanel.ts  # Mixpanel
│       ├── segment.ts  # Segment
│       └── event-tracker.ts  # Event tracking wrapper
```

### 2. Logging

```
packages/shared/src/
│
├── utils/
│   ├── logger/
│   │   ├── logger.ts  # Logger implementation
│   │   ├── log-levels.ts  # Log levels
│   │   ├── formatters/
│   │   │   ├── json.ts  # JSON formatter
│   │   │   └── pretty.ts  # Pretty formatter
│   │   └── transports/
│   │       ├── console.ts  # Console transport
│   │       ├── file.ts  # File transport
│   │       └── remote.ts  # Remote logging service
```

---

## X. Security Infrastructure

### 1. Security Headers & Middleware

```
apps/web/
│
├── middleware.ts  # Next.js middleware
│   # - CSP headers
│   # - CORS configuration
│   # - Rate limiting
│   # - Security headers (X-Frame-Options, etc.)
│
└── lib/
    ├── security/
    │   ├── csp.ts  # Content Security Policy
    │   ├── cors.ts  # CORS configuration
    │   ├── rate-limiter.ts  # Rate limiting
    │   └── headers.ts  # Security headers
```

### 2. Security Utilities

```
packages/shared/src/
│
├── security/
│   ├── csrf.ts  # CSRF token management
│   ├── sanitization.ts  # Input sanitization
│   │   # - HTML sanitization
│   │   # - SQL injection prevention
│   │   # - XSS prevention
│   ├── encryption.ts  # Client-side encryption
│   │   # - Encrypt sensitive data
│   │   # - Decrypt data
│   ├── validation.ts  # Input validation
│   │   # - Schema validation
│   │   # - Type checking
│   └── permissions.ts  # Permission checks
│       # - Role-based access
│       # - Feature flags
│       # - Resource ownership
```

---

## XI. Performance Optimization

### 1. Bundle Analysis

```
fe/
│
├── scripts/
│   ├── analyze-bundle.ts  # Bundle size analysis
│   ├── generate-bundle-report.ts  # Generate reports
│   └── check-bundle-limits.ts  # Enforce limits
│
└── .bundlewatch.config.json  # Bundle size limits
```

### 2. Image Optimization

```
apps/web/
│
├── lib/
│   ├── image-optimization/
│   │   ├── loader.ts  # Custom image loader
│   │   ├── placeholder.ts  # Blur placeholders
│   │   └── responsive.ts  # Responsive images
```

---

## Summary

This document completes the **final missing folder structure** based on all requirements in `fe-folder-structure-prompt.md`. 

### Major Areas Covered:

1. **Dashboard Routes (Additional)**:
   - Learning & Professional Development (courses, certifications, skill tests)
   - Community & Networking (forums, groups, events)
   - Enterprise/Agency Features (sub-accounts, talent pools, white-label, agency reporting)
   - Freelancer Specializations (Top Rated, Plus, Talent Cloud programs)

2. **Public/Marketing Routes**:
   - Solution pages (by role, by industry, enterprise, agencies)
   - Complete legal & compliance pages (terms, privacy, GDPR, CCPA, DMCA, accessibility)
   - Resource center (guides, case studies, whitepapers, webinars, tools)
   - Developer portal (API docs, SDKs, plugins, sample code)

3. **Mobile App Enhancements**:
   - Quick actions (quick apply, quick message, quick time entry)
   - Mobile onboarding (intro, permissions, biometric, tutorial)
   - Mobile-specific settings (app settings, security, offline, advanced)

4. **Shared Features**:
   - Webhooks management
   - Feature experimentation (A/B testing)
   - Real-time collaboration (presence, cursors, shared state)
   - Voice & video integration
   - Enhanced localization tools

5. **UI Components**:
   - Advanced inputs (rich text, code editor, markdown, signature pad, date range)
   - Data visualization (heatmap, Gantt, funnel, org chart)
   - AI/Smart components (AI assistant, smart suggestions, auto-complete, content generator)
   - Accessibility components (skip links, screen reader, focus trap, live region, keyboard shortcuts)

6. **Infrastructure**:
   - Testing utilities and E2E structure
   - Documentation (architecture, API, guides, components)
   - CI/CD configuration (GitHub Actions, environment config)
   - Monitoring & observability (error tracking, performance, analytics, logging)
   - Security infrastructure (headers, middleware, utilities)
   - Performance optimization (bundle analysis, image optimization)

All routes include proper backend mappings with microservice, domain, HTTP method, and endpoint information as specified in the original requirements.

---

**END OF FINAL COMPREHENSIVE MISSING FRONTEND FOLDER STRUCTURE**
