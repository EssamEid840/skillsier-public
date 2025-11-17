package ioc

import (
	"<module>/internal/application/initial_entity"
	initialentitydomain "<module>/internal/domain/initial_entity"
	"<module>/internal/infrastructure/messaging/kafka"
	"<module>/internal/infrastructure/messaging/outbox"
	"<module>/internal/infrastructure/persistence/postgres"
	"<module>/internal/interfaces/http/v1/handlers"
	
	"gorm.io/gorm"
)

func WireRepositories(db *gorm.DB) (initialentitydomain.Repository, *postgres.OutboxRepository) {
	initialEntityRepo := postgres.NewInitialEntityRepository(db)
	outboxRepo := postgres.NewOutboxRepository(db)
	
	return initialEntityRepo, outboxRepo
}

func WireServices(
	initialEntityRepo initialentitydomain.Repository,
	outboxRepo *postgres.OutboxRepository,
	db *gorm.DB,
) *initial_entity.Service {
	initialEntityService := initial_entity.NewService(initialEntityRepo, outboxRepo, db)
	
	return initialEntityService
}

func WireHandlers(
	initialEntityService *initial_entity.Service,
) (*handlers.InitialEntityHandler, *handlers.HealthHandler) {
	initialEntityHandler := handlers.NewInitialEntityHandler(initialEntityService)
	healthHandler := handlers.NewHealthHandler()
	
	return initialEntityHandler, healthHandler
}

func WireMessaging(
	db *gorm.DB,
	producer *kafka.Producer,
	outboxRepo *postgres.OutboxRepository,
) *outbox.Dispatcher {
	dispatcher := outbox.NewDispatcher(db, producer, outboxRepo)
	
	return dispatcher
}