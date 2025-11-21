import { useMutation, useQueryClient } from '@tanstack/react-query';
import { MockJobsClient } from '@skillsier/api';
import { mapJobDTOToDomain } from '@skillsier/types';
import type { UpdateJobRequest } from '@skillsier/types';
import { QUERY_KEYS } from '../config/query-client';

const jobsClient = new MockJobsClient();

export interface UpdateJobParams {
  jobId: string;
  data: UpdateJobRequest;
}

export const useUpdateJob = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async ({ jobId, data }: UpdateJobParams) => {
      const dto = await jobsClient.updateJob(jobId, data);
      return mapJobDTOToDomain(dto);
    },
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.job(variables.jobId) });
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.jobs });
    },
  });
};