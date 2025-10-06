import { colors, spacing, radii, typography } from "./tokens";

export const theme = {
  colors,
  spacing,
  radii,
  typography
};

export type Theme = typeof theme;

// convenience re-exports
export * from "./tokens";
export * from "./a11y";
