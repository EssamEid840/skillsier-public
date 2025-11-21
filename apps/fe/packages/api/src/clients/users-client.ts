import { BaseApiClient, type ApiClientConfig } from '../lib/base-client';
import type {
  UserDTO,
  UserProfileDTO,
  CreateUserRequest,
  UpdateUserRequest,
} from '@skillsier/types';

export class UsersClient extends BaseApiClient {
  constructor(config: ApiClientConfig) {
    super(config);
  }

  async getCurrentUser(): Promise<UserDTO> {
    return this.get<UserDTO>('/v1/users/me');
  }

  async getUser(userId: string): Promise<UserDTO> {
    return this.get<UserDTO>(`/v1/users/${userId}`);
  }

  async updateUser(userId: string, data: UpdateUserRequest): Promise<UserDTO> {
    return this.patch<UserDTO>(`/v1/users/${userId}`, data);
  }

  async getUserProfile(userId: string): Promise<UserProfileDTO> {
    return this.get<UserProfileDTO>(`/v1/users/${userId}/profile`);
  }

  async updateUserProfile(
    userId: string,
    data: Partial<UserProfileDTO>
  ): Promise<UserProfileDTO> {
    return this.patch<UserProfileDTO>(`/v1/users/${userId}/profile`, data);
  }

  async verifyEmail(userId: string, token: string): Promise<void> {
    return this.post<void>(`/v1/users/${userId}/verify-email`, { token });
  }

  async resendVerification(userId: string): Promise<void> {
    return this.post<void>(`/v1/users/${userId}/resend-verification`);
  }
}