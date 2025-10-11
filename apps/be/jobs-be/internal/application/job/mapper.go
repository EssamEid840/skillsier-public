package job

import "jobs-be/internal/domain/job"

func ToResponseDTO(j *job.Job) *JobResponseDTO {
	if j == nil {
		return nil
	}
	skills := make([]JobSkillDTO, len(j.Skills))
	for i, s := range j.Skills {
		skills[i] = JobSkillDTO{
			Name:     s.Name,
			Required: s.Required,
		}
	}
	return &JobResponseDTO{
		ID:              j.ID,
		ClientID:        j.ClientID,
		Title:           j.Title,
		Description:     j.Description,
		Category:        j.Category,
		Status:          j.Status,
		BudgetType:      j.BudgetType,
		BudgetAmount:    j.BudgetAmount,
		HourlyRateMin:   j.HourlyRateMin,
		HourlyRateMax:   j.HourlyRateMax,
		Duration:        j.Duration,
		ExperienceLevel: j.ExperienceLevel,
		ProposalCount:   j.ProposalCount,
		Skills:          skills,
		CreatedAt:       j.CreatedAt,
		UpdatedAt:       j.UpdatedAt,
	}
}

func ToListResponse(jobs []*job.Job, total int64, page, pageSize int) *ListJobsResponseDTO {
	dtos := make([]*JobResponseDTO, len(jobs))
	for i, j := range jobs {
		dtos[i] = ToResponseDTO(j)
	}
	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}
	return &ListJobsResponseDTO{
		Jobs:       dtos,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
}