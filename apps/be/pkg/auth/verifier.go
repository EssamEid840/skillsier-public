package auth

import "context"

// TokenVerifier is the domain interface for token verification
// Services depend only on this interface, not on specific implementations (e.g., Keycloak)
type TokenVerifier interface {
	// VerifyToken validates a JWT token and returns the principal
	// Returns ErrInvalidToken, ErrExpiredToken, ErrUnauthorized, etc. on failure
	VerifyToken(ctx context.Context, token string) (*Principal, error)
}