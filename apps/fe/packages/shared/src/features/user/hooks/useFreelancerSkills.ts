import { useQuery } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type { FreelancerSkill } from '@skillsier/types';

export function useFreelancerSkills() {
  return useQuery<FreelancerSkill[]>({
    queryKey: queryKeys.users.skills,
    queryFn: userApi.getSkills,
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
}