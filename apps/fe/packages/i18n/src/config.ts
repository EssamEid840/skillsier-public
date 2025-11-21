import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';

const loadLocaleResources = () => {
  const locales = ['en', 'ar'];
  const namespaces = ['common', 'auth', 'jobs'];
  const resources: Record<string, any> = {};

  locales.forEach(locale => {
    resources[locale] = {};
    namespaces.forEach(namespace => {
      try {
        resources[locale][namespace] = require(`./locales/${locale}/${namespace}.json`);
      } catch (error) {
        console.warn(`Failed to load ${locale}/${namespace}.json`);
        resources[locale][namespace] = {};
      }
    });
  });

  return resources;
};

i18n.use(initReactI18next).init({
  resources: loadLocaleResources(),
  lng: 'en',
  fallbackLng: 'en',
  defaultNS: 'common',
  ns: ['common', 'auth', 'jobs'],
  interpolation: {
    escapeValue: false,
  },
  react: {
    useSuspense: false,
  },
});

export { i18n };

export const SUPPORTED_LOCALES = ['en', 'ar'] as const;
export type SupportedLocale = (typeof SUPPORTED_LOCALES)[number];

export const RTL_LOCALES: SupportedLocale[] = ['ar'];

export const isRTL = (locale: string): boolean => {
  return RTL_LOCALES.includes(locale as SupportedLocale);
};