package ioc

import (
	"<module>/internal/application/initial_entity"
	"<module>/internal/domain/initial_entity"
	"<module>/internal/infrastructure/messaging/kafka"
	"<module>/internal/infrastructure/messaging/outbox"
	"<module>/internal/infrastructure/persistence/postgres"
	"<module>/internal/interfaces/http/v1/handlers"
	
	"gorm.io/gorm"
)

// WireRepositories creates all repository implementations
func WireRepositories(db *gorm.DB) (*postgres.InitialEntityRepository, *postgres.OutboxStore) {
	initialEntityRepo := postgres.NewInitialEntityRepository(db)
	outboxStore := postgres.NewOutboxStore(db)
	
	return initialEntityRepo, outboxStore
}

// WireServices creates all application services with their dependencies
func WireServices(
	initialEntityRepo initial_entity.Repository,
	outboxStore *postgres.OutboxStore,
) *initial_entity.Service {
	initialEntityService := initial_entity.NewService(initialEntityRepo, outboxStore)
	
	return initialEntityService
}

// WireHandlers creates all HTTP handlers with their dependencies
func WireHandlers(
	initialEntityService *initial_entity.Service,
) (*handlers.InitialEntityHandler, *handlers.HealthHandler) {
	initialEntityHandler := handlers.NewInitialEntityHandler(initialEntityService)
	healthHandler := handlers.NewHealthHandler()
	
	return initialEntityHandler, healthHandler
}

// WireMessaging creates Kafka producer and outbox dispatcher
func WireMessaging(
	db *gorm.DB,
	producer *kafka.Producer,
	outboxStore *postgres.OutboxStore,
	cfg interface{}, // Pass config for outbox settings
) *outbox.Dispatcher {
	dispatcher := outbox.NewDispatcher(db, producer, outboxStore)
	
	return dispatcher
}

// WireEventConsumers creates Kafka consumers for consuming events from other services
// Uncomment and implement when you need to consume events
/*
func WireEventConsumers(
	consumer *kafka.Consumer,
	initialEntityService *initial_entity.Service,
) {
	// Example: Register handlers for events from other services
	// consumer.Subscribe("other-service-events", handlers.HandleOtherServiceEvent)
}
*/