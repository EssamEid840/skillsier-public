# storage-be User Stories (Consolidated & Aligned)

**Complete Extended User Stories**  
Following domain-driven design with event-sourcing, CQRS projections, and platform alignment.  
This version keeps the **same style and sections** as the original document and **adds** the missing stories and event catalog items we identified.

---

## Global Conventions (applies to all sections)

**Event envelope (Kafka / outbox)**
```
event_id (ULID), event_type, version, occurred_at (UTC), actor_id, tenant_id,
correlation_id, causation_id, partition_key, schema_ref (hash),
nonpii_payload (domain DTO; no raw files or secrets)
```

**Idempotent write-path**
- All write handlers accept Idempotency-Key (UUIDv4). Server returns the original success payload on safe retries (TTL 24h).
- Natural keys prevent duplicates (e.g., {user_id, file_hash} for blobs; {file_id} for metadata).
- Outbox is exactly-once per event_id. Inbox performs dedupe on event_id + causation_id.

**Non-PII rules**
- Never store raw emails/phones/IM handles inside metadata; lint + redact at write.
- Files are content-addressed by hash; all metadata undergoes PII detection asynchronously.
- Audit logs scrub sensitive text; indices keep only hashed/fingerprinted fields where needed.

**Platform alignment**
- Follows provided folder structure (domain/application/interfaces layers).
- Topics grouped by feature (e.g., file.*, blob.*, media.*, scan.*, lifecycle.*).
- Projections are suffixed _read; commands/queries in application layer.

---

# 1 - CORE FILE MANAGEMENT

## 1.1 file/

### Stories
- As a **user**, I want to upload files with metadata (name, type, size, tags) so that I can store and organize my content.
- As a **user**, I want upload guardrails (max file size, allowed MIME types, quota checks) so that I don't exceed limits or upload malicious content.
- As a **system**, I want content-addressed storage with de-duplication so that identical files share the same blob and save storage space.
- As a **system**, I want dedupe within upload session for {user_id, file_hash} so that duplicate uploads are prevented.
- As a **user**, I want to retrieve file metadata and generate signed download URLs so that I can access my files securely.
- As a **user**, I want to soft-delete files with a restore window so that I can recover accidentally deleted files.
- **As a user, I want to rename files, edit metadata/tags, and change visibility so that I can manage how files appear and are shared.**  *(new)*
- **As a user, I want to copy/duplicate a file (same blob, new file row) so that I can reuse content without extra storage.**  *(new)*
- **As an admin/system, I want to transfer file ownership across users/tenants so that we can reassign content when accounts move.**  *(new)*
*   As a **user**, I want to **search my files** by name, tags, MIME, size, date, and extracted text (OCR) with sort & pagination so I can find things fast.

### Flow
- UploadFileCommand(user_id, file_stream, file_name, mime_type, size, tags?, folder_id?) → ValidateQuota() | ValidateMIMEType() | ValidateSize() | ComputeHash() | CheckDuplicate() → If exists: LinkToExistingBlob() ELSE: StoreBlobToMinIO() | CreateFileMetadata() → EnqueueVirusScan(file_id) | EnqueuePolicyCheck(file_id) → **Outbox:** file.uploaded.v1
- GetFileCommand(file_id) → RetrieveMetadata() | CheckPermissions() | GenerateSignedURL(ttl=1h) → Return file metadata + signed URL
- SoftDeleteFileCommand(file_id, user_id) → GuardOwnership() | MarkDeleted(deleted_at=now, restore_by=now+30d) | UnlinkBlob() → **Outbox:** file.soft_deleted.v1
- RestoreFileCommand(file_id, user_id) → GuardOwnership() | ValidateRestoreWindow() | RestoreFile() | RelinkBlob() → **Outbox:** file.restored.v1
- PermanentDeleteFileCommand(file_id) → GuardOwnership() | DeleteFileMetadata() | DecrementBlobRefCount() | EnqueueGC(blob_id) → **Outbox:** file.deleted.v1
- **RenameFileCommand(file_id, new_name, user_id) → GuardOwnership() | AntiPiiLint(new_name) | UpdateName() → Outbox: file.metadata.updated.v1**  *(new)*
- **UpdateFileMetadataCommand(file_id, metadata, user_id) → GuardOwnership() | ValidateKeys() | UpdateMetadata() → Outbox: file.metadata.updated.v1**  *(new)*
- **UpdateFileVisibilityCommand(file_id, visibility, user_id) → GuardOwnership() | ValidateVisibility() | UpdateVisibility() → Outbox: file.metadata.updated.v1**  *(new)*
- **CopyFileCommand(file_id, target_folder_id?, user_id) → GuardAccess() | LinkExistingBlob() | CreateNewFileRow() → Outbox: file.copied.v1**  *(new)*
- **TransferFileOwnershipCommand(file_id, to_owner_id, admin_id) → GuardAdmin() | ValidateReferences() | UpdateOwner() | UpdateNamespace() → Outbox: file.ownership.transferred.v1**  *(new)*
- SearchFilesQuery(user\_id, filters{name?, tags?, mime?, size\_range?, date\_range?, full\_text?}, sort?, page, limit) → GuardAccess() | Use index (trgm/GIN) + OCR text projection | Return paginated results.

### Projections
- files_read (per file: id, name, type, size, hash, owner, folder, tags, upload_date, scan_status, policy_status, visibility)
- file_metadata_cache (Redis: file_id → metadata, TTL 1h)
- user_files_read (user_id → file list with filters by folder/type/date)
- file\_search\_index\_read (file\_id → name, tags, mime, size, dates, OCR\_text\_ref, rank)

### Events
- file.uploaded.v1, file.soft_deleted.v1, file.restored.v1, file.deleted.v1, file.metadata.updated.v1, **file.copied.v1**, **file.ownership.transferred.v1**

### RBAC/SLO
- **RBAC:** OWNER for CRUD; ADMIN for force delete/ownership transfer
- **SLO:** upload P95 < 3s for files <50MB; metadata retrieval P95 < 150ms; signed URL generation P95 < 100ms

### Limits
- max 10GB per user (default); max 100MB per file; max 1000 files per user (plan overrides allowed)

### Idempotency
- by (user_id, file_hash) for uploads; by file_id for deletes/updates

---

## 1.2 folder/

### Stories
- As a **user**, I want to create folders with hierarchical structure so that I can organize my files logically.
- As a **user**, I want to move files between folders so that I can reorganize my content.
- As a **user**, I want to rename and delete folders (with contents) so that I can manage my file structure.
- As a **system**, I want to validate folder depth limits and prevent circular hierarchies so that the structure remains manageable.
- As a **user**, I want to list folder contents with pagination so that I can browse large collections efficiently.
- **As a user, I want to search folders by name/path so that I can quickly navigate large hierarchies.**  *(new)*
- **As a user, I want to copy/duplicate a folder (with optional deep copy of contents) so that I can clone structures.**  *(new)*

### Flow
- CreateFolderCommand(user_id, folder_name, parent_folder_id?) → ValidateDepth(max=10) | ValidateCircular() | AntiPiiLint(folder_name) → CreateFolder() → **Outbox:** folder.created.v1
- RenameFolderCommand(folder_id, new_name, user_id) → GuardOwnership() | AntiPiiLint(new_name) | RenameFolder() → **Outbox:** folder.renamed.v1
- MoveFolderCommand(folder_id, new_parent_id, user_id) → GuardOwnership() | ValidateDepth() | ValidateCircular() | MoveFolder() → **Outbox:** folder.moved.v1
- DeleteFolderCommand(folder_id, user_id, recursive=false) → GuardOwnership() → If recursive: SoftDeleteAllFiles() | DeleteAllSubfolders() ELSE: CheckEmpty() → DeleteFolder() → **Outbox:** folder.deleted.v1
- MoveFileToFolderCommand(file_id, folder_id, user_id) → GuardOwnership() | MoveFile() → **Outbox:** file.moved.v1
- **SearchFoldersQuery(query, user_id, page, limit) → GuardAccess() | Search(materialized_path, name) → return folders**  *(new)*
- **CopyFolderCommand(folder_id, deep:bool, user_id) → GuardOwnership() | CloneTree(deep?) | LinkBlobsForDeepCopy() → Outbox: folder.copied.v1**  *(new)*

### Projections
- folders_read (per folder: id, name, parent_id, owner, created_at, file_count, subfolder_count)
- folder_hierarchy_read (materialized path for quick ancestor/descendant queries)
- folder_cache (Redis: folder_id → metadata, TTL 30m)

### Events
- folder.created.v1, folder.renamed.v1, folder.moved.v1, folder.deleted.v1, **folder.copied.v1**

### RBAC/SLO
- **RBAC:** OWNER; **VIEWER** for list/search
- **SLO:** create/rename/move P95 < 200ms; list contents P95 < 150ms; delete (recursive) P95 < 5s

### Limits
- max depth 10 levels; max 10,000 folders per user; max 1000 files per folder

### Idempotency
- by (user_id, folder_name, parent_id) for create

---

## 1.3 upload/

### Stories
- As a **user**, I want resumable/chunked uploads for large files so that I can recover from network interruptions.
- As a **system**, I want to track upload sessions with expiry so that incomplete uploads don't linger indefinitely.
- As a **user**, I want to pause and resume uploads so that I can manage bandwidth usage.
- As a **system**, I want to validate chunk integrity and order so that uploaded files are complete and correct.
- As a **user**, I want upload progress tracking so that I know how much has been uploaded.
- **As a user, I want an explicit resume operation and server-advertised missing-chunk recovery so that resuming is reliable.**  *(new)*
- As a **system**, I **expire stale upload sessions** on schedule and clean partial chunks.

### Flow
- InitiateUploadCommand(user_id, file_name, total_size, chunk_count) → ValidateQuota() | CreateUploadSession(upload_id, expires_at=now+24h) → **Outbox:** upload.initiated.v1
- UploadChunkCommand(upload_id, chunk_index, chunk_data, chunk_hash) → ValidateSession() | ValidateChunkHash() | StoreChunk(temp_storage) | UpdateProgress() → **Outbox:** upload.chunk.uploaded.v1
- **ResumeUploadCommand(upload_id) → ValidateSession() | ListReceivedChunks() | ReturnMissingRanges() → Outbox: upload.resumed.v1**  *(new)*
- CompleteUploadCommand(upload_id) → ValidateAllChunks() | AssembleFile() | ComputeFinalHash() | CreateFileMetadata() | StoreBlobToMinIO() | DeleteTempChunks() → **Outbox:** upload.completed.v1 + file.uploaded.v1
- AbortUploadCommand(upload_id, user_id) → GuardOwnership() | DeleteTempChunks() | DeleteSession() → **Outbox:** upload.aborted.v1
- ExpireUploadSessionsJob() → Find sessions expires\_at < now && status IN (PENDING, IN\_PROGRESS) | Delete temp chunks | Close session → **Outbox:** upload.expired.v1

### Projections
- upload_sessions_read (upload_id → status, progress, chunks_received, expires_at)
- active_uploads_cache (Redis: upload_id → session metadata, TTL 24h)

### Events
- upload.initiated.v1, upload.chunk.uploaded.v1, upload.progress.updated.v1, upload.completed.v1, upload.aborted.v1, upload.expired.v1, **upload.resumed.v1**

### RBAC/SLO
- **RBAC:** OWNER
- **SLO:** initiate P95 < 200ms; chunk upload P95 < 1s; complete P95 < 3s; abort P95 < 300ms

### Limits
- max 5GB per upload; max 100 concurrent uploads per user; session TTL 24h; chunk size 5MB

### Idempotency
- by upload_id for all operations

---

## 1.4 version/

### Stories
- As a **user**, I want file versioning so that I can track changes and revert to previous versions.
- As a **user**, I want to list all versions of a file so that I can see the history.
- As a **user**, I want to restore a previous version so that I can recover from unwanted changes.
- As a **system**, I want version limits and retention policies so that storage doesn't grow unbounded.
- As a **user**, I want to delete specific versions (except current) so that I can clean up old versions.
- **As a user, I want to promote/alias a historical version as current without rewriting blobs so that I can switch quickly.**  *(new)*

### Flow
- CreateVersionCommand(file_id, user_id, change_description?) → GuardOwnership() | CreateNewVersion(blob_id, version_number++) | UpdateFileMetadata(current_version) → **Outbox:** file.version.created.v1
- ListVersionsQuery(file_id, user_id) → GuardAccess() | RetrieveVersionHistory() → Return version list
- RestoreVersionCommand(file_id, version_number, user_id) → GuardOwnership() | CreateNewVersionFromOld() | UpdateFileMetadata(current_version) → **Outbox:** file.version.restored.v1
- DeleteVersionCommand(file_id, version_number, user_id) → GuardOwnership() | CheckNotCurrent() | DeleteVersion() | DecrementBlobRefCount() → **Outbox:** file.version.deleted.v1
- SetVersionLimitCommand(user_id, max_versions) → ValidateLimit(1-100) | UpdateUserSettings() → **Outbox:** version.limit.updated.v1
- **PromoteVersionCommand(file_id, version_number, user_id) → GuardOwnership() | MarkAsCurrent() → Outbox: file.version.promoted.v1**  *(new)*

### Projections
- file_versions_read (file_id → version list with blob_id, size, created_at, change_description)
- version_history_cache (Redis: file_id → version list, TTL 1h)

### Events
- file.version.created.v1, file.version.restored.v1, file.version.deleted.v1, version.limit.updated.v1, **file.version.promoted.v1**

### RBAC/SLO
- **RBAC:** **OWNER**; **VIEWER** for list
- **SLO:** create version P95 < 500ms; list versions P95 < 150ms; restore P95 < 1s; delete P95 < 300ms

### Limits
- default 10 versions per file; max 100 versions; auto-cleanup versions older than 90 days

### Idempotency
- by (file_id, version_number)

---

# 2 - ACCESS CONTROL & SHARING

## 2.1 access_control/

### Stories
- As a **user**, I want to set granular permissions (read, write, delete) on files and folders so that I can control access.
- As a **user**, I want permission inheritance from folders to files so that access control is manageable.
- As a **user**, I want to revoke access so that I can remove permissions when needed.
- As a **system**, I want permission checks cached so that access validation is fast.
- **As an admin, I want to assign permissions to orgs/teams/groups (not only users) so that access scales.**  *(new)*
- **As a user/admin, I want to list effective permissions for a user on a resource so that I can troubleshoot access.**  *(new)*

### Flow
- GrantAccessCommand(file_id, subject(user|team|org)_id, permission_level, granted_by) → GuardOwnership() | CreateACLEntry(subject, permission) → InvalidateCache() → **Outbox:** access.granted.v1
- RevokeAccessCommand(file_id, subject_id, revoked_by) → GuardOwnership() | DeleteACLEntry() → InvalidateCache() → **Outbox:** access.revoked.v1
- UpdatePermissionCommand(file_id, subject_id, new_permission, updated_by) → GuardOwnership() | UpdateACL() → InvalidateCache() → **Outbox:** access.updated.v1
- CheckAccessQuery(file_id, user_id, required_permission) → CheckOwnership() OR CheckACL() OR CheckInheritedPermission() → Return boolean + reason
- ListAccessQuery(file_id) → GuardOwnership() | RetrieveACL() → Return permission list
- **ListEffectivePermissionsQuery(user_id, resource_id) → ResolveGroups+Inheritance() → Return effective scopes**  *(new)*

### Projections
- file_acl_read (file_id → user/team/org permissions list)
- folder_acl_read (folder_id → inherited permissions)
- access_cache (Redis: {file_id, user_id} → permission, TTL 15m)

### Events
- access.granted.v1, access.revoked.v1, access.updated.v1

### RBAC/SLO
- **RBAC:** **OWNER** for grant/revoke/update; **SYSTEM** for check
- **SLO:** grant/revoke P95 < 200ms; check P95 < 10ms (cached); list P95 < 150ms

### Limits
- max 100 ACL entries per file; max 1000 ACL entries per user

### Idempotency
- by (file_id, subject_id) for grant/revoke

---

## 2.2 share/

### Stories
- As a **user**, I want to create signed share links with expiry so that I can share files with anyone securely.
- As a **user**, I want password-protected shares so that only intended recipients can access.
- As a **user**, I want download limits on shares so that I can control usage.
- As a **system**, I want to track share access (views, downloads) so that users can see engagement.
- As a **user**, I want to revoke shares so that I can stop access immediately.
- **As an admin, I want to transfer share ownership so that shares can be re-assigned.**  *(new)*
- **As an admin, I want to bulk revoke shares by filters so that I can respond to incidents.**  *(new)*

### Flow
- CreateShareLinkCommand(file_id, user_id, expiry_hours, password?, max_downloads?, watermark?) → GuardOwnership() | GenerateToken() | HashPassword() | CreateShare(expires_at, settings) → **Outbox:** share.link.created.v1
- AccessShareLinkCommand(share_token, password?) → ValidateToken() | ValidateExpiry() | ValidatePassword() | CheckDownloadLimit() | LogAccess() | IncrementAccessCount() → GenerateSignedURL(ttl=1h) → **Outbox:** share.link.accessed.v1
- RevokeShareLinkCommand(share_id, user_id) → GuardOwnership() | MarkRevoked() → InvalidateToken() → **Outbox:** share.link.revoked.v1
- ListSharesQuery(file_id, user_id) → GuardOwnership() | RetrieveShares() → Return share list with analytics
- UpdateShareSettingsCommand(share_id, settings, user_id) → GuardOwnership() | UpdateShare() → **Outbox:** share.link.updated.v1
- **TransferShareOwnershipCommand(share_id, to_user_id, admin_id) → GuardAdmin() | UpdateOwner() → Outbox: share.ownership.transferred.v1**  *(new)*
- **BulkRevokeSharesCommand(filter, admin_id) → GuardAdmin() | RevokeAll() → Outbox: share.bulk_revoked.v1**  *(new)*

### Projections
- shares_read (share_id → file_id, owner, token, expires_at, access_count, download_count, settings)
- share_analytics_read (share_id → access logs, unique viewers, geo data)
- share_token_cache (Redis: token → share_id + metadata, TTL = expiry)

### Events
- share.link.created.v1, share.link.accessed.v1, share.link.revoked.v1, share.link.updated.v1, share.link.expired.v1, **share.ownership.transferred.v1**, **share.bulk_revoked.v1**

### RBAC/SLO
- **RBAC:** **OWNER** for create/revoke/update; **PUBLIC** for access (with token)
- **SLO:** create P95 < 200ms; access P95 < 150ms; revoke P95 < 100ms; list P95 < 150ms

### Limits
- max expiry 365 days; default expiry 7 days; max 1000 downloads per share

### Idempotency
- by (file_id, user_id, settings_hash) for create

---

## 2.3 linking/

### Stories
- As a **user**, I want to generate time-limited signed URLs for direct downloads so that I can share files without permanent links.
- As a **system**, I want signed URL validation with signature verification so that tampering is prevented.
- As a **system**, I want to track signed URL usage for audit purposes so that downloads are traceable.
- As a **user**, I want to generate batch signed URLs for multiple files so that I can create archives efficiently.
- **As an admin/compliance officer, I want to list and export signed-URL audit logs so that I can investigate usage.**  *(new)*
- As an **admin**, I want to **bulk revoke signed URLs** by filters (resource, creator, date range) to respond to incidents quickly.
- As a **user**, I can **resume downloads** via HTTP Range requests, benefit from ETag/If-None-Match caching, and get proper Content-Disposition filenames.
- As a **user**, I can **pause/resume** downloads (HTTP Range), leverage **ETag/If-None-Match**, and get correct Content-Disposition names.

### Flow
- GenerateSignedURLCommand(file_id, user_id, ttl_seconds=3600, purpose?) → GuardAccess() | RetrieveBlobLocation() | GenerateSignature(expires_at, file_id, secret) | ConstructURL() | LogGeneration() → **Outbox:** signed_url.generated.v1
- ValidateSignedURLCommand(url, signature) → ParseURL() | ExtractClaims() | ValidateExpiry() | ValidateSignature() | LogValidation() → Return validation result
- GenerateBatchSignedURLsCommand(file_ids[], user_id, ttl_seconds) → ForEach: GuardAccess() | GenerateSignedURL() → Return URL list → **Outbox:** signed_url.batch.generated.v1
- RevokeSignedURLCommand(url_id, user_id) → GuardOwnership() | InvalidateSignature() → **Outbox:** signed_url.revoked.v1
- **ListSignedURLAuditQuery(resource_id?, date_range?, admin_id) → GuardAdmin() | RetrieveAudit()**  *(new)*
- **ExportSignedURLAuditCommand(filters, format, admin_id) → GuardAdmin() | RetrieveAudit() | CreateExportArtifact() → Outbox: signed_url.audit.exported.v1**  *(new)*
- BulkRevokeSignedURLsCommand(filters{resource\_id?, created\_by?, date\_range?}, admin\_id) → GuardAdmin() | Invalidate signatures | **Outbox:** signed\_url.bulk\_revoked.v1
- (Download semantics) GenerateSignedURLCommand(...) → also set ETag=file\_hash | Content-Disposition from current file name; handlers support Range/206.
- Handler emits 206 for ranges; sets ETag=file\_hash, Last-Modified, and Content-Disposition="attachment; filename=". No new domain event.

### Projections
- signed_url_audit_read (url_id → file_id, generated_by, generated_at, expires_at, access_count)
- signed_url_cache (Redis: signature → {file_id, expires_at}, TTL = expiry)

### Events
- signed_url.generated.v1, signed_url.batch.generated.v1, signed_url.accessed.v1, signed_url.revoked.v1, signed_url.expired.v1, **signed_url.audit.exported.v1**, signed\_url.bulk\_revoked.v1

### RBAC/SLO
- **RBAC:** **OWNER/VIEWER** for generate; **SYSTEM** for validate; **ADMIN/COMPLIANCE** for audit
- **SLO:** generate P95 < 100ms; validate P95 < 50ms; batch generate P95 < 500ms; list/export P95 < 300ms

### Limits
- default TTL 1h; max TTL 24h; max 100 URLs per batch

### Idempotency
- by (file_id, user_id, ttl, purpose) for generate

---

## 2.4 lock/

### Stories
- As a **user**, I want to acquire exclusive locks on files so that I can edit without conflicts.
- As a **system**, I want lock expiry with heartbeat renewal so that abandoned locks don't block access indefinitely.
- As a **user**, I want to release locks manually so that I can unblock others when done.
- As a **system**, I want to detect and break stale locks so that files remain accessible.
- **As an admin, I want to list locks by user/file scope so that I can troubleshoot contention.**  *(new)*

### Flow
- AcquireLockCommand(file_id, user_id, lock_duration_seconds=300) → CheckExistingLock() | ValidateDuration(max=3600) | CreateLock(expires_at, lease_id) → CacheInRedis(TTL=duration) → **Outbox:** lock.acquired.v1
- RenewLockCommand(file_id, lease_id, user_id) → ValidateLease() | GuardOwnership() | ExtendLock(+duration) → UpdateCache() → **Outbox:** lock.renewed.v1
- ReleaseLockCommand(file_id, lease_id, user_id) → ValidateLease() | GuardOwnership() | DeleteLock() → InvalidateCache() → **Outbox:** lock.released.v1
- CheckLockQuery(file_id) → RetrieveLock() | ValidateExpiry() → Return lock status + owner
- ForceBreakLockCommand(file_id, admin_id, reason) → GuardAdmin() | DeleteLock() → InvalidateCache() → **Outbox:** lock.force_broken.v1
- **ListLocksQuery(filter_by_user?, filter_by_file?, admin_id) → GuardAdmin() | RetrieveLocks()**  *(new)*

### Projections
- file_locks_read (file_id → owner, lease_id, acquired_at, expires_at)
- lock_lease_cache (Redis: {file_id, lease_id} → lock metadata, TTL = expiry)

### Events
- lock.acquired.v1, lock.renewed.v1, lock.released.v1, lock.expired.v1, lock.force_broken.v1

### RBAC/SLO
- **RBAC:** **OWNER/EDITOR** for acquire/renew/release; **ADMIN** for force break/list
- **SLO:** acquire P95 < 150ms; renew P95 < 100ms; release P95 < 100ms; check P95 < 50ms

### Limits
- default duration 5min; max duration 1h; max 10 concurrent locks per user

### Idempotency
- by (file_id, user_id) for acquire

---

# 3 - CONTENT PIPELINE

## 3.1 media/

### Stories
- As a **user**, I want automatic image processing (resize, thumbnails, format conversion) so that my images are optimized for different use cases.
- As a **user**, I want video transcoding to multiple formats/resolutions so that videos play on all devices.
- As a **system**, I want processing pipelines with retry logic so that transient failures don't lose work.
- As a **user**, I want to track processing status and receive completion notifications so that I know when files are ready.
- As a **system**, I want configurable quality presets so that processing is consistent and efficient.
- **As an operator, I want to cancel media jobs and change their priority so that I can unblock queues.**  *(new)*
- **As a system, I want per-variant quota enforcement so that generated outputs count against usage.**  *(new)*
- As an **admin**, I can **configure media presets per tenant/plan** and have jobs pick the right preset automatically.

### Flow
- ProcessImageCommand(file_id, pipeline_config) → ValidateImageFormat() | LoadImageFromMinIO() | ApplyPipeline(resize, compress, convert) | GenerateThumbnails(sizes[]) | SaveVariants() | UpdateFileMetadata(variants[]) → **Outbox:** media.image.processed.v1 + thumbnail.generated.v1
- ProcessVideoCommand(file_id, transcode_config) → ValidateVideoFormat() | LoadVideoFromMinIO() | TranscodeVideo(formats[], resolutions[], bitrates[]) | GenerateThumbnails(timestamps[]) | SaveVariants() | UpdateFileMetadata(variants[]) → **Outbox:** media.video.processed.v1
- GenerateThumbnailCommand(file_id, dimensions, timestamp?) → LoadMedia() | ExtractFrame(timestamp) | ResizeImage(dimensions) | SaveThumbnail() | LinkToFile() → **Outbox:** thumbnail.generated.v1
- GetMediaJobQuery(job_id) → RetrieveJobStatus() → Return status + progress + outputs
- RetryMediaJobCommand(job_id) → ValidateRetryable() | ResetJob() | EnqueueJob() → **Outbox:** media.job.retried.v1
- **CancelMediaJobCommand(job_id) → GuardOperator() | CancelJob() → Outbox: media.job.cancelled.v1**  *(new)*
- **ChangeMediaJobPriorityCommand(job_id, priority) → GuardOperator() | UpdatePriority() → Outbox: media.job.priority.changed.v1**  *(new)*
- SetTenantMediaPresetsCommand(subject\_id, presets\[\], admin\_id) → GuardAdmin() | Save & cache → **Outbox:** media.presets.updated.v1
- Processing path: Resolve preset via (tenant/plan → default) before enqueue.

### Projections
- media_jobs_read (job_id → file_id, type, status, progress, attempts, error, outputs[], created_at)
- media_variants_read (file_id → variants list with format, size, resolution, url)
- processing_queue_read (pending jobs by priority)
- media\_presets\_read (subject\_id → presets\[\])

### Events
- media.processing.started.v1, media.image.processed.v1, media.video.processed.v1, thumbnail.generated.v1, media.processing.failed.v1, media.job.retried.v1, **media.job.cancelled.v1**, **media.job.priority.changed.v1**, media.presets.updated.v1

### RBAC/SLO
- **RBAC:** **OWNER** to request; **SYSTEM/OPERATOR** to process/cancel/prioritize
- **SLO:** image processing P95 < 5s; video processing P95 < 60s (1080p 1min video); thumbnail generation P95 < 2s

### Limits
- max 3 retries per job; max 10 concurrent jobs per user; max video size 2GB

### Idempotency
- by (file_id, pipeline_config_hash)

---

## 3.2 extraction/

### Stories
- As a **user**, I want OCR extraction from images/PDFs so that text is searchable.
- As a **user**, I want EXIF metadata extraction from photos so that I can see camera settings and location.
- As a **user**, I want text extraction from documents so that content is indexable.
- As a **system**, I want extraction results cached so that repeated queries are fast.
- As a **system**, I want extraction to be asynchronous so that large files don't block.
- **As a user/operator, I want to cancel extraction jobs and retry from UI so that I can recover from failures.**  *(new)*

### Flow
- StartExtractionCommand(file_id, extraction_type) → ValidateFileType() | CreateExtractionJob() | EnqueueJob() → **Outbox:** extraction.started.v1
- PerformOCRJob(file_id) → LoadFile() | RunOCREngine() | ParseText() | StoreResults(text, confidence) | UpdateFileMetadata(searchable_text) → **Outbox:** extraction.ocr.completed.v1
- ExtractEXIFJob(file_id) → LoadImage() | ReadEXIFTags() | ParseMetadata(camera, settings, location, timestamp) | StoreResults() → **Outbox:** extraction.exif.completed.v1
- ExtractTextJob(file_id) → LoadDocument() | ParseContent(pdf, docx, etc) | ExtractText() | StoreResults() → **Outbox:** extraction.text.completed.v1
- GetExtractionResultQuery(file_id, extraction_type) → RetrieveResults() → Return extracted data
- **CancelExtractionCommand(job_id) → GuardOwner/Operator() | CancelJob() → Outbox: extraction.cancelled.v1**  *(new)*
- **RetryExtractionCommand(job_id) → GuardOwner/Operator() | ResetJob() | EnqueueJob() → Outbox: extraction.retried.v1**  *(new)*

### Projections
- extraction_jobs_read (job_id → file_id, type, status, started_at, completed_at, error)
- extraction_results_read (file_id → {ocr_text, exif_data, extracted_text, metadata})
- extraction_cache (Redis: file_id → extraction results, TTL 7d)

### Events
- extraction.started.v1, extraction.ocr.completed.v1, extraction.exif.completed.v1, extraction.text.completed.v1, extraction.failed.v1, **extraction.cancelled.v1**, **extraction.retried.v1**

### RBAC/SLO
- **RBAC:** **OWNER** to request/cancel/retry; **SYSTEM/OPERATOR** to process
- **SLO:** OCR P95 < 10s (single page); EXIF P95 < 1s; text extraction P95 < 5s (10-page doc)

### Limits
- max 100 pages per OCR job; max 3 concurrent extraction jobs per user

### Idempotency
- by (file_id, extraction_type)

---

## 3.3 artifact/

### Stories
- As a **user**, I want to create temporary artifacts (zip archives, previews, reports) with auto-expiry so that I can share collections without permanent storage.
- As a **system**, I want TTL-based cleanup of expired artifacts so that storage doesn't grow unbounded.
- As a **user**, I want to renew artifact expiry so that I can extend access when needed.
- As a **system**, I want artifact generation to be idempotent so that duplicate requests don't create multiple artifacts.
- **As a user, I want to manually delete an artifact before TTL so that I can clean up on demand.**  *(new)*
- **As a user/admin, I want artifact access analytics so that I can see usage.**  *(new)*

### Flow
- CreateZipArtifactCommand(file_ids[], user_id, ttl_hours=24) → ValidateFiles() | GuardAccess() | CreateZipJob() | EnqueueJob() | CreateArtifact(expires_at, status=PENDING) → **Outbox:** artifact.creation.started.v1
- GenerateZipJob(artifact_id) → LoadFiles() | CreateZipArchive() | StoreBlobToMinIO() | UpdateArtifact(status=READY, blob_id) | GenerateSignedURL() → **Outbox:** artifact.zip.created.v1
- CreatePreviewArtifactCommand(file_id, user_id, ttl_hours=1) → ValidateFile() | GuardAccess() | GeneratePreview(low_res, watermark?) | StoreBlobToMinIO() | CreateArtifact() → **Outbox:** artifact.preview.created.v1
- RenewArtifactCommand(artifact_id, user_id, additional_hours) → GuardOwnership() | ExtendExpiry(+hours) | UpdateArtifact() → **Outbox:** artifact.renewed.v1
- ExpireArtifactJob(artifact_id) → DeleteBlob() | DeleteArtifact() → **Outbox:** artifact.expired.v1
- GetArtifactQuery(artifact_id, user_id) → GuardAccess() | RetrieveArtifact() | GenerateSignedURL() → Return artifact metadata + URL
- **DeleteArtifactCommand(artifact_id, user_id) → GuardOwnership() | DeleteNow() → Outbox: artifact.deleted.v1**  *(new)*

### Projections
- artifacts_read (artifact_id → type, file_ids[], blob_id, created_at, expires_at, status, download_count)
- artifact_cache (Redis: artifact_id → metadata + signed_url, TTL = expiry)
- **artifact_analytics_read (artifact_id → views/downloads, unique users, geo, last_access_at)**  *(new)*

### Events
- artifact.creation.started.v1, artifact.zip.created.v1, artifact.preview.created.v1, artifact.renewed.v1, artifact.expired.v1, **artifact.deleted.v1**

### RBAC/SLO
- **RBAC:** **OWNER** for create/renew/delete; **VIEWER** for access
- **SLO:** zip creation P95 < 10s (100 files, 500MB); preview generation P95 < 3s; renew/delete P95 < 150ms

### Limits
- default TTL 24h; max TTL 7d; max 1000 files per zip; max 5 concurrent artifact jobs per user

### Idempotency
- by (user_id, file_ids_hash, type)

---

# 4 - SAFETY & POLICY

## 4.1 policy/

### Stories
- As an **admin**, I want to define file policies (allowed types, max size, virus scan, DLP rules) so that platform content is safe and compliant.
- As a **system**, I want to evaluate files against policies automatically so that violations are detected early.
- As a **user**, I want policy violation reports so that I can remediate issues.
- As an **admin**, I want to update policies with versioning so that changes are trackable.
- As a **system**, I want policies applied at upload time so that bad content never enters storage.
- **As an admin, I want to enable/disable policies without deleting them so that I can toggle enforcement.**  *(new)*
- **As an admin, I want to dry-run/simulate a policy against file sets so that I can validate rules.**  *(new)*

### Flow
- CreatePolicyCommand(admin_id, name, rules) → ValidateRules(mime_types, max_size, dlp_patterns) | CreatePolicy(version=1) → CachePolicy() → **Outbox:** policy.created.v1
- UpdatePolicyCommand(policy_id, rules, admin_id) → ValidateRules() | IncrementVersion() | UpdatePolicy() | InvalidateCache() → **Outbox:** policy.updated.v1
- EvaluatePolicyCommand(file_id, policy_id) → LoadFile() | LoadPolicy() | CheckFileType() | CheckFileSize() | RunDLPChecks(patterns) | RunVirusScan() → RecordResult(violations[]) → **Outbox:** policy.evaluated.v1 + policy.violation.detected.v1 (if violations)
- GetPolicyQuery(policy_id) → RetrievePolicy() → Return policy rules
- ListViolationsQuery(file_id) → RetrieveViolations() → Return violation list
- **TogglePolicyCommand(policy_id, enabled, admin_id) → GuardAdmin() | UpdateActiveFlag() → Outbox: policy.toggled.v1**  *(new)*
- **SimulatePolicyCommand(policy_id, file_ids[], admin_id) → GuardAdmin() | EvaluateNoSideEffects() → Outbox: policy.simulation.completed.v1**  *(new)*

### Projections
- policies_read (policy_id → name, rules, version, active, created_at, updated_at)
- policy_violations_read (file_id → violations list with policy_id, rule, severity, detected_at)
- policy_cache (Redis: policy_id → rules, TTL 1h)

### Events
- policy.created.v1, policy.updated.v1, policy.evaluated.v1, policy.violation.detected.v1, policy.violation.resolved.v1, **policy.toggled.v1**, **policy.simulation.completed.v1**

### RBAC/SLO
- **RBAC:** **ADMIN** for create/update/toggle/simulate; **SYSTEM** for evaluate
- **SLO:** evaluate P95 < 2s; create/update/toggle P95 < 300ms; list violations P95 < 150ms

### Limits
- max 100 policies; max 50 DLP patterns per policy

### Idempotency
- by (policy_name) for create; by (policy_id, version) for update

---

## 4.2 scan/

### Stories
- As a **system**, I want to run antivirus scans on all uploaded files so that malware is detected.
- As a **system**, I want DLP scans to detect sensitive data (PII, credentials, secrets) so that leaks are prevented.
- As a **system**, I want scan results with severity levels so that high-risk files are prioritized.
- As a **user**, I want scan status visible in file metadata so that I know if files are safe.
- As a **system**, I want to queue scans and process them asynchronously so that uploads aren't blocked.
- **As an operator, I want to cancel queued scans, change priority, and bulk rescan so that I can manage incidents.**  *(new)*

### Flow
- StartScanCommand(file_id, scan_type) → ValidateFile() | CreateScanJob(type=AV/DLP, status=QUEUED) | EnqueueJob() → **Outbox:** scan.started.v1
- PerformAVScanJob(file_id) → LoadFileFromMinIO() | RunAVEngine() | ParseResults(threat_name, severity) | RecordScanResult(verdict=CLEAN/INFECTED/SUSPICIOUS) → If INFECTED: TriggerQuarantine() → **Outbox:** scan.av.completed.v1 + file.quarantined.v1 (if infected)
- PerformDLPScanJob(file_id) → LoadFile() | ExtractText() | RunDLPPatterns(ssn, credit_card, api_key) | RecordFindings(matches[], offsets[], confidence) → If findings: RecordResult(verdict=FLAGGED) → **Outbox:** scan.dlp.completed.v1
- GetScanResultQuery(file_id) → RetrieveScanResults() → Return scan history + current verdict
- RescanFileCommand(file_id, user_id) → GuardOwnership() | CreateScanJob() | EnqueueJob() → **Outbox:** scan.started.v1
- **CancelScanCommand(job_id, operator_id) → GuardOperator() | Cancel() → Outbox: scan.cancelled.v1**  *(new)*
- **ChangeScanPriorityCommand(job_id, priority, operator_id) → GuardOperator() | UpdatePriority() → Outbox: scan.priority.changed.v1**  *(new)*
- **BulkRescanCommand(filter, operator_id) → GuardOperator() | CreateBatchJobs() → Outbox: scan.bulk.started.v1**  *(new)*

### Projections
- scan_jobs_read (job_id → file_id, type, status, started_at, completed_at, verdict, findings[])
- scan_results_read (file_id → av_verdict, dlp_verdict, last_scan_at, findings_summary)
- scan_queue_read (pending scans by priority)

### Events
- scan.started.v1, scan.av.completed.v1, scan.dlp.completed.v1, scan.failed.v1, scan.verdict.changed.v1, **scan.cancelled.v1**, **scan.priority.changed.v1**, **scan.bulk.started.v1**

### RBAC/SLO
- **RBAC:** **SYSTEM** for scan; **OWNER** for rescan; **OPERATOR** for cancel/reprioritize/bulk
- **SLO:** AV scan P95 < 5s (10MB file); DLP scan P95 < 10s (10-page doc); queue wait P95 < 1min

### Limits
- max 3 scan retries; max 100MB for DLP scan; rescan cooldown 1h

### Idempotency
- by (file_id, scan_type, file_hash)

---

## 4.3 quarantine/

### Stories
- As a **system**, I want to automatically quarantine files with malware or severe policy violations so that they can't be accessed.
- As an **admin**, I want to review quarantined files and decide to release or permanently delete them.
- As a **user**, I want notifications when my files are quarantined so that I can take action.
- As a **system**, I want quarantine to block all access (download, share, preview) so that infected files can't spread.
- **As an admin, I want to extend quarantine windows and perform bulk actions so that I can manage large incidents.**  *(new)*

### Flow
- QuarantineFileCommand(file_id, reason, severity, placed_by) → MarkQuarantined() | BlockAccess() | NotifyOwner() | LogEvent() → **Outbox:** file.quarantined.v1
- ReleaseQuarantineCommand(file_id, admin_id, reason) → GuardAdmin() | ValidateClean() | UnmarkQuarantined() | RestoreAccess() | NotifyOwner() → **Outbox:** quarantine.released.v1
- DeleteQuarantinedFileCommand(file_id, admin_id, reason) → GuardAdmin() | PermanentDelete() | NotifyOwner() → **Outbox:** quarantine.deleted.v1
- ListQuarantinedFilesQuery(admin_id) → RetrieveQuarantinedFiles() → Return list with reasons + severity
- GetQuarantineDetailsQuery(file_id, admin_id) → RetrieveQuarantineInfo() → Return quarantine metadata + scan results
- **ExtendQuarantineCommand(file_id, days, admin_id) → GuardAdmin() | ExtendWindow(days) → Outbox: quarantine.extended.v1**  *(new)*
- **BulkQuarantineActionCommand(action, filters, admin_id) → GuardAdmin() | ApplyBulk() → Outbox: quarantine.bulk_actioned.v1**  *(new)*

### Projections
- quarantined_files_read (file_id → reason, severity, placed_by, placed_at, scan_results, owner_notified)
- quarantine_review_queue_read (pending review by severity)

### Events
- file.quarantined.v1, quarantine.released.v1, quarantine.deleted.v1, quarantine.review.requested.v1, **quarantine.extended.v1**, **quarantine.bulk_actioned.v1**

### RBAC/SLO
- **RBAC:** **SYSTEM** for quarantine; **ADMIN** for release/delete/extend/bulk; **OWNER** for view
- **SLO:** quarantine P95 < 500ms; release P95 < 1s; list P95 < 200ms

### Limits
- auto-delete quarantined files after 90 days; max 10,000 quarantined files

### Idempotency
- by (file_id) for quarantine

---

## 4.4 file_flag/

### Stories
- As a **user**, I want to flag files for review (malware, copyright violation, policy breach) so that problematic content is reviewed.
- As a **moderator**, I want to review flags and take action (resolve, dismiss, escalate) so that cases are closed.
- As a **system**, I want flag aggregation so that files with multiple flags are prioritized.
- As a **user**, I want flag status updates so that I know the outcome of my report.
- **As a moderator, I want bulk triage and SLA-based auto-escalation so that critical items surface.**  *(new)*

### Flow
- FlagFileCommand(file_id, flagged_by, reason, details) → ValidateReason(malware, copyright, policy_violation) | CreateFlag(status=OPEN) | NotifyModerators() → **Outbox:** file.flagged.v1
- ResolveFlagCommand(flag_id, moderator_id, action, resolution_notes) → GuardModerator() | UpdateFlag(status=RESOLVED, action) | ExecuteAction(quarantine, delete, clear) | NotifyFlagger() → **Outbox:** flag.resolved.v1
- DismissFlagCommand(flag_id, moderator_id, reason) → GuardModerator() | UpdateFlag(status=DISMISSED) | NotifyFlagger() → **Outbox:** flag.dismissed.v1
- ListFlagsQuery(moderator_id, status?) → RetrieveFlags(filters) → Return flag list with priority
- GetFlagDetailsQuery(flag_id, moderator_id) → RetrieveFlagDetails() → Return flag + file + scan results
- **BulkFlagTriageCommand(filters, action, moderator_id) → GuardModerator() | ApplyBulk() → Outbox: flag.bulk_actioned.v1**  *(new)*
- **AutoEscalateFlagJob(flag_id) → SLAWindowExceeded? → Escalate() → Outbox: flag.escalated.v1**  *(new)*

### Projections
- file_flags_read (flag_id → file_id, flagged_by, reason, status, created_at, resolved_at, moderator_id)
- flag_queue_read (open flags by priority/severity)
- flagged_files_read (file_id → flag_count, reasons[], status)

### Events
- file.flagged.v1, flag.resolved.v1, flag.dismissed.v1, flag.escalated.v1, **flag.bulk_actioned.v1**

### RBAC/SLO
- **RBAC:** **USER** for flag; **MODERATOR** for resolve/dismiss/bulk
- **SLO:** flag creation P95 < 300ms; resolve P95 < 500ms; list P95 < 200ms

### Limits
- max 10 flags per file; cooldown 1h per user per file

### Idempotency
- by (file_id, flagged_by, reason) for flag

---

# 5 - COMPLIANCE & LIFECYCLE

## 5.1 lifecycle/

### Stories
- As an **admin**, I want to define lifecycle rules (retention periods, auto-delete policies) so that data is managed according to regulations.
- As a **system**, I want to apply lifecycle rules automatically based on file age or state so that compliance is enforced.
- As an **admin**, I want to set legal holds on files so that they can't be deleted during investigations.
- As a **user**, I want to see lifecycle status on files so that I know when they'll be deleted.
- As a **system**, I want to execute lifecycle transitions (soft delete, permanent delete) on schedule.
- **As an admin, I want to archive/unarchive files (separate from soft-delete) so that long-term storage policies are applied.**  *(new)*

### Flow
- DefineLifecycleRuleCommand(admin_id, rule_name, conditions, actions) → ValidateRule(retention_days, state_transitions) | CreateRule() → **Outbox:** lifecycle.rule.created.v1
- UpdateLifecycleRuleCommand(rule_id, updates, admin_id) → ValidateUpdates() | UpdateRule() | RefreshAffectedFiles() → **Outbox:** lifecycle.rule.updated.v1
- ApplyLifecycleRulesJob() → LoadActiveRules() | ForEachRule: MatchFiles(conditions) | ExecuteActions(soft_delete, archive, delete) → **Outbox:** lifecycle.rule.applied.v1 + file.soft_deleted.v1/file.archived.v1/file.deleted.v1
- PlaceLegalHoldCommand(file_ids[], admin_id, case_id, reason) → ForEach: MarkLegalHold(expires_at=null) | SuspendLifecycle() → **Outbox:** legal_hold.placed.v1
- RemoveLegalHoldCommand(file_ids[], admin_id, case_id, reason) → ForEach: UnmarkLegalHold() | ResumeLifecycle() → **Outbox:** legal_hold.removed.v1
- **ArchiveFileCommand(file_id, admin_id) → GuardAdmin() | MarkArchived() → Outbox: file.archived.v1**  *(new)*
- **UnarchiveFileCommand(file_id, admin_id) → GuardAdmin() | ClearArchived() → Outbox: file.unarchived.v1**  *(new)*
- GetFileLifecycleQuery(file_id) → RetrieveLifecycleStatus() → Return rule, retention_end, legal_hold

### Projections
- lifecycle_rules_read (rule_id → name, conditions, actions, active, created_at)
- file_lifecycle_read (file_id → rule_id, retention_end, legal_hold, next_action, archived_at?)
- legal_holds_read (file_id → case_id, placed_by, placed_at, reason)

### Events
- lifecycle.rule.created.v1, lifecycle.rule.updated.v1, lifecycle.rule.applied.v1, legal_hold.placed.v1, legal_hold.removed.v1, **file.archived.v1**, **file.unarchived.v1**

### RBAC/SLO
- **RBAC:** **ADMIN** for rules/legal holds/archive ops; **SYSTEM** for apply
- **SLO:** define rule P95 < 300ms; apply rules batch P95 < 10s (1000 files); place/remove hold P95 < 500ms

### Limits
- max 50 active rules; max 10,000 files per legal hold

### Idempotency
- by (rule_name) for create; by (file_id, case_id) for legal hold

---

## 5.2 audit/

### Stories
- As an **admin**, I want comprehensive audit logs (upload, access, delete, share, scan) so that all file operations are traceable.
- As a **compliance officer**, I want to export audit logs in standard formats so that I can provide them for audits.
- As a **system**, I want append-only audit records so that logs can't be tampered with.
- As an **admin**, I want to search audit logs by user, file, action, date so that investigations are efficient.
- As a **system**, I want audit logs with PII minimization so that sensitive data is protected.
- **As a compliance officer, I want to filter by IP/UA and stream to SIEM/external sinks so that security teams can monitor.**  *(new)*
- As **compliance**, I want a **scheduled audit-log integrity check** (hash chain/sequence) with alerts on any break.

### Flow
- AppendAuditRecordCommand(action, actor_id, resource_id, resource_type, ip, user_agent, metadata) → ValidateAction() | RedactPII(metadata) | AppendToLog(immutable) → **Outbox:** audit.recorded.v1
- ListAuditRecordsQuery(filters, admin_id) → GuardAdmin() | RetrieveAuditLog(filters: user, file, action, date_range) → Return paginated records
- ExportAuditLogsCommand(filters, format, admin_id) → GuardAdmin() | RetrieveAuditLog(filters) | FormatOutput(JSON/CSV) | CreateExportArtifact() → **Outbox:** audit.export.created.v1
- SearchAuditLogsQuery(query, admin_id) → GuardAdmin() | FullTextSearch(actions, resources) → Return matches
- **EnableAuditStreamingCommand(target, admin_id) → GuardAdmin() | ConfigureStreaming() → Outbox: audit.streaming.enabled.v1**  *(new)*
- VerifyAuditChainJob() → Recompute chain/hash windows | If OK: **Outbox:** audit.integrity.verified.v1 | If broken: **Outbox:** audit.integrity.alerted.v1

### Projections
- audit_log_read (record_id → timestamp, action, actor, resource, ip, metadata, hash_chain)
- audit_export_artifacts_read (export_id → filters, format, artifact_url, created_at, expires_at)

### Events
- audit.recorded.v1, audit.export.created.v1, **audit.streaming.enabled.v1**
- audit.integrity.verified.v1, audit.integrity.alerted.v1

### RBAC/SLO
- **RBAC:** **SYSTEM** for append; **ADMIN/COMPLIANCE** for list/export
- **SLO:** append P95 < 50ms; list P95 < 300ms; export P95 < 10s (10k records)

### Limits
- retain audit logs 7 years; max 1M records per export

### Idempotency
- by (action, actor, resource, timestamp) for append

---

# 6 - STORAGE PRIMITIVES & OPS

## 6.1 blob/

### Stories
- As a **system**, I want content-addressed blob storage so that duplicate data is automatically de-duplicated.
- As a **system**, I want reference counting on blobs so that blobs are only deleted when no files reference them.
- As a **system**, I want to track blob storage class (hot, warm, cold) so that costs are optimized.
- As a **system**, I want blob integrity verification with checksums so that corruption is detected.
- **As a system, I want scheduled integrity verification and verify-on-read hooks so that silent corruption is caught.**  *(new)*
- As a **system**, I **auto-tier blobs** (HOT↔WARM↔COLD) based on last access, and allow admins to override storage class.

### Flow
- PutBlobCommand(data_stream, user_id) → ComputeHash(SHA256) | CheckExisting(hash) → If exists: IncrementRefCount() ELSE: StoreToMinIO(hash) | CreateBlobMetadata(hash, size, storage_class=HOT, ref_count=1) → **Outbox:** blob.created.v1
- LinkBlobToFileCommand(blob_id, file_id) → IncrementRefCount() | CreateReference() → **Outbox:** blob.linked.v1
- UnlinkBlobFromFileCommand(blob_id, file_id) → DecrementRefCount() | DeleteReference() → If ref_count=0: MarkForGC() → **Outbox:** blob.unlinked.v1
- GetBlobByHashQuery(hash) → RetrieveBlobMetadata() → Return blob location + metadata
- VerifyBlobIntegrityCommand(blob_id) → LoadBlobFromMinIO() | ComputeHash() | CompareWithStored() → **Outbox:** blob.verified.v1/blob.corrupted.v1
- **ScheduleIntegritySweepCommand(sample_rate) → CreatePlan() → Outbox: blob.integrity.sweep.scheduled.v1**  *(new)*
- **VerifyOnReadHook(file_id) → FeatureFlag? → VerifyBlobIntegrity() → emit verified/corrupted**  *(new)*
- AutoTieringJob() → Scan blobs\_read.last\_access\_at | Determine target class by policy | Move/rehydrate as needed | **Outbox:** blob.storage\_class.changed.v1
- OverrideBlobStorageClassCommand(blob\_id, class, admin\_id) → GuardAdmin() | Apply | **Outbox:** blob.storage\_class.changed.v1
    

### Projections
- blobs_read (blob_id → hash, size, storage_class, ref_count, created_at, location)
- blob_reference_counts_read (blob_id → file_count)
- gc_candidates_read (blobs with ref_count=0)
- blob\_access\_stats\_read (blob\_id → last\_access\_at, access\_count, current\_class)

### Events
- blob.created.v1, blob.linked.v1, blob.unlinked.v1, blob.verified.v1, blob.corrupted.v1, blob.gc_marked.v1, **blob.integrity.sweep.scheduled.v1**
- blob.storage\_class.changed.v1


### RBAC/SLO
- **RBAC:** **SYSTEM**
- **SLO:** put P95 < 2s (10MB); link/unlink P95 < 100ms; verify P95 < 3s (10MB)

### Limits
- max blob size 5GB; storage classes: HOT (7d), WARM (30d), COLD (365d+)

### Idempotency
- by (hash) for put

---

## 6.2 reference/

### Stories
- As a **system**, I want to track cross-service file references (user profile, job attachment, proposal doc) so that files can't be deleted while in use.
- As a **system**, I want reference validation before file deletion so that broken links are prevented.
- As a **service**, I want to add/remove references when attaching/detaching files so that tracking is automatic.
- **As a system, I want to list references by aggregate (service/entity) so that producers can audit usage.**  *(new)*
- As a **producer service**, I can **list file references by aggregate** (service/entity) to audit attachments.

### Flow
- AddReferenceCommand(aggregate_type, aggregate_id, file_id, purpose) → CreateReference(service_name, entity_id, purpose) → **Outbox:** reference.added.v1
- RemoveReferenceCommand(aggregate_type, aggregate_id, file_id) → DeleteReference() → **Outbox:** reference.removed.v1
- ListReferencesByFileQuery(file_id) → RetrieveReferences() → Return reference list (service, entity, purpose)
- ValidateNoReferencesCommand(file_id) → CountReferences() → If count > 0: Fail(REFERENCES_EXIST) ELSE: Allow
- **ListReferencesByAggregateQuery(aggregate_type, aggregate_id) → RetrieveRefs() → Return file list**  *(new)*
- ListReferencesByAggregateQuery(aggregate\_type, aggregate\_id) → Return file list.

### Projections
- file_references_read (file_id → references list with aggregate_type, aggregate_id, purpose, created_at)
- reference_counts_read (file_id → reference_count by service)

### Events
- reference.added.v1, reference.removed.v1

### RBAC/SLO
- **RBAC:** **SYSTEM**
- **SLO:** add/remove P95 < 150ms; list P95 < 200ms; validate P95 < 100ms

### Limits
- max 100 references per file

### Idempotency
- by (aggregate_type, aggregate_id, file_id) for add

---

## 6.3 quota/

### Stories
- As an **admin**, I want to set storage quotas per user/tenant so that usage is controlled.
- As a **system**, I want to enforce quotas at upload time so that overages are prevented.
- As a **user**, I want to see my quota usage so that I can manage my storage.
- As a **system**, I want quota warnings at 80% and 90% so that users can take action before hitting limits.
- **As an admin, I want to apply/change plan tiers that automatically set quotas so that plans stay in sync.**  *(new)*

### Flow
- SetQuotaCommand(user_id, quota_gb, admin_id) → ValidateQuota(min=1GB, max=1TB) | CreateOrUpdateQuota() → **Outbox:** quota.set.v1
- AdjustQuotaCommand(user_id, delta_bytes, operation) → LoadQuota() | If operation=ADD: IncrementUsage() ELSE: DecrementUsage() | CheckThresholds(80%, 90%, 100%) → If threshold crossed: **Outbox:** quota.threshold.crossed.v1
- GetQuotaQuery(user_id) → RetrieveQuota() → Return quota_total, quota_used, quota_remaining, percentage
- EnforceQuotaCommand(user_id, upload_size) → LoadQuota() | CheckRemaining(quota_used + upload_size <= quota_total) → If exceeded: Fail(QUOTA_EXCEEDED) ELSE: Allow
- **ApplyPlanTierCommand(subject_id, tier, admin_id) → GuardAdmin() | SetPresetLimits() → Outbox: quota.plan.applied.v1**  *(new)*
- **ChangePlanTierCommand(subject_id, new_tier, admin_id) → GuardAdmin() | UpdateLimits() → Outbox: quota.plan.changed.v1**  *(new)*

### Projections
- user_quotas_read (user_id → quota_total, quota_used, quota_remaining, last_updated)
- quota_usage_cache (Redis: user_id → quota_used, TTL 5m)
- quota_exceeded_users_read (users at 100%)

### Events
- quota.set.v1, quota.adjusted.v1, quota.threshold.crossed.v1, quota.exceeded.v1, **quota.plan.applied.v1**, **quota.plan.changed.v1**

### RBAC/SLO
- **RBAC:** **ADMIN** for set/apply/change; **SYSTEM** for adjust/enforce
- **SLO:** enforce P95 < 50ms (cached); adjust P95 < 100ms; get P95 < 100ms

### Limits
- default quota 10GB; max quota 1TB; threshold warnings at 80%, 90%

### Idempotency
- by (user_id, operation, timestamp)

---

## 6.4 namespace/

### Stories
- As a **system**, I want to map tenants to storage buckets/zones so that multi-tenancy is isolated.
- As an **admin**, I want to create namespaces with custom settings so that tenants have dedicated storage.
- As a **system**, I want namespace resolution to be fast so that storage operations aren't delayed.
- **As an admin, I want to disable/delete a namespace and migrate files between namespaces/zones so that I can manage residency.**  *(new)*
- As an **admin**, I can **list namespaces** and view settings/health (bucket, region, residency, status) to manage data residency.

### Flow
- CreateNamespaceCommand(tenant_id, bucket_name, region, admin_id) → ValidateBucket() | CreateNamespace(settings) → CacheNamespace() → **Outbox:** namespace.created.v1
- ResolveNamespaceQuery(tenant_id) → LoadNamespace() | CacheHit() → Return bucket + region + settings
- UpdateNamespaceCommand(namespace_id, settings, admin_id) → ValidateSettings() | UpdateNamespace() | InvalidateCache() → **Outbox:** namespace.updated.v1
- **DisableNamespaceCommand(namespace_id, admin_id) → GuardAdmin() | Disable() → Outbox: namespace.disabled.v1**  *(new)*
- **MigrateNamespaceCommand(namespace_id, target_zone_or_bucket, admin_id) → GuardAdmin() | Plan+Execute() → Outbox: namespace.migration.planned.v1 → namespace.migration.completed.v1**  *(new)*
- ListNamespacesQuery(filters{tenant\_id?, region?, status?}, page, limit) → Return paginated namespaces
- GetNamespaceDetailQuery(namespace\_id) → Return settings + health metrics

### Projections
- namespaces_read (namespace_id → tenant_id, bucket, region, settings, created_at)
- namespace_cache (Redis: tenant_id → namespace, TTL 1h)
- namespace\_admin\_read (namespace\_id → settings, health, last\_migration\_at)

### Events
- namespace.created.v1, namespace.updated.v1, **namespace.disabled.v1**, **namespace.migration.planned.v1**, **namespace.migration.completed.v1**

### RBAC/SLO
- **RBAC:** **ADMIN** for create/update/disable/migrate; **SYSTEM** for resolve
- **SLO:** resolve P95 < 50ms (cached); create/update P95 < 500ms

### Limits
- max 10,000 namespaces

### Idempotency
- by (tenant_id) for create

---

## 6.5 gc/

### Stories
- As a **system**, I want scheduled garbage collection to delete unreferenced blobs so that storage costs are minimized.
- As a **system**, I want GC to run during low-traffic periods so that performance isn't impacted.
- As an **admin**, I want GC reports so that I can see reclaimed storage.
- As a **system**, I want GC to be safe (verify zero references) so that in-use blobs aren't deleted.
- **As an operator, I want to cancel in-progress GC and run dry-run previews so that I can iterate safely.**  *(new)*

### Flow
- PlanGCJob(admin_id?) → LoadGCCandidates(ref_count=0, marked_at < now-7d) | CalculateStorageSavings() | CreateGCPlan() → **Outbox:** gc.planned.v1
- RunGCJob(gc_plan_id) → ForEachBlob: VerifyZeroReferences() | DeleteFromMinIO() | DeleteBlobMetadata() | RecordDeletion() → GenerateReport(deleted_count, reclaimed_bytes) → **Outbox:** gc.completed.v1
- GetGCReportQuery(gc_plan_id, admin_id) → RetrieveReport() → Return statistics
- ScheduleGCCommand(cron_expression, admin_id) → ValidateCron() | CreateSchedule() → **Outbox:** gc.scheduled.v1
- **DryRunGCCommand(scope, admin_id) → GuardAdmin() | SimulateDeletes() → Outbox: gc.dry_run.completed.v1**  *(new)*
- **CancelGCCommand(gc_plan_id, admin_id) → GuardAdmin() | CancelPlan() → Outbox: gc.cancelled.v1**  *(new)*

### Projections
- gc_plans_read (plan_id → created_at, started_at, completed_at, blob_count, reclaimed_bytes, status)
- gc_candidates_read (blobs eligible for GC with marked_at, ref_count)
- gc_reports_read (historical GC statistics)

### Events
- gc.planned.v1, gc.started.v1, gc.completed.v1, gc.failed.v1, gc.scheduled.v1, **gc.dry_run.completed.v1**, **gc.cancelled.v1**

### RBAC/SLO
- **RBAC:** **ADMIN** for plan/schedule/cancel/dry-run; **SYSTEM** for run
- **SLO:** plan P95 < 5s (10k blobs); run P95 < 10s per 1000 blobs; report P95 < 200ms

### Limits
- GC grace period 7 days; max 100k blobs per run; run frequency min 1h

### Idempotency
- by (gc_plan_id) for run

---

## 6.6 encryption/

### Stories
- As a **system**, I want to track encryption keys and their rotation schedule so that data is encrypted at rest.
- As an **admin**, I want to rotate encryption keys without downtime so that security is maintained.
- As a **system**, I want to re-encrypt files with new keys during rotation so that old keys can be retired.
- As a **compliance officer**, I want encryption audit logs so that key usage is traceable.
- **As a security admin, I want emergency key revoke and pause/resume rotation windows so that compromises can be handled.**  *(new)*

### Flow
- TrackKeyCommand(key_id, version, created_at, expires_at) → CreateKeyRecord() → **Outbox:** encryption.key.tracked.v1
- ScheduleKeyRotationCommand(key_id, rotation_date, admin_id) → ValidateDate() | CreateRotationJob() → **Outbox:** encryption.rotation.scheduled.v1
- RotateKeyJob(key_id) → GenerateNewKey() | ReEncryptFiles(old_key, new_key) | UpdateFileMetadata(key_version) | RetireOldKey() → **Outbox:** encryption.key.rotated.v1
- GetKeyStatusQuery(key_id, admin_id) → RetrieveKeyMetadata() → Return version, expiry, file_count
- AuditKeyUsageQuery(key_id, admin_id) → RetrieveKeyAudit() → Return usage logs
- **RevokeKeyCommand(key_id, security_admin) → GuardSecurity() | RevokeNow() | QueueReEncrypt() → Outbox: encryption.key.revoked.v1**  *(new)*
- **PauseRotationCommand(key_id, admin_id) → GuardAdmin() | Pause() → Outbox: encryption.rotation.paused.v1**  *(new)*
- **ResumeRotationCommand(key_id, admin_id) → GuardAdmin() | Resume() → Outbox: encryption.rotation.resumed.v1**  *(new)*

### Projections
- encryption_keys_read (key_id → version, created_at, expires_at, rotation_schedule, status)
- key_file_mapping_read (key_id → file_count)
- key_rotation_jobs_read (job_id → key_id, status, started_at, completed_at, files_reencrypted)

### Events
- encryption.key.tracked.v1, encryption.rotation.scheduled.v1, encryption.key.rotated.v1, encryption.key.retired.v1, **encryption.key.revoked.v1**, **encryption.rotation.paused.v1**, **encryption.rotation.resumed.v1**

### RBAC/SLO
- **RBAC:** **ADMIN** for schedule/rotate/pause/resume; **SYSTEM** for track; **SECURITY_ADMIN** for revoke
- **SLO:** track P95 < 100ms; schedule P95 < 300ms; rotation P95 < 1min per 1000 files

### Limits
- key expiry max 365 days; rotation window 30 days; max 10 concurrent rotations

### Idempotency
- by (key_id, version) for track

---

# 7 - INBOX EVENTS (Consumers)

## 7.1 Inbox: User Events
### Stories
- As a **system**, I want to consume user.created events to initialize user storage quotas.
- As a **system**, I want to consume user.deleted events to cleanup user files and enforce data retention.
- As a **system**, on **user.updated.v1** I refresh ACL/effective ownership contexts for impacted files/folders.
    
- As a **system**, on **contract.state.changed.v1** I apply lifecycle/holds (e.g., archive on completion, hold on dispute).

### Flow
- Consume: user.created.v1 → CreateUserQuota(user_id, plan_preset) → **Outbox:** quota.initialized.v1
- Consume: user.deleted.v1 → ListUserFiles(user_id) | SoftDeleteAllFiles() | SchedulePermanentDeletion(after=90d) → **Outbox:** user.files.deleted.v1
- Consume: user.updated.v1 → RebuildEffectiveACLs(user\_id) → warm caches → Outbox: access.updated.v1 (if deltas)
- Consume: contract.state.changed.v1 → Match policy (completed|disputed|terminated) → Archive or place legal hold → Outbox: file.archived.v1 / legal\_hold.placed.v1

### Projections
- user_storage_read (user_id → quota, used, file_count)

### Events
- (consumer)

### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 500ms per event

---

## 7.2 Inbox: Job Events
### Stories
- As a **system**, I want to consume job.attachment.added events to create file references.
- As a **system**, I want to consume job.deleted events to cleanup job attachments.

### Flow
- Consume: job.attachment.added.v1 → AddReference(aggregate_type=JOB, aggregate_id=job_id, file_id) → **Outbox:** reference.added.v1
- Consume: job.deleted.v1 → RemoveReferences(aggregate_type=JOB, aggregate_id=job_id) | CheckOrphanFiles() → **Outbox:** reference.removed.v1

### Projections
- job_file_references_read

### Events
- (consumer)

### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 300ms per event

---

## 7.3 Inbox: Proposal Events
### Stories
- As a **system**, I want to consume proposal.attachment.added events to create file references.
- As a **system**, I want to consume proposal.deleted events to cleanup proposal attachments.

### Flow
- Consume: proposal.attachment.added.v1 → AddReference(aggregate_type=PROPOSAL, aggregate_id=proposal_id, file_id) → **Outbox:** reference.added.v1
- Consume: proposal.deleted.v1 → RemoveReferences(aggregate_type=PROPOSAL, aggregate_id=proposal_id) | CheckOrphanFiles() → **Outbox:** reference.removed.v1

### Projections
- proposal_file_references_read

### Events
- (consumer)

### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 300ms per event

---

## 7.4 Inbox: Contract Events
### Stories
- As a **system**, I want to consume contract.deliverable.uploaded events to create file references.
- As a **system**, I want to consume contract.completed events to archive contract files.

### Flow
- Consume: contract.deliverable.uploaded.v1 → AddReference(aggregate_type=CONTRACT, aggregate_id=contract_id, file_id) → **Outbox:** reference.added.v1
- Consume: contract.completed.v1 → ArchiveContractFiles(contract_id, archive_policy=7y) | ApplyLifecycleRule() → **Outbox:** file.archived.v1

### Projections
- contract_file_references_read

### Events
- (consumer)

### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 500ms per event

---

## 7.5 Inbox: Admin Events
### Stories
- As a **system**, I want to consume admin.file.removed events to force delete files.
- As a **system**, I want to consume admin.policy.updated events to refresh policy cache.
- **As a system, I want to consume admin.moderation.actioned events to quarantine/release files and revoke shares/links.**  *(new)*

### Flow
- Consume: admin.file.removed.v1 → ForceDeleteFile(file_id) | NotifyOwner() → **Outbox:** file.deleted.v1
- Consume: admin.policy.updated.v1 → InvalidatePolicyCache(policy_id) | RefreshPolicy()
- **Consume: admin.moderation.actioned.v1 → ApplyModerationAction(quarantine|release|revoke_shares|revoke_links) → Outbox: file.quarantined.v1/quarantine.released.v1/share.bulk_revoked.v1/signed_url.revoked.v1**  *(new)*

### Projections
- admin_actions_audit_read

### Events
- (consumer)

### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 300ms per event

---

# APPENDIX

## A. Storage Classes
- **HOT:** Files accessed within last 7 days; stored in fast SSD; highest cost  
- **WARM:** Files accessed 8-30 days ago; stored in standard HDD; medium cost  
- **COLD:** Files accessed 31-365 days ago; stored in archive tier; lowest cost  
- **GLACIER:** Files older than 365 days; deep archive; retrieval delay

## B. File Types
- **DOCUMENT:** pdf, docx, xlsx, pptx, txt, md  
- **IMAGE:** jpg, png, gif, webp, svg, bmp  
- **VIDEO:** mp4, avi, mov, mkv, webm  
- **AUDIO:** mp3, wav, ogg, flac, m4a  
- **ARCHIVE:** zip, tar, gz, rar, 7z  
- **CODE:** js, py, go, java, cpp, rs

## C. MIME Type Whitelist
- documents: application/pdf, application/msword, application/vnd.openxmlformats-officedocument.*
- images: image/jpeg, image/png, image/gif, image/webp, image/svg+xml
- videos: video/mp4, video/mpeg, video/quicktime, video/x-msvideo
- audio: audio/mpeg, audio/wav, audio/ogg, audio/flac
- archives: application/zip, application/x-tar, application/gzip

## D. Default Quotas by User Type
- **FREE:** 5GB storage, 100MB max file size, 500 files  
- **BASIC:** 50GB storage, 500MB max file size, 5,000 files  
- **PRO:** 500GB storage, 2GB max file size, 50,000 files  
- **ENTERPRISE:** 5TB storage, 5GB max file size, unlimited files

## E. Scan Engines
- **Antivirus:** ClamAV (open-source)  
- **DLP Patterns:** Regex for SSN, credit cards, API keys, emails, phone numbers  
- **Content Moderation:** AI-based image classification for NSFW, violence, hate symbols

## F. Media Processing Presets
### Image Presets
- **thumbnail:** 150x150px, JPEG 80% quality
- **small:** 640x480px, JPEG 85% quality
- **medium:** 1280x720px, JPEG 90% quality
- **large:** 1920x1080px, JPEG 95% quality
- **webp_conversion:** WebP format, 85% quality

### Video Presets
- **mobile:** 480p, H.264, 500kbps
- **sd:** 720p, H.264, 1.5Mbps
- **hd:** 1080p, H.264, 5Mbps
- **thumbnails:** Extract frames at 0s, 25%, 50%, 75%, 100%

## G. Rate Limits
- **Upload:** 100 uploads/hour per user  
- **Download:** 1000 downloads/hour per user  
- **Share Link Creation:** 50 links/hour per user  
- **Scan Request:** 10 rescans/hour per user  
- **Flag Submission:** 10 flags/hour per user

## H. Retention Policies
- **Soft Deleted Files:** 30 days recovery window  
- **Audit Logs:** 7 years retention  
- **Quarantined Files:** 90 days before auto-delete  
- **Upload Sessions:** 24 hours before expiry  
- **Signed URLs:** Default 1 hour, max 24 hours  
- **Artifacts:** Default 24 hours, max 7 days

## I. Event Topics (canonical, singularized + new items)
- **file.\***: file.uploaded, file.deleted, file.soft_deleted, file.restored, file.metadata.updated, file.moved, file.quarantined, **file.archived**, **file.unarchived**, **file.copied**, **file.ownership.transferred**
- **blob.\***: blob.created, blob.linked, blob.unlinked, blob.verified, blob.corrupted, blob.gc_marked, **blob.integrity.sweep.scheduled**
- **media.\***: media.processing.started, media.image.processed, media.video.processed, media.processing.failed, thumbnail.generated, media.job.retried, **media.job.cancelled**, **media.job.priority.changed**
- **scan.\***: scan.started, scan.av.completed, scan.dlp.completed, scan.failed, scan.verdict.changed, **scan.cancelled**, **scan.priority.changed**, **scan.bulk.started**
- **quarantine.\***: file.quarantined, quarantine.released, quarantine.deleted, quarantine.review.requested, **quarantine.extended**, **quarantine.bulk_actioned**
- **policy.\***: policy.created, policy.updated, policy.evaluated, policy.violation.detected, policy.violation.resolved, **policy.toggled**, **policy.simulation.completed**
- **lifecycle.\***: lifecycle.rule.created, lifecycle.rule.updated, lifecycle.rule.applied
- **legal_hold.\***: legal_hold.placed, legal_hold.removed
- **share.\***: share.link.created, share.link.accessed, share.link.revoked, share.link.updated, share.link.expired, **share.ownership.transferred**, **share.bulk_revoked**
- **access.\***: access.granted, access.revoked, access.updated
- **lock.\***: lock.acquired, lock.renewed, lock.released, lock.expired, lock.force_broken
- **quota.\***: quota.set, quota.adjusted, quota.threshold.crossed, quota.exceeded, **quota.plan.applied**, **quota.plan.changed**
- **gc.\***: gc.planned, gc.started, gc.completed, gc.failed, gc.scheduled, **gc.dry_run.completed**, **gc.cancelled**
- **encryption.\***: encryption.key.tracked, encryption.rotation.scheduled, encryption.key.rotated, encryption.key.retired, **encryption.key.revoked**, **encryption.rotation.paused**, **encryption.rotation.resumed**
- **audit.\***: audit.recorded, audit.export.created, **audit.streaming.enabled**
- **reference.\***: reference.added, reference.removed
- **folder.\***: folder.created, folder.renamed, folder.moved, folder.deleted, **folder.copied**
- **upload.\***: upload.initiated, upload.chunk.uploaded, upload.progress.updated, upload.completed, upload.aborted, upload.expired, **upload.resumed**
- **artifact.\***: artifact.creation.started, artifact.zip.created, artifact.preview.created, artifact.renewed, artifact.expired, **artifact.deleted**
- **extraction.\***: extraction.started, extraction.ocr.completed, extraction.exif.completed, extraction.text.completed, extraction.failed, **extraction.cancelled**, **extraction.retried**
- **flag.\***: file.flagged, flag.resolved, flag.dismissed, flag.escalated, **flag.bulk_actioned**
- **namespace.\***: namespace.created, namespace.updated, **namespace.disabled**, **namespace.migration.planned**, **namespace.migration.completed**
- **signed_url.\***: signed_url.generated, signed_url.batch.generated, signed_url.accessed, signed_url.revoked, signed_url.expired, **signed_url.audit.exported**

---

**END OF storage-be USER STORIES (Aligned Style)**