import { createRemoteJWKSet, jwtVerify, JWTPayload } from "jose";

let jwks: ReturnType<typeof createRemoteJWKSet> | null = null;

export function getRemoteJWKS() {
  if (jwks) return jwks;
  const issuer = process.env.KEYCLOAK_ISSUER_URL!;
  jwks = createRemoteJWKSet(new URL(`${issuer}/protocol/openid-connect/certs`));
  return jwks;
}

export async function verifyJwt(token: string) {
  const issuer = process.env.KEYCLOAK_ISSUER_URL!;
  const { payload } = await jwtVerify(token, getRemoteJWKS(), { issuer });
  return payload as JWTPayload & { email?: string; preferred_username?: string; name?: string };
}
