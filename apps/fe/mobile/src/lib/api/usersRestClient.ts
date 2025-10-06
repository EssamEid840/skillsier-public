import { Profile } from "@skillsier/sdk";

const BASE = process.env.EXPO_PUBLIC_USERS_BE_URL ?? "https://your-backend.example.com";

export async function getMe(accessToken: string) {
  const r = await fetch(`${BASE}/users/me`, { headers: { Authorization: `Bearer ${accessToken}` } });
  if (!r.ok) throw new Error("Failed");
  return r.json();
}

export async function getProfile() {
  return Profile.getMyProfile();
}