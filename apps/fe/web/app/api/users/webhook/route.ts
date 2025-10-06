export const runtime = "nodejs";
import { NextRequest, NextResponse } from "next/server";
import crypto from "crypto";

export async function POST(req: NextRequest) {
  const secret = process.env.USERS_WEBHOOK_SECRET;
  const raw = Buffer.from(await req.arrayBuffer());

  if (secret) {
    const sig = req.headers.get("x-signature") || "";
    const mac = crypto.createHmac("sha256", secret).update(raw).digest("hex");
    if (sig !== mac) return NextResponse.json({ error: "invalid_signature" }, { status: 401 });
  }

  // TODO: handle events. For now, just 200 OK.
  return NextResponse.json({ ok: true });
}
