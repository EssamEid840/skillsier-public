package postgres

import (
	"log"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	log.Println("Running database migrations...")
	err := db.AutoMigrate(
		&proposal.Proposal{},
		&proposal.ProposalMilestone{},
		&outbox.Event{},
	)
	if err != nil {
		log.Printf("Failed to run migrations: %v", err)
		return err
	}
	log.Println("Database migrations completed successfully")
	return nil
}
