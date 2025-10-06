// Offline-friendly: use local SDL schema (Phase 10) instead of JSON introspection
const config = {
  schema: "packages/types/src/graphql/users/schema.graphql",
  documents: "packages/sdk/src/users/graphql/documents/**/*.graphql",
  generates: {
    "packages/sdk/src/users/graphql/generated/types.ts": {
      plugins: ["typescript"]
    },
    "packages/sdk/src/users/graphql/generated/hooks.ts": {
      plugins: ["typescript-operations", "typescript-react-apollo"],
      config: {
        withHooks: true,
        withHOC: false,
        withComponent: false
      }
    }
  },
  hooks: {}
};

export default config;
