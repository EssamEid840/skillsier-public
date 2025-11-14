package postgres

import (
	"log"

	"gorm.io/gorm"

	"users-be/internal/domain/certification"
	"users-be/internal/domain/client"
	"users-be/internal/domain/education"
	"users-be/internal/domain/experience"
	"users-be/internal/domain/freelancer"
	"users-be/internal/domain/outbox"
	"users-be/internal/domain/portfolio"
	"users-be/internal/domain/skill"
	"users-be/internal/domain/user"
)

func AutoMigrate(db *gorm.DB) error {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		// Core tables
		&user.User{},
		&outbox.Event{},
		
		// Profile management tables
		&skill.Skill{},
		&experience.WorkExperience{},
		&education.Education{},
		&certification.Certification{},
		&portfolio.Portfolio{},
		&portfolio.PortfolioImage{},
		&freelancer.FreelancerProfile{},
		&client.ClientProfile{},
	)

	if err != nil {
		log.Printf("Failed to run migrations: %v", err)
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}