export const runtime = "nodejs";

import { NextRequest, NextResponse } from "next/server";
import { reqEnv } from "../_utils/env";
import { setSessionCookies } from "../_utils/cookies";

export async function POST(req: NextRequest) {
  const { email, password } = await req.json().catch(() => ({}));
  if (!email || !password) {
    return NextResponse.json({ error: "Missing email/password" }, { status: 400 });
  }

  const issuer = reqEnv("KEYCLOAK_ISSUER_URL");
  const clientId = reqEnv("KEYCLOAK_CLIENT_ID");
  const clientSecret = process.env.KEYCLOAK_CLIENT_SECRET || "";

  const body = new URLSearchParams({
    grant_type: "password",
    username: email,
    password,
    client_id: clientId,
    scope: "openid email profile",
  });
  if (clientSecret) body.set("client_secret", clientSecret);

  const tokenRes = await fetch(`${issuer}/protocol/openid-connect/token`, {
    method: "POST",
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
    body
  });

  if (!tokenRes.ok) {
    const text = await tokenRes.text().catch(() => "");
    return NextResponse.json({ error: "invalid_credentials", details: text }, { status: 401 });
  }

  const tokens = await tokenRes.json();
  const res = NextResponse.json({ ok: true });
  setSessionCookies(res, tokens);
  return res;
}
