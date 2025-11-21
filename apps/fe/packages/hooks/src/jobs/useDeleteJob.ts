import { useMutation, useQueryClient } from '@tanstack/react-query';
import { MockJobsClient } from '@skillsier/api';
import { QUERY_KEYS } from '../config/query-client';

const jobsClient = new MockJobsClient();

export const useDeleteJob = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (jobId: string) => {
      await jobsClient.deleteJob(jobId);
    },
    onSuccess: (_, jobId) => {
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.jobs });
      queryClient.removeQueries({ queryKey: QUERY_KEYS.job(jobId) });
    },
  });
};