export const SUPPORTED_LOCALES = {
  en: {
    code: 'en',
    name: 'English',
    nativeName: 'English',
    direction: 'ltr',
    flag: '🇺🇸',
  },
  ar: {
    code: 'ar',
    name: 'Arabic',
    nativeName: 'العربية',
    direction: 'rtl',
    flag: '🇸🇦',
  },
} as const;

export type LocaleCode = keyof typeof SUPPORTED_LOCALES;

export const DEFAULT_LOCALE: LocaleCode = 'en';
export const FALLBACK_LOCALE: LocaleCode = 'en';