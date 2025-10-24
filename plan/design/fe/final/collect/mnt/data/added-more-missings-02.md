fe/
├── apps/
│   ├── web/
│   │   └── src/
│   │       ├── app/
│   │       │   └── [locale]/
│   │       │       ├── (admin)/
│   │       │       │   ├── admin-session/
│   │       │       │   │   └── page.tsx                          # JIT break-glass admin sessions
│   │       │       │   │                                          # BE: admin-be/admin_session — POST /v1/admin/sessions/jit
│   │       │       │   ├── audit-reporting/
│   │       │       │   │   ├── audit-logs/
│   │       │       │   │   │   └── page.tsx                      # Platform audit trail viewer
│   │       │       │   │   │                                      # BE: utility/audit — GET /v1/audit/logs
│   │       │       │   │   └── bi-exports/
│   │       │       │   │       └── page.tsx                      # BI extracts (CSV/Parquet) request/download
│   │       │       │   │                                          # BE: financial-be/reports — POST/GET /v1/reports/exports
│   │       │       │   ├── case-mgmt/
│   │       │       │   │   └── page.tsx                          # Disputes & appeals caseboard (cross-domain)
│   │       │       │   │                                          # BE: admin-be/case_mgmt — GET /v1/admin/cases ; PATCH /v1/admin/cases/{id}
│   │       │       │   ├── change-approval/
│   │       │       │   │   └── page.tsx                          # Two-person approvals (configs/financial)
│   │       │       │   │                                          # BE: admin-be/change_approval — GET/POST /v1/admin/change-approvals
│   │       │       │   ├── comms/
│   │       │       │   │   ├── broadcasts/
│   │       │       │   │   │   └── page.tsx                      # Broadcasts (email/push/SMS)
│   │       │       │   │   │                                      # BE: communications-be/broadcast — CRUD
│   │       │       │   │   ├── campaigns/
│   │       │       │   │   │   └── page.tsx                      # Lifecycle/marketing campaigns
│   │       │       │   │   │                                      # BE: communications-be/campaign — CRUD
│   │       │       │   │   ├── rate-limits/
│   │       │       │   │   │   └── page.tsx                      # Throughput & quotas
│   │       │       │   │   │                                      # BE: communications-be/admin — GET/PUT /v1/comms/admin/rate-limits
│   │       │       │   │   └── templates/
│   │       │       │   │       └── page.tsx                      # Template library (versioning/AB)
│   │       │       │   │                                          # BE: communications-be/template — CRUD
│   │       │       │   ├── experiments/
│   │       │       │   │   └── page.tsx                          # Experiments dashboard (assignments & metrics)
│   │       │       │   │                                          # BE: utility/experiments — GET /v1/experiments/active ; POST /v1/experiments/{id}/track
│   │       │       │   ├── feature-flags/
│   │       │       │   │   └── page.tsx                          # Kill switches & gradual rollout UI
│   │       │       │   │                                          # BE: utility/flags — GET/PUT /v1/flags
│   │       │       │   ├── financial-ops/
│   │       │       │   │   ├── chargebacks/
│   │       │       │   │   │   └── page.tsx                      # Chargeback handling & evidence
│   │       │       │   │   │                                      # BE: financial-be/chargeback — GET/POST /v1/chargebacks
│   │       │       │   │   ├── payouts-review/
│   │       │       │   │   │   └── page.tsx                      # Payout approvals & holds
│   │       │       │   │   │                                      # BE: financial-be/payout — GET /v1/payouts?status=PENDING ; POST /v1/payouts/{id}/approve
│   │       │       │   │   └── reconciliation/
│   │       │       │   │       └── page.tsx                      # Ledger ↔ gateway reconciliation
│   │       │       │   │                                          # BE: financial-be/recon — GET /v1/recon/runs ; POST /v1/recon/runs
│   │       │       │   ├── moderation/
│   │       │       │   │   ├── page.tsx                          # Reports queue & filters
│   │       │       │   │   │                                      # BE: users-be/moderation — GET /v1/admin/moderation/reports
│   │       │       │   │   └── [reportId]/
│   │       │       │   │       └── page.tsx                      # Review → warn/suspend/ban → notes & audit
│   │       │       │   │                                          # BE: users-be/moderation — POST /v1/admin/users/{id}/moderation/{action}
│   │       │       │   ├── ops/
│   │       │       │   │   └── refund-cases/
│   │       │       │   │       ├── page.tsx                      # Intake & queue
│   │       │       │   │       │                                  # BE: admin-be/refund_case — GET /v1/admin/refund-cases
│   │       │       │   │       └── [caseId]/
│   │       │       │   │           └── page.tsx                  # Case detail → approve/deny → post to ledger
│   │       │       │   │                                          # BE: admin-be/refund_case — POST /v1/admin/refund-cases/{id}/decision
│   │       │       │   │                                          # BE: financial-be/refund — POST /v1/refunds
│   │       │       │   └── search-quality/
│   │       │       │       ├── blacklists/
│   │       │       │       │   └── page.tsx                      # Blocked entities/terms
│   │       │       │       │                                      # BE: search-be/admin/safety_filters — POST/DELETE /v1/search/admin/safety
│   │       │       │       ├── boosts/
│   │       │       │       │   └── page.tsx                      # Query boosts & pinning
│   │       │       │       │                                      # BE: search-be/admin/ranking — PUT /v1/search/admin/boosts
│   │       │       │       ├── explainability/
│   │       │       │       │   └── page.tsx                      # Explain results for sample queries
│   │       │       │       │                                      # BE: search-be/admin/explain — GET (ES _explain proxy)
│   │       │       │       ├── facets/
│   │       │       │       │   └── page.tsx                      # Manage facets & weights
│   │       │       │       │                                      # BE: search-be/admin/facets — GET/PUT
│   │       │       │       ├── hygiene/
│   │       │       │       │   └── page.tsx                      # Dedup/archive/visibility ops
│   │       │       │       │                                      # BE: search-be/admin/hygiene — POST /v1/search/admin/hygiene/run
│   │       │       │       ├── indices/
│   │       │       │       │   └── page.tsx                      # Rollover/snapshots, lifecycle ops
│   │       │       │       │                                      # BE: search-be/admin/indices — POST /v1/search/admin/indices/*
│   │       │       │       ├── performance/
│   │       │       │       │   └── page.tsx                      # Metrics & alerts (slow queries, index health)
│   │       │       │       │                                      # BE: search-be/admin/performance — GET /v1/search/admin/performance
│   │       │       │       ├── query-logs/
│   │       │       │       │   └── page.tsx                      # Query logs & filters; view ES explain
│   │       │       │       │                                      # BE: search-be/admin/query-logs — GET /v1/search/admin/query-logs
│   │       │       │       ├── rewrites/
│   │       │       │       │   └── page.tsx                      # Query rewrites & pin rules
│   │       │       │       │                                      # BE: search-be/admin/rewrites — GET/POST
│   │       │       │       └── synonyms/
│   │       │       │           └── page.tsx                      # Manage synonyms, aliases, taxonomy
│   │       │       │                                              # BE: search-be/admin/taxonomy — POST/PUT /v1/search/admin/synonyms
│   │       │       ├── (dashboard)/
│   │       │       │   ├── budgets/
│   │       │       │   │   └── page.tsx                          # Client org budgets & spend controls
│   │       │       │   │                                          # BE: financial-be/budget — GET/POST /v1/budgets
│   │       │       │   ├── feed/
│   │       │       │   │   └── page.tsx                          # Personalized feed (jobs/talent/portfolio)
│   │       │       │   │                                          # BE: search-be/feed — GET /v1/search/feed?type=JOB|FREELANCER
│   │       │       │   │                                          # BE: search-be/feed-interactions — POST /v1/search/feed/{id}/interact
│   │       │       │   ├── orgs/
│   │       │       │   │   └── [orgId]/
│   │       │       │   │       └── settings/
│   │       │       │   │           └── billing/
│   │       │       │   │               └── tax-vat/
│   │       │       │   │                   └── page.tsx          # Org tax/VAT profile
│   │       │       │   │                                           # BE: admin-be/business_verification — GET/PUT
│   │       │       │   ├── search/
│   │       │       │   │   ├── alerts/
│   │       │       │   │   │   └── page.tsx                      # Saved-search alerts (channels/schedule)
│   │       │       │   │   │                                      # BE: search-be/saved_search — GET/PUT schedules
│   │       │       │   │   ├── personalization/
│   │       │       │   │   │   └── page.tsx                      # Pause/reset, prefer/hide entities
│   │       │       │   │   │                                      # BE: search-be/personalization — GET/PUT /v1/search/personalization
│   │       │       │   │   └── saved/
│   │       │       │   │       └── page.tsx                      # Saved searches list/CRUD
│   │       │       │   │                                          # BE: search-be/saved_search — GET/POST /v1/saved-searches
│   │       │       │   ├── status/
│   │       │       │   │   └── page.tsx                          # Public status & incidents
│   │       │       │   │                                          # BE: utility/status — GET /v1/status ; GET /v1/incidents
│   │       │       │   ├── support/
│   │       │       │   │   ├── help-center/
│   │       │       │   │   │   └── page.tsx                      # Help center (articles search)
│   │       │       │   │   │                                      # BE: communications-be/knowledge — GET /v1/help/articles?q=
│   │       │       │   │   └── tickets/
│   │       │       │   │       ├── new/
│   │       │       │   │       │   └── page.tsx                  # Create ticket (attachments)
│   │       │       │   │       │                                  # BE: communications-be/ticket — POST /v1/support/tickets
│   │       │       │   │       │                                  # BE: storage-be/asset — POST /v1/storage/uploads (signed URLs)
│   │       │       │   │       ├── page.tsx                      # My tickets list
│   │       │       │   │       │                                  # BE: communications-be/ticket — GET /v1/support/tickets?me
│   │       │       │   │       └── [ticketId]/
│   │       │       │   │           └── page.tsx                  # Ticket detail, replies, SLA
│   │       │       │   │                                          # BE: communications-be/ticket — GET/POST /v1/support/tickets/{id}/messages
│   │       │       │   ├── billing/
│   │       │       │   │   └── subscription/
│   │       │       │   │       ├── dunning/
│   │       │       │   │       │   └── page.tsx                  # Failed invoices & retries
│   │       │       │   │       │                                  # BE: financial-be/invoice — GET failed/retry invoices
│   │       │       │   │       ├── manage/
│   │       │       │   │       │   └── page.tsx                  # Manage current subscription
│   │       │       │   │       │                                  # BE: financial-be/subscription — GET /v1/subscriptions/me
│   │       │       │   │       └── plans/
│   │       │       │   │           └── page.tsx                  # Subscription plans & checkout
│   │       │       │   │                                           # BE: financial-be/subscription — GET /v1/subscriptions/plans
│   │       │       │   ├── contracts/
│   │       │       │   │   ├── [contractId]/
│   │       │       │   │   │   ├── acceptance/
│   │       │       │   │   │   │   └── page.tsx                  # Accept deliverables/milestones
│   │       │       │   │   │   │                                  # BE: contracts-be/deliverable — POST accept
│   │       │       │   │   │   ├── deliverables/
│   │       │       │   │   │   │   ├── page.tsx                  # Deliverables list
│   │       │       │   │   │   │   │                              # BE: contracts-be/deliverable — GET /v1/contracts/{id}/deliverables
│   │       │       │   │   │   │   └── [deliverableId]/
│   │       │       │   │   │   │       └── versions/
│   │       │       │   │   │   │           └── page.tsx          # File versions (download/preview)
│   │       │       │   │   │   │                                   # BE: storage-be/asset — GET versions
│   │       │       │   │   │   └── work-diary/
│   │       │       │   │   │       └── screenshots/
│   │       │       │   │   │           └── [entryId]/
│   │       │       │   │   │               └── page.tsx          # Screenshot detail (blur/delete if allowed)
│   │       │       │   │   │                                        # BE: contracts-be/work_diary — GET /v1/work-diary/entries/{id}
│   │       │       │   │   └── disputes/
│   │       │       │   │       └── [disputeId]/
│   │       │       │   │           └── page.tsx                  # Dispute detail & submit evidence
│   │       │       │   │                                          # BE: contracts-be/dispute — GET/POST
│   │       │       │   └── reviews/
│   │       │       │       ├── my/
│   │       │       │       │   └── page.tsx                      # My reviews
│   │       │       │       │                                      # BE: reviews-be/review — GET /v1/reviews/me
│   │       │       │       └── rate/
│   │       │       │           └── [targetId]/
│   │       │       │               └── page.tsx                  # Leave a review
│   │       │       │                                              # BE: reviews-be/review — POST /v1/reviews
│   │       │       └── settings/
│   │       │           ├── security/
│   │       │           │   ├── data-export/
│   │       │           │   │   └── page.tsx                      # Export my data
│   │       │           │   │                                      # BE: users-be/account — POST /v1/user/export
│   │       │           │   ├── data-delete/
│   │       │           │   │   └── page.tsx                      # Request delete
│   │       │           │   │                                      # BE: users-be/account — POST /v1/user/delete
│   │       │           │   ├── deactivate/
│   │       │           │   │   └── page.tsx                      # Deactivate account
│   │       │           │   │                                      # BE: users-be/account — POST /v1/user/deactivate
│   │       │           │   ├── devices-sessions/
│   │       │           │   │   └── page.tsx                      # Active devices & sessions
│   │       │           │   │                                      # BE: users-be/account — GET /v1/user/sessions
│   │       │           │   └── mfa/
│   │       │           │       └── page.tsx                      # MFA enable/disable
│   │       │           │                                          # BE: users-be/auth — PUT /v1/user/mfa
│   │       │           └── notifications/
│   │       │               └── digests-quiet-hours/
│   │       │                   └── page.tsx                      # Email digests & quiet hours
│   │       │                                                      # BE: users-be/preferences — GET/PUT (digest cadence, DND window)
│   │       └── lib/
│   │           └── api/
│   │               ├── audit/
│   │               │   ├── audit.ts                              # Audit logs client
│   │               │   └── exports.ts                            # BI export jobs client
│   │               ├── budgets/
│   │               │   └── budgets.ts                            # Budgets APIs
│   │               ├── moderation/
│   │               │   └── moderation.ts                         # Moderation actions & reports
│   │               ├── refunds/
│   │               │   └── refunds.ts                            # Refund cases & financial refunds
│   │               ├── search/
│   │               │   ├── personalization.ts                    # Personalization API client
│   │               │   └── saved-search.ts                       # Saved-search CRUD & alerts
│   │               ├── search-admin/
│   │               │   ├── hygiene.ts                            # Run hygiene jobs
│   │               │   ├── query-logs.ts                         # Query logs fetcher
│   │               │   └── search-admin.ts                       # Synonyms/boosts/facets/rewrites/perf
│   │               ├── settings/
│   │               │   └── notifications.ts                      # Channels, email prefs, quiet hours
│   │               ├── status/
│   │               │   └── status.ts                             # Status & incidents APIs
│   │               └── support/
│   │                   └── tickets.ts                            # Tickets CRUD & messages
│   └── mobile/
│       └── app/
│           ├── (dashboard)/
│           │   ├── billing/
│           │   │   └── subscription/
│           │   │       ├── manage.tsx                            # Manage subscription (mobile)
│           │   │       │                                          # BE: financial-be/subscription — GET /v1/subscriptions/me
│           │   │       └── plans.tsx                             # Subscription plans (mobile)
│           │   │                                                  # BE: financial-be/subscription — GET /v1/subscriptions/plans
│           │   ├── contracts/
│           │   │   └── [contractId]/
│           │   │       └── work-diary/
│           │   │           └── screenshots/
│           │   │               └── [entryId].tsx                 # Screenshot detail (mobile)
│           │   │                                                  # BE: contracts-be/work_diary — GET /v1/work-diary/entries/{id}
│           │   ├── feed/
│           │   │   └── index.tsx                                 # Mobile feed (same BE as web)
│           │   │                                                  # BE: search-be/feed — GET /v1/search/feed
│           │   └── support/
│           │       ├── help-center/
│           │       │   └── index.tsx                             # Help articles (mobile)
│           │       │                                              # BE: communications-be/knowledge — GET /v1/help/articles
│           │       └── tickets/
│           │           ├── index.tsx                             # My tickets list (mobile)
│           │           │                                          # BE: communications-be/ticket — GET /v1/support/tickets?me
│           │           └── [ticketId]/
│           │               └── index.tsx                         # Ticket detail & reply
│           │                                                      # BE: communications-be/ticket — POST /v1/support/tickets/{id}/messages
│           ├── settings/
│           │   ├── notifications/
│           │   │   ├── channels.tsx                              # Notification channels (mobile)
│           │   │   │                                              # BE: users-be/communication_channels — GET/PUT
│           │   │   └── digests-quiet-hours.tsx                   # Digests & quiet hours (mobile)
│           │   │                                                  # BE: users-be/preferences — GET/PUT
│           │   └── security/
│           │       ├── data-export-delete.tsx                    # Export / delete account (mobile)
│           │       │                                              # BE: users-be/account — POST /v1/user/export, /delete
│           │       └── devices-sessions.tsx                      # Active devices & sessions (mobile)
│           │                                                      # BE: users-be/account — GET /v1/user/sessions
│           └── status/
│               └── index.tsx                                     # Status & incidents (mobile)
│                                                                  # BE: utility/status — GET /v1/status, /v1/incidents
├── packages/
│   ├── shared/
│   │   └── src/
│   │       └── features/
│   │           ├── notifications/
│   │           │   └── store.ts                                  # UI-only toggles; server state via queries
│   │           └── subscriptions/
│   │               ├── query-keys.ts                             # ['subscription','plans'|'me']
│   │               └── store.ts                                  # Local UI state for plans/checkout
│   └── types/
│       └── src/
│           ├── audit.ts                                          # Audit Log/Export DTOs
│           ├── moderation.ts                                     # ModerationCase/Appeal types
│           ├── search-admin.ts                                   # Taxonomy/Facet/Rewrite/Speller DTOs
│           ├── subscription.ts                                   # Plan/Entitlement types
│           └── support.ts                                        # Ticket/Message types
