const BASE = process.env.EXPO_PUBLIC_WEB_API_BASE_URL ?? "http://localhost:3000";

export async function backendRegister(input: { fullName: string; email: string; password: string }) {
  const r = await fetch(`${BASE}/api/auth/keycloak/register`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(input),
  });
  if (!r.ok) throw new Error(`Register failed: ${r.status}`);
  return await r.json().catch(() => ({}));
}

export async function backendForgotPassword(email: string) {
  const r = await fetch(`${BASE}/api/auth/keycloak/forgot-password`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email }),
  });
  if (!r.ok) throw new Error(`Forgot password failed: ${r.status}`);
  return await r.json().catch(() => ({}));
}
