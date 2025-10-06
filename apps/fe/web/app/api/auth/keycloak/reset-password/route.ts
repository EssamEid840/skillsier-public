export const runtime = "nodejs";
import { NextResponse } from "next/server";

export async function POST() {
  return NextResponse.json(
    { error: "not_implemented", note: "Use the email link sent by Keycloak to reset password." },
    { status: 501 }
  );
}
