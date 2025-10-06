package outbox_dispatcher

import (
	"context"
	"os"
	"strconv"
	"time"


	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"

	"users.go/m/internal/infrastructure/event"
	"users.go/m/internal/infrastructure/event/common"
	"users.go/m/internal/models"
	"users.go/m/internal/repository/outbox"
)

type OutboxDispatcher struct {
	eventPublisher event.EventPublisher
	outboxRepo     *outbox.OutboxRepo
	interval       time.Duration
	batchSize      int
}

// --- Defaults (can be overridden via env) ---
const (
	defaultIntervalSeconds = 5
	defaultBatchSize       = 64
)

// Backward-compatible constructor (matches setup.go usage):
//   outboxDispatcher := outbox_dispatcher.NewOutboxDispatcher(eventPublisher, outboxRepo)
// Reads optional env overrides:
//   APP_OUTBOX__FETCH_INTERVAL_SECONDS (int, seconds)
//   APP_OUTBOX__BATCH_SIZE (int)
func NewOutboxDispatcher(pub event.EventPublisher, repo *outbox.OutboxRepo) *OutboxDispatcher {
	intervalSec := getEnvInt("APP_OUTBOX__FETCH_INTERVAL_SECONDS", defaultIntervalSeconds)
	batch := getEnvInt("APP_OUTBOX__BATCH_SIZE", defaultBatchSize)
	return &OutboxDispatcher{
		eventPublisher: pub,
		outboxRepo:     repo,
		interval:       time.Duration(intervalSec) * time.Second,
		batchSize:      batch,
	}
}

// Optional explicit constructor if you ever want to set values in code.
func NewOutboxDispatcherWith(pub event.EventPublisher, repo *outbox.OutboxRepo, interval time.Duration, batchSize int) *OutboxDispatcher {
	return &OutboxDispatcher{
		eventPublisher: pub,
		outboxRepo:     repo,
		interval:       interval,
		batchSize:      batchSize,
	}
}

func (o *OutboxDispatcher) Start(ctx context.Context) {
	t := time.NewTicker(o.interval)
	defer t.Stop()

	log.Info().Dur("interval", o.interval).Int("batch", o.batchSize).Msg("Starting outbox dispatcher")

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			o.processOnce(ctx)
		}
	}
}

// Backward-compat shim so existing setup.go can call Execute()
func (o *OutboxDispatcher) Execute() {
	o.Start(context.Background())
}

func (o *OutboxDispatcher) processOnce(ctx context.Context) {
	log.Info().Msg("Fetching outboxes")

	events, err := o.outboxRepo.FetchUnsent(ctx)
	if err != nil {
		log.Err(err).Msg("failed to fetch unsent outbox")
		return
	}
	if len(events) == 0 {
		return
	}

	// If you want to honor batchSize strictly, slice here:
	if len(events) > o.batchSize {
		events = events[:o.batchSize]
	}

	for _, ev := range events {
		o.publishOne(ctx, ev)
	}
}

func (o *OutboxDispatcher) publishOne(ctx context.Context, eventOutbox *models.Outbox) {
	log.Info().Str("id", eventOutbox.ID).Msg("Processing outbox...")

	// Map aggregate type to topic
	topic, ok := common.AggregateTypeToTopic[eventOutbox.AggregateType]
	if !ok {
		log.Warn().Str("aggregate_type", eventOutbox.AggregateType).Msg("no topic mapping for aggregate type")
		return
	}

	// IMPORTANT: partition key = aggregate_id (per-aggregate ordering)
	partitionKey := eventOutbox.AggregateID

	// Your payload now contains the event envelope JSON (backward compatible for consumers).
	body := []byte(eventOutbox.Payload)

	if err := o.eventPublisher.Publish(ctx, topic, partitionKey, body); err != nil {
		log.Err(err).
			Str("topic", topic).
			Str("key", partitionKey).
			Msg("failed to publish outbox")
		return
	}

	// Mark as sent (persist via repo)
	eventOutbox.SentAt = null.TimeFrom(time.Now().UTC())
	if err := o.outboxRepo.Save(ctx, nil, eventOutbox); err != nil {
		log.Err(err).Msg("failed to save outbox")
		return
	}

	log.Info().Str("id", eventOutbox.ID).Msg("outbox processed")
}

func getEnvInt(name string, def int) int {
	v := os.Getenv(name)
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil || n <= 0 {
		return def
	}
	return n
}
