package postgres

import (
	"log"

	"gorm.io/gorm"

	"contracts-be/internal/domain/contract"
	"contracts-be/internal/domain/outbox"
)

func AutoMigrate(db *gorm.DB) error {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&contract.Contract{},
		&contract.ContractMilestone{},
		&outbox.Event{},
	)

	if err != nil {
		log.Printf("Failed to run migrations: %v", err)
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}