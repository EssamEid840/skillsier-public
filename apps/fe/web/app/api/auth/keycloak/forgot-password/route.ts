export const runtime = "nodejs";

import { NextRequest, NextResponse } from "next/server";
import { adminFetch } from "../_utils/admin";

export async function POST(req: NextRequest) {
  const { email } = await req.json().catch(() => ({}));
  if (!email) return NextResponse.json({ error: "Missing email" }, { status: 400 });

  // find user id by email
  const q = await adminFetch(`/users?email=${encodeURIComponent(email)}&max=1`, { method: "GET" });
  const users = await q.json();
  const userId = users?.[0]?.id as string | undefined;

  // Always respond 200 to avoid account enumeration
  if (!userId) return NextResponse.json({ ok: true });

  await adminFetch(`/users/${userId}/execute-actions-email`, {
    method: "PUT",
    body: JSON.stringify(["UPDATE_PASSWORD"])
  }).catch(() => null);

  return NextResponse.json({ ok: true });
}
