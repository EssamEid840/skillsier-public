## Event Conventions (applies to all) — communications-be

- **Format:** aggregate.resource.action.past_tense.v1 (e.g., conversation.created.v1, message.sent.v1)
- **Envelope includes:** event_id, event_ts, aggregate_id, partition_key=conversation_id/notification_id, correlation_id, causation_id, actor{ id, role }, user_context{ ip, ua }, data_zone(EU|US), schema_ref, compliance_context{ pii_flags }
- **Batch ops:** Per-entity events + one *.summary.v1
- **PII:** Emit hashes/storage_ids/refs only (no raw PII—e.g., no plaintext emails, phone numbers, or file contents/names)

## Write-path Defaults — communications-be

- **Idempotency:** Header Idempotency-Key (or envelope key). Safe retries must return 200 with no duplicate events.
- **Transactions:** DB tx + outbox with (aggregate_id, event_type, idempotency_key) dedupe.
- **Retries/DLQ:** For external calls (storage-be, users-be, jobs-be, contracts-be, proposals-be, admin-be, WildDuck). Exponential backoff; poison messages to DLQ.
- **Projections:** _read views per domain; metric event_to_projector_lag_ms tracked.
- **Security/Perf:** RBAC enforced on commands/queries; SLO/SLA and rate limits as specified per endpoint/feature; field-level encryption for sensitive data; secrets never logged; typical **write P95 ≤ 300 ms**, **read P95 ≤ 250 ms** (unless noted).





## **========================= 💬 CORE CHAT PRIMITIVES =========================**

### 1) conversation/

#### 1.1 entity (Conversation Core) (Direct & Group Chat)

##### Stories

- As a **freelancer**, I want to create a direct conversation with a client so that we can discuss job opportunities.
- As a **client**, I want to create a group conversation with multiple freelancers so that I can manage team communication.
- As a **system**, I want to track conversation kind (direct, group, system) so that proper rules are enforced.
- As a **system**, I want to track conversation metadata (created_by, visibility, data_zone) so that compliance is maintained.
- As a **system**, I want to assign created_by so that ownership is clear.
- As a **system**, I want to validate conversation membership so that unauthorized access is prevented.
- As a **system**, I want to track visibility (public, private) so that access control works.
- As a **user**, I want to archive conversations so that my active list stays clean.
- As a **user**, I want to unarchive conversations so that I can restore them.
- As a **user**, I want to mute conversations so that I don't receive notifications for specific chats.
- As a **user**, I want to delete conversations so that unwanted chats are removed from my view.
- As a **admin**, I want to view all conversations for moderation so that platform safety is ensured.

##### Flow

1. **CreateConversationCommand**(kind, participants[], title?, visibility, tenant_id, data_zone, created_by) → ValidateParticipants() | CheckTenantQuota() | ValidateKind() | Persist() → **Outbox:** conversation.created.v1
2. **UpdateConversationCommand**(conversation_id, updates, updated_by) → AuthorizeOwner() | Validate() | Update() → **Outbox:** conversation.updated.v1
3. **ArchiveConversationCommand**(conversation_id, user_id) → AuthorizeParticipant() | Archive() → **Outbox:** conversation.archived.v1
4. **UnarchiveConversationCommand**(conversation_id, user_id) → AuthorizeParticipant() | Unarchive() → **Outbox:** conversation.unarchived.v1
5. **DeleteConversationCommand**(conversation_id, user_id, delete_for) → AuthorizeOwner() | SoftDelete() → **Outbox:** conversation.deleted.v1
6. **GetConversationQuery**(conversation_id) → AuthorizeAccess() | Fetch() → ConversationDTO
7. **ListConversationsQuery**(user_id, filters, pagination) → ApplyFilters() | Paginate() → ConversationListDTO
8. **SearchConversationsQuery**(user_id, query, filters) → FullTextSearch() → ConversationSearchDTO

##### Projections

- conversation_read

##### Events Published

- conversation.created.v1
- conversation.updated.v1
- conversation.archived.v1
- conversation.unarchived.v1
- conversation.deleted.v1

##### RBAC/SLO

- **RBAC:** OWNER (create/update/delete), PARTICIPANT (archive/unarchive/view), ADMIN (view all/moderate)
- **SLO:** P95 < 200ms (create), P95 < 150ms (update/archive), P95 < 120ms (read)

---

#### 1.2 participant

##### Stories

- As a **conversation owner**, I want to add participants so that more people can join the conversation.
- As a **conversation owner**, I want to remove participants so that I can manage membership.
- As a **conversation owner**, I want to assign participant roles (owner, admin, member) so that permissions are controlled.
- As a **participant**, I want to leave a conversation so that I stop receiving messages.
- As a **participant**, I want to see who is in a conversation so that I know the audience.
- As a **system**, I want to track last_read_msg_id per participant so that unread counts are accurate.
- As a **participant**, I want to pin/unpin conversations so that important chats stay at the top.
- As a **participant**, I want to mute conversations until a specific time so that I control notifications.

##### Flow

1. **AddParticipantCommand**(conversation_id, user_id, role, added_by) → AuthorizeOwner() | ValidateUser() | CheckLimit() | Add() → **Outbox:** conversation.member_added.v1
2. **RemoveParticipantCommand**(conversation_id, user_id, removed_by) → AuthorizeOwner() | Remove() → **Outbox:** conversation.member_removed.v1
3. **UpdateParticipantRoleCommand**(conversation_id, user_id, new_role, updated_by) → AuthorizeOwner() | UpdateRole() → **Outbox:** participant.role_updated.v1
4. **LeaveConversationCommand**(conversation_id, user_id) → AuthorizeParticipant() | RemoveSelf() → **Outbox:** participant.left.v1
5. **PinConversationCommand**(conversation_id, user_id) → AuthorizeParticipant() | Pin() → **Outbox:** conversation.pinned.v1
6. **UnpinConversationCommand**(conversation_id, user_id) → AuthorizeParticipant() | Unpin() → **Outbox:** conversation.unpinned.v1
7. **MuteConversationCommand**(conversation_id, user_id, muted_until) → AuthorizeParticipant() | Mute() → **Outbox:** conversation.muted.v1
8. **UnmuteConversationCommand**(conversation_id, user_id) → AuthorizeParticipant() | Unmute() → **Outbox:** conversation.unmuted.v1
9. **UpdateLastReadCommand**(conversation_id, user_id, last_read_msg_id) → AuthorizeParticipant() | Update() → **Outbox:** participant.last_read_updated.v1
10. **GetParticipantsQuery**(conversation_id) → AuthorizeAccess() | Fetch() → ParticipantListDTO

##### Projections

- conversation_participants_read

##### Events Published

- conversation.member_added.v1
- conversation.member_removed.v1
- participant.role_updated.v1
- participant.left.v1
- conversation.pinned.v1
- conversation.unpinned.v1
- conversation.muted.v1
- conversation.unmuted.v1
- participant.last_read_updated.v1

##### RBAC/SLO

- **RBAC:** OWNER (add/remove/update role), PARTICIPANT (leave/pin/mute/update last read)
- **SLO:** P95 < 160ms (add/remove), P95 < 120ms (pin/mute/last read)

---

#### 1.3 settings

##### Stories

- As a **conversation owner**, I want to set TTL policies so that messages auto-delete after a period.
- As a **system**, I want to enforce legal_hold so that TTL policies are overridden for compliance.
- As a **conversation owner**, I want to enable slow_mode so that message flooding is prevented.
- As a **conversation owner**, I want to control allow_replies so that reply permissions are managed.
- As a **conversation owner**, I want to control allow_files so that file sharing can be restricted.
- As a **system**, I want to persist ttl_policy_id so that retention rules are applied.

##### Flow

1. **SetTTLPolicyCommand**(conversation_id, ttl_days, set_by) → AuthorizeOwner() | ValidateTTL() | Set() → **Outbox:** ttl_policy.set.v1
2. **ApplyLegalHoldCommand**(conversation_id, admin_id, reason, hold_until) → AuthorizeAdmin() | ApplyHold() | DisableAutoDelete() → **Outbox:** legal_hold.applied.v1
3. **ReleaseLegalHoldCommand**(conversation_id, admin_id) → AuthorizeAdmin() | Release() | EnableAutoDelete() → **Outbox:** legal_hold.released.v1
4. **EnableSlowModeCommand**(conversation_id, delay_seconds, enabled_by) → AuthorizeOwner() | Enable() → **Outbox:** slow_mode.enabled.v1
5. **DisableSlowModeCommand**(conversation_id, disabled_by) → AuthorizeOwner() | Disable() → **Outbox:** slow_mode.disabled.v1
6. **UpdateConversationSettingsCommand**(conversation_id, allow_replies, allow_files, updated_by) → AuthorizeOwner() | Update() → **Outbox:** conversation.settings_updated.v1
7. **GetConversationSettingsQuery**(conversation_id) → AuthorizeAccess() | Fetch() → SettingsDTO

##### Projections

- conversation_settings_read
- legal_holds_read

##### Events Published

- ttl_policy.set.v1
- legal_hold.applied.v1
- legal_hold.released.v1
- slow_mode.enabled.v1
- slow_mode.disabled.v1
- conversation.settings_updated.v1

##### RBAC/SLO

- **RBAC:** OWNER (set TTL/slow mode/permissions), ADMIN (legal hold), PARTICIPANT (view settings)
- **SLO:** P95 < 160ms

---

#### 1.4 typing_indicator

##### Stories

- As a **user**, I want to broadcast typing indicators so that others know I'm composing a message.
- As a **user**, I want to see when others are typing so that I know to wait for their response.
- As a **system**, I want typing indicators to auto-expire (3s) so that stale indicators are cleared.
- As a **system**, I want to store typing TTL markers per conversation/user so that state is managed.

##### Flow

1. **BroadcastTypingCommand**(conversation_id, user_id) → AuthorizeParticipant() | SetTTL(3s) | BroadcastWebSocket() → **Outbox:** typing.started.v1
2. **StopTypingCommand**(conversation_id, user_id) → ClearIndicator() | BroadcastWebSocket() → **Outbox:** typing.stopped.v1
3. **GetTypingIndicatorsQuery**(conversation_id) → AuthorizeParticipant() | Fetch() → TypingIndicatorDTO

##### Projections

- typing_indicators_read (Redis TTL 3s)

##### Events Published

- typing.started.v1
- typing.stopped.v1

##### RBAC/SLO

- **RBAC:** PARTICIPANT (broadcast/stop/view)
- **SLO:** P95 < 80ms (WebSocket broadcast)

---

### 2) thread/

#### 2.1 entity (Thread/Sub-discussions)

##### Stories

- As a **user**, I want to create threads from messages so that sub-discussions are organized.
- As a **system**, I want to track thread id, conversation_id, root_message_id so that threading is correct.
- As a **user**, I want to set thread title so that topics are clear.
- As a **user**, I want to follow threads so that I receive notifications for updates.
- As a **user**, I want to unfollow threads so that I stop receiving notifications.
- As a **user**, I want to archive threads so that completed discussions are hidden.
- As a **user**, I want to rename threads so that titles stay relevant.

##### Flow

1. **CreateThreadCommand**(conversation_id, root_message_id, title?, user_id) → ValidateRootMessage() | AuthorizeParticipant() | Create() → **Outbox:** thread.created.v1
2. **RenameThreadCommand**(thread_id, new_title, user_id) → AuthorizeCreator() | Rename() → **Outbox:** thread.renamed.v1
3. **ArchiveThreadCommand**(thread_id, user_id) → AuthorizeCreator() | Archive() → **Outbox:** thread.archived.v1
4. **FollowThreadCommand**(thread_id, user_id) → AuthorizeParticipant() | AddFollower() → **Outbox:** thread.followed.v1
5. **UnfollowThreadCommand**(thread_id, user_id) → RemoveFollower() → **Outbox:** thread.unfollowed.v1
6. **GetThreadQuery**(thread_id) → AuthorizeAccess() | Fetch() → ThreadDTO
7. **ListThreadsQuery**(conversation_id, pagination) → AuthorizeAccess() | Fetch() → ThreadListDTO

##### Projections

- thread_read
- thread_followers_read

##### Events Published

- thread.created.v1
- thread.renamed.v1
- thread.archived.v1
- thread.followed.v1
- thread.unfollowed.v1

##### RBAC/SLO

- **RBAC:** CREATOR (rename/archive), PARTICIPANT (create/follow/unfollow/view)
- **SLO:** P95 < 180ms (create), P95 < 120ms (follow/rename)

---

### 3) message/

#### 3.1 entity (Message Core)

##### Stories

- As a **user**, I want to send text messages so that I can communicate with other users.
- As a **system**, I want to track message id, conversation_id, sender_id so that messages are identifiable.
- As a **system**, I want to store body (rich text) so that formatting is preserved.
- As a **user**, I want to reply to specific messages (reply_to_id) so that context is clear.
- As a **user**, I want to edit my sent messages so that I can correct mistakes.
- As a **system**, I want to track edited_at so that edits are timestamped.
- As a **user**, I want to delete messages so that I can remove unwanted content.
- As a **system**, I want to track deleted_at and redact_reason so that deletions are auditable.
- As a **system**, I want to validate message content so that profanity and PII are filtered.
- As a **system**, I want to enforce rate limits so that spam is prevented.

##### Flow

1. **SendMessageCommand**(conversation_id, sender_id, body, reply_to_id?, thread_id?) → ValidateParticipant() | AntiPiiLint() | ProfanityCheck() | RateLimitCheck() | Persist() | BroadcastWebSocket() → **Outbox:** message.sent.v1
2. **EditMessageCommand**(message_id, new_body, user_id) → AuthorizeSender() | ValidateEditWindow(15m) | AppendEditHistory() | Update() | BroadcastWebSocket() → **Outbox:** message.edited.v1
3. **DeleteMessageCommand**(message_id, user_id, delete_for, redact_reason?) → AuthorizeSender() | SoftDelete() | BroadcastWebSocket() → **Outbox:** message.deleted.v1
4. **GetMessagesQuery**(conversation_id, pagination, user_id) → AuthorizeAccess() | Fetch() → MessageListDTO
5. **GetMessageQuery**(message_id) → AuthorizeAccess() | Fetch() → MessageDTO
6. **GetThreadMessagesQuery**(thread_id, pagination) → AuthorizeAccess() | Fetch() → MessageListDTO

##### Projections

- message_read
- message_edit_history_read

##### Events Published

- message.sent.v1
- message.edited.v1
- message.deleted.v1

##### Events Consumed

- user.banned.v1 (delete all messages from user)
- admin.message.removed.v1 (moderator deletion)

##### RBAC/SLO

- **RBAC:** SENDER (edit/delete), PARTICIPANT (send/view)
- **SLO:** P95 < 250ms (send), P95 < 150ms (edit/delete), Rate Limit: 30 msgs/min/user

---

#### 3.2 attachment (Rich Media Attachments)

##### Stories

- As a **user**, I want to attach files to messages so that I can share share rich content, documents, images, and videos.
- As a **system**, I want to integrate with storage-be so that files are managed properly.
- As a **system**, I want to store storage-be asset refs (url, id, hash, type, size, thumb) so that files are managed externally.
- As a **system**, I want to track virus_status so that malware is blocked.
- As a **system**, I want to scan attachments for viruses so that security is maintained.
- As a **system**, I want to generate thumbnails for media so that previews are available.

##### Flow

1. **AttachFileCommand**(message_id, file_id, file_metadata) → ValidateFile(storage-be) | ScanVirus() | GenerateThumbnail() | Attach() → **Outbox:** attachment.added.v1
2. **RemoveAttachmentCommand**(attachment_id, user_id) → AuthorizeSender() | Remove() | DeleteFromStorage(storage-be) → **Outbox:** attachment.removed.v1
3. **GetAttachmentQuery**(attachment_id) → AuthorizeAccess() | Fetch() → AttachmentDTO
4. **GetMessageAttachmentsQuery**(message_id) → AuthorizeAccess() | Fetch() → AttachmentListDTO

##### Projections

- message_attachments_read
- attachment_scan_results_read

##### Events Published

- attachment.added.v1
- attachment.removed.v1
- attachment.scanned.v1

##### Events Consumed

- storage.file.uploaded.v1
- storage.file.deleted.v1
- storage.media.processed.v1 (virus scan results)

##### RBAC/SLO

- **RBAC:** SENDER (attach/remove), PARTICIPANT (view)
- **SLO:** P95 < 250ms (attach with virus scan)

---

#### 3.3 read_receipt

##### Stories

- As a **user**, I want to mark messages as read so that unread counts are accurate.
- As a **system**, I want to store message_id, user_id, read_at so that receipts are rollup-friendly.
- As a **user**, I want to see who has read my messages so that I know delivery status.
- As a **system**, I want to track read receipts per message so that analytics are available.
- As a **system**, I want to compact read receipts so that storage is optimized.

##### Flow

1. **MarkAsReadCommand**(message_ids[], user_id, conversation_id) → AuthorizeParticipant() | UpdateReadReceipts() | UpdateUnreadCount() | UpdateLastRead() → **Outbox:** message.read.v1
2. **GetReadReceiptsQuery**(message_id) → AuthorizeParticipant() | Fetch() → ReadReceiptDTO
3. **GetUnreadCountQuery**(user_id, conversation_id?) → Aggregate() → UnreadCountDTO
4. **CompactReadReceiptsCommand**(conversation_id) → Rollup() → Persist()

##### Projections

- message_read_receipts_read
- unread_count_read

##### Events Published

- message.read.v1
- read_receipts.compacted.v1

##### RBAC/SLO

- **RBAC:** PARTICIPANT (mark read/view receipts)
- **SLO:** P95 < 100ms

---

#### 3.4 reaction (Message Reactions & Emojis)

##### Stories

- As a **user**, I want to react to messages with emojis so that I can express quick feedback.
- As a **system**, I want to track emoji, user_id, reacted_at so that reactions are timestamped.
- As a **user**, I want to remove my reactions so that I can change my mind.
- As a **system**, I want to limit reactions per user per message (5) so that abuse is prevented.
- As a **user**, I want to see who reacted with each emoji so that context is clear.

##### Flow

1. **ReactToMessageCommand**(message_id, user_id, emoji) → ValidateEmoji() | CheckLimit(5) | AddReaction() | BroadcastWebSocket() → **Outbox:** message.reacted.v1
2. **RemoveReactionCommand**(message_id, user_id, emoji) → RemoveReaction() | BroadcastWebSocket() → **Outbox:** message.reaction_removed.v1
3. **GetReactionDetailsQuery**(message_id) → AuthorizeParticipant() | Fetch() → ReactionDetailsDTO
4. **GetReactionSummaryQuery**(message_id) → Aggregate() → ReactionSummaryDTO

##### Projections

- message_reactions_read

##### Events Published

- message.reacted.v1
- message.reaction_removed.v1

##### RBAC/SLO

- **RBAC:** PARTICIPANT (add/remove/view)
- **SLO:** P95 < 120ms

---

#### 3.5 mention

##### Stories

- As a **user**, I want to mention other participants in messages so that they are notified.
- As a **system**, I want to track mentioned_user_id and offsets so that mentions are highlighted.
- As a **system**, I want to validate mentioned users are participants so that mentions are valid.
- As a **user**, I want to see all messages where I was mentioned so that I can review context.

##### Flow

1. **ValidateMentionsCommand**(conversation_id, mentioned_user_ids[]) → CheckParticipants() → ValidationResult
2. **ParseMentionsCommand**(message_body) → ExtractMentions() → MentionList
3. **GetMentionedMessagesQuery**(user_id, conversation_id?, pagination) → Fetch() → MentionedMessageListDTO

##### Projections

- message_mentions_read

##### Events Published

- user.mentioned.v1

##### RBAC/SLO

- **RBAC:** PARTICIPANT (mention/view mentions)
- **SLO:** P95 < 120ms

---

### 4) draft/

#### 4.1 entity (Per-user Unsent Drafts)

##### Stories

- As a **user**, I want to save message drafts so that I don't lose work in progress.
- As a **system**, I want to track conversation_id, user_id, content, updated_at so that drafts are user-specific.
- As a **user**, I want to update drafts as I type so that changes are auto-saved.
- As a **user**, I want to delete drafts when I send messages so that drafts are cleaned up.
- As a **user**, I want to list my drafts so that I can resume composing.
- As a **system**, I want to auto-delete old drafts (30d) so that storage is optimized.

##### Flow

1. **SaveDraftCommand**(conversation_id, user_id, content) → AuthorizeParticipant() | Upsert() → **Outbox:** draft.saved.v1
2. **UpdateDraftCommand**(draft_id, content) → AuthorizeOwner() | Update() → **Outbox:** draft.updated.v1
3. **DeleteDraftCommand**(draft_id, user_id) → AuthorizeOwner() | Delete() → **Outbox:** draft.deleted.v1
4. **GetDraftQuery**(conversation_id, user_id) → AuthorizeOwner() | Fetch() → DraftDTO
5. **ListDraftsQuery**(user_id, pagination) → Fetch() → DraftListDTO
6. **CleanupOldDraftsCommand**(cutoff_date) → FetchExpired() | Delete() → **Outbox:** drafts.cleaned_up.v1

##### Projections

- draft_read

##### Events Published

- draft.saved.v1
- draft.updated.v1
- draft.deleted.v1
- drafts.cleaned_up.v1

##### RBAC/SLO

- **RBAC:** OWNER (save/update/delete/view)
- **SLO:** P95 < 100ms (save/update), Background cleanup

---

### 5) pin/

#### 5.1 entity (Pinned Highlights)

##### Stories

- As a **user**, I want to pin messages in a conversation so that important information is highlighted.
- As a **system**, I want to track conversation_id, message_id, pinned_by, pinned_at so that pins are timestamped.
- As a **user**, I want to unpin messages so that pins stay current.
- As a **user**, I want to list pinned messages so that I can quickly access important content.
- As a **system**, I want to enforce pin limits (5 per conversation) so that pins remain meaningful.

##### Flow

1. **PinMessageCommand**(conversation_id, message_id, pinned_by) → AuthorizeParticipant() | CheckPinLimit(5) | Pin() → **Outbox:** message.pinned.v1
2. **UnpinMessageCommand**(conversation_id, message_id, user_id) → AuthorizeParticipant() | Unpin() → **Outbox:** message.unpinned.v1
3. **GetPinnedMessagesQuery**(conversation_id) → AuthorizeAccess() | Fetch() → PinnedMessageListDTO
4. **ReorderPinsCommand**(conversation_id, pin_order[], user_id) → AuthorizeParticipant() | Reorder() → **Outbox:** pins.reordered.v1

##### Projections

- pinned_messages_read

##### Events Published

- message.pinned.v1
- message.unpinned.v1
- pins.reordered.v1

##### RBAC/SLO

- **RBAC:** PARTICIPANT (pin/unpin/reorder/view)
- **SLO:** P95 < 140ms

---

### 6) bookmark/

#### 6.1 entity (Private Bookmarks)

##### Stories

- As a **user**, I want to bookmark messages privately so that I can save important content for myself.
- As a **system**, I want to track user_id, message_id, note, created_at so that bookmarks are personal.
- As a **user**, I want to add notes to bookmarks so that I remember context.
- As a **user**, I want to remove bookmarks so that my list stays relevant.
- As a **user**, I want to list my bookmarks so that I can find saved content.
- As a **user**, I want to search bookmarks by note so that finding content is easy.

##### Flow

1. **BookmarkMessageCommand**(user_id, message_id, note?) → AuthorizeUser() | Create() → **Outbox:** message.bookmarked.v1
2. **UpdateBookmarkCommand**(bookmark_id, note, user_id) → AuthorizeOwner() | Update() → **Outbox:** bookmark.updated.v1
3. **RemoveBookmarkCommand**(bookmark_id, user_id) → AuthorizeOwner() | Delete() → **Outbox:** bookmark.removed.v1
4. **GetBookmarksQuery**(user_id, pagination) → Fetch() → BookmarkListDTO
5. **SearchBookmarksQuery**(user_id, query) → FullTextSearch() → BookmarkSearchDTO

##### Projections

- bookmarks_read

##### Events Published

- message.bookmarked.v1
- bookmark.updated.v1
- bookmark.removed.v1

##### RBAC/SLO

- **RBAC:** OWNER (create/update/delete/view)
- **SLO:** P95 < 120ms

---

### 8.5 Conversation Export

#### Stories

- As a **user**, I want to export conversation history so that I have offline records.
- As a **system**, I want to export in multiple formats (JSON, CSV, PDF) so that flexibility is provided.
- As a **system**, I want to queue exports so that heavy operations don't block.

#### Flow

1. **RequestExportCommand**(conversation_id, user_id, format, date_range) → AuthorizeParticipant() | QueueJob() → **Outbox:** export.requested.v1
2. **ProcessExportCommand**(export_id) → FetchMessages() | FormatOutput() | UploadToStorage(storage-be) | NotifyUser() → **Outbox:** export.completed.v1
3. **GetExportQuery**(export_id) → AuthorizeOwner() | Fetch() → ExportDTO
4. **DownloadExportCommand**(export_id, user_id) → AuthorizeOwner() | GenerateSignedURL(storage-be) → SignedURLDTO

#### Projections

- conversation_exports_read

#### Events Published

- export.requested.v1
- export.completed.v1
- export.failed.v1

#### RBAC/SLO

- **RBAC:** PARTICIPANT (request/download)
- **SLO:** P95 < 300ms (request), Background processing for generation
---





## **========================= 🚚 DELIVERY & READ STATE =========================**

### 7) delivery/

#### 7.1 status (Server→Device Delivery State)

##### Stories

- As a **system**, I want to track delivery status (queued→dispatched→ack) so that message delivery is verifiable.
- As a **system**, I want to store per-device/session acks so that delivery is tracked per endpoint.
- As a **system**, I want to retry failed deliveries so that reliability is ensured.
- As a **system**, I want to track delivery attempts so that debugging is possible.
- As a **user**, I want to see delivery status for my sent messages so that I know they were received.

##### Flow

1. **QueueDeliveryCommand**(message_id, recipient_ids[], devices[]) → Enqueue() → **Outbox:** delivery.queued.v1
2. **DispatchDeliveryCommand**(delivery_id, device_id, session_id) → SendToDevice() | MarkDispatched() → **Outbox:** delivery.dispatched.v1
3. **AcknowledgeDeliveryCommand**(delivery_id, device_id, acked_at) → MarkAcked() → **Outbox:** delivery.acknowledged.v1
4. **RetryFailedDeliveryCommand**(delivery_id) → Retry() | IncrementAttempts() → **Outbox:** delivery.retried.v1
5. **GetDeliveryStatusQuery**(message_id) → Fetch() → DeliveryStatusDTO
6. **GetDeliveryStatsQuery**(date_range) → Aggregate() → DeliveryStatsDTO

##### Projections

- delivery_status_read
- delivery_attempts_read

##### Events Published

- delivery.queued.v1
- delivery.dispatched.v1
- delivery.acknowledged.v1
- delivery.retried.v1
- delivery.failed.v1
- message.delivered.v1

##### RBAC/SLO

- **RBAC:** SENDER (view status), SYSTEM (queue/dispatch/ack)
- **SLO:** P95 < 150ms (queue/dispatch), P95 < 80ms (ack)

---

### 8) read_receipt/ (explicit "I read it")

#### 8.1 entity (Explicit Read Tracking - Compacted)

##### Stories

- As a **system**, I want to store compacted read receipts (message_id, user_id, read_at) so that storage is efficient.
- As a **system**, I want to rollup read receipts periodically so that performance is maintained.
- As a **user**, I want explicit read tracking so that read status is accurate.

##### Flow

1. **CompactReadReceiptsCommand**(conversation_id, cutoff_date) → FetchReceipts() | Rollup() | Persist() → **Outbox:** read_receipts.compacted.v1
2. **GetCompactedReceiptsQuery**(message_id) → Fetch() → CompactedReceiptDTO

##### Projections

- read_receipts_compacted_read

##### Events Published

- read_receipts.compacted.v1

##### RBAC/SLO

- **RBAC:** SYSTEM (compact), PARTICIPANT (view)
- **SLO:** P95 < 200ms (compact), P95 < 100ms (read)

---

## **========================= ⚡ EPHEMERAL REALTIME SIGNALS =========================**

### 9) presence/

#### 9.1 session (Online/Away/Offline)

##### Stories

- As a **user**, I want to show my online/away/offline status so that others know my availability.
- As a **system**, I want to track session_id, user_id, last_seen_at so that presence is session-based.
- As a **system**, I want to track ip, ua, device_kind so that session context is available.
- As a **user**, I want to see other users' presence status so that I know when they're available.
- As a **system**, I want to update presence based on activity so that status is accurate.
- As a **system**, I want presence to auto-expire (60s) so that stale statuses are cleared.
- As a **system**, I want to handle presence heartbeats so that sessions stay alive.

##### Flow

1. **JoinPresenceCommand**(user_id, session_id, device_kind, ip, ua) → RegisterSession() | SetTTL(60s) | BroadcastWebSocket() → **Outbox:** presence.joined.v1
2. **HeartbeatPresenceCommand**(session_id, user_id) → RefreshTTL(60s) | Update() → **Outbox:** presence.heartbeat.v1
3. **LeavePresenceCommand**(session_id, user_id) → UnregisterSession() | BroadcastWebSocket() → **Outbox:** presence.left.v1
4. **GetPresenceQuery**(user_ids[]) → Fetch() → PresenceDTO
5. **GetOnlineUsersQuery**(conversation_id) → Fetch() → OnlineUserListDTO
6. **GetUserSessionsQuery**(user_id) → Fetch() → SessionListDTO

##### Projections

- presence_read (Redis TTL 60s)
- user_sessions_read

##### Events Published

- presence.joined.v1
- presence.left.v1
- presence.heartbeat.v1

##### RBAC/SLO

- **RBAC:** USER (join/heartbeat/leave), PUBLIC (view presence)
- **SLO:** P95 < 70ms

---

### 10) typing/

#### 10.1 signal (Typing Indicators - Short TTL)

##### Stories

- As a **user**, I want to broadcast typing indicators so that others know I'm composing.
- As a **system**, I want to track conversation_id, user_id, started_at, stopped_at so that typing state is managed.
- As a **system**, I want typing signals to have short TTL (3s) so that they auto-expire.
- As a **user**, I want to see when others are typing so that I wait for responses.

##### Flow

1. **StartTypingCommand**(conversation_id, user_id, started_at) → SetTTL(3s) | BroadcastWebSocket() → **Outbox:** typing.started.v1
2. **StopTypingCommand**(conversation_id, user_id, stopped_at) → Clear() | BroadcastWebSocket() → **Outbox:** typing.stopped.v1
3. **GetTypingUsersQuery**(conversation_id) → Fetch() → TypingUserListDTO

##### Projections

- typing_signals_read (Redis TTL 3s)

##### Events Published

- typing.started.v1
- typing.stopped.v1

##### RBAC/SLO

- **RBAC:** PARTICIPANT (start/stop/view)
- **SLO:** P95 < 50ms (WebSocket broadcast)

---

## **========================= 🛡️ SAFETY & COMPLIANCE =========================**

### 11) moderation/

#### 11.1 flag (Message/Conversation Flagging)

##### Stories

- As a **user**, I want to flag inappropriate messages so that moderators can review them.
- As a **user**, I want to flag conversations so that policy violations are addressed.
- As a **system**, I want to track flag_id, resource_type, resource_id, reason, reporter_id so that flags are detailed.
- As a **moderator**, I want to review flagged content so that appropriate action can be taken.
- As a **moderator**, I want to resolve/dismiss flags so that cases are closed.
- As a **system**, I want to auto-flag high-risk content so that immediate action is taken.

##### Flow

1. **FlagMessageCommand**(message_id, user_id, reason, details) → ValidateReason() | Create() | NotifyModerators() → **Outbox:** message.flagged.v1
2. **FlagConversationCommand**(conversation_id, user_id, reason, details) → Create() | NotifyModerators() → **Outbox:** conversation.flagged.v1
3. **ReviewFlagCommand**(flag_id, moderator_id, action, notes) → AuthorizeModerator() | ApplyAction() | Resolve() → **Outbox:** flag.reviewed.v1
4. **DismissFlagCommand**(flag_id, moderator_id, reason) → AuthorizeModerator() | Dismiss() → **Outbox:** flag.dismissed.v1
5. **AutoFlagCommand**(resource_type, resource_id, reason, confidence_score) → Create() | NotifyModerators() → **Outbox:** content.auto_flagged.v1
6. **GetFlaggedContentQuery**(filters, pagination) → AuthorizeModerator() | Fetch() → FlaggedContentListDTO
7. **GetFlagDetailsQuery**(flag_id) → AuthorizeModerator() | Fetch() → FlagDTO

##### Projections

- moderation_flags_read
- moderation_queue_read

##### Events Published

- message.flagged.v1
- conversation.flagged.v1
- flag.reviewed.v1
- flag.dismissed.v1
- content.auto_flagged.v1

##### Events Consumed

- user.banned.v1 (auto-flag all content from user)
- admin.message.removed.v1

##### RBAC/SLO

- **RBAC:** USER (flag), MODERATOR (review/dismiss), SYSTEM (auto-flag)
- **SLO:** P95 < 180ms (flag/review)

---

#### 11.2 action (Moderation Actions)

##### Stories

- As a **moderator**, I want to review reported conversations so that policy violations are addressed.
- As a **moderator**, I want to remove messages so that harmful content is deleted.
- As a **moderator**, I want to quarantine messages so that they're hidden pending review.
- As a **moderator**, I want to freeze conversations so that harmful interactions are stopped.
- As a **moderator**, I want to delete conversations so that severe violations are removed.
- As a **system**, I want to track action_id, moderator_id, action_type, reason so that actions are auditable.
- As a **system**, I want to detect spam patterns in conversations so that automated action is taken.


##### Flow

1. **FlagConversationCommand**(conversation_id, user_id, reason, details) → Create() | NotifyModerators() → **Outbox:** conversation.flagged.v1
2. **ReviewConversationFlagCommand**(flag_id, moderator_id, action, notes) → AuthorizeModerator() | ApplyAction() | Resolve() → **Outbox:** conversation_flag.reviewed.v1
3. **GetFlaggedConversationsQuery**(filters, pagination) → AuthorizeModerator() | Fetch() → FlaggedConversationListDTO
4. **RemoveMessageCommand**(message_id, moderator_id, reason) → AuthorizeModerator() | Remove() | NotifyParties() → **Outbox:** message.removed.v1
5. **QuarantineMessageCommand**(message_id, moderator_id, reason) → Quarantine() | NotifyModerators() → **Outbox:** message.quarantined.v1
6. **ReleaseQuarantineCommand**(message_id, moderator_id) → Release() → **Outbox:** message.released.v1
7. **FreezeConversationCommand**(conversation_id, moderator_id, reason) → Freeze() | NotifyParticipants() → **Outbox:** conversation.frozen.v1
8. **UnfreezeConversationCommand**(conversation_id, moderator_id) → Unfreeze() | NotifyParticipants() → **Outbox:** conversation.unfrozen.v1
9. **DeleteConversationByModeratorCommand**(conversation_id, moderator_id, reason) → HardDelete() | NotifyParticipants() → **Outbox:** conversation.deleted_by_moderator.v1
10. **GetModerationActionsQuery**(filters, pagination) → AuthorizeModerator() | Fetch() → ModerationActionListDTO

##### Projections

- moderation_actions_read
- quarantined_content_read

##### Events Published

- message.removed.v1
- message.quarantined.v1
- message.released.v1
- conversation.frozen.v1
- conversation.unfrozen.v1
- conversation.deleted_by_moderator.v1

##### RBAC/SLO

- **RBAC:** MODERATOR (all actions), ADMIN (view audit)
- **SLO:** P95 < 220ms

---

### 12) spam/

#### 12.1 detector (Spam Detection)

##### Stories

- As a **system**, I want to detect spam messages using ML models so that spam is automatically flagged.
- As a **system**, I want to compute spam scores so that risk is quantified.
- As a **system**, I want to track user spam behavior so that repeat offenders are identified.
- As a **system**, I want to apply automated actions (shadowban, rate limit) so that spam is controlled.
- As a **moderator**, I want to review spam detections so that false positives are corrected.

##### Flow

1. **DetectSpamCommand**(message_id, content) → LoadModel() | ComputeSpamScore() | ApplyThreshold() | ApplyAction() → **Outbox:** spam.detected.v1
2. **UpdateUserSpamScoreCommand**(user_id, increment) → Update() | CheckThreshold() | ApplyAction() → **Outbox:** user.spam_score_updated.v1
3. **ShadowbanUserCommand**(user_id, moderator_id, reason, duration) → AuthorizeModerator() | Shadowban() → **Outbox:** user.shadowbanned.v1
4. **LiftShadowbanCommand**(user_id, moderator_id) → Lift() → **Outbox:** user.shadowban_lifted.v1
5. **GetSpamStatsQuery**(date_range) → AuthorizeModerator() | Aggregate() → SpamStatsDTO
6. **GetUserSpamScoreQuery**(user_id) → AuthorizeModerator() | Fetch() → SpamScoreDTO

##### Projections

- spam_detection_read
- user_spam_scores_read
- shadowbanned_users_read

##### Events Published

- spam.detected.v1
- user.spam_score_updated.v1
- user.shadowbanned.v1
- user.shadowban_lifted.v1

##### RBAC/SLO

- **RBAC:** SYSTEM (detect/update score), MODERATOR (shadowban/lift)
- **SLO:** P95 < 250ms (ML inference), P95 < 150ms (update score)

---

### 13) content_filter/

#### 13.1 rule (Content Filtering Rules)

##### Stories

- As a **system**, I want to filter profanity in messages so that inappropriate language is blocked.
- As a **system**, I want to detect hate speech so that harmful content is prevented.
- As a **system**, I want to redact PII automatically so that sensitive data is protected.
- As a **admin**, I want to configure filter rules so that platform policies are enforced.
- As a **system**, I want to apply filters in real-time so that content is cleaned before delivery.

##### Flow

1. **FilterContentCommand**(content, filter_types[]) → ApplyFilters() | GenerateCleanVersion() → FilterResultDTO
2. **UpdateFilterRulesCommand**(filter_type, rules[], admin_id) → AuthorizeAdmin() | Update() → **Outbox:** filter_rules.updated.v1
3. **GetFilterRulesQuery**(filter_type) → Fetch() → FilterRulesDTO
4. **TestFilterCommand**(content, filter_types[]) → ApplyFilters() → TestResultDTO

##### Projections

- content_filter_rules_read
- filter_violations_read

##### Events Published

- filter_rules.updated.v1
- content.filtered.v1
- pii.redacted.v1

##### RBAC/SLO

- **RBAC:** ADMIN (update rules/test), SYSTEM (filter)
- **SLO:** P95 < 120ms

---

### 14) blocklist/

#### 14.1 entry (User Blocklist)

##### Stories

- As a **user**, I want to block other users so that I don't receive messages from them.
- As a **user**, I want to unblock users so that I can restore communication.
- As a **system**, I want to track blocker_id, blocked_id, scope, reason, expiry so that blocks are managed.
- As a **system**, I want to enforce blocks in conversation creation so that blocked users can't message.
- As a **user**, I want to list blocked users so that I can manage my blocklist.
- As a **system**, I want to expire temporary blocks so that time-limited blocks are automatic.

##### Flow

1. **BlockUserCommand**(blocker_id, blocked_id, scope, reason, expiry?) → Validate() | Create() → **Outbox:** user.blocked.v1
2. **UnblockUserCommand**(blocker_id, blocked_id) → Remove() → **Outbox:** user.unblocked.v1
3. **CheckBlockedCommand**(user_a_id, user_b_id) → CheckBidirectional() → BlockStatusDTO
4. **GetBlocklistQuery**(user_id, pagination) → Fetch() → BlocklistDTO
5. **ExpireBlocksCommand**(cutoff_time) → FetchExpired() | Remove() → **Outbox:** blocks.expired.v1

##### Projections

- blocklist_read

##### Events Published

- user.blocked.v1
- user.unblocked.v1
- blocks.expired.v1

##### RBAC/SLO

- **RBAC:** OWNER (block/unblock/view), SYSTEM (check/expire)
- **SLO:** P95 < 140ms

---

### 15) url_safety/

#### 15.1 cache (URL Safety Cache)

##### Stories

- As a **system**, I want to scan URLs for malicious content so that phishing is prevented.
- As a **system**, I want to cache URL safety results so that repeated checks are fast.
- As a **system**, I want to update cache entries so that stale data is refreshed.
- As a **system**, I want to expire old cache entries so that storage is optimized.
- As a **system**, I want to track url, safety_status, scanned_at, expires_at so that caching works.

##### Flow

1. **ScanURLCommand**(url) → CheckCache() | ScanExternal() | UpdateCache() → **Outbox:** url.scanned.v1
2. **UpdateURLCacheCommand**(url, safety_status, expires_at) → Update() → **Outbox:** url_cache.updated.v1
3. **CheckURLSafetyQuery**(url) → CheckCache() → URLSafetyDTO
4. **ExpireURLCacheCommand**(cutoff_time) → FetchExpired() | Remove() → **Outbox:** url_cache.expired.v1

##### Projections

- url_safety_cache_read (Redis TTL 24h)

##### Events Published

- url.scanned.v1
- url_cache.updated.v1
- url_cache.expired.v1

##### RBAC/SLO

- **RBAC:** SYSTEM (scan/update/expire), PUBLIC (check via system)
- **SLO:** P95 < 150ms (scan), P95 < 50ms (cached check)

---


















## **========================= 📧 NOTIFICATION ORCHESTRATION =========================**

### 16) notification/

#### 16.1 in_app (In-App Notifications)

##### Stories

- As a **user**, I want to receive in-app notifications for important events so that I stay informed.
- As a **user**, I want to mark notifications as read so that I can track what I've seen.
- As a **user**, I want to dismiss notifications so that my notification list stays clean.
- As a **system**, I want to batch similar notifications so that users aren't overwhelmed.
- As a **system**, I want to prioritize notifications (low/normal/high/urgent) so that critical alerts stand out.
- As a **system**, I want to track notification_type, title, content, action_url, related_entity so that notifications are actionable.

##### Flow

1. **CreateInAppNotificationCommand**(user_id, type, title, content, action_url, priority, related_entity) → ValidateUser() | CheckPreferences() | Create() | BroadcastWebSocket() → **Outbox:** notification.in_app.created.v1
2. **MarkNotificationAsReadCommand**(notification_id, user_id) → AuthorizeOwner() | MarkRead() → **Outbox:** notification.read.v1
3. **DismissNotificationCommand**(notification_id, user_id) → AuthorizeOwner() | Dismiss() → **Outbox:** notification.dismissed.v1
4. **BatchCreateNotificationsCommand**(notifications[]) → ValidateBatch(≤100) | CreateBatch() | BroadcastWebSocket() → **Outbox:** notification.batch_created.v1
5. **GetNotificationsQuery**(user_id, filters, pagination) → AuthorizeOwner() | Fetch() → NotificationListDTO
6. **GetUnreadNotificationCountQuery**(user_id) → Aggregate() → UnreadCountDTO

##### Projections

- in_app_notifications_read
- notification_unread_counts_read

##### Events Published

- notification.in_app.created.v1
- notification.read.v1
- notification.dismissed.v1
- notification.batch_created.v1

##### RBAC/SLO

- **RBAC:** OWNER (mark read/dismiss/view), SYSTEM (create)
- **SLO:** P95 < 150ms (create), P95 < 100ms (mark read/dismiss)

---

#### 16.2 preference (User Notification Preferences)

##### Stories

- As a **user**, I want to configure notification preferences per channel (in-app/email/push/sms) so that I control how I'm notified.
- As a **user**, I want to enable/disable specific notification types so that I receive only relevant notifications.
- As a **user**, I want to set quiet hours so that I'm not disturbed during specific times.
- As a **user**, I want to enable email digests so that I receive batched updates.
- As a **user**, I want to mute all notifications temporarily so that I can focus.

##### Flow

1. **UpdateNotificationPreferencesCommand**(user_id, channel_prefs, type_prefs) → Validate() | Update() → **Outbox:** notification_preferences.updated.v1
2. **EnableQuietHoursCommand**(user_id, start_time, end_time, timezone) → Validate() | Enable() → **Outbox:** quiet_hours.enabled.v1
3. **DisableQuietHoursCommand**(user_id) → Disable() → **Outbox:** quiet_hours.disabled.v1
4. **EnableDigestCommand**(user_id, digest_type, schedule) → Validate() | Enable() → **Outbox:** digest.enabled.v1
5. **DisableDigestCommand**(user_id, digest_type) → Disable() → **Outbox:** digest.disabled.v1
6. **MuteAllNotificationsCommand**(user_id, muted_until) → Mute() → **Outbox:** notifications.muted.v1
7. **UnmuteAllNotificationsCommand**(user_id) → Unmute() → **Outbox:** notifications.unmuted.v1
8. **GetNotificationPreferencesQuery**(user_id) → AuthorizeOwner() | Fetch() → NotificationPreferencesDTO

##### Projections

- notification_preferences_read
- quiet_hours_read
- digest_settings_read

##### Events Published

- notification_preferences.updated.v1
- quiet_hours.enabled.v1
- quiet_hours.disabled.v1
- digest.enabled.v1
- digest.disabled.v1
- notifications.muted.v1
- notifications.unmuted.v1

##### RBAC/SLO

- **RBAC:** OWNER (update/view)
- **SLO:** P95 < 140ms

---


#### 16.3 template (Notification Templates)

##### Stories

- As an **admin**, I want to create notification templates so that messages are consistent across channels.
- As an **admin**, I want to version templates so that changes are tracked.
- As a **system**, I want to render templates with dynamic data so that notifications are personalized.
- As an **admin**, I want to test templates before publishing so that quality is ensured.
- As a **system**, I want to track template_id, name, channel, content, variables so that templates are reusable.

##### Flow

1. **CreateTemplateCommand**(name, channel, content, variables, category) → Validate() | Create() → **Outbox:** template.created.v1
2. **UpdateTemplateCommand**(template_id, updates, admin_id) → AuthorizeAdmin() | CreateVersion() | Update() → **Outbox:** template.updated.v1
3. **PublishTemplateCommand**(template_id, version, admin_id) → Validate() | Publish() → **Outbox:** template.published.v1
4. **TestTemplateCommand**(template_id, test_data, admin_id) → Render() | SendTest() → **Outbox:** template.tested.v1
5. **RenderTemplateCommand**(template_id, data) → Fetch() | Render() → RenderedContentDTO
6. **GetTemplateQuery**(template_id, version?) → AuthorizeAdmin() | Fetch() → TemplateDTO
7. **ListTemplatesQuery**(filters, pagination) → AuthorizeAdmin() | Fetch() → TemplateListDTO

##### Projections

- templates_read
- template_versions_read

##### Events Published

- template.created.v1
- template.updated.v1
- template.published.v1
- template.tested.v1

##### RBAC/SLO

- **RBAC:** ADMIN (create/update/publish/test/view), SYSTEM (render)
- **SLO:** P95 < 180ms (create/update), P95 < 100ms (render)

---

### 17) email/

#### 17.1 delivery (Email Delivery via WildDuck)

##### Stories

- As a **user**, I want to receive email notifications for important events so that I stay informed offline.
- As a **system**, I want to use WildDuck (self-hosted SMTP) so that email delivery is controlled.
- As a **system**, I want to track email delivery status (sent/delivered/bounced/opened/clicked) so that analytics are available.
- As a **system**, I want to handle bounces and complaints so that sender reputation is maintained.
- As a **user**, I want to unsubscribe from email categories so that I control what I receive.

##### Flow

1. **SendEmailCommand**(user_id, template_id, template_data, category) → ValidateUser() | CheckPreferences() | CheckUnsubscribe() | RenderTemplate() | SendViaWildDuck() | Persist() → **Outbox:** email.sent.v1
2. **ProcessBounceEventCommand**(email_id, bounce_type, bounce_reason) → UpdateStatus() | HandleBounce() | IncrementBounceCount() → **Outbox:** email.bounced.v1
3. **ProcessComplaintEventCommand**(email_id, complaint_type) → UpdateStatus() | HandleComplaint() | FlagForReview() → **Outbox:** email.complaint.received.v1
4. **UnsubscribeEmailCommand**(user_id, category, reason) → Unsubscribe() | UpdatePreferences() → **Outbox:** email.unsubscribed.v1
5. **GetEmailStatusQuery**(email_id) → Fetch() → EmailStatusDTO
6. **GetEmailDeliveryStatsQuery**(user_id, date_range) → Aggregate() → EmailStatsDTO

##### Projections

- email_delivery_read
- email_bounce_read
- email_unsubscribe_read

##### Events Published

- email.sent.v1
- email.delivered.v1
- email.bounced.v1
- email.complaint.received.v1
- email.unsubscribed.v1

##### Events Consumed

- user.created.v1 (send welcome email)
- contract.created.v1
- proposal.accepted.v1
- payment.processed.v1

##### RBAC/SLO

- **RBAC:** SYSTEM (send), OWNER (unsubscribe), ADMIN (view stats)
- **SLO:** P95 < 400ms (send via SMTP), P95 < 150ms (track events)

---

#### 17.2 bridge (Reply-by-Email)

##### Stories

- As a **user**, I want to reply to notifications via email so that I can respond without opening the app.
- As a **system**, I want to parse inbound emails so that replies are converted to messages.
- As a **system**, I want to validate reply addresses so that spoofing is prevented.
- As a **system**, I want to track inbound_email_id, conversation_id, sender_email so that routing works.

##### Flow

1. **ProcessInboundEmailCommand**(inbound_email_id, from, to, subject, body) → ValidateReplyAddress() | ParseContent() | ExtractConversationID() | CreateMessage() → **Outbox:** email.reply.processed.v1
2. **ValidateReplyAddressCommand**(reply_address) → DecodeToken() | CheckExpiry() | Validate() → ValidationResultDTO
3. **GenerateReplyAddressCommand**(conversation_id, user_id) → EncodeToken() | GenerateAddress() → ReplyAddressDTO

##### Projections

- email_bridge_read

##### Events Published

- email.reply.processed.v1
- email.reply.failed.v1

##### RBAC/SLO

- **RBAC:** SYSTEM (process/validate/generate)
- **SLO:** P95 < 300ms (process inbound)

---

#### 17.3 tracking (Email Provider Logs)

##### Stories

- As a **system**, I want to track email opens so that engagement is measured.
- As a **system**, I want to track link clicks so that CTR is calculated.
- As a **admin**, I want to view email tracking stats so that campaign performance is monitored.
- As a **system**, I want to store email_id, event_type, timestamp, metadata so that tracking is detailed.

##### Flow

1. **ProcessOpenEventCommand**(email_id, opened_at, user_agent, ip) → UpdateStatus() | TrackOpen() → **Outbox:** email.opened.v1
2. **ProcessClickEventCommand**(email_id, clicked_url, clicked_at, user_agent) → UpdateStatus() | TrackClick() → **Outbox:** email.link_clicked.v1
3. **GetEmailTrackingQuery**(email_id) → Fetch() → EmailTrackingDTO
4. **GetEmailEngagementStatsQuery**(date_range) → Aggregate() → EngagementStatsDTO

##### Projections

- email_tracking_read
- email_engagement_read

##### Events Published

- email.opened.v1
- email.link_clicked.v1

##### RBAC/SLO

- **RBAC:** SYSTEM (process), ADMIN (view stats)
- **SLO:** P95 < 120ms

---

### 18) Push Notifications (WebPush, FCM, APNS)

#### 18.1 device (Device Registration)

##### Stories

- As a **user**, I want to register devices for push notifications so that I receive alerts on all my devices.
- As a **user**, I want to unregister devices so that I stop receiving notifications on old devices.
- As a **system**, I want to track device_token, platform (FCM/APNS/WebPush), user_id so that targeting works.
- As a **system**, I want to validate device tokens so that invalid devices are removed.

##### Flow

1. **RegisterDeviceCommand**(user_id, device_token, platform, device_info) → Validate() | Register() → **Outbox:** device.registered.v1
2. **UnregisterDeviceCommand**(user_id, device_token) → Unregister() → **Outbox:** device.unregistered.v1
3. **UpdateDeviceCommand**(device_id, device_info) → Update() → **Outbox:** device.updated.v1
4. **ValidateDeviceTokenCommand**(device_token, platform) → ValidateWithProvider() → ValidationResultDTO
5. **GetRegisteredDevicesQuery**(user_id) → AuthorizeOwner() | Fetch() → DeviceListDTO
6. **CleanupInvalidDevicesCommand**() → FetchInvalid() | Remove() → **Outbox:** devices.cleaned_up.v1

##### Projections

- registered_devices_read

##### Events Published

- device.registered.v1
- device.unregistered.v1
- device.updated.v1
- devices.cleaned_up.v1

##### RBAC/SLO

- **RBAC:** OWNER (register/unregister/view), SYSTEM (validate/cleanup)
- **SLO:** P95 < 180ms

---

#### 18.2 delivery (Push Notification Delivery)

##### Stories

- As a **system**, I want to send push notifications via FCM/APNS/WebPush so that users receive real-time alerts.
- As a **system**, I want to track delivery status so that failed pushes can be retried.
- As a **system**, I want to handle platform-specific formatting so that notifications render correctly.
- As a **user**, I want to receive push notifications on my devices so that I'm alerted immediately.

##### Flow

1. **SendPushNotificationCommand**(user_id, title, body, data, priority, badge_count) → FetchDevices() | FormatForPlatform() | SendToPlatform(FCM/APNS/WebPush) | Track() → **Outbox:** push.sent.v1
2. **ProcessPushDeliveryEventCommand**(push_id, device_token, status, error?) → UpdateStatus() | HandleFailure() → **Outbox:** push.delivered.v1 or push.failed.v1
3. **ProcessPushClickEventCommand**(push_id, device_token, clicked_at) → Track() → **Outbox:** push.clicked.v1
4. **RetryFailedPushCommand**(push_id) → Retry() | IncrementAttempts() → **Outbox:** push.retried.v1
5. **GetPushStatusQuery**(push_id) → Fetch() → PushStatusDTO
6. **GetPushStatsQuery**(user_id, date_range) → Aggregate() → PushStatsDTO

##### Projections

- push_delivery_read
- push_stats_read

##### Events Published

- push.sent.v1
- push.delivered.v1
- push.failed.v1
- push.clicked.v1
- push.retried.v1

##### RBAC/SLO

- **RBAC:** SYSTEM (send/process/retry), OWNER (view stats)
- **SLO:** P95 < 300ms (send push)

---

#### 18.2 delivery (Push Notification Delivery)

##### Stories

- As a **system**, I want to send push notifications via FCM/APNS/WebPush so that users receive real-time alerts.
- As a **system**, I want to track delivery status so that failed pushes can be retried.
- As a **system**, I want to handle platform-specific formatting so that notifications render correctly.
- As a **user**, I want to receive push notifications on my devices so that I'm alerted immediately.

##### Flow

1. **SendPushNotificationCommand**(user_id, title, body, data, priority, badge_count) → FetchDevices() | FormatForPlatform() | SendToPlatform(FCM/APNS/WebPush) | Track() → **Outbox:** push.sent.v1
2. **ProcessPushDeliveryEventCommand**(push_id, device_token, status, error?) → UpdateStatus() | HandleFailure() → **Outbox:** push.delivered.v1 or push.failed.v1
3. **ProcessPushClickEventCommand**(push_id, device_token, clicked_at) → Track() → **Outbox:** push.clicked.v1
4. **RetryFailedPushCommand**(push_id) → Retry() | IncrementAttempts() → **Outbox:** push.retried.v1
5. **GetPushStatusQuery**(push_id) → Fetch() → PushStatusDTO
6. **GetPushStatsQuery**(user_id, date_range) → Aggregate() → PushStatsDTO

##### Projections

- push_delivery_read
- push_stats_read

##### Events Published

- push.sent.v1
- push.delivered.v1
- push.failed.v1
- push.clicked.v1
- push.retried.v1

##### RBAC/SLO

- **RBAC:** SYSTEM (send/process/retry), OWNER (view stats)
- **SLO:** P95 < 300ms (send push)

---



### 19) SMS Notifications (Optional)

#### Stories

- As a **user**, I want to receive SMS notifications for critical events so that I'm alerted even without internet.
- As a **system**, I want to integrate with SMS gateways (Twilio) so that SMS delivery is reliable.
- As a **system**, I want to track SMS delivery status so that analytics are available.
- As a **user**, I want to opt-in/opt-out of SMS notifications so that I control my experience.

#### Flow

1. **SendSMSCommand**(user_id, phone_number, message, category) → ValidatePhone() | CheckPreferences() | CheckOptIn() | SendViaTwilio() → **Outbox:** sms.sent.v1
2. **ProcessSMSDeliveryEventCommand**(sms_id, status) → UpdateStatus() → **Outbox:** sms.delivered.v1 or sms.failed.v1
3. **OptInSMSCommand**(user_id, phone_number) → Verify() | OptIn() → **Outbox:** sms.opt_in.v1
4. **OptOutSMSCommand**(user_id, phone_number) → OptOut() → **Outbox:** sms.opt_out.v1
5. **GetSMSStatusQuery**(sms_id) → Fetch() → SMSStatusDTO

#### Projections

- sms_delivery_read
- sms_opt_in_read

#### Events Published

- sms.sent.v1
- sms.delivered.v1
- sms.failed.v1
- sms.opt_in.v1
- sms.opt_out.v1

#### RBAC/SLO

- **RBAC:** OWNER (opt-in/opt-out), SYSTEM (send)
- **SLO:** P95 < 350ms (send SMS)

---







































































































---

### 20) notification_queue/

#### 20.1 queue (Notification Queueing)

##### Stories

- As a **system**, I want to queue notifications for delivery so that processing is asynchronous.
- As a **system**, I want to prioritize notifications so that urgent alerts are delivered first.
- As a **system**, I want to batch notifications so that delivery is efficient.
- As a **system**, I want to retry failed deliveries so that reliability is ensured.
- As a **system**, I want to track queue_id, notification_id, status, priority, attempts so that queue management works.

##### Flow

1. **EnqueueNotificationCommand**(notification_id, channel, priority, scheduled_for?) → Validate() | Enqueue() → **Outbox:** notification.queued.v1
2. **DequeueNotificationCommand**(channel) → FetchNext() | MarkProcessing() → NotificationDTO
3. **MarkNotificationProcessedCommand**(queue_id, status, delivered_at) → MarkProcessed() → **Outbox:** notification.processed.v1
4. **RetryFailedNotificationCommand**(queue_id) → IncrementAttempts() | Requeue() → **Outbox:** notification.retry.queued.v1
5. **MoveToDLQCommand**(queue_id, reason) → MoveToDLQ() → **Outbox:** notification.moved_to_dlq.v1
6. **GetQueueStatsQuery**(channel) → Aggregate() → QueueStatsDTO
7. **GetDLQItemsQuery**(filters, pagination) → Fetch() → DLQItemListDTO

##### Projections

- notification_queue_read
- notification_dlq_read
- queue_stats_read

##### Events Published

- notification.queued.v1
- notification.processed.v1
- notification.retry.queued.v1
- notification.moved_to_dlq.v1

##### RBAC/SLO

- **RBAC:** SYSTEM (enqueue/dequeue/process/retry/dlq)
- **SLO:** P95 < 100ms (enqueue), P95 < 50ms (dequeue)

---

### 21) suppression/

#### 21.1 list (Suppression List Management)

##### Stories

- As a **system**, I want to maintain suppression lists so that users who opted out don't receive notifications.
- As a **system**, I want to check suppression before sending so that opt-outs are respected.
- As a **user**, I want to be added to suppression list when I unsubscribe so that I stop receiving notifications.
- As a **admin**, I want to manage suppression lists so that entries can be reviewed and removed.
- As a **system**, I want to track user_id, channel, reason, suppressed_at so that suppressions are auditable.

##### Flow

1. **AddToSuppressionListCommand**(user_id, channel, reason) → Validate() | Add() → **Outbox:** suppression.added.v1
2. **RemoveFromSuppressionListCommand**(user_id, channel, admin_id) → AuthorizeAdmin() | Remove() → **Outbox:** suppression.removed.v1
3. **CheckSuppressionCommand**(user_id, channel) → Check() → SuppressionStatusDTO
4. **BulkAddToSuppressionCommand**(user_ids[], channel, reason) → ValidateBatch() | AddBatch() → **Outbox:** suppression.bulk_added.v1
5. **GetSuppressionListQuery**(channel, pagination) → AuthorizeAdmin() | Fetch() → SuppressionListDTO
6. **GetUserSuppressionStatusQuery**(user_id) → AuthorizeOwner() | Fetch() → UserSuppressionStatusDTO

##### Projections

- suppression_list_read

##### Events Published

- suppression.added.v1
- suppression.removed.v1
- suppression.bulk_added.v1

##### RBAC/SLO

- **RBAC:** SYSTEM (add/check), ADMIN (remove/view list), OWNER (view own status)
- **SLO:** P95 < 80ms (check), P95 < 150ms (add/remove)

---

### 22) system_message/

#### 22.1 feed (System Message Feed)

##### Stories

- As a **system**, I want to create system messages so that platform announcements are sent.
- As a **system**, I want to broadcast system messages to all users or specific segments so that important info is communicated.
- As a **user**, I want to view system messages so that I stay informed about platform updates.
- As a **user**, I want to filter system messages by category so that I see relevant announcements.
- As a **system**, I want to track message_id, category, title, body, target_audience, expires_at so that system messages are managed.

##### Flow

1. **CreateSystemMessageCommand**(category, title, body, target_audience, priority, expires_at, created_by) → AuthorizeAdmin() | Validate() | Create() | BroadcastToTargets() → **Outbox:** system_message.created.v1
2. **UpdateSystemMessageCommand**(message_id, updates, updated_by) → AuthorizeAdmin() | Update() → **Outbox:** system_message.updated.v1
3. **ExpireSystemMessageCommand**(message_id, admin_id) → AuthorizeAdmin() | Expire() → **Outbox:** system_message.expired.v1
4. **BroadcastSystemMessageCommand**(message_id) → FetchTargetUsers() | CreateNotifications() → **Outbox:** system_message.broadcasted.v1
5. **GetSystemMessagesQuery**(user_id, filters, pagination) → ApplyFilters() | Fetch() → SystemMessageListDTO
6. **GetSystemMessageQuery**(message_id) → Fetch() → SystemMessageDTO
7. **CleanupExpiredMessagesCommand**(cutoff_date) → FetchExpired() | Remove() → **Outbox:** system_messages.cleaned_up.v1

##### Projections

- system_message_feed_read

##### Events Published

- system_message.created.v1
- system_message.updated.v1
- system_message.expired.v1
- system_message.broadcasted.v1
- system_messages.cleaned_up.v1

##### RBAC/SLO

- **RBAC:** ADMIN (create/update/expire/broadcast), PUBLIC (view)
- **SLO:** P95 < 200ms (create), P95 < 150ms (read), Background broadcast

---

## **========================= 🔗 REAL-TIME COMMUNICATION DOMAIN - REALTIME CHANNELS =========================**

### 23) websocket/

#### 23.1 connection (WebSocket Connection Management)

##### Stories

- As a **user**, I want to establish WebSocket connections so that I receive real-time updates.
- As a **system**, I want to authenticate WebSocket connections so that unauthorized access is prevented.
- As a **system**, I want to handle WebSocket reconnections so that connections are resilient.
- As a **system**, I want to track connection_id, user_id, session_id, connected_at so that connections are managed.
- As a **system**, I want to broadcast messages to specific users/rooms so that real-time delivery is efficient.

##### Flow

1. **EstablishWebSocketCommand**(user_id, auth_token, connection_id, session_id) → ValidateToken() | RegisterConnection() | JoinUserRoom() → **Outbox:** websocket.connected.v1
2. **CloseWebSocketCommand**(connection_id, user_id) → UnregisterConnection() | LeaveRooms() → **Outbox:** websocket.disconnected.v1
3. **JoinRoomCommand**(connection_id, room_id) → AddToRoom() → **Outbox:** websocket.room_joined.v1
4. **LeaveRoomCommand**(connection_id, room_id) → RemoveFromRoom() → **Outbox:** websocket.room_left.v1
5. **BroadcastToRoomCommand**(room_id, event_type, payload) → FetchConnections() | BroadcastToSockets() → **Outbox:** websocket.broadcast.sent.v1
6. **BroadcastToUserCommand**(user_id, event_type, payload) → FetchConnections() | BroadcastToSockets() → **Outbox:** websocket.broadcast.sent.v1
7. **GetActiveConnectionsQuery**(user_id) → Fetch() → ConnectionListDTO
8. **GetRoomMembersQuery**(room_id) → Fetch() → RoomMemberListDTO

##### Projections

- active_websocket_connections_read (Redis TTL)
- websocket_rooms_read

##### Events Published

- websocket.connected.v1
- websocket.disconnected.v1
- websocket.room_joined.v1
- websocket.room_left.v1
- websocket.broadcast.sent.v1

##### RBAC/SLO

- **RBAC:** USER (connect/disconnect), SYSTEM (broadcast/manage rooms)
- **SLO:** P95 < 50ms (broadcast latency), Connection capacity: 10k concurrent per instance

---

### 24) sse/

#### 24.1 stream (Server-Sent Events Stream)

##### Stories

- As a **user**, I want to subscribe to SSE streams so that I receive updates without WebSocket complexity.
- As a **system**, I want to push events to SSE clients so that lightweight real-time updates are supported.
- As a **user**, I want to filter SSE events by type so that I receive only relevant updates.
- As a **system**, I want to track stream_id, user_id, event_types[], subscribed_at so that streams are managed.

##### Flow

1. **SubscribeSSECommand**(user_id, auth_token, event_types[], stream_id) → ValidateToken() | RegisterStream() → **Outbox:** sse.subscribed.v1
2. **UnsubscribeSSECommand**(user_id, stream_id) → UnregisterStream() → **Outbox:** sse.unsubscribed.v1
3. **PushSSEEventCommand**(user_id, event_type, payload) → FetchStreams() | FilterByEventType() | PushToStreams() → **Outbox:** sse.pushed.v1
4. **UpdateSSEFiltersCommand**(stream_id, event_types[]) → Update() → **Outbox:** sse.filters_updated.v1
5. **GetActiveSSEStreamsQuery**(user_id) → Fetch() → SSEStreamListDTO

##### Projections

- active_sse_streams_read (Redis TTL)

##### Events Published

- sse.subscribed.v1
- sse.unsubscribed.v1
- sse.pushed.v1
- sse.filters_updated.v1

##### RBAC/SLO

- **RBAC:** USER (subscribe/unsubscribe/update filters), SYSTEM (push)
- **SLO:** P95 < 80ms (push latency)

---

## **========================= 📞 VOICE/VIDEO CALLS =========================**

### 25) call/

#### 25.1 session (Call Session Management)

##### Stories

- As a **user**, I want to initiate voice/video calls so that I can have real-time conversations.
- As a **user**, I want to accept/reject call invitations so that I control when I join calls.
- As a **system**, I want to exchange WebRTC signaling so that peer connections are established.
- As a **system**, I want to track call_id, caller_id, recipient_id, call_type, status so that calls are managed.
- As a **system**, I want to track call duration and quality so that analytics are available.

##### Flow

1. **InitiateCallCommand**(caller_id, recipient_id, call_type, offer_sdp) → ValidateParticipants() | CreateCall() | SendInvitation() | BroadcastWebSocket() → **Outbox:** call.initiated.v1
2. **AcceptCallCommand**(call_id, recipient_id, answer_sdp) → ValidateCall() | UpdateStatus() | BroadcastWebSocket() → **Outbox:** call.accepted.v1
3. **RejectCallCommand**(call_id, recipient_id, reason) → UpdateStatus() | BroadcastWebSocket() → **Outbox:** call.rejected.v1
4. **ExchangeICECandidateCommand**(call_id, user_id, candidate) → ValidateCall() | BroadcastWebSocket() → **Outbox:** call.ice_candidate.exchanged.v1
5. **EndCallCommand**(call_id, user_id, end_reason) → UpdateStatus() | RecordDuration() | BroadcastWebSocket() → **Outbox:** call.ended.v1
6. **RecordCallQualityCommand**(call_id, quality_metrics) → Persist() → **Outbox:** call.quality_recorded.v1
7. **GetCallQuery**(call_id) → AuthorizeParticipant() | Fetch() → CallDTO
8. **GetCallHistoryQuery**(user_id, pagination) → Fetch() → CallHistoryDTO
9. **GetCallStatsQuery**(call_id) → Aggregate() → CallStatsDTO

##### Projections

- active_calls_read
- call_history_read
- call_stats_read

##### Events Published

- call.initiated.v1
- call.accepted.v1
- call.rejected.v1
- call.ice_candidate.exchanged.v1
- call.ended.v1
- call.missed.v1
- call.quality_recorded.v1

##### RBAC/SLO

- **RBAC:** PARTICIPANT (initiate/accept/reject/end/exchange ICE), OWNER (view history)
- **SLO:** P95 < 100ms (signaling latency)

---

### 26) calendar_invite/

#### 26.1 invite (Calendar Invite Management)

##### Stories

- As a **user**, I want to send calendar invites for scheduled calls so that meetings are organized.
- As a **user**, I want to update calendar invites when schedule changes so that participants are informed.
- As a **user**, I want to accept/decline calendar invites so that my availability is tracked.
- As a **system**, I want to track invite_id, call_id, organizer_id, invitees[], scheduled_at so that invites are managed.
- As a **system**, I want to send reminders before scheduled calls so that participants don't miss meetings.

##### Flow

1. **SendCalendarInviteCommand**(call_id, organizer_id, invitees[], scheduled_at, duration, title, description) → ValidateSchedule() | Create() | SendNotifications() → **Outbox:** calendar_invite.sent.v1
2. **UpdateCalendarInviteCommand**(invite_id, updates, updated_by) → AuthorizeOrganizer() | Update() | NotifyInvitees() → **Outbox:** calendar_invite.updated.v1
3. **CancelCalendarInviteCommand**(invite_id, reason, cancelled_by) → AuthorizeOrganizer() | Cancel() | NotifyInvitees() → **Outbox:** calendar_invite.cancelled.v1
4. **RSVPCalendarInviteCommand**(invite_id, user_id, response) → ValidateInvitee() | UpdateRSVP() | NotifyOrganizer() → **Outbox:** calendar_invite.rsvp.v1
5. **SendReminderCommand**(invite_id) → FetchInvitees() | SendReminders() → **Outbox:** calendar_invite.reminder.sent.v1
6. **GetCalendarInviteQuery**(invite_id) → AuthorizeAccess() | Fetch() → CalendarInviteDTO
7. **ListCalendarInvitesQuery**(user_id, filters, pagination) → Fetch() → CalendarInviteListDTO
8. **GetUpcomingCallsQuery**(user_id, date_range) → Fetch() → UpcomingCallsDTO

##### Projections

- calendar_invites_read
- upcoming_calls_read

##### Events Published

- calendar_invite.sent.v1
- calendar_invite.updated.v1
- calendar_invite.cancelled.v1
- calendar_invite.rsvp.v1
- calendar_invite.reminder.sent.v1

##### RBAC/SLO

- **RBAC:** ORGANIZER (send/update/cancel), INVITEE (rsvp/view), SYSTEM (send reminders)
- **SLO:** P95 < 200ms (send/update), P95 < 150ms (rsvp), Background reminders

---

### 27) interview/

#### 27.1 schedule (Interview Scheduling)

##### Stories

- As a **client**, I want to schedule interviews with freelancers so that hiring discussions are organized.
- As a **freelancer**, I want to confirm interview times so that availability is aligned.
- As a **client/freelancer**, I want to reschedule interviews so that conflicts are managed.
- As a **client/freelancer**, I want to cancel interviews so that time isn't wasted.
- As a **system**, I want to track interview_id, job_id, client_id, freelancer_id, scheduled_at so that interviews are managed.
- As a **system**, I want to mark interviews complete so that hiring progress is tracked.

##### Flow

1. **ScheduleInterviewCommand**(job_id, client_id, freelancer_id, scheduled_at, duration, interview_type, notes) → ValidateParticipants() | CheckAvailability() | Create() | SendNotifications() → **Outbox:** interview.scheduled.v1
2. **ConfirmInterviewCommand**(interview_id, user_id) → AuthorizeParticipant() | Confirm() | NotifyOtherParty() → **Outbox:** interview.confirmed.v1
3. **RescheduleInterviewCommand**(interview_id, new_scheduled_at, rescheduled_by, reason) → AuthorizeParticipant() | Reschedule() | NotifyParticipants() → **Outbox:** interview.rescheduled.v1
4. **CancelInterviewCommand**(interview_id, cancelled_by, reason) → AuthorizeParticipant() | Cancel() | NotifyParticipants() → **Outbox:** interview.cancelled.v1
5. **CompleteInterviewCommand**(interview_id, completed_by, outcome, notes) → AuthorizeParticipant() | Complete() → **Outbox:** interview.completed.v1
6. **SendInterviewReminderCommand**(interview_id) → FetchParticipants() | SendReminders() → **Outbox:** interview.reminder.sent.v1
7. **GetInterviewQuery**(interview_id) → AuthorizeAccess() | Fetch() → InterviewDTO
8. **ListInterviewsQuery**(user_id, filters, pagination) → Fetch() → InterviewListDTO

##### Projections

- interview_schedule_read
- upcoming_interviews_read

##### Events Published

- interview.scheduled.v1
- interview.confirmed.v1
- interview.rescheduled.v1
- interview.cancelled.v1
- interview.completed.v1
- interview.reminder.sent.v1

##### RBAC/SLO

- **RBAC:** CLIENT/FREELANCER (schedule/confirm/reschedule/cancel/complete), SYSTEM (send reminders)
- **SLO:** P95 < 220ms (schedule/reschedule), P95 < 150ms (confirm/cancel), Background reminders

---

## **========================= 🔔 PLATFORM ALERTS =========================**

### 28) platform_alert/

#### 28.1 alert (Platform-wide Alerts)

##### Stories

- As an **admin**, I want to send platform-wide alerts so that critical information reaches all users.
- As a **system**, I want to broadcast alerts for system maintenance so that users are informed.
- As a **user**, I want to dismiss alerts so that my interface stays clean.
- As a **admin**, I want to target alerts to specific user segments so that messaging is relevant.
- As a **system**, I want to track alert_id, alert_type, severity, title, body, target_segment so that alerts are managed.

##### Flow

1. **SendPlatformAlertCommand**(alert_type, severity, title, body, target_segment, expires_at, sent_by) → AuthorizeAdmin() | Validate() | Create() | BroadcastToTargets() → **Outbox:** platform_alert.sent.v1
2. **DismissAlertCommand**(alert_id, user_id) → AuthorizeUser() | MarkDismissed() → **Outbox:** platform_alert.dismissed.v1
3. **ExpireAlertCommand**(alert_id, admin_id) → AuthorizeAdmin() | Expire() → **Outbox:** platform_alert.expired.v1
4. **GetActiveAlertsQuery**(user_id) → FilterBySegment() | FetchActive() → AlertListDTO
5. **GetAlertQuery**(alert_id) → Fetch() → AlertDTO
6. **ListAllAlertsQuery**(filters, pagination) → AuthorizeAdmin() | Fetch() → AlertListDTO

##### Projections

- platform_alerts_read
- active_alerts_read

##### Events Published

- platform_alert.sent.v1
- platform_alert.dismissed.v1
- platform_alert.expired.v1

##### RBAC/SLO

- **RBAC:** ADMIN (send/expire/view all), USER (dismiss/view active)
- **SLO:** P95 < 180ms (send), P95 < 100ms (dismiss), Background broadcast

---

## **========================= 📊 ANALYTICS & AUDIT =========================**

### 29) analytics/

#### 29.1 message_stats (Message Analytics)

##### Stories

- As a **user**, I want to view my message statistics (sent/received/read rates) so that I understand my communication patterns.
- As an **admin**, I want to view platform-wide message metrics so that system health is monitored.
- As a **system**, I want to track messages sent/received/read rates so that issues are detected.
- As a **system**, I want to aggregate stats by time period so that trends are visible.

##### Flow

1. **GetUserMessageStatsQuery**(user_id, date_range) → AuthorizeOwner() | Aggregate() → MessageStatsDTO
2. **GetConversationStatsQuery**(conversation_id, date_range) → AuthorizeParticipant() | Aggregate() → ConversationStatsDTO
3. **GetPlatformMessageStatsQuery**(date_range) → AuthorizeAdmin() | Aggregate() → PlatformMessageStatsDTO
4. **GetMessageDeliveryRatesQuery**(date_range) → AuthorizeAdmin() | Aggregate() → DeliveryRatesDTO
5. **RefreshStatsCommand**(entity_type, entity_id) → Recompute() → **Outbox:** stats.refreshed.v1

##### Projections

- message_stats_read
- conversation_stats_read
- platform_stats_read

##### Events Published

- stats.refreshed.v1

##### RBAC/SLO

- **RBAC:** OWNER (user stats), PARTICIPANT (conversation stats), ADMIN (platform stats)
- **SLO:** P95 < 300ms (aggregations)

---

#### 29.2 notification_stats (Notification Analytics)

##### Stories

- As a **user**, I want to view notification delivery statistics so that I know what I've received.
- As an **admin**, I want to track notification delivery rates by channel so that delivery is optimized.
- As an **admin**, I want to monitor notification engagement (open rates, click rates) so that effectiveness is measured.

##### Flow

1. **GetUserNotificationStatsQuery**(user_id, date_range) → AuthorizeOwner() | Aggregate() → NotificationStatsDTO
2. **GetNotificationDeliveryStatsQuery**(date_range, channel) → AuthorizeAdmin() | Aggregate() → DeliveryStatsDTO
3. **GetNotificationEngagementQuery**(date_range) → AuthorizeAdmin() | Aggregate() → EngagementStatsDTO
4. **GetChannelPerformanceQuery**(channel, date_range) → AuthorizeAdmin() | Aggregate() → ChannelPerformanceDTO

##### Projections

- notification_stats_read
- notification_engagement_read
- channel_performance_read

##### Events Published

- None (read-only analytics)

##### RBAC/SLO

- **RBAC:** OWNER (user stats), ADMIN (platform stats)
- **SLO:** P95 < 280ms

---

### 30) audit/

#### 30.1 log (Audit Trail)

##### Stories

- As a **compliance officer**, I want comprehensive audit logs so that all actions are traceable.
- As a **system**, I want to log all sensitive operations so that security is maintained.
- As a **admin**, I want to search audit logs so that investigations are possible.
- As a **system**, I want to track actor, action, resource, timestamp, metadata so that logs are detailed.

##### Flow

1. **LogAuditEventCommand**(actor_id, actor_role, action_type, resource_type, resource_id, metadata) → Persist() → **Outbox:** audit.event.logged.v1
2. **SearchAuditLogsQuery**(filters, pagination) → AuthorizeComplianceOfficer() | Search() → AuditLogListDTO
3. **GetAuditLogQuery**(log_id) → AuthorizeComplianceOfficer() | Fetch() → AuditLogDTO
4. **GetUserAuditHistoryQuery**(user_id, date_range) → AuthorizeComplianceOfficer() | Fetch() → UserAuditHistoryDTO
5. **ExportAuditLogsCommand**(filters, format) → AuthorizeComplianceOfficer() | Export() | UploadToStorage() → **Outbox:** audit.logs.exported.v1

##### Projections

- audit_log_read

##### Events Published

- audit.event.logged.v1
- audit.logs.exported.v1

##### RBAC/SLO

- **RBAC:** SYSTEM (log), COMPLIANCE_OFFICER/ADMIN (search/view/export)
- **SLO:** P95 < 100ms (log), P95 < 400ms (search)

---

## **31 - DIGEST & BATCH PROCESSING DOMAIN**

### 31.1 Email Digests

#### Stories

- As a **user**, I want to receive daily/weekly email digests so that I get batched updates instead of individual emails.
- As a **system**, I want to aggregate digest content so that summaries are meaningful.
- As a **system**, I want to schedule digest delivery so that users receive them at optimal times.

#### Flow

1. **ScheduleDigestCommand**(user_id, digest_type, schedule) → Validate() | Schedule() → **Outbox:** digest.scheduled.v1
2. **GenerateDigestCommand**(user_id, digest_type, date_range) → AggregateContent() | RenderTemplate() | SendEmail() → **Outbox:** digest.sent.v1
3. **CancelDigestCommand**(user_id, digest_type) → Cancel() → **Outbox:** digest.cancelled.v1
4. **GetDigestPreviewQuery**(user_id, digest_type) → Generate() → DigestPreviewDTO

#### Projections

- digest_schedules_read
- digest_content_read

#### Events Published

- digest.scheduled.v1
- digest.sent.v1
- digest.cancelled.v1

#### RBAC/SLO

- **RBAC:** OWNER (schedule/cancel), SYSTEM (generate)
- **SLO:** P95 < 500ms (generate digest)

---

### 31.2 Notification Batching

#### Stories

- As a **system**, I want to batch similar notifications so that users aren't overwhelmed.
- As a **system**, I want to apply frequency limits so that notification fatigue is prevented.
- As a **system**, I want to group notifications by entity so that summaries are clear.

#### Flow

1. **BatchNotificationsCommand**(user_id, notifications[], grouping_key) → ValidateBatch() | Group() | CreateSummary() → **Outbox:** notification.batch_created.v1
2. **ApplyFrequencyLimitCommand**(user_id, notification_type) → CheckLastSent() | ApplyDelay() → **Outbox:** notification.delayed.v1
3. **GetBatchedNotificationsQuery**(user_id, date_range) → Fetch() → BatchedNotificationListDTO

#### Projections

- batched_notifications_read
- frequency_limits_read

#### Events Published

- notification.batch_created.v1
- notification.delayed.v1

#### RBAC/SLO

- **RBAC:** SYSTEM (batch/apply limits)
- **SLO:** P95 < 180ms

---


## **9 - ADVANCED FEATURES DOMAIN**

### 9.1 Message Search & Indexing

#### Stories

- As a **user**, I want to search messages by content so that I can find past conversations.
- As a **user**, I want to filter searches by date, sender, conversation so that results are precise.
- As a **system**, I want to index messages in Elasticsearch so that search is fast.
- As a **system**, I want to exclude PII from search indexes so that compliance is maintained.

#### Flow

1. **IndexMessageCommand**(message_id) → FetchMessage() | RedactPII() | IndexInElasticsearch() → **Outbox:** message.indexed.v1
2. **SearchMessagesQuery**(user_id, query, filters, pagination) → AuthorizeUser() | SearchElasticsearch() → MessageSearchResultsDTO
3. **ReindexMessagesCommand**(conversation_id) → FetchAll() | BulkIndex() → **Outbox:** messages.reindexed.v1

#### Projections

- message_search_index (Elasticsearch)

#### Events Published

- message.indexed.v1
- messages.reindexed.v1

#### RBAC/SLO

- **RBAC:** PARTICIPANT (search), SYSTEM (index), ADMIN (reindex)
- **SLO:** P95 < 200ms (search), P95 < 100ms (index)

---

























































































## **========================= ⚙️ SYSTEM UTILITIES =========================**

### 31) idempotency/

#### 31.1 key (Idempotency Key Management)

##### Stories

- As a **system**, I want to track idempotency keys so that duplicate requests are prevented.
- As a **system**, I want to cache responses for idempotent requests so that retries are fast.
- As a **system**, I want to expire old idempotency keys so that storage is optimized.
- As a **system**, I want to track key, request_hash, response, expires_at so that deduplication works.

##### Flow

1. **CheckIdempotencyKeyCommand**(key, request_hash) → CheckCache() → IdempotencyResultDTO
2. **StoreIdempotencyKeyCommand**(key, request_hash, response, ttl) → SetTTL() | Store() → **Outbox:** idempotency.key.stored.v1
3. **CleanupExpiredKeysCommand**(cutoff_time) → FetchExpired() | Remove() → **Outbox:** idempotency.keys.cleaned_up.v1

##### Projections

- idempotency_keys_read (Redis TTL 24h for messages, 7d for notifications)

##### Events Published

- idempotency.key.stored.v1
- idempotency.keys.cleaned_up.v1

##### RBAC/SLO

- **RBAC:** SYSTEM (check/store/cleanup)
- **SLO:** P95 < 50ms (check), Background cleanup

---

### 32) quota/

#### 32.1 limit (Rate Limits & Quotas)

##### Stories

- As a **system**, I want to enforce rate limits so that abuse is prevented.
- As a **system**, I want to track quota usage so that limits are respected.
- As a **admin**, I want to configure quota limits per user/tenant so that usage is controlled.
- As a **user**, I want to view my quota usage so that I know my limits.
- As a **system**, I want to track resource_type, user_id, limit, current_usage, window so that quotas work.

##### Flow

1. **CheckQuotaCommand**(user_id, resource_type) → CheckUsage() | CheckLimit() → QuotaStatusDTO
2. **IncrementQuotaCommand**(user_id, resource_type, amount) → Increment() | CheckThreshold() → **Outbox:** quota.incremented.v1
3. **ResetQuotaCommand**(user_id, resource_type) → Reset() → **Outbox:** quota.reset.v1
4. **SetQuotaLimitCommand**(user_id, resource_type, limit, admin_id) → AuthorizeAdmin() | SetLimit() → **Outbox:** quota.limit.set.v1
5. **GetQuotaUsageQuery**(user_2. **UnregisterDeviceCommand**(user_id, device_token) → Unregister() → **Outbox:** device.unregistered.v1
3. **UpdateDeviceCommand**(device_id, device_info) → Update() → **Outbox:** device.updated.v1
4. **ValidateDeviceTokenCommand**(device_token, platform) → ValidateWithProvider() → ValidationResultDTO
5. **GetRegisteredDevicesQuery**(user_id) → AuthorizeOwner() | Fetch() → DeviceListDTO
6. **CleanupInvalidDevicesCommand**() → FetchInvalid() | Remove() → **Outbox:** devices.cleaned_up.v1

##### Projections

- registered_devices_read

##### Events Published

- device.registered.v1
- device.unregistered.v1
- device.updated.v1
- devices.cleaned_up.v1

##### RBAC/SLO

- **RBAC:** OWNER (register/unregister/view), SYSTEM (validate/cleanup)
- **SLO:** P95 < 180ms

---

























































## **9 - INTEGRATION & WEBHOOK DOMAIN**

### 9.1 Webhook Subscriptions

#### Stories

- As a **developer**, I want to subscribe webhooks to communication events so that external systems are notified.
- As an **admin**, I want to view webhook delivery logs so that debugging is possible.
- As a **system**, I want to retry failed webhook deliveries so that reliability is ensured.

#### Flow

1. **SubscribeWebhookCommand**(url, events[], secret) → ValidateURL() | ValidateEvents() | Subscribe() → **Outbox:** webhook.subscribed.v1
2. **UnsubscribeWebhookCommand**(webhook_id) → Unsubscribe() → **Outbox:** webhook.unsubscribed.v1
3. **DeliverWebhookCommand**(webhook_id, event) → SignPayload() | SendHTTP() | LogDelivery() → **Outbox:** webhook.delivered.v1
4. **RetryFailedWebhookCommand**(delivery_id) → Retry() → **Outbox:** webhook.retried.v1
5. **GetWebhookLogsQuery**(webhook_id, pagination) → AuthorizeOwner() | Fetch() → WebhookLogListDTO

#### Projections

- webhook_subscriptions_read
- webhook_delivery_logs_read

#### Events Published

- webhook.subscribed.v1
- webhook.unsubscribed.v1
- webhook.delivered.v1
- webhook.failed.v1
- webhook.retried.v1

#### RBAC/SLO

- **RBAC:** ADMIN (subscribe/unsubscribe/view logs), SYSTEM (deliver)
- **SLO:** P95 < 300ms (delivery), Max retries: 5 with exponential backoff

---



## **3 - NOTIFICATION TRIGGER & ROUTING DOMAIN**

### 3.1 Event-Driven Notification Triggers

#### Stories

- As a **system**, I want to consume platform events (job.posted, proposal.accepted, payment.processed, etc.) so that relevant notifications are triggered.
- As a **system**, I want to route notifications to appropriate channels based on user preferences so that delivery is optimized.
- As a **system**, I want to enrich notification content with entity details so that messages are contextual.
- As a **system**, I want to apply notification rules (priority, frequency limits, grouping) so that user experience is maintained.

#### Flow

1. **ConsumeJobPostedEvent**(event) → ExtractJobDetails() | MatchPreferences() | EnrichContent() | RouteToChannels() → **Outbox:** notification.triggered.v1
2. **ConsumeProposalAcceptedEvent**(event) → ExtractProposalDetails() | NotifyFreelancer() | RouteToChannels() → **Outbox:** notification.triggered.v1
3. **ConsumePaymentProcessedEvent**(event) → ExtractPaymentDetails() | NotifyBothParties() | RouteToChannels() → **Outbox:** notification.triggered.v1
4. **ConsumeContractCreatedEvent**(event) → ExtractContractDetails() | NotifyBothParties() | RouteToChannels() → **Outbox:** notification.triggered.v1
5. **ConsumeMessageSentEvent**(event) → NotifyRecipients() | CheckOnlineStatus() | RouteToChannels() → **Outbox:** notification.triggered.v1

#### Projections

- notification_triggers_read
- notification_routing_rules_read

#### Events Consumed

- user.created.v1
- user.verified.v1
- job.posted.v1
- job.closed.v1
- job.invitation_sent.v1
- proposal.submitted.v1
- proposal.accepted.v1
- proposal.rejected.v1
- bid.placed.v1
- contract.created.v1
- contract.started.v1
- milestone.completed.v1
- milestone.approved.v1
- payment.processed.v1
- escrow.released.v1
- invoice.generated.v1
- review.submitted.v1
- review.responded.v1
- subscription.expiring.v1
- connects.running_low.v1

#### Events Published

- notification.triggered.v1

#### RBAC/SLO

- **RBAC:** SYSTEM
- **SLO:** P95 < 200ms (trigger and route)

---


## **10 - EVENT CONSUMERS DOMAIN**

### 10.1 User Events → Notification Triggers

#### Stories

- As a **system**, I want to consume user.created events so that welcome emails are sent.
- As a **system**, I want to consume user.verified events so that verification confirmation is sent.
- As a **system**, I want to consume user.banned events so that conversations are quarantined.

#### Flow

- Consume: `user.created.v1` → Send welcome email + in-app notification
- Consume: `user.verified.v1` → Send verification confirmation
- Consume: `user.banned.v1` → Quarantine all conversations, send notification to contacts

#### Projections

- user_notification_triggers_read

#### Events Consumed

- user.created.v1
- user.verified.v1
- user.suspended.v1
- user.banned.v1

#### RBAC/SLO

- **RBAC:** SYSTEM
- **SLO:** P95 < 180ms

---

### 10.2 Job Events → Notification Triggers

#### Stories

- As a **system**, I want to consume job.posted events so that matching freelancers are notified.
- As a **system**, I want to consume job.invitation_sent events so that invitees are notified.
- As a **system**, I want to consume job.closed events so that applicants are notified.

#### Flow

- Consume: `job.posted.v1` → Match job with freelancer preferences → Send notifications
- Consume: `job.invitation_sent.v1` → Send invitation notification
- Consume: `job.closed.v1` → Notify all applicants

#### Projections

- job_notification_triggers_read

#### Events Consumed

- job.posted.v1
- job.updated.v1
- job.closed.v1
- job.invitation_sent.v1

#### RBAC/SLO

- **RBAC:** SYSTEM
- **SLO:** P95 < 200ms

---

### 10.3 Proposal Events → Notification Triggers

#### Stories

- As a **system**, I want to consume proposal.submitted events so that clients are notified.
- As a **system**, I want to consume proposal.accepted events so that freelancers are notified.
- As a **system**, I want to consume bid.placed events so that outbid alerts are sent.

#### Flow

- Consume: `proposal.submitted.v1` → Notify client
- Consume: `proposal.accepted.v1` → Notify freelancer
- Consume: `bid.placed.v1` → Check if user is outbid → Send outbid alert

#### Projections

- proposal_notification_triggers_read

#### Events Consumed

- proposal.submitted.v1
- proposal.accepted.v1
- proposal.rejected.v1
- bid.placed.v1
- bid.updated.v1
- outbid_alert.triggered.v1

#### RBAC/SLO

- **RBAC:** SYSTEM
- **SLO:** P95 < 180ms

---

### 10.4 Contract Events → Notification Triggers

#### Stories

- As a **system**, I want to consume contract.created events so that both parties are notified.
- As a **system**, I want to consume milestone.completed events so that clients are notified.
- As a **system**, I want to consume contract.ended events so that final notifications are sent.

#### Flow

- Consume: `contract.created.v1` → Notify both parties
- Consume: `milestone.completed.v1` → Notify client for approval
- Consume: `contract.ended.v1` → Send completion notifications + review reminders

#### Projections

- contract_notification_triggers_read

#### Events Consumed

- contract.created.v1
- contract.started.v1
- milestone.created.v1
- milestone.completed.v1
- milestone.approved.v1
- timesheet.submitted.v1
- contract.ended.v1

#### RBAC/SLO

- **RBAC:** SYSTEM
- **SLO:** P95 < 180ms

---

### 10.5 Financial Events → Notification Triggers

#### Stories

- As a **system**, I want to consume payment.processed events so that payment receipts are sent.
- As a **system**, I want to consume escrow.released events so that both parties are notified.
- As a **system**, I want to consume invoice.generated events so that payment reminders are sent.

#### Flow

- Consume: `payment.processed.v1` → Send payment confirmation
- Consume: `escrow.released.v1` → Notify freelancer
- Consume: `invoice.generated.v1` → Send invoice notification

#### Projections

- financial_notification_triggers_read

#### Events Consumed

- payment.processed.v1
- payment.failed.v1
- escrow.held.v1
- escrow.released.v1
- payout.processed.v1
- invoice.generated.v1

#### RBAC/SLO

- **RBAC:** SYSTEM
- **SLO:** P95 < 180ms

---

### 10.6 Review Events → Notification Triggers

#### Stories

- As a **system**, I want to consume review.submitted events so that reviewed users are notified.
- As a **system**, I want to consume badge.awarded events so that users are congratulated.

#### Flow

- Consume: `review.submitted.v1` → Notify reviewed user
- Consume: `badge.awarded.v1` → Send congratulations notification

#### Projections

- review_notification_triggers_read

#### Events Consumed

- review.submitted.v1
- review.responded.v1
- badge.awarded.v1

#### RBAC/SLO

- **RBAC:** SYSTEM
- **SLO:** P95 < 150ms

---

### 10.7 Subscription Events → Notification Triggers

#### Stories

- As a **system**, I want to consume subscription.expiring events so that renewal reminders are sent.
- As a **system**, I want to consume connects.running_low events so that users are prompted to purchase.

#### Flow

- Consume: `subscription.expiring.v1` → Send renewal reminder
- Consume: `connects.running_low.v1` → Send purchase prompt

#### Projections

- subscription_notification_triggers_read

#### Events Consumed

- subscription.expiring.v1
- subscription.renewed.v1
- subscription.cancelled.v1
- connects.purchased.v1
- connects.running_low.v1

#### RBAC/SLO

- **RBAC:** SYSTEM
- **SLO:** P95 < 150ms

---











## **11 - COMPLIANCE & PRIVACY DOMAIN**
### 11.1 Data Retention Policies

#### Stories

- As an **admin**, I want to set retention policies per conversation type so that data is managed properly.
- As a **system**, I want to auto-delete expired messages so that storage is optimized.
- As a **user**, I want to configure retention for my conversations so that I control data lifecycle.
- As a **compliance officer**, I want to enforce legal hold on conversations so that litigation requirements are met.

#### Flow

1. **SetRetentionPolicyCommand**(conversation_id, retention_days, user_id) → AuthorizeOwner() | Validate() | Set() → **Outbox:** retention_policy.set.v1
2. **ApplyLegalHoldCommand**(conversation_id, admin_id, reason, hold_until) → AuthorizeAdmin() | ApplyHold() | DisableAutoDelete() → **Outbox:** legal_hold.applied.v1
3. **ReleaseLegalHoldCommand**(conversation_id, admin_id) → AuthorizeAdmin() | Release() | EnableAutoDelete() → **Outbox:** legal_hold.released.v1
4. **DeleteExpiredMessagesCommand**(batch_size) → FetchExpired() | CheckLegalHold() | SoftDelete() → **Outbox:** messages.expired.deleted.v1
5. **GetRetentionPolicyQuery**(conversation_id) → AuthorizeAccess() | Fetch() → RetentionPolicyDTO

#### Projections

- retention_policies_read
- legal_holds_read
- expired_messages_queue_read

#### Events Published

- retention_policy.set.v1
- legal_hold.applied.v1
- legal_hold.released.v1
- messages.expired.deleted.v1

#### RBAC/SLO

- **RBAC:** OWNER (set retention), ADMIN (legal hold), SYSTEM (delete expired)
- **SLO:** P95 < 200ms (set policy), Background processing for deletion

---

### 11.2 GDPR/CCPA Compliance

#### Stories

- As a **user**, I want to request all my communication data so that I can exercise my right to access.
- As a **user**, I want to request deletion of my communication data so that I can exercise my right to be forgotten.
- As a **system**, I want to anonymize user data on deletion so that conversations remain coherent while removing PII.
- As a **compliance officer**, I want audit logs of all data access/deletion requests so that compliance is proven.

#### Flow

1. **RequestDataExportCommand**(user_id) → AuthorizeOwner() | QueueExport() | Notify() → **Outbox:** data_export.requested.v1
2. **ProcessDataExportCommand**(export_id) → FetchAllData(conversations, messages, notifications) | Package() | UploadToStorage() | Notify() → **Outbox:** data_export.completed.v1
3. **RequestDataDeletionCommand**(user_id) → AuthorizeOwner() | QueueDeletion() | Notify() → **Outbox:** data_deletion.requested.v1
4. **ProcessDataDeletionCommand**(deletion_id) → AnonymizeUserData() | RemovePII() | UpdateReferences() | Audit() → **Outbox:** data_deletion.completed.v1
5. **GetDataAccessAuditQuery**(user_id, date_range) → AuthorizeComplianceOfficer() | Fetch() → AuditLogDTO

#### Projections

- data_export_requests_read
- data_deletion_requests_read
- compliance_audit_log_read

#### Events Published

- data_export.requested.v1
- data_export.completed.v1
- data_deletion.requested.v1
- data_deletion.completed.v1
- user_data.anonymized.v1

#### Events Consumed

- user.erasure_requested.v1

#### RBAC/SLO

- **RBAC:** OWNER (request export/deletion), SYSTEM (process), COMPLIANCE_OFFICER (view audit)
- **SLO:** Export ready within 48h, Deletion completed within 30 days

---

### 11.3 Encryption Management

#### Stories

- As a **user**, I want to enable end-to-end encryption for conversations so that privacy is maximized.
- As a **system**, I want to manage encryption keys securely so that data is protected.
- As a **user**, I want to verify encryption status so that I know conversations are secure.

#### Flow

1. **EnableE2EEncryptionCommand**(conversation_id, user_id, public_keys[]) → ValidateKeys() | Enable() | ExchangeKeys() → **Outbox:** e2e_encryption.enabled.v1
2. **DisableE2EEncryptionCommand**(conversation_id, user_id) → AuthorizeOwner() | Disable() → **Outbox:** e2e_encryption.disabled.v1
3. **RotateEncryptionKeyCommand**(conversation_id, new_key) → Validate() | Rotate() | ReEncryptRecent() → **Outbox:** encryption_key.rotated.v1
4. **GetEncryptionStatusQuery**(conversation_id) → AuthorizeParticipant() | Fetch() → EncryptionStatusDTO

#### Projections

- encryption_settings_read
- encryption_keys_read (KMS-encrypted)

#### Events Published

- e2e_encryption.enabled.v1
- e2e_encryption.disabled.v1
- encryption_key.rotated.v1

#### RBAC/SLO

- **RBAC:** OWNER (enable/disable), SYSTEM (rotate)
- **SLO:** P95 < 250ms

---




## **13 - AUTOMATION & BOT INTEGRATION DOMAIN**

### 13.1 Bot Management

#### Stories

- As a **developer**, I want to create bots that can send/receive messages so that automated workflows are supported.
- As a **system**, I want to identify bot messages so that users know they're interacting with automation.
- As a **user**, I want to block bots so that I control automated interactions.

#### Flow

1. **RegisterBotCommand**(bot_name, webhook_url, permissions) → Validate() | Register() | GenerateToken() → **Outbox:** bot.registered.v1
2. **SendBotMessageCommand**(bot_id, conversation_id, content) → AuthorizeBot() | RateLimitCheck() | Send() | MarkAsBot() → **Outbox:** bot_message.sent.v1
3. **BlockBotCommand**(user_id, bot_id) → Block() → **Outbox:** bot.blocked.v1
4. **GetBotQuery**(bot_id) → AuthorizeDeveloper() | Fetch() → BotDTO

#### Projections

- bots_read
- bot_permissions_read
- blocked_bots_read

#### Events Published

- bot.registered.v1
- bot.unregistered.v1
- bot_message.sent.v1
- bot.blocked.v1

#### RBAC/SLO

- **RBAC:** DEVELOPER (register/unregister), BOT (send messages), USER (block)
- **SLO:** P95 < 200ms, Rate limit: 60 msgs/min per bot

---

### 13.2 Automated Responses

#### Stories

- As a **user**, I want to set up auto-replies for when I'm unavailable so that contacts are informed.
- As a **system**, I want to trigger auto-replies based on conditions (time, keywords) so that responses are contextual.

#### Flow

1. **SetAutoReplyCommand**(user_id, message, conditions, enabled) → Validate() | Set() → **Outbox:** auto_reply.set.v1
2. **TriggerAutoReplyCommand**(conversation_id, incoming_message) → CheckConditions() | SendAutoReply() → **Outbox:** auto_reply.sent.v1
3. **DisableAutoReplyCommand**(user_id) → Disable() → **Outbox:** auto_reply.disabled.v1

#### Projections

- auto_reply_settings_read

#### Events Published

- auto_reply.set.v1
- auto_reply.sent.v1
- auto_reply.disabled.v1

#### RBAC/SLO

- **RBAC:** OWNER (set/disable), SYSTEM (trigger)
- **SLO:** P95 < 150ms

---

## **14 - COLLABORATION FEATURES DOMAIN**

### 14.1 Screen Sharing Signaling

#### Stories

- As a **user**, I want to share my screen during calls so that collaboration is enhanced.
- As a **system**, I want to exchange screen sharing signaling so that WebRTC streams are established.

#### Flow

1. **InitiateScreenShareCommand**(call_id, user_id, offer_sdp) → ValidateCall() | BroadcastWebSocket() → **Outbox:** screen_share.initiated.v1
2. **AcceptScreenShareCommand**(call_id, user_id, answer_sdp) → BroadcastWebSocket() → **Outbox:** screen_share.accepted.v1
3. **StopScreenShareCommand**(call_id, user_id) → BroadcastWebSocket() → **Outbox:** screen_share.stopped.v1

#### Projections

- active_screen_shares_read

#### Events Published

- screen_share.initiated.v1
- screen_share.accepted.v1
- screen_share.stopped.v1

#### RBAC/SLO

- **RBAC:** CALL_PARTICIPANT
- **SLO:** P95 < 80ms

---

### 14.2 File Collaboration

#### Stories

- As a **user**, I want to co-edit documents in conversations so that real-time collaboration is supported.
- As a **system**, I want to track document versions so that changes are preserved.

#### Flow

1. **StartCollaborationCommand**(conversation_id, file_id, user_id) → Validate() | CreateSession() → **Outbox:** collaboration.started.v1
2. **UpdateCollaborationCommand**(session_id, changes[], user_id) → ApplyChanges() | BroadcastWebSocket() → **Outbox:** collaboration.updated.v1
3. **EndCollaborationCommand**(session_id, user_id) → SaveVersion() | End() → **Outbox:** collaboration.ended.v1

#### Projections

- collaboration_sessions_read
- document_versions_read

#### Events Published

- collaboration.started.v1
- collaboration.updated.v1
- collaboration.ended.v1

#### RBAC/SLO

- **RBAC:** PARTICIPANT
- **SLO:** P95 < 150ms

---

## **15 - MONITORING & HEALTH DOMAIN**

### 15.1 Service Health Checks

#### Stories

- As an **operator**, I want health check endpoints so that service status is monitored.
- As a **system**, I want to report dependency health so that issues are detected.

#### Flow

1. **LivenessProbeQuery**() → CheckProcess() → HTTP 200 or 503
2. **ReadinessProbeQuery**() → CheckDB() | CheckRedis() | CheckKafka() → HTTP 200 or 503
3. **GetServiceMetricsQuery**() → FetchPrometheus() → MetricsDTO

#### Projections

- service_health_read

#### Events Published

- service.unhealthy.v1 (when dependencies fail)

#### RBAC/SLO

- **RBAC:** PUBLIC (health checks), OPERATOR (metrics)
- **SLO:** P95 < 50ms

---

### 15.2 Performance Monitoring

#### Stories

- As an **operator**, I want real-time performance metrics so that bottlenecks are identified.
- As an **operator**, I want alert thresholds so that incidents are detected early.

#### Flow

1. **GetPerformanceMetricsQuery**(service, date_range) → AuthorizeOperator() | FetchTimeSeries() → MetricsDTO
2. **SetAlertThresholdCommand**(metric, threshold, operator_id) → Set() → **Outbox:** alert_threshold.set.v1
3. **TriggerAlertCommand**(metric, current_value) → CheckThreshold() | SendAlert() → **Outbox:** alert.triggered.v1

#### Projections

- performance_metrics_read
- alert_thresholds_read

#### Events Published

- alert_threshold.set.v1
- alert.triggered.v1

#### RBAC/SLO

- **RBAC:** OPERATOR
- **SLO:** P95 < 100ms

---