
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
