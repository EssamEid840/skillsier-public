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

	"jobs-be/internal/application/job"
	"jobs-be/internal/infrastructure/messaging/kafka"
	"jobs-be/internal/infrastructure/outbox"
	"jobs-be/internal/infrastructure/persistence/postgres"
	httpInterface "jobs-be/internal/interfaces/http"
	"jobs-be/internal/interfaces/http/handlers"
	"jobs-be/internal/config"
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
	jobRepo := postgres.NewJobRepository(db)
	outboxRepo := postgres.NewOutboxRepository(db)

	// Initialize services
	jobService := job.NewService(jobRepo, outboxRepo, db)

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
		5*time.Second,
		100,
		5,
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

	// Initialize HTTP handlers
	jobHandler := handlers.NewJobHandler(jobService)
	healthHandler := handlers.NewHealthHandler(db)

	// Setup router
	router := httpInterface.SetupRouter(jobHandler, healthHandler)

	// Create HTTP server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.App.Port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in goroutine
	go func() {
		log.Printf("🚀 Server starting on port %s...", cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Printf("✓ jobs-be service is running on port %s", cfg.App.Port)
	log.Println("✓ Ready to accept requests")

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited successfully")
}