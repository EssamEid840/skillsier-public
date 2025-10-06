export function randomString(bytes = 32) {
  const arr = new Uint8Array(bytes);
  (globalThis.crypto || (require("node:crypto") as any).webcrypto).getRandomValues(arr);
  return Buffer.from(arr).toString("base64url");
}

export async function sha256Base64Url(input: string) {
  const enc = new TextEncoder().encode(input);
  const buf = await (globalThis.crypto || (require("node:crypto") as any).webcrypto).subtle.digest("SHA-256", enc);
  return Buffer.from(new Uint8Array(buf)).toString("base64url");
}
