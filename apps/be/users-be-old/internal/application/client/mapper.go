package client

import "users-be/internal/domain/client"

func ToResponseDTO(profile *client.ClientProfile) *ClientProfileResponseDTO {
	if profile == nil {
		return nil
	}
	return &ClientProfileResponseDTO{
		ID:          profile.ID,
		UserID:      profile.UserID,
		CompanyName: profile.CompanyName,
		CompanySize: profile.CompanySize,
		Industry:    profile.Industry,
		Website:     profile.Website,
		TotalSpent:  profile.TotalSpent,
		JobsPosted:  profile.JobsPosted,
		HireRate:    profile.HireRate,
		Rating:      profile.Rating,
		ReviewCount: profile.ReviewCount,
		IsVerified:  profile.IsVerified,
	}
}