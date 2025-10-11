package outbox

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"contracts-be/internal/domain/outbox"
	"contracts-be/internal/infrastructure/messaging/kafka"
)

// Processor handles the outbox pattern implementation
// It periodically polls the outbox table and publishes pending events to Kafka
type Processor struct {
	outboxRepo    outbox.Repository
	kafkaProducer *kafka.Producer
	pollInterval  time.Duration
	batchSize     int
	maxRetries    int
}

// NewProcessor creates a new outbox processor
func NewProcessor(
	outboxRepo outbox.Repository,
	kafkaProducer *kafka.Producer,
	pollInterval time.Duration,
	batchSize int,
	maxRetries int,
) *Processor {
	return &Processor{
		outboxRepo:    outboxRepo,
		kafkaProducer: kafkaProducer,
		pollInterval:  pollInterval,
		batchSize:     batchSize,
		maxRetries:    maxRetries,
	}
}

// Start begins processing outbox events
// This runs in a continuous loop, polling for pending events
func (p *Processor) Start(ctx context.Context) error {
	log.Println("Starting outbox processor...")

	ticker := time.NewTicker(p.pollInterval)
	defer ticker.Stop()

	// Process immediately on start
	if err := p.processBatch(ctx); err != nil {
		log.Printf("Error processing initial batch: %v", err)
	}

	// Then process on each tick
	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping outbox processor...")
			return ctx.Err()
		case <-ticker.C:
			if err := p.processBatch(ctx); err != nil {
				log.Printf("Error processing batch: %v", err)
				// Continue processing even on error
			}
		}
	}
}

// processBatch retrieves and processes a batch of pending events
func (p *Processor) processBatch(ctx context.Context) error {
	// Retrieve pending events from the outbox
	events, err := p.outboxRepo.FindPendingEvents(ctx, p.batchSize)
	if err != nil {
		return fmt.Errorf("failed to find pending events: %w", err)
	}

	if len(events) == 0 {
		// No pending events, nothing to do
		return nil
	}

	log.Printf("Processing %d pending events", len(events))

	// Process each event
	for _, event := range events {
		if err := p.processEvent(ctx, event); err != nil {
			log.Printf("Error processing event %s: %v", event.ID, err)
			// Continue with next event
		}
	}

	return nil
}

// processEvent publishes a single event to Kafka
func (p *Processor) processEvent(ctx context.Context, event *outbox.Event) error {
	// Check if we should retry this event
	if event.Status == outbox.EventStatusFailed {
		if !event.CanRetry(p.maxRetries) {
			log.Printf("Event %s has exceeded max retries or is not ready for retry", event.ID)
			return nil
		}
	}

	// Create the message payload
	// The payload combines the event data with metadata
	message := map[string]interface{}{
		"event_id":       event.ID,
		"event_type":     event.EventType,
		"aggregate_id":   event.AggregateID,
		"aggregate_type": event.AggregateType,
		"version":        event.Version,
		"timestamp":      event.CreatedAt,
		"payload":        json.RawMessage(event.Payload),
	}

	if event.Metadata != nil {
		message["metadata"] = json.RawMessage(event.Metadata)
	}

	// Serialize message to JSON
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return p.handlePublishError(ctx, event, fmt.Errorf("failed to marshal message: %w", err))
	}

	// Publish to Kafka
	// Use aggregate_id as the message key for partitioning
	key := []byte(event.AggregateID)
	if err := p.kafkaProducer.Publish(event.Topic, key, messageBytes); err != nil {
		return p.handlePublishError(ctx, event, fmt.Errorf("failed to publish to Kafka: %w", err))
	}

	// Mark event as published
	event.MarkAsPublished()
	if err := p.outboxRepo.Update(ctx, event); err != nil {
		log.Printf("Failed to mark event %s as published: %v", event.ID, err)
		// Event was published but we couldn't update the status
		// This might cause duplicate publishing, but that's better than losing events
		return err
	}

	log.Printf("Successfully published event %s to topic %s", event.ID, event.Topic)
	return nil
}

// handlePublishError handles errors that occur during event publishing
func (p *Processor) handlePublishError(ctx context.Context, event *outbox.Event, err error) error {
	log.Printf("Failed to publish event %s: %v", event.ID, err)

	// Mark event as failed
	event.MarkAsFailed(err.Error())

	// Update event in database
	if updateErr := p.outboxRepo.Update(ctx, event); updateErr != nil {
		log.Printf("Failed to update failed event %s: %v", event.ID, updateErr)
		return fmt.Errorf("publish error: %w, update error: %v", err, updateErr)
	}

	return err
}

// CleanupPublishedEvents removes old published events from the outbox
// This should be called periodically to prevent the outbox table from growing indefinitely
func (p *Processor) CleanupPublishedEvents(ctx context.Context, olderThanDays int) error {
	deleted, err := p.outboxRepo.DeletePublished(ctx, olderThanDays)
	if err != nil {
		return fmt.Errorf("failed to cleanup published events: %w", err)
	}

	if deleted > 0 {
		log.Printf("Cleaned up %d published events older than %d days", deleted, olderThanDays)
	}

	return nil
}
