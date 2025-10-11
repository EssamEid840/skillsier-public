package client

import "users-be/internal/domain/client"

func ToResponseDTO(profile *client.ClientProfile) *ClientProfileResponseDTO {
	return &ClientProfileResponseDTO{
		ID:              profile.ID,
		UserID:          profile.UserID,
		CompanyName:     profile.CompanyName,
		CompanySize:     profile.CompanySize,
		Industry:        profile.Industry,
		TotalSpent:      profile.TotalSpent,
		TotalJobsPosted: profile.TotalJobsPosted,
		TotalHired:      profile.TotalHired,
		PaymentVerified: profile.PaymentVerified,
		CreatedAt:       profile.CreatedAt,
		UpdatedAt:       profile.UpdatedAt,
	}
}