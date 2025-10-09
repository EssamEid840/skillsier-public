import { useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';

export function useDeleteEducation() {
  const queryClient = useQueryClient();

  return useMutation<void, Error, string>({
    mutationFn: userApi.deleteEducation,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.education });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.freelancerProfile });
    },
  });
}