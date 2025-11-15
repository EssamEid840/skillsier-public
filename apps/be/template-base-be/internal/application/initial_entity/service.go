package initial_entity

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	"<module>/internal/domain/initial_entity"
)

// Service implements business logic for InitialEntity
type Service struct {
	repo        initial_entity.Repository
	topicPrefix string
}

// NewService creates a new InitialEntity service
func NewService(repo initial_entity.Repository, topicPrefix string) *Service {
	return &Service{
		repo:        repo,
		topicPrefix: topicPrefix,
	}
}

// Create creates a new InitialEntity
func (s *Service) Create(ctx context.Context, dto *CreateInitialEntityDTO) (*InitialEntityResponseDTO, error) {
	// Validate DTO
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	// Map DTO to entity
	entity := MapCreateDTOToEntity(dto)

	// Validate entity
	if err := entity.Validate(); err != nil {
		return nil, err
	}

	// Create domain event
	event := initial_entity.NewInitialEntityCreatedEvent(entity)

	// Serialize event payload
	eventPayload, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize event: %w", err)
	}

	// Create entity with outbox (atomic transaction)
	topic := s.topicPrefix + ".initial_entity.created"
	if err := s.repo.CreateWithOutbox(ctx, entity, event.EventType(), eventPayload, topic); err != nil {
		return nil, fmt.Errorf("failed to create initial entity: %w", err)
	}

	// Map entity to response DTO
	return MapEntityToResponseDTO(entity), nil
}

// Update updates an existing InitialEntity
func (s *Service) Update(ctx context.Context, id uuid.UUID, dto *UpdateInitialEntityDTO) (*InitialEntityResponseDTO, error) {
	// Validate DTO
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	// Find existing entity
	entity, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Check if entity is deleted
	if entity.IsDeleted() {
		return nil, initial_entity.ErrCannotModifyDeleted
	}

	// Check if entity is archived
	if entity.Status == initial_entity.StatusArchived {
		return nil, initial_entity.ErrCannotModifyArchived
	}

	// Store old status for status change event
	oldStatus := entity.Status

	// Apply updates
	if dto.Name != nil {
		entity.Name = *dto.Name
	}
	if dto.Description != nil {
		entity.Description = *dto.Description
	}
	if dto.Status != nil {
		// Check if status transition is allowed
		if !entity.CanTransitionTo(*dto.Status) {
			return nil, initial_entity.ErrInvalidStatusTransition
		}
		entity.Status = *dto.Status
	}

	// Validate updated entity
	if err := entity.Validate(); err != nil {
		return nil, err
	}

	// Create domain event
	event := initial_entity.NewInitialEntityUpdatedEvent(entity)

	// Serialize event payload
	eventPayload, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize event: %w", err)
	}

	// Update entity with outbox (atomic transaction)
	topic := s.topicPrefix + ".initial_entity.updated"
	if err := s.repo.UpdateWithOutbox(ctx, entity, event.EventType(), eventPayload, topic); err != nil {
		return nil, fmt.Errorf("failed to update initial entity: %w", err)
	}

	// If status changed, emit status change event
	if oldStatus != entity.Status {
		statusEvent := initial_entity.NewInitialEntityStatusChangedEvent(entity.ID, oldStatus, entity.Status)
		statusPayload, _ := json.Marshal(statusEvent)
		statusTopic := s.topicPrefix + ".initial_entity.status_changed"
		
		// Note: This is a separate event, not critical if it fails
		// In production, you might want to handle this differently
		_ = s.repo.UpdateWithOutbox(ctx, entity, statusEvent.EventType(), statusPayload, statusTopic)
	}

	return MapEntityToResponseDTO(entity), nil
}

// Get retrieves an InitialEntity by ID
func (s *Service) Get(ctx context.Context, id uuid.UUID) (*InitialEntityResponseDTO, error) {
	entity, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return MapEntityToResponseDTO(entity), nil
}

// List retrieves InitialEntities with pagination and filters
func (s *Service) List(ctx context.Context, dto *ListInitialEntitiesDTO) (*ListInitialEntitiesResponseDTO, error) {
	// Validate DTO
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	// Map DTO to domain filter
	filter := MapListDTOToFilter(dto)

	// Fetch entities
	entities, total, err := s.repo.List(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to list initial entities: %w", err)
	}

	// Map to response DTOs
	items := make([]*InitialEntityResponseDTO, len(entities))
	for i, entity := range entities {
		items[i] = MapEntityToResponseDTO(entity)
	}

	// Build response
	response := &ListInitialEntitiesResponseDTO{
		Items: items,
		Pagination: PaginationResponse{
			Page:       dto.Page,
			PageSize:   dto.PageSize,
			TotalItems: total,
			TotalPages: (total + int64(dto.PageSize) - 1) / int64(dto.PageSize),
		},
	}

	return response, nil
}

// Delete soft-deletes an InitialEntity
func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	// Find existing entity to get details for event
	entity, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Check if already deleted
	if entity.IsDeleted() {
		return initial_entity.ErrAlreadyDeleted
	}

	// Create domain event
	event := initial_entity.NewInitialEntityDeletedEvent(entity, true) // true = soft delete

	// Serialize event payload
	eventPayload, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to serialize event: %w", err)
	}

	// Delete entity with outbox (atomic transaction)
	topic := s.topicPrefix + ".initial_entity.deleted"
	if err := s.repo.DeleteWithOutbox(ctx, id, event.EventType(), eventPayload, topic); err != nil {
		return fmt.Errorf("failed to delete initial entity: %w", err)
	}

	return nil
}

// Restore restores a soft-deleted InitialEntity
func (s *Service) Restore(ctx context.Context, id uuid.UUID) (*InitialEntityResponseDTO, error) {
	// Find entity (including deleted)
	entity, err := s.repo.FindByIDWithDeleted(ctx, id)
	if err != nil {
		return nil, err
	}

	// Check if it's deleted
	if !entity.IsDeleted() {
		return nil, fmt.Errorf("entity is not deleted")
	}

	// Restore entity
	if err := s.repo.Restore(ctx, id); err != nil {
		return nil, fmt.Errorf("failed to restore initial entity: %w", err)
	}

	// Fetch restored entity
	restoredEntity, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Create domain event
	event := initial_entity.NewInitialEntityRestoredEvent(restoredEntity)

	// Serialize event payload
	eventPayload, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize event: %w", err)
	}

	// Publish event (not critical if it fails)
	topic := s.topicPrefix + ".initial_entity.restored"
	_ = s.repo.UpdateWithOutbox(ctx, restoredEntity, event.EventType(), eventPayload, topic)

	return MapEntityToResponseDTO(restoredEntity), nil
}

// GetByOwner retrieves all InitialEntities owned by a specific user
func (s *Service) GetByOwner(ctx context.Context, ownerID uuid.UUID) ([]*InitialEntityResponseDTO, error) {
	entities, err := s.repo.FindByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get entities by owner: %w", err)
	}

	// Map to response DTOs
	items := make([]*InitialEntityResponseDTO, len(entities))
	for i, entity := range entities {
		items[i] = MapEntityToResponseDTO(entity)
	}

	return items, nil
}

// CountByStatus counts entities by status
func (s *Service) CountByStatus(ctx context.Context, status initial_entity.Status) (int64, error) {
	return s.repo.CountByStatus(ctx, status)
}