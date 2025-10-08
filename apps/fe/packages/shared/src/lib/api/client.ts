import axios, { AxiosError, AxiosRequestConfig } from 'axios';
import { API_CONFIG } from '../../constants/api';
import type { ApiError, ApiResponse } from '@skillsier/types';

const apiClient = axios.create({
  baseURL: API_CONFIG.BASE_URL,
  timeout: API_CONFIG.TIMEOUT,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor
apiClient.interceptors.request.use(
  (config) => {
    const token = getAccessToken();
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Response interceptor
apiClient.interceptors.response.use(
  (response) => response.data,
  async (error: AxiosError<ApiError>) => {
    const originalRequest = error.config as AxiosRequestConfig & { _retry?: boolean };

    // Handle 401 errors (token refresh)
    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;

      try {
        const refreshToken = getRefreshToken();
        if (refreshToken) {
          const { data } = await axios.post(`${API_CONFIG.BASE_URL}/auth/refresh`, {
            refreshToken,
          });

          setAccessToken(data.accessToken);
          setRefreshToken(data.refreshToken);

          if (originalRequest.headers) {
            originalRequest.headers.Authorization = `Bearer ${data.accessToken}`;
          }

          return apiClient(originalRequest);
        }
      } catch (refreshError) {
        // Refresh failed, logout user
        clearTokens();
        if (typeof window !== 'undefined') {
          window.location.href = '/login';
        }
        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  }
);

// Token management (platform-agnostic interface)
let tokenStorage: TokenStorage = {
  getAccessToken: () => null,
  setAccessToken: () => {},
  getRefreshToken: () => null,
  setRefreshToken: () => {},
  clearTokens: () => {},
};

export interface TokenStorage {
  getAccessToken: () => string | null;
  setAccessToken: (token: string) => void;
  getRefreshToken: () => string | null;
  setRefreshToken: (token: string) => void;
  clearTokens: () => void;
}

export const setTokenStorage = (storage: TokenStorage) => {
  tokenStorage = storage;
};

export const getAccessToken = () => tokenStorage.getAccessToken();
export const setAccessToken = (token: string) => tokenStorage.setAccessToken(token);
export const getRefreshToken = () => tokenStorage.getRefreshToken();
export const setRefreshToken = (token: string) => tokenStorage.setRefreshToken(token);
export const clearTokens = () => tokenStorage.clearTokens();

export default apiClient;