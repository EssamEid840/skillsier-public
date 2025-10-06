const env = ((globalThis as any)?.process?.env ?? {}) as Record<string, string | undefined>;
const cfg = {
  input: env.USERS_OPENAPI_URL || "packages/types/src/openapi/users.schema.json",
  output: "packages/types/src/openapi/users.d.ts"
};
export default cfg;
