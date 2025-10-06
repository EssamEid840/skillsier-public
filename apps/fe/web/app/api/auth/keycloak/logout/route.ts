import { cookies, headers } from "next/headers";
import { NextResponse } from "next/server";

/** Build http/https base from the current request */
function buildBase(h: Headers, req: Request) {
  const u = new URL(req.url);
  const proto = h.get("x-forwarded-proto") ?? u.protocol.replace(":", "") ?? "http";
  const host  = h.get("host") ?? u.host;
  return `${proto}://${host}`;
}

/** Delete every cookie we received (with both secure/non-secure to cover http/https) */
async function clearLocalCookies(res: NextResponse) {
  const jar = await cookies();
  for (const c of jar.getAll()) {
    for (const sec of [true, false]) {
      res.cookies.set({
        name: c.name,
        value: "",
        path: "/",
        httpOnly: true,
        sameSite: "lax",
        secure: sec,
        maxAge: 0,
      });
    }
  }
}

async function doLogout(req: Request) {
  const h = await headers();
  const base = buildBase(h, req);

  // Where to land AFTER Keycloak ends the SSO session
  const returnTo =
    process.env.AUTH_LOGOUT_REDIRECT_URI ?? `${base}/keycloak/login`;

  const issuer   = process.env.KEYCLOAK_ISSUER_URL; // e.g. https://keycloak.skillsier.com/realms/skillsier
  const clientId = process.env.KEYCLOAK_CLIENT_ID;  // e.g. skillsier-fe

  // If we don't even know the issuer, just clear local cookies and go to login.
  if (!issuer) {
    const res = NextResponse.redirect(returnTo, { status: 302 });
    await clearLocalCookies(res);
    return res;
  }

  const idToken = (await cookies()).get("kc_id_token")?.value;

  // Build the browser redirect to Keycloak logout (front-channel),
  // so Keycloak can clear its own SSO cookies on its domain.
  const qs = new URLSearchParams();
  qs.set("post_logout_redirect_uri", returnTo);
  if (idToken) {
    // Preferred: binds the exact session to terminate
    qs.set("id_token_hint", idToken);
  } else if (clientId) {
    // Fallback for realms/clients that accept client_id when id_token is missing
    qs.set("client_id", clientId);
  }

  const keycloakLogoutUrl = `${issuer}/protocol/openid-connect/logout?${qs.toString()}`;

  const res = NextResponse.redirect(keycloakLogoutUrl, { status: 302 });
  await clearLocalCookies(res);
  // Bust caches and hint the browser to clear site data (effective on HTTPS)
  res.headers.set("Cache-Control", "no-store, max-age=0");
  res.headers.set("Clear-Site-Data", '"cookies","storage"');
  return res;
}

export const GET  = doLogout;
export const POST = doLogout;
