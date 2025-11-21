import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import type { AuthUser } from '@skillsier/types';

interface AuthState {
  // State
  user: AuthUser | null;
  isAuthenticated: boolean;
  token: string | null;
  refreshToken: string | null;

  // Actions
  setUser: (user: AuthUser | null) => void;
  setTokens: (token: string, refreshToken?: string) => void;
  clearAuth: () => void;
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      // Initial state
      user: null,
      isAuthenticated: false,
      token: null,
      refreshToken: null,

      // Actions
      setUser: (user) =>
        set({
          user,
          isAuthenticated: !!user,
        }),

      setTokens: (token, refreshToken) =>
        set({
          token,
          refreshToken: refreshToken || null,
        }),

      clearAuth: () =>
        set({
          user: null,
          isAuthenticated: false,
          token: null,
          refreshToken: null,
        }),
    }),
    {
      name: 'skillsier-auth-storage',
    }
  )
);