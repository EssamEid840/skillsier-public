import { useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';

export const useUploadAvatar = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (file: File | Blob) => userApi.uploadAvatar(file),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.profile });
      queryClient.invalidateQueries({ queryKey: queryKeys.auth.me });
    },
  });
};

export const useDeleteAvatar = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: userApi.deleteAvatar,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.profile });
      queryClient.invalidateQueries({ queryKey: queryKeys.auth.me });
    },
  });
};