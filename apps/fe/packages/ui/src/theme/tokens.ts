/**
 * Design tokens shared by RN + RNW components.
 * Keep primitive values here; compose in theme/index.ts
 */
export const colors = {
  primary: "#0ea5e9",
  primaryText: "#ffffff",
  bg: "#ffffff",
  fg: "#111827",
  mutedBg: "#f3f4f6",
  mutedFg: "#4b5563",
  danger: "#ef4444",
  success: "#10b981",
  warning: "#f59e0b",
  border: "#e5e7eb",
};

export const spacing = {
  xs: 4,
  sm: 8,
  md: 12,
  lg: 16,
  xl: 24,
  "2xl": 32
};

export const radii = {
  sm: 6,
  md: 10,
  lg: 14,
  xl: 18
};

export const typography = {
  family: {
    regular: "System",
    medium: "System",
    bold: "System"
  },
  size: {
    xs: 12,
    sm: 14,
    md: 16,
    lg: 18,
    xl: 22,
    "2xl": 26
  }
};

export type Tokens = {
  colors: typeof colors;
  spacing: typeof spacing;
  radii: typeof radii;
  typography: typeof typography;
};
