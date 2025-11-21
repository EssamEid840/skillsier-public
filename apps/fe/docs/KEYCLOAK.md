# Keycloak Authentication Setup

Guide for configuring production Keycloak authentication.

## Overview

Skillsier uses a pluggable authentication system with two adapters:
- **Dev Adapter:** In-memory credentials for development
- **Keycloak Adapter:** OAuth2/OIDC for production

## Architecture
```
┌─────────────┐
│   Web App   │
│   Mobile    │
└──────┬──────┘
       │
       ▼
┌──────────────────┐
│  AuthProvider    │
│  (packages/auth) │
└──────┬───────────┘
       │
       ├─── Dev Adapter (development)
       └─── Keycloak Adapter (production)
                │
                ▼
         ┌──────────────┐
         │   Keycloak   │
         │   Server     │
         └──────────────┘
```

## Keycloak Server Setup

### 1. Install Keycloak
```bash
# Docker (recommended)
docker run -d \
  --name keycloak \
  -p 8080:8080 \
  -e KEYCLOAK_ADMIN=admin \
  -e KEYCLOAK_ADMIN_PASSWORD=admin \
  quay.io/keycloak/keycloak:23.0.0 \
  start-dev
```

### 2. Create Realm

1. Login to admin console: http://localhost:8080
2. Click "Create Realm"
3. Name: `skillsier`
4. Click "Create"

### 3. Create Client

1. Navigate to Clients → Create Client
2. **Client ID:** `skillsier-web` (for web) or `skillsier-mobile` (for mobile)
3. **Client Protocol:** openid-connect
4. Click "Next"

#### Client Settings:

**Web Client:**
```
Client Authentication: ON
Authorization: OFF
Standard Flow: ON
Direct Access Grants: ON
Valid Redirect URIs: http://localhost:3000/api/auth/callback/keycloak
Web Origins: http://localhost:3000
```

**Mobile Client:**
```
Client Authentication: OFF (public client)
Authorization: OFF
Standard Flow: ON
Direct Access Grants: ON
Valid Redirect URIs: skillsier://auth
```

### 4. Create Users

1. Navigate to Users → Add User
2. Create test accounts:
   - `admin@skillsier.com`
   - `client@skillsier.com`
   - `freelancer@skillsier.com`
3. Set passwords in Credentials tab

### 5. Create Roles

1. Navigate to Realm Roles → Create Role
2. Create roles:
   - `admin`
   - `client`
   - `freelancer`
3. Assign roles to users in Users → Role Mapping

## Frontend Configuration

### 1. Environment Variables

Update `.env`:
```env
# Switch to Keycloak
AUTH_PROVIDER=keycloak

# Keycloak Configuration
KEYCLOAK_ISSUER=http://localhost:8080/realms/skillsier
KEYCLOAK_CLIENT_ID=skillsier-web
KEYCLOAK_CLIENT_SECRET=your-client-secret-here

# NextAuth Configuration
NEXTAUTH_SECRET=generate-with-openssl-rand-base64-32
NEXTAUTH_URL=http://localhost:3000
```

### 2. Get Client Secret

1. In Keycloak Admin Console
2. Navigate to Clients → skillsier-web
3. Go to Credentials tab
4. Copy Client Secret
5. Add to `.env` as `KEYCLOAK_CLIENT_SECRET`

### 3. Mobile Configuration

Update `apps/mobile/.env`:
```env
AUTH_PROVIDER=keycloak
KEYCLOAK_ISSUER=http://localhost:8080/realms/skillsier
KEYCLOAK_CLIENT_ID=skillsier-mobile
```

## Testing Keycloak Integration

### Web
```bash
# Start web app
pnpm dev:web

# Navigate to http://localhost:3000
# Click "Sign In"
# You'll be redirected to Keycloak login
# Login with test account
# Redirected back to app
```

### Mobile
```bash
# Start mobile app
pnpm dev:mobile

# Tap "Sign In"
# Opens Keycloak in browser
# Login with test account
# Redirected back to app
```

## Token Management

### Access Token
```typescript
import { useAuth } from '@skillsier/auth';

function MyComponent() {
  const { getAccessToken } = useAuth();
  
  const fetchData = async () => {
    const token = await getAccessToken();
    const response = await fetch('/api/data', {
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
  };
}
```

### Token Refresh

Automatic token refresh is handled by the auth adapter:
- Access tokens expire in 5 minutes
- Refresh tokens expire in 30 days
- Auto-refresh happens 1 minute before expiry

## Production Deployment

### 1. Keycloak Server

Deploy Keycloak to production:
```bash
# Production URL
KEYCLOAK_ISSUER=https://auth.skillsier.com/realms/skillsier
```

### 2. Update Environment

Production `.env`:
```env
AUTH_PROVIDER=keycloak
KEYCLOAK_ISSUER=https://auth.skillsier.com/realms/skillsier
KEYCLOAK_CLIENT_ID=skillsier-web
KEYCLOAK_CLIENT_SECRET=prod-secret
NEXTAUTH_SECRET=prod-secret-min-32-chars
NEXTAUTH_URL=https://app.skillsier.com
```

### 3. Valid Redirect URIs

Update Keycloak client:
```
Web: https://app.skillsier.com/api/auth/callback/keycloak
Mobile: skillsier://auth
```

## Troubleshooting

### "Invalid redirect URI"

- Check KEYCLOAK_ISSUER matches realm URL exactly
- Verify Valid Redirect URIs in Keycloak client
- Ensure NEXTAUTH_URL matches callback domain

### "Invalid client credentials"

- Verify KEYCLOAK_CLIENT_ID matches client name
- Check KEYCLOAK_CLIENT_SECRET is correct
- Ensure Client Authentication is enabled

### Token errors
```bash
# Check token validity
curl -X POST \
  http://localhost:8080/realms/skillsier/protocol/openid-connect/token/introspect \
  -d "client_id=skillsier-web" \
  -d "client_secret=your-secret" \
  -d "token=your-token"
```

### Mobile redirect not working

- Verify scheme `skillsier://` is registered in app.json
- Check Keycloak client Valid Redirect URIs includes `skillsier://auth`
- Test redirect with: `skillsier://auth?code=test`

## Development vs Production

| Feature | Development | Production |
|---------|------------|------------|
| Provider | Dev Adapter | Keycloak |
| Auth Flow | In-memory | OAuth2 OIDC |
| Tokens | Fake JWT | Real JWT |
| Users | Seeded | Keycloak DB |
| MFA | ❌ | ✅ |
| SSO | ❌ | ✅ |

## Additional Resources

- [Keycloak Documentation](https://www.keycloak.org/documentation)
- [NextAuth.js Keycloak Provider](https://next-auth.js.org/providers/keycloak)
- [OAuth 2.0 Flow](https://oauth.net/2/)