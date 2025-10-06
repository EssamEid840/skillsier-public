/** @type {import('next').NextConfig} */
const nextConfig = {
  transpilePackages: [
    // we will add these deps in a later phase;
    // keeping them here avoids forgetting to configure when we do
    "react-native",
    "react-native-web",
    "react-native-reanimated",
    "react-native-gesture-handler",
    "nativewind",
    "@skillsier/ui",
    "@skillsier/sdk",
    "@skillsier/auth",
    "@skillsier/i18n",
    "@skillsier/analytics",
    "@skillsier/flags",
    "@skillsier/state"
  ],
  webpack: (config) => {
    config.resolve.alias = {
      ...(config.resolve.alias || {}),
      "react-native$": "react-native-web"
    };
    return config;
  }
};
export default nextConfig;
