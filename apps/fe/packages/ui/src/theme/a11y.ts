import { AccessibilityInfo } from "react-native";

/** Announce a message for screen readers. */
export function announce(msg: string) {
  try { AccessibilityInfo.announceForAccessibility?.(msg); } catch {}
}

/** Larger touch target by default. */
export const hitSlop = { top: 8, bottom: 8, left: 8, right: 8 };
