'use client';

import Link from 'next/link';
import { useTranslations } from 'next-intl';
import { Button } from '@skillsier/ui';
import { ArrowRight, Play } from 'lucide-react';

export function Hero() {
  const t = useTranslations('landing.hero');

  return (
    <section className="relative overflow-hidden bg-gradient-to-br from-primary-50 via-white to-purple-50 pt-20 pb-32">
      <div className="container mx-auto px-4 sm:px-6 lg:px-8">
        <div className="grid grid-cols-1 gap-12 lg:grid-cols-2 lg:gap-8 items-center">
          <div className="max-w-2xl">
            <h1 className="text-5xl font-bold tracking-tight text-gray-900 sm:text-6xl lg:text-7xl">
              {t('title')}{' '}
              <span className="bg-gradient-to-r from-primary-600 to-purple-600 bg-clip-text text-transparent">
                {t('titleHighlight')}
              </span>
            </h1>
            <p className="mt-6 text-lg leading-8 text-gray-600 sm:text-xl">
              {t('subtitle')}
            </p>
            <div className="mt-10 flex flex-col gap-4 sm:flex-row sm:gap-6">
              <Link href="/register">
                <Button size="lg" className="w-full sm:w-auto">
                  {t('getStarted')}
                  <ArrowRight className="ltr:ml-2 rtl:mr-2 h-5 w-5" />
                </Button>
              </Link>
              <Button variant="outline" size="lg" className="w-full sm:w-auto">
                <Play className="ltr:mr-2 rtl:ml-2 h-5 w-5" />
                {t('watchDemo')}
              </Button>
            </div>
            <div className="mt-8 flex items-center gap-8">
              <div>
                <p className="text-3xl font-bold text-gray-900">500K+</p>
                <p className="text-sm text-gray-600">{t('stats.learners')}</p>
              </div>
              <div className="h-12 w-px bg-gray-300" />
              <div>
                <p className="text-3xl font-bold text-gray-900">2,000+</p>
                <p className="text-sm text-gray-600">{t('stats.clients')}</p>
              </div>
              <div className="h-12 w-px bg-gray-300" />
              <div>
                <p className="text-3xl font-bold text-gray-900">98%</p>
                <p className="text-sm text-gray-600">{t('stats.satisfaction')}</p>
              </div>
            </div>
          </div>
          <div className="relative lg:ml-auto">
            <div className="relative rounded-2xl bg-gradient-to-br from-primary-500 to-purple-600 p-2 shadow-2xl">
              <img
                src="/images/dashboard-preview.png"
                alt="Skillsier Dashboard"
                className="rounded-xl"
                width={600}
                height={400}
              />
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}