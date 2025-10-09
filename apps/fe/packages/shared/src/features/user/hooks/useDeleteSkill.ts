import { useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';

export function useDeleteSkill() {
  const queryClient = useQueryClient();

  return useMutation<void, Error, string>({
    mutationFn: userApi.deleteSkill,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.skills });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.freelancerProfile });
    },
  });
}