## Missing Sections in apps/web/src/app/[locale]/ Structure

### 1. Missing (dashboard) Routes - Jobs Section Completion

```
apps/web/src/app/[locale]/(dashboard)/
│
├── jobs/
│   ├── [jobId]/
│   │   ├── analytics/
│   │   │   └── page.tsx  # Job analytics detail
│   │   │       # - View count trends
│   │   │       # - Proposal funnel
│   │   │       # - Demographic insights
│   │   │       # - Conversion metrics
│   │   │       # BE: jobs-be/analytics
│   │   │       # GET /v1/jobs/{job_id}/analytics
│   │   │
│   │   ├── duplicate/
│   │   │   └── page.tsx  # Duplicate job form
│   │   │       # - Pre-filled from original
│   │   │       # - Modify and post
│   │   │       # BE: jobs-be/job
│   │   │       # GET /v1/jobs/{job_id} (to fetch original)
│   │   │       # POST /v1/jobs (to create duplicate)
│   │   │
│   │   ├── history/
│   │   │   └── page.tsx  # Job edit history
│   │   │       # - Version timeline
│   │   │       # - Change diff viewer
│   │   │       # - Revert capability
│   │   │       # BE: jobs-be/audit
│   │   │       # GET /v1/jobs/{job_id}/history
│   │   │
│   │   └── repost/
│   │       └── page.tsx  # Repost expired job
│   │           # - Review and update
│   │           # - Set new deadline
│   │           # BE: jobs-be/job
│   │           # POST /v1/jobs/{job_id}/repost
│   │
│   ├── archived/
│   │   └── page.tsx  # Archived jobs list
│   │       # - View archived jobs
│   │       # - Restore capability
│   │       # BE: jobs-be/job
│   │       # GET /v1/jobs/archived
│   │
│   ├── templates/
│   │   ├── [templateId]/
│   │   │   ├── edit/
│   │   │   │   └── page.tsx  # Edit job template
│   │   │   │       # BE: jobs-be/template
│   │   │   │       # PUT /v1/jobs/templates/{template_id}
│   │   │   └── use/
│   │   │       └── page.tsx  # Use template to create job
│   │   │           # BE: jobs-be/template, jobs-be/job
│   │   │           # GET /v1/jobs/templates/{template_id}
│   │   │           # POST /v1/jobs (using template data)
│   │   └── page.tsx  # Job templates list
│   │       # - Saved templates
│   │       # - Create new template
│   │       # BE: jobs-be/template
│   │       # GET /v1/jobs/templates
│   │
│   └── saved-searches/
│       └── page.tsx  # Saved job searches
│           # - Manage saved searches
│           # - Email alerts toggle
│           # BE: search-be/saved-search
│           # GET /v1/search/saved-searches?type=jobs
```

### 2. Missing (dashboard) Routes - Proposals Section Completion

```
apps/web/src/app/[locale]/(dashboard)/
│
├── proposals/
│   ├── [proposalId]/
│   │   ├── analytics/
│   │   │   └── page.tsx  # Proposal performance analytics
│   │   │       # - View metrics
│   │   │       # - Comparison with others
│   │   │       # BE: proposals-be/analytics
│   │   │       # GET /v1/proposals/{proposal_id}/analytics
│   │   │
│   │   ├── revise/
│   │   │   └── page.tsx  # Revise proposal
│   │   │       # - Update terms
│   │   │       # - Adjust bid
│   │   │       # - Add clarifications
│   │   │       # BE: proposals-be/proposal
│   │   │       # PUT /v1/proposals/{proposal_id}/revise
│   │   │
│   │   └── versions/
│   │       └── page.tsx  # Proposal revision history
│   │           # - Version timeline
│   │           # - Change tracking
│   │           # BE: proposals-be/audit
│   │           # GET /v1/proposals/{proposal_id}/versions
│   │
│   ├── drafts/
│   │   ├── [draftId]/
│   │   │   └── edit/
│   │   │       └── page.tsx  # Edit proposal draft
│   │   │           # BE: proposals-be/draft
│   │   │           # PUT /v1/proposals/drafts/{draft_id}
│   │   └── page.tsx  # Proposal drafts list
│   │       # BE: proposals-be/draft
│   │       # GET /v1/proposals/drafts
│   │
│   ├── declined/
│   │   └── page.tsx  # Declined proposals
│   │       # - Feedback from client
│   │       # - Learn from rejections
│   │       # BE: proposals-be/proposal
│   │       # GET /v1/proposals?status=declined
│   │
│   ├── withdrawn/
│   │   └── page.tsx  # Withdrawn proposals
│   │       # - Self-withdrawn proposals
│   │       # - Reasons tracking
│   │       # BE: proposals-be/proposal
│   │       # GET /v1/proposals?status=withdrawn
│   │
│   └── templates/
│       ├── [templateId]/
│       │   ├── edit/
│       │   │   └── page.tsx  # Edit proposal template
│       │   │       # BE: proposals-be/template
│       │   │       # PUT /v1/proposals/templates/{template_id}
│       │   └── use/
│       │       └── page.tsx  # Use template for new proposal
│       │           # BE: proposals-be/template
│       │           # GET /v1/proposals/templates/{template_id}
│       └── page.tsx  # Proposal templates list
│           # - Manage templates
│           # - Create template from proposal
│           # BE: proposals-be/template
│           # GET /v1/proposals/templates
```

### 3. Missing (dashboard) Routes - Contracts Section Completion

```
apps/web/src/app/[locale]/(dashboard)/
│
├── contracts/
│   ├── [contractId]/
│   │   ├── amendments/
│   │   │   ├── [amendmentId]/
│   │   │   │   └── page.tsx  # Amendment detail
│   │   │   │       # - View changes
│   │   │   │       # - Approval status
│   │   │   │       # BE: contracts-be/amendment
│   │   │   │       # GET /v1/contracts/{contract_id}/amendments/{amendment_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create amendment
│   │   │   │       # - Modify terms
│   │   │   │       # - Add milestones
│   │   │   │       # - Change scope
│   │   │   │       # BE: contracts-be/amendment
│   │   │   │       # POST /v1/contracts/{contract_id}/amendments
│   │   │   └── page.tsx  # Amendments list
│   │   │       # BE: contracts-be/amendment
│   │   │       # GET /v1/contracts/{contract_id}/amendments
│   │   │
│   │   ├── deliverables/
│   │   │   ├── [deliverableId]/
│   │   │   │   ├── accept/
│   │   │   │   │   └── page.tsx  # Accept deliverable
│   │   │   │   │       # - Review checklist
│   │   │   │   │       # - Release milestone payment
│   │   │   │   │       # BE: contracts-be/deliverable
│   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/accept
│   │   │   │   ├── reject/
│   │   │   │   │   └── page.tsx  # Reject deliverable
│   │   │   │   │       # - Provide feedback
│   │   │   │   │       # - Request revisions
│   │   │   │   │       # BE: contracts-be/deliverable
│   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/reject
│   │   │   │   └── page.tsx  # Deliverable detail
│   │   │   │       # - Files and description
│   │   │   │       # - Review history
│   │   │   │       # BE: contracts-be/deliverable
│   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}
│   │   │   ├── submit/
│   │   │   │   └── page.tsx  # Submit deliverable
│   │   │   │       # - Upload files
│   │   │   │       # - Add notes
│   │   │   │       # BE: contracts-be/deliverable, storage-be/asset
│   │   │   │       # POST /v1/contracts/{contract_id}/deliverables
│   │   │   └── page.tsx  # Deliverables list
│   │   │       # BE: contracts-be/deliverable
│   │   │       # GET /v1/contracts/{contract_id}/deliverables
│   │   │
│   │   ├── extensions/
│   │   │   ├── [extensionId]/
│   │   │   │   └── page.tsx  # Extension request detail
│   │   │   │       # - Approve/reject
│   │   │   │       # - Negotiation
│   │   │   │       # BE: contracts-be/extension
│   │   │   │       # GET /v1/contracts/{contract_id}/extensions/{extension_id}
│   │   │   ├── request/
│   │   │   │   └── page.tsx  # Request extension
│   │   │   │       # - New deadline
│   │   │   │       # - Justification
│   │   │   │       # BE: contracts-be/extension
│   │   │   │       # POST /v1/contracts/{contract_id}/extensions
│   │   │   └── page.tsx  # Extensions list
│   │   │       # BE: contracts-be/extension
│   │   │       # GET /v1/contracts/{contract_id}/extensions
│   │   │
│   │   ├── invoices/
│   │   │   ├── [invoiceId]/
│   │   │   │   └── page.tsx  # Invoice detail
│   │   │   │       # - Line items
│   │   │   │       # - Payment status
│   │   │   │       # - Download PDF
│   │   │   │       # BE: financial-be/invoice
│   │   │   │       # GET /v1/invoices/{invoice_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create invoice
│   │   │   │       # - Add line items
│   │   │   │       # - Tax calculation
│   │   │   │       # BE: financial-be/invoice
│   │   │   │       # POST /v1/contracts/{contract_id}/invoices
│   │   │   └── page.tsx  # Invoices list
│   │   │       # BE: financial-be/invoice
│   │   │       # GET /v1/contracts/{contract_id}/invoices
│   │   │
│   │   ├── pause/
│   │   │   └── page.tsx  # Pause contract
│   │   │       # - Reason selection
│   │   │       # - Estimated resume date
│   │   │       # BE: contracts-be/contract
│   │   │       # POST /v1/contracts/{contract_id}/pause
│   │   │
│   │   ├── resume/
│   │   │   └── page.tsx  # Resume paused contract
│   │   │       # - Confirm resume
│   │   │       # - Adjust timeline
│   │   │       # BE: contracts-be/contract
│   │   │       # POST /v1/contracts/{contract_id}/resume
│   │   │
│   │   ├── sow/
│   │   │   ├── edit/
│   │   │   │   └── page.tsx  # Edit SOW (before signing)
│   │   │   │       # BE: contracts-be/sow
│   │   │   │       # PUT /v1/contracts/{contract_id}/sow
│   │   │   └── page.tsx  # SOW detail view
│   │   │       # - Scope of work
│   │   │       # - Deliverables
│   │   │       # - Timeline
│   │   │       # BE: contracts-be/sow
│   │   │       # GET /v1/contracts/{contract_id}/sow
│   │   │
│   │   ├── terminate/
│   │   │   └── page.tsx  # Terminate contract
│   │   │       # - Termination reason
│   │   │       # - Final settlement
│   │   │       # BE: contracts-be/contract
│   │   │       # POST /v1/contracts/{contract_id}/terminate
│   │   │
│   │   └── workdiary/
│   │       ├── manual-time/
│   │       │   └── page.tsx  # Add manual time entry
│   │       │       # - Date/time range
│   │       │       # - Description
│   │       │       # - Approval required
│   │       │       # BE: contracts-be/workdiary
│   │       │       # POST /v1/contracts/{contract_id}/workdiary/manual-time
│   │       └── reports/
│   │           └── page.tsx  # Work diary reports
│   │               # - Weekly summaries
│   │               # - Time breakdowns
│   │               # BE: contracts-be/workdiary
│   │               # GET /v1/contracts/{contract_id}/workdiary/reports
│   │
│   ├── archived/
│   │   └── page.tsx  # Archived contracts
│   │       # - Completed contracts
│   │       # - Historical data
│   │       # BE: contracts-be/contract
│   │       # GET /v1/contracts?status=archived
│   │
│   ├── paused/
│   │   └── page.tsx  # Paused contracts list
│   │       # - Temporarily paused
│   │       # - Resume options
│   │       # BE: contracts-be/contract
│   │       # GET /v1/contracts?status=paused
│   │
│   └── templates/
│       ├── [templateId]/
│       │   ├── edit/
│       │   │   └── page.tsx  # Edit contract template
│       │   │       # BE: contracts-be/template
│       │   │       # PUT /v1/contracts/templates/{template_id}
│       │   └── use/
│       │       └── page.tsx  # Use template for new contract
│       │           # BE: contracts-be/template
│       │           # GET /v1/contracts/templates/{template_id}
│       └── page.tsx  # Contract templates list
│           # - Standard contract templates
│           # - Create from existing
│           # BE: contracts-be/template
│           # GET /v1/contracts/templates
```

### 4. Missing (dashboard) Routes - Financial Section Completion

```
apps/web/src/app/[locale]/(dashboard)/
│
├── financial/
│   ├── billing/
│   │   ├── history/
│   │   │   └── page.tsx  # Billing history
│   │   │       # - Past invoices
│   │   │       # - Payment receipts
│   │   │       # BE: financial-be/billing
│   │   │       # GET /v1/billing/history
│   │   │
│   │   ├── payment-methods/
│   │   │   ├── [methodId]/
│   │   │   │   ├── edit/
│   │   │   │   │   └── page.tsx  # Edit payment method
│   │   │   │   │       # BE: financial-be/payment-method
│   │   │   │   │       # PUT /v1/payment-methods/{method_id}
│   │   │   │   └── page.tsx  # Payment method detail
│   │   │   │       # BE: financial-be/payment-method
│   │   │   │       # GET /v1/payment-methods/{method_id}
│   │   │   ├── add/
│   │   │   │   └── page.tsx  # Add payment method
│   │   │   │       # - Card/bank details
│   │   │   │       # - Verification
│   │   │   │       # BE: financial-be/payment-method
│   │   │   │       # POST /v1/payment-methods
│   │   │   └── page.tsx  # Payment methods list
│   │   │       # - Manage cards/banks
│   │   │       # - Set default
│   │   │       # BE: financial-be/payment-method
│   │   │       # GET /v1/payment-methods
│   │   │
│   │   └── subscriptions/
│   │       └── page.tsx  # Subscription billing
│   │           # - Current plan
│   │           # - Billing cycle
│   │           # - Upgrade/downgrade
│   │           # BE: financial-be/subscription
│   │           # GET /v1/subscriptions/billing
│   │
│   ├── disputes/
│   │   ├── [disputeId]/
│   │   │   ├── evidence/
│   │   │   │   ├── submit/
│   │   │   │   │   └── page.tsx  # Submit evidence
│   │   │   │   │       # - Upload documents
│   │   │   │   │       # - Add description
│   │   │   │   │       # BE: contracts-be/dispute, storage-be/asset
│   │   │   │   │       # POST /v1/disputes/{dispute_id}/evidence
│   │   │   │   └── page.tsx  # Evidence list
│   │   │   │       # BE: contracts-be/dispute
│   │   │   │       # GET /v1/disputes/{dispute_id}/evidence
│   │   │   ├── messages/
│   │   │   │   └── page.tsx  # Dispute messages
│   │   │   │       # - Communication thread
│   │   │   │       # - Mediation chat
│   │   │   │       # BE: communications-be/conversation
│   │   │   │       # GET /v1/disputes/{dispute_id}/messages
│   │   │   ├── resolution/
│   │   │   │   └── page.tsx  # Dispute resolution
│   │   │   │       # - Accept/reject resolution
│   │   │   │       # - Final settlement
│   │   │   │       # BE: contracts-be/dispute
│   │   │   │       # POST /v1/disputes/{dispute_id}/resolution
│   │   │   └── page.tsx  # Dispute detail
│   │   │       # - Status and timeline
│   │   │       # - Evidence
│   │   │       # - Messages
│   │   │       # BE: contracts-be/dispute
│   │   │       # GET /v1/disputes/{dispute_id}
│   │   ├── create/
│   │   │   └── page.tsx  # Create dispute
│   │   │       # - Select contract
│   │   │       # - Issue description
│   │   │       # - Initial evidence
│   │   │       # BE: contracts-be/dispute
│   │   │       # POST /v1/disputes
│   │   └── page.tsx  # Disputes list
│   │       # - Open disputes
│   │       # - Resolved disputes
│   │       # BE: contracts-be/dispute
│   │       # GET /v1/disputes
│   │
│   ├── escrow/
│   │   ├── [escrowId]/
│   │   │   ├── fund/
│   │   │   │   └── page.tsx  # Fund escrow
│   │   │   │       # - Select payment method
│   │   │   │       # - Amount confirmation
│   │   │   │       # BE: financial-be/escrow
│   │   │   │       # POST /v1/escrow/{escrow_id}/fund
│   │   │   ├── release/
│   │   │   │   └── page.tsx  # Release escrow funds
│   │   │   │       # - Release to freelancer
│   │   │   │       # - Partial/full release
│   │   │   │       # BE: financial-be/escrow
│   │   │   │       # POST /v1/escrow/{escrow_id}/release
│   │   │   └── page.tsx  # Escrow detail
│   │   │       # - Balance and transactions
│   │   │       # - Release history
│   │   │       # BE: financial-be/escrow
│   │   │       # GET /v1/escrow/{escrow_id}
│   │   └── page.tsx  # Escrow accounts list
│   │       # - Active escrows
│   │       # - Transaction history
│   │       # BE: financial-be/escrow
│   │       # GET /v1/escrow
│   │
│   ├── invoices/
│   │   ├── [invoiceId]/
│   │   │   ├── download/
│   │   │   │   └── page.tsx  # Download invoice PDF
│   │   │   │       # BE: financial-be/invoice
│   │   │   │       # GET /v1/invoices/{invoice_id}/download
│   │   │   ├── pay/
│   │   │   │   └── page.tsx  # Pay invoice
│   │   │   │       # - Select payment method
│   │   │   │       # - Confirm payment
│   │   │   │       # BE: financial-be/payment
│   │   │   │       # POST /v1/invoices/{invoice_id}/pay
│   │   │   └── page.tsx  # Invoice detail (already in combined, but completion)
│   │   │
│   │   └── overdue/
│   │       └── page.tsx  # Overdue invoices
│   │           # - Payment reminders
│   │           # - Late fees
│   │           # BE: financial-be/invoice
│   │           # GET /v1/invoices?status=overdue
│   │
│   ├── payouts/
│   │   ├── [payoutId]/
│   │   │   ├── details/
│   │   │   │   └── page.tsx  # Payout transaction detail
│   │   │   │       # - Transaction breakdown
│   │   │   │       # - Tax withholdings
│   │   │   │       # BE: financial-be/payout
│   │   │   │       # GET /v1/payouts/{payout_id}
│   │   │   └── receipt/
│   │   │       └── page.tsx  # Payout receipt
│   │   │           # - Download receipt
│   │   │           # - Tax information
│   │   │           # BE: financial-be/payout
│   │   │           # GET /v1/payouts/{payout_id}/receipt
│   │   ├── pending/
│   │   │   └── page.tsx  # Pending payouts
│   │   │       # - In-process withdrawals
│   │   │       # - Estimated dates
│   │   │       # BE: financial-be/payout
│   │   │       # GET /v1/payouts?status=pending
│   │   │
│   │   └── schedule/
│   │       └── page.tsx  # Schedule payout
│   │           # - Set payout frequency
│   │           # - Minimum threshold
│   │           # BE: financial-be/payout
│   │           # POST /v1/payouts/schedule
│   │
│   ├── refunds/
│   │   ├── [refundId]/
│   │   │   └── page.tsx  # Refund detail
│   │   │       # - Refund status
│   │   │       # - Processing timeline
│   │   │       # BE: financial-be/refund
│   │   │       # GET /v1/refunds/{refund_id}
│   │   ├── request/
│   │   │   └── page.tsx  # Request refund
│   │   │       # - Select transaction
│   │   │       # - Reason and evidence
│   │   │       # BE: financial-be/refund
│   │   │       # POST /v1/refunds
│   │   └── page.tsx  # Refunds list
│   │       # - Requested refunds
│   │       # - Completed refunds
│   │       # BE: financial-be/refund
│   │       # GET /v1/refunds
│   │
│   ├── reports/
│   │   ├── earnings/
│   │   │   └── page.tsx  # Earnings reports
│   │   │       # - Period selection
│   │   │       # - Breakdown by project
│   │   │       # - Export options
│   │   │       # BE: financial-be/reports
│   │   │       # GET /v1/reports/earnings
│   │   │
│   │   ├── expenses/
│   │   │   └── page.tsx  # Expense reports
│   │   │       # - Platform fees
│   │   │       # - Service charges
│   │   │       # BE: financial-be/reports
│   │   │       # GET /v1/reports/expenses
│   │   │
│   │   ├── tax/
│   │   │   ├── 1099/
│   │   │   │   └── page.tsx  # 1099 tax forms
│   │   │   │       # - Annual 1099s
│   │   │   │       # - Download PDFs
│   │   │   │       # BE: financial-be/tax
│   │   │   │       # GET /v1/tax/forms/1099
│   │   │   ├── vat/
│   │   │   │   └── page.tsx  # VAT reports
│   │   │   │       # - VAT breakdown
│   │   │   │       # - Export for filing
│   │   │   │       # BE: financial-be/tax
│   │   │   │       # GET /v1/tax/reports/vat
│   │   │   └── page.tsx  # Tax reports overview
│   │   │       # BE: financial-be/tax
│   │   │       # GET /v1/tax/reports
│   │   │
│   │   └── page.tsx  # Financial reports overview
│   │       # - Quick stats
│   │       # - Report categories
│   │       # BE: financial-be/reports
│   │       # GET /v1/reports
│   │
│   ├── tax/
│   │   ├── forms/
│   │   │   ├── w9/
│   │   │   │   └── page.tsx  # W-9 form management
│   │   │   │       # - Submit W-9
│   │   │   │       # - Update information
│   │   │   │       # BE: financial-be/tax
│   │   │   │       # POST /v1/tax/forms/w9
│   │   │   └── page.tsx  # Tax forms list
│   │   │       # BE: financial-be/tax
│   │   │       # GET /v1/tax/forms
│   │   │
│   │   ├── settings/
│   │   │   └── page.tsx  # Tax settings
│   │   │       # - Tax residency
│   │   │       # - Withholding preferences
│   │   │       # BE: financial-be/tax
│   │   │       # PUT /v1/tax/settings
│   │   │
│   │   └── page.tsx  # Tax overview
│   │       # - Tax liability
│   │       # - Forms required
│   │       # BE: financial-be/tax
│   │       # GET /v1/tax/overview
│   │
│   └── wallet/
│       ├── add-funds/
│       │   └── page.tsx  # Add funds to wallet
│       │       # - Amount input
│       │       # - Payment method selection
│       │       # BE: financial-be/wallet
│       │       # POST /v1/wallet/add-funds
│       │
│       └── withdraw/
│           └── page.tsx  # Withdraw from wallet
│               # - Withdrawal amount
│               # - Destination account
│               # BE: financial-be/wallet, financial-be/payout
│               # POST /v1/wallet/withdraw
```

### 5. Missing (dashboard) Routes - Messages/Communications Section

```
apps/web/src/app/[locale]/(dashboard)/
│
├── messages/
│   ├── archived/
│   │   └── page.tsx  # Archived conversations
│   │       # - View archived messages
│   │       # - Unarchive capability
│   │       # BE: communications-be/conversation
│   │       # GET /v1/conversations?archived=true
│   │
│   ├── compose/
│   │   └── page.tsx  # Compose new message
│   │       # - Recipient search
│   │       # - Subject and body
│   │       # - Attachments
│   │       # BE: communications-be/conversation
│   │       # POST /v1/conversations
│   │
│   ├── drafts/
│   │   └── page.tsx  # Message drafts
│   │       # - Saved drafts
│   │       # - Resume editing
│   │       # BE: communications-be/draft
│   │       # GET /v1/messages/drafts
│   │
│   └── starred/
│       └── page.tsx  # Starred messages
│           # - Important messages
│           # - Quick access
│           # BE: communications-be/conversation
│           # GET /v1/conversations?starred=true
```

### 6. Missing (dashboard) Routes - Notifications Section

```
apps/web/src/app/[locale]/(dashboard)/
│
├── notifications/
│   ├── all/
│   │   └── page.tsx  # All notifications
│   │       # - Complete history
│   │       # - Filter options
│   │       # BE: communications-be/notification
│   │       # GET /v1/notifications
│   │
│   ├── unread/
│   │   └── page.tsx  # Unread notifications only
│   │       # BE: communications-be/notification
│   │       # GET /v1/notifications?read=false
│   │
│   └── settings/
│       ├── channels/
│       │   └── page.tsx  # Notification channels
│       │       # - Email settings
│       │       # - Push notification settings
│       │       # - SMS settings
│       │       # BE: communications-be/preferences
│       │       # GET /v1/notifications/channels
│       │       # PUT /v1/notifications/channels
│       │
│       └── page.tsx  # Notification preferences (in combined already, ensuring completion)
```

### 7. Missing (dashboard) Routes - Portfolio Section Completion

```
apps/web/src/app/[locale]/(dashboard)/
│
├── portfolio/
│   ├── analytics/
│   │   └── page.tsx  # Portfolio analytics
│   │       # - View count trends
│   │       # - Engagement metrics
│   │       # - Profile strength score
│   │       # BE: users-be/analytics
│   │       # GET /v1/users/me/portfolio/analytics
│   │
│   ├── import/
│   │   └── page.tsx  # Import portfolio
│   │       # - LinkedIn import
│   │       # - Behance import
│   │       # - GitHub import
│   │       # BE: users-be/portfolio, storage-be/asset
│   │       # POST /v1/users/me/portfolio/import
│   │
│   └── templates/
│       └── page.tsx  # Portfolio templates
│           # - Pre-designed layouts
│           # - Customize template
│           # BE: users-be/portfolio
│           # GET /v1/portfolio/templates
```

### 8. Missing (dashboard) Routes - Profile Section Completion

```
apps/web/src/app/[locale]/(dashboard)/
│
├── profile/
│   ├── certifications/
│   │   ├── [certId]/
│   │   │   ├── edit/
│   │   │   │   └── page.tsx  # Edit certification
│   │   │   │       # BE: users-be/certifications
│   │   │   │       # PUT /v1/users/me/certifications/{cert_id}
│   │   │   └── page.tsx  # Certification detail
│   │   │       # BE: users-be/certifications
│   │   │       # GET /v1/users/me/certifications/{cert_id}
│   │   ├── add/
│   │   │   └── page.tsx  # Add certification
│   │   │       # - Certification name
│   │   │       # - Issuing organization
│   │   │       # - Credential URL/ID
│   │   │       # - Upload certificate
│   │   │       # BE: users-be/certifications, storage-be/asset
│   │   │       # POST /v1/users/me/certifications
│   │   └── verify/
│   │       └── page.tsx  # Verify certifications
│   │           # - Verification requests
│   │           # - Badge display
│   │           # BE: users-be/certifications
│   │           # POST /v1/users/me/certifications/{cert_id}/verify
│   │
│   ├── languages/
│   │   └── page.tsx  # Language proficiency
│   │       # - Add languages
│   │       # - Proficiency levels
│   │       # BE: users-be/profile
│   │       # PUT /v1/users/me/languages
│   │
│   ├── portfolio-items/
│   │   └── [itemId]/
│   │       ├── analytics/
│   │       │   └── page.tsx  # Portfolio item analytics
│   │       │       # - View count
│   │       │       # - Engagement rate
│   │       │       # BE: users-be/analytics
│   │       │       # GET /v1/users/me/portfolio/{item_id}/analytics
│   │       └── page.tsx  # Portfolio item detail (in combined, ensuring it exists)
│   │
│   ├── references/
│   │   ├── [referenceId]/
│   │   │   └── page.tsx  # Reference detail
│   │   │       # - Reference content
│   │   │       # - Relationship details
│   │   │       # BE: users-be/references
│   │   │       # GET /v1/users/me/references/{reference_id}
│   │   ├── request/
│   │   │   └── page.tsx  # Request reference
│   │   │       # - Select contact
│   │   │       # - Reference type
│   │   │       # BE: users-be/references
│   │   │       # POST /v1/users/me/references/request
│   │   └── page.tsx  # References list
│   │       # - Manage references
│   │       # - Request new
│   │       # BE: users-be/references
│   │       # GET /v1/users/me/references
│   │
│   ├── social-links/
│   │   └── page.tsx  # Social media links
│   │       # - LinkedIn, Twitter, GitHub
│   │       # - Portfolio websites
│   │       # BE: users-be/profile
│   │       # PUT /v1/users/me/social-links
│   │
│   └── visibility/
│       └── page.tsx  # Profile visibility settings
│           # - Public/private toggle
│           # - Search visibility
│           # - Profile sections visibility
│           # BE: users-be/profile
│           # PUT /v1/users/me/visibility
```

### 9. Missing (dashboard) Routes - Reviews Section Completion

```
apps/web/src/app/[locale]/(dashboard)/
│
├── reviews/
│   ├── dispute/
│   │   ├── [disputeId]/
│   │   │   └── page.tsx  # Review dispute detail
│   │   │       # - Dispute reason
│   │   │       # - Resolution status
│   │   │       # BE: reviews-be/dispute
│   │   │       # GET /v1/reviews/disputes/{dispute_id}
│   │   ├── submit/
│   │   │   └── page.tsx  # Dispute a review
│   │   │       # - Select review
│   │   │       # - Dispute reason
│   │   │       # - Evidence
│   │   │       # BE: reviews-be/dispute
│   │   │       # POST /v1/reviews/{review_id}/dispute
│   │   └── page.tsx  # Review disputes list
│   │       # BE: reviews-be/dispute
│   │       # GET /v1/reviews/disputes
│   │
│   ├── given/
│   │   └── page.tsx  # Reviews given by user
│   │       # - All reviews posted
│   │       # - Edit capability (time-limited)
│   │       # BE: reviews-be/review
│   │       # GET /v1/reviews/given
│   │
│   ├── pending/
│   │   └── page.tsx  # Pending reviews to write
│   │       # - Completed contracts
│   │       # - Review prompts
│   │       # BE: reviews-be/review
│   │       # GET /v1/reviews/pending
│   │
│   └── received/
│       └── page.tsx  # Reviews received
│           # - All reviews about user
│           # - Response capability
│           # BE: reviews-be/review
│           # GET /v1/reviews/received
```

### 10. Missing (dashboard) Routes - Search Section Completion

```
apps/web/src/app/[locale]/(dashboard)/
│
├── search/
│   ├── freelancers/
│   │   ├── advanced/
│   │   │   └── page.tsx  # Advanced freelancer search
│   │   │       # - Multiple filters
│   │   │       # - Boolean operators
│   │   │       # - Saved search option
│   │   │       # BE: search-be/query
│   │   │       # POST /v1/search/freelancers/advanced
│   │   └── page.tsx  # Basic freelancer search (in combined, ensuring it exists)
│   │
│   ├── jobs/
│   │   ├── advanced/
│   │   │   └── page.tsx  # Advanced job search
│   │   │       # - Complex filters
│   │   │       # - Save search
│   │   │       # BE: search-be/query
│   │   │       # POST /v1/search/jobs/advanced
│   │   └── page.tsx  # Basic job search (in combined)
│   │
│   └── portfolios/
│       └── page.tsx  # Search portfolios
│           # - Search by category
│           # - Filter by tags
│           # BE: search-be/query
│           # GET /v1/search/portfolios
```

### 11. Missing (dashboard) Routes - Settings Section Completion

```
apps/web/src/app/[locale]/(dashboard)/
│
├── settings/
│   ├── account/
│   │   ├── close/
│   │   │   └── page.tsx  # Close account
│   │   │       # - Confirmation steps
│   │   │       # - Data retention options
│   │   │       # - Final warning
│   │   │       # BE: users-be/account
│   │   │       # POST /v1/users/me/close-account
│   │   │
│   │   ├── data-export/
│   │   │   └── page.tsx  # Export user data (GDPR)
│   │   │       # - Request export
│   │   │       # - Download data
│   │   │       # BE: users-be/account
│   │   │       # POST /v1/users/me/data-export
│   │   │
│   │   ├── deactivate/
│   │   │   └── page.tsx  # Deactivate account
│   │   │       # - Temporary suspension
│   │   │       # - Reactivation info
│   │   │       # BE: users-be/account
│   │   │       # POST /v1/users/me/deactivate
│   │   │
│   │   └── reactivate/
│   │       └── page.tsx  # Reactivate account
│   │           # - Restore account
│   │           # - Update information
│   │           # BE: users-be/account
│   │           # POST /v1/users/me/reactivate
│   │
│   ├── accessibility/
│   │   └── page.tsx  # Accessibility preferences
│   │       # - Screen reader settings
│   │       # - Keyboard shortcuts
│   │       # - High contrast mode
│   │       # - Motion reduction
│   │       # BE: users-be/preferences
│   │       # PUT /v1/users/me/accessibility
│   │
│   ├── api/
│   │   ├── documentation/
│   │   │   └── page.tsx  # API documentation
│   │   │       # - API reference
│   │   │       # - Examples
│   │   │       # BE: None (static content)
│   │   │
│   │   └── page.tsx  # API settings (in combined, ensuring developer section exists)
│   │
│   ├── blocked-users/
│   │   └── page.tsx  # Blocked users management
│   │       # - View blocked users
│   │       # - Unblock users
│   │       # BE: users-be/blocked
│   │       # GET /v1/users/me/blocked-users
│   │       # DELETE /v1/users/me/blocked-users/{user_id}
│   │
│   ├── devices/
│   │   └── page.tsx  # Connected devices
│   │       # - Active sessions
│   │       # - Device trust management
│   │       # - Logout devices
│   │       # BE: users-be/sessions
│   │       # GET /v1/users/me/devices
│   │       # DELETE /v1/users/me/devices/{device_id}
│   │
│   ├── integrations/
│   │   ├── connected/
│   │   │   └── page.tsx  # Connected integrations
│   │   │       # - Active integrations
│   │   │       # - Disconnect options
│   │   │       # BE: users-be/integrations
│   │   │       # GET /v1/users/me/integrations
│   │   │
│   │   └── available/
│   │       └── page.tsx  # Available integrations
│   │           # - Browse integrations
│   │           # - Connect new
│   │           # BE: users-be/integrations
│   │           # GET /v1/integrations/available
│   │
│   ├── login-history/
│   │   └── page.tsx  # Login history
│   │       # - Recent logins
│   │       # - Location tracking
│   │       # - Security alerts
│   │       # BE: users-be/audit
│   │       # GET /v1/users/me/login-history
│   │
│   ├── preferences/
│   │   ├── appearance/
│   │   │   └── page.tsx  # Appearance preferences
│   │   │       # - Theme selection
│   │   │       # - Color customization
│   │   │       # - Layout preferences
│   │   │       # BE: users-be/preferences
│   │   │       # PUT /v1/users/me/preferences/appearance
│   │   │
│   │   ├── language/
│   │   │   └── page.tsx  # Language preferences
│   │   │       # - Interface language
│   │   │       # - Content language
│   │   │       # BE: users-be/preferences
│   │   │       # PUT /v1/users/me/preferences/language
│   │   │
│   │   └── timezone/
│   │       └── page.tsx  # Timezone settings
│   │           # - Timezone selection
│   │           # - Time format (12/24 hour)
│   │           # BE: users-be/preferences
│   │           # PUT /v1/users/me/preferences/timezone
│   │
│   ├── privacy/
│   │   ├── activity/
│   │   │   └── page.tsx  # Activity privacy
│   │   │       # - Who can see activity
│   │   │       # - Activity history settings
│   │   │       # BE: users-be/privacy
│   │   │       # PUT /v1/users/me/privacy/activity
│   │   │
│   │   ├── data-sharing/
│   │   │   └── page.tsx  # Data sharing preferences
│   │   │       # - Analytics opt-in/out
│   │   │       # - Third-party sharing
│   │   │       # BE: users-be/privacy
│   │   │       # PUT /v1/users/me/privacy/data-sharing
│   │   │
│   │   └── profile-visibility/
│   │       └── page.tsx  # Profile visibility settings
│   │           # - Search engine visibility
│   │           # - Profile sections
│   │           # BE: users-be/privacy
│   │           # PUT /v1/users/me/privacy/profile-visibility
│   │
│   └── two-factor/
│       ├── backup-codes/
│       │   └── page.tsx  # Backup codes
│       │       # - Generate codes
│       │       # - Download codes
│       │       # BE: users-be/mfa
│       │       # POST /v1/users/me/mfa/backup-codes
│       │
│       ├── disable/
│       │   └── page.tsx  # Disable 2FA
│       │       # - Verification required
│       │       # - Confirmation
│       │       # BE: users-be/mfa
│       │       # POST /v1/users/me/mfa/disable
│       │
│       └── methods/
│           ├── app/
│           │   └── page.tsx  # Authenticator app setup
│           │       # - QR code
│           │       # - Verification
│           │       # BE: users-be/mfa
│           │       # POST /v1/users/me/mfa/app
│           │
│           └── sms/
│               └── page.tsx  # SMS 2FA setup
│                   # - Phone verification
│                   # - Test code
│                   # BE: users-be/mfa
│                   # POST /v1/users/me/mfa/sms
```

### 12. Missing (dashboard) Routes - Subscriptions Section

```
apps/web/src/app/[locale]/(dashboard)/
│
├── subscriptions/
│   ├── billing/
│   │   └── page.tsx  # Subscription billing details
│   │       # - Current charges
│   │       # - Next billing date
│   │       # - Payment history
│   │       # BE: financial-be/subscription
│   │       # GET /v1/subscriptions/billing
│   │
│   ├── cancel/
│   │   └── page.tsx  # Cancel subscription
│   │       # - Cancellation reasons
│   │       # - Retention offers
│   │       # - Confirm cancellation
│   │       # BE: financial-be/subscription
│   │       # POST /v1/subscriptions/cancel
│   │
│   ├── change-plan/
│   │   └── page.tsx  # Change subscription plan
│   │       # - Plan comparison
│   │       # - Proration calculation
│   │       # - Confirm change
│   │       # BE: financial-be/subscription
│   │       # POST /v1/subscriptions/change-plan
│   │
│   ├── features/
│   │   └── page.tsx  # Subscription features
│   │       # - Active features
│   │       # - Feature limits
│   │       # - Usage tracking
│   │       # BE: financial-be/subscription
│   │       # GET /v1/subscriptions/features
│   │
│   ├── history/
│   │   └── page.tsx  # Subscription history
│   │       # - Past subscriptions
│   │       # - Billing history
│   │       # BE: financial-be/subscription
│   │       # GET /v1/subscriptions/history
│   │
│   ├── pause/
│   │   └── page.tsx  # Pause subscription
│   │       # - Pause duration
│   │       # - Resume date
│   │       # BE: financial-be/subscription
│   │       # POST /v1/subscriptions/pause
│   │
│   ├── reactivate/
│   │   └── page.tsx  # Reactivate subscription
│   │       # - Choose plan
│   │       # - Payment method
│   │       # BE: financial-be/subscription
│   │       # POST /v1/subscriptions/reactivate
│   │
│   ├── usage/
│   │   └── page.tsx  # Subscription usage metrics
│   │       # - Connects used
│   │       # - Posts remaining
│   │       # - Feature usage
│   │       # BE: financial-be/subscription
│   │       # GET /v1/subscriptions/usage
│   │
│   └── upgrade/
│       └── page.tsx  # Upgrade subscription (in combined, ensuring completeness)
```

### 13. Missing (dashboard) Routes - Teams/Organizations Section

```
apps/web/src/app/[locale]/(dashboard)/
│
├── teams/
│   ├── [teamId]/
│   │   ├── analytics/
│   │   │   └── page.tsx  # Team analytics
│   │   │       # - Team productivity
│   │   │       # - Spending metrics
│   │   │       # - Project outcomes
│   │   │       # BE: users-be/org, financial-be/reports
│   │   │       # GET /v1/teams/{team_id}/analytics
│   │   │
│   │   ├── audit-log/
│   │   │   └── page.tsx  # Team audit log
│   │   │       # - Activity history
│   │   │       # - Change tracking
│   │   │       # BE: users-be/audit
│   │   │       # GET /v1/teams/{team_id}/audit-log
│   │   │
│   │   ├── billing/
│   │   │   ├── invoices/
│   │   │   │   └── page.tsx  # Team invoices
│   │   │   │       # BE: financial-be/invoice
│   │   │   │       # GET /v1/teams/{team_id}/invoices
│   │   │   ├── payment-methods/
│   │   │   │   └── page.tsx  # Team payment methods
│   │   │   │       # BE: financial-be/payment-method
│   │   │   │       # GET /v1/teams/{team_id}/payment-methods
│   │   │   └── page.tsx  # Team billing overview
│   │   │       # BE: financial-be/billing
│   │   │       # GET /v1/teams/{team_id}/billing
│   │   │
│   │   ├── budgets/
│   │   │   ├── [budgetId]/
│   │   │   │   ├── edit/
│   │   │   │   │   └── page.tsx  # Edit budget
│   │   │   │   │       # BE: financial-be/budget
│   │   │   │   │       # PUT /v1/teams/{team_id}/budgets/{budget_id}
│   │   │   │   └── page.tsx  # Budget detail
│   │   │   │       # BE: financial-be/budget
│   │   │   │       # GET /v1/teams/{team_id}/budgets/{budget_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create budget
│   │   │   │       # - Budget amount
│   │   │   │       # - Period
│   │   │   │       # - Allocation
│   │   │   │       # BE: financial-be/budget
│   │   │   │       # POST /v1/teams/{team_id}/budgets
│   │   │   └── page.tsx  # Budgets list
│   │   │       # BE: financial-be/budget
│   │   │       # GET /v1/teams/{team_id}/budgets
│   │   │
│   │   ├── contracts/
│   │   │   └── page.tsx  # Team contracts
│   │   │       # - All team contracts
│   │   │       # - Filter by member
│   │   │       # BE: contracts-be/contract
│   │   │       # GET /v1/teams/{team_id}/contracts
│   │   │
│   │   ├── delete/
│   │   │   └── page.tsx  # Delete team
│   │   │       # - Confirmation
│   │   │       # - Data handling
│   │   │       # BE: users-be/org
│   │   │       # DELETE /v1/teams/{team_id}
│   │   │
│   │   ├── jobs/
│   │   │   └── page.tsx  # Team jobs
│   │   │       # - All team job postings
│   │   │       # - Filter by status
│   │   │       # BE: jobs-be/job
│   │   │       # GET /v1/teams/{team_id}/jobs
│   │   │
│   │   ├── members/
│   │   │   ├── [memberId]/
│   │   │   │   ├── permissions/
│   │   │   │   │   └── page.tsx  # Member permissions
│   │   │   │   │       # BE: users-be/role
│   │   │   │   │       # GET /v1/teams/{team_id}/members/{member_id}/permissions
│   │   │   │   │       # PUT /v1/teams/{team_id}/members/{member_id}/permissions
│   │   │   │   ├── remove/
│   │   │   │   │   └── page.tsx  # Remove member
│   │   │   │   │       # BE: users-be/org
│   │   │   │   │       # DELETE /v1/teams/{team_id}/members/{member_id}
│   │   │   │   └── page.tsx  # Member detail
│   │   │   │       # BE: users-be/org
│   │   │   │       # GET /v1/teams/{team_id}/members/{member_id}
│   │   │   ├── invite/
│   │   │   │   └── page.tsx  # Invite members
│   │   │   │       # - Email invitations
│   │   │   │       # - Role assignment
│   │   │   │       # BE: users-be/org, communications-be/notification
│   │   │   │       # POST /v1/teams/{team_id}/invitations
│   │   │   ├── invitations/
│   │   │   │   ├── [inviteId]/
│   │   │   │   │   └── page.tsx  # Invitation detail
│   │   │   │   │       # - Resend
│   │   │   │   │       # - Revoke
│   │   │   │   │       # BE: users-be/org
│   │   │   │   │       # GET /v1/teams/{team_id}/invitations/{invite_id}
│   │   │   │   └── page.tsx  # Pending invitations
│   │   │   │       # BE: users-be/org
│   │   │   │       # GET /v1/teams/{team_id}/invitations
│   │   │   └── page.tsx  # Members list (in combined)
│   │   │
│   │   ├── projects/
│   │   │   ├── [projectId]/
│   │   │   │   └── page.tsx  # Project detail
│   │   │   │       # - Project overview
│   │   │   │       # - Associated contracts
│   │   │   │       # BE: users-be/org (or projects domain if exists)
│   │   │   │       # GET /v1/teams/{team_id}/projects/{project_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create project
│   │   │   │       # BE: users-be/org
│   │   │   │       # POST /v1/teams/{team_id}/projects
│   │   │   └── page.tsx  # Projects list
│   │   │       # BE: users-be/org
│   │   │       # GET /v1/teams/{team_id}/projects
│   │   │
│   │   ├── reports/
│   │   │   ├── contracts/
│   │   │   │   └── page.tsx  # Contracts report
│   │   │   │       # BE: contracts-be/contract
│   │   │   │       # GET /v1/teams/{team_id}/reports/contracts
│   │   │   ├── expenses/
│   │   │   │   └── page.tsx  # Expenses report
│   │   │   │       # BE: financial-be/reports
│   │   │   │       # GET /v1/teams/{team_id}/reports/expenses
│   │   │   └── page.tsx  # Reports overview
│   │   │       # BE: users-be/org
│   │   │       # GET /v1/teams/{team_id}/reports
│   │   │
│   │   ├── roles/
│   │   │   ├── [roleId]/
│   │   │   │   ├── edit/
│   │   │   │   │   └── page.tsx  # Edit role
│   │   │   │   │       # BE: users-be/role
│   │   │   │   │       # PUT /v1/teams/{team_id}/roles/{role_id}
│   │   │   │   └── page.tsx  # Role detail
│   │   │   │       # BE: users-be/role
│   │   │   │       # GET /v1/teams/{team_id}/roles/{role_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create custom role
│   │   │   │       # BE: users-be/role
│   │   │   │       # POST /v1/teams/{team_id}/roles
│   │   │   └── page.tsx  # Roles list
│   │   │       # BE: users-be/role
│   │   │       # GET /v1/teams/{team_id}/roles
│   │   │
│   │   ├── settings/
│   │   │   ├── branding/
│   │   │   │   └── page.tsx  # Team branding
│   │   │   │       # - Logo upload
│   │   │   │       # - Color scheme
│   │   │   │       # BE: users-be/org, storage-be/asset
│   │   │   │       # PUT /v1/teams/{team_id}/branding
│   │   │   ├── general/
│   │   │   │   └── page.tsx  # General team settings
│   │   │   │       # BE: users-be/org
│   │   │   │       # GET /v1/teams/{team_id}/settings
│   │   │   │       # PUT /v1/teams/{team_id}/settings
│   │   │   └── page.tsx  # Team settings overview
│   │   │
│   │   └── vendors/
│   │       ├── [vendorId]/
│   │       │   └── page.tsx  # Vendor detail
│   │       │       # - Vendor profile
│   │       │       # - Contract history
│   │       │       # BE: users-be/org (vendor management)
│   │       │       # GET /v1/teams/{team_id}/vendors/{vendor_id}
│   │       ├── blacklist/
│   │       │   └── page.tsx  # Blacklisted vendors
│   │       │       # BE: users-be/org
│   │       │       # GET /v1/teams/{team_id}/vendors/blacklist
│   │       ├── preferred/
│   │       │   └── page.tsx  # Preferred vendors
│   │       │       # BE: users-be/org
│   │       │       # GET /v1/teams/{team_id}/vendors/preferred
│   │       └── page.tsx  # Vendors list
│   │           # BE: users-be/org
│   │           # GET /v1/teams/{team_id}/vendors
│   │
│   ├── create/
│   │   └── page.tsx  # Create team/organization (in combined)
│   │
│   └── invitations/
│       └── page.tsx  # Team invitations received
│           # - Accept/decline invitations
│           # BE: users-be/org
│           # GET /v1/users/me/team-invitations
```

### 14. Missing (dashboard) Routes - Admin Section

```
apps/web/src/app/[locale]/(dashboard)/
│
├── admin/
│   ├── analytics/
│   │   ├── platform/
│   │   │   └── page.tsx  # Platform-wide analytics
│   │   │       # - User growth
│   │   │       # - Transaction volume
│   │   │       # - System health
│   │   │       # BE: admin-be/analytics
│   │   │       # GET /v1/admin/analytics/platform
│   │   ├── revenue/
│   │   │   └── page.tsx  # Revenue analytics
│   │   │       # - Revenue trends
│   │   │       # - Fee breakdown
│   │   │       # BE: admin-be/analytics, financial-be/reports
│   │   │       # GET /v1/admin/analytics/revenue
│   │   └── page.tsx  # Analytics overview
│   │       # BE: admin-be/analytics
│   │       # GET /v1/admin/analytics
│   │
│   ├── audits/
│   │   ├── [auditId]/
│   │   │   └── page.tsx  # Audit detail
│   │   │       # - Full audit trail
│   │   │       # - User actions
│   │   │       # BE: admin-be/audit
│   │   │       # GET /v1/admin/audits/{audit_id}
│   │   ├── exports/
│   │   │   └── page.tsx  # Export audit logs
│   │   │       # - Date range
│   │   │       # - Format selection
│   │   │       # BE: admin-be/audit
│   │   │       # POST /v1/admin/audits/export
│   │   └── page.tsx  # Audit logs
│   │       # - System-wide audit trail
│   │       # - Filter capabilities
│   │       # BE: admin-be/audit
│   │       # GET /v1/admin/audits
│   │
│   ├── business-verification/
│   │   ├── [verificationId]/
│   │   │   ├── approve/
│   │   │   │   └── page.tsx  # Approve business
│   │   │   │       # BE: admin-be/business-verification
│   │   │   │       # POST /v1/admin/business-verifications/{verification_id}/approve
│   │   │   ├── documents/
│   │   │   │   └── page.tsx  # Verification documents
│   │   │   │       # - View uploaded docs
│   │   │   │       # - Request additional
│   │   │   │       # BE: admin-be/business-verification, storage-be/asset
│   │   │   │       # GET /v1/admin/business-verifications/{verification_id}/documents
│   │   │   ├── reject/
│   │   │   │   └── page.tsx  # Reject business
│   │   │   │       # - Rejection reason
│   │   │   │       # - Feedback
│   │   │   │       # BE: admin-be/business-verification
│   │   │   │       # POST /v1/admin/business-verifications/{verification_id}/reject
│   │   │   └── page.tsx  # Business verification detail
│   │   │       # - Company info
│   │   │       # - Documents
│   │   │       # - Decision history
│   │   │       # BE: admin-be/business-verification
│   │   │       # GET /v1/admin/business-verifications/{verification_id}
│   │   ├── pending/
│   │   │   └── page.tsx  # Pending verifications
│   │   │       # BE: admin-be/business-verification
│   │   │       # GET /v1/admin/business-verifications?status=pending
│   │   └── page.tsx  # Business verifications queue
│   │       # BE: admin-be/business-verification
│   │       # GET /v1/admin/business-verifications
│   │
│   ├── change-approvals/
│   │   ├── [approvalId]/
│   │   │   ├── approve/
│   │   │   │   └── page.tsx  # Approve change
│   │   │   │       # - Second approver
│   │   │   │       # - Apply change
│   │   │   │       # BE: admin-be/change-approval
│   │   │   │       # POST /v1/admin/change-approvals/{approval_id}/approve
│   │   │   ├── reject/
│   │   │   │   └── page.tsx  # Reject change
│   │   │   │       # BE: admin-be/change-approval
│   │   │   │       # POST /v1/admin/change-approvals/{approval_id}/reject
│   │   │   └── page.tsx  # Change approval detail
│   │   │       # - Change details
│   │   │       # - Risk assessment
│   │   │       # - Approval history
│   │   │       # BE: admin-be/change-approval
│   │   │       # GET /v1/admin/change-approvals/{approval_id}
│   │   ├── pending/
│   │   │   └── page.tsx  # Pending approvals
│   │   │       # BE: admin-be/change-approval
│   │   │       # GET /v1/admin/change-approvals?status=pending
│   │   └── page.tsx  # Change approvals list
│   │       # - Two-person rule changes
│   │       # - Risky operations
│   │       # BE: admin-be/change-approval
│   │       # GET /v1/admin/change-approvals
│   │
│   ├── communications/
│   │   ├── broadcasts/
│   │   │   ├── [broadcastId]/
│   │   │   │   ├── analytics/
│   │   │   │   │   └── page.tsx  # Broadcast analytics
│   │   │   │   │       # - Delivery rates
│   │   │   │   │       # - Engagement metrics
│   │   │   │   │       # BE: communications-be/broadcast
│   │   │   │   │       # GET /v1/admin/broadcasts/{broadcast_id}/analytics
│   │   │   │   └── page.tsx  # Broadcast detail
│   │   │   │       # BE: communications-be/broadcast
│   │   │   │       # GET /v1/admin/broadcasts/{broadcast_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create broadcast
│   │   │   │       # - Compose message
│   │   │   │       # - Target audience
│   │   │   │       # - Schedule
│   │   │   │       # BE: communications-be/broadcast
│   │   │   │       # POST /v1/admin/broadcasts
│   │   │   └── page.tsx  # Broadcasts list
│   │   │       # BE: communications-be/broadcast
│   │   │       # GET /v1/admin/broadcasts
│   │   ├── campaigns/
│   │   │   ├── [campaignId]/
│   │   │   │   └── page.tsx  # Campaign detail
│   │   │   │       # BE: communications-be/campaign
│   │   │   │       # GET /v1/admin/campaigns/{campaign_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create campaign
│   │   │   │       # BE: communications-be/campaign
│   │   │   │       # POST /v1/admin/campaigns
│   │   │   └── page.tsx  # Campaigns list
│   │   │       # BE: communications-be/campaign
│   │   │       # GET /v1/admin/campaigns
│   │   ├── rate-limits/
│   │   │   └── page.tsx  # Communication rate limits
│   │   │       # - Per-user limits
│   │   │       # - Global throttling
│   │   │       # BE: communications-be/preferences
│   │   │       # GET /v1/admin/communications/rate-limits
│   │   │       # PUT /v1/admin/communications/rate-limits
│   │   └── templates/
│   │       ├── [templateId]/
│   │       │   ├── edit/
│   │       │   │   └── page.tsx  # Edit template
│   │       │   │       # BE: communications-be/template
│   │       │   │       # PUT /v1/admin/templates/{template_id}
│   │       │   ├── preview/
│   │       │   │   └── page.tsx  # Preview template
│   │       │   │       # BE: communications-be/template
│   │       │   │       # POST /v1/admin/templates/{template_id}/preview
│   │       │   └── page.tsx  # Template detail
│   │       │       # BE: communications-be/template
│   │       │       # GET /v1/admin/templates/{template_id}
│   │       ├── create/
│   │       │   └── page.tsx  # Create template
│   │       │       # BE: communications-be/template
│   │       │       # POST /v1/admin/templates
│   │       └── page.tsx  # Templates list
│   │           # BE: communications-be/template
│   │           # GET /v1/admin/templates
│   │
│   ├── configurations/
│   │   ├── feature-flags/
│   │   │   ├── [flagId]/
│   │   │   │   └── page.tsx  # Feature flag detail
│   │   │   │       # - Toggle flag
│   │   │   │       # - Rollout percentage
│   │   │   │       # BE: admin-be/config (or utility-be/flags)
│   │   │   │       # GET /v1/admin/feature-flags/{flag_id}
│   │   │   │       # PUT /v1/admin/feature-flags/{flag_id}
│   │   │   └── page.tsx  # Feature flags list
│   │   │       # BE: admin-be/config
│   │   │       # GET /v1/admin/feature-flags
│   │   ├── experiments/
│   │   │   ├── [experimentId]/
│   │   │   │   ├── results/
│   │   │   │   │   └── page.tsx  # Experiment results
│   │   │   │   │       # - A/B test metrics
│   │   │   │   │       # - Statistical significance
│   │   │   │   │       # BE: admin-be/experiments
│   │   │   │   │       # GET /v1/admin/experiments/{experiment_id}/results
│   │   │   │   └── page.tsx  # Experiment detail
│   │   │   │       # BE: admin-be/experiments
│   │   │   │       # GET /v1/admin/experiments/{experiment_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create experiment
│   │   │   │       # BE: admin-be/experiments
│   │   │   │       # POST /v1/admin/experiments
│   │   │   └── page.tsx  # Experiments list
│   │   │       # BE: admin-be/experiments
│   │   │       # GET /v1/admin/experiments
│   │   └── page.tsx  # System configurations
│   │       # BE: admin-be/config
│   │       # GET /v1/admin/configurations
│   │
│   ├── content-moderation/
│   │   ├── [reportId]/
│   │   │   ├── actions/
│   │   │   │   └── page.tsx  # Moderation actions
│   │   │   │       # - Warning
│   │   │   │       # - Suspension
│   │   │   │       # - Ban
│   │   │   │       # BE: admin-be/moderation
│   │   │   │       # POST /v1/admin/reports/{report_id}/actions
│   │   │   └── page.tsx  # Report detail
│   │   │       # - Content review
│   │   │       # - Reporter info
│   │   │       # - History
│   │   │       # BE: admin-be/moderation
│   │   │       # GET /v1/admin/reports/{report_id}
│   │   ├── appeals/
│   │   │   ├── [appealId]/
│   │   │   │   ├── approve/
│   │   │   │   │   └── page.tsx  # Approve appeal
│   │   │   │   │       # BE: admin-be/moderation
│   │   │   │   │       # POST /v1/admin/appeals/{appeal_id}/approve
│   │   │   │   ├── reject/
│   │   │   │   │   └── page.tsx  # Reject appeal
│   │   │   │   │       # BE: admin-be/moderation
│   │   │   │   │       # POST /v1/admin/appeals/{appeal_id}/reject
│   │   │   │   └── page.tsx  # Appeal detail
│   │   │   │       # BE: admin-be/moderation
│   │   │   │       # GET /v1/admin/appeals/{appeal_id}
│   │   │   └── page.tsx  # Appeals queue
│   │   │       # BE: admin-be/moderation
│   │   │       # GET /v1/admin/appeals
│   │   ├── queue/
│   │   │   └── page.tsx  # Moderation queue
│   │   │       # - Pending reports
│   │   │       # - Priority sorting
│   │   │       # BE: admin-be/moderation
│   │   │       # GET /v1/admin/reports?status=pending
│   │   └── rules/
│   │       └── page.tsx  # Moderation rules
│   │           # - Auto-moderation rules
│   │           # - Keyword filters
│   │           # BE: admin-be/moderation
│   │           # GET /v1/admin/moderation/rules
│   │           # PUT /v1/admin/moderation/rules
│   │
│   ├── financial-ops/
│   │   ├── chargebacks/
│   │   │   ├── [chargebackId]/
│   │   │   │   ├── dispute/
│   │   │   │   │   └── page.tsx  # Dispute chargeback
│   │   │   │   │       # BE: financial-be/chargeback, admin-be/financial-ops
│   │   │   │   │       # POST /v1/admin/chargebacks/{chargeback_id}/dispute
│   │   │   │   └── page.tsx  # Chargeback detail
│   │   │   │       # BE: financial-be/chargeback
│   │   │   │       # GET /v1/admin/chargebacks/{chargeback_id}
│   │   │   └── page.tsx  # Chargebacks list
│   │   │       # BE: financial-be/chargeback
│   │   │       # GET /v1/admin/chargebacks
│   │   ├── goodwill-credits/
│   │   │   ├── [creditId]/
│   │   │   │   └── page.tsx  # Goodwill credit detail
│   │   │   │       # BE: admin-be/goodwill-credit
│   │   │   │       # GET /v1/admin/goodwill-credits/{credit_id}
│   │   │   ├── approve/
│   │   │   │   └── page.tsx  # Approve goodwill credit
│   │   │   │       # BE: admin-be/goodwill-credit
│   │   │   │       # POST /v1/admin/goodwill-credits/{credit_id}/approve
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create goodwill credit
│   │   │   │       # BE: admin-be/goodwill-credit
│   │   │   │       # POST /v1/admin/goodwill-credits
│   │   │   └── page.tsx  # Goodwill credits list
│   │   │       # BE: admin-be/goodwill-credit
│   │   │       # GET /v1/admin/goodwill-credits
│   │   ├── payouts/
│   │   │   ├── [payoutId]/
│   │   │   │   ├── approve/
│   │   │   │   │   └── page.tsx  # Approve payout
│   │   │   │   │       # BE: financial-be/payout
│   │   │   │   │       # POST /v1/admin/payouts/{payout_id}/approve
│   │   │   │   ├── hold/
│   │   │   │   │   └── page.tsx  # Hold payout
│   │   │   │   │       # BE: financial-be/payout
│   │   │   │   │       # POST /v1/admin/payouts/{payout_id}/hold
│   │   │   │   └── page.tsx  # Payout detail
│   │   │   │       # BE: financial-be/payout
│   │   │   │       # GET /v1/admin/payouts/{payout_id}
│   │   │   └── page.tsx  # Payouts review queue
│   │   │       # BE: financial-be/payout
│   │   │       # GET /v1/admin/payouts?status=pending
│   │   ├── reconciliation/
│   │   │   ├── [reconId]/
│   │   │   │   └── page.tsx  # Reconciliation detail
│   │   │   │       # BE: financial-be/reconciliation
│   │   │   │       # GET /v1/admin/reconciliation/{recon_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Start reconciliation
│   │   │   │       # BE: financial-be/reconciliation
│   │   │   │       # POST /v1/admin/reconciliation
│   │   │   └── page.tsx  # Reconciliation reports
│   │   │       # BE: financial-be/reconciliation
│   │   │       # GET /v1/admin/reconciliation
│   │   └── refund-cases/
│   │       ├── [caseId]/
│   │       │   ├── approve/
│   │       │   │   └── page.tsx  # Approve refund
│   │       │   │       # BE: admin-be/refund-case
│   │       │   │       # POST /v1/admin/refund-cases/{case_id}/approve
│   │       │   ├── investigation/
│   │       │   │   └── page.tsx  # Investigation notes
│   │       │   │       # BE: admin-be/refund-case
│   │       │   │       # GET /v1/admin/refund-cases/{case_id}/investigation
│   │       │   │       # POST /v1/admin/refund-cases/{case_id}/investigation
│   │       │   ├── reject/
│   │       │   │   └── page.tsx  # Reject refund
│   │       │   │       # BE: admin-be/refund-case
│   │       │   │       # POST /v1/admin/refund-cases/{case_id}/reject
│   │       │   └── page.tsx  # Refund case detail
│   │       │       # BE: admin-be/refund-case
│   │       │       # GET /v1/admin/refund-cases/{case_id}
│   │       └── page.tsx  # Refund cases queue
│   │           # BE: admin-be/refund-case
│   │           # GET /v1/admin/refund-cases
│   │
│   ├── kyc/
│   │   ├── [kycId]/
│   │   │   ├── approve/
│   │   │   │   └── page.tsx  # Approve KYC
│   │   │   │       # BE: admin-be/kyc-case
│   │   │   │       # POST /v1/admin/kyc/{kyc_id}/approve
│   │   │   ├── documents/
│   │   │   │   └── page.tsx  # KYC documents review
│   │   │   │       # - ID verification
│   │   │   │       # - Address proof
│   │   │   │       # BE: admin-be/kyc-case, storage-be/asset
│   │   │   │       # GET /v1/admin/kyc/{kyc_id}/documents
│   │   │   ├── reject/
│   │   │   │   └── page.tsx  # Reject KYC
│   │   │   │       # - Rejection reason
│   │   │   │       # - Request resubmission
│   │   │   │       # BE: admin-be/kyc-case
│   │   │   │       # POST /v1/admin/kyc/{kyc_id}/reject
│   │   │   ├── reopen/
│   │   │   │   └── page.tsx  # Reopen KYC case
│   │   │   │       # BE: admin-be/kyc-case
│   │   │   │       # POST /v1/admin/kyc/{kyc_id}/reopen
│   │   │   └── page.tsx  # KYC case detail
│   │   │       # - User information
│   │   │       # - Documents
│   │   │       # - Decision history
│   │   │       # BE: admin-be/kyc-case
│   │   │       # GET /v1/admin/kyc/{kyc_id}
│   │   ├── pending/
│   │   │   └── page.tsx  # Pending KYC cases
│   │   │       # BE: admin-be/kyc-case
│   │   │       # GET /v1/admin/kyc?status=pending
│   │   ├── rejected/
│   │   │   └── page.tsx  # Rejected KYC cases
│   │   │       # BE: admin-be/kyc-case
│   │   │       # GET /v1/admin/kyc?status=rejected
│   │   └── page.tsx  # KYC queue
│   │       # - Triage queue
│   │       # - Priority sorting
│   │       # BE: admin-be/kyc-case
│   │       # GET /v1/admin/kyc
│   │
│   ├── search-quality/
│   │   ├── blacklists/
│   │   │   └── page.tsx  # Search blacklists
│   │   │       # - Blacklisted terms
│   │   │       # - Blocked users
│   │   │       # BE: search-be/admin
│   │   │       # GET /v1/admin/search/blacklists
│   │   │       # PUT /v1/admin/search/blacklists
│   │   ├── boosts/
│   │   │   └── page.tsx  # Search boosts
│   │   │       # - Boosted content
│   │   │       # - Boost rules
│   │   │       # BE: search-be/admin
│   │   │       # GET /v1/admin/search/boosts
│   │   │       # PUT /v1/admin/search/boosts
│   │   ├── reindex/
│   │   │   └── page.tsx  # Reindex controls
│   │   │       # - Trigger reindex
│   │   │       # - Monitor progress
│   │   │       # BE: search-be/admin
│   │   │       # POST /v1/admin/search/reindex
│   │   └── synonyms/
│   │       └── page.tsx  # Search synonyms
│   │           # - Synonym management
│   │           # - Query expansion rules
│   │           # BE: search-be/admin
│   │           # GET /v1/admin/search/synonyms
│   │           # PUT /v1/admin/search/synonyms
│   │
│   ├── sessions/
│   │   ├── [sessionId]/
│   │   │   ├── approve/
│   │   │   │   └── page.tsx  # Approve JIT session
│   │   │   │       # - Second approver required
│   │   │   │       # - Time-box access
│   │   │   │       # BE: admin-be/admin-session
│   │   │   │       # POST /v1/admin/sessions/{session_id}/approve
│   │   │   ├── revoke/
│   │   │   │   └── page.tsx  # Revoke session
│   │   │   │       # BE: admin-be/admin-session
│   │   │   │       # POST /v1/admin/sessions/{session_id}/revoke
│   │   │   └── page.tsx  # Session detail
│   │   │       # - Session info
│   │   │       # - Actions performed
│   │   │       # BE: admin-be/admin-session
│   │   │       # GET /v1/admin/sessions/{session_id}
│   │   ├── request/
│   │   │   └── page.tsx  # Request JIT session
│   │   │       # - Justification
│   │   │       # - Requested duration
│   │   │       # BE: admin-be/admin-session
│   │   │       # POST /v1/admin/sessions/request
│   │   └── page.tsx  # Admin sessions list
│   │       # - Active sessions
│   │       # - Session history
│   │       # BE: admin-be/admin-session
│   │       # GET /v1/admin/sessions
│   │
│   ├── system-health/
│   │   ├── incidents/
│   │   │   ├── [incidentId]/
│   │   │   │   ├── postmortem/
│   │   │   │   │   └── page.tsx  # Incident postmortem
│   │   │   │   │       # BE: admin-be/incidents (or utility-be/status)
│   │   │   │   │       # GET /v1/admin/incidents/{incident_id}/postmortem
│   │   │   │   │       # POST /v1/admin/incidents/{incident_id}/postmortem
│   │   │   │   ├── resolve/
│   │   │   │   │   └── page.tsx  # Resolve incident
│   │   │   │   │       # BE: admin-be/incidents
│   │   │   │   │       # POST /v1/admin/incidents/{incident_id}/resolve
│   │   │   │   └── page.tsx  # Incident detail
│   │   │   │       # BE: admin-be/incidents
│   │   │   │       # GET /v1/admin/incidents/{incident_id}
│   │   │   ├── create/
│   │   │   │   └── page.tsx  # Create incident
│   │   │   │       # BE: admin-be/incidents
│   │   │   │       # POST /v1/admin/incidents
│   │   │   └── page.tsx  # Incidents list
│   │   │       # BE: admin-be/incidents
│   │   │       # GET /v1/admin/incidents
│   │   ├── maintenance/
│   │   │   ├── [maintenanceId]/
│   │   │   │   └── page.tsx  # Maintenance detail
│   │   │   │       # BE: admin-be/maintenance
│   │   │   │       # GET /v1/admin/maintenance/{maintenance_id}
│   │   │   ├── schedule/
│   │   │   │   └── page.tsx  # Schedule maintenance
│   │   │   │       # BE: admin-be/maintenance
│   │   │   │       # POST /v1/admin/maintenance
│   │   │   └── page.tsx  # Maintenance windows
│   │   │       # BE: admin-be/maintenance
│   │   │       # GET /v1/admin/maintenance
│   │   └── status/
│   │       └── page.tsx  # System status dashboard
│   │           # - Service health
│   │           # - Uptime metrics
│   │           # BE: admin-be/system (or utility-be/status)
│   │           # GET /v1/admin/system/status
│   │
│   ├── users/
│   │   ├── [userId]/
│   │   │   ├── ban/
│   │   │   │   └── page.tsx  # Ban user
│   │   │   │       # - Ban reason
│   │   │   │       # - Duration
│   │   │   │       # BE: users-be/account, admin-be/moderation
│   │   │   │       # POST /v1/admin/users/{user_id}/ban
│   │   │   ├── contracts/
│   │   │   │   └── page.tsx  # User contracts
│   │   │   │       # BE: contracts-be/contract
│   │   │   │       # GET /v1/admin/users/{user_id}/contracts
│   │   │   ├── financials/
│   │   │   │   └── page.tsx  # User financial history
│   │   │   │       # BE: financial-be/reports
│   │   │   │       # GET /v1/admin/users/{user_id}/financials
│   │   │   ├── impersonate/
│   │   │   │   └── page.tsx  # Impersonate user
│   │   │   │       # - Requires approval
│   │   │   │       # - Audit trail
│   │   │   │       # BE: admin-be/admin-session
│   │   │   │       # POST /v1/admin/users/{user_id}/impersonate
│   │   │   ├── suspend/
│   │   │   │   └── page.tsx  # Suspend user
│   │   │   │       # BE: users-be/account, admin-be/moderation
│   │   │   │       # POST /v1/admin/users/{user_id}/suspend
│   │   │   ├── unban/
│   │   │   │   └── page.tsx  # Unban user
│   │   │   │       # BE: users-be/account
│   │   │   │       # POST /v1/admin/users/{user_id}/unban
│   │   │   ├── verify/
│   │   │   │   └── page.tsx  # Manually verify user
│   │   │   │       # BE: users-be/account
│   │   │   │       # POST /v1/admin/users/{user_id}/verify
│   │   │   ├── warning/
│   │   │   │   └── page.tsx  # Issue warning
│   │   │   │       # BE: admin-be/moderation
│   │   │   │       # POST /v1/admin/users/{user_id}/warning
│   │   │   └── page.tsx  # User detail
│   │   │       # - Full profile
│   │   │       # - Activity history
│   │   │       # - Actions menu
│   │   │       # BE: users-be/profile
│   │   │       # GET /v1/admin/users/{user_id}
│   │   ├── banned/
│   │   │   └── page.tsx  # Banned users
│   │   │       # BE: users-be/account
│   │   │       # GET /v1/admin/users?status=banned
│   │   ├── search/
│   │   │   └── page.tsx  # Search users
│   │   │       # - Advanced search
│   │   │       # - Bulk actions
│   │   │       # BE: search-be/query
│   │   │       # POST /v1/admin/users/search
│   │   └── suspended/
│   │       └── page.tsx  # Suspended users
│   │           # BE: users-be/account
│   │           # GET /v1/admin/users?status=suspended
│   │
│   └── page.tsx  # Admin dashboard
│       # - Key metrics
│       # - Quick actions
│       # - Recent activities
│       # BE: admin-be/analytics
│       # GET /v1/admin/dashboard
```

### 15. Missing (public) Routes Section

```
apps/web/src/app/[locale]/(public)/
│
├── about/
│   ├── leadership/
│   │   └── page.tsx  # Leadership team
│   │       # BE: None (static)
│   ├── press/
│   │   └── page.tsx  # Press releases
│   │       # BE: None (static or CMS)
│   └── page.tsx  # About page (in combined)
│
├── case-studies/
│   ├── [slug]/
│   │   └── page.tsx  # Case study detail
│   │       # BE: None (static or CMS)
│   └── page.tsx  # Case studies list
│       # BE: None (static or CMS)
│
├── developers/
│   ├── api/
│   │   └── page.tsx  # API documentation
│   │       # BE: None (static)
│   ├── changelog/
│   │   └── page.tsx  # API changelog
│   │       # BE: None (static)
│   └── page.tsx  # Developer portal
│       # BE: None (static)
│
├── enterprise/
│   ├── contact/
│   │   └── page.tsx  # Enterprise contact
│   │       # BE: communications-be/messages
│   │       # POST /v1/contact/enterprise
│   ├── demo/
│   │   └── page.tsx  # Request demo
│   │       # BE: communications-be/messages
│   │       # POST /v1/contact/demo
│   └── page.tsx  # Enterprise solutions
│       # BE: None (static)
│
├── legal/
│   ├── cookies/
│   │   └── page.tsx  # Cookie policy
│   │       # BE: None (static)
│   ├── dmca/
│   │   └── page.tsx  # DMCA policy
│   │       # BE: None (static)
│   ├── privacy/
│   │   └── page.tsx  # Privacy policy
│   │       # BE: None (static)
│   └── terms/
│       └── page.tsx  # Terms of service
│           # BE: None (static)
│
├── partners/
│   ├── become-partner/
│   │   └── page.tsx  # Partner application
│   │       # BE: communications-be/messages
│   │       # POST /v1/partners/apply
│   ├── directory/
│   │   └── page.tsx  # Partners directory
│   │       # BE: None (static or CMS)
│   └── page.tsx  # Partners program
│       # BE: None (static)
│
├── pricing/
│   ├── compare/
│   │   └── page.tsx  # Plan comparison
│   │       # BE: financial-be/subscription
│   │       # GET /v1/subscriptions/plans/compare
│   ├── enterprise/
│   │   └── page.tsx  # Enterprise pricing
│   │       # BE: None (static)
│   └── page.tsx  # Pricing overview
│       # BE: financial-be/subscription
│       # GET /v1/subscriptions/plans
│
├── resources/
│   ├── guides/
│   │   ├── [slug]/
│   │   │   └── page.tsx  # Guide detail
│   │   │       # BE: None (static or CMS)
│   │   └── page.tsx  # Guides list
│   │       # BE: None (static or CMS)
│   ├── tutorials/
│   │   ├── [slug]/
│   │   │   └── page.tsx  # Tutorial detail
│   │   │       # BE: None (static or CMS)
│   │   └── page.tsx  # Tutorials list
│   │       # BE: None (static or CMS)
│   ├── webinars/
│   │   ├── [id]/
│   │   │   ├── register/
│   │   │   │   └── page.tsx  # Register for webinar
│   │   │   │       # BE: communications-be/events (if exists)
│   │   │   │       # POST /v1/webinars/{id}/register
│   │   │   └── page.tsx  # Webinar detail
│   │   │       # BE: None (static or CMS)
│   │   └── page.tsx  # Webinars list
│   │       # BE: None (static or CMS)
│   └── page.tsx  # Resources hub
│       # BE: None (static)
│
├── security/
│   └── page.tsx  # Security information
│       # - Security practices
│       # - Compliance certifications
│       # - Vulnerability disclosure
│       # BE: None (static)
│
├── sitemap.xml  # Dynamic sitemap
│   # BE: Multiple services for dynamic content
│
└── status/
    └── page.tsx  # System status page
        # - Service status
        # - Incident history
        # BE: admin-be/system (or utility-be/status)
        # GET /v1/status
```

---

## Missing Mobile App Routes (apps/mobile/app/)

### 1. Missing Mobile Auth Routes

```
apps/mobile/app/(auth)/
│
├── forgot-password.tsx  # Password reset flow
│   # - Email input
│   # - Code verification
│   # BE: users-be/auth
│   # POST /v1/auth/forgot-password
│
├── reset-password.tsx  # Reset password with code
│   # - New password input
│   # - Confirmation
│   # BE: users-be/auth
│   # POST /v1/auth/reset-password
│
└── verify-email.tsx  # Email verification
    # - Code input
    # - Resend code
    # BE: users-be/auth
    # POST /v1/auth/verify-email
```

### 2. Missing Mobile Main App Routes

```
apps/mobile/app/(tabs)/
│
├── admin/
│   ├── _layout.tsx  # Admin tab layout
│   ├── index.tsx  # Admin dashboard (mobile)
│   │   # BE: admin-be/analytics
│   │   # GET /v1/admin/dashboard
│   ├── kyc/
│   │   ├── [kycId].tsx  # KYC case detail (mobile)
│   │   │   # BE: admin-be/kyc-case
│   │   │   # GET /v1/admin/kyc/{kyc_id}
│   │   └── index.tsx  # KYC queue (mobile)
│   │       # BE: admin-be/kyc-case
│   │       # GET /v1/admin/kyc
│   ├── moderation/
│   │   ├── [reportId].tsx  # Report detail (mobile)
│   │   │   # BE: admin-be/moderation
│   │   │   # GET /v1/admin/reports/{report_id}
│   │   └── index.tsx  # Moderation queue (mobile)
│   │       # BE: admin-be/moderation
│   │       # GET /v1/admin/reports
│   └── users/
│       ├── [userId].tsx  # User detail (mobile admin)
│       │   # BE: users-be/profile
│       │   # GET /v1/admin/users/{user_id}
│       └── index.tsx  # Users search (mobile)
│           # BE: search-be/query
│           # POST /v1/admin/users/search
│
├── browse/
│   ├── _layout.tsx  # Browse tab layout
│   ├── categories.tsx  # Browse categories
│   │   # BE: jobs-be/categories
│   │   # GET /v1/jobs/categories
│   ├── freelancers/
│   │   ├── [userId].tsx  # Freelancer profile (mobile)
│   │   │   # BE: users-be/profile
│   │   │   # GET /v1/users/{user_id}
│   │   ├── filters.tsx  # Freelancer search filters
│   │   └── index.tsx  # Freelancers list (mobile)
│   │       # BE: search-be/query
│   │       # GET /v1/search/freelancers
│   └── jobs/
│       ├── [jobId].tsx  # Job detail (mobile)
│       │   # BE: jobs-be/job
│       │   # GET /v1/jobs/{job_id}
│       ├── filters.tsx  # Job search filters
│       └── index.tsx  # Jobs list (mobile)
│           # BE: search-be/query
│           # GET /v1/search/jobs
│
├── contracts/
│   ├── _layout.tsx  # Contracts tab layout
│   ├── [contractId]/
│   │   ├── chat.tsx  # Contract chat (mobile)
│   │   │   # BE: communications-be/conversation
│   │   │   # GET /v1/contracts/{contract_id}/messages
│   │   ├── deliverables.tsx  # Deliverables (mobile)
│   │   │   # BE: contracts-be/deliverable
│   │   │   # GET /v1/contracts/{contract_id}/deliverables
│   │   ├── details.tsx  # Contract details (mobile)
│   │   │   # BE: contracts-be/contract
│   │   │   # GET /v1/contracts/{contract_id}
│   │   ├── disputes.tsx  # Contract disputes (mobile)
│   │   │   # BE: contracts-be/dispute
│   │   │   # GET /v1/contracts/{contract_id}/disputes
│   │   ├── milestones.tsx  # Milestones (mobile)
│   │   │   # BE: contracts-be/milestone
│   │   │   # GET /v1/contracts/{contract_id}/milestones
│   │   └── workdiary.tsx  # Work diary (mobile)
│   │       # BE: contracts-be/workdiary
│   │       # GET /v1/contracts/{contract_id}/workdiary
│   ├── active.tsx  # Active contracts (mobile)
│   │   # BE: contracts-be/contract
│   │   # GET /v1/contracts?status=active
│   ├── completed.tsx  # Completed contracts (mobile)
│   │   # BE: contracts-be/contract
│   │   # GET /v1/contracts?status=completed
│   └── index.tsx  # All contracts (mobile)
│       # BE: contracts-be/contract
│       # GET /v1/contracts
│
├── dashboard/
│   ├── _layout.tsx  # Dashboard tab layout
│   ├── analytics.tsx  # User analytics (mobile)
│   │   # BE: users-be/analytics
│   │   # GET /v1/users/me/analytics
│   ├── earnings.tsx  # Earnings overview (mobile)
│   │   # BE: financial-be/reports
│   │   # GET /v1/reports/earnings
│   ├── notifications.tsx  # Notifications (mobile)
│   │   # BE: communications-be/notification
│   │   # GET /v1/notifications
│   └── index.tsx  # Main dashboard (mobile)
│       # BE: Multiple services
│       # Aggregated dashboard data
│
├── financial/
│   ├── _layout.tsx  # Financial tab layout
│   ├── disputes/
│   │   ├── [disputeId].tsx  # Dispute detail (mobile)
│   │   │   # BE: contracts-be/dispute
│   │   │   # GET /v1/disputes/{dispute_id}
│   │   └── index.tsx  # Disputes list (mobile)
│   │       # BE: contracts-be/dispute
│   │       # GET /v1/disputes
│   ├── invoices/
│   │   ├── [invoiceId].tsx  # Invoice detail (mobile)
│   │   │   # BE: financial-be/invoice
│   │   │   # GET /v1/invoices/{invoice_id}
│   │   └── index.tsx  # Invoices list (mobile)
│   │       # BE: financial-be/invoice
│   │       # GET /v1/invoices
│   ├── payouts/
│   │   ├── [payoutId].tsx  # Payout detail (mobile)
│   │   │   # BE: financial-be/payout
│   │   │   # GET /v1/payouts/{payout_id}
│   │   ├── request.tsx  # Request payout (mobile)
│   │   │   # BE: financial-be/payout
│   │   │   # POST /v1/payouts
│   │   └── index.tsx  # Payouts list (mobile)
│   │       # BE: financial-be/payout
│   │       # GET /v1/payouts
│   ├── transactions.tsx  # Transaction history (mobile)
│   │   # BE: financial-be/transaction
│   │   # GET /v1/transactions
│   └── wallet.tsx  # Wallet (mobile)
│       # BE: financial-be/wallet
│       # GET /v1/wallet
│
├── jobs/
│   ├── _layout.tsx  # Jobs tab layout
│   ├── [jobId]/
│   │   ├── apply.tsx  # Apply to job (mobile)
│   │   │   # BE: proposals-be/proposal
│   │   │   # POST /v1/proposals
│   │   ├── proposals.tsx  # Job proposals (client view, mobile)
│   │   │   # BE: proposals-be/proposal
│   │   │   # GET /v1/jobs/{job_id}/proposals
│   │   └── details.tsx  # Job details (mobile)
│   │       # BE: jobs-be/job
│   │       # GET /v1/jobs/{job_id}
│   ├── create.tsx  # Create job (mobile)
│   │   # BE: jobs-be/job
│   │   # POST /v1/jobs
│   ├── drafts.tsx  # Job drafts (mobile)
│   │   # BE: jobs-be/draft
│   │   # GET /v1/jobs/drafts
│   ├── my-jobs.tsx  # My posted jobs (mobile)
│   │   # BE: jobs-be/job
│   │   # GET /v1/jobs/my-jobs
│   ├── saved.tsx  # Saved jobs (mobile)
│   │   # BE: jobs-be/saved
│   │   # GET /v1/jobs/saved
│   └── index.tsx  # Browse jobs (mobile)
│       # BE: search-be/query
│       # GET /v1/search/jobs
│
├── messages/
│   ├── _layout.tsx  # Messages tab layout
│   ├── [conversationId]/
│   │   ├── details.tsx  # Conversation details (mobile)
│   │   │   # BE: communications-be/conversation
│   │   │   # GET /v1/conversations/{conversation_id}
│   │   └── chat.tsx  # Chat interface (mobile)
│   │       # BE: communications-be/conversation
│   │       # GET /v1/conversations/{conversation_id}/messages
│   │       # POST /v1/conversations/{conversation_id}/messages
│   ├── archived.tsx  # Archived conversations (mobile)
│   │   # BE: communications-be/conversation
│   │   # GET /v1/conversations?archived=true
│   ├── compose.tsx  # New message (mobile)
│   │   # BE: communications-be/conversation
│   │   # POST /v1/conversations
│   └── index.tsx  # Conversations list (mobile)
│       # BE: communications-be/conversation
│       # GET /v1/conversations
│
├── profile/
│   ├── _layout.tsx  # Profile tab layout
│   ├── edit/
│   │   ├── basic.tsx  # Edit basic info (mobile)
│   │   │   # BE: users-be/profile
│   │   │   # PUT /v1/users/me/profile
│   │   ├── education.tsx  # Edit education (mobile)
│   │   │   # BE: users-be/profile
│   │   │   # PUT /v1/users/me/education
│   │   ├── experience.tsx  # Edit experience (mobile)
│   │   │   # BE: users-be/profile
│   │   │   # PUT /v1/users/me/experience
│   │   ├── portfolio.tsx  # Edit portfolio (mobile)
│   │   │   # BE: users-be/portfolio, storage-be/asset
│   │   │   # PUT /v1/users/me/portfolio
│   │   └── skills.tsx  # Edit skills (mobile)
│   │       # BE: users-be/profile
│   │       # PUT /v1/users/me/skills
│   ├── portfolio/
│   │   ├── [itemId].tsx  # Portfolio item detail (mobile)
│   │   │   # BE: users-be/portfolio
│   │   │   # GET /v1/users/me/portfolio/{item_id}
│   │   ├── add.tsx  # Add portfolio item (mobile)
│   │   │   # BE: users-be/portfolio, storage-be/asset
│   │   │   # POST /v1/users/me/portfolio
│   │   └── index.tsx  # Portfolio list (mobile)
│   │       # BE: users-be/portfolio
│   │       # GET /v1/users/me/portfolio
│   ├── reviews.tsx  # User reviews (mobile)
│   │   # BE: reviews-be/review
│   │   # GET /v1/users/me/reviews
│   ├── settings.tsx  # Profile settings (mobile)
│   │   # BE: users-be/preferences
│   │   # GET /v1/users/me/preferences
│   └── index.tsx  # View profile (mobile)
│       # BE: users-be/profile
│       # GET /v1/users/me
│
├── proposals/
│   ├── _layout.tsx  # Proposals tab layout
│   ├── [proposalId]/
│   │   ├── details.tsx  # Proposal details (mobile)
│   │   │   # BE: proposals-be/proposal
│   │   │   # GET /v1/proposals/{proposal_id}
│   │   ├── edit.tsx  # Edit proposal (mobile)
│   │   │   # BE: proposals-be/proposal
│   │   │   # PUT /v1/proposals/{proposal_id}
│   │   └── withdraw.tsx  # Withdraw proposal (mobile)
│   │       # BE: proposals-be/proposal
│   │       # POST /v1/proposals/{proposal_id}/withdraw
│   ├── active.tsx  # Active proposals (mobile)
│   │   # BE: proposals-be/proposal
│   │   # GET /v1/proposals?status=active
│   ├── archived.tsx  # Archived proposals (mobile)
│   │   # BE: proposals-be/proposal
│   │   # GET /v1/proposals?status=archived
│   ├── drafts.tsx  # Proposal drafts (mobile)
│   │   # BE: proposals-be/draft
│   │   # GET /v1/proposals/drafts
│   └── index.tsx  # All proposals (mobile)
│       # BE: proposals-be/proposal
│       # GET /v1/proposals
│
├── search/
│   ├── _layout.tsx  # Search tab layout
│   ├── filters.tsx  # Search filters (mobile)
│   ├── freelancers.tsx  # Search freelancers (mobile)
│   │   # BE: search-be/query
│   │   # GET /v1/search/freelancers
│   ├── jobs.tsx  # Search jobs (mobile)
│   │   # BE: search-be/query
│   │   # GET /v1/search/jobs
│   ├── portfolios.tsx  # Search portfolios (mobile)
│   │   # BE: search-be/query
│   │   # GET /v1/search/portfolios
│   └── saved-searches.tsx  # Saved searches (mobile)
│       # BE: search-be/saved-search
│       # GET /v1/search/saved-searches
│
└── settings/
    ├── _layout.tsx  # Settings tab layout
    ├── account/
    │   ├── close.tsx  # Close account (mobile)
    │   │   # BE: users-be/account
    │   │   # POST /v1/users/me/close-account
    │   ├── email.tsx  # Change email (mobile)
    │   │   # BE: users-be/account
    │   │   # PUT /v1/users/me/email
    │   ├── password.tsx  # Change password (mobile)
    │   │   # BE: users-be/auth
    │   │   # PUT /v1/auth/password
    │   └── phone.tsx  # Change phone (mobile)
    │       # BE: users-be/account
    │       # PUT /v1/users/me/phone
    ├── billing/
    │   ├── payment-methods.tsx  # Payment methods (mobile)
    │   │   # BE: financial-be/payment-method
    │   │   # GET /v1/payment-methods
    │   └── subscription.tsx  # Subscription (mobile)
    │       # BE: financial-be/subscription
    │       # GET /v1/subscriptions
    ├── notifications.tsx  # Notification settings (mobile)
    │   # BE: communications-be/preferences
    │   # GET /v1/notifications/preferences
    ├── privacy.tsx  # Privacy settings (mobile)
    │   # BE: users-be/privacy
    │   # GET /v1/users/me/privacy
    ├── security.tsx  # Security settings (mobile)
    │   # - Two-factor auth
    │   # - Active sessions
    │   # BE: users-be/mfa, users-be/sessions
    │   # GET /v1/users/me/mfa
    │   # GET /v1/users/me/sessions
    └── index.tsx  # Settings overview (mobile)
```

---

## Missing packages/shared/src/features/ Structure

### Missing Feature Modules

```
packages/shared/src/features/
│
├── admin/
│   ├── api/
│   │   ├── admin-api.ts  # Admin API client
│   │   │   # BE: admin-be/*
│   │   ├── analytics-api.ts  # Admin analytics API
│   │   │   # BE: admin-be/analytics
│   │   ├── kyc-api.ts  # KYC management API
│   │   │   # BE: admin-be/kyc-case
│   │   └── moderation-api.ts  # Moderation API
│   │       # BE: admin-be/moderation
│   ├── hooks/
│   │   ├── useAdminSession.ts  # JIT admin session hook
│   │   ├── useKycCases.ts  # KYC cases management
│   │   ├── useModeration.ts  # Moderation actions
│   │   └── useSystemHealth.ts  # System health monitoring
│   ├── queries/
│   │   ├── admin-mutations.ts  # Admin mutations
│   │   └── admin-queries.ts  # Admin queries
│   └── types.ts  # Admin types
│
├── analytics/
│   ├── api/
│   │   └── analytics-api.ts  # Analytics API client
│   │       # BE: Multiple services (/analytics endpoints)
│   ├── hooks/
│   │   ├── useContractAnalytics.ts  # Contract analytics
│   │   ├── useJobAnalytics.ts  # Job analytics
│   │   ├── useProfileAnalytics.ts  # Profile analytics
│   │   └── useRevenueAnalytics.ts  # Revenue analytics
│   ├── queries/
│   │   └── analytics-queries.ts  # Analytics queries
│   └── types.ts  # Analytics types
│
├── disputes/
│   ├── api/
│   │   └── disputes-api.ts  # Disputes API client
│   │       # BE: contracts-be/dispute
│   ├── hooks/
│   │   ├── useDispute.ts  # Single dispute
│   │   ├── useDisputes.ts  # Disputes list
│   │   ├── useDisputeEvidence.ts  # Evidence management
│   │   └── useDisputeResolution.ts  # Resolution actions
│   ├── queries/
│   │   ├── disputes-mutations.ts  # Dispute mutations
│   │   └── disputes-queries.ts  # Dispute queries
│   └── types.ts  # Dispute types
│
├── organizations/
│   ├── api/
│   │   ├── organizations-api.ts  # Organizations API
│   │   │   # BE: users-be/org
│   │   ├── budgets-api.ts  # Budget management API
│   │   │   # BE: financial-be/budget
│   │   └── vendors-api.ts  # Vendor management API
│   │       # BE: users-be/org (vendor subdomain)
│   ├── hooks/
│   │   ├── useBudgets.ts  # Budget management
│   │   ├── useOrganization.ts  # Organization details
│   │   ├── useTeamMembers.ts  # Team member management
│   │   └── useVendors.ts  # Vendor management
│   ├── queries/
│   │   ├── organizations-mutations.ts  # Org mutations
│   │   └── organizations-queries.ts  # Org queries
│   └── types.ts  # Organization types
│
├── subscriptions/
│   ├── api/
│   │   └── subscriptions-api.ts  # Subscriptions API
│   │       # BE: financial-be/subscription
│   ├── hooks/
│   │   ├── useEntitlements.ts  # Feature entitlements
│   │   ├── usePlans.ts  # Subscription plans
│   │   ├── useSubscription.ts  # Current subscription
│   │   └── useUsage.ts  # Usage metrics
│   ├── queries/
│   │   ├── subscriptions-mutations.ts  # Subscription mutations
│   │   └── subscriptions-queries.ts  # Subscription queries
│   └── types.ts  # Subscription types
│
└── tax/
    ├── api/
    │   └── tax-api.ts  # Tax API client
    │       # BE: financial-be/tax
    ├── hooks/
    │   ├── useTaxForms.ts  # Tax forms management
    │   ├── useTaxReports.ts  # Tax reports
    │   └── useTaxSettings.ts  # Tax settings
    ├── queries/
    │   ├── tax-mutations.ts  # Tax mutations
    │   └── tax-queries.ts  # Tax queries
    └── types.ts  # Tax types
```
