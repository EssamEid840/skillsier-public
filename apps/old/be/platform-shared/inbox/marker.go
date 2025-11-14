package inbox

import "context"

// Marker marks messages as processed
type Marker struct {
	repo Repository
}

// NewMarker creates a new inbox marker
func NewMarker(repo Repository) *Marker {
	return &Marker{
		repo: repo,
	}
}

// MarkProcessed marks a message as processed
func (m *Marker) MarkProcessed(ctx context.Context, messageID, handler string, payload []byte) error {
	message := NewMessage(messageID, handler, payload)
	return m.repo.Create(ctx, message)
}