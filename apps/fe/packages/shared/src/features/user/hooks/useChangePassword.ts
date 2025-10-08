import { useMutation } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import type { ChangePasswordRequest } from '@skillsier/types';

export const useChangePassword = () => {
  return useMutation({
    mutationFn: (data: ChangePasswordRequest) => userApi.changePassword(data),
  });
};
