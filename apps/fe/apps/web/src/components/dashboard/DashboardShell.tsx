'use client';

import { ReactNode } from 'react';
import { useTranslations } from 'next-intl';

interface DashboardShellProps {
  children: ReactNode;
  title?: string;
  description?: string;
  action?: ReactNode;
}

export function DashboardShell({
  children,
  title,
  description,
  action,
}: DashboardShellProps) {
  const t = useTranslations('dashboard');

  return (
    <div className="space-y-6">
      {(title || description || action) && (
        <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
          <div>
            {title && (
              <h1 className="text-2xl font-bold text-gray-900 sm:text-3xl">
                {title}
              </h1>
            )}
            {description && (
              <p className="mt-2 text-sm text-gray-600">{description}</p>
            )}
          </div>
          {action && <div className="flex-shrink-0">{action}</div>}
        </div>
      )}
      <div>{children}</div>
    </div>
  );
}