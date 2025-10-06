// @ts-nocheck
/**
 * Smoke tests for users REST & GraphQL BFF routes.
 * Set E2E_BASE_URL to run, e.g.:
 *   E2E_BASE_URL=http://localhost:3000 pnpm jest web/tests
 */

const BASE: string | undefined = (globalThis as any)?.process?.env?.E2E_BASE_URL;
const maybe = (fn: any) => (BASE ? fn : it.skip);

describe("API — /api/users (smoke)", () => {
  maybe(it)("REST: /api/users/rest/[...path] responds", async () => {
    // Hit a harmless sub-path that should exist (catch-all).
    const res = await fetch(`${BASE}/api/users/rest/ping`, {
      method: "GET",
      credentials: "include" as any,
    }).catch(() => null);

    expect(res).toBeTruthy();
    // Depending on your handler it may be 200 (mock), 404 (no route), or 405 (method not allowed)
    expect([200, 404, 405, 401]).toContain(res!.status);

    const ct = res!.headers.get("content-type") || "";
    // If JSON is returned, it should be parseable
    if (ct.includes("application/json")) {
      await expect(res!.json()).resolves.toBeDefined();
    }
  });

  maybe(it)("GraphQL: /api/users-graphql responds", async () => {
    const res = await fetch(`${BASE}/api/users-graphql`, {
      method: "POST",
      headers: { "content-type": "application/json" },
      credentials: "include" as any,
      body: JSON.stringify({ query: "{ __typename }" }),
    }).catch(() => null);

    expect(res).toBeTruthy();
    // 200 expected for GraphQL; some setups may return 401 if auth is required
    expect([200, 401]).toContain(res!.status);

    if (res!.status === 200) {
      const data = await res!.json().catch(() => null);
      expect(data).toBeTruthy();
      // shape: { data?: ..., errors?: ... }
      expect(typeof data).toBe("object");
    }
  });
});
