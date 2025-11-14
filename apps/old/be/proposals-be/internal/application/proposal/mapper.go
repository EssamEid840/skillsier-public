package proposal

import (
	"proposals-be/internal/domain/proposal"
	"math"
)

func ToResponseDTO(p *proposal.Proposal) *ProposalResponseDTO {
	milestones := make([]MilestoneDTO, len(p.Milestones))
	for i, m := range p.Milestones {
		milestones[i] = MilestoneDTO{
			Description: m.Description,
			Amount:      m.Amount,
			DueDate:     m.DueDate,
		}
	}

	return &ProposalResponseDTO{
		ID:                p.ID,
		JobID:             p.JobID,
		FreelancerID:      p.FreelancerID,
		CoverLetter:       p.CoverLetter,
		BidAmount:         p.BidAmount,
		EstimatedDuration: p.EstimatedDuration,
		Status:            string(p.Status),
		Milestones:        milestones,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
	}
}

func ToResponseDTOList(proposals []*proposal.Proposal) []ProposalResponseDTO {
	result := make([]ProposalResponseDTO, len(proposals))
	for i, p := range proposals {
		result[i] = *ToResponseDTO(p)
	}
	return result
}

func ToListResponse(proposals []*proposal.Proposal, total int64, page, pageSize int) *ListProposalsResponseDTO {
	return &ListProposalsResponseDTO{
		Proposals:  ToResponseDTOList(proposals),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int(math.Ceil(float64(total) / float64(pageSize))),
	}
}
