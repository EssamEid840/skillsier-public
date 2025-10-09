import { useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';

export function useDeleteWorkExperience() {
  const queryClient = useQueryClient();

  return useMutation<void, Error, string>({
    mutationFn: userApi.deleteWorkExperience,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.workExperience });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.freelancerProfile });
    },
  });
}