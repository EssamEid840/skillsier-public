package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents the user domain entity
// This stores additional information beyond what Keycloak manages
type User struct {
	// ID is the internal database ID
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	
	// KeycloakID links this user to their Keycloak identity
	// This is the "sub" claim from Keycloak JWT tokens
	KeycloakID string `gorm:"type:varchar(255);uniqueIndex;not null" json:"keycloak_id"`
	
	// Basic Info (synced from Keycloak)
	Username string `gorm:"type:varchar(255);uniqueIndex;not null" json:"username"`
	Email    string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	
	// Additional Profile Information (managed by this service)
	FirstName   string  `gorm:"type:varchar(100)" json:"first_name,omitempty"`
	LastName    string  `gorm:"type:varchar(100)" json:"last_name,omitempty"`
	PhoneNumber *string `gorm:"type:varchar(20)" json:"phone_number,omitempty"`
	Bio         *string `gorm:"type:text" json:"bio,omitempty"`
	
	// Professional Information (for freelancer/client profiles)
	ProfileType    string  `gorm:"type:varchar(50)" json:"profile_type,omitempty"` // freelancer, client, both
	Profession     *string `gorm:"type:varchar(100)" json:"profession,omitempty"`
	HourlyRate     *float64 `gorm:"type:decimal(10,2)" json:"hourly_rate,omitempty"`
	AvailableHours *int    `gorm:"type:integer" json:"available_hours,omitempty"`
	
	// Location
	Country *string `gorm:"type:varchar(100)" json:"country,omitempty"`
	City    *string `gorm:"type:varchar(100)" json:"city,omitempty"`
	
	// Status
	IsActive       bool `gorm:"default:true" json:"is_active"`
	EmailVerified  bool `gorm:"default:false" json:"email_verified"`
	ProfileComplete bool `gorm:"default:false" json:"profile_complete"`
	
	// Timestamps
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName specifies the table name for GORM
func (User) TableName() string {
	return "users"
}

// BeforeCreate is a GORM hook that runs before creating a new user
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
	return nil
}

// BeforeUpdate is a GORM hook that runs before updating a user
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}

// IsProfileCompleted checks if user has filled all required profile fields
func (u *User) IsProfileCompleted() bool {
	return u.FirstName != "" &&
		u.LastName != "" &&
		u.ProfileType != "" &&
		u.Country != nil
}