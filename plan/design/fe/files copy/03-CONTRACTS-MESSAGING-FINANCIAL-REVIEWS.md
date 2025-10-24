# Skillsier Frontend - Complete Folder Structure
## Part 3: Contracts, Messaging, Financial & Reviews Modules

> **CRITICAL**: This document contains ONLY the folder structure, filenames, and inline backend API mappings.  
> **NO CODE IMPLEMENTATIONS** are included per the strict output policy.

---

## Contracts & Work Management Module

```
apps/web/src/app/[locale]/(dashboard)/contracts/
│
├── page.tsx                                # Contracts list
│                                            # - Active contracts
│                                            # - Completed contracts
│                                            # - Filter by status
│                                            # - Sort options
│                                            # BE: contracts-be/contract
│                                            # GET /v1/contracts?status=active
│
├── active/
│   └── page.tsx                            # Active contracts only
│                                            # BE: contracts-be/contract
│                                            # GET /v1/contracts?status=active
│
├── completed/
│   └── page.tsx                            # Completed contracts
│                                            # BE: contracts-be/contract
│                                            # GET /v1/contracts?status=completed
│
├── [contractId]/
│   ├── page.tsx                            # Contract overview
│   │                                        # - Contract details
│   │                                        # - Parties involved
│   │                                        # - Budget/rate
│   │                                        # - Timeline
│   │                                        # - Status
│   │                                        # - Quick actions (message, submit work, etc.)
│   │                                        # BE: contracts-be/contract
│   │                                        # GET /v1/contracts/{contract_id}
│   │
│   ├── details/
│   │   └── page.tsx                        # Full contract details
│   │                                        # - Contract terms
│   │                                        # - Scope of work (SOW)
│   │                                        # - Payment terms
│   │                                        # - Deadlines
│   │                                        # - Special clauses
│   │                                        # BE: contracts-be/contract
│   │                                        # GET /v1/contracts/{contract_id}/details
│   │                                        # BE: contracts-be/sow
│   │                                        # GET /v1/contracts/{contract_id}/sow
│   │
│   ├── milestones/
│   │   ├── page.tsx                        # Milestones list
│   │   │                                    # - All milestones
│   │   │                                    # - Status (pending, in_progress, completed)
│   │   │                                    # - Amount & due date
│   │   │                                    # - Approval status
│   │   │                                    # BE: contracts-be/milestone
│   │   │                                    # GET /v1/contracts/{contract_id}/milestones
│   │   │
│   │   ├── create/
│   │   │   └── page.tsx                    # Create milestone (if contract allows)
│   │   │                                    # - Milestone title
│   │   │                                    # - Description
│   │   │                                    # - Amount
│   │   │                                    # - Due date
│   │   │                                    # - Deliverables
│   │   │                                    # BE: contracts-be/milestone
│   │   │                                    # POST /v1/contracts/{contract_id}/milestones
│   │   │                                    # Publishes: MilestoneCreated event
│   │   │
│   │   └── [milestoneId]/
│   │       ├── page.tsx                    # Milestone detail
│   │       │                                # - Milestone info
│   │       │                                # - Deliverables submitted
│   │       │                                # - Approval status
│   │       │                                # - Feedback
│   │       │                                # BE: contracts-be/milestone
│   │       │                                # GET /v1/milestones/{milestone_id}
│   │       │
│   │       ├── submit/
│   │       │   └── page.tsx                # Submit deliverables (freelancer)
│   │       │                                # - Upload files
│   │       │                                # - Completion notes
│   │       │                                # BE: contracts-be/deliverable
│   │       │                                # POST /v1/milestones/{milestone_id}/deliverables
│   │       │                                # BE: storage-be/uploads
│   │       │                                # Publishes: MilestoneCompleted event
│   │       │
│   │       ├── approve/
│   │       │   └── page.tsx                # Approve milestone (client)
│   │       │                                # - Review deliverables
│   │       │                                # - Accept/request changes
│   │       │                                # - Approval notes
│   │       │                                # BE: contracts-be/milestone
│   │       │                                # POST /v1/milestones/{milestone_id}/approve
│   │       │                                # Publishes: MilestoneApproved event
│   │       │                                # Triggers: Escrow release (financial-be)
│   │       │
│   │       └── dispute/
│   │           └── page.tsx                # Dispute milestone
│   │                                        # - Reason for dispute
│   │                                        # - Evidence upload
│   │                                        # BE: contracts-be/dispute
│   │                                        # POST /v1/milestones/{milestone_id}/dispute
│   │
│   ├── timesheet/
│   │   ├── page.tsx                        # Timesheet view (hourly contracts)
│   │   │                                    # - Weekly/monthly view
│   │   │                                    # - Total hours
│   │   │                                    # - Approval status
│   │   │                                    # - Submit for review
│   │   │                                    # BE: contracts-be/timesheet
│   │   │                                    # GET /v1/contracts/{contract_id}/timesheets
│   │   │
│   │   ├── submit/
│   │   │   └── page.tsx                    # Submit timesheet (freelancer)
│   │   │                                    # - Hours worked per day
│   │   │                                    # - Task descriptions
│   │   │                                    # - Billable/non-billable
│   │   │                                    # BE: contracts-be/timesheet
│   │   │                                    # POST /v1/contracts/{contract_id}/timesheets
│   │   │                                    # Publishes: TimesheetSubmitted event
│   │   │
│   │   ├── approve/
│   │   │   └── page.tsx                    # Approve timesheet (client)
│   │   │                                    # - Review hours
│   │   │                                    # - Approve/request changes
│   │   │                                    # BE: contracts-be/timesheet
│   │   │                                    # POST /v1/timesheets/{timesheet_id}/approve
│   │   │
│   │   └── [timesheetId]/
│   │       └── page.tsx                    # Timesheet detail
│   │                                        # BE: contracts-be/timesheet
│   │                                        # GET /v1/timesheets/{timesheet_id}
│   │
│   ├── work-diary/
│   │   ├── page.tsx                        # Work diary overview
│   │   │                                    # - Daily activity logs
│   │   │                                    # - Screenshots (if enabled)
│   │   │                                    # - Productivity metrics
│   │   │                                    # - Calendar view
│   │   │                                    # BE: contracts-be/work_diary
│   │   │                                    # GET /v1/contracts/{contract_id}/work-diary
│   │   │
│   │   ├── add-entry/
│   │   │   └── page.tsx                    # Add work diary entry (freelancer)
│   │   │                                    # - Date & time
│   │   │                                    # - Hours worked
│   │   │                                    # - Description
│   │   │                                    # - Upload screenshot (optional)
│   │   │                                    # BE: contracts-be/work_diary
│   │   │                                    # POST /v1/contracts/{contract_id}/work-diary/entries
│   │   │                                    # BE: storage-be/uploads
│   │   │
│   │   └── [date]/
│   │       └── page.tsx                    # Work diary for specific date
│   │                                        # BE: contracts-be/work_diary
│   │                                        # GET /v1/contracts/{contract_id}/work-diary?date={date}
│   │
│   ├── deliverables/
│   │   ├── page.tsx                        # All deliverables
│   │   │                                    # - List all submitted deliverables
│   │   │                                    # - Status (pending, approved, rejected)
│   │   │                                    # BE: contracts-be/deliverable
│   │   │                                    # GET /v1/contracts/{contract_id}/deliverables
│   │   │
│   │   └── [deliverableId]/
│   │       └── page.tsx                    # Deliverable detail
│   │                                        # - File preview
│   │                                        # - Download
│   │                                        # - Approval status
│   │                                        # - Feedback
│   │                                        # BE: contracts-be/deliverable
│   │                                        # GET /v1/deliverables/{deliverable_id}
│   │                                        # BE: storage-be
│   │                                        # GET /v1/storage/download/{file_id}
│   │
│   ├── amendments/
│   │   ├── page.tsx                        # Contract amendments list
│   │   │                                    # - Proposed amendments
│   │   │                                    # - Accepted amendments
│   │   │                                    # BE: contracts-be/amendment
│   │   │                                    # GET /v1/contracts/{contract_id}/amendments
│   │   │
│   │   ├── propose/
│   │   │   └── page.tsx                    # Propose amendment
│   │   │                                    # - Change description
│   │   │                                    # - Updated terms
│   │   │                                    # - Reason
│   │   │                                    # BE: contracts-be/amendment
│   │   │                                    # POST /v1/contracts/{contract_id}/amendments
│   │   │
│   │   └── [amendmentId]/
│   │       ├── page.tsx                    # Amendment detail
│   │       │                                # BE: contracts-be/amendment
│   │       │                                # GET /v1/amendments/{amendment_id}
│   │       │
│   │       └── approve/
│   │           └── page.tsx                # Approve/reject amendment
│   │                                        # BE: contracts-be/amendment
│   │                                        # POST /v1/amendments/{amendment_id}/approve
│   │                                        # POST /v1/amendments/{amendment_id}/reject
│   │
│   ├── disputes/
│   │   ├── page.tsx                        # Disputes list
│   │   │                                    # - Active disputes
│   │   │                                    # - Resolved disputes
│   │   │                                    # BE: contracts-be/dispute
│   │   │                                    # GET /v1/contracts/{contract_id}/disputes
│   │   │
│   │   ├── open/
│   │   │   └── page.tsx                    # Open a dispute
│   │   │                                    # - Dispute reason
│   │   │                                    # - Description
│   │   │                                    # - Evidence upload
│   │   │                                    # - Desired resolution
│   │   │                                    # BE: contracts-be/dispute
│   │   │                                    # POST /v1/contracts/{contract_id}/disputes
│   │   │                                    # BE: storage-be/uploads
│   │   │                                    # Publishes: DisputeOpened event
│   │   │
│   │   └── [disputeId]/
│   │       ├── page.tsx                    # Dispute detail
│   │       │                                # - Dispute timeline
│   │       │                                # - Messages/responses
│   │       │                                # - Evidence
│   │       │                                # - Admin notes (if assigned)
│   │       │                                # - Resolution status
│   │       │                                # BE: contracts-be/dispute
│   │       │                                # GET /v1/disputes/{dispute_id}
│   │       │
│   │       ├── respond/
│   │       │   └── page.tsx                # Respond to dispute
│   │       │                                # - Response message
│   │       │                                # - Additional evidence
│   │       │                                # BE: contracts-be/dispute
│   │       │                                # POST /v1/disputes/{dispute_id}/responses
│   │       │
│   │       └── escalate/
│   │           └── page.tsx                # Escalate to admin/mediation
│   │                                        # BE: contracts-be/dispute
│   │                                        # POST /v1/disputes/{dispute_id}/escalate
│   │                                        # BE: admin-be
│   │                                        # Creates mediation case
│   │
│   ├── payments/
│   │   └── page.tsx                        # Contract payments
│   │                                        # - Payment schedule
│   │                                        # - Escrow status
│   │                                        # - Released payments
│   │                                        # - Pending payments
│   │                                        # BE: financial-be/escrow
│   │                                        # GET /v1/contracts/{contract_id}/escrow
│   │                                        # BE: financial-be/payment
│   │                                        # GET /v1/contracts/{contract_id}/payments
│   │
│   ├── messages/
│   │   └── page.tsx                        # Contract-specific messages
│   │                                        # - Threaded conversation
│   │                                        # - File sharing
│   │                                        # - Quick links to milestones/deliverables
│   │                                        # BE: communications-be/conversations
│   │                                        # GET /v1/contracts/{contract_id}/conversation
│   │
│   ├── feedback/
│   │   └── page.tsx                        # Ongoing feedback
│   │                                        # - Mid-contract feedback
│   │                                        # - Performance notes
│   │                                        # BE: contracts-be/feedback
│   │                                        # GET /v1/contracts/{contract_id}/feedback
│   │                                        # POST /v1/contracts/{contract_id}/feedback
│   │
│   ├── pause/
│   │   └── page.tsx                        # Pause contract
│   │                                        # - Reason
│   │                                        # - Expected resume date
│   │                                        # - Notify other party
│   │                                        # BE: contracts-be/contract
│   │                                        # POST /v1/contracts/{contract_id}/pause
│   │                                        # Publishes: ContractPaused event
│   │
│   ├── terminate/
│   │   └── page.tsx                        # Terminate contract
│   │                                        # - Termination reason
│   │                                        # - Early termination terms
│   │                                        # - Final deliverables
│   │                                        # - Escrow settlement
│   │                                        # BE: contracts-be/termination
│   │                                        # POST /v1/contracts/{contract_id}/terminate
│   │                                        # Publishes: ContractTerminated event
│   │
│   └── complete/
│       └── page.tsx                        # Complete contract
│                                            # - Confirm all deliverables received
│                                            # - Leave review
│                                            # - Final payment release
│                                            # BE: contracts-be/contract
│                                            # POST /v1/contracts/{contract_id}/complete
│                                            # Publishes: ContractCompleted event
│                                            # Redirects to: /reviews/create
│
├── templates/
│   ├── page.tsx                            # Contract templates (for recurring work)
│   │                                        # BE: contracts-be/template
│   │                                        # GET /v1/contract-templates
│   │
│   ├── create/
│   │   └── page.tsx                        # Create contract template
│   │                                        # BE: contracts-be/template
│   │                                        # POST /v1/contract-templates
│   │
│   └── [templateId]/
│       └── edit/
│           └── page.tsx                    # Edit template
│                                            # BE: contracts-be/template
│                                            # PUT /v1/contract-templates/{template_id}
│
└── recurring/
    ├── page.tsx                            # Recurring contracts
    │                                        # - List recurring contracts
    │                                        # - Renewal schedule
    │                                        # BE: contracts-be/recurring
    │                                        # GET /v1/contracts/recurring
    │
    └── [contractId]/
        └── renew/
            └── page.tsx                    # Renew recurring contract
                                            # BE: contracts-be/recurring
                                            # POST /v1/contracts/{contract_id}/renew
```

---

## Messaging Module

```
apps/web/src/app/[locale]/(dashboard)/messages/
│
├── page.tsx                                # Messages inbox
│                                            # - Conversation list
│                                            # - Unread indicator
│                                            # - Search conversations
│                                            # - Filter (all, unread, archived)
│                                            # - Real-time updates via WebSocket
│                                            # BE: communications-be/conversations
│                                            # GET /v1/conversations
│                                            # WebSocket: ws://communications-be/v1/realtime
│
├── [conversationId]/
│   ├── page.tsx                            # Conversation thread
│   │                                        # - Message history
│   │                                        # - Real-time new messages
│   │                                        # - Message composer
│   │                                        # - File attachments
│   │                                        # - Typing indicators
│   │                                        # - Read receipts
│   │                                        # - Quick actions (block, report)
│   │                                        # BE: communications-be/messages
│   │                                        # GET /v1/conversations/{conversation_id}/messages
│   │                                        # POST /v1/messages
│   │                                        # WebSocket: Real-time message delivery
│   │                                        # Publishes: MessageSent event
│   │
│   ├── info/
│   │   └── page.tsx                        # Conversation info
│   │                                        # - Participants
│   │                                        # - Related job/contract
│   │                                        # - Shared files
│   │                                        # - Search in conversation
│   │                                        # BE: communications-be/conversations
│   │                                        # GET /v1/conversations/{conversation_id}/info
│   │
│   └── archive/
│       └── page.tsx                        # Archive conversation
│                                            # BE: communications-be/conversations
│                                            # POST /v1/conversations/{conversation_id}/archive
│
├── archived/
│   └── page.tsx                            # Archived conversations
│                                            # BE: communications-be/conversations
│                                            # GET /v1/conversations?archived=true
│
└── new/
    └── page.tsx                            # Start new conversation
                                            # - Select recipient (search users)
                                            # - Initial message
                                            # BE: communications-be/conversations
                                            # POST /v1/conversations
```

---

## Notifications Module

```
apps/web/src/app/[locale]/(dashboard)/notifications/
│
├── page.tsx                                # Notifications center
│                                            # - All notifications list
│                                            # - Unread indicator
│                                            # - Mark all as read
│                                            # - Filter by type
│                                            # - Real-time updates
│                                            # BE: communications-be/notifications
│                                            # GET /v1/notifications
│                                            # POST /v1/notifications/read-all
│                                            # WebSocket: ws://communications-be/v1/notifications
│
├── unread/
│   └── page.tsx                            # Unread notifications only
│                                            # BE: communications-be/notifications
│                                            # GET /v1/notifications?unread=true
│
├── [notificationId]/
│   └── page.tsx                            # Notification detail (redirects to relevant page)
│                                            # - Mark as read
│                                            # - Navigate to related entity
│                                            # BE: communications-be/notifications
│                                            # POST /v1/notifications/{notif_id}/read
│
└── settings/
    └── page.tsx                            # Notification settings
                                            # - Email notifications
                                            # - Push notifications
                                            # - In-app notifications
                                            # - Notification preferences by type
                                            # - Frequency settings
                                            # - Quiet hours
                                            # BE: communications-be/preferences
                                            # GET /v1/notifications/preferences
                                            # PUT /v1/notifications/preferences
```

---

## Financial Management Module

```
apps/web/src/app/[locale]/(dashboard)/financials/
│
├── page.tsx                                # Financial overview
│                                            # - Wallet balance
│                                            # - Pending payments
│                                            # - Recent transactions
│                                            # - Earnings chart (freelancer)
│                                            # - Spending chart (client)
│                                            # BE: financial-be/wallet
│                                            # GET /v1/wallet/balance
│                                            # BE: financial-be/transaction
│                                            # GET /v1/transactions/recent
│
├── wallet/
│   ├── page.tsx                            # Wallet details
│   │                                        # - Available balance
│   │                                        # - Pending balance
│   │                                        # - Escrow balance
│   │                                        # - Add funds button (client)
│   │                                        # - Withdraw button (freelancer)
│   │                                        # - Transaction history
│   │                                        # BE: financial-be/wallet
│   │                                        # GET /v1/wallet
│   │
│   ├── add-funds/
│   │   └── page.tsx                        # Add funds (client)
│   │                                        # - Amount input
│   │                                        # - Payment method selection
│   │                                        # - Payment processing
│   │                                        # BE: financial-be/wallet
│   │                                        # POST /v1/wallet/add-funds
│   │                                        # BE: financial-be/payment
│   │                                        # POST /v1/payments (Stripe/PayPal integration)
│   │
│   └── withdraw/
│       └── page.tsx                        # Withdraw funds (freelancer)
│                                            # - Amount input
│                                            # - Payout method selection
│                                            # - Tax information
│                                            # - Withdrawal fees
│                                            # BE: financial-be/payout
│                                            # POST /v1/payouts/request
│                                            # Publishes: PayoutRequested event
│
├── transactions/
│   ├── page.tsx                            # Transaction history
│   │                                        # - All transactions
│   │                                        # - Filter by type (payment, payout, refund, etc.)
│   │                                        # - Filter by date range
│   │                                        # - Search by description
│   │                                        # - Export to CSV
│   │                                        # BE: financial-be/transaction
│   │                                        # GET /v1/transactions?filters={...}
│   │
│   └── [transactionId]/
│       └── page.tsx                        # Transaction detail
│                                            # - Full transaction info
│                                            # - Related contract/job
│                                            # - Receipt download
│                                            # BE: financial-be/transaction
│                                            # GET /v1/transactions/{transaction_id}
│
├── invoices/
│   ├── page.tsx                            # Invoices list
│   │                                        # - Sent invoices (freelancer)
│   │                                        # - Received invoices (client)
│   │                                        # - Filter by status (paid, pending, overdue)
│   │                                        # BE: financial-be/invoice
│   │                                        # GET /v1/invoices
│   │
│   ├── [invoiceId]/
│   │   ├── page.tsx                        # Invoice detail
│   │   │                                    # - Invoice information
│   │   │                                    # - Line items
│   │   │                                    # - Tax details
│   │   │                                    # - Payment status
│   │   │                                    # - Download PDF
│   │   │                                    # BE: financial-be/invoice
│   │   │                                    # GET /v1/invoices/{invoice_id}
│   │   │                                    # GET /v1/invoices/{invoice_id}/pdf
│   │   │
│   │   └── pay/
│   │       └── page.tsx                    # Pay invoice (client)
│   │                                        # - Invoice summary
│   │                                        # - Payment method selection
│   │                                        # - Process payment
│   │                                        # BE: financial-be/payment
│   │                                        # POST /v1/invoices/{invoice_id}/pay
│   │
│   └── create/
│       └── page.tsx                        # Create invoice (manual invoicing)
│                                            # - Client selection
│                                            # - Line items
│                                            # - Tax settings
│                                            # - Due date
│                                            # - Notes
│                                            # BE: financial-be/invoice
│                                            # POST /v1/invoices
│
├── payment-methods/
│   ├── page.tsx                            # Payment methods list
│   │                                        # - Saved credit cards
│   │                                        # - PayPal accounts
│   │                                        # - Bank accounts
│   │                                        # - Default payment method
│   │                                        # BE: financial-be/payment_method
│   │                                        # GET /v1/payment-methods
│   │
│   ├── add/
│   │   └── page.tsx                        # Add payment method
│   │                                        # - Card details (Stripe Elements)
│   │                                        # - PayPal connection
│   │                                        # - Bank account (ACH)
│   │                                        # - Set as default
│   │                                        # BE: financial-be/payment_method
│   │                                        # POST /v1/payment-methods
│   │
│   └── [methodId]/
│       ├── page.tsx                        # Payment method detail
│       │                                    # BE: financial-be/payment_method
│       │                                    # GET /v1/payment-methods/{method_id}
│       │
│       └── delete/
│           └── page.tsx                    # Delete payment method
│                                            # BE: financial-be/payment_method
│                                            # DELETE /v1/payment-methods/{method_id}
│
├── payout-methods/
│   ├── page.tsx                            # Payout methods list (freelancer)
│   │                                        # - Bank accounts
│   │                                        # - PayPal
│   │                                        # - Wire transfer details
│   │                                        # BE: financial-be/payout_method
│   │                                        # GET /v1/payout-methods
│   │
│   ├── add/
│   │   └── page.tsx                        # Add payout method
│   │                                        # - Bank account details
│   │                                        # - PayPal email
│   │                                        # - Wire transfer info
│   │                                        # - Tax forms (W-9, W-8BEN)
│   │                                        # BE: financial-be/payout_method
│   │                                        # POST /v1/payout-methods
│   │
│   └── [methodId]/
│       └── page.tsx                        # Payout method detail
│                                            # BE: financial-be/payout_method
│                                            # GET /v1/payout-methods/{method_id}
│                                            # DELETE /v1/payout-methods/{method_id}
│
├── tax/
│   ├── page.tsx                            # Tax information
│   │                                        # - Tax forms
│   │                                        # - Tax ID
│   │                                        # - VAT/GST number
│   │                                        # - Tax residency
│   │                                        # BE: financial-be/tax
│   │                                        # GET /v1/tax/info
│   │
│   ├── forms/
│   │   ├── page.tsx                        # Tax forms list
│   │   │                                    # - W-9, 1099, W-8BEN, etc.
│   │   │                                    # - Download forms
│   │   │                                    # BE: financial-be/tax
│   │   │                                    # GET /v1/tax/forms
│   │   │
│   │   └── upload/
│   │       └── page.tsx                    # Upload tax form
│   │                                        # BE: financial-be/tax
│   │                                        # POST /v1/tax/forms
│   │                                        # BE: storage-be/uploads
│   │
│   └── settings/
│       └── page.tsx                        # Tax settings
│                                            # - Tax information
│                                            # - VAT reverse charge
│                                            # - Tax exemptions
│                                            # BE: financial-be/tax
│                                            # PUT /v1/tax/settings
│
├── reports/
│   ├── page.tsx                            # Financial reports
│   │                                        # - Earnings report (freelancer)
│   │                                        # - Spending report (client)
│   │                                        # - Tax report
│   │                                        # - Date range selection
│   │                                        # - Export options
│   │                                        # BE: financial-be/reports
│   │                                        # GET /v1/reports/earnings
│   │                                        # GET /v1/reports/spending
│   │
│   ├── earnings/
│   │   └── page.tsx                        # Detailed earnings report
│   │                                        # - By project
│   │                                        # - By client
│   │                                        # - By time period
│   │                                        # BE: financial-be/reports
│   │                                        # GET /v1/reports/earnings/detailed
│   │
│   └── spending/
│       └── page.tsx                        # Detailed spending report
│                                            # - By project
│                                            # - By freelancer
│                                            # - By category
│                                            # BE: financial-be/reports
│                                            # GET /v1/reports/spending/detailed
│
└── escrow/
    ├── page.tsx                            # Escrow overview
    │                                        # - Active escrow accounts
    │                                        # - Total amount in escrow
    │                                        # - Pending releases
    │                                        # BE: financial-be/escrow
    │                                        # GET /v1/escrow
    │
    └── [escrowId]/
        └── page.tsx                        # Escrow detail
                                            # - Related contract
                                            # - Amount held
                                            # - Release schedule
                                            # - Transaction history
                                            # BE: financial-be/escrow
                                            # GET /v1/escrow/{escrow_id}
```

---

## Reviews & Ratings Module

```
apps/web/src/app/[locale]/(dashboard)/reviews/
│
├── page.tsx                                # Reviews overview
│                                            # - Reviews received
│                                            # - Reviews given
│                                            # - Overall rating stats
│                                            # - Badges earned
│                                            # BE: reviews-be/reviews
│                                            # GET /v1/reviews?user_id={current_user}
│                                            # BE: reviews-be/stats
│                                            # GET /v1/reviews/stats?user_id={current_user}
│
├── received/
│   ├── page.tsx                            # Reviews received list
│   │                                        # - All reviews received
│   │                                        # - Filter by rating
│   │                                        # - Filter by contract
│   │                                        # BE: reviews-be/reviews
│   │                                        # GET /v1/reviews?user_id={current_user}&type=received
│   │
│   └── [reviewId]/
│       ├── page.tsx                        # Review detail
│       │                                    # - Full review content
│       │                                    # - Reviewer info
│       │                                    # - Related contract
│       │                                    # - Response option
│       │                                    # BE: reviews-be/reviews
│       │                                    # GET /v1/reviews/{review_id}
│       │
│       └── respond/
│           └── page.tsx                    # Respond to review
│                                            # - Public response
│                                            # - Character limit
│                                            # BE: reviews-be/reviews
│                                            # POST /v1/reviews/{review_id}/respond
│                                            # Publishes: ReviewResponded event
│
├── given/
│   └── page.tsx                            # Reviews given list
│                                            # BE: reviews-be/reviews
│                                            # GET /v1/reviews?user_id={current_user}&type=given
│
├── create/
│   ├── [contractId]/
│   │   └── page.tsx                        # Create review
│   │                                        # - Rating (1-5 stars)
│   │                                        # - Multiple criteria ratings:
│   │                                        #   - Quality of work
│   │                                        #   - Communication
│   │                                        #   - Professionalism
│   │                                        #   - Deadlines
│   │                                        # - Written feedback (required)
│   │                                        # - Recommend to others?
│   │                                        # - Skills demonstrated
│   │                                        # - Make public/private
│   │                                        # BE: reviews-be/reviews
│   │                                        # POST /v1/reviews
│   │                                        # Body: { contract_id, rating, criteria_ratings, feedback, ... }
│   │                                        # Publishes: ReviewSubmitted event
│   │                                        # Triggers: Reputation update (users-be)
│   │
│   └── pending/
│       └── page.tsx                        # Pending reviews
│                                            # - Contracts awaiting review
│                                            # - Reminders
│                                            # BE: reviews-be/reviews
│                                            # GET /v1/reviews/pending
│
├── badges/
│   └── page.tsx                            # Badges & achievements
│                                            # - Earned badges
│                                            # - Badge criteria
│                                            # - Progress towards badges
│                                            # BE: reviews-be/badges
│                                            # GET /v1/reviews/badges?user_id={current_user}
│
└── stats/
    └── page.tsx                            # Detailed statistics
                                            # - Rating breakdown
                                            # - Review trends over time
                                            # - Category-specific ratings
                                            # - Comparison to platform average
                                            # BE: reviews-be/stats
                                            # GET /v1/reviews/stats/detailed?user_id={current_user}
```

---

## Query Key Patterns (TanStack Query)

### Contracts
```typescript
['contracts', 'list', status]               // Contract list by status
['contracts', 'detail', contractId]          // Contract detail
['contracts', 'active']                      // Active contracts
['contracts', 'completed']                   // Completed contracts
['contracts', 'milestones', contractId]      // Contract milestones
['contracts', 'timesheet', contractId]       // Contract timesheets
['contracts', 'work-diary', contractId]      // Work diary
['contracts', 'deliverables', contractId]    // Deliverables
['contracts', 'amendments', contractId]      // Amendments
['contracts', 'disputes', contractId]        // Disputes
['contracts', 'payments', contractId]        // Contract payments
```

### Messages
```typescript
['conversations', 'list']                    // Conversation list
['conversations', 'detail', conversationId]  // Conversation detail
['messages', conversationId]                 // Messages in conversation
['messages', 'unread-count']                 // Unread message count
```

### Notifications
```typescript
['notifications', 'list']                    // All notifications
['notifications', 'unread']                  // Unread notifications
['notifications', 'unread-count']            // Unread count
```

### Financial
```typescript
['wallet', 'balance']                        // Wallet balance
['wallet', 'detail']                         // Wallet details
['transactions', 'list', filters]            // Transaction list
['transactions', 'detail', transactionId]    // Transaction detail
['invoices', 'list', filters]                // Invoice list
['invoices', 'detail', invoiceId]            // Invoice detail
['payment-methods', 'list']                  // Payment methods
['payout-methods', 'list']                   // Payout methods
['tax', 'info']                              // Tax information
['tax', 'forms']                             // Tax forms
['reports', 'earnings', dateRange]           // Earnings report
['reports', 'spending', dateRange]           // Spending report
['escrow', 'list']                           // Escrow accounts
['escrow', 'detail', escrowId]               // Escrow detail
```

### Reviews
```typescript
['reviews', 'list', userId]                  // User reviews
['reviews', 'received', userId]              // Reviews received
['reviews', 'given', userId]                 // Reviews given
['reviews', 'detail', reviewId]              // Review detail
['reviews', 'pending']                       // Pending reviews
['reviews', 'stats', userId]                 // Review statistics
['reviews', 'badges', userId]                // User badges
```

---

**End of Part 3**

**Continue to Part 4 for:**
- Settings & Preferences
- Subscription Management
- Admin Panel
- Shared Components
- Features Structure
