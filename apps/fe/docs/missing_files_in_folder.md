skillsier-fe/

├── apps/
│   ├── web/                              # Next.js 15 Web App (React 19)
│   │   ├── public/
│   │   │   ├── images/
│   │   │   │   ├── dashboard-preview.png
│   │   │   │   └── avatars/
│   │   │   ├── icons/
│   │   │   └── fonts/
│   │
│   └── mobile/                           # React Native Expo App (React 19)
│       ├── .env
│       ├── app/                          # Expo Router (file-based)
│       │   ├── (tabs)/                   # Main app tabs
│       │   │   ├── courses.tsx           # Browse jobs/gigs (renamed context)
│       └── src/
│           └── assets/
│               ├── images/
│               ├── fonts/
│               └── icons/
│
├── packages/
│   ├── shared/                           # Shared business logic
│   │   └── src/
│   │       ├── features/
│   │       │   ├── auth/
│   │       │   │   └── utils/
│   │       │   │       └── token.ts
│   │       │   └── user/                   # ⭐ FREELANCING USER FEATURE
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