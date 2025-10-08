# 🚀 Skillsier - Enterprise Freelancing Platform

[![React](https://img.shields.io/badge/React-19.0.0-blue.svg)](https://reactjs.org/)
[![Next.js](https://img.shields.io/badge/Next.js-15.1.3-black.svg)](https://nextjs.org/)
[![React Native](https://img.shields.io/badge/React%20Native-0.76.6-blue.svg)](https://reactnative.dev/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.6.3-blue.svg)](https://www.typescriptlang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

> A modern, high-performance freelancing platform built with React 19, supporting 9 languages and 120 FPS mobile experience.

## ✨ Features

- 🌍 **9 Languages** - EN, AR, ZH, HI, DE, FR, TR, ES, RU
- 📱 **Cross-Platform** - Web (Next.js) + Mobile (iOS & Android)
- ⚡ **120 FPS** - Butter-smooth mobile experience
- 🎨 **Modern UI** - Tailwind CSS + NativeWind
- 🔐 **Secure Auth** - Keycloak integration
- 📊 **Real-time** - TanStack Query + Zustand
- 🚀 **Monorepo** - 80%+ code reuse
- 💼 **Freelancing** - Complete marketplace features

## 🎯 Quick Start

### Prerequisites

- Node.js 20.x or higher
- pnpm 9.x or higher
- iOS: Xcode (macOS only)
- Android: Android Studio

### Installation

```bash
# Clone repository
git clone <your-repo-url>
cd skillsier-fe

# Install dependencies
pnpm install

# Setup environment
./setup-env.sh

# Start development
pnpm dev              # Both web & mobile
pnpm dev:web          # Web only (localhost:3000)
pnpm dev:mobile       # Mobile only (Expo)
```

## 📱 Mobile Development

### Expo Go (Quick)
```bash
pnpm dev:mobile
# Scan QR code with Expo Go app
```

### Development Build (Full Features)
```bash
cd apps/mobile
npx expo prebuild --clean
npx expo run:ios        # iOS
npx expo run:android    # Android
```

## 🌍 Supported Languages

| Language | Code | Status |
|----------|------|--------|
| English | `en` | ✅ |
| Arabic | `ar` | ✅ |
| Chinese | `zh` | ✅ |
| Hindi | `hi` | ✅ |
| German | `de` | ✅ |
| French | `fr` | ✅ |
| Turkish | `tr` | ✅ |
| Spanish | `es` | ✅ |
| Russian | `ru` | ✅ |

## 📁 Project Structure

```
skillsier-fe/
├── apps/
│   ├── web/              # Next.js 15 web app
│   └── mobile/           # React Native Expo app
├── packages/
│   ├── shared/           # Business logic
│   ├── ui/               # UI components
│   ├── types/            # TypeScript types
│   └── config/           # Shared configs
└── docs/                 # Documentation
```

## 🛠️ Available Scripts

```bash
# Development
pnpm dev              # Start all
pnpm dev:web          # Web only
pnpm dev:mobile       # Mobile only

# Build
pnpm build            # Build all
pnpm build:web        # Web production
pnpm build:mobile     # Mobile production

# Quality
pnpm lint             # Lint all
pnpm type-check       # Type check
pnpm format           # Format code

# Clean
pnpm clean            # Clean artifacts
```

## 🎨 Tech Stack

### Web
- **Framework**: Next.js 15 (App Router)
- **UI**: React 19, Tailwind CSS
- **i18n**: next-intl
- **State**: Zustand, TanStack Query

### Mobile
- **Framework**: Expo 52 (New Architecture)
- **UI**: React Native, NativeWind
- **i18n**: i18next
- **Performance**: 120 FPS, FlashList, MMKV

### Shared
- **Language**: TypeScript 5.6
- **Monorepo**: Turborepo, pnpm
- **Testing**: Jest, React Testing Library (ready)
- **Linting**: ESLint, Prettier

## 📚 Documentation

- [Setup Guide](docs/SETUP.md) - Complete installation guide
- [Architecture](docs/ARCHITECTURE.md) - System architecture
- [Performance](docs/PERFORMANCE.md) - 120 FPS optimization
- [i18n Guide](docs/I18N.md) - Internationalization
- [Freelancing](docs/FREELANCING.md) - Platform features
- [Contributing](docs/CONTRIBUTING.md) - Contribution guide

## 🏗️ Building for Production

### Web
```bash
pnpm build:web
# Deploy to Vercel/Netlify
```

### Mobile
```bash
cd apps/mobile
eas build --platform ios
eas build --platform android
eas submit --platform all
```

## 🧪 Testing

```bash
# Unit tests
pnpm test

# E2E tests
pnpm test:e2e:web
pnpm test:e2e:mobile

# Coverage
pnpm test:coverage
```

## 🤝 Contributing

We welcome contributions! Please read our [Contributing Guide](docs/CONTRIBUTING.md) first.

### Development Workflow

1. Fork the repository
2. Create feature branch: `git checkout -b feat/my-feature`
3. Commit changes: `git commit -m "feat: add feature"`
4. Push: `git push origin feat/my-feature`
5. Open Pull Request

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details

## 🙏 Acknowledgments

- [Next.js](https://nextjs.org/)
- [Expo](https://expo.dev/)
- [React](https://react.dev/)
- [Tailwind CSS](https://tailwindcss.com/)

## 📞 Support

- 📧 Email: support@skillsier.com
- 💬 Discord: [Join our community](https://discord.gg/skillsier)
- 📖 Docs: [docs.skillsier.com](https://docs.skillsier.com)

---

**Built with ❤️ by the Skillsier Team**