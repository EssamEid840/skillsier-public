import { useQuery } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';

export interface Earning {
  id: string;
  amount: number;
  currency: string;
  jobId: string;
  jobTitle: string;
  clientName: string;
  date: string;
  status: 'PENDING' | 'COMPLETED' | 'WITHDRAWN';
}

export function useEarnings() {
  return useQuery<Earning[]>({
    queryKey: queryKeys.users.earnings,
    queryFn: userApi.getEarnings,
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
}