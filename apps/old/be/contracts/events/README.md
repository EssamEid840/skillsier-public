# Skillsier Event Contracts

Versioned event schemas for the Skillsier platform using Protocol Buffers.

## Overview

This module contains all event definitions for inter-service communication in the Skillsier platform. All events are defined using Protocol Buffers for:

- **Strong typing** - Compile-time type safety
- **Backward compatibility** - Controlled schema evolution
- **Cross-language support** - Use in Go, Python, Java, etc.
- **Performance** - Efficient binary serialization
- **Documentation** - Self-documenting schemas

## Structure

```
contracts/events/
├── buf.yaml                    # Buf configuration
├── buf.gen.yaml               # Code generation config
├── go.mod                     # Go module definition
├── README.md                  # This file
├── EVENTS.md                  # Complete event catalog
│
├── user/v1/                   # User lifecycle events
│   ├── user_created.proto
│   ├── user_updated.proto
│   ├── user_verified.proto
│   ├── user_suspended.proto
│   ├── freelancer_profile_completed.proto
│   └── client_profile_completed.proto
│
├── job/v1/                    # Job lifecycle events
│   ├── job_posted.proto
│   ├── job_updated.proto
│   ├── job_closed.proto
│   └── job_invitation_sent.proto
│
├── proposal/v1/               # Proposal & bidding events
│   ├── proposal_submitted.proto
│   ├── proposal_accepted.proto
│   ├── bid_placed.proto
│   ├── bid_updated.proto
│   └── outbid_alert.proto
│
├── contract/v1/               # Contract events
│   ├── contract_created.proto
│   ├── milestone_completed.proto
│   ├── timesheet_submitted.proto
│   └── dispute_opened.proto
│
├── payment/v1/                # Financial events
│   ├── payment_processed.proto
│   ├── escrow_held.proto
│   ├── escrow_released.proto
│   └── payout_requested.proto
│
├── review/v1/                 # Review events
│   ├── review_submitted.proto
│   └── badge_awarded.proto
│
├── subscription/v1/           # Subscription events
│   ├── subscription_created.proto
│   ├── connects_purchased.proto
│   └── connects_used.proto
│
├── message/v1/                # Communication events
│   ├── message_sent.proto
│   └── notification_delivered.proto
│
├── storage/v1/                # File storage events
│   ├── file_uploaded.proto
│   └── media_processed.proto
│
├── search/v1/                 # Search & recommendation events
│   ├── job_indexed.proto
│   └── recommendation_generated.proto
│
├── admin/v1/                  # Admin action events
│   ├── user_suspended.proto
│   ├── content_removed.proto
│   └── dispute_resolved.proto
│
└── gen/                       # Generated code
    └── go/
        └── skillsier/
            ├── user/v1/
            ├── job/v1/
            ├── proposal/v1/
            └── ...
```

## Installation

### For Go Services

```bash
go get skillsier.dev/contracts/events@latest
```

### In your go.mod

```go
require (
    skillsier.dev/contracts/events v1.0.0
)
```

## Usage

### Import Events in Go

```go
import (
    userv1 "skillsier.dev/contracts/events/gen/go/user/v1"
    jobv1 "skillsier.dev/contracts/events/gen/go/job/v1"
    "google.golang.org/protobuf/types/known/timestamppb"
)
```

### Create an Event

```go
event := &userv1.UserCreated{
    EventId:         uuid.New().String(),
    EventTimestamp:  timestamppb.Now(),
    AggregateId:     userID,
    EventVersion:    1,
    UserId:          userID,
    KeycloakId:      keycloakID,
    Username:        "john_doe",
    Email:           "john@example.com",
    EmailVerified:   true,
    FirstName:       "John",
    LastName:        "Doe",
    UserType:        userv1.UserType_USER_TYPE_FREELANCER,
    Status:          userv1.AccountStatus_ACCOUNT_STATUS_ACTIVE,
    CreatedAt:       timestamppb.Now(),
}
```

### Serialize to Bytes

```go
import "google.golang.org/protobuf/proto"

bytes, err := proto.Marshal(event)
if err != nil {
    log.Fatal(err)
}

// Publish to Kafka
err = publisher.Publish(ctx, "user.created", userID, bytes)
```

### Deserialize from Bytes

```go
var event userv1.UserCreated
err := proto.Unmarshal(bytes, &event)
if err != nil {
    log.Fatal(err)
}

// Use event
fmt.Printf("User %s created at %v\n", event.Username, event.CreatedAt.AsTime())
```

### Using with Outbox Pattern

```go
import (
    userv1 "skillsier.dev/contracts/events/gen/go/user/v1"
    "skillsier.dev/platform-shared/outbox"
)

// Create event
event := &userv1.UserCreated{ /* ... */ }

// Marshal to JSON (or protobuf bytes)
payload, _ := protojson.Marshal(event)

// Create outbox event
outboxEvent, _ := outbox.NewEvent(
    event.UserId,           // aggregate_id
    "user",                 // aggregate_type
    "user.created",         // event_type
    payload,                // payload
)

// Publish within transaction
publisher.PublishWithTx(tx, outboxEvent)
```

## Development

### Prerequisites

- Go 1.23+
- [buf CLI](https://buf.build/docs/installation)
- Protocol Buffers compiler (protoc)

### Generate Code

```bash
# Install buf (if not already installed)
brew install bufbuild/buf/buf

# Generate Go code from .proto files
buf generate

# Or using make
make generate
```

### Lint Proto Files

```bash
buf lint
```

### Check for Breaking Changes

```bash
# Compare against main branch
buf breaking --against '.git#branch=main'

# Compare against a tag
buf breaking --against '.git#tag=v1.0.0'
```

### Add a New Event

1. Create the `.proto` file in the appropriate domain directory
2. Follow the naming convention: `<event_name>.proto`
3. Use the standard event structure with common fields
4. Run `buf lint` to ensure compliance
5. Run `buf generate` to generate Go code
6. Update EVENTS.md with the new event details

### Version a Proto File

When making changes:

**Backward Compatible (Minor Version):**
- ✅ Add new optional fields
- ✅ Add new enum values (with proper defaults)
- ✅ Add new message types

**Breaking Changes (Major Version):**
- ❌ Remove fields
- ❌ Change field types
- ❌ Change field numbers
- ❌ Rename fields

For breaking changes:
1. Create new version: `user/v2/user_created.proto`
2. Update imports in services gradually
3. Deprecate old version
4. Remove after migration period

## Event Catalog

See [EVENTS.md](./EVENTS.md) for complete catalog with:
- All event types and fields
- Topic names and partition keys
- Consumer mapping
- Event flow diagrams
- Field completeness scores

## Best Practices

### 1. Always Include Base Fields

Every event MUST have:
- `event_id` (UUID)
- `event_timestamp` (timestamp)
- `aggregate_id` (primary entity ID)
- `event_version` (schema version number)

### 2. Use Enums with UNSPECIFIED Default

```protobuf
enum UserType {
  USER_TYPE_UNSPECIFIED = 0;  // Always have this
  USER_TYPE_FREELANCER = 1;
  USER_TYPE_CLIENT = 2;
}
```

### 3. Mark Deprecated Fields

```protobuf
string old_field = 5 [deprecated = true];
```

### 4. Use Well-Known Types

```protobuf
import "google/protobuf/timestamp.proto";
import "google/protobuf/duration.proto";
import "google/protobuf/wrappers.proto";

google.protobuf.Timestamp created_at = 1;
google.protobuf.Duration timeout = 2;
google.protobuf.StringValue optional_field = 3;
```

### 5. Document Fields

```protobuf
message UserCreated {
  // Unique event identifier for deduplication
  string event_id = 1;
  
  // ISO 3166-1 alpha-2 country code (e.g., US, UK, EG)
  string country_code = 2;
}
```

### 6. Use Appropriate Data Types

- IDs: `string` (UUIDs, external IDs)
- Money: Use custom `Money` message with amount + currency
- Timestamps: `google.protobuf.Timestamp`
- Durations: `google.protobuf.Duration`
- Percentages: `double` or `int32` (basis points)

### 7. Event Naming

- PascalCase for message names: `UserCreated`
- snake_case for field names: `user_id`
- Past tense for events: `Created`, not `Create`
- Specific and descriptive: `PaymentProcessed` not `Payment`

## Monitoring & Observability

### Event Metrics

Track these metrics per event type:
- Published count
- Failed publishes
- Serialization errors
- Schema validation errors
- Consumer lag
- Processing time

### Event Validation

All events are validated:
1. At publish time (schema validation)
2. By `buf lint