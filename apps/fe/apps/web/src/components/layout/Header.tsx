'use client';

import Link from 'next/link';
import { useState } from 'react';
import { useTranslations } from 'next-intl';
import { Button } from '@skillsier/ui';
import { Menu, X } from 'lucide-react';
import { useAuth } from '@skillsier/shared';
import { LanguageSwitcher } from './LanguageSwitcher';

export function Header() {
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);
  const { isAuthenticated } = useAuth();
  const t = useTranslations('auth');

  return (
    <header className="fixed top-0 left-0 right-0 z-50 bg-white/80 backdrop-blur-lg border-b border-gray-200">
      <nav className="container mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          <div className="flex items-center">
            <Link href="/" className="flex items-center gap-2">
              <div className="h-8 w-8 rounded-lg bg-gradient-to-br from-primary-600 to-purple-600" />
              <span className="text-xl font-bold text-gray-900">Skillsier</span>
            </Link>
          </div>

          <div className="hidden md:flex items-center gap-6">
            <LanguageSwitcher />
            {isAuthenticated ? (
              <Link href="/dashboard">
                <Button>Dashboard</Button>
              </Link>
            ) : (
              <>
                <Link href="/login">
                  <Button variant="ghost">{t('login')}</Button>
                </Link>
                <Link href="/register">
                  <Button>{t('register')}</Button>
                </Link>
              </>
            )}
          </div>

          <button
            className="md:hidden p-2"
            onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
          >
            {mobileMenuOpen ? <X className="h-6 w-6" /> : <Menu className="h-6 w-6" />}
          </button>
        </div>
      </nav>
    </header>
  );
}
