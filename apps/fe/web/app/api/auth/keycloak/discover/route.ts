export const runtime = "nodejs";
import { NextResponse } from "next/server";

export async function GET() {
  const issuer = process.env.KEYCLOAK_ISSUER_URL || "";
  const json = {
    issuer,
    authorization_endpoint: issuer ? `${issuer}/protocol/openid-connect/auth` : null,
    token_endpoint: issuer ? `${issuer}/protocol/openid-connect/token` : null,
    userinfo_endpoint: issuer ? `${issuer}/protocol/openid-connect/userinfo` : null,
    jwks_uri: issuer ? `${issuer}/protocol/openid-connect/certs` : null,
  };
  return NextResponse.json(json);
}
