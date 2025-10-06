// @ts-nocheck
/**
 * Smoke tests for public user-facing pages.
 * Set E2E_BASE_URL to run against a live dev server, e.g.:
 *   E2E_BASE_URL=http://localhost:3000 pnpm jest web/tests
 */

const BASE: string | undefined = (globalThis as any)?.process?.env?.E2E_BASE_URL;

const maybe = (fn: any) => (BASE ? fn : it.skip);

describe("Users — account pages (smoke)", () => {
  maybe(it)("GET / (features)/users root renders", async () => {
    const res = await fetch(`${BASE}/(features)/users`, { redirect: "manual" as any }).catch(() => null);
    expect(res).toBeTruthy();
    // 200 for public page, or 3xx if it redirects to login
    expect([200, 301, 302, 307, 308]).toContain(res!.status);
  });

  maybe(it)("GET /u/[handle] renders public profile", async () => {
    // Use a stable handle that exists in your dev DB; fallback to checking 200/404
    const res = await fetch(`${BASE}/u/example`, { redirect: "manual" as any }).catch(() => null);
    expect(res).toBeTruthy();
    // public profile may not exist -> 404 is acceptable
    expect([200, 404]).toContain(res!.status);
    // If it is 200, it should be HTML
    if (res!.status === 200) {
      const ct = res!.headers.get("content-type") || "";
      expect(ct.includes("text/html")).toBe(true);
    }
  });
});
