import { describe, it, expect, beforeEach } from 'vitest';
import { DevAuthAdapter } from '../src/adapters/dev-adapter';
import { UserRole } from '@skillsier/types';

describe('DevAuthAdapter', () => {
  let adapter: DevAuthAdapter;

  beforeEach(() => {
    adapter = new DevAuthAdapter();
  });

  it('logs in with valid credentials', async () => {
    const session = await adapter.login({
      email: 'freelancer@skillsier.dev',
      password: 'freelancer123',
    });

    expect(session.user).toBeDefined();
    expect(session.user.email).toBe('freelancer@skillsier.dev');
    expect(session.user.role).toBe(UserRole.FREELANCER);
    expect(session.tokens.accessToken).toBeDefined();
  });

  it('throws error with invalid credentials', async () => {
    await expect(
      adapter.login({
        email: 'wrong@email.com',
        password: 'wrongpass',
      })
    ).rejects.toThrow('Invalid credentials');
  });

  it('gets current user after login', async () => {
    await adapter.login({
      email: 'client@skillsier.dev',
      password: 'client123',
    });

    const user = await adapter.getCurrentUser();
    expect(user).toBeDefined();
    expect(user?.email).toBe('client@skillsier.dev');
    expect(user?.role).toBe(UserRole.CLIENT);
  });

  it('clears user after logout', async () => {
    await adapter.login({
      email: 'admin@skillsier.dev',
      password: 'admin123',
    });

    await adapter.logout();

    const user = await adapter.getCurrentUser();
    expect(user).toBeNull();
  });

  it('creates new user on signup', async () => {
    const session = await adapter.signup({
      email: 'newuser@test.com',
      password: 'password123',
      firstName: 'New',
      lastName: 'User',
      role: UserRole.FREELANCER,
    });

    expect(session.user.email).toBe('newuser@test.com');
    expect(session.user.firstName).toBe('New');
  });
});