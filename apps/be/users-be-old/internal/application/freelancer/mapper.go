package freelancer

import "users-be/internal/domain/freelancer"

func ToResponseDTO(profile *freelancer.FreelancerProfile) *FreelancerProfileResponseDTO {
	if profile == nil {
		return nil
	}
	return &FreelancerProfileResponseDTO{
		ID:                profile.ID,
		UserID:            profile.UserID,
		ProfessionalTitle: profile.ProfessionalTitle,
		Overview:          profile.Overview,
		HourlyRate:        profile.HourlyRate,
		Availability:      profile.Availability,
		TotalJobs:         profile.TotalJobs,
		TotalEarnings:     profile.TotalEarnings,
		SuccessRate:       profile.SuccessRate,
		Rating:            profile.Rating,
		ReviewCount:       profile.ReviewCount,
	}
}