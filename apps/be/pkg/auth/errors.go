package auth

import "fmt"

// Error types for authentication and authorization
var (
	// ErrUnauthorized indicates the request lacks valid authentication credentials
	ErrUnauthorized = NewAuthError("unauthorized", "authentication required")
	
	// ErrForbidden indicates the authenticated user lacks required permissions
	ErrForbidden = NewAuthError("forbidden", "insufficient permissions")
	
	// ErrInvalidToken indicates the token is malformed or invalid
	ErrInvalidToken = NewAuthError("invalid_token", "token is invalid or malformed")
	
	// ErrExpiredToken indicates the token has expired
	ErrExpiredToken = NewAuthError("expired_token", "token has expired")
	
	// ErrInvalidIssuer indicates the token issuer doesn't match expected issuer
	ErrInvalidIssuer = NewAuthError("invalid_issuer", "token issuer is invalid")
	
	// ErrInvalidAudience indicates the token audience doesn't match expected audience
	ErrInvalidAudience = NewAuthError("invalid_audience", "token audience is invalid")
)

// AuthError represents an authentication/authorization error
type AuthError struct {
	Code    string
	Message string
	Cause   error
}

// NewAuthError creates a new authentication error
func NewAuthError(code, message string) *AuthError {
	return &AuthError{
		Code:    code,
		Message: message,
	}
}

// Error implements the error interface
func (e *AuthError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s (cause: %v)", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// WithCause adds a cause to the error
func (e *AuthError) WithCause(cause error) *AuthError {
	e.Cause = cause
	return e
}

// Is checks if the error matches the target error
func (e *AuthError) Is(target error) bool {
	t, ok := target.(*AuthError)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

// ErrInvalidConfig returns a configuration error
func ErrInvalidConfig(message string) error {
	return NewAuthError("invalid_config", message)
}