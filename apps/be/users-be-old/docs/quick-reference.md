# Quick Reference Card - users-be

## 🚀 Quick Start (Copy-Paste)

```bash
# 1. Create Keycloak client 'users-be' in admin console
#    Copy the client secret

# 2. Setup in one command block
export KEYCLOAK_CLIENT_ID=users-be
export KEYCLOAK_CLIENT_SECRET=your-secret-from-keycloak
cd users-be
go mod download && go mod tidy
chmod +x scripts/*.sh
source scripts/get-secrets.sh
make run
```

## 📋 Keycloak Client Settings

```
Client ID: users-be
Client authentication: ON ✅
Service accounts roles: ON ✅
Direct access grants: ON ✅ (optional)
Standard flow: OFF
```

## 🔑 Secret Management Commands

```bash
# Create/Update Keycloak secret
export KEYCLOAK_CLIENT_ID=users-be
export KEYCLOAK_CLIENT_SECRET=xxx
source scripts/get-secrets.sh

# Verify secret exists
kubectl get secret keycloak-client-users-be -n keycloak

# View secret
kubectl get secret keycloak-client-users-be -n keycloak -o yaml

# Delete secret
kubectl delete secret keycloak-client-users-be -n keycloak

# Manual creation (if needed)
kubectl create secret generic keycloak-client-users-be -n keycloak \
  --from-literal=client-id=users-be \
  --from-literal=client-secret=xxx \
  --dry-run=client -o yaml | kubectl apply -f -
```

## 🛠️ Development Commands

```bash
# Setup
make deps                    # Install dependencies
make setup-local            # Setup local environment
source scripts/get-secrets.sh  # Load secrets

# Build & Run
make build-local            # Build binary
make run                    # Run application
make dev                    # Run with hot reload

# Test
make test                   # Run tests
make test-integration       # Integration tests
make lint                   # Run linter
make fmt                    # Format code

# Docker
make docker-build          # Build image
make docker-push           # Push to registry
make docker-run            # Run in Docker

# Kubernetes
make k8s-deploy            # Deploy to K8s
make k8s-delete            # Delete from K8s
make k8s-logs              # View logs
make k8s-status            # Check status
make k8s-restart           # Restart deployment
```

## 🔍 Testing & Verification

```bash
# Health check
curl http://localhost:8080/health

# Create user
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "keycloak_id": "test-123",
    "username": "testuser",
    "email": "test@example.com"
  }'

# Get user by Keycloak ID
curl http://localhost:8080/api/v1/users/keycloak/test-123

# List users
curl "http://localhost:8080/api/v1/users?page=1&page_size=10"

# Update user
curl -X PUT http://localhost:8080/api/v1/users/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "profile_type": "freelancer",
    "hourly_rate": 50.00
  }'
```

## 🗄️ Database Commands

```bash
# Connect to PostgreSQL
PGPASSWORD=$(kubectl get secret users-be-postgres -n skillsier \
  -o jsonpath='{.data.POSTGRES_PASSWORD}' | base64 -d) \
psql -h 173.212.218.251 -p 30432 -U usersuser -d usersdb

# View users
SELECT * FROM users ORDER BY created_at DESC LIMIT 10;

# View outbox events
SELECT id, event_type, status, created_at 
FROM outbox_events 
ORDER BY created_at DESC LIMIT 10;

# Check pending events
SELECT COUNT(*) FROM outbox_events WHERE status='pending';
```

## 📨 Kafka Commands

```bash
# Get Kafka password
KAFKA_PASSWORD=$(kubectl get secret admin-user -n kafka \
  -o jsonpath='{.data.password}' | base64 -d)

# Consume from keycloak-events
kcat -C \
  -b 173.212.218.251:31691 \
  -t keycloak-events \
  -X security.protocol=SASL_SSL \
  -X sasl.mechanism=SCRAM-SHA-512 \
  -X sasl.username=admin-user \
  -X sasl.password=$KAFKA_PASSWORD \
  -X enable.ssl.certificate.verification=false \
  -o beginning

# Consume from user-events
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

## 🐛 Troubleshooting

```bash
# Check all services
kubectl get pods -n skillsier
kubectl get pods -n kafka
kubectl get pods -n keycloak

# View application logs
kubectl logs -f -n skillsier -l app=users-be

# Check database connection
kubectl exec -it users-be-postgres-0 -n skillsier -- \
  psql -U usersuser -d usersdb -c "SELECT version();"

# Test Keycloak connectivity
# curl -I http://173.212.218.251:30080/realms/skillsier
curl -I https://keycloak.skillsier.com/realms/skillsier


# Test Kafka connectivity
telnet 173.212.218.251 31691

# View pod details
kubectl describe pod -n skillsier users-be-xxx-xxx

# View events
kubectl get events -n skillsier --sort-by='.lastTimestamp'
```

## 🔐 Keycloak Client Verification

```bash
# Test client credentials
# KEYCLOAK_URL=http://173.212.218.251:30080
KEYCLOAK_URL=https://keycloak.skillsier.com

REALM=skillsier

TOKEN_RESPONSE=$(curl -s -X POST \
  "${KEYCLOAK_URL}/realms/${REALM}/protocol/openid-connect/token" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "grant_type=client_credentials" \
  -d "client_id=users-be" \
  -d "client_secret=$KEYCLOAK_CLIENT_SECRET")

# Extract token
ACCESS_TOKEN=$(echo $TOKEN_RESPONSE | grep -o '"access_token":"[^"]*' | cut -d'"' -f4)

echo "Token: ${ACCESS_TOKEN:0:50}..."
# If you get a token, credentials are valid!
```

## 📊 Monitoring

```bash
# Application health
curl http://localhost:8080/health
curl http://localhost:8080/ready
curl http://localhost:8080/live

# Database stats
psql -h 173.212.218.251 -p 30432 -U usersuser -d usersdb -c \
  "SELECT 
     (SELECT COUNT(*) FROM users) as total_users,
     (SELECT COUNT(*) FROM users WHERE is_active=true) as active_users,
     (SELECT COUNT(*) FROM outbox_events WHERE status='pending') as pending_events;"

# Check consumer group lag
kafka-consumer-groups.sh --bootstrap-server 173.212.218.251:31691 \
  --group users-be-group --describe \
  --command-config kafka.properties
```

## 📁 File Locations

```
Configuration:
  - .env.example               Environment template
  - internal/config/config.go  Config loading
  
Scripts:
  - scripts/get-secrets.sh     Smart secret management
  - scripts/setup-local.sh     Local environment setup
  
K8s Manifests:
  - deployments/k8s/           All K8s resources
  
Documentation:
  - README.md                  Main documentation
  - GETTING_STARTED.md         Setup guide
  - KEYCLOAK_SETUP.md          Keycloak configuration
  - QUICK_FIX.md               Security updates
  - CHANGES_SUMMARY.md         Recent changes
```

## 🌐 Service URLs

```
Local Development:
  Application: http://localhost:8080
  Health:      http://localhost:8080/health
  API:         http://localhost:8080/api/v1

Production (NodePort):
  Application: http://173.212.218.251:30800
  Health:      http://173.212.218.251:30800/health
  
Infrastructure:
  <!-- Keycloak:    http://173.212.218.251:30080 -->
  Keycloak:    https://keycloak.skillsier.com
  PostgreSQL:  173.212.218.251:30432
  Kafka:       173.212.218.251:31691
```

## 🎯 Common Tasks

### Add New User Field
1. Update `internal/domain/user/entity.go`
2. Update `application/user/dto.go`
3. Update `application/user/mapper.go`
4. Restart app (auto-migration runs)

### Update Keycloak Secret
```bash
# Get new secret from Keycloak admin console
export KEYCLOAK_CLIENT_SECRET=new-secret
source scripts/get-secrets.sh
make k8s-restart
```

### Reset Database
```bash
make db-reset  # ⚠️ WARNING: Deletes all data!
```

### View Recent Logs
```bash
kubectl logs --tail=100 -n skillsier -l app=users-be
```

### Restart Service
```bash
# Local
Ctrl+C then make run

# Kubernetes
make k8s-restart
```

## ❗ Important Notes

- **Client Name:** Always use `users-be` (not `myapp`)
- **Service Accounts:** Must be enabled in Keycloak client
- **Security:** Updated dependencies fix CRITICAL vulnerabilities
- **Secrets:** Never commit to Git (in `.gitignore`)
- **Auto-Migration:** Runs on every application start
- **Outbox:** Processes events every 5 seconds

## 📞 Getting Help

1. Check logs: `kubectl logs -n skillsier -l app=users-be`
2. Review documentation in order:
   - QUICK_REFERENCE.md (this file)
   - README.md
   - GETTING_STARTED.md
   - KEYCLOAK_SETUP.md
3. Test connectivity to all services
4. Verify secrets exist and are valid

---

**Quick Setup:** Export Keycloak credentials → Run `source scripts/get-secrets.sh` → Run `make run` 🚀