package portfolio

import (
	"context"
	"encoding/json"
	"fmt"
	"users-be/internal/domain/outbox"
	"users-be/internal/domain/portfolio"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	portfolioRepo portfolio.Repository
	outboxRepo    outbox.Repository
	db            *gorm.DB
}

func NewService(portfolioRepo portfolio.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{portfolioRepo: portfolioRepo, outboxRepo: outboxRepo, db: db}
}

func (s *Service) CreatePortfolio(ctx context.Context, userID uuid.UUID, dto *CreatePortfolioDTO) (*PortfolioResponseDTO, error) {
	p := &portfolio.Portfolio{
		UserID:      userID,
		Title:       dto.Title,
		Description: dto.Description,
		ProjectURL:  dto.ProjectURL,
	}
	if err := p.Validate(); err != nil {
		return nil, err
	}
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.portfolioRepo.Create(ctx, p); err != nil {
			return err
		}
		event, _ := s.createPortfolioEvent("portfolio.added", p)
		return s.outboxRepo.Create(ctx, event)
	})
	if err != nil {
		return nil, err
	}
	return ToResponseDTO(p), nil
}

func (s *Service) GetAllPortfolios(ctx context.Context, userID uuid.UUID) (*ListPortfoliosResponseDTO, error) {
	portfolios, err := s.portfolioRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return ToListResponse(portfolios), nil
}

func (s *Service) UpdatePortfolio(ctx context.Context, id, userID uuid.UUID, dto *UpdatePortfolioDTO) (*PortfolioResponseDTO, error) {
	p, err := s.portfolioRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if p.UserID != userID {
		return nil, fmt.Errorf("unauthorized")
	}
	if dto.Title != nil {
		p.Title = *dto.Title
	}
	if dto.Description != nil {
		p.Description = *dto.Description
	}
	if dto.ProjectURL != nil {
		p.ProjectURL = *dto.ProjectURL
	}
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.portfolioRepo.Update(ctx, p); err != nil {
			return err
		}
		event, _ := s.createPortfolioEvent("portfolio.updated", p)
		return s.outboxRepo.Create(ctx, event)
	})
	if err != nil {
		return nil, err
	}
	return ToResponseDTO(p), nil
}

func (s *Service) DeletePortfolio(ctx context.Context, id, userID uuid.UUID) error {
	p, err := s.portfolioRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if p.UserID != userID {
		return fmt.Errorf("unauthorized")
	}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.portfolioRepo.Delete(ctx, id); err != nil {
			return err
		}
		event, _ := s.createPortfolioEvent("portfolio.removed", p)
		return s.outboxRepo.Create(ctx, event)
	})
}

func (s *Service) UploadImage(ctx context.Context, portfolioID, userID uuid.UUID, dto *UploadImageDTO) error {
	p, err := s.portfolioRepo.GetByID(ctx, portfolioID)
	if err != nil {
		return err
	}
	if p.UserID != userID {
		return fmt.Errorf("unauthorized")
	}
	image := &portfolio.PortfolioImage{
		PortfolioID:  portfolioID,
		ImageURL:     dto.ImageURL,
		Caption:      dto.Caption,
		DisplayOrder: dto.DisplayOrder,
	}
	return s.portfolioRepo.CreateImage(ctx, image)
}

func (s *Service) DeleteImage(ctx context.Context, imageID, userID uuid.UUID) error {
	return s.portfolioRepo.DeleteImage(ctx, imageID)
}

func (s *Service) createPortfolioEvent(eventType string, p *portfolio.Portfolio) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"portfolio_id": p.ID.String(),
		"user_id":      p.UserID.String(),
		"title":        p.Title,
	}
	payloadBytes, _ := json.Marshal(payload)
	metadata := map[string]interface{}{"source": "users-be"}
	metadataBytes, _ := json.Marshal(metadata)
	return &outbox.Event{
		AggregateID:   p.UserID.String(),
		AggregateType: "user",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}