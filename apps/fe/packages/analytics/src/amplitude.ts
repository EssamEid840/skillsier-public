import type { AnalyticsClient, AnalyticsEvent } from "./events";

/**
 * Tiny Amplitude wrapper. Delegates to window.amplitude if present.
 */
declare global { interface Window { amplitude?: { track: (e: string, p?: any) => void } } }

export const AmplitudeClient = (): AnalyticsClient => ({
  capture: (e: AnalyticsEvent) => {
    try {
      if (typeof window !== "undefined" && window.amplitude) {
        window.amplitude.track(e.type, e);
      }
    } catch {}
  }
});
