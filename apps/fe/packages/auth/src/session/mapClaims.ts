export type IdTokenClaims = {
  sub: string;
  email?: string;
  preferred_username?: string;
  name?: string;
  picture?: string;
  email_verified?: boolean;
  [k: string]: unknown;
};

export function mapToUser(claims: IdTokenClaims) {
  return {
    id: claims.sub,
    email: claims.email,
    displayName: claims.name ?? claims.preferred_username,
    picture: claims.picture,
    emailVerified: !!claims.email_verified
  };
}
