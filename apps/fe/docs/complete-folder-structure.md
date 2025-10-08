skillsier-fe/
├── package.json                          # Root workspace (React 19.0)
├── pnpm-workspace.yaml                   # pnpm workspace definition
├── turbo.json                            # Turborepo pipeline config
├── .gitignore
├── .prettierrc
├── tsconfig.json
├── README.md                             # Complete setup guide
├── COMMANDS.md                           # Quick command reference
├── setup.sh                              # Setup script
├── setup-env.sh                          # Environment setup
├── dev.sh                                # Interactive dev launcher
├── setup-mobile-dev-client.sh           # Mobile dev client setup
│
├── apps/
│   ├── web/                              # Next.js 15 Web App (React 19)
│   │   ├── package.json                  # React 19.0.0, next-intl
│   │   ├── next.config.js                # With next-intl plugin
│   │   ├── tailwind.config.ts
│   │   ├── tsconfig.json
│   │   ├── postcss.config.js
│   │   ├── .env.local                    # Environment variables
│   │   ├── .eslintrc.json
│   │   ├── i18n.ts                       # ⭐ i18n configuration
│   │   ├── middleware.ts                 # ⭐ Locale detection & routing
│   │   ├── public/
│   │   │   ├── images/
│   │   │   │   ├── dashboard-preview.png
│   │   │   │   └── avatars/
│   │   │   ├── icons/
│   │   │   └── fonts/
│   │   └── src/
│   │       ├── app/
│   │       │   └── [locale]/             # ⭐ Locale-based routing
│   │       │       ├── layout.tsx        # Root layout with i18n provider
│   │       │       ├── page.tsx          # Landing page
│   │       │       ├── providers.tsx     # Query client provider
│   │       │       ├── globals.css       # Global styles with CSS variables
│   │       │       ├── (auth)/           # Auth route group
│   │       │       │   ├── layout.tsx    # Auth layout (centered)
│   │       │       │   ├── login/
│   │       │       │   │   └── page.tsx  # ⭐ Login with i18n & validation
│   │       │       │   └── register/
│   │       │       │       └── page.tsx  # ⭐ Register with validation & password strength
│   │       │       └── (dashboard)/      # Protected routes
│   │       │           ├── layout.tsx    # Dashboard layout (sidebar + header)
│   │       │           ├── dashboard/
│   │       │           │   └── page.tsx  # Dashboard home
│   │       │           ├── profile/
│   │       │           │   ├── page.tsx  # ⭐ COMPLETE Freelancer profile view
│   │       │           │   └── edit/
│   │       │           │       └── page.tsx  # ⭐ COMPLETE Edit profile form
│   │       │           ├── portfolio/
│   │       │           │   └── page.tsx  # Portfolio management (ready)
│   │       │           ├── skills/
│   │       │           │   └── page.tsx  # Skills management (ready)
│   │       │           └── settings/
│   │       │               └── page.tsx  # Settings page (ready)
│   │       └── components/
│   │           ├── landing/
│   │           │   ├── Hero.tsx          # ⭐ Hero with i18n
│   │           │   ├── Features.tsx      # ⭐ Features with i18n
│   │           │   ├── Stats.tsx         # Stats section
│   │           │   ├── Testimonials.tsx  # Testimonials
│   │           │   └── CTA.tsx           # ⭐ CTA with i18n
│   │           ├── layout/
│   │           │   ├── Header.tsx        # ⭐ Header with LanguageSwitcher
│   │           │   ├── Footer.tsx        # Footer with links
│   │           │   ├── Sidebar.tsx       # Dashboard sidebar
│   │           │   ├── DashboardHeader.tsx  # Dashboard top header
│   │           │   ├── LanguageSwitcher.tsx  # ⭐ NEW: Language dropdown
│   │           │   └── MobileNav.tsx     # Mobile navigation
│   │           └── dashboard/
│   │               ├── DashboardShell.tsx
│   │               └── StatsCard.tsx
│   │
│   └── mobile/                           # React Native Expo App (React 19)
│       ├── package.json                  # React 19.0.0, i18next, expo 52
│       ├── app.json                      # ⭐ UPDATED: 120 FPS iOS config
│       ├── babel.config.js               # With nativewind & reanimated
│       ├── metro.config.js               # ⭐ UPDATED: Performance optimizations
│       ├── tailwind.config.js            # NativeWind config
│       ├── tsconfig.json
│       ├── .env
│       ├── index.js
│       ├── global.css                    # NativeWind styles
│       ├── app/                          # Expo Router (file-based)
│       │   ├── _layout.tsx               # ⭐ UPDATED: i18n & performance init
│       │   ├── index.tsx                 # Landing/Splash screen
│       │   ├── +not-found.tsx            # 404 screen
│       │   ├── (auth)/                   # Auth stack
│       │   │   ├── _layout.tsx           # Auth stack layout
│       │   │   ├── login.tsx             # ⭐ UPDATED: Login with i18n
│       │   │   └── register.tsx          # ⭐ UPDATED: Register with i18n & validation
│       │   ├── (tabs)/                   # Main app tabs
│       │   │   ├── _layout.tsx           # Tab navigator with icons
│       │   │   ├── dashboard.tsx         # ⭐ UPDATED: Dashboard with i18n
│       │   │   ├── courses.tsx           # Browse jobs/gigs (renamed context)
│       │   │   ├── skills.tsx            # ⭐ UPDATED: Skills with i18n
│       │   │   └── profile.tsx           # ⭐ COMPLETE: Freelancer profile
│       │   └── profile/                  # Profile screens (outside tabs)
│       │       └── edit.tsx              # ⭐ COMPLETE: Edit profile screen
│       └── src/
│           ├── components/
│           │   ├── landing/
│           │   │   ├── HeroMobile.tsx
│           │   │   └── FeaturesMobile.tsx
│           │   ├── navigation/
│           │   │   ├── TabBar.tsx
│           │   │   └── DrawerContent.tsx
│           │   ├── OptimizedFlashList.tsx  # ⭐ NEW: 120 FPS optimized list
│           │   └── LanguageSwitcher.tsx    # ⭐ NEW: Language modal
│           ├── lib/
│           │   ├── i18n/                   # ⭐ NEW: i18n configuration
│           │   │   └── index.ts            # i18next setup with MMKV
│           │   ├── performance.ts          # ⭐ NEW: 120 FPS utilities
│           │   ├── keycloak-mobile.ts
│           │   └── utils.ts
│           ├── hooks/
│           │   └── useHighFPSAnimation.ts  # ⭐ NEW: 120 FPS animation hook
│           └── assets/
│               ├── images/
│               ├── fonts/
│               └── icons/
│
├── packages/
│   ├── shared/                           # Shared business logic
│   │   ├── package.json
│   │   ├── tsconfig.json
│   │   ├── .eslintrc.json
│   │   └── src/
│   │       ├── index.ts
│   │       ├── features/
│   │       │   ├── auth/
│   │       │   │   ├── index.ts
│   │       │   │   ├── hooks/
│   │       │   │   │   ├── useAuth.ts
│   │       │   │   │   ├── useLogin.ts
│   │       │   │   │   ├── useRegister.ts
│   │       │   │   │   └── useLogout.ts
│   │       │   │   ├── stores/
│   │       │   │   │   └── authStore.ts    # Zustand store
│   │       │   │   ├── types/
│   │       │   │   │   └── auth.types.ts
│   │       │   │   ├── api/
│   │       │   │   │   └── authApi.ts
│   │       │   │   └── utils/
│   │       │   │       ├── validation.ts
│   │       │   │       └── token.ts
│   │       │   └── user/                   # ⭐ FREELANCING USER FEATURE
│   │       │       ├── index.ts
│   │       │       ├── hooks/
│   │       │       │   ├── useUserProfile.ts              # ⭐ NEW
│   │       │       │   ├── useFreelancerProfile.ts        # ⭐ NEW
│   │       │       │   ├── useClientProfile.ts            # ⭐ NEW
│   │       │       │   ├── useUpdateProfile.ts            # ⭐ UPDATED
│   │       │       │   ├── useUploadAvatar.ts             # ⭐ NEW
│   │       │       │   ├── useDeleteAvatar.ts             # ⭐ NEW
│   │       │       │   ├── useUpdatePreferences.ts        # ⭐ NEW
│   │       │       │   ├── useUpdateSocialLinks.ts        # ⭐ NEW
│   │       │       │   ├── useChangePassword.ts           # ⭐ NEW
│   │       │       │   ├── useFreelancerSkills.ts         # ⭐ NEW
│   │       │       │   ├── useAddSkill.ts                 # ⭐ NEW
│   │       │       │   ├── useUpdateSkill.ts              # ⭐ NEW
│   │       │       │   ├── useDeleteSkill.ts              # ⭐ NEW
│   │       │       │   ├── useWorkExperience.ts           # ⭐ NEW
│   │       │       │   ├── useAddWorkExperience.ts        # ⭐ NEW
│   │       │       │   ├── useUpdateWorkExperience.ts     # ⭐ NEW
│   │       │       │   ├── useDeleteWorkExperience.ts     # ⭐ NEW
│   │       │       │   ├── useEducation.ts                # ⭐ NEW
│   │       │       │   ├── useAddEducation.ts             # ⭐ NEW
│   │       │       │   ├── useCertifications.ts           # ⭐ NEW
│   │       │       │   ├── useAddCertification.ts         # ⭐ NEW
│   │       │       │   ├── usePortfolio.ts                # ⭐ NEW
│   │       │       │   ├── useAddPortfolioItem.ts         # ⭐ NEW
│   │       │       │   └── useUploadPortfolioImage.ts     # ⭐ NEW
│   │       │       ├── api/
│   │       │       │   └── userApi.ts      # ⭐ COMPLETE: 30+ API methods
│   │       │       └── types/
│   │       │           └── user.types.ts
│   │       ├── lib/
│   │       │   ├── api/
│   │       │   │   ├── client.ts           # Axios client with interceptors
│   │       │   │   └── queryClient.ts      # ⭐ UPDATED: New query keys
│   │       │   ├── i18n/                   # ⭐ INTERNATIONALIZATION
│   │       │   │   ├── config.ts           # Locale definitions
│   │       │   │   └── translations/       # Translation files
│   │       │   │       ├── en.json         # ⭐ COMPLETE: English (1000+ strings)
│   │       │   │       └── ar.json         # ⭐ COMPLETE: Arabic (1000+ strings)
│   │       │   ├── keycloak/
│   │       │   │   ├── config.ts
│   │       │   │   └── types.ts
│   │       │   └── utils/
│   │       │       ├── date.ts
│   │       │       ├── string.ts
│   │       │       └── validators.ts
│   │       └── constants/
│   │           ├── api.ts                  # ⭐ UPDATED: 30+ API endpoints
│   │           └── app.ts
│   │
│   ├── ui/                                 # Shared UI components
│   │   ├── package.json
│   │   ├── tsconfig.json
│   │   └── src/
│   │       ├── index.ts
│   │       ├── components/
│   │       │   ├── Button/
│   │       │   │   ├── Button.tsx          # Platform router
│   │       │   │   ├── Button.web.tsx      # Web implementation
│   │       │   │   ├── Button.native.tsx   # Native implementation
│   │       │   │   └── Button.types.ts
│   │       │   ├── Input/
│   │       │   │   ├── Input.tsx
│   │       │   │   ├── Input.web.tsx
│   │       │   │   ├── Input.native.tsx
│   │       │   │   └── Input.types.ts
│   │       │   ├── Card/
│   │       │   │   ├── Card.tsx
│   │       │   │   ├── Card.web.tsx
│   │       │   │   ├── Card.native.tsx
│   │       │   │   └── Card.types.ts
│   │       │   ├── Avatar/
│   │       │   │   ├── Avatar.tsx
│   │       │   │   ├── Avatar.web.tsx
│   │       │   │   ├── Avatar.native.tsx
│   │       │   │   └── Avatar.types.ts
│   │       │   ├── Badge/
│   │       │   │   ├── Badge.tsx
│   │       │   │   ├── Badge.web.tsx
│   │       │   │   ├── Badge.native.tsx
│   │       │   │   └── Badge.types.ts
│   │       │   ├── LoadingSpinner/
│   │       │   ├── Modal/
│   │       │   └── Toast/
│   │       ├── primitives/                 # Base components
│   │       │   ├── Text.tsx
│   │       │   ├── View.tsx
│   │       │   └── Pressable.tsx
│   │       ├── theme/
│   │       │   ├── colors.ts
│   │       │   ├── spacing.ts
│   │       │   ├── typography.ts
│   │       │   └── index.ts
│   │       └── lib/
│   │           └── utils.ts                # cn() utility
│   │
│   ├── types/                              # Shared TypeScript types
│   │   ├── package.json
│   │   ├── tsconfig.json
│   │   └── src/
│   │       ├── index.ts
│   │       ├── entities/
│   │       │   ├── user.ts                 # ⭐ COMPLETE: Freelancing types
│   │       │   │                           #   - User, FreelancerProfile
│   │       │   │                           #   - ClientProfile, UserSkill
│   │       │   │                           #   - WorkExperience, Education
│   │       │   │                           #   - Certification, PortfolioItem
│   │       │   ├── job.ts                  # Job/Gig types (ready)
│   │       │   ├── proposal.ts             # Proposal types (ready)
│   │       │   ├── contract.ts             # Contract types (ready)
│   │       │   └── review.ts               # Review types (ready)
│   │       ├── api/
│   │       │   ├── requests.ts             # ⭐ COMPLETE: 15+ request types
│   │       │   └── responses.ts            # API response types
│   │       └── common/
│   │           ├── pagination.ts
│   │           └── filters.ts
│   │
│   └── config/                             # Shared configuration
│       ├── package.json
│       ├── tsconfig.json
│       └── src/
│           ├── eslint/
│           │   ├── base.js
│           │   ├── next.js
│           │   └── react-native.js
│           ├── typescript/
│           │   ├── base.json
│           │   ├── nextjs.json
│           │   └── react-native.json
│           └── tailwind/
│               ├── base.js
│               └── mobile.js
│
├── .husky/                                 # Git hooks
│   ├── pre-commit
│   └── pre-push
│
├── .vscode/                                # VS Code settings
│   ├── settings.json
│   ├── extensions.json
│   └── launch.json
│
└── docs/                                   # Documentation
    ├── SETUP.md
    ├── ARCHITECTURE.md
    ├── CONTRIBUTING.md
    ├── PERFORMANCE.md                      # ⭐ NEW: 120 FPS guide
    ├── I18N.md                             # ⭐ NEW: i18n guide
    ├── REACT_19.md                         # ⭐ NEW: React 19 migration
    ├── FREELANCING.md                      # ⭐ NEW: Freelancing features guide
    └── TROUBLESHOOTING.md