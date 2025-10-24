```
fe/
└── fe/
    ├── (tabs)/
    │   ├── more/
    │   │   ├── _layout.tsx  # More menu stack
    │   │   ├── about.tsx  # About app
    │   │   │   # BE: none (static)
    │   │   │   # 2. Mobile-Specific Features
    │   │   ├── account.tsx  # Account settings
    │   │   │   # BE: users-be/account
    │   │   ├── help.tsx  # Help center
    │   │   │   # BE: CMS
    │   │   └── index.tsx  # More menu home
    │   ├── notifications/
    │   │   ├── [notificationId].tsx  # Notification detail
    │   │   │   # BE: communications-be/notification
    │   │   ├── _layout.tsx  # Notifications stack
    │   │   ├── index.tsx  # All notifications
    │   │   │   # BE: communications-be/notification
    │   │   └── settings.tsx  # Notification settings
    │   │       # BE: communications-be/preferences
    │   ├── search/
    │   │   ├── _layout.tsx  # Search stack navigator
    │   │   ├── filters.tsx  # Advanced filters
    │   │   │   # BE: search-be/facets
    │   │   ├── index.tsx  # Search home
    │   │   │   # BE: search-be/query
    │   │   ├── results.tsx  # Search results
    │   │   │   # BE: search-be/query
    │   │   └── saved.tsx  # Saved searches
    │   │       # BE: search-be/saved-search
    │   └── │
    ├── .github/  # GitHub workflows
    │   ├── actions/
    │   │   ├── build-mobile/
    │   │   │   └── action.yml
    │   │   ├── build-web/
    │   │   │   └── action.yml
    │   │   ├── cache-deps/
    │   │   │   └── action.yml
    │   │   ├── deploy-preview/
    │   │   │   └── action.yml
    │   │   └── setup-node/
    │   │       └── action.yml
    │   ├── workflows/
    │   │   ├── accessibility.yml  # Accessibility checks
    │   │   │   # - Run axe-core
    │   │   │   # - Check WCAG compliance
    │   │   ├── cd-mobile-production.yml  # Mobile production deployment
    │   │   │   # - Build
    │   │   │   # - Submit to App Store/Play Store
    │   │   ├── cd-mobile-staging.yml  # Mobile staging deployment
    │   │   │   # - Build
    │   │   │   # - Submit to TestFlight/Internal Testing
    │   │   ├── cd-mobile.yml  # Mobile deployment
    │   │   │   # - EAS build
    │   │   │   # - Submit to stores
    │   │   ├── cd-web-production.yml  # Web production deployment
    │   │   │   # - Build
    │   │   │   # - Deploy to production
    │   │   │   # - Run smoke tests
    │   │   │   # - Notify team
    │   │   ├── cd-web-staging.yml  # Web staging deployment
    │   │   │   # - Build
    │   │   │   # - Deploy to staging
    │   │   │   # - Run smoke tests
    │   │   ├── cd-web.yml  # Web deployment
    │   │   │   # - Build Next.js production
    │   │   │   # - Deploy to K8s
    │   │   ├── ci-mobile.yml  # Mobile CI pipeline
    │   │   │   # - Lint
    │   │   │   # - Type check
    │   │   │   # - Unit tests
    │   │   │   # - Build (iOS/Android)
    │   │   ├── ci-web.yml  # Web CI pipeline
    │   │   │   # - Lint
    │   │   │   # - Type check
    │   │   │   # - Unit tests
    │   │   │   # - Build
    │   │   │   # - Bundle size check
    │   │   ├── ci.yml  # Continuous Integration
    │   │   │   # - Lint, type-check, test
    │   │   │   # - Build verification
    │   │   │   # - Bundle size check
    │   │   ├── dependabot.yml  # Automated dependency updates
    │   │   ├── dependency-review.yml  # Dependency review
    │   │   │   # - Check for vulnerabilities
    │   │   │   # - License compliance
    │   │   ├── e2e-tests.yml  # E2E tests
    │   │   │   # - Setup test environment
    │   │   │   # - Run Playwright/Detox tests
    │   │   │   # - Upload test results
    │   │   ├── lighthouse.yml  # Performance audits
    │   │   │   # - Run Lighthouse CI
    │   │   │   # - Compare against budgets
    │   │   │   # - Comment on PR
    │   │   ├── release.yml  # Release automation
    │   │   │   # - Create changelog
    │   │   │   # - Tag release
    │   │   │   # - Create GitHub release
    │   │   └── security.yml  # Security scanning
    │   │       # - Dependency audit
    │   │       # - SAST scan
    │   │       # - License check
    │   └── CODEOWNERS  # Code ownership
    ├── .husky/  # Git hooks for quality gates
    │   ├── pre-commit  # Runs linting, type checking
    │   │   # Validates no console.log in prod code
    │   │   # Ensures all tests pass
    │   └── pre-push  # Runs full test suite before push
    │       # Bundle size check
    ├── .vscode/  # VS Code workspace configuration
    │   ├── extensions.json  # Recommended extensions list
    │   │   # - ESLint, Prettier, Tailwind IntelliSense
    │   │   # - i18n Ally, Error Lens
    │   ├── launch.json  # Debug configurations
    │   │   # - Next.js debug config
    │   │   # - Expo debug config
    │   │   # - Jest test debug config
    │   └── settings.json  # Workspace settings
    │       # - Format on save enabled
    │       # - Auto import organization
    │       # - Tailwind CSS IntelliSense config
    ├── analytics/
    │   ├── custom-reports/
    │   │   ├── [reportId]/
    │   │   │   ├── edit/
    │   │   │   │   └── page.tsx  # Edit custom report
    │   │   │   │       # BE: financial-be/reports
    │   │   │   │       # PUT /v1/reports/custom/{report_id}
    │   │   │   ├── page.tsx  # View custom report
    │   │   │   │   # BE: financial-be/reports
    │   │   │   │   # GET /v1/reports/custom/{report_id}
    │   │   │   └── │
    │   │   ├── new/
    │   │   │   └── page.tsx  # Create custom report
    │   │   │       # BE: financial-be/reports
    │   │   │       # POST /v1/reports/custom
    │   │   ├── page.tsx  # Custom reports list
    │   │   │   # BE: financial-be/reports
    │   │   │   # GET /v1/reports/custom
    │   │   └── │
    │   ├── earnings/
    │   │   ├── forecast/
    │   │   │   └── page.tsx  # Earnings forecast
    │   │   │       # - Projected income
    │   │   │       # - Pipeline value
    │   │   │       # BE: financial-be/analytics
    │   │   │       # GET /v1/analytics/earnings/forecast
    │   │   ├── page.tsx  # Earnings analytics
    │   │   │   # - Monthly trends
    │   │   │   # - Year-over-year comparison
    │   │   │   # - Client breakdown
    │   │   │   # BE: financial-be/analytics
    │   │   │   # GET /v1/analytics/earnings
    │   │   └── │
    │   ├── market-insights/
    │   │   └── page.tsx  # Market insights
    │   │       # - Skill demand trends
    │   │       # - Rate benchmarks
    │   │       # - Competition analysis
    │   │       # BE: search-be/analytics, jobs-be/analytics
    │   │       # GET /v1/analytics/market-insights
    │   ├── performance/
    │   │   └── page.tsx  # Performance analytics
    │   │       # - Response time metrics
    │   │       # - Proposal success rate
    │   │       # - Client satisfaction
    │   │       # BE: users-be/analytics, proposals-be/performance
    │   │       # GET /v1/users/me/analytics/performance
    │   └── │
    ├── apps/  # Application workspaces
    │   ├── mobile/  # React Native/Expo application
    │   │   ├── app/  # Expo Router file-based routing
    │   │   │   ├── (auth)/  # Auth screens
    │   │   │   │   ├── _layout.tsx  # Auth layout
    │   │   │   │   ├── callback.tsx  # OAuth callback
    │   │   │   │   │   # BE: Keycloak token exchange
    │   │   │   │   ├── forgot-password.tsx  # Password reset
    │   │   │   │   │   # BE: users-be/security/recovery
    │   │   │   │   │   # POST /v1/auth/forgot-password
    │   │   │   │   ├── login.tsx  # Login screen
    │   │   │   │   │   # - Email/password form
    │   │   │   │   │   # - Social login (Google, Apple)
    │   │   │   │   │   # - Biometric login (Face ID, Touch ID)
    │   │   │   │   │   # BE: Keycloak OAuth2
    │   │   │   │   │   # POST /v1/auth/login
    │   │   │   │   └── register.tsx  # Registration screen
    │   │   │   │       # BE: users-be/user
    │   │   │   │       # POST /v1/users/register
    │   │   │   ├── (onboarding)/  # Onboarding flow
    │   │   │   │   ├── _layout.tsx  # Onboarding layout
    │   │   │   │   ├── complete.tsx  # Onboarding complete
    │   │   │   │   ├── preferences.tsx  # Preferences
    │   │   │   │   ├── profile.tsx  # Basic profile setup
    │   │   │   │   ├── skills.tsx  # Skills selection
    │   │   │   │   └── welcome.tsx  # Welcome screen
    │   │   │   ├── (tabs)/  # Bottom tabs navigation
    │   │   │   │   ├── more/
    │   │   │   │   │   ├── _layout.tsx  # More menu stack
    │   │   │   │   │   ├── about.tsx  # About app
    │   │   │   │   │   │   # BE: none (static)
    │   │   │   │   │   ├── account.tsx  # Account settings
    │   │   │   │   │   │   # BE: users-be/account
    │   │   │   │   │   ├── help.tsx  # Help center
    │   │   │   │   │   │   # BE: CMS
    │   │   │   │   │   └── index.tsx  # More menu home
    │   │   │   │   ├── notifications/
    │   │   │   │   │   ├── [notificationId].tsx  # Notification detail
    │   │   │   │   │   │   # BE: communications-be/notification
    │   │   │   │   │   ├── _layout.tsx  # Notifications stack
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
    │   │   │   │   ├── _layout.tsx  # Tabs layout
    │   │   │   │   │   # - Bottom tab navigator
    │   │   │   │   │   # - Tab bar icons
    │   │   │   │   │   # - Badge indicators (messages, notifications)
    │   │   │   │   ├── index.tsx  # Home tab / Dashboard
    │   │   │   │   │   # - Dashboard overview
    │   │   │   │   │   # - Quick actions
    │   │   │   │   │   # - Recent activity
    │   │   │   │   │   # BE: Multiple services (same as web dashboard)
    │   │   │   │   ├── jobs.tsx  # Jobs tab
    │   │   │   │   │   # - Job listings (browse/my-jobs based on role)
    │   │   │   │   │   # - Search jobs
    │   │   │   │   │   # - Saved jobs
    │   │   │   │   │   # BE: jobs-be/job
    │   │   │   │   │   # GET /v1/jobs/browse (freelancer)
    │   │   │   │   │   # GET /v1/jobs/my-jobs (client)
    │   │   │   │   ├── messages.tsx  # Messages tab
    │   │   │   │   │   # - Conversation list
    │   │   │   │   │   # - Real-time updates
    │   │   │   │   │   # BE: communications-be/conversations
    │   │   │   │   │   # GET /v1/conversations
    │   │   │   │   │   # WebSocket connection
    │   │   │   │   ├── profile.tsx  # Profile tab
    │   │   │   │   │   # - Current user profile preview
    │   │   │   │   │   # - Quick links to settings
    │   │   │   │   │   # - Stats
    │   │   │   │   │   # BE: users-be/profile
    │   │   │   │   │   # GET /v1/users/me
    │   │   │   │   └── proposals.tsx  # Proposals tab
    │   │   │   │       # - My proposals (freelancer)
    │   │   │   │       # - Received proposals (client, redirects to job)
    │   │   │   │       # BE: proposals-be
    │   │   │   │       # GET /v1/proposals/my-proposals
    │   │   │   ├── contracts/
    │   │   │   │   ├── [id]/
    │   │   │   │   │   ├── index.tsx  # Contract overview
    │   │   │   │   │   │   # BE: contracts-be/contract
    │   │   │   │   │   │   # GET /v1/contracts/{contract_id}
    │   │   │   │   │   ├── messages.tsx  # Contract messages
    │   │   │   │   │   │   # BE: communications-be
    │   │   │   │   │   │   # GET /v1/contracts/{contract_id}/conversation
    │   │   │   │   │   ├── milestones.tsx  # Milestones
    │   │   │   │   │   │   # BE: contracts-be/milestone
    │   │   │   │   │   │   # GET /v1/contracts/{contract_id}/milestones
    │   │   │   │   │   ├── timesheet.tsx  # Timesheet
    │   │   │   │   │   │   # BE: contracts-be/timesheet
    │   │   │   │   │   │   # GET /v1/contracts/{contract_id}/timesheets
    │   │   │   │   │   └── work-diary.tsx  # Work diary
    │   │   │   │   │       # BE: contracts-be/work_diary
    │   │   │   │   │       # GET /v1/contracts/{contract_id}/work-diary
    │   │   │   │   └── index.tsx  # Contracts list
    │   │   │   │       # BE: contracts-be/contract
    │   │   │   │       # GET /v1/contracts
    │   │   │   ├── financials/
    │   │   │   │   ├── invoices.tsx  # Invoices
    │   │   │   │   │   # BE: financial-be/invoice
    │   │   │   │   │   # GET /v1/invoices
    │   │   │   │   ├── transactions.tsx  # Transaction history
    │   │   │   │   │   # BE: financial-be/transaction
    │   │   │   │   │   # GET /v1/transactions
    │   │   │   │   └── wallet.tsx  # Wallet
    │   │   │   │       # BE: financial-be/wallet
    │   │   │   │       # GET /v1/wallet/balance
    │   │   │   ├── jobs/
    │   │   │   │   ├── [id].tsx  # Job detail
    │   │   │   │   │   # BE: jobs-be/job
    │   │   │   │   │   # GET /v1/jobs/{job_id}
    │   │   │   │   ├── post.tsx  # Post job (client)
    │   │   │   │   │   # BE: jobs-be/job
    │   │   │   │   │   # POST /v1/jobs
    │   │   │   │   └── search.tsx  # Job search
    │   │   │   │       # BE: search-be
    │   │   │   │       # POST /v1/search/jobs
    │   │   │   ├── messages/
    │   │   │   │   └── [conversationId].tsx  # Conversation thread
    │   │   │   │       # - Message list
    │   │   │   │       # - Message composer
    │   │   │   │       # - Real-time updates
    │   │   │   │       # BE: communications-be/messages
    │   │   │   │       # GET /v1/conversations/{conversation_id}/messages
    │   │   │   │       # POST /v1/messages
    │   │   │   │       # WebSocket updates
    │   │   │   ├── notifications/
    │   │   │   │   └── index.tsx  # Notifications list
    │   │   │   │       # BE: communications-be/notifications
    │   │   │   │       # GET /v1/notifications
    │   │   │   ├── offline/
    │   │   │   │   ├── queue.tsx  # Offline actions queue
    │   │   │   │   │   # - Pending uploads
    │   │   │   │   │   # - Queued messages
    │   │   │   │   │   # - Draft proposals
    │   │   │   │   └── sync.tsx  # Sync status
    │   │   │   │       # - Sync progress
    │   │   │   │       # - Conflict resolution
    │   │   │   ├── profile/
    │   │   │   │   ├── edit/
    │   │   │   │   │   ├── experience.tsx  # Edit experience
    │   │   │   │   │   ├── index.tsx  # Edit profile
    │   │   │   │   │   ├── portfolio.tsx  # Edit portfolio
    │   │   │   │   │   └── skills.tsx  # Edit skills
    │   │   │   │   └── [userId].tsx  # User profile (public view)
    │   │   │   │       # BE: users-be/profile
    │   │   │   │       # GET /v1/users/{user_id}/profile
    │   │   │   ├── proposals/
    │   │   │   │   ├── submit/
    │   │   │   │   │   └── [jobId].tsx  # Submit proposal
    │   │   │   │   │       # BE: proposals-be
    │   │   │   │   │       # POST /v1/proposals
    │   │   │   │   └── [id].tsx  # Proposal detail
    │   │   │   │       # BE: proposals-be
    │   │   │   │       # GET /v1/proposals/{proposal_id}
    │   │   │   ├── reviews/
    │   │   │   │   ├── create/
    │   │   │   │   │   └── [contractId].tsx  # Create review
    │   │   │   │   │       # BE: reviews-be/reviews
    │   │   │   │   │       # POST /v1/reviews
    │   │   │   │   └── index.tsx  # Reviews list
    │   │   │   │       # BE: reviews-be/reviews
    │   │   │   │       # GET /v1/reviews
    │   │   │   ├── scanner/
    │   │   │   │   ├── document.tsx  # Document scanner
    │   │   │   │   │   # - Scan compliance docs
    │   │   │   │   │   # - OCR processing
    │   │   │   │   │   # BE: storage-be/asset, admin-be/business_verification
    │   │   │   │   └── qr-code.tsx  # QR code scanner
    │   │   │   │       # - Event check-in
    │   │   │   │       # - Profile sharing
    │   │   │   ├── settings/
    │   │   │   │   ├── about.tsx  # About & support
    │   │   │   │   ├── account.tsx  # Account settings
    │   │   │   │   ├── index.tsx  # Settings menu
    │   │   │   │   ├── notifications.tsx  # Notification settings
    │   │   │   │   ├── privacy.tsx  # Privacy settings
    │   │   │   │   └── security.tsx  # Security settings
    │   │   │   ├── subscription/
    │   │   │   │   ├── connects.tsx  # Connects management
    │   │   │   │   │   # BE: subscriptions-be
    │   │   │   │   ├── index.tsx  # Subscription overview
    │   │   │   │   ├── plans.tsx  # Available plans
    │   │   │   │   └── upgrade.tsx  # Upgrade plan
    │   │   │   ├── widgets/
    │   │   │   │   ├── quick-actions.tsx  # Quick actions widget
    │   │   │   │   │   # - Quick message
    │   │   │   │   │   # - Quick proposal
    │   │   │   │   └── time-tracker.tsx  # Home screen time tracker widget
    │   │   │   │       # BE: contracts-be/work_diary
    │   │   │   ├── +not-found.tsx  # 404 screen
    │   │   │   ├── _layout.tsx  # Root layout
    │   │   │   │   # - Auth provider
    │   │   │   │   # - Query client provider
    │   │   │   │   # - Theme provider
    │   │   │   │   # - Error boundary
    │   │   │   └── index.tsx  # App entry point
    │   │   │       # - Splash screen
    │   │   │       # - Initial route determination
    │   │   ├── assets/  # Mobile assets
    │   │   │   ├── fonts/  # Custom fonts
    │   │   │   ├── icons/  # App icons
    │   │   │   ├── images/  # Images
    │   │   │   └── splash/  # Splash screens
    │   │   ├── src/
    │   │   │   ├── components/  # Mobile-specific components
    │   │   │   │   ├── Auth/
    │   │   │   │   │   ├── BiometricButton.tsx  # Biometric auth button
    │   │   │   │   │   ├── LoginForm.tsx  # Login form component
    │   │   │   │   │   ├── RegisterForm.tsx  # Registration form
    │   │   │   │   │   └── SocialButtons.tsx  # Social login buttons
    │   │   │   │   ├── Common/
    │   │   │   │   │   ├── EmptyState.tsx  # Empty state component
    │   │   │   │   │   ├── ErrorBoundary.tsx  # Error boundary
    │   │   │   │   │   ├── Loading.tsx  # Loading spinner
    │   │   │   │   │   ├── OptimizedFlashList.tsx  # Optimized list (FlashList)
    │   │   │   │   │   └── PullToRefresh.tsx  # Pull to refresh
    │   │   │   │   ├── Contracts/
    │   │   │   │   │   ├── ContractCard.tsx  # Contract card
    │   │   │   │   │   ├── MilestoneItem.tsx  # Milestone list item
    │   │   │   │   │   └── TimesheetEntry.tsx  # Timesheet entry
    │   │   │   │   ├── Financial/
    │   │   │   │   │   ├── InvoiceCard.tsx  # Invoice card
    │   │   │   │   │   ├── TransactionItem.tsx  # Transaction list item
    │   │   │   │   │   └── WalletCard.tsx  # Wallet balance card
    │   │   │   │   ├── Jobs/
    │   │   │   │   │   ├── JobCard.tsx  # Job card component
    │   │   │   │   │   ├── JobDetail.tsx  # Job detail view
    │   │   │   │   │   ├── JobFilters.tsx  # Job filters bottom sheet
    │   │   │   │   │   └── JobList.tsx  # Job list
    │   │   │   │   ├── Messages/
    │   │   │   │   │   ├── ConversationCard.tsx  # Conversation list item
    │   │   │   │   │   ├── MessageBubble.tsx  # Message bubble
    │   │   │   │   │   ├── MessageComposer.tsx  # Message input
    │   │   │   │   │   └── TypingIndicator.tsx  # Typing indicator
    │   │   │   │   ├── Navigation/
    │   │   │   │   │   ├── Header.tsx  # Screen header
    │   │   │   │   │   └── TabBar.tsx  # Custom tab bar
    │   │   │   │   ├── Profile/
    │   │   │   │   │   ├── ExperienceItem.tsx  # Experience item
    │   │   │   │   │   ├── PortfolioItem.tsx  # Portfolio item
    │   │   │   │   │   ├── ProfileHeader.tsx  # Profile header
    │   │   │   │   │   └── SkillTag.tsx  # Skill tag
    │   │   │   │   ├── Proposals/
    │   │   │   │   │   ├── ProposalCard.tsx  # Proposal card
    │   │   │   │   │   ├── ProposalForm.tsx  # Proposal submission form
    │   │   │   │   │   └── ProposalList.tsx  # Proposal list
    │   │   │   │   └── UI/
    │   │   │   │       ├── Avatar.tsx  # Avatar component
    │   │   │   │       ├── Badge.tsx  # Badge component
    │   │   │   │       ├── BottomSheet.tsx  # Bottom sheet modal
    │   │   │   │       ├── Button.tsx  # Button component
    │   │   │   │       ├── Card.tsx  # Card component
    │   │   │   │       ├── Input.tsx  # Input component
    │   │   │   │       └── SearchBar.tsx  # Search bar
    │   │   │   ├── hooks/  # Mobile-specific hooks
    │   │   │   │   ├── useAppState.ts  # App state (foreground/background)
    │   │   │   │   ├── useBiometricAuth.ts  # Biometric authentication
    │   │   │   │   ├── useCamera.ts  # Camera access
    │   │   │   │   ├── useHighFPSAnimation.ts  # High FPS animations
    │   │   │   │   ├── useKeyboard.ts  # Keyboard handling
    │   │   │   │   ├── useLocation.ts  # Geolocation
    │   │   │   │   ├── useNetworkStatus.ts  # Network connectivity
    │   │   │   │   ├── useOrientation.ts  # Device orientation
    │   │   │   │   └── usePushNotifications.ts  # Push notifications
    │   │   │   ├── lib/
    │   │   │   │   ├── analytics.ts  # Mobile analytics
    │   │   │   │   │   # - Event tracking
    │   │   │   │   │   # - Screen tracking
    │   │   │   │   ├── deeplink.ts  # Deep link handling
    │   │   │   │   ├── error-tracking.ts  # Error tracking (Sentry)
    │   │   │   │   ├── keycloak-mobile.ts  # Keycloak mobile config
    │   │   │   │   │   # - OAuth2 with PKCE
    │   │   │   │   │   # - Token storage (SecureStore)
    │   │   │   │   │   # - Refresh token flow
    │   │   │   │   ├── performance.ts  # Performance utilities
    │   │   │   │   │   # - FPS monitoring
    │   │   │   │   │   # - Memory optimization
    │   │   │   │   ├── push-notifications.ts  # Push notification setup
    │   │   │   │   │   # - Firebase/Expo notifications
    │   │   │   │   │   # - Token registration
    │   │   │   │   │   # BE: communications-be/push
    │   │   │   │   │   # POST /v1/push-tokens
    │   │   │   │   ├── storage.ts  # Secure storage wrapper
    │   │   │   │   │   # - Token storage
    │   │   │   │   │   # - Biometric keys
    │   │   │   │   └── utils.ts  # General utilities
    │   │   │   ├── lib/i18n/
    │   │   │   │   └── index.ts  # i18n configuration (mobile)
    │   │   │   │       # Uses packages/shared i18n resources
    │   │   │   ├── stores/  # Mobile-specific Zustand stores
    │   │   │   │   ├── biometric-store.ts  # Biometric settings
    │   │   │   │   ├── camera-store.ts  # Camera state
    │   │   │   │   └── offline-queue-store.ts  # Offline action queue
    │   │   │   └── types/  # Mobile-specific types
    │   │   │       ├── biometric.ts  # Biometric types
    │   │   │       └── navigation.ts  # Navigation types
    │   │   ├── .env  # Environment variables
    │   │   ├── .eslintrc.json  # ESLint config
    │   │   ├── app.json  # Expo config
    │   │   │   # - App name, slug
    │   │   │   # - Icons, splash screens
    │   │   │   # - Permissions
    │   │   │   # - Build settings (EAS)
    │   │   ├── babel.config.js  # Babel config
    │   │   ├── eas.json  # EAS Build config
    │   │   │   # - Development build
    │   │   │   # - Preview build
    │   │   │   # - Production build
    │   │   │   # - Credentials
    │   │   ├── global.css  # Global styles (NativeWind)
    │   │   ├── index.js  # Entry point
    │   │   ├── metro.config.js  # Metro bundler config
    │   │   │   # - Monorepo support
    │   │   │   # - Asset resolution
    │   │   ├── package.json  # Mobile dependencies
    │   │   │   # - Expo SDK
    │   │   │   # - React Native
    │   │   │   # - Expo Router
    │   │   │   # - NativeWind
    │   │   │   # - TanStack Query
    │   │   │   # - Zustand
    │   │   ├── tailwind.config.js  # Tailwind config (NativeWind)
    │   │   └── tsconfig.json  # TypeScript config
    │   └── web/  # Next.js web application
    │       └── src/
    │           └── app/
    │               └── [locale]/  # Internationalized routing (en, ar, zh, hi, de, fr, tr, es, ru)
    │                   ├── (auth)/  # Authentication pages (no dashboard layout)
    │                   │   ├── callback/
    │                   │   │   └── page.tsx  # OAuth callback handler
    │                   │   │       # - Keycloak callback
    │                   │   │       # - Google OAuth callback
    │                   │   │       # - GitHub OAuth callback
    │                   │   │       # - LinkedIn OAuth callback
    │                   │   │       # - Token exchange
    │                   │   │       # BE: Keycloak token exchange
    │                   │   │       # POST /oauth2/token (Keycloak)
    │                   │   │       # Returns: access_token, refresh_token
    │                   │   ├── forgot-password/
    │                   │   │   └── page.tsx  # Password reset request
    │                   │   │       # - Email input
    │                   │   │       # - Send reset link
    │                   │   │       # BE: users-be/security/recovery
    │                   │   │       # POST /v1/auth/forgot-password
    │                   │   │       # Body: { email }
    │                   │   │       # Returns: reset_email_sent
    │                   │   ├── login/
    │                   │   │   └── page.tsx  # Login page
    │                   │   │       # - Email/password form
    │                   │   │       # - Social login buttons (Google, GitHub, LinkedIn)
    │                   │   │       # - Remember me option
    │                   │   │       # - Forgot password link
    │                   │   │       # - Register link
    │                   │   │       # BE: Keycloak OAuth2 flow
    │                   │   │       # POST /v1/auth/login (users-be)
    │                   │   │       # Returns: JWT access_token, refresh_token
    │                   │   │       # Social: OAuth2 redirect to Keycloak
    │                   │   ├── mfa/
    │                   │   │   ├── setup/
    │                   │   │   │   └── page.tsx  # MFA setup
    │                   │   │   │       # - QR code display
    │                   │   │   │       # - Backup codes
    │                   │   │   │       # - Verify setup
    │                   │   │   │       # BE: users-be/security/mfa
    │                   │   │   │       # POST /v1/auth/mfa/setup
    │                   │   │   │       # GET /v1/auth/mfa/qrcode
    │                   │   │   │       # POST /v1/auth/mfa/verify-setup
    │                   │   │   └── verify/
    │                   │   │       └── page.tsx  # MFA verification
    │                   │   │           # - OTP input
    │                   │   │           # - Backup code option
    │                   │   │           # - Trust device option
    │                   │   │           # BE: users-be/security/mfa
    │                   │   │           # POST /v1/auth/mfa/verify
    │                   │   │           # Body: { code, trust_device }
    │                   │   ├── register/
    │                   │   │   ├── verification/
    │                   │   │   │   └── page.tsx  # Email verification callback
    │                   │   │   │       # - Verify token from email link
    │                   │   │   │       # - Success/error messaging
    │                   │   │   │       # - Auto-redirect to onboarding
    │                   │   │   │       # BE: users-be/user
    │                   │   │   │       # POST /v1/users/verify-email
    │                   │   │   │       # Body: { token }
    │                   │   │   │       # Returns: verification_status
    │                   │   │   │       # Publishes: UserVerified event
    │                   │   │   └── page.tsx  # Registration page
    │                   │   │       # - User type selection (Freelancer/Client)
    │                   │   │       # - Email, password, name fields
    │                   │   │       # - Terms acceptance
    │                   │   │       # - Social registration
    │                   │   │       # BE: users-be/user
    │                   │   │       # POST /v1/users/register
    │                   │   │       # Body: { email, password, first_name, last_name, user_type, terms_accepted }
    │                   │   │       # Returns: user_id, verification_email_sent
    │                   │   │       # Publishes: UserCreated event
    │                   │   ├── reset-password/
    │                   │   │   └── page.tsx  # Password reset form
    │                   │   │       # - New password input
    │                   │   │       # - Confirm password
    │                   │   │       # - Token validation
    │                   │   │       # BE: users-be/security/recovery
    │                   │   │       # POST /v1/auth/reset-password
    │                   │   │       # Body: { token, new_password }
    │                   │   │       # Returns: password_reset_success
    │                   │   └── layout.tsx  # Auth pages layout
    │                   │       # - Minimal layout (logo + form)
    │                   │       # - No header/footer
    │                   │       # - Language switcher
    │                   ├── (dashboard)/  # Main authenticated dashboard
    │                   │   ├── admin/  # Admin panel (SUPER_ADMIN only)
    │                   │   │   ├── audit-logs/
    │                   │   │   │   └── page.tsx  # Audit logs
    │                   │   │   │       # - All admin actions
    │                   │   │   │       # - Filter by admin
    │                   │   │   │       # - Filter by action type
    │                   │   │   │       # - Search logs
    │                   │   │   │       # BE: admin-be/audit
    │                   │   │   │       # GET /v1/admin/audit-logs
    │                   │   │   ├── business-verification/
    │                   │   │   │   ├── [caseId]/
    │                   │   │   │   │   ├── review/
    │                   │   │   │   │   │   └── page.tsx  # Review business documents
    │                   │   │   │   │   │       # BE: admin-be/business_verification
    │                   │   │   │   │   │       # POST /v1/admin/business-verification/{case_id}/review
    │                   │   │   │   │   └── page.tsx  # Business verification case
    │                   │   │   │   │       # BE: admin-be/business_verification
    │                   │   │   │   │       # GET /v1/admin/business-verification/{case_id}
    │                   │   │   │   └── page.tsx  # Business verification queue
    │                   │   │   │       # BE: admin-be/business_verification
    │                   │   │   │       # GET /v1/admin/business-verification/cases
    │                   │   │   ├── change-approvals/
    │                   │   │   │   ├── [requestId]/
    │                   │   │   │   │   ├── approve/
    │                   │   │   │   │   │   └── page.tsx  # Approve/reject change
    │                   │   │   │   │   │       # BE: admin-be/change_approval
    │                   │   │   │   │   │       # POST /v1/admin/change-approvals/{request_id}/approve
    │                   │   │   │   │   │       # POST /v1/admin/change-approvals/{request_id}/reject
    │                   │   │   │   │   └── page.tsx  # Change request detail
    │                   │   │   │   │       # BE: admin-be/change_approval
    │                   │   │   │   │       # GET /v1/admin/change-approvals/{request_id}
    │                   │   │   │   └── page.tsx  # Change approval queue (Two-Person Rule)
    │                   │   │   │       # - Pending approvals
    │                   │   │   │       # - My requests
    │                   │   │   │       # - History
    │                   │   │   │       # BE: admin-be/change_approval
    │                   │   │   │       # GET /v1/admin/change-approvals
    │                   │   │   ├── compliance/
    │                   │   │   │   ├── aml-kyc/
    │                   │   │   │   │   ├── monitoring/
    │                   │   │   │   │   │   └── page.tsx  # AML monitoring dashboard
    │                   │   │   │   │   │       # - Suspicious activity
    │                   │   │   │   │   │       # - Transaction patterns
    │                   │   │   │   │   │       # - Risk scores
    │                   │   │   │   │   │       # BE: admin-be/kyc_case, financial-be/transaction
    │                   │   │   │   │   │       # GET /v1/kyc/monitoring/suspicious-activity
    │                   │   │   │   │   ├── reports/
    │                   │   │   │   │   │   ├── [reportId]/
    │                   │   │   │   │   │   │   └── page.tsx  # SAR (Suspicious Activity Report) detail
    │                   │   │   │   │   │   │       # BE: admin-be/kyc_case
    │                   │   │   │   │   │   │       # GET /v1/kyc/reports/{report_id}
    │                   │   │   │   │   │   └── page.tsx  # AML reports list
    │                   │   │   │   │   │       # - Filed reports
    │                   │   │   │   │   │       # - Pending reports
    │                   │   │   │   │   │       # BE: admin-be/kyc_case
    │                   │   │   │   │   │       # GET /v1/kyc/reports
    │                   │   │   │   │   └── risk-assessment/
    │                   │   │   │   │       └── page.tsx  # Risk assessment tools
    │                   │   │   │   │           # - User risk profiles
    │                   │   │   │   │           # - Country risk matrix
    │                   │   │   │   │           # - Enhanced due diligence
    │                   │   │   │   │           # BE: admin-be/kyc_case
    │                   │   │   │   │           # GET /v1/kyc/risk-assessment
    │                   │   │   │   ├── data-retention/
    │                   │   │   │   │   ├── audit/
    │                   │   │   │   │   │   └── page.tsx  # Retention audit log
    │                   │   │   │   │   │       # - Deletion history
    │                   │   │   │   │   │       # - Policy compliance
    │                   │   │   │   │   │       # BE: utility/audit
    │                   │   │   │   │   │       # GET /v1/audit/retention
    │                   │   │   │   │   ├── policies/
    │                   │   │   │   │   │   └── page.tsx  # Retention policies
    │                   │   │   │   │   │       # - Policy definitions
    │                   │   │   │   │   │       # - Data categories
    │                   │   │   │   │   │       # - Retention periods
    │                   │   │   │   │   │       # BE: admin-be/data_retention (if exists) or utility/config
    │                   │   │   │   │   │       # GET /v1/retention/policies
    │                   │   │   │   │   └── schedule/
    │                   │   │   │   │       └── page.tsx  # Deletion schedule
    │                   │   │   │   │           # - Upcoming deletions
    │                   │   │   │   │           # - Retention expirations
    │                   │   │   │   │           # BE: admin-be/data_retention
    │                   │   │   │   │           # GET /v1/retention/schedule
    │                   │   │   │   ├── document-verification/
    │                   │   │   │   │   ├── [documentId]/
    │                   │   │   │   │   │   └── page.tsx  # Document review interface
    │                   │   │   │   │   │       # - Document viewer
    │                   │   │   │   │   │       # - Verification checks
    │                   │   │   │   │   │       # - Approve/reject/request-more
    │                   │   │   │   │   │       # BE: admin-be/business_verification, storage-be/asset
    │                   │   │   │   │   │       # GET /v1/verification/documents/{document_id}
    │                   │   │   │   │   │       # PUT /v1/verification/documents/{document_id}/review
    │                   │   │   │   │   ├── automated-checks/
    │                   │   │   │   │   │   └── page.tsx  # Automated verification rules
    │                   │   │   │   │   │       # - OCR settings
    │                   │   │   │   │   │       # - Validation rules
    │                   │   │   │   │   │       # - ML model performance
    │                   │   │   │   │   │       # BE: admin-be/business_verification
    │                   │   │   │   │   │       # GET /v1/verification/automation-rules
    │                   │   │   │   │   └── queue/
    │                   │   │   │   │       └── page.tsx  # Document verification queue
    │                   │   │   │   │           # - Pending documents
    │                   │   │   │   │           # - Priority sorting
    │                   │   │   │   │           # - Auto-verification status
    │                   │   │   │   │           # BE: admin-be/business_verification, storage-be/asset
    │                   │   │   │   │           # GET /v1/verification/documents/queue
    │                   │   │   │   └── gdpr/
    │                   │   │   │       ├── consent-management/
    │                   │   │   │       │   └── page.tsx  # Consent logs & management
    │                   │   │   │       │       # - User consent history
    │                   │   │   │       │       # - Consent versions
    │                   │   │   │       │       # - Audit trail
    │                   │   │   │       │       # BE: users-be/consent, utility/audit
    │                   │   │   │       │       # GET /v1/users/consent-logs
    │                   │   │   │       ├── deletion-requests/
    │                   │   │   │       │   ├── [requestId]/
    │                   │   │   │       │   │   └── page.tsx  # Deletion request detail
    │                   │   │   │       │   │       # - Data preview
    │                   │   │   │       │   │       # - Retention check
    │                   │   │   │       │   │       # - Process deletion
    │                   │   │   │       │   │       # BE: admin-be/privacy, users-be/user
    │                   │   │   │       │   │       # GET /v1/privacy/deletion-requests/{request_id}
    │                   │   │   │       │   │       # POST /v1/privacy/deletion-requests/{request_id}/process
    │                   │   │   │       │   └── page.tsx  # Deletion requests queue
    │                   │   │   │       │       # BE: admin-be/privacy
    │                   │   │   │       │       # GET /v1/privacy/deletion-requests
    │                   │   │   │       ├── export-requests/
    │                   │   │   │       │   ├── [requestId]/
    │                   │   │   │       │   │   └── page.tsx  # Export request detail
    │                   │   │   │       │   │       # - Request review
    │                   │   │   │       │   │       # - Generate export
    │                   │   │   │       │   │       # - Approve/deny
    │                   │   │   │       │   │       # BE: admin-be/privacy, users-be/user
    │                   │   │   │       │   │       # GET /v1/privacy/export-requests/{request_id}
    │                   │   │   │       │   │       # POST /v1/privacy/export-requests/{request_id}/process
    │                   │   │   │       │   └── page.tsx  # Export requests queue
    │                   │   │   │       │       # BE: admin-be/privacy
    │                   │   │   │       │       # GET /v1/privacy/export-requests
    │                   │   │   │       └── reports/
    │                   │   │   │           └── page.tsx  # GDPR compliance reports
    │                   │   │   │               # - Processing activities
    │                   │   │   │               # - Data inventory
    │                   │   │   │               # - Breach reports
    │                   │   │   │               # BE: admin-be/privacy, utility/audit
    │                   │   │   │               # GET /v1/privacy/reports
    │                   │   │   ├── disputes/
    │                   │   │   │   ├── [disputeId]/
    │                   │   │   │   │   ├── assign/
    │                   │   │   │   │   │   └── page.tsx  # Assign dispute to admin
    │                   │   │   │   │   │       # BE: admin-be/disputes
    │                   │   │   │   │   │       # POST /v1/admin/disputes/{dispute_id}/assign
    │                   │   │   │   │   ├── escalate/
    │                   │   │   │   │   │   └── page.tsx  # Escalate dispute
    │                   │   │   │   │   │       # BE: admin-be/disputes
    │                   │   │   │   │   │       # POST /v1/admin/disputes/{dispute_id}/escalate
    │                   │   │   │   │   └── resolve/
    │                   │   │   │   │       └── page.tsx  # Resolve dispute
    │                   │   │   │   │           # - Resolution decision
    │                   │   │   │   │           # - Financial settlement
    │                   │   │   │   │           # - Explanation
    │                   │   │   │   │           # BE: admin-be/disputes
    │                   │   │   │   │           # POST /v1/admin/disputes/{dispute_id}/resolve
    │                   │   │   │   │           # Publishes: DisputeResolved event
    │                   │   │   │   └── page.tsx  # Disputes management
    │                   │   │   │       # - Open disputes
    │                   │   │   │       # - Assigned to me
    │                   │   │   │       # - Resolved disputes
    │                   │   │   │       # BE: admin-be/disputes
    │                   │   │   │       # GET /v1/admin/disputes
    │                   │   │   ├── incidents/
    │                   │   │   │   ├── [incidentId]/
    │                   │   │   │   │   ├── edit/
    │                   │   │   │   │   │   └── page.tsx  # Edit incident details
    │                   │   │   │   │   │       # - Update status
    │                   │   │   │   │   │       # - Add postmortem
    │                   │   │   │   │   │       # - Affected services
    │                   │   │   │   │   │       # BE: utility/status
    │                   │   │   │   │   │       # PUT /v1/status/incidents/{incident_id}
    │                   │   │   │   │   ├── timeline/
    │                   │   │   │   │   │   └── page.tsx  # Incident timeline
    │                   │   │   │   │   │       # - Event log
    │                   │   │   │   │   │       # - Update history
    │                   │   │   │   │   │       # BE: utility/status
    │                   │   │   │   │   │       # GET /v1/status/incidents/{incident_id}/timeline
    │                   │   │   │   │   └── page.tsx  # Incident detail
    │                   │   │   │   │       # - Current status
    │                   │   │   │   │       # - Impact assessment
    │                   │   │   │   │       # - Resolution steps
    │                   │   │   │   │       # BE: utility/status
    │                   │   │   │   │       # GET /v1/status/incidents/{incident_id}
    │                   │   │   │   ├── create/
    │                   │   │   │   │   └── page.tsx  # Create new incident
    │                   │   │   │   │       # - Incident type selection
    │                   │   │   │   │       # - Severity level
    │                   │   │   │   │       # - Affected services
    │                   │   │   │   │       # BE: utility/status
    │                   │   │   │   │       # POST /v1/status/incidents
    │                   │   │   │   ├── history/
    │                   │   │   │   │   └── page.tsx  # Historical incidents
    │                   │   │   │   │       # - Past incidents archive
    │                   │   │   │   │       # - Postmortems
    │                   │   │   │   │       # - Lessons learned
    │                   │   │   │   │       # BE: utility/status
    │                   │   │   │   │       # GET /v1/status/incidents/history
    │                   │   │   │   └── page.tsx  # Active incidents dashboard
    │                   │   │   │       # - Current incidents
    │                   │   │   │       # - Quick actions
    │                   │   │   │       # - Status board
    │                   │   │   │       # BE: utility/status
    │                   │   │   │       # GET /v1/status/incidents?status=active
    │                   │   │   ├── kyc/
    │                   │   │   │   ├── [caseId]/
    │                   │   │   │   │   ├── approve/
    │                   │   │   │   │   │   └── page.tsx  # Approve KYC
    │                   │   │   │   │   │       # BE: admin-be/kyc_case
    │                   │   │   │   │   │       # POST /v1/admin/kyc/cases/{case_id}/approve
    │                   │   │   │   │   │       # Publishes: KYCApproved event
    │                   │   │   │   │   └── reject/
    │                   │   │   │   │       └── page.tsx  # Reject KYC
    │                   │   │   │   │           # - Rejection reason
    │                   │   │   │   │           # - Required actions
    │                   │   │   │   │           # BE: admin-be/kyc_case
    │                   │   │   │   │           # POST /v1/admin/kyc/cases/{case_id}/reject
    │                   │   │   │   └── page.tsx  # KYC cases queue
    │                   │   │   │       # - Pending cases
    │                   │   │   │       # - Approved/rejected cases
    │                   │   │   │       # BE: admin-be/kyc_case
    │                   │   │   │       # GET /v1/admin/kyc/cases
    │                   │   │   ├── maintenance/
    │                   │   │   │   ├── [maintenanceId]/
    │                   │   │   │   │   ├── edit/
    │                   │   │   │   │   │   └── page.tsx  # Edit maintenance window
    │                   │   │   │   │   │       # BE: utility/status
    │                   │   │   │   │   │       # PUT /v1/status/maintenance/{maintenance_id}
    │                   │   │   │   │   └── page.tsx  # Maintenance detail
    │                   │   │   │   │       # BE: utility/status
    │                   │   │   │   │       # GET /v1/status/maintenance/{maintenance_id}
    │                   │   │   │   ├── schedule/
    │                   │   │   │   │   └── page.tsx  # Schedule maintenance
    │                   │   │   │   │       # - Date/time selection
    │                   │   │   │   │       # - Affected services
    │                   │   │   │   │       # - Notification plan
    │                   │   │   │   │       # BE: utility/status
    │                   │   │   │   │       # POST /v1/status/maintenance
    │                   │   │   │   └── page.tsx  # Maintenance calendar
    │                   │   │   │       # - Upcoming maintenance
    │                   │   │   │       # - Impact windows
    │                   │   │   │       # BE: utility/status
    │                   │   │   │       # GET /v1/status/maintenance
    │                   │   │   ├── moderation/
    │                   │   │   │   ├── jobs/
    │                   │   │   │   │   ├── [jobId]/
    │                   │   │   │   │   │   └── review/
    │                   │   │   │   │   │       └── page.tsx  # Review flagged job
    │                   │   │   │   │   │           # - Job content
    │                   │   │   │   │   │           # - Flag reason
    │                   │   │   │   │   │           # - Actions: Approve/Remove/Hide
    │                   │   │   │   │   │           # BE: admin-be/moderation
    │                   │   │   │   │   │           # POST /v1/admin/moderation/jobs/{job_id}/review
    │                   │   │   │   │   └── page.tsx  # Flagged jobs
    │                   │   │   │   │       # BE: admin-be/moderation
    │                   │   │   │   │       # GET /v1/admin/moderation/jobs
    │                   │   │   │   ├── messages/
    │                   │   │   │   │   └── [messageId]/
    │                   │   │   │   │       └── review/
    │                   │   │   │   │           └── page.tsx  # Review flagged message
    │                   │   │   │   │               # BE: admin-be/moderation
    │                   │   │   │   │               # POST /v1/admin/moderation/messages/{message_id}/review
    │                   │   │   │   ├── proposals/
    │                   │   │   │   │   └── [proposalId]/
    │                   │   │   │   │       └── review/
    │                   │   │   │   │           └── page.tsx  # Review flagged proposal
    │                   │   │   │   │               # BE: admin-be/moderation
    │                   │   │   │   │               # POST /v1/admin/moderation/proposals/{proposal_id}/review
    │                   │   │   │   └── reviews/
    │                   │   │   │       └── [reviewId]/
    │                   │   │   │           └── review/
    │                   │   │   │               └── page.tsx  # Review flagged review
    │                   │   │   │                   # BE: admin-be/moderation
    │                   │   │   │                   # POST /v1/admin/moderation/reviews/{review_id}/review
    │                   │   │   ├── platform-config/
    │                   │   │   │   ├── integrations/
    │                   │   │   │   │   ├── [integrationId]/
    │                   │   │   │   │   │   ├── configure/
    │                   │   │   │   │   │   │   └── page.tsx  # Configure integration
    │                   │   │   │   │   │   │       # BE: admin-be/integrations (if exists)
    │                   │   │   │   │   │   │       # PUT /v1/integrations/{integration_id}/config
    │                   │   │   │   │   │   ├── logs/
    │                   │   │   │   │   │   │   └── page.tsx  # Integration logs
    │                   │   │   │   │   │   │       # BE: utility/audit
    │                   │   │   │   │   │   │       # GET /v1/integrations/{integration_id}/logs
    │                   │   │   │   │   │   └── page.tsx  # Integration detail
    │                   │   │   │   │   │       # BE: admin-be/integrations
    │                   │   │   │   │   │       # GET /v1/integrations/{integration_id}
    │                   │   │   │   │   └── page.tsx  # Integrations list
    │                   │   │   │   │       # - Payment providers
    │                   │   │   │   │       # - Email services
    │                   │   │   │   │       # - Storage providers
    │                   │   │   │   │       # - Auth providers
    │                   │   │   │   │       # BE: admin-be/integrations
    │                   │   │   │   │       # GET /v1/integrations
    │                   │   │   │   ├── limits/
    │                   │   │   │   │   └── page.tsx  # Platform limits configuration
    │                   │   │   │   │       # - Rate limits
    │                   │   │   │   │       # - Upload limits
    │                   │   │   │   │       # - API quotas
    │                   │   │   │   │       # - Subscription limits
    │                   │   │   │   │       # BE: utility/config, admin-be/platform_config (if exists)
    │                   │   │   │   │       # GET /v1/config/limits
    │                   │   │   │   │       # PUT /v1/config/limits
    │                   │   │   │   ├── localization/
    │                   │   │   │   │   ├── languages/
    │                   │   │   │   │   │   └── page.tsx  # Language management
    │                   │   │   │   │   │       # - Enabled languages
    │                   │   │   │   │   │       # - Default language
    │                   │   │   │   │   │       # - RTL settings
    │                   │   │   │   │   │       # BE: utility/i18n
    │                   │   │   │   │   │       # GET /v1/config/languages
    │                   │   │   │   │   │       # PUT /v1/config/languages
    │                   │   │   │   │   └── regions/
    │                   │   │   │   │       └── page.tsx  # Regional settings
    │                   │   │   │   │           # - Timezone defaults
    │                   │   │   │   │           # - Currency settings
    │                   │   │   │   │           # - Date/time formats
    │                   │   │   │   │           # BE: utility/config
    │                   │   │   │   │           # GET /v1/config/regions
    │                   │   │   │   │           # PUT /v1/config/regions
    │                   │   │   │   ├── notifications/
    │                   │   │   │   │   ├── settings/
    │                   │   │   │   │   │   └── page.tsx  # Notification settings
    │                   │   │   │   │   │       # - Default preferences
    │                   │   │   │   │   │       # - Delivery channels
    │                   │   │   │   │   │       # - Retry policies
    │                   │   │   │   │   │       # BE: communications-be/config
    │                   │   │   │   │   │       # GET /v1/notifications/config
    │                   │   │   │   │   │       # PUT /v1/notifications/config
    │                   │   │   │   │   └── templates/
    │                   │   │   │   │       ├── [templateId]/
    │                   │   │   │   │       │   ├── edit/
    │                   │   │   │   │       │   │   └── page.tsx  # Edit notification template
    │                   │   │   │   │       │   │       # BE: communications-be/template
    │                   │   │   │   │       │   │       # PUT /v1/notifications/templates/{template_id}
    │                   │   │   │   │       │   ├── preview/
    │                   │   │   │   │       │   │   └── page.tsx  # Preview template
    │                   │   │   │   │       │   │       # BE: communications-be/template
    │                   │   │   │   │       │   │       # POST /v1/notifications/templates/{template_id}/preview
    │                   │   │   │   │       │   └── page.tsx  # Template detail
    │                   │   │   │   │       │       # BE: communications-be/template
    │                   │   │   │   │       │       # GET /v1/notifications/templates/{template_id}
    │                   │   │   │   │       └── page.tsx  # Template library
    │                   │   │   │   │           # BE: communications-be/template
    │                   │   │   │   │           # GET /v1/notifications/templates
    │                   │   │   │   └── pricing/
    │                   │   │   │       └── page.tsx  # Pricing configuration
    │                   │   │   │           # - Commission rates
    │                   │   │   │           # - Subscription pricing
    │                   │   │   │           # - Regional pricing
    │                   │   │   │           # BE: financial-be/pricing_config (if exists)
    │                   │   │   │           # GET /v1/config/pricing
    │                   │   │   │           # PUT /v1/config/pricing
    │                   │   │   │           # Note: Requires change_approval
    │                   │   │   ├── refunds/
    │                   │   │   │   ├── [caseId]/
    │                   │   │   │   │   ├── approve/
    │                   │   │   │   │   │   └── page.tsx  # Approve refund
    │                   │   │   │   │   │       # BE: admin-be/refund_case
    │                   │   │   │   │   │       # POST /v1/admin/refunds/{case_id}/approve
    │                   │   │   │   │   └── reject/
    │                   │   │   │   │       └── page.tsx  # Reject refund
    │                   │   │   │   │           # BE: admin-be/refund_case
    │                   │   │   │   │           # POST /v1/admin/refunds/{case_id}/reject
    │                   │   │   │   └── page.tsx  # Refund cases
    │                   │   │   │       # - Pending refund requests
    │                   │   │   │       # - Processed refunds
    │                   │   │   │       # BE: admin-be/refund_case
    │                   │   │   │       # GET /v1/admin/refunds
    │                   │   │   ├── reports/
    │                   │   │   │   ├── financial/
    │                   │   │   │   │   └── page.tsx  # Financial reports
    │                   │   │   │   │       # BE: admin-be/reports
    │                   │   │   │   │       # GET /v1/admin/reports/financial
    │                   │   │   │   ├── moderation/
    │                   │   │   │   │   └── page.tsx  # Moderation reports
    │                   │   │   │   │       # BE: admin-be/reports
    │                   │   │   │   │       # GET /v1/admin/reports/moderation
    │                   │   │   │   ├── users/
    │                   │   │   │   │   └── page.tsx  # User reports
    │                   │   │   │   │       # BE: admin-be/reports
    │                   │   │   │   │       # GET /v1/admin/reports/users
    │                   │   │   │   └── page.tsx  # Admin reports
    │                   │   │   │       # - Platform metrics
    │                   │   │   │       # - User growth
    │                   │   │   │       # - Revenue reports
    │                   │   │   │       # - Moderation stats
    │                   │   │   │       # BE: admin-be/reports
    │                   │   │   │       # GET /v1/admin/reports
    │                   │   │   ├── system/
    │                   │   │   │   ├── announcements/
    │                   │   │   │   │   ├── create/
    │                   │   │   │   │   │   └── page.tsx  # Create announcement
    │                   │   │   │   │   │       # BE: communications-be/announcements
    │                   │   │   │   │   │       # POST /v1/admin/announcements
    │                   │   │   │   │   └── page.tsx  # System announcements
    │                   │   │   │   │       # BE: communications-be/announcements
    │                   │   │   │   │       # GET /v1/admin/announcements
    │                   │   │   │   ├── feature-flags/
    │                   │   │   │   │   ├── [flagId]/
    │                   │   │   │   │   │   └── edit/
    │                   │   │   │   │   │       └── page.tsx  # Edit feature flag
    │                   │   │   │   │   │           # BE: subscriptions-be/feature_toggles
    │                   │   │   │   │   │           # PUT /v1/admin/feature-flags/{flag_id}
    │                   │   │   │   │   └── page.tsx  # Feature flags management
    │                   │   │   │   │       # - List flags
    │                   │   │   │   │       # - Toggle flags
    │                   │   │   │   │       # - Rollout percentage
    │                   │   │   │   │       # BE: subscriptions-be/feature_toggles
    │                   │   │   │   │       # GET /v1/admin/feature-flags
    │                   │   │   │   ├── maintenance/
    │                   │   │   │   │   └── page.tsx  # Maintenance mode
    │                   │   │   │   │       # - Enable/disable maintenance
    │                   │   │   │   │       # - Maintenance message
    │                   │   │   │   │       # BE: admin-be/system
    │                   │   │   │   │       # POST /v1/admin/system/maintenance
    │                   │   │   │   └── page.tsx  # System settings
    │                   │   │   │       # - Feature flags
    │                   │   │   │       # - System configuration
    │                   │   │   │       # BE: admin-be/system
    │                   │   │   │       # GET /v1/admin/system/config
    │                   │   │   ├── system-health/
    │                   │   │   │   ├── metrics/
    │                   │   │   │   │   └── page.tsx  # System metrics dashboard
    │                   │   │   │   │       # - CPU/Memory usage
    │                   │   │   │   │       # - Database performance
    │                   │   │   │   │       # - Queue depths
    │                   │   │   │   │       # BE: utility/metrics (or monitoring service)
    │                   │   │   │   │       # GET /v1/metrics/system
    │                   │   │   │   ├── services/
    │                   │   │   │   │   └── page.tsx  # Service health overview
    │                   │   │   │   │       # - All microservices status
    │                   │   │   │   │       # - Uptime metrics
    │                   │   │   │   │       # - Response times
    │                   │   │   │   │       # BE: utility/status
    │                   │   │   │   │       # GET /v1/status/services
    │                   │   │   │   └── page.tsx  # Health dashboard
    │                   │   │   │       # - Overall system health
    │                   │   │   │       # - Critical alerts
    │                   │   │   │       # - Performance trends
    │                   │   │   │       # BE: utility/status
    │                   │   │   │       # GET /v1/status/health
    │                   │   │   ├── users/
    │                   │   │   │   ├── [userId]/
    │                   │   │   │   │   ├── ban/
    │                   │   │   │   │   │   └── page.tsx  # Ban user
    │                   │   │   │   │   │       # - Ban reason
    │                   │   │   │   │   │       # - Permanent/temporary
    │                   │   │   │   │   │       # - Related accounts
    │                   │   │   │   │   │       # BE: admin-be/users
    │                   │   │   │   │   │       # POST /v1/admin/users/{user_id}/ban
    │                   │   │   │   │   │       # Publishes: UserBanned event
    │                   │   │   │   │   ├── suspend/
    │                   │   │   │   │   │   └── page.tsx  # Suspend user
    │                   │   │   │   │   │       # - Suspension reason
    │                   │   │   │   │   │       # - Duration
    │                   │   │   │   │   │       # - Notify user
    │                   │   │   │   │   │       # BE: admin-be/users
    │                   │   │   │   │   │       # POST /v1/admin/users/{user_id}/suspend
    │                   │   │   │   │   │       # Publishes: UserSuspended event
    │                   │   │   │   │   ├── verify/
    │                   │   │   │   │   │   └── page.tsx  # Verify user (manual)
    │                   │   │   │   │   │       # BE: admin-be/users
    │                   │   │   │   │   │       # POST /v1/admin/users/{user_id}/verify
    │                   │   │   │   │   └── warn/
    │                   │   │   │   │       └── page.tsx  # Warn user
    │                   │   │   │   │           # BE: admin-be/users
    │                   │   │   │   │           # POST /v1/admin/users/{user_id}/warn
    │                   │   │   │   ├── bulk-actions/
    │                   │   │   │   │   └── page.tsx  # Bulk user actions
    │                   │   │   │   │       # BE: admin-be/users
    │                   │   │   │   │       # POST /v1/admin/users/bulk-action
    │                   │   │   │   └── page.tsx  # Users management
    │                   │   │   │       # - User list
    │                   │   │   │       # - Search users
    │                   │   │   │       # - Filter by status/type
    │                   │   │   │       # - Bulk actions
    │                   │   │   │       # BE: admin-be/users
    │                   │   │   │       # GET /v1/admin/users?filters={...}
    │                   │   │   ├── layout.tsx  # Admin layout (RBAC guard)
    │                   │   │   │   # - Only ADMIN, SUPER_ADMIN, MODERATOR
    │                   │   │   │   # - Admin navigation sidebar
    │                   │   │   └── page.tsx  # Admin dashboard
    │                   │   │       # - Key metrics
    │                   │   │       # - Pending moderation queue
    │                   │   │       # - Recent admin actions
    │                   │   │       # - System alerts
    │                   │   │       # BE: admin-be/dashboard
    │                   │   │       # GET /v1/admin/dashboard
    │                   │   ├── analytics/
    │                   │   │   ├── clients/
    │                   │   │   │   └── page.tsx  # Client analytics (freelancer)
    │                   │   │   │       # - Top clients
    │                   │   │   │       # - Client retention rate
    │                   │   │   │       # - Repeat business
    │                   │   │   │       # - Client lifetime value
    │                   │   │   │       # BE: users-be/analytics
    │                   │   │   │       # GET /v1/users/me/analytics/clients
    │                   │   │   ├── earnings/
    │                   │   │   │   ├── forecast/
    │                   │   │   │   │   └── page.tsx  # Earnings forecast
    │                   │   │   │   │       # - Projected earnings
    │                   │   │   │   │       # - Based on current trends
    │                   │   │   │   │       # - Scenario analysis
    │                   │   │   │   │       # BE: users-be/analytics
    │                   │   │   │   │       # GET /v1/users/me/analytics/earnings/forecast
    │                   │   │   │   └── page.tsx  # Earnings analytics
    │                   │   │   │       # - Monthly earnings
    │                   │   │   │       # - Yearly earnings
    │                   │   │   │       # - Client breakdown
    │                   │   │   │       # - Tax estimates
    │                   │   │   │       # BE: users-be/analytics
    │                   │   │   │       # GET /v1/users/me/analytics/earnings
    │                   │   │   ├── freelancers/
    │                   │   │   │   └── page.tsx  # Freelancer analytics (client)
    │                   │   │   │       # - Top freelancers
    │                   │   │   │       # - Freelancer retention
    │                   │   │   │       # - Performance scores
    │                   │   │   │       # - Cost efficiency
    │                   │   │   │       # BE: users-be/analytics
    │                   │   │   │       # GET /v1/users/me/analytics/freelancers
    │                   │   │   ├── performance/
    │                   │   │   │   └── page.tsx  # Performance analytics
    │                   │   │   │       # - Response time metrics
    │                   │   │   │       # - Proposal success rate
    │                   │   │   │       # - Client satisfaction
    │                   │   │   │       # BE: users-be/analytics, proposals-be/performance
    │                   │   │   │       # GET /v1/users/me/analytics/performance
    │                   │   │   ├── projects/
    │                   │   │   │   └── page.tsx  # Project analytics
    │                   │   │   │       # - Completion rates
    │                   │   │   │       # - Average duration
    │                   │   │   │       # - Budget adherence
    │                   │   │   │       # - Success metrics
    │                   │   │   │       # BE: users-be/analytics
    │                   │   │   │       # GET /v1/users/me/analytics/projects
    │                   │   │   ├── spending/
    │                   │   │   │   ├── forecast/
    │                   │   │   │   │   └── page.tsx  # Spending forecast
    │                   │   │   │   │       # - Projected spending
    │                   │   │   │   │       # - Based on active contracts
    │                   │   │   │   │       # - Budget overrun warnings
    │                   │   │   │   │       # BE: users-be/analytics
    │                   │   │   │   │       # GET /v1/users/me/analytics/spending/forecast
    │                   │   │   │   └── page.tsx  # Spending analytics (client)
    │                   │   │   │       # - Monthly spending
    │                   │   │   │       # - Yearly spending
    │                   │   │   │       # - Freelancer breakdown
    │                   │   │   │       # - Project breakdown
    │                   │   │   │       # BE: users-be/analytics
    │                   │   │   │       # GET /v1/users/me/analytics/spending
    │                   │   │   └── page.tsx  # Analytics dashboard
    │                   │   │       # - Overview metrics
    │                   │   │       # - Customizable widgets
    │                   │   │       # - Time range selector
    │                   │   │       # - Export reports
    │                   │   │       # BE: users-be/analytics
    │                   │   │       # GET /v1/users/me/analytics/dashboard
    │                   │   ├── availability/
    │                   │   │   ├── calendar/
    │                   │   │   │   └── page.tsx  # Availability calendar
    │                   │   │   │       # - Mark available/busy
    │                   │   │   │       # - Recurring patterns
    │                   │   │   │       # - Sync with external calendar
    │                   │   │   │       # BE: users-be/availability (if exists) or users-be/profile
    │                   │   │   │       # GET /v1/users/me/availability
    │                   │   │   │       # PUT /v1/users/me/availability
    │                   │   │   ├── settings/
    │                   │   │   │   └── page.tsx  # Availability settings
    │                   │   │   │       # - Working hours
    │                   │   │   │       # - Timezone preferences
    │                   │   │   │       # - Auto-reply settings
    │                   │   │   │       # BE: users-be/settings
    │                   │   │   │       # PUT /v1/users/me/availability-settings
    │                   │   │   └── page.tsx  # Availability dashboard
    │                   │   │       # - Current status
    │                   │   │       # - Upcoming commitments
    │                   │   │       # BE: users-be/availability
    │                   │   │       # GET /v1/users/me/availability-overview
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
    │                   │   ├── contracts/  # Contracts management
    │                   │   │   ├── [contractId]/
    │                   │   │   │   ├── amendments/
    │                   │   │   │   │   ├── [amendmentId]/
    │                   │   │   │   │   │   ├── approve/
    │                   │   │   │   │   │   │   └── page.tsx  # Approve/reject amendment
    │                   │   │   │   │   │   │       # BE: contracts-be/amendment
    │                   │   │   │   │   │   │       # POST /v1/amendments/{amendment_id}/approve
    │                   │   │   │   │   │   │       # POST /v1/amendments/{amendment_id}/reject
    │                   │   │   │   │   │   └── page.tsx  # Amendment detail
    │                   │   │   │   │   │       # BE: contracts-be/amendment
    │                   │   │   │   │   │       # GET /v1/amendments/{amendment_id}
    │                   │   │   │   │   ├── propose/
    │                   │   │   │   │   │   └── page.tsx  # Propose amendment
    │                   │   │   │   │   │       # - Change description
    │                   │   │   │   │   │       # - Updated terms
    │                   │   │   │   │   │       # - Reason
    │                   │   │   │   │   │       # BE: contracts-be/amendment
    │                   │   │   │   │   │       # POST /v1/contracts/{contract_id}/amendments
    │                   │   │   │   │   └── page.tsx  # Contract amendments list
    │                   │   │   │   │       # - Proposed amendments
    │                   │   │   │   │       # - Accepted amendments
    │                   │   │   │   │       # BE: contracts-be/amendment
    │                   │   │   │   │       # GET /v1/contracts/{contract_id}/amendments
    │                   │   │   │   ├── complete/
    │                   │   │   │   │   └── page.tsx  # Complete contract
    │                   │   │   │   │       # - Confirm all deliverables received
    │                   │   │   │   │       # - Leave review
    │                   │   │   │   │       # - Final payment release
    │                   │   │   │   │       # BE: contracts-be/contract
    │                   │   │   │   │       # POST /v1/contracts/{contract_id}/complete
    │                   │   │   │   │       # Publishes: ContractCompleted event
    │                   │   │   │   │       # Redirects to: /reviews/create
    │                   │   │   │   ├── deliverables/
    │                   │   │   │   │   ├── [deliverableId]/
    │                   │   │   │   │   │   └── page.tsx  # Deliverable detail
    │                   │   │   │   │   │       # - File preview
    │                   │   │   │   │   │       # - Download
    │                   │   │   │   │   │       # - Approval status
    │                   │   │   │   │   │       # - Feedback
    │                   │   │   │   │   │       # BE: contracts-be/deliverable
    │                   │   │   │   │   │       # GET /v1/deliverables/{deliverable_id}
    │                   │   │   │   │   │       # BE: storage-be
    │                   │   │   │   │   │       # GET /v1/storage/download/{file_id}
    │                   │   │   │   │   └── page.tsx  # All deliverables
    │                   │   │   │   │       # - List all submitted deliverables
    │                   │   │   │   │       # - Status (pending, approved, rejected)
    │                   │   │   │   │       # BE: contracts-be/deliverable
    │                   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables
    │                   │   │   │   ├── details/
    │                   │   │   │   │   └── page.tsx  # Full contract details
    │                   │   │   │   │       # - Contract terms
    │                   │   │   │   │       # - Scope of work (SOW)
    │                   │   │   │   │       # - Payment terms
    │                   │   │   │   │       # - Deadlines
    │                   │   │   │   │       # - Special clauses
    │                   │   │   │   │       # BE: contracts-be/contract
    │                   │   │   │   │       # GET /v1/contracts/{contract_id}/details
    │                   │   │   │   │       # BE: contracts-be/sow
    │                   │   │   │   │       # GET /v1/contracts/{contract_id}/sow
    │                   │   │   │   ├── disputes/
    │                   │   │   │   │   ├── [disputeId]/
    │                   │   │   │   │   │   ├── escalate/
    │                   │   │   │   │   │   │   └── page.tsx  # Escalate to admin/mediation
    │                   │   │   │   │   │   │       # BE: contracts-be/dispute
    │                   │   │   │   │   │   │       # POST /v1/disputes/{dispute_id}/escalate
    │                   │   │   │   │   │   │       # BE: admin-be
    │                   │   │   │   │   │   │       # Creates mediation case
    │                   │   │   │   │   │   ├── respond/
    │                   │   │   │   │   │   │   └── page.tsx  # Respond to dispute
    │                   │   │   │   │   │   │       # - Response message
    │                   │   │   │   │   │   │       # - Additional evidence
    │                   │   │   │   │   │   │       # BE: contracts-be/dispute
    │                   │   │   │   │   │   │       # POST /v1/disputes/{dispute_id}/responses
    │                   │   │   │   │   │   └── page.tsx  # Dispute detail
    │                   │   │   │   │   │       # - Dispute timeline
    │                   │   │   │   │   │       # - Messages/responses
    │                   │   │   │   │   │       # - Evidence
    │                   │   │   │   │   │       # - Admin notes (if assigned)
    │                   │   │   │   │   │       # - Resolution status
    │                   │   │   │   │   │       # BE: contracts-be/dispute
    │                   │   │   │   │   │       # GET /v1/disputes/{dispute_id}
    │                   │   │   │   │   ├── open/
    │                   │   │   │   │   │   └── page.tsx  # Open a dispute
    │                   │   │   │   │   │       # - Dispute reason
    │                   │   │   │   │   │       # - Description
    │                   │   │   │   │   │       # - Evidence upload
    │                   │   │   │   │   │       # - Desired resolution
    │                   │   │   │   │   │       # BE: contracts-be/dispute
    │                   │   │   │   │   │       # POST /v1/contracts/{contract_id}/disputes
    │                   │   │   │   │   │       # BE: storage-be/uploads
    │                   │   │   │   │   │       # Publishes: DisputeOpened event
    │                   │   │   │   │   └── page.tsx  # Disputes list
    │                   │   │   │   │       # - Active disputes
    │                   │   │   │   │       # - Resolved disputes
    │                   │   │   │   │       # BE: contracts-be/dispute
    │                   │   │   │   │       # GET /v1/contracts/{contract_id}/disputes
    │                   │   │   │   ├── feedback/
    │                   │   │   │   │   └── page.tsx  # Ongoing feedback
    │                   │   │   │   │       # - Mid-contract feedback
    │                   │   │   │   │       # - Performance notes
    │                   │   │   │   │       # BE: contracts-be/feedback
    │                   │   │   │   │       # GET /v1/contracts/{contract_id}/feedback
    │                   │   │   │   │       # POST /v1/contracts/{contract_id}/feedback
    │                   │   │   │   ├── messages/
    │                   │   │   │   │   └── page.tsx  # Contract-specific messages
    │                   │   │   │   │       # - Threaded conversation
    │                   │   │   │   │       # - File sharing
    │                   │   │   │   │       # - Quick links to milestones/deliverables
    │                   │   │   │   │       # BE: communications-be/conversations
    │                   │   │   │   │       # GET /v1/contracts/{contract_id}/conversation
    │                   │   │   │   ├── milestones/
    │                   │   │   │   │   ├── [milestoneId]/
    │                   │   │   │   │   │   ├── approve/
    │                   │   │   │   │   │   │   └── page.tsx  # Approve milestone (client)
    │                   │   │   │   │   │   │       # - Review deliverables
    │                   │   │   │   │   │   │       # - Accept/request changes
    │                   │   │   │   │   │   │       # - Approval notes
    │                   │   │   │   │   │   │       # BE: contracts-be/milestone
    │                   │   │   │   │   │   │       # POST /v1/milestones/{milestone_id}/approve
    │                   │   │   │   │   │   │       # Publishes: MilestoneApproved event
    │                   │   │   │   │   │   │       # Triggers: Escrow release (financial-be)
    │                   │   │   │   │   │   ├── dispute/
    │                   │   │   │   │   │   │   └── page.tsx  # Dispute milestone
    │                   │   │   │   │   │   │       # - Reason for dispute
    │                   │   │   │   │   │   │       # - Evidence upload
    │                   │   │   │   │   │   │       # BE: contracts-be/dispute
    │                   │   │   │   │   │   │       # POST /v1/milestones/{milestone_id}/dispute
    │                   │   │   │   │   │   └── submit/
    │                   │   │   │   │   │       └── page.tsx  # Submit deliverables (freelancer)
    │                   │   │   │   │   │           # - Upload files
    │                   │   │   │   │   │           # - Completion notes
    │                   │   │   │   │   │           # BE: contracts-be/deliverable
    │                   │   │   │   │   │           # POST /v1/milestones/{milestone_id}/deliverables
    │                   │   │   │   │   │           # BE: storage-be/uploads
    │                   │   │   │   │   │           # Publishes: MilestoneCompleted event
    │                   │   │   │   │   ├── create/
    │                   │   │   │   │   │   └── page.tsx  # Create milestone (if contract allows)
    │                   │   │   │   │   │       # - Milestone title
    │                   │   │   │   │   │       # - Description
    │                   │   │   │   │   │       # - Amount
    │                   │   │   │   │   │       # - Due date
    │                   │   │   │   │   │       # - Deliverables
    │                   │   │   │   │   │       # BE: contracts-be/milestone
    │                   │   │   │   │   │       # POST /v1/contracts/{contract_id}/milestones
    │                   │   │   │   │   │       # Publishes: MilestoneCreated event
    │                   │   │   │   │   └── page.tsx  # Milestones list
    │                   │   │   │   │       # - All milestones
    │                   │   │   │   │       # - Status (pending, in_progress, completed)
    │                   │   │   │   │       # - Amount & due date
    │                   │   │   │   │       # - Approval status
    │                   │   │   │   │       # BE: contracts-be/milestone
    │                   │   │   │   │       # GET /v1/contracts/{contract_id}/milestones
    │                   │   │   │   ├── pause/
    │                   │   │   │   │   └── page.tsx  # Pause contract
    │                   │   │   │   │       # - Reason
    │                   │   │   │   │       # - Expected resume date
    │                   │   │   │   │       # - Notify other party
    │                   │   │   │   │       # BE: contracts-be/contract
    │                   │   │   │   │       # POST /v1/contracts/{contract_id}/pause
    │                   │   │   │   │       # Publishes: ContractPaused event
    │                   │   │   │   ├── payments/
    │                   │   │   │   │   └── page.tsx  # Contract payments
    │                   │   │   │   │       # - Payment schedule
    │                   │   │   │   │       # - Escrow status
    │                   │   │   │   │       # - Released payments
    │                   │   │   │   │       # - Pending payments
    │                   │   │   │   │       # BE: financial-be/escrow
    │                   │   │   │   │       # GET /v1/contracts/{contract_id}/escrow
    │                   │   │   │   │       # BE: financial-be/payment
    │                   │   │   │   │       # GET /v1/contracts/{contract_id}/payments
    │                   │   │   │   ├── terminate/
    │                   │   │   │   │   └── page.tsx  # Terminate contract
    │                   │   │   │   │       # - Termination reason
    │                   │   │   │   │       # - Early termination terms
    │                   │   │   │   │       # - Final deliverables
    │                   │   │   │   │       # - Escrow settlement
    │                   │   │   │   │       # BE: contracts-be/termination
    │                   │   │   │   │       # POST /v1/contracts/{contract_id}/terminate
    │                   │   │   │   │       # Publishes: ContractTerminated event
    │                   │   │   │   ├── timesheet/
    │                   │   │   │   │   ├── [timesheetId]/
    │                   │   │   │   │   │   └── page.tsx  # Timesheet detail
    │                   │   │   │   │   │       # BE: contracts-be/timesheet
    │                   │   │   │   │   │       # GET /v1/timesheets/{timesheet_id}
    │                   │   │   │   │   ├── approve/
    │                   │   │   │   │   │   └── page.tsx  # Approve timesheet (client)
    │                   │   │   │   │   │       # - Review hours
    │                   │   │   │   │   │       # - Approve/request changes
    │                   │   │   │   │   │       # BE: contracts-be/timesheet
    │                   │   │   │   │   │       # POST /v1/timesheets/{timesheet_id}/approve
    │                   │   │   │   │   ├── submit/
    │                   │   │   │   │   │   └── page.tsx  # Submit timesheet (freelancer)
    │                   │   │   │   │   │       # - Hours worked per day
    │                   │   │   │   │   │       # - Task descriptions
    │                   │   │   │   │   │       # - Billable/non-billable
    │                   │   │   │   │   │       # BE: contracts-be/timesheet
    │                   │   │   │   │   │       # POST /v1/contracts/{contract_id}/timesheets
    │                   │   │   │   │   │       # Publishes: TimesheetSubmitted event
    │                   │   │   │   │   └── page.tsx  # Timesheet view (hourly contracts)
    │                   │   │   │   │       # - Weekly/monthly view
    │                   │   │   │   │       # - Total hours
    │                   │   │   │   │       # - Approval status
    │                   │   │   │   │       # - Submit for review
    │                   │   │   │   │       # BE: contracts-be/timesheet
    │                   │   │   │   │       # GET /v1/contracts/{contract_id}/timesheets
    │                   │   │   │   ├── work-diary/
    │                   │   │   │   │   ├── [date]/
    │                   │   │   │   │   │   └── page.tsx  # Work diary for specific date
    │                   │   │   │   │   │       # BE: contracts-be/work_diary
    │                   │   │   │   │   │       # GET /v1/contracts/{contract_id}/work-diary?date={date}
    │                   │   │   │   │   ├── add-entry/
    │                   │   │   │   │   │   └── page.tsx  # Add work diary entry (freelancer)
    │                   │   │   │   │   │       # - Date & time
    │                   │   │   │   │   │       # - Hours worked
    │                   │   │   │   │   │       # - Description
    │                   │   │   │   │   │       # - Upload screenshot (optional)
    │                   │   │   │   │   │       # BE: contracts-be/work_diary
    │                   │   │   │   │   │       # POST /v1/contracts/{contract_id}/work-diary/entries
    │                   │   │   │   │   │       # BE: storage-be/uploads
    │                   │   │   │   │   └── page.tsx  # Work diary overview
    │                   │   │   │   │       # - Daily activity logs
    │                   │   │   │   │       # - Screenshots (if enabled)
    │                   │   │   │   │       # - Productivity metrics
    │                   │   │   │   │       # - Calendar view
    │                   │   │   │   │       # BE: contracts-be/work_diary
    │                   │   │   │   │       # GET /v1/contracts/{contract_id}/work-diary
    │                   │   │   │   └── page.tsx  # Contract overview
    │                   │   │   │       # - Contract details
    │                   │   │   │       # - Parties involved
    │                   │   │   │       # - Budget/rate
    │                   │   │   │       # - Timeline
    │                   │   │   │       # - Status
    │                   │   │   │       # - Quick actions (message, submit work, etc.)
    │                   │   │   │       # BE: contracts-be/contract
    │                   │   │   │       # GET /v1/contracts/{contract_id}
    │                   │   │   ├── active/
    │                   │   │   │   └── page.tsx  # Active contracts only
    │                   │   │   │       # BE: contracts-be/contract
    │                   │   │   │       # GET /v1/contracts?status=active
    │                   │   │   ├── completed/
    │                   │   │   │   └── page.tsx  # Completed contracts
    │                   │   │   │       # BE: contracts-be/contract
    │                   │   │   │       # GET /v1/contracts?status=completed
    │                   │   │   ├── recurring/
    │                   │   │   │   ├── [contractId]/
    │                   │   │   │   │   └── renew/
    │                   │   │   │   │       └── page.tsx  # Renew recurring contract
    │                   │   │   │   │           # BE: contracts-be/recurring
    │                   │   │   │   │           # POST /v1/contracts/{contract_id}/renew
    │                   │   │   │   └── page.tsx  # Recurring contracts
    │                   │   │   │       # - List recurring contracts
    │                   │   │   │       # - Renewal schedule
    │                   │   │   │       # BE: contracts-be/recurring
    │                   │   │   │       # GET /v1/contracts/recurring
    │                   │   │   ├── templates/
    │                   │   │   │   ├── [templateId]/
    │                   │   │   │   │   └── edit/
    │                   │   │   │   │       └── page.tsx  # Edit template
    │                   │   │   │   │           # BE: contracts-be/template
    │                   │   │   │   │           # PUT /v1/contract-templates/{template_id}
    │                   │   │   │   ├── create/
    │                   │   │   │   │   └── page.tsx  # Create contract template
    │                   │   │   │   │       # BE: contracts-be/template
    │                   │   │   │   │       # POST /v1/contract-templates
    │                   │   │   │   └── page.tsx  # Contract templates (for recurring work)
    │                   │   │   │       # BE: contracts-be/template
    │                   │   │   │       # GET /v1/contract-templates
    │                   │   │   └── page.tsx  # Contracts list
    │                   │   │       # - Active contracts
    │                   │   │       # - Completed contracts
    │                   │   │       # - Filter by status
    │                   │   │       # - Sort options
    │                   │   │       # BE: contracts-be/contract
    │                   │   │       # GET /v1/contracts?status=active
    │                   │   ├── dashboard/  # Dashboard home (role-based view)
    │                   │   │   └── page.tsx  # Main dashboard
    │                   │   │       # Freelancer view:
    │                   │   │       # - Active proposals
    │                   │   │       # - Active contracts
    │                   │   │       # - Earnings overview
    │                   │   │       # - Job recommendations
    │                   │   │       # - Profile completion score
    │                   │   │       # Client view:
    │                   │   │       # - Active jobs
    │                   │   │       # - Spending overview
    │                   │   │       # - Recent proposals
    │                   │   │       # - Talent recommendations
    │                   │   │       # BE: users-be/user
    │                   │   │       # GET /v1/users/me
    │                   │   │       # BE: Multiple services for dashboard data
    │                   │   │       # GET /v1/analytics/dashboard (analytics service)
    │                   │   │       # GET /v1/jobs/my-jobs (jobs-be)
    │                   │   │       # GET /v1/proposals/my-proposals (proposals-be)
    │                   │   │       # GET /v1/contracts/active (contracts-be)
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
    │                   │   ├── financials/  # Financial management
    │                   │   │   ├── escrow/
    │                   │   │   │   ├── [escrowId]/
    │                   │   │   │   │   └── page.tsx  # Escrow detail
    │                   │   │   │   │       # - Related contract
    │                   │   │   │   │       # - Amount held
    │                   │   │   │   │       # - Release schedule
    │                   │   │   │   │       # - Transaction history
    │                   │   │   │   │       # BE: financial-be/escrow
    │                   │   │   │   │       # GET /v1/escrow/{escrow_id}
    │                   │   │   │   └── page.tsx  # Escrow overview
    │                   │   │   │       # - Active escrow accounts
    │                   │   │   │       # - Total amount in escrow
    │                   │   │   │       # - Pending releases
    │                   │   │   │       # BE: financial-be/escrow
    │                   │   │   │       # GET /v1/escrow
    │                   │   │   ├── invoices/
    │                   │   │   │   ├── [invoiceId]/
    │                   │   │   │   │   ├── pay/
    │                   │   │   │   │   │   └── page.tsx  # Pay invoice (client)
    │                   │   │   │   │   │       # - Invoice summary
    │                   │   │   │   │   │       # - Payment method selection
    │                   │   │   │   │   │       # - Process payment
    │                   │   │   │   │   │       # BE: financial-be/payment
    │                   │   │   │   │   │       # POST /v1/invoices/{invoice_id}/pay
    │                   │   │   │   │   └── page.tsx  # Invoice detail
    │                   │   │   │   │       # - Invoice information
    │                   │   │   │   │       # - Line items
    │                   │   │   │   │       # - Tax details
    │                   │   │   │   │       # - Payment status
    │                   │   │   │   │       # - Download PDF
    │                   │   │   │   │       # BE: financial-be/invoice
    │                   │   │   │   │       # GET /v1/invoices/{invoice_id}
    │                   │   │   │   │       # GET /v1/invoices/{invoice_id}/pdf
    │                   │   │   │   ├── create/
    │                   │   │   │   │   └── page.tsx  # Create invoice (manual invoicing)
    │                   │   │   │   │       # - Client selection
    │                   │   │   │   │       # - Line items
    │                   │   │   │   │       # - Tax settings
    │                   │   │   │   │       # - Due date
    │                   │   │   │   │       # - Notes
    │                   │   │   │   │       # BE: financial-be/invoice
    │                   │   │   │   │       # POST /v1/invoices
    │                   │   │   │   └── page.tsx  # Invoices list
    │                   │   │   │       # - Sent invoices (freelancer)
    │                   │   │   │       # - Received invoices (client)
    │                   │   │   │       # - Filter by status (paid, pending, overdue)
    │                   │   │   │       # BE: financial-be/invoice
    │                   │   │   │       # GET /v1/invoices
    │                   │   │   ├── payment-methods/
    │                   │   │   │   ├── [methodId]/
    │                   │   │   │   │   ├── delete/
    │                   │   │   │   │   │   └── page.tsx  # Delete payment method
    │                   │   │   │   │   │       # BE: financial-be/payment_method
    │                   │   │   │   │   │       # DELETE /v1/payment-methods/{method_id}
    │                   │   │   │   │   └── page.tsx  # Payment method detail
    │                   │   │   │   │       # BE: financial-be/payment_method
    │                   │   │   │   │       # GET /v1/payment-methods/{method_id}
    │                   │   │   │   ├── add/
    │                   │   │   │   │   └── page.tsx  # Add payment method
    │                   │   │   │   │       # - Card details (Stripe Elements)
    │                   │   │   │   │       # - PayPal connection
    │                   │   │   │   │       # - Bank account (ACH)
    │                   │   │   │   │       # - Set as default
    │                   │   │   │   │       # BE: financial-be/payment_method
    │                   │   │   │   │       # POST /v1/payment-methods
    │                   │   │   │   └── page.tsx  # Payment methods list
    │                   │   │   │       # - Saved credit cards
    │                   │   │   │       # - PayPal accounts
    │                   │   │   │       # - Bank accounts
    │                   │   │   │       # - Default payment method
    │                   │   │   │       # BE: financial-be/payment_method
    │                   │   │   │       # GET /v1/payment-methods
    │                   │   │   ├── payout-methods/
    │                   │   │   │   ├── [methodId]/
    │                   │   │   │   │   └── page.tsx  # Payout method detail
    │                   │   │   │   │       # BE: financial-be/payout_method
    │                   │   │   │   │       # GET /v1/payout-methods/{method_id}
    │                   │   │   │   │       # DELETE /v1/payout-methods/{method_id}
    │                   │   │   │   ├── add/
    │                   │   │   │   │   └── page.tsx  # Add payout method
    │                   │   │   │   │       # - Bank account details
    │                   │   │   │   │       # - PayPal email
    │                   │   │   │   │       # - Wire transfer info
    │                   │   │   │   │       # - Tax forms (W-9, W-8BEN)
    │                   │   │   │   │       # BE: financial-be/payout_method
    │                   │   │   │   │       # POST /v1/payout-methods
    │                   │   │   │   └── page.tsx  # Payout methods list (freelancer)
    │                   │   │   │       # - Bank accounts
    │                   │   │   │       # - PayPal
    │                   │   │   │       # - Wire transfer details
    │                   │   │   │       # BE: financial-be/payout_method
    │                   │   │   │       # GET /v1/payout-methods
    │                   │   │   ├── reports/
    │                   │   │   │   ├── earnings/
    │                   │   │   │   │   └── page.tsx  # Detailed earnings report
    │                   │   │   │   │       # - By project
    │                   │   │   │   │       # - By client
    │                   │   │   │   │       # - By time period
    │                   │   │   │   │       # BE: financial-be/reports
    │                   │   │   │   │       # GET /v1/reports/earnings/detailed
    │                   │   │   │   ├── spending/
    │                   │   │   │   │   └── page.tsx  # Detailed spending report
    │                   │   │   │   │       # - By project
    │                   │   │   │   │       # - By freelancer
    │                   │   │   │   │       # - By category
    │                   │   │   │   │       # BE: financial-be/reports
    │                   │   │   │   │       # GET /v1/reports/spending/detailed
    │                   │   │   │   └── page.tsx  # Financial reports
    │                   │   │   │       # - Earnings report (freelancer)
    │                   │   │   │       # - Spending report (client)
    │                   │   │   │       # - Tax report
    │                   │   │   │       # - Date range selection
    │                   │   │   │       # - Export options
    │                   │   │   │       # BE: financial-be/reports
    │                   │   │   │       # GET /v1/reports/earnings
    │                   │   │   │       # GET /v1/reports/spending
    │                   │   │   ├── tax/
    │                   │   │   │   ├── forms/
    │                   │   │   │   │   ├── upload/
    │                   │   │   │   │   │   └── page.tsx  # Upload tax form
    │                   │   │   │   │   │       # BE: financial-be/tax
    │                   │   │   │   │   │       # POST /v1/tax/forms
    │                   │   │   │   │   │       # BE: storage-be/uploads
    │                   │   │   │   │   └── page.tsx  # Tax forms list
    │                   │   │   │   │       # - W-9, 1099, W-8BEN, etc.
    │                   │   │   │   │       # - Download forms
    │                   │   │   │   │       # BE: financial-be/tax
    │                   │   │   │   │       # GET /v1/tax/forms
    │                   │   │   │   ├── settings/
    │                   │   │   │   │   └── page.tsx  # Tax settings
    │                   │   │   │   │       # - Tax information
    │                   │   │   │   │       # - VAT reverse charge
    │                   │   │   │   │       # - Tax exemptions
    │                   │   │   │   │       # BE: financial-be/tax
    │                   │   │   │   │       # PUT /v1/tax/settings
    │                   │   │   │   └── page.tsx  # Tax information
    │                   │   │   │       # - Tax forms
    │                   │   │   │       # - Tax ID
    │                   │   │   │       # - VAT/GST number
    │                   │   │   │       # - Tax residency
    │                   │   │   │       # BE: financial-be/tax
    │                   │   │   │       # GET /v1/tax/info
    │                   │   │   ├── transactions/
    │                   │   │   │   ├── [transactionId]/
    │                   │   │   │   │   └── page.tsx  # Transaction detail
    │                   │   │   │   │       # - Full transaction info
    │                   │   │   │   │       # - Related contract/job
    │                   │   │   │   │       # - Receipt download
    │                   │   │   │   │       # BE: financial-be/transaction
    │                   │   │   │   │       # GET /v1/transactions/{transaction_id}
    │                   │   │   │   └── page.tsx  # Transaction history
    │                   │   │   │       # - All transactions
    │                   │   │   │       # - Filter by type (payment, payout, refund, etc.)
    │                   │   │   │       # - Filter by date range
    │                   │   │   │       # - Search by description
    │                   │   │   │       # - Export to CSV
    │                   │   │   │       # BE: financial-be/transaction
    │                   │   │   │       # GET /v1/transactions?filters={...}
    │                   │   │   ├── wallet/
    │                   │   │   │   ├── add-funds/
    │                   │   │   │   │   └── page.tsx  # Add funds (client)
    │                   │   │   │   │       # - Amount input
    │                   │   │   │   │       # - Payment method selection
    │                   │   │   │   │       # - Payment processing
    │                   │   │   │   │       # BE: financial-be/wallet
    │                   │   │   │   │       # POST /v1/wallet/add-funds
    │                   │   │   │   │       # BE: financial-be/payment
    │                   │   │   │   │       # POST /v1/payments (Stripe/PayPal integration)
    │                   │   │   │   ├── withdraw/
    │                   │   │   │   │   └── page.tsx  # Withdraw funds (freelancer)
    │                   │   │   │   │       # - Amount input
    │                   │   │   │   │       # - Payout method selection
    │                   │   │   │   │       # - Tax information
    │                   │   │   │   │       # - Withdrawal fees
    │                   │   │   │   │       # BE: financial-be/payout
    │                   │   │   │   │       # POST /v1/payouts/request
    │                   │   │   │   │       # Publishes: PayoutRequested event
    │                   │   │   │   └── page.tsx  # Wallet details
    │                   │   │   │       # - Available balance
    │                   │   │   │       # - Pending balance
    │                   │   │   │       # - Escrow balance
    │                   │   │   │       # - Add funds button (client)
    │                   │   │   │       # - Withdraw button (freelancer)
    │                   │   │   │       # - Transaction history
    │                   │   │   │       # BE: financial-be/wallet
    │                   │   │   │       # GET /v1/wallet
    │                   │   │   └── page.tsx  # Financial overview
    │                   │   │       # - Wallet balance
    │                   │   │       # - Pending payments
    │                   │   │       # - Recent transactions
    │                   │   │       # - Earnings chart (freelancer)
    │                   │   │       # - Spending chart (client)
    │                   │   │       # BE: financial-be/wallet
    │                   │   │       # GET /v1/wallet/balance
    │                   │   │       # BE: financial-be/transaction
    │                   │   │       # GET /v1/transactions/recent
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
    │                   │   ├── job-alerts/
    │                   │   │   ├── [alertId]/
    │                   │   │   │   ├── edit/
    │                   │   │   │   │   └── page.tsx  # Edit job alert
    │                   │   │   │   │       # BE: search-be/alert (if exists) or search-be/saved-search
    │                   │   │   │   │       # PUT /v1/job-alerts/{alert_id}
    │                   │   │   │   ├── history/
    │                   │   │   │   │   └── page.tsx  # Alert history
    │                   │   │   │   │       # - Jobs matched
    │                   │   │   │   │       # - Notifications sent
    │                   │   │   │   │       # BE: search-be/alert
    │                   │   │   │   │       # GET /v1/job-alerts/{alert_id}/history
    │                   │   │   │   └── page.tsx  # Alert detail
    │                   │   │   │       # BE: search-be/alert
    │                   │   │   │       # GET /v1/job-alerts/{alert_id}
    │                   │   │   ├── create/
    │                   │   │   │   └── page.tsx  # Create job alert
    │                   │   │   │       # - Search criteria
    │                   │   │   │       # - Notification frequency
    │                   │   │   │       # - Delivery channel
    │                   │   │   │       # BE: search-be/alert
    │                   │   │   │       # POST /v1/job-alerts
    │                   │   │   └── page.tsx  # Job alerts list
    │                   │   │       # - Active alerts
    │                   │   │       # - Pause/resume
    │                   │   │       # BE: search-be/alert
    │                   │   │       # GET /v1/job-alerts
    │                   │   ├── jobs/  # Jobs management
    │                   │   │   ├── [jobId]/
    │                   │   │   │   ├── analytics/
    │                   │   │   │   │   └── page.tsx  # Job analytics (client)
    │                   │   │   │   │       # - Views
    │                   │   │   │   │       # - Proposals received
    │                   │   │   │   │       # - Proposal conversion rate
    │                   │   │   │   │       # - Time to hire
    │                   │   │   │   │       # BE: jobs-be/analytics
    │                   │   │   │   │       # GET /v1/jobs/{job_id}/analytics
    │                   │   │   │   ├── bidding/
    │                   │   │   │   │   ├── place-bid/
    │                   │   │   │   │   │   └── page.tsx  # Place/update bid (freelancer)
    │                   │   │   │   │   │       # - Current bid amount
    │                   │   │   │   │   │       # - Minimum bid
    │                   │   │   │   │   │       # - Place new bid
    │                   │   │   │   │   │       # - Bid increment rules
    │                   │   │   │   │   │       # - Outbid warning
    │                   │   │   │   │   │       # BE: proposals-be/bidding
    │                   │   │   │   │   │       # POST /v1/jobs/{job_id}/bids
    │                   │   │   │   │   │       # PUT /v1/bids/{bid_id}
    │                   │   │   │   │   │       # Publishes: BidPlaced, BidUpdated, OutbidAlert events
    │                   │   │   │   │   └── page.tsx  # Active bids on job (client view)
    │                   │   │   │   │       # - Real-time bid updates
    │                   │   │   │   │       # - Current lowest bid
    │                   │   │   │   │       # - Bid history
    │                   │   │   │   │       # - Accept bid
    │                   │   │   │   │       # BE: proposals-be/bidding
    │                   │   │   │   │       # GET /v1/jobs/{job_id}/bids
    │                   │   │   │   │       # WebSocket: ws://proposals-be/v1/jobs/{job_id}/bids
    │                   │   │   │   ├── close/
    │                   │   │   │   │   └── page.tsx  # Close job
    │                   │   │   │   │       # - Reason for closing
    │                   │   │   │   │       # - Notify applicants
    │                   │   │   │   │       # BE: jobs-be/job
    │                   │   │   │   │       # POST /v1/jobs/{job_id}/close
    │                   │   │   │   │       # Publishes: JobClosed event
    │                   │   │   │   ├── edit/
    │                   │   │   │   │   └── page.tsx  # Edit job (client only)
    │                   │   │   │   │       # - Same form as post job
    │                   │   │   │   │       # - Cannot edit if has accepted proposals
    │                   │   │   │   │       # BE: jobs-be/job
    │                   │   │   │   │       # PUT /v1/jobs/{job_id}
    │                   │   │   │   │       # Publishes: JobUpdated event
    │                   │   │   │   ├── invite/
    │                   │   │   │   │   └── page.tsx  # Invite freelancers (client)
    │                   │   │   │   │       # - Search freelancers
    │                   │   │   │   │       # - Send invitation with message
    │                   │   │   │   │       # BE: jobs-be/invitations
    │                   │   │   │   │       # POST /v1/jobs/{job_id}/invitations
    │                   │   │   │   │       # BE: search-be
    │                   │   │   │   │       # POST /v1/search/freelancers
    │                   │   │   │   │       # BE: communications-be
    │                   │   │   │   │       # Publishes: JobInvitationSent event
    │                   │   │   │   ├── proposals/
    │                   │   │   │   │   ├── [proposalId]/
    │                   │   │   │   │   │   └── page.tsx  # Proposal detail
    │                   │   │   │   │   │       # - Full proposal view
    │                   │   │   │   │   │       # - Freelancer profile preview
    │                   │   │   │   │   │       # - Accept/Reject buttons
    │                   │   │   │   │   │       # - Shortlist button
    │                   │   │   │   │   │       # - Message freelancer
    │                   │   │   │   │   │       # BE: proposals-be
    │                   │   │   │   │   │       # GET /v1/proposals/{proposal_id}
    │                   │   │   │   │   │       # POST /v1/proposals/{proposal_id}/accept
    │                   │   │   │   │   │       # POST /v1/proposals/{proposal_id}/reject
    │                   │   │   │   │   │       # POST /v1/proposals/{proposal_id}/shortlist
    │                   │   │   │   │   └── page.tsx  # Proposals received (client)
    │                   │   │   │   │       # - List all proposals
    │                   │   │   │   │       # - Filter (all/shortlisted/archived)
    │                   │   │   │   │       # - Sort (date, rate, rating)
    │                   │   │   │   │       # - Proposal cards with key info
    │                   │   │   │   │       # BE: proposals-be
    │                   │   │   │   │       # GET /v1/proposals?job_id={job_id}
    │                   │   │   │   └── page.tsx  # Job detail page
    │                   │   │   │       # - Full job description
    │                   │   │   │       # - Client info
    │                   │   │   │       # - Skills required
    │                   │   │   │       # - Budget/rate
    │                   │   │   │       # - Proposals count
    │                   │   │   │       # - Similar jobs
    │                   │   │   │       # - "Submit Proposal" button (freelancer)
    │                   │   │   │       # - Save job button
    │                   │   │   │       # BE: jobs-be/job
    │                   │   │   │       # GET /v1/jobs/{job_id}
    │                   │   │   │       # BE: proposals-be
    │                   │   │   │       # GET /v1/proposals/count?job_id={job_id}
    │                   │   │   │       # BE: search-be/similarity
    │                   │   │   │       # GET /v1/similarity/jobs/{job_id}
    │                   │   │   ├── browse/
    │                   │   │   │   └── page.tsx  # Job listings with filters
    │                   │   │   │       # - Category filters
    │                   │   │   │       # - Budget range
    │                   │   │   │       # - Experience level
    │                   │   │   │       # - Job type (fixed/hourly)
    │                   │   │   │       # - Location preferences
    │                   │   │   │       # - Skills required
    │                   │   │   │       # - Posted date
    │                   │   │   │       # - Saved jobs indicator
    │                   │   │   │       # - "Best Matches" tab
    │                   │   │   │       # BE: jobs-be/job
    │                   │   │   │       # GET /v1/jobs/browse?filters={...}
    │                   │   │   │       # BE: search-be/query
    │                   │   │   │       # POST /v1/search/jobs (for advanced search)
    │                   │   │   ├── categories/
    │                   │   │   │   ├── [categoryId]/
    │                   │   │   │   │   └── page.tsx  # Jobs in category
    │                   │   │   │   │       # BE: jobs-be/job
    │                   │   │   │   │       # GET /v1/jobs?category_id={category_id}
    │                   │   │   │   └── page.tsx  # Browse by category
    │                   │   │   │       # - Category grid
    │                   │   │   │       # - Subcategories
    │                   │   │   │       # BE: jobs-be/categories
    │                   │   │   │       # GET /v1/jobs/categories
    │                   │   │   ├── drafts/
    │                   │   │   │   ├── [draftId]/
    │                   │   │   │   │   └── edit/
    │                   │   │   │   │       └── page.tsx  # Edit draft
    │                   │   │   │   │           # BE: jobs-be/draft
    │                   │   │   │   │           # PUT /v1/jobs/drafts/{draft_id}
    │                   │   │   │   │           # DELETE /v1/jobs/drafts/{draft_id}
    │                   │   │   │   └── page.tsx  # Job drafts list
    │                   │   │   │       # BE: jobs-be/draft
    │                   │   │   │       # GET /v1/jobs/drafts
    │                   │   │   ├── invitations/
    │                   │   │   │   └── page.tsx  # Job invitations received
    │                   │   │   │       # BE: jobs-be/invitations
    │                   │   │   │       # GET /v1/jobs/invitations
    │                   │   │   ├── my-jobs/
    │                   │   │   │   ├── active/
    │                   │   │   │   │   └── page.tsx  # Active jobs only
    │                   │   │   │   │       # BE: jobs-be/job
    │                   │   │   │   │       # GET /v1/jobs/my-jobs?status=active
    │                   │   │   │   ├── closed/
    │                   │   │   │   │   └── page.tsx  # Closed jobs
    │                   │   │   │   │       # BE: jobs-be/job
    │                   │   │   │   │       # GET /v1/jobs/my-jobs?status=closed
    │                   │   │   │   └── page.tsx  # All posted jobs
    │                   │   │   │       # - Active jobs
    │                   │   │   │       # - Closed jobs
    │                   │   │   │       # - Drafts
    │                   │   │   │       # BE: jobs-be/job
    │                   │   │   │       # GET /v1/jobs/my-jobs?status=active
    │                   │   │   ├── post/
    │                   │   │   │   └── page.tsx  # Post a new job (client)
    │                   │   │   │       # - Job title
    │                   │   │   │       # - Job description (rich text editor)
    │                   │   │   │       # - Category selection
    │                   │   │   │       # - Required skills (autocomplete)
    │                   │   │   │       # - Experience level
    │                   │   │   │       # - Job type (fixed/hourly)
    │                   │   │   │       # - Budget/rate
    │                   │   │   │       # - Duration
    │                   │   │   │       # - Attachments
    │                   │   │   │       # - Screening questions
    │                   │   │   │       # - Visibility (public/private/invited)
    │                   │   │   │       # - Save as draft
    │                   │   │   │       # BE: jobs-be/job
    │                   │   │   │       # POST /v1/jobs
    │                   │   │   │       # Body: { title, description, category_id, skills, budget, ... }
    │                   │   │   │       # BE: jobs-be/attachments
    │                   │   │   │       # POST /v1/jobs/{job_id}/attachments
    │                   │   │   │       # BE: jobs-be/screening
    │                   │   │   │       # POST /v1/jobs/{job_id}/screening-questions
    │                   │   │   │       # BE: storage-be/uploads
    │                   │   │   │       # Publishes: JobPosted event
    │                   │   │   ├── recommendations/
    │                   │   │   │   └── page.tsx  # Recommended jobs (freelancer)
    │                   │   │   │       # - ML-powered job recommendations
    │                   │   │   │       # - Based on skills, history, preferences
    │                   │   │   │       # BE: search-be/recommendations
    │                   │   │   │       # GET /v1/recommendations/jobs
    │                   │   │   ├── saved/
    │                   │   │   │   └── page.tsx  # Saved/bookmarked jobs
    │                   │   │   │       # BE: jobs-be/saved_jobs
    │                   │   │   │       # GET /v1/jobs/saved
    │                   │   │   │       # DELETE /v1/jobs/saved/{job_id}
    │                   │   │   └── page.tsx  # Jobs list (role-based)
    │                   │   │       # Freelancer view: Browse available jobs
    │                   │   │       # Client view: My posted jobs
    │                   │   │       # - Filters (category, budget, skills, etc.)
    │                   │   │       # - Search bar
    │                   │   │       # - Sort options
    │                   │   │       # - Pagination
    │                   │   │       # BE: jobs-be/job
    │                   │   │       # GET /v1/jobs?filters=...&page=1&limit=20
    │                   │   │       # Freelancer: GET /v1/jobs/browse
    │                   │   │       # Client: GET /v1/jobs/my-jobs
    │                   │   ├── learning/
    │                   │   │   ├── achievements/
    │                   │   │   │   └── page.tsx  # Achievements & badges
    │                   │   │   │       # - Earned badges
    │                   │   │   │       # - Progress to next level
    │                   │   │   │       # - Leaderboard
    │                   │   │   │       # BE: users-be/achievement
    │                   │   │   │       # GET /v1/users/me/achievements
    │                   │   │   │       # Learning achievements
    │                   │   │   │       # - Certificates earned
    │                   │   │   │       # - Badges
    │                   │   │   │       # - Skill verification
    │                   │   │   ├── assessments/
    │                   │   │   │   ├── [assessmentId]/
    │                   │   │   │   │   ├── results/
    │                   │   │   │   │   │   └── page.tsx  # Assessment results
    │                   │   │   │   │   │       # BE: learning-be/assessment
    │                   │   │   │   │   │       # GET /v1/learning/assessments/{assessment_id}/results
    │                   │   │   │   │   ├── take/
    │                   │   │   │   │   │   └── page.tsx  # Take assessment
    │                   │   │   │   │   │       # BE: learning-be/assessment
    │                   │   │   │   │   │       # POST /v1/learning/assessments/{assessment_id}/submit
    │                   │   │   │   │   └── page.tsx  # Assessment detail
    │                   │   │   │   │       # BE: learning-be/assessment
    │                   │   │   │   │       # GET /v1/learning/assessments/{assessment_id}
    │                   │   │   │   └── page.tsx  # Assessments list
    │                   │   │   │       # BE: learning-be/assessment
    │                   │   │   │       # GET /v1/learning/assessments
    │                   │   │   ├── certifications/
    │                   │   │   │   └── page.tsx  # Manage certifications
    │                   │   │   │       # - Upload certificates
    │                   │   │   │       # - Verification status
    │                   │   │   │       # - Expiry tracking
    │                   │   │   │       # BE: users-be/credential
    │                   │   │   │       # GET /v1/users/me/credentials?type=certification
    │                   │   │   ├── courses/
    │                   │   │   │   ├── [courseId]/
    │                   │   │   │   │   ├── lessons/
    │                   │   │   │   │   │   └── [lessonId]/
    │                   │   │   │   │   │       └── page.tsx  # Lesson content
    │                   │   │   │   │   │           # BE: learning-be/lesson (if exists) or external LMS
    │                   │   │   │   │   │           # GET /v1/learning/courses/{course_id}/lessons/{lesson_id}
    │                   │   │   │   │   ├── progress/
    │                   │   │   │   │   │   └── page.tsx  # Course progress
    │                   │   │   │   │   │       # BE: learning-be/progress
    │                   │   │   │   │   │       # GET /v1/learning/courses/{course_id}/progress
    │                   │   │   │   │   └── page.tsx  # Course detail
    │                   │   │   │   │       # BE: learning-be/course
    │                   │   │   │   │       # GET /v1/learning/courses/{course_id}
    │                   │   │   │   └── page.tsx  # Courses catalog
    │                   │   │   │       # BE: learning-be/course
    │                   │   │   │       # GET /v1/learning/courses
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
    │                   │   ├── messages/  # Messaging
    │                   │   │   ├── [conversationId]/
    │                   │   │   │   ├── archive/
    │                   │   │   │   │   └── page.tsx  # Archive conversation
    │                   │   │   │   │       # BE: communications-be/conversations
    │                   │   │   │   │       # POST /v1/conversations/{conversation_id}/archive
    │                   │   │   │   ├── info/
    │                   │   │   │   │   └── page.tsx  # Conversation info
    │                   │   │   │   │       # - Participants
    │                   │   │   │   │       # - Related job/contract
    │                   │   │   │   │       # - Shared files
    │                   │   │   │   │       # - Search in conversation
    │                   │   │   │   │       # BE: communications-be/conversations
    │                   │   │   │   │       # GET /v1/conversations/{conversation_id}/info
    │                   │   │   │   └── page.tsx  # Conversation thread
    │                   │   │   │       # - Message history
    │                   │   │   │       # - Real-time new messages
    │                   │   │   │       # - Message composer
    │                   │   │   │       # - File attachments
    │                   │   │   │       # - Typing indicators
    │                   │   │   │       # - Read receipts
    │                   │   │   │       # - Quick actions (block, report)
    │                   │   │   │       # BE: communications-be/messages
    │                   │   │   │       # GET /v1/conversations/{conversation_id}/messages
    │                   │   │   │       # POST /v1/messages
    │                   │   │   │       # WebSocket: Real-time message delivery
    │                   │   │   │       # Publishes: MessageSent event
    │                   │   │   ├── archived/
    │                   │   │   │   └── page.tsx  # Archived conversations
    │                   │   │   │       # BE: communications-be/conversations
    │                   │   │   │       # GET /v1/conversations?archived=true
    │                   │   │   └── new/
    │                   │   │       └── page.tsx  # Start new conversation
    │                   │   │           # - Select recipient (search users)
    │                   │   │           # - Initial message
    │                   │   │           # BE: communications-be/conversations
    │                   │   │           # POST /v1/conversations
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
    │                   │   ├── notifications/  # Notifications center
    │                   │   │   ├── [notificationId]/
    │                   │   │   │   └── page.tsx  # Notification detail (redirects to relevant page)
    │                   │   │   │       # - Mark as read
    │                   │   │   │       # - Navigate to related entity
    │                   │   │   │       # BE: communications-be/notifications
    │                   │   │   │       # POST /v1/notifications/{notif_id}/read
    │                   │   │   ├── settings/
    │                   │   │   │   └── page.tsx  # Notification settings
    │                   │   │   │       # - Email notifications
    │                   │   │   │       # - Push notifications
    │                   │   │   │       # - In-app notifications
    │                   │   │   │       # - Notification preferences by type
    │                   │   │   │       # - Frequency settings
    │                   │   │   │       # - Quiet hours
    │                   │   │   │       # BE: communications-be/preferences
    │                   │   │   │       # GET /v1/notifications/preferences
    │                   │   │   │       # PUT /v1/notifications/preferences
    │                   │   │   ├── unread/
    │                   │   │   │   └── page.tsx  # Unread notifications only
    │                   │   │   │       # BE: communications-be/notifications
    │                   │   │   │       # GET /v1/notifications?unread=true
    │                   │   │   └── page.tsx  # Notifications center
    │                   │   │       # - All notifications list
    │                   │   │       # - Unread indicator
    │                   │   │       # - Mark all as read
    │                   │   │       # - Filter by type
    │                   │   │       # - Real-time updates
    │                   │   │       # BE: communications-be/notifications
    │                   │   │       # GET /v1/notifications
    │                   │   │       # POST /v1/notifications/read-all
    │                   │   │       # WebSocket: ws://communications-be/v1/notifications
    │                   │   ├── organization/  # Organization management (for clients)
    │                   │   │   ├── analytics/
    │                   │   │   │   └── page.tsx  # Organization analytics
    │                   │   │   │       # - Hiring metrics
    │                   │   │   │       # - Freelancer performance
    │                   │   │   │       # - Cost per hire
    │                   │   │   │       # - Time to hire
    │                   │   │   │       # BE: analytics-be
    │                   │   │   │       # GET /v1/analytics/organization/{org_id}
    │                   │   │   ├── settings/
    │                   │   │   │   ├── billing/
    │                   │   │   │   │   └── page.tsx  # Organization billing
    │                   │   │   │   │       # - Billing profile
    │                   │   │   │   │       # - Tax information
    │                   │   │   │   │       # - Payment methods
    │                   │   │   │   │       # BE: financial-be/billing_profile
    │                   │   │   │   │       # GET /v1/organizations/{org_id}/billing-profile
    │                   │   │   │   └── page.tsx  # Organization settings
    │                   │   │   │       # - Company name
    │                   │   │   │       # - Industry
    │                   │   │   │       # - Company size
    │                   │   │   │       # - Website
    │                   │   │   │       # - Logo
    │                   │   │   │       # BE: users-be/organization
    │                   │   │   │       # PATCH /v1/organizations/{org_id}
    │                   │   │   ├── spending/
    │                   │   │   │   ├── budgets/
    │                   │   │   │   │   ├── create/
    │                   │   │   │   │   │   └── page.tsx  # Create budget
    │                   │   │   │   │   │       # BE: financial-be/budget
    │                   │   │   │   │   │       # POST /v1/organizations/{org_id}/budgets
    │                   │   │   │   │   └── page.tsx  # Budget management
    │                   │   │   │   │       # - Set budgets
    │                   │   │   │   │       # - Budget alerts
    │                   │   │   │   │       # BE: financial-be/budget
    │                   │   │   │   │       # GET /v1/organizations/{org_id}/budgets
    │                   │   │   │   └── page.tsx  # Spending overview
    │                   │   │   │       # - Total spending
    │                   │   │   │       # - By project
    │                   │   │   │       # - By freelancer
    │                   │   │   │       # - By time period
    │                   │   │   │       # BE: financial-be/reports
    │                   │   │   │       # GET /v1/organizations/{org_id}/spending
    │                   │   │   ├── team/
    │                   │   │   │   ├── [memberId]/
    │                   │   │   │   │   ├── edit/
    │                   │   │   │   │   │   └── page.tsx  # Edit member
    │                   │   │   │   │   │       # - Change role
    │                   │   │   │   │   │       # - Update permissions
    │                   │   │   │   │   │       # BE: users-be/team
    │                   │   │   │   │   │       # PATCH /v1/organizations/{org_id}/members/{member_id}
    │                   │   │   │   │   └── remove/
    │                   │   │   │   │       └── page.tsx  # Remove member
    │                   │   │   │   │           # BE: users-be/team
    │                   │   │   │   │           # DELETE /v1/organizations/{org_id}/members/{member_id}
    │                   │   │   │   ├── invite/
    │                   │   │   │   │   └── page.tsx  # Invite team member
    │                   │   │   │   │       # - Email address
    │                   │   │   │   │       # - Role selection
    │                   │   │   │   │       # - Permissions
    │                   │   │   │   │       # BE: users-be/team
    │                   │   │   │   │       # POST /v1/organizations/{org_id}/members/invite
    │                   │   │   │   │       # BE: communications-be
    │                   │   │   │   │       # Sends invitation email
    │                   │   │   │   ├── roles/
    │                   │   │   │   │   ├── [roleId]/
    │                   │   │   │   │   │   └── edit/
    │                   │   │   │   │   │       └── page.tsx  # Edit role
    │                   │   │   │   │   │           # BE: users-be/role
    │                   │   │   │   │   │           # PUT /v1/organizations/{org_id}/roles/{role_id}
    │                   │   │   │   │   │           # DELETE /v1/organizations/{org_id}/roles/{role_id}
    │                   │   │   │   │   └── create/
    │                   │   │   │   │       └── page.tsx  # Create custom role
    │                   │   │   │   │           # - Role name
    │                   │   │   │   │           # - Permissions selection
    │                   │   │   │   │           # BE: users-be/role
    │                   │   │   │   │           # POST /v1/organizations/{org_id}/roles
    │                   │   │   │   └── page.tsx  # Team members list
    │                   │   │   │       # - Active members
    │                   │   │   │       # - Pending invitations
    │                   │   │   │       # - Roles
    │                   │   │   │       # BE: users-be/team
    │                   │   │   │       # GET /v1/organizations/{org_id}/members
    │                   │   │   └── page.tsx  # Organization overview
    │                   │   │       # - Company details
    │                   │   │       # - Team members
    │                   │   │       # - Spending overview
    │                   │   │       # - Active contracts
    │                   │   │       # BE: users-be/organization
    │                   │   │       # GET /v1/organizations/{org_id}
    │                   │   ├── profile/  # Current user profile management
    │                   │   │   ├── availability/
    │                   │   │   │   └── page.tsx  # Availability management
    │                   │   │   │       # - Calendar view
    │                   │   │   │       # - Set available hours
    │                   │   │   │       # - Time zone
    │                   │   │   │       # - Vacation mode
    │                   │   │   │       # - Max concurrent projects
    │                   │   │   │       # BE: users-be/availability
    │                   │   │   │       # GET /v1/users/{id}/availability
    │                   │   │   │       # POST /v1/users/{id}/availability
    │                   │   │   │       # PATCH /v1/users/{id}/availability
    │                   │   │   ├── certifications/
    │                   │   │   │   ├── add/
    │                   │   │   │   │   └── page.tsx  # Add certification
    │                   │   │   │   │       # - Certification name
    │                   │   │   │   │       # - Issuing organization
    │                   │   │   │   │       # - Issue date
    │                   │   │   │   │       # - Expiry date (if any)
    │                   │   │   │   │       # - Credential ID
    │                   │   │   │   │       # - Credential URL
    │                   │   │   │   │       # - Certificate upload
    │                   │   │   │   │       # BE: users-be/credentials
    │                   │   │   │   │       # POST /v1/users/{id}/certifications
    │                   │   │   │   │       # BE: storage-be/uploads
    │                   │   │   │   ├── verify/
    │                   │   │   │   │   └── [certificationId]/
    │                   │   │   │   │       └── page.tsx  # Verification request
    │                   │   │   │   │           # - Submit for verification
    │                   │   │   │   │           # - Upload proof
    │                   │   │   │   │           # BE: users-be/credentials
    │                   │   │   │   │           # POST /v1/users/{id}/certifications/{cert_id}/verify
    │                   │   │   │   └── page.tsx  # Certifications list
    │                   │   │   │       # - External certifications
    │                   │   │   │       # - Platform certifications
    │                   │   │   │       # - Badges earned
    │                   │   │   │       # BE: users-be/credentials
    │                   │   │   │       # GET /v1/users/{id}/credentials
    │                   │   │   ├── edit/
    │                   │   │   │   └── page.tsx  # Edit profile form
    │                   │   │   │       # - Basic info (name, title, bio)
    │                   │   │   │       # - Profile photo upload
    │                   │   │   │       # - Location
    │                   │   │   │       # - Languages
    │                   │   │   │       # - Hourly rate (freelancer)
    │                   │   │   │       # - Professional headline
    │                   │   │   │       # BE: users-be/profile
    │                   │   │   │       # PATCH /v1/users/{id}/profile
    │                   │   │   │       # BE: storage-be/uploads
    │                   │   │   │       # POST /v1/storage/upload (photo)
    │                   │   │   ├── education/
    │                   │   │   │   ├── [educationId]/
    │                   │   │   │   │   └── edit/
    │                   │   │   │   │       └── page.tsx  # Edit education form
    │                   │   │   │   │           # BE: users-be/education
    │                   │   │   │   │           # PUT /v1/users/{id}/education/{edu_id}
    │                   │   │   │   │           # DELETE /v1/users/{id}/education/{edu_id}
    │                   │   │   │   ├── add/
    │                   │   │   │   │   └── page.tsx  # Add education form
    │                   │   │   │   │       # - School/university
    │                   │   │   │   │       # - Degree/qualification
    │                   │   │   │   │       # - Field of study
    │                   │   │   │   │       # - Graduation year
    │                   │   │   │   │       # - GPA (optional)
    │                   │   │   │   │       # - Description
    │                   │   │   │   │       # BE: users-be/education
    │                   │   │   │   │       # POST /v1/users/{id}/education
    │                   │   │   │   └── page.tsx  # Education list
    │                   │   │   │       # BE: users-be/education
    │                   │   │   │       # GET /v1/users/{id}/education
    │                   │   │   ├── experience/
    │                   │   │   │   ├── [experienceId]/
    │                   │   │   │   │   └── edit/
    │                   │   │   │   │       └── page.tsx  # Edit experience form
    │                   │   │   │   │           # BE: users-be/experience
    │                   │   │   │   │           # PUT /v1/users/{id}/experience/{exp_id}
    │                   │   │   │   │           # DELETE /v1/users/{id}/experience/{exp_id}
    │                   │   │   │   ├── add/
    │                   │   │   │   │   └── page.tsx  # Add experience form
    │                   │   │   │   │       # - Company name
    │                   │   │   │   │       # - Position title
    │                   │   │   │   │       # - Start/end dates
    │                   │   │   │   │       # - Current position checkbox
    │                   │   │   │   │       # - Description (rich text)
    │                   │   │   │   │       # - Skills used
    │                   │   │   │   │       # BE: users-be/experience
    │                   │   │   │   │       # POST /v1/users/{id}/experience
    │                   │   │   │   └── page.tsx  # Work experience list
    │                   │   │   │       # - List all experience entries
    │                   │   │   │       # - Add new experience button
    │                   │   │   │       # - Edit/delete actions
    │                   │   │   │       # BE: users-be/experience
    │                   │   │   │       # GET /v1/users/{id}/experience
    │                   │   │   ├── portfolio/
    │                   │   │   │   ├── [portfolioId]/
    │                   │   │   │   │   ├── edit/
    │                   │   │   │   │   │   └── page.tsx  # Edit portfolio item
    │                   │   │   │   │   │       # BE: users-be/portfolio
    │                   │   │   │   │   │       # PUT /v1/users/{id}/portfolio/{item_id}
    │                   │   │   │   │   │       # DELETE /v1/users/{id}/portfolio/{item_id}
    │                   │   │   │   │   └── page.tsx  # Portfolio item detail
    │                   │   │   │   │       # BE: users-be/portfolio
    │                   │   │   │   │       # GET /v1/users/{id}/portfolio/{item_id}
    │                   │   │   │   ├── add/
    │                   │   │   │   │   └── page.tsx  # Add portfolio item
    │                   │   │   │   │       # - Project title
    │                   │   │   │   │       # - Description
    │                   │   │   │   │       # - Media upload (images/videos)
    │                   │   │   │   │       # - Project URL
    │                   │   │   │   │       # - Skills used
    │                   │   │   │   │       # - Client (optional)
    │                   │   │   │   │       # - Completion date
    │                   │   │   │   │       # BE: users-be/portfolio
    │                   │   │   │   │       # POST /v1/users/{id}/portfolio
    │                   │   │   │   │       # BE: storage-be/uploads
    │                   │   │   │   │       # POST /v1/storage/upload (media)
    │                   │   │   │   ├── reorder/
    │                   │   │   │   │   └── page.tsx  # Reorder portfolio items
    │                   │   │   │   │       # - Drag & drop interface
    │                   │   │   │   │       # - Set featured items
    │                   │   │   │   │       # BE: users-be/portfolio
    │                   │   │   │   │       # PATCH /v1/users/{id}/portfolio/reorder
    │                   │   │   │   │       # Body: { item_ids: string[] }
    │                   │   │   │   └── page.tsx  # Portfolio items list
    │                   │   │   │       # - Grid/list view
    │                   │   │   │       # - Featured items
    │                   │   │   │       # - Reorder items (drag & drop)
    │                   │   │   │       # BE: users-be/portfolio
    │                   │   │   │       # GET /v1/users/{id}/portfolio
    │                   │   │   ├── service-catalog/
    │                   │   │   │   ├── [serviceId]/
    │                   │   │   │   │   └── edit/
    │                   │   │   │   │       └── page.tsx  # Edit service
    │                   │   │   │   │           # BE: users-be/service_catalog
    │                   │   │   │   │           # PUT /v1/users/{id}/services/{service_id}
    │                   │   │   │   │           # DELETE /v1/users/{id}/services/{service_id}
    │                   │   │   │   ├── add/
    │                   │   │   │   │   └── page.tsx  # Add service
    │                   │   │   │   │       # - Service name
    │                   │   │   │   │       # - Description
    │                   │   │   │   │       # - Capabilities required
    │                   │   │   │   │       # - Delivery time
    │                   │   │   │   │       # - Pricing
    │                   │   │   │   │       # - Packages (Basic/Standard/Premium)
    │                   │   │   │   │       # BE: users-be/service_catalog
    │                   │   │   │   │       # POST /v1/users/{id}/services
    │                   │   │   │   └── page.tsx  # Service catalog management (freelancer)
    │                   │   │   │       # - List offered services
    │                   │   │   │       # - Service packages
    │                   │   │   │       # - Pricing tiers
    │                   │   │   │       # BE: users-be/service_catalog
    │                   │   │   │       # GET /v1/users/{id}/service-catalog
    │                   │   │   ├── skills/
    │                   │   │   │   ├── specializations/
    │                   │   │   │   │   └── page.tsx  # Specializations & niche expertise
    │                   │   │   │   │       # - Add specializations
    │                   │   │   │   │       # - Verification status
    │                   │   │   │   │       # - Niche expertise tags
    │                   │   │   │   │       # BE: users-be/capabilities
    │                   │   │   │   │       # GET /v1/users/{id}/specializations
    │                   │   │   │   │       # POST /v1/users/{id}/specializations
    │                   │   │   │   └── page.tsx  # Skills management
    │                   │   │   │       # - List current skills with levels
    │                   │   │   │       # - Add new skills (autocomplete)
    │                   │   │   │       # - Edit skill proficiency
    │                   │   │   │       # - Remove skills
    │                   │   │   │       # - Primary skills (max 5)
    │                   │   │   │       # BE: users-be/capabilities
    │                   │   │   │       # GET /v1/users/{id}/skills
    │                   │   │   │       # POST /v1/users/{id}/skills
    │                   │   │   │       # PUT /v1/users/{id}/skills/{skill_id}
    │                   │   │   │       # DELETE /v1/users/{id}/skills/{skill_id}
    │                   │   │   │       # BE: search-be/taxonomy
    │                   │   │   │       # GET /v1/taxonomy/skills (autocomplete)
    │                   │   │   ├── verification/
    │                   │   │   │   ├── identity/
    │                   │   │   │   │   └── page.tsx  # ID verification
    │                   │   │   │   │       # - Upload ID document
    │                   │   │   │   │       # - Selfie verification
    │                   │   │   │   │       # - Address proof
    │                   │   │   │   │       # BE: users-be/verification/kyc
    │                   │   │   │   │       # POST /v1/users/{id}/kyc/submit
    │                   │   │   │   │       # BE: storage-be/uploads
    │                   │   │   │   │       # BE: admin-be/kyc_case (creates case)
    │                   │   │   │   ├── phone/
    │                   │   │   │   │   └── page.tsx  # Phone verification
    │                   │   │   │   │       # - Enter phone number
    │                   │   │   │   │       # - Receive OTP
    │                   │   │   │   │       # - Verify OTP
    │                   │   │   │   │       # BE: users-be/verification
    │                   │   │   │   │       # POST /v1/users/{id}/verify-phone/send
    │                   │   │   │   │       # POST /v1/users/{id}/verify-phone/verify
    │                   │   │   │   └── page.tsx  # Verification status
    │                   │   │   │       # - Email verified
    │                   │   │   │       # - Phone verified
    │                   │   │   │       # - ID verification status
    │                   │   │   │       # - Payment method verified
    │                   │   │   │       # BE: users-be/verification
    │                   │   │   │       # GET /v1/users/{id}/verification-status
    │                   │   │   └── page.tsx  # Profile overview / public view
    │                   │   │       # - Profile header (photo, name, title, location)
    │                   │   │       # - Stats (rating, jobs completed, earnings)
    │                   │   │       # - Skills showcase
    │                   │   │       # - Portfolio highlights
    │                   │   │       # - Recent reviews
    │                   │   │       # - Availability calendar
    │                   │   │       # BE: users-be/profile
    │                   │   │       # GET /v1/users/{id}/profile
    │                   │   │       # BE: users-be/capabilities
    │                   │   │       # GET /v1/users/{id}/skills
    │                   │   │       # BE: users-be/portfolio
    │                   │   │       # GET /v1/users/{id}/portfolio
    │                   │   │       # BE: reviews-be
    │                   │   │       # GET /v1/reviews?user_id={id}&limit=5
    │                   │   │       # BE: users-be/availability
    │                   │   │       # GET /v1/users/{id}/availability
    │                   │   ├── proposals/  # Proposals management
    │                   │   │   ├── [proposalId]/
    │                   │   │   │   ├── bidding/
    │                   │   │   │   │   └── page.tsx  # Bid status for proposal
    │                   │   │   │   │       # - Your current bid
    │                   │   │   │   │       # - Current lowest bid
    │                   │   │   │   │       # - Update bid
    │                   │   │   │   │       # - Bid history
    │                   │   │   │   │       # BE: proposals-be/bidding
    │                   │   │   │   │       # GET /v1/proposals/{proposal_id}/bid
    │                   │   │   │   ├── edit/
    │                   │   │   │   │   └── page.tsx  # Edit proposal
    │                   │   │   │   │       # - Only if status = DRAFT or PENDING
    │                   │   │   │   │       # - Update cover letter, rate, timeline
    │                   │   │   │   │       # BE: proposals-be
    │                   │   │   │   │       # PUT /v1/proposals/{proposal_id}
    │                   │   │   │   ├── withdraw/
    │                   │   │   │   │   └── page.tsx  # Withdraw proposal
    │                   │   │   │   │       # - Confirmation dialog
    │                   │   │   │   │       # - Reason for withdrawal
    │                   │   │   │   │       # - Connects refund info
    │                   │   │   │   │       # BE: proposals-be
    │                   │   │   │   │       # POST /v1/proposals/{proposal_id}/withdraw
    │                   │   │   │   │       # Publishes: ProposalWithdrawn event
    │                   │   │   │   └── page.tsx  # Proposal detail
    │                   │   │   │       # - View submitted proposal
    │                   │   │   │       # - Proposal status
    │                   │   │   │       # - Client messages/feedback
    │                   │   │   │       # - Withdraw option (if pending)
    │                   │   │   │       # BE: proposals-be
    │                   │   │   │       # GET /v1/proposals/{proposal_id}
    │                   │   │   ├── accepted/
    │                   │   │   │   └── page.tsx  # Accepted proposals
    │                   │   │   │       # BE: proposals-be
    │                   │   │   │       # GET /v1/proposals/my-proposals?status=accepted
    │                   │   │   ├── analytics/
    │                   │   │   │   └── page.tsx  # Proposal analytics (freelancer)
    │                   │   │   │       # - Total proposals submitted
    │                   │   │   │       # - Acceptance rate
    │                   │   │   │       # - Average response time
    │                   │   │   │       # - Connects spent
    │                   │   │   │       # BE: proposals-be/analytics
    │                   │   │   │       # GET /v1/proposals/analytics
    │                   │   │   ├── pending/
    │                   │   │   │   └── page.tsx  # Pending proposals
    │                   │   │   │       # BE: proposals-be
    │                   │   │   │       # GET /v1/proposals/my-proposals?status=pending
    │                   │   │   ├── rejected/
    │                   │   │   │   └── page.tsx  # Rejected proposals
    │                   │   │   │       # BE: proposals-be
    │                   │   │   │       # GET /v1/proposals/my-proposals?status=rejected
    │                   │   │   ├── submit/
    │                   │   │   │   └── [jobId]/
    │                   │   │   │       └── page.tsx  # Submit proposal (freelancer)
    │                   │   │   │           # - Cover letter (required)
    │                   │   │   │           # - Proposed rate/budget
    │                   │   │   │           # - Proposed timeline
    │                   │   │   │           # - Answer screening questions
    │                   │   │   │           # - Attachments (portfolio samples)
    │                   │   │   │           # - Milestones (for fixed-price)
    │                   │   │   │           # - Terms acceptance
    │                   │   │   │           # - Connects deduction warning
    │                   │   │   │           # BE: proposals-be
    │                   │   │   │           # POST /v1/proposals
    │                   │   │   │           # Body: { job_id, cover_letter, rate, timeline, ... }
    │                   │   │   │           # BE: subscriptions-be/connects
    │                   │   │   │           # POST /v1/connects/deduct
    │                   │   │   │           # BE: storage-be/uploads
    │                   │   │   │           # Publishes: ProposalSubmitted event
    │                   │   │   ├── templates/
    │                   │   │   │   ├── [templateId]/
    │                   │   │   │   │   └── edit/
    │                   │   │   │   │       └── page.tsx  # Edit template
    │                   │   │   │   │           # BE: proposals-be/templates
    │                   │   │   │   │           # PUT /v1/proposals/templates/{template_id}
    │                   │   │   │   │           # DELETE /v1/proposals/templates/{template_id}
    │                   │   │   │   ├── create/
    │                   │   │   │   │   └── page.tsx  # Create template
    │                   │   │   │   │       # - Template name
    │                   │   │   │   │       # - Cover letter template
    │                   │   │   │   │       # - Default rate/terms
    │                   │   │   │   │       # BE: proposals-be/templates
    │                   │   │   │   │       # POST /v1/proposals/templates
    │                   │   │   │   └── page.tsx  # Proposal templates
    │                   │   │   │       # - List saved templates
    │                   │   │   │       # - Create new template
    │                   │   │   │       # BE: proposals-be/templates
    │                   │   │   │       # GET /v1/proposals/templates
    │                   │   │   └── page.tsx  # Proposals list
    │                   │   │       # Freelancer: Submitted proposals
    │                   │   │       # Client: Received proposals (redirect to jobs)
    │                   │   │       # BE: proposals-be
    │                   │   │       # GET /v1/proposals/my-proposals
    │                   │   ├── reviews/  # Reviews & ratings
    │                   │   │   ├── analytics/
    │                   │   │   │   └── page.tsx  # Review analytics
    │                   │   │   │       # - Rating trends
    │                   │   │   │       # - Response rate
    │                   │   │   │       # - Sentiment analysis
    │                   │   │   │       # BE: reviews-be/review
    │                   │   │   │       # GET /v1/users/me/reviews/analytics
    │                   │   │   ├── badges/
    │                   │   │   │   └── page.tsx  # Badges & achievements
    │                   │   │   │       # - Earned badges
    │                   │   │   │       # - Badge criteria
    │                   │   │   │       # - Progress towards badges
    │                   │   │   │       # BE: reviews-be/badges
    │                   │   │   │       # GET /v1/reviews/badges?user_id={current_user}
    │                   │   │   ├── create/
    │                   │   │   │   ├── [contractId]/
    │                   │   │   │   │   └── page.tsx  # Create review
    │                   │   │   │   │       # - Rating (1-5 stars)
    │                   │   │   │   │       # - Multiple criteria ratings:
    │                   │   │   │   │       # - Quality of work
    │                   │   │   │   │       # - Communication
    │                   │   │   │   │       # - Professionalism
    │                   │   │   │   │       # - Deadlines
    │                   │   │   │   │       # - Written feedback (required)
    │                   │   │   │   │       # - Recommend to others?
    │                   │   │   │   │       # - Skills demonstrated
    │                   │   │   │   │       # - Make public/private
    │                   │   │   │   │       # BE: reviews-be/reviews
    │                   │   │   │   │       # POST /v1/reviews
    │                   │   │   │   │       # Body: { contract_id, rating, criteria_ratings, feedback, ... }
    │                   │   │   │   │       # Publishes: ReviewSubmitted event
    │                   │   │   │   │       # Triggers: Reputation update (users-be)
    │                   │   │   │   └── pending/
    │                   │   │   │       └── page.tsx  # Pending reviews
    │                   │   │   │           # - Contracts awaiting review
    │                   │   │   │           # - Reminders
    │                   │   │   │           # BE: reviews-be/reviews
    │                   │   │   │           # GET /v1/reviews/pending
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
    │                   │   │   │   └── page.tsx  # Reviews given list
    │                   │   │   │       # BE: reviews-be/reviews
    │                   │   │   │       # GET /v1/reviews?user_id={current_user}&type=given
    │                   │   │   │       # Given reviews list
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
    │                   │   │   ├── received/
    │                   │   │   │   ├── [reviewId]/
    │                   │   │   │   │   ├── respond/
    │                   │   │   │   │   │   └── page.tsx  # Respond to review
    │                   │   │   │   │   │       # - Public response
    │                   │   │   │   │   │       # - Character limit
    │                   │   │   │   │   │       # BE: reviews-be/reviews
    │                   │   │   │   │   │       # POST /v1/reviews/{review_id}/respond
    │                   │   │   │   │   │       # Publishes: ReviewResponded event
    │                   │   │   │   │   │       # BE: reviews-be/review
    │                   │   │   │   │   │       # POST /v1/reviews/{review_id}/response
    │                   │   │   │   │   └── page.tsx  # Review detail
    │                   │   │   │   │       # - Full review content
    │                   │   │   │   │       # - Reviewer info
    │                   │   │   │   │       # - Related contract
    │                   │   │   │   │       # - Response option
    │                   │   │   │   │       # BE: reviews-be/reviews
    │                   │   │   │   │       # GET /v1/reviews/{review_id}
    │                   │   │   │   │       # Review details
    │                   │   │   │   │       # BE: reviews-be/review
    │                   │   │   │   └── page.tsx  # Reviews received list
    │                   │   │   │       # - All reviews received
    │                   │   │   │       # - Filter by rating
    │                   │   │   │       # - Filter by contract
    │                   │   │   │       # BE: reviews-be/reviews
    │                   │   │   │       # GET /v1/reviews?user_id={current_user}&type=received
    │                   │   │   │       # Received reviews list
    │                   │   │   │       # BE: reviews-be/review
    │                   │   │   │       # GET /v1/users/me/reviews/received
    │                   │   │   ├── stats/
    │                   │   │   │   └── page.tsx  # Detailed statistics
    │                   │   │   │       # - Rating breakdown
    │                   │   │   │       # - Review trends over time
    │                   │   │   │       # - Category-specific ratings
    │                   │   │   │       # - Comparison to platform average
    │                   │   │   │       # BE: reviews-be/stats
    │                   │   │   │       # GET /v1/reviews/stats/detailed?user_id={current_user}
    │                   │   │   └── page.tsx  # Reviews overview
    │                   │   │       # - Reviews received
    │                   │   │       # - Reviews given
    │                   │   │       # - Overall rating stats
    │                   │   │       # - Badges earned
    │                   │   │       # BE: reviews-be/reviews
    │                   │   │       # GET /v1/reviews?user_id={current_user}
    │                   │   │       # BE: reviews-be/stats
    │                   │   │       # GET /v1/reviews/stats?user_id={current_user}
    │                   │   ├── search/  # Advanced search functionality
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
    │                   │   │   │       # - Experience level
    │                   │   │   │       # - Hourly rate range
    │                   │   │   │       # - Location
    │                   │   │   │       # - Availability
    │                   │   │   │       # - Rating
    │                   │   │   │       # - Portfolio keywords
    │                   │   │   │       # BE: search-be/query
    │                   │   │   │       # POST /v1/search/freelancers
    │                   │   │   │       # Body: { query, filters: {...}, sort, page }
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
    │                   │   │   ├── portfolio/
    │                   │   │   │   └── page.tsx  # Search portfolios
    │                   │   │   │       # - Search by project keywords
    │                   │   │   │       # - Filter by skills used
    │                   │   │   │       # - Filter by industry
    │                   │   │   │       # BE: search-be/portfolio
    │                   │   │   │       # POST /v1/search/portfolios
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
    │                   │   │   │   └── page.tsx  # Saved searches list (may be in combined, ensuring here)
    │                   │   │   ├── saved-searches/
    │                   │   │   │   ├── [searchId]/
    │                   │   │   │   │   ├── edit/
    │                   │   │   │   │   │   └── page.tsx  # Edit saved search
    │                   │   │   │   │   │       # BE: search-be/saved_search
    │                   │   │   │   │   │       # PUT /v1/saved-searches/{search_id}
    │                   │   │   │   │   └── page.tsx  # Execute saved search
    │                   │   │   │   │       # BE: search-be/saved_search
    │                   │   │   │   │       # GET /v1/saved-searches/{search_id}/results
    │                   │   │   │   ├── create/
    │                   │   │   │   │   └── page.tsx  # Create saved search
    │                   │   │   │   │       # - Name the search
    │                   │   │   │   │       # - Set alert frequency
    │                   │   │   │   │       # - Save filters
    │                   │   │   │   │       # BE: search-be/saved_search
    │                   │   │   │   │       # POST /v1/saved-searches
    │                   │   │   │   └── page.tsx  # Saved searches list
    │                   │   │   │       # - List all saved searches
    │                   │   │   │       # - Email alerts toggle
    │                   │   │   │       # - Edit search
    │                   │   │   │       # - Delete search
    │                   │   │   │       # BE: search-be/saved_search
    │                   │   │   │       # GET /v1/saved-searches
    │                   │   │   └── trending/
    │                   │   │       └── page.tsx  # Trending searches and jobs
    │                   │   │           # BE: search-be/trending
    │                   │   │           # GET /v1/trending/jobs
    │                   │   │           # GET /v1/trending/skills
    │                   │   ├── settings/  # User settings
    │                   │   │   ├── account/
    │                   │   │   │   ├── delete/
    │                   │   │   │   │   └── page.tsx  # Delete account
    │                   │   │   │   │       # - Reason for deletion
    │                   │   │   │   │       # - Data export option
    │                   │   │   │   │       # - Confirmation (password + checkbox)
    │                   │   │   │   │       # - GDPR compliance
    │                   │   │   │   │       # BE: users-be/account
    │                   │   │   │   │       # POST /v1/users/{id}/delete-account
    │                   │   │   │   │       # Publishes: AccountDeleted event
    │                   │   │   │   ├── email/
    │                   │   │   │   │   └── change/
    │                   │   │   │   │       └── page.tsx  # Change email
    │                   │   │   │   │           # - New email input
    │                   │   │   │   │           # - Password confirmation
    │                   │   │   │   │           # - Verification email sent
    │                   │   │   │   │           # BE: users-be/user
    │                   │   │   │   │           # POST /v1/users/{id}/change-email
    │                   │   │   │   ├── phone/
    │                   │   │   │   │   └── change/
    │                   │   │   │   │       └── page.tsx  # Change phone number
    │                   │   │   │   │           # - New phone input
    │                   │   │   │   │           # - OTP verification
    │                   │   │   │   │           # BE: users-be/user
    │                   │   │   │   │           # POST /v1/users/{id}/change-phone
    │                   │   │   │   └── username/
    │                   │   │   │       └── change/
    │                   │   │   │           └── page.tsx  # Change username
    │                   │   │   │               # - New username
    │                   │   │   │               # - Availability check
    │                   │   │   │               # BE: users-be/user
    │                   │   │   │               # POST /v1/users/{id}/change-username
    │                   │   │   ├── developer/
    │                   │   │   │   ├── api-keys/
    │                   │   │   │   │   ├── create/
    │                   │   │   │   │   │   └── page.tsx  # Create API key
    │                   │   │   │   │   │       # - Key name
    │                   │   │   │   │   │       # - Permissions/scopes
    │                   │   │   │   │   │       # - Expiration
    │                   │   │   │   │   │       # BE: users-be/developer
    │                   │   │   │   │   │       # POST /v1/users/{id}/api-keys
    │                   │   │   │   │   └── page.tsx  # API keys list
    │                   │   │   │   │       # - Active keys
    │                   │   │   │   │       # - Create new key
    │                   │   │   │   │       # - Revoke key
    │                   │   │   │   │       # BE: users-be/developer
    │                   │   │   │   │       # GET /v1/users/{id}/api-keys
    │                   │   │   │   ├── oauth-apps/
    │                   │   │   │   │   ├── [appId]/
    │                   │   │   │   │   │   └── page.tsx  # OAuth app detail
    │                   │   │   │   │   │       # - Client ID/secret
    │                   │   │   │   │   │       # - Edit/delete
    │                   │   │   │   │   │       # - Usage stats
    │                   │   │   │   │   │       # BE: users-be/developer
    │                   │   │   │   │   │       # GET /v1/users/{id}/oauth-apps/{app_id}
    │                   │   │   │   │   └── create/
    │                   │   │   │   │       └── page.tsx  # Create OAuth app
    │                   │   │   │   │           # - App name
    │                   │   │   │   │           # - Redirect URIs
    │                   │   │   │   │           # - Scopes
    │                   │   │   │   │           # BE: users-be/developer
    │                   │   │   │   │           # POST /v1/users/{id}/oauth-apps
    │                   │   │   │   └── page.tsx  # Developer settings
    │                   │   │   │       # - API keys
    │                   │   │   │       # - OAuth applications
    │                   │   │   │       # - API usage stats
    │                   │   │   │       # BE: users-be/developer
    │                   │   │   │       # GET /v1/users/{id}/developer
    │                   │   │   ├── integrations/
    │                   │   │   │   ├── calendar/
    │                   │   │   │   │   └── page.tsx  # Calendar integration
    │                   │   │   │   │       # - Google Calendar
    │                   │   │   │   │       # - Outlook Calendar
    │                   │   │   │   │       # - Sync settings
    │                   │   │   │   │       # BE: users-be/integrations
    │                   │   │   │   │       # POST /v1/users/{id}/integrations/calendar
    │                   │   │   │   ├── slack/
    │                   │   │   │   │   └── page.tsx  # Slack integration
    │                   │   │   │   │       # - Connect Slack workspace
    │                   │   │   │   │       # - Notification channels
    │                   │   │   │   │       # BE: users-be/integrations
    │                   │   │   │   │       # POST /v1/users/{id}/integrations/slack
    │                   │   │   │   └── webhooks/
    │                   │   │   │       ├── [webhookId]/
    │                   │   │   │       │   └── page.tsx  # Webhook detail
    │                   │   │   │       │       # - Delivery logs
    │                   │   │   │       │       # - Test webhook
    │                   │   │   │       │       # - Edit/delete
    │                   │   │   │       │       # BE: users-be/integrations
    │                   │   │   │       │       # GET /v1/users/{id}/webhooks/{webhook_id}
    │                   │   │   │       │       # PUT /v1/users/{id}/webhooks/{webhook_id}
    │                   │   │   │       │       # DELETE /v1/users/{id}/webhooks/{webhook_id}
    │                   │   │   │       └── create/
    │                   │   │   │           └── page.tsx  # Create webhook
    │                   │   │   │               # - Webhook URL
    │                   │   │   │               # - Secret key
    │                   │   │   │               # - Events to subscribe
    │                   │   │   │               # BE: users-be/integrations
    │                   │   │   │               # POST /v1/users/{id}/webhooks
    │                   │   │   ├── notifications/
    │                   │   │   │   ├── digest/
    │                   │   │   │   │   └── page.tsx  # Email digest settings
    │                   │   │   │   │       # - Daily/weekly/monthly
    │                   │   │   │   │       # - Content preferences
    │                   │   │   │   │       # BE: communications-be/preferences
    │                   │   │   │   │       # PUT /v1/notifications/digest-preferences
    │                   │   │   │   ├── email/
    │                   │   │   │   │   └── page.tsx  # Email notification settings
    │                   │   │   │   │       # - Per-category toggles
    │                   │   │   │   │       # - Digest preferences
    │                   │   │   │   │       # BE: communications-be/preferences
    │                   │   │   │   │       # PUT /v1/notifications/email-preferences
    │                   │   │   │   └── push/
    │                   │   │   │       └── page.tsx  # Push notification settings
    │                   │   │   │           # - Device management
    │                   │   │   │           # - Per-category toggles
    │                   │   │   │           # BE: communications-be/preferences
    │                   │   │   │           # PUT /v1/notifications/push-preferences
    │                   │   │   ├── preferences/
    │                   │   │   │   ├── accessibility/
    │                   │   │   │   │   └── page.tsx  # Accessibility settings
    │                   │   │   │   │       # - Screen reader optimizations
    │                   │   │   │   │       # - High contrast mode
    │                   │   │   │   │       # - Font size
    │                   │   │   │   │       # - Keyboard shortcuts
    │                   │   │   │   │       # - Motion preferences
    │                   │   │   │   │       # BE: users-be/preferences
    │                   │   │   │   │       # PATCH /v1/users/{id}/preferences/accessibility
    │                   │   │   │   └── language/
    │                   │   │   │       └── page.tsx  # Language settings
    │                   │   │   │           # - Interface language
    │                   │   │   │           # - Content languages
    │                   │   │   │           # BE: users-be/preferences
    │                   │   │   │           # PATCH /v1/users/{id}/preferences/language
    │                   │   │   ├── privacy/
    │                   │   │   │   ├── blocked-users/
    │                   │   │   │   │   ├── add/
    │                   │   │   │   │   │   └── page.tsx  # Block user
    │                   │   │   │   │   │       # BE: users-be/privacy
    │                   │   │   │   │   │       # POST /v1/users/{id}/blocked-users
    │                   │   │   │   │   └── page.tsx  # Blocked users list
    │                   │   │   │   │       # - List blocked users
    │                   │   │   │   │       # - Unblock option
    │                   │   │   │   │       # BE: users-be/privacy
    │                   │   │   │   │       # GET /v1/users/{id}/blocked-users
    │                   │   │   │   │       # DELETE /v1/users/{id}/blocked-users/{user_id}
    │                   │   │   │   └── data-export/
    │                   │   │   │       ├── request/
    │                   │   │   │       │   └── page.tsx  # Request new data export
    │                   │   │   │       │       # - Select data categories
    │                   │   │   │       │       # - Format (JSON, CSV, PDF)
    │                   │   │   │       │       # BE: users-be/privacy
    │                   │   │   │       │       # POST /v1/users/{id}/data-export/request
    │                   │   │   │       └── page.tsx  # Data export (GDPR)
    │                   │   │   │           # - Request data export
    │                   │   │   │           # - Export history
    │                   │   │   │           # - Download exports
    │                   │   │   │           # BE: users-be/privacy
    │                   │   │   │           # POST /v1/users/{id}/data-export
    │                   │   │   │           # GET /v1/users/{id}/data-exports
    │                   │   │   ├── security/
    │                   │   │   │   ├── login-history/
    │                   │   │   │   │   └── page.tsx  # Login history
    │                   │   │   │   │       # - Recent logins
    │                   │   │   │   │       # - Failed attempts
    │                   │   │   │   │       # - Device/location info
    │                   │   │   │   │       # BE: users-be/security/audit
    │                   │   │   │   │       # GET /v1/users/{id}/login-history
    │                   │   │   │   ├── password/
    │                   │   │   │   │   └── change/
    │                   │   │   │   │       └── page.tsx  # Change password
    │                   │   │   │   │           # - Current password
    │                   │   │   │   │           # - New password
    │                   │   │   │   │           # - Password strength meter
    │                   │   │   │   │           # BE: users-be/security
    │                   │   │   │   │           # POST /v1/users/{id}/change-password
    │                   │   │   │   ├── sessions/
    │                   │   │   │   │   ├── revoke-all/
    │                   │   │   │   │   │   └── page.tsx  # Revoke all sessions (except current)
    │                   │   │   │   │   │       # BE: users-be/security/session
    │                   │   │   │   │   │       # POST /v1/users/{id}/sessions/revoke-all
    │                   │   │   │   │   └── page.tsx  # Active sessions
    │                   │   │   │   │       # - List all active sessions
    │                   │   │   │   │       # - Device info
    │                   │   │   │   │       # - Location
    │                   │   │   │   │       # - Last active
    │                   │   │   │   │       # - Revoke session
    │                   │   │   │   │       # BE: users-be/security/session
    │                   │   │   │   │       # GET /v1/users/{id}/sessions
    │                   │   │   │   │       # DELETE /v1/users/{id}/sessions/{session_id}
    │                   │   │   │   └── two-factor/
    │                   │   │   │       ├── disable/
    │                   │   │   │       │   └── page.tsx  # Disable 2FA
    │                   │   │   │       │       # - Password confirmation
    │                   │   │   │       │       # - 2FA code verification
    │                   │   │   │       │       # BE: users-be/security/mfa
    │                   │   │   │       │       # POST /v1/users/{id}/mfa/disable
    │                   │   │   │       └── enable/
    │                   │   │   │           └── page.tsx  # Enable 2FA
    │                   │   │   │               # - QR code scan
    │                   │   │   │               # - Verify setup
    │                   │   │   │               # - Save backup codes
    │                   │   │   │               # BE: users-be/security/mfa
    │                   │   │   │               # POST /v1/users/{id}/mfa/enable
    │                   │   │   └── page.tsx  # Settings overview
    │                   │   │       # - Quick access to all settings
    │                   │   │       # - Profile completion indicator
    │                   │   │       # - Recently changed settings
    │                   │   │       # BE: users-be/preferences
    │                   │   │       # GET /v1/users/{id}/preferences
    │                   │   ├── sourcing/
    │                   │   │   ├── campaigns/
    │                   │   │   │   ├── [campaignId]/
    │                   │   │   │   │   ├── analytics/
    │                   │   │   │   │   │   └── page.tsx  # Campaign analytics
    │                   │   │   │   │   │       # - Reach
    │                   │   │   │   │   │       # - Engagement
    │                   │   │   │   │   │       # - Conversions
    │                   │   │   │   │   │       # BE: jobs-be/analytics
    │                   │   │   │   │   │       # GET /v1/sourcing/campaigns/{campaign_id}/analytics
    │                   │   │   │   │   ├── edit/
    │                   │   │   │   │   │   └── page.tsx  # Edit sourcing campaign
    │                   │   │   │   │   │       # BE: jobs-be/campaign (if exists) or jobs-be/job
    │                   │   │   │   │   │       # PUT /v1/sourcing/campaigns/{campaign_id}
    │                   │   │   │   │   └── page.tsx  # Campaign detail
    │                   │   │   │   │       # BE: jobs-be/campaign
    │                   │   │   │   │       # GET /v1/sourcing/campaigns/{campaign_id}
    │                   │   │   │   ├── create/
    │                   │   │   │   │   └── page.tsx  # Create sourcing campaign
    │                   │   │   │   │       # - Target criteria
    │                   │   │   │   │       # - Budget allocation
    │                   │   │   │   │       # - Messaging
    │                   │   │   │   │       # BE: jobs-be/campaign
    │                   │   │   │   │       # POST /v1/sourcing/campaigns
    │                   │   │   │   └── page.tsx  # Campaigns list
    │                   │   │   │       # BE: jobs-be/campaign
    │                   │   │   │       # GET /v1/sourcing/campaigns
    │                   │   │   ├── invitations/
    │                   │   │   │   ├── [invitationId]/
    │                   │   │   │   │   └── page.tsx  # Invitation detail
    │                   │   │   │   │       # - View status
    │                   │   │   │   │       # - Resend invitation
    │                   │   │   │   │       # BE: communications-be/invitation (if exists) or jobs-be/job
    │                   │   │   │   │       # GET /v1/sourcing/invitations/{invitation_id}
    │                   │   │   │   └── page.tsx  # Sent invitations
    │                   │   │   │       # - Invitation history
    │                   │   │   │       # - Response rates
    │                   │   │   │       # BE: communications-be/invitation
    │                   │   │   │       # GET /v1/sourcing/invitations
    │                   │   │   └── talent-pools/
    │                   │   │       ├── [poolId]/
    │                   │   │       │   ├── edit/
    │                   │   │       │   │   └── page.tsx  # Edit talent pool
    │                   │   │       │   │       # BE: users-be/talent_pool (if exists) or search-be/saved-search
    │                   │   │       │   │       # PUT /v1/sourcing/talent-pools/{pool_id}
    │                   │   │       │   ├── members/
    │                   │   │       │   │   └── page.tsx  # Pool members
    │                   │   │       │   │       # - Add/remove members
    │                   │   │       │   │       # - Bulk actions
    │                   │   │       │   │       # BE: users-be/talent_pool
    │                   │   │       │   │       # GET /v1/sourcing/talent-pools/{pool_id}/members
    │                   │   │       │   └── page.tsx  # Talent pool detail
    │                   │   │       │       # BE: users-be/talent_pool
    │                   │   │       │       # GET /v1/sourcing/talent-pools/{pool_id}
    │                   │   │       ├── create/
    │                   │   │       │   └── page.tsx  # Create talent pool
    │                   │   │       │       # BE: users-be/talent_pool
    │                   │   │       │       # POST /v1/sourcing/talent-pools
    │                   │   │       └── page.tsx  # Talent pools list
    │                   │   │           # BE: users-be/talent_pool
    │                   │   │           # GET /v1/sourcing/talent-pools
    │                   │   ├── spend-analytics/
    │                   │   │   ├── by-category/
    │                   │   │   │   └── page.tsx  # Spending by category
    │                   │   │   │       # - Job category breakdown
    │                   │   │   │       # - Trend analysis
    │                   │   │   │       # BE: financial-be/analytics (if exists) or financial-be/invoice
    │                   │   │   │       # GET /v1/analytics/spend/by-category
    │                   │   │   ├── by-department/
    │                   │   │   │   └── page.tsx  # Spending by department
    │                   │   │   │       # - Cost center breakdown
    │                   │   │   │       # - Budget vs actual
    │                   │   │   │       # BE: financial-be/budget, financial-be/invoice
    │                   │   │   │       # GET /v1/analytics/spend/by-department
    │                   │   │   ├── by-vendor/
    │                   │   │   │   └── page.tsx  # Spending by vendor
    │                   │   │   │       # - Top vendors
    │                   │   │   │       # - Vendor concentration risk
    │                   │   │   │       # BE: financial-be/invoice, users-be/profile
    │                   │   │   │       # GET /v1/analytics/spend/by-vendor
    │                   │   │   ├── forecasting/
    │                   │   │   │   └── page.tsx  # Spend forecasting
    │                   │   │   │       # - Projected spend
    │                   │   │   │       # - Budget burn rate
    │                   │   │   │       # - Alerts for overages
    │                   │   │   │       # BE: financial-be/forecast (if exists) or financial-be/analytics
    │                   │   │   │       # GET /v1/analytics/spend/forecast
    │                   │   │   └── page.tsx  # Spend analytics dashboard
    │                   │   │       # - Overview metrics
    │                   │   │       # - Charts & visualizations
    │                   │   │       # BE: financial-be/analytics
    │                   │   │       # GET /v1/analytics/spend
    │                   │   ├── subscription/  # Subscription management
    │                   │   │   ├── addons/
    │                   │   │   │   └── [addonId]/
    │                   │   │   │       ├── cancel/
    │                   │   │   │       │   └── page.tsx  # Cancel addon
    │                   │   │   │       │       # BE: subscriptions-be/addons
    │                   │   │   │       │       # DELETE /v1/subscriptions/{sub_id}/addons/{addon_id}
    │                   │   │   │       └── purchase/
    │                   │   │   │           └── page.tsx  # Purchase addon
    │                   │   │   │               # BE: subscriptions-be/addons
    │                   │   │   │               # POST /v1/subscriptions/{sub_id}/addons
    │                   │   │   ├── billing-history/
    │                   │   │   │   ├── [invoiceId]/
    │                   │   │   │   │   └── page.tsx  # Invoice detail
    │                   │   │   │   │       # - Invoice details
    │                   │   │   │   │       # - Download PDF
    │                   │   │   │   │       # BE: subscriptions-be/invoices
    │                   │   │   │   │       # GET /v1/invoices/{invoice_id}
    │                   │   │   │   │       # GET /v1/invoices/{invoice_id}/pdf
    │                   │   │   │   └── page.tsx  # Billing history
    │                   │   │   │       # - Past invoices
    │                   │   │   │       # - Payment history
    │                   │   │   │       # - Download invoices
    │                   │   │   │       # BE: subscriptions-be/invoices
    │                   │   │   │       # GET /v1/subscriptions/{sub_id}/invoices
    │                   │   │   ├── cancel/
    │                   │   │   │   └── page.tsx  # Cancel subscription
    │                   │   │   │       # - Cancellation reason
    │                   │   │   │       # - Feedback
    │                   │   │   │       # - Immediate vs. end of period
    │                   │   │   │       # - Refund eligibility
    │                   │   │   │       # - Data retention info
    │                   │   │   │       # BE: subscriptions-be/subscriptions
    │                   │   │   │       # POST /v1/subscriptions/{sub_id}/cancel
    │                   │   │   │       # Publishes: SubscriptionCancelled event
    │                   │   │   ├── connects/
    │                   │   │   │   ├── history/
    │                   │   │   │   │   └── page.tsx  # Connects usage history
    │                   │   │   │   │       # - Connects spent (proposals)
    │                   │   │   │   │       # - Connects added (purchases/plan)
    │                   │   │   │   │       # - Balance over time
    │                   │   │   │   │       # BE: subscriptions-be/connects
    │                   │   │   │   │       # GET /v1/connects/transactions
    │                   │   │   │   ├── purchase/
    │                   │   │   │   │   └── page.tsx  # Purchase connects
    │                   │   │   │   │       # - Select package
    │                   │   │   │   │       # - Pricing options
    │                   │   │   │   │       # - Bulk discounts
    │                   │   │   │   │       # - Payment method
    │                   │   │   │   │       # BE: subscriptions-be/connects
    │                   │   │   │   │       # POST /v1/connects/purchase
    │                   │   │   │   │       # Body: { package_id, quantity }
    │                   │   │   │   │       # Publishes: ConnectsPurchased event
    │                   │   │   │   └── page.tsx  # Connects overview
    │                   │   │   │       # - Current balance
    │                   │   │   │       # - Usage history
    │                   │   │   │       # - Included in plan
    │                   │   │   │       # - Purchase more
    │                   │   │   │       # BE: subscriptions-be/connects
    │                   │   │   │       # GET /v1/connects/balance
    │                   │   │   │       # GET /v1/connects/usage
    │                   │   │   ├── downgrade/
    │                   │   │   │   └── page.tsx  # Downgrade plan
    │                   │   │   │       # - Select new plan
    │                   │   │   │       # - Effective date (end of billing period)
    │                   │   │   │       # - Feature comparison
    │                   │   │   │       # - Confirm downgrade
    │                   │   │   │       # BE: subscriptions-be/subscriptions
    │                   │   │   │       # POST /v1/subscriptions/downgrade
    │                   │   │   │       # Publishes: SubscriptionDowngraded event
    │                   │   │   ├── plans/
    │                   │   │   │   ├── compare/
    │                   │   │   │   │   └── page.tsx  # Plan comparison
    │                   │   │   │   │       # - Side-by-side comparison
    │                   │   │   │   │       # - Feature highlights
    │                   │   │   │   │       # BE: subscriptions-be/plans
    │                   │   │   │   │       # GET /v1/plans/compare
    │                   │   │   │   └── page.tsx  # Available plans
    │                   │   │   │       # - Plan comparison table
    │                   │   │   │       # - Feature matrix
    │                   │   │   │       # - Pricing tiers
    │                   │   │   │       # - Billing periods (monthly, annual)
    │                   │   │   │       # - Free trial info
    │                   │   │   │       # BE: subscriptions-be/plans
    │                   │   │   │       # GET /v1/plans
    │                   │   │   ├── reactivate/
    │                   │   │   │   └── page.tsx  # Reactivate subscription
    │                   │   │   │       # - Select plan
    │                   │   │   │       # - Payment method
    │                   │   │   │       # BE: subscriptions-be/subscriptions
    │                   │   │   │       # POST /v1/subscriptions/reactivate
    │                   │   │   ├── trial/
    │                   │   │   │   ├── convert/
    │                   │   │   │   │   └── page.tsx  # Convert trial to paid
    │                   │   │   │   │       # - Select plan
    │                   │   │   │   │       # - Payment method
    │                   │   │   │   │       # - Apply promotion code
    │                   │   │   │   │       # BE: subscriptions-be/trials
    │                   │   │   │   │       # POST /v1/trials/{trial_id}/convert
    │                   │   │   │   └── page.tsx  # Trial status
    │                   │   │   │       # - Trial end date
    │                   │   │   │       # - Days remaining
    │                   │   │   │       # - Trial features
    │                   │   │   │       # - Upgrade prompt
    │                   │   │   │       # BE: subscriptions-be/trials
    │                   │   │   │       # GET /v1/trials/current
    │                   │   │   ├── upgrade/
    │                   │   │   │   ├── confirm/
    │                   │   │   │   │   └── page.tsx  # Confirm upgrade
    │                   │   │   │   │       # - Upgrade summary
    │                   │   │   │   │       # - Payment processing
    │                   │   │   │   │       # BE: subscriptions-be/subscriptions
    │                   │   │   │   │       # POST /v1/subscriptions/{sub_id}/confirm-upgrade
    │                   │   │   │   └── page.tsx  # Upgrade plan
    │                   │   │   │       # - Select new plan
    │                   │   │   │       # - Billing period
    │                   │   │   │       # - Proration calculation
    │                   │   │   │       # - Payment method
    │                   │   │   │       # - Confirm upgrade
    │                   │   │   │       # BE: subscriptions-be/subscriptions
    │                   │   │   │       # POST /v1/subscriptions/upgrade
    │                   │   │   │       # Body: { plan_id, billing_period }
    │                   │   │   │       # Publishes: SubscriptionUpgraded event
    │                   │   │   ├── usage/
    │                   │   │   │   └── page.tsx  # Usage statistics
    │                   │   │   │       # - Jobs posted
    │                   │   │   │       # - Proposals submitted
    │                   │   │   │       # - Storage used
    │                   │   │   │       # - API calls
    │                   │   │   │       # - Usage vs. limits
    │                   │   │   │       # BE: subscriptions-be/usage
    │                   │   │   │       # GET /v1/subscriptions/usage
    │                   │   │   └── page.tsx  # Subscription overview
    │                   │   │       # - Current plan details
    │                   │   │       # - Usage stats
    │                   │   │       # - Next billing date
    │                   │   │       # - Upgrade/downgrade options
    │                   │   │       # - Connects balance
    │                   │   │       # BE: subscriptions-be/subscriptions
    │                   │   │       # GET /v1/subscriptions/current
    │                   │   │       # BE: subscriptions-be/entitlements
    │                   │   │       # GET /v1/entitlements
    │                   │   │       # BE: subscriptions-be/connects
    │                   │   │       # GET /v1/connects/balance
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
    │                   │   │   ├── approve/
    │                   │   │   │   └── page.tsx  # Timesheets pending approval (client)
    │                   │   │   │       # BE: contracts-be/timesheet
    │                   │   │   │       # GET /v1/timesheets/pending-approval
    │                   │   │   └── page.tsx  # All timesheets overview
    │                   │   │       # BE: contracts-be/timesheet
    │                   │   │       # GET /v1/timesheets
    │                   │   ├── vendor-management/
    │                   │   │   ├── blacklist/
    │                   │   │   │   ├── [userId]/
    │                   │   │   │   │   └── page.tsx  # Blacklist entry detail
    │                   │   │   │   │       # - Reason for blacklist
    │                   │   │   │   │       # - Remove option
    │                   │   │   │   │       # BE: users-be/org (blacklist subdomain)
    │                   │   │   │   │       # GET /v1/vendors/blacklist/{user_id}
    │                   │   │   │   │       # DELETE /v1/vendors/blacklist/{user_id}
    │                   │   │   │   └── page.tsx  # Blacklisted vendors
    │                   │   │   │       # BE: users-be/org
    │                   │   │   │       # GET /v1/vendors/blacklist
    │                   │   │   ├── compliance-docs/
    │                   │   │   │   ├── [vendorId]/
    │                   │   │   │   │   └── page.tsx  # Vendor compliance documents
    │                   │   │   │   │       # - W-9/W-8BEN forms
    │                   │   │   │   │       # - Insurance certificates
    │                   │   │   │   │       # - Background checks
    │                   │   │   │   │       # BE: admin-be/business_verification, storage-be/asset
    │                   │   │   │   │       # GET /v1/vendors/{vendor_id}/compliance-docs
    │                   │   │   │   └── page.tsx  # Compliance tracking
    │                   │   │   │       # - Expiring documents
    │                   │   │   │       # - Missing documents
    │                   │   │   │       # BE: admin-be/business_verification
    │                   │   │   │       # GET /v1/vendors/compliance-status
    │                   │   │   └── preferred/
    │                   │   │       ├── [vendorId]/
    │                   │   │       │   ├── history/
    │                   │   │       │   │   └── page.tsx  # Work history with vendor
    │                   │   │       │   │       # - Past contracts
    │                   │   │       │   │       # - Total spend
    │                   │   │       │   │       # - Reviews given
    │                   │   │       │   │       # BE: contracts-be/contract, financial-be/invoice
    │                   │   │       │   │       # GET /v1/vendors/{vendor_id}/history
    │                   │   │       │   ├── performance/
    │                   │   │       │   │   └── page.tsx  # Vendor performance metrics
    │                   │   │       │   │       # - Success rate
    │                   │   │       │   │       # - Average delivery time
    │                   │   │       │   │       # - Quality scores
    │                   │   │       │   │       # BE: users-be/org (vendor subdomain), contracts-be/contract
    │                   │   │       │   │       # GET /v1/vendors/{vendor_id}/performance
    │                   │   │       │   └── page.tsx  # Vendor detail
    │                   │   │       │       # BE: users-be/org (vendor subdomain)
    │                   │   │       │       # GET /v1/vendors/{vendor_id}
    │                   │   │       └── page.tsx  # Preferred vendors list
    │                   │   │           # - Star ratings
    │                   │   │           # - Quick invite
    │                   │   │           # BE: users-be/org
    │                   │   │           # GET /v1/vendors/preferred
    │                   │   ├── work-diary/
    │                   │   │   ├── [contractId]/
    │                   │   │   │   ├── calendar/
    │                   │   │   │   │   └── page.tsx  # Calendar view of work diary
    │                   │   │   │   │       # BE: contracts-be/work_diary
    │                   │   │   │   │       # GET /v1/contracts/{contract_id}/work-diary/calendar
    │                   │   │   │   ├── screenshots/
    │                   │   │   │   │   └── page.tsx  # Screenshots management
    │                   │   │   │   │       # - View all screenshots
    │                   │   │   │   │       # - Delete sensitive ones
    │                   │   │   │   │       # - Privacy settings
    │                   │   │   │   │       # BE: contracts-be/work_diary, storage-be/asset
    │                   │   │   │   │       # GET /v1/contracts/{contract_id}/work-diary/screenshots
    │                   │   │   │   └── page.tsx  # Work diary detail
    │                   │   │   │       # BE: contracts-be/work_diary
    │                   │   │   │       # GET /v1/contracts/{contract_id}/work-diary
    │                   │   │   └── page.tsx  # Work diary overview (all contracts)
    │                   │   │       # BE: contracts-be/work_diary
    │                   │   │       # GET /v1/work-diary
    │                   │   └── layout.tsx  # Dashboard layout
    │                   │       # - Authenticated header with user menu
    │                   │       # - Sidebar navigation (responsive)
    │                   │       # - Notification bell with unread count
    │                   │       # - Messages indicator
    │                   │       # - Breadcrumbs
    │                   │       # - Footer
    │                   │       # BE: communications-be
    │                   │       # GET /v1/notifications/unread-count
    │                   │       # WebSocket: ws://communications-be/v1/realtime
    │                   ├── (onboarding)/  # Onboarding flow (post-registration)
    │                   │   ├── client/  # Client onboarding
    │                   │   │   ├── billing/
    │                   │   │   │   └── page.tsx  # Billing setup
    │                   │   │   │       # - Billing address
    │                   │   │   │       # - Tax information
    │                   │   │   │       # - VAT/GST number
    │                   │   │   │       # - Payment method
    │                   │   │   │       # BE: financial-be/billing_profile
    │                   │   │   │       # POST /v1/billing-profiles
    │                   │   │   │       # BE: financial-be/payment_method
    │                   │   │   │       # POST /v1/payment-methods
    │                   │   │   ├── company/
    │                   │   │   │   └── page.tsx  # Company info
    │                   │   │   │       # - Company name
    │                   │   │   │       # - Industry
    │                   │   │   │       # - Company size
    │                   │   │   │       # - Website
    │                   │   │   │       # - Logo upload
    │                   │   │   │       # BE: users-be/organization
    │                   │   │   │       # POST /v1/organizations
    │                   │   │   │       # BE: storage-be/uploads
    │                   │   │   ├── preferences/
    │                   │   │   │   └── page.tsx  # Hiring preferences
    │                   │   │   │       # - Typical project types
    │                   │   │   │       # - Budget ranges
    │                   │   │   │       # - Notification settings
    │                   │   │   │       # BE: users-be/preferences
    │                   │   │   │       # POST /v1/users/{id}/preferences
    │                   │   │   │       # Publishes: ClientProfileCompleted event
    │                   │   │   ├── team/
    │                   │   │   │   └── page.tsx  # Team setup (optional)
    │                   │   │   │       # - Invite team members
    │                   │   │   │       # - Assign roles
    │                   │   │   │       # BE: users-be/team
    │                   │   │   │       # POST /v1/organizations/{id}/members
    │                   │   │   │       # BE: communications-be
    │                   │   │   │       # POST /v1/invitations/team
    │                   │   │   └── verification/
    │                   │   │       └── page.tsx  # Business verification
    │                   │   │           # - Upload business documents
    │                   │   │           # - Company registration
    │                   │   │           # - Tax documents
    │                   │   │           # BE: admin-be/business_verification
    │                   │   │           # POST /v1/admin/business-verification
    │                   │   │           # BE: storage-be/uploads
    │                   │   ├── freelancer/  # Freelancer onboarding
    │                   │   │   ├── experience/
    │                   │   │   │   └── page.tsx  # Work experience
    │                   │   │   │       # - Add previous positions
    │                   │   │   │       # - Company, role, dates, description
    │                   │   │   │       # BE: users-be/experience
    │                   │   │   │       # POST /v1/users/{id}/experience
    │                   │   │   ├── portfolio/
    │                   │   │   │   └── page.tsx  # Portfolio items
    │                   │   │   │       # - Upload work samples
    │                   │   │   │       # - Project descriptions
    │                   │   │   │       # - Links to live work
    │                   │   │   │       # BE: users-be/portfolio
    │                   │   │   │       # POST /v1/users/{id}/portfolio
    │                   │   │   │       # BE: storage-be/uploads
    │                   │   │   │       # POST /v1/storage/upload (signed URL)
    │                   │   │   ├── preferences/
    │                   │   │   │   └── page.tsx  # Job preferences
    │                   │   │   │       # - Job categories
    │                   │   │   │       # - Work availability
    │                   │   │   │       # - Preferred job types
    │                   │   │   │       # - Notification settings
    │                   │   │   │       # BE: users-be/preferences
    │                   │   │   │       # POST /v1/users/{id}/preferences
    │                   │   │   │       # Publishes: FreelancerProfileCompleted event
    │                   │   │   ├── profile/
    │                   │   │   │   └── page.tsx  # Basic profile
    │                   │   │   │       # - Professional title
    │                   │   │   │       # - Bio
    │                   │   │   │       # - Location
    │                   │   │   │       # - Profile photo
    │                   │   │   │       # BE: users-be/profile
    │                   │   │   │       # PATCH /v1/users/{id}/profile
    │                   │   │   ├── rates/
    │                   │   │   │   └── page.tsx  # Rate setting
    │                   │   │   │       # - Hourly rate
    │                   │   │   │       # - Preferred project budget range
    │                   │   │   │       # - Currency
    │                   │   │   │       # BE: users-be/freelancer
    │                   │   │   │       # PATCH /v1/users/{id}/rates
    │                   │   │   └── skills/
    │                   │   │       └── page.tsx  # Skills selection
    │                   │   │           # - Skill search & autocomplete
    │                   │   │           # - Skill level (Beginner/Intermediate/Expert)
    │                   │   │           # - Primary skills (minimum 3)
    │                   │   │           # BE: users-be/capabilities
    │                   │   │           # POST /v1/users/{id}/skills
    │                   │   │           # Body: { skills: [{ skill_id, level }] }
    │                   │   ├── welcome/
    │                   │   │   └── page.tsx  # Welcome message
    │                   │   │       # - User type confirmation
    │                   │   │       # - Next steps preview
    │                   │   │       # BE: users-be/user
    │                   │   │       # GET /v1/users/me
    │                   │   └── layout.tsx  # Onboarding layout
    │                   │       # - Progress indicator
    │                   │       # - Skip options
    │                   │       # - Help sidebar
    │                   ├── (public)/  # Public pages (no auth required)
    │                   │   ├── about/
    │                   │   │   └── page.tsx  # About page
    │                   │   │       # - Company story
    │                   │   │       # - Team showcase
    │                   │   │       # - Mission & values
    │                   │   │       # BE: None (static)
    │                   │   ├── blog/
    │                   │   │   ├── [slug]/
    │                   │   │   │   └── page.tsx  # Blog post detail
    │                   │   │   ├── category/
    │                   │   │   │   └── [category]/
    │                   │   │   │       └── page.tsx  # Category listing
    │                   │   │   │           # BE: CMS or separate service
    │                   │   │   └── page.tsx  # Blog listing
    │                   │   ├── careers/
    │                   │   │   ├── [slug]/
    │                   │   │   │   └── page.tsx  # Individual job posting
    │                   │   │   │       # BE: None (static or CMS)
    │                   │   │   └── page.tsx  # Careers overview
    │                   │   ├── contact/
    │                   │   │   └── page.tsx  # Contact form
    │                   │   │       # - Support inquiries
    │                   │   │       # - Partnership requests
    │                   │   │       # - Media contacts
    │                   │   │       # BE: communications-be/messages
    │                   │   │       # POST /v1/contact
    │                   │   │       # Body: { name, email, subject, message, type }
    │                   │   ├── help/
    │                   │   │   ├── [category]/
    │                   │   │   │   └── page.tsx  # Help category
    │                   │   │   ├── article/
    │                   │   │   │   └── [slug]/
    │                   │   │   │       └── page.tsx  # Help article
    │                   │   │   │           # BE: communications-be/knowledge-base
    │                   │   │   │           # GET /v1/kb/articles
    │                   │   │   │           # GET /v1/kb/articles/{slug}
    │                   │   │   └── page.tsx  # Help center home
    │                   │   ├── how-it-works/
    │                   │   │   ├── clients/
    │                   │   │   │   └── page.tsx  # For clients
    │                   │   │   │       # - Posting jobs
    │                   │   │   │       # - Hiring process
    │                   │   │   │       # - Payment & escrow
    │                   │   │   │       # BE: None (static)
    │                   │   │   ├── freelancers/
    │                   │   │   │   └── page.tsx  # For freelancers
    │                   │   │   │       # - Getting started
    │                   │   │   │       # - Finding jobs
    │                   │   │   │       # - Earning money
    │                   │   │   │       # BE: None (static)
    │                   │   │   └── page.tsx  # How it works overview
    │                   │   ├── legal/
    │                   │   │   ├── compliance/
    │                   │   │   │   └── page.tsx  # Compliance information
    │                   │   │   │       # BE: None (static/versioned)
    │                   │   │   ├── cookies/
    │                   │   │   │   └── page.tsx  # Cookie Policy
    │                   │   │   ├── privacy/
    │                   │   │   │   └── page.tsx  # Privacy Policy
    │                   │   │   └── terms/
    │                   │   │       └── page.tsx  # Terms of Service
    │                   │   ├── pricing/
    │                   │   │   └── page.tsx  # Pricing page
    │                   │   │       # - Plan comparison
    │                   │   │       # - Feature matrix
    │                   │   │       # - FAQ
    │                   │   │       # BE: subscriptions-be/plans
    │                   │   │       # GET /v1/plans?public=true
    │                   │   │       # Returns: plans with public pricing
    │                   │   ├── trust-safety/
    │                   │   │   └── page.tsx  # Trust & Safety
    │                   │   │       # - Security measures
    │                   │   │       # - Payment protection
    │                   │   │       # - Dispute resolution
    │                   │   │       # BE: None (static)
    │                   │   ├── layout.tsx  # Public pages layout
    │                   │   │   # - Public header (no auth UI)
    │                   │   │   # - Footer
    │                   │   │   # - Language switcher
    │                   │   └── page.tsx  # Homepage / Landing
    │                   │       # - Hero section
    │                   │       # - Features overview
    │                   │       # - Stats showcase
    │                   │       # - Testimonials
    │                   │       # - CTA sections
    │                   │       # BE: None (static/cached content)
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
    │                   ├── transparency/
    │                   │   └── page.tsx  # Transparency report
    │                   │       # - User statistics
    │                   │       # - Moderation actions
    │                   │       # - Government requests
    │                   │       # BE: admin-be/reporting
    │                   │       # GET /v1/public/transparency-report
    │                   └── layout.tsx  # Root layout for locale
    │                       # - HTML lang attribute (i18n)
    │                       # - Head meta tags
    │                       # - Body layout
    │                       # - Font loading
    ├── apps/mobile/app/
    ├── apps/web/src/app/[locale]/
    ├── apps/web/src/app/[locale]/(dashboard)/
    ├── auction/
    │   ├── AuctionTimer.native.tsx
    │   ├── AuctionTimer.tsx  # Countdown timer
    │   ├── AuctionTimer.web.tsx
    │   ├── BidHistoryChart.native.tsx
    │   ├── BidHistoryChart.tsx  # Bid history visualization
    │   ├── BidHistoryChart.web.tsx
    │   ├── LiveBidFeed.native.tsx
    │   ├── LiveBidFeed.tsx  # Real-time bid feed
    │   └── LiveBidFeed.web.tsx
    ├── auctions/
    │   ├── api/
    │   │   └── auctions-api.ts  # Auctions API client
    │   │       # BE: proposals-be/auction
    │   ├── hooks/
    │   │   ├── useActiveAuctions.ts  # Active auctions list
    │   │   ├── useAuction.ts  # Single auction
    │   │   ├── useAuctionBid.ts  # Place bid
    │   │   └── useAuctionHistory.ts  # Bid history
    │   ├── queries/
    │   │   ├── auctions-mutations.ts  # Auction mutations
    │   │   └── auctions-queries.ts  # Auction queries
    │   ├── store/
    │   │   └── auction-store.ts  # Real-time auction state (Zustand)
    │   └── types.ts  # Auction types
    │       # 2. Negotiations Module
    ├── bidding/
    │   ├── analytics/
    │   │   └── page.tsx  # Bidding analytics
    │   │       # - Win rate
    │   │       # - Average bid amount
    │   │       # - Competition analysis
    │   │       # BE: proposals-be/bid-strategy
    │   │       # GET /v1/bid-strategies/analytics
    │   │       # 4. Invitations Management
    │   ├── api/
    │   │   ├── bid-api.ts  # Bid placement API
    │   │   │   # BE: proposals-be/bid
    │   │   └── bid-strategy-api.ts  # Bid strategy API
    │   │       # BE: proposals-be/bid-strategy
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
    │   │   ├── page.tsx  # Active auctions list
    │   │   │   # BE: proposals-be/auction
    │   │   │   # GET /v1/auctions/active
    │   │   └── │
    │   ├── hooks/
    │   │   ├── useBidAnalytics.ts  # Bid analytics
    │   │   ├── useBidHistory.ts  # Bid history
    │   │   ├── useBidStrategies.ts  # List strategies
    │   │   ├── useBidStrategy.ts  # Bid strategy management
    │   │   └── usePlaceBid.ts  # Place bid
    │   ├── queries/
    │   │   ├── bidding-mutations.ts  # Bidding mutations
    │   │   └── bidding-queries.ts  # Bidding queries
    │   ├── strategies/
    │   │   ├── [strategyId]/
    │   │   │   ├── edit/
    │   │   │   │   └── page.tsx  # Edit bid strategy
    │   │   │   │       # BE: proposals-be/bid-strategy
    │   │   │   │       # PUT /v1/bid-strategies/{strategy_id}
    │   │   │   └── page.tsx  # View bid strategy details
    │   │   │       # BE: proposals-be/bid-strategy
    │   │   │       # GET /v1/bid-strategies/{strategy_id}
    │   │   ├── new/
    │   │   │   └── page.tsx  # Create new bid strategy
    │   │   │       # BE: proposals-be/bid-strategy
    │   │   │       # POST /v1/bid-strategies
    │   │   ├── page.tsx  # Bid strategies list
    │   │   │   # - Auto-bid rules
    │   │   │   # - Price ranges
    │   │   │   # - Category targeting
    │   │   │   # BE: proposals-be/bid-strategy
    │   │   │   # GET /v1/bid-strategies
    │   │   └── │
    │   ├── types.ts  # Bidding types
    │   │   # 12. Invitations Module
    │   └── │
    ├── charts/
    │   ├── EarningsChart.native.tsx
    │   ├── EarningsChart.tsx  # Earnings visualization
    │   ├── EarningsChart.web.tsx
    │   ├── PerformanceChart.native.tsx
    │   ├── PerformanceChart.tsx  # Performance metrics
    │   ├── PerformanceChart.web.tsx
    │   ├── TrendChart.native.tsx
    │   ├── TrendChart.tsx  # Trend visualization
    │   └── TrendChart.web.tsx
    ├── collaboration/
    │   ├── CollaborationPanel.native.tsx
    │   ├── CollaborationPanel.tsx  # Team collaboration
    │   ├── CollaborationPanel.web.tsx
    │   ├── GroupCard.native.tsx
    │   ├── GroupCard.tsx  # User group card
    │   ├── GroupCard.web.tsx
    │   ├── MentorCard.native.tsx
    │   ├── MentorCard.tsx  # Mentor profile card
    │   └── MentorCard.web.tsx
    ├── compliance/
    │   ├── api/
    │   │   ├── compliance-api.ts  # Compliance API
    │   │   │   # BE: users-be/compliance
    │   │   └── tax-profile-api.ts  # Tax profile API
    │   │       # BE: users-be/compliance, financial-be/tax
    │   ├── documents/
    │   │   ├── [documentId]/
    │   │   │   └── page.tsx  # Compliance document details
    │   │   │       # BE: storage-be/asset, admin-be/business_verification
    │   │   │       # GET /v1/compliance/documents/{document_id}
    │   │   ├── upload/
    │   │   │   └── page.tsx  # Upload compliance documents
    │   │   │       # BE: storage-be/asset, admin-be/business_verification
    │   │   │       # POST /v1/compliance/documents/upload
    │   │   ├── page.tsx  # Compliance documents list
    │   │   │   # BE: admin-be/business_verification
    │   │   │   # GET /v1/compliance/documents
    │   │   └── │
    │   ├── hooks/
    │   │   ├── useComplianceDocuments.ts  # Document management
    │   │   ├── useComplianceProfile.ts  # Compliance profile
    │   │   ├── useTaxProfile.ts  # Tax profile management
    │   │   └── useTaxReports.ts  # Tax reports
    │   ├── queries/
    │   │   ├── compliance-mutations.ts  # Compliance mutations
    │   │   └── compliance-queries.ts  # Compliance queries
    │   ├── reports/
    │   │   ├── tax-summary/
    │   │   │   └── page.tsx  # Annual tax summary
    │   │   │       # BE: financial-be/tax
    │   │   │       # GET /v1/tax/reports/annual-summary
    │   │   ├── page.tsx  # Compliance reports
    │   │   │   # - Income reports
    │   │   │   # - Tax withholding
    │   │   │   # - Payment history
    │   │   │   # BE: financial-be/reports
    │   │   │   # GET /v1/reports/compliance
    │   │   │   # 12. Analytics & Insights (Enhanced)
    │   │   └── │
    │   ├── tax-profile/
    │   │   ├── edit/
    │   │   │   └── page.tsx  # Edit tax profile
    │   │   │       # BE: users-be/compliance, financial-be/tax
    │   │   │       # PUT /v1/users/me/compliance/tax-profile
    │   │   ├── page.tsx  # Tax profile overview
    │   │   │   # - Tax ID
    │   │   │   # - Tax forms
    │   │   │   # - Withholding settings
    │   │   │   # BE: users-be/compliance
    │   │   │   # GET /v1/users/me/compliance/tax-profile
    │   │   └── │
    │   ├── DocumentUploader.native.tsx
    │   ├── DocumentUploader.tsx  # Compliance doc uploader
    │   ├── DocumentUploader.web.tsx
    │   ├── types.ts  # Compliance types
    │   │   # 10. Search & Discovery Module (Enhanced)
    │   ├── VerificationStatus.native.tsx
    │   ├── VerificationStatus.tsx  # Verification status badge
    │   ├── VerificationStatus.web.tsx
    │   └── │
    ├── config/
    │   ├── constants/
    │   │   ├── api-endpoints.ts  # API endpoint constants
    │   │   ├── app-config.ts  # App configuration
    │   │   ├── performance-budgets.ts  # Performance budgets
    │   │   └── third-party-keys.ts  # Third-party service keys
    │   ├── environments/
    │   │   ├── development.ts  # Dev config
    │   │   ├── production.ts  # Production config
    │   │   └── staging.ts  # Staging config
    │   └── feature-flags.ts  # Feature flags config
    ├── connects/
    │   ├── api/
    │   │   └── connects-api.ts  # Connects API client
    │   │       # BE: proposals-be/connect
    │   ├── hooks/
    │   │   ├── useConnectPackages.ts  # Available packages
    │   │   ├── useConnectRefund.ts  # Request refund
    │   │   ├── useConnects.ts  # Connects balance and history
    │   │   └── usePurchaseConnects.ts  # Purchase connects
    │   ├── purchase/
    │   │   └── page.tsx  # Purchase connects
    │   │       # - Select package
    │   │       # - Payment processing
    │   │       # BE: proposals-be/connect, financial-be/payment
    │   │       # GET /v1/connects/packages
    │   │       # POST /v1/connects/purchase
    │   ├── queries/
    │   │   ├── connects-mutations.ts  # Connect mutations
    │   │   └── connects-queries.ts  # Connect queries
    │   ├── usage/
    │   │   └── page.tsx  # Connects usage analytics
    │   │       # - Spending patterns
    │   │       # - Refund history
    │   │       # - ROI tracking
    │   │       # BE: proposals-be/connect
    │   │       # GET /v1/connects/usage-analytics
    │   ├── page.tsx  # Connects dashboard
    │   │   # - Current balance
    │   │   # - Transaction history
    │   │   # - Refund requests
    │   │   # BE: proposals-be/connect
    │   │   # GET /v1/connects
    │   │   # GET /v1/connects/balance
    │   │   # 3. Bidding & Strategy Management (Freelancer)
    │   ├── types.ts  # Connect types
    │   │   # 5. Work Tracking Module
    │   └── │
    ├── deliverables/
    │   ├── [contractId]/
    │   │   ├── [deliverableId]/
    │   │   │   ├── review/
    │   │   │   │   └── page.tsx  # Review deliverable (client)
    │   │   │   │       # - Approve/reject
    │   │   │   │       # - Request changes
    │   │   │   │       # - Add comments
    │   │   │   │       # BE: contracts-be/deliverable
    │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/review
    │   │   │   ├── revisions/
    │   │   │   │   ├── [revisionId]/
    │   │   │   │   │   └── page.tsx  # Revision detail
    │   │   │   │   │       # BE: contracts-be/deliverable
    │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/revisions/{revision_id}
    │   │   │   │   ├── page.tsx  # Revision history
    │   │   │   │   │   # BE: contracts-be/deliverable
    │   │   │   │   │   # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/revisions
    │   │   │   │   └── │
    │   │   │   ├── upload/
    │   │   │   │   └── page.tsx  # Upload new version
    │   │   │   │       # BE: contracts-be/deliverable, storage-be/asset
    │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/upload
    │   │   │   ├── page.tsx  # Deliverable details
    │   │   │   │   # - File viewer
    │   │   │   │   # - Download
    │   │   │   │   # - Metadata
    │   │   │   │   # - Comments thread
    │   │   │   │   # BE: contracts-be/deliverable, storage-be/asset
    │   │   │   │   # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}
    │   │   │   └── │
    │   │   ├── new/
    │   │   │   └── page.tsx  # Submit new deliverable
    │   │   │       # BE: contracts-be/deliverable, storage-be/asset
    │   │   │       # POST /v1/contracts/{contract_id}/deliverables
    │   │   ├── page.tsx  # Contract deliverables list
    │   │   │   # BE: contracts-be/deliverable
    │   │   │   # GET /v1/contracts/{contract_id}/deliverables
    │   │   └── │
    │   ├── api/
    │   │   └── deliverables-api.ts  # Deliverables API client
    │   │       # BE: contracts-be/deliverable, storage-be/asset
    │   ├── hooks/
    │   │   ├── useDeliverable.ts  # Single deliverable
    │   │   ├── useDeliverableRevisions.ts  # Revision management
    │   │   ├── useDeliverables.ts  # Deliverables list
    │   │   ├── useReviewDeliverable.ts  # Review deliverable (client)
    │   │   └── useUploadDeliverable.ts  # Upload deliverable
    │   ├── pending-review/
    │   │   └── page.tsx  # Deliverables pending client review
    │   │       # BE: contracts-be/deliverable
    │   │       # GET /v1/deliverables/pending-review
    │   ├── queries/
    │   │   ├── deliverables-mutations.ts  # Deliverable mutations
    │   │   └── deliverables-queries.ts  # Deliverable queries
    │   ├── page.tsx  # All deliverables overview
    │   │   # BE: contracts-be/deliverable
    │   │   # GET /v1/deliverables
    │   │   # 8. Reviews & Reputation Management
    │   ├── types.ts  # Deliverable types
    │   │   # 7. Learning Module
    │   └── │
    ├── deploy/  # Deployment configurations
    │   ├── docker/
    │   │   ├── mobile-build.Dockerfile  # EAS build container
    │   │   ├── nginx.conf  # NGINX config for static serving
    │   │   ├── web.Dockerfile  # Multi-stage Next.js build
    │   │   └── web.dockerignore
    │   └── k8s/  # Kubernetes manifests
    │       ├── mobile-api/  # Mobile API gateway (if separate)
    │       ├── monitoring/
    │       │   ├── grafana-dashboard.yaml  # Dashboards
    │       │   └── prometheus.yaml  # Metrics scraping
    │       └── web/
    │           ├── configmap.yaml  # Environment config
    │           ├── deployment.yaml  # Web app deployment
    │           ├── hpa.yaml  # Horizontal Pod Autoscaler
    │           ├── ingress.yaml  # NGINX ingress
    │           └── service.yaml  # ClusterIP service
    ├── developers/
    │   ├── api-reference/
    │   │   ├── [endpoint]/
    │   │   │   └── page.tsx  # API endpoint reference
    │   │   │       # BE: static (OpenAPI spec)
    │   │   ├── page.tsx  # API reference home
    │   │   │   # BE: static (OpenAPI spec)
    │   │   └── │
    │   ├── docs/
    │   │   ├── [section]/
    │   │   │   └── page.tsx  # Documentation section
    │   │   │       # BE: static or CMS
    │   │   │       # GET /v1/content/docs/{section}
    │   │   ├── page.tsx  # API documentation home
    │   │   │   # BE: static
    │   │   └── │
    │   ├── sandbox/
    │   │   └── page.tsx  # API sandbox/playground
    │   │       # BE: developer API
    │   │       # POST /v1/developer/sandbox/execute
    │   │       # 5. Platform Status & Trust
    │   ├── sdks/
    │   │   └── page.tsx  # SDK downloads and docs
    │   │       # BE: static
    │   ├── webhooks/
    │   │   └── page.tsx  # Webhooks documentation
    │   │       # BE: static
    │   └── │
    ├── docs/  # Documentation
    │   ├── adr/  # Architecture Decision Records
    │   │   ├── ...
    │   │   ├── 001-monorepo-structure.md
    │   │   ├── 002-state-management.md
    │   │   ├── 003-authentication-approach.md
    │   │   └── 004-component-library.md
    │   ├── api-integration/
    │   │   ├── admin-be.md
    │   │   ├── communications-be.md
    │   │   ├── contracts-be.md
    │   │   ├── financial-be.md
    │   │   ├── jobs-be.md
    │   │   ├── proposals-be.md
    │   │   ├── reviews-be.md
    │   │   ├── search-be.md
    │   │   ├── storage-be.md
    │   │   ├── subscriptions-be.md
    │   │   ├── users-be.md
    │   │   └── utility-services.md
    │   ├── architecture/
    │   │   ├── authentication-flow.md
    │   │   ├── data-fetching-patterns.md
    │   │   ├── frontend-architecture.md
    │   │   ├── microservices-integration.md
    │   │   ├── routing-strategy.md
    │   │   └── state-management.md
    │   ├── components/
    │   │   ├── accessibility.md
    │   │   ├── component-library.md
    │   │   ├── design-system.md
    │   │   └── theming.md
    │   ├── features/
    │   │   ├── admin-tools.md
    │   │   ├── authentication.md
    │   │   ├── contract-management.md
    │   │   ├── job-management.md
    │   │   ├── messaging.md
    │   │   ├── notifications.md
    │   │   ├── payment-system.md
    │   │   ├── proposal-system.md
    │   │   └── search-discovery.md
    │   ├── guides/
    │   │   ├── contributing.md
    │   │   ├── deployment.md
    │   │   ├── development-workflow.md
    │   │   ├── getting-started.md
    │   │   ├── testing-guide.md
    │   │   └── troubleshooting.md
    │   ├── ARCHITECTURE.md  # System architecture
    │   ├── CONTRIBUTING.md  # Contribution guidelines
    │   ├── DEPLOYMENT.md  # Deployment procedures
    │   ├── MICROSERVICES_MAPPING.md  # BE microservices integration
    │   ├── PERFORMANCE.md  # Performance optimization guide
    │   ├── README.md  # Project overview
    │   ├── SETUP.md  # Development setup guide
    │   ├── STATE_MANAGEMENT.md  # TanStack Query + Zustand patterns
    │   └── TESTING.md  # Testing strategy
    ├── enterprise/
    │   ├── case-studies/
    │   │   └── page.tsx  # Enterprise case studies
    │   │       # BE: CMS
    │   │       # GET /v1/content/case-studies?type=enterprise
    │   ├── contact/
    │   │   └── page.tsx  # Enterprise contact/demo request
    │   │       # BE: communications-be
    │   │       # POST /v1/contact/enterprise
    │   │       # 4. Developers Portal
    │   ├── pricing/
    │   │   └── page.tsx  # Enterprise pricing
    │   │       # BE: financial-be/subscription
    │   │       # GET /v1/subscriptions/plans?type=enterprise
    │   ├── solutions/
    │   │   ├── managed-services/
    │   │   │   └── page.tsx  # Managed services offering
    │   │   │       # BE: none (marketing content)
    │   │   ├── staffing/
    │   │   │   └── page.tsx  # Enterprise staffing solutions
    │   │   │       # BE: none (marketing content)
    │   │   ├── page.tsx  # Enterprise solutions overview
    │   │   │   # BE: none (marketing content)
    │   │   └── │
    │   └── │
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
    │       # 15. Real-time Updates Module
    ├── interviews/
    │   ├── api/
    │   │   └── interviews-api.ts  # Interviews API client
    │   │       # BE: proposals-be/interview
    │   ├── hooks/
    │   │   ├── useInterview.ts  # Single interview
    │   │   ├── useInterviewFeedback.ts  # Interview feedback
    │   │   ├── useInterviews.ts  # Interviews list
    │   │   └── useScheduleInterview.ts  # Schedule interview
    │   ├── queries/
    │   │   ├── interviews-mutations.ts  # Interview mutations
    │   │   └── interviews-queries.ts  # Interview queries
    │   └── types.ts  # Interview types
    │       # 4. Connects Module
    ├── invitations/
    │   ├── api/
    │   │   ├── job-invitations-api.ts  # Job invitations API (client)
    │   │   │   # BE: jobs-be/invitation
    │   │   └── proposal-invites-api.ts  # Proposal invites API (freelancer)
    │   │       # BE: proposals-be/invite
    │   ├── hooks/
    │   │   ├── useAcceptInvite.ts  # Accept invite (freelancer)
    │   │   ├── useDeclineInvite.ts  # Decline invite (freelancer)
    │   │   ├── useInvitationAnalytics.ts  # Invitation metrics
    │   │   ├── useInvitations.ts  # Invitations management
    │   │   └── useSendInvitation.ts  # Send invitation (client)
    │   ├── queries/
    │   │   ├── invitations-mutations.ts  # Invitation mutations
    │   │   └── invitations-queries.ts  # Invitation queries
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
    │   │   ├── page.tsx  # Received invitations list
    │   │   │   # BE: proposals-be/invite
    │   │   │   # GET /v1/invites/received
    │   │   └── │
    │   ├── sent/
    │   │   ├── [inviteId]/
    │   │   │   └── page.tsx  # Sent invitation tracking
    │   │   │       # - Delivery status
    │   │   │       # - Response tracking
    │   │   │       # BE: jobs-be/invitation
    │   │   │       # GET /v1/jobs/{job_id}/invitations/{invite_id}
    │   │   ├── page.tsx  # Sent invitations list (client)
    │   │   │   # BE: jobs-be/invitation
    │   │   │   # GET /v1/jobs/{job_id}/invitations
    │   │   └── │
    │   ├── page.tsx  # Invitations overview
    │   │   # - Pending actions
    │   │   # - Response rate (client)
    │   │   # - Conversion metrics
    │   │   # BE: proposals-be/invite OR jobs-be/invitation (based on role)
    │   │   # 5. Talent Sourcing & Shortlists (Client)
    │   ├── types.ts  # Invitation types
    │   │   # 13. Shortlists Module
    │   └── │
    ├── learning/
    │   ├── achievements/
    │   │   └── page.tsx  # Achievements & badges
    │   │       # - Earned badges
    │   │       # - Progress to next level
    │   │       # - Leaderboard
    │   │       # BE: users-be/achievement
    │   │       # GET /v1/users/me/achievements
    │   ├── api/
    │   │   ├── learning-paths-api.ts  # Learning paths API
    │   │   │   # BE: users-be/learning_path
    │   │   └── mentorship-api.ts  # Mentorship API
    │   │       # BE: users-be/mentorship
    │   ├── certifications/
    │   │   └── page.tsx  # Manage certifications
    │   │       # - Upload certificates
    │   │       # - Verification status
    │   │       # - Expiry tracking
    │   │       # BE: users-be/credential
    │   │       # GET /v1/users/me/credentials?type=certification
    │   │       # 11. Compliance & Tax Management
    │   ├── hooks/
    │   │   ├── useAchievements.ts  # Achievements/badges
    │   │   ├── useLearningPath.ts  # Single learning path
    │   │   ├── useLearningPaths.ts  # Learning paths list
    │   │   ├── useLearningProgress.ts  # Track progress
    │   │   └── useMentorship.ts  # Mentorship management
    │   ├── mentorship/
    │   │   ├── [sessionId]/
    │   │   │   └── page.tsx  # Mentorship session details
    │   │   │       # BE: users-be/mentorship
    │   │   │       # GET /v1/users/me/mentorship/{session_id}
    │   │   ├── find-mentor/
    │   │   │   └── page.tsx  # Find a mentor
    │   │   │       # BE: users-be/mentorship, search-be/query
    │   │   │       # POST /v1/search/mentors
    │   │   ├── my-mentees/
    │   │   │   └── page.tsx  # Manage mentees
    │   │   │       # BE: users-be/mentorship
    │   │   │       # GET /v1/users/me/mentorship/mentees
    │   │   ├── page.tsx  # Mentorship dashboard
    │   │   │   # BE: users-be/mentorship
    │   │   │   # GET /v1/users/me/mentorship
    │   │   └── │
    │   ├── paths/
    │   │   ├── [pathId]/
    │   │   │   ├── progress/
    │   │   │   │   └── page.tsx  # Learning path progress
    │   │   │   │       # BE: users-be/learning_path
    │   │   │   │       # GET /v1/users/me/learning-path/{path_id}/progress
    │   │   │   ├── page.tsx  # Learning path details
    │   │   │   │   # - Courses
    │   │   │   │   # - Milestones
    │   │   │   │   # - Resources
    │   │   │   │   # BE: users-be/learning_path
    │   │   │   │   # GET /v1/users/me/learning-path/{path_id}
    │   │   │   └── │
    │   │   ├── discover/
    │   │   │   └── page.tsx  # Discover learning paths
    │   │   │       # BE: users-be/learning_path
    │   │   │       # GET /v1/learning-paths/discover
    │   │   ├── page.tsx  # My learning paths
    │   │   │   # BE: users-be/learning_path
    │   │   │   # GET /v1/users/me/learning-path
    │   │   └── │
    │   ├── queries/
    │   │   ├── learning-mutations.ts  # Learning mutations
    │   │   └── learning-queries.ts  # Learning queries
    │   ├── AchievementBadge.native.tsx
    │   ├── AchievementBadge.tsx  # Achievement badge
    │   ├── AchievementBadge.web.tsx
    │   ├── LearningPathCard.native.tsx
    │   ├── LearningPathCard.tsx  # Learning path card
    │   ├── LearningPathCard.web.tsx
    │   ├── ProgressTracker.native.tsx
    │   ├── ProgressTracker.tsx  # Progress visualization
    │   ├── ProgressTracker.web.tsx
    │   ├── types.ts  # Learning types
    │   │   # 8. Networking Module
    │   └── │
    ├── legal/
    │   ├── accessibility/
    │   │   └── page.tsx  # Accessibility statement
    │   │       # BE: none (static content)
    │   │       # 2. Resources & Help Center
    │   ├── compliance/
    │   │   ├── ccpa/
    │   │   │   └── page.tsx  # CCPA compliance
    │   │   │       # BE: none (static content)
    │   │   ├── gdpr/
    │   │   │   └── page.tsx  # GDPR compliance
    │   │   │       # BE: none (static content)
    │   │   ├── page.tsx  # Compliance overview
    │   │   │   # BE: none (static content)
    │   │   └── │
    │   ├── dmca/
    │   │   └── page.tsx  # DMCA policy
    │   │       # BE: none (static content)
    │   ├── ip-policy/
    │   │   └── page.tsx  # Intellectual property policy
    │   │       # BE: none (static content)
    │   ├── privacy/
    │   │   ├── cookie-policy/
    │   │   │   └── page.tsx  # Cookie policy
    │   │   │       # BE: none (static content)
    │   │   ├── data-processing/
    │   │   │   └── page.tsx  # Data processing agreement
    │   │   │       # BE: none (static content)
    │   │   ├── page.tsx  # Privacy policy
    │   │   │   # BE: none (static content)
    │   │   └── │
    │   ├── terms/
    │   │   ├── client/
    │   │   │   └── page.tsx  # Client terms of service
    │   │   │       # BE: none (static content)
    │   │   ├── freelancer/
    │   │   │   └── page.tsx  # Freelancer terms of service
    │   │   │       # BE: none (static content with version from CMS)
    │   │   ├── page.tsx  # General terms
    │   │   │   # BE: none (static content)
    │   │   └── │
    │   └── │
    ├── negotiations/
    │   ├── api/
    │   │   └── negotiations-api.ts  # Negotiations API client
    │   │       # BE: proposals-be/negotiation
    │   ├── hooks/
    │   │   ├── useNegotiation.ts  # Single negotiation
    │   │   ├── useNegotiationHistory.ts  # Negotiation history
    │   │   └── useNegotiationOffer.ts  # Make/accept/reject offer
    │   ├── queries/
    │   │   ├── negotiations-mutations.ts  # Negotiation mutations
    │   │   └── negotiations-queries.ts  # Negotiation queries
    │   └── types.ts  # Negotiation types
    │       # 3. Interviews Module
    ├── network/
    │   ├── connections/
    │   │   ├── [userId]/
    │   │   │   └── page.tsx  # Connection profile view
    │   │   │       # BE: users-be/profile, users-be/connection
    │   │   │       # GET /v1/users/{user_id}
    │   │   │       # GET /v1/users/me/connections/{user_id}
    │   │   ├── pending/
    │   │   │   └── page.tsx  # Pending connection requests
    │   │   │       # BE: users-be/connection
    │   │   │       # GET /v1/users/me/connections/pending
    │   │   ├── page.tsx  # Connections list
    │   │   │   # BE: users-be/connection
    │   │   │   # GET /v1/users/me/connections
    │   │   └── │
    │   ├── groups/
    │   │   ├── [groupId]/
    │   │   │   ├── members/
    │   │   │   │   └── page.tsx  # Group members
    │   │   │   │       # BE: users-be/user_group
    │   │   │   │       # GET /v1/groups/{group_id}/members
    │   │   │   ├── page.tsx  # Group details
    │   │   │   │   # - Posts
    │   │   │   │   # - Events
    │   │   │   │   # - Resources
    │   │   │   │   # BE: users-be/user_group
    │   │   │   │   # GET /v1/groups/{group_id}
    │   │   │   └── │
    │   │   ├── discover/
    │   │   │   └── page.tsx  # Discover groups
    │   │   │       # BE: users-be/user_group
    │   │   │       # GET /v1/groups/discover
    │   │   ├── page.tsx  # My groups
    │   │   │   # BE: users-be/user_group
    │   │   │   # GET /v1/users/me/groups
    │   │   └── │
    │   ├── recommendations/
    │   │   └── page.tsx  # Connection recommendations
    │   │       # - People you may know
    │   │       # - Similar professionals
    │   │       # BE: search-be/recommendation
    │   │       # GET /v1/recommendations/connections
    │   │       # 10. Learning & Professional Development
    │   ├── referrals/
    │   │   ├── dashboard/
    │   │   │   └── page.tsx  # Referral dashboard
    │   │   │       # - Total referrals
    │   │   │       # - Earnings
    │   │   │       # - Conversion rate
    │   │   │       # BE: users-be/referral
    │   │   │       # GET /v1/users/me/referral-code
    │   │   │       # GET /v1/referrals/analytics
    │   │   ├── page.tsx  # Referrals overview
    │   │   │   # - Share referral code
    │   │   │   # - Track referrals
    │   │   │   # BE: users-be/referral
    │   │   │   # GET /v1/referrals
    │   │   └── │
    │   └── │
    ├── networking/
    │   ├── api/
    │   │   ├── connections-api.ts  # Connections API
    │   │   │   # BE: users-be/connection
    │   │   ├── groups-api.ts  # Groups API
    │   │   │   # BE: users-be/user_group
    │   │   └── referrals-api.ts  # Referrals API
    │   │       # BE: users-be/referral
    │   ├── hooks/
    │   │   ├── useConnectionRequest.ts  # Send/accept/reject
    │   │   ├── useConnections.ts  # Connections management
    │   │   ├── useGroups.ts  # Groups management
    │   │   ├── useNetworkRecommendations.ts  # Connection recommendations
    │   │   └── useReferrals.ts  # Referral management
    │   ├── queries/
    │   │   ├── networking-mutations.ts  # Networking mutations
    │   │   └── networking-queries.ts  # Networking queries
    │   └── types.ts  # Networking types
    │       # 9. Compliance Module
    ├── offline/
    │   ├── queue.tsx  # Offline actions queue
    │   │   # - Pending uploads
    │   │   # - Queued messages
    │   │   # - Draft proposals
    │   └── sync.tsx  # Sync status
    │       # - Sync progress
    │       # - Conflict resolution
    ├── packages/  # Shared libraries
    │   ├── config/  # Shared configurations
    │   │   ├── eslint-config/
    │   │   │   ├── index.js  # Base ESLint config
    │   │   │   ├── next.js  # Next.js-specific rules
    │   │   │   ├── package.json
    │   │   │   └── react.js  # React-specific rules
    │   │   ├── tailwind-config/
    │   │   │   ├── themes/
    │   │   │   │   ├── dark.js
    │   │   │   │   └── light.js
    │   │   │   ├── index.js  # Base Tailwind config
    │   │   │   │   # - Design tokens
    │   │   │   │   # - Color palette
    │   │   │   │   # - Spacing scale
    │   │   │   │   # - Typography
    │   │   │   │   # - Breakpoints
    │   │   │   │   # - Animations
    │   │   │   └── package.json
    │   │   └── typescript-config/
    │   │       ├── base.json  # Base TS config
    │   │       ├── nextjs.json  # Next.js TS config
    │   │       ├── package.json
    │   │       └── react-native.json  # React Native TS config
    │   ├── shared/  # Business logic, hooks, utilities
    │   │   # Shared business logic, hooks, utilities
    │   │   ├── src/
    │   │   │   ├── accessibility/
    │   │   │   │   ├── testing/
    │   │   │   │   │   ├── a11y-test-utils.ts  # Testing utilities
    │   │   │   │   │   └── axe-config.ts  # axe-core configuration
    │   │   │   │   ├── a11y-utils.ts  # Accessibility utilities
    │   │   │   │   ├── aria-utils.ts  # ARIA utilities
    │   │   │   │   ├── focus-management.ts  # Focus management
    │   │   │   │   ├── keyboard-navigation.ts  # Keyboard navigation
    │   │   │   │   └── screen-reader.ts  # Screen reader utilities
    │   │   │   ├── activity/
    │   │   │   │   ├── activity-client.ts  # Activity API client
    │   │   │   │   │   # BE: utility/activity (if exists) or communications-be/notification
    │   │   │   │   │   # GET /v1/activity/feed
    │   │   │   │   │   # GET /v1/activity/user/{user_id}
    │   │   │   │   └── types.ts
    │   │   │   ├── compliance/
    │   │   │   │   ├── compliance-client.ts  # Compliance API client
    │   │   │   │   │   # BE: admin-be/privacy, admin-be/kyc_case
    │   │   │   │   │   # GET /v1/privacy/export-requests
    │   │   │   │   │   # POST /v1/privacy/export-requests/{id}/process
    │   │   │   │   │   # GET /v1/privacy/deletion-requests
    │   │   │   │   │   # POST /v1/privacy/deletion-requests/{id}/process
    │   │   │   │   └── types.ts
    │   │   │   ├── experiments/
    │   │   │   │   ├── experiments-client.ts  # Experiments API client
    │   │   │   │   │   # BE: utility/experiments (if exists) or utility/flags
    │   │   │   │   │   # GET /v1/experiments/active
    │   │   │   │   │   # POST /v1/experiments/{id}/track
    │   │   │   │   └── types.ts
    │   │   │   ├── features/  # Feature-specific shared code
    │   │   │   │   ├── admin/  # Admin feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useAdminUsers.ts
    │   │   │   │   │   │   ├── useDisputes.ts
    │   │   │   │   │   │   ├── useKYCCases.ts
    │   │   │   │   │   │   └── useModeration.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── admin-mutations.ts
    │   │   │   │   │   │   └── admin-queries.ts  # BE: admin-be
    │   │   │   │   │   ├── admin-api.ts
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── auctions/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── auctions-api.ts  # Auctions API client
    │   │   │   │   │   │       # BE: proposals-be/auction
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useActiveAuctions.ts  # Active auctions list
    │   │   │   │   │   │   ├── useAuction.ts  # Single auction
    │   │   │   │   │   │   ├── useAuctionBid.ts  # Place bid
    │   │   │   │   │   │   └── useAuctionHistory.ts  # Bid history
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── auctions-mutations.ts  # Auction mutations
    │   │   │   │   │   │   └── auctions-queries.ts  # Auction queries
    │   │   │   │   │   ├── store/
    │   │   │   │   │   │   └── auction-store.ts  # Real-time auction state (Zustand)
    │   │   │   │   │   └── types.ts  # Auction types
    │   │   │   │   ├── auth/  # Authentication
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useAuth.ts  # Main auth hook
    │   │   │   │   │   │   │   # - isAuthenticated
    │   │   │   │   │   │   │   # - user
    │   │   │   │   │   │   │   # - login, logout
    │   │   │   │   │   │   │   # - refresh token
    │   │   │   │   │   │   ├── useKeycloak.ts  # Keycloak integration
    │   │   │   │   │   │   ├── usePermissions.ts  # RBAC permissions
    │   │   │   │   │   │   └── useSession.ts  # Session management
    │   │   │   │   │   ├── stores/
    │   │   │   │   │   │   └── auth-store.ts  # Zustand auth store
    │   │   │   │   │   │       # - user
    │   │   │   │   │   │       # - tokens
    │   │   │   │   │   │       # - login/logout actions
    │   │   │   │   │   │       # - Persisted to storage
    │   │   │   │   │   ├── utils/
    │   │   │   │   │   │   ├── rbac.ts  # RBAC utilities
    │   │   │   │   │   │   ├── session.ts  # Session utilities
    │   │   │   │   │   │   └── token.ts  # Token utilities
    │   │   │   │   │   └── types.ts  # Auth types
    │   │   │   │   ├── bidding/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── bid-api.ts  # Bid placement API
    │   │   │   │   │   │   │   # BE: proposals-be/bid
    │   │   │   │   │   │   └── bid-strategy-api.ts  # Bid strategy API
    │   │   │   │   │   │       # BE: proposals-be/bid-strategy
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useBidAnalytics.ts  # Bid analytics
    │   │   │   │   │   │   ├── useBidHistory.ts  # Bid history
    │   │   │   │   │   │   ├── useBidStrategies.ts  # List strategies
    │   │   │   │   │   │   ├── useBidStrategy.ts  # Bid strategy management
    │   │   │   │   │   │   └── usePlaceBid.ts  # Place bid
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── bidding-mutations.ts  # Bidding mutations
    │   │   │   │   │   │   └── bidding-queries.ts  # Bidding queries
    │   │   │   │   │   └── types.ts  # Bidding types
    │   │   │   │   ├── compliance/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── compliance-api.ts  # Compliance API
    │   │   │   │   │   │   │   # BE: users-be/compliance
    │   │   │   │   │   │   └── tax-profile-api.ts  # Tax profile API
    │   │   │   │   │   │       # BE: users-be/compliance, financial-be/tax
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useComplianceDocuments.ts  # Document management
    │   │   │   │   │   │   ├── useComplianceProfile.ts  # Compliance profile
    │   │   │   │   │   │   ├── useTaxProfile.ts  # Tax profile management
    │   │   │   │   │   │   └── useTaxReports.ts  # Tax reports
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── compliance-mutations.ts  # Compliance mutations
    │   │   │   │   │   │   └── compliance-queries.ts  # Compliance queries
    │   │   │   │   │   └── types.ts  # Compliance types
    │   │   │   │   ├── connects/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── connects-api.ts  # Connects API client
    │   │   │   │   │   │       # BE: proposals-be/connect
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useConnectPackages.ts  # Available packages
    │   │   │   │   │   │   ├── useConnectRefund.ts  # Request refund
    │   │   │   │   │   │   ├── useConnects.ts  # Connects balance and history
    │   │   │   │   │   │   └── usePurchaseConnects.ts  # Purchase connects
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── connects-mutations.ts  # Connect mutations
    │   │   │   │   │   │   └── connects-queries.ts  # Connect queries
    │   │   │   │   │   └── types.ts  # Connect types
    │   │   │   │   ├── contracts/  # Contracts feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── contracts-api.ts
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useContract.ts
    │   │   │   │   │   │   ├── useContracts.ts
    │   │   │   │   │   │   ├── useDeliverables.ts
    │   │   │   │   │   │   ├── useDisputes.ts
    │   │   │   │   │   │   ├── useMilestones.ts
    │   │   │   │   │   │   ├── useTimesheets.ts
    │   │   │   │   │   │   └── useWorkDiary.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── contracts-mutations.ts
    │   │   │   │   │   │   └── contracts-queries.ts  # BE: contracts-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── deliverables/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── deliverables-api.ts  # Deliverables API client
    │   │   │   │   │   │       # BE: contracts-be/deliverable, storage-be/asset
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useDeliverable.ts  # Single deliverable
    │   │   │   │   │   │   ├── useDeliverableRevisions.ts  # Revision management
    │   │   │   │   │   │   ├── useDeliverables.ts  # Deliverables list
    │   │   │   │   │   │   ├── useReviewDeliverable.ts  # Review deliverable (client)
    │   │   │   │   │   │   └── useUploadDeliverable.ts  # Upload deliverable
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── deliverables-mutations.ts  # Deliverable mutations
    │   │   │   │   │   │   └── deliverables-queries.ts  # Deliverable queries
    │   │   │   │   │   └── types.ts  # Deliverable types
    │   │   │   │   ├── feature-flags/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── flags-api.ts  # Feature flags API
    │   │   │   │   │   │       # BE: utility/flags
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useFeatureFlag.ts  # Check single flag
    │   │   │   │   │   │   ├── useFeatureFlags.ts  # Get all flags
    │   │   │   │   │   │   └── useFeatureFlagVariant.ts  # A/B test variant
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   └── flags-queries.ts  # Flag queries
    │   │   │   │   │   ├── store/
    │   │   │   │   │   │   └── flags-store.ts  # Flags state (Zustand)
    │   │   │   │   │   └── types.ts  # Flag types
    │   │   │   │   ├── financial/  # Financial feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── financial-api.ts
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useEscrow.ts
    │   │   │   │   │   │   ├── useInvoices.ts
    │   │   │   │   │   │   ├── usePaymentMethods.ts
    │   │   │   │   │   │   ├── usePayoutMethods.ts
    │   │   │   │   │   │   ├── useTransactions.ts
    │   │   │   │   │   │   └── useWallet.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── financial-mutations.ts
    │   │   │   │   │   │   └── financial-queries.ts  # BE: financial-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── interviews/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── interviews-api.ts  # Interviews API client
    │   │   │   │   │   │       # BE: proposals-be/interview
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useInterview.ts  # Single interview
    │   │   │   │   │   │   ├── useInterviewFeedback.ts  # Interview feedback
    │   │   │   │   │   │   ├── useInterviews.ts  # Interviews list
    │   │   │   │   │   │   └── useScheduleInterview.ts  # Schedule interview
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── interviews-mutations.ts  # Interview mutations
    │   │   │   │   │   │   └── interviews-queries.ts  # Interview queries
    │   │   │   │   │   └── types.ts  # Interview types
    │   │   │   │   ├── invitations/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── job-invitations-api.ts  # Job invitations API (client)
    │   │   │   │   │   │   │   # BE: jobs-be/invitation
    │   │   │   │   │   │   └── proposal-invites-api.ts  # Proposal invites API (freelancer)
    │   │   │   │   │   │       # BE: proposals-be/invite
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useAcceptInvite.ts  # Accept invite (freelancer)
    │   │   │   │   │   │   ├── useDeclineInvite.ts  # Decline invite (freelancer)
    │   │   │   │   │   │   ├── useInvitationAnalytics.ts  # Invitation metrics
    │   │   │   │   │   │   ├── useInvitations.ts  # Invitations management
    │   │   │   │   │   │   └── useSendInvitation.ts  # Send invitation (client)
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── invitations-mutations.ts  # Invitation mutations
    │   │   │   │   │   │   └── invitations-queries.ts  # Invitation queries
    │   │   │   │   │   └── types.ts  # Invitation types
    │   │   │   │   ├── jobs/  # Jobs feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── jobs-api.ts  # Jobs API client
    │   │   │   │   │   │       # Axios/Fetch wrapper
    │   │   │   │   │   │       # All jobs-be endpoints
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useCreateJob.ts  # Create job mutation
    │   │   │   │   │   │   ├── useDeleteJob.ts  # Delete job mutation
    │   │   │   │   │   │   ├── useJob.ts  # Single job query
    │   │   │   │   │   │   ├── useJobFilters.ts  # Job filters state
    │   │   │   │   │   │   ├── useJobs.ts  # List jobs query
    │   │   │   │   │   │   ├── useSaveJob.ts  # Save/bookmark job
    │   │   │   │   │   │   └── useUpdateJob.ts  # Update job mutation
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── jobs-mutations.ts  # TanStack Query mutations
    │   │   │   │   │   │   │   # BE: jobs-be
    │   │   │   │   │   │   │   # POST /v1/jobs
    │   │   │   │   │   │   │   # PUT /v1/jobs/{id}
    │   │   │   │   │   │   └── jobs-queries.ts  # TanStack Query queries
    │   │   │   │   │   │       # BE: jobs-be
    │   │   │   │   │   │       # GET /v1/jobs
    │   │   │   │   │   │       # GET /v1/jobs/{id}
    │   │   │   │   │   ├── utils/
    │   │   │   │   │   │   ├── job-helpers.ts  # Job utility functions
    │   │   │   │   │   │   └── job-validation.ts  # Job validation (Zod)
    │   │   │   │   │   └── types.ts  # Jobs types
    │   │   │   │   ├── learning/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── learning-paths-api.ts  # Learning paths API
    │   │   │   │   │   │   │   # BE: users-be/learning_path
    │   │   │   │   │   │   └── mentorship-api.ts  # Mentorship API
    │   │   │   │   │   │       # BE: users-be/mentorship
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useAchievements.ts  # Achievements/badges
    │   │   │   │   │   │   ├── useLearningPath.ts  # Single learning path
    │   │   │   │   │   │   ├── useLearningPaths.ts  # Learning paths list
    │   │   │   │   │   │   ├── useLearningProgress.ts  # Track progress
    │   │   │   │   │   │   └── useMentorship.ts  # Mentorship management
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── learning-mutations.ts  # Learning mutations
    │   │   │   │   │   │   └── learning-queries.ts  # Learning queries
    │   │   │   │   │   └── types.ts  # Learning types
    │   │   │   │   ├── messages/  # Messaging feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── messages-api.ts
    │   │   │   │   │   │   └── websocket.ts  # WebSocket client
    │   │   │   │   │   │       # WS: ws://communications-be/v1/realtime
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useConversation.ts
    │   │   │   │   │   │   ├── useConversations.ts
    │   │   │   │   │   │   ├── useMessages.ts
    │   │   │   │   │   │   ├── useRealtimeMessages.ts  # WebSocket
    │   │   │   │   │   │   └── useSendMessage.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── messages-mutations.ts
    │   │   │   │   │   │   └── messages-queries.ts  # BE: communications-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── negotiations/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── negotiations-api.ts  # Negotiations API client
    │   │   │   │   │   │       # BE: proposals-be/negotiation
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useNegotiation.ts  # Single negotiation
    │   │   │   │   │   │   ├── useNegotiationHistory.ts  # Negotiation history
    │   │   │   │   │   │   └── useNegotiationOffer.ts  # Make/accept/reject offer
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── negotiations-mutations.ts  # Negotiation mutations
    │   │   │   │   │   │   └── negotiations-queries.ts  # Negotiation queries
    │   │   │   │   │   └── types.ts  # Negotiation types
    │   │   │   │   ├── networking/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── connections-api.ts  # Connections API
    │   │   │   │   │   │   │   # BE: users-be/connection
    │   │   │   │   │   │   ├── groups-api.ts  # Groups API
    │   │   │   │   │   │   │   # BE: users-be/user_group
    │   │   │   │   │   │   └── referrals-api.ts  # Referrals API
    │   │   │   │   │   │       # BE: users-be/referral
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useConnectionRequest.ts  # Send/accept/reject
    │   │   │   │   │   │   ├── useConnections.ts  # Connections management
    │   │   │   │   │   │   ├── useGroups.ts  # Groups management
    │   │   │   │   │   │   ├── useNetworkRecommendations.ts  # Connection recommendations
    │   │   │   │   │   │   └── useReferrals.ts  # Referral management
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── networking-mutations.ts  # Networking mutations
    │   │   │   │   │   │   └── networking-queries.ts  # Networking queries
    │   │   │   │   │   └── types.ts  # Networking types
    │   │   │   │   ├── notifications/  # Notifications feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── notifications-api.ts
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useMarkAsRead.ts
    │   │   │   │   │   │   ├── useNotifications.ts
    │   │   │   │   │   │   ├── useRealtimeNotifications.ts  # WebSocket
    │   │   │   │   │   │   └── useUnreadCount.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── notifications-mutations.ts
    │   │   │   │   │   │   └── notifications-queries.ts  # BE: communications-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── profile/  # Profile feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── profile-api.ts
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useEducation.ts
    │   │   │   │   │   │   ├── useExperience.ts
    │   │   │   │   │   │   ├── usePortfolio.ts
    │   │   │   │   │   │   ├── useProfile.ts
    │   │   │   │   │   │   ├── useServiceCatalog.ts
    │   │   │   │   │   │   ├── useSkills.ts
    │   │   │   │   │   │   └── useUpdateProfile.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── profile-mutations.ts
    │   │   │   │   │   │   └── profile-queries.ts  # BE: users-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── proposals/  # Proposals feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── proposals-api.ts
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useBidding.ts  # Bidding hooks
    │   │   │   │   │   │   ├── useProposal.ts
    │   │   │   │   │   │   ├── useProposals.ts
    │   │   │   │   │   │   ├── useSubmitProposal.ts
    │   │   │   │   │   │   └── useWithdrawProposal.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── proposals-mutations.ts
    │   │   │   │   │   │   └── proposals-queries.ts  # BE: proposals-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── realtime/
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── usePresence.ts  # User presence (online/offline)
    │   │   │   │   │   │   ├── useRealtimeAuction.ts  # Real-time auction updates
    │   │   │   │   │   │   ├── useRealtimeMessages.ts  # Real-time messages
    │   │   │   │   │   │   ├── useRealtimeNotifications.ts  # Real-time notifications
    │   │   │   │   │   │   └── useWebSocket.ts  # WebSocket connection
    │   │   │   │   │   ├── store/
    │   │   │   │   │   │   └── realtime-store.ts  # Real-time state (Zustand)
    │   │   │   │   │   ├── websocket/
    │   │   │   │   │   │   ├── client.ts  # WebSocket client
    │   │   │   │   │   │   ├── heartbeat.ts  # Connection health
    │   │   │   │   │   │   └── reconnection.ts  # Reconnection logic
    │   │   │   │   │   └── types.ts  # Real-time types
    │   │   │   │   ├── reviews/  # Reviews feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── reviews-api.ts
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useBadges.ts
    │   │   │   │   │   │   ├── useCreateReview.ts
    │   │   │   │   │   │   ├── useReviews.ts
    │   │   │   │   │   │   └── useReviewStats.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── reviews-mutations.ts
    │   │   │   │   │   │   └── reviews-queries.ts  # BE: reviews-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── search/  # Search feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   ├── recommendations-api.ts  # Recommendations API
    │   │   │   │   │   │   │   # BE: search-be/recommendation
    │   │   │   │   │   │   ├── saved-searches-api.ts  # Saved searches API
    │   │   │   │   │   │   │   # BE: search-be/saved-search
    │   │   │   │   │   │   ├── search-api.ts  # Search API (already may exist, ensuring completeness)
    │   │   │   │   │   │   │   # BE: search-be/query
    │   │   │   │   │   │   └── trending-api.ts  # Trending API
    │   │   │   │   │   │       # BE: search-be/trending
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useFreelancerSearch.ts
    │   │   │   │   │   │   ├── useJobSearch.ts
    │   │   │   │   │   │   ├── useRecommendations.ts  # Recommendations
    │   │   │   │   │   │   ├── useSavedSearches.ts  # Saved searches
    │   │   │   │   │   │   ├── useSearch.ts  # Search execution
    │   │   │   │   │   │   ├── useSearchHistory.ts  # Search history
    │   │   │   │   │   │   ├── useSearchSuggestions.ts  # Auto-complete suggestions
    │   │   │   │   │   │   └── useTrending.ts  # Trending items
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── search-mutations.ts  # Search mutations
    │   │   │   │   │   │   └── search-queries.ts  # BE: search-be
    │   │   │   │   │   │       # Search queries
    │   │   │   │   │   ├── store/
    │   │   │   │   │   │   └── search-store.ts  # Search UI state (filters, etc.)
    │   │   │   │   │   └── types.ts  # Search types
    │   │   │   │   ├── shortlists/
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── shortlists-api.ts  # Shortlists API
    │   │   │   │   │   │       # BE: jobs-be/shortlist
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useAddToShortlist.ts  # Add candidate
    │   │   │   │   │   │   ├── useRemoveFromShortlist.ts  # Remove candidate
    │   │   │   │   │   │   ├── useShortlist.ts  # Single shortlist
    │   │   │   │   │   │   └── useShortlists.ts  # Shortlists management
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── shortlists-mutations.ts  # Shortlist mutations
    │   │   │   │   │   │   └── shortlists-queries.ts  # Shortlist queries
    │   │   │   │   │   └── types.ts  # Shortlist types
    │   │   │   │   ├── storage/  # File storage feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── storage-api.ts  # BE: storage-be
    │   │   │   │   │   │       # POST /v1/storage/upload
    │   │   │   │   │   │       # GET /v1/storage/presign
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useFileDownload.ts
    │   │   │   │   │   │   ├── usePresignedUrl.ts
    │   │   │   │   │   │   └── useUpload.ts
    │   │   │   │   │   └── types.ts
    │   │   │   │   ├── subscriptions/  # Subscriptions feature
    │   │   │   │   │   ├── api/
    │   │   │   │   │   │   └── subscriptions-api.ts
    │   │   │   │   │   ├── hooks/
    │   │   │   │   │   │   ├── useConnects.ts
    │   │   │   │   │   │   ├── useEntitlements.ts
    │   │   │   │   │   │   ├── usePlans.ts
    │   │   │   │   │   │   ├── useSubscription.ts
    │   │   │   │   │   │   └── useUpgrade.ts
    │   │   │   │   │   ├── queries/
    │   │   │   │   │   │   ├── subscriptions-mutations.ts
    │   │   │   │   │   │   └── subscriptions-queries.ts  # BE: subscriptions-be
    │   │   │   │   │   └── types.ts
    │   │   │   │   └── work-tracking/
    │   │   │   │       ├── api/
    │   │   │   │       │   ├── timesheet-api.ts  # Timesheet API
    │   │   │   │       │   │   # BE: contracts-be/timesheet
    │   │   │   │       │   └── work-diary-api.ts  # Work diary API
    │   │   │   │       │       # BE: contracts-be/work_diary
    │   │   │   │       ├── hooks/
    │   │   │   │       │   ├── useApproveTimesheet.ts  # Approve timesheet (client)
    │   │   │   │       │   ├── useTimesheet.ts  # Timesheet management
    │   │   │   │       │   ├── useTimeTracking.ts  # Real-time time tracking
    │   │   │   │       │   └── useWorkDiary.ts  # Work diary entries
    │   │   │   │       ├── queries/
    │   │   │   │       │   ├── timesheet-mutations.ts  # Timesheet mutations
    │   │   │   │       │   ├── timesheet-queries.ts  # Timesheet queries
    │   │   │   │       │   ├── work-diary-mutations.ts  # Work diary mutations
    │   │   │   │       │   └── work-diary-queries.ts  # Work diary queries
    │   │   │   │       ├── store/
    │   │   │   │       │   └── time-tracker-store.ts  # Time tracker state (Zustand)
    │   │   │   │       └── types.ts  # Work tracking types
    │   │   │   ├── flags/
    │   │   │   │   ├── flags-client.ts  # Feature flags API client
    │   │   │   │   │   # BE: utility/flags
    │   │   │   │   │   # GET /v1/flags/user
    │   │   │   │   │   # GET /v1/flags/organization
    │   │   │   │   └── types.ts
    │   │   │   ├── hooks/  # General hooks
    │   │   │   │   ├── useClickOutside.ts
    │   │   │   │   ├── useCopyToClipboard.ts
    │   │   │   │   ├── useDebounce.ts
    │   │   │   │   ├── useIntersectionObserver.ts
    │   │   │   │   ├── useLocalStorage.ts
    │   │   │   │   ├── useMediaQuery.ts
    │   │   │   │   └── useToggle.ts
    │   │   │   ├── i18n/  # Internationalization
    │   │   │   │   ├── locales/  # Locale files
    │   │   │   │   │   ├── ar/  # Arabic
    │   │   │   │   │   ├── de/  # German
    │   │   │   │   │   ├── en/
    │   │   │   │   │   │   ├── admin.json
    │   │   │   │   │   │   ├── auth.json
    │   │   │   │   │   │   ├── common.json
    │   │   │   │   │   │   ├── contracts.json
    │   │   │   │   │   │   ├── errors.json
    │   │   │   │   │   │   ├── financial.json
    │   │   │   │   │   │   ├── jobs.json
    │   │   │   │   │   │   ├── messages.json
    │   │   │   │   │   │   ├── profile.json
    │   │   │   │   │   │   ├── proposals.json
    │   │   │   │   │   │   ├── reviews.json
    │   │   │   │   │   │   ├── settings.json
    │   │   │   │   │   │   └── subscription.json
    │   │   │   │   │   ├── es/  # Spanish
    │   │   │   │   │   ├── fr/  # French
    │   │   │   │   │   ├── hi/  # Hindi
    │   │   │   │   │   ├── ru/  # Russian
    │   │   │   │   │   ├── tr/  # Turkish
    │   │   │   │   │   └── zh/  # Chinese
    │   │   │   │   └── config.ts  # i18n configuration
    │   │   │   ├── incidents/
    │   │   │   │   ├── incidents-client.ts  # Incidents API client
    │   │   │   │   │   # BE: utility/status
    │   │   │   │   │   # GET /v1/status/incidents
    │   │   │   │   │   # POST /v1/status/incidents
    │   │   │   │   │   # PUT /v1/status/incidents/{id}
    │   │   │   │   │   # GET /v1/status/maintenance
    │   │   │   │   └── types.ts
    │   │   │   ├── learning/
    │   │   │   │   ├── learning-client.ts  # Learning API client
    │   │   │   │   │   # BE: learning-be (if exists) or external LMS
    │   │   │   │   │   # GET /v1/learning/courses
    │   │   │   │   │   # GET /v1/learning/courses/{id}
    │   │   │   │   │   # POST /v1/learning/assessments/{id}/submit
    │   │   │   │   └── types.ts
    │   │   │   ├── lib/  # Shared utilities
    │   │   │   │   ├── api/
    │   │   │   │   │   ├── client.ts  # Axios/Fetch client setup
    │   │   │   │   │   │   # - Base URL
    │   │   │   │   │   │   # - Auth interceptors
    │   │   │   │   │   │   # - Error handling
    │   │   │   │   │   │   # - Request/response logging
    │   │   │   │   │   ├── endpoints.ts  # API endpoint constants
    │   │   │   │   │   └── error-handler.ts  # Global error handler
    │   │   │   │   ├── constants/
    │   │   │   │   │   ├── api.ts  # API constants
    │   │   │   │   │   ├── app.ts  # App constants
    │   │   │   │   │   ├── permissions.ts  # RBAC permissions
    │   │   │   │   │   └── routes.ts  # Route constants
    │   │   │   │   ├── formatting/
    │   │   │   │   │   ├── currency.ts  # Currency formatting
    │   │   │   │   │   ├── date.ts  # Date formatting
    │   │   │   │   │   ├── number.ts  # Number formatting
    │   │   │   │   │   └── string.ts  # String utilities
    │   │   │   │   ├── validation/
    │   │   │   │   │   ├── schemas.ts  # Zod schemas
    │   │   │   │   │   └── validators.ts  # Custom validators
    │   │   │   │   └── websocket/
    │   │   │   │       ├── client.ts  # WebSocket client
    │   │   │   │       │   # WS: ws://communications-be/v1/realtime
    │   │   │   │       │   # - Connection management
    │   │   │   │       │   # - Reconnection logic
    │   │   │   │       │   # - Event subscriptions
    │   │   │   │       └── events.ts  # WebSocket event types
    │   │   │   ├── moderation/
    │   │   │   │   ├── moderation-client.ts  # Moderation API client
    │   │   │   │   │   # BE: admin-be/moderation (if exists)
    │   │   │   │   │   # POST /v1/moderation/report
    │   │   │   │   │   # POST /v1/moderation/content-check
    │   │   │   │   └── types.ts
    │   │   │   ├── monitoring/
    │   │   │   │   ├── analytics/
    │   │   │   │   │   ├── providers/
    │   │   │   │   │   │   ├── amplitude.ts  # Amplitude
    │   │   │   │   │   │   ├── google-analytics.ts  # GA4
    │   │   │   │   │   │   └── mixpanel.ts  # Mixpanel
    │   │   │   │   │   ├── tracking/
    │   │   │   │   │   │   ├── event-tracker.ts
    │   │   │   │   │   │   ├── page-tracker.ts
    │   │   │   │   │   │   └── user-tracker.ts
    │   │   │   │   │   └── analytics-config.ts  # Analytics configuration
    │   │   │   │   ├── logging/
    │   │   │   │   │   ├── transports/
    │   │   │   │   │   │   ├── console-transport.ts
    │   │   │   │   │   │   ├── file-transport.ts  # Mobile only
    │   │   │   │   │   │   └── remote-transport.ts
    │   │   │   │   │   ├── log-levels.ts  # Log level configuration
    │   │   │   │   │   └── logger.ts  # Application logger
    │   │   │   │   ├── performance/
    │   │   │   │   │   ├── api-monitor.ts  # API performance monitoring
    │   │   │   │   │   ├── bundle-monitor.ts  # Bundle size monitoring
    │   │   │   │   │   ├── performance-observer.ts  # Performance API
    │   │   │   │   │   └── web-vitals.ts  # Web Vitals monitoring
    │   │   │   │   └── sentry/
    │   │   │   │       ├── error-boundary.tsx  # Sentry error boundary
    │   │   │   │       ├── sentry-config.ts  # Sentry configuration
    │   │   │   │       ├── sentry-init.native.ts  # Mobile initialization
    │   │   │   │       └── sentry-init.web.ts  # Web initialization
    │   │   │   ├── presence/
    │   │   │   │   ├── presence-client.ts  # Presence API client
    │   │   │   │   │   # BE: communications-be/presence (if exists)
    │   │   │   │   │   # POST /v1/presence/heartbeat
    │   │   │   │   │   # GET /v1/presence/users/{id}
    │   │   │   │   └── types.ts
    │   │   │   ├── security/
    │   │   │   │   ├── auth/
    │   │   │   │   │   ├── device-trust.ts  # Device trust
    │   │   │   │   │   ├── session-manager.ts  # Session management
    │   │   │   │   │   ├── token-manager.ts  # Token management
    │   │   │   │   │   └── token-refresh.ts  # Auto token refresh
    │   │   │   │   ├── encryption/
    │   │   │   │   │   ├── crypto-utils.ts  # Encryption utilities
    │   │   │   │   │   ├── key-management.ts  # Key management
    │   │   │   │   │   └── secure-storage.ts  # Secure storage
    │   │   │   │   ├── headers/
    │   │   │   │   │   ├── cors.ts  # CORS configuration
    │   │   │   │   │   ├── csp.ts  # Content Security Policy
    │   │   │   │   │   └── security-headers.ts  # Security headers config
    │   │   │   │   ├── monitoring/
    │   │   │   │   │   ├── anomaly-detector.ts  # Anomaly detection
    │   │   │   │   │   ├── breach-detector.ts  # Breach detection
    │   │   │   │   │   └── security-monitor.ts  # Security monitoring
    │   │   │   │   └── validation/
    │   │   │   │       ├── input-validator.ts  # Input validation
    │   │   │   │       ├── sanitizer.ts  # HTML/input sanitization
    │   │   │   │       ├── sql-injection-guard.ts  # SQL injection protection
    │   │   │   │       └── xss-protection.ts  # XSS protection
    │   │   │   ├── sourcing/
    │   │   │   │   ├── sourcing-client.ts  # Sourcing API client
    │   │   │   │   │   # BE: jobs-be/campaign (if exists) or jobs-be/job
    │   │   │   │   │   # GET /v1/sourcing/campaigns
    │   │   │   │   │   # POST /v1/sourcing/campaigns
    │   │   │   │   │   # GET /v1/sourcing/talent-pools
    │   │   │   │   └── types.ts
    │   │   │   └── webhooks/
    │   │   │       ├── types.ts
    │   │   │       └── webhooks-client.ts  # Webhooks API client
    │   │   │           # BE: utility/webhooks (if exists) or admin-be
    │   │   │           # GET /v1/webhooks
    │   │   │           # POST /v1/webhooks
    │   │   │           # PUT /v1/webhooks/{id}
    │   │   │           # DELETE /v1/webhooks/{id}
    │   │   ├── package.json
    │   │   ├── README.md
    │   │   └── tsconfig.json
    │   ├── testing/
    │   │   ├── factories/
    │   │   │   ├── ...  # Other factories
    │   │   │   ├── job-factory.ts  # Job factory
    │   │   │   ├── proposal-factory.ts  # Proposal factory
    │   │   │   └── user-factory.ts  # User factory
    │   │   ├── fixtures/
    │   │   │   ├── ...  # Other fixtures
    │   │   │   ├── auth.ts  # Auth fixtures
    │   │   │   ├── jobs.ts  # Job fixtures
    │   │   │   └── users.ts  # User fixtures
    │   │   ├── mocks/
    │   │   │   ├── api/
    │   │   │   │   ├── ...  # Other service mocks
    │   │   │   │   ├── communications-mock.ts  # Mock communications API
    │   │   │   │   ├── contracts-mock.ts  # Mock contracts API
    │   │   │   │   ├── financial-mock.ts  # Mock financial API
    │   │   │   │   ├── jobs-mock.ts  # Mock jobs API
    │   │   │   │   ├── proposals-mock.ts  # Mock proposals API
    │   │   │   │   └── users-mock.ts  # Mock users API
    │   │   │   ├── data/
    │   │   │   │   ├── ...  # Other mock data
    │   │   │   │   ├── jobs.ts  # Mock job data
    │   │   │   │   ├── proposals.ts  # Mock proposal data
    │   │   │   │   └── users.ts  # Mock user data
    │   │   │   └── handlers/
    │   │   │       ├── ...  # Other MSW handlers
    │   │   │       ├── auth-handlers.ts  # MSW auth handlers
    │   │   │       └── users-handlers.ts  # MSW users handlers
    │   │   ├── setup/
    │   │   │   ├── msw-setup.ts  # MSW server setup
    │   │   │   ├── test-providers.tsx  # Test providers wrapper
    │   │   │   └── test-setup.ts  # Test environment setup
    │   │   └── utils/
    │   │       ├── accessibility-checker.ts  # A11y testing utils
    │   │       ├── render-with-providers.tsx  # Custom render
    │   │       └── wait-for-async.ts  # Async utilities
    │   ├── types/  # TypeScript type definitions
    │   │   # Shared TypeScript types
    │   │   ├── src/
    │   │   │   ├── api/  # API response types
    │   │   │   │   ├── performance/
    │   │   │   │   │   └── index.ts  # Performance types
    │   │   │   │   │       # - Web Vitals
    │   │   │   │   │       # - Performance metric
    │   │   │   │   │       # - Error event
    │   │   │   │   ├── realtime/
    │   │   │   │   │   └── index.ts  # Real-time types
    │   │   │   │   │       # - WebSocket message
    │   │   │   │   │       # - Presence update
    │   │   │   │   │       # - Event subscription
    │   │   │   │   ├── webhooks/
    │   │   │   │   │   └── index.ts  # Webhook types
    │   │   │   │   │       # - Webhook event
    │   │   │   │   │       # - Webhook subscription
    │   │   │   │   │       # - Webhook delivery
    │   │   │   │   ├── admin.ts
    │   │   │   │   ├── common.ts  # Common types (pagination, filters, etc.)
    │   │   │   │   ├── contracts.ts
    │   │   │   │   ├── financial.ts
    │   │   │   │   ├── jobs.ts
    │   │   │   │   ├── messages.ts
    │   │   │   │   ├── notifications.ts
    │   │   │   │   ├── proposals.ts
    │   │   │   │   ├── review.ts
    │   │   │   │   ├── search.ts
    │   │   │   │   ├── storage.ts
    │   │   │   │   ├── subscriptions.ts
    │   │   │   │   └── users.ts
    │   │   │   ├── domains/
    │   │   │   │   ├── achievements/
    │   │   │   │   │   └── index.ts  # Achievement types
    │   │   │   │   │       # - Achievement definition
    │   │   │   │   │       # - User achievement
    │   │   │   │   │       # - Badge
    │   │   │   │   │       # - Points
    │   │   │   │   ├── activity/
    │   │   │   │   │   └── index.ts  # Activity feed types
    │   │   │   │   │       # - Activity event
    │   │   │   │   │       # - Activity feed item
    │   │   │   │   │       # - Activity filters
    │   │   │   │   ├── compliance/
    │   │   │   │   │   └── index.ts  # Compliance types
    │   │   │   │   │       # - KYC case
    │   │   │   │   │       # - Business verification
    │   │   │   │   │       # - Document verification
    │   │   │   │   │       # - Privacy request
    │   │   │   │   ├── experiments/
    │   │   │   │   │   └── index.ts  # Experiment types
    │   │   │   │   │       # - Experiment definition
    │   │   │   │   │       # - Variant
    │   │   │   │   │       # - User assignment
    │   │   │   │   ├── flags/
    │   │   │   │   │   └── index.ts  # Feature flag types
    │   │   │   │   │       # - Flag definition
    │   │   │   │   │       # - Flag value
    │   │   │   │   │       # - Targeting rules
    │   │   │   │   ├── incidents/
    │   │   │   │   │   └── index.ts  # Incident types
    │   │   │   │   │       # - Incident
    │   │   │   │   │       # - Maintenance window
    │   │   │   │   │       # - System health
    │   │   │   │   ├── learning/
    │   │   │   │   │   └── index.ts  # Learning types
    │   │   │   │   │       # - Course
    │   │   │   │   │       # - Lesson
    │   │   │   │   │       # - Assessment
    │   │   │   │   │       # - Progress
    │   │   │   │   ├── moderation/
    │   │   │   │   │   └── index.ts  # Moderation types
    │   │   │   │   │       # - Report
    │   │   │   │   │       # - Moderation action
    │   │   │   │   │       # - Content check
    │   │   │   │   ├── presence/
    │   │   │   │   │   └── index.ts  # Presence types
    │   │   │   │   │       # - Online status
    │   │   │   │   │       # - Last seen
    │   │   │   │   │       # - Typing indicator
    │   │   │   │   └── sourcing/
    │   │   │   │       └── index.ts  # Sourcing types
    │   │   │   │           # - Campaign
    │   │   │   │           # - Talent pool
    │   │   │   │           # - Invitation
    │   │   │   ├── entities/  # Domain entities
    │   │   │   │   ├── contract.ts
    │   │   │   │   ├── invoice.ts
    │   │   │   │   ├── job.ts
    │   │   │   │   ├── message.ts
    │   │   │   │   ├── notification.ts
    │   │   │   │   ├── proposal.ts
    │   │   │   │   ├── review.ts
    │   │   │   │   ├── subscription.ts
    │   │   │   │   ├── transaction.ts
    │   │   │   │   └── user.ts
    │   │   │   ├── enums/  # Enums
    │   │   │   │   ├── contract-status.ts
    │   │   │   │   ├── job-status.ts
    │   │   │   │   ├── payment-status.ts
    │   │   │   │   ├── proposal-status.ts
    │   │   │   │   ├── review-rating.ts
    │   │   │   │   └── user-type.ts
    │   │   │   ├── shared/
    │   │   │   │   ├── analytics/
    │   │   │   │   │   └── index.ts  # Analytics types
    │   │   │   │   │       # - Event
    │   │   │   │   │       # - Page view
    │   │   │   │   │       # - Conversion
    │   │   │   │   └── geolocation/
    │   │   │   │       └── index.ts  # Geolocation types
    │   │   │   │           # - Coordinates
    │   │   │   │           # - Location
    │   │   │   │           # - Distance
    │   │   │   └── index.ts  # Export all types
    │   │   ├── package.json
    │   │   ├── README.md
    │   │   └── tsconfig.json
    │   └── ui/  # Cross-platform component library
    │       # Cross-platform UI component library
    │       ├── src/
    │       │   ├── a11y/
    │       │   │   ├── Announcer/
    │       │   │   │   ├── LiveAnnouncer.native.tsx
    │       │   │   │   ├── LiveAnnouncer.tsx  # Live region announcements
    │       │   │   │   ├── LiveAnnouncer.types.ts
    │       │   │   │   └── LiveAnnouncer.web.tsx
    │       │   │   ├── FocusTrap/
    │       │   │   │   ├── FocusTrap.tsx
    │       │   │   │   ├── FocusTrap.types.ts
    │       │   │   │   └── FocusTrap.web.tsx
    │       │   │   ├── SkipLink/
    │       │   │   │   ├── SkipLink.tsx
    │       │   │   │   ├── SkipLink.types.ts
    │       │   │   │   └── SkipLink.web.tsx
    │       │   │   └── VisuallyHidden/
    │       │   │       ├── VisuallyHidden.native.tsx
    │       │   │       ├── VisuallyHidden.tsx
    │       │   │       ├── VisuallyHidden.types.ts
    │       │   │       └── VisuallyHidden.web.tsx
    │       │   ├── components/
    │       │   │   ├── Accordion/
    │       │   │   ├── Alert/
    │       │   │   ├── auction/
    │       │   │   │   ├── AuctionTimer.native.tsx
    │       │   │   │   ├── AuctionTimer.tsx  # Countdown timer
    │       │   │   │   ├── AuctionTimer.web.tsx
    │       │   │   │   ├── BidHistoryChart.native.tsx
    │       │   │   │   ├── BidHistoryChart.tsx  # Bid history visualization
    │       │   │   │   ├── BidHistoryChart.web.tsx
    │       │   │   │   ├── LiveBidFeed.native.tsx
    │       │   │   │   ├── LiveBidFeed.tsx  # Real-time bid feed
    │       │   │   │   └── LiveBidFeed.web.tsx
    │       │   │   ├── Avatar/
    │       │   │   ├── Badge/
    │       │   │   ├── Breadcrumb/
    │       │   │   ├── Button/
    │       │   │   │   ├── Button.native.tsx  # Native-specific overrides
    │       │   │   │   ├── Button.stories.tsx  # Storybook stories
    │       │   │   │   ├── Button.test.tsx  # Component tests
    │       │   │   │   ├── Button.tsx  # Base button component
    │       │   │   │   └── Button.web.tsx  # Web-specific overrides
    │       │   │   ├── Card/
    │       │   │   ├── charts/
    │       │   │   │   ├── EarningsChart.native.tsx
    │       │   │   │   ├── EarningsChart.tsx  # Earnings visualization
    │       │   │   │   ├── EarningsChart.web.tsx
    │       │   │   │   ├── PerformanceChart.native.tsx
    │       │   │   │   ├── PerformanceChart.tsx  # Performance metrics
    │       │   │   │   ├── PerformanceChart.web.tsx
    │       │   │   │   ├── TrendChart.native.tsx
    │       │   │   │   ├── TrendChart.tsx  # Trend visualization
    │       │   │   │   └── TrendChart.web.tsx
    │       │   │   ├── Checkbox/
    │       │   │   ├── CodeEditor/
    │       │   │   │   ├── CodeEditor.native.tsx
    │       │   │   │   ├── CodeEditor.tsx
    │       │   │   │   ├── CodeEditor.types.ts
    │       │   │   │   └── CodeEditor.web.tsx  # Web (e.g., Monaco/CodeMirror)
    │       │   │   ├── collaboration/
    │       │   │   │   ├── CollaborationPanel.native.tsx
    │       │   │   │   ├── CollaborationPanel.tsx  # Team collaboration
    │       │   │   │   ├── CollaborationPanel.web.tsx
    │       │   │   │   ├── GroupCard.native.tsx
    │       │   │   │   ├── GroupCard.tsx  # User group card
    │       │   │   │   ├── GroupCard.web.tsx
    │       │   │   │   ├── MentorCard.native.tsx
    │       │   │   │   ├── MentorCard.tsx  # Mentor profile card
    │       │   │   │   └── MentorCard.web.tsx
    │       │   │   ├── compliance/
    │       │   │   │   ├── DocumentUploader.native.tsx
    │       │   │   │   ├── DocumentUploader.tsx  # Compliance doc uploader
    │       │   │   │   ├── DocumentUploader.web.tsx
    │       │   │   │   ├── VerificationStatus.native.tsx
    │       │   │   │   ├── VerificationStatus.tsx  # Verification status badge
    │       │   │   │   └── VerificationStatus.web.tsx
    │       │   │   ├── DataDisplay/
    │       │   │   │   ├── Kanban/
    │       │   │   │   │   ├── KanbanBoard.native.tsx  # Touch gestures
    │       │   │   │   │   ├── KanbanBoard.tsx
    │       │   │   │   │   ├── KanbanBoard.types.ts
    │       │   │   │   │   └── KanbanBoard.web.tsx  # Drag & drop
    │       │   │   │   ├── List/
    │       │   │   │   │   ├── VirtualList/
    │       │   │   │   │   │   ├── VirtualList.native.tsx  # FlashList
    │       │   │   │   │   │   ├── VirtualList.tsx
    │       │   │   │   │   │   └── VirtualList.web.tsx
    │       │   │   │   │   ├── List.native.tsx
    │       │   │   │   │   ├── List.tsx
    │       │   │   │   │   ├── List.types.ts
    │       │   │   │   │   └── List.web.tsx
    │       │   │   │   ├── Table/
    │       │   │   │   │   ├── DataTable/
    │       │   │   │   │   │   ├── DataTable.tsx  # With sorting/filtering
    │       │   │   │   │   │   └── DataTable.types.ts
    │       │   │   │   │   ├── VirtualTable/
    │       │   │   │   │   │   ├── VirtualTable.tsx  # Virtualized table
    │       │   │   │   │   │   └── VirtualTable.types.ts
    │       │   │   │   │   ├── Table.tsx  # Base table
    │       │   │   │   │   ├── Table.types.ts
    │       │   │   │   │   └── Table.web.tsx  # Full featured table
    │       │   │   │   └── Timeline/
    │       │   │   │       ├── Timeline.native.tsx
    │       │   │   │       ├── Timeline.tsx
    │       │   │   │       ├── Timeline.types.ts
    │       │   │   │       └── Timeline.web.tsx
    │       │   │   ├── DataTable/
    │       │   │   ├── DatePicker/
    │       │   │   ├── Dropdown/
    │       │   │   ├── Feedback/
    │       │   │   │   ├── Alert/
    │       │   │   │   │   ├── Alert.native.tsx
    │       │   │   │   │   ├── Alert.tsx
    │       │   │   │   │   ├── Alert.types.ts
    │       │   │   │   │   └── Alert.web.tsx
    │       │   │   │   ├── EmptyState/
    │       │   │   │   │   ├── EmptyState.native.tsx
    │       │   │   │   │   ├── EmptyState.tsx
    │       │   │   │   │   ├── EmptyState.types.ts
    │       │   │   │   │   └── EmptyState.web.tsx
    │       │   │   │   ├── Notification/
    │       │   │   │   │   ├── Notification.native.tsx
    │       │   │   │   │   ├── Notification.tsx
    │       │   │   │   │   ├── Notification.types.ts
    │       │   │   │   │   └── Notification.web.tsx  # Toast notifications
    │       │   │   │   ├── Progress/
    │       │   │   │   │   ├── ProgressBar/
    │       │   │   │   │   │   ├── ProgressBar.native.tsx
    │       │   │   │   │   │   ├── ProgressBar.tsx
    │       │   │   │   │   │   └── ProgressBar.web.tsx
    │       │   │   │   │   ├── ProgressCircle/
    │       │   │   │   │   │   ├── ProgressCircle.native.tsx
    │       │   │   │   │   │   ├── ProgressCircle.tsx
    │       │   │   │   │   │   └── ProgressCircle.web.tsx
    │       │   │   │   │   └── Progress.types.ts
    │       │   │   │   └── Skeleton/
    │       │   │   │       ├── Skeleton.native.tsx
    │       │   │   │       ├── Skeleton.tsx
    │       │   │   │       ├── Skeleton.types.ts
    │       │   │   │       └── Skeleton.web.tsx
    │       │   │   ├── FileUpload/
    │       │   │   │   ├── ImageUpload/
    │       │   │   │   │   ├── ImageCropper.tsx  # Image cropping
    │       │   │   │   │   ├── ImageUpload.native.tsx  # Camera/gallery
    │       │   │   │   │   ├── ImageUpload.tsx
    │       │   │   │   │   └── ImageUpload.web.tsx
    │       │   │   │   ├── MultiFileUpload/
    │       │   │   │   │   ├── MultiFileUpload.native.tsx
    │       │   │   │   │   ├── MultiFileUpload.tsx
    │       │   │   │   │   └── MultiFileUpload.web.tsx
    │       │   │   │   ├── FileUpload.native.tsx  # Document picker
    │       │   │   │   ├── FileUpload.tsx  # Base file upload
    │       │   │   │   ├── FileUpload.types.ts
    │       │   │   │   └── FileUpload.web.tsx  # Drag & drop support
    │       │   │   ├── Forms/
    │       │   │   │   ├── Checkbox/
    │       │   │   │   │   ├── Checkbox.native.tsx
    │       │   │   │   │   ├── Checkbox.tsx
    │       │   │   │   │   ├── Checkbox.types.ts
    │       │   │   │   │   └── Checkbox.web.tsx
    │       │   │   │   ├── FormField/
    │       │   │   │   │   ├── FormError.tsx
    │       │   │   │   │   ├── FormField.native.tsx
    │       │   │   │   │   ├── FormField.tsx  # Form field wrapper
    │       │   │   │   │   ├── FormField.types.ts
    │       │   │   │   │   ├── FormField.web.tsx
    │       │   │   │   │   ├── FormHelperText.tsx
    │       │   │   │   │   └── FormLabel.tsx
    │       │   │   │   ├── Radio/
    │       │   │   │   │   ├── Radio.types.ts
    │       │   │   │   │   ├── RadioGroup.native.tsx
    │       │   │   │   │   ├── RadioGroup.tsx
    │       │   │   │   │   └── RadioGroup.web.tsx
    │       │   │   │   ├── Select/
    │       │   │   │   │   ├── MultiSelect/
    │       │   │   │   │   │   ├── MultiSelect.native.tsx
    │       │   │   │   │   │   ├── MultiSelect.tsx
    │       │   │   │   │   │   └── MultiSelect.web.tsx
    │       │   │   │   │   ├── Select.native.tsx
    │       │   │   │   │   ├── Select.tsx
    │       │   │   │   │   ├── Select.types.ts
    │       │   │   │   │   └── Select.web.tsx
    │       │   │   │   ├── Slider/
    │       │   │   │   │   ├── RangeSlider/
    │       │   │   │   │   │   ├── RangeSlider.native.tsx
    │       │   │   │   │   │   ├── RangeSlider.tsx
    │       │   │   │   │   │   └── RangeSlider.web.tsx
    │       │   │   │   │   ├── Slider.native.tsx
    │       │   │   │   │   ├── Slider.tsx
    │       │   │   │   │   ├── Slider.types.ts
    │       │   │   │   │   └── Slider.web.tsx
    │       │   │   │   └── Switch/
    │       │   │   │       ├── Switch.native.tsx
    │       │   │   │       ├── Switch.tsx
    │       │   │   │       ├── Switch.types.ts
    │       │   │   │       └── Switch.web.tsx
    │       │   │   ├── Input/
    │       │   │   │   ├── Input.native.tsx
    │       │   │   │   ├── Input.test.tsx
    │       │   │   │   ├── Input.tsx
    │       │   │   │   └── Input.web.tsx
    │       │   │   ├── learning/
    │       │   │   │   ├── AchievementBadge.native.tsx
    │       │   │   │   ├── AchievementBadge.tsx  # Achievement badge
    │       │   │   │   ├── AchievementBadge.web.tsx
    │       │   │   │   ├── LearningPathCard.native.tsx
    │       │   │   │   ├── LearningPathCard.tsx  # Learning path card
    │       │   │   │   ├── LearningPathCard.web.tsx
    │       │   │   │   ├── ProgressTracker.native.tsx
    │       │   │   │   ├── ProgressTracker.tsx  # Progress visualization
    │       │   │   │   └── ProgressTracker.web.tsx
    │       │   │   ├── Modal/
    │       │   │   ├── Navigation/
    │       │   │   │   ├── Breadcrumb/
    │       │   │   │   │   ├── Breadcrumb.tsx
    │       │   │   │   │   ├── Breadcrumb.types.ts
    │       │   │   │   │   └── Breadcrumb.web.tsx
    │       │   │   │   ├── Pagination/
    │       │   │   │   │   ├── Pagination.native.tsx
    │       │   │   │   │   ├── Pagination.tsx
    │       │   │   │   │   ├── Pagination.types.ts
    │       │   │   │   │   └── Pagination.web.tsx
    │       │   │   │   ├── Stepper/
    │       │   │   │   │   ├── Stepper.native.tsx
    │       │   │   │   │   ├── Stepper.tsx
    │       │   │   │   │   ├── Stepper.types.ts
    │       │   │   │   │   └── Stepper.web.tsx
    │       │   │   │   └── Tabs/
    │       │   │   │       ├── Tabs.native.tsx
    │       │   │   │       ├── Tabs.tsx
    │       │   │   │       ├── Tabs.types.ts
    │       │   │   │       └── Tabs.web.tsx
    │       │   │   ├── Pagination/
    │       │   │   ├── Popover/
    │       │   │   ├── Progress/
    │       │   │   ├── Radio/
    │       │   │   ├── Rating/
    │       │   │   ├── Select/
    │       │   │   ├── Skeleton/
    │       │   │   ├── Slider/
    │       │   │   ├── Stepper/
    │       │   │   ├── Switch/
    │       │   │   ├── Tabs/
    │       │   │   ├── Textarea/
    │       │   │   ├── Timeline/
    │       │   │   ├── TimePicker/
    │       │   │   ├── Toast/
    │       │   │   ├── Tooltip/
    │       │   │   ├── tracking/
    │       │   │   │   ├── TimesheetTable.native.tsx
    │       │   │   │   ├── TimesheetTable.tsx  # Timesheet grid
    │       │   │   │   ├── TimesheetTable.web.tsx
    │       │   │   │   ├── TimeTracker.native.tsx
    │       │   │   │   ├── TimeTracker.tsx  # Time tracking widget
    │       │   │   │   ├── TimeTracker.web.tsx
    │       │   │   │   ├── WorkDiaryEntry.native.tsx
    │       │   │   │   ├── WorkDiaryEntry.tsx  # Work diary card
    │       │   │   │   └── WorkDiaryEntry.web.tsx
    │       │   │   └── video/
    │       │   │       ├── VideoPlayer.native.tsx
    │       │   │       ├── VideoPlayer.tsx  # Video player
    │       │   │       ├── VideoPlayer.web.tsx
    │       │   │       ├── VideoUploader.native.tsx
    │       │   │       ├── VideoUploader.tsx  # Video upload
    │       │   │       └── VideoUploader.web.tsx
    │       │   ├── forms/  # Form components
    │       │   │   ├── FormError/
    │       │   │   ├── FormField/
    │       │   │   ├── FormGroup/
    │       │   │   ├── FormHelper/
    │       │   │   └── FormLabel/
    │       │   ├── icons/  # Icon components
    │       │   │   └── index.ts  # Export all icons
    │       │   └── layout/  # Layout components
    │       │       ├── Container/
    │       │       ├── Divider/
    │       │       ├── Grid/
    │       │       ├── Spacer/
    │       │       └── Stack/
    │       ├── package.json
    │       ├── README.md
    │       └── tsconfig.json
    ├── packages/shared/src/features/
    ├── packages/ui/src/
    ├── public/  # Static assets
    │   ├── animations/  # Lottie/animation files
    │   │   ├── empty-state.json
    │   │   ├── error.json
    │   │   ├── loading.json
    │   │   └── success.json
    │   ├── fonts/  # Web fonts
    │   │   ├── inter-var-latin.woff2  # Latin subset
    │   │   ├── inter-var.woff2  # Variable font
    │   │   └── noto-sans-arabic.woff2  # Arabic font
    │   ├── images/
    │   │   ├── payment-providers/  # Payment method icons
    │   │   │   ├── mastercard.svg
    │   │   │   ├── paypal.svg
    │   │   │   ├── stripe.svg
    │   │   │   └── visa.svg
    │   │   ├── social/  # Social login icons
    │   │   │   ├── github.svg
    │   │   │   ├── google.svg
    │   │   │   └── linkedin.svg
    │   │   ├── favicon.ico
    │   │   ├── hero-bg-mobile.webp  # Mobile hero
    │   │   ├── hero-bg.webp  # Hero background
    │   │   ├── logo-dark.svg  # Dark mode logo
    │   │   ├── logo.svg  # Skillsier logo
    │   │   ├── placeholder-avatar.png  # Default avatar
    │   │   └── placeholder-company.png  # Default company logo
    │   ├── locales/  # Public locale files (deprecated - moved to packages/shared)
    │   ├── browserconfig.xml  # Windows tile config
    │   ├── manifest.json  # PWA manifest
    │   ├── robots.txt  # Search engine directives
    │   ├── sitemap-dynamic.xml  # Dynamic routes sitemap
    │   ├── sitemap.xml  # Static sitemap
    │   └── sw.js  # Service worker (if using PWA)
    ├── realtime/
    │   ├── hooks/
    │   │   ├── usePresence.ts  # User presence (online/offline)
    │   │   ├── useRealtimeAuction.ts  # Real-time auction updates
    │   │   ├── useRealtimeMessages.ts  # Real-time messages
    │   │   ├── useRealtimeNotifications.ts  # Real-time notifications
    │   │   └── useWebSocket.ts  # WebSocket connection
    │   ├── store/
    │   │   └── realtime-store.ts  # Real-time state (Zustand)
    │   ├── websocket/
    │   │   ├── client.ts  # WebSocket client
    │   │   ├── heartbeat.ts  # Connection health
    │   │   └── reconnection.ts  # Reconnection logic
    │   └── types.ts  # Real-time types
    ├── resources/
    │   ├── blog/
    │   │   ├── [postId]/
    │   │   │   └── page.tsx  # Blog post
    │   │   │       # BE: CMS
    │   │   │       # GET /v1/content/blog/{post_id}
    │   │   ├── category/
    │   │   │   └── [categoryId]/
    │   │   │       └── page.tsx  # Blog category
    │   │   │           # BE: CMS
    │   │   │           # GET /v1/content/blog?category={category_id}
    │   │   ├── page.tsx  # Blog home
    │   │   │   # BE: CMS
    │   │   │   # GET /v1/content/blog
    │   │   └── │
    │   ├── case-studies/
    │   │   ├── [caseStudyId]/
    │   │   │   └── page.tsx  # Case study detail
    │   │   │       # BE: CMS
    │   │   │       # GET /v1/content/case-studies/{case_study_id}
    │   │   ├── page.tsx  # Case studies list
    │   │   │   # BE: CMS
    │   │   │   # GET /v1/content/case-studies
    │   │   └── │
    │   ├── faq/
    │   │   └── page.tsx  # Frequently asked questions
    │   │       # BE: CMS
    │   │       # GET /v1/content/faq
    │   │       # 3. Enterprise & Business Solutions
    │   ├── guides/
    │   │   ├── [guideId]/
    │   │   │   └── page.tsx  # Guide detail
    │   │   │       # BE: CMS or static
    │   │   │       # GET /v1/content/guides/{guide_id}
    │   │   ├── client/
    │   │   │   └── page.tsx  # Client guides
    │   │   │       # BE: CMS
    │   │   │       # GET /v1/content/guides?category=client
    │   │   ├── freelancer/
    │   │   │   └── page.tsx  # Freelancer guides
    │   │   │       # BE: CMS
    │   │   │       # GET /v1/content/guides?category=freelancer
    │   │   ├── page.tsx  # All guides
    │   │   │   # BE: CMS
    │   │   │   # GET /v1/content/guides
    │   │   └── │
    │   ├── tutorials/
    │   │   ├── [tutorialId]/
    │   │   │   └── page.tsx  # Tutorial detail
    │   │   │       # BE: CMS
    │   │   │       # GET /v1/content/tutorials/{tutorial_id}
    │   │   ├── page.tsx  # Tutorials list
    │   │   │   # BE: CMS
    │   │   │   # GET /v1/content/tutorials
    │   │   └── │
    │   ├── webinars/
    │   │   ├── [webinarId]/
    │   │   │   └── page.tsx  # Webinar detail & registration
    │   │   │       # BE: CMS + registration system
    │   │   │       # GET /v1/content/webinars/{webinar_id}
    │   │   │       # POST /v1/webinars/{webinar_id}/register
    │   │   ├── page.tsx  # Upcoming webinars
    │   │   │   # BE: CMS
    │   │   │   # GET /v1/content/webinars
    │   │   └── │
    │   └── │
    ├── reviews/
    │   ├── analytics/
    │   │   └── page.tsx  # Review analytics
    │   │       # - Rating trends
    │   │       # - Response rate
    │   │       # - Sentiment analysis
    │   │       # BE: reviews-be/review
    │   │       # GET /v1/users/me/reviews/analytics
    │   │       # 9. Networking & Connections
    │   ├── disputes/
    │   │   ├── [disputeId]/
    │   │   │   └── page.tsx  # Review dispute details
    │   │   │       # - Evidence submission
    │   │   │       # - Admin review status
    │   │   │       # BE: reviews-be/review, admin-be/case_mgmt
    │   │   │       # GET /v1/reviews/{review_id}/dispute
    │   │   ├── page.tsx  # Review disputes list
    │   │   │   # BE: reviews-be/review
    │   │   │   # GET /v1/reviews/disputes
    │   │   └── │
    │   ├── given/
    │   │   ├── [reviewId]/
    │   │   │   ├── edit/
    │   │   │   │   └── page.tsx  # Edit given review
    │   │   │   │       # BE: reviews-be/review
    │   │   │   │       # PUT /v1/reviews/{review_id}
    │   │   │   ├── page.tsx  # Given review details
    │   │   │   │   # BE: reviews-be/review
    │   │   │   │   # GET /v1/reviews/{review_id}
    │   │   │   └── │
    │   │   ├── page.tsx  # Given reviews list
    │   │   │   # BE: reviews-be/review
    │   │   │   # GET /v1/users/me/reviews/given
    │   │   └── │
    │   ├── pending/
    │   │   ├── [contractId]/
    │   │   │   └── page.tsx  # Leave review form
    │   │   │       # BE: reviews-be/review, contracts-be/contract
    │   │   │       # GET /v1/contracts/{contract_id}
    │   │   │       # POST /v1/reviews
    │   │   ├── page.tsx  # Pending reviews to complete
    │   │   │   # BE: reviews-be/review
    │   │   │   # GET /v1/reviews/pending
    │   │   └── │
    │   ├── received/
    │   │   ├── [reviewId]/
    │   │   │   ├── respond/
    │   │   │   │   └── page.tsx  # Respond to review
    │   │   │   │       # BE: reviews-be/review
    │   │   │   │       # POST /v1/reviews/{review_id}/response
    │   │   │   ├── page.tsx  # Review details
    │   │   │   │   # BE: reviews-be/review
    │   │   │   │   # GET /v1/reviews/{review_id}
    │   │   │   └── │
    │   │   ├── page.tsx  # Received reviews list
    │   │   │   # BE: reviews-be/review
    │   │   │   # GET /v1/users/me/reviews/received
    │   │   └── │
    │   └── │
    ├── scanner/
    │   ├── document.tsx  # Document scanner
    │   │   # - Scan compliance docs
    │   │   # - OCR processing
    │   │   # BE: storage-be/asset, admin-be/business_verification
    │   └── qr-code.tsx  # QR code scanner
    │       # - Event check-in
    │       # - Profile sharing
    ├── scripts/  # Automation scripts
    │   ├── build/
    │   │   ├── analyze-bundle.sh  # Bundle analysis
    │   │   ├── build-all.sh  # Build all apps
    │   │   ├── build-mobile.sh  # Build mobile
    │   │   └── build-web.sh  # Build web
    │   ├── ci/
    │   │   ├── lint-all.sh  # Lint all code
    │   │   ├── pre-commit.sh  # Pre-commit hook
    │   │   ├── pre-push.sh  # Pre-push hook
    │   │   └── verify-types.sh  # Type check
    │   ├── deploy/
    │   │   ├── deploy-mobile-prod.sh
    │   │   ├── deploy-mobile-staging.sh
    │   │   ├── deploy-web-prod.sh
    │   │   └── deploy-web-staging.sh
    │   ├── dev/
    │   │   ├── reset-db.sh  # Reset local DB
    │   │   ├── start-all.sh  # Start all services
    │   │   ├── start-mobile.sh  # Start mobile only
    │   │   └── start-web.sh  # Start web only
    │   ├── maintenance/
    │   │   ├── clean-all.sh  # Clean all artifacts
    │   │   ├── generate-types.sh  # Generate API types
    │   │   ├── sync-translations.sh  # Sync i18n
    │   │   └── update-deps.sh  # Update dependencies
    │   ├── test/
    │   │   ├── test-all.sh  # Run all tests
    │   │   ├── test-coverage.sh  # Coverage report
    │   │   ├── test-e2e-mobile.sh  # E2E mobile tests
    │   │   ├── test-e2e-web.sh  # E2E web tests
    │   │   ├── test-integration.sh  # Integration tests
    │   │   └── test-unit.sh  # Unit tests
    │   ├── utils/
    │   │   ├── check-ports.sh  # Check port availability
    │   │   ├── doctor.sh  # Environment health check
    │   │   └── setup-env.sh  # Setup environment
    │   ├── build-all.sh  # Build all apps
    │   ├── check-bundle-size.sh  # Bundle size analysis
    │   ├── clean.sh  # Clean build artifacts
    │   ├── db-seed.sh  # Seed local dev data
    │   ├── dev.sh  # Start all dev servers
    │   ├── generate-types.sh  # Generate types from OpenAPI
    │   ├── setup.sh  # Initial project setup
    │   └── test-all.sh  # Run all tests
    ├── search/
    │   ├── advanced/
    │   │   └── page.tsx  # Advanced search interface
    │   │       # - Complex filters builder
    │   │       # - Boolean operators
    │   │       # - Saved search management
    │   │       # BE: search-be/query
    │   │       # POST /v1/search/advanced
    │   ├── api/
    │   │   ├── recommendations-api.ts  # Recommendations API
    │   │   │   # BE: search-be/recommendation
    │   │   ├── saved-searches-api.ts  # Saved searches API
    │   │   │   # BE: search-be/saved-search
    │   │   ├── search-api.ts  # Search API (already may exist, ensuring completeness)
    │   │   │   # BE: search-be/query
    │   │   └── trending-api.ts  # Trending API
    │   │       # BE: search-be/trending
    │   ├── history/
    │   │   └── page.tsx  # Search history
    │   │       # BE: search-be/query
    │   │       # GET /v1/search/history
    │   │       # 2. Connects & Credits Management (Freelancer)
    │   ├── hooks/
    │   │   ├── useRecommendations.ts  # Recommendations
    │   │   ├── useSavedSearches.ts  # Saved searches
    │   │   ├── useSearch.ts  # Search execution
    │   │   ├── useSearchHistory.ts  # Search history
    │   │   ├── useSearchSuggestions.ts  # Auto-complete suggestions
    │   │   └── useTrending.ts  # Trending items
    │   ├── queries/
    │   │   ├── search-mutations.ts  # Search mutations
    │   │   └── search-queries.ts  # Search queries
    │   ├── recommendations/
    │   │   └── page.tsx  # Personalized recommendations
    │   │       # - AI-powered job matches
    │   │       # - Talent suggestions
    │   │       # BE: search-be/recommendation
    │   │       # GET /v1/recommendations/personalized
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
    │   ├── store/
    │   │   └── search-store.ts  # Search UI state (filters, etc.)
    │   ├── trending/
    │   │   └── page.tsx  # Trending searches and jobs
    │   │       # BE: search-be/trending
    │   │       # GET /v1/trending/jobs
    │   │       # GET /v1/trending/skills
    │   ├── types.ts  # Search types
    │   │   # 11. Bidding Module
    │   └── │
    ├── security/
    │   ├── bug-bounty/
    │   │   └── page.tsx  # Bug bounty program
    │   │       # BE: none (static content)
    │   ├── certifications/
    │   │   └── page.tsx  # Security certifications (SOC2, ISO, etc.)
    │   │       # BE: none (static content)
    │   ├── overview/
    │   │   └── page.tsx  # Security overview
    │   │       # BE: none (static content)
    │   ├── responsible-disclosure/
    │   │   └── page.tsx  # Responsible disclosure policy
    │   │       # BE: none (static content)
    │   ├── auth-guard.ts
    │   ├── csrf-protection.ts
    │   ├── rate-limiter.ts
    │   └── │
    ├── shortlists/
    │   ├── api/
    │   │   └── shortlists-api.ts  # Shortlists API
    │   │       # BE: jobs-be/shortlist
    │   ├── hooks/
    │   │   ├── useAddToShortlist.ts  # Add candidate
    │   │   ├── useRemoveFromShortlist.ts  # Remove candidate
    │   │   ├── useShortlist.ts  # Single shortlist
    │   │   └── useShortlists.ts  # Shortlists management
    │   ├── queries/
    │   │   ├── shortlists-mutations.ts  # Shortlist mutations
    │   │   └── shortlists-queries.ts  # Shortlist queries
    │   └── types.ts  # Shortlist types
    │       # 14. Feature Flags Module
    ├── status/
    │   ├── current/
    │   │   └── page.tsx  # Current system status
    │   │       # BE: utility/status
    │   │       # GET /v1/status/current
    │   ├── history/
    │   │   └── page.tsx  # Status history
    │   │       # BE: utility/status
    │   │       # GET /v1/status/history
    │   ├── subscribe/
    │   │   └── page.tsx  # Subscribe to status updates
    │   │       # BE: communications-be
    │   │       # POST /v1/notifications/status-subscribe
    │   └── │
    ├── talent/
    │   ├── browse/
    │   │   └── page.tsx  # Browse talent
    │   │       # - Search freelancers
    │   │       # - Filters (skills, rate, location)
    │   │       # - Save to shortlist
    │   │       # BE: search-be/query, users-be/profile
    │   │       # POST /v1/search/freelancers
    │   │       # GET /v1/search/freelancers?filters=...
    │   ├── recommendations/
    │   │   └── page.tsx  # AI-recommended talent for jobs
    │   │       # BE: search-be/recommendation
    │   │       # GET /v1/recommendations/talent?job_id={job_id}
    │   ├── saved/
    │   │   └── page.tsx  # Saved talent profiles
    │   │       # BE: users-be/profile
    │   │       # GET /v1/users/me/saved-profiles
    │   │       # 6. Work Tracking & Time Management
    │   ├── shortlists/
    │   │   ├── [shortlistId]/
    │   │   │   ├── edit/
    │   │   │   │   └── page.tsx  # Edit shortlist
    │   │   │   │       # BE: jobs-be/shortlist
    │   │   │   │       # PUT /v1/jobs/{job_id}/shortlists/{shortlist_id}
    │   │   │   ├── page.tsx  # Shortlist details
    │   │   │   │   # - View candidates
    │   │   │   │   # - Send invitations
    │   │   │   │   # - Compare profiles
    │   │   │   │   # BE: jobs-be/shortlist
    │   │   │   │   # GET /v1/jobs/{job_id}/shortlists/{shortlist_id}
    │   │   │   └── │
    │   │   ├── new/
    │   │   │   └── page.tsx  # Create shortlist
    │   │   │       # BE: jobs-be/shortlist
    │   │   │       # POST /v1/jobs/{job_id}/shortlists
    │   │   ├── page.tsx  # Shortlists overview
    │   │   │   # BE: jobs-be/shortlist
    │   │   │   # GET /v1/jobs/{job_id}/shortlists
    │   │   └── │
    │   └── │
    ├── tests/
    │   ├── e2e/
    │   │   ├── fixtures/
    │   │   │   ├── ...
    │   │   │   ├── jobs.json
    │   │   │   └── users.json
    │   │   ├── mobile/
    │   │   │   ├── auth/
    │   │   │   ├── client/
    │   │   │   ├── freelancer/
    │   │   │   └── offline/
    │   │   └── web/
    │   │       ├── admin/
    │   │       │   ├── financial-ops.spec.ts
    │   │       │   ├── kyc.spec.ts
    │   │       │   └── moderation.spec.ts
    │   │       ├── auth/
    │   │       │   ├── login.spec.ts
    │   │       │   ├── password-reset.spec.ts
    │   │       │   └── register.spec.ts
    │   │       ├── client/
    │   │       │   ├── hiring.spec.ts
    │   │       │   ├── job-posting.spec.ts
    │   │       │   ├── payments.spec.ts
    │   │       │   └── talent-search.spec.ts
    │   │       └── freelancer/
    │   │           ├── contracts.spec.ts
    │   │           ├── earnings.spec.ts
    │   │           ├── profile.spec.ts
    │   │           └── proposals.spec.ts
    │   ├── integration/
    │   │   ├── api/
    │   │   │   ├── ...
    │   │   │   ├── jobs-api.test.ts
    │   │   │   └── users-api.test.ts
    │   │   └── features/
    │   │       ├── ...
    │   │       ├── auth.test.ts
    │   │       ├── jobs.test.ts
    │   │       └── proposals.test.ts
    │   └── performance/
    │       ├── lighthouse/
    │       │   ├── dashboard.spec.ts
    │       │   ├── home.spec.ts
    │       │   └── job-detail.spec.ts
    │       └── load/
    │           ├── api-load.spec.ts
    │           └── user-simulation.spec.ts
    ├── timesheets/
    │   ├── [contractId]/
    │   │   ├── [timesheetId]/
    │   │   │   ├── edit/
    │   │   │   │   └── page.tsx  # Edit timesheet
    │   │   │   │       # BE: contracts-be/timesheet
    │   │   │   │       # PUT /v1/contracts/{contract_id}/timesheets/{timesheet_id}
    │   │   │   ├── page.tsx  # Timesheet details
    │   │   │   │   # - Hours breakdown
    │   │   │   │   # - Approval status
    │   │   │   │   # - Dispute options
    │   │   │   │   # BE: contracts-be/timesheet
    │   │   │   │   # GET /v1/contracts/{contract_id}/timesheets/{timesheet_id}
    │   │   │   └── │
    │   │   ├── new/
    │   │   │   └── page.tsx  # Create timesheet
    │   │   │       # BE: contracts-be/timesheet
    │   │   │       # POST /v1/contracts/{contract_id}/timesheets
    │   │   ├── page.tsx  # Contract timesheets list
    │   │   │   # BE: contracts-be/timesheet
    │   │   │   # GET /v1/contracts/{contract_id}/timesheets
    │   │   └── │
    │   ├── approve/
    │   │   └── page.tsx  # Timesheets pending approval (client)
    │   │       # BE: contracts-be/timesheet
    │   │       # GET /v1/timesheets/pending-approval
    │   ├── page.tsx  # All timesheets overview
    │   │   # BE: contracts-be/timesheet
    │   │   # GET /v1/timesheets
    │   │   # 7. Deliverables Management
    │   └── │
    ├── tracking/
    │   ├── TimesheetTable.native.tsx
    │   ├── TimesheetTable.tsx  # Timesheet grid
    │   ├── TimesheetTable.web.tsx
    │   ├── TimeTracker.native.tsx
    │   ├── TimeTracker.tsx  # Time tracking widget
    │   ├── TimeTracker.web.tsx
    │   ├── WorkDiaryEntry.native.tsx
    │   ├── WorkDiaryEntry.tsx  # Work diary card
    │   └── WorkDiaryEntry.web.tsx
    ├── transparency/
    │   └── page.tsx  # Transparency report
    │       # - User statistics
    │       # - Moderation actions
    │       # - Government requests
    │       # BE: admin-be/reporting
    │       # GET /v1/public/transparency-report
    ├── video/
    │   ├── VideoPlayer.native.tsx
    │   ├── VideoPlayer.tsx  # Video player
    │   ├── VideoPlayer.web.tsx
    │   ├── VideoUploader.native.tsx
    │   ├── VideoUploader.tsx  # Video upload
    │   └── VideoUploader.web.tsx
    ├── widgets/
    │   ├── quick-actions.tsx  # Quick actions widget
    │   │   # - Quick message
    │   │   # - Quick proposal
    │   └── time-tracker.tsx  # Home screen time tracker widget
    │       # BE: contracts-be/work_diary
    ├── work-diary/
    │   ├── [contractId]/
    │   │   ├── calendar/
    │   │   │   └── page.tsx  # Calendar view of work diary
    │   │   │       # BE: contracts-be/work_diary
    │   │   │       # GET /v1/contracts/{contract_id}/work-diary/calendar
    │   │   ├── screenshots/
    │   │   │   └── page.tsx  # Screenshots management
    │   │   │       # - View all screenshots
    │   │   │       # - Delete sensitive ones
    │   │   │       # - Privacy settings
    │   │   │       # BE: contracts-be/work_diary, storage-be/asset
    │   │   │       # GET /v1/contracts/{contract_id}/work-diary/screenshots
    │   │   ├── page.tsx  # Work diary detail
    │   │   │   # BE: contracts-be/work_diary
    │   │   │   # GET /v1/contracts/{contract_id}/work-diary
    │   │   └── │
    │   ├── page.tsx  # Work diary overview (all contracts)
    │   │   # BE: contracts-be/work_diary
    │   │   # GET /v1/work-diary
    │   └── │
    ├── work-tracking/
    │   ├── api/
    │   │   ├── timesheet-api.ts  # Timesheet API
    │   │   │   # BE: contracts-be/timesheet
    │   │   └── work-diary-api.ts  # Work diary API
    │   │       # BE: contracts-be/work_diary
    │   ├── hooks/
    │   │   ├── useApproveTimesheet.ts  # Approve timesheet (client)
    │   │   ├── useTimesheet.ts  # Timesheet management
    │   │   ├── useTimeTracking.ts  # Real-time time tracking
    │   │   └── useWorkDiary.ts  # Work diary entries
    │   ├── queries/
    │   │   ├── timesheet-mutations.ts  # Timesheet mutations
    │   │   ├── timesheet-queries.ts  # Timesheet queries
    │   │   ├── work-diary-mutations.ts  # Work diary mutations
    │   │   └── work-diary-queries.ts  # Work diary queries
    │   ├── store/
    │   │   └── time-tracker-store.ts  # Time tracker state (Zustand)
    │   └── types.ts  # Work tracking types
    │       # 6. Deliverables Module
    ├── ---  # Additional Public Routes
    │   # 1. Legal & Compliance Pages
    │   # Additional Shared Features (packages/shared/src/features/)
    │   # 1. Auction System Module
    │   # Additional Mobile App Routes (apps/mobile/app/)
    │   # 1. Enhanced Mobile Navigation
    │   # Additional UI Components (packages/ui/src/)
    │   # 1. Advanced Components
    ├── .env.development  # Development environment
    │   # Local development
    ├── .env.example  # Environment variables template
    │   # Template
    ├── .env.local  # Local environment variables
    ├── .env.local.example  # Local overrides template
    ├── .env.production  # Production environment
    ├── .env.staging  # Staging environment
    ├── .eslintrc.json  # Root ESLint config
    │   # Web-specific ESLint rules
    ├── .gitignore  # Git ignore patterns
    ├── .nvmrc  # Node.js version (v20.x)
    ├── .prettierignore  # Prettier ignore patterns
    ├── .prettierrc  # Prettier configuration
    ├── error.tsx  # Global error boundary
    ├── globals.css  # Global styles
    │   # - Tailwind base, components, utilities
    │   # - CSS custom properties for theming
    │   # - Dark mode variables
    ├── jest.config.js  # Root Jest configuration
    │   # Web Jest configuration
    ├── layout.tsx  # Root layout
    │   # - HTML lang attribute (i18n)
    │   # - Head meta tags
    │   # - Body layout
    │   # - Font loading
    ├── loading.tsx  # Global loading state
    ├── middleware.ts  # Next.js middleware
    │   # - Security headers
    │   # - CSRF protection
    │   # - Rate limiting (client-side)
    │   # - Auth checks
    ├── next.config.js  # Next.js configuration
    │   # - i18n config
    │   # - Image optimization
    │   # - Bundle analyzer
    │   # - Security headers
    │   # - Rewrites/redirects
    ├── not-found.tsx  # 404 page
    ├── package.json  # Root package (workspace manager)
    │   # Scripts: dev, build, test, lint, type-check
    │   # Web dependencies
    ├── page.tsx  # Root page (redirects to /[locale])
    ├── pnpm-lock.yaml  # Locked dependencies
    ├── pnpm-workspace.yaml  # pnpm workspace configuration
    ├── postcss.config.js  # PostCSS configuration
    ├── README.md  # Root README
    │   # Web app documentation
    ├── tailwind.config.js  # Tailwind configuration
    │   # - Design tokens
    │   # - Custom utilities
    │   # - Plugin configuration
    ├── tsconfig.base.json  # Base TypeScript configuration
    │   # Shared compiler options
    ├── tsconfig.json  # Root TypeScript config
    │   # Web TypeScript config
    │   # Extends tsconfig.base.json
    ├── turbo.json  # Turborepo pipeline configuration
    │   # Build cache, task dependencies
    └── │
```