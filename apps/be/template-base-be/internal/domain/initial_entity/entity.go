package initial_entity

import (
	"time"

	"github.com/google/uuid"
)

// InitialEntity represents the core domain entity for this microservice
// This is a template entity - replace with your actual domain entity
type InitialEntity struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	Name        string     `gorm:"type:varchar(255);not null;index" json:"name"`
	Description string     `gorm:"type:text" json:"description"`
	Status      Status     `gorm:"type:varchar(50);not null;default:'active';index" json:"status"`
	OwnerID     uuid.UUID  `gorm:"type:uuid;not null;index" json:"owner_id"`
	Metadata    Metadata   `gorm:"embedded;embeddedPrefix:metadata_" json:"metadata"`
	CreatedAt   time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at,omitempty"` // Soft delete
}

// Metadata contains additional entity metadata
type Metadata struct {
	Tags       []string          `gorm:"type:jsonb" json:"tags,omitempty"`
	Properties map[string]string `gorm:"type:jsonb" json:"properties,omitempty"`
	Version    int               `gorm:"not null;default:1" json:"version"`
}

// TableName returns the table name for GORM
func (InitialEntity) TableName() string {
	return "initial_entities"
}

// BeforeCreate is a GORM hook called before creating a new entity
func (e *InitialEntity) BeforeCreate() error {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	
	now := time.Now()
	e.CreatedAt = now
	e.UpdatedAt = now
	
	// Default status if not set
	if e.Status == "" {
		e.Status = StatusActive
	}
	
	// Initialize metadata version
	if e.Metadata.Version == 0 {
		e.Metadata.Version = 1
	}
	
	return nil
}

// BeforeUpdate is a GORM hook called before updating an entity
func (e *InitialEntity) BeforeUpdate() error {
	e.UpdatedAt = time.Now()
	e.Metadata.Version++
	return nil
}

// Validate performs domain-level validation
func (e *InitialEntity) Validate() error {
	if e.Name == "" {
		return ErrNameRequired
	}
	
	if len(e.Name) < 3 {
		return ErrNameTooShort
	}
	
	if len(e.Name) > 255 {
		return ErrNameTooLong
	}
	
	if e.OwnerID == uuid.Nil {
		return ErrOwnerIDRequired
	}
	
	if !e.Status.IsValid() {
		return ErrInvalidStatus
	}
	
	return nil
}

// CanTransitionTo checks if the entity can transition to the given status
func (e *InitialEntity) CanTransitionTo(newStatus Status) bool {
	// Define allowed status transitions
	allowedTransitions := map[Status][]Status{
		StatusActive: {StatusInactive, StatusArchived},
		StatusInactive: {StatusActive, StatusArchived},
		StatusArchived: {}, // Cannot transition from archived
	}
	
	allowed, exists := allowedTransitions[e.Status]
	if !exists {
		return false
	}
	
	for _, s := range allowed {
		if s == newStatus {
			return true
		}
	}
	
	return false
}

// IsActive returns true if the entity is active
func (e *InitialEntity) IsActive() bool {
	return e.Status == StatusActive
}

// IsDeleted returns true if the entity is soft-deleted
func (e *InitialEntity) IsDeleted() bool {
	return e.DeletedAt != nil
}

// SoftDelete marks the entity as deleted
func (e *InitialEntity) SoftDelete() {
	now := time.Now()
	e.DeletedAt = &now
	e.UpdatedAt = now
}

// Restore restores a soft-deleted entity
func (e *InitialEntity) Restore() {
	e.DeletedAt = nil
	e.UpdatedAt = time.Now()
}

// AddTag adds a tag to the entity's metadata
func (e *InitialEntity) AddTag(tag string) {
	if e.Metadata.Tags == nil {
		e.Metadata.Tags = []string{}
	}
	
	// Avoid duplicates
	for _, t := range e.Metadata.Tags {
		if t == tag {
			return
		}
	}
	
	e.Metadata.Tags = append(e.Metadata.Tags, tag)
	e.UpdatedAt = time.Now()
}

// RemoveTag removes a tag from the entity's metadata
func (e *InitialEntity) RemoveTag(tag string) {
	if e.Metadata.Tags == nil {
		return
	}
	
	filtered := []string{}
	for _, t := range e.Metadata.Tags {
		if t != tag {
			filtered = append(filtered, t)
		}
	}
	
	e.Metadata.Tags = filtered
	e.UpdatedAt = time.Now()
}

// SetProperty sets a property in the entity's metadata
func (e *InitialEntity) SetProperty(key, value string) {
	if e.Metadata.Properties == nil {
		e.Metadata.Properties = make(map[string]string)
	}
	
	e.Metadata.Properties[key] = value
	e.UpdatedAt = time.Now()
}

// GetProperty gets a property from the entity's metadata
func (e *InitialEntity) GetProperty(key string) (string, bool) {
	if e.Metadata.Properties == nil {
		return "", false
	}
	
	value, exists := e.Metadata.Properties[key]
	return value, exists
}

// Clone creates a deep copy of the entity
func (e *InitialEntity) Clone() *InitialEntity {
	clone := *e
	
	// Deep copy slices and maps
	if e.Metadata.Tags != nil {
		clone.Metadata.Tags = make([]string, len(e.Metadata.Tags))
		copy(clone.Metadata.Tags, e.Metadata.Tags)
	}
	
	if e.Metadata.Properties != nil {
		clone.Metadata.Properties = make(map[string]string)
		for k, v := range e.Metadata.Properties {
			clone.Metadata.Properties[k] = v
		}
	}
	
	return &clone
}