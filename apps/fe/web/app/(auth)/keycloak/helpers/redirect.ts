export function redirectTo(url: string) {
  if (typeof window !== "undefined") {
    window.location.assign(url);
  }
}
