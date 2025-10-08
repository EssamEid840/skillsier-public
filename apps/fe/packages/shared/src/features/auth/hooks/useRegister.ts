import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useAuthStore } from '../stores/authStore';
import { authApi } from '../api/authApi';
import { setAccessToken, setRefreshToken } from '../../../lib/api/client';
import { queryKeys } from '../../../lib/api/queryClient';
import type { RegisterRequest } from '@skillsier/types';

export const useRegister = () => {
  const queryClient = useQueryClient();
  const setUser = useAuthStore((state) => state.setUser);

  return useMutation({
    mutationFn: (data: RegisterRequest) => authApi.register(data),
    onSuccess: (data) => {
      setAccessToken(data.accessToken);
      setRefreshToken(data.refreshToken);
      setUser(data.user);
      queryClient.setQueryData(queryKeys.auth.me, data.user);
    },
  });
};