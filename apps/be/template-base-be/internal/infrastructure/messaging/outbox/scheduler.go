package outbox

import (
	"math"
	"time"

	"<module>/internal/config"
)

// Scheduler handles retry scheduling with exponential backoff
type Scheduler struct {
	config config.OutboxConfig
}

// NewScheduler creates a new retry scheduler
func NewScheduler(cfg config.OutboxConfig) *Scheduler {
	return &Scheduler{
		config: cfg,
	}
}

// CalculateNextRetry calculates the next retry time using exponential backoff
func (s *Scheduler) CalculateNextRetry(retryCount int) time.Time {
	// Exponential backoff formula: base * (2 ^ retryCount)
	// With jitter to prevent thundering herd
	backoff := s.calculateBackoff(retryCount)
	
	// Add jitter (±10% randomness)
	jitter := s.calculateJitter(backoff)
	finalBackoff := backoff + jitter

	// Cap at max backoff
	if finalBackoff > s.config.RetryBackoffMax {
		finalBackoff = s.config.RetryBackoffMax
	}

	return time.Now().Add(finalBackoff)
}

// calculateBackoff calculates the base backoff duration
func (s *Scheduler) calculateBackoff(retryCount int) time.Duration {
	// Exponential backoff: base * (2 ^ retryCount)
	multiplier := math.Pow(2, float64(retryCount))
	backoff := time.Duration(float64(s.config.RetryBackoffBase) * multiplier)
	
	return backoff
}

// calculateJitter adds random jitter to prevent thundering herd
func (s *Scheduler) calculateJitter(backoff time.Duration) time.Duration {
	// Add ±10% jitter
	jitterRange := float64(backoff) * 0.1
	jitter := (time.Now().UnixNano() % int64(jitterRange*2)) - int64(jitterRange)
	
	return time.Duration(jitter)
}

// ShouldRetry determines if an event should be retried
func (s *Scheduler) ShouldRetry(retryCount int) bool {
	return retryCount < s.config.MaxRetries
}

// GetRetrySchedule returns the full retry schedule for debugging
func (s *Scheduler) GetRetrySchedule() []time.Duration {
	schedule := make([]time.Duration, s.config.MaxRetries)
	
	for i := 0; i < s.config.MaxRetries; i++ {
		backoff := s.calculateBackoff(i)
		if backoff > s.config.RetryBackoffMax {
			backoff = s.config.RetryBackoffMax
		}
		schedule[i] = backoff
	}
	
	return schedule
}

// PrintRetrySchedule prints the retry schedule for debugging
func (s *Scheduler) PrintRetrySchedule() {
	schedule := s.GetRetrySchedule()
	
	println("Retry Schedule:")
	for i, backoff := range schedule {
		println("  Retry", i+1, "after", backoff.String())
	}
}