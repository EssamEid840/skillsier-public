package auth

import (
	"context"
	"fmt"
	"strings"
)

// Authorizer provides RBAC (Role-Based Access Control) functionality
type Authorizer struct {
	verifier TokenVerifier
}

// NewAuthorizer creates a new Authorizer
func NewAuthorizer(verifier TokenVerifier) *Authorizer {
	return &Authorizer{
		verifier: verifier,
	}
}

// RequireRoles returns a middleware function that checks if the authenticated user has ANY of the required roles
func (a *Authorizer) RequireRoles(roles ...string) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		principal := GetPrincipal(ctx)
		if principal == nil {
			return ErrUnauthorized
		}
		
		if !principal.HasAnyRole(roles...) {
			return ErrForbidden.WithCause(
				fmt.Errorf("required roles: %s, user roles: %s",
					strings.Join(roles, ", "),
					strings.Join(principal.Roles, ", "),
				),
			)
		}
		
		return nil
	}
}

// RequireAllRoles returns a middleware function that checks if the authenticated user has ALL of the required roles
func (a *Authorizer) RequireAllRoles(roles ...string) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		principal := GetPrincipal(ctx)
		if principal == nil {
			return ErrUnauthorized
		}
		
		if !principal.HasAllRoles(roles...) {
			return ErrForbidden.WithCause(
				fmt.Errorf("required all roles: %s, user roles: %s",
					strings.Join(roles, ", "),
					strings.Join(principal.Roles, ", "),
				),
			)
		}
		
		return nil
	}
}

// RequirePermissions returns a middleware function that checks if the authenticated user has ANY of the required permissions
func (a *Authorizer) RequirePermissions(permissions ...string) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		principal := GetPrincipal(ctx)
		if principal == nil {
			return ErrUnauthorized
		}
		
		if !principal.HasAnyPermission(permissions...) {
			return ErrForbidden.WithCause(
				fmt.Errorf("required permissions: %s, user permissions: %s",
					strings.Join(permissions, ", "),
					strings.Join(principal.Permissions, ", "),
				),
			)
		}
		
		return nil
	}
}

// RequireAllPermissions returns a middleware function that checks if the authenticated user has ALL of the required permissions
func (a *Authorizer) RequireAllPermissions(permissions ...string) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		principal := GetPrincipal(ctx)
		if principal == nil {
			return ErrUnauthorized
		}
		
		if !principal.HasAllPermissions(permissions...) {
			return ErrForbidden.WithCause(
				fmt.Errorf("required all permissions: %s, user permissions: %s",
					strings.Join(permissions, ", "),
					strings.Join(principal.Permissions, ", "),
				),
			)
		}
		
		return nil
	}
}

// RequireAuthentication ensures a valid principal exists in the context
func (a *Authorizer) RequireAuthentication() func(ctx context.Context) error {
	return func(ctx context.Context) error {
		principal := GetPrincipal(ctx)
		if principal == nil || !principal.IsAuthenticated() {
			return ErrUnauthorized
		}
		return nil
	}
}