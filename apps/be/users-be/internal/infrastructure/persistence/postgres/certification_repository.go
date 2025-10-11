package postgres

import (
	"context"
	"users-be/internal/domain/certification"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type certificationRepository struct {
	db *gorm.DB
}

func NewCertificationRepository(db *gorm.DB) certification.Repository {
	return &certificationRepository{db: db}
}

func (r *certificationRepository) Create(ctx context.Context, cert *certification.Certification) error {
	return r.db.WithContext(ctx).Create(cert).Error
}

func (r *certificationRepository) GetByID(ctx context.Context, id uuid.UUID) (*certification.Certification, error) {
	var cert certification.Certification
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&cert).Error
	if err == gorm.ErrRecordNotFound {
		return nil, certification.ErrCertificationNotFound
	}
	return &cert, err
}

func (r *certificationRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*certification.Certification, error) {
	var certs []*certification.Certification
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).
		Order("issue_date DESC").Find(&certs).Error
	return certs, err
}

func (r *certificationRepository) Update(ctx context.Context, cert *certification.Certification) error {
	return r.db.WithContext(ctx).Model(cert).Updates(cert).Error
}

func (r *certificationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&certification.Certification{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return certification.ErrCertificationNotFound
	}
	return result.Error
}

func (r *certificationRepository) CountByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&certification.Certification{}).
		Where("user_id = ?", userID).Count(&count).Error
	return count, err
}
