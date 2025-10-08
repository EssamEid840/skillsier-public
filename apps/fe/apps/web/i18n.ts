// apps/web/i18n.ts
import { getRequestConfig } from 'next-intl/server';
import { notFound } from 'next/navigation';

export const locales = ['en', 'ar', 'zh', 'hi', 'de', 'fr', 'tr', 'es', 'ru'] as const;
export type Locale = (typeof locales)[number];

export default getRequestConfig(async ({ locale }) => {
  if (!locales.includes(locale as Locale)) notFound();

  return {
    messages: (await import(`../../packages/shared/src/lib/i18n/translations/${locale}.json`))
      .default,
  };
});