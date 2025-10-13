package skill

import (
	"users-be/internal/domain/skill"
)

func ToResponseDTO(s *skill.Skill) *SkillResponseDTO {
	if s == nil {
		return nil
	}
	return &SkillResponseDTO{
		ID:                s.ID,
		UserID:            s.UserID,
		Name:              s.Name,
		Category:          s.Category,
		Level:             s.Level,
		YearsOfExperience: s.YearsOfExperience,
		CreatedAt:         s.CreatedAt,
		UpdatedAt:         s.UpdatedAt,
	}
}

func ToListResponse(skills []*skill.Skill) *ListSkillsResponseDTO {
	dtos := make([]*SkillResponseDTO, len(skills))
	for i, s := range skills {
		dtos[i] = ToResponseDTO(s)
	}
	return &ListSkillsResponseDTO{
		Skills: dtos,
		Total:  len(skills),
	}
}