import { headers, cookies } from "next/headers";
import { NextResponse } from "next/server";

type StartState = { ru?: string };

// Edge-safe base64url decoder
function b64urlDecodeToString(b64url: string): string {
  const b64 = b64url.replace(/-/g, "+").replace(/_/g, "/") + "===".slice((b64url.length + 3) % 4);
  const bin = atob(b64);
  const bytes = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; i++) bytes[i] = bin.charCodeAt(i);
  return new TextDecoder().decode(bytes);
}
function decodeState(raw?: string | null): StartState {
  if (!raw) return {};
  try {
    return JSON.parse(b64urlDecodeToString(raw));
  } catch {
    return {};
  }
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

/** Make a same-origin absolute URL (Edge requires absolute URLs) */
function toSameOriginAbsolute(target: string, base: string): string {
  try {
    const abs = new URL(target, base);       // resolves relative or absolute
    const b   = new URL(base);
    // Enforce same-origin to avoid open redirect
    if (abs.origin !== b.origin) {
      return b.origin + abs.pathname + abs.search + abs.hash;
    }
    return abs.toString();
  } catch {
    return base + "/"; // fallback
  }
}

function setCookie(res: NextResponse, name: string, value: string, maxAgeSec: number, secure: boolean) {
  for (const sec of [secure, true, false]) {
    res.cookies.set({
      name,
      value,
      path: "/",
      httpOnly: true,
      sameSite: "lax",
      secure: sec,
      maxAge: Math.max(0, maxAgeSec),
    });
  }
}

export async function GET(req: Request) {
  try {
    const issuer       = process.env.KEYCLOAK_ISSUER_URL!;
    const clientId     = process.env.KEYCLOAK_CLIENT_ID!;
    const clientSecret = process.env.KEYCLOAK_CLIENT_SECRET!;
    if (!issuer || !clientId || !clientSecret) {
      return new NextResponse(
        JSON.stringify({ error: "config", details: "Missing KEYCLOAK envs" }),
        { status: 500, headers: { "Content-Type": "application/json" } }
      );
    }

    const url   = new URL(req.url);
    const code  = url.searchParams.get("code");
    const state = url.searchParams.get("state");
    if (!code) {
      return new NextResponse(
        JSON.stringify({ error: "invalid_request", details: "Missing code" }),
        { status: 400, headers: { "Content-Type": "application/json" } }
      );
    }

    const redirectUri = await resolveRedirectUri(req);
    const form = new URLSearchParams({
      grant_type: "authorization_code",
      client_id: clientId,
      client_secret: clientSecret,
      code,
      redirect_uri: redirectUri,
    });

    const tokenUrl = `${issuer}/protocol/openid-connect/token`;
    const tr = await fetch(tokenUrl, {
      method: "POST",
      headers: { "Content-Type": "application/x-www-form-urlencoded" },
      body: form,
      cache: "no-store",
    });

    if (!tr.ok) {
      const details = await tr.text().catch(() => "");
      return new NextResponse(
        JSON.stringify({ error: "token_exchange_failed", details }),
        { status: 502, headers: { "Content-Type": "application/json" } }
      );
    }

    const tj = await tr.json() as {
      access_token: string;
      refresh_token?: string;
      id_token?: string;
      expires_in: number;
    };

    const h = await headers();
    const secure = (h.get("x-forwarded-proto") ?? new URL(req.url).protocol.replace(":", "")) === "https";

    const base = await resolveBase(req);
    const { ru } = decodeState(state);
    const fallback = process.env.AUTH_AFTER_LOGIN_REDIRECT_URI ?? "/";
    const redirectToAbs = toSameOriginAbsolute(ru ?? fallback, base);

    const res = NextResponse.redirect(redirectToAbs, { status: 302 });

    // Persist tokens (names match earlier routes)
    setCookie(res, "kc_access_token", tj.access_token, tj.expires_in ?? 300, secure);
    if (tj.refresh_token) setCookie(res, "kc_refresh_token", tj.refresh_token, 30 * 24 * 3600, secure);
    if (tj.id_token)      setCookie(res, "kc_id_token", tj.id_token, tj.expires_in ?? 300, secure);
    setCookie(res, "kc_session", "1", tj.expires_in ?? 300, secure);

    return res;
  } catch (e: any) {
    return new NextResponse(
      JSON.stringify({ error: "callback_failed", details: String(e?.message ?? e) }),
      { status: 500, headers: { "Content-Type": "application/json" } }
    );
  }
}
