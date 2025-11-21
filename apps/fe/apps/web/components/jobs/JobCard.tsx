'use client';

import { Card } from '@skillsier/ui';
import type { Job } from '@skillsier/types';
import Link from 'next/link';

interface JobCardProps {
  job: Job;
}

export function JobCard({ job }: JobCardProps) {
  return (
    <Link href={`/(authenticated)/(dashboard)/jobs/${job.id}`}>
      <Card className="p-6 hover:shadow-lg transition-shadow cursor-pointer">
        <div className="flex items-start justify-between mb-3">
          <h3 className="text-xl font-semibold text-gray-900">{job.title}</h3>
          <span className="px-3 py-1 text-xs font-medium rounded-full bg-primary-100 text-primary-700">
            {job.status}
          </span>
        </div>

        <p className="text-gray-600 mb-4 line-clamp-2">{job.description}</p>

        <div className="flex items-center justify-between">
          <div className="flex items-center gap-4 text-sm text-gray-500">
            {job.budget && (
              <span className="font-semibold text-primary-500">
                ${job.budget.toLocaleString()}
              </span>
            )}
            {job.createdAt && (
              <span>Posted {new Date(job.createdAt).toLocaleDateString()}</span>
            )}
          </div>
          
          {job.proposalCount !== undefined && (
            <span className="text-sm text-gray-500">
              {job.proposalCount} proposals
            </span>
          )}
        </div>
      </Card>
    </Link>
  );
}