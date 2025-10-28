fe/
└── fe/
    ├── apps/
    │   └── mobile/
    │       └── app/
    │           └── (tabs)/
    │               ├── contests/  # ❌ MISSING ENTIRE FEATURE
    │               │   ├── [contestId]/
    │               │   │   ├── details.tsx  # Contest details (mobile)
    │               │   │   │   # - Contest title, description, rules
    │               │   │   │   # - Prize details and distribution
    │               │   │   │   # - Entry requirements
    │               │   │   │   # - Submission deadline
    │               │   │   │   # - Current entries count
    │               │   │   │   # - Leaderboard preview
    │               │   │   │   # BE: contests-be/contest
    │               │   │   │   # GET /v1/contests/{contest_id}
    │               │   │   ├── entries.tsx  # Contest entries list (mobile)
    │               │   │   │   # - All contest entries
    │               │   │   │   # - Entry previews
    │               │   │   │   # - Voting/rating interface
    │               │   │   │   # - Filter by criteria
    │               │   │   │   # - Sort by votes/date
    │               │   │   │   # BE: contests-be/entry
    │               │   │   │   # GET /v1/contests/{contest_id}/entries
    │               │   │   ├── leaderboard.tsx  # Contest leaderboard (mobile)
    │               │   │   │   # - Top entries ranked
    │               │   │   │   # - Current standings
    │               │   │   │   # - Score/vote counts
    │               │   │   │   # - Winner indicators
    │               │   │   │   # - Real-time updates
    │               │   │   │   # BE: contests-be/leaderboard
    │               │   │   │   # GET /v1/contests/{contest_id}/leaderboard
    │               │   │   └── submit-entry.tsx  # Submit contest entry (mobile)
    │               │   │       # - Entry submission form
    │               │   │       # - File upload (portfolio/work samples)
    │               │   │       # - Entry description
    │               │   │       # - Terms acceptance
    │               │   │       # - Submit button
    │               │   │       # BE: contests-be/entry
    │               │   │       # POST /v1/contests/{contest_id}/entries
    │               │   ├── active/
    │               │   │   └── index.tsx  # Active contests list (mobile)
    │               │   │       # - All active contests
    │               │   │       # - Contest cards with key info
    │               │   │       # - Filter by category
    │               │   │       # - Sort by deadline/prize
    │               │   │       # - Quick enter button
    │               │   │       # BE: contests-be/contest
    │               │   │       # GET /v1/contests?status=active
    │               │   ├── browse/
    │               │   │   └── index.tsx  # Browse all contests (mobile)
    │               │   │       # - All contests (active/upcoming/past)
    │               │   │       # - Search and filters
    │               │   │       # - Category navigation
    │               │   │       # - Contest previews
    │               │   │       # BE: contests-be/contest
    │               │   │       # GET /v1/contests
    │               │   ├── create/
    │               │   │   └── index.tsx  # Create new contest (mobile)
    │               │   │       # - Contest creation form
    │               │   │       # - Title, description, rules
    │               │   │       # - Prize setup
    │               │   │       # - Submission requirements
    │               │   │       # - Duration and deadlines
    │               │   │       # - Publish button
    │               │   │       # BE: contests-be/contest
    │               │   │       # POST /v1/contests
    │               │   └── my-entries/
    │               │       └── index.tsx  # My contest entries (mobile)
    │               │           # - All my submitted entries
    │               │           # - Entry status (pending/accepted/rejected)
    │               │           # - Voting/ranking position
    │               │           # - Edit/withdraw options
    │               │           # BE: contests-be/entry
    │               │           # GET /v1/contests/my-entries
    │               │
    │               ├── escrow/  # ❌ MISSING ENTIRE FEATURE
    │               │   ├── [escrowId]/
    │               │   │   ├── dispute.tsx  # File escrow dispute (mobile)
    │               │   │   │   # - Dispute reason form
    │               │   │   │   # - Evidence upload
    │               │   │   │   # - Issue description
    │               │   │   │   # - Submit dispute button
    │               │   │   │   # BE: escrow-be/dispute
    │               │   │   │   # POST /v1/escrow/{escrow_id}/dispute
    │               │   │   ├── release.tsx  # Release escrow funds (mobile)
    │               │   │   │   # - Release confirmation
    │               │   │   │   # - Amount breakdown
    │               │   │   │   # - Release conditions review
    │               │   │   │   # - Signature/authorization
    │               │   │   │   # - Confirm release button
    │               │   │   │   # BE: escrow-be/release
    │               │   │   │   # POST /v1/escrow/{escrow_id}/release
    │               │   │   └── details.tsx  # Escrow details (mobile)
    │               │   │       # - Escrow information
    │               │   │       # - Amount and parties
    │               │   │       # - Status and conditions
    │               │   │       # - Timeline
    │               │   │       # - Action buttons (release/dispute)
    │               │   │       # BE: escrow-be/escrow
    │               │   │       # GET /v1/escrow/{escrow_id}
    │               │   ├── active/
    │               │   │   └── index.tsx  # Active escrow list (mobile)
    │               │   │       # - All active escrows
    │               │   │       # - Filter by role (buyer/seller)
    │               │   │       # - Status indicators
    │               │   │       # - Quick actions
    │               │   │       # BE: escrow-be/escrow
    │               │   │       # GET /v1/escrow?status=active
    │               │   └── history/
    │               │       └── index.tsx  # Escrow history (mobile)
    │               │           # - Completed escrows
    │               │           # - Transaction details
    │               │           # - Filter by date/status
    │               │           # - Export records
    │               │           # BE: escrow-be/escrow
    │               │           # GET /v1/escrow?status=completed,disputed,cancelled
    │               │
    │               ├── milestones/  # ❌ MISSING ENTIRE FEATURE (as separate tab)
    │               │   ├── [milestoneId]/
    │               │   │   ├── details.tsx  # Milestone details (mobile)
    │               │   │   │   # - Milestone information
    │               │   │   │   # - Description and requirements
    │               │   │   │   # - Amount and deadline
    │               │   │   │   # - Status and progress
    │               │   │   │   # - Deliverables list
    │               │   │   │   # - Action buttons
    │               │   │   │   # BE: milestones-be/milestone
    │               │   │   │   # GET /v1/milestones/{milestone_id}
    │               │   │   ├── submit.tsx  # Submit milestone deliverable (mobile)
    │               │   │   │   # - Deliverable upload
    │               │   │   │   # - Description
    │               │   │   │   # - Evidence/proof of work
    │               │   │   │   # - Submit for review button
    │               │   │   │   # BE: milestones-be/deliverable
    │               │   │   │   # POST /v1/milestones/{milestone_id}/submit
    │               │   │   └── approve.tsx  # Approve milestone (mobile)
    │               │   │       # - Review deliverable
    │               │   │       # - Approval checklist
    │               │   │       # - Request changes option
    │               │   │       # - Approve/reject buttons
    │               │   │       # - Payment release confirmation
    │               │   │       # BE: milestones-be/approval
    │               │   │       # POST /v1/milestones/{milestone_id}/approve
    │               │   ├── pending/
    │               │   │   └── index.tsx  # Pending milestones (mobile)
    │               │   │       # - All pending milestones
    │               │   │       # - Filter by contract
    │               │   │       # - Sort by deadline
    │               │   │       # - Quick submit action
    │               │   │       # BE: milestones-be/milestone
    │               │   │       # GET /v1/milestones?status=pending
    │               │   ├── in-review/
    │               │   │   └── index.tsx  # In-review milestones (mobile)
    │               │   │       # - Milestones awaiting approval
    │               │   │       # - Submitted deliverables
    │               │   │       # - Review status
    │               │   │       # - Action required indicators
    │               │   │       # BE: milestones-be/milestone
    │               │   │       # GET /v1/milestones?status=in_review
    │               │   └── completed/
    │               │       └── index.tsx  # Completed milestones (mobile)
    │               │           # - All completed milestones
    │               │           # - Payment history
    │               │           # - Filter by date/contract
    │               │           # - Export records
    │               │           # BE: milestones-be/milestone
    │               │           # GET /v1/milestones?status=completed
    │               │
    │               ├── talent-cloud/  # ❌ MISSING ENTIRE FEATURE
    │               │   ├── agencies/
    │               │   │   ├── [agencyId]/
    │               │   │   │   └── details.tsx  # Agency details (mobile)
    │               │   │   │       # - Agency profile
    │               │   │   │       # - Team members
    │               │   │   │       # - Portfolio and projects
    │               │   │   │       # - Reviews and ratings
    │               │   │   │       # - Contact/hire button
    │               │   │   │       # BE: talent-cloud-be/agency
    │               │   │   │       # GET /v1/agencies/{agency_id}
    │               │   │   └── index.tsx  # Agencies list (mobile)
    │               │   │       # - Browse all agencies
    │               │   │       # - Search and filters
    │               │   │       # - Sort by rating/size
    │               │   │       # - Agency cards
    │               │   │       # BE: talent-cloud-be/agency
    │               │   │       # GET /v1/agencies
    │               │   ├── projects/
    │               │   │   ├── [projectId]/
    │               │   │   │   └── details.tsx  # Project details (mobile)
    │               │   │   │       # - Project information
    │               │   │   │       # - Team assigned
    │               │   │   │       # - Timeline and milestones
    │               │   │   │       # - Status and progress
    │               │   │   │       # - Documents/deliverables
    │               │   │   │       # BE: talent-cloud-be/project
    │               │   │   │       # GET /v1/projects/{project_id}
    │               │   │   └── index.tsx  # Projects list (mobile)
    │               │   │       # - All projects
    │               │   │       # - Filter by status/team
    │               │   │       # - Project cards
    │               │   │       # - Create new button
    │               │   │       # BE: talent-cloud-be/project
    │               │   │       # GET /v1/projects
    │               │   └── teams/
    │               │       ├── [teamId]/
    │               │       │   ├── details.tsx  # Team details (mobile)
    │               │       │   │   # - Team information
    │               │       │   │   # - Skills and expertise
    │               │       │   │   # - Availability
    │               │       │   │   # - Projects history
    │               │       │   │   # - Hire team button
    │               │       │   │   # BE: talent-cloud-be/team
    │               │       │   │   # GET /v1/teams/{team_id}
    │               │       │   └── members.tsx  # Team members (mobile)
    │               │       │       # - All team members
    │               │       │       # - Member profiles
    │               │       │       # - Roles and skills
    │               │       │       # - Contact information
    │               │       │       # BE: talent-cloud-be/team-member
    │               │       │       # GET /v1/teams/{team_id}/members
    │               │       └── index.tsx  # Teams list (mobile)
    │               │           # - Browse all teams
    │               │           # - Search and filters
    │               │           # - Team cards
    │               │           # - Create team button
    │               │           # BE: talent-cloud-be/team
    │               │           # GET /v1/teams
    │               │
    │               ├── timesheet/  # ❌ MISSING ENTIRE FEATURE (as separate tab)
    │               │   ├── [timesheetId]/
    │               │   │   ├── details.tsx  # Timesheet details (mobile)
    │               │   │   │   # - Timesheet information
    │               │   │   │   # - Time entries
    │               │   │   │   # - Total hours
    │               │   │   │   # - Status (draft/submitted/approved)
    │               │   │   │   # - Edit/submit buttons
    │               │   │   │   # BE: timesheet-be/timesheet
    │               │   │   │   # GET /v1/timesheets/{timesheet_id}
    │               │   │   └── edit.tsx  # Edit timesheet (mobile)
    │               │   │       # - Edit time entries
    │               │   │       # - Add/remove entries
    │               │   │       # - Update hours
    │               │   │       # - Notes/descriptions
    │               │   │       # - Save/submit buttons
    │               │   │       # BE: timesheet-be/timesheet
    │               │   │       # PUT /v1/timesheets/{timesheet_id}
    │               │   ├── manual-time/
    │               │   │   └── index.tsx  # Manual time entry (mobile)
    │               │   │       # - Add time entry form
    │               │   │       # - Date/time selection
    │               │   │       # - Duration input
    │               │   │       # - Task/project selection
    │               │   │       # - Description
    │               │   │       # - Save button
    │               │   │       # BE: timesheet-be/time-entry
    │               │   │       # POST /v1/timesheets/manual-entry
    │               │   ├── work-diary/
    │               │   │   ├── [date]/
    │               │   │   │   └── index.tsx  # Work diary for date (mobile)
    │               │   │   │       # - Daily time entries
    │               │   │   │       # - Screenshots (if enabled)
    │               │   │   │       # - Activity levels
    │               │   │   │       # - Tasks completed
    │               │   │   │       # - Edit entries
    │               │   │   │       # BE: timesheet-be/work-diary
    │               │   │   │       # GET /v1/timesheets/work-diary/{date}
    │               │   │   └── index.tsx  # Work diary overview (mobile)
    │               │   │       # - Weekly work diary
    │               │   │       # - Daily summaries
    │               │   │       # - Activity tracking
    │               │   │       # - Screenshot gallery
    │               │   │       # - Navigate by date
    │               │   │       # BE: timesheet-be/work-diary
    │               │   │       # GET /v1/timesheets/work-diary
    │               │   ├── weekly/
    │               │   │   └── index.tsx  # Weekly timesheet (mobile)
    │               │   │       # - Current week timesheet
    │               │   │       # - Daily breakdown
    │               │   │       # - Total hours
    │               │   │       # - Submit for approval
    │               │   │       # BE: timesheet-be/timesheet
    │               │   │       # GET /v1/timesheets/weekly
    │               │   └── index.tsx  # Timesheet overview (mobile)
    │               │       # - Recent timesheets
    │               │       # - Quick time entry
    │               │       # - Pending approvals
    │               │       # - Weekly summary
    │               │       # BE: timesheet-be/timesheet
    │               │       # GET /v1/timesheets
    │               │
    │               ├── certifications/  # ❌ MISSING ENTIRE FEATURE
    │               │   ├── [certId]/
    │               │   │   ├── details.tsx  # Certification details (mobile)
    │               │   │   │   # - Certification information
    │               │   │   │   # - Issuing authority
    │               │   │   │   # - Issue/expiry dates
    │               │   │   │   # - Verification status
    │               │   │   │   # - Certificate document
    │               │   │   │   # - Share button
    │               │   │   │   # BE: certifications-be/certification
    │               │   │   │   # GET /v1/certifications/{cert_id}
    │               │   │   └── verify.tsx  # Verify certification (mobile)
    │               │   │       # - Verification form
    │               │   │       # - Document upload
    │               │   │       # - Issuer information
    │               │   │       # - Request verification button
    │               │   │       # BE: certifications-be/verification
    │               │   │       # POST /v1/certifications/{cert_id}/verify
    │               │   ├── add/
    │               │   │   └── index.tsx  # Add certification (mobile)
    │               │   │       # - Add certification form
    │               │   │       # - Title and issuer
    │               │   │       # - Issue/expiry dates
    │               │   │       # - Certificate number
    │               │   │       # - Document upload
    │               │   │       # - Save button
    │               │   │       # BE: certifications-be/certification
    │               │   │       # POST /v1/certifications
    │               │   ├── pending-verification/
    │               │   │   └── index.tsx  # Pending verification (mobile)
    │               │   │       # - Certifications awaiting verification
    │               │   │       # - Verification status
    │               │   │       # - Required actions
    │               │   │       # - Resubmit option
    │               │   │       # BE: certifications-be/certification
    │               │   │       # GET /v1/certifications?status=pending_verification
    │               │   └── index.tsx  # Certifications overview (mobile)
    │               │       # - All certifications
    │               │       # - Verified/unverified status
    │               │       # - Expiring soon alerts
    │               │       # - Add new button
    │               │       # BE: certifications-be/certification
    │               │       # GET /v1/certifications
    │               │
    │               ├── skills-tests/  # ❌ MISSING ENTIRE FEATURE
    │               │   ├── [testId]/
    │               │   │   ├── start.tsx  # Start skills test (mobile)
    │               │   │   │   # - Test overview
    │               │   │   │   # - Instructions
    │               │   │   │   # - Duration and format
    │               │   │   │   # - Prerequisites
    │               │   │   │   # - Start test button
    │               │   │   │   # BE: skills-tests-be/test
    │               │   │   │   # POST /v1/skills-tests/{test_id}/start
    │               │   │   ├── test.tsx  # Take skills test (mobile)
    │               │   │   │   # - Test questions
    │               │   │   │   # - Answer inputs
    │               │   │   │   # - Timer
    │               │   │   │   # - Progress indicator
    │               │   │   │   # - Submit answers
    │               │   │   │   # BE: skills-tests-be/test-session
    │               │   │   │   # POST /v1/skills-tests/{test_id}/submit-answer
    │               │   │   └── results.tsx  # Test results (mobile)
    │               │   │       # - Score and performance
    │               │   │       # - Correct/incorrect breakdown
    │               │   │       # - Percentile ranking
    │               │   │       # - Badge earned
    │               │   │       # - Share results
    │               │   │       # BE: skills-tests-be/test-result
    │               │   │       # GET /v1/skills-tests/{test_id}/results
    │               │   ├── available/
    │               │   │   └── index.tsx  # Available tests (mobile)
    │               │   │       # - Browse all tests
    │               │   │       # - Filter by category/skill
    │               │   │       # - Difficulty levels
    │               │   │       # - Test previews
    │               │   │       # - Take test button
    │               │   │       # BE: skills-tests-be/test
    │               │   │       # GET /v1/skills-tests/available
    │               │   ├── completed/
    │               │   │   └── index.tsx  # Completed tests (mobile)
    │               │   │       # - All completed tests
    │               │   │       # - Scores and badges
    │               │   │       # - Retake options
    │               │   │       # - View results
    │               │   │       # BE: skills-tests-be/test-result
    │               │   │       # GET /v1/skills-tests/completed
    │               │   └── index.tsx  # Skills tests overview (mobile)
    │               │       # - Recommended tests
    │               │       # - Recent results
    │               │       # - Badges earned
    │               │       # - Browse all tests
    │               │       # BE: skills-tests-be/test
    │               │       # GET /v1/skills-tests
    │               │
    │               ├── groups/  # ❌ MISSING ENTIRE FEATURE
    │               │   ├── [groupId]/
    │               │   │   ├── details.tsx  # Group details (mobile)
    │               │   │   │   # - Group information
    │               │   │   │   # - Description and rules
    │               │   │   │   # - Member count
    │               │   │   │   # - Activity feed
    │               │   │   │   # - Join/leave button
    │               │   │   │   # BE: groups-be/group
    │               │   │   │   # GET /v1/groups/{group_id}
    │               │   │   ├── members.tsx  # Group members (mobile)
    │               │   │   │   # - All group members
    │               │   │   │   # - Member profiles
    │               │   │   │   # - Roles (admin/moderator/member)
    │               │   │   │   # - Search members
    │               │   │   │   # - Invite button
    │               │   │   │   # BE: groups-be/group-member
    │               │   │   │   # GET /v1/groups/{group_id}/members
    │               │   │   ├── posts/
    │               │   │   │   ├── [postId]/
    │               │   │   │   │   └── index.tsx  # Group post detail (mobile)
    │               │   │   │   │       # - Full post content
    │               │   │   │   │       # - Comments thread
    │               │   │   │   │       # - Like/react
    │               │   │   │   │       # - Share post
    │               │   │   │   │       # - Edit/delete (if owner)
    │               │   │   │   │       # BE: groups-be/group-post
    │               │   │   │   │       # GET /v1/groups/{group_id}/posts/{post_id}
    │               │   │   │   └── create.tsx  # Create group post (mobile)
    │               │   │   │       # - Post creation form
    │               │   │   │       # - Title and content
    │               │   │   │       # - Media upload
    │               │   │   │       # - Tags
    │               │   │   │       # - Publish button
    │               │   │   │       # BE: groups-be/group-post
    │               │   │   │       # POST /v1/groups/{group_id}/posts
    │               │   │   └── events/
    │               │   │       └── index.tsx  # Group events (mobile)
    │               │   │           # - Upcoming group events
    │               │   │           # - Past events
    │               │   │           # - RSVP status
    │               │   │           # - Create event button
    │               │   │           # BE: groups-be/group-event
    │               │   │           # GET /v1/groups/{group_id}/events
    │               │   ├── discover/
    │               │   │   └── index.tsx  # Discover groups (mobile)
    │               │   │       # - Browse all groups
    │               │   │       # - Search and filters
    │               │   │       # - Recommended groups
    │               │   │       # - Popular groups
    │               │   │       # - Join buttons
    │               │   │       # BE: groups-be/group
    │               │   │       # GET /v1/groups/discover
    │               │   ├── my-groups/
    │               │   │   └── index.tsx  # My groups (mobile)
    │               │   │       # - Groups I'm member of
    │               │   │       # - Recent activity
    │               │   │       # - Notifications
    │               │   │       # - Leave group option
    │               │   │       # BE: groups-be/group
    │               │   │       # GET /v1/groups/my-groups
    │               │   └── create/
    │               │       └── index.tsx  # Create group (mobile)
    │               │           # - Create group form
    │               │           # - Name and description
    │               │           # - Privacy settings
    │               │           # - Category/tags
    │               │           # - Cover image
    │               │           # - Create button
    │               │           # BE: groups-be/group
    │               │           # POST /v1/groups
    │
    ├── packages/
    │   ├── hooks/
    │   │   ├── contests/  # ❌ MISSING ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-contest.ts  # Single contest hook
    │   │   │   │   # Hook: useContest(contestId)
    │   │   │   │   # Returns: contest data, loading, error
    │   │   │   │   # BE: contests-be/contest GET /v1/contests/{contest_id}
    │   │   │   ├── use-contests.ts  # Contests list hook
    │   │   │   │   # Hook: useContests(filters)
    │   │   │   │   # Returns: contests[], pagination, loading, error
    │   │   │   │   # BE: contests-be/contest GET /v1/contests
    │   │   │   ├── use-contest-entries.ts  # Contest entries hook
    │   │   │   │   # Hook: useContestEntries(contestId)
    │   │   │   │   # Returns: entries[], loading, error
    │   │   │   │   # BE: contests-be/entry GET /v1/contests/{contest_id}/entries
    │   │   │   ├── use-submit-entry.ts  # Submit entry mutation
    │   │   │   │   # Hook: useSubmitEntry()
    │   │   │   │   # Returns: submitEntry(), loading, error
    │   │   │   │   # BE: contests-be/entry POST /v1/contests/{contest_id}/entries
    │   │   │   └── use-contest-leaderboard.ts  # Leaderboard hook
    │   │   │       # Hook: useContestLeaderboard(contestId)
    │   │   │       # Returns: leaderboard[], loading, error
    │   │   │       # BE: contests-be/leaderboard GET /v1/contests/{contest_id}/leaderboard
    │   │   │
    │   │   ├── talent-cloud/  # ❌ MISSING ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-agency.ts  # Single agency hook
    │   │   │   │   # Hook: useAgency(agencyId)
    │   │   │   │   # Returns: agency data, loading, error
    │   │   │   │   # BE: talent-cloud-be/agency GET /v1/agencies/{agency_id}
    │   │   │   ├── use-agencies.ts  # Agencies list hook
    │   │   │   │   # Hook: useAgencies(filters)
    │   │   │   │   # Returns: agencies[], pagination, loading, error
    │   │   │   │   # BE: talent-cloud-be/agency GET /v1/agencies
    │   │   │   ├── use-agency-projects.ts  # Agency projects hook
    │   │   │   │   # Hook: useAgencyProjects(agencyId)
    │   │   │   │   # Returns: projects[], loading, error
    │   │   │   │   # BE: talent-cloud-be/project GET /v1/agencies/{agency_id}/projects
    │   │   │   ├── use-team.ts  # Single team hook
    │   │   │   │   # Hook: useTeam(teamId)
    │   │   │   │   # Returns: team data, loading, error
    │   │   │   │   # BE: talent-cloud-be/team GET /v1/teams/{team_id}
    │   │   │   └── use-teams.ts  # Teams list hook
    │   │   │       # Hook: useTeams(filters)
    │   │   │       # Returns: teams[], pagination, loading, error
    │   │   │       # BE: talent-cloud-be/team GET /v1/teams
    │   │   │
    │   │   ├── skills-tests/  # ❌ MISSING ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-test.ts  # Single test hook
    │   │   │   │   # Hook: useTest(testId)
    │   │   │   │   # Returns: test data, loading, error
    │   │   │   │   # BE: skills-tests-be/test GET /v1/skills-tests/{test_id}
    │   │   │   ├── use-tests.ts  # Tests list hook
    │   │   │   │   # Hook: useTests(filters)
    │   │   │   │   # Returns: tests[], pagination, loading, error
    │   │   │   │   # BE: skills-tests-be/test GET /v1/skills-tests
    │   │   │   ├── use-start-test.ts  # Start test mutation
    │   │   │   │   # Hook: useStartTest()
    │   │   │   │   # Returns: startTest(), loading, error
    │   │   │   │   # BE: skills-tests-be/test POST /v1/skills-tests/{test_id}/start
    │   │   │   ├── use-submit-test.ts  # Submit test mutation
    │   │   │   │   # Hook: useSubmitTest()
    │   │   │   │   # Returns: submitTest(), loading, error
    │   │   │   │   # BE: skills-tests-be/test-session POST /v1/skills-tests/{test_id}/submit
    │   │   │   └── use-test-results.ts  # Test results hook
    │   │   │       # Hook: useTestResults(testId)
    │   │   │       # Returns: results data, loading, error
    │   │   │       # BE: skills-tests-be/test-result GET /v1/skills-tests/{test_id}/results
    │   │   │
    │   │   ├── certifications/  # ❌ MISSING ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-certification.ts  # Single certification hook
    │   │   │   │   # Hook: useCertification(certId)
    │   │   │   │   # Returns: certification data, loading, error
    │   │   │   │   # BE: certifications-be/certification GET /v1/certifications/{cert_id}
    │   │   │   ├── use-certifications.ts  # Certifications list hook
    │   │   │   │   # Hook: useCertifications(filters)
    │   │   │   │   # Returns: certifications[], loading, error
    │   │   │   │   # BE: certifications-be/certification GET /v1/certifications
    │   │   │   ├── use-add-certification.ts  # Add certification mutation
    │   │   │   │   # Hook: useAddCertification()
    │   │   │   │   # Returns: addCertification(), loading, error
    │   │   │   │   # BE: certifications-be/certification POST /v1/certifications
    │   │   │   └── use-verify-certification.ts  # Verify certification mutation
    │   │   │       # Hook: useVerifyCertification()
    │   │   │       # Returns: verifyCertification(), loading, error
    │   │   │       # BE: certifications-be/verification POST /v1/certifications/{cert_id}/verify
    │   │   │
    │   │   ├── groups/  # ❌ MISSING ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-group.ts  # Single group hook
    │   │   │   │   # Hook: useGroup(groupId)
    │   │   │   │   # Returns: group data, loading, error
    │   │   │   │   # BE: groups-be/group GET /v1/groups/{group_id}
    │   │   │   ├── use-groups.ts  # Groups list hook
    │   │   │   │   # Hook: useGroups(filters)
    │   │   │   │   # Returns: groups[], pagination, loading, error
    │   │   │   │   # BE: groups-be/group GET /v1/groups
    │   │   │   ├── use-group-members.ts  # Group members hook
    │   │   │   │   # Hook: useGroupMembers(groupId)
    │   │   │   │   # Returns: members[], loading, error
    │   │   │   │   # BE: groups-be/group-member GET /v1/groups/{group_id}/members
    │   │   │   ├── use-group-posts.ts  # Group posts hook
    │   │   │   │   # Hook: useGroupPosts(groupId)
    │   │   │   │   # Returns: posts[], loading, error
    │   │   │   │   # BE: groups-be/group-post GET /v1/groups/{group_id}/posts
    │   │   │   ├── use-create-group.ts  # Create group mutation
    │   │   │   │   # Hook: useCreateGroup()
    │   │   │   │   # Returns: createGroup(), loading, error
    │   │   │   │   # BE: groups-be/group POST /v1/groups
    │   │   │   └── use-join-group.ts  # Join group mutation
    │   │   │       # Hook: useJoinGroup()
    │   │   │       # Returns: joinGroup(), loading, error
    │   │   │       # BE: groups-be/group POST /v1/groups/{group_id}/join
    │   │   │
    │   │   ├── events/  # ❌ MISSING ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-event.ts  # Single event hook
    │   │   │   │   # Hook: useEvent(eventId)
    │   │   │   │   # Returns: event data, loading, error
    │   │   │   │   # BE: events-be/event GET /v1/events/{event_id}
    │   │   │   ├── use-events.ts  # Events list hook
    │   │   │   │   # Hook: useEvents(filters)
    │   │   │   │   # Returns: events[], pagination, loading, error
    │   │   │   │   # BE: events-be/event GET /v1/events
    │   │   │   ├── use-event-registration.ts  # Event registration mutation
    │   │   │   │   # Hook: useEventRegistration()
    │   │   │   │   # Returns: registerForEvent(), loading, error
    │   │   │   │   # BE: events-be/registration POST /v1/events/{event_id}/register
    │   │   │   ├── use-event-attendees.ts  # Event attendees hook
    │   │   │   │   # Hook: useEventAttendees(eventId)
    │   │   │   │   # Returns: attendees[], loading, error
    │   │   │   │   # BE: events-be/attendee GET /v1/events/{event_id}/attendees
    │   │   │   └── use-create-event.ts  # Create event mutation
    │   │   │       # Hook: useCreateEvent()
    │   │   │       # Returns: createEvent(), loading, error
    │   │   │       # BE: events-be/event POST /v1/events
    │   │   │
    │   │   ├── learning/  # ❌ MISSING ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-course.ts  # Single course hook
    │   │   │   │   # Hook: useCourse(courseId)
    │   │   │   │   # Returns: course data, loading, error
    │   │   │   │   # BE: learning-be/course GET /v1/courses/{course_id}
    │   │   │   ├── use-courses.ts  # Courses list hook
    │   │   │   │   # Hook: useCourses(filters)
    │   │   │   │   # Returns: courses[], pagination, loading, error
    │   │   │   │   # BE: learning-be/course GET /v1/courses
    │   │   │   ├── use-lesson.ts  # Single lesson hook
    │   │   │   │   # Hook: useLesson(courseId, lessonId)
    │   │   │   │   # Returns: lesson data, loading, error
    │   │   │   │   # BE: learning-be/lesson GET /v1/courses/{course_id}/lessons/{lesson_id}
    │   │   │   ├── use-lesson-progress.ts  # Lesson progress hook
    │   │   │   │   # Hook: useLessonProgress(lessonId)
    │   │   │   │   # Returns: progress data, markComplete(), loading, error
    │   │   │   │   # BE: learning-be/progress GET /v1/lessons/{lesson_id}/progress
    │   │   │   ├── use-learning-path.ts  # Learning path hook
    │   │   │   │   # Hook: useLearningPath(pathId)
    │   │   │   │   # Returns: path data, loading, error
    │   │   │   │   # BE: learning-be/learning-path GET /v1/learning-paths/{path_id}
    │   │   │   └── use-certificate.ts  # Certificate hook
    │   │   │       # Hook: useCertificate(courseId)
    │   │   │       # Returns: certificate data, loading, error
    │   │   │       # BE: learning-be/certificate GET /v1/courses/{course_id}/certificate
    │   │   │
    │   │   ├── referrals/  # ❌ MISSING ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-referral-code.ts  # Referral code hook
    │   │   │   │   # Hook: useReferralCode()
    │   │   │   │   # Returns: referral code, loading, error
    │   │   │   │   # BE: referrals-be/referral GET /v1/referrals/code
    │   │   │   ├── use-referral-stats.ts  # Referral stats hook
    │   │   │   │   # Hook: useReferralStats()
    │   │   │   │   # Returns: stats data, loading, error
    │   │   │   │   # BE: referrals-be/referral GET /v1/referrals/stats
    │   │   │   ├── use-referral-earnings.ts  # Referral earnings hook
    │   │   │   │   # Hook: useReferralEarnings()
    │   │   │   │   # Returns: earnings data, loading, error
    │   │   │   │   # BE: referrals-be/earnings GET /v1/referrals/earnings
    │   │   │   └── use-send-referral.ts  # Send referral mutation
    │   │   │       # Hook: useSendReferral()
    │   │   │       # Returns: sendReferral(), loading, error
    │   │   │       # BE: referrals-be/referral POST /v1/referrals/send
    │   │   │
    │   │   ├── service-catalog/  # ❌ MISSING ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-service.ts  # Single service hook
    │   │   │   │   # Hook: useService(serviceId)
    │   │   │   │   # Returns: service data, loading, error
    │   │   │   │   # BE: service-catalog-be/service GET /v1/services/{service_id}
    │   │   │   ├── use-services.ts  # Services list hook
    │   │   │   │   # Hook: useServices(filters)
    │   │   │   │   # Returns: services[], pagination, loading, error
    │   │   │   │   # BE: service-catalog-be/service GET /v1/services
    │   │   │   ├── use-create-service.ts  # Create service mutation
    │   │   │   │   # Hook: useCreateService()
    │   │   │   │   # Returns: createService(), loading, error
    │   │   │   │   # BE: service-catalog-be/service POST /v1/services
    │   │   │   ├── use-book-service.ts  # Book service mutation
    │   │   │   │   # Hook: useBookService()
    │   │   │   │   # Returns: bookService(), loading, error
    │   │   │   │   # BE: service-catalog-be/booking POST /v1/services/{service_id}/book
    │   │   │   └── use-service-reviews.ts  # Service reviews hook
    │   │   │       # Hook: useServiceReviews(serviceId)
    │   │   │       # Returns: reviews[], loading, error
    │   │   │       # BE: service-catalog-be/review GET /v1/services/{service_id}/reviews
    │   │   │
    │   │   ├── timesheets/  # ❌ MISSING ENTIRE DOMAIN
    │   │   │   ├── index.ts  # Barrel export
    │   │   │   ├── package.json
    │   │   │   ├── tsconfig.json
    │   │   │   ├── use-timesheet.ts  # Single timesheet hook
    │   │   │   │   # Hook: useTimesheet(timesheetId)
    │   │   │   │   # Returns: timesheet data, loading, error
    │   │   │   │   # BE: timesheet-be/timesheet GET /v1/timesheets/{timesheet_id}
    │   │   │   ├── use-timesheets.ts  # Timesheets list hook
    │   │   │   │   # Hook: useTimesheets(filters)
    │   │   │   │   # Returns: timesheets[], loading, error
    │   │   │   │   # BE: timesheet-be/timesheet GET /v1/timesheets
    │   │   │   ├── use-log-time.ts  # Log time mutation
    │   │   │   │   # Hook: useLogTime()
    │   │   │   │   # Returns: logTime(), loading, error
    │   │   │   │   # BE: timesheet-be/time-entry POST /v1/timesheets/log-time
    │   │   │   ├── use-work-diary.ts  # Work diary hook
    │   │   │   │   # Hook: useWorkDiary(date)
    │   │   │   │   # Returns: diary data, loading, error
    │   │   │   │   # BE: timesheet-be/work-diary GET /v1/timesheets/work-diary/{date}
    │   │   │   └── use-manual-time.ts  # Manual time entry mutation
    │   │   │       # Hook: useManualTime()
    │   │   │       # Returns: addManualTime(), loading, error
    │   │   │       # BE: timesheet-be/time-entry POST /v1/timesheets/manual-entry
    │   │   │
    │   │   └── disputes/  # ❌ MISSING ENTIRE DOMAIN
    │   │       ├── index.ts  # Barrel export
    │   │       ├── package.json
    │   │       ├── tsconfig.json
    │   │       ├── use-dispute.ts  # Single dispute hook
    │   │       │   # Hook: useDispute(disputeId)
    │   │       │   # Returns: dispute data, loading, error
    │   │       │   # BE: disputes-be/dispute GET /v1/disputes/{dispute_id}
    │   │       ├── use-disputes.ts  # Disputes list hook
    │   │       │   # Hook: useDisputes(filters)
    │   │       │   # Returns: disputes[], loading, error
    │   │       │   # BE: disputes-be/dispute GET /v1/disputes
    │   │       ├── use-file-dispute.ts  # File dispute mutation
    │   │       │   # Hook: useFileDispute()
    │   │       │   # Returns: fileDispute(), loading, error
    │   │       │   # BE: disputes-be/dispute POST /v1/disputes
    │   │       ├── use-dispute-evidence.ts  # Dispute evidence hook
    │   │       │   # Hook: useDisputeEvidence(disputeId)
    │   │       │   # Returns: evidence[], addEvidence(), loading, error
    │   │       │   # BE: disputes-be/evidence GET /v1/disputes/{dispute_id}/evidence
    │   │       └── use-dispute-mediation.ts  # Dispute mediation hook
    │   │           # Hook: useDisputeMediation(disputeId)
    │   │           # Returns: mediation data, loading, error
    │   │           # BE: disputes-be/mediation GET /v1/disputes/{dispute_id}/mediation
    │   │
    │   └── ui/
    │       └── src/
    │           └── components/
    │               ├── contests/  # ❌ MISSING ENTIRE SECTION
    │               │   ├── ContestCard.native.tsx  # Contest card (native)
    │               │   │   # - Contest preview card
    │               │   │   # - Title, prize, deadline
    │               │   │   # - Entry count
    │               │   │   # - Enter button
    │               │   ├── ContestCard.tsx  # Contest card (base)
    │               │   │   # - Shared contest card logic
    │               │   │   # - Props interface
    │               │   ├── ContestCard.web.tsx  # Contest card (web)
    │               │   │   # - Contest preview card
    │               │   │   # - Hover effects
    │               │   │   # - Responsive layout
    │               │   ├── EntryCard.native.tsx  # Entry card (native)
    │               │   │   # - Contest entry card
    │               │   │   # - Entry preview
    │               │   │   # - Vote button
    │               │   │   # - Score display
    │               │   ├── EntryCard.tsx  # Entry card (base)
    │               │   │   # - Shared entry card logic
    │               │   │   # - Props interface
    │               │   ├── EntryCard.web.tsx  # Entry card (web)
    │               │   │   # - Contest entry card
    │               │   │   # - Hover effects
    │               │   │   # - Vote interactions
    │               │   ├── Leaderboard.native.tsx  # Leaderboard (native)
    │               │   │   # - Contest leaderboard
    │               │   │   # - Ranked entries
    │               │   │   # - Scrollable list
    │               │   │   # - Winner highlights
    │               │   ├── Leaderboard.tsx  # Leaderboard (base)
    │               │   │   # - Shared leaderboard logic
    │               │   │   # - Props interface
    │               │   └── Leaderboard.web.tsx  # Leaderboard (web)
    │               │       # - Contest leaderboard
    │               │       # - Sortable columns
    │               │       # - Winner highlights
    │               │
    │               ├── escrow/  # ❌ MISSING ENTIRE SECTION
    │               │   ├── EscrowStatus.native.tsx  # Escrow status (native)
    │               │   │   # - Escrow status display
    │               │   │   # - Status badge
    │               │   │   # - Amount info
    │               │   │   # - Timeline
    │               │   ├── EscrowStatus.tsx  # Escrow status (base)
    │               │   │   # - Shared escrow status logic
    │               │   │   # - Props interface
    │               │   ├── EscrowStatus.web.tsx  # Escrow status (web)
    │               │   │   # - Escrow status display
    │               │   │   # - Status badge
    │               │   │   # - Detailed timeline
    │               │   ├── ReleaseButton.native.tsx  # Release button (native)
    │               │   │   # - Escrow release button
    │               │   │   # - Confirmation modal
    │               │   │   # - Loading state
    │               │   ├── ReleaseButton.tsx  # Release button (base)
    │               │   │   # - Shared release button logic
    │               │   │   # - Props interface
    │               │   └── ReleaseButton.web.tsx  # Release button (web)
    │               │       # - Escrow release button
    │               │       # - Confirmation modal
    │               │       # - Hover effects
    │               │
    │               ├── skills-tests/  # ❌ MISSING ENTIRE SECTION
    │               │   ├── TestCard.native.tsx  # Test card (native)
    │               │   │   # - Skills test card
    │               │   │   # - Test info
    │               │   │   # - Difficulty badge
    │               │   │   # - Take test button
    │               │   ├── TestCard.tsx  # Test card (base)
    │               │   │   # - Shared test card logic
    │               │   │   # - Props interface
    │               │   ├── TestCard.web.tsx  # Test card (web)
    │               │   │   # - Skills test card
    │               │   │   # - Hover effects
    │               │   │   # - Responsive layout
    │               │   ├── TestQuestion.native.tsx  # Test question (native)
    │               │   │   # - Test question display
    │               │   │   # - Answer options
    │               │   │   # - Selection handling
    │               │   │   # - Navigation buttons
    │               │   ├── TestQuestion.tsx  # Test question (base)
    │               │   │   # - Shared question logic
    │               │   │   # - Props interface
    │               │   ├── TestQuestion.web.tsx  # Test question (web)
    │               │   │   # - Test question display
    │               │   │   # - Answer options
    │               │   │   # - Keyboard navigation
    │               │   ├── TestResults.native.tsx  # Test results (native)
    │               │   │   # - Test results display
    │               │   │   # - Score visualization
    │               │   │   # - Breakdown
    │               │   │   # - Share button
    │               │   ├── TestResults.tsx  # Test results (base)
    │               │   │   # - Shared results logic
    │               │   │   # - Props interface
    │               │   └── TestResults.web.tsx  # Test results (web)
    │               │       # - Test results display
    │               │       # - Charts/graphs
    │               │       # - Detailed breakdown
    │               │
    │               ├── groups/  # ❌ MISSING ENTIRE SECTION
    │               │   ├── GroupCard.native.tsx  # Group card (native)
    │               │   │   # - Group preview card
    │               │   │   # - Group info
    │               │   │   # - Member count
    │               │   │   # - Join button
    │               │   ├── GroupCard.tsx  # Group card (base)
    │               │   │   # - Shared group card logic
    │               │   │   # - Props interface
    │               │   ├── GroupCard.web.tsx  # Group card (web)
    │               │   │   # - Group preview card
    │               │   │   # - Hover effects
    │               │   │   # - Responsive layout
    │               │   ├── GroupPost.native.tsx  # Group post (native)
    │               │   │   # - Group post card
    │               │   │   # - Post content
    │               │   │   # - Like/comment buttons
    │               │   │   # - Author info
    │               │   ├── GroupPost.tsx  # Group post (base)
    │               │   │   # - Shared post logic
    │               │   │   # - Props interface
    │               │   └── GroupPost.web.tsx  # Group post (web)
    │               │       # - Group post card
    │               │       # - Rich interactions
    │               │       # - Comment thread
    │               │
    │               ├── events/  # ❌ MISSING ENTIRE SECTION
    │               │   ├── EventCard.native.tsx  # Event card (native)
    │               │   │   # - Event preview card
    │               │   │   # - Date/time/location
    │               │   │   # - Register button
    │               │   │   # - Attendees count
    │               │   ├── EventCard.tsx  # Event card (base)
    │               │   │   # - Shared event card logic
    │               │   │   # - Props interface
    │               │   ├── EventCard.web.tsx  # Event card (web)
    │               │   │   # - Event preview card
    │               │   │   # - Hover effects
    │               │   │   # - Responsive layout
    │               │   ├── EventCalendar.native.tsx  # Event calendar (native)
    │               │   │   # - Calendar view
    │               │   │   # - Event markers
    │               │   │   # - Date navigation
    │               │   │   # - Touch interactions
    │               │   ├── EventCalendar.tsx  # Event calendar (base)
    │               │   │   # - Shared calendar logic
    │               │   │   # - Props interface
    │               │   └── EventCalendar.web.tsx  # Event calendar (web)
    │               │       # - Calendar view
    │               │       # - Event markers
    │               │       # - Click interactions
    │               │
    │               ├── learning/  # ❌ MISSING ENTIRE SECTION
    │               │   ├── CourseCard.native.tsx  # Course card (native)
    │               │   │   # - Course preview card
    │               │   │   # - Course info
    │               │   │   # - Progress bar
    │               │   │   # - Enroll button
    │               │   ├── CourseCard.tsx  # Course card (base)
    │               │   │   # - Shared course card logic
    │               │   │   # - Props interface
    │               │   ├── CourseCard.web.tsx  # Course card (web)
    │               │   │   # - Course preview card
    │               │   │   # - Hover effects
    │               │   │   # - Responsive layout
    │               │   ├── LessonPlayer.native.tsx  # Lesson player (native)
    │               │   │   # - Video/content player
    │               │   │   # - Playback controls
    │               │   │   # - Progress tracking
    │               │   │   # - Next lesson
    │               │   ├── LessonPlayer.tsx  # Lesson player (base)
    │               │   │   # - Shared player logic
    │               │   │   # - Props interface
    │               │   ├── LessonPlayer.web.tsx  # Lesson player (web)
    │               │   │   # - Video/content player
    │               │   │   # - Full controls
    │               │   │   # - Keyboard shortcuts
    │               │   ├── ProgressTracker.native.tsx  # Progress tracker (native)
    │               │   │   # - Learning progress display
    │               │   │   # - Progress bars
    │               │   │   # - Milestones
    │               │   │   # - Achievements
    │               │   ├── ProgressTracker.tsx  # Progress tracker (base)
    │               │   │   # - Shared progress logic
    │               │   │   # - Props interface
    │               │   └── ProgressTracker.web.tsx  # Progress tracker (web)
    │               │       # - Learning progress display
    │               │       # - Charts/graphs
    │               │       # - Detailed stats
    │               │
    │               └── referrals/  # ❌ MISSING ENTIRE SECTION
    │                   ├── ReferralCard.native.tsx  # Referral card (native)
    │                   │   # - Referral info card
    │                   │   # - Referral status
    │                   │   # - Earnings display
    │                   │   # - Share button
    │                   ├── ReferralCard.tsx  # Referral card (base)
    │                   │   # - Shared referral card logic
    │                   │   # - Props interface
    │                   ├── ReferralCard.web.tsx  # Referral card (web)
    │                   │   # - Referral info card
    │                   │   # - Hover effects
    │                   │   # - Responsive layout
    │                   ├── ReferralStats.native.tsx  # Referral stats (native)
    │                   │   # - Statistics display
    │                   │   # - Total referrals
    │                   │   # - Earnings chart
    │                   │   # - Conversion rate
    │                   ├── ReferralStats.tsx  # Referral stats (base)
    │                   │   # - Shared stats logic
    │                   │   # - Props interface
    │                   └── ReferralStats.web.tsx  # Referral stats (web)
    │                       # - Statistics display
    │                       # - Interactive charts
    │                       # - Detailed analytics
