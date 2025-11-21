import { useTranslation as useI18nextTranslation } from 'react-i18next';
import type { SupportedLocale } from '../config';

export const useTranslation = (namespace: string = 'common') => {
  const { t, i18n } = useI18nextTranslation(namespace);

  const changeLanguage = async (locale: SupportedLocale) => {
    await i18n.changeLanguage(locale);
    
    if (typeof document !== 'undefined') {
      const isRTL = locale === 'ar';
      document.documentElement.dir = isRTL ? 'rtl' : 'ltr';
      document.documentElement.lang = locale;
    }
  };

  return {
    t,
    i18n,
    changeLanguage,
    currentLanguage: i18n.language as SupportedLocale,
    isRTL: i18n.dir() === 'rtl',
  };
};