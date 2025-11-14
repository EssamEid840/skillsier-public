package portfolio

import "users-be/internal/domain/portfolio"

func ToResponseDTO(p *portfolio.Portfolio) *PortfolioResponseDTO {
	if p == nil {
		return nil
	}
	images := make([]*PortfolioImageDTO, len(p.Images))
	for i, img := range p.Images {
		images[i] = &PortfolioImageDTO{
			ID:           img.ID,
			ImageURL:     img.ImageURL,
			Caption:      img.Caption,
			DisplayOrder: img.DisplayOrder,
		}
	}
	return &PortfolioResponseDTO{
		ID:          p.ID,
		UserID:      p.UserID,
		Title:       p.Title,
		Description: p.Description,
		ProjectURL:  p.ProjectURL,
		Images:      images,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func ToListResponse(portfolios []*portfolio.Portfolio) *ListPortfoliosResponseDTO {
	dtos := make([]*PortfolioResponseDTO, len(portfolios))
	for i, p := range portfolios {
		dtos[i] = ToResponseDTO(p)
	}
	return &ListPortfoliosResponseDTO{Portfolios: dtos, Total: len(portfolios)}
}