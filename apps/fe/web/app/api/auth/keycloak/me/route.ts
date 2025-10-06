import { NextResponse } from "next/server";
import { cookies } from "next/headers";
import { verifyJwt } from "@/src/lib/auth/jwks";
import { mapClaimsToUser } from "@/src/lib/auth/map-claims";

export async function GET() {
  try {
    const jar = await cookies();
    const token = jar.get(process.env.AUTH_COOKIE_NAME || "kc_session")?.value;
    if (!token) {
      return NextResponse.json({ authenticated: false }, { status: 200 });
    }
    const payload = await verifyJwt(token);
    const user = mapClaimsToUser(payload);
    return NextResponse.json({ authenticated: true, user }, { status: 200 });
  } catch {
    return NextResponse.json({ authenticated: false }, { status: 200 });
  }
}