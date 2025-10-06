package setup

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"users.go/m/internal/infrastructure/database"
	"users.go/m/internal/infrastructure/event"
	"users.go/m/internal/infrastructure/event/dapr/daprpub"
	"users.go/m/internal/infrastructure/event/kafka"

	"users.go/m/internal/repository/outbox"
	"users.go/m/internal/repository/users"
	"users.go/m/internal/routes"
	uow2 "users.go/m/internal/uow"
	users2 "users.go/m/internal/usecases/users"
	"users.go/m/internal/worker/outbox_dispatcher"
)

type App struct {
	database database.Database
	router   *gin.Engine
}

func strToBool(s string) bool {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "1", "true", "t", "yes", "y", "on":
		return true
	default:
		return false
	}
}

func strToIntDefault(s string, def int) int {
	if v, err := strconv.Atoi(strings.TrimSpace(s)); err == nil {
		return v
	}
	return def
}

func New() *App {
	container, err := BuildContainer()
	if err != nil {
		panic(err)
	}

	userDb, err := container.Database.GetConnection("users")
	if err != nil {
		panic(err)
	}

	router := container.Engine
	uow := uow2.NewUnitOfWork(userDb)

	usersRepo := users.NewUsersRepo(userDb)
	outboxRepo := outbox.NewOutboxRepo(userDb)

	// --- config ---
	topic := container.Config.GetString("kafka.topic")
	if topic == "" {
		topic = "user"
	}
	daprEnabled := strToBool(container.Config.GetString("dapr.enabled"))
	daprHTTPPort := strToIntDefault(container.Config.GetString("dapr.http_port"), 3501)
	daprPubsub := container.Config.GetString("dapr.pubsub")
	if daprPubsub == "" {
		daprPubsub = "kafka-pubsub"
	}

	// --- publisher selection (Dapr or native Kafka) ---
	var eventPublisher event.EventPublisher

	if daprEnabled {
		baseURL := fmt.Sprintf("http://localhost:%d", daprHTTPPort)
		eventPublisher = daprpub.New(daprpub.Config{
			BaseURL:    baseURL,
			PubsubName: daprPubsub,
		})
	} else {
		k := kafka.NewKafkaPublisher(kafka.KafkaProducerConfig{
			Host:    container.Config.GetString("kafka.host"),
			GroupID: container.Config.GetString("kafka.group_id"),
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		if err := k.EnsureTopics(ctx, []string{topic}); err != nil {
			log.Warn().Err(err).Str("topic", topic).Msg("failed ensuring Kafka topic")
		}
		cancel()

		eventPublisher = k
	}

	// Use cases & routes
	createUserUseCase := users2.NewCreateUserUseCase(usersRepo, uow, outboxRepo, eventPublisher)
	routes.NewUsersRoutes(router, createUserUseCase)

	// --- Outbox dispatcher gating (CDC vs in-app dispatcher) ---
	mode := os.Getenv("APP_OUTBOX__MODE") // "cdc" or "dispatcher"
	if mode == "" {
		mode = "cdc" // default to CDC
	}

	if mode == "dispatcher" {
		// Outbox dispatcher works with either publisher (Kafka or Dapr)
		outboxDispatcher := outbox_dispatcher.NewOutboxDispatcher(eventPublisher, outboxRepo)
		go outboxDispatcher.Execute()
	} else {
		log.Info().Msg("CDC mode enabled; in-app outbox dispatcher is DISABLED")
	}

	for _, route := range router.Routes() {
		fmt.Printf("%s %s → %s\n", route.Method, route.Path, route.Handler)
	}

	return &App{
		database: container.Database,
		router:   router,
	}
}

func (s *App) CloseDb() {
	s.database.CloseAll()
}

func (s *App) Run() error {
	return s.router.Run()
}

