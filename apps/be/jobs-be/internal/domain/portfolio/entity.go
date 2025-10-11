package portfolio

import (
	"time"
	"github.com/google/uuid"
)

type Portfolio struct {
	ID          uuid.UUID        `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID      uuid.UUID        `gorm:"type:uuid;not null;index" json:"user_id"`
	Title       string           `gorm:"type:varchar(200);not null" json:"title"`
	Description *string          `gorm:"type:text" json:"description,omitempty"`
	ProjectURL  *string          `gorm:"type:text" json:"project_url,omitempty"`
	StartDate   *time.Time       `gorm:"type:date" json:"start_date,omitempty"`
	EndDate     *time.Time       `gorm:"type:date" json:"end_date,omitempty"`
	Images      []PortfolioImage `gorm:"foreignKey:PortfolioID" json:"images,omitempty"`
	CreatedAt   time.Time        `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time        `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type PortfolioImage struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	PortfolioID  uuid.UUID `gorm:"type:uuid;not null;index" json:"portfolio_id"`
	ImageURL     string    `gorm:"type:text;not null" json:"image_url"`
	DisplayOrder int       `gorm:"type:int;default:0" json:"display_order"`
	CreatedAt    time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (Portfolio) TableName() string {
	return "portfolios"
}

func (PortfolioImage) TableName() string {
	return "portfolio_images"
}

func (p *Portfolio) Validate() error {
	if p.Title == "" {
		return ErrTitleRequired
	}
	return nil
}