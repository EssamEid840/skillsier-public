'use client';

import { Button, Card } from '@skillsier/ui';
import { useJob } from '@skillsier/hooks';
import Link from 'next/link';
import { useParams } from 'next/navigation';

export function JobDetails() {
  const params = useParams();
  const jobId = params?.id as string;
  const { data: job, isLoading } = useJob(jobId);

  if (isLoading) {
    return (
      <div className="animate-pulse space-y-4">
        <div className="h-8 bg-gray-200 rounded w-3/4"></div>
        <div className="h-4 bg-gray-200 rounded w-full"></div>
        <div className="h-4 bg-gray-200 rounded w-5/6"></div>
      </div>
    );
  }

  if (!job) {
    return (
      <Card className="p-6 text-center">
        <p className="text-gray-500">Job not found</p>
      </Card>
    );
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <Card className="p-6">
        <div className="flex items-start justify-between mb-4">
          <div>
            <h1 className="text-3xl font-bold text-gray-900 mb-2">
              {job.title}
            </h1>
            <div className="flex items-center gap-4 text-sm text-gray-500">
              {job.createdAt && (
                <span>Posted {new Date(job.createdAt).toLocaleDateString()}</span>
              )}
              <span className="px-3 py-1 rounded-full bg-primary-100 text-primary-700 font-medium">
                {job.status}
              </span>
            </div>
          </div>
          <div className="flex items-center gap-2">
            <Link href={`/(authenticated)/(dashboard)/jobs/${job.id}/edit`}>
              <Button variant="outline">Edit</Button>
            </Link>
          </div>
        </div>

        <div className="prose max-w-none">
          <p className="text-gray-700 whitespace-pre-wrap">{job.description}</p>
        </div>
      </Card>

      {/* Job Info */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <Card className="p-6">
          <h3 className="text-sm font-medium text-gray-500 mb-2">Budget</h3>
          <p className="text-2xl font-bold text-primary-500">
            ${job.budget?.toLocaleString() || 'N/A'}
          </p>
        </Card>

        <Card className="p-6">
          <h3 className="text-sm font-medium text-gray-500 mb-2">Proposals</h3>
          <p className="text-2xl font-bold text-gray-900">
            {job.proposalCount || 0}
          </p>
        </Card>

        <Card className="p-6">
          <h3 className="text-sm font-medium text-gray-500 mb-2">Duration</h3>
          <p className="text-2xl font-bold text-gray-900">
            {job.duration || 'N/A'}
          </p>
        </Card>
      </div>

      {/* Skills */}
      {job.skills && job.skills.length > 0 && (
        <Card className="p-6">
          <h3 className="text-lg font-semibold text-gray-900 mb-3">
            Required Skills
          </h3>
          <div className="flex flex-wrap gap-2">
            {job.skills.map((skill, index) => (
              <span
                key={index}
                className="px-3 py-1 bg-gray-100 text-gray-700 rounded-full text-sm"
              >
                {skill}
              </span>
            ))}
          </div>
        </Card>
      )}

      {/* Actions */}
      <Card className="p-6">
        <div className="flex items-center gap-4">
          <Link href={`/(authenticated)/(dashboard)/jobs/${job.id}/proposals`}>
            <Button>View Proposals</Button>
          </Link>
          <Link href={`/(authenticated)/(dashboard)/jobs/${job.id}/applicants`}>
            <Button variant="outline">View Applicants</Button>
          </Link>
        </div>
      </Card>
    </div>
  );
}