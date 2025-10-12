package auth

import "context"

type contextKey string

const principalKey contextKey = "auth:principal"

// WithPrincipal adds a Principal to the context
func WithPrincipal(ctx context.Context, principal *Principal) context.Context {
	return context.WithValue(ctx, principalKey, principal)
}

// GetPrincipal retrieves the Principal from the context
// Returns nil if no principal is found
func GetPrincipal(ctx context.Context) *Principal {
	principal, ok := ctx.Value(principalKey).(*Principal)
	if !ok {
		return nil
	}
	return principal
}

// MustGetPrincipal retrieves the Principal from the context
// Panics if no principal is found (use only when you're certain a principal exists)
func MustGetPrincipal(ctx context.Context) *Principal {
	principal := GetPrincipal(ctx)
	if principal == nil {
		panic("auth: no principal in context")
	}
	return principal
}

// GetUserID is a convenience method to get the user ID from context
func GetUserID(ctx context.Context) string {
	principal := GetPrincipal(ctx)
	if principal == nil {
		return ""
	}
	return principal.UserID
}

// GetSubject is a convenience method to get the subject from context
func GetSubject(ctx context.Context) string {
	principal := GetPrincipal(ctx)
	if principal == nil {
		return ""
	}
	return principal.Subject
}