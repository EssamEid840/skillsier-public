package portfolio

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Portfolio struct {
	ID          uuid.UUID        `gorm:"type:uuid;primary_key" json:"id"`
	UserID      uuid.UUID        `gorm:"type:uuid;not null;index" json:"user_id"`
	Title       string           `gorm:"type:varchar(200);not null" json:"title"`
	Description string           `gorm:"type:text" json:"description"`
	ProjectURL  string           `gorm:"type:varchar(500)" json:"project_url"`
	Images      []PortfolioImage `gorm:"foreignKey:PortfolioID" json:"images,omitempty"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	DeletedAt   gorm.DeletedAt   `gorm:"index" json:"-"`
}

func (Portfolio) TableName() string {
	return "portfolios"
}

func (p *Portfolio) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

func (p *Portfolio) Validate() error {
	if p.UserID == uuid.Nil {
		return ErrInvalidUserID
	}
	if p.Title == "" {
		return ErrTitleRequired
	}
	return nil
}