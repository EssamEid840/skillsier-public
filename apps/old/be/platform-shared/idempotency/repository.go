package idempotency

import "context"

// Repository defines the interface for idempotency record storage
type Repository interface {
	// Get retrieves an idempotency record by key
	Get(ctx context.Context, key string) (*Record, error)
	
	// Create stores a new idempotency record
	Create(ctx context.Context, record *Record) error
	
	// Delete removes an idempotency record
	Delete(ctx context.Context, key string) error
	
	// DeleteExpired removes expired idempotency records
	DeleteExpired(ctx context.Context) error
}