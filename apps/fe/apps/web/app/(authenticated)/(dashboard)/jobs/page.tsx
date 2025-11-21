'use client';

import * as React from 'react';
import Link from 'next/link';
import { useJobs } from '@skillsier/hooks';
import { useJobsStore } from '@skillsier/stores';
import { Button, Card, CardHeader, CardTitle, CardContent } from '@skillsier/ui';
import { JobCard } from '@/components/jobs/JobCard';
import { JobFilters } from '@/components/jobs/JobFilters';

export default function JobsPage() {
  const { filters, page, limit } = useJobsStore();
  const { data, isLoading, error } = useJobs(filters, page, limit);

  if (isLoading) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">Loading jobs...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center text-error">
          Error loading jobs: {error.message}
        </div>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="mb-6 flex items-center justify-between">
        <h1 className="text-3xl font-bold">Browse Jobs</h1>
        <Link href="/dashboard/jobs/create">
          <Button>Post a Job</Button>
        </Link>
      </div>

      <div className="grid gap-6 lg:grid-cols-4">
        <aside className="lg:col-span-1">
          <JobFilters />
        </aside>

        <main className="lg:col-span-3">
          {data?.jobs.length === 0 ? (
            <Card>
              <CardContent className="py-8 text-center">
                <p className="text-secondary-600">No jobs found</p>
              </CardContent>
            </Card>
          ) : (
            <div className="space-y-4">
              {data?.jobs.map((job) => (
                <JobCard key={job.id} job={job} />
              ))}
            </div>
          )}

          {data && data.pagination.totalPages > 1 && (
            <div className="mt-6 flex justify-center gap-2">
              <Button
                variant="outline"
                disabled={page === 1}
                onClick={() => useJobsStore.getState().setPage(page - 1)}
              >
                Previous
              </Button>
              <span className="flex items-center px-4">
                Page {page} of {data.pagination.totalPages}
              </span>
              <Button
                variant="outline"
                disabled={page === data.pagination.totalPages}
                onClick={() => useJobsStore.getState().setPage(page + 1)}
              >
                Next
              </Button>
            </div>
          )}
        </main>
      </div>
    </div>
  );
}