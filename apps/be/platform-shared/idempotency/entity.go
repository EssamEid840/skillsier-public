package idempotency

import "time"

// Record represents an idempotency record for HTTP requests
type Record struct {
	// Key is the idempotency key (from Idempotency-Key header)
	Key string
	
	// StatusCode is the HTTP status code of the original response
	StatusCode int
	
	// ResponseBody is the response body of the original request
	ResponseBody []byte
	
	// ResponseHeaders are the response headers of the original request
	ResponseHeaders map[string]string
	
	// CreatedAt is when the record was created
	CreatedAt time.Time
	
	// ExpiresAt is when the record expires and can be deleted
	ExpiresAt time.Time
}

// NewRecord creates a new idempotency record
func NewRecord(key string, statusCode int, body []byte, headers map[string]string, ttl time.Duration) *Record {
	now := time.Now()
	return &Record{
		Key:             key,
		StatusCode:      statusCode,
		ResponseBody:    body,
		ResponseHeaders: headers,
		CreatedAt:       now,
		ExpiresAt:       now.Add(ttl),
	}
}

// IsExpired returns true if the record has expired
func (r *Record) IsExpired() bool {
	return time.Now().After(r.ExpiresAt)
}