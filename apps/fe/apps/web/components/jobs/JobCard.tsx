import Link from 'next/link';
import type { Job } from '@skillsier/types';
import { Card, CardHeader, CardTitle, CardContent } from '@skillsier/ui';

interface JobCardProps {
  job: Job;
}

export function JobCard({ job }: JobCardProps) {
  return (
    <Link href={`/dashboard/jobs/${job.id}`}>
      <Card className="transition-shadow hover:shadow-md">
        <CardHeader>
          <CardTitle className="text-lg">{job.title}</CardTitle>
          <div className="flex gap-2">
            <span className="text-xs text-secondary-600">{job.category}</span>
            <span className="text-xs text-secondary-400">•</span>
            <span className="text-xs text-secondary-600">
              {job.proposalCount} proposals
            </span>
          </div>
        </CardHeader>
        <CardContent>
          <p className="mb-4 line-clamp-2 text-sm text-secondary-600">
            {job.description}
          </p>

          <div className="mb-4 flex flex-wrap gap-2">
            {job.skills.slice(0, 4).map((skill) => (
              <span
                key={skill}
                className="rounded-lg bg-secondary-100 px-2 py-1 text-xs"
              >
                {skill}
              </span>
            ))}
            {job.skills.length > 4 && (
              <span className="rounded-lg bg-secondary-100 px-2 py-1 text-xs">
                +{job.skills.length - 4} more
              </span>
            )}
          </div>

          <div className="flex items-center justify-between">
            <div className="font-semibold text-primary">
              {job.budgetType === 'HOURLY'
                ? `$${job.budgetMin}-$${job.budgetMax}/hr`
                : `$${job.budgetAmount}`}
            </div>
            <span className="text-xs text-secondary-500">
              Posted {new Date(job.createdAt).toLocaleDateString()}
            </span>
          </div>
        </CardContent>
      </Card>
    </Link>
  );
}