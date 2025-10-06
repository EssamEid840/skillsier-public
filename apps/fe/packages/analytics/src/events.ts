export type AnalyticsEvent =
  | { type: "auth_login"; method: "password" | "google"; }
  | { type: "auth_register"; method: "password" | "google"; }
  | { type: "profile_saved"; }
  | { type: "portfolio_item_created"; id: string; };

export type AnalyticsClient = {
  capture: (e: AnalyticsEvent) => void;
};

export function combine(...clients: AnalyticsClient[]): AnalyticsClient {
  return { capture: (e) => clients.forEach(c => c.capture(e)) };
}
