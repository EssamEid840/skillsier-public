package outbox

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"

	"<module>/internal/config"
	"<module>/internal/infrastructure/messaging/kafka"
)

// Dispatcher polls the outbox table and publishes pending events to Kafka
type Dispatcher struct {
	db        *gorm.DB
	producer  *kafka.Producer
	scheduler *Scheduler
	config    config.OutboxConfig
	stopCh    chan struct{}
}

// NewDispatcher creates a new outbox dispatcher
func NewDispatcher(db *gorm.DB, producer *kafka.Producer, cfg config.OutboxConfig) *Dispatcher {
	return &Dispatcher{
		db:        db,
		producer:  producer,
		scheduler: NewScheduler(cfg),
		config:    cfg,
		stopCh:    make(chan struct{}),
	}
}

// Start starts the outbox dispatcher
func (d *Dispatcher) Start(ctx context.Context) {
	log.Println("→ Starting outbox dispatcher...")
	log.Printf("  Poll interval: %v", d.config.PollInterval)
	log.Printf("  Batch size: %d", d.config.BatchSize)
	log.Printf("  Max retries: %d", d.config.MaxRetries)

	ticker := time.NewTicker(d.config.PollInterval)
	defer ticker.Stop()

	// Start cleanup goroutine if enabled
	if d.config.CleanupEnabled {
		go d.runCleanup(ctx)
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("✓ Outbox dispatcher stopped (context cancelled)")
			return
		case <-d.stopCh:
			log.Println("✓ Outbox dispatcher stopped")
			return
		case <-ticker.C:
			if err := d.processBatch(ctx); err != nil {
				log.Printf("⚠ Error processing outbox batch: %v", err)
			}
		}
	}
}

// Stop stops the outbox dispatcher
func (d *Dispatcher) Stop() {
	close(d.stopCh)
}

// processBatch processes a batch of pending outbox events
func (d *Dispatcher) processBatch(ctx context.Context) error {
	// Fetch pending events
	events, err := d.fetchPendingEvents(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch pending events: %w", err)
	}

	if len(events) == 0 {
		return nil // No events to process
	}

	log.Printf("→ Processing %d outbox events...", len(events))

	successCount := 0
	failureCount := 0

	for _, event := range events {
		if err := d.processEvent(ctx, event); err != nil {
			log.Printf("⚠ Failed to process event %s: %v", event.ID, err)
			failureCount++
		} else {
			successCount++
		}
	}

	log.Printf("✓ Outbox batch processed: %d succeeded, %d failed", successCount, failureCount)

	return nil
}

// fetchPendingEvents fetches pending events from the outbox table
func (d *Dispatcher) fetchPendingEvents(ctx context.Context) ([]*OutboxEvent, error) {
	var events []*OutboxEvent

	now := time.Now().Unix()

	// Fetch events that are:
	// 1. Status = 'pending'
	// 2. RetryCount < MaxRetries
	// 3. NextRetryAt is NULL or <= now
	err := d.db.WithContext(ctx).
		Where("status = ?", "pending").
		Where("retry_count < ?", d.config.MaxRetries).
		Where("next_retry_at IS NULL OR next_retry_at <= ?", now).
		Order("created_at ASC").
		Limit(d.config.BatchSize).
		Find(&events).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch pending events: %w", err)
	}

	return events, nil
}

// processEvent processes a single outbox event
func (d *Dispatcher) processEvent(ctx context.Context, event *OutboxEvent) error {
	// Publish to Kafka
	if err := d.producer.Publish(ctx, event.Topic, event.Payload); err != nil {
		// Publishing failed - schedule retry
		return d.handlePublishFailure(ctx, event, err)
	}

	// Publishing succeeded - mark as sent
	return d.markAsSent(ctx, event)
}

// handlePublishFailure handles a failed publish attempt
func (d *Dispatcher) handlePublishFailure(ctx context.Context, event *OutboxEvent, publishErr error) error {
	event.RetryCount++
	event.ErrorMessage = publishErr.Error()
	event.UpdatedAt = time.Now().Unix()

	// Check if max retries exceeded
	if event.RetryCount >= d.config.MaxRetries {
		log.Printf("⚠ Event %s exceeded max retries (%d), moving to DLQ", event.ID, d.config.MaxRetries)
		return d.moveToDLQ(ctx, event)
	}

	// Calculate next retry time using exponential backoff
	nextRetry := d.scheduler.CalculateNextRetry(event.RetryCount)
	nextRetryUnix := nextRetry.Unix()
	event.NextRetryAt = &nextRetryUnix

	// Update event with retry info
	if err := d.db.WithContext(ctx).Save(event).Error; err != nil {
		return fmt.Errorf("failed to update event retry info: %w", err)
	}

	log.Printf("→ Event %s scheduled for retry %d at %s", event.ID, event.RetryCount, nextRetry.Format(time.RFC3339))

	return nil
}

// markAsSent marks an event as successfully sent
func (d *Dispatcher) markAsSent(ctx context.Context, event *OutboxEvent) error {
	now := time.Now().Unix()
	event.Status = "sent"
	event.PublishedAt = &now
	event.UpdatedAt = now
	event.ErrorMessage = ""

	if err := d.db.WithContext(ctx).Save(event).Error; err != nil {
		return fmt.Errorf("failed to mark event as sent: %w", err)
	}

	log.Printf("✓ Event %s published successfully to topic %s", event.ID, event.Topic)

	return nil
}

// moveToDLQ moves a failed event to the dead letter queue
func (d *Dispatcher) moveToDLQ(ctx context.Context, event *OutboxEvent) error {
	event.Status = "failed"
	event.UpdatedAt = time.Now().Unix()

	if err := d.db.WithContext(ctx).Save(event).Error; err != nil {
		return fmt.Errorf("failed to move event to DLQ: %w", err)
	}

	// TODO: Optionally publish to a DLQ topic for manual review
	// dlqTopic := event.Topic + ".dlq"
	// d.producer.Publish(ctx, dlqTopic, event.Payload)

	return nil
}

// runCleanup periodically cleans up old sent events
func (d *Dispatcher) runCleanup(ctx context.Context) {
	log.Printf("→ Starting outbox cleanup (interval: %v, retention: %d days)", 
		d.config.CleanupInterval, d.config.CleanupRetentionDays)

	ticker := time.NewTicker(d.config.CleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-d.stopCh:
			return
		case <-ticker.C:
			if err := d.cleanup(ctx); err != nil {
				log.Printf("⚠ Error during outbox cleanup: %v", err)
			}
		}
	}
}

// cleanup removes old sent events
func (d *Dispatcher) cleanup(ctx context.Context) error {
	retentionDuration := time.Duration(d.config.CleanupRetentionDays) * 24 * time.Hour
	cutoffTime := time.Now().Add(-retentionDuration).Unix()

	result := d.db.WithContext(ctx).
		Where("status = ?", "sent").
		Where("published_at < ?", cutoffTime).
		Delete(&OutboxEvent{})

	if result.Error != nil {
		return fmt.Errorf("failed to cleanup old events: %w", result.Error)
	}

	if result.RowsAffected > 0 {
		log.Printf("✓ Cleaned up %d old outbox events", result.RowsAffected)
	}

	return nil
}

// GetStats returns outbox statistics
func (d *Dispatcher) GetStats(ctx context.Context) (*Stats, error) {
	stats := &Stats{}

	// Count by status
	if err := d.db.WithContext(ctx).Model(&OutboxEvent{}).
		Where("status = ?", "pending").Count(&stats.Pending).Error; err != nil {
		return nil, err
	}

	if err := d.db.WithContext(ctx).Model(&OutboxEvent{}).
		Where("status = ?", "sent").Count(&stats.Sent).Error; err != nil {
		return nil, err
	}

	if err := d.db.WithContext(ctx).Model(&OutboxEvent{}).
		Where("status = ?", "failed").Count(&stats.Failed).Error; err != nil {
		return nil, err
	}

	// Oldest pending event
	var oldestPending OutboxEvent
	if err := d.db.WithContext(ctx).
		Where("status = ?", "pending").
		Order("created_at ASC").
		First(&oldestPending).Error; err == nil {
		createdAt := time.Unix(oldestPending.CreatedAt, 0)
		stats.OldestPendingAge = time.Since(createdAt)
	}

	return stats, nil
}

// Stats represents outbox statistics
type Stats struct {
	Pending         int64
	Sent            int64
	Failed          int64
	OldestPendingAge time.Duration
}

// OutboxEvent represents an outbox event (same as in migrations.go)
type OutboxEvent struct {
	ID            string `gorm:"type:uuid;primary_key"`
	AggregateID   string `gorm:"type:varchar(255);not null;index"`
	AggregateType string `gorm:"type:varchar(100);not null"`
	EventType     string `gorm:"type:varchar(100);not null;index"`
	EventVersion  int    `gorm:"not null;default:1"`
	Payload       []byte `gorm:"type:jsonb;not null"`
	Metadata      []byte `gorm:"type:jsonb"`
	Status        string `gorm:"type:varchar(20);not null;default:'pending';index"`
	Topic         string `gorm:"type:varchar(100);not null"`
	ErrorMessage  string `gorm:"type:text"`
	PublishedAt   *int64 `gorm:"index"`
	CreatedAt     int64  `gorm:"not null;index"`
	UpdatedAt     int64  `gorm:"not null"`
	RetryCount    int    `gorm:"default:0"`
	NextRetryAt   *int64 `gorm:"index"`
}

// TableName returns the table name for GORM
func (OutboxEvent) TableName() string {
	return "outbox_events"
}