'use client';

import { useLocale } from 'next-intl';
import { useRouter, usePathname } from 'next/navigation';
import { Globe } from 'lucide-react';

const languages = [
  { code: 'en', name: 'English', flag: '🇺🇸' },
  { code: 'ar', name: 'العربية', flag: '🇸🇦' },
];

export function LanguageSwitcher() {
  const locale = useLocale();
  const router = useRouter();
  const pathname = usePathname();

  const switchLanguage = (newLocale: string) => {
    const segments = pathname.split('/');
    segments[1] = newLocale;
    router.push(segments.join('/'));
  };

  return (
    <div className="relative group">
      <button className="flex items-center gap-2 px-3 py-2 rounded-lg hover:bg-gray-100 transition-colors">
        <Globe className="h-5 w-5" />
        <span className="text-sm font-medium">
          {languages.find((l) => l.code === locale)?.flag}
        </span>
      </button>
      <div className="absolute top-full mt-2 right-0 bg-white rounded-lg shadow-lg border border-gray-200 py-1 min-w-[150px] opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all">
        {languages.map((lang) => (
          <button
            key={lang.code}
            onClick={() => switchLanguage(lang.code)}
            className={`w-full px-4 py-2 text-left hover:bg-gray-50 flex items-center gap-2 ${
              locale === lang.code ? 'bg-primary-50 text-primary-600' : 'text-gray-700'
            }`}
          >
            <span>{lang.flag}</span>
            <span className="text-sm">{lang.name}</span>
          </button>
        ))}
      </div>
    </div>
  );
}
