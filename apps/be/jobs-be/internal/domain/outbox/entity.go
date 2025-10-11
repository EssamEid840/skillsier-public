package outbox

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// EventStatus represents the processing status of an outbox event
type EventStatus string

const (
	// EventStatusPending means the event is waiting to be published
	EventStatusPending EventStatus = "pending"
	// EventStatusPublished means the event has been successfully published to Kafka
	EventStatusPublished EventStatus = "published"
	// EventStatusFailed means publishing the event failed
	EventStatusFailed EventStatus = "failed"
)

// Event represents an outbox event that will be published to Kafka
// This implements the Transactional Outbox Pattern
// Events are saved in the same transaction as business data, then published asynchronously
type Event struct {
	// ID is the unique identifier for this event
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	
	// AggregateID is the ID of the entity this event relates to (e.g., user ID)
	AggregateID string `gorm:"type:varchar(255);not null;index" json:"aggregate_id"`
	
	// AggregateType identifies the type of aggregate (e.g., "user", "project")
	AggregateType string `gorm:"type:varchar(50);not null;index" json:"aggregate_type"`
	
	// EventType describes what happened (e.g., "user.created", "user.updated")
	EventType string `gorm:"type:varchar(100);not null;index" json:"event_type"`
	
	// Version supports event versioning for future event sourcing
	// Increment this when the event schema changes
	Version int `gorm:"type:integer;not null;default:1" json:"version"`
	
	// Payload contains the event data as JSON
	// This is the actual event content that will be published
	Payload json.RawMessage `gorm:"type:jsonb;not null" json:"payload"`
	
	// Metadata stores additional contextual information
	Metadata json.RawMessage `gorm:"type:jsonb" json:"metadata,omitempty"`
	
	// Status tracks whether the event has been published
	Status EventStatus `gorm:"type:varchar(20);not null;default:'pending';index" json:"status"`
	
	// Topic is the Kafka topic where this event should be published
	Topic string `gorm:"type:varchar(100);not null" json:"topic"`
	
	// ErrorMessage stores any error that occurred during publishing
	ErrorMessage *string `gorm:"type:text" json:"error_message,omitempty"`
	
	// PublishedAt is when the event was successfully published
	PublishedAt *time.Time `json:"published_at,omitempty"`
	
	// Timestamps
	CreatedAt time.Time `gorm:"not null;index" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	
	// RetryCount tracks how many times we've tried to publish this event
	RetryCount int `gorm:"type:integer;default:0" json:"retry_count"`
	
	// NextRetryAt indicates when to retry publishing if it failed
	NextRetryAt *time.Time `gorm:"index" json:"next_retry_at,omitempty"`
}

// TableName specifies the table name for GORM
func (Event) TableName() string {
	return "outbox_events"
}

// BeforeCreate is a GORM hook that runs before creating a new event
func (e *Event) BeforeCreate(tx *gorm.DB) error {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	now := time.Now()
	e.CreatedAt = now
	e.UpdatedAt = now
	
	if e.Status == "" {
		e.Status = EventStatusPending
	}
	
	if e.Version == 0 {
		e.Version = 1
	}
	
	return nil
}

// BeforeUpdate is a GORM hook that runs before updating an event
func (e *Event) BeforeUpdate(tx *gorm.DB) error {
	e.UpdatedAt = time.Now()
	return nil
}

// MarkAsPublished marks the event as successfully published
func (e *Event) MarkAsPublished() {
	now := time.Now()
	e.Status = EventStatusPublished
	e.PublishedAt = &now
	e.UpdatedAt = now
}

// MarkAsFailed marks the event as failed with an error message
func (e *Event) MarkAsFailed(errorMsg string) {
	e.Status = EventStatusFailed
	e.ErrorMessage = &errorMsg
	e.RetryCount++
	e.UpdatedAt = time.Now()
	
	// Calculate exponential backoff for retry
	// Wait 2^retryCount minutes before next retry (max 60 minutes)
	backoffMinutes := 1 << uint(e.RetryCount) // 2, 4, 8, 16, 32, 64...
	if backoffMinutes > 60 {
		backoffMinutes = 60
	}
	nextRetry := time.Now().Add(time.Duration(backoffMinutes) * time.Minute)
	e.NextRetryAt = &nextRetry
}

// CanRetry checks if this event should be retried
func (e *Event) CanRetry(maxRetries int) bool {
	if e.Status != EventStatusFailed {
		return false
	}
	
	if e.RetryCount >= maxRetries {
		return false
	}
	
	if e.NextRetryAt != nil && time.Now().Before(*e.NextRetryAt) {
		return false
	}
	
	return true
}
