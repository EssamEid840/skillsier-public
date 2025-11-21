import { useMutation, useQueryClient } from '@tanstack/react-query';
import { MockJobsClient } from '@skillsier/api';
import { mapJobDTOToDomain } from '@skillsier/types';
import type { CreateJobRequest } from '@skillsier/types';
import { QUERY_KEYS } from '../config/query-client';

const jobsClient = new MockJobsClient();

export const useCreateJob = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (data: CreateJobRequest) => {
      const dto = await jobsClient.createJob(data);
      return mapJobDTOToDomain(dto);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.jobs });
    },
  });
};