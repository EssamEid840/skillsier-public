package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	auth "users.go/m/internal/auth"
	konf "users.go/m/internal/infrastructure/config"
)

func getBool(cfg konf.Config, key string) bool {
	// Try to interpret the string value as a boolean.
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
		log.Fatal("config error: both Dapr and native Kafka enabled; enable only one")
	case !useDapr && !useKafka:
		log.Fatal("config error: neither Dapr nor Kafka enabled")
	}
}

func main() {
	cfg := konf.NewConfig(konf.ConfigProviderKoanfs)
	mustOnePath(cfg)

	issuer := cfg.GetString("auth.oidc.issuer")
	aud := cfg.GetString("auth.oidc.audience")
	jwtmw, err := auth.Middleware(issuer, aud)
	if err != nil {
		log.Fatalf("jwt middleware init failed: %v", err)
	}

	mux := http.NewServeMux()

	// Example protected endpoint
	mux.Handle("/v1/users", jwtmw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	})))

	port := cfg.GetString("http.port")
	if port == "" {
		port = ":8080"
	} else if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	log.Printf("API listening on %s", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
