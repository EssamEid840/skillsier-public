package user

import (
	"context"

	"github.com/google/uuid"
)

// Repository defines the interface for user data persistence
// This follows the repository pattern, keeping domain logic independent of infrastructure
type Repository interface {
	// Create creates a new user in the database
	Create(ctx context.Context, user *User) error
	
	// FindByID retrieves a user by their internal ID
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
	
	// FindByKeycloakID retrieves a user by their Keycloak ID (sub claim)
	FindByKeycloakID(ctx context.Context, keycloakID string) (*User, error)
	
	// FindByEmail retrieves a user by their email address
	FindByEmail(ctx context.Context, email string) (*User, error)
	
	// FindByUsername retrieves a user by their username
	FindByUsername(ctx context.Context, username string) (*User, error)
	
	// Update updates an existing user's information
	Update(ctx context.Context, user *User) error
	
	// Delete soft-deletes a user (sets DeletedAt timestamp)
	Delete(ctx context.Context, id uuid.UUID) error
	
	// List retrieves a paginated list of users
	List(ctx context.Context, limit, offset int) ([]*User, int64, error)
	
	// ExistsByKeycloakID checks if a user with the given Keycloak ID exists
	ExistsByKeycloakID(ctx context.Context, keycloakID string) (bool, error)
}