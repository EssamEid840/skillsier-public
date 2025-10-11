package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"users-be/internal/domain/certification"
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
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, certification.ErrCertificationNotFound
		}
		return nil, err
	}
	return &cert, nil
}

func (r *certificationRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*certification.Certification, error) {
	var certifications []*certification.Certification
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("issue_date DESC").
		Find(&certifications).Error
	return certifications, err
}

func (r *certificationRepository) Update(ctx context.Context, cert *certification.Certification) error {
	return r.db.WithContext(ctx).Save(cert).Error
}

func (r *certificationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&certification.Certification{}, "id = ?", id).Error
}