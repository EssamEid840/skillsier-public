/** RFC7636 PKCE: code_verifier & code_challenge (S256). */
export async function createPkcePair() {
  const random = crypto.getRandomValues(new Uint8Array(32));
  const verifier = btoa(String.fromCharCode(...random))
    .replace(/=/g, "").replace(/\+/g, "-").replace(/\//g, "_");

  const data = new TextEncoder().encode(verifier);
  const digest = await crypto.subtle.digest("SHA-256", data);
  const hash = btoa(String.fromCharCode(...new Uint8Array(digest)))
    .replace(/=/g, "").replace(/\+/g, "-").replace(/\//g, "_");

  return { verifier, challenge: hash, method: "S256" as const };
}
