package postgres

import (
	"log"

	"gorm.io/gorm"

	"jobs-be/internal/domain/job"
	"jobs-be/internal/domain/outbox"
)

func AutoMigrate(db *gorm.DB) error {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&job.Job{},
		&job.JobSkill{},
		&outbox.Event{},
	)

	if err != nil {
		log.Printf("Failed to run migrations: %v", err)
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}