package outbox

import (
	"math"
	"time"
)

// Scheduler handles retry scheduling with exponential backoff
type Scheduler struct {
	config *Config
}

// NewScheduler creates a new scheduler
func NewScheduler(config *Config) *Scheduler {
	return &Scheduler{
		config: config,
	}
}

// NextRetryTime calculates the next retry time using exponential backoff
// Formula: baseDelay * (2 ^ attempt) + jitter
func (s *Scheduler) NextRetryTime(attempt int) time.Time {
	// Calculate delay with exponential backoff
	delay := float64(s.config.RetryBaseDelay) * math.Pow(2, float64(attempt))
	
	// Cap at max delay
	if delay > float64(s.config.RetryMaxDelay) {
		delay = float64(s.config.RetryMaxDelay)
	}
	
	// Add jitter (random value between 0 and 1 second)
	jitter := time.Duration(float64(time.Second) * float64(time.Now().UnixNano()%1000) / 1000)
	
	return time.Now().Add(time.Duration(delay) + jitter)
}

// ShouldRetry determines if an event should be retried based on attempt count
func (s *Scheduler) ShouldRetry(attempt, maxAttempts int) bool {
	return attempt < maxAttempts
}