import type { AnalyticsClient, AnalyticsEvent } from "./events";

/**
 * Tiny PostHog wrapper. If posthog-js is loaded globally, we delegate to it.
 * Otherwise, we no-op. No hard dependency added.
 */
declare global { interface Window { posthog?: { capture: (e: string, p?: any) => void } } }

export const PosthogClient = (): AnalyticsClient => ({
  capture: (e: AnalyticsEvent) => {
    try {
      if (typeof window !== "undefined" && window.posthog) {
        window.posthog.capture(e.type, e);
      }
    } catch {}
  }
});
