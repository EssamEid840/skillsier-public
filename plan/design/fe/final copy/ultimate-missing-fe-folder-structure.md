# Ultimate Missing Frontend Folder Structure for Skillsier Application
## Final Comprehensive Coverage of All Remaining Requirements

> **SCOPE**: This document contains ONLY the folder structure elements that are:
> 1. Required by `fe-folder-structure-prompt.md`
> 2. NOT present in any of the previous documents:
>    - `combined-folder-structure.md`
>    - `missing-fe-folder-structure.md`
>    - `additional-missing-fe-folder-structure.md`
>    - `remaining-missing-fe-folder-structure.md`
>    - `final-missing-fe-folder-structure.md`
>    - `final-comprehensive-missing-fe-structure.md`

---

## I. Critical Dashboard Routes - Deep Feature Coverage

### 1. Proposals - Advanced Negotiation & Version Control

```
apps/web/src/app/[locale]/(dashboard)/
│
├── proposals/
│   ├── [proposalId]/
│   │   ├── compare/
│   │   │   └── [compareWith]/
│   │   │       └── page.tsx  # Compare two proposals side-by-side
│   │   │           # - Terms comparison
│   │   │           # - Pricing comparison
│   │   │           # - Deliverables comparison
│   │   │           # - Timeline comparison
│   │   │           # BE: proposals-be/proposal
│   │   │           # GET /v1/proposals/{proposal_id}
│   │   │           # GET /v1/proposals/{compare_with}
│   │   │
│   │   ├── negotiation/
│   │   │   ├── counter-offer/
│   │   │   │   └── page.tsx  # Create counter-offer
│   │   │   │       # - Modify terms
│   │   │   │       # - Adjust pricing
│   │   │   │       # - Add conditions
│   │   │   │       # BE: proposals-be/proposal
│   │   │   │       # POST /v1/proposals/{proposal_id}/counter-offer
│   │   │   │
│   │   │   ├── history/
│   │   │   │   └── page.tsx  # Negotiation timeline
│   │   │   │       # - All offers/counter-offers
│   │   │   │       # - Changes tracking
│   │   │   │       # - Decision points
│   │   │   │       # BE: proposals-be/proposal
│   │   │   │       # GET /v1/proposals/{proposal_id}/negotiation-history
│   │   │   │
│   │   │   └── page.tsx  # Active negotiation dashboard
│   │   │       # - Current offer status
│   │   │       # - Next steps
│   │   │       # - Quick actions
│   │   │       # BE: proposals-be/proposal
│   │   │       # GET /v1/proposals/{proposal_id}/negotiation-status
│   │   │
│   │   ├── milestones/
│   │   │   ├── edit/
│   │   │   │   └── page.tsx  # Edit proposal milestones
│   │   │   │       # - Add/remove milestones
│   │   │   │       # - Adjust payment schedule
│   │   │   │       # BE: proposals-be/milestone
│   │   │   │       # PUT /v1/proposals/{proposal_id}/milestones
│   │   │   │
│   │   │   └── page.tsx  # Milestones overview
│   │   │       # - Timeline view
│   │   │       # - Payment breakdown
│   │   │       # BE: proposals-be/milestone
│   │   │       # GET /v1/proposals/{proposal_id}/milestones
│   │   │
│   │   ├── questions/
│   │   │   ├── [questionId]/
│   │   │   │   └── page.tsx  # Answer single question
│   │   │   │       # BE: proposals-be/question
│   │   │   │       # GET /v1/proposals/{proposal_id}/questions/{question_id}
│   │   │   │       # POST /v1/proposals/{proposal_id}/questions/{question_id}/answer
│   │   │   │
│   │   │   └── page.tsx  # All proposal questions
│   │   │       # - Q&A thread
│   │   │       # - Ask new question
│   │   │       # BE: proposals-be/question
│   │   │       # GET /v1/proposals/{proposal_id}/questions
│   │   │       # POST /v1/proposals/{proposal_id}/questions
│   │   │
│   │   └── collaborators/
│   │       └── page.tsx  # Team collaboration on proposal
│   │           # - Invite team members
│   │           # - Internal notes
│   │           # - Review assignments
│   │           # BE: proposals-be/proposal, users-be/team
│   │           # GET /v1/proposals/{proposal_id}/collaborators
│   │           # POST /v1/proposals/{proposal_id}/collaborators
│   │
│   ├── insights/
│   │   ├── win-rate/
│   │   │   └── page.tsx  # Win rate analytics
│   │   │       # - Success by job type
│   │   │       # - Success by budget range
│   │   │       # - Improvement suggestions
│   │   │       # BE: proposals-be/analytics
│   │   │       # GET /v1/proposals/insights/win-rate
│   │   │
│   │   ├── pricing-analysis/
│   │   │   └── page.tsx  # Pricing competitiveness
│   │   │       # - Market rate comparison
│   │   │       # - Your pricing vs. competitors
│   │   │       # - Pricing recommendations
│   │   │       # BE: proposals-be/analytics, search-be/market-data
│   │   │       # GET /v1/proposals/insights/pricing
│   │   │       # GET /v1/market/rates
│   │   │
│   │   └── response-time/
│   │       └── page.tsx  # Response time analytics
│   │           # - Average response time
│   │           # - Impact on success rate
│   │           # - Optimization tips
│   │           # BE: proposals-be/analytics
│   │           # GET /v1/proposals/insights/response-time
│   │
│   └── portfolio-showcases/
│       ├── [showcaseId]/
│       │   ├── edit/
│       │   │   └── page.tsx  # Edit showcase
│       │   │       # BE: proposals-be/showcase
│       │   │       # PUT /v1/proposals/showcases/{showcase_id}
│       │   │
│       │   └── page.tsx  # Showcase detail
│       │       # BE: proposals-be/showcase
│       │       # GET /v1/proposals/showcases/{showcase_id}
│       │
│       └── page.tsx  # Manage showcases
│           # - Create showcase
│           # - Link to proposals
│           # BE: proposals-be/showcase
│           # GET /v1/proposals/showcases
│           # POST /v1/proposals/showcases
```

### 2. Contracts - Comprehensive Work Management

```
apps/web/src/app/[locale]/(dashboard)/
│
├── contracts/
│   ├── [contractId]/
│   │   ├── work-diary/
│   │   │   ├── bulk-entry/
│   │   │   │   └── page.tsx  # Bulk time entry
│   │   │   │       # - Multiple entries at once
│   │   │   │       # - Copy from previous week
│   │   │   │       # BE: contracts-be/workdiary
│   │   │   │       # POST /v1/contracts/{contract_id}/workdiary/bulk
│   │   │   │
│   │   │   ├── screenshots/
│   │   │   │   └── page.tsx  # Screenshot gallery
│   │   │   │       # - View all screenshots
│   │   │   │       # - Delete/flag screenshots
│   │   │   │       # BE: contracts-be/workdiary, storage-be/asset
│   │   │   │       # GET /v1/contracts/{contract_id}/workdiary/screenshots
│   │   │   │
│   │   │   ├── activity-levels/
│   │   │   │   └── page.tsx  # Activity level tracking
│   │   │   │       # - Keyboard/mouse metrics
│   │   │   │       # - Focus time
│   │   │   │       # - Idle time detection
│   │   │   │       # BE: contracts-be/workdiary
│   │   │   │       # GET /v1/contracts/{contract_id}/workdiary/activity
│   │   │   │
│   │   │   └── corrections/
│   │   │       └── page.tsx  # Time entry corrections
│   │   │           # - Request corrections
│   │   │           # - Approve corrections
│   │   │           # BE: contracts-be/workdiary
│   │   │           # POST /v1/contracts/{contract_id}/workdiary/corrections
│   │   │           # GET /v1/contracts/{contract_id}/workdiary/corrections
│   │   │
│   │   ├── milestones/
│   │   │   ├── [milestoneId]/
│   │   │   │   ├── submission/
│   │   │   │   │   └── page.tsx  # Submit milestone
│   │   │   │   │       # - Upload deliverables
│   │   │   │   │       # - Add notes
│   │   │   │   │       # BE: contracts-be/milestone, storage-be/asset
│   │   │   │   │       # POST /v1/contracts/{contract_id}/milestones/{milestone_id}/submit
│   │   │   │   │
│   │   │   │   ├── review/
│   │   │   │   │   └── page.tsx  # Review milestone
│   │   │   │   │       # - Approve/reject
│   │   │   │   │       # - Request changes
│   │   │   │   │       # BE: contracts-be/milestone
│   │   │   │   │       # POST /v1/contracts/{contract_id}/milestones/{milestone_id}/review
│   │   │   │   │
│   │   │   │   └── revisions/
│   │   │   │       └── page.tsx  # Milestone revision history
│   │   │   │           # BE: contracts-be/milestone
│   │   │   │           # GET /v1/contracts/{contract_id}/milestones/{milestone_id}/revisions
│   │   │   │
│   │   │   └── reorder/
│   │   │       └── page.tsx  # Reorder milestones
│   │   │           # - Drag and drop
│   │   │           # - Adjust dependencies
│   │   │           # BE: contracts-be/milestone
│   │   │           # PUT /v1/contracts/{contract_id}/milestones/reorder
│   │   │
│   │   ├── deliverables/
│   │   │   ├── [deliverableId]/
│   │   │   │   ├── versions/
│   │   │   │   │   ├── [versionId]/
│   │   │   │   │   │   └── page.tsx  # Specific version detail
│   │   │   │   │   │       # BE: contracts-be/deliverable
│   │   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/versions/{version_id}
│   │   │   │   │   │
│   │   │   │   │   └── compare/
│   │   │   │   │       └── page.tsx  # Compare versions
│   │   │   │   │           # BE: contracts-be/deliverable
│   │   │   │   │           # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/versions/compare
│   │   │   │   │
│   │   │   │   ├── feedback/
│   │   │   │   │   └── page.tsx  # Deliverable feedback
│   │   │   │   │       # - Annotate files
│   │   │   │   │       # - Comment threads
│   │   │   │   │       # BE: contracts-be/deliverable, communications-be/comment
│   │   │   │   │       # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/feedback
│   │   │   │   │       # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/feedback
│   │   │   │   │
│   │   │   │   └── approvals/
│   │   │   │       └── page.tsx  # Approval workflow
│   │   │   │           # - Multi-step approvals
│   │   │   │           # - Stakeholder sign-offs
│   │   │   │           # BE: contracts-be/deliverable
│   │   │   │           # GET /v1/contracts/{contract_id}/deliverables/{deliverable_id}/approvals
│   │   │   │           # POST /v1/contracts/{contract_id}/deliverables/{deliverable_id}/approve
│   │   │   │
│   │   │   └── bulk-operations/
│   │   │       └── page.tsx  # Bulk deliverable actions
│   │   │           # BE: contracts-be/deliverable
│   │   │           # POST /v1/contracts/{contract_id}/deliverables/bulk
│   │   │
│   │   ├── change-requests/
│   │   │   ├── [requestId]/
│   │   │   │   ├── negotiate/
│   │   │   │   │   └── page.tsx  # Negotiate change request
│   │   │   │   │       # - Counter proposals
│   │   │   │   │       # - Impact analysis
│   │   │   │   │       # BE: contracts-be/change-request
│   │   │   │   │       # POST /v1/contracts/{contract_id}/change-requests/{request_id}/negotiate
│   │   │   │   │
│   │   │   │   ├── impact-analysis/
│   │   │   │   │   └── page.tsx  # Change impact analysis
│   │   │   │   │       # - Cost impact
│   │   │   │   │       # - Timeline impact
│   │   │   │   │       # - Resource impact
│   │   │   │   │       # BE: contracts-be/change-request
│   │   │   │   │       # GET /v1/contracts/{contract_id}/change-requests/{request_id}/impact
│   │   │   │   │
│   │   │   │   └── approval-chain/
│   │   │   │       └── page.tsx  # Change request approvals
│   │   │   │           # - Approval hierarchy
│   │   │   │           # - Pending approvals
│   │   │   │           # BE: contracts-be/change-request
│   │   │   │           # GET /v1/contracts/{contract_id}/change-requests/{request_id}/approvals
│   │   │   │
│   │   │   └── templates/
│   │   │       └── page.tsx  # Change request templates
│   │   │           # BE: contracts-be/change-request
│   │   │           # GET /v1/contracts/change-requests/templates
│   │   │
│   │   ├── compliance/
│   │   │   ├── documents/
│   │   │   │   └── page.tsx  # Compliance documents
│   │   │   │       # - Certifications
│   │   │   │       # - Insurance documents
│   │   │   │       # - Legal requirements
│   │   │   │       # BE: contracts-be/compliance, storage-be/asset
│   │   │   │       # GET /v1/contracts/{contract_id}/compliance/documents
│   │   │   │       # POST /v1/contracts/{contract_id}/compliance/documents
│   │   │   │
│   │   │   ├── audits/
│   │   │   │   ├── [auditId]/
│   │   │   │   │   └── page.tsx  # Audit detail
│   │   │   │   │       # BE: contracts-be/compliance
│   │   │   │   │       # GET /v1/contracts/{contract_id}/compliance/audits/{audit_id}
│   │   │   │   │
│   │   │   │   └── page.tsx  # Audits list
│   │   │   │       # BE: contracts-be/compliance
│   │   │   │       # GET /v1/contracts/{contract_id}/compliance/audits
│   │   │   │
│   │   │   └── reports/
│   │   │       └── page.tsx  # Compliance reports
│   │   │           # - Generate reports
│   │   │           # - Export for auditors
│   │   │           # BE: contracts-be/compliance
│   │   │           # GET /v1/contracts/{contract_id}/compliance/reports
│   │   │           # POST /v1/contracts/{contract_id}/compliance/reports/generate
│   │   │
│   │   ├── risks/
│   │   │   ├── register/
│   │   │   │   └── page.tsx  # Risk register
│   │   │   │       # - Identify risks
│   │   │   │       # - Risk assessment
│   │   │   │       # - Mitigation plans
│   │   │   │       # BE: contracts-be/risk
│   │   │   │       # GET /v1/contracts/{contract_id}/risks
│   │   │   │       # POST /v1/contracts/{contract_id}/risks
│   │   │   │
│   │   │   ├── monitoring/
│   │   │   │   └── page.tsx  # Risk monitoring
│   │   │   │       # - Active risks
│   │   │   │       # - Risk trends
│   │   │   │       # - Alerts
│   │   │   │       # BE: contracts-be/risk
│   │   │   │       # GET /v1/contracts/{contract_id}/risks/monitoring
│   │   │   │
│   │   │   └── reports/
│   │   │       └── page.tsx  # Risk reports
│   │   │           # BE: contracts-be/risk
│   │   │           # GET /v1/contracts/{contract_id}/risks/reports
│   │   │
│   │   ├── quality/
│   │   │   ├── metrics/
│   │   │   │   └── page.tsx  # Quality metrics
│   │   │   │       # - Code quality
│   │   │   │       # - Deliverable quality
│   │   │   │       # - Process quality
│   │   │   │       # BE: contracts-be/quality
│   │   │   │       # GET /v1/contracts/{contract_id}/quality/metrics
│   │   │   │
│   │   │   ├── reviews/
│   │   │   │   └── page.tsx  # Quality reviews
│   │   │   │       # - Schedule reviews
│   │   │   │       # - Review results
│   │   │   │       # BE: contracts-be/quality
│   │   │   │       # GET /v1/contracts/{contract_id}/quality/reviews
│   │   │   │       # POST /v1/contracts/{contract_id}/quality/reviews
│   │   │   │
│   │   │   └── standards/
│   │   │       └── page.tsx  # Quality standards
│   │   │           # - Set standards
│   │   │           # - Compliance tracking
│   │   │           # BE: contracts-be/quality
│   │   │           # GET /v1/contracts/{contract_id}/quality/standards
│   │   │
│   │   └── knowledge-transfer/
│   │       ├── sessions/
│   │       │   ├── [sessionId]/
│   │       │   │   └── page.tsx  # Session detail
│   │       │   │       # BE: contracts-be/knowledge-transfer
│   │       │   │       # GET /v1/contracts/{contract_id}/knowledge-transfer/sessions/{session_id}
│   │       │   │
│   │       │   └── page.tsx  # Sessions list
│   │       │       # - Schedule sessions
│   │       │       # - Session recordings
│   │       │       # BE: contracts-be/knowledge-transfer
│   │       │       # GET /v1/contracts/{contract_id}/knowledge-transfer/sessions
│   │       │
│   │       ├── documentation/
│   │       │   └── page.tsx  # Transfer documentation
│   │       │       # - Create docs
│   │       │       # - Track completeness
│   │       │       # BE: contracts-be/knowledge-transfer, storage-be/asset
│   │       │       # GET /v1/contracts/{contract_id}/knowledge-transfer/documentation
│   │       │       # POST /v1/contracts/{contract_id}/knowledge-transfer/documentation
│   │       │
│   │       └── checklist/
│   │           └── page.tsx  # Transfer checklist
│   │               # - Required items
│   │               # - Completion status
│   │               # BE: contracts-be/knowledge-transfer
│   │               # GET /v1/contracts/{contract_id}/knowledge-transfer/checklist
│   │
│   └── benchmarking/
│       ├── performance/
│       │   └── page.tsx  # Performance benchmarking
│       │       # - Compare against similar contracts
│       │       # - Industry standards
│       │       # BE: contracts-be/analytics
│       │       # GET /v1/contracts/benchmarks/performance
│       │
│       ├── costs/
│       │   └── page.tsx  # Cost benchmarking
│       │       # - Rate comparison
│       │       # - Budget efficiency
│       │       # BE: contracts-be/analytics, financial-be/analytics
│       │       # GET /v1/contracts/benchmarks/costs
│       │
│       └── quality/
│           └── page.tsx  # Quality benchmarking
│               # - Deliverable quality comparison
│               # - Client satisfaction
│               # BE: contracts-be/analytics, reviews-be/analytics
│               # GET /v1/contracts/benchmarks/quality
```

### 3. Financial - Advanced Treasury & Risk Management

```
apps/web/src/app/[locale]/(dashboard)/
│
├── financial/
│   ├── treasury/
│   │   ├── dashboard/
│   │   │   └── page.tsx  # Treasury dashboard
│   │   │       # - Cash position
│   │   │       # - Liquidity forecast
│   │   │       # - Working capital
│   │   │       # BE: financial-be/treasury
│   │   │       # GET /v1/treasury/dashboard
│   │   │
│   │   ├── cash-flow/
│   │   │   ├── forecast/
│   │   │   │   └── page.tsx  # Cash flow forecasting
│   │   │   │       # - 30/60/90 day forecast
│   │   │   │       # - Scenario planning
│   │   │   │       # BE: financial-be/treasury
│   │   │   │       # GET /v1/treasury/cash-flow/forecast
│   │   │   │
│   │   │   ├── analysis/
│   │   │   │   └── page.tsx  # Cash flow analysis
│   │   │   │       # - Historical trends
│   │   │   │       # - Variance analysis
│   │   │   │       # BE: financial-be/treasury
│   │   │   │       # GET /v1/treasury/cash-flow/analysis
│   │   │   │
│   │   │   └── reports/
│   │   │       └── page.tsx  # Cash flow reports
│   │   │           # BE: financial-be/treasury
│   │   │           # GET /v1/treasury/cash-flow/reports
│   │   │
│   │   ├── liquidity/
│   │   │   ├── management/
│   │   │   │   └── page.tsx  # Liquidity management
│   │   │   │       # - Current ratio
│   │   │   │       # - Quick ratio
│   │   │   │       # - Reserve requirements
│   │   │   │       # BE: financial-be/treasury
│   │   │   │       # GET /v1/treasury/liquidity
│   │   │   │
│   │   │   └── alerts/
│   │   │       └── page.tsx  # Liquidity alerts
│   │   │           # - Threshold monitoring
│   │   │           # - Alert configuration
│   │   │           # BE: financial-be/treasury
│   │   │           # GET /v1/treasury/liquidity/alerts
│   │   │           # PUT /v1/treasury/liquidity/alerts
│   │   │
│   │   └── investments/
│   │       └── page.tsx  # Short-term investments
│   │           # - Investment tracking
│   │           # - Returns analysis
│   │           # BE: financial-be/treasury
│   │           # GET /v1/treasury/investments
│   │
│   ├── risk-management/
│   │   ├── exposure/
│   │   │   ├── currency/
│   │   │   │   └── page.tsx  # Currency exposure
│   │   │   │       # - FX risk analysis
│   │   │   │       # - Hedging strategies
│   │   │   │       # BE: financial-be/risk
│   │   │   │       # GET /v1/risk/exposure/currency
│   │   │   │
│   │   │   ├── credit/
│   │   │   │   └── page.tsx  # Credit exposure
│   │   │   │       # - Counterparty risk
│   │   │   │       # - Concentration risk
│   │   │   │       # BE: financial-be/risk
│   │   │   │       # GET /v1/risk/exposure/credit
│   │   │   │
│   │   │   └── operational/
│   │   │       └── page.tsx  # Operational risk
│   │   │           # - Process risks
│   │   │           # - System risks
│   │   │           # BE: financial-be/risk
│   │   │           # GET /v1/risk/exposure/operational
│   │   │
│   │   ├── limits/
│   │   │   └── page.tsx  # Risk limits
│   │   │       # - Set risk limits
│   │   │       # - Monitor breaches
│   │   │       # BE: financial-be/risk
│   │   │       # GET /v1/risk/limits
│   │   │       # PUT /v1/risk/limits
│   │   │
│   │   └── reports/
│   │       ├── var/
│   │       │   └── page.tsx  # Value at Risk
│   │       │       # BE: financial-be/risk
│   │       │       # GET /v1/risk/reports/var
│   │       │
│   │       └── stress-testing/
│   │           └── page.tsx  # Stress testing
│   │               # - Scenario analysis
│   │               # - Impact assessment
│   │               # BE: financial-be/risk
│   │               # GET /v1/risk/reports/stress-testing
│   │
│   ├── reconciliation/
│   │   ├── bank/
│   │   │   ├── [accountId]/
│   │   │   │   ├── auto-match/
│   │   │   │   │   └── page.tsx  # Auto-matching
│   │   │   │   │       # - Rule-based matching
│   │   │   │   │       # - AI-powered suggestions
│   │   │   │   │       # BE: financial-be/reconciliation
│   │   │   │   │       # POST /v1/reconciliation/bank/{account_id}/auto-match
│   │   │   │   │
│   │   │   │   ├── discrepancies/
│   │   │   │   │   └── page.tsx  # Discrepancy resolution
│   │   │   │   │       # BE: financial-be/reconciliation
│   │   │   │   │       # GET /v1/reconciliation/bank/{account_id}/discrepancies
│   │   │   │   │
│   │   │   │   └── page.tsx  # Account reconciliation
│   │   │   │       # BE: financial-be/reconciliation
│   │   │   │       # GET /v1/reconciliation/bank/{account_id}
│   │   │   │
│   │   │   └── page.tsx  # Bank reconciliation list
│   │   │       # BE: financial-be/reconciliation
│   │   │       # GET /v1/reconciliation/bank
│   │   │
│   │   ├── merchant/
│   │   │   └── page.tsx  # Merchant account reconciliation
│   │   │       # - Payment processor reconciliation
│   │   │       # - Fee reconciliation
│   │   │       # BE: financial-be/reconciliation
│   │   │       # GET /v1/reconciliation/merchant
│   │   │
│   │   ├── intercompany/
│   │   │   └── page.tsx  # Intercompany reconciliation
│   │   │       # BE: financial-be/reconciliation
│   │   │       # GET /v1/reconciliation/intercompany
│   │   │
│   │   └── reports/
│   │       └── page.tsx  # Reconciliation reports
│   │           # - Status reports
│   │           # - Exception reports
│   │           # BE: financial-be/reconciliation
│   │           # GET /v1/reconciliation/reports
│   │
│   ├── credit-management/
│   │   ├── limits/
│   │   │   ├── [clientId]/
│   │   │   │   └── page.tsx  # Client credit limit
│   │   │   │       # - Set/adjust limit
│   │   │   │       # - Usage tracking
│   │   │   │       # BE: financial-be/credit
│   │   │   │       # GET /v1/credit/limits/{client_id}
│   │   │   │       # PUT /v1/credit/limits/{client_id}
│   │   │   │
│   │   │   └── page.tsx  # All credit limits
│   │   │       # BE: financial-be/credit
│   │   │       # GET /v1/credit/limits
│   │   │
│   │   ├── scoring/
│   │   │   └── page.tsx  # Credit scoring
│   │   │       # - Credit assessment
│   │   │       # - Risk rating
│   │   │       # BE: financial-be/credit
│   │   │       # GET /v1/credit/scoring
│   │   │       # POST /v1/credit/scoring/assess
│   │   │
│   │   ├── collections/
│   │   │   ├── aging/
│   │   │   │   └── page.tsx  # Aging analysis
│   │   │   │       # - 30/60/90 day buckets
│   │   │   │       # - Collection priority
│   │   │   │       # BE: financial-be/credit
│   │   │   │       # GET /v1/credit/collections/aging
│   │   │   │
│   │   │   ├── actions/
│   │   │   │   └── page.tsx  # Collection actions
│   │   │   │       # - Reminders
│   │   │   │       # - Escalations
│   │   │   │       # BE: financial-be/credit, communications-be
│   │   │   │       # POST /v1/credit/collections/actions
│   │   │   │
│   │   │   └── page.tsx  # Collections dashboard
│   │   │       # BE: financial-be/credit
│   │   │       # GET /v1/credit/collections
│   │   │
│   │   └── disputes/
│   │       └── page.tsx  # Credit disputes
│   │           # BE: financial-be/credit
│   │           # GET /v1/credit/disputes
│   │
│   ├── budgets/
│   │   ├── [budgetId]/
│   │   │   ├── allocations/
│   │   │   │   └── page.tsx  # Budget allocations
│   │   │   │       # - Department allocations
│   │   │   │       # - Project allocations
│   │   │   │       # BE: financial-be/budget
│   │   │   │       # GET /v1/budgets/{budget_id}/allocations
│   │   │   │       # POST /v1/budgets/{budget_id}/allocations
│   │   │   │
│   │   │   ├── variance/
│   │   │   │   └── page.tsx  # Budget variance
│   │   │   │       # - Actual vs. planned
│   │   │   │       # - Variance analysis
│   │   │   │       # BE: financial-be/budget
│   │   │   │       # GET /v1/budgets/{budget_id}/variance
│   │   │   │
│   │   │   ├── adjustments/
│   │   │   │   └── page.tsx  # Budget adjustments
│   │   │   │       # - Reallocation requests
│   │   │   │       # - Approval workflow
│   │   │   │       # BE: financial-be/budget
│   │   │   │       # GET /v1/budgets/{budget_id}/adjustments
│   │   │   │       # POST /v1/budgets/{budget_id}/adjustments
│   │   │   │
│   │   │   └── forecasts/
│   │   │       └── page.tsx  # Budget forecasts
│   │   │           # - Rolling forecasts
│   │   │           # - Projection scenarios
│   │   │           # BE: financial-be/budget
│   │   │           # GET /v1/budgets/{budget_id}/forecasts
│   │   │
│   │   ├── templates/
│   │   │   └── page.tsx  # Budget templates
│   │   │       # - Standard templates
│   │   │       # - Create from template
│   │   │       # BE: financial-be/budget
│   │   │       # GET /v1/budgets/templates
│   │   │
│   │   └── consolidation/
│   │       └── page.tsx  # Budget consolidation
│   │           # - Roll-up view
│   │           # - Cross-department view
│   │           # BE: financial-be/budget
│   │           # GET /v1/budgets/consolidation
│   │
│   ├── forecasting/
│   │   ├── revenue/
│   │   │   ├── models/
│   │   │   │   └── page.tsx  # Revenue models
│   │   │   │       # - Historical models
│   │   │   │       # - Predictive models
│   │   │   │       # BE: financial-be/forecasting
│   │   │   │       # GET /v1/forecasting/revenue/models
│   │   │   │
│   │   │   ├── scenarios/
│   │   │   │   └── page.tsx  # Revenue scenarios
│   │   │   │       # - Best/worst/likely cases
│   │   │   │       # - Sensitivity analysis
│   │   │   │       # BE: financial-be/forecasting
│   │   │   │       # GET /v1/forecasting/revenue/scenarios
│   │   │   │       # POST /v1/forecasting/revenue/scenarios
│   │   │   │
│   │   │   └── page.tsx  # Revenue forecast
│   │   │       # BE: financial-be/forecasting
│   │   │       # GET /v1/forecasting/revenue
│   │   │
│   │   ├── expenses/
│   │   │   └── page.tsx  # Expense forecast
│   │   │       # BE: financial-be/forecasting
│   │   │       # GET /v1/forecasting/expenses
│   │   │
│   │   ├── profitability/
│   │   │   └── page.tsx  # Profitability forecast
│   │   │       # - Margin analysis
│   │   │       # - Break-even analysis
│   │   │       # BE: financial-be/forecasting
│   │   │       # GET /v1/forecasting/profitability
│   │   │
│   │   └── validation/
│   │       └── page.tsx  # Forecast validation
│   │           # - Accuracy tracking
│   │           # - Model refinement
│   │           # BE: financial-be/forecasting
│   │           # GET /v1/forecasting/validation
│   │
│   ├── compliance/
│   │   ├── aml/
│   │   │   ├── monitoring/
│   │   │   │   └── page.tsx  # AML monitoring
│   │   │   │       # - Transaction monitoring
│   │   │   │       # - Suspicious activity
│   │   │   │       # BE: financial-be/compliance
│   │   │   │       # GET /v1/compliance/aml/monitoring
│   │   │   │
│   │   │   ├── reports/
│   │   │   │   └── page.tsx  # AML reports
│   │   │   │       # - SAR filing
│   │   │   │       # - Compliance reports
│   │   │   │       # BE: financial-be/compliance
│   │   │   │       # GET /v1/compliance/aml/reports
│   │   │   │       # POST /v1/compliance/aml/reports/sar
│   │   │   │
│   │   │   └── page.tsx  # AML dashboard
│   │   │       # BE: financial-be/compliance
│   │   │       # GET /v1/compliance/aml
│   │   │
│   │   ├── kyc/
│   │   │   ├── verification/
│   │   │   │   └── page.tsx  # KYC verification
│   │   │   │       # BE: financial-be/compliance, admin-be/kyc
│   │   │   │       # GET /v1/compliance/kyc/verification
│   │   │   │
│   │   │   └── due-diligence/
│   │   │       └── page.tsx  # Enhanced due diligence
│   │   │           # BE: financial-be/compliance
│   │   │           # GET /v1/compliance/kyc/due-diligence
│   │   │
│   │   ├── sanctions/
│   │   │   ├── screening/
│   │   │   │   └── page.tsx  # Sanctions screening
│   │   │   │       # - Watchlist screening
│   │   │   │       # - PEP screening
│   │   │   │       # BE: financial-be/compliance
│   │   │   │       # POST /v1/compliance/sanctions/screen
│   │   │   │
│   │   │   └── alerts/
│   │   │       └── page.tsx  # Sanctions alerts
│   │   │           # BE: financial-be/compliance
│   │   │           # GET /v1/compliance/sanctions/alerts
│   │   │
│   │   └── audits/
│   │       ├── schedule/
│   │       │   └── page.tsx  # Audit schedule
│   │       │       # BE: financial-be/compliance
│   │       │       # GET /v1/compliance/audits/schedule
│   │       │
│   │       ├── findings/
│   │       │   └── page.tsx  # Audit findings
│   │       │       # BE: financial-be/compliance
│   │       │       # GET /v1/compliance/audits/findings
│   │       │
│   │       └── remediation/
│   │           └── page.tsx  # Remediation tracking
│   │               # BE: financial-be/compliance
│   │               # GET /v1/compliance/audits/remediation
│   │
│   └── analytics/
│       ├── profitability/
│       │   ├── by-client/
│       │   │   └── page.tsx  # Client profitability
│       │   │       # BE: financial-be/analytics
│       │   │       # GET /v1/analytics/profitability/by-client
│       │   │
│       │   ├── by-project/
│       │   │   └── page.tsx  # Project profitability
│       │   │       # BE: financial-be/analytics
│       │   │       # GET /v1/analytics/profitability/by-project
│       │   │
│       │   └── by-service/
│       │       └── page.tsx  # Service line profitability
│       │           # BE: financial-be/analytics
│       │           # GET /v1/analytics/profitability/by-service
│       │
│       ├── margins/
│       │   └── page.tsx  # Margin analysis
│       │       # - Gross margin
│       │       # - Operating margin
│       │       # - Net margin
│       │       # BE: financial-be/analytics
│       │       # GET /v1/analytics/margins
│       │
│       └── kpis/
│           ├── dashboard/
│           │   └── page.tsx  # Financial KPI dashboard
│           │       # - Custom KPIs
│           │       # - Benchmarking
│           │       # BE: financial-be/analytics
│           │       # GET /v1/analytics/kpis/dashboard
│           │
│           └── custom/
│               └── page.tsx  # Custom KPI builder
│                   # BE: financial-be/analytics
│                   # POST /v1/analytics/kpis/custom
```

### 4. Teams & Organizations - Enterprise Features

```
apps/web/src/app/[locale]/(dashboard)/
│
├── teams/
│   ├── [teamId]/
│   │   ├── hierarchy/
│   │   │   ├── org-chart/
│   │   │   │   └── page.tsx  # Organization chart
│   │   │   │       # - Visual hierarchy
│   │   │   │       # - Drag-and-drop reorg
│   │   │   │       # BE: users-be/team
│   │   │   │       # GET /v1/teams/{team_id}/hierarchy
│   │   │   │
│   │   │   ├── reporting-lines/
│   │   │   │   └── page.tsx  # Reporting structure
│   │   │   │       # BE: users-be/team
│   │   │   │       # GET /v1/teams/{team_id}/reporting-lines
│   │   │   │
│   │   │   └── page.tsx  # Team hierarchy overview
│   │   │       # BE: users-be/team
│   │   │       # GET /v1/teams/{team_id}/hierarchy
│   │   │
│   │   ├── capacity/
│   │   │   ├── planning/
│   │   │   │   └── page.tsx  # Capacity planning
│   │   │   │       # - Resource allocation
│   │   │   │       # - Utilization forecasting
│   │   │   │       # BE: users-be/team, contracts-be
│   │   │   │       # GET /v1/teams/{team_id}/capacity/planning
│   │   │   │
│   │   │   ├── utilization/
│   │   │   │   └── page.tsx  # Utilization tracking
│   │   │   │       # - Current utilization
│   │   │   │       # - Historical trends
│   │   │   │       # BE: users-be/team
│   │   │   │       # GET /v1/teams/{team_id}/capacity/utilization
│   │   │   │
│   │   │   └── forecasting/
│   │   │       └── page.tsx  # Capacity forecasting
│   │   │           # BE: users-be/team
│   │   │           # GET /v1/teams/{team_id}/capacity/forecast
│   │   │
│   │   ├── workflows/
│   │   │   ├── [workflowId]/
│   │   │   │   ├── edit/
│   │   │   │   │   └── page.tsx  # Edit workflow
│   │   │   │   │       # BE: users-be/workflow
│   │   │   │   │       # PUT /v1/teams/{team_id}/workflows/{workflow_id}
│   │   │   │   │
│   │   │   │   ├── analytics/
│   │   │   │   │   └── page.tsx  # Workflow analytics
│   │   │   │   │       # - Completion rates
│   │   │   │   │       # - Bottlenecks
│   │   │   │   │       # BE: users-be/workflow
│   │   │   │   │       # GET /v1/teams/{team_id}/workflows/{workflow_id}/analytics
│   │   │   │   │
│   │   │   │   └── page.tsx  # Workflow detail
│   │   │   │       # BE: users-be/workflow
│   │   │   │       # GET /v1/teams/{team_id}/workflows/{workflow_id}
│   │   │   │
│   │   │   └── page.tsx  # Workflows list
│   │   │       # BE: users-be/workflow
│   │   │       # GET /v1/teams/{team_id}/workflows
│   │   │
│   │   ├── policies/
│   │   │   ├── [policyId]/
│   │   │   │   ├── versions/
│   │   │   │   │   └── page.tsx  # Policy versions
│   │   │   │   │       # BE: users-be/policy
│   │   │   │   │       # GET /v1/teams/{team_id}/policies/{policy_id}/versions
│   │   │   │   │
│   │   │   │   ├── attestations/
│   │   │   │   │   └── page.tsx  # Policy attestations
│   │   │   │   │       # - Track acknowledgments
│   │   │   │   │       # - Compliance tracking
│   │   │   │   │       # BE: users-be/policy
│   │   │   │   │       # GET /v1/teams/{team_id}/policies/{policy_id}/attestations
│   │   │   │   │
│   │   │   │   └── page.tsx  # Policy detail
│   │   │   │       # BE: users-be/policy
│   │   │   │       # GET /v1/teams/{team_id}/policies/{policy_id}
│   │   │   │
│   │   │   └── page.tsx  # Policies list
│   │   │       # BE: users-be/policy
│   │   │       # GET /v1/teams/{team_id}/policies
│   │   │
│   │   ├── training/
│   │   │   ├── programs/
│   │   │   │   ├── [programId]/
│   │   │   │   │   ├── enroll/
│   │   │   │   │   │   └── page.tsx  # Enroll in program
│   │   │   │   │   │       # BE: users-be/training
│   │   │   │   │   │       # POST /v1/teams/{team_id}/training/programs/{program_id}/enroll
│   │   │   │   │   │
│   │   │   │   │   ├── progress/
│   │   │   │   │   │   └── page.tsx  # Training progress
│   │   │   │   │   │       # BE: users-be/training
│   │   │   │   │   │       # GET /v1/teams/{team_id}/training/programs/{program_id}/progress
│   │   │   │   │   │
│   │   │   │   │   └── page.tsx  # Program detail
│   │   │   │   │       # BE: users-be/training
│   │   │   │   │       # GET /v1/teams/{team_id}/training/programs/{program_id}
│   │   │   │   │
│   │   │   │   └── page.tsx  # Training programs
│   │   │   │       # BE: users-be/training
│   │   │   │       # GET /v1/teams/{team_id}/training/programs
│   │   │   │
│   │   │   ├── certifications/
│   │   │   │   └── page.tsx  # Team certifications
│   │   │   │       # - Track certifications
│   │   │   │       # - Expiration alerts
│   │   │   │       # BE: users-be/training
│   │   │   │       # GET /v1/teams/{team_id}/training/certifications
│   │   │   │
│   │   │   └── compliance/
│   │   │       └── page.tsx  # Training compliance
│   │   │           # - Completion tracking
│   │   │           # - Compliance reports
│   │   │           # BE: users-be/training
│   │   │           # GET /v1/teams/{team_id}/training/compliance
│   │   │
│   │   ├── procurement/
│   │   │   ├── requests/
│   │   │   │   ├── [requestId]/
│   │   │   │   │   ├── approval/
│   │   │   │   │   │   └── page.tsx  # Approve procurement
│   │   │   │   │   │       # BE: users-be/procurement
│   │   │   │   │   │       # POST /v1/teams/{team_id}/procurement/requests/{request_id}/approve
│   │   │   │   │   │
│   │   │   │   │   └── page.tsx  # Request detail
│   │   │   │   │       # BE: users-be/procurement
│   │   │   │   │       # GET /v1/teams/{team_id}/procurement/requests/{request_id}
│   │   │   │   │
│   │   │   │   └── page.tsx  # Procurement requests
│   │   │   │       # BE: users-be/procurement
│   │   │   │       # GET /v1/teams/{team_id}/procurement/requests
│   │   │   │       # POST /v1/teams/{team_id}/procurement/requests
│   │   │   │
│   │   │   ├── vendors/
│   │   │   │   ├── preferred/
│   │   │   │   │   └── page.tsx  # Preferred vendors
│   │   │   │   │       # BE: users-be/procurement
│   │   │   │   │       # GET /v1/teams/{team_id}/procurement/vendors/preferred
│   │   │   │   │
│   │   │   │   ├── evaluation/
│   │   │   │   │   └── page.tsx  # Vendor evaluation
│   │   │   │   │       # BE: users-be/procurement
│   │   │   │   │       # GET /v1/teams/{team_id}/procurement/vendors/evaluation
│   │   │   │   │
│   │   │   │   └── page.tsx  # Vendor management
│   │   │   │       # BE: users-be/procurement
│   │   │   │       # GET /v1/teams/{team_id}/procurement/vendors
│   │   │   │
│   │   │   └── contracts/
│   │   │       └── page.tsx  # Procurement contracts
│   │   │           # BE: users-be/procurement, contracts-be
│   │   │           # GET /v1/teams/{team_id}/procurement/contracts
│   │   │
│   │   └── knowledge-base/
│   │       ├── articles/
│   │       │   ├── [articleId]/
│   │       │   │   ├── edit/
│   │       │   │   │   └── page.tsx  # Edit article
│   │       │   │   │       # BE: users-be/knowledge-base
│   │       │   │   │       # PUT /v1/teams/{team_id}/kb/articles/{article_id}
│   │       │   │   │
│   │       │   │   ├── versions/
│   │       │   │   │   └── page.tsx  # Article versions
│   │       │   │   │       # BE: users-be/knowledge-base
│   │       │   │   │       # GET /v1/teams/{team_id}/kb/articles/{article_id}/versions
│   │       │   │   │
│   │       │   │   └── page.tsx  # Article detail
│   │       │   │       # BE: users-be/knowledge-base
│   │       │   │       # GET /v1/teams/{team_id}/kb/articles/{article_id}
│   │       │   │
│   │       │   └── page.tsx  # Articles list
│   │       │       # BE: users-be/knowledge-base
│   │       │       # GET /v1/teams/{team_id}/kb/articles
│   │       │
│   │       ├── categories/
│   │       │   └── page.tsx  # KB categories
│   │       │       # BE: users-be/knowledge-base
│   │       │       # GET /v1/teams/{team_id}/kb/categories
│   │       │
│   │       └── search/
│   │           └── page.tsx  # KB search
│   │               # BE: users-be/knowledge-base, search-be
│   │               # GET /v1/teams/{team_id}/kb/search
│   │
│   └── cross-team/
│       ├── collaboration/
│       │   └── page.tsx  # Cross-team collaboration
│       │       # - Shared projects
│       │       # - Resource sharing
│       │       # BE: users-be/team
│       │       # GET /v1/teams/collaboration
│       │
│       └── benchmarking/
│           └── page.tsx  # Inter-team benchmarking
│               # - Performance comparison
│               # - Best practices sharing
│               # BE: users-be/team
│               # GET /v1/teams/benchmarking
```

### 5. Settings - Enterprise & Developer Features

```
apps/web/src/app/[locale]/(dashboard)/
│
├── settings/
│   ├── developer/
│   │   ├── api-keys/
│   │   │   ├── [keyId]/
│   │   │   │   ├── rotate/
│   │   │   │   │   └── page.tsx  # Rotate API key
│   │   │   │   │       # BE: admin-be/api-keys
│   │   │   │   │       # POST /v1/developer/api-keys/{key_id}/rotate
│   │   │   │   │
│   │   │   │   ├── permissions/
│   │   │   │   │   └── page.tsx  # Key permissions
│   │   │   │   │       # BE: admin-be/api-keys
│   │   │   │   │       # GET /v1/developer/api-keys/{key_id}/permissions
│   │   │   │   │       # PUT /v1/developer/api-keys/{key_id}/permissions
│   │   │   │   │
│   │   │   │   └── logs/
│   │   │   │       └── page.tsx  # API key usage logs
│   │   │   │           # BE: admin-be/api-keys
│   │   │   │           # GET /v1/developer/api-keys/{key_id}/logs
│   │   │   │
│   │   │   └── page.tsx  # API keys management
│   │   │       # BE: admin-be/api-keys
│   │   │       # GET /v1/developer/api-keys
│   │   │       # POST /v1/developer/api-keys
│   │   │
│   │   ├── webhooks/
│   │   │   ├── [webhookId]/
│   │   │   │   ├── test/
│   │   │   │   │   └── page.tsx  # Test webhook
│   │   │   │   │       # BE: admin-be/webhooks
│   │   │   │   │       # POST /v1/developer/webhooks/{webhook_id}/test
│   │   │   │   │
│   │   │   │   ├── deliveries/
│   │   │   │   │   ├── [deliveryId]/
│   │   │   │   │   │   └── page.tsx  # Delivery detail
│   │   │   │   │   │       # BE: admin-be/webhooks
│   │   │   │   │   │       # GET /v1/developer/webhooks/{webhook_id}/deliveries/{delivery_id}
│   │   │   │   │   │
│   │   │   │   │   └── page.tsx  # Webhook deliveries
│   │   │   │   │       # - Delivery history
│   │   │   │   │       # - Retry failed deliveries
│   │   │   │   │       # BE: admin-be/webhooks
│   │   │   │   │       # GET /v1/developer/webhooks/{webhook_id}/deliveries
│   │   │   │   │
│   │   │   │   ├── events/
│   │   │   │   │   └── page.tsx  # Webhook event config
│   │   │   │   │       # BE: admin-be/webhooks
│   │   │   │   │       # GET /v1/developer/webhooks/{webhook_id}/events
│   │   │   │   │       # PUT /v1/developer/webhooks/{webhook_id}/events
│   │   │   │   │
│   │   │   │   └── page.tsx  # Webhook detail
│   │   │   │       # BE: admin-be/webhooks
│   │   │   │       # GET /v1/developer/webhooks/{webhook_id}
│   │   │   │
│   │   │   └── page.tsx  # Webhooks management
│   │   │       # BE: admin-be/webhooks
│   │   │       # GET /v1/developer/webhooks
│   │   │       # POST /v1/developer/webhooks
│   │   │
│   │   ├── oauth-apps/
│   │   │   ├── [appId]/
│   │   │   │   ├── credentials/
│   │   │   │   │   └── page.tsx  # OAuth credentials
│   │   │   │   │       # - Client ID
│   │   │   │   │       # - Client secret rotation
│   │   │   │   │       # BE: admin-be/oauth
│   │   │   │   │       # GET /v1/developer/oauth-apps/{app_id}/credentials
│   │   │   │   │       # POST /v1/developer/oauth-apps/{app_id}/rotate-secret
│   │   │   │   │
│   │   │   │   ├── scopes/
│   │   │   │   │   └── page.tsx  # OAuth scopes
│   │   │   │   │       # BE: admin-be/oauth
│   │   │   │   │       # GET /v1/developer/oauth-apps/{app_id}/scopes
│   │   │   │   │       # PUT /v1/developer/oauth-apps/{app_id}/scopes
│   │   │   │   │
│   │   │   │   ├── authorizations/
│   │   │   │   │   └── page.tsx  # App authorizations
│   │   │   │   │       # BE: admin-be/oauth
│   │   │   │   │       # GET /v1/developer/oauth-apps/{app_id}/authorizations
│   │   │   │   │
│   │   │   │   └── page.tsx  # OAuth app detail
│   │   │   │       # BE: admin-be/oauth
│   │   │   │       # GET /v1/developer/oauth-apps/{app_id}
│   │   │   │
│   │   │   └── page.tsx  # OAuth apps
│   │   │       # BE: admin-be/oauth
│   │   │       # GET /v1/developer/oauth-apps
│   │   │       # POST /v1/developer/oauth-apps
│   │   │
│   │   ├── sandbox/
│   │   │   ├── environments/
│   │   │   │   └── page.tsx  # Sandbox environments
│   │   │   │       # BE: admin-be/sandbox
│   │   │   │       # GET /v1/developer/sandbox/environments
│   │   │   │       # POST /v1/developer/sandbox/environments
│   │   │   │
│   │   │   └── test-data/
│   │   │       └── page.tsx  # Test data management
│   │   │           # BE: admin-be/sandbox
│   │   │           # GET /v1/developer/sandbox/test-data
│   │   │           # POST /v1/developer/sandbox/test-data
│   │   │
│   │   └── documentation/
│   │       └── page.tsx  # API documentation
│   │           # - Interactive docs
│   │           # - Code samples
│   │           # BE: None (static with API client examples)
│   │
│   ├── advanced/
│   │   ├── data-retention/
│   │   │   └── page.tsx  # Data retention policies
│   │   │       # - Retention rules
│   │   │       # - Deletion schedules
│   │   │       # BE: admin-be/data-governance
│   │   │       # GET /v1/settings/data-retention
│   │   │       # PUT /v1/settings/data-retention
│   │   │
│   │   ├── audit-logs/
│   │   │   ├── export/
│   │   │   │   └── page.tsx  # Export audit logs
│   │   │   │       # BE: admin-be/audit
│   │   │   │       # POST /v1/settings/audit-logs/export
│   │   │   │
│   │   │   └── page.tsx  # Audit log settings
│   │   │       # - Retention period
│   │   │       # - Log level
│   │   │       # BE: admin-be/audit
│   │   │       # GET /v1/settings/audit-logs
│   │   │       # PUT /v1/settings/audit-logs
│   │   │
│   │   ├── ip-whitelist/
│   │   │   └── page.tsx  # IP whitelist
│   │   │       # - Add/remove IPs
│   │   │       # - CIDR ranges
│   │   │       # BE: admin-be/security
│   │   │       # GET /v1/settings/ip-whitelist
│   │   │       # PUT /v1/settings/ip-whitelist
│   │   │
│   │   ├── session-management/
│   │   │   └── page.tsx  # Session settings
│   │   │       # - Timeout duration
│   │   │       # - Concurrent sessions
│   │   │       # - Force logout
│   │   │       # BE: admin-be/security
│   │   │       # GET /v1/settings/sessions
│   │   │       # PUT /v1/settings/sessions
│   │   │
│   │   └── rate-limiting/
│   │       └── page.tsx  # Rate limit configuration
│   │           # - API rate limits
│   │           # - Custom rules
│   │           # BE: admin-be/security
│   │           # GET /v1/settings/rate-limits
│   │           # PUT /v1/settings/rate-limits
│   │
│   └── labs/
│       └── page.tsx  # Experimental features
│           # - Beta features toggle
│           # - A/B test participation
│           # BE: admin-be/feature-flags
│           # GET /v1/settings/labs
│           # PUT /v1/settings/labs
```

---

## II. Admin Panel - Complete Enterprise Admin

### 1. Financial Operations - Deep Admin Control

```
apps/web/src/app/[locale]/(admin)/
│
├── financial/
│   ├── settlement/
│   │   ├── batches/
│   │   │   ├── [batchId]/
│   │   │   │   ├── review/
│   │   │   │   │   └── page.tsx  # Review settlement batch
│   │   │   │   │       # BE: financial-be/settlement
│   │   │   │   │       # GET /v1/admin/settlement/batches/{batch_id}
│   │   │   │   │
│   │   │   │   ├── approve/
│   │   │   │   │   └── page.tsx  # Approve settlement
│   │   │   │   │       # BE: financial-be/settlement, admin-be/change-approval
│   │   │   │   │       # POST /v1/admin/settlement/batches/{batch_id}/approve
│   │   │   │   │
│   │   │   │   └── page.tsx  # Batch detail
│   │   │   │       # BE: financial-be/settlement
│   │   │   │       # GET /v1/admin/settlement/batches/{batch_id}
│   │   │   │
│   │   │   └── page.tsx  # Settlement batches
│   │   │       # BE: financial-be/settlement
│   │   │       # GET /v1/admin/settlement/batches
│   │   │
│   │   ├── holds/
│   │   │   ├── [holdId]/
│   │   │   │   ├── release/
│   │   │   │   │   └── page.tsx  # Release hold
│   │   │   │   │       # BE: financial-be/settlement
│   │   │   │   │       # POST /v1/admin/settlement/holds/{hold_id}/release
│   │   │   │   │
│   │   │   │   └── page.tsx  # Hold detail
│   │   │   │       # BE: financial-be/settlement
│   │   │   │       # GET /v1/admin/settlement/holds/{hold_id}
│   │   │   │
│   │   │   └── page.tsx  # Payment holds
│   │   │       # BE: financial-be/settlement
│   │   │       # GET /v1/admin/settlement/holds
│   │   │
│   │   └── rules/
│   │       └── page.tsx  # Settlement rules
│   │           # - Auto-hold rules
│   │           # - Risk thresholds
│   │           # BE: financial-be/settlement
│   │           # GET /v1/admin/settlement/rules
│   │           # PUT /v1/admin/settlement/rules
│   │
│   ├── reserves/
│   │   ├── calculation/
│   │   │   └── page.tsx  # Reserve calculation
│   │   │       # - Reserve requirements
│   │   │       # - Rolling reserves
│   │   │       # BE: financial-be/reserves
│   │   │       # GET /v1/admin/reserves/calculation
│   │   │
│   │   ├── adjustments/
│   │   │   └── page.tsx  # Reserve adjustments
│   │   │       # BE: financial-be/reserves
│   │   │       # GET /v1/admin/reserves/adjustments
│   │   │       # POST /v1/admin/reserves/adjustments
│   │   │
│   │   └── releases/
│   │       └── page.tsx  # Reserve releases
│   │           # BE: financial-be/reserves
│   │           # GET /v1/admin/reserves/releases
│   │
│   ├── fraud/
│   │   ├── detection/
│   │   │   ├── rules/
│   │   │   │   └── page.tsx  # Fraud detection rules
│   │   │   │       # BE: financial-be/fraud
│   │   │   │       # GET /v1/admin/fraud/rules
│   │   │   │       # POST /v1/admin/fraud/rules
│   │   │   │
│   │   │   ├── ml-models/
│   │   │   │   └── page.tsx  # ML fraud models
│   │   │   │       # - Model performance
│   │   │   │       # - Model tuning
│   │   │   │       # BE: financial-be/fraud
│   │   │   │       # GET /v1/admin/fraud/ml-models
│   │   │   │
│   │   │   └── alerts/
│   │   │       ├── [alertId]/
│   │   │       │   └── page.tsx  # Alert investigation
│   │   │       │       # BE: financial-be/fraud
│   │   │       │       # GET /v1/admin/fraud/alerts/{alert_id}
│   │   │       │       # POST /v1/admin/fraud/alerts/{alert_id}/investigate
│   │   │       │
│   │   │       └── page.tsx  # Fraud alerts
│   │   │           # BE: financial-be/fraud
│   │   │           # GET /v1/admin/fraud/alerts
│   │   │
│   │   ├── cases/
│   │   │   ├── [caseId]/
│   │   │   │   ├── investigation/
│   │   │   │   │   └── page.tsx  # Case investigation
│   │   │   │   │       # BE: financial-be/fraud
│   │   │   │   │       # GET /v1/admin/fraud/cases/{case_id}/investigation
│   │   │   │   │
│   │   │   │   ├── evidence/
│   │   │   │   │   └── page.tsx  # Case evidence
│   │   │   │   │       # BE: financial-be/fraud, storage-be
│   │   │   │   │       # GET /v1/admin/fraud/cases/{case_id}/evidence
│   │   │   │   │
│   │   │   │   └── resolution/
│   │   │   │       └── page.tsx  # Case resolution
│   │   │   │           # BE: financial-be/fraud
│   │   │   │           # POST /v1/admin/fraud/cases/{case_id}/resolve
│   │   │   │
│   │   │   └── page.tsx  # Fraud cases
│   │   │       # BE: financial-be/fraud
│   │   │       # GET /v1/admin/fraud/cases
│   │   │
│   │   └── patterns/
│   │       └── page.tsx  # Fraud pattern analysis
│   │           # BE: financial-be/fraud
│   │           # GET /v1/admin/fraud/patterns
│   │
│   ├── currency/
│   │   ├── rates/
│   │   │   ├── manual-override/
│   │   │   │   └── page.tsx  # Manual rate override
│   │   │   │       # BE: financial-be/currency
│   │   │   │       # POST /v1/admin/currency/rates/override
│   │   │   │
│   │   │   └── page.tsx  # Exchange rates
│   │   │       # BE: financial-be/currency
│   │   │       # GET /v1/admin/currency/rates
│   │   │
│   │   ├── conversions/
│   │   │   └── page.tsx  # Currency conversions
│   │   │       # - Conversion history
│   │   │       # - Spread analysis
│   │   │       # BE: financial-be/currency
│   │   │       # GET /v1/admin/currency/conversions
│   │   │
│   │   └── hedging/
│   │       └── page.tsx  # Currency hedging
│   │           # BE: financial-be/currency
│   │           # GET /v1/admin/currency/hedging
│   │
│   └── fees/
│       ├── structures/
│       │   ├── [structureId]/
│       │   │   └── page.tsx  # Fee structure detail
│       │   │       # BE: financial-be/fees
│       │   │       # GET /v1/admin/fees/structures/{structure_id}
│       │   │
│       │   └── page.tsx  # Fee structures
│       │       # BE: financial-be/fees
│       │       # GET /v1/admin/fees/structures
│       │       # POST /v1/admin/fees/structures
│       │
│       ├── promotions/
│       │   ├── [promotionId]/
│       │   │   └── page.tsx  # Promotion detail
│       │   │       # BE: financial-be/fees
│       │   │       # GET /v1/admin/fees/promotions/{promotion_id}
│       │   │
│       │   └── page.tsx  # Fee promotions
│       │       # BE: financial-be/fees
│       │       # GET /v1/admin/fees/promotions
│       │       # POST /v1/admin/fees/promotions
│       │
│       └── overrides/
│           └── page.tsx  # Fee overrides
│               # - Custom fee arrangements
│               # - Volume discounts
│               # BE: financial-be/fees
│               # GET /v1/admin/fees/overrides
│               # POST /v1/admin/fees/overrides
```

### 2. Trust & Safety - Advanced Moderation

```
apps/web/src/app/[locale]/(admin)/
│
├── trust-safety/
│   ├── content-moderation/
│   │   ├── queue/
│   │   │   ├── priority/
│   │   │   │   └── page.tsx  # Priority queue
│   │   │   │       # BE: admin-be/moderation
│   │   │   │       # GET /v1/admin/moderation/queue?priority=high
│   │   │   │
│   │   │   ├── categories/
│   │   │   │   └── [category]/
│   │   │   │       └── page.tsx  # Category-specific queue
│   │   │   │           # BE: admin-be/moderation
│   │   │   │           # GET /v1/admin/moderation/queue?category={category}
│   │   │   │
│   │   │   └── page.tsx  # Moderation queue
│   │   │       # BE: admin-be/moderation
│   │   │       # GET /v1/admin/moderation/queue
│   │   │
│   │   ├── ml-assistance/
│   │   │   ├── predictions/
│   │   │   │   └── page.tsx  # ML predictions
│   │   │   │       # - Auto-classification
│   │   │   │       # - Confidence scores
│   │   │   │       # BE: admin-be/moderation
│   │   │   │       # GET /v1/admin/moderation/ml-predictions
│   │   │   │
│   │   │   ├── training/
│   │   │   │   └── page.tsx  # Model training
│   │   │   │       # - Feedback loop
│   │   │   │       # - Model retraining
│   │   │   │       # BE: admin-be/moderation
│   │   │   │       # POST /v1/admin/moderation/ml-training
│   │   │   │
│   │   │   └── accuracy/
│   │   │       └── page.tsx  # Model accuracy
│   │   │           # BE: admin-be/moderation
│   │   │           # GET /v1/admin/moderation/ml-accuracy
│   │   │
│   │   ├── automation/
│   │   │   ├── rules/
│   │   │   │   └── page.tsx  # Auto-moderation rules
│   │   │   │       # BE: admin-be/moderation
│   │   │   │       # GET /v1/admin/moderation/automation/rules
│   │   │   │       # POST /v1/admin/moderation/automation/rules
│   │   │   │
│   │   │   └── actions/
│   │   │       └── page.tsx  # Automated actions
│   │   │           # BE: admin-be/moderation
│   │   │           # GET /v1/admin/moderation/automation/actions
│   │   │
│   │   └── appeals/
│   │       ├── [appealId]/
│   │       │   └── page.tsx  # Appeal review
│   │       │       # BE: admin-be/moderation
│   │       │       # GET /v1/admin/moderation/appeals/{appeal_id}
│   │       │       # POST /v1/admin/moderation/appeals/{appeal_id}/decide
│   │       │
│   │       └── page.tsx  # Appeals queue
│   │           # BE: admin-be/moderation
│   │           # GET /v1/admin/moderation/appeals
│   │
│   ├── risk-scoring/
│   │   ├── models/
│   │   │   └── page.tsx  # Risk scoring models
│   │   │       # - User risk scores
│   │   │       # - Transaction risk scores
│   │   │       # BE: admin-be/risk-scoring
│   │   │       # GET /v1/admin/risk-scoring/models
│   │   │
│   │   ├── thresholds/
│   │   │   └── page.tsx  # Risk thresholds
│   │   │       # BE: admin-be/risk-scoring
│   │   │       # GET /v1/admin/risk-scoring/thresholds
│   │   │       # PUT /v1/admin/risk-scoring/thresholds
│   │   │
│   │   └── monitoring/
│   │       └── page.tsx  # Risk monitoring
│   │           # BE: admin-be/risk-scoring
│   │           # GET /v1/admin/risk-scoring/monitoring
│   │
│   ├── identity-verification/
│   │   ├── providers/
│   │   │   └── page.tsx  # Verification providers
│   │   │       # - Provider configuration
│   │   │       # - Fallback rules
│   │   │       # BE: admin-be/identity-verification
│   │   │       # GET /v1/admin/identity-verification/providers
│   │   │
│   │   ├── manual-review/
│   │   │   └── page.tsx  # Manual ID review
│   │   │       # BE: admin-be/identity-verification
│   │   │       # GET /v1/admin/identity-verification/manual-review
│   │   │
│   │   └── statistics/
│   │       └── page.tsx  # Verification statistics
│   │           # - Pass rates
│   │           # - Rejection reasons
│   │           # BE: admin-be/identity-verification
│   │           # GET /v1/admin/identity-verification/statistics
│   │
│   └── watchlists/
│       ├── global/
│       │   └── page.tsx  # Global watchlists
│       │       # - OFAC
│       │       # - UN sanctions
│       │       # - PEP lists
│       │       # BE: admin-be/watchlists
│       │       # GET /v1/admin/watchlists/global
│       │
│       ├── custom/
│       │   └── page.tsx  # Custom watchlists
│       │       # - Internal blacklist
│       │       # - High-risk users
│       │       # BE: admin-be/watchlists
│       │       # GET /v1/admin/watchlists/custom
│       │       # POST /v1/admin/watchlists/custom
│       │
│       └── monitoring/
│           └── page.tsx  # Watchlist monitoring
│               # - Screening results
│               # - False positives
│               # BE: admin-be/watchlists
│               # GET /v1/admin/watchlists/monitoring
```

### 3. Platform Operations - Deep System Control

```
apps/web/src/app/[locale]/(admin)/
│
├── operations/
│   ├── search-quality/
│   │   ├── relevance/
│   │   │   ├── tuning/
│   │   │   │   └── page.tsx  # Relevance tuning
│   │   │   │       # - Weights adjustment
│   │   │   │       # - Field boosting
│   │   │   │       # BE: search-be/admin
│   │   │   │       # GET /v1/admin/search/relevance/tuning
│   │   │   │       # PUT /v1/admin/search/relevance/tuning
│   │   │   │
│   │   │   ├── testing/
│   │   │   │   └── page.tsx  # Relevance testing
│   │   │   │       # - A/B testing
│   │   │   │       # - Test queries
│   │   │   │       # BE: search-be/admin
│   │   │   │       # POST /v1/admin/search/relevance/test
│   │   │   │
│   │   │   └── metrics/
│   │   │       └── page.tsx  # Relevance metrics
│   │   │           # - NDCG scores
│   │   │           # - Click-through rates
│   │   │           # BE: search-be/admin
│   │   │           # GET /v1/admin/search/relevance/metrics
│   │   │
│   │   ├── synonyms/
│   │   │   ├── [synonymId]/
│   │   │   │   └── page.tsx  # Synonym detail
│   │   │   │       # BE: search-be/admin
│   │   │   │       # GET /v1/admin/search/synonyms/{synonym_id}
│   │   │   │       # PUT /v1/admin/search/synonyms/{synonym_id}
│   │   │   │
│   │   │   └── page.tsx  # Synonym management
│   │   │       # BE: search-be/admin
│   │   │       # GET /v1/admin/search/synonyms
│   │   │       # POST /v1/admin/search/synonyms
│   │   │
│   │   ├── boosts/
│   │   │   └── page.tsx  # Search boosts
│   │   │       # - Term boosts
│   │   │       # - Document boosts
│   │   │       # BE: search-be/admin
│   │   │       # GET /v1/admin/search/boosts
│   │   │       # POST /v1/admin/search/boosts
│   │   │
│   │   └── reindexing/
│   │       ├── schedule/
│   │       │   └── page.tsx  # Reindex scheduling
│   │       │       # BE: search-be/admin
│   │       │       # POST /v1/admin/search/reindex/schedule
│   │       │
│   │       └── status/
│   │           └── page.tsx  # Reindex status
│   │               # BE: search-be/admin
│   │               # GET /v1/admin/search/reindex/status
│   │
│   ├── messaging/
│   │   ├── templates/
│   │   │   ├── [templateId]/
│   │   │   │   ├── versions/
│   │   │   │   │   └── page.tsx  # Template versions
│   │   │   │   │       # BE: communications-be/templates
│   │   │   │   │       # GET /v1/admin/messaging/templates/{template_id}/versions
│   │   │   │   │
│   │   │   │   ├── preview/
│   │   │   │   │   └── page.tsx  # Template preview
│   │   │   │   │       # BE: communications-be/templates
│   │   │   │   │       # POST /v1/admin/messaging/templates/{template_id}/preview
│   │   │   │   │
│   │   │   │   └── page.tsx  # Template detail
│   │   │   │       # BE: communications-be/templates
│   │   │   │       # GET /v1/admin/messaging/templates/{template_id}
│   │   │   │
│   │   │   └── page.tsx  # Message templates
│   │   │       # BE: communications-be/templates
│   │   │       # GET /v1/admin/messaging/templates
│   │   │       # POST /v1/admin/messaging/templates
│   │   │
│   │   ├── campaigns/
│   │   │   ├── [campaignId]/
│   │   │   │   ├── analytics/
│   │   │   │   │   └── page.tsx  # Campaign analytics
│   │   │   │   │       # - Open rates
│   │   │   │   │       # - Click rates
│   │   │   │   │       # - Conversions
│   │   │   │   │       # BE: communications-be/campaigns
│   │   │   │   │       # GET /v1/admin/messaging/campaigns/{campaign_id}/analytics
│   │   │   │   │
│   │   │   │   ├── test/
│   │   │   │   │   └── page.tsx  # Test campaign
│   │   │   │   │       # BE: communications-be/campaigns
│   │   │   │   │       # POST /v1/admin/messaging/campaigns/{campaign_id}/test
│   │   │   │   │
│   │   │   │   └── page.tsx  # Campaign detail
│   │   │   │       # BE: communications-be/campaigns
│   │   │   │       # GET /v1/admin/messaging/campaigns/{campaign_id}
│   │   │   │
│   │   │   └── page.tsx  # Campaigns
│   │   │       # BE: communications-be/campaigns
│   │   │       # GET /v1/admin/messaging/campaigns
│   │   │       # POST /v1/admin/messaging/campaigns
│   │   │
│   │   ├── rate-limits/
│   │   │   └── page.tsx  # Message rate limits
│   │   │       # - Per user limits
│   │   │       # - Global limits
│   │   │       # BE: communications-be/rate-limits
│   │   │       # GET /v1/admin/messaging/rate-limits
│   │   │       # PUT /v1/admin/messaging/rate-limits
│   │   │
│   │   └── deliverability/
│   │       ├── reputation/
│   │       │   └── page.tsx  # Sender reputation
│   │       │       # BE: communications-be/deliverability
│   │       │       # GET /v1/admin/messaging/deliverability/reputation
│   │       │
│   │       └── bounces/
│   │           └── page.tsx  # Bounce management
│   │               # BE: communications-be/deliverability
│   │               # GET /v1/admin/messaging/deliverability/bounces
│   │
│   ├── storage/
│   │   ├── quotas/
│   │   │   └── page.tsx  # Storage quotas
│   │   │       # - Per-user quotas
│   │   │       # - Org quotas
│   │   │       # BE: storage-be/admin
│   │   │       # GET /v1/admin/storage/quotas
│   │   │       # PUT /v1/admin/storage/quotas
│   │   │
│   │   ├── lifecycle/
│   │   │   └── page.tsx  # Storage lifecycle
│   │   │       # - Archival rules
│   │   │       # - Deletion policies
│   │   │       # BE: storage-be/admin
│   │   │       # GET /v1/admin/storage/lifecycle
│   │   │       # POST /v1/admin/storage/lifecycle
│   │   │
│   │   ├── virus-scanning/
│   │   │   └── page.tsx  # Virus scanning config
│   │   │       # BE: storage-be/admin
│   │   │       # GET /v1/admin/storage/virus-scanning
│   │   │
│   │   └── cdn/
│   │       └── page.tsx  # CDN configuration
│   │           # - Purge cache
│   │           # - Edge locations
│   │           # BE: storage-be/admin
│   │           # GET /v1/admin/storage/cdn
│   │           # POST /v1/admin/storage/cdn/purge
│   │
│   └── feature-flags/
│       ├── flags/
│       │   ├── [flagId]/
│       │   │   ├── targeting/
│       │   │   │   └── page.tsx  # Flag targeting
│       │   │   │       # - User segments
│       │   │   │       # - Percentage rollout
│       │   │   │       # BE: admin-be/feature-flags
│       │   │   │       # GET /v1/admin/feature-flags/{flag_id}/targeting
│       │   │   │       # PUT /v1/admin/feature-flags/{flag_id}/targeting
│       │   │   │
│       │   │   ├── history/
│       │   │   │   └── page.tsx  # Flag change history
│       │   │   │       # BE: admin-be/feature-flags
│       │   │   │       # GET /v1/admin/feature-flags/{flag_id}/history
│       │   │   │
│       │   │   └── page.tsx  # Flag detail
│       │   │       # BE: admin-be/feature-flags
│       │   │       # GET /v1/admin/feature-flags/{flag_id}
│       │   │
│       │   └── page.tsx  # Feature flags
│       │       # BE: admin-be/feature-flags
│       │       # GET /v1/admin/feature-flags
│       │       # POST /v1/admin/feature-flags
│       │
│       └── experiments/
│           ├── [experimentId]/
│           │   ├── results/
│           │   │   └── page.tsx  # Experiment results
│           │   │       # BE: admin-be/experiments
│           │   │       # GET /v1/admin/experiments/{experiment_id}/results
│           │   │
│           │   └── page.tsx  # Experiment detail
│           │       # BE: admin-be/experiments
│           │       # GET /v1/admin/experiments/{experiment_id}
│           │
│           └── page.tsx  # A/B experiments
│               # BE: admin-be/experiments
│               # GET /v1/admin/experiments
│               # POST /v1/admin/experiments
```

---

## III. Shared Features - Enterprise Infrastructure

### 1. Real-time & Events Infrastructure

```
packages/shared/src/
│
├── realtime/
│   ├── websocket/
│   │   ├── connection-manager.ts  # WebSocket connection management
│   │   │   # - Connection pooling
│   │   │   # - Reconnection logic
│   │   │   # - Heartbeat monitoring
│   │   │   # BE: communications-be/websocket
│   │   │
│   │   ├── message-handler.ts  # Message routing
│   │   │   # - Message type handling
│   │   │   # - Event dispatch
│   │   │
│   │   ├── subscription-manager.ts  # Subscription management
│   │   │   # - Channel subscriptions
│   │   │   # - Topic filtering
│   │   │
│   │   └── error-recovery.ts  # Error handling & recovery
│   │
│   ├── presence/
│   │   ├── presence-tracker.ts  # User presence tracking
│   │   │   # BE: communications-be/presence
│   │   │   # POST /v1/presence/update
│   │   │
│   │   ├── hooks/
│   │   │   ├── usePresence.ts  # Presence hook
│   │   │   └── useOnlineUsers.ts  # Online users hook
│   │   │
│   │   └── types.ts  # Presence types
│   │
│   ├── typing-indicators/
│   │   ├── typing-manager.ts  # Typing indicator management
│   │   │   # BE: communications-be/typing
│   │   │   # POST /v1/conversations/{id}/typing
│   │   │
│   │   └── hooks/
│   │       └── useTypingIndicator.ts  # Typing hook
│   │
│   └── live-updates/
│       ├── event-stream.ts  # Server-sent events
│       ├── polling-fallback.ts  # Polling fallback
│       └── hooks/
│           ├── useLiveQuery.ts  # Live query hook
│           └── useLiveDocument.ts  # Live document hook
```

### 2. Advanced Analytics & Telemetry

```
packages/shared/src/
│
├── analytics/
│   ├── collectors/
│   │   ├── page-view-collector.ts  # Page view collection
│   │   ├── interaction-collector.ts  # Interaction tracking
│   │   ├── performance-collector.ts  # Performance metrics
│   │   └── error-collector.ts  # Error tracking
│   │
│   ├── enrichment/
│   │   ├── context-enricher.ts  # Event context enrichment
│   │   ├── user-enricher.ts  # User data enrichment
│   │   └── session-enricher.ts  # Session data enrichment
│   │
│   ├── batching/
│   │   ├── event-batcher.ts  # Event batching
│   │   ├── flush-strategies.ts  # Flush strategies
│   │   └── queue-manager.ts  # Queue management
│   │
│   ├── sampling/
│   │   ├── sampler.ts  # Event sampling
│   │   └── strategies/
│   │       ├── percentage-sampling.ts
│   │       ├── user-based-sampling.ts
│   │       └── adaptive-sampling.ts
│   │
│   └── privacy/
│       ├── pii-scrubber.ts  # PII removal
│       ├── data-minimization.ts  # Data minimization
│       └── consent-manager.ts  # Consent management
```

### 3. Performance Monitoring

```
packages/shared/src/
│
├── performance/
│   ├── web-vitals/
│   │   ├── collectors/
│   │   │   ├── lcp-collector.ts  # Largest Contentful Paint
│   │   │   ├── fid-collector.ts  # First Input Delay
│   │   │   ├── cls-collector.ts  # Cumulative Layout Shift
│   │   │   ├── ttfb-collector.ts  # Time to First Byte
│   │   │   └── inp-collector.ts  # Interaction to Next Paint
│   │   │
│   │   ├── attribution.ts  # Performance attribution
│   │   └── reporting.ts  # Web Vitals reporting
│   │
│   ├── resource-timing/
│   │   ├── api-timing.ts  # API request timing
│   │   ├── asset-timing.ts  # Asset load timing
│   │   └── navigation-timing.ts  # Navigation timing
│   │
│   ├── custom-metrics/
│   │   ├── time-to-interactive.ts  # Custom TTI
│   │   ├── route-change-duration.ts  # Route timing
│   │   └── component-render-time.ts  # Component perf
│   │
│   └── monitoring/
│       ├── performance-observer.ts  # Performance Observer API
│       ├── long-tasks.ts  # Long task detection
│       └── memory-usage.ts  # Memory monitoring
```

### 4. Offline & Synchronization (Mobile Focus)

```
packages/shared/src/
│
├── offline/
│   ├── queue/
│   │   ├── operation-queue.ts  # Offline operation queue
│   │   ├── queue-processor.ts  # Queue processing
│   │   ├── retry-strategies.ts  # Retry logic
│   │   └── conflict-resolution.ts  # Conflict handling
│   │
│   ├── storage/
│   │   ├── indexed-db.ts  # IndexedDB wrapper
│   │   ├── async-storage.ts  # React Native AsyncStorage
│   │   └── cache-manager.ts  # Cache management
│   │
│   ├── sync/
│   │   ├── sync-engine.ts  # Synchronization engine
│   │   ├── sync-strategies/
│   │   │   ├── full-sync.ts  # Full synchronization
│   │   │   ├── incremental-sync.ts  # Incremental sync
│   │   │   └── conflict-merge.ts  # Merge strategies
│   │   │
│   │   ├── change-detection.ts  # Change detection
│   │   └── version-tracking.ts  # Version management
│   │
│   ├── hooks/
│   │   ├── useOfflineQueue.ts  # Offline queue hook
│   │   ├── useSyncStatus.ts  # Sync status hook
│   │   ├── useOnlineStatus.ts  # Online status hook
│   │   └── useOfflineStorage.ts  # Offline storage hook
│   │
│   └── network/
│       ├── network-monitor.ts  # Network status monitoring
│       ├── bandwidth-estimator.ts  # Bandwidth estimation
│       └── connection-quality.ts  # Connection quality
```

### 5. Geolocation & Internationalization

```
packages/shared/src/
│
├── geolocation/
│   ├── ip-detection/
│   │   ├── ip-resolver.ts  # IP-based location
│   │   │   # BE: utility-be/geolocation
│   │   │   # GET /v1/geo/ip-lookup
│   │   │
│   │   └── cache.ts  # Location cache
│   │
│   ├── browser-api/
│   │   ├── geolocation-api.ts  # Browser Geolocation API
│   │   ├── permissions.ts  # Permission handling
│   │   └── fallback.ts  # Fallback strategies
│   │
│   ├── timezone/
│   │   ├── timezone-detector.ts  # Timezone detection
│   │   ├── timezone-converter.ts  # Timezone conversion
│   │   └── dst-handler.ts  # Daylight saving handling
│   │
│   ├── country-detection/
│   │   ├── detector.ts  # Country detection
│   │   ├── locale-mapping.ts  # Country to locale
│   │   └── currency-mapping.ts  # Country to currency
│   │
│   └── hooks/
│       ├── useGeolocation.ts  # Geolocation hook
│       ├── useTimezone.ts  # Timezone hook
│       └── useCountry.ts  # Country hook
```

### 6. Experimentation & Feature Management

```
packages/shared/src/
│
├── experiments/
│   ├── ab-testing/
│   │   ├── experiment-engine.ts  # A/B test engine
│   │   │   # BE: admin-be/experiments
│   │   │   # GET /v1/experiments/active
│   │   │   # POST /v1/experiments/{id}/track
│   │   │
│   │   ├── variant-selection.ts  # Variant assignment
│   │   ├── bucketing.ts  # User bucketing
│   │   └── tracking.ts  # Experiment tracking
│   │
│   ├── feature-variants/
│   │   ├── variant-manager.ts  # Feature variants
│   │   ├── rollout-strategies.ts  # Rollout strategies
│   │   └── targeting.ts  # User targeting
│   │
│   ├── hooks/
│   │   ├── useExperiment.ts  # Experiment hook
│   │   ├── useVariant.ts  # Variant hook
│   │   └── useFeatureVariant.ts  # Feature variant hook
│   │
│   └── analytics/
│       ├── conversion-tracking.ts  # Conversion tracking
│       ├── metric-collection.ts  # Metric collection
│       └── statistical-analysis.ts  # Statistical analysis
```

### 7. Gamification & Rewards

```
packages/shared/src/
│
├── gamification/
│   ├── achievements/
│   │   ├── achievement-engine.ts  # Achievement system
│   │   │   # BE: users-be/achievements
│   │   │   # GET /v1/achievements
│   │   │   # POST /v1/achievements/{id}/claim
│   │   │
│   │   ├── achievement-tracker.ts  # Progress tracking
│   │   ├── achievement-notifier.ts  # Achievement notifications
│   │   └── types.ts  # Achievement types
│   │
│   ├── badges/
│   │   ├── badge-system.ts  # Badge management
│   │   │   # BE: users-be/badges
│   │   │   # GET /v1/badges
│   │   │
│   │   ├── badge-criteria.ts  # Badge criteria
│   │   └── badge-display.ts  # Badge rendering
│   │
│   ├── leaderboards/
│   │   ├── leaderboard-engine.ts  # Leaderboard system
│   │   │   # BE: users-be/leaderboards
│   │   │   # GET /v1/leaderboards/{type}
│   │   │
│   │   ├── ranking-algorithms.ts  # Ranking logic
│   │   └── time-windows.ts  # Time-based rankings
│   │
│   ├── points/
│   │   ├── points-system.ts  # Points management
│   │   │   # BE: users-be/points
│   │   │   # GET /v1/points/balance
│   │   │   # POST /v1/points/earn
│   │   │
│   │   ├── earning-rules.ts  # Point earning rules
│   │   └── redemption.ts  # Point redemption
│   │
│   └── hooks/
│       ├── useAchievements.ts  # Achievements hook
│       ├── useBadges.ts  # Badges hook
│       ├── useLeaderboard.ts  # Leaderboard hook
│       └── usePoints.ts  # Points hook
```

### 8. Moderation & Content Safety

```
packages/shared/src/
│
├── moderation/
│   ├── content-validation/
│   │   ├── profanity-filter.ts  # Profanity detection
│   │   ├── spam-detector.ts  # Spam detection
│   │   ├── link-validator.ts  # Link validation
│   │   └── format-validator.ts  # Format validation
│   │
│   ├── reporting/
│   │   ├── report-submission.ts  # Report submission
│   │   │   # BE: admin-be/moderation
│   │   │   # POST /v1/reports
│   │   │
│   │   ├── report-types.ts  # Report categories
│   │   └── evidence-collection.ts  # Evidence gathering
│   │
│   ├── auto-moderation/
│   │   ├── rule-engine.ts  # Auto-moderation rules
│   │   ├── action-executor.ts  # Automated actions
│   │   └── escalation.ts  # Escalation logic
│   │
│   └── hooks/
│       ├── useContentValidation.ts  # Validation hook
│       ├── useReporting.ts  # Reporting hook
│       └── useModerationStatus.ts  # Status hook
```

---

## IV. Mobile App - Enhanced Mobile Features

### 1. Mobile Onboarding Flow

```
apps/mobile/app/
│
├── onboarding/
│   ├── _layout.tsx  # Onboarding layout
│   │
│   ├── welcome.tsx  # Welcome screen
│   │   # - App introduction
│   │   # - Value propositions
│   │   # BE: None (static)
│   │
│   ├── permissions.tsx  # Permission requests
│   │   # - Camera access
│   │   # - Notifications
│   │   # - Location (optional)
│   │   # BE: None (device-level)
│   │
│   ├── user-type.tsx  # Select user type
│   │   # - Freelancer
│   │   # - Client
│   │   # BE: None (local state)
│   │
│   ├── profile-setup/
│   │   ├── basic-info.tsx  # Basic information
│   │   │   # BE: users-be/profile
│   │   │   # POST /v1/onboarding/profile
│   │   │
│   │   ├── skills.tsx  # Skills selection
│   │   │   # BE: users-be/skills
│   │   │   # POST /v1/onboarding/skills
│   │   │
│   │   ├── photo.tsx  # Profile photo
│   │   │   # BE: storage-be/asset, users-be/profile
│   │   │   # POST /v1/storage/uploads
│   │   │   # PUT /v1/users/me/photo
│   │   │
│   │   └── preferences.tsx  # Initial preferences
│   │       # BE: users-be/preferences
│   │       # POST /v1/onboarding/preferences
│   │
│   ├── verification.tsx  # Identity verification
│   │   # - Document upload
│   │   # - Selfie verification
│   │   # BE: admin-be/kyc
│   │   # POST /v1/kyc/submit
│   │
│   ├── notifications-setup.tsx  # Notification preferences
│   │   # BE: communications-be/preferences
│   │   # POST /v1/notifications/preferences
│   │
│   └── complete.tsx  # Onboarding complete
│       # - Success message
│       # - Next steps
│       # BE: users-be/onboarding
│       # POST /v1/onboarding/complete
```

### 2. Mobile Quick Actions

```
apps/mobile/app/
│
├── quick-actions/
│   ├── quick-apply.tsx  # Quick job application
│   │   # - Minimal form
│   │   # - Saved templates
│   │   # - Quick submit
│   │   # BE: proposals-be/proposal
│   │   # POST /v1/proposals/quick-apply
│   │
│   ├── quick-message.tsx  # Quick messaging
│   │   # - Contact from anywhere
│   │   # - Voice messages
│   │   # - Quick replies
│   │   # BE: communications-be/messages
│   │   # POST /v1/messages/quick-send
│   │
│   ├── quick-time-entry.tsx  # Quick time logging
│   │   # - One-tap timer
│   │   # - Recent tasks
│   │   # - Auto-fill
│   │   # BE: contracts-be/workdiary
│   │   # POST /v1/contracts/time-entry/quick
│   │
│   └── quick-invoice.tsx  # Quick invoice creation
│       # - Templates
│       # - Recent clients
│       # - Quick send
│       # BE: financial-be/invoice
│       # POST /v1/invoices/quick-create
```

### 3. Mobile Settings & Preferences

```
apps/mobile/app/
│
├── settings/
│   ├── app-settings/
│   │   ├── appearance.tsx  # Appearance settings
│   │   │   # - Theme (light/dark/auto)
│   │   │   # - Font size
│   │   │   # - Display density
│   │   │   # BE: None (local storage)
│   │   │
│   │   ├── biometric-auth.tsx  # Biometric authentication
│   │   │   # - Face ID
│   │   │   # - Fingerprint
│   │   │   # - Setup/disable
│   │   │   # BE: None (device-level)
│   │   │
│   │   ├── haptics.tsx  # Haptic feedback
│   │   │   # - Enable/disable
│   │   │   # - Intensity
│   │   │   # BE: None (local storage)
│   │   │
│   │   ├── quick-actions.tsx  # Quick action config
│   │   │   # - Customize quick actions
│   │   │   # - Shortcuts
│   │   │   # BE: None (local storage)
│   │   │
│   │   └── offline-mode.tsx  # Offline settings
│   │       # - Auto-sync
│   │       # - Storage limits
│   │       # - Download quality
│   │       # BE: None (local storage)
│   │
│   └── mobile-specific/
│       ├── data-usage.tsx  # Data usage settings
│       │   # - WiFi only mode
│       │   # - Image quality
│       │   # - Video autoplay
│       │   # BE: None (local storage)
│       │
│       ├── battery-optimization.tsx  # Battery settings
│       │   # - Power saving mode
│       │   # - Background sync
│       │   # BE: None (local storage)
│       │
│       └── cache-management.tsx  # Cache management
│           # - Clear cache
│           # - Cache size
│           # BE: None (local storage)
```

---

## V. UI Components - Advanced & Specialized

### 1. Advanced Form Components

```
packages/ui/src/components/
│
├── Form/
│   ├── RichTextEditor/
│   │   ├── RichTextEditor.tsx  # Base editor
│   │   ├── RichTextEditor.web.tsx  # Web (TipTap/Slate)
│   │   ├── RichTextEditor.native.tsx  # Native implementation
│   │   ├── toolbar/
│   │   │   ├── Toolbar.tsx
│   │   │   ├── FormatButtons.tsx
│   │   │   └── InsertButtons.tsx
│   │   └── RichTextEditor.types.ts
│   │
│   ├── CodeEditor/
│   │   ├── CodeEditor.tsx  # Base code editor
│   │   ├── CodeEditor.web.tsx  # Monaco/CodeMirror
│   │   ├── CodeEditor.native.tsx  # Native code editor
│   │   ├── syntax-highlighting/
│   │   │   ├── languages.ts
│   │   │   └── themes.ts
│   │   └── CodeEditor.types.ts
│   │
│   ├── MarkdownEditor/
│   │   ├── MarkdownEditor.tsx
│   │   ├── MarkdownEditor.web.tsx
│   │   ├── MarkdownEditor.native.tsx
│   │   ├── preview/
│   │   │   └── MarkdownPreview.tsx
│   │   └── MarkdownEditor.types.ts
│   │
│   ├── SignaturePad/
│   │   ├── SignaturePad.tsx
│   │   ├── SignaturePad.web.tsx  # Canvas-based
│   │   ├── SignaturePad.native.tsx  # React Native Skia
│   │   └── SignaturePad.types.ts
│   │
│   └── DateRangePicker/
│       ├── DateRangePicker.tsx
│       ├── DateRangePicker.web.tsx
│       ├── DateRangePicker.native.tsx
│       ├── presets/
│       │   └── DatePresets.tsx
│       └── DateRangePicker.types.ts
```

### 2. Data Visualization Components

```
packages/ui/src/components/
│
├── Visualization/
│   ├── HeatMap/
│   │   ├── HeatMap.tsx
│   │   ├── HeatMap.web.tsx
│   │   ├── HeatMap.native.tsx
│   │   └── HeatMap.types.ts
│   │
│   ├── GanttChart/
│   │   ├── GanttChart.tsx
│   │   ├── GanttChart.web.tsx
│   │   ├── GanttChart.native.tsx
│   │   ├── timeline/
│   │   │   ├── Timeline.tsx
│   │   │   └── TimelineItem.tsx
│   │   └── GanttChart.types.ts
│   │
│   ├── KanbanBoard/
│   │   ├── KanbanBoard.tsx
│   │   ├── KanbanBoard.web.tsx
│   │   ├── KanbanBoard.native.tsx
│   │   ├── column/
│   │   │   └── KanbanColumn.tsx
│   │   ├── card/
│   │   │   └── KanbanCard.tsx
│   │   └── KanbanBoard.types.ts
│   │
│   ├── TreeView/
│   │   ├── TreeView.tsx
│   │   ├── TreeView.web.tsx
│   │   ├── TreeView.native.tsx
│   │   ├── node/
│   │   │   └── TreeNode.tsx
│   │   └── TreeView.types.ts
│   │
│   └── OrgChart/
│       ├── OrgChart.tsx
│       ├── OrgChart.web.tsx
│       ├── OrgChart.native.tsx
│       ├── node/
│       │   └── OrgNode.tsx
│       └── OrgChart.types.ts
```

### 3. AI-Powered Components

```
packages/ui/src/components/
│
├── AI/
│   ├── AIAssistant/
│   │   ├── AIAssistant.tsx  # AI chat assistant
│   │   ├── AIAssistant.web.tsx
│   │   ├── AIAssistant.native.tsx
│   │   ├── chat/
│   │   │   ├── ChatBubble.tsx
│   │   │   └── ChatInput.tsx
│   │   └── AIAssistant.types.ts
│   │
│   ├── SmartSuggestions/
│   │   ├── SmartSuggestions.tsx  # AI suggestions
│   │   ├── SmartSuggestions.web.tsx
│   │   ├── SmartSuggestions.native.tsx
│   │   └── SmartSuggestions.types.ts
│   │
│   ├── AutoComplete/
│   │   ├── AIAutoComplete.tsx  # AI-powered autocomplete
│   │   ├── AIAutoComplete.web.tsx
│   │   ├── AIAutoComplete.native.tsx
│   │   └── AIAutoComplete.types.ts
│   │
│   └── ContentGeneration/
│       ├── ContentGenerator.tsx  # AI content generation
│       ├── templates/
│       │   └── GenerationTemplates.tsx
│       └── ContentGeneration.types.ts
```

---

## VI. Testing Infrastructure

### 1. Test Utilities

```
packages/shared/src/testing/
│
├── test-utils/
│   ├── render-utils.tsx  # React Testing Library utils
│   ├── mock-providers.tsx  # Mock providers
│   ├── test-wrapper.tsx  # Test wrapper component
│   └── custom-matchers.ts  # Custom Jest matchers
│
├── mock-data/
│   ├── users.ts  # Mock user data
│   ├── jobs.ts  # Mock job data
│   ├── proposals.ts  # Mock proposal data
│   ├── contracts.ts  # Mock contract data
│   ├── payments.ts  # Mock payment data
│   └── factories/
│       ├── user-factory.ts  # User factory
│       ├── job-factory.ts  # Job factory
│       └── proposal-factory.ts  # Proposal factory
│
├── mocks/
│   ├── api/
│   │   ├── handlers.ts  # MSW handlers
│   │   ├── server.ts  # MSW server
│   │   └── responses/
│   │       ├── success-responses.ts
│   │       └── error-responses.ts
│   │
│   ├── storage/
│   │   └── local-storage-mock.ts
│   │
│   └── websocket/
│       └── websocket-mock.ts
│
└── setup/
    ├── jest.setup.ts  # Jest configuration
    ├── test-env.ts  # Test environment
    └── global-setup.ts  # Global test setup
```

### 2. E2E Test Structure

```
apps/web/
│
├── e2e/
│   ├── tests/
│   │   ├── auth/
│   │   │   ├── login.spec.ts
│   │   │   ├── registration.spec.ts
│   │   │   └── password-reset.spec.ts
│   │   │
│   │   ├── jobs/
│   │   │   ├── job-posting.spec.ts
│   │   │   ├── job-search.spec.ts
│   │   │   └── job-application.spec.ts
│   │   │
│   │   ├── proposals/
│   │   │   ├── proposal-submission.spec.ts
│   │   │   ├── proposal-negotiation.spec.ts
│   │   │   └── proposal-acceptance.spec.ts
│   │   │
│   │   ├── contracts/
│   │   │   ├── contract-creation.spec.ts
│   │   │   ├── milestone-tracking.spec.ts
│   │   │   └── deliverable-submission.spec.ts
│   │   │
│   │   └── payments/
│   │       ├── payment-methods.spec.ts
│   │       ├── invoice-generation.spec.ts
│   │       └── payout-processing.spec.ts
│   │
│   ├── fixtures/
│   │   ├── users.json
│   │   ├── jobs.json
│   │   └── contracts.json
│   │
│   ├── page-objects/
│   │   ├── auth/
│   │   │   ├── LoginPage.ts
│   │   │   └── RegisterPage.ts
│   │   │
│   │   ├── jobs/
│   │   │   ├── JobListPage.ts
│   │   │   └── JobDetailPage.ts
│   │   │
│   │   └── dashboard/
│   │       └── DashboardPage.ts
│   │
│   └── config/
│       ├── playwright.config.ts
│       └── test-ids.ts
```

---

## VII. Documentation Structure

### 1. API Integration Documentation

```
docs/
│
├── api-integration/
│   ├── authentication/
│   │   ├── keycloak-integration.md
│   │   ├── token-management.md
│   │   └── refresh-flow.md
│   │
│   ├── microservices/
│   │   ├── users-be-integration.md
│   │   ├── jobs-be-integration.md
│   │   ├── proposals-be-integration.md
│   │   ├── contracts-be-integration.md
│   │   ├── financial-be-integration.md
│   │   ├── communications-be-integration.md
│   │   ├── search-be-integration.md
│   │   ├── storage-be-integration.md
│   │   ├── reviews-be-integration.md
│   │   ├── admin-be-integration.md
│   │   └── subscriptions-be-integration.md
│   │
│   ├── error-handling/
│   │   ├── error-codes.md
│   │   ├── retry-strategies.md
│   │   └── fallback-mechanisms.md
│   │
│   └── best-practices/
│       ├── api-client-patterns.md
│       ├── caching-strategies.md
│       └── rate-limiting.md
```

### 2. Component Documentation

```
docs/
│
├── components/
│   ├── overview.md  # Component library overview
│   │
│   ├── design-system/
│   │   ├── design-tokens.md
│   │   ├── color-system.md
│   │   ├── typography.md
│   │   ├── spacing.md
│   │   └── breakpoints.md
│   │
│   ├── usage-guidelines/
│   │   ├── component-composition.md
│   │   ├── theming-guide.md
│   │   ├── responsive-design.md
│   │   └── accessibility.md
│   │
│   └── storybook/
│       ├── setup.md
│       ├── writing-stories.md
│       └── testing-stories.md
```

---

## VIII. CI/CD Configuration

### 1. GitHub Actions Workflows

```
.github/
│
├── workflows/
│   ├── ci.yml  # Main CI pipeline
│   │   # - Lint
│   │   # - Type check
│   │   # - Unit tests
│   │   # - Build
│   │
│   ├── web-deploy-staging.yml  # Web staging deployment
│   ├── web-deploy-production.yml  # Web production deployment
│   ├── mobile-build-android.yml  # Android build
│   ├── mobile-build-ios.yml  # iOS build
│   │
│   ├── e2e-tests.yml  # E2E test pipeline
│   ├── visual-regression.yml  # Visual regression tests
│   ├── performance-tests.yml  # Performance testing
│   │
│   ├── security-scan.yml  # Security scanning
│   ├── dependency-update.yml  # Automated dependency updates
│   │
│   └── release.yml  # Release automation
│
└── actions/
    ├── setup-node/
    │   └── action.yml
    ├── setup-pnpm/
    │   └── action.yml
    └── notify-deployment/
        └── action.yml
```

### 2. Environment Configuration

```
.github/
│
├── environments/
│   ├── staging.yml  # Staging environment config
│   ├── production.yml  # Production environment config
│   └── development.yml  # Development environment config
│
└── secrets/
    └── README.md  # Secrets documentation
```

---

## IX. Monitoring & Observability

### 1. Error Tracking

```
packages/shared/src/monitoring/
│
├── error-tracking/
│   ├── sentry-config.ts  # Sentry configuration
│   ├── error-boundary-setup.ts  # Error boundary setup
│   ├── error-enrichment.ts  # Context enrichment
│   └── custom-integrations.ts  # Custom integrations
│
├── performance/
│   ├── web-vitals-reporter.ts  # Web Vitals reporting
│   ├── custom-metrics.ts  # Custom performance metrics
│   └── profiling.ts  # Performance profiling
│
└── analytics/
    ├── segment-config.ts  # Segment configuration
    ├── mixpanel-config.ts  # Mixpanel configuration
    └── event-tracking.ts  # Event tracking setup
```

### 2. Logging Infrastructure

```
packages/shared/src/logging/
│
├── logger.ts  # Base logger
├── log-levels.ts  # Log level configuration
├── formatters/
│   ├── json-formatter.ts
│   ├── pretty-formatter.ts
│   └── structured-formatter.ts
│
├── transports/
│   ├── console-transport.ts
│   ├── file-transport.ts
│   └── remote-transport.ts
│
└── context/
    ├── request-context.ts  # Request context
    ├── user-context.ts  # User context
    └── trace-context.ts  # Trace context
```

---

## X. Security Infrastructure

### 1. Security Utilities

```
packages/shared/src/security/
│
├── csrf/
│   ├── csrf-token.ts  # CSRF token management
│   ├── csrf-validation.ts  # CSRF validation
│   └── double-submit-cookie.ts  # Double submit pattern
│
├── sanitization/
│   ├── html-sanitizer.ts  # HTML sanitization
│   ├── input-sanitizer.ts  # Input sanitization
│   ├── xss-prevention.ts  # XSS prevention
│   └── sql-injection-prevention.ts  # SQL injection prevention
│
├── encryption/
│   ├── client-encryption.ts  # Client-side encryption
│   ├── key-management.ts  # Key management
│   └── secure-storage.ts  # Secure storage
│
└── validation/
    ├── input-validation.ts  # Input validation
    ├── schema-validation.ts  # Schema validation
    └── file-validation.ts  # File validation
```

### 2. Security Headers & Policies

```
apps/web/
│
├── middleware.ts  # Next.js middleware for security
│
└── security/
    ├── headers/
    │   ├── csp.ts  # Content Security Policy
    │   ├── hsts.ts  # HTTP Strict Transport Security
    │   ├── x-frame-options.ts  # X-Frame-Options
    │   └── permissions-policy.ts  # Permissions Policy
    │
    ├── cors/
    │   ├── cors-config.ts  # CORS configuration
    │   └── origin-validation.ts  # Origin validation
    │
    └── rate-limiting/
        ├── rate-limiter.ts  # Rate limiting
        └── ddos-protection.ts  # DDoS protection
```

---

## Summary

This document provides the **final comprehensive missing folder structure** that covers all remaining requirements from `fe-folder-structure-prompt.md` that were not present in any previous documents.

### Key Areas Covered:

#### Dashboard Routes (Deep Features):
1. **Proposals** - Advanced negotiation, version control, milestones, questions, collaborators, insights (win rate, pricing, response time), portfolio showcases
2. **Contracts** - Work diary (bulk entry, screenshots, activity levels, corrections), milestones (submission, review, revisions), deliverables (versions, feedback, approvals), change requests, compliance (documents, audits, reports), risks, quality, knowledge transfer, benchmarking
3. **Financial** - Treasury (dashboard, cash flow, liquidity, investments), risk management (exposure, limits, reports), reconciliation (bank, merchant, intercompany), credit management (limits, scoring, collections), budgets, forecasting, compliance (AML, KYC, sanctions), analytics (profitability, margins, KPIs)
4. **Teams** - Hierarchy (org chart, reporting lines), capacity planning, workflows, policies, training, procurement, knowledge base, cross-team collaboration
5. **Settings** - Developer tools (API keys, webhooks, OAuth apps, sandbox), advanced settings (data retention, audit logs, IP whitelist, sessions, rate limiting), labs

#### Admin Panel (Enterprise):
1. **Financial Ops** - Settlement (batches, holds, rules), reserves, fraud detection, currency management, fee structures
2. **Trust & Safety** - Content moderation (queue, ML assistance, automation, appeals), risk scoring, identity verification, watchlists
3. **Platform Operations** - Search quality (relevance, synonyms, boosts, reindexing), messaging (templates, campaigns, rate limits, deliverability), storage (quotas, lifecycle, virus scanning, CDN), feature flags & experiments

#### Shared Features (Infrastructure):
1. **Real-time** - WebSocket, presence, typing indicators, live updates
2. **Analytics** - Collectors, enrichment, batching, sampling, privacy
3. **Performance** - Web Vitals, resource timing, custom metrics, monitoring
4. **Offline** - Operation queue, storage, sync engine, network monitoring
5. **Geolocation** - IP detection, browser API, timezone, country detection
6. **Experiments** - A/B testing, feature variants, conversion tracking
7. **Gamification** - Achievements, badges, leaderboards, points
8. **Moderation** - Content validation, reporting, auto-moderation

#### Mobile Enhancements:
1. **Onboarding** - Welcome, permissions, user type, profile setup, verification
2. **Quick Actions** - Quick apply, quick message, quick time entry, quick invoice
3. **Mobile Settings** - Appearance, biometric auth, haptics, offline mode, data usage, battery optimization

#### UI Components:
1. **Advanced Forms** - Rich text editor, code editor, markdown editor, signature pad, date range picker
2. **Visualization** - Heat map, Gantt chart, Kanban board, tree view, org chart
3. **AI Components** - AI assistant, smart suggestions, AI autocomplete, content generation

#### Infrastructure:
1. **Testing** - Test utilities, mock data, E2E structure
2. **Documentation** - API integration, component docs
3. **CI/CD** - GitHub Actions, environments
4. **Monitoring** - Error tracking, logging
5. **Security** - Security utilities, headers & policies

All routes and components include comprehensive backend mappings with microservice, domain, HTTP method, and endpoint information as specified in the requirements.
