```
fe/
├── apps/
│   ├── mobile/
│   │   └── app/
│   │       ├── (auth)/
│   │       │   └── onboarding/
│   │       │       └── kyc/
│   │       │           └── index.tsx                        # KYC onboarding (mobile)
│   │       │                                                # - Upload ID docs, liveness
│   │       │                                                # - Track review status
│   │       │                                                # BE: admin-be/kyc_case
│   │       │                                                # POST /v1/kyc/cases
│   │       │                                                # GET  /v1/kyc/cases/{id}
│   │       ├── (billing)/
│   │       │   └── wallet/
│   │       │       └── index.tsx                            # Wallet (mobile)
│   │       │                                                # - Balance, transactions
│   │       │                                                # - Request payout
│   │       │                                                # BE: financial-be/wallet|payout
│   │       │                                                # GET  /v1/wallet
│   │       │                                                # GET  /v1/wallet/transactions
│   │       │                                                # POST /v1/payouts
│   │       ├── (inbox)/
│   │       │   ├── messages/
│   │       │   │   ├── [conversationId]/
│   │       │   │   │   └── index.tsx                        # Messages (mobile - thread)
│   │       │   │   │                                        # - View/send messages, attachments
│   │       │   │   │                                        # BE: communications-be/conversation|message
│   │       │   │   │                                        # GET  /v1/conversations/{id}
│   │       │   │   │                                        # GET  /v1/conversations/{id}/messages
│   │       │   │   │                                        # POST /v1/messages
│   │       │   │   └── index.tsx                            # Messages (mobile - inbox)
│   │       │   │                                            # - List user conversations
│   │       │   │                                            # BE: communications-be/conversation
│   │       │   │                                            # GET /v1/conversations?mine=1
│   │       │   └── proposals/
│   │       │       └── index.tsx                            # Proposals inbox (mobile)
│   │       │                                                # - List/detail, statuses
│   │       │                                                # BE: proposals-be/proposal
│   │       │                                                # GET /v1/proposals?mine=1
│   │       ├── (market)/
│   │       │   └── jobs/
│   │       │       └── index.tsx                            # Job discovery (mobile)
│   │       │                                                # - Feed, trending, filters
│   │       │                                                # BE: search-be/feed|trending|query
│   │       │                                                # GET /v1/feed
│   │       │                                                # GET /v1/trending
│   │       ├── (work)/
│   │       │   └── contracts/
│   │       │       ├── [id]/
│   │       │       │   └── work-diary/
│   │       │       │       └── index.tsx                    # Work diary & timesheets (mobile)
│   │       │       │                                        # - Log time/screenshots
│   │       │       │                                        # BE: contracts-be/workdiary|timesheet
│   │       │       │                                        # GET  /v1/contracts/{id}/work-diary
│   │       │       │                                        # POST /v1/contracts/{id}/timesheets
│   │       │       └── [id]/
│   │       │           └── milestones/
│   │       │               └── index.tsx                    # Milestones (mobile)
│   │       │                                                # - View progress, submit for review
│   │       │                                                # BE: contracts-be/milestone
│   │       │                                                # GET  /v1/contracts/{id}/milestones
│   │       ├── (contracts)/
│   │       │   └── offers/
│   │       │       └── [id]/
│   │       │           └── index.tsx                        # Offer (mobile - accept/decline)
│   │       │                                                # - Review terms, accept/decline
│   │       │                                                # BE: contracts-be/offer (proposed if missing)
│   │       │                                                # GET  /v1/offers/{id}
│   │       │                                                # POST /v1/offers/{id}/accept|decline
│   │       ├── billing/
│   │       │   └── subscriptions/
│   │       │       └── index.tsx                            # Subscriptions (mobile)
│   │       │                                                # - View current plan and status
│   │       │                                                # - Upgrade/downgrade/cancel (with proration)
│   │       │                                                # - View invoices/history
│   │       │                                                # BE: financial-be/subscription
│   │       │                                                # GET  /v1/subscriptions/me
│   │       │                                                # POST /v1/subscriptions/change
│   │       │                                                # POST /v1/subscriptions/cancel
│   │       │                                                # GET  /v1/subscriptions/{id}/invoices
│   │       ├── notifications/
│   │       │   └── index.tsx                                # Notifications center (mobile)
│   │       │                                                # - In-app alerts list, mark read
│   │       │                                                # BE: communications-be/notification
│   │       │                                                # GET  /v1/notifications?mine=1
│   │       │                                                # POST /v1/notifications/{id}/read
│   │       ├── search/
│   │       │   ├── alerts/
│   │       │   │   └── index.tsx                            # Saved search alerts (mobile)
│   │       │   │                                            # - Create/manage alerts
│   │       │   │                                            # BE: search-be/alert (proposed if missing)
│   │       │   │                                            # GET/POST/DELETE /v1/search/alerts
│   │       │   └── saved/
│   │       │       └── index.tsx                            # Saved searches (mobile)
│   │       │                                                # - List/run/delete saved queries
│   │       │                                                # BE: search-be/saved_query (proposed if missing)
│   │       │                                                # GET/POST/DELETE /v1/search/saved
│   │       ├── (billing)/
│   │       │   └── payout-methods/
│   │       │       └── index.tsx                            # Payout methods (mobile)
│   │       │                                                # - Add/edit bank/card
│   │       │                                                # BE: financial-be/payout_method
│   │       │                                                # GET/POST/DELETE /v1/payout-methods
│   │       ├── settings/
│   │       │   └── privacy/
│   │       │       └── gdpr/
│   │       │           └── index.tsx                        # GDPR (mobile)
│   │       │                                                # - Request data export
│   │       │                                                # - Request account erasure
│   │       │                                                # - Track request status
│   │       │                                                # BE: users-be/privacy
│   │       │                                                # POST /v1/privacy/export
│   │       │                                                # POST /v1/privacy/erase
│   │       │                                                # GET  /v1/privacy/requests/{id}
│   │       └── support/
│   │           └── tickets/
│   │               └── index.tsx                            # Support tickets (mobile)
│   │                                                        # - Create ticket, reply with attachments
│   │                                                        # - View ticket list and details
│   │                                                        # BE: admin-be/support_ticket
│   │                                                        # GET  /v1/support/tickets (mine)
│   │                                                        # POST /v1/support/tickets
│   │                                                        # POST /v1/support/tickets/{id}/messages
│   │                                                        # BE: storage-be/asset (uploads)
│   │                                                        # POST /v1/storage/uploads (signed URL) → PUT file → POST /v1/storage/commit
│   └── web/
│       └── src/
│           ├── app/
│           │   └── [locale]/
│           │       ├── (admin)/
│           │       │   ├── business-verification/
│           │       │   │   └── page.tsx                     # Business Verification cases
│           │       │   │                                   # - Queue, details, evidence review
│           │       │   │                                   # - Approve/Reject with notes
│           │       │   │                                   # BE: admin-be/business_verification
│           │       │   │                                   # GET  /v1/admin/business-verifications
│           │       │   │                                   # POST /v1/admin/business-verifications/{id}/approve
│           │       │   │                                   # POST /v1/admin/business-verifications/{id}/reject
│           │       │   ├── kyc-cases/
│           │       │   │   └── page.tsx                     # KYC case management
│           │       │   │                                   # - Triage, document review
│           │       │   │                                   # - Approve/Reject/Request-info
│           │       │   │                                   # BE: admin-be/kyc_case
│           │       │   │                                   # GET  /v1/admin/kyc/cases
│           │       │   │                                   # POST /v1/admin/kyc/cases/{id}/approve
│           │       │   │                                   # POST /v1/admin/kyc/cases/{id}/reject
│           │       │   │                                   # POST /v1/admin/kyc/cases/{id}/request-info
│           │       │   ├── ops/
│           │       │   │   ├── admin-session/
│           │       │   │   │   └── page.tsx                # JIT “break-glass” admin session
│           │       │   │   │                               # - Request time-boxed access
│           │       │   │   │                               # - Two-person approval
│           │       │   │   │                               # - Full audit trail
│           │       │   │   │                               # BE: admin-be/admin_session
│           │       │   │   │                               # POST /v1/admin/sessions {reason,duration}
│           │       │   │   │                               # GET  /v1/admin/sessions
│           │       │   │   │                               # POST /v1/admin/sessions/{id}/approve
│           │       │   │   └── change-approval/
│           │       │   │       └── page.tsx                # Two-person change approvals
│           │       │   │                                   # - Risky change queue
│           │       │   │                                   # - Approve / Rollback
│           │       │   │                                   # BE: admin-be/change_approval
│           │       │   │                                   # GET  /v1/admin/change-approvals
│           │       │   │                                   # POST /v1/admin/change-approvals/{id}/approve
│           │       │   │                                   # POST /v1/admin/change-approvals/{id}/rollback
│           │       │   ├── refund-cases/
│           │       │   │   └── page.tsx                    # Refund & goodwill credits
│           │       │   │                                   # - Intake, investigation, decision
│           │       │   │                                   # BE: admin-be/refund_case | admin-be/goodwill_credit
│           │       │   │                                   # GET  /v1/admin/refunds
│           │       │   │                                   # POST /v1/admin/refunds/{id}/approve
│           │       │   │                                   # POST /v1/admin/refunds/{id}/deny
│           │       │   ├── moderation/
│           │       │   │   ├── actions/
│           │       │   │   │   └── page.tsx                # Enforcement actions
│           │       │   │   │                               # - Warn / Suspend / Ban
│           │       │   │   │                               # BE: users-be/moderation
│           │       │   │   │                               # POST /v1/admin/warning
│           │       │   │   │                               # POST /v1/admin/suspension
│           │       │   │   │                               # POST /v1/admin/ban
│           │       │   │   ├── appeals/
│           │       │   │   │   └── page.tsx                # Appeals review
│           │       │   │   │                               # - View appeals, decide outcomes
│           │       │   │   │                               # BE: users-be/appeal
│           │       │   │   │                               # GET  /v1/admin/appeals
│           │       │   │   │                               # POST /v1/admin/appeals/{id}/decide
│           │       │   │   └── review-queue/
│           │       │   │       └── page.tsx                # Moderation queue
│           │       │   │                                   # - Content & report review
│           │       │   │                                   # BE: users-be/moderation
│           │       │   │                                   # GET /v1/admin/moderation/queue
│           │       │   ├── search-quality/
│           │       │   │   ├── facets-filters/
│           │       │   │   │   └── page.tsx                # Facets & filters config
│           │       │   │   │                               # - Define schemas & options
│           │       │   │   │                               # BE: search-be/facets | search-be/filters
│           │       │   │   │                               # GET /v1/search/admin/facets
│           │       │   │   │                               # PUT /v1/search/admin/facets
│           │       │   │   │                               # GET /v1/search/admin/filters
│           │       │   │   │                               # PUT /v1/search/admin/filters
│           │       │   │   ├── indexing/
│           │       │   │   │   └── page.tsx                # Indexing & backfills
│           │       │   │   │                               # - Trigger reindex/backfills
│           │       │   │   │                               # - Monitor progress/errors
│           │       │   │   │                               # BE: search-be/index
│           │       │   │   │                               # POST /v1/search/admin/reindex
│           │       │   │   │                               # GET  /v1/search/admin/index/jobs
│           │       │   │   ├── language/
│           │       │   │   │   └── page.tsx                # Language analyzers
│           │       │   │   │                               # - Stopwords, stemming, per-locale
│           │       │   │   │                               # BE: search-be/language
│           │       │   │   │                               # GET /v1/search/admin/languages
│           │       │   │   │                               # PUT /v1/search/admin/languages
│           │       │   │   ├── metrics/
│           │       │   │   │   └── page.tsx                # Search metrics
│           │       │   │   │                               # - Query logs, CTR, latency, drift
│           │       │   │   │                               # BE: search-be/metrics
│           │       │   │   │                               # GET /v1/search/admin/metrics?range=…
│           │       │   │   ├── rewrites/
│           │       │   │   │   └── page.tsx                # Query rewrites & synonyms
│           │       │   │   │                               # - Edit dictionaries
│           │       │   │   │                               # - A/B preview before publish
│           │       │   │   │                               # BE: search-be/rewrite
│           │       │   │   │                               # GET/PUT /v1/search/admin/rewrites
│           │       │   │   │                               # POST    /v1/search/admin/preview
│           │       │   │   └── speller/
│           │       │   │       └── page.tsx                # Speller configuration
│           │       │   │                                   # - Dictionaries & thresholds
│           │       │   │                                   # BE: search-be/speller
│           │       │   │                                   # GET/PUT /v1/search/admin/speller
│           │       │   ├── communications/
│           │       │   │   ├── broadcasts/
│           │       │   │   │   └── page.tsx                # Broadcast messages
│           │       │   │   │                               # - Create/send announcements
│           │       │   │   │                               # - Rate limits & compliance
│           │       │   │   │                               # BE: communications-be/broadcast
│           │       │   │   │                               # POST /v1/broadcasts
│           │       │   │   │                               # GET  /v1/broadcasts
│           │       │   │   ├── templates/
│           │       │   │   │   └── page.tsx                # Templates management
│           │       │   │   │                               # - Email/SMS/in-app templates
│           │       │   │   │                               # BE: communications-be/template
│           │       │   │   │                               # GET/POST/PUT/DELETE /v1/templates
│           │       │   │   └── campaigns/
│           │       │   │       └── page.tsx                # Multi-step campaigns
│           │       │   │                                   # - Audience + schedule + analytics
│           │       │   │                                   # BE: communications-be/campaign
│           │       │   │                                   # GET/POST/PUT/DELETE /v1/campaigns
│           │       │   ├── audit-reporting/
│           │       │   │   └── page.tsx                    # Audit & reporting
│           │       │   │                                   # - System logs, CSV/BI exports
│           │       │   │                                   # BE: utility/audit | financial-be/reports
│           │       │   │                                   # GET /v1/admin/audit
│           │       │   │                                   # GET /v1/admin/reports?type=…
│           │       │   ├── status-incidents/
│           │       │   │   └── page.tsx                    # Incidents & maintenance
│           │       │   │                                   # - Open/resolve incidents
│           │       │   │                                   # - Post maintenance notes
│           │       │   │                                   # BE: utility/status | communications-be/broadcast
│           │       │   │                                   # GET  /v1/status
│           │       │   │                                   # POST /v1/broadcasts
│           │       │   ├── storage/
│           │       │   │   └── page.tsx                    # Storage lifecycle (admin)
│           │       │   │                                   # - Retention, soft-delete, restore
│           │       │   │                                   # BE: storage-be/lifecycle
│           │       │   │                                   # GET /v1/storage/lifecycle
│           │       │   ├── financial/
│           │       │   │   └── tax/
│           │       │   │       └── page.tsx                # Tax ops (admin)
│           │       │   │                                   # - Forms queues, reports
│           │       │   │                                   # BE: financial-be/tax
│           │       │   │                                   # GET /v1/tax/forms
│           │       │   │                                   # GET /v1/tax/reports
│           │       │   └── search/
│           │       │       └── explain/
│           │       │           └── page.tsx                # Search explainability (admin)
│           │       │                                       # - "Why this result" tooling
│           │       │                                       # BE: search-be/explainability
│           │       │                                       # GET /v1/explain/{docId}
│           │       ├── notifications/
│           │       │   └── page.tsx                        # Notifications center (web)
│           │       │                                      # - In-app notifications, bulk mark read
│           │       │                                      # BE: communications-be/notification
│           │       │                                      # GET  /v1/notifications?mine=1
│           │       │                                      # POST /v1/notifications/{id}/read
│           │       ├── search/
│           │       │   ├── assist/
│           │       │   │   └── page.tsx                    # Search assist
│           │       │   │                                   # - Suggestions, speller, rewrites, languages
│           │       │   │                                   # BE: search-be/suggestions|speller|rewrites|languages
│           │       │   │                                   # GET /v1/suggestions
│           │       │   │                                   # GET /v1/speller
│           │       │   │                                   # GET /v1/rewrites
│           │       │   │                                   # GET /v1/languages
│           │       │   ├── feed/
│           │       │   │   └── page.tsx                    # Personalized feed
│           │       │   │                                   # - Jobs/talent recommendations
│           │       │   │                                   # BE: search-be/feed|similarity|trending
│           │       │   │                                   # GET /v1/feed
│           │       │   │                                   # GET /v1/trending
│           │       │   ├── portfolios/
│           │       │   │   └── page.tsx                    # Portfolio search
│           │       │   │                                   # - Filter by skills/tags
│           │       │   │                                   # BE: search-be/portfolio_index
│           │       │   │                                   # GET /v1/search/portfolios
│           │       │   ├── saved/
│           │       │   │   └── page.tsx                    # Saved searches (web)
│           │       │   │                                   # - List/create/delete saved queries
│           │       │   │                                   # BE: search-be/saved_query (proposed if missing)
│           │       │   │                                   # GET/POST/DELETE /v1/search/saved
│           │       │   └── alerts/
│           │       │       └── page.tsx                    # Search alerts (web)
│           │       │                                       # - Manage email/push alerts for queries
│           │       │                                       # BE: search-be/alert (proposed if missing)
│           │       │                                       # GET/POST/DELETE /v1/search/alerts
│           │       ├── (billing)/
│           │       │   ├── invoices/
│           │       │   │   └── page.tsx                    # Invoices
│           │       │   │                                   # - List/pay/download invoices
│           │       │   │                                   # BE: financial-be/invoice|payment
│           │       │   │                                   # GET  /v1/invoices
│           │       │   │                                   # POST /v1/invoices/{id}/pay
│           │       │   ├── payouts/
│           │       │   │   └── page.tsx                    # Payouts
│           │       │   │                                   # - Destinations & requests
│           │       │   │                                   # BE: financial-be/payout
│           │       │   │                                   # GET/POST /v1/payouts
│           │       │   ├── tax-forms/
│           │       │   │   └── page.tsx                    # Tax forms (freelancer/client)
│           │       │   │                                   # - W-8/W-9/VAT info, upload forms
│           │       │   │                                   # BE: financial-be/tax_form (proposed if missing)
│           │       │   │                                   # GET/POST /v1/tax/forms (mine)
│           │       │   ├── reports/
│           │       │   │   └── page.tsx                    # Billing reports (client)
│           │       │   │                                   # - Spend by team/project, exports
│           │       │   │                                   # BE: financial-be/reports
│           │       │   │                                   # GET /v1/reports?scope=client
│           │       │   └── wallet/
│           │       │       └── page.tsx                    # Wallet
│           │       │                                       # - Balance & transactions
│           │       │                                       # BE: financial-be/wallet
│           │       │                                       # GET /v1/wallet
│           │       ├── (client)/
│           │       │   ├── billing/
│           │       │   │   └── profile/
│           │       │   │       └── page.tsx                # Billing profile (client)
│           │       │   │                                   # - Payment methods, tax/VAT
│           │       │   │                                   # BE: users-be/financial_profile|payment_method
│           │       │   │                                   # GET/PUT /v1/financial-profile
│           │       │   │                                   # CRUD    /v1/payment-methods
│           │       │   ├── jobs/
│           │       │   │   ├── manage/
│           │       │   │   │   └── page.tsx                # Job management
│           │       │   │   │                               # - Eligibility, budget, visibility
│           │       │   │   │                               # BE: jobs-be/job (+ core handlers)
│           │       │   │   │                               # GET/PATCH /v1/jobs/{id}/eligibility|budget|visibility
│           │       │   │   └── post/
│           │       │   │       └── page.tsx                # Post a job
│           │       │   │                                   # - Draft, attachments, publish
│           │       │   │                                   # BE: jobs-be/job | storage-be/asset
│           │       │   │                                   # POST /v1/jobs
│           │       │   │                                   # PATCH /v1/jobs/{id}
│           │       │   │                                   # POST /v1/jobs/{id}/publish
│           │       │   ├── offers/
│           │       │   │   └── new/
│           │       │   │       └── [proposalId]/
│           │       │   │           └── page.tsx            # Create offer from proposal
│           │       │   │                                   # - Terms, milestones, start date
│           │       │   │                                   # BE: contracts-be/offer (proposed if missing)
│           │       │   │                                   # POST /v1/offers
│           │       │   ├── org/
│           │       │   │   └── page.tsx                    # Organization & teams
│           │       │   │                                   # - Members, roles, invites
│           │       │   │                                   # BE: users-be/org|team|invite
│           │       │   │                                   # GET/POST/PATCH /v1/orgs, /v1/orgs/{id}/members, /v1/invites
│           │       │   ├── shortlists/
│           │       │   │   └── page.tsx                    # Talent shortlists
│           │       │   │                                   # - Create/manage candidate lists
│           │       │   │                                   # BE: proposals-be/shortlist (proposed)
│           │       │   │                                   # GET/POST/DELETE /v1/shortlists
│           │       │   └── talent/
│           │       │       └── invite/
│           │       │           └── [userId]/
│           │       │               └── page.tsx            # Invite freelancer to apply
│           │       │                                       # - Send invite, track status
│           │       │                                       # BE: proposals-be/invite
│           │       │                                       # POST /v1/invites
│           │       ├── (contracts)/
│           │       │   ├── contracts/
│           │       │   │   └── [id]/
│           │       │   │       └── page.tsx                # Contract detail & SOW
│           │       │   │                                   # - Sign/accept, milestones
│           │       │   │                                   # BE: contracts-be/contract|sow
│           │       │   │                                   # GET /v1/contracts/{id}
│           │       │   │                                   # POST /v1/contracts/{id}/sign
│           │       │   ├── escrow/
│           │       │   │   └── [contractId]/
│           │       │   │       ├── fund/
│           │       │   │       │   └── page.tsx            # Fund escrow (client)
│           │       │   │       │                           # - Amount, payment method
│           │       │   │       │                           # BE: financial-be/escrow
│           │       │   │       │                           # POST /v1/escrow/{contractId}/fund
│           │       │   │       └── release/
│           │       │   │           └── page.tsx            # Release escrow (client)
│           │       │   │                                   # - Partial/full release
│           │       │   │                                   # BE: financial-be/escrow
│           │       │   │                                   # POST /v1/escrow/{contractId}/release
│           │       │   └── workroom/
│           │       │       └── [id]/
│           │       │           ├── approvals/
│           │       │           │   └── page.tsx            # Timesheet approvals (client)
│           │       │           │                           # - Approve/reject, comments
│           │       │           │                           # BE: contracts-be/timesheet_approval (proposed if missing)
│           │       │           │                           # POST /v1/contracts/{id}/timesheets/{tsId}/approve|reject
│           │       │           ├── files/
│           │       │           │   └── page.tsx            # Workroom files tab
│           │       │           │                           # - Shared files, versions
│           │       │           │                           # BE: storage-be/asset
│           │       │           │                           # GET /v1/contracts/{id}/files
│           │       │           └── milestones/
│           │       │               └── page.tsx            # Milestones (web)
│           │       │                                       # - Create/edit/submit, approvals
│           │       │                                       # BE: contracts-be/milestone
│           │       │                                       # GET/POST/PATCH /v1/contracts/{id}/milestones
│           │       ├── (delivery)/
│           │       │   └── deliverables/
│           │       │       └── [contractId]/
│           │       │           └── page.tsx                # Deliverables & versions
│           │       │                                       # - Upload, request changes
│           │       │                                       # BE: contracts-be/deliverable | storage-be/asset
│           │       │                                       # GET/POST /v1/contracts/{id}/deliverables
│           │       ├── (reviews)/
│           │       │   └── reviews/
│           │       │       └── new/
│           │       │           └── [contractId]/
│           │       │               └── page.tsx            # Write a review
│           │       │                                       # - Post-review, edit within window
│           │       │                                       # BE: reviews-be/review
│           │       │                                       # POST /v1/reviews
│           │       ├── (support)/
│           │       │   └── disputes/
│           │       │       └── [contractId]/
│           │       │           └── page.tsx                # Dispute center
│           │       │                                       # - Open case, attach evidence
│           │       │                                       # BE: contracts-be/dispute | financial-be/refund | admin-be/refund_case
│           │       │                                       # POST /v1/contracts/{id}/disputes
│           │       ├── (dashboard)/
│           │       │   ├── billing/
│           │       │   │   ├── budgets/
│           │       │   │   │   └── page.tsx                # Budgets
│           │       │   │   │                               # - Create/manage budgets & alerts
│           │       │   │   │                               # BE: financial-be/budget
│           │       │   │   │                               # GET  /v1/budgets
│           │       │   │   │                               # POST /v1/budgets
│           │       │   │   └── subscriptions/
│           │       │   │       ├── history/
│           │       │   │       │   └── page.tsx            # Subscription history
│           │       │   │       │                           # - Invoices & entitlement history
│           │       │   │       │                           # BE: financial-be/subscription
│           │       │   │       │                           # GET /v1/subscriptions/{id}/invoices
│           │       │   │       └── manage/
│           │       │   │           └── page.tsx            # Manage subscription
│           │       │   │                                   # - Plan selection, upgrade/downgrade
│           │       │   │                                   # - Cancel & dunning status
│           │       │   │                                   # BE: financial-be/subscription
│           │       │   │                                   # GET  /v1/subscriptions/me
│           │       │   │                                   # POST /v1/subscriptions/change
│           │       │   │                                   # POST /v1/subscriptions/cancel
│           │       │   ├── messages/
│           │       │   │   ├── [conversationId]/
│           │       │   │   │   └── page.tsx                # Conversation
│           │       │   │   │                               # - Thread, read receipts, uploads
│           │       │   │   │                               # BE: communications-be/conversation | message
│           │       │   │   │                               # GET  /v1/conversations/{id}
│           │       │   │   │                               # GET  /v1/conversations/{id}/messages
│           │       │   │   │                               # POST /v1/messages
│           │       │   │   ├── page.tsx                    # Inbox
│           │       │   │   │                               # - User conversations list
│           │       │   │   │                               # BE: communications-be/conversation
│           │       │   │   │                               # GET /v1/conversations?mine=1
│           │       │   │   └── settings/
│           │       │   │       └── notifications/
│           │       │   │           └── page.tsx            # Notification preferences
│           │       │   │                                   # - Email/SMS/push, digests
│           │       │   │                                   # BE: communications-be/notification
│           │       │   │                                   # GET/PUT /v1/notifications/preferences
│           │       │   ├── profile/
│           │       │   │   └── reviews/
│           │       │   │       └── page.tsx                # Reviews (from profile)
│           │       │   │                                   # - View received/given
│           │       │   │                                   # - Submit review
│           │       │   │                                   # BE: reviews-be/review
│           │       │   │                                   # GET  /v1/reviews?subject_id=…
│           │       │   │                                   # POST /v1/reviews
│           │       │   └── settings/
│           │       │       └── privacy/
│           │       │           └── gdpr/
│           │       │               ├── delete/
│           │       │               │   └── page.tsx        # GDPR — Delete account
│           │       │               │                       # - Submit erasure request
│           │       │               │                       # - Track status
│           │       │               │                       # BE: users-be/privacy
│           │       │               │                       # POST /v1/privacy/erase
│           │       │               │                       # GET  /v1/privacy/requests/{id}
│           │       │               └── export/
│           │       │                   └── page.tsx        # GDPR — Data export
│           │       │                                       # - Request export
│           │       │                                       # - Track status & download link
│           │       │                                       # BE: users-be/privacy
│           │       │                                       # POST /v1/privacy/export
│           │       │                                       # GET  /v1/privacy/requests/{id}
│           │       ├── (freelancer)/
│           │       │   ├── portfolio/
│           │       │   │   └── manage/
│           │       │   │       └── page.tsx                # Manage portfolio
│           │       │   │                                   # - Add/edit items, media uploads
│           │       │   │                                   # BE: users-be/portfolio | storage-be/asset
│           │       │   │                                   # GET/POST/PATCH /v1/portfolio
│           │       │   └── proposals/
│           │       │       └── new/
│           │       │           └── page.tsx                # New proposal (freelancer)
│           │       │                                       # - Compose, attach files
│           │       │                                       # BE: proposals-be/proposal | storage-be/asset
│           │       │                                       # POST /v1/proposals
│           │       ├── (social)/
│           │       │   └── network/
│           │       │       └── page.tsx                    # Network & groups
│           │       │                                       # - Connections, groups, referrals
│           │       │                                       # BE: users-be/professional_network|user_group|referral
│           │       │                                       # GET/POST /v1/network|/v1/groups|/v1/referrals
│           │       ├── pricing/
│           │       │   └── page.tsx                        # Pricing (public)
│           │       │                                      # - Plans overview, compare
│           │       │                                      # BE: financial-be/subscription (plans)
│           │       │                                      # GET /v1/subscriptions/plans
│           │       ├── settings/
│           │       │   ├── account/
│           │       │   │   └── saved-items/
│           │       │   │       └── page.tsx                # Saved items
│           │       │   │                                   # - Saved jobs/talent, remove
│           │       │   │                                   # BE: users-be/saved_item
│           │       │   │                                   # GET/DELETE /v1/saved-items
│           │       │   └── security/
│           │       │       ├── devices/
│           │       │       │   └── page.tsx                # Devices & sessions
│           │       │       │                               # - Active sessions, revoke device
│           │       │       │                               # BE: users-be/session
│           │       │       │                               # GET/DELETE /v1/sessions
│           │       │       ├── mfa/
│           │       │       │   └── page.tsx                # MFA settings
│           │       │       │                               # - Enroll/disable TOTP/SMS
│           │       │       │                               # BE: users-be/mfa
│           │       │       │                               # GET/POST/DELETE /v1/mfa
│           │       │       └── password/
│           │       │           └── page.tsx                # Change password
│           │       │                                       # - Update password
│           │       │                                       # BE: users-be/account
│           │       │                                       # POST /v1/account/password
│           │       └── status/
│           │           └── page.tsx                        # Status page (public)
│           │                                              # - Current incidents & components
│           │                                              # BE: utility/status
│           │                                              # GET /v1/status
│           └── features/
│               ├── budget/
│               │   └── api/
│               │       └── mutations.ts                   # Budget mutations
│               │                                          # - Create/update budgets
│               │                                          # - Invalidate ['budget','list']
│               │                                          # BE: financial-be/budget
│               │                                          # POST /v1/budgets
│               ├── interviews/
│               │   └── api/
│               │       └── mutations.ts                   # Interview scheduling
│               │                                          # - Create/reschedule/cancel
│               │                                          # BE: proposals-be/interview
│               │                                          # POST /v1/proposals/{id}/interviews
│               ├── moderation/
│               │   └── api/
│               │       └── mutations.ts                   # Moderation actions
│               │                                          # - Warn / Suspend / Ban users
│               │                                          # BE: users-be/moderation
│               │                                          # POST /v1/admin/warning
│               │                                          # POST /v1/admin/suspension
│               │                                          # POST /v1/admin/ban
│               ├── notifications/
│               │   └── api/
│               │       ├── mutations.ts                   # Notifications mutations
│               │       │                                  # - Mark read / mark all
│               │       │                                  # BE: communications-be/notification
│               │       │                                  # POST /v1/notifications/{id}/read
│               │       └── queries.ts                     # Notifications queries
│               │                                          # - List notifications
│               │                                          # BE: communications-be/notification
│               │                                          # GET /v1/notifications?mine=1
│               ├── search/
│               │   └── api/
│               │       ├── alerts.ts                      # Search alerts API
│               │       │                                  # BE: search-be/alert (proposed if missing)
│               │       │                                  # GET/POST/DELETE /v1/search/alerts
│               │       └── saved.ts                       # Saved queries API
│               │                                          # BE: search-be/saved_query (proposed if missing)
│               │                                          # GET/POST/DELETE /v1/search/saved
│               ├── subscriptions/
│               │   └── api/
│               │       ├── mutations.ts                   # Subscription mutations
│               │       │                                  # - Change plan / Cancel
│               │       │                                  # - Invalidate ['subscriptions','me']
│               │       │                                  # BE: financial-be/subscription
│               │       │                                  # POST /v1/subscriptions/change
│               │       │                                  # POST /v1/subscriptions/cancel
│               │       └── queries.ts                     # Subscription queries
│               │                                          # - Current subscription
│               │                                          # - Invoices history
│               │                                          # BE: financial-be/subscription
│               │                                          # GET /v1/subscriptions/me
│               │                                          # GET /v1/subscriptions/{id}/invoices
│               ├── support/
│               │   └── api/
│               │       ├── mutations.ts                   # Support ticket mutations
│               │       │                                  # - Create ticket, add message, attach files
│               │       │                                  # - Invalidate ['support','tickets']
│               │       │                                  # BE: admin-be/support_ticket
│               │       │                                  # POST /v1/support/tickets
│               │       │                                  # POST /v1/support/tickets/{id}/messages
│               │       │                                  # BE: storage-be/asset (uploads)
│               │       │                                  # POST /v1/storage/uploads (signed URL) → PUT file → POST /v1/storage/commit
│               │       └── queries.ts                     # Support ticket queries
│               │                                          # - List & detail
│               │                                          # BE: admin-be/support_ticket
│               │                                          # GET /v1/support/tickets
│               │                                          # GET /v1/support/tickets/{id}
│               ├── offers/
│               │   └── api/
│               │       └── mutations.ts                   # Offer create/respond
│               │                                          # - Create, accept, decline
│               │                                          # BE: contracts-be/offer (proposed if missing)
│               │                                          # POST /v1/offers; POST /v1/offers/{id}/accept|decline
│               └── reviews/
│                   ├── api/
│                   │   ├── mutations.ts                   # Review mutations
│                   │   │                                  # - Submit / edit review
│                   │   │                                  # - Invalidate ['reviews','list',subjectId]
│                   │   │                                  # BE: reviews-be/review
│                   │   │                                  # POST /v1/reviews
│                   │   └── queries.ts                     # Review queries
│                   │                                      # - List reviews for subject
│                   │                                      # BE: reviews-be/review
│                   │                                      # GET /v1/reviews?subject_id=…
│                   └── components/
│                       └── ReviewForm.tsx                 # Review form (presentational)
│                                                          # - Consumes typed props only
│                                                          # BE: none (presentational)
├── packages/
│   ├── shared/
│   │   └── src/
│   │       ├── api/
│   │       │   ├── communications/
│   │       │   │   ├── client.ts                           # Communications API client
│   │       │   │   # - Conversations, messages, notifications
│   │       │   │   # BE: communications-be
│   │       │   │   ├── conversations.ts                    # /v1/conversations, /v1/messages
│   │       │   │   └── notifications.ts                    # /v1/notifications
│   │       │   ├── contracts/
│   │       │   │   ├── client.ts                           # Contracts API client
│   │       │   │   # - Contract, SOW, milestones, diary, deliverables, disputes, escrow, offers
│   │       │   │   # BE: contracts-be | financial-be/escrow
│   │       │   │   ├── contracts.ts                        # /v1/contracts
│   │       │   │   ├── milestones.ts                       # /v1/contracts/{id}/milestones
│   │       │   │   ├── diary.ts                            # /v1/contracts/{id}/work-diary|timesheets
│   │       │   │   ├── deliverables.ts                     # /v1/contracts/{id}/deliverables
│   │       │   │   ├── disputes.ts                         # /v1/contracts/{id}/disputes
│   │       │   │   ├── escrow.ts                           # /v1/escrow/{contractId}/fund|release
│   │       │   │   └── offers.ts                           # /v1/offers (proposed)
│   │       │   ├── financial/
│   │       │   │   ├── client.ts                           # Financial API client
│   │       │   │   # - Base URL, interceptors, retries
│   │       │   │   # BE: financial-be (all)
│   │       │   │   # Propagates trace & idempotency keys
│   │       │   │   ├── invoices.ts                         # Invoices endpoints
│   │       │   │   # - GET /v1/invoices; POST /v1/invoices/{id}/pay
│   │       │   │   ├── payouts.ts                          # Payouts endpoints
│   │       │   │   # - GET/POST /v1/payouts
│   │       │   │   ├── refunds.ts                          # Refunds endpoints
│   │       │   │   # - GET/POST /v1/refunds
│   │       │   │   ├── subscription.ts                     # Subscription endpoints
│   │       │   │   # - GET /v1/subscriptions/me; POST /v1/subscriptions/change|cancel
│   │       │   │   └── wallet.ts                           # Wallet endpoints
│   │       │   │       # - GET /v1/wallet; GET /v1/wallet/transactions
│   │       │   ├── jobs/
│   │       │   │   ├── client.ts                           # Jobs API client
│   │       │   │   # - Post/edit/publish jobs; invites
│   │       │   │   # BE: jobs-be
│   │       │   │   ├── jobs.ts                              # /v1/jobs + subroutes
│   │       │   │   └── invites.ts                           # /v1/invites
│   │       │   ├── proposals/
│   │       │   │   ├── client.ts                           # Proposals API client
│   │       │   │   # - Retries, error mapping
│   │       │   │   # BE: proposals-be
│   │       │   │   ├── interviews.ts                       # Interviews scheduling
│   │       │   │   # - POST /v1/proposals/{id}/interviews
│   │       │   │   └── proposals.ts                        # Proposals CRUD
│   │       │   │       # - POST /v1/proposals; GET /v1/proposals?mine=1
│   │       │   ├── search/
│   │       │   │   ├── client.ts                           # Search API client
│   │       │   │   # - Query/feed/portfolios + assist + saved/alerts
│   │       │   │   # BE: search-be
│   │       │   │   ├── feed.ts                              # /v1/feed|trending|similarity
│   │       │   │   ├── portfolios.ts                        # /v1/search/portfolios
│   │       │   │   ├── query.ts                             # /v1/search
│   │       │   │   ├── alerts.ts                            # /v1/search/alerts (proposed)
│   │       │   │   └── saved.ts                             # /v1/search/saved (proposed)
│   │       │   ├── storage/
│   │       │   │   ├── client.ts                           # Storage API client
│   │       │   │   # - Signed upload, commit, list
│   │       │   │   # BE: storage-be
│   │       │   │   └── assets.ts                            # /v1/storage/uploads|commit|assets
│   │       │   └── users/
│   │       │       ├── client.ts                           # Users API client
│   │       │       # - Org/team, profile, security, saved items
│   │       │       # BE: users-be
│   │       │       ├── org.ts                               # /v1/orgs, /v1/invites
│   │       │       ├── profile.ts                           # /v1/profile, /v1/users/{handle}
│   │       │       ├── security.ts                          # /v1/sessions, /v1/mfa
│   │       │       └── saved.ts                             # /v1/saved-items
│   │       └── features/
│   │           └── notifications/
│   │               └── store.ts                             # Notifications store (Zustand)
│   │                                                      # - Unread counts, last fetched
│   │                                                      # BE: none (consumes API in features/notifications)
│   └── types/
│       └── src/
│           └── entities/
│               ├── contracts/
│               │   ├── dispute.ts                           # Dispute entity types
│               │   ├── milestone.ts                         # Milestone entity types
│               │   ├── offer.ts                             # Offer entity types
│               │   └── timesheet.ts                         # Timesheet entity types
│               ├── financial/
│               │   ├── escrow.ts                            # Escrow DTOs
│               │   ├── report.ts                            # Billing report row DTOs
│               │   └── tax.ts                               # Tax DTOs
│               ├── notification.ts                           # Notification DTOs
│               └── search/
│                   └── saved_query.ts                        # Saved query/alert DTOs
├── packages/
│   └── ui/
│       └── src/
│           └── components/
│               ├── FileUploader/
│               │   ├── FileUploader.native.tsx            # Uploader (native)
│               │   │                                      # - Signed URL upload flow
│               │   │                                      # BE: storage-be/asset
│               │   │                                      # POST /v1/storage/uploads (signed) → PUT file → POST /v1/storage/commit
│               │   ├── FileUploader.tsx                   # Uploader (shared)
│               │   │                                      # - Abstraction wrapper
│               │   │                                      # BE: storage-be/asset (same flow)
│               │   └── FileUploader.web.tsx               # Uploader (web)
│               │                                          # - Drag & drop, previews
│               │                                          # BE: storage-be/asset (same flow)
│               ├── LoadingSpinner/
│               │   ├── LoadingSpinner.native.tsx          # Spinner (native)
│               │   ├── LoadingSpinner.tsx                 # Spinner (shared)
│               │   └── LoadingSpinner.web.tsx             # Spinner (web)
│               │                                          # BE: none (presentational)
│               ├── Modal/
│               │   ├── Modal.native.tsx                   # Modal (native)
│               │   ├── Modal.tsx                          # Modal (shared)
│               │   └── Modal.web.tsx                      # Modal (web)
│               │                                          # BE: none (presentational)
│               └── Toast/
│                   ├── Toast.native.tsx                   # Toasts (native)
│                   ├── Toast.tsx                          # Toasts (shared)
│                   └── Toast.web.tsx                      # Toasts (web)
│                                                          # BE: none (presentational)
```