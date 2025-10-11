# Users-BE Phase 1 - Complete Implementation Guide

## 📋 Overview
This guide covers all files that need to be created or modified to add profile management features to users-be microservice.

## 🗂️ New Files to Create (60+ files)

### 1. Domain Layer - Entities & Repositories

#### Skills
```
internal/domain/skill/
├── entity.go          # Skill entity with validation
├── repository.go      # Skill repository interface
└── errors.go          # Skill-specific errors
```

#### Work Experience
```
internal/domain/experience/
├── entity.go          # WorkExperience entity
├── repository.go      # Experience repository interface
└── errors.go          # Experience-specific errors
```

#### Education
```
internal/domain/education/
├── entity.go          # Education entity
├── repository.go      # Education repository interface
└── errors.go          # Education-specific errors
```

#### Certifications
```
internal/domain/certification/
├── entity.go          # Certification entity
├── repository.go      # Certification repository interface
└── errors.go          # Certification-specific errors
```

#### Portfolio
```
internal/domain/portfolio/
├── entity.go          # Portfolio entity
├── image.go           # PortfolioImage entity
├── repository.go      # Portfolio repository interface
└── errors.go          # Portfolio-specific errors
```

#### Freelancer Profile
```
internal/domain/freelancer/
├── entity.go          # FreelancerProfile entity
├── repository.go      # Freelancer repository interface
└── errors.go          # Freelancer-specific errors
```

#### Client Profile
```
internal/domain/client/
├── entity.go          # ClientProfile entity
├── repository.go      # Client repository interface
└── errors.go          # Client-specific errors
```

### 2. Application Layer - Services & DTOs

#### Skills Service
```
internal/application/skill/
├── service.go         # Skills business logic
├── dto.go             # Skills DTOs
└── mapper.go          # Entity-DTO mapping
```

#### Experience Service
```
internal/application/experience/
├── service.go         # Experience business logic
├── dto.go             # Experience DTOs
└── mapper.go          # Entity-DTO mapping
```

#### Education Service
```
internal/application/education/
├── service.go         # Education business logic
├── dto.go             # Education DTOs
└── mapper.go          # Entity-DTO mapping
```

#### Certification Service
```
internal/application/certification/
├── service.go         # Certification business logic
├── dto.go             # Certification DTOs
└── mapper.go          # Entity-DTO mapping
```

#### Portfolio Service
```
internal/application/portfolio/
├── service.go         # Portfolio business logic
├── dto.go             # Portfolio DTOs
└── mapper.go          # Entity-DTO mapping
```

#### Freelancer Service
```
internal/application/freelancer/
├── service.go         # Freelancer profile logic
├── dto.go             # Freelancer DTOs
└── mapper.go          # Entity-DTO mapping
```

#### Client Service
```
internal/application/client/
├── service.go         # Client profile logic
├── dto.go             # Client DTOs
└── mapper.go          # Entity-DTO mapping
```

### 3. Infrastructure Layer - Persistence

#### PostgreSQL Repositories
```
internal/infrastructure/persistence/postgres/
├── skill_repository.go          # Skills repo implementation
├── experience_repository.go     # Experience repo implementation
├── education_repository.go      # Education repo implementation
├── certification_repository.go  # Certification repo implementation
├── portfolio_repository.go      # Portfolio repo implementation
├── freelancer_repository.go     # Freelancer profile repo
├── client_repository.go         # Client profile repo
└── file_storage.go              # File upload handling
```

#### Storage Service
```
internal/infrastructure/storage/
├── local.go           # Local file storage
└── s3.go              # S3 storage (future)
```

### 4. HTTP Interface - Handlers & Routes

#### HTTP Handlers
```
internal/interfaces/http/handlers/
├── skill_handler.go          # Skills CRUD endpoints
├── experience_handler.go     # Experience CRUD endpoints
├── education_handler.go      # Education CRUD endpoints
├── certification_handler.go  # Certification CRUD endpoints
├── portfolio_handler.go      # Portfolio CRUD endpoints
├── freelancer_handler.go     # Freelancer profile endpoints
├── client_handler.go         # Client profile endpoints
├── avatar_handler.go         # Avatar upload/delete
└── stats_handler.go          # User statistics
```

### 5. Configuration Updates

#### Configuration
```
internal/config/config.go      # Add storage config
```

#### Environment Variables
```
.env.example                   # Add new env vars
```

### 6. Database Migrations

#### Migration Update
```
internal/infrastructure/persistence/postgres/migrations.go
```
Add migrations for:
- skills table
- work_experiences table
- educations table
- certifications table
- portfolios table
- portfolio_images table
- freelancer_profiles table
- client_profiles table

### 7. Update Existing Files

#### Update User Entity
```
internal/domain/user/entity.go
```
Add fields:
- AvatarURL
- PhoneNumber
- Bio
- Profession
- HourlyRate
- AvailableHours
- Country
- City
- ProfileComplete

#### Update User Service
```
internal/application/user/service.go
```
Add methods:
- UploadAvatar
- DeleteAvatar
- UpdatePreferences
- ChangePassword
- GetStats
- GetEarnings

#### Update User Handler
```
internal/interfaces/http/handlers/user_handler.go
```
Add endpoints:
- POST /users/profile/avatar
- DELETE /users/profile/avatar
- PATCH /users/profile/preferences
- POST /users/profile/password
- GET /users/profile/stats

#### Update Router
```
internal/interfaces/http/router.go
```
Add new routes for all new handlers

#### Update Main
```
cmd/api/main.go
```
Initialize new repositories and services

## 🔨 Implementation Steps

### Step 1: Create Domain Entities (Day 1)

Create all entity files with:
- Proper GORM tags
- Validation methods
- Table names
- JSON tags

### Step 2: Create Repository Interfaces (Day 1)

Define repository interfaces for each entity with:
- CRUD operations
- Custom queries
- Error definitions

### Step 3: Implement PostgreSQL Repositories (Day 2)

Implement all repository interfaces with:
- GORM operations
- Transactions support
- Error handling
- Indexing for performance

### Step 4: Update Database Migrations (Day 2)

Add auto-migration for all new tables with:
- Foreign keys
- Indexes
- Constraints
- Default values

### Step 5: Create DTOs and Mappers (Day 3)

Create DTOs for each entity:
- CreateDTO
- UpdateDTO
- ResponseDTO
- ListResponseDTO

Implement mappers:
- Entity to DTO
- DTO to Entity
- Null handling

### Step 6: Implement Services (Day 4-5)

Implement business logic services:
- Input validation
- Business rules
- Outbox event creation
- Transaction management

### Step 7: Create HTTP Handlers (Day 6)

Implement HTTP handlers:
- Request binding
- Input validation
- Service calls
- Response formatting
- Error handling

### Step 8: Update Routes (Day 6)

Add all new routes:
- Skills routes
- Experience routes
- Education routes
- Certification routes
- Portfolio routes
- Profile routes

### Step 9: File Upload Implementation (Day 7)

Implement file upload:
- Avatar upload
- Portfolio image upload
- File validation
- Image processing
- Storage service

### Step 10: Testing & Integration (Day 8-10)

- Unit tests for services
- Integration tests for handlers
- API testing with Postman
- Load testing
- Documentation updates

## 📊 Database Schema Updates

### New Tables:

```sql
-- Skills table
CREATE TABLE skills (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    level VARCHAR(50) NOT NULL,
    years_of_experience INT,
    endorsements INT DEFAULT 0,
    is_primary BOOLEAN DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_skills_user_id (user_id)
);

-- Work experiences table
CREATE TABLE work_experiences (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    company VARCHAR(200) NOT NULL,
    location VARCHAR(200),
    start_date DATE NOT NULL,
    end_date DATE,
    is_current BOOLEAN DEFAULT false,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_work_exp_user_id (user_id)
);

-- Educations table
CREATE TABLE educations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    school VARCHAR(200) NOT NULL,
    degree VARCHAR(200) NOT NULL,
    field_of_study VARCHAR(200),
    start_date DATE NOT NULL,
    end_date DATE,
    is_current BOOLEAN DEFAULT false,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_education_user_id (user_id)
);

-- Certifications table
CREATE TABLE certifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(200) NOT NULL,
    issuing_organization VARCHAR(200) NOT NULL,
    issue_date DATE NOT NULL,
    expiry_date DATE,
    credential_id VARCHAR(200),
    credential_url TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_cert_user_id (user_id)
);

-- Portfolios table
CREATE TABLE portfolios (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    project_url TEXT,
    start_date DATE,
    end_date DATE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_portfolio_user_id (user_id)
);

-- Portfolio images table
CREATE TABLE portfolio_images (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    portfolio_id UUID NOT NULL REFERENCES portfolios(id) ON DELETE CASCADE,
    image_url TEXT NOT NULL,
    display_order INT DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_portfolio_img_portfolio_id (portfolio_id)
);

-- Freelancer profiles table
CREATE TABLE freelancer_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(200),
    overview TEXT,
    hourly_rate DECIMAL(10,2),
    available_hours INT,
    response_time INT, -- in hours
    total_jobs INT DEFAULT 0,
    total_earnings DECIMAL(12,2) DEFAULT 0,
    success_rate DECIMAL(5,2) DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Client profiles table
CREATE TABLE client_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    company_name VARCHAR(200),
    company_size VARCHAR(50),
    industry VARCHAR(100),
    total_spent DECIMAL(12,2) DEFAULT 0,
    total_jobs_posted INT DEFAULT 0,
    total_hired INT DEFAULT 0,
    payment_verified BOOLEAN DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```

## 🔄 Event Updates

### New Events to Publish:

```go
// Skill events
- skill.created
- skill.updated
- skill.deleted

// Experience events
- experience.added
- experience.updated
- experience.deleted

// Education events
- education.added
- education.updated
- education.deleted

// Certification events
- certification.added
- certification.updated
- certification.deleted

// Portfolio events
- portfolio.item.created
- portfolio.item.updated
- portfolio.item.deleted
- portfolio.image.uploaded

// Profile events
- freelancer.profile.updated
- client.profile.updated
- profile.completed
```

## 🧪 Testing Checklist

- [ ] Unit tests for all services
- [ ] Repository integration tests
- [ ] HTTP handler tests
- [ ] File upload tests
- [ ] Validation tests
- [ ] Error handling tests
- [ ] Event publishing tests
- [ ] Transaction rollback tests

## 📚 API Documentation

Update README.md with all new endpoints and their request/response formats.

## 🚀 Deployment

After implementation:
1. Update Dockerfile if needed
2. Update K8s manifests with new env vars
3. Deploy to staging
4. Run integration tests
5. Deploy to production

## 📝 Summary

**Total New Files**: ~60 files
**Lines of Code**: ~8,000-10,000 lines
**Estimated Time**: 8-10 days for one developer
**Database Tables**: 8 new tables
**API Endpoints**: 30+ new endpoints

Would you like me to generate the complete code for any specific component?