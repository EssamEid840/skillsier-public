import { useQuery } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type { Education } from '@skillsier/types';

export function useEducation() {
  return useQuery<Education[]>({
    queryKey: queryKeys.users.education,
    queryFn: userApi.getEducation,
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
}