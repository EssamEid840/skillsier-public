package initial_entity

import (
	"time"

	"github.com/google/uuid"
)

// DomainEvent represents a domain event for InitialEntity
type DomainEvent interface {
	EventType() string
	AggregateID() uuid.UUID
	OccurredAt() time.Time
}

// BaseEvent contains common event fields
type BaseEvent struct {
	EventID      uuid.UUID `json:"event_id"`
	EntityID     uuid.UUID `json:"entity_id"`
	EventTime    time.Time `json:"event_time"`
	EventVersion int       `json:"event_version"`
}

// NewBaseEvent creates a new base event
func NewBaseEvent(entityID uuid.UUID) BaseEvent {
	return BaseEvent{
		EventID:      uuid.New(),
		EntityID:     entityID,
		EventTime:    time.Now(),
		EventVersion: 1,
	}
}

// InitialEntityCreated is emitted when a new InitialEntity is created
type InitialEntityCreated struct {
	BaseEvent
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Status      Status            `json:"status"`
	OwnerID     uuid.UUID         `json:"owner_id"`
	Tags        []string          `json:"tags,omitempty"`
	Properties  map[string]string `json:"properties,omitempty"`
}

// EventType returns the event type
func (e InitialEntityCreated) EventType() string {
	return "initial_entity.created.v1"
}

// AggregateID returns the aggregate ID
func (e InitialEntityCreated) AggregateID() uuid.UUID {
	return e.EntityID
}

// OccurredAt returns when the event occurred
func (e InitialEntityCreated) OccurredAt() time.Time {
	return e.EventTime
}

// NewInitialEntityCreatedEvent creates a new InitialEntityCreated event
func NewInitialEntityCreatedEvent(entity *InitialEntity) *InitialEntityCreated {
	return &InitialEntityCreated{
		BaseEvent:   NewBaseEvent(entity.ID),
		Name:        entity.Name,
		Description: entity.Description,
		Status:      entity.Status,
		OwnerID:     entity.OwnerID,
		Tags:        entity.Metadata.Tags,
		Properties:  entity.Metadata.Properties,
	}
}

// InitialEntityUpdated is emitted when an InitialEntity is updated
type InitialEntityUpdated struct {
	BaseEvent
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Status      Status            `json:"status"`
	Tags        []string          `json:"tags,omitempty"`
	Properties  map[string]string `json:"properties,omitempty"`
	Version     int               `json:"version"`
}

// EventType returns the event type
func (e InitialEntityUpdated) EventType() string {
	return "initial_entity.updated.v1"
}

// AggregateID returns the aggregate ID
func (e InitialEntityUpdated) AggregateID() uuid.UUID {
	return e.EntityID
}

// OccurredAt returns when the event occurred
func (e InitialEntityUpdated) OccurredAt() time.Time {
	return e.EventTime
}

// NewInitialEntityUpdatedEvent creates a new InitialEntityUpdated event
func NewInitialEntityUpdatedEvent(entity *InitialEntity) *InitialEntityUpdated {
	return &InitialEntityUpdated{
		BaseEvent:   NewBaseEvent(entity.ID),
		Name:        entity.Name,
		Description: entity.Description,
		Status:      entity.Status,
		Tags:        entity.Metadata.Tags,
		Properties:  entity.Metadata.Properties,
		Version:     entity.Metadata.Version,
	}
}

// InitialEntityStatusChanged is emitted when an InitialEntity's status changes
type InitialEntityStatusChanged struct {
	BaseEvent
	OldStatus Status `json:"old_status"`
	NewStatus Status `json:"new_status"`
}

// EventType returns the event type
func (e InitialEntityStatusChanged) EventType() string {
	return "initial_entity.status_changed.v1"
}

// AggregateID returns the aggregate ID
func (e InitialEntityStatusChanged) AggregateID() uuid.UUID {
	return e.EntityID
}

// OccurredAt returns when the event occurred
func (e InitialEntityStatusChanged) OccurredAt() time.Time {
	return e.EventTime
}

// NewInitialEntityStatusChangedEvent creates a new InitialEntityStatusChanged event
func NewInitialEntityStatusChangedEvent(entityID uuid.UUID, oldStatus, newStatus Status) *InitialEntityStatusChanged {
	return &InitialEntityStatusChanged{
		BaseEvent: NewBaseEvent(entityID),
		OldStatus: oldStatus,
		NewStatus: newStatus,
	}
}

// InitialEntityDeleted is emitted when an InitialEntity is deleted
type InitialEntityDeleted struct {
	BaseEvent
	OwnerID uuid.UUID `json:"owner_id"`
	Name    string    `json:"name"`
	SoftDelete bool   `json:"soft_delete"` // true for soft delete, false for hard delete
}

// EventType returns the event type
func (e InitialEntityDeleted) EventType() string {
	return "initial_entity.deleted.v1"
}

// AggregateID returns the aggregate ID
func (e InitialEntityDeleted) AggregateID() uuid.UUID {
	return e.EntityID
}

// OccurredAt returns when the event occurred
func (e InitialEntityDeleted) OccurredAt() time.Time {
	return e.EventTime
}

// NewInitialEntityDeletedEvent creates a new InitialEntityDeleted event
func NewInitialEntityDeletedEvent(entity *InitialEntity, softDelete bool) *InitialEntityDeleted {
	return &InitialEntityDeleted{
		BaseEvent:  NewBaseEvent(entity.ID),
		OwnerID:    entity.OwnerID,
		Name:       entity.Name,
		SoftDelete: softDelete,
	}
}

// InitialEntityRestored is emitted when a soft-deleted InitialEntity is restored
type InitialEntityRestored struct {
	BaseEvent
	OwnerID uuid.UUID `json:"owner_id"`
	Name    string    `json:"name"`
}

// EventType returns the event type
func (e InitialEntityRestored) EventType() string {
	return "initial_entity.restored.v1"
}

// AggregateID returns the aggregate ID
func (e InitialEntityRestored) AggregateID() uuid.UUID {
	return e.EntityID
}

// OccurredAt returns when the event occurred
func (e InitialEntityRestored) OccurredAt() time.Time {
	return e.EventTime
}

// NewInitialEntityRestoredEvent creates a new InitialEntityRestored event
func NewInitialEntityRestoredEvent(entity *InitialEntity) *InitialEntityRestored {
	return &InitialEntityRestored{
		BaseEvent: NewBaseEvent(entity.ID),
		OwnerID:   entity.OwnerID,
		Name:      entity.Name,
	}
}