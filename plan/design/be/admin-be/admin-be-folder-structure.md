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
