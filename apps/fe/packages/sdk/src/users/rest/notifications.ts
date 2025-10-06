import { usersFetch } from "../clients/rest";

export type NotificationSettings = { email: boolean; push: boolean };

export async function getNotificationSettings() {
  return usersFetch<NotificationSettings>("rest/notifications/settings", { method: "GET" });
}

export async function updateNotificationSettings(patch: Partial<NotificationSettings>) {
  return usersFetch<NotificationSettings>("rest/notifications/settings", {
    method: "PATCH",
    body: JSON.stringify(patch)
  });
}
