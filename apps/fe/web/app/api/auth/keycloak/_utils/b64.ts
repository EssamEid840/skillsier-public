import crypto from "crypto";

export function b64url(input: Buffer | string) {
  const buf = Buffer.isBuffer(input) ? input : Buffer.from(input);
  return buf.toString("base64").replace(/\+/g, "-").replace(/\//g, "_").replace(/=+$/g, "");
}

export function randomUrlSafe(size = 32) {
  return b64url(crypto.randomBytes(size));
}

export async function sha256B64Url(verifier: string) {
  const hash = crypto.createHash("sha256").update(verifier).digest();
  return b64url(hash);
}

export function parseJwtPayload<T = any>(jwt: string): T | null {
  try {
    const [, payload] = jwt.split(".");
    return JSON.parse(Buffer.from(payload, "base64").toString("utf8"));
  } catch {
    return null;
  }
}
