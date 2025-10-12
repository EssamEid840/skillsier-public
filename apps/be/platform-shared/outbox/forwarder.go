package outbox

import (
	"context"
	"fmt"
	"time"

	"skillsier.dev/platform-shared/logging"
)

// MessagePublisher is the interface for publishing messages to a message broker (Kafka, RabbitMQ, etc.)
type MessagePublisher interface {
	Publish(ctx context.Context, topic string, key string, payload []byte) error
}

// Forwarder polls the outbox table and publishes pending events to the message broker
type Forwarder struct {
	repo      Repository
	publisher MessagePublisher
	scheduler *Scheduler
	logger    *logging.Logger
	config    *Config
}

// NewForwarder creates a new outbox forwarder
func NewForwarder(
	repo Repository,
	publisher MessagePublisher,
	config *Config,
	logger *logging.Logger,
) *Forwarder {
	return &Forwarder{
		repo:      repo,
		publisher: publisher,
		scheduler: NewScheduler(config),
		logger:    logger,
		config:    config,
	}
}

// Start begins processing outbox events
func (f *Forwarder) Start(ctx context.Context) error {
	f.logger.Info("Starting outbox forwarder...")
	
	ticker := time.NewTicker(f.config.PollInterval)
	defer ticker.Stop()
	
	// Process immediately on start
	if err := f.processBatch(ctx); err != nil {
		f.logger.WithError(err).Error("Error processing initial batch")
	}
	
	// Then process on each tick
	for {
		select {
		case <-ctx.Done():
			f.logger.Info("Stopping outbox forwarder...")
			return ctx.Err()
		case <-ticker.C:
			if err := f.processBatch(ctx); err != nil {
				f.logger.WithError(err).Error("Error processing batch")
			}
		}
	}
}

// processBatch retrieves and processes a batch of pending events
func (f *Forwarder) processBatch(ctx context.Context) error {
	// Retrieve pending events
	events, err := f.repo.FindPending(ctx, f.config.BatchSize)
	if err != nil {
		return fmt.Errorf("failed to retrieve pending events: %w", err)
	}
	
	if len(events) == 0 {
		return nil
	}
	
	f.logger.Infof("Processing %d pending events", len(events))
	
	// Process each event
	for _, event := range events {
		if err := f.processEvent(ctx, event); err != nil {
			f.logger.
				WithError(err).
				WithField("event_id", event.ID).
				WithField("event_type", event.EventType).
				Error("Failed to process event")
		}
	}
	
	return nil
}

// processEvent processes a single outbox event
func (f *Forwarder) processEvent(ctx context.Context, event *Event) error {
	// Check if event should be processed now
	if !event.ShouldProcess() {
		return nil
	}
	
	// Increment attempts
	if err := f.repo.IncrementAttempts(ctx, event.ID); err != nil {
		return fmt.Errorf("failed to increment attempts: %w", err)
	}
	
	// Publish to message broker
	// Topic name is typically the event type (e.g., "user.created")
	topic := event.EventType
	key := event.AggregateID // Use aggregate ID as message key for partitioning
	
	err := f.publisher.Publish(ctx, topic, key, event.Payload)
	if err != nil {
		// Check if we should retry
		if event.CanRetry() {
			// Calculate next retry time with exponential backoff
			nextRetry := f.scheduler.NextRetryTime(event.Attempts)
			event.ScheduledFor = &nextRetry
			
			f.logger.
				WithField("event_id", event.ID).
				WithField("attempt", event.Attempts).
				WithField("next_retry", nextRetry).
				Warn("Event publish failed, will retry")
			
			return nil // Don't mark as failed yet
		}
		
		// Max retries exceeded, mark as failed
		if err := f.repo.MarkFailed(ctx, event.ID, err); err != nil {
			return fmt.Errorf("failed to mark event as failed: %w", err)
		}
		
		f.logger.
			WithError(err).
			WithField("event_id", event.ID).
			WithField("event_type", event.EventType).
			Error("Event failed after max retries")
		
		return err
	}
	
	// Mark as published
	if err := f.repo.MarkPublished(ctx, event.ID); err != nil {
		return fmt.Errorf("failed to mark event as published: %w", err)
	}
	
	f.logger.
		WithField("event_id", event.ID).
		WithField("event_type", event.EventType).
		WithField("aggregate_id", event.AggregateID).
		Info("Event published successfully")
	
	return nil
}