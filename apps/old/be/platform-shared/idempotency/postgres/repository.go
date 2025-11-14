package postgres

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
	"skillsier.dev/platform-shared/idempotency"
)

// StringMap is a custom type for storing map[string]string in PostgreSQL
type StringMap map[string]string

// Scan implements the sql.Scanner interface
func (s *StringMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan StringMap")
	}
	return json.Unmarshal(bytes, s)
}

// Value implements the driver.Valuer interface
func (s StringMap) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// IdempotencyRecord is the GORM model for idempotency records
type IdempotencyRecord struct {
	Key             string    `gorm:"type:varchar(255);primaryKey"`
	StatusCode      int       `gorm:"not null"`
	ResponseBody    []byte    `gorm:"type:bytea"`
	ResponseHeaders StringMap `gorm:"type:jsonb"`
	CreatedAt       time.Time `gorm:"not null"`
	ExpiresAt       time.Time `gorm:"not null;index"`
}

// TableName specifies the table name for GORM
func (IdempotencyRecord) TableName() string {
	return "idempotency_records"
}

// Repository implements idempotency.Repository using PostgreSQL
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new PostgreSQL idempotency repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Get retrieves an idempotency record by key
func (r *Repository) Get(ctx context.Context, key string) (*idempotency.Record, error) {
	var model IdempotencyRecord
	err := r.db.WithContext(ctx).Where("key = ?", key).First(&model).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	
	return &idempotency.Record{
		Key:             model.Key,
		StatusCode:      model.StatusCode,
		ResponseBody:    model.ResponseBody,
		ResponseHeaders: model.ResponseHeaders,
		CreatedAt:       model.CreatedAt,
		ExpiresAt:       model.ExpiresAt,
	}, nil
}

// Create stores a new idempotency record
func (r *Repository) Create(ctx context.Context, record *idempotency.Record) error {
	model := &IdempotencyRecord{
		Key:             record.Key,
		StatusCode:      record.StatusCode,
		ResponseBody:    record.ResponseBody,
		ResponseHeaders: record.ResponseHeaders,
		CreatedAt:       record.CreatedAt,
		ExpiresAt:       record.ExpiresAt,
	}
	
	return r.db.WithContext(ctx).Create(model).Error
}

// Delete removes an idempotency record
func (r *Repository) Delete(ctx context.Context, key string) error {
	return r.db.WithContext(ctx).
		Where("key = ?", key).
		Delete(&IdempotencyRecord{}).Error
}

// DeleteExpired removes expired idempotency records
func (r *Repository) DeleteExpired(ctx context.Context) error {
	return r.db.WithContext(ctx).
		Where("expires_at < ?", time.Now()).
		Delete(&IdempotencyRecord{}).Error
}

// AutoMigrate runs database migrations for the idempotency table
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&IdempotencyRecord{})
}