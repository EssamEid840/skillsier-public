# 📦 **financial-be - Financial Management Service - Complete User Stories**

---

## **Global Conventions**

### Event Envelope (All Events)
```json
{
  "event_id": "uuid-v7",
  "event_type": "payment.processed.v1",
  "event_version": "v1",
  "aggregate_id": "payment-uuid",
  "aggregate_type": "Payment",
  "occurred_at": "ISO8601",
  "causation_id": "parent-event-id",
  "correlation_id": "trace-id",
  "metadata": {
    "user_id": "...",
    "service": "financial-be",
    "idempotency_key": "..."
  },
  "payload": { ... }
}
```

### Idempotent Write-Path
- All commands accept `idempotency_key` parameter
- Store `(idempotency_key, aggregate_id, event_version)` in `idempotency_log` table
- TTL: 7 days
- Return existing result if duplicate detected within TTL window

### Non-PII Event Payloads
- Events NEVER contain PII in payload (no card numbers, bank accounts, SSNs)
- Events reference IDs only: `payment_id`, `wallet_id`, `user_id`, `transaction_id`
- Consumers fetch PII via authenticated API calls if needed
- Example: `payment.processed.v1` contains `payment_id`, `payer_id`, `payee_id` but NOT names, emails, or full card numbers

### Folder Structure Alignment
- All domain entities map to `internal/domain/{context}/`
- All repositories map to `internal/infrastructure/persistence/postgres/{context}_repository.go`
- All services map to `internal/application/{context}/service.go`
- All handlers map to `internal/interfaces/http/v1/handlers/{context}_handler.go`
- All routes map to `internal/interfaces/http/v1/routes/{context}_routes.go`

### Events Catalog Integration
- All events published are registered in `contracts/events/financial/` catalog
- Event schemas versioned with semantic versioning (v1, v2, etc.)
- Breaking changes require new event version
- Consumers subscribe via Dapr pub/sub with scopes: `["financial-be"]`

### Caching Strategy
- Cache keys follow pattern: `financial:{user_id}:{context}:{version}`
- TTLs defined in `internal/infrastructure/cache/redis/keys.go`
- Invalidation rules map events to cache keys in `invalidation_rules.go`
- Singleflight prevents cache stampedes for hot keys (wallet balances, FX rates)

### Observability
- All commands/queries emit spans with OpenTelemetry
- Metrics tracked: P95 latency, error rate, event publish lag, payment_success_rate, payout_processing_time, escrow_hold_duration, reconciliation_accuracy
- Structured logging with correlation_id for tracing
- Health checks: `/healthz/live` (liveness), `/healthz/ready` (readiness)

### Security & Compliance
- All endpoints require JWT authentication via Keycloak
- RBAC enforced at service layer (USER, ADMIN, SYSTEM, FINANCE_ADMIN)
- PII encrypted at rest using KMS envelope encryption
- PCI-DSS compliance for payment data
- Sensitive fields redacted in logs via PII redactor
- SOC 2 compliance for financial operations

### Data Retention & Erasure
- Financial data retention: 10 years (legal/tax requirement)
- Event logs retention: 90 days (projections can replay)
- GDPR/CCPA: pseudonymize user IDs but maintain financial records
- Audit trail: immutable for 10 years

---

## **1 - CORE WALLET & BALANCE DOMAIN**

### 1.1 wallet/

#### User Stories
- As a **user**, I want to **create a wallet** so that I can hold funds on the platform.
- As a **user**, I want to **view my wallet balance** (available, pending, reserved) so that I know my funds status.
- As a **user**, I want to **deposit funds** to my wallet so that I can use them for payments.
- As a **user**, I want to **withdraw funds** from my wallet so that I can cash out.
- As a **user**, I want to **transfer funds** between wallets so that I can move money.
- As a **system**, I want to **reserve funds** for pending transactions so that double-spending is prevented.
- As a **system**, I want to **release reserved funds** after transaction completion so that balance is accurate.
- As a **system**, I want to **support multi-currency wallets** so that global users can transact.
- As a **system**, I want to **track balance history** so that all movements are auditable.
- As a **user**, I want to **set low balance alerts** so that I'm notified when funds are low.
- As a **admin**, I want to **freeze wallets** for compliance so that suspicious activity is stopped.

#### Flow
1. **CreateWalletCommand**(user_id, currency, wallet_type) → ValidateUser() | CreateWallet() | Persist() → **Outbox:** wallet.created.v1
2. **DepositFundsCommand**(wallet_id, amount, source, reference, deposited_by) → ValidateWallet() | ProcessDeposit() | UpdateBalance() | CreateTransaction() → **Outbox:** wallet.deposit.completed.v1
3. **WithdrawFundsCommand**(wallet_id, amount, destination, reference, withdrawn_by) → ValidateSufficientFunds() | ProcessWithdrawal() | UpdateBalance() | CreateTransaction() → **Outbox:** wallet.withdrawal.completed.v1
4. **TransferFundsCommand**(from_wallet_id, to_wallet_id, amount, reference, transferred_by) → ValidateBothWallets() | ValidateFunds() | ExecuteTransfer() | CreateTransactions() → **Outbox:** wallet.transfer.completed.v1
5. **ReserveFundsCommand**(wallet_id, amount, purpose, reserved_for) → ValidateAvailableFunds() | Reserve() | UpdateBalance() → **Outbox:** wallet.funds.reserved.v1
6. **ReleaseFundsCommand**(wallet_id, reservation_id, release_reason) → ValidateReservation() | Release() | UpdateBalance() → **Outbox:** wallet.funds.released.v1
7. **FreezeWalletCommand**(wallet_id, reason, frozen_by) → AuthorizeAdmin() | Freeze() | NotifyUser(communications-be) → **Outbox:** wallet.frozen.v1
8. **UnfreezeWalletCommand**(wallet_id, unfrozen_by) → AuthorizeAdmin() | Unfreeze() | NotifyUser() → **Outbox:** wallet.unfrozen.v1
9. **GetWalletBalanceQuery**(wallet_id) → AuthorizeAccess() | Fetch() → WalletBalanceDTO
10. **GetWalletHistoryQuery**(wallet_id, date_range, filters) → ApplyFilters() | Paginate() → WalletHistoryDTO
11. **SetLowBalanceAlertCommand**(wallet_id, threshold, alert_method) → AuthorizeOwner() | Set() → **Outbox:** wallet.alert.configured.v1

#### Projections
- wallet_read
- wallet_balance_read
- wallet_history_read
- wallet_reservations_read
- wallet_alerts_read

#### Events Published
- wallet.created.v1
- wallet.deposit.completed.v1
- wallet.withdrawal.completed.v1
- wallet.transfer.completed.v1
- wallet.funds.reserved.v1
- wallet.funds.released.v1
- wallet.frozen.v1
- wallet.unfrozen.v1
- wallet.alert.triggered.v1
- wallet.balance.low.v1

#### Events Consumed
- user.created.v1 (auto-create wallet)
- user.verified.v1 (enable full wallet features)
- payment.processed.v1 (update balance)

#### RBAC/SLO
- **RBAC:** USER (create/deposit/withdraw/transfer/view own), ADMIN (freeze/unfreeze/view all), SYSTEM (reserve/release)
- **SLO:** P95 < 200ms (read), P95 < 300ms (write), P99 < 500ms (transfer)

---

### 1.2 balance/

#### User Stories
- As a **user**, I want to **see breakdown** (available, pending, reserved) so that I understand my funds.
- As a **system**, I want to **atomic balance updates** so that race conditions don't occur.
- As a **system**, I want to **track balance snapshots** daily so that reporting is accurate.
- As a **user**, I want to **view balance chart** over time so that I can see trends.

#### Flow
1. **UpdateBalanceCommand**(wallet_id, adjustment_type, amount) → AcquireLock() | ValidateBalance() | ApplyAdjustment() | ReleaseLock() → **Outbox:** balance.updated.v1
2. **CreateBalanceSnapshotCommand**(wallet_id, snapshot_date) → CaptureSnapshot() | Persist() → **Outbox:** balance.snapshot.created.v1
3. **GetBalanceBreakdownQuery**(wallet_id) → Fetch() → BalanceBreakdownDTO
4. **GetBalanceChartQuery**(wallet_id, date_range, interval) → AggregateSnapshots() → BalanceChartDTO

#### Projections
- balance_breakdown_read
- balance_snapshots_read
- balance_trends_read

#### Events Published
- balance.updated.v1
- balance.snapshot.created.v1

#### RBAC/SLO
- **RBAC:** USER (view own), SYSTEM (update), ADMIN (view all)
- **SLO:** P95 < 150ms (read), P95 < 250ms (update - includes locking)

---

## **2 - TRANSACTION & LEDGER DOMAIN**

### 2.1 transaction/

#### User Stories
- As a **user**, I want to **view all transactions** so that I can audit my financial activity.
- As a **system**, I want to **create transactions** with double-entry bookkeeping so that accounting is accurate.
- As a **system**, I want to **reverse transactions** when errors occur so that corrections are possible.
- As a **system**, I want to **reconcile transactions** with bank statements so that discrepancies are found.
- As a **user**, I want to **search transactions** by type, date, amount so that finding is easy.
- As a **user**, I want to **export transactions** to CSV/PDF so that records can be saved.
- As a **system**, I want to **link transactions** to contracts/milestones so that context is preserved.

#### Flow
1. **CreateTransactionCommand**(wallet_id, amount, type, reference, metadata) → ValidateWallet() | CreateLedgerEntries() | Persist() | UpdateBalance() → **Outbox:** transaction.created.v1
2. **ReverseTransactionCommand**(transaction_id, reason, reversed_by) → AuthorizeAdmin() | ValidateReversible() | CreateReversalTransaction() | UpdateBalances() → **Outbox:** transaction.reversed.v1
3. **ReconcileTransactionsCommand**(start_date, end_date, bank_statement_file) → ParseStatement() | MatchTransactions() | FlagDiscrepancies() | GenerateReport() → **Outbox:** transactions.reconciled.v1
4. **GetTransactionQuery**(transaction_id) → AuthorizeAccess() | Fetch() → TransactionDTO
5. **SearchTransactionsQuery**(wallet_id, filters, pagination) → ApplyFilters() | Search() → TransactionListDTO
6. **ExportTransactionsCommand**(wallet_id, date_range, format, exported_by) → FetchTransactions() | GenerateExport() | Upload(storage-be) → **Outbox:** transactions.exported.v1

#### Projections
- transaction_read
- transaction_ledger_read
- transaction_reconciliation_read

#### Events Published
- transaction.created.v1
- transaction.reversed.v1
- transactions.reconciled.v1
- transactions.exported.v1
- transaction.reconciliation.discrepancy.v1

#### RBAC/SLO
- **RBAC:** USER (view own/export), ADMIN (reverse/reconcile/view all), SYSTEM (create)
- **SLO:** P95 < 180ms (read), P95 < 300ms (create), P95 < 500ms (reverse), P95 < 2000ms (reconcile)

---

### 2.2 ledger_journal/

#### User Stories
- As a **system**, I want to **maintain append-only ledger** so that financial history is immutable.
- As a **system**, I want to **validate debits equal credits** so that accounting is balanced.
- As a **system**, I want to **chain ledger entries** with cryptographic hashes so that tampering is detectable.
- As a **admin**, I want to **create manual adjustments** with approval workflow so that corrections are controlled.
- As a **auditor**, I want to **audit ledger trail** so that compliance is verified.
- As a **system**, I want to **verify ledger integrity** so that data corruption is detected.

#### Flow
1. **AppendJournalEntryCommand**(debit_account, credit_account, amount, reference) → ValidateBalance() | ComputeHash() | Append() → **Outbox:** journal.entry.recorded.v1
2. **CreateAdjustmentCommand**(account, amount, reason, created_by) → AuthorizeMaker() | CreateProposal() → **Outbox:** adjustment.proposed.v1
3. **ApproveAdjustmentCommand**(adjustment_id, approved_by) → AuthorizeChecker() | ApplyAdjustment() | AppendJournalEntry() → **Outbox:** adjustment.approved.v1
4. **RejectAdjustmentCommand**(adjustment_id, reason, rejected_by) → AuthorizeChecker() | Reject() → **Outbox:** adjustment.rejected.v1
5. **GetAuditTrailQuery**(account, date_range) → Fetch() | VerifyHashChain() → AuditTrailDTO
6. **VerifyLedgerIntegrityCommand**() → VerifyHashChain() | CheckBalances() | GenerateReport() → **Outbox:** ledger.integrity.verified.v1

#### Projections
- ledger_journal_read
- ledger_adjustments_read
- ledger_audit_trail_read

#### Events Published
- journal.entry.recorded.v1
- adjustment.proposed.v1
- adjustment.approved.v1
- adjustment.rejected.v1
- ledger.integrity.verified.v1
- ledger.hash.mismatch.detected.v1

#### RBAC/SLO
- **RBAC:** SYSTEM (append), MAKER (propose adjustment), CHECKER (approve/reject), AUDITOR (audit trail/verify integrity)
- **SLO:** P95 < 200ms (append), P95 < 300ms (adjustment operations), P95 < 1000ms (integrity check)

---

## **3 - PAYMENT PROCESSING DOMAIN**

### 3.1 payment/

#### User Stories
- As a **client**, I want to **process payments** for milestones/timesheets so that freelancers get paid.
- As a **system**, I want to **support multiple payment methods** (credit card, bank transfer, PayPal, wallet) so that flexibility is provided.
- As a **system**, I want to **integrate with payment gateways** (Stripe, PayPal) so that processing is handled.
- As a **system**, I want to **retry failed payments** with backoff so that transient failures are handled.
- As a **user**, I want to **view payment status** so that I can track processing.
- As a **system**, I want to **handle 3D Secure** for card payments so that authentication is completed.
- As a **system**, I want to **calculate fees** (platform, gateway, tax) so that deductions are accurate.
- As a **system**, I want to **send payment receipts** so that users have records.

#### Flow
1. **ProcessPaymentCommand**(payer_id, payee_id, amount, currency, payment_method, reference) → ValidateParties() | CalculateFees() | ProcessViaGateway() | CreateTransaction() | UpdateWallets() → **Outbox:** payment.processed.v1
2. **AuthorizePaymentCommand**(payment_id, payment_method_token) → Authorize(gateway) | HoldFunds() → **Outbox:** payment.authorized.v1
3. **CapturePaymentCommand**(payment_id) → Capture(gateway) | CompleteFundsTransfer() → **Outbox:** payment.captured.v1
4. **RetryFailedPaymentCommand**(payment_id, retry_count) → ValidateRetryable() | ProcessPayment() → **Outbox:** payment.retried.v1
5. **Handle3DSChallengeCommand**(payment_id, challenge_response) → Verify3DS(gateway) | CompletePayment() → **Outbox:** payment.3ds.completed.v1
6. **GetPaymentQuery**(payment_id) → AuthorizeAccess() | Fetch() → PaymentDTO
7. **ListPaymentsQuery**(user_id, filters) → ApplyFilters() → PaymentListDTO
8. **GenerateReceiptCommand**(payment_id) → FetchPaymentDetails() | GeneratePDF() | Send(communications-be) → **Outbox:** payment.receipt.generated.v1

#### Projections
- payment_read
- payment_status_read
- payment_history_read
- payment_fees_read

#### Events Published
- payment.initiated.v1
- payment.authorized.v1
- payment.captured.v1
- payment.processed.v1
- payment.failed.v1
- payment.retried.v1
- payment.3ds.completed.v1
- payment.receipt.generated.v1

#### Events Consumed
- milestone.approved.v1 (trigger payment)
- timesheet.approved.v1 (trigger payment)
- contract.terminated.v1 (handle refunds)

#### RBAC/SLO
- **RBAC:** CLIENT (initiate payment), FREELANCER (view payments received), SYSTEM (process/retry), ADMIN (view all)
- **SLO:** P95 < 500ms (authorization), P95 < 800ms (capture), P95 < 1000ms (full payment processing), P99 < 2000ms

---

### 3.2 payment_method/

#### User Stories
- As a **user**, I want to **add payment methods** (cards, bank accounts) so that I can pay/receive funds.
- As a **user**, I want to **verify payment methods** so that they're activated.
- As a **user**, I want to **set default payment method** so that checkout is faster.
- As a **user**, I want to **delete payment methods** so that unused ones are removed.
- As a **system**, I want to **tokenize payment details** via gateway so that PCI compliance is maintained.
- As a **system**, I want to **validate payment methods** before use so that failures are prevented.

#### Flow
1. **AddPaymentMethodCommand**(user_id, method_type, method_details, added_by) → Tokenize(gateway) | ValidateMethod() | Persist() → **Outbox:** payment_method.added.v1
2. **VerifyPaymentMethodCommand**(method_id, verification_data) → VerifyViaGateway() | ActivateMethod() → **Outbox:** payment_method.verified.v1
3. **SetDefaultPaymentMethodCommand**(user_id, method_id) → AuthorizeOwner() | UpdateDefault() → **Outbox:** payment_method.default.set.v1
4. **DeletePaymentMethodCommand**(method_id, deleted_by) → AuthorizeOwner() | RemoveFromGateway() | SoftDelete() → **Outbox:** payment_method.deleted.v1
5. **GetPaymentMethodsQuery**(user_id) → AuthorizeOwner() | Fetch() → PaymentMethodListDTO
6. **GetPaymentMethodQuery**(method_id) → AuthorizeOwner() | Fetch() → PaymentMethodDTO

#### Projections
- payment_methods_read

#### Events Published
- payment_method.added.v1
- payment_method.verified.v1
- payment_method.default.set.v1
- payment_method.deleted.v1
- payment_method.expired.v1

#### RBAC/SLO
- **RBAC:** USER (add/verify/set default/delete own), ADMIN (view all)
- **SLO:** P95 < 300ms (add - includes gateway tokenization), P95 < 250ms (verify), P95 < 150ms (read)

---

## **4 - ESCROW MANAGEMENT DOMAIN**

### 4.1 escrow/

#### User Stories
- As a **client**, I want to **fund escrow** for contract so that freelancer has payment assurance.
- As a **system**, I want to **hold funds in escrow** until milestone approval so that both parties are protected.
- As a **system**, I want to **release escrow** upon approval so that freelancer receives payment.
- As a **system**, I want to **refund escrow** on contract termination so that client recovers funds.
- As a **system**, I want to **partial release escrow** for milestone-based contracts so that progressive payment is supported.
- As a **admin**, I want to **manually release escrow** in disputes so that resolutions can be enforced.
- As a **system**, I want to **track escrow holds** so that funds are properly accounted.

#### Flow
1. **CreateEscrowCommand**(contract_id, amount, currency, created_by) → ValidateContract() | CreateEscrowAccount() | Persist() → **Outbox:** escrow.created.v1
2. **FundEscrowCommand**(escrow_id, amount, payment_method_id, funded_by) → ValidatePaymentMethod() | ProcessPayment(gateway) | HoldFunds() → **Outbox:** escrow.funded.v1
3. **PlaceEscrowHoldCommand**(escrow_id, milestone_id, hold_amount) → ValidateFunds() | Reserve() → **Outbox:** escrow.hold.placed.v1
4. **ReleaseEscrowCommand**(escrow_id, milestone_id, release_amount, released_by) → ValidateHold() | ReleaseToFreelancer() | ProcessPayout() → **Outbox:** escrow.released.v1
5. **PartialReleaseEscrowCommand**(escrow_id, release_amount, pro_rata_basis, released_by) → CalculateProRata() | ReleasePartial() → **Outbox:** escrow.partial.released.v1
6. **RefundEscrowCommand**(escrow_id, refund_amount, reason, refunded_by) → CalculateRefund() | ProcessRefund() | ReleaseHolds() → **Outbox:** escrow.refunded.v1
7. **ManualReleaseEscrowCommand**(escrow_id, amount, admin_reason, released_by) → AuthorizeAdmin() | Release() → **Outbox:** escrow.manual.released.v1
8. **GetEscrowStatusQuery**(contract_id) → Fetch() → EscrowStatusDTO
9. **GetEscrowHistoryQuery**(escrow_id) → Fetch() → EscrowHistoryDTO

#### Projections
- escrow_read
- escrow_holds_read
- escrow_history_read
- escrow_status_read

#### Events Published
- escrow.created.v1
- escrow.funded.v1
- escrow.hold.placed.v1
- escrow.released.v1
- escrow.partial.released.v1
- escrow.refunded.v1
- escrow.manual.released.v1

#### Events Consumed
- contract.activated.v1 (create escrow)
- milestone.approved.v1 (release hold)
- contract.terminated.v1 (refund escrow)
- dispute.resolved.v1 (manual release)

#### RBAC/SLO
- **RBAC:** CLIENT (fund), SYSTEM (hold/release), ADMIN (manual release), PUBLIC (view own)
- **SLO:** P95 < 400ms (fund - includes payment processing), P95 < 250ms (hold), P95 < 350ms (release), P95 < 150ms (read)

---

## **5 - PAYOUT PROCESSING DOMAIN**

### 5.1 payout/

#### User Stories
- As a **freelancer**, I want to **request payouts** so that I can withdraw earnings.
- As a **freelancer**, I want to **choose payout method** (bank transfer, PayPal, Wise, crypto) so that I receive funds conveniently.
- As a **system**, I want to **batch payouts** by method/currency so that processing is efficient.
- As a **system**, I want to **validate minimum payout thresholds** so that small payouts are prevented.
- As a **system**, I want to **calculate payout fees** so that costs are transparent.
- As a **system**, I want to **process payouts on schedule** (daily/weekly/instant) so that timing is predictable.
- As a **freelancer**, I want to **track payout status** so that I know when funds arrive.
- As a **admin**, I want to **cancel payouts** if needed so that errors can be corrected.

#### Flow
1. **RequestPayoutCommand**(user_id, amount, payout_method_id, requested_by) → ValidateBalance() | ValidateMinimum() | CalculateFees() | QueuePayout() → **Outbox:** payout.requested.v1
2. **BatchPayoutsCommand**(payout_ids[], batch_date) → GroupByMethod() | ProcessBatch(gateway) | UpdateStatuses() → **Outbox:** payout.batch.processed.v1
3. **ProcessPayoutCommand**(payout_id) → ValidateStatus() | ProcessViaGateway() | UpdateWallet() | GenerateReceipt() → **Outbox:** payout.processed.v1
4. **CancelPayoutCommand**(payout_id, reason, cancelled_by) → ValidateStatus() | CancelInGateway() | RefundFees() | UpdateStatus() → **Outbox:** payout.cancelled.v1
5. **RetryFailedPayoutCommand**(payout_id, retry_count) → ValidateRetryable() | ProcessPayout() → **Outbox:** payout.retried.v1
6. **GetPayoutQuery**(payout_id) → AuthorizeAccess() | Fetch() → PayoutDTO
7. **ListPayoutsQuery**(user_id, filters) → ApplyFilters() → PayoutListDTO
8. **GetPayoutScheduleQuery**(user_id) → CalculateNextPayout() → PayoutScheduleDTO

#### Projections
- payout_read
- payout_queue_read
- payout_batch_read
- payout_schedule_read

#### Events Published
- payout.requested.v1
- payout.queued.v1
- payout.batch.processed.v1
- payout.processed.v1
- payout.failed.v1
- payout.cancelled.v1
- payout.retried.v1

#### Events Consumed
- payment.processed.v1 (update available balance)
- escrow.released.v1 (trigger payout)

#### RBAC/SLO
- **RBAC:** FREELANCER (request/view own), ADMIN (cancel/view all/process batches), SYSTEM (batch/process)
- **SLO:** P95 < 300ms (request), P95 < 800ms (process individual), P95 < 5000ms (batch), P95 < 150ms (read)

---

### 5.2 payout_method/

#### User Stories
- As a **freelancer**, I want to **add payout destinations** so that I can receive funds.
- As a **freelancer**, I want to **verify payout methods** so that they're activated.
- As a **freelancer**, I want to **set default payout method** so that withdrawals are easier.
- As a **system**, I want to **validate payout details** (bank routing, PayPal email) so that transfers succeed.
- As a **system**, I want to **comply with KYC/KYB** before enabling payouts so that regulations are followed.

#### Flow
1. **AddPayoutMethodCommand**(user_id, method_type, method_details, added_by) → ValidateDetails() | VerifyKYC() | Persist() → **Outbox:** payout_method.added.v1
2. **VerifyPayoutMethodCommand**(method_id, verification_data) → SendMicroDeposit() | VerifyAmount() | Activate() → **Outbox:** payout_method.verified.v1
3. **SetDefaultPayoutMethodCommand**(user_id, method_id) → AuthorizeOwner() | UpdateDefault() → **Outbox:** payout_method.default.set.v1
4. **DeletePayoutMethodCommand**(method_id, deleted_by) → AuthorizeOwner() | SoftDelete() → **Outbox:** payout_method.deleted.v1
5. **GetPayoutMethodsQuery**(user_id) → AuthorizeOwner() | Fetch() → PayoutMethodListDTO

#### Projections
- payout_methods_read

#### Events Published
- payout_method.added.v1
- payout_method.verified.v1
- payout_method.default.set.v1
- payout_method.deleted.v1

#### RBAC/SLO
- **RBAC:** FREELANCER (add/verify/set default/delete own), ADMIN (view all)
- **SLO:** P95 < 250ms (add), P95 < 300ms (verify), P95 < 150ms (read)

---

## **6 - INVOICE MANAGEMENT DOMAIN**

### 6.1 invoice/

#### User Stories
- As a **system**, I want to **generate invoices** for milestones/timesheets so that billing is documented.
- As a **freelancer**, I want to **view invoices** so that I can track what was billed.
- As a **client**, I want to **pay invoices** so that I fulfill obligations.
- As a **system**, I want to **send invoice reminders** when overdue so that collection is improved.
- As a **system**, I want to **apply discounts/coupons** to invoices so that promotions are honored.
- As a **system**, I want to **calculate taxes** (VAT, sales tax) so that invoices are compliant.
- As a **system**, I want to **generate invoice PDFs** so that records can be downloaded.

#### Flow
1. **GenerateInvoiceCommand**(contract_id, line_items[], due_date, generated_by) → CalculateTotals() | CalculateTaxes() | AssignInvoiceNumber() | Persist() | GeneratePDF() → **Outbox:** invoice.generated.v1
2. **SendInvoiceCommand**(invoice_id) → FetchInvoice() | Send(communications-be) | MarkSent() → **Outbox:** invoice.sent.v1
3. **MarkInvoicePaidCommand**(invoice_id, payment_id, paid_by) → ValidatePayment() | MarkPaid() | UpdateAccounting() → **Outbox:** invoice.paid.v1
4. **CancelInvoiceCommand**(invoice_id, reason, cancelled_by) → ValidateStatus() | Cancel() | RefundPayment() → **Outbox:** invoice.cancelled.v1
5. **ApplyDiscountCommand**(invoice_id, discount_code) → ValidateDiscount() | ApplyDiscount() | Recalculate() → **Outbox:** invoice.discount.applied.v1
6. **SendReminderCommand**(invoice_id) → CheckOverdue() | Send(communications-be) | LogReminder() → **Outbox:** invoice.reminder.sent.v1
7. **GetInvoiceQuery**(invoice_id) → AuthorizeAccess() | Fetch() | GeneratePDF() → InvoiceDTO
8. **ListInvoicesQuery**(user_id, filters) → ApplyFilters() → InvoiceListDTO
9. **DownloadInvoicePDFQuery**(invoice_id) → AuthorizeAccess() | GeneratePDF() | PresignURL(storage-be) → InvoicePDFURLDTO

#### Projections
- invoice_read
- invoice_line_items_read
- invoice_payment_status_read
- invoice_reminders_read

#### Events Published
- invoice.generated.v1
- invoice.sent.v1
- invoice.paid.v1
- invoice.overdue.v1
- invoice.cancelled.v1
- invoice.discount.applied.v1
- invoice.reminder.sent.v1

#### Events Consumed
- milestone.approved.v1 (generate invoice)
- timesheet.approved.v1 (generate invoice)
- payment.processed.v1 (mark paid)

#### RBAC/SLO
- **RBAC:** FREELANCER (view own/generate), CLIENT (pay/view own), SYSTEM (generate/send reminders), ADMIN (view all/cancel)
- **SLO:** P95 < 400ms (generate - includes PDF), P95 < 250ms (send), P95 < 200ms (mark paid), P95 < 150ms (read)

---

### 6.2 line_item/

#### User Stories
- As a **system**, I want to **add line items** to invoices so that charges are itemized.
- As a **system**, I want to **calculate line item totals** with taxes so that amounts are accurate.
- As a **user**, I want to **view line item breakdown** so that charges are clear.

#### Flow
1. **AddLineItemCommand**(invoice_id, description, quantity, unit_price, tax_rate) → CalculateTotal() | Add() → **Outbox:** invoice.line_item.added.v1
2. **UpdateLineItemCommand**(line_item_id, updates) → Recalculate() | Update() → **Outbox:** invoice.line_item.updated.v1
3. **RemoveLineItemCommand**(line_item_id) → Remove() | RecalculateInvoice() → **Outbox:** invoice.line_item.removed.v1
4. **GetLineItemsQuery**(invoice_id) → Fetch() → LineItemListDTO

#### Projections
- line_items_read

#### Events Published
- invoice.line_item.added.v1
- invoice.line_item.updated.v1
- invoice.line_item.removed.v1

#### RBAC/SLO
- **RBAC:** SYSTEM (add/update/remove), USER (view)
- **SLO:** P95 < 180ms

---

## **7 - PLATFORM FEES DOMAIN**

### 7.1 fee/

#### User Stories
- As a **platform**, I want to **calculate platform fees** based on transaction type so that revenue is collected.
- As a **system**, I want to **support tiered fee structures** (by volume, user type) so that pricing is flexible.
- As a **system**, I want to **apply minimum and maximum caps** so that fees are reasonable.
- As a **admin**, I want to **configure fee rules** so that business model can evolve.
- As a **user**, I want to **see fee breakdown** so that charges are transparent.
- As a **system**, I want to **waive fees** for promotions so that marketing campaigns work.

#### Flow
1. **CalculateFeeCommand**(transaction_id, user_id, amount, transaction_type) → DetermineUserTier() | ApplyRules() | CalculateFee() → **Outbox:** fee.calculated.v1
2. **ApplyFeeCommand**(transaction_id, fee_amount) → DeductFee() | UpdateAccounting() → **Outbox:** fee.applied.v1
3. **WaiveFeeCommand**(transaction_id, reason, waived_by) → AuthorizeAdmin() | Waive() → **Outbox:** fee.waived.v1
4. **UpdateFeeRulesCommand**(rule_set, updated_by) → AuthorizeAdmin() | ValidateRules() | Update() → **Outbox:** fee.rules.updated.v1
5. **GetFeeBreakdownQuery**(transaction_id) → Calculate() → FeeBreakdownDTO
6. **GetUserFeeTierQuery**(user_id) → CalculateTier() → FeeTierDTO

#### Projections
- fee_rules_read
- fee_calculation_read
- user_fee_tiers_read

#### Events Published
- fee.calculated.v1
- fee.applied.v1
- fee.waived.v1
- fee.rules.updated.v1

#### RBAC/SLO
- **RBAC:** SYSTEM (calculate/apply), ADMIN (waive/update rules), USER (view breakdown)
- **SLO:** P95 < 150ms (calculate), P95 < 200ms (apply)

---

### 7.2 fee_v2/

#### User Stories
- As a **system**, I want to **support volume discounts** so that high-volume users pay less.
- As a **system**, I want to **apply country-specific fees** so that local regulations are followed.
- As a **system**, I want to **apply coupon codes** to fees so that promotions work.
- As a **system**, I want to **run A/B experiments** on fee structures so that optimization is possible.

#### Flow
1. **ApplyCouponCommand**(transaction_id, coupon_code) → ValidateCoupon() | CalculateDiscount() | Apply() → **Outbox:** fee.coupon.applied.v1
2. **CalculateVolumeDiscountCommand**(user_id, period) → AggregateVolume() | ApplyTiers() → **Outbox:** fee.volume_discount.calculated.v1
3. **ApplyCountryExceptionCommand**(country_code, fee_override) → AuthorizeAdmin() | Apply() → **Outbox:** fee.country_exception.applied.v1
4. **StartFeeExperimentCommand**(experiment_config) → AuthorizeAdmin() | AssignVariants() | Start() → **Outbox:** fee.experiment.started.v1

#### Projections
- fee_v2_rules_read
- fee_coupons_read
- fee_experiments_read
- volume_discounts_read

#### Events Published
- fee.coupon.applied.v1
- fee.volume_discount.calculated.v1
- fee.country_exception.applied.v1
- fee.experiment.started.v1

#### RBAC/SLO
- **RBAC:** SYSTEM (calculate), ADMIN (configure experiments/exceptions), USER (apply coupon)
- **SLO:** P95 < 200ms

---

## **8 - REFUND PROCESSING DOMAIN**

### 8.1 refund/

#### User Stories
- As a **client**, I want to **request refunds** so that I can recover funds.
- As a **system**, I want to **process full/partial refunds** so that flexibility is provided.
- As a **system**, I want to **calculate pro-rata refunds** based on work completed so that fairness is maintained.
- As a **admin**, I want to **approve refund requests** so that abuse is prevented.
- As a **system**, I want to **reverse fees** on refunds so that users aren't charged for returned funds.
- As a **user**, I want to **track refund status** so that I know when funds are returned.

#### Flow
1. **RequestRefundCommand**(payment_id, refund_amount, reason, requested_by) → ValidatePayment() | CalculateEligibility() | CreateRefundRequest() → **Outbox:** refund.requested.v1
2. **ApproveRefundCommand**(refund_id, approved_by) → AuthorizeAdmin() | Approve() | QueueProcessing() → **Outbox:** refund.approved.v1
3. **ProcessRefundCommand**(refund_id) → ProcessViaGateway() | ReverseFees() | UpdateWallets() → **Outbox:** refund.processed.v1
4. **ProcessPartialRefundCommand**(refund_id, refund_amount, pro_rata_basis) → CalculateProRata() | Process() → **Outbox:** refund.partial.processed.v1
5. **RejectRefundCommand**(refund_id, reason, rejected_by) → AuthorizeAdmin() | Reject() | NotifyUser(communications-be) → **Outbox:** refund.rejected.v1
6. **CancelRefundCommand**(refund_id, cancelled_by) → ValidateStatus() | Cancel() → **Outbox:** refund.cancelled.v1
7. **GetRefundQuery**(refund_id) → AuthorizeAccess() | Fetch() → RefundDTO
8. **ListRefundsQuery**(user_id, filters) → ApplyFilters() → RefundListDTO

#### Projections
- refund_read
- refund_requests_read
- refund_eligibility_read

#### Events Published
- refund.requested.v1
- refund.approved.v1
- refund.rejected.v1
- refund.processed.v1
- refund.partial.processed.v1
- refund.cancelled.v1
- refund.failed.v1

#### Events Consumed
- contract.terminated.v1 (trigger refund)
- dispute.resolved.v1 (process approved refunds)

#### RBAC/SLO
- **RBAC:** CLIENT (request), ADMIN (approve/reject), SYSTEM (process), USER (view own)
- **SLO:** P95 < 250ms (request), P95 < 200ms (approve/reject), P95 < 600ms (process), P95 < 150ms (read)

---

## **9 - TAX MANAGEMENT DOMAIN**

### 9.1 tax/

#### User Stories
- As a **system**, I want to **calculate taxes** (VAT, sales tax, withholding) so that invoices are compliant.
- As a **system**, I want to **support tax jurisdictions** (country, state, city) so that rates are accurate.
- As a **system**, I want to **handle reverse charge** for B2B EU transactions so that rules are followed.
- As a **system**, I want to **generate tax reports** (1099, VAT returns) so that compliance is automated.
- As a **freelancer**, I want to **submit tax forms** (W9, W8) so that withholding is correct.
- As a **system**, I want to **track tax thresholds** ($600 for 1099) so that reporting is triggered.

#### Flow
1. **CalculateTaxCommand**(amount, user_id, jurisdiction, tax_type) → DetermineJurisdiction() | LookupRate() | Calculate() → **Outbox:** tax.calculated.v1
2. **ApplyWithholdingCommand**(payment_id, withholding_rate) → CalculateWithholding() | Withhold() | UpdatePayment() → **Outbox:** tax.withheld.v1
3. **SubmitTaxFormCommand**(user_id, form_type, form_data, submitted_by) → ValidateForm() | Persist() | UpdateStatus() → **Outbox:** tax.form.submitted.v1
4. **Generate1099Command**(freelancer_id, tax_year) → AggregateEarnings() | CheckThreshold() | GenerateForm() | Send(communications-be) → **Outbox:** tax.1099.generated.v1
5. **GenerateVATReturnCommand**(period, jurisdiction) → AggregateVAT() | GenerateReturn() | Export() → **Outbox:** tax.vat_return.generated.v1
6. **UpdateTaxRatesCommand**(jurisdiction, rates, updated_by) → AuthorizeAdmin() | Update() | Invalidate Cache() → **Outbox:** tax.rates.updated.v1
7. **GetTaxCalculationQuery**(amount, jurisdiction) → Calculate() → TaxCalculationDTO
8. **GetTaxSummaryQuery**(user_id, tax_year) → Aggregate() → TaxSummaryDTO

#### Projections
- tax_rates_read
- tax_forms_read
- tax_calculations_read
- tax_thresholds_read

#### Events Published
- tax.calculated.v1
- tax.withheld.v1
- tax.form.submitted.v1
- tax.1099.generated.v1
- tax.vat_return.generated.v1
- tax.rates.updated.v1
- tax.threshold.reached.v1

#### Events Consumed
- payment.processed.v1 (track earnings for thresholds)
- user.profile.updated.v1 (update tax jurisdiction)

#### RBAC/SLO
- **RBAC:** SYSTEM (calculate/withhold/generate), USER (submit forms/view summary), ADMIN (update rates/generate returns)
- **SLO:** P95 < 180ms (calculate), P95 < 300ms (generate forms), P95 < 150ms (read)

---

## **10 - FOREIGN EXCHANGE DOMAIN**

### 10.1 fx/

#### User Stories
- As a **system**, I want to **convert currencies** for cross-border payments so that users can transact globally.
- As a **system**, I want to **fetch real-time exchange rates** from providers so that rates are current.
- As a **system**, I want to **cache exchange rates** with TTL so that API costs are controlled.
- As a **system**, I want to **apply FX margins** so that platform earns on currency conversion.
- As a **user**, I want to **lock exchange rates** for duration so that pricing is predictable.
- As a **user**, I want to **view FX fee breakdown** so that costs are transparent.

#### Flow
1. **ConvertCurrencyCommand**(amount, from_currency, to_currency, conversion_purpose) → FetchRate() | ApplyMargin() | Convert() | CreateTransaction() → **Outbox:** currency.converted.v1
2. **UpdateExchangeRatesCommand**() → FetchFromProvider() | ValidateRates() | Update() | InvalidateCache() → **Outbox:** fx.rates.updated.v1
3. **LockExchangeRateCommand**(user_id, currency_pair, duration, locked_by) → ValidateDuration() | LockRate() | ScheduleExpiry() → **Outbox:** fx.rate.locked.v1
4. **GetExchangeRateQuery**(from_currency, to_currency) → FetchCached() | Calculate() → ExchangeRateDTO
5. **GetConversionHistoryQuery**(user_id, filters) → Fetch() → ConversionHistoryDTO
6. **GetFXFeeBreakdownQuery**(amount, currency_pair) → Calculate() → FXFeeBreakdownDTO

#### Projections
- fx_rates_read
- fx_conversions_read
- fx_locks_read

#### Events Published
- currency.converted.v1
- fx.rates.updated.v1
- fx.rate.locked.v1
- fx.rate.lock.expired.v1

#### RBAC/SLO
- **RBAC:** SYSTEM (convert/update rates), USER (lock rate/view history), PUBLIC (get rates)
- **SLO:** P95 < 200ms (convert), P95 < 300ms (update rates - includes provider call), P95 < 100ms (get rate - cached)

---

## **11 - RISK & FRAUD DOMAIN**

### 11.1 risk/

#### User Stories
- As a **system**, I want to **detect fraud patterns** so that losses are prevented.
- As a **system**, I want to **calculate risk scores** for transactions so that high-risk ones are flagged.
- As a **system**, I want to **place holds** on suspicious transactions so that review can happen.
- As a **system**, I want to **monitor velocity** (transactions per hour) so that abnormal activity is detected.
- As a **admin**, I want to **review flagged transactions** so that decisions can be made.
- As a **system**, I want to **integrate with fraud detection APIs** so that advanced checks are performed.

#### Flow
1. **CalculateRiskScoreCommand**(transaction_id, user_id, amount, metadata) → AnalyzePatterns() | CheckVelocity() | CalculateScore() → **Outbox:** risk.score.calculated.v1
2. **PlaceRiskHoldCommand**(transaction_id, risk_reason, hold_duration) → PlaceHold() | NotifyParties(communications-be) | QueueReview() → **Outbox:** risk.hold.placed.v1
3. **ReleaseRiskHoldCommand**(hold_id, release_reason, released_by) → AuthorizeAdmin() | Release() | ProcessTransaction() → **Outbox:** risk.hold.released.v1
4. **FlagTransactionCommand**(transaction_id, flag_type, evidence) → Flag() | NotifyRiskTeam() | CreateReviewCase() → **Outbox:** transaction.flagged.v1
5. **ReviewFlaggedTransactionCommand**(flag_id, decision, reviewed_by) → AuthorizeRiskAdmin() | ApplyDecision() | UpdateStatus() → **Outbox:** transaction.review.completed.v1
6. **MonitorVelocityCommand**(user_id, time_window) → CountTransactions() | CheckThresholds() | TriggerAlerts() → **Outbox:** velocity.alert.triggered.v1
7. **GetRiskScoreQuery**(transaction_id) → Calculate() → RiskScoreDTO
8. **GetFlaggedTransactionsQuery**(filters) → AuthorizeRiskAdmin() | Fetch() → FlaggedTransactionListDTO

#### Projections
- risk_scores_read
- risk_holds_read
- flagged_transactions_read
- velocity_metrics_read

#### Events Published
- risk.score.calculated.v1
- risk.hold.placed.v1
- risk.hold.released.v1
- transaction.flagged.v1
- transaction.review.completed.v1
- velocity.alert.triggered.v1
- fraud.detected.v1

#### Events Consumed
- payment.processed.v1 (calculate risk score)
- user.suspended.v1 (place holds on user transactions)

#### RBAC/SLO
- **RBAC:** SYSTEM (calculate/monitor/flag), RISK_ADMIN (review/release holds), ADMIN (view all)
- **SLO:** P95 < 200ms (calculate score), P95 < 250ms (place hold), P95 < 300ms (review)

---

### 11.2 chargeback/

#### User Stories
- As a **system**, I want to **handle chargebacks** from payment gateways so that disputes are tracked.
- As a **system**, I want to **place reserve** for chargeback risk so that losses are covered.
- As a **admin**, I want to **respond to chargebacks** with evidence so that wins are possible.
- As a **system**, I want to **track chargeback ratio** so that high-risk accounts are identified.
- As a **system**, I want to **deduct funds** on chargeback loss so that accounting is correct.

#### Flow
1. **CreateChargebackCommand**(payment_id, chargeback_amount, reason, gateway_ref) → ValidatePayment() | PlaceReserve() | CreateCase() → **Outbox:** chargeback.created.v1
2. **RespondToChargebackCommand**(chargeback_id, evidence_urls[], response_text, responded_by) → AuthorizeAdmin() | SubmitToGateway() | UpdateStatus() → **Outbox:** chargeback.responded.v1
3. **ResolveChargebackCommand**(chargeback_id, outcome, resolution_amount) → ApplyOutcome() | ReleaseOrDeductReserve() | UpdateAccounting() → **Outbox:** chargeback.resolved.v1
4. **GetChargebackQuery**(chargeback_id) → AuthorizeAccess() | Fetch() → ChargebackDTO
5. **GetChargebackRatioQuery**(user_id, period) → Calculate() → ChargebackRatioDTO

#### Projections
- chargeback_read
- chargeback_reserves_read
- chargeback_ratios_read

#### Events Published
- chargeback.created.v1
- chargeback.responded.v1
- chargeback.resolved.v1
- chargeback.won.v1
- chargeback.lost.v1

#### Events Consumed
- payment.gateway.chargeback.notification (from webhook)

#### RBAC/SLO
- **RBAC:** SYSTEM (create), ADMIN (respond), RISK_ADMIN (resolve), USER (view own)
- **SLO:** P95 < 300ms (create), P95 < 400ms (respond), P95 < 350ms (resolve)

---

## **12 - BONUS & INCENTIVES DOMAIN**

### 12.1 bonus/

#### User Stories
- As a **platform**, I want to **issue bonuses** to users so that engagement is rewarded.
- As a **system**, I want to **calculate performance bonuses** based on KPIs so that rewards are merit-based.
- As a **system**, I want to **issue referral bonuses** so that growth is incentivized.
- As a **system**, I want to **apply bonus vesting** so that retention is improved.
- As a **user**, I want to **view bonus history** so that earnings are tracked.

#### Flow
1. **IssueBonusCommand**(user_id, amount, bonus_type, reason, issued_by) → ValidateUser() | CalculateAmount() | CreditWallet() → **Outbox:** bonus.issued.v1
2. **CalculatePerformanceBonusCommand**(user_id, period, kpis[]) → EvaluateKPIs() | CalculateAmount() | IssueBonus() → **Outbox:** performance_bonus.calculated.v1
3. **IssueReferralBonusCommand**(referrer_id, referred_id, bonus_tier) → ValidateReferral() | CalculateAmount() | IssueBonus() → **Outbox:** referral_bonus.issued.v1
4. **VestBonusCommand**(bonus_id) → CheckVestingSchedule() | Vest() | MakeAvailable() → **Outbox:** bonus.vested.v1
5. **GetBonusHistoryQuery**(user_id, filters) → Fetch() → BonusHistoryDTO
6. **GetVestingScheduleQuery**(user_id) → Fetch() → VestingScheduleDTO

#### Projections
- bonus_read
- bonus_history_read
- vesting_schedule_read

#### Events Published
- bonus.issued.v1
- performance_bonus.calculated.v1
- referral_bonus.issued.v1
- bonus.vested.v1

#### Events Consumed
- contract.completed.v1 (trigger performance bonus)
- referral.converted.v1 (issue referral bonus)

#### RBAC/SLO
- **RBAC:** SYSTEM (issue/calculate/vest), USER (view own), ADMIN (issue manual bonuses)
- **SLO:** P95 < 250ms (issue), P95 < 300ms (calculate performance), P95 < 150ms (read)

---

## **13 - EXPENSE REIMBURSEMENT DOMAIN**

### 13.1 expense/

#### User Stories
- As a **freelancer**, I want to **submit expenses** for reimbursement so that costs are covered.
- As a **freelancer**, I want to **attach receipts** to expenses so that proof is provided.
- As a **client**, I want to **approve/reject expenses** so that legitimate costs are paid.
- As a **system**, I want to **validate expense limits** per contract so that budgets are maintained.
- As a **system**, I want to **reimburse approved expenses** so that freelancers are paid.

#### Flow
1. **SubmitExpenseCommand**(contract_id, amount, category, description, receipt_urls[], submitted_by) → ValidateContract() | CheckLimit() | Persist() | NotifyClient(communications-be) → **Outbox:** expense.submitted.v1
2. **ApproveExpenseCommand**(expense_id, approved_by) → AuthorizeClient() | Approve() | QueueReimbursement() → **Outbox:** expense.approved.v1
3. **RejectExpenseCommand**(expense_id, rejection_reason, rejected_by) → AuthorizeClient() | Reject() | NotifyFreelancer(communications-be) → **Outbox:** expense.rejected.v1
4. **ReimburseExpenseCommand**(expense_id) → ValidateApproval() | ProcessPayment() | MarkReimbursed() → **Outbox:** expense.reimbursed.v1
5. **GetExpenseQuery**(expense_id) → AuthorizeAccess() | Fetch() → ExpenseDTO
6. **ListExpensesQuery**(contract_id, filters) → ApplyFilters() → ExpenseListDTO

#### Projections
- expense_read
- expense_approvals_read
- expense_limits_read

#### Events Published
- expense.submitted.v1
- expense.approved.v1
- expense.rejected.v1
- expense.reimbursed.v1
- expense.limit.exceeded.v1

#### RBAC/SLO
- **RBAC:** FREELANCER (submit/view own), CLIENT (approve/reject), SYSTEM (reimburse)
- **SLO:** P95 < 250ms (submit), P95 < 200ms (approve/reject), P95 < 400ms (reimburse)

---

## **14 - PAYMENT SCHEDULES & AUTOMATION DOMAIN**

### 14.1 payment_schedule/

#### User Stories
- As a **client**, I want to **set up recurring payments** so that retainers are automated.
- As a **system**, I want to **process scheduled payments** automatically so that manual work is reduced.
- As a **system**, I want to **send payment reminders** before due date so that failures are prevented.
- As a **system**, I want to **retry failed scheduled payments** so that transient issues are handled.
- As a **user**, I want to **cancel payment schedules** so that subscriptions can be stopped.

#### Flow
1. **CreatePaymentScheduleCommand**(contract_id, amount, frequency, start_date, end_date, created_by) → ValidateContract() | Schedule() | Persist() → **Outbox:** payment_schedule.created.v1
2. **ProcessScheduledPaymentCommand**(schedule_id) → FetchSchedule() | ProcessPayment() | UpdateNextRun() → **Outbox:** scheduled_payment.processed.v1
3. **SendPaymentReminderCommand**(schedule_id, reminder_days_before) → FetchSchedule() | Send(communications-be) | LogReminder() → **Outbox:** payment_reminder.sent.v1
4. **RetryFailedScheduledPaymentCommand**(schedule_id, retry_count) → ValidateRetryable() | ProcessPayment() → **Outbox:** scheduled_payment.retried.v1
5. **CancelPaymentScheduleCommand**(schedule_id, reason, cancelled_by) → ValidateOwner() | Cancel() | NotifyParties() → **Outbox:** payment_schedule.cancelled.v1
6. **GetPaymentScheduleQuery**(schedule_id) → AuthorizeAccess() | Fetch() → PaymentScheduleDTO
7. **ListPaymentSchedulesQuery**(contract_id, filters) → ApplyFilters() → PaymentScheduleListDTO

#### Projections
- payment_schedule_read
- scheduled_payment_history_read

#### Events Published
- payment_schedule.created.v1
- payment_schedule.updated.v1
- scheduled_payment.processed.v1
- scheduled_payment.failed.v1
- scheduled_payment.retried.v1
- payment_schedule.cancelled.v1
- payment_reminder.sent.v1

#### Events Consumed
- contract.activated.v1 (create schedule if recurring)
- contract.terminated.v1 (cancel schedules)

#### RBAC/SLO
- **RBAC:** CLIENT (create/cancel), SYSTEM (process/retry/send reminders), USER (view own)
- **SLO:** P95 < 250ms (create), P95 < 600ms (process - includes payment), P95 < 150ms (read)

---

## **15 - RECONCILIATION & REPORTING DOMAIN**

### 15.1 reconciliation/

#### User Stories
- As a **finance admin**, I want to **reconcile daily transactions** with bank statements so that accuracy is verified.
- As a **system**, I want to **detect discrepancies** automatically so that issues are flagged.
- As a **system**, I want to **generate reconciliation reports** so that auditing is easier.
- As a **admin**, I want to **resolve discrepancies** with adjustments so that books balance.

#### Flow
1. **RunDailyReconciliationCommand**(reconciliation_date, bank_statement_file) → ParseStatement() | MatchTransactions() | FlagDiscrepancies() | GenerateReport() → **Outbox:** reconciliation.completed.v1
2. **FlagDiscrepancyCommand**(transaction_id, discrepancy_type, details) → Flag() | NotifyFinanceTeam() | CreateCase() → **Outbox:** reconciliation.discrepancy.flagged.v1
3. **ResolveDiscrepancyCommand**(discrepancy_id, resolution_type, adjustment_amount, resolved_by) → AuthorizeFinanceAdmin() | ApplyAdjustment() | Resolve() → **Outbox:** reconciliation.discrepancy.resolved.v1
4. **GetReconciliationReportQuery**(date_range) → AuthorizeFinanceAdmin() | GenerateReport() → ReconciliationReportDTO
5. **GetDiscrepanciesQuery**(filters) → AuthorizeFinanceAdmin() | Fetch() → DiscrepancyListDTO

#### Projections
- reconciliation_reports_read
- reconciliation_discrepancies_read
- reconciliation_matches_read

#### Events Published
- reconciliation.started.v1
- reconciliation.completed.v1
- reconciliation.discrepancy.flagged.v1
- reconciliation.discrepancy.resolved.v1

#### RBAC/SLO
- **RBAC:** FINANCE_ADMIN (run/resolve), SYSTEM (run daily), AUDITOR (view reports)
- **SLO:** P95 < 10000ms (daily reconciliation - large data set), P95 < 300ms (flag), P95 < 400ms (resolve)

---

### 15.2 reporting/

#### User Stories
- As a **admin**, I want to **generate financial reports** (revenue, payouts, fees) so that business health is tracked.
- As a **admin**, I want to **export reports** to various formats (CSV, PDF, Excel) so that sharing is easy.
- As a **system**, I want to **schedule automated reports** so that stakeholders receive updates regularly.
- As a **admin**, I want to **view real-time dashboards** so that current status is visible.

#### Flow
1. **GenerateFinancialReportCommand**(report_type, date_range, filters, generated_by) → AuthorizeAdmin() | AggregateData() | GenerateReport() | Upload(storage-be) → **Outbox:** report.generated.v1
2. **ExportReportCommand**(report_id, export_format, exported_by) → AuthorizeAdmin() | FetchData() | ConvertFormat() | Upload(storage-be) → **Outbox:** report.exported.v1
3. **ScheduleReportCommand**(report_config, frequency, recipients[], scheduled_by) → AuthorizeAdmin() | Schedule() → **Outbox:** report.scheduled.v1
4. **GetDashboardMetricsQuery**(metric_types[], date_range) → AuthorizeAdmin() | AggregateRealTime() → DashboardMetricsDTO
5. **GetReportQuery**(report_id) → AuthorizeAdmin() | Fetch() → ReportDTO

#### Projections
- financial_reports_read
- report_schedules_read
- dashboard_metrics_read

#### Events Published
- report.generated.v1
- report.exported.v1
- report.scheduled.v1
- report.sent.v1

#### RBAC/SLO
- **RBAC:** ADMIN (generate/export/schedule/view), FINANCE_ADMIN (all operations)
- **SLO:** P95 < 5000ms (generate - large data aggregation), P95 < 2000ms (export), P95 < 500ms (dashboard)

---

## **16 - SUBSCRIPTION BILLING DOMAIN**

### 16.1 subscription/

#### User Stories
- As a **user**, I want to **subscribe to plans** so that I get premium features.
- As a **system**, I want to **charge recurring subscription fees** automatically so that billing is hands-off.
- As a **system**, I want to **prorate subscription changes** so that charges are fair.
- As a **system**, I want to **handle failed subscription payments** with retry logic so that subscriptions aren't cancelled prematurely.
- As a **user**, I want to **cancel subscriptions** so that I can stop charges.
- As a **system**, I want to **send renewal reminders** before billing so that surprises are avoided.

#### Flow
1. **CreateSubscriptionCommand**(user_id, plan_id, payment_method_id, created_by) → ValidatePlan() | ChargeSetupFee() | CreateSchedule() → **Outbox:** subscription.created.v1
2. **ChargeSubscriptionCommand**(subscription_id) → FetchSubscription() | ProcessPayment() | UpdatePeriod() | ScheduleNextCharge() → **Outbox:** subscription.charged.v1
3. **UpgradeSubscriptionCommand**(subscription_id, new_plan_id, upgraded_by) → CalculateProration() | ChargeUpgrade() | UpdatePlan() → **Outbox:** subscription.upgraded.v1
4. **DowngradeSubscriptionCommand**(subscription_id, new_plan_id, downgraded_by) → ApplyCredit() | UpdatePlan() | ScheduleChange() → **Outbox:** subscription.downgraded.v1
5. **CancelSubscriptionCommand**(subscription_id, cancellation_reason, cancel_at_period_end, cancelled_by) → AuthorizeOwner() | Cancel() | ProcessRefund() → **Outbox:** subscription.cancelled.v1
6. **RetryFailedSubscriptionPaymentCommand**(subscription_id, retry_count) → ProcessPayment() | UpdateStatus() → **Outbox:** subscription.payment.retried.v1
7. **SendRenewalReminderCommand**(subscription_id, reminder_days_before) → Send(communications-be) | LogReminder() → **Outbox:** subscription.renewal.reminder.sent.v1
8. **GetSubscriptionQuery**(subscription_id) → AuthorizeAccess() | Fetch() → SubscriptionDTO
9. **ListSubscriptionsQuery**(user_id, filters) → ApplyFilters() → SubscriptionListDTO

#### Projections
- subscription_read
- subscription_billing_history_read
- subscription_status_read

#### Events Published
- subscription.created.v1
- subscription.charged.v1
- subscription.upgraded.v1
- subscription.downgraded.v1
- subscription.cancelled.v1
- subscription.payment.failed.v1
- subscription.payment.retried.v1
- subscription.renewal.reminder.sent.v1
- subscription.expired.v1

#### Events Consumed
- payment.failed.v1 (handle failed subscription payment)
- user.deleted.v1 (cancel subscriptions)

#### RBAC/SLO
- **RBAC:** USER (create/cancel/view own), SYSTEM (charge/retry), ADMIN (view all/force cancel)
- **SLO:** P95 < 400ms (create), P95 < 600ms (charge), P95 < 300ms (upgrade/downgrade), P95 < 150ms (read)

---

### 16.2 connects_purchase/

#### User Stories
- As a **freelancer**, I want to **purchase connect packages** so that I can submit proposals.
- As a **system**, I want to **track connect balance** so that usage is limited.
- As a **system**, I want to **deduct connects** on proposal submission so that usage is charged.
- As a **system**, I want to **refund connects** when proposals aren't viewed so that fairness is maintained.
- As a **system**, I want to **apply bonus connects** for promotions so that engagement is boosted.

#### Flow
1. **PurchaseConnectsCommand**(user_id, package_id, payment_method_id, purchased_by) → ValidatePackage() | ProcessPayment() | CreditConnects() → **Outbox:** connects.purchased.v1
2. **DeductConnectsCommand**(user_id, amount, proposal_id, reason) → ValidateBalance() | Deduct() | UpdateBalance() → **Outbox:** connects.deducted.v1
3. **RefundConnectsCommand**(user_id, amount, reason, refunded_by) → Validate() | Credit() | UpdateBalance() → **Outbox:** connects.refunded.v1
4. **AddBonusConnectsCommand**(user_id, amount, reason, added_by) → Credit() | SetExpiry() → **Outbox:** connects.bonus.added.v1
5. **GetConnectBalanceQuery**(user_id) → AuthorizeOwner() | Fetch() → ConnectBalanceDTO
6. **GetConnectHistoryQuery**(user_id, filters) → Fetch() → ConnectHistoryDTO

#### Projections
- connect_balance_read
- connect_history_read
- connect_packages_read

#### Events Published
- connects.purchased.v1
- connects.deducted.v1
- connects.refunded.v1
- connects.bonus.added.v1
- connects.balance.low.v1

#### Events Consumed
- proposal.submitted.v1 (deduct connects)
- proposal.not_viewed.v1 (refund connects)
- subscription.upgraded.v1 (add bonus connects)

#### RBAC/SLO
- **RBAC:** FREELANCER (purchase/view own), SYSTEM (deduct/refund/add bonus), ADMIN (view all)
- **SLO:** P95 < 400ms (purchase - includes payment), P95 < 150ms (deduct), P95 < 200ms (refund)

---

## **17 - WITHDRAWAL LIMITS & COMPLIANCE DOMAIN**

### 17.1 withdrawal_limit/

#### User Stories
- As a **system**, I want to **enforce withdrawal limits** based on KYC level so that compliance is maintained.
- As a **system**, I want to **track daily/monthly withdrawal amounts** so that limits are checked.
- As a **admin**, I want to **override withdrawal limits** for verified users so that flexibility exists.
- As a **system**, I want to **require additional verification** for large withdrawals so that security is maintained.

#### Flow
1. **CheckWithdrawalLimitCommand**(user_id, amount) → GetKYCLevel() | GetCurrentUsage() | ValidateLimit() → WithdrawalLimitResultDTO
2. **TrackWithdrawalCommand**(user_id, amount, withdrawal_id) → UpdateDailyTotal() | UpdateMonthlyTotal() → **Outbox:** withdrawal.tracked.v1
3. **OverrideLimitCommand**(user_id, new_limit, reason, overridden_by) → AuthorizeAdmin() | UpdateLimit() | SetExpiry() → **Outbox:** withdrawal_limit.overridden.v1
4. **RequestHighValueWithdrawalCommand**(user_id, amount, requested_by) → ValidateLimit() | RequireVerification() | CreateRequest() → **Outbox:** high_value_withdrawal.requested.v1
5. **ApproveHighValueWithdrawalCommand**(request_id, approved_by) → AuthorizeAdmin() | Approve() | ProcessWithdrawal() → **Outbox:** high_value_withdrawal.approved.v1
6. **GetWithdrawalLimitsQuery**(user_id) → Fetch() → WithdrawalLimitsDTO
7. **GetWithdrawalUsageQuery**(user_id, period) → Calculate() → WithdrawalUsageDTO

#### Projections
- withdrawal_limits_read
- withdrawal_usage_read
- high_value_requests_read

#### Events Published
- withdrawal.tracked.v1
- withdrawal_limit.reached.v1
- withdrawal_limit.overridden.v1
- high_value_withdrawal.requested.v1
- high_value_withdrawal.approved.v1

#### Events Consumed
- wallet.withdrawal.completed.v1 (track usage)
- user.kyc.updated.v1 (adjust limits)

#### RBAC/SLO
- **RBAC:** SYSTEM (check/track), ADMIN (override/approve), USER (view own/request)
- **SLO:** P95 < 150ms (check), P95 < 180ms (track), P95 < 250ms (override)

---

## **18 - CURRENCY WALLET DOMAIN**

### 18.1 multi_currency/

#### User Stories
- As a **user**, I want to **hold multiple currencies** in separate sub-wallets so that FX fees are minimized.
- As a **user**, I want to **transfer between currency wallets** so that balances are managed.
- As a **system**, I want to **default to user's preferred currency** for payments so that UX is good.
- As a **system**, I want to **support 100+ currencies** so that global coverage is complete.

#### Flow
1. **CreateCurrencyWalletCommand**(user_id, currency, created_by) → ValidateCurrency() | CreateSubWallet() → **Outbox:** currency_wallet.created.v1
2. **TransferBetweenCurrenciesCommand**(user_id, from_currency, to_currency, amount) → ConvertCurrency() | Transfer() → **Outbox:** currency_transfer.completed.v1
3. **SetPreferredCurrencyCommand**(user_id, currency) → AuthorizeOwner() | Update() → **Outbox:** preferred_currency.set.v1
4. **GetCurrencyWalletsQuery**(user_id) → Fetch() → CurrencyWalletListDTO

#### Projections
- currency_wallets_read
- currency_preferences_read

#### Events Published
- currency_wallet.created.v1
- currency_transfer.completed.v1
- preferred_currency.set.v1

#### RBAC/SLO
- **RBAC:** USER (create/transfer/set preference on own), ADMIN (view all)
- **SLO:** P95 < 250ms (create), P95 < 400ms (transfer - includes FX)

---

## **19 - PAYMENT GATEWAY INTEGRATION DOMAIN**

### 19.1 gateway_webhook/

#### User Stories
- As a **system**, I want to **receive webhook notifications** from gateways so that payment status is updated.
- As a **system**, I want to **verify webhook signatures** so that authenticity is ensured.
- As a **system**, I want to **retry webhook processing** on failure so that events aren't lost.
- As a **system**, I want to **log all webhook events** so that debugging is possible.

#### Flow
1. **ProcessWebhookCommand**(gateway, event_type, payload, signature) → VerifySignature() | ParseEvent() | ProcessEvent() | AcknowledgeWebhook() → **Outbox:** webhook.processed.v1
2. **RetryWebhookProcessingCommand**(webhook_id, retry_count) → FetchWebhook() | ProcessEvent() → **Outbox:** webhook.retry.processed.v1
3. **GetWebhookQuery**(webhook_id) → AuthorizeAdmin() | Fetch() → WebhookDTO
4. **ListWebhooksQuery**(gateway, filters) → AuthorizeAdmin() | Fetch() → WebhookListDTO

#### Projections
- webhook_log_read
- webhook_failures_read

#### Events Published
- webhook.received.v1
- webhook.processed.v1
- webhook.retry.processed.v1
- webhook.processing.failed.v1
- webhook.signature.invalid.v1

#### RBAC/SLO
- **RBAC:** SYSTEM (process/retry), ADMIN (view logs)
- **SLO:** P95 < 300ms (process), P95 < 200ms (verify signature)

---

## **20 - DISPUTE & RESOLUTION DOMAIN**

### 20.1 payment_dispute/

#### User Stories
- As a **user**, I want to **dispute payments** when there are issues so that resolution is possible.
- As a **admin**, I want to **investigate disputes** so that fair decisions are made.
- As a **system**, I want to **place holds during disputes** so that funds are protected.
- As a **admin**, I want to **resolve disputes** with decisions so that parties know outcomes.

#### Flow
1. **CreatePaymentDisputeCommand**(payment_id, dispute_reason, evidence_urls[], created_by) → ValidatePayment() | PlaceHold() | CreateCase() | Notify(admin-be) → **Outbox:** payment_dispute.created.v1
2. **AddDisputeEvidenceCommand**(dispute_id, evidence_urls[], description, added_by) → AuthorizeParty() | Attach() → **Outbox:** dispute_evidence.added.v1
3. **InvestigateDisputeCommand**(dispute_id, investigated_by) → AuthorizeAdmin() | GatherInfo() | UpdateStatus() → **Outbox:** dispute.under_investigation.v1
4. **ResolvePaymentDisputeCommand**(dispute_id, resolution_type, refund_amount, resolved_by) → AuthorizeAdmin() | ApplyResolution() | ReleaseHold() | ProcessRefund() → **Outbox:** payment_dispute.resolved.v1
5. **GetDisputeQuery**(dispute_id) → AuthorizeAccess() | Fetch() → PaymentDisputeDTO
6. **ListDisputesQuery**(filters) → AuthorizeParty() | Fetch() → DisputeListDTO

#### Projections
- payment_dispute_read
- dispute_evidence_read
- dispute_resolution_read

#### Events Published
- payment_dispute.created.v1
- dispute_evidence.added.v1
- dispute.under_investigation.v1
- payment_dispute.resolved.v1
- dispute.escalated.v1

#### Events Consumed
- contract.dispute.opened.v1 (may trigger payment dispute)

#### RBAC/SLO
- **RBAC:** USER (create/add evidence/view own), ADMIN (investigate/resolve/view all)
- **SLO:** P95 < 300ms (create), P95 < 250ms (add evidence), P95 < 400ms (resolve)

---

## **21 - MARKETPLACE FEES & COMMISSIONS DOMAIN**

### 21.1 commission/

#### User Stories
- As a **platform**, I want to **calculate commissions** on transactions so that revenue is earned.
- As a **system**, I want to **support split payments** (freelancer + platform) so that distribution is automated.
- As a **system**, I want to **apply commission tiers** based on volume so that pricing is flexible.
- As a **admin**, I want to **configure commission rates** per category so that business model adapts.

#### Flow
1. **CalculateCommissionCommand**(transaction_id, amount, transaction_type) → DetermineTier() | ApplyRules() | Calculate() → **Outbox:** commission.calculated.v1
2. **ApplyCommissionCommand**(transaction_id, commission_amount) → SplitPayment() | DistributeFunds() → **Outbox:** commission.applied.v1
3. **UpdateCommissionRatesCommand**(category, rates, updated_by) → AuthorizeAdmin() | Update() | Invalidate Cache() → **Outbox:** commission_rates.updated.v1
4. **GetCommissionBreakdownQuery**(transaction_id) → Calculate() → CommissionBreakdownDTO
5. **GetCommissionReportQuery**(date_range, filters) → AuthorizeAdmin() | Aggregate() → CommissionReportDTO

#### Projections
- commission_rates_read
- commission_history_read
- commission_reports_read

#### Events Published
- commission.calculated.v1
- commission.applied.v1
- commission_rates.updated.v1

#### RBAC/SLO
- **RBAC:** SYSTEM (calculate/apply), ADMIN (update rates/view reports), USER (view breakdown)
- **SLO:** P95 < 150ms (calculate), P95 < 200ms (apply)

---

## **22 - PROMOTIONAL CREDITS & COUPONS DOMAIN**

### 22.1 promo_credit/

#### User Stories
- As a **marketing**, I want to **issue promotional credits** to users so that acquisition/retention is boosted.
- As a **user**, I want to **redeem coupon codes** so that discounts are applied.
- As a **system**, I want to **validate coupon constraints** (min purchase, user eligibility) so that abuse is prevented.
- As a **system**, I want to **expire promotional credits** after time period so that usage is time-bound.
- As a **system**, I want to **track promo redemption** so that campaign ROI is measured.

#### Flow
1. **IssuePromoCreditCommand**(user_id, amount, credit_type, expiry_date, campaign_id, issued_by) → ValidateUser() | CreditWallet() | SetExpiry() → **Outbox:** promo_credit.issued.v1
2. **RedeemCouponCommand**(user_id, coupon_code, transaction_id) → ValidateCoupon() | CheckConstraints() | ApplyDiscount() | MarkRedeemed() → **Outbox:** coupon.redeemed.v1
3. **ExpirePromoCreditCommand**(credit_id) → Expire() | DebitUnused() → **Outbox:** promo_credit.expired.v1
4. **GetPromoCreditBalanceQuery**(user_id) → Fetch() → PromoCreditBalanceDTO
5. **GetCouponDetailsQuery**(coupon_code) → Fetch() → CouponDTO
6. **GetRedemptionReportQuery**(campaign_id, date_range) → AuthorizeMarketing() | Aggregate() → RedemptionReportDTO

#### Projections
- promo_credits_read
- coupons_read
- redemption_history_read

#### Events Published
- promo_credit.issued.v1
- promo_credit.expired.v1
- coupon.redeemed.v1
- coupon.invalid.v1

#### RBAC/SLO
- **RBAC:** MARKETING (issue/view reports), USER (redeem/view balance), SYSTEM (expire)
- **SLO:** P95 < 200ms (issue), P95 < 250ms (redeem), P95 < 150ms (read)

---

## **23 - BANK ACCOUNT VERIFICATION DOMAIN**

### 23.1 bank_verification/

#### User Stories
- As a **user**, I want to **verify bank accounts** via micro-deposits so that payouts are enabled.
- As a **system**, I want to **send micro-deposits** (two small amounts) so that ownership is proven.
- As a **user**, I want to **confirm micro-deposit amounts** so that verification completes.
- As a **system**, I want to **instantly verify** via Plaid/Stripe so that onboarding is faster.

#### Flow
1. **InitiateBankVerificationCommand**(user_id, bank_account_details, verification_method) → ValidateDetails() | SendMicroDeposits() | CreateVerificationRequest() → **Outbox:** bank_verification.initiated.v1
2. **ConfirmMicroDepositsCommand**(verification_id, amount1, amount2, confirmed_by) → ValidateAmounts() | ActivateAccount() → **Outbox:** bank_verification.completed.v1
3. **InstantVerifyCommand**(user_id, plaid_token) → VerifyViaPlaid() | ActivateAccount() → **Outbox:** bank_verification.instant.completed.v1
4. **GetVerificationStatusQuery**(verification_id) → AuthorizeOwner() | Fetch() → VerificationStatusDTO

#### Projections
- bank_verification_read

#### Events Published
- bank_verification.initiated.v1
- bank_verification.completed.v1
- bank_verification.instant.completed.v1
- bank_verification.failed.v1

#### RBAC/SLO
- **RBAC:** USER (initiate/confirm own), SYSTEM (send deposits)
- **SLO:** P95 < 300ms (initiate), P95 < 200ms (confirm), P95 < 500ms (instant verify - includes Plaid)

---

## **24 - FINANCIAL ANALYTICS & INSIGHTS DOMAIN**

### 24.1 analytics/

#### User Stories
- As a **user**, I want to **view earnings analytics** (trends, averages, projections) so that financial health is understood.
- As a **user**, I want to **view spending analytics** so that budget is managed.
- As a **admin**, I want to **platform-wide analytics** (GMV, take rate, churn) so that business metrics are tracked.
- As a **system**, I want to **generate insights** (savings opportunities, cashflow predictions) so that users make better decisions.

#### Flow
1. **GenerateEarningsAnalyticsCommand**(user_id, date_range) → AggregateEarnings() | CalculateTrends() | GenerateInsights() → **Outbox:** earnings_analytics.generated.v1
2. **GenerateSpendingAnalyticsCommand**(user_id, date_range) → AggregateSpending() | CategorizeExpenses() | GenerateInsights() → **Outbox:** spending_analytics.generated.v1
3. **GeneratePlatformAnalyticsCommand**(date_range, metrics[]) → AuthorizeAdmin() | AggregateData() | CalculateKPIs() → **Outbox:** platform_analytics.generated.v1
4. **GetEarningsAnalyticsQuery**(user_id, date_range) → Fetch() → EarningsAnalyticsDTO
5. **GetSpendingAnalyticsQuery**(user_id, date_range) → Fetch() → SpendingAnalyticsDTO
6. **GetPlatformKPIsQuery**(date_range) → AuthorizeAdmin() | Fetch() → PlatformKPIsDTO

#### Projections
- earnings_analytics_read
- spending_analytics_read
- platform_analytics_read

#### Events Published
- earnings_analytics.generated.v1
- spending_analytics.generated.v1
- platform_analytics.generated.v1

#### RBAC/SLO
- **RBAC:** USER (view own analytics), ADMIN (platform analytics), SYSTEM (generate)
- **SLO:** P95 < 1000ms (generate - aggregation intensive), P95 < 300ms (read)

---

## **25 - INBOX: EVENT CONSUMERS**

### 25.1 Contract Events → Payment Processing

#### User Stories
- As a **system**, I want to **consume milestone.approved events** so that payments are triggered.
- As a **system**, I want to **consume timesheet.approved events** so that hourly payments are processed.
- As a **system**, I want to **consume contract.terminated events** so that refunds are handled.

#### Flow
- Consume: `milestone.approved.v1` → Trigger `ProcessPaymentCommand`
- Consume: `timesheet.approved.v1` → Trigger `ProcessPaymentCommand`
- Consume: `contract.activated.v1` → Trigger `CreateEscrowCommand`
- Consume: `contract.terminated.v1` → Trigger `RefundEscrowCommand`

#### Projections
- payment_triggers_read

#### Events Consumed
- milestone.approved.v1
- timesheet.approved.v1
- contract.activated.v1
- contract.terminated.v1

#### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 200ms (processing lag)

---

### 25.2 User Events → Wallet Management

#### User Stories
- As a **system**, I want to **consume user.created events** so that wallets are auto-created.
- As a **system**, I want to **consume user.verified events** so that withdrawal limits are increased.
- As a **system**, I want to **consume user.suspended events** so that wallets are frozen.

#### Flow
- Consume: `user.created.v1` → Auto-create wallet
- Consume: `user.verified.v1` → Increase withdrawal limits
- Consume: `user.suspended.v1` → Freeze wallet
- Consume: `user.deleted.v1` → Process final payout, close wallet

#### Projections
- wallet_status_read

#### Events Consumed
- user.created.v1
- user.verified.v1
- user.suspended.v1
- user.deleted.v1

#### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 180ms

---

### 25.3 Subscription Events → Billing

#### User Stories
- As a **system**, I want to **consume subscription events** so that billing is triggered.
- As a **system**, I want to **consume connects.purchased events** so that balance is updated.

#### Flow
- Consume: `subscription.created.v1` → Schedule first charge
- Consume: `subscription.renewed.v1` → Process renewal payment
- Consume: `connects.purchased.v1` → Credit connect balance

#### Projections
- subscription_billing_read

#### Events Consumed
- subscription.created.v1
- subscription.renewed.v1
- connects.purchased.v1

#### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 200ms

---

### 25.4 Dispute Events → Financial Holds

#### User Stories
- As a **system**, I want to **consume dispute.opened events** so that payment holds are placed.
- As a **system**, I want to **consume dispute.resolved events** so that holds are released.

#### Flow
- Consume: `dispute.opened.v1` → Place payment hold
- Consume: `dispute.resolved.v1` → Release hold, process refund if needed

#### Projections
- dispute_holds_read

#### Events Consumed
- dispute.opened.v1
- dispute.resolved.v1

#### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 150ms

---

### 25.5 Admin Events → Platform Operations

#### User Stories
- As a **system**, I want to **consume admin.config.updated events** so that fee rules are refreshed.
- As a **system**, I want to **consume admin.feature_flag events** so that features are toggled.

#### Flow
- Consume: `admin.config.updated.v1` → Refresh config cache
- Consume: `admin.feature_flag.updated.v1` → Refresh feature flags
- Consume: `admin.payment.refunded.v1` → Process manual refund

#### Projections
- service_config_read

#### Events Consumed
- admin.config.updated.v1
- admin.feature_flag.updated.v1
- admin.payment.refunded.v1

#### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 120ms

---

## **END OF FINANCIAL-BE USER STORIES**

**Total Domains Covered:** 25  
**Total Sections:** 60+  
**Total User Stories:** 500+  
**Total Flows:** 400+  
**Total Events:** 300+  

All stories follow the pattern: **Stories → Flow → Projections → Events → RBAC/SLO**  
All flows include: **idempotent write-path, event envelope, non-PII payloads**  
All components align with: **folder structure, events catalog, platform conventions**

### Summary of Coverage

✅ **Core Wallet & Balance** - Multi-currency wallets, balance tracking, reservations  
✅ **Transaction & Ledger** - Double-entry bookkeeping, immutable journal, reconciliation  
✅ **Payment Processing** - Multiple gateways, 3DS, retries, receipts  
✅ **Escrow Management** - Fund holds, releases, pro-rata, refunds  
✅ **Payout Processing** - Batch processing, multiple methods, scheduling  
✅ **Invoice Management** - Generation, PDF, reminders, discounts  
✅ **Platform Fees** - Tiered fees, volume discounts, coupons, experiments  
✅ **Refund Processing** - Full/partial refunds, pro-rata calculations  
✅ **Tax Management** - VAT, withholding, 1099 generation, jurisdictions  
✅ **Foreign Exchange** - Currency conversion, rate locking, margins  
✅ **Risk & Fraud** - Fraud detection, risk scores, holds, chargebacks  
✅ **Bonus & Incentives** - Performance bonuses, referrals, vesting  
✅ **Expense Reimbursement** - Submission, approval, reimbursement  
✅ **Payment Schedules** - Recurring payments, automation, reminders  
✅ **Reconciliation** - Daily reconciliation, discrepancy detection  
✅ **Reporting** - Financial reports, dashboards, exports  
✅ **Subscription Billing** - Recurring charges, prorations, upgrades  
✅ **Connects Purchase** - Package sales, balance tracking, refunds  
✅ **Withdrawal Limits** - KYC-based limits, tracking, overrides  
✅ **Multi-Currency** - Currency wallets, transfers, preferences  
✅ **Gateway Integration** - Webhooks, signature verification, retry  
✅ **Payment Disputes** - Dispute creation, investigation, resolution  
✅ **Marketplace Commissions** - Split payments, tiered rates  
✅ **Promotional Credits** - Credits issuance, coupon redemption, expiry  
✅ **Bank Verification** - Micro-deposits, instant verification  
✅ **Financial Analytics** - Earnings/spending analytics, platform KPIs  
✅ **Event Consumers** - Integration with all platform services