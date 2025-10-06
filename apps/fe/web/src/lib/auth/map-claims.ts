type AppUser = { sub: string; email?: string; name?: string; username?: string };

export function mapClaimsToUser(payload: any): AppUser {
  return {
    sub: payload.sub,
    email: payload.email,
    name: payload.name || payload.given_name || payload.preferred_username,
    username: payload.preferred_username,
  };
}
