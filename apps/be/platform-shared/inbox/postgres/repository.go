package postgres

import (
	"context"
	"time"

	"gorm.io/gorm"
	"skillsier.dev/platform-shared/inbox"
)

// InboxMessage is the GORM model for inbox messages
type InboxMessage struct {
	ID          string    `gorm:"type:varchar(255);not null"`
	Handler     string    `gorm:"type:varchar(255);not null"`
	ProcessedAt time.Time `gorm:"not null;index"`
	Payload     []byte    `gorm:"type:jsonb"`
}

// TableName specifies the table name for GORM
func (InboxMessage) TableName() string {
	return "inbox_messages"
}

// Repository implements inbox.Repository using PostgreSQL
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new PostgreSQL inbox repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Exists checks if a message has already been processed
func (r *Repository) Exists(ctx context.Context, messageID, handler string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&InboxMessage{}).
		Where("id = ? AND handler = ?", messageID, handler).
		Count(&count).Error
	
	if err != nil {
		return false, err
	}
	
	return count > 0, nil
}

// Create records a processed message
func (r *Repository) Create(ctx context.Context, message *inbox.Message) error {
	model := &InboxMessage{
		ID:          message.ID,
		Handler:     message.Handler,
		ProcessedAt: message.ProcessedAt,
		Payload:     message.Payload,
	}
	
	return r.db.WithContext(ctx).Create(model).Error
}

// Delete removes a message record
func (r *Repository) Delete(ctx context.Context, messageID, handler string) error {
	return r.db.WithContext(ctx).
		Where("id = ? AND handler = ?", messageID, handler).
		Delete(&InboxMessage{}).Error
}

// DeleteOld removes message records older than the specified duration
func (r *Repository) DeleteOld(ctx context.Context, olderThanDays int) error {
	cutoff := time.Now().AddDate(0, 0, -olderThanDays)
	
	return r.db.WithContext(ctx).
		Where("processed_at < ?", cutoff).
		Delete(&InboxMessage{}).Error
}

// AutoMigrate runs database migrations for the inbox table
func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&InboxMessage{}); err != nil {
		return err
	}
	
	// Add composite unique index
	return db.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS idx_inbox_messages_id_handler 
		ON inbox_messages(id, handler)
	`).Error
}