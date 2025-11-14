package idempotency

import (
	"context"
	"time"
)

// Handler provides business logic for idempotency
type Handler struct {
	repo Repository
	ttl  time.Duration
}

// NewHandler creates a new idempotency handler
func NewHandler(repo Repository, ttl time.Duration) *Handler {
	if ttl == 0 {
		ttl = DefaultTTL
	}
	return &Handler{
		repo: repo,
		ttl:  ttl,
	}
}

// CheckAndStore checks if a request with the given key was already processed
// Returns the cached record if found, nil otherwise
func (h *Handler) CheckAndStore(ctx context.Context, key string, statusCode int, body []byte, headers map[string]string) (*Record, error) {
	// Check if already exists
	existing, err := h.repo.Get(ctx, key)
	if err == nil && existing != nil && !existing.IsExpired() {
		return existing, nil
	}

	// Create new record
	record := NewRecord(key, statusCode, body, headers, h.ttl)
	if err := h.repo.Create(ctx, record); err != nil {
		return nil, err
	}

	return nil, nil
}

// Get retrieves an idempotency record by key
func (h *Handler) Get(ctx context.Context, key string) (*Record, error) {
	return h.repo.Get(ctx, key)
}

// Delete removes an idempotency record
func (h *Handler) Delete(ctx context.Context, key string) error {
	return h.repo.Delete(ctx, key)
}

// CleanupExpired removes expired idempotency records
// Should be called periodically (e.g., via cron job)
func (h *Handler) CleanupExpired(ctx context.Context) (int64, error) {
	return h.repo.DeleteExpired(ctx, time.Now())
}