package portfolio

import (
	"github.com/google/uuid"
	"time"
)

type CreatePortfolioDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	ProjectURL  string `json:"project_url"`
}

type UpdatePortfolioDTO struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	ProjectURL  *string `json:"project_url,omitempty"`
}

type PortfolioImageDTO struct {
	ID           uuid.UUID `json:"id"`
	ImageURL     string    `json:"image_url"`
	Caption      string    `json:"caption"`
	DisplayOrder int       `json:"display_order"`
}

type PortfolioResponseDTO struct {
	ID          uuid.UUID            `json:"id"`
	UserID      uuid.UUID            `json:"user_id"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	ProjectURL  string               `json:"project_url"`
	Images      []*PortfolioImageDTO `json:"images"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

type ListPortfoliosResponseDTO struct {
	Portfolios []*PortfolioResponseDTO `json:"portfolios"`
	Total      int                     `json:"total"`
}

type UploadImageDTO struct {
	ImageURL     string `json:"image_url" binding:"required"`
	Caption      string `json:"caption"`
	DisplayOrder int    `json:"display_order"`
}
