export function startGoogle() {
  window.location.assign("/api/auth/keycloak/oauth/start/google");
}
export async function logout() {
  await fetch("/api/auth/keycloak/logout", { method: "POST" });
  window.location.assign("/");
}
export async function me() {
  const r = await fetch("/api/auth/keycloak/me", { cache: "no-store" });
  return r.json();
}
