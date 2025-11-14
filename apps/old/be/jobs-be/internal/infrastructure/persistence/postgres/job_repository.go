package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"jobs-be/internal/domain/job"
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
	err := r.db.WithContext(ctx).Preload("Skills").Where("id = ?", id).First(&j).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, job.ErrJobNotFound
		}
		return nil, err
	}
	return &j, nil
}

func (r *jobRepository) GetAll(ctx context.Context, limit, offset int, filters map[string]interface{}) ([]*job.Job, int64, error) {
	var jobs []*job.Job
	var total int64

	query := r.db.WithContext(ctx).Model(&job.Job{}).Preload("Skills")

	// Apply filters
	if category, ok := filters["category"]; ok && category != "" {
		query = query.Where("category = ?", category)
	}
	if status, ok := filters["status"]; ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if budgetType, ok := filters["budget_type"]; ok && budgetType != "" {
		query = query.Where("budget_type = ?", budgetType)
	}
	if search, ok := filters["search"]; ok && search != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+search.(string)+"%", "%"+search.(string)+"%")
	}
	if skills, ok := filters["skills"]; ok && skills != nil {
		skillList := skills.([]string)
		if len(skillList) > 0 {
			query = query.Joins("JOIN job_skills ON job_skills.job_id = jobs.id").
				Where("job_skills.name IN ?", skillList)
		}
	}

	// Get total count
	query.Count(&total)

	// Get paginated results
	err := query.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&jobs).Error

	return jobs, total, err
}

func (r *jobRepository) GetByClientID(ctx context.Context, clientID uuid.UUID, limit, offset int) ([]*job.Job, int64, error) {
	var jobs []*job.Job
	var total int64

	query := r.db.WithContext(ctx).
		Model(&job.Job{}).
		Preload("Skills").
		Where("client_id = ?", clientID)

	query.Count(&total)

	err := query.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&jobs).Error

	return jobs, total, err
}

func (r *jobRepository) Update(ctx context.Context, j *job.Job) error {
	return r.db.WithContext(ctx).Save(j).Error
}

func (r *jobRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&job.Job{}, "id = ?", id).Error
}

func (r *jobRepository) IncrementProposalCount(ctx context.Context, jobID uuid.UUID) error {
	return r.db.WithContext(ctx).Model(&job.Job{}).
		Where("id = ?", jobID).
		UpdateColumn("proposal_count", gorm.Expr("proposal_count + ?", 1)).Error
}