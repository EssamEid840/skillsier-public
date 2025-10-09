import { useQuery } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';

export interface UserStats {
  totalEarnings: number;
  completedJobs: number;
  activeJobs: number;
  rating: number;
  successRate: number;
  responseTime: number; // in hours
  profileViews: number;
}

export function useUserStats() {
  return useQuery<UserStats>({
    queryKey: queryKeys.users.stats,
    queryFn: userApi.getUserStats,
    staleTime: 2 * 60 * 1000, // 2 minutes
  });
}