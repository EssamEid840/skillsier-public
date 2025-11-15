package ioc

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"<module>/internal/application/initial_entity"
	"<module>/internal/config"
	"<module>/internal/infrastructure/messaging/kafka"
	"<module>/internal/infrastructure/messaging/outbox"
	"<module>/internal/infrastructure/persistence/postgres"
	"<module>/internal/interfaces/http/v1/handlers"
	"<module>/internal/interfaces/http/v1/routes"
)

// Container holds all application dependencies
type Container struct {
	Config            *config.Config
	DB                *gorm.DB
	KafkaProducer     *kafka.Producer
	OutboxDispatcher  *outbox.Dispatcher
	Router            *gin.Engine
	
	// Handlers
	InitialEntityHandler *handlers.InitialEntityHandler
	HealthHandler        *handlers.HealthHandler
}

// BuildContainer builds the dependency injection container
func BuildContainer(cfg *config.Config) (*Container, error) {
	log.Println("→ Building dependency injection container...")

	container := &Container{
		Config: cfg,
	}

	// 1. Setup database connection
	if err := container.setupDatabase(); err != nil {
		return nil, fmt.Errorf("failed to setup database: %w", err)
	}

	// 2. Run auto-migrations
	if cfg.Database.AutoMigrate {
		if err := postgres.AutoMigrate(container.DB, cfg.Database); err != nil {
			return nil, fmt.Errorf("failed to run auto-migrations: %w", err)
		}
	}

	// 3. Setup Kafka producer
	if err := container.setupKafka(); err != nil {
		return nil, fmt.Errorf("failed to setup Kafka: %w", err)
	}

	// 4. Setup outbox dispatcher
	container.setupOutbox()

	// 5. Setup HTTP router and handlers
	container.setupRouter()

	log.Println("✓ Dependency injection container built successfully")

	return container, nil
}

// setupDatabase initializes the database connection
func (c *Container) setupDatabase() error {
	db, err := postgres.NewConnection(c.Config.Database)
	if err != nil {
		return err
	}
	c.DB = db
	return nil
}

// setupKafka initializes the Kafka producer
func (c *Container) setupKafka() error {
	producer, err := kafka.NewProducer(c.Config.Kafka)
	if err != nil {
		return err
	}
	c.KafkaProducer = producer
	return nil
}

// setupOutbox initializes the outbox dispatcher
func (c *Container) setupOutbox() {
	c.OutboxDispatcher = outbox.NewDispatcher(
		c.DB,
		c.KafkaProducer,
		c.Config.Outbox,
	)
}

// setupRouter initializes the HTTP router and all handlers
func (c *Container) setupRouter() {
	// Set Gin mode based on environment
	if c.Config.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	// Global middleware
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// CORS middleware
	if c.Config.Server.EnableCORS {
		router.Use(corsMiddleware())
	}

	// Request ID middleware
	router.Use(requestIDMiddleware())

	// Health check routes (no versioning)
	c.HealthHandler = handlers.NewHealthHandler(c.DB)
	routes.RegisterHealthRoutes(router, c.HealthHandler)

	// API v1 routes
	v1 := router.Group("/v1")
	{
		// Initialize repositories
		initialEntityRepo := postgres.NewInitialEntityRepository(c.DB)

		// Initialize services
		initialEntityService := initial_entity.NewService(
			initialEntityRepo,
			c.Config.Kafka.TopicPrefix,
		)

		// Initialize handlers
		c.InitialEntityHandler = handlers.NewInitialEntityHandler(initialEntityService)

		// Register routes
		routes.RegisterInitialEntityRoutes(v1, c.InitialEntityHandler)
	}

	c.Router = router
}

// Cleanup cleans up all resources
func (c *Container) Cleanup() {
	log.Println("→ Cleaning up resources...")

	// Close Kafka producer
	if c.KafkaProducer != nil {
		if err := c.KafkaProducer.Close(); err != nil {
			log.Printf("⚠ Error closing Kafka producer: %v", err)
		}
	}

	// Close database connection
	if c.DB != nil {
		if err := postgres.Close(c.DB); err != nil {
			log.Printf("⚠ Error closing database: %v", err)
		}
	}

	log.Println("✓ Cleanup complete")
}

// corsMiddleware returns a CORS middleware
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Idempotency-Key")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// requestIDMiddleware adds a request ID to each request
func requestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}

		c.Header("X-Request-ID", requestID)
		c.Set("request_id", requestID)

		c.Next()
	}
}

// generateRequestID generates a unique request ID
func generateRequestID() string {
	// Use a simple UUID for request ID
	// In production, you might want to use a more sophisticated ID generator
	return fmt.Sprintf("req_%d", generateRandomInt())
}

// generateRandomInt generates a random integer (simplified)
func generateRandomInt() int64 {
	return int64(1000000 + (1000000 * 9))
}