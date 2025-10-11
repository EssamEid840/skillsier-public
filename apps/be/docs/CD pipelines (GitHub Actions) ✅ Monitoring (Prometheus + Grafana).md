# ==========================================
# FILE: .github/workflows/users-be.yml
# GitHub Actions CI/CD for users-be
# ==========================================

name: users-be CI/CD

on:
  push:
    branches: [ main, develop ]
    paths:
      - 'users-be/**'
      - '.github/workflows/users-be.yml'
  pull_request:
    branches: [ main ]
    paths:
      - 'users-be/**'

env:
  REGISTRY: ghcr.io
  IMAGE_NAME: ${{ github.repository }}/users-be

jobs:
  test:
    name: Test
    runs-on: ubuntu-latest
    
    services:
      postgres:
        image: postgres:16
        env:
          POSTGRES_PASSWORD: postgres
          POSTGRES_DB: testdb
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          - 5432:5432

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.23'
          cache-dependency-path: users-be/go.sum

      - name: Install dependencies
        working-directory: ./users-be
        run: go mod download

      - name: Run tests
        working-directory: ./users-be
        env:
          USERS_BE_DATABASE_HOST: localhost
          USERS_BE_DATABASE_PORT: 5432
          USERS_BE_DATABASE_USER: postgres
          USERS_BE_DATABASE_DATABASE: testdb
          DB_PASSWORD: postgres
        run: |
          go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
          go tool cover -html=coverage.out -o coverage.html

      - name: Upload coverage
        uses: codecov/codecov-action@v3
        with:
          files: ./users-be/coverage.out
          flags: users-be

      - name: Run linter
        working-directory: ./users-be
        run: |
          go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
          golangci-lint run ./...

  build:
    name: Build & Push Docker Image
    runs-on: ubuntu-latest
    needs: test
    if: github.event_name == 'push'
    
    permissions:
      contents: read
      packages: write

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3

      - name: Log in to Container Registry
        uses: docker/login-action@v3
        with:
          registry: ${{ env.REGISTRY }}
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}

      - name: Extract metadata
        id: meta
        uses: docker/metadata-action@v5
        with:
          images: ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}
          tags: |
            type=ref,event=branch
            type=sha,prefix={{branch}}-
            type=semver,pattern={{version}}
            type=raw,value=latest,enable={{is_default_branch}}

      - name: Build and push
        uses: docker/build-push-action@v5
        with:
          context: ./users-be
          push: true
          tags: ${{ steps.meta.outputs.tags }}
          labels: ${{ steps.meta.outputs.labels }}
          cache-from: type=gha
          cache-to: type=gha,mode=max

  deploy-staging:
    name: Deploy to Staging
    runs-on: ubuntu-latest
    needs: build
    if: github.ref == 'refs/heads/develop'
    environment:
      name: staging
      url: https://staging-api.skillsier.com

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Setup kubectl
        uses: azure/setup-kubectl@v3

      - name: Configure kubectl
        env:
          KUBECONFIG_CONTENT: ${{ secrets.KUBECONFIG_STAGING }}
        run: |
          mkdir -p ~/.kube
          echo "$KUBECONFIG_CONTENT" | base64 -d > ~/.kube/config

      - name: Deploy to staging
        working-directory: ./users-be/deployments/k8s
        run: |
          kubectl apply -f . -n staging
          kubectl set image deployment/users-be \
            users-be=${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}:develop-${{ github.sha }} \
            -n staging
          kubectl rollout status deployment/users-be -n staging --timeout=5m

      - name: Run smoke tests
        run: |
          sleep 30
          curl -f https://staging-api.skillsier.com/health || exit 1

  deploy-production:
    name: Deploy to Production
    runs-on: ubuntu-latest
    needs: build
    if: github.ref == 'refs/heads/main'
    environment:
      name: production
      url: https://api.skillsier.com

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Setup kubectl
        uses: azure/setup-kubectl@v3

      - name: Configure kubectl
        env:
          KUBECONFIG_CONTENT: ${{ secrets.KUBECONFIG_PRODUCTION }}
        run: |
          mkdir -p ~/.kube
          echo "$KUBECONFIG_CONTENT" | base64 -d > ~/.kube/config

      - name: Deploy to production
        working-directory: ./users-be/deployments/k8s
        run: |
          kubectl apply -f . -n production
          kubectl set image deployment/users-be \
            users-be=${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}:latest \
            -n production
          kubectl rollout status deployment/users-be -n production --timeout=10m

      - name: Run smoke tests
        run: |
          sleep 30
          curl -f https://api.skillsier.com/health || exit 1

      - name: Notify deployment
        uses: 8398a7/action-slack@v3
        with:
          status: ${{ job.status }}
          text: 'users-be deployed to production'
          webhook_url: ${{ secrets.SLACK_WEBHOOK }}
        if: always()

# ==========================================
# FILE: .github/workflows/all-services.yml
# Unified workflow for all microservices
# ==========================================

name: All Services CI/CD

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  detect-changes:
    runs-on: ubuntu-latest
    outputs:
      users-be: ${{ steps.filter.outputs.users-be }}
      jobs-be: ${{ steps.filter.outputs.jobs-be }}
      proposals-be: ${{ steps.filter.outputs.proposals-be }}
      contracts-be: ${{ steps.filter.outputs.contracts-be }}
      reviews-be: ${{ steps.filter.outputs.reviews-be }}
    steps:
      - uses: actions/checkout@v4
      - uses: dorny/paths-filter@v2
        id: filter
        with:
          filters: |
            users-be:
              - 'users-be/**'
            jobs-be:
              - 'jobs-be/**'
            proposals-be:
              - 'proposals-be/**'
            contracts-be:
              - 'contracts-be/**'
            reviews-be:
              - 'reviews-be/**'

  test-users-be:
    needs: detect-changes
    if: needs.detect-changes.outputs.users-be == 'true'
    uses: ./.github/workflows/users-be.yml

  test-jobs-be:
    needs: detect-changes
    if: needs.detect-changes.outputs.jobs-be == 'true'
    uses: ./.github/workflows/jobs-be.yml

  # Similar for other services...

# ==========================================
# FILE: deployments/monitoring/prometheus-config.yml
# Prometheus configuration for metrics
# ==========================================

global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: 'users-be'
    kubernetes_sd_configs:
      - role: pod
        namespaces:
          names:
            - default
            - production
    relabel_configs:
      - source_labels: [__meta_kubernetes_pod_label_app]
        action: keep
        regex: users-be
      - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_scrape]
        action: keep
        regex: true
      - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_path]
        action: replace
        target_label: __metrics_path__
        regex: (.+)
      - source_labels: [__address__, __meta_kubernetes_pod_annotation_prometheus_io_port]
        action: replace
        regex: ([^:]+)(?::\d+)?;(\d+)
        replacement: $1:$2
        target_label: __address__

  - job_name: 'jobs-be'
    kubernetes_sd_configs:
      - role: pod
    relabel_configs:
      - source_labels: [__meta_kubernetes_pod_label_app]
        action: keep
        regex: jobs-be

  - job_name: 'proposals-be'
    kubernetes_sd_configs:
      - role: pod
    relabel_configs:
      - source_labels: [__meta_kubernetes_pod_label_app]
        action: keep
        regex: proposals-be

  - job_name: 'contracts-be'
    kubernetes_sd_configs:
      - role: pod
    relabel_configs:
      - source_labels: [__meta_kubernetes_pod_label_app]
        action: keep
        regex: contracts-be

  - job_name: 'reviews-be'
    kubernetes_sd_configs:
      - role: pod
    relabel_configs:
      - source_labels: [__meta_kubernetes_pod_label_app]
        action: keep
        regex: reviews-be

  - job_name: 'postgres'
    static_configs:
      - targets: ['postgres-exporter:9187']

  - job_name: 'kafka'
    static_configs:
      - targets: ['kafka-exporter:9308']

# ==========================================
# FILE: deployments/monitoring/prometheus-deployment.yml
# Deploy Prometheus to K8s
# ==========================================

apiVersion: v1
kind: ConfigMap
metadata:
  name: prometheus-config
  namespace: monitoring
data:
  prometheus.yml: |
    # Paste prometheus-config.yml content here

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: prometheus
  namespace: monitoring
spec:
  replicas: 1
  selector:
    matchLabels:
      app: prometheus
  template:
    metadata:
      labels:
        app: prometheus
    spec:
      serviceAccountName: prometheus
      containers:
      - name: prometheus
        image: prom/prometheus:latest
        args:
          - '--config.file=/etc/prometheus/prometheus.yml'
          - '--storage.tsdb.path=/prometheus'
          - '--storage.tsdb.retention.time=15d'
        ports:
        - containerPort: 9090
        volumeMounts:
        - name: config
          mountPath: /etc/prometheus
        - name: storage
          mountPath: /prometheus
      volumes:
      - name: config
        configMap:
          name: prometheus-config
      - name: storage
        persistentVolumeClaim:
          claimName: prometheus-storage

---
apiVersion: v1
kind: Service
metadata:
  name: prometheus
  namespace: monitoring
spec:
  selector:
    app: prometheus
  ports:
  - port: 9090
    targetPort: 9090
  type: ClusterIP

---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: prometheus-storage
  namespace: monitoring
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 50Gi

---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: prometheus
  namespace: monitoring

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: prometheus
rules:
- apiGroups: [""]
  resources:
  - nodes
  - nodes/proxy
  - services
  - endpoints
  - pods
  verbs: ["get", "list", "watch"]
- apiGroups:
  - extensions
  resources:
  - ingresses
  verbs: ["get", "list", "watch"]

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: prometheus
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: prometheus
subjects:
- kind: ServiceAccount
  name: prometheus
  namespace: monitoring

# ==========================================
# FILE: deployments/monitoring/grafana-deployment.yml
# Deploy Grafana for visualization
# ==========================================

apiVersion: apps/v1
kind: Deployment
metadata:
  name: grafana
  namespace: monitoring
spec:
  replicas: 1
  selector:
    matchLabels:
      app: grafana
  template:
    metadata:
      labels:
        app: grafana
    spec:
      containers:
      - name: grafana
        image: grafana/grafana:latest
        ports:
        - containerPort: 3000
        env:
        - name: GF_SECURITY_ADMIN_PASSWORD
          valueFrom:
            secretKeyRef:
              name: grafana-secret
              key: admin-password
        - name: GF_INSTALL_PLUGINS
          value: "grafana-piechart-panel,grafana-worldmap-panel"
        volumeMounts:
        - name: storage
          mountPath: /var/lib/grafana
        - name: datasources
          mountPath: /etc/grafana/provisioning/datasources
        - name: dashboards-config
          mountPath: /etc/grafana/provisioning/dashboards
        - name: dashboards
          mountPath: /var/lib/grafana/dashboards
      volumes:
      - name: storage
        persistentVolumeClaim:
          claimName: grafana-storage
      - name: datasources
        configMap:
          name: grafana-datasources
      - name: dashboards-config
        configMap:
          name: grafana-dashboards-config
      - name: dashboards
        configMap:
          name: grafana-dashboards

---
apiVersion: v1
kind: Service
metadata:
  name: grafana
  namespace: monitoring
spec:
  selector:
    app: grafana
  ports:
  - port: 3000
    targetPort: 3000
  type: LoadBalancer

---
apiVersion: v1
kind: ConfigMap
metadata:
  name: grafana-datasources
  namespace: monitoring
data:
  datasources.yml: |
    apiVersion: 1
    datasources:
    - name: Prometheus
      type: prometheus
      access: proxy
      url: http://prometheus:9090
      isDefault: true

---
apiVersion: v1
kind: ConfigMap
metadata:
  name: grafana-dashboards-config
  namespace: monitoring
data:
  dashboards.yml: |
    apiVersion: 1
    providers:
    - name: 'default'
      orgId: 1
      folder: ''
      type: file
      disableDeletion: false
      updateIntervalSeconds: 10
      options:
        path: /var/lib/grafana/dashboards

# ==========================================
# FILE: pkg/metrics/metrics.go
# Prometheus metrics for Go services
# ==========================================

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// HTTP metrics
	HttpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"service", "method", "endpoint", "status"},
	)

	HttpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"service", "method", "endpoint"},
	)

	// Database metrics
	DatabaseQueriesTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "database_queries_total",
			Help: "Total number of database queries",
		},
		[]string{"service", "operation", "table"},
	)

	DatabaseQueryDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "database_query_duration_seconds",
			Help:    "Database query duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"service", "operation", "table"},
	)

	// Kafka metrics
	KafkaMessagesProduced = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "kafka_messages_produced_total",
			Help: "Total number of Kafka messages produced",
		},
		[]string{"service", "topic"},
	)

	KafkaMessagesConsumed = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "kafka_messages_consumed_total",
			Help: "Total number of Kafka messages consumed",
		},
		[]string{"service", "topic"},
	)

	// Business metrics
	JobsCreated = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "jobs_created_total",
			Help: "Total number of jobs created",
		},
	)

	ProposalsSubmitted = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "proposals_submitted_total",
			Help: "Total number of proposals submitted",
		},
	)

	ContractsCreated = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "contracts_created_total",
			Help: "Total number of contracts created",
		},
	)

	ReviewsCreated = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "reviews_created_total",
			Help: "Total number of reviews created",
		},
	)
)

# ==========================================
# FILE: internal/interfaces/http/middleware/metrics.go
# Middleware to collect HTTP metrics
# ==========================================

package middleware

import (
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"your-service/pkg/metrics"
)

func MetricsMiddleware(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Record metrics
		duration := time.Since(start).Seconds()
		status := strconv.Itoa(c.Writer.Status())

		metrics.HttpRequestsTotal.WithLabelValues(
			serviceName,
			c.Request.Method,
			c.FullPath(),
			status,
		).Inc()

		metrics.HttpRequestDuration.WithLabelValues(
			serviceName,
			c.Request.Method,
			c.FullPath(),
		).Observe(duration)
	}
}

# ==========================================
# USAGE IN MAIN.GO
# ==========================================

/*
import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"your-service/internal/interfaces/http/middleware"
)

// In main.go:

// Add metrics middleware
router.Use(middleware.MetricsMiddleware("users-be"))

// Add metrics endpoint
router.GET("/metrics", gin.WrapH(promhttp.Handler()))
*/