import { useQuery } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type { ClientProfile } from '@skillsier/types';

export function useClientProfile() {
  return useQuery<ClientProfile>({
    queryKey: queryKeys.users.clientProfile,
    queryFn: userApi.getClientProfile,
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
}