import * as React from 'react';
import { useTranslation } from '../hooks/useTranslation';
import { SUPPORTED_LOCALES, type SupportedLocale } from '../config';

export interface LocaleSwitcherProps {
  className?: string;
}

export const LocaleSwitcher: React.FC<LocaleSwitcherProps> = ({
  className,
}) => {
  const { currentLanguage, changeLanguage, t } = useTranslation();

  const handleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    changeLanguage(e.target.value as SupportedLocale);
  };

  return (
    <select
      value={currentLanguage}
      onChange={handleChange}
      className={className}
      aria-label={t('selectLanguage')}
    >
      {SUPPORTED_LOCALES.map(locale => (
        <option key={locale} value={locale}>
          {locale === 'en' ? 'English' : 'العربية'}
        </option>
      ))}
    </select>
  );
};