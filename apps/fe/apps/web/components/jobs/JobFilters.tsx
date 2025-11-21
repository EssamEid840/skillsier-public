'use client';

import { Button, Card } from '@skillsier/ui';
import { useJobsStore } from '@skillsier/stores';
import { JobStatus } from '@skillsier/types';

export function JobFilters() {
  const { filters, setFilters, resetFilters } = useJobsStore();

  const statuses = [
    { value: JobStatus.DRAFT, label: 'Draft' },
    { value: JobStatus.ACTIVE, label: 'Active' },
    { value: JobStatus.CLOSED, label: 'Closed' },
    { value: JobStatus.CANCELLED, label: 'Cancelled' },
  ];

  return (
    <Card className="p-6">
      <div className="flex items-center justify-between mb-4">
        <h3 className="text-lg font-semibold text-gray-900">Filters</h3>
        <Button
          variant="ghost"
          size="sm"
          onClick={resetFilters}
          className="text-primary-500"
        >
          Reset
        </Button>
      </div>

      <div className="space-y-4">
        {/* Status Filter */}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Status
          </label>
          <div className="space-y-2">
            {statuses.map((status) => (
              <label key={status.value} className="flex items-center">
                <input
                  type="checkbox"
                  checked={filters.status?.includes(status.value)}
                  onChange={(e) => {
                    const currentStatuses = filters.status || [];
                    if (e.target.checked) {
                      setFilters({
                        status: [...currentStatuses, status.value],
                      });
                    } else {
                      setFilters({
                        status: currentStatuses.filter((s) => s !== status.value),
                      });
                    }
                  }}
                  className="rounded border-gray-300 text-primary-500 focus:ring-primary-500"
                />
                <span className="ml-2 text-sm text-gray-700">{status.label}</span>
              </label>
            ))}
          </div>
        </div>

        {/* Budget Range */}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Budget Range
          </label>
          <div className="flex items-center gap-2">
            <input
              type="number"
              placeholder="Min"
              value={filters.budgetMin || ''}
              onChange={(e) =>
                setFilters({ budgetMin: Number(e.target.value) || undefined })
              }
              className="flex-1 rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500"
            />
            <span className="text-gray-500">-</span>
            <input
              type="number"
              placeholder="Max"
              value={filters.budgetMax || ''}
              onChange={(e) =>
                setFilters({ budgetMax: Number(e.target.value) || undefined })
              }
              className="flex-1 rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500"
            />
          </div>
        </div>

        {/* Search */}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Search
          </label>
          <input
            type="text"
            placeholder="Search jobs..."
            value={filters.search || ''}
            onChange={(e) => setFilters({ search: e.target.value || undefined })}
            className="w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500"
          />
        </div>
      </div>
    </Card>
  );
}