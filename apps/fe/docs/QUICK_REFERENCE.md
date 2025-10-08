# 🚀 Skillsier Quick Reference Card

## 📦 Version Summary

| Package | Version | Purpose |
|---------|---------|---------|
| React | **19.0.0** | UI library with new compiler |
| Next.js | 15.1.3 | Web framework |
| React Native | 0.76.6 | Mobile framework |
| Expo | 52.0.20 | Mobile tooling (New Arch) |

## ⚡ Performance Features

### Mobile (120 FPS Ready)
```typescript
✅ New Architecture (Fabric + TurboModules)
✅ 120Hz ProMotion support (iPhone 13 Pro+)
✅ FlashList (10x faster than FlatList)
✅ MMKV (zero-bridge storage)
✅ Reanimated 3 (UI thread animations)
✅ expo-image (optimized caching)
✅ Hermes engine
```

### Web
```typescript
✅ Server Components (Next.js)
✅ Code splitting
✅ Image optimization
✅ React 19 compiler (auto-memoization)
```

## 🌍 i18n Cheat Sheet

### Supported Languages
- 🇺🇸 **English** (en) - LTR
- 🇸🇦 **Arabic** (ar) - RTL

### Usage

**Web:**
```typescript
import { useTranslations } from 'next-intl';

const t = useTranslations('auth');
<button>{t('login')}</button>
```

**Mobile:**
```typescript
import { useTranslation } from 'react-i18next';

const { t } = useTranslation();
<Text>{t('auth.login')}</Text>
```

### Key Translations

| Key | English | Arabic |
|-----|---------|--------|
| `common.welcome` | Welcome | مرحباً |
| `auth.login` | Sign In | تسجيل الدخول |
| `auth.register` | Sign Up | إنشاء حساب |
| `auth.logout` | Sign Out | تسجيل الخروج |
| `dashboard.title` | Dashboard | لوحة التحكم |
| `courses.title` | Courses | الدورات |
| `profile.title` | Profile | الملف الشخصي |

## 🎨 RTL Support

### Tailwind Classes
```css
/* Margin */
ltr:ml-4 rtl:mr-4  /* margin-left in LTR, margin-right in RTL */

/* Padding */
ltr:pl-4 rtl:pr-4  /* padding-left in LTR, padding-right in RTL */

/* Text Alignment */
ltr:text-left rtl:text-right

/* Flex Direction */
ltr:flex-row rtl:flex-row-reverse

/* Logical Properties (Recommended) */
ms-4  /* margin-start (auto RTL) */
me-4  /* margin-end (auto RTL) */
ps-4  /* padding-start (auto RTL) */
pe-4  /* padding-end (auto RTL) */
```

## 🔧 Common Commands

### Development
```bash
pnpm dev              # Both apps
pnpm dev:web          # Web only
pnpm dev:mobile       # Mobile only
```

### Mobile Specific
```bash
npx expo prebuild --clean           # Rebuild native
npx expo run:ios                    # Run iOS
npx expo run:android                # Run Android
npx expo start --clear              # Clear cache
```

### Build
```bash
pnpm build:web                      # Web production
eas build --platform ios            # iOS build
eas build --platform android        # Android build
```

### Language Testing
```bash
# Web
http://localhost:3000/en            # English
http://localhost:3000/ar            # Arabic (RTL)

# Mobile
# Use in-app language switcher in Profile
```

## 🎯 Performance Optimization Checklist

### Mobile
- [ ] Use `OptimizedFlashList` instead of `FlatList`
- [ ] Store data in MMKV, not AsyncStorage
- [ ] Use `useHighFPSAnimation` for animations
- [ ] Enable `removeClippedSubviews` on lists
- [ ] Use `React.memo` for expensive components
- [ ] Use `expo-image` instead of `Image`
- [ ] Avoid inline functions in renderItem
- [ ] Provide `estimatedItemSize` to FlashList

### Web
- [ ] Use Server Components where possible
- [ ] Lazy load routes with dynamic imports
- [ ] Optimize images with next/image
- [ ] Use React 19's automatic memoization
- [ ] Code split by route
- [ ] Preload critical resources

## 📱 Frame Rate Detection

```typescript
import { getOptimalFrameRate, FRAME_RATE } from '@/lib/performance';

const frameRate = getOptimalFrameRate();
// Returns: 60, 90, or 120 based on device

if (frameRate === FRAME_RATE.ULTRA) {
  console.log('Running at 120 FPS! 🚀');
}
```

## 🌐 Language Switching

### Web
```typescript
// Automatic via URL
router.push('/ar/dashboard')  // Switch to Arabic
router.push('/en/dashboard')  // Switch to English

// Or use LanguageSwitcher component
<LanguageSwitcher />
```

### Mobile
```typescript
import { setLanguage } from '@/lib/i18n';

// Programmatic
await setLanguage('ar');  // App restarts for RTL

// Or use LanguageSwitcher component
<LanguageSwitcher />
```

## 🐛 Quick Fixes

### "Translations not loading"
```bash
# Web
rm -rf .next && pnpm dev:web

# Mobile
npx expo start --clear
```

### "120 FPS not working"
```bash
# Check device supports 120Hz
# iOS: Settings > Display > ProMotion
# Android: Settings > Display > Screen refresh rate

# Rebuild app
npx expo prebuild --clean
npx expo run:ios
```

### "RTL not switching"
```bash
# Mobile requires app restart
# This is expected behavior
# Or use Updates.reloadAsync()
```

### "React 19 type errors"
```bash
pnpm add -D @types/react@19.0.0 @types/react-dom@19.0.0
```

## 📊 Monitoring

### Check FPS (Mobile)
```typescript
import { PerformanceObserver } from 'react-native';

const observer = new PerformanceObserver((list) => {
  list.getEntries().forEach((entry) => {
    if (entry.duration > 8.33) {  // 120 FPS threshold
      console.warn('Frame drop:', entry.duration);
    }
  });
});
```

### Check Bundle Size (Web)
```bash
pnpm build:web
# Check .next/analyze for bundle size
```

## 🎓 Learning Resources

- **React 19**: https://react.dev/blog/2024/12/05/react-19
- **120 FPS Guide**: See Performance section in upgrade guide
- **i18n Best Practices**: See Translation section in upgrade guide
- **RTL Guidelines**: https://rtlstyling.com/posts/rtl-styling

## 💡 Pro Tips

1. **Use translation keys consistently**: `feature.component.text`
2. **Test on real 120Hz devices**: Simulators cap at 60 FPS
3. **Always provide estimatedItemSize**: Critical for FlashList
4. **Use logical properties**: `ms-4` instead of `ml-4/mr-4`
5. **Memoize expensive computations**: React 19 helps but not everything
6. **Profile before optimizing**: Use Chrome DevTools / Flipper
7. **Test RTL thoroughly**: Many bugs only show in RTL
8. **Keep translations in sync**: Use a tool like i18n-ally VSCode extension

## 🚀 Next Steps

1. ✅ Upgrade to React 19
2. ✅ Enable 120 FPS support
3. ✅ Implement i18n (EN + AR)
4. 🔲 Add French/Spanish
5. 🔲 Implement offline mode
6. 🔲 Add push notifications
7. 🔲 Optimize bundle size further
8. 🔲 Add performance monitoring

---

**You're all set! Build something amazing! 🎉**