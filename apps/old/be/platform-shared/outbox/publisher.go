package outbox

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

// Publisher publishes events within a database transaction
// This ensures atomicity - either both the business operation and event creation succeed, or both fail
type Publisher struct {
	repo Repository
}

// NewPublisher creates a new outbox publisher
func NewPublisher(repo Repository) *Publisher {
	return &Publisher{
		repo: repo,
	}
}

// Publish creates an outbox event within the current transaction
// The event will be picked up by the forwarder and published to the message broker
func (p *Publisher) Publish(ctx context.Context, event *Event) error {
	return p.repo.Create(ctx, event)
}

// PublishWithTx publishes an event within a GORM transaction
func (p *Publisher) PublishWithTx(tx *gorm.DB, event *Event) error {
	return p.repo.Create(tx.Statement.Context, event)
}

// PublishWithSQLTx publishes an event within a sql.Tx transaction
func (p *Publisher) PublishWithSQLTx(ctx context.Context, tx *sql.Tx, event *Event) error {
	return p.repo.Create(ctx, event)
}

// PublishMultiple publishes multiple events within a transaction
func (p *Publisher) PublishMultiple(ctx context.Context, events []*Event) error {
	for _, event := range events {
		if err := p.repo.Create(ctx, event); err != nil {
			return err
		}
	}
	return nil
}