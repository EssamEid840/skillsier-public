'use client';

import * as React from 'react';
import { useRouter } from 'next/navigation';
import { useAuth } from '@skillsier/auth';
import Link from 'next/link';
import { Button } from '@skillsier/ui';

export default function AuthenticatedLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const router = useRouter();
  const { user, isAuthenticated, isLoading, logout } = useAuth();

  React.useEffect(() => {
    if (!isLoading && !isAuthenticated) {
      router.push('/login');
    }
  }, [isAuthenticated, isLoading, router]);

  if (isLoading) {
    return (
      <div className="flex min-h-screen items-center justify-center">
        <div className="text-lg">Loading...</div>
      </div>
    );
  }

  if (!isAuthenticated) {
    return null;
  }

  return (
    <div className="min-h-screen">
      <nav className="border-b bg-white">
        <div className="container mx-auto flex h-16 items-center justify-between px-4">
          <Link href="/dashboard" className="text-2xl font-bold text-primary">
            Skillsier
          </Link>
          <div className="flex items-center gap-4">
            <Link href="/dashboard/jobs">
              <Button variant="ghost">Jobs</Button>
            </Link>
            <span className="text-sm text-secondary-600">
              {user?.firstName} {user?.lastName}
            </span>
            <Button variant="outline" onClick={() => logout()}>
              Logout
            </Button>
          </div>
        </div>
      </nav>
      <main>{children}</main>
    </div>
  );
}