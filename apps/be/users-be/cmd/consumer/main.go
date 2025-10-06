package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	konf "users.go/m/internal/infrastructure/config"
)

type CloudEvent struct {
	ID          string          `json:"id"`
	Source      string          `json:"source"`
	SpecVersion string          `json:"specversion"`
	Type        string          `json:"type"`
	Data        json.RawMessage `json:"data"`
	Subject     string          `json:"subject,omitempty"`
	Time        string          `json:"time,omitempty"`
}

func getBool(cfg konf.Config, key string) bool {
	s := strings.TrimSpace(strings.ToLower(cfg.GetString(key)))
	switch s {
	case "1", "t", "true", "y", "yes", "on":
		return true
	case "0", "f", "false", "n", "no", "off", "":
		return false
	}
	if n, err := strconv.Atoi(s); err == nil {
		return n != 0
	}
	return false
}

func mustOnePath(cfg konf.Config) {
	useDapr := getBool(cfg, "dapr.enabled")
	useKafka := getBool(cfg, "kafka.enabled")

	switch {
	case useDapr && useKafka:
		log.Fatal("config error: both Dapr and native Kafka enabled")
	case !useDapr && !useKafka:
		log.Fatal("config error: neither Dapr nor native Kafka enabled")
	}
}

func main() {
	cfg := konf.NewConfig(konf.ConfigProviderKoanfs)
	mustOnePath(cfg)

	useDapr := getBool(cfg, "dapr.enabled")
	if useDapr {
		appPort := cfg.GetString("dapr.consumer.app_port")
		if appPort == "" {
			appPort = "8082"
		}
		if !strings.HasPrefix(appPort, ":") {
			appPort = ":" + appPort
		}

		pubsub := cfg.GetString("dapr.pubsub")
		if pubsub == "" {
			pubsub = "kafka-pubsub"
		}

		// Topic here is only for logs; actual routing is via Dapr Subscriptions YAML.
		topic := cfg.GetString("kafka.topic")
		if topic == "" {
			topic = "skillsier.user-events"
		}
		log.Printf("Dapr subscriber on %s (pubsub=%s topic=%s)", appPort, pubsub, topic)

		mux := http.NewServeMux()

		mux.HandleFunc("/events/users", func(w http.ResponseWriter, r *http.Request) {
			var evt CloudEvent
			if err := json.NewDecoder(r.Body).Decode(&evt); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// TODO: handle users topic payload in evt.Data
			w.WriteHeader(http.StatusOK)
		})

		mux.HandleFunc("/events/keycloak/realm", func(w http.ResponseWriter, r *http.Request) {
			var evt CloudEvent
			if err := json.NewDecoder(r.Body).Decode(&evt); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// TODO: handle Keycloak realm events in evt.Data
			w.WriteHeader(http.StatusOK)
		})

		mux.HandleFunc("/events/keycloak/admin", func(w http.ResponseWriter, r *http.Request) {
			var evt CloudEvent
			if err := json.NewDecoder(r.Body).Decode(&evt); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// TODO: handle Keycloak admin events in evt.Data
			w.WriteHeader(http.StatusOK)
		})

		log.Fatal(http.ListenAndServe(appPort, mux))
		return
	}

	// Native Kafka path intentionally omitted for remote-dev
}
