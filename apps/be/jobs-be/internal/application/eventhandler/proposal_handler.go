// ==========================================
// INTER-SERVICE EVENT HANDLERS
// ==========================================

// ==========================================
// FILE: jobs-be/internal/application/eventhandler/proposal_handler.go
// Handle events from proposals-be
// ==========================================
package eventhandler

import (
	"context"
	"encoding/json"
	"log"
	"jobs-be/internal/domain/job"
	"github.com/google/uuid"
)

type ProposalEventHandler struct {
	jobRepo job.Repository
}

func NewProposalEventHandler(jobRepo job.Repository) *ProposalEventHandler {
	return &ProposalEventHandler{
		jobRepo: jobRepo,
	}
}

func (h *ProposalEventHandler) HandleMessage(ctx context.Context, message []byte) error {
	var envelope struct {
		EventType string          `json:"event_type"`
		Payload   json.RawMessage `json:"payload"`
	}

	if err := json.Unmarshal(message, &envelope); err != nil {
		log.Printf("Failed to unmarshal event envelope: %v", err)
		return err
	}

	switch envelope.EventType {
	case "proposal.created":
		return h.handleProposalCreated(ctx, envelope.Payload)
	case "proposal.accepted":
		return h.handleProposalAccepted(ctx, envelope.Payload)
	default:
		log.Printf("Unknown event type: %s", envelope.EventType)
	}

	return nil
}

func (h *ProposalEventHandler) handleProposalCreated(ctx context.Context, payload json.RawMessage) error {
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

	// Increment proposal count
	if err := h.jobRepo.IncrementProposalCount(ctx, jobID); err != nil {
		log.Printf("Failed to increment proposal count for job %s: %v", jobID, err)
		return err
	}

	log.Printf("Incremented proposal count for job %s", jobID)
	return nil
}

func (h *ProposalEventHandler) handleProposalAccepted(ctx context.Context, payload json.RawMessage) error {
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

	// Update job status to in_progress
	j, err := h.jobRepo.GetByID(ctx, jobID)
	if err != nil {
		return err
	}

	j.Status = job.JobStatusInProgress
	if err := h.jobRepo.Update(ctx, j); err != nil {
		log.Printf("Failed to update job status for job %s: %v", jobID, err)
		return err
	}

	log.Printf("Updated job %s status to in_progress", jobID)
	return nil
}