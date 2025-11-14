package experience

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrWorkExperienceNotFound = errors.New("work experience not found")
	ErrInvalidUserID          = errors.New("invalid user ID")
	ErrTitleRequired          = errors.New("title is required")
	ErrCompanyRequired        = errors.New("company is required")
	ErrStartDateRequired      = errors.New("start date is required")
	ErrEndDateRequired        = errors.New("end date is required for past positions")
	ErrInvalidDateRange       = errors.New("end date must be after start date")
)

type Repository interface {
	Create(ctx context.Context, experience *WorkExperience) error
	GetByID(ctx context.Context, id uuid.UUID) (*WorkExperience, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*WorkExperience, error)
	Update(ctx context.Context, experience *WorkExperience) error
	Delete(ctx context.Context, id uuid.UUID) error
}