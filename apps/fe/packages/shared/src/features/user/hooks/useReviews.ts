import { useQuery } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';

export interface Review {
  id: string;
  rating: number;
  comment: string;
  reviewerName: string;
  reviewerAvatar?: string;
  jobTitle: string;
  date: string;
}

export function useReviews() {
  return useQuery<Review[]>({
    queryKey: queryKeys.users.reviews,
    queryFn: userApi.getReviews,
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
}