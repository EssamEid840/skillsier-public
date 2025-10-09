import { useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';

export function useDeleteAvatar() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: userApi.deleteAvatar,
    onSuccess: () => {
      // Invalidate profile queries to refetch updated data
      queryClient.invalidateQueries({ queryKey: queryKeys.users.profile });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.freelancerProfile });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.clientProfile });
    },
  });
}