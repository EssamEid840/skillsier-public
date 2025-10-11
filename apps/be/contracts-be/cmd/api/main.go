package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"contracts-be/internal/application/contract"
	"contracts-be/internal/config"
	"contracts-be/internal/infrastructure/messaging/kafka"
	"contracts-be/internal/infrastructure/outbox"
	"contracts-be/internal/infrastructure/persistence/postgres"
	httpInterface "contracts-be/internal/interfaces/http"
	"contracts-be/internal/interfaces/http/handlers"
)

func main() {
	cfg := config.Load()

	db, err := postgres.NewConnection(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := postgres.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	contractRepo := postgres.NewContractRepository(db)
	outboxRepo := postgres.NewOutboxRepository(db)

	contractService := contract.NewService(contractRepo, outboxRepo, db)
	contractHandler := handlers.NewContractHandler(contractService)

	router := httpInterface.SetupRouter(contractHandler)

	producer, err := kafka.NewProducer(cfg.Kafka)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	outboxProcessor := outbox.NewProcessor(db, producer, outboxRepo)
	go outboxProcessor.Start(context.Background())

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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}