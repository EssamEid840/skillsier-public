package outbox

import (
	"context"

	"github.com/google/uuid"
)

// Repository defines the interface for outbox event persistence
type Repository interface {
	// Create creates a new outbox event
	// This should be called in the same transaction as the business operation
	Create(ctx context.Context, event *Event) error
	
	// FindPendingEvents retrieves events that need to be published
	// Returns events with status 'pending' or 'failed' events that are ready to retry
	FindPendingEvents(ctx context.Context, limit int) ([]*Event, error)
	
	// Update updates an existing outbox event
	// Used to mark events as published or failed
	Update(ctx context.Context, event *Event) error
	
	// Delete removes an event from the outbox
	// Typically used to clean up old published events
	Delete(ctx context.Context, id uuid.UUID) error
	
	// FindByAggregateID retrieves all events for a specific aggregate
	// Useful for event sourcing and audit trails
	FindByAggregateID(ctx context.Context, aggregateID string, limit, offset int) ([]*Event, error)
	
	// DeletePublished removes published events older than the specified duration
	// Used for periodic cleanup to prevent the outbox table from growing indefinitely
	DeletePublished(ctx context.Context, olderThanDays int) (int64, error)
}