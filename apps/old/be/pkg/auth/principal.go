package auth

// Principal represents the authenticated identity (normalized across all auth providers)
type Principal struct {
	// UserID is the internal application user ID
	UserID string
	
	// Subject is the unique identifier from the auth provider (e.g., Keycloak sub claim)
	Subject string
	
	// Username is the human-readable username
	Username string
	
	// Email is the user's email address
	Email string
	
	// EmailVerified indicates if the email has been verified
	EmailVerified bool
	
	// Roles contains the user's role names (e.g., ["freelancer", "premium"])
	Roles []string
	
	// Permissions contains fine-grained permissions (e.g., ["jobs:create", "proposals:submit"])
	Permissions []string
	
	// Metadata contains additional claims from the token
	Metadata map[string]interface{}
}

// HasRole checks if the principal has a specific role
func (p *Principal) HasRole(role string) bool {
	for _, r := range p.Roles {
		if r == role {
			return true
		}
	}
	return false
}

// HasAnyRole checks if the principal has any of the specified roles
func (p *Principal) HasAnyRole(roles ...string) bool {
	for _, role := range roles {
		if p.HasRole(role) {
			return true
		}
	}
	return false
}

// HasAllRoles checks if the principal has all of the specified roles
func (p *Principal) HasAllRoles(roles ...string) bool {
	for _, role := range roles {
		if !p.HasRole(role) {
			return false
		}
	}
	return true
}

// HasPermission checks if the principal has a specific permission
func (p *Principal) HasPermission(permission string) bool {
	for _, p := range p.Permissions {
		if p == permission {
			return true
		}
	}
	return false
}

// HasAnyPermission checks if the principal has any of the specified permissions
func (p *Principal) HasAnyPermission(permissions ...string) bool {
	for _, permission := range permissions {
		if p.HasPermission(permission) {
			return true
		}
	}
	return false
}

// HasAllPermissions checks if the principal has all of the specified permissions
func (p *Principal) HasAllPermissions(permissions ...string) bool {
	for _, permission := range permissions {
		if !p.HasPermission(permission) {
			return false
		}
	}
	return true
}

// IsAuthenticated returns true if the principal represents an authenticated user
func (p *Principal) IsAuthenticated() bool {
	return p.Subject != ""
}