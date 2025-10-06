export const runtime = "nodejs";
import { NextRequest, NextResponse } from "next/server";

export async function POST(req: NextRequest) {
  const upstream = process.env.USERS_BE_URL;
  if (!upstream) return NextResponse.json({ error: "USERS_BE_URL missing" }, { status: 500 });

  const cookieName = process.env.AUTH_COOKIE_NAME || "kc_session";
  const at = req.cookies.get(cookieName)?.value;

  const form = await req.formData(); // keeps file streams
  const res = await fetch(`${upstream.replace(/\/$/, "")}/upload`, {
    method: "POST",
    headers: { ...(at ? { Authorization: `Bearer ${at}` } : {}) },
    body: form as any,
  });

  const data = await res.text();
  return new NextResponse(data, { status: res.status, headers: { "content-type": res.headers.get("content-type") || "application/json" } });
}
