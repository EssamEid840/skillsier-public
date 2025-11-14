package review

import (
	"reviews-be/internal/domain/review"
	"math"
)

func ToResponseDTO(r *review.Review) *ReviewResponseDTO {
	return &ReviewResponseDTO{
		ID:                r.ID,
		ContractID:        r.ContractID,
		ReviewerID:        r.ReviewerID,
		RevieweeID:        r.RevieweeID,
		Rating:            r.Rating,
		Comment:           r.Comment,
		QualityOfWork:     r.QualityOfWork,
		Communication:     r.Communication,
		Professionalism:   r.Professionalism,
		DeadlineAdherence: r.DeadlineAdherence,
		CreatedAt:         r.CreatedAt,
		UpdatedAt:         r.UpdatedAt,
	}
}

func ToResponseDTOList(reviews []*review.Review) []ReviewResponseDTO {
	result := make([]ReviewResponseDTO, len(reviews))
	for i, r := range reviews {
		result[i] = *ToResponseDTO(r)
	}
	return result
}

func ToListResponse(reviews []*review.Review, total int64, page, pageSize int, avgRating float64, reviewCount int64) *ListReviewsResponseDTO {
	return &ListReviewsResponseDTO{
		Reviews:       ToResponseDTOList(reviews),
		Total:         total,
		Page:          page,
		PageSize:      pageSize,
		TotalPages:    int(math.Ceil(float64(total) / float64(pageSize))),
		AverageRating: avgRating,
		ReviewCount:   reviewCount,
	}
}
