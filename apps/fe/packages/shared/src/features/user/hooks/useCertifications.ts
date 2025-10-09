import { useQuery } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type { Certification } from '@skillsier/types';

export function useCertifications() {
  return useQuery<Certification[]>({
    queryKey: queryKeys.users.certifications,
    queryFn: userApi.getCertifications,
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
}