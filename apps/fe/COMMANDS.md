# 📝 Skillsier Quick Command Reference

## 🚀 Development

### Start Servers
```bash
pnpm dev              # Both web and mobile
pnpm dev:web          # Web only (localhost:3000)
pnpm dev:mobile       # Mobile only (Expo dev server)
./dev.sh              # Interactive launcher
```

### Web Development
```bash
cd apps/web
pnpm dev              # Start Next.js dev server
pnpm build            # Production build
pnpm start            # Start production server
pnpm lint             # Lint code
```

### Mobile Development
```bash
cd apps/mobile

# Development
pnpm dev              # Start Expo dev server
npx expo start        # Alternative Expo start
npx expo start --clear # Clear cache

# Run on devices
npx expo run:ios      # iOS simulator
npx expo run:android  # Android emulator

# Prebuild (first time)
npx expo prebuild --clean
```

## 📦 Package Management

### Install Dependencies
```bash
pnpm install                              # Install all
pnpm add <package>                        # Add to root
pnpm add <package> --filter=web           # Add to web
pnpm add <package> --filter=mobile        # Add to mobile
pnpm add <package> --filter=@skillsier/shared  # Add to shared
```

### Update Dependencies
```bash
pnpm update                   # Update all
pnpm update --latest          # Update to latest versions
pnpm outdated                 # Check outdated packages
```

## 🏗️ Building

### Web Production
```bash
pnpm build:web                # Build for production
cd apps/web && pnpm start     # Test production build
```

### Mobile Production
```bash
cd apps/mobile

# Configure EAS (first time only)
eas login
eas build:configure

# Build
eas build --platform ios      # iOS build
eas build --platform android  # Android build
eas build --platform all      # Both platforms

# Submit to stores
eas submit --platform ios
eas submit --platform android
```

## 🧹 Cleaning

### Clear Caches
```bash
pnpm clean                    # Clean build artifacts
rm -rf node_modules           # Remove dependencies
pnpm install                  # Reinstall

# Mobile specific
cd apps/mobile
npx expo start --clear        # Clear Metro bundler cache
rm -rf node_modules .expo     # Deep clean
```

### Reset Mobile Project
```bash
cd apps/mobile
rm -rf node_modules ios android .expo
npx expo prebuild --clean
```

## 🧪 Testing

### Unit Tests
```bash
pnpm test                     # Run all tests
pnpm test:watch               # Watch mode
pnpm test:coverage            # Coverage report
```

### E2E Tests
```bash
pnpm test:e2e:web             # Web E2E
pnpm test:e2e:mobile          # Mobile E2E
```

## 🔍 Code Quality

### Linting
```bash
pnpm lint                     # Lint all packages
pnpm lint --filter=web        # Lint web only
pnpm lint --fix               # Auto-fix issues
```

### Type Checking
```bash
pnpm type-check               # Check all packages
pnpm type-check --filter=@skillsier/shared  # Check shared only
```

### Formatting
```bash
pnpm format                   # Format all code
pnpm format:check             # Check formatting
```

## 🌍 Language Management

### Add New Language
```bash
# 1. Create translation file
cp packages/shared/src/lib/i18n/translations/en.json \
   packages/shared/src/lib/i18n/translations/ja.json

# 2. Update config
# Edit: packages/shared/src/lib/i18n/config.ts
# Edit: apps/web/i18n.ts
# Edit: apps/mobile/src/lib/i18n/index.ts
```

### Test Languages
```bash
# Web
http://localhost:3000/en     # English
http://localhost:3000/ar     # Arabic (RTL)
http://localhost:3000/zh     # Chinese

# Mobile - Use in-app language switcher
```

## 🔧 Troubleshooting

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

### Fix Metro Bundler
```bash
cd apps/mobile
npx expo start --clear
# or
npx react-native start --reset-cache
```

### Fix TypeScript Errors
```bash
pnpm clean
rm -rf node_modules
rm -rf apps/*/node_modules
rm -rf packages/*/node_modules
pnpm install
pnpm type-check
```

### Fix Next.js Cache
```bash
cd apps/web
rm -rf .next
pnpm dev
```

## 📱 Mobile Specific

### iOS Development
```bash
# Open in Xcode
cd apps/mobile/ios
open Skillsier.xcworkspace

# Run specific simulator
npx expo run:ios --simulator="iPhone 15 Pro"

# View logs
npx react-native log-ios
```

### Android Development
```bash
# Open in Android Studio
cd apps/mobile/android
studio .

# Run specific device
npx expo run:android --device

# View logs
npx react-native log-android
```

### Development Build
```bash
# Install development client
npx expo install expo-dev-client

# Build
npx expo prebuild --clean
npx expo run:ios
npx expo run:android
```

## 🚀 Deployment

### Web Deployment (Vercel)
```bash
# Install Vercel CLI
npm i -g vercel

# Deploy
vercel --prod

# Or use GitHub integration
git push origin main  # Auto-deploys
```

### Mobile Deployment (EAS)
```bash
cd apps/mobile

# Production build
eas build --platform all --profile production

# Submit to stores
eas submit --platform ios
eas submit --platform android

# Create update (OTA)
eas update --branch production --message "Bug fixes"
```

## 🔐 Environment Variables

### Setup
```bash
./setup-env.sh              # Create env files
```

### Check
```bash
# Web
cat apps/web/.env.local

# Mobile
cat apps/mobile/.env
```

## 📊 Monitoring

### Bundle Size
```bash
# Web
cd apps/web
pnpm build
# Check .next/analyze

# Mobile
npx expo export
# Check dist folder
```

### Performance
```bash
# Lighthouse (web)
lighthouse http://localhost:3000

# React DevTools
# Open in browser/React Native debugger
```

## 🎨 Customization

### Change Branding
```bash
# Update logo
apps/web/public/images/logo.png
apps/mobile/assets/icon.png

# Update colors
packages/ui/src/theme/colors.ts
apps/web/tailwind.config.ts
apps/mobile/tailwind.config.js
```

### Update App Name
```bash
# Mobile
apps/mobile/app.json  # Update "name" and "slug"

# Web
apps/web/src/app/[locale]/layout.tsx  # Update metadata
```

## 📚 Useful Aliases

Add to your `~/.zshrc` or `~/.bashrc`:

```bash
# Skillsier shortcuts
alias sk="cd ~/skillsier-fe"
alias sk-dev="cd ~/skillsier-fe && ./dev.sh"
alias sk-web="cd ~/skillsier-fe && pnpm dev:web"
alias sk-mobile="cd ~/skillsier-fe && pnpm dev:mobile"
alias sk-ios="cd ~/skillsier-fe/apps/mobile && npx expo run:ios"
alias sk-android="cd ~/skillsier-fe/apps/mobile && npx expo run:android"
alias sk-clean="cd ~/skillsier-fe && pnpm clean && rm -rf node_modules && pnpm install"
```

## 🆘 Quick Fixes

### "Command not found: pnpm"
```bash
npm install -g pnpm@9.15.0
```

### "React 19 type errors"
```bash
pnpm add -D @types/react@19.0.0 @types/react-dom@19.0.0
```

### "Metro bundler won't start"
```bash
cd apps/mobile
npx expo start --clear
# or
pkill -f "react-native" && npx expo start
```

### "Can't resolve @skillsier/..."
```bash
pnpm install
cd packages/shared && pnpm build  # if needed
```

### "Expo prebuild fails"
```bash
cd apps/mobile
rm -rf ios android
npx expo prebuild --clean
```

---

**💡 Tip**: Use `./dev.sh` for an interactive menu to choose what to run!