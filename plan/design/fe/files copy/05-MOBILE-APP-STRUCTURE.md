# Skillsier Frontend - Complete Folder Structure
## Part 5: Mobile App (React Native/Expo) Structure

> **CRITICAL**: This document contains ONLY the folder structure, filenames, and inline backend API mappings.  
> **NO CODE IMPLEMENTATIONS** are included per the strict output policy.

---

## Mobile App Structure

```
apps/mobile/
│
├── app/                                     # Expo Router file-based routing
│   │
│   ├── (tabs)/                             # Bottom tabs navigation
│   │   ├── _layout.tsx                     # Tabs layout
│   │   │                                    # - Bottom tab navigator
│   │   │                                    # - Tab bar icons
│   │   │                                    # - Badge indicators (messages, notifications)
│   │   │
│   │   ├── index.tsx                       # Home tab / Dashboard
│   │   │                                    # - Dashboard overview
│   │   │                                    # - Quick actions
│   │   │                                    # - Recent activity
│   │   │                                    # BE: Multiple services (same as web dashboard)
│   │   │
│   │   ├── jobs.tsx                        # Jobs tab
│   │   │                                    # - Job listings (browse/my-jobs based on role)
│   │   │                                    # - Search jobs
│   │   │                                    # - Saved jobs
│   │   │                                    # BE: jobs-be/job
│   │   │                                    # GET /v1/jobs/browse (freelancer)
│   │   │                                    # GET /v1/jobs/my-jobs (client)
│   │   │
│   │   ├── proposals.tsx                   # Proposals tab
│   │   │                                    # - My proposals (freelancer)
│   │   │                                    # - Received proposals (client, redirects to job)
│   │   │                                    # BE: proposals-be
│   │   │                                    # GET /v1/proposals/my-proposals
│   │   │
│   │   ├── messages.tsx                    # Messages tab
│   │   │                                    # - Conversation list
│   │   │                                    # - Real-time updates
│   │   │                                    # BE: communications-be/conversations
│   │   │                                    # GET /v1/conversations
│   │   │                                    # WebSocket connection
│   │   │
│   │   └── profile.tsx                     # Profile tab
│   │                                        # - Current user profile preview
│   │                                        # - Quick links to settings
│   │                                        # - Stats
│   │                                        # BE: users-be/profile
│   │                                        # GET /v1/users/me
│   │
│   ├── (auth)/                             # Auth screens
│   │   ├── _layout.tsx                     # Auth layout
│   │   ├── login.tsx                       # Login screen
│   │   │                                    # - Email/password form
│   │   │                                    # - Social login (Google, Apple)
│   │   │                                    # - Biometric login (Face ID, Touch ID)
│   │   │                                    # BE: Keycloak OAuth2
│   │   │                                    # POST /v1/auth/login
│   │   │
│   │   ├── register.tsx                    # Registration screen
│   │   │                                    # BE: users-be/user
│   │   │                                    # POST /v1/users/register
│   │   │
│   │   ├── forgot-password.tsx             # Password reset
│   │   │                                    # BE: users-be/security/recovery
│   │   │                                    # POST /v1/auth/forgot-password
│   │   │
│   │   └── callback.tsx                    # OAuth callback
│   │                                        # BE: Keycloak token exchange
│   │
│   ├── (onboarding)/                       # Onboarding flow
│   │   ├── _layout.tsx                     # Onboarding layout
│   │   ├── welcome.tsx                     # Welcome screen
│   │   ├── profile.tsx                     # Basic profile setup
│   │   ├── skills.tsx                      # Skills selection
│   │   ├── preferences.tsx                 # Preferences
│   │   └── complete.tsx                    # Onboarding complete
│   │
│   ├── jobs/
│   │   ├── [id].tsx                        # Job detail
│   │   │                                    # BE: jobs-be/job
│   │   │                                    # GET /v1/jobs/{job_id}
│   │   │
│   │   ├── search.tsx                      # Job search
│   │   │                                    # BE: search-be
│   │   │                                    # POST /v1/search/jobs
│   │   │
│   │   └── post.tsx                        # Post job (client)
│   │                                        # BE: jobs-be/job
│   │                                        # POST /v1/jobs
│   │
│   ├── proposals/
│   │   ├── [id].tsx                        # Proposal detail
│   │   │                                    # BE: proposals-be
│   │   │                                    # GET /v1/proposals/{proposal_id}
│   │   │
│   │   └── submit/
│   │       └── [jobId].tsx                 # Submit proposal
│   │                                        # BE: proposals-be
│   │                                        # POST /v1/proposals
│   │
│   ├── contracts/
│   │   ├── [id]/
│   │   │   ├── index.tsx                   # Contract overview
│   │   │   │                                # BE: contracts-be/contract
│   │   │   │                                # GET /v1/contracts/{contract_id}
│   │   │   │
│   │   │   ├── milestones.tsx              # Milestones
│   │   │   │                                # BE: contracts-be/milestone
│   │   │   │                                # GET /v1/contracts/{contract_id}/milestones
│   │   │   │
│   │   │   ├── timesheet.tsx               # Timesheet
│   │   │   │                                # BE: contracts-be/timesheet
│   │   │   │                                # GET /v1/contracts/{contract_id}/timesheets
│   │   │   │
│   │   │   ├── work-diary.tsx              # Work diary
│   │   │   │                                # BE: contracts-be/work_diary
│   │   │   │                                # GET /v1/contracts/{contract_id}/work-diary
│   │   │   │
│   │   │   └── messages.tsx                # Contract messages
│   │   │                                    # BE: communications-be
│   │   │                                    # GET /v1/contracts/{contract_id}/conversation
│   │   │
│   │   └── index.tsx                       # Contracts list
│   │                                        # BE: contracts-be/contract
│   │                                        # GET /v1/contracts
│   │
│   ├── messages/
│   │   └── [conversationId].tsx            # Conversation thread
│   │                                        # - Message list
│   │                                        # - Message composer
│   │                                        # - Real-time updates
│   │                                        # BE: communications-be/messages
│   │                                        # GET /v1/conversations/{conversation_id}/messages
│   │                                        # POST /v1/messages
│   │                                        # WebSocket updates
│   │
│   ├── notifications/
│   │   └── index.tsx                       # Notifications list
│   │                                        # BE: communications-be/notifications
│   │                                        # GET /v1/notifications
│   │
│   ├── profile/
│   │   ├── [userId].tsx                    # User profile (public view)
│   │   │                                    # BE: users-be/profile
│   │   │                                    # GET /v1/users/{user_id}/profile
│   │   │
│   │   └── edit/
│   │       ├── index.tsx                   # Edit profile
│   │       ├── skills.tsx                  # Edit skills
│   │       ├── experience.tsx              # Edit experience
│   │       └── portfolio.tsx               # Edit portfolio
│   │
│   ├── settings/
│   │   ├── index.tsx                       # Settings menu
│   │   ├── account.tsx                     # Account settings
│   │   ├── security.tsx                    # Security settings
│   │   ├── notifications.tsx               # Notification settings
│   │   ├── privacy.tsx                     # Privacy settings
│   │   └── about.tsx                       # About & support
│   │
│   ├── financials/
│   │   ├── wallet.tsx                      # Wallet
│   │   │                                    # BE: financial-be/wallet
│   │   │                                    # GET /v1/wallet/balance
│   │   │
│   │   ├── transactions.tsx                # Transaction history
│   │   │                                    # BE: financial-be/transaction
│   │   │                                    # GET /v1/transactions
│   │   │
│   │   └── invoices.tsx                    # Invoices
│   │                                        # BE: financial-be/invoice
│   │                                        # GET /v1/invoices
│   │
│   ├── reviews/
│   │   ├── index.tsx                       # Reviews list
│   │   │                                    # BE: reviews-be/reviews
│   │   │                                    # GET /v1/reviews
│   │   │
│   │   └── create/
│   │       └── [contractId].tsx            # Create review
│   │                                        # BE: reviews-be/reviews
│   │                                        # POST /v1/reviews
│   │
│   ├── subscription/
│   │   ├── index.tsx                       # Subscription overview
│   │   ├── plans.tsx                       # Available plans
│   │   ├── upgrade.tsx                     # Upgrade plan
│   │   └── connects.tsx                    # Connects management
│   │                                        # BE: subscriptions-be
│   │
│   ├── +not-found.tsx                      # 404 screen
│   ├── _layout.tsx                         # Root layout
│   │                                        # - Auth provider
│   │                                        # - Query client provider
│   │                                        # - Theme provider
│   │                                        # - Error boundary
│   │
│   └── index.tsx                           # App entry point
│                                            # - Splash screen
│                                            # - Initial route determination
│
├── src/
│   ├── components/                         # Mobile-specific components
│   │   ├── Auth/
│   │   │   ├── LoginForm.tsx              # Login form component
│   │   │   ├── RegisterForm.tsx           # Registration form
│   │   │   ├── SocialButtons.tsx          # Social login buttons
│   │   │   └── BiometricButton.tsx        # Biometric auth button
│   │   │
│   │   ├── Common/
│   │   │   ├── ErrorBoundary.tsx          # Error boundary
│   │   │   ├── Loading.tsx                # Loading spinner
│   │   │   ├── EmptyState.tsx             # Empty state component
│   │   │   ├── OptimizedFlashList.tsx     # Optimized list (FlashList)
│   │   │   └── PullToRefresh.tsx          # Pull to refresh
│   │   │
│   │   ├── Jobs/
│   │   │   ├── JobCard.tsx                # Job card component
│   │   │   ├── JobList.tsx                # Job list
│   │   │   ├── JobDetail.tsx              # Job detail view
│   │   │   └── JobFilters.tsx             # Job filters bottom sheet
│   │   │
│   │   ├── Proposals/
│   │   │   ├── ProposalCard.tsx           # Proposal card
│   │   │   ├── ProposalList.tsx           # Proposal list
│   │   │   └── ProposalForm.tsx           # Proposal submission form
│   │   │
│   │   ├── Contracts/
│   │   │   ├── ContractCard.tsx           # Contract card
│   │   │   ├── MilestoneItem.tsx          # Milestone list item
│   │   │   └── TimesheetEntry.tsx         # Timesheet entry
│   │   │
│   │   ├── Messages/
│   │   │   ├── ConversationCard.tsx       # Conversation list item
│   │   │   ├── MessageBubble.tsx          # Message bubble
│   │   │   ├── MessageComposer.tsx        # Message input
│   │   │   └── TypingIndicator.tsx        # Typing indicator
│   │   │
│   │   ├── Profile/
│   │   │   ├── ProfileHeader.tsx          # Profile header
│   │   │   ├── SkillTag.tsx               # Skill tag
│   │   │   ├── ExperienceItem.tsx         # Experience item
│   │   │   └── PortfolioItem.tsx          # Portfolio item
│   │   │
│   │   ├── Financial/
│   │   │   ├── WalletCard.tsx             # Wallet balance card
│   │   │   ├── TransactionItem.tsx        # Transaction list item
│   │   │   └── InvoiceCard.tsx            # Invoice card
│   │   │
│   │   ├── Navigation/
│   │   │   ├── TabBar.tsx                 # Custom tab bar
│   │   │   └── Header.tsx                 # Screen header
│   │   │
│   │   └── UI/
│   │       ├── Button.tsx                 # Button component
│   │       ├── Input.tsx                  # Input component
│   │       ├── Card.tsx                   # Card component
│   │       ├── Badge.tsx                  # Badge component
│   │       ├── Avatar.tsx                 # Avatar component
│   │       ├── BottomSheet.tsx            # Bottom sheet modal
│   │       └── SearchBar.tsx              # Search bar
│   │
│   ├── hooks/                              # Mobile-specific hooks
│   │   ├── useHighFPSAnimation.ts         # High FPS animations
│   │   ├── useBiometricAuth.ts            # Biometric authentication
│   │   ├── usePushNotifications.ts        # Push notifications
│   │   ├── useKeyboard.ts                 # Keyboard handling
│   │   ├── useOrientation.ts              # Device orientation
│   │   ├── useAppState.ts                 # App state (foreground/background)
│   │   ├── useNetworkStatus.ts            # Network connectivity
│   │   ├── useLocation.ts                 # Geolocation
│   │   └── useCamera.ts                   # Camera access
│   │
│   ├── lib/
│   │   ├── keycloak-mobile.ts             # Keycloak mobile config
│   │   │                                    # - OAuth2 with PKCE
│   │   │                                    # - Token storage (SecureStore)
│   │   │                                    # - Refresh token flow
│   │   │
│   │   ├── performance.ts                 # Performance utilities
│   │   │                                    # - FPS monitoring
│   │   │                                    # - Memory optimization
│   │   │
│   │   ├── storage.ts                     # Secure storage wrapper
│   │   │                                    # - Token storage
│   │   │                                    # - Biometric keys
│   │   │
│   │   ├── push-notifications.ts          # Push notification setup
│   │   │                                    # - Firebase/Expo notifications
│   │   │                                    # - Token registration
│   │   │                                    # BE: communications-be/push
│   │   │                                    # POST /v1/push-tokens
│   │   │
│   │   ├── analytics.ts                   # Mobile analytics
│   │   │                                    # - Event tracking
│   │   │                                    # - Screen tracking
│   │   │
│   │   ├── error-tracking.ts              # Error tracking (Sentry)
│   │   ├── deeplink.ts                    # Deep link handling
│   │   └── utils.ts                       # General utilities
│   │
│   ├── lib/i18n/
│   │   └── index.ts                       # i18n configuration (mobile)
│   │                                        # Uses packages/shared i18n resources
│   │
│   ├── stores/                             # Mobile-specific Zustand stores
│   │   ├── offline-queue-store.ts         # Offline action queue
│   │   ├── camera-store.ts                # Camera state
│   │   └── biometric-store.ts             # Biometric settings
│   │
│   └── types/                              # Mobile-specific types
│       ├── navigation.ts                  # Navigation types
│       └── biometric.ts                   # Biometric types
│
├── assets/                                 # Mobile assets
│   ├── fonts/                             # Custom fonts
│   ├── images/                            # Images
│   ├── icons/                             # App icons
│   └── splash/                            # Splash screens
│
├── .env                                    # Environment variables
├── .eslintrc.json                         # ESLint config
├── app.json                               # Expo config
│                                            # - App name, slug
│                                            # - Icons, splash screens
│                                            # - Permissions
│                                            # - Build settings (EAS)
│
├── babel.config.js                        # Babel config
├── eas.json                               # EAS Build config
│                                            # - Development build
│                                            # - Preview build
│                                            # - Production build
│                                            # - Credentials
│
├── global.css                             # Global styles (NativeWind)
├── index.js                               # Entry point
├── metro.config.js                        # Metro bundler config
│                                            # - Monorepo support
│                                            # - Asset resolution
│
├── package.json                           # Mobile dependencies
│                                            # - Expo SDK
│                                            # - React Native
│                                            # - Expo Router
│                                            # - NativeWind
│                                            # - TanStack Query
│                                            # - Zustand
│
├── tailwind.config.js                     # Tailwind config (NativeWind)
└── tsconfig.json                          # TypeScript config
```

---

## Mobile App Features

### Key Mobile-Specific Features

1. **Biometric Authentication**
   - Face ID (iOS)
   - Touch ID (iOS)
   - Fingerprint (Android)
   - Secure token storage (Keychain/KeyStore)

2. **Push Notifications**
   - Real-time job alerts
   - Message notifications
   - Proposal status updates
   - Contract updates
   - Payment notifications
   - BE: communications-be/push
   - POST /v1/push-tokens (register device token)

3. **Offline Support**
   - TanStack Query persistence
   - Offline action queue
   - Sync when online
   - Optimistic updates

4. **Camera Integration**
   - Profile photo capture
   - Document scanning (KYC)
   - Work diary screenshots
   - Portfolio uploads

5. **Location Services**
   - Location-based job search
   - Time zone detection
   - Freelancer location

6. **Deep Linking**
   - Job deep links: skillsier://jobs/{job_id}
   - Contract deep links: skillsier://contracts/{contract_id}
   - Message deep links: skillsier://messages/{conversation_id}
   - Universal links support

7. **Performance Optimizations**
   - FlashList for large lists (60fps scrolling)
   - Image caching (expo-image)
   - Code splitting
   - Lazy loading
   - Memory optimization

8. **Accessibility**
   - Screen reader support
   - Voice control
   - High contrast mode
   - Font scaling

---

## Mobile State Management

### TanStack Query (Same as Web)
```typescript
// Same query keys as web
// Offline persistence enabled
// Background refetch on app foreground
```

### Zustand Stores (Mobile-Specific)
```typescript
// offline-queue-store
// - Queue mutations when offline
// - Sync when online

// camera-store
// - Camera permissions
// - Photo/video capture state

// biometric-store
// - Biometric availability
// - User preference for biometric auth
```

---

## Mobile Navigation Patterns

### Bottom Tabs
- Home/Dashboard
- Jobs (Browse/My Jobs)
- Proposals (My Proposals)
- Messages
- Profile/More

### Stack Navigation
- Job Detail → Submit Proposal → Proposal Submitted
- Contract Detail → Milestone Detail → Submit Deliverable
- Message List → Conversation → User Profile

### Modal Screens
- Filters (Bottom Sheet)
- Share
- Photo/Camera
- File picker

---

## Mobile Push Notification Types

All notification types from web, delivered via:
- **iOS**: Apple Push Notification Service (APNS)
- **Android**: Firebase Cloud Messaging (FCM)
- **Expo**: Expo Notifications API

BE Integration:
```
Communications-be publishes event → 
  Notification service → 
    Push notification provider (FCM/APNS) → 
      Mobile device
```

---

## Mobile Permissions Required

### iOS (Info.plist)
- Camera (NSCameraUsageDescription)
- Photo Library (NSPhotoLibraryUsageDescription)
- Location (NSLocationWhenInUseUsageDescription)
- Face ID (NSFaceIDUsageDescription)
- Notifications (User Notifications)

### Android (AndroidManifest.xml)
- CAMERA
- READ_EXTERNAL_STORAGE
- WRITE_EXTERNAL_STORAGE
- ACCESS_FINE_LOCATION
- USE_BIOMETRIC
- POST_NOTIFICATIONS (Android 13+)
- INTERNET
- ACCESS_NETWORK_STATE

---

## Mobile Build Configurations

### Development Build
- Development API endpoints
- Debug logging enabled
- Hot reload
- Dev tools

### Preview Build
- Staging API endpoints
- Limited logging
- No dev tools
- TestFlight/Internal testing

### Production Build
- Production API endpoints
- Error tracking only
- Optimized bundles
- App Store/Play Store release

---

**End of Part 5**

**Continue to Part 6 for:**
- Shared Packages (UI, Shared, Types, Config)
- Feature Modules Structure
- Query/Mutation Helpers
- WebSocket Integration
