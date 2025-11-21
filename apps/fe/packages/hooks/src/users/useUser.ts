import { useQuery } from '@tanstack/react-query';
import { UsersClient } from '@skillsier/api';
import { mapUserDTOToDomain } from '@skillsier/types';
import { QUERY_KEYS } from '../config/query-client';

const usersClient = new UsersClient({
  baseURL: process.env.API_BASE_URL || 'http://localhost:8080',
});

export const useUser = (userId: string, enabled: boolean = true) => {
  return useQuery({
    queryKey: QUERY_KEYS.user(userId),
    queryFn: async () => {
      const dto = await usersClient.getUser(userId);
      return mapUserDTOToDomain(dto);
    },
    enabled: enabled && !!userId,
  });
};

export const useCurrentUser = () => {
  return useQuery({
    queryKey: QUERY_KEYS.currentUser,
    queryFn: async () => {
      const dto = await usersClient.getCurrentUser();
      return mapUserDTOToDomain(dto);
    },
  });
};