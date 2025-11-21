/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  transpilePackages: [
    '@skillsier/api',
    '@skillsier/auth',
    '@skillsier/hooks',
    '@skillsier/i18n',
    '@skillsier/stores',
    '@skillsier/types',
    '@skillsier/ui',
  ],
  images: {
    domains: ['api.dicebear.com'],
  },
};

module.exports = nextConfig;