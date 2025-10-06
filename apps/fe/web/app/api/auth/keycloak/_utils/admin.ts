import { reqEnv, optEnv } from "./env";

// app/api/auth/keycloak/_utils/admin.ts
export async function getMgmtAccessToken() {
  const issuer = process.env.KEYCLOAK_ISSUER_URL!;
  const id = process.env.KEYCLOAK_MGMT_CLIENT_ID!;
  const secret = process.env.KEYCLOAK_MGMT_CLIENT_SECRET!;
  const body = new URLSearchParams({
    grant_type: "client_credentials",
    client_id: id,
    client_secret: secret,
  });
  const r = await fetch(`${issuer}/protocol/openid-connect/token`, {
    method: "POST",
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
    body,
    cache: "no-store",
  });
  if (!r.ok) {
    const text = await r.text();           // 👈 show why it failed
    console.error("Mgmt token error:", r.status, text);
    throw new Error("Failed to obtain mgmt token");
  }
  const j = await r.json();
  return j.access_token as string;
}


export async function adminFetch(path: string, init?: RequestInit) {
  const issuer = reqEnv("KEYCLOAK_ISSUER_URL");
  const realm = issuer.split("/realms/")[1];
  const base = issuer.replace(`/realms/${realm}`, "");
  const token = await getMgmtAccessToken();
  const url = `${base}/admin/realms/${realm}${path}`;
  const r = await fetch(url, {
    ...init,
    headers: {
      "Authorization": `Bearer ${token}`,
      "Content-Type": "application/json",
      ...(init?.headers || {})
    }
  });
  return r;
}
