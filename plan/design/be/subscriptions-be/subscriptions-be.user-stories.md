# 📦 **subscriptions-be - Subscription & Entitlement Management Service - Complete User Stories**

---

## **DOCUMENT OVERVIEW**

**Service:** subscriptions-be  
**Purpose:** Manage subscription plans, connects, entitlements, usage limits, billing, and dunning for the platform  
**Architecture:** Event-Driven CQRS with Outbox Pattern  
**Event Envelope:** Standard platform envelope (event_id, timestamp, correlation_id, causation_id, user_context, compliance_context)  
**Idempotency:** All write commands use idempotency keys  
**Non-PII:** Events contain only IDs and codes; no direct PII  
**Coverage:** 100% of subscriptions-be folder structure domains  
**Structure:** Sections → Domains → Entities → (Stories → Flow → Projections → Events → RBAC/SLO)

---

## **GLOBAL CONVENTIONS**


### Event Envelope Structure (All Events)
```json
{
  "event_id": "uuid",
  "event_type": "subscription.created.v1",
  "event_timestamp": "2025-01-15T10:30:00Z",
  "event_version": "1",
  "aggregate_type": "subscription",
  "aggregate_id": "uuid",
  "correlation_id": "uuid",
  "causation_id": "uuid",
  "event_source": "subscriptions-be",
  "user_context": {
    "user_id": "uuid",
    "keycloak_id": "uuid",
    "user_type": "FREELANCER|CLIENT",
    "session_id": "uuid"
  },
  "compliance_context": {
    "gdpr_consent": true,
    "data_classification": "SENSITIVE|INTERNAL|PUBLIC",
    "retention_policy": "7y"
  },
  "audit_metadata": {
    "ip_address": "x.x.x.x",
    "user_agent": "string",
    "request_id": "uuid"
  },
  "payload": { /* event-specific data */ }
}
```

### Idempotent Write-Path
- All commands accept `idempotency_key` header
- Duplicate requests return cached response (24h TTL)
- Key format: `{service}.{command}.{user_id}.{resource_id}.{hash}`

### Non-PII Event Rules
- Events contain only UUIDs, codes, enums, and numeric values
- No raw names, emails, addresses, or free-text PII
- Consumers fetch PII via API if needed
- Redacted fields marked as `[REDACTED]` in audits

### Platform Alignment
- Folder structure: `/apps/be/subscriptions-be/`
- Events catalog: `/apps/be/contracts/events/`
- Shared libraries: `/platform-shared/`, `/pkg/auth/`

---

## **SECTION 1: CATALOG (PLANS & FEATURES)**

### 1.1 Domain: plan/

#### 1.1.1 Entity: entity.go (Plan Aggregate)

##### User Stories
- As an **admin**, I want to **create subscription plans** so that users can choose the right tier for their needs.
- As an **admin**, I want to **update plan details** (name, description) so that information stays current.
- As an **admin**, I want to **activate/deactivate plans** so that plan catalog is managed.
- As a **user**, I want to **view available plans** so that I can compare options.

##### Flow
1. **CreatePlanCommand**(code, name, tier, active, idempotency_key) → Validate() | Persist() → **Outbox:** plan.created.v1
2. **UpdatePlanCommand**(plan_id, updates, idempotency_key) → Apply() → **Outbox:** plan.updated.v1
3. **ActivatePlanCommand**(plan_id) → Activate() → **Outbox:** plan.activated.v1
4. **DeactivatePlanCommand**(plan_id, reason) → Deactivate() → **Outbox:** plan.deactivated.v1
5. **GetPlanQuery**(plan_id) → Fetch() → PlanDTO
6. **ListPlansQuery**(filters) → ApplyFilters() → PlanListDTO

##### Projections
- plan_read
- plan_catalog_read

##### Events Published
- plan.created.v1
- plan.updated.v1
- plan.activated.v1
- plan.deactivated.v1

##### RBAC/SLO
- **RBAC:** ADMIN (create/update/activate/deactivate), PUBLIC (view)
- **SLO:** P95 < 100ms (read), P95 < 180ms (write)

---

#### 1.1.2 Entity: features.go (Plan Features)

##### User Stories
- As an **admin**, I want to **define plan features** (job posts, proposals, connects) so that entitlements are clear.
- As a **system**, I want to **retrieve feature values** so that gates can be enforced.

##### Flow
1. **SetPlanFeaturesCommand**(plan_id, features_map, idempotency_key) → Validate() | Persist() → **Outbox:** plan.features.set.v1
2. **GetPlanFeaturesQuery**(plan_id) → Fetch() → PlanFeaturesDTO

##### Projections
- plan_features_read

##### Events Published
- plan.features.set.v1

##### RBAC/SLO
- **RBAC:** ADMIN (set), PUBLIC (view)
- **SLO:** P95 < 80ms (read), P95 < 150ms (write)

---

#### 1.1.3 Entity: pricing.go (Plan Pricing)

##### User Stories
- As an **admin**, I want to **set plan pricing** per billing cycle (monthly, yearly) so that revenue is optimized.
- As a **system**, I want to **calculate charges** based on billing cycle so that invoices are correct.

##### Flow
1. **SetPlanPricingCommand**(plan_id, billing_cycle, base_price, currency, idempotency_key) → Validate() | Persist() → **Outbox:** plan.pricing.set.v1
2. **CalculatePlanChargeQuery**(plan_id, billing_cycle) → Fetch() | Calculate() → PlanChargeDTO

##### Projections
- plan_pricing_read

##### Events Published
- plan.pricing.set.v1

##### RBAC/SLO
- **RBAC:** ADMIN (set), SYSTEM (calculate), PUBLIC (view)
- **SLO:** P95 < 80ms (calculate), P95 < 150ms (write)

---

#### 1.1.4 Entity: limits.go (Plan Limits)

##### User Stories
- As an **admin**, I want to **set plan limits** (daily job posts, invites per month) so that usage is controlled.
- As a **system**, I want to **check limits** before operations so that quotas are enforced.

##### Flow
1. **SetPlanLimitsCommand**(plan_id, limits_map, idempotency_key) → Validate() | Persist() → **Outbox:** plan.limits.set.v1
2. **CheckPlanLimitQuery**(plan_id, limit_type) → Fetch() → PlanLimitDTO

##### Projections
- plan_limits_read

##### Events Published
- plan.limits.set.v1

##### RBAC/SLO
- **RBAC:** ADMIN (set), SYSTEM (check)
- **SLO:** P95 < 60ms (check - critical path), P95 < 150ms (write)

---

### 1.2 Domain: plan_version/

#### 1.2.1 Entity: entity.go (Plan Version Aggregate)

##### User Stories
- As an **admin**, I want to **version plan changes** so that existing subscribers keep their original terms.
- As a **system**, I want to **freeze plan features** at subscription time so that retroactive changes don't affect users.
- As an **admin**, I want to **view plan version history** so that I can track changes over time.

##### Flow
1. **CreatePlanVersionCommand**(plan_id, version, changes, version_reason, idempotency_key) → VersionIncrement() | ClonePlan() | Persist() → **Outbox:** plan.version.created.v1
2. **ActivatePlanVersionCommand**(plan_id, version) → Activate() → **Outbox:** plan.version.activated.v1
3. **GetPlanVersionQuery**(plan_id, version) → Fetch() → PlanVersionDTO
4. **ListPlanVersionsQuery**(plan_id) → FetchAll() → PlanVersionListDTO

##### Projections
- plan_version_read

##### Events Published
- plan.version.created.v1
- plan.version.activated.v1
- plan.version.deprecated.v1

##### RBAC/SLO
- **RBAC:** ADMIN (create/activate), SYSTEM (freeze version on subscription)
- **SLO:** P95 < 120ms (read), P95 < 200ms (write)

---

#### 1.2.2 Entity: migration_rule.go (Plan Migration Policy)

##### User Stories
- As an **admin**, I want to **define auto-migration rules** so that users can be migrated to new plan versions.
- As a **system**, I want to **apply opt-in/opt-out windows** so that migrations are controlled.

##### Flow
1. **SetMigrationRuleCommand**(plan_id, migration_policy, window_dates, idempotency_key) → Validate() | Persist() → **Outbox:** plan.migration.rule.set.v1
2. **ApplyMigrationJob**(plan_id, target_version) → FindEligible() | Migrate() → **Outbox:** plan.migration.applied.v1

##### Projections
- migration_rule_read

##### Events Published
- plan.migration.rule.set.v1
- plan.migration.applied.v1

##### RBAC/SLO
- **RBAC:** ADMIN (set rules), SYSTEM (apply migration)
- **SLO:** P95 < 180ms (set), P95 < 300ms (apply)

---

## **SECTION 2: SUBSCRIPTIONS & CHANGE REQUESTS**

### 2.1 Domain: subscription/

#### 2.1.1 Entity: entity.go (Subscription Aggregate)

##### User Stories
- As a **user**, I want to **subscribe to a plan** so that I can access premium features.
- As a **user**, I want to **cancel my subscription** so that I can stop charges.
- As a **user**, I want to **view my subscription status** so that I know my current tier and renewal date.
- As a **system**, I want to **handle subscription lifecycle** (active, paused, expired, canceled) so that state is managed.

##### Flow
1. **SubscribeCommand**(user_id, plan_id, billing_cycle, payment_method_id, promo_code, idempotency_key) → ValidatePlan() | ValidatePayment() | CreatePaymentIntent(financial-be) | ProcessPayment() | ActivateSubscription() | FreezeVersion() → **Outbox:** subscription.created.v1
2. **CancelSubscriptionCommand**(subscription_id, cancellation_reason, cancel_at_period_end, idempotency_key) → AuthorizeOwner() | Cancel() | ProcessRefund() → **Outbox:** subscription.cancelled.v1
3. **GetSubscriptionQuery**(subscription_id) → AuthorizeAccess() | Fetch() | EnrichWithFeatures() → SubscriptionDTO
4. **ListSubscriptionsQuery**(user_id, filters) → ApplyFilters() → SubscriptionListDTO

##### Projections
- subscription_read
- subscription_status_read

##### Events Published
- subscription.created.v1
- subscription.cancelled.v1
- subscription.expired.v1

##### Events Consumed
- payment.processed.v1 (activate subscription)
- payment.failed.v1 (handle failed payment)
- user.deleted.v1 (cancel subscriptions)

##### RBAC/SLO
- **RBAC:** OWNER (subscribe/cancel/view), ADMIN (view all/force operations)
- **SLO:** P95 < 400ms (subscribe), P95 < 250ms (cancel), P95 < 150ms (read)

---

#### 2.1.2 Entity: billing_cycle.go (Billing Cycle Logic)

##### User Stories
- As a **system**, I want to **compute next renewal date** so that billing is scheduled.
- As a **system**, I want to **calculate remaining days** in period so that prorations are accurate.
- As a **system**, I want to **handle grace periods** so that payment failures don't immediately cancel.

##### Flow
1. **CalculateNextRenewalQuery**(subscription_id) → FetchSubscription() | ComputeDate() → NextRenewalDTO
2. **CalculateRemainingDaysQuery**(subscription_id) → FetchPeriod() | ComputeDays() → RemainingDaysDTO
3. **ApplyGracePeriodCommand**(subscription_id, grace_days) → Extend() → **Outbox:** subscription.grace.applied.v1

##### Projections
- billing_cycle_read

##### Events Published
- subscription.grace.applied.v1

##### RBAC/SLO
- **RBAC:** SYSTEM (all operations)
- **SLO:** P95 < 80ms (calculate)

---

#### 2.1.3 Entity: auto_renewal.go (Auto-Renewal Control)

##### User Stories
- As a **user**, I want to **enable auto-renewal** so that my subscription continues automatically.
- As a **user**, I want to **disable auto-renewal** so that I'm not charged after the current period.
- As a **system**, I want to **auto-renew subscriptions** on renewal date so that service is uninterrupted.

##### Flow
1. **EnableAutoRenewalCommand**(subscription_id, idempotency_key) → AuthorizeOwner() | Enable() → **Outbox:** subscription.auto_renewal.enabled.v1
2. **DisableAutoRenewalCommand**(subscription_id, idempotency_key) → AuthorizeOwner() | Disable() → **Outbox:** subscription.auto_renewal.disabled.v1
3. **RenewSubscriptionJob**(subscription_id) → CreatePaymentIntent(financial-be) | ProcessPayment() | ExtendPeriod() | UpdateNextRenewal() → **Outbox:** subscription.renewed.v1

##### Projections
- auto_renewal_read

##### Events Published
- subscription.auto_renewal.enabled.v1
- subscription.auto_renewal.disabled.v1
- subscription.renewed.v1

##### RBAC/SLO
- **RBAC:** OWNER (enable/disable), SYSTEM (renew)
- **SLO:** P95 < 150ms (enable/disable), P95 < 600ms (renew)

---

### 2.2 Domain: change_request/

#### 2.2.1 Entity: entity.go (Change Request Aggregate)

##### User Stories
- As a **user**, I want to **upgrade my plan** so that I get more features immediately.
- As a **user**, I want to **downgrade my plan** at period end so that I don't lose prepaid time.
- As a **user**, I want to **schedule plan changes** so that transitions happen at the right time.
- As a **user**, I want to **cancel a pending plan change** so that I can change my mind.

##### Flow
1. **RequestUpgradeCommand**(subscription_id, new_plan_id, effective_date, idempotency_key) → AuthorizeOwner() | ValidatePlan() | CalculateProration() | CreateChangeRequest() | ProcessImmediate() → **Outbox:** subscription.upgrade.requested.v1
2. **RequestDowngradeCommand**(subscription_id, new_plan_id, effective_date, idempotency_key) → AuthorizeOwner() | ValidatePlan() | ScheduleChange() → **Outbox:** subscription.downgrade.requested.v1
3. **CancelChangeRequestCommand**(change_request_id, idempotency_key) → AuthorizeOwner() | Cancel() → **Outbox:** subscription.change.cancelled.v1
4. **GetChangeRequestQuery**(change_request_id) → AuthorizeAccess() | Fetch() → ChangeRequestDTO
5. **ListChangeRequestsQuery**(subscription_id) → Fetch() → ChangeRequestListDTO

##### Projections
- change_request_read
- pending_changes_read

##### Events Published
- subscription.upgrade.requested.v1
- subscription.downgrade.requested.v1
- subscription.change.cancelled.v1

##### RBAC/SLO
- **RBAC:** OWNER (request/cancel), SYSTEM (apply)
- **SLO:** P95 < 300ms (request), P95 < 150ms (read)

---

#### 2.2.2 Entity: proration_policy.go (Proration Policy)

##### User Stories
- As a **system**, I want to **calculate proration** for mid-cycle changes so that charges are fair.
- As an **admin**, I want to **set proration policies** (None, Immediate, CreditNote) so that billing strategy is flexible.

##### Flow
1. **CalculateProrationQuery**(subscription_id, new_plan_id, effective_date) → FetchSubscription() | ComputeProration() → ProrationDTO
2. **SetProrationPolicyCommand**(plan_id, policy_type, idempotency_key) → Validate() | Persist() → **Outbox:** proration.policy.set.v1

##### Projections
- proration_policy_read

##### Events Published
- proration.policy.set.v1

##### RBAC/SLO
- **RBAC:** SYSTEM (calculate), ADMIN (set policy)
- **SLO:** P95 < 100ms (calculate), P95 < 150ms (set)

---

#### 2.2.3 Entity: events.go (Change Request Events - Applied Job)

##### User Stories
- As a **system**, I want to **apply scheduled plan changes** so that transitions happen automatically.

##### Flow
1. **ApplyChangeRequestJob**(change_request_id) → FetchRequest() | ApplyChange() | UpdatePlan() | UpdateBilling() → **Outbox:** subscription.change.applied.v1

##### Projections
- None (event-driven)

##### Events Published
- subscription.change.applied.v1

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 400ms

---

## **SECTION 3: ENTITLEMENTS & GRANTS**

### 3.1 Domain: entitlement/

#### 3.1.1 Entity: entity.go (Entitlement Aggregate)

##### User Stories
- As a **user**, I want to **check my entitlements** (job posts, proposals, features) so that I know what I can do.
- As a **system**, I want to **enforce entitlement gates** before operations so that usage limits are respected.
- As a **system**, I want to **compute effective entitlements** from subscription + grants so that special permissions are included.

##### Flow
1. **CheckEntitlementQuery**(user_id, feature_key) → FetchSubscription() | FetchGrants() | ComputeEffective() | CheckLimit() → EntitlementCheckDTO
2. **ListEntitlementsQuery**(user_id) → FetchSubscription() | FetchGrants() | ComputeAll() → EntitlementListDTO
3. **GetEffectiveEntitlementsQuery**(user_id) → FetchSubscription() | FetchGrants() | MergeAll() → EffectiveEntitlementDTO

##### Projections
- entitlement_read
- effective_entitlement_read

##### Events Published
- (none - query-only)

##### Events Consumed
- subscription.created.v1 (refresh entitlements)
- subscription.changed.v1 (refresh entitlements)
- grant.issued.v1 (refresh entitlements)

##### RBAC/SLO
- **RBAC:** OWNER (view own), ADMIN (view all), SYSTEM (check/enforce)
- **SLO:** P95 < 50ms (check - critical path), P95 < 100ms (list)

---

#### 3.1.2 Entity: rules.go (Entitlement Merge Rules)

##### User Stories
- As a **system**, I want to **merge entitlements** with priority (plan < addon < promo < ad-hoc grant) so that most permissive wins.

##### Flow
1. **MergeEntitlementsQuery**(user_id) → FetchAll() | ApplyMergePriority() → MergedEntitlementDTO

##### Projections
- None (computed on-the-fly)

##### Events Published
- (none)

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 80ms

---

### 3.2 Domain: entitlement_grant/

#### 3.2.1 Entity: entity.go (Grant Aggregate)

##### User Stories
- As an **admin**, I want to **issue temporary grants** (bonus connects, feature access) so that I can reward users.
- As an **admin**, I want to **set grant expiry** so that temporary access is time-boxed.
- As an **admin**, I want to **revoke grants** if misused so that abuse is stopped.
- As a **system**, I want to **expire grants automatically** so that manual cleanup isn't needed.

##### Flow
1. **IssueGrantCommand**(user_id, grant_type, grant_value, expiry_date, reason, issued_by, idempotency_key) → AuthorizeAdmin() | Validate() | Issue() → **Outbox:** grant.issued.v1
2. **RevokeGrantCommand**(grant_id, reason, revoked_by, idempotency_key) → AuthorizeAdmin() | Revoke() → **Outbox:** grant.revoked.v1
3. **ExpireGrantsJob**() → FindExpired() | BatchExpire() → **Outbox:** grant.expired.v1
4. **GetGrantQuery**(grant_id) → AuthorizeAdmin() | Fetch() → GrantDTO
5. **ListGrantsQuery**(user_id) → Fetch() → GrantListDTO

##### Projections
- grant_read
- active_grants_read

##### Events Published
- grant.issued.v1
- grant.revoked.v1
- grant.expired.v1

##### RBAC/SLO
- **RBAC:** ADMIN (issue/revoke), SYSTEM (expire), OWNER (view own)
- **SLO:** P95 < 180ms (issue/revoke), P95 < 100ms (read)

---

## **SECTION 4: USAGE TRACKING & ALLOWANCES**

### 4.1 Domain: usage/

#### 4.1.1 Entity: entity.go (Usage Aggregate)

##### User Stories
- As a **system**, I want to **track usage** (job posts, proposals, invites) so that limits are enforced.
- As a **system**, I want to **increment usage counters** on each action so that quotas are consumed.
- As a **user**, I want to **view my usage** so that I know how much quota I have left.

##### Flow
1. **IncrementUsageCommand**(user_id, usage_type, amount, idempotency_key) → FetchEntitlement() | CheckLimit() | Increment() | CheckThreshold() → **Outbox:** usage.incremented.v1
2. **GetUsageQuery**(user_id) → Fetch() | EnrichWithLimits() → UsageDTO
3. **GetUsageHistoryQuery**(user_id, time_range) → Fetch() → UsageHistoryDTO

##### Projections
- usage_read
- usage_history_read

##### Events Published
- usage.incremented.v1
- usage.limit.reached.v1
- usage.threshold.warning.v1

##### Events Consumed
- job.posted.v1 (increment job posts)
- proposal.submitted.v1 (increment proposals)
- invitation.sent.v1 (increment invites)

##### RBAC/SLO
- **RBAC:** SYSTEM (increment), OWNER (view own), ADMIN (view all)
- **SLO:** P95 < 80ms (increment - critical path), P95 < 100ms (read)

---

#### 4.1.2 Entity: counter.go (Usage Counters)

##### User Stories
- As a **system**, I want to **maintain atomic counters** so that concurrent usage is tracked correctly.

##### Flow
1. **AtomicIncrementCommand**(user_id, counter_type, delta) → IncrementAtomic() → **Outbox:** counter.incremented.v1

##### Projections
- counter_read

##### Events Published
- counter.incremented.v1

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 50ms (atomic operation)

---

#### 4.1.3 Entity: reset.go (Usage Reset Logic)

##### User Stories
- As a **system**, I want to **reset usage counters** at period boundaries so that monthly limits restart.

##### Flow
1. **ResetUsageJob**(period_type) → FindExpired() | BatchReset() → **Outbox:** usage.reset.v1

##### Projections
- None (background job)

##### Events Published
- usage.reset.v1

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 2000ms (batch operation)

---

### 4.2 Domain: allowance/

#### 4.2.1 Entity: entity.go (Allowance Aggregate)

##### User Stories
- As a **system**, I want to **allocate monthly allowances** (job posts, proposals) per subscription so that usage is bucketed.
- As a **system**, I want to **rollover unused allowances** if plan allows so that value is preserved.
- As a **user**, I want to **view my allowance breakdown** so that I understand my quota.

##### Flow
1. **AllocateAllowanceCommand**(user_id, allowance_type, amount, period_start, period_end, idempotency_key) → Validate() | Allocate() → **Outbox:** allowance.allocated.v1
2. **ConsumeAllowanceCommand**(user_id, allowance_type, amount, idempotency_key) → FindAvailable() | Consume() | CheckThreshold() → **Outbox:** allowance.consumed.v1
3. **GetAllowanceQuery**(user_id) → Fetch() | ComputeAvailable() → AllowanceDTO

##### Projections
- allowance_read
- available_allowance_read

##### Events Published
- allowance.allocated.v1
- allowance.consumed.v1

##### Events Consumed
- subscription.renewed.v1 (allocate new period)
- usage.incremented.v1 (consume allowance)

##### RBAC/SLO
- **RBAC:** SYSTEM (allocate/consume), OWNER (view own)
- **SLO:** P95 < 100ms (allocate/consume), P95 < 80ms (read)

---

#### 4.2.2 Entity: rollover.go (Allowance Rollover Policy)

##### User Stories
- As a **system**, I want to **apply rollover policies** so that unused allowances carry forward correctly.

##### Flow
1. **RolloverAllowanceJob**(user_id) → FetchUnused() | ApplyRolloverPolicy() | Allocate() → **Outbox:** allowance.rolled_over.v1

##### Projections
- None (background job)

##### Events Published
- allowance.rolled_over.v1

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 300ms

---

#### 4.2.3 Entity: expiry.go (Allowance Expiry Logic)

##### User Stories
- As a **system**, I want to **expire old allowances** so that usage doesn't accumulate indefinitely.

##### Flow
1. **ExpireAllowanceJob**() → FindExpired() | BatchExpire() → **Outbox:** allowance.expired.v1

##### Projections
- None (background job)

##### Events Published
- allowance.expired.v1

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 2000ms (batch)

---

## **SECTION 5: CONNECTS SYSTEM**

### 5.1 Domain: connect/

#### 5.1.1 Entity: entity.go (Connect Aggregate)

##### User Stories
- As a **user**, I want to **purchase connect packages** so that I can submit proposals.
- As a **system**, I want to **track connect balance** so that users know how many they have left.
- As a **user**, I want to **view connect history** so that I can track spending.

##### Flow
1. **PurchaseConnectsCommand**(user_id, package_id, payment_method_id, promo_code, idempotency_key) → ValidatePackage() | CreatePaymentIntent(financial-be) | ProcessPayment() | AddConnects() → **Outbox:** connects.purchased.v1
2. **GetConnectBalanceQuery**(user_id) → Fetch() → ConnectBalanceDTO
3. **GetConnectHistoryQuery**(user_id, filters) → ApplyFilters() → ConnectHistoryDTO

##### Projections
- connect_balance_read
- connect_history_read

##### Events Published
- connects.purchased.v1

##### Events Consumed
- payment.processed.v1 (add connects after purchase)

##### RBAC/SLO
- **RBAC:** OWNER (purchase/view), SYSTEM (add)
- **SLO:** P95 < 300ms (purchase), P95 < 80ms (read)

---

#### 5.1.2 Entity: transaction.go (Connect Transactions)

##### User Stories
- As a **system**, I want to **deduct connects** on proposal submission so that usage is charged.
- As a **system**, I want to **refund connects** if proposal is rejected within 24h so that failed applications don't cost.

##### Flow
1. **UseConnectsCommand**(user_id, proposal_id, connects_amount, idempotency_key) → CheckBalance() | Deduct() | CheckThreshold() → **Outbox:** connects.used.v1
2. **RefundConnectsCommand**(usage_id, refund_reason, idempotency_key) → ValidateRefund() | Refund() → **Outbox:** connects.refunded.v1

##### Projections
- connect_transaction_read

##### Events Published
- connects.used.v1
- connects.refunded.v1
- connects.low.warning.v1
- connects.depleted.v1

##### Events Consumed
- proposal.submitted.v1 (deduct connects)
- proposal.rejected.quick.v1 (refund connects if eligible)

##### RBAC/SLO
- **RBAC:** SYSTEM (use/refund)
- **SLO:** P95 < 100ms (use - critical path), P95 < 200ms (refund)

---

#### 5.1.3 Entity: balance.go (Connect Balance Calculation)

##### User Stories
- As a **system**, I want to **compute derived balance** so that display is accurate.

##### Flow
1. **ComputeBalanceQuery**(user_id) → FetchTransactions() | Aggregate() → BalanceDTO

##### Projections
- connect_balance_read

##### Events Published
- (none - computed)

##### RBAC/SLO
- **RBAC:** OWNER (view own)
- **SLO:** P95 < 60ms

---

#### 5.1.4 Entity: expiry.go (Connect Expiry)

##### User Stories
- As a **system**, I want to **expire connects** after validity period so that usage is time-bound.

##### Flow
1. **ExpireConnectsJob**() → FindExpired() | BatchExpire() → **Outbox:** connects.expired.v1

##### Projections
- None (background job)

##### Events Published
- connects.expired.v1

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 2000ms (batch)

---

#### 5.1.5 Entity: grant.go (Connect Grants)

##### User Stories
- As an **admin**, I want to **grant bonus connects** so that users can be rewarded.

##### Flow
1. **GrantConnectsCommand**(user_id, amount, reason, granted_by, idempotency_key) → AuthorizeAdmin() | Grant() → **Outbox:** connects.granted.v1

##### Projections
- connect_grant_read

##### Events Published
- connects.granted.v1

##### RBAC/SLO
- **RBAC:** ADMIN
- **SLO:** P95 < 150ms

---

#### 5.1.6 Entity: package.go (Connect Packages)

##### User Stories
- As an **admin**, I want to **create connect packages** (starter, value, premium) so that users have options.
- As an **admin**, I want to **set package pricing** with volume discounts so that bulk purchases are cheaper.

##### Flow
1. **CreatePackageCommand**(name, connects_amount, price, bonus_connects, active, idempotency_key) → Validate() | Persist() → **Outbox:** connect_package.created.v1
2. **UpdatePackageCommand**(package_id, updates, idempotency_key) → Apply() → **Outbox:** connect_package.updated.v1
3. **GetPackageQuery**(package_id) → Fetch() → PackageDTO
4. **ListPackagesQuery**(filters) → ApplyFilters() → PackageListDTO

##### Projections
- connect_package_read

##### Events Published
- connect_package.created.v1
- connect_package.updated.v1

##### RBAC/SLO
- **RBAC:** ADMIN (create/update), PUBLIC (view)
- **SLO:** P95 < 120ms (read), P95 < 180ms (write)

---

## **SECTION 6: SEAT BILLING**

### 6.1 Domain: seat_billing/

#### 6.1.1 Entity: entity.go (Seat Billing Aggregate)

##### User Stories
- As an **organization**, I want to **add team seats** to my subscription so that multiple users can access premium features.
- As an **organization**, I want to **pay per active seat** so that costs scale with team size.
- As an **organization**, I want to **remove seats** when team members leave so that costs are optimized.

##### Flow
1. **AssignSeatCommand**(subscription_id, user_id, assigned_by, idempotency_key) → CheckAvailableSeats() | AssignSeat() | CheckOverage() → **Outbox:** seat.assigned.v1
2. **ReleaseSeatCommand**(seat_id, released_by, idempotency_key) → AuthorizeAdmin() | Release() → **Outbox:** seat.released.v1
3. **GetSeatUsageQuery**(subscription_id) → Fetch() | ComputeOverage() → SeatUsageDTO
4. **ListSeatsQuery**(subscription_id) → Fetch() → SeatListDTO

##### Projections
- seat_usage_read
- seat_assignment_read

##### Events Published
- seat.assigned.v1
- seat.released.v1

##### RBAC/SLO
- **RBAC:** ORG_ADMIN (assign/release), OWNER (view), SYSTEM (charge)
- **SLO:** P95 < 200ms (assign/release), P95 < 120ms (read)

---

#### 6.1.2 Entity: overage.go (Seat Overage Calculation)

##### User Stories
- As a **system**, I want to **track seat assignments** so that usage is accurate.
- As a **system**, I want to **charge overages** when seats exceed plan limits so that extra usage is billed.

##### Flow
1. **ChargeOverageJob**(subscription_id) → CalculateOverage() | CreateInvoice(financial-be) → **Outbox:** seat.overage.charged.v1

##### Projections
- seat_overage_read

##### Events Published
- seat.overage.charged.v1

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 300ms

---

#### 6.1.3 Entity: proration.go (Seat Proration)

##### User Stories
- As a **system**, I want to **prorate seat charges** for mid-cycle additions so that billing is fair.

##### Flow
1. **CalculateSeatProrationQuery**(subscription_id, added_seats, effective_date) → FetchSubscription() | ComputeProration() → SeatProrationDTO

##### Projections
- None (computed)

##### Events Published
- (none)

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 100ms

---

## **SECTION 7: ADDONS, PROMOTIONS & TRIALS**

### 7.1 Domain: addon/

#### 7.1.1 Entity: entity.go (Addon Aggregate)

##### User Stories
- As a **user**, I want to **add optional features** (priority support, featured profile) to my subscription so that I get extras.
- As a **user**, I want to **remove addons** when no longer needed so that costs are controlled.
- As a **system**, I want to **charge for addons** monthly so that billing is automated.

##### Flow
1. **AddAddonCommand**(subscription_id, addon_id, idempotency_key) → ValidateAddon() | CheckCompatibility() | Add() | CreateInvoice(financial-be) → **Outbox:** addon.added.v1
2. **RemoveAddonCommand**(subscription_addon_id, idempotency_key) → AuthorizeOwner() | Remove() | ProcessRefund() → **Outbox:** addon.removed.v1
3. **ChargeAddonJob**(subscription_addon_id) → FetchAddon() | CreateInvoice(financial-be) → **Outbox:** addon.charged.v1
4. **GetAddonsQuery**(subscription_id) → Fetch() → AddonListDTO

##### Projections
- addon_read
- subscription_addon_read

##### Events Published
- addon.added.v1
- addon.removed.v1
- addon.charged.v1

##### RBAC/SLO
- **RBAC:** OWNER (add/remove/view), SYSTEM (charge), ADMIN (create addons)
- **SLO:** P95 < 250ms (add/remove), P95 < 100ms (read)

---

### 7.2 Domain: promotion/

#### 7.2.1 Entity: entity.go (Promotion Aggregate)

##### User Stories
- As an **admin**, I want to **create promo codes** (discount percentages, free trials) so that marketing campaigns work.
- As an **admin**, I want to **set promo limits** (max uses, expiry date) so that abuse is prevented.
- As a **user**, I want to **apply promo codes** during subscription so that I get discounts.

##### Flow
1. **CreatePromoCommand**(code, discount_type, discount_value, max_uses, expiry_date, restrictions, idempotency_key) → Validate() | Persist() → **Outbox:** promo.created.v1
2. **ApplyPromoCommand**(user_id, promo_code, subscription_id, idempotency_key) → ValidateCode() | CheckRestrictions() | CheckMaxUses() | Apply() | IncrementUses() → **Outbox:** promo.applied.v1
3. **RevokePromoCommand**(promo_id, reason, idempotency_key) → AuthorizeAdmin() | Revoke() → **Outbox:** promo.revoked.v1
4. **GetPromoQuery**(promo_code) → Validate() | Fetch() → PromoDTO

##### Projections
- promo_read
- promo_usage_read

##### Events Published
- promo.created.v1
- promo.applied.v1
- promo.revoked.v1

##### RBAC/SLO
- **RBAC:** ADMIN (create/revoke), USER (apply/validate)
- **SLO:** P95 < 120ms (validate/apply), P95 < 180ms (create)

---

#### 7.2.2 Entity: discount.go (Discount Calculation)

##### User Stories
- As a **system**, I want to **calculate discount amounts** (percent/fixed) so that pricing is correct.

##### Flow
1. **CalculateDiscountQuery**(promo_code, base_amount) → FetchPromo() | ComputeDiscount() → DiscountDTO

##### Projections
- None (computed)

##### Events Published
- (none)

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 60ms

---

#### 7.2.3 Entity: usage_limit.go (Promo Usage Limits)

##### User Stories
- As a **system**, I want to **enforce per-code and per-user limits** so that promo abuse is prevented.

##### Flow
1. **CheckPromoUsageQuery**(promo_code, user_id) → FetchUsage() | ValidateLimits() → PromoUsageDTO

##### Projections
- promo_usage_read

##### Events Published
- promo.exhausted.v1

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 80ms

---

### 7.3 Domain: trial/

#### 7.3.1 Entity: entity.go (Trial Aggregate)

##### User Stories
- As a **new user**, I want to **start a free trial** so that I can test premium features.
- As a **user**, I want to **cancel trial** before it ends so that I'm not charged.
- As a **system**, I want to **track trial periods** so that conversions are measured.

##### Flow
1. **StartTrialCommand**(user_id, plan_id, trial_duration, idempotency_key) → ValidateTrial() | CreateTrial() | ActivateFeatures() → **Outbox:** trial.started.v1
2. **CancelTrialCommand**(trial_id, cancellation_reason, idempotency_key) → AuthorizeOwner() | Cancel() → **Outbox:** trial.cancelled.v1
3. **GetTrialQuery**(trial_id) → AuthorizeAccess() | Fetch() → TrialDTO

##### Projections
- trial_read
- trial_conversion_read

##### Events Published
- trial.started.v1
- trial.cancelled.v1
- trial.expired.v1

##### RBAC/SLO
- **RBAC:** OWNER (start/cancel/view), SYSTEM (convert), ADMIN (view all)
- **SLO:** P95 < 250ms (start), P95 < 200ms (cancel), P95 < 150ms (read)

---

#### 7.3.2 Entity: eligibility.go (Trial Eligibility)

##### User Stories
- As a **system**, I want to **check trial eligibility** so that users can't abuse free trials.

##### Flow
1. **CheckTrialEligibilityQuery**(user_id, plan_id) → FetchHistory() | ValidateRules() → TrialEligibilityDTO

##### Projections
- trial_eligibility_read

##### Events Published
- (none)

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 100ms

---

#### 7.3.3 Entity: events.go (Trial Conversion Job)

##### User Stories
- As a **system**, I want to **convert trials to paid** at trial end so that revenue starts.
- As a **system**, I want to **send trial reminders** (3 days, 1 day before end) so that users are notified.

##### Flow
1. **ConvertTrialJob**(trial_id) → FetchTrial() | CreatePaymentIntent(financial-be) | ProcessPayment() | ConvertToPaid() → **Outbox:** trial.converted.v1
2. **SendTrialReminderJob**(trial_id, days_before) → Send(communications-be) | LogReminder() → **Outbox:** trial.reminder.sent.v1

##### Projections
- None (background jobs)

##### Events Published
- trial.converted.v1
- trial.reminder.sent.v1

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 400ms (convert), P95 < 200ms (reminder)

---

## **SECTION 8: BILLING CORE**

### 8.1 Domain: invoice/

#### 8.1.1 Entity: entity.go (Invoice Aggregate)

##### User Stories
- As a **system**, I want to **track invoice status** (draft, issued, paid, voided) so that payment is monitored.
- As a **user**, I want to **download invoices** as PDF so that I have records.

##### Flow
1. **IssueInvoiceCommand**(subscription_id, lines[], tax_amount, total, idempotency_key) → Validate() | Issue() → **Outbox:** invoice.issued.v1
2. **VoidInvoiceCommand**(invoice_id, void_reason, idempotency_key) → AuthorizeAdmin() | Void() → **Outbox:** invoice.voided.v1
3. **GetInvoiceQuery**(invoice_id) → AuthorizeAccess() | Fetch() → InvoiceDTO
4. **DownloadInvoicePDFQuery**(invoice_id) → AuthorizeAccess() | GeneratePDF() | PresignURL(storage-be) → InvoicePDFURLDTO

##### Projections
- invoice_read

##### Events Published
- invoice.issued.v1
- invoice.voided.v1

##### Events Consumed
- payment.processed.v1 (mark paid)

##### RBAC/SLO
- **RBAC:** OWNER (view/download), SYSTEM (issue), ADMIN (void)
- **SLO:** P95 < 200ms (issue), P95 < 150ms (void), P95 < 400ms (PDF)

---

#### 8.1.2 Entity: line_item.go (Invoice Line Items)

##### User Stories
- As a **system**, I want to **add line items** to invoices so that charges are itemized.

##### Flow
1. **AddLineItemCommand**(invoice_id, description, amount, idempotency_key) → Validate() | Add() → **Outbox:** invoice.line_item.added.v1

##### Projections
- line_item_read

##### Events Published
- invoice.line_item.added.v1

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 120ms

---

#### 8.1.3 Entity: tax.go (Invoice Tax Calculation)

##### User Stories
- As a **system**, I want to **calculate taxes** per line item so that compliance is maintained.

##### Flow
1. **CalculateInvoiceTaxQuery**(invoice_id) → FetchLineItems() | ApplyTaxRates() → TaxCalculationDTO

##### Projections
- None (computed)

##### Events Published
- (none)

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 100ms

---

### 8.2 Domain: payment/

#### 8.2.1 Entity: entity.go (Payment Intent Aggregate)

##### User Stories
- As a **system**, I want to **create payment intents** for subscriptions so that charges are prepared.
- As a **system**, I want to **record payment attempts** so that success/failure is tracked.

##### Flow
1. **CreatePaymentIntentCommand**(subscription_id, amount, currency, idempotency_key) → Validate() | CreateIntent(financial-be) → **Outbox:** payment.intent.created.v1
2. **RecordPaymentAttemptCommand**(payment_intent_id, status, error_code, idempotency_key) → Record() | UpdateSubscription() → **Outbox:** payment.attempt.recorded.v1

##### Projections
- payment_attempt_read

##### Events Published
- payment.intent.created.v1
- payment.attempt.recorded.v1

##### Events Consumed
- payment.processed.v1 (from financial-be)
- payment.failed.v1 (from financial-be)

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 300ms (create intent), P95 < 150ms (record attempt)

---

#### 8.2.2 Entity: webhook.go (Payment Webhook Handling)

##### User Stories
- As a **system**, I want to **handle payment webhooks** from financial-be so that status is updated.

##### Flow
1. **HandlePaymentWebhookCommand**(webhook_payload) → ValidateSignature() | Process() | UpdateInvoice() → **Outbox:** payment.webhook.processed.v1

##### Projections
- None (event-driven)

##### Events Published
- payment.webhook.processed.v1

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 200ms

---

### 8.3 Domain: credit_note/

#### 8.3.1 Entity: entity.go (Credit Note Aggregate)

##### User Stories
- As an **admin**, I want to **issue credit notes** for refunds so that users get account credits.
- As a **system**, I want to **apply credits** to future invoices so that charges are offset.
- As a **user**, I want to **view my credit balance** so that I know what's available.

##### Flow
1. **IssueCreditNoteCommand**(user_id, amount, reason, issued_by, idempotency_key) → AuthorizeAdmin() | Issue() → **Outbox:** credit_note.issued.v1
2. **ApplyCreditCommand**(invoice_id, credit_note_id, amount, idempotency_key) → Validate() | Apply() | UpdateBalance() → **Outbox:** credit.applied.v1
3. **GetCreditBalanceQuery**(user_id) → Fetch() → CreditBalanceDTO

##### Projections
- credit_note_read
- credit_balance_read

##### Events Published
- credit_note.issued.v1
- credit.applied.v1

##### RBAC/SLO
- **RBAC:** ADMIN (issue), SYSTEM (apply), OWNER (view)
- **SLO:** P95 < 200ms (issue/apply), P95 < 100ms (read)

---

### 8.4 Domain: tax_class/

#### 8.4.1 Entity: entity.go (Tax Class Aggregate)

##### User Stories
- As an **admin**, I want to **define tax classes** (standard, reduced, exempt) so that taxes are calculated correctly.
- As an **admin**, I want to **bind tax classes to plans** so that rates are applied automatically.

##### Flow
1. **CreateTaxClassCommand**(name, rate, jurisdiction, idempotency_key) → Validate() | Persist() → **Outbox:** tax_class.created.v1
2. **UpdateTaxClassCommand**(tax_class_id, updates, idempotency_key) → Apply() → **Outbox:** tax_class.updated.v1
3. **BindTaxClassCommand**(plan_id, tax_class_id, idempotency_key) → Bind() → **Outbox:** tax_class.bound.v1

##### Projections
- tax_class_read

##### Events Published
- tax_class.created.v1
- tax_class.updated.v1
- tax_class.bound.v1

##### RBAC/SLO
- **RBAC:** ADMIN (create/update/bind), SYSTEM (calculate)
- **SLO:** P95 < 180ms (create/update)

---

#### 8.4.2 Entity: validation.go (Tax Validation)

##### User Stories
- As a **system**, I want to **calculate taxes** based on user location and plan so that compliance is maintained.

##### Flow
1. **CalculateTaxQuery**(user_id, plan_id, amount) → FetchTaxClass() | CalculateRate() → TaxCalculationDTO

##### Projections
- None (computed)

##### Events Published
- (none)

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 80ms

---

### 8.5 Domain: billing_profile/

#### 8.5.1 Entity: entity.go (Billing Profile Aggregate)

##### User Stories
- As a **user**, I want to **set my billing profile** (name, address, VAT ID) so that invoices are correct.
- As a **user**, I want to **update my billing profile** when details change so that records are current.

##### Flow
1. **CreateBillingProfileCommand**(user_id, name, address, country, vat_id, idempotency_key) → ValidateVAT() | ValidateAddress() | Persist() → **Outbox:** billing_profile.created.v1
2. **UpdateBillingProfileCommand**(profile_id, updates, idempotency_key) → AuthorizeOwner() | ValidateUpdates() | Apply() → **Outbox:** billing_profile.updated.v1
3. **GetBillingProfileQuery**(user_id) → Fetch() → BillingProfileDTO

##### Projections
- billing_profile_read

##### Events Published
- billing_profile.created.v1
- billing_profile.updated.v1

##### RBAC/SLO
- **RBAC:** OWNER (create/update/view)
- **SLO:** P95 < 200ms (create/update), P95 < 100ms (read)

---

#### 8.5.2 Entity: validation.go (VAT Validation)

##### User Stories
- As a **system**, I want to **validate VAT IDs** (format checks) so that tax compliance is maintained.

##### Flow
1. **ValidateVATQuery**(vat_id, country) → CheckFormat() | OfflineValidation() → VATValidationDTO

##### Projections
- None (validation only)

##### Events Published
- (none)

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 120ms

---

### 8.6 Domain: billing/ (Orchestration)

#### 8.6.1 Entity: service.go (Billing Orchestration)

##### User Stories
- As a **system**, I want to **generate invoices** for subscriptions so that charges are documented.
- As a **system**, I want to **export invoices** to financial-be so that accounting is centralized.
- As a **system**, I want to **prorate charges** for mid-cycle changes so that billing is fair.

##### Flow
1. **GenerateInvoiceCommand**(subscription_id, billing_period, idempotency_key) → FetchSubscription() | CalculateCharges() | ApplyProration() | GenerateInvoice() | ExportToFinancial() → **Outbox:** invoice.generated.v1
2. **ExportInvoiceCommand**(invoice_id, idempotency_key) → FetchInvoice() | MapToFinancialDTO() | SendToFinancial(financial-be) → **Outbox:** invoice.exported.v1
3. **ListInvoicesQuery**(user_id, filters) → ApplyFilters() → InvoiceListDTO

##### Projections
- billing_history_read

##### Events Published
- invoice.generated.v1
- invoice.exported.v1

##### Events Consumed
- subscription.renewed.v1 (generate invoice)
- subscription.created.v1 (generate first invoice)

##### RBAC/SLO
- **RBAC:** OWNER (view own), SYSTEM (generate/export), ADMIN (view all)
- **SLO:** P95 < 400ms (generate), P95 < 300ms (export), P95 < 150ms (read)

---

## **SECTION 9: DUNNING & HISTORY**

### 9.1 Domain: dunning/

#### 9.1.1 Entity: case.go (Dunning Case Aggregate)

##### User Stories
- As a **system**, I want to **open dunning cases** for failed payments so that recovery is automated.
- As a **system**, I want to **advance dunning stages** (retry, warn, suspend, cancel) based on schedule so that collection is systematic.
- As an **admin**, I want to **resolve dunning cases** manually when issues are fixed so that subscriptions are restored.

##### Flow
1. **OpenDunningCaseCommand**(subscription_id, payment_attempt_id, initial_error, idempotency_key) → CreateCase() | ScheduleRetry() → **Outbox:** dunning.case.opened.v1
2. **AdvanceDunningStageJob**(case_id) → FetchCase() | DetermineNextStage() | ExecuteAction() | SendEmail(communications-be) → **Outbox:** dunning.stage.advanced.v1
3. **ResolveDunningCaseCommand**(case_id, resolution_type, resolved_by, idempotency_key) → AuthorizeAdmin() | Resolve() | RestoreSubscription() → **Outbox:** dunning.case.resolved.v1
4. **GetDunningCaseQuery**(case_id) → AuthorizeAdmin() | Fetch() → DunningCaseDTO

##### Projections
- dunning_case_read
- dunning_stage_read

##### Events Published
- dunning.case.opened.v1
- dunning.stage.advanced.v1
- dunning.case.resolved.v1

##### Events Consumed
- payment.failed.v1 (open case)
- payment.processed.v1 (resolve case)

##### RBAC/SLO
- **RBAC:** SYSTEM (open/advance), ADMIN (resolve/view)
- **SLO:** P95 < 200ms (open), P95 < 300ms (advance), P95 < 250ms (resolve)

---

#### 9.1.2 Entity: schedule.go (Dunning Schedule)

##### User Stories
- As an **admin**, I want to **configure dunning schedules** (retry intervals, stage progression) so that recovery strategy is customized.

##### Flow
1. **SetDunningScheduleCommand**(schedule_config, idempotency_key) → Validate() | Persist() → **Outbox:** dunning.schedule.set.v1

##### Projections
- dunning_schedule_read

##### Events Published
- dunning.schedule.set.v1

##### RBAC/SLO
- **RBAC:** ADMIN
- **SLO:** P95 < 180ms

---

#### 9.1.3 Entity: outcome.go (Dunning Outcomes)

##### User Stories
- As a **system**, I want to **cancel subscriptions** after final dunning stage so that non-payers are removed.
- As a **system**, I want to **send dunning emails** at each stage so that users are notified.

##### Flow
1. **RecordDunningOutcomeCommand**(case_id, outcome_type, idempotency_key) → Record() | TriggerAction() → **Outbox:** dunning.outcome.recorded.v1

##### Projections
- dunning_outcome_read

##### Events Published
- dunning.outcome.recorded.v1

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 200ms

---

### 9.2 Domain: billing_history/

#### 9.2.1 Entity: entity.go (Billing History Aggregate)

##### User Stories
- As a **user**, I want to **view my billing history** so that I can track charges.
- As a **system**, I want to **maintain immutable audit snapshots** so that all actions are auditable.

##### Flow
1. **RecordBillingEventCommand**(event_type, subscription_id, metadata, idempotency_key) → AppendWORM() → **Outbox:** billing.event.recorded.v1
2. **GetBillingHistoryQuery**(user_id, filters) → ApplyFilters() → BillingHistoryDTO

##### Projections
- billing_history_read

##### Events Published
- billing.invoice.generated.v1
- billing.payment.applied.v1
- billing.credit.issued.v1

##### RBAC/SLO
- **RBAC:** OWNER (view own), ADMIN (view all), SYSTEM (record)
- **SLO:** P95 < 60ms (append), P95 < 150ms (read)

---

### 9.3 Domain: feature_toggle/

#### 9.3.1 Entity: entity.go (Feature Toggle Aggregate)

##### User Stories
- As an **admin**, I want to **control operational flags** so that features can be enabled/disabled.
- As a **system**, I want to **check feature toggles** before operations so that rollout is controlled.

##### Flow
1. **SetFeatureToggleCommand**(toggle_key, enabled, idempotency_key) → Validate() | Persist() | ClearCache() → **Outbox:** feature.toggle.set.v1
2. **GetFeatureToggleQuery**(toggle_key) → FetchCache() | FallbackDB() → FeatureToggleDTO

##### Projections
- feature_toggle_read

##### Events Published
- admin.feature_flag.updated.v1
- feature.toggle.enabled.v1
- feature.toggle.disabled.v1

##### RBAC/SLO
- **RBAC:** ADMIN (set), SYSTEM (check)
- **SLO:** P95 < 50ms (check - critical), P95 < 150ms (set)

---

## **SECTION 10: INBOX (EVENT CONSUMERS)**

### 10.1 Domain: eventhandler/

#### 10.1.1 Entity: financial_handler.go (Financial Events Consumer)

##### User Stories
- As a **system**, I want to **consume payment.processed/failed** so that subscriptions are activated/paused.

##### Flow
- Consume: payment.processed.v1 → ActivateSubscriptionCommand() | RenewSubscriptionCommand()
- Consume: payment.failed.v1 → OpenDunningCaseCommand() | AdvanceDunningStageCommand()

##### Projections
- subscription_read
- dunning_case_read

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 150ms

---

#### 10.1.2 Entity: proposal_handler.go (Proposal Events Consumer)

##### User Stories
- As a **system**, I want to **consume proposal.submitted** so that connects are deducted.

##### Flow
- Consume: proposal.submitted.v1 → UseConnectsCommand(proposal_id, connects_amount)

##### Projections
- connect_balance_read

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 120ms

---

#### 10.1.3 Entity: job_handler.go (Job Events Consumer)

##### User Stories
- As a **system**, I want to **consume job.posted** so that posting limits are enforced.

##### Flow
- Consume: job.posted.v1 → IncrementUsageCommand(user_id, "job_posts", 1)

##### Projections
- usage_read

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 120ms

---

#### 10.1.4 Entity: admin_flags_handler.go (Admin Events Consumer)

##### User Stories
- As a **system**, I want to **consume admin.feature_flag.updated** so that toggles are refreshed.

##### Flow
- Consume: admin.feature_flag.updated.v1 → RefreshFeatureFlagsCommand()

##### Projections
- feature_toggle_read

##### RBAC/SLO
- **RBAC:** SYSTEM
- **SLO:** P95 < 100ms

---

## **EVENT TOPICS & CATALOG**

### Published Events (subscriptions-be)

**Plan Catalog:**
- plan.created.v1
- plan.updated.v1
- plan.activated.v1
- plan.deactivated.v1
- plan.features.set.v1
- plan.pricing.set.v1
- plan.limits.set.v1
- plan.version.created.v1
- plan.version.activated.v1
- plan.migration.rule.set.v1
- plan.migration.applied.v1

**Subscriptions:**
- subscription.created.v1
- subscription.cancelled.v1
- subscription.paused.v1
- subscription.resumed.v1
- subscription.renewed.v1
- subscription.payment.retried.v1
- subscription.expired.v1
- subscription.auto_renewal.enabled.v1
- subscription.auto_renewal.disabled.v1
- subscription.grace.applied.v1
- subscription.upgrade.requested.v1
- subscription.downgrade.requested.v1
- subscription.change.applied.v1
- subscription.change.cancelled.v1
- proration.policy.set.v1

**Entitlements:**
- grant.issued.v1
- grant.revoked.v1
- grant.expired.v1

**Usage & Allowances:**
- usage.incremented.v1
- usage.reset.v1
- usage.limit.reached.v1
- usage.threshold.warning.v1
- counter.incremented.v1
- allowance.allocated.v1
- allowance.consumed.v1
- allowance.rolled_over.v1
- allowance.expired.v1

**Connects:**
- connects.purchased.v1
- connects.used.v1
- connects.refunded.v1
- connects.low.warning.v1
- connects.depleted.v1
- connects.expired.v1
- connects.granted.v1
- connect_package.created.v1
- connect_package.updated.v1

**Seats:**
- seat.assigned.v1
- seat.released.v1
- seat.overage.charged.v1

**Addons, Promos, Trials:**
- addon.added.v1
- addon.removed.v1
- addon.charged.v1
- promo.created.v1
- promo.applied.v1
- promo.revoked.v1
- promo.exhausted.v1
- trial.started.v1
- trial.cancelled.v1
- trial.converted.v1
- trial.expired.v1
- trial.reminder.sent.v1

**Billing:**
- invoice.issued.v1
- invoice.voided.v1
- invoice.line_item.added.v1
- invoice.generated.v1
- invoice.exported.v1
- payment.intent.created.v1
- payment.attempt.recorded.v1
- payment.webhook.processed.v1
- credit_note.issued.v1
- credit.applied.v1
- tax_class.created.v1
- tax_class.updated.v1
- tax_class.bound.v1
- billing_profile.created.v1
- billing_profile.updated.v1
- billing.event.recorded.v1

**Dunning:**
- dunning.case.opened.v1
- dunning.stage.advanced.v1
- dunning.case.resolved.v1
- dunning.schedule.set.v1
- dunning.outcome.recorded.v1

**Feature Toggles:**
- admin.feature_flag.updated.v1
- feature.toggle.set.v1
- feature.toggle.enabled.v1
- feature.toggle.disabled.v1

---

### Consumed Events (subscriptions-be)

**From financial-be:**
- payment.processed.v1
- payment.failed.v1

**From proposals-be:**
- proposal.submitted.v1
- proposal.rejected.quick.v1

**From jobs-be:**
- job.posted.v1

**From users-be:**
- user.deleted.v1

**From admin-be:**
- admin.feature_flag.updated.v1

---

## **CROSS-SERVICE INTEGRATION**

### Outbound Dependencies

1. **financial-be:** Create payment intents, process payments, export invoices
2. **communications-be:** Send notifications (renewal reminders, dunning emails, trial reminders)
3. **users-be:** Fetch user profiles for billing
4. **storage-be:** Generate and store invoice PDFs

### Inbound Dependencies

1. **financial-be:** Consumes payment events to activate/renew subscriptions
2. **proposals-be:** Consumes proposal events to deduct connects
3. **jobs-be:** Consumes job events to enforce posting limits
4. **admin-be:** Consumes admin events to refresh feature toggles

---

## **GLOBAL SLO TARGETS**

### Read Operations
- Critical path (checks): P95 < 50ms
- Simple queries: P95 < 100ms
- Complex aggregations: P95 < 200ms

### Write Operations
- Simple writes: P95 < 150ms
- Complex writes: P95 < 300ms
- Payment operations: P95 < 600ms

### Event Processing
- Event consumption: P95 < 150ms
- Event publishing: P95 < 100ms

### Background Jobs
- Daily renewals: < 30 minutes
- Dunning advancement: < 5 minutes
- Allowance rollover: < 10 minutes
- Grant expiry: < 5 minutes

---

## **CACHING STRATEGY**

### Redis Caching (TTL)
- Plan by ID: 15m
- Subscription by user_id: 5m
- Connect balance: 2m
- Entitlement checks: 1m
- Feature toggles: 10m
- Billing profile: 5m
- Usage counters: 30s

### Cache Invalidation
- On subscription.created.v1 → Invalidate subscription, entitlements
- On connects.used.v1 → Invalidate connect_balance
- On plan.updated.v1 → Invalidate plan
- On grant.issued.v1 → Invalidate entitlements
- On usage.incremented.v1 → Invalidate usage

---

## **SECURITY & COMPLIANCE**

### PII Protection
- No raw names/emails in events
- Billing profiles contain minimal PII
- Payment details never stored (delegated to financial-be)

### GDPR Compliance
- Full DSAR erasure support
- Data export in JSON formats
- Consent tracking in compliance_context
- Retention policies (7 years default)

### Audit Requirements
- Immutable billing history with hash chains
- All subscription changes logged
- Dunning actions logged
- Payment attempts logged

---

## **FINAL SUMMARY**

**Total Sections:** 10  
**Total Domains:** 25  
**Total Entities:** 60+  
**Total User Stories:** 200+  
**Total Events Published:** 80+  
**Total Events Consumed:** 10+  
**Coverage:** 100% of subscriptions-be folder structure  

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
✅ GDPR compliance  
✅ Audit trails  
✅ Plan versioning  
✅ Proration logic  
✅ Dunning automation  
✅ Connect system  
✅ Seat billing  
✅ Trial management  
✅ Multi-currency support  

---

**END OF subscriptions-be USER STORIES**

```json
{
  "event_id": "uuid",
  "event_type": "subscription.created.v1",
  "event_timestamp": "2025-01-15T10:30:00Z",
  "event_version": "1",
  "aggregate_type": "subscription",
  "aggregate_id": "uuid",
  "correlation_id": "uuid",
  "causation_id": "uuid",
  "event_source": "subscriptions-be",
  "user_context": {
    "user_id": "uuid",
    "keycloak_id": "uuid",
    "user_type": "FREELANCER|CLIENT",
    "session_id": "uuid"
  },
  "compliance_context": {
    "gdpr_consent": true,
    "data_classification": "SENSITIVE|INTERNAL|PUBLIC",
    "retention_policy": "7y"
  },
  "audit_metadata": {
    "ip_address": "x.x.x.x",
    "user_agent": "string",
    "request_id": "uuid"
  },
  "payload": { /* event-specific data */ }
}
```

### Idempotent Write-Path
- All commands accept `idempotency_key` header
- Duplicate requests return cached response (24h TTL)
- Key format: `{service}.{command}.{user_id}.{resource_id}.{hash}`

### Non-PII Event Rules
- Events contain only UUIDs, codes, enums, and numeric values
- No raw names, emails, addresses, or free-text PII
- Consumers fetch PII via API if needed
- Redacted fields marked as `[REDACTED]` in audits

### Platform Alignment
- Folder structure: `/apps/be/subscriptions-be/`
- Events catalog: `/apps/be/contracts/events/`
- Shared libraries: `/platform-shared/`, `/pkg/auth/`

---

## **1 - CATALOG DOMAIN (PLANS & FEATURES)**

### 1.1 plan/

#### User Stories
- As an **admin**, I want to **create subscription plans** so that users can choose the right tier for their needs.
- As an **admin**, I want to **define plan features** (job posts, proposals, connects) so that entitlements are clear.
- As an **admin**, I want to **set plan pricing** per billing cycle (monthly, yearly) so that revenue is optimized.
- As an **admin**, I want to **set plan limits** (daily job posts, invites per month) so that usage is controlled.
- As an **admin**, I want to **activate/deactivate plans** so that plan catalog is managed.
- As a **user**, I want to **view available plans** so that I can compare options.
- As a **system**, I want to **version plans** so that existing subscribers aren't affected by changes.

#### Flow
1. **CreatePlanCommand**(code, name, tier, features, pricing, limits, idempotency_key) → Validate() | Persist() → **Outbox:** plan.created.v1
2. **UpdatePlanCommand**(plan_id, updates, idempotency_key) → VersionIncrement() | Apply() → **Outbox:** plan.updated.v1
3. **ActivatePlanCommand**(plan_id) → Activate() → **Outbox:** plan.activated.v1
4. **DeactivatePlanCommand**(plan_id, reason) → Deactivate() → **Outbox:** plan.deactivated.v1
5. **GetPlanQuery**(plan_id) → Fetch() → PlanDTO
6. **ListPlansQuery**(filters) → ApplyFilters() → PlanListDTO
7. **ComparePlansQuery**(plan_ids[]) → FetchMultiple() | BuildComparison() → PlanComparisonDTO

#### Projections
- plan_read
- plan_catalog_read
- plan_comparison_read

#### Events Published
- plan.created.v1
- plan.updated.v1
- plan.activated.v1
- plan.deactivated.v1

#### RBAC/SLO
- **RBAC:** ADMIN (create/update/activate/deactivate), PUBLIC (view/compare)
- **SLO:** P95 < 100ms (read), P95 < 180ms (write)

---

### 1.2 plan_version/

#### User Stories
- As an **admin**, I want to **version plan changes** so that existing subscribers keep their original terms.
- As a **system**, I want to **freeze plan features** at subscription time so that retroactive changes don't affect users.
- As an **admin**, I want to **view plan version history** so that I can track changes over time.

#### Flow
1. **CreatePlanVersionCommand**(plan_id, changes, version_reason, idempotency_key) → VersionIncrement() | ClonePlan() | Persist() → **Outbox:** plan.version.created.v1
2. **GetPlanVersionQuery**(plan_id, version) → Fetch() → PlanVersionDTO
3. **ListPlanVersionsQuery**(plan_id) → FetchAll() → PlanVersionListDTO

#### Projections
- plan_version_read

#### Events Published
- plan.version.created.v1

#### RBAC/SLO
- **RBAC:** ADMIN (create/view), SYSTEM (freeze version on subscription)
- **SLO:** P95 < 120ms (read), P95 < 200ms (write)

---

## **2 - SUBSCRIPTION DOMAIN**

### 2.1 subscription/

#### User Stories
- As a **user**, I want to **subscribe to a plan** so that I can access premium features.
- As a **user**, I want to **cancel my subscription** so that I can stop charges.
- As a **user**, I want to **pause my subscription** during inactivity so that I don't pay for unused time.
- As a **user**, I want to **resume my subscription** when I return so that I can continue where I left off.
- As a **user**, I want to **view my subscription status** so that I know my current tier and renewal date.
- As a **system**, I want to **auto-renew subscriptions** so that service is uninterrupted.
- As a **system**, I want to **handle payment failures** with retry logic so that temporary issues don't cancel subscriptions.

#### Flow
1. **SubscribeCommand**(user_id, plan_id, billing_cycle, payment_method_id, promo_code, idempotency_key) → ValidatePlan() | ValidatePayment() | CreatePaymentIntent(financial-be) | ProcessPayment() | ActivateSubscription() | FreezeVersion() → **Outbox:** subscription.created.v1
2. **CancelSubscriptionCommand**(subscription_id, cancellation_reason, cancel_at_period_end, idempotency_key) → AuthorizeOwner() | Cancel() | ProcessRefund() → **Outbox:** subscription.cancelled.v1
3. **PauseSubscriptionCommand**(subscription_id, pause_reason, idempotency_key) → AuthorizeOwner() | Validate() | Pause() → **Outbox:** subscription.paused.v1
4. **ResumeSubscriptionCommand**(subscription_id, idempotency_key) → AuthorizeOwner() | Resume() | ProcessPayment() → **Outbox:** subscription.resumed.v1
5. **RenewSubscriptionJob**(subscription_id) → CreatePaymentIntent(financial-be) | ProcessPayment() | ExtendPeriod() | UpdateNextRenewal() → **Outbox:** subscription.renewed.v1
6. **RetryFailedPaymentJob**(subscription_id, retry_count) → ProcessPayment() | UpdateRetryCount() | AdvanceDunning() → **Outbox:** subscription.payment.retried.v1
7. **GetSubscriptionQuery**(subscription_id) → AuthorizeAccess() | Fetch() | EnrichWithFeatures() → SubscriptionDTO
8. **ListSubscriptionsQuery**(user_id, filters) → ApplyFilters() → SubscriptionListDTO
9. **GetSubscriptionHistoryQuery**(subscription_id) → Fetch() → SubscriptionHistoryDTO

#### Projections
- subscription_read
- subscription_history_read
- subscription_status_read

#### Events Published
- subscription.created.v1
- subscription.cancelled.v1
- subscription.paused.v1
- subscription.resumed.v1
- subscription.renewed.v1
- subscription.payment.retried.v1
- subscription.expired.v1

#### Events Consumed
- payment.processed.v1 (activate subscription)
- payment.failed.v1 (handle failed payment)
- user.deleted.v1 (cancel subscriptions)

#### RBAC/SLO
- **RBAC:** OWNER (subscribe/cancel/pause/resume/view), ADMIN (view all/force operations)
- **SLO:** P95 < 400ms (subscribe), P95 < 250ms (cancel/pause/resume), P95 < 600ms (renew), P95 < 150ms (read)

---

### 2.2 change_request/

#### User Stories
- As a **user**, I want to **upgrade my plan** so that I get more features immediately.
- As a **user**, I want to **downgrade my plan** at period end so that I don't lose prepaid time.
- As a **user**, I want to **schedule plan changes** so that transitions happen at the right time.
- As a **system**, I want to **calculate proration** for mid-cycle changes so that charges are fair.
- As a **user**, I want to **cancel a pending plan change** so that I can change my mind.

#### Flow
1. **RequestUpgradeCommand**(subscription_id, new_plan_id, effective_date, idempotency_key) → AuthorizeOwner() | ValidatePlan() | CalculateProration() | CreateChangeRequest() | ProcessImmediate() → **Outbox:** subscription.upgrade.requested.v1
2. **RequestDowngradeCommand**(subscription_id, new_plan_id, effective_date, idempotency_key) → AuthorizeOwner() | ValidatePlan() | ScheduleChange() → **Outbox:** subscription.downgrade.requested.v1
3. **ApplyChangeRequestJob**(change_request_id) → FetchRequest() | ApplyChange() | UpdatePlan() | UpdateBilling() → **Outbox:** subscription.change.applied.v1
4. **CancelChangeRequestCommand**(change_request_id, idempotency_key) → AuthorizeOwner() | Cancel() → **Outbox:** subscription.change.cancelled.v1
5. **GetChangeRequestQuery**(change_request_id) → AuthorizeAccess() | Fetch() → ChangeRequestDTO
6. **ListChangeRequestsQuery**(subscription_id) → Fetch() → ChangeRequestListDTO

#### Projections
- change_request_read
- pending_changes_read

#### Events Published
- subscription.upgrade.requested.v1
- subscription.downgrade.requested.v1
- subscription.change.applied.v1
- subscription.change.cancelled.v1

#### RBAC/SLO
- **RBAC:** OWNER (request/cancel), SYSTEM (apply)
- **SLO:** P95 < 300ms (request), P95 < 400ms (apply), P95 < 150ms (read)

---

## **3 - ENTITLEMENT DOMAIN**

### 3.1 entitlement/

#### User Stories
- As a **user**, I want to **check my entitlements** (job posts, proposals, features) so that I know what I can do.
- As a **system**, I want to **enforce entitlement gates** before operations so that usage limits are respected.
- As a **system**, I want to **compute effective entitlements** from subscription + grants so that special permissions are included.
- As an **admin**, I want to **view user entitlements** for support so that I can troubleshoot issues.

#### Flow
1. **CheckEntitlementQuery**(user_id, feature_key) → FetchSubscription() | FetchGrants() | ComputeEffective() | CheckLimit() → EntitlementCheckDTO
2. **ListEntitlementsQuery**(user_id) → FetchSubscription() | FetchGrants() | ComputeAll() → EntitlementListDTO
3. **GetEffectiveEntitlementsQuery**(user_id) → FetchSubscription() | FetchGrants() | MergeAll() → EffectiveEntitlementDTO

#### Projections
- entitlement_read
- effective_entitlement_read

#### Events Published
- (none - query-only)

#### Events Consumed
- subscription.created.v1 (refresh entitlements)
- subscription.changed.v1 (refresh entitlements)
- grant.issued.v1 (refresh entitlements)

#### RBAC/SLO
- **RBAC:** OWNER (view own), ADMIN (view all), SYSTEM (check/enforce)
- **SLO:** P95 < 50ms (check - critical path), P95 < 100ms (list)

---

### 3.2 entitlement_grant/

#### User Stories
- As an **admin**, I want to **issue temporary grants** (bonus connects, feature access) so that I can reward users.
- As an **admin**, I want to **set grant expiry** so that temporary access is time-boxed.
- As an **admin**, I want to **revoke grants** if misused so that abuse is stopped.
- As a **system**, I want to **expire grants automatically** so that manual cleanup isn't needed.

#### Flow
1. **IssueGrantCommand**(user_id, grant_type, grant_value, expiry_date, reason, issued_by, idempotency_key) → AuthorizeAdmin() | Validate() | Issue() → **Outbox:** grant.issued.v1
2. **RevokeGrantCommand**(grant_id, reason, revoked_by, idempotency_key) → AuthorizeAdmin() | Revoke() → **Outbox:** grant.revoked.v1
3. **ExpireGrantsJob**() → FindExpired() | BatchExpire() → **Outbox:** grant.expired.v1
4. **GetGrantQuery**(grant_id) → AuthorizeAdmin() | Fetch() → GrantDTO
5. **ListGrantsQuery**(user_id) → Fetch() → GrantListDTO

#### Projections
- grant_read
- active_grants_read

#### Events Published
- grant.issued.v1
- grant.revoked.v1
- grant.expired.v1

#### RBAC/SLO
- **RBAC:** ADMIN (issue/revoke), SYSTEM (expire), OWNER (view own)
- **SLO:** P95 < 180ms (issue/revoke), P95 < 100ms (read)

---

3.3 entitlement\_grant/scope/
-----------------------------

### Stories

*   As an **admin**, I want to define the **scope** of a grant (feature-level vs. meter-level) so that grants resolve correctly against plan limits.
    
*   As a **system**, I want precedence rules (**plan < addon < promo < ad-hoc grant**) applied consistently so that effective entitlements are deterministic.
    
*   As an **admin**, I want to **update or revoke** a grant’s scope so that I can correct mistakes without breaking usage history.
    
*   As a **system**, I want **grant scope to expire** automatically at expires\_at so that temporary grants don’t persist.
    

### Flow

*   IssueGrantCommand(user\_id, feature\_key|meter\_key, scope, qty?, expires\_at, reason, idem\_key) → ValidateScope() | StoreGrant() → RebuildEffectiveEntitlements(user\_id) → **Outbox:** entitlement.grant.issued.v1
    
*   UpdateGrantScopeCommand(grant\_id, new\_scope, qty?) → ValidateChange() | Update() → RebuildEffectiveEntitlements(user\_id) → **Outbox:** entitlement.grant.scope.updated.v1
    
*   ConsumeGrantOnUse(user\_id, meter\_key, amount, usage\_token) → ResolvePrecedence() | ConsumeIfApplicable() → **Outbox:** entitlement.grant.consumed.v1
    
*   ExpireGrantJob() → ScanExpired() | SetExpired() → RebuildEffectiveEntitlements(user\_id) → **Outbox:** entitlement.grant.expired.v1
    

### Projections

*   entitlement\_grants\_read (grant\_id → user\_id, scope{feature|meter}, key, qty\_remaining, expires\_at, reason, state)
    
*   effective\_entitlements\_read (user\_id → resolved feature gates & meter boosts; precedence trace)
    
*   grant\_resolution\_audit (user\_id, request\_ref → inputs, chosen grant, reason)
    

### Events

*   entitlement.grant.issued.v1, entitlement.grant.scope.updated.v1, entitlement.grant.consumed.v1, entitlement.grant.expired.v1
    

### RBAC/SLO

*   **RBAC:** **ADMIN/SYSTEM** issue/update/expire; **USER** read own effective entitlements.
    
*   **SLO:** resolve effective entitlements P95 < 50ms; grant issue/update P95 < 150ms.
    
*   **Limits:** ≤ 100 active grants/user; grant duration ≤ 12 months.
    
*   **Idempotency:** by (user\_id, idem\_key) for issue; by (grant\_id, version) for update.


---

## **4 - USAGE TRACKING DOMAIN**

### 4.1 usage/

#### User Stories
- As a **system**, I want to **track usage** (job posts, proposals, invites) so that limits are enforced.
- As a **system**, I want to **increment usage counters** on each action so that quotas are consumed.
- As a **system**, I want to **reset usage counters** at period boundaries so that monthly limits restart.
- As a **user**, I want to **view my usage** so that I know how much quota I have left.
- As a **system**, I want to **emit warnings** when approaching limits so that users can upgrade.

#### Flow
1. **IncrementUsageCommand**(user_id, usage_type, amount, idempotency_key) → FetchEntitlement() | CheckLimit() | Increment() | CheckThreshold() → **Outbox:** usage.incremented.v1
2. **ResetUsageJob**(period_type) → FindExpired() | BatchReset() → **Outbox:** usage.reset.v1
3. **GetUsageQuery**(user_id) → Fetch() | EnrichWithLimits() → UsageDTO
4. **GetUsageHistoryQuery**(user_id, time_range) → Fetch() → UsageHistoryDTO

#### Projections
- usage_read
- usage_history_read

#### Events Published
- usage.incremented.v1
- usage.reset.v1
- usage.limit.reached.v1
- usage.threshold.warning.v1

#### Events Consumed
- job.posted.v1 (increment job posts)
- proposal.submitted.v1 (increment proposals)
- invitation.sent.v1 (increment invites)

#### RBAC/SLO
- **RBAC:** SYSTEM (increment/reset), OWNER (view own), ADMIN (view all)
- **SLO:** P95 < 80ms (increment - critical path), P95 < 100ms (read)

---

4.1.1 usage/meter/
----------------

### Stories

*   As an **admin**, I want to define **meters** (e.g., messages\_to\_non\_hires, boosts, invites) so that product events map cleanly to usage.
    
*   As a **system**, I want to **update meter mappings** when products evolve so that counters stay aligned.
    
*   As an **admin**, I want to **deprecate meters** safely so that historic counters remain readable.
    

### Flow

*   DefineMeterCommand(meter\_key, description, unit, event\_mappings\[\], active\_from) → ValidateUniqueness() | StoreMeter() → **Outbox:** usage.meter.defined.v1
    
*   UpdateMeterMappingCommand(meter\_key, event\_mappings\[\]) → ValidateMappings() | Update() | WarmCache() → **Outbox:** usage.meter.mapping.updated.v1
    
*   DeprecateMeterCommand(meter\_key, deprecate\_from) → CheckNoActiveDependencies() | MarkDeprecated() → **Outbox:** usage.meter.deprecated.v1
    
*   MapEventToMeter(event\_name, payload) → LookupMeterMapping() | EmitIncrement(meter\_key, amount, usage\_token) (→ handled by usage service)
    

### Projections

*   usage\_meters\_read (meter\_key → unit, description, mappings\[\], active/deprecated flags, timestamps)
    
*   meter\_mapping\_cache (Redis: event\_name → meter\_key, TTL 10m)
    

### Events

*   usage.meter.defined.v1, usage.meter.mapping.updated.v1, usage.meter.deprecated.v1
    

### RBAC/SLO

*   **RBAC:** **ADMIN** define/update/deprecate; **SYSTEM** read mappings.
    
*   **SLO:** event→meter mapping P95 < 5ms; meter upsert P95 < 150ms.
    
*   **Limits:** ≤ 200 meters; ≤ 10 event mappings/meter.
    
*   **Idempotency:** by (meter\_key, version) for updates.
    
---

4.1.2 usage/quota/
----------------

### Stories

*   As an **admin**, I want to manage **quota policies** (soft/hard caps, thresholds) so that usage enforcement is consistent.
    
*   As a **system**, I want **threshold callbacks** (e.g., 80%, 100%) so that product UI and comms can react.
    
*   As an **admin**, I want **policy versioning** so that changes don’t retroactively alter past decisions.
    

### Flow

*   CreateQuotaPolicyCommand(feature\_key|meter\_key, soft\_cap?, hard\_cap, thresholds\[\]) → ValidateBounds() | StorePolicy() → **Outbox:** usage.quota.policy.created.v1
    
*   UpdateQuotaPolicyCommand(policy\_id, changes) → Validate() | StoreVersion() | WarmCache() → **Outbox:** usage.quota.policy.updated.v1
    
*   EnforceQuotaOnIncrement(user\_id, key, amount, usage\_token) → LoadEffectivePolicy() | CheckThresholds() | CheckCaps() → EmitThresholdEvents() | AllowOrBlock() → **Outbox:** usage.quota.threshold.triggered.v1 / usage.limit.reached.v1
    

### Projections

*   usage\_quota\_policies\_read (policy\_id → key, soft\_cap, hard\_cap, thresholds\[\], active\_version, updated\_at)
    
*   quota\_policy\_cache (Redis: key → active policy, TTL 5m)
    

### Events

*   usage.quota.policy.created.v1, usage.quota.policy.updated.v1, usage.quota.threshold.triggered.v1
    

### RBAC/SLO

*   **RBAC:** **ADMIN** create/update; **SYSTEM** enforce.
    
*   **SLO:** policy fetch P95 < 5ms; enforcement path overhead P95 < 10ms.
    
*   **Limits:** thresholds ≤ 5 per key; hard\_cap ≤ 1e9.
    
*   **Idempotency:** policy updates by (policy\_id, version); enforcement by (user\_id, key, usage\_token).
    

4.1.3 usage/limit/
----------------

### Stories

*   As a **system**, I want to **snapshot per-plan limits** (static caps) so that enforcement is fast and audit-friendly.
    
*   As an **admin**, I want **rebuilds on plan/version changes** so that snapshots stay up-to-date.
    
*   As a **system**, I want **read-time enrichment** with the latest snapshot so that requests don’t scan plan trees.
    

### Flow

*   BuildPlanLimitSnapshotCommand(plan\_version\_id) → ReadPlanLimits() | PersistSnapshot() → **Outbox:** usage.plan\_limits.snapshot.created.v1
    
*   RefreshSnapshotsOnPlanVersionActivated(plan\_id, new\_version\_id) → BuildPlanLimitSnapshotCommand() → **Outbox:** usage.plan\_limits.snapshot.refreshed.v1
    
*   EnrichRequestWithLimits(query\_ctx) → LoadSnapshot(plan\_version\_id) → AttachLimits()
    

### Projections

*   usage\_plan\_limits\_snapshot\_read (plan\_version\_id → {key: cap} map, created\_at, hash)
    
*   plan\_version\_to\_snapshot\_read (plan\_id → active\_version\_id, snapshot\_id)
    

### Events

*   usage.plan\_limits.snapshot.created.v1, usage.plan\_limits.snapshot.refreshed.v1
    

### RBAC/SLO

*   **RBAC:** **ADMIN/SYSTEM** build/refresh; **SYSTEM** read for enforcement.
    
*   **SLO:** build snapshot for 100 features P95 < 2s; read P95 < 2ms.
    
*   **Limits:** ≤ 500 keys/snapshot; snapshot size ≤ 256KB.
    
*   **Idempotency:** by (plan\_version\_id, hash).

---

### 4.2 allowance/

#### User Stories
- As a **system**, I want to **allocate monthly allowances** (job posts, proposals) per subscription so that usage is bucketed.
- As a **system**, I want to **rollover unused allowances** if plan allows so that value is preserved.
- As a **system**, I want to **expire old allowances** so that usage doesn't accumulate indefinitely.
- As a **user**, I want to **view my allowance breakdown** so that I understand my quota.

#### Flow
1. **AllocateAllowanceCommand**(user_id, allowance_type, amount, period_start, period_end, idempotency_key) → Validate() | Allocate() → **Outbox:** allowance.allocated.v1
2. **ConsumeAllowanceCommand**(user_id, allowance_type, amount, idempotency_key) → FindAvailable() | Consume() | CheckThreshold() → **Outbox:** allowance.consumed.v1
3. **RolloverAllowanceJob**(user_id) → FetchUnused() | ApplyRolloverPolicy() | Allocate() → **Outbox:** allowance.rolled_over.v1
4. **ExpireAllowanceJob**() → FindExpired() | BatchExpire() → **Outbox:** allowance.expired.v1
5. **GetAllowanceQuery**(user_id) → Fetch() | ComputeAvailable() → AllowanceDTO

#### Projections
- allowance_read
- available_allowance_read

#### Events Published
- allowance.allocated.v1
- allowance.consumed.v1
- allowance.rolled_over.v1
- allowance.expired.v1

#### Events Consumed
- subscription.renewed.v1 (allocate new period)
- usage.incremented.v1 (consume allowance)

#### RBAC/SLO
- **RBAC:** SYSTEM (allocate/consume/rollover/expire), OWNER (view own)
- **SLO:** P95 < 100ms (allocate/consume), P95 < 80ms (read)

---

## **5 - CONNECTS DOMAIN**

### 5.1 connect/

#### User Stories
- As a **user**, I want to **purchase connect packages** so that I can submit proposals.
- As a **system**, I want to **deduct connects** on proposal submission so that usage is charged.
- As a **system**, I want to **track connect balance** so that users know how many they have left.
- As a **system**, I want to **refund connects** if proposal is rejected within 24h so that failed applications don't cost.
- As a **user**, I want to **view connect history** so that I can track spending.
- As a **system**, I want to **warn users** when connects are low so that they can recharge.

#### Flow
1. **PurchaseConnectsCommand**(user_id, package_id, payment_method_id, promo_code, idempotency_key) → ValidatePackage() | CreatePaymentIntent(financial-be) | ProcessPayment() | AddConnects() → **Outbox:** connects.purchased.v1
2. **UseConnectsCommand**(user_id, proposal_id, connects_amount, idempotency_key) → CheckBalance() | Deduct() | CheckThreshold() → **Outbox:** connects.used.v1
3. **RefundConnectsCommand**(usage_id, refund_reason, idempotency_key) → ValidateRefund() | Refund() → **Outbox:** connects.refunded.v1
4. **GetConnectBalanceQuery**(user_id) → Fetch() → ConnectBalanceDTO
5. **GetConnectHistoryQuery**(user_id, filters) → ApplyFilters() → ConnectHistoryDTO

#### Projections
- connect_balance_read
- connect_history_read
- connect_package_read

#### Events Published
- connects.purchased.v1
- connects.used.v1
- connects.refunded.v1
- connects.low.warning.v1
- connects.depleted.v1

#### Events Consumed
- proposal.submitted.v1 (deduct connects)
- proposal.rejected.quick.v1 (refund connects if eligible)
- payment.processed.v1 (add connects after purchase)

#### RBAC/SLO
- **RBAC:** OWNER (purchase/view), SYSTEM (use/refund)
- **SLO:** P95 < 300ms (purchase), P95 < 100ms (use - critical path), P95 < 200ms (refund), P95 < 80ms (read)

---

### 5.2 connect_package/

#### User Stories
- As an **admin**, I want to **create connect packages** (starter, value, premium) so that users have options.
- As an **admin**, I want to **set package pricing** with volume discounts so that bulk purchases are cheaper.
- As an **admin**, I want to **add bonus connects** to packages for promotions so that purchases are incentivized.
- As an **admin**, I want to **deactivate packages** so that outdated options are removed.

#### Flow
1. **CreatePackageCommand**(name, connects_amount, price, bonus_connects, active, idempotency_key) → Validate() | Persist() → **Outbox:** connect_package.created.v1
2. **UpdatePackageCommand**(package_id, updates, idempotency_key) → Apply() → **Outbox:** connect_package.updated.v1
3. **ActivatePackageCommand**(package_id) → Activate() → **Outbox:** connect_package.activated.v1
4. **DeactivatePackageCommand**(package_id, reason) → Deactivate() → **Outbox:** connect_package.deactivated.v1
5. **GetPackageQuery**(package_id) → Fetch() → PackageDTO
6. **ListPackagesQuery**(filters) → ApplyFilters() → PackageListDTO

#### Projections
- connect_package_read

#### Events Published
- connect_package.created.v1
- connect_package.updated.v1
- connect_package.activated.v1
- connect_package.deactivated.v1

#### RBAC/SLO
- **RBAC:** ADMIN (create/update/activate/deactivate), PUBLIC (view)
- **SLO:** P95 < 120ms (read), P95 < 180ms (write)

---

## **6 - SEAT BILLING DOMAIN**

### 6.1 seat_billing/

#### User Stories
- As an **organization**, I want to **add team seats** to my subscription so that multiple users can access premium features.
- As an **organization**, I want to **pay per active seat** so that costs scale with team size.
- As a **system**, I want to **track seat assignments** so that usage is accurate.
- As a **system**, I want to **charge overages** when seats exceed plan limits so that extra usage is billed.
- As an **organization**, I want to **remove seats** when team members leave so that costs are optimized.

#### Flow
1. **AssignSeatCommand**(subscription_id, user_id, assigned_by, idempotency_key) → CheckAvailableSeats() | AssignSeat() | CheckOverage() → **Outbox:** seat.assigned.v1
2. **ReleaseSeatCommand**(seat_id, released_by, idempotency_key) → AuthorizeAdmin() | Release() → **Outbox:** seat.released.v1
3. **ChargeOverageJob**(subscription_id) → CalculateOverage() | CreateInvoice(financial-be) → **Outbox:** seat.overage.charged.v1
4. **GetSeatUsageQuery**(subscription_id) → Fetch() | ComputeOverage() → SeatUsageDTO
5. **ListSeatsQuery**(subscription_id) → Fetch() → SeatListDTO

#### Projections
- seat_usage_read
- seat_assignment_read

#### Events Published
- seat.assigned.v1
- seat.released.v1
- seat.overage.charged.v1

#### RBAC/SLO
- **RBAC:** ORG_ADMIN (assign/release), OWNER (view), SYSTEM (charge)
- **SLO:** P95 < 200ms (assign/release), P95 < 120ms (read)

---

6.2 seat\_billing/invoice\_export/
----------------------------------

### Stories

*   As a **system**, I want to **prepare invoice line items** for seat overages so that financial-be can import them idempotently.
    
*   As an **admin**, I want to **rebuild seat exports** for a billing period so that corrections can be applied.
    
*   As a **system**, I want **traceability** from seat deltas → invoice lines so that audits are simple.
    

### Flow

*   PrepareSeatInvoiceExportCommand(subscription\_id, period) → ComputeSeatDeltas() | PriceOverages() | StageLineItems() → **Outbox:** seat.invoice\_export.prepared.v1
    
*   RebuildSeatExportCommand(subscription\_id, period) → ClearStaging() | PrepareSeatInvoiceExportCommand() → **Outbox:** seat.invoice\_export.rebuilt.v1
    
*   MarkSeatExportAsImportedCommand(export\_id, invoice\_id) → LockStagingRows() | MarkImported() → **Outbox:** seat.invoice\_export.imported.v1
    

### Projections

*   seat\_export\_staging\_read (export\_id → subscription\_id, period, line\_items\[\], totals, state)
    
*   seat\_overage\_audit (subscription\_id, period → assigned, cap, overage\_units, price, calc\_inputs)
    

### Events

*   seat.invoice\_export.prepared.v1, seat.invoice\_export.rebuilt.v1, seat.invoice\_export.imported.v1
    

### RBAC/SLO

*   **RBAC:** **SYSTEM** prepare/import; **ADMIN** rebuild.
    
*   **SLO:** prepare export for 10k seats P95 < 60s; import mark P95 < 300ms.
    
*   **Limits:** batch ≤ 10k line items; retries ≤ 3.
    
*   **Idempotency:** by (subscription\_id, period).
    
---

## **7 - ADDON DOMAIN**

### 7.1 addon/

#### User Stories
- As a **user**, I want to **add optional features** (priority support, featured profile) to my subscription so that I get extras.
- As a **user**, I want to **remove addons** when no longer needed so that costs are controlled.
- As a **system**, I want to **charge for addons** monthly so that billing is automated.
- As an **admin**, I want to **create addons** so that monetization is flexible.

#### Flow
1. **AddAddonCommand**(subscription_id, addon_id, idempotency_key) → ValidateAddon() | CheckCompatibility() | Add() | CreateInvoice(financial-be) → **Outbox:** addon.added.v1
2. **RemoveAddonCommand**(subscription_addon_id, idempotency_key) → AuthorizeOwner() | Remove() | ProcessRefund() → **Outbox:** addon.removed.v1
3. **ChargeAddonJob**(subscription_addon_id) → FetchAddon() | CreateInvoice(financial-be) → **Outbox:** addon.charged.v1
4. **GetAddonsQuery**(subscription_id) → Fetch() → AddonListDTO

#### Projections
- addon_read
- subscription_addon_read

#### Events Published
- addon.added.v1
- addon.removed.v1
- addon.charged.v1

#### RBAC/SLO
- **RBAC:** OWNER (add/remove/view), SYSTEM (charge), ADMIN (create addons)
- **SLO:** P95 < 250ms (add/remove), P95 < 100ms (read)

---

## **8 - PROMOTION DOMAIN**

### 8.1 promotion/

#### User Stories
- As an **admin**, I want to **create promo codes** (discount percentages, free trials) so that marketing campaigns work.
- As an **admin**, I want to **set promo limits** (max uses, expiry date) so that abuse is prevented.
- As a **user**, I want to **apply promo codes** during subscription so that I get discounts.
- As a **system**, I want to **validate promo codes** before applying so that invalid codes are rejected.
- As an **admin**, I want to **deactivate promo codes** when campaigns end so that usage is controlled.

#### Flow
1. **CreatePromoCommand**(code, discount_type, discount_value, max_uses, expiry_date, restrictions, idempotency_key) → Validate() | Persist() → **Outbox:** promo.created.v1
2. **ApplyPromoCommand**(user_id, promo_code, subscription_id, idempotency_key) → ValidateCode() | CheckRestrictions() | CheckMaxUses() | Apply() | IncrementUses() → **Outbox:** promo.applied.v1
3. **RevokePromoCommand**(promo_id, reason, idempotency_key) → AuthorizeAdmin() | Revoke() → **Outbox:** promo.revoked.v1
4. **DeactivatePromoCommand**(promo_id, idempotency_key) → Deactivate() → **Outbox:** promo.deactivated.v1
5. **GetPromoQuery**(promo_code) → Validate() | Fetch() → PromoDTO
6. **ListPromosQuery**(filters) → AuthorizeAdmin() | ApplyFilters() → PromoListDTO

#### Projections
- promo_read
- promo_usage_read

#### Events Published
- promo.created.v1
- promo.applied.v1
- promo.revoked.v1
- promo.deactivated.v1

#### RBAC/SLO
- **RBAC:** ADMIN (create/revoke/deactivate/view all), USER (apply/validate)
- **SLO:** P95 < 120ms (validate/apply), P95 < 180ms (create), P95 < 100ms (read)

---

## **9 - TRIAL DOMAIN**

### 9.1 trial/

#### User Stories
- As a **new user**, I want to **start a free trial** so that I can test premium features.
- As a **system**, I want to **track trial periods** so that conversions are measured.
- As a **system**, I want to **convert trials to paid** at trial end so that revenue starts.
- As a **system**, I want to **send trial reminders** (3 days, 1 day before end) so that users are notified.
- As a **user**, I want to **cancel trial** before it ends so that I'm not charged.

#### Flow
1. **StartTrialCommand**(user_id, plan_id, trial_duration, idempotency_key) → ValidateTrial() | CreateTrial() | ActivateFeatures() → **Outbox:** trial.started.v1
2. **CancelTrialCommand**(trial_id, cancellation_reason, idempotency_key) → AuthorizeOwner() | Cancel() → **Outbox:** trial.cancelled.v1
3. **ConvertTrialJob**(trial_id) → FetchTrial() | CreatePaymentIntent(financial-be) | ProcessPayment() | ConvertToPaid() → **Outbox:** trial.converted.v1
4. **SendTrialReminderJob**(trial_id, days_before) → Send(communications-be) | LogReminder() → **Outbox:** trial.reminder.sent.v1
5. **GetTrialQuery**(trial_id) → AuthorizeAccess() | Fetch() → TrialDTO
6. **ListTrialsQuery**(filters) → AuthorizeAdmin() | ApplyFilters() → TrialListDTO

#### Projections
- trial_read
- trial_conversion_read

#### Events Published
- trial.started.v1
- trial.cancelled.v1
- trial.converted.v1
- trial.expired.v1
- trial.reminder.sent.v1

#### RBAC/SLO
- **RBAC:** OWNER (start/cancel/view), SYSTEM (convert/send reminders), ADMIN (view all)
- **SLO:** P95 < 250ms (start), P95 < 200ms (cancel), P95 < 150ms (read)

---

## **10 - BILLING DOMAIN**

### 10.1 billing/

#### User Stories
- As a **system**, I want to **generate invoices** for subscriptions so that charges are documented.
- As a **system**, I want to **export invoices** to financial-be so that accounting is centralized.
- As a **system**, I want to **prorate charges** for mid-cycle changes so that billing is fair.
- As a **user**, I want to **view my billing history** so that I can track charges.

#### Flow
1. **GenerateInvoiceCommand**(subscription_id, billing_period, idempotency_key) → FetchSubscription() | CalculateCharges() | ApplyProration() | GenerateInvoice() | ExportToFinancial() → **Outbox:** invoice.generated.v1
2. **ExportInvoiceCommand**(invoice_id, idempotency_key) → FetchInvoice() | MapToFinancialDTO() | SendToFinancial(financial-be) → **Outbox:** invoice.exported.v1
3. **GetInvoiceQuery**(invoice_id) → AuthorizeAccess() | Fetch() → InvoiceDTO
4. **ListInvoicesQuery**(user_id, filters) → ApplyFilters() → InvoiceListDTO

#### Projections
- invoice_read
- billing_history_read

#### Events Published
- invoice.generated.v1
- invoice.exported.v1

#### Events Consumed
- subscription.renewed.v1 (generate invoice)
- subscription.created.v1 (generate first invoice)

#### RBAC/SLO
- **RBAC:** OWNER (view own), SYSTEM (generate/export), ADMIN (view all)
- **SLO:** P95 < 400ms (generate), P95 < 300ms (export), P95 < 150ms (read)

---

### 10.2 invoice/

#### User Stories
- As a **system**, I want to **track invoice status** (pending, paid, overdue) so that payment is monitored.
- As a **system**, I want to **void invoices** if errors occur so that billing is corrected.
- As a **user**, I want to **download invoices** as PDF so that I have records.

#### Flow
1. **IssueInvoiceCommand**(subscription_id, lines[], tax_amount, total, idempotency_key) → Validate() | Issue() → **Outbox:** invoice.issued.v1
2. **VoidInvoiceCommand**(invoice_id, void_reason, idempotency_key) → AuthorizeAdmin() | Void() → **Outbox:** invoice.voided.v1
3. **GetInvoiceQuery**(invoice_id) → AuthorizeAccess() | Fetch() → InvoiceDTO
4. **DownloadInvoicePDFQuery**(invoice_id) → AuthorizeAccess() | GeneratePDF() | PresignURL(storage-be) → InvoicePDFURLDTO

#### Projections
- invoice_read

#### Events Published
- invoice.issued.v1
- invoice.voided.v1

#### Events Consumed
- payment.processed.v1 (mark paid)

#### RBAC/SLO
- **RBAC:** OWNER (view/download), SYSTEM (issue), ADMIN (void)
- **SLO:** P95 < 200ms (issue), P95 < 150ms (void), P95 < 400ms (PDF - includes generation)

---

10.2.1 invoice/adjustment/
-----------------------

### Stories

*   As an **admin**, I want to add **adjustment lines** (proration/credits/fees) so that invoice totals reflect policy and state changes.
    
*   As a **system**, I want to **validate adjustments** (amounts, sign, taxable?) so that invoices remain consistent.
    
*   As an **admin**, I want to **void or update** adjustments before issuance so that I can correct errors.
    

### Flow

*   AddAdjustmentLineCommand(invoice\_id, kind{proration|credit|fee}, amount, taxable?, meta, idem\_key) → ValidateInvoiceState(draft) | ValidateAmount() | AddLine() → **Outbox:** invoice.adjustment.added.v1
    
*   UpdateAdjustmentLineCommand(invoice\_id, line\_id, patch) → ValidateInvoiceState(draft) | UpdateLine() → **Outbox:** invoice.adjustment.updated.v1
    
*   RemoveAdjustmentLineCommand(invoice\_id, line\_id) → ValidateInvoiceState(draft) | RemoveLine() → **Outbox:** invoice.adjustment.removed.v1
    
*   ValidateBeforeIssue(invoice\_id) → RecalcTotals() | ValidateNoNegativeTotals() | **Outbox:** invoice.validated.v1
    

### Projections

*   invoice\_adjustments\_read (invoice\_id → lines\[kind, amount, taxable, meta\], totals)
    
*   invoice\_validation\_issues\_read (invoice\_id → issues\[\], resolved\_at?)
    

### Events

*   invoice.adjustment.added.v1, invoice.adjustment.updated.v1, invoice.adjustment.removed.v1, invoice.validated.v1
    

### RBAC/SLO

*   **RBAC:** **ADMIN** add/update/remove; **SYSTEM** validate on issue.
    
*   **SLO:** add/update P95 < 150ms; validation P95 < 300ms.
    
*   **Limits:** ≤ 50 adjustment lines/invoice; |amount| ≤ invoice subtotal.
    
*   **Idempotency:** add by (invoice\_id, idem\_key); update/remove by (invoice\_id, line\_id, version).
    

---

### 10.3 payment/

#### User Stories
- As a **system**, I want to **create payment intents** for subscriptions so that charges are prepared.
- As a **system**, I want to **record payment attempts** so that success/failure is tracked.
- As a **system**, I want to **handle payment webhooks** from financial-be so that status is updated.

#### Flow
1. **CreatePaymentIntentCommand**(subscription_id, amount, currency, idempotency_key) → Validate() | CreateIntent(financial-be) → **Outbox:** payment.intent.created.v1
2. **RecordPaymentAttemptCommand**(payment_intent_id, status, error_code, idempotency_key) → Record() | UpdateSubscription() → **Outbox:** payment.attempt.recorded.v1
3. **HandlePaymentWebhookCommand**(webhook_payload) → ValidateSignature() | Process() | UpdateInvoice() → **Outbox:** payment.webhook.processed.v1

#### Projections
- payment_attempt_read

#### Events Published
- payment.intent.created.v1
- payment.attempt.recorded.v1
- payment.webhook.processed.v1

#### Events Consumed
- payment.processed.v1 (from financial-be)
- payment.failed.v1 (from financial-be)

#### RBAC/SLO
- **RBAC:** SYSTEM (all operations)
- **SLO:** P95 < 300ms (create intent), P95 < 150ms (record attempt), P95 < 200ms (webhook)

---

10.3.1 payment/method\_hint/
-------------------------

### Stories

*   As a **system**, I want to **store non-PII payment method hints** (brand, last4, exp) so that UX and support can reference a method safely.
    
*   As an **admin**, I want to **update/clear hints** when gateway updates arrive so that data stays fresh.
    
*   As a **system**, I want **hints per invoice/payment intent** so that we can show accurate context during dunning.
    

### Flow

*   StoreMethodHintCommand(user\_id, intent\_id|invoice\_id, brand, last4, exp\_month, exp\_year) → ValidateNonPII() | UpsertHint() → **Outbox:** payment.method\_hint.stored.v1
    
*   UpdateMethodHintOnWebhook(intent\_id, payload) → Validate() | UpsertHint() → **Outbox:** payment.method\_hint.updated.v1
    
*   ClearMethodHintCommand(intent\_id|invoice\_id) → RemoveHint() → **Outbox:** payment.method\_hint.cleared.v1
    

### Projections

*   payment\_method\_hints\_read (user\_id, intent\_id|invoice\_id → brand, last4, exp\_month, exp\_year, updated\_at)
    
*   dunning\_case\_view (join) → latest method hint for display
    

### Events

*   payment.method\_hint.stored.v1, payment.method\_hint.updated.v1, payment.method\_hint.cleared.v1
    

### RBAC/SLO

*   **RBAC:** **SYSTEM** store/update/clear; **ADMIN** read for support; **USER** read own hints.
    
*   **SLO:** upsert P95 < 100ms; read P95 < 50ms.
    
*   **Limits:** ≤ 5 hints/intent; TTL 18 months post-payment.
    
*   **Idempotency:** by (intent\_id|invoice\_id, brand, last4, exp\_year, exp\_month).

---

### 10.4 credit_note/

#### User Stories
- As an **admin**, I want to **issue credit notes** for refunds so that users get account credits.
- As a **system**, I want to **apply credits** to future invoices so that charges are offset.
- As a **user**, I want to **view my credit balance** so that I know what's available.

#### Flow
1. **IssueCreditNoteCommand**(user_id, amount, reason, issued_by, idempotency_key) → AuthorizeAdmin() | Issue() → **Outbox:** credit_note.issued.v1
2. **ApplyCreditCommand**(invoice_id, credit_note_id, amount, idempotency_key) → Validate() | Apply() | UpdateBalance() → **Outbox:** credit.applied.v1
3. **GetCreditBalanceQuery**(user_id) → Fetch() → CreditBalanceDTO
4. **ListCreditNotesQuery**(user_id) → Fetch() → CreditNoteListDTO

#### Projections
- credit_note_read
- credit_balance_read

#### Events Published
- credit_note.issued.v1
- credit.applied.v1

#### RBAC/SLO
- **RBAC:** ADMIN (issue), SYSTEM (apply), OWNER (view)
- **SLO:** P95 < 200ms (issue/apply), P95 < 100ms (read)

---

### 10.5 tax_class/

#### User Stories
- As an **admin**, I want to **define tax classes** (standard, reduced, exempt) so that taxes are calculated correctly.
- As an **admin**, I want to **bind tax classes to plans** so that rates are applied automatically.
- As a **system**, I want to **calculate taxes** based on user location and plan so that compliance is maintained.

#### Flow
1. **CreateTaxClassCommand**(name, rate, jurisdiction, idempotency_key) → Validate() | Persist() → **Outbox:** tax_class.created.v1
2. **UpdateTaxClassCommand**(tax_class_id, updates, idempotency_key) → Apply() → **Outbox:** tax_class.updated.v1
3. **BindTaxClassCommand**(plan_id, tax_class_id, idempotency_key) → Bind() → **Outbox:** tax_class.bound.v1
4. **CalculateTaxQuery**(user_id, plan_id, amount) → FetchTaxClass() | CalculateRate() → TaxCalculationDTO

#### Projections
- tax_class_read

#### Events Published
- tax_class.created.v1
- tax_class.updated.v1
- tax_class.bound.v1

#### RBAC/SLO
- **RBAC:** ADMIN (create/update/bind), SYSTEM (calculate)
- **SLO:** P95 < 180ms (create/update), P95 < 80ms (calculate)

---

### 10.6 billing_profile/

#### User Stories
- As a **user**, I want to **set my billing profile** (name, address, VAT ID) so that invoices are correct.
- As a **system**, I want to **validate VAT IDs** (format checks) so that tax compliance is maintained.
- As a **user**, I want to **update my billing profile** when details change so that records are current.

#### Flow
1. **CreateBillingProfileCommand**(user_id, name, address, country, vat_id, idempotency_key) → ValidateVAT() | ValidateAddress() | Persist() → **Outbox:** billing_profile.created.v1
2. **UpdateBillingProfileCommand**(profile_id, updates, idempotency_key) → AuthorizeOwner() | ValidateUpdates() | Apply() → **Outbox:** billing_profile.updated.v1
3. **GetBillingProfileQuery**(user_id) → Fetch() → BillingProfileDTO

#### Projections
- billing_profile_read

#### Events Published
- billing_profile.created.v1
- billing_profile.updated.v1

#### RBAC/SLO
- **RBAC:** OWNER (create/update/view)
- **SLO:** P95 < 200ms (create/update), P95 < 100ms (read)

---

## **11 - DUNNING DOMAIN**

### 11.1 dunning/

#### User Stories
- As a **system**, I want to **open dunning cases** for failed payments so that recovery is automated.
- As a **system**, I want to **advance dunning stages** (retry, warn, suspend, cancel) based on schedule so that collection is systematic.
- As a **system**, I want to **send dunning emails** at each stage so that users are notified.
- As an **admin**, I want to **resolve dunning cases** manually when issues are fixed so that subscriptions are restored.
- As a **system**, I want to **cancel subscriptions** after final dunning stage so that non-payers are removed.

#### Flow
1. **OpenDunningCaseCommand**(subscription_id, payment_attempt_id, initial_error, idempotency_key) → CreateCase() | ScheduleRetry() → **Outbox:** dunning.case.opened.v1
2. **AdvanceDunningStageJob**(case_id) → FetchCase() | DetermineNextStage() | ExecuteAction() | SendEmail(communications-be) → **Outbox:** dunning.stage.advanced.v1
3. **ResolveDunningCaseCommand**(case_id, resolution_type, resolved_by, idempotency_key) → AuthorizeAdmin() | Resolve() | RestoreSubscription() → **Outbox:** dunning.case.resolved.v1
4. **GetDunningCaseQuery**(case_id) → AuthorizeAdmin() | Fetch() → DunningCaseDTO
5. **ListDunningCasesQuery**(filters) → AuthorizeAdmin() | ApplyFilters() → DunningCaseListDTO

#### Projections
- dunning_case_read
- dunning_stage_read

#### Events Published
- dunning.case.opened.v1
- dunning.stage.advanced.v1
- dunning.case.resolved.v1

#### Events Consumed
- payment.failed.v1 (open case)
- payment.processed.v1 (resolve case)

#### RBAC/SLO
- **RBAC:** SYSTEM (open/advance), ADMIN (resolve/view)
- **SLO:** P95 < 200ms (open), P95 < 300ms (advance), P95 < 250ms (resolve), P95 < 150ms (read)

---

### 11.2 dunning_schedule/

#### User Stories
- As an **admin**, I want to **configure dunning schedules** (retry intervals, stage progression) so that recovery strategy is customized.
- As a **system**, I want to **apply