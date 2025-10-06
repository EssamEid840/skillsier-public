export const runtime = "nodejs";
import { NextRequest, NextResponse } from "next/server";
import { reqEnv } from "../_utils/env";
import { setSessionCookies } from "../_utils/cookies";

export async function POST(req: NextRequest) {
  const cookieName = process.env.AUTH_COOKIE_NAME || "kc_session";
  const rt = req.cookies.get(`${cookieName}_rt`)?.value;
  if (!rt) return NextResponse.json({ error: "no_refresh_token" }, { status: 401 });

  const issuer = reqEnv("KEYCLOAK_ISSUER_URL");
  const clientId = reqEnv("KEYCLOAK_CLIENT_ID");
  const clientSecret = process.env.KEYCLOAK_CLIENT_SECRET || "";
  const body = new URLSearchParams({
    grant_type: "refresh_token",
    refresh_token: rt,
    client_id: clientId
  });
  if (clientSecret) body.set("client_secret", clientSecret);

  const r = await fetch(`${issuer}/protocol/openid-connect/token`, {
    method: "POST",
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
    body
  });

  if (!r.ok) return NextResponse.json({ error: "refresh_failed" }, { status: 401 });

  const tokens = await r.json();
  const res = NextResponse.json({ ok: true });
  setSessionCookies(res, tokens);
  return res;
}
