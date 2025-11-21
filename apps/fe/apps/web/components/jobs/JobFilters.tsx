'use client';

import * as React from 'react';
import { useJobsStore } from '@skillsier/stores';
import { JobStatus, JobBudgetType } from '@skillsier/types';
import { Button, Input, Card, CardHeader, CardTitle, CardContent } from '@skillsier/ui';

export function JobFilters() {
  const { filters, setFilters, resetFilters } = useJobsStore();
  const [search, setSearch] = React.useState(filters.search || '');

  const handleSearchSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    setFilters({ search: search || undefined });
  };

  return (
    <Card>
      <CardHeader>
        <CardTitle className="text-lg">Filters</CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        <form onSubmit={handleSearchSubmit}>
          <Input
            placeholder="Search jobs..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
          />
        </form>

        <div>
          <label className="mb-2 block text-sm font-medium">Status</label>
          <select
            className="w-full rounded-lg border border-secondary-300 px-3 py-2"
            value={filters.status || ''}
            onChange={(e) =>
              setFilters({
                status: e.target.value ? (e.target.value as JobStatus) : undefined,
              })
            }
          >
            <option value="">All</option>
            <option value={JobStatus.OPEN}>Open</option>
            <option value={JobStatus.IN_PROGRESS}>In Progress</option>
            <option value={JobStatus.CLOSED}>Closed</option>
          </select>
        </div>

        <div>
          <label className="mb-2 block text-sm font-medium">Budget Type</label>
          <select
            className="w-full rounded-lg border border-secondary-300 px-3 py-2"
            value={filters.budgetType || ''}
            onChange={(e) =>
              setFilters({
                budgetType: e.target.value
                  ? (e.target.value as JobBudgetType)
                  : undefined,
              })
            }
          >
            <option value="">All</option>
            <option value={JobBudgetType.HOURLY}>Hourly</option>
            <option value={JobBudgetType.FIXED}>Fixed</option>
          </select>
        </div>

        <Button variant="outline" className="w-full" onClick={resetFilters}>
          Reset Filters
        </Button>
      </CardContent>
    </Card>
  );
}