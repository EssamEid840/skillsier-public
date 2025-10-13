# Keycloak Client Setup for users-be

## Overview

The `users-be` service needs a Keycloak client for authentication. This guide shows you how to:
1. Create the client in Keycloak
2. Get the client secret
3. Automatically create the Kubernetes secret using environment variables

## Step 1: Create Keycloak Client

### 1.1 Access Keycloak Admin Console

```bash
# Open in browser
# http://173.212.218.251:30080
# OR
https://keycloak.skillsier.com
```

Login with admin credentials.

### 1.2 Select Realm

- Click on the realm dropdown (top-left)
- Select **skillsier** realm

### 1.3 Create New Client

1. In the left sidebar, click **Clients**
2. Click **"Create client"** button

**General Settings:**
- Client type: `OpenID Connect`
- Client ID: `users-be`
- Name: `Users Backend Service`
- Description: `Backend microservice for user management`
- Click **"Next"**

**Capability Config:**
- Client authentication: **ON** ✅ (This makes it a confidential client)
- Authorization: **OFF**
- Authentication flow:
  - Standard flow: **OFF** (we don't need browser login)
  - Direct access grants: **ON** ✅ (for password grant if needed)
  - Service accounts roles: **ON** ✅ (IMPORTANT: for client credentials)
  - Implicit flow: **OFF**
- Click **"Next"**

**Login Settings:**
- Root URL: `http://users-be.skillsier.svc.cluster.local:8080`
- Valid redirect URIs: `http://users-be.skillsier.svc.cluster.local:8080/*`
- Click **"Save"**

### 1.4 Get Client Secret

1. After saving, you'll be on the client details page
2. Click the **"Credentials"** tab at the top
3. You'll see **"Client secret"** with a value like: `7X9kLmP4qR2wN8vT5yU3hJ6aZ1bC0dE`
4. Click the **copy icon** to copy it
5. **IMPORTANT:** Save this somewhere secure - you'll need it in the next step

### 1.5 Configure Service Account (Optional but Recommended)

1. Click the **"Service accounts roles"** tab
2. Click **"Assign role"**
3. Search for and assign necessary roles (e.g., `view-users`, `manage-users`)

## Step 2: Setup Local Environment

### 2.1 Export Environment Variables

```bash
# Set the client credentials you got from Keycloak
export KEYCLOAK_CLIENT_ID=users-be
export KEYCLOAK_CLIENT_SECRET=7X9kLmP4qR2wN8vT5yU3hJ6aZ1bC0dE  # Replace with your actual secret

# Optional: Set Keycloak URL if different
# export KEYCLOAK_URL=http://173.212.218.251:30080
export KEYCLOAK_URL=https://keycloak.skillsier.com
export KEYCLOAK_REALM=skillsier
```

### 2.2 Run the Smart Secret Script

The script will automatically:
1. Check if the secret exists in Kubernetes
2. If it exists, verify the credentials work
3. If it doesn't exist or credentials are invalid, create/update the secret using your environment variables

```bash
# Run the script (it will source environment variables)
source scripts/get-secrets.sh
```

**Expected Output:**

```
Fetching secrets from Kubernetes...
Checking Keycloak client credentials...
⚠️  Keycloak secret does not exist, creating...

⚠️  Keycloak client secret not found or credentials are invalid.

To create the secret, you need to provide:
  1. KEYCLOAK_CLIENT_ID (should be 'users-be')
  2. KEYCLOAK_CLIENT_SECRET (get from Keycloak admin console)

Found environment variables:
  KEYCLOAK_CLIENT_ID: users-be
  KEYCLOAK_CLIENT_SECRET: 7X9kLmP4qR***

Verifying credentials with Keycloak...
✓ Credentials verified successfully!

Creating Kubernetes secret...
secret/keycloak-client-users-be created
✓ Secret created successfully!

Fetching other secrets...

✓ All secrets loaded successfully!
========================================
  DB Password:         Redis***
  Kafka Password:      WNUy3***
  Keycloak Client ID:  users-be
  Keycloak Secret:     7X9kLmP4qR***
========================================

Environment variables exported. You can now run:
  make run
  OR
  go run cmd/api/main.go
```

### 2.3 Verify Secret in Kubernetes

```bash
# Check the secret was created
kubectl get secret keycloak-client-users-be -n keycloak

# View the secret (base64 encoded)
kubectl get secret keycloak-client-users-be -n keycloak -o yaml

# Decode and view (optional)
kubectl get secret keycloak-client-users-be -n keycloak -o jsonpath='{.data.client-id}' | base64 -d
echo ""
kubectl get secret keycloak-client-users-be -n keycloak -o jsonpath='{.data.client-secret}' | base64 -d
echo ""
```

## Step 3: Test the Setup

### 3.1 Run the Application

```bash
# The environment variables are already exported from the script
go run cmd/api/main.go
```

### 3.2 Verify Connection

Check the logs for successful Keycloak connection:

```
<!-- Connecting to Keycloak issuer: http://173.212.218.251:30080/realms/skillsier -->
Connecting to Keycloak issuer: https://keycloak.skillsier.com/realms/skillsier

✓ Keycloak authentication configured successfully
  Client ID: users-be
  Realm: skillsier
```
### 3.3 Test with Keycloak API (Optional)

```bash
# Get an access token using client credentials
# KEYCLOAK_URL=http://173.212.218.251:30080
KEYCLOAK_URL=https://keycloak.skillsier.com
REALM=skillsier

TOKEN_RESPONSE=$(curl -s -X POST \
  "${KEYCLOAK_URL}/realms/${REALM}/protocol/openid-connect/token" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "grant_type=client_credentials" \
  -d "client_id=${KEYCLOAK_CLIENT_ID}" \
  -d "client_secret=${KEYCLOAK_CLIENT_SECRET}")

# Extract access token
ACCESS_TOKEN=$(echo $TOKEN_RESPONSE | grep -o '"access_token":"[^"]*' | cut -d'"' -f4)

echo "Access Token: ${ACCESS_TOKEN:0:50}..."

# If you get a token, the client is configured correctly!
```

## Step 4: Deploy to Kubernetes

### 4.1 Verify Secret Exists

```bash
# The secret should already be created from Step 2
kubectl get secret keycloak-client-users-be -n keycloak
```

### 4.2 Deploy the Application

```bash
# Deploy users-be to Kubernetes
make k8s-deploy

# Check pods are using the secret
kubectl get pods -n skillsier -l app=users-be
kubectl logs -n skillsier -l app=users-be | grep -i keycloak
```

## Troubleshooting

### Issue: "ERROR: Environment variables not set!"

**Solution:**
```bash
# Make sure you export the variables before running the script
export KEYCLOAK_CLIENT_ID=users-be
export KEYCLOAK_CLIENT_SECRET=your-actual-secret

# Then run the script again
source scripts/get-secrets.sh
```

### Issue: "ERROR: Credentials verification failed!"

**Possible causes:**
1. Wrong client secret
2. Keycloak not accessible
3. Client 'users-be' doesn't exist in realm 'skillsier'
4. Service accounts not enabled

**Solution:**
```bash
# Test Keycloak connectivity
# curl -I http://173.212.218.251:30080/realms/skillsier
curl -I https://keycloak.skillsier.com/realms/skillsier

# Verify client exists
# Go to Keycloak admin console → Clients → search for "users-be"

# Check Service accounts roles is enabled
# Keycloak admin → Clients → users-be → Settings → Service accounts roles: ON

# Get a new client secret
# Keycloak admin → Clients → users-be → Credentials → Regenerate secret
```

### Issue: Secret exists but credentials are invalid

The script will automatically detect this and recreate the secret:

```bash
# Just export new credentials and run the script
export KEYCLOAK_CLIENT_ID=users-be
export KEYCLOAK_CLIENT_SECRET=new-secret-from-keycloak

source scripts/get-secrets.sh
# Script will detect invalid credentials and recreate the secret
```

### Issue: Want to manually create/update the secret

```bash
# Delete existing secret (if any)
kubectl delete secret keycloak-client-users-be -n keycloak

# Create new secret
kubectl create secret generic keycloak-client-users-be -n keycloak \
  --from-literal=client-id=users-be \
  --from-literal=client-secret=your-actual-secret

# Verify
kubectl get secret keycloak-client-users-be -n keycloak -o yaml
```

## Script Behavior Summary

The `get-secrets.sh` script follows this logic:

```
1. Check if secret exists in K8s
   ├─ YES: Retrieve credentials
   │   └─ Verify credentials work with Keycloak
   │       ├─ Valid: Export and use
   │       └─ Invalid: Recreate from env variables
   └─ NO: Create from env variables
       └─ Verify credentials before creating
           ├─ Valid: Create secret and export
           └─ Invalid: Show error and exit

2. Export all credentials to environment
3. Ready to run application
```

## Security Best Practices

1. **Never commit secrets to Git** - They're in `.gitignore`
2. **Rotate secrets regularly** - Update in Keycloak and re-run script
3. **Use different secrets for dev/staging/prod**
4. **In production, use sealed secrets or external secret managers** (e.g., HashiCorp Vault, AWS Secrets Manager)
5. **Limit client permissions** - Only assign necessary roles

## Quick Reference Commands

```bash
# Setup: Export credentials and create secret
export KEYCLOAK_CLIENT_ID=users-be
export KEYCLOAK_CLIENT_SECRET=your-secret
source scripts/get-secrets.sh

# Verify: Check secret exists
kubectl get secret keycloak-client-users-be -n keycloak

# Update: Regenerate and update secret
# 1. Regenerate in Keycloak admin console
# 2. Export new secret
# 3. Run script again
export KEYCLOAK_CLIENT_SECRET=new-secret
source scripts/get-secrets.sh

# Delete: Remove secret
kubectl delete secret keycloak-client-users-be -n keycloak

# Run: Start the application
make run
```

## Next Steps

After setup:
1. ✅ Test user registration in Keycloak
2. ✅ Verify users-be receives events
3. ✅ Check user is created in database
4. ✅ Deploy to production: `make k8s-deploy`