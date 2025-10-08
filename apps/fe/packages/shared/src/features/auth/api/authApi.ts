import apiClient from '../../../lib/api/client';
import { API_ENDPOINTS } from '../../../constants/api';
import type { AuthResponse, User, RegisterRequest, LoginRequest } from '@skillsier/types';

export const authApi = {
  login: async (credentials: LoginRequest): Promise<AuthResponse> => {
    return apiClient.post(API_ENDPOINTS.AUTH.LOGIN, credentials);
  },

  register: async (data: RegisterRequest): Promise<AuthResponse> => {
    return apiClient.post(API_ENDPOINTS.AUTH.REGISTER, data);
  },

  logout: async (): Promise<void> => {
    return apiClient.post(API_ENDPOINTS.AUTH.LOGOUT);
  },

  getCurrentUser: async (): Promise<User> => {
    return apiClient.get(API_ENDPOINTS.AUTH.ME);
  },

  refreshToken: async (refreshToken: string): Promise<AuthResponse> => {
    return apiClient.post(API_ENDPOINTS.AUTH.REFRESH, { refreshToken });
  },
};