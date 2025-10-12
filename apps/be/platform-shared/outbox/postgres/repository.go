package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"skillsier.dev/platform-shared/outbox"
)

// OutboxEvent is the GORM model for outbox events
type OutboxEvent struct {
	ID            string    `gorm:"type:uuid;primaryKey"`
	AggregateID   string    `gorm:"type:varchar(255);not null;index"`
	AggregateType string    `gorm:"type:varchar(100);not null"`
	EventType     string    `gorm:"type:varchar(255);not null;index"`
	Payload       []byte    `gorm:"type:jsonb;not null"`
	Status        string    `gorm:"type:varchar(20);not null;index;default:'pending'"`
	Attempts      int       `gorm:"not null;default:0"`
	MaxAttempts   int       `gorm:"not null;default:5"`
	LastError     string    `gorm:"type:text"`
	CreatedAt     time.Time `gorm:"not null;index"`
	UpdatedAt     time.Time `gorm:"not null"`
	PublishedAt   *time.Time
	ScheduledFor  *time.Time `gorm:"index"`
}

// TableName specifies the table name for GORM
func (OutboxEvent) TableName() string {
	return "outbox_events"
}

// Repository implements outbox.Repository using PostgreSQL
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new PostgreSQL outbox repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Create creates a new outbox event
func (r *Repository) Create(ctx context.Context, event *outbox.Event) error {
	// Generate ID if not set
	if event.ID == "" {
		event.ID = uuid.New().String()
	}
	
	model := &OutboxEvent{
		ID:            event.ID,
		AggregateID:   event.AggregateID,
		AggregateType: event.AggregateType,
		EventType:     event.EventType,
		Payload:       event.Payload,
		Status:        string(event.Status),
		Attempts:      event.Attempts,
		MaxAttempts:   event.MaxAttempts,
		LastError:     event.LastError,
		CreatedAt:     event.CreatedAt,
		UpdatedAt:     event.UpdatedAt,
		PublishedAt:   event.PublishedAt,
		ScheduledFor:  event.ScheduledFor,
	}
	
	return r.db.WithContext(ctx).Create(model).Error
}

// FindPending retrieves pending events up to the specified limit
func (r *Repository) FindPending(ctx context.Context, limit int) ([]*outbox.Event, error) {
	var models []OutboxEvent
	
	err := r.db.WithContext(ctx).
		Where("status = ?", outbox.StatusPending).
		Where("(scheduled_for IS NULL OR scheduled_for <= ?)", time.Now()).
		Order("created_at ASC").
		Limit(limit).
		Find(&models).Error
	
	if err != nil {
		return nil, err
	}
	
	events := make([]*outbox.Event, len(models))
	for i, model := range models {
		events[i] = toEvent(&model)
	}
	
	return events, nil
}

// MarkPublished marks an event as published
func (r *Repository) MarkPublished(ctx context.Context, eventID string) error {
	now := time.Now()
	return r.db.WithContext(ctx).
		Model(&OutboxEvent{}).
		Where("id = ?", eventID).
		Updates(map[string]interface{}{
			"status":       outbox.StatusPublished,
			"published_at": now,
			"updated_at":   now,
		}).Error
}

// MarkFailed marks an event as failed
func (r *Repository) MarkFailed(ctx context.Context, eventID string, err error) error {
	return r.db.WithContext(ctx).
		Model(&OutboxEvent{}).
		Where("id = ?", eventID).
		Updates(map[string]interface{}{
			"status":     outbox.StatusFailed,
			"last_error": err.Error(),
			"updated_at": time.Now(),
		}).Error
}

// IncrementAttempts increments the attempt counter for an event
func (r *Repository) IncrementAttempts(ctx context.Context, eventID string) error {
	return r.db.WithContext(ctx).
		Model(&OutboxEvent{}).
		Where("id = ?", eventID).
		Updates(map[string]interface{}{
			"attempts":   gorm.Expr("attempts + 1"),
			"updated_at": time.Now(),
		}).Error
}

// Delete removes an event
func (r *Repository) Delete(ctx context.Context, eventID string) error {
	return r.db.WithContext(ctx).
		Where("id = ?", eventID).
		Delete(&OutboxEvent{}).Error
}

// DeletePublished removes published events older than the specified duration
func (r *Repository) DeletePublished(ctx context.Context, olderThanDays int) error {
	cutoff := time.Now().AddDate(0, 0, -olderThanDays)
	
	return r.db.WithContext(ctx).
		Where("status = ?", outbox.StatusPublished).
		Where("published_at < ?", cutoff).
		Delete(&OutboxEvent{}).Error
}

// toEvent converts a GORM model to a domain event
func toEvent(model *OutboxEvent) *outbox.Event {
	return &outbox.Event{
		ID:            model.ID,
		AggregateID:   model.AggregateID,
		AggregateType: model.AggregateType,
		EventType:     model.EventType,
		Payload:       model.Payload,
		Status:        outbox.Status(model.Status),
		Attempts:      model.Attempts,
		MaxAttempts:   model.MaxAttempts,
		LastError:     model.LastError,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		PublishedAt:   model.PublishedAt,
		ScheduledFor:  model.ScheduledFor,
	}
}

// AutoMigrate runs database migrations for the outbox table
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&OutboxEvent{})
}