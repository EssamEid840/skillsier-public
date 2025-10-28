ARCHITECTURE RULES:
1. NO direct Keycloak imports - use pkg/auth only
2. NO copy-pasted middleware - use platform-shared
3. NO magic strings for Dapr components - use constants  
4. NO event schema definitions in services - use contracts/events
5. EVERY service MUST use the standard config loader
6. shared packages MUST remain "leaf" dependencies (no service imports)
7. Roll  Back strategy
Final punch-list
----------------

*   **Health & shutdown**
    
    *   Each service exposes /healthz + readiness/liveness.
        
    *   Graceful shutdown with timeouts (Read/Write/IdleTimeout), HTTP clients have sane timeouts.
        
*   **Config parity**
    
    *   Single loader (flags → env → file → defaults) wired in **every** service.
        
    *   One --config flag everywhere; CONFIGURATION.md present per service.
        
*   **Auth & RBAC**
    
    *   Shared pkg/auth in use (no per-service JWT parsing left).
        
    *   Issuer/audience enforced; role→permission mapping centralized; /me returns normalized principal.
        
*   **Cross-cutting utils**
    
    *   Logging/tracing/error envelope/pagination via platform-shared/ (local duplicates removed).
        
    *   Request ID middleware on all entrypoints; logs include req\_id.
        
*   **Dapr components**
    
    *   dapr/components/ co-located per service, with scopes set and secrets via secretKeyRef.
        
    *   Code uses **constants** for component/topic names (no magic strings).
        
*   **Eventing**
    
    *   **contracts/events** module exists; topics + protobuf/avro types generated.
        
    *   EVENTS.md lists topics, keys, owners, and BC policy; publishers/consumers use generated types.
        
*   **Outbox (Watermill)**
    
    *   TX publisher used inside DB transactions; forwarder running (or deployable) and monitored.
        
    *   Schema initialized at boot/forwarder (not inside TX). Retry/backoff metrics visible.
        
*   **Kafka hygiene**
    
    *   Topic configs defined declaratively (partitions, retention/compaction, ACLs).
        
    *   Producers use keyed partitioning; consumers idempotent (dedupe by message ID).
        
*   **Security & runtime hardening**
    
    *   K8s: resource requests/limits, PodSecurity, runAsNonRoot, readOnlyRootFS, NetworkPolicy.
        
    *   Secrets via SealedSecrets/Vault; no secrets in images or ConfigMaps.
        
    *   Key/Cert rotation plan (Keycloak, TLS, Kafka SCRAM).
        
*   **Observability**
    
    *   RED metrics + traces for HTTP/Kafka; /metrics exposed; dashboards for:
        
        *   error rate, latency, RPS
            
        *   db saturation
            
        *   outbox queue length / forwarder lag
            
        *   Kafka consumer lag
            
    *   Alerts for SLO breaches (latency, 5xx, consumer lag, outbox retries).
        
*   **Data & migrations**
    
    *   Deterministic, reversible migrations; rollback tested.
        
    *   Backups + restore test for Postgres; retention documented.

    *   Migrations & data policy. You have migrator commands—great. Add conventions: deterministic ordering, idempotent seeds, and rollbacks per service. Consider golang-migrate with Makefile targets (migrate.up/down). 

        
*   **CI/CD & supply chain**
    
    *   Pipelines: fmt/lint/tests → build → SBOM & image scan → sign (cosign) → deploy.
        
    *   Per-service SemVer + CHANGELOG; prod manifests are GitOps-managed.
        
*   **Docs & runbooks**
    
    *   ARCHITECTURE.md (high level), RUNBOOK.md per service (how to debug, rotate secrets, restart forwarder).
        
    *   DECISIONS/ ADRs for auth, outbox, Dapr, contracts.

    ----------------------------------------------
*  **Makefile & scripts.**

    *   Standardize targets: run, test, lint, migrate.up/down, docker.build, k8s.apply, k8s.rollout. Reuse a common make/include/common.mk.
------------------------------------------
### Examples
6. Config standardization. Unify config structs (Viper/env), default overrides, and a single --config flag in each cmd/*. Add a CONFIGURATION.md with env var names and precedence.

Add this in each folder structure and write note to exlain and force me do that
├── internal/
│   └── config/
│       ├── schema.go        # typed Config struct used by the service
│       └── docs/
│           └── CONFIGURATION.md


example
Typed config (one struct per service)

internal/config/schema.go

package config

import "time"

type Config struct {
  App struct {
    Name string `mapstructure:"name"` // e.g. "users-be"
    Env  string `mapstructure:"env"`  // dev|staging|prod
  } `mapstructure:"app"`

  Server struct {
    Port         int           `mapstructure:"port"`           // 8080
    ReadTimeout  time.Duration `mapstructure:"read_timeout"`   // "10s"
    WriteTimeout time.Duration `mapstructure:"write_timeout"`  // "10s"
  } `mapstructure:"server"`

  Postgres struct {
    DSN           string `mapstructure:"dsn"`
    MaxOpenConns  int    `mapstructure:"max_open_conns"`
    MaxIdleConns  int    `mapstructure:"max_idle_conns"`
  } `mapstructure:"postgres"`

  Kafka struct {
    Brokers     []string `mapstructure:"brokers"`
    SASLUser    string   `mapstructure:"sasl_user"`
    SASLPass    string   `mapstructure:"sasl_pass"`
    TLSCAPath   string   `mapstructure:"tls_ca_path"`
    ClientID    string   `mapstructure:"client_id"`
  } `mapstructure:"kafka"`

  Auth struct {
    IssuerURL   string        `mapstructure:"issuer_url"`
    Audience    string        `mapstructure:"audience"`
    JWKSURL     string        `mapstructure:"jwks_url"`
    AllowedAlgs []string      `mapstructure:"allowed_algs"`
    CacheTTL    time.Duration `mapstructure:"cache_ttl"`
    ClockSkew   time.Duration `mapstructure:"clock_skew"`
  } `mapstructure:"auth"`

  Telemetry struct {
    PrometheusEnabled bool   `mapstructure:"prometheus_enabled"`
    OTLPEndpoint      string `mapstructure:"otlp_endpoint"`
    LogLevel          string `mapstructure:"log_level"` // debug|info|warn|error
  } `mapstructure:"telemetry"`

  Dapr struct {
    Enabled     bool   `mapstructure:"enabled"`
    PubsubName  string `mapstructure:"pubsub_name"`
  } `mapstructure:"dapr"`
}

Default file (checked into Git)

config/default.yaml

app:
  name: users-be
  env: dev

server:
  port: 8080
  read_timeout: 10s
  write_timeout: 10s

postgres:
  dsn: "postgres://usersuser:ChangeMe@localhost:5432/usersdb?sslmode=disable"
  max_open_conns: 10
  max_idle_conns: 5

kafka:
  brokers: ["localhost:9092"]
  client_id: "users-be"

auth:
  issuer_url: "https://keycloak.example.com/realms/skillsier"
  audience: "users-be"
  jwks_url: "https://keycloak.example.com/realms/skillsier/protocol/openid-connect/certs"
  allowed_algs: ["RS256"]
  cache_ttl: 10m
  clock_skew: 60s

telemetry:
  prometheus_enabled: true
  otlp_endpoint: ""
  log_level: debug

dapr:
  enabled: true
  pubsub_name: "kafka-pubsub"

One loader (Viper + pflag), same in every service

In each cmd/api/main.go:

package main

import (
  "embed"
  "fmt"
  "log"
  "os"
  "strings"

  "github.com/spf13/pflag"
  "github.com/spf13/viper"

  svc "users-be/internal/config"
)

var (
  //go:embed ../../config/*.yaml
  defaultsFS embed.FS
)

func loadConfig() (svc.Config, error) {
  var cfg svc.Config

  // 1) flags
  var cfgPath string
  pflag.StringVar(&cfgPath, "config", "", "Path to config file (YAML/TOML/JSON)")
  pflag.Parse()

  // 2) viper setup
  v := viper.New()
  v.SetConfigType("yaml")
  v.SetEnvPrefix("USERS_BE") // <- change per service
  v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
  v.AutomaticEnv()

  // 3) embedded defaults
  if b, err := defaultsFS.ReadFile("../../config/default.yaml"); err == nil {
    if err := v.MergeConfig(strings.NewReader(string(b))); err != nil {
      return cfg, err
    }
  }

  // 4) explicit config file (if provided)
  if cfgPath != "" {
    v.SetConfigFile(cfgPath)
    if err := v.MergeInConfig(); err != nil {
      return cfg, fmt.Errorf("read --config: %w", err)
    }
  } else {
    // fallback search: ./config/<env>.yaml or ./config/local.yaml
    // (env may come from defaults or env vars)
    env := v.GetString("app.env")
    if env == "" { env = "dev" }
    v.AddConfigPath(".")
    v.AddConfigPath("./config")
    v.SetConfigName(env) // dev.yaml / prod.yaml if present
    _ = v.MergeInConfig() // ok if not found
  }

  // 5) flags override (bind after Merge so flags win)
  _ = v.BindPFlags(pflag.CommandLine)

  // 6) unmarshal
  if err := v.Unmarshal(&cfg); err != nil {
    return cfg, err
  }
  return cfg, nil
}

func main() {
  cfg, err := loadConfig()
  if err != nil {
    log.Fatalf("config error: %v", err)
  }
  // Now use cfg.Server.Port, cfg.Postgres.DSN, cfg.Auth.*, etc.
  _ = cfg
  // ...
}


ENV naming rule (automatic via viper):
USERS_BE_POSTGRES_DSN, USERS_BE_SERVER_PORT, USERS_BE_AUTH_JWKS_URL, etc.
(Just uppercase and replace dots with underscores.)



