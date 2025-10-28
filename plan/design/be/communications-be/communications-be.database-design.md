# COMMUNICATIONS-BE DATABASE DESIGN (Combined)
**Skillsier Platform – Enterprise Scale (Upwork-like)**  
**PostgreSQL 16+**

> This file combines your existing schema with the missing domains you listed.  
> It strictly follows your CRITICAL ALIGNMENT RULES:
> 1) each `internal/domain/{domain}/` → **one** main table named exactly `{domain}`,  
> 2) sub-entities use `{domain}_{sub}`,  
> 3) all domains from the folder structure are covered,  
> 4) fields & indexes are production-ready for large scale.

---

## Global extensions

```sql
-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";
CREATE EXTENSION IF NOT EXISTS "btree_gin";
CREATE EXTENSION IF NOT EXISTS "citext";
CREATE EXTENSION IF NOT EXISTS "btree_gist";
```

---

## SECTION 1–30 (Existing Schema)


=========================================
##  SECTION 1: CORE CONVERSATION DOMAIN
```sql
-- Domain: internal/domain/conversation/
-- Entity: conversation/entity.go
-- =========================================

CREATE TABLE conversations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kind VARCHAR(20) NOT NULL CHECK (kind IN ('DIRECT', 'GROUP', 'SYSTEM', 'SUPPORT', 'BROADCAST')),
    title VARCHAR(200),
    description TEXT,
    tenant_id UUID NOT NULL,
    created_by UUID NOT NULL,
    visibility VARCHAR(20) DEFAULT 'PRIVATE' CHECK (visibility IN ('PRIVATE', 'PUBLIC', 'ORGANIZATION', 'TEAM')),
    data_zone VARCHAR(10) DEFAULT 'US' CHECK (data_zone IN ('US', 'EU', 'ASIA')),
    ttl_policy_id UUID,
    legal_hold BOOLEAN DEFAULT FALSE,
    allow_replies BOOLEAN DEFAULT TRUE,
    allow_files BOOLEAN DEFAULT TRUE,
    slow_mode_seconds INTEGER DEFAULT 0,
    is_encrypted BOOLEAN DEFAULT FALSE,
    encryption_key_id VARCHAR(255),
    metadata JSONB,
    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE', 'ARCHIVED', 'LOCKED', 'DELETED')),
    archived_at TIMESTAMPTZ,
    archived_by UUID,
    locked_at TIMESTAMPTZ,
    locked_by UUID,
    lock_reason TEXT,
    participant_count INTEGER DEFAULT 0,
    message_count INTEGER DEFAULT 0,
    last_message_at TIMESTAMPTZ,
    last_message_preview TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ,
    version INTEGER DEFAULT 1 NOT NULL
);

CREATE INDEX idx_conversations_tenant ON conversations (tenant_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_conversations_created_by ON conversations (created_by) WHERE deleted_at IS NULL;
CREATE INDEX idx_conversations_kind ON conversations (kind, status);
CREATE INDEX idx_conversations_data_zone ON conversations (data_zone);
CREATE INDEX idx_conversations_last_message ON conversations (last_message_at DESC) WHERE status = 'ACTIVE';
CREATE INDEX idx_conversations_metadata ON conversations USING gin(metadata);

COMMENT ON TABLE conversations IS 'Conversations - maps to internal/domain/conversation/entity.go';

```
=========================================
##  SECTION 2: CONVERSATION PARTICIPANTS
```sql
-- Domain: internal/domain/conversation/participant.go
-- =========================================

CREATE TABLE participants (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID NOT NULL,
    user_id UUID NOT NULL,
    role VARCHAR(20) DEFAULT 'MEMBER' CHECK (role IN ('OWNER', 'ADMIN', 'MODERATOR', 'MEMBER', 'GUEST', 'BOT')),
    can_send_messages BOOLEAN DEFAULT TRUE,
    can_add_members BOOLEAN DEFAULT FALSE,
    can_remove_members BOOLEAN DEFAULT FALSE,
    can_manage_settings BOOLEAN DEFAULT FALSE,
    last_read_message_id UUID,
    last_read_seq BIGINT DEFAULT 0,
    unread_count INTEGER DEFAULT 0,
    pinned BOOLEAN DEFAULT FALSE,
    pinned_at TIMESTAMPTZ,
    muted_until TIMESTAMPTZ,
    notification_preference VARCHAR(20) DEFAULT 'ALL' CHECK (notification_preference IN ('ALL', 'MENTIONS_ONLY', 'MUTED')),
    custom_title VARCHAR(200),
    custom_emoji VARCHAR(50),
    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE', 'INVITED', 'LEFT', 'REMOVED', 'BANNED')),
    invited_by UUID,
    invited_at TIMESTAMPTZ,
    joined_at TIMESTAMPTZ,
    left_at TIMESTAMPTZ,
    removed_at TIMESTAMPTZ,
    removed_by UUID,
    removed_reason TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT fk_participants_conversation FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE,
    CONSTRAINT uk_participants_conv_user UNIQUE (conversation_id, user_id)
);

CREATE INDEX idx_participants_conversation ON participants (conversation_id, status);
CREATE INDEX idx_participants_user ON participants (user_id, status);
CREATE INDEX idx_participants_unread ON participants (user_id, unread_count) WHERE unread_count > 0;
CREATE INDEX idx_participants_pinned ON participants (user_id, pinned) WHERE pinned = TRUE;
CREATE INDEX idx_participants_role ON participants (conversation_id, role);

COMMENT ON TABLE participants IS 'Conversation participants - maps to internal/domain/conversation/participant.go';

```
=========================================
##  SECTION 3: TYPING INDICATORS
```sql
-- Domain: internal/domain/conversation/typing_indicator.go
-- =========================================

CREATE TABLE typing_indicators (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID NOT NULL,
    user_id UUID NOT NULL,
    is_typing BOOLEAN DEFAULT TRUE,
    expires_at TIMESTAMPTZ NOT NULL,
    started_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT fk_typing_conversation FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE,
    CONSTRAINT uk_typing_conv_user UNIQUE (conversation_id, user_id)
);

CREATE INDEX idx_typing_conversation ON typing_indicators (conversation_id, expires_at);
CREATE INDEX idx_typing_expires ON typing_indicators (expires_at);

COMMENT ON TABLE typing_indicators IS 'Real-time typing indicators - maps to internal/domain/conversation/typing_indicator.go';

```
=========================================
##  SECTION 4: THREADS (SUB-DISCUSSIONS)
```sql
-- Domain: internal/domain/thread/
-- Entity: thread/entity.go
-- =========================================

CREATE TABLE threads (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID NOT NULL,
    root_message_id UUID NOT NULL,
    title VARCHAR(200),
    message_count INTEGER DEFAULT 0,
    participant_count INTEGER DEFAULT 0,
    last_message_at TIMESTAMPTZ,
    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE', 'LOCKED', 'ARCHIVED')),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    archived_at TIMESTAMPTZ,
    archived_by UUID,
    CONSTRAINT fk_threads_conversation FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
);

CREATE INDEX idx_threads_conversation ON threads (conversation_id, status);
CREATE INDEX idx_threads_root_message ON threads (root_message_id);
CREATE INDEX idx_threads_last_message ON threads (last_message_at DESC) WHERE status = 'ACTIVE';

COMMENT ON TABLE threads IS 'Message threads - maps to internal/domain/thread/entity.go';

CREATE TABLE thread_followers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    thread_id UUID NOT NULL,
    user_id UUID NOT NULL,
    notification_enabled BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT fk_thread_followers_thread FOREIGN KEY (thread_id) REFERENCES threads(id) ON DELETE CASCADE,
    CONSTRAINT uk_thread_followers UNIQUE (thread_id, user_id)
);

CREATE INDEX idx_thread_followers_thread ON thread_followers (thread_id);
CREATE INDEX idx_thread_followers_user ON thread_followers (user_id);

COMMENT ON TABLE thread_followers IS 'Thread followers for notifications';

```
=========================================
##  SECTION 5: MESSAGES (CORE MESSAGING)
```sql
-- Domain: internal/domain/message/
-- Entity: message/entity.go
-- =========================================

CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID NOT NULL,
    thread_id UUID,
    sender_id UUID NOT NULL,
    seq BIGINT NOT NULL,
    body TEXT NOT NULL,
    body_format VARCHAR(20) DEFAULT 'PLAIN' CHECK (body_format IN ('PLAIN', 'MARKDOWN', 'HTML', 'RICH')),
    reply_to_message_id UUID,
    reply_to_seq BIGINT,
    message_type VARCHAR(30) DEFAULT 'TEXT' CHECK (
        message_type IN ('TEXT', 'SYSTEM', 'FILE', 'IMAGE', 'VIDEO', 'AUDIO','LOCATION', 'CONTACT', 'POLL', 'CALL', 'EVENT')
    ),
    system_event_type VARCHAR(50),
    system_event_data JSONB,
    is_edited BOOLEAN DEFAULT FALSE,
    edited_at TIMESTAMPTZ,
    edit_count INTEGER DEFAULT 0,
    original_body TEXT,
    is_deleted BOOLEAN DEFAULT FALSE,
    deleted_at TIMESTAMPTZ,
    deleted_by UUID,
    delete_type VARCHAR(20) CHECK (delete_type IN ('SOFT', 'FOR_ME', 'FOR_EVERYONE')),
    is_flagged BOOLEAN DEFAULT FALSE,
    flagged_at TIMESTAMPTZ,
    flagged_by UUID,
    flag_reason TEXT,
    moderation_status VARCHAR(20) CHECK (moderation_status IN ('PENDING', 'APPROVED', 'REJECTED', 'AUTO_REMOVED')),
    moderated_at TIMESTAMPTZ,
    moderated_by UUID,
    is_quarantined BOOLEAN DEFAULT FALSE,
    quarantine_reason TEXT,
    quarantined_at TIMESTAMPTZ,
    reaction_count INTEGER DEFAULT 0,
    reactions_summary JSONB,
    mentioned_user_ids UUID[],
    mention_count INTEGER DEFAULT 0,
    delivery_status VARCHAR(20) DEFAULT 'SENT' CHECK (delivery_status IN ('PENDING', 'SENT', 'DELIVERED', 'READ', 'FAILED')),
    read_count INTEGER DEFAULT 0,
    read_by_all BOOLEAN DEFAULT FALSE,
    is_encrypted BOOLEAN DEFAULT FALSE,
    encryption_key_version VARCHAR(50),
    metadata JSONB,
    search_vector tsvector GENERATED ALWAYS AS (to_tsvector('english', COALESCE(body, ''))) STORED,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT fk_messages_conversation FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE,
    CONSTRAINT fk_messages_thread FOREIGN KEY (thread_id) REFERENCES threads(id) ON DELETE SET NULL,
    CONSTRAINT uk_messages_conv_seq UNIQUE (conversation_id, seq)
);

CREATE INDEX idx_messages_conversation ON messages (conversation_id, seq DESC) WHERE is_deleted = FALSE;
CREATE INDEX idx_messages_sender ON messages (sender_id, created_at DESC);
CREATE INDEX idx_messages_thread ON messages (thread_id, seq DESC) WHERE thread_id IS NOT NULL;
CREATE INDEX idx_messages_reply_to ON messages (reply_to_message_id) WHERE reply_to_message_id IS NOT NULL;
CREATE INDEX idx_messages_search ON messages USING gin(search_vector);
CREATE INDEX idx_messages_flagged ON messages (is_flagged, flagged_at) WHERE is_flagged = TRUE;
CREATE INDEX idx_messages_type ON messages (message_type, conversation_id);
CREATE INDEX idx_messages_delivery ON messages (delivery_status, created_at);

COMMENT ON TABLE messages IS 'Messages - maps to internal/domain/message/entity.go';

CREATE TABLE message_edit_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    message_id UUID NOT NULL,
    previous_body TEXT NOT NULL,
    previous_body_format VARCHAR(20),
    edited_by UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT fk_message_edit_history_message FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE
);

CREATE INDEX idx_message_edit_history_message ON message_edit_history (message_id, created_at DESC);

COMMENT ON TABLE message_edit_history IS 'Message edit audit trail';

```
=========================================
##  SECTION 6: MESSAGE REACTIONS
```sql
-- Domain: internal/domain/message/reaction.go
-- =========================================

CREATE TABLE reactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    message_id UUID NOT NULL,
    user_id UUID NOT NULL,
    emoji VARCHAR(50) NOT NULL,
    emoji_unicode VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT fk_reactions_message FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE,
    CONSTRAINT uk_reactions_message_user_emoji UNIQUE (message_id, user_id, emoji)
);

CREATE INDEX idx_reactions_message ON reactions (message_id, emoji);
CREATE INDEX idx_reactions_user ON reactions (user_id, created_at DESC);

COMMENT ON TABLE reactions IS 'Message reactions - maps to internal/domain/message/reaction.go';

```
=========================================
##  SECTION 7: MESSAGE ATTACHMENTS
```sql
-- Domain: internal/domain/message/attachment.go
-- =========================================

CREATE TABLE attachments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    message_id UUID NOT NULL,
    file_ref_id UUID NOT NULL,
    storage_url TEXT NOT NULL,
    filename VARCHAR(500) NOT NULL,
    file_type VARCHAR(100) NOT NULL,
    mime_type VARCHAR(100),
    file_size BIGINT NOT NULL,
    file_hash VARCHAR(64) NOT NULL,
    thumbnail_url TEXT,
    thumbnail_width INTEGER,
    thumbnail_height INTEGER,
    width INTEGER,
    height INTEGER,
    duration_seconds INTEGER,
    virus_scan_status VARCHAR(20) DEFAULT 'PENDING' CHECK (virus_scan_status IN ('PENDING', 'CLEAN', 'INFECTED', 'FAILED')),
    virus_scan_result TEXT,
    scanned_at TIMESTAMPTZ,
    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE', 'REMOVED', 'QUARANTINED')),
    removed_at TIMESTAMPTZ,
    removed_by UUID,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT fk_attachments_message FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE
);

CREATE INDEX idx_attachments_message ON attachments (message_id);
CREATE INDEX idx_attachments_hash ON attachments (file_hash);
CREATE INDEX idx_attachments_file_ref ON attachments (file_ref_id);
CREATE INDEX idx_attachments_scan_status ON attachments (virus_scan_status, created_at);

COMMENT ON TABLE attachments IS 'Message attachments - maps to internal/domain/message/attachment.go';

```
=========================================
##  SECTION 8: READ RECEIPTS
```sql
-- Domain: internal/domain/message/read_receipt.go
-- =========================================

CREATE TABLE read_receipts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    message_id UUID NOT NULL,
    user_id UUID NOT NULL,
    status VARCHAR(20) DEFAULT 'DELIVERED' CHECK (status IN ('SENT', 'DELIVERED', 'READ')),
    delivered_at TIMESTAMPTZ,
    read_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT fk_read_receipts_message FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE,
    CONSTRAINT uk_read_receipts_message_user UNIQUE (message_id, user_id)
);

CREATE INDEX idx_read_receipts_message ON read_receipts (message_id, status);
CREATE INDEX idx_read_receipts_user ON read_receipts (user_id, read_at DESC);

COMMENT ON TABLE read_receipts IS 'Message read receipts - maps to internal/domain/message/read_receipt.go';

```
=========================================
##  SECTION 9: MENTIONS
```sql
-- Domain: internal/domain/mention/
-- Entity: mention/entity.go
-- =========================================

CREATE TABLE mentions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    message_id UUID NOT NULL,
    mentioned_user_id UUID NOT NULL,
    mentioned_by UUID NOT NULL,
    conversation_id UUID NOT NULL,
    mention_text VARCHAR(200),
    mention_position INTEGER,
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMPTZ,
    notification_sent BOOLEAN DEFAULT FALSE,
    notification_sent_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT fk_mentions_message FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE,
    CONSTRAINT fk_mentions_conversation FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
);

CREATE INDEX idx_mentions_user ON mentions (mentioned_user_id, is_read, created_at DESC);
CREATE INDEX idx_mentions_message ON mentions (message_id);
CREATE INDEX idx_mentions_conversation ON mentions (conversation_id, created_at DESC);

COMMENT ON TABLE mentions IS 'User mentions - maps to internal/domain/mention/entity.go';

```
=========================================
##  SECTION 10: CONVERSATION EXPORT
```sql
-- Domain: internal/domain/export/
-- Entity: export/entity.go
-- =========================================

CREATE TABLE conversation_exports (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID NOT NULL,
    requested_by UUID NOT NULL,
    format VARCHAR(20) NOT NULL CHECK (format IN ('JSON', 'CSV', 'PDF', 'HTML')),
    date_range_start TIMESTAMPTZ,
    date_range_end TIMESTAMPTZ,
    include_attachments BOOLEAN DEFAULT FALSE,
    filters JSONB,
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'PROCESSING', 'COMPLETED', 'FAILED', 'EXPIRED')),
    progress_percentage INTEGER DEFAULT 0,
    file_ref_id UUID,
    download_url TEXT,
    file_size BIGINT,
    file_hash VARCHAR(64),
    expires_at TIMESTAMPTZ,
    error_message TEXT,
    retry_count INTEGER DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    completed_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT fk_exports_conversation FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
);

CREATE INDEX idx_exports_conversation ON conversation_exports (conversation_id, created_at DESC);
CREATE INDEX idx_exports_requested_by ON conversation_exports (requested_by, status, created_at DESC);
CREATE INDEX idx_exports_status ON conversation_exports (status, created_at);
CREATE INDEX idx_exports_expiration ON conversation_exports (expires_at) WHERE status = 'COMPLETED';

COMMENT ON TABLE conversation_exports IS 'Conversation exports - maps to internal/domain/export/entity.go';

```
=========================================
##  SECTION 11: NOTIFICATIONS (IN-APP)
```sql
-- Domain: internal/domain/notification/
-- Entity: notification/entity.go
-- =========================================

CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    notification_type VARCHAR(50) NOT NULL,
    category VARCHAR(30) NOT NULL CHECK (category IN ('MESSAGE', 'JOB', 'PROPOSAL', 'CONTRACT', 'PAYMENT', 'REVIEW', 'SYSTEM', 'SECURITY', 'SUBSCRIPTION')),
    priority VARCHAR(20) DEFAULT 'NORMAL' CHECK (priority IN ('LOW', 'NORMAL', 'HIGH', 'URGENT')),
    title VARCHAR(200) NOT NULL,
    body TEXT NOT NULL,
    icon_url TEXT,
    image_url TEXT,
    action_url TEXT,
    action_label VARCHAR(100),
    action_data JSONB,
    reference_type VARCHAR(30),
    reference_id UUID,
    group_key VARCHAR(200),
    group_count INTEGER DEFAULT 1,
    actor_id UUID,
    actor_name VARCHAR(200),
    actor_avatar_url TEXT,
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMPTZ,
    is_archived BOOLEAN DEFAULT FALSE,
    archived_at TIMESTAMPTZ,
    is_delivered BOOLEAN DEFAULT FALSE,
    delivered_at TIMESTAMPTZ,
    delivery_channels TEXT[],
    metadata JSONB,
    expires_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT chk_notification_expiration CHECK (expires_at > created_at)
);

CREATE INDEX idx_notifications_user ON notifications (user_id, is_read, created_at DESC);
CREATE INDEX idx_notifications_user_unread ON notifications (user_id, created_at DESC) WHERE is_read = FALSE AND is_archived = FALSE;
CREATE INDEX idx_notifications_type ON notifications (notification_type, created_at DESC);
CREATE INDEX idx_notifications_category ON notifications (category, user_id, created_at DESC);
CREATE INDEX idx_notifications_group ON notifications (user_id, group_key, created_at DESC) WHERE group_key IS NOT NULL;
CREATE INDEX idx_notifications_reference ON notifications (reference_type, reference_id);
CREATE INDEX idx_notifications_priority ON notifications (user_id, priority, created_at DESC) WHERE is_read = FALSE;
CREATE INDEX idx_notifications_expiration ON notifications (expires_at) WHERE expires_at IS NOT NULL;

COMMENT ON TABLE notifications IS 'In-app notifications - maps to internal/domain/notification/entity.go';

```
=========================================
##  SECTION 12: NOTIFICATION PREFERENCES
```sql
-- Domain: internal/domain/notification/preference.go
-- =========================================

CREATE TABLE notification_preferences (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL UNIQUE,
    in_app_enabled BOOLEAN DEFAULT TRUE,
    email_enabled BOOLEAN DEFAULT TRUE,
    push_enabled BOOLEAN DEFAULT TRUE,
    sms_enabled BOOLEAN DEFAULT FALSE,
    message_notifications BOOLEAN DEFAULT TRUE,
    job_notifications BOOLEAN DEFAULT TRUE,
    proposal_notifications BOOLEAN DEFAULT TRUE,
    contract_notifications BOOLEAN DEFAULT TRUE,
    payment_notifications BOOLEAN DEFAULT TRUE,
    review_notifications BOOLEAN DEFAULT TRUE,
    system_notifications BOOLEAN DEFAULT TRUE,
    security_notifications BOOLEAN DEFAULT TRUE,
    subscription_notifications BOOLEAN DEFAULT TRUE,
    notification_frequency VARCHAR(20) DEFAULT 'REALTIME' CHECK (notification_frequency IN ('REALTIME', 'HOURLY', 'DAILY', 'WEEKLY')),
    quiet_hours_enabled BOOLEAN DEFAULT FALSE,
    quiet_hours_start TIME,
    quiet_hours_end TIME,
    quiet_hours_timezone VARCHAR(50),
    email_digest_enabled BOOLEAN DEFAULT FALSE,
    email_digest_frequency VARCHAR(20) CHECK (email_digest_frequency IN ('DAILY', 'WEEKLY', 'MONTHLY')),
    email_digest_time TIME,
    group_notifications BOOLEAN DEFAULT TRUE,
    sound_enabled BOOLEAN DEFAULT TRUE,
    vibration_enabled BOOLEAN DEFAULT TRUE,
    dnd_enabled BOOLEAN DEFAULT FALSE,
    dnd_until TIMESTAMPTZ,
    metadata JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_notification_prefs_user ON notification_preferences (user_id);
CREATE INDEX idx_notification_prefs_dnd ON notification_preferences (dnd_enabled, dnd_until) WHERE dnd_enabled = TRUE;

COMMENT ON TABLE notification_preferences IS 'User notification preferences - maps to internal/domain/notification/preference.go';

```
=========================================
##  SECTION 13: EMAIL NOTIFICATIONS
```sql
-- Domain: internal/domain/email/
-- Entity: email/entity.go
-- =========================================

CREATE TABLE email_notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    to_email CITEXT NOT NULL,
    subject VARCHAR(500) NOT NULL,
    body_html TEXT,
    body_text TEXT NOT NULL,
    template_id VARCHAR(100),
    template_data JSONB,
    notification_id UUID,
    priority VARCHAR(20) DEFAULT 'NORMAL' CHECK (priority IN ('LOW', 'NORMAL', 'HIGH', 'URGENT')),
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'QUEUED', 'SENDING', 'SENT', 'DELIVERED', 'BOUNCED', 'COMPLAINED', 'FAILED')),
    message_id VARCHAR(255),
    provider_message_id VARCHAR(255),
    queued_at TIMESTAMPTZ,
    sent_at TIMESTAMPTZ,
    delivered_at TIMESTAMPTZ,
    opened_at TIMESTAMPTZ,
    clicked_at TIMESTAMPTZ,
    bounced_at TIMESTAMPTZ,
    complained_at TIMESTAMPTZ,
    error_message TEXT,
    bounce_type VARCHAR(50),
    retry_count INTEGER DEFAULT 0,
    max_retries INTEGER DEFAULT 3,
    open_tracking_enabled BOOLEAN DEFAULT TRUE,
    click_tracking_enabled BOOLEAN DEFAULT TRUE,
    unsubscribe_token UUID DEFAULT uuid_generate_v4(),
    metadata JSONB,
    headers JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_email_notifications_user ON email_notifications (user_id, created_at DESC);
CREATE INDEX idx_email_notifications_status ON email_notifications (status, created_at);
CREATE INDEX idx_email_notifications_queued ON email_notifications (queued_at) WHERE status = 'QUEUED';
CREATE INDEX idx_email_notifications_notification ON email_notifications (notification_id);
CREATE INDEX idx_email_notifications_message_id ON email_notifications (message_id);

COMMENT ON TABLE email_notifications IS 'Email notifications - maps to internal/domain/email/entity.go';

CREATE TABLE email_templates (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    template_id VARCHAR(100) NOT NULL UNIQUE,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    category VARCHAR(50),
    subject VARCHAR(500) NOT NULL,
    body_html TEXT NOT NULL,
    body_text TEXT,
    variables JSONB,
    version INTEGER DEFAULT 1,
    is_active BOOLEAN DEFAULT TRUE,
    test_data JSONB,
    metadata JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by UUID
);

CREATE INDEX idx_email_templates_template_id ON email_templates (template_id, is_active);
CREATE INDEX idx_email_templates_category ON email_templates (category);

COMMENT ON TABLE email_templates IS 'Email templates for notifications';

```
=========================================
##  SECTION 14: PUSH NOTIFICATIONS (WebPush)
```sql
-- Domain: internal/domain/push/
-- Entity: push/entity.go
-- =========================================

CREATE TABLE push_notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    device_id UUID NOT NULL,
    title VARCHAR(200) NOT NULL,
    body TEXT NOT NULL,
    icon_url TEXT,
    badge_url TEXT,
    image_url TEXT,
    action_url TEXT,
    actions JSONB,
    notification_id UUID,
    priority VARCHAR(20) DEFAULT 'NORMAL' CHECK (priority IN ('LOW', 'NORMAL', 'HIGH', 'URGENT')),
    urgency VARCHAR(20) DEFAULT 'normal' CHECK (urgency IN ('very-low', 'low', 'normal', 'high')),
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'QUEUED', 'SENDING', 'SENT', 'DELIVERED', 'FAILED', 'EXPIRED')),
    endpoint TEXT,
    subscription_id UUID,
    queued_at TIMESTAMPTZ,
    sent_at TIMESTAMPTZ,
    delivered_at TIMESTAMPTZ,
    clicked_at TIMESTAMPTZ,
    expired_at TIMESTAMPTZ,
    error_message TEXT,
    error_code VARCHAR(50),
    retry_count INTEGER DEFAULT 0,
    max_retries INTEGER DEFAULT 3,
    ttl_seconds INTEGER DEFAULT 86400,
    metadata JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_push_notifications_user ON push_notifications (user_id, created_at DESC);
CREATE INDEX idx_push_notifications_device ON push_notifications (device_id, status);
CREATE INDEX idx_push_notifications_status ON push_notifications (status, created_at);
CREATE INDEX idx_push_notifications_queued ON push_notifications (queued_at) WHERE status = 'QUEUED';
CREATE INDEX idx_push_notifications_notification ON push_notifications (notification_id);

COMMENT ON TABLE push_notifications IS 'Push notifications - maps to internal/domain/push/entity.go';

CREATE TABLE push_subscriptions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    device_id UUID NOT NULL,
    endpoint TEXT NOT NULL,
    keys_p256dh TEXT NOT NULL,
    keys_auth TEXT NOT NULL,
    device_type VARCHAR(20) CHECK (device_type IN ('WEB', 'MOBILE', 'DESKTOP')),
    browser VARCHAR(100),
    os VARCHAR(100),
    is_active BOOLEAN DEFAULT TRUE,
    expires_at TIMESTAMPTZ,
    metadata JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    last_used_at TIMESTAMPTZ,
    CONSTRAINT uk_push_subscriptions_endpoint UNIQUE (endpoint)
);

CREATE INDEX idx_push_subscriptions_user ON push_subscriptions (user_id, is_active);
CREATE INDEX idx_push_subscriptions_device ON push_subscriptions (device_id, is_active);
CREATE INDEX idx_push_subscriptions_expires ON push_subscriptions (expires_at) WHERE is_active = TRUE;

COMMENT ON TABLE push_subscriptions IS 'WebPush subscriptions';

```
=========================================
##  SECTION 15: SMS NOTIFICATIONS (Optional)
```sql
-- Domain: internal/domain/sms/
-- Entity: sms/entity.go
-- =========================================

CREATE TABLE sms_notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    phone_number VARCHAR(50) NOT NULL,
    body TEXT NOT NULL,
    notification_id UUID,
    priority VARCHAR(20) DEFAULT 'NORMAL' CHECK (priority IN ('LOW', 'NORMAL', 'HIGH', 'URGENT')),
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'QUEUED', 'SENDING', 'SENT', 'DELIVERED', 'FAILED', 'UNDELIVERED')),
    provider VARCHAR(50),
    provider_message_id VARCHAR(255),
    queued_at TIMESTAMPTZ,
    sent_at TIMESTAMPTZ,
    delivered_at TIMESTAMPTZ,
    error_message TEXT,
    error_code VARCHAR(50),
    retry_count INTEGER DEFAULT 0,
    max_retries INTEGER DEFAULT 3,
    cost_amount DECIMAL(10, 4),
    cost_currency CHAR(3) DEFAULT 'USD',
    metadata JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_sms_notifications_user ON sms_notifications (user_id, created_at DESC);
CREATE INDEX idx_sms_notifications_status ON sms_notifications (status, created_at);
CREATE INDEX idx_sms_notifications_queued ON sms_notifications (queued_at) WHERE status = 'QUEUED';
CREATE INDEX idx_sms_notifications_notification ON sms_notifications (notification_id);
CREATE INDEX idx_sms_notifications_phone ON sms_notifications (phone_number, created_at DESC);

COMMENT ON TABLE sms_notifications IS 'SMS notifications - maps to internal/domain/sms/entity.go';

```
=========================================
##  SECTION 16: WEBSOCKET CONNECTIONS
```sql
-- Domain: internal/domain/websocket/
-- Entity: websocket/entity.go
-- =========================================

CREATE TABLE websocket_connections (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    session_id UUID NOT NULL,
    connection_id VARCHAR(100) NOT NULL UNIQUE,
    remote_addr INET,
    user_agent TEXT,
    device_type VARCHAR(20),
    device_id UUID,
    status VARCHAR(20) DEFAULT 'CONNECTED' CHECK (status IN ('CONNECTING', 'CONNECTED', 'DISCONNECTED', 'RECONNECTING')),
    last_heartbeat_at TIMESTAMPTZ,
    heartbeat_interval_seconds INTEGER DEFAULT 30,
    subscribed_channels TEXT[],
    messages_sent INTEGER DEFAULT 0,
    messages_received INTEGER DEFAULT 0,
    errors_count INTEGER DEFAULT 0,
    queue_size INTEGER DEFAULT 0,
    max_queue_size INTEGER DEFAULT 1000,
    metadata JSONB,
    connected_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    disconnected_at TIMESTAMPTZ,
    last_activity_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_websocket_connections_user ON websocket_connections (user_id, status);
CREATE INDEX idx_websocket_connections_session ON websocket_connections (session_id);
CREATE INDEX idx_websocket_connections_connection_id ON websocket_connections (connection_id);
CREATE INDEX idx_websocket_connections_status ON websocket_connections (status, last_heartbeat_at);
CREATE INDEX idx_websocket_connections_heartbeat ON websocket_connections (last_heartbeat_at) WHERE status = 'CONNECTED';

COMMENT ON TABLE websocket_connections IS 'WebSocket connections - maps to internal/domain/websocket/entity.go';

```
=========================================
##  SECTION 17: REAL-TIME PRESENCE
```sql
-- Domain: internal/domain/presence/
-- Entity: presence/entity.go
-- =========================================

CREATE TABLE user_presence (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL UNIQUE,
    status VARCHAR(20) DEFAULT 'OFFLINE' CHECK (status IN ('ONLINE', 'AWAY', 'BUSY', 'OFFLINE', 'INVISIBLE')),
    custom_status VARCHAR(200),
    custom_emoji VARCHAR(50),
    custom_status_expires_at TIMESTAMPTZ,
    last_activity_at TIMESTAMPTZ,
    is_active BOOLEAN DEFAULT FALSE,
    active_device_count INTEGER DEFAULT 0,
    metadata JSONB,
    status_changed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_user_presence_user ON user_presence (user_id);
CREATE INDEX idx_user_presence_status ON user_presence (status, last_activity_at DESC);
CREATE INDEX idx_user_presence_active ON user_presence (is_active, last_activity_at DESC);

COMMENT ON TABLE user_presence IS 'User presence status - maps to internal/domain/presence/entity.go';

CREATE TABLE presence_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    from_status VARCHAR(20),
    to_status VARCHAR(20) NOT NULL,
    duration_seconds INTEGER,
    changed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_presence_history_user ON presence_history (user_id, changed_at DESC);

COMMENT ON TABLE presence_history IS 'Presence status change history';

```
=========================================
##  SECTION 18: COLLABORATION (REAL-TIME EDITING)
```sql
-- Domain: internal/domain/collaboration/
-- Entity: collaboration/entity.go
-- =========================================

CREATE TABLE collaboration_sessions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    file_ref_id UUID NOT NULL,
    document_type VARCHAR(50),
    session_name VARCHAR(200),
    owner_id UUID NOT NULL,
    status VARCHAR(20) DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE', 'LOCKED', 'ENDED')),
    participant_count INTEGER DEFAULT 0,
    active_participant_count INTEGER DEFAULT 0,
    current_version INTEGER DEFAULT 1,
    metadata JSONB,
    started_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ended_at TIMESTAMPTZ,
    last_activity_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_collaboration_sessions_file_ref ON collaboration_sessions (file_ref_id);
CREATE INDEX idx_collaboration_sessions_owner ON collaboration_sessions (owner_id);
CREATE INDEX idx_collaboration_sessions_status ON collaboration_sessions (status, last_activity_at DESC);

COMMENT ON TABLE collaboration_sessions IS 'Real-time collaboration sessions - maps to internal/domain/collaboration/entity.go';

CREATE TABLE collaboration_participants (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    session_id UUID NOT NULL,
    user_id UUID NOT NULL,
    role VARCHAR(20) DEFAULT 'EDITOR' CHECK (role IN ('OWNER', 'EDITOR', 'VIEWER')),
    is_active BOOLEAN DEFAULT FALSE,
    cursor_position JSONB,
    selection JSONB,
    last_activity_at TIMESTAMPTZ,
    joined_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    left_at TIMESTAMPTZ,
    CONSTRAINT fk_collaboration_participants_session FOREIGN KEY (session_id) REFERENCES collaboration_sessions(id) ON DELETE CASCADE,
    CONSTRAINT uk_collaboration_participants UNIQUE (session_id, user_id)
);

CREATE INDEX idx_collaboration_participants_session ON collaboration_participants (session_id, is_active);
CREATE INDEX idx_collaboration_participants_user ON collaboration_participants (user_id, is_active);

COMMENT ON TABLE collaboration_participants IS 'Collaboration session participants';

CREATE TABLE document_versions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    session_id UUID NOT NULL,
    version_number INTEGER NOT NULL,
    content_snapshot JSONB,
    content_hash VARCHAR(64),
    changes_description TEXT,
    created_by UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT fk_document_versions_session FOREIGN KEY (session_id) REFERENCES collaboration_sessions(id) ON DELETE CASCADE,
    CONSTRAINT uk_document_versions UNIQUE (session_id, version_number)
);

CREATE INDEX idx_document_versions_session ON document_versions (session_id, version_number DESC);

COMMENT ON TABLE document_versions IS 'Document version history';

```
=========================================
##  SECTION 19: EMAIL DIGEST
```sql
-- Domain: internal/domain/digest/
-- Entity: digest/entity.go
-- =========================================

CREATE TABLE email_digests (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    frequency VARCHAR(20) NOT NULL CHECK (frequency IN ('DAILY', 'WEEKLY', 'MONTHLY')),
    period_start TIMESTAMPTZ NOT NULL,
    period_end TIMESTAMPTZ NOT NULL,
    notification_count INTEGER DEFAULT 0,
    notifications_summary JSONB,
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'GENERATED', 'SENT', 'FAILED')),
    email_notification_id UUID,
    scheduled_at TIMESTAMPTZ NOT NULL,
    generated_at TIMESTAMPTZ,
    sent_at TIMESTAMPTZ,
    error_message TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_email_digests_user ON email_digests (user_id, scheduled_at DESC);
CREATE INDEX idx_email_digests_status ON email_digests (status, scheduled_at);
CREATE INDEX idx_email_digests_scheduled ON email_digests (scheduled_at) WHERE status = 'PENDING';

COMMENT ON TABLE email_digests IS 'Email digest notifications - maps to internal/domain/digest/entity.go';

```
=========================================
##  SECTION 20: WEBHOOKS
```sql
-- Domain: internal/domain/webhook/
-- Entity: webhook/entity.go
-- =========================================

CREATE TABLE webhooks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    tenant_id UUID,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    url TEXT NOT NULL,
    subscribed_events TEXT[] NOT NULL,
    secret_key VARCHAR(255) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    rate_limit_per_minute INTEGER DEFAULT 60,
    max_retries INTEGER DEFAULT 3,
    retry_backoff_seconds INTEGER DEFAULT 60,
    total_deliveries INTEGER DEFAULT 0,
    successful_deliveries INTEGER DEFAULT 0,
    failed_deliveries INTEGER DEFAULT 0,
    last_delivery_at TIMESTAMPTZ,
    last_delivery_status VARCHAR(20),
    metadata JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by UUID
);

CREATE INDEX idx_webhooks_user ON webhooks (user_id, is_active);
CREATE INDEX idx_webhooks_tenant ON webhooks (tenant_id, is_active);

COMMENT ON TABLE webhooks IS 'Webhook subscriptions - maps to internal/domain/webhook/entity.go';

CREATE TABLE webhook_deliveries (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    webhook_id UUID NOT NULL,
    event_type VARCHAR(100) NOT NULL,
    event_id UUID NOT NULL,
    payload JSONB NOT NULL,
    signature VARCHAR(255) NOT NULL,
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'SENDING', 'SUCCESS', 'FAILED', 'RETRYING')),
    http_status_code INTEGER,
    response_body TEXT,
    response_headers JSONB,
    duration_ms INTEGER,
    retry_count INTEGER DEFAULT 0,
    next_retry_at TIMESTAMPTZ,
    error_message TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    sent_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    CONSTRAINT fk_webhook_deliveries_webhook FOREIGN KEY (webhook_id) REFERENCES webhooks(id) ON DELETE CASCADE
);

CREATE INDEX idx_webhook_deliveries_webhook ON webhook_deliveries (webhook_id, created_at DESC);
CREATE INDEX idx_webhook_deliveries_status ON webhook_deliveries (status, created_at);
CREATE INDEX idx_webhook_deliveries_event ON webhook_deliveries (event_type, event_id);
CREATE INDEX idx_webhook_deliveries_retry ON webhook_deliveries (next_retry_at) WHERE status = 'RETRYING';

COMMENT ON TABLE webhook_deliveries IS 'Webhook delivery attempts';

```
=========================================
##  SECTION 21: MESSAGE SEARCH INDEX
```sql
-- Domain: internal/domain/search/
-- Entity: search/entity.go
-- =========================================

CREATE TABLE message_search_index (
    message_id UUID PRIMARY KEY,
    conversation_id UUID NOT NULL,
    sender_id UUID NOT NULL,
    content_text TEXT NOT NULL,
    content_normalized TEXT,
    conversation_title VARCHAR(200),
    sender_name VARCHAR(200),
    message_type VARCHAR(30),
    has_attachments BOOLEAN DEFAULT FALSE,
    has_mentions BOOLEAN DEFAULT FALSE,
    sent_date DATE,
    sent_time TIME,
    search_vector tsvector,
    search_rank INTEGER DEFAULT 0,
    metadata JSONB,
    indexed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_message_search_conversation ON message_search_index (conversation_id, sent_date DESC);
CREATE INDEX idx_message_search_sender ON message_search_index (sender_id, sent_date DESC);
CREATE INDEX idx_message_search_vector ON message_search_index USING gin(search_vector);
CREATE INDEX idx_message_search_filters ON message_search_index (message_type, has_attachments, has_mentions);

COMMENT ON TABLE message_search_index IS 'Message search index - maps to internal/domain/search/entity.go';

```
=========================================
##  SECTION 22: COMPLIANCE & RETENTION
```sql
-- Domain: internal/domain/compliance/
-- Entity: compliance/entity.go
-- =========================================

CREATE TABLE retention_policies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    policy_name VARCHAR(200) NOT NULL,
    description TEXT,
    applies_to VARCHAR(30) CHECK (applies_to IN ('ALL', 'CONVERSATION', 'MESSAGE', 'NOTIFICATION', 'ATTACHMENT')),
    retention_days INTEGER NOT NULL,
    action_on_expiry VARCHAR(20) DEFAULT 'DELETE' CHECK (action_on_expiry IN ('DELETE', 'ARCHIVE', 'ANONYMIZE')),
    bypass_legal_hold BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    metadata JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by UUID
);

CREATE INDEX idx_retention_policies_applies_to ON retention_policies (applies_to, is_active);

COMMENT ON TABLE retention_policies IS 'Data retention policies - maps to internal/domain/compliance/entity.go';

CREATE TABLE gdpr_requests (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    request_type VARCHAR(30) NOT NULL CHECK (request_type IN ('ACCESS', 'EXPORT', 'DELETE', 'RECTIFICATION', 'RESTRICTION')),
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'PROCESSING', 'COMPLETED', 'REJECTED', 'FAILED')),
    progress_percentage INTEGER DEFAULT 0,
    result_file_ref_id UUID,
    result_data JSONB,
    requested_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    processing_started_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    error_message TEXT,
    verified_by UUID,
    verified_at TIMESTAMPTZ,
    verification_method VARCHAR(50),
    metadata JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_gdpr_requests_user ON gdpr_requests (user_id, requested_at DESC);
CREATE INDEX idx_gdpr_requests_status ON gdpr_requests (status, requested_at);

COMMENT ON TABLE gdpr_requests IS 'GDPR data subject requests';

```
=========================================
##  SECTION 23: RATE LIMITING & QUOTAS
```sql
-- Domain: internal/domain/quota/
-- Entity: quota/entity.go
-- =========================================

CREATE TABLE rate_limits (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID,
    tenant_id UUID,
    resource_type VARCHAR(50) NOT NULL,
    limit_per_minute INTEGER,
    limit_per_hour INTEGER,
    limit_per_day INTEGER,
    limit_per_month INTEGER,
    usage_this_minute INTEGER DEFAULT 0,
    usage_this_hour INTEGER DEFAULT 0,
    usage_this_day INTEGER DEFAULT 0,
    usage_this_month INTEGER DEFAULT 0,
    minute_reset_at TIMESTAMPTZ,
    hour_reset_at TIMESTAMPTZ,
    day_reset_at TIMESTAMPTZ,
    month_reset_at TIMESTAMPTZ,
    burst_limit INTEGER,
    is_throttled BOOLEAN DEFAULT FALSE,
    throttled_until TIMESTAMPTZ,
    metadata JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT uk_rate_limits UNIQUE (user_id, tenant_id, resource_type)
);

CREATE INDEX idx_rate_limits_user ON rate_limits (user_id, resource_type);
CREATE INDEX idx_rate_limits_tenant ON rate_limits (tenant_id, resource_type);
CREATE INDEX idx_rate_limits_throttled ON rate_limits (is_throttled, throttled_until) WHERE is_throttled = TRUE;

COMMENT ON TABLE rate_limits IS 'Rate limiting quotas - maps to internal/domain/quota/entity.go';

CREATE TABLE usage_metrics (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID,
    tenant_id UUID,
    resource_type VARCHAR(50) NOT NULL,
    count INTEGER NOT NULL DEFAULT 1,
    period_start TIMESTAMPTZ NOT NULL,
    period_end TIMESTAMPTZ NOT NULL,
    period_type VARCHAR(20) CHECK (period_type IN ('MINUTE', 'HOUR', 'DAY', 'MONTH')),
    metadata JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_usage_metrics_user ON usage_metrics (user_id, resource_type, period_start DESC);
CREATE INDEX idx_usage_metrics_tenant ON usage_metrics (tenant_id, resource_type, period_start DESC);
CREATE INDEX idx_usage_metrics_period ON usage_metrics (period_start, period_type);

COMMENT ON TABLE usage_metrics IS 'Usage metrics aggregation';

```
=========================================
##  SECTION 24: EVENT SOURCING & OUTBOX
```sql
-- Domain: internal/domain/outbox/
-- Entity: outbox/entity.go
-- =========================================

CREATE TABLE outbox_events (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    event_id UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    event_type VARCHAR(100) NOT NULL,
    event_version VARCHAR(10) DEFAULT 'v1',
    aggregate_type VARCHAR(50) NOT NULL,
    aggregate_id UUID NOT NULL,
    payload JSONB NOT NULL,
    correlation_id UUID,
    causation_id UUID,
    user_id UUID,
    tenant_id UUID,
    topic VARCHAR(100) NOT NULL,
    partition_key VARCHAR(255),
    data_zone VARCHAR(10) DEFAULT 'US',
    signature VARCHAR(255),
    status VARCHAR(20) DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'PROCESSING', 'PUBLISHED', 'FAILED')),
    retry_count INTEGER DEFAULT 0,
    max_retries INTEGER DEFAULT 5,
    next_retry_at TIMESTAMPTZ,
    error_message TEXT,
    occurred_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    published_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_outbox_events_status ON outbox_events (status, created_at);
CREATE INDEX idx_outbox_events_aggregate ON outbox_events (aggregate_type, aggregate_id, occurred_at DESC);
CREATE INDEX idx_outbox_events_topic ON outbox_events (topic, status);
CREATE INDEX idx_outbox_events_retry ON outbox_events (next_retry_at) WHERE status = 'FAILED';
CREATE INDEX idx_outbox_events_correlation ON outbox_events (correlation_id);

COMMENT ON TABLE outbox_events IS 'Transactional outbox for event publishing';

CREATE TABLE dead_letter_queue (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    original_event_id UUID NOT NULL,
    event_type VARCHAR(100) NOT NULL,
    payload JSONB NOT NULL,
    error_message TEXT NOT NULL,
    error_stack_trace TEXT,
    retry_count INTEGER NOT NULL,
    source VARCHAR(100),
    status VARCHAR(20) DEFAULT 'UNRESOLVED' CHECK (status IN ('UNRESOLVED', 'INVESTIGATING', 'RESOLVED', 'DISCARDED')),
    resolved_at TIMESTAMPTZ,
    resolved_by UUID,
    resolution_notes TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_dlq_status ON dead_letter_queue (status, created_at DESC);
CREATE INDEX idx_dlq_event_type ON dead_letter_queue (event_type, created_at DESC);

COMMENT ON TABLE dead_letter_queue IS 'Dead letter queue for failed events';

```
=========================================
##  SECTION 25: AUDIT LOGS
```sql

-- =========================================

CREATE TABLE audit_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    entity_type VARCHAR(100) NOT NULL,
    entity_id UUID NOT NULL,
    action VARCHAR(100) NOT NULL,
    action_category VARCHAR(50),
    actor_user_id UUID,
    actor_type VARCHAR(20),
    actor_ip INET,
    actor_user_agent TEXT,
    old_values JSONB,
    new_values JSONB,
    changed_fields TEXT[],
    request_id UUID,
    correlation_id UUID,
    gdpr_relevant BOOLEAN DEFAULT FALSE,
    pii_accessed BOOLEAN DEFAULT FALSE,
    occurred_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_audit_logs_entity ON audit_logs (entity_type, entity_id, occurred_at DESC);
CREATE INDEX idx_audit_logs_actor ON audit_logs (actor_user_id, occurred_at DESC);
CREATE INDEX idx_audit_logs_action ON audit_logs (action, occurred_at DESC);
CREATE INDEX idx_audit_logs_compliance ON audit_logs (occurred_at DESC) WHERE gdpr_relevant = TRUE OR pii_accessed = TRUE;

COMMENT ON TABLE audit_logs IS 'Comprehensive audit trail';

```
=========================================
##  SECTION 26: READ MODELS (CQRS)

```sql

-- =========================================

CREATE TABLE conversation_read_model (
    conversation_id UUID PRIMARY KEY,
    kind VARCHAR(20),
    title VARCHAR(200),
    visibility VARCHAR(20),
    created_by UUID,
    status VARCHAR(20),
    participant_count INTEGER,
    message_count INTEGER,
    unread_count INTEGER,
    last_message_at TIMESTAMPTZ,
    last_message_preview TEXT,
    last_sender_name VARCHAR(200),
    is_pinned BOOLEAN DEFAULT FALSE,
    is_muted BOOLEAN DEFAULT FALSE,
    is_archived BOOLEAN DEFAULT FALSE,
    metadata JSONB,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_conversation_read_user ON conversation_read_model (created_by, last_message_at DESC);
CREATE INDEX idx_conversation_read_status ON conversation_read_model (status, last_message_at DESC);

COMMENT ON TABLE conversation_read_model IS 'CQRS read model for conversations';

CREATE TABLE user_notification_summary (
    user_id UUID PRIMARY KEY,
    total_unread INTEGER DEFAULT 0,
    message_unread INTEGER DEFAULT 0,
    job_unread INTEGER DEFAULT 0,
    proposal_unread INTEGER DEFAULT 0,
    contract_unread INTEGER DEFAULT 0,
    payment_unread INTEGER DEFAULT 0,
    review_unread INTEGER DEFAULT 0,
    system_unread INTEGER DEFAULT 0,
    last_notification_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_user_notification_summary_user ON user_notification_summary (user_id);

COMMENT ON TABLE user_notification_summary IS 'User notification summary for quick access';

```
=========================================
##  SECTION 27: ANALYTICS & METRICS

```sql
-- =========================================

CREATE TABLE communication_metrics (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    metric_date DATE NOT NULL,
    metric_hour INTEGER,
    user_id UUID,
    tenant_id UUID,
    conversation_id UUID,
    messages_sent INTEGER DEFAULT 0,
    messages_received INTEGER DEFAULT 0,
    messages_read INTEGER DEFAULT 0,
    notifications_sent INTEGER DEFAULT 0,
    notifications_delivered INTEGER DEFAULT 0,
    notifications_read INTEGER DEFAULT 0,
    emails_sent INTEGER DEFAULT 0,
    emails_delivered INTEGER DEFAULT 0,
    emails_opened INTEGER DEFAULT 0,
    emails_clicked INTEGER DEFAULT 0,
    emails_bounced INTEGER DEFAULT 0,
    push_sent INTEGER DEFAULT 0,
    push_delivered INTEGER DEFAULT 0,
    push_clicked INTEGER DEFAULT 0,
    active_users INTEGER DEFAULT 0,
    active_conversations INTEGER DEFAULT 0,
    metadata JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT uk_communication_metrics UNIQUE (metric_date, metric_hour, user_id, tenant_id, conversation_id)
);

CREATE INDEX idx_communication_metrics_date ON communication_metrics (metric_date DESC, metric_hour DESC);
CREATE INDEX idx_communication_metrics_user ON communication_metrics (user_id, metric_date DESC);
CREATE INDEX idx_communication_metrics_tenant ON communication_metrics (tenant_id, metric_date DESC);
CREATE INDEX idx_communication_metrics_conversation ON communication_metrics (conversation_id, metric_date DESC);

COMMENT ON TABLE communication_metrics IS 'Communication analytics and metrics';

```
=========================================
##  SECTION 28: VIEWS FOR DASHBOARDS

```sql
-- =========================================

CREATE OR REPLACE VIEW v_active_conversations AS
SELECT
    c.id AS conversation_id,
    c.kind,
    c.title,
    c.visibility,
    c.created_by,
    c.participant_count,
    c.message_count,
    c.last_message_at,
    c.last_message_preview,
    c.created_at
FROM conversations c
WHERE c.status = 'ACTIVE'
  AND c.deleted_at IS NULL
ORDER BY c.last_message_at DESC;

CREATE OR REPLACE VIEW v_user_unread_messages AS
SELECT
    p.user_id,
    p.conversation_id,
    c.title AS conversation_title,
    c.kind AS conversation_kind,
    p.unread_count,
    c.last_message_at,
    c.last_message_preview
FROM participants p
JOIN conversations c ON p.conversation_id = c.id
WHERE p.status = 'ACTIVE'
  AND p.unread_count > 0
  AND c.status = 'ACTIVE'
  AND c.deleted_at IS NULL
ORDER BY c.last_message_at DESC;

CREATE OR REPLACE VIEW v_pending_notifications AS
SELECT
    n.id AS notification_id,
    n.user_id,
    n.notification_type,
    n.category,
    n.priority,
    n.title,
    n.created_at
FROM notifications n
WHERE n.is_read = FALSE
  AND n.is_archived = FALSE
  AND (n.expires_at IS NULL OR n.expires_at > CURRENT_TIMESTAMP)
ORDER BY n.priority DESC, n.created_at DESC;

```
=========================================
##  SECTION 29: TRIGGERS

```sql
-- =========================================

CREATE OR REPLACE FUNCTION update_conversation_last_message()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE conversations
    SET last_message_at = NEW.created_at,
        last_message_preview = LEFT(NEW.body, 200),
        message_count = message_count + 1,
        updated_at = CURRENT_TIMESTAMP
    WHERE id = NEW.conversation_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_update_conversation_last_message
AFTER INSERT ON messages
FOR EACH ROW EXECUTE FUNCTION update_conversation_last_message();

CREATE OR REPLACE FUNCTION update_participant_unread_count()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE participants
    SET unread_count = unread_count + 1
    WHERE conversation_id = NEW.conversation_id
      AND user_id != NEW.sender_id
      AND status = 'ACTIVE';
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_update_participant_unread_count
AFTER INSERT ON messages
FOR EACH ROW
WHEN (NEW.message_type = 'TEXT')
EXECUTE FUNCTION update_participant_unread_count();

CREATE OR REPLACE FUNCTION update_message_reaction_count()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' THEN
        UPDATE messages
        SET reaction_count = reaction_count + 1,
            reactions_summary = jsonb_set(
                COALESCE(reactions_summary, '{}'::jsonb),
                ARRAY[NEW.emoji],
                to_jsonb(COALESCE((reactions_summary->>NEW.emoji)::int, 0) + 1)
            )
        WHERE id = NEW.message_id;
    ELSIF TG_OP = 'DELETE' THEN
        UPDATE messages
        SET reaction_count = GREATEST(reaction_count - 1, 0),
            reactions_summary = CASE
                WHEN (reactions_summary->>OLD.emoji)::int > 1 THEN
                    jsonb_set(reactions_summary, ARRAY[OLD.emoji], to_jsonb((reactions_summary->>OLD.emoji)::int - 1))
                ELSE
                    reactions_summary - OLD.emoji
            END
        WHERE id = OLD.message_id;
    END IF;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_update_message_reaction_count
AFTER INSERT OR DELETE ON reactions
FOR EACH ROW EXECUTE FUNCTION update_message_reaction_count();

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_conversations_updated_at BEFORE UPDATE ON conversations FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER trg_participants_updated_at BEFORE UPDATE ON participants FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER trg_messages_updated_at BEFORE UPDATE ON messages FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER trg_notifications_updated_at BEFORE UPDATE ON notifications FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

```
=========================================
##  SECTION 30: DATABASE STATISTICS

```sql
-- =========================================

CREATE OR REPLACE VIEW v_table_sizes AS
SELECT 
    schemaname,
    tablename,
    pg_size_pretty(pg_total_relation_size(schemaname||'.'||tablename)) AS size,
    pg_total_relation_size(schemaname||'.'||tablename) AS size_bytes
FROM pg_tables
WHERE schemaname = 'public'
ORDER BY pg_total_relation_size(schemaname||'.'||tablename) DESC;

CREATE OR REPLACE VIEW v_index_usage AS
SELECT
    schemaname,
    tablename,
    indexname,
    idx_scan AS index_scans,
    idx_tup_read AS tuples_read,
    idx_tup_fetch AS tuples_fetched,
    pg_size_pretty(pg_relation_size(indexrelid)) AS index_size
FROM pg_stat_user_indexes
WHERE schemaname = 'public'
ORDER BY idx_scan DESC;

```
=========================================
## COMMENTS (from existing)

```sql

COMMENT ON TABLE threads IS 'Message threads - maps to internal/domain/thread/entity.go';
COMMENT ON TABLE thread_followers IS 'Thread followers for notifications';
COMMENT ON TABLE message_edit_history IS 'Message edit audit trail';
COMMENT ON TABLE reactions IS 'Message reactions - maps to internal/domain/message/reaction.go';
COMMENT ON TABLE attachments IS 'Message attachments - maps to internal/domain/message/attachment.go';
COMMENT ON TABLE read_receipts IS 'Message read receipts - maps to internal/domain/message/read_receipt.go';
COMMENT ON TABLE mentions IS 'User mentions - maps to internal/domain/mention/entity.go';
COMMENT ON TABLE conversation_exports IS 'Conversation exports - maps to internal/domain/export/entity.go';
COMMENT ON TABLE notification_preferences IS 'User notification preferences - maps to internal/domain/notification/preference.go';
COMMENT ON TABLE email_notifications IS 'Email notifications - maps to internal/domain/email/entity.go';
COMMENT ON TABLE email_templates IS 'Email templates for notifications';
COMMENT ON TABLE push_notifications IS 'Push notifications - maps to internal/domain/push/entity.go';
COMMENT ON TABLE push_subscriptions IS 'WebPush subscriptions';
COMMENT ON TABLE sms_notifications IS 'SMS notifications - maps to internal/domain/sms/entity.go';
COMMENT ON TABLE websocket_connections IS 'WebSocket connections - maps to internal/domain/websocket/entity.go';
COMMENT ON TABLE user_presence IS 'User presence status - maps to internal/domain/presence/entity.go';
COMMENT ON TABLE presence_history IS 'Presence status change history';
COMMENT ON TABLE collaboration_sessions IS 'Real-time collaboration sessions - maps to internal/domain/collaboration/entity.go';
COMMENT ON TABLE collaboration_participants IS 'Collaboration session participants';
COMMENT ON TABLE document_versions IS 'Document version history';
COMMENT ON TABLE email_digests IS 'Email digest notifications - maps to internal/domain/digest/entity.go';
COMMENT ON TABLE webhooks IS 'Webhook subscriptions - maps to internal/domain/webhook/entity.go';
COMMENT ON TABLE webhook_deliveries IS 'Webhook delivery attempts';
COMMENT ON TABLE message_search_index IS 'Message search index - maps to internal/domain/search/entity.go';
COMMENT ON TABLE retention_policies IS 'Data retention policies - maps to internal/domain/compliance/entity.go';
COMMENT ON TABLE gdpr_requests IS 'GDPR data subject requests';
COMMENT ON TABLE rate_limits IS 'Rate limiting quotas - maps to internal/domain/quota/entity.go';
COMMENT ON TABLE usage_metrics IS 'Usage metrics aggregation';
COMMENT ON TABLE outbox_events IS 'Transactional outbox for event publishing';
COMMENT ON TABLE dead_letter_queue IS 'Dead letter queue for failed events';
COMMENT ON TABLE audit_logs IS 'Comprehensive audit trail';
COMMENT ON TABLE conversation_read_model IS 'CQRS read model for conversations';
COMMENT ON TABLE user_notification_summary IS 'User notification summary for quick access';
COMMENT ON TABLE communication_metrics IS 'Communication analytics and metrics';
```

---

## SECTION 31+ (Added domains to cover all folders)



=========================================
##  SECTION 31: draft/
```sql

CREATE TABLE draft (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID NOT NULL,
    user_id UUID NOT NULL,
    content TEXT NOT NULL,
    content_format VARCHAR(20) DEFAULT 'PLAIN' CHECK (content_format IN ('PLAIN','MARKDOWN','HTML','RICH')),
    attachments JSONB,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uk_draft_conv_user UNIQUE (conversation_id, user_id),
    CONSTRAINT fk_draft_conversation FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
);
CREATE INDEX idx_draft_user ON draft (user_id, updated_at DESC);

```
=========================================
##  SECTION 32: pin/
```sql

CREATE TABLE pin (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID NOT NULL,
    message_id UUID NOT NULL,
    pinned_by UUID NOT NULL,
    note VARCHAR(500),
    pinned_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uk_pin UNIQUE (conversation_id, message_id),
    CONSTRAINT fk_pin_conversation FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE,
    CONSTRAINT fk_pin_message FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE
);
CREATE INDEX idx_pin_conv_time ON pin (conversation_id, pinned_at DESC);
CREATE INDEX idx_pin_by ON pin (pinned_by, pinned_at DESC);

```
=========================================
##  SECTION 33: bookmark/

```sql
CREATE TABLE bookmark (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    message_id UUID NOT NULL,
    note VARCHAR(500),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uk_bookmark_user_msg UNIQUE (user_id, message_id),
    CONSTRAINT fk_bookmark_message FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE
);
CREATE INDEX idx_bookmark_user_time ON bookmark (user_id, created_at DESC);

```
=========================================
##  SECTION 34: read_state/
```sql

CREATE TABLE read_state (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID NOT NULL,
    user_id UUID NOT NULL,
    last_read_seq BIGINT NOT NULL DEFAULT 0,
    last_read_message_id UUID,
    last_read_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uk_read_state UNIQUE (conversation_id, user_id),
    CONSTRAINT fk_read_state_conversation FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
);
CREATE INDEX idx_read_state_user ON read_state (user_id, updated_at DESC);

```
=========================================
##  SECTION 35: delivery/
```sql

CREATE TABLE delivery (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    message_id UUID NOT NULL,
    user_id UUID NOT NULL,
    device_id UUID,
    session_id UUID,
    channel VARCHAR(20) NOT NULL CHECK (channel IN ('IN_APP','WEBPUSH','EMAIL','SMS','PUSH')),
    status VARCHAR(20) NOT NULL DEFAULT 'QUEUED' CHECK (status IN ('QUEUED','DISPATCHED','ACK','FAILED','EXPIRED')),
    attempt_count INTEGER NOT NULL DEFAULT 0,
    first_dispatched_at TIMESTAMPTZ,
    last_dispatched_at TIMESTAMPTZ,
    acked_at TIMESTAMPTZ,
    expired_at TIMESTAMPTZ,
    error_code VARCHAR(50),
    error_message TEXT,
    metadata JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_delivery_message FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE
);
CREATE INDEX idx_delivery_user_status ON delivery (user_id, status, updated_at DESC);
CREATE INDEX idx_delivery_msg ON delivery (message_id, channel, status);

```
=========================================
##  SECTION 36: notification_queue/
```sql

CREATE TABLE notification_queue (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    notification_id UUID NOT NULL,
    channel VARCHAR(20) NOT NULL CHECK (channel IN ('IN_APP','WEBPUSH','EMAIL','SMS','PUSH')),
    priority VARCHAR(20) NOT NULL DEFAULT 'NORMAL' CHECK (priority IN ('LOW','NORMAL','HIGH','URGENT')),
    scheduled_for TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING' CHECK (status IN ('PENDING','ENQUEUED','DEQUEUED','PROCESSING','SENT','FAILED','DEADLETTER')),
    retries INTEGER NOT NULL DEFAULT 0,
    max_retries INTEGER NOT NULL DEFAULT 5,
    partition_key VARCHAR(255),
    data_zone VARCHAR(10) DEFAULT 'US' CHECK (data_zone IN ('US','EU','ASIA')),
    metadata JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_notification_queue_notification FOREIGN KEY (notification_id) REFERENCES notifications(id) ON DELETE CASCADE
);
CREATE INDEX idx_nq_when ON notification_queue (status, scheduled_for);
CREATE INDEX idx_nq_notif ON notification_queue (notification_id, channel);
CREATE INDEX idx_nq_priority ON notification_queue (priority, scheduled_for DESC) WHERE status IN ('PENDING','ENQUEUED','PROCESSING');

```
=========================================
##  SECTION 37: delivery_log/
```sql

CREATE TABLE delivery_log (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    channel VARCHAR(20) NOT NULL CHECK (channel IN ('IN_APP','WEBPUSH','EMAIL','SMS','PUSH','WEBHOOK')),
    entity_type VARCHAR(30) NOT NULL,
    entity_id UUID NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('QUEUED','SENT','DELIVERED','BOUNCED','FAILED','EXPIRED','CLICKED','OPENED')),
    http_status INTEGER,
    latency_ms INTEGER,
    provider_id VARCHAR(255),
    error_code VARCHAR(50),
    error_message TEXT,
    occurred_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    metadata JSONB
);
CREATE INDEX idx_dl_entity ON delivery_log (entity_type, entity_id, occurred_at DESC);
CREATE INDEX idx_dl_channel_status ON delivery_log (channel, status, occurred_at DESC);

```
=========================================
##  SECTION 38: moderation/
```sql

CREATE TABLE moderation (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    scope VARCHAR(20) NOT NULL CHECK (scope IN ('MESSAGE','CONVERSATION')),
    target_id UUID NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING' CHECK (status IN ('PENDING','REVIEWING','RESOLVED','REJECTED')),
    resolution VARCHAR(30),
    resolved_by UUID,
    resolved_at TIMESTAMPTZ,
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_moderation_target ON moderation (scope, target_id, created_at DESC);
CREATE INDEX idx_moderation_status ON moderation (status, created_at DESC);

CREATE TABLE moderation_flag (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    moderation_id UUID NOT NULL,
    reporter_id UUID NOT NULL,
    reason VARCHAR(50) NOT NULL,
    details TEXT,
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING' CHECK (status IN ('PENDING','REVIEWED','DISMISSED')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_mflag_case FOREIGN KEY (moderation_id) REFERENCES moderation(id) ON DELETE CASCADE
);
CREATE INDEX idx_mflag_case ON moderation_flag (moderation_id, status);

CREATE TABLE moderation_quarantine (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    moderation_id UUID NOT NULL,
    started_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ended_at TIMESTAMPTZ,
    reason VARCHAR(100),
    CONSTRAINT fk_mq_case FOREIGN KEY (moderation_id) REFERENCES moderation(id) ON DELETE CASCADE
);

```
=========================================
##  SECTION 39: blocklist/
```sql

CREATE TABLE blocklist (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    scope VARCHAR(20) NOT NULL CHECK (scope IN ('USER','TENANT','GLOBAL')),
    subject_type VARCHAR(20) NOT NULL CHECK (subject_type IN ('USER','PHRASE','DOMAIN','URL')),
    subject_value TEXT NOT NULL,
    reason VARCHAR(200),
    created_by UUID,
    expires_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uk_blocklist UNIQUE (scope, subject_type, subject_value)
);
CREATE INDEX idx_blocklist_exp ON blocklist (expires_at);

```
=========================================
##  SECTION 40: url_safety/
```sql

CREATE TABLE url_safety (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    url_hash VARCHAR(64) NOT NULL UNIQUE,
    reputation VARCHAR(20) NOT NULL DEFAULT 'UNKNOWN' CHECK (reputation IN ('UNKNOWN','SAFE','SUSPICIOUS','MALICIOUS')),
    source VARCHAR(50),
    last_scanned_at TIMESTAMPTZ,
    ttl_seconds INTEGER DEFAULT 86400,
    metadata JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_url_safety_rep ON url_safety (reputation, updated_at DESC);

```
=========================================
##  SECTION 41: suppression/
```sql

CREATE TABLE suppression (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    channel VARCHAR(10) NOT NULL CHECK (channel IN ('EMAIL','SMS')),
    address_hash VARCHAR(64) NOT NULL,
    domain VARCHAR(255),
    reason VARCHAR(20) NOT NULL CHECK (reason IN ('BOUNCE','COMPLAINT','INVALID','MANUAL')),
    details TEXT,
    expires_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uk_suppression UNIQUE (channel, address_hash)
);
CREATE INDEX idx_suppression_exp ON suppression (expires_at);

```
=========================================
##  SECTION 42: unsubscribe/
```sql

CREATE TABLE unsubscribe (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID,
    channel VARCHAR(10) NOT NULL CHECK (channel IN ('EMAIL','SMS','PUSH','IN_APP','WEBPUSH')),
    topic VARCHAR(50) NOT NULL,
    reason VARCHAR(200),
    unsubscribed_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uk_unsubscribe UNIQUE (COALESCE(user_id, '00000000-0000-0000-0000-000000000000'::uuid), channel, topic)
);
CREATE INDEX idx_unsubscribe_user ON unsubscribe (user_id, channel, topic);

```
=========================================
##  SECTION 43: encryption/
```sql

CREATE TABLE encryption (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID NOT NULL UNIQUE,
    enabled BOOLEAN NOT NULL DEFAULT FALSE,
    policy VARCHAR(20) DEFAULT 'STRICT' CHECK (policy IN ('STRICT','LENIENT')),
    search_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    automod_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_encryption_conversation FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
);

CREATE TABLE encryption_conversation_key (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID NOT NULL,
    key_version INTEGER NOT NULL,
    kms_key_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    rotated_at TIMESTAMPTZ,
    CONSTRAINT uk_enc_room_version UNIQUE (conversation_id, key_version),
    CONSTRAINT fk_enc_roomkey_conversation FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
);

CREATE TABLE encryption_participant_key (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID NOT NULL,
    user_id UUID NOT NULL,
    public_key TEXT NOT NULL,
    added_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    revoked_at TIMESTAMPTZ,
    CONSTRAINT uk_enc_participant UNIQUE (conversation_id, user_id),
    CONSTRAINT fk_enc_part_conversation FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
);
CREATE INDEX idx_enc_part_user ON encryption_participant_key (user_id, added_at DESC);

```
=========================================
##  SECTION 44: in_app_notification/
```sql

CREATE TABLE in_app_notification (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    notification_id UUID NOT NULL,
    user_id UUID NOT NULL,
    displayed_at TIMESTAMPTZ,
    dismissed_at TIMESTAMPTZ,
    clicked_at TIMESTAMPTZ,
    badge_count_delta INTEGER DEFAULT 0,
    cta_label VARCHAR(100),
    cta_url TEXT,
    cta_action JSONB,
    metadata JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uk_inapp UNIQUE (notification_id, user_id),
    CONSTRAINT fk_inapp_notification FOREIGN KEY (notification_id) REFERENCES notifications(id) ON DELETE CASCADE
);
CREATE INDEX idx_inapp_user ON in_app_notification (user_id, created_at DESC);

```
=========================================
##  SECTION 45: notification_template/
```sql

CREATE TABLE notification_template (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    key VARCHAR(100) NOT NULL,
    channel VARCHAR(20) NOT NULL CHECK (channel IN ('IN_APP','WEBPUSH','PUSH')),
    locale VARCHAR(10) NOT NULL DEFAULT 'en',
    version INTEGER NOT NULL DEFAULT 1,
    title VARCHAR(200),
    body TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    variables JSONB,
    metadata JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uk_ntemplate UNIQUE (key, channel, locale, version)
);
CREATE INDEX idx_ntemplate_active ON notification_template (key, channel, locale) WHERE is_active = TRUE;

```
=========================================
##  SECTION 46: push_device/
```sql

CREATE TABLE push_device (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    platform VARCHAR(10) NOT NULL CHECK (platform IN ('IOS','ANDROID')),
    device_token TEXT NOT NULL,
    device_id UUID,
    app_version VARCHAR(50),
    locale VARCHAR(10),
    last_used_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uk_push_device UNIQUE (platform, device_token)
);
CREATE INDEX idx_push_device_user ON push_device (user_id, last_used_at DESC);

```
=========================================
##  SECTION 47: system_message/
```sql

CREATE TABLE system_message (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID NOT NULL,
    kind VARCHAR(30) NOT NULL,
    payload JSONB NOT NULL,
    created_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_sysmsg_conversation FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
);
CREATE INDEX idx_sysmsg_conv_time ON system_message (conversation_id, created_at DESC);

```
=========================================
##  SECTION 48: call/
```sql

CREATE TABLE call (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID,
    organizer_id UUID NOT NULL,
    starts_at TIMESTAMPTZ NOT NULL,
    ends_at TIMESTAMPTZ,
    status VARCHAR(20) NOT NULL DEFAULT 'SCHEDULED' CHECK (status IN ('SCHEDULED','STARTED','ENDED','CANCELED','MISSED')),
    link TEXT,
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_call_conv_time ON call (conversation_id, starts_at DESC);
CREATE INDEX idx_call_org_time ON call (organizer_id, starts_at DESC);

```
=========================================
##  SECTION 49: calendar_invite/
```sql

CREATE TABLE calendar_invite (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    call_id UUID,
    ical TEXT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'SENT' CHECK (status IN ('SENT','ACCEPTED','DECLINED','TENTATIVE','BOUNCED')),
    sent_to JSONB,
    sent_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    responded_at TIMESTAMPTZ,
    metadata JSONB,
    CONSTRAINT fk_calinvite_call FOREIGN KEY (call_id) REFERENCES call(id) ON DELETE CASCADE
);
CREATE INDEX idx_calinvite_status ON calendar_invite (status, sent_at DESC);

```
=========================================
##  SECTION 50: email_bridge/
```sql

CREATE TABLE email_bridge (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID NOT NULL,
    inbound_address VARCHAR(320) NOT NULL,
    message_id_header VARCHAR(255),
    created_by UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMPTZ,
    CONSTRAINT uk_email_bridge_addr UNIQUE (inbound_address),
    CONSTRAINT fk_ebridge_conversation FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
);
CREATE INDEX idx_ebridge_exp ON email_bridge (expires_at);

CREATE TABLE email_bridge_inbound (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email_bridge_id UUID NOT NULL,
    provider_message_id VARCHAR(255),
    from_address VARCHAR(320) NOT NULL,
    subject TEXT,
    body_text TEXT,
    body_html TEXT,
    received_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    parsed_ok BOOLEAN DEFAULT TRUE,
    error_message TEXT,
    CONSTRAINT fk_ebridge_inbound FOREIGN KEY (email_bridge_id) REFERENCES email_bridge(id) ON DELETE CASCADE
);
CREATE INDEX idx_ebridge_inbound_time ON email_bridge_inbound (received_at DESC);

```
=========================================
##  SECTION 51: mail_tracking/
```sql

CREATE TABLE mail_tracking (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email_notification_id UUID,
    event_type VARCHAR(20) NOT NULL CHECK (event_type IN ('DELIVERED','DEFERRED','BOUNCED','COMPLAINED','OPENED','CLICKED')),
    provider_message_id VARCHAR(255),
    ts TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    details JSONB,
    CONSTRAINT fk_mailtrack_email FOREIGN KEY (email_notification_id) REFERENCES email_notifications(id) ON DELETE SET NULL
);
CREATE INDEX idx_mailtrack_event ON mail_tracking (event_type, ts DESC);
CREATE INDEX idx_mailtrack_provider ON mail_tracking (provider_message_id);

```
=========================================
##  SECTION 52: interview/
```sql

CREATE TABLE interview (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID,
    client_id UUID NOT NULL,
    freelancer_id UUID NOT NULL,
    scheduled_at TIMESTAMPTZ NOT NULL,
    scheduled_tz VARCHAR(50),
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING' CHECK (status IN ('PENDING','CONFIRMED','CANCELED','COMPLETED','NO_SHOW')),
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_interview_time ON interview (scheduled_at DESC, status);
CREATE INDEX idx_interview_parties ON interview (client_id, freelancer_id, scheduled_at DESC);

CREATE TABLE interview_participant (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    interview_id UUID NOT NULL,
    user_id UUID NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('CLIENT','FREELANCER','OBSERVER')),
    availability JSONB,
    joined_at TIMESTAMPTZ,
    left_at TIMESTAMPTZ,
    CONSTRAINT uk_interview_participant UNIQUE (interview_id, user_id),
    CONSTRAINT fk_interview_part FOREIGN KEY (interview_id) REFERENCES interview(id) ON DELETE CASCADE
);
CREATE INDEX idx_interview_part_user ON interview_participant (user_id, interview_id);

```
=========================================
##  SECTION 53: platform_alert/
```sql

CREATE TABLE platform_alert (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    severity VARCHAR(10) NOT NULL CHECK (severity IN ('INFO','WARN','ERROR','CRITICAL')),
    target VARCHAR(20) NOT NULL CHECK (target IN ('ALL','FREELANCERS','CLIENTS','ADMINS','TENANT')),
    tenant_id UUID,
    title VARCHAR(200) NOT NULL,
    message TEXT NOT NULL,
    starts_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ends_at TIMESTAMPTZ,
    status VARCHAR(20) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','EXPIRED','CANCELED')),
    metadata JSONB,
    created_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_palert_window ON platform_alert (status, starts_at, ends_at);

CREATE TABLE platform_alert_delivery (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    platform_alert_id UUID NOT NULL,
    user_id UUID NOT NULL,
    delivered_at TIMESTAMPTZ,
    dismissed_at TIMESTAMPTZ,
    channel VARCHAR(20),
    CONSTRAINT uk_palert_delivery UNIQUE (platform_alert_id, user_id),
    CONSTRAINT fk_palert_delivery FOREIGN KEY (platform_alert_id) REFERENCES platform_alert(id) ON DELETE CASCADE
);
CREATE INDEX idx_palert_delivery_user ON platform_alert_delivery (user_id, delivered_at DESC);

```
=========================================
##  SECTION 54: spam_detection/
```sql

CREATE TABLE spam_detection (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    scope VARCHAR(20) NOT NULL CHECK (scope IN ('MESSAGE','EMAIL')),
    target_id UUID,
    spam_score NUMERIC(5,2) NOT NULL,
    detected_patterns JSONB,
    action VARCHAR(20) NOT NULL CHECK (action IN ('NONE','FLAG','QUARANTINE','BLOCK')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_spamdet_scope ON spam_detection (scope, target_id, created_at DESC);
CREATE INDEX idx_spamdet_score ON spam_detection (spam_score DESC, created_at DESC);

CREATE TABLE spam_detection_rule (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    pattern TEXT NOT NULL,
    severity VARCHAR(10) NOT NULL DEFAULT 'LOW' CHECK (severity IN ('LOW','MEDIUM','HIGH','CRITICAL')),
    action VARCHAR(20) NOT NULL DEFAULT 'FLAG' CHECK (action IN ('FLAG','QUARANTINE','BLOCK')),
    enabled BOOLEAN NOT NULL DEFAULT TRUE,
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_spamrule_enabled ON spam_detection_rule (enabled, severity);

```
=========================================
##  SECTION 55: Extra updated_at triggers for new tables

```sql

-- =========================================
CREATE TRIGGER trg_draft_updated_at                 BEFORE UPDATE ON draft                 FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER trg_delivery_updated_at              BEFORE UPDATE ON delivery              FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER trg_read_state_updated_at            BEFORE UPDATE ON read_state            FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER trg_encryption_updated_at            BEFORE UPDATE ON encryption            FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER trg_ntemplate_updated_at             BEFORE UPDATE ON notification_template FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER trg_url_safety_updated_at            BEFORE UPDATE ON url_safety            FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER trg_push_device_updated_at           BEFORE UPDATE ON push_device           FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
```

---

## Comments for new domain tables

```sql
COMMENT ON TABLE draft                     IS 'User drafts per conversation - maps to internal/domain/draft/';
COMMENT ON TABLE pin                       IS 'Pinned messages per conversation - maps to internal/domain/pin/';
COMMENT ON TABLE bookmark                  IS 'User bookmarks of messages - maps to internal/domain/bookmark/';
COMMENT ON TABLE read_state                IS 'Per-user last read pointers - maps to internal/domain/read_state/';
COMMENT ON TABLE delivery                  IS 'Per-device/server→device deliveries - maps to internal/domain/delivery/';
COMMENT ON TABLE notification_queue        IS 'Queue for outbound notifications - maps to internal/domain/notification_queue/';
COMMENT ON TABLE delivery_log              IS 'Cross-channel delivery telemetry - maps to internal/domain/delivery_log/';
COMMENT ON TABLE moderation                IS 'Moderation cases - maps to internal/domain/moderation/';
COMMENT ON TABLE moderation_flag           IS 'User-submitted flags - maps to internal/domain/moderation/';
COMMENT ON TABLE moderation_quarantine     IS 'Quarantine windows - maps to internal/domain/moderation/';
COMMENT ON TABLE blocklist                 IS 'Blocklist entries - maps to internal/domain/blocklist/';
COMMENT ON TABLE url_safety                IS 'URL reputation cache - maps to internal/domain/url_safety/';
COMMENT ON TABLE suppression               IS 'Channel suppression (bounce/complaint) - maps to internal/domain/suppression/';
COMMENT ON TABLE unsubscribe               IS 'Opt-out registry by user/channel/topic - maps to internal/domain/unsubscribe/';
COMMENT ON TABLE encryption                IS 'E2EE settings per conversation - maps to internal/domain/encryption/';
COMMENT ON TABLE encryption_conversation_key IS 'Rotating room keys - maps to internal/domain/encryption/';
COMMENT ON TABLE encryption_participant_key IS 'Participant public keys - maps to internal/domain/encryption/';
COMMENT ON TABLE in_app_notification       IS 'Per-user display/CTA state - maps to internal/domain/in_app_notification/';
COMMENT ON TABLE notification_template     IS 'Templates for in-app/push - maps to internal/domain/notification_template/';
COMMENT ON TABLE push_device               IS 'Native push device registry - maps to internal/domain/push_device/';
COMMENT ON TABLE system_message            IS 'Structured system feed - maps to internal/domain/system_message/';
COMMENT ON TABLE call                      IS 'Scheduled/live calls - maps to internal/domain/call/';
COMMENT ON TABLE calendar_invite           IS 'Calendar invites (iCal) - maps to internal/domain/calendar_invite/';
COMMENT ON TABLE email_bridge              IS 'Inbound email mapping/aliases - maps to internal/domain/email_bridge/';
COMMENT ON TABLE email_bridge_inbound      IS 'Inbound email log - maps to internal/domain/email_bridge/';
COMMENT ON TABLE mail_tracking             IS 'Provider event logs - maps to internal/domain/mail_tracking/';
COMMENT ON TABLE interview                 IS 'Interview slots - maps to internal/domain/interview/';
COMMENT ON TABLE interview_participant     IS 'Interview attendees - maps to internal/domain/interview/';
COMMENT ON TABLE platform_alert            IS 'Broadcast/platform alerts - maps to internal/domain/platform_alert/';
COMMENT ON TABLE platform_alert_delivery   IS 'Per-user alert deliveries - maps to internal/domain/platform_alert/';
COMMENT ON TABLE spam_detection            IS 'Spam detection outcomes - maps to internal/domain/spam_detection/';
COMMENT ON TABLE spam_detection_rule       IS 'Spam rule set - maps to internal/domain/spam_detection/';
```

---

## Final summary

- **Total Domains Covered:** conversation, thread, message, mention, export, notification, notification_preferences, email, push, sms, websocket, presence, collaboration, digest, webhook, search, compliance, quota, outbox, analytics, views, stats, plus the **newly added**: **draft, pin, bookmark, read_state, delivery, notification_queue, delivery_log, moderation, blocklist, url_safety, suppression, unsubscribe, encryption, in_app_notification, notification_template, push_device, system_message, call, calendar_invite, email_bridge, mail_tracking, interview, platform_alert, spam_detection**.  
- **Alignment:** every `internal/domain/{domain}/` has **one main table** named exactly `{domain}`, with sub-entities as `{domain}_{sub}`.  
- **Scale:** rich fields, indexes, queues, logs, suppression/unsubscribe registries, E2EE keys, device registries, and provider telemetry added for an Upwork-scale workload.  
- **Compliance & Ops:** GDPR zones, suppression + unsubscribe, moderation + spam, delivery telemetry, outbox + DLQ, audit, and dashboards.  

---

## 📌 Additions (Large-Scale Completeness)

Below are the **new tables** added to fully align with the folder structure and production needs.  
Each block includes the **section number (suggested placement)** and the **domain path** to maintain the CRITICAL ALIGNMENT RULES.

### ✅ Where to place them in the main doc
- **Section 31 — Domain: `internal/domain/idempotency/`** → Main table: `idempotency` (request key registry for exactly-once semantics).  
  _Tip:_ In the main document, you can place this right after **Section 23: Rate Limiting & Quotas** and **before the Outbox** (since idempotency is often checked before emitting events). If you prefer to keep numbering stable, keep it as **Section 31**.
- **Section 32 — Domain: `internal/domain/message/` (sub-entity: `sequence.go`)** → Sub-table: `message_sequence` (per-conversation monotonic allocator).  
  _Tip:_ You may nest this logically under the **Messages** section (as a sub-entity) or keep it as **Section 32** for clarity.

---

=========================================
### SECTION 31: IDEMPOTENCY KEYS
```sql
-- Domain: internal/domain/idempotency/
### -- Entity: idempotency/key.go
### -- =========================================


CREATE TABLE idempotency (
    -- Primary Key
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    -- Identity
    scope VARCHAR(100) NOT NULL,                     -- e.g., 'message.send', 'notification.send'
    idempotency_key VARCHAR(255) NOT NULL,           -- request key from client / upstream

    -- Actor / Context
    user_id UUID,
    tenant_id UUID,
    client_ip INET,
    request_method VARCHAR(10),
    request_path TEXT,
    request_headers JSONB,
    request_body_hash VARCHAR(64),                   -- SHA-256 of normalized request body/payload
    checksum VARCHAR(64),                            -- additional checksum/fingerprint if needed

    -- Lifecycle
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING' CHECK (
        status IN ('PENDING','COMPLETED','FAILED','EXPIRED')
    ),
    first_seen_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_seen_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMPTZ NOT NULL,

    -- Concurrency guard (optional)
    locked_by UUID,
    locked_at TIMESTAMPTZ,
    lock_expires_at TIMESTAMPTZ,

    -- Response snapshot (optional, for safe retries)
    response_status INTEGER,
    response_headers JSONB,
    response_body_ref UUID,                          -- reference to storage-be for large responses
    response_body_preview TEXT,                      -- truncated preview for quick debugging

    -- Metadata
    metadata JSONB,

    -- Audit
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT uk_idempotency_scope_key UNIQUE (scope, idempotency_key)
);

CREATE INDEX idx_idempotency_expires ON idempotency (expires_at);
CREATE INDEX idx_idempotency_status ON idempotency (status, expires_at);
CREATE INDEX idx_idempotency_actor ON idempotency (tenant_id, user_id);
CREATE INDEX idx_idempotency_path ON idempotency (request_method, request_path) WHERE status <> 'EXPIRED';

COMMENT ON TABLE idempotency IS 'Idempotency request key registry - maps to internal/domain/idempotency/';
COMMENT ON COLUMN idempotency.scope IS 'Functional scope for the key (ex: message.send, notification.send)';
COMMENT ON COLUMN idempotency.idempotency_key IS 'Client-supplied or server-generated idempotency key';

-- Auto-update updated_at
CREATE TRIGGER trg_idempotency_updated_at
BEFORE UPDATE ON idempotency
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
```

---

=========================================
### SECTION 32: MESSAGE SEQUENCE ALLOCATOR (SUB-ENTITY)
```sql
-- Domain: internal/domain/message/ (sub-entity: sequence.go)
### -- =========================================


CREATE TABLE message_sequence (
    -- One row per conversation
    conversation_id UUID PRIMARY KEY,
    next_seq BIGINT NOT NULL DEFAULT 1,              -- next sequence to allocate (monotonic per conversation)

    -- Allocator diagnostics
    allocator_node VARCHAR(100),                     -- which node/process allocated last
    last_allocated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version INTEGER NOT NULL DEFAULT 1,

    -- Audit
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_message_sequence_conversation FOREIGN KEY (conversation_id)
        REFERENCES conversations(id) ON DELETE CASCADE
);

CREATE INDEX idx_message_sequence_updated_at ON message_sequence (updated_at DESC);

COMMENT ON TABLE message_sequence IS 'Per-conversation monotonic sequence allocator - maps to internal/domain/message/sequence.go';

-- Auto-update updated_at
CREATE TRIGGER trg_message_sequence_updated_at
BEFORE UPDATE ON message_sequence
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Optional helper for atomic allocation (safe across replicas)
CREATE OR REPLACE FUNCTION reserve_next_message_seq(p_conversation_id UUID)
RETURNS BIGINT AS $$
DECLARE
    v_next BIGINT;
BEGIN
    LOOP
        UPDATE message_sequence
        SET next_seq = next_seq + 1,
            last_allocated_at = CURRENT_TIMESTAMP
        WHERE conversation_id = p_conversation_id
        RETURNING next_seq - 1 INTO v_next;
        IF FOUND THEN
            RETURN v_next;
        END IF;

        -- If row missing, create a seed record and retry
        BEGIN
            INSERT INTO message_sequence (conversation_id, next_seq)
            VALUES (p_conversation_id, 1)
            ON CONFLICT (conversation_id) DO NOTHING;
        EXCEPTION WHEN unique_violation THEN
            -- ignore: another allocator raced us; loop and retry
        END;
    END LOOP;
END;
$$ LANGUAGE plpgsql;
```

---
## Final summary

*   **Total Domains Covered:** 
- **conversation, thread, message, mention, export, notification, notification\_preferences, email, push, sms, websocket, presence, collaboration, digest, webhook, search, compliance, quota, outbox, analytics, views, stats, plus the newly added: draft, pin, bookmark, read\_state, delivery, notification\_queue, delivery\_log, moderation, blocklist, url\_safety, suppression, unsubscribe, encryption, in\_app\_notification, notification\_template, push\_device, system\_message, call, calendar\_invite, email\_bridge, mail\_tracking, interview, platform\_alert, spam\_detection.**
    
*   **Alignment:** Every internal/domain/{domain}/ has **one main table** named exactly {domain}, with sub-entities as {domain}\_{sub}.
    
*   **Scale:** Rich fields, indexes, queues, logs, suppression/unsubscribe registries, E2EE keys, device registries, and provider telemetry added for an **Upwork-scale workload**.
    
*   **Compliance & Ops:** **GDPR zones**, suppression + unsubscribe, moderation + spam, delivery telemetry, outbox + DLQ, audit, and dashboards.
### 🔄 Summary Addendum

- **Added Tables**
  - `idempotency` (Domain: `internal/domain/idempotency/`) — **Section 31** (suggested to sit before Outbox in the main doc).
  - `message_sequence` (Domain: `internal/domain/message/`, sub-entity: `sequence.go`) — **Section 32** (or embed under the Messages section as a sub-entity block).

- **Alignment Updates**
  - ✅ **idempotency/** → `idempotency` table (request key registry for exactly-once semantics).
  - ✅ **message/sequence.go** → `message_sequence` table (per-conversation monotonic allocator).

- **Why these matter at scale**
  - `idempotency` prevents duplicate effects on retried requests and enables **exactly-once** semantics across retries/timeouts.
  - `message_sequence` ensures **global per-room ordering** across multiple replicas, keeping the monotonic `seq` allocation contention-free and observable.
