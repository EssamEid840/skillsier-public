package keycloak

import "strings"

// RoleMapper maps Keycloak roles to application-specific permissions
// This provides a layer of abstraction between Keycloak roles and fine-grained permissions
type RoleMapper struct {
	// roleToPermissions maps role names to permission lists
	roleToPermissions map[string][]string
}

// NewRoleMapper creates a new RoleMapper with default mappings for Skillsier platform
func NewRoleMapper() *RoleMapper {
	return &RoleMapper{
		roleToPermissions: map[string][]string{
			// Admin roles
			"admin": {
				"*", // Wildcard - admin can do everything
			},
			"super_admin": {
				"*", // Super admin has all permissions
			},
			"moderator": {
				"users:suspend", "users:ban", "users:warn", "users:verify",
				"jobs:remove", "jobs:feature", "jobs:flag:review",
				"proposals:remove", "proposals:flag:review",
				"contracts:view_all", "contracts:dispute:resolve",
				"messages:remove", "messages:flag:review",
				"reviews:remove", "reviews:flag:review",
				"content:moderate",
			},
			"support": {
				"users:view", "users:search",
				"tickets:view", "tickets:respond", "tickets:assign",
				"disputes:view", "disputes:comment",
				"reports:view",
			},

			// Freelancer roles
			"freelancer": {
				"profile:create", "profile:update", "profile:view",
				"proposals:create", "proposals:update", "proposals:withdraw",
				"contracts:accept", "contracts:view", "contracts:work",
				"messages:send", "messages:read",
				"reviews:respond",
				"wallet:view", "wallet:withdraw",
			},
			"premium_freelancer": {
				// Inherits all freelancer permissions plus:
				"profile:create", "profile:update", "profile:view",
				"proposals:create", "proposals:update", "proposals:withdraw",
				"proposals:boost", "proposals:featured",
				"contracts:accept", "contracts:view", "contracts:work",
				"messages:send", "messages:read",
				"reviews:respond",
				"wallet:view", "wallet:withdraw",
				"jobs:advanced_search", "jobs:saved_searches",
				"analytics:view",
			},
			"verified_freelancer": {
				// Verified badge - same as freelancer but with trust signals
				"profile:create", "profile:update", "profile:view",
				"proposals:create", "proposals:update", "proposals:withdraw",
				"contracts:accept", "contracts:view", "contracts:work",
				"messages:send", "messages:read",
				"reviews:respond",
				"wallet:view", "wallet:withdraw",
			},

			// Client roles
			"client": {
				"jobs:create", "jobs:update", "jobs:close",
				"proposals:view", "proposals:accept", "proposals:reject",
				"contracts:create", "contracts:view", "contracts:manage",
				"messages:send", "messages:read",
				"reviews:create",
				"payments:add_funds", "payments:process",
				"wallet:view",
			},
			"premium_client": {
				// Inherits all client permissions plus:
				"jobs:create", "jobs:update", "jobs:close",
				"jobs:feature", "jobs:priority",
				"proposals:view", "proposals:accept", "proposals:reject",
				"contracts:create", "contracts:view", "contracts:manage",
				"messages:send", "messages:read",
				"reviews:create",
				"payments:add_funds", "payments:process",
				"wallet:view",
				"analytics:view", "reports:generate",
			},
			"verified_client": {
				// Verified badge - same as client
				"jobs:create", "jobs:update", "jobs:close",
				"proposals:view", "proposals:accept", "proposals:reject",
				"contracts:create", "contracts:view", "contracts:manage",
				"messages:send", "messages:read",
				"reviews:create",
				"payments:add_funds", "payments:process",
				"wallet:view",
			},

			// Special roles
			"agency": {
				"agency:create", "agency:manage",
				"jobs:create", "jobs:update", "jobs:close", "jobs:bulk",
				"proposals:view", "proposals:accept", "proposals:reject",
				"contracts:create", "contracts:view", "contracts:manage", "contracts:assign",
				"team:invite", "team:manage",
				"messages:send", "messages:read",
				"payments:add_funds", "payments:process",
				"wallet:view",
				"reports:generate", "analytics:view",
			},
			"enterprise": {
				// Enterprise client with advanced features
				"jobs:create", "jobs:update", "jobs:close", "jobs:bulk", "jobs:templates",
				"proposals:view", "proposals:accept", "proposals:reject",
				"contracts:create", "contracts:view", "contracts:manage", "contracts:templates",
				"messages:send", "messages:read",
				"reviews:create",
				"payments:add_funds", "payments:process", "payments:invoice",
				"wallet:view",
				"analytics:view", "reports:generate", "reports:scheduled",
				"api:access",
			},
		},
	}
}

// MapRolesToPermissions converts a list of roles to a unique list of permissions
func (m *RoleMapper) MapRolesToPermissions(roles []string) []string {
	permissionSet := make(map[string]bool)

	for _, role := range roles {
		role = strings.ToLower(strings.TrimSpace(role))
		
		// Check if wildcard permission exists
		if perms, ok := m.roleToPermissions[role]; ok {
			for _, perm := range perms {
				if perm == "*" {
					// Wildcard permission - return immediately
					return []string{"*"}
				}
				permissionSet[perm] = true
			}
		}
	}

	// Convert set to slice
	permissions := make([]string, 0, len(permissionSet))
	for perm := range permissionSet {
		permissions = append(permissions, perm)
	}

	return permissions
}

// AddRoleMapping adds a new role-to-permissions mapping
func (m *RoleMapper) AddRoleMapping(role string, permissions []string) {
	role = strings.ToLower(strings.TrimSpace(role))
	m.roleToPermissions[role] = permissions
}

// GetPermissionsForRole returns the permissions for a specific role
func (m *RoleMapper) GetPermissionsForRole(role string) []string {
	role = strings.ToLower(strings.TrimSpace(role))
	if perms, ok := m.roleToPermissions[role]; ok {
		return perms
	}
	return []string{}
}