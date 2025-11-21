'use client';

import { Card } from '@skillsier/ui';
import { useParams } from 'next/navigation';

export function JobApplicants() {
  const params = useParams();
  const jobId = params?.id as string;

  // TODO: Implement useJobApplicants hook
  const applicants: any[] = [];
  const isLoading = false;

  if (isLoading) {
    return (
      <div className="space-y-4">
        {[...Array(3)].map((_, i) => (
          <div key={i} className="animate-pulse">
            <Card className="p-6">
              <div className="flex items-center gap-4">
                <div className="w-12 h-12 bg-gray-200 rounded-full"></div>
                <div className="flex-1 space-y-2">
                  <div className="h-4 bg-gray-200 rounded w-1/4"></div>
                  <div className="h-3 bg-gray-200 rounded w-1/2"></div>
                </div>
              </div>
            </Card>
          </div>
        ))}
      </div>
    );
  }

  if (applicants.length === 0) {
    return (
      <Card className="p-12 text-center">
        <p className="text-gray-500">No applicants yet</p>
        <p className="text-sm text-gray-400 mt-2">
          Applicants will appear here once they submit proposals
        </p>
      </Card>
    );
  }

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between mb-6">
        <h2 className="text-2xl font-bold text-gray-900">
          Applicants ({applicants.length})
        </h2>
      </div>

      {applicants.map((applicant) => (
        <Card key={applicant.id} className="p-6">
          <div className="flex items-start justify-between">
            <div className="flex items-center gap-4">
              <div className="w-12 h-12 bg-primary-100 rounded-full flex items-center justify-center">
                <span className="text-primary-500 font-semibold">
                  {applicant.name?.charAt(0) || 'U'}
                </span>
              </div>
              <div>
                <h3 className="text-lg font-semibold text-gray-900">
                  {applicant.name}
                </h3>
                <p className="text-sm text-gray-500">{applicant.title}</p>
              </div>
            </div>
            <span className="text-sm text-gray-500">
              Applied {new Date(applicant.appliedAt).toLocaleDateString()}
            </span>
          </div>
        </Card>
      ))}
    </div>
  );
}