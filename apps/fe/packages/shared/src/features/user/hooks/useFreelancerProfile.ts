import { useQuery } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type { FreelancerProfile } from '@skillsier/types';

export function useFreelancerProfile() {
  return useQuery<FreelancerProfile>({
    queryKey: queryKeys.users.freelancerProfile,
    queryFn: userApi.getFreelancerProfile,
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
}