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
