import { DevAuthAdapter } from './dev-adapter';
import { KeycloakAuthAdapter } from './keycloak-adapter';
import type { AuthAdapter } from './AuthAdapter';

export type AuthProvider = 'dev' | 'keycloak';

export const createAuthAdapter = (
  provider: AuthProvider = 'dev'
): AuthAdapter => {
  switch (provider) {
    case 'keycloak':
      return new KeycloakAuthAdapter();
    case 'dev':
    default:
      return new DevAuthAdapter();
  }
};

export { AuthAdapter, DevAuthAdapter, KeycloakAuthAdapter };