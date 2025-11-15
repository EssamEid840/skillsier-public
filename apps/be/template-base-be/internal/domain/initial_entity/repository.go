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
	
	// CreateWithOutbox creates a new InitialEntity and publishes an event to the outbox (atomic transaction)
	CreateWithOutbox(ctx context.Context, entity *InitialEntity, eventType string, eventPayload []byte, topic string) error
	
	// Update updates an existing InitialEntity
	Update(ctx context.Context, entity *InitialEntity) error
	
	// UpdateWithOutbox updates an InitialEntity and publishes an event to the outbox (atomic transaction)
	UpdateWithOutbox(ctx context.Context, entity *InitialEntity, eventType string, eventPayload []byte, topic string) error
	
	// Delete soft-deletes an InitialEntity
	Delete(ctx context.Context, id uuid.UUID) error
	
	// DeleteWithOutbox soft-deletes an InitialEntity and publishes an event to the outbox (atomic transaction)
	DeleteWithOutbox(ctx context.Context, id uuid.UUID, eventType string, eventPayload []byte, topic string) error
	
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
	Search    string     `json:"search,omitempty"`    // Search in name and description
	Tags      []string   `json:"tags,omitempty"`      // Filter by tags
	CreatedAfter  *string `json:"created_after,omitempty"`  // RFC3339 format
	CreatedBefore *string `json:"created_before,omitempty"` // RFC3339 format
	
	// Sorting
	SortBy    string `json:"sort_by"`    // Field to sort by (e.g., "name", "created_at")
	SortOrder string `json:"sort_order"` // "asc" or "desc"
	
	// Options
	IncludeDeleted bool `json:"include_deleted"` // Include soft-deleted entities
}

// Validate validates the list filter
func (f *ListFilter) Validate() error {
	// Set defaults
	if f.Page < 1 {
		f.Page = 1
	}
	
	if f.PageSize < 1 {
		f.PageSize = 10
	}
	
	if f.PageSize > 100 {
		f.PageSize = 100
	}
	
	if f.SortBy == "" {
		f.SortBy = "created_at"
	}
	
	if f.SortOrder == "" {
		f.SortOrder = "desc"
	}
	
	// Validate sort order
	if f.SortOrder != "asc" && f.SortOrder != "desc" {
		return NewValidationError("sort_order", "must be 'asc' or 'desc'")
	}
	
	// Validate sort field
	validSortFields := map[string]bool{
		"name":       true,
		"status":     true,
		"created_at": true,
		"updated_at": true,
	}
	
	if !validSortFields[f.SortBy] {
		return NewValidationError("sort_by", "invalid sort field")
	}
	
	// Validate status if provided
	if f.Status != nil && !f.Status.IsValid() {
		return NewValidationError("status", "invalid status value")
	}
	
	return nil
}

// GetOffset returns the offset for pagination
func (f *ListFilter) GetOffset() int {
	return (f.Page - 1) * f.PageSize
}

// GetLimit returns the limit for pagination
func (f *ListFilter) GetLimit() int {
	return f.PageSize
}