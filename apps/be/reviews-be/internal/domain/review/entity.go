package review

import (
	"time"
	"github.com/google/uuid"
)

type Review struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	ContractID   uuid.UUID `gorm:"type:uuid;not null;index" json:"contract_id"`
	ReviewerID   uuid.UUID `gorm:"type:uuid;not null;index" json:"reviewer_id"`   // Who gave the review
	RevieweeID   uuid.UUID `gorm:"type:uuid;not null;index" json:"reviewee_id"`   // Who received the review
	Rating       int       `gorm:"type:int;not null" json:"rating"`                // 1-5
	Comment      *string   `gorm:"type:text" json:"comment,omitempty"`
	// Skill ratings (optional detailed breakdown)
	QualityOfWork      *int `gorm:"type:int" json:"quality_of_work,omitempty"`      // 1-5
	Communication      *int `gorm:"type:int" json:"communication,omitempty"`        // 1-5
	Professionalism    *int `gorm:"type:int" json:"professionalism,omitempty"`      // 1-5
	DeadlineAdherence  *int `gorm:"type:int" json:"deadline_adherence,omitempty"`   // 1-5
	CreatedAt    time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Review) TableName() string {
	return "reviews"
}

func (r *Review) Validate() error {
	if r.Rating < 1 || r.Rating > 5 {
		return ErrInvalidRating
	}
	if r.QualityOfWork != nil && (*r.QualityOfWork < 1 || *r.QualityOfWork > 5) {
		return ErrInvalidDetailedRating
	}
	if r.Communication != nil && (*r.Communication < 1 || *r.Communication > 5) {
		return ErrInvalidDetailedRating
	}
	if r.Professionalism != nil && (*r.Professionalism < 1 || *r.Professionalism > 5) {
		return ErrInvalidDetailedRating
	}
	if r.DeadlineAdherence != nil && (*r.DeadlineAdherence < 1 || *r.DeadlineAdherence > 5) {
		return ErrInvalidDetailedRating
	}
	return nil
}