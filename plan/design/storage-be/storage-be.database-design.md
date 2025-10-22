# STORAGE-BE DATABASE DESIGN (Comprehensive, Production-Ready)
**Skillsier Platform – Enterprise Scale (Upwork-like)**
**PostgreSQL 16+**

> This schema aligns 1:1 with your `apps/be/storage-be/` folder structure and the consolidated Storage user stories.
> It strictly follows your **CRITICAL ALIGNMENT RULES** and is ready for large scale, mobile + web, and cross‑service integrations.

---

## CRITICAL ALIGNMENT RULES

1. Each domain folder in `internal/domain/{domain}/` = **ONE** main table named **exactly** `{domain}`.  
2. Table names match domain folder names **exactly**.  
3. Sub-entities create related tables with `{domain}_{sub}` naming.  
4. **All** domains from the folder structure are covered.  
5. Rich, production-ready fields, strong constraints, indexes, and partitioning for scale.

---

## Global Extensions & Settings

```sql
-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";
CREATE EXTENSION IF NOT EXISTS "btree_gin";
CREATE EXTENSION IF NOT EXISTS "citext";
CREATE EXTENSION IF NOT EXISTS "btree_gist";

-- Time zone and timestamps
SET TIME ZONE 'UTC';
```

---

## Global Enums (Domain-Scoped)

```sql
-- ===== FILE PRIMITIVES =====
DO $$ BEGIN
  CREATE TYPE visibility_t AS ENUM ('private','team','org','tenant','public');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;

DO $$ BEGIN
  CREATE TYPE file_status_t AS ENUM ('active','soft_deleted','quarantined','archived');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;

DO $$ BEGIN
  CREATE TYPE file_type_t AS ENUM ('generic','image','video','audio','document','archive','code','dataset','other');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;

-- ===== ACCESS / SHARING =====
DO $$ BEGIN
  CREATE TYPE share_scope_t AS ENUM ('view','download','comment','edit','owner');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;

DO $$ BEGIN
  CREATE TYPE link_action_t AS ENUM ('view','download','upload');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;

-- ===== CONTENT PIPELINE =====
DO $$ BEGIN
  CREATE TYPE media_kind_t AS ENUM ('image','video');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;

DO $$ BEGIN
  CREATE TYPE job_status_t AS ENUM ('queued','running','succeeded','failed','canceled','expired');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;

-- ===== SAFETY & POLICY =====
DO $$ BEGIN
  CREATE TYPE severity_t AS ENUM ('info','low','medium','high','critical');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;

-- ===== COMPLIANCE & LIFECYCLE =====
DO $$ BEGIN
  CREATE TYPE legal_hold_state_t AS ENUM ('active','released');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;

-- ===== STORAGE PRIMITIVES & OPS =====
DO $$ BEGIN
  CREATE TYPE storage_class_t AS ENUM ('standard','infrequent_access','archive');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;

DO $$ BEGIN
  CREATE TYPE data_zone_t AS ENUM ('EU','US','APAC','ME','GLOBAL');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;
```

---

## GLOBAL TABLES (Outbox / Inbox / Idempotency)

> Platform-shared provides wrappers; we keep service-local tables for reliability and recovery.

```sql
-- Outbox events to Kafka (exactly-once via event_id)
CREATE TABLE IF NOT EXISTS outbox_event (
  event_id        UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  event_type      text NOT NULL,
  version         int2 NOT NULL DEFAULT 1,
  occurred_at     timestamptz NOT NULL DEFAULT now(),
  partition_key   text NOT NULL,
  actor_id        uuid NULL,
  tenant_id       uuid NOT NULL,
  correlation_id  uuid NULL,
  causation_id    uuid NULL,
  schema_ref      text NOT NULL,
  nonpii_payload  jsonb NOT NULL,
  headers         jsonb NOT NULL DEFAULT '{}'::jsonb,
  published_at    timestamptz NULL,
  retry_count     int2 NOT NULL DEFAULT 0
);
CREATE INDEX IF NOT EXISTS idx_outbox_event_tenant_published ON outbox_event (tenant_id, published_at);

-- Inbox (dedupe consumer for external events)
CREATE TABLE IF NOT EXISTS inbox_event (
  event_id        uuid PRIMARY KEY,
  received_at     timestamptz NOT NULL DEFAULT now(),
  source          text NOT NULL,
  handler         text NOT NULL,
  status          text NOT NULL CHECK (status IN ('processed','skipped','failed')),
  attempts        int2 NOT NULL DEFAULT 1,
  last_error      text NULL
);

-- Idempotency keys for write API
CREATE TABLE IF NOT EXISTS idempotency_key (
  key             uuid PRIMARY KEY,
  created_at      timestamptz NOT NULL DEFAULT now(),
  expires_at      timestamptz NOT NULL,
  request_hash    text NOT NULL,
  response_hash   text NULL
);
CREATE INDEX IF NOT EXISTS idx_idem_expires ON idempotency_key (expires_at);
```

---

# SECTION 1 – CORE FILE PRIMITIVES

## 1.1 `file` (main) & `file_metadata` (sub)

```sql
CREATE TABLE IF NOT EXISTS file (
  id                uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id         uuid NOT NULL,
  namespace_id      uuid NOT NULL,                 -- FK namespace.id
  owner_id          uuid NOT NULL,                 -- user id (users-be)
  folder_id         uuid NULL,                     -- FK folder.id (nullable = root)
  name              text NOT NULL,                 -- original name without path
  ext               text NULL,                     -- derived extension (lowercased)
  mime              text NOT NULL,
  type              file_type_t NOT NULL DEFAULT 'generic',
  size_bytes        bigint NOT NULL CHECK (size_bytes >= 0),
  hash_sha256       text NOT NULL,                 -- content hash of current head version
  visibility        visibility_t NOT NULL DEFAULT 'private',
  status            file_status_t NOT NULL DEFAULT 'active',
  latest_version_id uuid NULL,                     -- FK version.id (head pointer)
  blob_id           uuid NULL,                     -- direct fast path (head blob) for reads
  policy_id         uuid NULL,                     -- FK policy.id applied at creation/update
  quarantine_id     uuid NULL,                     -- FK quarantine.id if quarantined
  deleted_at        timestamptz NULL,              -- soft delete marker & restore window enforced by lifecycle
  restore_by        timestamptz NULL,
  created_by        uuid NOT NULL,
  updated_by        uuid NULL,
  created_at        timestamptz NOT NULL DEFAULT now(),
  updated_at        timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT uq_file_hash_head UNIQUE (tenant_id, hash_sha256, status) DEFERRABLE INITIALLY DEFERRED
);
CREATE INDEX IF NOT EXISTS idx_file_tenant_folder ON file (tenant_id, folder_id, status);
CREATE INDEX IF NOT EXISTS idx_file_owner ON file (tenant_id, owner_id);
CREATE INDEX IF NOT EXISTS idx_file_name_trgm ON file USING gin (name gin_trgm_ops);
CREATE INDEX IF NOT EXISTS idx_file_mime ON file (mime);
CREATE INDEX IF NOT EXISTS idx_file_visibility ON file (tenant_id, visibility);

-- Derived/optional metadata (EXIF/dimensions/duration/pages etc.)
CREATE TABLE IF NOT EXISTS file_metadata (
  file_id           uuid PRIMARY KEY REFERENCES file(id) ON DELETE CASCADE,
  width_px          int4 NULL CHECK (width_px >= 0),
  height_px         int4 NULL CHECK (height_px >= 0),
  duration_ms       int8 NULL CHECK (duration_ms >= 0),
  page_count        int4 NULL CHECK (page_count >= 0),
  exif              jsonb NULL,
  custom            jsonb NULL,
  updated_at        timestamptz NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_file_metadata_custom ON file_metadata USING gin (custom);
```

### Integrity Hooks

```sql
-- auto-update updated_at
CREATE OR REPLACE FUNCTION touch_updated_at() RETURNS TRIGGER AS $$
BEGIN NEW.updated_at = now(); RETURN NEW; END; $$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_file_touch ON file;
CREATE TRIGGER trg_file_touch BEFORE UPDATE ON file
FOR EACH ROW EXECUTE FUNCTION touch_updated_at();
```

---

## 1.2 `folder` (main)

```sql
CREATE TABLE IF NOT EXISTS folder (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     uuid NOT NULL,
  namespace_id  uuid NOT NULL,
  owner_id      uuid NOT NULL,
  parent_id     uuid NULL REFERENCES folder(id) ON DELETE CASCADE,
  path          text NOT NULL,                   -- normalized absolute path string
  depth         int2 NOT NULL DEFAULT 0 CHECK (depth >= 0),
  name          text NOT NULL,
  created_at    timestamptz NOT NULL DEFAULT now(),
  updated_at    timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT uq_folder_path UNIQUE (tenant_id, namespace_id, path)
);
CREATE INDEX IF NOT EXISTS idx_folder_parent ON folder (tenant_id, parent_id);
CREATE INDEX IF NOT EXISTS idx_folder_name_trgm ON folder USING gin (name gin_trgm_ops);
```

---

## 1.3 `upload` (main) & `upload_chunk` (sub)

```sql
CREATE TABLE IF NOT EXISTS upload (
  id              uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id       uuid NOT NULL,
  namespace_id    uuid NOT NULL,
  owner_id        uuid NOT NULL,
  file_id         uuid NULL REFERENCES file(id) ON DELETE SET NULL,
  status          job_status_t NOT NULL DEFAULT 'queued',    -- started→running→succeeded|failed|canceled
  total_size      bigint NULL CHECK (total_size IS NULL OR total_size >= 0),
  received_bytes  bigint NOT NULL DEFAULT 0 CHECK (received_bytes >= 0),
  expected_parts  int4 NULL CHECK (expected_parts IS NULL OR expected_parts >= 0),
  etag            text NULL,                                  -- for resumable validation
  created_at      timestamptz NOT NULL DEFAULT now(),
  updated_at      timestamptz NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_upload_owner ON upload (tenant_id, owner_id, status);

CREATE TABLE IF NOT EXISTS upload_chunk (
  upload_id     uuid NOT NULL REFERENCES upload(id) ON DELETE CASCADE,
  part_no       int4 NOT NULL CHECK (part_no > 0),
  size_bytes    int8 NOT NULL CHECK (size_bytes >= 0),
  checksum_md5  text NOT NULL,
  received_at   timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY (upload_id, part_no)
);
```

---

## 1.4 `version` (main)

```sql
CREATE TABLE IF NOT EXISTS version (
  id              uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id       uuid NOT NULL,
  file_id         uuid NOT NULL REFERENCES file(id) ON DELETE CASCADE,
  blob_id         uuid NOT NULL,                     -- FK blob.id (below)
  number          int4 NOT NULL CHECK (number >= 1), -- 1..N
  author_id       uuid NOT NULL,
  promoted_at     timestamptz NOT NULL DEFAULT now(),
  created_at      timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT uq_version_no UNIQUE (file_id, number)
);
CREATE INDEX IF NOT EXISTS idx_version_file ON version (file_id);
```

---

# SECTION 2 – ACCESS & SHARING

## 2.1 `access_control` (main)

```sql
CREATE TABLE IF NOT EXISTS access_control (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     uuid NOT NULL,
  resource_type text NOT NULL CHECK (resource_type IN ('file','folder')),
  resource_id   uuid NOT NULL,                     -- file.id or folder.id
  subject_type  text NOT NULL CHECK (subject_type IN ('user','team','org','role')),
  subject_id    uuid NOT NULL,
  actions       text[] NOT NULL,                   -- e.g., {'view','download','edit'}
  granted_by    uuid NOT NULL,
  granted_at    timestamptz NOT NULL DEFAULT now(),
  revoked_at    timestamptz NULL,
  CONSTRAINT uq_acl UNIQUE (tenant_id, resource_type, resource_id, subject_type, subject_id)
);
CREATE INDEX IF NOT EXISTS idx_acl_resource ON access_control (tenant_id, resource_type, resource_id);
CREATE INDEX IF NOT EXISTS idx_acl_subject ON access_control (tenant_id, subject_type, subject_id);
```

---

## 2.2 `share` (main) & `share_link` (sub)

```sql
CREATE TABLE IF NOT EXISTS share (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     uuid NOT NULL,
  resource_type text NOT NULL CHECK (resource_type IN ('file','folder')),
  resource_id   uuid NOT NULL,
  grantee_type  text NOT NULL CHECK (grantee_type IN ('user','team','org','link')),
  grantee_id    uuid NULL,                         -- NULL when grantee_type = 'link'
  scopes        share_scope_t[] NOT NULL,          -- {'view','download','edit',...}
  expires_at    timestamptz NULL,
  created_by    uuid NOT NULL,
  created_at    timestamptz NOT NULL DEFAULT now(),
  revoked_at    timestamptz NULL
);
CREATE INDEX IF NOT EXISTS idx_share_resource ON share (tenant_id, resource_type, resource_id);
CREATE INDEX IF NOT EXISTS idx_share_expiry ON share (tenant_id, expires_at);

CREATE TABLE IF NOT EXISTS share_link (
  share_id      uuid PRIMARY KEY REFERENCES share(id) ON DELETE CASCADE,
  token_hash    text NOT NULL,                     -- store hash, not raw token
  usage_count   int4 NOT NULL DEFAULT 0 CHECK (usage_count >= 0),
  max_usage     int4 NULL CHECK (max_usage IS NULL OR max_usage > 0),
  last_used_at  timestamptz NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS uq_share_token ON share_link (token_hash);
```

---

## 2.3 `linking` (main) & `linking_audit` (sub)

```sql
CREATE TABLE IF NOT EXISTS linking (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     uuid NOT NULL,
  file_id       uuid NOT NULL REFERENCES file(id) ON DELETE CASCADE,
  action        link_action_t NOT NULL,            -- view | download | upload
  expires_at    timestamptz NOT NULL,
  signature     text NOT NULL,                     -- HMAC/signature ref
  revoked_at    timestamptz NULL,
  created_by    uuid NOT NULL,
  created_at    timestamptz NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_linking_file ON linking (tenant_id, file_id, action);
CREATE INDEX IF NOT EXISTS idx_linking_expiry ON linking (tenant_id, expires_at);

CREATE TABLE IF NOT EXISTS linking_audit (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     uuid NOT NULL,
  file_id       uuid NOT NULL REFERENCES file(id) ON DELETE CASCADE,
  linking_id    uuid NULL REFERENCES linking(id) ON DELETE SET NULL,
  user_id       uuid NULL,
  ip            inet NULL,
  user_agent    text NULL,
  action        link_action_t NOT NULL,
  occurred_at   timestamptz NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_linking_audit_file ON linking_audit (tenant_id, file_id, occurred_at DESC);
```

---

## 2.4 `lock` (main)

```sql
CREATE TABLE IF NOT EXISTS lock (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     uuid NOT NULL,
  file_id       uuid NOT NULL REFERENCES file(id) ON DELETE CASCADE,
  owner_id      uuid NOT NULL,
  lease_id      uuid NOT NULL,
  expires_at    timestamptz NOT NULL,
  created_at    timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT uq_lock_file UNIQUE (file_id)
);
CREATE INDEX IF NOT EXISTS idx_lock_expires ON lock (tenant_id, expires_at);
```

---

# SECTION 3 – CONTENT PIPELINE

## 3.1 `media` (main), `media_thumbnail`, `media_variant` (subs)

```sql
CREATE TABLE IF NOT EXISTS media (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     uuid NOT NULL,
  file_id       uuid NOT NULL REFERENCES file(id) ON DELETE CASCADE,
  kind          media_kind_t NOT NULL,
  pipeline      text NOT NULL,                     -- preset name
  status        job_status_t NOT NULL DEFAULT 'queued',
  attempts      int2 NOT NULL DEFAULT 0,
  error         text NULL,
  created_at    timestamptz NOT NULL DEFAULT now(),
  started_at    timestamptz NULL,
  finished_at   timestamptz NULL
);
CREATE INDEX IF NOT EXISTS idx_media_file ON media (tenant_id, file_id);
CREATE INDEX IF NOT EXISTS idx_media_status ON media (tenant_id, status);

CREATE TABLE IF NOT EXISTS media_thumbnail (
  media_id      uuid NOT NULL REFERENCES media(id) ON DELETE CASCADE,
  size_label    text NOT NULL,                     -- e.g., 'sm','md','lg'
  blob_id       uuid NOT NULL,                     -- FK blob.id
  created_at    timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY (media_id, size_label)
);

CREATE TABLE IF NOT EXISTS media_variant (
  media_id      uuid NOT NULL REFERENCES media(id) ON DELETE CASCADE,
  codec         text NOT NULL,
  bitrate_kbps  int4 NULL CHECK (bitrate_kbps IS NULL OR bitrate_kbps >= 0),
  width_px      int4 NULL CHECK (width_px IS NULL OR width_px >= 0),
  height_px     int4 NULL CHECK (height_px IS NULL OR height_px >= 0),
  blob_id       uuid NOT NULL,                     -- FK blob.id
  created_at    timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY (media_id, codec, width_px, height_px)
);
```

---

## 3.2 `extraction` (main) & `extraction_result` (sub)

```sql
CREATE TABLE IF NOT EXISTS extraction (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     uuid NOT NULL,
  file_id       uuid NOT NULL REFERENCES file(id) ON DELETE CASCADE,
  tool          text NOT NULL,                     -- 'tesseract','exiftool','pdftext',...
  status        job_status_t NOT NULL DEFAULT 'queued',
  attempts      int2 NOT NULL DEFAULT 0,
  error         text NULL,
  created_at    timestamptz NOT NULL DEFAULT now(),
  started_at    timestamptz NULL,
  finished_at   timestamptz NULL
);
CREATE INDEX IF NOT EXISTS idx_extraction_file ON extraction (tenant_id, file_id);

CREATE TABLE IF NOT EXISTS extraction_result (
  extraction_id uuid PRIMARY KEY REFERENCES extraction(id) ON DELETE CASCADE,
  text_content  text NULL,                         -- normalized extracted text
  metadata      jsonb NULL,                        -- normalized exif/attrs
  created_at    timestamptz NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_extraction_text_trgm ON extraction_result USING gin (text_content gin_trgm_ops);
CREATE INDEX IF NOT EXISTS idx_extraction_metadata ON extraction_result USING gin (metadata);
```

---

## 3.3 `artifact` (main)

```sql
CREATE TABLE IF NOT EXISTS artifact (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     uuid NOT NULL,
  type          text NOT NULL CHECK (type IN ('preview','zip','report','dataset')),
  file_id       uuid NULL REFERENCES file(id) ON DELETE CASCADE,
  blob_id       uuid NOT NULL,
  ttl_seconds   int4 NOT NULL CHECK (ttl_seconds > 0),
  expires_at    timestamptz NOT NULL,
  created_by    uuid NOT NULL,
  created_at    timestamptz NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_artifact_expiry ON artifact (tenant_id, expires_at);
```

---

# SECTION 4 – SAFETY & POLICY

## 4.1 `policy` (main), `policy_dlp_pattern`, `policy_result` (subs)

```sql
CREATE TABLE IF NOT EXISTS policy (
  id              uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id       uuid NOT NULL,
  name            text NOT NULL,
  max_size_mb     int4 NULL CHECK (max_size_mb IS NULL OR max_size_mb > 0),
  allow_mime      text[] NULL,
  block_mime      text[] NULL,
  virus_scan      boolean NOT NULL DEFAULT true,
  dlp_enabled     boolean NOT NULL DEFAULT false,
  data_zone       data_zone_t NOT NULL DEFAULT 'GLOBAL',
  created_at      timestamptz NOT NULL DEFAULT now(),
  updated_at      timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT uq_policy_name UNIQUE (tenant_id, name)
);
CREATE INDEX IF NOT EXISTS idx_policy_zone ON policy (tenant_id, data_zone);

CREATE TABLE IF NOT EXISTS policy_dlp_pattern (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  policy_id     uuid NOT NULL REFERENCES policy(id) ON DELETE CASCADE,
  name          text NOT NULL,
  pattern       text NOT NULL,                     -- regex or detector key
  created_at    timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS policy_result (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     uuid NOT NULL,
  file_id       uuid NOT NULL REFERENCES file(id) ON DELETE CASCADE,
  policy_id     uuid NOT NULL REFERENCES policy(id) ON DELETE CASCADE,
  violations    jsonb NOT NULL DEFAULT '[]'::jsonb,
  evaluated_at  timestamptz NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_policy_result_file ON policy_result (tenant_id, file_id, evaluated_at DESC);
```

---

## 4.2 `scan` (main) & `scan_result` (sub)

```sql
CREATE TABLE IF NOT EXISTS scan (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     uuid NOT NULL,
  file_id       uuid NOT NULL REFERENCES file(id) ON DELETE CASCADE,
  kind          text NOT NULL CHECK (kind IN ('av','dlp')),
  status        job_status_t NOT NULL DEFAULT 'queued',
  severity      severity_t NULL,
  created_at    timestamptz NOT NULL DEFAULT now(),
  started_at    timestamptz NULL,
  finished_at   timestamptz NULL,
  error         text NULL
);
CREATE INDEX IF NOT EXISTS idx_scan_file_kind ON scan (tenant_id, file_id, kind);

CREATE TABLE IF NOT EXISTS scan_result (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  scan_id       uuid NOT NULL REFERENCES scan(id) ON DELETE CASCADE,
  finding_code  text NOT NULL,                     -- e.g., 'MALWARE.EICAR' or 'PCI.PAN'
  match_hash    text NULL,
  byte_offset   int8 NULL CHECK (byte_offset IS NULL OR byte_offset >= 0),
  snippet_hash  text NULL,                         -- hashed window instead of raw content
  reason        text NULL
);
CREATE INDEX IF NOT EXISTS idx_scan_result_scan ON scan_result (scan_id);
```

---

## 4.3 `quarantine` (main)

```sql
CREATE TABLE IF NOT EXISTS quarantine (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     uuid NOT NULL,
  file_id       uuid NOT NULL UNIQUE REFERENCES file(id) ON DELETE CASCADE,
  reason        text NOT NULL,
  placed_by     uuid NOT NULL,
  placed_at     timestamptz NOT NULL DEFAULT now(),
  released_at   timestamptz NULL
);
CREATE INDEX IF NOT EXISTS idx_quarantine_tenant ON quarantine (tenant_id);
```

---

## 4.4 `file_flag` (main)

```sql
CREATE TABLE IF NOT EXISTS file_flag (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     uuid NOT NULL,
  file_id       uuid NOT NULL REFERENCES file(id) ON DELETE CASCADE,
  reporter_id   uuid NOT NULL,
  reason        text NOT NULL CHECK (reason IN ('malware','copyright','policy_violation','other')),
  state         text NOT NULL CHECK (state IN ('open','resolved','dismissed')),
  created_at    timestamptz NOT NULL DEFAULT now(),
  resolved_at   timestamptz NULL,
  resolved_by   uuid NULL
);
CREATE INDEX IF NOT EXISTS idx_flag_file_state ON file_flag (tenant_id, file_id, state);
```

---

# SECTION 5 – COMPLIANCE & LIFECYCLE

## 5.1 `lifecycle` (main) & `lifecycle_legal_hold` (sub)

```sql
CREATE TABLE IF NOT EXISTS lifecycle (
  id                  uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id           uuid NOT NULL,
  name                text NOT NULL,
  rule_state          text NOT NULL CHECK (rule_state IN ('active','soft_deleted','quarantined','archived')),
  retention_days      int4 NULL CHECK (retention_days IS NULL OR retention_days >= 0),
  restore_window_days int4 NULL CHECK (restore_window_days IS NULL OR restore_window_days >= 0),
  legal_hold          boolean NOT NULL DEFAULT false,
  created_at          timestamptz NOT NULL DEFAULT now(),
  updated_at          timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT uq_lifecycle_name UNIQUE (tenant_id, name)
);

CREATE TABLE IF NOT EXISTS lifecycle_legal_hold (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     uuid NOT NULL,
  file_id       uuid NOT NULL REFERENCES file(id) ON DELETE CASCADE,
  state         legal_hold_state_t NOT NULL DEFAULT 'active',
  reason        text NULL,
  placed_by     uuid NOT NULL,
  placed_at     timestamptz NOT NULL DEFAULT now(),
  released_at   timestamptz NULL
);
CREATE INDEX IF NOT EXISTS idx_legal_hold_file ON lifecycle_legal_hold (tenant_id, file_id, state);
```

---

## 5.2 `audit` (main)  — **Monthly Partitioned**

```sql
CREATE TABLE IF NOT EXISTS audit (
  id            uuid NOT NULL,
  tenant_id     uuid NOT NULL,
  resource_type text NOT NULL CHECK (resource_type IN ('file','folder','policy','share','link','blob','namespace')),
  resource_id   uuid NOT NULL,
  action        text NOT NULL,
  actor_id      uuid NULL,
  ip            inet NULL,
  user_agent    text NULL,
  details       jsonb NULL,
  occurred_at   timestamptz NOT NULL,
  PRIMARY KEY (id, occurred_at)  -- for partitioning
) PARTITION BY RANGE (occurred_at);

-- rolling 24 months by default (adjust in ops)
CREATE TABLE IF NOT EXISTS audit_2025_10 PARTITION OF audit
  FOR VALUES FROM ('2025-10-01') TO ('2025-11-01');

-- Create future partitions via ops scripts / worker
CREATE INDEX IF NOT EXISTS idx_audit_tenant_resource ON audit_2025_10 (tenant_id, resource_type, resource_id, occurred_at DESC);
CREATE INDEX IF NOT EXISTS idx_audit_actor ON audit_2025_10 (tenant_id, actor_id, occurred_at DESC);
```

---

# SECTION 6 – STORAGE PRIMITIVES & OPS

## 6.1 `blob` (main) — **Hash-Partitioned by tenant**

```sql
CREATE TABLE IF NOT EXISTS blob (
  id              uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id       uuid NOT NULL,
  sha256          text NOT NULL,
  size_bytes      bigint NOT NULL CHECK (size_bytes >= 0),
  storage_class   storage_class_t NOT NULL DEFAULT 'standard',
  ref_count       int8 NOT NULL DEFAULT 0 CHECK (ref_count >= 0),
  location        text NOT NULL,                 -- bucket/key prefix or provider ref
  created_at      timestamptz NOT NULL DEFAULT now(),
  UNIQUE (tenant_id, sha256)
) PARTITION BY HASH (tenant_id);

-- example partitions (tune by scale)
CREATE TABLE IF NOT EXISTS blob_p0 PARTITION OF blob FOR VALUES WITH (MODULUS 8, REMAINDER 0);
CREATE TABLE IF NOT EXISTS blob_p1 PARTITION OF blob FOR VALUES WITH (MODULUS 8, REMAINDER 1);
CREATE TABLE IF NOT EXISTS blob_p2 PARTITION OF blob FOR VALUES WITH (MODULUS 8, REMAINDER 2);
CREATE TABLE IF NOT EXISTS blob_p3 PARTITION OF blob FOR VALUES WITH (MODULUS 8, REMAINDER 3);
CREATE TABLE IF NOT EXISTS blob_p4 PARTITION OF blob FOR VALUES WITH (MODULUS 8, REMAINDER 4);
CREATE TABLE IF NOT EXISTS blob_p5 PARTITION OF blob FOR VALUES WITH (MODULUS 8, REMAINDER 5);
CREATE TABLE IF NOT EXISTS blob_p6 PARTITION OF blob FOR VALUES WITH (MODULUS 8, REMAINDER 6);
CREATE TABLE IF NOT EXISTS blob_p7 PARTITION OF blob FOR VALUES WITH (MODULUS 8, REMAINDER 7);

CREATE INDEX IF NOT EXISTS idx_blob_sha ON blob (tenant_id, sha256);
```

### Blob Refcount Hooks

```sql
-- increase ref_count when a version links a blob
CREATE OR REPLACE FUNCTION blob_ref_inc() RETURNS TRIGGER AS $$
BEGIN
  UPDATE blob SET ref_count = ref_count + 1 WHERE id = NEW.blob_id;
  RETURN NEW;
END $$ LANGUAGE plpgsql;

-- decrease ref_count when version row is deleted
CREATE OR REPLACE FUNCTION blob_ref_dec() RETURNS TRIGGER AS $$
BEGIN
  UPDATE blob SET ref_count = GREATEST(0, ref_count - 1) WHERE id = OLD.blob_id;
  RETURN OLD;
END $$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_version_blob_inc ON version;
CREATE TRIGGER trg_version_blob_inc AFTER INSERT ON version
FOR EACH ROW EXECUTE FUNCTION blob_ref_inc();

DROP TRIGGER IF EXISTS trg_version_blob_dec ON version;
CREATE TRIGGER trg_version_blob_dec AFTER DELETE ON version
FOR EACH ROW EXECUTE FUNCTION blob_ref_dec();
```

---

## 6.2 `reference` (main)

```sql
CREATE TABLE IF NOT EXISTS reference (
  id              uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id       uuid NOT NULL,
  aggregate_type  text NOT NULL,               -- e.g., 'job','proposal','contract','message','ticket'
  aggregate_id    uuid NOT NULL,
  file_id         uuid NOT NULL REFERENCES file(id) ON DELETE CASCADE,
  purpose         text NOT NULL,               -- 'attachment','avatar','deliverable','evidence',...
  created_at      timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT uq_reference UNIQUE (tenant_id, aggregate_type, aggregate_id, file_id, purpose)
);
CREATE INDEX IF NOT EXISTS idx_reference_agg ON reference (tenant_id, aggregate_type, aggregate_id);
CREATE INDEX IF NOT EXISTS idx_reference_file ON reference (tenant_id, file_id);
```

---

## 6.3 `quota` (main)

```sql
CREATE TABLE IF NOT EXISTS quota (
  id              uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id       uuid NOT NULL,
  subject_type    text NOT NULL CHECK (subject_type IN ('user','team','org','tenant')),
  subject_id      uuid NULL,                    -- NULL when tenant-wide
  hard_bytes      bigint NOT NULL CHECK (hard_bytes >= 0),
  soft_bytes      bigint NULL CHECK (soft_bytes IS NULL OR soft_bytes >= 0),
  used_bytes      bigint NOT NULL DEFAULT 0 CHECK (used_bytes >= 0),
  window          text NULL,                    -- e.g., 'monthly','rolling_30'
  updated_at      timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT uq_quota UNIQUE (tenant_id, subject_type, subject_id)
);
CREATE INDEX IF NOT EXISTS idx_quota_subject ON quota (tenant_id, subject_type, subject_id);
```

---

## 6.4 `namespace` (main)

```sql
CREATE TABLE IF NOT EXISTS namespace (
  id                uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id         uuid NOT NULL,
  bucket            text NOT NULL,
  data_zone         data_zone_t NOT NULL DEFAULT 'GLOBAL',
  encryption_policy jsonb NULL,                 -- KMS/key policy refs
  created_at        timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT uq_namespace UNIQUE (tenant_id, bucket, data_zone)
);
CREATE INDEX IF NOT EXISTS idx_namespace_tenant ON namespace (tenant_id, data_zone);
```

---

## 6.5 `gc` (main)

```sql
CREATE TABLE IF NOT EXISTS gc (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     uuid NOT NULL,
  scope         text NOT NULL CHECK (scope IN ('blob','artifact','soft_deleted_files')),
  state         text NOT NULL CHECK (state IN ('planned','running','completed','failed')),
  stats         jsonb NOT NULL DEFAULT '{}'::jsonb,
  started_at    timestamptz NULL,
  finished_at   timestamptz NULL,
  created_at    timestamptz NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_gc_state ON gc (tenant_id, state, created_at DESC);
```

---

## 6.6 `encryption` (main) & `encryption_rotation` (sub)

```sql
CREATE TABLE IF NOT EXISTS encryption (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     uuid NOT NULL,
  file_id       uuid NOT NULL UNIQUE REFERENCES file(id) ON DELETE CASCADE,
  key_id        text NOT NULL,
  version       int4 NOT NULL DEFAULT 1,
  rotated_at    timestamptz NULL,
  created_at    timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS encryption_rotation (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     uuid NOT NULL,
  file_id       uuid NOT NULL REFERENCES file(id) ON DELETE CASCADE,
  from_key      text NOT NULL,
  to_key        text NOT NULL,
  rotated_by    uuid NOT NULL,
  rotated_at    timestamptz NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_encryption_rotation_file ON encryption_rotation (tenant_id, file_id, rotated_at DESC);
```

---

# SECTION 7 – RELATIONSHIPS & FK MAP (Cross-Service)

- **users-be**: `file.owner_id`, `created_by`, `updated_by`, `reporter_id`, `placed_by`, etc. reference **users** (logical FK).  
- **contracts-be / jobs-be / proposals-be / communications-be**: linked through `reference(aggregate_type, aggregate_id)` for attachments, deliverables, message files.  
- **financial-be**: artifacts/reports (e.g., invoice PDFs) linked via `reference`.  
- **admin / policy**: `policy` and `policy_result` tie to admin policy actions; consumed events update caches.  
- **data residency**: `namespace.data_zone` governs route of blobs and signed URLs.

> All external references are **logical** (not hard FKs) to avoid cross-service coupling. Validate at application layer and via asynchronous compensations.

---

# SECTION 8 – PERFORMANCE: Indexing, Views, and Caching

```sql
-- Read-optimized view for listing files in folders
CREATE OR REPLACE VIEW file_read AS
SELECT  f.id, f.tenant_id, f.namespace_id, f.owner_id, f.folder_id, f.name, f.ext,
        f.mime, f.type, f.size_bytes, f.visibility, f.status, f.latest_version_id,
        fm.width_px, fm.height_px, fm.page_count, fm.duration_ms,
        f.created_at, f.updated_at
FROM file f
LEFT JOIN file_metadata fm ON fm.file_id = f.id;

-- Search helpers (trgm already added on names and text content)
```

- **Redis** caches: hot `file` metadata, policy decisions, short‑TTL lock leases, signed URL pre-checks.  
- **GIN/Trigram**: names, extraction text/metadata for in-app search.  
- **Partitioning**: `blob` (hash by tenant), `audit` (monthly), optionally `linking_audit` by month at extreme scale.

---

# SECTION 9 – SECURITY & COMPLIANCE

- **PII-safe**: no raw DLP matches or file contents; only hashes and metadata.  
- **Soft delete**: `file.deleted_at` + `restore_by` enforced by `lifecycle`; hard-deletes via GC tasks.  
- **Quarantine** isolates risky files; shared links auto-revoked on quarantine.  
- **Data residency**: enforced via `namespace.data_zone` and policy checks.  
- **Signatures**: `linking.signature` (HMAC/Ed25519) references key material outside DB.  
- **Auditing**: all critical actions append to `audit` (partitioned).

---

# SECTION 10 – SAMPLE GRANTS & (Optional) RLS

> RLS is optional; many teams implement at service layer. If desired, enable per table with tenant scoping.

```sql
-- Example: enable RLS for file
ALTER TABLE file ENABLE ROW LEVEL SECURITY;

-- Each tenant sees only its rows
CREATE POLICY file_tenant_isolation ON file
  USING (tenant_id = current_setting('app.tenant_id')::uuid);

-- Repeat for other main tables as needed.
```

---

# SECTION 11 – MIGRATIONS & SEEDING ORDER

1. **Enums** → Global tables (outbox/inbox/idempotency).  
2. **namespace** → **policy** → **folder** → **file** → **file_metadata**.  
3. **blob** → **version** (triggers for ref_count).  
4. **access_control / share / linking (+ audit)**.  
5. **media / extraction / artifact**.  
6. **scan / scan_result / quarantine / file_flag**.  
7. **lifecycle / lifecycle_legal_hold / audit partitions**.  
8. **reference / quota / gc / encryption / encryption_rotation**.  

---

# SECTION 12 – TOPIC MAP (Publish/Consume)

- **Publish**: `file.*`, `version.*`, `blob.*`, `media.*`, `scan.*`, `artifact.*`, `policy.*`, `lifecycle.*`, `linking.*`, `share.*`, `quota.*`  
- **Consume**: `user.*`, `admin.policy.*`, `admin.moderation.*`, `contract.*`

---

## Notes

- Table names match your domain directories exactly; sub-entities follow `{domain}_{sub}`.  
- All constraints favor **safety** (check constraints), **scale** (indexes, partitions), and **observability** (audit, outbox).  
- Adjust partition counts and retention at 10B+ rows scale.
