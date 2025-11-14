package job

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobSkill struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	JobID     uuid.UUID      `gorm:"type:uuid;not null;index" json:"job_id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	Required  bool           `gorm:"default:false" json:"required"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (JobSkill) TableName() string {
	return "job_skills"
}

func (js *JobSkill) BeforeCreate(tx *gorm.DB) error {
	if js.ID == uuid.Nil {
		js.ID = uuid.New()
	}
	return nil
}