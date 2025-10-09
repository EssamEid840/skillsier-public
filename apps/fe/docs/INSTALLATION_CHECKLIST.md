# ✅ Skillsier Frontend - Installation Checklist

Follow these steps to complete your Skillsier frontend setup.

---

## 📋 Pre-Installation

### ✅ Prerequisites Check
- [ ] Node.js 20.x or higher installed
- [ ] pnpm 9.x installed (`npm install -g pnpm@9`)
- [ ] Git installed
- [ ] For iOS: macOS with Xcode
- [ ] For Android: Android Studio with SDK

---

## 📁 Step 1: Copy All Files

Copy these files from the artifacts to your project:

### Web Components
- [ ] `apps/web/src/components/layout/MobileNav.tsx`
- [ ] `apps/web/src/components/dashboard/DashboardShell.tsx`
- [ ] `apps/web/src/components/dashboard/StatsCard.tsx`
- [ ] `apps/web/src/app/[locale]/(auth)/register/page.tsx`

### Mobile Components
- [ ] `apps/mobile/index.js`
- [ ] `apps/mobile/app/+not-found.tsx`
- [ ] `apps/mobile/app/(auth)/register.tsx`
- [ ] `apps/mobile/app/(tabs)/_layout.tsx`
- [ ] `apps/mobile/app/(tabs)/courses.tsx`
- [ ] `apps/mobile/src/components/landing/HeroMobile.tsx`
- [ ] `apps/mobile/src/components/landing/FeaturesMobile.tsx`
- [ ] `apps/mobile/src/components/navigation/TabBar.tsx`
- [ ] `apps/mobile/src/components/navigation/DrawerContent.tsx`

### Shared Hooks
- [ ] `packages/shared/package.json`
- [ ] `packages/shared/src/features/user/hooks/index.ts`
- [ ] `packages/shared/src/features/user/hooks/useClientProfile.ts`
- [ ] `packages/shared/src/features/user/hooks/useDeleteAvatar.ts`
- [ ] `packages/shared/src/features/user/hooks/useFreelancerSkills.ts`
- [ ] `packages/shared/src/features/user/hooks/useAddSkill.ts`
- [ ] `packages/shared/src/features/user/hooks/useUpdateSkill.ts`
- [ ] `packages/shared/src/features/user/hooks/useDeleteSkill.ts`
- [ ] `packages/shared/src/features/user/hooks/useWorkExperience.ts`
- [ ] `packages/shared/src/features/user/hooks/useAddWorkExperience.ts`
- [ ] `packages/shared/src/features/user/hooks/useUpdateWorkExperience.ts`
- [ ] `packages/shared/src/features/user/hooks/useDeleteWorkExperience.ts`
- [ ] `packages/shared/src/features/user/hooks/useEducation.ts`
- [ ] `packages/shared/src/features/user/hooks/useAddEducation.ts`
- [ ] `packages/shared/src/features/user/hooks/useUpdateEducation.ts`
- [ ] `packages/shared/src/features/user/hooks/useDeleteEducation.ts`
- [ ] `packages/shared/src/features/user/hooks/useCertifications.ts`
- [ ] `packages/shared/src/features/user/hooks/useAddCertification.ts`
- [ ] `packages/shared/src/features/user/hooks/useUpdateCertification.ts`
- [ ] `packages/shared/src/features/user/hooks/useDeleteCertification.ts`
- [ ] `packages/shared/src/features/user/hooks/usePortfolio.ts`
- [ ] `packages/shared/src/features/user/hooks/useAddPortfolioItem.ts`
- [ ] `packages/shared/src/features/user/hooks/useUpdatePortfolioItem.ts`
- [ ] `packages/shared/src/features/user/hooks/useDeletePortfolioItem.ts`
- [ ] `packages/shared/src/features/user/hooks/useUploadPortfolioImage.ts`
- [ ] `packages/shared/src/features/user/hooks/useUserStats.ts`
- [ ] `packages/shared/src/features/user/hooks/useEarnings.ts`
- [ ] `packages/shared/src/features/user/hooks/useReviews.ts`

---

## 📦 Step 2: Install Dependencies

```bash
cd skillsier-fe

# Install all dependencies
pnpm install

# This will install dependencies for:
# - Root workspace
# - apps/web
# - apps/mobile
# - packages/shared
# - packages/ui
# - packages/types
# - packages/config
```

---

## 🔧 Step 3: Environment Variables

### Web Environment (.env.local)
Create `apps/web/.env.local`:
```bash
NEXT_PUBLIC_API_URL=http://localhost:8080/api
NEXT_PUBLIC_KEYCLOAK_URL=http://localhost:8080
NEXT_PUBLIC_KEYCLOAK_REALM=skillsier
NEXT_PUBLIC_KEYCLOAK_CLIENT_ID=skillsier-web
```

### Mobile Environment (.env)
Create `apps/mobile/.env`:
```bash
EXPO_PUBLIC_API_URL=http://localhost:8080/api
EXPO_PUBLIC_KEYCLOAK_URL=http://localhost:8080
EXPO_PUBLIC_KEYCLOAK_REALM=skillsier
EXPO_PUBLIC_KEYCLOAK_CLIENT_ID=skillsier-mobile
```

---

## 🚀 Step 4: Start Development

### Option A: Start Both Apps
```bash
pnpm dev
```

### Option B: Start Individually

**Web Only:**
```bash
pnpm dev:web
# Open http://localhost:3000
```

**Mobile Only:**
```bash
pnpm dev:mobile
# Scan QR code with Expo Go app
```

---

## 📱 Step 5: Mobile Development Build (Optional)

For full native features:

```bash
cd apps/mobile

# Clean and rebuild native code
npx expo prebuild --clean

# Run on iOS
npx expo run:ios

# Run on Android
npx expo run:android
```

---

## ✅ Step 6: Verify Installation

### Web Application Tests
- [ ] Navigate to `http://localhost:3000`
- [ ] Landing page loads correctly
- [ ] Language switcher works (top right)
- [ ] Switch to Arabic - RTL layout works
- [ ] Register page accessible at `/register`
- [ ] Login page accessible at `/login`
- [ ] Form validation works
- [ ] Mobile menu opens on small screens

### Mobile Application Tests
- [ ] App launches without errors
- [ ] Landing screen shows
- [ ] Register screen works
- [ ] Login screen works
- [ ] Tab navigation works (Dashboard, Jobs, Skills, Profile)
- [ ] Language switcher in profile works
- [ ] Pull to refresh works on jobs page
- [ ] Profile screen displays user info

---

## 🔍 Step 7: Test API Integration

### Backend Connection
- [ ] Backend API is running on `http://localhost:8080`
- [ ] Keycloak is configured and accessible
- [ ] Test registration creates user in backend
- [ ] Test login returns valid token
- [ ] Profile API endpoints respond correctly

### Test Endpoints
```bash
# Check backend health
curl http://localhost:8080/health

# Check API endpoints
curl http://localhost:8080/api/users/profile
```

---

## 🎨 Step 8: Customization (Optional)

### Brand Colors
Edit `packages/config/src/tailwind/base.js`:
```javascript
colors: {
  primary: {
    50: '#your-color',
    // ... customize
  }
}
```

### Translations
Add/edit translations in:
- `packages/shared/src/lib/i18n/translations/en.json`
- `packages/shared/src/lib/i18n/translations/ar.json`
- etc.

---

## 🐛 Troubleshooting

### Issue: "Cannot find module '@skillsier/shared'"
**Solution:**
```bash
pnpm install
pnpm build:packages
```

### Issue: "Type errors in TypeScript"
**Solution:**
```bash
pnpm type-check
# Fix any type errors shown
```

### Issue: Metro bundler cache issues (Mobile)
**Solution:**
```bash
cd apps/mobile
npx expo start --clear
```

### Issue: Web app not loading
**Solution:**
```bash
cd apps/web
rm -rf .next
pnpm dev:web
```

### Issue: Hooks not found
**Solution:**
Make sure you copied all hook files and the `index.ts` file that exports them.

### Issue: Translation keys not found
**Solution:**
Check that translation files exist in `packages/shared/src/lib/i18n/translations/`

---

## 📊 Step 9: Verify Features

### Authentication Features
- [ ] User can register
- [ ] User can login
- [ ] User can logout
- [ ] Protected routes redirect to login
- [ ] Token persists after refresh

### Profile Features
- [ ] View freelancer profile
- [ ] View client profile
- [ ] Upload avatar
- [ ] Delete avatar
- [ ] Edit profile information

### Skills Management
- [ ] View all skills
- [ ] Add new skill
- [ ] Edit skill
- [ ] Delete skill

### Work Experience
- [ ] View work experience
- [ ] Add experience
- [ ] Edit experience
- [ ] Delete experience

### Education
- [ ] View education
- [ ] Add education
- [ ] Edit education
- [ ] Delete education

### Certifications
- [ ] View certifications
- [ ] Add certification
- [ ] Edit certification
- [ ] Delete certification

### Portfolio
- [ ] View portfolio
- [ ] Add portfolio item
- [ ] Edit portfolio item
- [ ] Delete portfolio item
- [ ] Upload images

### Jobs/Gigs
- [ ] Browse jobs
- [ ] Search jobs
- [ ] Filter jobs (UI ready)
- [ ] View job details

---

## 🎯 Step 10: Production Build

### Web Production Build
```bash
cd apps/web
pnpm build

# Preview production build
pnpm start
```

### Mobile Production Build
```bash
cd apps/mobile

# Configure app.json with your credentials
# Then build:

# iOS
eas build --platform ios

# Android
eas build --platform android

# Submit to stores
eas submit --platform all
```

---

## 📈 Performance Checklist

- [ ] Images are optimized
- [ ] API calls are cached (TanStack Query)
- [ ] Components use React.memo where needed
- [ ] Large lists use virtualization
- [ ] Bundle size is reasonable (<500KB initial)
- [ ] Lighthouse score > 90 (Web)
- [ ] App starts in < 3 seconds (Mobile)

---

## 🔒 Security Checklist

- [ ] API endpoints use HTTPS in production
- [ ] Tokens stored securely (localStorage/MMKV)
- [ ] Sensitive data not logged
- [ ] Input validation on all forms
- [ ] XSS protection enabled
- [ ] CSRF tokens in place (if applicable)
- [ ] Environment variables not committed to git

---

## 📚 Documentation

Additional documentation available in:
- `apps/fe/docs/SUMMARY.md` - Complete overview
- `apps/fe/docs/ARCHITECTURE.md` - Architecture details
- `apps/fe/docs/I18N.md` - Internationalization guide
- `apps/fe/docs/REACT_19.md` - React 19 features
- `apps/fe/docs/FREELANCING.md` - Freelancing features
- `apps/fe/docs/PERFORMANCE.md` - Performance guide

---

## ✨ Success Criteria

Your installation is successful when:

✅ Web app runs on `http://localhost:3000`  
✅ Mobile app runs in Expo Go or simulator  
✅ User can register and login  
✅ Navigation works on both platforms  
✅ Language switching works  
✅ Profile management works  
✅ All API endpoints connect successfully  
✅ No console errors  
✅ TypeScript compiles without errors  
✅ Tests pass (if implemented)  

---

## 🎉 You're Done!

Congratulations! Your Skillsier frontend is now fully set up and ready for development.

### Next Steps:
1. **Connect to your backend** - Update API endpoints in env files
2. **Customize branding** - Update colors, logos, and content
3. **Add more features** - Jobs, proposals, contracts, reviews
4. **Test thoroughly** - Manual and automated testing
5. **Deploy** - Web to Vercel/Netlify, Mobile to App Store/Play Store

---

## 🆘 Need Help?

If you encounter any issues:

1. **Check this checklist** - Make sure all steps are complete
2. **Review documentation** - Check the docs folder
3. **Verify backend** - Ensure backend API is running
4. **Check console** - Look for error messages
5. **Clear cache** - Try clearing build caches

---

## 📞 Support Resources

- **Project Documentation**: `apps/fe/docs/`
- **API Documentation**: Check your backend docs
- **React Query Docs**: https://tanstack.com/query
- **Next.js Docs**: https://nextjs.org/docs
- **Expo Docs**: https://docs.expo.dev

---

**Built with ❤️ for Skillsier**

Last Updated: October 2025