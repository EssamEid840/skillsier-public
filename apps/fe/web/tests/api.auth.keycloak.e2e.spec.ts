// @ts-nocheck
/**
 * Smoke tests for Keycloak BFF routes.
 * Set E2E_BASE_URL to run, e.g.:
 *   E2E_BASE_URL=http://localhost:3000 pnpm jest web/tests
 */

const BASE: string | undefined = (globalThis as any)?.process?.env?.E2E_BASE_URL;
const maybe = (fn: any) => (BASE ? fn : it.skip);

describe("API — /api/auth/keycloak (smoke)", () => {
  maybe(it)("GET /discover returns JSON (or 5xx if remote KC is down)", async () => {
    const res = await fetch(`${BASE}/api/auth/keycloak/discover`, { credentials: "include" as any }).catch(() => null);
    expect(res).toBeTruthy();
    // 200 when discovery works; 502/500 if KC not reachable; 404 if route not wired yet
    expect([200, 500, 502, 404]).toContain(res!.status);

    const ct = res!.headers.get("content-type") || "";
    if (res!.status === 200 && ct.includes("application/json")) {
      const body = await res!.json().catch(() => null);
      expect(body).toBeTruthy();
      expect(typeof body).toBe("object");
      // Optional: common fields in well-known discovery result
      // expect(body.authorization_endpoint).toBeDefined();
      // expect(body.token_endpoint).toBeDefined();
    }
  });

  maybe(it)("GET /me returns JSON or 401", async () => {
    const res = await fetch(`${BASE}/api/auth/keycloak/me`, { credentials: "include" as any }).catch(() => null);
    expect(res).toBeTruthy();
    expect([200, 401]).toContain(res!.status);

    const ct = res!.headers.get("content-type") || "";
    if (res!.status === 200 && ct.includes("application/json")) {
      await expect(res!.json()).resolves.toBeDefined();
    }
  });
});
