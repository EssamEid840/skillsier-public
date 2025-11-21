import * as React from 'react';
import type {
  AuthState,
  AuthUser,
  LoginCredentials,
  SignupCredentials,
} from '../types/auth.types';
import type { AuthAdapter } from '../adapters/AuthAdapter';

interface AuthContextValue extends AuthState {
  login: (credentials: LoginCredentials) => Promise<void>;
  signup: (credentials: SignupCredentials) => Promise<void>;
  logout: () => Promise<void>;
  refreshAuth: () => Promise<void>;
}

export const AuthContext = React.createContext<AuthContextValue | null>(null);

interface AuthProviderProps {
  adapter: AuthAdapter;
  children: React.ReactNode;
}

export const AuthProvider: React.FC<AuthProviderProps> = ({
  adapter,
  children,
}) => {
  const [state, setState] = React.useState<AuthState>({
    user: null,
    isAuthenticated: false,
    isLoading: true,
    error: null,
  });

  const login = React.useCallback(
    async (credentials: LoginCredentials) => {
      try {
        setState(prev => ({ ...prev, isLoading: true, error: null }));
        const session = await adapter.login(credentials);
        setState({
          user: session.user,
          isAuthenticated: true,
          isLoading: false,
          error: null,
        });
      } catch (error) {
        setState(prev => ({
          ...prev,
          isLoading: false,
          error: error instanceof Error ? error.message : 'Login failed',
        }));
        throw error;
      }
    },
    [adapter]
  );

  const signup = React.useCallback(
    async (credentials: SignupCredentials) => {
      try {
        setState(prev => ({ ...prev, isLoading: true, error: null }));
        const session = await adapter.signup(credentials);
        setState({
          user: session.user,
          isAuthenticated: true,
          isLoading: false,
          error: null,
        });
      } catch (error) {
        setState(prev => ({
          ...prev,
          isLoading: false,
          error: error instanceof Error ? error.message : 'Signup failed',
        }));
        throw error;
      }
    },
    [adapter]
  );

  const logout = React.useCallback(async () => {
    try {
      await adapter.logout();
      setState({
        user: null,
        isAuthenticated: false,
        isLoading: false,
        error: null,
      });
    } catch (error) {
      setState(prev => ({
        ...prev,
        error: error instanceof Error ? error.message : 'Logout failed',
      }));
      throw error;
    }
  }, [adapter]);

  const refreshAuth = React.useCallback(async () => {
    try {
      setState(prev => ({ ...prev, isLoading: true }));
      const user = await adapter.getCurrentUser();
      setState({
        user,
        isAuthenticated: !!user,
        isLoading: false,
        error: null,
      });
    } catch (error) {
      setState({
        user: null,
        isAuthenticated: false,
        isLoading: false,
        error: null,
      });
    }
  }, [adapter]);

  React.useEffect(() => {
    refreshAuth();
  }, [refreshAuth]);

  const value = React.useMemo(
    () => ({
      ...state,
      login,
      signup,
      logout,
      refreshAuth,
    }),
    [state, login, signup, logout, refreshAuth]
  );

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};