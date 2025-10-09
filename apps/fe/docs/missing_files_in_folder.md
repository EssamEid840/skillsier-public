skillsier-fe/

├── setup-env.sh                          # Environment setup
├── dev.sh                                # Interactive dev launcher
├── setup-mobile-dev-client.sh           # Mobile dev client setup
├── apps/
│   ├── web/                              # Next.js 15 Web App (React 19)
│   │   ├── tsconfig.json
│   │   ├── postcss.config.js
│   │   ├── .env.local                    # Environment variables
│   │   ├── .eslintrc.json
│   │   ├── public/
│   │   │   ├── images/
│   │   │   │   ├── dashboard-preview.png
│   │   │   │   └── avatars/
│   │   │   ├── icons/
│   │   │   └── fonts/
│   │   └── src/
│   │       ├── app/
│   │       │   └── [locale]/             # ⭐ Locale-based routing
│   │       │       ├── globals.css       # Global styles with CSS variables
│   │       │       ├── (auth)/           # Auth route group
│   │       │       │   ├── login/
│   │       └── components/
│   │           ├── layout/
│   │           │   └── MobileNav.tsx     # Mobile navigation
│   │           └── dashboard/
│   │               ├── DashboardShell.tsx
│   │               └── StatsCard.tsx
│   │
│   └── mobile/                           # React Native Expo App (React 19)
│       ├── tsconfig.json
│       ├── .env
│       ├── app/                          # Expo Router (file-based)
│       │   ├── +not-found.tsx            # 404 screen
│       │   ├── (auth)/                   # Auth stack
│       │   │   └── register.tsx          # ⭐ UPDATED: Register with i18n & validation
│       │   ├── (tabs)/                   # Main app tabs
│       │   │   ├── _layout.tsx           # Tab navigator with icons
│       │   │   ├── dashboard.tsx         # ⭐ UPDATED: Dashboard with i18n
│       │   │   ├── courses.tsx           # Browse jobs/gigs (renamed context)
│       │   │   ├── skills.tsx            # ⭐ UPDATED: Skills with i18n
│       └── src/
│           ├── components/
│           │   ├── landing/
│           │   │   ├── HeroMobile.tsx
│           │   │   └── FeaturesMobile.tsx
│           │   ├── navigation/
│           │   │   ├── TabBar.tsx
│           │   │   └── DrawerContent.tsx
│           ├── lib/
│           │   ├── keycloak-mobile.ts
│           │   └── utils.ts
│           └── assets/
│               ├── images/
│               ├── fonts/
│               └── icons/
│
├── packages/
│   ├── shared/                           # Shared business logic
│   │   ├── tsconfig.json
│   │   ├── .eslintrc.json
│   │   └── src/
│   │       ├── features/
│   │       │   ├── auth/
│   │       │   │   └── utils/
│   │       │   │       ├── validation.ts
│   │       │   │       └── token.ts
│   │       │   └── user/                   # ⭐ FREELANCING USER FEATURE
│   │       │       ├── hooks/
│   │       │       │   ├── useClientProfile.ts            # ⭐ NEW
│   │       │       │   ├── useDeleteAvatar.ts             # ⭐ NEW
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
│   │       │       └── types/
│   │       │           └── user.types.ts
│   │       ├── lib/
│   │       │   ├── keycloak/
│   │       │   │   ├── config.ts
│   │       │   │   └── types.ts
│   │       │   └── utils/
│   │       │       ├── date.ts
│   │       │       ├── string.ts
│   │       │       └── validators.ts
│   │       └── constants/
│   │           └── app.ts
│   │
│   ├── ui/                                 # Shared UI components
│   │   ├── tsconfig.json
│   │   └── src/
│   │       ├── components/
│   │       │   ├── LoadingSpinner/
│   │       │   ├── Modal/
│   │       │   └── Toast
│   │       ├── theme/
│   │       │   └── index.ts
│   │
│   ├── types/                              # Shared TypeScript types
│   │   ├── tsconfig.json
│   │   └── src/
│   │       ├── entities/
│   │       │   ├── job.ts                  # Job/Gig types (ready)
│   │       │   ├── proposal.ts             # Proposal types (ready)
│   │       │   ├── contract.ts             # Contract types (ready)
│   │       │   └── review.ts               # Review types (ready)
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
│   └── launch.json
│

```