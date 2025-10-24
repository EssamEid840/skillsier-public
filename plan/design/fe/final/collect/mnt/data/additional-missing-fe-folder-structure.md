fe/
├── .husky/
│   ├── pre-commit
│   └── pre-push
├── .vscode/
│   ├── extensions.json
│   ├── launch.json
│   └── settings.json
├── apps/
│   ├── mobile/
│   │   ├── app/
│   │   │   ├── (auth)/
│   │   │   │   ├── _layout.tsx
│   │   │   │   ├── callback.tsx
│   │   │   │   ├── login.tsx
│   │   │   │   └── register.tsx
│   │   │   ├── (tabs)/
│   │   │   │   ├── notifications/
│   │   │   │   │   ├── _layout.tsx  # Notifications stack
│   │   │   │   │   ├── [notificationId].tsx  # Notification detail
│   │   │   │   │   │   # BE: communications-be/notification
│   │   │   │   │   ├── index.tsx  # All notifications
│   │   │   │   │   │   # BE: communications-be/notification
│   │   │   │   │   └── settings.tsx  # Notification settings
│   │   │   │   │       # BE: communications-be/preferences
│   │   │   │   ├── search/
│   │   │   │   │   ├── _layout.tsx  # Search stack navigator
│   │   │   │   │   ├── filters.tsx  # Advanced filters
│   │   │   │   │   │   # BE: search-be/facets
│   │   │   │   │   ├── index.tsx  # Search home
│   │   │   │   │   │   # BE: search-be/query
│   │   │   │   │   ├── results.tsx  # Search results
│   │   │   │   │   │   # BE: search-be/query
│   │   │   │   │   └── saved.tsx  # Saved searches
│   │   │   │   │       # BE: search-be/saved-search
│   │   │   │   └── more/
│   │   │   │       ├── _layout.tsx  # More menu stack
│   │   │   │       ├── about.tsx  # About app
│   │   │   │       │   # BE: none (static)
│   │   │   │       ├── account.tsx  # Account settings
│   │   │   │       │   # BE: users-be/account
│   │   │   │       └── help.tsx  # Help center
│   │   │   │           # BE: CMS
│   │   │   ├── +not-found.tsx
│   │   │   ├── _layout.tsx
│   │   │   └── index.tsx
│   │   ├── offline/
│   │   │   ├── queue.tsx  # Offline actions queue
│   │   │   │   # - Pending uploads
│   │   │   │   # - Queued messages
│   │   │   │   # - Draft proposals
│   │   │   └── sync.tsx  # Sync status
│   │   │       # - Sync progress
│   │   │       # - Conflict resolution
│   │   ├── scanner/
│   │   │   ├── document.tsx  # Document scanner
│   │   │   │   # - Scan compliance docs
│   │   │   │   # - OCR processing
│   │   │   │   # BE: storage-be/asset, admin-be/business_verification
│   │   │   └── qr-code.tsx  # QR code scanner
│   │   │       # - Event check-in
│   │   │       # - Profile sharing
│   │   └── widgets/
│   │       ├── quick-actions.tsx  # Quick actions widget
│   │       │   # - Quick message
│   │       │   # - Quick proposal
│   │       └── time-tracker.tsx  # Home screen time tracker widget
│   │           # BE: contracts-be/work_diary
│   └── web/
│       └── src/
│           └── app/
│               └── [locale]/
│                   ├── (dashboard)/
│                   │   ├── analytics/
│                   │   │   ├── custom-reports/
│                   │   │   │   ├── [reportId]/
│                   │   │   │   │   ├── edit/
│                   │   │   │   │   │   └── page.tsx  # Edit custom report
│                   │   │   │   │   │       # BE: financial-be/reports
│                   │   │   │   │   │       # PUT /v1/reports/custom/{report_id}
│                   │   │   │   │   └── page.tsx  # View custom report
│                   │   │   │   │       # BE: financial-be/reports
│                   │   │   │   │       # GET /v1/reports/custom/{report_id}
│                   │   │   │   ├── new/
│                   │   │   │   │   └── page.tsx  # Create custom report
│                   │   │   │   │       # BE: financial-be/reports
│                   │   │   │   │       # POST /v1/reports/custom
│                   │   │   │   └── page.tsx  # Custom reports list
│                   │   │   │       # BE: financial-be/reports
│                   │   │   │       # GET /v1/reports/custom
│                   │   │   ├── earnings/
│                   │   │   │   ├── forecast/
│                   │   │   │   │   └── page.tsx  # Earnings forecast
│                   │   │   │   │       # - Projected income
│                   │   │   │   │       # - Pipeline value
│                   │   │   │   │       # BE: financial-be/analytics
│                   │   │   │   │       # GET /v1/analytics/earnings/forecast
│                   │   │   │   └── page.tsx  # Earnings analytics
│                   │   │   │       # - Monthly trends
│                   │   │   │       # - Year-over-year comparison
│                   │   │   │       # - Client breakdown
│                   │   │   │       # BE: financial-be/analytics
│                   │   │   │       # GET /v1/analytics/earnings
│                   │   │   ├── market-insights/
│                   │   │   │   └── page.tsx  # Market insights
│                   │   │   │       # - Skill demand trends
│                   │   │   │       # - Rate benchmarks
│                   │   │   │       # - Competition analysis
│                   │   │   │       # BE: search-be/analytics, jobs-be/analytics
│                   │   │   │       # GET /v1/analytics/market-insights
│                   │   │   └── performance/
│                   │   │       └── page.tsx  # Performance analytics
│                   │   │           # - Response time metrics
│                   │   │           # - Proposal success rate
│                   │   │           # - Client satisfaction
│                   │   │           # BE: users-be/analytics, proposals-be/performance
│                   │   │           # GET /v1/users/me/analytics/performance
│                   │   ├── bidding/
│                   │   │   ├── analytics/
│                   │   │   │   └── page.tsx  # Bidding analytics
│                   │   │   │       # - Win rate
│                   │   │   │       # - Average bid amount
│                   │   │   │       # - Competition analysis
│                   │   │   │       # BE: proposals-be/bid-strategy
│                   │   │   │       # GET /v1/bid-strategies/analytics
│                   │   │   ├── auctions/
│                   │   │   │   ├── [auctionId]/
│                   │   │   │   │   └── page.tsx  # Auction participation
│                   │   │   │   │       # - Real-time bidding
│                   │   │   │   │       # - Bid history
│                   │   │   │   │       # - Competitor activity
│                   │   │   │   │       # BE: proposals-be/auction
│                   │   │   │   │       # GET /v1/jobs/{job_id}/auction
│                   │   │   │   │       # POST /v1/jobs/{job_id}/auction/bid
│                   │   │   │   │       # WebSocket: Real-time updates
│                   │   │   │   └── page.tsx  # Active auctions list
│                   │   │   │       # BE: proposals-be/auction
│                   │   │   │       # GET /v1/auctions/active
│                   │   │   └── strategies/
│                   │   │       ├── [strategyId]/
│                   │   │       │   ├── edit/
│                   │   │       │   │   └── page.tsx  # Edit bid strategy
│                   │   │       │   │       # BE: proposals-be/bid-strategy
│                   │   │       │   │       # PUT /v1/bid-strategies/{strategy_id}
│                   │   │       │   └── page.tsx  # View bid strategy details
│                   │   │       │       # BE: proposals-be/bid-strategy
│                   │   │       │       # GET /v1/bid-strategies/{strategy_id}
│                   │   │       ├── new/
│                   │   │       │   └── page.tsx  # Create new bid strategy
│                   │   │       │       # BE: proposals-be/bid-strategy
│                   │   │       │       # POST /v1/bid-strategies
│                   │   │       └── page.tsx  # Bid strategies list
│                   │   │           # - Auto-bid rules
│                   │   │           # - Price ranges
│                   │   │           # - Category targeting
│                   │   │           # BE: proposals-be/bid-strategy
│                   │   │           # GET /v1/bid-strategies
│                   │   ├── compliance/
│                   │   │   ├── documents/
│                   │   │   │   ├── [documentId]/
│                   │   │   │   │   └── page.tsx  # Compliance document details
│                   │   │   │   │       # BE: storage-be/asset, admin-be/business_verification
│                   │   │   │   │       # GET /v1/compliance/documents/{document_id}
│                   │   │   │   ├── upload/
│                   │   │   │   │   └── page.tsx  # Upload compliance documents
│                   │   │   │   │       # BE: storage-be/asset, admin-be/business_verification
│                   │   │   │   │       # POST /v1/compliance/documents/upload
│                   │   │   │   └── page.tsx  # Compliance documents list
│                   │   │   │       # BE: admin-be/business_verification
│                   │   │   │       # GET /v1/compliance/documents
│                   │   │   ├── reports/
│                   │   │   │   ├── tax-summary/
│                   │   │   │   │   └── page.tsx  # Annual tax summary
│                   │   │   │   │       # BE: financial-be/tax
│                   │   │   │   │       # GET /v1/tax/reports/annual-summary
│                   │   │   │   └── page.tsx  # Compliance reports
│                   │   │   │       # - Income reports
│                   │   │   │       # - Tax withholding
│                   │   │   │       # - Payment history
│                   │   │   │       # BE: financial-be/reports
│                   │   │   │       # GET /v1/reports/compliance
│                   │   │   └── tax-profile/
│                   │   │       ├── edit/
│                   │   │       │   └── page.tsx  # Edit tax profile
│                   │   │       │       # BE: users-be/compliance, financial-be/tax
│                   │   │       │       # PUT /v1/users/me/compliance/tax-profile
│                   │   │       └── page.tsx  # Tax profile overview
│                   │   │           # - Tax ID
│                   │   │           # - Tax forms
│                   │   │           # - Withholding settings
│                   │   │           # BE: users-be/compliance
│                   │   │           # GET /v1/users/me/compliance/tax-profile
│                   │   ├── connects/
│                   │   │   ├── purchase/
│                   │   │   │   └── page.tsx  # Purchase connects
│                   │   │   │       # - Select package
│                   │   │   │       # - Payment processing
│                   │   │   │       # BE: proposals-be/connect, financial-be/payment
│                   │   │   │       # GET /v1/connects/packages
│                   │   │   │       # POST /v1/connects/purchase
│                   │   │   ├── usage/
│                   │   │   │   └── page.tsx  # Connects usage analytics
│                   │   │   │       # - Spending patterns
│                   │   │   │       # - Refund history
│                   │   │   │       # - ROI tracking
│                   │   │   │       # BE: proposals-be/connect
│                   │   │   │       # GET /v1/connects/usage-analytics
│                   │   │   └── page.tsx  # Connects dashboard
│                   │   │       # - Current balance
│                   │   │       # - Transaction history
│                   │   │       # - Refund requests
│                   │   │       # BE: proposals-be/connect
│                   │   │       # GET /v1/connects
│                   │   │       # GET /v1/connects/balance
│                   │   ├── deliverables/
│                   │   │   ├── [contractId]/
│                   │   │   │   ├── [deliverableId]/
│                   │   │   │   │   ├── review/
│                   │   │   │   │   │   └── page.tsx  # Review deliverable (client)
│                   │   │   │   │   │       # - Approve/reject
│                   │   │   │   │   │       # - Request changes
│                   │   │   │   │   │       # - Add comments
│                   │   │   │   │   │       # BE: contracts-be/deliverable
│                   │   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/review
│                   │   │   │   │   ├── revisions/
│                   │   │   │   │   │   ├── [revisionId]/
│                   │   │   │   │   │   │   └── page.tsx  # Revision detail
│                   │   │   │   │   │   │       # BE: contracts-be/deliverable
│                   │   │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/revisions/{revision_id}
│                   │   │   │   │   │   └── page.tsx  # Revision history
│                   │   │   │   │   │       # BE: contracts-be/deliverable
│                   │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/revisions
│                   │   │   │   │   ├── upload/
│                   │   │   │   │   │   └── page.tsx  # Upload new version
│                   │   │   │   │   │       # BE: contracts-be/deliverable, storage-be/asset
│                   │   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/upload
│                   │   │   │   │   └── page.tsx  # Deliverable details
│                   │   │   │   │       # - File viewer
│                   │   │   │   │       # - Download
│                   │   │   │   │       # - Metadata
│                   │   │   │   │       # - Comments thread
│                   │   │   │   │       # BE: contracts-be/deliverable, storage-be/asset
│                   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}
│                   │   │   │   ├── new/
│                   │   │   │   │   └── page.tsx  # Submit new deliverable
│                   │   │   │   │       # BE: contracts-be/deliverable, storage-be/asset
│                   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables
│                   │   │   │   └── page.tsx  # Contract deliverables list
│                   │   │   │       # BE: contracts-be/deliverable
│                   │   │   │       # GET /v1/contracts/{contract_id}/deliverables
│                   │   │   ├── pending-review/
│                   │   │   │   └── page.tsx  # Deliverables pending client review
│                   │   │   │       # BE: contracts-be/deliverable
│                   │   │   │       # GET /v1/deliverables/pending-review
│                   │   │   └── page.tsx  # All deliverables overview
│                   │   │       # BE: contracts-be/deliverable
│                   │   │       # GET /v1/deliverables
│                   │   ├── invitations/
│                   │   │   ├── received/
│                   │   │   │   ├── [inviteId]/
│                   │   │   │   │   └── page.tsx  # Invitation details
│                   │   │   │   │       # - Job details
│                   │   │   │   │       # - Accept/decline
│                   │   │   │   │       # - Proposal draft
│                   │   │   │   │       # BE: proposals-be/invite, jobs-be/job
│                   │   │   │   │       # GET /v1/invites/{invite_id}
│                   │   │   │   │       # POST /v1/invites/{invite_id}/accept
│                   │   │   │   │       # POST /v1/invites/{invite_id}/decline
│                   │   │   │   └── page.tsx  # Received invitations list
│                   │   │   │       # BE: proposals-be/invite
│                   │   │   │       # GET /v1/invites/received
│                   │   │   ├── sent/
│                   │   │   │   ├── [inviteId]/
│                   │   │   │   │   └── page.tsx  # Sent invitation tracking
│                   │   │   │   │       # - Delivery status
│                   │   │   │   │       # - Response tracking
│                   │   │   │   │       # BE: jobs-be/invitation
│                   │   │   │   │       # GET /v1/jobs/{job_id}/invitations/{invite_id}
│                   │   │   │   └── page.tsx  # Sent invitations list (client)
│                   │   │   │       # BE: jobs-be/invitation
│                   │   │   │       # GET /v1/jobs/{job_id}/invitations
│                   │   │   └── page.tsx  # Invitations overview
│                   │   │       # - Pending actions
│                   │   │       # - Response rate (client)
│                   │   │       # - Conversion metrics
│                   │   │       # BE: proposals-be/invite OR jobs-be/invitation (based on role)
│                   │   ├── learning/
│                   │   │   ├── achievements/
│                   │   │   │   └── page.tsx  # Achievements & badges
│                   │   │   │       # - Earned badges
│                   │   │   │       # - Progress to next level
│                   │   │   │       # - Leaderboard
│                   │   │   │       # BE: users-be/achievement
│                   │   │   │       # GET /v1/users/me/achievements
│                   │   │   ├── certifications/
│                   │   │   │   └── page.tsx  # Manage certifications
│                   │   │   │       # - Upload certificates
│                   │   │   │       # - Verification status
│                   │   │   │       # - Expiry tracking
│                   │   │   │       # BE: users-be/credential
│                   │   │   │       # GET /v1/users/me/credentials?type=certification
│                   │   │   ├── mentorship/
│                   │   │   │   ├── [sessionId]/
│                   │   │   │   │   └── page.tsx  # Mentorship session details
│                   │   │   │   │       # BE: users-be/mentorship
│                   │   │   │   │       # GET /v1/users/me/mentorship/{session_id}
│                   │   │   │   ├── find-mentor/
│                   │   │   │   │   └── page.tsx  # Find a mentor
│                   │   │   │   │       # BE: users-be/mentorship, search-be/query
│                   │   │   │   │       # POST /v1/search/mentors
│                   │   │   │   ├── my-mentees/
│                   │   │   │   │   └── page.tsx  # Manage mentees
│                   │   │   │   │       # BE: users-be/mentorship
│                   │   │   │   │       # GET /v1/users/me/mentorship/mentees
│                   │   │   │   └── page.tsx  # Mentorship dashboard
│                   │   │   │       # BE: users-be/mentorship
│                   │   │   │       # GET /v1/users/me/mentorship
│                   │   │   └── paths/
│                   │   │       ├── [pathId]/
│                   │   │       │   ├── progress/
│                   │   │       │   │   └── page.tsx  # Learning path progress
│                   │   │       │   │       # BE: users-be/learning_path
│                   │   │       │   │       # GET /v1/users/me/learning-path/{path_id}/progress
│                   │   │       │   └── page.tsx  # Learning path details
│                   │   │       │       # - Courses
│                   │   │       │       # - Milestones
│                   │   │       │       # - Resources
│                   │   │       │       # BE: users-be/learning_path
│                   │   │       │       # GET /v1/users/me/learning-path/{path_id}
│                   │   │       ├── discover/
│                   │   │       │   └── page.tsx  # Discover learning paths
│                   │   │       │       # BE: users-be/learning_path
│                   │   │       │       # GET /v1/learning-paths/discover
│                   │   │       └── page.tsx  # My learning paths
│                   │   │           # BE: users-be/learning_path
│                   │   │           # GET /v1/users/me/learning-path
│                   │   ├── network/
│                   │   │   ├── connections/
│                   │   │   │   ├── [userId]/
│                   │   │   │   │   └── page.tsx  # Connection profile view
│                   │   │   │   │       # BE: users-be/profile, users-be/connection
│                   │   │   │   │       # GET /v1/users/{user_id}
│                   │   │   │   │       # GET /v1/users/me/connections/{user_id}
│                   │   │   │   ├── pending/
│                   │   │   │   │   └── page.tsx  # Pending connection requests
│                   │   │   │   │       # BE: users-be/connection
│                   │   │   │   │       # GET /v1/users/me/connections/pending
│                   │   │   │   └── page.tsx  # Connections list
│                   │   │   │       # BE: users-be/connection
│                   │   │   │       # GET /v1/users/me/connections
│                   │   │   ├── groups/
│                   │   │   │   ├── [groupId]/
│                   │   │   │   │   ├── members/
│                   │   │   │   │   │   └── page.tsx  # Group members
│                   │   │   │   │   │       # BE: users-be/user_group
│                   │   │   │   │   │       # GET /v1/groups/{group_id}/members
│                   │   │   │   │   └── page.tsx  # Group details
│                   │   │   │   │       # - Posts
│                   │   │   │   │       # - Events
│                   │   │   │   │       # - Resources
│                   │   │   │   │       # BE: users-be/user_group
│                   │   │   │   │       # GET /v1/groups/{group_id}
│                   │   │   │   ├── discover/
│                   │   │   │   │   └── page.tsx  # Discover groups
│                   │   │   │   │       # BE: users-be/user_group
│                   │   │   │   │       # GET /v1/groups/discover
│                   │   │   │   └── page.tsx  # My groups
│                   │   │   │       # BE: users-be/user_group
│                   │   │   │       # GET /v1/users/me/groups
│                   │   │   ├── recommendations/
│                   │   │   │   └── page.tsx  # Connection recommendations
│                   │   │   │       # - People you may know
│                   │   │   │       # - Similar professionals
│                   │   │   │       # BE: search-be/recommendation
│                   │   │   │       # GET /v1/recommendations/connections
│                   │   │   └── referrals/
│                   │   │       ├── dashboard/
│                   │   │       │   └── page.tsx  # Referral dashboard
│                   │   │       │       # - Total referrals
│                   │   │       │       # - Earnings
│                   │   │       │       # - Conversion rate
│                   │   │       │       # BE: users-be/referral
│                   │   │       │       # GET /v1/users/me/referral-code
│                   │   │       │       # GET /v1/referrals/analytics
│                   │   │       └── page.tsx  # Referrals overview
│                   │   │           # - Share referral code
│                   │   │           # - Track referrals
│                   │   │           # BE: users-be/referral
│                   │   │           # GET /v1/referrals
│                   │   ├── reviews/
│                   │   │   ├── analytics/
│                   │   │   │   └── page.tsx  # Review analytics
│                   │   │   │       # - Rating trends
│                   │                   # - Response rate
│                   │                   # - Sentiment analysis
│                   │                   # BE: reviews-be/review
│                   │                   # GET /v1/users/me/reviews/analytics
│                   │   │   ├── disputes/
│                   │   │   │   ├── [disputeId]/
│                   │   │   │   │   └── page.tsx  # Review dispute details
│                   │   │   │   │       # - Evidence submission
│                   │   │   │   │       # - Admin review status
│                   │   │   │   │       # BE: reviews-be/review, admin-be/case_mgmt
│                   │   │   │   │       # GET /v1/reviews/{review_id}/dispute
│                   │   │   │   └── page.tsx  # Review disputes list
│                   │   │   │       # BE: reviews-be/review
│                   │   │   │       # GET /v1/reviews/disputes
│                   │   │   ├── given/
│                   │   │   │   ├── [reviewId]/
│                   │   │   │   │   ├── edit/
│                   │   │   │   │   │   └── page.tsx  # Edit given review
│                   │   │   │   │   │       # BE: reviews-be/review
│                   │   │   │   │   │       # PUT /v1/reviews/{review_id}
│                   │   │   │   │   └── page.tsx  # Given review details
│                   │   │   │   │       # BE: reviews-be/review
│                   │   │   │   │       # GET /v1/reviews/{review_id}
│                   │   │   │   └── page.tsx  # Given reviews list
│                   │   │   │       # BE: reviews-be/review
│                   │   │   │       # GET /v1/users/me/reviews/given
│                   │   │   ├── pending/
│                   │   │   │   ├── [contractId]/
│                   │   │   │   │   └── page.tsx  # Leave review form
│                   │   │   │   │       # BE: reviews-be/review, contracts-be/contract
│                   │   │   │   │       # GET /v1/contracts/{contract_id}
│                   │   │   │   │       # POST /v1/reviews
│                   │   │   │   └── page.tsx  # Pending reviews to complete
│                   │   │   │       # BE: reviews-be/review
│                   │   │   │       # GET /v1/reviews/pending
│                   │   │   └── received/
│                   │   │       ├── [reviewId]/
│                   │   │       │   ├── respond/
│                   │   │       │   │   └── page.tsx  # Respond to review
│                   │   │       │   │       # BE: reviews-be/review
│                   │   │       │   │       # POST /v1/reviews/{review_id}/response
│                   │   │       │   └── page.tsx  # Review details
│                   │   │       │       # BE: reviews-be/review
│                   │   │       │       # GET /v1/reviews/{review_id}
│                   │   │       └── page.tsx  # Received reviews list
│                   │   │           # BE: reviews-be/review
│                   │   │           # GET /v1/users/me/reviews/received
│                   │   ├── search/
│                   │   │   ├── advanced/
│                   │   │   │   └── page.tsx  # Advanced search interface
│                   │   │   │       # - Complex filters builder
│                   │   │   │       # - Boolean operators
│                   │   │   │       # - Saved search management
│                   │   │   │       # BE: search-be/query
│                   │   │   │       # POST /v1/search/advanced
│                   │   │   ├── freelancers/
│                   │   │   │   └── page.tsx  # Advanced freelancer search (client)
│                   │   │   │       # - Search by skills
│                   │   │   ├── history/
│                   │   │   │   └── page.tsx  # Search history
│                   │   │   │       # BE: search-be/query
│                   │   │   │       # GET /v1/search/history
│                   │   │   ├── jobs/
│                   │   │   │   └── page.tsx  # Advanced job search
│                   │   │   │       # - Full-text search
│                   │   │   │       # - Faceted filters
│                   │   │   │       # - Autocomplete suggestions
│                   │   │   │       # - Search history
│                   │   │   │       # - Save search
│                   │   │   │       # BE: search-be/query
│                   │   │   │       # POST /v1/search/jobs
│                   │   │   │       # Body: { query, filters: {...}, sort, page }
│                   │   │   │       # BE: search-be/suggestions
│                   │   │   │       # GET /v1/suggestions?q={query}
│                   │   │   ├── recommendations/
│                   │   │   │   └── page.tsx  # Personalized recommendations
│                   │   │   │       # - AI-powered job matches
│                   │   │   │       # - Talent suggestions
│                   │   │   │       # BE: search-be/recommendation
│                   │   │   │       # GET /v1/recommendations/personalized
│                   │   │   ├── saved/
│                   │   │   │   ├── [searchId]/
│                   │   │   │   │   ├── edit/
│                   │   │   │   │   │   └── page.tsx  # Edit saved search
│                   │   │   │   │   │       # BE: search-be/saved-search
│                   │   │   │   │   │       # PUT /v1/search/saved-searches/{search_id}
│                   │   │   │   │   └── results/
│                   │   │   │   │       └── page.tsx  # View results from saved search
│                   │   │   │   │           # BE: search-be/saved-search, search-be/query
│                   │   │   │   │           # GET /v1/search/saved-searches/{search_id}/results
│                   │   │   │   └── page.tsx  # Saved searches list
│                   │   │   └── trending/
│                   │   │       └── page.tsx  # Trending searches and jobs
│                   │   │           # BE: search-be/trending
│                   │   │           # GET /v1/trending/jobs
│                   │   │           # GET /v1/trending/skills
│                   │   ├── talent/
│                   │   │   ├── browse/
│                   │   │   │   └── page.tsx  # Browse talent
│                   │   │   │       # - Search freelancers
│                   │   │   │       # - Filters (skills, rate, location)
│                   │   │   │       # - Save to shortlist
│                   │   │   │       # BE: search-be/query, users-be/profile
│                   │   │   │       # POST /v1/search/freelancers
│                   │   │   │       # GET /v1/search/freelancers?filters=...
│                   │   │   ├── recommendations/
│                   │   │   │   └── page.tsx  # AI-recommended talent for jobs
│                   │   │   │       # BE: search-be/recommendation
│                   │   │   │       # GET /v1/recommendations/talent?job_id={job_id}
│                   │   │   ├── saved/
│                   │   │   │   └── page.tsx  # Saved talent profiles
│                   │   │   │       # BE: users-be/profile
│                   │   │   │       # GET /v1/users/me/saved-profiles
│                   │   │   └── shortlists/
│                   │   │       ├── [shortlistId]/
│                   │   │       │   ├── edit/
│                   │   │       │   │   └── page.tsx  # Edit shortlist
│                   │   │       │   │       # BE: jobs-be/shortlist
│                   │   │       │   │       # PUT /v1/jobs/{job_id}/shortlists/{shortlist_id}
│                   │   │       │   └── page.tsx  # Shortlist details
│                   │   │       │       # - View candidates
│                   │   │       │       # - Send invitations
│                   │   │       │       # - Compare profiles
│                   │   │       │       # BE: jobs-be/shortlist
│                   │   │       │       # GET /v1/jobs/{job_id}/shortlists/{shortlist_id}
│                   │   │       ├── new/
│                   │   │       │   └── page.tsx  # Create shortlist
│                   │   │       │       # BE: jobs-be/shortlist
│                   │   │       │       # POST /v1/jobs/{job_id}/shortlists
│                   │   │       └── page.tsx  # Shortlists overview
│                   │   │           # BE: jobs-be/shortlist
│                   │   │           # GET /v1/jobs/{job_id}/shortlists
│                   │   ├── timesheets/
│                   │   │   ├── approve/
│                   │   │   │   └── page.tsx  # Timesheets pending approval (client)
│                   │   │   │       # BE: contracts-be/timesheet
│                   │   │   │       # GET /v1/timesheets/pending-approval
│                   │   │   ├── [contractId]/
│                   │   │   │   ├── [timesheetId]/
│                   │   │   │   │   ├── edit/
│                   │   │   │   │   │   └── page.tsx  # Edit timesheet
│                   │   │   │   │   │       # BE: contracts-be/timesheet
│                   │   │   │   │   │       # PUT /v1/contracts/{contract_id}/timesheets/{timesheet_id}
│                   │   │   │   │   └── page.tsx  # Timesheet details
│                   │   │   │   │       # - Hours breakdown
│                   │   │   │   │       # - Approval status
│                   │   │   │   │       # - Dispute options
│                   │   │   │   │       # BE: contracts-be/timesheet
│                   │   │   │   │       # GET /v1/contracts/{contract_id}/timesheets/{timesheet_id}
│                   │   │   │   ├── new/
│                   │   │   │   │   └── page.tsx  # Create timesheet
│                   │   │   │   │       # BE: contracts-be/timesheet
│                   │   │   │   │       # POST /v1/contracts/{contract_id}/timesheets
│                   │   │   │   └── page.tsx  # Contract timesheets list
│                   │   │   │       # BE: contracts-be/timesheet
│                   │   │   │       # GET /v1/contracts/{contract_id}/timesheets
│                   │   │   └── page.tsx  # All timesheets overview
│                   │   │       # BE: contracts-be/timesheet
│                   │   │       # GET /v1/timesheets
│                   │   └── work-diary/
│                   │       ├── [contractId]/
│                   │       │   ├── calendar/
│                   │       │   │   └── page.tsx  # Calendar view of work diary
│                   │       │   │       # BE: contracts-be/work_diary
│                   │       │   │       # GET /v1/contracts/{contract_id}/work-diary/calendar
│                   │       │   ├── screenshots/
│                   │       │   │   └── page.tsx  # Screenshots management
│                   │       │   │       # - View all screenshots
│                   │       │   │       # - Delete sensitive ones
│                   │       │   │       # - Privacy settings
│                   │       │   │       # BE: contracts-be/work_diary, storage-be/asset
│                   │       │   │       # GET /v1/contracts/{contract_id}/work-diary/screenshots
│                   │       │   └── page.tsx  # Work diary detail
│                   │       │       # BE: contracts-be/work_diary
│                   │       │       # GET /v1/contracts/{contract_id}/work-diary
│                   │       └── page.tsx  # Work diary overview (all contracts)
│                   │           # BE: contracts-be/work_diary
│                   │           # GET /v1/work-diary
│                   ├── developers/
│                   │   ├── api-reference/
│                   │   │   ├── [endpoint]/
│                   │   │   │   └── page.tsx  # API endpoint reference
│                   │   │   │       # BE: static (OpenAPI spec)
│                   │   │   └── page.tsx  # API reference home
│                   │   │       # BE: static (OpenAPI spec)
│                   │   ├── docs/
│                   │   │   ├── [section]/
│                   │   │   │   └── page.tsx  # Documentation section
│                   │   │   │       # BE: static or CMS
│                   │   │   │       # GET /v1/content/docs/{section}
│                   │   │   └── page.tsx  # API documentation home
│                   │   │       # BE: static
│                   │   ├── sandbox/
│                   │   │   └── page.tsx  # API sandbox/playground
│                   │   │       # BE: developer API
│                   │   │       # POST /v1/developer/sandbox/execute
│                   │   ├── sdks/
│                   │   │   └── page.tsx  # SDK downloads and docs
│                   │   │       # BE: static
│                   │   └── webhooks/
│                   │       └── page.tsx  # Webhooks documentation
│                   │           # BE: static
│                   ├── enterprise/
│                   │   ├── case-studies/
│                   │   │   └── page.tsx  # Enterprise case studies
│                   │   │       # BE: CMS
│                   │   │       # GET /v1/content/case-studies?type=enterprise
│                   │   ├── contact/
│                   │   │   └── page.tsx  # Enterprise contact/demo request
│                   │   │       # BE: communications-be
│                   │   │       # POST /v1/contact/enterprise
│                   │   ├── pricing/
│                   │   │   └── page.tsx  # Enterprise pricing
│                   │   │       # BE: financial-be/subscription
│                   │   │       # GET /v1/subscriptions/plans?type=enterprise
│                   │   └── solutions/
│                   │       ├── managed-services/
│                   │       │   └── page.tsx  # Managed services offering
│                   │       │       # BE: none (marketing content)
│                   │       ├── staffing/
│                   │       │   └── page.tsx  # Enterprise staffing solutions
│                   │       │       # BE: none (marketing content)
│                   │       └── page.tsx  # Enterprise solutions overview
│                   │           # BE: none (marketing content)
│                   ├── legal/
│                   │   ├── accessibility/
│                   │   │   └── page.tsx  # Accessibility statement
│                   │   │       # BE: none (static content)
│                   │   ├── compliance/
│                   │   │   ├── ccpa/
│                   │   │   │   └── page.tsx  # CCPA compliance
│                   │   │   │       # BE: none (static content)
│                   │   │   ├── gdpr/
│                   │   │   │   └── page.tsx  # GDPR compliance
│                   │   │   │       # BE: none (static content)
│                   │   │   └── page.tsx  # Compliance overview
│                   │   │       # BE: none (static content)
│                   │   ├── dmca/
│                   │   │   └── page.tsx  # DMCA policy
│                   │   │       # BE: none (static content)
│                   │   ├── ip-policy/
│                   │   │   └── page.tsx  # Intellectual property policy
│                   │   │       # BE: none (static content)
│                   │   ├── privacy/
│                   │   │   ├── cookie-policy/
│                   │   │   │   └── page.tsx  # Cookie policy
│                   │   │   │       # BE: none (static content)
│                   │   │   ├── data-processing/
│                   │   │   │   └── page.tsx  # Data processing agreement
│                   │   │   │       # BE: none (static content)
│                   │   │   └── page.tsx  # Privacy policy
│                   │   │       # BE: none (static content)
│                   │   └── terms/
│                   │       ├── client/
│                   │       │   └── page.tsx  # Client terms of service
│                   │       │       # BE: none (static content)
│                   │       ├── freelancer/
│                   │       │   └── page.tsx  # Freelancer terms of service
│                   │       │       # BE: none (static content with version from CMS)
│                   │       └── page.tsx  # General terms
│                   │           # BE: none (static content)
│                   ├── resources/
│                   │   ├── blog/
│                   │   │   ├── [postId]/
│                   │   │   │   └── page.tsx  # Blog post
│                   │   │   │       # BE: CMS
│                   │   │   │       # GET /v1/content/blog/{post_id}
│                   │   │   ├── category/
│                   │   │   │   └── [categoryId]/
│                   │   │   │       └── page.tsx  # Blog category
│                   │   │   │           # BE: CMS
│                   │   │   │           # GET /v1/content/blog?category={category_id}
│                   │   │   └── page.tsx  # Blog home
│                   │   │       # BE: CMS
│                   │   │       # GET /v1/content/blog
│                   │   ├── case-studies/
│                   │   │   ├── [caseStudyId]/
│                   │   │   │   └── page.tsx  # Case study detail
│                   │   │   │       # BE: CMS
│                   │   │   │       # GET /v1/content/case-studies/{case_study_id}
│                   │   │   └── page.tsx  # Case studies list
│                   │   │       # BE: CMS
│                   │   │       # GET /v1/content/case-studies
│                   │   ├── faq/
│                   │   │   └── page.tsx  # Frequently asked questions
│                   │   │       # BE: CMS
│                   │   │       # GET /v1/content/faq
│                   │   ├── guides/
│                   │   │   ├── [guideId]/
│                   │   │   │   └── page.tsx  # Guide detail
│                   │   │   │       # BE: CMS or static
│                   │   │   │       # GET /v1/content/guides/{guide_id}
│                   │   │   ├── client/
│                   │   │   │   └── page.tsx  # Client guides
│                   │   │   │       # BE: CMS
│                   │   │   │       # GET /v1/content/guides?category=client
│                   │   │   ├── freelancer/
│                   │   │   │   └── page.tsx  # Freelancer guides
│                   │   │   │       # BE: CMS
│                   │   │   │       # GET /v1/content/guides?category=freelancer
│                   │   │   └── page.tsx  # All guides
│                   │   │       # BE: CMS
│                   │   │       # GET /v1/content/guides
│                   │   ├── tutorials/
│                   │   │   ├── [tutorialId]/
│                   │   │   │   └── page.tsx  # Tutorial detail
│                   │   │   │       # BE: CMS
│                   │   │   │       # GET /v1/content/tutorials/{tutorial_id}
│                   │   │   └── page.tsx  # Tutorials list
│                   │   │       # BE: CMS
│                   │   │       # GET /v1/content/tutorials
│                   │   └── webinars/
│                   │       ├── [webinarId]/
│                   │       │   └── page.tsx  # Webinar detail & registration
│                   │       │       # BE: CMS + registration system
│                   │       │       # GET /v1/content/webinars/{webinar_id}
│                   │       │       # POST /v1/webinars/{webinar_id}/register
│                   │       └── page.tsx  # Upcoming webinars
│                   │           # BE: CMS
│                   │           # GET /v1/content/webinars
│                   ├── security/
│                   │   ├── bug-bounty/
│                   │   │   └── page.tsx  # Bug bounty program
│                   │   │       # BE: none (static content)
│                   │   ├── certifications/
│                   │   │   └── page.tsx  # Security certifications (SOC2, ISO, etc.)
│                   │   │       # BE: none (static content)
│                   │   ├── overview/
│                   │   │   └── page.tsx  # Security overview
│                   │   │       # BE: none (static content)
│                   │   └── responsible-disclosure/
│                   │       └── page.tsx  # Responsible disclosure policy
│                   │           # BE: none (static content)
│                   ├── status/
│                   │   ├── current/
│                   │   │   └── page.tsx  # Current system status
│                   │   │       # BE: utility/status
│                   │   │       # GET /v1/status/current
│                   │   ├── history/
│                   │   │   └── page.tsx  # Status history
│                   │   │       # BE: utility/status
│                   │   │       # GET /v1/status/history
│                   │   └── subscribe/
│                   │       └── page.tsx  # Subscribe to status updates
│                   │           # BE: communications-be
│                   │           # POST /v1/notifications/status-subscribe
│                   └── transparency/
│                       └── page.tsx  # Transparency report
│                           # - User statistics
│                           # - Moderation actions
│                           # - Government requests
│                           # BE: admin-be/reporting
│                           # GET /v1/public/transparency-report
├── packages/
│   ├── shared/
│   │   └── src/
│   │       └── features/
│   │           ├── auctions/
│   │           │   ├── api/
│   │           │   │   └── auctions-api.ts  # Auctions API client
│   │           │   │       # BE: proposals-be/auction
│   │           │   ├── hooks/
│   │           │   │   ├── useActiveAuctions.ts  # Active auctions list
│   │           │   │   ├── useAuction.ts  # Single auction
│   │           │   │   ├── useAuctionBid.ts  # Place bid
│   │           │   │   └── useAuctionHistory.ts  # Bid history
│   │           │   ├── queries/
│   │           │   │   ├── auctions-mutations.ts  # Auction mutations
│   │           │   │   └── auctions-queries.ts  # Auction queries
│   │           │   ├── store/
│   │           │   │   └── auction-store.ts  # Real-time auction state (Zustand)
│   │           │   └── types.ts  # Auction types
│   │           ├── bidding/
│   │           │   ├── api/
│   │           │   │   ├── bid-api.ts  # Bid placement API
│   │           │   │   │   # BE: proposals-be/bid
│   │           │   │   └── bid-strategy-api.ts  # Bid strategy API
│   │           │   │       # BE: proposals-be/bid-strategy
│   │           │   ├── hooks/
│   │           │   │   ├── useBidAnalytics.ts  # Bid analytics
│   │           │   │   ├── useBidHistory.ts  # Bid history
│   │           │   │   ├── useBidStrategies.ts  # List strategies
│   │           │   │   ├── useBidStrategy.ts  # Bid strategy management
│   │           │   │   └── usePlaceBid.ts  # Place bid
│   │           │   ├── queries/
│   │           │   │   ├── bidding-mutations.ts  # Bidding mutations
│   │           │   │   └── bidding-queries.ts  # Bidding queries
│   │           │   └── types.ts  # Bidding types
│   │           ├── compliance/
│   │           │   ├── api/
│   │           │   │   ├── compliance-api.ts  # Compliance API
│   │           │   │   │   # BE: users-be/compliance
│   │           │   │   └── tax-profile-api.ts  # Tax profile API
│   │           │   │       # BE: users-be/compliance, financial-be/tax
│   │           │   ├── hooks/
│   │           │   │   ├── useComplianceDocuments.ts  # Document management
│   │           │   │   ├── useComplianceProfile.ts  # Compliance profile
│   │           │   │   ├── useTaxProfile.ts  # Tax profile management
│   │           │   │   └── useTaxReports.ts  # Tax reports
│   │           │   ├── queries/
│   │           │   │   ├── compliance-mutations.ts  # Compliance mutations
│   │           │   │   └── compliance-queries.ts  # Compliance queries
│   │           │   └── types.ts  # Compliance types
│   │           ├── connects/
│   │           │   ├── api/
│   │           │   │   └── connects-api.ts  # Connects API client
│   │           │   │       # BE: proposals-be/connect
│   │           │   ├── hooks/
│   │           │   │   ├── useConnectPackages.ts  # Available packages
│   │           │   │   ├── useConnectRefund.ts  # Request refund
│   │           │   │   ├── useConnects.ts  # Connects balance and history
│   │           │   │   └── usePurchaseConnects.ts  # Purchase connects
│   │           │   ├── queries/
│   │           │   │   ├── connects-mutations.ts  # Connect mutations
│   │           │   │   └── connects-queries.ts  # Connect queries
│   │           │   └── types.ts  # Connect types
│   │           ├── deliverables/
│   │           │   ├── api/
│   │           │   │   └── deliverables-api.ts  # Deliverables API client
│   │           │   │       # BE: contracts-be/deliverable, storage-be/asset
│   │           │   ├── hooks/
│   │           │   │   ├── useDeliverable.ts  # Single deliverable
│   │           │   │   ├── useDeliverableRevisions.ts  # Revision management
│   │           │   │   ├── useDeliverables.ts  # Deliverables list
│   │           │   │   ├── useReviewDeliverable.ts  # Review deliverable (client)
│   │           │   │   └── useUploadDeliverable.ts  # Upload deliverable
│   │           │   ├── queries/
│   │           │   │   ├── deliverables-mutations.ts  # Deliverable mutations
│   │           │   │   └── deliverables-queries.ts  # Deliverable queries
│   │           │   └── types.ts  # Deliverable types
│   │           ├── invitations/
│   │           │   ├── api/
│   │           │   │   ├── job-invitations-api.ts  # Job invitations API (client)
│   │           │   │   │   # BE: jobs-be/invitation
│   │           │   │   └── proposal-invites-api.ts  # Proposal invites API (freelancer)
│   │           │   │       # BE: proposals-be/invite
│   │           │   ├── hooks/
│   │           │   │   ├── useAcceptInvite.ts  # Accept invite (freelancer)
│   │           │   │   ├── useDeclineInvite.ts  # Decline invite (freelancer)
│   │           │   │   ├── useInvitationAnalytics.ts  # Invitation metrics
│   │           │   │   ├── useInvitations.ts  # Invitations management
│   │           │   │   └── useSendInvitation.ts  # Send invitation (client)
│   │           │   ├── queries/
│   │           │   │   ├── invitations-mutations.ts  # Invitation mutations
│   │           │   │   └── invitations-queries.ts  # Invitation queries
│   │           │   └── types.ts  # Invitation types
│   │           ├── interviews/
│   │           │   ├── api/
│   │           │   │   └── interviews-api.ts  # Interviews API client
│   │           │   │       # BE: proposals-be/interview
│   │           │   ├── hooks/
│   │           │   │   ├── useInterview.ts  # Single interview
│   │           │   │   ├── useInterviewFeedback.ts  # Interview feedback
│   │           │   │   ├── useScheduleInterview.ts  # Schedule interview
│   │           │   │   └── useInterviews.ts  # Interviews list
│   │           │   ├── queries/
│   │           │   │   ├── interviews-mutations.ts  # Interview mutations
│   │           │   │   └── interviews-queries.ts  # Interview queries
│   │           │   └── types.ts  # Interview types
│   │           ├── learning/
│   │           │   ├── api/
│   │           │   │   ├── learning-paths-api.ts  # Learning paths API
│   │           │   │   │   # BE: users-be/learning_path
│   │           │   │   └── mentorship-api.ts  # Mentorship API
│   │           │   │       # BE: users-be/mentorship
│   │           │   ├── hooks/
│   │           │   │   ├── useAchievements.ts  # Achievements/badges
│   │           │   │   ├── useLearningPath.ts  # Single learning path
│   │           │   │   ├── useLearningPaths.ts  # Learning paths list
│   │           │   │   ├── useLearningProgress.ts  # Track progress
│   │           │   │   └── useMentorship.ts  # Mentorship management
│   │           │   ├── queries/
│   │           │   │   ├── learning-mutations.ts  # Learning mutations
│   │           │   │   └── learning-queries.ts  # Learning queries
│   │           │   └── types.ts  # Learning types
│   │           ├── networking/
│   │           │   ├── api/
│   │           │   │   ├── connections-api.ts  # Connections API
│   │           │   │   │   # BE: users-be/connection
│   │           │   │   ├── groups-api.ts  # Groups API
│   │           │   │   │   # BE: users-be/user_group
│   │           │   │   └── referrals-api.ts  # Referrals API
│   │           │   │       # BE: users-be/referral
│   │           │   ├── hooks/
│   │           │   │   ├── useConnectionRequest.ts  # Send/accept/reject
│   │           │   │   ├── useConnections.ts  # Connections management
│   │           │   │   ├── useGroups.ts  # Groups management
│   │           │   │   ├── useNetworkRecommendations.ts  # Connection recommendations
│   │           │   │   └── useReferrals.ts  # Referral management
│   │           │   ├── queries/
│   │           │   │   ├── networking-mutations.ts  # Networking mutations
│   │           │   │   └── networking-queries.ts  # Networking queries
│   │           │   └── types.ts  # Networking types
│   │           ├── realtime/
│   │           │   ├── hooks/
│   │           │   │   ├── usePresence.ts  # User presence (online/offline)
│   │           │   │   ├── useRealtimeAuction.ts  # Real-time auction updates
│   │           │   │   ├── useRealtimeMessages.ts  # Real-time messages
│   │           │   │   ├── useRealtimeNotifications.ts  # Real-time notifications
│   │           │   │   └── useWebSocket.ts  # WebSocket connection
│   │           │   ├── store/
│   │           │   │   └── realtime-store.ts  # Real-time state (Zustand)
│   │           │   ├── types.ts  # Real-time types
│   │           │   └── websocket/
│   │           │       ├── client.ts  # WebSocket client
│   │           │       ├── heartbeat.ts  # Connection health
│   │           │       └── reconnection.ts  # Reconnection logic
│   │           ├── search/
│   │           │   ├── api/
│   │           │   │   ├── recommendations-api.ts  # Recommendations API
│   │           │   │   │   # BE: search-be/recommendation
│   │           │   │   ├── saved-searches-api.ts  # Saved searches API
│   │           │   │   │   # BE: search-be/saved-search
│   │           │   │   ├── search-api.ts  # Search API
│   │           │   │   │   # BE: search-be/query
│   │           │   │   └── trending-api.ts  # Trending API
│   │           │   │       # BE: search-be/trending
│   │           │   ├── hooks/
│   │           │   │   ├── useRecommendations.ts  # Recommendations
│   │           │   │   ├── useSearch.ts  # Search execution
│   │           │   │   ├── useSearchHistory.ts  # Search history
│   │           │   │   ├── useSearchSuggestions.ts  # Auto-complete suggestions
│   │           │   │   └── useSavedSearches.ts  # Saved searches
│   │           │   ├── queries/
│   │           │   │   ├── search-mutations.ts  # Search mutations
│   │           │   │   └── search-queries.ts  # Search queries
│   │           │   ├── store/
│   │           │   │   └── search-store.ts  # Search UI state (filters, etc.)
│   │           │   └── types.ts  # Search types
│   │           ├── shortlists/
│   │           │   ├── api/
│   │           │   │   └── shortlists-api.ts  # Shortlists API
│   │           │   │       # BE: jobs-be/shortlist
│   │           │   ├── hooks/
│   │           │   │   ├── useAddToShortlist.ts  # Add candidate
│   │           │   │   ├── useRemoveFromShortlist.ts  # Remove candidate
│   │           │   │   ├── useShortlist.ts  # Single shortlist
│   │           │   │   └── useShortlists.ts  # Shortlists management
│   │           │   ├── queries/
│   │           │   │   ├── shortlists-mutations.ts  # Shortlist mutations
│   │           │   │   └── shortlists-queries.ts  # Shortlist queries
│   │           │   └── types.ts  # Shortlist types
│   │           ├── feature-flags/
│   │           │   ├── api/
│   │           │   │   └── flags-api.ts  # Feature flags API
│   │           │   │       # BE: utility/flags
│   │           │   ├── hooks/
│   │           │   │   ├── useFeatureFlag.ts  # Check single flag
│   │           │   │   ├── useFeatureFlags.ts  # Get all flags
│   │           │   │   └── useFeatureFlagVariant.ts  # A/B test variant
│   │           │   ├── queries/
│   │           │   │   └── flags-queries.ts  # Flag queries
│   │           │   ├── store/
│   │           │   │   └── flags-store.ts  # Flags state (Zustand)
│   │           │   └── types.ts  # Flag types
│   │           └── work-tracking/
│   │               ├── api/
│   │               │   ├── timesheet-api.ts  # Timesheet API
│   │               │   │   # BE: contracts-be/timesheet
│   │               │   └── work-diary-api.ts  # Work diary API
│   │               │       # BE: contracts-be/work_diary
│   │               ├── hooks/
│   │               │   ├── useApproveTimesheet.ts  # Approve timesheet (client)
│   │               │   ├── useTimeTracking.ts  # Real-time time tracking
│   │               │   ├── useTimesheet.ts  # Timesheet management
│   │               │   └── useWorkDiary.ts  # Work diary entries
│   │               ├── queries/
│   │               │   ├── timesheet-mutations.ts  # Timesheet mutations
│   │               │   ├── timesheet-queries.ts  # Timesheet queries
│   │               │   ├── work-diary-mutations.ts  # Work diary mutations
│   │               │   └── work-diary-queries.ts  # Work diary queries
│   │               ├── store/
│   │               │   └── time-tracker-store.ts  # Time tracker state (Zustand)
│   │               └── types.ts  # Work tracking types
│   └── ui/
│       └── src/
│           ├── auction/
│           │   ├── AuctionTimer.native.tsx
│           │   ├── AuctionTimer.tsx  # Countdown timer
│           │   ├── AuctionTimer.web.tsx
│           │   ├── BidHistoryChart.native.tsx
│           │   ├── BidHistoryChart.tsx  # Bid history visualization
│           │   ├── BidHistoryChart.web.tsx
│           │   ├── LiveBidFeed.native.tsx
│           │   ├── LiveBidFeed.tsx  # Real-time bid feed
│           │   └── LiveBidFeed.web.tsx
│           ├── charts/
│           │   ├── EarningsChart.native.tsx
│           │   ├── EarningsChart.tsx  # Earnings visualization
│           │   ├── EarningsChart.web.tsx
│           │   ├── PerformanceChart.native.tsx
│           │   ├── PerformanceChart.tsx  # Performance metrics
│           │   ├── PerformanceChart.web.tsx
│           │   ├── TrendChart.native.tsx
│           │   ├── TrendChart.tsx  # Trend visualization
│           │   └── TrendChart.web.tsx
│           ├── collaboration/
│           │   ├── CollaborationPanel.native.tsx
│           │   ├── CollaborationPanel.tsx  # Team collaboration
│           │   ├── CollaborationPanel.web.tsx
│           │   ├── GroupCard.native.tsx
│           │   ├── GroupCard.tsx  # User group card
│           │   ├── GroupCard.web.tsx
│           │   ├── MentorCard.native.tsx
│           │   ├── MentorCard.tsx  # Mentor profile card
│           │   └── MentorCard.web.tsx
│           ├── compliance/
│           │   ├── DocumentUploader.native.tsx
│           │   ├── DocumentUploader.tsx  # Compliance doc uploader
│           │   ├── DocumentUploader.web.tsx
│           │   ├── VerificationStatus.native.tsx
│           │   ├── VerificationStatus.tsx  # Verification status badge
│           │   └── VerificationStatus.web.tsx
│           ├── learning/
│           │   ├── AchievementBadge.native.tsx
│           │   ├── AchievementBadge.tsx  # Achievement badge
│           │   ├── AchievementBadge.web.tsx
│           │   ├── LearningPathCard.native.tsx
│           │   ├── LearningPathCard.tsx  # Learning path card
│           │   ├── LearningPathCard.web.tsx
│           │   ├── ProgressTracker.native.tsx
│           │   ├── ProgressTracker.tsx  # Progress visualization
│           │   └── ProgressTracker.web.tsx
│           ├── tracking/
│           │   ├── TimeTracker.native.tsx
│           │   ├── TimeTracker.tsx  # Time tracking widget
│           │   ├── TimeTracker.web.tsx
│           │   ├── TimesheetTable.native.tsx
│           │   ├── TimesheetTable.tsx  # Timesheet grid
│           │   ├── TimesheetTable.web.tsx
│           │   ├── WorkDiaryEntry.native.tsx
│           │   ├── WorkDiaryEntry.tsx  # Work diary card
│           │   └── WorkDiaryEntry.web.tsx
│           └── video/
│               ├── VideoPlayer.native.tsx
│               ├── VideoPlayer.tsx  # Video player
│               ├── VideoPlayer.web.tsx
│               ├── VideoUploader.native.tsx  # Video upload
│               ├── VideoUploader.tsx  # Video upload
│               └── VideoUploader.web.tsx
