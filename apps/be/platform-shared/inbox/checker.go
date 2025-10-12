package inbox

import "context"

// Checker checks if messages have already been processed
type Checker struct {
	repo Repository
}

// NewChecker creates a new inbox checker
func NewChecker(repo Repository) *Checker {
	return &Checker{
		repo: repo,
	}
}

// IsProcessed checks if a message has already been processed by the specified handler
func (c *Checker) IsProcessed(ctx context.Context, messageID, handler string) (bool, error) {
	return c.repo.Exists(ctx, messageID, handler)
}