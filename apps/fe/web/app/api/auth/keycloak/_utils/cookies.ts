import { NextResponse } from "next/server";
import { optEnv } from "./env";

export function setSessionCookies(res: NextResponse, tokens: {
  access_token: string;
  refresh_token?: string;
  expires_in?: number;
}) {
  const cookieName = optEnv("AUTH_COOKIE_NAME", "kc_session");
  const secure = optEnv("AUTH_COOKIE_SECURE", process.env.NODE_ENV === "production" ? "true" : "false") === "true";
  const sameSite = (optEnv("AUTH_COOKIE_SAMESITE", "lax") as "lax" | "strict" | "none");

  const maxAge = Math.max(60, Math.min(60 * 60 * 8, (tokens.expires_in || 300)));

  res.cookies.set(cookieName, tokens.access_token, {
    httpOnly: true,
    sameSite,
    secure,
    path: "/",
    maxAge,
  });

  if (tokens.refresh_token) {
    res.cookies.set(`${cookieName}_rt`, tokens.refresh_token, {
      httpOnly: true,
      sameSite,
      secure,
      path: "/",
      maxAge: 60 * 60 * 24 * 30, // 30 days
    });
  }
}

export function clearSessionCookies(res: NextResponse) {
  const cookieName = optEnv("AUTH_COOKIE_NAME", "kc_session");
  res.cookies.set(cookieName, "", { httpOnly: true, path: "/", maxAge: 0 });
  res.cookies.set(`${cookieName}_rt`, "", { httpOnly: true, path: "/", maxAge: 0 });
}
