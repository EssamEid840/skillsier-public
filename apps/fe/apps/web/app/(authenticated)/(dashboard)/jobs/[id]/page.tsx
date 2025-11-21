'use client';

import * as React from 'react';
import { useParams, useRouter } from 'next/navigation';
import { useJob } from '@skillsier/hooks';
import { Button, Card, CardHeader, CardTitle, CardContent } from '@skillsier/ui';

export default function JobDetailsPage() {
  const params = useParams();
  const router = useRouter();
  const jobId = params.id as string;
  const { data: job, isLoading, error } = useJob(jobId);

  if (isLoading) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">Loading job details...</div>
      </div>
    );
  }

  if (error || !job) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center text-error">Job not found</div>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <Button
        variant="ghost"
        onClick={() => router.back()}
        className="mb-4"
      >
        ← Back
      </Button>

      <Card>
        <CardHeader>
          <div className="flex items-start justify-between">
            <div>
              <CardTitle className="text-2xl">{job.title}</CardTitle>
              <div className="mt-2 flex gap-2">
                <span className="rounded-full bg-primary-100 px-3 py-1 text-xs font-medium text-primary">
                  {job.status}
                </span>
                <span className="rounded-full bg-secondary-100 px-3 py-1 text-xs font-medium">
                  {job.budgetType}
                </span>
              </div>
            </div>
            <Button>Apply Now</Button>
          </div>
        </CardHeader>
        <CardContent className="space-y-6">
          <div>
            <h3 className="mb-2 font-semibold">Description</h3>
            <p className="text-secondary-600">{job.description}</p>
          </div>

          <div>
            <h3 className="mb-2 font-semibold">Budget</h3>
            <p className="text-lg font-semibold">
              {job.budgetType === 'HOURLY'
                ? `$${job.budgetMin}-$${job.budgetMax}/hr`
                : `$${job.budgetAmount}`}
            </p>
          </div>

          <div>
            <h3 className="mb-2 font-semibold">Skills Required</h3>
            <div className="flex flex-wrap gap-2">
              {job.skills.map((skill) => (
                <span
                  key={skill}
                  className="rounded-lg bg-secondary-100 px-3 py-1 text-sm"
                >
                  {skill}
                </span>
              ))}
            </div>
          </div>

          <div className="grid gap-4 md:grid-cols-3">
            <div>
              <h4 className="text-sm text-secondary-600">Duration</h4>
              <p className="font-medium">{job.duration.replace(/_/g, ' ')}</p>
            </div>
            <div>
              <h4 className="text-sm text-secondary-600">Experience Level</h4>
              <p className="font-medium">{job.experienceLevel}</p>
            </div>
            <div>
              <h4 className="text-sm text-secondary-600">Proposals</h4>
              <p className="font-medium">{job.proposalCount}</p>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}