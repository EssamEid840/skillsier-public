/** Parse a cookie header into a simple map. */
export function parseCookieHeader(header: string | null | undefined) {
  const out: Record<string, string> = {};
  if (!header) return out;
  header.split(";").forEach(part => {
    const [k, ...rest] = part.trim().split("=");
    out[k] = decodeURIComponent(rest.join("=") ?? "");
  });
  return out;
}
