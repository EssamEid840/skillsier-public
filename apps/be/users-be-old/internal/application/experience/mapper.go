package experience

import (
	"users-be/internal/domain/experience"
)

func ToResponseDTO(exp *experience.WorkExperience) *WorkExperienceResponseDTO {
	if exp == nil {
		return nil
	}
	return &WorkExperienceResponseDTO{
		ID:          exp.ID,
		UserID:      exp.UserID,
		Title:       exp.Title,
		Company:     exp.Company,
		Location:    exp.Location,
		StartDate:   exp.StartDate,
		EndDate:     exp.EndDate,
		IsCurrent:   exp.IsCurrent,
		Description: exp.Description,
		Skills:      exp.Skills,
		CreatedAt:   exp.CreatedAt,
		UpdatedAt:   exp.UpdatedAt,
	}
}

func ToListResponse(experiences []*experience.WorkExperience) *ListWorkExperiencesResponseDTO {
	dtos := make([]*WorkExperienceResponseDTO, len(experiences))
	for i, exp := range experiences {
		dtos[i] = ToResponseDTO(exp)
	}
	return &ListWorkExperiencesResponseDTO{
		Experiences: dtos,
		Total:       len(experiences),
	}
}