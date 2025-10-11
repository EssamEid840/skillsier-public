package freelancer

import "users-be/internal/domain/freelancer"

func ToResponseDTO(profile *freelancer.FreelancerProfile) *FreelancerProfileResponseDTO {
	return &FreelancerProfileResponseDTO{
		ID:             profile.ID,
		UserID:         profile.UserID,
		Title:          profile.Title,
		Overview:       profile.Overview,
		HourlyRate:     profile.HourlyRate,
		AvailableHours: profile.AvailableHours,
		ResponseTime:   profile.ResponseTime,
		TotalJobs:      profile.TotalJobs,
		TotalEarnings:  profile.TotalEarnings,
		SuccessRate:    profile.SuccessRate,
		CreatedAt:      profile.CreatedAt,
		UpdatedAt:      profile.UpdatedAt,
	}
}
