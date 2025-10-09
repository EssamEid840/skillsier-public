import { useQuery } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type { WorkExperience } from '@skillsier/types';

export function useWorkExperience() {
  return useQuery<WorkExperience[]>({
    queryKey: queryKeys.users.workExperience,
    queryFn: userApi.getWorkExperience,
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
}