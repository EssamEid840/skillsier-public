package inbox

import "context"

// Repository defines the interface for inbox message storage
type Repository interface {
	// Exists checks if a message has already been processed
	Exists(ctx context.Context, messageID, handler string) (bool, error)
	
	// Create records a processed message
	Create(ctx context.Context, message *Message) error
	
	// Delete removes a message record (for cleanup)
	Delete(ctx context.Context, messageID, handler string) error
	
	// DeleteOld removes message records older than the specified duration (in days)
	DeleteOld(ctx context.Context, olderThanDays int) error
}