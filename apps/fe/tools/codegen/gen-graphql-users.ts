// @ts-nocheck


/**
 * Programmatic GraphQL codegen runner with zero type deps at edit time.
 * After Phase 14 you’ll have proper types via @graphql-codegen/* and graphql.
 */

import config from "../../config/graphql/codegen";

async function run() {
  try {
    console.log("Generating GraphQL types/hooks for users...");
    const { generate } = await import("@graphql-codegen/cli");
    await generate(config as any, true);
    console.log("✅ GraphQL types/hooks generated.");
  } catch (e) {
    console.error("❌ GraphQL codegen failed:", e);
    (globalThis as any)?.process?.exit?.(1);
  }
}

run();
