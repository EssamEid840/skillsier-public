package postgres

import (
	"context"
	"jobs-be/internal/domain/job"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type jobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) job.Repository {
	return &jobRepository{db: db}
}

func (r *jobRepository) Create(ctx context.Context, j *job.Job) error {
	return r.db.WithContext(ctx).Create(j).Error
}

func (r *jobRepository) GetByID(ctx context.Context, id uuid.UUID) (*job.Job, error) {
	var j job.Job
	err := r.db.WithContext(ctx).Preload("RequiredSkills").Where("id = ?", id).First(&j).Error
	if err == gorm.ErrRecordNotFound {
		return nil, job.ErrJobNotFound
	}
	return &j, err
}

func (r *jobRepository) List(ctx context.Context, filters *job.ListFilters, limit, offset int) ([]*job.Job, int64, error) {
	var jobs []*job.Job
	var total int64

	query := r.db.WithContext(ctx).Model(&job.Job{}).Preload("RequiredSkills")

	// Apply filters
	if filters != nil {
		if filters.Category != nil {
			query = query.Where("category = ?", *filters.Category)
		}
		if filters.BudgetType != nil {
			query = query.Where("budget_type = ?", *filters.BudgetType)
		}
		if filters.Status != nil {
			query = query.Where("status = ?", *filters.Status)
		} else {
			// Default to open jobs only
			query = query.Where("status = ?", job.JobStatusOpen)
		}
		if filters.ExperienceLevel != nil {
			query = query.Where("experience_level = ?", *filters.ExperienceLevel)
		}
		if filters.SearchTerm != nil {
			searchPattern := "%" + *filters.SearchTerm + "%"
			query = query.Where("title ILIKE ? OR description ILIKE ?", searchPattern, searchPattern)
		}
		if filters.MinBudget != nil {
			query = query.Where("budget_amount >= ?", *filters.MinBudget)
		}
		if filters.MaxBudget != nil {
			query = query.Where("budget_amount <= ?", *filters.MaxBudget)
		}
		// Skills filter - join with job_skills table
		if len(filters.Skills) > 0 {
			query = query.Joins("JOIN job_skills ON job_skills.job_id = jobs.id").
				Where("job_skills.name IN ?", filters.Skills).
				Group("jobs.id")
		}
	}

	// Count total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&jobs).Error
	return jobs, total, err
}

func (r *jobRepository) Update(ctx context.Context, j *job.Job) error {
	return r.db.WithContext(ctx).Model(j).Updates(j).Error
}

func (r *jobRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&job.Job{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return job.ErrJobNotFound
	}
	return result.Error
}

func (r *jobRepository) GetByClientID(ctx context.Context, clientID uuid.UUID, limit, offset int) ([]*job.Job, int64, error) {
	var jobs []*job.Job
	var total int64

	query := r.db.WithContext(ctx).Model(&job.Job{}).Where("client_id = ?", clientID).Preload("RequiredSkills")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&jobs).Error
	return jobs, total, err
}

func (r *jobRepository) IncrementProposalCount(ctx context.Context, jobID uuid.UUID) error {
	return r.db.WithContext(ctx).Model(&job.Job{}).Where("id = ?", jobID).
		UpdateColumn("proposal_count", gorm.Expr("proposal_count + ?", 1)).Error
}