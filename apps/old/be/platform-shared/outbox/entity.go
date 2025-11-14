package outbox

import (
	"encoding/json"
	"time"
)

// Status represents the processing status of an outbox event
type Status string

const (
	// StatusPending indicates the event is waiting to be published
	StatusPending Status = "pending"
	
	// StatusPublished indicates the event was successfully published
	StatusPublished Status = "published"
	
	// StatusFailed indicates the event failed to publish after retries
	StatusFailed Status = "failed"
)

// Event represents an outbox event that will be published to the message broker
type Event struct {
	// ID is the unique identifier for this outbox event
	ID string
	
	// AggregateID is the ID of the aggregate that generated this event
	AggregateID string
	
	// AggregateType is the type of aggregate (e.g., "user", "job", "proposal")
	AggregateType string
	
	// EventType is the type of event (e.g., "user.created", "job.posted")
	EventType string
	
	// Payload is the event data as JSON
	Payload json.RawMessage
	
	// Status is the current processing status
	Status Status
	
	// Attempts is the number of publish attempts
	Attempts int
	
	// MaxAttempts is the maximum number of retry attempts
	MaxAttempts int
	
	// LastError stores the last error message if publishing failed
	LastError string
	
	// CreatedAt is when the event was created
	CreatedAt time.Time
	
	// UpdatedAt is when the event was last updated
	UpdatedAt time.Time
	
	// PublishedAt is when the event was successfully published (null if not published)
	PublishedAt *time.Time
	
	// ScheduledFor allows delaying event publishing (null for immediate)
	ScheduledFor *time.Time
}

// NewEvent creates a new outbox event
func NewEvent(aggregateID, aggregateType, eventType string, payload interface{}) (*Event, error) {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	
	now := time.Now()
	
	return &Event{
		AggregateID:   aggregateID,
		AggregateType: aggregateType,
		EventType:     eventType,
		Payload:       payloadJSON,
		Status:        StatusPending,
		Attempts:      0,
		MaxAttempts:   5,
		CreatedAt:     now,
		UpdatedAt:     now,
	}, nil
}

// IsPending returns true if the event is pending
func (e *Event) IsPending() bool {
	return e.Status == StatusPending
}

// IsPublished returns true if the event was published
func (e *Event) IsPublished() bool {
	return e.Status == StatusPublished
}

// IsFailed returns true if the event failed
func (e *Event) IsFailed() bool {
	return e.Status == StatusFailed
}

// CanRetry returns true if the event can be retried
func (e *Event) CanRetry() bool {
	return e.Attempts < e.MaxAttempts
}

// ShouldProcess returns true if the event should be processed now
func (e *Event) ShouldProcess() bool {
	if !e.IsPending() {
		return false
	}
	
	if e.ScheduledFor == nil {
		return true
	}
	
	return time.Now().After(*e.ScheduledFor)
}

// MarkPublished marks the event as successfully published
func (e *Event) MarkPublished() {
	now := time.Now()
	e.Status = StatusPublished
	e.PublishedAt = &now
	e.UpdatedAt = now
}

// MarkFailed marks the event as failed with an error message
func (e *Event) MarkFailed(err error) {
	e.Status = StatusFailed
	e.LastError = err.Error()
	e.UpdatedAt = time.Now()
}

// IncrementAttempts increments the attempt counter
func (e *Event) IncrementAttempts() {
	e.Attempts++
	e.UpdatedAt = time.Now()
}