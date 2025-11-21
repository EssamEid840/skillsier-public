# Skillsier Frontend Monorepo

Production-ready freelancing platform frontend built with Next.js 15, Expo, and pnpm workspaces.

## Tech Stack

- **Monorepo:** pnpm 10.x + Turborepo 3.x
- **Web:** Next.js 15.1.x (App Router), React 19.x
- **Mobile:** Expo SDK 52.x + Expo Router 4.x
- **UI:** Tailwind CSS 4.x + NativeWind 4.x + shadcn/ui
- **State:** TanStack Query 6.x + Zustand 5.x
- **Auth:** Dev adapter (Keycloak-ready)

## Prerequisites

- Node.js 20.18.1+ (use nvm: `nvm use`)
- pnpm 10.0.0+

## Quick Start

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

## Scripts

```bash
pnpm dev              # Run all apps in dev mode
pnpm dev:web          # Run web app only
pnpm dev:mobile       # Run mobile app only
pnpm build            # Build all apps
pnpm lint             # Lint all packages
pnpm typecheck        # TypeScript check
pnpm test             # Run tests
pnpm test:e2e         # Run E2E tests
pnpm format           # Format code
```

## Project Structure

```
skillsier-fe/
├── apps/
│   ├── web/          # Next.js web application
│   └── mobile/       # Expo mobile application
├── packages/
│   ├── ui/           # Shared UI components
│   ├── api/          # API clients
│   ├── auth/         # Authentication
│   ├── types/        # TypeScript types
│   ├── hooks/        # React hooks
│   ├── stores/       # Zustand stores
│   └── i18n/         # Internationalization
└── .github/          # CI/CD workflows
```

## Development

- Web runs on `http://localhost:3000`
- Mobile Metro bundler on `http://localhost:8081`
- API base URL: `http://localhost:8080` (configure in `.env`)

## Authentication

Default uses dev adapter with seeded accounts:
- Admin: `admin@skillsier.dev` / `admin123`
- Client: `client@skillsier.dev` / `client123`
- Freelancer: `freelancer@skillsier.dev` / `freelancer123`

To switch to Keycloak, see `docs/KEYCLOAK.md`.

## License

Proprietary