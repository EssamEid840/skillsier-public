import { headers } from "next/headers";
import { NextResponse } from "next/server";

type StartState = { ru?: string };

// base helpers (Edge-safe; no Buffer)
function b64urlEncode(str: string): string {
  // encode UTF-8 to base64url
  const bytes = new TextEncoder().encode(str);
  let bin = "";
  bytes.forEach((b) => (bin += String.fromCharCode(b)));
  const b64 = btoa(bin);
  return b64.replace(/\+/g, "-").replace(/\//g, "_").replace(/=+$/g, "");
}
function encodeState(obj: StartState) {
  return b64urlEncode(JSON.stringify(obj));
}

async function resolveBase(req: Request) {
  const h = await headers();
  const u = new URL(req.url);
  const proto = h.get("x-forwarded-proto") ?? u.protocol.replace(":", "") ?? "http";
  const host  = h.get("host") ?? u.host;
  return `${proto}://${host}`;
}
async function resolveRedirectUri(req: Request) {
  const base = await resolveBase(req);
  return process.env.AUTH_REDIRECT_URI
    ?? `${base}/api/auth/keycloak/oauth/callback/google`;
}

export async function GET(req: Request) {
  try {
    const issuer   = process.env.KEYCLOAK_ISSUER_URL!;
    const clientId = process.env.KEYCLOAK_CLIENT_ID!;
    if (!issuer || !clientId) {
      return new NextResponse(
        JSON.stringify({ error: "config", details: "Missing KEYCLOAK_ISSUER_URL or KEYCLOAK_CLIENT_ID" }),
        { status: 500, headers: { "Content-Type": "application/json" } }
      );
    }

    const url = new URL(req.url);
    // ru from query OR env fallback OR /
    const ru =
      url.searchParams.get("ru")
      ?? process.env.AUTH_AFTER_LOGIN_REDIRECT_URI
      ?? "/";

    const redirectUri = await resolveRedirectUri(req);
    const qp = new URLSearchParams({
      client_id: clientId,
      response_type: "code",
      scope: "openid email profile",
      redirect_uri: redirectUri,
      kc_idp_hint: "google",
      // Show Google chooser every time:
      prompt: "login select_account",
      // encode return URL inside state
      state: encodeState({ ru }),
    });

    return NextResponse.redirect(
      `${issuer}/protocol/openid-connect/auth?${qp.toString()}`,
      { status: 302 }
    );
  } catch (e: any) {
    return new NextResponse(
      JSON.stringify({ error: "start_failed", details: String(e?.message ?? e) }),
      { status: 500, headers: { "Content-Type": "application/json" } }
    );
  }
}
