export const runtime = "nodejs";

import { NextRequest, NextResponse } from "next/server";
import { adminFetch } from "../_utils/admin";
import { optEnv } from "../_utils/env";
import { parseJwtPayload } from "../_utils/b64";

export async function POST(req: NextRequest) {
  const cookieName = optEnv("AUTH_COOKIE_NAME", "kc_session");
  const token = req.cookies.get(cookieName)?.value;
  if (!token) return NextResponse.json({ error: "not_authenticated" }, { status: 401 });

  const payload = parseJwtPayload<{ sub: string }>(token);
  const userId = payload?.sub;
  if (!userId) return NextResponse.json({ error: "invalid_token" }, { status: 401 });

  await adminFetch(`/users/${userId}/execute-actions-email`, {
    method: "PUT",
    body: JSON.stringify(["VERIFY_EMAIL"])
  }).catch(() => null);

  return NextResponse.json({ ok: true });
}
