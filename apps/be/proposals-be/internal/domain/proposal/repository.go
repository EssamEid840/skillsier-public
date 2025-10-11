package proposal

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrProposalNotFound      = errors.New("proposal not found")
	ErrUnauthorized          = errors.New("unauthorized access")
	ErrCoverLetterRequired   = errors.New("cover letter is required")
	ErrInvalidBidAmount      = errors.New("bid amount must be greater than zero")
	ErrDurationRequired      = errors.New("estimated duration is required")
	ErrJobAlreadySubmitted   = errors.New("you have already submitted a proposal for this job")
)

type Repository interface {
	Create(ctx context.Context, proposal *Proposal) error
	GetByID(ctx context.Context, id uuid.UUID) (*Proposal, error)
	List(ctx context.Context, filters *ListFilters, limit, offset int) ([]*Proposal, int64, error)
	Update(ctx context.Context, proposal *Proposal) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByJobID(ctx context.Context, jobID uuid.UUID, limit, offset int) ([]*Proposal, int64, error)
	GetByFreelancerID(ctx context.Context, freelancerID uuid.UUID, limit, offset int) ([]*Proposal, int64, error)
	CheckExisting(ctx context.Context, jobID uuid.UUID, freelancerID uuid.UUID) (bool, error)
}

type ListFilters struct {
	JobID        *uuid.UUID
	FreelancerID *uuid.UUID
	Status       *ProposalStatus
}
