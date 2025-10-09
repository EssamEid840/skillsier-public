import { useQuery } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type { User } from '@skillsier/types';

export function useUserProfile() {
  return useQuery<User>({
    queryKey: queryKeys.users.profile,
    queryFn: userApi.getProfile,
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
}