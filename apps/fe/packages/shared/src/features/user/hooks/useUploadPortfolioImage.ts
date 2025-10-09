import { useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';

interface UploadPortfolioImageVariables {
  itemId: string;
  file: File | FormData;
}

export function useUploadPortfolioImage() {
  const queryClient = useQueryClient();

  return useMutation<string, Error, UploadPortfolioImageVariables>({
    mutationFn: ({ itemId, file }) => userApi.uploadPortfolioImage(itemId, file),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.portfolio });
    },
  });
}