const createNextIntlPlugin = require('next-intl/plugin');
const withNextIntl = createNextIntlPlugin('./i18n.ts');

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  transpilePackages: ['@skillsier/shared', '@skillsier/ui', '@skillsier/types'],
  images: {
    domains: ['localhost'],
  },
  experimental: {
    optimizePackageImports: ['@skillsier/ui'],
  },
};

module.exports = withNextIntl(nextConfig);
