#!/bin/bash

# ============================================
# SKILLSIER COMPLETE SETUP SCRIPT
# ============================================

```
set -e

echo "🚀 Starting Skillsier Monorepo Setup..."
echo ""

# Check Node.js version
echo "📋 Checking prerequisites..."
NODE_VERSION=$(node -v | cut -d'v' -f2 | cut -d'.' -f1)
if [ "$NODE_VERSION" -lt 20 ]; then
    echo "❌ Node.js 20.x or higher is required (found $(node -v))"
    exit 1
fi
echo "✅ Node.js version: $(node -v)"

# Install pnpm if not present
if ! command -v pnpm &> /dev/null; then
    echo "📦 Installing pnpm..."
    npm install -g pnpm@9.15.0
fi
echo "✅ pnpm version: $(pnpm -v)"

# Create project structure
echo ""
echo "📁 Creating project structure..."

mkdir -p skillsier-fe
cd skillsier-fe

# Create root files
cat > package.json << 'EOF'
{
  "name": "skillsier-fe",
  "version": "1.0.0",
  "private": true,
  "scripts": {
    "dev": "turbo run dev",
    "dev:web": "turbo run dev --filter=web",
    "dev:mobile": "turbo run dev --filter=mobile",
    "build": "turbo run build",
    "build:web": "turbo run build --filter=web",
    "build:mobile": "turbo run build --filter=mobile",
    "lint": "turbo run lint",
    "type-check": "turbo run type-check",
    "clean": "turbo run clean && rm -rf node_modules",
    "format": "prettier --write \"**/*.{ts,tsx,md,json}\""
  },
  "devDependencies": {
    "@turbo/gen": "^2.3.3",
    "eslint": "^8.57.0",
    "prettier": "^3.2.5",
    "turbo": "^2.3.3",
    "typescript": "^5.6.3"
  },
  "packageManager": "pnpm@9.15.0",
  "engines": {
    "node": ">=20.0.0",
    "pnpm": ">=9.0.0"
  }
}
EOF

cat > pnpm-workspace.yaml << 'EOF'
packages:
  - 'apps/*'
  - 'packages/*'
EOF

cat > turbo.json << 'EOF'
{
  "$schema": "https://turbo.build/schema.json",
  "globalDependencies": ["**/.env.*local"],
  "pipeline": {
    "build": {
      "dependsOn": ["^build"],
      "outputs": [".next/**", "!.next/cache/**", "dist/**", ".expo/**"]
    },
    "dev": {
      "cache": false,
      "persistent": true
    },
    "lint": {
      "dependsOn": ["^lint"]
    },
    "type-check": {
      "dependsOn": ["^type-check"]
    }
  }
}
EOF

cat > .gitignore << 'EOF'
node_modules/
.next/
.expo/
dist/
.env*.local
.env
.DS_Store
.turbo/
*.tsbuildinfo
EOF

cat > .prettierrc << 'EOF'
{
  "semi": true,
  "trailingComma": "es5",
  "singleQuote": true,
  "printWidth": 100,
  "tabWidth": 2
}
EOF

cat > tsconfig.json << 'EOF'
{
  "compilerOptions": {
    "target": "ES2020",
    "lib": ["ES2020"],
    "module": "ESNext",
    "moduleResolution": "bundler",
    "strict": true,
    "skipLibCheck": true,
    "esModuleInterop": true
  }
}
EOF

echo "✅ Root configuration created"

# Install root dependencies
echo ""
echo "📦 Installing root dependencies..."
pnpm install

echo ""
echo "✅ Skillsier monorepo setup complete!"
echo ""
echo "📋 Next steps:"
echo ""
echo "1. Set up environment variables:"
echo "   - Create apps/web/.env.local"
echo "   - Create apps/mobile/.env"
echo ""
echo "2. Start development:"
echo "   pnpm dev:web      # Web app on http://localhost:3000"
echo "   pnpm dev:mobile   # Mobile app with Expo"
echo ""
echo "3. Build for production:"
echo "   pnpm build:web"
echo "   pnpm build:mobile"
echo ""
echo "📚 Documentation: ./README.md"
echo ""
```
# ============================================
# ENVIRONMENT SETUP HELPER
# ============================================
```
cat > setup-env.sh << 'EOF'
#!/bin/bash

echo "🔧 Setting up environment files..."

# Web environment
cat > apps/web/.env.local << 'WEBENV'
NEXT_PUBLIC_API_URL=http://localhost:3000/api
NEXT_PUBLIC_KEYCLOAK_URL=http://localhost:8080
NEXT_PUBLIC_KEYCLOAK_REALM=skillsier
NEXT_PUBLIC_KEYCLOAK_CLIENT_ID=skillsier-web
WEBENV

# Mobile environment
cat > apps/mobile/.env << 'MOBILEENV'
EXPO_PUBLIC_API_URL=http://localhost:3000/api
EXPO_PUBLIC_KEYCLOAK_URL=http://localhost:8080
EXPO_PUBLIC_KEYCLOAK_REALM=skillsier
EXPO_PUBLIC_KEYCLOAK_CLIENT_ID=skillsier-mobile
MOBILEENV

echo "✅ Environment files created"
echo ""
echo "Web environment: apps/web/.env.local"
echo "Mobile environment: apps/mobile/.env"
echo ""
echo "⚠️  Update these files with your actual values before running the apps"
EOF

chmod +x setup-env.sh

# ============================================
# DEVELOPMENT HELPERS
# ============================================

cat > dev.sh << 'EOF'
#!/bin/bash

echo "🚀 Starting Skillsier Development Servers..."
echo ""
echo "Choose an option:"
echo "1. Web only (Next.js)"
echo "2. Mobile only (Expo)"
echo "3. Both (recommended)"
echo ""
read -p "Enter choice [1-3]: " choice

case $choice in
  1)
    echo "Starting web server..."
    pnpm dev:web
    ;;
  2)
    echo "Starting mobile server..."
    pnpm dev:mobile
    ;;
  3)
    echo "Starting both servers..."
    pnpm dev
    ;;
  *)
    echo "Invalid choice"
    exit 1
    ;;
esac
EOF

chmod +x dev.sh
```
# ============================================
# MOBILE DEV CLIENT SETUP
# ============================================
```
cat > setup-mobile-dev-client.sh << 'EOF'
#!/bin/bash

echo "📱 Setting up React Native Development Client..."
echo ""

cd apps/mobile

echo "1. Prebuild native projects..."
npx expo prebuild --clean

echo ""
echo "2. Install iOS dependencies (Mac only)..."
if [[ "$OSTYPE" == "darwin"* ]]; then
    cd ios
    pod install
    cd ..
    echo "✅ iOS dependencies installed"
else
    echo "⏭️  Skipped (not on macOS)"
fi

echo ""
echo "✅ Mobile development client setup complete!"
echo ""
echo "Run the app:"
echo "  iOS:     npx expo run:ios"
echo "  Android: npx expo run:android"
EOF

chmod +x setup-mobile-dev-client.sh
```
# ============================================
# QUICK COMMANDS REFERENCE
# ============================================
```
cat > COMMANDS.md << 'EOF'
# Skillsier Quick Command Reference
```
## Development

### Start Development Servers
```bash
pnpm dev              # Both web and mobile
pnpm dev:web          # Web only (localhost:3000)
pnpm dev:mobile       # Mobile only (Expo dev server)
./dev.sh              # Interactive launcher
```

### Mobile Development
```bash
cd apps/mobile

# First time setup
npx expo prebuild --clean
./setup-mobile-dev-client.sh

# Run on devices
npx expo run:ios            # iOS simulator
npx expo run:android        # Android emulator
npx expo start              # QR code for Expo Go
```

## Building

### Web Production Build
```bash
pnpm build:web
cd apps/web && pnpm start   # Test production build
```

### Mobile Production Build
```bash
cd apps/mobile

# Configure EAS (first time)
eas login
eas build:configure

# Build
eas build --platform ios
eas build --platform android
eas build --platform all

# Submit to stores
eas submit --platform ios
eas submit --platform android
```

## Package Management

### Install Dependencies
```bash
pnpm install                # Install all
pnpm add <package>          # Add to root
pnpm add <package> --filter=web        # Add to web
pnpm add <package> --filter=mobile     # Add to mobile
pnpm add <package> --filter=@skillsier/shared  # Add to shared
```

### Update Dependencies
```bash
pnpm update                 # Update all
pnpm update --latest        # Update to latest versions
```

## Code Quality

### Linting
```bash
pnpm lint                   # Lint all packages
pnpm lint --filter=web      # Lint web only
```

### Type Checking
```bash
pnpm type-check             # Check all packages
pnpm type-check --filter=@skillsier/shared
```

### Formatting
```bash
pnpm format                 # Format all code
```

## Troubleshooting

### Clear All Caches
```bash
pnpm clean                  # Clean build artifacts
rm -rf node_modules         # Remove dependencies
pnpm install                # Reinstall

# Mobile specific
cd apps/mobile
npx expo start --clear      # Clear Metro bundler cache
```

### Reset Mobile Project
```bash
cd apps/mobile
rm -rf node_modules ios android .expo
npx expo prebuild --clean
```

### Fix iOS Issues
```bash
cd apps/mobile/ios
pod deintegrate
pod install
cd ..
npx expo run:ios
```

### Fix Android Issues
```bash
cd apps/mobile/android
./gradlew clean
cd ..
npx expo run:android
```

## Environment Management

### Setup Environment
```bash
./setup-env.sh              # Create .env files
```

### Check Environment
```bash
# Web
cat apps/web/.env.local

# Mobile
cat apps/mobile/.env
```

## Package Scripts

### Shared Package (@skillsier/shared)
- Hooks: `packages/shared/src/features/*/hooks`
- Stores: `packages/shared/src/features/*/stores`
- API: `packages/shared/src/features/*/api`

### UI Package (@skillsier/ui)
- Components: `packages/ui/src/components`
- Theme: `packages/ui/src/theme`

### Types Package (@skillsier/types)
- Entities: `packages/types/src/entities`
- API Types: `packages/types/src/api`

## Useful Aliases

Add to your `~/.zshrc` or `~/.bashrc`:

```bash
alias sk-dev="cd ~/skillsier-fe && ./dev.sh"
alias sk-web="cd ~/skillsier-fe && pnpm dev:web"
alias sk-mobile="cd ~/skillsier-fe && pnpm dev:mobile"
alias sk-ios="cd ~/skillsier-fe/apps/mobile && npx expo run:ios"
alias sk-android="cd ~/skillsier-fe/apps/mobile && npx expo run:android"
alias sk-lint="cd ~/skillsier-fe && pnpm lint"
alias sk-clean="cd ~/skillsier-fe && pnpm clean"

EOF

echo ""
echo "✅ Helper scripts created:"
echo "   - setup-env.sh (create environment files)"
echo "   - dev.sh (interactive dev server launcher)"
echo "   - setup-mobile-dev-client.sh (setup React Native dev client)"
echo "   - COMMANDS.md (quick reference)"
echo ""
```