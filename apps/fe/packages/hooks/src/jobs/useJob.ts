import { useQuery } from '@tanstack/react-query';
import { MockJobsClient } from '@skillsier/api';
import { mapJobDTOToDomain } from '@skillsier/types';
import { QUERY_KEYS } from '../config/query-client';

const jobsClient = new MockJobsClient();

export const useJob = (jobId: string, enabled: boolean = true) => {
  return useQuery({
    queryKey: QUERY_KEYS.job(jobId),
    queryFn: async () => {
      const dto = await jobsClient.getJob(jobId);
      return mapJobDTOToDomain(dto);
    },
    enabled: enabled && !!jobId,
  });
};