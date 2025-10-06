#!/usr/bin/env node
// @ts-nocheck

/**
 * Generates TypeScript definitions from the Users OpenAPI schema using openapi-typescript.
 * If remote fetch fails, falls back to a minimal stub so builds don’t break offline.
 */

import cfg from "../../config/openapi/openapi-users.config";

function joinPosix(...parts: string[]) {
  return parts.join("/").replace(/\/+/g, "/").replace(/\/\.\//g, "/");
}

async function ensureDir(filePath: string) {
  const { promises: fs } = await import("fs");
  const segs = filePath.split("/").slice(0, -1);
  let cur = "";
  for (const s of segs) {
    cur = cur ? `${cur}/${s}` : s;
    try { await fs.mkdir(cur); } catch {}
  }
}

async function writeStub(outPath: string) {
  const { promises: fs } = await import("fs");
  const stub = `/* AUTO-GENERATED (stub) - BE schema not reachable
Replace by running: pnpm codegen:openapi:users when your BE is up */
declare namespace SkillsierUsersApi {}
export = SkillsierUsersApi;`;
  await ensureDir(outPath);
  await fs.writeFile(outPath, stub, "utf8");
  console.warn("⚠️ Wrote OpenAPI stub types (backend not reachable).");
}

async function run() {
  const input = cfg.input;
  const outputPath = joinPosix(cfg.output);

  console.log(`Generating OpenAPI types:\n  input:  ${input}\n  output: ${outputPath}`);

  try {
    const { default: openapiTS } = await import("openapi-typescript");
    const { promises: fs } = await import("fs");

    let dts: string | null = null;
    try {
      dts = await openapiTS(input, {
        transform(schemaObject: any) { return schemaObject; }
      });
    } catch (e) {
      console.warn("⚠️ Fetch/parse failed for input. Will try local file path or stub.", e?.message || e);
      try {
        const raw = await fs.readFile(input, "utf8");
        dts = await openapiTS(JSON.parse(raw));
      } catch {
        dts = null;
      }
    }

    if (!dts) {
      await writeStub(outputPath);
      return;
    }

    await ensureDir(outputPath);
    await fs.writeFile(outputPath, `/* AUTO-GENERATED - DO NOT EDIT */\n${dts}\n`, "utf8");
    console.log("✅ OpenAPI types generated.");
  } catch (e) {
    console.error("❌ OpenAPI generation failed fatally, writing stub:", e);
    await writeStub(outputPath);
  }
}

run();
