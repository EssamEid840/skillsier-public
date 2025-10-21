# FINANCIAL-BE: Post-Review Database Modifications

This document packages the **schema fixes and enhancements** we discussed into a single, copy‑pasteable reference.  
Each section includes runnable SQL in the **recommended execution order** (safe to run multiple times thanks to `IF EXISTS` guards and idempotent patterns).

---

## 0) Pre-reqs
```sql
-- Ensure required extensions for hashing and UUID helpers are present.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
```
---

## 1) `fraud_alerts` → Fix incorrect FK and add proper column

> The original schema accidentally linked `fraud_alerts.id` to `risk_assessments(id)`.  
> We drop any existing FK with that name and add the correct nullable `risk_assessment_id` column + FK.

```sql
ALTER TABLE fraud_alerts
  DROP CONSTRAINT IF EXISTS fk_fraud_alerts_risk_assessment;

ALTER TABLE fraud_alerts
  ADD COLUMN IF NOT EXISTS risk_assessment_id UUID,
  ADD CONSTRAINT fk_fraud_alerts_risk_assessment
    FOREIGN KEY (risk_assessment_id) REFERENCES risk_assessments(id) ON DELETE SET NULL;
```
---

## 2) Transactions immutability trigger (correct field names, immutable-core enforcement)

> Prevents post-insert edits to immutable core fields; allows lifecycle/status/timestamps/metadata updates.

```sql
CREATE OR REPLACE FUNCTION enforce_transaction_immutability()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.is_immutable THEN
    IF NEW.transaction_number <> OLD.transaction_number
       OR NEW.amount <> OLD.amount
       OR NEW.currency <> OLD.currency
       OR COALESCE(NEW.debit_wallet_id,'00000000-0000-0000-0000-000000000000') <>
          COALESCE(OLD.debit_wallet_id,'00000000-0000-0000-0000-000000000000')
       OR COALESCE(NEW.credit_wallet_id,'00000000-0000-0000-0000-000000000000') <>
          COALESCE(OLD.credit_wallet_id,'00000000-0000-0000-0000-000000000000')
       OR NEW.fee_amount <> OLD.fee_amount
       OR NEW.net_amount <> OLD.net_amount
       OR COALESCE(NEW.payment_method_id,'00000000-0000-0000-0000-000000000000') <>
          COALESCE(OLD.payment_method_id,'00000000-0000-0000-0000-000000000000')
       OR NEW.payment_gateway <> OLD.payment_gateway
       OR NEW.gateway_transaction_id <> OLD.gateway_transaction_id
       OR NEW.initiated_by <> OLD.initiated_by
       OR NEW.reference_type <> OLD.reference_type
       OR COALESCE(NEW.reference_id,'00000000-0000-0000-0000-000000000000') <>
          COALESCE(OLD.reference_id,'00000000-0000-0000-0000-000000000000')
    THEN
      RAISE EXCEPTION 'Cannot modify immutable transaction fields';
    END IF;
  END IF;
  RETURN NEW;
END; $$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_enforce_transaction_immutability ON transactions;
CREATE TRIGGER trg_enforce_transaction_immutability
  BEFORE UPDATE ON transactions
  FOR EACH ROW
  EXECUTE FUNCTION enforce_transaction_immutability();
```
---

## 3) Wallet balance update trigger (fire on first transition to COMPLETED)

> Makes balance updates run on insert **and** when status transitions to `COMPLETED` (idempotent).

```sql
CREATE OR REPLACE FUNCTION update_wallet_balance()
RETURNS TRIGGER AS $$
BEGIN
  IF (TG_OP = 'INSERT' AND NEW.status = 'COMPLETED')
     OR (TG_OP = 'UPDATE' AND OLD.status IS DISTINCT FROM 'COMPLETED' AND NEW.status = 'COMPLETED') THEN

    IF NEW.debit_wallet_id IS NOT NULL THEN
      UPDATE wallets
         SET available_balance = available_balance - NEW.amount,
             updated_at = CURRENT_TIMESTAMP
       WHERE id = NEW.debit_wallet_id;
    END IF;

    IF NEW.credit_wallet_id IS NOT NULL THEN
      UPDATE wallets
         SET available_balance = available_balance + NEW.net_amount,
             updated_at = CURRENT_TIMESTAMP
       WHERE id = NEW.credit_wallet_id;
    END IF;
  END IF;
  RETURN NEW;
END; $$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_transaction_update_wallet ON transactions;
CREATE TRIGGER trg_transaction_update_wallet
  AFTER INSERT OR UPDATE OF status ON transactions
  FOR EACH ROW
  WHEN (NEW.status = 'COMPLETED')
  EXECUTE FUNCTION update_wallet_balance();
```
---

## 4) Prevent overdrafts at completion time

> Hard stop when debit wallet lacks funds at the moment of completion.

```sql
CREATE OR REPLACE FUNCTION prevent_overdrafts()
RETURNS TRIGGER AS $$
DECLARE bal BIGINT;
BEGIN
  IF (TG_OP = 'INSERT' AND NEW.status = 'COMPLETED')
     OR (TG_OP = 'UPDATE' AND OLD.status IS DISTINCT FROM 'COMPLETED' AND NEW.status = 'COMPLETED') THEN
    IF NEW.debit_wallet_id IS NOT NULL THEN
      SELECT available_balance INTO bal FROM wallets WHERE id = NEW.debit_wallet_id FOR UPDATE;
      IF bal < NEW.amount THEN
        RAISE EXCEPTION 'Insufficient funds in debit wallet %', NEW.debit_wallet_id;
      END IF;
    END IF;
  END IF;
  RETURN NEW;
END; $$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_transactions_prevent_overdrafts ON transactions;
CREATE TRIGGER trg_transactions_prevent_overdrafts
  BEFORE INSERT OR UPDATE OF status ON transactions
  FOR EACH ROW
  EXECUTE FUNCTION prevent_overdrafts();
```
---

## 5) Arithmetic consistency checks (net/fees)

> Locks in definitions to prevent drift.

```sql
ALTER TABLE transactions
  DROP CONSTRAINT IF EXISTS chk_transactions_net_amount,
  ADD  CONSTRAINT chk_transactions_net_amount CHECK (net_amount = amount - fee_amount);

ALTER TABLE payments
  DROP CONSTRAINT IF EXISTS chk_payments_net_amount,
  DROP CONSTRAINT IF EXISTS chk_payments_total_fees,
  ADD  CONSTRAINT chk_payments_total_fees   CHECK (total_fees = platform_fee + processing_fee),
  ADD  CONSTRAINT chk_payments_net_amount   CHECK (net_amount = amount - total_fees);
```
---

## 6) Strengthen referential integrity (payment methods & wallets)

```sql
-- Payment method references
ALTER TABLE payments
  ADD CONSTRAINT IF NOT EXISTS fk_payments_payment_method
  FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id) ON DELETE RESTRICT;

ALTER TABLE payouts
  ADD CONSTRAINT IF NOT EXISTS fk_payouts_payment_method
  FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id) ON DELETE RESTRICT;

-- Wallet references used by transactions & escrow flows
ALTER TABLE transactions
  ADD CONSTRAINT IF NOT EXISTS fk_transactions_debit_wallet
    FOREIGN KEY (debit_wallet_id)  REFERENCES wallets(id) ON DELETE RESTRICT,
  ADD CONSTRAINT IF NOT EXISTS fk_transactions_credit_wallet
    FOREIGN KEY (credit_wallet_id) REFERENCES wallets(id) ON DELETE RESTRICT;

ALTER TABLE escrow_releases
  ADD CONSTRAINT IF NOT EXISTS fk_escrow_releases_wallet
  FOREIGN KEY (released_to) REFERENCES wallets(id) ON DELETE RESTRICT;

ALTER TABLE escrow_holds
  ADD CONSTRAINT IF NOT EXISTS fk_escrow_holds_wallet
  FOREIGN KEY (released_to) REFERENCES wallets(id) ON DELETE SET NULL;
```
---

## 7) Payment attempts uniqueness (retry tracking)

```sql
ALTER TABLE payment_attempts
  ADD CONSTRAINT IF NOT EXISTS uk_payment_attempts UNIQUE (payment_id, attempt_number);
```
---

## 8) Ledger journal hash chain (auto-compute `entry_hash`/`prev_hash`)

> Computes a deterministic SHA‑256 over key fields + `prev_hash`.  
> Uses the last entry’s hash as `prev_hash` when absent.

```sql
CREATE OR REPLACE FUNCTION ledger_journal_hash()
RETURNS TRIGGER AS $$
DECLARE payload TEXT;
BEGIN
  IF NEW.prev_hash IS NULL THEN
    SELECT entry_hash INTO NEW.prev_hash
    FROM ledger_journal_entries
    ORDER BY entry_number DESC
    LIMIT 1;
  END IF;

  payload := COALESCE(NEW.transaction_id::text,'') || '|' ||
             NEW.debit_account || '|' || NEW.credit_account || '|' ||
             NEW.amount::text || '|' || NEW.currency || '|' ||
             NEW.effective_at::text || '|' || COALESCE(NEW.prev_hash,'');

  NEW.entry_hash := encode(digest(payload, 'sha256'), 'hex');
  RETURN NEW;
END; $$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_ledger_journal_hash ON ledger_journal_entries;
CREATE TRIGGER trg_ledger_journal_hash
  BEFORE INSERT ON ledger_journal_entries
  FOR EACH ROW
  EXECUTE FUNCTION ledger_journal_hash();
```
---

## 9) PCI storage types (ciphertext as BYTEA)

> Store encrypted secrets/tokens as `BYTEA`. (Ciphertext productionization left to app layer/key mgmt.)

```sql
ALTER TABLE payment_methods
  ALTER COLUMN gateway_token TYPE BYTEA USING gateway_token::bytea;

ALTER TABLE gateway_configurations
  ALTER COLUMN api_key_encrypted TYPE BYTEA USING api_key_encrypted::bytea,
  ALTER COLUMN api_secret_encrypted TYPE BYTEA USING api_secret_encrypted::bytea;
```
---

## 10) (Optional) Robustness & hygiene notes

- `risk_assessments` uniqueness on `(assessment_type, subject_id, assessed_at)` can collide at high throughput. Consider a surrogate `assessment_seq` (BIGSERIAL) or lower‑precision bucketing.  
- Add idempotency scope guard if your `idempotency_key` is per‑actor (e.g., `UNIQUE (initiated_by, idempotency_key)`).
- Review FK `ON DELETE` behaviors to match lifecycle expectations (e.g., `SET NULL` vs `RESTRICT`).

---

## Execution order (recommended)

1. **0 Pre-reqs**  
2. **1 fraud_alerts FK fix**  
3. **5 arithmetic checks**  
4. **6 referential integrity**  
5. **7 payment attempts uniqueness**  
6. **2 immutability trigger**  
7. **3 wallet balance trigger**  
8. **4 prevent overdrafts trigger**  
9. **8 ledger journal hash trigger**  
10. **9 PCI storage types**  
11. **10 notes (review-only)**

---

### Rollback helpers (quick remove)
```sql
-- Drop added constraints/triggers if you need to roll back quickly.
ALTER TABLE payment_attempts DROP CONSTRAINT IF EXISTS uk_payment_attempts;

ALTER TABLE payments  DROP CONSTRAINT IF EXISTS fk_payments_payment_method;
ALTER TABLE payouts   DROP CONSTRAINT IF EXISTS fk_payouts_payment_method;

ALTER TABLE transactions
  DROP CONSTRAINT IF EXISTS fk_transactions_debit_wallet,
  DROP CONSTRAINT IF EXISTS fk_transactions_credit_wallet;

ALTER TABLE escrow_releases DROP CONSTRAINT IF EXISTS fk_escrow_releases_wallet;
ALTER TABLE escrow_holds    DROP CONSTRAINT IF EXISTS fk_escrow_holds_wallet;

DROP TRIGGER IF EXISTS trg_enforce_transaction_immutability ON transactions;
DROP FUNCTION IF EXISTS enforce_transaction_immutability();

DROP TRIGGER IF EXISTS trg_transaction_update_wallet ON transactions;
DROP FUNCTION IF EXISTS update_wallet_balance();

DROP TRIGGER IF EXISTS trg_transactions_prevent_overdrafts ON transactions;
DROP FUNCTION IF EXISTS prevent_overdrafts();

DROP TRIGGER IF EXISTS trg_ledger_journal_hash ON ledger_journal_entries;
DROP FUNCTION IF EXISTS ledger_journal_hash();

ALTER TABLE fraud_alerts DROP CONSTRAINT IF EXISTS fk_fraud_alerts_risk_assessment;
ALTER TABLE fraud_alerts DROP COLUMN IF EXISTS risk_assessment_id;

-- (PCI type changes are destructive; plan reversions carefully.)
```
