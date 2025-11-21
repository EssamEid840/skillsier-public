# Setup Guide

Complete guide for setting up the Skillsier frontend development environment.

## Prerequisites

### Required Software

1. **Node.js 20.18.1+**
```bash
   # Install via nvm (recommended)
   curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
   nvm install 20.18.1
   nvm use 20.18.1
```

2. **pnpm 10.0.0+**
```bash
   npm install -g pnpm@10
```

3. **Git**
```bash
   # macOS
   brew install git
   
   # Ubuntu/Debian
   sudo apt-get install git
```

### Mobile Development (Optional)

#### Android Development

1. **Android Studio**
   - Download from https://developer.android.com/studio
   - Install Android SDK Platform 34
   - Configure Android SDK path: `~/Android/Sdk`

2. **Environment Variables**
```bash
   # Add to ~/.zshrc or ~/.bashrc
   export ANDROID_HOME=$HOME/Android/Sdk
   export PATH=$PATH:$ANDROID_HOME/emulator
   export PATH=$PATH:$ANDROID_HOME/platform-tools
```

#### iOS Development (macOS only)

1. **Xcode**
```bash
   xcode-select --install
```

2. **CocoaPods**
```bash
   sudo gem install cocoapods
```

## Installation

### 1. Clone Repository
```bash
git clone https://github.com/your-org/skillsier-fe.git
cd skillsier-fe
```

### 2. Install Dependencies
```bash
# Use Node 20.18.1
nvm use

# Install all dependencies
pnpm install
```

### 3. Environment Configuration
```bash
# Copy environment template
cp .env.example .env

# Edit .env with your values
vim .env
```

**Required Environment Variables:**
```env
# API Configuration
API_BASE_URL=http://localhost:8080

# Auth Provider (dev | keycloak)
AUTH_PROVIDER=dev

# Dev Auth Credentials (for AUTH_PROVIDER=dev)
DEV_ADMIN_EMAIL=admin@skillsier.dev
DEV_ADMIN_PASSWORD=admin123
DEV_CLIENT_EMAIL=client@skillsier.dev
DEV_CLIENT_PASSWORD=client123
DEV_FREELANCER_EMAIL=freelancer@skillsier.dev
DEV_FREELANCER_PASSWORD=freelancer123
```

### 4. Verify Installation
```bash
# Lint check
pnpm lint

# Type check
pnpm typecheck

# Run tests
pnpm test
```

## Running Applications

### Web Application
```bash
# Development mode
pnpm dev:web

# Production build
pnpm build:web
pnpm --filter web start
```

Access at: http://localhost:3000

### Mobile Application
```bash
# Start Metro bundler
pnpm dev:mobile

# In separate terminals:
# Android
pnpm --filter mobile android

# iOS
pnpm --filter mobile ios

# Web preview
pnpm --filter mobile web
```

## Development Tools

### VS Code Extensions (Recommended)

Install these extensions for best experience:
```json
{
  "recommendations": [
    "dbaeumer.vscode-eslint",
    "esbenp.prettier-vscode",
    "bradlc.vscode-tailwindcss",
    "ms-vscode.vscode-typescript-next",
    "expo.vscode-expo-tools"
  ]
}
```

### Git Hooks

Husky is configured for pre-commit and pre-push hooks:

- **pre-commit:** Runs lint-staged, typecheck
- **pre-push:** Runs full test suite, build

To skip hooks (not recommended):
```bash
git commit --no-verify
git push --no-verify
```

## Troubleshooting

### pnpm install fails
```bash
# Clear pnpm cache
pnpm store prune

# Remove node_modules and reinstall
rm -rf node_modules
pnpm install --frozen-lockfile
```

### Metro bundler issues (Mobile)
```bash
# Clear Metro cache
pnpm --filter mobile start --clear

# Reset Expo cache
rm -rf apps/mobile/.expo
```

### Build errors
```bash
# Clean all builds
pnpm clean

# Rebuild everything
pnpm install
pnpm build
```

### Type errors
```bash
# Check specific package
pnpm --filter @skillsier/types typecheck

# Rebuild types
pnpm --filter @skillsier/types build
```

## Database Setup (Backend)

The frontend expects the following backend services:
```bash
# Backend services should be running on:
API Gateway: http://localhost:8080
users-be: http://localhost:8081
jobs-be: http://localhost:8082
proposals-be: http://localhost:8083
contracts-be: http://localhost:8084
```

See backend repository for setup instructions.

## CI/CD Setup

### GitHub Actions

Workflows are in `.github/workflows/`:
- `ci-web.yml` - Web app CI
- `ci-mobile.yml` - Mobile app CI
- `cd-web.yml` - Web deployment

### Required Secrets

Add these to GitHub repository secrets:
```
VERCEL_TOKEN
VERCEL_ORG_ID
VERCEL_PROJECT_ID
API_BASE_URL
KEYCLOAK_ISSUER
KEYCLOAK_CLIENT_ID
KEYCLOAK_CLIENT_SECRET
NEXTAUTH_SECRET
NEXTAUTH_URL
```

## Next Steps

- Read [ARCHITECTURE.md](./ARCHITECTURE.md) for system overview
- Read [KEYCLOAK.md](./KEYCLOAK.md) for production auth
- Read [I18N.md](./I18N.md) for adding languages
- Start developing! 🚀