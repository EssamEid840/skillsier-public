// fe/packages/ui/tailwind.config.ts
import type { Config } from "tailwindcss";

// This is a preset, not a full project config → use Partial<Config>
const preset = {
  theme: {
    extend: {
      colors: {
        brand: {
          DEFAULT: "#0ea5e9",
          50: "#f0f9ff",
          100: "#e0f2fe",
          200: "#bae6fd",
          300: "#7dd3fc",
          400: "#38bdf8",
          500: "#0ea5e9",
          600: "#0284c7",
          700: "#0369a1",
          800: "#075985",
          900: "#0c4a6e"
        }
      },
      borderRadius: {
        "2xl": "1rem"
      }
    }
  },
  plugins: []
} satisfies Partial<Config>;

export default preset;
