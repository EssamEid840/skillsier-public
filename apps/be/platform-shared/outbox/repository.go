package outbox

import "context"

// Repository defines the interface for outbox event storage
type Repository interface {
	// Create creates a new outbox event
	Create(ctx context.Context, event *Event) error
	
	// FindPending retrieves pending events up to the specified limit
	FindPending(ctx context.Context, limit int) ([]*Event, error)
	
	// MarkPublished marks an event as published
	MarkPublished(ctx context.Context, eventID string) error
	
	// MarkFailed marks an event as failed
	MarkFailed(ctx context.Context, eventID string, err error) error
	
	// IncrementAttempts increments the attempt counter for an event
	IncrementAttempts(ctx context.Context, eventID string) error
	
	// Delete removes an event (for cleanup of old published events)
	Delete(ctx context.Context, eventID string) error
	
	// DeletePublished removes published events older than the specified duration
	DeletePublished(ctx context.Context, olderThan int) error
}