package proposal

import (
	"context"
	"errors"
)

var (
	ErrProposalNotFound     = errors.New("proposal not found")
	ErrDuplicateProposal    = errors.New("proposal already exists for this job")
	ErrUnauthorized         = errors.New("unauthorized")
)

type Repository interface {
	Create(ctx context.Context, proposal *Proposal) error
	GetByID(ctx context.Context, id uuid.UUID) (*Proposal, error)
	GetByJobID(ctx context.Context, jobID uuid.UUID) ([]*Proposal, error)
	GetByFreelancerID(ctx context.Context, freelancerID uuid.UUID, limit, offset int) ([]*Proposal, int64, error)
	Update(ctx context.Context, proposal *Proposal) error
	Delete(ctx context.Context, id uuid.UUID) error
	CheckDuplicate(ctx context.Context, jobID, freelancerID uuid.UUID) (bool, error)
}