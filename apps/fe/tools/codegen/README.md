# Codegen tools

This folder contains small scripts to generate:
- **OpenAPI** TypeScript declarations for Users REST API
- **GraphQL** types and React hooks for Users GraphQL API

## OpenAPI (users)
- Config: `config/openapi/openapi-users.config.ts`
- Script: `tools/codegen/gen-openapi-users.ts`

Run (after Phase 14 install):
```bash
pnpm tsx tools/codegen/gen-openapi-users.ts
