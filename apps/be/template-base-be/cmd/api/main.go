package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"<module>/internal/config"
	"<module>/internal/infrastructure/messaging/kafka"
	"<module>/internal/infrastructure/persistence/postgres"
	"<module>/internal/ioc"
	httpInterface "<module>/internal/interfaces/http"
	
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

	// Run migrations
	if err := postgres.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Wire repositories using DI
	initialEntityRepo, outboxRepo := ioc.WireRepositories(db)

	// Wire services using DI
	initialEntityService := ioc.WireServices(initialEntityRepo, outboxRepo, db)

	// Wire handlers using DI
	initialEntityHandler, healthHandler := ioc.WireHandlers(initialEntityService)

	// Initialize Kafka producer
	producer, err := kafka.NewProducer(cfg.Kafka)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	// Wire messaging (outbox dispatcher)
	dispatcher := ioc.WireMessaging(db, producer, outboxRepo)
	go dispatcher.Start(context.Background())

	// Setup HTTP router
	router := gin.Default()
	httpInterface.RegisterRoutes(router, initialEntityHandler, healthHandler)

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

// printBanner prints startup information
func printBanner(cfg *config.Config) {
	banner := `
╔════════════════════════════════════════════════════════════╗
║                                                            ║
║              <Service> Microservice                        ║
║                                                            ║
╚════════════════════════════════════════════════════════════╝

  Version:     %s
  Environment: %s
  Port:        %s
  Database:    %s:%d/%s
  Kafka:       %v
  Auto-Migrate: %t
  
  Logs Level:  %s
  
  Press Ctrl+C to stop...
`

	fmt.Printf(banner,
		cfg.App.Version,
		cfg.App.Environment,
		cfg.Server.Port,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
		cfg.Kafka.Brokers,
		cfg.Database.AutoMigrate,
		cfg.App.LogLevel,
	)
}