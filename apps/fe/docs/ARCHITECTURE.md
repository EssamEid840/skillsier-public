# Architecture Documentation

Technical architecture and design decisions for Skillsier frontend.

## System Overview
```
┌─────────────────────────────────────────────┐
│           Skillsier Frontend                │
│                                             │
│  ┌──────────┐         ┌─────────────┐     │
│  │   Web    │         │   Mobile    │     │
│  │ Next.js  │         │    Expo     │     │
│  └────┬─────┘         └──────┬──────┘     │
│       │                      │             │
│       └──────────┬───────────┘             │
│                  │                         │
│         ┌────────▼────────┐               │
│         │   Shared Pkgs   │               │
│         │  UI, API, Auth  │               │
│         │  Types, Hooks   │               │
│         └────────┬────────┘               │
│                  │                         │
└──────────────────┼─────────────────────────┘
                   │
                   ▼
         ┌─────────────────┐
         │  Backend APIs   │
         │  Microservices  │
         └─────────────────┘
```

## Technology Decisions

### Monorepo - pnpm Workspaces

**Why:**
- Share code between web/mobile
- Consistent dependencies
- Faster CI/CD with Turborepo
- Type safety across packages

**Structure:**
```
apps/          # Applications
packages/      # Shared packages
```

### Web - Next.js 15 App Router

**Why:**
- Server Components for performance
- File-based routing
- Built-in optimizations
- API routes for BFF pattern

**Key Features:**
- Route groups: `(auth)`, `(dashboard)`
- Layouts for nested UI
- Loading/error states
- Middleware for auth

### Mobile - Expo + Expo Router

**Why:**
- Fast development with Expo
- File-based routing (like Next.js)
- EAS Build for native builds
- Cross-platform (iOS/Android/Web)

**Key Features:**
- Tab navigation
- Stack navigation
- Deep linking
- Over-the-air updates

### UI - Tailwind CSS + NativeWind

**Why:**
- Consistent styling syntax
- Works on web and mobile
- Utility-first approach
- Small bundle size

**Platform Variants:**
```typescript
// Button.tsx - shared logic
// Button.web.tsx - web implementation
// Button.native.tsx - mobile implementation
```

### State - TanStack Query + Zustand

**Why TanStack Query:**
- Server state management
- Automatic caching
- Background refetching
- Optimistic updates

**Why Zustand:**
- Client state management
- Simple API
- No boilerplate
- DevTools support

**Division:**
- TanStack Query: API data (jobs, users, proposals)
- Zustand: UI state (filters, modals, preferences)

### Auth - Pluggable Adapters

**Why:**
- Environment-specific auth
- Easy to swap providers
- Testable with mocks
- Production-ready Keycloak

**Adapters:**
```typescript
interface AuthAdapter {
  login(email, password): Promise<User>;
  logout(): Promise<void>;
  getUser(): Promise<User | null>;
  getAccessToken(): Promise<string>;
}

// Dev: In-memory credentials
// Prod: Keycloak OAuth2
```

## Package Architecture

### @skillsier/ui

Shared component library:
```
ui/
├── tokens/        # Design tokens
├── components/
│   ├── Button/
│   │   ├── Button.tsx        # Shared logic
│   │   ├── Button.web.tsx    # Web impl
│   │   └── Button.native.tsx # Mobile impl
```

**Exports:**
```typescript
import { Button, Input, Card } from '@skillsier/ui';
```

### @skillsier/types

TypeScript definitions:
```
types/
├── domains/     # Business logic types
├── models/      # API DTOs
├── enums/       # Enums
└── mappers/     # DTO → Domain transformers
```

**Pattern:**
```typescript
// API returns snake_case
interface JobModel {
  job_id: string;
  created_at: string;
}

// Frontend uses camelCase
interface Job {
  id: string;
  createdAt: Date;
}

// Mapper transforms
const toJob = (model: JobModel): Job => ({
  id: model.job_id,
  createdAt: new Date(model.created_at),
});
```

### @skillsier/api

HTTP clients:
```
api/
├── lib/
│   ├── base-client.ts    # Axios instance
│   ├── interceptors.ts   # JWT, errors
├── clients/
│   ├── jobs-client.ts    # Jobs API
│   ├── users-client.ts   # Users API
└── mocks/
    └── jobs.mock.ts      # Mock data
```

**Usage:**
```typescript
import { jobsClient } from '@skillsier/api';

const jobs = await jobsClient.getJobs({ status: 'ACTIVE' });
```

### @skillsier/hooks

React hooks with TanStack Query:
```
hooks/
├── jobs/
│   ├── useJobs.ts        # List
│   ├── useJob.ts         # Detail
│   ├── useCreateJob.ts   # Mutation
```

**Pattern:**
```typescript
export const useJobs = (filters: JobFilters) => {
  return useQuery({
    queryKey: ['jobs', filters],
    queryFn: () => jobsClient.getJobs(filters),
    staleTime: 5 * 60 * 1000, // 5 min
  });
};
```

### @skillsier/stores

Zustand stores:
```
stores/
├── jobs.store.ts     # Job filters
├── auth.store.ts     # Auth state
└── ui.store.ts       # UI state
```

**Pattern:**
```typescript
export const useJobsStore = create<JobsState>((set) => ({
  filters: {},
  setFilters: (filters) => set({ filters }),
  resetFilters: () => set({ filters: {} }),
}));
```

### @skillsier/i18n

Internationalization:
```
i18n/
├── locales/
│   ├── en/
│   └── ar/
└── hooks/
    └── useTranslation.ts
```

**Usage:**
```typescript
const { t, locale, setLocale } = useTranslation();
t('common.welcome')  // "Welcome"
```

## Data Flow

### Query Flow (Read)
```
Component
  ↓ useJobs()
TanStack Query
  ↓ queryFn
API Client
  ↓ HTTP GET
Backend
  ↓ Response
Mapper (model → domain)
  ↓
TanStack Query Cache
  ↓
Component Re-render
```

### Mutation Flow (Write)
```
Component
  ↓ useCreateJob()
TanStack Query Mutation
  ↓ mutationFn
API Client
  ↓ HTTP POST
Backend
  ↓ Response
Mapper
  ↓
Query Cache Invalidation
  ↓
Automatic Refetch
  ↓
Component Update
```

## Authentication Flow

### Web (Next.js)
```
1. User visits /dashboard
2. Middleware checks auth
3. No auth → redirect to /login
4. Login → AuthProvider.login()
5. Set cookies/session
6. Redirect to /dashboard
7. Middleware validates session
8. Render protected page
```

### Mobile (Expo)
```
1. App starts
2. Check stored token
3. No token → show login
4. Login → AuthProvider.login()
5. Store token in SecureStore
6. Navigate to tabs
7. API calls use token
```

## Routing Architecture

### Web Routes
```
app/
├── (auth)/
│   └── login/
│       └── page.tsx
├── (authenticated)/
│   └── (dashboard)/
│       ├── page.tsx
│       └── jobs/
│           ├── page.tsx
│           ├── create/
│           └── [id]/
└── layout.tsx
```

**Route Groups:**
- `(auth)` - Public routes
- `(authenticated)` - Protected routes
- `(dashboard)` - Dashboard layout

### Mobile Routes
```
app/
├── (auth)/
│   └── login.tsx
├── (authenticated)/
│   └── (tabs)/
│       ├── index.tsx
│       └── jobs/
│           ├── index.tsx
│           ├── create.tsx
│           └── [id].tsx
└── _layout.tsx
```

## Performance Optimizations

### Code Splitting

- Next.js automatic code splitting
- Expo lazy loading with `React.lazy()`
- Dynamic imports for large components

### Caching Strategy

**TanStack Query:**
- `staleTime: 5 minutes` - Data fresh for 5 min
- `cacheTime: 30 minutes` - Keep in cache 30 min
- Background refetch on window focus

**API Responses:**
- ETags for conditional requests
- Compression (gzip/brotli)

### Image Optimization

- Next.js Image component
- Expo Image with caching
- WebP format with fallbacks

## Error Handling

### API Errors
```typescript
try {
  const job = await jobsClient.getJob(id);
} catch (error) {
  if (error.status === 404) {
    // Handle not found
  } else if (error.status === 401) {
    // Handle unauthorized
  } else {
    // Handle generic error
  }
}
```

### Error Boundaries
```typescript
// app/error.tsx (Next.js)
'use client';

export default function Error({ error, reset }) {
  return (
    <div>
      <h2>Something went wrong!</h2>
      <button onClick={reset}>Try again</button>
    </div>
  );
}
```

## Testing Strategy

### Unit Tests (Vitest)

- Component logic
- Hooks behavior
- Utility functions
- Mappers

### Integration Tests (Vitest + Testing Library)

- Component + hooks
- Form submissions
- User interactions

### E2E Tests (Playwright)

- Critical user journeys
- Authentication flows
- Job creation/editing

### Coverage Goals

- Utilities: 90%+
- Components: 70%+
- E2E: Critical paths

## Security Considerations

### Authentication

- JWT tokens in HTTP-only cookies (web)
- Secure token storage (mobile)
- Auto token refresh
- CSRF protection

### Authorization

- Role-based access control
- Route-level guards
- Component-level permissions

### Data Validation

- Zod schemas for forms
- API response validation
- Type safety with TypeScript

## Deployment

### Web (Vercel)
```
main branch → Production
develop branch → Staging
PR → Preview deployment
```

### Mobile (EAS)
```
eas build --platform android --profile production
eas build --platform ios --profile production
eas submit
```

## Monitoring

### Error Tracking

- Sentry for error monitoring
- Source maps uploaded
- User context attached

### Analytics

- PostHog for product analytics
- Custom events tracking
- Funnel analysis

### Performance

- Web Vitals monitoring
- Core Web Vitals tracking
- Mobile performance metrics

## Future Enhancements

### Planned

- [ ] Offline support (mobile)
- [ ] Push notifications
- [ ] Advanced search filters
- [ ] Real-time updates (WebSocket)
- [ ] File uploads (multi-part)

### Under Consideration

- [ ] Server-Side Rendering (SSR)
- [ ] GraphQL migration
- [ ] Micro-frontends
- [ ] Native mobile (React Native CLI)

## References

- [Next.js Documentation](https://nextjs.org/docs)
- [Expo Documentation](https://docs.expo.dev/)
- [TanStack Query](https://tanstack.com/query)
- [Zustand](https://github.com/pmndrs/zustand)
- [Tailwind CSS](https://tailwindcss.com/)