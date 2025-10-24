# Additional Missing Frontend Folder Structure for Skillsier Application
## Based on fe-folder-structure-prompt.md Requirements

> **Note**: This document contains ONLY the additional missing folder structure elements that are NOT present in either `combined-folder-structure.md` or `missing-fe-folder-structure.md` but are required according to `fe-folder-structure-prompt.md`

---

## Additional Missing Dashboard Routes

### 1. Search & Discovery Routes (Enhanced)

```
apps/web/src/app/[locale]/(dashboard)/
│
├── search/
│   ├── advanced/
│   │   └── page.tsx  # Advanced search interface
│   │       # - Complex filters builder
│   │       # - Boolean operators
│   │       # - Saved search management
│   │       # BE: search-be/query
│   │       # POST /v1/search/advanced
│   │
│   ├── saved/
│   │   ├── [searchId]/
│   │   │   ├── edit/
│   │   │   │   └── page.tsx  # Edit saved search
│   │   │   │       # BE: search-be/saved-search
│   │   │   │       # PUT /v1/search/saved-searches/{search_id}
│   │   │   └── results/
│   │   │       └── page.tsx  # View results from saved search
│   │   │           # BE: search-be/saved-search, search-be/query
│   │   │           # GET /v1/search/saved-searches/{search_id}/results
│   │   └── page.tsx  # Saved searches list (may be in combined, ensuring here)
│   │
│   ├── recommendations/
│   │   └── page.tsx  # Personalized recommendations
│   │       # - AI-powered job matches
│   │       # - Talent suggestions
│   │       # BE: search-be/recommendation
│   │       # GET /v1/recommendations/personalized
│   │
│   ├── trending/
│   │   └── page.tsx  # Trending searches and jobs
│   │       # BE: search-be/trending
│   │       # GET /v1/trending/jobs
│   │       # GET /v1/trending/skills
│   │
│   └── history/
│       └── page.tsx  # Search history
│           # BE: search-be/query
│           # GET /v1/search/history
```

### 2. Connects & Credits Management (Freelancer)

```
apps/web/src/app/[locale]/(dashboard)/
│
├── connects/
│   ├── purchase/
│   │   └── page.tsx  # Purchase connects
│   │       # - Select package
│   │       # - Payment processing
│   │       # BE: proposals-be/connect, financial-be/payment
│   │       # GET /v1/connects/packages
│   │       # POST /v1/connects/purchase
│   │
│   ├── usage/
│   │   └── page.tsx  # Connects usage analytics
│   │       # - Spending patterns
│   │       # - Refund history
│   │       # - ROI tracking
│   │       # BE: proposals-be/connect
│   │       # GET /v1/connects/usage-analytics
│   │
│   └── page.tsx  # Connects dashboard
│       # - Current balance
│       # - Transaction history
│       # - Refund requests
│       # BE: proposals-be/connect
│       # GET /v1/connects
│       # GET /v1/connects/balance
```

### 3. Bidding & Strategy Management (Freelancer)

```
apps/web/src/app/[locale]/(dashboard)/
│
├── bidding/
│   ├── strategies/
│   │   ├── [strategyId]/
│   │   │   ├── edit/
│   │   │   │   └── page.tsx  # Edit bid strategy
│   │   │   │       # BE: proposals-be/bid-strategy
│   │   │   │       # PUT /v1/bid-strategies/{strategy_id}
│   │   │   └── page.tsx  # View bid strategy details
│   │   │       # BE: proposals-be/bid-strategy
│   │   │       # GET /v1/bid-strategies/{strategy_id}
│   │   │
│   │   ├── new/
│   │   │   └── page.tsx  # Create new bid strategy
│   │   │       # BE: proposals-be/bid-strategy
│   │   │       # POST /v1/bid-strategies
│   │   │
│   │   └── page.tsx  # Bid strategies list
│   │       # - Auto-bid rules
│   │       # - Price ranges
│   │       # - Category targeting
│   │       # BE: proposals-be/bid-strategy
│   │       # GET /v1/bid-strategies
│   │
│   ├── auctions/
│   │   ├── [auctionId]/
│   │   │   └── page.tsx  # Auction participation
│   │   │       # - Real-time bidding
│   │   │       # - Bid history
│   │   │       # - Competitor activity
│   │   │       # BE: proposals-be/auction
│   │   │       # GET /v1/jobs/{job_id}/auction
│   │   │       # POST /v1/jobs/{job_id}/auction/bid
│   │   │       # WebSocket: Real-time updates
│   │   │
│   │   └── page.tsx  # Active auctions list
│   │       # BE: proposals-be/auction
│   │       # GET /v1/auctions/active
│   │
│   └── analytics/
│       └── page.tsx  # Bidding analytics
│           # - Win rate
│           # - Average bid amount
│           # - Competition analysis
│           # BE: proposals-be/bid-strategy
│           # GET /v1/bid-strategies/analytics
```

### 4. Invitations Management

```
apps/web/src/app/[locale]/(dashboard)/
│
├── invitations/
│   ├── received/
│   │   ├── [inviteId]/
│   │   │   └── page.tsx  # Invitation details
│   │   │       # - Job details
│   │   │       # - Accept/decline
│   │   │       # - Proposal draft
│   │   │       # BE: proposals-be/invite, jobs-be/job
│   │   │       # GET /v1/invites/{invite_id}
│   │   │       # POST /v1/invites/{invite_id}/accept
│   │   │       # POST /v1/invites/{invite_id}/decline
│   │   │
│   │   └── page.tsx  # Received invitations list
│   │       # BE: proposals-be/invite
│   │       # GET /v1/invites/received
│   │
│   ├── sent/
│   │   ├── [inviteId]/
│   │   │   └── page.tsx  # Sent invitation tracking
│   │   │       # - Delivery status
│   │   │       # - Response tracking
│   │   │       # BE: jobs-be/invitation
│   │   │       # GET /v1/jobs/{job_id}/invitations/{invite_id}
│   │   │
│   │   └── page.tsx  # Sent invitations list (client)
│   │       # BE: jobs-be/invitation
│   │       # GET /v1/jobs/{job_id}/invitations
│   │
│   └── page.tsx  # Invitations overview
│       # - Pending actions
│       # - Response rate (client)
│       # - Conversion metrics
│       # BE: proposals-be/invite OR jobs-be/invitation (based on role)
```

### 5. Talent Sourcing & Shortlists (Client)

```
apps/web/src/app/[locale]/(dashboard)/
│
├── talent/
│   ├── browse/
│   │   └── page.tsx  # Browse talent
│   │       # - Search freelancers
│   │       # - Filters (skills, rate, location)
│   │       # - Save to shortlist
│   │       # BE: search-be/query, users-be/profile
│   │       # POST /v1/search/freelancers
│   │       # GET /v1/search/freelancers?filters=...
│   │
│   ├── shortlists/
│   │   ├── [shortlistId]/
│   │   │   ├── edit/
│   │   │   │   └── page.tsx  # Edit shortlist
│   │   │   │       # BE: jobs-be/shortlist
│   │   │   │       # PUT /v1/jobs/{job_id}/shortlists/{shortlist_id}
│   │   │   │
│   │   │   └── page.tsx  # Shortlist details
│   │   │       # - View candidates
│   │   │       # - Send invitations
│   │   │       # - Compare profiles
│   │   │       # BE: jobs-be/shortlist
│   │   │       # GET /v1/jobs/{job_id}/shortlists/{shortlist_id}
│   │   │
│   │   ├── new/
│   │   │   └── page.tsx  # Create shortlist
│   │   │       # BE: jobs-be/shortlist
│   │   │       # POST /v1/jobs/{job_id}/shortlists
│   │   │
│   │   └── page.tsx  # Shortlists overview
│   │       # BE: jobs-be/shortlist
│   │       # GET /v1/jobs/{job_id}/shortlists
│   │
│   ├── recommendations/
│   │   └── page.tsx  # AI-recommended talent for jobs
│   │       # BE: search-be/recommendation
│   │       # GET /v1/recommendations/talent?job_id={job_id}
│   │
│   └── saved/
│       └── page.tsx  # Saved talent profiles
│           # BE: users-be/profile
│           # GET /v1/users/me/saved-profiles
```

### 6. Work Tracking & Time Management

```
apps/web/src/app/[locale]/(dashboard)/
│
├── work-diary/
│   ├── [contractId]/
│   │   ├── calendar/
│   │   │   └── page.tsx  # Calendar view of work diary
│   │   │       # BE: contracts-be/work_diary
│   │   │       # GET /v1/contracts/{contract_id}/work-diary/calendar
│   │   │
│   │   ├── screenshots/
│   │   │   └── page.tsx  # Screenshots management
│   │   │       # - View all screenshots
│   │   │       # - Delete sensitive ones
│   │   │       # - Privacy settings
│   │   │       # BE: contracts-be/work_diary, storage-be/asset
│   │   │       # GET /v1/contracts/{contract_id}/work-diary/screenshots
│   │   │
│   │   └── page.tsx  # Work diary detail
│   │       # BE: contracts-be/work_diary
│   │       # GET /v1/contracts/{contract_id}/work-diary
│   │
│   └── page.tsx  # Work diary overview (all contracts)
│       # BE: contracts-be/work_diary
│       # GET /v1/work-diary
│
├── timesheets/
│   ├── [contractId]/
│   │   ├── [timesheetId]/
│   │   │   ├── edit/
│   │   │   │   └── page.tsx  # Edit timesheet
│   │   │   │       # BE: contracts-be/timesheet
│   │   │   │       # PUT /v1/contracts/{contract_id}/timesheets/{timesheet_id}
│   │   │   │
│   │   │   └── page.tsx  # Timesheet details
│   │   │       # - Hours breakdown
│   │   │       # - Approval status
│   │   │       # - Dispute options
│   │   │       # BE: contracts-be/timesheet
│   │   │       # GET /v1/contracts/{contract_id}/timesheets/{timesheet_id}
│   │   │
│   │   ├── new/
│   │   │   └── page.tsx  # Create timesheet
│   │   │       # BE: contracts-be/timesheet
│   │   │       # POST /v1/contracts/{contract_id}/timesheets
│   │   │
│   │   └── page.tsx  # Contract timesheets list
│   │       # BE: contracts-be/timesheet
│   │       # GET /v1/contracts/{contract_id}/timesheets
│   │
│   ├── approve/
│   │   └── page.tsx  # Timesheets pending approval (client)
│   │       # BE: contracts-be/timesheet
│   │       # GET /v1/timesheets/pending-approval
│   │
│   └── page.tsx  # All timesheets overview
│       # BE: contracts-be/timesheet
│       # GET /v1/timesheets
```

### 7. Deliverables Management

```
apps/web/src/app/[locale]/(dashboard)/
│
├── deliverables/
│   ├── [contractId]/
│   │   ├── [deliverableId]/
│   │   │   ├── revisions/
│   │   │   │   ├── [revisionId]/
│   │   │   │   │   └── page.tsx  # Revision detail
│   │   │   │   │       # BE: contracts-be/deliverable
│   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/revisions/{revision_id}
│   │   │   │   │
│   │   │   │   └── page.tsx  # Revision history
│   │   │   │       # BE: contracts-be/deliverable
│   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/revisions
│   │   │   │
│   │   │   ├── review/
│   │   │   │   └── page.tsx  # Review deliverable (client)
│   │   │   │       # - Approve/reject
│   │   │   │       # - Request changes
│   │   │   │       # - Add comments
│   │   │   │       # BE: contracts-be/deliverable
│   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/review
│   │   │   │
│   │   │   ├── upload/
│   │   │   │   └── page.tsx  # Upload new version
│   │   │   │       # BE: contracts-be/deliverable, storage-be/asset
│   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/upload
│   │   │   │
│   │   │   └── page.tsx  # Deliverable details
│   │   │       # - File viewer
│   │   │       # - Download
│   │   │       # - Metadata
│   │   │       # - Comments thread
│   │   │       # BE: contracts-be/deliverable, storage-be/asset
│   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}
│   │   │
│   │   ├── new/
│   │   │   └── page.tsx  # Submit new deliverable
│   │   │       # BE: contracts-be/deliverable, storage-be/asset
│   │   │       # POST /v1/contracts/{contract_id}/deliverables
│   │   │
│   │   └── page.tsx  # Contract deliverables list
│   │       # BE: contracts-be/deliverable
│   │       # GET /v1/contracts/{contract_id}/deliverables
│   │
│   ├── pending-review/
│   │   └── page.tsx  # Deliverables pending client review
│   │       # BE: contracts-be/deliverable
│   │       # GET /v1/deliverables/pending-review
│   │
│   └── page.tsx  # All deliverables overview
│       # BE: contracts-be/deliverable
│       # GET /v1/deliverables
```

### 8. Reviews & Reputation Management

```
apps/web/src/app/[locale]/(dashboard)/
│
├── reviews/
│   ├── received/
│   │   ├── [reviewId]/
│   │   │   ├── respond/
│   │   │   │   └── page.tsx  # Respond to review
│   │   │   │       # BE: reviews-be/review
│   │   │   │       # POST /v1/reviews/{review_id}/response
│   │   │   │
│   │   │   └── page.tsx  # Review details
│   │   │       # BE: reviews-be/review
│   │   │       # GET /v1/reviews/{review_id}
│   │   │
│   │   └── page.tsx  # Received reviews list
│   │       # BE: reviews-be/review
│   │       # GET /v1/users/me/reviews/received
│   │
│   ├── given/
│   │   ├── [reviewId]/
│   │   │   ├── edit/
│   │   │   │   └── page.tsx  # Edit given review
│   │   │   │       # BE: reviews-be/review
│   │   │   │       # PUT /v1/reviews/{review_id}
│   │   │   │
│   │   │   └── page.tsx  # Given review details
│   │   │       # BE: reviews-be/review
│   │   │       # GET /v1/reviews/{review_id}
│   │   │
│   │   └── page.tsx  # Given reviews list
│   │       # BE: reviews-be/review
│   │       # GET /v1/users/me/reviews/given
│   │
│   ├── pending/
│   │   ├── [contractId]/
│   │   │   └── page.tsx  # Leave review form
│   │   │       # BE: reviews-be/review, contracts-be/contract
│   │   │       # GET /v1/contracts/{contract_id}
│   │   │       # POST /v1/reviews
│   │   │
│   │   └── page.tsx  # Pending reviews to complete
│   │       # BE: reviews-be/review
│   │       # GET /v1/reviews/pending
│   │
│   ├── disputes/
│   │   ├── [disputeId]/
│   │   │   └── page.tsx  # Review dispute details
│   │   │       # - Evidence submission
│   │   │       # - Admin review status
│   │   │       # BE: reviews-be/review, admin-be/case_mgmt
│   │   │       # GET /v1/reviews/{review_id}/dispute
│   │   │
│   │   └── page.tsx  # Review disputes list
│   │       # BE: reviews-be/review
│   │       # GET /v1/reviews/disputes
│   │
│   └── analytics/
│       └── page.tsx  # Review analytics
│           # - Rating trends
│           # - Response rate
│           # - Sentiment analysis
│           # BE: reviews-be/review
│           # GET /v1/users/me/reviews/analytics
```

### 9. Networking & Connections

```
apps/web/src/app/[locale]/(dashboard)/
│
├── network/
│   ├── connections/
│   │   ├── [userId]/
│   │   │   └── page.tsx  # Connection profile view
│   │   │       # BE: users-be/profile, users-be/connection
│   │   │       # GET /v1/users/{user_id}
│   │   │       # GET /v1/users/me/connections/{user_id}
│   │   │
│   │   ├── pending/
│   │   │   └── page.tsx  # Pending connection requests
│   │   │       # BE: users-be/connection
│   │   │       # GET /v1/users/me/connections/pending
│   │   │
│   │   └── page.tsx  # Connections list
│   │       # BE: users-be/connection
│   │       # GET /v1/users/me/connections
│   │
│   ├── referrals/
│   │   ├── dashboard/
│   │   │   └── page.tsx  # Referral dashboard
│   │   │       # - Total referrals
│   │   │       # - Earnings
│   │   │       # - Conversion rate
│   │   │       # BE: users-be/referral
│   │   │       # GET /v1/users/me/referral-code
│   │   │       # GET /v1/referrals/analytics
│   │   │
│   │   └── page.tsx  # Referrals overview
│   │       # - Share referral code
│   │       # - Track referrals
│   │       # BE: users-be/referral
│   │       # GET /v1/referrals
│   │
│   ├── groups/
│   │   ├── [groupId]/
│   │   │   ├── members/
│   │   │   │   └── page.tsx  # Group members
│   │   │   │       # BE: users-be/user_group
│   │   │   │       # GET /v1/groups/{group_id}/members
│   │   │   │
│   │   │   └── page.tsx  # Group details
│   │   │       # - Posts
│   │   │       # - Events
│   │   │       # - Resources
│   │   │       # BE: users-be/user_group
│   │   │       # GET /v1/groups/{group_id}
│   │   │
│   │   ├── discover/
│   │   │   └── page.tsx  # Discover groups
│   │   │       # BE: users-be/user_group
│   │   │       # GET /v1/groups/discover
│   │   │
│   │   └── page.tsx  # My groups
│   │       # BE: users-be/user_group
│   │       # GET /v1/users/me/groups
│   │
│   └── recommendations/
│       └── page.tsx  # Connection recommendations
│           # - People you may know
│           # - Similar professionals
│           # BE: search-be/recommendation
│           # GET /v1/recommendations/connections
```

### 10. Learning & Professional Development

```
apps/web/src/app/[locale]/(dashboard)/
│
├── learning/
│   ├── paths/
│   │   ├── [pathId]/
│   │   │   ├── progress/
│   │   │   │   └── page.tsx  # Learning path progress
│   │   │   │       # BE: users-be/learning_path
│   │   │   │       # GET /v1/users/me/learning-path/{path_id}/progress
│   │   │   │
│   │   │   └── page.tsx  # Learning path details
│   │   │       # - Courses
│   │   │       # - Milestones
│   │   │       # - Resources
│   │   │       # BE: users-be/learning_path
│   │   │       # GET /v1/users/me/learning-path/{path_id}
│   │   │
│   │   ├── discover/
│   │   │   └── page.tsx  # Discover learning paths
│   │   │       # BE: users-be/learning_path
│   │   │       # GET /v1/learning-paths/discover
│   │   │
│   │   └── page.tsx  # My learning paths
│   │       # BE: users-be/learning_path
│   │       # GET /v1/users/me/learning-path
│   │
│   ├── mentorship/
│   │   ├── [sessionId]/
│   │   │   └── page.tsx  # Mentorship session details
│   │   │       # BE: users-be/mentorship
│   │   │       # GET /v1/users/me/mentorship/{session_id}
│   │   │
│   │   ├── find-mentor/
│   │   │   └── page.tsx  # Find a mentor
│   │   │       # BE: users-be/mentorship, search-be/query
│   │   │       # POST /v1/search/mentors
│   │   │
│   │   ├── my-mentees/
│   │   │   └── page.tsx  # Manage mentees
│   │   │       # BE: users-be/mentorship
│   │   │       # GET /v1/users/me/mentorship/mentees
│   │   │
│   │   └── page.tsx  # Mentorship dashboard
│   │       # BE: users-be/mentorship
│   │       # GET /v1/users/me/mentorship
│   │
│   ├── achievements/
│   │   └── page.tsx  # Achievements & badges
│   │       # - Earned badges
│   │       # - Progress to next level
│   │       # - Leaderboard
│   │       # BE: users-be/achievement
│   │       # GET /v1/users/me/achievements
│   │
│   └── certifications/
│       └── page.tsx  # Manage certifications
│           # - Upload certificates
│           # - Verification status
│           # - Expiry tracking
│           # BE: users-be/credential
│           # GET /v1/users/me/credentials?type=certification
```

### 11. Compliance & Tax Management

```
apps/web/src/app/[locale]/(dashboard)/
│
├── compliance/
│   ├── tax-profile/
│   │   ├── edit/
│   │   │   └── page.tsx  # Edit tax profile
│   │   │       # BE: users-be/compliance, financial-be/tax
│   │   │       # PUT /v1/users/me/compliance/tax-profile
│   │   │
│   │   └── page.tsx  # Tax profile overview
│   │       # - Tax ID
│   │       # - Tax forms
│   │       # - Withholding settings
│   │       # BE: users-be/compliance
│   │       # GET /v1/users/me/compliance/tax-profile
│   │
│   ├── documents/
│   │   ├── [documentId]/
│   │   │   └── page.tsx  # Compliance document details
│   │   │       # BE: storage-be/asset, admin-be/business_verification
│   │   │       # GET /v1/compliance/documents/{document_id}
│   │   │
│   │   ├── upload/
│   │   │   └── page.tsx  # Upload compliance documents
│   │   │       # BE: storage-be/asset, admin-be/business_verification
│   │   │       # POST /v1/compliance/documents/upload
│   │   │
│   │   └── page.tsx  # Compliance documents list
│   │       # BE: admin-be/business_verification
│   │       # GET /v1/compliance/documents
│   │
│   └── reports/
│       ├── tax-summary/
│       │   └── page.tsx  # Annual tax summary
│       │       # BE: financial-be/tax
│       │       # GET /v1/tax/reports/annual-summary
│       │
│       └── page.tsx  # Compliance reports
│           # - Income reports
│           # - Tax withholding
│           # - Payment history
│           # BE: financial-be/reports
│           # GET /v1/reports/compliance
```

### 12. Analytics & Insights (Enhanced)

```
apps/web/src/app/[locale]/(dashboard)/
│
├── analytics/
│   ├── performance/
│   │   └── page.tsx  # Performance analytics
│   │       # - Response time metrics
│   │       # - Proposal success rate
│   │       # - Client satisfaction
│   │       # BE: users-be/analytics, proposals-be/performance
│   │       # GET /v1/users/me/analytics/performance
│   │
│   ├── earnings/
│   │   ├── forecast/
│   │   │   └── page.tsx  # Earnings forecast
│   │   │       # - Projected income
│   │   │       # - Pipeline value
│   │   │       # BE: financial-be/analytics
│   │   │       # GET /v1/analytics/earnings/forecast
│   │   │
│   │   └── page.tsx  # Earnings analytics
│   │       # - Monthly trends
│   │       # - Year-over-year comparison
│   │       # - Client breakdown
│   │       # BE: financial-be/analytics
│   │       # GET /v1/analytics/earnings
│   │
│   ├── market-insights/
│   │   └── page.tsx  # Market insights
│   │       # - Skill demand trends
│   │       # - Rate benchmarks
│   │       # - Competition analysis
│   │       # BE: search-be/analytics, jobs-be/analytics
│   │       # GET /v1/analytics/market-insights
│   │
│   └── custom-reports/
│       ├── [reportId]/
│       │   ├── edit/
│       │   │   └── page.tsx  # Edit custom report
│       │   │       # BE: financial-be/reports
│       │   │       # PUT /v1/reports/custom/{report_id}
│       │   │
│       │   └── page.tsx  # View custom report
│       │       # BE: financial-be/reports
│       │       # GET /v1/reports/custom/{report_id}
│       │
│       ├── new/
│       │   └── page.tsx  # Create custom report
│       │       # BE: financial-be/reports
│       │       # POST /v1/reports/custom
│       │
│       └── page.tsx  # Custom reports list
│           # BE: financial-be/reports
│           # GET /v1/reports/custom
```

---

## Additional Public Routes

### 1. Legal & Compliance Pages

```
apps/web/src/app/[locale]/
│
├── legal/
│   ├── terms/
│   │   ├── freelancer/
│   │   │   └── page.tsx  # Freelancer terms of service
│   │   │       # BE: none (static content with version from CMS)
│   │   │
│   │   ├── client/
│   │   │   └── page.tsx  # Client terms of service
│   │   │       # BE: none (static content)
│   │   │
│   │   └── page.tsx  # General terms
│   │       # BE: none (static content)
│   │
│   ├── privacy/
│   │   ├── cookie-policy/
│   │   │   └── page.tsx  # Cookie policy
│   │   │       # BE: none (static content)
│   │   │
│   │   ├── data-processing/
│   │   │   └── page.tsx  # Data processing agreement
│   │   │       # BE: none (static content)
│   │   │
│   │   └── page.tsx  # Privacy policy
│   │       # BE: none (static content)
│   │
│   ├── compliance/
│   │   ├── gdpr/
│   │   │   └── page.tsx  # GDPR compliance
│   │   │       # BE: none (static content)
│   │   │
│   │   ├── ccpa/
│   │   │   └── page.tsx  # CCPA compliance
│   │   │       # BE: none (static content)
│   │   │
│   │   └── page.tsx  # Compliance overview
│   │       # BE: none (static content)
│   │
│   ├── ip-policy/
│   │   └── page.tsx  # Intellectual property policy
│   │       # BE: none (static content)
│   │
│   ├── dmca/
│   │   └── page.tsx  # DMCA policy
│   │       # BE: none (static content)
│   │
│   └── accessibility/
│       └── page.tsx  # Accessibility statement
│           # BE: none (static content)
```

### 2. Resources & Help Center

```
apps/web/src/app/[locale]/
│
├── resources/
│   ├── guides/
│   │   ├── [guideId]/
│   │   │   └── page.tsx  # Guide detail
│   │   │       # BE: CMS or static
│   │   │       # GET /v1/content/guides/{guide_id}
│   │   │
│   │   ├── freelancer/
│   │   │   └── page.tsx  # Freelancer guides
│   │   │       # BE: CMS
│   │   │       # GET /v1/content/guides?category=freelancer
│   │   │
│   │   ├── client/
│   │   │   └── page.tsx  # Client guides
│   │   │       # BE: CMS
│   │   │       # GET /v1/content/guides?category=client
│   │   │
│   │   └── page.tsx  # All guides
│   │       # BE: CMS
│   │       # GET /v1/content/guides
│   │
│   ├── tutorials/
│   │   ├── [tutorialId]/
│   │   │   └── page.tsx  # Tutorial detail
│   │   │       # BE: CMS
│   │   │       # GET /v1/content/tutorials/{tutorial_id}
│   │   │
│   │   └── page.tsx  # Tutorials list
│   │       # BE: CMS
│   │       # GET /v1/content/tutorials
│   │
│   ├── blog/
│   │   ├── [postId]/
│   │   │   └── page.tsx  # Blog post
│   │   │       # BE: CMS
│   │   │       # GET /v1/content/blog/{post_id}
│   │   │
│   │   ├── category/
│   │   │   └── [categoryId]/
│   │   │       └── page.tsx  # Blog category
│   │   │           # BE: CMS
│   │   │           # GET /v1/content/blog?category={category_id}
│   │   │
│   │   └── page.tsx  # Blog home
│   │       # BE: CMS
│   │       # GET /v1/content/blog
│   │
│   ├── case-studies/
│   │   ├── [caseStudyId]/
│   │   │   └── page.tsx  # Case study detail
│   │   │       # BE: CMS
│   │   │       # GET /v1/content/case-studies/{case_study_id}
│   │   │
│   │   └── page.tsx  # Case studies list
│   │       # BE: CMS
│   │       # GET /v1/content/case-studies
│   │
│   ├── webinars/
│   │   ├── [webinarId]/
│   │   │   └── page.tsx  # Webinar detail & registration
│   │   │       # BE: CMS + registration system
│   │   │       # GET /v1/content/webinars/{webinar_id}
│   │   │       # POST /v1/webinars/{webinar_id}/register
│   │   │
│   │   └── page.tsx  # Upcoming webinars
│   │       # BE: CMS
│   │       # GET /v1/content/webinars
│   │
│   └── faq/
│       └── page.tsx  # Frequently asked questions
│           # BE: CMS
│           # GET /v1/content/faq
```

### 3. Enterprise & Business Solutions

```
apps/web/src/app/[locale]/
│
├── enterprise/
│   ├── solutions/
│   │   ├── staffing/
│   │   │   └── page.tsx  # Enterprise staffing solutions
│   │   │       # BE: none (marketing content)
│   │   │
│   │   ├── managed-services/
│   │   │   └── page.tsx  # Managed services offering
│   │   │       # BE: none (marketing content)
│   │   │
│   │   └── page.tsx  # Enterprise solutions overview
│   │       # BE: none (marketing content)
│   │
│   ├── pricing/
│   │   └── page.tsx  # Enterprise pricing
│   │       # BE: financial-be/subscription
│   │       # GET /v1/subscriptions/plans?type=enterprise
│   │
│   ├── case-studies/
│   │   └── page.tsx  # Enterprise case studies
│   │       # BE: CMS
│   │       # GET /v1/content/case-studies?type=enterprise
│   │
│   └── contact/
│       └── page.tsx  # Enterprise contact/demo request
│           # BE: communications-be
│           # POST /v1/contact/enterprise
```

### 4. Developers Portal

```
apps/web/src/app/[locale]/
│
├── developers/
│   ├── docs/
│   │   ├── [section]/
│   │   │   └── page.tsx  # Documentation section
│   │   │       # BE: static or CMS
│   │   │       # GET /v1/content/docs/{section}
│   │   │
│   │   └── page.tsx  # API documentation home
│   │       # BE: static
│   │
│   ├── api-reference/
│   │   ├── [endpoint]/
│   │   │   └── page.tsx  # API endpoint reference
│   │   │       # BE: static (OpenAPI spec)
│   │   │
│   │   └── page.tsx  # API reference home
│   │       # BE: static (OpenAPI spec)
│   │
│   ├── sdks/
│   │   └── page.tsx  # SDK downloads and docs
│   │       # BE: static
│   │
│   ├── webhooks/
│   │   └── page.tsx  # Webhooks documentation
│   │       # BE: static
│   │
│   └── sandbox/
│       └── page.tsx  # API sandbox/playground
│           # BE: developer API
│           # POST /v1/developer/sandbox/execute
```

### 5. Platform Status & Trust

```
apps/web/src/app/[locale]/
│
├── status/
│   ├── current/
│   │   └── page.tsx  # Current system status
│   │       # BE: utility/status
│   │       # GET /v1/status/current
│   │
│   ├── history/
│   │   └── page.tsx  # Status history
│   │       # BE: utility/status
│   │       # GET /v1/status/history
│   │
│   └── subscribe/
│       └── page.tsx  # Subscribe to status updates
│           # BE: communications-be
│           # POST /v1/notifications/status-subscribe
│
├── security/
│   ├── overview/
│   │   └── page.tsx  # Security overview
│   │       # BE: none (static content)
│   │
│   ├── bug-bounty/
│   │   └── page.tsx  # Bug bounty program
│   │       # BE: none (static content)
│   │
│   ├── responsible-disclosure/
│   │   └── page.tsx  # Responsible disclosure policy
│   │       # BE: none (static content)
│   │
│   └── certifications/
│       └── page.tsx  # Security certifications (SOC2, ISO, etc.)
│           # BE: none (static content)
│
└── transparency/
    └── page.tsx  # Transparency report
        # - User statistics
        # - Moderation actions
        # - Government requests
        # BE: admin-be/reporting
        # GET /v1/public/transparency-report
```

---

## Additional Shared Features (packages/shared/src/features/)

### 1. Auction System Module

```
packages/shared/src/features/
│
├── auctions/
│   ├── api/
│   │   └── auctions-api.ts  # Auctions API client
│   │       # BE: proposals-be/auction
│   ├── hooks/
│   │   ├── useAuction.ts  # Single auction
│   │   ├── useAuctionBid.ts  # Place bid
│   │   ├── useAuctionHistory.ts  # Bid history
│   │   └── useActiveAuctions.ts  # Active auctions list
│   ├── queries/
│   │   ├── auctions-mutations.ts  # Auction mutations
│   │   └── auctions-queries.ts  # Auction queries
│   ├── store/
│   │   └── auction-store.ts  # Real-time auction state (Zustand)
│   └── types.ts  # Auction types
```

### 2. Negotiations Module

```
packages/shared/src/features/
│
├── negotiations/
│   ├── api/
│   │   └── negotiations-api.ts  # Negotiations API client
│   │       # BE: proposals-be/negotiation
│   ├── hooks/
│   │   ├── useNegotiation.ts  # Single negotiation
│   │   ├── useNegotiationOffer.ts  # Make/accept/reject offer
│   │   └── useNegotiationHistory.ts  # Negotiation history
│   ├── queries/
│   │   ├── negotiations-mutations.ts  # Negotiation mutations
│   │   └── negotiations-queries.ts  # Negotiation queries
│   └── types.ts  # Negotiation types
```

### 3. Interviews Module

```
packages/shared/src/features/
│
├── interviews/
│   ├── api/
│   │   └── interviews-api.ts  # Interviews API client
│   │       # BE: proposals-be/interview
│   ├── hooks/
│   │   ├── useInterview.ts  # Single interview
│   │   ├── useScheduleInterview.ts  # Schedule interview
│   │   ├── useInterviewFeedback.ts  # Interview feedback
│   │   └── useInterviews.ts  # Interviews list
│   ├── queries/
│   │   ├── interviews-mutations.ts  # Interview mutations
│   │   └── interviews-queries.ts  # Interview queries
│   └── types.ts  # Interview types
```

### 4. Connects Module

```
packages/shared/src/features/
│
├── connects/
│   ├── api/
│   │   └── connects-api.ts  # Connects API client
│   │       # BE: proposals-be/connect
│   ├── hooks/
│   │   ├── useConnects.ts  # Connects balance and history
│   │   ├── usePurchaseConnects.ts  # Purchase connects
│   │   ├── useConnectRefund.ts  # Request refund
│   │   └── useConnectPackages.ts  # Available packages
│   ├── queries/
│   │   ├── connects-mutations.ts  # Connect mutations
│   │   └── connects-queries.ts  # Connect queries
│   └── types.ts  # Connect types
```

### 5. Work Tracking Module

```
packages/shared/src/features/
│
├── work-tracking/
│   ├── api/
│   │   ├── work-diary-api.ts  # Work diary API
│   │   │   # BE: contracts-be/work_diary
│   │   └── timesheet-api.ts  # Timesheet API
│   │       # BE: contracts-be/timesheet
│   ├── hooks/
│   │   ├── useWorkDiary.ts  # Work diary entries
│   │   ├── useTimesheet.ts  # Timesheet management
│   │   ├── useTimeTracking.ts  # Real-time time tracking
│   │   └── useApproveTimesheet.ts  # Approve timesheet (client)
│   ├── queries/
│   │   ├── work-diary-mutations.ts  # Work diary mutations
│   │   ├── work-diary-queries.ts  # Work diary queries
│   │   ├── timesheet-mutations.ts  # Timesheet mutations
│   │   └── timesheet-queries.ts  # Timesheet queries
│   ├── store/
│   │   └── time-tracker-store.ts  # Time tracker state (Zustand)
│   └── types.ts  # Work tracking types
```

### 6. Deliverables Module

```
packages/shared/src/features/
│
├── deliverables/
│   ├── api/
│   │   └── deliverables-api.ts  # Deliverables API client
│   │       # BE: contracts-be/deliverable, storage-be/asset
│   ├── hooks/
│   │   ├── useDeliverable.ts  # Single deliverable
│   │   ├── useDeliverables.ts  # Deliverables list
│   │   ├── useUploadDeliverable.ts  # Upload deliverable
│   │   ├── useReviewDeliverable.ts  # Review deliverable (client)
│   │   └── useDeliverableRevisions.ts  # Revision management
│   ├── queries/
│   │   ├── deliverables-mutations.ts  # Deliverable mutations
│   │   └── deliverables-queries.ts  # Deliverable queries
│   └── types.ts  # Deliverable types
```

### 7. Learning Module

```
packages/shared/src/features/
│
├── learning/
│   ├── api/
│   │   ├── learning-paths-api.ts  # Learning paths API
│   │   │   # BE: users-be/learning_path
│   │   └── mentorship-api.ts  # Mentorship API
│   │       # BE: users-be/mentorship
│   ├── hooks/
│   │   ├── useLearningPath.ts  # Single learning path
│   │   ├── useLearningPaths.ts  # Learning paths list
│   │   ├── useLearningProgress.ts  # Track progress
│   │   ├── useMentorship.ts  # Mentorship management
│   │   └── useAchievements.ts  # Achievements/badges
│   ├── queries/
│   │   ├── learning-mutations.ts  # Learning mutations
│   │   └── learning-queries.ts  # Learning queries
│   └── types.ts  # Learning types
```

### 8. Networking Module

```
packages/shared/src/features/
│
├── networking/
│   ├── api/
│   │   ├── connections-api.ts  # Connections API
│   │   │   # BE: users-be/connection
│   │   ├── referrals-api.ts  # Referrals API
│   │   │   # BE: users-be/referral
│   │   └── groups-api.ts  # Groups API
│   │       # BE: users-be/user_group
│   ├── hooks/
│   │   ├── useConnections.ts  # Connections management
│   │   ├── useConnectionRequest.ts  # Send/accept/reject
│   │   ├── useReferrals.ts  # Referral management
│   │   ├── useGroups.ts  # Groups management
│   │   └── useNetworkRecommendations.ts  # Connection recommendations
│   ├── queries/
│   │   ├── networking-mutations.ts  # Networking mutations
│   │   └── networking-queries.ts  # Networking queries
│   └── types.ts  # Networking types
```

### 9. Compliance Module

```
packages/shared/src/features/
│
├── compliance/
│   ├── api/
│   │   ├── compliance-api.ts  # Compliance API
│   │   │   # BE: users-be/compliance
│   │   └── tax-profile-api.ts  # Tax profile API
│   │       # BE: users-be/compliance, financial-be/tax
│   ├── hooks/
│   │   ├── useComplianceProfile.ts  # Compliance profile
│   │   ├── useTaxProfile.ts  # Tax profile management
│   │   ├── useComplianceDocuments.ts  # Document management
│   │   └── useTaxReports.ts  # Tax reports
│   ├── queries/
│   │   ├── compliance-mutations.ts  # Compliance mutations
│   │   └── compliance-queries.ts  # Compliance queries
│   └── types.ts  # Compliance types
```

### 10. Search & Discovery Module (Enhanced)

```
packages/shared/src/features/
│
├── search/
│   ├── api/
│   │   ├── search-api.ts  # Search API (already may exist, ensuring completeness)
│   │   │   # BE: search-be/query
│   │   ├── saved-searches-api.ts  # Saved searches API
│   │   │   # BE: search-be/saved-search
│   │   ├── recommendations-api.ts  # Recommendations API
│   │   │   # BE: search-be/recommendation
│   │   └── trending-api.ts  # Trending API
│   │       # BE: search-be/trending
│   ├── hooks/
│   │   ├── useSearch.ts  # Search execution
│   │   ├── useSavedSearches.ts  # Saved searches
│   │   ├── useRecommendations.ts  # Recommendations
│   │   ├── useTrending.ts  # Trending items
│   │   ├── useSearchHistory.ts  # Search history
│   │   └── useSearchSuggestions.ts  # Auto-complete suggestions
│   ├── queries/
│   │   ├── search-mutations.ts  # Search mutations
│   │   └── search-queries.ts  # Search queries
│   ├── store/
│   │   └── search-store.ts  # Search UI state (filters, etc.)
│   └── types.ts  # Search types
```

### 11. Bidding Module

```
packages/shared/src/features/
│
├── bidding/
│   ├── api/
│   │   ├── bid-strategy-api.ts  # Bid strategy API
│   │   │   # BE: proposals-be/bid-strategy
│   │   └── bid-api.ts  # Bid placement API
│   │       # BE: proposals-be/bid
│   ├── hooks/
│   │   ├── useBidStrategy.ts  # Bid strategy management
│   │   ├── useBidStrategies.ts  # List strategies
│   │   ├── usePlaceBid.ts  # Place bid
│   │   ├── useBidAnalytics.ts  # Bid analytics
│   │   └── useBidHistory.ts  # Bid history
│   ├── queries/
│   │   ├── bidding-mutations.ts  # Bidding mutations
│   │   └── bidding-queries.ts  # Bidding queries
│   └── types.ts  # Bidding types
```

### 12. Invitations Module

```
packages/shared/src/features/
│
├── invitations/
│   ├── api/
│   │   ├── job-invitations-api.ts  # Job invitations API (client)
│   │   │   # BE: jobs-be/invitation
│   │   └── proposal-invites-api.ts  # Proposal invites API (freelancer)
│   │       # BE: proposals-be/invite
│   ├── hooks/
│   │   ├── useInvitations.ts  # Invitations management
│   │   ├── useSendInvitation.ts  # Send invitation (client)
│   │   ├── useAcceptInvite.ts  # Accept invite (freelancer)
│   │   ├── useDeclineInvite.ts  # Decline invite (freelancer)
│   │   └── useInvitationAnalytics.ts  # Invitation metrics
│   ├── queries/
│   │   ├── invitations-mutations.ts  # Invitation mutations
│   │   └── invitations-queries.ts  # Invitation queries
│   └── types.ts  # Invitation types
```

### 13. Shortlists Module

```
packages/shared/src/features/
│
├── shortlists/
│   ├── api/
│   │   └── shortlists-api.ts  # Shortlists API
│   │       # BE: jobs-be/shortlist
│   ├── hooks/
│   │   ├── useShortlist.ts  # Single shortlist
│   │   ├── useShortlists.ts  # Shortlists management
│   │   ├── useAddToShortlist.ts  # Add candidate
│   │   └── useRemoveFromShortlist.ts  # Remove candidate
│   ├── queries/
│   │   ├── shortlists-mutations.ts  # Shortlist mutations
│   │   └── shortlists-queries.ts  # Shortlist queries
│   └── types.ts  # Shortlist types
```

### 14. Feature Flags Module

```
packages/shared/src/features/
│
├── feature-flags/
│   ├── api/
│   │   └── flags-api.ts  # Feature flags API
│   │       # BE: utility/flags
│   ├── hooks/
│   │   ├── useFeatureFlag.ts  # Check single flag
│   │   ├── useFeatureFlags.ts  # Get all flags
│   │   └── useFeatureFlagVariant.ts  # A/B test variant
│   ├── queries/
│   │   └── flags-queries.ts  # Flag queries
│   ├── store/
│   │   └── flags-store.ts  # Flags state (Zustand)
│   └── types.ts  # Flag types
```

### 15. Real-time Updates Module

```
packages/shared/src/features/
│
├── realtime/
│   ├── websocket/
│   │   ├── client.ts  # WebSocket client
│   │   ├── reconnection.ts  # Reconnection logic
│   │   └── heartbeat.ts  # Connection health
│   ├── hooks/
│   │   ├── useWebSocket.ts  # WebSocket connection
│   │   ├── useRealtimeMessages.ts  # Real-time messages
│   │   ├── useRealtimeNotifications.ts  # Real-time notifications
│   │   ├── useRealtimeAuction.ts  # Real-time auction updates
│   │   └── usePresence.ts  # User presence (online/offline)
│   ├── store/
│   │   └── realtime-store.ts  # Real-time state (Zustand)
│   └── types.ts  # Real-time types
```

---

## Additional Mobile App Routes (apps/mobile/app/)

### 1. Enhanced Mobile Navigation

```
apps/mobile/app/
│
├── (tabs)/
│   ├── search/
│   │   ├── _layout.tsx  # Search stack navigator
│   │   ├── index.tsx  # Search home
│   │   │   # BE: search-be/query
│   │   ├── results.tsx  # Search results
│   │   │   # BE: search-be/query
│   │   ├── filters.tsx  # Advanced filters
│   │   │   # BE: search-be/facets
│   │   └── saved.tsx  # Saved searches
│   │       # BE: search-be/saved-search
│   │
│   ├── notifications/
│   │   ├── _layout.tsx  # Notifications stack
│   │   ├── index.tsx  # All notifications
│   │   │   # BE: communications-be/notification
│   │   ├── settings.tsx  # Notification settings
│   │   │   # BE: communications-be/preferences
│   │   └── [notificationId].tsx  # Notification detail
│   │       # BE: communications-be/notification
│   │
│   └── more/
│       ├── _layout.tsx  # More menu stack
│       ├── index.tsx  # More menu home
│       ├── account.tsx  # Account settings
│       │   # BE: users-be/account
│       ├── help.tsx  # Help center
│       │   # BE: CMS
│       └── about.tsx  # About app
│           # BE: none (static)
```

### 2. Mobile-Specific Features

```
apps/mobile/app/
│
├── offline/
│   ├── queue.tsx  # Offline actions queue
│   │   # - Pending uploads
│   │   # - Queued messages
│   │   # - Draft proposals
│   └── sync.tsx  # Sync status
│       # - Sync progress
│       # - Conflict resolution
│
├── scanner/
│   ├── document.tsx  # Document scanner
│   │   # - Scan compliance docs
│   │   # - OCR processing
│   │   # BE: storage-be/asset, admin-be/business_verification
│   └── qr-code.tsx  # QR code scanner
│       # - Event check-in
│       # - Profile sharing
│
└── widgets/
    ├── time-tracker.tsx  # Home screen time tracker widget
    │   # BE: contracts-be/work_diary
    └── quick-actions.tsx  # Quick actions widget
        # - Quick message
        # - Quick proposal
```

---

## Additional UI Components (packages/ui/src/)

### 1. Advanced Components

```
packages/ui/src/
│
├── auction/
│   ├── AuctionTimer.tsx  # Countdown timer
│   ├── AuctionTimer.web.tsx
│   ├── AuctionTimer.native.tsx
│   ├── BidHistoryChart.tsx  # Bid history visualization
│   ├── BidHistoryChart.web.tsx
│   ├── BidHistoryChart.native.tsx
│   ├── LiveBidFeed.tsx  # Real-time bid feed
│   ├── LiveBidFeed.web.tsx
│   └── LiveBidFeed.native.tsx
│
├── charts/
│   ├── EarningsChart.tsx  # Earnings visualization
│   ├── EarningsChart.web.tsx
│   ├── EarningsChart.native.tsx
│   ├── PerformanceChart.tsx  # Performance metrics
│   ├── PerformanceChart.web.tsx
│   ├── PerformanceChart.native.tsx
│   ├── TrendChart.tsx  # Trend visualization
│   ├── TrendChart.web.tsx
│   └── TrendChart.native.tsx
│
├── compliance/
│   ├── DocumentUploader.tsx  # Compliance doc uploader
│   ├── DocumentUploader.web.tsx
│   ├── DocumentUploader.native.tsx
│   ├── VerificationStatus.tsx  # Verification status badge
│   ├── VerificationStatus.web.tsx
│   └── VerificationStatus.native.tsx
│
├── tracking/
│   ├── TimeTracker.tsx  # Time tracking widget
│   ├── TimeTracker.web.tsx
│   ├── TimeTracker.native.tsx
│   ├── WorkDiaryEntry.tsx  # Work diary card
│   ├── WorkDiaryEntry.web.tsx
│   ├── WorkDiaryEntry.native.tsx
│   ├── TimesheetTable.tsx  # Timesheet grid
│   ├── TimesheetTable.web.tsx
│   └── TimesheetTable.native.tsx
│
├── collaboration/
│   ├── CollaborationPanel.tsx  # Team collaboration
│   ├── CollaborationPanel.web.tsx
│   ├── CollaborationPanel.native.tsx
│   ├── MentorCard.tsx  # Mentor profile card
│   ├── MentorCard.web.tsx
│   ├── MentorCard.native.tsx
│   ├── GroupCard.tsx  # User group card
│   ├── GroupCard.web.tsx
│   └── GroupCard.native.tsx
│
├── learning/
│   ├── LearningPathCard.tsx  # Learning path card
│   ├── LearningPathCard.web.tsx
│   ├── LearningPathCard.native.tsx
│   ├── ProgressTracker.tsx  # Progress visualization
│   ├── ProgressTracker.web.tsx
│   ├── ProgressTracker.native.tsx
│   ├── AchievementBadge.tsx  # Achievement badge
│   ├── AchievementBadge.web.tsx
│   └── AchievementBadge.native.tsx
│
└── video/
    ├── VideoPlayer.tsx  # Video player
    ├── VideoPlayer.web.tsx
    ├── VideoPlayer.native.tsx
    ├── VideoUploader.tsx  # Video upload
    ├── VideoUploader.web.tsx
    └── VideoUploader.native.tsx
```

---

## Summary of Additional Missing Elements

This document adds the following categories that were not fully covered:

### Dashboard Routes:
1. **Enhanced Search & Discovery** - Advanced search, saved searches with alerts, recommendations, trending
2. **Connects & Credits** - Purchase, usage analytics, refunds
3. **Bidding & Strategy** - Bid strategies, auctions, analytics
4. **Invitations** - Received/sent invitations, tracking
5. **Talent Sourcing** - Browse talent, shortlists, recommendations (client-side)
6. **Work Tracking** - Work diary with calendar/screenshots, timesheets with approvals
7. **Deliverables** - Upload, revisions, review process, comments
8. **Reviews** - Received/given reviews, responses, disputes, analytics
9. **Networking** - Connections, referrals, groups, recommendations
10. **Learning** - Learning paths, mentorship, achievements, certifications
11. **Compliance & Tax** - Tax profile, compliance documents, tax reports
12. **Enhanced Analytics** - Performance, earnings forecast, market insights, custom reports

### Public Routes:
1. **Legal & Compliance** - Terms, privacy, GDPR/CCPA, IP policy, DMCA, accessibility
2. **Resources** - Guides, tutorials, blog, case studies, webinars, FAQ
3. **Enterprise** - Solutions, pricing, case studies, contact
4. **Developers** - API docs, reference, SDKs, webhooks, sandbox
5. **Platform Status** - Current status, history, subscriptions
6. **Security** - Overview, bug bounty, responsible disclosure, certifications
7. **Transparency** - Public transparency reports

### Shared Features (packages/shared):
1. **Auctions** - Real-time bidding
2. **Negotiations** - Offer management
3. **Interviews** - Scheduling and feedback
4. **Connects** - Purchase and management
5. **Work Tracking** - Work diary and timesheets
6. **Deliverables** - File management and revisions
7. **Learning** - Paths, mentorship, achievements
8. **Networking** - Connections, referrals, groups
9. **Compliance** - Tax profiles and documents
10. **Enhanced Search** - Saved searches, recommendations, trending
11. **Bidding** - Bid strategies and analytics
12. **Invitations** - Job invites management
13. **Shortlists** - Candidate shortlisting
14. **Feature Flags** - A/B testing
15. **Real-time** - WebSocket connections and presence

### Mobile Enhancements:
1. **Enhanced Navigation** - Search, notifications, more menu
2. **Mobile-Specific** - Offline queue, document scanner, QR scanner, widgets

### UI Components:
1. **Auction Components** - Timer, bid history, live feed
2. **Charts** - Earnings, performance, trends
3. **Compliance Components** - Document uploader, verification status
4. **Tracking Components** - Time tracker, work diary, timesheets
5. **Collaboration** - Team panels, mentor cards, group cards
6. **Learning** - Path cards, progress trackers, badges
7. **Video** - Player and uploader

All routes and components include proper backend mappings with microservice, domain, HTTP method, and endpoint information as required.
