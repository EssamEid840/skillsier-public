import { useQuery } from '@tanstack/react-query';
import { MockJobsClient } from '@skillsier/api';
import { mapJobListDTOToDomain } from '@skillsier/types';
import type { JobListRequest, JobFilters } from '@skillsier/types';
import { QUERY_KEYS } from '../config/query-client';

const jobsClient = new MockJobsClient();

export const useJobs = (filters?: JobFilters, page: number = 1, limit: number = 10) => {
  const params: JobListRequest = {
    page,
    limit,
    search: filters?.search,
    status: filters?.status,
    budget_type: filters?.budgetType,
    budget_min: filters?.budgetMin,
    budget_max: filters?.budgetMax,
    skills: filters?.skills,
    category: filters?.category,
    experience_level: filters?.experienceLevel,
    duration: filters?.duration,
  };

  return useQuery({
    queryKey: QUERY_KEYS.jobsList(params),
    queryFn: async () => {
      const dto = await jobsClient.listJobs(params);
      return mapJobListDTOToDomain(dto);
    },
  });
};