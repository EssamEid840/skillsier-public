'use client';

import { useAuth } from '@skillsier/auth';
import Link from 'next/link';
import { Button, Card, CardHeader, CardTitle, CardContent } from '@skillsier/ui';

export default function DashboardPage() {
  const { user } = useAuth();

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="mb-8 text-3xl font-bold">
        Welcome back, {user?.firstName}!
      </h1>

      <div className="grid gap-6 md:grid-cols-3">
        <Card>
          <CardHeader>
            <CardTitle>Active Jobs</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">0</div>
            <p className="text-sm text-secondary-600">No active jobs yet</p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Proposals</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">0</div>
            <p className="text-sm text-secondary-600">No proposals submitted</p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Earnings</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">$0</div>
            <p className="text-sm text-secondary-600">Total earnings</p>
          </CardContent>
        </Card>
      </div>

      <div className="mt-8">
        <h2 className="mb-4 text-2xl font-bold">Quick Actions</h2>
        <div className="flex gap-4">
          <Link href="/dashboard/jobs">
            <Button>Browse Jobs</Button>
          </Link>
          <Link href="/dashboard/jobs/create">
            <Button variant="outline">Post a Job</Button>
          </Link>
        </div>
      </div>
    </div>
  );
}