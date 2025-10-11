package skill

import "users-be/internal/domain/skill"

func ToResponseDTO(s *skill.Skill) *SkillResponseDTO {
	return &SkillResponseDTO{
		ID:           s.ID,
		UserID:       s.UserID,
		Name:         s.Name,
		Level:        s.Level,
		YearsOfExp:   s.YearsOfExp,
		Endorsements: s.Endorsements,
		IsPrimary:    s.IsPrimary,
		CreatedAt:    s.CreatedAt,
		UpdatedAt:    s.UpdatedAt,
	}
}

func ToResponseDTOList(skills []*skill.Skill) []*SkillResponseDTO {
	result := make([]*SkillResponseDTO, len(skills))
	for i, s := range skills {
		result[i] = ToResponseDTO(s)
	}
	return result
}
