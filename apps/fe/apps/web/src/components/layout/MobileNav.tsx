'use client';

import { useState } from 'react';
import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { useTranslations } from 'next-intl';
import { Menu, X, Home, BookOpen, Award, Users, Settings, User } from 'lucide-react';
import { Button } from '@skillsier/ui';
import { useAuth } from '@skillsier/shared';

export function MobileNav() {
  const [isOpen, setIsOpen] = useState(false);
  const pathname = usePathname();
  const t = useTranslations('common');
  const { user, logout } = useAuth();

  const navigation = [
    { name: t('dashboard'), href: '/dashboard', icon: Home },
    { name: t('courses'), href: '/dashboard/courses', icon: BookOpen },
    { name: t('skills'), href: '/dashboard/skills', icon: Award },
    { name: t('community'), href: '/dashboard/community', icon: Users },
    { name: t('settings'), href: '/dashboard/settings', icon: Settings },
  ];

  const toggleMenu = () => setIsOpen(!isOpen);

  return (
    <>
      {/* Mobile menu button */}
      <button
        type="button"
        className="lg:hidden inline-flex items-center justify-center p-2 rounded-lg text-gray-700 hover:text-gray-900 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-primary-500"
        onClick={toggleMenu}
        aria-label="Toggle menu"
      >
        {isOpen ? <X className="h-6 w-6" /> : <Menu className="h-6 w-6" />}
      </button>

      {/* Mobile menu overlay */}
      {isOpen && (
        <div
          className="fixed inset-0 bg-black bg-opacity-50 z-40 lg:hidden"
          onClick={toggleMenu}
          aria-hidden="true"
        />
      )}

      {/* Mobile menu panel */}
      <div
        className={`fixed top-0 ${
          typeof window !== 'undefined' && document.dir === 'rtl' ? 'left-0' : 'right-0'
        } bottom-0 w-64 bg-white shadow-xl z-50 transform transition-transform duration-300 ease-in-out lg:hidden ${
          isOpen ? 'translate-x-0' : typeof window !== 'undefined' && document.dir === 'rtl' ? '-translate-x-full' : 'translate-x-full'
        }`}
      >
        <div className="h-full flex flex-col">
          {/* Header */}
          <div className="flex items-center justify-between p-4 border-b border-gray-200">
            <div className="flex items-center gap-2">
              <div className="h-8 w-8 rounded-lg bg-gradient-to-br from-primary-600 to-purple-600" />
              <span className="text-lg font-bold text-gray-900">Skillsier</span>
            </div>
            <button
              onClick={toggleMenu}
              className="p-2 rounded-lg text-gray-500 hover:bg-gray-100"
              aria-label="Close menu"
            >
              <X className="h-5 w-5" />
            </button>
          </div>

          {/* User info */}
          {user && (
            <div className="p-4 border-b border-gray-200">
              <div className="flex items-center gap-3">
                <div className="h-12 w-12 rounded-full bg-gradient-to-br from-primary-500 to-purple-600 flex items-center justify-center text-white font-semibold">
                  {user.firstName?.[0]}{user.lastName?.[0]}
                </div>
                <div className="flex-1 min-w-0">
                  <p className="text-sm font-medium text-gray-900 truncate">
                    {user.firstName} {user.lastName}
                  </p>
                  <p className="text-xs text-gray-500 truncate">{user.email}</p>
                </div>
              </div>
            </div>
          )}

          {/* Navigation */}
          <nav className="flex-1 px-3 py-4 space-y-1 overflow-y-auto">
            {navigation.map((item) => {
              const isActive = pathname === item.href;
              return (
                <Link
                  key={item.name}
                  href={item.href}
                  onClick={toggleMenu}
                  className={`flex items-center gap-3 px-3 py-2 rounded-lg text-sm font-medium transition-colors ${
                    isActive
                      ? 'bg-primary-50 text-primary-700'
                      : 'text-gray-700 hover:bg-gray-50'
                  }`}
                >
                  <item.icon className="h-5 w-5" />
                  {item.name}
                </Link>
              );
            })}
          </nav>

          {/* Footer */}
          <div className="p-4 border-t border-gray-200 space-y-2">
            <Link href="/profile" onClick={toggleMenu}>
              <Button variant="outline" fullWidth className="justify-start">
                <User className="h-4 w-4 ltr:mr-2 rtl:ml-2" />
                {t('profile')}
              </Button>
            </Link>
            <Button
              variant="ghost"
              fullWidth
              onClick={() => {
                logout();
                toggleMenu();
              }}
              className="justify-start text-red-600 hover:text-red-700 hover:bg-red-50"
            >
              {t('logout')}
            </Button>
          </div>
        </div>
      </div>
    </>
  );
}