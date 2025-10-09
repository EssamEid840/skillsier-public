# 🚀 Skillsier Frontend - Quick Reference Card

## ⚡ Quick Commands

### Development
```bash
pnpm dev              # Start both web & mobile
pnpm dev:web          # Start web only (port 3000)
pnpm dev:mobile       # Start mobile only (Expo)
```

### Building
```bash
pnpm build            # Build all apps
pnpm build:web        # Build web for production
```

### Testing
```bash
pnpm lint             # Lint all code
pnpm type-check       # TypeScript check
pnpm format           # Format code
```

### Cleaning
```bash
pnpm clean            # Clean all build artifacts
```

---

## 📁 File Locations

### Web Components
```
apps/web/src/components/
├── layout/
│   └── MobileNav.tsx           ← Mobile drawer menu
└── dashboard/
    ├── DashboardShell.tsx      ← Dashboard container
    └── StatsCard.tsx           ← Statistics card
```

### Mobile Components
```
apps/mobile/src/components/
├── landing/
│   ├── HeroMobile.tsx          ← Hero section
│   └── FeaturesMobile.tsx      ← Features section
└── navigation/
    ├── TabBar.tsx              ← Bottom tabs
    └── DrawerContent.tsx       ← Drawer menu
```

### Shared Hooks
```
packages/shared/src/features/user/hooks/
├── index.ts                    ← Central exports
├── useClientProfile.ts
├── useFreelancerSkills.ts
├── useAddSkill.ts
├── useWorkExperience.ts
├── useEducation.ts
├── useCertifications.ts
├── usePortfolio.ts
└── ... (27 total hooks)
```

---

## 🎣 Hook Usage Examples

### Fetch Data
```typescript
import { useFreelancerProfile, useFreelancerSkills } from '@skillsier/shared';

const { data: profile, isLoading } = useFreelancerProfile();
const { data: skills, isLoading: skillsLoading } = useFreelancerSkills();
```

### Create/Update
```typescript
import { useAddSkill, useUpdateSkill } from '@skillsier/shared';

const { mutate: addSkill, isPending } = useAddSkill();
const { mutate: updateSkill } = useUpdateSkill();

// Add skill
addSkill({
  name: 'React',
  category: 'Frontend',
  level: 'EXPERT',
  yearsOfExperience: 5
});

// Update skill
updateSkill({
  skillId: '123',
  data: { level: 'EXPERT' }
});
```

### Delete
```typescript
import { useDeleteSkill } from '@skillsier/shared';

const { mutate: deleteSkill } = useDeleteSkill();

deleteSkill('skill-id-123');
```

---

## 🌍 Translation Keys

### Common
```typescript
t('common.search')           // Search
t('common.save')             // Save
t('common.cancel')           // Cancel
t('common.loading')          // Loading...
```

### Authentication
```typescript
t('auth.login')              // Sign In
t('auth.register')           // Sign Up
t('auth.logout')             // Sign Out
t('auth.email')              // Email
t('auth.password')           // Password
```

### Profile
```typescript
t('profile.title')           // Profile
t('profile.editProfile')     // Edit Profile
t('profile.firstName')       // First Name
t('profile.skills')          // Skills
t('profile.experience')      // Work Experience
```

### Jobs
```typescript
t('jobs.findWork')           // Find Work
t('jobs.title')              // Jobs
t('jobs.budget')             // Budget
t('jobs.deadline')           // Deadline
```

---

## 🎨 Tailwind Classes

### Colors
```css
bg-primary-600          /* Indigo background */
text-primary-700        /* Indigo text */
border-primary-500      /* Indigo border */

bg-success-600          /* Green */
bg-warning-600          /* Yellow */
bg-error-600            /* Red */
```

### Spacing
```css
p-4                     /* Padding 16px */
m-6                     /* Margin 24px */
gap-3                   /* Gap 12px */
space-y-4               /* Vertical spacing 16px */
```

### Responsive
```css
sm:text-lg              /* Small screens */
md:grid-cols-2          /* Medium screens */
lg:block                /* Large screens */
```

### RTL Support
```css
ltr:ml-4 rtl:mr-4       /* Margin left/right */
ltr:pl-3 rtl:pr-3       /* Padding left/right */
```

---

## 🔧 Environment Variables

### Web (.env.local)
```bash
NEXT_PUBLIC_API_URL=http://localhost:8080/api
NEXT_PUBLIC_KEYCLOAK_URL=http://localhost:8080
NEXT_PUBLIC_KEYCLOAK_REALM=skillsier
NEXT_PUBLIC_KEYCLOAK_CLIENT_ID=skillsier-web
```

### Mobile (.env)
```bash
EXPO_PUBLIC_API_URL=http://localhost:8080/api
EXPO_PUBLIC_KEYCLOAK_URL=http://localhost:8080
EXPO_PUBLIC_KEYCLOAK_REALM=skillsier
EXPO_PUBLIC_KEYCLOAK_CLIENT_ID=skillsier-mobile
```

---

## 📱 Mobile Commands

### Expo
```bash
npx expo start           # Start dev server
npx expo start --clear   # Clear cache and start
npx expo prebuild        # Generate native code
npx expo run:ios         # Run on iOS
npx expo run:android     # Run on Android
```

### EAS (Production)
```bash
eas build --platform ios       # Build iOS
eas build --platform android   # Build Android
eas submit --platform all      # Submit to stores
```

---

## 🐛 Common Issues & Solutions

### Issue: Module not found
```bash
pnpm install
pnpm build:packages
```

### Issue: Type errors
```bash
pnpm type-check
# Fix reported errors
```

### Issue: Metro cache (Mobile)
```bash
cd apps/mobile
npx expo start --clear
```

### Issue: Next.js cache (Web)
```bash
cd apps/web
rm -rf .next
pnpm dev:web
```

### Issue: Hooks not working
Check: `packages/shared/src/features/user/hooks/index.ts` exists and exports all hooks

---

## 📊 API Endpoints

### Authentication
```
POST /auth/register
POST /auth/login
POST /auth/logout
POST /auth/refresh
```

### User Profile
```
GET    /users/profile
PATCH  /users/profile
POST   /users/profile/avatar
DELETE /users/profile/avatar
```

### Skills
```
GET    /users/profile/skills
POST   /users/profile/skills
PATCH  /users/profile/skills/:id
DELETE /users/profile/skills/:id
```

### Work Experience
```
GET    /users/profile/experience
POST   /users/profile/experience
PATCH  /users/profile/experience/:id
DELETE /users/profile/experience/:id
```

### Portfolio
```
GET    /users/profile/portfolio
POST   /users/profile/portfolio
PATCH  /users/profile/portfolio/:id
DELETE /users/profile/portfolio/:id
POST   /users/profile/portfolio/:id/images
```

---

## 🎯 Key Files to Remember

### Entry Points
- `apps/web/src/app/[locale]/page.tsx` - Web landing
- `apps/mobile/app/index.tsx` - Mobile landing
- `apps/mobile/index.js` - Mobile entry

### Layouts
- `apps/web/src/app/[locale]/layout.tsx` - Root layout
- `apps/web/src/app/[locale]/(dashboard)/layout.tsx` - Dashboard
- `apps/mobile/app/(tabs)/_layout.tsx` - Tabs layout

### Configuration
- `packages/shared/package.json` - Shared dependencies
- `packages/shared/src/lib/i18n/config.ts` - i18n config
- `packages/shared/src/constants/api.ts` - API endpoints

---

## 💻 VS Code Shortcuts

### Navigation
```
Ctrl/Cmd + P          # Quick file open
Ctrl/Cmd + Shift + F  # Search in files
Ctrl/Cmd + Click      # Go to definition
```

### Editing
```
Alt + Up/Down         # Move line up/down
Ctrl/Cmd + D          # Select next occurrence
Ctrl/Cmd + /          # Toggle comment
```

### Terminal
```
Ctrl/Cmd + `          # Toggle terminal
Ctrl/Cmd + Shift + ` # New terminal
```

---

## 📚 Important Links

- **Next.js**: https://nextjs.org/docs
- **Expo**: https://docs.expo.dev
- **React Query**: https://tanstack.com/query
- **Tailwind**: https://tailwindcss.com/docs
- **TypeScript**: https://www.typescriptlang.org/docs

---

## 🎨 Component Patterns

### Web Component
```typescript
'use client'; // For client components

import { useTranslations } from 'next-intl';
import { Button } from '@skillsier/ui';

export function MyComponent() {
  const t = useTranslations('namespace');
  
  return (
    <div className="p-6">
      <h1>{t('title')}</h1>
      <Button onClick={handleClick}>
        {t('action')}
      </Button>
    </div>
  );
}
```

### Mobile Component
```typescript
import { View, Text } from 'react-native';
import { useTranslation } from 'react-i18next';
import { Button } from '@skillsier/ui';

export function MyComponent() {
  const { t } = useTranslation();
  
  return (
    <View className="p-6">
      <Text className="text-xl font-bold">
        {t('title')}
      </Text>
      <Button onPress={handlePress}>
        {t('action')}
      </Button>
    </View>
  );
}
```

---

## 🔄 State Management

### Using Zustand Store
```typescript
import { useAuth } from '@skillsier/shared';

const { user, isAuthenticated, login, logout } = useAuth();
```

### Using TanStack Query
```typescript
import { useQuery, useMutation } from '@tanstack/react-query';

// Query
const { data, isLoading, error } = useQuery({
  queryKey: ['key'],
  queryFn: fetchFn,
});

// Mutation
const { mutate, isPending } = useMutation({
  mutationFn: updateFn,
  onSuccess: () => {
    // Handle success
  },
});
```

---

## 🎯 Testing Checklist

### Before Committing
- [ ] `pnpm lint` passes
- [ ] `pnpm type-check` passes
- [ ] No console errors
- [ ] Features work on web
- [ ] Features work on mobile
- [ ] Responsive design works
- [ ] Translations present

### Before Deploying
- [ ] Production build succeeds
- [ ] Environment variables set
- [ ] API endpoints correct
- [ ] Images optimized
- [ ] Performance tested
- [ ] Security reviewed

---

## 📱 Platform Differences

### File Extensions
- `.tsx` - Web & Mobile
- `.web.tsx` - Web only
- `.native.tsx` - Mobile only

### Styling
- **Web**: className with Tailwind
- **Mobile**: className with NativeWind

### Navigation
- **Web**: `import { Link } from 'next/link'`
- **Mobile**: `import { Link } from 'expo-router'`

### Storage
- **Web**: `localStorage`
- **Mobile**: `MMKV` (via shared package)

---

## 🔍 Debugging

### Web Console
```javascript
console.log('Debug:', data);
console.error('Error:', error);
console.table(arrayData);
```

### Mobile Console
```javascript
import { Alert } from 'react-native';

Alert.alert('Title', 'Message');
console.log('Debug:', data);
```

### React Query DevTools
```typescript
// Web only - auto-included in dev
// Check browser for React Query panel
```

---

## 📦 Package Management

### Add Dependency
```bash
# Root
pnpm add <package>

# Specific workspace
pnpm add <package> --filter=web
pnpm add <package> --filter=mobile
pnpm add <package> --filter=@skillsier/shared
```

### Update Dependencies
```bash
pnpm update              # Update all
pnpm update <package>    # Update specific
```

### Remove Dependency
```bash
pnpm remove <package> --filter=workspace
```

---

## 🚀 Deployment

### Web (Vercel)
```bash
# Install Vercel CLI
npm i -g vercel

# Deploy
cd apps/web
vercel --prod
```

### Mobile (EAS)
```bash
# Install EAS CLI
npm i -g eas-cli

# Configure
cd apps/mobile
eas build:configure

# Build
eas build --platform all

# Submit
eas submit --platform all
```

---

## 🎨 Brand Customization

### Colors
Edit `packages/config/src/tailwind/base.js`:
```javascript
primary: {
  50: '#eef2ff',
  600: '#6366f1',
  // ...
}
```

### Logo
- Web: `apps/web/public/images/logo.svg`
- Mobile: `apps/mobile/assets/images/logo.png`

### App Name
- Web: `apps/web/src/app/[locale]/layout.tsx`
- Mobile: `apps/mobile/app.json`

---

## 📊 Performance Tips

### Web
- Use `'use client'` only when needed
- Implement lazy loading
- Optimize images (use Next Image)
- Enable caching headers

### Mobile
- Use FlashList for long lists
- Optimize image sizes
- Minimize re-renders
- Use memoization

### Both
- Implement pagination
- Add loading skeletons
- Cache API responses
- Optimize bundle size

---

## 🔒 Security Best Practices

### Never Commit
- Environment variables
- API keys
- Secrets
- Passwords
- Private keys

### Always Do
- Validate inputs
- Sanitize data
- Use HTTPS in production
- Implement rate limiting
- Log security events

---

## 📈 Monitoring

### Error Tracking
Add Sentry or similar:
```typescript
// apps/web/src/app/error.tsx
// apps/mobile/app/_layout.tsx
```

### Analytics
Add Google Analytics or similar:
```typescript
// Track page views
// Track user actions
// Monitor performance
```

---

## 🎓 Learning Path

### Beginner
1. Understand monorepo structure
2. Learn component patterns
3. Study hook usage
4. Practice with forms

### Intermediate
1. Master state management
2. Implement new features
3. Optimize performance
4. Handle edge cases

### Advanced
1. Architecture decisions
2. Performance profiling
3. Security hardening
4. Testing strategies

---

## 💡 Pro Tips

### 1. Use IntelliSense
TypeScript provides autocomplete for everything.

### 2. Check DevTools
React Query and React DevTools are your friends.

### 3. Read Error Messages
Error messages usually tell you exactly what's wrong.

### 4. Use Git Branches
Create feature branches for new work.

### 5. Document As You Go
Add comments for complex logic.

---

## 🆘 Emergency Commands

### Complete Reset
```bash
# Clean everything
pnpm clean

# Delete node_modules
rm -rf node_modules apps/*/node_modules packages/*/node_modules

# Reinstall
pnpm install

# Rebuild
pnpm build
```

### Fix Package Issues
```bash
# Clear pnpm cache
pnpm store prune

# Reinstall
pnpm install --force
```

### Fix Mobile Issues
```bash
cd apps/mobile

# Clear Metro cache
npx expo start --clear

# Clear iOS build
rm -rf ios/

# Clear Android build
rm -rf android/

# Rebuild native
npx expo prebuild --clean
```

---

## 📞 Support Checklist

Before asking for help:
1. [ ] Read error message completely
2. [ ] Check this reference card
3. [ ] Review documentation
4. [ ] Google the error
5. [ ] Check Stack Overflow
6. [ ] Try restarting dev server
7. [ ] Clear caches
8. [ ] Check git status

---

## ✨ Success Indicators

You're doing well if:
- ✅ No console errors
- ✅ TypeScript compiles cleanly
- ✅ Tests pass
- ✅ Features work as expected
- ✅ Code is readable
- ✅ Git history is clean
- ✅ Documentation is updated

---

## 🎉 Quick Wins

### Add a New Page (Web)
1. Create `apps/web/src/app/[locale]/newpage/page.tsx`
2. Add translations
3. Add navigation link
4. Done!

### Add a New Screen (Mobile)
1. Create `apps/mobile/app/newscreen.tsx`
2. Add to navigation
3. Add translations
4. Done!

### Add a New Hook
1. Create hook file in `packages/shared/src/features/user/hooks/`
2. Export from `index.ts`
3. Use in components
4. Done!

---

**Keep this card handy for quick reference! 📌**

Last Updated: October 9, 2025