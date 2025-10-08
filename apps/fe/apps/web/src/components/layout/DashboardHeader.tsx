'use client';

import { useAuth, useLogout } from '@skillsier/shared';
import { Avatar, Button } from '@skillsier/ui';
import { Bell, Search, Menu } from 'lucide-react';
import { useState } from 'react';

export function DashboardHeader() {
  const { user } = useAuth();
  const { mutate: logout } = useLogout();
  const [showUserMenu, setShowUserMenu] = useState(false);

  return (
    <header className="bg-white border-b border-gray-200 h-16 flex items-center px-4 sm:px-6 lg:px-8">
      <div className="flex items-center justify-between w-full">
        <div className="flex items-center gap-4 flex-1">
          <button className="lg:hidden p-2">
            <Menu className="h-5 w-5" />
          </button>
          <div className="relative flex-1 max-w-md">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-5 w-5 text-gray-400" />
            <input
              type="search"
              placeholder="Search courses, skills..."
              className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent"
            />
          </div>
        </div>

        <div className="flex items-center gap-4">
          <button className="relative p-2 text-gray-600 hover:text-gray-900">
            <Bell className="h-5 w-5" />
            <span className="absolute top-1 right-1 h-2 w-2 bg-red-500 rounded-full" />
          </button>

          <div className="relative">
            <button
              onClick={() => setShowUserMenu(!showUserMenu)}
              className="flex items-center gap-2"
            >
              <Avatar
                src={user?.avatar}
                alt={user?.username || 'User'}
                size="md"
              />
            </button>

            {showUserMenu && (
              <div className="absolute right-0 mt-2 w-48 bg-white rounded-lg shadow-lg border border-gray-200 py-1 z-50">
                <div className="px-4 py-2 border-b border-gray-200">
                  <p className="font-medium text-gray-900">{user?.firstName} {user?.lastName}</p>
                  <p className="text-sm text-gray-600">{user?.email}</p>
                </div>
                <a href="/dashboard/profile" className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-50">
                  Profile
                </a>
                <a href="/dashboard/settings" className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-50">
                  Settings
                </a>
                <button
                  onClick={() => logout()}
                  className="w-full text-left px-4 py-2 text-sm text-red-600 hover:bg-red-50"
                >
                  Sign out
                </button>
              </div>
            )}
          </div>
        </div>
      </div>
    </header>
  );
}