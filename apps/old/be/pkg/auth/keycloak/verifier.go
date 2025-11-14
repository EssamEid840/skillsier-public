package keycloak

import (
	"context"
	"fmt"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/golang-jwt/jwt/v5"
	"skillsier.dev/pkg/auth"
)

// Verifier implements auth.TokenVerifier for Keycloak JWT tokens
type Verifier struct {
	config *Config
	jwks   *keyfunc.JWKS
}

// NewVerifier creates a new Keycloak token verifier
func NewVerifier(config *Config) (*Verifier, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	// Initialize JWKS (JSON Web Key Set) client
	jwks, err := keyfunc.Get(config.Auth.JWKSURL, keyfunc.Options{
		RefreshInterval:   config.Auth.CacheTTL,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Second * 10,
		RefreshErrorHandler: func(err error) {
			// Log the error but don't fail - use cached keys
			fmt.Printf("JWKS refresh error: %v\n", err)
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize JWKS: %w", err)
	}

	return &Verifier{
		config: config,
		jwks:   jwks,
	}, nil
}

// VerifyToken validates a Keycloak JWT token and returns the principal
func (v *Verifier) VerifyToken(ctx context.Context, tokenString string) (*auth.Principal, error) {
	// Parse and verify the token
	token, err := jwt.Parse(tokenString, v.jwks.Keyfunc, jwt.WithValidMethods(v.config.Auth.AllowedAlgorithms))
	if err != nil {
		if err == jwt.ErrTokenExpired {
			return nil, auth.ErrExpiredToken.WithCause(err)
		}
		return nil, auth.ErrInvalidToken.WithCause(err)
	}

	if !token.Valid {
		return nil, auth.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, auth.ErrInvalidToken.WithCause(fmt.Errorf("invalid claims type"))
	}

	// Verify issuer
	if !claims.VerifyIssuer(v.config.Auth.Issuer, true) {
		return nil, auth.ErrInvalidIssuer
	}

	// Verify audience
	if !claims.VerifyAudience(v.config.Auth.Audience, true) {
		return nil, auth.ErrInvalidAudience
	}

	// Extract principal from claims
	principal, err := v.extractPrincipal(claims)
	if err != nil {
		return nil, auth.ErrInvalidToken.WithCause(err)
	}

	return principal, nil
}

// extractPrincipal extracts Principal from Keycloak JWT claims
func (v *Verifier) extractPrincipal(claims jwt.MapClaims) (*auth.Principal, error) {
	principal := &auth.Principal{
		Metadata: make(map[string]interface{}),
	}

	// Extract subject (required)
	if sub, ok := claims["sub"].(string); ok {
		principal.Subject = sub
		principal.UserID = sub // Use subject as UserID by default
	} else {
		return nil, fmt.Errorf("missing or invalid 'sub' claim")
	}

	// Extract username (preferred_username in Keycloak)
	if username, ok := claims["preferred_username"].(string); ok {
		principal.Username = username
	}

	// Extract email
	if email, ok := claims["email"].(string); ok {
		principal.Email = email
	}

	// Extract email_verified
	if emailVerified, ok := claims["email_verified"].(bool); ok {
		principal.EmailVerified = emailVerified
	}

	// Extract roles from realm_access
	if realmAccess, ok := claims["realm_access"].(map[string]interface{}); ok {
		if roles, ok := realmAccess["roles"].([]interface{}); ok {
			principal.Roles = make([]string, 0, len(roles))
			for _, role := range roles {
				if r, ok := role.(string); ok {
					principal.Roles = append(principal.Roles, r)
				}
			}
		}
	}

	// Extract client roles (resource_access)
	if resourceAccess, ok := claims["resource_access"].(map[string]interface{}); ok {
		if clientAccess, ok := resourceAccess[v.config.ClientID].(map[string]interface{}); ok {
			if roles, ok := clientAccess["roles"].([]interface{}); ok {
				for _, role := range roles {
					if r, ok := role.(string); ok {
						principal.Roles = append(principal.Roles, r)
					}
				}
			}
		}
	}

	// Map roles to permissions using RoleMapper
	mapper := NewRoleMapper()
	principal.Permissions = mapper.MapRolesToPermissions(principal.Roles)

	// Store additional metadata
	for key, value := range claims {
		if key != "sub" && key != "preferred_username" && key != "email" && 
		   key != "email_verified" && key != "realm_access" && key != "resource_access" {
			principal.Metadata[key] = value
		}
	}

	return principal, nil
}

// Close cleans up resources
func (v *Verifier) Close() error {
	if v.jwks != nil {
		v.jwks.EndBackground()
	}
	return nil
}