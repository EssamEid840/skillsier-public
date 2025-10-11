# Microservice Template & Generator

## 🎯 Purpose

This template allows you to quickly create new microservices (jobs-be, proposals-be, contracts-be, reviews-be) following the same architecture as users-be.

---

## 📁 Standard Microservice Structure

```
<service-name>-be/
├── cmd/api/
│   └── main.go                       # Entry point
├── internal/
│   ├── domain/                       # Business entities
│   │   ├── <entity>/
│   │   │   ├── entity.go
│   │   │   └── repository.go
│   │   └── outbox/
│   │       ├── entity.go
│   │       └── repository.go
│   ├── application/                  # Business logic
│   │   ├── <entity>/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   └── eventhandler/
│   │       └── <event>_handler.go
│   ├── infrastructure/               # External dependencies
│   │   ├── persistence/postgres/
│   │   │   ├── connection.go
│   │   │   ├── migrations.go
│   │   │   ├── <entity>_repository.go
│   │   │   └── outbox_repository.go
│   │   ├── messaging/kafka/
│   │   │   ├── consumer.go
│   │   │   ├── producer.go
│   │   │   └── scram.go
│   │   └── outbox/
│   │       └── processor.go
│   ├── interfaces/                   # External interfaces
│   │   └── http/
│   │       ├── handlers/
│   │       │   └── <entity>_handler.go
│   │       ├── middleware/
│   │       │   ├── auth.go
│   │       │   ├── cors.go
│   │       │   └── logger.go
│   │       └── router.go
│   └── config/
│       └── config.go                 # Configuration
├── pkg/                              # Shared utilities
│   └── utils/
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       └── dapr-component.yaml
├── scripts/
│   ├── setup-local.sh
│   └── get-secrets.sh
├── .env.example
├── Dockerfile
├── Makefile
├── go.mod
├── go.sum
└── README.md
```

---

## 🛠️ Creating a New Microservice

### Step 1: Clone users-be as Template

```bash
# Clone users-be structure
cp -r users-be jobs-be

cd jobs-be

# Clean up
rm -rf bin/
rm -rf vendor/
```

### Step 2: Update Service Identifiers

**Search and replace in ALL files:**
- `users-be` → `jobs-be`
- `usersuser` → `jobsuser`
- `usersdb` → `jobsdb`
- `users-be-group` → `jobs-be-group`
- `user-events` → `job-events`

### Step 3: Update go.mod

```bash
cd jobs-be

# Update module name
sed -i 's/module users-be/module jobs-be/g' go.mod

# Clean up
go mod tidy
```

### Step 4: Define Domain Entity

Update `internal/domain/<entity>/entity.go` with your business entity.

---

## 📊 Example: jobs-be Implementation

### Domain Entity

```go
// internal/domain/job/entity.go
package job

import (
	"time"
	"github.com/google/uuid"
)

type JobStatus string
type JobType string
type BudgetType string

const (
	JobStatusOpen       JobStatus = "open"
	JobStatusInProgress JobStatus = "in_progress"
	JobStatusCompleted  JobStatus = "completed"
	JobStatusCancelled  JobStatus = "cancelled"
	
	JobTypeOneTime   JobType = "one_time"
	JobTypeOngoing   JobType = "ongoing"
	
	BudgetTypeFixed  BudgetType = "fixed_price"
	BudgetTypeHourly BudgetType = "hourly"
)

type Job struct {
	ID                uuid.UUID   `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	ClientID          uuid.UUID   `gorm:"type:uuid;not null;index" json:"client_id"`
	Title             string      `gorm:"type:varchar(200);not null" json:"title"`
	Description       string      `gorm:"type:text;not null" json:"description"`
	Category          string      `gorm:"type:varchar(100)" json:"category"`
	BudgetType        BudgetType  `gorm:"type:varchar(50);not null" json:"budget_type"`
	BudgetAmount      *float64    `gorm:"type:decimal(12,2)" json:"budget_amount,omitempty"`
	HourlyRateMin     *float64    `gorm:"type:decimal(10,2)" json:"hourly_rate_min,omitempty"`
	HourlyRateMax     *float64    `gorm:"type:decimal(10,2)" json:"hourly_rate_max,omitempty"`
	Duration          *string     `gorm:"type:varchar(100)" json:"duration,omitempty"`
	ExperienceLevel   string      `gorm:"type:varchar(50)" json:"experience_level"`
	Status            JobStatus   `gorm:"type:varchar(50);not null;default:'open'" json:"status"`
	ProposalCount     int         `gorm:"type:int;default:0" json:"proposal_count"`
	RequiredSkills    []JobSkill  `gorm:"foreignKey:JobID" json:"required_skills,omitempty"`
	CreatedAt         time.Time   `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time   `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	ClosedAt          *time.Time  `gorm:"type:timestamp" json:"closed_at,omitempty"`
}

type JobSkill struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	JobID  uuid.UUID `gorm:"type:uuid;not null;index" json:"job_id"`
	Name   string    `gorm:"type:varchar(100);not null" json:"name"`
	Level  string    `gorm:"type:varchar(50)" json:"level"`
}

func (Job) TableName() string {
	return "jobs"
}

func (JobSkill) TableName() string {
	return "job_skills"
}
```

### Repository Interface

```go
// internal/domain/job/repository.go
package job

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrJobNotFound = errors.New("job not found")
	ErrUnauthorized = errors.New("unauthorized access")
)

type Repository interface {
	Create(ctx context.Context, job *Job) error
	GetByID(ctx context.Context, id uuid.UUID) (*Job, error)
	List(ctx context.Context, filters ListFilters, limit, offset int) ([]*Job, int64, error)
	Update(ctx context.Context, job *Job) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByClientID(ctx context.Context, clientID uuid.UUID, limit, offset int) ([]*Job, int64, error)
}

type ListFilters struct {
	Category        *string
	BudgetType      *BudgetType
	MinBudget       *float64
	MaxBudget       *float64
	ExperienceLevel *string
	Status          *JobStatus
	Skills          []string
}
```

### Service Layer

```go
// internal/application/job/service.go
package job

import (
	"context"
	"encoding/json"
	"fmt"
	"jobs-be/internal/domain/job"
	"jobs-be/internal/domain/outbox"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	jobRepo    job.Repository
	outboxRepo outbox.Repository
	db         *gorm.DB
}

func NewService(jobRepo job.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{
		jobRepo:    jobRepo,
		outboxRepo: outboxRepo,
		db:         db,
	}
}

func (s *Service) CreateJob(ctx context.Context, clientID uuid.UUID, dto *CreateJobDTO) (*JobResponseDTO, error) {
	newJob := &job.Job{
		ClientID:        clientID,
		Title:           dto.Title,
		Description:     dto.Description,
		Category:        dto.Category,
		BudgetType:      job.BudgetType(dto.BudgetType),
		BudgetAmount:    dto.BudgetAmount,
		HourlyRateMin:   dto.HourlyRateMin,
		HourlyRateMax:   dto.HourlyRateMax,
		Duration:        dto.Duration,
		ExperienceLevel: dto.ExperienceLevel,
		Status:          job.JobStatusOpen,
	}

	// Create skills
	if len(dto.RequiredSkills) > 0 {
		for _, skill := range dto.RequiredSkills {
			newJob.RequiredSkills = append(newJob.RequiredSkills, job.JobSkill{
				Name:  skill.Name,
				Level: skill.Level,
			})
		}
	}

	// Transaction: Create job + outbox event
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.jobRepo.Create(ctx, newJob); err != nil {
			return err
		}

		event, err := s.createJobEvent("job.created", newJob)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(newJob), nil
}

func (s *Service) createJobEvent(eventType string, j *job.Job) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"job_id":      j.ID.String(),
		"client_id":   j.ClientID.String(),
		"title":       j.Title,
		"status":      string(j.Status),
		"budget_type": string(j.BudgetType),
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	metadata := map[string]interface{}{"source": "jobs-be"}
	metadataBytes, err := json.Marshal(metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal metadata: %w", err)
	}

	return &outbox.Event{
		AggregateID:   j.ID.String(),
		AggregateType: "job",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}
```

### HTTP Handler

```go
// internal/interfaces/http/handlers/job_handler.go
package handlers

import (
	"net/http"
	"jobs-be/internal/application/job"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type JobHandler struct {
	jobService *job.Service
}

func NewJobHandler(jobService *job.Service) *JobHandler {
	return &JobHandler{jobService: jobService}
}

func (h *JobHandler) CreateJob(c *gin.Context) {
	clientID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	var dto job.CreateJobDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.jobService.CreateJob(c.Request.Context(), clientID, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *JobHandler) ListJobs(c *gin.Context) {
	// Parse query parameters for filtering
	var filters job.ListFilters
	// ... parse filters from query params ...

	page := 1
	pageSize := 20
	// ... parse pagination ...

	jobs, total, err := h.jobService.ListJobs(c.Request.Context(), filters, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"jobs": jobs,
		"pagination": gin.H{
			"total": total,
			"page": page,
			"page_size": pageSize,
		},
	})
}
```

---

## 🗄️ Database Setup

### Create Database User and Database

```sql
-- In PostgreSQL (via K8s or local)
CREATE USER jobsuser WITH PASSWORD 'secure_password';
CREATE DATABASE jobsdb OWNER jobsuser;
GRANT ALL PRIVILEGES ON DATABASE jobsdb TO jobsuser;
```

### Environment Variables

```bash
# .env.example for jobs-be
JOBS_BE_APP_NAME=jobs-be
JOBS_BE_APP_ENVIRONMENT=development
JOBS_BE_APP_PORT=8081

# Database
JOBS_BE_DATABASE_HOST=173.212.218.251
JOBS_BE_DATABASE_PORT=30432
JOBS_BE_DATABASE_USER=jobsuser
JOBS_BE_DATABASE_DATABASE=jobsdb
JOBS_BE_DATABASE_SSL_MODE=disable
DB_PASSWORD=

# Kafka
JOBS_BE_KAFKA_BROKERS=173.212.218.251:31691
JOBS_BE_KAFKA_CONSUMER_GROUP=jobs-be-group
JOBS_BE_KAFKA_USER_EVENTS_TOPIC=user-events
JOBS_BE_KAFKA_JOB_EVENTS_TOPIC=job-events
KAFKA_PASSWORD=

# Keycloak
JOBS_BE_KEYCLOAK_URL=https://keycloak.skillsier.com
JOBS_BE_KEYCLOAK_REALM=skillsier
KEYCLOAK_CLIENT_ID=jobs-be
KEYCLOAK_CLIENT_SECRET=
```

---

## ☸️ Kubernetes Deployment

### Create K8s Secret for Keycloak Client

```bash
# In Keycloak Admin Console
# Create client: jobs-be
# Copy client secret

# Create K8s secret
kubectl create secret generic keycloak-client-jobs-be \
  --from-literal=client-id=jobs-be \
  --from-literal=client-secret=<your-secret>
```

### Deployment Manifest

```yaml
# deployments/k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: jobs-be
  labels:
    app: jobs-be
spec:
  replicas: 2
  selector:
    matchLabels:
      app: jobs-be
  template:
    metadata:
      labels:
        app: jobs-be
      annotations:
        dapr.io/enabled: "true"
        dapr.io/app-id: "jobs-be"
        dapr.io/app-port: "8081"
    spec:
      containers:
      - name: jobs-be
        image: skillsier/jobs-be:latest
        ports:
        - containerPort: 8081
        env:
        - name: JOBS_BE_APP_PORT
          value: "8081"
        # ... other env vars from configmap/secrets ...
```

---

## 🔄 Event Consumption

### Consuming User Events (Optional)

If your microservice needs to react to user events:

```go
// internal/application/eventhandler/user_handler.go
package eventhandler

import (
	"context"
	"encoding/json"
	"log"
)

type UserEventHandler struct {
	// your dependencies
}

func NewUserEventHandler() *UserEventHandler {
	return &UserEventHandler{}
}

func (h *UserEventHandler) HandleMessage(ctx context.Context, message []byte) error {
	var event struct {
		EventType string          `json:"event_type"`
		Payload   json.RawMessage `json:"payload"`
	}

	if err := json.Unmarshal(message, &event); err != nil {
		return err
	}

	switch event.EventType {
	case "user.created":
		return h.handleUserCreated(ctx, event.Payload)
	case "user.deleted":
		return h.handleUserDeleted(ctx, event.Payload)
	default:
		log.Printf("Unknown event type: %s", event.EventType)
	}

	return nil
}
```

---

## 🚀 Quick Creation Checklist

For each new microservice:

- [ ] 1. Copy users-be directory
- [ ] 2. Rename directory (jobs-be, proposals-be, etc.)
- [ ] 3. Update go.mod module name
- [ ] 4. Search/replace service identifiers
- [ ] 5. Define domain entities
- [ ] 6. Implement repositories
- [ ] 7. Implement services with outbox
- [ ] 8. Create HTTP handlers
- [ ] 9. Update router
- [ ] 10. Create database and user in PostgreSQL
- [ ] 11. Update .env.example
- [ ] 12. Create Keycloak client
- [ ] 13. Update K8s manifests
- [ ] 14. Test locally
- [ ] 15. Deploy to K8s

---

## 📊 Microservices Summary

### jobs-be
- **Port**: 8081
- **Database**: jobsdb
- **Events Published**: job.created, job.updated, job.deleted, job.closed
- **Events Consumed**: user.deleted (cascade delete jobs)

### proposals-be
- **Port**: 8082
- **Database**: proposalsdb
- **Events Published**: proposal.created, proposal.updated, proposal.withdrawn, proposal.accepted
- **Events Consumed**: job.deleted (cascade delete proposals)

### contracts-be
- **Port**: 8083
- **Database**: contractsdb
- **Events Published**: contract.created, contract.accepted, contract.completed, milestone.submitted, milestone.approved
- **Events Consumed**: proposal.accepted (create contract)

### reviews-be
- **Port**: 8084
- **Database**: reviewsdb
- **Events Published**: review.created
- **Events Consumed**: contract.completed (enable review creation)

---

## 🎯 Next Steps

Ready to create specific microservices? Let me know which one to generate first:

1. **jobs-be** - Complete implementation
2. **proposals-be** - Complete implementation
3. **contracts-be** - Complete implementation
4. **reviews-be** - Complete implementation

I can provide the complete code for each!