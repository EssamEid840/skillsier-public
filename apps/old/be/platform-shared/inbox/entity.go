package inbox

import "time"

// Message represents a processed message to prevent duplicate processing
type Message struct {
	// ID is the unique message ID (from Kafka, RabbitMQ, etc.)
	ID string
	
	// Handler is the name of the handler that processed this message
	Handler string
	
	// ProcessedAt is when the message was successfully processed
	ProcessedAt time.Time
	
	// Payload stores the original message payload (optional, for debugging)
	Payload []byte
}

// NewMessage creates a new inbox message record
func NewMessage(id, handler string, payload []byte) *Message {
	return &Message{
		ID:          id,
		Handler:     handler,
		ProcessedAt: time.Now(),
		Payload:     payload,
	}
}