export const runtime = "nodejs";
import { NextRequest, NextResponse } from "next/server";

export async function POST(req: NextRequest) {
  const upstream = process.env.USERS_GRAPHQL_URL;
  if (!upstream) return NextResponse.json({ error: "USERS_GRAPHQL_URL missing" }, { status: 500 });

  const cookieName = process.env.AUTH_COOKIE_NAME || "kc_session";
  const at = req.cookies.get(cookieName)?.value;

  const body = await req.text();
  const res = await fetch(upstream, {
    method: "POST",
    headers: {
      "content-type": "application/json",
      ...(at ? { Authorization: `Bearer ${at}` } : {}),
    },
    body,
  });

  return new NextResponse(await res.text(), {
    status: res.status,
    headers: { "content-type": res.headers.get("content-type") || "application/json" },
  });
}
