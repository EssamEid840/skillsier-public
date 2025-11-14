package postgres

import (
	"log"

	"gorm.io/gorm"

	"reviews-be/internal/domain/outbox"
	"reviews-be/internal/domain/review"
)

func AutoMigrate(db *gorm.DB) error {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&review.Review{},
		&outbox.Event{},
	)

	if err != nil {
		log.Printf("Failed to run migrations: %v", err)
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}