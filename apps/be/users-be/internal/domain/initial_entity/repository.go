package initial_entity

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Repository defines the interface for InitialEntity persistence operations
type Repository interface {
	// Create creates a new InitialEntity
	Create(ctx context.Context, entity *InitialEntity) error
	
	// Update updates an existing InitialEntity
	Update(ctx context.Context, entity *InitialEntity) error
	
	// Delete soft-deletes an InitialEntity
	Delete(ctx context.Context, id uuid.UUID) error
	
	// FindByID retrieves an InitialEntity by ID
	FindByID(ctx context.Context, id uuid.UUID) (*InitialEntity, error)
	
	// FindByIDWithDeleted retrieves an InitialEntity by ID including soft-deleted entities
	FindByIDWithDeleted(ctx context.Context, id uuid.UUID) (*InitialEntity, error)
	
	// FindByOwnerID retrieves all InitialEntities owned by a specific user
	FindByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]*InitialEntity, error)
	
	// List retrieves InitialEntities with pagination and filters
	List(ctx context.Context, filter *ListFilter) ([]*InitialEntity, int64, error)
	
	// Exists checks if an InitialEntity with the given ID exists
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
	
	// Restore restores a soft-deleted InitialEntity
	Restore(ctx context.Context, id uuid.UUID) error
	
	// HardDelete permanently deletes an InitialEntity
	HardDelete(ctx context.Context, id uuid.UUID) error
	
	// CountByStatus counts entities by status
	CountByStatus(ctx context.Context, status Status) (int64, error)
	
	// WithTx returns a new repository instance bound to the given transaction
	WithTx(tx *gorm.DB) Repository
}

// ListFilter defines filtering and pagination options for listing entities
type ListFilter struct {
	// Pagination
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	
	// Filters
	Status    *Status    `json:"status,omitempty"`
	OwnerID   *uuid.UUID `json:"owner_id,omitempty"`
	Search    string     `json:"search,omitempty"`
	Tags      []string   `json:"tags,omitempty"`

	SortBy        string                 `form:"sort_by" binding:"omitempty,oneof=name status created_at updated_at"`
	SortOrder     string                 `form:"sort_order" binding:"omitempty,oneof=asc desc"`
	IncludeDeleted bool                  `form:"include_deleted"`
}