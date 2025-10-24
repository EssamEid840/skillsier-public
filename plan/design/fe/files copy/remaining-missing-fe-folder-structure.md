# Remaining Missing Frontend Folder Structure for Skillsier Application
## Completing All Requirements from fe-folder-structure-prompt.md

> **Note**: This document contains ONLY the folder structure elements that are:
> 1. Required by `fe-folder-structure-prompt.md`
> 2. NOT present in `combined-folder-structure.md`
> 3. NOT present in `missing-fe-folder-structure.md`
> 4. NOT present in `additional-missing-fe-folder-structure.md`

---

## I. Missing Dashboard Routes - Enhanced Features

### 1. Jobs Section - Additional Features

```
apps/web/src/app/[locale]/(dashboard)/
│
├── jobs/
│   ├── insights/
│   │   └── page.tsx  # Market insights for job posting
│   │       # - Recommended rates
│   │       # - Competition analysis
│   │       # - Time-to-hire estimates
│   │       # - Skill demand trends
│   │       # BE: jobs-be/analytics, search-be/trending
│   │       # GET /v1/jobs/market-insights
│   │       # GET /v1/trending/skills
│   │
│   ├── batch/
│   │   ├── create/
│   │   │   └── page.tsx  # Bulk job creation
│   │   │       # - CSV upload
│   │   │       # - Template selection
│   │   │       # - Preview and publish
│   │   │       # BE: jobs-be/job
│   │   │       # POST /v1/jobs/batch
│   │   │
│   │   └── manage/
│   │       └── page.tsx  # Batch operations
│   │           # - Bulk edit
│   │           # - Bulk status change
│   │           # - Bulk archive
│   │           # BE: jobs-be/job
│   │           # PATCH /v1/jobs/batch
│   │
│   └── scheduling/
│       └── page.tsx  # Schedule job postings
│           # - Future publish dates
│           # - Auto-repost settings
│           # - Expiration reminders
│           # BE: jobs-be/job
│           # POST /v1/jobs/{job_id}/schedule
```

### 2. Proposals Section - Enhanced Features

```
apps/web/src/app/[locale]/(dashboard)/
│
├── proposals/
│   ├── [proposalId]/
│   │   ├── negotiations/
│   │   │   └── page.tsx  # Proposal negotiations history
│   │   │       # - Negotiation timeline
│   │   │       # - Counter-offers
│   │   │       # - Terms evolution
│   │   │       # BE: proposals-be/negotiation
│   │   │       # GET /v1/proposals/{proposal_id}/negotiations
│   │   │
│   │   ├── compare/
│   │   │   └── page.tsx  # Compare with other proposals
│   │   │       # - Side-by-side comparison
│   │   │       # - Highlight differences
│   │   │       # BE: proposals-be/proposal
│   │   │       # GET /v1/proposals/compare?ids=...
│   │   │
│   │   └── feedback/
│   │       └── page.tsx  # Proposal feedback
│   │           # - Client feedback
│   │           # - Improvement suggestions
│   │           # BE: proposals-be/proposal
│   │           # GET /v1/proposals/{proposal_id}/feedback
│   │
│   ├── insights/
│   │   └── page.tsx  # Proposal insights
│   │       # - Win rate analysis
│   │       # - Optimal pricing insights
│   │       # - Response time impact
│   │       # - Proposal quality score
│   │       # BE: proposals-be/analytics
│   │       # GET /v1/proposals/insights
│   │
│   ├── benchmarking/
│   │   └── page.tsx  # Benchmark against market
│   │       # - Rate comparisons
│   │       # - Proposal quality metrics
│   │       # - Success factors
│   │       # BE: proposals-be/analytics, search-be/trending
│   │       # GET /v1/proposals/benchmarking
│   │
│   └── archived/
│       └── page.tsx  # Archived proposals
│           # - Old proposals
│           # - Historical reference
│           # BE: proposals-be/proposal
│           # GET /v1/proposals?status=archived
```

### 3. Contracts Section - Additional Features

```
apps/web/src/app/[locale]/(dashboard)/
│
├── contracts/
│   ├── [contractId]/
│   │   ├── change-orders/
│   │   │   ├── [orderId]/
│   │   │   │   ├── approve/
│   │   │   │   │   └── page.tsx  # Approve change order
│   │   │   │   │       # BE: contracts-be/change_order
│   │   │   │   │       # POST /v1/contracts/{contract_id}/change-orders/{order_id}/approve
│   │   │   │   ├── reject/
│   │   │   │   │   └── page.tsx  # Reject change order
│   │   │   │   │       # BE: contracts-be/change_order
│   │   │   │   │       # POST /v1/contracts/{contract_id}/change-orders/{order_id}/reject
│   │   │   │   └── page.tsx  # Change order details
│   │   │   │       # BE: contracts-be/change_order
│   │   │   │       # GET /v1/contracts/{contract_id}/change-orders/{order_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create change order
│   │   │   │       # - Scope modifications
│   │   │   │       # - Budget adjustments
│   │   │   │       # - Timeline changes
│   │   │   │       # BE: contracts-be/change_order
│   │   │   │       # POST /v1/contracts/{contract_id}/change-orders
│   │   │   └── page.tsx  # Change orders list
│   │   │       # BE: contracts-be/change_order
│   │   │       # GET /v1/contracts/{contract_id}/change-orders
│   │   │
│   │   ├── compliance/
│   │   │   └── page.tsx  # Contract compliance tracking
│   │   │       # - KPIs monitoring
│   │   │       # - SLA compliance
│   │   │       # - Penalty tracking
│   │   │       # BE: contracts-be/compliance
│   │   │       # GET /v1/contracts/{contract_id}/compliance
│   │   │
│   │   └── audit-trail/
│   │       └── page.tsx  # Complete contract audit trail
│   │           # - All changes
│   │           # - Approval history
│   │           # - Access logs
│   │           # BE: contracts-be/contract, utility-be/audit
│   │           # GET /v1/contracts/{contract_id}/audit-trail
│   │
│   ├── calendar/
│   │   └── page.tsx  # Contracts calendar view
│   │       # - Milestone timeline
│   │       # - Payment schedules
│   │       # - Deliverable deadlines
│   │       # BE: contracts-be/contract, contracts-be/milestone
│   │       # GET /v1/contracts/calendar
│   │
│   └── benchmarking/
│       └── page.tsx  # Contract performance benchmarking
│           # - Industry comparisons
│           # - Performance metrics
│           # - Best practices
│           # BE: contracts-be/analytics
│           # GET /v1/contracts/benchmarking
```

### 4. Financial Section - Enhanced Features

```
apps/web/src/app/[locale]/(dashboard)/
│
├── financial/
│   ├── reconciliation/
│   │   └── page.tsx  # Financial reconciliation
│   │       # - Match transactions
│   │       # - Resolve discrepancies
│   │       # - Bank statement upload
│   │       # BE: financial-be/reconciliation
│   │       # GET /v1/financial/reconciliation
│   │       # POST /v1/financial/reconciliation/upload
│   │
│   ├── forecasting/
│   │   └── page.tsx  # Revenue/expense forecasting
│   │       # - Projected earnings
│   │       # - Cash flow predictions
│   │       # - Scenario planning
│   │       # BE: financial-be/analytics
│   │       # GET /v1/financial/forecasts
│   │
│   ├── chargebacks/
│   │   ├── [chargebackId]/
│   │   │   ├── respond/
│   │   │   │   └── page.tsx  # Respond to chargeback
│   │   │   │       # - Upload evidence
│   │   │   │       # - Add defense statement
│   │   │   │       # BE: financial-be/chargeback
│   │   │   │       # POST /v1/financial/chargebacks/{chargeback_id}/respond
│   │   │   └── page.tsx  # Chargeback details
│   │   │       # BE: financial-be/chargeback
│   │   │       # GET /v1/financial/chargebacks/{chargeback_id}
│   │   └── page.tsx  # Chargebacks list
│   │       # - Active disputes
│   │       # - Resolution status
│   │       # BE: financial-be/chargeback
│   │       # GET /v1/financial/chargebacks
│   │
│   ├── payment-methods/
│   │   ├── [methodId]/
│   │   │   ├── edit/
│   │   │   │   └── page.tsx  # Edit payment method
│   │   │   │       # BE: financial-be/payment_method
│   │   │   │       # PUT /v1/financial/payment-methods/{method_id}
│   │   │   ├── remove/
│   │   │   │   └── page.tsx  # Remove payment method
│   │   │   │       # BE: financial-be/payment_method
│   │   │   │       # DELETE /v1/financial/payment-methods/{method_id}
│   │   │   └── verify/
│   │   │       └── page.tsx  # Verify payment method
│   │   │           # - Micro-deposit verification
│   │   │           # - Card verification
│   │   │           # BE: financial-be/payment_method
│   │   │           # POST /v1/financial/payment-methods/{method_id}/verify
│   │   ├── add/
│   │   │   └── page.tsx  # Add payment method
│   │   │       # - Credit/debit card
│   │   │       # - Bank account (ACH)
│   │   │       # - Digital wallet
│   │   │       # BE: financial-be/payment_method
│   │   │       # POST /v1/financial/payment-methods
│   │   └── page.tsx  # Payment methods list
│   │       # BE: financial-be/payment_method
│   │       # GET /v1/financial/payment-methods
│   │
│   ├── payout-methods/
│   │   ├── [methodId]/
│   │   │   ├── edit/
│   │   │   │   └── page.tsx  # Edit payout method
│   │   │   │       # BE: financial-be/payout_method
│   │   │   │       # PUT /v1/financial/payout-methods/{method_id}
│   │   │   ├── remove/
│   │   │   │   └── page.tsx  # Remove payout method
│   │   │   │       # BE: financial-be/payout_method
│   │   │   │       # DELETE /v1/financial/payout-methods/{method_id}
│   │   │   └── verify/
│   │   │       └── page.tsx  # Verify payout method
│   │   │           # BE: financial-be/payout_method
│   │   │           # POST /v1/financial/payout-methods/{method_id}/verify
│   │   ├── add/
│   │   │   └── page.tsx  # Add payout method
│   │   │       # - Bank account
│   │   │       # - PayPal
│   │   │       # - Wire transfer
│   │   │       # BE: financial-be/payout_method
│   │   │       # POST /v1/financial/payout-methods
│   │   └── page.tsx  # Payout methods list
│   │       # BE: financial-be/payout_method
│   │       # GET /v1/financial/payout-methods
│   │
│   └── cost-centers/
│       ├── [centerId]/
│       │   ├── edit/
│       │   │   └── page.tsx  # Edit cost center
│       │   │       # BE: financial-be/cost_center
│       │   │       # PUT /v1/financial/cost-centers/{center_id}
│       │   ├── analytics/
│       │   │   └── page.tsx  # Cost center analytics
│       │   │       # BE: financial-be/cost_center, financial-be/analytics
│       │   │       # GET /v1/financial/cost-centers/{center_id}/analytics
│       │   └── page.tsx  # Cost center details
│       │       # BE: financial-be/cost_center
│       │       # GET /v1/financial/cost-centers/{center_id}
│       ├── create/
│       │   └── page.tsx  # Create cost center
│       │       # BE: financial-be/cost_center
│       │       # POST /v1/financial/cost-centers
│       └── page.tsx  # Cost centers list
│           # - Department/project budgets
│           # - Spend tracking
│           # BE: financial-be/cost_center
│           # GET /v1/financial/cost-centers
```

### 5. Teams/Organizations - Additional Features

```
apps/web/src/app/[locale]/(dashboard)/
│
├── teams/
│   ├── [teamId]/
│   │   ├── spending-controls/
│   │   │   └── page.tsx  # Team spending controls
│   │   │       # - Approval workflows
│   │   │       # - Spending limits
│   │   │       # - Auto-approval rules
│   │   │       # BE: users-be/team, financial-be/budget
│   │   │       # GET /v1/teams/{team_id}/spending-controls
│   │   │       # PUT /v1/teams/{team_id}/spending-controls
│   │   │
│   │   ├── compliance/
│   │   │   └── page.tsx  # Team compliance dashboard
│   │   │       # - Document status
│   │   │       # - Training completion
│   │   │       # - Policy acknowledgments
│   │   │       # BE: users-be/team, admin-be/business_verification
│   │   │       # GET /v1/teams/{team_id}/compliance
│   │   │
│   │   ├── performance/
│   │   │   └── page.tsx  # Team performance metrics
│   │   │       # - Productivity stats
│   │   │       # - Quality metrics
│   │   │       # - Member contributions
│   │   │       # BE: users-be/team, contracts-be/analytics
│   │   │       # GET /v1/teams/{team_id}/performance
│   │   │
│   │   └── hierarchy/
│   │       └── page.tsx  # Team hierarchy management
│   │           # - Organizational chart
│   │           # - Reporting structure
│   │           # - Role relationships
│   │           # BE: users-be/team
│   │           # GET /v1/teams/{team_id}/hierarchy
│   │
│   └── integrations/
│       ├── [integrationId]/
│       │   ├── configure/
│       │   │   └── page.tsx  # Configure integration
│       │   │       # BE: users-be/integration
│       │   │       # PUT /v1/teams/integrations/{integration_id}
│       │   ├── logs/
│       │   │   └── page.tsx  # Integration logs
│       │   │       # BE: users-be/integration, utility-be/audit
│       │   │       # GET /v1/teams/integrations/{integration_id}/logs
│       │   └── page.tsx  # Integration details
│       │       # BE: users-be/integration
│       │       # GET /v1/teams/integrations/{integration_id}
│       ├── available/
│       │   └── page.tsx  # Available integrations
│       │       # - Slack, JIRA, etc.
│       │       # - Feature descriptions
│       │       # BE: users-be/integration
│       │       # GET /v1/teams/integrations/available
│       └── page.tsx  # Active integrations list
│           # BE: users-be/integration
│           # GET /v1/teams/integrations
```

### 6. Settings - Additional Features

```
apps/web/src/app/[locale]/(dashboard)/
│
├── settings/
│   ├── developer/
│   │   ├── api-keys/
│   │   │   ├── [keyId]/
│   │   │   │   ├── regenerate/
│   │   │   │   │   └── page.tsx  # Regenerate API key
│   │   │   │   │       # BE: users-be/api_key
│   │   │   │   │       # POST /v1/developer/api-keys/{key_id}/regenerate
│   │   │   │   ├── revoke/
│   │   │   │   │   └── page.tsx  # Revoke API key
│   │   │   │   │       # BE: users-be/api_key
│   │   │   │   │       # DELETE /v1/developer/api-keys/{key_id}
│   │   │   │   └── page.tsx  # API key details
│   │   │   │       # BE: users-be/api_key
│   │   │   │       # GET /v1/developer/api-keys/{key_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create API key
│   │   │   │       # - Name and permissions
│   │   │   │       # - Expiration settings
│   │   │   │       # BE: users-be/api_key
│   │   │   │       # POST /v1/developer/api-keys
│   │   │   └── page.tsx  # API keys list
│   │   │       # BE: users-be/api_key
│   │   │       # GET /v1/developer/api-keys
│   │   │
│   │   ├── webhooks/
│   │   │   ├── [webhookId]/
│   │   │   │   ├── edit/
│   │   │   │   │   └── page.tsx  # Edit webhook
│   │   │   │   │       # BE: users-be/webhook
│   │   │   │   │       # PUT /v1/developer/webhooks/{webhook_id}
│   │   │   │   ├── test/
│   │   │   │   │   └── page.tsx  # Test webhook
│   │   │   │   │       # BE: users-be/webhook
│   │   │   │   │       # POST /v1/developer/webhooks/{webhook_id}/test
│   │   │   │   ├── logs/
│   │   │   │   │   └── page.tsx  # Webhook delivery logs
│   │   │   │   │       # BE: users-be/webhook
│   │   │   │   │       # GET /v1/developer/webhooks/{webhook_id}/logs
│   │   │   │   └── page.tsx  # Webhook details
│   │   │   │       # BE: users-be/webhook
│   │   │   │       # GET /v1/developer/webhooks/{webhook_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create webhook
│   │   │   │       # - Endpoint URL
│   │   │   │       # - Event subscriptions
│   │   │   │       # - Secret key
│   │   │   │       # BE: users-be/webhook
│   │   │   │       # POST /v1/developer/webhooks
│   │   │   └── page.tsx  # Webhooks list
│   │   │       # BE: users-be/webhook
│   │   │       # GET /v1/developer/webhooks
│   │   │
│   │   ├── oauth-apps/
│   │   │   ├── [appId]/
│   │   │   │   ├── edit/
│   │   │   │   │   └── page.tsx  # Edit OAuth app
│   │   │   │   │       # BE: users-be/oauth_app
│   │   │   │   │       # PUT /v1/developer/oauth-apps/{app_id}
│   │   │   │   ├── credentials/
│   │   │   │   │   └── page.tsx  # OAuth credentials
│   │   │   │   │       # - Client ID/Secret
│   │   │   │   │       # - Regenerate secret
│   │   │   │   │       # BE: users-be/oauth_app
│   │   │   │   │       # POST /v1/developer/oauth-apps/{app_id}/regenerate-secret
│   │   │   │   └── page.tsx  # OAuth app details
│   │   │   │       # BE: users-be/oauth_app
│   │   │   │       # GET /v1/developer/oauth-apps/{app_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create OAuth app
│   │   │   │       # BE: users-be/oauth_app
│   │   │   │       # POST /v1/developer/oauth-apps
│   │   │   └── page.tsx  # OAuth apps list
│   │   │       # BE: users-be/oauth_app
│   │   │       # GET /v1/developer/oauth-apps
│   │   │
│   │   └── page.tsx  # Developer settings hub
│   │       # BE: None (navigation)
│   │
│   ├── labs/
│   │   └── page.tsx  # Experimental features
│   │       # - Beta feature toggles
│   │       # - Early access programs
│   │       # BE: utility-be/feature_flags
│   │       # GET /v1/labs/features
│   │       # PUT /v1/labs/features/{feature_id}/toggle
│   │
│   ├── authorized-apps/
│   │   ├── [appId]/
│   │   │   └── revoke/
│   │   │       └── page.tsx  # Revoke app access
│   │   │           # BE: users-be/oauth_token
│   │   │           # DELETE /v1/settings/authorized-apps/{app_id}
│   │   └── page.tsx  # Authorized apps list
│   │       # - Third-party app permissions
│   │       # - Last used dates
│   │       # BE: users-be/oauth_token
│   │       # GET /v1/settings/authorized-apps
│   │
│   └── advanced/
│       └── page.tsx  # Advanced settings
│           # - Debug mode
│           # - Experimental APIs
│           # - Performance monitoring opt-in
│           # BE: users-be/preferences
│           # GET /v1/settings/advanced
│           # PUT /v1/settings/advanced
```

---

## II. Missing Admin Panel Routes

### 1. Admin Panel - Core Management

```
apps/web/src/app/[locale]/(admin)/
│
├── admin/
│   ├── dashboard/
│   │   └── page.tsx  # Admin dashboard
│   │       # - Key metrics
│   │       # - Pending actions
│   │       # - System alerts
│   │       # BE: admin-be/dashboard
│   │       # GET /v1/admin/dashboard
│   │
│   ├── break-glass/
│   │   ├── request/
│   │   │   └── page.tsx  # Request break-glass access
│   │   │       # - Justification
│   │   │       # - Duration request
│   │   │       # BE: admin-be/admin_session
│   │   │       # POST /v1/admin/break-glass/request
│   │   │
│   │   ├── approve/
│   │   │   └── page.tsx  # Approve break-glass requests
│   │   │       # - Two-person rule
│   │   │       # - Request review
│   │   │       # BE: admin-be/admin_session
│   │   │       # POST /v1/admin/break-glass/approve/{request_id}
│   │   │
│   │   └── active/
│   │       └── page.tsx  # Active admin sessions
│   │           # - Time-boxed access monitoring
│   │           # - Force termination
│   │           # BE: admin-be/admin_session
│   │           # GET /v1/admin/break-glass/active
│   │
│   ├── two-person-rules/
│   │   ├── [ruleId]/
│   │   │   ├── approve/
│   │   │   │   └── page.tsx  # Approve rule change
│   │   │   │       # BE: admin-be/change_approval
│   │   │   │       # POST /v1/admin/two-person-rules/{rule_id}/approve
│   │   │   └── page.tsx  # Rule details
│   │   │       # BE: admin-be/change_approval
│   │   │       # GET /v1/admin/two-person-rules/{rule_id}
│   │   ├── pending/
│   │   │   └── page.tsx  # Pending approvals
│   │   │       # BE: admin-be/change_approval
│   │   │       # GET /v1/admin/two-person-rules?status=pending
│   │   └── page.tsx  # Two-person rules dashboard
│   │       # BE: admin-be/change_approval
│   │       # GET /v1/admin/two-person-rules
│   │
│   ├── kyc-cases/
│   │   ├── [caseId]/
│   │   │   ├── review/
│   │   │   │   └── page.tsx  # Review KYC case
│   │   │   │       # - Document verification
│   │   │   │       # - Approve/reject/escalate
│   │   │   │       # BE: admin-be/kyc_case
│   │   │   │       # POST /v1/admin/kyc-cases/{case_id}/review
│   │   │   ├── reopen/
│   │   │   │   └── page.tsx  # Reopen KYC case
│   │   │   │       # BE: admin-be/kyc_case
│   │   │   │       # POST /v1/admin/kyc-cases/{case_id}/reopen
│   │   │   ├── documents/
│   │   │   │   └── page.tsx  # Case documents viewer
│   │   │   │       # BE: admin-be/kyc_case, storage-be/asset
│   │   │   │       # GET /v1/admin/kyc-cases/{case_id}/documents
│   │   │   └── page.tsx  # KYC case details
│   │   │       # BE: admin-be/kyc_case
│   │   │       # GET /v1/admin/kyc-cases/{case_id}
│   │   ├── queue/
│   │   │   └── page.tsx  # KYC queue
│   │   │       # - Prioritization
│   │   │       # - Assignment
│   │   │       # BE: admin-be/kyc_case
│   │   │       # GET /v1/admin/kyc-cases/queue
│   │   └── page.tsx  # KYC cases dashboard
│   │       # BE: admin-be/kyc_case
│   │       # GET /v1/admin/kyc-cases
│   │
│   ├── business-verification/
│   │   ├── [verificationId]/
│   │   │   ├── review/
│   │   │   │   └── page.tsx  # Review business verification
│   │   │   │       # - Company evidence review
│   │   │   │       # - Decision making
│   │   │   │       # BE: admin-be/business_verification
│   │   │   │       # POST /v1/admin/business-verification/{verification_id}/review
│   │   │   ├── reverify/
│   │   │   │   └── page.tsx  # Request reverification
│   │   │   │       # BE: admin-be/business_verification
│   │   │   │       # POST /v1/admin/business-verification/{verification_id}/reverify
│   │   │   └── page.tsx  # Verification details
│   │   │       # BE: admin-be/business_verification
│   │   │       # GET /v1/admin/business-verification/{verification_id}
│   │   ├── pending/
│   │   │   └── page.tsx  # Pending verifications
│   │   │       # BE: admin-be/business_verification
│   │   │       # GET /v1/admin/business-verification?status=pending
│   │   └── page.tsx  # Business verification dashboard
│   │       # BE: admin-be/business_verification
│   │       # GET /v1/admin/business-verification
│   │
│   ├── refund-cases/
│   │   ├── [caseId]/
│   │   │   ├── investigate/
│   │   │   │   └── page.tsx  # Investigate refund case
│   │   │   │       # - Evidence review
│   │   │   │       # - Transaction history
│   │   │   │       # BE: admin-be/refund_case
│   │   │   │       # GET /v1/admin/refund-cases/{case_id}/investigation
│   │   │   ├── decision/
│   │   │   │   └── page.tsx  # Make refund decision
│   │   │   │       # - Approve/deny
│   │   │   │       # - Partial refund
│   │   │   │       # BE: admin-be/refund_case
│   │   │   │       # POST /v1/admin/refund-cases/{case_id}/decision
│   │   │   └── page.tsx  # Refund case details
│   │   │       # BE: admin-be/refund_case
│   │   │       # GET /v1/admin/refund-cases/{case_id}
│   │   ├── pending/
│   │   │   └── page.tsx  # Pending refund cases
│   │   │       # BE: admin-be/refund_case
│   │   │       # GET /v1/admin/refund-cases?status=pending
│   │   └── page.tsx  # Refund cases dashboard
│   │       # BE: admin-be/refund_case
│   │       # GET /v1/admin/refund-cases
│   │
│   ├── goodwill-credits/
│   │   ├── [creditId]/
│   │   │   ├── approve/
│   │   │   │   └── page.tsx  # Approve goodwill credit
│   │   │   │       # BE: admin-be/goodwill_credit
│   │   │   │       # POST /v1/admin/goodwill-credits/{credit_id}/approve
│   │   │   └── page.tsx  # Goodwill credit details
│   │   │       # BE: admin-be/goodwill_credit
│   │   │       # GET /v1/admin/goodwill-credits/{credit_id}
│   │   ├── issue/
│   │   │   └── page.tsx  # Issue goodwill credit
│   │   │       # - User selection
│   │   │       # - Amount and reason
│   │   │       # BE: admin-be/goodwill_credit
│   │   │       # POST /v1/admin/goodwill-credits
│   │   └── page.tsx  # Goodwill credits dashboard
│   │       # BE: admin-be/goodwill_credit
│   │       # GET /v1/admin/goodwill-credits
│   │
│   ├── moderation/
│   │   ├── reports/
│   │   │   ├── [reportId]/
│   │   │   │   ├── review/
│   │   │   │   │   └── page.tsx  # Review report
│   │   │   │   │       # - Content review
│   │   │   │   │       # - Take action
│   │   │   │   │       # BE: admin-be/moderation
│   │   │   │   │       # POST /v1/admin/moderation/reports/{report_id}/review
│   │   │   │   └── page.tsx  # Report details
│   │   │   │       # BE: admin-be/moderation
│   │   │   │       # GET /v1/admin/moderation/reports/{report_id}
│   │   │   ├── queue/
│   │   │   │   └── page.tsx  # Moderation queue
│   │   │   │       # BE: admin-be/moderation
│   │   │   │       # GET /v1/admin/moderation/reports/queue
│   │   │   └── page.tsx  # Reports dashboard
│   │   │       # BE: admin-be/moderation
│   │   │       # GET /v1/admin/moderation/reports
│   │   │
│   │   ├── actions/
│   │   │   ├── [actionId]/
│   │   │   │   ├── appeal/
│   │   │   │   │   └── page.tsx  # Review appeal
│   │   │   │   │       # BE: admin-be/moderation
│   │   │   │   │       # POST /v1/admin/moderation/actions/{action_id}/appeal
│   │   │   │   └── page.tsx  # Action details
│   │   │   │       # BE: admin-be/moderation
│   │   │   │       # GET /v1/admin/moderation/actions/{action_id}
│   │   │   └── page.tsx  # Moderation actions
│   │   │       # - Warnings
│   │   │       # - Suspensions
│   │   │       # - Bans
│   │   │       # BE: admin-be/moderation
│   │   │       # GET /v1/admin/moderation/actions
│   │   │
│   │   └── patterns/
│   │       └── page.tsx  # Abuse patterns detection
│   │           # - Pattern analysis
│   │           # - Risk scoring
│   │           # BE: admin-be/moderation
│   │           # GET /v1/admin/moderation/patterns
│   │
│   ├── financial-ops/
│   │   ├── reconciliation/
│   │   │   ├── [reconciliationId]/
│   │   │   │   ├── resolve/
│   │   │   │   │   └── page.tsx  # Resolve reconciliation
│   │   │   │   │       # BE: financial-be/reconciliation
│   │   │   │   │       # POST /v1/admin/reconciliation/{reconciliation_id}/resolve
│   │   │   │   └── page.tsx  # Reconciliation details
│   │   │   │       # BE: financial-be/reconciliation
│   │   │   │       # GET /v1/admin/reconciliation/{reconciliation_id}
│   │   │   ├── pending/
│   │   │   │   └── page.tsx  # Pending reconciliations
│   │   │   │       # BE: financial-be/reconciliation
│   │   │   │       # GET /v1/admin/reconciliation?status=pending
│   │   │   └── page.tsx  # Reconciliation dashboard
│   │   │       # BE: financial-be/reconciliation
│   │   │       # GET /v1/admin/reconciliation
│   │   │
│   │   ├── payouts/
│   │   │   ├── [payoutId]/
│   │   │   │   ├── review/
│   │   │   │   │   └── page.tsx  # Review payout
│   │   │   │   │       # BE: financial-be/payout
│   │   │   │   │       # POST /v1/admin/payouts/{payout_id}/review
│   │   │   │   ├── retry/
│   │   │   │   │   └── page.tsx  # Retry failed payout
│   │   │   │   │       # BE: financial-be/payout
│   │   │   │   │       # POST /v1/admin/payouts/{payout_id}/retry
│   │   │   │   └── page.tsx  # Payout details
│   │   │   │       # BE: financial-be/payout
│   │   │   │       # GET /v1/admin/payouts/{payout_id}
│   │   │   ├── pending/
│   │   │   │   └── page.tsx  # Pending payouts
│   │   │   │       # BE: financial-be/payout
│   │   │   │       # GET /v1/admin/payouts?status=pending
│   │   │   ├── failed/
│   │   │   │   └── page.tsx  # Failed payouts
│   │   │   │       # BE: financial-be/payout
│   │   │   │       # GET /v1/admin/payouts?status=failed
│   │   │   └── page.tsx  # Payouts dashboard
│   │   │       # BE: financial-be/payout
│   │   │       # GET /v1/admin/payouts
│   │   │
│   │   ├── tax-forms/
│   │   │   ├── [formId]/
│   │   │   │   ├── review/
│   │   │   │   │   └── page.tsx  # Review tax form
│   │   │   │   │       # BE: financial-be/tax
│   │   │   │   │       # POST /v1/admin/tax-forms/{form_id}/review
│   │   │   │   └── page.tsx  # Tax form details
│   │   │   │       # BE: financial-be/tax
│   │   │   │       # GET /v1/admin/tax-forms/{form_id}
│   │   │   ├── generate/
│   │   │   │   └── page.tsx  # Generate tax forms
│   │   │   │       # - Bulk 1099 generation
│   │   │   │       # - Tax year selection
│   │   │   │       # BE: financial-be/tax
│   │   │   │       # POST /v1/admin/tax-forms/generate
│   │   │   └── page.tsx  # Tax forms dashboard
│   │   │       # BE: financial-be/tax
│   │   │       # GET /v1/admin/tax-forms
│   │   │
│   │   └── disputes/
│   │       ├── [disputeId]/
│   │       │   ├── mediate/
│   │       │   │   └── page.tsx  # Mediate payment dispute
│   │       │   │       # BE: admin-be/dispute_resolution
│   │       │   │       # POST /v1/admin/financial-disputes/{dispute_id}/mediate
│   │       │   └── page.tsx  # Dispute details
│   │       │       # BE: admin-be/dispute_resolution
│   │       │       # GET /v1/admin/financial-disputes/{dispute_id}
│   │       └── page.tsx  # Financial disputes
│   │           # BE: admin-be/dispute_resolution
│   │           # GET /v1/admin/financial-disputes
│   │
│   ├── communications-ops/
│   │   ├── broadcasts/
│   │   │   ├── [broadcastId]/
│   │   │   │   ├── edit/
│   │   │   │   │   └── page.tsx  # Edit broadcast
│   │   │   │   │       # BE: communications-be/broadcast
│   │   │   │   │       # PUT /v1/admin/broadcasts/{broadcast_id}
│   │   │   │   ├── schedule/
│   │   │   │   │   └── page.tsx  # Schedule broadcast
│   │   │   │   │       # BE: communications-be/broadcast
│   │   │   │   │       # POST /v1/admin/broadcasts/{broadcast_id}/schedule
│   │   │   │   ├── analytics/
│   │   │   │   │   └── page.tsx  # Broadcast analytics
│   │   │   │   │       # BE: communications-be/broadcast
│   │   │   │   │       # GET /v1/admin/broadcasts/{broadcast_id}/analytics
│   │   │   │   └── page.tsx  # Broadcast details
│   │   │   │       # BE: communications-be/broadcast
│   │   │   │       # GET /v1/admin/broadcasts/{broadcast_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create broadcast
│   │   │   │       # - Target audience
│   │   │   │       # - Message content
│   │   │   │       # - Delivery channels
│   │   │   │       # BE: communications-be/broadcast
│   │   │   │       # POST /v1/admin/broadcasts
│   │   │   └── page.tsx  # Broadcasts dashboard
│   │   │       # BE: communications-be/broadcast
│   │   │       # GET /v1/admin/broadcasts
│   │   │
│   │   ├── templates/
│   │   │   ├── [templateId]/
│   │   │   │   ├── edit/
│   │   │   │   │   └── page.tsx  # Edit template
│   │   │   │   │       # BE: communications-be/template
│   │   │   │   │       # PUT /v1/admin/templates/{template_id}
│   │   │   │   ├── test/
│   │   │   │   │   └── page.tsx  # Test template
│   │   │   │   │       # BE: communications-be/template
│   │   │   │   │       # POST /v1/admin/templates/{template_id}/test
│   │   │   │   └── page.tsx  # Template details
│   │   │   │       # BE: communications-be/template
│   │   │   │       # GET /v1/admin/templates/{template_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create template
│   │   │   │       # BE: communications-be/template
│   │   │   │       # POST /v1/admin/templates
│   │   │   └── page.tsx  # Templates dashboard
│   │   │       # BE: communications-be/template
│   │   │       # GET /v1/admin/templates
│   │   │
│   │   ├── campaigns/
│   │   │   ├── [campaignId]/
│   │   │   │   ├── edit/
│   │   │   │   │   └── page.tsx  # Edit campaign
│   │   │   │   │       # BE: communications-be/campaign
│   │   │   │   │       # PUT /v1/admin/campaigns/{campaign_id}
│   │   │   │   ├── analytics/
│   │   │   │   │   └── page.tsx  # Campaign analytics
│   │   │   │   │       # BE: communications-be/campaign
│   │   │   │   │       # GET /v1/admin/campaigns/{campaign_id}/analytics
│   │   │   │   └── page.tsx  # Campaign details
│   │   │   │       # BE: communications-be/campaign
│   │   │   │       # GET /v1/admin/campaigns/{campaign_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create campaign
│   │   │   │       # BE: communications-be/campaign
│   │   │   │       # POST /v1/admin/campaigns
│   │   │   └── page.tsx  # Campaigns dashboard
│   │   │       # BE: communications-be/campaign
│   │   │       # GET /v1/admin/campaigns
│   │   │
│   │   └── rate-limits/
│   │       └── page.tsx  # Communication rate limits
│   │           # - Configure limits
│   │           # - Monitor usage
│   │           # BE: communications-be/rate_limit
│   │           # GET /v1/admin/rate-limits
│   │           # PUT /v1/admin/rate-limits
│   │
│   ├── search-quality/
│   │   ├── synonyms/
│   │   │   ├── [synonymId]/
│   │   │   │   ├── edit/
│   │   │   │   │   └── page.tsx  # Edit synonym
│   │   │   │   │       # BE: search-be/admin
│   │   │   │   │       # PUT /v1/admin/search/synonyms/{synonym_id}
│   │   │   │   └── page.tsx  # Synonym details
│   │   │   │       # BE: search-be/admin
│   │   │   │       # GET /v1/admin/search/synonyms/{synonym_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create synonym
│   │   │   │       # BE: search-be/admin
│   │   │   │       # POST /v1/admin/search/synonyms
│   │   │   └── page.tsx  # Synonyms management
│   │   │       # BE: search-be/admin
│   │   │       # GET /v1/admin/search/synonyms
│   │   │
│   │   ├── boosts/
│   │   │   ├── [boostId]/
│   │   │   │   ├── edit/
│   │   │   │   │   └── page.tsx  # Edit boost rule
│   │   │   │   │       # BE: search-be/admin
│   │   │   │   │       # PUT /v1/admin/search/boosts/{boost_id}
│   │   │   │   └── page.tsx  # Boost details
│   │   │   │       # BE: search-be/admin
│   │   │   │       # GET /v1/admin/search/boosts/{boost_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create boost rule
│   │   │   │       # BE: search-be/admin
│   │   │   │       # POST /v1/admin/search/boosts
│   │   │   └── page.tsx  # Boost rules management
│   │   │       # BE: search-be/admin
│   │   │       # GET /v1/admin/search/boosts
│   │   │
│   │   ├── blacklists/
│   │   │   ├── [blacklistId]/
│   │   │   │   └── page.tsx  # Blacklist entry details
│   │   │   │       # BE: search-be/admin
│   │   │   │       # GET /v1/admin/search/blacklists/{blacklist_id}
│   │   │   ├── add/
│   │   │   │   └── page.tsx  # Add blacklist entry
│   │   │   │       # BE: search-be/admin
│   │   │   │       # POST /v1/admin/search/blacklists
│   │   │   └── page.tsx  # Blacklist management
│   │   │       # BE: search-be/admin
│   │   │       # GET /v1/admin/search/blacklists
│   │   │
│   │   ├── reindex/
│   │   │   └── page.tsx  # Reindex operations
│   │   │       # - Full reindex
│   │   │       # - Selective reindex
│   │   │       # - Progress monitoring
│   │   │       # BE: search-be/admin
│   │   │       # POST /v1/admin/search/reindex
│   │   │
│   │   └── analytics/
│   │       └── page.tsx  # Search analytics
│   │           # - Query performance
│   │           # - Zero-result queries
│   │           # - Popular searches
│   │           # BE: search-be/admin
│   │           # GET /v1/admin/search/analytics
│   │
│   ├── system/
│   │   ├── health/
│   │   │   └── page.tsx  # System health dashboard
│   │   │       # - Service status
│   │   │       # - Resource usage
│   │   │       # - Error rates
│   │   │       # BE: utility-be/health
│   │   │       # GET /v1/admin/system/health
│   │   │
│   │   ├── metrics/
│   │   │   └── page.tsx  # System metrics
│   │   │       # - Performance metrics
│   │   │       # - Custom dashboards
│   │   │       # BE: utility-be/metrics
│   │   │       # GET /v1/admin/system/metrics
│   │   │
│   │   ├── logs/
│   │   │   └── page.tsx  # System logs viewer
│   │   │       # - Real-time logs
│   │   │       # - Search and filter
│   │   │       # BE: utility-be/logs
│   │   │       # GET /v1/admin/system/logs
│   │   │
│   │   ├── feature-flags/
│   │   │   ├── [flagId]/
│   │   │   │   ├── edit/
│   │   │   │   │   └── page.tsx  # Edit feature flag
│   │   │   │   │       # BE: utility-be/feature_flags
│   │   │   │   │       # PUT /v1/admin/feature-flags/{flag_id}
│   │   │   │   ├── rollout/
│   │   │   │   │   └── page.tsx  # Manage rollout
│   │   │   │   │       # BE: utility-be/feature_flags
│   │   │   │   │       # POST /v1/admin/feature-flags/{flag_id}/rollout
│   │   │   │   └── page.tsx  # Feature flag details
│   │   │   │       # BE: utility-be/feature_flags
│   │   │   │       # GET /v1/admin/feature-flags/{flag_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create feature flag
│   │   │   │       # BE: utility-be/feature_flags
│   │   │   │       # POST /v1/admin/feature-flags
│   │   │   └── page.tsx  # Feature flags dashboard
│   │   │       # BE: utility-be/feature_flags
│   │   │       # GET /v1/admin/feature-flags
│   │   │
│   │   ├── experiments/
│   │   │   ├── [experimentId]/
│   │   │   │   ├── edit/
│   │   │   │   │   └── page.tsx  # Edit experiment
│   │   │   │   │       # BE: utility-be/experiments
│   │   │   │   │       # PUT /v1/admin/experiments/{experiment_id}
│   │   │   │   ├── results/
│   │   │   │   │   └── page.tsx  # Experiment results
│   │   │   │   │       # BE: utility-be/experiments
│   │   │   │   │       # GET /v1/admin/experiments/{experiment_id}/results
│   │   │   │   └── page.tsx  # Experiment details
│   │   │   │       # BE: utility-be/experiments
│   │   │   │       # GET /v1/admin/experiments/{experiment_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create A/B experiment
│   │   │   │       # BE: utility-be/experiments
│   │   │   │       # POST /v1/admin/experiments
│   │   │   └── page.tsx  # Experiments dashboard
│   │   │       # BE: utility-be/experiments
│   │   │       # GET /v1/admin/experiments
│   │   │
│   │   └── configuration/
│   │       └── page.tsx  # System configuration
│   │           # - Global settings
│   │           # - Environment variables
│   │           # BE: utility-be/config
│   │           # GET /v1/admin/system/config
│   │           # PUT /v1/admin/system/config
│   │
│   └── audit-logs/
│       ├── [logId]/
│       │   └── page.tsx  # Audit log details
│       │       # BE: utility-be/audit
│       │       # GET /v1/admin/audit-logs/{log_id}
│       └── page.tsx  # Audit logs viewer
│           # - System-wide audit trail
│           # - Filter by entity/action
│           # - Export capabilities
│           # BE: utility-be/audit
│           # GET /v1/admin/audit-logs
```

---

## III. Additional Shared Features (packages/shared/src/features/)

### 1. Events & Real-time Module

```
packages/shared/src/features/
│
├── events/
│   ├── api/
│   │   └── events-api.ts  # Events API client
│   │       # BE: communications-be/events (if exists)
│   ├── hooks/
│   │   ├── useWebSocket.ts  # WebSocket connection
│   │   ├── useRealTimeUpdates.ts  # Real-time data sync
│   │   ├── usePresence.ts  # User presence tracking
│   │   └── useTypingIndicator.ts  # Typing indicators
│   ├── providers/
│   │   └── WebSocketProvider.tsx  # WebSocket context
│   └── types.ts  # Event types
```

### 2. Offline & Sync Module (Mobile Focus)

```
packages/shared/src/features/
│
├── offline/
│   ├── hooks/
│   │   ├── useOfflineQueue.ts  # Offline action queue
│   │   ├── useOfflineSync.ts  # Data synchronization
│   │   ├── useNetworkStatus.ts  # Network status
│   │   └── useOfflineStorage.ts  # Local storage
│   ├── store/
│   │   ├── offline-store.ts  # Offline state management
│   │   └── sync-store.ts  # Sync state management
│   └── types.ts  # Offline types
```

### 3. Performance Monitoring Module

```
packages/shared/src/features/
│
├── performance/
│   ├── hooks/
│   │   ├── usePerformanceMetrics.ts  # Web vitals tracking
│   │   ├── useErrorTracking.ts  # Error monitoring
│   │   └── useAnalyticsTracking.ts  # Analytics events
│   ├── utils/
│   │   ├── metrics-collector.ts  # Metrics collection
│   │   ├── error-reporter.ts  # Error reporting
│   │   └── trace-headers.ts  # Distributed tracing
│   └── types.ts  # Performance types
```

### 4. Experiments & Testing Module

```
packages/shared/src/features/
│
├── experiments/
│   ├── api/
│   │   └── experiments-api.ts  # Experiments API
│   │       # BE: utility-be/experiments
│   ├── hooks/
│   │   ├── useExperiment.ts  # A/B test variant
│   │   ├── useFeatureVariant.ts  # Feature variant
│   │   └── useExperimentTracking.ts  # Track experiment events
│   ├── utils/
│   │   ├── variant-assignment.ts  # Variant logic
│   │   └── experiment-context.ts  # Experiment context
│   └── types.ts  # Experiment types
```

### 5. Geolocation Module

```
packages/shared/src/features/
│
├── geolocation/
│   ├── api/
│   │   └── geolocation-api.ts  # Geolocation API
│   │       # BE: utility-be/geolocation (if exists)
│   ├── hooks/
│   │   ├── useGeolocation.ts  # Device location
│   │   ├── useTimezone.ts  # User timezone
│   │   └── useCountry.ts  # Country detection
│   ├── utils/
│   │   ├── geocoding.ts  # Geocoding utilities
│   │   └── distance-calculator.ts  # Distance calculations
│   └── types.ts  # Geolocation types
```

### 6. Content Moderation Module

```
packages/shared/src/features/
│
├── moderation/
│   ├── api/
│   │   └── moderation-api.ts  # Moderation API
│   │       # BE: admin-be/moderation
│   ├── hooks/
│   │   ├── useReportContent.ts  # Report content
│   │   ├── useContentStatus.ts  # Content status check
│   │   └── useModerationActions.ts  # Moderation actions
│   ├── utils/
│   │   ├── content-validator.ts  # Content validation
│   │   └── profanity-filter.ts  # Profanity filtering (client-side)
│   └── types.ts  # Moderation types
```

### 7. Gamification Module

```
packages/shared/src/features/
│
├── gamification/
│   ├── api/
│   │   ├── achievements-api.ts  # Achievements API
│   │   │   # BE: users-be/achievement
│   │   ├── badges-api.ts  # Badges API
│   │   │   # BE: users-be/badge
│   │   └── leaderboards-api.ts  # Leaderboards API
│   │       # BE: users-be/leaderboard
│   ├── hooks/
│   │   ├── useAchievements.ts  # User achievements
│   │   ├── useBadges.ts  # User badges
│   │   ├── useLeaderboard.ts  # Leaderboard data
│   │   └── usePoints.ts  # Points/reputation
│   ├── queries/
│   │   ├── gamification-mutations.ts  # Gamification mutations
│   │   └── gamification-queries.ts  # Gamification queries
│   └── types.ts  # Gamification types
```

### 8. Referrals & Rewards Module

```
packages/shared/src/features/
│
├── referrals/
│   ├── api/
│   │   └── referrals-api.ts  # Referrals API (enhanced)
│   │       # BE: users-be/referral
│   ├── hooks/
│   │   ├── useReferralProgram.ts  # Referral program details
│   │   ├── useReferralCode.ts  # User referral code
│   │   ├── useReferralStats.ts  # Referral statistics
│   │   └── useRewards.ts  # Rewards management
│   ├── queries/
│   │   ├── referrals-mutations.ts  # Referral mutations
│   │   └── referrals-queries.ts  # Referral queries
│   └── types.ts  # Referral types
```

---

## IV. Additional Mobile Routes (apps/mobile/app/)

### 1. Mobile Onboarding Flow

```
apps/mobile/app/
│
├── onboarding/
│   ├── _layout.tsx  # Onboarding layout
│   ├── welcome.tsx  # Welcome screen
│   ├── features.tsx  # Features showcase
│   ├── permissions.tsx  # Request permissions
│   │   # - Notifications
│   │   # - Camera
│   │   # - Location
│   ├── profile-setup.tsx  # Quick profile setup
│   │   # BE: users-be/profile
│   │   # POST /v1/users/me/profile
│   └── complete.tsx  # Onboarding complete
```

### 2. Mobile Quick Actions

```
apps/mobile/app/
│
├── quick-actions/
│   ├── quick-apply.tsx  # Quick proposal submission
│   │   # - Minimal form
│   │   # - Draft save
│   │   # BE: proposals-be/proposal
│   │   # POST /v1/proposals/quick
│   │
│   ├── quick-message.tsx  # Quick message
│   │   # - Contact selection
│   │   # - Quick send
│   │   # BE: communications-be/message
│   │   # POST /v1/messages
│   │
│   └── quick-time-entry.tsx  # Quick time logging
│       # - Current contract
│       # - Duration input
│       # BE: contracts-be/workdiary
│       # POST /v1/contracts/{contract_id}/workdiary/quick
```

### 3. Mobile Settings

```
apps/mobile/app/
│
├── mobile-settings/
│   ├── app-settings.tsx  # App-specific settings
│   │   # - Cache management
│   │   # - Offline mode
│   │   # - Data usage
│   ├── biometric.tsx  # Biometric authentication
│   │   # - FaceID/TouchID setup
│   │   # BE: users-be/auth
│   │   # POST /v1/auth/biometric/register
│   └── haptics.tsx  # Haptic feedback settings
```

---

## V. Additional UI Components (packages/ui/src/)

### 1. Advanced Form Components

```
packages/ui/src/
│
├── forms/
│   ├── RichTextEditor/
│   │   ├── RichTextEditor.tsx
│   │   ├── RichTextEditor.web.tsx
│   │   └── RichTextEditor.native.tsx
│   ├── CodeEditor/
│   │   ├── CodeEditor.tsx
│   │   ├── CodeEditor.web.tsx
│   │   └── CodeEditor.native.tsx
│   ├── MarkdownEditor/
│   │   ├── MarkdownEditor.tsx
│   │   ├── MarkdownEditor.web.tsx
│   │   └── MarkdownEditor.native.tsx
│   ├── SignatureInput/
│   │   ├── SignatureInput.tsx
│   │   ├── SignatureInput.web.tsx
│   │   └── SignatureInput.native.tsx
│   └── DateRangePicker/
│       ├── DateRangePicker.tsx
│       ├── DateRangePicker.web.tsx
│       └── DateRangePicker.native.tsx
```

### 2. Visualization Components

```
packages/ui/src/
│
├── visualization/
│   ├── Heatmap/
│   │   ├── Heatmap.tsx
│   │   ├── Heatmap.web.tsx
│   │   └── Heatmap.native.tsx
│   ├── Gantt/
│   │   ├── GanttChart.tsx
│   │   ├── GanttChart.web.tsx
│   │   └── GanttChart.native.tsx
│   ├── Kanban/
│   │   ├── KanbanBoard.tsx
│   │   ├── KanbanBoard.web.tsx
│   │   └── KanbanBoard.native.tsx
│   └── OrgChart/
│       ├── OrganizationChart.tsx
│       ├── OrganizationChart.web.tsx
│       └── OrganizationChart.native.tsx
```

### 3. AI & Machine Learning Components

```
packages/ui/src/
│
├── ai/
│   ├── AIAssistant/
│   │   ├── AIAssistant.tsx
│   │   ├── AIAssistant.web.tsx
│   │   └── AIAssistant.native.tsx
│   ├── SmartSuggestions/
│   │   ├── SmartSuggestions.tsx
│   │   ├── SmartSuggestions.web.tsx
│   │   └── SmartSuggestions.native.tsx
│   └── AutoComplete/
│       ├── AIAutoComplete.tsx
│       ├── AIAutoComplete.web.tsx
│       └── AIAutoComplete.native.tsx
```

### 4. Accessibility Components

```
packages/ui/src/
│
├── accessibility/
│   ├── SkipLinks/
│   │   ├── SkipLinks.tsx
│   │   └── SkipLinks.web.tsx
│   ├── ScreenReaderAnnouncer/
│   │   ├── ScreenReaderAnnouncer.tsx
│   │   ├── ScreenReaderAnnouncer.web.tsx
│   │   └── ScreenReaderAnnouncer.native.tsx
│   └── FocusTrap/
│       ├── FocusTrap.tsx
│       ├── FocusTrap.web.tsx
│       └── FocusTrap.native.tsx
```

---

## VI. Testing & Quality Infrastructure

### 1. Test Utilities

```
packages/shared/src/
│
├── testing/
│   ├── test-utils.ts  # Common test utilities
│   ├── mock-data/
│   │   ├── users.ts  # Mock user data
│   │   ├── jobs.ts  # Mock job data
│   │   ├── proposals.ts  # Mock proposal data
│   │   ├── contracts.ts  # Mock contract data
│   │   └── messages.ts  # Mock message data
│   ├── mock-api/
│   │   ├── handlers.ts  # MSW handlers
│   │   └── server.ts  # MSW server setup
│   └── providers/
│       └── TestProviders.tsx  # Test provider wrapper
```

### 2. E2E Test Structure

```
apps/web/
│
├── e2e/
│   ├── auth/
│   │   ├── login.spec.ts
│   │   ├── register.spec.ts
│   │   └── password-reset.spec.ts
│   ├── jobs/
│   │   ├── job-posting.spec.ts
│   │   ├── job-search.spec.ts
│   │   └── job-application.spec.ts
│   ├── proposals/
│   │   ├── proposal-submission.spec.ts
│   │   └── proposal-review.spec.ts
│   ├── contracts/
│   │   ├── contract-creation.spec.ts
│   │   └── contract-execution.spec.ts
│   ├── payments/
│   │   ├── payment-methods.spec.ts
│   │   └── escrow.spec.ts
│   └── fixtures/
│       ├── users.json
│       ├── jobs.json
│       └── contracts.json
```

---

## VII. Documentation Structure

### 1. Architecture Documentation

```
docs/
│
├── architecture/
│   ├── overview.md  # System architecture
│   ├── frontend-architecture.md  # FE architecture details
│   ├── state-management.md  # State management patterns
│   ├── routing.md  # Routing strategy
│   ├── authentication.md  # Auth flow
│   ├── data-fetching.md  # Data fetching patterns
│   └── performance.md  # Performance optimization
│
├── api/
│   ├── introduction.md  # API integration guide
│   ├── authentication.md  # API authentication
│   ├── error-handling.md  # Error handling
│   ├── rate-limiting.md  # Rate limiting
│   └── microservices/
│       ├── users-be.md
│       ├── jobs-be.md
│       ├── proposals-be.md
│       ├── contracts-be.md
│       ├── financial-be.md
│       ├── communications-be.md
│       ├── search-be.md
│       ├── storage-be.md
│       ├── reviews-be.md
│       ├── admin-be.md
│       └── subscriptions-be.md
│
├── guides/
│   ├── getting-started.md
│   ├── development.md
│   ├── testing.md
│   ├── deployment.md
│   ├── contributing.md
│   └── troubleshooting.md
│
└── components/
    ├── component-library.md
    ├── design-tokens.md
    ├── theming.md
    └── accessibility.md
```

---

## VIII. CI/CD & DevOps

### 1. GitHub Actions Workflows

```
.github/
│
├── workflows/
│   ├── ci.yml  # Continuous integration
│   ├── cd-web.yml  # Web deployment
│   ├── cd-mobile.yml  # Mobile deployment
│   ├── lighthouse.yml  # Performance audits
│   ├── accessibility.yml  # A11y checks
│   ├── security.yml  # Security scanning
│   ├── dependency-review.yml  # Dependency checks
│   └── release.yml  # Release automation
│
└── actions/
    ├── setup-node/
    │   └── action.yml
    ├── cache-dependencies/
    │   └── action.yml
    └── deploy-preview/
        └── action.yml
```

### 2. Environment Configuration

```
fe/
│
├── .env.example  # Environment variables template
├── .env.development
├── .env.staging
├── .env.production
└── config/
    ├── environments/
    │   ├── development.ts
    │   ├── staging.ts
    │   └── production.ts
    └── feature-flags.ts
```

---

## IX. Monitoring & Observability

### 1. Monitoring Setup

```
packages/shared/src/
│
├── monitoring/
│   ├── sentry.ts  # Error tracking setup
│   ├── analytics.ts  # Analytics setup
│   ├── web-vitals.ts  # Performance monitoring
│   └── logger.ts  # Logging configuration
```

### 2. Analytics Events

```
packages/shared/src/
│
├── analytics/
│   ├── events/
│   │   ├── auth-events.ts  # Auth analytics
│   │   ├── job-events.ts  # Job analytics
│   │   ├── proposal-events.ts  # Proposal analytics
│   │   ├── contract-events.ts  # Contract analytics
│   │   ├── payment-events.ts  # Payment analytics
│   │   └── user-events.ts  # User behavior
│   ├── trackers/
│   │   ├── page-view-tracker.ts
│   │   ├── interaction-tracker.ts
│   │   └── conversion-tracker.ts
│   └── utils/
│       ├── event-builder.ts
│       └── anonymize.ts  # PII anonymization
```

---

## X. Security Infrastructure

### 1. Security Utilities

```
packages/shared/src/
│
├── security/
│   ├── csrf.ts  # CSRF protection
│   ├── sanitization.ts  # Input sanitization
│   ├── encryption.ts  # Client-side encryption
│   ├── validation.ts  # Input validation
│   └── permissions.ts  # Permission checks
```

### 2. Security Headers

```
apps/web/
│
├── middleware.ts  # Next.js middleware for security headers
└── security/
    ├── headers.ts  # Security headers config
    ├── csp.ts  # Content Security Policy
    └── cors.ts  # CORS configuration
```

---

## Summary

This document completes the **remaining missing folder structure** based on `fe-folder-structure-prompt.md` requirements. The additions cover:

### Dashboard Routes (Enhanced):
1. **Jobs** - Insights, batch operations, scheduling, templates management
2. **Proposals** - Negotiations, comparisons, insights, benchmarking, archives
3. **Contracts** - Change orders, compliance tracking, audit trails, calendar view, benchmarking
4. **Financial** - Reconciliation, forecasting, chargebacks, payment/payout method management, cost centers
5. **Teams** - Spending controls, compliance, performance metrics, hierarchy, integrations
6. **Settings** - Developer tools (API keys, webhooks, OAuth apps), labs, authorized apps, advanced settings

### Admin Panel:
1. **Core Management** - Dashboard, break-glass access, two-person rules
2. **KYC & Verification** - Cases queue, business verification, document review
3. **Financial Operations** - Refunds, goodwill credits, reconciliation, payouts, tax forms, disputes
4. **Moderation** - Reports, actions, appeals, patterns
5. **Communications Ops** - Broadcasts, templates, campaigns, rate limits
6. **Search Quality** - Synonyms, boosts, blacklists, reindex, analytics
7. **System Management** - Health, metrics, logs, feature flags, experiments, configuration, audit logs

### Shared Features:
1. **Events & Real-time** - WebSocket, presence, typing indicators
2. **Offline & Sync** - Queue, sync, network status (mobile focus)
3. **Performance** - Metrics, error tracking, analytics
4. **Experiments** - A/B testing, feature variants
5. **Geolocation** - Location, timezone, country detection
6. **Moderation** - Content reporting, validation
7. **Gamification** - Achievements, badges, leaderboards, points
8. **Referrals & Rewards** - Enhanced referral management

### Mobile Enhancements:
1. **Onboarding** - Complete mobile onboarding flow
2. **Quick Actions** - Quick apply, quick message, quick time entry
3. **Mobile Settings** - App settings, biometric auth, haptics

### UI Components:
1. **Advanced Forms** - Rich text, code editor, markdown, signature, date range
2. **Visualization** - Heatmap, Gantt, Kanban, org chart
3. **AI Components** - AI assistant, smart suggestions, autocomplete
4. **Accessibility** - Skip links, screen reader, focus trap

### Infrastructure:
1. **Testing** - Test utilities, mock data, E2E structure
2. **Documentation** - Architecture, API, guides, components
3. **CI/CD** - GitHub Actions, environment config
4. **Monitoring** - Error tracking, analytics, web vitals
5. **Security** - Security utilities, headers, CSRF, sanitization

All routes and components include proper backend mappings with microservice, domain, HTTP method, and endpoint information as specified in the original requirements.
