/** @type {import('jest').Config} */
module.exports = {
  testEnvironment: "jsdom",
  transform: {
    "^.+\\.(t|j)sx?$": ["@swc/jest"]
  },
  moduleFileExtensions: ["ts", "tsx", "js", "jsx", "json"],
  moduleNameMapper: {
    // Allow importing images and styles in tests
    "\\.(css|less|scss|sass)$": "identity-obj-proxy"
  },
  setupFilesAfterEnv: [],
  testPathIgnorePatterns: ["/node_modules/", "/dist/", "/.next/"],
  coveragePathIgnorePatterns: ["/node_modules/", "/dist/", "/.next/"]
};
