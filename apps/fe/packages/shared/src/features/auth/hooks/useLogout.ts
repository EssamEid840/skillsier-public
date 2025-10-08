import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useAuthStore } from '../stores/authStore';
import { authApi } from '../api/authApi';
import { clearTokens } from '../../../lib/api/client';

export const useLogout = () => {
  const queryClient = useQueryClient();
  const logout = useAuthStore((state) => state.logout);

  return useMutation({
    mutationFn: authApi.logout,
    onSuccess: () => {
      clearTokens();
      logout();
      queryClient.clear();
    },
    onError: () => {
      // Even if API call fails, clear local state
      clearTokens();
      logout();
      queryClient.clear();
    },
  });
};