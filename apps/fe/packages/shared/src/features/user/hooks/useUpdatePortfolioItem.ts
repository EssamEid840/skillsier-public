import { useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type { UpdatePortfolioItemRequest, PortfolioItem } from '@skillsier/types';

interface UpdatePortfolioItemVariables {
  itemId: string;
  data: UpdatePortfolioItemRequest;
}

export function useUpdatePortfolioItem() {
  const queryClient = useQueryClient();

  return useMutation<PortfolioItem, Error, UpdatePortfolioItemVariables>({
    mutationFn: ({ itemId, data }) => userApi.updatePortfolioItem(itemId, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.portfolio });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.freelancerProfile });
    },
  });
}