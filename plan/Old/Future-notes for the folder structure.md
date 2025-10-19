1. Migrations & data policy. You have migrator commands—great. Add conventions: deterministic ordering, idempotent seeds, and rollbacks per service. Consider golang-migrate with Makefile targets (migrate.up/down). 



3. Ensure internal/* never imports deployments/ or scripts; and domain never imports infrastructure.

4. Security defaults. Secret sourcing from Vault/SealedSecrets, DB least-privilege users, and mTLS (or Dapr S2S auth) between services.

5. API versioning. Create /docs/api.md per service (present) and add explicit v1 routers plus deprecation notes and changelog. => Supposed to be done


7. K8s overlays. You have deployments/k8s/; introduce kustomize overlays (shared/, local/, prod/), add NetworkPolicy, PodSecurity, resources/limits, PDBs, and readiness/liveness on all services for safe rollouts.