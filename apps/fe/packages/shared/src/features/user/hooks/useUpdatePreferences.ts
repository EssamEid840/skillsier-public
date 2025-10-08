import { useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type { UpdatePreferencesRequest } from '@skillsier/types';

export const useUpdatePreferences = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: UpdatePreferencesRequest) => userApi.updatePreferences(data),
    onSuccess: (data) => {
      queryClient.setQueryData(queryKeys.users.profile, data);
    },
  });
};
