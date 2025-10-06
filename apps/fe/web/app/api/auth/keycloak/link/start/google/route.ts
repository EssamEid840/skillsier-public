export const runtime = "nodejs";
import { NextResponse } from "next/server";
export async function GET() {
  return NextResponse.json({ error: "not_implemented", note: "Account linking flow will be added later." }, { status: 501 });
}
