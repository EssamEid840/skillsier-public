# 📦 **admin-be - Admin & Moderation Service - Complete User Stories**

---

## **DOCUMENT OVERVIEW**

**Service:** admin-be  
**Purpose:** Manage admin operations, support ticketing, content moderation, user management, disputes, compliance, risk management, and platform configuration for the Skillsier platform  
**Architecture:** Event-Driven CQRS with Outbox Pattern  
**Event Envelope:** Standard platform envelope (event_id, timestamp, correlation_id, causation_id, user_context, compliance_context)  
**Idempotency:** All write commands use idempotency keys  
**Non-PII:** Events contain only IDs and codes; no direct PII  
**Coverage:** 100% of admin-be folder structure domains  
**Structure:** Sections → Domains → Entities → (Stories → Flow → Projections → Events → RBAC/SLO)

---

## **GLOBAL CONVENTIONS**

### Event Envelope Structure (All Events)
```json
{
  "event_id": "uuid",
  "event_type": "admin.user.created.v1",
  "event_timestamp": "2025-01-15T10:30:00Z",
  "event_version": "1",
  "aggregate_type": "admin_user",
  "aggregate_id": "uuid",
  "correlation_id": "uuid",
  "causation_id": "uuid",
  "event_source": "admin-be",
  "user_context": {
    "user_id": "uuid",
    "keycloak_id": "uuid",
    "user_type": "ADMIN",
    "session_id": "uuid",
    "admin_role": "SUPER_ADMIN|MODERATOR|SUPPORT|COMPLIANCE"
  },
  "compliance_context": {
    "gdpr_consent": true,
    "data_classification": "SENSITIVE|INTERNAL|PUBLIC",
    "retention_policy": "10y",
    "audit_required": true
  },
  "audit_metadata": {
    "ip_address": "x.x.x.x",
    "user_agent": "string",
    "request_id": "uuid",
    "action_category": "USER_ACTION|CONTENT_MODERATION|FINANCIAL|SYSTEM_CONFIG"
  },
  "payload": { /* event-specific data */ }
}
```

### Idempotent Write-Path
- All commands accept `idempotency_key` header
- Duplicate requests return cached response (24h TTL)
- Key format: `{service}.{command}.{admin_id}.{resource_id}.{hash}`

### Non-PII Event Rules
- Events contain only UUIDs, codes, enums, and numeric values
- No raw names, emails, addresses, or free-text PII
- Consumers fetch PII via API if needed
- Redacted fields marked as `[REDACTED]` in audits
- All admin actions fully audited with immutable logs

### Platform Alignment
- Folder structure: `/apps/be/admin-be/`
- Events catalog: `/apps/be/contracts/events/`
- Shared libraries: `/platform-shared/`, `/pkg/auth/`

### Path Defaults (All Domains)
- **Idempotency:** HTTP Header `Idempotency-Key` or envelope-level key; safe retries return 200 with no duplicate events
- **Transactions:** Database transaction + outbox pattern with (aggregate_id, event_type, idempotency_key) deduplication
- **Retries/DLQ:** For external service calls (users-be, jobs-be, proposals-be, contracts-be, financial-be, subscriptions-be, communications-be, storage-be, search-be, reviews-be); exponential backoff; poison messages to DLQ
- **Projections:** `_read` materialized views per domain; metric `event_to_projector_lag_ms` tracked
- **Security:** RBAC enforced on all commands/queries; field-level encryption for sensitive data; secrets never logged; all actions audited
- **Performance:** Typical write P95 ≤ 400ms, read P95 ≤ 300ms (unless specified otherwise)
- **Rate Limiting:** Per-endpoint rate limits enforced (e.g., 100 req/min/admin for moderation actions)
- **Audit:** All admin actions logged with full context; immutable audit trail; 10-year retention

### PII Handling (All Domains)
- **NO raw PII in events:** Emit hashes, storage_ids, references only
- **Examples:** No plaintext emails, phone numbers, file contents, file names
- **Compliance:** PII flags in envelope; anonymization tracked; data residency respected
- **Break-Glass Access:** PII access requires explicit approval and session tracking

### Integration Patterns
- **Async Event-Driven:** Primary integration via Kafka events
- **Sync Queries:** REST APIs for read operations with caching
- **Command Validation:** External service validation before persistence
- **Circuit Breakers:** Implemented for all external service calls
- **Saga Pattern:** Used for multi-service workflows (e.g., user suspension → content removal → notification)

---

## **SECTION 1: CORE ADMIN**

### 1.1 Domain: admin_user/

#### 1.1.1 Entity: entity.go (AdminUser Aggregate)

##### User Stories
- As a **super admin**, I want to **create admin accounts** so that staff can access admin tools.
- As a **super admin**, I want to **assign roles** (SuperAdmin/Moderator/Support/Compliance) so that permissions are appropriate.
- As a **super admin**, I want to **grant/revoke permissions** so that access control is granular.
- As an **admin**, I want to **enable MFA** on my account so that security is enhanced.
- As a **super admin**, I want to **deactivate admin accounts** so that former staff lose access.
- As an **admin**, I want to **view my activity log** so that I can track my actions.
- As a **super admin**, I want to **audit all admin actions** so that accountability is maintained.
- As a **system**, I want to **log all admin logins** so that access patterns are tracked.

##### Flow
1. **CreateAdminUserCommand**(email, roles[], permissions[], mfa_required, created_by, idempotency_key) → ValidateEmail() | ValidateRoles() | HashPassword() | Persist() → **Outbox:** admin.user.created.v1
2. **UpdateAdminUserCommand**(admin_id, updates, updated_by, idempotency_key) → AuthorizeSuperAdmin() | ValidateUpdates() | Apply() → **Outbox:** admin.user.updated.v1
3. **AssignRoleCommand**(admin_id, role, assigned_by, reason, idempotency_key) → AuthorizeSuperAdmin() | ValidateRole() | CheckConflicts() | Assign() → **Outbox:** admin.user.role.changed.v1
4. **GrantPermissionCommand**(admin_id, permission, granted_by, reason, idempotency_key) → AuthorizeSuperAdmin() | ValidatePermission() | Grant() → **Outbox:** admin.user.permissions.updated.v1
5. **RevokePermissionCommand**(admin_id, permission, revoked_by, reason, idempotency_key) → AuthorizeSuperAdmin() | Revoke() → **Outbox:** admin.user.permissions.updated.v1
6. **EnableMFACommand**(admin_id, mfa_type, secret) → Generate() | Enable() → **Outbox:** admin.user.mfa.enabled.v1
7. **DeactivateAdminCommand**(admin_id, deactivated_by, reason, idempotency_key) → AuthorizeSuperAdmin() | RevokeAllSessions() | Deactivate() → **Outbox:** admin.user.deactivated.v1
8. **LogAdminLoginCommand**(admin_id, ip, user_agent, session_id) → Record() → **Outbox:** admin.user.login.logged.v1
9. **GetAdminUserQuery**(admin_id) → AuthorizeOwnerOrSuperAdmin() | Fetch() → AdminUserDTO
10. **ListAdminUsersQuery**(filters, pagination) → AuthorizeSuperAdmin() | ApplyFilters() → AdminUserListDTO
11. **SearchAdminUsersQuery**(search_term, filters) → AuthorizeSuperAdmin() | Search() → AdminUserListDTO
12. **GetAdminActivityLogQuery**(admin_id, date_range, filters) → AuthorizeOwnerOrSuperAdmin() | Fetch() → ActivityLogDTO

##### Projections
- admin_user_read
- admin_roles_read
- admin_permissions_read
- admin_activity_log_read

##### Events Published
- admin.user.created.v1
- admin.user.updated.v1
- admin.user.role.changed.v1
- admin.user.permissions.updated.v1
- admin.user.mfa.enabled.v1
- admin.user.mfa.disabled.v1
- admin.user.deactivated.v1
- admin.user.reactivated.v1
- admin.user.login.logged.v1
- admin.user.logout.logged.v1
- admin.user.password.changed.v1

##### RBAC/SLO
- **RBAC:** SUPER_ADMIN (create/update/deactivate/assign roles/grant permissions), OWNER (view self/update self), ALL_ADMINS (view activity log)
- **SLO:** P95 < 200ms (create/update), P95 < 150ms (role assignment), P95 < 100ms (login logging), P95 < 150ms (read)

---

### 1.2 Domain: admin_user/ (continued)

#### 1.2.1 Entity: role.go (Role Model)

##### User Stories
- As a **super admin**, I want to **define role hierarchies** so that permissions inherit correctly.
- As a **super admin**, I want to **set role restrictions** so that conflicts are prevented.
- As a **system**, I want to **validate role assignments** so that invariants are maintained.

##### Flow
1. **DefineRoleCommand**(role_name, permissions[], hierarchy_level, restrictions, idempotency_key) → AuthorizeSuperAdmin() | ValidateHierarchy() | Persist() → **Outbox:** admin.role.defined.v1
2. **UpdateRoleCommand**(role_id, updates, updated_by, idempotency_key) → AuthorizeSuperAdmin() | Apply() → **Outbox:** admin.role.updated.v1
3. **ValidateRoleAssignmentCommand**(admin_id, role) → CheckConflicts() | CheckHierarchy() → ValidationResult
4. **GetRoleQuery**(role_id) → Fetch() → RoleDTO
5. **ListRolesQuery**() → FetchAll() → RoleListDTO

##### Projections
- role_definition_read
- role_hierarchy_read

##### Events Published
- admin.role.defined.v1
- admin.role.updated.v1
- admin.role.deleted.v1

##### RBAC/SLO
- **RBAC:** SUPER_ADMIN (define/update/delete), ALL_ADMINS (view)
- **SLO:** P95 < 150ms (define/update), P95 < 80ms (validate), P95 < 100ms (read)

---

#### 1.2.2 Entity: permission.go (Permission Bitset Helpers)

##### User Stories
- As a **super admin**, I want to **manage permission bitsets** so that access control is efficient.
- As a **system**, I want to **check permissions** on every operation so that unauthorized access is blocked.
- As a **super admin**, I want to **view effective permissions** for an admin so that access is clear.

##### Flow
1. **GrantPermissionsCommand**(admin_id, permissions[], granted_by, reason, idempotency_key) → AuthorizeSuperAdmin() | Grant() | UpdateBitset() → **Outbox:** admin.permissions.granted.v1
2. **RevokePermissionsCommand**(admin_id, permissions[], revoked_by, reason, idempotency_key) → AuthorizeSuperAdmin() | Revoke() | UpdateBitset() → **Outbox:** admin.permissions.revoked.v1
3. **CheckPermissionCommand**(admin_id, permission) → FetchBitset() | Check() → PermissionCheckResult
4. **GetEffectivePermissionsQuery**(admin_id) → FetchRoles() | ExpandPermissions() → EffectivePermissionsDTO

##### Projections
- permission_bitset_read
- effective_permissions_read

##### Events Published
- admin.permissions.granted.v1
- admin.permissions.revoked.v1
- admin.permissions.expanded.v1

##### RBAC/SLO
- **RBAC:** SUPER_ADMIN (grant/revoke), SYSTEM (check permissions - hot path)
- **SLO:** P95 < 30ms (check permission - critical path), P95 < 150ms (grant/revoke), P95 < 100ms (get effective)

---

#### 1.2.3 Entity: activity_log.go (Staff Action Trail)

##### User Stories
- As an **admin**, I want to **view my action history** so that I can track what I've done.
- As a **super admin**, I want to **audit all admin actions** so that compliance is maintained.
- As a **compliance officer**, I want to **export audit logs** so that regulatory requirements are met.
- As a **system**, I want to **append all actions** to immutable log so that tampering is prevented.
- As an **auditor**, I want to **search audit logs** by admin/action/date so that investigations are efficient.

##### Flow
1. **LogActionCommand**(admin_id, action_type, resource_type, resource_id, details, ip, user_agent, idempotency_key) → Persist() → **Outbox:** admin.action.logged.v1
2. **GetActivityLogQuery**(admin_id, date_range, filters) → AuthorizeOwnerOrSuperAdmin() | Fetch() → ActivityLogDTO
3. **SearchAuditLogsQuery**(search_params, filters, pagination) → AuthorizeAuditor() | Search() → AuditLogListDTO
4. **ExportAuditLogsCommand**(date_range, filters, format, exported_by) → AuthorizeCompliance() | GenerateExport() | Upload(storage-be) → **Outbox:** admin.audit.exported.v1
5. **GetActionDetailsQuery**(action_id) → AuthorizeAuditor() | Fetch() → ActionDetailsDTO

##### Projections
- activity_log_read
- audit_trail_read
- action_statistics_read

##### Events Published
- admin.action.logged.v1
- admin.audit.exported.v1
- admin.action.flagged.v1

##### RBAC/SLO
- **RBAC:** OWNER (view self), SUPER_ADMIN (view all), AUDITOR/COMPLIANCE (search/export)
- **SLO:** P95 < 50ms (log action - hot path), P95 < 200ms (search), P95 < 500ms (export)

---

## **SECTION 2: SUPPORT & CASEWORK**

### 2.1 Domain: support_ticket/

#### 2.1.1 Entity: entity.go (Ticket Aggregate)

##### User Stories
- As a **user**, I want to **open support tickets** so that I can get help with issues.
- As a **support agent**, I want to **view my assigned tickets** so that I can work on them.
- As a **support agent**, I want to **update ticket status** so that progress is tracked.
- As a **support agent**, I want to **assign tickets to other agents** so that workload is balanced.
- As a **support manager**, I want to **escalate tickets** so that complex issues get attention.
- As a **support agent**, I want to **see SLA clocks** so that deadlines are met.
- As a **system**, I want to **auto-escalate overdue tickets** so that SLAs are maintained.
- As a **user**, I want to **view my ticket history** so that I can track resolution.
- As a **support manager**, I want to **generate ticket reports** so that performance is measured.

##### Flow
1. **OpenTicketCommand**(subject, category, priority, description, user_id, attachments[], idempotency_key) → ValidateCategory() | DeterminePriority() | AssignQueue() | StartSLAClock() | Persist() | Notify(communications-be) → **Outbox:** admin.case.opened.v1
2. **UpdateTicketCommand**(ticket_id, updates, updated_by, idempotency_key) → AuthorizeAgentOrOwner() | Apply() | UpdateSLA() → **Outbox:** admin.case.updated.v1
3. **AssignTicketCommand**(ticket_id, agent_id, assigned_by, reason, idempotency_key) → AuthorizeManager() | ValidateAgent() | CheckWorkload() | Assign() | Notify(communications-be) → **Outbox:** admin.case.assigned.v1
4. **ReassignTicketCommand**(ticket_id, new_agent_id, reassigned_by, reason, idempotency_key) → AuthorizeManager() | Reassign() | UpdateSLA() → **Outbox:** admin.case.reassigned.v1
5. **EscalateTicketCommand**(ticket_id, escalation_level, escalated_by, reason, idempotency_key) → AuthorizeAgent() | EscalatePriority() | ReassignQueue() | Notify(communications-be) → **Outbox:** admin.case.escalated.v1
6. **ResolveTicketCommand**(ticket_id, resolution_notes, resolved_by, idempotency_key) → AuthorizeAgent() | Resolve() | StopSLA() | Notify(communications-be) → **Outbox:** admin.case.resolved.v1
7. **CloseTicketCommand**(ticket_id, closed_by, feedback_requested, idempotency_key) → AuthorizeAgent() | Close() | RequestFeedback(communications-be) → **Outbox:** admin.case.closed.v1
8. **ReopenTicketCommand**(ticket_id, reopened_by, reason, idempotency_key) → AuthorizeUser() | Reopen() | ResetSLA() → **Outbox:** admin.case.reopened.v1
9. **AutoEscalateTicketCommand**(ticket_id) → CheckSLABreach() | Escalate() | Notify(communications-be) → **Outbox:** admin.case.auto_escalated.v1
10. **GetTicketQuery**(ticket_id) → AuthorizeAgentOrOwner() | Fetch() → TicketDTO
11. **ListTicketsQuery**(filters, pagination) → AuthorizeAgent() | ApplyFilters() | FetchWithSLA() → TicketListDTO
12. **SearchTicketsQuery**(search_term, filters) → AuthorizeAgent() | Search() → TicketListDTO
13. **GetMyTicketsQuery**(agent_id, status_filter) → Fetch() → MyTicketsDTO
14. **GetTicketMetricsQuery**(date_range, agent_id) → AuthorizeManager() | Aggregate() → TicketMetricsDTO

##### Projections
- support_ticket_read
- ticket_queue_read
- ticket_sla_read
- agent_workload_read
- ticket_metrics_read

##### Events Published
- admin.case.opened.v1
- admin.case.updated.v1
- admin.case.assigned.v1
- admin.case.reassigned.v1
- admin.case.escalated.v1
- admin.case.auto_escalated.v1
- admin.case.resolved.v1
- admin.case.closed.v1
- admin.case.reopened.v1
- admin.case.sla.breached.v1
- admin.case.sla.warning.v1

##### Events Consumed
- user.support.request.created.v1 (from users-be - auto-create ticket)
- payment.disputed.v1 (from financial-be - create billing ticket)
- contract.dispute.opened.v1 (from contracts-be - create dispute ticket)

##### RBAC/SLO
- **RBAC:** USER (open/view own/reopen), AGENT (view assigned/update/resolve/close), MANAGER (assign/escalate/view all/metrics)
- **SLO:** P95 < 300ms (open), P95 < 200ms (update/assign/resolve/close), P95 < 150ms (read), P95 < 400ms (metrics)

---

### 2.2 Domain: support_ticket/ (continued)

#### 2.2.1 Entity: priority.go (Priority Enum + Escalation Thresholds)

##### User Stories
- As a **support manager**, I want to **define priority levels** so that urgency is clear.
- As a **system**, I want to **auto-assign priority** based on category so that triage is automated.
- As a **support manager**, I want to **set escalation thresholds** so that tickets escalate automatically.
- As a **support agent**, I want to **change priority** when needed so that urgency is adjusted.

##### Flow
1. **SetPriorityCommand**(ticket_id, priority, set_by, reason, idempotency_key) → AuthorizeAgent() | Validate() | Apply() | UpdateSLA() → **Outbox:** admin.case.priority.changed.v1
2. **DefinePriorityLevelCommand**(priority_name, sla_hours, escalation_threshold, auto_assign_rules, idempotency_key) → AuthorizeManager() | Persist() → **Outbox:** admin.priority.defined.v1
3. **AutoAssignPriorityCommand**(ticket_id, category) → ApplyRules() | Assign() → **Outbox:** admin.case.priority.auto_assigned.v1
4. **GetPriorityDefinitionsQuery**() → FetchAll() → PriorityListDTO

##### Projections
- priority_definition_read
- priority_thresholds_read

##### Events Published
- admin.case.priority.changed.v1
- admin.priority.defined.v1
- admin.case.priority.auto_assigned.v1

##### RBAC/SLO
- **RBAC:** AGENT (set priority), MANAGER (define priority levels)
- **SLO:** P95 < 100ms (set/auto-assign), P95 < 80ms (read)

---

#### 2.2.2 Entity: status.go (Status FSM)

##### User Stories
- As a **system**, I want to **enforce status transitions** so that invalid state changes are blocked.
- As a **support agent**, I want to **view valid next statuses** so that workflow is clear.
- As a **support manager**, I want to **define status workflows** so that processes are consistent.

##### Flow
1. **TransitionStatusCommand**(ticket_id, new_status, transitioned_by, notes, idempotency_key) → ValidateTransition() | Apply() | UpdateSLA() → **Outbox:** admin.case.status.changed.v1
2. **GetValidTransitionsQuery**(ticket_id, current_status) → FetchRules() → ValidTransitionsDTO
3. **DefineStatusWorkflowCommand**(workflow_rules, idempotency_key) → AuthorizeManager() | Persist() → **Outbox:** admin.workflow.defined.v1

##### Projections
- status_workflow_read

##### Events Published
- admin.case.status.changed.v1
- admin.workflow.defined.v1

##### RBAC/SLO
- **RBAC:** AGENT (transition status), MANAGER (define workflow)
- **SLO:** P95 < 150ms (transition), P95 < 80ms (get valid transitions)

---

#### 2.2.3 Entity: category.go (Ticket Categories)

##### User Stories
- As a **support manager**, I want to **define ticket categories** so that tickets are organized.
- As a **support agent**, I want to **categorize tickets** so that routing is correct.
- As a **system**, I want to **route tickets by category** so that specialists handle them.

##### Flow
1. **DefineCategoryCommand**(category_name, description, routing_queue, sla_override, idempotency_key) → AuthorizeManager() | Persist() → **Outbox:** admin.category.defined.v1
2. **UpdateCategoryCommand**(category_id, updates, updated_by, idempotency_key) → AuthorizeManager() | Apply() → **Outbox:** admin.category.updated.v1
3. **SetTicketCategoryCommand**(ticket_id, category, set_by, idempotency_key) → AuthorizeAgent() | Apply() | RouteToQueue() → **Outbox:** admin.case.category.changed.v1
4. **GetCategoriesQuery**() → FetchAll() → CategoryListDTO

##### Projections
- category_definition_read
- category_routing_read

##### Events Published
- admin.category.defined.v1
- admin.category.updated.v1
- admin.case.category.changed.v1

##### RBAC/SLO
- **RBAC:** AGENT (set category), MANAGER (define/update categories)
- **SLO:** P95 < 120ms (set/define/update), P95 < 80ms (read)

---

#### 2.2.4 Entity: assignment.go (Assignment VO)

##### User Stories
- As a **support manager**, I want to **define assignment rules** so that tickets are auto-assigned.
- As a **system**, I want to **balance workload** across agents so that no one is overloaded.
- As a **support agent**, I want to **see my queue** so that I know what to work on.
- As a **support manager**, I want to **override auto-assignment** when needed so that special cases are handled.

##### Flow
1. **DefineAssignmentRulesCommand**(rules, idempotency_key) → AuthorizeManager() | Validate() | Persist() → **Outbox:** admin.assignment.rules.defined.v1
2. **AutoAssignTicketCommand**(ticket_id) → ApplyRules() | CheckWorkload() | SelectAgent() | Assign() → **Outbox:** admin.case.auto_assigned.v1
3. **BalanceWorkloadCommand**() → CalculateLoad() | Redistribute() → **Outbox:** admin.workload.balanced.v1
4. **GetAgentQueueQuery**(agent_id) → Fetch() → AgentQueueDTO
5. **GetWorkloadQuery**(agent_id) → Calculate() → WorkloadDTO

##### Projections
- assignment_rules_read
- agent_queue_read
- workload_balance_read

##### Events Published
- admin.assignment.rules.defined.v1
- admin.case.auto_assigned.v1
- admin.workload.balanced.v1

##### RBAC/SLO
- **RBAC:** MANAGER (define rules/balance workload), AGENT (view queue)
- **SLO:** P95 < 200ms (auto-assign), P95 < 150ms (balance), P95 < 100ms (read)

---

### 2.3 Domain: ticket_message/

#### 2.3.1 Entity: entity.go (Ticket Message)

##### User Stories
- As a **support agent**, I want to **add notes to tickets** so that context is preserved.
- As a **support agent**, I want to **communicate with users** via ticket messages so that issues are resolved.
- As a **support agent**, I want to **mark messages as internal** so that users don't see them.
- As a **support agent**, I want to **edit my messages** so that corrections are possible.
- As a **support manager**, I want to **delete inappropriate messages** so that quality is maintained.
- As a **user**, I want to **reply to tickets** so that I can provide additional information.

##### Flow
1. **AddNoteCommand**(ticket_id, author_id, body, internal, attachments[], idempotency_key) → ValidateAuthor() | ValidateBody() | Persist() | Notify(communications-be if public) → **Outbox:** ticket.note.added.v1
2. **EditNoteCommand**(note_id, new_body, edited_by, idempotency_key) → AuthorizeAuthor() | ValidateEdit() | Apply() → **Outbox:** ticket.note.edited.v1
3. **DeleteNoteCommand**(note_id, deleted_by, reason, idempotency_key) → AuthorizeManagerOrAuthor() | SoftDelete() → **Outbox:** ticket.note.deleted.v1
4. **AddAttachmentCommand**(note_id, file_ref, file_name, file_size, uploaded_by, idempotency_key) → ValidateSize() | ScanVirus() | Persist() → **Outbox:** ticket.attachment.added.v1
5. **RemoveAttachmentCommand**(attachment_id, removed_by, reason, idempotency_key) → AuthorizeAgent() | Remove() → **Outbox:** ticket.attachment.removed.v1
6. **GetTicketMessagesQuery**(ticket_id, include_internal) → AuthorizeAgentOrOwner() | FilterInternal() | Fetch() → MessageListDTO
7. **GetMessageQuery**(message_id) → AuthorizeAgentOrOwner() | Fetch() → MessageDTO

##### Projections
- ticket_message_read
- ticket_conversation_read

##### Events Published
- ticket.note.added.v1
- ticket.note.edited.v1
- ticket.note.deleted.v1
- ticket.attachment.added.v1
- ticket.attachment.removed.v1

##### RBAC/SLO
- **RBAC:** AGENT (add/edit/delete notes/attachments), USER (add public notes/view public), MANAGER (delete any)
- **SLO:** P95 < 250ms (add note), P95 < 200ms (edit/delete), P95 < 150ms (read)

---

#### 2.3.2 Entity: attachment.go (Attachment)

##### User Stories
- As a **support agent**, I want to **attach files to tickets** so that documentation is complete.
- As a **system**, I want to **scan attachments for viruses** so that security is maintained.
- As a **support agent**, I want to **view attachment metadata** so that files are identifiable.
- As a **support manager**, I want to **remove malicious attachments** so that security is protected.

##### Flow
1. **UploadAttachmentCommand**(ticket_id, file_data, file_name, uploaded_by, idempotency_key) → ValidateSize() | Upload(storage-be) | ScanVirus() | Persist() → **Outbox:** ticket.attachment.uploaded.v1
2. **ScanAttachmentCommand**(attachment_id) → FetchFile(storage-be) | ScanVirus() | UpdateStatus() → **Outbox:** ticket.attachment.scanned.v1
3. **QuarantineAttachmentCommand**(attachment_id, quarantined_by, reason, idempotency_key) → AuthorizeManager() | Quarantine() | Notify(communications-be) → **Outbox:** ticket.attachment.quarantined.v1
4. **GetAttachmentQuery**(attachment_id) → AuthorizeAgentOrOwner() | Fetch() → AttachmentDTO
5. **DownloadAttachmentCommand**(attachment_id, downloaded_by) → AuthorizeAgentOrOwner() | GenerateURL(storage-be) | Log() → DownloadURLDTO

##### Projections
- attachment_read
- attachment_scan_status_read

##### Events Published
- ticket.attachment.uploaded.v1
- ticket.attachment.scanned.v1
- ticket.attachment.quarantined.v1
- ticket.attachment.downloaded.v1

##### Events Consumed
- storage.file.uploaded.v1 (complete attachment upload)
- storage.file.virus.detected.v1 (quarantine attachment)

##### RBAC/SLO
- **RBAC:** AGENT/USER (upload/download), MANAGER (quarantine)
- **SLO:** P95 < 300ms (upload), P95 < 500ms (scan), P95 < 150ms (download URL generation)

---

### 2.4 Domain: support_agent/

#### 2.4.1 Entity: entity.go (Agent Profile)

##### User Stories
- As a **support manager**, I want to **create agent profiles** so that capabilities are tracked.
- As a **support agent**, I want to **update my skills** so that assignment is accurate.
- As a **support manager**, I want to **assign agents to queues** so that routing works.
- As a **support agent**, I want to **set my availability** so that assignment reflects capacity.
- As a **support manager**, I want to **view agent KPIs** so that performance is measured.
- As a **system**, I want to **track agent workload** so that balance is maintained.

##### Flow
1. **CreateAgentProfileCommand**(admin_id, skills[], queues[], languages[], idempotency_key) → ValidateAgent() | Persist() → **Outbox:** support.agent.created.v1
2. **UpdateAgentProfileCommand**(agent_id, updates, updated_by, idempotency_key) → AuthorizeOwnerOrManager() | Apply() → **Outbox:** support.agent.updated.v1
3. **AddAgentSkillCommand**(agent_id, skill, added_by, idempotency_key) → AuthorizeManager() | Add() → **Outbox:** support.agent.skill.added.v1
4. **AssignAgentToQueueCommand**(agent_id, queue, assigned_by, idempotency_key) → AuthorizeManager() | Assign() → **Outbox:** support.agent.queue.assigned.v1
5. **SetAvailabilityCommand**(agent_id, status, reason) → Update() → **Outbox:** support.agent.status.changed.v1
6. **UpdateWorkloadCommand**(agent_id) → CalculateLoad() | Update() → **Outbox:** support.agent.workload.updated.v1
7. **GetAgentProfileQuery**(agent_id) → Fetch() → AgentProfileDTO
8. **GetAgentKPIsQuery**(agent_id, date_range) → Aggregate() → AgentKPIsDTO
9. **ListAgentsQuery**(filters, pagination) → AuthorizeManager() | ApplyFilters() → AgentListDTO

##### Projections
- agent_profile_read
- agent_skills_read
- agent_kpis_read
- agent_availability_read

##### Events Published
- support.agent.created.v1
- support.agent.updated.v1
- support.agent.skill.added.v1
- support.agent.skill.removed.v1
- support.agent.queue.assigned.v1
- support.agent.queue.unassigned.v1
- support.agent.status.changed.v1
- support.agent.workload.updated.v1

##### Events Consumed
- admin.case.assigned.v1 (update workload)
- admin.case.resolved.v1 (update KPIs)

##### RBAC/SLO
- **RBAC:** AGENT (update self/set availability), MANAGER (create/update/assign/view all)
- **SLO:** P95 < 200ms (create/update), P95 < 150ms (set availability), P95 < 300ms (KPIs), P95 < 150ms (read)

---

#### 2.4.2 Entity: availability.go (Presence State)

##### User Stories
- As a **support agent**, I want to **set my status** (Online/Busy/Offline) so that assignment reflects availability.
- As a **system**, I want to **auto-set offline** after timeout so that stale statuses are prevented.
- As a **support manager**, I want to **see team availability** so that coverage is managed.

##### Flow
1. **SetAgentStatusCommand**(agent_id, status, reason) → Validate() | Update() | RebalanceIfOffline() → **Outbox:** support.agent.availability.changed.v1
2. **AutoSetOfflineCommand**(agent_id) → CheckTimeout() | SetOffline() → **Outbox:** support.agent.auto_offline.v1
3. **GetTeamAvailabilityQuery**(team_id) → FetchAll() → TeamAvailabilityDTO

##### Projections
- agent_availability_read
- team_availability_read

##### Events Published
- support.agent.availability.changed.v1
- support.agent.auto_offline.v1

##### RBAC/SLO
- **RBAC:** AGENT (set status), MANAGER (view team)
- **SLO:** P95 < 100ms (set status), P95 < 150ms (read)

---

#### 2.4.3 Entity: stats.go (Rolling KPIs)

##### User Stories
- As a **support agent**, I want to **view my performance metrics** so that I can improve.
- As a **support manager**, I want to **compare agent metrics** so that coaching is targeted.
- As a **system**, I want to **calculate rolling KPIs** so that trends are visible.

##### Flow
1. **CalculateAgentKPIsCommand**(agent_id, date_range) → AggregateTickets() | CalculateFRT() | CalculateART() | CalculateCSAT() | CalculateResolutionRate() | Persist() → **Outbox:** support.agent.kpis.calculated.v1
2. **GetAgentStatsQuery**(agent_id, date_range) → Fetch() → AgentStatsDTO
3. **CompareAgentStatsQuery**(agent_ids[], date_range) → AuthorizeManager() | FetchMultiple() | Compare() → AgentComparisonDTO

##### Projections
- agent_kpis_read
- agent_stats_comparison_read

##### Events Published
- support.agent.kpis.calculated.v1
- support.agent.kpis.updated.v1

##### RBAC/SLO
- **RBAC:** AGENT (view self), MANAGER (view all/compare)
- **SLO:** P95 < 300ms (calculate), P95 < 200ms (read)

---

## **SECTION 3: CONTENT & KNOWLEDGE**

### 3.1 Domain: canned_response/

#### 3.1.1 Entity: entity.go (Prewritten Response)

##### User Stories
- As a **support manager**, I want to **create canned responses** so that agents respond consistently.
- As a **support agent**, I want to **use canned responses** so that replies are faster.
- As a **support manager**, I want to **organize responses by category** so that finding them is easy.
- As a **support manager**, I want to **localize responses** so that multiple languages are supported.
- As a **support manager**, I want to **track response usage** so that effectiveness is measured.

##### Flow
1. **CreateCannedResponseCommand**(title, body, category, locale, tags[], created_by, idempotency_key) → Validate() | Persist() → **Outbox:** canned_response.created.v1
2. **UpdateCannedResponseCommand**(response_id, updates, updated_by, idempotency_key) → AuthorizeManager() | Apply() → **Outbox:** canned_response.updated.v1
3. **ArchiveCannedResponseCommand**(response_id, archived_by, reason, idempotency_key) → AuthorizeManager() | Archive() → **Outbox:** canned_response.archived.v1
4. **UseCannedResponseCommand**(response_id, ticket_id, used_by) → Log() | IncrementUsage() → **Outbox:** canned_response.used.v1
5. **GetCannedResponseQuery**(response_id) → Fetch() → CannedResponseDTO
6. **SearchCannedResponsesQuery**(search_term, category, locale) → Search() → CannedResponseListDTO
7. **GetCannedResponseUsageQuery**(response_id, date_range) → AuthorizeManager() | Aggregate() → UsageStatsDTO

##### Projections
- canned_response_read
- canned_response_usage_read
- canned_response_category_read

##### Events Published
- canned_response.created.v1
- canned_response.updated.v1
- canned_response.archived.v1
- canned_response.unarchived.v1
- canned_response.used.v1

##### RBAC/SLO
- **RBAC:** MANAGER (create/update/archive), AGENT (view/search/use)
- **SLO:** P95 < 150ms (create/update/use), P95 < 100ms (search)

---

#### 3.1.2 Entity: category.go (Response Grouping)

##### User Stories
- As a **support manager**, I want to **organize responses into folders** so that categorization is clear.
- As a **support agent**, I want to **browse responses by category** so that finding them is easy.

##### Flow
1. **CreateResponseCategoryCommand**(name, parent_id, created_by, idempotency_key) → Validate() | Persist() → **Outbox:** canned_response.category.created.v1
2. **UpdateResponseCategoryCommand**(category_id, updates, updated_by, idempotency_key) → AuthorizeManager() | Apply() → **Outbox:** canned_response.category.updated.v1
3. **GetCategoryTreeQuery**() → FetchHierarchy() → CategoryTreeDTO

##### Projections
- response_category_read
- category_tree_read

##### Events Published
- canned_response.category.created.v1
- canned_response.category.updated.v1
- canned_response.category.deleted.v1

##### RBAC/SLO
- **RBAC:** MANAGER (create/update/delete), AGENT (view)
- **SLO:** P95 < 150ms (create/update), P95 < 100ms (read)

---

### 3.2 Domain: knowledge_base/

#### 3.2.1 Entity: entity.go (KB Article)

##### User Stories
- As a **support manager**, I want to **create KB articles** so that users can self-serve.
- As a **support agent**, I want to **link KB articles to tickets** so that users get relevant help.
- As a **user**, I want to **search the KB** so that I can find answers.
- As a **support manager**, I want to **version KB articles** so that history is preserved.
- As a **support manager**, I want to **publish/unpublish articles** so that drafts aren't visible.
- As a **system**, I want to **track article views** so that popularity is measured.

##### Flow
1. **CreateKBArticleCommand**(title, body, category, tags[], author_id, idempotency_key) → Validate() | Persist() | Index(search-be) → **Outbox:** kb.article.created.v1
2. **UpdateKBArticleCommand**(article_id, updates, updated_by, idempotency_key) → AuthorizeManager() | IncrementVersion() | Apply() | Reindex(search-be) → **Outbox:** kb.article.updated.v1
3. **PublishKBArticleCommand**(article_id, published_by, idempotency_key) → AuthorizeManager() | Publish() | Reindex(search-be) → **Outbox:** kb.article.published.v1
4. **UnpublishKBArticleCommand**(article_id, unpublished_by, reason, idempotency_key) → AuthorizeManager() | Unpublish() | Reindex(search-be) → **Outbox:** kb.article.unpublished.v1
5. **ViewKBArticleCommand**(article_id, viewer_id) → IncrementViews() | Log() → **Outbox:** kb.article.viewed.v1
6. **GetKBArticleQuery**(article_id, version) → Fetch() → KBArticleDTO
7. **SearchKBQuery**(search_term, filters) → Search() → KBArticleListDTO
8. **GetKBArticleVersionsQuery**(article_id) → FetchVersions() → VersionListDTO
9. **GetKBArticleStatsQuery**(article_id, date_range) → AuthorizeManager() | Aggregate() → ArticleStatsDTO

##### Projections
- kb_article_read
- kb_article_version_read
- kb_article_stats_read

##### Events Published
- kb.article.created.v1
- kb.article.updated.v1
- kb.article.published.v1
- kb.article.unpublished.v1
- kb.article.archived.v1
- kb.article.viewed.v1
- kb.article.version.created.v1

##### Events Consumed
- search.indexed.v1 (confirm indexing)

##### RBAC/SLO
- **RBAC:** MANAGER (create/update/publish), AGENT (view/search/link), USER (search/view published)
- **SLO:** P95 < 250ms (create/update/publish), P95 < 150ms (search), P95 < 100ms (view)

---

#### 3.2.2 Entity: category.go (KB Category)

##### User Stories
- As a **support manager**, I want to **organize KB articles into categories** so that browsing is easy.
- As a **user**, I want to **browse KB by category** so that topics are organized.

##### Flow
1. **CreateKBCategoryCommand**(name, parent_id, icon, order, created_by, idempotency_key) → Validate() | Persist() → **Outbox:** kb.category.created.v1
2. **UpdateKBCategoryCommand**(category_id, updates, updated_by, idempotency_key) → AuthorizeManager() | Apply() → **Outbox:** kb.category.updated.v1
3. **ReorderKBCategoriesCommand**(category_ids[], reordered_by, idempotency_key) → AuthorizeManager() | Reorder() → **Outbox:** kb.categories.reordered.v1
4. **GetKBCategoryTreeQuery**() → FetchHierarchy() → CategoryTreeDTO

##### Projections
- kb_category_read
- kb_category_tree_read

##### Events Published
- kb.category.created.v1
- kb.category.updated.v1
- kb.category.deleted.v1
- kb.categories.reordered.v1

##### RBAC/SLO
- **RBAC:** MANAGER (create/update/reorder/delete), ALL (view)
- **SLO:** P95 < 150ms (create/update/reorder), P95 < 100ms (read)

---

#### 3.2.3 Entity: related_article.go (Article Links)

##### User Stories
- As a **support manager**, I want to **link related articles** so that users discover more content.
- As a **user**, I want to **see related articles** so that I can explore topics.

##### Flow
1. **LinkRelatedArticlesCommand**(article_id, related_article_ids[], linked_by, idempotency_key) → AuthorizeManager() | Validate() | Link() → **Outbox:** kb.articles.linked.v1
2. **UnlinkRelatedArticleCommand**(article_id, related_article_id, unlinked_by, idempotency_key) → AuthorizeManager() | Unlink() → **Outbox:** kb.articles.unlinked.v1
3. **GetRelatedArticlesQuery**(article_id) → Fetch() → RelatedArticleListDTO

##### Projections
- related_articles_read

##### Events Published
- kb.articles.linked.v1
- kb.articles.unlinked.v1

##### RBAC/SLO
- **RBAC:** MANAGER (link/unlink), ALL (view)
- **SLO:** P95 < 150ms (link/unlink), P95 < 100ms (read)

---

#### 3.2.4 Entity: search_service.go (KB Search Adapter)

##### User Stories
- As a **user**, I want to **search KB articles** with full-text search so that relevant content is found.
- As a **support agent**, I want to **suggest KB articles** based on ticket content so that self-serve is encouraged.

##### Flow
1. **SearchKBCommand**(query, filters) → Query(search-be) | RankResults() → SearchResultsDTO
2. **SuggestArticlesCommand**(ticket_id) → ExtractKeywords() | Search(search-be) | RankBySimilarity() → SuggestionListDTO
3. **IndexArticleCommand**(article_id) → Fetch() | Index(search-be) → **Outbox:** kb.article.indexed.v1

##### Projections
- kb_search_index_read

##### Events Published
- kb.article.indexed.v1
- kb.article.reindexed.v1

##### Events Consumed
- search.indexed.v1 (confirm indexing)

##### RBAC/SLO
- **RBAC:** ALL (search), SYSTEM (index)
- **SLO:** P95 < 200ms (search), P95 < 300ms (suggest), P95 < 500ms (index)

---

### 3.3 Domain: faq/

#### 3.3.1 Entity: entity.go (FAQ)

##### User Stories
- As a **support manager**, I want to **create FAQs** so that common questions are answered.
- As a **user**, I want to **browse FAQs** so that quick answers are available.
- As a **support manager**, I want to **reorder FAQs** so that important ones are first.
- As a **support manager**, I want to **publish/unpublish FAQs** so that drafts aren't visible.

##### Flow
1. **CreateFAQCommand**(question, answer, category, locale, order, created_by, idempotency_key) → Validate() | Persist() → **Outbox:** faq.created.v1
2. **UpdateFAQCommand**(faq_id, updates, updated_by, idempotency_key) → AuthorizeManager() | Apply() → **Outbox:** faq.updated.v1
3. **PublishFAQCommand**(faq_id, published_by, idempotency_key) → AuthorizeManager() | Publish() → **Outbox:** faq.published.v1
4. **UnpublishFAQCommand**(faq_id, unpublished_by, reason, idempotency_key) → AuthorizeManager() | Unpublish() → **Outbox:** faq.unpublished.v1
5. **ReorderFAQsCommand**(faq_ids[], reordered_by, idempotency_key) → AuthorizeManager() | Reorder() → **Outbox:** faqs.reordered.v1
6. **GetFAQQuery**(faq_id) → Fetch() → FAQDTO
7. **ListFAQsQuery**(category, locale, published_only) → ApplyFilters() | Sort() → FAQListDTO

##### Projections
- faq_read
- faq_category_read

##### Events Published
- faq.created.v1
- faq.updated.v1
- faq.published.v1
- faq.unpublished.v1
- faq.archived.v1
- faqs.reordered.v1

##### RBAC/SLO
- **RBAC:** MANAGER (create/update/publish/reorder), ALL (view published)
- **SLO:** P95 < 150ms (create/update/publish/reorder), P95 < 100ms (read)

---

## **SECTION 4: SAFETY & MODERATION**

### 4.1 Domain: moderation/

#### 4.1.1 Entity: entity.go (Moderation Queue)

##### User Stories
- As a **moderator**, I want to **view flagged content** so that I can review it.
- As a **moderator**, I want to **approve/reject content** so that decisions are made.
- As a **moderator**, I want to **filter queue by severity** so that urgent items are prioritized.
- As a **system**, I want to **auto-flag content** based on rules so that moderation is efficient.
- As a **moderator**, I want to **escalate complex cases** so that senior review happens.
- As a **compliance officer**, I want to **audit moderation decisions** so that quality is maintained.

##### Flow
1. **FlagContentForModerationCommand**(content_type, content_id, flag_reason, severity, flagged_by, evidence, idempotency_key) → Validate() | Enqueue() | Notify(communications-be) → **Outbox:** admin.content.flagged.v1
2. **ApproveContentCommand**(queue_item_id, approved_by, notes, idempotency_key) → AuthorizeModerator() | Approve() | UpdateContent() → **Outbox:** admin.moderation.approved.v1
3. **RejectContentCommand**(queue_item_id, rejected_by, reason, action, idempotency_key) → AuthorizeModerator() | Reject() | TakeAction() | Notify(communications-be) → **Outbox:** admin.moderation.rejected.v1
4. **RemoveContentCommand**(queue_item_id, removed_by, reason, idempotency_key) → AuthorizeModerator() | Remove() | NotifyOwner(communications-be) → **Outbox:** admin.content.removed.v1
5. **EscalateModerationCommand**(queue_item_id, escalated_by, reason, idempotency_key) → AuthorizeModerator() | Escalate() | NotifyManager() → **Outbox:** admin.moderation.escalated.v1
6. **AutoFlagContentCommand**(content_type, content_id, rule_id, confidence_score) → ValidateRule() | Enqueue() → **Outbox:** admin.content.auto_flagged.v1
7. **GetModerationQueueQuery**(filters, pagination) → AuthorizeModerator() | ApplyFilters() | Sort() → QueueDTO
8. **GetQueueItemQuery**(item_id) → AuthorizeModerator() | Fetch() → QueueItemDTO
9. **GetModerationStatsQuery**(moderator_id, date_range) → AuthorizeManagerOrSelf() | Aggregate() → ModerationStatsDTO

##### Projections
- moderation_queue_read
- moderation_stats_read
- moderation_rules_read

##### Events Published
- admin.content.flagged.v1
- admin.content.auto_flagged.v1
- admin.moderation.approved.v1
- admin.moderation.rejected.v1
- admin.content.removed.v1
- admin.content.hidden.v1
- admin.moderation.escalated.v1

##### Events Consumed
- user.flagged.v1 (from users-be)
- job.flagged.v1 (from jobs-be)
- proposal.flagged.v1 (from proposals-be)
- review.flagged.v1 (from reviews-be)
- message.flagged.v1 (from communications-be)

##### RBAC/SLO
- **RBAC:** MODERATOR (view/approve/reject/remove), MANAGER (escalate/view all stats), SYSTEM (auto-flag)
- **SLO:** P95 < 250ms (flag/approve/reject/remove), P95 < 200ms (queue), P95 < 300ms (stats)

---

#### 4.1.2 Entity: queue_manager.go (Assignment & Aging Logic)

##### User Stories
- As a **system**, I want to **assign moderation items** to moderators so that workload is balanced.
- As a **system**, I want to **age items in queue** so that old items are prioritized.
- As a **moderation manager**, I want to **view queue health** so that bottlenecks are visible.

##### Flow
1. **AssignQueueItemCommand**(item_id, moderator_id, assigned_by) → CheckWorkload() | Assign() → **Outbox:** admin.moderation.assigned.v1
2. **AgeQueueItemsCommand**() → CalculateAge() | IncreasePriority() → **Outbox:** admin.moderation.items.aged.v1
3. **GetQueueHealthQuery**() → AuthorizeManager() | CalculateMetrics() → QueueHealthDTO

##### Projections
- queue_assignments_read
- queue_health_read

##### Events Published
- admin.moderation.assigned.v1
- admin.moderation.items.aged.v1

##### RBAC/SLO
- **RBAC:** SYSTEM (assign/age), MANAGER (view health)
- **SLO:** P95 < 200ms (assign), P95 < 500ms (age batch), P95 < 150ms (health)

---

#### 4.1.3 Entity: auto_moderator.go (Heuristics Rules)

##### User Stories
- As a **moderation manager**, I want to **define auto-moderation rules** so that obvious cases are handled automatically.
- As a **system**, I want to **apply rules to new content** so that flagging is automated.
- As a **moderation manager**, I want to **disable rules** so that false positives are reduced.

##### Flow
1. **CreateAutoModRuleCommand**(name, conditions, action, confidence_threshold, created_by, idempotency_key) → Validate() | Persist() → **Outbox:** admin.automod.rule.created.v1
2. **UpdateAutoModRuleCommand**(rule_id, updates, updated_by, idempotency_key) → AuthorizeManager() | Apply() → **Outbox:** admin.automod.rule.updated.v1
3. **EnableAutoModRuleCommand**(rule_id, enabled_by, idempotency_key) → AuthorizeManager() | Enable() → **Outbox:** admin.automod.rule.enabled.v1
4. **DisableAutoModRuleCommand**(rule_id, disabled_by, reason, idempotency_key) → AuthorizeManager() | Disable() → **Outbox:** admin.automod.rule.disabled.v1
5. **ApplyAutoModRulesCommand**(content_type, content_id, content_data) → FetchRules() | EvaluateEach() | FlagIfMatches() → **Outbox:** admin.automod.applied.v1
6. **GetAutoModRulesQuery**() → AuthorizeManager() | FetchAll() → RuleListDTO

##### Projections
- automod_rules_read
- automod_stats_read

##### Events Published
- admin.automod.rule.created.v1
- admin.automod.rule.updated.v1
- admin.automod.rule.enabled.v1
- admin.automod.rule.disabled.v1
- admin.automod.applied.v1

##### RBAC/SLO
- **RBAC:** MANAGER (create/update/enable/disable), SYSTEM (apply)
- **SLO:** P95 < 200ms (create/update), P95 < 150ms (apply - hot path)

---

#### 4.1.4 Entity: content_scanner.go (Lightweight Rule-Based Scanner Hooks)

##### User Stories
- As a **system**, I want to **scan content for profanity** so that inappropriate language is flagged.
- As a **system**, I want to **detect spam patterns** so that junk is filtered.
- As a **system**, I want to **check for PII exposure** so that privacy is protected.

##### Flow
1. **ScanContentCommand**(content_type, content_id, content_text) → ScanProfanity() | ScanSpam() | ScanPII() | FlagIfViolates() → **Outbox:** admin.content.scanned.v1
2. **GetScanResultQuery**(content_id) → Fetch() → ScanResultDTO

##### Projections
- content_scan_results_read

##### Events Published
- admin.content.scanned.v1
- admin.content.scan.violation.detected.v1

##### RBAC/SLO
- **RBAC:** SYSTEM (scan)
- **SLO:** P95 < 150ms (scan - hot path)

---

### 4.2 Domain: user_management/

#### 4.2.1 Entity: entity.go (User Admin Actions)

##### User Stories
- As a **moderator**, I want to **suspend user accounts** so that violations are addressed.
- As a **moderator**, I want to **ban user accounts** so that severe violations are prevented.
- As a **moderator**, I want to **warn users** so that behavior improves.
- As a **moderator**, I want to **verify user identities** so that trust is established.
- As a **moderator**, I want to **remove verification** so that fraudulent users are demoted.
- As a **system**, I want to **notify users of actions** so that transparency is maintained.
- As a **user**, I want to **appeal suspensions** so that mistakes can be corrected.

##### Flow
1. **SuspendUserCommand**(user_id, suspended_by, reason, duration, affected_content_action, idempotency_key) → AuthorizeModerator() | Suspend() | HideContent() | Notify(communications-be) | Publish(users-be) → **Outbox:** admin.user.suspended.v1
2. **UnsuspendUserCommand**(user_id, unsuspended_by, reason, idempotency_key) → AuthorizeModerator() | Unsuspend() | RestoreContent() | Notify(communications-be) | Publish(users-be) → **Outbox:** admin.user.unsuspended.v1
3. **BanUserCommand**(user_id, banned_by, reason, permanent, idempotency_key) → AuthorizeManager() | Ban() | RemoveContent() | CloseContracts(contracts-be) | Notify(communications-be) | Publish(users-be) → **Outbox:** admin.user.banned.v1
4. **UnbanUserCommand**(user_id, unbanned_by, reason, idempotency_key) → AuthorizeManager() | Unban() | Notify(communications-be) | Publish(users-be) → **Outbox:** admin.user.unbanned.v1
5. **WarnUserCommand**(user_id, warned_by, warning_type, message, idempotency_key) → AuthorizeModerator() | IssueWarning() | Notify(communications-be) | Publish(users-be) → **Outbox:** admin.user.warned.v1
6. **VerifyUserCommand**(user_id, verified_by, verification_type, evidence, idempotency_key) → AuthorizeModerator() | Verify() | Notify(communications-be) | Publish(users-be) → **Outbox:** admin.user.verified.v1
7. **RemoveVerificationCommand**(user_id, removed_by, reason, idempotency_key) → AuthorizeModerator() | RemoveVerification() | Notify(communications-be) | Publish(users-be) → **Outbox:** admin.user.verification.removed.v1
8. **GetUserAdminViewQuery**(user_id) → AuthorizeModerator() | FetchUserData(users-be) | FetchActions() → UserAdminViewDTO
9. **SearchUsersQuery**(search_term, filters) → AuthorizeModerator() | Search() → UserListDTO

##### Projections
- user_admin_actions_read
- user_admin_view_read

##### Events Published
- admin.user.suspended.v1
- admin.user.unsuspended.v1
- admin.user.banned.v1
- admin.user.unbanned.v1
- admin.user.warned.v1
- admin.user.verified.v1
- admin.user.verification.removed.v1

##### Events Consumed
- user.flagged.v1 (from users-be - trigger review)
- user.appeal.submitted.v1 (from users-be - review suspension)

##### RBAC/SLO
- **RBAC:** MODERATOR (suspend/warn/verify/remove verification), MANAGER (ban/unban/unsuspend major cases)
- **SLO:** P95 < 400ms (suspend/ban - saga), P95 < 300ms (warn/verify), P95 < 250ms (search)

---

#### 4.2.2 Entity: action_validator.go (Risk/Appeal/Ownership Guards)

##### User Stories
- As a **system**, I want to **validate action preconditions** so that invalid actions are blocked.
- As a **system**, I want to **check user ownership** so that cross-user actions are prevented.
- As a **system**, I want to **assess appeal rights** so that fair process is ensured.

##### Flow
1. **ValidateUserActionCommand**(user_id, action_type, action_params) → CheckRisk() | CheckOwnership() | CheckAppealRights() | CheckPreviousActions() → ValidationResult
2. **GetActionRisksQuery**(user_id, action_type) → AssessRisk() → RiskAssessmentDTO

##### Projections
- action_validation_read
- user_risk_assessment_read

##### Events Published
- admin.action.validated.v1
- admin.action.blocked.v1

##### RBAC/SLO
- **RBAC:** SYSTEM (validate)
- **SLO:** P95 < 150ms (validate - hot path)

---

### 4.3 Domain: content_management/

#### 4.3.1 Entity: entity.go (Content Admin Actions)

##### User Stories
- As a **moderator**, I want to **remove jobs** so that violations are addressed.
- As a **moderator**, I want to **hide proposals** so that inappropriate content is not visible.
- As a **moderator**, I want to **feature content** so that quality is promoted.
- As a **moderator**, I want to **remove reviews** so that fake reviews are eliminated.
- As a **moderator**, I want to **remove messages** so that abuse is prevented.
- As a **system**, I want to **notify content owners** so that actions are transparent.

##### Flow
1. **RemoveJobCommand**(job_id, removed_by, reason, notify_owner, idempotency_key) → AuthorizeModerator() | Remove() | NotifyOwner(communications-be) | Publish(jobs-be) → **Outbox:** admin.job.removed.v1
2. **HideJobCommand**(job_id, hidden_by, reason, idempotency_key) → AuthorizeModerator() | Hide() | Publish(jobs-be) → **Outbox:** admin.job.hidden.v1
3. **FeatureJobCommand**(job_id, featured_by, duration, idempotency_key) → AuthorizeModerator() | Feature() | Publish(jobs-be) → **Outbox:** admin.job.featured.v1
4. **RemoveProposalCommand**(proposal_id, removed_by, reason, notify_owner, idempotency_key) → AuthorizeModerator() | Remove() | NotifyOwner(communications-be) | Publish(proposals-be) → **Outbox:** admin.proposal.removed.v1
5. **RemoveReviewCommand**(review_id, removed_by, reason, idempotency_key) → AuthorizeModerator() | Remove() | RecalculateRatings(reviews-be) | Publish(reviews-be) → **Outbox:** admin.review.removed.v1
6. **RemoveMessageCommand**(message_id, removed_by, reason, idempotency_key) → AuthorizeModerator() | Remove() | Publish(communications-be) → **Outbox:** admin.message.removed.v1
7. **RemoveFileCommand**(file_id, removed_by, reason, idempotency_key) → AuthorizeModerator() | Remove() | Publish(storage-be) → **Outbox:** admin.file.removed.v1
8. **GetContentQuery**(content_type, content_id) → AuthorizeModerator() | Fetch() → ContentDTO
9. **SearchContentQuery**(content_type, search_term, filters) → AuthorizeModerator() | Search() → ContentListDTO

##### Projections
- content_admin_actions_read
- content_admin_view_read

##### Events Published
- admin.job.removed.v1
- admin.job.hidden.v1
- admin.job.unhidden.v1
- admin.job.featured.v1
- admin.proposal.removed.v1
- admin.proposal.hidden.v1
- admin.review.removed.v1
- admin.message.removed.v1
- admin.file.removed.v1

##### RBAC/SLO
- **RBAC:** MODERATOR (remove/hide/feature)
- **SLO:** P95 < 300ms (remove/hide/feature)

---

## **SECTION 5: DISPUTES & LEGAL**

### 5.1 Domain: dispute_resolution/

#### 5.1.1 Entity: entity.go (Dispute Lifecycle)

##### User Stories
- As a **user**, I want to **open disputes** so that conflicts are resolved.
- As a **moderator**, I want to **review disputes** so that fair decisions are made.
- As a **moderator**, I want to **request evidence** from parties so that decisions are informed.
- As a **moderator**, I want to **make decisions** on disputes so that resolution happens.
- As a **moderator**, I want to **escalate complex disputes** so that senior review occurs.
- As a **system**, I want to **track dispute SLAs** so that resolution is timely.
- As a **compliance officer**, I want to **audit dispute resolutions** so that fairness is verified.

##### Flow
1. **OpenDisputeCommand**(dispute_type, parties[], subject, description, evidence[], opened_by, idempotency_key) → Validate() | Enqueue() | StartSLA() | NotifyParties(communications-be) | Persist() → **Outbox:** admin.dispute.opened.v1
2. **AssignDisputeCommand**(dispute_id, moderator_id, assigned_by, idempotency_key) → AuthorizeManager() | Assign() | Notify(communications-be) → **Outbox:** admin.dispute.assigned.v1
3. **RequestEvidenceCommand**(dispute_id, party_id, evidence_type, deadline, requested_by, idempotency_key) → AuthorizeModerator() | Request() | Notify(communications-be) → **Outbox:** admin.dispute.evidence.requested.v1
4. **SubmitEvidenceCommand**(dispute_id, party_id, evidence[], submitted_by, idempotency_key) → ValidateParty() | Attach() | Notify(communications-be) → **Outbox:** admin.dispute.evidence.submitted.v1
5. **DecideDisputeCommand**(dispute_id, decision, reasoning, decided_by, actions[], idempotency_key) → AuthorizeModerator() | Decide() | ExecuteActions() | NotifyParties(communications-be) | CloseSLA() → **Outbox:** admin.dispute.resolved.v1
6. **EscalateDisputeCommand**(dispute_id, escalated_by, reason, idempotency_key) → AuthorizeModerator() | Escalate() | NotifyManager() → **Outbox:** admin.dispute.escalated.v1
7. **CloseDisputeCommand**(dispute_id, closed_by, outcome, idempotency_key) → AuthorizeModerator() | Close() | NotifyParties(communications-be) → **Outbox:** admin.dispute.closed.v1
8. **GetDisputeQuery**(dispute_id) → AuthorizeModeratorOrParty() | Fetch() → DisputeDTO
9. **ListDisputesQuery**(filters, pagination) → AuthorizeModerator() | ApplyFilters() → DisputeListDTO
10. **GetDisputeStatsQuery**(date_range) → AuthorizeManager() | Aggregate() → DisputeStatsDTO

##### Projections
- dispute_read
- dispute_queue_read
- dispute_evidence_read
- dispute_stats_read

##### Events Published
- admin.dispute.opened.v1
- admin.dispute.assigned.v1
- admin.dispute.evidence.requested.v1
- admin.dispute.evidence.submitted.v1
- admin.dispute.resolved.v1
- admin.dispute.escalated.v1
- admin.dispute.closed.v1
- admin.dispute.sla.breached.v1

##### Events Consumed
- contract.dispute.opened.v1 (from contracts-be)
- payment.disputed.v1 (from financial-be)

##### RBAC/SLO
- **RBAC:** USER (open/submit evidence), MODERATOR (assign/review/request evidence/decide/close), MANAGER (escalate/view stats)
- **SLO:** P95 < 300ms (open/assign/decide/close), P95 < 200ms (evidence operations), P95 < 150ms (read)

---

#### 5.1.2 Entity: evidence.go (Evidence Management)

##### User Stories
- As a **moderator**, I want to **organize evidence** so that review is structured.
- As a **moderator**, I want to **verify evidence authenticity** so that decisions are based on facts.
- As a **system**, I want to **archive evidence** so that records are preserved.

##### Flow
1. **OrganizeEvidenceCommand**(dispute_id, evidence_ids[], organization_scheme, organized_by, idempotency_key) → AuthorizeModerator() | Organize() → **Outbox:** admin.dispute.evidence.organized.v1
2. **VerifyEvidenceCommand**(evidence_id, verified_by, verification_notes, idempotency_key) → AuthorizeModerator() | Verify() → **Outbox:** admin.dispute.evidence.verified.v1
3. **ArchiveEvidenceCommand**(dispute_id, archived_by, idempotency_key) → AuthorizeModerator() | Archive(storage-be) → **Outbox:** admin.dispute.evidence.archived.v1

##### Projections
- evidence_read
- evidence_verification_read

##### Events Published
- admin.dispute.evidence.organized.v1
- admin.dispute.evidence.verified.v1
- admin.dispute.evidence.archived.v1

##### RBAC/SLO
- **RBAC:** MODERATOR (organize/verify/archive)
- **SLO:** P95 < 200ms (organize/verify), P95 < 400ms (archive)

---

#### 5.1.3 Entity: decision.go (Decision State Transitions)

##### User Stories
- As a **moderator**, I want to **record decision reasoning** so that justifications are documented.
- As a **system**, I want to **enforce decision finality** so that disputes can't be reopened arbitrarily.
- As a **compliance officer**, I want to **audit decisions** so that fairness is verified.

##### Flow
1. **RecordDecisionCommand**(dispute_id, decision_type, reasoning, evidence_refs[], decided_by, idempotency_key) → AuthorizeModerator() | Record() → **Outbox:** admin.dispute.decision.recorded.v1
2. **FinalizeDecisionCommand**(dispute_id, finalized_by, idempotency_key) → AuthorizeManager() | Finalize() → **Outbox:** admin.dispute.decision.finalized.v1
3. **GetDecisionHistoryQuery**(dispute_id) → AuthorizeModerator() | FetchHistory() → DecisionHistoryDTO

##### Projections
- decision_history_read

##### Events Published
- admin.dispute.decision.recorded.v1
- admin.dispute.decision.finalized.v1

##### RBAC/SLO
- **RBAC:** MODERATOR (record), MANAGER (finalize), COMPLIANCE (audit)
- **SLO:** P95 < 200ms (record/finalize), P95 < 150ms (read)

---

### 5.2 Domain: financial_dispute/

#### 5.2.1 Entity: entity.go (Financial Disputes)

##### User Stories
- As a **user**, I want to **dispute charges** so that incorrect charges are corrected.
- As a **financial moderator**, I want to **review payment disputes** so that refunds are processed.
- As a **financial moderator**, I want to **coordinate with financial-be** so that holds/refunds are executed.
- As a **system**, I want to **track chargeback cases** so that merchant accounts are protected.

##### Flow
1. **OpenPaymentDisputeCommand**(transaction_id, dispute_reason, evidence[], opened_by, idempotency_key) → Validate() | PlaceHold(financial-be) | Enqueue() | NotifyParties(communications-be) → **Outbox:** admin.payment.dispute.opened.v1
2. **ResolvePaymentDisputeCommand**(dispute_id, resolution, refund_amount, resolved_by, idempotency_key) → AuthorizeFinancialModerator() | ProcessRefund(financial-be) | ReleaseHold(financial-be) | Notify(communications-be) → **Outbox:** admin.payment.dispute.resolved.v1
3. **HandleChargebackCommand**(chargeback_id, response, evidence[], handled_by, idempotency_key) → AuthorizeFinancialManager() | SubmitResponse(financial-be) | Track() → **Outbox:** admin.chargeback.handled.v1
4. **GetPaymentDisputeQuery**(dispute_id) → AuthorizeFinancialModeratorOrParty() | Fetch() → PaymentDisputeDTO
5. **ListPaymentDisputesQuery**(filters, pagination) → AuthorizeFinancialModerator() | ApplyFilters() → PaymentDisputeListDTO

##### Projections
- payment_dispute_read
- chargeback_read

##### Events Published
- admin.payment.dispute.opened.v1
- admin.payment.dispute.resolved.v1
- admin.payment.refunded.v1
- admin.chargeback.opened.v1
- admin.chargeback.handled.v1
- admin.chargeback.resolved.v1

##### Events Consumed
- payment.disputed.v1 (from financial-be)
- payment.chargeback.received.v1 (from financial-be)

##### RBAC/SLO
- **RBAC:** FINANCIAL_MODERATOR (review/resolve), FINANCIAL_MANAGER (handle chargebacks)
- **SLO:** P95 < 400ms (open/resolve - saga with financial-be), P95 < 200ms (read)

---

### 5.3 Domain: legal_hold/

#### 5.3.1 Entity: entity.go (Legal Hold)

##### User Stories
- As a **legal officer**, I want to **place legal holds** so that data is preserved for litigation.
- As a **legal officer**, I want to **define hold scope** (user/content/contract) so that preservation is targeted.
- As a **legal officer**, I want to **release holds** when cases close so that data lifecycle resumes.
- As a **compliance officer**, I want to **audit holds** so that legal compliance is verified.

##### Flow
1. **PlaceLegalHoldCommand**(scope_type, scope_ids[], reason, case_number, placed_by, idempotency_key) → AuthorizeLegalOfficer() | PlaceHold() | NotifyServices() | FreezeData() → **Outbox:** admin.legal_hold.placed.v1
2. **ReleaseLegalHoldCommand**(hold_id, released_by, reason, idempotency_key) → AuthorizeLegalOfficer() | Release() | NotifyServices() | ResumeLifecycle() → **Outbox:** admin.legal_hold.released.v1
3. **CreateExportJobCommand**(hold_id, export_format, requester_id, idempotency_key) → AuthorizeLegalOfficer() | CreateJob() | GatherData() | Export(storage-be) → **Outbox:** admin.ediscovery.export.created.v1
4. **GetLegalHoldQuery**(hold_id) → AuthorizeLegalOfficer() | Fetch() → LegalHoldDTO
5. **ListLegalHoldsQuery**(filters, pagination) → AuthorizeLegalOfficer() | ApplyFilters() → LegalHoldListDTO

##### Projections
- legal_hold_read
- export_job_read

##### Events Published
- admin.legal_hold.placed.v1
- admin.legal_hold.released.v1
- admin.ediscovery.export.created.v1
- admin.ediscovery.export.completed.v1

##### RBAC/SLO
- **RBAC:** LEGAL_OFFICER (place/release/export), COMPLIANCE (audit)
- **SLO:** P95 < 500ms (place/release - saga), P95 < 2000ms (create export job), P95 < 200ms (read)

---

#### 5.3.2 Entity: export_job.go (Export Job Metadata)

##### User Stories
- As a **legal officer**, I want to **track export job progress** so that completion is monitored.
- As a **legal officer**, I want to **download export files** so that data is reviewed.

##### Flow
1. **TrackExportJobCommand**(job_id) → FetchStatus() | UpdateProgress() → **Outbox:** admin.ediscovery.export.progress.v1
2. **CompleteExportJobCommand**(job_id, file_refs[], completed_at) → MarkComplete() | Notify(communications-be) → **Outbox:** admin.ediscovery.export.completed.v1
3. **DownloadExportCommand**(job_id, downloaded_by) → AuthorizeLegalOfficer() | GenerateURL(storage-be) | Log() → DownloadURLDTO

##### Projections
- export_job_progress_read

##### Events Published
- admin.ediscovery.export.progress.v1
- admin.ediscovery.export.completed.v1
- admin.ediscovery.export.downloaded.v1

##### RBAC/SLO
- **RBAC:** LEGAL_OFFICER (track/download)
- **SLO:** P95 < 200ms (track), P95 < 300ms (download)

---

### 5.4 Domain: privacy_request/

#### 5.4.1 Entity: entity.go (GDPR/CCPA DSAR)

##### User Stories
- As a **user**, I want to **request data access** (DSAR) so that I can see my data.
- As a **user**, I want to **request data erasure** so that my data is deleted.
- As a **compliance officer**, I want to **verify identity** before fulfilling requests so that fraud is prevented.
- As a **compliance officer**, I want to **fulfill requests within SLA** so that legal compliance is maintained.
- As a **system**, I want to **coordinate erasure across services** so that deletion is complete.

##### Flow
1. **SubmitPrivacyRequestCommand**(request_type, user_id, evidence[], submitted_by, idempotency_key) → Validate() | Enqueue() | StartSLA() | Notify(communications-be) → **Outbox:** admin.privacy.requested.v1
2. **VerifyPrivacyRequestCommand**(request_id, verified_by, verification_notes, idempotency_key) → AuthorizeCompliance() | VerifyIdentity() | Approve() → **Outbox:** admin.privacy.approved.v1
3. **FulfillDataAccessCommand**(request_id, fulfilled_by, idempotency_key) → AuthorizeCompliance() | GatherData() | Export(storage-be) | Notify(communications-be) → **Outbox:** admin.privacy.fulfilled.v1
4. **FulfillDataErasureCommand**(request_id, fulfilled_by, idempotency_key) → AuthorizeCompliance() | EraseData(users-be, jobs-be, proposals-be, contracts-be, financial-be) | Verify() | Notify(communications-be) → **Outbox:** admin.privacy.fulfilled.v1
5. **DenyPrivacyRequestCommand**(request_id, denied_by, reason, idempotency_key) → AuthorizeCompliance() | Deny() | Notify(communications-be) → **Outbox:** admin.privacy.denied.v1
6. **GetPrivacyRequestQuery**(request_id) → AuthorizeComplianceOrOwner() | Fetch() → PrivacyRequestDTO
7. **ListPrivacyRequestsQuery**(filters, pagination) → AuthorizeCompliance() | ApplyFilters() → PrivacyRequestListDTO

##### Projections
- privacy_request_read
- privacy_request_sla_read

##### Events Published
- admin.privacy.requested.v1
- admin.privacy.approved.v1
- admin.privacy.fulfilled.v1
- admin.privacy.denied.v1
- admin.privacy.erasure.completed.v1

##### Events Consumed
- user.privacy.request.submitted.v1 (from users-be)

##### RBAC/SLO
- **RBAC:** USER (submit), COMPLIANCE (verify/fulfill/deny)
- **SLO:** P95 < 300ms (submit/verify/deny), P95 < 5000ms (fulfill - multi-service saga), 30 days total SLA

---

#### 5.4.2 Entity: evidence.go (Identity Proof & Consent Trail)

##### User Stories
- As a **compliance officer**, I want to **collect identity proof** so that requests are legitimate.
- As a **system**, I want to **track consent** so that GDPR compliance is maintained.

##### Flow
1. **CollectIdentityProofCommand**(request_id, proof_type, proof_data, collected_by, idempotency_key) → AuthorizeCompliance() | Validate() | Store(storage-be) → **Outbox:** admin.privacy.evidence.collected.v1
2. **VerifyIdentityProofCommand**(request_id, verified_by, verification_result, idempotency_key) → AuthorizeCompliance() | Verify() → **Outbox:** admin.privacy.identity.verified.v1

##### Projections
- identity_proof_read

##### Events Published
- admin.privacy.evidence.collected.v1
- admin.privacy.identity.verified.v1

##### RBAC/SLO
- **RBAC:** COMPLIANCE (collect/verify)
- **SLO:** P95 < 250ms (collect/verify)

---

### 5.5 Domain: pii_access/

#### 5.5.1 Entity: entity.go (Break-Glass PII Access)

##### User Stories
- As an **admin**, I want to **request PII access** with justification so that break-glass is audited.
- As a **super admin**, I want to **approve PII requests** so that access is controlled.
- As a **system**, I want to **grant temporary PII access** so that sessions are time-limited.
- As a **compliance officer**, I want to **audit PII access** so that misuse is detected.
- As a **system**, I want to **auto-expire PII grants** so that access doesn't persist.

##### Flow
1. **RequestPIIAccessCommand**(purpose, scope_type, scope_ids[], duration, requested_by, justification, idempotency_key) → Validate() | Enqueue() | NotifyApprovers() → **Outbox:** admin.pii.access.requested.v1
2. **ApprovePIIAccessCommand**(request_id, approved_by, conditions[], idempotency_key) → AuthorizeSuperAdmin() | Grant() | StartSession() → **Outbox:** admin.pii.access.granted.v1
3. **DenyPIIAccessCommand**(request_id, denied_by, reason, idempotency_key) → AuthorizeSuperAdmin() | Deny() | Notify(communications-be) → **Outbox:** admin.pii.access.denied.v1
4. **AccessPIICommand**(grant_id, accessed_by, resource_id, idempotency_key) → ValidateGrant() | Log() | UnmaskData() → **Outbox:** admin.pii.accessed.v1
5. **RevokeAccessCommand**(grant_id, revoked_by, reason, idempotency_key) → AuthorizeSuperAdmin() | Revoke() | Notify(communications-be) → **Outbox:** admin.pii.access.revoked.v1
6. **ExpireAccessCommand**(grant_id) → CheckExpiry() | Expire() → **Outbox:** admin.pii.access.expired.v1
7. **GetPIIAccessRequestsQuery**(filters, pagination) → AuthorizeSuperAdmin() | ApplyFilters() → PIIAccessRequestListDTO
8. **GetPIIAccessAuditQuery**(date_range, admin_id) → AuthorizeCompliance() | FetchAudit() → PIIAccessAuditDTO

##### Projections
- pii_access_request_read
- pii_access_grant_read
- pii_access_audit_read

##### Events Published
- admin.pii.access.requested.v1
- admin.pii.access.granted.v1
- admin.pii.access.denied.v1
- admin.pii.accessed.v1
- admin.pii.access.revoked.v1
- admin.pii.access.expired.v1

##### RBAC/SLO
- **RBAC:** ADMIN (request), SUPER_ADMIN (approve/deny/revoke), COMPLIANCE (audit)
- **SLO:** P95 < 200ms (request/approve/deny/revoke), P95 < 150ms (access - hot path), P95 < 300ms (audit)

---

#### 5.5.2 Entity: grant.go (Grant Management)

##### User Stories
- As a **system**, I want to **track active grants** so that access is monitored.
- As a **system**, I want to **enforce grant expiry** so that access is time-bound.

##### Flow
1. **CreateGrantCommand**(request_id, scope, duration, conditions[], created_by) → Create() | Schedule Expiry() → **Outbox:** admin.pii.grant.created.v1
2. **GetActiveGrantsQuery**(admin_id) → Fetch() → ActiveGrantsDTO

##### Projections
- active_grants_read

##### Events Published
- admin.pii.grant.created.v1
- admin.pii.grant.expired.v1

##### RBAC/SLO
- **RBAC:** SYSTEM (create/expire)
- **SLO:** P95 < 150ms (create), P95 < 100ms (read)

---

#### 5.5.3 Entity: policy.go (Masking Policy)

##### User Stories
- As a **compliance officer**, I want to **define masking policies** so that PII is protected by default.
- As a **system**, I want to **apply masking rules** so that PII is never exposed without grants.

##### Flow
1. **DefineMaskingPolicyCommand**(policy_name, fields[], redaction_rules[], idempotency_key) → AuthorizeCompliance() | Validate() | Persist() → **Outbox:** admin.pii.policy.defined.v1
2. **ApplyMaskingCommand**(data, policy_id) → FetchPolicy() | ApplyRules() | RedactPII() → MaskedData
3. **GetMaskingPoliciesQuery**() → AuthorizeCompliance() | FetchAll() → PolicyListDTO

##### Projections
- masking_policy_read

##### Events Published
- admin.pii.policy.defined.v1
- admin.pii.policy.updated.v1

##### RBAC/SLO
- **RBAC:** COMPLIANCE (define/update), SYSTEM (apply)
- **SLO:** P95 < 150ms (define/update), P95 < 80ms (apply - hot path)

---

### 5.6 Domain: business_verification/

#### 5.6.1 Entity: entity.go (Business Profile Verification)

##### User Stories
- As a **business client**, I want to **submit business verification** so that I can hire at scale.
- As a **compliance officer**, I want to **review business documents** so that legitimacy is verified.
- As a **compliance officer**, I want to **approve/reject verification** so that only legitimate businesses proceed.
- As a **system**, I want to **track verification status** so that hiring limits are enforced.

##### Flow
1. **SubmitBusinessVerificationCommand**(user_id, business_name, registration_number, tax_id, documents[], submitted_by, idempotency_key) → Validate() | Enqueue() | Notify(communications-be) → **Outbox:** admin.biz_verification.requested.v1
2. **ReviewBusinessVerificationCommand**(verification_id, reviewed_by, notes, idempotency_key) → AuthorizeCompliance() | Review() → **Outbox:** admin.biz_verification.reviewed.v1
3. **ApproveBusinessVerificationCommand**(verification_id, approved_by, verification_level, idempotency_key) → AuthorizeCompliance() | Approve() | UpdateUserProfile(users-be) | Notify(communications-be) → **Outbox:** admin.biz_verification.approved.v1
4. **RejectBusinessVerificationCommand**(verification_id, rejected_by, reason, idempotency_key) → AuthorizeCompliance() | Reject() | Notify(communications-be) → **Outbox:** admin.biz_verification.rejected.v1
5. **GetBusinessVerificationQuery**(verification_id) → AuthorizeComplianceOrOwner() | Fetch() → BusinessVerificationDTO
6. **ListBusinessVerificationsQuery**(filters, pagination) → AuthorizeCompliance() | ApplyFilters() → BusinessVerificationListDTO

##### Projections
- business_verification_read
- business_verification_queue_read

##### Events Published
- admin.biz_verification.requested.v1
- admin.biz_verification.reviewed.v1
- admin.biz_verification.approved.v1
- admin.biz_verification.rejected.v1

##### Events Consumed
- user.business.verification.submitted.v1 (from users-be)

##### RBAC/SLO
- **RBAC:** USER (submit), COMPLIANCE (review/approve/reject)
- **SLO:** P95 < 300ms (submit/review/approve/reject), P95 < 200ms (read)

---

### 5.7 Domain: sanctions_screening/

#### 5.7.1 Entity: entity.go (Watchlist Screening)

##### User Stories
- As a **compliance system**, I want to **screen users against sanctions lists** so that prohibited entities are blocked.
- As a **compliance officer**, I want to **review screening hits** so that false positives are cleared.
- As a **compliance officer**, I want to **escalate true hits** so that accounts are blocked.
- As a **system**, I want to **re-screen periodically** so that changes are detected.

##### Flow
1. **RunSanctionsScreeningCommand**(user_id, screening_lists[], idempotency_key) → FetchUserData(users-be) | ScreenAgainstLists() | PersistResults() → **Outbox:** admin.sanctions.screening.completed.v1
2. **ReviewSanctionsHitCommand**(hit_id, reviewed_by, disposition, notes, idempotency_key) → AuthorizeCompliance() | Review() | UpdateDisposition() → **Outbox:** admin.sanctions.hit.reviewed.v1
3. **ClearSanctionsHitCommand**(hit_id, cleared_by, reason, idempotency_key) → AuthorizeCompliance() | Clear() | Notify(communications-be) → **Outbox:** admin.sanctions.hit.cleared.v1
4. **EscalateSanctionsHitCommand**(hit_id, escalated_by, reason, idempotency_key) → AuthorizeCompliance() | Escalate() | BlockUser(users-be) | Notify(communications-be) → **Outbox:** admin.sanctions.hit.escalated.v1
5. **GetSanctionsScreeningQuery**(screening_id) → AuthorizeCompliance() | Fetch() → ScreeningDTO
6. **ListSanctionsHitsQuery**(filters, pagination) → AuthorizeCompliance() | ApplyFilters() → SanctionsHitListDTO

##### Projections
- sanctions_screening_read
- sanctions_hit_read

##### Events Published
- admin.sanctions.screening.completed.v1
- admin.sanctions.hit.detected.v1
- admin.sanctions.hit.reviewed.v1
- admin.sanctions.hit.cleared.v1
- admin.sanctions.hit.escalated.v1

##### Events Consumed
- user.created.v1 (from users-be - auto-screen new users)

##### RBAC/SLO
- **RBAC:** SYSTEM (run screening), COMPLIANCE (review/clear/escalate)
- **SLO:** P95 < 500ms (run screening), P95 < 250ms (review/clear/escalate), P95 < 200ms (read)

---

## **SECTION 6: REPORTING & METRICS**

### 6.1 Domain: reporting/

#### 6.1.1 Entity: entity.go (Admin Reports)

##### User Stories
- As an **admin manager**, I want to **generate performance reports** so that team productivity is measured.
- As an **admin manager**, I want to **export reports to CSV/PDF** so that sharing is easy.
- As an **admin manager**, I want to **schedule automated reports** so that stakeholders are updated regularly.
- As an **executive**, I want to **view platform health dashboards** so that status is visible.

##### Flow
1. **GenerateReportCommand**(report_type, date_range, filters, generated_by, idempotency_key) → AuthorizeManager() | AggregateData() | GenerateReport() | Upload(storage-be) → **Outbox:** admin.report.generated.v1
2. **ExportReportCommand**(report_id, export_format, exported_by, idempotency_key) → AuthorizeManager() | FetchData() | ConvertFormat() | Upload(storage-be) → **Outbox:** admin.report.exported.v1
3. **ScheduleReportCommand**(report_config, frequency, recipients[], scheduled_by, idempotency_key) → AuthorizeManager() | Schedule() → **Outbox:** admin.report.scheduled.v1
4. **GetReportQuery**(report_id) → AuthorizeManager() | Fetch() → ReportDTO
5. **ListReportsQuery**(filters, pagination) → AuthorizeManager() | ApplyFilters() → ReportListDTO

##### Projections
- report_read
- report_schedule_read

##### Events Published
- admin.report.generated.v1
- admin.report.exported.v1
- admin.report.scheduled.v1
- admin.report.sent.v1

##### RBAC/SLO
- **RBAC:** MANAGER (generate/export/schedule), EXECUTIVE (view dashboards)
- **SLO:** P95 < 5000ms (generate - large aggregation), P95 < 2000ms (export), P95 < 200ms (read)

---

### 6.2 Domain: metrics/

#### 6.2.1 Entity: entity.go (Platform Metrics)

##### User Stories
- As an **admin manager**, I want to **track ticket resolution metrics** so that performance is measured.
- As an **admin manager**, I want to **track moderation metrics** so that queue health is monitored.
- As an **executive**, I want to **view aggregate platform metrics** so that health is assessed.
- As a **system**, I want to **calculate metrics in real-time** so that dashboards are current.

##### Flow
1. **CalculateMetricsCommand**(metric_type, date_range) → AggregateData() | Calculate() | Persist() → **Outbox:** admin.metrics.calculated.v1
2. **GetTicketMetricsQuery**(date_range, filters) → AuthorizeManager() | Aggregate() → TicketMetricsDTO
3. **GetModerationMetricsQuery**(date_range, filters) → AuthorizeManager() | Aggregate() → ModerationMetricsDTO
4. **GetPlatformHealthQuery**() → AuthorizeExecutive() | FetchMetrics() → PlatformHealthDTO

##### Projections
- ticket_metrics_read
- moderation_metrics_read
- platform_health_read

##### Events Published
- admin.metrics.calculated.v1
- admin.metrics.updated.v1

##### RBAC/SLO
- **RBAC:** MANAGER (view team metrics), EXECUTIVE (view platform metrics)
- **SLO:** P95 < 300ms (calculate), P95 < 200ms (read)

---

## **SECTION 7: CONFIGURATION & POLICY**

### 7.1 Domain: system_config/

#### 7.1.1 Entity: entity.go (System Configuration)

##### User Stories
- As an **admin manager**, I want to **manage feature flags** so that features are controlled.
- As an **admin manager**, I want to **configure platform settings** so that behavior is customized.
- As an **admin manager**, I want to **version configuration changes** so that rollback is possible.
- As a **system**, I want to **cache config** so that performance is optimized.

##### Flow
1. **UpdateConfigCommand**(config_key, config_value, updated_by, reason, idempotency_key) → AuthorizeManager() | Validate() | VersionIncrement() | Apply() | InvalidateCache() → **Outbox:** admin.config.updated.v1
2. **CreateFeatureFlagCommand**(flag_name, description, default_value, created_by, idempotency_key) → AuthorizeManager() | Persist() → **Outbox:** admin.feature_flag.created.v1
3. **UpdateFeatureFlagCommand**(flag_id, enabled, updated_by, idempotency_key) → AuthorizeManager() | Apply() | InvalidateCache() | NotifyServices() → **Outbox:** admin.feature_flag.updated.v1
4. **RollbackConfigCommand**(config_key, target_version, rolled_back_by, idempotency_key) → AuthorizeManager() | Rollback() | InvalidateCache() → **Outbox:** admin.config.rolled_back.v1
5. **GetConfigQuery**(config_key) → Fetch() → ConfigDTO
6. **GetConfigHistoryQuery**(config_key) → AuthorizeManager() | FetchVersions() → ConfigHistoryDTO
7. **GetFeatureFlagsQuery**() → FetchAll() → FeatureFlagListDTO

##### Projections
- system_config_read
- config_history_read
- feature_flag_read

##### Events Published
- admin.config.updated.v1
- admin.config.rolled_back.v1
- admin.feature_flag.created.v1
- admin.feature_flag.updated.v1
- admin.feature_flag.deleted.v1

##### RBAC/SLO
- **RBAC:** MANAGER (update/rollback), ALL_ADMINS (view)
- **SLO:** P95 < 200ms (update/rollback), P95 < 80ms (read - hot path with caching)

---

### 7.2 Domain: policy_doc/

#### 7.2.1 Entity: entity.go (Policy Documents)

##### User Stories
- As a **compliance officer**, I want to **publish policy documents** (Terms, Privacy Policy) so that users are informed.
- As a **compliance officer**, I want to **version policies** so that history is preserved.
- As a **compliance officer**, I want to **require acceptance** on policy updates so that consent is tracked.
- As a **system**, I want to **notify users of policy changes** so that awareness is ensured.

##### Flow
1. **CreatePolicyDocumentCommand**(policy_type, title, content, effective_date, created_by, idempotency_key) → Validate() | Persist() → **Outbox:** admin.policy.created.v1
2. **UpdatePolicyDocumentCommand**(policy_id, updates, updated_by, idempotency_key) → AuthorizeCompliance() | IncrementVersion() | Apply() → **Outbox:** admin.policy.updated.v1
3. **PublishPolicyCommand**(policy_id, published_by, require_acceptance, idempotency_key) → AuthorizeCompliance() | Publish() | NotifyUsers(communications-be) → **Outbox:** admin.policy.published.v1
4. **GetPolicyDocumentQuery**(policy_id, version) → Fetch() → PolicyDocumentDTO
5. **GetLatestPolicyQuery**(policy_type) → FetchLatest() → PolicyDocumentDTO
6. **GetPolicyVersionsQuery**(policy_id) → FetchVersions() → PolicyVersionListDTO

##### Projections
- policy_doc_read
- policy_version_read

##### Events Published
- admin.policy.created.v1
- admin.policy.updated.v1
- admin.policy.published.v1
- admin.policy.version.created.v1

##### RBAC/SLO
- **RBAC:** COMPLIANCE (create/update/publish), ALL (view)
- **SLO:** P95 < 250ms (create/update/publish), P95 < 100ms (read)

---

### 7.3 Domain: experiment/

#### 7.3.1 Entity: entity.go (A/B Tests & Experiments)

##### User Stories
- As a **product manager**, I want to **create experiments** so that features are tested.
- As a **product manager**, I want to **define experiment variants** so that options are compared.
- As a **product manager**, I want to **ramp experiments** so that rollout is controlled.
- As a **product manager**, I want to **analyze experiment results** so that decisions are data-driven.

##### Flow
1. **CreateExperimentCommand**(name, description, variants[], created_by, idempotency_key) → Validate() | Persist() → **Outbox:** admin.experiment.created.v1
2. **UpdateExperimentCommand**(experiment_id, updates, updated_by, idempotency_key) → AuthorizeProductManager() | Apply() → **Outbox:** admin.experiment.updated.v1
3. **RampExperimentCommand**(experiment_id, target_percentage, ramped_by, idempotency_key) → AuthorizeProductManager() | Ramp() → **Outbox:** admin.experiment.ramped.v1
4. **CompleteExperimentCommand**(experiment_id, winning_variant, completed_by, idempotency_key) → AuthorizeProductManager() | Complete() → **Outbox:** admin.experiment.completed.v1
5. **GetExperimentQuery**(experiment_id) → Fetch() → ExperimentDTO
6. **GetExperimentResultsQuery**(experiment_id) → Aggregate() → ExperimentResultsDTO

##### Projections
- experiment_read
- experiment_results_read

##### Events Published
- admin.experiment.created.v1
- admin.experiment.updated.v1
- admin.experiment.ramped.v1
- admin.experiment.completed.v1

##### RBAC/SLO
- **RBAC:** PRODUCT_MANAGER (create/update/ramp/complete), ALL_ADMINS (view)
- **SLO:** P95 < 200ms (create/update/ramp/complete), P95 < 300ms (results)

---

### 7.4 Domain: search_policy_admin/

#### 7.4.1 Entity: entity.go (Search Policy Bundles)

##### User Stories
- As an **admin manager**, I want to **configure search policies** so that ranking is controlled.
- As an **admin manager**, I want to **adjust search weights** so that relevance is tuned.
- As an **admin manager**, I want to **manage search filters** so that results are appropriate.

##### Flow
1. **UpdateSearchPolicyCommand**(policy_bundle, updated_by, idempotency_key) → AuthorizeManager() | Validate() | Apply() | Publish(search-be) → **Outbox:** admin.search.policy.updated.v1
2. **GetSearchPolicyQuery**() → Fetch() → SearchPolicyDTO

##### Projections
- search_policy_read

##### Events Published
- admin.search.policy.updated.v1

##### RBAC/SLO
- **RBAC:** MANAGER (update), ALL_ADMINS (view)
- **SLO:** P95 < 200ms (update), P95 < 100ms (read)

---

### 7.5 Domain: throttle_policy/

#### 7.5.1 Entity: entity.go (Throttle Policies)

##### User Stories
- As an **admin manager**, I want to **define rate limits** so that abuse is prevented.
- As an **admin manager**, I want to **create exceptions** for trusted users so that legitimate traffic isn't blocked.
- As a **system**, I want to **enforce rate limits** so that resources are protected.

##### Flow
1. **CreateThrottlePolicyCommand**(resource_type, limit, window, created_by, idempotency_key) → Validate() | Persist() → **Outbox:** admin.throttle.policy.created.v1
2. **UpdateThrottlePolicyCommand**(policy_id, updates, updated_by, idempotency_key) → AuthorizeManager() | Apply() | InvalidateCache() → **Outbox:** admin.throttle.policy.updated.v1
3. **CreateExceptionCommand**(user_id, resource_type, custom_limit, reason, created_by, idempotency_key) → AuthorizeManager() | Create() → **Outbox:** admin.throttle.exception.created.v1
4. **GetThrottlePoliciesQuery**() → FetchAll() → ThrottlePolicyListDTO

##### Projections
- throttle_policy_read
- throttle_exception_read

##### Events Published
- admin.throttle.policy.created.v1
- admin.throttle.policy.updated.v1
- admin.throttle.exception.created.v1
- admin.throttle.exception.deleted.v1

##### RBAC/SLO
- **RBAC:** MANAGER (create/update/exception)
- **SLO:** P95 < 150ms (create/update), P95 < 80ms (read)

---

### 7.6 Domain: quota_override/

#### 7.6.1 Entity: entity.go (Quota Overrides)

##### User Stories
- As an **admin manager**, I want to **override user quotas** so that special cases are handled.
- As an **admin manager**, I want to **set expiry on overrides** so that temporary exceptions are automatic.
- As a **system**, I want to **expire overrides** so that normal limits resume.

##### Flow
1. **CreateQuotaOverrideCommand**(user_id, quota_type, override_value, expires_at, reason, created_by, idempotency_key) → AuthorizeManager() | Validate() | Create() | Notify(subscriptions-be) → **Outbox:** admin.quota.override.created.v1
2. **UpdateQuotaOverrideCommand**(override_id, updates, updated_by, idempotency_key) → AuthorizeManager() | Apply() → **Outbox:** admin.quota.override.updated.v1
3. **ExpireQuotaOverrideCommand**(override_id) → CheckExpiry() | Expire() | Notify(subscriptions-be) → **Outbox:** admin.quota.override.expired.v1
4. **GetQuotaOverridesQuery**(user_id) → FetchAll() → QuotaOverrideListDTO

##### Projections
- quota_override_read

##### Events Published
- admin.quota.override.created.v1
- admin.quota.override.updated.v1
- admin.quota.override.expired.v1
- admin.quota.override.deleted.v1

##### RBAC/SLO
- **RBAC:** MANAGER (create/update/delete)
- **SLO:** P95 < 200ms (create/update/expire), P95 < 100ms (read)

---

## **SECTION 8: RISK, FRAUD & INCIDENTS**

### 8.1 Domain: fraud_review/

#### 8.1.1 Entity: entity.go (Fraud Cases)

##### User Stories
- As a **fraud analyst**, I want to **review fraud alerts** so that risky accounts are investigated.
- As a **fraud analyst**, I want to **approve/reject users** after review so that decisions are made.
- As a **fraud analyst**, I want to **place holds** on suspicious accounts so that funds are protected.
- As a **system**, I want to **consume fraud signals** from other services so that cases are created.

##### Flow
1. **CreateFraudCaseCommand**(user_id, case_type, risk_score, evidence[], created_by, idempotency_key) → Validate() | Enqueue() | PlaceHold(financial-be if high risk) → **Outbox:** admin.fraud.case.created.v1
2. **AssignFraudCaseCommand**(case_id, analyst_id, assigned_by, idempotency_key) → AuthorizeManager() | Assign() → **Outbox:** admin.fraud.case.assigned.v1
3. **ReviewFraudCaseCommand**(case_id, reviewed_by, findings, idempotency_key) → AuthorizeAnalyst() | Review() → **Outbox:** admin.fraud.case.reviewed.v1
4. **ApproveFraudCaseCommand**(case_id, approved_by, notes, idempotency_key) → AuthorizeAnalyst() | Approve() | ReleaseHold(financial-be) | Notify(communications-be) → **Outbox:** admin.fraud.case.approved.v1
5. **RejectFraudCaseCommand**(case_id, rejected_by, reason, action, idempotency_key) → AuthorizeAnalyst() | Reject() | TakeAction(ban/suspend) | Notify(communications-be) → **Outbox:** admin.fraud.case.rejected.v1
6. **GetFraudCaseQuery**(case_id) → AuthorizeAnalyst() | Fetch() → FraudCaseDTO
7. **ListFraudCasesQuery**(filters, pagination) → AuthorizeAnalyst() | ApplyFilters() → FraudCaseListDTO

##### Projections
- fraud_case_read
- fraud_queue_read

##### Events Published
- admin.fraud.case.created.v1
- admin.fraud.case.assigned.v1
- admin.fraud.case.reviewed.v1
- admin.fraud.case.approved.v1
- admin.fraud.case.rejected.v1

##### Events Consumed
- fraud.alert.triggered.v1 (from financial-be)
- risk.alert.triggered.v1 (from financial-be)

##### RBAC/SLO
- **RBAC:** FRAUD_ANALYST (review/approve/reject), FRAUD_MANAGER (assign/view all)
- **SLO:** P95 < 300ms (create/assign/review/approve/reject), P95 < 200ms (read)

---

### 8.2 Domain: risk_hold/

#### 8.2.1 Entity: entity.go (Risk Holds)

##### User Stories
- As a **fraud analyst**, I want to **place financial holds** so that risky transactions are frozen.
- As a **fraud analyst**, I want to **release holds** when cleared so that funds flow.
- As a **system**, I want to **track hold duration** so that limits are enforced.

##### Flow
1. **PlaceRiskHoldCommand**(user_id, hold_type, amount, reason, placed_by, idempotency_key) → AuthorizeAnalyst() | PlaceHold(financial-be) | Persist() | Notify(communications-be) → **Outbox:** admin.risk.hold.placed.v1
2. **ReleaseRiskHoldCommand**(hold_id, released_by, reason, idempotency_key) → AuthorizeAnalyst() | ReleaseHold(financial-be) | Notify(communications-be) → **Outbox:** admin.risk.hold.released.v1
3. **GetRiskHoldsQuery**(user_id) → AuthorizeAnalyst() | Fetch() → RiskHoldListDTO

##### Projections
- risk_hold_read

##### Events Published
- admin.risk.hold.placed.v1
- admin.risk.hold.released.v1

##### RBAC/SLO
- **RBAC:** FRAUD_ANALYST (place/release)
- **SLO:** P95 < 300ms (place/release - saga with financial-be), P95 < 150ms (read)

---

### 8.3 Domain: risk_reserve/

#### 8.3.1 Entity: entity.go (Risk Reserves)

##### User Stories
- As a **fraud analyst**, I want to **set risk reserves** so that high-risk payments have buffers.
- As a **system**, I want to **track reserve duration** so that releases are timely.

##### Flow
1. **SetRiskReserveCommand**(user_id, percentage, duration, reason, set_by, idempotency_key) → AuthorizeAnalyst() | SetReserve(financial-be) | Persist() → **Outbox:** admin.risk.reserve.set.v1
2. **ReleaseRiskReserveCommand**(reserve_id, released_by, idempotency_key) → AuthorizeAnalyst() | ReleaseReserve(financial-be) → **Outbox:** admin.risk.reserve.released.v1

##### Projections
- risk_reserve_read

##### Events Published
- admin.risk.reserve.set.v1
- admin.risk.reserve.released.v1

##### RBAC/SLO
- **RBAC:** FRAUD_ANALYST (set/release)
- **SLO:** P95 < 300ms (set/release)

---

### 8.4 Domain: risk_chargeback/

#### 8.4.1 Entity: entity.go (Chargeback Cases)

##### User Stories
- As a **fraud analyst**, I want to **track chargebacks** so that patterns are identified.
- As a **fraud analyst**, I want to **respond to chargebacks** so that disputes are defended.
- As a **system**, I want to **link chargebacks to fraud cases** so that investigations are complete.

##### Flow
1. **CreateChargebackCaseCommand**(transaction_id, chargeback_reason, amount, created_by, idempotency_key) → Validate() | Link(financial-be) | Persist() → **Outbox:** admin.chargeback.case.created.v1
2. **RespondToChargebackCommand**(case_id, response, evidence[], responded_by, idempotency_key) → AuthorizeAnalyst() | SubmitResponse(financial-be) → **Outbox:** admin.chargeback.responded.v1
3. **ResolveChargebackCommand**(case_id, outcome, resolved_by, idempotency_key) → AuthorizeAnalyst() | Resolve() | UpdateMetrics() → **Outbox:** admin.chargeback.resolved.v1

##### Projections
- chargeback_case_read

##### Events Published
- admin.chargeback.case.created.v1
- admin.chargeback.responded.v1
- admin.chargeback.resolved.v1

##### Events Consumed
- payment.chargeback.received.v1 (from financial-be)

##### RBAC/SLO
- **RBAC:** FRAUD_ANALYST (respond/resolve)
- **SLO:** P95 < 300ms (create/respond/resolve)

---

### 8.5 Domain: risk_velocity_alert/

#### 8.5.1 Entity: entity.go (Velocity Alerts)

##### User Stories
- As a **system**, I want to **detect velocity anomalies** so that suspicious patterns are flagged.
- As a **fraud analyst**, I want to **review velocity alerts** so that rapid activity is investigated.

##### Flow
1. **CreateVelocityAlertCommand**(user_id, alert_type, threshold_exceeded, details, idempotency_key) → Validate() | Persist() | Notify(communications-be) → **Outbox:** admin.velocity.alert.created.v1
2. **ReviewVelocityAlertCommand**(alert_id, reviewed_by, disposition, idempotency_key) → AuthorizeAnalyst() | Review() → **Outbox:** admin.velocity.alert.reviewed.v1

##### Projections
- velocity_alert_read

##### Events Published
- admin.velocity.alert.created.v1
- admin.velocity.alert.reviewed.v1

##### RBAC/SLO
- **RBAC:** SYSTEM (create), FRAUD_ANALYST (review)
- **SLO:** P95 < 200ms (create), P95 < 250ms (review)

---

### 8.6 Domain: risk_country_rate_anomaly/

#### 8.6.1 Entity: entity.go (Geo/Rate Anomalies)

##### User Stories
- As a **system**, I want to **detect geographic anomalies** so that unusual patterns are flagged.
- As a **fraud analyst**, I want to **review anomalies** so that fraudulent activity is investigated.

##### Flow
1. **CreateAnomalyCommand**(user_id, anomaly_type, details, risk_score, idempotency_key) → Validate() | Persist() | Notify(communications-be) → **Outbox:** admin.anomaly.created.v1
2. **ReviewAnomalyCommand**(anomaly_id, reviewed_by, disposition, idempotency_key) → AuthorizeAnalyst() | Review() → **Outbox:** admin.anomaly.reviewed.v1

##### Projections
- anomaly_read

##### Events Published
- admin.anomaly.created.v1
- admin.anomaly.reviewed.v1

##### RBAC/SLO
- **RBAC:** SYSTEM (create), FRAUD_ANALYST (review)
- **SLO:** P95 < 200ms (create/review)

---

### 8.7 Domain: incident/

#### 8.7.1 Entity: entity.go (Ops Incidents & Postmortems)

##### User Stories
- As an **ops engineer**, I want to **open incidents** so that outages are tracked.
- As an **ops engineer**, I want to **update incident status** so that progress is visible.
- As an **incident commander**, I want to **assign roles** so that response is coordinated.
- As an **ops engineer**, I want to **record timeline events** so that chronology is preserved.
- As an **ops manager**, I want to **conduct postmortems** so that learning happens.
- As an **ops manager**, I want to **track action items** so that improvements are implemented.

##### Flow
1. **OpenIncidentCommand**(title, severity, description, detected_by, idempotency_key) → Validate() | Persist() | NotifyTeam(communications-be) → **Outbox:** admin.incident.opened.v1
2. **UpdateIncidentCommand**(incident_id, updates, updated_by, idempotency_key) → AuthorizeOps() | Apply() → **Outbox:** admin.incident.updated.v1
3. **AssignCommanderCommand**(incident_id, commander_id, assigned_by, idempotency_key) → AuthorizeOpsManager() | Assign() | Notify(communications-be) → **Outbox:** admin.incident.commander.assigned.v1
4. **AddTimelineEventCommand**(incident_id, event_type, description, added_by, idempotency_key) → AuthorizeOps() | Append() → **Outbox:** admin.incident.timeline.added.v1
5. **ResolveIncidentCommand**(incident_id, resolution, resolved_by, idempotency_key) → AuthorizeCommander() | Resolve() | NotifyTeam(communications-be) → **Outbox:** admin.incident.resolved.v1
6. **CreatePostmortemCommand**(incident_id, rca, action_items[], created_by, idempotency_key) → AuthorizeOpsManager() | Create() → **Outbox:** admin.incident.postmortem.created.v1
7. **TrackActionItemCommand**(action_item_id, status, assignee, due_date, idempotency_key) → AuthorizeOps() | Update() → **Outbox:** admin.incident.action_item.updated.v1
8. **GetIncidentQuery**(incident_id) → AuthorizeOps() | Fetch() → IncidentDTO
9. **ListIncidentsQuery**(filters, pagination) → AuthorizeOps() | ApplyFilters() → IncidentListDTO
10. **GetIncidentMetricsQuery**(date_range) → AuthorizeOpsManager() | Aggregate() → IncidentMetricsDTO

##### Projections
- incident_read
- incident_timeline_read
- incident_postmortem_read
- incident_metrics_read

##### Events Published
- admin.incident.opened.v1
- admin.incident.updated.v1
- admin.incident.commander.assigned.v1
- admin.incident.timeline.added.v1
- admin.incident.resolved.v1
- admin.incident.postmortem.created.v1
- admin.incident.action_item.created.v1
- admin.incident.action_item.updated.v1
- admin.incident.action_item.completed.v1

##### RBAC/SLO
- **RBAC:** OPS (open/update/add timeline), COMMANDER (resolve), OPS_MANAGER (assign commander/create postmortem/metrics)
- **SLO:** P95 < 250ms (open/update/resolve), P95 < 150ms (add timeline/track action item), P95 < 200ms (read)

---

## **SECTION 9: INTEGRATIONS & BULK OPS**

### 9.1 Domain: integrations_admin/

#### 9.1.1 Entity: entity.go (Third-Party Integrations)

##### User Stories
- As an **admin manager**, I want to **add integrations** so that third-party services connect.
- As an **admin manager**, I want to **configure integration endpoints** so that data flows.
- As an **admin manager**, I want to **disable integrations** so that unused services are disconnected.
- As an **admin manager**, I want to **monitor integration health** so that failures are detected.

##### Flow
1. **AddIntegrationCommand**(vendor, integration_type, endpoints[], scopes[], added_by, idempotency_key) → Validate() | Persist() | Test() → **Outbox:** admin.integration.added.v1
2. **UpdateIntegrationCommand**(integration_id, updates, updated_by, idempotency_key) → AuthorizeManager() | Apply() → **Outbox:** admin.integration.updated.v1
3. **DisableIntegrationCommand**(integration_id, disabled_by, reason, idempotency_key) → AuthorizeManager() | Disable() → **Outbox:** admin.integration.disabled.v1
4. **TestIntegrationCommand**(integration_id, tested_by) → SendTestRequest() | ValidateResponse() → IntegrationTestResultDTO
5. **GetIntegrationQuery**(integration_id) → Fetch() → IntegrationDTO
6. **ListIntegrationsQuery**(filters) → ApplyFilters() → IntegrationListDTO

##### Projections
- integration_read
- integration_health_read

##### Events Published
- admin.integration.added.v1
- admin.integration.updated.v1
- admin.integration.disabled.v1
- admin.integration.enabled.v1
- admin.integration.test.completed.v1

##### RBAC/SLO
- **RBAC:** MANAGER (add/update/disable/test)
- **SLO:** P95 < 300ms (add/update/disable), P95 < 500ms (test), P95 < 150ms (read)

---

#### 9.1.2 Entity: api_key.go (API Keys)

##### User Stories
- As an **admin manager**, I want to **issue API keys** so that integrations authenticate.
- As an **admin manager**, I want to **revoke API keys** so that compromised keys are disabled.
- As an **admin manager**, I want to **rotate API keys** so that security is maintained.
- As a **system**, I want to **track key usage** so that anomalies are detected.

##### Flow
1. **IssueAPIKeyCommand**(integration_id, scopes[], expiry, issued_by, idempotency_key) → GenerateKey() | HashKey() | Persist() → **Outbox:** admin.api_key.issued.v1
2. **RevokeAPIKeyCommand**(key_id, revoked_by, reason, idempotency_key) → AuthorizeManager() | Revoke() → **Outbox:** admin.api_key.revoked.v1
3. **RotateAPIKeyCommand**(key_id, rotated_by, idempotency_key) → AuthorizeManager() | GenerateNew() | RevokeOld() → **Outbox:** admin.api_key.rotated.v1
4. **TrackKeyUsageCommand**(key_id, request_count, last_used_at) → Update() → **Outbox:** admin.api_key.usage.tracked.v1
5. **GetAPIKeysQuery**(integration_id) → AuthorizeManager() | Fetch() → APIKeyListDTO

##### Projections
- api_key_read
- api_key_usage_read

##### Events Published
- admin.api_key.issued.v1
- admin.api_key.revoked.v1
- admin.api_key.rotated.v1
- admin.api_key.usage.tracked.v1

##### RBAC/SLO
- **RBAC:** MANAGER (issue/revoke/rotate)
- **SLO:** P95 < 200ms (issue/revoke/rotate), P95 < 100ms (track usage)

---

#### 9.1.3 Entity: webhook_endpoint.go (Webhook Endpoints)

##### User Stories
- As an **admin manager**, I want to **configure webhook endpoints** so that events are sent to integrations.
- As an **admin manager**, I want to **test webhooks** so that delivery is verified.
- As a **system**, I want to **retry failed webhooks** so that delivery is reliable.

##### Flow
1. **AddWebhookEndpointCommand**(integration_id, url, secret, events[], added_by, idempotency_key) → Validate() | Persist() → **Outbox:** admin.webhook.endpoint.added.v1
2. **UpdateWebhookEndpointCommand**(endpoint_id, updates, updated_by, idempotency_key) → AuthorizeManager() | Apply() → **Outbox:** admin.webhook.endpoint.updated.v1
3. **TestWebhookCommand**(endpoint_id, tested_by) → SendTestEvent() | ValidateResponse() → WebhookTestResultDTO
4. **GetWebhookEndpointsQuery**(integration_id) → Fetch() → WebhookEndpointListDTO

##### Projections
- webhook_endpoint_read

##### Events Published
- admin.webhook.endpoint.added.v1
- admin.webhook.endpoint.updated.v1
- admin.webhook.endpoint.deleted.v1
- admin.webhook.test.completed.v1

##### RBAC/SLO
- **RBAC:** MANAGER (add/update/delete/test)
- **SLO:** P95 < 250ms (add/update/delete), P95 < 500ms (test), P95 < 150ms (read)

---

### 9.2 Domain: bulk_action/

#### 9.2.1 Entity: entity.go (Bulk Operations)

##### User Stories
- As an **admin manager**, I want to **preview bulk operations** so that impact is understood before execution.
- As an **admin manager**, I want to **execute bulk operations** so that mass changes are efficient.
- As an **admin manager**, I want to **rollback bulk operations** so that mistakes are corrected.
- As a **system**, I want to **track bulk job progress** so that status is visible.

##### Flow
1. **StartBulkJobCommand**(operation_type, target_query, parameters, preview_mode, started_by, idempotency_key) → Validate() | EnqueueJob() → **Outbox:** admin.bulk.started.v1
2. **PreviewBulkJobCommand**(job_id) → ExecuteQuery() | GeneratePreview() → BulkJobPreviewDTO
3. **CommitBulkJobCommand**(job_id, committed_by, idempotency_key) → AuthorizeManager() | ValidatePreview() | ExecuteBatches() | TrackProgress() → **Outbox:** admin.bulk.committed.v1
4. **RollbackBulkJobCommand**(job_id, rolled_back_by, reason, idempotency_key) → AuthorizeManager() | RevertChanges() → **Outbox:** admin.bulk.rolled_back.v1
5. **GetBulkJobQuery**(job_id) → Fetch() → BulkJobDTO
6. **GetBulkJobProgressQuery**(job_id) → FetchProgress() → BulkJobProgressDTO
7. **ListBulkJobsQuery**(filters, pagination) → AuthorizeManager() | ApplyFilters() → BulkJobListDTO

##### Projections
- bulk_job_read
- bulk_job_progress_read

##### Events Published
- admin.bulk.started.v1
- admin.bulk.preview.generated.v1
- admin.bulk.committed.v1
- admin.bulk.progressed.v1
- admin.bulk.completed.v1
- admin.bulk.rolled_back.v1
- admin.bulk.failed.v1

##### RBAC/SLO
- **RBAC:** MANAGER (start/preview/commit/rollback)
- **SLO:** P95 < 300ms (start/preview), P95 < 5000ms (commit - depends on batch size), P95 < 3000ms (rollback), P95 < 200ms (read)

---

## **SECTION 10: SESSIONS & APPROVALS**

### 10.1 Domain: admin_session/

#### 10.1.1 Entity: entity.go (JIT / Break-Glass Sessions)

##### User Stories
- As an **admin**, I want to **request elevated access** with justification so that sensitive operations are controlled.
- As a **super admin**, I want to **approve access requests** so that elevation is authorized.
- As a **system**, I want to **auto-expire sessions** so that temporary access doesn't persist.
- As a **compliance officer**, I want to **audit all elevated sessions** so that misuse is detected.

##### Flow
1. **StartSessionCommand**(admin_id, purpose, scope, duration, requested_by, idempotency_key) → Validate() | RequestApproval() | Notify(communications-be) → **Outbox:** admin.session.requested.v1
2. **GrantScopeCommand**(session_id, scope, granted_by, conditions[], idempotency_key) → AuthorizeSuperAdmin() | Grant() | StartTimer() → **Outbox:** admin.break_glass.granted.v1
3. **EndSessionCommand**(session_id, ended_by, reason, idempotency_key) → RevokeScopes() | Terminate() → **Outbox:** admin.session.ended.v1
4. **AutoExpireSessionCommand**(session_id) → CheckExpiry() | Expire() | Notify(communications-be) → **Outbox:** admin.session.expired.v1
5. **GetSessionQuery**(session_id) → AuthorizeOwnerOrSuperAdmin() | Fetch() → SessionDTO
6. **ListActiveSessionsQuery**() → AuthorizeSuperAdmin() | FetchActive() → SessionListDTO
7. **GetSessionAuditQuery**(date_range, admin_id) → AuthorizeCompliance() | FetchAudit() → SessionAuditDTO

##### Projections
- admin_session_read
- active_sessions_read
- session_audit_read

##### Events Published
- admin.session.requested.v1
- admin.session.started.v1
- admin.break_glass.granted.v1
- admin.break_glass.denied.v1
- admin.session.ended.v1
- admin.session.expired.v1

##### RBAC/SLO
- **RBAC:** ADMIN (start/end own), SUPER_ADMIN (grant/deny)
- **SLO:** P95 < 200ms (start/grant/end/expire), P95 < 150ms (read)

---

### 10.2 Domain: change_approval/

#### 10.2.1 Entity: entity.go (Two-Person Rule)

##### User Stories
- As an **admin**, I want to **request approval** for sensitive changes so that two-person rule is enforced.
- As a **super admin**, I want to **approve/reject change requests** so that safety is ensured.
- As a **system**, I want to **expire unapproved requests** so that stale requests are cleared.
- As a **compliance officer**, I want to **audit approvals** so that policy compliance is verified.

##### Flow
1. **RequestChangeCommand**(change_type, resource_id, proposed_changes, justification, requested_by, idempotency_key) → Validate() | Enqueue() | NotifyApprovers() → **Outbox:** admin.change.requested.v1
2. **ApproveChangeCommand**(request_id, approved_by, conditions[], idempotency_key) → AuthorizeSuperAdmin() | Approve() | ExecuteChange() → **Outbox:** admin.change.approved.v1
3. **RejectChangeCommand**(request_id, rejected_by, reason, idempotency_key) → AuthorizeSuperAdmin() | Reject() | Notify(communications-be) → **Outbox:** admin.change.rejected.v1
4. **ExpireChangeCommand**(request_id) → CheckExpiry() | Expire() → **Outbox:** admin.change.expired.v1
5. **GetChangeRequestQuery**(request_id) → AuthorizeRequesterOrApprover() | Fetch() → ChangeRequestDTO
6. **ListChangeRequestsQuery**(filters, pagination) → AuthorizeSuperAdmin() | ApplyFilters() → ChangeRequestListDTO
7. **GetApprovalAuditQuery**(date_range) → AuthorizeCompliance() | FetchAudit() → ApprovalAuditDTO

##### Projections
- change_request_read
- pending_approvals_read
- approval_audit_read

##### Events Published
- admin.change.requested.v1
- admin.change.approved.v1
- admin.change.rejected.v1
- admin.change.expired.v1
- admin.change.executed.v1

##### RBAC/SLO
- **RBAC:** ADMIN (request), SUPER_ADMIN (approve/reject)
- **SLO:** P95 < 250ms (request/approve/reject), P95 < 150ms (expire), P95 < 200ms (read)

---

## **CROSS-SERVICE INTEGRATION**

### Outbound Dependencies

1. **users-be:** Fetch user profiles, publish user actions (suspend/ban/verify), update user status
2. **jobs-be:** Remove/hide/feature jobs, fetch job data for moderation
3. **proposals-be:** Remove/hide proposals, fetch proposal data for moderation
4. **contracts-be:** Resolve disputes, close contracts on bans, fetch contract data
5. **financial-be:** Place/release holds, process refunds, handle chargebacks, fetch transaction data
6. **subscriptions-be:** Override quotas, publish feature flag updates
7. **communications-be:** Send notifications (ticket updates, moderation actions, user suspensions, dispute resolutions)
8. **storage-be:** Upload reports/exports, manage file attachments
9. **search-be:** Publish search policy updates, index KB articles
10. **reviews-be:** Remove reviews, recalculate ratings

### Inbound Dependencies (Events Consumed)

1. **users-be:** user.flagged.v1, user.appeal.submitted.v1, user.support.request.created.v1, user.privacy.request.submitted.v1, user.created.v1 (sanctions screening)
2. **jobs-be:** job.flagged.v1
3. **proposals-be:** proposal.flagged.v1
4. **contracts-be:** contract.dispute.opened.v1
5. **financial-be:** payment.disputed.v1, payment.chargeback.received.v1, fraud.alert.triggered.v1, risk.alert.triggered.v1
6. **reviews-be:** review.flagged.v1
7. **communications-be:** message.flagged.v1
8. **storage-be:** storage.file.uploaded.v1, storage.file.virus.detected.v1
9. **search-be:** search.indexed.v1

---

## **GLOBAL SLO TARGETS**

### Read Operations
- Critical path (permissions): P95 < 30ms
- Simple queries: P95 < 100ms
- Complex queries: P95 < 200ms
- Reports/aggregations: P95 < 500ms

### Write Operations
- Simple writes: P95 < 200ms
- Complex writes: P95 < 400ms
- Multi-service sagas: P95 < 600ms

### Event Processing
- Event consumption: P95 < 150ms
- Event publishing: P95 < 100ms

### Background Jobs
- Auto-escalation: P95 < 300ms
- Session expiry: P95 < 200ms
- Sanctions screening: P95 < 500ms

---

## **CACHING STRATEGY**

### Redis Caching (TTL)
- Admin permissions: 5m
- Feature flags: 10m
- System config: 15m
- Canned responses: 30m
- KB articles: 10m
- Moderation queue: 2m
- Active sessions: 1m
- Throttle policies: 5m

### Cache Invalidation
- On admin.user.permissions.updated.v1 → Invalidate permissions
- On admin.feature_flag.updated.v1 → Invalidate flags
- On admin.config.updated.v1 → Invalidate config
- On canned_response.updated.v1 → Invalidate response
- On kb.article.updated.v1 → Invalidate article
- On admin.session.ended.v1 → Invalidate session

---

## **SECURITY & COMPLIANCE**

### PII Protection
- No raw PII in events
- Break-glass access for PII with approval
- All PII access audited
- Masking policies enforced

### GDPR/CCPA Compliance
- Full DSAR fulfillment (access + erasure)
- 30-day SLA for privacy requests
- Identity verification required
- Consent tracking

### Audit Requirements
- All admin actions logged (immutable)
- 10-year retention
- Full audit trail with context
- Legal hold support

### SOX Compliance
- Two-person rule for sensitive changes
- Change approval workflow
- Immutable change logs
- Separation of duties enforced

---

## **FINAL SUMMARY**

**Total Sections:** 10  
**Total Domains:** 40+  
**Total Entities:** 120+  
**Total User Stories:** 500+  
**Total Events Published:** 200+  
**Total Events Consumed:** 20+  
**Coverage:** 100% of admin-be folder structure  

**Pattern Compliance:**
✅ Event-Driven Architecture  
✅ CQRS with Projections  
✅ Outbox Pattern for Events  
✅ Idempotent Commands  
✅ Non-PII Events  
✅ RBAC per Operation  
✅ SLO per Operation  
✅ Platform Alignment  
✅ Sections → Domains → Entities → (Stories/Flow/Projections/Events/RBAC/SLO)

**Production Ready:**
✅ Complete domain coverage  
✅ Event sourcing  
✅ GDPR/CCPA compliance  
✅ Immutable audit trails  
✅ Break-glass access control  
✅ Two-person rule  
✅ Legal hold support  
✅ Privacy request handling  
✅ Sanctions screening  
✅ Fraud detection  
✅ Incident management  
✅ Bulk operations  
✅ Integration management  
✅ Policy management  
✅ Reporting & metrics  

---

**END OF admin-be USER STORIES**
