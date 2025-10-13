package postgres

import (
	"context"
	"errors"
	"fmt"

	"users-be/internal/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// userRepository implements the user.Repository interface using PostgreSQL
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of the user repository
func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{db: db}
}

// Create creates a new user in the database
func (r *userRepository) Create(ctx context.Context, u *user.User) error {
	if err := r.db.WithContext(ctx).Create(u).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// FindByID retrieves a user by their internal ID
func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).
		Where("id = ? AND deleted_at IS NULL", id).
		First(&u).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to find user by ID: %w", err)
	}
	
	return &u, nil
}

// FindByKeycloakID retrieves a user by their Keycloak ID
func (r *userRepository) FindByKeycloakID(ctx context.Context, keycloakID string) (*user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).
		Where("keycloak_id = ? AND deleted_at IS NULL", keycloakID).
		First(&u).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to find user by Keycloak ID: %w", err)
	}
	
	return &u, nil
}

// FindByEmail retrieves a user by their email address
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).
		Where("email = ? AND deleted_at IS NULL", email).
		First(&u).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}
	
	return &u, nil
}

// FindByUsername retrieves a user by their username
func (r *userRepository) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).
		Where("username = ? AND deleted_at IS NULL", username).
		First(&u).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to find user by username: %w", err)
	}
	
	return &u, nil
}

// Update updates an existing user's information
func (r *userRepository) Update(ctx context.Context, u *user.User) error {
	if err := r.db.WithContext(ctx).Save(u).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// Delete soft-deletes a user (sets DeletedAt timestamp)
func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).
		Model(&user.User{}).
		Where("id = ?", id).
		Update("deleted_at", gorm.Expr("NOW()")).Error; err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

// List retrieves a paginated list of users
func (r *userRepository) List(ctx context.Context, limit, offset int) ([]*user.User, int64, error) {
	var users []*user.User
	var total int64
	
	// Count total records
	if err := r.db.WithContext(ctx).
		Model(&user.User{}).
		Where("deleted_at IS NULL").
		Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count users: %w", err)
	}
	
	// Get paginated results
	if err := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&users).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to list users: %w", err)
	}
	
	return users, total, nil
}

// ExistsByKeycloakID checks if a user with the given Keycloak ID exists
func (r *userRepository) ExistsByKeycloakID(ctx context.Context, keycloakID string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&user.User{}).
		Where("keycloak_id = ? AND deleted_at IS NULL", keycloakID).
		Count(&count).Error
	
	if err != nil {
		return false, fmt.Errorf("failed to check user existence: %w", err)
	}
	
	return count > 0, nil
}