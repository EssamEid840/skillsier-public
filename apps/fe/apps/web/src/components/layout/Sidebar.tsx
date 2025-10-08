'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { Home, BookOpen, Award, Users, Settings, BarChart3 } from 'lucide-react';

const navigation = [
  { name: 'Dashboard', href: '/dashboard', icon: Home },
  { name: 'My Courses', href: '/dashboard/courses', icon: BookOpen },
  { name: 'Skills', href: '/dashboard/skills', icon: Award },
  { name: 'Community', href: '/dashboard/community', icon: Users },
  { name: 'Analytics', href: '/dashboard/analytics', icon: BarChart3 },
  { name: 'Settings', href: '/dashboard/settings', icon: Settings },
];

export function Sidebar() {
  const pathname = usePathname();

  return (
    <aside className="hidden lg:flex lg:flex-col w-64 bg-white border-r border-gray-200">
      <div className="flex items-center gap-2 h-16 px-6 border-b border-gray-200">
        <div className="h-8 w-8 rounded-lg bg-gradient-to-br from-primary-600 to-purple-600" />
        <span className="text-xl font-bold text-gray-900">Skillsier</span>
      </div>

      <nav className="flex-1 px-3 py-4 space-y-1">
        {navigation.map((item) => {
          const isActive = pathname === item.href;
          return (
            <Link
              key={item.name}
              href={item.href}
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

      <div className="p-4 border-t border-gray-200">
        <div className="bg-gradient-to-r from-primary-500 to-purple-600 rounded-lg p-4 text-white">
          <p className="font-semibold mb-1">Upgrade to Pro</p>
          <p className="text-sm text-primary-100 mb-3">
            Unlock premium features and courses
          </p>
          <button className="w-full bg-white text-primary-600 rounded-lg px-4 py-2 text-sm font-medium hover:bg-primary-50 transition-colors">
            Learn More
          </button>
        </div>
      </div>
    </aside>
  );
}
