# Skillsier Frontend - Complete Folder Structure
## Part 2: Profile, Jobs, Proposals & Search Modules

> **CRITICAL**: This document contains ONLY the folder structure, filenames, and inline backend API mappings.  
> **NO CODE IMPLEMENTATIONS** are included per the strict output policy.

---

## Profile Management Module

```
apps/web/src/app/[locale]/(dashboard)/profile/
│
├── page.tsx                                # Profile overview / public view
│                                            # - Profile header (photo, name, title, location)
│                                            # - Stats (rating, jobs completed, earnings)
│                                            # - Skills showcase
│                                            # - Portfolio highlights
│                                            # - Recent reviews
│                                            # - Availability calendar
│                                            # BE: users-be/profile
│                                            # GET /v1/users/{id}/profile
│                                            # BE: users-be/capabilities
│                                            # GET /v1/users/{id}/skills
│                                            # BE: users-be/portfolio
│                                            # GET /v1/users/{id}/portfolio
│                                            # BE: reviews-be
│                                            # GET /v1/reviews?user_id={id}&limit=5
│                                            # BE: users-be/availability
│                                            # GET /v1/users/{id}/availability
│
├── edit/
│   └── page.tsx                            # Edit profile form
│                                            # - Basic info (name, title, bio)
│                                            # - Profile photo upload
│                                            # - Location
│                                            # - Languages
│                                            # - Hourly rate (freelancer)
│                                            # - Professional headline
│                                            # BE: users-be/profile
│                                            # PATCH /v1/users/{id}/profile
│                                            # BE: storage-be/uploads
│                                            # POST /v1/storage/upload (photo)
│
├── skills/
│   ├── page.tsx                            # Skills management
│   │                                        # - List current skills with levels
│   │                                        # - Add new skills (autocomplete)
│   │                                        # - Edit skill proficiency
│   │                                        # - Remove skills
│   │                                        # - Primary skills (max 5)
│   │                                        # BE: users-be/capabilities
│   │                                        # GET /v1/users/{id}/skills
│   │                                        # POST /v1/users/{id}/skills
│   │                                        # PUT /v1/users/{id}/skills/{skill_id}
│   │                                        # DELETE /v1/users/{id}/skills/{skill_id}
│   │                                        # BE: search-be/taxonomy
│   │                                        # GET /v1/taxonomy/skills (autocomplete)
│   │
│   └── specializations/
│       └── page.tsx                        # Specializations & niche expertise
│                                            # - Add specializations
│                                            # - Verification status
│                                            # - Niche expertise tags
│                                            # BE: users-be/capabilities
│                                            # GET /v1/users/{id}/specializations
│                                            # POST /v1/users/{id}/specializations
│
├── experience/
│   ├── page.tsx                            # Work experience list
│   │                                        # - List all experience entries
│   │                                        # - Add new experience button
│   │                                        # - Edit/delete actions
│   │                                        # BE: users-be/experience
│   │                                        # GET /v1/users/{id}/experience
│   │
│   ├── add/
│   │   └── page.tsx                        # Add experience form
│   │                                        # - Company name
│   │                                        # - Position title
│   │                                        # - Start/end dates
│   │                                        # - Current position checkbox
│   │                                        # - Description (rich text)
│   │                                        # - Skills used
│   │                                        # BE: users-be/experience
│   │                                        # POST /v1/users/{id}/experience
│   │
│   └── [experienceId]/
│       └── edit/
│           └── page.tsx                    # Edit experience form
│                                            # BE: users-be/experience
│                                            # PUT /v1/users/{id}/experience/{exp_id}
│                                            # DELETE /v1/users/{id}/experience/{exp_id}
│
├── education/
│   ├── page.tsx                            # Education list
│   │                                        # BE: users-be/education
│   │                                        # GET /v1/users/{id}/education
│   │
│   ├── add/
│   │   └── page.tsx                        # Add education form
│   │                                        # - School/university
│   │                                        # - Degree/qualification
│   │                                        # - Field of study
│   │                                        # - Graduation year
│   │                                        # - GPA (optional)
│   │                                        # - Description
│   │                                        # BE: users-be/education
│   │                                        # POST /v1/users/{id}/education
│   │
│   └── [educationId]/
│       └── edit/
│           └── page.tsx                    # Edit education form
│                                            # BE: users-be/education
│                                            # PUT /v1/users/{id}/education/{edu_id}
│                                            # DELETE /v1/users/{id}/education/{edu_id}
│
├── certifications/
│   ├── page.tsx                            # Certifications list
│   │                                        # - External certifications
│   │                                        # - Platform certifications
│   │                                        # - Badges earned
│   │                                        # BE: users-be/credentials
│   │                                        # GET /v1/users/{id}/credentials
│   │
│   ├── add/
│   │   └── page.tsx                        # Add certification
│   │                                        # - Certification name
│   │                                        # - Issuing organization
│   │                                        # - Issue date
│   │                                        # - Expiry date (if any)
│   │                                        # - Credential ID
│   │                                        # - Credential URL
│   │                                        # - Certificate upload
│   │                                        # BE: users-be/credentials
│   │                                        # POST /v1/users/{id}/certifications
│   │                                        # BE: storage-be/uploads
│   │
│   └── verify/
│       └── [certificationId]/
│           └── page.tsx                    # Verification request
│                                            # - Submit for verification
│                                            # - Upload proof
│                                            # BE: users-be/credentials
│                                            # POST /v1/users/{id}/certifications/{cert_id}/verify
│
├── portfolio/
│   ├── page.tsx                            # Portfolio items list
│   │                                        # - Grid/list view
│   │                                        # - Featured items
│   │                                        # - Reorder items (drag & drop)
│   │                                        # BE: users-be/portfolio
│   │                                        # GET /v1/users/{id}/portfolio
│   │
│   ├── add/
│   │   └── page.tsx                        # Add portfolio item
│   │                                        # - Project title
│   │                                        # - Description
│   │                                        # - Media upload (images/videos)
│   │                                        # - Project URL
│   │                                        # - Skills used
│   │                                        # - Client (optional)
│   │                                        # - Completion date
│   │                                        # BE: users-be/portfolio
│   │                                        # POST /v1/users/{id}/portfolio
│   │                                        # BE: storage-be/uploads
│   │                                        # POST /v1/storage/upload (media)
│   │
│   ├── [portfolioId]/
│   │   ├── page.tsx                        # Portfolio item detail
│   │   │                                    # BE: users-be/portfolio
│   │   │                                    # GET /v1/users/{id}/portfolio/{item_id}
│   │   │
│   │   └── edit/
│   │       └── page.tsx                    # Edit portfolio item
│   │                                        # BE: users-be/portfolio
│   │                                        # PUT /v1/users/{id}/portfolio/{item_id}
│   │                                        # DELETE /v1/users/{id}/portfolio/{item_id}
│   │
│   └── reorder/
│       └── page.tsx                        # Reorder portfolio items
│                                            # - Drag & drop interface
│                                            # - Set featured items
│                                            # BE: users-be/portfolio
│                                            # PATCH /v1/users/{id}/portfolio/reorder
│                                            # Body: { item_ids: string[] }
│
├── service-catalog/
│   ├── page.tsx                            # Service catalog management (freelancer)
│   │                                        # - List offered services
│   │                                        # - Service packages
│   │                                        # - Pricing tiers
│   │                                        # BE: users-be/service_catalog
│   │                                        # GET /v1/users/{id}/service-catalog
│   │
│   ├── add/
│   │   └── page.tsx                        # Add service
│   │                                        # - Service name
│   │                                        # - Description
│   │                                        # - Capabilities required
│   │                                        # - Delivery time
│   │                                        # - Pricing
│   │                                        # - Packages (Basic/Standard/Premium)
│   │                                        # BE: users-be/service_catalog
│   │                                        # POST /v1/users/{id}/services
│   │
│   └── [serviceId]/
│       └── edit/
│           └── page.tsx                    # Edit service
│                                            # BE: users-be/service_catalog
│                                            # PUT /v1/users/{id}/services/{service_id}
│                                            # DELETE /v1/users/{id}/services/{service_id}
│
├── availability/
│   └── page.tsx                            # Availability management
│                                            # - Calendar view
│                                            # - Set available hours
│                                            # - Time zone
│                                            # - Vacation mode
│                                            # - Max concurrent projects
│                                            # BE: users-be/availability
│                                            # GET /v1/users/{id}/availability
│                                            # POST /v1/users/{id}/availability
│                                            # PATCH /v1/users/{id}/availability
│
├── verification/
│   ├── page.tsx                            # Verification status
│   │                                        # - Email verified
│   │                                        # - Phone verified
│   │                                        # - ID verification status
│   │                                        # - Payment method verified
│   │                                        # BE: users-be/verification
│   │                                        # GET /v1/users/{id}/verification-status
│   │
│   ├── phone/
│   │   └── page.tsx                        # Phone verification
│   │                                        # - Enter phone number
│   │                                        # - Receive OTP
│   │                                        # - Verify OTP
│   │                                        # BE: users-be/verification
│   │                                        # POST /v1/users/{id}/verify-phone/send
│   │                                        # POST /v1/users/{id}/verify-phone/verify
│   │
│   └── identity/
│       └── page.tsx                        # ID verification
│                                            # - Upload ID document
│                                            # - Selfie verification
│                                            # - Address proof
│                                            # BE: users-be/verification/kyc
│                                            # POST /v1/users/{id}/kyc/submit
│                                            # BE: storage-be/uploads
│                                            # BE: admin-be/kyc_case (creates case)
│
└── public-view/
    └── [userId]/
        └── page.tsx                        # Public profile view (for others to see)
                                            # - Same as profile overview but for any user
                                            # - "Hire" button (if freelancer)
                                            # - "Contact" button
                                            # BE: users-be/profile
                                            # GET /v1/users/{user_id}/profile
```

---

## Jobs Module

```
apps/web/src/app/[locale]/(dashboard)/jobs/
│
├── page.tsx                                # Jobs list (role-based)
│                                            # Freelancer view: Browse available jobs
│                                            # Client view: My posted jobs
│                                            # - Filters (category, budget, skills, etc.)
│                                            # - Search bar
│                                            # - Sort options
│                                            # - Pagination
│                                            # BE: jobs-be/job
│                                            # GET /v1/jobs?filters=...&page=1&limit=20
│                                            # Freelancer: GET /v1/jobs/browse
│                                            # Client: GET /v1/jobs/my-jobs
│
├── browse/                                 # Browse jobs (freelancer)
│   ├── page.tsx                            # Job listings with filters
│   │                                        # - Category filters
│   │                                        # - Budget range
│   │                                        # - Experience level
│   │                                        # - Job type (fixed/hourly)
│   │                                        # - Location preferences
│   │                                        # - Skills required
│   │                                        # - Posted date
│   │                                        # - Saved jobs indicator
│   │                                        # - "Best Matches" tab
│   │                                        # BE: jobs-be/job
│   │                                        # GET /v1/jobs/browse?filters={...}
│   │                                        # BE: search-be/query
│   │                                        # POST /v1/search/jobs (for advanced search)
│   │
│   ├── saved/
│   │   └── page.tsx                        # Saved/bookmarked jobs
│   │                                        # BE: jobs-be/saved_jobs
│   │                                        # GET /v1/jobs/saved
│   │                                        # DELETE /v1/jobs/saved/{job_id}
│   │
│   └── invitations/
│       └── page.tsx                        # Job invitations received
│                                            # BE: jobs-be/invitations
│                                            # GET /v1/jobs/invitations
│
├── post/
│   └── page.tsx                            # Post a new job (client)
│                                            # - Job title
│                                            # - Job description (rich text editor)
│                                            # - Category selection
│                                            # - Required skills (autocomplete)
│                                            # - Experience level
│                                            # - Job type (fixed/hourly)
│                                            # - Budget/rate
│                                            # - Duration
│                                            # - Attachments
│                                            # - Screening questions
│                                            # - Visibility (public/private/invited)
│                                            # - Save as draft
│                                            # BE: jobs-be/job
│                                            # POST /v1/jobs
│                                            # Body: { title, description, category_id, skills, budget, ... }
│                                            # BE: jobs-be/attachments
│                                            # POST /v1/jobs/{job_id}/attachments
│                                            # BE: jobs-be/screening
│                                            # POST /v1/jobs/{job_id}/screening-questions
│                                            # BE: storage-be/uploads
│                                            # Publishes: JobPosted event
│
├── drafts/
│   ├── page.tsx                            # Job drafts list
│   │                                        # BE: jobs-be/draft
│   │                                        # GET /v1/jobs/drafts
│   │
│   └── [draftId]/
│       └── edit/
│           └── page.tsx                    # Edit draft
│                                            # BE: jobs-be/draft
│                                            # PUT /v1/jobs/drafts/{draft_id}
│                                            # DELETE /v1/jobs/drafts/{draft_id}
│
├── [jobId]/
│   ├── page.tsx                            # Job detail page
│   │                                        # - Full job description
│   │                                        # - Client info
│   │                                        # - Skills required
│   │                                        # - Budget/rate
│   │                                        # - Proposals count
│   │                                        # - Similar jobs
│   │                                        # - "Submit Proposal" button (freelancer)
│   │                                        # - Save job button
│   │                                        # BE: jobs-be/job
│   │                                        # GET /v1/jobs/{job_id}
│   │                                        # BE: proposals-be
│   │                                        # GET /v1/proposals/count?job_id={job_id}
│   │                                        # BE: search-be/similarity
│   │                                        # GET /v1/similarity/jobs/{job_id}
│   │
│   ├── edit/
│   │   └── page.tsx                        # Edit job (client only)
│   │                                        # - Same form as post job
│   │                                        # - Cannot edit if has accepted proposals
│   │                                        # BE: jobs-be/job
│   │                                        # PUT /v1/jobs/{job_id}
│   │                                        # Publishes: JobUpdated event
│   │
│   ├── proposals/
│   │   ├── page.tsx                        # Proposals received (client)
│   │   │                                    # - List all proposals
│   │   │                                    # - Filter (all/shortlisted/archived)
│   │   │                                    # - Sort (date, rate, rating)
│   │   │                                    # - Proposal cards with key info
│   │   │                                    # BE: proposals-be
│   │   │                                    # GET /v1/proposals?job_id={job_id}
│   │   │
│   │   └── [proposalId]/
│   │       └── page.tsx                    # Proposal detail
│   │                                        # - Full proposal view
│   │                                        # - Freelancer profile preview
│   │                                        # - Accept/Reject buttons
│   │                                        # - Shortlist button
│   │                                        # - Message freelancer
│   │                                        # BE: proposals-be
│   │                                        # GET /v1/proposals/{proposal_id}
│   │                                        # POST /v1/proposals/{proposal_id}/accept
│   │                                        # POST /v1/proposals/{proposal_id}/reject
│   │                                        # POST /v1/proposals/{proposal_id}/shortlist
│   │
│   ├── invite/
│   │   └── page.tsx                        # Invite freelancers (client)
│   │                                        # - Search freelancers
│   │                                        # - Send invitation with message
│   │                                        # BE: jobs-be/invitations
│   │                                        # POST /v1/jobs/{job_id}/invitations
│   │                                        # BE: search-be
│   │                                        # POST /v1/search/freelancers
│   │                                        # BE: communications-be
│   │                                        # Publishes: JobInvitationSent event
│   │
│   ├── analytics/
│   │   └── page.tsx                        # Job analytics (client)
│   │                                        # - Views
│   │                                        # - Proposals received
│   │                                        # - Proposal conversion rate
│   │                                        # - Time to hire
│   │                                        # BE: jobs-be/analytics
│   │                                        # GET /v1/jobs/{job_id}/analytics
│   │
│   └── close/
│       └── page.tsx                        # Close job
│                                            # - Reason for closing
│                                            # - Notify applicants
│                                            # BE: jobs-be/job
│                                            # POST /v1/jobs/{job_id}/close
│                                            # Publishes: JobClosed event
│
├── my-jobs/                                # Client's posted jobs
│   ├── page.tsx                            # All posted jobs
│   │                                        # - Active jobs
│   │                                        # - Closed jobs
│   │                                        # - Drafts
│   │                                        # BE: jobs-be/job
│   │                                        # GET /v1/jobs/my-jobs?status=active
│   │
│   ├── active/
│   │   └── page.tsx                        # Active jobs only
│   │                                        # BE: jobs-be/job
│   │                                        # GET /v1/jobs/my-jobs?status=active
│   │
│   └── closed/
│       └── page.tsx                        # Closed jobs
│                                            # BE: jobs-be/job
│                                            # GET /v1/jobs/my-jobs?status=closed
│
├── categories/
│   ├── page.tsx                            # Browse by category
│   │                                        # - Category grid
│   │                                        # - Subcategories
│   │                                        # BE: jobs-be/categories
│   │                                        # GET /v1/jobs/categories
│   │
│   └── [categoryId]/
│       └── page.tsx                        # Jobs in category
│                                            # BE: jobs-be/job
│                                            # GET /v1/jobs?category_id={category_id}
│
└── recommendations/
    └── page.tsx                            # Recommended jobs (freelancer)
                                            # - ML-powered job recommendations
                                            # - Based on skills, history, preferences
                                            # BE: search-be/recommendations
                                            # GET /v1/recommendations/jobs
```

---

## Proposals Module

```
apps/web/src/app/[locale]/(dashboard)/proposals/
│
├── page.tsx                                # Proposals list
│                                            # Freelancer: Submitted proposals
│                                            # Client: Received proposals (redirect to jobs)
│                                            # BE: proposals-be
│                                            # GET /v1/proposals/my-proposals
│
├── submit/
│   └── [jobId]/
│       └── page.tsx                        # Submit proposal (freelancer)
│                                            # - Cover letter (required)
│                                            # - Proposed rate/budget
│                                            # - Proposed timeline
│                                            # - Answer screening questions
│                                            # - Attachments (portfolio samples)
│                                            # - Milestones (for fixed-price)
│                                            # - Terms acceptance
│                                            # - Connects deduction warning
│                                            # BE: proposals-be
│                                            # POST /v1/proposals
│                                            # Body: { job_id, cover_letter, rate, timeline, ... }
│                                            # BE: subscriptions-be/connects
│                                            # POST /v1/connects/deduct
│                                            # BE: storage-be/uploads
│                                            # Publishes: ProposalSubmitted event
│
├── [proposalId]/
│   ├── page.tsx                            # Proposal detail
│   │                                        # - View submitted proposal
│   │                                        # - Proposal status
│   │                                        # - Client messages/feedback
│   │                                        # - Withdraw option (if pending)
│   │                                        # BE: proposals-be
│   │                                        # GET /v1/proposals/{proposal_id}
│   │
│   ├── edit/
│   │   └── page.tsx                        # Edit proposal
│   │                                        # - Only if status = DRAFT or PENDING
│   │                                        # - Update cover letter, rate, timeline
│   │                                        # BE: proposals-be
│   │                                        # PUT /v1/proposals/{proposal_id}
│   │
│   └── withdraw/
│       └── page.tsx                        # Withdraw proposal
│                                            # - Confirmation dialog
│                                            # - Reason for withdrawal
│                                            # - Connects refund info
│                                            # BE: proposals-be
│                                            # POST /v1/proposals/{proposal_id}/withdraw
│                                            # Publishes: ProposalWithdrawn event
│
├── pending/
│   └── page.tsx                            # Pending proposals
│                                            # BE: proposals-be
│                                            # GET /v1/proposals/my-proposals?status=pending
│
├── accepted/
│   └── page.tsx                            # Accepted proposals
│                                            # BE: proposals-be
│                                            # GET /v1/proposals/my-proposals?status=accepted
│
├── rejected/
│   └── page.tsx                            # Rejected proposals
│                                            # BE: proposals-be
│                                            # GET /v1/proposals/my-proposals?status=rejected
│
├── analytics/
│   └── page.tsx                            # Proposal analytics (freelancer)
│                                            # - Total proposals submitted
│                                            # - Acceptance rate
│                                            # - Average response time
│                                            # - Connects spent
│                                            # BE: proposals-be/analytics
│                                            # GET /v1/proposals/analytics
│
└── templates/
    ├── page.tsx                            # Proposal templates
    │                                        # - List saved templates
    │                                        # - Create new template
    │                                        # BE: proposals-be/templates
    │                                        # GET /v1/proposals/templates
    │
    ├── create/
    │   └── page.tsx                        # Create template
    │                                        # - Template name
    │                                        # - Cover letter template
    │                                        # - Default rate/terms
    │                                        # BE: proposals-be/templates
    │                                        # POST /v1/proposals/templates
    │
    └── [templateId]/
        └── edit/
            └── page.tsx                    # Edit template
                                            # BE: proposals-be/templates
                                            # PUT /v1/proposals/templates/{template_id}
                                            # DELETE /v1/proposals/templates/{template_id}
```

---

## Search & Discovery Module

```
apps/web/src/app/[locale]/(dashboard)/search/
│
├── jobs/
│   └── page.tsx                            # Advanced job search
│                                            # - Full-text search
│                                            # - Faceted filters
│                                            # - Autocomplete suggestions
│                                            # - Search history
│                                            # - Save search
│                                            # BE: search-be/query
│                                            # POST /v1/search/jobs
│                                            # Body: { query, filters: {...}, sort, page }
│                                            # BE: search-be/suggestions
│                                            # GET /v1/suggestions?q={query}
│
├── freelancers/
│   └── page.tsx                            # Advanced freelancer search (client)
│                                            # - Search by skills
│                                            # - Experience level
│                                            # - Hourly rate range
│                                            # - Location
│                                            # - Availability
│                                            # - Rating
│                                            # - Portfolio keywords
│                                            # BE: search-be/query
│                                            # POST /v1/search/freelancers
│                                            # Body: { query, filters: {...}, sort, page }
│
├── saved-searches/
│   ├── page.tsx                            # Saved searches list
│   │                                        # - List all saved searches
│   │                                        # - Email alerts toggle
│   │                                        # - Edit search
│   │                                        # - Delete search
│   │                                        # BE: search-be/saved_search
│   │                                        # GET /v1/saved-searches
│   │
│   ├── create/
│   │   └── page.tsx                        # Create saved search
│   │                                        # - Name the search
│   │                                        # - Set alert frequency
│   │                                        # - Save filters
│   │                                        # BE: search-be/saved_search
│   │                                        # POST /v1/saved-searches
│   │
│   └── [searchId]/
│       ├── page.tsx                        # Execute saved search
│       │                                    # BE: search-be/saved_search
│       │                                    # GET /v1/saved-searches/{search_id}/results
│       │
│       └── edit/
│           └── page.tsx                    # Edit saved search
│                                            # BE: search-be/saved_search
│                                            # PUT /v1/saved-searches/{search_id}
│                                            # DELETE /v1/saved-searches/{search_id}
│
└── portfolio/
    └── page.tsx                            # Search portfolios
                                            # - Search by project keywords
                                            # - Filter by skills used
                                            # - Filter by industry
                                            # BE: search-be/portfolio
                                            # POST /v1/search/portfolios
```

---

## Bidding System (Proposals Enhancement)

```
apps/web/src/app/[locale]/(dashboard)/jobs/[jobId]/bidding/
│
├── page.tsx                                # Active bids on job (client view)
│                                            # - Real-time bid updates
│                                            # - Current lowest bid
│                                            # - Bid history
│                                            # - Accept bid
│                                            # BE: proposals-be/bidding
│                                            # GET /v1/jobs/{job_id}/bids
│                                            # WebSocket: ws://proposals-be/v1/jobs/{job_id}/bids
│
└── place-bid/
    └── page.tsx                            # Place/update bid (freelancer)
                                            # - Current bid amount
                                            # - Minimum bid
                                            # - Place new bid
                                            # - Bid increment rules
                                            # - Outbid warning
                                            # BE: proposals-be/bidding
                                            # POST /v1/jobs/{job_id}/bids
                                            # PUT /v1/bids/{bid_id}
                                            # Publishes: BidPlaced, BidUpdated, OutbidAlert events

apps/web/src/app/[locale]/(dashboard)/proposals/[proposalId]/bidding/
│
└── page.tsx                                # Bid status for proposal
                                            # - Your current bid
                                            # - Current lowest bid
                                            # - Update bid
                                            # - Bid history
                                            # BE: proposals-be/bidding
                                            # GET /v1/proposals/{proposal_id}/bid
```

---

## Query Key Patterns (TanStack Query)

### Jobs
```typescript
['jobs', 'list', filters]                    // Job list with filters
['jobs', 'detail', jobId]                    // Job detail
['jobs', 'my-jobs', status]                  // User's posted jobs
['jobs', 'browse', filters]                  // Browse jobs (freelancer)
['jobs', 'saved']                            // Saved jobs
['jobs', 'invitations']                      // Job invitations
['jobs', 'drafts']                           // Job drafts
['jobs', 'categories']                       // Job categories
['jobs', 'analytics', jobId]                 // Job analytics
['jobs', 'recommendations']                  // Job recommendations
['jobs', 'similar', jobId]                   // Similar jobs
```

### Proposals
```typescript
['proposals', 'list', filters]               // Proposal list
['proposals', 'detail', proposalId]          // Proposal detail
['proposals', 'my-proposals', status]        // User's proposals
['proposals', 'job', jobId]                  // Proposals for a job
['proposals', 'count', jobId]                // Proposal count for job
['proposals', 'analytics']                   // Proposal analytics
['proposals', 'templates']                   // Proposal templates
```

### Bidding
```typescript
['bids', 'job', jobId]                       // Bids on a job
['bids', 'proposal', proposalId]             // Bid for proposal
['bids', 'realtime', jobId]                  // Real-time bid updates
```

### Search
```typescript
['search', 'jobs', query, filters]           // Job search results
['search', 'freelancers', query, filters]    // Freelancer search
['search', 'portfolios', query, filters]     // Portfolio search
['search', 'saved-searches']                 // Saved searches list
['search', 'suggestions', query]             // Autocomplete suggestions
['search', 'history']                        // Search history
```

### Profile
```typescript
['profile', 'me']                            // Current user profile
['profile', 'detail', userId]                // Other user profile
['profile', 'skills', userId]                // User skills
['profile', 'experience', userId]            // User experience
['profile', 'education', userId]             // User education
['profile', 'certifications', userId]        // User certifications
['profile', 'portfolio', userId]             // User portfolio
['profile', 'service-catalog', userId]       // Service catalog
['profile', 'availability', userId]          // Availability
['profile', 'verification', userId]          // Verification status
```

---

**End of Part 2**

**Continue to Part 3 for:**
- Contracts & Work Management
- Messaging & Notifications
- Financial Management (Wallet, Transactions, Invoices)
- Reviews & Ratings
