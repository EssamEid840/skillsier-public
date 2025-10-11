# Skillsier Platform - Complete Microservices Folder Structures (Updated)



## 1️⃣ users-be (REFERENCE - UNCHANGED)

```
users-be/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── domain/
│   │   ├── user/
│   │   │   ├── entity.go              # User aggregate root
│   │   │   ├── value_objects.go       # Email, Phone, Address
│   │   │   ├── enums.go               # UserType, AccountStatus, VerificationStatus
│   │   │   └── repository.go
│   │   ├── profile/
│   │   │   ├── entity.go              # Extended profile info
│   │   │   ├── preferences.go         # User preferences, settings
│   │   │   ├── availability.go        # Work availability, timezone
│   │   │   └── repository.go
│   │   ├── skill/
│   │   │   ├── entity.go              # User skills
│   │   │   ├── proficiency.go         # Skill level enum
│   │   │   └── repository.go
│   │   ├── experience/
│   │   │   ├── entity.go              # Work experience
│   │   │   └── repository.go
│   │   ├── education/
│   │   │   ├── entity.go              # Educational background
│   │   │   └── repository.go
│   │   ├── certification/
│   │   │   ├── entity.go              # Certifications
│   │   │   ├── verification.go        # Verification status
│   │   │   └── repository.go
│   │   ├── portfolio/
│   │   │   ├── entity.go              # Portfolio items
│   │   │   ├── item.go                # Portfolio item details
│   │   │   ├── media.go               # Associated media
│   │   │   └── repository.go
│   │   ├── language/
│   │   │   ├── entity.go              # Language proficiency
│   │   │   └── repository.go
│   │   ├── freelancer/
│   │   │   ├── entity.go              # Freelancer-specific data
│   │   │   ├── profile.go             # Freelancer profile
│   │   │   ├── rates.go               # Hourly/fixed rates
│   │   │   ├── stats.go               # Job stats, earnings
│   │   │   └── repository.go
│   │   ├── client/
│   │   │   ├── entity.go              # Client-specific data
│   │   │   ├── profile.go             # Client profile
│   │   │   ├── company.go             # Company info
│   │   │   ├── stats.go               # Hiring stats, spending
│   │   │   └── repository.go
│   │   ├── verification/
│   │   │   ├── entity.go              # Identity verification
│   │   │   ├── document.go            # Verification documents
│   │   │   └── repository.go
│   │   ├── settings/
│   │   │   ├── entity.go              # User settings
│   │   │   ├── notification_prefs.go  # Notification preferences
│   │   │   ├── privacy_settings.go    # Privacy settings
│   │   │   └── repository.go
│   │   ├── saved_items/
│   │   │   ├── entity.go              # Saved jobs, freelancers
│   │   │   └── repository.go
│   │   ├── blocked_users/
│   │   │   ├── entity.go              # Blocked user relationships
│   │   │   └── repository.go
│   │   ├── user_suspension/
│   │   │   ├── entity.go              # Track user suspensions
│   │   │   ├── reason.go              # Suspension reasons
│   │   │   ├── duration.go            # Suspension duration
│   │   │   └── repository.go
│   │   ├── user_ban/
│   │   │   ├── entity.go              # Track user bans
│   │   │   ├── reason.go              # Ban reasons
│   │   │   ├── permanent.go           # Permanent vs temporary
│   │   │   └── repository.go
│   │   ├── user_warning/
│   │   │   ├── entity.go              # Track user warnings
│   │   │   ├── reason.go              # Warning reasons
│   │   │   ├── severity.go            # Warning severity level
│   │   │   └── repository.go
│   │   └── outbox/
│   │       ├── entity.go
│   │       └── repository.go
│   │
│   ├── application/
│   │   ├── user/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── profile/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── skill/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── experience/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── education/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── certification/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── portfolio/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── language/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── freelancer/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── stats_calculator.go
│   │   ├── client/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── stats_calculator.go
│   │   ├── verification/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── settings/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── saved_items/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── suspension/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Suspend, Unsuspend
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── ban/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Ban, Unban
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── warning/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Issue warning
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   └── eventhandler/
│   │       ├── keycloak_handler.go
│   │       ├── contract_handler.go
│   │       ├── review_handler.go
│   │       ├── transaction_handler.go
│   │       ├── badge_handler.go
│   │       └── admin_handler.go       # Handle admin actions on users
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go
│   │   │       ├── transaction.go
│   │   │       ├── migrations.go          # Auto-migrate logic
│   │   │       ├── user_repository.go
│   │   │       ├── profile_repository.go
│   │   │       ├── skill_repository.go
│   │   │       ├── experience_repository.go
│   │   │       ├── education_repository.go
│   │   │       ├── certification_repository.go
│   │   │       ├── portfolio_repository.go
│   │   │       ├── language_repository.go
│   │   │       ├── freelancer_repository.go
│   │   │       ├── client_repository.go
│   │   │       ├── verification_repository.go
│   │   │       ├── settings_repository.go
│   │   │       ├── saved_items_repository.go
│   │   │       ├── blocked_users_repository.go
│   │   │       ├── user_suspension_repository.go
│   │   │       ├── user_ban_repository.go
│   │   │       ├── user_warning_repository.go
│   │   │       └── outbox_repository.go
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go
│   │   │       ├── user_cache.go
│   │   │       └── profile_cache.go
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go
│   │   │       ├── producer.go
│   │   │       ├── topics.go
│   │   │       └── scram.go
│   │   ├── storage/
│   │   │   ├── client.go              # Storage service client
│   │   │   └── local.go
│   │   ├── keycloak/
│   │   │   ├── client.go
│   │   │   └── sync.go
│   │   └── outbox/
│   │       ├── processor.go
│   │       └── scheduler.go
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── user_handler.go
│   │       │   ├── profile_handler.go
│   │       │   ├── skill_handler.go
│   │       │   ├── experience_handler.go
│   │       │   ├── education_handler.go
│   │       │   ├── certification_handler.go
│   │       │   ├── portfolio_handler.go
│   │       │   ├── language_handler.go
│   │       │   ├── freelancer_handler.go
│   │       │   ├── client_handler.go
│   │       │   ├── verification_handler.go
│   │       │   ├── settings_handler.go
│   │       │   ├── saved_items_handler.go
│   │       │   ├── suspension_handler.go
│   │       │   ├── ban_handler.go
│   │       │   ├── warning_handler.go
│   │       │   └── health_handler.go
│   │       ├── middleware/
│   │       │   ├── auth.go
│   │       │   ├── rbac.go
│   │       │   ├── cors.go
│   │       │   ├── rate_limit.go
│   │       │   ├── logging.go
│   │       │   ├── error_handler.go
│   │       │   └── request_id.go
│   │       ├── responses/
│   │       │   ├── success.go
│   │       │   ├── error.go
│   │       │   └── pagination.go
│   │       └── router.go
│   │
│   └── config/
│       ├── config.go
│       ├── database.go
│       ├── kafka.go
│       ├── redis.go
│       └── keycloak.go
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   └── codes.go
│   ├── logger/
│   │   └── logger.go
│   ├── utils/
│   │   ├── validator.go
│   │   ├── slug.go
│   │   └── sanitizer.go
│   └── constants/
│       ├── events.go
│       └── topics.go
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   └── seed-data.sh
│
├── tests/
│   ├── unit/
│   │   ├── domain/
│   │   ├── application/
│   │   └── infrastructure/
│   ├── integration/
│   │   ├── handlers/
│   │   └── repositories/
│   └── e2e/
│       └── scenarios/
│
├── docs/
│   ├── README.md
│   ├── api.md
│   ├── events.md
│   └── architecture.md
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md
```

---

## 2️⃣ jobs-be (UPDATED)

```
jobs-be/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── domain/
│   │   ├── job/
│   │   │   ├── entity.go              # Job aggregate root
│   │   │   ├── value_objects.go       # Budget, Duration
│   │   │   ├── enums.go               # JobType, Status, ExperienceLevel, BudgetType
│   │   │   ├── visibility.go          # Public, Private, Invite-only
│   │   │   └── repository.go
│   │   ├── category/
│   │   │   ├── entity.go              # Job categories
│   │   │   ├── subcategory.go         # Nested categories
│   │   │   └── repository.go
│   │   ├── skill/
│   │   │   ├── entity.go              # Skill taxonomy
│   │   │   ├── category.go            # Skill categories
│   │   │   └── repository.go
│   │   ├── job_skill/
│   │   │   ├── entity.go              # Skills required for job
│   │   │   ├── requirement_level.go   # Required, Preferred
│   │   │   └── repository.go
│   │   ├── job_question/
│   │   │   ├── entity.go              # Screening questions
│   │   │   ├── question_type.go       # Text, MultiChoice, File
│   │   │   └── repository.go
│   │   ├── job_attachment/
│   │   │   ├── entity.go              # Job attachments
│   │   │   └── repository.go
│   │   ├── job_invitation/
│   │   │   ├── entity.go              # Invitations to freelancers
│   │   │   ├── status.go              # Invitation status
│   │   │   └── repository.go
│   │   ├── job_view/
│   │   │   ├── entity.go              # Job view tracking
│   │   │   └── repository.go
│   │   ├── saved_search/
│   │   │   ├── entity.go              # Saved search filters
│   │   │   └── repository.go
│   │   ├── job_flag/
│   │   │   ├── entity.go              # Flagged jobs
│   │   │   ├── reason.go              # Flag reasons
│   │   │   ├── status.go              # Flag status (Pending, Reviewed, Resolved)
│   │   │   └── repository.go
│   │   └── outbox/
│   │       ├── entity.go
│   │       └── repository.go
│   │
│   ├── application/
│   │   ├── job/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Update, Close, Delete, Repost
│   │   │   ├── queries.go             # Get, List, Search, Filter
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── business_rules.go
│   │   ├── category/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── tree_builder.go
│   │   ├── skill/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── suggestion_service.go
│   │   ├── invitation/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── search/
│   │   │   ├── service.go
│   │   │   ├── filter_builder.go
│   │   │   └── dto.go
│   │   ├── flag/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Flag, Unflag job
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   └── eventhandler/
│   │       ├── user_handler.go
│   │       ├── proposal_handler.go
│   │       ├── contract_handler.go
│   │       ├── subscription_handler.go
│   │       └── admin_handler.go       # Handle admin actions on jobs
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go
│   │   │       ├── transaction.go
│   │   │       ├── migrations.go          # Auto-migrate logic
│   │   │       ├── job_repository.go
│   │   │       ├── category_repository.go
│   │   │       ├── skill_repository.go
│   │   │       ├── job_skill_repository.go
│   │   │       ├── job_question_repository.go
│   │   │       ├── job_attachment_repository.go
│   │   │       ├── job_invitation_repository.go
│   │   │       ├── job_view_repository.go
│   │   │       ├── saved_search_repository.go
│   │   │       ├── job_flag_repository.go
│   │   │       └── outbox_repository.go
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go
│   │   │       ├── job_cache.go
│   │   │       ├── category_cache.go
│   │   │       └── skill_cache.go
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go
│   │   │       ├── producer.go
│   │   │       ├── topics.go
│   │   │       └── scram.go
│   │   ├── search/
│   │   │   └── client.go              # Search service client
│   │   ├── storage/
│   │   │   └── client.go
│   │   └── outbox/
│   │       ├── processor.go
│   │       └── scheduler.go
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── job_handler.go
│   │       │   ├── category_handler.go
│   │       │   ├── skill_handler.go
│   │       │   ├── invitation_handler.go
│   │       │   ├── search_handler.go
│   │       │   ├── saved_search_handler.go
│   │       │   ├── flag_handler.go
│   │       │   └── health_handler.go
│   │       ├── middleware/
│   │       │   ├── auth.go
│   │       │   ├── rbac.go
│   │       │   ├── cors.go
│   │       │   ├── rate_limit.go
│   │       │   ├── logging.go
│   │       │   ├── error_handler.go
│   │       │   └── request_id.go
│   │       ├── responses/
│   │       │   ├── success.go
│   │       │   ├── error.go
│   │       │   └── pagination.go
│   │       └── router.go
│   │
│   └── config/
│       ├── config.go
│       ├── database.go
│       ├── kafka.go
│       ├── redis.go
│       └── services.go
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   └── codes.go
│   ├── logger/
│   │   └── logger.go
│   ├── utils/
│   │   ├── validator.go
│   │   ├── slug.go
│   │   └── sanitizer.go
│   └── constants/
│       ├── events.go
│       └── topics.go
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   ├── seed-categories.sh
│   └── seed-skills.sh
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── api.md
│   ├── events.md
│   ├── categories.md
│   └── skills-taxonomy.md
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md
```

---

## 3️⃣ proposals-be (UPDATED WITH BIDDING SYSTEM)

```
proposals-be/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── domain/
│   │   ├── proposal/
│   │   │   ├── entity.go              # Proposal aggregate root
│   │   │   ├── value_objects.go       # Amount, Duration
│   │   │   ├── enums.go               # Status, Type, ProposalType
│   │   │   └── repository.go
│   │   ├── bid/
│   │   │   ├── entity.go              # Bidding system
│   │   │   ├── bid_amount.go          # Bid amount management
│   │   │   ├── bid_history.go         # Bid history tracking
│   │   │   ├── auto_bid.go            # Auto-bidding logic
│   │   │   ├── bid_rank.go            # Bid ranking
│   │   │   └── repository.go
│   │   ├── bid_strategy/
│   │   │   ├── entity.go              # Bidding strategies
│   │   │   ├── aggressive.go          # Aggressive bidding
│   │   │   ├── conservative.go        # Conservative bidding
│   │   │   ├── auto_strategy.go       # Auto-bid strategy
│   │   │   └── repository.go
│   │   ├── bid_notification/
│   │   │   ├── entity.go              # Bid notifications
│   │   │   ├── outbid_alert.go        # Outbid alerts
│   │   │   └── repository.go
│   │   ├── milestone/
│   │   │   ├── entity.go              # Proposal milestones
│   │   │   ├── milestone_group.go
│   │   │   └── repository.go
│   │   ├── cover_letter/
│   │   │   ├── entity.go              # Cover letter content
│   │   │   └── repository.go
│   │   ├── proposal_attachment/
│   │   │   ├── entity.go              # Proposal attachments
│   │   │   └── repository.go
│   │   ├── proposal_question_answer/
│   │   │   ├── entity.go              # Answers to job questions
│   │   │   └── repository.go
│   │   ├── template/
│   │   │   ├── entity.go              # Proposal templates
│   │   │   ├── template_category.go
│   │   │   └── repository.go
│   │   ├── connect/
│   │   │   ├── entity.go              # Connect usage tracking
│   │   │   ├── transaction.go         # Connect transactions
│   │   │   └── repository.go
│   │   ├── boost/
│   │   │   ├── entity.go              # Boosted proposals
│   │   │   ├── boost_level.go         # Boost levels
│   │   │   └── repository.go
│   │   ├── proposal_analytics/
│   │   │   ├── entity.go              # Proposal performance analytics
│   │   │   ├── view_tracking.go       # Track views
│   │   │   └── repository.go
│   │   ├── proposal_flag/
│   │   │   ├── entity.go              # Flagged proposals
│   │   │   ├── reason.go              # Flag reasons
│   │   │   ├── status.go              # Flag status
│   │   │   └── repository.go
│   │   └── outbox/
│   │       ├── entity.go
│   │       └── repository.go
│   │
│   ├── application/
│   │   ├── proposal/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Submit, Withdraw, Update
│   │   │   ├── queries.go             # Get, List, Filter
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── business_rules.go
│   │   ├── bid/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # PlaceBid, UpdateBid, WithdrawBid
│   │   │   ├── queries.go             # GetBidStatus, GetBidHistory
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── bid_calculator.go      # Calculate optimal bids
│   │   │   ├── bid_validator.go       # Validate bid amounts
│   │   │   ├── auto_bid_manager.go    # Auto-bidding management
│   │   │   └── ranking_engine.go      # Rank bids
│   │   ├── bid_strategy/
│   │   │   ├── service.go
│   │   │   ├── strategy_executor.go   # Execute bidding strategies
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── milestone/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── template/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── connect/
│   │   │   ├── service.go
│   │   │   ├── calculator.go          # Calculate connect cost
│   │   │   └── dto.go
│   │   ├── boost/
│   │   │   ├── service.go
│   │   │   └── dto.go
│   │   ├── analytics/
│   │   │   ├── service.go
│   │   │   ├── performance_tracker.go
│   │   │   └── dto.go
│   │   ├── flag/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Flag, Unflag proposal
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   └── eventhandler/
│   │       ├── job_handler.go
│   │       ├── user_handler.go
│   │       ├── contract_handler.go
│   │       ├── subscription_handler.go
│   │       ├── payment_handler.go
│   │       └── admin_handler.go       # Handle admin actions on proposals
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go
│   │   │       ├── transaction.go
│   │   │       ├── migrations.go          # Auto-migrate logic
│   │   │       ├── proposal_repository.go
│   │   │       ├── bid_repository.go
│   │   │       ├── bid_strategy_repository.go
│   │   │       ├── bid_notification_repository.go
│   │   │       ├── milestone_repository.go
│   │   │       ├── cover_letter_repository.go
│   │   │       ├── attachment_repository.go
│   │   │       ├── question_answer_repository.go
│   │   │       ├── template_repository.go
│   │   │       ├── connect_repository.go
│   │   │       ├── boost_repository.go
│   │   │       ├── analytics_repository.go
│   │   │       ├── proposal_flag_repository.go
│   │   │       └── outbox_repository.go
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go
│   │   │       ├── proposal_cache.go
│   │   │       └── bid_cache.go
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go
│   │   │       ├── producer.go
│   │   │       ├── topics.go
│   │   │       └── scram.go
│   │   ├── storage/
│   │   │   └── client.go
│   │   └── outbox/
│   │       ├── processor.go
│   │       └── scheduler.go
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── proposal_handler.go
│   │       │   ├── bid_handler.go
│   │       │   ├── bid_strategy_handler.go
│   │       │   ├── milestone_handler.go
│   │       │   ├── template_handler.go
│   │       │   ├── connect_handler.go
│   │       │   ├── boost_handler.go
│   │       │   ├── analytics_handler.go
│   │       │   ├── flag_handler.go
│   │       │   └── health_handler.go
│   │       ├── middleware/
│   │       │   ├── auth.go
│   │       │   ├── rbac.go
│   │       │   ├── cors.go
│   │       │   ├── rate_limit.go
│   │       │   ├── logging.go
│   │       │   ├── error_handler.go
│   │       │   └── request_id.go
│   │       ├── responses/
│   │       │   ├── success.go
│   │       │   ├── error.go
│   │       │   └── pagination.go
│   │       └── router.go
│   │
│   └── config/
│       ├── config.go
│       ├── database.go
│       ├── kafka.go
│       ├── redis.go
│       └── services.go
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   └── codes.go
│   ├── logger/
│   │   └── logger.go
│   ├── utils/
│   │   ├── validator.go
│   │   ├── sanitizer.go
│   │   └── bid_calculator.go
│   └── constants/
│       ├── events.go
│       ├── topics.go
│       └── bid_types.go
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   └── seed-data.sh
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── api.md
│   ├── events.md
│   ├── bidding-system.md
│   ├── connects-system.md
│   └── auto-bidding.md
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md
```

---

## 4️⃣ contracts-be (UPDATED)

```
contracts-be/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── domain/
│   │   ├── contract/
│   │   │   ├── entity.go              # Contract aggregate root
│   │   │   ├── terms.go               # Contract terms
│   │   │   ├── enums.go               # Type, Status, PaymentType
│   │   │   ├── value_objects.go       # Amount, Duration
│   │   │   └── repository.go
│   │   ├── milestone/
│   │   │   ├── entity.go              # Contract milestones
│   │   │   ├── delivery.go            # Milestone deliverables
│   │   │   ├── status.go              # Milestone status tracking
│   │   │   └── repository.go
│   │   ├── timesheet/
│   │   │   ├── entity.go              # Time entries
│   │   │   ├── week.go                # Weekly timesheets
│   │   │   ├── time_entry.go          # Individual time entry
│   │   │   ├── screenshot.go          # Work diary screenshots
│   │   │   └── repository.go
│   │   ├── workdiary/
│   │   │   ├── entity.go              # Work diary entries
│   │   │   ├── activity_level.go      # Activity tracking
│   │   │   └── repository.go
│   │   ├── template/
│   │   │   ├── entity.go              # Contract templates
│   │   │   ├── clause.go              # Contract clauses
│   │   │   └── repository.go
│   │   ├── amendment/
│   │   │   ├── entity.go              # Contract amendments
│   │   │   ├── change_request.go      # Change requests
│   │   │   └── repository.go
│   │   ├── deliverable/
│   │   │   ├── entity.go              # Contract deliverables
│   │   │   ├── submission.go          # Deliverable submissions
│   │   │   ├── revision.go            # Revision requests
│   │   │   └── repository.go
│   │   ├── pause/
│   │   │   ├── entity.go              # Contract pause/resume
│   │   │   └── repository.go
│   │   ├── termination/
│   │   │   ├── entity.go              # Contract termination
│   │   │   ├── reason.go              # Termination reasons
│   │   │   └── repository.go
│   │   ├── dispute/
│   │   │   ├── entity.go              # Contract disputes
│   │   │   ├── evidence.go            # Evidence submission
│   │   │   ├── resolution.go          # Dispute resolution
│   │   │   ├── status.go              # Dispute status
│   │   │   └── repository.go
│   │   └── outbox/
│   │       ├── entity.go
│   │       └── repository.go
│   │
│   ├── application/
│   │   ├── contract/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Update, Pause, End
│   │   │   ├── queries.go             # Get, List, Filter
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── business_rules.go
│   │   ├── milestone/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Complete, Request
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── timesheet/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Log time, Submit week
│   │   │   ├── queries.go             # Get timesheets, Calculate
│   │   │   ├── calculator.go          # Calculate hours, earnings
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── workdiary/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── template/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── amendment/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── deliverable/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── termination/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── dispute/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Open, Submit evidence, Resolve
│   │   │   ├── queries.go             # Get disputes
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   └── eventhandler/
│   │       ├── proposal_handler.go
│   │       ├── user_handler.go
│   │       ├── payment_handler.go
│   │       ├── escrow_handler.go
│   │       ├── dispute_handler.go
│   │       └── admin_handler.go       # Handle admin dispute resolutions
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go
│   │   │       ├── transaction.go
│   │   │       ├── migrations.go          # Auto-migrate logic
│   │   │       ├── contract_repository.go
│   │   │       ├── milestone_repository.go
│   │   │       ├── timesheet_repository.go
│   │   │       ├── workdiary_repository.go
│   │   │       ├── template_repository.go
│   │   │       ├── amendment_repository.go
│   │   │       ├── deliverable_repository.go
│   │   │       ├── pause_repository.go
│   │   │       ├── termination_repository.go
│   │   │       ├── dispute_repository.go
│   │   │       └── outbox_repository.go
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go
│   │   │       └── contract_cache.go
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go
│   │   │       ├── producer.go
│   │   │       ├── topics.go
│   │   │       └── scram.go
│   │   ├── storage/
│   │   │   └── client.go
│   │   └── outbox/
│   │       ├── processor.go
│   │       └── scheduler.go
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── contract_handler.go
│   │       │   ├── milestone_handler.go
│   │       │   ├── timesheet_handler.go
│   │       │   ├── workdiary_handler.go
│   │       │   ├── template_handler.go
│   │       │   ├── amendment_handler.go
│   │       │   ├── deliverable_handler.go
│   │       │   ├── termination_handler.go
│   │       │   ├── dispute_handler.go
│   │       │   └── health_handler.go
│   │       ├── middleware/
│   │       │   ├── auth.go
│   │       │   ├── rbac.go
│   │       │   ├── cors.go
│   │       │   ├── rate_limit.go
│   │       │   ├── logging.go
│   │       │   ├── error_handler.go
│   │       │   └── request_id.go
│   │       ├── responses/
│   │       │   ├── success.go
│   │       │   ├── error.go
│   │       │   └── pagination.go
│   │       └── router.go
│   │
│   └── config/
│       ├── config.go
│       ├── database.go
│       ├── kafka.go
│       ├── redis.go
│       └── services.go
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   └── codes.go
│   ├── logger/
│   │   └── logger.go
│   ├── utils/
│   │   ├── validator.go
│   │   ├── time_calculator.go
│   │   └── date_utils.go
│   └── constants/
│       ├── events.go
│       └── topics.go
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   └── seed-data.sh
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── api.md
│   ├── events.md
│   ├── contract-lifecycle.md
│   └── timesheet-system.md
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md
```

---

## 5️⃣ financial-be (UPDATED)

```
financial-be/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── domain/
│   │   ├── wallet/
│   │   │   ├── entity.go              # User wallet
│   │   │   ├── balance.go             # Balance tracking
│   │   │   ├── currency.go            # Multi-currency support
│   │   │   └── repository.go
│   │   ├── transaction/
│   │   │   ├── entity.go              # Financial transactions
│   │   │   ├── enums.go               # Type, Status, Category
│   │   │   ├── ledger.go              # Double-entry ledger
│   │   │   └── repository.go
│   │   ├── payment/
│   │   │   ├── entity.go              # Payment records
│   │   │   ├── payment_method.go      # Credit card, PayPal, Bank transfer
│   │   │   ├── gateway.go             # Payment gateway abstraction
│   │   │   └── repository.go
│   │   ├── escrow/
│   │   │   ├── entity.go              # Escrow accounts
│   │   │   ├── hold.go                # Fund holds
│   │   │   ├── release.go             # Fund release conditions
│   │   │   └── repository.go
│   │   ├── payout/
│   │   │   ├── entity.go              # Payout requests
│   │   │   ├── method.go              # Bank transfer, PayPal, Payoneer
│   │   │   ├── schedule.go            # Payout scheduling
│   │   │   └── repository.go
│   │   ├── invoice/
│   │   │   ├── entity.go              # Invoice generation
│   │   │   ├── line_item.go           # Invoice line items
│   │   │   ├── tax.go                 # Tax calculations
│   │   │   └── repository.go
│   │   ├── fee/
│   │   │   ├── entity.go              # Platform fees
│   │   │   ├── calculator.go          # Fee calculation rules
│   │   │   ├── tier.go                # Fee tiers
│   │   │   └── repository.go
│   │   ├── refund/
│   │   │   ├── entity.go              # Refund processing
│   │   │   ├── policy.go              # Refund policies
│   │   │   └── repository.go
│   │   ├── dispute_payment/
│   │   │   ├── entity.go              # Payment disputes
│   │   │   ├── chargeback.go          # Chargeback handling
│   │   │   └── repository.go
│   │   ├── tax/
│   │   │   ├── entity.go              # Tax records
│   │   │   ├── form.go                # Tax forms (W9, 1099)
│   │   │   └── repository.go
│   │   ├── bank_account/
│   │   │   ├── entity.go              # Bank account info
│   │   │   ├── verification.go        # Micro-deposit verification
│   │   │   └── repository.go
│   │   ├── payment_schedule/
│   │   │   ├── entity.go              # Recurring payment schedules
│   │   │   └── repository.go
│   │   └── outbox/
│   │       ├── entity.go
│   │       └── repository.go
│   │
│   ├── application/
│   │   ├── wallet/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Deposit, Withdraw, Transfer
│   │   │   ├── queries.go             # Get balance, History
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── transaction/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Reverse, Reconcile
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── reconciliation.go
│   │   ├── payment/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Process, Capture, Void
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── stripe_processor.go
│   │   │   ├── paypal_processor.go
│   │   │   └── processor_factory.go
│   │   ├── escrow/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Hold, Release, Refund
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── release_manager.go
│   │   ├── payout/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Request, Process, Cancel
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── batch_processor.go
│   │   ├── invoice/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Generate, Send, Mark paid
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── generator.go           # PDF generation
│   │   │   └── tax_calculator.go
│   │   ├── fee/
│   │   │   ├── service.go
│   │   │   ├── calculator.go
│   │   │   ├── dto.go
│   │   │   └── rules_engine.go
│   │   ├── refund/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── tax/
│   │   │   ├── service.go
│   │   │   ├── form_generator.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── bank_account/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Add, Verify, Remove
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   └── eventhandler/
│   │       ├── contract_handler.go
│   │       ├── milestone_handler.go
│   │       ├── dispute_handler.go
│   │       ├── user_handler.go
│   │       ├── subscription_handler.go
│   │       └── admin_handler.go       # Handle admin actions on payments
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go
│   │   │       ├── transaction.go
│   │   │       ├── migrations.go          # Auto-migrate logic
│   │   │       ├── wallet_repository.go
│   │   │       ├── transaction_repository.go
│   │   │       ├── payment_repository.go
│   │   │       ├── escrow_repository.go
│   │   │       ├── payout_repository.go
│   │   │       ├── invoice_repository.go
│   │   │       ├── fee_repository.go
│   │   │       ├── refund_repository.go
│   │   │       ├── dispute_payment_repository.go
│   │   │       ├── tax_repository.go
│   │   │       ├── bank_account_repository.go
│   │   │       ├── payment_schedule_repository.go
│   │   │       └── outbox_repository.go
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go
│   │   │       ├── wallet_cache.go
│   │   │       └── rate_cache.go
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go
│   │   │       ├── producer.go
│   │   │       ├── topics.go
│   │   │       └── scram.go
│   │   ├── payment_gateway/
│   │   │   ├── stripe/
│   │   │   │   ├── client.go
│   │   │   │   ├── webhook_handler.go
│   │   │   │   └── mapper.go
│   │   │   ├── paypal/
│   │   │   │   ├── client.go
│   │   │   │   ├── webhook_handler.go
│   │   │   │   └── mapper.go
│   │   │   └── factory.go
│   │   ├── pdf/
│   │   │   └── generator.go           # PDF invoice generation
│   │   └── outbox/
│   │       ├── processor.go
│   │       └── scheduler.go
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── wallet_handler.go
│   │       │   ├── transaction_handler.go
│   │       │   ├── payment_handler.go
│   │       │   ├── escrow_handler.go
│   │       │   ├── payout_handler.go
│   │       │   ├── invoice_handler.go
│   │       │   ├── fee_handler.go
│   │       │   ├── refund_handler.go
│   │       │   ├── tax_handler.go
│   │       │   ├── bank_account_handler.go
│   │       │   ├── webhook_handler.go
│   │       │   └── health_handler.go
│   │       ├── middleware/
│   │       │   ├── auth.go
│   │       │   ├── rbac.go
│   │       │   ├── cors.go
│   │       │   ├── rate_limit.go
│   │       │   ├── logging.go
│   │       │   ├── error_handler.go
│   │       │   ├── request_id.go
│   │       │   └── idempotency.go
│   │       ├── responses/
│   │       │   ├── success.go
│   │       │   ├── error.go
│   │       │   └── pagination.go
│   │       └── router.go
│   │
│   └── config/
│       ├── config.go
│       ├── database.go
│       ├── kafka.go
│       ├── redis.go
│       ├── stripe.go
│       ├── paypal.go
│       └── services.go
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   ├── codes.go
│   │   └── payment_errors.go
│   ├── logger/
│   │   └── logger.go
│   ├── utils/
│   │   ├── validator.go
│   │   ├── currency.go
│   │   ├── decimal.go
│   │   └── encryption.go
│   └── constants/
│       ├── events.go
│       ├── topics.go
│       ├── currencies.go
│       └── payment_methods.go
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       ├── networkpolicy.yaml
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   ├── seed-data.sh
│   └── reconciliation.sh
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── api.md
│   ├── events.md
│   ├── payment-flows.md
│   ├── escrow-system.md
│   ├── fee-structure.md
│   └── compliance.md
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md
```

---

## 6️⃣ communications-be (UPDATED - FREE & SELF-HOSTED)

```
communications-be/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── domain/
│   │   ├── conversation/
│   │   │   ├── entity.go              # Chat conversations
│   │   │   ├── participant.go         # Conversation participants
│   │   │   ├── settings.go            # Conversation settings
│   │   │   ├── typing_indicator.go    # Typing indicators
│   │   │   └── repository.go
│   │   ├── message/
│   │   │   ├── entity.go              # Chat messages
│   │   │   ├── attachment.go          # Message attachments
│   │   │   ├── read_receipt.go        # Read receipts
│   │   │   ├── reaction.go            # Message reactions
│   │   │   ├── mention.go             # User mentions
│   │   │   └── repository.go
│   │   ├── notification/
│   │   │   ├── entity.go              # All notifications
│   │   │   ├── enums.go               # Type, Priority, Category
│   │   │   ├── preferences.go         # User notification preferences
│   │   │   ├── settings.go            # Notification settings per type
│   │   │   └── repository.go
│   │   ├── in_app_notification/
│   │   │   ├── entity.go              # In-app real-time notifications
│   │   │   ├── badge_count.go         # Unread notification badges
│   │   │   ├── group.go               # Notification grouping
│   │   │   ├── action.go              # Notification actions (CTA)
│   │   │   └── repository.go
│   │   ├── notification_template/
│   │   │   ├── entity.go              # Notification templates
│   │   │   ├── variable.go            # Template variables
│   │   │   ├── localization.go        # Multi-language templates
│   │   │   └── repository.go
│   │   ├── email/
│   │   │   ├── entity.go              # Email records
│   │   │   ├── template.go            # Email templates
│   │   │   ├── batch.go               # Batch email sending
│   │   │   └── repository.go
│   │   ├── notification_queue/
│   │   │   ├── entity.go              # Queued notifications
│   │   │   ├── priority_queue.go      # Priority-based queue
│   │   │   └── repository.go
│   │   ├── delivery_log/
│   │   │   ├── entity.go              # Delivery tracking
│   │   │   ├── status.go              # Delivery status
│   │   │   └── repository.go
│   │   ├── unsubscribe/
│   │   │   ├── entity.go              # Unsubscribe management
│   │   │   └── repository.go
│   │   ├── online_status/
│   │   │   ├── entity.go              # User online/offline status
│   │   │   └── repository.go
│   │   ├── message_flag/
│   │   │   ├── entity.go              # Flagged messages
│   │   │   ├── reason.go              # Flag reasons
│   │   │   ├── status.go              # Flag status
│   │   │   └── repository.go
│   │   └── outbox/
│   │       ├── entity.go
│   │       └── repository.go
│   │
│   ├── application/
│   │   ├── conversation/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Archive, Mute, Delete
│   │   │   ├── queries.go             # Get, List conversations
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── message/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Send, Edit, Delete, React
│   │   │   ├── queries.go             # Get messages, Search
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── realtime_service.go    # WebSocket handling
│   │   ├── notification/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Send, Mark read, Delete, Clear all
│   │   │   ├── queries.go             # Get notifications, Count unread
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── orchestrator.go        # Multi-channel orchestration
│   │   │   ├── preferences_service.go # Manage preferences
│   │   │   └── aggregator.go          # Aggregate notifications
│   │   ├── in_app_notification/
│   │   │   ├── service.go
│   │   │   ├── real_time_sender.go    # Real-time notification delivery
│   │   │   ├── badge_manager.go       # Badge count management
│   │   │   ├── grouping_engine.go     # Group similar notifications
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── email/
│   │   │   ├── service.go
│   │   │   ├── sender.go              # Email sending logic
│   │   │   ├── template_renderer.go   # Render email templates
│   │   │   ├── batch_sender.go        # Batch email sending
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── wildduck_client.go     # WildDuck SMTP integration
│   │   ├── template/
│   │   │   ├── service.go
│   │   │   ├── renderer.go            # Template rendering engine
│   │   │   ├── variable_injector.go   # Inject dynamic variables
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── online_status/
│   │   │   ├── service.go
│   │   │   ├── tracker.go             # Track online/offline
│   │   │   ├── presence_manager.go    # Manage user presence
│   │   │   └── dto.go
│   │   ├── flag/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Flag, Unflag message
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   └── eventhandler/
│   │       ├── user_handler.go        # User events (registration, profile updates)
│   │       ├── job_handler.go         # Job events (new job, job closed)
│   │       ├── proposal_handler.go    # Proposal events (new proposal, accepted)
│   │       ├── bid_handler.go         # Bid events (new bid, outbid alert)
│   │       ├── contract_handler.go    # Contract events (created, milestone completed)
│   │       ├── payment_handler.go     # Payment notifications
│   │       ├── review_handler.go      # Review notifications
│   │       ├── message_handler.go     # New message notifications
│   │       ├── subscription_handler.go # Subscription notifications
│   │       ├── system_handler.go      # System notifications
│   │       └── admin_handler.go       # Handle admin actions on messages
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go
│   │   │       ├── transaction.go
│   │   │       ├── migrations.go          # Auto-migrate logic
│   │   │       ├── conversation_repository.go
│   │   │       ├── message_repository.go
│   │   │       ├── notification_repository.go
│   │   │       ├── in_app_notification_repository.go
│   │   │       ├── template_repository.go
│   │   │       ├── email_repository.go
│   │   │       ├── notification_queue_repository.go
│   │   │       ├── delivery_log_repository.go
│   │   │       ├── unsubscribe_repository.go
│   │   │       ├── online_status_repository.go
│   │   │       ├── message_flag_repository.go
│   │   │       └── outbox_repository.go
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go
│   │   │       ├── conversation_cache.go
│   │   │       ├── online_status_cache.go # Cache online status
│   │   │       ├── typing_indicator_cache.go # Cache typing indicators
│   │   │       ├── notification_cache.go  # Cache unread counts
│   │   │       └── presence_cache.go
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go
│   │   │       ├── producer.go
│   │   │       ├── topics.go
│   │   │       └── scram.go
│   │   ├── realtime/
│   │   │   ├── websocket/
│   │   │   │   ├── hub.go             # WebSocket hub (connection manager)
│   │   │   │   ├── client.go          # WebSocket client
│   │   │   │   ├── handler.go         # WebSocket message handler
│   │   │   │   ├── broadcaster.go     # Broadcast to multiple clients
│   │   │   │   └── room.go            # WebSocket rooms/channels
│   │   │   └── sse/
│   │   │       ├── handler.go         # Server-Sent Events
│   │   │       └── stream.go          # SSE stream management
│   │   ├── email/
│   │   │   ├── wildduck/
│   │   │   │   ├── client.go          # WildDuck SMTP/API client
│   │   │   │   ├── smtp_sender.go     # SMTP sending
│   │   │   │   ├── api_client.go      # WildDuck REST API
│   │   │   │   └── config.go
│   │   │   └── smtp/
│   │   │       ├── client.go          # Generic SMTP fallback
│   │   │       └── config.go
│   │   ├── storage/
│   │   │   └── client.go              # Storage service client
│   │   └── outbox/
│   │       ├── processor.go
│   │       └── scheduler.go
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── conversation_handler.go
│   │       │   ├── message_handler.go
│   │       │   ├── notification_handler.go
│   │       │   ├── in_app_notification_handler.go
│   │       │   ├── email_handler.go
│   │       │   ├── template_handler.go
│   │       │   ├── preferences_handler.go
│   │       │   ├── websocket_handler.go
│   │       │   ├── sse_handler.go
│   │       │   ├── online_status_handler.go
│   │       │   ├── unsubscribe_handler.go
│   │       │   ├── flag_handler.go
│   │       │   └── health_handler.go
│   │       ├── middleware/
│   │       │   ├── auth.go
│   │       │   ├── rbac.go
│   │       │   ├── cors.go
│   │       │   ├── rate_limit.go
│   │       │   ├── logging.go
│   │       │   ├── error_handler.go
│   │       │   └── request_id.go
│   │       ├── responses/
│   │       │   ├── success.go
│   │       │   ├── error.go
│   │       │   └── pagination.go
│   │       └── router.go
│   │
│   └── config/
│       ├── config.go
│       ├── database.go
│       ├── kafka.go
│       ├── redis.go
│       ├── wildduck.go
│       ├── websocket.go
│       └── services.go
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   └── codes.go
│   ├── logger/
│   │   └── logger.go
│   ├── utils/
│   │   ├── validator.go
│   │   ├── template_engine.go
│   │   ├── sanitizer.go
│   │   └── html_to_text.go
│   └── constants/
│       ├── events.go
│       ├── topics.go
│       ├── notification_types.go
│       └── websocket_events.go
│
├── templates/
│   ├── email/
│   │   ├── base.html                  # Base email template
│   │   ├── welcome.html
│   │   ├── job_alert.html
│   │   ├── new_proposal.html
│   │   ├── proposal_accepted.html
│   │   ├── bid_received.html
│   │   ├── outbid_alert.html
│   │   ├── contract_created.html
│   │   ├── milestone_completed.html
│   │   ├── payment_received.html
│   │   ├── payment_sent.html
│   │   ├── review_request.html
│   │   ├── review_received.html
│   │   ├── new_message.html
│   │   ├── subscription_expiring.html
│   │   ├── password_reset.html
│   │   ├── verify_email.html
│   │   └── weekly_summary.html
│   └── notification/
│       ├── job_posted.json
│       ├── proposal_received.json
│       ├── contract_created.json
│       ├── milestone_completed.json
│       ├── payment_received.json
│       └── review_received.json
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   ├── seed-templates.sh
│   └── seed-data.sh
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── api.md
│   ├── events.md
│   ├── websocket-protocol.md
│   ├── notification-system.md
│   ├── in-app-notifications.md
│   ├── email-templates.md
│   └── wildduck-integration.md
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md
```

---

## 7️⃣ storage-be (UPDATED)

```
storage-be/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── domain/
│   │   ├── file/
│   │   │   ├── entity.go              # File metadata
│   │   │   ├── enums.go               # FileType, Status, Visibility
│   │   │   ├── metadata.go            # File metadata
│   │   │   └── repository.go
│   │   ├── folder/
│   │   │   ├── entity.go              # Folder structure
│   │   │   └── repository.go
│   │   ├── upload/
│   │   │   ├── entity.go              # Upload sessions
│   │   │   ├── chunk.go               # Chunked uploads
│   │   │   ├── resumable.go           # Resumable upload support
│   │   │   └── repository.go
│   │   ├── media/
│   │   │   ├── entity.go              # Media processing records
│   │   │   ├── thumbnail.go           # Thumbnail generation
│   │   │   ├── variant.go             # Image variants
│   │   │   └── repository.go
│   │   ├── access_control/
│   │   │   ├── entity.go              # File access permissions
│   │   │   └── repository.go
│   │   ├── version/
│   │   │   ├── entity.go              # File versioning
│   │   │   └── repository.go
│   │   ├── share/
│   │   │   ├── entity.go              # File sharing
│   │   │   ├── link.go                # Share links
│   │   │   └── repository.go
│   │   ├── file_flag/
│   │   │   ├── entity.go              # Flagged files
│   │   │   ├── reason.go              # Flag reasons
│   │   │   ├── status.go              # Flag status
│   │   │   └── repository.go
│   │   └── outbox/
│   │       ├── entity.go
│   │       └── repository.go
│   │
│   ├── application/
│   │   ├── file/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Upload, Delete, Move, Copy
│   │   │   ├── queries.go             # Get, List, Search
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── upload/
│   │   │   ├── service.go
│   │   │   ├── chunked_upload.go
│   │   │   ├── resumable.go
│   │   │   └── dto.go
│   │   ├── media/
│   │   │   ├── service.go
│   │   │   ├── image_processor.go
│   │   │   ├── video_processor.go
│   │   │   ├── thumbnail_generator.go
│   │   │   └── dto.go
│   │   ├── folder/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── share/
│   │   │   ├── service.go
│   │   │   ├── link_generator.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── version/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── flag/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Flag, Unflag file
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   └── eventhandler/
│   │       ├── user_handler.go
│   │       ├── contract_handler.go
│   │       ├── portfolio_handler.go
│   │       └── admin_handler.go       # Handle admin actions on files
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go
│   │   │       ├── transaction.go
│   │   │       ├── migrations.go          # Auto-migrate logic
│   │   │       ├── file_repository.go
│   │   │       ├── folder_repository.go
│   │   │       ├── upload_repository.go
│   │   │       ├── media_repository.go
│   │   │       ├── access_control_repository.go
│   │   │       ├── version_repository.go
│   │   │       ├── share_repository.go
│   │   │       ├── file_flag_repository.go
│   │   │       └── outbox_repository.go
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go
│   │   │       └── file_cache.go
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go
│   │   │       ├── producer.go
│   │   │       ├── topics.go
│   │   │       └── scram.go
│   │   ├── object_storage/
│   │   │   ├── local/
│   │   │   │   ├── storage.go         # Local file system storage
│   │   │   │   └── config.go
│   │   │   ├── minio/
│   │   │   │   ├── client.go          # Self-hosted MinIO
│   │   │   │   └── config.go
│   │   │   └── provider.go
│   │   ├── media_processing/
│   │   │   ├── image/
│   │   │   │   ├── resizer.go
│   │   │   │   ├── optimizer.go
│   │   │   │   └── watermark.go
│   │   │   ├── video/
│   │   │   │   ├── transcoder.go
│   │   │   │   └── thumbnail.go
│   │   │   └── processor.go
│   │   ├── virus_scan/
│   │   │   └── clamav.go              # ClamAV integration
│   │   └── outbox/
│   │       ├── processor.go
│   │       └── scheduler.go
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── file_handler.go
│   │       │   ├── upload_handler.go
│   │       │   ├── download_handler.go
│   │       │   ├── folder_handler.go
│   │       │   ├── share_handler.go
│   │       │   ├── media_handler.go
│   │       │   ├── flag_handler.go
│   │       │   └── health_handler.go
│   │       ├── middleware/
│   │       │   ├── auth.go
│   │       │   ├── rbac.go
│   │       │   ├── cors.go
│   │       │   ├── rate_limit.go
│   │       │   ├── logging.go
│   │       │   ├── error_handler.go
│   │       │   ├── request_id.go
│   │       │   └── file_size_limit.go
│   │       ├── responses/
│   │       │   ├── success.go
│   │       │   ├── error.go
│   │       │   └── pagination.go
│   │       └── router.go
│   │
│   └── config/
│       ├── config.go
│       ├── database.go
│       ├── kafka.go
│       ├── redis.go
│       ├── minio.go
│       └── storage.go
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   └── codes.go
│   ├── logger/
│   │   └── logger.go
│   ├── utils/
│   │   ├── validator.go
│   │   ├── file_utils.go
│   │   ├── mime_detector.go
│   │   └── hash.go
│   └── constants/
│       ├── events.go
│       ├── topics.go
│       └── mime_types.go
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       ├── pvc.yaml                    # Persistent Volume Claim
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   ├── seed-data.sh
│   └── cleanup-orphaned.sh
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── api.md
│   ├── events.md
│   ├── upload-flow.md
│   └── media-processing.md
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md
```

---

## 8️⃣ search-be (UPDATED WITH ENHANCED RECOMMENDATIONS)

```
search-be/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── domain/
│   │   ├── search_index/
│   │   │   ├── entity.go              # Search index metadata
│   │   │   ├── job_index.go           # Job search document
│   │   │   ├── user_index.go          # User/Freelancer document
│   │   │   └── repository.go
│   │   ├── search_query/
│   │   │   ├── entity.go              # Search query logs
│   │   │   ├── filters.go             # Search filters
│   │   │   └── repository.go
│   │   ├── recommendation/
│   │   │   ├── entity.go              # Recommendation records
│   │   │   ├── score.go               # Scoring algorithm
│   │   │   ├── reason.go              # Recommendation reasons
│   │   │   ├── feedback.go            # User feedback on recommendations
│   │   │   └── repository.go
│   │   ├── recommendation_model/
│   │   │   ├── entity.go              # ML model metadata
│   │   │   ├── feature.go             # Feature vectors
│   │   │   ├── training_data.go       # Training data
│   │   │   └── repository.go
│   │   ├── user_preference/
│   │   │   ├── entity.go              # User preferences for recommendations
│   │   │   ├── implicit_signals.go    # Implicit user signals (clicks, views)
│   │   │   ├── explicit_preferences.go # Explicit preferences
│   │   │   └── repository.go
│   │   ├── matching/
│   │   │   ├── entity.go              # Job-Freelancer matches
│   │   │   ├── criteria.go            # Matching criteria
│   │   │   ├── score_breakdown.go     # Detailed match scoring
│   │   │   └── repository.go
│   │   ├── feed/
│   │   │   ├── entity.go              # User feeds
│   │   │   ├── item.go                # Feed items
│   │   │   ├── personalization.go     # Personalization data
│   │   │   └── repository.go
│   │   ├── trending/
│   │   │   ├── entity.go              # Trending jobs/skills
│   │   │   ├── calculator.go          # Calculate trending items
│   │   │   └── repository.go
│   │   ├── saved_search/
│   │   │   ├── entity.go              # Saved searches
│   │   │   ├── alert.go               # Search alerts
│   │   │   └── repository.go
│   │   ├── similarity/
│   │   │   ├── entity.go              # Similar jobs/users
│   │   │   ├── vector.go              # Similarity vectors
│   │   │   └── repository.go
│   │   └── outbox/
│   │       ├── entity.go
│   │       └── repository.go
│   │
│   ├── application/
│   │   ├── search/
│   │   │   ├── service.go
│   │   │   ├── job_search.go
│   │   │   ├── freelancer_search.go
│   │   │   ├── query_builder.go
│   │   │   ├── facet_builder.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── indexing/
│   │   │   ├── service.go
│   │   │   ├── job_indexer.go
│   │   │   ├── user_indexer.go
│   │   │   ├── bulk_indexer.go
│   │   │   └── dto.go
│   │   ├── recommendation/
│   │   │   ├── service.go
│   │   │   ├── job_recommender.go     # Recommend jobs to freelancers
│   │   │   ├── freelancer_recommender.go # Recommend freelancers to clients
│   │   │   ├── collaborative_filtering.go # Collaborative filtering algorithm
│   │   │   ├── content_based.go       # Content-based filtering
│   │   │   ├── hybrid_recommender.go  # Hybrid recommendation approach
│   │   │   ├── scoring_engine.go      # Calculate recommendation scores
│   │   │   ├── personalization.go     # Personalization logic
│   │   │   ├── diversity_optimizer.go # Ensure diverse recommendations
│   │   │   ├── cold_start_handler.go  # Handle new users/jobs
│   │   │   ├── dto.go
│   │   │   └── ml_model.go            # ML model integration
│   │   ├── matching/
│   │   │   ├── service.go
│   │   │   ├── matcher.go             # Job-Freelancer matching
│   │   │   ├── criteria_evaluator.go  # Evaluate match criteria
│   │   │   ├── skill_matcher.go       # Match based on skills
│   │   │   ├── experience_matcher.go  # Match based on experience
│   │   │   ├── rate_matcher.go        # Match based on rates
│   │   │   ├── availability_matcher.go # Match based on availability
│   │   │   ├── score_calculator.go    # Calculate match score
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── feed/
│   │   │   ├── service.go
│   │   │   ├── generator.go           # Generate personalized feed
│   │   │   ├── ranking.go             # Rank feed items
│   │   │   ├── freshness_scorer.go    # Score by freshness
│   │   │   ├── relevance_scorer.go    # Score by relevance
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── trending/
│   │   │   ├── service.go
│   │   │   ├── calculator.go          # Calculate trending items
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── similarity/
│   │   │   ├── service.go
│   │   │   ├── job_similarity.go      # Find similar jobs
│   │   │   ├── user_similarity.go     # Find similar users
│   │   │   ├── vector_calculator.go   # Calculate similarity vectors
│   │   │   └── dto.go
│   │   ├── suggestion/
│   │   │   ├── service.go             # Autocomplete suggestions
│   │   │   ├── dto.go
│   │   │   └── cache_warmer.go
│   │   └── eventhandler/
│   │       ├── job_handler.go
│   │       ├── user_handler.go
│   │       ├── proposal_handler.go
│   │       ├── contract_handler.go
│   │       ├── review_handler.go
│   │       ├── skill_handler.go
│   │       ├── interaction_handler.go # Track user interactions
│   │       └── admin_handler.go       # Handle content removal from index
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go
│   │   │       ├── transaction.go
│   │   │       ├── migrations.go          # Auto-migrate logic
│   │   │       ├── search_query_repository.go
│   │   │       ├── recommendation_repository.go
│   │   │       ├── recommendation_model_repository.go
│   │   │       ├── user_preference_repository.go
│   │   │       ├── matching_repository.go
│   │   │       ├── feed_repository.go
│   │   │       ├── trending_repository.go
│   │   │       ├── saved_search_repository.go
│   │   │       ├── similarity_repository.go
│   │   │       └── outbox_repository.go
│   │   ├── elasticsearch/
│   │   │   ├── client.go
│   │   │   ├── index_manager.go
│   │   │   ├── job_mapper.go
│   │   │   ├── user_mapper.go
│   │   │   └── config.go
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go
│   │   │       ├── search_cache.go
│   │   │       ├── suggestion_cache.go
│   │   │       ├── feed_cache.go
│   │   │       ├── recommendation_cache.go
│   │   │       └── trending_cache.go
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go
│   │   │       ├── producer.go
│   │   │       ├── topics.go
│   │   │       └── scram.go
│   │   ├── ml/
│   │   │   ├── model_loader.go        # Load ML models
│   │   │   ├── predictor.go           # Make predictions
│   │   │   ├── feature_extractor.go   # Extract features
│   │   │   ├── trainer.go             # Train models
│   │   │   └── evaluator.go           # Evaluate model performance
│   │   └── outbox/
│   │       ├── processor.go
│   │       └── scheduler.go
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── search_handler.go
│   │       │   ├── recommendation_handler.go
│   │       │   ├── matching_handler.go
│   │       │   ├── feed_handler.go
│   │       │   ├── trending_handler.go
│   │       │   ├── similarity_handler.go
│   │       │   ├── suggestion_handler.go
│   │       │   ├── indexing_handler.go # Admin endpoint
│   │       │   └── health_handler.go
│   │       ├── middleware/
│   │       │   ├── auth.go
│   │       │   ├── rbac.go
│   │       │   ├── cors.go
│   │       │   ├── rate_limit.go
│   │       │   ├── logging.go
│   │       │   ├── error_handler.go
│   │       │   └── request_id.go
│   │       ├── responses/
│   │       │   ├── success.go
│   │       │   ├── error.go
│   │       │   └── pagination.go
│   │       └── router.go
│   │
│   └── config/
│       ├── config.go
│       ├── database.go
│       ├── kafka.go
│       ├── redis.go
│       ├── elasticsearch.go
│       └── ml_config.go
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   └── codes.go
│   ├── logger/
│   │   └── logger.go
│   ├── utils/
│   │   ├── validator.go
│   │   ├── text_analyzer.go
│   │   ├── normalizer.go
│   │   └── vector_math.go
│   └── constants/
│       ├── events.go
│       ├── topics.go
│       └── indices.go
│
├── elasticsearch/
│   ├── mappings/
│   │   ├── jobs.json              # Job index mapping
│   │   └── users.json             # User index mapping
│   └── analyzers/
│       └── custom_analyzers.json
│
├── ml_models/
│   ├── job_recommendation/
│   │   ├── model.pkl
│   │   ├── features.json
│   │   └── metadata.json
│   ├── freelancer_recommendation/
│   │   ├── model.pkl
│   │   ├── features.json
│   │   └── metadata.json
│   └── matching/
│       ├── model.pkl
│       ├── features.json
│       └── metadata.json
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   ├── seed-data.sh
│   ├── create-indices.sh
│   ├── reindex-all.sh
│   └── train-models.sh
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── api.md
│   ├── events.md
│   ├── search-algorithms.md
│   ├── recommendation-engine.md
│   ├── recommendation-types.md
│   ├── matching-algorithm.md
│   ├── ml-models.md
│   └── elasticsearch-setup.md
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md
```

---

## 9️⃣ reviews-be (UPDATED - COVERS RATINGS)

```
reviews-be/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── domain/
│   │   ├── review/
│   │   │   ├── entity.go              # Review aggregate root
│   │   │   ├── rating.go              # Rating breakdown (overall, quality, communication, etc.)
│   │   │   ├── enums.go               # ReviewType, Status, ReviewCategory
│   │   │   ├── response.go            # Review responses
│   │   │   ├── helpful.go             # Helpful votes
│   │   │   └── repository.go
│   │   ├── rating/
│   │   │   ├── entity.go              # Rating system
│   │   │   ├── criteria.go            # Rating criteria
│   │   │   ├── aggregation.go         # Aggregate ratings
│   │   │   └── repository.go
│   │   ├── badge/
│   │   │   ├── entity.go              # Achievement badges
│   │   │   ├── criteria.go            # Badge criteria
│   │   │   ├── types.go               # Badge types (Top Rated, Rising Talent, etc.)
│   │   │   ├── level.go               # Badge levels
│   │   │   └── repository.go
│   │   ├── user_badge/
│   │   │   ├── entity.go              # User badge assignments
│   │   │   ├── achievement_date.go    # When badge was earned
│   │   │   └── repository.go
│   │   ├── reputation/
│   │   │   ├── entity.go              # Reputation scores
│   │   │   ├── score_calculator.go    # Score calculation
│   │   │   ├── history.go             # Reputation history
│   │   │   └── repository.go
│   │   ├── feedback/
│   │   │   ├── entity.go              # Private feedback
│   │   │   ├── category.go            # Feedback categories
│   │   │   └── repository.go
│   │   ├── flag/
│   │   │   ├── entity.go              # Flagged reviews
│   │   │   ├── reason.go              # Flag reasons
│   │   │   └── repository.go
│   │   ├── stats/
│   │   │   ├── entity.go              # Review statistics
│   │   │   ├── aggregates.go          # Aggregated stats
│   │   │   └── repository.go
│   │   ├── review_reminder/
│   │   │   ├── entity.go              # Review reminders
│   │   │   └── repository.go
│   │   └── outbox/
│   │       ├── entity.go
│   │       └── repository.go
│   │
│   ├── application/
│   │   ├── review/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Update, Delete, Respond
│   │   │   ├── queries.go             # Get, List, Filter
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── business_rules.go
│   │   ├── rating/
│   │   │   ├── service.go
│   │   │   ├── aggregator.go          # Aggregate ratings
│   │   │   ├── calculator.go          # Calculate average ratings
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── badge/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Award, Revoke badges
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── achievement_checker.go # Check badge criteria
│   │   ├── reputation/
│   │   │   ├── service.go
│   │   │   ├── calculator.go          # Calculate reputation
│   │   │   ├── updater.go             # Update reputation
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── feedback/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── stats/
│   │   │   ├── service.go
│   │   │   ├── aggregator.go          # Aggregate stats
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   └── eventhandler/
│   │       ├── contract_handler.go    # Enable review after contract
│   │       ├── user_handler.go        # Handle user events
│   │       ├── job_handler.go         # Update job stats
│   │       ├── proposal_handler.go    # Update proposal stats
│   │       ├── payment_handler.go     # Ensure payment before review
│   │       └── admin_handler.go       # Handle admin actions on reviews
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go
│   │   │       ├── transaction.go
│   │   │       ├── migrations.go          # Auto-migrate logic
│   │   │       ├── review_repository.go
│   │   │       ├── rating_repository.go
│   │   │       ├── badge_repository.go
│   │   │       ├── user_badge_repository.go
│   │   │       ├── reputation_repository.go
│   │   │       ├── feedback_repository.go
│   │   │       ├── flag_repository.go
│   │   │       ├── stats_repository.go
│   │   │       ├── review_reminder_repository.go
│   │   │       └── outbox_repository.go
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go
│   │   │       ├── review_cache.go
│   │   │       ├── badge_cache.go
│   │   │       ├── reputation_cache.go
│   │   │       └── stats_cache.go
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go
│   │   │       ├── producer.go
│   │   │       ├── topics.go
│   │   │       └── scram.go
│   │   └── outbox/
│   │       ├── processor.go
│   │       └── scheduler.go
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── review_handler.go
│   │       │   ├── rating_handler.go
│   │       │   ├── badge_handler.go
│   │       │   ├── reputation_handler.go
│   │       │   ├── feedback_handler.go
│   │       │   ├── stats_handler.go
│   │       │   ├── flag_handler.go
│   │       │   └── health_handler.go
│   │       ├── middleware/
│   │       │   ├── auth.go
│   │       │   ├── rbac.go
│   │       │   ├── cors.go
│   │       │   ├── rate_limit.go
│   │       │   ├── logging.go
│   │       │   ├── error_handler.go
│   │       │   └── request_id.go
│   │       ├── responses/
│   │       │   ├── success.go
│   │       │   ├── error.go
│   │       │   └── pagination.go
│   │       └── router.go
│   │
│   └── config/
│       ├── config.go
│       ├── database.go
│       ├── kafka.go
│       ├── redis.go
│       └── services.go
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   └── codes.go
│   ├── logger/
│   │   └── logger.go
│   ├── utils/
│   │   ├── validator.go
│   │   ├── sanitizer.go
│   │   └── sentiment_analyzer.go
│   └── constants/
│       ├── events.go
│       ├── topics.go
│       └── badges.go
│
├── seeds/
│   └── badges.sql                 # Seed initial badges
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   ├── seed-badges.sh
│   └── recalculate-reputation.sh
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── api.md
│   ├── events.md
│   ├── badge-system.md
│   ├── reputation-algorithm.md
│   ├── rating-system.md
│   └── review-guidelines.md
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md
```

---

## 🔟 subscriptions-be (UPDATED)

```
subscriptions-be/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── domain/
│   │   ├── plan/
│   │   │   ├── entity.go              # Subscription plans
│   │   │   ├── features.go            # Plan features
│   │   │   ├── pricing.go             # Pricing tiers
│   │   │   ├── limits.go              # Plan limits
│   │   │   └── repository.go
│   │   ├── subscription/
│   │   │   ├── entity.go              # User subscriptions
│   │   │   ├── billing_cycle.go       # Billing cycle management
│   │   │   ├── enums.go               # Status, Type
│   │   │   ├── auto_renewal.go        # Auto-renewal logic
│   │   │   └── repository.go
│   │   ├── connect/
│   │   │   ├── entity.go              # Connects/Credits
│   │   │   ├── package.go             # Connect packages
│   │   │   ├── transaction.go         # Connect transactions
│   │   │   ├── balance.go             # Balance tracking
│   │   │   └── repository.go
│   │   ├── usage/
│   │   │   ├── entity.go              # Usage tracking
│   │   │   ├── quota.go               # Usage quotas
│   │   │   ├── limit.go               # Usage limits
│   │   │   └── repository.go
│   │   ├── addon/
│   │   │   ├── entity.go              # Plan add-ons
│   │   │   └── repository.go
│   │   ├── promotion/
│   │   │   ├── entity.go              # Promotional codes
│   │   │   ├── discount.go            # Discount rules
│   │   │   ├── usage_limit.go         # Usage limits for promos
│   │   │   └── repository.go
│   │   ├── trial/
│   │   │   ├── entity.go              # Free trials
│   │   │   ├── eligibility.go         # Trial eligibility
│   │   │   └── repository.go
│   │   ├── billing_history/
│   │   │   ├── entity.go              # Billing history
│   │   │   └── repository.go
│   │   ├── feature_toggle/
│   │   │   ├── entity.go              # Feature toggles per plan
│   │   │   └── repository.go
│   │   └── outbox/
│   │       ├── entity.go
│   │       └── repository.go
│   │
│   ├── application/
│   │   ├── plan/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Update plans
│   │   │   ├── queries.go             # Get, List plans
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── subscription/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Subscribe, Cancel, Change, Pause
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── lifecycle_manager.go   # Manage subscription lifecycle
│   │   │   └── renewal_manager.go     # Handle renewals
│   │   ├── connect/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Purchase, Use, Transfer, Refund
│   │   │   ├── queries.go             # Get balance, History
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── calculator.go          # Calculate connect costs
│   │   ├── usage/
│   │   │   ├── service.go
│   │   │   ├── tracker.go             # Track usage
│   │   │   ├── quota_checker.go       # Check quotas
│   │   │   ├── limiter.go             # Enforce limits
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── addon/
│   │   │   ├── service.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── promotion/
│   │   │   ├── service.go
│   │   │   ├── validator.go           # Validate promo codes
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── trial/
│   │   │   ├── service.go
│   │   │   ├── eligibility_checker.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── billing/
│   │   │   ├── service.go
│   │   │   ├── invoice_generator.go   # Generate invoices
│   │   │   ├── payment_processor.go   # Process payments
│   │   │   └── dto.go
│   │   ├── feature_toggle/
│   │   │   ├── service.go
│   │   │   ├── checker.go             # Check feature access
│   │   │   └── dto.go
│   │   └── eventhandler/
│   │       ├── user_handler.go        # Handle new user events
│   │       ├── payment_handler.go     # Handle payment events
│   │       ├── proposal_handler.go    # Deduct connects on proposal
│   │       ├── job_handler.go         # Check posting limits
│   │       └── admin_handler.go       # Handle admin subscription actions
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go
│   │   │       ├── transaction.go
│   │   │       ├── migrations.go          # Auto-migrate logic
│   │   │       ├── plan_repository.go
│   │   │       ├── subscription_repository.go
│   │   │       ├── connect_repository.go
│   │   │       ├── usage_repository.go
│   │   │       ├── addon_repository.go
│   │   │       ├── promotion_repository.go
│   │   │       ├── trial_repository.go
│   │   │       ├── billing_history_repository.go
│   │   │       ├── feature_toggle_repository.go
│   │   │       └── outbox_repository.go
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go
│   │   │       ├── subscription_cache.go
│   │   │       ├── plan_cache.go
│   │   │       ├── connect_cache.go
│   │   │       └── feature_toggle_cache.go
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go
│   │   │       ├── producer.go
│   │   │       ├── topics.go
│   │   │       └── scram.go
│   │   ├── scheduler/
│   │   │   ├── cron.go                # Cron jobs for renewals
│   │   │   └── jobs.go
│   │   ├── payment/
│   │   │   └── client.go              # Financial service client
│   │   └── outbox/
│   │       ├── processor.go
│   │       └── scheduler.go
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── plan_handler.go
│   │       │   ├── subscription_handler.go
│   │       │   ├── connect_handler.go
│   │       │   ├── usage_handler.go
│   │       │   ├── addon_handler.go
│   │       │   ├── promotion_handler.go
│   │       │   ├── trial_handler.go
│   │       │   ├── billing_handler.go
│   │       │   ├── feature_toggle_handler.go
│   │       │   └── health_handler.go
│   │       ├── middleware/
│   │       │   ├── auth.go
│   │       │   ├── rbac.go
│   │       │   ├── cors.go
│   │       │   ├── rate_limit.go
│   │       │   ├── logging.go
│   │       │   ├── error_handler.go
│   │       │   ├── request_id.go
│   │       │   └── feature_gate.go    # Feature access control
│   │       ├── responses/
│   │       │   ├── success.go
│   │       │   ├── error.go
│   │       │   └── pagination.go
│   │       └── router.go
│   │
│   └── config/
│       ├── config.go
│       ├── database.go
│       ├── kafka.go
│       ├── redis.go
│       └── services.go
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   └── codes.go
│   ├── logger/
│   │   └── logger.go
│   ├── utils/
│   │   ├── validator.go
│   │   ├── billing_calculator.go
│   │   └── proration.go
│   └── constants/
│       ├── events.go
│       ├── topics.go
│       ├── plans.go
│       └── features.go
│
├── seeds/
│   ├── plans.sql                  # Seed subscription plans
│   └── connect_packages.sql       # Seed connect packages
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       ├── cronjob-renewal.yaml   # Renewal cron job
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   ├── seed-plans.sh
│   ├── seed-data.sh
│   └── process-renewals.sh
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── docs/
│   ├── README.md
│   ├── api.md
│   ├── events.md
│   ├── subscription-plans.md
│   ├── connects-system.md
│   ├── billing-logic.md
│   └── feature-toggles.md
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md
```

## 1️⃣1️⃣ admin-be (NEW - COMPREHENSIVE)

```
admin-be/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── domain/
│   │   ├── admin_user/
│   │   │   ├── entity.go              # Admin user accounts
│   │   │   ├── role.go                # Admin roles (SuperAdmin, Moderator, Support)
│   │   │   ├── permission.go          # Granular permissions
│   │   │   ├── activity_log.go        # Admin activity tracking
│   │   │   └── repository.go
│   │   ├── support_ticket/
│   │   │   ├── entity.go              # Support tickets
│   │   │   ├── priority.go            # Priority levels (Low, Medium, High, Urgent)
│   │   │   ├── status.go              # Status (Open, InProgress, Resolved, Closed)
│   │   │   ├── category.go            # Ticket categories
│   │   │   ├── assignment.go          # Agent assignment
│   │   │   └── repository.go
│   │   ├── ticket_message/
│   │   │   ├── entity.go              # Ticket conversation messages
│   │   │   ├── attachment.go          # Message attachments
│   │   │   └── repository.go
│   │   ├── support_agent/
│   │   │   ├── entity.go              # Support agent profiles
│   │   │   ├── availability.go        # Agent availability/online status
│   │   │   ├── stats.go               # Agent performance stats
│   │   │   └── repository.go
│   │   ├── canned_response/
│   │   │   ├── entity.go              # Predefined responses
│   │   │   ├── category.go            # Response categories
│   │   │   └── repository.go
│   │   ├── knowledge_base/
│   │   │   ├── entity.go              # KB articles
│   │   │   ├── category.go            # Article categories
│   │   │   ├── tag.go                 # Article tags
│   │   │   ├── version.go             # Article versioning
│   │   │   └── repository.go
│   │   ├── faq/
│   │   │   ├── entity.go              # Frequently asked questions
│   │   │   ├── category.go            # FAQ categories
│   │   │   └── repository.go
│   │   ├── moderation_queue/
│   │   │   ├── entity.go              # Moderation queue items
│   │   │   ├── content_type.go        # Job, User, Review, Message, etc.
│   │   │   ├── flag_reason.go         # Reasons for flagging
│   │   │   ├── action.go              # Moderation actions taken
│   │   │   └── repository.go
│   │   ├── user_action/
│   │   │   ├── entity.go              # Admin actions on users
│   │   │   ├── action_type.go         # Suspend, Ban, Verify, Warn, etc.
│   │   │   ├── reason.go              # Action reasons
│   │   │   └── repository.go
│   │   ├── content_action/
│   │   │   ├── entity.go              # Admin actions on content
│   │   │   ├── action_type.go         # Remove, Hide, Approve, Reject
│   │   │   └── repository.go
│   │   ├── dispute_resolution/
│   │   │   ├── entity.go              # Dispute cases
│   │   │   ├── evidence.go            # Evidence submitted
│   │   │   ├── decision.go            # Admin decision
│   │   │   └── repository.go
│   │   ├── system_config/
│   │   │   ├── entity.go              # System configuration
│   │   │   ├── feature_flag.go        # Feature flags
│   │   │   ├── maintenance.go         # Maintenance mode settings
│   │   │   └── repository.go
│   │   ├── announcement/
│   │   │   ├── entity.go              # Platform announcements
│   │   │   ├── target.go              # Target audience (All, Freelancers, Clients)
│   │   │   ├── schedule.go            # Scheduled announcements
│   │   │   └── repository.go
│   │   ├── report/
│   │   │   ├── entity.go              # Generated reports
│   │   │   ├── report_type.go         # Report types (Users, Revenue, Activity)
│   │   │   ├── schedule.go            # Scheduled reports
│   │   │   └── repository.go
│   │   ├── analytics_dashboard/
│   │   │   ├── entity.go              # Dashboard configurations
│   │   │   ├── widget.go              # Dashboard widgets
│   │   │   ├── metric.go              # Custom metrics
│   │   │   └── repository.go
│   │   ├── audit_log/
│   │   │   ├── entity.go              # Complete audit trail
│   │   │   ├── action.go              # Actions performed
│   │   │   ├── resource.go            # Resources affected
│   │   │   └── repository.go
│   │   ├── notification_blast/
│   │   │   ├── entity.go              # Bulk notifications/emails
│   │   │   ├── audience.go            # Target audience
│   │   │   ├── schedule.go            # Scheduled blasts
│   │   │   └── repository.go
│   │   ├── platform_stats/
│   │   │   ├── entity.go              # Platform statistics
│   │   │   ├── realtime.go            # Real-time stats
│   │   │   └── repository.go
│   │   └── outbox/
│   │       ├── entity.go
│   │       └── repository.go
│   │
│   ├── application/
│   │   ├── admin_user/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Update, Deactivate
│   │   │   ├── queries.go             # Get, List admins
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── permission_manager.go  # Manage permissions
│   │   ├── support_ticket/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Assign, Resolve, Close
│   │   │   ├── queries.go             # Get, List, Search tickets
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   ├── assignment_engine.go   # Auto-assign tickets
│   │   │   ├── escalation_manager.go  # Escalate urgent tickets
│   │   │   └── sla_tracker.go         # Track SLA compliance
│   │   ├── ticket_message/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── support_agent/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── stats_calculator.go    # Calculate agent stats
│   │   ├── canned_response/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── knowledge_base/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Update, Publish
│   │   │   ├── queries.go             # Search, Get articles
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── search_service.go      # KB search
│   │   ├── faq/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── moderation/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Approve, Reject, Remove
│   │   │   ├── queries.go             # Get queue, Filter
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   ├── queue_manager.go       # Manage moderation queue
│   │   │   ├── auto_moderator.go      # Automatic moderation rules
│   │   │   └── content_scanner.go     # Scan for violations
│   │   ├── user_management/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Suspend, Ban, Verify, Warn
│   │   │   ├── queries.go             # Search users, Get details
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── action_validator.go    # Validate admin actions
│   │   ├── content_management/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Remove, Hide, Feature
│   │   │   ├── queries.go             # List content, Search
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── dispute_resolution/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Review, Decide, Close
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── decision_engine.go     # Help make decisions
│   │   ├── system_config/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Update configs
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── announcement/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Schedule, Send
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── validators.go
│   │   ├── report/
│   │   │   ├── service.go
│   │   │   ├── commands.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   ├── generator.go           # Generate reports
│   │   │   ├── scheduler.go           # Schedule reports
│   │   │   └── exporters/
│   │   │       ├── pdf_exporter.go
│   │   │       ├── csv_exporter.go
│   │   │       └── excel_exporter.go
│   │   ├── analytics/
│   │   │   ├── service.go
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── aggregator.go          # Aggregate analytics
│   │   │   ├── metrics_calculator.go  # Calculate KPIs
│   │   │   └── dashboard_builder.go   # Build dashboards
│   │   ├── notification_blast/
│   │   │   ├── service.go
│   │   │   ├── commands.go            # Create, Schedule, Send
│   │   │   ├── queries.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── validators.go
│   │   │   └── audience_selector.go   # Select target users
│   │   ├── audit/
│   │   │   ├── service.go
│   │   │   ├── queries.go             # Query audit logs
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   └── logger.go              # Log admin actions
│   │   └── eventhandler/
│   │       ├── user_handler.go        # Handle user events
│   │       ├── job_handler.go         # Handle job events
│   │       ├── proposal_handler.go    # Handle proposal events
│   │       ├── contract_handler.go    # Handle contract events
│   │       ├── payment_handler.go     # Handle payment events
│   │       ├── review_handler.go      # Handle review flags
│   │       ├── message_handler.go     # Handle message flags
│   │       ├── file_handler.go        # Handle file flags
│   │       └── system_handler.go      # Handle system events
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go
│   │   │       ├── transaction.go
│   │   │       ├── migrations.go          # Auto-migrate logic
│   │   │       ├── admin_user_repository.go
│   │   │       ├── support_ticket_repository.go
│   │   │       ├── ticket_message_repository.go
│   │   │       ├── support_agent_repository.go
│   │   │       ├── canned_response_repository.go
│   │   │       ├── knowledge_base_repository.go
│   │   │       ├── faq_repository.go
│   │   │       ├── moderation_queue_repository.go
│   │   │       ├── user_action_repository.go
│   │   │       ├── content_action_repository.go
│   │   │       ├── dispute_resolution_repository.go
│   │   │       ├── system_config_repository.go
│   │   │       ├── announcement_repository.go
│   │   │       ├── report_repository.go
│   │   │       ├── analytics_dashboard_repository.go
│   │   │       ├── audit_log_repository.go
│   │   │       ├── notification_blast_repository.go
│   │   │       ├── platform_stats_repository.go
│   │   │       └── outbox_repository.go
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go
│   │   │       ├── admin_cache.go
│   │   │       ├── ticket_cache.go
│   │   │       ├── stats_cache.go
│   │   │       └── config_cache.go
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go
│   │   │       ├── producer.go
│   │   │       ├── topics.go
│   │   │       └── scram.go
│   │   ├── external_services/
│   │   │   ├── users_client.go        # Users service client
│   │   │   ├── jobs_client.go         # Jobs service client
│   │   │   ├── proposals_client.go    # Proposals service client
│   │   │   ├── contracts_client.go    # Contracts service client
│   │   │   ├── financial_client.go    # Financial service client
│   │   │   ├── reviews_client.go      # Reviews service client
│   │   │   ├── communications_client.go # Communications service client
│   │   │   ├── search_client.go       # Search service client
│   │   │   ├── storage_client.go      # Storage service client
│   │   │   └── subscriptions_client.go # Subscriptions service client
│   │   ├── keycloak/
│   │   │   ├── admin_client.go        # Keycloak admin operations
│   │   │   └── user_manager.go        # Manage users via Keycloak
│   │   ├── reporting/
│   │   │   ├── pdf_generator.go
│   │   │   ├── csv_generator.go
│   │   │   └── excel_generator.go
│   │   └── outbox/
│   │       ├── processor.go
│   │       └── scheduler.go
│   │
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── admin_user_handler.go
│   │       │   ├── support_ticket_handler.go
│   │       │   ├── ticket_message_handler.go
│   │       │   ├── support_agent_handler.go
│   │       │   ├── canned_response_handler.go
│   │       │   ├── knowledge_base_handler.go
│   │       │   ├── faq_handler.go
│   │       │   ├── moderation_handler.go
│   │       │   ├── user_management_handler.go
│   │       │   ├── content_management_handler.go
│   │       │   ├── dispute_resolution_handler.go
│   │       │   ├── system_config_handler.go
│   │       │   ├── announcement_handler.go
│   │       │   ├── report_handler.go
│   │       │   ├── analytics_handler.go
│   │       │   ├── notification_blast_handler.go
│   │       │   ├── audit_log_handler.go
│   │       │   ├── dashboard_handler.go
│   │       │   └── health_handler.go
│   │       ├── middleware/
│   │       │   ├── auth.go
│   │       │   ├── admin_auth.go      # Admin-specific auth
│   │       │   ├── permission_check.go # Check admin permissions
│   │       │   ├── audit_logger.go    # Auto-log all admin actions
│   │       │   ├── cors.go
│   │       │   ├── rate_limit.go
│   │       │   ├── logging.go
│   │       │   ├── error_handler.go
│   │       │   └── request_id.go
│   │       ├── responses/
│   │       │   ├── success.go
│   │       │   ├── error.go
│   │       │   └── pagination.go
│   │       └── router.go
│   │
│   └── config/
│       ├── config.go
│       ├── database.go
│       ├── kafka.go
│       ├── redis.go
│       ├── keycloak.go
│       └── services.go
│
├── pkg/
│   ├── errors/
│   │   ├── errors.go
│   │   └── codes.go
│   ├── logger/
│   │   └── logger.go
│   ├── utils/
│   │   ├── validator.go
│   │   ├── sanitizer.go
│   │   ├── permission_checker.go
│   │   └── report_formatter.go
│   └── constants/
│       ├── events.go
│       ├── topics.go
│       ├── permissions.go
│       └── moderation_actions.go
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       ├── secrets.yaml
│       ├── hpa.yaml
│       ├── pdb.yaml
│       ├── rbac.yaml                  # Admin RBAC policies
│       └── servicemonitor.yaml
│
├── scripts/
│   ├── setup-local.sh
│   ├── get-secrets.sh
│   ├── seed-admin-users.sh
│   ├── seed-canned-responses.sh
│   └── seed-data.sh
│
├── tests/
│   ├── unit/
│   │   ├── domain/
│   │   │   ├── admin_user_test.go
│   │   │   ├── support_ticket_test.go
│   │   │   └── moderation_queue_test.go
│   │   ├── application/
│   │   │   ├── admin_user_service_test.go
│   │   │   ├── support_ticket_service_test.go
│   │   │   └── moderation_service_test.go
│   │   └── infrastructure/
│   │       ├── postgres_repository_test.go
│   │       └── kafka_producer_test.go
│   ├── integration/
│   │   ├── handlers/
│   │   │   ├── admin_user_handler_test.go
│   │   │   ├── support_ticket_handler_test.go
│   │   │   └── moderation_handler_test.go
│   │   └── repositories/
│   │       ├── admin_user_repository_test.go
│   │       └── support_ticket_repository_test.go
│   └── e2e/
│       └── scenarios/
│           ├── ticket_workflow_test.go
│           ├── moderation_workflow_test.go
│           └── dispute_resolution_test.go
│
├── docs/
│   ├── README.md
│   ├── api.md
│   ├── events.md
│   ├── admin-roles.md
│   ├── permissions.md
│   ├── moderation-guide.md
│   ├── support-workflows.md
│   └── reporting.md
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md

```



---

# Complete Event Flow & Integration Map

## Key Events Between Services

### 1. User Lifecycle Events
```
users-be → All Services
- UserCreated
- UserUpdated
- UserVerified
- FreelancerProfileCompleted
- ClientProfileCompleted
```

### 2. Job Lifecycle Events
```
jobs-be → proposals-be, search-be, communications-be, subscriptions-be
- JobPosted
- JobUpdated
- JobClosed
- JobInvitationSent
```

### 3. Proposal & Bidding Events
```
proposals-be → jobs-be, contracts-be, communications-be, subscriptions-be, financial-be
- ProposalSubmitted
- ProposalAccepted
- ProposalRejected
- BidPlaced
- BidUpdated
- OutbidAlert
- ConnectUsed
```

### 4. Contract Events
```
contracts-be → financial-be, communications-be, reviews-be, users-be
- ContractCreated
- ContractStarted
- MilestoneCreated
- MilestoneCompleted
- MilestoneApproved
- TimesheetSubmitted
- ContractPaused
- ContractEnded
```

### 5. Financial Events
```
financial-be → contracts-be, users-be, communications-be, subscriptions-be
- PaymentProcessed
- PaymentFailed
- EscrowHeld
- EscrowReleased
- PayoutRequested
- PayoutProcessed
- InvoiceGenerated
```

### 6. Communication Events
```
communications-be → All Services
- MessageSent
- NotificationDelivered
- EmailSent
- InAppNotificationSent
```

### 7. Review & Rating Events
```
reviews-be → users-be, search-be, communications-be
- ReviewSubmitted
- ReviewResponded
- BadgeAwarded
- ReputationUpdated
- RatingCalculated
```

### 8. Search & Recommendation Events
```
search-be → communications-be
- JobIndexed
- UserIndexed
- RecommendationGenerated
- MatchFound
```

### 9. Subscription Events
```
subscriptions-be → users-be, jobs-be, proposals-be, communications-be, financial-be
- SubscriptionCreated
- SubscriptionRenewed
- SubscriptionCancelled
- ConnectsPurchased
- ConnectsUsed
- UsageLimitReached
```

### 10. Storage Events
```
storage-be → users-be, jobs-be, contracts-be, proposals-be
- FileUploaded
- FileDeleted
- MediaProcessed
```

## 11. Events Published by admin-be

```
// User Actions
- admin.user.suspended
- admin.user.unsuspended
- admin.user.banned
- admin.user.unbanned
- admin.user.warned
- admin.user.verified

// Content Actions
- admin.job.removed
- admin.job.hidden
- admin.job.featured
- admin.proposal.removed
- admin.review.removed
- admin.message.removed
- admin.file.removed

// Moderation
- admin.moderation.approved
- admin.moderation.rejected
- admin.flag.resolved

// Disputes
- admin.dispute.resolved
- admin.dispute.escalated

// Financial
- admin.payment.refunded
- admin.chargeback.resolved

// Subscriptions
- admin.subscription.cancelled
- admin.subscription.extended
- admin.connects.added

// System
- admin.config.updated
- admin.announcement.published
- admin.maintenance.scheduled
```

---

## Events Subscribed by admin-be

```
// Flagging Events
- user.flagged
- job.flagged
- proposal.flagged
- contract.dispute.opened
- payment.disputed
- review.flagged
- message.flagged
- file.flagged

// System Events
- system.error
- system.alert
- system.abuse.detected

// Support Events
- support.ticket.created
- support.ticket.escalated
```



---

## In-App Notification Types (communications-be)

### Job-Related Notifications
1. **New Job Posted** (matching your skills)
2. **Job Invitation Received**
3. **Job You Applied To Closed**
4. **Similar Jobs Available**

### Proposal & Bidding Notifications
1. **Your Proposal Was Viewed**
2. **Your Proposal Was Accepted**
3. **Your Proposal Was Rejected**
4. **New Bid Placed on Your Job**
5. **You've Been Outbid**
6. **Bid Accepted**

### Contract Notifications
1. **New Contract Created**
2. **Contract Started**
3. **Milestone Completed**
4. **Milestone Approved**
5. **Timesheet Submitted**
6. **Work Diary Updated**
7. **Contract Ending Soon**

### Financial Notifications
1. **Payment Received**
2. **Payment Sent**
3. **Funds in Escrow**
4. **Escrow Released**
5. **Payout Processed**
6. **Invoice Generated**
7. **Low Balance Alert**

### Message Notifications
1. **New Message Received**
2. **Message Read**
3. **User Is Typing**

### Review Notifications
1. **New Review Received**
2. **Review Response Posted**
3. **Badge Earned**

### Subscription Notifications
1. **Subscription Expiring Soon**
2. **Subscription Renewed**
3. **Connects Running Low**
4. **New Connects Purchased**

### System Notifications
1. **Profile Verification Approved/Rejected**
2. **Account Security Alert**
3. **Terms of Service Updated**

---

## Technology Stack Summary

### Backend
- **Language**: Go
- **Framework**: Standard library + Gin/Echo/Chi
- **Architecture**: Clean Architecture (DDD)

### Data Layer
- **Primary DB**: PostgreSQL (per service)
- **Caching**: Redis
- **Search**: Elasticsearch/OpenSearch
- **Object Storage**: MinIO (self-hosted)

### Messaging & Events
- **Event Bus**: Apache Kafka
- **Patterns**: Outbox Pattern, Event Sourcing, CQRS

### Authentication & Authorization
- **Auth Service**: Keycloak
- **Token**: JWT

### Communication
- **Email**: WildDuck (self-hosted SMTP)
- **WebSocket**: Native Go WebSocket
- **Real-time**: Server-Sent Events (SSE)

### Infrastructure
- **Container**: Docker
- **Orchestration**: Kubernetes
- **Service Mesh**: Istio/Linkerd (optional)
- **API Gateway**: Kong/NGINX

### Observability
- **Logging**: ELK Stack or Grafana Loki
- **Monitoring**: Prometheus + Grafana
- **Tracing**: Jaeger

### CI/CD
- **Version Control**: Git
- **CI/CD**: GitHub Actions / GitLab CI
- **GitOps**: ArgoCD

---

## Resource Estimation (Minimum K8s Cluster)

```yaml
# Minimal Production Setup
Nodes: 3 nodes × 8GB RAM = 24GB total

Application Services (10 services):
- users-be: 2 replicas × 512MB = 1GB
- jobs-be: 2 replicas × 512MB = 1GB
- proposals-be: 2 replicas × 512MB = 1GB
- contracts-be: 2 replicas × 512MB = 1GB
- financial-be: 3 replicas × 1GB = 3GB (critical)
- communications-be: 2 replicas × 512MB = 1GB
- storage-be: 2 replicas × 512MB = 1GB
- search-be: 2 replicas × 1GB = 2GB
- reviews-be: 2 replicas × 512MB = 1GB
- subscriptions-be: 2 replicas × 512MB = 1GB

Subtotal: ~14GB

Infrastructure:
- PostgreSQL: 2GB
- Redis: 1GB
- Elasticsearch: 2GB
- Kafka: 2GB
- Keycloak: 1GB
- MinIO: 1GB

Subtotal: ~9GB

Total: ~23GB (fits in 24GB cluster)
```

---

This completes the comprehensive microservices architecture for your Skillsier platform! All services now include:

✅ **Bidding system** in proposals-be
✅ **In-app notifications** with comprehensive coverage in communications-be
✅ **Free & self-hosted** communications (WildDuck, WebSocket, SSE)
✅ **Auto-migrations** following your users-be pattern
✅ **Enhanced recommendation system** in search-be with ML models
✅ **Reviews covering ratings** in reviews-be
✅ Complete event flows between all services
✅ Enterprise-level folder structures

Each service is production-ready and follows Go best practices with Clean Architecture!