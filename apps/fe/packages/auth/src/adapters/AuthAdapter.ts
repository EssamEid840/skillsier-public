import type {
  AuthUser,
  AuthTokens,
  LoginCredentials,
  SignupCredentials,
  AuthSession,
} from '../types/auth.types';

export interface AuthAdapter {
  login(credentials: LoginCredentials): Promise<AuthSession>;
  
  signup(credentials: SignupCredentials): Promise<AuthSession>;
  
  logout(): Promise<void>;
  
  refreshToken(refreshToken: string): Promise<AuthTokens>;
  
  getCurrentUser(): Promise<AuthUser | null>;
  
  verifyToken(token: string): Promise<boolean>;
  
  resetPassword(email: string): Promise<void>;
  
  changePassword(oldPassword: string, newPassword: string): Promise<void>;
}