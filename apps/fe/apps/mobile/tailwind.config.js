/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './app/**/*.{js,jsx,ts,tsx}',
    './components/**/*.{js,jsx,ts,tsx}',
    '../../packages/ui/src/**/*.{js,jsx,ts,tsx}',
  ],
  presets: [require('nativewind/preset')],
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: '#E60023',
          50: '#FFEBEE',
          100: '#FFC7CE',
          200: '#FF8A95',
          300: '#FF4D5C',
          400: '#FF1A2F',
          500: '#E60023',
          600: '#CC001F',
          700: '#B3001B',
          800: '#990017',
          900: '#800013',
        },
      },
    },
  },
  plugins: [],
};