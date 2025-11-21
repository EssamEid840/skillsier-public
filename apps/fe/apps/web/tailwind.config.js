/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './app/**/*.{js,ts,jsx,tsx,mdx}',
    './components/**/*.{js,ts,jsx,tsx,mdx}',
    '../../packages/ui/src/**/*.{js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: '#E60023',
          50: '#FFE5E9',
          100: '#FFCCD4',
          200: '#FF99AA',
          300: '#FF667F',
          400: '#FF3355',
          500: '#E60023',
          600: '#B3001C',
          700: '#800015',
          800: '#4D000D',
          900: '#1A0004',
        },
        secondary: {
          DEFAULT: '#111111',
          50: '#F5F5F5',
          100: '#E5E5E5',
          200: '#CCCCCC',
          300: '#B3B3B3',
          400: '#999999',
          500: '#808080',
          600: '#666666',
          700: '#4D4D4D',
          800: '#333333',
          900: '#1A1A1A',
        },
      },
    },
  },
  plugins: [],
};