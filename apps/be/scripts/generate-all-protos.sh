#!/bin/bash

# Script to generate all remaining protobuf files for Skillsier events
# This creates 60+ proto files based on the EVENTS.md catalog

set -e

EVENTS_DIR="contracts/events"
PROTO_BASE_HEADER='syntax = "proto3";

import "google/protobuf/timestamp.proto";
'

echo "🚀 Generating all protobuf event files..."

# Create directories
mkdir -p "$EVENTS_DIR"/{user,job,proposal,contract,payment,review,subscription,message,storage,search,admin}/v1

# Function to create a proto file
create_proto() {
  local domain=$1
  local version=$2
  local event_name=$3
  local package_name=$4
  
  local file_path="$EVENTS_DIR/$domain/$version/${event_name}.proto"
  
  if [ -f "$file_path" ]; then
    echo "✓ Skipping $file_path (already exists)"
    return
  fi
  
  echo "📝 Creating $file_path"
  
  cat > "$file_path" << EOF
syntax = "proto3";

package skillsier.${domain}.${version};

option go_package = "skillsier.dev/contracts/events/gen/go/${domain}/${version};${domain}${version}";

import "google/protobuf/timestamp.proto";

// ${event_name} - Auto-generated, fill in fields from EVENTS.md
message $(echo $event_name | sed 's/_/ /g' | awk '{for(i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) tolower(substr($i,2));}1' | sed 's/ //g') {
  // Event metadata
  string event_id = 1;
  google.protobuf.Timestamp event_timestamp = 2;
  string aggregate_id = 3;
  int32 event_version = 4;
  
  // TODO: Add fields from EVENTS.md catalog
  // Refer to contracts/events/EVENTS.md for complete field definitions
}
EOF
}

# USER EVENTS
echo "📁 Creating User events..."
# Already created: user_created, user_updated, user_verified, user_suspended, user_banned, freelancer_profile_completed, client_profile_completed

# JOB EVENTS
echo "📁 Creating Job events..."
# Already created: job_posted
create_proto "job" "v1" "job_updated" "job"
create_proto "job" "v1" "job_closed" "job"
create_proto "job" "v1" "job_invitation_sent" "job"
create_proto "job" "v1" "job_removed" "job"
create_proto "job" "v1" "job_flagged" "job"

# PROPOSAL EVENTS
echo "📁 Creating Proposal events..."
create_proto "proposal" "v1" "proposal_submitted" "proposal"
create_proto "proposal" "v1" "proposal_accepted" "proposal"
create_proto "proposal" "v1" "proposal_rejected" "proposal"
create_proto "proposal" "v1" "proposal_withdrawn" "proposal"
create_proto "proposal" "v1" "bid_placed" "proposal"
create_proto "proposal" "v1" "bid_updated" "proposal"
create_proto "proposal" "v1" "outbid_alert" "proposal"
create_proto "proposal" "v1" "connect_used" "proposal"
create_proto "proposal" "v1" "proposal_flagged" "proposal"

# CONTRACT EVENTS
echo "📁 Creating Contract events..."
create_proto "contract" "v1" "contract_created" "contract"
create_proto "contract" "v1" "contract_started" "contract"
create_proto "contract" "v1" "contract_paused" "contract"
create_proto "contract" "v1" "contract_ended" "contract"
create_proto "contract" "v1" "milestone_created" "contract"
create_proto "contract" "v1" "milestone_completed" "contract"
create_proto "contract" "v1" "milestone_approved" "contract"
create_proto "contract" "v1" "timesheet_submitted" "contract"
create_proto "contract" "v1" "dispute_opened" "contract"

# PAYMENT EVENTS
echo "📁 Creating Payment events..."
create_proto "payment" "v1" "payment_processed" "payment"
create_proto "payment" "v1" "payment_failed" "payment"
create_proto "payment" "v1" "escrow_held" "payment"
create_proto "payment" "v1" "escrow_released" "payment"
create_proto "payment" "v1" "payout_requested" "payment"
create_proto "payment" "v1" "payout_processed" "payment"
create_proto "payment" "v1" "invoice_generated" "payment"
create_proto "payment" "v1" "refund_processed" "payment"

# REVIEW EVENTS
echo "📁 Creating Review events..."
create_proto "review" "v1" "review_submitted" "review"
create_proto "review" "v1" "review_responded" "review"
create_proto "review" "v1" "badge_awarded" "review"
create_proto "review" "v1" "reputation_updated" "review"
create_proto "review" "v1" "review_flagged" "review"

# SUBSCRIPTION EVENTS
echo "📁 Creating Subscription events..."
create_proto "subscription" "v1" "subscription_created" "subscription"
create_proto "subscription" "v1" "subscription_renewed" "subscription"
create_proto "subscription" "v1" "subscription_cancelled" "subscription"
create_proto "subscription" "v1" "subscription_expired" "subscription"
create_proto "subscription" "v1" "connects_purchased" "subscription"
create_proto "subscription" "v1" "connects_used" "subscription"
create_proto "subscription" "v1" "usage_limit_reached" "subscription"

# MESSAGE EVENTS
echo "📁 Creating Message events..."
create_proto "message" "v1" "message_sent" "message"
create_proto "message" "v1" "notification_delivered" "message"
create_proto "message" "v1" "email_sent"