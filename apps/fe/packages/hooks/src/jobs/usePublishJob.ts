import { useMutation, useQueryClient } from '@tanstack/react-query';
import { MockJobsClient } from '@skillsier/api';
import { mapJobDTOToDomain } from '@skillsier/types';
import { QUERY_KEYS } from '../config/query-client';

const jobsClient = new MockJobsClient();

export const usePublishJob = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (jobId: string) => {
      const dto = await jobsClient.publishJob(jobId);
      return mapJobDTOToDomain(dto);
    },
    onSuccess: (_, jobId) => {
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.job(jobId) });
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.jobs });
    },
  });
};