package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"users-be/internal/application/certification"
	"users-be/internal/application/client"
	"users-be/internal/application/education"
	"users-be/internal/application/eventhandler"
	"users-be/internal/application/experience"
	"users-be/internal/application/freelancer"
	"users-be/internal/application/portfolio"
	"users-be/internal/application/skill"
	"users-be/internal/application/user"
	"users-be/internal/config"
	"users-be/internal/infrastructure/messaging/kafka"
	"users-be/internal/infrastructure/outbox"
	"users-be/internal/infrastructure/persistence/postgres"
	httpInterface "users-be/internal/interfaces/http"
	"users-be/internal/interfaces/http/handlers"
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

	// Initialize all repositories
	userRepo := postgres.NewUserRepository(db)
	outboxRepo := postgres.NewOutboxRepository(db)
	skillRepo := postgres.NewSkillRepository(db)
	experienceRepo := postgres.NewExperienceRepository(db)
	educationRepo := postgres.NewEducationRepository(db)
	certificationRepo := postgres.NewCertificationRepository(db)
	portfolioRepo := postgres.NewPortfolioRepository(db)
	freelancerRepo := postgres.NewFreelancerRepository(db)
	clientRepo := postgres.NewClientRepository(db)

	// Initialize all services
	userService := user.NewService(userRepo, outboxRepo, db)
	skillService := skill.NewService(skillRepo, outboxRepo, db)
	experienceService := experience.NewService(experienceRepo, outboxRepo, db)
	educationService := education.NewService(educationRepo, outboxRepo, db)
	certificationService := certification.NewService(certificationRepo, outboxRepo, db)
	portfolioService := portfolio.NewService(portfolioRepo, outboxRepo, db)
	freelancerService := freelancer.NewService(freelancerRepo, outboxRepo, db)
	clientService := client.NewService(clientRepo, outboxRepo, db)

	// Initialize all handlers
	userHandler := handlers.NewUserHandler(userService)
	skillHandler := handlers.NewSkillHandler(skillService)
	experienceHandler := handlers.NewExperienceHandler(experienceService)
	educationHandler := handlers.NewEducationHandler(educationService)
	certificationHandler := handlers.NewCertificationHandler(certificationService)
	portfolioHandler := handlers.NewPortfolioHandler(portfolioService)
	freelancerHandler := handlers.NewFreelancerHandler(freelancerService)
	clientHandler := handlers.NewClientHandler(clientService)

	// Setup router with all handlers
	router := httpInterface.SetupRouter(
		userHandler,
		skillHandler,
		experienceHandler,
		educationHandler,
		certificationHandler,
		portfolioHandler,
		freelancerHandler,
		clientHandler,
	)

	// Initialize Kafka producer
	producer, err := kafka.NewProducer(cfg.Kafka)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	// Start outbox processor
	outboxProcessor := outbox.NewProcessor(db, producer, outboxRepo)
	go outboxProcessor.Start(context.Background())

	// Initialize Kafka consumer for Keycloak events
	consumer, err := kafka.NewConsumer(cfg.Kafka, "users-service-group")
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer consumer.Close()

	// Start event handler
	keycloakHandler := eventhandler.NewKeycloakEventHandler(userService)
	go func() {
		if err := consumer.Subscribe([]string{"keycloak-events"}, keycloakHandler.HandleEvent); err != nil {
			log.Fatalf("Failed to subscribe to Keycloak events: %v", err)
		}
	}()

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