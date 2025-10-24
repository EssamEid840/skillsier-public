fe/
├── .husky/
│   ├── pre-commit
│   └── pre-push
├── .vscode/
│   ├── extensions.json
│   ├── launch.json
│   └── settings.json
├── apps/
│   ├── mobile/
│   │   ├── app/
│   │   │   ├── (auth)/
│   │   │   │   ├── _layout.tsx
│   │   │   │   ├── callback.tsx
│   │   │   │   ├── login.tsx
│   │   │   │   └── register.tsx
│   │   │   ├── (dashboard)/
│   │   │   │   ├── contracts/
│   │   │   │   │   └── page.tsx                           # Contracts list
│   │   │   │   │                                          # - Active & completed, filters, sort
│   │   │   │   │                                          # BE: contracts-be/contract — GET /v1/contracts?status=active
│   │   │   │   ├── dashboard/
│   │   │   │   │   └── page.tsx                           # Dashboard home (role-based)
│   │   │   │   │                                          # Freelancer: active proposals/contracts, earnings, recs
│   │   │   │   │                                          # Client: active jobs, spend overview, recent proposals, recs
│   │   │   │   │                                          # BE: users-be/user — GET /v1/users/me
│   │   │   │   │                                          # BE: analytics — GET /v1/analytics/dashboard
│   │   │   │   │                                          # BE: jobs-be — GET /v1/jobs/my-jobs
│   │   │   │   │                                          # BE: proposals-be — GET /v1/proposals/my-proposals
│   │   │   │   └── deliverables/
│   │   │   │       ├── submit/
│   │   │   │       │   └── page.tsx                       # Submit deliverable (upload files, notes)
│   │   │   │       │                                      # BE: contracts-be/deliverable — POST /v1/contracts/{id}/deliverables
│   │   │   │       │                                      # BE: storage-be/asset — POST /v1/storage/upload
│   │   │   │       └── page.tsx                           # All deliverables (status, history)
│   │   │   │                                              # BE: contracts-be/deliverable — GET /v1/contracts/{id}/deliverables
│   │   │   │                                              # BE: storage-be/asset — GET /v1/storage/download/{file_id}
│   │   │   ├── (shared)/
│   │   │   │   └── deliverables/
│   │   │   │       ├── components/
│   │   │   │       │   └── DeliverableCard.tsx            # BE: none (typed props), actions wired to mutations
│   │   │   │       ├── mutations.ts                       # BE: contracts-be/deliverable — POST /v1/contracts/{id}/deliverables
│   │   │   │       └── queries.ts                         # BE: contracts-be/deliverable — GET /v1/contracts/{id}/deliverables
│   │   │   ├── +not-found.tsx
│   │   │   ├── _layout.tsx
│   │   │   └── index.tsx
│   │   └── src/
│   └── web/
│       └── src/
│           └── app/
│               └── [locale]/
│                   └── (dashboard)/
│                       ├── notifications/
│                       │   ├── settings/
│                       │   │   └── page.tsx                # Notification preferences (email/push/in-app)
│                       │   │                               # BE: users-be/preferences — PATCH /v1/users/me/preferences
│                       │   │                               # BE: communications-be/subscription — POST /v1/push/subscribe
│                       │   └── page.tsx                    # Notifications inbox (unread, categories, mark read)
│                       │                                   # BE: communications-be/notification — GET /v1/notifications?status=unread
│                       ├── search/
│                       │   ├── jobs/
│                       │   │   └── page.tsx                # Advanced job search
│                       │   │                               # - Full-text search
│                       │   │                               # - Faceted filters
│                       │   │                               # - Autocomplete suggestions
│                       │   │                               # - Search history
│                       │   │                               # - Save search
│                       │   │                               # BE: search-be/query
│                       │   │                               # POST /v1/search/jobs
│                       │   │                               # Body: { query, filters: {...}, sort, page }
│                       │   │                               # BE: search-be/suggestions
│                       │   │                               # GET /v1/suggestions?q={query}
│                       │   └── freelancers/
│                       │       └── page.tsx                # Advanced freelancer search (client)
│                       │                                   # - Search by skills
│                       └── settings/
│                           └── web-push/
│                               └── page.tsx                # Web Push: subscribe/unsubscribe, device listing
│                                                           # BE: communications-be/subscription — POST /v1/push/subscribe
│                                                           # BE: communications-be/subscription — DELETE /v1/push/{id}
