fe/
├── apps/
│   └── web/
│       └── src/
│           └── app/
│               └── [locale]/
│                   └── (dashboard)/
│                       ├── profile/
│                       │   ├── edit/
│                       │   │   └── page.tsx                            # Edit profile form
│                       │   │                                            # - Basic info (name, title, bio)
│                       │   │                                            # - Profile photo upload
│                       │   │                                            # - Location
│                       │   │                                            # - Languages
│                       │   │                                            # - Hourly rate (freelancer)
│                       │   │                                            # - Professional headline
│                       │   │                                            # BE: users-be/profile
│                       │   │                                            # PATCH /v1/users/{id}/profile
│                       │   │                                            # BE: storage-be/uploads
│                       │   │                                            # POST /v1/storage/upload (photo)
│                       │   ├── skills/
│                       │   │   ├── specializations/
│                       │   │   │   └── page.tsx                        # Specializations & niche expertise
│                       │   │   │                                        # - Add specializations
│                       │   │   │                                        # - Verification status
│                       │   │   │                                        # - Niche expertise tags
│                       │   │   │                                        # BE: users-be/capabilities
│                       │   │   │                                        # GET /v1/users/{id}/specializations
│                       │   │   │                                        # POST /v1/users/{id}/specializations
│                       │   │   └── page.tsx                            # Skills management
│                       │   │                                            # - List current skills with levels
│                       │   │                                            # - Add new skills (autocomplete)
│                       │   │                                            # - Edit skill proficiency
│                       │   │                                            # - Remove skills
│                       │   │                                            # - Primary skills (max 5)
│                       │   │                                            # BE: users-be/capabilities
│                       │   │                                            # GET /v1/users/{id}/skills
│                       │   │                                            # POST /v1/users/{id}/skills
│                       │   │                                            # PUT /v1/users/{id}/skills/{skill_id}
│                       │   │                                            # DELETE /v1/users/{id}/skills/{skill_id}
│                       │   │                                            # BE: search-be/taxonomy
│                       │   │                                            # GET /v1/taxonomy/skills (autocomplete)
│                       │   ├── experience/
│                       │   │   ├── add/
│                       │   │   │   └── page.tsx                        # Add experience form
│                       │   │   │                                        # - Company name
│                       │   │   │                                        # - Position title
│                       │   │   │                                        # - Start/end dates
│                       │   │   │                                        # - Current position checkbox
│                       │   │   │                                        # - Description (rich text)
│                       │   │   │                                        # - Skills used
│                       │   │   │                                        # BE: users-be/experience
│                       │   │   │                                        # POST /v1/users/{id}/experience
│                       │   │   ├── [experienceId]/
│                       │   │   │   └── edit/
│                       │   │   │       └── page.tsx                    # Edit experience form
│                       │   │   │                                        # BE: users-be/experience
│                       │   │   │                                        # PUT /v1/users/{id}/experience/{exp_id}
│                       │   │   │                                        # DELETE /v1/users/{id}/experience/{exp_id}
│                       │   │   └── page.tsx                            # Work experience list
│                       │   │                                            # - List all experience entries
│                       │   │                                            # - Add new experience button
│                       │   │                                            # - Edit/delete actions
│                       │   │                                            # BE: users-be/experience
│                       │   │                                            # GET /v1/users/{id}/experience
│                       │   ├── education/
│                       │   │   ├── add/
│                       │   │   │   └── page.tsx                        # Add education form
│                       │   │   │                                        # - School/university
│                       │   │   │                                        # - Degree/qualification
│                       │   │   │                                        # - Field of study
│                       │   │   │                                        # - Graduation year
│                       │   │   │                                        # - GPA (optional)
│                       │   │   │                                        # - Description
│                       │   │   │                                        # BE: users-be/education
│                       │   │   │                                        # POST /v1/users/{id}/education
│                       │   │   ├── [educationId]/
│                       │   │   │   └── edit/
│                       │   │   │       └── page.tsx                    # Edit education form
│                       │   │   │                                        # BE: users-be/education
│                       │   │   │                                        # PUT /v1/users/{id}/education/{edu_id}
│                       │   │   │                                        # DELETE /v1/users/{id}/education/{edu_id}
│                       │   │   └── page.tsx                            # Education list
│                       │   │                                            # BE: users-be/education
│                       │   │                                            # GET /v1/users/{id}/education
│                       │   ├── certifications/
│                       │   │   ├── add/
│                       │   │   │   └── page.tsx                        # Add certification
│                       │   │   │                                        # - Certification name
│                       │   │   │                                        # - Issuing organization
│                       │   │   │                                        # - Issue date
│                       │   │   │                                        # - Expiry date (if any)
│                       │   │   │                                        # - Credential ID
│                       │   │   │                                        # - Credential URL
│                       │   │   │                                        # - Certificate upload
│                       │   │   │                                        # BE: users-be/credentials
│                       │   │   │                                        # POST /v1/users/{id}/certifications
│                       │   │   │                                        # BE: storage-be/uploads
│                       │   │   ├── verify/
│                       │   │   │   └── [certificationId]/
│                       │   │   │       └── page.tsx                    # Verification request
│                       │   │   │                                        # - Submit for verification
│                       │   │   │                                        # - Upload proof
│                       │   │   │                                        # BE: users-be/credentials
│                       │   │   │                                        # POST /v1/users/{id}/certifications/{cert_id}/verify
│                       │   │   └── page.tsx                            # Certifications list
│                       │   │                                            # - External certifications
│                       │   │                                            # - Platform certifications
│                       │   │                                            # - Badges earned
│                       │   │                                            # BE: users-be/credentials
│                       │   │                                            # GET /v1/users/{id}/credentials
│                       │   ├── portfolio/
│                       │   │   ├── add/
│                       │   │   │   └── page.tsx                        # Add portfolio item
│                       │   │   │                                        # - Project title
│                       │   │   │                                        # - Description
│                       │   │   │                                        # - Media upload (images/videos)
│                       │   │   │                                        # - Project URL
│                       │   │   │                                        # - Skills used
│                       │   │   │                                        # - Client (optional)
│                       │   │   │                                        # - Completion date
│                       │   │   │                                        # BE: users-be/portfolio
│                       │   │   │                                        # POST /v1/users/{id}/portfolio
│                       │   │   │                                        # BE: storage-be/uploads
│                       │   │   │                                        # POST /v1/storage/upload (media)
│                       │   │   ├── [portfolioId]/
│                       │   │   │   ├── edit/
│                       │   │   │   │   └── page.tsx                    # Edit portfolio item
│                       │   │   │   │                                    # BE: users-be/portfolio
│                       │   │   │   │                                    # PUT /v1/users/{id}/portfolio/{item_id}
│                       │   │   │   │                                    # DELETE /v1/users/{id}/portfolio/{item_id}
│                       │   │   │   └── page.tsx                        # Portfolio item detail
│                       │   │   │                                        # BE: users-be/portfolio
│                       │   │   │                                        # GET /v1/users/{id}/portfolio/{item_id}
│                       │   │   ├── reorder/
│                       │   │   │   └── page.tsx                        # Reorder portfolio items
│                       │   │   │                                        # - Drag & drop interface
│                       │   │   │                                        # - Set featured items
│                       │   │   │                                        # BE: users-be/portfolio
│                       │   │   │                                        # PATCH /v1/users/{id}/portfolio/reorder
│                       │   │   │                                        # Body: { item_ids: string[] }
│                       │   │   └── page.tsx                            # Portfolio items list
│                       │   │                                            # - Grid/list view
│                       │   │                                            # - Featured items
│                       │   │                                            # - Reorder items (drag & drop)
│                       │   │                                            # BE: users-be/portfolio
│                       │   │                                            # GET /v1/users/{id}/portfolio
│                       │   ├── service-catalog/
│                       │   │   ├── add/
│                       │   │   │   └── page.tsx                        # Add service
│                       │   │   │                                        # - Service name
│                       │   │   │                                        # - Description
│                       │   │   │                                        # - Capabilities required
│                       │   │   │                                        # - Delivery time
│                       │   │   │                                        # - Pricing
│                       │   │   │                                        # - Packages (Basic/Standard/Premium)
│                       │   │   │                                        # BE: users-be/service_catalog
│                       │   │   │                                        # POST /v1/users/{id}/services
│                       │   │   ├── [serviceId]/
│                       │   │   │   └── edit/
│                       │   │   │       └── page.tsx                    # Edit service
│                       │   │   │                                        # BE: users-be/service_catalog
│                       │   │   │                                        # PUT /v1/users/{id}/services/{service_id}
│                       │   │   │                                        # DELETE /v1/users/{id}/services/{service_id}
│                       │   │   └── page.tsx                            # Service catalog management (freelancer)
│                       │   │                                            # - List offered services
│                       │   │                                            # - Service packages
│                       │   │                                            # - Pricing tiers
│                       │   │                                            # BE: users-be/service_catalog
│                       │   │                                            # GET /v1/users/{id}/service-catalog
│                       │   ├── availability/
│                       │   │   └── page.tsx                            # Availability management
│                       │   │                                            # - Calendar view
│                       │   │                                            # - Set available hours
│                       │   │                                            # - Time zone
│                       │   │                                            # - Vacation mode
│                       │   │                                            # - Max concurrent projects
│                       │   │                                            # BE: users-be/availability
│                       │   │                                            # GET /v1/users/{id}/availability
│                       │   │                                            # POST /v1/users/{id}/availability
│                       │   │                                            # PATCH /v1/users/{id}/availability
│                       │   ├── verification/
│                       │   │   ├── phone/
│                       │   │   │   └── page.tsx                        # Phone verification
│                       │   │   │                                        # - Enter phone number
│                       │   │   │                                        # - Receive OTP
│                       │   │   │                                        # - Verify OTP
│                       │   │   │                                        # BE: users-be/verification
│                       │   │   │                                        # POST /v1/users/{id}/verify-phone/send
│                       │   │   │                                        # POST /v1/users/{id}/verify-phone/verify
│                       │   │   ├── identity/
│                       │   │   │   └── page.tsx                        # ID verification
│                       │   │   │                                        # - Upload ID document
│                       │   │   │                                        # - Selfie verification
│                       │   │   │                                        # - Address proof
│                       │   │   │                                        # BE: users-be/verification/kyc
│                       │   │   │                                        # POST /v1/users/{id}/kyc/submit
│                       │   │   │                                        # BE: storage-be/uploads
│                       │   │   │                                        # BE: admin-be/kyc_case (creates case)
│                       │   │   └── page.tsx                            # Verification status
│                       │   │                                            # - Email verified
│                       │   │                                            # - Phone verified
│                       │   │                                            # - ID verification status
│                       │   │                                            # - Payment method verified
│                       │   │                                            # BE: users-be/verification
│                       │   │                                            # GET /v1/users/{id}/verification-status
│                       │   └── page.tsx                                # Profile overview / public view
│                       │                                                # - Profile header (photo, name, title, location)
│                       │                                                # - Stats (rating, jobs completed, earnings)
│                       │                                                # - Skills showcase
│                       │                                                # - Portfolio highlights
│                       │                                                # - Recent reviews
│                       │                                                # - Availability calendar
│                       │                                                # BE: users-be/profile
│                       │                                                # GET /v1/users/{id}/profile
│                       │                                                # BE: users-be/capabilities
│                       │                                                # GET /v1/users/{id}/skills
│                       │                                                # BE: users-be/portfolio
│                       │                                                # GET /v1/users/{id}/portfolio
│                       │                                                # BE: reviews-be
│                       │                                                # GET /v1/reviews?user_id={id}&limit=5
│                       │                                                # BE: users-be/availability
│                       │                                                # GET /v1/users/{id}/availability
│                       ├── jobs/
│                       │   ├── browse/
│                       │   │   └── page.tsx                            # Job listings with filters
│                       │   │                                            # - Category filters
│                       │   │                                            # - Budget range
│                       │   │                                            # - Experience level
│                       │   │                                            # - Job type (fixed/hourly)
│                       │   │                                            # - Location preferences
│                       │   │                                            # - Skills required
│                       │   │                                            # - Posted date
│                       │   │                                            # - Saved jobs indicator
│                       │   │                                            # - "Best Matches" tab
│                       │   │                                            # BE: jobs-be/job
│                       │   │                                            # GET /v1/jobs/browse?filters={...}
│                       │   │                                            # BE: search-be/query
│                       │   │                                            # POST /v1/search/jobs (for advanced search)
│                       │   ├── saved/
│                       │   │   └── page.tsx                            # Saved/bookmarked jobs
│                       │   │                                            # BE: jobs-be/saved_jobs
│                       │   │                                            # GET /v1/jobs/saved
│                       │   │                                            # DELETE /v1/jobs/saved/{job_id}
│                       │   ├── invitations/
│                       │   │   └── page.tsx                            # Job invitations received
│                       │   │                                            # BE: jobs-be/invitations
│                       │   │                                            # GET /v1/jobs/invitations
│                       │   ├── post/
│                       │   │   └── page.tsx                            # Post a new job (client)
│                       │   │                                            # - Job title
│                       │   │                                            # - Job description (rich text editor)
│                       │   │                                            # - Category selection
│                       │   │                                            # - Required skills (autocomplete)
│                       │   │                                            # - Experience level
│                       │   │                                            # - Job type (fixed/hourly)
│                       │   │                                            # - Budget/rate
│                       │   │                                            # - Duration
│                       │   │                                            # - Attachments
│                       │   │                                            # - Screening questions
│                       │   │                                            # - Visibility (public/private/invited)
│                       │   │                                            # - Save as draft
│                       │   │                                            # BE: jobs-be/job
│                       │   │                                            # POST /v1/jobs
│                       │   │                                            # Body: { title, description, category_id, skills, budget, ... }
│                       │   │                                            # BE: jobs-be/attachments
│                       │   │                                            # POST /v1/jobs/{job_id}/attachments
│                       │   │                                            # BE: jobs-be/screening
│                       │   │                                            # POST /v1/jobs/{job_id}/screening-questions
│                       │   │                                            # BE: storage-be/uploads
│                       │   │                                            # Publishes: JobPosted event
│                       │   ├── drafts/
│                       │   │   ├── [draftId]/
│                       │   │   │   └── edit/
│                       │   │   │       └── page.tsx                    # Edit draft
│                       │   │   │                                        # BE: jobs-be/draft
│                       │   │   │                                        # PUT /v1/jobs/drafts/{draft_id}
│                       │   │   │                                        # DELETE /v1/jobs/drafts/{draft_id}
│                       │   │   └── page.tsx                            # Job drafts list
│                       │   │                                            # BE: jobs-be/draft
│                       │   │                                            # GET /v1/jobs/drafts
│                       │   ├── [jobId]/
│                       │   │   ├── bidding/
│                       │   │   │   ├── place-bid/
│                       │   │   │   │   └── page.tsx                    # Place/update bid (freelancer)
│                       │   │   │   │                                    # - Current bid amount
│                       │   │   │   │                                    # - Minimum bid
│                       │   │   │   │                                    # - Place new bid
│                       │   │   │   │                                    # - Bid increment rules
│                       │   │   │   │                                    # - Outbid warning
│                       │   │   │   │                                    # BE: proposals-be/bidding
│                       │   │   │   │                                    # POST /v1/jobs/{job_id}/bids
│                       │   │   │   │                                    # PUT /v1/bids/{bid_id}
│                       │   │   │   │                                    # Publishes: BidPlaced, BidUpdated, OutbidAlert events
│                       │   │   │   └── page.tsx                        # Active bids on job (client view)
│                       │   │   │                                        # - Real-time bid updates
│                       │   │   │                                        # - Current lowest bid
│                       │   │   │                                        # - Bid history
│                       │   │   │                                        # - Accept bid
│                       │   │   │                                        # BE: proposals-be/bidding
│                       │   │   │                                        # GET /v1/jobs/{job_id}/bids
│                       │   │   │                                        # WebSocket: ws://proposals-be/v1/jobs/{job_id}/bids
│                       │   │   ├── proposals/
│                       │   │   │   ├── [proposalId]/
│                       │   │   │   │   └── page.tsx                    # Proposal detail
│                       │   │   │   │                                    # - Full proposal view
│                       │   │   │   │                                    # - Freelancer profile preview
│                       │   │   │   │                                    # - Accept/Reject buttons
│                       │   │   │   │                                    # - Shortlist button
│                       │   │   │   │                                    # - Message freelancer
│                       │   │   │   │                                    # BE: proposals-be
│                       │   │   │   │                                    # GET /v1/proposals/{proposal_id}
│                       │   │   │   │                                    # POST /v1/proposals/{proposal_id}/accept
│                       │   │   │   │                                    # POST /v1/proposals/{proposal_id}/reject
│                       │   │   │   │                                    # POST /v1/proposals/{proposal_id}/shortlist
│                       │   │   │   └── page.tsx                        # Proposals received (client)
│                       │   │   │                                        # - List all proposals
│                       │   │   │                                        # - Filter (all/shortlisted/archived)
│                       │   │   │                                        # - Sort (date, rate, rating)
│                       │   │   │                                        # - Proposal cards with key info
│                       │   │   │                                        # BE: proposals-be
│                       │   │   │                                        # GET /v1/proposals?job_id={job_id}
│                       │   │   ├── invite/
│                       │   │   │   └── page.tsx                        # Invite freelancers (client)
│                       │   │   │                                        # - Search freelancers
│                       │   │   │                                        # - Send invitation with message
│                       │   │   │                                        # BE: jobs-be/invitations
│                       │   │   │                                        # POST /v1/jobs/{job_id}/invitations
│                       │   │   │                                        # BE: search-be
│                       │   │   │                                        # POST /v1/search/freelancers
│                       │   │   │                                        # BE: communications-be
│                       │   │   │                                        # Publishes: JobInvitationSent event
│                       │   │   ├── analytics/
│                       │   │   │   └── page.tsx                        # Job analytics (client)
│                       │   │   │                                        # - Views
│                       │   │   │                                        # - Proposals received
│                       │   │   │                                        # - Proposal conversion rate
│                       │   │   │                                        # - Time to hire
│                       │   │   │                                        # BE: jobs-be/analytics
│                       │   │   │                                        # GET /v1/jobs/{job_id}/analytics
│                       │   │   ├── edit/
│                       │   │   │   └── page.tsx                        # Edit job (client only)
│                       │   │   │                                        # - Same form as post job
│                       │   │   │                                        # - Cannot edit if has accepted proposals
│                       │   │   │                                        # BE: jobs-be/job
│                       │   │   │                                        # PUT /v1/jobs/{job_id}
│                       │   │   │                                        # Publishes: JobUpdated event
│                       │   │   ├── close/
│                       │   │   │   └── page.tsx                        # Close job
│                       │   │   │                                        # - Reason for closing
│                       │   │   │                                        # - Notify applicants
│                       │   │   │                                        # BE: jobs-be/job
│                       │   │   │                                        # POST /v1/jobs/{job_id}/close
│                       │   │   │                                        # Publishes: JobClosed event
│                       │   │   └── page.tsx                            # Job detail page
│                       │   │                                            # - Full job description
│                       │   │                                            # - Client info
│                       │   │                                            # - Skills required
│                       │   │                                            # - Budget/rate
│                       │   │                                            # - Proposals count
│                       │   │                                            # - Similar jobs
│                       │   │                                            # - "Submit Proposal" button (freelancer)
│                       │   │                                            # - Save job button
│                       │   │                                            # BE: jobs-be/job
│                       │   │                                            # GET /v1/jobs/{job_id}
│                       │   │                                            # BE: proposals-be
│                       │   │                                            # GET /v1/proposals/count?job_id={job_id}
│                       │                       │                        # BE: search-be/similarity
│                       │   │                                            # GET /v1/similarity/jobs/{job_id}
│                       │   ├── my-jobs/
│                       │   │   ├── active/
│                       │   │   │   └── page.tsx                        # Active jobs only
│                       │   │   │                                        # BE: jobs-be/job
│                       │   │   │                                        # GET /v1/jobs/my-jobs?status=active
│                       │   │   ├── closed/
│                       │   │   │   └── page.tsx                        # Closed jobs
│                       │   │   │                                        # BE: jobs-be/job
│                       │   │   │                                        # GET /v1/jobs/my-jobs?status=closed
│                       │   │   └── page.tsx                            # All posted jobs
│                       │   │                                            # - Active jobs
│                       │   │                                            # - Closed jobs
│                       │   │                                            # - Drafts
│                       │   │                                            # BE: jobs-be/job
│                       │   │                                            # GET /v1/jobs/my-jobs?status=active
│                       │   ├── categories/
│                       │   │   ├── [categoryId]/
│                       │   │   │   └── page.tsx                        # Jobs in category
│                       │   │   │                                        # BE: jobs-be/job
│                       │   │   │                                        # GET /v1/jobs?category_id={category_id}
│                       │   │   └── page.tsx                            # Browse by category
│                       │   │                                            # - Category grid
│                       │   │                                            # - Subcategories
│                       │   │                                            # BE: jobs-be/categories
│                       │   │                                            # GET /v1/jobs/categories
│                       │   ├── recommendations/
│                       │   │   └── page.tsx                            # Recommended jobs (freelancer)
│                       │   │                                            # - ML-powered job recommendations
│                       │   │                                            # - Based on skills, history, preferences
│                       │   │                                            # BE: search-be/recommendations
│                       │   │                                            # GET /v1/recommendations/jobs
│                       │   └── page.tsx                                # Jobs list (role-based)
│                       │                                                # Freelancer view: Browse available jobs
│                       │                                                # Client view: My posted jobs
│                       │                                                # - Filters (category, budget, skills, etc.)
│                       │                                                # - Search bar
│                       │                                                # - Sort options
│                       │                                                # - Pagination
│                       │                                                # BE: jobs-be/job
│                       │                                                # GET /v1/jobs?filters=...&page=1&limit=20
│                       │                                                # Freelancer: GET /v1/jobs/browse
│                       │                                                # Client: GET /v1/jobs/my-jobs
│                       ├── proposals/
│                       │   ├── submit/
│                       │   │   ├── [jobId]/
│                       │   │   │   └── page.tsx                        # Submit proposal (freelancer)
│                       │   │   │                                        # - Cover letter (required)
│                       │   │   │                                        # - Proposed rate/budget
│                       │   │   │                                        # - Proposed timeline
│                       │   │   │                                        # - Answer screening questions
│                       │   │   │                                        # - Attachments (portfolio samples)
│                       │   │   │                                        # - Milestones (for fixed-price)
│                       │   │   │                                        # - Terms acceptance
│                       │   │   │                                        # - Connects deduction warning
│                       │   │   │                                        # BE: proposals-be
│                       │   │   │                                        # POST /v1/proposals
│                       │   │   │                                        # Body: { job_id, cover_letter, rate, timeline, ... }
│                       │   │   │                                        # BE: subscriptions-be/connects
│                       │   │   │                                        # POST /v1/connects/deduct
│                       │   │   │                                        # BE: storage-be/uploads
│                       │   │   │                                        # Publishes: ProposalSubmitted event
│                       │   ├── [proposalId]/
│                       │   │   ├── bidding/
│                       │   │   │   └── page.tsx                        # Bid status for proposal
│                       │   │   │                                        # - Your current bid
│                       │   │   │                                        # - Current lowest bid
│                       │   │   │                                        # - Update bid
│                       │   │   │                                        # - Bid history
│                       │   │   │                                        # BE: proposals-be/bidding
│                       │   │   │                                        # GET /v1/proposals/{proposal_id}/bid
│                       │   │   ├── edit/
│                       │   │   │   └── page.tsx                        # Edit proposal
│                       │   │   │                                        # - Only if status = DRAFT or PENDING
│                       │   │   │                                        # - Update cover letter, rate, timeline
│                       │   │   │                                        # BE: proposals-be
│                       │   │   │                                        # PUT /v1/proposals/{proposal_id}
│                       │   │   ├── withdraw/
│                       │   │   │   └── page.tsx                        # Withdraw proposal
│                       │   │   │                                        # - Confirmation dialog
│                       │   │   │                                        # - Reason for withdrawal
│                       │   │   │                                        # - Connects refund info
│                       │   │   │                                        # BE: proposals-be
│                       │   │   │                                        # POST /v1/proposals/{proposal_id}/withdraw
│                       │   │   │                                        # Publishes: ProposalWithdrawn event
│                       │   │   └── page.tsx                            # Proposal detail
│                       │   │                                            # - View submitted proposal
│                       │   │                                            # - Proposal status
│                       │   │                                            # - Client messages/feedback
│                       │   │                                            # - Withdraw option (if pending)
│                       │   │                                            # BE: proposals-be
│                       │   │                                            # GET /v1/proposals/{proposal_id}
│                       │   ├── pending/
│                       │   │   └── page.tsx                            # Pending proposals
│                       │   │                                            # BE: proposals-be
│                       │   │                                            # GET /v1/proposals/my-proposals?status=pending
│                       │   ├── accepted/
│                       │   │   └── page.tsx                            # Accepted proposals
│                       │   │                                            # BE: proposals-be
│                       │   │                                            # GET /v1/proposals/my-proposals?status=accepted
│                       │   ├── rejected/
│                       │   │   └── page.tsx                            # Rejected proposals
│                       │   │                                            # BE: proposals-be
│                       │   │                                            # GET /v1/proposals/my-proposals?status=rejected
│                       │   ├── analytics/
│                       │   │   └── page.tsx                            # Proposal analytics (freelancer)
│                       │   │                                            # - Total proposals submitted
│                       │   │                                            # - Acceptance rate
│                       │   │                                            # - Average response time
│                       │   │                                            # - Connects spent
│                       │   │                                            # BE: proposals-be/analytics
│                       │   │                                            # GET /v1/proposals/analytics
│                       │   ├── templates/
│                       │   │   ├── create/
│                       │   │   │   └── page.tsx                        # Create template
│                       │   │   │                                        # - Template name
│                       │   │   │                                        # - Cover letter template
│                       │   │   │                                        # - Default rate/terms
│                       │   │   │                                        # BE: proposals-be/templates
│                       │   │   │                                        # POST /v1/proposals/templates
│                       │   │   ├── [templateId]/
│                       │   │   │   └── edit/
│                       │   │   │       └── page.tsx                    # Edit template
│                       │   │   │                                        # BE: proposals-be/templates
│                       │   │   │                                        # PUT /v1/proposals/templates/{template_id}
│                       │   │   │                                        # DELETE /v1/proposals/templates/{template_id}
│                       │   │   └── page.tsx                            # Proposal templates
│                       │   │                                            # - List saved templates
│                       │   │                                            # - Create new template
│                       │   │                                            # BE: proposals-be/templates
│                       │   │                                            # GET /v1/proposals/templates
│                       │   └── page.tsx                                # Proposals list
│                       │                                                # Freelancer: Submitted proposals
│                       │                                                # Client: Received proposals (redirect to jobs)
│                       │                                                # BE: proposals-be
│                       │                                                # GET /v1/proposals/my-proposals
│                       └── search/
│                           ├── saved-searches/
│                           │   ├── create/
│                           │   │   └── page.tsx                        # Create saved search
│                           │   │                                        # - Name the search
│                           │   │                                        # - Set alert frequency
│                           │   │                                        # - Save filters
│                           │   │                                        # BE: search-be/saved_search
│                           │   │                                        # POST /v1/saved-searches
│                           │   ├── [searchId]/
│                           │   │   ├── edit/
│                           │   │   │   └── page.tsx                    # Edit saved search
│                           │   │   │                                    # BE: search-be/saved_search
│                           │   │   │                                    # PUT /v1/saved-searches/{search_id}
│                           │   │   │                                    # DELETE /v1/saved-searches/{search_id}
│                           │   │   └── page.tsx                        # Execute saved search
│                           │   │                                        # BE: search-be/saved_search
│                           │   │                                        # GET /v1/saved-searches/{search_id}/results
│                           │   └── page.tsx                            # Saved searches list
│                           │                                            # - List all saved searches
│                           │                                            # - Email alerts toggle
│                           │                                            # - Edit search
│                           │                                            # - Delete search
│                           │                                            # BE: search-be/saved_search
│                           │                                            # GET /v1/saved-searches
│                           ├── portfolio/
│                           │   └── page.tsx                            # Search portfolios
│                           │                                            # - Search by project keywords
│                           │                                            # - Filter by skills used
│                           │                                            # - Filter by industry
│                           │                                            # BE: search-be/portfolio
│                           │                                            # POST /v1/search/portfolios
│                           ├── freelancers/
│                           │   └── page.tsx                            # Advanced freelancer search (client)
│                           │                                            # - Search by skills
│                           │                                            # - Experience level
│                           │                                            # - Hourly rate range
│                           │                                            # - Location
│                           │                                            # - Availability
│                           │                                            # - Rating
│                           │                                            # - Portfolio keywords
│                           │                                            # BE: search-be/query
│                           │                                            # POST /v1/search/freelancers
│                           │                                            # Body: { query, filters: {...}, sort, page }
│                           └── jobs/
│                               └── page.tsx                            # Advanced job search
│                                                                        # - Full-text search
│                                                                        # - Faceted filters
│                                                                        # - Autocomplete suggestions
│                                                                        # - Search history
│                                                                        # - Save search
│                                                                        # BE: search-be/query
│                                                                        # POST /v1/search/jobs
│                                                                        # Body: { query, filters: {...}, sort, page }
│                                                                        # BE: search-be/suggestions
│                                                                        # GET /v1/suggestions?q={query}
