import { SignJWT, jwtVerify } from 'jose';
import { UserRole } from '@skillsier/types';
import type { AuthAdapter } from './AuthAdapter';
import type {
  AuthUser,
  AuthTokens,
  LoginCredentials,
  SignupCredentials,
  AuthSession,
} from '../types/auth.types';

interface DevUser {
  id: string;
  email: string;
  password: string;
  firstName: string;
  lastName: string;
  role: UserRole;
  profilePicture?: string;
}

const DEV_USERS: DevUser[] = [
  {
    id: 'dev-admin-1',
    email: 'admin@skillsier.dev',
    password: 'admin123',
    firstName: 'Admin',
    lastName: 'User',
    role: UserRole.ADMIN,
  },
  {
    id: 'dev-client-1',
    email: 'client@skillsier.dev',
    password: 'client123',
    firstName: 'Client',
    lastName: 'User',
    role: UserRole.CLIENT,
  },
  {
    id: 'dev-freelancer-1',
    email: 'freelancer@skillsier.dev',
    password: 'freelancer123',
    firstName: 'Freelancer',
    lastName: 'User',
    role: UserRole.FREELANCER,
    profilePicture: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Freelancer',
  },
];

const JWT_SECRET = new TextEncoder().encode(
  process.env.NEXTAUTH_SECRET || 'dev-secret-change-in-production-min-32-chars'
);

export class DevAuthAdapter implements AuthAdapter {
  private currentUser: AuthUser | null = null;
  private currentTokens: AuthTokens | null = null;

  async login(credentials: LoginCredentials): Promise<AuthSession> {
    await new Promise(resolve => setTimeout(resolve, 500));

    const user = DEV_USERS.find(
      u => u.email === credentials.email && u.password === credentials.password
    );

    if (!user) {
      throw new Error('Invalid credentials');
    }

    const authUser: AuthUser = {
      id: user.id,
      email: user.email,
      firstName: user.firstName,
      lastName: user.lastName,
      role: user.role,
      profilePicture: user.profilePicture,
    };

    const tokens = await this.generateTokens(authUser);
    const expiresAt = new Date(Date.now() + tokens.expiresIn * 1000);

    this.currentUser = authUser;
    this.currentTokens = tokens;

    if (typeof window !== 'undefined') {
      localStorage.setItem('dev_auth_user', JSON.stringify(authUser));
      localStorage.setItem('dev_auth_tokens', JSON.stringify(tokens));
    }

    return {
      user: authUser,
      tokens,
      expiresAt,
    };
  }

  async signup(credentials: SignupCredentials): Promise<AuthSession> {
    await new Promise(resolve => setTimeout(resolve, 500));

    const existingUser = DEV_USERS.find(u => u.email === credentials.email);
    if (existingUser) {
      throw new Error('Email already exists');
    }

    const newUser: DevUser = {
      id: `dev-user-${Date.now()}`,
      email: credentials.email,
      password: credentials.password,
      firstName: credentials.firstName,
      lastName: credentials.lastName,
      role: credentials.role,
    };

    DEV_USERS.push(newUser);

    return this.login({
      email: credentials.email,
      password: credentials.password,
    });
  }

  async logout(): Promise<void> {
    await new Promise(resolve => setTimeout(resolve, 200));

    this.currentUser = null;
    this.currentTokens = null;

    if (typeof window !== 'undefined') {
      localStorage.removeItem('dev_auth_user');
      localStorage.removeItem('dev_auth_tokens');
    }
  }

  async refreshToken(refreshToken: string): Promise<AuthTokens> {
    await new Promise(resolve => setTimeout(resolve, 300));

    if (!this.currentUser) {
      throw new Error('No user logged in');
    }

    return this.generateTokens(this.currentUser);
  }

  async getCurrentUser(): Promise<AuthUser | null> {
    if (this.currentUser) {
      return this.currentUser;
    }

    if (typeof window !== 'undefined') {
      const storedUser = localStorage.getItem('dev_auth_user');
      if (storedUser) {
        this.currentUser = JSON.parse(storedUser);
        return this.currentUser;
      }
    }

    return null;
  }

  async verifyToken(token: string): Promise<boolean> {
    try {
      await jwtVerify(token, JWT_SECRET);
      return true;
    } catch {
      return false;
    }
  }

  async resetPassword(email: string): Promise<void> {
    await new Promise(resolve => setTimeout(resolve, 500));
    console.log(`Password reset email sent to ${email}`);
  }

  async changePassword(
    oldPassword: string,
    newPassword: string
  ): Promise<void> {
    await new Promise(resolve => setTimeout(resolve, 500));
    
    if (!this.currentUser) {
      throw new Error('Not authenticated');
    }

    const user = DEV_USERS.find(u => u.id === this.currentUser!.id);
    if (!user || user.password !== oldPassword) {
      throw new Error('Invalid old password');
    }

    user.password = newPassword;
  }

  private async generateTokens(user: AuthUser): Promise<AuthTokens> {
    const accessToken = await new SignJWT({
      sub: user.id,
      email: user.email,
      role: user.role,
    })
      .setProtectedHeader({ alg: 'HS256' })
      .setIssuedAt()
      .setExpirationTime('1h')
      .sign(JWT_SECRET);

    const refreshToken = await new SignJWT({ sub: user.id })
      .setProtectedHeader({ alg: 'HS256' })
      .setIssuedAt()
      .setExpirationTime('7d')
      .sign(JWT_SECRET);

    return {
      accessToken,
      refreshToken,
      expiresIn: 3600,
    };
  }
}