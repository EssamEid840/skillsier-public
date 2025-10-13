package education

import (
	"users-be/internal/domain/education"
)

func ToResponseDTO(edu *education.Education) *EducationResponseDTO {
	if edu == nil {
		return nil
	}
	return &EducationResponseDTO{
		ID:           edu.ID,
		UserID:       edu.UserID,
		Degree:       edu.Degree,
		Institution:  edu.Institution,
		FieldOfStudy: edu.FieldOfStudy,
		StartDate:    edu.StartDate,
		EndDate:      edu.EndDate,
		Description:  edu.Description,
		CreatedAt:    edu.CreatedAt,
		UpdatedAt:    edu.UpdatedAt,
	}
}

func ToListResponse(educations []*education.Education) *ListEducationsResponseDTO {
	dtos := make([]*EducationResponseDTO, len(educations))
	for i, edu := range educations {
		dtos[i] = ToResponseDTO(edu)
	}
	return &ListEducationsResponseDTO{
		Educations: dtos,
		Total:      len(educations),
	}
}