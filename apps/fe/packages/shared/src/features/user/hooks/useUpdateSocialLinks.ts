import { useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type { UpdateSocialLinksRequest } from '@skillsier/types';

export const useUpdateSocialLinks = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: UpdateSocialLinksRequest) => userApi.updateSocialLinks(data),
    onSuccess: (data) => {
      queryClient.setQueryData(queryKeys.users.profile, data);
    },
  });
};
