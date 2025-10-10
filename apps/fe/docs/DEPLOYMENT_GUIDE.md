# 🚀 Skillsier Frontend Deployment Guide

## Overview

This guide covers deploying the Skillsier frontend in different environments:

1. Local Development
2. Kubernetes Production
3. Mobile App Stores

---

## 📋 Pre-Deployment Checklist

### Keycloak Configuration

- [ ] Production Keycloak running at `https://keycloak.skillsier.com`
- [ ] Realm `skillsier` created
- [ ] Client `skillsier-fe` configured with correct redirect URIs
- [ ] Client `skillsier-bff` configured with service account roles
- [ ] Client `skillsier-mobile` configured for mobile apps
- [ ] Google Identity Provider configured with correct redirect URI
- [ ] All client secrets obtained and documented securely

### Backend API

- [ ] Backend API deployed and accessible
- [ ] API health endpoint responding
- [ ] CORS configured to allow frontend domains
- [ ] JWT validation configured with Keycloak

### DNS & Certificates

- [ ] DNS records pointing to cluster
- [ ] SSL certificates configured (cert-manager)
- [ ] Ingress controller installed (nginx)

---

## 🌐 Local Development Deployment

### 1. Setup Environment

```bash
# Run setup script
./setup-env.sh

# Choose option 1 (Local development)
# This creates:
# - apps/web/.env.local
# - apps/mobile/.env
```

### 2. Update Keycloak Secrets

Edit `apps/web/.env.local`:

```bash
# Get from Keycloak Admin Console
KEYCLOAK_CLIENT_SECRET=actual-secret-from-keycloak
KEYCLOAK_MGMT_CLIENT_SECRET=actual-mgmt-secret-from-keycloak
```

### 3. Configure Keycloak Redirect URIs

In Keycloak Admin → Clients → skillsier-fe → Settings:

```
Valid Redirect URIs:
  http://localhost:3000/api/auth/keycloak/oauth/callback/google
  http://localhost:3000/api/auth/keycloak/oauth/callback/local

Web Origins:
  http://localhost:3000

Valid Post Logout Redirect URIs:
  http://localhost:3000/*
```

### 4. Start Development Servers

```bash
# Interactive launcher
./dev.sh

# Or manually
pnpm dev:web      # Web on :3000
pnpm dev:mobile   # Mobile with Expo
```

### 5. Test Authentication

1. Navigate to http://localhost:3000/login
2. Test Google SSO
3. Test email/password registration
4. Verify dashboard access

---

## ☸️ Kubernetes Production Deployment

### 1. Prepare Docker Image

```bash
cd apps/web

# Build image
docker build -t skillsier/web-fe:1.0.0 -f Dockerfile .

# Tag for registry
docker tag skillsier/web-fe:1.0.0 your-registry/skillsier/web-fe:1.0.0

# Push to registry
docker push your-registry/skillsier/web-fe:1.0.0
```

### 2. Create Kubernetes Secrets

```bash
# Create secret with Keycloak credentials
kubectl create secret generic web-fe-secrets \
  --from-literal=KEYCLOAK_CLIENT_SECRET='your-actual-client-secret' \
  --from-literal=KEYCLOAK_MGMT_CLIENT_SECRET='your-actual-mgmt-secret' \
  -n default \
  --dry-run=client -o yaml | kubectl apply -f -
```

### 3. Update ConfigMap

Edit `k8s/web-deployment.yaml`:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: web-fe-config
data:
  NEXT_PUBLIC_API_URL: "http://users-be-service:8080/api"
  NEXT_PUBLIC_KEYCLOAK_URL: "https://keycloak.skillsier.com"
  NEXT_PUBLIC_APP_URL: "https://skillsier.com"
  # ... other vars
```

### 4. Configure Production Keycloak Redirect URIs

In Keycloak Admin → Clients → skillsier-fe → Settings:

```
Valid Redirect URIs:
  https://skillsier.com/api/auth/keycloak/oauth/callback/google
  https://skillsier.com/api/auth/keycloak/oauth/callback/local
  https://www.skillsier.com/api/auth/keycloak/oauth/callback/*

Web Origins:
  https://skillsier.com
  https://www.skillsier.com

Valid Post Logout Redirect URIs:
  https://skillsier.com/*
  https://www.skillsier.com/*
```

### 5. Deploy to Kubernetes

```bash
# Apply deployment
kubectl apply -f k8s/web-deployment.yaml

# Check deployment status
kubectl get pods -n default
kubectl get deployment web-fe -n default

# Check logs
kubectl logs -f deployment/web-fe -n default

# Check service
kubectl get svc web-fe-service -n default

# Check ingress
kubectl get ingress web-fe-ingress -n default
```

### 6. Verify Deployment

```bash
# Check pod health
kubectl get pods -l app=web-fe -n default

# Check logs for errors
kubectl logs -f deployment/web-fe -n default

# Test health endpoint
curl https://skillsier.com/api/health

# Should return:
# {"status":"healthy","timestamp":"..."}
```

### 7. Test Production Authentication

1. Navigate to https://skillsier.com/login
2. Test Google SSO flow
3. Test registration
4. Verify dashboard access
5. Check browser console for errors

---

## 📱 Mobile App Deployment

### Option A: Expo EAS Build (Recommended)

#### 1. Install EAS CLI

```bash
npm install -g eas-cli
```

#### 2. Login to Expo

```bash
cd apps/mobile
eas login
```

#### 3. Configure EAS Build

```bash
# Initialize EAS configuration
eas build:configure
```

This creates `eas.json`:

```json
{
  "build": {
    "development": {
      "developmentClient": true,
      "distribution": "internal"
    },
    "preview": {
      "distribution": "internal"
    },
    "production": {
      "autoIncrement": true
    }
  },
  "submit": {
    "production": {}
  }
}
```

#### 4. Update Environment for Production

Edit `apps/mobile/.env`:

```bash
EXPO_PUBLIC_API_URL=https://api.skillsier.com/api
EXPO_PUBLIC_KEYCLOAK_URL=https://keycloak.skillsier.com
EXPO_PUBLIC_KEYCLOAK_REALM=skillsier
EXPO_PUBLIC_KEYCLOAK_CLIENT_ID=skillsier-mobile
```

#### 5. Build for Production

```bash
# Build for both platforms
eas build --platform all --profile production

# Or build separately
eas build --platform ios --profile production
eas build --platform android --profile production
```

#### 6. Submit to App Stores

```bash
# iOS App Store
eas submit --platform ios --latest

# Google Play Store
eas submit --platform android --latest
```

### Option B: Local Build

#### iOS (Mac only)

```bash
cd apps/mobile

# Generate native projects
npx expo prebuild

# Build with Xcode
cd ios
pod install
cd ..

# Open in Xcode
open ios/Skillsier.xcworkspace

# Build and archive in Xcode
# Upload to App Store Connect
```

#### Android

```bash
cd apps/mobile

# Generate native projects
npx expo prebuild

# Build release APK
cd android
./gradlew assembleRelease

# Or build AAB for Play Store
./gradlew bundleRelease

# Output:
# android/app/build/outputs/apk/release/app-release.apk
# android/app/build/outputs/bundle/release/app-release.aab
```

---

## 🔄 Update/Rollback Strategy

### Web Application

#### Zero-Downtime Updates

```bash
# Build new version
docker build -t skillsier/web-fe:1.1.0 apps/web
docker push skillsier/web-fe:1.1.0

# Update deployment
kubectl set image deployment/web-fe \
  web-fe=skillsier/web-fe:1.1.0 \
  -n default

# Watch rollout
kubectl rollout status deployment/web-fe -n default
```

#### Rollback

```bash
# Rollback to previous version
kubectl rollout undo deployment/web-fe -n default

# Or rollback to specific revision
kubectl rollout history deployment/web-fe -n default
kubectl rollout undo deployment/web-fe --to-revision=2 -n default
```

### Mobile Application

#### Over-The-Air (OTA) Updates

```bash
cd apps/mobile

# Create update
eas update --branch production --message "Bug fixes"

# Users get update automatically without app store
```

#### Full App Store Update

```bash
# Increment version in app.json
# Build new version
eas build --platform all --profile production

# Submit to stores
eas submit --platform all --latest
```

---

## 📊 Monitoring & Health Checks

### Web Application Health

```bash
# Check health endpoint
curl https://skillsier.com/api/health

# Check Kubernetes health
kubectl get pods -l app=web-fe
kubectl describe pod <pod-name>
kubectl logs -f deployment/web-fe
```

### Application Metrics

```bash
# Pod resource usage
kubectl top pods -l app=web-fe

# View pod events
kubectl get events --field-selector involvedObject.name=web-fe

# Check ingress
kubectl describe ingress web-fe-ingress
```

---

## 🐛 Troubleshooting Production Issues

### Issue: Pods in CrashLoopBackOff

```bash
# Check logs
kubectl logs -f deployment/web-fe --tail=100

# Check events
kubectl describe pod <pod-name>

# Common causes:
# 1. Missing environment variables
# 2. Cannot connect to backend API
# 3. Memory limit too low
```

### Issue: 502 Bad Gateway

```bash
# Check if pods are running
kubectl get pods -l app=web-fe

# Check service endpoints
kubectl get endpoints web-fe-service

# Check ingress
kubectl describe ingress web-fe-ingress

# Check pod logs
kubectl logs -f deployment/web-fe
```

### Issue: Authentication Not Working

1. Check Keycloak redirect URIs match exactly
2. Verify client secrets in Kubernetes secret
3. Check browser console for errors
4. Verify cookies are being set
5. Check CORS configuration

```bash
# Get current secrets
kubectl get secret web-fe-secrets -o yaml

# Update secrets
kubectl create secret generic web-fe-secrets \
  --from-literal=KEYCLOAK_CLIENT_SECRET='new-secret' \
  --dry-run=client -o yaml | kubectl apply -f -

# Restart pods to pick up new secrets
kubectl rollout restart deployment/web-fe
```

---

## 🔒 Security Considerations

### Production Checklist

- [ ] All secrets stored in Kubernetes secrets (not ConfigMaps)
- [ ] HTTPS enforced (no HTTP)
- [ ] Security headers configured
- [ ] CORS properly configured
- [ ] Rate limiting enabled
- [ ] DDoS protection enabled
- [ ] Regular security updates
- [ ] Backup strategy in place
- [ ] Monitoring and alerting configured

### Rotate Secrets

```bash
# Generate new client secret in Keycloak
# Update Kubernetes secret
kubectl create secret generic web-fe-secrets \
  --from-literal=KEYCLOAK_CLIENT_SECRET='new-secret' \
  --from-literal=KEYCLOAK_MGMT_CLIENT_SECRET='new-mgmt-secret' \
  --dry-run=client -o yaml | kubectl apply -f -

# Restart deployment
kubectl rollout restart deployment/web-fe
```

---

## ✅ Post-Deployment Verification

### 1. Functional Testing

- [ ] Homepage loads
- [ ] Login with Google works
- [ ] Login with email/password works
- [ ] Registration works
- [ ] Dashboard accessible after login
- [ ] Logout works
- [ ] Profile pages load
- [ ] API calls succeed

### 2. Performance Testing

- [ ] Page load time < 3 seconds
- [ ] API response time < 500ms
- [ ] No console errors
- [ ] Mobile app smooth (60 FPS)

### 3. Security Testing

- [ ] HTTPS redirect works
- [ ] Security headers present
- [ ] No sensitive data in logs
- [ ] Authentication required for protected routes
- [ ] CSRF protection working

---

## 📞 Support & Escalation

If deployment issues occur:

1. Check pod logs: `kubectl logs -f deployment/web-fe`
2. Check Keycloak configuration
3. Verify all secrets are correct
4. Test health endpoint
5. Review this guide's troubleshooting section

---

**Deployment Guide Version**: 1.0.0  
**Last Updated**: October 9, 2025