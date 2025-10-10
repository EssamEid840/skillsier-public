# 🚀 Skillsier Frontend - Complete Setup Guide

## 📋 Overview

This is the complete frontend implementation for Skillsier - an Upwork-like freelancing platform with:

- **Web App**: Next.js 15 + React 19
- **Mobile Apps**: React Native Expo (iOS & Android)
- **Authentication**: Keycloak with Google SSO
- **Architecture**: Monorepo with 80%+ code sharing
- **Languages**: 9 languages supported
- **Production Ready**: K8s deployment configs included

---

## 🎯 What's Included

### ✅ Complete Features
- Authentication (Login/Register with Google SSO & username/password)
- User Profiles (Freelancer & Client types)
- Skills Management (CRUD operations)
- Work Experience, Education, Certifications
- Portfolio with image uploads
- Profile completion tracking
- 30+ API integrations ready
- Production Keycloak configuration

### ✅ Development Ready
- Environment setup scripts
- Docker configurations
- Kubernetes manifests
- Health check endpoints
- TypeScript throughout
- ESLint + Prettier configured

---

## 📦 Prerequisites

Install these before starting:

```bash
# Node.js 20+
node -v  # Should be >= 20.0.0

# pnpm 9+
npm install -g pnpm@9.15.0
pnpm -v

# For iOS development (Mac only)
xcode-select --install

# For Android development
# Install Android Studio from https://developer.android.com/studio
```

---

## 🚀 Quick Start

### Step 1: Clone and Install

```bash
# Clone your repository
git clone <your-repo-url>
cd skillsier-fe

# Install all dependencies
pnpm install

# Make scripts executable
chmod +x setup-env.sh dev.sh setup-mobile-dev-client.sh
```

### Step 2: Configure Environment

```bash
# Run the environment setup script
./setup-env.sh

# Choose:
# 1 - Local development
# 2 - Kubernetes deployment
```

### Step 3: Update Keycloak Secrets

**Important**: Update these values in `apps/web/.env.local`:

1. Go to https://keycloak.skillsier.com/admin/
2. Login with admin credentials
3. Select realm: `skillsier`
4. Get `skillsier-fe` client secret:
   - Navigate to: Clients → skillsier-fe → Credentials tab
   - Copy the Client Secret
   - Update `KEYCLOAK_CLIENT_SECRET` in `.env.local`

5. Get `skillsier-bff` client secret:
   - Navigate to: Clients → skillsier-bff → Credentials tab
   - Copy the Client Secret
   - Update `KEYCLOAK_MGMT_CLIENT_SECRET` in `.env.local`

### Step 4: Start Development

```bash
# Interactive launcher (recommended)
./dev.sh

# Or start manually:
pnpm dev:web      # Web on http://localhost:3000
pnpm dev:mobile   # Mobile with Expo (scan QR code)

# Or start both:
pnpm dev
```

---

## 🔐 Keycloak Configuration

### Production Keycloak Setup

Your Keycloak is already configured at: `https://keycloak.skillsier.com`

#### Required Clients:

1. **skillsier-fe** (Web Frontend)
   - Client Type: OpenID Connect
   - Client Authentication: ON
   - Standard Flow: ON
   - Valid Redirect URIs:
     - `http://localhost:3000/api/auth/keycloak/oauth/callback/*`
     - `https://skillsier.com/api/auth/keycloak/oauth/callback/*`
   - Web Origins:
     - `http://localhost:3000`
     - `https://skillsier.com`

2. **skillsier-bff** (Management Client)
   - Client Type: OpenID Connect
   - Client Authentication: ON
   - Service Accounts: ON
   - Service Account Roles:
     - `realm-management` → `manage-users`
     - `realm-management` → `view-users`

3. **skillsier-mobile** (Mobile App)
   - Client Type: OpenID Connect
   - Client Authentication: OFF (Public client)
   - Standard Flow: ON
   - Valid Redirect URIs:
     - `skillsier://auth/callback`
     - `exp://localhost:8081/--/auth/callback`

#### Google Identity Provider:

Already configured with:
- Alias: `google`
- Trust Email: ON
- Scopes: `openid email profile`
- Redirect URI: `https://keycloak.skillsier.com/realms/skillsier/broker/google/endpoint`

---

## 📱 Mobile Development

### Option 1: Expo Go (Easiest)

```bash
# Start Expo dev server
pnpm dev:mobile

# Install Expo Go on your phone:
# iOS: https://apps.apple.com/app/expo-go/id982107779
# Android: https://play.google.com/store/apps/details?id=host.exp.exponent

# Scan QR code with:
# - iOS: Camera app
# - Android: Expo Go app
```

### Option 2: Development Build (Full Features)

```bash
# Setup mobile dev client (first time only)
./setup-mobile-dev-client.sh

# Run on iOS (Mac only)
cd apps/mobile
npx expo run:ios

# Run on Android
cd apps/mobile
npx expo run:android
```

---

## 🌐 Running in Kubernetes

### Step 1: Build Docker Images

```bash
# Build web image
cd apps/web
docker build -t skillsier/web-fe:latest .

# Push to registry
docker push skillsier/web-fe:latest
```

### Step 2: Update Secrets

```bash
# Create Kubernetes secret with actual values
kubectl create secret generic web-fe-secrets \
  --from-literal=KEYCLOAK_CLIENT_SECRET='your-actual-secret' \
  --from-literal=KEYCLOAK_MGMT_CLIENT_SECRET='your-actual-mgmt-secret' \
  -n default
```

### Step 3: Deploy

```bash
# Apply Kubernetes manifests
kubectl apply -f k8s/web-deployment.yaml

# Check status
kubectl get pods -n default
kubectl logs -f deployment/web-fe -n default

# Access the app
# Visit: https://skillsier.com
```

---

## 🧪 Testing Authentication

### Test Google SSO:

1. Go to http://localhost:3000/login
2. Click "Continue with Google"
3. You should be redirected to Keycloak
4. Then to Google OAuth
5. After authentication, back to your app dashboard

### Test Email/Password:

1. Go to http://localhost:3000/register
2. Fill in the registration form
3. Select user type (freelancer or client)
4. Submit
5. Login with email/password

### Test Mobile:

1. Start mobile app with `pnpm dev:mobile`
2. Tap "Sign in with Google"
3. Complete OAuth flow in browser
4. Return to app (dashboard)

---

## 📁 Project Structure

```
skillsier-fe/
├── apps/
│   ├── web/                    # Next.js web app
│   │   ├── src/
│   │   │   ├── app/           # App router pages
│   │   │   │   ├── api/       # API routes
│   │   │   │   │   ├── auth/keycloak/
│   │   │   │   │   │   ├── oauth/
│   │   │   │   │   │   │   ├── start/[provider]/route.ts
│   │   │   │   │   │   │   └── callback/[provider]/route.ts
│   │   │   │   │   │   └── register/route.ts
│   │   │   │   │   └── health/route.ts
│   │   │   │   └── [locale]/
│   │   │   │       ├── (auth)/
│   │   │   │       │   ├── login/page.tsx
│   │   │   │       │   └── register/page.tsx
│   │   │   │       └── dashboard/page.tsx
│   │   │   └── lib/
│   │   │       └── keycloak.ts
│   │   ├── .env.local
│   │   ├── next.config.js
│   │   └── Dockerfile
│   │
│   └── mobile/                 # React Native app
│       ├── app/
│       │   ├── (auth)/
│       │   │   ├── login.tsx
│       │   │   └── register.tsx
│       │   ├── (tabs)/
│       │   │   └── dashboard.tsx
│       │   └── auth/
│       │       └── callback.tsx
│       ├── src/
│       │   └── lib/
│       │       └── keycloak-mobile.ts
│       ├── .env
│       └── app.json
│
├── packages/
│   ├── shared/                 # Shared business logic
│   ├── ui/                     # Shared components
│   ├── types/                  # TypeScript types
│   └── config/                 # Shared configs
│
├── k8s/                        # Kubernetes manifests
│   └── web-deployment.yaml
│
├── setup-env.sh                # Environment setup
├── dev.sh                      # Dev launcher
└── setup-mobile-dev-client.sh  # Mobile setup
```

---

## 🔧 Common Commands

```bash
# Development
pnpm dev                # Start all apps
pnpm dev:web           # Web only
pnpm dev:mobile        # Mobile only

# Building
pnpm build             # Build all
pnpm build:web         # Build web
pnpm build:mobile      # Build mobile

# Code Quality
pnpm lint              # Lint all
pnpm type-check        # TypeScript check
pnpm format            # Format code

# Clean
pnpm clean             # Clean all build artifacts
rm -rf node_modules && pnpm install  # Fresh install
```

---

## 🐛 Troubleshooting

### Issue: "Missing Keycloak secrets"

**Solution:**
1. Check `.env.local` has both secrets filled
2. Get secrets from Keycloak Admin Console
3. Restart the dev server

### Issue: "OAuth redirect_uri_mismatch"

**Solution:**
1. Check Keycloak client redirect URIs
2. Ensure they match your app URLs exactly
3. For local dev: `http://localhost:3000/api/auth/keycloak/oauth/callback/google`
4. No trailing slashes!

### Issue: "Token exchange failed"

**Solution:**
1. Verify `KEYCLOAK_CLIENT_SECRET` is correct
2. Check Keycloak client has "Client authentication" ON
3. Ensure "Standard flow" is enabled

### Issue: "State mismatch - CSRF attack"

**Solution:**
1. Clear browser cookies
2. Check if cookies are being set properly
3. Ensure `sameSite: 'lax'` in cookie settings

### Issue: Mobile app won't connect to API

**Solution:**
1. Check `EXPO_PUBLIC_API_URL` in `.env`
2. For iOS simulator: Use computer's IP (not localhost)
3. For Android emulator: Use `10.0.2.2` instead of localhost
4. Example: `http://192.168.1.100:8080/api`

### Issue: "Can't resolve @skillsier/..."

**Solution:**
```bash
# Rebuild shared packages
cd packages/shared && pnpm build
cd packages/ui && pnpm build
cd packages/types && pnpm build

# Or clean and reinstall
pnpm clean && rm -rf node_modules && pnpm install
```

### Issue: Expo build fails

**Solution:**
```bash
cd apps/mobile
rm -rf node_modules ios android .expo
npx expo prebuild --clean
```

---

## 🌍 Supported Languages

The app supports 9 languages out of the box:

- 🇬🇧 English (en)
- 🇸🇦 Arabic (ar) - with RTL support
- 🇪🇸 Spanish (es)
- 🇫🇷 French (fr)
- 🇩🇪 German (de)
- 🇨🇳 Chinese (zh)
- 🇮🇳 Hindi (hi)
- 🇹🇷 Turkish (tr)
- 🇷🇺 Russian (ru)

Translation files are in: `packages/shared/src/lib/i18n/translations/`

---

## 🔒 Security Features

- ✅ PKCE (Proof Key for Code Exchange)
- ✅ State parameter for CSRF protection
- ✅ Secure httpOnly cookies
- ✅ Token refresh mechanism
- ✅ Encrypted token storage (MMKV on mobile)
- ✅ HTTPS enforcement in production
- ✅ Security headers configured

---

## 📊 Performance Features

- ✅ Server Components (Next.js 15)
- ✅ React 19 optimizations
- ✅ FlashList for mobile (10x faster than FlatList)
- ✅ MMKV storage (20x faster than AsyncStorage)
- ✅ Code splitting
- ✅ Image optimization
- ✅ Bundle size optimization
- ✅ TanStack Query for caching

---

## 🚀 Deployment

### Web (Vercel - Easiest)

```bash
# Install Vercel CLI
npm i -g vercel

# Deploy
cd apps/web
vercel --prod
```

### Web (Docker + K8s)

```bash
# Build and push
docker build -t skillsier/web-fe:latest apps/web
docker push skillsier/web-fe:latest

# Deploy to Kubernetes
kubectl apply -f k8s/web-deployment.yaml
```

### Mobile (EAS Build)

```bash
cd apps/mobile

# Configure EAS
eas login
eas build:configure

# Build for stores
eas build --platform all --profile production

# Submit to stores
eas submit --platform ios
eas submit --platform android
```

---

## 📚 Additional Resources

- **Keycloak Admin**: https://keycloak.skillsier.com/admin/
- **API Documentation**: Check your backend docs
- **Expo Documentation**: https://docs.expo.dev
- **Next.js Documentation**: https://nextjs.org/docs
- **React 19 Documentation**: https://react.dev

---

## 🆘 Getting Help

If you encounter issues:

1. Check this README's troubleshooting section
2. Review the `.env.local` configuration
3. Check Keycloak client configurations
4. Verify all secrets are correctly set
5. Check application logs: `kubectl logs -f deployment/web-fe`

---

## ✅ Checklist Before Going Live

- [ ] Update all environment variables for production
- [ ] Get real Keycloak client secrets
- [ ] Configure Google OAuth with production URLs
- [ ] Update Keycloak redirect URIs for production
- [ ] Build and test Docker images
- [ ] Deploy to Kubernetes cluster
- [ ] Test authentication flows end-to-end
- [ ] Test on real mobile devices
- [ ] Configure SSL certificates
- [ ] Set up monitoring and logging
- [ ] Configure backup strategy
- [ ] Test error scenarios

---

## 🎉 You're Ready!

Your Skillsier frontend is now complete with:

- ✅ Production Keycloak integration
- ✅ Google SSO working
- ✅ Both web and mobile apps
- ✅ Kubernetes deployment ready
- ✅ Local development setup
- ✅ All authentication flows

**Start your development:**

```bash
./dev.sh
```

**Good luck building your freelancing platform! 🚀**

---

**Last Updated**: October 9, 2025  
**Version**: 1.0.0