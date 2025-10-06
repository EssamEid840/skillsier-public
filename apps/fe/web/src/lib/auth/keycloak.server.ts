import { cookies } from "next/headers";
import { verifyJwt } from "./jwks";
import { mapClaimsToUser } from "./map-claims";

export async function getServerUser() {
  const jar = await cookies()
  const token = jar.get(process.env.AUTH_COOKIE_NAME || "kc_session")?.value;
  if (!token) return null;
  try {
    const payload = await verifyJwt(token);
    return mapClaimsToUser(payload);
  } catch {
    return null;
  }
}
