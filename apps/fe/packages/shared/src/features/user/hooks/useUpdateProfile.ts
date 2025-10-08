import { useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import { useAuthStore } from '../../auth/stores/authStore';
import type { UpdateProfileRequest } from '@skillsier/types';

export const useUpdateProfile = () => {
  const queryClient = useQueryClient();
  const setUser = useAuthStore((state) => state.setUser);

  return useMutation({
    mutationFn: (data: UpdateProfileRequest) => userApi.updateProfile(data),
    onSuccess: (data) => {
      queryClient.setQueryData(queryKeys.users.profile, data);
      queryClient.invalidateQueries({ queryKey: queryKeys.auth.me });
      setUser(data);
    },
  });
};
