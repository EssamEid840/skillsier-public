package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"skillsier.dev/apps/be/users-be/internal/config"
	"skillsier.dev/apps/be/users-be/internal/infrastructure/messaging/kafka"
	"skillsier.dev/apps/be/users-be/internal/infrastructure/persistence/postgres"
	"skillsier.dev/apps/be/users-be/internal/ioc"
	
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to database
	db, err := postgres.NewConnection(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations with config
	if err := postgres.AutoMigrate(db, cfg.Database); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Wire authentication (optional - uncomment to enable)
	// verifier, authMiddleware, err := ioc.WireAuth(cfg)
	// if err != nil {
	// 	log.Fatalf("Failed to setup authentication: %v", err)
	// }
	// defer verifier.(*keycloak.Verifier).Close()

	// Wire repositories using DI
	initialEntityRepo := ioc.WireRepositories(db)

	// Wire services using DI
	initialEntityService := ioc.WireServices(initialEntityRepo, db)

	// Wire handlers using DI
	initialEntityHandler, healthHandler := ioc.WireHandlers(initialEntityService, db)

	// Initialize Kafka producer
	producer, err := kafka.NewProducer(cfg.Kafka)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	// Wire messaging (outbox dispatcher using platform-shared)
	dispatcher := ioc.WireMessaging(db, producer)
	go dispatcher.Start(context.Background())

	// Setup HTTP router
	router := gin.Default()
	
	// Note: To enable auth, uncomment WireAuth above and pass authMiddleware here
	ioc.RegisterRoutes(router, initialEntityHandler, healthHandler)

	// Start HTTP server
	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}

	go func() {
		log.Printf("Starting server on port %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}