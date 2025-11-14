# pkg/auth - Centralized Authentication Package

Centralized Keycloak authentication for all Skillsier microservices. This package provides JWT token verification, RBAC (Role-Based Access Control), and permission management.

## Features

- ✅ JWT token verification using JWKS
- ✅ Keycloak integration with automatic JWKS refresh
- ✅ Role-based access control (RBAC)
- ✅ Fine-grained permissions
- ✅ Context-based principal management
- ✅ Keycloak Admin API client
- ✅ Extensible role-to-permission mapping

## Architecture

```
pkg/auth/
├── config.go           # Base auth configuration
├── errors.go           # Auth-specific errors
├── principal.go        # Normalized identity model
├── context.go          # Context helpers for principal
├── verifier.go         # TokenVerifier interface
├── authorizer.go       # RBAC middleware
└── keycloak/
    ├── config.go       # Keycloak-specific config
    ├── verifier.go     # JWT/JWKS implementation
    ├── client.go       # Admin API client
    └── role_mapper.go  # Role → Permission mapping
```

## Usage

### 1. Initialize Keycloak Verifier

```go
import (
    "skillsier.dev/pkg/auth/keycloak"
)

// Create config
config := keycloak.NewConfig(
    "https://keycloak.skillsier.com",  // Base URL
    "skillsier",                         // Realm
    "users-be",                          // Client ID
    "your-client-secret",                // Client Secret
)

// Create verifier
verifier, err := keycloak.NewVerifier(config)
if err != nil {
    log.Fatal(err)
}
defer verifier.Close()
```

### 2. Verify Tokens

```go
import "skillsier.dev/pkg/auth"

// Extract token from request header
token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

// Verify and extract principal
principal, err := verifier.VerifyToken(r.Context(), token)
if err != nil {
    // Handle auth errors
    if errors.Is(err, auth.ErrExpiredToken) {
        // Token expired
    } else if errors.Is(err, auth.ErrInvalidToken) {
        // Invalid token
    }
    return
}

// Add principal to context
ctx := auth.WithPrincipal(r.Context(), principal)
```

### 3. Use RBAC Middleware

```go
import "skillsier.dev/pkg/auth"

authorizer := auth.NewAuthorizer(verifier)

// Require any of the specified roles
requireFreelancer := authorizer.RequireRoles("freelancer", "premium_freelancer")

// Require all specified roles
requireAdminMod := authorizer.RequireAllRoles("admin", "moderator")

// Require permissions
requireJobCreate := authorizer.RequirePermissions("jobs:create")

// Use in middleware
if err := requireFreelancer(ctx); err != nil {
    // User doesn't have required role
    return err
}
```

### 4. Access Principal in Handlers

```go
import "skillsier.dev/pkg/auth"

func MyHandler(w http.ResponseWriter, r *http.Request) {
    // Get principal from context
    principal := auth.GetPrincipal(r.Context())
    if principal == nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Use principal
    userID := principal.UserID
    username := principal.Username
    email := principal.Email

    // Check roles
    if principal.HasRole("admin") {
        // Admin logic
    }

    // Check permissions
    if principal.HasPermission("jobs:create") {
        // Can create jobs
    }
}
```

### 5. Custom Role Mappings

```go
import "skillsier.dev/pkg/auth/keycloak"

mapper := keycloak.NewRoleMapper()

// Add custom role mapping
mapper.AddRoleMapping("custom_role", []string{
    "jobs:create",
    "proposals:submit",
    "contracts:view",
})

// Use in verifier (modify verifier.go to accept custom mapper)
```

## Role Definitions

### Admin Roles
- **admin**: Full access (wildcard permission)
- **super_admin**: Complete platform control
- **moderator**: Content moderation, user management
- **support**: Support tickets, dispute viewing

### Freelancer Roles
- **freelancer**: Basic freelancer permissions
- **premium_freelancer**: Enhanced features + analytics
- **verified_freelancer**: Trust badge

### Client Roles
- **client**: Basic client permissions
- **premium_client**: Advanced features + analytics
- **verified_client**: Trust badge
- **agency**: Team management + bulk operations
- **enterprise**: Full API access + custom integrations

## Permission System

Permissions follow the format: `resource:action`

Examples:
- `jobs:create` - Can create jobs
- `proposals:submit` - Can submit proposals
- `contracts:manage` - Can manage contracts
- `users:suspend` - Can suspend users
- `*` - Wildcard (all permissions)

## Error Handling

```go
import "skillsier.dev/pkg/auth"

principal, err := verifier.VerifyToken(ctx, token)
if err != nil {
    switch {
    case errors.Is(err, auth.ErrUnauthorized):
        // No valid credentials
    case errors.Is(err, auth.ErrForbidden):
        // Lacks permissions
    case errors.Is(err, auth.ErrInvalidToken):
        // Malformed token
    case errors.Is(err, auth.ErrExpiredToken):
        // Token expired
    case errors.Is(err, auth.ErrInvalidIssuer):
        // Wrong issuer
    case errors.Is(err, auth.ErrInvalidAudience):
        // Wrong audience
    default:
        // Other error
    }
}
```

## Configuration

### Environment Variables

```bash
# Keycloak Configuration
KEYCLOAK_URL=https://keycloak.skillsier.com
KEYCLOAK_REALM=skillsier
KEYCLOAK_CLIENT_ID=users-be
KEYCLOAK_CLIENT_SECRET=your-secret

# Auth Configuration (optional, has defaults)
AUTH_CACHE_TTL=10m
AUTH_CLOCK_SKEW=60s
AUTH_ALLOWED_ALGORITHMS=RS256
```

### Code Configuration

```go
config := &keycloak.Config{
    Auth: &auth.Config{
        Issuer:            "https://keycloak.skillsier.com/realms/skillsier",
        Audience:          "users-be",
        JWKSURL:           "https://keycloak.skillsier.com/realms/skillsier/protocol/openid-connect/certs",
        AllowedAlgorithms: []string{"RS256"},
        CacheTTL:          10 * time.Minute,
        ClockSkew:         60 * time.Second,
    },
    Realm:        "skillsier",
    ClientID:     "users-be",
    ClientSecret: "your-secret",
    BaseURL:      "https://keycloak.skillsier.com",
}
```

## Testing

```go
import (
    "testing"
    "skillsier.dev/pkg/auth"
)

func TestPrincipalPermissions(t *testing.T) {
    principal := &auth.Principal{
        UserID:      "123",
        Subject:     "sub-123",
        Roles:       []string{"freelancer"},
        Permissions: []string{"jobs:view", "proposals:submit"},
    }

    if !principal.HasPermission("jobs:view") {
        t.Error("Expected permission jobs:view")
    }

    if principal.HasPermission("jobs:delete") {
        t.Error("Should not have permission jobs:delete")
    }
}
```

## Integration with Services

Each service should import this package:

```go
import "skillsier.dev/pkg/auth"
import "skillsier.dev/pkg/auth/keycloak"
```

**DO NOT**:
- Copy JWT verification code into services
- Implement custom authentication logic
- Directly import Keycloak libraries in services

**DO**:
- Use pkg/auth for all authentication needs
- Depend on the TokenVerifier interface
- Use the provided RBAC middleware

## Best Practices

1. **Always verify tokens on every request** - Don't cache authentication results
2. **Use the most restrictive permission check** - Require specific permissions when possible
3. **Handle all error types** - Different errors need different responses
4. **Add principal to context early** - Do it in middleware
5. **Use permission checks over role checks** - More flexible
6. **Keep role mappings updated** - As new features are added

## Troubleshooting

### JWKS Refresh Failures
- Check network connectivity to Keycloak
- Verify JWKS URL is correct
- Check if Keycloak is accessible

### Token Verification Failures
- Verify issuer and audience match Keycloak configuration
- Check token hasn't expired
- Ensure Keycloak is using RS256 algorithm

### Permission Denied
- Check role mappings in role_mapper.go
- Verify user has correct roles in Keycloak
- Ensure permissions are correctly defined

## Dependencies

- `github.com/MicahParks/keyfunc/v2` - JWKS handling
- `github.com/golang-jwt/jwt/v5` - JWT parsing

## License

Proprietary - Skillsier Platform