package portfolio

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PortfolioImage struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	PortfolioID uuid.UUID      `gorm:"type:uuid;not null;index" json:"portfolio_id"`
	ImageURL    string         `gorm:"type:varchar(500);not null" json:"image_url"`
	Caption     string         `gorm:"type:varchar(200)" json:"caption"`
	DisplayOrder int           `gorm:"default:0" json:"display_order"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (PortfolioImage) TableName() string {
	return "portfolio_images"
}

func (pi *PortfolioImage) BeforeCreate(tx *gorm.DB) error {
	if pi.ID == uuid.Nil {
		pi.ID = uuid.New()
	}
	return nil
}