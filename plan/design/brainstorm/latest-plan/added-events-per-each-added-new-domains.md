
# Microservices Event Overview

---

## **users-be**
**New capabilities:** `org`, `security_center`, `compliance`, `risk_signals`, `profile_depth`

### Publishes
- `user.org.created` / `user.org.updated` / `user.org.member.added` / `user.org.member.removed`  
  → `admin-be`, `jobs-be`, `contracts-be`, `financial-be`

- `user.security.finding.opened` / `user.security.finding.resolved`  
  → `admin-be`, `security ops`

- `user.compliance.status.updated` (KYC/KYB/Tax)  
  → `financial-be`, `admin-be`, `contracts-be`

- `user.risk.signal.emitted` (device, velocity, ip_reputation)  
  → `admin-be[risk]`, `financial-be[risk]`

- `user.profile.depth.updated`  
  → `search-be[pers]`, `reviews-be[eligibility]`, `subscriptions-be[entitlements]`

### Consumes
- `admin.compliance.policy.updated` (`admin-be` → enforce KYC/KYB rules)
- `financial.chargeback.created` / `financial.risk.hold.placed` (`financial-be` → raise user risk)
- `contracts.state.changed` (`contracts-be` → adjust compliance checks)

---

## **jobs-be**
**New capabilities:** `job_template`, `screening_compliance`, `sourcing_modes`, `budget_controls`, `visibility_lifecycle`

### Publishes
- `job.template.created` / `job.template.updated`  
  → `search-be[hygiene]`, `admin-be[kb]`

- `job.screening.compliance.failed` / `job.screening.compliance.passed`  
  → `admin-be[moderation]`, `users-be`

- `job.sourcing.mode.changed` (organic, invite-only, agency)  
  → `search-be[facets]`, `subscriptions-be[usage]`

- `job.budget.control.updated` (caps, spend_rate)  
  → `financial-be`, `admin-be`

- `job.visibility.state.changed` (hidden, boosted, expired)  
  → `search-be[index]`, `admin-be[moderation]`

### Consumes
- `subscriptions.entitlement.updated` (`subscriptions-be` → gate boosts/invites)
- `admin.moderation.action.applied` (`admin-be` → hide/remove job)
- `financial.invoice.overdue` (`financial-be` → restrict posting/boosts)

---

## **proposals-be**
**New capabilities:** `negotiation`, `invite`, `invite_flow`, `rate_card`

### Publishes
- `proposal.negotiation.started` / `updated` / `concluded`  
  → `contracts-be`, `admin-be[audit]`

- `proposal.invite.sent` / `accepted` / `declined`  
  → `communications-be`, `search-be[signals]`, `admin-be`

- `proposal.invite.flow.abandoned`  
  → `analytics`, `admin-be`

- `proposal.rate_card.updated`  
  → `contracts-be`, `financial-be[quotes]`

### Consumes
- `subscriptions.connects.debited` / `credited` (`subscriptions-be` → enforce send limits)
- `jobs.visibility.state.changed` (`jobs-be` → allow boosts/invites)
- `admin.moderation.action.applied` (`admin-be` → freeze negotiation thread)

---

## **contracts-be**
**New capabilities:** `sow`, `financial_hold`

### Publishes
- `contract.sow.created` / `updated` / `approved`  
  → `financial-be[invoicing]`, `admin-be[disputes]`

- `contract.financial_hold.placed` / `released`  
  → `financial-be[risk]`, `admin-be[risk]`, `users-be[status]`

### Consumes
- `proposal.negotiation.concluded` (`proposals-be` → finalize terms)
- `financial.chargeback.created` / `reserve.updated` (`financial-be` → place/release holds)
- `admin.dispute.decision.made` (`admin-be` → amend SOW / apply hold)

---

## **financial-be**
**New capabilities:** `ledger_journal`, `fee_v2`, `fx`, `risk`

### Publishes
- `financial.ledger.journal.posted`  
  → `admin-be[reporting]`, `subscriptions-be[billing]`, `contracts-be`

- `financial.fee.schedule.updated` (`fee_v2`)  
  → `jobs-be[budget]`, `subscriptions-be[billing]`

- `financial.fx.rate.updated`  
  → `contracts-be`, `subscriptions-be`, `search-be[pricing facets]`

- `financial.risk.alert.emitted` (velocity, anomaly)  
  → `admin-be[risk]`, `users-be[risk]`

- `financial.chargeback.created` / `updated`  
  → `admin-be[risk]`, `contracts-be`, `users-be`

### Consumes
- `admin.risk.rule.updated` (`admin-be` → recompute alerts)
- `contracts.financial_hold.placed` / `released` (`contracts-be` → account state)
- `subscriptions.billing.invoice.requested` (`subscriptions-be` → generate/export)

---

## **communications-be**
**New capabilities:** `in_app_notification`, `notification_queue`, `delivery_log`, `system_message`, `call`, `calendar_invite`

### Publishes
- `comm.in_app.sent` / `read`  
  → `admin-be[audit]`, `analytics`

- `comm.queue.enqueued` / `dequeued`  
  → `admin-be[ops]`

- `comm.delivery.logged` (status, channel)  
  → `admin-be[audit]`, `users-be[security notifications]`

- `comm.system.message.published`  
  → `admin-be`, `users-be`

- `comm.call.started` / `ended` / `recording.ready`  
  → `admin-be[audit]`, `reviews-be[proof]`

- `comm.calendar.invite.sent` / `response.received`  
  → `contracts-be`, `admin-be`

### Consumes
- `proposal.invite.sent` (`proposals-be` → notify)
- `admin.announcement.created` (`admin-be` → fan-out)
- `reviews.double_blind.window.opened` (`reviews-be` → remind)

---

## **storage-be**
**New capabilities:** `policy`, `lifecycle`, `linking`, `preview`

### Publishes
- `file.policy.updated` / `violation.detected`
- `file.lifecycle.soft_deleted` / `restored`
- `file.lifecycle.legal_hold.placed` / `removed`
- `file.link.signed_url.created` / `revoked`
- `file.download.logged`
- `file.preview.generated` (thumbnails/variants ready)  
  → `search-be[facets]`, `admin-be`

### Consumes
- `user.updated` / `contract.state.changed` (enforce retention/holds)
- `admin.policy.updated` (centralized DLP patterns)

---

## **search-be**
**New capabilities:** `taxonomy`, `ltr`, `facets`, `personalization`, `index_hygiene`

### Publishes
- `search.taxonomy.synonym.updated`  
  → `jobs-be`, `proposals-be`, `admin-be`

- `search.ltr.model.signal.recorded`  
  → `analytics`, `admin-be`

- `search.facets.schema.updated`  
  → `web/apps cache warmers`

- `search.personalization.profile.updated`  
  → recommendations (if any)

- `search.index.hygiene.de_dupe.performed` / `archive.marked`  
  → `admin-be[audit]`

### Consumes
- `users.profile.depth.updated` / `users.risk.signal.emitted` (`users-be` → ranking/personalization)
- `jobs.visibility.state.changed` / `job.template.updated` (`jobs-be` → reindex)
- `storage.file.preview.generated` (`storage-be` → enrich search doc)
- `admin.experiment.toggled` (`admin-be` → feature flags for ranking)

---

## **reviews-be**
**New capabilities:** `double_blind`, `weighting`, `public_response`, `private_feedback`, `reputation_model`, `abuse_controls`

### Publishes
- `review.double_blind.window.opened` / `closed`  
  → `communications-be` (reminders)

- `review.weighting.schema.updated`  
  → `analytics`, `admin-be`

- `review.public_response.added`  
  → `admin-be[audit]`, `users-be[profile]`

- `review.private_feedback.submitted`  
  → `admin-be[risk]`, `analytics`

- `reputation.score.updated` / `eligibility.top_rated.updated`  
  → `users-be`, `admin-be[badges]`

- `review.abuse.auto_flagged`  
  → `admin-be[moderation]`

### Consumes
- `contracts.state.changed` (`contracts-be` → open review window)
- `admin.moderation.action.applied` (`admin-be` → hide/remove review)
- `communications.delivery.logged` (`communications-be` → SLA tracking on reminders)

---

## **subscriptions-be**
**New capabilities:** `entitlement`, `seat_billing`, `billing_export`, `plan/connect/usage extensions`

### Publishes
- `subscriptions.entitlement.updated` (features: boosts, connects, invites)  
  → `jobs-be`, `proposals-be`, `search-be`

- `connects.purchased` / `debited` / `refunded` / `expired`  
  → `proposals-be`, `admin-be[audit]`

- `billing.seat.allocated` / `deallocated` / `overage.incurred`  
  → `financial-be[invoice]`, `admin-be`

- `billing.invoice.generated` / `exported`  
  → `financial-be`, `admin-be[reporting]`

- `usage.counter.incremented` / `limit.reached`  
  → `jobs-be`, `proposals-be`, `admin-be`

### Consumes
- `admin.feature_flag.updated` (`admin-be` → toggle entitlements)
- `financial.chargeback.created` (`financial-be` → reverse entitlements/connects)
- `users.org.member.added` / `removed` (`users-be` → seat updates)

---

## **admin-be**
**New capabilities:** `fraud_review`, `user_report`, `risk_management`, `system_config`, `data_export`, `ticket_note`

### Publishes
- `admin.case.fraud_review.opened` / `updated` / `closed`  
  → `financial-be[risk]`, `users-be[risk]`

- `admin.case.user_report.created` / `triaged` / `actioned` / `dismissed`  
  → `users-be`, `reviews-be`, `communications-be`

- `admin.risk.hold.placed` / `released` / `reserve.set` / `chargeback.review.requested`  
  → `financial-be`, `contracts-be`

- `admin.config.updated` / `admin.feature_flag.updated`  
  → `search-be`, `subscriptions-be`, `users-be`, `jobs-be`

- `admin.audit.action.logged`  
  → everyone (audit streams as needed)

- `admin.data_export.requested` / `approved` / `generated` / `delivered` / `revoked`  
  → `storage-be` (if artifact), `users-be` (notify subject)

- `admin.ticket.note.added`  
  → support workflows, `communications-be` (if notify)

### Consumes
- `financial.risk.alert.emitted` / `chargeback.created` (`financial-be` → open fraud review)
- `reviews.abuse.auto_flagged` (`reviews-be` → queue moderation)
- `search.index.hygiene.de_dupe.performed` (`search-be` → audit trail)
- `storage.file.policy.violation.detected` (`storage-be` → moderation)
- `users.risk.signal.emitted` (`users-be` → risk dashboard)

---

## **Cross-Domain Event Summary**

### **Feature Flags & Experiments**
- **Source:** `admin.feature_flag.updated`
- **Consumers:** `search-be`, `subscriptions-be`, `users-be`, `jobs-be`

### **Risk & Financial Holds**
- `admin.risk.hold.placed` / `hold.released`  
  → Published by `admin-be`, consumed by `financial-be`, `contracts-be`.

- `financial.risk.alert.emitted`  
  → Published by `financial-be`, consumed by `admin-be`.

### **Exports & Audit**
- **Data Export Flow:** `admin.data_export.*` from `admin-be`.
- **Audit Trail Feeds to admin-be:**  
  - `file.download.logged` (`storage-be`)  
  - `comm.delivery.logged` (`communications-be`)
