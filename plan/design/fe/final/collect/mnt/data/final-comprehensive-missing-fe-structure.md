fe/
├── .husky/
│   ├── pre-commit
│   └── pre-push
├── .vscode/
│   ├── extensions.json
│   ├── launch.json
│   └── settings.json
├── .github/
│   ├── CODEOWNERS  # Code ownership
│   ├── actions/
│   │   ├── cache-dependencies/
│   │   │   └── action.yml
│   │   ├── deploy/
│   │   │   └── action.yml
│   │   └── setup-pnpm/
│   │       └── action.yml
│   └── workflows/
│       ├── bundle-analysis.yml  # Bundle size checks
│       ├── ci.yml  # Main CI pipeline
│       ├── dependency-update.yml  # Dependabot automation
│       ├── deploy-mobile-production.yml
│       ├── deploy-mobile-staging.yml
│       ├── deploy-web-production.yml
│       ├── deploy-web-staging.yml
│       ├── lighthouse.yml  # Performance checks
│       ├── lint.yml  # Linting
│       ├── security-scan.yml  # Security scanning
│       ├── test.yml  # Test runner
│       └── type-check.yml  # TypeScript checks
├── config/
│   ├── environments/
│   │   ├── development.ts
│   │   ├── production.ts
│   │   └── staging.ts
│   └── index.ts  # Config loader
├── apps/
│   ├── mobile/
│   │   ├── app/
│   │   │   ├── (auth)/
│   │   │   │   ├── _layout.tsx
│   │   │   │   ├── callback.tsx
│   │   │   │   ├── login.tsx
│   │   │   │   └── register.tsx
│   │   │   ├── _layout.tsx
│   │   │   ├── +not-found.tsx
│   │   │   ├── index.tsx
│   │   │   ├── onboarding/
│   │   │   │   ├── biometric/
│   │   │   │   │   └── setup.tsx  # Biometric auth setup
│   │   │   │   │       # - Face ID
│   │   │   │   │       # - Touch ID
│   │   │   │   │       # BE: users-be/auth
│   │   │   │   │       # POST /v1/auth/biometric/enable
│   │   │   │   ├── intro/
│   │   │   │   │   └── index.tsx  # App intro screens
│   │   │   │   │       # - Swipeable intro
│   │   │   │   │       # - Feature highlights
│   │   │   │   ├── permissions/
│   │   │   │   │   ├── camera.tsx  # Camera permission
│   │   │   │   │   ├── index.tsx  # All permissions
│   │   │   │   │   ├── location.tsx  # Location permission
│   │   │   │   │   └── notifications.tsx  # Notification permission
│   │   │   │   └── tutorial/
│   │   │   │       ├── apply.tsx  # Tutorial: Apply
│   │   │   │       ├── browse-jobs.tsx  # Tutorial: Browse jobs
│   │   │   │       ├── time-tracking.tsx  # Tutorial: Time tracking
│   │   │   │       └── index.tsx  # Interactive tutorial
│   │   │   ├── quick-actions/
│   │   │   │   ├── quick-apply/
│   │   │   │   │   └── [jobId].tsx  # Quick apply to job
│   │   │   │   │       # - Pre-filled proposal
│   │   │   │   │       # - One-tap apply
│   │   │   │   │       # - Connect deduction
│   │   │   │   │       # BE: proposals-be/proposal
│   │   │   │   │       # POST /v1/proposals/quick-apply
│   │   │   │   ├── quick-message/
│   │   │   │   │   └── [userId].tsx  # Quick message to user
│   │   │   │   │       # - Template messages
│   │   │   │   │       # - Voice-to-text
│   │   │   │   │       # BE: communications-be/conversation
│   │   │   │   │       # POST /v1/conversations/{conversation_id}/messages/quick
│   │   │   │   └── quick-time-entry/
│   │   │   │       └── [contractId].tsx  # Quick time logging
│   │   │   │           # - Timer widget
│   │   │   │           # - Quick notes
│   │   │   │           # - One-tap submit
│   │   │   │           # BE: contracts-be/work_diary
│   │   │   │           # POST /v1/work-diary/quick-entry
│   │   └── src/
│   └── web/
│       ├── e2e/
│       │   ├── fixtures/
│       │   │   ├── test-data.ts
│       │   │   ├── test-jobs.json
│       │   │   └── test-users.json
│       │   ├── helpers/
│       │   │   ├── assertions.ts
│       │   │   ├── auth.ts
│       │   │   └── navigation.ts
│       │   ├── tests/
│       │   │   ├── auth/
│       │   │   │   ├── login.spec.ts
│       │   │   │   ├── oauth.spec.ts
│       │   │   │   ├── password-reset.spec.ts
│       │   │   │   └── register.spec.ts
│       │   │   ├── contracts/
│       │   │   │   ├── create-contract.spec.ts
│       │   │   │   ├── dispute.spec.ts
│       │   │   │   ├── submit-deliverable.spec.ts
│       │   │   │   └── time-tracking.spec.ts
│       │   │   ├── jobs/
│       │   │   │   ├── apply-to-job.spec.ts
│       │   │   │   ├── manage-jobs.spec.ts
│       │   │   │   ├── post-job.spec.ts
│       │   │   │   └── search-jobs.spec.ts
│       │   │   ├── messaging/
│       │   │   │   ├── notifications.spec.ts
│       │   │   │   ├── real-time-chat.spec.ts
│       │   │   │   └── send-message.spec.ts
│       │   │   ├── payments/
│       │   │   │   ├── add-payment-method.spec.ts
│       │   │   │   ├── escrow-funding.spec.ts
│       │   │   │   ├── release-payment.spec.ts
│       │   │   │   └── withdraw.spec.ts
│       │   │   └── proposals/
│       │   │       ├── create-proposal.spec.ts
│       │   │       ├── edit-proposal.spec.ts
│       │   │       └── withdraw-proposal.spec.ts
│       │   └── playwright.config.ts
│       ├── lib/
│       │   ├── image-optimization/
│       │   │   ├── loader.ts  # Custom image loader
│       │   │   ├── placeholder.ts  # Blur placeholders
│       │   │   └── responsive.ts  # Responsive images
│       │   └── security/
│       │       ├── csp.ts  # Content Security Policy
│       │       ├── cors.ts  # CORS configuration
│       │       ├── headers.ts  # Security headers
│       │       └── rate-limiter.ts  # Rate limiting
│       ├── middleware.ts  # Next.js middleware
│       │   # - CSP headers
│       │   # - CORS configuration
│       │   # - Rate limiting
│       │   # - Security headers (X-Frame-Options, etc.)
│       └── src/
│           └── app/
│               └── [locale]/
│                   ├── (dashboard)/
│                   │   ├── agency/
│                   │   │   ├── overview/
│                   │   │   │   └── page.tsx  # Agency dashboard
│                   │   │   │       # - Revenue overview
│                   │   │   │       # - Active projects
│                   │   │   │       # - Team utilization
│                   │   │   │       # - Client list
│                   │   │   │       # BE: users-be/org, financial-be/analytics
│                   │   │   │       # GET /v1/agencies/me/overview
│                   │   │   ├── reporting/
│                   │   │   │   ├── clients/
│                   │   │   │   │   └── page.tsx  # Client reports
│                   │   │   │   │       # - Per-client spending
│                   │   │   │   │       # - Project status
│                   │   │   │   │       # BE: financial-be/analytics
│                   │   │   │   │       # GET /v1/agencies/reports/clients
│                   │   │   │   ├── financial/
│                   │   │   │   │   └── page.tsx  # Financial reports
│                   │   │   │   │       # - Revenue by period
│                   │   │   │   │       # - Margins
│                   │   │   │   │       # - Forecasts
│                   │   │   │   │       # BE: financial-be/analytics
│                   │   │   │   │       # GET /v1/agencies/reports/financial
│                   │   │   │   └── team/
│                   │   │   │       └── page.tsx  # Team performance
│                   │   │   │           # - Utilization rates
│                   │   │   │           # - Project assignments
│                   │   │   │           # BE: users-be/org, jobs-be/analytics
│                   │   │   │           # GET /v1/agencies/reports/team
│                   │   │   ├── sub-accounts/
│                   │   │   │   ├── [subAccountId]/
│                   │   │   │   │   ├── settings/
│                   │   │   │   │   │   └── page.tsx  # Sub-account settings
│                   │   │   │   │   │       # BE: users-be/org
│                   │   │   │   │   │       # PUT /v1/agencies/sub-accounts/{sub_account_id}
│                   │   │   │   │   └── page.tsx  # Sub-account detail
│                   │   │   │   │       # - Jobs posted
│                   │   │   │   │       # - Contracts
│                   │   │   │   │       # - Spending
│                   │   │   │   │       # BE: users-be/org
│                   │   │   │   │       # GET /v1/agencies/sub-accounts/{sub_account_id}
│                   │   │   │   ├── create/
│                   │   │   │   │   └── page.tsx  # Create sub-account
│                   │   │   │   │       # BE: users-be/org
│                   │   │   │   │       # POST /v1/agencies/sub-accounts
│                   │   │   │   └── page.tsx  # All sub-accounts
│                   │   │   │       # - List all
│                   │   │   │       # - Manage
│                   │   │   │       # BE: users-be/org
│                   │   │   │       # GET /v1/agencies/sub-accounts
│                   │   │   ├── talent-pool/
│                   │   │   │   ├── [poolId]/
│                   │   │   │   │   ├── members/
│                   │   │   │   │   │   └── page.tsx  # Pool members
│                   │   │   │   │   │       # - Freelancer list
│                   │   │   │   │   │       # - Add/remove
│                   │   │   │   │   │       # BE: users-be/org, search-be/talent-pool
│                   │   │   │   │   │       # GET /v1/talent-pools/{pool_id}/members
│                   │   │   │   │   └── page.tsx  # Talent pool detail
│                   │   │   │   │       # BE: search-be/talent-pool
│                   │   │   │   │       # GET /v1/talent-pools/{pool_id}
│                   │   │   │   ├── create/
│                   │   │   │   │   └── page.tsx  # Create talent pool
│                   │   │   │   │       # BE: search-be/talent-pool
│                   │   │   │   │       # POST /v1/talent-pools
│                   │   │   │   └── page.tsx  # All talent pools
│                   │   │   │       # BE: search-be/talent-pool
│                   │   │   │       # GET /v1/talent-pools
│                   │   │   └── white-label/
│                   │   │       └── page.tsx  # White-label settings
│                   │   │           # - Branding
│                   │   │           # - Custom domain
│                   │   │           # - Logo/colors
│                   │   │           # BE: users-be/org
│                   │   │           # PUT /v1/agencies/white-label
│                   │   ├── community/
│                   │   │   ├── events/
│                   │   │   │   ├── [eventId]/
│                   │   │   │   │   ├── register/
│                   │   │   │   │   │   └── page.tsx  # Event registration
│                   │   │   │   │   │       # - RSVP form
│                   │   │   │   │   │       # - Add to calendar
│                   │   │   │   │   │       # BE: communications-be/events
│                   │   │   │   │   │       # POST /v1/events/{event_id}/register
│                   │   │   │   │   └── page.tsx  # Event detail
│                   │   │   │   │       # - Info
│                   │   │   │   │       # - Attendees
│                   │   │   │   │       # - Check-in (QR code)
│                   │   │   │   │       # BE: communications-be/events
│                   │   │   │   │       # GET /v1/events/{event_id}
│                   │   │   │   ├── my-events/
│                   │   │   │   │   └── page.tsx  # My registered events
│                   │   │   │   │       # - Attending
│                   │   │   │   │       # - Past attended
│                   │   │   │   │       # BE: communications-be/events
│                   │   │   │   │       # GET /v1/users/me/events
│                   │   │   │   └── upcoming/
│                   │   │   │       └── page.tsx  # Upcoming events
│                   │   │   │           # BE: communications-be/events
│                   │   │   │           # GET /v1/events?status=upcoming
│                   │   │   ├── forums/
│                   │   │   │   ├── [forumId]/
│                   │   │   │   │   ├── threads/
│                   │   │   │   │   │   └── [threadId]/
│                   │   │   │   │   │       ├── page.tsx  # Thread view
│                   │   │   │   │   │       │   # - Posts
│                   │   │   │   │   │       │   # - Reply
│                   │   │   │   │   │       │   # - Voting
│                   │   │   │   │   │       │   # BE: communications-be/forums (if exists) OR community service
│                   │   │   │   │   │       │   # GET /v1/forums/{forum_id}/threads/{thread_id}
│                   │   │   │   │   │       └── reply/
│                   │   │   │   │   │           └── page.tsx  # Reply to thread
│                   │   │   │   │   │               # BE: communications-be/forums
│                   │   │   │   │   │               # POST /v1/threads/{thread_id}/replies
│                   │   │   │   │   └── page.tsx  # Forum overview
│                   │   │   │   │       # - Thread list
│                   │   │   │   │       # - Create thread
│                   │   │   │   │       # BE: communications-be/forums
│                   │   │   │   │       # GET /v1/forums/{forum_id}/threads
│                   │   │   │   └── page.tsx  # All forums
│                   │   │   │       # - Categories
│                   │   │   │       # - Popular threads
│                   │   │   │       # BE: communications-be/forums
│                   │   │   │       # GET /v1/forums
│                   │   │   ├── groups/
│                   │   │   │   ├── [groupId]/
│                   │   │   │   │   ├── discussions/
│                   │   │   │   │   │   └── page.tsx  # Group discussions
│                   │   │   │   │   │       # - Posts feed
│                   │   │   │   │   │       # - Comment
│                   │   │   │   │   │       # BE: communications-be/discussions
│                   │   │   │   │   │       # GET /v1/groups/{group_id}/discussions
│                   │   │   │   │   ├── events/
│                   │   │   │   │   │   ├── [eventId]/
│                   │   │   │   │   │   │   └── page.tsx  # Event detail
│                   │   │   │   │   │   │       # - RSVP
│                   │   │   │   │   │   │       # - Calendar add
│                   │   │   │   │   │   │       # BE: communications-be/events OR community service
│                   │   │   │   │   │   │       # GET /v1/events/{event_id}
│                   │   │   │   │   │   └── page.tsx  # Group events
│                   │   │   │   │   │       # - Upcoming
│                   │   │   │   │   │       # - Past
│                   │   │   │   │   │       # BE: communications-be/events
│                   │   │   │   │   │       # GET /v1/groups/{group_id}/events
│                   │   │   │   │   ├── members/
│                   │   │   │   │   │   └── page.tsx  # Group members
│                   │   │   │   │   │       # - Member list
│                   │   │   │   │   │       # - Invite
│                   │   │   │   │   │       # - Roles
│                   │   │   │   │   │       # BE: users-be/groups (if exists) OR community service
│                   │   │   │   │   │       # GET /v1/groups/{group_id}/members
│                   │   │   │   │   └── page.tsx  # Group overview
│                   │   │   │   │       # - About
│                   │   │   │   │       # - Join/leave
│                   │   │   │   │       # - Activity feed
│                   │   │   │   │       # BE: users-be/groups
│                   │   │   │   │       # GET /v1/groups/{group_id}
│                   │   ├── learning/
│                   │   │   ├── certifications/
│                   │   │   │   ├── [certId]/
│                   │   │   │   │   ├── verify/
│                   │   │   │   │   │   └── page.tsx  # Verify certificate
│                   │   │   │   │   │       # BE: utility-be/learning
│                   │   │   │   │   │       # GET /v1/certifications/{cert_id}/verify
│                   │   │   │   │   └── page.tsx  # Certificate detail
│                   │   │   │   │       # - Download PDF
│                   │   │   │   │       # - Share link
│                   │   │   │   │       # - Add to profile
│                   │   │   │   │       # BE: utility-be/learning
│                   │   │   │   │       # GET /v1/certifications/{cert_id}
│                   │   │   │   └── page.tsx  # All certifications
│                   │   │   │       # - Earned certificates
│                   │   │   │       # - Available certifications
│                   │   │   │       # BE: utility-be/learning
│                   │   │   │       # GET /v1/users/me/certifications
│                   │   │   ├── courses/
│                   │   │   │   ├── [courseId]/
│                   │   │   │   │   ├── assessments/
│                   │   │   │   │   │   ├── [assessmentId]/
│                   │   │   │   │   │   │   ├── results/
│                   │   │   │   │   │   │   │   └── page.tsx  # View results
│                   │   │   │   │   │   │   │       # BE: utility-be/learning
│                   │   │   │   │   │   │   │       # GET /v1/assessments/{assessment_id}/results
│                   │   │   │   │   │   │   ├── start/
│                   │   │   │   │   │   │   │   └── page.tsx  # Start assessment
│                   │   │   │   │   │   │   │       # BE: utility-be/learning
│                   │   │   │   │   │   │   │       # POST /v1/assessments/{assessment_id}/start
│                   │   │   │   │   │   │   └── submit/
│                   │   │   │   │   │   │       └── page.tsx  # Submit answers
│                   │   │   │   │   │   │           # BE: utility-be/learning
│                   │   │   │   │   │   │           # POST /v1/assessments/{assessment_id}/submit
│                   │   │   │   │   ├── lessons/
│                   │   │   │   │   │   └── [lessonId]/
│                   │   │   │   │   │       └── page.tsx  # Lesson view
│                   │   │   │   │   │           # - Video/content player
│                   │   │   │   │   │           # - Notes taking
│                   │   │   │   │   │           # - Progress tracking
│                   │   │   │   │   │           # BE: utility-be/learning OR external LMS
│                   │   │   │   │   │           # GET /v1/courses/{course_id}/lessons/{lesson_id}
│                   │   │   │   │   └── page.tsx  # Course overview
│                   │   │   │   │       # - Syllabus
│                   │   │   │   │       # - Progress
│                   │   │   │   │       # - Enroll button
│                   │   │   │   │       # BE: utility-be/learning
│                   │   │   │   │       # GET /v1/courses/{course_id}
│                   │   │   │   │       # POST /v1/courses/{course_id}/enroll
│                   │   │   │   ├── browse/
│                   │   │   │   │   └── page.tsx  # Browse all courses
│                   │   │   │   │       # - Filter by skill
│                   │   │   │   │       # - Search
│                   │   │   │   │       # - Recommendations
│                   │   │   │   │       # BE: utility-be/learning
│                   │   │   │   │       # GET /v1/courses
│                   │   │   │   └── my-courses/
│                   │   │   │       └── page.tsx  # Enrolled courses
│                   │   │   │           # - In progress
│                   │   │   │           # - Completed
│                   │   │   │           # - Certificates
│                   │   │   │           # BE: utility-be/learning
│                   │   │   │           # GET /v1/users/me/courses
│                   │   │   └── skill-tests/
│                   │   │       ├── [testId]/
│                   │   │       │   ├── instructions/
│                   │   │       │   │   └── page.tsx  # Test instructions
│                   │   │       │   │       # BE: users-be/capabilities
│                   │   │       │   │       # GET /v1/skill-tests/{test_id}
│                   │   │       │   ├── results/
│                   │   │       │   │   └── page.tsx  # Test results
│                   │   │       │   │       # - Score
│                   │   │       │   │       # - Percentile
│                   │   │       │   │       # - Badge earned
│                   │   │       │   │       # BE: users-be/capabilities
│                   │   │       │   │       # GET /v1/skill-tests/{test_id}/results
│                   │   │       │   └── take-test/
│                   │   │       │       └── page.tsx  # Take the test
│                   │   │       │           # - Timed test interface
│                   │   │       │           # - Submit answers
│                   │   │       │           # BE: users-be/capabilities
│                   │   │       │           # POST /v1/skill-tests/{test_id}/submit
│                   │   │       └── page.tsx  # Available skill tests
│                   │   │           # - Browse by skill
│                   │   │           # - Test history
│                   │   │           # - Top scores
│                   │   │           # BE: users-be/capabilities
│                   │   │           # GET /v1/skill-tests
│                   │   ├── search/
│                   │   │   ├── freelancers/
│                   │   │   │   └── page.tsx  # Advanced freelancer search (client)
│                   │   │   │       # - Search by skills
│                   │   │   └── jobs/
│                   │   │       └── page.tsx  # Advanced job search
│                   │   │           # - Full-text search
│                   │   │           # - Faceted filters
│                   │   │           # - Autocomplete suggestions
│                   │   │           # - Search history
│                   │   │           # - Save search
│                   │   │           # BE: search-be/query
│                   │   │           # POST /v1/search/jobs
│                   │   │           # Body: { query, filters: {...}, sort, page }
│                   │   │           # BE: search-be/suggestions
│                   │   │           # GET /v1/suggestions?q={query}
│                   │   └── specialized/
│                   │       ├── plus/
│                   │       │   ├── subscribe/
│                   │       │   │   └── page.tsx  # Subscribe to Plus
│                   │       │   │       # - Plans
│                   │       │   │       # - Payment
│                   │       │   │       # BE: subscriptions-be/subscription (financial-be)
│                   │       │   │       # POST /v1/subscriptions/plus
│                   │       │   └── page.tsx  # Plus membership overview
│                   │       │       # - Benefits
│                   │       │       # - Pricing
│                   │       │       # - Comparison
│                   │       │       # BE: subscriptions-be/plans
│                   │       │       # GET /v1/subscriptions/plans/plus
│                   │       ├── talent-cloud/
│                   │       │   ├── projects/
│                   │       │   │   └── page.tsx  # Talent Cloud exclusive projects
│                   │       │   │       # - High-value projects
│                   │       │   │       # - Direct invites
│                   │       │   │       # BE: jobs-be/job (with exclusive flag)
│                   │       │   │       # GET /v1/jobs/talent-cloud
│                   │       │   └── page.tsx  # Talent Cloud info
│                   │       │       # - Qualification
│                   │       │       # - Apply
│                   │       │       # - Status
│                   │       │       # BE: users-be/badges
│                   │       │       # GET /v1/programs/talent-cloud
│                   │       └── top-rated/
│                   │           ├── application/
│                   │           │   └── page.tsx  # Apply for Top Rated
│                   │           │       # - Eligibility check
│                   │           │       # - Submit application
│                   │           │       # BE: users-be/badges OR reviews-be/reputation
│                   │           │       # POST /v1/users/me/top-rated/apply
│                   │           └── page.tsx  # Top Rated program info
│                   │               # - Benefits
│                   │               # - Requirements
│                   │               # - Application status
│                   │               # BE: users-be/badges
│                   │               # GET /v1/programs/top-rated
│                   └── (public)/
│                       ├── developers/
│                       │   ├── api/
│                       │   │   ├── authentication/
│                       │   │   │   └── page.tsx  # OAuth & API keys
│                       │   │   ├── endpoints/
│                       │   │   │   ├── [category]/
│                       │   │   │   │   └── page.tsx  # Endpoint category
│                       │   │   │   └── page.tsx  # All endpoints
│                       │   │   ├── getting-started/
│                       │   │   │   └── page.tsx  # API getting started
│                       │   │   ├── rate-limits/
│                       │   │   │   └── page.tsx  # Rate limiting info
│                       │   │   ├── webhooks/
│                       │   │   │   └── page.tsx  # Webhook documentation
│                       │   │   └── page.tsx  # API overview
│                       │   ├── community/
│                       │   │   ├── forum/
│                       │   │   │   └── page.tsx  # Developer forum
│                       │   │   └── github/
│                       │   │       └── page.tsx  # GitHub repos
│                       │   ├── plugins/
│                       │   │   ├── shopify/
│                       │   │   │   └── page.tsx  # Shopify app
│                       │   │   └── wordpress/
│                       │   │       └── page.tsx  # WordPress plugin
│                       │   └── sdks/
│                       │       ├── javascript/
│                       │       │   └── page.tsx  # JavaScript SDK
│                       │       ├── python/
│                       │       │   └── page.tsx  # Python SDK
│                       │       ├── ruby/
│                       │       │   └── page.tsx  # Ruby SDK
│                       │       └── page.tsx  # All SDKs
│                       ├── legal/
│                       │   ├── accessibility/
│                       │   │   └── page.tsx  # Accessibility statement
│                       │   ├── compliance/
│                       │   │   ├── aml/
│                       │   │   │   └── page.tsx  # AML policy
│                       │   │   ├── kyc/
│                       │   │   │   └── page.tsx  # KYC policy
│                       │   │   └── page.tsx  # Compliance overview
│                       │   ├── dmca/
│                       │   │   └── page.tsx  # DMCA policy
│                       │   ├── licenses/
│                       │   │   └── page.tsx  # Open source licenses
│                       │   ├── privacy/
│                       │   │   ├── ccpa/
│                       │   │   │   └── page.tsx  # CCPA information
│                       │   │   ├── cookies/
│                       │   │   │   └── page.tsx  # Cookie Policy
│                       │   │   ├── gdpr/
│                       │   │   │   └── page.tsx  # GDPR information
│                       │   │   └── policy/
│                       │   │       └── page.tsx  # Privacy Policy
│                       │   └── terms/
│                       │       ├── client/
│                       │       │   └── page.tsx  # Client Agreement
│                       │       ├── freelancer/
│                       │       │   └── page.tsx  # Freelancer Agreement
│                       │       ├── service/
│                       │       │   └── page.tsx  # Terms of Service
│                       │       └── page.tsx  # Legal terms overview
│                       ├── resources/
│                       │   ├── api/
│                       │   │   ├── changelog/
│                       │   │   │   └── page.tsx  # API changelog
│                       │   │   ├── docs/
│                       │   │   │   └── page.tsx  # API documentation
│                       │   │   └── reference/
│                       │   │       └── page.tsx  # API reference
│                       │   ├── case-studies/
│                       │   │   ├── [slug]/
│                       │   │   │   └── page.tsx  # Case study detail
│                       │   │   └── page.tsx  # All case studies
│                       │   ├── guides/
│                       │   │   ├── [slug]/
│                       │   │   │   └── page.tsx  # Guide detail
│                       │   │   └── page.tsx  # All guides
│                       │   │       # - Getting started
│                       │   │       # - Best practices
│                       │   │       # - How-tos
│                       │   ├── tools/
│                       │   │   ├── budget-estimator/
│                       │   │   │   └── page.tsx  # Project budget estimator
│                       │   │   ├── rate-calculator/
│                       │   │   │   └── page.tsx  # Hourly rate calculator
│                       │   │   └── roi-calculator/
│                       │   │       └── page.tsx  # ROI calculator
│                       │   └── webinars/
│                       │       ├── on-demand/
│                       │       │   └── page.tsx  # On-demand recordings
│                       │       ├── upcoming/
│                       │       │   └── page.tsx  # Upcoming webinars
│                       │       └── [slug]/
│                       │           ├── register/
│                       │           │   └── page.tsx  # Webinar registration
│                       │           │       # BE: communications-be/events
│                       │           │       # POST /v1/webinars/{webinar_id}/register
│                       │           └── page.tsx  # Webinar detail
│                       └── solutions/
│                           ├── agencies/
│                           │   └── page.tsx  # Agency solutions
│                           │       # - White-label
│                           │       # - Multi-client
│                           │       # - Revenue share
│                           ├── by-industry/
│                           │   ├── finance/
│                           │   │   └── page.tsx  # Finance solutions
│                           │   ├── marketing/
│                           │   │   └── page.tsx  # Marketing solutions
│                           │   ├── tech/
│                           │   │   └── page.tsx  # Tech industry solutions
│                           │   └── page.tsx  # Browse by industry
│                           ├── by-role/
│                           │   ├── designers/
│                           │   │   └── page.tsx  # Designer hiring solutions
│                           │   ├── developers/
│                           │   │   └── page.tsx  # Developer hiring solutions
│                           │   ├── writers/
│                           │   │   └── page.tsx  # Writer hiring solutions
│                           │   └── page.tsx  # Browse by role
│                           └── enterprise/
│                               ├── contact-sales/
│                               │   └── page.tsx  # Enterprise contact form
│                               │       # BE: communications-be/contact
│                               │       # POST /v1/contact/enterprise
│                               └── page.tsx  # Enterprise solutions
│                                   # - MSA contracts
│                                   # - Dedicated support
│                                   # - Volume pricing
├── packages/
│   ├── shared/
│   │   └── src/
│   │       ├── features/
│   │       │   ├── collaboration/
│   │       │   │   ├── api/
│   │       │   │   │   └── collaboration-api.ts  # Collaboration API
│   │       │   │   ├── components/
│   │       │   │   │   ├── ActiveUsers/
│   │       │   │   │   │   ├── ActiveUsers.tsx
│   │       │   │   │   │   └── ActiveUsers.types.ts
│   │       │   │   │   ├── CollaboratorCursor/
│   │       │   │   │   │   ├── CollaboratorCursor.tsx
│   │       │   │   │   │   └── CollaboratorCursor.types.ts
│   │       │   │   │   └── PresenceIndicator/
│   │       │   │   │       ├── PresenceIndicator.tsx
│   │       │   │   │       └── PresenceIndicator.types.ts
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useCollaboration.ts  # Collaboration session
│   │       │   │   │   ├── useCursors.ts  # Cursor tracking
│   │       │   │   │   ├── usePresence.ts  # User presence
│   │       │   │   │   └── useSharedState.ts  # Shared state sync
│   │       │   │   ├── providers/
│   │       │   │   │   └── CollaborationProvider.tsx  # Collab context
│   │       │   │   └── types.ts  # Collaboration types
│   │       │   ├── experiments/
│   │       │   │   ├── api/
│   │       │   │   │   └── experiments-api.ts  # Experiments API
│   │       │   │   │       # BE: utility-be/experiments
│   │       │   │   │       # GET /v1/experiments/active
│   │       │   │   │       # POST /v1/experiments/{experiment_id}/track
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useExperiment.ts  # Get experiment variant
│   │       │   │   │   ├── useExperimentTracking.ts  # Track experiment events
│   │       │   │   │   └── useFeatureVariant.ts  # Feature variant
│   │       │   │   ├── providers/
│   │       │   │   │   └── ExperimentProvider.tsx  # Experiment context
│   │       │   │   ├── types.ts  # Experiment types
│   │       │   │   └── utils/
│   │       │   │       ├── experiment-storage.ts  # Store assignments
│   │       │   │       ├── tracking.ts  # Track events
│   │       │   │       └── variant-assignment.ts  # Assign variant
│   │       │   ├── i18n/
│   │       │   │   ├── components/
│   │       │   │   │   ├── FormattedMessage/
│   │       │   │   │   │   └── FormattedMessage.tsx
│   │       │   │   │   ├── LocaleSwitcher/
│   │       │   │   │   │   ├── LocaleSwitcher.native.tsx
│   │       │   │   │   │   ├── LocaleSwitcher.tsx
│   │       │   │   │   │   └── LocaleSwitcher.web.tsx
│   │       │   │   │   └── TranslationProvider/
│   │       │   │   │       └── TranslationProvider.tsx
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useCurrencyFormat.ts  # Currency formatting
│   │       │   │   │   ├── useDateFormat.ts  # Date formatting
│   │       │   │   │   ├── useNumberFormat.ts  # Number formatting
│   │       │   │   │   ├── useRTL.ts  # RTL detection
│   │       │   │   │   └── useTranslation.ts  # (already exists, enhanced)
│   │       │   │   ├── tools/
│   │       │   │   │   ├── currency-formatter.ts  # Currency formatting
│   │       │   │   │   ├── missing-keys-detector.ts  # Detect missing keys
│   │       │   │   │   ├── pluralization-helper.ts  # Pluralization
│   │       │   │   │   └── translation-manager.ts  # Manage translations
│   │       │   │   └── utils/
│   │       │   │       ├── fallback-loader.ts  # Fallback translations
│   │       │   │       ├── interpolation.ts  # Message interpolation
│   │       │   │       └── locale-detection.ts  # Auto-detect locale
│   │       │   ├── video/
│   │       │   │   ├── api/
│   │       │   │   │   └── video-api.ts  # Video call API
│   │       │   │   │       # BE: communications-be/video
│   │       │   │   │       # POST /v1/video/rooms
│   │       │   │   │       # GET /v1/video/rooms/{room_id}/token
│   │       │   │   ├── components/
│   │       │   │   │   ├── ParticipantGrid/
│   │       │   │   │   │   ├── ParticipantGrid.tsx
│   │       │   │   │   │   └── ParticipantGrid.types.ts
│   │       │   │   │   ├── VideoControls/
│   │       │   │   │   │   ├── VideoControls.tsx
│   │       │   │   │   │   └── VideoControls.types.ts
│   │       │   │   │   └── VideoRoom/
│   │       │   │   │       ├── VideoRoom.native.tsx
│   │       │   │   │       ├── VideoRoom.tsx
│   │       │   │   │       ├── VideoRoom.web.tsx
│   │       │   │   │       └── VideoRoom.types.ts
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useRecording.ts  # Call recording
│   │       │   │   │   ├── useScreenShare.ts  # Screen sharing
│   │       │   │   │   ├── useVideoCall.ts  # Video call management
│   │       │   │   │   └── useVideoDevices.ts  # Device selection
│   │       │   │   └── types.ts  # Video types
│   │       │   └── webhooks/
│   │       │       ├── api/
│   │       │       │   └── webhooks-api.ts  # Webhooks API client
│   │       │       │       # BE: utility-be/webhooks OR admin-be
│   │       │       │       # GET /v1/webhooks
│   │       │       │       # POST /v1/webhooks
│   │       │       │       # PUT /v1/webhooks/{webhook_id}
│   │       │       │       # DELETE /v1/webhooks/{webhook_id}
│   │       │       │       # POST /v1/webhooks/{webhook_id}/test
│   │       │       ├── components/
│   │       │       │   ├── EventSelector/
│   │       │       │   │   ├── EventSelector.tsx
│   │       │       │   │   └── EventSelector.types.ts
│   │       │       │   ├── WebhookForm/
│   │       │       │   │   ├── WebhookForm.tsx
│   │       │       │   │   ├── WebhookForm.types.ts
│   │       │       │   │   └── WebhookForm.web.tsx
│   │       │       │   └── WebhookLogs/
│   │       │       │       ├── WebhookLogs.tsx
│   │       │       │       └── WebhookLogs.types.ts
│   │       │       ├── hooks/
│   │       │       │   ├── useWebhook.ts  # Single webhook
│   │       │       │   ├── useWebhookLogs.ts  # Webhook delivery logs
│   │       │       │   ├── useWebhookTest.ts  # Test webhook
│   │       │       │   └── useWebhooks.ts  # List webhooks
│   │       │       ├── queries/
│   │       │       │   ├── webhooks-mutations.ts  # Webhook mutations
│   │       │       │   └── webhooks-queries.ts  # Webhook queries
│   │       │       └── types.ts  # Webhook types
│   │       ├── monitoring/
│   │       │   ├── analytics/
│   │       │   │   ├── event-tracker.ts  # Event tracking wrapper
│   │       │   │   ├── google-analytics.ts  # GA4
│   │       │   │   ├── mixpanel.ts  # Mixpanel
│   │       │   │   └── segment.ts  # Segment
│   │       │   ├── error-tracking/
│   │       │   │   ├── error-boundary.tsx  # Error boundary HOC
│   │       │   │   ├── error-filters.ts  # Filter errors
│   │       │   │   ├── error-reporter.ts  # Custom error reporter
│   │       │   │   └── sentry.ts  # Sentry configuration
│   │       │   └── performance/
│   │       │       ├── metrics.ts  # Custom metrics
│   │       │       ├── performance-observer.ts  # Performance monitoring
│   │       │       └── web-vitals.ts  # Core Web Vitals
│   │       ├── security/
│   │       │   ├── csrf.ts  # CSRF token management
│   │       │   ├── encryption.ts  # Client-side encryption
│   │       │   │   # - Encrypt sensitive data
│   │       │   │   # - Decrypt data
│   │       │   ├── permissions.ts  # Permission checks
│   │       │   │   # - Role-based access
│   │       │   │   # - Feature flags
│   │       │   │   # - Resource ownership
│   │       │   ├── sanitization.ts  # Input sanitization
│   │       │   │   # - HTML sanitization
│   │       │   │   # - SQL injection prevention
│   │       │   │   # - XSS prevention
│   │       │   └── validation.ts  # Input validation
│   │       ├── testing/
│   │       │   ├── fixtures/
│   │       │   │   ├── contract.ts  # Contract fixtures
│   │       │   │   ├── job.ts  # Job fixtures
│   │       │   │   ├── proposal.ts  # Proposal fixtures
│   │       │   │   └── user.ts  # User fixtures
│   │       │   ├── mocks/
│   │       │   │   ├── auth.ts  # Auth mocks
│   │       │   │   ├── contracts.ts  # Contract mocks
│   │       │   │   ├── financial.ts  # Financial mocks
│   │       │   │   ├── jobs.ts  # Job data mocks
│   │       │   │   ├── proposals.ts  # Proposal mocks
│   │       │   │   └── users.ts  # User data mocks
│   │       │   ├── setup/
│   │       │   │   ├── jest.setup.ts  # Jest configuration
│   │       │   │   ├── msw.setup.ts  # MSW setup
│   │       │   │   └── testing-library.setup.ts  # Testing Library setup
│   │       │   └── utils/
│   │       │       ├── mock-api.ts  # Mock API responses
│   │       │       ├── mock-providers.tsx  # Mock context providers
│   │       │       ├── render.tsx  # Custom render with providers
│   │       │       ├── test-data-generators.ts  # Generate test data
│   │       │       └── test-helpers.ts  # Helper functions
│   │       └── utils/
│   │           └── logger/
│   │               ├── formatters/
│   │               │   ├── json.ts  # JSON formatter
│   │               │   └── pretty.ts  # Pretty formatter
│   │               ├── log-levels.ts  # Log levels
│   │               ├── logger.ts  # Logger implementation
│   │               └── transports/
│   │                   ├── console.ts  # Console transport
│   │                   ├── file.ts  # File transport
│   │                   └── remote.ts  # Remote logging service
│   └── ui/
│       └── src/
│           └── components/
│               ├── AI/
│               │   ├── AIAssistant/
│               │   │   ├── AIAssistant.native.tsx
│               │   │   ├── AIAssistant.tsx
│               │   │   ├── AIAssistant.types.ts
│               │   │   ├── AIAssistant.web.tsx
│               │   │   └── ChatInterface/
│               │   │       └── ChatInterface.tsx
│               │   ├── AutoComplete/
│               │   │   ├── SmartAutoComplete.native.tsx
│               │   │   ├── SmartAutoComplete.tsx
│               │   │   ├── SmartAutoComplete.types.ts
│               │   │   └── SmartAutoComplete.web.tsx
│               │   └── SmartSuggestions/
│               │       ├── SmartSuggestions.tsx
│               │       ├── SmartSuggestions.types.ts
│               │       └── SuggestionCard/
│               │           └── SuggestionCard.tsx
│               ├── Accessibility/
│               │   ├── FocusTrap/
│               │   │   ├── FocusTrap.tsx
│               │   │   ├── FocusTrap.types.ts
│               │   │   └── FocusTrap.web.tsx
│               │   ├── KeyboardShortcuts/
│               │   │   ├── KeyboardShortcuts.tsx
│               │   │   ├── KeyboardShortcuts.types.ts
│               │   │   └── ShortcutDialog/
│               │   │       └── ShortcutDialog.tsx
│               │   ├── LiveRegion/
│               │   │   ├── LiveRegion.native.tsx
│               │   │   ├── LiveRegion.tsx
│               │   │   ├── LiveRegion.types.ts
│               │   │   └── LiveRegion.web.tsx
│               │   └── ScreenReaderOnly/
│               │       ├── ScreenReaderOnly.tsx
│               │       └── ScreenReaderOnly.types.ts
│               ├── Charts/
│               │   ├── FunnelChart/
│               │   │   ├── FunnelChart.tsx
│               │   │   └── FunnelChart.types.ts
│               │   ├── GanttChart/
│               │   │   ├── GanttChart.tsx
│               │   │   ├── GanttChart.types.ts
│               │   │   └── Task/
│               │   │       └── Task.tsx
│               │   ├── HeatMap/
│               │   │   ├── HeatMap.tsx
│               │   │   ├── HeatMap.types.ts
│               │   │   └── HeatMap.web.tsx  # D3/Recharts
│               │   └── OrgChart/
│               │       ├── Node/
│               │       │   └── OrgNode.tsx
│               │       ├── OrgChart.tsx
│               │       ├── OrgChart.types.ts
│               │       └── OrgChart.web.tsx
│               └── Form/
│                   ├── CodeEditor/
│                   │   ├── CodeEditor.tsx
│                   │   ├── CodeEditor.types.ts
│                   │   └── LanguageSelector/
│                   │       └── LanguageSelector.tsx
│                   ├── DateRangePicker/
│                   │   ├── DateRangePicker.native.tsx
│                   │   ├── DateRangePicker.tsx
│                   │   ├── DateRangePicker.types.ts
│                   │   └── DateRangePicker.web.tsx
│                   ├── MarkdownEditor/
│                   │   ├── MarkdownEditor.tsx
│                   │   ├── MarkdownEditor.types.ts
│                   │   └── Preview/
│                   │       └── MarkdownPreview.tsx
│                   ├── RichTextEditor/
│                   │   ├── RichTextEditor.native.tsx  # Limited rich text
│                   │   ├── RichTextEditor.tsx
│                   │   ├── RichTextEditor.types.ts
│                   │   └── Toolbar/
│                   │       ├── Toolbar.tsx
│                   │       └── Toolbar.types.ts
│                   └── SignaturePad/
│                       ├── SignaturePad.native.tsx  # React Native Canvas
│                       ├── SignaturePad.tsx
│                       ├── SignaturePad.types.ts
│                       └── SignaturePad.web.tsx  # Canvas-based
├── docs/
│   ├── api/
│   │   ├── authentication.md
│   │   ├── endpoints/
│   │   │   ├── contracts.md
│   │   │   ├── jobs.md
│   │   │   ├── messages.md
│   │   │   ├── notifications.md
│   │   │   ├── payments.md
│   │   │   ├── proposals.md
│   │   │   └── users.md
│   │   ├── errors.md
│   │   ├── getting-started.md
│   │   ├── rate-limiting.md
│   │   └── webhooks.md
│   ├── architecture/
│   │   ├── backend-integration/
│   │   │   ├── api-patterns.md  # API patterns
│   │   │   ├── caching.md  # Caching strategy
│   │   │   ├── error-handling.md  # Error handling
│   │   │   └── real-time.md  # WebSocket/SSE
│   │   ├── diagrams/
│   │   │   ├── auth-flow.mmd
│   │   │   ├── data-flow.mmd
│   │   │   └── system-architecture.mmd
│   │   ├── frontend/
│   │   │   ├── authentication.md  # Auth flow
│   │   │   ├── routing.md  # Routing strategy
│   │   │   ├── state-management.md  # State management
│   │   │   └── structure.md  # Folder structure
│   │   ├── mobile/
│   │   │   ├── architecture.md  # Mobile architecture
│   │   │   ├── offline-support.md  # Offline capabilities
│   │   │   └── performance.md  # Performance optimization
│   │   └── overview.md  # System overview
│   ├── components/
│   │   ├── atoms/
│   │   │   ├── button.md
│   │   │   ├── input.md
│   │   │   └── ...
│   │   ├── design-system.md
│   │   ├── molecules/
│   │   │   ├── card.md
│   │   │   ├── form-field.md
│   │   │   └── ...
│   │   └── organisms/
│   │       ├── header.md
│   │       ├── sidebar.md
│   │       └── ...
│   └── guides/
│       ├── contributing/
│       │   ├── code-review.md
│       │   ├── getting-started.md
│       │   └── pull-requests.md
│       ├── development/
│       │   ├── coding-standards.md
│       │   ├── debugging.md
│       │   ├── git-workflow.md
│       │   └── testing.md
│       ├── deployment/
│       │   ├── ci-cd.md
│       │   ├── mobile.md
│       │   └── web.md
│       └── setup/
│           ├── environment-variables.md
│           ├── local-development.md
│           └── troubleshooting.md
├── scripts/
│   ├── analyze-bundle.ts  # Bundle size analysis
│   ├── check-bundle-limits.ts  # Enforce limits
│   └── generate-bundle-report.ts  # Generate reports
├── .bundlewatch.config.json  # Bundle size limits
├── .env.development  # Development environment
├── .env.example  # Example environment variables
├── .env.local  # Local development (git-ignored)
├── .env.production  # Production environment
└── .env.staging  # Staging environment
