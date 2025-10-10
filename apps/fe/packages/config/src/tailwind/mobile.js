// packages/config/src/tailwind/mobile.js
const baseConfig = require('./base');

module.exports = {
  ...baseConfig,
  content: [
    './app/**/*.{js,jsx,ts,tsx}',
    './src/**/*.{js,jsx,ts,tsx}',
  ],
};