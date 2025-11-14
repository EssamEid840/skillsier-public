package postgres

import (
	"context"
	"fmt"
	"time"

	"users-be/internal/domain/outbox"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// outboxRepository implements the outbox.Repository interface using PostgreSQL
type outboxRepository struct {
	db *gorm.DB
}

// NewOutboxRepository creates a new instance of the outbox repository
func NewOutboxRepository(db *gorm.DB) outbox.Repository {
	return &outboxRepository{db: db}
}

// Create creates a new outbox event
func (r *outboxRepository) Create(ctx context.Context, event *outbox.Event) error {
	if err := r.db.WithContext(ctx).Create(event).Error; err != nil {
		return fmt.Errorf("failed to create outbox event: %w", err)
	}
	return nil
}

// FindPendingEvents retrieves events that need to be published
// Returns pending events and failed events that are ready to retry
func (r *outboxRepository) FindPendingEvents(ctx context.Context, limit int) ([]*outbox.Event, error) {
	var events []*outbox.Event
	now := time.Now()
	
	err := r.db.WithContext(ctx).
		Where("status = ?", outbox.EventStatusPending).
		Or("(status = ? AND (next_retry_at IS NULL OR next_retry_at <= ?))", 
			outbox.EventStatusFailed, now).
		Order("created_at ASC").
		Limit(limit).
		Find(&events).Error
	
	if err != nil {
		return nil, fmt.Errorf("failed to find pending events: %w", err)
	}
	
	return events, nil
}

// Update updates an existing outbox event
func (r *outboxRepository) Update(ctx context.Context, event *outbox.Event) error {
	if err := r.db.WithContext(ctx).Save(event).Error; err != nil {
		return fmt.Errorf("failed to update outbox event: %w", err)
	}
	return nil
}

// Delete removes an event from the outbox
func (r *outboxRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).
		Delete(&outbox.Event{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete outbox event: %w", err)
	}
	return nil
}

// FindByAggregateID retrieves all events for a specific aggregate
func (r *outboxRepository) FindByAggregateID(ctx context.Context, aggregateID string, limit, offset int) ([]*outbox.Event, error) {
	var events []*outbox.Event
	
	err := r.db.WithContext(ctx).
		Where("aggregate_id = ?", aggregateID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&events).Error
	
	if err != nil {
		return nil, fmt.Errorf("failed to find events by aggregate ID: %w", err)
	}
	
	return events, nil
}

// DeletePublished removes published events older than the specified duration
func (r *outboxRepository) DeletePublished(ctx context.Context, olderThanDays int) (int64, error) {
	cutoffDate := time.Now().AddDate(0, 0, -olderThanDays)
	
	result := r.db.WithContext(ctx).
		Where("status = ? AND published_at < ?", outbox.EventStatusPublished, cutoffDate).
		Delete(&outbox.Event{})
	
	if result.Error != nil {
		return 0, fmt.Errorf("failed to delete published events: %w", result.Error)
	}
	
	return result.RowsAffected, nil
}