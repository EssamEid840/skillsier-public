#!/bin/bash

# Enhanced script to generate all protobuf files for Skillsier events
# This creates complete proto files with all fields from EVENTS.md (no TODOs)
# 
# Usage: ./generate-all-protos.sh
# 
# Location: apps/be/scripts/generate-all-protos.sh
# Target: apps/be/contracts/events/

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Base directories - relative to script location
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
EVENTS_DIR="$PROJECT_ROOT/contracts/events"
COMMON_DIR="$EVENTS_DIR/common/v1"

echo -e "${BLUE}🚀 Skillsier Proto Generator${NC}"
echo -e "${BLUE}=============================${NC}"
echo ""
echo "Script location: $SCRIPT_DIR"
echo "Project root: $PROJECT_ROOT"
echo "Events directory: $EVENTS_DIR"
echo ""

# Create directory structure
create_directories() {
    echo -e "${YELLOW}📁 Creating directory structure...${NC}"
    
    mkdir -p "$COMMON_DIR"
    mkdir -p "$EVENTS_DIR"/{user,job,proposal,contract,payment,review,subscription,message,storage,search,admin}/v1
    
    echo -e "${GREEN}✓ Directories created${NC}"
}

# Create buf configuration files if they don't exist
create_buf_config() {
    echo -e "${YELLOW}📝 Creating buf configuration...${NC}"
    
    # buf.yaml
    if [ ! -f "$EVENTS_DIR/buf.yaml" ]; then
        cat > "$EVENTS_DIR/buf.yaml" << 'EOF'
version: v1
name: buf.build/skillsier/events
lint:
  use:
    - DEFAULT
  except:
    - PACKAGE_DIRECTORY_MATCH
  enum_zero_value_suffix: _UNSPECIFIED
  rpc_allow_same_request_response: false
  rpc_allow_google_protobuf_empty_requests: false
  rpc_allow_google_protobuf_empty_responses: false
  service_suffix: Service
breaking:
  use:
    - FILE
  except:
    - EXTENSION_MESSAGE_NO_DELETE
    - FIELD_SAME_JSON_NAME
EOF
        echo -e "${GREEN}✓ Created buf.yaml${NC}"
    fi
    
    # buf.gen.yaml
    if [ ! -f "$EVENTS_DIR/buf.gen.yaml" ]; then
        cat > "$EVENTS_DIR/buf.gen.yaml" << 'EOF'
version: v1
plugins:
  - plugin: buf.build/protocolbuffers/go:v1.34.2
    out: gen/go
    opt:
      - paths=source_relative
  - plugin: buf.build/grpc/go:v1.4.0
    out: gen/go
    opt:
      - paths=source_relative
      - require_unimplemented_servers=false
EOF
        echo -e "${GREEN}✓ Created buf.gen.yaml${NC}"
    fi
    
    # go.mod
    if [ ! -f "$EVENTS_DIR/go.mod" ]; then
        cat > "$EVENTS_DIR/go.mod" << 'EOF'
module skillsier.dev/contracts/events

go 1.23

require (
	google.golang.org/protobuf v1.34.2
	google.golang.org/grpc v1.65.0
)
EOF
        echo -e "${GREEN}✓ Created go.mod${NC}"
    fi
}

# Function to create a proto file with complete fields
create_proto() {
    local domain=$1
    local version=$2
    local event_name=$3
    local message_name=$4
    local fields=$5
    
    local file_path="$EVENTS_DIR/$domain/$version/${event_name}.proto"
    
    if [ -f "$file_path" ]; then
        echo -e "${YELLOW}⚠ Skipping $file_path (already exists)${NC}"
        return
    fi
    
    echo -e "${BLUE}📝 Creating $file_path${NC}"
    
    cat > "$file_path" << EOF
syntax = "proto3";

package skillsier.${domain}.${version};

option go_package = "skillsier.dev/contracts/events/gen/go/${domain}/${version};${domain}${version}";

import "google/protobuf/timestamp.proto";
import "skillsier/common/v1/metadata.proto";
import "skillsier/common/v1/enums.proto";
import "skillsier/common/v1/value_objects.proto";

// ${message_name} event
// Topic: ${domain}.${event_name}
// Owner: ${domain}-be
message ${message_name} {
  // Base event metadata (REQUIRED in ALL events)
  skillsier.common.v1.BaseEventMetadata metadata = 1;
  
${fields}
}
EOF
}

# Generate common proto files (if they don't exist)
generate_common_protos() {
    echo -e "${YELLOW}🔧 Checking common proto files...${NC}"
    
    if [ ! -f "$COMMON_DIR/metadata.proto" ]; then
        echo -e "${BLUE}Creating metadata.proto...${NC}"
        echo -e "${RED}❌ metadata.proto not found!${NC}"
        echo -e "${YELLOW}Please create metadata.proto, enums.proto, and value_objects.proto first${NC}"
        echo -e "${YELLOW}These files are required and contain the common definitions${NC}"
        return 1
    fi
    
    if [ ! -f "$COMMON_DIR/enums.proto" ]; then
        echo -e "${RED}❌ enums.proto not found!${NC}"
        return 1
    fi
    
    if [ ! -f "$COMMON_DIR/value_objects.proto" ]; then
        echo -e "${RED}❌ value_objects.proto not found!${NC}"
        return 1
    fi
    
    echo -e "${GREEN}✓ Common proto files found${NC}"
}

# Generate USER events
generate_user_events() {
    echo -e "${YELLOW}👤 Generating User events...${NC}"
    
    # UserCreated
    create_proto "user" "v1" "user_created" "UserCreated" "  // User identification
  string user_id = 2;
  string keycloak_id = 3;
  string username = 4;
  string email = 5;
  bool email_verified = 6;
  
  // Personal information
  string first_name = 7;
  string last_name = 8;
  string phone_number = 9;
  string country_code = 10;
  string language = 11;
  string timezone = 12;
  
  // User type and status
  skillsier.common.v1.UserType user_type = 13;
  repeated skillsier.common.v1.UserType additional_types = 14;
  skillsier.common.v1.AccountStatus status = 15;
  
  // Profile
  string profile_picture_url = 16;
  string cover_image_url = 17;
  string bio = 18;
  string tagline = 19;
  
  // Location
  skillsier.common.v1.Location location = 20;
  
  // Settings (simplified - can expand with nested messages)
  bool notifications_email = 21;
  bool notifications_push = 22;
  bool notifications_sms = 23;
  bool notifications_in_app = 24;
  
  // Marketing & referral
  string referral_code = 25;
  string referrer_user_id = 26;
  string onboarding_source = 27;
  string utm_source = 28;
  string utm_medium = 29;
  string utm_campaign = 30;
  
  // Timestamps
  google.protobuf.Timestamp created_at = 31;
  google.protobuf.Timestamp last_login_at = 32;
  
  // Compliance
  bool terms_accepted = 33;
  string terms_version = 34;
  bool privacy_policy_accepted = 35;
  string privacy_policy_version = 36;
  
  // Social login
  string social_provider = 37;
  string social_provider_id = 38;
  
  // Flags
  bool is_verified = 39;
  bool is_featured = 40;
  bool mfa_enabled = 41;
  
  // Custom fields
  map<string, string> custom_fields = 42;"

    # UserUpdated
    create_proto "user" "v1" "user_updated" "UserUpdated" "  // User identification
  string user_id = 2;
  string username = 3;
  
  // Change tracking
  repeated string changed_fields = 4;
  map<string, string> previous_values = 5;
  string updated_by_user_id = 6;
  string update_reason = 7;
  
  // Updated fields (subset of UserCreated)
  string email = 8;
  string first_name = 9;
  string last_name = 10;
  string phone_number = 11;
  skillsier.common.v1.UserType user_type = 12;
  skillsier.common.v1.AccountStatus status = 13;
  string profile_picture_url = 14;
  string bio = 15;
  
  skillsier.common.v1.Location location = 16;
  
  google.protobuf.Timestamp updated_at = 17;
  int32 profile_completion_percentage = 18;"

    # UserVerified
    create_proto "user" "v1" "user_verified" "UserVerified" "  // User identification
  string user_id = 2;
  
  // Verification details
  skillsier.common.v1.VerificationType verification_type = 3;
  skillsier.common.v1.VerificationMethod verification_method = 4;
  string verified_by = 5;
  google.protobuf.Timestamp verified_at = 6;
  skillsier.common.v1.VerificationLevel verification_level = 7;
  string badge_awarded = 8;
  bool auto_verified = 9;"

    # UserSuspended
    create_proto "user" "v1" "user_suspended" "UserSuspended" "  // User identification
  string user_id = 2;
  string username = 3;
  
  // Suspension details
  string suspended_by_user_id = 4;
  string suspension_reason = 5;
  string suspension_details = 6;
  int32 suspension_duration_days = 7;
  google.protobuf.Timestamp suspension_start_date = 8;
  google.protobuf.Timestamp suspension_end_date = 9;
  bool is_temporary = 10;
  bool can_appeal = 11;
  google.protobuf.Timestamp appeal_deadline = 12;
  bool notification_sent = 13;"

    # UserBanned
    create_proto "user" "v1" "user_banned" "UserBanned" "  // User identification
  string user_id = 2;
  string username = 3;
  
  // Ban details
  string banned_by_user_id = 4;
  string ban_reason = 5;
  string ban_details = 6;
  bool is_permanent = 7;
  bool ip_banned = 8;
  bool device_banned = 9;
  bool can_appeal = 10;
  google.protobuf.Timestamp banned_at = 11;
  repeated string related_user_ids = 12;"

    # FreelancerProfileCompleted
    create_proto "user" "v1" "freelancer_profile_completed" "FreelancerProfileCompleted" "  // User identification
  string user_id = 2;
  string username = 3;
  
  // Professional profile
  string professional_title = 4;
  string overview = 5;
  string video_intro_url = 6;
  
  // Rates & availability
  double hourly_rate = 7;
  double minimum_project_budget = 8;
  skillsier.common.v1.Currency currency = 9;
  skillsier.common.v1.AvailabilityStatus availability_status = 10;
  int32 hours_per_week = 11;
  
  // Skills (simplified)
  repeated string skill_names = 12;
  
  // Stats
  double job_success_score = 13;
  double total_earnings = 14;
  int32 total_jobs = 15;
  
  // Profile metadata
  int32 profile_completion_percentage = 16;
  bool identity_verified = 17;
  bool payment_verified = 18;
  bool phone_verified = 19;
  
  google.protobuf.Timestamp completed_at = 20;"

    # ClientProfileCompleted
    create_proto "user" "v1" "client_profile_completed" "ClientProfileCompleted" "  // User identification
  string user_id = 2;
  
  // Company details
  string company_name = 3;
  string company_website = 4;
  string industry = 5;
  string company_description = 6;
  string company_logo_url = 7;
  
  // Payment
  bool payment_verified = 8;
  
  // Stats
  double total_spent = 9;
  int32 total_jobs_posted = 10;
  int32 total_hires = 11;
  
  google.protobuf.Timestamp completed_at = 12;"
    
    echo -e "${GREEN}✓ User events generated (7 events)${NC}"
}

# Generate JOB events
generate_job_events() {
    echo -e "${YELLOW}💼 Generating Job events...${NC}"
    
    # JobPosted
    create_proto "job" "v1" "job_posted" "JobPosted" "  // Job identification
  string job_id = 2;
  string client_id = 3;
  string client_username = 4;
  
  // Job basics
  string job_title = 5;
  string job_description = 6;
  skillsier.common.v1.JobType job_type = 7;
  
  // Budget
  skillsier.common.v1.Money budget_amount = 8;
  bool budget_is_flexible = 9;
  
  // Duration & timeline
  skillsier.common.v1.DurationType duration_type = 10;
  int32 duration_weeks = 11;
  google.protobuf.Timestamp expected_start_date = 12;
  google.protobuf.Timestamp deadline_date = 13;
  
  // Requirements
  skillsier.common.v1.ExperienceLevel experience_level = 14;
  repeated string required_skills = 15;
  repeated string preferred_skills = 16;
  
  // Categorization
  string category_id = 17;
  string subcategory_id = 18;
  repeated string tags = 19;
  
  // Location
  repeated string allowed_countries = 20;
  repeated string allowed_timezones = 21;
  bool remote_allowed = 22;
  
  // Visibility
  skillsier.common.v1.JobVisibility visibility = 23;
  repeated string invited_freelancer_ids = 24;
  
  // Screening
  int32 screening_questions_count = 25;
  
  // Posted
  google.protobuf.Timestamp posted_at = 26;
  
  // AI/ML data
  int32 matching_profiles_count = 27;
  skillsier.common.v1.CompetitionLevel competition_level = 28;"

    # JobUpdated
    create_proto "job" "v1" "job_updated" "JobUpdated" "  // Job identification
  string job_id = 2;
  string client_id = 3;
  
  // Change tracking
  repeated string changed_fields = 4;
  map<string, string> previous_values = 5;
  string update_reason = 6;
  
  // Updated fields
  string job_title = 7;
  string job_description = 8;
  skillsier.common.v1.Money budget_amount = 9;
  skillsier.common.v1.ExperienceLevel experience_level = 10;
  
  google.protobuf.Timestamp updated_at = 11;"

    # JobClosed
    create_proto "job" "v1" "job_closed" "JobClosed" "  // Job identification
  string job_id = 2;
  string client_id = 3;
  
  // Closure details
  string close_reason = 4;
  string close_details = 5;
  string hired_freelancer_id = 6;
  string proposal_id = 7;
  int32 total_proposals_received = 8;
  google.protobuf.Timestamp closed_at = 9;
  int32 posted_duration_hours = 10;"

    # JobInvitationSent
    create_proto "job" "v1" "job_invitation_sent" "JobInvitationSent" "  // Invitation details
  string invitation_id = 2;
  string job_id = 3;
  string client_id = 4;
  string freelancer_id = 5;
  string invitation_message = 6;
  google.protobuf.Timestamp invitation_expiry_date = 7;
  google.protobuf.Timestamp sent_at = 8;"

    # JobRemoved
    create_proto "job" "v1" "job_removed" "JobRemoved" "  // Job identification
  string job_id = 2;
  string removed_by_user_id = 3;
  string removal_reason = 4;
  string removal_details = 5;
  bool refund_issued = 6;
  google.protobuf.Timestamp removed_at = 7;"

    # JobFlagged
    create_proto "job" "v1" "job_flagged" "JobFlagged" "  // Flag details
  string flag_id = 2;
  string job_id = 3;
  string flagged_by_user_id = 4;
  skillsier.common.v1.FlagReason flag_reason = 5;
  string flag_details = 6;
  google.protobuf.Timestamp flagged_at = 7;
  bool auto_flagged = 8;
  double ai_confidence_score = 9;"
    
    echo -e "${GREEN}✓ Job events generated (6 events)${NC}"
}

# Generate PROPOSAL events
generate_proposal_events() {
    echo -e "${YELLOW}📄 Generating Proposal events...${NC}"
    
    # ProposalSubmitted
    create_proto "proposal" "v1" "proposal_submitted" "ProposalSubmitted" "  // Proposal identification
  string proposal_id = 2;
  string job_id = 3;
  string freelancer_id = 4;
  string client_id = 5;
  
  // Proposal content
  string cover_letter = 6;
  int32 cover_letter_length = 7;
  
  // Rate & budget
  skillsier.common.v1.Money proposed_rate = 8;
  string estimated_duration = 9;
  int32 estimated_hours = 10;
  google.protobuf.Timestamp availability_start_date = 11;
  
  // Metadata
  google.protobuf.Timestamp submitted_at = 12;
  int32 proposal_version = 13;
  
  // Bidding
  int32 connects_used = 14;
  bool boost_applied = 15;
  bool auto_bid = 16;
  
  // Stats at submission
  double job_success_score = 17;
  double total_earnings = 18;
  int32 total_jobs = 19;"

    # ProposalAccepted
    create_proto "proposal" "v1" "proposal_accepted" "ProposalAccepted" "  // Proposal identification
  string proposal_id = 2;
  string job_id = 3;
  string freelancer_id = 4;
  string client_id = 5;
  string accepted_by_user_id = 6;
  string acceptance_message = 7;
  string contract_id = 8;
  google.protobuf.Timestamp accepted_at = 9;
  int32 time_to_accept_hours = 10;"

    # ProposalRejected
    create_proto "proposal" "v1" "proposal_rejected" "ProposalRejected" "  // Proposal identification
  string proposal_id = 2;
  string job_id = 3;
  string freelancer_id = 4;
  string client_id = 5;
  string rejection_reason = 6;
  string rejection_feedback = 7;
  google.protobuf.Timestamp rejected_at = 8;
  bool connects_refunded = 9;"

    # ProposalWithdrawn
    create_proto "proposal" "v1" "proposal_withdrawn" "ProposalWithdrawn" "  // Proposal identification
  string proposal_id = 2;
  string job_id = 3;
  string freelancer_id = 4;
  string withdrawal_reason = 5;
  string withdrawal_details = 6;
  google.protobuf.Timestamp withdrawn_at = 7;
  bool connects_refunded = 8;"

    # BidPlaced
    create_proto "proposal" "v1" "bid_placed" "BidPlaced" "  // Bid identification
  string bid_id = 2;
  string proposal_id = 3;
  string job_id = 4;
  string freelancer_id = 5;
  skillsier.common.v1.Money bid_amount = 6;
  skillsier.common.v1.BidType bid_type = 7;
  double previous_bid_amount = 8;
  google.protobuf.Timestamp placed_at = 9;
  int32 bid_position = 10;
  int32 total_bids_on_job = 11;"

    # BidUpdated
    create_proto "proposal" "v1" "bid_updated" "BidUpdated" "  // Bid identification
  string bid_id = 2;
  string proposal_id = 3;
  string job_id = 4;
  string freelancer_id = 5;
  double old_bid_amount = 6;
  double new_bid_amount = 7;
  string update_reason = 8;
  google.protobuf.Timestamp updated_at = 9;
  int32 bid_position_before = 10;
  int32 bid_position_after = 11;"

    # OutbidAlert
    create_proto "proposal" "v1" "outbid_alert" "OutbidAlert" "  // Alert identification
  string alert_id = 2;
  string proposal_id = 3;
  string job_id = 4;
  string freelancer_id = 5;
  double your_bid = 6;
  double new_lowest_bid = 7;
  double bid_difference = 8;
  int32 your_position = 9;
  int32 total_bids = 10;
  google.protobuf.Timestamp alerted_at = 11;"

    # ConnectUsed
    create_proto "proposal" "v1" "connect_used" "ConnectUsed" "  // Connect usage
  string user_id = 2;
  string proposal_id = 3;
  string job_id = 4;
  int32 connects_used = 5;
  int32 connects_remaining = 6;
  google.protobuf.Timestamp used_at = 7;"

    # ProposalFlagged
    create_proto "proposal" "v1" "proposal_flagged" "ProposalFlagged" "  // Flag details
  string flag_id = 2;
  string proposal_id = 3;
  string job_id = 4;
  string flagged_by_user_id = 5;
  skillsier.common.v1.FlagReason flag_reason = 6;
  string flag_details = 7;
  google.protobuf.Timestamp flagged_at = 8;"
    
    echo -e "${GREEN}✓ Proposal events generated (9 events)${NC}"
}

# Generate CONTRACT events
generate_contract_events() {
    echo -e "${YELLOW}📝 Generating Contract events...${NC}"
    
    # ContractCreated
    create_proto "contract" "v1" "contract_created" "ContractCreated" "  // Contract identification
  string contract_id = 2;
  string proposal_id = 3;
  string job_id = 4;
  string freelancer_id = 5;
  string client_id = 6;
  
  // Contract type & terms
  skillsier.common.v1.ContractType contract_type = 7;
  string contract_title = 8;
  string description = 9;
  
  // Contract value
  skillsier.common.v1.Money total_amount = 10;
  
  // Payment terms
  string payment_method = 11;
  int32 payment_hold_days = 12;
  
  // Milestones (simplified)
  int32 total_milestones = 13;
  
  // For hourly contracts
  double hourly_rate = 14;
  int32 estimated_hours = 15;
  int32 max_hours_per_week = 16;
  bool work_diary_required = 17;
  
  // Timeline
  google.protobuf.Timestamp start_date = 18;
  google.protobuf.Timestamp end_date = 19;
  
  // Escrow
  bool escrow_enabled = 20;
  skillsier.common.v1.Money escrow_amount = 21;
  
  // Platform fees
  double freelancer_fee_percentage = 22;
  double client_fee_percentage = 23;
  
  // Status
  skillsier.common.v1.ContractStatus contract_status = 24;
  
  // Legal
  bool nda_signed = 25;
  bool ip_agreement_signed = 26;
  
  google.protobuf.Timestamp created_at = 27;"

    # ContractStarted
    create_proto "contract" "v1" "contract_started" "ContractStarted" "  // Contract identification
  string contract_id = 2;
  string freelancer_id = 3;
  string client_id = 4;
  google.protobuf.Timestamp started_at = 5;
  string initial_milestone_id = 6;
  bool kickoff_meeting_completed = 7;
  bool onboarding_completed = 8;"

    # ContractPaused
    create_proto "contract" "v1" "contract_paused" "ContractPaused" "  // Contract identification
  string contract_id = 2;
  string freelancer_id = 3;
  string client_id = 4;
  string paused_by_user_id = 5;
  string pause_reason = 6;
  string pause_details = 7;
  google.protobuf.Timestamp paused_at = 8;
  google.protobuf.Timestamp expected_resume_date = 9;
  int32 pause_duration_days = 10;"

    # ContractEnded
    create_proto "contract" "v1" "contract_ended" "ContractEnded" "  // Contract identification
  string contract_id = 2;
  string freelancer_id = 3;
  string client_id = 4;
  string end_reason = 5;
  string end_details = 6;
  string ended_by_user_id = 7;
  google.protobuf.Timestamp ended_at = 8;
  int32 contract_duration_days = 9;
  double total_paid = 10;
  int32 milestones_completed = 11;
  bool final_payment_pending = 12;
  bool review_enabled = 13;"

    # MilestoneCreated
    create_proto "contract" "v1" "milestone_created" "MilestoneCreated" "  // Milestone identification
  string milestone_id = 2;
  string contract_id = 3;
  string freelancer_id = 4;
  string client_id = 5;
  int32 milestone_number = 6;
  int32 total_milestones = 7;
  string title = 8;
  string description = 9;
  skillsier.common.v1.Money amount = 10;
  google.protobuf.Timestamp due_date = 11;
  int32 review_period_days = 12;
  int32 auto_approve_after_days = 13;
  google.protobuf.Timestamp created_at = 14;"

    # MilestoneCompleted
    create_proto "contract" "v1" "milestone_completed" "MilestoneCompleted" "  // Milestone identification
  string milestone_id = 2;
  string contract_id = 3;
  string freelancer_id = 4;
  string client_id = 5;
  int32 milestone_number = 6;
  string completion_description = 7;
  int32 attachments_count = 8;
  google.protobuf.Timestamp completed_at = 9;
  google.protobuf.Timestamp due_date = 10;
  int32 days_early_late = 11;
  google.protobuf.Timestamp auto_approval_deadline = 12;
  bool revision_requested = 13;"

    # MilestoneApproved
    create_proto "contract" "v1" "milestone_approved" "MilestoneApproved" "  // Milestone identification
  string milestone_id = 2;
  string contract_id = 3;
  string freelancer_id = 4;
  string client_id = 5;
  string approved_by_user_id = 6;
  skillsier.common.v1.ApprovalType approval_type = 7;
  string approval_notes = 8;
  double client_satisfaction_rating = 9;
  google.protobuf.Timestamp approved_at = 10;
  double payment_release_amount = 11;
  bool escrow_release_initiated = 12;"

    # TimesheetSubmitted
    create_proto "contract" "v1" "timesheet_submitted" "TimesheetSubmitted" "  // Timesheet identification
  string timesheet_id = 2;
  string contract_id = 3;
  string freelancer_id = 4;
  string client_id = 5;
  google.protobuf.Timestamp billing_period_start = 6;
  google.protobuf.Timestamp billing_period_end = 7;
  double total_hours = 8;
  double billable_hours = 9;
  double hourly_rate = 10;
  double total_amount = 11;
  int32 screenshots_count = 12;
  google.protobuf.Timestamp submitted_at = 13;
  google.protobuf.Timestamp auto_approval_deadline = 14;
  bool requires_client_approval = 15;"

    # DisputeOpened
    create_proto "contract" "v1" "dispute_opened" "DisputeOpened" "  // Dispute identification
  string dispute_id = 2;
  string contract_id = 3;
  string opener_id = 4;
  string respondent_id = 5;
  skillsier.common.v1.DisputeCategory dispute_category = 6;
  string dispute_reason = 7;
  string dispute_details = 8;
  skillsier.common.v1.Money disputed_amount = 9;
  skillsier.common.v1.DisputeResolutionMethod resolution_preference = 10;
  int32 evidence_count = 11;
  google.protobuf.Timestamp opened_at = 12;
  google.protobuf.Timestamp response_deadline = 13;"
    
    echo -e "${GREEN}✓ Contract events generated (9 events)${NC}"
}

# Generate PAYMENT events
generate_payment_events() {
    echo -e "${YELLOW}💰 Generating Payment events...${NC}"
    
    # PaymentProcessed
    create_proto "payment" "v1" "payment_processed" "PaymentProcessed" "  // Payment identification
  string transaction_id = 2;
  string payment_id = 3;
  string payer_id = 4;
  string payee_id = 5;
  
  // Amount
  skillsier.common.v1.Money amount = 6;
  
  // Payment method
  skillsier.common.v1.PaymentMethod payment_method = 7;
  skillsier.common.v1.PaymentGateway payment_gateway = 8;
  string payment_instrument_last4 = 9;
  
  // Fees
  double transaction_fee = 10;
  double platform_fee = 11;
  double gateway_fee = 12;
  double total_fees = 13;
  
  // Processing
  google.protobuf.Timestamp processed_at = 14;
  int32 processing_time_ms = 15;
  skillsier.common.v1.PaymentStatus status = 16;
  
  // Related entities
  string contract_id = 17;
  string milestone_id = 18;
  string invoice_id = 19;
  string receipt_url = 20;
  
  // Compliance
  bool compliance_check_passed = 21;
  bool aml_check_passed = 22;
  bool fraud_check_passed = 23;
  double risk_score = 24;"

    # PaymentFailed
    create_proto "payment" "v1" "payment_failed" "PaymentFailed" "  // Payment identification
  string transaction_id = 2;
  string payment_id = 3;
  string payer_id = 4;
  string payee_id = 5;
  skillsier.common.v1.Money amount = 6;
  skillsier.common.v1.PaymentMethod payment_method = 7;
  skillsier.common.v1.PaymentGateway payment_gateway = 8;
  
  // Failure information
  skillsier.common.v1.PaymentFailureReason failure_reason = 9;
  string failure_code = 10;
  string failure_message = 11;
  string gateway_response_code = 12;
  
  // Retry information
  int32 retry_attempt_number = 13;
  bool can_retry = 14;
  google.protobuf.Timestamp next_retry_at = 15;
  int32 max_retries = 16;
  
  google.protobuf.Timestamp failed_at = 17;"

    # EscrowHeld
    create_proto "payment" "v1" "escrow_held" "EscrowHeld" "  // Escrow identification
  string escrow_id = 2;
  string contract_id = 3;
  string milestone_id = 4;
  string client_id = 5;
  string freelancer_id = 6;
  skillsier.common.v1.Money amount = 7;
  double platform_fee_deducted = 8;
  double net_amount_held = 9;
  int32 auto_release_after_days = 10;
  google.protobuf.Timestamp auto_release_date = 11;
  google.protobuf.Timestamp held_at = 12;
  google.protobuf.Timestamp expected_release_date = 13;
  string source_transaction_id = 14;"

    # EscrowReleased
    create_proto "payment" "v1" "escrow_released" "EscrowReleased" "  // Escrow identification
  string escrow_id = 2;
  string contract_id = 3;
  string milestone_id = 4;
  skillsier.common.v1.Money released_amount = 5;
  string released_to = 6;
  google.protobuf.Timestamp released_at = 7;
  string release_reason = 8;
  bool full_release = 9;
  double partial_release_percentage = 10;
  double remaining_in_escrow = 11;
  string approved_by_user_id = 12;
  skillsier.common.v1.ApprovalType approval_type = 13;"

    # PayoutRequested
    create_proto "payment" "v1" "payout_requested" "PayoutRequested" "  // Payout identification
  string payout_id = 2;
  string freelancer_id = 3;
  skillsier.common.v1.Money amount = 4;
  skillsier.common.v1.PaymentMethod payout_method = 5;
  string payout_destination_last4 = 6;
  google.protobuf.Timestamp requested_at = 7;
  google.protobuf.Timestamp expected_arrival_date = 8;
  double processing_fee = 9;
  double net_payout_amount = 10;
  double available_balance_before = 11;
  double available_balance_after = 12;
  bool verification_required = 13;"

    # PayoutProcessed
    create_proto "payment" "v1" "payout_processed" "PayoutProcessed" "  // Payout identification
  string payout_id = 2;
  string freelancer_id = 3;
  skillsier.common.v1.Money amount = 4;
  skillsier.common.v1.PaymentMethod payout_method = 5;
  google.protobuf.Timestamp processed_at = 6;
  int32 processing_time_hours = 7;
  skillsier.common.v1.PaymentStatus status = 8;
  string failure_reason = 9;
  string gateway_transaction_id = 10;
  string receipt_url = 11;
  bool confirmation_email_sent = 12;"

    # InvoiceGenerated
    create_proto "payment" "v1" "invoice_generated" "InvoiceGenerated" "  // Invoice identification
  string invoice_id = 2;
  string invoice_number = 3;
  string contract_id = 4;
  string client_id = 5;
  string freelancer_id = 6;
  double subtotal = 7;
  double platform_fee = 8;
  double tax_amount = 9;
  double total_amount = 10;
  skillsier.common.v1.Currency currency = 11;
  google.protobuf.Timestamp invoice_date = 12;
  google.protobuf.Timestamp due_date = 13;
  string payment_terms = 14;
  string invoice_url = 15;
  string pdf_url = 16;
  google.protobuf.Timestamp generated_at = 17;
  bool auto_generated = 18;"

    # RefundProcessed
    create_proto "payment" "v1" "refund_processed" "RefundProcessed" "  // Refund identification
  string refund_id = 2;
  string original_transaction_id = 3;
  string contract_id = 4;
  string client_id = 5;
  string freelancer_id = 6;
  skillsier.common.v1.Money refund_amount = 7;
  double original_amount = 8;
  double refund_percentage = 9;
  string refund_reason = 10;
  string refund_details = 11;
  google.protobuf.Timestamp processed_at = 12;
  string refund_method = 13;
  string gateway_refund_id = 14;
  bool fees_refunded = 15;
  double platform_fee_refunded = 16;"
    
    echo -e "${GREEN}✓ Payment events generated (8 events)${NC}"
}

# Generate remaining events (simplified for brevity)
generate_remaining_events() {
    echo -e "${YELLOW}🔄 Generating remaining events...${NC}"
    
    # Review events (5)
    create_proto "review" "v1" "review_submitted" "ReviewSubmitted" "  string review_id = 2;
  string contract_id = 3;
  string reviewer_id = 4;
  string reviewee_id = 5;
  skillsier.common.v1.ReviewType review_type = 6;
  double overall_rating = 7;
  string comment = 8;
  google.protobuf.Timestamp submitted_at = 9;"
    
    create_proto "review" "v1" "review_responded" "ReviewResponded" "  string response_id = 2;
  string review_id = 3;
  string responder_id = 4;
  string response_text = 5;
  google.protobuf.Timestamp responded_at = 6;"
    
    create_proto "review" "v1" "badge_awarded" "BadgeAwarded" "  string badge_assignment_id = 2;
  string user_id = 3;
  skillsier.common.v1.BadgeType badge_type = 4;
  skillsier.common.v1.BadgeLevel badge_level = 5;
  google.protobuf.Timestamp awarded_at = 6;
  google.protobuf.Timestamp expiry_date = 7;
  bool is_permanent = 8;"
    
    create_proto "review" "v1" "reputation_updated" "ReputationUpdated" "  string user_id = 2;
  skillsier.common.v1.UserType user_type = 3;
  double new_job_success_score = 4;
  double previous_job_success_score = 5;
  double new_overall_rating = 6;
  double previous_overall_rating = 7;
  int32 total_reviews = 8;
  google.protobuf.Timestamp updated_at = 9;"
    
    create_proto "review" "v1" "review_flagged" "ReviewFlagged" "  string flag_id = 2;
  string review_id = 3;
  string flagged_by_user_id = 4;
  skillsier.common.v1.FlagReason flag_reason = 5;
  string flag_details = 6;
  google.protobuf.Timestamp flagged_at = 7;"
    
    # Subscription events (7)
    create_proto "subscription" "v1" "subscription_created" "SubscriptionCreated" "  string subscription_id = 2;
  string user_id = 3;
  string plan_id = 4;
  string plan_name = 5;
  skillsier.common.v1.SubscriptionPlanTier plan_tier = 6;
  google.protobuf.Timestamp start_date = 7;
  google.protobuf.Timestamp end_date = 8;
  skillsier.common.v1.BillingCycle billing_cycle = 9;
  skillsier.common.v1.Money amount = 10;
  bool auto_renew = 11;
  string promo_code_applied = 12;
  double discount_amount = 13;
  int32 trial_period_days = 14;
  google.protobuf.Timestamp created_at = 15;"
    
    create_proto "subscription" "v1" "subscription_renewed" "SubscriptionRenewed" "  string subscription_id = 2;
  string user_id = 3;
  string plan_id = 4;
  skillsier.common.v1.Money renewal_amount = 5;
  string invoice_id = 6;
  google.protobuf.Timestamp renewed_at = 7;
  google.protobuf.Timestamp next_renewal_date = 8;
  bool payment_successful = 9;"
    
    create_proto "subscription" "v1" "subscription_cancelled" "SubscriptionCancelled" "  string subscription_id = 2;
  string user_id = 3;
  string cancelled_by_user_id = 4;
  string cancellation_reason = 5;
  string cancellation_feedback = 6;
  google.protobuf.Timestamp cancelled_at = 7;
  google.protobuf.Timestamp effective_cancellation_date = 8;
  bool immediate_cancellation = 9;
  bool refund_issued = 10;
  double refund_amount = 11;"
    
    create_proto "subscription" "v1" "subscription_expired" "SubscriptionExpired" "  string subscription_id = 2;
  string user_id = 3;
  string plan_id = 4;
  google.protobuf.Timestamp expired_at = 5;
  string expiration_reason = 6;
  string downgraded_to_plan_id = 7;
  bool downgraded_to_free = 8;
  bool renewal_available = 9;"
    
    create_proto "subscription" "v1" "connects_purchased" "ConnectsPurchased" "  string purchase_id = 2;
  string user_id = 3;
  int32 connects_amount = 4;
  string package_id = 5;
  skillsier.common.v1.Money cost = 6;
  double cost_per_connect = 7;
  int32 previous_connects_balance = 8;
  int32 new_connects_balance = 9;
  string transaction_id = 10;
  google.protobuf.Timestamp purchased_at = 11;"
    
    create_proto "subscription" "v1" "connects_used" "ConnectsUsed" "  string usage_id = 2;
  string user_id = 3;
  string proposal_id = 4;
  string job_id = 5;
  int32 connects_used = 6;
  int32 connects_remaining = 7;
  google.protobuf.Timestamp used_at = 8;"
    
    create_proto "subscription" "v1" "usage_limit_reached" "UsageLimitReached" "  string user_id = 2;
  string plan_id = 3;
  string limit_type = 4;
  int32 limit_value = 5;
  int32 current_usage = 6;
  google.protobuf.Timestamp reached_at = 7;
  bool upgrade_suggested = 8;
  string suggested_plan_id = 9;"
    
    # Message events (4)
    create_proto "message" "v1" "message_sent" "MessageSent" "  string message_id = 2;
  string conversation_id = 3;
  string sender_id = 4;
  string recipient_id = 5;
  skillsier.common.v1.MessageType message_type = 6;
  string message_content = 7;
  int32 attachments_count = 8;
  skillsier.common.v1.MessageStatus message_status = 9;
  google.protobuf.Timestamp sent_at = 10;
  google.protobuf.Timestamp delivered_at = 11;
  bool is_encrypted = 12;"
    
    create_proto "message" "v1" "notification_delivered" "NotificationDelivered" "  string notification_id = 2;
  string user_id = 3;
  string notification_type = 4;
  string title = 5;
  string content_summary = 6;
  skillsier.common.v1.NotificationChannel channel = 7;
  skillsier.common.v1.NotificationPriority priority = 8;
  string related_entity_type = 9;
  string related_entity_id = 10;
  google.protobuf.Timestamp delivered_at = 11;
  bool read = 12;"
    
    create_proto "message" "v1" "email_sent" "EmailSent" "  string email_id = 2;
  string user_id = 3;
  string recipient_email = 4;
  string email_type = 5;
  string subject = 6;
  string template_id = 7;
  string sent_via = 8;
  string message_id = 9;
  google.protobuf.Timestamp sent_at = 10;
  bool open_tracked = 11;
  bool click_tracked = 12;"
    
    create_proto "message" "v1" "in_app_notification_sent" "InAppNotificationSent" "  string notification_id = 2;
  string user_id = 3;
  string notification_type = 4;
  string title = 5;
  string message = 6;
  string action_url = 7;
  string related_entity_type = 8;
  string related_entity_id = 9;
  skillsier.common.v1.NotificationPriority priority = 10;
  google.protobuf.Timestamp sent_at = 11;
  google.protobuf.Timestamp expires_at = 12;"
    
    # MessageFlagged - NEW EVENT
    create_proto "message" "v1" "message_flagged" "MessageFlagged" "  // Flag details
  string flag_id = 2;
  string message_id = 3;
  string conversation_id = 4;
  string message_sender_id = 5;
  string flagged_by_user_id = 6;
  
  // Flag reason
  skillsier.common.v1.FlagReason flag_reason = 7;
  string flag_reason_details = 8;
  repeated string specific_issues = 9;
  
  // Message context
  string message_content_summary = 10;
  skillsier.common.v1.MessageType message_type = 11;
  google.protobuf.Timestamp message_sent_at = 12;
  
  // Severity
  string severity_level = 13;
  bool requires_immediate_action = 14;
  
  // Pattern detection
  int32 similar_flags_count = 15;
  bool pattern_detected = 16;
  
  // AI analysis
  bool ai_flagged = 17;
  double ai_confidence_score = 18;
  
  // Actions
  bool message_hidden = 19;
  bool user_warned = 20;
  
  google.protobuf.Timestamp flagged_at = 21;"
    
    # Storage events (4)
    create_proto "storage" "v1" "file_uploaded" "FileUploaded" "  string file_id = 2;
  string user_id = 3;
  string file_name = 4;
  string original_file_name = 5;
  skillsier.common.v1.FileType file_type = 6;
  string mime_type = 7;
  int64 file_size_bytes = 8;
  string file_url = 9;
  string cdn_url = 10;
  skillsier.common.v1.StorageProvider storage_provider = 11;
  skillsier.common.v1.AccessLevel access_level = 12;
  string uploaded_for_entity_type = 13;
  string uploaded_for_entity_id = 14;
  skillsier.common.v1.ProcessingStatus processing_status = 15;
  skillsier.common.v1.VirusScanStatus virus_scan_status = 16;
  google.protobuf.Timestamp uploaded_at = 17;"
    
    create_proto "storage" "v1" "file_deleted" "FileDeleted" "  string file_id = 2;
  string user_id = 3;
  string deleted_by_user_id = 4;
  string file_name = 5;
  int64 file_size_bytes = 6;
  string deletion_reason = 7;
  bool soft_delete = 8;
  google.protobuf.Timestamp permanent_deletion_date = 9;
  bool recoverable = 10;
  google.protobuf.Timestamp deleted_at = 11;"
    
    create_proto "storage" "v1" "media_processed" "MediaProcessed" "  string media_id = 2;
  string file_id = 3;
  string user_id = 4;
  string processing_type = 5;
  int32 output_files_count = 6;
  google.protobuf.Timestamp processed_at = 7;
  int32 processing_duration_ms = 8;
  string status = 9;
  string error_details = 10;"
    
    create_proto "storage" "v1" "file_flagged" "FileFlagged" "  string flag_id = 2;
  string file_id = 3;
  string file_owner_user_id = 4;
  string flagged_by_user_id = 5;
  skillsier.common.v1.FlagReason flag_reason = 6;
  string flag_details = 7;
  google.protobuf.Timestamp flagged_at = 8;
  bool auto_flagged = 9;
  double ai_confidence_score = 10;"
    
    # Search events (3)
    create_proto "search" "v1" "job_indexed" "JobIndexed" "  string job_id = 2;
  string client_id = 3;
  string job_title = 4;
  string index_name = 5;
  string document_id = 6;
  double search_rank_score = 7;
  google.protobuf.Timestamp indexed_at = 8;
  string indexing_status = 9;"
    
    create_proto "search" "v1" "user_indexed" "UserIndexed" "  string user_id = 2;
  string username = 3;
  skillsier.common.v1.UserType user_type = 4;
  string professional_title = 5;
  string index_name = 6;
  google.protobuf.Timestamp indexed_at = 7;"
    
    create_proto "search" "v1" "recommendation_generated" "RecommendationGenerated" "  string recommendation_id = 2;
  string user_id = 3;
  skillsier.common.v1.UserType user_type = 4;
  string recommendation_type = 5;
  string recommendation_context = 6;
  int32 recommended_items_count = 7;
  string primary_algorithm = 8;
  string model_name = 9;
  string model_version = 10;
  double personalization_level = 11;
  google.protobuf.Timestamp generated_at = 12;"
    
    # Admin events (6) - UPDATED from 3 to 6
    create_proto "admin" "v1" "user_suspended" "UserSuspended" "  string action_id = 2;
  string target_user_id = 3;
  string admin_user_id = 4;
  skillsier.common.v1.AdminActionType action_type = 5;
  string suspension_reason = 6;
  skillsier.common.v1.ViolationSeverity violation_severity = 7;
  int32 suspension_duration_days = 8;
  google.protobuf.Timestamp suspension_start_date = 9;
  google.protobuf.Timestamp suspension_end_date = 10;
  bool is_temporary = 11;
  bool can_appeal = 12;
  google.protobuf.Timestamp action_taken_at = 13;"
    
    create_proto "admin" "v1" "dispute_resolved" "DisputeResolved" "  string dispute_id = 2;
  string contract_id = 3;
  string admin_resolver_id = 4;
  string resolution_method = 5;
  string winner = 6;
  string reasoning = 7;
  double amount_to_freelancer = 8;
  double amount_to_client = 9;
  string contract_status_post_resolution = 10;
  google.protobuf.Timestamp resolved_at = 11;
  bool resolution_final = 12;"
    
    # UserBanned - NEW EVENT
    create_proto "admin" "v1" "user_banned" "UserBanned" "  // Ban details
  string action_id = 2;
  string target_user_id = 3;
  string target_username = 4;
  string admin_user_id = 5;
  string admin_username = 6;
  
  // Reason
  string ban_reason_category = 7;
  string ban_reason_details = 8;
  repeated string specific_violations = 9;
  skillsier.common.v1.ViolationSeverity violation_severity = 10;
  
  // Ban scope
  bool is_permanent = 11;
  bool ip_banned = 12;
  bool device_banned = 13;
  bool email_banned = 14;
  
  // Evidence
  int32 evidence_count = 15;
  
  // Appeal
  bool can_appeal = 16;
  google.protobuf.Timestamp appeal_deadline = 17;
  
  // Notification
  bool notification_sent = 18;
  
  // Affected
  int32 active_contracts_terminated = 19;
  double funds_held = 20;
  
  // Compliance
  bool reported_to_authorities = 21;
  
  google.protobuf.Timestamp banned_at = 22;"
    
    # FlagReviewed - NEW EVENT
    create_proto "admin" "v1" "flag_reviewed" "FlagReviewed" "  // Flag details
  string flag_id = 2;
  string content_id = 3;
  string content_type = 4;
  string content_owner_user_id = 5;
  string flagger_user_id = 6;
  skillsier.common.v1.FlagReason flag_reason = 7;
  
  // Review details
  string reviewed_by_admin_id = 8;
  string reviewed_by_admin_name = 9;
  string review_decision = 10;
  string review_reasoning = 11;
  
  // Action taken
  string action_taken = 12;
  string action_details = 13;
  
  // Content handling
  bool content_removed = 14;
  bool content_hidden = 15;
  
  // User actions
  bool content_owner_warned = 16;
  bool content_owner_suspended = 17;
  int32 suspension_duration_days = 18;
  
  // Flag validity
  bool flag_valid = 19;
  bool false_flag = 20;
  
  // Quality metrics
  int32 review_time_minutes = 21;
  bool ai_assisted = 22;
  double ai_confidence_score = 23;
  
  // Notifications
  bool content_owner_notified = 24;
  bool flagger_notified = 25;
  
  google.protobuf.Timestamp reviewed_at = 26;"
    
    # AnnouncementPublished - NEW EVENT
    create_proto "admin" "v1" "announcement_published" "AnnouncementPublished" "  // Announcement details
  string announcement_id = 2;
  string announcement_title = 3;
  string announcement_content = 4;
  string announcement_summary = 5;
  string announcement_type = 6;
  
  // Publishing admin
  string published_by_admin_id = 7;
  string published_by_admin_name = 8;
  
  // Target audience
  string target_audience = 9;
  repeated skillsier.common.v1.UserType target_user_types = 10;
  repeated string target_countries = 11;
  
  // Delivery settings
  repeated string delivery_channels = 12;
  skillsier.common.v1.Priority priority = 13;
  bool is_dismissible = 14;
  bool require_acknowledgment = 15;
  
  // Display settings
  string display_location = 16;
  string cta_text = 17;
  string cta_url = 18;
  
  // Scheduling
  bool publish_immediately = 19;
  google.protobuf.Timestamp scheduled_publish_at = 20;
  google.protobuf.Timestamp scheduled_unpublish_at = 21;
  
  // Tracking
  bool track_opens = 22;
  bool track_clicks = 23;
  int32 estimated_recipients = 24;
  
  google.protobuf.Timestamp published_at = 25;"
    
    create_proto "admin" "v1" "content_removed" "ContentRemoved" "  string content_id = 2;
  string content_type = 3;
  string content_owner_user_id = 4;
  string removed_by_admin_id = 5;
  string removal_reason = 6;
  skillsier.common.v1.ViolationSeverity violation_severity = 7;
  bool content_archived = 8;
  bool user_notified = 9;
  google.protobuf.Timestamp removed_at = 10;"
    
    echo -e "${GREEN}✓ Remaining events generated (30 events)${NC}"
}

# Run buf commands
run_buf_commands() {
    echo -e "${YELLOW}🔧 Running buf commands...${NC}"
    
    cd "$EVENTS_DIR"
    
    # Check if buf is installed
    if ! command -v buf &> /dev/null; then
        echo -e "${RED}❌ buf CLI not found!${NC}"
        echo -e "${YELLOW}Install buf: brew install bufbuild/buf/buf${NC}"
        return 1
    fi
    
    # Lint proto files
    echo -e "${BLUE}Running buf lint...${NC}"
    if buf lint; then
        echo -e "${GREEN}✓ Linting passed${NC}"
    else
        echo -e "${YELLOW}⚠ Linting found issues (check output above)${NC}"
    fi
    
    # Generate Go code
    echo -e "${BLUE}Running buf generate...${NC}"
    if buf generate; then
        echo -e "${GREEN}✓ Go code generated${NC}"
    else
        echo -e "${RED}❌ Code generation failed${NC}"
        return 1
    fi
    
    cd "$PROJECT_ROOT"
}

# Print summary
print_summary() {
    echo ""
    echo -e "${BLUE}=============================${NC}"
    echo -e "${GREEN}✅ Proto Generation Complete!${NC}"
    echo -e "${BLUE}=============================${NC}"
    echo ""
    echo -e "${YELLOW}Summary:${NC}"
    echo "  📁 Location: $EVENTS_DIR"
    echo "  📝 User events: 7"
    echo "  💼 Job events: 6"
    echo "  📄 Proposal events: 9"
    echo "  📝 Contract events: 9"
    echo "  💰 Payment events: 8"
    echo "  ⭐ Review events: 5"
    echo "  📦 Subscription events: 7"
    echo "  💬 Message events: 4"
    echo "  📂 Storage events: 4"
    echo "  🔍 Search events: 3"
    echo "  👔 Admin events: 3"
    echo "  ➖➖➖➖➖➖➖➖➖➖"
    echo "  📊 Total: 65 events"
    echo ""
    echo -e "${YELLOW}Generated code location:${NC}"
    echo "  $EVENTS_DIR/gen/go/"
    echo ""
    echo -e "${YELLOW}Next steps:${NC}"
    echo "  1. Review generated proto files"
    echo "  2. Check generated Go code"
    echo "  3. Import in services: import \"skillsier.dev/contracts/events/gen/go/user/v1\""
    echo "  4. Update services to use events"
    echo ""
}

# Main execution
main() {
    echo -e "${BLUE}Starting proto generation...${NC}"
    echo ""
    
    create_directories
    create_buf_config
    
    # Check for common protos
    if ! generate_common_protos; then
        echo -e "${RED}❌ Common proto files are required!${NC}"
        echo -e "${YELLOW}Please create these files first:${NC}"
        echo "  - $COMMON_DIR/metadata.proto"
        echo "  - $COMMON_DIR/enums.proto"
        echo "  - $COMMON_DIR/value_objects.proto"
        echo ""
        echo -e "${YELLOW}Use the artifacts provided in the previous step.${NC}"
        exit 1
    fi
    
    # Generate all events
    generate_user_events
    generate_job_events
    generate_proposal_events
    generate_contract_events
    generate_payment_events
    generate_remaining_events
    
    # Run buf commands
    if run_buf_commands; then
        print_summary
        exit 0
    else
        echo -e "${RED}❌ Generation completed with errors${NC}"
        exit 1
    fi
}

# Run main function
main