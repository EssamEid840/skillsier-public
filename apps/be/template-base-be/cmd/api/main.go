package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"<module>/internal/config"
	"<module>/internal/ioc"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Print startup banner
	printBanner(cfg)

	// Build dependency injection container
	container, err := ioc.BuildContainer(cfg)
	if err != nil {
		log.Fatalf("Failed to build DI container: %v", err)
	}
	defer container.Cleanup()

	// Start outbox dispatcher (if enabled)
	if cfg.Outbox.Enabled {
		log.Println("Starting outbox dispatcher...")
		go container.OutboxDispatcher.Start(context.Background())
	}

	// Start HTTP server
	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      container.Router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	// Start server in goroutine
	go func() {
		log.Printf("Starting HTTP server on port %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
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