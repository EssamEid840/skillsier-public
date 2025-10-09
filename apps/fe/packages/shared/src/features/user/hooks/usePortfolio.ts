import { useQuery } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type { PortfolioItem } from '@skillsier/types';

export function usePortfolio() {
  return useQuery<PortfolioItem[]>({
    queryKey: queryKeys.users.portfolio,
    queryFn: userApi.getPortfolio,
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
}