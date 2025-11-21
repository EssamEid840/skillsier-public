package ioc

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"skillsier.dev/apps/be/users-be/internal/application/initial_entity"
	"skillsier.dev/apps/be/users-be/internal/config"
	domainInitialEntity "skillsier.dev/apps/be/users-be/internal/domain/initial_entity"
	"skillsier.dev/apps/be/users-be/internal/infrastructure/messaging/kafka"
	"skillsier.dev/apps/be/users-be/internal/infrastructure/messaging/outbox"
	"skillsier.dev/apps/be/users-be/internal/infrastructure/persistence/postgres"
	"skillsier.dev/apps/be/users-be/internal/interfaces/http/middleware"
	"skillsier.dev/apps/be/users-be/internal/interfaces/http/v1/handlers"
	
	"skillsier.dev/pkg/auth"
	"skillsier.dev/pkg/auth/keycloak"
	platformOutbox "skillsier.dev/platform-shared/outbox"
	platformOutboxPostgres "skillsier.dev/platform-shared/outbox/postgres"
	"skillsier.dev/platform-shared/logging"
)

// WireRepositories creates and wires all repositories
func WireRepositories(db *gorm.DB) domainInitialEntity.Repository {
	return postgres.NewInitialEntityRepository(db)
}

// WireServices creates and wires all services
func WireServices(
	initialEntityRepo domainInitialEntity.Repository,
	db *gorm.DB,
) *initial_entity.Service {
	return initial_entity.NewService(initialEntityRepo, db)
}

// WireHandlers creates and wires all HTTP handlers
func WireHandlers(
	initialEntityService *initial_entity.Service,
	db *gorm.DB,
) (*handlers.InitialEntityHandler, *handlers.HealthHandler) {
	initialEntityHandler := handlers.NewInitialEntityHandler(initialEntityService)
	healthHandler := handlers.NewHealthHandler(db)
	
	return initialEntityHandler, healthHandler
}

// WireMessaging creates and wires messaging components (outbox forwarder)
func WireMessaging(
	db *gorm.DB,
	kafkaProducer *kafka.Producer,
) *platformOutbox.Forwarder {
	// Use platform-shared outbox repository
	outboxRepo := platformOutboxPostgres.NewRepository(db)
	
	// Create Kafka publisher adapter
	publisher := outbox.NewKafkaPublisher(kafkaProducer)
	
	// Create forwarder config
	config := &platformOutbox.Config{
		PollInterval: 5 * time.Second,
		BatchSize:    100,
	}
	
	// Create logger
	logger := &logging.Logger{} // Simple logger instance
	
	// Create and return forwarder using platform-shared
	return platformOutbox.NewForwarder(outboxRepo, publisher, config, logger)
}

// WireAuth creates and wires authentication components
func WireAuth(cfg *config.Config) (auth.TokenVerifier, *middleware.AuthMiddleware, error) {
	// Create Keycloak config
	keycloakConfig := keycloak.NewConfig(
		cfg.Keycloak.BaseURL,
		cfg.Keycloak.Realm,
		cfg.Keycloak.ClientID,
		cfg.Keycloak.ClientSecret,
	)
	
	// Create token verifier
	verifier, err := keycloak.NewVerifier(keycloakConfig)
	if err != nil {
		return nil, nil, err
	}
	
	// Create auth middleware
	authMiddleware := middleware.NewAuthMiddleware(verifier)
	
	return verifier, authMiddleware, nil
}

// RegisterRoutes registers all HTTP routes
func RegisterRoutes(
	router *gin.Engine,
	initialEntityHandler *handlers.InitialEntityHandler,
	healthHandler *handlers.HealthHandler,
) {
	// Health check
	router.GET("/health", healthHandler.Health)
	router.GET("/ready", healthHandler.Ready)
	
	// API v1
	v1 := router.Group("/v1")
	{
		// Initial entity routes
		entities := v1.Group("/initial-entities")
		{
			entities.POST("", initialEntityHandler.Create)
			entities.GET("/:id", initialEntityHandler.Get)
			entities.GET("", initialEntityHandler.List)
			entities.PUT("/:id", initialEntityHandler.Update)
			entities.DELETE("/:id", initialEntityHandler.Delete)
		}
	}
}