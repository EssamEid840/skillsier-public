export type AuthProvider = { id: "email" | "google"; label: string; enabled: boolean };

export function getAuthProviders(): AuthProvider[] {
  return [
    { id: "email", label: "Email & Password", enabled: true },
    { id: "google", label: "Google", enabled: true }, // set false until configured if you want
  ];
}
