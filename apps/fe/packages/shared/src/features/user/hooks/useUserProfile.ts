import { useQuery } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';

export const useUserProfile = () => {
  return useQuery({
    queryKey: queryKeys.users.profile,
    queryFn: userApi.getProfile,
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
};