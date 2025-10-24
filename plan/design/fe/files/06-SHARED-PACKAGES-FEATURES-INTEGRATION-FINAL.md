fe/
└── packages/
    ├── ui/                                      # Cross-platform UI component library
    │   ├── src/
    │   │   ├── components/
    │   │   │   ├── Button/
    │   │   │   │   ├── Button.tsx             # Base button component
    │   │   │   │   ├── Button.web.tsx         # Web-specific overrides
    │   │   │   │   ├── Button.native.tsx      # Native-specific overrides
    │   │   │   │   ├── Button.stories.tsx     # Storybook stories
    │   │   │   │   └── Button.test.tsx        # Component tests
    │   │   │   ├── Input/
    │   │   │   │   ├── Input.tsx
    │   │   │   │   ├── Input.web.tsx
    │   │   │   │   ├── Input.native.tsx
    │   │   │   │   └── Input.test.tsx
    │   │   │   ├── Card/
    │   │   │   ├── Badge/
    │   │   │   ├── Avatar/
    │   │   │   ├── Dropdown/
    │   │   │   ├── Modal/
    │   │   │   ├── Toast/
    │   │   │   ├── Tabs/
    │   │   │   ├── Accordion/
    │   │   │   ├── Pagination/
    │   │   │   ├── Rating/
    │   │   │   ├── Slider/
    │   │   │   ├── Switch/
    │   │   │   ├── Checkbox/
    │   │   │   ├── Radio/
    │   │   │   ├── Select/
    │   │   │   ├── Textarea/
    │   │   │   ├── DatePicker/
    │   │   │   ├── TimePicker/
    │   │   │   ├── FileUpload/
    │   │   │   ├── Progress/
    │   │   │   ├── Skeleton/
    │   │   │   ├── Tooltip/
    │   │   │   ├── Popover/
    │   │   │   ├── Alert/
    │   │   │   ├── Breadcrumb/
    │   │   │   ├── Stepper/
    │   │   │   ├── Timeline/
    │   │   │   └── DataTable/
    │   │   ├── forms/                          # Form components
    │   │   │   ├── FormField/
    │   │   │   ├── FormGroup/
    │   │   │   ├── FormLabel/
    │   │   │   ├── FormError/
    │   │   │   └── FormHelper/
    │   │   ├── layout/                         # Layout components
    │   │   │   ├── Container/
    │   │   │   ├── Grid/
    │   │   │   ├── Stack/
    │   │   │   ├── Divider/
    │   │   │   └── Spacer/
    │   │   └── icons/                          # Icon components
    │   │       └── index.ts                    # Export all icons
    │   ├── package.json
    │   ├── tsconfig.json
    │   └── README.md
    │
    ├── shared/                                  # Shared business logic, hooks, utilities
    │   ├── src/
    │   │   ├── features/                       # Feature-specific shared code
    │   │   │   ├── auth/                       # Authentication
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useAuth.ts          # Main auth hook
    │   │   │   │   │   │                       # - isAuthenticated
    │   │   │   │   │   │                       # - user
    │   │   │   │   │   │                       # - login, logout
    │   │   │   │   │   │                       # - refresh token
    │   │   │   │   │   ├── useSession.ts       # Session management
    │   │   │   │   │   ├── usePermissions.ts   # RBAC permissions
    │   │   │   │   │   └── useKeycloak.ts      # Keycloak integration
    │   │   │   │   ├── stores/
    │   │   │   │   │   └── auth-store.ts       # Zustand auth store
    │   │   │   │   │                           # - user
    │   │   │   │   │                           # - tokens
    │   │   │   │   │                           # - login/logout actions
    │   │   │   │   │                           # - Persisted to storage
    │   │   │   │   ├── utils/
    │   │   │   │   │   ├── token.ts            # Token utilities
    │   │   │   │   │   ├── session.ts          # Session utilities
    │   │   │   │   │   └── rbac.ts             # RBAC utilities
    │   │   │   │   └── types.ts                # Auth types
    │   │   │   ├── jobs/                       # Jobs feature
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useJobs.ts          # List jobs query
    │   │   │   │   │   ├── useJob.ts           # Single job query
    │   │   │   │   │   ├── useCreateJob.ts     # Create job mutation
    │   │   │   │   │   ├── useUpdateJob.ts     # Update job mutation
    │   │   │   │   │   ├── useDeleteJob.ts     # Delete job mutation
    │   │   │   │   │   ├── useSaveJob.ts       # Save/bookmark job
    │   │   │   │   │   └── useJobFilters.ts    # Job filters state
    │   │   │   │   ├── queries/
    │   │   │   │   │   ├── jobs-queries.ts     # TanStack Query queries
    │   │   │   │   │   │                       # BE: jobs-be
    │   │   │   │   │   │                       # GET /v1/jobs
    │   │   │   │   │   └── jobs-mutations.ts   # TanStack Query mutations
    │   │   │   │   │                           # BE: jobs-be
    │   │   │   │   │                           # POST /v1/jobs
    │   │   │   │   │                           # PUT /v1/jobs/{id}
    │   │   │   │   ├── api/
    │   │   │   │   │   └── jobs-api.ts         # Jobs API client
    │   │   │   │   │                           # Axios/Fetch wrapper
    │   │   │   │   │                           # All jobs-be endpoints
    │   │   │   │   ├── utils/
    │   │   │   │   │   ├── job-helpers.ts      # Job utility functions
    │   │   │   │   │   └── job-validation.ts   # Job validation (Zod)
    │   │   │   │   └── types.ts                # Jobs types
    │   │   │   ├── proposals/                  # Proposals feature
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useProposals.ts
    │   │   │   │   │   ├── useProposal.ts
    │   │   │   │   │   ├── useSubmitProposal.ts
    │   │   │   │   │   ├── useWithdrawProposal.ts
    │   │   │   │   │   └── useBidding.ts       # Bidding hooks
    │   │   │   │   ├── queries/
    │   │   │   │   │   ├── proposals-queries.ts # BE: proposals-be
    │   │   │   │   │   └── proposals-mutations.ts
    │   │   │   │   ├── api/
    │   │   │   │   │   └── proposals-api.ts
    │   │   │   │   └── types.ts
    │   │   │   ├── contracts/                  # Contracts feature
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useContracts.ts
    │   │   │   │   │   ├── useContract.ts
    │   │   │   │   │   ├── useMilestones.ts
    │   │   │   │   │   ├── useTimesheets.ts
    │   │   │   │   │   ├── useWorkDiary.ts
    │   │   │   │   │   ├── useDeliverables.ts
    │   │   │   │   │   └── useDisputes.ts
    │   │   │   │   ├── queries/
    │   │   │   │   │   ├── contracts-queries.ts # BE: contracts-be
    │   │   │   │   │   └── contracts-mutations.ts
    │   │   │   │   ├── api/
    │   │   │   │   │   └── contracts-api.ts
    │   │   │   │   └── types.ts
    │   │   │   ├── messages/                   # Messaging feature
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useConversations.ts
    │   │   │   │   │   ├── useConversation.ts
    │   │   │   │   │   ├── useMessages.ts
    │   │   │   │   │   ├── useSendMessage.ts
    │   │   │   │   │   └── useRealtimeMessages.ts # WebSocket
    │   │   │   │   ├── queries/
    │   │   │   │   │   ├── messages-queries.ts  # BE: communications-be
    │   │   │   │   │   └── messages-mutations.ts
    │   │   │   │   ├── api/
    │   │   │   │   │   ├── messages-api.ts
    │   │   │   │   │   └── websocket.ts        # WebSocket client
    │   │   │   │   │                           # WS: ws://communications-be/v1/realtime
    │   │   │   │   └── types.ts
    │   │   │   ├── notifications/              # Notifications feature
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useNotifications.ts
    │   │   │   │   │   ├── useUnreadCount.ts
    │   │   │   │   │   ├── useMarkAsRead.ts
    │   │   │   │   │   └── useRealtimeNotifications.ts # WebSocket
    │   │   │   │   ├── queries/
    │   │   │   │   │   ├── notifications-queries.ts # BE: communications-be
    │   │   │   │   │   └── notifications-mutations.ts
    │   │   │   │   ├── api/
    │   │   │   │   │   └── notifications-api.ts
    │   │   │   │   └── types.ts
    │   │   │   ├── profile/                    # Profile feature
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useProfile.ts
    │   │   │   │   │   ├── useUpdateProfile.ts
    │   │   │   │   │   ├── useSkills.ts
    │   │   │   │   │   ├── useExperience.ts
    │   │   │   │   │   ├── useEducation.ts
    │   │   │   │   │   ├── usePortfolio.ts
    │   │   │   │   │   └── useServiceCatalog.ts
    │   │   │   │   ├── queries/
    │   │   │   │   │   ├── profile-queries.ts   # BE: users-be
    │   │   │   │   │   └── profile-mutations.ts
    │   │   │   │   ├── api/
    │   │   │   │   │   └── profile-api.ts
    │   │   │   │   └── types.ts
    │   │   │   ├── financial/                  # Financial feature
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useWallet.ts
    │   │   │   │   │   ├── useTransactions.ts
    │   │   │   │   │   ├── useInvoices.ts
    │   │   │   │   │   ├── usePaymentMethods.ts
    │   │   │   │   │   ├── usePayoutMethods.ts
    │   │   │   │   │   └── useEscrow.ts
    │   │   │   │   ├── queries/
    │   │   │   │   │   ├── financial-queries.ts # BE: financial-be
    │   │   │   │   │   └── financial-mutations.ts
    │   │   │   │   ├── api/
    │   │   │   │   │   └── financial-api.ts
    │   │   │   │   └── types.ts
    │   │   │   ├── reviews/                    # Reviews feature
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useReviews.ts
    │   │   │   │   │   ├── useCreateReview.ts
    │   │   │   │   │   ├── useReviewStats.ts
    │   │   │   │   │   └── useBadges.ts
    │   │   │   │   ├── queries/
    │   │   │   │   │   ├── reviews-queries.ts   # BE: reviews-be
    │   │   │   │   │   └── reviews-mutations.ts
    │   │   │   │   ├── api/
    │   │   │   │   │   └── reviews-api.ts
    │   │   │   │   └── types.ts
    │   │   │   ├── search/                     # Search feature
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useJobSearch.ts
    │   │   │   │   │   ├── useFreelancerSearch.ts
    │   │   │   │   │   ├── useSavedSearches.ts
    │   │   │   │   │   └── useSearchSuggestions.ts
    │   │   │   │   ├── queries/
    │   │   │   │   │   ├── search-queries.ts    # BE: search-be
    │   │   │   │   │   └── search-mutations.ts
    │   │   │   │   ├── api/
    │   │   │   │   │   └── search-api.ts
    │   │   │   │   └── types.ts
    │   │   │   ├── subscriptions/              # Subscriptions feature
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useSubscription.ts
    │   │   │   │   │   ├── usePlans.ts
    │   │   │   │   │   ├── useConnects.ts
    │   │   │   │   │   ├── useEntitlements.ts
    │   │   │   │   │   └── useUpgrade.ts
    │   │   │   │   ├── queries/
    │   │   │   │   │   ├── subscription-queries.ts # BE: subscriptions-be
    │   │   │   │   │   └── subscription-mutations.ts
    │   │   │   │   ├── api/
    │   │   │   │   │   └── subscriptions-api.ts
    │   │   │   │   └── types.ts
    │   │   │   ├── storage/                    # File storage feature
    │   │   │   │   ├── hooks/
    │   │   │   │   │   ├── useUpload.ts
    │   │   │   │   │   ├── usePresignedUrl.ts
    │   │   │   │   │   └── useFileDownload.ts
    │   │   │   │   ├── api/
    │   │   │   │   │   └── storage-api.ts       # BE: storage-be
    │   │   │   │   │                           # POST /v1/storage/upload
    │   │   │   │   │                           # GET /v1/storage/presign
    │   │   │   │   └── types.ts
    │   │   │   └── admin/                      # Admin feature
    │   │   │       ├── hooks/
    │   │   │       │   ├── useAdminUsers.ts
    │   │   │       │   ├── useModeration.ts
    │   │   │       │   ├── useKYCCases.ts
    │   │   │       │   └── useDisputes.ts
    │   │   │       ├── queries/
    │   │   │       │   ├── admin-queries.ts     # BE: admin-be
    │   │   │       │   └── admin-mutations.ts
    │   │   │       ├── api/
    │   │   │       │   └── admin-api.ts
    │   │   │       └── types.ts
    │   │   ├── hooks/                          # General hooks
    │   │   │   ├── useDebounce.ts
    │   │   │   ├── useLocalStorage.ts
    │   │   │   ├── useMediaQuery.ts
    │   │   │   ├── useClickOutside.ts
    │   │   │   ├── useCopyToClipboard.ts
    │   │   │   ├── useIntersectionObserver.ts
    │   │   │   └── useToggle.ts
    │   │   ├── lib/                            # Shared utilities
    │   │   │   ├── api/
    │   │   │   │   ├── client.ts               # Axios/Fetch client setup
    │   │   │   │   │                           # - Base URL
    │   │   │   │   │                           # - Auth interceptors
    │   │   │   │   │                           # - Error handling
    │   │   │   │   │                           # - Request/response logging
    │   │   │   │   ├── endpoints.ts            # API endpoint constants
    │   │   │   │   └── error-handler.ts        # Global error handler
    │   │   │   ├── websocket/
    │   │   │   │   ├── client.ts               # WebSocket client
    │   │   │   │   │                           # WS: ws://communications-be/v1/realtime
    │   │   │   │   │                           # - Connection management
    │   │   │   │   │                           # - Reconnection logic
    │   │   │   │   │                           # - Event subscriptions
    │   │   │   │   └── events.ts               # WebSocket event types
    │   │   │   ├── validation/
    │   │   │   │   ├── schemas.ts              # Zod schemas
    │   │   │   │   └── validators.ts           # Custom validators
    │   │   │   ├── formatting/
    │   │   │   │   ├── date.ts                 # Date formatting
    │   │   │   │   ├── currency.ts             # Currency formatting
    │   │   │   │   ├── number.ts               # Number formatting
    │   │   │   │   └── string.ts               # String utilities
    │   │   │   └── constants/
    │   │   │       ├── api.ts                  # API constants
    │   │   │       ├── app.ts                  # App constants
    │   │   │       ├── routes.ts               # Route constants
    │   │   │       └── permissions.ts          # RBAC permissions
    │   │   └── i18n/                           # Internationalization
    │   │       ├── locales/                    # Locale files
    │   │       │   ├── en/
    │   │       │   │   ├── common.json
    │   │       │   │   ├── auth.json
    │   │       │   │   ├── jobs.json
    │   │       │   │   ├── proposals.json
    │   │       │   │   ├── contracts.json
    │   │       │   │   ├── messages.json
    │   │       │   │   ├── profile.json
    │   │       │   │   ├── financial.json
    │   │       │   │   ├── reviews.json
    │   │       │   │   ├── settings.json
    │   │       │   │   ├── subscription.json
    │   │       │   │   ├── admin.json
    │   │       │   │   └── errors.json
    │   │       │   ├── ar/                     # Arabic
    │   │       │   ├── zh/                     # Chinese
    │   │       │   ├── hi/                     # Hindi
    │   │       │   ├── de/                     # German
    │   │       │   ├── fr/                     # French
    │   │       │   ├── tr/                     # Turkish
    │   │       │   ├── es/                     # Spanish
    │   │       │   └── ru/                     # Russian
    │   │       └── config.ts                   # i18n configuration
    │   ├── package.json
    │   ├── tsconfig.json
    │   └── README.md
    │
    ├── types/                                   # Shared TypeScript types
    │   ├── src/
    │   │   ├── api/                            # API response types
    │   │   │   ├── jobs.ts
    │   │   │   ├── proposals.ts
    │   │   │   ├── contracts.ts
    │   │   │   ├── messages.ts
    │   │   │   ├── notifications.ts
    │   │   │   ├── users.ts
    │   │   │   ├── financial.ts
    │   │   │   ├── reviews.ts
    │   │   │   ├── search.ts
    │   │   │   ├── subscriptions.ts
    │   │   │   ├── storage.ts
    │   │   │   ├── admin.ts
    │   │   │   └── common.ts                   # Common types (pagination, filters, etc.)
    │   │   ├── entities/                       # Domain entities
    │   │   │   ├── user.ts
    │   │   │   ├── job.ts
    │   │   │   ├── proposal.ts
    │   │   │   ├── contract.ts
    │   │   │   ├── message.ts
    │   │   │   ├── notification.ts
    │   │   │   ├── transaction.ts
    │   │   │   ├── invoice.ts
    │   │   │   ├── review.ts
    │   │   │   └── subscription.ts
    │   │   ├── enums/                          # Enums
    │   │   │   ├── job-status.ts
    │   │   │   ├── proposal-status.ts
    │   │   │   ├── contract-status.ts
    │   │   │   ├── user-type.ts
    │   │   │   ├── payment-status.ts
    │   │   │   └── review-rating.ts
    │   │   └── index.ts                        # Export all types
    │   ├── package.json
    │   ├── tsconfig.json
    │   └── README.md
    │
    └── config/                                  # Shared configurations
        ├── eslint-config/
        │   ├── index.js                        # Base ESLint config
        │   ├── react.js                        # React-specific rules
        │   ├── next.js                         # Next.js-specific rules
        │   └── package.json
        ├── typescript-config/
        │   ├── base.json                       # Base TS config
        │   ├── nextjs.json                     # Next.js TS config
        │   ├── react-native.json               # React Native TS config
        │   └── package.json
        └── tailwind-config/
            ├── themes/
            │   ├── light.js
            │   └── dark.js
            ├── index.js                        # Base Tailwind config
            │                                    # - Design tokens
            │                                    # - Color palette
            │                                    # - Spacing scale
            │                                    # - Typography
            │                                    # - Breakpoints
            │                                    # - Animations
            └── package.json
