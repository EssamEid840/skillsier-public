export const runtime = "nodejs";

import { NextRequest, NextResponse } from "next/server";
import { adminFetch } from "../_utils/admin";

export async function POST(req: NextRequest) {
  const { name, email, password } = await req.json().catch(() => ({}));
  if (!name || !email || !password) {
    return NextResponse.json({ error: "Missing name/email/password" }, { status: 400 });
  }

  // 1) create user
  const r = await adminFetch(`/users`, {
    method: "POST",
    body: JSON.stringify({
      username: email,
      email,
      firstName: name,
      enabled: true,
      emailVerified: false
    })
  });

  if (r.status !== 201 && r.status !== 409) { // 409 already exists
    const err = await r.text().catch(() => "");
    return NextResponse.json({ error: "create_user_failed", details: err }, { status: 500 });
  }

  // fetch ID (either newly created or existing)
  const q = await adminFetch(`/users?email=${encodeURIComponent(email)}&max=1`, { method: "GET" });
  const users = await q.json();
  const userId = users?.[0]?.id as string | undefined;
  if (!userId) return NextResponse.json({ error: "user_lookup_failed" }, { status: 500 });

  // 2) set password
  const pw = await adminFetch(`/users/${userId}/reset-password`, {
    method: "PUT",
    body: JSON.stringify({ type: "password", value: password, temporary: false })
  });
  if (!pw.ok) {
    const err = await pw.text().catch(() => "");
    return NextResponse.json({ error: "set_password_failed", details: err }, { status: 500 });
  }

  // 3) ask Keycloak to send verification email (optional)
  await adminFetch(`/users/${userId}/execute-actions-email`, {
    method: "PUT",
    body: JSON.stringify(["VERIFY_EMAIL"])
  }).catch(() => null);

  return NextResponse.json({ ok: true });
}
