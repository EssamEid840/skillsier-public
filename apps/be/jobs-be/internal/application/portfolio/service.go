package portfolio

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"users-be/internal/domain/portfolio"
	"users-be/internal/domain/outbox"
	"users-be/internal/infrastructure/storage"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	MaxPortfolioPerUser = 20
	MaxImagesPerPortfolio = 5
)

type Service struct {
	portfolioRepo portfolio.Repository
	outboxRepo    outbox.Repository
	storage       *storage.LocalStorage
	db            *gorm.DB
}

func NewService(
	portfolioRepo portfolio.Repository,
	outboxRepo outbox.Repository,
	storage *storage.LocalStorage,
	db *gorm.DB,
) *Service {
	return &Service{
		portfolioRepo: portfolioRepo,
		outboxRepo:    outboxRepo,
		storage:       storage,
		db:            db,
	}
}

func (s *Service) CreatePortfolio(ctx context.Context, userID uuid.UUID, dto *CreatePortfolioDTO) (*PortfolioResponseDTO, error) {
	count, err := s.portfolioRepo.CountByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to count portfolio items: %w", err)
	}
	if count >= MaxPortfolioPerUser {
		return nil, portfolio.ErrMaxPortfolioExceeded
	}

	newPortfolio := &portfolio.Portfolio{
		UserID:      userID,
		Title:       dto.Title,
		Description: dto.Description,
		ProjectURL:  dto.ProjectURL,
		StartDate:   dto.StartDate,
		EndDate:     dto.EndDate,
	}

	if err := newPortfolio.Validate(); err != nil {
		return nil, err
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.portfolioRepo.Create(ctx, newPortfolio); err != nil {
			return err
		}

		event, err := s.createEvent("portfolio.item.created", newPortfolio)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(newPortfolio), nil
}

func (s *Service) GetPortfolios(ctx context.Context, userID uuid.UUID) ([]*PortfolioResponseDTO, error) {
	portfolios, err := s.portfolioRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return ToResponseDTOList(portfolios), nil
}

func (s *Service) UpdatePortfolio(ctx context.Context, id uuid.UUID, userID uuid.UUID, dto *UpdatePortfolioDTO) (*PortfolioResponseDTO, error) {
	p, err := s.portfolioRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if p.UserID != userID {
		return nil, portfolio.ErrPortfolioNotFound
	}

	if dto.Title != nil {
		p.Title = *dto.Title
	}
	if dto.Description != nil {
		p.Description = dto.Description
	}
	if dto.ProjectURL != nil {
		p.ProjectURL = dto.ProjectURL
	}
	if dto.StartDate != nil {
		p.StartDate = dto.StartDate
	}
	if dto.EndDate != nil {
		p.EndDate = dto.EndDate
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.portfolioRepo.Update(ctx, p); err != nil {
			return err
		}

		event, err := s.createEvent("portfolio.item.updated", p)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(p), nil
}

func (s *Service) DeletePortfolio(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	p, err := s.portfolioRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if p.UserID != userID {
		return portfolio.ErrPortfolioNotFound
	}

	// Delete images from storage
	for _, img := range p.Images {
		s.storage.DeleteFile(img.ImageURL)
	}

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.portfolioRepo.Delete(ctx, id); err != nil {
			return err
		}

		event, err := s.createEvent("portfolio.item.deleted", p)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})
}

func (s *Service) UploadImage(ctx context.Context, portfolioID uuid.UUID, userID uuid.UUID, file *multipart.FileHeader) (*ImageResponseDTO, error) {
	// Verify ownership
	p, err := s.portfolioRepo.GetByID(ctx, portfolioID)
	if err != nil {
		return nil, err
	}
	if p.UserID != userID {
		return nil, portfolio.ErrPortfolioNotFound
	}

	// Check max images
	count, err := s.portfolioRepo.CountImages(ctx, portfolioID)
	if err != nil {
		return nil, err
	}
	if count >= MaxImagesPerPortfolio {
		return nil, portfolio.ErrMaxImagesExceeded
	}

	// Save file
	imageURL, err := s.storage.SavePortfolioImage(file)
	if err != nil {
		return nil, fmt.Errorf("failed to save image: %w", err)
	}

	// Create image record
	image := &portfolio.PortfolioImage{
		PortfolioID:  portfolioID,
		ImageURL:     imageURL,
		DisplayOrder: int(count),
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.portfolioRepo.AddImage(ctx, image); err != nil {
			// Rollback: delete uploaded file
			s.storage.DeleteFile(imageURL)
			return err
		}

		event, err := s.createImageEvent("portfolio.image.uploaded", image)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return &ImageResponseDTO{
		ID:           image.ID,
		PortfolioID:  image.PortfolioID,
		ImageURL:     image.ImageURL,
		DisplayOrder: image.DisplayOrder,
		CreatedAt:    image.CreatedAt,
	}, nil
}

func (s *Service) createEvent(eventType string, p *portfolio.Portfolio) (*outbox.Event, error) {
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
		AggregateType: "portfolio",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}

func (s *Service) createImageEvent(eventType string, img *portfolio.PortfolioImage) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"image_id":     img.ID.String(),
		"portfolio_id": img.PortfolioID.String(),
		"image_url":    img.ImageURL,
	}
	payloadBytes, _ := json.Marshal(payload)
	metadata := map[string]interface{}{"source": "users-be"}
	metadataBytes, _ := json.Marshal(metadata)

	return &outbox.Event{
		AggregateID:   img.PortfolioID.String(),
		AggregateType: "portfolio_image",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}
