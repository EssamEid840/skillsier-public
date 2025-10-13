# Getting Started with Users-BE

## Complete Setup Guide

This guide will walk you through setting up and running the users-be microservice in both local and production modes.

## Project Structure Overview

```
users-be/
├── cmd/api/main.go                          # Application entry point
├── internal/
│   ├── domain/                              # Business logic layer
│   │   ├── user/                            # User domain
│   │   │   ├── entity.go                    # User entity
│   │   │   └── repository.go                # Repository interface
│   │   └── outbox/                          # Outbox pattern
│   │       ├── entity.go                    # Event entity
│   │       └── repository.go                # Repository interface
│   ├── application/                         # Application services
│   │   ├── user/                            # User service
│   │   │   ├── service.go                   # Business logic
│   │   │   ├── dto.go                       # Data transfer objects
│   │   │   └── mapper.go                    # Entity-DTO mapping
│   │   └── eventhandler/                    # Event handlers
│   │       └── keycloak_handler.go          # Keycloak events
│   ├── infrastructure/                      # External integrations
│   │   ├── persistence/postgres/            # Database
│   │   │   ├── connection.go
│   │   │   ├── migrations.go
│   │   │   ├── user_repository.go
│   │   │   └── outbox_repository.go
│   │   ├── messaging/kafka/                 # Kafka
│   │   │   ├── consumer.go
│   │   │   ├── producer.go
│   │   │   └── scram.go
│   │   └── outbox/                          # Outbox processor
│   │       └── processor.go
│   ├── interfaces/http/                     # HTTP layer
│   │   ├── handlers/
│   │   │   ├── user_handler.go
│   │   │   └── health_handler.go
│   │   ├── middleware/
│   │   │   ├── logging.go
│   │   │   └── cors.go
│   │   └── router.go
│   └── config/                              # Configuration
│       └── config.go
├── deployments/k8s/                         # Kubernetes manifests
├── scripts/                                 # Helper scripts
├── Makefile                                 # Build automation
├── Dockerfile                               # Container image
├── go.mod                                   # Go dependencies
└── README.md                                # Documentation
```

## Step 1: Initial Setup

### 1.1 Create Keycloak Client

Before anything else, create the Keycloak client that users-be will use for authentication.

**Go to Keycloak Admin Console:**
```
<!-- http://173.212.218.251:30080 -->
https://keycloak.skillsier.com
```

**Create Client:**
1. Navigate to: **Clients** → **Create client**
2. **Client ID:** `users-be`
3. **Client authentication:** ON ✅
4. **Service accounts roles:** ON ✅
5. Save and go to **Credentials** tab
6. Copy the **Client Secret**

**Full instructions:** See [KEYCLOAK_SETUP.md](KEYCLOAK_SETUP.md)

### 1.2 Create Project Directory

```bash
mkdir -p users-be
cd users-be
```

### 1.3 Initialize Go Module

```bash
go mod init users-be
```

### 1.4 Copy All Source Files

Copy all the provided code files into their respective directories following the structure in FILE_STRUCTURE.md.

### 1.5 Install Dependencies

```bash
make deps
```

This will download all required Go packages with the latest secure versions.

## Step 2: Prepare Kubernetes Infrastructure

### 2.1 Ensure PostgreSQL is Running

```bash
# Check if PostgreSQL is deployed
kubectl get statefulset users-be-postgres -n skillsier

# If not, apply the PostgreSQL manifests
kubectl apply -f postgres-secret.yaml
kubectl apply -f postgres-statefulset.yaml
kubectl apply -f postgres-service.yaml
kubectl apply -f postgres-nodeport.yaml
```

### 2.2 Verify Kafka is Accessible

```bash
# Check Kafka service
kubectl get svc -n kafka | grep external

# Test Kafka connectivity
KAFKA_PASSWORD=$(kubectl get secret admin-user -n kafka -o jsonpath='{.data.password}' | base64 -d)
echo "test" | kcat -P \
  -b 173.212.218.251:31691 \
  -t smoke \
  -X security.protocol=SASL_SSL \
  -X sasl.mechanism=SCRAM-SHA-512 \
  -X sasl.username=admin-user \
  -X sasl.password=$KAFKA_PASSWORD \
  -X enable.ssl.certificate.verification=false
```

### 2.3 Verify Keycloak is Accessible

```bash
# Check Keycloak service
kubectl get svc keycloak-external -n keycloak

# Test accessibility
# curl -I http://173.212.218.251:30080/realms/skillsier
curl -I https://keycloak.skillsier.com/realms/skillsier
```

## Step 3: Setup Keycloak Client Secret

### 3.1 Export Client Credentials

Use the credentials you got from Keycloak in Step 1:

```bash
# Export the Keycloak client credentials
export KEYCLOAK_CLIENT_ID=users-be
export KEYCLOAK_CLIENT_SECRET=7X9kLmP4qR2wN8vT5yU3hJ6aZ1bC0dE  # Replace with your actual secret

# Optional: Set custom Keycloak URL
# export KEYCLOAK_URL=http://173.212.218.251:30080
export KEYCLOAK_URL=https://keycloak.skillsier.com

export KEYCLOAK_REALM=skillsier
```

### 3.2 Run Smart Secret Setup

```bash
# Make scripts executable
chmod +x scripts/*.sh

# Run the intelligent secret setup script
source scripts/get-secrets.sh
```

**What the script does:**
1. ✅ Checks if `keycloak-client-users-be` secret exists in K8s
2. ✅ If exists, verifies credentials work with Keycloak
3. ✅ If not exists or invalid, creates/updates from your env variables
4. ✅ Fetches all other secrets (PostgreSQL, Kafka)
5. ✅ Exports everything to your environment

**Expected output:**
```
Checking Keycloak client credentials...
⚠️  Keycloak secret does not exist, creating...
Found environment variables:
  KEYCLOAK_CLIENT_ID: users-be
  KEYCLOAK_CLIENT_SECRET: 7X9kLmP4qR***

Verifying credentials with Keycloak...
✓ Credentials verified successfully!

Creating Kubernetes secret...
secret/keycloak-client-users-be created
✓ Secret created successfully!

✓ All secrets loaded successfully!
========================================
  DB Password:         Redis***
  Kafka Password:      WNUy3***
  Keycloak Client ID:  users-be
  Keycloak Secret:     7X9kLmP4qR***
========================================
```

### 3.3 Verify Secret in Kubernetes

```bash
# Check the secret was created
kubectl get secret keycloak-client-users-be -n keycloak

# View secret details
kubectl describe secret keycloak-client-users-be -n keycloak
```

## Step 4: Run Locally

```bash
# Option 1: Run directly with Go
make run

# Option 2: Build and run binary
make build-local
./bin/users-be
```

### 3.4 Verify It's Working

```bash
# Health check
curl http://localhost:8080/health

# Expected output:
# {
#   "status": "healthy",
#   "database": "connected",
#   "timestamp": "2025-01-07T12:00:00Z"
# }
```

## Step 4: Test the Integration

### 4.1 Register a User in Keycloak

Go to Keycloak admin console and register a new user. The users-be service should automatically:
1. Receive the Keycloak event
2. Create the user in its database
3. Publish a user.created event to Kafka

### 4.2 Check Logs

You should see logs like:
```
Received Keycloak event from topic=keycloak-events
Processing Keycloak event: type=user_register user=john email=john@example.com
Successfully created user: id=... username=john email=john@example.com
Successfully published event ... to topic user-events
```

### 4.3 Verify User in Database

```bash
# Connect to PostgreSQL
PGPASSWORD=$(kubectl get secret users-be-postgres -n skillsier -o jsonpath='{.data.POSTGRES_PASSWORD}' | base64 -d) \
psql -h 173.212.218.251 -p 30432 -U usersuser -d usersdb

# Query users
SELECT id, username, email, keycloak_id, created_at FROM users;

# Check outbox events
SELECT id, event_type, aggregate_id, status, created_at FROM outbox_events ORDER BY created_at DESC LIMIT 10;
```

### 4.4 Verify Events in Kafka

```bash
# Consume from user-events topic
KAFKA_PASSWORD=$(kubectl get secret admin-user -n kafka -o jsonpath='{.data.password}' | base64 -d)

kcat -C \
  -b 173.212.218.251:31691 \
  -t user-events \
  -X security.protocol=SASL_SSL \
  -X sasl.mechanism=SCRAM-SHA-512 \
  -X sasl.username=admin-user \
  -X sasl.password=$KAFKA_PASSWORD \
  -X enable.ssl.certificate.verification=false \
  -o beginning
```

## Step 5: Test API Endpoints

### 5.1 Create User Manually (for testing)

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "keycloak_id": "test-sub-123",
    "username": "testuser",
    "email": "test@skillsier.com",
    "first_name": "Test",
    "last_name": "User",
    "email_verified": true
  }'
```

### 5.2 Get User by Keycloak ID

```bash
curl http://localhost:8080/api/v1/users/keycloak/test-sub-123
```

### 5.3 Update User Profile

```bash
# Get user ID from previous response
USER_ID="..."

curl -X PUT http://localhost:8080/api/v1/users/$USER_ID \
  -H "Content-Type: application/json" \
  -