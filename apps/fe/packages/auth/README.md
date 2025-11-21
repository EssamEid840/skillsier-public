# @skillsier/auth

Authentication package for Skillsier platform with pluggable adapter architecture.

## Features

- Pluggable auth adapters (Dev, Keycloak)
- React hooks and context
- Token management
- Type-safe authentication flow

## Usage

### Setup Auth Provider

```tsx
import { AuthProvider, createAuthAdapter } from '@skillsier/auth';

const authAdapter = createAuthAdapter(
  process.env.AUTH_PROVIDER as 'dev' | 'keycloak'
);

function App() {
  return (
    <AuthProvider adapter={authAdapter}>
      {/* Your app */}
    </AuthProvider>
  );
}
```

### Use Auth Hook

```tsx
import { useAuth } from '@skillsier/auth';

function LoginPage() {
  const { login, isLoading, error, user, isAuthenticated } = useAuth();

  const handleLogin = async (email: string, password: string) => {
    await login({ email, password });
  };

  return (
    // Your login UI
  );
}
```

## Auth Adapters

### Dev Adapter (Default)

Development-only adapter with seeded accounts:
- `admin@skillsier.dev` / `admin123`
- `client@skillsier.dev` / `client123`
- `freelancer@skillsier.dev` / `freelancer123`

### Keycloak Adapter

Production-ready Keycloak integration (stub - needs implementation).

To enable Keycloak:

1. Set `AUTH_PROVIDER=keycloak` in `.env`
2. Configure Keycloak environment variables:
   ```env
   KEYCLOAK_ISSUER=https://keycloak.example.com/realms/skillsier
   KEYCLOAK_CLIENT_ID=skillsier-web
   KEYCLOAK_CLIENT_SECRET=your-secret
   ```
3. Implement methods in `src/adapters/keycloak-adapter.ts`

## API

### `useAuth()`

Returns:
- `user: AuthUser | null` - Current authenticated user
- `isAuthenticated: boolean` - Authentication status
- `isLoading: boolean` - Loading state
- `error: string | null` - Error message
- `login(credentials)` - Login function
- `signup(credentials)` - Signup function
- `logout()` - Logout function
- `refreshAuth()` - Refresh authentication state

### `createAuthAdapter(provider)`

Creates auth adapter instance.

Arguments:
- `provider: 'dev' | 'keycloak'` - Auth provider to use

## Development

```bash
pnpm install
pnpm build
pnpm test
```