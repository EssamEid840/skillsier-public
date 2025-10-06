export const runtime = "nodejs";
import { NextRequest, NextResponse } from "next/server";

const UPSTREAM = process.env.USERS_BE_URL!;

function authHeaderFromCookie(req: NextRequest) {
  const cookieName = process.env.AUTH_COOKIE_NAME || "kc_session";
  const at = req.cookies.get(cookieName)?.value;
  return at ? { Authorization: `Bearer ${at}` } : {};
}

async function proxy(req: NextRequest, path: string[]) {
  if (!UPSTREAM) return NextResponse.json({ error: "USERS_BE_URL missing" }, { status: 500 });

  const url = new URL(path.join("/"), UPSTREAM.endsWith("/") ? UPSTREAM : `${UPSTREAM}/`);
  // Preserve query string
  const q = new URL(req.url).search;
  const target = `${url.toString()}${q || ""}`;

  const headers: Record<string, string> = {
    ...authHeaderFromCookie(req),
  };
  // forward content-type if present
  const ct = req.headers.get("content-type");
  if (ct) headers["content-type"] = ct;

  let body: BodyInit | undefined = undefined;
  if (req.method !== "GET" && req.method !== "HEAD") {
    // for JSON / text / x-www-form-urlencoded this is fine
    const buf = await req.arrayBuffer();
    body = Buffer.from(buf);
  }

  const upstream = await fetch(target, {
    method: req.method,
    headers,
    body,
  });

  const res = new NextResponse(upstream.body, {
    status: upstream.status,
  });

  // forward content-type
  const uct = upstream.headers.get("content-type");
  if (uct) res.headers.set("content-type", uct);

  return res;
}

export async function GET(req: NextRequest, ctx: { params: { path: string[] } }) {
  return proxy(req, ctx.params.path);
}
export async function POST(req: NextRequest, ctx: { params: { path: string[] } }) {
  return proxy(req, ctx.params.path);
}
export async function PUT(req: NextRequest, ctx: { params: { path: string[] } }) {
  return proxy(req, ctx.params.path);
}
export async function PATCH(req: NextRequest, ctx: { params: { path: string[] } }) {
  return proxy(req, ctx.params.path);
}
export async function DELETE(req: NextRequest, ctx: { params: { path: string[] } }) {
  return proxy(req, ctx.params.path);
}
