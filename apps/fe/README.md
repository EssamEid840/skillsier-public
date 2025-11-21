# Skillsier Frontend Monorepo

Production-ready freelancing platform frontend built with Next.js 15, Expo, and pnpm workspaces.

## 🚀 Tech Stack

- **Monorepo:** pnpm 10.x + Turborepo 3.x
- **Web:** Next.js 15.1.x (App Router), React 19.x
- **Mobile:** Expo SDK 52.x + Expo Router 4.x
- **UI:** Tailwind CSS 4.x + NativeWind 4.x + shadcn/ui
- **State:** TanStack Query 6.x + Zustand 5.x
- **Auth:** Dev adapter (Keycloak-ready)
- **Forms:** React Hook Form 7.x + Zod 3.x
- **Testing:** Vitest 3.x + Playwright 1.x

## 📋 Prerequisites

- Node.js 20.18.1+ (use nvm: `nvm use`)
- pnpm 10.0.0+
- For mobile: Expo CLI, Android Studio / Xcode

## ⚡ Quick Start
```bash
# Install dependencies
pnpm install

# Copy environment variables
cp .env.example .env

# Run web app
pnpm dev:web

# Run mobile app
pnpm dev:mobile

# Run all apps
pnpm dev
```

## 📱 Mobile Development
```bash
# Start Metro bundler
pnpm --filter mobile start

# Run on Android
pnpm --filter mobile android

# Run on iOS
pnpm --filter mobile ios

# Build for production
pnpm --filter mobile build:android
pnpm --filter mobile build:ios
```

## 🧪 Testing
```bash
# Unit tests
pnpm test

# Watch mode
pnpm test:watch

# E2E tests (Playwright)
pnpm test:e2e

# Test coverage
pnpm test -- --coverage
```

## 🔨 Development Scripts
```bash
pnpm dev              # Run all apps in dev mode
pnpm dev:web          # Run web app only
pnpm dev:mobile       # Run mobile app only
pnpm build            # Build all apps
pnpm build:web        # Build web app
pnpm build:mobile     # Build mobile app
pnpm lint             # Lint all packages
pnpm typecheck        # TypeScript check
pnpm test             # Run tests
pnpm test:e2e         # Run E2E tests
pnpm clean            # Clean all builds
pnpm format           # Format code
```

## 📁 Project Structure
```
skillsier-fe/
├── apps/
│   ├── web/                    # Next.js web application
│   │   ├── app/               # App Router pages
│   │   ├── components/        # Web components
│   │   └── public/            # Static assets
│   └── mobile/                # Expo mobile application
│       ├── app/               # Expo Router pages
│       └── components/        # Mobile components
├── packages/
│   ├── ui/                    # Shared UI components
│   ├── api/                   # API clients
│   ├── auth/                  # Authentication
│   ├── types/                 # TypeScript types
│   ├── hooks/                 # React hooks
│   ├── stores/                # Zustand stores
│   └── i18n/                  # Internationalization
├── docs/                      # Documentation
└── .github/                   # CI/CD workflows
```

## 🔐 Authentication

Default dev adapter with seeded accounts:
- **Admin:** `admin@skillsier.dev` / `admin123`
- **Client:** `client@skillsier.dev` / `client123`
- **Freelancer:** `freelancer@skillsier.dev` / `freelancer123`

To switch to Keycloak production auth, see [docs/KEYCLOAK.md](./docs/KEYCLOAK.md)

## 🌍 Internationalization

Supports English (en) and Arabic (ar) with RTL support.
See [docs/I18N.md](./docs/I18N.md) for adding new languages.

## 📚 Documentation

- [Setup Guide](./docs/SETUP.md) - Detailed setup instructions
- [Keycloak Setup](./docs/KEYCLOAK.md) - Production authentication
- [Internationalization](./docs/I18N.md) - Adding languages
- [Architecture](./docs/ARCHITECTURE.md) - System architecture

## 🏗️ Architecture

- **Monorepo:** Workspace packages with pnpm
- **Type Safety:** Strict TypeScript throughout
- **Component Library:** Shared UI with platform variants
- **State Management:** TanStack Query + Zustand
- **API Layer:** Typed clients with mock adapters
- **Auth:** Pluggable adapters (dev/Keycloak)

See [docs/ARCHITECTURE.md](./docs/ARCHITECTURE.md) for details.

## 🔄 Development Workflow

1. Create feature branch from `develop`
2. Make changes and commit (pre-commit hooks run)
3. Push (pre-push hooks run tests)
4. Create PR to `develop`
5. CI runs lint, typecheck, tests, build
6. Merge to `develop` for staging
7. Merge to `main` for production deploy

## 🚢 Deployment

- **Web:** Vercel (auto-deploy from `main`)
- **Mobile:** EAS Build (manual trigger)

See [docs/SETUP.md](./docs/SETUP.md) for deployment details.

## 📄 License

Proprietary - All rights reserved

## 🤝 Contributing

1. Follow the established patterns
2. Write tests for new features
3. Update documentation
4. Ensure all checks pass before PR

## 🆘 Support

For issues or questions:
- Check documentation in `/docs`
- Review existing issues
- Contact development team