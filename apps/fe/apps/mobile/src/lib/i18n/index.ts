import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import * as Localization from 'expo-localization';
import { I18nManager } from 'react-native';

// Import translations
import en from '@skillsier/shared/lib/i18n/translations/en.json';
import ar from '@skillsier/shared/lib/i18n/translations/ar.json';

const resources = {
  en: { translation: en },
  ar: { translation: ar },
};

// Detect device locale
const deviceLocale = Localization.getLocales()[0]?.languageCode || 'en';

i18n
  .use(initReactI18next)
  .init({
    resources,
    lng: deviceLocale,
    fallbackLng: 'en',
    compatibilityJSON: 'v3',
    interpolation: {
      escapeValue: false,
    },
    react: {
      useSuspense: false,
    },
  });

// Handle RTL for Arabic
export const setLanguage = async (language: string) => {
  await i18n.changeLanguage(language);
  const isRTL = language === 'ar';
  
  if (I18nManager.isRTL !== isRTL) {
    I18nManager.forceRTL(isRTL);
    // Note: App needs to restart for RTL change to take effect
  }
};

export default i18n;
