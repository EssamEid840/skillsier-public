skillsier-fe/
├── package.json                          # Root workspace (React 19.0)
├── pnpm-workspace.yaml                   # pnpm workspace definition
├── turbo.json                            # Turborepo pipeline config
├── .gitignore
├── .prettierrc
├── tsconfig.json
├── setup.sh                              # Setup script
├── apps/
│   ├── web/                              # Next.js 15 Web App (React 19)
│   │   ├── package.json                  # React 19.0.0, next-intl
│   │   ├── next.config.js                # With next-intl plugin
│   │   ├── tailwind.config.ts
│   │   ├── i18n.ts                       # ⭐ i18n configuration
│   │   ├── middleware.ts                 # ⭐ Locale detection & routing
│   │   └── src/
│   │       ├── app/
│   │       │   └── [locale]/             # ⭐ Locale-based routing
│   │       │       ├── layout.tsx        # Root layout with i18n provider
│   │       │       ├── page.tsx          # Landing page
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
│   │
│   └── mobile/                           # React Native Expo App (React 19)
│       ├── package.json                  # React 19.0.0, i18next, expo 52
│       ├── app.json                      # ⭐ UPDATED: 120 FPS iOS config
│       ├── babel.config.js               # With nativewind & reanimated
│       ├── metro.config.js               # ⭐ UPDATED: Performance optimizations
│       ├── tailwind.config.js            # NativeWind config
│       ├── index.js
│       ├── global.css                    # NativeWind styles
│       ├── app/                          # Expo Router (file-based)
│       │   ├── _layout.tsx               # ⭐ UPDATED: i18n & performance init
│       │   ├── index.tsx                 # Landing/Splash screen
│       │   ├── +not-found.tsx            # 404 screen
│       │   ├── (auth)/                   # Auth stack
│       │   │   ├── _layout.tsx           # Auth stack layout
│       │   │   ├── login.tsx             # ⭐ UPDATED: Login with i18n
│       │   │   └── profile.tsx           # ⭐ COMPLETE: Freelancer profile
│       │   └── profile/                  # Profile screens (outside tabs)
│       │       └── edit.tsx              # ⭐ COMPLETE: Edit profile screen
│       └── src/
│           ├── components/
│           │   ├── OptimizedFlashList.tsx  # ⭐ NEW: 120 FPS optimized list
│           │   └── LanguageSwitcher.tsx    # ⭐ NEW: Language modal
│           ├── lib/
│           │   ├── i18n/                   # ⭐ NEW: i18n configuration
│           │   │   └── index.ts            # i18next setup with MMKV
│           │   ├── performance.ts          # ⭐ NEW: 120 FPS utilities
│           ├── hooks/
│           │   └── useHighFPSAnimation.ts  # ⭐ NEW: 120 FPS animation hook
│
├── packages/
│   ├── shared/                           # Shared business logic
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
│   │       │   └── user/                   # ⭐ FREELANCING USER FEATURE
│   │       │       ├── index.ts
│   │       │       ├── hooks/
│   │       │       │   ├── useUserProfile.ts              # ⭐ NEW
│   │       │       │   ├── useFreelancerProfile.ts        # ⭐ NEW
│   │       │       │   ├── useUpdateProfile.ts            # ⭐ UPDATED
│   │       │       │   ├── useUploadAvatar.ts             # ⭐ NEW
│   │       │       │   ├── useUpdatePreferences.ts        # ⭐ NEW
│   │       │       │   ├── useUpdateSocialLinks.ts        # ⭐ NEW
│   │       │       │   ├── useChangePassword.ts           # ⭐ NEW
│   │       │       │   └── useUserSkills.ts               # ⭐ NEW
│   │       │       │   └── index.ts               # ⭐ NEW
│   │       │       ├── api/
│   │       │       │   └── userApi.ts      # ⭐ COMPLETE: 30+ API methods
│   │       ├── lib/
│   │       │   ├── api/
│   │       │   │   ├── client.ts           # Axios client with interceptors
│   │       │   │   └── queryClient.ts      # ⭐ UPDATED: New query keys
│   │       │   ├── i18n/                   # ⭐ INTERNATIONALIZATION
│   │       │   │   ├── config.ts           # Locale definitions
│   │       │   │   └── translations/       # Translation files
│   │       │   │       ├── en.json         # ⭐ COMPLETE: English (1000+ strings)
│   │       │   │       └── ar.json         # ⭐ COMPLETE: Arabic (1000+ strings)
│   │       └── constants/
│   │           ├── api.ts                  # ⭐ UPDATED: 30+ API endpoints
│   │           └── app.ts
│   │
│   ├── ui/                                 # Shared UI components
│   │   ├── package.json
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
│   │       ├── primitives/                 # Base components
│   │       │   ├── Text.tsx
│   │       │   ├── View.tsx
│   │       │   └── Pressable.tsx
│   │       │   └── index.ts
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
│   │       ├── api/
│   │       │   ├── requests.ts             # ⭐ COMPLETE: 15+ request types
│   │       │   └── responses.ts            # API response types
│   │       └── common/
│   │           ├── pagination.ts
│   │           └── filters.ts
│
├── .vscode/                                # VS Code settings
│   ├── settings.json
│   ├── extensions.json
```