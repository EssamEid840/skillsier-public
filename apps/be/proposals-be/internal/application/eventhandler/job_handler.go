// ==========================================
// FILE: proposals-be/internal/application/eventhandler/job_handler.go
// Handle events from jobs-be
// ==========================================
package eventhandler

import (
	"context"
	"encoding/json"
	"log"
	"proposals-be/internal/domain/proposal"
	"github.com/google/uuid"
)

type JobEventHandler struct {
	proposalRepo proposal.Repository
}

func NewJobEventHandler(proposalRepo proposal.Repository) *JobEventHandler {
	return &JobEventHandler{
		proposalRepo: proposalRepo,
	}
}

func (h *JobEventHandler) HandleMessage(ctx context.Context, message []byte) error {
	var envelope struct {
		EventType string          `json:"event_type"`
		Payload   json.RawMessage `json:"payload"`
	}

	if err := json.Unmarshal(message, &envelope); err != nil {
		log.Printf("Failed to unmarshal event envelope: %v", err)
		return err
	}

	switch envelope.EventType {
	case "job.deleted":
		return h.handleJobDeleted(ctx, envelope.Payload)
	case "job.closed":
		return h.handleJobClosed(ctx, envelope.Payload)
	default:
		log.Printf("Unknown event type: %s", envelope.EventType)
	}

	return nil
}

func (h *JobEventHandler) handleJobDeleted(ctx context.Context, payload json.RawMessage) error {
	var data struct {
		JobID string `json:"job_id"`
	}

	if err := json.Unmarshal(payload, &data); err != nil {
		return err
	}

	jobID, err := uuid.Parse(data.JobID)
	if err != nil {
		return err
	}

	// Mark all proposals for this job as withdrawn
	// Implementation would iterate through proposals and update them
	log.Printf("Handling deletion of job %s - withdrawing related proposals", jobID)
	
	// TODO: Implement batch update of proposals
	// This is a simplified version - in production, you'd want to:
	// 1. Get all pending proposals for this job
	// 2. Update their status to withdrawn
	// 3. Publish events for each withdrawal

	return nil
}

func (h *JobEventHandler) handleJobClosed(ctx context.Context, payload json.RawMessage) error {
	var data struct {
		JobID string `json:"job_id"`
	}

	if err := json.Unmarshal(payload, &data); err != nil {
		return err
	}

	jobID, err := uuid.Parse(data.JobID)
	if err != nil {
		return err
	}

	log.Printf("Handling closure of job %s - rejecting pending proposals", jobID)
	
	// Mark all pending proposals as rejected
	// Similar to handleJobDeleted but with rejected status

	return nil
}