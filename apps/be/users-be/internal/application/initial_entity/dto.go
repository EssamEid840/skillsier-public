package initial_entity

import (
	"time"

	"github.com/google/uuid"

	"skillsier.dev/apps/be/users-be/internal/domain/initial_entity"
)

// CreateInitialEntityDTO represents the request to create an InitialEntity
type CreateInitialEntityDTO struct {
	Name        string                       `json:"name" binding:"required,min=3,max=255"`
	Description string                       `json:"description" binding:"max=1000"`
	Status      initial_entity.Status        `json:"status" binding:"omitempty,oneof=active inactive"`
	OwnerID     uuid.UUID                    `json:"owner_id" binding:"required"`
	Tags        []string                     `json:"tags" binding:"omitempty,dive,min=1,max=50"`
	Properties  map[string]string            `json:"properties" binding:"omitempty"`
}

// Validate validates the CreateInitialEntityDTO
func (dto *CreateInitialEntityDTO) Validate() error {
	if dto.Name == "" {
		return initial_entity.ErrNameRequired
	}
	if len(dto.Name) < 3 {
		return initial_entity.ErrNameTooShort
	}
	if len(dto.Name) > 255 {
		return initial_entity.ErrNameTooLong
	}
	if dto.OwnerID == uuid.Nil {
		return initial_entity.ErrOwnerIDRequired
	}
	if dto.Status != "" && !dto.Status.IsValid() {
		return initial_entity.ErrInvalidStatus
	}
	return nil
}

// UpdateInitialEntityDTO represents the request to update an InitialEntity
type UpdateInitialEntityDTO struct {
	Name        *string                `json:"name" binding:"omitempty,min=3,max=255"`
	Description *string                `json:"description" binding:"omitempty,max=1000"`
	Status      *initial_entity.Status `json:"status" binding:"omitempty,oneof=active inactive archived"`
	Tags        []string               `json:"tags" binding:"omitempty,dive,min=1,max=50"`
	Properties  map[string]string      `json:"properties" binding:"omitempty"`
}

// Validate validates the UpdateInitialEntityDTO
func (dto *UpdateInitialEntityDTO) Validate() error {
	if dto.Name != nil {
		if len(*dto.Name) < 3 {
			return initial_entity.ErrNameTooShort
		}
		if len(*dto.Name) > 255 {
			return initial_entity.ErrNameTooLong
		}
	}
	if dto.Status != nil && !dto.Status.IsValid() {
		return initial_entity.ErrInvalidStatus
	}
	return nil
}

// ListInitialEntitiesDTO represents the request to list InitialEntities with filters
type ListInitialEntitiesDTO struct {
	Page          int                    `form:"page" binding:"omitempty,min=1"`
	PageSize      int                    `form:"page_size" binding:"omitempty,min=1,max=100"`
	Status        *initial_entity.Status `form:"status" binding:"omitempty,oneof=active inactive archived"`
	OwnerID       *uuid.UUID             `form:"owner_id" binding:"omitempty"`
	Search        string                 `form:"search" binding:"omitempty,max=255"`
	Tags          []string               `form:"tags" binding:"omitempty"`
	SortBy        string                 `form:"sort_by" binding:"omitempty,oneof=name status created_at updated_at"`
	SortOrder     string                 `form:"sort_order" binding:"omitempty,oneof=asc desc"`
	IncludeDeleted bool                  `form:"include_deleted"`
}

// Validate validates the ListInitialEntitiesDTO
func (dto *ListInitialEntitiesDTO) Validate() error {
	// Set defaults
	if dto.Page < 1 {
		dto.Page = 1
	}
	if dto.PageSize < 1 {
		dto.PageSize = 10
	}
	if dto.PageSize > 100 {
		dto.PageSize = 100
	}
	if dto.SortBy == "" {
		dto.SortBy = "created_at"
	}
	if dto.SortOrder == "" {
		dto.SortOrder = "desc"
	}

	// Validate status if provided
	if dto.Status != nil && !dto.Status.IsValid() {
		return initial_entity.ErrInvalidStatus
	}

	return nil
}

// InitialEntityResponseDTO represents the response for an InitialEntity
type InitialEntityResponseDTO struct {
	ID          uuid.UUID                `json:"id"`
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Status      initial_entity.Status    `json:"status"`
	OwnerID     uuid.UUID                `json:"owner_id"`
	Tags        []string                 `json:"tags,omitempty"`
	Properties  map[string]string        `json:"properties,omitempty"`
	Version     int                      `json:"version"`
	CreatedAt   time.Time                `json:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at"`
	DeletedAt   *time.Time               `json:"deleted_at,omitempty"`
}

// ListInitialEntitiesResponseDTO represents the response for listing InitialEntities
type ListInitialEntitiesResponseDTO struct {
	Items      []*InitialEntityResponseDTO `json:"items"`
	Pagination PaginationResponse          `json:"pagination"`
}

// PaginationResponse represents pagination metadata
type PaginationResponse struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalItems int64 `json:"total_items"`
	TotalPages int64 `json:"total_pages"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string                 `json:"error"`
	Message string                 `json:"message"`
	Code    string                 `json:"code,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// SuccessResponse represents a generic success response
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}