package review

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	ContractID   uuid.UUID      `gorm:"type:uuid;not null;index" json:"contract_id"`
	ReviewerID   uuid.UUID      `gorm:"type:uuid;not null;index" json:"reviewer_id"`
	RevieweeID   uuid.UUID      `gorm:"type:uuid;not null;index" json:"reviewee_id"`
	Rating       int            `gorm:"not null" json:"rating"`
	Comment      string         `gorm:"type:text" json:"comment"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Review) TableName() string {
	return "reviews"
}

func (r *Review) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}

func (r *Review) Validate() error {
	if r.Rating < 1 || r.Rating > 5 {
		return ErrInvalidRating
	}
	if r.ReviewerID == uuid.Nil || r.RevieweeID == uuid.Nil {
		return ErrInvalidUserID
	}
	return nil
}
