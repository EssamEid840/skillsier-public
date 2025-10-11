package job

import (
	"jobs-be/internal/domain/job"
	"math"
)

func ToResponseDTO(j *job.Job) *JobResponseDTO {
	skills := make([]SkillDTO, len(j.RequiredSkills))
	for i, s := range j.RequiredSkills {
		skills[i] = SkillDTO{Name: s.Name, Level: s.Level}
	}

	return &JobResponseDTO{
		ID:              j.ID,
		ClientID:        j.ClientID,
		Title:           j.Title,
		Description:     j.Description,
		Category:        j.Category,
		BudgetType:      string(j.BudgetType),
		BudgetAmount:    j.BudgetAmount,
		HourlyRateMin:   j.HourlyRateMin,
		HourlyRateMax:   j.HourlyRateMax,
		Duration:        j.Duration,
		ExperienceLevel: j.ExperienceLevel,
		Status:          string(j.Status),
		ProposalCount:   j.ProposalCount,
		RequiredSkills:  skills,
		CreatedAt:       j.CreatedAt,
		UpdatedAt:       j.UpdatedAt,
		ClosedAt:        j.ClosedAt,
	}
}

func ToResponseDTOList(jobs []*job.Job) []JobResponseDTO {
	result := make([]JobResponseDTO, len(jobs))
	for i, j := range jobs {
		result[i] = *ToResponseDTO(j)
	}
	return result
}

func ToListResponse(jobs []*job.Job, total int64, page, pageSize int) *ListJobsResponseDTO {
	return &ListJobsResponseDTO{
		Jobs:       ToResponseDTOList(jobs),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int(math.Ceil(float64(total) / float64(pageSize))),
	}
}
