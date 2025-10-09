import { useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';

export function useDeleteCertification() {
  const queryClient = useQueryClient();

  return useMutation<void, Error, string>({
    mutationFn: userApi.deleteCertification,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.certifications });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.freelancerProfile });
    },
  });
}