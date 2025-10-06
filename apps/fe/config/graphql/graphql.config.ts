// No types import; plain object to avoid needing 'graphql-config' types now
const config = {
  projects: {
    users: {
      schema: "packages/sdk/src/users/graphql/generated/graphql.schema.json",
      documents: ["packages/sdk/src/users/graphql/documents/**/*.graphql"],
      extensions: {
        codegen: {
          // editor hint only
        }
      }
    }
  }
};

export default config;
