# Keycloak with Kafka Integration (Kubernetes)

Complete setup for integrating Keycloak events with Kafka using the webhook plugin and a lightweight HTTP-to-Kafka bridge.

## Architecture

```
┌─────────────┐       ┌──────────────┐       ┌─────────────┐
│  Keycloak   │──────>│ Kafka Bridge │──────>│    Kafka    │
│  (Webhook)  │ HTTP  │  (HTTP→SASL) │ SASL  │   Cluster   │
└─────────────┘       └──────────────┘       └─────────────┘
       │                                             │
       │                                             │
       v                                             v
┌─────────────┐                            ┌─────────────┐
│ PostgreSQL  │                            │   Topics    │
└─────────────┘                            │ - keycloak- │
                                           │   events    │
                                           └─────────────┘
```

## Components

1. **PostgreSQL** - Keycloak database backend
2. **Keycloak** - Identity and access management with webhook plugin
3. **Kafka Bridge** - Lightweight Node.js service that converts HTTP webhooks to Kafka messages
4. **Kafka Cluster** - Event streaming platform (pre-installed via Strimzi)

## Prerequisites

- Kubernetes cluster (K3s) up and running
- kubectl configured and connected to your cluster
- Kafka cluster installed and running (using Strimzi operator)
- KafkaUser `keycloak-user` created with SCRAM-SHA-512 authentication
- At least 2GB RAM and 2 CPU cores available

## Files Overview

```
.
├── keycloak-deployment.yaml    # Main Keycloak deployment with Kafka Bridge
├── postgres-plain.yaml         # PostgreSQL StatefulSet
├── install-keycloak.sh         # Installation script
├── test-and-monitor.sh         # Testing and monitoring utility
└── README.md                   # This file
```

## Installation

### Step 1: Ensure Kafka is Running

First, verify your Kafka cluster is installed and running:

```bash
# Check Kafka cluster
kubectl -n kafka get kafka kafka-cluster

# Check Kafka user exists
kubectl -n kafka get kafkauser keycloak-user

# Get Kafka user password (save this)
kubectl -n kafka get secret keycloak-user -o jsonpath='{.data.password}' | base64 -d
```

### Step 2: (Optional) Update Domain

If you want to use a custom domain for Keycloak:

```bash
export DOMAIN="keycloak.yourdomain.com"
```

Or edit `keycloak-deployment.yaml` and replace `keycloak.yourdomain.com` with your domain.

### Step 3: Run Installation Script

Make the script executable and run it:

```bash
chmod +x install-keycloak.sh
./install-keycloak.sh
```

The script will:
- Create the `keycloak` namespace
- Deploy PostgreSQL database
- Copy Kafka credentials
- Deploy Kafka Bridge service
- Deploy Keycloak with webhook plugin
- Create the `keycloak-events` Kafka topic
- Display connection details

Installation takes approximately 3-5 minutes.

## Configuration

### Environment Variables

You can customize the installation using environment variables:

```bash
# Change namespace (default: keycloak)
export NAMESPACE="my-keycloak"

# Change Kafka namespace (default: kafka)
export KAFKA_NAMESPACE="my-kafka"

# Set custom domain
export DOMAIN="auth.mydomain.com"

# Set timeout (default: 600s)
export TIMEOUT="900s"

# Run installation
./install-keycloak.sh
```

### Keycloak Webhook Events

By default, the following events are captured:

- `LOGIN` - User login
- `LOGOUT` - User logout
- `REGISTER` - User registration
- `UPDATE_EMAIL` - Email update
- `UPDATE_PROFILE` - Profile update
- `UPDATE_PASSWORD` - Password change
- `VERIFY_EMAIL` - Email verification
- `REMOVE_TOTP` - 2FA removal
- `UPDATE_TOTP` - 2FA update
- `CODE_TO_TOKEN` - Token exchange
- `LOGIN_ERROR` - Failed login
- `REGISTER_ERROR` - Failed registration

To modify which events are captured, edit the `WEBHOOK_EVENTS_TAKEN` environment variable in `keycloak-deployment.yaml`:

```yaml
- name: WEBHOOK_EVENTS_TAKEN
  value: "LOGIN,LOGOUT,REGISTER"  # Customize this list
```

Remove this variable entirely to capture ALL Keycloak events.

## Testing and Monitoring

### Interactive Menu

Use the test and monitoring script:

```bash
chmod +x test-and-monitor.sh
./test-and-monitor.sh
```

This provides an interactive menu with options to:
1. Check system status
2. Test Kafka Bridge connection
3. Monitor events in real-time
4. Send test events
5. View recent messages
6. Check logs
7. Get connection details
8. Cleanup test resources

### Manual Testing

#### Access Keycloak Admin Console

Port forward to access locally:

```bash
kubectl -n keycloak port-forward svc/keycloak 8080:8080
```

Then open: http://localhost:8080
- Username: `admin`
- Password: `admin`

#### Send Test Event to Kafka Bridge

```bash
kubectl -n keycloak run test-event --rm -ti --restart=Never \
  --image=curlimages/curl:8.5.0 -- \
  curl -X POST http://kafka-bridge:3000/events \
  -H "Content-Type: application/json" \
  -H "Authorization: Basic $(echo -n 'webhook_user:webhook_secret_2024' | base64)" \
  -d '{"type":"TEST","message":"Hello Kafka"}'
```

#### Monitor Kafka Events

```bash
kubectl -n kafka run kafka-consumer --rm -ti --restart=Never \
  --image=quay.io/strimzi/kafka:0.47.0-kafka-3.9.1 -- \
  bin/kafka-console-consumer.sh \
  --bootstrap-server kafka-cluster-kafka-bootstrap:9092 \
  --topic keycloak-events \
  --from-beginning
```

#### Check Keycloak Logs

```bash
kubectl -n keycloak logs -f deployment/keycloak
```

#### Check Kafka Bridge Logs

```bash
kubectl -n keycloak logs -f deployment/kafka-bridge
```

## Resource Usage

### Expected Resource Consumption

| Component | Memory | CPU | Storage |
|-----------|--------|-----|---------|
| PostgreSQL | 256-512 MB | 200-500m | 10 GB |
| Keycloak | 512 MB - 1 GB | 250m-1000m | - |
| Kafka Bridge | 128-256 MB | 100-500m | - |
| **Total** | **~1-2 GB** | **~0.5-2 cores** | **10 GB** |

### Scaling

To scale Keycloak for higher load:

```bash
# Scale Keycloak pods
kubectl -n keycloak scale deployment/keycloak --replicas=2

# Scale Kafka Bridge
kubectl -n keycloak scale deployment/kafka-bridge --replicas=2
```

## Troubleshooting

### Keycloak Pod Not Starting

Check logs:
```bash
kubectl -n keycloak describe pod -l app=keycloak
kubectl -n keycloak logs -l app=keycloak
```

Common issues:
- PostgreSQL not ready: Wait for PostgreSQL StatefulSet
- Plugin download failed: Check internet connectivity
- Database connection failed: Verify PostgreSQL credentials

### Kafka Bridge Connection Issues

Check if Kafka user credentials are correct:
```bash
# Verify secret exists
kubectl -n keycloak get secret keycloak-user

# Test Kafka connection from bridge pod
kubectl -n keycloak exec -it deployment/kafka-bridge -- sh
```

### Events Not Appearing in Kafka

1. Check webhook configuration:
```bash
kubectl -n keycloak exec -it deployment/keycloak -- sh
env | grep WEBHOOK
```

2. Verify Kafka Bridge health:
```bash
kubectl -n keycloak run test-health --rm -ti --restart=Never \
  --image=curlimages/curl -- \
  curl http://kafka-bridge:3000/health
```

3. Check Kafka Bridge logs for errors:
```bash
kubectl -n keycloak logs -f deployment/kafka-bridge
```

4. Ensure Kafka topic exists:
```bash
kubectl -n kafka get kafkatopic keycloak-events
```

### Plugin Not Loading

Check if plugins were downloaded:
```bash
kubectl -n keycloak exec -it deployment/keycloak -- ls -lh /opt/keycloak/providers/
```

Should show:
- `keycloak-webhook-provider-core-0.9.1-all.jar`
- `keycloak-webhook-provider-http-0.9.1-all.jar`

## Security Considerations

### Production Recommendations

1. **Change default passwords**:
   - Update Keycloak admin password
   - Update PostgreSQL password in `postgres-plain.yaml`
   - Update Kafka Bridge auth credentials in `keycloak-deployment.yaml`

2. **Enable TLS**:
   - Configure cert-manager for automatic TLS certificates
   - Update ingress with proper TLS configuration

3. **Network Policies**:
   - The deployment includes basic NetworkPolicies
   - Review and adjust based on your security requirements

4. **Resource Limits**:
   - Adjust resource requests/limits based on your load
   - Monitor actual usage and tune accordingly

5. **Backup**:
   - Implement backup for PostgreSQL data
   - Consider using VolumeSnapshots for PVC backups

## Event Message Format

Events sent to Kafka have the following structure:

```json
{
  "type": "LOGIN",
  "realmId": "master",
  "userId": "user-123",
  "time": 1704067200000,
  "ipAddress": "192.168.1.100",
  "details": {
    "username": "john.doe",
    "auth_method": "openid-connect"
  },
  "bridgeMetadata": {
    "receivedAt": "2024-01-01T00:00:00.000Z",
    "source": "keycloak-webhook"
  }
}
```

## Integration with Other Services

### Consuming Events with Spring Boot

```java
@KafkaListener(topics = "keycloak-events", groupId = "my-app")
public void handleKeycloakEvent(String message) {
    // Process event
}
```

### Consuming Events with Node.js

```javascript
const { Kafka } = require('kafkajs');

const kafka = new Kafka({
  brokers: ['kafka-cluster-kafka-bootstrap.kafka.svc:9092']
});

const consumer = kafka.consumer({ groupId: 'my-app' });
await consumer.subscribe({ topic: 'keycloak-events' });

await consumer.run({
  eachMessage: async ({ message }) => {
    const event = JSON.parse(message.value.toString());
    console.log('Received event:', event);
  },
});
```

## Uninstallation

To remove all components:

```bash
# Delete Keycloak and Kafka Bridge
kubectl delete -f keycloak-deployment.yaml

# Delete PostgreSQL (WARNING: This deletes data!)
kubectl delete -f postgres-plain.yaml

# Delete the namespace (removes everything)
kubectl delete namespace keycloak

# Remove Kafka topic
kubectl -n kafka delete kafkatopic keycloak-events
```

## Performance Tuning

### For High-Volume Scenarios

1. **Increase Kafka Bridge replicas**:
```yaml
spec:
  replicas: 3
```

2. **Tune Kafka producer settings** in `keycloak-deployment.yaml`:
```javascript
const producer = kafka.producer({
  retry: { retries: 5 },
  batch: { size: 16384 },
  compression: 'snappy',
});
```

3. **Increase Kafka topic partitions**:
```bash
kubectl -n kafka patch kafkatopic keycloak-events --type merge \
  -p '{"spec":{"partitions":6}}'
```

## Support and Contributing

For issues or questions:
- Check the troubleshooting section above
- Review Keycloak logs and Kafka Bridge logs
- Ensure all prerequisites are met
- Verify network connectivity between components

## License

This setup uses:
- Keycloak (Apache 2.0)
- PostgreSQL (PostgreSQL License)
- Strimzi Kafka Operator (Apache 2.0)
- vymalo/keycloak-webhook (MIT)