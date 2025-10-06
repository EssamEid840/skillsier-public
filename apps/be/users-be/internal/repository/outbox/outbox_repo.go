package outbox

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"

	"users.go/m/internal/models"
)

type OutboxRepo struct {
	db *sql.DB
}

func NewOutboxRepo(db *sql.DB) *OutboxRepo {
	return &OutboxRepo{db: db}
}

// ---- Event envelope (embedded in the existing payload column) ----

const currentSchemaVersion = 1

type eventEnvelope struct {
	ID            string            `json:"id"`
	AggregateType string            `json:"aggregate_type"`
	AggregateID   string            `json:"aggregate_id"`
	Version       int               `json:"version"`
	Type          string            `json:"type"`
	OccurredAt    time.Time         `json:"occurred_at"`
	SchemaVersion int               `json:"schema_version"`
	Metadata      map[string]string `json:"metadata,omitempty"`
	CorrelationID *string           `json:"correlation_id,omitempty"`
	CausationID   *string           `json:"causation_id,omitempty"`

	// Back-compat fields for your current consumer shape
	Status    string          `json:"status,omitempty"`
	UserID    string          `json:"user_id,omitempty"`
	CreatedAt *time.Time      `json:"created_at,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
}

// Save upserts the outbox row and (1) wraps the original payload in an envelope,
// (2) computes a soft per-aggregate version, and (3) preserves back-compat fields.
func (r *OutboxRepo) Save(ctx context.Context, tx *sql.Tx, e *models.Outbox) error {
	// choose the executor (either tx or db); both implement boil.ContextExecutor
	var exec boil.ContextExecutor = r.db
	if tx != nil {
		exec = tx
	}

	// Compute next version (soft, stored inside the payload for now).
	nextVer, err := r.nextVersion(ctx, tx, e.AggregateID)
	if err != nil {
		log.Err(err).Str("aggregate_id", e.AggregateID).Msg("failed to compute next version")
		return err
	}

	var data json.RawMessage
	if len(e.Payload) > 0 {
		data = json.RawMessage(e.Payload)
	}

	env := eventEnvelope{
		ID:            e.ID,
		AggregateType: e.AggregateType,
		AggregateID:   e.AggregateID,
		Version:       nextVer,
		Type:          e.Type,
		SchemaVersion: currentSchemaVersion,
		Data:          data,
	}

	// OccurredAt default
	if e.OccurredAt.IsZero() {
		e.OccurredAt = null.TimeFrom(time.Now().UTC())
	}
	env.OccurredAt = e.OccurredAt.Time

	// Mirror some top-level fields (status, user_id, created_at) for current consumer.
	if len(data) > 0 {
		var tmp map[string]any
		if err := json.Unmarshal(data, &tmp); err == nil {
			if v, ok := tmp["user_id"].(string); ok {
				env.UserID = v
			}
			if v, ok := tmp["status"].(string); ok && v != "" {
				env.Status = v
			} else {
				env.Status = "created"
			}
			if v, ok := tmp["created_at"].(string); ok && v != "" {
				if t, perr := time.Parse(time.RFC3339, v); perr == nil {
					env.CreatedAt = &t
				}
			}
		}
	}
	if env.CreatedAt == nil {
		t := e.OccurredAt.Time
		env.CreatedAt = &t
	}

	// Wrap the original payload in the envelope (stored in the same column).
	wrapped, merr := json.Marshal(&env)
	if merr != nil {
		log.Err(merr).Msg("failed to marshal envelope")
		return merr
	}
	e.Payload = types.JSON(wrapped)

	// ✅ sqlboiler v4 Upsert signature:
	// Upsert(ctx, exec, updateOnConflict, conflictColumns, updateColumns, insertColumns, ...opts)
	if err := e.Upsert(ctx, exec, true, []string{models.OutboxColumns.ID}, boil.Infer(), boil.Infer()); err != nil {
		log.Err(err).Msg("failed to upsert outbox")
		return err
	}
	return nil
}

// nextVersion returns current count+1 for the aggregate (soft versioning inside payload).
func (r *OutboxRepo) nextVersion(ctx context.Context, tx *sql.Tx, aggregateID string) (int, error) {
	var exec boil.ContextExecutor = r.db
	if tx != nil {
		exec = tx
	}
	n, err := models.Outboxes(qm.Where("aggregate_id = ?", aggregateID)).Count(ctx, exec)
	if err != nil {
		return 0, err
	}
	return int(n) + 1, nil
}

func (repo *OutboxRepo) FetchUnsent(ctx context.Context) ([]*models.Outbox, error) {
	return models.Outboxes(
		models.OutboxWhere.SentAt.IsNull(),
	).All(ctx, repo.db)
}
