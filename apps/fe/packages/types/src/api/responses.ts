// ============================================
// FILE: packages/types/src/api/responses.ts

import { User } from "../entities/user";

// ============================================
export interface ApiResponse<T> {
  success: boolean;
  data: T;
  message?: string;
}

export interface ApiError {
  success: false;
  error: {
    code: string;
    message: string;
    details?: Record<string, string[]>;
  };
}

export interface AuthResponse {
  accessToken: string;
  refreshToken: string;
  expiresIn: number;
  tokenType: string;
  user: User;
}

export interface PaginatedResponse<T> {
  items: T[];
  pagination: PaginationMeta;
}

export interface PaginationMeta {
  page: number;
  pageSize: number;
  totalItems: number;
  totalPages: number;
  hasNext: boolean;
  hasPrevious: boolean;
}