package initial_entity

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"skillsier.dev/apps/be/users-be/internal/domain/initial_entity"
	
	platformOutbox "skillsier.dev/platform-shared/outbox"
	platformOutboxPostgres "skillsier.dev/platform-shared/outbox/postgres"
)

// Service handles business logic for InitialEntity
type Service struct {
	repo            initial_entity.Repository
	outboxPublisher *platformOutbox.Publisher
	db              *gorm.DB
}

// NewService creates a new InitialEntity service
func NewService(
	repo initial_entity.Repository,
	db *gorm.DB,
) *Service {
	// Create outbox repository using platform-shared
	outboxRepo := platformOutboxPostgres.NewRepository(db)
	
	// Create outbox publisher using platform-shared
	outboxPublisher := platformOutbox.NewPublisher(outboxRepo)
	
	return &Service{
		repo:            repo,
		outboxPublisher: outboxPublisher,
		db:              db,
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

	// Begin transaction
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Save entity
		if err := tx.Create(entity).Error; err != nil {
			return err
		}

		// Create domain event
		event := initial_entity.NewInitialEntityCreatedEvent(entity)
		eventPayload, err := json.Marshal(event)
		if err != nil {
			return err
		}

		// Publish event to outbox using platform-shared
		outboxEvent := &platformOutbox.Event{
			ID:            uuid.New().String(),
			AggregateID:   entity.ID.String(),
			AggregateType: "initial_entity",
			EventType:     event.EventType(),
			Payload:       eventPayload,
			Status:        platformOutbox.StatusPending,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			MaxAttempts:   5,
		}

		return s.outboxPublisher.PublishWithTx(tx, outboxEvent)
	})

	if err != nil {
		return nil, err
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

	var updatedEntity *initial_entity.InitialEntity
	
	// Begin transaction
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Find existing entity
		entity, err := s.repo.FindByID(ctx, id)
		if err != nil {
			return err
		}

		// Check if entity is deleted
		if entity.IsDeleted() {
			return initial_entity.ErrCannotModifyDeleted
		}

		// Check if entity is archived
		if entity.Status == initial_entity.StatusArchived {
			return initial_entity.ErrCannotModifyArchived
		}

		// Store old status for status change event
		oldStatus := entity.Status

		// Apply updates from DTO
		if dto.Name != nil {
			entity.Name = *dto.Name
		}
		if dto.Description != nil {
			entity.Description = *dto.Description
		}
		if dto.Status != nil {
			// Check if status transition is allowed
			if !entity.CanTransitionTo(*dto.Status) {
				return initial_entity.ErrInvalidStatusTransition
			}
			entity.Status = *dto.Status
		}

		// Validate updated entity
		if err := entity.Validate(); err != nil {
			return err
		}

		// Save entity
		if err := tx.Save(entity).Error; err != nil {
			return err
		}

		// Create domain event
		event := initial_entity.NewInitialEntityUpdatedEvent(entity)
		eventPayload, err := json.Marshal(event)
		if err != nil {
			return err
		}

		// Publish event to outbox
		outboxEvent := &platformOutbox.Event{
			ID:            uuid.New().String(),
			AggregateID:   entity.ID.String(),
			AggregateType: "initial_entity",
			EventType:     event.EventType(),
			Payload:       eventPayload,
			Status:        platformOutbox.StatusPending,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			MaxAttempts:   5,
		}

		if err := s.outboxPublisher.PublishWithTx(tx, outboxEvent); err != nil {
			return err
		}

		// If status changed, emit status change event
		if oldStatus != entity.Status {
			statusEvent := initial_entity.NewInitialEntityStatusChangedEvent(entity.ID, oldStatus, entity.Status)
			statusPayload, _ := json.Marshal(statusEvent)
			
			statusOutboxEvent := &platformOutbox.Event{
				ID:            uuid.New().String(),
				AggregateID:   entity.ID.String(),
				AggregateType: "initial_entity",
				EventType:     statusEvent.EventType(),
				Payload:       statusPayload,
				Status:        platformOutbox.StatusPending,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
				MaxAttempts:   5,
			}
			
			_ = s.outboxPublisher.PublishWithTx(tx, statusOutboxEvent)
		}

		updatedEntity = entity
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Map entity to response DTO
	return MapEntityToResponseDTO(updatedEntity), nil
}

// Get retrieves an InitialEntity by ID
func (s *Service) Get(ctx context.Context, id uuid.UUID) (*InitialEntityResponseDTO, error) {
	entity, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Map entity to response DTO
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
		return nil, err
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
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Find existing entity
		entity, err := s.repo.FindByID(ctx, id)
		if err != nil {
			return err
		}

		// Check if already deleted
		if entity.IsDeleted() {
			return initial_entity.ErrAlreadyDeleted
		}

		// Soft delete
		now := time.Now()
		entity.DeletedAt = &now

		if err := tx.Save(entity).Error; err != nil {
			return err
		}

		// Create domain event
		event := initial_entity.NewInitialEntityDeletedEvent(entity, true)
		eventPayload, err := json.Marshal(event)
		if err != nil {
			return err
		}

		// Publish event to outbox
		outboxEvent := &platformOutbox.Event{
			ID:            uuid.New().String(),
			AggregateID:   entity.ID.String(),
			AggregateType: "initial_entity",
			EventType:     event.EventType(),
			Payload:       eventPayload,
			Status:        platformOutbox.StatusPending,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			MaxAttempts:   5,
		}

		return s.outboxPublisher.PublishWithTx(tx, outboxEvent)
	})
}

// Restore restores a soft-deleted InitialEntity
func (s *Service) Restore(ctx context.Context, id uuid.UUID) (*InitialEntityResponseDTO, error) {
	var restoredEntity *initial_entity.InitialEntity
	
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Find entity (including deleted)
		entity, err := s.repo.FindByIDWithDeleted(ctx, id)
		if err != nil {
			return err
		}

		// Check if it's deleted
		if !entity.IsDeleted() {
			return initial_entity.ErrNotDeleted
		}

		// Restore entity
		entity.DeletedAt = nil
		if err := tx.Save(entity).Error; err != nil {
			return err
		}

		// Create domain event
		event := initial_entity.NewInitialEntityRestoredEvent(entity)
		eventPayload, err := json.Marshal(event)
		if err != nil {
			return err
		}

		// Publish event to outbox
		outboxEvent := &platformOutbox.Event{
			ID:            uuid.New().String(),
			AggregateID:   entity.ID.String(),
			AggregateType: "initial_entity",
			EventType:     event.EventType(),
			Payload:       eventPayload,
			Status:        platformOutbox.StatusPending,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			MaxAttempts:   5,
		}

		if err := s.outboxPublisher.PublishWithTx(tx, outboxEvent); err != nil {
			return err
		}

		restoredEntity = entity
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Map entity to response DTO
	return MapEntityToResponseDTO(restoredEntity), nil
}

// GetByOwner retrieves all InitialEntities owned by a specific user
func (s *Service) GetByOwner(ctx context.Context, ownerID uuid.UUID) ([]*InitialEntityResponseDTO, error) {
	entities, err := s.repo.FindByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, err
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