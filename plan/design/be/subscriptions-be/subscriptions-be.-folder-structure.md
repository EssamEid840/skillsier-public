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
