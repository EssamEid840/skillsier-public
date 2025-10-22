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
│   │   ├── schema.go                             # Typed Config (App, Server, Postgres, Kafka, Redis, Auth, WildDuck, WebSocket, WebPush, # 🆕 EnvelopeSign/Verify Keys, # 🆕 ES, # 🆕 KMS, # 🆕 SMS/SMPP, # 🆕 APNS/FCM)
│   │   ├── loader.go                             # Viper loader (flags → env → file → defaults)
│   │   └── docs/CONFIGURATION.md                 # ENV vars, defaults, examples
│   │
│   ├── domain/                                   # 🏛️ Domain Layer (business rules, aggregates, events)
│   │   # =========================
│   │   # 💬 CORE CHAT PRIMITIVES
│   │   # =========================
│   │   ├── conversation/
│   │   │   ├── entity.go                         # id, kind(direct,group,system), tenant_id, created_by, visibility, data_zone
│   │   │   ├── participant.go                    # user_id, role(owner,member), last_read_msg_id (kept), pinned, muted_until
│   │   │   ├── settings.go                       # ttl_policy_id, legal_hold, slow_mode, allow_replies, allow_files
│   │   │   ├── typing_indicator.go               # (kept) typing TTL markers per conversation/user
│   │   │   ├── errors.go                         # ConversationNotFound, ParticipantNotFound, UnauthorizedAccess
│   │   │   ├── repository.go                     # Create, FindByID, AddParticipant, RemoveParticipant
│   │   │   └── events.go                         # conversation.created/updated/archived/member_added/removed.v1
│   │   ├── thread/                               # 🆕 sub-discussions
│   │   │   ├── entity.go                         # id, conversation_id, root_message_id, title, followers[]
│   │   │   └── events.go                         # thread.created/renamed/archived.v1
│   │   ├── message/
│   │   │   ├── entity.go                         # id, conversation_id, sender_id, body(rich), reply_to_id, edited_at, deleted_at, redact_reason, # 🆕 seq
│   │   │   ├── attachment.go                     # storage-be asset refs (url,id,hash,type,size,thumb), virus_status
│   │   │   ├── read_receipt.go                   # message_id, user_id, read_at (rollup-friendly)
│   │   │   ├── reaction.go                       # emoji, user_id, reacted_at
│   │   │   ├── mention.go                        # mentioned_user_id, offsets
│   │   │   ├── errors.go                         # MessageNotFound, InvalidContent, MessageTooLong
│   │   │   ├── repository.go                     # Create, Update, Delete, FindByConversation, MarkAsRead
│   │   │   ├── events.go                         # message.sent/edited/deleted/reacted/mentioned.v1
│   │   │   └── sequence.go                       # 🆕 ReserveNextSequence(conversation_id) for monotonic ordering
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
│   │   ├── read_receipt/                         # explicit “I read it”
│   │   │   └── entity.go                         # message_id, user_id, read_at (compacted)
│   │   └── read_state/                           # 🆕 monotonic sequence pointers
│   │       ├── pointer.go                        # conversation_id, user_id, last_read_seq
│   │       └── events.go                         # message.read.v1 (payload: up_to_seq)
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
│   │   │   ├── quarantine.go                     # hide pending review; emits admin case
│   │   │   └── actions.go                        # 🆕 message.removed.v1, conversation.frozen/unfrozen.v1
│   │   ├── retention/
│   │   │   └── policy.go                         # per-room TTL, dispute_hold; purge windows
│   │   ├── blocklist/                            # 🆕 block phrases/users/domains
│   │   │   └── entity.go                         # scope(user/tenant/global), subject, reason, expires_at
│   │   ├── url_safety/                           # 🆕 reputation cache
│   │   │   ├── cache.go                          # open-source feeds only; refresh schedule
│   │   │   └── events.go                         # 🆕 url.scanned.v1, url_cache.updated/expired.v1
│   │   └── encryption/                           # 🆕 E2EE feature-scoped
│   │       ├── settings.go                       # enabled, key_ids, participants’ pub_keys
│   │       ├── events.go                         # e2e_encryption.enabled/disabled, encryption_key.rotated.v1
│   │       └── policy.go                         # gates: Search/AutoMod/Digests off when E2EE on
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
│   │   │   └── events.go                         # notification.created/updated/read/deleted/delayed.v1 🆕(delayed)
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
│   │   │   ├── events.go                         # email.queued/sent/delivered/bounced/failed.v1
│   │   │   └── tracking.go                       # 🆕 email.opened.v1 (best_effort), email.link_clicked.v1
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
│   │   │   └── events.go                         # webpush.subscription.added/removed/expired.v1
│   │   ├── sms/                                  # 🆕 SMS channel (opt-in/out + delivery)
│   │   │   ├── entity.go                         # opt-in records (e164_hash), sends, status timeline
│   │   │   ├── events.go                         # sms.opt_in/opt_out/sent/delivered/failed.v1
│   │   │   └── repository.go                     # opt-in/out, send, track status
│   │   ├── push/                                 # 🆕 Mobile push device registry (FCM/APNS; disabled by default)
│   │   │   ├── device.go                         # device_token, platform, user_id, attrs
│   │   │   ├── events.go                         # device.registered/unregistered/updated.v1
│   │   │   └── repository.go
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
│   │   │   ├── funnel.go                         # requested→sent→delivered→ack/read; histograms; SLOs
│   │   │   ├── message_stats.go                  # 🆕 user/conversation/platform message metrics
│   │   │   └── notification_stats.go             # 🆕 delivery & engagement by channel
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
│   │       ├── events.go                         # spam.detected/quarantined/reviewed.v1
│   │       └── rules_engine.go                   # Rule-based detection (keywords, links, repetition; no paid ML)
│   │
│   │   # =========================
│   │   # 🧩 PLATFORM EVENTS (ENVELOPE & REALTIME)
│   │   # =========================
│   │   ├── realtime/
│   │   │   └── events.go                         # 🆕 broadcast.sent.v1, broadcast.dropped.v1
│   │   └── webhook/                              # 🆕 outbound webhook subscriptions
│   │       ├── subscription.go                   # url, events[], secret
│   │       ├── events.go                         # webhook.subscribed/unsubscribed/delivered/failed/retried.v1
│   │       └── repository.go
│   │
│   ├── application/                              # 📋 Application Layer (use cases, orchestrators, consumers)
│   │   # =========================
│   │   # 📡 EVENT CONSUMERS (INBOX)
│   │   # =========================
│   │   ├── eventhandler/
│   │   │   ├── user_handler.go                   # user.created → welcome message/email
│   │   │   ├── job_handler.go                    # job.posted → notify matching freelancers
│   │   │   ├── proposal_handler.go               # proposal.submitted → notify client
│   │   │   ├── contract_handler.go               # contract.created → notify both parties
│   │   │   ├── payment_handler.go                # payment.processed → receipt notification
│   │   │   ├── review_handler.go                 # review.submitted / double_blind window nudges
│   │   │   ├── admin_case_handler.go             # admin.case.* → subject notifications as needed
│   │   │   ├── delivery_logger_handler.go        # emit comm.delivery.logged → admin audit
│   │   │   └── partitioning_notes.go             # (comments) partition keys per stream
│   │   # =========================
│   │   # 🧠 USE CASES (COMMANDS/QUERIES)
│   │   # =========================
│   │   # =========================
│   │   # 💬 CORE CHAT PRIMITIVES
│   │   # =========================
│   │   ├── conversation/
│   │   │   ├── service.go                        # Conversation business logic (Create, Archive, Mute, Delete)
│   │   │   ├── commands.go                       # CreateConversation, ArchiveConversation, MuteConversation, DeleteConversation
│   │   │   ├── queries.go                        # GetConversation, ListConversations, SearchConversations
│   │   │   ├── dto.go                            # ConversationDTO, CreateConversationDTO, ConversationListDTO
│   │   │   ├── mapper.go                         # Entity ↔ DTO mapping for conversations
│   │   │   └── validators.go                     # Input validation (membership, visibility, TTL policies)
│   │   ├── message/
│   │   │   ├── service.go                        # Message business logic (Send, Edit, Delete, React, MarkAsRead)
│   │   │   ├── commands.go                       # SendMessage, EditMessage, DeleteMessage, ReactToMessage, MarkAsRead
│   │   │   ├── queries.go                        # GetMessages, SearchMessages, GetUnreadCount
│   │   │   ├── dto.go                            # MessageDTO, SendMessageDTO, MessageListDTO
│   │   │   ├── mapper.go                         # Message entity ↔ DTO mappers
│   │   │   ├── validators.go                     # Content length, attachment limits, mention bounds
│   │   │   └── realtime_service.go               # 🆕 WebSocket handling (broadcast, typing, presence fan-out)
│   │   ├── thread/
│   │   │   ├── service.go                        # 🆕 Create/Archive thread, follow/unfollow
│   │   │   ├── commands.go                       # 🆕 CreateThread, ArchiveThread, FollowThread, UnfollowThread
│   │   │   ├── queries.go                        # 🆕 GetThread, ListThreadsForConversation
│   │   │   ├── dto.go                            # 🆕 ThreadDTO
│   │   │   ├── mapper.go                         # 🆕 Thread entity ↔ DTO mappers
│   │   │   └── validators.go                     # 🆕 Root message existence, membership checks
│   │   ├── pin/
│   │   │   ├── service.go                        # 🆕 PinMessage, UnpinMessage, ListPins
│   │   │   ├── commands.go                       # 🆕 PinMessage, UnpinMessage
│   │   │   ├── queries.go                        # 🆕 GetPinsForConversation
│   │   │   ├── dto.go                            # 🆕 PinDTO
│   │   │   ├── mapper.go                         # 🆕 Pin entity ↔ DTO mappers
│   │   │   └── validators.go                     # 🆕 Role/visibility checks
│   │   ├── bookmark/
│   │   │   ├── service.go                        # 🆕 AddBookmark, RemoveBookmark, ListBookmarks
│   │   │   ├── commands.go                       # 🆕 AddBookmark, RemoveBookmark
│   │   │   ├── queries.go                        # 🆕 GetBookmarksForUser
│   │   │   ├── dto.go                            # 🆕 BookmarkDTO
│   │   │   ├── mapper.go                         # 🆕 Bookmark entity ↔ DTO mappers
│   │   │   └── validators.go                     # 🆕 Ownership checks
│   │   ├── draft/
│   │   │   ├── service.go                        # 🆕 SaveDraft, ClearDraft, GetDraft
│   │   │   ├── commands.go                       # 🆕 SaveDraft, ClearDraft
│   │   │   ├── queries.go                        # 🆕 GetDraftForConversation
│   │   │   ├── dto.go                            # 🆕 DraftDTO
│   │   │   ├── mapper.go                         # 🆕 Draft entity ↔ DTO mappers
│   │   │   └── validators.go                     # 🆕 Size limits, content sanitization
│   │   # =========================
│   │   # 🚚 DELIVERY & READ STATE
│   │   # =========================
│   │   ├── read_receipt/
│   │   │   ├── service.go                        # 🆕 RecordRead, GetLatestRead, ListReaders
│   │   │   ├── commands.go                       # 🆕 RecordRead
│   │   │   ├── queries.go                        # 🆕 GetReadReceiptsForMessage
│   │   │   ├── dto.go                            # 🆕 ReadReceiptDTO
│   │   │   ├── mapper.go                         # 🆕 Read receipt entity ↔ DTO mappers
│   │   │   └── validators.go                     # 🆕 Membership & ordering checks
│   │   ├── delivery/
│   │   │   ├── service.go                        # 🆕 MarkDispatched, AckDelivery, GetDeliveryStatus
│   │   │   ├── commands.go                       # 🆕 MarkDispatched, AckDelivery
│   │   │   ├── queries.go                        # 🆕 GetDeliveriesForMessage
│   │   │   ├── dto.go                            # 🆕 DeliveryStatusDTO
│   │   │   ├── mapper.go                         # 🆕 Delivery entity ↔ DTO mappers
│   │   │   └── validators.go                     # 🆕 Session authenticity, idempotency keys
│   │   └── read_state/
│   │       ├── service.go                        # 🆕 MarkReadUpTo (advance pointer), GetUnreadCount
│   │       ├── commands.go                       # 🆕 MarkReadUpTo
│   │       └── queries.go                        # 🆕 GetUnreadCount
│   │   # =========================
│   │   # ⚡ EPHEMERAL REALTIME SIGNALS
│   │   # =========================
│   │   ├── online_status/
│   │   │   ├── service.go                        # Presence state machine (online, away, busy, offline)
│   │   │   ├── commands.go                       # SetOnline, SetAway, SetBusy, SetOffline
│   │   │   ├── queries.go                        # GetUserStatus, GetOnlineUsers
│   │   │   ├── validators.go                     # TTL bounds, allowed transitions
│   │   │   ├── tracker.go                        # Heartbeat ingestion & expiry logic
│   │   │   ├── presence_manager.go               # Session fan-out & dedupe across devices
│   │   │   └── dto.go                            # OnlineStatusDTO
│   │   └── typing/
│   │       ├── service.go                        # 🆕 StartTyping, StopTyping, GetTypingUsers
│   │       ├── commands.go                       # 🆕 StartTyping, StopTyping
│   │       ├── queries.go                        # 🆕 GetTypingForConversation
│   │       ├── dto.go                            # 🆕 TypingDTO
│   │       ├── mapper.go                         # 🆕 Typing signal ↔ DTO mappers
│   │       └── validators.go                     # 🆕 Rate limits, membership checks
│   │   # =========================
│   │   # 🛡️ SAFETY & COMPLIANCE
│   │   # =========================
│   │   ├── flag/
│   │   │   ├── service.go                        # Message flag lifecycle (Flag, Unflag, Resolve)
│   │   │   ├── commands.go                       # FlagMessage, UnflagMessage, ResolveFlag
│   │   │   ├── queries.go                        # GetFlags, GetFlag
│   │   │   ├── validators.go                     # Reason whitelist, reviewer role checks
│   │   │   ├── dto.go                            # FlagDTO, FlagMessageDTO
│   │   │   └── mapper.go                         # Flag entity ↔ DTO mappers
│   │   ├── moderation/
│   │   │   ├── service.go                        # 🆕 EvaluateRules, ApplyActions (quarantine/mask/notify/freeze/remove)
│   │   │   ├── commands.go                       # 🆕 UpsertRule, RemoveRule
│   │   │   ├── queries.go                        # 🆕 ListRules, GetRule
│   │   │   ├── dto.go                            # 🆕 ModerationRuleDTO
│   │   │   ├── mapper.go                         # 🆕 Rule entity ↔ DTO mappers
│   │   │   └── validators.go                     # 🆕 Pattern safety, action constraints
│   │   ├── retention_policy/
│   │   │   ├── service.go                        # 🆕 SetPolicy, GetPolicy, EnforcePolicy
│   │   │   ├── commands.go                       # 🆕 SetRetentionPolicy
│   │   │   ├── queries.go                        # 🆕 GetRetentionPolicy
│   │   │   ├── dto.go                            # 🆕 RetentionPolicyDTO
│   │   │   ├── mapper.go                         # 🆕 Policy entity ↔ DTO mappers
│   │   │   └── validators.go                     # 🆕 TTL bounds, hold precedence
│   │   └── blocklist/
│   │       ├── service.go                        # 🆕 AddBlock, RemoveBlock, IsBlocked
│   │       ├── commands.go                       # 🆕 AddBlock, RemoveBlock
│   │       ├── queries.go                        # 🆕 GetBlocksForScope
│   │       ├── dto.go                            # 🆕 BlockDTO
│   │       ├── mapper.go                         # 🆕 Block entity ↔ DTO mappers
│   │       └── validators.go                     # 🆕 Scope & expiry checks
│   │   # =========================
│   │   # 🔔 NOTIFS & USER FEEDS
│   │   # =========================
│   │   ├── notification/
│   │   │   ├── service.go                        # Notification business logic (Send, MarkAsRead, Delete, ClearAll)
│   │   │   ├── commands.go                       # SendNotification, MarkAsRead, DeleteNotification, ClearAllNotifications
│   │   │   ├── queries.go                        # GetNotifications, GetUnreadCount, GetNotificationHistory
│   │   │   ├── dto.go                            # NotificationDTO, SendNotificationDTO, NotificationListDTO
│   │   │   ├── mapper.go                         # Notification mappers
│   │   │   ├── validators.go                     # Type/category checks, payload schema validation
│   │   │   ├── orchestrator.go                   # 🆕 Multi-channel orchestration (in-app, email, webpush, sms, push)
│   │   │   ├── preferences_service.go            # 🆕 Manage per-topic/channel user preferences + DND windows
│   │   │   ├── aggregator.go                     # 🆕 Aggregate & group similar notifications (collapse keys/time window)
│   │   │   └── routing_policy.go                 # 🆕 urgency × quiet hours × channel → notification.delayed.v1
│   │   ├── notification_preferences/
│   │   │   ├── service.go                        # UpdatePreferences, GetPreferences, SetQuietHours, SetDigestSchedule
│   │   │   ├── commands.go                       # UpdatePreferences, SetQuietHours, SetDigestSchedule
│   │   │   ├── queries.go                        # GetPreferences, GetEffectivePreferences
│   │   │   ├── validators.go                     # Channel validation, timezone-safe windows
│   │   │   ├── dto.go                            # NotificationPreferencesDTO, QuietHoursDTO, DigestScheduleDTO
│   │   │   └── mapper.go                         # Preferences entity ↔ DTO mappers
│   │   ├── in_app_notification/
│   │   │   ├── service.go                        # In-app notification logic (render, deliver, state transitions)
│   │   │   ├── commands.go                       # PushInAppNotification, DismissInAppNotification, ClickInAppAction
│   │   │   ├── queries.go                        # GetInAppNotifications, GetBadgeCount
│   │   │   ├── validators.go                     # CTA validation, throttling & dedupe checks
│   │   │   ├── real_time_sender.go               # Real-time delivery via WS/SSE (user/room fan-out)
│   │   │   ├── badge_manager.go                  # Badge count calc/update/reset with cache hinting
│   │   │   ├── grouping_engine.go                # Grouping similar items (collapse keys)
│   │   │   ├── dto.go                            # InAppNotificationDTO
│   │   │   └── mapper.go                         # In-app notification mappers
│   │   ├── template/
│   │   │   ├── service.go                        # Create/Update/Render templates (versioned + i18n)
│   │   │   ├── commands.go                       # CreateTemplate, UpdateTemplate
│   │   │   ├── queries.go                        # GetTemplate, ListTemplates
│   │   │   ├── validators.go                     # Placeholder whitelist, locale fallback checks
│   │   │   ├── renderer.go                       # Safe renderer (HTML sanitizer, checksum)
│   │   │   ├── variable_injector.go              # Merge dynamic variables from context
│   │   │   ├── dto.go                            # TemplateDTO, RenderTemplateDTO
│   │   │   └── mapper.go                         # Template entity ↔ DTO mappers
│   │   ├── email/
│   │   │   ├── service.go                        # Email orchestration (Send, SendBatch, CheckStatus)
│   │   │   ├── commands.go                       # SendEmail, SendEmailBatch
│   │   │   ├── queries.go                        # GetEmailStatus, ListEmailsForUser
│   │   │   ├── validators.go                     # Address format, batch sizes, template variables
│   │   │   ├── sender.go                         # SMTP send workflow (retries, backoff, idempotency keys)
│   │   │   ├── template_renderer.go              # Render HTML/text templates with variable injection
│   │   │   ├── batch_sender.go                   # Batch queueing & rate limiting
│   │   │   ├── dto.go                            # EmailDTO, SendEmailDTO, BatchEmailDTO
│   │   │   ├── mapper.go                         # Email entity ↔ DTO mappers
│   │   │   └── wildduck_client.go                # 🆕 WildDuck SMTP/API integration (self-hosted MTA)
│   │   ├── sms/                                  # 🆕
│   │   │   ├── service.go                        # OptIn, OptOut, SendSMS, ProcessDLR
│   │   │   ├── commands.go                       # OptInSMS, OptOutSMS, SendSMS
│   │   │   ├── queries.go                        # GetSMSStatus
│   │   │   └── validators.go
│   │   └── push_device/                          # 🆕
│   │       ├── service.go                        # Register/Unregister devices (FCM/APNS; FF off)
│   │       ├── commands.go                       # RegisterDevice, UnregisterDevice
│   │       └── validators.go
│   │   # =========================
│   │   # ✉️ EMAIL BRIDGE (SELF-HOSTED)
│   │   # =========================
│   │   ├── email_bridge/
│   │   │   ├── inbound_processor.go              # 🆕 Parse inbound (plus-address) → conversation; trim signatures; auth checks
│   │   │   └── outbound_processor.go             # 🆕 Generate threading headers; queue to MTA; dedupe by message-id
│   │   # =========================
│   │   # 🔍 SEARCH & INDEXING
│   │   # =========================
│   │   └── search/                               # 🆕
│   │       ├── indexer.go                        # IndexMessage, ReindexConversation
│   │       ├── eraser.go                         # EraseUserData → push delete to ES
│   │       └── redactor.go                       # PII allowlist redaction before indexing
│   │   # =========================
│   │   # 🔐 ENCRYPTION (E2EE)
│   │   # =========================
│   │   ├── encryption/                           # 🆕
│   │   │   ├── service.go                        # EnableE2EE, DisableE2EE, RotateKey
│   │   │   └── validators.go
│   │   # =========================
│   │   # 🌐 WEBHOOKS (OUTBOUND)
│   │   # =========================
│   │   ├── webhook/                              # 🆕
│   │   │   ├── service.go                        # Subscribe, Unsubscribe, Deliver, Retry
│   │   │   └── validators.go
│   │   # =========================
│   │   # ⚖️ COMPLIANCE (EXPORT/ERASURE)
│   │   # =========================
│   │   └── compliance/                           # 🆕
│   │       ├── export_service.go                 # export.requested/completed
│   │       └── erasure_service.go                # data_deletion.requested/completed
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
│   │   │       ├── message_sequence_store.go        # 🆕 ReserveNextSequence store
│   │   │       # 🚚 DELIVERY & READ STATE
│   │   │       ├── read_receipt_repository.go       # 🆕 ReadReceiptRepository (record/read rollups)
│   │   │       ├── delivery_repository.go           # 🆕 DeliveryRepository (queued→dispatched→ack states)
│   │   │       ├── read_state_repository.go         # 🆕 Last-read pointer repository
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
│   │   │       ├── suppression_repository.go        # 🆕 SuppressionRepository (bounces/complaints; channel=email|sms)
│   │   │       ├── webpush_subscription_repository.go # 🆕 WebPushSubscriptionRepository (endpoint/keys/scope/expiry)
│   │   │       ├── sms_repository.go                # 🆕 SMS opt-in/out + delivery status
│   │   │       ├── push_device_repository.go        # 🆕 Device registry
│   │   │       # ✉️ EMAIL BRIDGE (SELF-HOSTED)
│   │   │       ├── email_repository.go              # EmailRepository (status/outbox/history)
│   │   │       ├── mail_tracking_repository.go      # 🆕 MailTrackingRepository (delivered/deferred/bounced/complained)
│   │   │       # 📅 SCHEDULING & CALLS
│   │   │       ├── system_message_repository.go     # SystemMessageRepository (contract/dispute feeds)
│   │   │       ├── call_repository.go               # CallRepository (links/schedule/state)
│   │   │       ├── calendar_invite_repository.go    # CalendarInviteRepository (iCal/status)
│   │   │       # 📊 OPS & GOVERNANCE
│   │   │       ├── quota_repository.go              # 🆕 QuotaRepository (per-user/topic/channel usage)
│   │   │       ├── analytics_repository.go          # 🆕 AnalyticsRepository (funnel counters, lag aggregates)
│   │   │       ├── webhook_subscription_repository.go # 🆕 Webhook subscriptions & logs
│   │   │       └── compliance_repository.go         # 🆕 Export/erasure requests
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
│   │   │       ├── scram.go                        # SASL/SCRAM-256
│   │   │       ├── middleware.go                   # 🆕 sign on produce; verify on consume (envelope)
│   │   │       └── zone_router.go                  # 🆕 publish eu.* / us.* by data_zone
│   │   # =========================
│   │   # 🔴 REALTIME (WS/SSE)
│   │   # =========================
│   │   ├── realtime/
│   │   │   ├── websocket/
│   │   │   │   ├── hub.go                          # connection manager; rooms; user fan-out
│   │   │   │   ├── client.go                       # per-conn read/write goroutines
│   │   │   │   ├── handler.go                      # HTTP→WS upgrade & auth
│   │   │   │   ├── broadcaster.go                  # send to all/user/room
│   │   │   │   ├── room.go                         # room registry (conversation_id)
│   │   │   │   └── backpressure.go                 # 🆕 queue budgets, drop policy, metrics
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
│   │   # 📱 PUSH / SMS / WEBPUSH
│   │   # =========================
│   │   ├── webpush/
│   │   │   └── vapid/
│   │   │       ├── signer.go                       # 🆕 VAPID JWT/EC keys
│   │   │       └── sender.go                       # 🆕 push send with retries
│   │   ├── sms/                                    # 🆕
│   │   │   ├── twilio_client.go                    # webhook signature verify (optional)
│   │   │   ├── smpp_client.go                      # self-hosted SMPP (Jasmin/Kannel)
│   │   │   └── router.go                           # provider selection (feature flags)
│   │   └── push/                                   # 🆕
│   │       ├── fcm_client.go                       # disabled by default
│   │       └── apns_client.go                      # disabled by default
│   │   # =========================
│   │   # 🔍 SEARCH / KMS / STORAGE
│   │   # =========================
│   │   ├── search/
│   │   │   └── elasticsearch/
│   │   │       ├── client.go
│   │   │       └── mapper.go
│   │   ├── kms/
│   │   │   └── client.go                           # 🆕 wrap/unwrap room keys
│   │   └── storage/
│   │       └── client.go                           # Storage service client (upload message attachments via HTTP API)
│   │   # =========================
│   │   # 🔐 PLATFORM (SECURITY & ENVELOPE)
│   │   # =========================
│   │   └── platform/
│   │       ├── events/
│   │       │   ├── envelope.go                     # tenant_id, data_zone, traceparent, schema_crc
│   │       │   ├── signer.go                       # Ed25519 sign (rotating keys)
│   │       │   └── verifier.go                     # verify + <=5m replay window
│   │       ├── policy/
│   │       │   └── residency.go                    # 🆕 consumer allowlist / residency enforcement
│   │       └── security/
│   │           ├── envelope_encryptor.go           # 🆕 per-tenant envelope encryption; Postgres RLS helpers
│   │           ├── webhook_signer.go               # 🆕 Ed25519 {t,v1} headers
│   │           ├── webhook_verifier.go             # 🆕 inbound verification
│   │           └── log_scrubber.go                 # 🆕 PII denylist/sampling
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
│   │           │   ├── sms_handler.go                  # 🆕 /sms opt-in/out, send, DLR webhook
│   │           │   ├── push_device_handler.go          # 🆕 register/unregister device tokens
│   │           │   ├── unsubscribe_handler.go          # unsubscribe flows
│   │           │   ├── email_handler.go                # send email / batch status
│   │           │   ├── email_tracking_handler.go       # 🆕 open/click endpoints (best-effort opens)
│   │           │   # ✉️ EMAIL BRIDGE (SELF-HOSTED)
│   │           │   ├── mail_tracking_handler.go        # 🆕 provider webhooks: delivered/deferred/bounced/complained
│   │           │   # 📅 SCHEDULING & CALLS
│   │           │   ├── system_message_handler.go       # system feed
│   │           │   ├── call_handler.go                 # call links & scheduling
│   │           │   ├── calendar_invite_handler.go      # invites
│   │           │   # 🌐 WEBHOOKS / COMPLIANCE
│   │           │   ├── webhook_handler.go              # 🆕 subscribe/unsubscribe; delivery logs
│   │           │   ├── compliance_handler.go           # 🆕 data export/erasure requests
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
│   │               ├── sms_routes.go                   # 🆕 /sms/*
│   │               ├── push_device_routes.go           # 🆕 /push/devices/*
│   │               ├── unsubscribe_routes.go           # /unsubscribe/*
│   │               ├── email_routes.go                 # /emails/*
│   │               ├── email_tracking_routes.go        # 🆕 /email/tracking/*
│   │               # ✉️ EMAIL BRIDGE (SELF-HOSTED)
│   │               ├── mail_tracking_routes.go         # 🆕 /mail/tracking/*
│   │               # 📅 SCHEDULING & CALLS
│   │               ├── system_message_routes.go        # /system-messages/*
│   │               ├── call_routes.go                  # /calls/*
│   │               ├── calendar_invite_routes.go       # /calendar-invites/*
│   │               # 🌐 WEBHOOKS / COMPLIANCE
│   │               ├── webhook_routes.go               # 🆕 /webhooks/*
│   │               └── compliance_routes.go            # 🆕 /compliance/*
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
│   ├── EVENTS.md                                  # 📝 published: message.sent, notification.delivered; consumed: users/jobs/proposals/contracts/payments/reviews/admin + 🆕 envelope, webhooks, sms, residency
│   ├── ARCHITECTURE.md                            # High-level diagrams
│   ├── MIGRATIONS.md                              # Migration history
│   ├── SCHEMA.md                                  # Database schema
│   ├── RUNBOOK.md                                 # Operational procedures (DLQ, rate-limit, replayer, 🆕 ES reindex, WS/SSE drain)
│   ├── websocket-protocol.md                      # WebSocket protocol documentation
│   ├── notification-system.md                     # Notification system overview (opens best_effort=true)
│   ├── in-app-notifications.md                    # In-app notifications guide
│   ├── wildduck-integration.md                    # WildDuck integration guide
│   ├── e2ee.md                                    # 🆕 End-to-End Encryption gating & KMS
│   ├── data-residency.md                          # 🆕 Zone topics, sanitized replication
│   └── security.md                                # 🆕 Envelope signatures, webhook signing, log scrubbing
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
│   ├── constants/
│   │   ├── notification_types.go                  # Notification type constants
│   │   └── websocket_events.go                    # WebSocket event types
│   └── metrics/                                   # 🆕 Observability helpers
│       ├── counters.go                            # idempotency_hits, ws_queue_depth, dlq_age, digest_backlog
│       └── histograms.go                          # send_latency, ack_latency
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
│   # =============================
│   # 🚀 APP ENTRYPOINTS
│   # =============================
│   ├── api/
│   │   └── main.go                               # 📝 API entrypoint - Gin+Dapr+Postgres (uses internal/config & platform-shared/logging)
│   └── worker/                                    # 🆕 background workers
│       └── main.go                                # 🆕 boot DI; run GC planner, quarantine sweeps, scan/outbox dispatchers; leader election
│
├── internal/
│   # =============================
│   # 🔧 CONFIGURATION (LOAD FIRST)
│   # =============================
│   ├── config/
│   │   ├── schema.go                              # Typed Config (App, Server, Postgres, Kafka, Redis, MinIO, Storage, DLP/AV, Residency)
│   │   ├── loader.go                              # Viper loader (flags → env → file → defaults)
│   │   └── docs/CONFIGURATION.md                  # ENV vars, defaults, examples
│   │
│   # =============================
│   # 🏛️ DOMAIN LAYER (DDD)
│   # =============================
│   ├── domain/
│   │   # =========================
│   │   # 📁 CORE FILE PRIMITIVES
│   │   # =========================
│   │   ├── file/
│   │   │   ├── entity.go                          # id, owner, namespace_id, name, ext, mime, size, hash, visibility, status, created_at
│   │   │   ├── enums.go                           # FileType, Status, Visibility
│   │   │   ├── metadata.go                        # Derived metadata (dimensions, duration, pages, exif)
│   │   │   ├── errors.go                          # FileNotFound, InvalidMimeType, FileTooLarge
│   │   │   ├── repository.go                      # FileRepository (CRUD, Move, Copy, Search)
│   │   │   └── events.go                          # file.created/updated/moved/deleted.v1
│   │   ├── folder/
│   │   │   ├── entity.go                          # parent_id, path, depth, owner, namespace_id
│   │   │   ├── errors.go                          # FolderNotFound, NameConflict
│   │   │   ├── repository.go                      # FolderRepository (CRUD, Tree ops)
│   │   │   └── events.go                          # folder.created/renamed/moved/deleted.v1
│   │   ├── upload/
│   │   │   ├── entity.go                          # resumable sessions (id, namespace_id, owner, status)
│   │   │   ├── chunk.go                           # chunks (part_no, size, checksum)
│   │   │   ├── resumable.go                       # offsets, etags, resumable rules
│   │   │   ├── errors.go                          # UploadSessionNotFound, ChunkOutOfOrder
│   │   │   ├── repository.go                      # UploadRepository
│   │   │   └── events.go                          # upload.started/chunk_appended/resumed/completed/aborted.v1
│   │   ├── version/
│   │   │   ├── entity.go                          # immutable versions (blob_id, promoted_at, author)
│   │   │   ├── errors.go                           # VersionNotFound, ImmutableVersion
│   │   │   ├── repository.go                      # VersionRepository
│   │   │   └── events.go                          # file_version.created/promoted/restored.v1
│   │
│   │   # =========================
│   │   # 🔑 ACCESS & SHARING
│   │   # =========================
│   │   ├── access_control/
│   │   │   ├── entity.go                          # ACL (subject, scope, action, resource)
│   │   │   ├── errors.go                          # PermissionDenied, ACLNotFound
│   │   │   ├── repository.go                      # ACL repository
│   │   │   └── events.go                          # access.granted/revoked/updated.v1
│   │   ├── share/
│   │   │   ├── entity.go                          # target=file/folder, grantee, scopes, expiry
│   │   │   ├── link.go                            # link token, expires_at, usage_count
│   │   │   ├── errors.go                          # ShareNotFound, ShareExpired
│   │   │   ├── repository.go                      # ShareRepository
│   │   │   └── events.go                          # share.created/revoked/expired.v1
│   │   ├── linking/
│   │   │   ├── signed_url.go                      # file_id, action, expires_at, signature, revoked_at
│   │   │   ├── audit_log.go                       # download audit: file_id, user_id, ip, ua, at
│   │   │   ├── errors.go                          # SignedURLInvalid, SignedURLExpired
│   │   │   ├── repository.go                      # LinkingRepository
│   │   │   └── events.go                          # signed_url.created/revoked; file.download.logged.v1
│   │   ├── lock/                                   # 🆕 short leases for co-edit/version ops
│   │   │   ├── entity.go                          # file_lock (file_id, owner, lease_id, expires_at)
│   │   │   ├── lease.go                           # short TTL helpers
│   │   │   └── repository.go                      # LockRepository
│   │
│   │   # =========================
│   │   # 🧪 CONTENT PIPELINE
│   │   # =========================
│   │   ├── media/
│   │   │   ├── entity.go                          # job (kind=image/video, pipeline, status, attempts)
│   │   │   ├── thumbnail.go                       # thumb generation plan
│   │   │   ├── variant.go                         # variants (size, codec, bitrate)
│   │   │   ├── errors.go                          # ProcessingFailed, UnsupportedFormat
│   │   │   ├── repository.go                      # MediaJob repository
│   │   │   └── events.go                          # media.processing_started/succeeded/failed; thumbnail.generated.v1
│   │   ├── extraction/                             # 🆕 OCR/EXIF/text extraction
│   │   │   ├── entity.go                          # ExtractionJob (tool, status)
│   │   │   ├── results.go                         # normalized text/metadata blobs
│   │   │   ├── repository.go                      # ExtractionRepository
│   │   │   └── events.go                          # extraction.started/succeeded/failed.v1
│   │   ├── artifact/                               # 🆕 derived/temporary outputs
│   │   │   ├── entity.go                          # type=preview/zip/report, file_id?, blob_id, ttl
│   │   │   ├── ttl.go                             # expiry/renewal calc
│   │   │   ├── repository.go                      # ArtifactRepository
│   │   │   └── events.go                          # artifact.created/expired.v1
│   │
│   │   # =========================
│   │   # 🛡️ SAFETY & POLICY
│   │   # =========================
│   │   ├── policy/
│   │   │   ├── entity.go                          # name, max_size_mb, allow_mime[], block_mime[], virus_scan, dlp
│   │   │   ├── dlp_pattern.go                     # regex/detectors
│   │   │   ├── result.go                          # violations, reasons
│   │   │   ├── errors.go                          # PolicyInvalid, PatternInvalid
│   │   │   ├── repository.go                      # PolicyRepository
│   │   │   └── events.go                          # file_policy.created/updated/violation_detected.v1
│   │   ├── scan/                                   # 🆕 AV + DLP scan domain
│   │   │   ├── entity.go                          # ScanJob (kind=av/dlp, status, severity, findings)
│   │   │   ├── results.go                         # findings (hash, match, offset, reason)
│   │   │   ├── repository.go                      # ScanRepository
│   │   │   └── events.go                          # file.scanned/quarantined/cleared.v1
│   │   ├── quarantine/                             # 🆕 isolation for suspicious files
│   │   │   ├── entity.go                          # Quarantine (file_id, reason, placed_by, released_at)
│   │   │   ├── repository.go                      # QuarantineRepository
│   │   │   └── events.go                          # quarantine.placed/released.v1
│   │   ├── file_flag/
│   │   │   ├── entity.go                          # reporter, reason, state
│   │   │   ├── reason.go                          # malware, copyright, policy_violation
│   │   │   ├── status.go                          # open, resolved, dismissed
│   │   │   ├── errors.go                          # FlagNotFound, InvalidFlagReason
│   │   │   ├── repository.go                      # FlagRepository
│   │   │   └── events.go                          # file_flag.submitted/resolved/dismissed.v1
│   │
│   │   # =========================
│   │   # 🧾 COMPLIANCE & LIFECYCLE
│   │   # =========================
│   │   ├── lifecycle/
│   │   │   ├── entity.go                          # Rule (state→retention_days, legal_hold, restore_window_days)
│   │   │   ├── soft_delete.go                     # deleted_at, restore_by
│   │   │   ├── legal_hold.go                      # placed_by, reason, expires_at
│   │   │   ├── errors.go                          # LifecycleRuleNotFound, RestoreWindowExceeded
│   │   │   ├── repository.go                      # LifecycleRepository
│   │   │   └── events.go                          # file.soft_deleted/restored; legal_hold.placed/removed.v1
│   │   ├── audit/                                  # 🆕 full operations trail (beyond download)
│   │   │   ├── entity.go                          # action, actor, resource, ip, ua, ts
│   │   │   ├── writer.go                          # append-only writer
│   │   │   └── queries.go                         # lookups/export
│   │
│   │   # =========================
│   │   # 🧱 STORAGE PRIMITIVES & OPS
│   │   # =========================
│   │   ├── blob/                                   # 🆕 content-addressed objects + de-dup
│   │   │   ├── entity.go                          # sha256, size, storage_class, ref_count, location
│   │   │   ├── repository.go                      # Put, GetByHash, Link, Unlink, MarkGC
│   │   │   └── events.go                          # blob.created/linked/unlinked/gc.v1
│   │   ├── reference/                              # 🆕 cross-service file refs
│   │   │   ├── entity.go                          # aggregate_type, aggregate_id, file_id, purpose
│   │   │   ├── repository.go                      # Add, Remove, ListByFile
│   │   │   └── events.go                          # reference.added/removed.v1
│   │   ├── quota/                                  # 🆕 tenant/user/org quotas
│   │   │   ├── entity.go                          # subject, hard_bytes, soft_bytes, used_bytes, window
│   │   │   ├── rules.go                           # enforce on upload/variant/artifact
│   │   │   ├── repository.go                      # QuotaRepository
│   │   │   └── events.go                          # quota.exceeded/adjusted.v1
│   │   ├── namespace/                              # 🆕 tenancy + data residency
│   │   │   ├── entity.go                          # tenant_id, bucket, data_zone, encryption_policy
│   │   │   ├── resolver.go                        # route writes by zone/bucket
│   │   │   └── repository.go                      # NamespaceRepository
│   │   ├── gc/                                     # 🆕 garbage collector
│   │   │   ├── entity.go                          # task (scope, state, started_at, finished_at, stats)
│   │   │   ├── planner.go                         # mark-sweep planner via ref_count/reference
│   │   │   └── events.go                          # gc.planned/run_started/run_completed.v1
│   │   └── encryption/                             # 🆕 key refs/rotation
│   │       ├── entity.go                          # file_id, key_id, version, rotated_at
│   │       ├── rotation.go                        # rotation jobs/state
│   │       └── repository.go                      # EncryptionRepository
│   │
│   # =============================
│   # 📋 APPLICATION LAYER (CQRS)
│   # =============================
│   ├── application/
│   │   # =========================
│   │   # 📡 EVENT CONSUMERS (INBOX)
│   │   # =========================
│   │   ├── eventhandler/
│   │   │   ├── user_handler.go                    # user.updated → refresh ACL/ownership context
│   │   │   ├── contract_handler.go                # contract.state.changed → lifecycle/holds
│   │   │   ├── admin_policy_handler.go            # admin.policy.updated → policy/DLP cache
│   │   │   └── admin_moderation_handler.go        # admin.moderation.actioned → quarantine/restore/revoke links
│   │
│   │   # =========================
│   │   # 🧠 USE CASES (COMMANDS/QUERIES)
│   │   # =========================
│   │   # -------- 📁 CORE FILE PRIMITIVES --------
│   │   ├── file/
│   │   │   ├── service.go                         # Create/Update/Delete/Move/Copy
│   │   │   ├── commands.go                        # CreateFile, UpdateFile, DeleteFile, MoveFile, CopyFile
│   │   │   ├── queries.go                         # GetFile, ListFiles, SearchFiles
│   │   │   ├── dto.go                             # FileDTO, SearchDTO
│   │   │   ├── mapper.go                          # Entity ↔ DTO
│   │   │   └── validators.go                      # names, size, policy
│   │   ├── folder/
│   │   │   ├── service.go                         # Create/Rename/Move/Delete
│   │   │   ├── commands.go                        # CreateFolder, RenameFolder, MoveFolder, DeleteFolder
│   │   │   ├── queries.go                         # GetFolder, ListFolderContents, SearchFolders
│   │   │   ├── dto.go                             # FolderDTO
│   │   │   ├── mapper.go                          # Entity ↔ DTO
│   │   │   └── validators.go                      # name/cycle/path depth
│   │   ├── upload/
│   │   │   ├── service.go                         # resumable/chunked workflows
│   │   │   ├── commands.go                        # StartUpload, AppendChunk, CompleteUpload, AbortUpload
│   │   │   ├── queries.go                         # GetUploadSession, GetUploadProgress
│   │   │   ├── chunked_upload.go                  # Append/Merge/Verify
│   │   │   ├── resumable.go                       # offsets/ETag
│   │   │   ├── dto.go                             # UploadSessionDTO, ProgressDTO
│   │   │   └── validators.go                      # order/size/checksum
│   │   ├── version/
│   │   │   ├── service.go                         # Create/Restore/Delete versions
│   │   │   ├── commands.go                        # CreateVersion, RestoreVersion, DeleteVersion
│   │   │   ├── queries.go                         # GetVersion, ListVersions
│   │   │   ├── dto.go                             # VersionDTO
│   │   │   ├── mapper.go                          # Entity ↔ DTO
│   │   │   └── validators.go                      # immutability rules
│   │   # -------- 🔑 ACCESS & SHARING --------
│   │   ├── access_control/
│   │   │   ├── service.go                         # Grant/Revoke ACL
│   │   │   ├── commands.go                        # GrantAccess, RevokeAccess
│   │   │   ├── queries.go                         # GetACL, ListACLs
│   │   │   ├── dto.go                             # ACLDTO
│   │   │   └── validators.go                      # scope/action checks
│   │   ├── share/
│   │   │   ├── service.go                         # Create/Revoke/Update share links
│   │   │   ├── commands.go                        # CreateShare, RevokeShare, UpdateShare
│   │   │   ├── queries.go                         # GetShare, ListSharesByFile
│   │   │   ├── dto.go                             # ShareDTO
│   │   │   └── validators.go                      # expiry/scopes/access
│   │   ├── linking/
│   │   │   ├── service.go                         # Create/ revoke signed URLs; audit logging
│   │   │   ├── commands.go                        # CreateSignedURL, RevokeSignedURL
│   │   │   ├── queries.go                         # GetSignedURL, ListSignedURLs, GetAuditLogs
│   │   │   ├── dto.go                             # SignedURLDTO, AuditLogDTO
│   │   │   └── validators.go                      # expiry/actions/scopes
│   │   ├── lock/
│   │   │   ├── service.go                         # Acquire/Renew/Release
│   │   │   ├── commands.go                        # AcquireLock, ReleaseLock
│   │   │   └── validators.go                      # lease bounds, ownership
│   │   # -------- 🧪 CONTENT PIPELINE --------
│   │   ├── media/
│   │   │   ├── service.go                         # ProcessImage/Video, GenerateThumbnail
│   │   │   ├── image_processor.go                 # image pipelines
│   │   │   ├── video_processor.go                 # video pipelines
│   │   │   ├── thumbnail_generator.go             # thumbnails
│   │   │   ├── commands.go                        # ProcessImage, ProcessVideo, GenerateThumbnail
│   │   │   ├── queries.go                         # GetMediaJob, ListMediaJobs
│   │   │   ├── dto.go                             # MediaJobDTO
│   │   │   └── validators.go                      # dimensions/codec/bitrate
│   │   ├── extraction/
│   │   │   ├── service.go                         # Start/Track extraction
│   │   │   ├── commands.go                        # StartExtraction
│   │   │   └── queries.go                         # GetExtractionJob
│   │   ├── artifact/
│   │   │   ├── service.go                         # Create zip/preview/report with TTL
│   │   │   ├── commands.go                        # CreateArtifact, ExpireArtifact
│   │   │   └── queries.go                         # GetArtifactsByFile
│   │   # -------- 🛡️ SAFETY & POLICY --------
│   │   ├── policy/
│   │   │   ├── service.go                         # Evaluate/Update/Get policies
│   │   │   ├── commands.go                        # SetPolicy, EnableDLP, DisableDLP
│   │   │   ├── queries.go                         # GetPolicy, ListPolicies
│   │   │   ├── dto.go                             # PolicyDTO, DLPResultDTO
│   │   │   └── validators.go                      # size/type/regexes
│   │   ├── scan/
│   │   │   ├── service.go                         # Queue scans, persist results, trigger quarantine
│   │   │   ├── commands.go                        # StartScan, RecordScanResult
│   │   │   └── queries.go                         # GetScanJob, ListFindings
│   │   ├── quarantine/
│   │   │   ├── service.go                         # Place/Release quarantine
│   │   │   └── commands.go                        # QuarantineFile, ReleaseQuarantine
│   │   ├── flag/
│   │   │   ├── service.go                         # Flag/Resolve/Dismiss
│   │   │   ├── commands.go                        # FlagFile, ResolveFlag, DismissFlag
│   │   │   ├── queries.go                         # GetFlags, GetFlag
│   │   │   └── validators.go                      # reason/state transitions
│   │   # -------- 🧾 COMPLIANCE & LIFECYCLE --------
│   │   ├── lifecycle/
│   │   │   ├── service.go                         # ApplyRules, SoftDelete, Restore, Place/RemoveLegalHold
│   │   │   ├── commands.go                        # DefineRule, UpdateRule, DeleteRule, SoftDeleteFile, RestoreFile, PlaceLegalHold
│   │   │   ├── queries.go                         # GetRules, GetFileLifecycle, GetLegalHolds
│   │   │   ├── dto.go                             # LifecycleRuleDTO, LegalHoldDTO
│   │   │   └── validators.go                      # retention/restore bounds
│   │   ├── audit/
│   │   │   ├── service.go                         # Append audit records, export
│   │   │   └── queries.go                         # ListAuditByResource
│   │   # -------- 🧱 STORAGE PRIMITIVES & OPS --------
│   │   ├── blob/
│   │   │   ├── service.go                         # PutFromStream, GetByHash, LinkToFile, UnlinkFromFile
│   │   │   ├── commands.go                        # PutBlob, LinkBlob, UnlinkBlob
│   │   │   └── queries.go                         # GetBlobByHash
│   │   ├── reference/
│   │   │   ├── service.go                         # Track references for safe delete
│   │   │   ├── commands.go                        # AddReference, RemoveReference
│   │   │   └── queries.go                         # ListReferencesByFile
│   │   ├── quota/
│   │   │   ├── service.go                         # Enforce/Adjust quotas
│   │   │   ├── commands.go                        # SetQuota, AdjustQuota
│   │   │   └── queries.go                         # GetQuota
│   │   ├── namespace/
│   │   │   ├── service.go                         # Resolve tenant → bucket/zone
│   │   │   └── queries.go                         # GetNamespace
│   │   ├── gc/
│   │   │   ├── service.go                         # Plan & run GC sweeps
│   │   │   └── commands.go                        # RunGC
│   │   └── encryption/
│   │       ├── service.go                         # Track key refs; schedule rotations
│   │       └── commands.go                        # RotateKeyRef
│   │
│   # =============================
│   # 🔌 INFRASTRUCTURE LAYER
│   # =============================
│   ├── infrastructure/
│   │   # =========================
│   │   # 🗄️ PERSISTENCE (POSTGRES)
│   │   # =========================
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       # 🧱 COMMON (DB BOOTSTRAP)
│   │   │       ├── connection.go                   # DSN & pooling
│   │   │       ├── transaction.go                  # TX helpers
│   │   │       ├── migrations.go                   # auto-migrate (versioned)
│   │   │       ├── version.go                      # schema version table
│   │   │       ├── safety.go                       # pre-flight checks (env/disk)
│   │   │       # 📁 CORE FILE PRIMITIVES
│   │   │       ├── file_repository.go              # FileRepository (GORM)
│   │   │       ├── folder_repository.go            # FolderRepository
│   │   │       ├── upload_repository.go            # UploadRepository
│   │   │       ├── version_repository.go           # VersionRepository
│   │   │       # 🔑 ACCESS & SHARING
│   │   │       ├── access_control_repository.go    # ACL
│   │   │       ├── share_repository.go             # Shares
│   │   │       ├── linking_repository.go           # Signed URLs + audit logs
│   │   │       ├── lock_repository.go              # 🆕 Locks
│   │   │       # 🧪 CONTENT PIPELINE
│   │   │       ├── media_repository.go             # Media jobs
│   │   │       ├── extraction_repository.go        # 🆕 Extraction jobs/results
│   │   │       ├── artifact_repository.go          # 🆕 Artifacts
│   │   │       # 🛡️ SAFETY & POLICY
│   │   │       ├── policy_repository.go            # Policies/DLP
│   │   │       ├── scan_repository.go              # 🆕 Scan jobs/results
│   │   │       ├── quarantine_repository.go        # 🆕 Quarantine
│   │   │       ├── file_flag_repository.go         # Flags
│   │   │       # 🧾 COMPLIANCE & LIFECYCLE
│   │   │       ├── lifecycle_repository.go         # Lifecycle rules/holds
│   │   │       ├── audit_repository.go             # 🆕 Audit records
│   │   │       # 🧱 STORAGE PRIMITIVES & OPS
│   │   │       ├── blob_repository.go              # 🆕 Blobs
│   │   │       ├── reference_repository.go         # 🆕 References
│   │   │       ├── quota_repository.go             # 🆕 Quotas
│   │   │       ├── namespace_repository.go         # 🆕 Namespaces
│   │   │       ├── gc_repository.go                # 🆕 GC tasks
│   │   │       └── encryption_repository.go        # 🆕 Encryption key refs
│   │   # =========================
│   │   # ⚡ CACHE (REDIS)
│   │   # =========================
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go                  # connection (pooling, retry)
│   │   │       ├── file_cache.go                  # file metadata (Get/Set/Invalidate, TTL)
│   │   │       ├── policy_cache.go                # 🆕 DLP/policy cache
│   │   │       ├── quota_cache.go                 # 🆕 hot-path usage counters
│   │   │       ├── lock_lease_cache.go            # 🆕 lock leases (short TTL)
│   │   │       └── signed_url_cache.go            # 🆕 presigned URL cache
│   │   # =========================
│   │   # 📨 MESSAGING (KAFKA)
│   │   # =========================
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go                    # uses platform-shared/inbox (dedupe, offsets)
│   │   │       ├── producer.go                    # uses platform-shared/outbox (reliable publishing)
│   │   │       ├── topics.go                      # topic constants (contracts/events; file.*, media.*, scan.*)
│   │   │       └── scram.go                       # SASL/SCRAM-256
│   │   # =========================
│   │   # 🪣 OBJECT STORAGE PROVIDERS
│   │   # =========================
│   │   ├── object_storage/
│   │   │   ├── local/
│   │   │   │   ├── storage.go                     # local FS (dev/test)
│   │   │   │   └── config.go                      # local config
│   │   │   ├── minio/
│   │   │   │   ├── client.go                      # self-hosted MinIO client (upload/download/presign)
│   │   │   │   └── config.go                      # endpoint/creds/buckets
│   │   │   ├── signer.go                          # abstraction to create/revoke signed URLs
│   │   │   └── provider.go                        # provider abstraction (local/minio)
│   │   # =========================
│   │   # 🎞️ MEDIA PROCESSING
│   │   # =========================
│   │   ├── media_processing/
│   │   │   ├── image/
│   │   │   │   ├── resizer.go                     # resize
│   │   │   │   ├── optimizer.go                   # compress/optimize
│   │   │   │   └── watermark.go                   # watermark
│   │   │   ├── video/
│   │   │   │   ├── transcoder.go                  # transcode
│   │   │   │   └── thumbnail.go                   # video thumbnails
│   │   │   └── processor.go                       # job orchestration
│   │   # =========================
│   │   # 🔐 SECURITY / SCANNERS
│   │   # =========================
│   │   ├── virus_scan/
│   │   │   └── clamav.go                          # ClamAV integration (stream/hash)
│   │   ├── dlp/
│   │   │   ├── regex_engine.go                    # regex-based detectors (PII/PCI)
│   │   │   └── provider.go                        # pluggable DLP providers (custom/3rd-party)
│   │   # =========================
│   │   # 📤 OUTBOX (PLATFORM-SHARED)
│   │   # =========================
│   │   └── outbox/
│   │       ├── processor.go                       # ❌ REMOVED → use platform-shared/outbox/forwarder.go
│   │       └── scheduler.go                       # ❌ REMOVED → use platform-shared/outbox/scheduler.go
│   │
│   # =============================
│   # 🌐 HTTP INTERFACE (v1)
│   # =============================
│   ├── interfaces/
│   │   └── http/
│   │       └── v1/
│   │           # =========================
│   │           # 🧭 HANDLERS
│   │           # =========================
│   │           ├── handlers/
│   │           # -------- 📁 CORE FILE PRIMITIVES --------
│   │           │   ├── file_handler.go               # files (GET, POST, PATCH, DELETE /files)
│   │           │   ├── upload_handler.go             # resumable/chunked
│   │           │   ├── download_handler.go           # GET /download/:id
│   │           │   ├── folder_handler.go             # folder CRUD/navigation
│   │           │   ├── version_handler.go            # file versions
│   │           # -------- 🔑 ACCESS & SHARING --------
│   │           │   ├── access_handler.go             # ACL grant/revoke
│   │           │   ├── share_handler.go              # share links
│   │           │   ├── linking_handler.go            # signed URLs & audit logs
│   │           │   ├── lock_handler.go               # 🆕 locks
│   │           # -------- 🧪 CONTENT PIPELINE --------
│   │           │   ├── media_handler.go              # process/thumbnail/status
│   │           │   ├── extraction_handler.go         # 🆕 start/inspect extraction
│   │           │   ├── artifact_handler.go           # 🆕 create/list/expire artifacts
│   │           # -------- 🛡️ SAFETY & POLICY --------
│   │           │   ├── policy_handler.go             # policies/DLP
│   │           │   ├── scan_handler.go               # 🆕 trigger/inspect scans
│   │           │   ├── quarantine_handler.go         # 🆕 place/release quarantine
│   │           │   ├── flag_handler.go               # moderation flags
│   │           # -------- 🧾 COMPLIANCE & LIFECYCLE --------
│   │           │   ├── lifecycle_handler.go          # soft delete/restore/legal holds
│   │           │   ├── audit_handler.go              # 🆕 admin audit export
│   │           # -------- 🧱 STORAGE PRIMITIVES & OPS --------
│   │           │   ├── quota_handler.go              # 🆕 quotas (admin/user views)
│   │           │   ├── namespace_handler.go          # 🆕 namespaces (tenant/bucket/zone)
│   │           │   └── health_handler.go             # /health, /ready, /live
│   │           # =========================
│   │           # 🧰 MIDDLEWARE
│   │           # =========================
│   │           ├── middleware/
│   │           │   ├── auth.go                       # JWT auth (pkg/auth)
│   │           │   ├── rbac.go                       # role checks (pkg/auth authorizer)
│   │           │   ├── cors.go                       # CORS (platform-shared/ginx)
│   │           │   ├── rate_limit.go                 # token-bucket rate limiting
│   │           │   ├── logging.go                    # structured logs (platform-shared/ginx/logging)
│   │           │   ├── error_handler.go              # unified error responses
│   │           │   ├── request_id.go                 # request ID (platform-shared/ginx/requestid)
│   │           │   └── file_size_limit.go            # enforce max upload size
│   │           # =========================
│   │           # 📨 RESPONSES
│   │           # =========================
│   │           ├── responses/
│   │           │   ├── success.go                    # success wrappers (platform-shared/httpx/response)
│   │           │   ├── error.go                      # error mapping (platform-shared/httpx/errors)
│   │           │   └── pagination.go                 # pagination (platform-shared/httpx/pagination)
│   │           # =========================
│   │           # 🗺️ ROUTES
│   │           # =========================
│   │           └── routes/
│   │               # 📁 CORE FILE PRIMITIVES
│   │               ├── file_routes.go                # /files/*
│   │               ├── upload_routes.go              # /upload/*
│   │               ├── download_routes.go            # /download/*
│   │               ├── folder_routes.go              # /folders/*
│   │               ├── version_routes.go             # /versions/*
│   │               # 🔑 ACCESS & SHARING
│   │               ├── access_routes.go              # /acl/*
│   │               ├── share_routes.go               # /shares/*
│   │               ├── linking_routes.go             # /links/*
│   │               ├── lock_routes.go                # /locks/*
│   │               # 🧪 CONTENT PIPELINE
│   │               ├── media_routes.go               # /media/*
│   │               ├── extraction_routes.go          # /extractions/*
│   │               ├── artifact_routes.go            # /artifacts/*
│   │               # 🛡️ SAFETY & POLICY
│   │               ├── policy_routes.go              # /policies/*
│   │               ├── scan_routes.go                # /scans/*
│   │               ├── quarantine_routes.go          # /quarantine/*
│   │               ├── flag_routes.go                # /flags/*
│   │               # 🧾 COMPLIANCE & LIFECYCLE
│   │               ├── lifecycle_routes.go           # /lifecycle/*
│   │               ├── audit_routes.go               # /audit/*
│   │               # 🧱 STORAGE PRIMITIVES & OPS
│   │               ├── quota_routes.go               # /quota/*
│   │               ├── namespace_routes.go           # /namespaces/*
│   │               └── router.go                     # Gin router wiring + common middleware
│
├── config/
│   ├── default.yaml                                 # Default configuration
│   ├── dev.yaml                                     # Development overrides
│   └── prod.yaml                                    # Production overrides
│
├── dapr/                                            # Dapr components split by environment
│   ├── local/                                       # For dapr run
│   │   ├── pubsub.yaml                              # Kafka pub/sub component
│   │   └── statestore.yaml                          # State store component
│   └── k8s/                                         # For kubectl apply
│       ├── pubsub.yaml                              # Kafka with scopes: ["storage-be"]
│       ├── statestore.yaml                          # State store with scopes
│       └── secrets.yaml                             # Dapr secret store
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                          # Kubernetes Deployment
│       ├── service.yaml                             # Kubernetes Service
│       ├── configmap.yaml                           # ConfigMap
│       ├── secrets.yaml                             # Secrets
│       ├── hpa.yaml                                 # HPA
│       ├── pdb.yaml                                 # PDB
│       ├── pvc.yaml                                 # PersistentVolumeClaim
│       └── servicemonitor.yaml                      # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                               # Setup local environment
│   ├── get-secrets.sh                               # Fetch secrets
│   ├── seed-data.sh                                 # Seed sample data
│   └── cleanup-orphaned.sh                          # Sweep unreferenced dev files
│
├── tests/
│   # =============================
│   # ✅ TEST SUITES
│   # =============================
│   ├── unit/
│   │   ├── domain/                                 # domain unit tests (file, blob, quota, policy, etc.)
│   │   ├── application/                            # service-level tests (commands/queries)
│   │   └── infrastructure/                         # repos, caches
│   ├── integration/
│   │   ├── handlers/                               # HTTP integration tests
│   │   └── repositories/                           # Postgres repo integration tests
│   └── e2e/
│       └── scenarios/                              # upload→scan→quarantine→share→download flows
│
├── docs/
│   ├── README.md                                   # Service overview
│   ├── API.md                                      # API documentation
│   ├── EVENTS.md                                   # published: file.*, blob.*, artifact.*, scan.*; consumed: user.*, admin.*, contract.*
│   ├── MIGRATIONS.md                               # Migration history
│   ├── SCHEMA.md                                   # Database schema & ERD
│   ├── upload-flow.md                              # Resumable/chunked flows
│   ├── media-processing.md                         # Pipelines & presets
│   └── RUNBOOK.md                                  # Ops (GC, quarantine, scan queues)
│
├── pkg/
│   # =============================
│   # 🧰 LOCAL UTILITIES
│   # =============================
│   ├── errors/
│   │   ├── errors.go                               # Service-specific errors
│   │   └── codes.go                                # Error codes
│   ├── utils/
│   │   ├── validator.go                            # Validation helpers
│   │   ├── file_utils.go                           # Path manipulation, extension extraction
│   │   ├── mime_detector.go                        # MIME detection
│   │   └── hash.go                                 # Hash calc (MD5, SHA256)
│   └── constants/
│       ├── mime_types.go                           # Supported MIME types
│       └── README.md                               # Constants provenance (contracts/events elsewhere)
│
├── .github/
│   └── workflows/
│       ├── ci.yml                                  # CI workflow
│       └── cd.yml                                  # CD workflow
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
│   # =============================
│   # 🚀 APP ENTRYPOINTS
│   # =============================
│   ├── api/
│   │   └── main.go                               # 📝 Gin + Dapr + Elasticsearch; uses internal/config & platform-shared/logging
│   └── worker/                                    # 🆕 background runner
│       └── main.go                                # 🆕 runs hygiene/backfill/snapshot jobs; leader election safe
│
├── internal/
│   # =============================
│   # 🔧 CONFIGURATION (LOAD FIRST)
│   # =============================
│   ├── config/
│   │   ├── schema.go                              # Typed Config (App, Server, Postgres, Kafka, Redis, Elasticsearch, ML, IndexLifecycle)
│   │   ├── loader.go                              # Viper loader (flags → env → file → defaults)
│   │   └── docs/
│   │       └── CONFIGURATION.md                   # All ENV vars, defaults, examples
│   │
│   # =============================
│   # 🏛️ DOMAIN LAYER (DDD)
│   # =============================
│   ├── domain/
│   │   # =========================
│   │   # 📚 CORE INDEX ARTIFACTS
│   │   # =========================
│   │   ├── search_index/
│   │   │   ├── entity.go                          # Index meta (kind, alias, version, settings, mappings_hash, visibility)
│   │   │   ├── job_index.go                       # Job search document fields
│   │   │   ├── user_index.go                      # User/Freelancer search document fields
│   │   │   ├── errors.go                          # IndexNotFound, DocumentConflict
│   │   │   ├── repository.go                      # IndexRepository interface
│   │   │   └── events.go                          # search.document.indexed/reindexed; search.index.visibility.changed/archived.v1
│   │   ├── portfolio_index/                       # 🆕 discoverable portfolios/works
│   │   │   ├── entity.go                          # Title, skills[], media_refs, engagement, recency
│   │   │   ├── repository.go                      # PortfolioIndexRepository
│   │   │   └── events.go                          # search.portfolio.indexed/reindexed/removed.v1
│   │
│   │   # =========================
│   │   # 🔍 QUERY INPUT & LOGGING
│   │   # =========================
│   │   ├── search_query/
│   │   │   ├── entity.go                          # Query logs (filters/sort/lang/latency, anon user hash)
│   │   │   ├── filters.go                         # Filter structs (rate, skills, languages, location, badges)
│   │   │   ├── errors.go                          # InvalidFilter, BadSort
│   │   │   ├── repository.go                      # SearchQueryRepository
│   │   │   └── events.go                          # search.query.logged / search.query.alert.triggered.v1
│   │   ├── saved_search/                          # ✅ keep existing (explicit saved queries)
│   │   │   ├── entity.go                          # SavedSearch{user_id,name,query_json,schedule,active}
│   │   │   ├── alert.go                           # Alert settings (window, channel)
│   │   │   ├── errors.go                          # SavedSearchNotFound, DuplicateName
│   │   │   ├── repository.go                      # SavedSearchRepository
│   │   │   └── events.go                          # search.saved_search.created/updated/deleted/alert.sent.v1
│   │   ├── multi_language/                        # 🆕 language & analyzer profiles
│   │   │   ├── entity.go                          # DetectedLang{code,confidence}, AnalyzerProfile{id,lang,tokenizer,filters}
│   │   │   ├── detector.go                        # LangDetect(text) → DetectedLang
│   │   │   ├── transliteration.go                 # ar↔en transliteration helpers (names/skills)
│   │   │   ├── repository.go                      # AnalyzerProfileRepository
│   │   │   └── events.go                          # search.lang.profile.updated.v1
│   │   ├── speller/                               # 🆕 spelling & “did you mean”
│   │   │   ├── entity.go                          # SpellingCandidate{term,score,source}
│   │   │   ├── dictionary.go                      # BK-tree / ES suggesters; build/merge
│   │   │   ├── repository.go                      # SpellerRepository
│   │   │   └── events.go                          # search.speller.dictionary.updated.v1
│   │   ├── query_rewrite/                         # 🆕 synonyms/stopwords/rewrites
│   │   │   ├── entity.go                          # RewriteRule{pattern,action,weight,lang,enabled}
│   │   │   ├── rules_engine.go                    # ApplyRules(query, lang) → rewritten query
│   │   │   ├── repository.go                      # RewriteRuleRepository
│   │   │   └── events.go                          # search.query_rewrite.updated.v1
│   │
│   │   # =========================
│   │   # 🧠 RANKING, LTR & CONTROLS
│   │   # =========================
│   │   ├── ltr/                                   # Learning-to-Rank signals
│   │   │   ├── entity.go                          # LTRSignal{doc_id, features, label, effective_at}
│   │   │   ├── signal_source.go                   # Event→feature extractors
│   │   │   ├── feature_store.go                   # Persisted features snapshots
│   │   │   ├── errors.go                          # SignalOutOfRange, FeatureStale
│   │   │   ├── repository.go                      # LTRRepository
│   │   │   └── events.go                          # search.ltr.signal.recorded/features.updated.v1
│   │   ├── promotion/                             # 🆕 editorial boosts/demotions (non-paid)
│   │   │   ├── entity.go                          # Promotion{scope,subject_id,boost,expires_at,reason}
│   │   │   ├── policy.go                          # Guardrails (no paid placements)
│   │   │   ├── repository.go                      # PromotionRepository
│   │   │   └── events.go                          # search.promotion.upserted/expired.v1
│   │
│   │   # =========================
│   │   # 👤 PERSONALIZATION & RECS
│   │   # =========================
│   │   ├── recommendation/
│   │   │   ├── entity.go                          # Recommendation records
│   │   │   ├── score.go                           # Score formula components
│   │   │   ├── reason.go                          # Human-readable reasons
│   │   │   ├── feedback.go                        # User feedback on recs
│   │   │   ├── errors.go                          # ModelUnavailable, ScoringFailed
│   │   │   ├── repository.go                      # RecommendationRepository
│   │   │   └── events.go                          # search.recommendation.generated/feedback.recorded.v1
│   │   ├── recommendation_model/
│   │   │   ├── entity.go                          # ML model metadata/versions
│   │   │   ├── feature.go                         # Feature vectors catalog
│   │   │   ├── training_data.go                   # Training snapshots
│   │   │   ├── errors.go                          # VersionNotFound, FeatureMismatch
│   │   │   ├── repository.go                      # RecommendationModelRepository
│   │   │   └── events.go                          # search.model.version.registered/deployed/rolled_back/training_data.updated.v1
│   │   ├── user_preference/
│   │   │   ├── entity.go                          # Explicit user prefs
│   │   │   ├── implicit_signals.go                # Clicks/views (summaries)
│   │   │   ├── explicit_preferences.go            # Explicit fields (rate, availability)
│   │   │   ├── errors.go                          # PreferenceNotFound
│   │   │   ├── repository.go                      # UserPreferenceRepository
│   │   │   └── events.go                          # search.preference.updated.v1
│   │   ├── personalization/                       # 🆕 profile rollups
│   │   │   ├── entity.go                          # Recent activity, preferred rates, tags
│   │   │   ├── cold_start.go                      # Defaults for new users
│   │   │   ├── errors.go                          # Personalization errors
│   │   │   ├── repository.go                      # PersonalizationRepository
│   │   │   └── events.go                          # search.personalization.cold_start.applied/updated.v1
│   │
│   │   # =========================
│   │   # 🔎 DISCOVERY & SIMILARITY
│   │   # =========================
│   │   ├── matching/
│   │   │   ├── entity.go                          # Job↔Freelancer matches (summary)
│   │   │   ├── criteria.go                        # Criteria (skills, rate, availability)
│   │   │   ├── score_breakdown.go                 # Factor-level scores
│   │   │   ├── errors.go                          # CriteriaInvalid
│   │   │   ├── repository.go                      # MatchingRepository
│   │   │   └── events.go                          # search.match.calculated/accepted/dismissed.v1
│   │   ├── similarity/
│   │   │   ├── entity.go                          # Similar jobs/users (links)
│   │   │   ├── vector.go                          # Vector fields (embeddings)
│   │   │   ├── errors.go                          # Similarity errors
│   │   │   ├── repository.go                      # SimilarityRepository
│   │   │   └── events.go                          # search.similarity.computed/model.updated.v1
│   │   ├── feed/
│   │   │   ├── entity.go                          # User feeds (items, ranks)
│   │   │   ├── item.go                            # Feed item data
│   │   │   ├── personalization.go                 # Feed-specific personalization
│   │   │   ├── errors.go                          # FeedNotFound
│   │   │   ├── repository.go                      # FeedRepository
│   │   │   └── events.go                          # search.feed.item.added/removed/updated.v1
│   │   ├── trending/
│   │   │   ├── entity.go                          # Trending jobs/skills
│   │   │   ├── calculator.go                      # Calculate trending
│   │   │   ├── errors.go                          # Trending errors
│   │   │   ├── repository.go                      # TrendingRepository
│   │   │   └── events.go                          # search.trending.calculated/updated.v1
│   │   ├── suggestion/
│   │   │   └── entity.go                          # Placeholder for suggestion types (lives mostly in application)
│   │
│   │   # =========================
│   │   # 🧭 TAXONOMY & FACETS
│   │   # =========================
│   │   ├── taxonomy/
│   │   │   ├── entity.go                          # Skills/Categories; normalized
│   │   │   ├── category.go                        # Category tree
│   │   │   ├── synonym.go                         # Synonyms & aliases
│   │   │   ├── typos.go                           # Typo tolerance rules
│   │   │   ├── errors.go                          # SkillNotFound, AliasConflict
│   │   │   ├── repository.go                      # TaxonomyRepository
│   │   │   └── events.go                          # search.taxonomy.updated/synonym.changed/typo_rules.updated.v1
│   │   ├── facets/
│   │   │   ├── entity.go                          # Facet defs (price bands, availability, languages, badges)
│   │   │   ├── banding.go                         # Banding logic
│   │   │   ├── tz_overlap.go                      # Timezone overlap
│   │   │   ├── errors.go                          # InvalidBand, UnknownBadge
│   │   │   ├── repository.go                      # FacetRepository
│   │   │   └── events.go                          # search.facets.definition.updated/banding.updated/tz_overlap.updated.v1
│   │
│   │   # =========================
│   │   # 🛡️ SAFETY, HYGIENE & COMPLIANCE
│   │   # =========================
│   │   ├── hygiene/
│   │   │   ├── entity.go                          # Hygiene tasks (incremental, dedupe, archival, visibility)
│   │   │   ├── incremental.go                     # Change markers (version, changed fields)
│   │   │   ├── dedupe.go                          # Fingerprints / duplicates
│   │   │   ├── visibility.go                      # Visibility states
│   │   │   ├── errors.go                          # Hygiene errors
│   │   │   ├── repository.go                      # HygieneRepository
│   │   │   └── events.go                          # search.index.hygiene.*.v1
│   │   ├── compliance/
│   │   │   ├── entity.go                          # Erasure tasks / holds
│   │   │   ├── erasure.go                         # Remove docs on request
│   │   │   └── events.go                          # compliance.erasure.requested/completed.v1
│   │   ├── safety_filters/                        # 🆕 query-time visibility gates
│   │   │   ├── entity.go                          # SafetyRule{kind,subject,action,ttl}
│   │   │   ├── engine.go                          # allow/deny/mask evaluation
│   │   │   └── repository.go                      # SafetyFiltersRepository
│   │
│   │   # =========================
│   │   # 📍 GEO & INTENT
│   │   # =========================
│   │   ├── geo/
│   │   │   ├── entity.go                          # GeoPoint{lat,lon}, GeoPolicy{radius,max_results}
│   │   │   ├── scorer.go                          # Distance decay helpers
│   │   │   └── repository.go                      # GeoRepository
│   │   ├── query_intent/
│   │   │   ├── entity.go                          # Intent{job|talent|navigational, confidence}
│   │   │   ├── classifier.go                      # Heuristics/ML-lite classifier
│   │   │   └── repository.go                      # QueryIntentRepository
│   │
│   │   # =========================
│   │   # 🛠️ OPERATIONS & EXPLAINABILITY
│   │   # =========================
│   │   ├── index_lifecycle/
│   │   │   ├── entity.go                          # IndexSchema{kind,version,alias,mappings_hash}
│   │   │   ├── rollover.go                        # Create v{n+1}, reindex, alias swap
│   │   │   ├── snapshot.go                        # Snapshot/restore via MinIO/local
│   │   │   ├── repository.go                      # IndexLifecycleRepository
│   │   │   └── events.go                          # search.index.rolled_over/snapshotted/restored.v1
│   │   ├── backfill/
│   │   │   ├── entity.go                          # BackfillRun{id,scope,state,counters}
│   │   │   ├── planner.go                         # Partition planning (id/time windows)
│   │   │   └── events.go                          # search.backfill.started/completed.v1
│   │   └── explainability/
│   │       ├── entity.go                          # Explanation{doc_id,factors[],scores[]}
│   │       ├── builder.go                         # Human-readable “why” strings
│   │       └── events.go                          # search.result.explained.v1
│   │
│   # =============================
│   # 📋 APPLICATION LAYER (CQRS)
│   # =============================
│   ├── application/
│   │   # =========================
│   │   # 📡 EVENT CONSUMERS (INBOX)
│   │   # =========================
│   │   ├── eventhandler/
│   │   │   ├── job_handler.go                     # Consumes: job.posted/updated/closed → index jobs & refresh facets
│   │   │   ├── user_handler.go                    # Consumes: user.updated → refresh freelancer docs (badges, visibility)
│   │   │   ├── review_handler.go                  # Consumes: review.* → update rating aggregates, LTR signals
│   │   │   ├── entitlement_handler.go             # Consumes: subscription.feature.changed → gate facets/features
│   │   │   ├── admin_content_handler.go           # Consumes: admin.content.actioned → hide/unhide docs
│   │   │   ├── admin_flags_handler.go             # Consumes: admin.feature_flag/threshold/experiment.updated → refresh toggles
│   │   │   ├── storage_lifecycle_handler.go       # Consumes: file.lifecycle.soft_deleted/restored → sync asset visibility in docs
│   │   │   ├── compliance_handler.go              # Consumes: user.erasure.requested → remove indexed docs
│   │   │   └── taxonomy_handler.go                # Consumes: admin.taxonomy.updated → refresh synonyms/rewrites
│   │
│   │   # =========================
│   │   # 🧠 USE CASES (COMMANDS/QUERIES)
│   │   # =========================
│   │   # ---- 🔍 SEARCH EXECUTION ----
│   │   ├── search/
│   │   │   ├── service.go                         # Run job/talent searches end-to-end
│   │   │   ├── job_search.go                      # ES DSL for jobs
│   │   │   ├── freelancer_search.go               # ES DSL for users
│   │   │   ├── query_builder.go                   # Query + filter composition
│   │   │   ├── facet_builder.go                   # Aggregations/facets builder
│   │   │   ├── dto.go                             # Request/Response DTOs
│   │   │   ├── mapper.go                          # Map ES hits → DTO
│   │   │   ├── commands.go                        # ExecuteSearch, SaveSearch
│   │   │   ├── queries.go                         # GetSearchResults, GetSearchSuggestions
│   │   │   └── validators.go                      # Validate filters/sorts/facets
│   │   # ---- 🧾 INDEXING ----
│   │   ├── indexing/
│   │   │   ├── service.go                         # Index/update/remove docs
│   │   │   ├── job_indexer.go                     # Job document mappers
│   │   │   ├── user_indexer.go                    # User document mappers
│   │   │   ├── bulk_indexer.go                    # Bulk ops & backpressure
│   │   │   ├── dto.go                             # Indexing DTOs
│   │   │   ├── commands.go                        # IndexJob, IndexUser, ReindexAll
│   │   │   ├── queries.go                         # GetIndexStatus, GetDocByID
│   │   │   └── validators.go                      # Payload validation
│   │   # ---- 💾 SAVED SEARCHES ----
│   │   ├── saved_search/
│   │   │   ├── service.go                         # Create/Update/Delete saved searches + alerts
│   │   │   ├── commands.go                        # CreateSavedSearch, UpdateSavedSearch, DeleteSavedSearch
│   │   │   ├── queries.go                         # GetSavedSearches, GetSavedSearch
│   │   │   ├── validators.go                      # Name uniqueness, schedule bounds
│   │   │   ├── dto.go                             
│   │   │   └── mapper.go                          
│   │   # ---- 👤 PERSONALIZATION & RECS ----
│   │   ├── recommendation/
│   │   │   ├── service.go                         # Generate & explain recs
│   │   │   ├── job_recommender.go                 # Jobs → freelancer
│   │   │   ├── freelancer_recommender.go          # Freelancers → client
│   │   │   ├── collaborative_filtering.go         # CF signals
│   │   │   ├── content_based.go                   # Content-based signals
│   │   │   ├── hybrid_recommender.go              # Hybrid approach
│   │   │   ├── scoring_engine.go                  # Score fusion
│   │   │   ├── personalization.go                 # Per-user personalization
│   │   │   ├── diversity_optimizer.go             # Result diversity
│   │   │   ├── cold_start_handler.go              # New users/jobs
│   │   │   ├── dto.go
│   │   │   ├── ml_model.go                        # Model selection
│   │   │   ├── commands.go                        # GenerateRecommendations, RecordFeedback
│   │   │   ├── queries.go                         # GetRecommendations, GetReasons
│   │   │   └── validators.go                      # Limits & params
│   │   ├── user_preference/
│   │   │   ├── service.go                         # Update/Get prefs
│   │   │   ├── commands.go                        # UpdatePreferences
│   │   │   └── queries.go                         # GetPreferences
│   │   ├── personalization/
│   │   │   ├── service.go                         # UpdateProfile, ComputeDefaults
│   │   │   ├── commands.go                        # UpdatePersonalizationProfile
│   │   │   └── queries.go                         # GetPersonalizationProfile
│   │   # ---- 🔎 DISCOVERY ----
│   │   ├── matching/
│   │   │   ├── service.go                         # Orchestrate criteria evaluators
│   │   │   ├── matcher.go                         # Match pipeline
│   │   │   ├── criteria_evaluator.go              # Evaluate match criteria
│   │   │   ├── skill_matcher.go                   # Skill overlap
│   │   │   ├── experience_matcher.go              # Experience fit
│   │   │   ├── rate_matcher.go                    # Rate fit
│   │   │   ├── availability_matcher.go            # Availability/tz fit
│   │   │   ├── score_calculator.go                # Overall score
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                        # CreateMatchRun
│   │   │   ├── queries.go                         # GetMatchesForJob/User
│   │   │   └── validators.go                      # Criteria completeness
│   │   ├── similarity/
│   │   │   ├── service.go                         # Similar jobs/users
│   │   │   ├── job_similarity.go
│   │   │   ├── user_similarity.go
│   │   │   ├── vector_calculator.go               # Embedding generation
│   │   │   ├── dto.go
│   │   │   ├── commands.go                        # RebuildSimilarityVectors
│   │   │   ├── queries.go                         # GetSimilarJobs/Users
│   │   │   └── validators.go                      # k/threshold
│   │   ├── feed/
│   │   │   ├── service.go                         # Generate personalized feed
│   │   │   ├── generator.go                       # Candidate gather
│   │   │   ├── ranking.go                         # Ranker
│   │   │   ├── freshness_scorer.go                # Freshness factor
│   │   │   ├── relevance_scorer.go                # Relevance factor
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                        # GenerateFeed
│   │   │   ├── queries.go                         # GetFeed
│   │   │   └── validators.go                      # Window/size checks
│   │   ├── trending/
│   │   │   ├── service.go                         # Trending recompute
│   │   │   ├── calculator.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                        # RecomputeTrending
│   │   │   ├── queries.go                         # GetTrending
│   │   │   └── validators.go                      # Window/minSupport
│   │   ├── suggestion/
│   │   │   ├── service.go                         # Autocomplete suggestions
│   │   │   ├── dto.go
│   │   │   ├── cache_warmer.go                    # Warm cache
│   │   │   ├── commands.go                        # WarmSuggestionCache
│   │   │   ├── queries.go                         # GetSuggestions
│   │   │   └── validators.go                      # Prefix/lang checks
│   │   ├── portfolio_index/
│   │   │   ├── service.go                         # Index & search portfolios
│   │   │   ├── commands.go                        # IndexPortfolioDoc
│   │   │   └── queries.go                         # SearchPortfolios
│   │   # ---- 🧭 TAXONOMY & FACETS ----
│   │   ├── taxonomy/
│   │   │   ├── service.go                         # Upsert skills/categories/aliases
│   │   │   ├── commands.go                        # UpsertSkill/Category, Add/RemoveAlias
│   │   │   ├── queries.go                         # GetSkill, ListSkills, GetCategoryTree
│   │   │   ├── validators.go                      # Alias conflicts, edit distance
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── facets/
│   │   │   ├── service.go                         # Build facets; apply banding
│   │   │   ├── commands.go                        # DefineFacet, UpdateFacetBands
│   │   │   ├── queries.go                         # GetFacetsForQuery
│   │   │   ├── validators.go                      # Band & tz rules
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   # ---- 🛡️ SAFETY & HYGIENE ----
│   │   ├── hygiene/
│   │   │   ├── service.go                         # Incrementals, dedupe, archive/visibility
│   │   │   ├── commands.go                        # RunIncrementalUpdate, RunDedup, ArchiveDoc, SetVisibility
│   │   │   ├── queries.go                         # GetHygieneStatus, GetDocHistory
│   │   │   ├── validators.go                      # State transitions
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── compliance/
│   │   │   ├── service.go                         # Process erasure
│   │   │   ├── commands.go                        # RequestErasure
│   │   │   └── queries.go                         # GetErasureStatus
│   │   ├── safety_filters/
│   │   │   ├── service.go                         # ApplySafetyRules to result sets
│   │   │   ├── commands.go                        # UpsertSafetyRule, RemoveSafetyRule
│   │   │   └── queries.go                         # ListSafetyRules
│   │   # ---- 📍 GEO & INTENT ----
│   │   ├── geo/
│   │   │   ├── service.go                         # ApplyGeoFilters, ScoreByDistance
│   │   │   └── validators.go                      # Radius/bounds checks
│   │   ├── query_intent/
│   │   │   ├── service.go                         # ClassifyIntent (job vs talent vs navigational)
│   │   │   └── queries.go                         # GetIntent
│   │   # ---- 🛠️ OPERATIONS ----
│   │   ├── index_lifecycle/
│   │   │   ├── service.go                         # Rollover, Snapshot, Restore
│   │   │   ├── commands.go                        # RolloverIndex, SnapshotIndex, RestoreIndex
│   │   │   └── queries.go                         # GetIndexSchema
│   │   ├── backfill/
│   │   │   ├── service.go                         # Plan & run backfills
│   │   │   └── commands.go                        # StartBackfill
│   │   └── explainability/
│   │       ├── service.go                         # BuildExplanation (ES _explain wrappers)
│   │       └── queries.go                         # GetExplanation
│   │
│   # =============================
│   # 🔌 INFRASTRUCTURE LAYER
│   # =============================
│   ├── infrastructure/
│   │   # =========================
│   │   # 🗄️ PERSISTENCE (POSTGRES)
│   │   # =========================
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       # 🧱 COMMON (DB BOOTSTRAP)
│   │   │       ├── connection.go                   # DSN & pooling
│   │   │       ├── transaction.go                  # TX helpers
│   │   │       ├── migrations.go                   # Auto-migrate (versioned)
│   │   │       ├── version.go                      # Schema version table
│   │   │       ├── safety.go                       # Pre-flight checks (env/disk)
│   │   │       # 🔍 QUERY INPUT & LOGGING
│   │   │       ├── search_query_repository.go      # SearchQueryRepository impl
│   │   │       ├── saved_search_repository.go      # SavedSearchRepository impl
│   │   │       ├── speller_repository.go           # 🆕 SpellerRepository impl
│   │   │       ├── query_rewrite_repository.go     # 🆕 RewriteRuleRepository impl
│   │   │       ├── multi_language_repository.go    # 🆕 AnalyzerProfileRepository impl
│   │   │       # 👤 PERSONALIZATION & RECS
│   │   │       ├── recommendation_repository.go
│   │   │       ├── recommendation_model_repository.go
│   │   │       ├── user_preference_repository.go
│   │   │       ├── personalization_repository.go   # 🆕
│   │   │       ├── ltr_repository.go               # 🆕 LTRRepository impl
│   │   │       # 🔎 DISCOVERY
│   │   │       ├── matching_repository.go
│   │   │       ├── similarity_repository.go
│   │   │       ├── feed_repository.go
│   │   │       ├── trending_repository.go
│   │   │       ├── portfolio_repository.go         # 🆕 portfolios index meta (PG)
│   │   │       # 🧭 TAXONOMY & FACETS
│   │   │       ├── taxonomy_repository.go          # 🆕
│   │   │       ├── facets_repository.go            # 🆕
│   │   │       # 🛡️ SAFETY/HYGIENE/COMPLIANCE
│   │   │       ├── hygiene_repository.go           # 🆕
│   │   │       ├── compliance_repository.go        # 🆕
│   │   │       ├── safety_filters_repository.go    # 🆕
│   │   │       # 📍 GEO & INTENT
│   │   │       ├── geo_repository.go               # 🆕
│   │   │       └── query_intent_repository.go      # 🆕
│   │   # =========================
│   │   # 🔎 ELASTICSEARCH
│   │   # =========================
│   │   ├── elasticsearch/
│   │   │   ├── client.go                          # ES client (connection, retry)
│   │   │   ├── index_manager.go                   # Create/update mappings/settings
│   │   │   ├── alias_router.go                    # Read/write aliases (single-tenant)
│   │   │   ├── snapshot_client.go                 # Snapshot/restore to MinIO/local
│   │   │   ├── job_mapper.go                      # Job entity → ES document
│   │   │   ├── user_mapper.go                     # User entity → ES document
│   │   │   └── config.go                          # Hosts, auth, timeouts
│   │   # =========================
│   │   # ⚡ CACHE (REDIS)
│   │   # =========================
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go                  # Pooling, retries, instrumentation
│   │   │       ├── search_cache.go                # Cache search results (TTL)
│   │   │       ├── suggestion_cache.go            # Autocomplete cache
│   │   │       ├── feed_cache.go                  # User feeds
│   │   │       ├── recommendation_cache.go        # Recommendations
│   │   │       ├── taxonomy_cache.go              # Skills/categories & aliases
│   │   │       ├── ltr_cache.go                   # LTR feature snapshots
│   │   │       ├── speller_cache.go               # Did-you-mean cache
│   │   │       ├── rewrite_rules_cache.go         # Query rewrite rules
│   │   │       └── promotion_cache.go             # Active promotions
│   │   # =========================
│   │   # 📨 MESSAGING (KAFKA)
│   │   # =========================
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go                    # 📝 uses platform-shared/inbox for dedupe & offsets
│   │   │       ├── producer.go                    # 📝 uses platform-shared/outbox for reliable publishing
│   │   │       ├── topics.go                      # 📝 contracts/events: job.*, user.*, review.*, experiment.*, compliance.*
│   │   │       └── scram.go                       # SASL/SCRAM-256 auth
│   │   # =========================
│   │   # 🤖 ML UTILITIES
│   │   # =========================
│   │   └── ml/
│   │       ├── model_loader.go                    # Load models (disk/env)
│   │       ├── predictor.go                       # Predictions/rerank
│   │       ├── feature_extractor.go               # Extract features
│   │       ├── trainer.go                         # Offline training jobs
│   │       └── evaluator.go                       # Evaluate model performance
│   │
│   # =============================
│   # 🌐 HTTP INTERFACE (v1)
│   # =============================
│   ├── interfaces/
│   │   └── http/
│   │       └── v1/
│   │           # =========================
│   │           # 🧭 HANDLERS
│   │           # =========================
│   │           ├── handlers/
│   │           # ---- SEARCH & INDEXING ----
│   │           │   ├── search_handler.go           # POST /search/jobs, /search/freelancers
│   │           │   ├── indexing_handler.go         # POST/GET admin index ops
│   │           │   ├── saved_search_handler.go     # 🆕 CRUD saved searches & alerts
│   │           # ---- RECS / DISCOVERY ----
│   │           │   ├── recommendation_handler.go   # GET /recommendations/*
│   │           │   ├── matching_handler.go         # POST/GET /matching/*
│   │           │   ├── feed_handler.go             # GET /feed/*
│   │           │   ├── trending_handler.go         # GET /trending/*
│   │           │   ├── similarity_handler.go       # GET /similarity/*
│   │           │   ├── suggestion_handler.go       # GET /suggestions/*
│   │           │   ├── portfolio_handler.go        # 🆕 GET /search/portfolios
│   │           # ---- TAXONOMY & FACETS ----
│   │           │   ├── taxonomy_handler.go         # GET/PUT /taxonomy/*
│   │           │   ├── facets_handler.go           # GET /facets/*
│   │           # ---- INPUT HELPERS ----
│   │           │   ├── speller_handler.go          # 🆕 GET /speller/did-you-mean
│   │           │   ├── rewrite_handler.go          # 🆕 POST /rewrites/preview
│   │           │   ├── language_handler.go         # 🆕 GET /languages/supported
│   │           # ---- SAFETY & OPS ----
│   │           │   ├── hygiene_handler.go          # POST/GET /hygiene/*
│   │           │   ├── compliance_handler.go       # POST /compliance/erasure
│   │           │   ├── lifecycle_handler.go        # POST /indices/rollover|snapshot|restore
│   │           │   ├── promotion_handler.go        # PUT/DELETE /promotions/*
│   │           │   ├── explain_handler.go          # GET /debug/explain
│   │           │   └── health_handler.go           # GET /health, /ready, /live
│   │           # =========================
│   │           # 🧰 MIDDLEWARE
│   │           # =========================
│   │           ├── middleware/
│   │           │   ├── auth.go                     # JWT verification (pkg/auth)
│   │           │   ├── rbac.go                     # Role checks
│   │           │   ├── cors.go                     # CORS (platform-shared/ginx)
│   │           │   ├── rate_limit.go               # Token bucket rate limiting
│   │           │   ├── logging.go                  # Structured request logging
│   │           │   ├── error_handler.go            # Error → HTTP mapping
│   │           │   └── request_id.go               # X-Request-ID propagation
│   │           # =========================
│   │           # 📨 RESPONSES
│   │           # =========================
│   │           ├── responses/
│   │           │   ├── success.go                  # platform-shared/httpx/response
│   │           │   ├── error.go                    # platform-shared/httpx/errors
│   │           │   └── pagination.go               # platform-shared/httpx/pagination
│   │           # =========================
│   │           # 🗺️ ROUTES (SECTIONED)
│   │           # =========================
│   │           ├── routes/
│   │           │   # ---- SEARCH & INDEXING ----
│   │           │   ├── search_routes.go            # /search/* (jobs|freelancers)
│   │           │   ├── indexing_routes.go          # /indexing/* (admin)
│   │           │   ├── saved_search_routes.go      # 🆕 /saved-searches/*
│   │           │   # ---- RECS / DISCOVERY ----
│   │           │   ├── recommendation_routes.go    # /recommendations/*
│   │           │   ├── matching_routes.go          # /matching/*
│   │           │   ├── feed_routes.go              # /feed/*
│   │           │   ├── trending_routes.go          # /trending/*
│   │           │   ├── similarity_routes.go        # /similarity/*
│   │           │   ├── suggestion_routes.go        # /suggestions/*
│   │           │   ├── portfolio_routes.go         # 🆕 /search/portfolios/*
│   │           │   # ---- TAXONOMY & FACETS ----
│   │           │   ├── taxonomy_routes.go          # /taxonomy/*
│   │           │   ├── facets_routes.go            # /facets/*
│   │           │   # ---- INPUT HELPERS ----
│   │           │   ├── speller_routes.go           # 🆕 /speller/*
│   │           │   ├── rewrite_routes.go           # 🆕 /rewrites/*
│   │           │   ├── language_routes.go          # 🆕 /languages/*
│   │           │   # ---- SAFETY & OPS ----
│   │           │   ├── hygiene_routes.go           # /hygiene/*
│   │           │   ├── compliance_routes.go        # /compliance/*
│   │           │   ├── lifecycle_routes.go         # /indices/*
│   │           │   ├── promotion_routes.go         # /promotions/*
│   │           │   └── websocket_routes.go         # (optional) /ws if needed later
│   │           └── router.go                       # Gin engine wiring + common middleware
│   │
│   └── (end internal)
│
├── config/
│   ├── default.yaml                               # Defaults
│   ├── dev.yaml                                   # Dev overrides
│   └── prod.yaml                                  # Prod overrides
│
├── dapr/
│   ├── local/
│   │   ├── pubsub.yaml                            # Kafka pub/sub
│   │   └── statestore.yaml                        # State store
│   └── k8s/
│       ├── pubsub.yaml                            # Scopes: ["search-be"]
│       ├── statestore.yaml                        # Scopes
│       └── secrets.yaml                           # Secret store
│
├── elasticsearch/
│   # =============================
│   # 🧭 ES ARTIFACTS (STATIC)
│   # =============================
│   ├── mappings/
│   │   ├── jobs.json                               # Job index mapping
│   │   ├── users.json                              # User index mapping
│   │   └── portfolios.json                         # 🆕 Portfolio index mapping
│   └── analyzers/
│       └── custom_analyzers.json                   # ICU folding, ar/en analyzers, n-grams
│
├── ml_models/
│   ├── job_recommendation/
│   │   ├── model.pkl
│   │   ├── features.json
│   │   └── metadata.json
│   ├── freelancer_recommendation/
│   │   ├── model.pkl
│   │   ├── features.json
│   │   └── metadata.json
│   └── matching/
│       ├── model.pkl
│       ├── features.json
│       └── metadata.json
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                        # Deployment
│       ├── service.yaml                           # Service
│       ├── configmap.yaml                         # ConfigMap
│       ├── secrets.yaml                           # Secrets
│       ├── hpa.yaml                               # HPA
│       ├── pdb.yaml                               # PDB
│       └── servicemonitor.yaml                    # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                             # Local environment bootstrap
│   ├── get-secrets.sh                             # Fetch secrets
│   ├── seed-data.sh                               # Seed PG + ES for dev
│   ├── create-indices.sh                          # Create ES indices & aliases
│   ├── reindex-all.sh                             # Rollover & reindex pipeline
│   └── train-models.sh                            # Offline training
│
├── tests/
│   # =============================
│   # ✅ TEST SUITES
│   # =============================
│   ├── unit/
│   │   ├── domain/                                # Rewrite/speller/facets/taxonomy/LTR unit tests
│   │   ├── application/                           # Search/indexing/recs services tests
│   │   └── infrastructure/                        # Repos, ES mappers tests
│   ├── integration/
│   │   ├── handlers/                              # HTTP integration tests
│   │   └── repositories/                          # Postgres repositories tests
│   └── e2e/
│       └── scenarios/                             # Search→click→LTR signal flows
│
├── docs/
│   ├── README.md                                  # Service overview
│   ├── api.md                                     # API reference
│   ├── events.md                                  # Published/Consumed events (contracts/events)
│   ├── search-algorithms.md                       # Ranking & filtering details
│   ├── recommendation-engine.md                   # Recommenders & signals
│   ├── recommendation-types.md                    # Placement taxonomy
│   ├── matching-algorithm.md                      # Matching pipeline
│   ├── ml-models.md                               # Models & features
│   ├── elasticsearch-setup.md                     # ES setup & snapshots
│   ├── MIGRATIONS.md                              # Migration history
│   ├── SCHEMA.md                                  # Database schema
│   └── RUNBOOK.md                                 # Ops procedures (reindex, hygiene, snapshots, backfills)
│
├── pkg/
│   # =============================
│   # 🧰 LOCAL UTILITIES
│   # =============================
│   ├── errors/
│   │   ├── errors.go                               # Service-specific error helpers
│   │   └── codes.go                                # Error codes
│   ├── utils/
│   │   ├── validator.go                            # Validation utilities
│   │   ├── text_analyzer.go                        # Tokenize/normalize helpers
│   │   ├── normalizer.go                           # Data normalization
│   │   └── vector_math.go                          # Vector ops & distances
│   └── constants/
│       ├── indices.go                               # Index/alias names
│       └── README.md                                # (Note) events/topics come from contracts/events
│
├── .github/
│   └── workflows/
│       ├── ci.yml                                  # CI pipeline
│       └── cd.yml                                  # CD pipeline
│
├── go.mod                                          # 📝 imports pkg/auth, platform-shared, contracts/events
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
│   # =============================
│   # 🚀 APP ENTRYPOINTS
│   # =============================
│   ├── api/
│   │   └── main.go                               # 📝 Gin + Dapr + Postgres bootstrap; loadConfig(); platform-shared/logging
│   └── worker/                                    # 🆕 background workers (cron-ish)
│       └── main.go                                # 🆕 open/close double-blind windows, send reminders, rebuild stats/reputation
│
├── internal/
│   # =============================
│   # 🔧 CONFIGURATION (LOAD FIRST)
│   # =============================
│   ├── config/
│   │   ├── schema.go                              # Typed Config: App, Server, Postgres, Kafka, Redis, Windows, Limits
│   │   ├── loader.go                              # Viper loader: flags → env → file → defaults
│   │   └── docs/CONFIGURATION.md                  # ENV vars, defaults, examples, sample commands
│   │
│   # =============================
│   # 🏛️ DOMAIN LAYER (DDD)
│   # =============================
│   ├── domain/
│   │   # =========================
│   │   # ⭐ CORE REVIEW PRIMITIVES
│   │   # =========================
│   │   ├── review/
│   │   │   ├── entity.go                           # Aggregate: ids, body, ratings ver, status(draft|published|withdrawn), created/edited
│   │   │   ├── rating.go                           # Per-dimension values + normalized score
│   │   │   ├── enums.go                            # ReviewType(client→freelancer|freelancer→client), Status, Category
│   │   │   ├── response.go                         # Single public response from subject (when policy allows)
│   │   │   ├── helpful.go                          # Helpful votes with actor dedupe & 24h change window
│   │   │   ├── errors.go                           # 🆕 ReviewNotFound, WindowClosed, AlreadySubmitted, NotAuthor
│   │   │   ├── repository.go                       # CRUD + list by contract/user + helpful ops (interface)
│   │   │   └── events.go                           # 🆕 review.submitted/edited/published/withdrawn/helpful.v1
│   │   ├── eligibility/                            # 🆕 who may review whom (and when)
│   │   │   ├── entity.go                           # EligibilityResult{allowed,bool,reason,not_before,not_after}
│   │   │   ├── policy.go                           # Contract/payment state checks; daily limits; cooldowns
│   │   │   ├── errors.go                           # NotEligible, WindowNotOpen, RateLimited
│   │   │   ├── repository.go                       # Persist last decision per (author,contract) for auditability
│   │   │   └── events.go                           # 🆕 review.eligibility.policy.updated.v1
│   │   ├── review_draft/                           # 🆕 save-before-submit
│   │   │   ├── entity.go                           # Draft{author,contract_id,body,rating_map,updated_at}
│   │   │   ├── errors.go                           # DraftNotFound, TooLarge
│   │   │   └── repository.go                       # Upsert/Get/Delete
│   │
│   │   # =========================
│   │   # 🧮 RATING & CRITERIA
│   │   # =========================
│   │   ├── rating/
│   │   │   ├── entity.go                           # Criteria set (id, category, version, dims with weights & scales)
│   │   │   ├── criteria.go                         # Versioned criteria definitions per job category
│   │   │   ├── aggregation.go                      # Rollups per subject (mean, weighted mean, Wilson)
│   │   │   ├── versioning.go                       # 🆕 freeze criteria version on each review
│   │   │   ├── errors.go                           # CriteriaNotFound, InvalidWeight, VersionMismatch
│   │   │   ├── repository.go                       # Store criteria versions & aggregates
│   │   │   └── events.go                           # 🆕 rating.criteria.updated / rating.aggregated.v1
│   │
│   │   # =========================
│   │   # 🏅 BADGES & REPUTATION
│   │   # =========================
│   │   ├── badge/
│   │   │   ├── entity.go                           # Badge meta (id, name, levels, description, active)
│   │   │   ├── criteria.go                         # Rule expressions referencing aggregates/reputation
│   │   │   ├── types.go                            # Built-ins: Top Rated, Rising Talent, …
│   │   │   ├── level.go                            # Thresholds per level (Bronze/Silver/Gold)
│   │   │   ├── errors.go                           # InvalidRule, BadgeNotFound
│   │   │   ├── repository.go                       # Badge & criteria CRUD
│   │   │   └── events.go                           # 🆕 badge.created/criteria.updated/level.updated/archived.v1
│   │   ├── user_badge/
│   │   │   ├── entity.go                           # Assignment{user_id,badge_id,level,awarded_at,expires_at}
│   │   │   ├── achievement_date.go                  # Derived dates (renewals)
│   │   │   ├── errors.go                            # AlreadyHasBadge, NotEligible
│   │   │   ├── repository.go                        # Assign/Revoke/ListByUser
│   │   │   └── events.go                            # 🆕 badge.awarded/revoked.v1
│   │   ├── reputation/
│   │   │   ├── entity.go                            # Score components + total
│   │   │   ├── score_calculator.go                  # Bayesian prior + decay + recency boost
│   │   │   ├── history.go                           # Time series for trends
│   │   │   ├── decay.go                              # Outlier dampening & rolling windows
│   │   │   ├── top_rated_rules.go                    # Eligibility rules for "Top Rated"
│   │   │   ├── errors.go                             
│   │   │   ├── repository.go                        # Upsert score/history
│   │   │   └── events.go                             # 🆕 reputation.updated / top_rated.eligibility.changed.v1
│   │
│   │   # =========================
│   │   # 🔒 DOUBLE-BLIND & WINDOWS
│   │   # =========================
│   │   ├── double_blind/
│   │   │   ├── entity.go                            # Window{contract_id, open_at, close_at, submitted_client, submitted_freelancer}
│   │   │   ├── window.go                            # Publish when both sides submit or on timeout
│   │   │   ├── policy.go                            # Durations, grace periods, extension rules
│   │   │   ├── errors.go                            # WindowNotOpen, PublishBlocked
│   │   │   ├── repository.go                        # Window persistence
│   │   │   └── events.go                            # 🆕 double_blind.submitted/published/timeout.v1
│   │   ├── review_reminder/
│   │   │   ├── entity.go                            # Reminder{contract_id,user_id,scheduled_for,status}
│   │   │   ├── errors.go                            # ReminderNotFound, AlreadySent
│   │   │   ├── repository.go                        # Schedule/MarkSent/Failures
│   │   │   └── events.go                            # 🆕 review.reminder.scheduled/sent/failed.v1
│   │
│   │   # =========================
│   │   # 🗣️ PRIVATE FEEDBACK
│   │   # =========================
│   │   ├── feedback/
│   │   │   ├── entity.go                            # Private feedback (hidden from public)
│   │   │   ├── category.go                          # Confidential categories
│   │   │   ├── nps.go                                # 🆕 NPS-style numeric score
│   │   │   ├── errors.go
│   │   │   ├── repository.go
│   │   │   └── events.go                             # 🆕 private_feedback.submitted/updated.v1
│   │
│   │   # =========================
│   │   # 🛡️ SAFETY & GOVERNANCE
│   │   # =========================
│   │   ├── flag/
│   │   │   ├── entity.go                            # Flag{review_id,reason,reporter, state}
│   │   │   ├── reason.go                            # Reason enum + custom text
│   │   │   ├── errors.go
│   │   │   ├── repository.go                        # Open/Resolve/Dismiss
│   │   │   └── events.go                             # 🆕 review.flag.submitted/resolved/dismissed.v1
│   │   ├── moderation/
│   │   │   ├── entity.go                            # Moderation state machine
│   │   │   ├── auto_flag.go                         # Heuristics: bursts, swaps, duplicate IPs
│   │   │   ├── signals.go                           # Evidence signals kept for decisioning
│   │   │   ├── errors.go
│   │   │   ├── repository.go                        # Set state; store decisions
│   │   │   └── events.go                             # 🆕 review.abuse.auto_flagged/state_changed.v1
│   │   ├── appeal/                                  # 🆕 structured challenges
│   │   │   ├── entity.go                            # Appeal{review_id, raised_by, reason, state, outcome}
│   │   │   ├── policy.go                            # Time windows + required evidence
│   │   │   ├── repository.go
│   │   │   └── events.go                             # 🆕 review.appeal.opened/decided.v1
│   │   ├── evidence/                                # 🆕 file attachments for moderation/appeals
│   │   │   ├── entity.go                            # Evidence{file_id(storage-be), type, notes}
│   │   │   └── repository.go
│   │   ├── redaction/                               # 🆕 PII masking
│   │   │   ├── policy.go                            # Allowlist/denylist + length caps
│   │   │   ├── masker.go                            # Apply redaction to review/response text
│   │   │   └── events.go                             # 🆕 review.redacted.v1
│   │   ├── compliance/                              # 🆕 DSAR/erasure/export
│   │   │   ├── entity.go                            # ErasureTask, ExportTask, Status
│   │   │   ├── erasure.go                           # Hard/soft delete policy + tombstones
│   │   │   └── events.go                             # 🆕 compliance.erasure.requested/completed.v1
│   │   ├── audit/                                   # 🆕 immutable trail
│   │   │   └── trail.go                             # who/what/when for edits, moderation, publishes
│   │
│   │   # =========================
│   │   # 📊 STATS & EXPOSURE
│   │   # =========================
│   │   ├── stats/
│   │   │   ├── entity.go                            # Counts, histograms, recency
│   │   │   ├── aggregates.go                        # Recalc logic & snapshotting
│   │   │   ├── errors.go
│   │   │   ├── repository.go
│   │   │   └── events.go                             # 🆕 review.stats.updated/recalculated.v1
│   │   ├── featured/                                # 🆕 editorial (non-paid) highlights
│   │   │   ├── entity.go                            # FeaturedReview{review_id, boost, expires_at, reason}
│   │   │   ├── policy.go                            # Guardrails (no pay-to-boost ever)
│   │   │   └── repository.go
│   │
│   │   └── outbox/
│   │       ├── entity.go                            # ❌ REMOVED (use platform-shared/outbox/entity.go)
│   │       └── repository.go                        # ❌ REMOVED (use platform-shared/outbox/repository.go)
│   │
│   # =============================
│   # 📋 APPLICATION LAYER (CQRS)
│   # =============================
│   ├── application/
│   │   # =========================
│   │   # 📡 EVENT CONSUMERS (INBOX)
│   │   # =========================
│   │   ├── eventhandler/
│   │   │   ├── contract_handler.go                  # Consumes: contract.ended → open double-blind window for both parties
│   │   │   ├── payment_handler.go                   # Consumes: payment.captured → qualify/open window (hourly/fixed)
│   │   │   ├── admin_moderation_handler.go          # Consumes: admin.moderation.actioned → remove/restore; recalc aggregates
│   │   │   ├── admin_flags_handler.go               # Consumes: admin.feature_flag/threshold/experiment.updated → refresh toggles
│   │   │   ├── search_projection_handler.go         # 🆕 On publish/withdraw → emit aggregates to search-be
│   │   │   └── compliance_handler.go                # 🆕 Consumes: user.erasure.requested → cascade DSAR into reviews
│   │   #
│   │   # =========================
│   │   # 🧠 USE CASES BY DOMAIN
│   │   # =========================
│   │   # ---- CORE REVIEWS ----
│   │   ├── review/
│   │   │   ├── service.go                           # Orchestrates create/edit/publish/withdraw & helpful votes
│   │   │   ├── commands.go                          # CreateReview, EditReview, PublishReview, WithdrawReview, VoteHelpful
│   │   │   ├── queries.go                           # GetReview, ListByUser, ListByContract, SearchReviews
│   │   │   ├── validators.go                        # Body length, PII mask, eligibility gate, double-blind state
│   │   │   ├── mapper.go                            # Entity ↔ DTO (includes rating version binding)
│   │   │   ├── dto.go                                # ReviewDTO, NewReviewDTO, ResponseDTO, HelpfulDTO
│   │   │   └── business_rules.go                    # Cross-cutting: one-per-contract-per-side, response window policy
│   │   ├── review_draft/                            # 🆕
│   │   │   ├── service.go                           # Save/Clear/Load draft (redaction applied on read)
│   │   │   ├── commands.go                          # SaveDraft, ClearDraft
│   │   │   ├── queries.go                           # GetDraft
│   │   │   ├── validators.go                        # Max size, allowed placeholders
│   │   │   ├── mapper.go                            # Draft ↔ DTO
│   │   │   └── dto.go                                # ReviewDraftDTO
│   │   ├── eligibility/                             # 🆕
│   │   │   ├── service.go                           # CheckEligibility(author, contract)
│   │   │   ├── queries.go                           # GetEligibility
│   │   │   ├── validators.go                        # Rate limit keying, window bounds
│   │   │   └── dto.go                                # EligibilityDTO
│   │   # ---- RATING/CRITERIA ----
│   │   ├── rating/
│   │   │   ├── service.go                           # Criteria CRUD + aggregate pipeline
│   │   │   ├── aggregator.go                        # Rebuild aggregates for user/contract
│   │   │   ├── calculator.go                        # Weighted means & Wilson intervals
│   │   │   ├── commands.go                          # UpsertCriteria, RecalculateForContract
│   │   │   ├── queries.go                           # GetCriteria, GetAggregates
│   │   │   ├── validators.go                        # Weight sum=1, value in scale bounds
│   │   │   ├── mapper.go                            # Criteria/aggregate ↔ DTO
│   │   │   └── dto.go                                # CriteriaDTO, AggregateDTO
│   │   # ---- BADGES & USER BADGES ----
│   │   ├── badge/
│   │   │   ├── service.go                           # Evaluate rules; emit award/revoke
│   │   │   ├── achievement_checker.go               # Check badge eligibility from aggregates/reputation
│   │   │   ├── commands.go                          # AwardBadge, RevokeBadge, UpsertBadge
│   │   │   ├── queries.go                           # GetBadge, ListBadges
│   │   │   ├── validators.go                        # Rule safety, level thresholds
│   │   │   ├── mapper.go                            # Badge ↔ DTO
│   │   │   └── dto.go                                # BadgeDTO, AwardDTO
│   │   ├── user_badge/                              # 🆕
│   │   │   ├── service.go                           # Read/refresh assignments
│   │   │   ├── queries.go                           # ListUserBadges
│   │   │   ├── mapper.go                            # Assignment ↔ DTO
│   │   │   └── dto.go                                # UserBadgeDTO
│   │   # ---- REPUTATION ----
│   │   ├── reputation/
│   │   │   ├── service.go                           # Compute + persist score
│   │   │   ├── calculator.go                        # Totals + components
│   │   │   ├── updater.go                           # Apply decay & outlier handling
│   │   │   ├── commands.go                          # RecomputeUserReputation, ApplyDecay
│   │   │   ├── queries.go                           # GetReputation, GetHistory
│   │   │   ├── validators.go                        # Decay params bounds
│   │   │   ├── mapper.go                            # Reputation ↔ DTO
│   │   │   └── dto.go                                # ReputationDTO, HistoryPointDTO
│   │   # ---- DOUBLE-BLIND & REMINDERS ----
│   │   ├── double_blind/
│   │   │   ├── service.go                           # OpenWindow, SubmitSide, PublishIfReady
│   │   │   ├── commands.go                          # OpenWindow, SubmitClient, SubmitFreelancer, ForcePublish
│   │   │   ├── queries.go                           # GetWindow, GetPublishState
│   │   │   ├── validators.go                        # Time & state transitions
│   │   │   └── dto.go                                # WindowDTO, PublishStateDTO
│   │   ├── review_reminder/
│   │   │   ├── service.go                           # Plan & send nudges
│   │   │   ├── commands.go                          # ScheduleReminder, SendReminder
│   │   │   └── dto.go                                # ReminderDTO
│   │   # ---- PRIVATE FEEDBACK ----
│   │   ├── feedback/
│   │   │   ├── service.go
│   │   │   ├── commands.go                          # SubmitPrivateFeedback
│   │   │   ├── queries.go                           # GetPrivateFeedback
│   │   │   ├── validators.go                        # NPS bounds; forbid PII
│   │   │   ├── mapper.go                            # Feedback ↔ DTO
│   │   │   └── dto.go                                # PrivateFeedbackDTO
│   │   # ---- SAFETY/GOVERNANCE ----
│   │   ├── flag/
│   │   │   ├── service.go                           # File/resolve/dismiss
│   │   │   ├── commands.go                          # FlagReview, ResolveFlag, DismissFlag
│   │   │   ├── queries.go                           # GetFlags, GetFlag
│   │   │   ├── validators.go                        # Reason whitelist; reviewer role
│   │   │   ├── mapper.go                            # Flag ↔ DTO
│   │   │   └── dto.go                                # FlagDTO
│   │   ├── moderation/
│   │   │   ├── service.go                           # AutoFlag & SetState
│   │   │   ├── commands.go                          # AutoFlagReview, SetModerationState
│   │   │   ├── queries.go                           # GetModerationState, ListFlags
│   │   │   ├── validators.go                        # Threshold/transition checks
│   │   │   ├── mapper.go                            # Moderation ↔ DTO
│   │   │   └── dto.go                                # ModerationDTO
│   │   ├── appeal/
│   │   │   ├── service.go                           # Open/decide
│   │   │   ├── commands.go                          # OpenAppeal, DecideAppeal
│   │   │   ├── queries.go                           # GetAppeal, ListAppeals
│   │   │   ├── validators.go                        # Window & evidence checks
│   │   │   ├── mapper.go                            # Appeal ↔ DTO
│   │   │   └── dto.go                                # AppealDTO
│   │   ├── evidence/
│   │   │   ├── service.go                           # Link to storage-be
│   │   │   ├── commands.go                          # AttachEvidence, RemoveEvidence
│   │   │   ├── queries.go                           # ListEvidenceForReview
│   │   │   ├── validators.go                        # File kind/size; ownership
│   │   │   ├── mapper.go                            # Evidence ↔ DTO
│   │   │   └── dto.go                                # EvidenceDTO
│   │   ├── redaction/
│   │   │   ├── service.go                           # ApplyMask(review/response)
│   │   │   ├── validators.go                        # Allowlist/denylist compile
│   │   │   └── dto.go                                # RedactedTextDTO
│   │   ├── compliance/
│   │   │   ├── service.go                           # DSAR: erasure/export orchestration
│   │   │   ├── commands.go                          # RequestErasure, ExportData
│   │   │   ├── queries.go                           # GetTaskStatus
│   │   │   ├── validators.go                        # Subject ownership check
│   │   │   └── dto.go                                # ErasureTaskDTO, ExportTaskDTO
│   │   # ---- STATS & EXPOSURE ----
│   │   ├── stats/
│   │   │   ├── service.go                           # Rebuild/serve aggregates
│   │   │   ├── commands.go                          # RebuildStats
│   │   │   ├── queries.go                           # GetStats (by user/time window)
│   │   │   ├── mapper.go                            # Stats ↔ DTO
│   │   │   └── dto.go                                # StatsDTO
│   │   ├── featured/
│   │   │   ├── service.go                           # Editorial pin/unpin (policy-guarded)
│   │   │   ├── commands.go                          # FeatureReview, UnfeatureReview
│   │   │   ├── queries.go                           # ListFeatured
│   │   │   ├── validators.go                        # No paid bias; expiry in future
│   │   │   ├── mapper.go                            # Featured ↔ DTO
│   │   │   └── dto.go                                # FeaturedReviewDTO
│   │   # ---- AUDIT READS ----
│   │   ├── audit/                                   # 🆕 read-only views
│   │   │   ├── service.go                           # Fetch audit trail entries
│   │   │   └── queries.go                           # GetAuditTrailForReview
│   │
│   # =============================
│   # 🔌 INFRASTRUCTURE LAYER
│   # =============================
│   ├── infrastructure/
│   │   # =========================
│   │   # 🗄️ PERSISTENCE (POSTGRES)
│   │   # =========================
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       # 🧱 COMMON (DB BOOTSTRAP)
│   │   │       ├── connection.go                    # DSN & pooling (env from config)
│   │   │       ├── transaction.go                   # WithTransaction helpers
│   │   │       ├── migrations.go                    # 📝 Auto-migrate with versioning
│   │   │       ├── version.go                       # SchemaVersion table
│   │   │       ├── safety.go                        # Env & disk checks before migration
│   │   │       # ⭐ CORE & WINDOWS
│   │   │       ├── review_repository.go
│   │   │       ├── review_draft_repository.go       # 🆕
│   │   │       ├── eligibility_repository.go        # 🆕
│   │   │       ├── double_blind_repository.go
│   │   │       ├── review_reminder_repository.go
│   │   │       # 🧮 RATING / REPUTATION / BADGES
│   │   │       ├── rating_repository.go
│   │   │       ├── badge_repository.go
│   │   │       ├── user_badge_repository.go
│   │   │       ├── reputation_repository.go
│   │   │       # 🗣️ PRIVATE FEEDBACK
│   │   │       ├── feedback_repository.go
│   │   │       # 🛡️ SAFETY & GOVERNANCE
│   │   │       ├── flag_repository.go
│   │   │       ├── moderation_repository.go
│   │   │       ├── appeal_repository.go            # 🆕
│   │   │       ├── evidence_repository.go          # 🆕
│   │   │       ├── redaction_repository.go         # 🆕 policies
│   │   │       ├── compliance_repository.go        # 🆕 erasure/export tasks
│   │   │       ├── audit_trail_repository.go       # 🆕 immutable audit log
│   │   │       # 📊 STATS & EXPOSURE
│   │   │       ├── stats_repository.go
│   │   │       └── featured_repository.go          # 🆕
│   │   # =========================
│   │   # ⚡ CACHE (REDIS)
│   │   # =========================
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go                   # Pooling, retry, timeouts
│   │   │       ├── review_cache.go                 # Review by ID
│   │   │       ├── reputation_cache.go             # Reputation scores
│   │   │       ├── stats_cache.go                  # Aggregates by user
│   │   │       ├── double_blind_cache.go           # Window state
│   │   │       ├── eligibility_cache.go            # 🆕 Eligibility result (short TTL)
│   │   │       └── badge_cache.go                  # Badge lookups
│   │   # =========================
│   │   # 📨 MESSAGING (KAFKA)
│   │   # =========================
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go                     # 📝 platform-shared/inbox (dedupe, ack)
│   │   │       ├── producer.go                     # 📝 platform-shared/outbox (exactly-once-ish)
│   │   │       ├── topics.go                       # 📝 contracts/events: review.*, badge.*, reputation.*, compliance.*
│   │   │       └── scram.go                        # SASL/SCRAM-256 auth
│   │   # =========================
│   │   # 🔗 EXTERNAL CLIENTS
│   │   # =========================
│   │   └── storage/
│   │       └── client.go                           # Minimal HTTP client to storage-be for evidence files
│   │
│   # =============================
│   # 🌐 HTTP INTERFACE (v1)
│   # =============================
│   ├── interfaces/
│   │   └── http/
│   │       └── v1/
│   │           # =========================
│   │           # 🧭 HANDLERS
│   │           # =========================
│   │           ├── handlers/
│   │           # ---- CORE REVIEWS ----
│   │           │   ├── review_handler.go            # POST/PUT/DELETE/GET reviews + helpful vote
│   │           │   ├── draft_handler.go             # 🆕 save/clear/get drafts
│   │           │   ├── eligibility_handler.go       # 🆕 GET eligibility for contract
│   │           # ---- RATING/REPUTATION/BADGES ----
│   │           │   ├── rating_handler.go            # CRUD criteria; get aggregates
│   │           │   ├── reputation_handler.go        # GET score/history
│   │           │   ├── badge_handler.go             # CRUD badges; list
│   │           │   ├── user_badge_handler.go        # 🆕 list badges for user
│   │           # ---- DOUBLE-BLIND & REMINDERS ----
│   │           │   ├── double_blind_handler.go      # open window; submit; publish-if-ready
│   │           │   ├── reminder_handler.go          # schedule/send reminders
│   │           # ---- PRIVATE FEEDBACK ----
│   │           │   ├── feedback_handler.go          # submit/get private feedback
│   │           # ---- SAFETY & GOVERNANCE ----
│   │           │   ├── flag_handler.go              # flag/resolve/dismiss
│   │           │   ├── moderation_handler.go        # set state; list flags
│   │           │   ├── appeal_handler.go            # open/decide appeals
│   │           │   ├── evidence_handler.go          # attach/remove evidence files
│   │           │   ├── redaction_handler.go         # preview/apply redaction
│   │           │   ├── compliance_handler.go        # DSAR erasure/export requests
│   │           # ---- STATS & EXPOSURE ----
│   │           │   ├── stats_handler.go             # get aggregates
│   │           │   ├── featured_handler.go          # editorial features
│   │           # ---- MISC ----
│   │           │   └── health_handler.go            # /health, /ready, /live
│   │           # =========================
│   │           # 🧰 MIDDLEWARE
│   │           # =========================
│   │           ├── middleware/
│   │           │   ├── auth.go                      # JWT verify (pkg/auth)
│   │           │   ├── rbac.go                      # role checks
│   │           │   ├── cors.go                      # platform-shared/ginx
│   │           │   ├── rate_limit.go                # token bucket
│   │           │   ├── logging.go                   # structured logs
│   │           │   ├── error_handler.go             # error → HTTP
│   │           │   └── request_id.go                # X-Request-ID
│   │           # =========================
│   │           # 📨 RESPONSES
│   │           # =========================
│   │           ├── responses/
│   │           │   ├── success.go                   # platform-shared/httpx/response
│   │           │   ├── error.go                     # platform-shared/httpx/errors
│   │           │   └── pagination.go                # platform-shared/httpx/pagination
│   │           # =========================
│   │           # 🗺️ ROUTES (SECTIONED)
│   │           # =========================
│   │           ├── routes/
│   │           │   # ---- CORE REVIEWS ----
│   │           │   ├── review_routes.go             # /reviews/*
│   │           │   ├── draft_routes.go              # 🆕 /reviews/:id/draft/*
│   │           │   ├── eligibility_routes.go        # 🆕 /eligibility/*
│   │           │   # ---- RATING/REPUTATION/BADGES ----
│   │           │   ├── rating_routes.go             # /ratings/*
│   │           │   ├── reputation_routes.go         # /reputation/*
│   │           │   ├── badge_routes.go              # /badges/*
│   │           │   ├── user_badge_routes.go         # 🆕 /user-badges/*
│   │           │   # ---- DOUBLE-BLIND & REMINDERS ----
│   │           │   ├── double_blind_routes.go       # /double-blind/*
│   │           │   ├── reminder_routes.go           # /reminders/*
│   │           │   # ---- PRIVATE FEEDBACK ----
│   │           │   ├── feedback_routes.go           # /feedback/*
│   │           │   # ---- SAFETY & GOVERNANCE ----
│   │           │   ├── flag_routes.go               # /flags/*
│   │           │   ├── moderation_routes.go         # /moderation/*
│   │           │   ├── appeal_routes.go             # /appeals/*
│   │           │   ├── evidence_routes.go           # /evidence/*
│   │           │   ├── redaction_routes.go          # /redactions/*
│   │           │   ├── compliance_routes.go         # /compliance/*
│   │           │   # ---- STATS & EXPOSURE ----
│   │           │   ├── stats_routes.go              # /stats/*
│   │           │   └── featured_routes.go           # /featured/*
│   │           └── router.go                        # Gin engine + common middleware wiring
│   │
│   └── (end internal)
│
├── config/
│   ├── default.yaml                               # Defaults (ports, DSN templates, kafka/redis hosts)
│   ├── dev.yaml                                   # Dev overrides
│   └── prod.yaml                                  # Prod overrides (timeouts, quotas)
│
├── dapr/
│   ├── local/
│   │   ├── pubsub.yaml                            # Kafka pub/sub component
│   │   └── statestore.yaml                        # State store component
│   └── k8s/
│       ├── pubsub.yaml                            # Scopes: ["reviews-be"]
│       ├── statestore.yaml                        # Scopes
│       └── secrets.yaml                           # Secret store
│
├── pkg/
│   # =============================
│   # 🧰 LOCAL UTILITIES
│   # =============================
│   ├── errors/
│   │   ├── errors.go                              # Service-specific errors
│   │   └── codes.go                               # Error codes (e.g., REVIEW_NOT_FOUND)
│   ├── utils/
│   │   ├── validator.go                           # Reusable validations
│   │   ├── sanitizer.go                           # HTML sanitization; guard XSS
│   │   └── sentiment_analyzer.go                  # Optional sentiment signal for moderation
│   └── constants/
│       ├── badges.go                              # Badge ids/types/levels
│       └── README.md                              # Events/topics come from contracts/events
│
├── seeds/
│   ├── badges.sql                                 # Seed default badges & levels
│   ├── rating_criteria.sql                        # 🆕 Seed criteria & versions per category
│   └── top_rated_rules.sql                        # 🆕 Seed reputation thresholds
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
│   ├── seed-criteria.sh                           # 🆕 load criteria versions
│   └── recalculate-reputation.sh                  # Rebuild reputation for all users
│
├── tests/
│   # =============================
│   # ✅ TEST SUITES
│   # =============================
│   ├── unit/                                      # domain & services (eligibility, double_blind, rating, reputation, moderation)
│   ├── integration/                               # handlers + repositories (Postgres/Redis)
│   └── e2e/                                       # contract end → blind window → publish → search-be emit
│
├── docs/
│   ├── README.md                                  # Overview
│   ├── api.md                                     # HTTP API
│   ├── events.md                                  # 📝 review.*, badge.*, reputation.*, compliance.*
│   ├── rating-system.md                           # Criteria, weights, versioning
│   ├── reputation-algorithm.md                    # Bayesian + decay details
│   ├── badge-system.md                            # Badges, levels, rules
│   ├── double-blind.md                            # Windows & publish policy
│   ├── moderation.md                               # Abuse controls & flows
│   ├── compliance.md                               # DSAR/erasure/export
│   ├── MIGRATIONS.md                               # History of schema changes
│   ├── SCHEMA.md                                   # Database schema snapshot
│   └── RUNBOOK.md                                  # Ops (DLQ, replayer, rebuilds)
│
├── .github/
│   └── workflows/
│       ├── ci.yml                                 # Lint, tests, docker build
│       └── cd.yml                                 # Deploy to K8s (staging/prod)
│
├── go.mod                                         # 📝 imports pkg/auth, platform-shared, contracts/events
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
│   # =============================
│   # 🚀 EXECUTABLES
│   # =============================
│   ├── api/
│   │   └── main.go                             # HTTP API bootstrap (Gin + Dapr). Loads config, mounts /v1 routes, health, metrics.
│   └── worker/                                 # 🆕 Background worker process (optional but handy)
│       └── main.go                             # 🆕 Cron-like loops: renewals, dunning retries, allowance rollovers, invoice exports.
│
├── internal/
│   # =============================
│   # 🔧 CONFIGURATION (LOAD FIRST)
│   # =============================
│   ├── config/
│   │   ├── schema.go                           # App, Server, Postgres, Kafka, Redis, Payments, Seats, Dunning.
│   │   ├── loader.go                           # Viper: flags → env → file → defaults; validates; logs effective config.
│   │   └── docs/
│   │       └── CONFIGURATION.md                # ENV keys, defaults, examples (local & k8s).
│   │
│   # =============================
│   # 🏛️ DOMAIN LAYER (DDD ENTITIES & RULES)
│   # =============================
│   ├── domain/
│   │   # -------- CATALOG: PLANS --------
│   │   ├── plan/
│   │   │   ├── entity.go                       # Plan{id, code, name, active, created_at}.
│   │   │   ├── features.go                     # feature_key → value (bool/number/cap).
│   │   │   ├── pricing.go                      # Period (monthly/yearly), base_price, currency (ISO-4217).
│   │   │   ├── limits.go                       # Declarative per-plan numeric caps (posts/day, invites, messages_to_non_hires).
│   │   │   ├── errors.go                       # PlanNotFound, InvalidFeature, LimitExceeded.
│   │   │   ├── repository.go                   # PlanRepository interface (GetByCode, ListActive, Upsert, Archive).
│   │   │   └── events.go                       # plan.created/updated/archived.v1
│   │   ├── plan_version/                       # 🆕 Immutable plan versioning for repricing/audit
│   │   │   ├── entity.go                       # PlanVersion{plan_id, version, features, pricing, active_from, notes}.
│   │   │   ├── migration_rule.go               # Optional auto-migration policy (opt-in/out windows).
│   │   │   ├── errors.go                       # VersionNotFound, MigrationBlocked.
│   │   │   ├── repository.go                   # PlanVersionRepository.
│   │   │   └── events.go                       # plan.version.created/activated/deprecated.v1
│   │
│   │   # -------- SUBSCRIPTIONS & CHANGE REQUESTS --------
│   │   ├── subscription/
│   │   │   ├── entity.go                       # Subscription{user_id, plan_id, status, period_start/end, cancel_at_period_end}.
│   │   │   ├── billing_cycle.go                # Compute next renewal, remaining days, grace windows.
│   │   │   ├── enums.go                        # Status{Active, PastDue, Paused, Canceled}; Type{Recurring, Trial}.
│   │   │   ├── auto_renewal.go                 # Enable/disable auto-renew + invariants (idempotent).
│   │   │   ├── errors.go                       # AlreadyActive, PastDue, Canceled.
│   │   │   ├── repository.go                   # SubscriptionRepository (Get, Save, FindDueForRenewal).
│   │   │   └── events.go                       # subscription.created/changed/paused/canceled/renewed.v1
│   │   ├── change_request/                     # 🆕 Schedule plan change (now/later) with proration policy
│   │   │   ├── entity.go                       # ChangeRequest{sub_id, new_plan_id, effective_at, policy}.
│   │   │   ├── proration_policy.go             # None | Immediate | CreditNote.
│   │   │   ├── errors.go                       # ChangeConflict, InvalidEffectiveDate.
│   │   │   ├── repository.go                   # ChangeRequestRepository.
│   │   │   └── events.go                       # subscription.change.scheduled/applied/canceled.v1
│   │
│   │   # -------- ENTITLEMENTS / GRANTS / USAGE / ALLOWANCE --------
│   │   ├── entitlement/
│   │   │   ├── entity.go                       # Entitlement{user_id, feature_key, allowed, scope}.
│   │   │   ├── rules.go                        # Merge order: plan < addon < promo < ad-hoc grant.
│   │   │   ├── errors.go                       # FeatureDenied, NotInPlan.
│   │   │   ├── repository.go                   # EntitlementRepository.
│   │   │   └── events.go                       # entitlement.feature.enabled/disabled.v1
│   │   ├── entitlement_grant/                  # 🆕 Ad-hoc feature grants (campaign/support gestures)
│   │   │   ├── entity.go                       # Grant{user_id, feature_key, qty?, expires_at, reason}.
│   │   │   ├── scope.go                        # Feature-level vs meter-level grants.
│   │   │   ├── errors.go                       # GrantExhausted, GrantExpired.
│   │   │   ├── repository.go                   # EntitlementGrantRepository.
│   │   │   └── events.go                       # entitlement.grant.issued/consumed/expired.v1
│   │   ├── usage/
│   │   │   ├── entity.go                       # UsageCounter{user_id, feature_key, period_key, value}.
│   │   │   ├── quota.go                        # Declarative caps & soft/hard logic.
│   │   │   ├── limit.go                        # Static per-plan limits snapshot.
│   │   │   ├── meter.go                        # 🆕 Meters: messages_to_non_hires, boosts, invites.
│   │   │   ├── errors.go                       # LimitReached, CounterNotFound.
│   │   │   ├── repository.go                   # UsageRepository.
│   │   │   └── events.go                       # usage.meter.incremented / usage.limit.reached.v1
│   │   ├── allowance/                          # 🆕 Rolling monthly buckets (grant + carryover)
│   │   │   ├── bucket.go                       # AllowanceBucket{feature_key, period, granted, carried_over, consumed}.
│   │   │   ├── rollover_rule.go                # Carryover caps & months.
│   │   │   ├── errors.go                       # BucketNotFound, RolloverNotAllowed.
│   │   │   ├── repository.go                   # AllowanceRepository.
│   │   │   └── events.go                       # allowance.granted/rolled_over/reset.v1
│   │
│   │   # -------- CONNECTS / SEATS / ADDONS / PROMOS / TRIALS --------
│   │   ├── connect/
│   │   │   ├── entity.go                       # User balance + ledger summary.
│   │   │   ├── package.go                      # Pack definitions (qty, price).
│   │   │   ├── transaction.go                  # +/− ledger items (idempotency keys).
│   │   │   ├── balance.go                      # Derived balance calculation.
│   │   │   ├── expiry.go                       # 🆕 Effective & expiry windows; rollover rules.
│   │   │   ├── grant.go                        # 🆕 Promo/connect grants.
│   │   │   ├── errors.go                       # InsufficientBalance, Expired, RolloverNotAllowed.
│   │   │   ├── repository.go                   # ConnectRepository.
│   │   │   └── events.go                       # connects.purchased/used/expired/granted.v1
│   │   ├── seat_billing/
│   │   │   ├── entity.go                       # Seats per subscription; assigned & cap.
│   │   │   ├── overage.go                      # Overage math (per-seat above cap).
│   │   │   ├── proration.go                    # Mid-cycle seat changes.
│   │   │   ├── invoice_export.go               # Exportable invoice lines (financial-be).
│   │   │   ├── errors.go                       # Seat billing errors.
│   │   │   ├── repository.go                   # SeatBillingRepository.
│   │   │   └── events.go                       # seat.overage.incurred / billing.proration.applied.v1
│   │   ├── addon/
│   │   │   ├── entity.go                       # Feature-pack SKU.
│   │   │   ├── errors.go                       # IncompatibleAddon.
│   │   │   ├── repository.go                   # AddonRepository.
│   │   │   └── events.go                       # addon.added/removed/updated.v1
│   │   ├── promotion/
│   │   │   ├── entity.go                       # Promo code with windows & usage caps.
│   │   │   ├── discount.go                     # Calc helpers: percent/fixed.
│   │   │   ├── usage_limit.go                  # Per-code & per-user limits.
│   │   │   ├── errors.go                       # InvalidCode, Exhausted, Ineligible.
│   │   │   ├── repository.go                   # PromotionRepository.
│   │   │   └── events.go                       # promo.created/applied/revoked/exhausted.v1
│   │   ├── trial/
│   │   │   ├── entity.go                       # Trial state & source.
│   │   │   ├── eligibility.go                  # Simple rule checks; no external calls.
│   │   │   ├── errors.go                       # NotEligible, AlreadyTrialed.
│   │   │   ├── repository.go                   # TrialRepository.
│   │   │   └── events.go                       # trial.started/ended/eligibility.updated.v1
│   │
│   │   # -------- BILLING CORE: INVOICE / PAYMENT / CREDIT / TAX / PROFILE --------
│   │   ├── invoice/                            # 🆕 Vendor-agnostic invoice schema
│   │   │   ├── entity.go                       # Invoice{status: draft|issued|paid|voided, totals, currency}.
│   │   │   ├── line_item.go                    # Lines for plan/addon/seats/connects/credit.
│   │   │   ├── tax.go                          # Simple per-line & total tax fields.
│   │   │   ├── adjustment.go                   # Proration & credits (adjustment lines).
│   │   │   ├── errors.go                       # InvoiceNotFound, InvalidState.
│   │   │   ├── repository.go                   # InvoiceRepository.
│   │   │   └── events.go                       # invoice.created/issued/paid/voided.v1
│   │   ├── payment/                            # 🆕 Payment intents & attempts
│   │   │   ├── intent.go                       # PaymentIntent{invoice_id, amount, status, retry_after}.
│   │   │   ├── attempt.go                      # Attempt results (success/failure, error_code, gateway_ref).
│   │   │   ├── method_hint.go                  # Non-PII hints (brand, last4, exp).
│   │   │   ├── errors.go                       # AlreadyCaptured, RetryWindowClosed.
│   │   │   ├── repository.go                   # PaymentRepository.
│   │   │   └── events.go                       # payment.intent.created/attempted/succeeded/failed.v1
│   │   ├── credit_note/                        # 🆕 Refund/credit records
│   │   │   ├── entity.go                       # CreditNote{amount, reason, status, remaining}.
│   │   │   ├── allocation.go                   # Apply credit to invoice or lines.
│   │   │   ├── errors.go                       # CreditExceedsBalance, AlreadyApplied.
│   │   │   ├── repository.go                   # CreditNoteRepository.
│   │   │   └── events.go                       # credit_note.issued/applied/voided.v1
│   │   ├── tax_class/                          # 🆕 Basic tax classification (no external service)
│   │   │   ├── entity.go                       # TaxClass{code, description}.
│   │   │   ├── binding.go                      # Binding{subject_kind, subject_id, class_code}.
│   │   │   ├── errors.go                       # TaxClassNotFound, BindingConflict.
│   │   │   ├── repository.go                   # TaxClassRepository.
│   │   │   └── events.go                       # tax_class.created/updated/bound.v1
│   │   ├── billing_profile/                    # 🆕 “Invoice to” identity
│   │   │   ├── entity.go                       # name, address, country, vat_id (strings, local format checks).
│   │   │   ├── validation.go                   # VAT format & address sanity checks (offline).
│   │   │   ├── errors.go                       # ProfileNotFound, InvalidVATFormat.
│   │   │   ├── repository.go                   # BillingProfileRepository.
│   │   │   └── events.go                       # billing_profile.created/updated.v1
│   │
│   │   # -------- DUNNING / HISTORY / TOGGLES --------
│   │   ├── dunning/
│   │   │   ├── case.go                          # DunningCase{stage, next_action_at, last_error}.
│   │   │   ├── schedule.go                      # Retry cadence/backoff policy.
│   │   │   ├── outcome.go                       # Resolved reason (paid/canceled/writeoff).
│   │   │   ├── errors.go                        # DunningCaseNotFound, InvalidTransition.
│   │   │   ├── repository.go                    # DunningRepository.
│   │   │   └── events.go                        # dunning.case.opened/advanced/resolved.v1
│   │   ├── billing_history/
│   │   │   ├── entity.go                        # Immutable audit snapshots (invoice/payment events).
│   │   │   ├── errors.go                        # HistoryNotFound.
│   │   │   ├── repository.go                    # BillingHistoryRepository.
│   │   │   └── events.go                        # billing.invoice.generated/payment.applied/credit.issued.v1
│   │   ├── feature_toggle/
│   │   │   ├── entity.go                        # On/off flags per plan for ops safety.
│   │   │   ├── errors.go                        # ToggleNotFound.
│   │   │   ├── repository.go                    # FeatureToggleRepository.
│   │   │   └── events.go                        # admin.feature_flag.updated / feature.toggle.enabled/disabled.v1
│   │   └── outbox/
│   │       ├── entity.go                        # ❌ REMOVED → use platform-shared/outbox/entity.go
│   │       └── repository.go                    # ❌ REMOVED → use platform-shared/outbox/repository.go
│   │
│   # =============================
│   # 📋 APPLICATION LAYER (USE CASES & ORCHESTRATION)
│   # =============================
│   ├── application/
│   │   # -------- EVENT CONSUMERS (INBOX) --------
│   │   ├── eventhandler/
│   │   │   ├── financial_handler.go             # payment.processed/failed → activate/renew/pause; dunning transitions.
│   │   │   ├── proposal_handler.go              # proposal.submitted → consume connects; enforce usage caps.
│   │   │   ├── job_handler.go                   # job.posted → posting limits; may consume connects.
│   │   │   └── admin_flags_handler.go           # admin.feature_flag/threshold/experiment.updated → refresh gates.
│   │
│   │   # -------- PLANS --------
│   │   ├── plan/
│   │   │   ├── service.go                       # CRUD with validations; emits plan.* events.
│   │   │   ├── commands.go                      # CreatePlan, UpdatePlan, ArchivePlan.
│   │   │   ├── queries.go                       # GetPlan, ListPlans.
│   │   │   ├── dto.go                           # PlanDTO, FeatureDTO, LimitDTO.
│   │   │   ├── mapper.go                        # Domain ↔ DTO mapping.
│   │   │   └── validators.go                    # Feature/limit/pricing guards.
│   │   ├── plan_version/                         # 🆕
│   │   │   ├── service.go                        # Create/Activate versions; apply migration rules.
│   │   │   ├── commands.go                       # CreatePlanVersion, ActivatePlanVersion.
│   │   │   ├── queries.go                        # GetVersion, ListVersions.
│   │   │   ├── dto.go                             # PlanVersionDTO, MigrationRuleDTO.
│   │   │   ├── mapper.go                          # Version ↔ DTO.
│   │   │   └── validators.go                      # Version/migration invariants.
│   │
│   │   # -------- SUBSCRIPTIONS --------
│   │   ├── subscription/
│   │   │   ├── service.go                       # Subscribe/Cancel/Pause/Resume/Renew orchestrations.
│   │   │   ├── lifecycle_manager.go             # Renewal pipeline (invoice→payment→entitlements), idempotent.
│   │   │   ├── renewal_manager.go               # Due-subscriptions fetch; backoff & jitter.
│   │   │   ├── commands.go                      # Subscribe, Cancel, ChangePlan, Pause, Resume.
│   │   │   ├── queries.go                       # GetSubscription, ListSubscriptions.
│   │   │   ├── dto.go                            # SubscriptionDTO, ChangePlanDTO.
│   │   │   ├── mapper.go                         # Domain ↔ DTO.
│   │   │   └── validators.go                     # State transitions, proration inputs.
│   │   ├── change_request/                       # 🆕
│   │   │   ├── service.go                        # Schedule/apply/cancel changes.
│   │   │   ├── commands.go                       # ScheduleChange, ApplyNow, CancelChange.
│   │   │   ├── queries.go                        # GetPendingChanges.
│   │   │   ├── dto.go                             # ChangeRequestDTO.
│   │   │   ├── mapper.go                          # Domain ↔ DTO.
│   │   │   └── validators.go                      # Effective_at, conflict checks.
│   │
│   │   # -------- ENTITLEMENTS / GRANTS / USAGE / ALLOWANCE --------
│   │   ├── entitlement/
│   │   │   ├── service.go                       # Resolve effective gates; CheckFeature.
│   │   │   ├── commands.go                      # GrantFeature, RevokeFeature (admin/system).
│   │   │   ├── queries.go                       # GetEntitlements, CheckFeature.
│   │   │   ├── dto.go                            # EntitlementDTO, CheckResultDTO.
│   │   │   ├── mapper.go                         # Domain ↔ DTO.
│   │   │   └── validators.go                     # Feature key validity; scope checks.
│   │   ├── entitlement_grant/                    # 🆕
│   │   │   ├── service.go                        # Issue/expire grants; consume on use.
│   │   │   ├── commands.go                       # IssueGrant, ExpireGrant.
│   │   │   ├── queries.go                        # ListGrants, GetGrant.
│   │   │   ├── dto.go                             # GrantDTO.
│   │   │   ├── mapper.go                          # Domain ↔ DTO.
│   │   │   └── validators.go                      # Amounts, expiries, scopes.
│   │   ├── usage/
│   │   │   ├── service.go                       # Increment meters & enforce caps (idempotency token).
│   │   │   ├── tracker.go                        # Records usage with dedupe keys (counter+token).
│   │   │   ├── quota_checker.go                   # Soft/hard cap evaluation.
│   │   │   ├── limiter.go                         # Hard-stop 4xx decisions for gated features.
│   │   │   ├── commands.go                        # IncrementUsage, ResetUsage.
│   │   │   ├── queries.go                         # GetUsage, GetLimits.
│   │   │   ├── dto.go                              # UsageDTO, LimitsDTO.
│   │   │   ├── mapper.go                           # Domain ↔ DTO.
│   │   │   └── validators.go                       # Counter existence, range checks.
│   │   ├── allowance/                              # 🆕
│   │   │   ├── service.go                          # Grant/rollover/read buckets.
│   │   │   ├── commands.go                         # GrantAllowance, RolloverAllowance, ResetAllowance.
│   │   │   ├── queries.go                          # GetAllowance, ListAllowances.
│   │   │   ├── dto.go                               # AllowanceDTO, RolloverDTO.
│   │   │   ├── mapper.go                            # Domain ↔ DTO.
│   │   │   └── validators.go                        # Rollover caps & periods.
│   │
│   │   # -------- CONNECTS / SEATS / ADDONS / PROMOS / TRIALS --------
│   │   ├── connect/
│   │   │   ├── service.go                       # Purchase/Use/Transfer/Refund; emits connects.*.
│   │   │   ├── calculator.go                    # Cost & rollover helpers.
│   │   │   ├── commands.go                      # PurchaseConnects, UseConnects, RefundConnects.
│   │   │   ├── queries.go                       # GetBalance, GetHistory.
│   │   │   ├── dto.go                            # ConnectBalanceDTO, ConnectTxnDTO.
│   │   │   ├── mapper.go                         # Domain ↔ DTO.
│   │   │   └── validators.go                     # Amounts, expiries, idempotency keys.
│   │   ├── seat_billing/
│   │   │   ├── service.go                       # Assign/Release seats; compute overages; export lines.
│   │   │   ├── commands.go                      # AssignSeat, ReleaseSeat, SetSeatCap, RecalculateOverages.
│   │   │   ├── queries.go                       # GetSeatSummary, GetOverages.
│   │   │   ├── dto.go                            # SeatSummaryDTO, OverageDTO.
│   │   │   ├── mapper.go                         # Domain ↔ DTO.
│   │   │   └── validators.go                     # Seat counts/policies.
│   │   ├── addon/
│   │   │   ├── service.go                       # Add/remove addons on subscriptions.
│   │   │   ├── commands.go                      # AddAddon, RemoveAddon.
│   │   │   ├── queries.go                       # GetAddons, GetAddon.
│   │   │   ├── dto.go                            # AddonDTO.
│   │   │   ├── mapper.go                         # Domain ↔ DTO.
│   │   │   └── validators.go                     # Compatibility checks.
│   │   ├── promotion/
│   │   │   ├── service.go                       # Create/apply/ revoke promo codes.
│   │   │   ├── validator.go                     # Legacy internal checks (kept).
│   │   │   ├── validators.go                    # Alias for naming consistency.
│   │   │   ├── commands.go                      # CreatePromo, ApplyPromo, RevokePromo.
│   │   │   ├── queries.go                       # GetPromo, ListPromos.
│   │   │   ├── dto.go                            # PromoDTO, ApplyResultDTO.
│   │   │   └── mapper.go                         # Domain ↔ DTO.
│   │   ├── trial/
│   │   │   ├── service.go                       # Start/End trials.
│   │   │   ├── eligibility_checker.go           # Lightweight rules (no vendor calls).
│   │   │   ├── commands.go                      # StartTrial, EndTrial.
│   │   │   ├── queries.go                       # GetTrial, IsEligible.
│   │   │   ├── dto.go                            # TrialDTO, EligibilityDTO.
│   │   │   ├── mapper.go                         # Domain ↔ DTO.
│   │   │   └── validators.go                     # Eligibility/rules checks.
│   │
│   │   # -------- BILLING ORCHESTRATION --------
│   │   ├── billing/
│   │   │   ├── service.go                       # Orchestrates invoice→payment→export; idempotent via keys.
│   │   │   ├── invoice_generator.go             # Build invoice + lines (plan/addon/seats/connects/proration).
│   │   │   ├── payment_processor.go             # Create payment intents; record attempts; capture success.
│   │   │   ├── exporter.go                      # Export issued invoices to financial-be (UPSERT downstream).
│   │   │   ├── commands.go                      # GenerateInvoice, ExportInvoice, CapturePayment.
│   │   │   ├── queries.go                       # GetInvoice, ListInvoices.
│   │   │   └── validators.go                    # Seat counts, proration math, export invariants.
│   │   ├── invoice/                              # 🆕 Focused invoice helpers
│   │   │   ├── service.go                        # Issue/Void/Get invoice; state machine guards.
│   │   │   ├── commands.go                       # IssueInvoice, VoidInvoice.
│   │   │   ├── queries.go                        # GetInvoiceByID, ListInvoicesByUser.
│   │   │   ├── dto.go                             # InvoiceDTO, LineItemDTO.
│   │   │   ├── mapper.go                          # Domain ↔ DTO.
│   │   │   └── validators.go                      # Valid states, totals.
│   │   ├── payment/                              # 🆕 Focused payment helpers
│   │   │   ├── service.go                        # Create intents, record attempts, finalize.
│   │   │   ├── commands.go                       # CreateIntent, RecordAttempt, FinalizeIntent.
│   │   │   ├── queries.go                        # GetIntent, ListAttempts.
│   │   │   ├── dto.go                             # PaymentIntentDTO, AttemptDTO.
│   │   │   ├── mapper.go                          # Domain ↔ DTO.
│   │   │   └── validators.go                      # Amount/currency/status checks.
│   │   ├── credit_note/                          # 🆕 Refund/credit flows
│   │   │   ├── service.go                        # Issue/apply credits; prevent over-allocation.
│   │   │   ├── commands.go                       # IssueCredit, ApplyCredit, VoidCredit.
│   │   │   ├── queries.go                        # GetCreditNote, ListCredits.
│   │   │   ├── dto.go                             # CreditNoteDTO, AllocationDTO.
│   │   │   ├── mapper.go                          # Domain ↔ DTO.
│   │   │   └── validators.go                      # Amount bounds, remaining balance checks.
│   │   ├── tax_class/                            # 🆕 Taxability mgmt
│   │   │   ├── service.go                        # CRUD tax classes; bind subjects.
│   │   │   ├── commands.go                       # CreateTaxClass, BindTaxClass.
│   │   │   ├── queries.go                        # GetTaxClass, ListTaxClasses.
│   │   │   ├── dto.go                             # TaxClassDTO, BindingDTO.
│   │   │   ├── mapper.go                          # Domain ↔ DTO.
│   │   │   └── validators.go                      # Binding/exists checks.
│   │   └── billing_profile/                       # 🆕 “Invoice to” identity
│   │       ├── service.go                          # CRUD profiles; simple VAT format checks.
│   │       ├── commands.go                         # UpsertBillingProfile.
│   │       ├── queries.go                          # GetBillingProfile.
│   │       ├── dto.go                               # BillingProfileDTO.
│   │       ├── mapper.go                            # Domain ↔ DTO.
│   │       └── validators.go                        # Address/VAT pattern checks.
│   │
│   │   # -------- DUNNING --------
│   │   └── dunning/
│   │       ├── service.go                         # Open/Advance/Resolve cases.
│   │       ├── workflow.go                        # Decide next action time; jitter/backoff.
│   │       ├── commands.go                         # OpenCase, AdvanceCase, ResolveCase.
│   │       ├── queries.go                          # GetCase, ListCases.
│   │       ├── dto.go                               # DunningCaseDTO.
│   │       ├── mapper.go                            # Domain ↔ DTO.
│   │       └── validators.go                        # Stage transitions.
│   │
│   # =============================
│   # 🔌 INFRASTRUCTURE
│   # =============================
│   ├── infrastructure/
│   │   # -------- 🗄️ PERSISTENCE (POSTGRES) --------
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       # =========================
│   │   │       # 🧱 COMMON (DB BOOTSTRAP)
│   │   │       # =========================
│   │   │       ├── connection.go                  # DSN, pooling, timeouts.
│   │   │       ├── transaction.go                 # WithTx helpers; savepoints for nested ops.
│   │   │       ├── migrations.go                  # Auto-migrate + ordered version steps.
│   │   │       ├── version.go                     # schema_version helpers.
│   │   │       ├── safety.go                      # Env/disk sanity checks pre-migration.
│   │   │       # =========================
│   │   │       # 📚 CATALOG (PLANS)
│   │   │       # =========================
│   │   │       ├── plan_repository.go
│   │   │       ├── plan_version_repository.go      # 🆕
│   │   │       # =========================
│   │   │       # 🔁 SUBSCRIPTIONS & CHANGES
│   │   │       # =========================
│   │   │       ├── subscription_repository.go
│   │   │       ├── change_request_repository.go    # 🆕
│   │   │       # =========================
│   │   │       # 🎫 ENTITLEMENTS • USAGE • ALLOWANCE
│   │   │       # =========================
│   │   │       ├── entitlement_repository.go
│   │   │       ├── entitlement_grant_repository.go # 🆕
│   │   │       ├── usage_repository.go
│   │   │       ├── allowance_repository.go         # 🆕
│   │   │       # =========================
│   │   │       # 💼 COMMERCIAL ADD-ONS
│   │   │       # (Connects, Seats, Addons, Promotions, Trials)
│   │   │       # =========================
│   │   │       ├── connect_repository.go
│   │   │       ├── seat_billing_repository.go
│   │   │       ├── addon_repository.go
│   │   │       ├── promotion_repository.go
│   │   │       ├── trial_repository.go
│   │   │       # =========================
│   │   │       # 🧾 BILLING SUITE
│   │   │       # (Invoices, Payments, Credits, Tax, Billing Profiles)
│   │   │       # =========================
│   │   │       ├── invoice_repository.go           # 🆕
│   │   │       ├── payment_repository.go           # 🆕
│   │   │       ├── credit_note_repository.go       # 🆕
│   │   │       ├── tax_class_repository.go         # 🆕
│   │   │       ├── billing_profile_repository.go   # 🆕
│   │   │       # =========================
│   │   │       # 🗃️ HISTORY & TOGGLES
│   │   │       # =========================
│   │   │       ├── billing_history_repository.go
│   │   │       ├── feature_toggle_repository.go
│   │   │       # =========================
│   │   │       # 📨 OUTBOX (REMOVED)
│   │   │       # =========================
│   │   │       └── outbox_repository.go            # ❌ REMOVED → platform-shared/outbox/postgres
│   │   # -------- ⚡ CACHE (REDIS) --------
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go                  # Pooled client; exponential backoff on boot.
│   │   │       ├── subscription_cache.go          # Current plan/status hot path.
│   │   │       ├── plan_cache.go                  # Plan & features snapshot.
│   │   │       ├── connect_cache.go               # Balance & TTL for UX.
│   │   │       ├── entitlement_cache.go           # 🆕 Effective gates by user_id.
│   │   │       ├── feature_toggle_cache.go        # Ops toggles.
│   │   │       ├── invoice_cache.go               # 🆕 Payment webhook speedup.
│   │   │       ├── dunning_cache.go               # 🆕 Stage/next_action hints.
│   │   │       ├── allowance_cache.go             # 🆕 Current-period buckets.
│   │   │       └── billing_profile_cache.go       # 🆕 Quick profile/VAT access.
│   │   # -------- 📬 MESSAGING (KAFKA) --------
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go                    # Uses platform-shared/inbox (dedupe, offsets).
│   │   │       ├── producer.go                    # Uses platform-shared/outbox (idempotent publish).
│   │   │       ├── topics.go                      # Import constants from contracts/events.
│   │   │       └── scram.go                       # SASL/SCRAM-256 auth.
│   │   # -------- ⏰ SCHEDULER --------
│   │   ├── scheduler/
│   │   │   ├── cron.go                            # Minimal ticker + jitter harness.
│   │   │   └── jobs.go                            # Renewals, dunning retries, allowance rollover, grant expiry.
│   │   # -------- 💳 FINANCIAL-BE CLIENTS --------
│   │   ├── payment/
│   │   │   └── client.go                          # REST/gRPC to financial-be: intents, capture.
│   │   ├── billing_export/
│   │   │   ├── client.go                          # Export issued invoices (idempotent).
│   │   │   └── mapper.go                          # Local invoice/lines → financial-be DTO.
│   │   # -------- 📦 OUTBOX (REMOVED) --------
│   │   └── outbox/
│   │       ├── processor.go                       # ❌ REMOVED (use platform-shared/outbox/forwarder.go).
│   │       └── scheduler.go                       # ❌ REMOVED (use platform-shared/outbox/scheduler.go).
│
│   # =============================
│   # 🌐 HTTP INTERFACE — API v1
│   # =============================
│   ├── interfaces/
│   │   └── http/
│   │       └── v1/
│   │           # -------- 🧭 HANDLERS --------
│   │           ├── handlers/
│   │           │   # Catalog
│   │           │   ├── plan_handler.go              # Admin CRUD for plans.
│   │           │   ├── plan_version_handler.go      # 🆕 Manage plan versions.
│   │           │   # Subscriptions
│   │           │   ├── subscription_handler.go      # Subscribe/Cancel/Pause/Resume; read status.
│   │           │   ├── change_request_handler.go    # 🆕 Schedule/apply/cancel plan changes.
│   │           │   # Entitlements & usage
│   │           │   ├── entitlement_handler.go       # 🆕 Check/List effective features.
│   │           │   ├── entitlement_grant_handler.go # 🆕 Issue/expire grants (admin/system).
│   │           │   ├── usage_handler.go             # Increment/inspect usage counters.
│   │           │   ├── allowance_handler.go         # 🆕 Grant/rollover/read buckets.
│   │           │   # Connects / seats / addons / promos / trials
│   │           │   ├── connect_handler.go           # Purchase/use/refund; balance & history.
│   │           │   ├── seat_billing_handler.go      # 🆕 Seats assign/release & overages.
│   │           │   ├── addon_handler.go             # Add/remove addons.
│   │           │   ├── promotion_handler.go         # Create/apply/revoke promo codes.
│   │           │   ├── trial_handler.go             # Start/inspect trials.
│   │           │   # Billing suite
│   │           │   ├── billing_handler.go           # Orchestrated generate/export/capture.
│   │           │   ├── invoice_handler.go           # 🆕 Issue/Void/Get invoices.
│   │           │   ├── payment_handler.go           # 🆕 Create intents; record attempts.
│   │           │   ├── credit_note_handler.go       # 🆕 Issue/apply credits.
│   │           │   ├── tax_class_handler.go         # 🆕 CRUD tax classes & bindings.
│   │           │   ├── billing_profile_handler.go   # 🆕 CRUD invoice-to profiles.
│   │           │   # Ops & health
│   │           │   ├── dunning_handler.go           # 🆕 List/advance/resolve dunning cases.
│   │           │   └── health_handler.go            # /health, /ready, /live
│   │           #
│   │           # -------- 🗺️ ROUTES (SMALL GROUPS + SECTION HEADERS) --------
│   │           ├── routes/
│   │           │   # Catalog
│   │           │   ├── plan_routes.go               # /v1/plans/*
│   │           │   ├── plan_version_routes.go       # 🆕 /v1/plans/:id/versions/*
│   │           │   # Subscriptions
│   │           │   ├── subscription_routes.go       # /v1/subscriptions/*
│   │           │   ├── change_request_routes.go     # 🆕 /v1/subscriptions/:id/changes/*
│   │           │   # Entitlements & usage
│   │           │   ├── entitlement_routes.go        # 🆕 /v1/entitlements/*
│   │           │   ├── entitlement_grant_routes.go  # 🆕 /v1/grants/*
│   │           │   ├── usage_routes.go              # /v1/usage/*
│   │           │   ├── allowance_routes.go          # 🆕 /v1/allowances/*
│   │           │   # Connects / seats / addons / promos / trials
│   │           │   ├── connect_routes.go            # /v1/connects/*
│   │           │   ├── seat_billing_routes.go       # 🆕 /v1/seats/*
│   │           │   ├── addon_routes.go              # /v1/addons/*
│   │           │   ├── promotion_routes.go          # /v1/promotions/*
│   │           │   ├── trial_routes.go              # /v1/trials/*
│   │           │   # Billing suite
│   │           │   ├── billing_routes.go            # /v1/billing/*
│   │           │   ├── invoice_routes.go            # 🆕 /v1/invoices/*
│   │           │   ├── payment_routes.go            # 🆕 /v1/payments/*
│   │           │   ├── credit_note_routes.go        # 🆕 /v1/credit-notes/*
│   │           │   ├── tax_class_routes.go          # 🆕 /v1/tax-classes/*
│   │           │   └── billing_profile_routes.go    # 🆕 /v1/billing-profiles/*
│   │           #
│   │           # -------- 🧱 MIDDLEWARE & RESPONSES --------
│   │           ├── middleware/
│   │           │   ├── auth.go                      # JWT verification (pkg/auth) → userID, roles.
│   │           │   ├── rbac.go                      # Role checks (admin/system/user) at route level.
│   │           │   ├── cors.go                      # platform-shared/ginx CORS.
│   │           │   ├── rate_limit.go                # Token-bucket per user/IP.
│   │           │   ├── logging.go                   # Structured access logs (trace/span if present).
│   │           │   ├── error_handler.go             # JSON error mapping (stable codes).
│   │           │   ├── request_id.go                # X-Request-ID passthrough/generation.
│   │           │   └── feature_gate.go              # 403 early for gated endpoints (plan/feature).
│   │           ├── responses/
│   │           │   ├── success.go                   # platform-shared/httpx/response.go wrappers.
│   │           │   ├── error.go                     # platform-shared/httpx/errors.go wrappers.
│   │           │   └── pagination.go                # platform-shared/httpx/pagination.go helpers.
│   │           └── router.go                        # Mount /v1 groups per section; attach middlewares; health routes.
│
├── config/
│   ├── default.yaml                            # Local-safe defaults.
│   ├── dev.yaml                                # Dev overrides.
│   └── prod.yaml                               # Production tuning (timeouts, retries).
│
├── dapr/
│   ├── local/
│   │   ├── pubsub.yaml                         # Kafka pub/sub (scope: subscriptions-be).
│   │   └── statestore.yaml                     # State store for idempotency tokens (short TTL).
│   └── k8s/
│       ├── pubsub.yaml                         # Dapr component with scopes: ["subscriptions-be"].
│       ├── statestore.yaml                     # Redis/StateStore config.
│       └── secrets.yaml                        # Dapr secret store bindings.
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                           # Service-level error helpers (wrap domain → HTTP).
│   │   └── codes.go                            # Stable error codes (SUBS_*).
│   ├── logger/
│   │   └── logger.go                           # ❌ REMOVED — use platform-shared/logging.
│   ├── utils/
│   │   ├── validator.go                        # Small validation helpers.
│   │   ├── billing_calculator.go               # Period fractions & seat proration math.
│   │   └── proration.go                         # Charge split helpers across date ranges.
│   └── constants/
│       ├── events.go                           # ❌ REMOVED — use contracts/events.
│       ├── topics.go                           # ❌ REMOVED — use contracts/events.
│       ├── plans.go                            # Seed/test plan codes.
│       └── features.go                         # Canonical feature keys.
│
├── seeds/
│   ├── plans.sql                                # Base plans & limits.
│   ├── connect_packages.sql                     # Connect bundles.
│   └── promo_codes.sql                          # 🆕 Starter promos for QA.
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                      # Probes: /health, /ready, /live.
│       ├── service.yaml                         # ClusterIP service.
│       ├── configmap.yaml                       # Non-secret config.
│       ├── secrets.yaml                         # DSN/API keys.
│       ├── hpa.yaml                             # Scale by CPU/RAM.
│       ├── pdb.yaml                             # Disruption budget.
│       ├── cronjob-renewal.yaml                 # Nightly renewals/dunning/rollover.
│       └── servicemonitor.yaml                  # Prometheus scrape config.
│
├── scripts/
│   ├── setup-local.sh                           # Create DB, migrate, seed.
│   ├── get-secrets.sh                           # Pull local secrets placeholder.
│   ├── seed-plans.sh                            # Seed plans & features.
│   ├── seed-data.sh                             # Small playground dataset.
│   ├── process-renewals.sh                      # Kick renewal pipeline.
│   ├── process-dunning.sh                       # 🆕 Advance dunning for test users.
│   └── export-invoices.sh                       # 🆕 Re-export last N invoices to financial-be.
│
├── tests/
│   ├── unit/                                    # Domain & service tests (no I/O).
│   ├── integration/                             # Repo + HTTP handler tests (real DB/Redis).
│   └── e2e/                                     # Subscribe→renew→dunning happy-path flows.
│
├── docs/
│   ├── README.md                                # Service scope & boundaries.
│   ├── api.md                                   # Endpoint reference (paths, payloads, errors).
│   ├── events.md                                # Published/consumed events; envelope examples.
│   ├── subscription-plans.md                    # Plan architecture & versioning.
│   ├── connects-system.md                       # Connects lifecycle & rollover.
│   ├── billing-logic.md                         # Invoices, proration, credits, exports.
│   ├── feature-toggles.md                       # Operational flags & rollout playbook.
│   ├── dunning.md                               # Stages, cadence, comms expectations.
│   ├── invoices.md                              # Invoice/lines/taxes; state machine.
│   ├── taxes-and-tax-classes.md                 # Basic tax classes; no external tax svc.
│   ├── credits-and-refunds.md                   # Credit note policy & flows.
│   ├── MIGRATIONS.md                            # Migration history & guardrails.
│   ├── SCHEMA.md                                # ERD & key indices.
│   └── RUNBOOK.md                               # Ops tasks: renewals, dunning, exports, rollovers.
│
├── .github/
│   └── workflows/
│       ├── ci.yml                               # Lint, unit, integration (dockerized DB/Redis).
│       └── cd.yml                               # Build & deploy to k8s (image tag = git SHA).
│
├── go.mod
├── go.sum
├── .env.example                                 # Local env template; never commit real secrets.
├── Makefile                                     # make run|test|lint|migrate|seed
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md                                    # Quickstart, scope, links to docs.

```


---

## **📦1️⃣1️⃣ admin-be (NEW - COMPREHENSIVE**

```
apps/be/admin-be/
│
├── cmd/
│   └── api/
│       └── main.go                                  # Entry: init Gin (+/v1), Dapr, Postgres, Redis, Kafka; load config; wire platform-shared logging.
│
├── internal/
│   # =====================================================================
│   # 🏛️ DOMAIN LAYER — Aggregates, Value Objects, Repositories, Events
│   # =====================================================================
│   └── domain/
│       # ------------------------- CORE ADMIN --------------------------------
│       ├── admin_user/
│       │   ├── entity.go                           # AdminUser aggregate (id, email, roles, mfa, status).
│       │   ├── role.go                             # Role model (SuperAdmin/Moderator/Support) + invariants.
│       │   ├── permission.go                       # Permission bitset helpers (grant/revoke/check).
│       │   ├── activity_log.go                     # Staff action trail (who/what/when/ip/ua).
│       │   ├── errors.go                           # AdminUserNotFound, PermissionDenied, RoleConflict.
│       │   ├── repository.go                       # Persist/find admins; mutate roles/permissions; append audit.
│       │   └── events.go                           # admin.user.created/updated/role.changed/permissions.updated/login.logged.v1
│
│       # ------------------------- SUPPORT & CASEWORK -------------------------
│       ├── support_ticket/
│       │   ├── entity.go                           # Ticket (subject, category, priority, status, SLA clocks).
│       │   ├── priority.go                         # Priority enum + escalation thresholds.
│       │   ├── status.go                           # Status FSM (Open/InProgress/Resolved/Closed).
│       │   ├── category.go                         # Ticket categories (billing, KYC, abuse, etc.).
│       │   ├── assignment.go                       # Assignment VO (agent, queue, takeover rules).
│       │   ├── errors.go                           # TicketNotFound, InvalidStatus, SLAExceeded.
│       │   ├── repository.go                       # CRUD, assign, transition, SLA stamping.
│       │   └── events.go                           # admin.case.opened/updated/assigned/closed.v1
│       ├── ticket_message/
│       │   ├── entity.go                           # Ticket message (author, body, internal/public).
│       │   ├── attachment.go                       # Attachment (fileRef, checksum, virusStatus).
│       │   ├── errors.go                           # MessageNotFound, AttachmentTooLarge, NotAuthor.
│       │   ├── repository.go                       # Add/edit/delete messages; manage attachments.
│       │   └── events.go                           # ticket.note.added/edited/deleted/attachment.added.v1
│       ├── support_agent/
│       │   ├── entity.go                           # Agent profile (skills, queues, workload).
│       │   ├── availability.go                     # Presence state (Online/Busy/Offline) with timeouts.
│       │   ├── stats.go                            # Rolling KPIs (FRT, ART, CSAT, resolutions).
│       │   ├── errors.go                           # AgentNotFound, AgentUnavailable, Overallocated.
│       │   ├── repository.go                       # Upsert agent, set availability, update metrics.
│       │   └── events.go                           # support.agent.status.changed/assigned/workload.updated.v1
│
│       # ------------------------- CONTENT & KNOWLEDGE ------------------------
│       ├── canned_response/
│       │   ├── entity.go                           # Prewritten response (title, body, locale, tags).
│       │   ├── category.go                         # Response grouping (folders/scopes).
│       │   ├── errors.go                           # ResponseNotFound, DuplicateTitle.
│       │   ├── repository.go                       # CRUD responses/categories.
│       │   └── events.go                           # canned_response.created/updated/archived.v1
│       ├── knowledge_base/
│       │   ├── entity.go                           # KB article (title, body, status, versionId).
│       │   ├── category.go                         # KB categories (nested tree).
│       │   ├── tag.go                              # Article tags for search facets.
│       │   ├── version.go                          # Versioned content + diffs metadata.
│       │   ├── errors.go                           # ArticleNotFound, VersionConflict, NotPublished.
│       │   ├── repository.go                       # CRUD, version, publish/unpublish.
│       │   └── events.go                           # kb.article.created/updated/versioned/published/unpublished.v1
│       ├── faq/
│       │   ├── entity.go                           # FAQ (question/answer, locale, ordering).
│       │   ├── category.go                         # FAQ categories.
│       │   ├── errors.go                           # FAQNotFound, CategoryConflict.
│       │   ├── repository.go                       # CRUD FAQs/categories.
│       │   └── events.go                           # faq.created/updated/published/unpublished.v1
│
│       # ------------------------- SAFETY & MODERATION ------------------------
│       ├── moderation_queue/
│       │   ├── entity.go                           # Queue item (contentRef, type, reason, state).
│       │   ├── content_type.go                     # Target types (Job/User/Review/Message/Asset).
│       │   ├── flag_reason.go                      # Canonical reasons + weights.
│       │   ├── action.go                           # Decision taken (hide/remove/ban/strike).
│       │   ├── errors.go                           # QueueItemNotFound, AlreadyActioned.
│       │   ├── repository.go                       # Enqueue, assign, decide, list.
│       │   └── events.go                           # admin.moderation.enqueued/assigned/state.changed.v1
│       ├── user_action/
│       │   ├── entity.go                           # Staff action on user (suspend/ban/verify/warn) + evidence.
│       │   ├── action_type.go                      # Action enum + semantics.
│       │   ├── reason.go                           # Structured reason codes.
│       │   ├── errors.go                           # ActionNotFound, AlreadyReversed.
│       │   ├── repository.go                       # Apply, reverse, enumerate actions.
│       │   └── events.go                           # admin.user_action.applied/reversed.v1
│       ├── content_action/
│       │   ├── entity.go                           # Action on content (remove/hide/approve/reject) + scope.
│       │   ├── action_type.go                      # Decision enums with constraints.
│       │   ├── errors.go                           # ContentActionNotFound, InvalidTransition.
│       │   ├── repository.go                       # Decide/restore/list content actions.
│       │   └── events.go                           # admin.content.actioned.v1
│
│       # ------------------------- DISPUTES & CASES ---------------------------
│       ├── dispute_resolution/
│       │   ├── entity.go                           # Dispute case (parties, contractRef, state timeline).
│       │   ├── evidence.go                         # Evidence refs & integrity info.
│       │   ├── decision.go                         # Outcomes & remedies awarded.
│       │   ├── errors.go                           # DisputeNotFound, InvalidState.
│       │   ├── repository.go                       # Open/update/decide/close; add evidence.
│       │   └── events.go                           # admin.case.opened/updated/closed/admin.dispute.decision.made.v1
│       ├── appeal/                                 # 🆕 User appeal flow
│       │   ├── entity.go                           # Appeal (target action, arguments, attachments, state).
│       │   ├── errors.go                           # AppealNotFound, InvalidOutcome.
│       │   ├── repository.go                       # File/decide/escalate/close appeals.
│       │   └── events.go                           # admin.appeal.opened/decided/escalated.v1
│       ├── case_linking/                           # 🆕 Cross-case relations
│       │   ├── entity.go                           # CaseLink (type: related/duplicate/blocking) between case IDs.
│       │   ├── errors.go                           # CaseLinkNotFound, InvalidLinkType.
│       │   ├── repository.go                       # Link/unlink/list case relationships.
│       │   └── events.go                           # admin.case.linked/unlinked.v1
│
│       # ------------------------- IDENTITY & VERIFICATION --------------------
│       ├── kyc_case/                               # 🆕 KYC for people/entities
│       │   ├── entity.go                           # KYCCase (applicant, doc set, status, notes).
│       │   ├── document.go                         # Document metadata (type, country, hash, reviewState).
│       │   ├── decision.go                         # VerificationDecision (approved/rejected/reopened).
│       │   ├── errors.go                           # KYCCaseNotFound, InvalidDocument.
│       │   ├── repository.go                       # Create/update/review KYC cases & docs.
│       │   └── events.go                           # admin.kyc.submitted/approved/rejected/reopened.v1
│       ├── business_verification/                  # 🆕 For agencies/companies
│       │   ├── entity.go                           # BusinessProfile (legalName, regNo, country, status).
│       │   ├── evidence.go                         # Evidence (docs/links) + checks.
│       │   ├── decision.go                         # Decision state transitions.
│       │   ├── errors.go                           # BusinessProfileNotFound, EvidenceInvalid.
│       │   ├── repository.go                       # Intake/review/decision persistence.
│       │   └── events.go                           # admin.biz_verification.requested/approved/rejected.v1
│       ├── sanctions_screening/                    # 🆕 Watchlist hits
│       │   ├── entity.go                           # ScreeningRun (sourceList, time, status).
│       │   ├── hit.go                               # Hit (entityRef, matchScore, list, disposition).
│       │   ├── disposition.go                      # Disposition (cleared/escalated/manual_review).
│       │   ├── errors.go                           # ScreeningRunNotFound, InvalidDisposition.
│       │   ├── repository.go                       # Store runs/hits; update dispositions.
│       │   └── events.go                           # admin.sanctions.hit.detected/cleared/escalated.v1
│
│       # ------------------------- LEGAL & PRIVACY ----------------------------
│       ├── legal_hold/                             # 🆕 Hold + eDiscovery
│       │   ├── entity.go                           # Hold (scope user/content/contract, placedBy, reason).
│       │   ├── release.go                          # Release record (who/when/why).
│       │   ├── export_job.go                       # Export job metadata for discovery.
│       │   ├── errors.go                           # HoldNotFound, AlreadyReleased.
│       │   ├── repository.go                       # Place/release holds; manage export jobs.
│       │   └── events.go                           # admin.legal_hold.placed/released, admin.ediscovery.export.created/completed.v1
│       ├── privacy_request/                        # 🆕 GDPR/CCPA DSAR
│       │   ├── entity.go                           # Request (type: erasure/access; subject; scope; status).
│       │   ├── evidence.go                         # Identity proof & consent trail.
│       │   ├── errors.go                           # PrivacyRequestNotFound, EvidenceInsufficient.
│       │   ├── repository.go                       # Request intake, approval, fulfillment tracking.
│       │   └── events.go                           # admin.privacy.requested/approved/fulfilled/denied.v1
│       ├── pii_access/                             # 🆕 Break-glass unmasking
│       │   ├── entity.go                           # PIIRequest (purpose, scope, approvers, ttl).
│       │   ├── grant.go                            # Grant (who/what/when; reason; mask profile).
│       │   ├── policy.go                           # MaskingPolicy (fields, redaction rules).
│       │   ├── errors.go                           # PIIRequestNotFound, PolicyViolation.
│       │   ├── repository.go                       # Request/grant/audit unmask ops.
│       │   └── events.go                           # admin.pii.unmask.requested/granted/denied/audited.v1
│       ├── ip_claim/                               # 🆕 DMCA/IP cases
│       │   ├── entity.go                           # Claim (owner, work, allegation, status).
│       │   ├── counternotice.go                    # CounterNotice (respondent, grounds).
│       │   ├── decision.go                         # Decision (remove/restore/deny).
│       │   ├── deadline.go                         # SLA/deadline tracking.
│       │   ├── errors.go                           # ClaimNotFound, DeadlineMissed.
│       │   ├── repository.go                       # File/validate/decide/close claims.
│       │   └── events.go                           # admin.ip.claim.filed/validated/removed/closed.v1
│
│       # ------------------------- BILLING REMEDIES ---------------------------
│       ├── refund_case/                            # 🆕 Refund workflows
│       │   ├── entity.go                           # RefundCase (reasonCode, amount, state).
│       │   ├── errors.go                           # RefundCaseNotFound, InvalidState.
│       │   ├── repository.go                       # Request/approve/deny/markPaid lifecycle.
│       │   └── events.go                           # admin.refund.requested/approved/denied/paid.v1
│       ├── goodwill_credit/                        # 🆕 Goodwill credits
│       │   ├── entity.go                           # CreditGrant (subject, amount, expiry).
│       │   ├── reason_code.go                      # ReasonCode catalog (ops remediation).
│       │   ├── errors.go                           # CreditNotFound, AlreadyRevoked.
│       │   ├── repository.go                       # Issue/revoke/list credit grants.
│       │   └── events.go                           # admin.credit.issued/revoked.v1
│
│       # ------------------------- CONFIG & POLICY ----------------------------
│       ├── system_config/
│       │   ├── entity.go                           # Config kv (key, value, scope, updatedBy).
│       │   ├── feature_flag.go                     # FeatureFlag (key, on/off, rollout).
│       │   ├── maintenance.go                      # MaintenanceWindow (from/until/message).
│       │   ├── errors.go                           # ConfigNotFound, ImmutableKey.
│       │   ├── repository.go                       # Get/set config & flags.
│       │   └── events.go                           # admin.feature_flag.updated/admin.threshold.updated/admin.experiment.updated/admin.maintenance.window.set.v1
│       ├── policy_doc/                             # 🆕 TOS/Privacy versions
│       │   ├── entity.go                           # PolicyDoc (kind, currentVersionId).
│       │   ├── version.go                          # Version (hash, changelog, effectiveAt).
│       │   ├── window.go                           # EffectiveWindow + enforcement scope.
│       │   ├── notice.go                           # Notice (who was notified and when).
│       │   ├── errors.go                           # PolicyVersionNotFound, OverlapWindow.
│       │   ├── repository.go                       # Publish/retire versions; record notices.
│       │   └── events.go                           # admin.policy.version.published/retired, admin.policy.notice.sent.v1
│       ├── experiment/                             # 🆕 Feature experiments
│       │   ├── entity.go                           # Experiment (goal, metrics, guardrails).
│       │   ├── variant.go                          # Variant config & weights.
│       │   ├── ramp_schedule.go                    # RampSchedule (phased rollout).
│       │   ├── guardrail.go                        # Guardrail definitions (stop if metric regresses).
│       │   ├── errors.go                           # ExperimentNotFound, RampConflict.
│       │   ├── repository.go                       # Create/activate/update/end experiments.
│       │   └── events.go                           # admin.experiment.created/activated/updated/ended.v1
│       ├── search_policy_admin/                    # 🆕 Search quality knobs
│       │   ├── entity.go                           # SynonymSet, StopwordList, BoostRule, SafetyRule.
│       │   ├── errors.go                           # SearchPolicyNotFound, RuleConflict.
│       │   ├── repository.go                       # Upsert/rollback policy bundles.
│       │   └── events.go                           # admin.search.policy.updated/published/rolled_back.v1
│       ├── throttle_policy/                        # 🆕 Rate-limit overrides
│       │   ├── entity.go                           # ThrottlePolicy (feature, rate, window).
│       │   ├── exception.go                        # Exception (subject, duration).
│       │   ├── window.go                           # Window VO (size/units).
│       │   ├── errors.go                           # ThrottlePolicyNotFound, InvalidWindow.
│       │   ├── repository.go                       # Create/update/disable policies & exceptions.
│       │   └── events.go                           # admin.throttle.policy.created/updated/disabled.v1
│       ├── quota_override/                         # 🆕 Temporary caps/boosts
│       │   ├── entity.go                           # Override (feature, cap, effective/expiry).
│       │   ├── reason.go                           # Reason VO (justification & ticket).
│       │   ├── errors.go                           # OverrideNotFound, AlreadyExpired.
│       │   ├── repository.go                       # Apply/revoke/expire overrides.
│       │   └── events.go                           # admin.quota.override.applied/expired/revoked.v1
│
│       # ------------------------- RISK, FRAUD & INCIDENTS --------------------
│       ├── fraud_review/                           # 🆕 Fraud investigations
│       │   ├── entity.go                           # Case (subject, risk signals, state, owner).
│       │   ├── status.go                           # Pending/Investigating/Cleared/Confirmed.
│       │   ├── severity.go                         # Low/Medium/High/Critical severity.
│       │   ├── reason.go                           # Trigger reason catalog.
│       │   ├── notes.go                            # Investigator notes log.
│       │   ├── sla.go                              # SLA checkpoints for steps.
│       │   ├── errors.go                           # FraudReviewNotFound, InvalidTransition.
│       │   └── repository.go                       # Open/assign/note/setStatus/close.
│       ├── risk_management/                        # 🆕 Financial risk knobs
│       │   ├── hold.go                              # Risk hold (type, amount, reason, status).
│       │   ├── reserve.go                           # Reserve settings & changes.
│       │   ├── chargeback.go                        # Chargeback case lifecycle.
│       │   ├── velocity_alert.go                    # Velocity rule hit records.
│       │   ├── country_rate_anomaly.go              # Geo/rate anomaly record.
│       │   ├── errors.go                            # RiskItemNotFound, RuleConflict.
│       │   ├── repository.go                        # Persist dashboards & actions.
│       │   └── events.go                            # risk.* (admin-side streams) snapshot/updated.v1
│       ├── incident/                                # 🆕 Ops incidents & postmortems
│       │   ├── entity.go                           # Incident (title, sev, status, commander).
│       │   ├── severity.go                         # Sev levels + paging policy hints.
│       │   ├── timeline_event.go                   # Timeline entries (when/what/who).
│       │   ├── rca.go                              # Root-cause analysis doc.
│       │   ├── action_item.go                      # Postmortem action items & follow-up dates.
│       │   ├── errors.go                           # IncidentNotFound, InvalidState.
│       │   ├── repository.go                       # Open/update/resolve; record RCA/actions.
│       │   └── events.go                           # admin.incident.opened/updated/resolved/action_item.created/completed.v1
│
│       # ------------------------- INTEGRATIONS & BULK OPS --------------------
│       ├── integrations_admin/                     # 🆕 Third-party & keys
│       │   ├── entity.go                           # Integration (vendor, scopes, status).
│       │   ├── api_key.go                          # ApiKey (token hash, scopes, expiry, rotation).
│       │   ├── webhook_endpoint.go                 # Webhook endpoint (url, secret, status).
│       │   ├── secret_rotation.go                  # Rotation policy (interval, lastRotatedAt).
│       │   ├── errors.go                           # IntegrationNotFound, KeyRevoked.
│       │   ├── repository.go                       # Add/update/disable; issue/revoke/rotate keys.
│       │   └── events.go                           # admin.integration.added/updated/disabled, admin.api_key.issued/revoked/rotated.v1
│       ├── bulk_action/                            # 🆕 Safe mass operations
│       │   ├── entity.go                           # BulkJob (query, preview, status).
│       │   ├── execution.go                        # Execution batches & progress.
│       │   ├── rollback.go                         # Rollback plan & checkpoints.
│       │   ├── errors.go                           # BulkJobNotFound, PreviewMismatch.
│       │   ├── repository.go                       # Start/preview/commit/rollback flows.
│       │   └── events.go                           # admin.bulk.started/progressed/completed/rolled_back.v1
│
│       # ------------------------- SESSIONS & APPROVALS -----------------------
│       ├── admin_session/                          # 🆕 Just-in-time / break-glass
│       │   ├── entity.go                           # Session (actor, reason, scope grants, expiry).
│       │   ├── scope_grant.go                      # ScopeGrant (resource scope & ttl).
│       │   ├── reason.go                           # Reason enumeration (incident, audit, court-order).
│       │   ├── errors.go                           # SessionNotFound, GrantInvalid, Denied.
│       │   ├── repository.go                       # Start/end sessions; grant/revoke scopes; audit.
│       │   └── events.go                           # admin.session.started/ended, admin.break_glass.granted/denied.v1
│       ├── change_approval/                        # 🆕 Two-person rule
│       │   ├── entity.go                           # ChangeRequest (resource, diff, risk, state).
│       │   ├── approval.go                         # Approval (approver, decision, rationale).
│       │   ├── policy.go                           # Policy (which actions require approval).
│       │   ├── errors.go                           # ChangeRequestNotFound, Expired, PolicyMiss.
│       │   ├── repository.go                       # Request/approve/reject/expire tracking.
│       │   └── events.go                           # admin.change.requested/approved/rejected/expired.v1
│
│       # ------------------------- OUTBOX (MOVED) -----------------------------
│       └── outbox/
│           ├── entity.go                           # ❌ REMOVED – use platform-shared/outbox/entity.go
│           └── repository.go                        # ❌ REMOVED – use platform-shared/outbox/repository.go
│
│   # =====================================================================
│   # 📋 APPLICATION LAYER — Use-cases, Orchestrators, DTOs, Validators
│   # =====================================================================
│   └── application/
│       # ------------------------- EVENT CONSUMERS (INBOX) -------------------
│       ├── eventhandler/
│       │   ├── storage_handler.go                  # Ingest storage-be signals → moderation/DLP/audit updates.
│       │   ├── search_handler.go                   # Track search-be taxonomy/LTR/hygiene/index ops for QA dashboards.
│       │   ├── reviews_handler.go                  # Mirror double-blind & moderation states into admin cases.
│       │   ├── subscriptions_handler.go            # Reflect subs/entitlements/usage/billing into admin views.
│       │   ├── contracts_handler.go                # Intake contract hold/state changes for risk/dispute linking.
│       │   ├── financial_handler.go                # Payments/chargebacks/holds/reserves → risk dashboards.
│       │   ├── users_handler.go                    # User updates & cross-linked disputes to user case pages.
│       │   ├── reports_exports_handler.go          # Drive admin data-export lifecycle & audit.
│       │   └── communications_handler.go           # Append comm.delivery.logged into audit streams.
│
│       # ------------------------- CORE ADMIN --------------------------------
│       ├── admin_user/
│       │   ├── service.go                          # CRUD admins; manage roles/perms; write audit.
│       │   ├── commands.go                         # CreateAdmin, UpdateAdmin, DeactivateAdmin, SetPermissions.
│       │   ├── queries.go                          # GetAdmin, ListAdmins, SearchAdmins.
│       │   ├── dto.go                               # Admin DTOs (read/write).
│       │   ├── mapper.go                            # Entity ↔ DTO transforms.
│       │   ├── validators.go                        # Email/role/permission constraints.
│       │   └── permission_manager.go                # Role/permission expansion + conflict detection.
│
│       # ------------------------- SUPPORT & CASEWORK -------------------------
│       ├── support_ticket/
│       │   ├── service.go                          # Ticket lifecycle; SLA; assignment policies.
│       │   ├── commands.go                         # OpenTicket, AssignTicket, ResolveTicket, CloseTicket.
│       │   ├── queries.go                          # GetTicket, ListTickets, SearchTickets.
│       │   ├── dto.go                               # Ticket DTOs.
│       │   ├── mapper.go                            # Ticket ↔ DTO mapping.
│       │   ├── validators.go                        # Category/priority/status transitions.
│       │   ├── assignment_engine.go                 # Auto-assign by skills/load/priority.
│       │   ├── escalation_manager.go                # Escalation rules & timers.
│       │   └── sla_tracker.go                       # SLA clock calculation & breach events.
│       ├── ticket_message/
│       │   ├── service.go                          # Post/edit/delete; attachment auth checks.
│       │   ├── commands.go                         # AddNote, EditNote, DeleteNote, AddAttachment.
│       │   ├── queries.go                          # ListNotesForTicket, GetNote.
│       │   ├── dto.go                               # Message/Attachment DTOs.
│       │   ├── mapper.go                            # Map entities ↔ DTOs.
│       │   └── validators.go                        # Body/visibility/attachment size validation.
│       ├── support_agent/
│       │   ├── service.go                          # Manage availability & workloads; compute stats.
│       │   ├── commands.go                         # SetAvailability, AssignQueue.
│       │   ├── queries.go                          # GetAgent, ListAgents, GetStats.
│       │   ├── dto.go                               # Agent DTOs.
│       │   ├── mapper.go                            # Map agent/profile/stats.
│       │   ├── validators.go                        # Skill/queue constraints.
│       │   └── stats_calculator.go                  # KPI computations windowed/weighted.
│
│       # ------------------------- CONTENT & KNOWLEDGE ------------------------
│       ├── canned_response/
│       │   ├── service.go                          # CRUD + category management.
│       │   ├── commands.go                         # CreateResponse, UpdateResponse, ArchiveResponse.
│       │   ├── queries.go                          # GetResponse, ListResponses.
│       │   ├── dto.go                               # Response DTOs.
│       │   ├── mapper.go                            # Map response/category.
│       │   └── validators.go                        # Uniqueness, locale bounds.
│       ├── knowledge_base/
│       │   ├── service.go                          # Draft/version/publish KB.
│       │   ├── commands.go                         # CreateArticle, UpdateArticle, PublishArticle.
│       │   ├── queries.go                          # GetArticle, SearchArticles.
│       │   ├── dto.go                               # Article/Version DTOs.
│       │   ├── mapper.go                            # Map article/version.
│       │   ├── validators.go                        # Version/locale checks.
│       │   └── search_service.go                    # KB search (ES/PG trigram) adapter.
│       ├── faq/
│       │   ├── service.go                          # Manage FAQs & categories.
│       │   ├── commands.go                         # CreateFAQ, UpdateFAQ, PublishFAQ.
│       │   ├── queries.go                          # GetFAQ, ListFAQs.
│       │   ├── dto.go                               # FAQ DTOs.
│       │   ├── mapper.go                            # Map FAQ/category.
│       │   └── validators.go                        # Order/locale checks.
│
│       # ------------------------- SAFETY & MODERATION ------------------------
│       ├── moderation/
│       │   ├── service.go                          # Queue ops; auto-mod & decisions.
│       │   ├── commands.go                         # ApproveContent, RejectContent, RemoveContent.
│       │   ├── queries.go                          # GetQueue, FilterQueue.
│       │   ├── dto.go                               # Queue item/decision DTOs.
│       │   ├── mapper.go                            # Map queue/action.
│       │   ├── validators.go                        # Reason thresholds & policy checks.
│       │   ├── queue_manager.go                     # Assignment & aging logic.
│       │   ├── auto_moderator.go                    # Heuristics rules (no ML).
│       │   └── content_scanner.go                   # Lightweight rule-based scanner hooks.
│       ├── user_management/
│       │   ├── service.go                          # Suspend/ban/verify/warn with audit.
│       │   ├── commands.go                         # SuspendUser, BanUser, VerifyUser, WarnUser.
│       │   ├── queries.go                          # SearchUsers, GetUserAdminView.
│       │   ├── dto.go                               # User admin DTOs.
│       │   ├── mapper.go                            # Map identity & flags.
│       │   ├── validators.go                        # Action preconditions.
│       │   └── action_validator.go                  # Risk/appeal/ownership guards.
│       ├── content_management/
│       │   ├── service.go                          # Remove/hide/feature content safely.
│       │   ├── commands.go                         # RemoveContent, HideContent, FeatureContent.
│       │   ├── queries.go                          # ListContent, SearchContent.
│       │   ├── dto.go                               # Content DTOs.
│       │   ├── mapper.go                            # Content mapping helpers.
│       │   └── validators.go                        # Decision & scope validation.
│
│       # ------------------------- DISPUTES & LEGAL ---------------------------
│       ├── dispute_resolution/
│       │   ├── service.go                          # Orchestrate dispute lifecycle.
│       │   ├── commands.go                         # OpenDispute, AddEvidence, DecideDispute, CloseDispute.
│       │   ├── queries.go                          # GetDispute, ListDisputes.
│       │   ├── dto.go                               # Dispute DTOs.
│       │   ├── mapper.go                            # Map evidence/decision.
│       │   ├── validators.go                        # State transitions & SLA checks.
│       │   └── decision_engine.go                   # Helper: decision templates & checks.
│       ├── appeal/
│       │   ├── service.go                          # Manage appeals; enforce windows.
│       │   ├── commands.go                         # OpenAppeal, DecideAppeal, EscalateAppeal.
│       │   ├── queries.go                          # GetAppeal, ListAppeals.
│       │   ├── dto.go                               # Appeal DTOs.
│       │   ├── mapper.go                            # Map appeal/outcome.
│       │   └── validators.go                        # Window, target, role checks.
│       ├── legal_hold/
│       │   ├── service.go                          # Place/release holds; manage exports.
│       │   ├── commands.go                         # PlaceHold, ReleaseHold, CreateExport.
│       │   ├── queries.go                          # GetHold, ListHolds, GetExport.
│       │   ├── dto.go                               # Hold/Export DTOs.
│       │   ├── mapper.go                            # Map hold/export.
│       │   └── validators.go                        # Scope/precedence checks.
│       ├── privacy_request/
│       │   ├── service.go                          # DSAR intake → fulfill/deny paths.
│       │   ├── commands.go                         # RequestErasure, RequestAccess, Approve, Fulfill, Deny.
│       │   ├── queries.go                          # GetRequest, ListRequests.
│       │   ├── dto.go                               # DSAR DTOs.
│       │   ├── mapper.go                            # Map DSAR/evidence.
│       │   └── validators.go                        # Identity proof & scope checks.
│       ├── pii_access/
│       │   ├── service.go                          # Break-glass flow + tight audit.
│       │   ├── commands.go                         # RequestUnmask, ApproveUnmask, DenyUnmask.
│       │   ├── queries.go                          # GetPIIRequest, ListPIIRequests.
│       │   ├── dto.go                               # PII request/grant DTOs.
│       │   ├── mapper.go                            # Map request/grant/policy.
│       │   └── validators.go                        # Purpose/least-privilege checks.
│       ├── ip_claim/
│       │   ├── service.go                          # DMCA/IP claim lifecycle & deadlines.
│       │   ├── commands.go                         # FileClaim, ValidateClaim, DecideClaim, CloseClaim.
│       │   ├── queries.go                          # GetClaim, ListClaims.
│       │   ├── dto.go                               # Claim/CounterNotice DTOs.
│       │   ├── mapper.go                            # Map claim/decision.
│       │   └── validators.go                        # Evidence & SLA checks.
│
│       # ------------------------- BILLING REMEDIES ---------------------------
│       ├── refund_case/
│       │   ├── service.go                          # Refund case orchestration & approvals.
│       │   ├── commands.go                         # RequestRefund, ApproveRefund, DenyRefund, MarkPaid.
│       │   ├── queries.go                          # GetRefund, ListRefunds.
│       │   ├── dto.go                               # Refund DTOs.
│       │   ├── mapper.go                            # Map refund entities.
│       │   └── validators.go                        # Amount/eligibility checks.
│       ├── goodwill_credit/
│       │   ├── service.go                          # Issue/revoke goodwill credits.
│       │   ├── commands.go                         # IssueCredit, RevokeCredit.
│       │   ├── queries.go                          # GetCredit, ListCredits.
│       │   ├── dto.go                               # Credit DTOs.
│       │   ├── mapper.go                            # Map credit & reasons.
│       │   └── validators.go                        # Cap/expiry validation.
│
│       # ------------------------- CONFIG & POLICY ----------------------------
│       ├── system_config/
│       │   ├── service.go                          # Update flags/config; schedule maintenance.
│       │   ├── commands.go                         # SetFlag, SetConfig, SetMaintenanceWindow.
│       │   ├── queries.go                          # GetFlag, GetConfig.
│       │   ├── dto.go                               # Config/Flag DTOs.
│       │   ├── mapper.go                            # Map configs.
│       │   └── validators.go                        # Key immutability & scope checks.
│       ├── policy_doc/
│       │   ├── service.go                          # Publish/retire policy versions + notices.
│       │   ├── commands.go                         # PublishVersion, RetireVersion, SendNotice.
│       │   ├── queries.go                          # GetPolicy, GetVersion, ActiveWindow.
│       │   ├── dto.go                               # Policy/Version/Notice DTOs.
│       │   ├── mapper.go                            # Policy mapping helpers.
│       │   └── validators.go                        # Window overlap/version checks.
│       ├── experiment/
│       │   ├── service.go                          # Create/update; ramp & guardrails.
│       │   ├── commands.go                         # CreateExperiment, Activate, Update, End.
│       │   ├── queries.go                          # GetExperiment, ListExperiments.
│       │   ├── dto.go                               # Experiment/Variant DTOs.
│       │   ├── mapper.go                            # Map experiment artifacts.
│       │   └── validators.go                        # Ramp schedule & guardrail bounds.
│       ├── search_policy_admin/
│       │   ├── service.go                          # Manage synonym/stopword/boost/safety sets.
│       │   ├── commands.go                         # UpsertSynonyms, UpsertStopwords, SetBoostRules.
│       │   ├── queries.go                          # GetSearchPolicyBundle.
│       │   ├── dto.go                               # Policy bundle DTOs.
│       │   ├── mapper.go                            # Map to search-be contract.
│       │   └── validators.go                        # Conflict/locale checks.
│       ├── throttle_policy/
│       │   ├── service.go                          # Author & apply throttle overrides/exceptions.
│       │   ├── commands.go                         # CreatePolicy, UpdatePolicy, DisablePolicy, AddException.
│       │   ├── queries.go                          # GetPolicy, ListPolicies.
│       │   ├── dto.go                               # Throttle policy DTOs.
│       │   ├── mapper.go                            # Map windows & rates.
│       │   └── validators.go                        # Window/rate sanity checks.
│       ├── quota_override/
│       │   ├── service.go                          # Temporary caps/boosts lifecycle.
│       │   ├── commands.go                         # ApplyOverride, RevokeOverride.
│       │   ├── queries.go                          # GetOverride, ListOverrides.
│       │   ├── dto.go                               # Override DTOs.
│       │   ├── mapper.go                            # Map overrides.
│       │   └── validators.go                        # Duration/feature gating checks.
│
│       # ------------------------- RISK, FRAUD & INCIDENTS --------------------
│       ├── fraud_review/
│       │   ├── service.go                          # Case intake, triage, notes, status changes.
│       │   ├── commands.go                         # OpenReview, Assign, AddNote, SetStatus, LinkEvidence.
│       │   ├── queries.go                          # GetReview, ListReviews, SearchReviews.
│       │   ├── dto.go                               # Fraud case DTOs.
│       │   ├── mapper.go                            # Map risk artifacts.
│       │   └── validators.go                        # Transition & SLA checks.
│       ├── risk/
│       │   ├── service.go                          # Holds/reserves/chargebacks dashboards.
│       │   ├── commands.go                         # PlaceHold, ReleaseHold, SetReserve, RecordChargeback.
│       │   ├── queries.go                          # GetRiskSummary, ListAlerts, Anomalies.
│       │   ├── dto.go                               # Risk DTOs.
│       │   ├── mapper.go                            # Map risk items.
│       │   └── validators.go                        # Amount/policy checks.
│       ├── incident/
│       │   ├── service.go                          # Incident lifecycle & postmortems.
│       │   ├── commands.go                         # OpenIncident, UpdateIncident, ResolveIncident, AddActionItem.
│       │   ├── queries.go                          # GetIncident, ListIncidents.
│       │   ├── dto.go                               # Incident & RCA DTOs.
│       │   ├── mapper.go                            # Incident mapping.
│       │   └── validators.go                        # Severity/state checks.
│
│       # ------------------------- INTEGRATIONS & BULK OPS --------------------
│       ├── integrations_admin/
│       │   ├── service.go                          # Manage integrations, API keys, webhooks.
│       │   ├── commands.go                         # AddIntegration, UpdateIntegration, DisableIntegration, IssueKey, RevokeKey, RotateKey.
│       │   ├── queries.go                          # GetIntegration, ListIntegrations.
│       │   ├── dto.go                               # Integration/API key DTOs.
│       │   ├── mapper.go                            # Map integrations & endpoints.
│       │   └── validators.go                        # Scope/rotation/secret rules.
│       ├── bulk_action/
│       │   ├── service.go                          # Preview then execute with rollback.
│       │   ├── commands.go                         # StartBulk, CommitBulk, RollbackBulk.
│       │   ├── queries.go                          # GetBulkJob, ListBulkJobs.
│       │   ├── dto.go                               # Bulk job DTOs.
│       │   ├── mapper.go                            # Map previews/executions.
│       │   └── validators.go                        # Preview/result set alignment checks.
│
│       # ------------------------- SESSIONS & APPROVALS -----------------------
│       ├── admin_session/
│       │   ├── service.go                          # JIT/break-glass session orchestration.
│       │   ├── commands.go                         # StartSession, GrantScope, EndSession.
│       │   ├── queries.go                          # GetSession, ListSessions.
│       │   ├── dto.go                               # Session & scope DTOs.
│       │   ├── mapper.go                            # Map session/grants.
│       │   └── validators.go                        # Purpose/scope/timebox checks.
│       ├── change_approval/
│       │   ├── service.go                          # Two-person approval workflow.
│       │   ├── commands.go                         # RequestChange, ApproveChange, RejectChange, ExpireChange.
│       │   ├── queries.go                          # GetChangeRequest, ListChangeRequests.
│       │   ├── dto.go                               # Change request/approval DTOs.
│       │   ├── mapper.go                            # Map policy & diffs.
│       │   └── validators.go                        # Policy coverage & risk checks.
│
│   # =====================================================================
│   # 🔌 INFRASTRUCTURE — DB, Cache, Messaging, External Clients, etc.
│   # =====================================================================
│   └── infrastructure/
│       # ------------------------- 🗄️ PERSISTENCE (POSTGRES) -------------------
│       ├── persistence/
│       │   └── postgres/
│       │       ├── connection.go                     # DSN, pooling, timeouts, observability tags.
│       │       ├── transaction.go                    # WithTx helpers; savepoints for nested ops.
│       │       ├── migrations.go                     # Auto-migrate + ordered version steps.
│       │       ├── version.go                        # schema_version helpers and recorders.
│       │       ├── safety.go                         # Env/disk sanity checks pre-migration.
│       │       # ---- CORE ADMIN
│       │       ├── admin_user_repository.go          # Admin users + audits repo.
│       │       # ---- SUPPORT & CASEWORK
│       │       ├── support_ticket_repository.go      # Tickets repo.
│       │       ├── ticket_message_repository.go      # Ticket messages repo.
│       │       ├── support_agent_repository.go       # Agents/availability/stats repo.
│       │       # ---- CONTENT & KNOWLEDGE
│       │       ├── canned_response_repository.go     # Canned responses repo.
│       │       ├── knowledge_base_repository.go      # KB articles/versions repo.
│       │       ├── faq_repository.go                 # FAQ repo.
│       │       # ---- SAFETY & MODERATION
│       │       ├── moderation_queue_repository.go    # Moderation queue repo.
│       │       ├── user_action_repository.go         # User actions repo.
│       │       ├── content_action_repository.go      # Content actions repo.
│       │       # ---- DISPUTES & LEGAL
│       │       ├── dispute_resolution_repository.go  # Disputes repo.
│       │       ├── appeal_repository.go              # 🆕 Appeals repo.
│       │       ├── case_link_repository.go           # 🆕 Case linking repo.
│       │       ├── legal_hold_repository.go          # 🆕 Holds & exports repo.
│       │       ├── privacy_request_repository.go     # 🆕 DSAR requests repo.
│       │       ├── pii_access_repository.go          # 🆕 PII unmask requests/grants repo.
│       │       ├── ip_claim_repository.go            # 🆕 IP/DMCA claims repo.
│       │       # ---- IDENTITY & VERIFICATION
│       │       ├── kyc_case_repository.go            # 🆕 KYC cases/documents repo.
│       │       ├── business_verification_repository.go # 🆕 Business verification repo.
│       │       ├── sanctions_screening_repository.go # 🆕 Screening runs/hits repo.
│       │       # ---- BILLING REMEDIES
│       │       ├── refund_case_repository.go         # 🆕 Refund cases repo.
│       │       ├── goodwill_credit_repository.go     # 🆕 Goodwill credits repo.
│       │       # ---- CONFIG & POLICY
│       │       ├── system_config_repository.go       # System config/flags repo.
│       │       ├── policy_doc_repository.go          # 🆕 Policy documents/versions repo.
│       │       ├── experiment_repository.go          # 🆕 Experiments & ramps repo.
│       │       ├── search_policy_admin_repository.go # 🆕 Search policy bundles repo.
│       │       ├── throttle_policy_repository.go     # 🆕 Throttle policies/exceptions repo.
│       │       ├── quota_override_repository.go      # 🆕 Quota overrides repo.
│       │       # ---- RISK, FRAUD & INCIDENTS
│       │       ├── fraud_review_repository.go        # 🆕 Fraud cases repo.
│       │       ├── risk_hold_repository.go           # 🆕 Risk holds repo.
│       │       ├── risk_reserve_repository.go        # 🆕 Risk reserves repo.
│       │       ├── risk_chargeback_repository.go     # 🆕 Chargeback cases repo.
│       │       ├── risk_velocity_alert_repository.go # 🆕 Velocity alerts repo.
│       │       ├── risk_country_rate_anomaly_repository.go # 🆕 Geo/rate anomalies repo.
│       │       ├── incident_repository.go            # 🆕 Incidents/RCA repo.
│       │       # ---- INTEGRATIONS & BULK
│       │       ├── integrations_admin_repository.go  # 🆕 Integrations/keys/webhooks repo.
│       │       ├── bulk_action_repository.go         # 🆕 Bulk jobs/executions/rollbacks repo.
│       │       # ---- SESSIONS & APPROVALS
│       │       ├── admin_session_repository.go       # 🆕 JIT/break-glass sessions repo.
│       │       ├── change_approval_repository.go     # 🆕 Change requests/approvals repo.
│       │       # ---- OUTBOX (MOVED)
│       │       └── outbox_repository.go              # ❌ REMOVED → platform-shared/outbox/postgres
│
│       # ------------------------- ⚡ CACHE (REDIS) ----------------------------
│       ├── cache/
│       │   └── redis/
│       │       ├── connection.go                     # Redis client & pool; health pings.
│       │       # ---- SUPPORT & CASEWORK
│       │       ├── ticket_cache.go                   # Ticket hot fields & SLA clocks cache.
│       │       ├── admin_cache.go                    # Admin profile/perm snapshot cache.
│       │       ├── stats_cache.go                    # Agent/ticket KPI snapshots.
│       │       # ---- CONFIG & POLICY
│       │       ├── config_cache.go                   # Feature flags/configs snapshot.
│       │       ├── search_policy_cache.go            # 🆕 Search knobs bundle cache.
│       │       ├── throttle_policy_cache.go          # 🆕 Throttle policies hot-path cache.
│       │       ├── quota_override_cache.go           # 🆕 Overrides with TTL eviction.
│       │       # ---- SESSIONS & APPROVALS
│       │       ├── admin_session_cache.go            # 🆕 Active JIT sessions & grants.
│       │       └── change_approval_cache.go          # 🆕 Pending approvals quick lookup.
│
│       # ------------------------- 📨 MESSAGING (KAFKA) ------------------------
│       ├── messaging/
│       │   └── kafka/
│       │       ├── consumer.go                       # Uses platform-shared/inbox for dedupe & offset mgmt.
│       │       ├── producer.go                       # Uses platform-shared/outbox for reliable publish.
│       │       ├── topics.go                         # Imports contracts/events topic constants.
│       │       └── scram.go                          # SASL/SCRAM-256 setup.
│
│       # ------------------------- EXTERNAL CLIENTS ---------------------------
│       ├── external_services/
│       │   ├── users_client.go                       # Users service client (admin reads).
│       │   ├── jobs_client.go                        # Jobs service client.
│       │   ├── proposals_client.go                   # Proposals client.
│       │   ├── contracts_client.go                   # Contracts client.
│       │   ├── financial_client.go                   # Financial/payments client.
│       │   ├── reviews_client.go                     # Reviews client.
│       │   ├── communications_client.go              # Communications client.
│       │   ├── search_client.go                      # Search client.
│       │   ├── storage_client.go                     # Storage client.
│       │   └── subscriptions_client.go               # Subscriptions client.
│       ├── keycloak/
│       │   ├── admin_client.go                       # Keycloak admin REST wrapper.
│       │   └── user_manager.go                       # Manage user states/roles.
│       ├── reporting/
│       │   ├── pdf_generator.go                      # PDF renderer (reports/exports).
│       │   ├── csv_generator.go                      # CSV exporter.
│       │   └── excel_generator.go                    # XLSX exporter.
│       └── outbox/
│           ├── processor.go                          # ❌ REMOVED – use platform-shared/outbox/forwarder.go
│           └── scheduler.go                          # ❌ REMOVED – use platform-shared/outbox/scheduler.go
│
│   # =====================================================================
│   # 🌐 INTERFACES — HTTP (v1), Middleware, Responses, Router
│   # =====================================================================
│   └── interfaces/
│       └── http/
│           └── v1/
│               # ---------------------- 🧭 HANDLERS -------------------------
│               ├── handlers/
│               │   # ---- CORE ADMIN
│               │   ├── admin_user_handler.go            # CRUD admins; roles & permissions endpoints.
│               │   # ---- SUPPORT & CASEWORK
│               │   ├── support_ticket_handler.go        # Tickets CRUD/assign/resolve.
│               │   ├── ticket_message_handler.go        # Ticket notes & attachments.
│               │   ├── support_agent_handler.go         # Agent availability & stats.
│               │   # ---- CONTENT & KNOWLEDGE
│               │   ├── canned_response_handler.go       # Canned responses CRUD.
│               │   ├── knowledge_base_handler.go        # KB articles/versions/search.
│               │   ├── faq_handler.go                   # FAQs CRUD.
│               │   # ---- SAFETY & MODERATION
│               │   ├── moderation_handler.go            # Moderation queue & actions.
│               │   ├── user_management_handler.go       # Suspend/ban/verify/warn.
│               │   ├── content_management_handler.go    # Remove/hide/feature content.
│               │   # ---- DISPUTES & LEGAL
│               │   ├── dispute_resolution_handler.go    # Disputes lifecycle.
│               │   ├── appeal_handler.go                # 🆕 Appeals endpoints.
│               │   ├── legal_hold_handler.go            # 🆕 Holds & eDiscovery exports.
│               │   ├── privacy_request_handler.go       # 🆕 DSAR endpoints.
│               │   ├── pii_access_handler.go            # 🆕 PII unmask requests.
│               │   ├── ip_claim_handler.go              # 🆕 IP/DMCA claims.
│               │   # ---- IDENTITY & VERIFICATION
│               │   ├── kyc_case_handler.go              # 🆕 KYC cases & decisions.
│               │   ├── business_verification_handler.go # 🆕 Business verification.
│               │   ├── sanctions_screening_handler.go   # 🆕 Screening runs/hits.
│               │   # ---- BILLING REMEDIES
│               │   ├── refund_case_handler.go           # 🆕 Refund case endpoints.
│               │   ├── goodwill_credit_handler.go       # 🆕 Goodwill credits.
│               │   # ---- CONFIG & POLICY
│               │   ├── system_config_handler.go         # Flags/config/maintenance.
│               │   ├── policy_doc_handler.go            # 🆕 Policy versions & notices.
│               │   ├── experiment_handler.go            # 🆕 Experiments & ramps.
│               │   ├── search_policy_admin_handler.go   # 🆕 Search policy bundles.
│               │   ├── throttle_policy_handler.go       # 🆕 Throttle overrides.
│               │   ├── quota_override_handler.go        # 🆕 Quota overrides.
│               │   # ---- RISK, FRAUD & INCIDENTS
│               │   ├── fraud_review_handler.go          # 🆕 Fraud case ops.
│               │   ├── risk_handler.go                  # 🆕 Risk dashboards/actions.
│               │   ├── incident_handler.go              # 🆕 Incidents & RCA.
│               │   # ---- INTEGRATIONS & BULK OPS
│               │   ├── integrations_admin_handler.go    # 🆕 Integrations, API keys, webhooks.
│               │   ├── bulk_action_handler.go           # 🆕 Bulk jobs previews/execs.
│               │   # ---- SESSIONS & APPROVALS
│               │   ├── admin_session_handler.go         # 🆕 Start/end JIT sessions; grants.
│               │   ├── change_approval_handler.go       # 🆕 Change requests/approvals.
│               │   # ---- HEALTH
│               │   └── health_handler.go                # /health, /ready, /live endpoints.
│               │
│               # ---------------------- 🗺️ ROUTES ---------------------------
│               ├── routes/
│               │   # ---- CORE ADMIN
│               │   ├── admin_user_routes.go             # /v1/admin/users/*
│               │   # ---- SUPPORT & CASEWORK
│               │   ├── support_ticket_routes.go         # /v1/admin/tickets/*
│               │   ├── ticket_message_routes.go         # /v1/admin/tickets/:id/messages/*
│               │   ├── support_agent_routes.go          # /v1/admin/agents/*
│               │   # ---- CONTENT & KNOWLEDGE
│               │   ├── canned_response_routes.go        # /v1/admin/canned-responses/*
│               │   ├── knowledge_base_routes.go         # /v1/admin/kb/*
│               │   ├── faq_routes.go                    # /v1/admin/faqs/*
│               │   # ---- SAFETY & MODERATION
│               │   ├── moderation_routes.go             # /v1/admin/moderation/*
│               │   ├── user_management_routes.go        # /v1/admin/users/actions/*
│               │   ├── content_management_routes.go     # /v1/admin/content/*
│               │   # ---- DISPUTES & LEGAL
│               │   ├── dispute_resolution_routes.go     # /v1/admin/disputes/*
│               │   ├── appeal_routes.go                 # 🆕 /v1/admin/appeals/*
│               │   ├── legal_hold_routes.go             # 🆕 /v1/admin/legal-holds/*
│               │   ├── privacy_request_routes.go        # 🆕 /v1/admin/privacy-requests/*
│               │   ├── pii_access_routes.go             # 🆕 /v1/admin/pii/*
│               │   ├── ip_claim_routes.go               # 🆕 /v1/admin/ip-claims/*
│               │   # ---- IDENTITY & VERIFICATION
│               │   ├── kyc_case_routes.go               # 🆕 /v1/admin/kyc/*
│               │   ├── business_verification_routes.go  # 🆕 /v1/admin/business-verifications/*
│               │   ├── sanctions_screening_routes.go    # 🆕 /v1/admin/sanctions/*
│               │   # ---- BILLING REMEDIES
│               │   ├── refund_case_routes.go            # 🆕 /v1/admin/refunds/*
│               │   ├── goodwill_credit_routes.go        # 🆕 /v1/admin/credits/*
│               │   # ---- CONFIG & POLICY
│               │   ├── system_config_routes.go          # /v1/admin/config/*
│               │   ├── policy_doc_routes.go             # 🆕 /v1/admin/policies/*
│               │   ├── experiment_routes.go             # 🆕 /v1/admin/experiments/*
│               │   ├── search_policy_admin_routes.go    # 🆕 /v1/admin/search-policy/*
│               │   ├── throttle_policy_routes.go        # 🆕 /v1/admin/throttles/*
│               │   ├── quota_override_routes.go         # 🆕 /v1/admin/overrides/*
│               │   # ---- RISK, FRAUD & INCIDENTS
│               │   ├── fraud_review_routes.go           # 🆕 /v1/admin/fraud-reviews/*
│               │   ├── risk_routes.go                   # 🆕 /v1/admin/risk/*
│               │   ├── incident_routes.go               # 🆕 /v1/admin/incidents/*
│               │   # ---- INTEGRATIONS & BULK OPS
│               │   ├── integrations_admin_routes.go     # 🆕 /v1/admin/integrations/*
│               │   ├── bulk_action_routes.go            # 🆕 /v1/admin/bulk/*
│               │   # ---- SESSIONS & APPROVALS
│               │   ├── admin_session_routes.go          # 🆕 /v1/admin/sessions/*
│               │   └── change_approval_routes.go        # 🆕 /v1/admin/change-approvals/*
│               │
│               # ---------------------- 🧱 MIDDLEWARE ------------------------
│               ├── middleware/
│               │   ├── auth.go                          # JWT auth (pkg/auth).
│               │   ├── admin_auth.go                    # Admin-only guard (realm/roles).
│               │   ├── permission_check.go              # Check fine-grained permissions.
│               │   ├── audit_logger.go                  # Auto-log admin actions to audit trail.
│               │   ├── cors.go                          # CORS config (platform-shared/ginx).
│               │   ├── rate_limit.go                    # Token-bucket rate limit.
│               │   ├── logging.go                       # Structured request logs.
│               │   ├── error_handler.go                 # Uniform error responses.
│               │   └── request_id.go                    # X-Request-ID injector.
│               # ---------------------- 📤 RESPONSES -------------------------
│               ├── responses/
│               │   ├── success.go                       # Success envelope (platform-shared/httpx).
│               │   ├── error.go                         # Error envelope (platform-shared/httpx).
│               │   └── pagination.go                    # Pagination helpers (platform-shared/httpx).
│               # ---------------------- 🚦 ROUTER ----------------------------
│               └── router.go                            # Build Gin engine; mount /v1; apply middleware & registrars.
│
│   # =====================================================================
│   # 🔧 CONFIG (typed schema + loader + docs)
│   # =====================================================================
│   └── config/
│       ├── schema.go                                  # App/Server/Postgres/Kafka/Redis/Keycloak typed config.
│       ├── loader.go                                  # Viper loader: flags → env → file → defaults.
│       └── docs/
│           └── CONFIGURATION.md                       # ENV vars, defaults, examples.
│
├── config/
│   ├── default.yaml                                   # Baseline configuration.
│   ├── dev.yaml                                       # Local/dev overrides.
│   └── prod.yaml                                      # Production overrides.
│
├── dapr/
│   ├── local/
│   │   ├── pubsub.yaml                                # Kafka pub/sub component.
│   │   └── statestore.yaml                            # Dapr state store.
│   └── k8s/
│       ├── pubsub.yaml                                # Kafka with scopes: ["admin-be"].
│       ├── statestore.yaml                            # Namespaced state store.
│       └── secrets.yaml                               # Secret store decl.
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                                  # Service-scoped error helpers.
│   │   └── codes.go                                   # Canonical error codes.
│   ├── logger/
│   │   └── logger.go                                  # ❌ REMOVED – use platform-shared/logging.
│   ├── utils/
│   │   ├── validator.go                               # Common input validation helpers.
│   │   ├── sanitizer.go                               # Input sanitizer (HTML/URL/PII scrubbing).
│   │   ├── permission_checker.go                      # Compose/verify permissions for handlers.
│   │   └── report_formatter.go                        # Report formatting helpers.
│   └── constants/
│       ├── events.go                                  # ❌ REMOVED – use contracts/events.
│       ├── topics.go                                  # ❌ REMOVED – use contracts/events.
│       ├── permissions.go                             # Permission constants (by feature).
│       └── moderation_actions.go                      # Moderation action constants.
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                            # Deployment (env, probes, resources).
│       ├── service.yaml                               # ClusterIP/ports.
│       ├── configmap.yaml                             # App config as CM.
│       ├── secrets.yaml                               # Secret refs.
│       ├── hpa.yaml                                   # Autoscaling policy.
│       ├── pdb.yaml                                   # PodDisruptionBudget.
│       ├── rbac.yaml                                  # Admin RBAC for cluster ops.
│       └── servicemonitor.yaml                        # Prometheus ServiceMonitor.
│
├── scripts/
│   ├── setup-local.sh                                 # Bootstrap local dev env.
│   ├── get-secrets.sh                                 # Pull secrets for local/dev.
│   ├── seed-admin-users.sh                            # Seed baseline admin accounts.
│   ├── seed-canned-responses.sh                       # Seed canned responses.
│   └── seed-data.sh                                   # Populate sample data.
│
├── tests/
│   ├── unit/
│   │   ├── domain/
│   │   │   ├── admin_user_test.go                    # Admin domain tests.
│   │   │   ├── support_ticket_test.go                # Ticket domain tests.
│   │   │   └── moderation_queue_test.go              # Moderation domain tests.
│   │   ├── application/
│   │   │   ├── admin_user_service_test.go            # Admin service tests.
│   │   │   ├── support_ticket_service_test.go        # Ticket service tests.
│   │   │   └── moderation_service_test.go            # Moderation service tests.
│   │   └── infrastructure/
│   │       ├── postgres_repository_test.go           # PG repo tests.
│   │       └── kafka_producer_test.go                # Kafka producer tests.
│   ├── integration/
│   │   ├── handlers/
│   │   │   ├── admin_user_handler_test.go            # Admin HTTP tests.
│   │   │   ├── support_ticket_handler_test.go        # Ticket HTTP tests.
│   │   │   └── moderation_handler_test.go            # Moderation HTTP tests.
│   │   └── repositories/
│   │       ├── admin_user_repository_test.go         # Admin repo integration tests.
│   │       └── support_ticket_repository_test.go     # Ticket repo integration tests.
│   └── e2e/
│       └── scenarios/
│           ├── ticket_workflow_test.go               # Open→assign→resolve flow.
│           ├── moderation_workflow_test.go           # Flag→review→action flow.
│           └── dispute_resolution_test.go            # Dispute end-to-end.
│
├── docs/
│   ├── README.md                                     # Service overview & responsibilities.
│   ├── api.md                                        # HTTP API reference.
│   ├── events.md                                     # Published/consumed event catalog.
│   ├── admin-roles.md                                # Roles & permissions matrix.
│   ├── permissions.md                                # Permission model & usage.
│   ├── moderation-guide.md                           # Moderation runbook.
│   ├── support-workflows.md                          # Support SLAs & flows.
│   ├── reporting.md                                  # Reports & exports reference.
│   ├── MIGRATIONS.md                                 # Schema migration history.
│   ├── SCHEMA.md                                     # Logical schema overview.
│   └── RUNBOOK.md                                    # Ops runbook (alerts, rollbacks, PII access).
│
├── .github/
│   └── workflows/
│       ├── ci.yml                                    # Build, tests, lint.
│       └── cd.yml                                    # Container build & deploy.
│
├── go.mod                                            # Imports pkg/auth, platform-shared, contracts/events.
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md                                         # Quickstart + local run instructions.


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