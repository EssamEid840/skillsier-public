# Skillsier Platform - Complete Updated Folder Structure
## All Microservices with Full Comments (Old + New) ✨

---

## 📋 **Change Impact Summary**

| Area | Impact | Changes |
|------|--------|---------|
| **Auth (pkg/auth)** | 🔴 MAJOR | New root `pkg/auth/` module - all services import this |
| **Platform-shared** | 🔴 MAJOR | New root `platform-shared/` module - utilities, outbox, etc. |
| **Contracts** | 🔴 MAJOR | New root `contracts/events/` - protobuf event schemas |
| **Config Standardization** | 🟡 MEDIUM | Per-service `config/` dir + `internal/config/schema.go` |
| **Dapr Components** | 🟡 MEDIUM | Per-service `dapr/local/` and `dapr/k8s/` split |
| **Migrations** | 🟢 MINOR | Enhanced with version tracking, safety checks, MIGRATIONS.md |
| **K8s Overlays** | 🟢 MINOR | Kustomize structure: `base/`, `overlays/local/`, `overlays/prod/` |

**Legend:**
- 🆕 = New additions
- 📝 = Updated files
- ❌ = Removed (use shared modules)

---

## 🗂️ **ROOT MONOREPO STRUCTURE**

```
skillsier/                                    # Monorepo root
│
├── apps/
│   └── be/
│       ├── users-be/                         # User management service
│       ├── jobs-be/                          # Job posting and management service
│       ├── proposals-be/                     # Freelancer proposal submission service
│       ├── contracts-be/                     # Contract management service
│       ├── financial-be/                     # Payments and wallet management service
│       ├── communications-be/                # Real-time messaging and notifications service
│       ├── storage-be/                       # File storage and management service
│       ├── search-be/                        # Search and discovery service
│       ├── reviews-be/                       # Rating and review service
│       ├── subscriptions-be/                 # Subscription and premium features service
│       └── admin-be/                         # Admin and moderation service
│
├── pkg/                                      # 🆕 MAJOR - Shared packages (leaf dependencies)
│   └── auth/                                 # 🆕 Centralized Keycloak authentication
│       ├── go.mod                            # module: skillsier.dev/pkg/auth
│       ├── go.sum
│       ├── README.md                         # Usage documentation
│       ├── authorizer.go                     # RequireRoles / RequirePermissions middleware (RBAC)
│       ├── config.go                         # Auth configuration struct (Issuer, Audience, JWKS URL)
│       ├── context.go                        # Put/Get Principal on request context
│       ├── principal.go                      # Normalized identity model (UserID, Roles, Permissions)
│       ├── verifier.go                       # TokenVerifier interface (domain depends on this only)
│       ├── errors.go                         # Auth-specific errors (Unauthorized, Forbidden, Invalid Token)
│       └── keycloak/
│           ├── verifier.go                   # Keycloak JWT/JWKS implementation
│           ├── client.go                     # Keycloak admin API client
│           ├── role_mapper.go                # Map Keycloak roles → app permissions
│           └── config.go                     # Keycloak-specific config (Realm, Client ID)
│
├── platform-shared/                          # 🆕 MAJOR - Cross-cutting utilities (leaf module)
│   ├── go.mod                                # module: skillsier.dev/platform-shared
│   ├── go.sum
│   ├── README.md                             # Usage documentation
│   │
│   ├── logging/                              # 🆕 Structured logging
│   │   ├── logger.go                         # Structured logging (zerolog/zap)
│   │   ├── context.go                        # Logger from context
│   │   ├── fields.go                         # Common log fields (RequestID, UserID, TraceID)
│   │   └── config.go                         # Logging configuration (Level, Format, Output)
│   │
│   ├── tracing/                              # 🆕 Distributed tracing
│   │   ├── otel.go                           # OpenTelemetry setup
│   │   ├── tracer.go                         # Tracer helpers
│   │   ├── span.go                           # Span helpers (Start, End, AddEvent)
│   │   └── config.go                         # Tracing configuration (Endpoint, Sampling)
│   │
│   ├── metrics/                              # 🆕 Prometheus metrics
│   │   ├── metrics.go                        # Prometheus metrics helpers
│   │   ├── http.go                           # HTTP metrics (RED: Rate, Errors, Duration)
│   │   ├── collectors.go                     # Custom collectors (DB pool, queue depth)
│   │   └── config.go                         # Metrics configuration (Namespace, Subsystem)
│   │
│   ├── httpx/                                # 🆕 HTTP-agnostic helpers
│   │   ├── errors.go                         # Error envelope (Code, Message, Details)
│   │   ├── pagination.go                     # Pagination helpers (Limit, Offset, Total, HasNext)
│   │   ├── response.go                       # Standard response wrappers (Success, Error)
│   │   └── validator.go                      # Common validation helpers
│   │
│   ├── ginx/                                 # 🆕 Gin-specific middleware
│   │   ├── requestid.go                      # Request ID middleware (X-Request-ID header)
│   │   ├── logging.go                        # Logging middleware (structured logs per request)
│   │   ├── recover.go                        # Panic recovery middleware
│   │   ├── otel.go                           # OpenTelemetry middleware (trace propagation)
│   │   ├── cors.go                           # CORS middleware
│   │   └── timeout.go                        # Request timeout middleware
│   │
│   ├── outbox/                               # 🆕 Outbox pattern (moved from services)
│   │   ├── publisher.go                      # Watermill TX publisher (publish within DB transaction)
│   │   ├── forwarder.go                      # Background processor (polls outbox, publishes to Kafka)
│   │   ├── scheduler.go                      # Scheduling logic (retry, backoff)
│   │   ├── entity.go                         # Outbox event entity (ID, AggregateID, Payload, Status)
│   │   ├── repository.go                     # Outbox repository interface
│   │   ├── postgres/
│   │   │   └── repository.go                 # Postgres implementation
│   │   └── config.go                         # Outbox configuration (PollInterval, BatchSize)
│   │
│   ├── inbox/                                # 🆕 Message deduplication (prevent duplicate processing)
│   │   ├── checker.go                        # Check if message processed (by MessageID)
│   │   ├── marker.go                         # Mark message as processed
│   │   ├── entity.go                         # Inbox entity (MessageID, ProcessedAt, Handler)
│   │   ├── repository.go                     # Inbox repository interface
│   │   └── postgres/
│   │       └── repository.go                 # Postgres implementation
│   │
│   └── idempotency/                          # 🆕 HTTP request deduplication
│       ├── middleware.go                     # Idempotency-Key middleware (Gin)
│       ├── handler.go                        # Idempotency handler logic
│       ├── entity.go                         # Idempotency record (Key, Response, ExpiresAt)
│       ├── repository.go                     # Repository interface
│       └── postgres/
│           └── repository.go                 # Postgres implementation
│
├── contracts/                                # 🆕 MAJOR - Event schemas (versioned, breaking-change protected)
│   └── events/                               # 🆕 Event contracts module
│       ├── go.mod                            # module: skillsier.dev/contracts/events
│       ├── go.sum
│       ├── README.md                         # Usage and versioning guide
│       ├── buf.yaml                          # Buf configuration (linting, breaking change detection)
│       ├── buf.gen.yaml                      # Code generation config (protoc plugins)
│       ├── buf.lock                          # Dependency lock file
│       ├── EVENTS.md                         # 🆕 Event catalog (Topics, Owners, BC Policy, Changelog)
│       │
│       ├── user/v1/                          # User lifecycle events
│       │   ├── user_created.proto            # New user registered
│       │   ├── user_updated.proto            # User profile updated
│       │   ├── user_verified.proto           # Email/identity verified
│       │   ├── user_suspended.proto          # User suspended by admin
│       │   ├── user_banned.proto             # User banned
│       │   ├── freelancer_profile_completed.proto  # Freelancer profile complete
│       │   └── client_profile_completed.proto      # Client profile complete
│       │
│       ├── job/v1/                           # Job lifecycle events
│       │   ├── job_posted.proto              # New job posted
│       │   ├── job_updated.proto             # Job details updated
│       │   ├── job_closed.proto              # Job closed (filled or cancelled)
│       │   ├── job_invitation_sent.proto     # Invitation sent to freelancer
│       │   ├── job_removed.proto             # Job removed by admin
│       │   └── job_flagged.proto             # Job flagged for review
│       │
│       ├── proposal/v1/                      # Proposal and bidding events
│       │   ├── proposal_submitted.proto      # Proposal submitted
│       │   ├── proposal_accepted.proto       # Proposal accepted by client
│       │   ├── proposal_rejected.proto       # Proposal rejected
│       │   ├── proposal_withdrawn.proto      # Proposal withdrawn by freelancer
│       │   ├── bid_placed.proto              # New bid placed
│       │   ├── bid_updated.proto             # Bid amount updated
│       │   ├── outbid_alert.proto            # Freelancer outbid notification
│       │   ├── connect_used.proto            # Connect/credit used
│       │   └── proposal_flagged.proto        # Proposal flagged for review
│       │
│       ├── contract/v1/                      # Contract lifecycle events
│       │   ├── contract_created.proto        # Contract created
│       │   ├── contract_started.proto        # Work started
│       │   ├── contract_paused.proto         # Contract paused
│       │   ├── contract_ended.proto          # Contract completed/terminated
│       │   ├── milestone_created.proto       # Milestone created
│       │   ├── milestone_completed.proto     # Milestone completed
│       │   ├── milestone_approved.proto      # Milestone approved by client
│       │   ├── timesheet_submitted.proto     # Weekly timesheet submitted
│       │   └── dispute_opened.proto          # Dispute opened
│       │
│       ├── payment/v1/                       # Financial events
│       │   ├── payment_processed.proto       # Payment successfully processed
│       │   ├── payment_failed.proto          # Payment failed
│       │   ├── escrow_held.proto             # Funds held in escrow
│       │   ├── escrow_released.proto         # Escrow released to freelancer
│       │   ├── payout_requested.proto        # Payout requested
│       │   ├── payout_processed.proto        # Payout completed
│       │   ├── invoice_generated.proto       # Invoice generated
│       │   └── refund_processed.proto        # Refund issued
│       │
│       ├── review/v1/                        # Review and rating events
│       │   ├── review_submitted.proto        # Review submitted
│       │   ├── review_responded.proto        # Response to review posted
│       │   ├── badge_awarded.proto           # Achievement badge awarded
│       │   ├── reputation_updated.proto      # Reputation score updated
│       │   └── review_flagged.proto          # Review flagged for moderation
│       │
│       ├── subscription/v1/                  # Subscription events
│       │   ├── subscription_created.proto    # New subscription
│       │   ├── subscription_renewed.proto    # Subscription renewed
│       │   ├── subscription_cancelled.proto  # Subscription cancelled
│       │   ├── subscription_expired.proto    # Subscription expired
│       │   ├── connects_purchased.proto      # Connects purchased
│       │   ├── connects_used.proto           # Connect used
│       │   └── usage_limit_reached.proto     # Plan limit reached
│       │
│       ├── message/v1/                       # Communication events
│       │   ├── message_sent.proto            # Chat message sent
│       │   ├── notification_delivered.proto  # Notification delivered
│       │   ├── email_sent.proto              # Email sent
│       │   └── message_flagged.proto         # Message flagged
│       │   └── in_app_notification_sent.proto         # in_app_notification_sent
│       │
│       ├── storage/v1/                       # File storage events
│       │   ├── file_uploaded.proto           # File uploaded
│       │   ├── file_deleted.proto            # File deleted
│       │   ├── media_processed.proto         # Media processing complete
│       │   └── file_flagged.proto            # File flagged for review
│       │
│       ├── search/v1/                        # Search and recommendation events
│       │   ├── job_indexed.proto             # Job added to search index
│       │   ├── user_indexed.proto            # User added to search index
│       │   └── recommendation_generated.proto # Recommendation created
│       │
│       ├── admin/v1/                         # Admin action events
│       │   ├── user_suspended.proto          # Admin suspended user
│       │   ├── user_banned.proto             # Admin banned user
│       │   ├── content_removed.proto         # Admin removed content
│       │   ├── dispute_resolved.proto        # Admin resolved dispute
│       │   ├── flag_reviewed.proto           # Flag reviewed and actioned
│       │   └── announcement_published.proto  # Platform announcement published
│       │
│       └── gen/                              # 🆕 Generated Go code (checked into git or built in CI)
│           └── go/
│               └── skillsier/
│                   ├── user/v1/
│                   │   ├── user_created.pb.go
│                   │   ├── user_updated.pb.go
│                   │   └── ...
│                   ├── job/v1/
│                   ├── proposal/v1/
│                   ├── contract/v1/
│                   ├── payment/v1/
│                   ├── review/v1/
│                   ├── subscription/v1/
│                   ├── message/v1/
│                   ├── storage/v1/
│                   ├── search/v1/
│                   └── admin/v1/
│
├── deployments/                              # 📝 UPDATED - Kustomize-based deployment
│   └── k8s/
│       ├── base/                             # 🆕 Base Kubernetes manifests (shared)
│       │   ├── kustomization.yaml            # Kustomize base config
│       │   ├── namespace.yaml                # Namespace definition
│       │   ├── networkpolicy.yaml            # 🆕 Default network policies (deny all, allow ingress)
│       │   └── podsecurity.yaml              # 🆕 Pod security standards (restricted)
│       │
│       ├── overlays/                         # 🆕 Environment-specific customizations
│       │   ├── local/                        # Local development environment
│       │   │   ├── kustomization.yaml        # Local overlay config
│       │   │   ├── configmap-patch.yaml      # Local config overrides (debug mode, local URLs)
│       │   │   └── resources-patch.yaml      # Lower resource limits for local (256MB RAM)
│       │   │
│       │   └── prod/                         # Production environment
│       │       ├── kustomization.yaml        # Production overlay config
│       │       ├── configmap-patch.yaml      # Production configs (prod URLs, log levels)
│       │       ├── resources-patch.yaml      # Production resource limits (2GB RAM, 1 CPU)
│       │       ├── replicas-patch.yaml       # Replica counts (min 2, max 10)
│       │       ├── hpa.yaml                  # Horizontal Pod Autoscaler (CPU/Memory based)
│       │       ├── pdb.yaml                  # Pod Disruption Budget (maxUnavailable: 1)
│       │       └── ingress.yaml              # Production ingress (TLS, rate limiting)
│       │
│       └── components/                       # 🆕 Reusable Kustomize components
│           ├── monitoring/
│           │   ├── kustomization.yaml
│           │   └── servicemonitor.yaml       # Prometheus ServiceMonitor
│           └── secrets/
│               ├── kustomization.yaml
│               └── sealed-secrets.yaml       # SealedSecrets for secret management
│
├── docs/                                     # 🆕 Root-level documentation
│   ├── README.md                             # Platform overview
│   ├── ARCHITECTURE.md                       # 🆕 High-level architecture (services, data flow, patterns)
│   ├── SAGAS.md                              # 🆕 Saga patterns & compensations (workflows, rollback)
│   ├── RUNBOOK.md                            # 🆕 Operational procedures (restarts, debugging, incidents)
│   └── decisions/                            # 🆕 Architecture Decision Records (ADRs)
│       ├── README.md                         # ADR index
│       ├── 001-auth-centralization.md        # Why centralize auth in pkg/auth
│       ├── 002-outbox-pattern.md             # Why use outbox for reliable events
│       ├── 003-event-contracts.md            # Why separate event contracts module
│       ├── 004-dapr-adoption.md              # Why use Dapr for service mesh
│       ├── 005-platform-shared.md            # Why create platform-shared module
│       └── 006-config-standardization.md     # Why standardize config across services
│
├── scripts/                                  # 🆕 Root-level automation scripts
│   ├── setup-dev-env.sh                      # Setup entire dev environment (deps, DBs, Kafka)
│   ├── generate-proto.sh                     # Generate protobuf code from contracts
│   ├── test-all.sh                           # Run all tests across all services
│   ├── lint-all.sh                           # Lint all services
│   └── k8s-apply-all.sh                      # Deploy all services to Kubernetes
│
├── .github/
│   └── workflows/
│       ├── contracts-ci.yml                  # 🆕 Contracts CI/CD (buf lint, breaking check, generate)
│       ├── platform-shared-ci.yml            # 🆕 Platform-shared CI/CD (test, lint, publish)
│       ├── auth-ci.yml                       # 🆕 Auth package CI/CD (test, lint, publish)
│       └── services-ci.yml                   # All services CI/CD (test, build, deploy)
│
├── Makefile                                  # 🆕 Root-level orchestration (make test, make deploy-all)
├── .gitignore                                # Git ignore patterns
├── .editorconfig                             # Editor configuration
├── go.work                                   # 🆕 Go workspace (optional, for multi-module development)
└── README.md                                 # Root README with quick start
```

---

## 📦 **6️⃣ communications-be (Real-time Messaging and Notifications Service)**

```
apps/be/communications-be/
│
├── cmd/
│   ├── api/
│   │   └── main.go                               # 📝 API entrypoint - Gin+Dapr+Postgres (uses internal/config & platform-shared/logging)
│   └── worker/                                   # 🆕 background workers
│       └── main.go                               # 🆕 boot DI; run digest scheduler, DLQ replayer, bounce/complaint consumers
│
├── internal/
│   ├── config/                                   # 🔧 Configuration (Load First)
│   │   ├── schema.go                             # Typed Config (App, Server, Postgres, Kafka, Redis, Auth, WildDuck, WebSocket, WebPush)
│   │   ├── loader.go                             # Viper loader (flags → env → file → defaults)
│   │   └── docs/CONFIGURATION.md                 # ENV vars, defaults, examples
│   │
│   ├── domain/                                   # 🏛️ Domain Layer (business rules, aggregates, events)
│   │   # =========================
│   │   # 💬 CORE CHAT PRIMITIVES
│   │   # =========================
│   │   ├── conversation/
│   │   │   ├── entity.go                         # id, kind(direct,group,system), tenant_id, created_by, visibility, data_zone
│   │   │   ├── participant.go                    # user_id, role(owner,member), last_read_msg_id, pinned, muted_until
│   │   │   ├── settings.go                       # ttl_policy_id, legal_hold, slow_mode, allow_replies, allow_files
│   │   │   ├── typing_indicator.go               # (kept) typing TTL markers per conversation/user
│   │   │   ├── errors.go                         # ConversationNotFound, ParticipantNotFound, UnauthorizedAccess
│   │   │   ├── repository.go                     # Create, FindByID, AddParticipant, RemoveParticipant
│   │   │   └── events.go                         # conversation.created/updated/archived/member_added/removed.v1
│   │   ├── thread/                               # 🆕 sub-discussions
│   │   │   ├── entity.go                         # id, conversation_id, root_message_id, title, followers[]
│   │   │   └── events.go                         # thread.created/renamed/archived.v1
│   │   ├── message/
│   │   │   ├── entity.go                         # id, conversation_id, sender_id, body(rich), reply_to_id, edited_at, deleted_at, redact_reason
│   │   │   ├── attachment.go                     # storage-be asset refs (url,id,hash,type,size,thumb), virus_status
│   │   │   ├── read_receipt.go                   # message_id, user_id, read_at (rollup-friendly)
│   │   │   ├── reaction.go                       # emoji, user_id, reacted_at
│   │   │   ├── mention.go                        # mentioned_user_id, offsets
│   │   │   ├── errors.go                         # MessageNotFound, InvalidContent, MessageTooLong
│   │   │   ├── repository.go                     # Create, Update, Delete, FindByConversation, MarkAsRead
│   │   │   └── events.go                         # message.sent/edited/deleted/reacted/mentioned.v1
│   │   ├── draft/                                # 🆕 per-user unsent drafts
│   │   │   └── entity.go                         # conversation_id, user_id, content, updated_at
│   │   ├── pin/                                  # 🆕 pinned highlights
│   │   │   └── entity.go                         # conversation_id, message_id, pinned_by, pinned_at
│   │   └── bookmark/                             # 🆕 private bookmarks
│   │       └── entity.go                         # user_id, message_id, note, created_at
│   │   # =========================
│   │   # 🚚 DELIVERY & READ STATE
│   │   # =========================
│   │   ├── delivery/                             # server→device delivery state
│   │   │   └── status.go                         # queued→dispatched→ack; per-device/session acks
│   │   └── read_receipt/                         # explicit “I read it”
│   │       └── entity.go                         # message_id, user_id, read_at (compacted)
│   │   # =========================
│   │   # ⚡ EPHEMERAL REALTIME SIGNALS
│   │   # =========================
│   │   ├── presence/                             # online/away/offline
│   │   │   ├── session.go                        # session_id, user_id, last_seen_at, ip, ua, device_kind
│   │   │   └── events.go                         # presence.joined/left/heartbeat.v1
│   │   └── typing/                               # typing indicators (short TTL)
│   │       └── signal.go                         # conversation_id, user_id, started_at, stopped_at
│   │   # =========================
│   │   # 🛡️ SAFETY & COMPLIANCE
│   │   # =========================
│   │   ├── moderation/
│   │   │   ├── automod_rule.go                   # regex/keyword heuristics; actions(quarantine, mask, notify)
│   │   │   ├── flag.go                           # reporter_id, reason, status(pending,reviewed,resolved)
│   │   │   └── quarantine.go                     # hide pending review; emits admin case
│   │   ├── retention/
│   │   │   └── policy.go                         # per-room TTL, dispute_hold; purge windows
│   │   ├── blocklist/                            # 🆕 block phrases/users/domains
│   │   │   └── entity.go                         # scope(user/tenant/global), subject, reason, expires_at
│   │   └── url_safety/                           # 🆕 reputation cache
│   │       └── cache.go                          # open-source feeds only; refresh schedule
│   │   # =========================
│   │   # 🔔 NOTIFS & USER FEEDS
│   │   # =========================
│   │   ├── notification/
│   │   │   ├── entity.go                         # id, user_id, type, title, body, data, priority, read_at, created_at
│   │   │   ├── enums.go                          # type/priorities/categories
│   │   │   ├── preferences.go                    # per-type enablement (email,push,inapp)
│   │   │   ├── settings.go                       # instant/daily/weekly/muted
│   │   │   ├── quiet_hours.go                    # quiet hours per user (tz aware)
│   │   │   ├── repository.go                     # Create, MarkAsRead, Delete, FindByUser, UnreadCount
│   │   │   └── events.go                         # notification.created/updated/read/deleted.v1
│   │   ├── in_app_notification/
│   │   │   ├── entity.go                         # id, notification_id, user_id, displayed_at, dismissed_at, clicked_at
│   │   │   ├── badge_count.go                    # per-user badge counters
│   │   │   ├── group.go                          # grouping rules
│   │   │   ├── action.go                         # CTAs (url,label,type)
│   │   │   └── events.go                         # inapp.displayed/dismissed/clicked/badge.updated.v1
│   │   ├── notification_template/
│   │   │   ├── entity.go                         # versioned templates + i18n
│   │   │   ├── variable.go                       # placeholders whitelist
│   │   │   └── events.go                         # template.created/updated/localized/deactivated.v1
│   │   ├── email/
│   │   │   ├── entity.go                         # id, to_id, to_email, subject, body, template_id, status, sent_at, delivered_at
│   │   │   ├── batch.go                          # batch send tracking
│   │   │   └── events.go                         # email.queued/sent/delivered/bounced/failed.v1
│   │   ├── notification_queue/                   # queues & priorities
│   │   │   ├── entity.go                         # notification_id, priority, scheduled_for, status, retries
│   │   │   └── events.go                         # notification.enqueued/dequeued/retry/deadletter.v1
│   │   ├── delivery_log/
│   │   │   ├── entity.go                         # channel, status, delivered_at, failure_reason, retry_count, latency_ms
│   │   │   └── events.go                         # delivery.logged/failed/bounced.v1
│   │   ├── unsubscribe/
│   │   │   └── entity.go                         # user_id, type, unsubscribed_at, reason
│   │   ├── preference/                           # 🆕 explicit module for prefs (extends existing)
│   │   │   ├── entity.go                         # topic/channel prefs + DND, tz, fallbacks
│   │   │   └── policy.go                         # inheritance global→tenant→user
│   │   ├── suppression/                          # 🆕 bounce/complaint suppression
│   │   │   └── rule.go                           # by address/domain; expiry/notes
│   │   ├── webpush/                              # 🆕 VAPID push (no vendor)
│   │   │   ├── subscription.go                   # endpoint, p256dh, auth, scope, device_info, expires_at
│   │   │   └── events.go                         # webpush.subscription.added/expired.v1
│   │   └── digest/                               # batched notifs
│   │       └── window.go                         # daily/weekly windows; locale cutoffs
│   │   # =========================
│   │   # ✉️ EMAIL BRIDGE (SELF-HOSTED)
│   │   # =========================
│   │   ├── email_bridge/                         # 🆕 reply-by-email
│   │   │   ├── inbound.go                        # plus-address → conversation mapping; signature trim
│   │   │   └── outbound.go                       # queue → MTA; threading headers
│   │   └── mail_tracking/                        # 🆕 delivery/complaint signals (self-hosted logs/webhooks)
│   │       └── event.go                          # delivered, deferred, bounced, complained → suppression
│   │   # =========================
│   │   # 📅 SCHEDULING & CALLS
│   │   # =========================
│   │   ├── system_message/
│   │   │   ├── entity.go                         # system feed items (milestones, disputes, approvals)
│   │   │   └── events.go                         # system_message.created/broadcasted.v1
│   │   ├── call/
│   │   │   ├── entity.go                         # call link, starts_at, ends_at, status
│   │   │   └── events.go                         # call.scheduled/started/ended/canceled.v1
│   │   └── calendar_invite/
│   │       ├── entity.go                         # invite (ical, status)
│   │       └── events.go                         # calendar_invite.sent/accepted/declined/bounced.v1
│   │   # =========================
│   │   # 📊 OPS & GOVERNANCE
│   │   # =========================
│   │   ├── quota/
│   │   │   └── token_bucket.go                   # per-user/topic/channel sliding-window limits
│   │   ├── analytics/
│   │   │   └── funnel.go                         # requested→sent→delivered→ack/read; histograms; SLOs
│   │   ├── audit/
│   │   │   └── trail.go                          # who/what/when for prefs, suppressions, impersonations
│   │   ├── idempotency/
│   │   │   └── key.go                            # keys for messages/notifs; replay-safe contracts
│   │   └── tenancy/
│   │       └── scope.go                          # tenant_id, data_zone, RLS helpers & partitioning hints
│   │   # =========================
│   │   # 🆕 ADDED FOR FREELANCING PLATFORM (UPWORK-LIKE)
│   │   # =========================
│   │   ├── interview/
│   │   │   ├── entity.go                         # id, conversation_id, client_id, freelancer_id, scheduled_at, status, notes
│   │   │   ├── participant.go                    # participant details (availability, timezone)
│   │   │   ├── errors.go                         # InterviewConflict, InvalidSchedule, NoAvailability
│   │   │   ├── repository.go                     # CreateInterview, FindByConversation, UpdateStatus
│   │   │   └── events.go                         # interview.scheduled/confirmed/cancelled/completed.v1
│   │   ├── platform_alert/
│   │   │   ├── entity.go                         # id, severity, message, targets (all/freelancers/clients), expires_at
│   │   │   ├── errors.go                         # InvalidTarget, ExpiredAlert
│   │   │   ├── repository.go                     # CreateAlert, FindActiveAlerts, MarkExpired
│   │   │   └── events.go                         # platform_alert.sent/dismissed.v1
│   │   └── spam_detection/
│   │       ├── entity.go                         # spam_score, detected_patterns, quarantine_status
│   │       ├── errors.go                         # SpamDetected, FalsePositive
│   │       ├── repository.go                     # LogSpamAttempt, GetSpamHistory
│   │       └── events.go                          # spam.detected/quarantined/reviewed.v1
│   │
    └── application/                              # 📋 Application Layer (use cases, orchestrators, consumers)
        │   # =========================
        │   # 📡 EVENT CONSUMERS (INBOX)
        │   # =========================
        ├── eventhandler/
        │   ├── user_handler.go                   # user.created → welcome message/email
        │   ├── job_handler.go                    # job.posted → notify matching freelancers
        │   ├── proposal_handler.go               # proposal.submitted → notify client
        │   ├── contract_handler.go               # contract.created → notify both parties
        │   ├── payment_handler.go                # payment.processed → receipt notification
        │   ├── review_handler.go                 # review.submitted / double_blind window nudges
        │   ├── admin_case_handler.go             # admin.case.* → subject notifications as needed
        │   ├── delivery_logger_handler.go        # emit comm.delivery.logged → admin audit
        │   └── partitioning_notes.go             # (comments) partition keys per stream
        │
        │   # =========================
        │   # 🧠 USE CASES (COMMANDS/QUERIES)
        │   # =========================
        │   # =========================
        │   # 💬 CORE CHAT PRIMITIVES
        │   # =========================
        ├── conversation/
        │   ├── service.go                        # Conversation business logic (Create, Archive, Mute, Delete)
        │   ├── commands.go                       # CreateConversation, ArchiveConversation, MuteConversation, DeleteConversation
        │   ├── queries.go                        # GetConversation, ListConversations, SearchConversations
        │   ├── dto.go                            # ConversationDTO, CreateConversationDTO, ConversationListDTO
        │   ├── mapper.go                         # Entity ↔ DTO mapping for conversations
        │   └── validators.go                     # Input validation (membership, visibility, TTL policies)
        ├── message/
        │   ├── service.go                        # Message business logic (Send, Edit, Delete, React, MarkAsRead)
        │   ├── commands.go                       # SendMessage, EditMessage, DeleteMessage, ReactToMessage, MarkAsRead
        │   ├── queries.go                        # GetMessages, SearchMessages, GetUnreadCount
        │   ├── dto.go                            # MessageDTO, SendMessageDTO, MessageListDTO
        │   ├── mapper.go                         # Message entity ↔ DTO mappers
        │   ├── validators.go                     # Content length, attachment limits, mention bounds
        │   └── realtime_service.go               # 🆕 WebSocket handling (broadcast, typing, presence fan-out)
        ├── thread/
        │   ├── service.go                        # 🆕 Create/Archive thread, follow/unfollow
        │   ├── commands.go                       # 🆕 CreateThread, ArchiveThread, FollowThread, UnfollowThread
        │   ├── queries.go                        # 🆕 GetThread, ListThreadsForConversation
        │   ├── dto.go                            # 🆕 ThreadDTO
        │   ├── mapper.go                         # 🆕 Thread entity ↔ DTO mappers
        │   └── validators.go                     # 🆕 Root message existence, membership checks
        ├── pin/
        │   ├── service.go                        # 🆕 PinMessage, UnpinMessage, ListPins
        │   ├── commands.go                       # 🆕 PinMessage, UnpinMessage
        │   ├── queries.go                        # 🆕 GetPinsForConversation
        │   ├── dto.go                            # 🆕 PinDTO
        │   ├── mapper.go                         # 🆕 Pin entity ↔ DTO mappers
        │   └── validators.go                     # 🆕 Role/visibility checks
        ├── bookmark/
        │   ├── service.go                        # 🆕 AddBookmark, RemoveBookmark, ListBookmarks
        │   ├── commands.go                       # 🆕 AddBookmark, RemoveBookmark
        │   ├── queries.go                        # 🆕 GetBookmarksForUser
        │   ├── dto.go                            # 🆕 BookmarkDTO
        │   ├── mapper.go                         # 🆕 Bookmark entity ↔ DTO mappers
        │   └── validators.go                     # 🆕 Ownership checks
        ├── draft/
        │   ├── service.go                        # 🆕 SaveDraft, ClearDraft, GetDraft
        │   ├── commands.go                       # 🆕 SaveDraft, ClearDraft
        │   ├── queries.go                        # 🆕 GetDraftForConversation
        │   ├── dto.go                            # 🆕 DraftDTO
        │   ├── mapper.go                         # 🆕 Draft entity ↔ DTO mappers
        │   └── validators.go                     # 🆕 Size limits, content sanitization
        │
        │   # =========================
        │   # 🚚 DELIVERY & READ STATE
        │   # =========================
        ├── read_receipt/
        │   ├── service.go                        # 🆕 RecordRead, GetLatestRead, ListReaders
        │   ├── commands.go                       # 🆕 RecordRead
        │   ├── queries.go                        # 🆕 GetReadReceiptsForMessage
        │   ├── dto.go                            # 🆕 ReadReceiptDTO
        │   ├── mapper.go                         # 🆕 Read receipt entity ↔ DTO mappers
        │   └── validators.go                     # 🆕 Membership & ordering checks
        ├── delivery/
        │   ├── service.go                        # 🆕 MarkDispatched, AckDelivery, GetDeliveryStatus
        │   ├── commands.go                       # 🆕 MarkDispatched, AckDelivery
        │   ├── queries.go                        # 🆕 GetDeliveriesForMessage
        │   ├── dto.go                            # 🆕 DeliveryStatusDTO
        │   ├── mapper.go                         # 🆕 Delivery entity ↔ DTO mappers
        │   └── validators.go                     # 🆕 Session authenticity, idempotency keys
        │
        │   # =========================
        │   # ⚡ EPHEMERAL REALTIME SIGNALS
        │   # =========================
        ├── online_status/
        │   ├── service.go                        # Presence state machine (online, away, busy, offline)
        │   ├── commands.go                       # SetOnline, SetAway, SetBusy, SetOffline
        │   ├── queries.go                        # GetUserStatus, GetOnlineUsers
        │   ├── validators.go                     # TTL bounds, allowed transitions
        │   ├── tracker.go                        # Heartbeat ingestion & expiry logic
        │   ├── presence_manager.go               # Session fan-out & dedupe across devices
        │   └── dto.go                            # OnlineStatusDTO
        ├── typing/
        │   ├── service.go                        # 🆕 StartTyping, StopTyping, GetTypingUsers
        │   ├── commands.go                       # 🆕 StartTyping, StopTyping
        │   ├── queries.go                        # 🆕 GetTypingForConversation
        │   ├── dto.go                            # 🆕 TypingDTO
        │   ├── mapper.go                         # 🆕 Typing signal ↔ DTO mappers
        │   └── validators.go                     # 🆕 Rate limits, membership checks
        │
        │   # =========================
        │   # 🛡️ SAFETY & COMPLIANCE
        │   # =========================
        ├── flag/
        │   ├── service.go                        # Message flag lifecycle (Flag, Unflag, Resolve)
        │   ├── commands.go                       # FlagMessage, UnflagMessage, ResolveFlag
        │   ├── queries.go                        # GetFlags, GetFlag
        │   ├── validators.go                     # Reason whitelist, reviewer role checks
        │   ├── dto.go                            # FlagDTO, FlagMessageDTO
        │   └── mapper.go                         # Flag entity ↔ DTO mappers
        ├── moderation/
        │   ├── service.go                        # 🆕 EvaluateRules, ApplyActions (quarantine/mask/notify)
        │   ├── commands.go                       # 🆕 UpsertRule, RemoveRule
        │   ├── queries.go                        # 🆕 ListRules, GetRule
        │   ├── dto.go                            # 🆕 ModerationRuleDTO
        │   ├── mapper.go                         # 🆕 Rule entity ↔ DTO mappers
        │   └── validators.go                     # 🆕 Pattern safety, action constraints
        ├── retention_policy/
        │   ├── service.go                        # 🆕 SetPolicy, GetPolicy, EnforcePolicy
        │   ├── commands.go                       # 🆕 SetRetentionPolicy
        │   ├── queries.go                        # 🆕 GetRetentionPolicy
        │   ├── dto.go                            # 🆕 RetentionPolicyDTO
        │   ├── mapper.go                         # 🆕 Policy entity ↔ DTO mappers
        │   └── validators.go                     # 🆕 TTL bounds, hold precedence
        ├── blocklist/
        │   ├── service.go                        # 🆕 AddBlock, RemoveBlock, IsBlocked
        │   ├── commands.go                       # 🆕 AddBlock, RemoveBlock
        │   ├── queries.go                        # 🆕 GetBlocksForScope
        │   ├── dto.go                            # 🆕 BlockDTO
        │   ├── mapper.go                         # 🆕 Block entity ↔ DTO mappers
        │   └── validators.go                     # 🆕 Scope & expiry checks
        │
        │   # =========================
        │   # 🔔 NOTIFS & USER FEEDS
        │   # =========================
        ├── notification/
        │   ├── service.go                        # Notification business logic (Send, MarkAsRead, Delete, ClearAll)
        │   ├── commands.go                       # SendNotification, MarkAsRead, DeleteNotification, ClearAllNotifications
        │   ├── queries.go                        # GetNotifications, GetUnreadCount, GetNotificationHistory
        │   ├── dto.go                            # NotificationDTO, SendNotificationDTO, NotificationListDTO
        │   ├── mapper.go                         # Notification mappers
        │   ├── validators.go                     # Type/category checks, payload schema validation
        │   ├── orchestrator.go                   # 🆕 Multi-channel orchestration (in-app, email, webpush)
        │   ├── preferences_service.go            # 🆕 Manage per-topic/channel user preferences + DND windows
        │   └── aggregator.go                     # 🆕 Aggregate & group similar notifications (collapse keys/time window)
        ├── notification_preferences/
        │   ├── service.go                        # UpdatePreferences, GetPreferences, SetQuietHours, SetDigestSchedule
        │   ├── commands.go                       # UpdatePreferences, SetQuietHours, SetDigestSchedule
        │   ├── queries.go                        # GetPreferences, GetEffectivePreferences
        │   ├── validators.go                     # Channel validation, timezone-safe windows
        │   ├── dto.go                            # NotificationPreferencesDTO, QuietHoursDTO, DigestScheduleDTO
        │   └── mapper.go                         # Preferences entity ↔ DTO mappers
        ├── in_app_notification/
        │   ├── service.go                        # In-app notification logic (render, deliver, state transitions)
        │   ├── commands.go                       # PushInAppNotification, DismissInAppNotification, ClickInAppAction
        │   ├── queries.go                        # GetInAppNotifications, GetBadgeCount
        │   ├── validators.go                     # CTA validation, throttling & dedupe checks
        │   ├── real_time_sender.go               # Real-time delivery via WS/SSE (user/room fan-out)
        │   ├── badge_manager.go                  # Badge count calc/update/reset with cache hinting
        │   ├── grouping_engine.go                # Grouping similar items (collapse keys)
        │   ├── dto.go                            # InAppNotificationDTO
        │   └── mapper.go                         # In-app notification mappers
        ├── template/
        │   ├── service.go                        # Create/Update/Render templates (versioned + i18n)
        │   ├── commands.go                       # CreateTemplate, UpdateTemplate
        │   ├── queries.go                        # GetTemplate, ListTemplates
        │   ├── validators.go                     # Placeholder whitelist, locale fallback checks
        │   ├── renderer.go                       # Safe renderer (HTML sanitizer, checksum)
        │   ├── variable_injector.go              # Merge dynamic variables from context
        │   ├── dto.go                            # TemplateDTO, RenderTemplateDTO
        │   └── mapper.go                         # Template entity ↔ DTO mappers
        ├── email/
        │   ├── service.go                        # Email orchestration (Send, SendBatch, CheckStatus)
        │   ├── commands.go                       # SendEmail, SendEmailBatch
        │   ├── queries.go                        # GetEmailStatus, ListEmailsForUser
        │   ├── validators.go                     # Address format, batch sizes, template variables
        │   ├── sender.go                         # SMTP send workflow (retries, backoff, idempotency keys)
        │   ├── template_renderer.go              # Render HTML/text templates with variable injection
        │   ├── batch_sender.go                   # Batch queueing & rate limiting
        │   ├── dto.go                            # EmailDTO, SendEmailDTO, BatchEmailDTO
        │   ├── mapper.go                         # Email entity ↔ DTO mappers
        │   └── wildduck_client.go                # 🆕 WildDuck SMTP/API integration (self-hosted MTA)
        │
        │   # =========================
        │   # ✉️ EMAIL BRIDGE (SELF-HOSTED)
        │   # =========================
        ├── email_bridge/
        │   ├── inbound_processor.go              # 🆕 Parse inbound (plus-address) → conversation; trim signatures; auth checks
        │   └── outbound_processor.go             # 🆕 Generate threading headers; queue to MTA; dedupe by message-id
        │
        │   # =========================
        │   # 📅 SCHEDULING & CALLS
        │   # =========================
        ├── system_message/
        │   ├── service.go                        # Publish system messages (milestones, disputes, approvals)
        │   ├── commands.go                       # CreateSystemMessage
        │   ├── queries.go                        # GetSystemMessagesForConversation
        │   ├── validators.go                     # Type safety & payload schema checks
        │   ├── dto.go                            # SystemMessageDTO
        │   └── mapper.go                         # System message entity ↔ DTO mappers
        ├── call/
        │   ├── service.go                        # Create/Cancel/Reschedule call links
        │   ├── commands.go                       # CreateCall, CancelCall, RescheduleCall
        │   ├── queries.go                        # GetCall, ListCallsForConversation
        │   ├── validators.go                     # Timezone-safe windows, overlap rules
        │   ├── dto.go                            # CallDTO
        │   └── mapper.go                         # Call entity ↔ DTO mappers
        ├── calendar_invite/
        │   ├── service.go                        # Generate iCal, SendInvites, UpdateInviteStatus
        │   ├── commands.go                       # SendCalendarInvite
        │   ├── queries.go                        # GetInvitesForCall
        │   ├── validators.go                     # Email, timezone, iCal bounds
        │   ├── dto.go                            # CalendarInviteDTO
        │   └── mapper.go                         # Calendar invite entity ↔ DTO mappers
        │
        │   # =========================
        │   # 📊 OPS & GOVERNANCE
        │   # =========================
        ├── quota/
        │   ├── service.go                        # 🆕 CheckQuota, ConsumeQuota, ResetQuota
        │   ├── commands.go                       # 🆕 ConsumeQuota, ResetQuota
        │   ├── queries.go                        # 🆕 GetQuotaUsage
        │   ├── dto.go                            # 🆕 QuotaUsageDTO
        │   ├── mapper.go                         # 🆕 Quota entity ↔ DTO mappers
        │   └── validators.go                     # 🆕 Window sizes, per-topic caps
        ├── analytics/
        │   ├── service.go                        # 🆕 RecordFunnelStep, ReportLagDistributions
        │   ├── commands.go                       # 🆕 RecordRequested, RecordSent, RecordDelivered, RecordRead
        │   ├── queries.go                        # 🆕 GetFunnel, GetLagHistogram
        │   ├── dto.go                            # 🆕 FunnelDTO, LagHistogramDTO
        │   ├── mapper.go                         # 🆕 Analytics model ↔ DTO mappers
        │   └── validators.go                     # 🆕 Sampling & PII guards
        │
        │   # =========================
        │   # 🆕 ADDED FOR FREELANCING PLATFORM (UPWORK-LIKE)
        │   # =========================
        ├── interview/
        │   ├── service.go                        # 🆕 ScheduleInterview, ConfirmInterview, CancelInterview, CompleteInterview
        │   ├── commands.go                       # 🆕 ScheduleInterview, ConfirmInterview, CancelInterview
        │   ├── queries.go                        # 🆕 GetInterview, ListInterviewsForConversation
        │   ├── dto.go                            # 🆕 InterviewDTO, ScheduleInterviewDTO
        │   ├── mapper.go                         # 🆕 Interview entity ↔ DTO mappers
        │   └── validators.go                     # 🆕 Availability checks, timezone validation, participant consent
        ├── platform_alert/
        │   ├── service.go                        # 🆕 SendPlatformAlert, DismissAlert, GetActiveAlerts
        │   ├── commands.go                       # 🆕 SendPlatformAlert, DismissAlert
        │   ├── queries.go                        # 🆕 GetActiveAlertsForUser
        │   ├── dto.go                            # 🆕 PlatformAlertDTO
        │   ├── mapper.go                         # 🆕 Alert entity ↔ DTO mappers
        │   └── validators.go                     # 🆕 Target segmentation, expiry bounds
        └── spam_detection/
            ├── service.go                        # 🆕 DetectSpam, QuarantineMessage, ReviewSpam
            ├── commands.go                       # 🆕 QuarantineMessage, ReviewSpam
            ├── queries.go                        # 🆕 GetSpamHistoryForUser
            ├── dto.go                            # 🆕 SpamDetectionDTO
            ├── mapper.go                         # 🆕 Spam entity ↔ DTO mappers
            ├── rules_engine.go                   # 🆕 Rule-based detection (keywords, links, repetition; no paid ML)
            └── validators.go                     # 🆕 Score thresholds, false-positive overrides

│   │
│   ├── infrastructure/                          # 🔌 Infrastructure (DB, cache, messaging, realtime, email, push)
│   │   # =========================
│   │   # 🗄️ PERSISTENCE (POSTGRES)
│   │   # =========================
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       # 🧱 COMMON (DB BOOTSTRAP)
│   │   │       ├── connection.go                    # DSN & pooling
│   │   │       ├── transaction.go                   # TX helpers
│   │   │       ├── migrations.go                    # auto-migrate (versioned)
│   │   │       ├── version.go                       # schema version table
│   │   │       ├── safety.go                        # pre-flight checks
│   │   │       # 💬 CORE CHAT PRIMITIVES
│   │   │       ├── conversation_repository.go       # ConversationRepository implementation (CRUD, members, settings)
│   │   │       ├── message_repository.go            # MessageRepository (send/edit/delete, reads/reactions)
│   │   │       ├── thread_repository.go             # 🆕 ThreadRepository (create/archive, follow/unfollow)
│   │   │       ├── pin_repository.go                # 🆕 PinRepository (pin/unpin, list pins)
│   │   │       ├── bookmark_repository.go           # 🆕 BookmarkRepository (user private bookmarks)
│   │   │       ├── draft_repository.go              # 🆕 DraftRepository (per-user unsent drafts)
│   │   │       # 🚚 DELIVERY & READ STATE
│   │   │       ├── read_receipt_repository.go       # 🆕 ReadReceiptRepository (record/read rollups)
│   │   │       ├── delivery_repository.go           # 🆕 DeliveryRepository (queued→dispatched→ack states)
│   │   │       # ⚡ EPHEMERAL REALTIME SIGNALS
│   │   │       ├── online_status_repository.go      # OnlineStatusRepository (sessions/heartbeats)
│   │   │       # 🛡️ SAFETY & COMPLIANCE
│   │   │       ├── moderation_flag_repository.go    # 🆕 ModerationFlagRepository (reports, statuses, evidence)
│   │   │       ├── retention_policy_repository.go   # 🆕 RetentionPolicyRepository (TTL/legal holds per conversation)
│   │   │       ├── blocklist_repository.go          # 🆕 BlocklistRepository (user/phrase/domain blocks)
│   │   │       ├── audit_trail_repository.go        # 🆕 AuditTrailRepository (immutable ops/security trail)
│   │   │       # 🔔 NOTIFS & USER FEEDS
│   │   │       ├── notification_repository.go       # NotificationRepository (feed + counters)
│   │   │       ├── in_app_notification_repository.go# InAppNotificationRepository (badge/grouping)
│   │   │       ├── notification_queue_repository.go # NotificationQueueRepository (priority/ETA/retries)
│   │   │       ├── delivery_log_repository.go       # DeliveryLogRepository (status/latency)
│   │   │       ├── template_repository.go           # TemplateRepository (versioned + i18n templates)
│   │   │       ├── unsubscribe_repository.go        # UnsubscribeRepository (per-type/channel)
│   │   │       ├── suppression_repository.go        # 🆕 SuppressionRepository (bounces/complaints)
│   │   │       ├── webpush_subscription_repository.go # 🆕 WebPushSubscriptionRepository (VAPID endpoint/keys/scope/expiry)
│   │   │       # ✉️ EMAIL BRIDGE (SELF-HOSTED)
│   │   │       ├── email_repository.go              # EmailRepository (status/outbox/history)
│   │   │       ├── mail_tracking_repository.go      # 🆕 MailTrackingRepository (delivered/deferred/bounced/complained)
│   │   │       # 📅 SCHEDULING & CALLS
│   │   │       ├── system_message_repository.go     # SystemMessageRepository (contract/dispute feeds)
│   │   │       ├── call_repository.go               # CallRepository (links/schedule/state)
│   │   │       ├── calendar_invite_repository.go    # CalendarInviteRepository (iCal/status)
│   │   │       # 📊 OPS & GOVERNANCE
│   │   │       ├── quota_repository.go              # 🆕 QuotaRepository (per-user/topic/channel usage)
│   │   │       └── analytics_repository.go          # 🆕 AnalyticsRepository (funnel counters, lag aggregates)
│   │   # =========================
│   │   # ⚡ CACHE (REDIS)
│   │   # =========================
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go                   # Redis connection setup (pooling, retry logic)
│   │   │       ├── conversation_cache.go           # Conversation caching (Get, Set, Invalidate with TTL)
│   │   │       ├── online_status_cache.go          # fast presence reads
│   │   │       ├── typing_indicator_cache.go       # short TTL ephemeral typing
│   │   │       ├── notification_cache.go           # unread/badge counts
│   │   │       └── presence_cache.go               # session sets per user/device
│   │   # =========================
│   │   # 📨 MESSAGING (KAFKA)
│   │   # =========================
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go                     # 📝 UPDATED: uses platform-shared/inbox (dedupe, offsets)
│   │   │       ├── producer.go                     # 📝 UPDATED: uses platform-shared/outbox (reliable publishing)
│   │   │       ├── topics.go                       # 📝 UPDATED: topic constants imported from contracts/events
│   │   │       └── scram.go                        # SASL/SCRAM-256
│   │   # =========================
│   │   # 🔴 REALTIME (WS/SSE)
│   │   # =========================
│   │   ├── realtime/
│   │   │   ├── websocket/
│   │   │   │   ├── hub.go                          # connection manager; rooms; user fan-out
│   │   │   │   ├── client.go                       # per-conn read/write goroutines
│   │   │   │   ├── handler.go                      # HTTP→WS upgrade & auth
│   │   │   │   ├── broadcaster.go                  # send to all/user/room
│   │   │   │   └── room.go                         # room registry (conversation_id)
│   │   │   └── sse/
│   │   │       ├── handler.go                      # Server-Sent Events handler
│   │   │       └── stream.go                       # stream registry & writes
│   │   # =========================
│   │   # ✉️ EMAIL (SELF-HOSTED)
│   │   # =========================
│   │   ├── email/
│   │   │   ├── wildduck/
│   │   │   │   ├── client.go                       # WildDuck SMTP/API client
│   │   │   │   ├── smtp_sender.go                  # SMTP send
│   │   │   │   ├── api_client.go                   # mailbox mgmt / filters
│   │   │   │   └── config.go                       # SMTP creds/hosts
│   │   │   └── smtp/
│   │   │       ├── client.go                       # generic SMTP fallback
│   │   │       └── config.go
│   │   # =========================
│   │   # 🔔 WEBPUSH (VAPID)
│   │   # =========================
│   │   └── webpush/
│   │       └── vapid/
│   │           ├── signer.go                       # 🆕 VAPID JWT/EC keys
│   │           └── sender.go                       # 🆕 push send with retries
│   │   # =========================
│   │   # 📦 INTEGRATIONS
│   │   # =========================
│   │   └── storage/
│   │       └── client.go                           # Storage service client (upload message attachments via HTTP API)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       └── v1/                                 # 🌐 HTTP Interface
│   │           # =========================
│   │           # 🧭 HANDLERS
│   │           # =========================
│   │           ├── handlers/
│   │           │   # 💬 CORE CHAT PRIMITIVES
│   │           │   ├── conversation_handler.go         # CRUD & member ops
│   │           │   ├── message_handler.go              # send/edit/delete/react/mark-read
│   │           │   ├── thread_handler.go               # 🆕 create/archive threads; follow/unfollow
│   │           │   ├── pin_handler.go                  # 🆕 pin/unpin messages; list pins
│   │           │   ├── bookmark_handler.go             # 🆕 add/remove/list user bookmarks
│   │           │   ├── draft_handler.go                # 🆕 save/clear/get conversation draft
│   │           │   # 🚚 DELIVERY & READ STATE
│   │           │   ├── read_receipt_handler.go         # 🆕 record read; list readers; latest read
│   │           │   ├── delivery_handler.go             # 🆕 mark dispatched/ack; delivery status
│   │           │   # ⚡ EPHEMERAL REALTIME SIGNALS
│   │           │   ├── online_status_handler.go        # GET/SET presence
│   │           │   ├── typing_handler.go               # 🆕 start/stop typing; list typers
│   │           │   # 🛡️ SAFETY & COMPLIANCE
│   │           │   ├── flag_handler.go                 # message flags
│   │           │   ├── moderation_handler.go           # 🆕 rules CRUD; evaluate; dry-run
│   │           │   ├── retention_policy_handler.go     # 🆕 set/get retention & legal holds
│   │           │   ├── blocklist_handler.go            # 🆕 add/remove/list user/phrase/domain blocks
│   │           │   # 🔔 NOTIFS & USER FEEDS
│   │           │   ├── notification_handler.go         # list, unread, mark-read, delete
│   │           │   ├── in_app_notification_handler.go  # list in-app; badge count
│   │           │   ├── preferences_handler.go          # GET/PUT user prefs & DND
│   │           │   ├── template_handler.go             # CRUD templates; render preview
│   │           │   ├── webpush_handler.go              # 🆕 subscribe/unsubscribe push
│   │           │   ├── unsubscribe_handler.go          # unsubscribe flows
│   │           │   ├── email_handler.go                # send email / batch status
│   │           │   # ✉️ EMAIL BRIDGE (SELF-HOSTED)
│   │           │   ├── mail_tracking_handler.go        # 🆕 provider webhooks: delivered/deferred/bounced/complained
│   │           │   # 📅 SCHEDULING & CALLS
│   │           │   ├── system_message_handler.go       # system feed
│   │           │   ├── call_handler.go                 # call links & scheduling
│   │           │   ├── calendar_invite_handler.go      # invites
│   │           │   # 📊 OPS & GOVERNANCE
│   │           │   └── health_handler.go               # /health, /ready, /live
│   │           # =========================
│   │           # 🗺️ ROUTES
│   │           # =========================
│   │           └── routes/
│   │               # 💬 CORE CHAT PRIMITIVES
│   │               ├── conversation_routes.go          # /conversations/*
│   │               ├── message_routes.go               # /conversations/:id/messages
│   │               ├── thread_routes.go                # 🆕 /conversations/:id/threads/*
│   │               ├── pin_routes.go                   # 🆕 /conversations/:id/pins/*
│   │               ├── bookmark_routes.go              # 🆕 /bookmarks/*
│   │               ├── draft_routes.go                 # 🆕 /conversations/:id/draft
│   │               # 🚚 DELIVERY & READ STATE
│   │               ├── read_receipt_routes.go          # 🆕 /messages/:id/read-receipts/*
│   │               ├── delivery_routes.go              # 🆕 /messages/:id/deliveries/*
│   │               # ⚡ EPHEMERAL REALTIME SIGNALS
│   │               ├── online_status_routes.go         # /status/*
│   │               ├── typing_routes.go                # 🆕 /conversations/:id/typing/*
│   │               ├── websocket_routes.go             # /ws
│   │               ├── sse_routes.go                   # /sse/*
│   │               # 🛡️ SAFETY & COMPLIANCE
│   │               ├── flag_routes.go                  # /messages/:id/flags/*
│   │               ├── moderation_routes.go            # 🆕 /moderation/rules/*
│   │               ├── retention_policy_routes.go      # 🆕 /conversations/:id/retention
│   │               ├── blocklist_routes.go             # 🆕 /blocklist/*
│   │               # 🔔 NOTIFS & USER FEEDS
│   │               ├── notification_routes.go          # /notifications/*
│   │               ├── in_app_notification_routes.go   # /notifications/in-app/*
│   │               ├── preferences_routes.go           # /preferences/*
│   │               ├── template_routes.go              # /templates/*
│   │               ├── webpush_routes.go               # 🆕 /webpush/*
│   │               ├── unsubscribe_routes.go           # /unsubscribe/*
│   │               ├── email_routes.go                 # /emails/*
│   │               # ✉️ EMAIL BRIDGE (SELF-HOSTED)
│   │               ├── mail_tracking_routes.go         # 🆕 /mail/tracking/*
│   │               # 📅 SCHEDULING & CALLS
│   │               ├── system_message_routes.go        # /system-messages/*
│   │               ├── call_routes.go                  # /calls/*
│   │               └── calendar_invite_routes.go       # /calendar-invites/*
│
├── templates/
│   ├── email/                                      # Email HTML templates
│   │   ├── base.html                               # Base email template (header, footer, styles)
│   │   ├── welcome.html                            # Welcome email
│   │   ├── job_alert.html                          # Job alert email
│   │   ├── new_proposal.html                       # New proposal notification
│   │   ├── proposal_accepted.html                  # Proposal accepted notification
│   │   ├── bid_received.html                       # New bid notification
│   │   ├── outbid_alert.html                       # Outbid alert
│   │   ├── contract_created.html                   # Contract created notification
│   │   ├── milestone_completed.html                # Milestone completed notification
│   │   ├── payment_received.html                   # Payment received notification
│   │   ├── payment_sent.html                       # Payment sent notification
│   │   ├── review_request.html                     # Review request
│   │   ├── review_received.html                    # Review received notification
│   │   ├── new_message.html                        # New message notification
│   │   ├── password_reset.html                     # Password reset email
│   │   ├── verify_email.html                       # Email verification
│   │   └── weekly_summary.html                     # Weekly activity summary
│   └── notification/                               # In-app notification templates (JSON)
│       ├── job_posted.json
│       ├── proposal_received.json
│       ├── contract_created.json
│       ├── milestone_completed.json
│       ├── payment_received.json
│       └── review_received.json
│
├── config/
│   ├── default.yaml                                # Default configuration
│   ├── dev.yaml                                    # Development overrides
│   └── prod.yaml                                   # Production overrides
│
├── dapr/                                           # Dapr components split by environment
│   ├── local/                                      # For dapr run
│   │   ├── pubsub.yaml                             # Kafka pub/sub component
│   │   └── statestore.yaml                         # State store component
│   └── k8s/                                        # For kubectl apply
│       ├── pubsub.yaml                             # Kafka with scopes: ["communications-be"]
│       ├── statestore.yaml                         # State store with scopes
│       └── secrets.yaml                            # Dapr secret store
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                         # Kubernetes Deployment
│       ├── service.yaml                            # Kubernetes Service
│       ├── configmap.yaml                          # ConfigMap
│       ├── secrets.yaml                            # Secrets
│       ├── hpa.yaml                                # HPA
│       ├── pdb.yaml                                # PDB
│       └── servicemonitor.yaml                     # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                              # Setup local environment
│   ├── get-secrets.sh                              # Fetch secrets
│   ├── seed-templates.sh                           # Seed message/email templates
│   └── seed-data.sh                                # Seed test data
│
├── tests/
│   ├── unit/
│   │   ├── domain/
│   │   │   ├── conversation_test.go               # Conversation domain tests
│   │   │   ├── message_test.go                    # Message domain tests
│   │   │   └── notification_test.go               # Notification domain tests
│   │   ├── application/
│   │   │   ├── message_service_test.go            # Message service tests
│   │   │   ├── notification_service_test.go       # Notification service tests
│   │   │   └── moderation_service_test.go         # Moderation service tests
│   │   └── infrastructure/
│   │       ├── postgres_repository_test.go        # Postgres repository tests
│   │       └── kafka_producer_test.go             # Kafka producer tests
│   ├── integration/
│   │   ├── handlers/
│   │   │   ├── conversation_handler_test.go       # Conversation HTTP tests
│   │   │   ├── message_handler_test.go            # Message HTTP tests
│   │   │   └── notification_handler_test.go       # Notification HTTP tests
│   │   └── repositories/
│   │       ├── message_repository_test.go         # Message repo tests
│   │       └── notification_repository_test.go    # Notification repo tests
│   └── e2e/
│       └── scenarios/
│           ├── messaging_flow_test.go             # Send/edit/delete/read flow
│           ├── notification_flow_test.go          # Orchestrated multi-channel flow
│           └── moderation_flow_test.go            # Report → quarantine → resolution
│
├── docs/
│   ├── README.md                                  # Service overview
│   ├── API.md                                     # API documentation
│   ├── EVENTS.md                                  # 📝 published: message.sent, notification.delivered; consumed: users/jobs/proposals/contracts/payments/reviews/admin
│   ├── ARCHITECTURE.md                            # High-level diagrams
│   ├── MIGRATIONS.md                              # Migration history
│   ├── SCHEMA.md                                  # Database schema
│   ├── RUNBOOK.md                                 # Operational procedures (DLQ, rate-limit, replayer)
│   ├── websocket-protocol.md                      # WebSocket protocol documentation
│   ├── notification-system.md                     # Notification system overview
│   ├── in-app-notifications.md                    # In-app notifications guide
│   └── wildduck-integration.md                    # WildDuck integration guide
│
├── .github/
│   └── workflows/
│       ├── ci.yml                                 # CI workflow
│       └── cd.yml                                 # CD workflow
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                              # Service-specific errors
│   │   └── codes.go                               # Error codes (MESSAGE_NOT_FOUND, CONVERSATION_NOT_FOUND)
│   ├── utils/
│   │   ├── validator.go                           # Local validation utilities
│   │   ├── template_engine.go                     # Template rendering utilities
│   │   ├── sanitizer.go                           # Sanitize message content (prevent XSS)
│   │   └── html_to_text.go                        # Convert HTML to plain text (for email fallback)
│   └── constants/
│       ├── notification_types.go                  # Notification type constants
│       └── websocket_events.go                    # WebSocket event types
│
├── go.mod                                         # 📝 UPDATED: Imports pkg/auth, platform-shared, contracts/events
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md


```
---

---

## **📦 7️⃣ storage-be (UPDATED)**
```
apps/be/storage-be/
│
├── cmd/
│   └── api/
│       └── main.go                           # 📝 UPDATED: Application entry point - initializes Gin, Dapr, connects to Postgres (now uses loadConfig from internal/config, platform-shared/logging)
│
├── internal/
│   ├── domain/
│   │   ├── file/
│   │   │   ├── entity.go                     # File metadata
│   │   │   ├── enums.go                      # FileType, Status, Visibility
│   │   │   ├── metadata.go                   # File metadata
│   │   │   ├── errors.go                     # 🆕 File domain errors (FileNotFound, InvalidMimeType, FileTooLarge)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: FileCreated, FileUpdated, FileMoved, FileDeleted (hard)
│   │   │
│   │   ├── folder/
│   │   │   ├── entity.go                     # Folder structure
│   │   │   ├── errors.go                     # 🆕 Folder domain errors (FolderNotFound, NameConflict)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: FolderCreated, FolderRenamed, FolderMoved, FolderDeleted
│   │   │
│   │   ├── upload/
│   │   │   ├── entity.go                     # Upload sessions
│   │   │   ├── chunk.go                      # Chunked uploads
│   │   │   ├── resumable.go                  # Resumable upload support
│   │   │   ├── errors.go                     # 🆕 Upload domain errors (UploadSessionNotFound, ChunkOutOfOrder)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: UploadStarted, UploadChunkAppended, UploadResumed, UploadCompleted, UploadAborted
│   │   │
│   │   ├── media/
│   │   │   ├── entity.go                     # Media processing records
│   │   │   ├── thumbnail.go                  # Thumbnail generation
│   │   │   ├── variant.go                    # Image variants
│   │   │   ├── errors.go                     # 🆕 Media domain errors (ProcessingFailed, UnsupportedFormat)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: MediaProcessingStarted, MediaProcessingSucceeded, MediaProcessingFailed, ThumbnailGenerated
│   │   │
│   │   ├── access_control/
│   │   │   ├── entity.go                     # File access permissions
│   │   │   ├── errors.go                     # 🆕 Access control errors (PermissionDenied, ACLNotFound)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: AccessGranted, AccessRevoked, ACLUpdated
│   │   │
│   │   ├── version/
│   │   │   ├── entity.go                     # File versioning
│   │   │   ├── errors.go                     # 🆕 Versioning errors (VersionNotFound, ImmutableVersion)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: FileVersionCreated, FileVersionPromoted, FileVersionRestored
│   │   │
│   │   ├── share/
│   │   │   ├── entity.go                     # File sharing
│   │   │   ├── link.go                       # Share links
│   │   │   ├── errors.go                     # 🆕 Share errors (ShareNotFound, ShareExpired)
│   │   │   ├── repository.go                 # ShareRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: ShareCreated, ShareRevoked, ShareExpired
│   │   │
│   │   ├── file_flag/
│   │   │   ├── entity.go                     # Flagged files
│   │   │   ├── reason.go                     # Flag reasons
│   │   │   ├── status.go                     # Flag status
│   │   │   ├── errors.go                     # 🆕 Flag errors (FlagNotFound, InvalidFlagReason)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: FileFlagSubmitted, FileFlagResolved, FileFlagDismissed
│   │   │
│   │   ├── policy/                           # 🆕 File policies: size/type, virus scan, DLP
│   │   │   ├── entity.go                     # Policy (ID, Name, MaxSizeMB, AllowedMIMEs, BlockedMIMEs, EnforceVirusScan, EnforceDLP)
│   │   │   ├── dlp_pattern.go                # DLP patterns for PII/PCI (regexes, detectors)
│   │   │   ├── result.go                     # Policy evaluation results (violations, reasons)
│   │   │   ├── errors.go                     # Policy errors
│   │   │   ├── repository.go                 # PolicyRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: FilePolicyCreated, FilePolicyUpdated, FilePolicyViolationDetected
│   │   │
│   │   ├── lifecycle/                        # 🆕 Lifecycle rules: retention, legal hold, soft-delete
│   │   │   ├── entity.go                     # LifecycleRule (ContractState, RetentionDays, LegalHold, RestoreWindowDays)
│   │   │   ├── soft_delete.go                # Soft-delete markers (DeletedAt, RestoreBy)
│   │   │   ├── legal_hold.go                 # Legal hold records (PlacedBy, Reason, ExpiresAt)
│   │   │   ├── errors.go                     # Lifecycle errors
│   │   │   ├── repository.go                 # LifecycleRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: FileSoftDeleted, FileRestored, LegalHoldPlaced, LegalHoldRemoved
│   │   │
│   │   ├── linking/                          # 🆕 Versioning & linking: signed URLs, download audit logs
│   │   │   ├── signed_url.go                 # Signed URL (FileID, Action, ExpiresAt, Signature)
│   │   │   ├── audit_log.go                  # Download audit log (FileID, UserID, IP, UserAgent, At)
│   │   │   ├── errors.go                     # Linking errors
│   │   │   ├── repository.go                 # LinkingRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: SignedURLCreated, SignedURLRevoked, FileDownloadLogged
│   │   │
│   │   └── outbox/
│   │       ├── entity.go                     # ❌ REMOVED (use platform-shared/outbox/entity.go)
│   │       └── repository.go                 # ❌ REMOVED (use platform-shared/outbox/repository.go)
│   │
│   ├── application/
│   │   ├── eventhandler/
│   │   │   ├── user_handler.go                # Consumes: user.updated
│   │   │   │                                     # Purpose: refresh ownership/ACL context (e.g., org/role changes) used by access_control/policy.
│   │   │   │
│   │   │   ├── contract_handler.go            # Consumes: contract.state.changed
│   │   │   │                                     # Purpose: enforce lifecycle/retention & legal holds by contract state
│   │   │   │                                     #           (place/remove holds; adjust restore windows/retention markers).
│   │   │   │
│   │   │   ├── admin_policy_handler.go        # Consumes: admin.policy.updated
│   │   │   │                                     # Purpose: update storage policy cache (size/type limits, virus-scan & DLP patterns/rules).
│   │   │   │
│   │   │   └── admin_moderation_handler.go    # Consumes: admin.moderation.actioned (approve/reject/remove/hide)
│   │   │                                         # Purpose: reflect moderation outcomes in storage:
│   │   │                                         #           - quarantine or restore files,
│   │   │                                         #           - revoke or reissue signed URLs,
│   │   │                                         #           - toggle visibility flags for indexed artifacts.
│   │   ├── file/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Upload, Delete, Move, Copy
│   │   │   ├── queries.go                    # Get, List, Search
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── upload/
│   │   │   ├── service.go
│   │   │   ├── chunked_upload.go
│   │   │   ├── resumable.go
│   │   │   ├── dto.go
│   │   │   ├── commands.go                   # 🆕 StartUpload, AppendChunk, CompleteUpload, AbortUpload
│   │   │   ├── queries.go                    # 🆕 GetUploadSession, GetUploadProgress
│   │   │   └── validators.go                 # 🆕 Validate chunk order/size, checksum
│   │   ├── media/
│   │   │   ├── service.go
│   │   │   ├── image_processor.go
│   │   │   ├── video_processor.go
│   │   │   ├── thumbnail_generator.go
│   │   │   ├── dto.go
│   │   │   ├── commands.go                   # 🆕 ProcessImage, ProcessVideo, GenerateThumbnail
│   │   │   ├── queries.go                    # 🆕 GetMediaJob, ListMediaJobs
│   │   │   └── validators.go                 # 🆕 Validate dimensions/codec/bitrate
│   │   ├── folder/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                   # 🆕 CreateFolder, RenameFolder, MoveFolder, DeleteFolder
│   │   │   ├── queries.go                    # 🆕 GetFolder, ListFolderContents, SearchFolders
│   │   │   └── validators.go                 # 🆕 Validate names, cycles, path depth
│   │   ├── share/
│   │   │   ├── service.go
│   │   │   ├── link_generator.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                   # 🆕 CreateShare, RevokeShare, UpdateShare
│   │   │   ├── queries.go                    # 🆕 GetShare, ListSharesByFile
│   │   │   └── validators.go                 # 🆕 Validate expiry/scopes/access
│   │   ├── version/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   │   ├── commands.go                   # 🆕 CreateVersion, RestoreVersion, DeleteVersion
│   │   │   ├── queries.go                    # 🆕 GetVersion, ListVersions
│   │   │   └── validators.go                 # 🆕 Validate immutable rules
│   │   ├── flag/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Flag, Unflag file
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go                 # 🆕 Validate flag reason/state transitions
│   │   ├── policy/                           # 🆕 File policy application layer
│   │   │   ├── service.go                    # EvaluatePolicy, UpdatePolicy, GetPolicy
│   │   │   ├── commands.go                   # SetPolicy, EnableDLP, DisableDLP
│   │   │   ├── queries.go                    # GetPolicy, ListPolicies
│   │   │   ├── validators.go                 # Validate size/type lists, regexes
│   │   │   ├── dto.go                        # PolicyDTO, DLPResultDTO
│   │   │   └── mapper.go                     # Policy mappers
│   │   ├── lifecycle/                        # 🆕 Lifecycle application layer
│   │   │   ├── service.go                    # ApplyRules, SoftDelete, Restore, PlaceLegalHold, RemoveLegalHold
│   │   │   ├── commands.go                   # DefineRule, UpdateRule, DeleteRule, SoftDeleteFile, RestoreFile, PlaceLegalHold
│   │   │   ├── queries.go                    # GetRules, GetFileLifecycle, GetLegalHolds
│   │   │   ├── validators.go                 # Validate retention windows, restore periods
│   │   │   ├── dto.go                        # LifecycleRuleDTO, LegalHoldDTO
│   │   │   └── mapper.go                     # Lifecycle mappers
│   │   ├── linking/                          # 🆕 Signed URLs & audit logs
│   │   │   ├── service.go                    # CreateSignedURL, RevokeSignedURL, LogDownload, ListAuditLogs
│   │   │   ├── commands.go                   # CreateSignedURL, RevokeSignedURL
│   │   │   ├── queries.go                    # GetSignedURL, ListSignedURLs, GetAuditLogs
│   │   │   ├── validators.go                 # Validate expiry, actions, scopes
│   │   │   ├── dto.go                        # SignedURLDTO, AuditLogDTO
│   │   │   └── mapper.go                     # Linking mappers
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection setup (DSN from config, connection pooling)
│   │   │       ├── transaction.go            # Transaction helpers (Begin, Commit, Rollback, WithTransaction wrapper)
│   │   │       ├── migrations.go             # 📝 UPDATED: Auto-migrate logic (now with version tracking, GORM AutoMigrate for all tables)
│   │   │       ├── version.go                # 🆕 Schema version tracking (SchemaVersion table, RecordMigration function)
│   │   │       ├── safety.go                 # 🆕 Pre-migration safety checks (environment validation, disk space check)
│   │   │       ├── file_repository.go        # FileRepository implementation with GORM
│   │   │       ├── folder_repository.go
│   │   │       ├── upload_repository.go
│   │   │       ├── media_repository.go
│   │   │       ├── access_control_repository.go
│   │   │       ├── version_repository.go
│   │   │       ├── share_repository.go
│   │   │       ├── file_flag_repository.go
│   │   │       ├── policy_repository.go      # 🆕 PolicyRepository implementation
│   │   │       ├── lifecycle_repository.go   # 🆕 LifecycleRepository implementation
│   │   │       └── linking_repository.go     # 🆕 LinkingRepository implementation (signed URLs, audit logs)
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go             # Redis connection setup (connection pooling, retry logic)
│   │   │       └── file_cache.go             # File metadata caching (Get, Set, Invalidate with TTL)
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go               # 📝 UPDATED: Kafka consumer (now uses platform-shared/inbox for message deduplication)
│   │   │       ├── producer.go               # 📝 UPDATED: Kafka producer (now uses platform-shared/outbox for reliable publishing)
│   │   │       ├── topics.go                 # 📝 UPDATED: Topic constants (now imported from contracts/events - file.uploaded, file.deleted, media.processed)
│   │   │       └── scram.go                  # SCRAM authentication for Kafka (SASL/SCRAM-SHA-256)
│   │   ├── object_storage/
│   │   │   ├── local/
│   │   │   │   ├── storage.go                # Local file system storage
│   │   │   │   └── config.go                 # Local storage config
│   │   │   ├── minio/
│   │   │   │   ├── client.go                 # Self-hosted MinIO client (upload, download, presigned URLs)
│   │   │   │   └── config.go                 # MinIO configuration (endpoint, credentials, bucket names)
│   │   │   ├── signer.go                     # 🆕 Signed URL signer abstraction (create/revoke)
│   │   │   └── provider.go                   # Storage provider abstraction
│   │   ├── media_processing/
│   │   │   ├── image/
│   │   │   │   ├── resizer.go                # Image resizing logic
│   │   │   │   ├── optimizer.go              # Image optimization (compression)
│   │   │   │   └── watermark.go              # Watermarking images
│   │   │   ├── video/
│   │   │   │   ├── transcoder.go             # Video transcoding
│   │   │   │   └── thumbnail.go              # Video thumbnail generation
│   │   │   └── processor.go                  # General media processor
│   │   ├── virus_scan/
│   │   │   └── clamav.go                     # ClamAV integration for virus scanning
│   │   ├── dlp/                              # 🆕 DLP engine adapters
│   │   │   ├── regex_engine.go               # Regex-based detectors (PII/PCI patterns)
│   │   │   └── provider.go                   # Pluggable DLP providers (custom/third-party)
│   │   └── outbox/
│   │       ├── processor.go                  # ❌ REMOVED (use platform-shared/outbox/forwarder.go)
│   │       └── scheduler.go                  # ❌ REMOVED (use platform-shared/outbox/scheduler.go)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── file_handler.go           # File HTTP handlers (GET, POST, DELETE /files)
│   │       │   ├── upload_handler.go         # Upload handlers (POST /upload, chunked support)
│   │       │   ├── download_handler.go       # Download handlers (GET /download/:id)
│   │       │   ├── folder_handler.go         # Folder handlers
│   │       │   ├── share_handler.go          # Share handlers
│   │       │   ├── media_handler.go          # Media processing handlers
│   │       │   ├── flag_handler.go           # Flag handlers
│   │       │   ├── policy_handler.go         # 🆕 Policy endpoints (GET/PUT policies, evaluate)
│   │       │   ├── lifecycle_handler.go      # 🆕 Lifecycle endpoints (soft-delete/restore/legal hold)
│   │       │   ├── linking_handler.go        # 🆕 Signed URL + audit log endpoints
│   │       │   └── health_handler.go         # Health check endpoints (/health, /ready, /live)
│   │       ├── middleware/
│   │       │   ├── auth.go                   # 📝 UPDATED: Authentication middleware (uses pkg/auth for JWT verification)
│   │       │   ├── rbac.go                   # 📝 UPDATED: RBAC middleware (uses pkg/auth authorizer for role checks)
│   │       │   ├── cors.go                   # 📝 UPDATED: CORS middleware (uses platform-shared/ginx CORS)
│   │       │   ├── rate_limit.go             # Rate limiting middleware (token bucket algorithm)
│   │       │   ├── logging.go                # 📝 UPDATED: Logging middleware (uses platform-shared/ginx/logging)
│   │       │   ├── error_handler.go          # Error handling middleware
│   │       │   ├── request_id.go             # 📝 UPDATED: Request ID middleware (uses platform-shared/ginx/requestid)
│   │       │   └── file_size_limit.go        # File size limit middleware
│   │       ├── responses/
│   │       │   ├── success.go                # 📝 UPDATED: Success response wrappers (use platform-shared/httpx/response.go)
│   │       │   ├── error.go                  # 📝 UPDATED: Error responses (use platform-shared/httpx/errors.go)
│   │       │   └── pagination.go             # 📝 UPDATED: Pagination (use platform-shared/httpx/pagination.go)
│   │       ├── routes/                       # 🆕 Route registrars mapped 1:1 to handlers
│   │       │   ├── file_routes.go            # /files/*
│   │       │   ├── upload_routes.go          # /upload/*
│   │       │   ├── download_routes.go        # /download/*
│   │       │   ├── folder_routes.go          # /folders/*
│   │       │   ├── share_routes.go           # /shares/*
│   │       │   ├── media_routes.go           # /media/*
│   │       │   ├── flag_routes.go            # /flags/*
│   │       │   ├── policy_routes.go          # 🆕 /policies/*
│   │       │   ├── lifecycle_routes.go       # 🆕 /lifecycle/*
│   │       │   └── linking_routes.go         # 🆕 /links/* (signed URLs, audit logs)
│   │       └── router.go                     # 📝 UPDATED: HTTP router setup (uses Gin, applies platform-shared/ginx middleware)
│   │
│   └── config/                               # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                         # 🆕 Typed Config struct (App, Server, Postgres, Kafka, Redis, MinIO, Storage)
│       ├── loader.go                         # 🆕 Config loader using Viper (precedence: flags → env → file → defaults)
│       └── docs/
│           └── CONFIGURATION.md              # 🆕 Configuration documentation (all ENV vars, defaults, examples)
│
├── config/
│   ├── default.yaml
│   ├── dev.yaml
│   └── prod.yaml
│
├── dapr/
│   ├── local/
│   │   ├── pubsub.yaml
│   │   └── statestore.yaml
│   └── k8s/
│       ├── pubsub.yaml                       # Kafka with scopes: ["storage-be"]
│       ├── statestore.yaml                   # State store with scopes
│       └── secrets.yaml                      # Dapr secret store
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                         # Service-specific errors
│   │   └── codes.go                          # Error codes
│   ├── logger/
│   │   └── logger.go                         # ❌ REMOVED: Use platform-shared/logging
│   ├── utils/
│   │   ├── validator.go                      # Local validation utilities
│   │   ├── file_utils.go                     # File utilities (path manipulation, extension extraction)
│   │   ├── mime_detector.go                  # MIME type detection
│   │   └── hash.go                           # File hash calculation (MD5, SHA256)
│   └── constants/
│       ├── events.go                         # ❌ REMOVED: Use contracts/events
│       ├── topics.go                         # ❌ REMOVED: Use contracts/events
│       └── mime_types.go                     # Supported MIME types
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       ├── pvc.yaml
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   ├── seed-data.sh
│   └── cleanup-orphaned.sh
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── api.md
│   ├── events.md                             # 📝 UPDATED: Events (published: file.uploaded, file.deleted, consumed: user.created, etc.)
│   ├── upload-flow.md
│   ├── media-processing.md
│   ├── MIGRATIONS.md
│   ├── SCHEMA.md
│   └── RUNBOOK.md
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md

```
---

## **📦8️⃣ search-be (UPDATED WITH ENHANCED RECOMMENDATIONS)**

```
apps/be/search-be/
│
├── cmd/
│   └── api/
│       └── main.go                           # 📝 UPDATED: Application entry point - initializes Gin, Dapr, connects to Elasticsearch (now uses loadConfig from internal/config, platform-shared/logging)
│
├── internal/
│   ├── domain/
│   │   ├── search_index/
│   │   │   ├── entity.go                     # Search index metadata
│   │   │   ├── job_index.go                  # Job search document
│   │   │   ├── user_index.go                 # User/Freelancer document
│   │   │   ├── errors.go                     # 🆕 Index errors (IndexNotFound, DocumentConflict)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: search.document.indexed / search.document.reindexed / search.index.visibility.changed / search.index.archived
│   │   │
│   │   ├── search_query/
│   │   │   ├── entity.go                     # Search query logs
│   │   │   ├── filters.go                    # Search filters
│   │   │   ├── errors.go                     # 🆕 Query errors (InvalidFilter, BadSort)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: search.query.logged / search.query.alert.triggered
│   │   │
│   │   ├── recommendation/
│   │   │   ├── entity.go                     # Recommendation records
│   │   │   ├── score.go                      # Scoring algorithm
│   │   │   ├── reason.go                     # Recommendation reasons
│   │   │   ├── feedback.go                   # User feedback on recommendations
│   │   │   ├── errors.go                     # 🆕 Recommendation errors (ModelUnavailable, ScoringFailed)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: search.recommendation.generated / search.recommendation.feedback.recorded
│   │   │
│   │   ├── recommendation_model/
│   │   │   ├── entity.go                     # ML model metadata
│   │   │   ├── feature.go                    # Feature vectors
│   │   │   ├── training_data.go              # Training data
│   │   │   ├── errors.go                     # 🆕 Model errors (VersionNotFound, FeatureMismatch)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: search.model.version.registered / search.model.deployed / search.model.rolled_back / search.model.training_data.updated
│   │   │
│   │   ├── user_preference/
│   │   │   ├── entity.go                     # User preferences for recommendations
│   │   │   ├── implicit_signals.go           # Implicit user signals (clicks, views)
│   │   │   ├── explicit_preferences.go       # Explicit preferences
│   │   │   ├── errors.go                     # 🆕 Preference errors (PreferenceNotFound)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: search.personalization.updated / search.preference.updated
│   │   │
│   │   ├── matching/
│   │   │   ├── entity.go                     # Job-Freelancer matches
│   │   │   ├── criteria.go                   # Matching criteria
│   │   │   ├── score_breakdown.go            # Detailed match scoring
│   │   │   ├── errors.go                     # 🆕 Matching errors (CriteriaInvalid)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: search.match.calculated / search.match.accepted / search.match.dismissed
│   │   │
│   │   ├── feed/
│   │   │   ├── entity.go                     # User feeds
│   │   │   ├── item.go                        # Feed items
│   │   │   ├── personalization.go             # Personalization data
│   │   │   ├── errors.go                      # 🆕 Feed errors (FeedNotFound)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: search.feed.item.added / search.feed.item.removed / search.feed.updated
│   │   │
│   │   ├── trending/
│   │   │   ├── entity.go                      # Trending jobs/skills
│   │   │   ├── calculator.go                  # Calculate trending items
│   │   │   ├── errors.go                      # 🆕 Trending errors
│   │   │   ├── repository.go
│   │   │   └── events.go                      # 🆕 Domain events: search.trending.calculated / search.trending.updated
│   │   │
│   │   ├── saved_search/
│   │   │   ├── entity.go                      # Saved searches
│   │   │   ├── alert.go                        # Search alerts
│   │   │   ├── errors.go                       # 🆕 Saved search errors
│   │   │   ├── repository.go
│   │   │   └── events.go                       # 🆕 Domain events: search.saved_search.created / search.saved_search.updated / search.saved_search.deleted / search.saved_search.alert.sent
│   │   │
│   │   ├── similarity/
│   │   │   ├── entity.go                      # Similar jobs/users
│   │   │   ├── vector.go                      # Similarity vectors
│   │   │   ├── errors.go                      # 🆕 Similarity errors
│   │   │   ├── repository.go
│   │   │   └── events.go                      # 🆕 Domain events: search.similarity.computed / search.similarity.model.updated
│   │   │
│   │   ├── taxonomy/                          # 🆕 Taxonomy & synonyms
│   │   │   ├── entity.go                       # Normalized skills/categories (SkillID, Name, CategoryID, Aliases)
│   │   │   ├── category.go                     # Categories (ID, Name, ParentID, Path)
│   │   │   ├── synonym.go                      # Synonyms & alias mapping (SkillID, Alias, ToleranceScore)
│   │   │   ├── typos.go                        # Typo tolerance rules (edit distance thresholds)
│   │   │   ├── errors.go                        # Taxonomy errors (SkillNotFound, AliasConflict)
│   │   │   ├── repository.go                    # TaxonomyRepository interface
│   │   │   └── events.go                        # 🆕 Domain events: search.taxonomy.updated / search.taxonomy.synonym.changed / search.taxonomy.typo_rules.updated
│   │   │
│   │   ├── ltr/                                # 🆕 Learning-to-rank hooks
│   │   │   ├── entity.go                       # LTR signals (docID, conversions, responseSpeed, jobSuccess, intent)
│   │   │   ├── signal_source.go                # Signal sources (events → features)
│   │   │   ├── feature_store.go                # Persisted LTR features (effectiveAt timestamps)
│   │   │   ├── errors.go                       # LTR errors (SignalOutOfRange, FeatureStale)
│   │   │   ├── repository.go                   # LTRRepository interface
│   │   │   └── events.go                       # 🆕 Domain events: search.ltr.signal.recorded / search.ltr.features.updated
│   │   │
│   │   ├── facets/                             # 🆕 Facets & filters
│   │   │   ├── entity.go                       # Facet definitions (price bands, availability, tz overlap, languages, badges)
│   │   │   ├── banding.go                      # Price/rate band logic
│   │   │   ├── tz_overlap.go                   # Timezone overlap calculator
│   │   │   ├── errors.go                       # Facet errors (InvalidBand, UnknownBadge)
│   │   │   ├── repository.go                   # FacetRepository interface
│   │   │   └── events.go                       # 🆕 Domain events: search.facets.definition.updated / search.facets.banding.updated / search.facets.tz_overlap.updated
│   │   │
│   │   ├── personalization/                    # 🆕 Personalization signals
│   │   │   ├── entity.go                       # Profiles (recent activity, saved searches, preferred rates)
│   │   │   ├── cold_start.go                   # Cold-start defaults
│   │   │   ├── errors.go                       # Personalization errors
│   │   │   ├── repository.go                   # PersonalizationRepository interface
│   │   │   └── events.go                       # 🆕 Domain events: search.personalization.updated / search.personalization.cold_start.applied
│   │   │
│   │   ├── hygiene/                            # 🆕 Index hygiene
│   │   │   ├── entity.go                       # Hygiene tasks (incremental updates, dedup, archival, visibility)
│   │   │   ├── incremental.go                  # Incremental update markers (version, changed fields)
│   │   │   ├── dedupe.go                       # Dedupe fingerprints
│   │   │   ├── visibility.go                   # Visibility states
│   │   │   ├── errors.go                       # Hygiene errors
│   │   │   ├── repository.go                   # HygieneRepository interface
│   │   │   └── events.go                       # 🆕 Domain events: search.index.hygiene.deduplicated / search.index.hygiene.archived / search.index.hygiene.visibility.changed
│   │   │
│   │   └── outbox/
│   │       ├── entity.go                       # ❌ REMOVED (use platform-shared/outbox/entity.go)
│   │       └── repository.go                    # ❌ REMOVED (use platform-shared/outbox/repository.go)
│   │
│   ├── application/
│   │   ├── eventhandler/
│   │   │   ├── job_handler.go                 # Consumes: job.posted | job.updated | job.closed
│   │   │   │                                     # Purpose: create/update/remove job documents in the search index; refresh facets.
│   │   │   │
│   │   │   ├── user_handler.go                # Consumes: user.updated
│   │   │   │                                     # Purpose: refresh user/freelancer search documents (profile fields, visibility, badges).
│   │   │   │
│   │   │   ├── review_handler.go              # Consumes: review.double_blind.published | review.published
│   │   │   │                                     # Purpose: update public rating aggregates on indexed user/job docs; recalc ranking features.
│   │   │   │
│   │   │   ├── entitlement_handler.go         # Consumes: subscription.feature.changed
│   │   │   │                                     # Purpose: adjust facet/visibility rules and gated search features tied to plan changes.
│   │   │   │
│   │   │   ├── admin_content_handler.go       # Consumes: admin.content.actioned (approve/reject/remove/hide)
│   │   │   │                                     # Purpose: hide/unhide or restore affected indexed documents promptly.
│   │   │   │
│   │   │   ├── admin_flags_handler.go         # Consumes: admin.feature_flag.updated | admin.threshold.updated | admin.experiment.updated
│   │   │   │                                     # Purpose: refresh runtime toggles that control search facets, boosts, filters, and experiments.
│   │   │   │
│   │   │   ├── storage_lifecycle_handler.go   # Consumes: file.lifecycle.soft_deleted | file.lifecycle.restored
│   │   │   │                                     # Purpose: sync file-backed assets’ visibility in indexed documents (thumbnails, attachments).
│   │   │   │
│   │   │   └── _partitioning_notes.go         # (comments only) Partition keys used by handlers:
│   │   │                                         #   job.* → key by job_id
│   │   │                                         #   user.* → key by user_id
│   │   │                                         #   review.* → key by subject_user_id (or contract_id when provided)
│   │   │                                         #   admin.* → key by affected_resource_id
│   │   │                                         #   file.lifecycle.* → key by file_id
│   │   ├── search/
│   │   │   ├── service.go
│   │   │   ├── job_search.go
│   │   │   ├── freelancer_search.go
│   │   │   ├── query_builder.go
│   │   │   ├── facet_builder.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                    # 🆕 ExecuteSearch (jobs/freelancers), SaveSearch
│   │   │   ├── queries.go                     # 🆕 GetSearchResults, GetSearchSuggestions
│   │   │   └── validators.go                  # 🆕 Validate filters/sorts/facets
│   │   ├── indexing/
│   │   │   ├── service.go
│   │   │   ├── job_indexer.go
│   │   │   ├── user_indexer.go
│   │   │   ├── bulk_indexer.go
│   │   │   ├── dto.go
│   │   │   ├── commands.go                    # 🆕 IndexJob, IndexUser, ReindexAll
│   │   │   ├── queries.go                     # 🆕 GetIndexStatus, GetDocByID
│   │   │   └── validators.go                  # 🆕 Validate index payloads
│   │   ├── recommendation/
│   │   │   ├── service.go
│   │   │   ├── job_recommender.go             # Recommend jobs to freelancers
│   │   │   ├── freelancer_recommender.go      # Recommend freelancers to clients
│   │   │   ├── collaborative_filtering.go     # Collaborative filtering algorithm
│   │   │   ├── content_based.go               # Content-based filtering
│   │   │   ├── hybrid_recommender.go          # Hybrid recommendation approach
│   │   │   ├── scoring_engine.go              # Calculate recommendation scores
│   │   │   ├── personalization.go             # Personalization logic
│   │   │   ├── diversity_optimizer.go         # Ensure diverse recommendations
│   │   │   ├── cold_start_handler.go          # Handle new users/jobs
│   │   │   ├── dto.go
│   │   │   ├── ml_model.go                    # ML model integration
│   │   │   ├── commands.go                    # 🆕 GenerateRecommendations, RecordFeedback
│   │   │   ├── queries.go                     # 🆕 GetRecommendations, GetReasons
│   │   │   └── validators.go                  # 🆕 Validate rec limits, diversity params
│   │   ├── matching/
│   │   │   ├── service.go
│   │   │   ├── matcher.go                     # Job-Freelancer matching
│   │   │   ├── criteria_evaluator.go          # Evaluate match criteria
│   │   │   ├── skill_matcher.go               # Match based on skills
│   │   │   ├── experience_matcher.go          # Match based on experience
│   │   │   ├── rate_matcher.go                # Match based on rates
│   │   │   ├── availability_matcher.go        # Match based on availability
│   │   │   ├── score_calculator.go            # Calculate match score
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                    # 🆕 CreateMatchRun
│   │   │   ├── queries.go                     # 🆕 GetMatchesForJob, GetMatchesForUser
│   │   │   └── validators.go                  # 🆕 Validate criteria completeness
│   │   ├── feed/
│   │   │   ├── service.go
│   │   │   ├── generator.go                   # Generate personalized feed
│   │   │   ├── ranking.go                     # Rank feed items
│   │   │   ├── freshness_scorer.go            # Score by freshness
│   │   │   ├── relevance_scorer.go            # Score by relevance
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                    # 🆕 GenerateFeed
│   │   │   ├── queries.go                     # 🆕 GetFeed
│   │   │   └── validators.go                  # 🆕 Validate feed window/size
│   │   ├── trending/
│   │   │   ├── service.go
│   │   │   ├── calculator.go                  # Calculate trending items
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                    # 🆕 RecomputeTrending
│   │   │   ├── queries.go                     # 🆕 GetTrending
│   │   │   └── validators.go                  # 🆕 Validate window/minSupport
│   │   ├── similarity/
│   │   │   ├── service.go
│   │   │   ├── job_similarity.go              # Find similar jobs
│   │   │   ├── user_similarity.go             # Find similar users
│   │   │   ├── vector_calculator.go           # Calculate similarity vectors
│   │   │   ├── dto.go
│   │   │   ├── commands.go                    # 🆕 RebuildSimilarityVectors
│   │   │   ├── queries.go                     # 🆕 GetSimilarJobs, GetSimilarUsers
│   │   │   └── validators.go                  # 🆕 Validate k/threshold
│   │   ├── suggestion/
│   │   │   ├── service.go                     # Autocomplete suggestions
│   │   │   ├── dto.go
│   │   │   ├── cache_warmer.go
│   │   │   ├── commands.go                    # 🆕 WarmSuggestionCache
│   │   │   ├── queries.go                     # 🆕 GetSuggestions
│   │   │   └── validators.go                  # 🆕 Validate prefix/lang
│   │   ├── taxonomy/                          # 🆕 Taxonomy application layer
│   │   │   ├── service.go                      # Manage skills/categories, synonyms, typo rules
│   │   │   ├── commands.go                     # UpsertSkill, UpsertCategory, AddAlias, RemoveAlias
│   │   │   ├── queries.go                      # GetSkill, ListSkills, GetCategoryTree
│   │   │   ├── validators.go                   # Validate alias conflicts, edit distance
│   │   │   ├── dto.go                           # TaxonomyDTO, CategoryDTO, AliasDTO
│   │   │   └── mapper.go                        # Taxonomy mappers
│   │   ├── ltr/                                # 🆕 LTR application layer
│   │   │   ├── service.go                      # Ingest signals, build feature vectors, rerank
│   │   │   ├── commands.go                     # IngestSignalBatch, SnapshotFeatures
│   │   │   ├── queries.go                      # GetFeaturesForDoc, PreviewRerank
│   │   │   ├── validators.go                   # Validate signal bounds, effectiveAt
│   │   │   ├── dto.go                           # SignalDTO, FeatureDTO, RerankPreviewDTO
│   │   │   └── mapper.go                        # LTR mappers
│   │   ├── facets/                             # 🆕 Facets application layer
│   │   │   ├── service.go                      # Build facets, apply filters, banding
│   │   │   ├── commands.go                     # DefineFacet, UpdateFacetBands
│   │   │   ├── queries.go                      # GetFacetsForQuery
│   │   │   ├── validators.go                   # Validate bands, tz overlap rules
│   │   │   ├── dto.go                           # FacetDTO, BandDTO
│   │   │   └── mapper.go                        # Facet mappers
│   │   ├── personalization/                    # 🆕 Personalization application layer
│   │   │   ├── service.go                      # Update profile, compute defaults
│   │   │   ├── commands.go                     # UpdatePersonalizationProfile
│   │   │   ├── queries.go                      # GetPersonalizationProfile
│   │   │   ├── validators.go                   # Validate rate prefs, recent activity bounds
│   │   │   ├── dto.go                           # PersonalizationDTO
│   │   │   └── mapper.go                        # Personalization mappers
│   │   ├── hygiene/                            # 🆕 Index hygiene application layer
│   │   │   ├── service.go                      # Incremental updates, dedupe, archive/visibility
│   │   │   ├── commands.go                     # RunIncrementalUpdate, RunDedup, ArchiveDoc, SetVisibility
│   │   │   ├── queries.go                      # GetHygieneStatus, GetDocHistory
│   │   │   ├── validators.go                   # Validate dedupe fingerprint size, state transitions
│   │   │   ├── dto.go                           # HygieneTaskDTO, VisibilityDTO
│   │   │   └── mapper.go                        # Hygiene mappers
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go              # PostgreSQL connection setup (DSN from config, connection pooling)
│   │   │       ├── transaction.go             # Transaction helpers (Begin, Commit, Rollback, WithTransaction wrapper)
│   │   │       ├── migrations.go              # 📝 UPDATED: Auto-migrate logic (now with version tracking, GORM AutoMigrate for all tables)
│   │   │       ├── version.go                 # 🆕 Schema version tracking (SchemaVersion table, RecordMigration function)
│   │   │       ├── safety.go                  # 🆕 Pre-migration safety checks (environment validation, disk space check)
│   │   │       ├── search_query_repository.go
│   │   │       ├── recommendation_repository.go
│   │   │       ├── recommendation_model_repository.go
│   │   │       ├── user_preference_repository.go
│   │   │       ├── matching_repository.go
│   │   │       ├── feed_repository.go
│   │   │       ├── trending_repository.go
│   │   │       ├── saved_search_repository.go
│   │   │       ├── similarity_repository.go
│   │   │       ├── taxonomy_repository.go     # 🆕 TaxonomyRepository implementation
│   │   │       ├── ltr_repository.go          # 🆕 LTRRepository implementation
│   │   │       ├── facets_repository.go       # 🆕 FacetRepository implementation
│   │   │       ├── personalization_repository.go # 🆕 PersonalizationRepository implementation
│   │   │       └── hygiene_repository.go      # 🆕 HygieneRepository implementation
│   │   ├── elasticsearch/
│   │   │   ├── client.go                      # Elasticsearch client
│   │   │   ├── index_manager.go               # Index management
│   │   │   ├── job_mapper.go                  # Job mapping to ES document
│   │   │   ├── user_mapper.go                 # User mapping to ES document
│   │   │   └── config.go                      # ES configuration (hosts, auth)
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go              # Redis connection setup (connection pooling, retry logic)
│   │   │       ├── search_cache.go            # Cache search results (Get, Set, Invalidate with TTL)
│   │   │       ├── suggestion_cache.go        # Cache autocomplete suggestions
│   │   │       ├── feed_cache.go              # Cache user feeds
│   │   │       ├── recommendation_cache.go    # Cache recommendations
│   │   │       ├── taxonomy_cache.go          # 🆕 Cache skills/categories & aliases
│   │   │       └── ltr_cache.go               # 🆕 Cache LTR features snapshot
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go                # 📝 UPDATED: Kafka consumer (now uses platform-shared/inbox for message deduplication)
│   │   │       ├── producer.go                # 📝 UPDATED: Kafka producer (now uses platform-shared/outbox for reliable publishing)
│   │   │       ├── topics.go                  # 📝 UPDATED: Topic constants (now imported from contracts/events - job.indexed, user.indexed, recommendation.generated)
│   │   │       └── scram.go                   # SCRAM authentication for Kafka (SASL/SCRAM-SHA-256)
│   │   ├── ml/
│   │   │   ├── model_loader.go                # Load ML models
│   │   │   ├── predictor.go                   # Make predictions
│   │   │   ├── feature_extractor.go           # Extract features
│   │   │   ├── trainer.go                     # Train models
│   │   │   └── evaluator.go                   # Evaluate model performance
│   │   └── outbox/
│   │       ├── processor.go                   # ❌ REMOVED (use platform-shared/outbox/forwarder.go)
│   │       └── scheduler.go                   # ❌ REMOVED (use platform-shared/outbox/scheduler.go)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── search_handler.go          # Search HTTP handlers (POST /search/jobs, /search/freelancers)
│   │       │   ├── recommendation_handler.go  # Recommendation handlers (GET /recommendations/jobs, /recommendations/freelancers)
│   │       │   ├── matching_handler.go        # Matching handlers
│   │       │   ├── feed_handler.go            # Feed handlers
│   │       │   ├── trending_handler.go        # Trending handlers
│   │       │   ├── similarity_handler.go      # Similarity handlers
│   │       │   ├── suggestion_handler.go      # Autocomplete handlers
│   │       │   ├── indexing_handler.go        # Admin endpoint for indexing
│   │       │   ├── taxonomy_handler.go        # 🆕 Taxonomy handlers (skills/categories, aliases)
│   │       │   ├── ltr_handler.go             # 🆕 LTR handlers (ingest signals, preview rerank)
│   │       │   ├── facets_handler.go          # 🆕 Facet handlers (definitions, bands)
│   │       │   ├── personalization_handler.go # 🆕 Personalization handlers (get/update profile)
│   │       │   ├── hygiene_handler.go         # 🆕 Hygiene handlers (incrementals/dedup/visibility)
│   │       │   └── health_handler.go          # Health check endpoints (/health, /ready, /live)
│   │       ├── routes/                         # 🆕 Route registrars mapped 1:1 to handlers
│   │       │   ├── search_routes.go            # /search/*
│   │       │   ├── recommendation_routes.go    # /recommendations/*
│   │       │   ├── matching_routes.go          # /matching/*
│   │       │   ├── feed_routes.go              # /feed/*
│   │       │   ├── trending_routes.go          # /trending/*
│   │       │   ├── similarity_routes.go        # /similarity/*
│   │       │   ├── suggestion_routes.go        # /suggestions/*
│   │       │   ├── indexing_routes.go          # /indexing/*
│   │       │   ├── taxonomy_routes.go          # 🆕 /taxonomy/*
│   │       │   ├── ltr_routes.go               # 🆕 /ltr/*
│   │       │   ├── facets_routes.go            # 🆕 /facets/*
│   │       │   ├── personalization_routes.go   # 🆕 /personalization/*
│   │       │   └── hygiene_routes.go           # 🆕 /hygiene/*
│   │       ├── middleware/
│   │       │   ├── auth.go                     # 📝 UPDATED: Authentication middleware (uses pkg/auth for JWT verification)
│   │       │   ├── rbac.go                     # 📝 UPDATED: RBAC middleware (uses pkg/auth authorizer for role checks)
│   │       │   ├── cors.go                     # 📝 UPDATED: CORS middleware (uses platform-shared/ginx CORS)
│   │       │   ├── rate_limit.go               # Rate limiting middleware (token bucket algorithm)
│   │       │   ├── logging.go                  # 📝 UPDATED: Logging middleware (structured logs per request, uses platform-shared/ginx/logging)
│   │       │   ├── error_handler.go            # Error handling middleware
│   │       │   └── request_id.go               # 📝 UPDATED: Request ID middleware (X-Request-ID header, uses platform-shared/ginx/requestid)
│   │       ├── responses/
│   │       │   ├── success.go                  # 📝 UPDATED: Success response wrappers (use platform-shared/httpx/response.go)
│   │       │   ├── error.go                    # 📝 UPDATED: Error responses (use platform-shared/httpx/errors.go)
│   │       │   └── pagination.go               # 📝 UPDATED: Pagination (use platform-shared/httpx/pagination.go)
│   │       └── router.go                       # 📝 UPDATED: HTTP router setup (uses Gin, applies platform-shared/ginx middleware)
│   │
│   └── config/                                 # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                           # 🆕 Typed Config struct (App, Server, Postgres, Kafka, Redis, Elasticsearch, ML)
│       ├── loader.go                           # 🆕 Config loader using Viper (precedence: flags → env → file → defaults)
│       └── docs/
│           └── CONFIGURATION.md                # 🆕 Configuration documentation (all ENV vars, defaults, examples)
│
├── config/                                     # 🆕 MEDIUM - Configuration files
│   ├── default.yaml                            # 🆕 Default configuration
│   ├── dev.yaml                                # 🆕 Development overrides
│   └── prod.yaml                               # 🆕 Production overrides
│
├── dapr/                                       # 🆕 MEDIUM - Dapr components split by environment
│   ├── local/
│   │   ├── pubsub.yaml                         # Kafka pub/sub component
│   │   └── statestore.yaml                     # State store component
│   └── k8s/
│       ├── pubsub.yaml                         # Kafka with scopes: ["search-be"]
│       ├── statestore.yaml                     # State store with scopes
│       └── secrets.yaml                        # Dapr secret store
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                           # Service-specific errors
│   │   └── codes.go                            # Error codes
│   ├── logger/
│   │   └── logger.go                           # ❌ REMOVED: Use platform-shared/logging
│   ├── utils/
│   │   ├── validator.go                        # Local validation utilities
│   │   ├── text_analyzer.go                    # Text analysis utilities
│   │   ├── normalizer.go                       # Data normalization
│   │   └── vector_math.go                      # Vector math for similarity
│   └── constants/
│       ├── events.go                           # ❌ REMOVED: Use contracts/events
│       ├── topics.go                           # ❌ REMOVED: Use contracts/events
│       └── indices.go                          # Index names constants
│
├── elasticsearch/
│   ├── mappings/
│   │   ├── jobs.json                           # Job index mapping
│   │   └── users.json                          # User index mapping
│   └── analyzers/
│       └── custom_analyzers.json               # Custom analyzers
│
├── ml_models/
│   ├── job_recommendation/
│   │   ├── model.pkl                           # ML model for job recommendations
│   │   ├── features.json                       # Features list
│   │   └── metadata.json                       # Model metadata
│   ├── freelancer_recommendation/
│   │   ├── model.pkl                           # ML model for freelancer recommendations
│   │   ├── features.json                       # Features list
│   │   └── metadata.json                       # Model metadata
│   └── matching/
│       ├── model.pkl                           # ML model for matching
│       ├── features.json                       # Features list
│       └── metadata.json                       # Model metadata
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                     # Kubernetes Deployment
│       ├── service.yaml                        # Kubernetes Service
│       ├── configmap.yaml                      # ConfigMap
│       ├── secrets.yaml                        # Secrets
│       ├── hpa.yaml                            # HPA
│       ├── pdb.yaml                            # PDB
│       └── servicemonitor.yaml                 # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                          # Setup local environment
│   ├── get-secrets.sh                          # Fetch secrets
│   ├── seed-data.sh                            # Seed test data
│   ├── create-indices.sh                       # Create ES indices
│   ├── reindex-all.sh                          # Reindex all data
│   └── train-models.sh                         # Train ML models
│
├── tests/
│   ├── unit/                                   # Unit tests
│   ├── integration/                            # Integration tests
│   └── e2e/                                    # End-to-end tests
│
├── docs/
│   ├── README.md                               # Service overview
│   ├── api.md                                  # API documentation
│   ├── events.md                               # 📝 UPDATED: Events (published: job.indexed, user.indexed, recommendation.generated, consumed: job.posted, user.updated, etc.)
│   ├── search-algorithms.md                    # Search algorithms
│   ├── recommendation-engine.md                # Recommendation engine details
│   ├── recommendation-types.md                 # Recommendation types
│   ├── matching-algorithm.md                   # Matching algorithm
│   ├── ml-models.md                            # ML models documentation
│   ├── elasticsearch-setup.md                  # Elasticsearch setup
│   ├── MIGRATIONS.md                           # 🆕 Migration history
│   ├── SCHEMA.md                               # 🆕 Database schema
│   └── RUNBOOK.md                              # 🆕 Operational procedures
│
├── .github/
│   └── workflows/
│       ├── ci.yml                              # CI workflow
│       └── cd.yml                              # CD workflow
│
├── go.mod                                      # 📝 UPDATED: Imports pkg/auth, platform-shared, contracts/events
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md

```

---

## **📦9️⃣ reviews-be (UPDATED - COVERS RATINGS)**

```
apps/be/reviews-be/
│
├── cmd/
│   └── api/
│       └── main.go                           # 📝 UPDATED: Application entry point - initializes Gin, Dapr, connects to Postgres (now uses loadConfig from internal/config, platform-shared/logging)
│
├── internal/
│   ├── domain/
│   │   ├── review/
│   │   │   ├── entity.go                     # Review aggregate root
│   │   │   ├── rating.go                     # Rating breakdown (overall, quality, communication, etc.)
│   │   │   ├── enums.go                      # ReviewType, Status, ReviewCategory
│   │   │   ├── response.go                   # Review responses
│   │   │   ├── helpful.go                    # Helpful votes
│   │   │   ├── errors.go                     # 🆕 Domain errors (ReviewNotFound, WindowClosed, AlreadySubmitted)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: ReviewSubmitted, ReviewEdited, ReviewPublished, ReviewWithdrawn
│   │   ├── rating/
│   │   │   ├── entity.go                     # Rating system
│   │   │   ├── criteria.go                   # Rating criteria
│   │   │   ├── aggregation.go                # Aggregate ratings
│   │   │   ├── errors.go                     # 🆕 Rating errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: RatingCriteriaUpdated, RatingAggregated
│   │   ├── badge/
│   │   │   ├── entity.go                     # Achievement badges
│   │   │   ├── criteria.go                   # Badge criteria
│   │   │   ├── types.go                      # Badge types (Top Rated, Rising Talent, etc.)
│   │   │   ├── level.go                      # Badge levels
│   │   │   ├── errors.go                     # 🆕 Badge errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: BadgeCreated, BadgeCriteriaUpdated, BadgeLevelUpdated, BadgeArchived
│   │   ├── user_badge/
│   │   │   ├── entity.go                     # User badge assignments
│   │   │   ├── achievement_date.go           # When badge was earned
│   │   │   ├── errors.go                     # 🆕 UserBadge errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: BadgeAwardedToUser, BadgeRevokedFromUser
│   │   ├── reputation/
│   │   │   ├── entity.go                     # Reputation scores
│   │   │   ├── score_calculator.go           # Score calculation
│   │   │   ├── history.go                    # Reputation history
│   │   │   ├── decay.go                      # 🆕 Rolling averages, decay, outlier dampening
│   │   │   ├── top_rated_rules.go            # 🆕 “Top-rated” eligibility rules
│   │   │   ├── errors.go                     # 🆕 Reputation errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: ReputationUpdated, TopRatedEligibilityChanged
│   │   ├── feedback/
│   │   │   ├── entity.go                     # Private feedback
│   │   │   ├── category.go                   # Feedback categories
│   │   │   ├── nps.go                        # 🆕 Confidential NPS-style scores
│   │   │   ├── errors.go                     # 🆕 Feedback errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: PrivateFeedbackSubmitted, PrivateFeedbackUpdated
│   │   ├── flag/
│   │   │   ├── entity.go                     # Flagged reviews
│   │   │   ├── reason.go                     # Flag reasons
│   │   │   ├── errors.go                     # 🆕 Flag errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: ReviewFlagSubmitted, ReviewFlagResolved, ReviewFlagDismissed
│   │   ├── stats/
│   │   │   ├── entity.go                     # Review statistics
│   │   │   ├── aggregates.go                 # Aggregated stats
│   │   │   ├── errors.go                     # 🆕 Stats errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: ReviewStatsUpdated, ReviewStatsRecalculated
│   │   ├── review_reminder/
│   │   │   ├── entity.go                     # Review reminders
│   │   │   ├── errors.go                     # 🆕 Reminder errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: ReviewReminderScheduled, ReviewReminderSent, ReviewReminderFailed
│   │   ├── double_blind/                     # 🆕 Double-blind reviews domain
│   │   │   ├── entity.go                     # Windows (contractID, openAt, closeAt, submitted{client,freelancer})
│   │   │   ├── window.go                     # Publish after both submit or timeout
│   │   │   ├── policy.go                     # Rules (durations, grace periods)
│   │   │   ├── errors.go                     # Double-blind errors (WindowNotOpen, PublishBlocked)
│   │   │   ├── repository.go                 # DoubleBlindRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: DoubleBlindSubmitted, DoubleBlindPublished, DoubleBlindTimeout
│   │   ├── weighting/                        # 🆕 Dimensions & weights
│   │   │   ├── entity.go                     # Weight profiles per category (quality, comms, timeliness)
│   │   │   ├── profile.go                    # Category-specific weighting resolution
│   │   │   ├── errors.go                     # Weighting errors (ProfileNotFound)
│   │   │   ├── repository.go                 # WeightingRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: WeightProfileCreated, WeightProfileUpdated
│   │   ├── public_response/                  # 🆕 Public response rights
│   │   │   ├── entity.go                     # Right-of-reply records (reviewID, authorID, body, postedAt)
│   │   │   ├── policy.go                     # Response windows/length limits
│   │   │   ├── errors.go                     # Public response errors
│   │   │   ├── repository.go                 # PublicResponseRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: PublicResponsePosted, PublicResponseEdited, PublicResponseRemoved
│   │   ├── moderation/                       # 🆕 Abuse controls
│   │   │   ├── entity.go                     # Moderation states (Pending, UnderReview, Removed, Restored)
│   │   │   ├── auto_flag.go                  # Auto-flags: swapping, bursts, suspicious patterns
│   │   │   ├── signals.go                    # Detection signals
│   │   │   ├── errors.go                     # Moderation errors
│   │   │   ├── repository.go                 # ModerationRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: ReviewAbuseAutoFlagged, ReviewModerationStateChanged
│   │   └── outbox/
│   │       ├── entity.go                     # ❌ REMOVED (use platform-shared/outbox/entity.go)
│   │       └── repository.go                 # ❌ REMOVED (use platform-shared/outbox/repository.go)
│   │
│   ├── application/
│   │   ├── eventhandler/
│   │   │   ├── contract_handler.go            # Consumes: contract.ended
│   │   │   │                                     # Purpose: opens the double-blind review window for both parties at contract end.
│   │   │   │
│   │   │   ├── payment_handler.go             # Consumes: payment.captured
│   │   │   │                                     # Purpose: starts/qualifies the double-blind window when payment is captured (hourly/fixed).
│   │   │   │
│   │   │   ├── admin_moderation_handler.go    # Consumes: admin.moderation.actioned (approve/reject/remove/hide)
│   │   │   │                                     # Purpose: reflect moderation decisions on reviews (remove/restore visibility, recalc aggregates if needed).
│   │   │   │
│   │   │   ├── badge_handler.go               # Consumes: badge.criteria.met  # (optional downstream hook)
│   │   │   │                                     # Purpose: optional: feed reputation/badge eligibility workflows that affect review surfaces.
│   │   │   │
│   │   │   └── admin_flags_handler.go         # Consumes: admin.feature_flag.updated | admin.threshold.updated | admin.experiment.updated
│   │   │                                         # Purpose: refresh runtime toggles/config used by review windows, moderation thresholds,
│   │   │                                         #           decay/outlier dampening, and publishing policies.
│   │   ├── review/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Create, Update, Delete, Respond
│   │   │   ├── queries.go                    # Get, List, Filter
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── business_rules.go
│   │   ├── rating/
│   │   │   ├── service.go
│   │   │   ├── aggregator.go                 # Aggregate ratings
│   │   │   ├── calculator.go                 # Calculate average ratings
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                   # 🆕 UpsertCriteria, RecalculateForContract
│   │   │   ├── queries.go                    # 🆕 GetCriteria, GetAggregates
│   │   │   └── validators.go                 # 🆕 Validate criteria weights & bounds
│   │   ├── badge/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Award, Revoke badges
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── achievement_checker.go        # Check badge criteria
│   │   ├── reputation/
│   │   │   ├── service.go
│   │   │   ├── calculator.go                 # Calculate reputation
│   │   │   ├── updater.go                    # Update reputation
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                   # 🆕 RecomputeUserReputation, ApplyDecay
│   │   │   ├── queries.go                    # 🆕 GetReputation, GetHistory
│   │   │   └── validators.go                 # 🆕 Validate decay/outlier params
│   │   ├── feedback/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                   # 🆕 SubmitPrivateFeedback
│   │   │   ├── queries.go                    # 🆕 GetPrivateFeedback
│   │   │   └── validators.go                 # 🆕 Validate NPS bounds & confidentiality
│   │   ├── stats/
│   │   │   ├── service.go
│   │   │   ├── aggregator.go                 # Aggregate stats
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                   # 🆕 RebuildStats
│   │   │   ├── queries.go                    # 🆕 GetStats
│   │   │   └── validators.go                 # 🆕 Validate windows & min samples
│   │   ├── double_blind/                     # 🆕 App layer for double-blind reviews
│   │   │   ├── service.go                    # OpenWindow, SubmitSide, PublishIfReady
│   │   │   ├── commands.go                   # OpenWindow, SubmitClient, SubmitFreelancer, ForcePublish
│   │   │   ├── queries.go                    # GetWindow, GetPublishState
│   │   │   └── validators.go                 # Validate window times & state transitions
│   │   ├── weighting/                        # 🆕 App layer for dimensions & weights
│   │   │   ├── service.go                    # ResolveWeightProfile, UpsertProfile
│   │   │   ├── commands.go                   # UpsertWeightProfile
│   │   │   ├── queries.go                    # GetWeightProfile, ListProfiles
│   │   │   └── validators.go                 # Validate weight sums & ranges
│   │   ├── public_response/                  # 🆕 App layer for right-of-reply
│   │   │   ├── service.go                    # PostResponse, UpdateResponse
│   │   │   ├── commands.go                   # CreateResponse, EditResponse
│   │   │   ├── queries.go                    # GetResponse
│   │   │   └── validators.go                 # Validate length & window
│   │   ├── moderation/                       # 🆕 Abuse controls
│   │   │   ├── service.go                    # AutoFlag, ReviewModeration, ChangeState
│   │   │   ├── commands.go                   # AutoFlagReview, SetModerationState
│   │   │   ├── queries.go                    # GetModerationState, ListFlags
│   │   │   └── validators.go                 # Validate reasons & thresholds
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection setup (DSN from config, connection pooling)
│   │   │       ├── transaction.go            # Transaction helpers (Begin, Commit, Rollback, WithTransaction wrapper)
│   │   │       ├── migrations.go             # 📝 UPDATED: Auto-migrate logic (now with version tracking, GORM AutoMigrate for all tables)
│   │   │       ├── version.go                # 🆕 Schema version tracking (SchemaVersion table, RecordMigration function)
│   │   │       ├── safety.go                 # 🆕 Pre-migration safety checks (environment validation, disk space check)
│   │   │       ├── review_repository.go
│   │   │       ├── rating_repository.go
│   │   │       ├── badge_repository.go
│   │   │       ├── user_badge_repository.go
│   │   │       ├── reputation_repository.go
│   │   │       ├── feedback_repository.go
│   │   │       ├── flag_repository.go
│   │   │       ├── stats_repository.go
│   │   │       ├── review_reminder_repository.go
│   │   │       ├── double_blind_repository.go # 🆕 DoubleBlindRepository implementation
│   │   │       ├── weighting_repository.go    # 🆕 WeightingRepository implementation
│   │   │       ├── public_response_repository.go # 🆕 PublicResponseRepository implementation
│   │   │       └── moderation_repository.go   # 🆕 ModerationRepository implementation
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go             # Redis connection setup (connection pooling, retry logic)
│   │   │       ├── review_cache.go           # Cache reviews (Get, Set, Invalidate with TTL)
│   │   │       ├── badge_cache.go            # Cache badges
│   │   │       ├── reputation_cache.go       # Cache reputation scores
│   │   │       ├── stats_cache.go            # Cache stats
│   │   │       ├── double_blind_cache.go     # 🆕 Cache windows/publish state
│   │   │       └── weighting_cache.go        # 🆕 Cache weight profiles
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go               # 📝 UPDATED: Kafka consumer (now uses platform-shared/inbox for message deduplication)
│   │   │       ├── producer.go               # 📝 UPDATED: Kafka producer (now uses platform-shared/outbox for reliable publishing)
│   │   │       ├── topics.go                 # 📝 UPDATED: Topic constants (imported from contracts/events - review.submitted, review.published, badge.awarded)
│   │   │       └── scram.go                  # SCRAM authentication for Kafka (SASL/SCRAM-SHA-256)
│   │   └── outbox/
│   │       ├── processor.go                  # ❌ REMOVED (use platform-shared/outbox/forwarder.go)
│   │       └── scheduler.go                  # ❌ REMOVED (use platform-shared/outbox/scheduler.go)
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── review_handler.go         # Review HTTP handlers (POST /reviews, GET /reviews/:id)
│   │       │   ├── rating_handler.go         # Rating handlers
│   │       │   ├── badge_handler.go          # Badge handlers
│   │       │   ├── reputation_handler.go     # Reputation handlers
│   │       │   ├── feedback_handler.go       # Feedback handlers
│   │       │   ├── stats_handler.go          # Stats handlers
│   │       │   ├── flag_handler.go           # Flag handlers
│   │       │   ├── double_blind_handler.go   # 🆕 Double-blind endpoints
│   │       │   ├── weighting_handler.go      # 🆕 Weight profiles endpoints
│   │       │   ├── public_response_handler.go# 🆕 Right-of-reply endpoints
│   │       │   └── moderation_handler.go     # 🆕 Moderation endpoints
│   │       ├── routes/                       # 🆕 Route registrars mapped 1:1 to handlers
│   │       │   ├── review_routes.go          # /reviews/*
│   │       │   ├── rating_routes.go          # /ratings/*
│   │       │   ├── badge_routes.go           # /badges/*
│   │       │   ├── reputation_routes.go      # /reputation/*
│   │       │   ├── feedback_routes.go        # /feedback/*
│   │       │   ├── stats_routes.go           # /stats/*
│   │       │   ├── flag_routes.go            # /flags/*
│   │       │   ├── double_blind_routes.go    # 🆕 /double-blind/*
│   │       │   ├── weighting_routes.go       # 🆕 /weights/*
│   │       │   ├── public_response_routes.go # 🆕 /responses/*
│   │       │   └── moderation_routes.go      # 🆕 /moderation/*
│   │       ├── middleware/
│   │       │   ├── auth.go                   # 📝 UPDATED: Authentication middleware (uses pkg/auth for JWT verification)
│   │       │   ├── rbac.go                   # 📝 UPDATED: RBAC middleware (uses pkg/auth authorizer for role checks)
│   │       │   ├── cors.go                   # 📝 UPDATED: CORS middleware (uses platform-shared/ginx CORS)
│   │       │   ├── rate_limit.go             # Rate limiting middleware (token bucket algorithm)
│   │       │   ├── logging.go                # 📝 UPDATED: Logging middleware (structured logs per request, uses platform-shared/ginx/logging)
│   │       │   ├── error_handler.go          # Error handling middleware
│   │       │   └── request_id.go             # 📝 UPDATED: Request ID middleware (X-Request-ID header, uses platform-shared/ginx/requestid)
│   │       ├── responses/
│   │       │   ├── success.go                # 📝 UPDATED: Success response wrappers (use platform-shared/httpx/response.go)
│   │       │   ├── error.go                  # 📝 UPDATED: Error responses (use platform-shared/httpx/errors.go)
│   │       │   └── pagination.go             # 📝 UPDATED: Pagination (use platform-shared/httpx/pagination.go)
│   │       └── router.go                     # 📝 UPDATED: HTTP router setup (uses Gin, applies platform-shared/ginx middleware)
│   │
│   └── config/                               # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                         # 🆕 Typed Config struct (App, Server, Postgres, Kafka, Redis)
│       ├── loader.go                         # 🆕 Config loader using Viper (precedence: flags → env → file → defaults)
│       └── docs/
│           └── CONFIGURATION.md              # 🆕 Configuration documentation (all ENV vars, defaults, examples)
│
├── config/
│   ├── default.yaml                          # 🆕 Default configuration
│   ├── dev.yaml                              # 🆕 Development overrides
│   └── prod.yaml                             # 🆕 Production overrides
│
├── dapr/
│   ├── local/
│   │   ├── pubsub.yaml                       # Kafka pub/sub component
│   │   └── statestore.yaml                   # State store component
│   └── k8s/
│       ├── pubsub.yaml                       # Kafka with scopes: ["reviews-be"]
│       ├── statestore.yaml                   # State store with scopes
│       └── secrets.yaml                      # Dapr secret store
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                         # Service-specific errors
│   │   └── codes.go                          # Error codes
│   ├── logger/
│   │   └── logger.go                         # ❌ REMOVED: Use platform-shared/logging
│   ├── utils/
│   │   ├── validator.go                      # Local validation utilities
│   │   ├── sanitizer.go                      # Sanitize review content (prevent XSS)
│   │   └── sentiment_analyzer.go             # Sentiment analysis for reviews
│   └── constants/
│       ├── events.go                         # ❌ REMOVED: Use contracts/events
│       ├── topics.go                         # ❌ REMOVED: Use contracts/events
│       └── badges.go                         # Badge constants
│
├── seeds/
│   └── badges.sql                            # Seed initial badges
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   ├── seed-badges.sh
│   └── recalculate-reputation.sh
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── api.md
│   ├── events.md                             # 📝 UPDATED: (review.submitted, review.published, badge.awarded, reputation.updated)
│   ├── badge-system.md
│   ├── reputation-algorithm.md
│   ├── rating-system.md
│   ├── double-blind.md                       # 🆕 Double-blind policy & flows
│   ├── moderation.md                          # 🆕 Abuse controls & moderation states
│   ├── MIGRATIONS.md                          # 🆕
│   ├── SCHEMA.md                              # 🆕
│   └── RUNBOOK.md                             # 🆕
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md

```


---

## **📦🔟 subscriptions-be (UPDATED)**

```
apps/be/subscriptions-be/
│
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── domain/
│   │   ├── plan/
│   │   │   ├── entity.go                     # Subscription plans
│   │   │   ├── features.go                   # Plan features
│   │   │   ├── pricing.go                    # Pricing tiers
│   │   │   ├── limits.go                     # Plan limits
│   │   │   ├── errors.go                     # 🆕 Domain errors (PlanNotFound, InvalidFeature, LimitExceeded)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: plan.created / plan.updated / plan.archived
│   │   ├── subscription/
│   │   │   ├── entity.go                     # User subscriptions
│   │   │   ├── billing_cycle.go              # Billing cycle management
│   │   │   ├── enums.go                      # Status, Type
│   │   │   ├── auto_renewal.go               # Auto-renewal logic
│   │   │   ├── errors.go                     # 🆕 Subscription errors (AlreadyActive, PastDue, Canceled)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: subscription.created / subscription.changed / subscription.paused / subscription.canceled / subscription.renewed
│   │   ├── connect/
│   │   │   ├── entity.go                     # Connects/Credits
│   │   │   ├── package.go                    # Connect packages
│   │   │   ├── transaction.go                # Connect transactions
│   │   │   ├── balance.go                    # Balance tracking
│   │   │   ├── expiry.go                     # 🆕 Expiry & rollover rules (effective/expiry windows)
│   │   │   ├── grant.go                      # 🆕 Promotional grants (campaign/ref IDs)
│   │   │   ├── errors.go                     # 🆕 Connect errors (InsufficientBalance, Expired, RolloverNotAllowed)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: connects.purchased / connects.used / connects.expired / connects.granted
│   │   ├── usage/
│   │   │   ├── entity.go                     # Usage tracking
│   │   │   ├── quota.go                      # Usage quotas
│   │   │   ├── limit.go                      # Usage limits
│   │   │   ├── meter.go                      # 🆕 Per-feature meters (messages_to_non_hires, boosts, invites)
│   │   │   ├── errors.go                     # 🆕 Usage errors (LimitReached, CounterNotFound)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: usage.meter.incremented / usage.limit.reached
│   │   ├── entitlement/                      # 🆕 Plans & entitlements / feature flags
│   │   │   ├── entity.go                     # Entitlement (User/OrgID, PlanID, FeatureKey, Allowed, Scope)
│   │   │   ├── rules.go                      # Rule evaluation (plan→features, addons, promos)
│   │   │   ├── errors.go                     # 🆕 Entitlement errors (FeatureDenied, NotInPlan)
│   │   │   ├── repository.go                 # EntitlementRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: entitlement.feature.enabled / entitlement.feature.disabled
│   │   ├── seat_billing/                     # 🆕 Seat-based billing (clients/agencies)
│   │   │   ├── entity.go                     # Seat allocation (SubscriptionID, Seats, Assigned, OveragePolicy)
│   │   │   ├── overage.go                    # Overage calculations (per-seat over cap)
│   │   │   ├── proration.go                  # Proration rules for mid-cycle changes
│   │   │   ├── invoice_export.go             # Invoice export model (→ Financial-BE)
│   │   │   ├── errors.go                     # 🆕 Seat billing errors
│   │   │   ├── repository.go                 # SeatBillingRepository interface
│   │   │   └── events.go                     # 🆕 Domain events: billing.proration.applied / billing.invoice.exported / seat.overage.incurred / seat.allocated
│   │   ├── addon/
│   │   │   ├── entity.go                     # Plan add-ons
│   │   │   ├── errors.go                     # 🆕 Addon errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: addon.added / addon.removed / addon.updated
│   │   ├── promotion/
│   │   │   ├── entity.go                     # Promotional codes
│   │   │   ├── discount.go                   # Discount rules
│   │   │   ├── usage_limit.go                # Usage limits for promos
│   │   │   ├── errors.go                     # 🆕 Promotion errors (InvalidCode, Exhausted, Ineligible)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: promo.created / promo.applied / promo.revoked / promo.exhausted
│   │   ├── trial/
│   │   │   ├── entity.go                     # Free trials
│   │   │   ├── eligibility.go                # Trial eligibility
│   │   │   ├── errors.go                     # 🆕 Trial errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: trial.started / trial.ended / trial.eligibility.updated
│   │   ├── billing_history/
│   │   │   ├── entity.go                     # Billing history
│   │   │   ├── errors.go                     # 🆕 Billing history errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: billing.invoice.generated / billing.payment.applied / billing.credit.issued
│   │   ├── feature_toggle/
│   │   │   ├── entity.go                     # Feature toggles per plan
│   │   │   ├── errors.go                     # 🆕 Feature toggle errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.feature_flag.updated (plan scope) / feature.toggle.enabled / feature.toggle.disabled
│   │   └── outbox/
│   │       ├── entity.go                     # ❌ REMOVED (use platform-shared/outbox/entity.go)
│   │       └── repository.go                 # ❌ REMOVED (use platform-shared/outbox/repository.go)
│   │
│   ├── application/
│   │   ├── eventhandler/
│   │   │   ├── financial_handler.go           # Consumes: payment.processed | payment.failed    (from financial-be)
│   │   │   │                                     # Purpose: post-payment workflows (activate/renew/pause as needed),
│   │   │   │                                     #          dunning/reactivation on failures.
│   │   │   │
│   │   │   ├── proposal_handler.go            # Consumes: proposal.submitted            (from proposals-be)
│   │   │   │                                     # Purpose: deduct connects / enforce submit limits based on plan/entitlements.
│   │   │   │
│   │   │   ├── job_handler.go                 # Consumes: job.posted                    (from jobs-be)
│   │   │   │                                     # Purpose: check posting limits, apply plan gating, consume connects if required.
│   │   │   │
│   │   │   └── admin_flags_handler.go         # Consumes: admin.feature_flag.updated |
│   │   │                                         #           admin.threshold.updated |
│   │   │                                         #           admin.experiment.updated          (from admin-be)
│   │   │                                         # Purpose: refresh runtime flags/thresholds/experiments that gate
│   │   │                                         #          features like boosts, extra invites, posting quotas.
│   │   ├── plan/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Create, Update plans
│   │   │   ├── queries.go                    # Get, List plans
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go                 # 🆕 Validate features/limits/pricing
│   │   ├── subscription/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Subscribe, Cancel, Change, Pause
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── lifecycle_manager.go          # Manage subscription lifecycle
│   │   │   ├── renewal_manager.go            # Handle renewals
│   │   │   └── validators.go                 # 🆕 Validate status transitions, proration inputs
│   │   ├── connect/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Purchase, Use, Transfer, Refund
│   │   │   ├── queries.go                    # Get balance, History
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── calculator.go                 # Calculate connect costs
│   │   │   └── validators.go                 # 🆕 Validate expiries/rollovers/amounts
│   │   ├── usage/
│   │   │   ├── service.go
│   │   │   ├── tracker.go                    # Track usage
│   │   │   ├── quota_checker.go              # Check quotas
│   │   │   ├── limiter.go                    # Enforce limits
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go                 # 🆕 Validate counters & limits
│   │   ├── entitlement/                      # 🆕 Plans & entitlements
│   │   │   ├── service.go                    # Resolve entitlements, check feature access
│   │   │   ├── commands.go                   # GrantEntitlement, RevokeEntitlement (admin/system)
│   │   │   ├── queries.go                    # GetEntitlements, CheckFeature
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go                 # Validate feature scopes and subjects
│   │   ├── seat_billing/                     # 🆕 Seats + overages + proration + export
│   │   │   ├── service.go                    # Allocate seats, compute overages, proration, export invoice
│   │   │   ├── commands.go                   # AssignSeat, ReleaseSeat, SetSeatCap, RecalculateOverages
│   │   │   ├── queries.go                    # GetSeatSummary, GetOverages
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go                 # Validate seat counts and policies
│   │   ├── addon/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                   # 🆕 AddAddon, RemoveAddon
│   │   │   ├── queries.go                    # 🆕 GetAddons, GetAddon
│   │   │   └── validators.go                 # 🆕 Validate addon compatibility
│   │   ├── promotion/
│   │   │   ├── service.go
│   │   │   ├── validator.go                  # Validate promo codes (kept as-is)
│   │   │   ├── validators.go                 # 🆕 Aliases to keep naming consistency
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── trial/
│   │   │   ├── service.go
│   │   │   ├── eligibility_checker.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                   # 🆕 StartTrial, EndTrial
│   │   │   ├── queries.go                    # 🆕 GetTrial, IsEligible
│   │   │   └── validators.go                 # 🆕 Validate trial eligibility/rules
│   │   ├── billing/
│   │   │   ├── service.go
│   │   │   ├── invoice_generator.go          # Generate invoices
│   │   │   ├── payment_processor.go          # Process payments
│   │   │   ├── exporter.go                   # 🆕 Export invoices to Financial-BE
│   │   │   ├── dto.go
│   │   │   ├── commands.go                   # 🆕 GenerateInvoice, ExportInvoice, CapturePayment
│   │   │   ├── queries.go                    # 🆕 GetInvoice, ListInvoices
│   │   │   └── validators.go                 # 🆕 Validate seat counts, proration, exports
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go
│   │   │       ├── transaction.go
│   │   │       ├── migrations.go
│   │   │       ├── version.go
│   │   │       ├── safety.go
│   │   │       ├── plan_repository.go
│   │   │       ├── subscription_repository.go
│   │   │       ├── connect_repository.go
│   │   │       ├── usage_repository.go
│   │   │       ├── entitlement_repository.go      # 🆕 Implementation
│   │   │       ├── seat_billing_repository.go     # 🆕 Implementation
│   │   │       ├── addon_repository.go
│   │   │       ├── promotion_repository.go
│   │   │       ├── trial_repository.go
│   │   │       ├── billing_history_repository.go
│   │   │       ├── feature_toggle_repository.go
│   │   │       └── outbox_repository.go           # ❌ REMOVED (use platform-shared/outbox/postgres/repository.go)
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go
│   │   │       ├── subscription_cache.go
│   │   │       ├── plan_cache.go
│   │   │       ├── connect_cache.go
│   │   │       ├── entitlement_cache.go           # 🆕 Cache entitlements
│   │   │       └── feature_toggle_cache.go
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go
│   │   │       ├── producer.go
│   │   │       ├── topics.go
│   │   │       └── scram.go
│   │   ├── scheduler/
│   │   │   ├── cron.go
│   │   │   └── jobs.go
│   │   ├── payment/
│   │   │   └── client.go                         # Financial-BE client (unchanged)
│   │   ├── billing_export/                        # 🆕 Explicit Financial-BE export client
│   │   │   ├── client.go                          # REST/gRPC client for invoice export
│   │   │   └── mapper.go                          # Map local invoices/overages → Financial-BE DTO
│   │   └── outbox/
│   │       ├── processor.go                       # ❌ REMOVED
│   │       └── scheduler.go                       # ❌ REMOVED
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── plan_handler.go
│   │       │   ├── subscription_handler.go
│   │       │   ├── connect_handler.go
│   │       │   ├── usage_handler.go
│   │       │   ├── entitlement_handler.go        # 🆕 Entitlement endpoints
│   │       │   ├── seat_billing_handler.go       # 🆕 Seat mgmt & overages
│   │       │   ├── addon_handler.go
│   │       │   ├── promotion_handler.go
│   │       │   ├── trial_handler.go
│   │       │   ├── billing_handler.go
│   │       │   └── health_handler.go
│   │       ├── middleware/
│   │       │   ├── auth.go
│   │       │   ├── rbac.go
│   │       │   ├── cors.go
│   │       │   ├── rate_limit.go
│   │       │   ├── logging.go
│   │       │   ├── error_handler.go
│   │       │   ├── request_id.go
│   │       │   └── feature_gate.go
│   │       ├── responses/
│   │       │   ├── success.go
│   │       │   ├── error.go
│   │       │   └── pagination.go
│   │       ├── routes/                           # Route registrars mapped 1:1 to handlers
│   │       │   ├── plan_routes.go                # /plans/*
│   │       │   ├── subscription_routes.go        # /subscriptions/*
│   │       │   ├── connect_routes.go             # /connects/*
│   │       │   ├── usage_routes.go               # /usage/*
│   │       │   ├── entitlement_routes.go         # 🆕 /entitlements/*
│   │       │   ├── seat_billing_routes.go        # 🆕 /seats/*
│   │       │   ├── addon_routes.go               # /addons/*
│   │       │   ├── promotion_routes.go           # /promotions/*
│   │       │   ├── trial_routes.go               # /trials/*
│   │       │   └── billing_routes.go             # /billing/*
│   │       └── router.go
│   │
│   └── config/
│       ├── schema.go
│       ├── loader.go
│       └── docs/
│           └── CONFIGURATION.md
│
├── config/
│   ├── default.yaml
│   ├── dev.yaml
│   └── prod.yaml
│
├── dapr/
│   ├── local/
│   │   ├── pubsub.yaml
│   │   └── statestore.yaml
│   └── k8s/
│       ├── pubsub.yaml
│       ├── statestore.yaml
│       └── secrets.yaml
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   └── codes.go
│   ├── logger/
│   │   └── logger.go
│   ├── utils/
│   │   ├── validator.go
│   │   ├── billing_calculator.go
│   │   └── proration.go
│   └── constants/
│       ├── events.go
│       ├── topics.go
│       ├── plans.go
│       └── features.go
│
├── seeds/
│   ├── plans.sql
│   └── connect_packages.sql
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       ├── cronjob-renewal.yaml
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   ├── seed-plans.sh
│   ├── seed-data.sh
│   └── process-renewals.sh
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── api.md
│   ├── events.md
│   ├── subscription-plans.md
│   ├── connects-system.md
│   ├── billing-logic.md
│   ├── feature-toggles.md
│   ├── MIGRATIONS.md
│   ├── SCHEMA.md
│   └── RUNBOOK.md
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md

```


---

## **📦1️⃣1️⃣ admin-be (NEW - COMPREHENSIVE**

```
apps/be/admin-be/
│
├── cmd/
│   └── api/
│       └── main.go                           # 📝 UPDATED: Application entry point - initializes Gin, Dapr, connects to Postgres (now uses loadConfig from internal/config, platform-shared/logging)
│
├── internal/
│   ├── domain/
│   │   ├── admin_user/
│   │   │   ├── entity.go                     # Admin user accounts
│   │   │   ├── role.go                       # Admin roles (SuperAdmin, Moderator, Support)
│   │   │   ├── permission.go                 # Granular permissions
│   │   │   ├── activity_log.go               # Admin activity tracking
│   │   │   ├── errors.go                     # 🆕 Domain errors (AdminUserNotFound, PermissionDenied)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.user.created / admin.user.updated / admin.user.role.changed / admin.user.permissions.updated / admin.user.login.logged
│   │   ├── support_ticket/
│   │   │   ├── entity.go                     # Support tickets
│   │   │   ├── priority.go                   # Priority levels (Low, Medium, High, Urgent)
│   │   │   ├── status.go                     # Status (Open, InProgress, Resolved, Closed)
│   │   │   ├── category.go                   # Ticket categories
│   │   │   ├── assignment.go                 # Agent assignment
│   │   │   ├── errors.go                     # 🆕 Ticket errors (TicketNotFound, InvalidStatus)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.case.opened / admin.case.updated / admin.case.assigned / admin.case.closed
│   │   ├── ticket_message/
│   │   │   ├── entity.go                     # Ticket conversation messages
│   │   │   ├── attachment.go                 # Message attachments
│   │   │   ├── errors.go                     # 🆕 Message errors (MessageNotFound, AttachmentTooLarge)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: ticket.note.added / ticket.note.edited / ticket.note.deleted / ticket.attachment.added
│   │   ├── support_agent/
│   │   │   ├── entity.go                     # Support agent profiles
│   │   │   ├── availability.go               # Agent availability/online status
│   │   │   ├── stats.go                      # Agent performance stats
│   │   │   ├── errors.go                     # 🆕 Agent errors (AgentNotFound, AgentUnavailable)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: support.agent.status.changed / support.agent.assigned / support.agent.workload.updated
│   │   ├── canned_response/
│   │   │   ├── entity.go                     # Predefined responses
│   │   │   ├── category.go                   # Response categories
│   │   │   ├── errors.go                     # 🆕 Canned response errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: canned_response.created / canned_response.updated / canned_response.archived
│   │   ├── knowledge_base/
│   │   │   ├── entity.go                     # KB articles
│   │   │   ├── category.go                   # Article categories
│   │   │   ├── tag.go                        # Article tags
│   │   │   ├── version.go                    # Article versioning
│   │   │   ├── errors.go                     # 🆕 KB errors (ArticleNotFound, VersionConflict)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: kb.article.created / kb.article.updated / kb.article.versioned / kb.article.published / kb.article.unpublished
│   │   ├── faq/
│   │   │   ├── entity.go                     # Frequently asked questions
│   │   │   ├── category.go                   # FAQ categories
│   │   │   ├── errors.go                     # 🆕 FAQ errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: faq.created / faq.updated / faq.published / faq.unpublished
│   │   ├── moderation_queue/
│   │   │   ├── entity.go                     # Moderation queue items
│   │   │   ├── content_type.go               # Job, User, Review, Message, etc.
│   │   │   ├── flag_reason.go                # Reasons for flagging
│   │   │   ├── action.go                     # Moderation actions taken
│   │   │   ├── errors.go                     # 🆕 Moderation errors (QueueItemNotFound)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.moderation.enqueued / admin.moderation.assigned / admin.moderation.state.changed
│   │   ├── user_action/
│   │   │   ├── entity.go                     # Admin actions on users
│   │   │   ├── action_type.go                # Suspend, Ban, Verify, Warn, etc.
│   │   │   ├── reason.go                     # Action reasons
│   │   │   ├── errors.go                     # 🆕 User action errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.user_action.applied (suspend/ban/verify/warn) / admin.user_action.reversed
│   │   ├── content_action/
│   │   │   ├── entity.go                     # Admin actions on content
│   │   │   ├── action_type.go                # Remove, Hide, Approve, Reject
│   │   │   ├── errors.go                     # 🆕 Content action errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.content.actioned (approve/reject/remove/hide)
│   │   ├── dispute_resolution/
│   │   │   ├── entity.go                     # Dispute cases
│   │   │   ├── evidence.go                   # Evidence submitted
│   │   │   ├── decision.go                   # Admin decision
│   │   │   ├── errors.go                     # 🆕 Dispute errors (DisputeNotFound)
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.case.opened (dispute) / admin.case.updated / admin.case.closed / admin.dispute.decision.made
│   │   ├── system_config/
│   │   │   ├── entity.go                     # System configuration
│   │   │   ├── feature_flag.go               # Feature flags
│   │   │   ├── maintenance.go                # Maintenance mode settings
│   │   │   ├── errors.go                     # 🆕 Config errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.feature_flag.updated / admin.threshold.updated / admin.experiment.updated / admin.maintenance.window.set
│   │   ├── announcement/
│   │   │   ├── entity.go                     # Platform announcements
│   │   │   ├── target.go                     # Target audience (All, Freelancers, Clients)
│   │   │   ├── schedule.go                   # Scheduled announcements
│   │   │   ├── errors.go                     # 🆕 Announcement errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.announcement.created / admin.announcement.scheduled / admin.announcement.published / admin.announcement.canceled
│   │   ├── report/
│   │   │   ├── entity.go                     # Generated reports
│   │   │   ├── report_type.go                # Report types (Users, Revenue, Activity)
│   │   │   ├── schedule.go                   # Scheduled reports
│   │   │   ├── errors.go                     # 🆕 Report errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.report.requested / admin.report.generated / admin.report.delivery.failed
│   │   ├── analytics_dashboard/
│   │   │   ├── entity.go                     # Dashboard configurations
│   │   │   ├── widget.go                     # Dashboard widgets
│   │   │   ├── metric.go                     # Custom metrics
│   │   │   ├── errors.go                     # 🆕 Analytics dashboard errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.dashboard.created / admin.dashboard.updated / admin.dashboard.widget.added / admin.dashboard.widget.removed
│   │   ├── audit_log/
│   │   │   ├── entity.go                     # Complete audit trail
│   │   │   ├── action.go                     # Actions performed
│   │   │   ├── resource.go                   # Resources affected
│   │   │   ├── errors.go                     # 🆕 Audit log errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.audit.logged (staff action appended)
│   │   ├── notification_blast/
│   │   │   ├── entity.go                     # Bulk notifications/emails
│   │   │   ├── audience.go                   # Target audience
│   │   │   ├── schedule.go                   # Scheduled blasts
│   │   │   ├── errors.go                     # 🆕 Blast errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.blast.scheduled / admin.blast.sent / admin.blast.failed / admin.blast.canceled
│   │   ├── platform_stats/
│   │   │   ├── entity.go                     # Platform statistics
│   │   │   ├── realtime.go                   # Real-time stats
│   │   │   ├── errors.go                     # 🆕 Platform stats errors
│   │   │   ├── repository.go
│   │   │   └── events.go                     # 🆕 Domain events: admin.platform.stats.snapshot.created / admin.platform.stats.realtime.updated
│   │   └── outbox/
│   │       ├── entity.go                     # ❌ REMOVED (use platform-shared/outbox/entity.go)
│   │       └── repository.go                 # ❌ REMOVED (use platform-shared/outbox/repository.go)
│   │
│   │   # --------------------------- NEW DOMAINS (ADD-ONLY) ---------------------------
│   │   ├── fraud_review/                     # 🆕 Case mgmt: fraud reviews
│   │   │   ├── entity.go
│   │   │   ├── status.go                     # Pending, Investigating, Cleared, Confirmed
│   │   │   ├── severity.go                   # Low, Medium, High, Critical
│   │   │   ├── reason.go                     # SuspiciousActivity, StolenCard, etc.
│   │   │   ├── notes.go                      # Investigator notes
│   │   │   ├── sla.go                        # SLA definitions for review steps
│   │   │   ├── errors.go                     # 🆕 FraudReviewNotFound, InvalidTransition
│   │   │   └── repository.go
│   │   ├── user_report/                      # 🆕 Case mgmt: user reports
│   │   │   ├── entity.go                     # Reports filed by users about users/content
│   │   │   ├── category.go                   # Spam, Harassment, IPViolation, etc.
│   │   │   ├── status.go                     # Open, Triaged, Actioned, Dismissed
│   │   │   ├── notes.go                      # Moderator notes
│   │   │   ├── sla.go                        # SLA targets for first-response/resolve
│   │   │   ├── errors.go                     # 🆕 UserReportNotFound, InvalidCategory
│   │   │   └── repository.go
│   │   ├── risk_management/                  # 🆕 Risk dashboards
│   │   │   ├── hold.go                       # Payment/account holds
│   │   │   ├── reserve.go                    # Rolling/reserve configs
│   │   │   ├── chargeback.go                 # Chargeback cases & states
│   │   │   ├── velocity_alert.go             # Velocity rule hits
│   │   │   ├── country_rate_anomaly.go       # Country/rate anomaly entity
│   │   │   ├── errors.go                     # 🆕 RiskItemNotFound, RuleConflict
│   │   │   └── repository.go
│   │   ├── ticket_note/                      # 🆕 Case mgmt: internal ticket notes
│   │   │   ├── entity.go
│   │   │   ├── errors.go                     # 🆕 TicketNoteNotFound
│   │   │   └── repository.go
│   │   └── data_export/                      # 🆕 Audit & export (GDPR/CCPA)
│   │       ├── entity.go                     # Export job (type, requester, scope)
│   │       ├── request.go                    # Data subject request metadata
│   │       ├── status.go                     # Requested, Approved, Generating, Delivered
│   │       ├── format.go                     # ZIP, CSV, JSONL, Parquet
│   │       ├── errors.go                     # 🆕 ExportNotFound, AccessDenied
│   │       └── repository.go
│   │   # ---------------------------------------------------------------------------
│   │
│   ├── application/
│   │   ├── eventhandler/
│   │   │   ├── storage_handler.go             # Consumes (from storage-be):
│   │   │   │                                     #   file.policy.updated             — update moderation/DLP config views;
│   │   │   │                                     #   file.policy.violation.detected    — open/route moderation case; risk signal;
│   │   │   │                                     #   file.lifecycle.soft_deleted       — reflect visibility change in dashboards;
│   │   │   │                                     #   file.lifecycle.restored           — restore visibility state;
│   │   │   │                                     #   file.lifecycle.legal_hold.placed  — show/propagate legal hold in case mgmt;
│   │   │   │                                     #   file.lifecycle.legal_hold.removed — clear legal hold indicators;
│   │   │   │                                     #   file.link.signed_url.created      — append to audit trail;
│   │   │   │                                     #   file.link.signed_url.revoked      — append to audit trail / revoke confirmations;
│   │   │   │                                     #   file.download.logged              — append to audit (SIEM/data-lake feed).
│   │   │   │
│   │   │   ├── search_handler.go              # Consumes (from search-be):
│   │   │   │                                     #   search.taxonomy.updated           — taxonomy quality oversight dashboards;
│   │   │   │                                     #   search.ltr.signal.recorded        — LTR quality monitoring;
│   │   │   │                                     #   search.personalization.updated    — privacy/compliance review surfaces;
│   │   │   │                                     #   search.index.hygiene.deduplicated — hygiene queue visibility;
│   │   │   │                                     #   search.index.hygiene.archived     — archive tracking;
│   │   │   │                                     #   search.index.hygiene.visibility.changed — moderation/visibility change intake;
│   │   │   │                                     #   search.document.indexed           — index ops audit;
│   │   │   │                                     #   search.document.reindexed         — reindex ops audit.
│   │   │   │
│   │   │   ├── reviews_handler.go             # Consumes (from reviews-be):
│   │   │   │                                     #   review.double_blind.submitted     — case intake; SLA timers;
│   │   │   │                                     #   review.double_blind.published     — publish audit / analytics;
│   │   │   │                                     #   review.double_blind.timeout       — timeout audit, follow-ups;
│   │   │   │                                     #   review.response.posted            — right-of-reply audit/visibility;
│   │   │   │                                     #   review.feedback.private_submitted — research/analytics sink;
│   │   │   │                                     #   reputation.updated                — eligibility/quality dashboards;
│   │   │   │                                     #   review.abuse.auto_flagged         — auto-moderation queue intake;
│   │   │   │                                     #   review.moderation.state.changed   — reflect action history & state.
│   │   │   │
│   │   │   ├── subscriptions_handler.go       # Consumes (from subscriptions-be):
│   │   │   │                                     #   subscription.created              — customer lifecycle dashboards;
│   │   │   │                                     #   subscription.changed              — same;
│   │   │   │                                     #   subscription.paused               — alerts & eligibility views;
│   │   │   │                                     #   subscription.canceled             — churn dashboards;
│   │   │   │                                     #   subscription.renewed              — renewal audit;
│   │   │   │                                     #   entitlement.feature.enabled       — propagate plan feature flags to runtime registry;
│   │   │   │                                     #   entitlement.feature.disabled      — revoke flags in registry;
│   │   │   │                                     #   connects.purchased                — credits oversight;
│   │   │   │                                     #   connects.used                     — consumption monitoring;
│   │   │   │                                     #   connects.expired                  — expiry monitoring;
│   │   │   │                                     #   connects.granted                  — promo/grant audit;
│   │   │   │                                     #   billing.invoice.generated         — billing pipeline visibility;
│   │   │   │                                     #   billing.invoice.exported          — export audit to financial-be;
│   │   │   │                                     #   billing.proration.applied         — proration audit;
│   │   │   │                                     #   usage.meter.incremented           — usage alerts dashboards;
│   │   │   │                                     #   usage.limit.reached               — limit alerts dashboards.
│   │   │   │
│   │   │   ├── contracts_handler.go           # Consumes (from contracts-be):
│   │   │   │                                     #   contract.financial_hold.placed    — open/annotate risk case;
│   │   │   │                                     #   contract.financial_hold.released  — close/annotate case;
│   │   │   │                                     #   contract.state.changed            — visibility/policy reactions (read-only intake).
│   │   │   │
│   │   │   ├── financial_handler.go           # Consumes (from financial-be & payments rails):
│   │   │   │                                     #   payment.processed | payment.failed  — dunning/collections dashboards;
│   │   │   │                                     #   payment.chargeback.created        — chargeback case intake;
│   │   │   │                                     #   payment.chargeback.updated        — case status updates;
│   │   │   │                                     #   risk.hold.placed | risk.hold.released — risk dashboard streams;
│   │   │   │                                     #   risk.reserve.set | risk.reserve.updated — reserve monitoring;
│   │   │   │                                     #   risk.chargeback.recorded | updated  — chargeback metrics;
│   │   │   │                                     #   risk.velocity.alert               — anomaly alerts.
│   │   │   │
│   │   │   ├── users_handler.go               # Consumes (from users-be / identity flows):
│   │   │   │                                     #   user.updated                      — reflect flags/verifications in case views;
│   │   │   │                                     #   contract.dispute.opened (user-side) — cross-link disputes to user case.
│   │   │   │
│   │   │   ├── reports_exports_handler.go     # Consumes (cross-service signals to drive admin exports/notifications):
│   │   │   │                                     #   admin.data_export.requested       — (when originated by other surfaces) intake/export flow;
│   │   │   │                                     #   admin.data_export.approved        — approval audit;
│   │   │   │                                     #   admin.data_export.generated       — delivery orchestration audit;
│   │   │   │                                     #   admin.data_export.delivered       — completion audit;
│   │   │   │                                     #   admin.data_export.revoked         — revoke audit.
│   │   │   │
│   │   │   └── communications_handler.go      # Consumes (from communications-be):
│   │   │                                         #   comm.delivery.logged   
│   │   ├── admin_user/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Create, Update, Deactivate
│   │   │   ├── queries.go                    # Get, List admins
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── permission_manager.go         # Manage permissions
│   │   ├── support_ticket/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Create, Assign, Resolve, Close
│   │   │   ├── queries.go                    # Get, List, Search tickets
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   ├── assignment_engine.go          # Auto-assign tickets
│   │   │   ├── escalation_manager.go         # Escalate urgent tickets
│   │   │   └── sla_tracker.go                # Track SLA compliance
│   │   ├── ticket_message/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go                 # 🆕 Input validation for messages/attachments
│   │   ├── support_agent/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── stats_calculator.go           # Calculate agent stats
│   │   ├── canned_response/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── knowledge_base/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Create, Update, Publish
│   │   │   ├── queries.go                    # Search, Get articles
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── search_service.go             # KB search
│   │   ├── faq/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── moderation/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Approve, Reject, Remove
│   │   │   ├── queries.go                    # Get queue, Filter
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   ├── queue_manager.go              # Manage moderation queue
│   │   │   ├── auto_moderator.go             # Automatic moderation rules
│   │   │   └── content_scanner.go            # Scan for violations
│   │   ├── user_management/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Suspend, Ban, Verify, Warn
│   │   │   ├── queries.go                    # Search users, Get details
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── action_validator.go           # Validate admin actions
│   │   ├── content_management/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Remove, Hide, Feature
│   │   │   ├── queries.go                    # List content, Search
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── dispute_resolution/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Review, Decide, Close
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── decision_engine.go            # Help make decisions
│   │   ├── system_config/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Update configs
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── announcement/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Create, Schedule, Send
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── report/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   ├── generator.go                  # Generate reports
│   │   │   ├── scheduler.go                  # Schedule reports
│   │   │   └── exporters/
│   │   │       ├── pdf_exporter.go           # PDF exporter
│   │   │       ├── csv_exporter.go           # CSV exporter
│   │   │       └── excel_exporter.go         # Excel exporter
│   │   ├── analytics/
│   │   │   ├── service.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── aggregator.go                 # Aggregate analytics
│   │   │   ├── metrics_calculator.go         # Calculate KPIs
│   │   │   ├── dashboard_builder.go          # Build dashboards
│   │   │   └── validators.go                 # 🆕 Validate analytics query params
│   │   ├── notification_blast/
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Create, Schedule, Send
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── audience_selector.go          # Select target users
│   │   ├── audit/
│   │   │   ├── service.go
│   │   │   ├── queries.go                    # Query audit logs
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go                 # 🆕 Validate audit queries/filters
│   │
│   │   # --------------------------- NEW APPLICATION LAYERS (ADD-ONLY) ----------------
│   │   ├── fraud_review/                     # 🆕
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # OpenReview, Assign, AddNote, SetStatus, LinkEvidence
│   │   │   ├── queries.go                    # Get, List, Search by user/payment/risk-signal
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── user_report/                      # 🆕
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # Create, Triage, Reassign, AddNote, Resolve, Dismiss
│   │   │   ├── queries.go                    # Get, List, Filter by category/status/SLA
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── risk/                             # 🆕 dashboards & actions
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # PlaceHold, ReleaseHold, SetReserve, RecordChargeback
│   │   │   ├── queries.go                    # Dashboards: holds/reserves/chargebacks/alerts/anomalies
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── ticket_note/                      # 🆕
│   │   │   ├── service.go
│   │   │   ├── commands.go                   # AddNote, EditNote, DeleteNote
│   │   │   ├── queries.go                    # ListNotesForTicket
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   └── data_export/                      # 🆕
│   │       ├── service.go
│   │       ├── commands.go                   # RequestExport, Approve, Generate, Deliver, Revoke
│   │       ├── queries.go                    # GetExport, ListExports, GetRequest
│   │       ├── dto.go
│   │       ├── mapper.go
│   │       └── validators.go
│   │   # ---------------------------------------------------------------------------
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection setup (DSN from config, connection pooling)
│   │   │       ├── transaction.go            # Transaction helpers (Begin, Commit, Rollback, WithTransaction wrapper)
│   │   │       ├── migrations.go             # 📝 UPDATED: Auto-migrate logic (now with version tracking, GORM AutoMigrate for all tables)
│   │   │       ├── version.go                # 🆕 Schema version tracking (SchemaVersion table, RecordMigration function)
│   │   │       ├── safety.go                 # 🆕 Pre-migration safety checks (environment validation, disk space check)
│   │   │       ├── admin_user_repository.go
│   │   │       ├── support_ticket_repository.go
│   │   │       ├── ticket_message_repository.go
│   │   │       ├── support_agent_repository.go
│   │   │       ├── canned_response_repository.go
│   │   │       ├── knowledge_base_repository.go
│   │   │       ├── faq_repository.go
│   │   │       ├── moderation_queue_repository.go
│   │   │       ├── user_action_repository.go
│   │   │       ├── content_action_repository.go
│   │   │       ├── dispute_resolution_repository.go
│   │   │       ├── system_config_repository.go
│   │   │       ├── announcement_repository.go
│   │   │       ├── report_repository.go
│   │   │       ├── analytics_dashboard_repository.go
│   │   │       ├── audit_log_repository.go
│   │   │       ├── notification_blast_repository.go
│   │   │       ├── platform_stats_repository.go
│   │   │       └── outbox_repository.go      # ❌ REMOVED (use platform-shared/outbox/postgres/repository.go)
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go             # Redis connection setup (connection pooling, retry logic)
│   │   │       ├── admin_cache.go            # Cache admin data (Get, Set, Invalidate with TTL)
│   │   │       ├── ticket_cache.go           # Cache tickets
│   │   │       ├── stats_cache.go            # Cache stats
│   │   │       └── config_cache.go           # Cache configs
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go               # 📝 UPDATED: Kafka consumer (now uses platform-shared/inbox for message deduplication)
│   │   │       ├── producer.go               # 📝 UPDATED: Kafka producer (now uses platform-shared/outbox for reliable publishing)
│   │   │       ├── topics.go                 # 📝 UPDATED: Topic constants (now imported from contracts/events - admin.user.suspended, admin.dispute.resolved)
│   │   │       └── scram.go                  # SCRAM authentication for Kafka (SASL/SCRAM-SHA-256)
│   │   ├── external_services/
│   │   │   ├── users_client.go               # Users service client
│   │   │   ├── jobs_client.go                # Jobs service client
│   │   │   ├── proposals_client.go           # Proposals service client
│   │   │   ├── contracts_client.go           # Contracts service client
│   │   │   ├── financial_client.go           # Financial service client
│   │   │   ├── reviews_client.go             # Reviews service client
│   │   │   ├── communications_client.go      # Communications service client
│   │   │   ├── search_client.go              # Search service client
│   │   │   ├── storage_client.go             # Storage service client
│   │   │   └── subscriptions_client.go       # Subscriptions service client
│   │   ├── keycloak/
│   │   │   ├── admin_client.go               # Keycloak admin operations
│   │   │   └── user_manager.go               # Manage users via Keycloak
│   │   ├── reporting/
│   │   │   ├── pdf_generator.go              # PDF generator
│   │   │   ├── csv_generator.go              # CSV generator
│   │   │   └── excel_generator.go            # Excel generator
│   │   └── outbox/
│   │       ├── processor.go                  # ❌ REMOVED (use platform-shared/outbox/forwarder.go)
│   │       └── scheduler.go                  # ❌ REMOVED (use platform-shared/outbox/scheduler.go)
│   │
│   │   # --------------------------- NEW POSTGRES REPOSITORIES (ADD-ONLY) -------------
│   │   └── persistence/
│   │       └── postgres/
│   │           ├── fraud_review_repository.go              # 🆕
│   │           ├── user_report_repository.go               # 🆕
│   │           ├── risk_hold_repository.go                 # 🆕
│   │           ├── risk_reserve_repository.go              # 🆕
│   │           ├── risk_chargeback_repository.go           # 🆕
│   │           ├── risk_velocity_alert_repository.go       # 🆕
│   │           ├── risk_country_rate_anomaly_repository.go # 🆕
│   │           ├── ticket_note_repository.go               # 🆕
│   │           └── data_export_repository.go               # 🆕
│   │   # ---------------------------------------------------------------------------
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── admin_user_handler.go
│   │       │   ├── support_ticket_handler.go
│   │       │   ├── ticket_message_handler.go
│   │       │   ├── support_agent_handler.go
│   │       │   ├── canned_response_handler.go
│   │       │   ├── knowledge_base_handler.go
│   │       │   ├── faq_handler.go
│   │       │   ├── moderation_handler.go
│   │       │   ├── user_management_handler.go
│   │       │   ├── content_management_handler.go
│   │       │   ├── dispute_resolution_handler.go
│   │       │   ├── system_config_handler.go
│   │       │   ├── announcement_handler.go
│   │       │   ├── report_handler.go
│   │       │   ├── analytics_handler.go
│   │       │   ├── notification_blast_handler.go
│   │       │   ├── audit_log_handler.go
│   │       │   ├── dashboard_handler.go
│   │       │   └── health_handler.go          # Health check endpoints (/health, /ready, /live)
│   │       ├── middleware/
│   │       │   ├── auth.go                    # 📝 UPDATED: Authentication middleware (uses pkg/auth for JWT verification)
│   │       │   ├── admin_auth.go              # 📝 UPDATED: Admin-specific auth (uses pkg/auth)
│   │       │   ├── permission_check.go        # 📝 UPDATED: Check admin permissions (uses pkg/auth)
│   │       │   ├── audit_logger.go            # Auto-log all admin actions (platform-shared/logging)
│   │       │   ├── cors.go                    # 📝 UPDATED: CORS (platform-shared/ginx)
│   │       │   ├── rate_limit.go              # Rate limiting
│   │       │   ├── logging.go                 # 📝 UPDATED: Structured logs (platform-shared/ginx/logging)
│   │       │   ├── error_handler.go           # Error handling
│   │       │   └── request_id.go              # 📝 UPDATED: X-Request-ID (platform-shared/ginx/requestid)
│   │       ├── responses/
│   │       │   ├── success.go                 # 📝 UPDATED: use platform-shared/httpx/response.go
│   │       │   ├── error.go                   # 📝 UPDATED: use platform-shared/httpx/errors.go
│   │       │   └── pagination.go              # 📝 UPDATED: use platform-shared/httpx/pagination.go
│   │       ├── routes/                        # 🆕 Route registrars mapped 1:1 to handlers
│   │       │   ├── admin_user_routes.go           # /admin/users/*
│   │       │   ├── support_ticket_routes.go       # /admin/tickets/*
│   │       │   ├── ticket_message_routes.go       # /admin/tickets/:id/messages/*
│   │       │   ├── support_agent_routes.go        # /admin/agents/*
│   │       │   ├── canned_response_routes.go      # /admin/canned-responses/*
│   │       │   ├── knowledge_base_routes.go       # /admin/kb/*
│   │       │   ├── faq_routes.go                  # /admin/faqs/*
│   │       │   ├── moderation_routes.go           # /admin/moderation/*
│   │       │   ├── user_management_routes.go      # /admin/users/actions/*
│   │       │   ├── content_management_routes.go   # /admin/content/*
│   │       │   ├── dispute_resolution_routes.go   # /admin/disputes/*
│   │       │   ├── system_config_routes.go        # /admin/config/*
│   │       │   ├── announcement_routes.go         # /admin/announcements/*
│   │       │   ├── report_routes.go               # /admin/reports/*
│   │       │   ├── analytics_routes.go            # /admin/analytics/*
│   │       │   ├── notification_blast_routes.go   # /admin/blasts/*
│   │       │   ├── audit_log_routes.go            # /admin/audit/*
│   │       │   ├── dashboard_routes.go            # /admin/dashboard/*
│   │       │   # --------------------------- NEW ROUTES (ADD-ONLY) -------------------
│   │       │   ├── fraud_review_routes.go         # 🆕 /admin/fraud-reviews/*
│   │       │   ├── user_report_routes.go          # 🆕 /admin/user-reports/*
│   │       │   ├── risk_routes.go                 # 🆕 /admin/risk/*
│   │       │   ├── data_export_routes.go          # 🆕 /admin/exports/*
│   │       │   └── ticket_note_routes.go          # 🆕 /admin/tickets/:id/notes/*
│   │       └── router.go                          # 📝 UPDATED: HTTP router setup (uses Gin, applies platform-shared/ginx middleware)
│   │
│   └── config/                                   # 🆕 MEDIUM - Standardized configuration
│       ├── schema.go                             # 🆕 Typed Config struct (App, Server, Postgres, Kafka, Redis, Keycloak)
│       ├── loader.go                             # 🆕 Config loader using Viper (precedence: flags → env → file → defaults)
│       └── docs/
│           └── CONFIGURATION.md                  # 🆕 Configuration documentation (all ENV vars, defaults, examples)
│
├── config/                                       # 🆕 MEDIUM - Configuration files
│   ├── default.yaml                              # 🆕 Default configuration
│   ├── dev.yaml                                  # 🆕 Development overrides
│   └── prod.yaml                                 # 🆕 Production overrides
│
├── dapr/                                         # 🆕 MEDIUM - Dapr components split by environment
│   ├── local/                                    # 🆕 For dapr run
│   │   ├── pubsub.yaml                           # Kafka pub/sub component
│   │   └── statestore.yaml                       # State store component
│   └── k8s/                                      # 🆕 For kubectl apply
│       ├── pubsub.yaml                           # Kafka with scopes: ["admin-be"]
│       ├── statestore.yaml                       # State store with scopes
│       └── secrets.yaml                          # Dapr secret store
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                             # Service-specific errors
│   │   └── codes.go                              # Error codes
│   ├── logger/
│   │   └── logger.go                             # ❌ REMOVED: Use platform-shared/logging
│   ├── utils/
│   │   ├── validator.go                          # Local validation utilities
│   │   ├── sanitizer.go                          # Sanitize input
│   │   ├── permission_checker.go                 # Permission checking utilities
│   │   └── report_formatter.go                   # Report formatting
│   └── constants/
│       ├── events.go                             # ❌ REMOVED: Use contracts/events
│       ├── topics.go                             # ❌ REMOVED: Use contracts/events
│       ├── permissions.go                        # Permissions constants
│       └── moderation_actions.go                 # Moderation actions constants
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                       # Kubernetes Deployment
│       ├── service.yaml                          # Kubernetes Service
│       ├── configmap.yaml                        # ConfigMap
│       ├── secrets.yaml                          # Secrets
│       ├── hpa.yaml                              # HPA
│       ├── pdb.yaml                              # PDB
│       ├── rbac.yaml                             # Admin RBAC policies
│       └── servicemonitor.yaml                   # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                            # Setup local environment
│   ├── get-secrets.sh                            # Fetch secrets
│   ├── seed-admin-users.sh                       # Seed admin users
│   ├── seed-canned-responses.sh                  # Seed canned responses
│   └── seed-data.sh                              # Seed test data
│
├── tests/
│   ├── unit/
│   │   ├── domain/
│   │   │   ├── admin_user_test.go               # Admin user tests
│   │   │   ├── support_ticket_test.go           # Support ticket tests
│   │   │   └── moderation_queue_test.go         # Moderation queue tests
│   │   ├── application/
│   │   │   ├── admin_user_service_test.go       # Admin user service tests
│   │   │   ├── support_ticket_service_test.go   # Support ticket service tests
│   │   │   └── moderation_service_test.go       # Moderation service tests
│   │   └── infrastructure/
│   │       ├── postgres_repository_test.go      # Postgres repository tests
│   │       └── kafka_producer_test.go           # Kafka producer tests
│   ├── integration/
│   │   ├── handlers/
│   │   │   ├── admin_user_handler_test.go      # Admin user handler tests
│   │   │   ├── support_ticket_handler_test.go  # Support ticket handler tests
│   │   │   └── moderation_handler_test.go      # Moderation handler tests
│   │   └── repositories/
│   │       ├── admin_user_repository_test.go   # Admin user repository tests
│   │       └── support_ticket_repository_test.go # Support ticket repository tests
│   └── e2e/
│       └── scenarios/
│           ├── ticket_workflow_test.go         # Ticket workflow tests
│           ├── moderation_workflow_test.go     # Moderation workflow tests
│           └── dispute_resolution_test.go      # Dispute resolution tests
│
├── docs/
│   ├── README.md                               # Service overview
│   ├── api.md                                  # API documentation
│   ├── events.md                               # 📝 UPDATED: Events (published: admin.user.suspended, admin.dispute.resolved, consumed: user.flagged, contract.dispute.opened, etc.)
│   ├── admin-roles.md                          # Admin roles
│   ├── permissions.md                          # Permissions
│   ├── moderation-guide.md                     # Moderation guide
│   ├── support-workflows.md                    # Support workflows
│   ├── reporting.md                            # Reporting
│   ├── MIGRATIONS.md                           # 🆕 Migration history
│   ├── SCHEMA.md                               # 🆕 Database schema
│   └── RUNBOOK.md                              # 🆕 Operational procedures
│
├── .github/
│   └── workflows/
│       ├── ci.yml                              # CI workflow
│       └── cd.yml                              # CD workflow
│
├── go.mod                                      # 📝 UPDATED: Imports pkg/auth, platform-shared, contracts/events
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md

```


---

# Complete Event Flow & Integration Map

## Key Events Between Services

### 1. User Lifecycle Events
```
users-be → All Services
- UserCreated
- UserUpdated
- UserVerified
- FreelancerProfileCompleted
- ClientProfileCompleted
```

### 2. Job Lifecycle Events
```
jobs-be → proposals-be, search-be, communications-be, subscriptions-be
- JobPosted
- JobUpdated
- JobClosed
- JobInvitationSent
```

### 3. Proposal & Bidding Events
```
proposals-be → jobs-be, contracts-be, communications-be, subscriptions-be, financial-be
- ProposalSubmitted
- ProposalAccepted
- ProposalRejected
- BidPlaced
- BidUpdated
- OutbidAlert
- ConnectUsed
```

### 4. Contract Events
```
contracts-be → financial-be, communications-be, reviews-be, users-be
- ContractCreated
- ContractStarted
- MilestoneCreated
- MilestoneCompleted
- MilestoneApproved
- TimesheetSubmitted
- ContractPaused
- ContractEnded
```

### 5. Financial Events
```
financial-be → contracts-be, users-be, communications-be, subscriptions-be
- PaymentProcessed
- PaymentFailed
- EscrowHeld
- EscrowReleased
- PayoutRequested
- PayoutProcessed
- InvoiceGenerated
```

### 6. Communication Events
```
communications-be → All Services
- MessageSent
- NotificationDelivered
- EmailSent
- InAppNotificationSent
```

### 7. Review & Rating Events
```
reviews-be → users-be, search-be, communications-be
- ReviewSubmitted
- ReviewResponded
- BadgeAwarded
- ReputationUpdated
- RatingCalculated
```

### 8. Search & Recommendation Events
```
search-be → communications-be
- JobIndexed
- UserIndexed
- RecommendationGenerated
- MatchFound
```

### 9. Subscription Events
```
subscriptions-be → users-be, jobs-be, proposals-be, communications-be, financial-be
- SubscriptionCreated
- SubscriptionRenewed
- SubscriptionCancelled
- ConnectsPurchased
- ConnectsUsed
- UsageLimitReached
```

### 10. Storage Events
```
storage-be → users-be, jobs-be, contracts-be, proposals-be
- FileUploaded
- FileDeleted
- MediaProcessed
```

## 11. Events Published by admin-be

```
// User Actions
- admin.user.suspended
- admin.user.unsuspended
- admin.user.banned
- admin.user.unbanned
- admin.user.warned
- admin.user.verified

// Content Actions
- admin.job.removed
- admin.job.hidden
- admin.job.featured
- admin.proposal.removed
- admin.review.removed
- admin.message.removed
- admin.file.removed

// Moderation
- admin.moderation.approved
- admin.moderation.rejected
- admin.flag.resolved

// Disputes
- admin.dispute.resolved
- admin.dispute.escalated

// Financial
- admin.payment.refunded
- admin.chargeback.resolved

// Subscriptions
- admin.subscription.cancelled
- admin.subscription.extended
- admin.connects.added

// System
- admin.config.updated
- admin.announcement.published
- admin.maintenance.scheduled
```

---

## Events Subscribed by admin-be

```
// Flagging Events
- user.flagged
- job.flagged
- proposal.flagged
- contract.dispute.opened
- payment.disputed
- review.flagged
- message.flagged
- file.flagged

// System Events
- system.error
- system.alert
- system.abuse.detected

// Support Events
- support.ticket.created
- support.ticket.escalated
```



---

## In-App Notification Types (communications-be)

### Job-Related Notifications
1. **New Job Posted** (matching your skills)
2. **Job Invitation Received**
3. **Job You Applied To Closed**
4. **Similar Jobs Available**

### Proposal & Bidding Notifications
1. **Your Proposal Was Viewed**
2. **Your Proposal Was Accepted**
3. **Your Proposal Was Rejected**
4. **New Bid Placed on Your Job**
5. **You've Been Outbid**
6. **Bid Accepted**

### Contract Notifications
1. **New Contract Created**
2. **Contract Started**
3. **Milestone Completed**
4. **Milestone Approved**
5. **Timesheet Submitted**
6. **Work Diary Updated**
7. **Contract Ending Soon**

### Financial Notifications
1. **Payment Received**
2. **Payment Sent**
3. **Funds in Escrow**
4. **Escrow Released**
5. **Payout Processed**
6. **Invoice Generated**
7. **Low Balance Alert**

### Message Notifications
1. **New Message Received**
2. **Message Read**
3. **User Is Typing**

### Review Notifications
1. **New Review Received**
2. **Review Response Posted**
3. **Badge Earned**

### Subscription Notifications
1. **Subscription Expiring Soon**
2. **Subscription Renewed**
3. **Connects Running Low**
4. **New Connects Purchased**

### System Notifications
1. **Profile Verification Approved/Rejected**
2. **Account Security Alert**
3. **Terms of Service Updated**

---

## Technology Stack Summary

### Backend
- **Language**: Go
- **Framework**: Standard library + Gin/Echo/Chi
- **Architecture**: Clean Architecture (DDD)

### Data Layer
- **Primary DB**: PostgreSQL (per service)
- **Caching**: Redis
- **Search**: Elasticsearch/OpenSearch
- **Object Storage**: MinIO (self-hosted)

### Messaging & Events
- **Event Bus**: Apache Kafka
- **Patterns**: Outbox Pattern, Event Sourcing, CQRS

### Authentication & Authorization
- **Auth Service**: Keycloak
- **Token**: JWT

### Communication
- **Email**: WildDuck (self-hosted SMTP)
- **WebSocket**: Native Go WebSocket
- **Real-time**: Server-Sent Events (SSE)

### Infrastructure
- **Container**: Docker
- **Orchestration**: Kubernetes
- **Service Mesh**: Istio/Linkerd (optional)
- **API Gateway**: Kong/NGINX

### Observability
- **Logging**: ELK Stack or Grafana Loki
- **Monitoring**: Prometheus + Grafana
- **Tracing**: Jaeger

### CI/CD
- **Version Control**: Git
- **CI/CD**: GitHub Actions / GitLab CI
- **GitOps**: ArgoCD

---

## Resource Estimation (Minimum K8s Cluster)

```yaml
# Minimal Production Setup
Nodes: 3 nodes × 8GB RAM = 24GB total

Application Services (10 services):
- users-be: 2 replicas × 512MB = 1GB
- jobs-be: 2 replicas × 512MB = 1GB
- proposals-be: 2 replicas × 512MB = 1GB
- contracts-be: 2 replicas × 512MB = 1GB
- financial-be: 3 replicas × 1GB = 3GB (critical)
- communications-be: 2 replicas × 512MB = 1GB
- storage-be: 2 replicas × 512MB = 1GB
- search-be: 2 replicas × 1GB = 2GB
- reviews-be: 2 replicas × 512MB = 1GB
- subscriptions-be: 2 replicas × 512MB = 1GB

Subtotal: ~14GB

Infrastructure:
- PostgreSQL: 2GB
- Redis: 1GB
- Elasticsearch: 2GB
- Kafka: 2GB
- Keycloak: 1GB
- MinIO: 1GB

Subtotal: ~9GB

Total: ~23GB (fits in 24GB cluster)
```

---

This completes the comprehensive microservices architecture for your Skillsier platform! All services now include:

✅ **Bidding system** in proposals-be
✅ **In-app notifications** with comprehensive coverage in communications-be
✅ **Free & self-hosted** communications (WildDuck, WebSocket, SSE)
✅ **Auto-migrations** following your users-be pattern
✅ **Enhanced recommendation system** in search-be with ML models
✅ **Reviews covering ratings** in reviews-be
✅ Complete event flows between all services
✅ Enterprise-level folder structures

Each service is production-ready and follows Go best practices with Clean Architecture!