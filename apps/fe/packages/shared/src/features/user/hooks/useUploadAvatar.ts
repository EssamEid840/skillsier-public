import { useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';

export function useUploadAvatar() {
  const queryClient = useQueryClient();

  return useMutation<string, Error, File | FormData>({
    mutationFn: userApi.uploadAvatar,
    onSuccess: () => {
      // Invalidate profile queries to refetch updated avatar
      queryClient.invalidateQueries({ queryKey: queryKeys.users.profile });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.freelancerProfile });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.clientProfile });
      queryClient.invalidateQueries({ queryKey: queryKeys.auth.me });
    },
  });
}