package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"users-be/internal/application/user"
	"users-be/internal/infrastructure/messaging/kafka"
	"users-be/internal/infrastructure/outbox"
	"users-be/internal/infrastructure/persistence/postgres"
	httpInterface "users-be/internal/interfaces/http"
	"users-be/internal/interfaces/http/handlers"
	"users-be/internal/application/eventhandler"
	"users-be/internal/config"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	log.Printf("Starting %s service in %s mode...", cfg.App.Name, cfg.App.Environment)

	// Setup database connection
	db, err := postgres.NewConnection(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer postgres.Close(db)

	log.Println("✓ Database connected successfully")

	// Run auto-migrations
	if err := postgres.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("✓ Database migrations completed")

	// Initialize repositories
	userRepo := postgres.NewUserRepository(db)
	outboxRepo := postgres.NewOutboxRepository(db)

	// Initialize services
	userService := user.NewService(userRepo, outboxRepo, db)

	// Initialize Kafka producer for outbox
	kafkaProducer, err := kafka.NewProducer(&cfg.Kafka)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer kafkaProducer.Close()
	log.Println("✓ Kafka producer initialized")

	// Initialize outbox processor
	outboxProcessor := outbox.NewProcessor(
		outboxRepo,
		kafkaProducer,
		5*time.Second,  // Poll every 5 seconds
		100,            // Batch size
		5,              // Max retries
	)

	// Start outbox processor in background
	outboxCtx, cancelOutbox := context.WithCancel(context.Background())
	defer cancelOutbox()
	
	go func() {
		if err := outboxProcessor.Start(outboxCtx); err != nil && err != context.Canceled {
			log.Printf("Outbox processor error: %v", err)
		}
	}()
	log.Println("✓ Outbox processor started")

	// Initialize Keycloak event handler
	keycloakHandler := eventhandler.NewKeycloakEventHandler(userService)

	// Initialize Kafka consumer for Keycloak events
	kafkaConsumer, err := kafka.NewConsumer(
		&cfg.Kafka,
		[]string{cfg.Kafka.KeycloakEventsTopic},
		keycloakHandler.HandleMessage,
	)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer kafkaConsumer.Close()
	log.Println("✓ Kafka consumer initialized")

	// Start Kafka consumer in background
	consumerCtx, cancelConsumer := context.WithCancel(context.Background())
	defer cancelConsumer()
	
	go func() {
		if err := kafkaConsumer.Start(consumerCtx); err != nil && err != context.Canceled {
			log.Printf("Kafka consumer error: %v", err)
		}
	}()
	log.Printf("✓ Kafka consumer started (topic: %s)", cfg.Kafka.KeycloakEventsTopic)

	// Initialize HTTP handlers
	userHandler := handlers.NewUserHandler(userService)
	healthHandler := handlers.NewHealthHandler(db)

	// Setup router
	router := httpInterface.SetupRouter(userHandler, healthHandler)

	// Create HTTP server
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	server := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	// Start HTTP server in background
	go func() {
		log.Printf("✓ HTTP server listening on %s", addr)
		log.Println("===========================================")
		log.Println("Users-BE Service is Ready!")
		log.Println("===========================================")
		log.Println("")
		log.Println("API Endpoints:")
		log.Println("  GET    /health                          - Health check")
		log.Println("  GET    /ready                           - Readiness probe")
		log.Println("  GET    /live                            - Liveness probe")
		log.Println("  POST   /api/v1/users                    - Create user")
		log.Println("  GET    /api/v1/users                    - List users")
		log.Println("  GET    /api/v1/users/:id                - Get user by ID")
		log.Println("  GET    /api/v1/users/keycloak/:id       - Get user by Keycloak ID")
		log.Println("  PUT    /api/v1/users/:id                - Update user")
		log.Println("  DELETE /api/v1/users/:id                - Delete user")
		log.Println("")
		log.Println("Event Sources:")
		log.Printf("  Kafka Consumer: %s\n", cfg.Kafka.KeycloakEventsTopic)
		log.Printf("  Kafka Producer: %s\n", cfg.Kafka.UserEventsTopic)
		log.Println("")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Shutdown HTTP server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	// Cancel background workers
	cancelOutbox()
	cancelConsumer()

	// Cleanup outbox (remove old published events)
	cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cleanupCancel()
	
	if err := outboxProcessor.CleanupPublishedEvents(cleanupCtx, 7); err != nil {
		log.Printf("Failed to cleanup outbox: %v", err)
	}

	log.Println("Server exited")
}