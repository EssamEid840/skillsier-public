"use server";

import { Profile, Auth } from "@skillsier/sdk";

export async function noopAction() {
  return { ok: true };
}

export async function loadMe() {
  const me = await Auth.Me.me();
  const profile = await Profile.getMyProfile();
  return { me, profile };
}