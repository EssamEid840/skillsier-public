import { useQuery } from '@tanstack/react-query';
import { useAuthStore } from '../stores/authStore';
import { authApi } from '../api/authApi';
import { queryKeys } from '../../../lib/api/queryClient';

export const useAuth = () => {
  const { user, isAuthenticated, setUser, logout: logoutStore } = useAuthStore();

  const { data, isLoading, error } = useQuery({
    queryKey: queryKeys.auth.me,
    queryFn: authApi.getCurrentUser,
    enabled: isAuthenticated,
    retry: false,
    staleTime: Infinity,
  });

  return {
    user: data || user,
    isAuthenticated,
    isLoading,
    error,
    logout: logoutStore,
  };
};