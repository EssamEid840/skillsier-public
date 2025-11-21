import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { UsersClient } from '@skillsier/api';
import { mapUserProfileDTOToDomain } from '@skillsier/types';
import type { UserProfileDTO } from '@skillsier/types';
import { QUERY_KEYS } from '../config/query-client';

const usersClient = new UsersClient({
  baseURL: process.env.API_BASE_URL || 'http://localhost:8080',
});

export const useUserProfile = (userId: string, enabled: boolean = true) => {
  return useQuery({
    queryKey: QUERY_KEYS.userProfile(userId),
    queryFn: async () => {
      const dto = await usersClient.getUserProfile(userId);
      return mapUserProfileDTOToDomain(dto);
    },
    enabled: enabled && !!userId,
  });
};

export interface UpdateUserProfileParams {
  userId: string;
  data: Partial<UserProfileDTO>;
}

export const useUpdateUserProfile = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async ({ userId, data }: UpdateUserProfileParams) => {
      const dto = await usersClient.updateUserProfile(userId, data);
      return mapUserProfileDTOToDomain(dto);
    },
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.userProfile(variables.userId),
      });
    },
  });
};