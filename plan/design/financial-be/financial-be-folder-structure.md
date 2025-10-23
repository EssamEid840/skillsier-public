## 📦 **financial-be (Financial Management Service) - REFACTORED**

```

apps/be/financial-be/
│
├── cmd/
│   ├── api/
│   │   └── main.go                           # 📝 API entrypoint - initializes Gin, Dapr, Postgres (uses platform-shared/logging, internal/config)
│   │                                         # ♻️ Serve versioned routes (/v1/*) & enable ETag middleware when flag is on
│   │                                         # NO background jobs here - delegated to cmd/worker
│   │
│   └── worker/
│       └── main.go                           # 🆕 Worker entrypoint - background jobs, inbox/outbox, leader election
│                                             # Runs: payout batches, payment schedules, reminders, reconciliation, outbox dispatcher
│                                             # Uses: infrastructure/coordination for leader election & distributed locks
│
├── internal/
│   │
│   # ──────────────────────────────────────────────────────────────────────────────────
│   # 🔧 CONFIGURATION (Load First)
│   # ──────────────────────────────────────────────────────────────────────────────────
│   ├── config/
│   │   ├── schema.go                         # ♻️ Typed Config struct (App, Server, Postgres, Kafka, Redis, Auth, Stripe/PayPal keys, FX providers, risk thresholds)
│   │   │                                     # Group flags under Config.FeatureFlags
│   │   ├── feature_flags.go                  # 🆕 Central toggles (enable_ledger_audit, enable_etag, enable_fx_auto_update, enable_risk_holds, etc.)
│   │   ├── loader.go                         # Config loader using Viper (CLI → ENV → file → defaults)
│   │   └── docs/
│   │       └── CONFIGURATION.md              # Configuration documentation
│   │
│   # ──────────────────────────────────────────────────────────────────────────────────
│   # 🧩 DEPENDENCY INJECTION CONTAINER
│   # ──────────────────────────────────────────────────────────────────────────────────
│   ├── ioc/
│   │   ├── container.go                      # 🆕 DI graph: constructs DB/Redis/Kafka clients, repositories, services, handlers, schedulers
│   │   │                                     # Wire outbox, coordination, observers based on env
│   │   └── wiring.go                         # 🆕 Env-driven wiring & feature flags: selects implementations (local vs cloud)
│   │                                         # Wire outbox, coordination, observers based on env
│   │
│   # ──────────────────────────────────────────────────────────────────────────────────
│   # 🔧 INFRASTRUCTURE LAYER (Load Second)
│   # ──────────────────────────────────────────────────────────────────────────────────
│   ├── infrastructure/
│   │   │
│   │   # ===== DISTRIBUTED COORDINATION =====
│   │   ├── coordination/
│   │   │   ├── leader_election.go            # 🆕 Single-active cron/worker (ensures only one worker processes jobs)
│   │   │   │                                 # Used by: payout batches, payment schedules, reminders, reconciliation
│   │   │   └── distlock.go                   # 🆕 Redis/PG advisory locks for cron tasks (prevents duplicate execution)
│   │   │
│   │   # ===== PERSISTENCE LAYER =====
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go             # PostgreSQL connection (pooling, tracing, retries)
│   │   │       ├── transaction.go            # ♻️ Transaction helpers (unit-of-work, retry on serialization)
│   │   │       │                             # Ensure outbox write is in same DB tx
│   │   │       ├── migrations.go             # 📝 Auto-migration logic with version tracking
│   │   │       ├── version.go                # Schema version tracking
│   │   │       ├── safety.go                 # Pre-migration safety checks
│   │   │       │
│   │   │       ├── migrations/               # 🆕 SQL-first, forward-only migrations (prod)
│   │   │       │   ├── wallet/               # 🆕 Per-schema migration folders
│   │   │       │   ├── payment/              # 🆕
│   │   │       │   ├── escrow/               # 🆕
│   │   │       │   ├── tax/                  # 🆕
│   │   │       │   └── risk/                 # 🆕
│   │   │       │
│   │   │       # ===== CORE WALLET & LEDGER REPOSITORIES =====
│   │   │       ├── wallet_repository.go      # WalletRepository implementation (balances, reserves, CRUD)
│   │   │       ├── transaction_repository.go # TransactionRepository implementation (journal linkage, persist, lookup, list)
│   │   │       ├── ledger_journal_repository.go # LedgerJournalRepository (append-only journal, audit trail, hash chain verification)
│   │   │       │
│   │   │       # ===== PAYMENT & ESCROW REPOSITORIES =====
│   │   │       ├── payment_repository.go     # PaymentRepository implementation (provider refs, states, search)
│   │   │       ├── escrow_repository.go      # EscrowRepository implementation (holds, releases, refunds, pro-rata)
│   │   │       ├── payout_repository.go      # PayoutRepository implementation (queue, batch, process, cancel)
│   │   │       │
│   │   │       # ===== INVOICE & FEE REPOSITORIES =====
│   │   │       ├── invoice_repository.go     # InvoiceRepository implementation (lines, taxes, generation, send, mark paid)
│   │   │       ├── fee_repository.go         # FeeRepository implementation (fee rows, audits, calculations)
│   │   │       ├── fee_rules_repository.go   # FeeRulesRepository implementation (v2 rulesets: tiers, coupons, locale exceptions)
│   │   │       │
│   │   │       # ===== REFUND & DISPUTE REPOSITORIES =====
│   │   │       ├── refund_repository.go      # RefundRepository implementation (process, cancel, partial refunds, states, audit)
│   │   │       ├── dispute_payment_repository.go # DisputePaymentRepository implementation (chargebacks, representment, resolution)
│   │   │       │
│   │   │       # ===== TAX & FX REPOSITORIES =====
│   │   │       ├── tax_repository.go         # TaxRepository implementation (records, forms, filing, VAT/GST, 1099-K)
│   │   │       ├── fx_repository.go          # FXRepository implementation (time-based rates, quote/settlement, rounding rules)
│   │   │       │
│   │   │       # ===== RISK MANAGEMENT REPOSITORIES =====
│   │   │       ├── risk_repository.go        # RiskRepository implementation (holds, reserves, chargeback workflows, negative balances)
│   │   │       │
│   │   │       # ===== PROTECTION & FEE UPDATE REPOSITORIES =====
│   │   │       ├── protection_plan_repository.go  # ProtectionPlanRepository implementation (plans, claims, eligibility, payouts)
│   │   │       ├── fee_update_repository.go       # FeeUpdateRepository implementation (fee versions, migrations, impact calculations)
│   │   │       │
│   │   │       # ===== INTERNATIONAL PAYMENTS REPOSITORY =====
│   │   │       ├── international_payment_repository.go # InternationalPaymentRepository (routing, compliance checks, local rails, FX adjustments)
│   │   │       │
│   │   │       # ===== BONUS & EXPENSE REPOSITORIES =====
│   │   │       ├── bonus_repository.go            # BonusRepository implementation (award, pay, reverse, conditions)
│   │   │       ├── expense_repository.go          # ExpenseRepository implementation (submit, approve, reimburse, receipts)
│   │   │       │
│   │   │       # ===== PAYMENT SCHEDULE & REMINDER REPOSITORIES =====
│   │   │       ├── payment_schedule_repository.go # PaymentScheduleRepository (schedules, frequency, due windows, automation)
│   │   │       ├── reminder_repository.go         # ReminderRepository implementation (triggers, escalation, dunning, templates)
│   │   │       │
│   │   │       # ===== INSURANCE & TAX FORM REPOSITORIES =====
│   │   │       ├── insurance_repository.go        # InsuranceRepository implementation (policies, claims, coverage, providers)
│   │   │       ├── tax_form_repository.go         # TaxFormRepository implementation (W9, 1099, VAT returns, validation, reporting)
│   │   │       │
│   │   │       # ===== PAYROLL & CURRENCY REPOSITORIES =====
│   │   │       ├── payroll_repository.go          # PayrollRepository implementation (process, withholding, pay periods, reporting)
│   │   │       ├── currency_repository.go         # CurrencyRepository implementation (preferences, rate locks, conversions)
│   │   │       │
│   │   │       # ===== BANK ACCOUNT REPOSITORY =====
│   │   │       ├── bank_account_repository.go     # 🆕 BankAccountRepository implementation (CRUD, verification, default payout method)
│   │   │       │
│   │   │       # ===== OUTBOX REPOSITORY =====
│   │   │       └── outbox_store.go           # 🆕 Outbox repository (tx-coupled outbox table access for exactly-once publishing)
│   │   │
│   │   # ===== CACHING LAYER =====
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go             # Redis connection (pooling, timeouts)
│   │   │       ├── keys.go                   # 🆕 Canonical cache keys & TTLs (no magic strings)
│   │   │       ├── singleflight.go           # 🆕 Stampede protection for hot keys (wallet balances, FX rates)
│   │   │       ├── invalidation_rules.go     # 🆕 Map events → keys to drop (documented alongside)
│   │   │       │
│   │   │       # ===== DOMAIN-SPECIFIC CACHES =====
│   │   │       ├── wallet_cache.go           # Wallet balance caching (hot path reads)
│   │   │       ├── rate_cache.go             # Exchange rate caching (FX provider TTL)
│   │   │       ├── risk_cache.go             # Holds/reserve snapshots for quick checks
│   │   │       └── schedule_cache.go         # NextDue caches for payment schedules (cron efficiency)
│   │   │
│   │   # ===== EXTERNAL SERVICE CLIENTS =====
│   │   ├── clients/
│   │   │   ├── users/                        # Users-BE client (tax profiles, KYC/KYB statuses)
│   │   │   │   └── client.go                 # Fetch tax profiles & KYC statuses; circuit breaker/retry
│   │   │   │
│   │   │   ├── coupons/                      # Coupon/Promo validation backend (if separate service)
│   │   │   │   └── client.go                 # Validate & redeem promo codes; idempotent redemption
│   │   │   │
│   │   │   └── contracts/                    # 🆕 Contracts-BE client (contract status, milestones)
│   │   │       └── client.go                 # Query contract status for escrow releases & payment schedules
│   │   │
│   │   # ===== MESSAGING LAYER =====
│   │   ├── messaging/
│   │   │   ├── kafka/
│   │   │   │   ├── consumer.go               # Kafka consumer (uses platform-shared/inbox, DLQ, retries)
│   │   │   │   ├── producer.go               # Kafka producer (uses platform-shared/outbox, exactly-once semantics)
│   │   │   │   ├── topics.go                 # ♻️ Topics: payment.*, escrow.*, payout.*, fee.*, tax.*, fx.*, risk.*, bonus.*, expense.*, schedules.*, reminders.*
│   │   │   │   │                             # Thin re-export from contracts/events with per-context keys (financial.*)
│   │   │   │   └── scram.go                  # SCRAM authentication (SASL/SCRAM-SHA-512)
│   │   │   │
│   │   │   ├── outbox/                       # 🆕 Exactly-once-ish publisher
│   │   │   │   ├── dispatcher.go             # 🆕 Reads outbox → Kafka (with retries/DLQ, leader-elected)
│   │   │   │   └── metrics.go                # 🆕 Publish lag, failures, retries (Prometheus metrics)
│   │   │   │
│   │   │   └── bootstrap.go                  # 🆕 Consumer/producer wiring (initialize subscriptions, configure handlers)
│   │   │
│   │   # ===== PAYMENT GATEWAY INTEGRATIONS =====
│   │   ├── payment_gateway/
│   │   │   ├── stripe/
│   │   │   │   ├── client.go                 # Stripe API client (auth, retries, idempotency keys)
│   │   │   │   ├── webhook_handler.go        # Stripe webhook handler (payment_intent.succeeded, charge.failed)
│   │   │   │   │                             # 🆕 Includes signature verification middleware
│   │   │   │   └── mapper.go                 # Stripe → internal event mapper (normalize payloads)
│   │   │   │
│   │   │   ├── paypal/
│   │   │   │   ├── client.go                 # PayPal API client (orders/captures/refunds)
│   │   │   │   ├── webhook_handler.go        # PayPal webhook handler (capture completed/denied)
│   │   │   │   │                             # 🆕 Includes signature verification middleware
│   │   │   │   └── mapper.go                 # PayPal → internal event mapper
│   │   │   │
│   │   │   └── factory.go                    # Payment gateway factory (select provider by method/region, fallback/retry logic)
│   │   │
│   │   # ===== PDF GENERATION =====
│   │   ├── pdf/
│   │   │   └── generator.go                  # PDF invoice generation (wkhtmltopdf/Chrome headless wrapper)
│   │   │
│   │   # ===== SCHEDULER (BACKGROUND JOBS) =====
│   │   ├── scheduler/
│   │   │   ├── cron.go                       # ♻️ Cron registry: schedules, idempotency guards, safe shutdown
│   │   │   │                                 # Wrap each task with distlock + jitter
│   │   │   └── tasks/
│   │   │       ├── task_guard.go             # 🆕 Idempotency tokens + last-run watermark
│   │   │       ├── payout_batch_task.go      # 🆕 Batch payout processing (grouped by method/currency, runs in worker)
│   │   │       ├── payment_schedule_task.go  # 🆕 Run due payment schedules (auto-payment triggers)
│   │   │       ├── reminder_task.go          # 🆕 Send reminders (invoice due, tax form due, escalation)
│   │   │       ├── reconciliation_task.go    # 🆕 Daily reconciliation (bank vs journal, generates reports)
│   │   │       ├── fx_rate_update_task.go    # 🆕 Update FX rates from providers (periodic refresh)
│   │   │       └── risk_review_task.go       # 🆕 Review holds/reserves for auto-release (SLA windows)
│   │   │
│   │   # ===== OBSERVABILITY =====
│   │   ├── observability/                    # 🆕 Consolidated telemetry utilities
│   │   │   ├── metrics.go                    # 🆕 RED/USE metrics helpers (payment latency, escrow holds, payout success rate)
│   │   │   ├── tracing.go                    # 🆕 Common attrs (hashed user_id, org_id, event_id, payment_id)
│   │   │   └── slo_monitor.go                # 🆕 SLO monitoring (payment P99, payout batch lag, reconciliation delay)
│   │   │
│   │   # ===== HTTP UTILITIES =====
│   │   └── http/
│   │       ├── idempotency_adapter.go        # Bind platform-shared idempotency to Gin (critical for payments!)
│   │       └── etag.go                       # 🆕 Middleware-level ETag (pairs with utils/etag, optional feature flag)
│   │
│   # ──────────────────────────────────────────────────────────────────────────────────
│   # 🏛️ DOMAIN LAYER (Business Logic & Entities - Load Third)
│   # ──────────────────────────────────────────────────────────────────────────────────
│   ├── domain/
│   │   │
│   │   # ===== CORE WALLET & BALANCE =====
│   │   ├── wallet/
│   │   │   ├── entity.go                     # User wallet (UserID, Balance, Currency, Status, CreatedAt)
│   │   │   ├── balance.go                    # Balance tracking (Available, Pending, Reserved) with atomic adjustments
│   │   │   ├── currency.go                   # Multi-currency support (USD, EUR, GBP, etc.) and default currency rules
│   │   │   ├── errors.go                     # Wallet errors (InsufficientFunds, WalletNotFound, WalletClosed)
│   │   │   ├── repository.go                 # WalletRepository interface (CRUD, Reserve, Release, GetBalance)
│   │   │   └── events.go                     # Domain events: WalletCreated, WalletUpdated, WalletClosed, FundsReserved, FundsReleased
│   │   │
│   │   # ===== TRANSACTION & LEDGER =====
│   │   ├── transaction/
│   │   │   ├── entity.go                     # Financial transactions (ID, WalletID, Amount, Type, Status, Reference, CreatedAt)
│   │   │   ├── enums.go                      # Type (Deposit, Withdrawal, Transfer, Payment, Refund), Status (Pending, Completed, Failed), Category
│   │   │   ├── ledger.go                     # Double-entry ledger (debit/credit accounting) + balancing checks
│   │   │   ├── errors.go                     # Transaction errors (TransactionFailed, DuplicateTransaction, ImbalancedEntry)
│   │   │   ├── repository.go                 # TransactionRepository interface (persist, lookup, list, reconcile)
│   │   │   └── events.go                     # Domain events: TransactionInitiated, TransactionPosted, TransactionFailed, TransactionReversed
│   │   │
│   │   ├── ledger_journal/                   # Wallet ledger (double-entry) - immutable journal & audit
│   │   │   ├── entity.go                     # JournalEntry (ID, DebitAccount, CreditAccount, Amount, Currency, EffectiveAt, Hash, PrevHash)
│   │   │   ├── transfer.go                   # Transfers between accounts (builds 2-sided entries; validates debits=credits)
│   │   │   ├── adjustment.go                 # Manual adjustments with maker/checker approvals & notes
│   │   │   ├── audit.go                      # Audit trail (entry proofs, tamper-evidence via hash chain)
│   │   │   ├── errors.go                     # Ledger errors (Imbalance, ImmutableViolation, HashMismatch)
│   │   │   ├── repository.go                 # LedgerJournalRepository interface (append, list, audit)
│   │   │   └── events.go                     # Domain events: JournalEntryRecorded, JournalAdjustmentApproved, JournalAdjustmentRejected
│   │   │
│   │   # ===== PAYMENT PROCESSING =====
│   │   ├── payment/
│   │   │   ├── entity.go                     # Payment records (ID, Amount, Currency, PayerID, PayeeID, Method, Status, ProviderRef)
│   │   │   ├── payment_method.go             # Payment methods (CreditCard, PayPal, BankTransfer, Wallet)
│   │   │   ├── method_profile.go             # Payout preferences per method (limits, cutoffs, fees, destination IDs)
│   │   │   ├── onboarding.go                 # Onboarding KYC/KYB statuses per method/provider (required docs, states)
│   │   │   ├── gateway.go                    # Payment gateway abstraction (Stripe, PayPal interface) & retries
│   │   │   ├── errors.go                     # Payment errors (PaymentFailed, InvalidCard, GatewayTimeout, InsufficientFunds)
│   │   │   ├── repository.go                 # PaymentRepository interface (persist, update, search, reconcile)
│   │   │   └── events.go                     # Domain events: PaymentAuthorized, PaymentCaptured, PaymentFailed, PaymentRefundInitiated
│   │   │
│   │   # ===== ESCROW MANAGEMENT =====
│   │   ├── escrow/
│   │   │   ├── entity.go                     # Escrow accounts (ID, ContractID, Amount, Status, HeldAt, ReleasedAt)
│   │   │   ├── hold.go                       # Fund holds (amount reserved until milestone completion or dispute close)
│   │   │   ├── release.go                    # Fund release conditions (milestone approved, dispute resolved, partial releases)
│   │   │   ├── pro_rata.go                   # Pro-rata partial releases by milestone & refunds (distribution calc)
│   │   │   ├── errors.go                     # Escrow errors (EscrowNotFound, InsufficientEscrow, EscrowLocked)
│   │   │   ├── repository.go                 # EscrowRepository interface (create, hold, release, refund, partial release)
│   │   │   └── events.go                     # Domain events: EscrowFunded, EscrowPartiallyReleased, EscrowReleased, EscrowRefunded
│   │   │
│   │   # ===== PAYOUT PROCESSING =====
│   │   ├── payout/
│   │   │   ├── entity.go                     # Payout requests (ID, UserID, Amount, Method, Status, RequestedAt, ProcessedAt)
│   │   │   ├── method.go                     # Payout methods (BankTransfer, PayPal, Payoneer, Wire) & metadata
│   │   │   ├── schedule.go                   # Payout scheduling (instant, daily, weekly, monthly) with cut-off logic
│   │   │   ├── errors.go                     # Payout errors (PayoutFailed, BelowMinimum, MethodUnavailable, InsufficientBalance)
│   │   │   ├── repository.go                 # PayoutRepository interface (queue, batch, update, cancel, process)
│   │   │   └── events.go                     # Domain events: PayoutRequested, PayoutScheduled, PayoutProcessed, PayoutFailed
│   │   │
│   │   # ===== INVOICE MANAGEMENT =====
│   │   ├── invoice/
│   │   │   ├── entity.go                     # Invoice generation (ID, ContractID, Number, Amount, DueDate, Status, PaidAt)
│   │   │   ├── line_item.go                  # Invoice line items (description, quantity, price, total, tax)
│   │   │   ├── tax.go                        # Tax calculations (VAT, sales tax by jurisdiction & reverse charge flags)
│   │   │   ├── errors.go                     # Invoice errors (InvoiceNotFound, AlreadyPaid, InvalidTotals, InvalidTax)
│   │   │   ├── repository.go                 # InvoiceRepository interface (create, send, update, lookup, mark paid, cancel)
│   │   │   └── events.go                     # Domain events: InvoiceIssued, InvoiceUpdated, InvoicePaid, InvoiceOverdue, InvoiceCanceled
│   │   │
│   │   # ===== PLATFORM FEES =====
│   │   ├── fee/
│   │   │   ├── entity.go                     # Platform fees (ID, TransactionID, Amount, Type, Rate, Currency)
│   │   │   ├── calculator.go                 # Fee calculation rules (percentage, tiered rates, min/max caps)
│   │   │   ├── tier.go                       # Fee tiers (based on subscription, volume, user type)
│   │   │   │
│   │   │   ├── v2/                           # Fee engine v2 (advanced rules)
│   │   │   │   ├── rules.go                  # Rules model (volume discounts, locale overrides, coupons, experiments)
│   │   │   │   ├── coupon.go                 # Coupon/promo code entity & redemption logic (single-use, stacking)
│   │   │   │   ├── country_exceptions.go     # Country/regional exceptions (local regulation overrides)
│   │   │   │   ├── errors.go                 # Fee v2 errors (InvalidCoupon, IneligibleTier, ExpiredRule)
│   │   │   │   └── repository.go             # FeeRulesRepository interface (CRUD for rule sets)
│   │   │   │
│   │   │   ├── errors.go                     # Fee errors (FeeConfigMissing, NegativeFee, InvalidTier)
│   │   │   ├── repository.go                 # FeeRepository interface (persist & audit fee applications)
│   │   │   └── events.go                     # Domain events: FeeCalculated, FeeAdjusted, CouponApplied, CouponRevoked
│   │   │
│   │   # ===== REFUND PROCESSING =====
│   │   ├── refund/
│   │   │   ├── entity.go                     # Refund processing (ID, PaymentID, Amount, Reason, Status, ProcessedAt)
│   │   │   ├── policy.go                     # Refund policies (full, partial, time limits, dispute linkage)
│   │   │   ├── pro_rata.go                   # Pro-rata refunds aligned to milestones (safe allocation)
│   │   │   ├── errors.go                     # Refund errors (RefundNotAllowed, RefundExpired, OverRefund, InvalidAmount)
│   │   │   ├── repository.go                 # RefundRepository interface (queue, process, audit, cancel)
│   │   │   └── events.go                     # Domain events: RefundRequested, RefundApproved, RefundDeclined, RefundProcessed
│   │   │
│   │   # ===== PAYMENT DISPUTES =====
│   │   ├── dispute_payment/
│   │   │   ├── entity.go                     # Payment disputes (ID, PaymentID, Reason, FiledBy, Status, Resolution)
│   │   │   ├── chargeback.go                 # Chargeback handling (card network flows, representment)
│   │   │   ├── errors.go                     # Dispute errors (DisputeNotFound, ChargebackInvalidState, AlreadyResolved)
│   │   │   ├── repository.go                 # DisputePaymentRepository interface (open, update, resolve, representment)
│   │   │   └── events.go                     # Domain events: PaymentDisputeOpened, ChargebackRecorded, PaymentDisputeResolved
│   │   │
│   │   # ===== TAX MANAGEMENT =====
│   │   ├── tax/
│   │   │   ├── entity.go                     # Tax records (ID, UserID, Year, Type, Amount, Status, Jurisdiction)
│   │   │   ├── vat_gst.go                    # VAT/GST per locale & reverse-charge logic (EU/UK/CA, etc.)
│   │   │   ├── reverse_charge.go             # Reverse-charge mechanics & validations (B2B cross-border)
│   │   │   ├── forms_1099k.go                # 1099-K forms generation & thresholds (US)
│   │   │   ├── profile_link.go               # Link to Users-BE tax profile (TIN/VAT ID; sync contract-level data)
│   │   │   ├── form.go                       # Tax forms (W9, 1099, VAT returns; states & versions)
│   │   │   ├── errors.go                     # Tax errors (TaxProfileMissing, InvalidVATID, FormInvalid, ThresholdNotMet)
│   │   │   ├── repository.go                 # TaxRepository interface (create, update, file, generate forms)
│   │   │   └── events.go                     # Domain events: TaxRecordCreated, TaxFormGenerated, TaxProfileLinked, VATReverseChargeApplied
│   │   │
│   │   # ===== FOREIGN EXCHANGE (FX) =====
│   │   ├── fx/
│   │   │   ├── rate.go                       # FXRate (Base, Quote, Rate, EffectiveFrom, Provider, Precision, Source)
│   │   │   ├── quote_settlement.go           # Quote vs. settlement currency handling (timing, spreads)
│   │   │   ├── rounding.go                   # Rounding rules per currency/amount type (bankers, up/down)
│   │   │   ├── errors.go                     # FX errors (RateNotFound, StaleRate, PrecisionExceeded, InvalidPair)
│   │   │   ├── repository.go                 # FXRepository interface (upsert rate, query time-bounded, get effective rate)
│   │   │   └── events.go                     # Domain events: FXRateUpdated, FXQuoteCreated, FXSettlementCalculated
│   │   │
│   │   # ===== RISK MANAGEMENT =====
│   │   ├── risk/
│   │   │   ├── reserve.go                    # Reserve balances & rolling reserve schedules (percent/time windows)
│   │   │   ├── hold_workflow.go              # Holds/auto-release based on risk signals & time (SLA windows)
│   │   │   ├── chargeback_workflow.go        # Chargeback workflows & recovery (fees, evidence, outcomes)
│   │   │   ├── negative_balance.go           # Negative balance handling & collections (repayment plans)
│   │   │   ├── errors.go                     # Risk errors (ReserveConflict, HoldNotFound, NegativeBalanceLimit)
│   │   │   ├── repository.go                 # RiskRepository interface (holds, reserves, risk metrics, negative balance tracking)
│   │   │   └── events.go                     # Domain events: ReserveSet, ReserveUpdated, RiskHoldPlaced, RiskHoldReleased, NegativeBalanceCreated
│   │   │
│   │   # ===== PAYMENT PROTECTION PLANS =====
│   │   ├── protection_plan/
│   │   │   ├── entity.go                     # ProtectionPlan aggregate (ContractID, Type, CoverageAmount, Premium, Status, Period)
│   │   │   ├── type.go                       # Plan types (HourlyProtection, FixedPriceProtection) & defaults
│   │   │   ├── claim.go                      # Protection claims (filing, review, payout; integrates with risk & escrow)
│   │   │   ├── eligibility.go                # Eligibility checks (contract type, verification, history)
│   │   │   ├── errors.go                     # Protection errors (NotEligible, ClaimDenied, CoverageExceeded, AlreadyClaimed)
│   │   │   ├── repository.go                 # ProtectionPlanRepository interface (plan lifecycle, claims, eligibility checks)
│   │   │   └── events.go                     # Domain events: ProtectionApplied, ClaimFiled, ClaimApproved, ClaimPaid
│   │   │
│   │   # ===== FEE UPDATES & MIGRATIONS =====
│   │   ├── fee_update/
│   │   │   ├── entity.go                     # FeeUpdate aggregate (Version, EffectiveDate, Rules[], Impact, Status)
│   │   │   ├── version.go                    # Fee versions (e.g., flat 10% for freelancers, 3–5% for clients)
│   │   │   ├── impact.go                     # Impact calculations (per segment; notifications, proration)
│   │   │   ├── migration.go                  # Fee migration logic (from old to new rule sets; rollback)
│   │   │   ├── errors.go                     # FeeUpdate errors (VersionConflict, InvalidRate, MigrationFailed, ActiveVersionExists)
│   │   │   ├── repository.go                 # FeeUpdateRepository interface (publish, activate, archive, rollback)
│   │   │   └── events.go                     # Domain events: FeeUpdated, ImpactCalculated, Migrated, RolledBack
│   │   │
│   │   # ===== INTERNATIONAL PAYMENTS =====
│   │   ├── international_payment/
│   │   │   ├── entity.go                     # InternationalPayment (TransactionID, LocalCurrency, ComplianceChecks, Routing)
│   │   │   ├── compliance.go                 # OFAC/AML/Sanctions checks; links user/tax profile for KYC/KYB
│   │   │   ├── local_method.go               # Local payout methods (SEPA, ACH, local banks) + routing preferences
│   │   │   ├── fee_adjustment.go             # Cross-border fee adjustments (FX spreads, local rails fees)
│   │   │   ├── errors.go                     # InternationalPayment errors (ComplianceFailed, MethodUnavailable, RoutingError)
│   │   │   ├── repository.go                 # InternationalPaymentRepository interface (route, track status, compliance checks)
│   │   │   └── events.go                     # Domain events: InternationalPaymentInitiated, InternationalPaymentCompliant, InternationalPaymentProcessed
│   │   │
│   │   # ===== BONUS PAYMENTS =====
│   │   ├── bonus/
│   │   │   ├── entity.go                     # Bonus aggregate (ContractID, Amount, Reason, AwardedBy, AwardedAt, Status)
│   │   │   ├── type.go                       # Bonus types (Performance, Completion, Referral) & default flows
│   │   │   ├── condition.go                  # Bonus conditions (milestone met, on-time delivery, KPI triggers)
│   │   │   ├── errors.go                     # Bonus errors (BonusNotEligible, AlreadyAwarded, InvalidReason, AmountExceedsLimit)
│   │   │   ├── repository.go                 # BonusRepository interface (award, reverse, list, get by contract)
│   │   │   └── events.go                     # Domain events: BonusRequested, BonusAwarded, BonusRejected, BonusPaid
│   │   │
│   │   # ===== EXPENSE REIMBURSEMENTS =====
│   │   ├── expense/
│   │   │   ├── entity.go                     # Expense aggregate (ContractID, ItemID, Amount, Description, SubmittedBy, Status, SubmittedAt)
│   │   │   ├── receipt.go                    # Receipt attachments and verification (OCR hooks)
│   │   │   ├── approval.go                   # Approval workflow (client approval required; levels)
│   │   │   ├── policy.go                     # Reimbursement policies (caps, eligible categories, per diem)
│   │   │   ├── errors.go                     # Expense errors (ExpenseNotApproved, OverCap, InvalidReceipt, MissingDocumentation)
│   │   │   ├── repository.go                 # ExpenseRepository interface (submit, approve, reimburse, reject)
│   │   │   └── events.go                     # Domain events: ExpenseSubmitted, ExpenseApproved, ExpenseRejected, ExpenseReimbursed
│   │   │
│   │   # ===== PAYMENT SCHEDULES =====
│   │   ├── payment_schedule/
│   │   │   ├── entity.go                     # PaymentSchedule aggregate (ContractID, Frequency, Amount, NextDue, Status)
│   │   │   ├── frequency.go                  # Payment frequency (Weekly, BiWeekly, Monthly) & cut-off rules
│   │   │   ├── adjustment.go                 # Schedule adjustments (prorating, skips, deferrals)
│   │   │   ├── automation.go                 # Auto-payment triggers (cron, Dapr bindings)
│   │   │   ├── errors.go                     # Schedule errors (ScheduleConflict, PaymentMissed, DisabledSchedule, InsufficientFunds)
│   │   │   ├── repository.go                 # PaymentScheduleRepository interface (create, update, list, due, process)
│   │   │   └── events.go                     # Domain events: ScheduleCreated, ScheduleUpdated, PaymentDue, PaymentProcessed
│   │   │
│   │   # ===== AUTOMATED REMINDERS =====
│   │   ├── reminder/
│   │   │   ├── entity.go                     # Reminder aggregate (ContractID, Type, Trigger, SentAt, NextAt, Attempts)
│   │   │   ├── type.go                       # Reminder types (InvoicePay, PayoutDue, TaxFormDue, ScheduleRun)
│   │   │   ├── template.go                   # Reminder templates (customizable messages; locale-aware)
│   │   │   ├── escalation.go                 # Escalation paths (repeated reminders, penalties, dunning)
│   │   │   ├── errors.go                     # Reminder errors (ReminderNotSent, ChannelUnavailable, ThrottleLimitReached)
│   │   │   ├── repository.go                 # ReminderRepository interface (queue, mark sent, escalate, list)
│   │   │   └── events.go                     # Domain events: ReminderTriggered, ReminderSent, ReminderEscalated
│   │   │
│   │   # ===== INSURANCE & PROTECTION =====
│   │   ├── insurance/
│   │   │   ├── entity.go                     # Insurance aggregate (ContractID, PolicyID, Coverage, Premium, Status)
│   │   │   ├── coverage.go                   # Coverage details (payment protection, liability, limits)
│   │   │   ├── claim.go                      # Insurance claims (filing, status, payout; external provider codes)
│   │   │   ├── provider.go                   # Integration with insurance providers (quote/bind/claim)
│   │   │   ├── errors.go                     # Insurance errors (PolicyNotActive, ClaimDenied, ProviderUnavailable, CoverageExceeded)
│   │   │   ├── repository.go                 # InsuranceRepository interface (create policy, claim lifecycle, provider integration)
│   │   │   └── events.go                     # Domain events: InsuranceApplied, ClaimFiled, ClaimApproved, ClaimPaid
│   │   │
│   │   # ===== TAX FORMS =====
│   │   ├── tax_form/
│   │   │   ├── entity.go                     # TaxForm aggregate (ContractID, FormType, SubmittedBy, Status, SubmittedAt)
│   │   │   ├── type.go                       # Form types (W9, 1099, VAT ID, 1099-K support docs)
│   │   │   ├── validation.go                 # Form validation and verification (format & identity checks)
│   │   │   ├── reporting.go                  # Tax reporting integration (export to tax authorities/providers)
│   │   │   ├── errors.go                     # Tax form errors (FormInvalid, NotSubmitted, VerificationFailed, IncompleteData)
│   │   │   ├── repository.go                 # TaxFormRepository interface (submit, verify, fetch, list, report)
│   │   │   └── events.go                     # Domain events: TaxFormSubmitted, TaxFormVerified, TaxReportGenerated
│   │   │
│   │   # ===== PAYROLL INTEGRATION =====
│   │   ├── payroll/
│   │   │   ├── entity.go                     # Payroll aggregate (ContractID, WorkerStatus, PayPeriods[], Taxes, Deductions)
│   │   │   ├── status.go                     # Worker classification (Freelancer, Employee) + validation hooks
│   │   │   ├── pay_period.go                 # Pay periods (bi-weekly, monthly; with deductions & employer taxes)
│   │   │   ├── tax.go                        # Tax withholding and reporting (per jurisdiction)
│   │   │   ├── errors.go                     # Payroll errors (ClassificationMismatch, PeriodClosed, InsufficientDeductions)
│   │   │   ├── repository.go                 # PayrollRepository interface (process, post, report, withhold)
│   │   │   └── events.go                     # Domain events: PayrollProcessed, TaxWithheld, PayPeriodClosed
│   │   │
│   │   # ===== CURRENCY MANAGEMENT =====
│   │   ├── currency/
│   │   │   ├── entity.go                     # Currency aggregate (ContractID, PreferredCurrency, ExchangeRates, LockedAt)
│   │   │   ├── conversion.go                 # Real-time conversion logic (preferred currency, rounding)
│   │   │   ├── rate_lock.go                  # Rate locking for payments (validity windows, expirations)
│   │   │   ├── errors.go                     # Currency errors (ConversionFailed, RateNotAvailable, LockExpired, InvalidCurrency)
│   │   │   ├── repository.go                 # CurrencyRepository interface (locks, preferences, histories, conversions)
│   │   │   └── events.go                     # Domain events: CurrencyChanged, RateLocked, Converted
│   │   │
│   │   # ===== BANK ACCOUNT MANAGEMENT =====
│   │   └── bank_account/                     # 🆕 Bank account domain (for exposed endpoints)
│   │       ├── entity.go                     # BankAccount aggregate (UserID, AccountNumber, RoutingNumber, BankName, IsDefault, Status, VerifiedAt)
│   │       ├── verifier.go                   # Bank account verification (micro-deposits, Plaid integration)
│   │       ├── errors.go                     # BankAccount errors (AccountNotFound, VerificationFailed, InvalidRouting, AlreadyExists)
│   │       ├── repository.go                 # BankAccountRepository interface (CRUD, verify, set default, list by user)
│   │       └── events.go                     # Domain events: BankAccountAdded, BankAccountVerified, BankAccountSetDefault, BankAccountRemoved
│   │
│   # ──────────────────────────────────────────────────────────────────────────────────
│   # 📋 APPLICATION LAYER (Use Cases & Orchestration - Load Fourth)
│   # ──────────────────────────────────────────────────────────────────────────────────
│   ├── application/
│   │   │
│   │   # ===== EVENT HANDLERS (Inbound Events) =====
│   │   ├── eventhandler/
│   │   │   ├── _common/
│   │   │   │   └── idempotency.go            # 🆕 Store last-seen event version per aggregate (prevent duplicate processing)
│   │   │   │
│   │   │   ├── billing_handler.go            # Consumes: subscriptions-be → billing.invoice.exported (ingest AR for charging)
│   │   │   ├── admin_risk_handler.go         # Consumes: risk.* (hold/reserve/chargeback/velocity alerts → apply effects)
│   │   │   ├── admin_flags_handler.go        # Consumes: admin.feature_flag/threshold/experiment.updated (runtime behaviors)
│   │   │   ├── seats_billing_handler.go      # Consumes: seat usage/billing streams (seat changes → invoice lines)
│   │   │   ├── contract_handler.go           # 🆕 Consumes: contract.* (contract status changes → escrow releases, payment schedules)
│   │   │   └── dispute_handler.go            # 🆕 Consumes: dispute.* (dispute events → escrow holds, refund triggers)
│   │   │
│   │   # ===== AUTHORIZATION LAYER =====
│   │   ├── authz/                            # 🆕 Service-level permission checks (defense-in-depth)
│   │   │   ├── policies.go                   # 🆕 Map roles → actions (admin, wallet_owner, contract_party, finance_admin)
│   │   │   └── guards.go                     # 🆕 Helpers used by services (CanAccessWallet, CanInitiateRefund, CanViewInvoice)
│   │   │
│   │   # ===== CORE WALLET & LEDGER SERVICES =====
│   │   ├── wallet/
│   │   │   ├── service.go                    # Wallet logic (Create, Deposit, Withdraw, Transfer, GetBalance, Reserve, Release)
│   │   │   ├── commands.go                   # DepositCommand, WithdrawCommand, TransferCommand, CreateWallet, ReserveFunds, ReleaseFunds
│   │   │   ├── queries.go                    # GetBalance, GetHistory, GetTransactions (filter/pagination)
│   │   │   ├── validators.go                 # Validate currency codes, positive amounts, sufficient balance, wallet status
│   │   │   ├── dto.go                        # WalletDTO, TransactionDTO, BalanceDTO (in/out mapping contracts)
│   │   │   └── mapper.go                     # Wallet ⇄ DTO mapping helpers
│   │   │
│   │   ├── transaction/
│   │   │   ├── service.go                    # Transaction logic (Create, Reverse, Reconcile with journal)
│   │   │   ├── commands.go                   # CreateTransaction, ReverseTransaction, ReconcileTransactions
│   │   │   ├── queries.go                    # GetTransaction, ListTransactions, GetLedger (with filters)
│   │   │   ├── validators.go                 # Immutable journal checks, debits=credits, idempotency keys
│   │   │   ├── dto.go                        # TransactionDTO, LedgerDTO, ReconciliationReportDTO
│   │   │   ├── mapper.go                     # Transaction ⇄ DTO mappers
│   │   │   └── reconciliation.go             # Reconciliation logic (match payments with bank statements)
│   │   │
│   │   ├── ledger_journal/
│   │   │   ├── service.go                    # AppendEntry, Transfer, Adjust, AuditTrail orchestration
│   │   │   ├── commands.go                   # AppendJournalEntry, TransferFunds, CreateAdjustment (maker/checker)
│   │   │   ├── queries.go                    # GetEntry, ListEntries, GetAuditTrail (hash continuity)
│   │   │   ├── validators.go                 # Entry hash chain, debits=credits, approval gates
│   │   │   ├── dto.go                        # JournalEntryDTO, TransferDTO, AdjustmentDTO, AuditTrailDTO
│   │   │   └── mapper.go                     # Journal ⇄ DTO mappers
│   │   │
│   │   # ===== PAYMENT PROCESSING SERVICES =====
│   │   ├── payment/
│   │   │   ├── service.go                    # Process, Capture, Void, Refund (routes to provider processors)
│   │   │   ├── commands.go                   # ProcessPayment, CapturePayment, VoidPayment, RefundPayment
│   │   │   ├── queries.go                    # GetPayment, ListPayments (status/provider filters)
│   │   │   ├── validators.go                 # Method eligibility, onboarding status, amount/limits, risk holds
│   │   │   ├── dto.go                        # PaymentDTO, ProcessPaymentDTO, PaymentResultDTO
│   │   │   ├── mapper.go                     # Payment ⇄ DTO mappers
│   │   │   ├── stripe_processor.go           # Stripe payment processor implementation (auth/capture/refund)
│   │   │   ├── paypal_processor.go           # PayPal payment processor implementation (orders/captures/refunds)
│   │   │   └── processor_factory.go          # Payment processor factory (select by method/provider; fallback/retry)
│   │   │
│   │   # ===== ESCROW MANAGEMENT SERVICES =====
│   │   ├── escrow/
│   │   │   ├── service.go                    # Hold, Release, Refund flows (interacts with contracts-be events)
│   │   │   ├── commands.go                   # HoldEscrow, ReleaseEscrow, RefundEscrow, PartialReleaseEscrow
│   │   │   ├── queries.go                    # GetEscrow, ListEscrows, GetEscrowHistory
│   │   │   ├── validators.go                 # Coverage checks, release conditions, pro-rata math, dispute gates
│   │   │   ├── dto.go                        # EscrowDTO, HoldEscrowDTO, ReleaseEscrowDTO, PartialReleaseDTO
│   │   │   ├── mapper.go                     # Escrow ⇄ DTO mappers
│   │   │   └── pro_rata_release_manager.go   # Dispute-driven conditional releases computation engine
│   │   │
│   │   # ===== PAYOUT PROCESSING SERVICES =====
│   │   ├── payout/
│   │   │   ├── service.go                    # Request, Schedule/Batch, Process, Cancel, GetHistory
│   │   │   ├── commands.go                   # RequestPayout, ProcessPayout, CancelPayout, SchedulePayouts
│   │   │   ├── queries.go                    # GetPayouts, GetPayoutHistory (time/status filters)
│   │   │   ├── validators.go                 # Method limits, compliance holds, minimum amounts, KYC checks
│   │   │   ├── dto.go                        # PayoutDTO, RequestPayoutDTO, PayoutBatchDTO
│   │   │   ├── mapper.go                     # Payout ⇄ DTO mappers
│   │   │   └── batch_processor.go            # Batch payout processor (grouped by method/currency, runs in worker)
│   │   │
│   │   # ===== INVOICE MANAGEMENT SERVICES =====
│   │   ├── invoice/
│   │   │   ├── service.go                    # Generate, Send, MarkPaid, Cancel; PDF rendering
│   │   │   ├── commands.go                   # GenerateInvoice, SendInvoice, MarkInvoicePaid, CancelInvoice
│   │   │   ├── queries.go                    # GetInvoice, ListInvoices (filters: contract/user/status/date)
│   │   │   ├── validators.go                 # Line totals, tax rounding, currency consistency, unique number
│   │   │   ├── dto.go                        # InvoiceDTO, LineItemDTO, GenerateInvoiceDTO
│   │   │   ├── mapper.go                     # Invoice ⇄ DTO mappers
│   │   │   ├── generator.go                  # Invoice PDF generation orchestrator
│   │   │   └── tax_calculator.go             # Tax calculation (VAT, sales tax; reverse charge)
│   │   │
│   │   # ===== FEE CALCULATION SERVICES =====
│   │   ├── fee/
│   │   │   ├── service.go                    # Calculate, Apply, Waive fees to transactions/payouts
│   │   │   ├── calculator.go                 # Base calculator (percentage, tiered, flat; min/max caps)
│   │   │   ├── validators.go                 # Fee caps, non-negative, eligibility & stacking rules
│   │   │   ├── dto.go                        # FeeDTO, FeeRequestDTO, FeeCalculationResultDTO
│   │   │   └── rules_engine.go               # Rules engine (user type, volume, geography)
│   │   │
│   │   ├── fee_v2/
│   │   │   ├── service.go                    # Tiered/volume, locale exceptions, coupons, experiments
│   │   │   ├── commands.go                   # ApplyCoupon, RevokeCoupon, UpdateFeeRules
│   │   │   ├── queries.go                    # GetEffectiveFee, GetUserFeeTier, GetCoupon
│   │   │   ├── validators.go                 # Coupon validity, stacking, locale overrides, experiment flags
│   │   │   ├── dto.go                        # FeeV2DTO, CouponDTO, FeeRuleDTO
│   │   │   └── mapper.go                     # Fee v2 ⇄ DTO mappers
│   │   │
│   │   # ===== REFUND PROCESSING SERVICES =====
│   │   ├── refund/
│   │   │   ├── service.go                    # Process & Cancel refunds (full/partial) with audit trail
│   │   │   ├── commands.go                   # ProcessRefund, CancelRefund, ProcessPartialRefund
│   │   │   ├── queries.go                    # GetRefund, ListRefunds (by payment/user/contract)
│   │   │   ├── validators.go                 # Pro-rata vs eligibility windows, idempotency, limits
│   │   │   ├── dto.go                        # RefundDTO, RefundRequestDTO, PartialRefundDTO
│   │   │   └── mapper.go                     # Refund ⇄ DTO mappers
│   │   │
│   │   # ===== TAX SERVICES =====
│   │   ├── tax/
│   │   │   ├── service.go                    # Calculate, Generate forms, File (or export)
│   │   │   ├── form_generator.go             # Generate tax forms (W9, 1099, 1099-K, VAT returns)
│   │   │   ├── commands.go                   # SyncTaxProfileFromUsersBE, Generate1099K, FileVATReturn
│   │   │   ├── queries.go                    # GetTaxProfile, GetVATRate, GetReverseChargeEligibility
│   │   │   ├── validators.go                 # VAT ID formats per locale, threshold checks, 1099-K thresholds
│   │   │   ├── dto.go                        # TaxDTO, TaxFormDTO, VATRateDTO
│   │   │   └── mapper.go                     # Tax ⇄ DTO mappers
│   │   │
│   │   # ===== FOREIGN EXCHANGE SERVICES =====
│   │   ├── fx/
│   │   │   ├── service.go                    # Quote→settlement conversion, fetch/apply FX rates, rounding
│   │   │   ├── commands.go                   # UpsertFXRate, SetRoundingRule
│   │   │   ├── queries.go                    # GetFXRate, ConvertAmount, GetEffectiveRateAt
│   │   │   ├── validators.go                 # Effective timestamps, precision, allowed currency pairs
│   │   │   ├── dto.go                        # FXRateDTO, ConversionRequestDTO, ConversionResultDTO
│   │   │   └── mapper.go                     # FX ⇄ DTO mappers
│   │   │
│   │   # ===== RISK MANAGEMENT SERVICES =====
│   │   ├── risk/
│   │   │   ├── service.go                    # Holds, reserves, chargebacks, negative balances orchestration
│   │   │   ├── commands.go                   # CreateReserve, ReleaseReserve, PlaceHold, RemoveHold, RecordChargeback
│   │   │   ├── queries.go                    # GetReserves, GetHolds, GetRiskScore (if available)
│   │   │   ├── validators.go                 # Rolling reserve schedules, chargeback states, policy gates
│   │   │   ├── dto.go                        # RiskDTO, ReserveDTO, HoldDTO, ChargebackDTO
│   │   │   └── mapper.go                     # Risk ⇄ DTO mappers
│   │   │
│   │   # ===== PROTECTION PLAN SERVICES =====
│   │   ├── protection_plan/
│   │   │   ├── service.go                    # Apply plan, file/approve claims, compute payouts
│   │   │   ├── commands.go                   # ApplyProtection, FileClaim, ApproveClaim, PayClaim
│   │   │   ├── queries.go                    # GetPlan, GetClaims, GetEligibility
│   │   │   ├── validators.go                 # Eligibility checks & coverage caps
│   │   │   ├── dto.go                        # ProtectionPlanDTO, ClaimDTO, EligibilityDTO
│   │   │   └── mapper.go                     # Protection plan ⇄ DTO mappers
│   │   │
│   │   # ===== FEE UPDATE & MIGRATION SERVICES =====
│   │   ├── fee_update/
│   │   │   ├── service.go                    # Manage fee versions & migrations; compute impact
│   │   │   ├── commands.go                   # PublishFeeVersion, RunFeeMigration, ActivateFeeVersion, RollbackFeeVersion
│   │   │   ├── queries.go                    # GetActiveFeeVersion, GetImpact, ListFeeVersions
│   │   │   ├── validators.go                 # Version monotonicity, rate bounds, window checks
│   │   │   ├── dto.go                        # FeeUpdateDTO, FeeImpactDTO, FeeVersionDTO
│   │   │   └── mapper.go                     # Fee update ⇄ DTO mappers
│   │   │
│   │   # ===== INTERNATIONAL PAYMENT SERVICES =====
│   │   ├── international_payment/
│   │   │   ├── service.go                    # Run compliance, convert FX, route to local rails
│   │   │   ├── commands.go                   # InitiateInternationalPayment, MarkCompliant, RoutePayment
│   │   │   ├── queries.go                    # GetInternationalPayment, ListInternationalPayments
│   │   │   ├── validators.go                 # OFAC/AML gates, supported corridors
│   │   │   ├── dto.go                        # InternationalPaymentDTO, ComplianceDTO
│   │   │   └── mapper.go                     # Intl payment ⇄ DTO mappers
│   │   │
│   │   # ===== BONUS SERVICES =====
│   │   ├── bonus/
│   │   │   ├── service.go                    # Award & pay bonuses, enforce conditions
│   │   │   ├── commands.go                   # AwardBonus, RejectBonus, PayBonus
│   │   │   ├── queries.go                    # GetBonuses (by contract/user), GetBonusPolicy
│   │   │   ├── validators.go                 # Condition checks, caps, eligibility
│   │   │   ├── dto.go                        # BonusDTO, AwardBonusDTO
│   │   │   └── mapper.go                     # Bonus ⇄ DTO mappers
│   │   │
│   │   # ===== EXPENSE SERVICES =====
│   │   ├── expense/
│   │   │   ├── service.go                    # Submit/approve/reimburse expenses
│   │   │   ├── commands.go                   # SubmitExpense, ApproveExpense, RejectExpense, ReimburseExpense
│   │   │   ├── queries.go                    # GetExpenses, GetExpenseByID
│   │   │   ├── validators.go                 # Policy caps, categories, receipt validation
│   │   │   ├── dto.go                        # ExpenseDTO, ReceiptDTO
│   │   │   └── mapper.go                     # Expense ⇄ DTO mappers
│   │   │
│   │   # ===== PAYMENT SCHEDULE SERVICES =====
│   │   ├── payment_schedule/
│   │   │   ├── service.go                    # Create/update schedules, run due payments
│   │   │   ├── commands.go                   # CreateSchedule, UpdateSchedule, RunDuePayments
│   │   │   ├── queries.go                    # GetSchedules, GetNextDue
│   │   │   ├── validators.go                 # Conflicts, proration, disabled checks
│   │   │   ├── dto.go                        # PaymentScheduleDTO, FrequencyDTO
│   │   │   └── mapper.go                     # Schedule ⇄ DTO mappers
│   │   │
│   │   # ===== REMINDER SERVICES =====
│   │   ├── reminder/
│   │   │   ├── service.go                    # Trigger & escalate reminders (dunning support)
│   │   │   ├── commands.go                   # TriggerReminder, EscalateReminder
│   │   │   ├── queries.go                    # GetReminders, GetReminderTemplates
│   │   │   ├── validators.go                 # Channel availability, throttle
│   │   │   ├── dto.go                        # ReminderDTO, ReminderTemplateDTO
│   │   │   └── mapper.go                     # Reminder ⇄ DTO mappers
│   │   │
│   │   # ===== INSURANCE SERVICES =====
│   │   ├── insurance/
│   │   │   ├── service.go                    # Apply policy, file/approve claims, coordinate payouts
│   │   │   ├── commands.go                   # ApplyInsurance, FileClaim, ApproveClaim, PayInsuranceClaim
│   │   │   ├── queries.go                    # GetInsurance, GetClaims, GetCoverage
│   │   │   ├── validators.go                 # Policy eligibility, coverage limits
│   │   │   ├── dto.go                        # InsuranceDTO, InsuranceClaimDTO
│   │   │   └── mapper.go                     # Insurance ⇄ DTO mappers
│   │   │
│   │   # ===== TAX FORM SERVICES =====
│   │   ├── tax_form/
│   │   │   ├── service.go                    # Submit/verify forms; link to tax repo
│   │   │   ├── commands.go                   # SubmitTaxForm, VerifyTaxForm
│   │   │   ├── queries.go                    # GetTaxForms, GetTaxFormByID
│   │   │   ├── validators.go                 # Form completeness, identity checks
│   │   │   ├── dto.go                        # TaxFormDTO
│   │   │   └── mapper.go                     # Tax form ⇄ DTO mappers
│   │   │
│   │   # ===== PAYROLL SERVICES =====
│   │   ├── payroll/
│   │   │   ├── service.go                    # Process payroll; apply withholdings & deductions
│   │   │   ├── commands.go                   # ProcessPayroll, WithholdTax
│   │   │   ├── queries.go                    # GetPayroll, GetPayPeriods
│   │   │   ├── validators.go                 # Classification rules, closed periods
│   │   │   ├── dto.go                        # PayrollDTO, PayPeriodDTO
│   │   │   └── mapper.go                     # Payroll ⇄ DTO mappers
│   │   │
│   │   # ===== CURRENCY SERVICES =====
│   │   ├── currency/
│   │   │   ├── service.go                    # Manage preferred currency & conversions
│   │   │   ├── commands.go                   # LockRate, ConvertNow, SetPreferredCurrency
│   │   │   ├── queries.go                    # GetPreferredCurrency, GetRateLock
│   │   │   ├── validators.go                 # Lock windows, precision, supported pairs
│   │   │   ├── dto.go                        # CurrencyPreferenceDTO, RateLockDTO
│   │   │   └── mapper.go                     # Currency ⇄ DTO mappers
│   │   │
│   │   # ===== BANK ACCOUNT SERVICES =====
│   │   └── bank_account/                     # 🆕 Bank account app module
│   │       ├── service.go                    # CRUD bank accounts, verify, set default payout method
│   │       ├── commands.go                   # AddBankAccount, VerifyBankAccount, SetDefaultBankAccount, RemoveBankAccount
│   │       ├── queries.go                    # GetBankAccount, ListBankAccounts, GetDefaultBankAccount
│   │       ├── validators.go                 # Validate routing number, account number, bank details
│   │       ├── dto.go                        # BankAccountDTO, AddBankAccountDTO, VerificationDTO
│   │       └── mapper.go                     # BankAccount ⇄ DTO mappers
│   │
│   # ──────────────────────────────────────────────────────────────────────────────────
│   # 🌐 INTERFACES LAYER (HTTP/API - Load Fifth)
│   # ──────────────────────────────────────────────────────────────────────────────────
│   └── interfaces/
│       └── http/
│           │
│           # ===== API VERSIONING =====
│           ├── v1/                           # 🆕 Versioned API surface (all handlers/routes under /v1)
│           │   │
│           │   # ===== HANDLERS =====
│           │   ├── handlers/
│           │   │   │
│           │   │   # ===== CORE WALLET & LEDGER HANDLERS =====
│           │   │   ├── wallet_handler.go         # GET /v1/wallets/:id, POST /v1/wallets/:id/deposit, /withdraw, /transfer
│           │   │   ├── transaction_handler.go    # GET /v1/transactions, GET /v1/transactions/:id
│           │   │   ├── ledger_journal_handler.go # POST /v1/ledger/entry, POST /v1/ledger/adjustment, GET /v1/ledger/audit
│           │   │   │
│           │   │   # ===== PAYMENT & ESCROW HANDLERS =====
│           │   │   ├── payment_handler.go        # POST /v1/payments, POST /v1/payments/:id/capture, /void, /refund, GET /v1/payments/:id
│           │   │   ├── escrow_handler.go         # POST /v1/escrow/hold, POST /v1/escrow/release, POST /v1/escrow/partial-release
│           │   │   ├── payout_handler.go         # POST /v1/payouts, GET /v1/payouts/:id, POST /v1/payouts/:id/cancel
│           │   │   │
│           │   │   # ===== INVOICE & FEE HANDLERS =====
│           │   │   ├── invoice_handler.go        # POST /v1/invoices, GET /v1/invoices/:id, POST /v1/invoices/:id/send, /pay
│           │   │   ├── fee_handler.go            # GET /v1/fees/calculate
│           │   │   ├── fee_v2_handler.go         # GET /v1/fees/v2/effective, POST /v1/fees/v2/coupon
│           │   │   │
│           │   │   # ===== REFUND & TAX HANDLERS =====
│           │   │   ├── refund_handler.go         # POST /v1/refunds, POST /v1/refunds/:id/cancel
│           │   │   ├── tax_handler.go            # GET /v1/tax/forms, POST /v1/tax/1099k
│           │   │   │
│           │   │   # ===== FX & RISK HANDLERS =====
│           │   │   ├── fx_handler.go             # GET /v1/fx/rate?base=&quote=, POST /v1/fx/rate
│           │   │   ├── risk_handler.go           # POST /v1/risk/reserve, POST /v1/risk/hold, GET /v1/risk/holds
│           │   │   │
│           │   │   # ===== PROTECTION & FEE UPDATE HANDLERS =====
│           │   │   ├── protection_plan_handler.go    # POST /v1/protection/apply, POST /v1/protection/claims, GET /v1/protection/:contractId
│           │   │   ├── fee_update_handler.go         # POST /v1/fees/updates, POST /v1/fees/activate, GET /v1/fees/active
│           │   │   │
│           │   │   # ===== INTERNATIONAL PAYMENT HANDLERS =====
│           │   │   ├── international_payment_handler.go # POST /v1/intl/payments, GET /v1/intl/payments/:id
│           │   │   │
│           │   │   # ===== BONUS & EXPENSE HANDLERS =====
│           │   │   ├── bonus_handler.go              # POST /v1/bonus/award, GET /v1/bonus?contractId=
│           │   │   ├── expense_handler.go            # POST /v1/expenses, PUT /v1/expenses/:id/approve, PUT /v1/expenses/:id/reject
│           │   │   │
│           │   │   # ===== SCHEDULE & REMINDER HANDLERS =====
│           │   │   ├── payment_schedule_handler.go   # POST /v1/schedules, PUT /v1/schedules/:id, POST /v1/schedules/run-due
│           │   │   ├── reminder_handler.go           # POST /v1/reminders/trigger, POST /v1/reminders/escalate, GET /v1/reminders
│           │   │   │
│           │   │   # ===== INSURANCE & TAX FORM HANDLERS =====
│           │   │   ├── insurance_handler.go          # POST /v1/insurance/apply, POST /v1/insurance/claims, GET /v1/insurance/:contractId
│           │   │   ├── tax_form_handler.go           # POST /v1/tax/forms, GET /v1/tax/forms?contractId=
│           │   │   │
│           │   │   # ===== PAYROLL & BANK ACCOUNT HANDLERS =====
│           │   │   ├── payroll_handler.go            # POST /v1/payroll/process, GET /v1/payroll?contractId=
│           │   │   ├── bank_account_handler.go       # 🆕 POST /v1/bank-accounts, GET /v1/bank-accounts/:id, DELETE /v1/bank-accounts/:id
│           │   │   │
│           │   │   # ===== WEBHOOK HANDLERS =====
│           │   │   ├── webhook_handler.go        # POST /v1/webhooks/stripe, POST /v1/webhooks/paypal
│           │   │   │                             # 🆕 Includes signature verification middleware
│           │   │   │
│           │   │   # ===== HEALTH HANDLERS =====
│           │   │   └── health_handler.go         # 🆕 GET /v1/healthz/live, GET /v1/healthz/ready
│           │   │
│           │   # ===== ROUTES =====
│           │   ├── routes/
│           │   │   │
│           │   │   # ===== CORE ROUTES =====
│           │   │   ├── wallet_routes.go          # /v1/wallets/*
│           │   │   ├── transaction_routes.go     # /v1/transactions/*
│           │   │   ├── ledger_journal_routes.go  # /v1/ledger/*
│           │   │   │
│           │   │   # ===== PAYMENT & ESCROW ROUTES =====
│           │   │   ├── payment_routes.go         # /v1/payments/*
│           │   │   ├── escrow_routes.go          # /v1/escrow/*
│           │   │   ├── payout_routes.go          # /v1/payouts/*
│           │   │   │
│           │   │   # ===== INVOICE & FEE ROUTES =====
│           │   │   ├── invoice_routes.go         # /v1/invoices/*
│           │   │   ├── fee_routes.go             # /v1/fees/*
│           │   │   ├── fee_v2_routes.go          # /v1/fees/v2/*
│           │   │   │
│           │   │   # ===== REFUND & TAX ROUTES =====
│           │   │   ├── refund_routes.go          # /v1/refunds/*
│           │   │   ├── tax_routes.go             # /v1/tax/*
│           │   │   │
│           │   │   # ===== FX & RISK ROUTES =====
│           │   │   ├── fx_routes.go              # /v1/fx/*
│           │   │   ├── risk_routes.go            # /v1/risk/*
│           │   │   │
│           │   │   # ===== PROTECTION & FEE UPDATE ROUTES =====
│           │   │   ├── protection_plan_routes.go     # /v1/protection/*
│           │   │   ├── fee_update_routes.go          # /v1/fees/updates/*
│           │   │   │
│           │   │   # ===== INTERNATIONAL PAYMENT ROUTES =====
│           │   │   ├── international_payment_routes.go # /v1/intl/*
│           │   │   │
│           │   │   # ===== BONUS & EXPENSE ROUTES =====
│           │   │   ├── bonus_routes.go               # /v1/bonus/*
│           │   │   ├── expense_routes.go             # /v1/expenses/*
│           │   │   │
│           │   │   # ===== SCHEDULE & REMINDER ROUTES =====
│           │   │   ├── payment_schedule_routes.go    # /v1/schedules/*
│           │   │   ├── reminder_routes.go            # /v1/reminders/*
│           │   │   │
│           │   │   # ===== INSURANCE & TAX FORM ROUTES =====
│           │   │   ├── insurance_routes.go           # /v1/insurance/*
│           │   │   ├── tax_form_routes.go            # /v1/tax/forms/*
│           │   │   │
│           │   │   # ===== PAYROLL & BANK ACCOUNT ROUTES =====
│           │   │   ├── payroll_routes.go             # /v1/payroll/*
│           │   │   ├── bank_account_routes.go        # 🆕 /v1/bank-accounts/*
│           │   │   │
│           │   │   # ===== WEBHOOK ROUTES =====
│           │   │   ├── webhook_routes.go         # 🆕 /v1/webhooks/*
│           │   │   │
│           │   │   # ===== HEALTH ROUTES =====
│           │   │   └── health_routes.go          # 🆕 /v1/healthz/*
│           │   │
│           │   # ===== OPENAPI SPEC =====
│           │   └── openapi/
│           │       ├── openapi.yaml              # 🆕 OpenAPI 3.0 spec (served as /v1/openapi)
│           │       └── generator.go              # 🆕 Serves /v1/swagger and /v1/openapi.json (dev only)
│           │
│           # ===== MIDDLEWARE =====
│           ├── middleware/
│           │   ├── requestid.go              # 🆕 Wraps platform-shared request id
│           │   ├── logging.go                # 🆕 Wraps platform-shared logging
│           │   ├── recovery.go               # 🆕 Wraps platform-shared recovery
│           │   ├── tracing.go                # 🆕 Wraps platform-shared otel middleware
│           │   ├── auth.go                   # Uses pkg/auth (Keycloak/OIDC token verification)
│           │   ├── rbac.go                   # ♻️ Uses pkg/auth (role checks & resource scoping)
│           │   │                             # Ensure aligns with application/authz
│           │   ├── cors.go                   # Uses platform-shared/ginx (CORS config)
│           │   ├── rate_limit.go             # Rate limiting (per-IP/per-user, burst)
│           │   ├── idempotency.go            # Uses platform-shared/idempotency (CRITICAL for payments!)
│           │   └── webhook_signature.go      # 🆕 Webhook signature verification (Stripe/PayPal)
│           │
│           # ===== ERROR PRESENTERS =====
│           ├── presenters/
│           │   └── errors.go                 # 🆕 Maps domain/service errors → HTTP status & problem+json
│           │
│           # ===== ROUTER =====
│           └── router.go                     # ♻️ Uses platform-shared/ginx; mounts v1 route registrars & middleware
│
├── db/                                       # 🆕 Developer-friendly entrypoint for SQL
│   └── migrations/                           # 🆕 Symlink or mirror of internal/.../migrations (optional)
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go                         # Base error helpers (wrap, with code/context)
│   │   ├── codes.go                          # Error codes: INSUFFICIENT_FUNDS, PAYMENT_FAILED, ESCROW_LOCK, etc.
│   │   └── payment_errors.go                 # Payment-specific error types (GatewayDecline, 3DSRequired, RiskHold)
│   │
│   ├── logger/                               # ❌ REMOVED
│   │   └── README.md                         # Placeholder: use platform-shared logging
│   │
│   ├── utils/
│   │   ├── validator.go                      # Common validation helpers (UUIDs, enums, ranges)
│   │   ├── currency.go                       # Currency conversion helpers (symbols, ISO codes)
│   │   ├── decimal.go                        # Decimal math for money (fixed-point ops, rounding)
│   │   ├── encryption.go                     # Encryption utilities (field-level encryption, PCI scope)
│   │   └── etag.go                           # 🆕 ETag generation helpers (pairs with infrastructure/http/etag)
│   │
│   └── constants/
│       ├── events.go                         # ❌ REMOVED (use platform-shared constants)
│       ├── topics.go                         # ❌ REMOVED (use platform-shared topics)
│       ├── currencies.go                     # Currency constants (USD, EUR, GBP, supported list)
│       └── payment_methods.go                # Payment method constants (Stripe, PayPal, BankTransfer, Wallet)
│
├── config/                                   # Runtime config profiles
│   ├── default.yaml                          # Base config (safe defaults)
│   ├── dev.yaml                              # Development overrides (local gateways, debug flags)
│   └── prod.yaml                             # Production overrides (timeouts, pools, feature flags)
│
├── dapr/                                     # Dapr components
│   ├── local/
│   │   ├── pubsub.yaml                       # Local pub/sub (scoped to financial-be)
│   │   └── statestore.yaml                   # Local state store (idempotency keys, outbox)
│   └── k8s/
│       ├── pubsub.yaml                       # Scopes: ["financial-be"] (topics for payments/escrow/payouts/etc.)
│       ├── statestore.yaml                   # State store for production (HA, TTLs)
│       └── secrets.yaml                      # References to K8s secrets (API keys, tokens)
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                   # Deployment spec (resources, env, probes)
│       ├── service.yaml                      # Service spec (ports, selectors)
│       ├── configmap.yaml                    # App config mounted as files/env
│       ├── secrets.yaml                      # Contains Stripe/PayPal API keys, DB creds (sealed/encrypted)
│       ├── hpa.yaml                          # Horizontal Pod Autoscaler (CPU/RAM/requests-per-sec)
│       ├── pdb.yaml                          # PodDisruptionBudget (availability during maintenance)
│       ├── networkpolicy.yaml                # Extra security (ingress/egress restrictions)
│       └── servicemonitor.yaml               # Prometheus ServiceMonitor (metrics scraping)
│
├── scripts/
│   ├── setup-local.sh                        # Local dev bootstrap (migrations, seeds, Dapr components)
│   ├── migrate.sh                            # 🆕 SQL migrations (dev/prod flags)
│   ├── openapi-diff.sh                       # 🆕 Fail CI on breaking API changes
│   ├── schema-diff.sh                        # 🆕 DB schema guardrail (pg_dump + mig verify)
│   ├── generate-sdks.sh                      # 🆕 Produce /sdk clients from OpenAPI
│   ├── get-secrets.sh                        # Script to fetch secrets to local env (vault/secret manager)
│   ├── seed-data.sh                          # Seed demo data (wallets, rates, fee rules)
│   ├── reconciliation.sh                     # Daily reconciliation script (bank vs. journal; reports)
│   └── sbom-sign.sh                          # 🆕 Build SBOM + cosign sign
│
├── tests/
│   ├── unit/                                 # Unit tests (domain/application)
│   ├── integration/                          # Integration tests (DB/cache/gateways)
│   ├── e2e/                                  # End-to-end tests (HTTP flows, idempotency, race conditions)
│   │   └── chaos_event_delivery_test.go      # 🆕 Simulate duplicates, reordering, delays
│   │
│   ├── reliability/                          # 🆕 Reliability tests
│   │   ├── projections_replay_test.go        # 🆕 Rebuild read models from event logs
│   │   └── outbox_dispatcher_test.go         # 🆕 At-least-once + idempotency assertions
│   │
│   └── property/                             # 🆕 Property-based tests
│       └── fee_calculation_property_test.go  # 🆕 Property/fuzz tests for fee calculation kernel
│
├── docs/
│   ├── README.md                             # Overview & getting started
│   ├── API.md                                # HTTP API surface (endpoints, payloads, examples)
│   ├── API_VERSIONING.md                     # 🆕 HTTP contract versioning & deprecation policy
│   ├── OPENAPI.md                            # 🆕 How to regenerate SDKs from OpenAPI spec
│   ├── EVENTS.md                             # Events (payment.processed, escrow.held/released, payout.processed, fee.updated, intl.payment.*)
│   ├── ARCHITECTURE.md                       # Service architecture (layers, dependencies, flows)
│   ├── MIGRATIONS.md                         # DB migrations & versioning strategy
│   ├── SCHEMA.md                             # Domain schemas (journal, fees v2, fx, protection/insurance/schedules)
│   ├── RUNBOOK.md                            # Ops runbook (alerts, dashboards, SLOs, incident steps)
│   ├── SLOS.md                               # 🆕 Target P99s, projection lag, queue delay
│   ├── OUTBOX.md                             # 🆕 Exactly-once-ish design, retries, DLQ
│   ├── CACHING.md                            # 🆕 Keys, TTLs, SWR, invalidation events
│   ├── DATA_RETENTION.md                     # 🆕 Per-domain retention windows
│   ├── ERASURE.md                            # 🆕 GDPR/CCPA erasure hooks & playbooks
│   ├── RELEASE_CHECKLIST.md                  # 🆕 Preflight (openapi diff, schema diff, migrations plan)
│   ├── NAMING.md                             # 🆕 Package snake_case; HTTP kebab-case conventions
│   ├── payment-flows.md                      # Payment flow documentation (auth/capture/refund)
│   ├── escrow-system.md                      # Escrow system documentation (holds/releases/disputes)
│   ├── fee-structure.md                      # Platform fee structure (old engine)
│   ├── fee-v2.md                             # Fee engine v2 rules & examples (tiers, coupons, locales)
│   ├── fx-and-rounding.md                    # FX, effective timestamps, rounding rules
│   ├── risk-holds.md                         # Reserves, holds, chargebacks (policies & flows)
│   ├── intl-payments.md                      # International payments flows & compliance checks
│   ├── protection-plans.md                   # Protection plans & claims flows (eligibility, payouts)
│   ├── schedules-and-reminders.md            # PaymentSchedules & Reminders (cron windows, dunning)
│   └── payroll-and-taxforms.md               # Payroll + tax_form integration notes (jurisdictions)
│
├── sdk/                                      # 🆕 Generated clients (optional but handy)
│   ├── go/                                   # 🆕 Generated Go client from OpenAPI
│   └── ts/                                   # 🆕 Generated TypeScript client from OpenAPI
│
├── .github/
│   └── workflows/
│       ├── ci.yml                            # CI pipeline (lint, test, build)
│       ├── cd.yml                            # CD pipeline (image build/push, deploy)
│       ├── contract-ci.yml                   # 🆕 OpenAPI-diff + event schema checks
│       ├── security.yml                      # 🆕 golangci-lint, govulncheck, trivy, cosign verify
│       └── load-tests.yml                    # 🆕 k6/gatling smoke against /healthz & hot paths
│
├── go.mod                                    # Imports pkg/auth, platform-shared libs, contracts/events schemas
├── go.sum                                    # Module dependency checksums
├── .env.example                              # Example environment variables (non-secret)
├── .golangci.yml                             # 🆕 Linter config (CI parity)
├── CODEOWNERS                                # 🆕 Explicit ownership per context (ease future splits)
├── Makefile                                  # Common dev tasks (run, test, lint, migrate)
├── Dockerfile                                # Container build (multi-stage; minimal runtime)
├── .dockerignore                             # Docker build context exclusions
├── .gitignore                                # Git ignore rules
└── README.md                                 # High-level service description & quickstart


```

---