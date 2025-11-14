// ==========================================
// FILE: contracts-be/internal/application/eventhandler/proposal_handler.go
// Handle events from proposals-be to create contracts
// ==========================================
package eventhandler

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"contracts-be/internal/domain/contract"
	"contracts-be/internal/domain/outbox"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProposalEventHandler struct {
	contractRepo contract.Repository
	outboxRepo   outbox.Repository
	db           *gorm.DB
}

func NewProposalEventHandler(
	contractRepo contract.Repository,
	outboxRepo outbox.Repository,
	db *gorm.DB,
) *ProposalEventHandler {
	return &ProposalEventHandler{
		contractRepo: contractRepo,
		outboxRepo:   outboxRepo,
		db:           db,
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
	case "proposal.accepted":
		return h.handleProposalAccepted(ctx, envelope.Payload)
	default:
		log.Printf("Unknown event type: %s", envelope.EventType)
	}

	return nil
}

func (h *ProposalEventHandler) handleProposalAccepted(ctx context.Context, payload json.RawMessage) error {
	var data struct {
		ProposalID   string  `json:"proposal_id"`
		JobID        string  `json:"job_id"`
		ClientID     string  `json:"client_id"`
		FreelancerID string  `json:"freelancer_id"`
		BidAmount    float64 `json:"bid_amount"`
		// Additional fields from proposal
	}

	if err := json.Unmarshal(payload, &data); err != nil {
		return err
	}

	proposalID, _ := uuid.Parse(data.ProposalID)
	jobID, _ := uuid.Parse(data.JobID)
	clientID, _ := uuid.Parse(data.ClientID)
	freelancerID, _ := uuid.Parse(data.FreelancerID)

	// Create contract
	newContract := &contract.Contract{
		JobID:        jobID,
		ProposalID:   proposalID,
		ClientID:     clientID,
		FreelancerID: freelancerID,
		Title:        "Contract for accepted proposal", // Should get from proposal
		Description:  "Auto-generated contract",
		TotalAmount:  data.BidAmount,
		Status:       contract.ContractStatusActive,
		StartDate:    time.Now(),
		Terms:        "Standard terms apply",
	}

	// Transaction: Create contract + publish event
	err := h.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := h.contractRepo.Create(ctx, newContract); err != nil {
			return err
		}

		// Create outbox event
		event, err := h.createContractEvent("contract.created", newContract)
		if err != nil {
			return err
		}

		return h.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		log.Printf("Failed to create contract from proposal %s: %v", proposalID, err)
		return err
	}

	log.Printf("Created contract %s from proposal %s", newContract.ID, proposalID)
	return nil
}

func (h *ProposalEventHandler) createContractEvent(eventType string, c *contract.Contract) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"contract_id":   c.ID.String(),
		"job_id":        c.JobID.String(),
		"proposal_id":   c.ProposalID.String(),
		"client_id":     c.ClientID.String(),
		"freelancer_id": c.FreelancerID.String(),
		"status":        string(c.Status),
	}

	payloadBytes, _ := json.Marshal(payload)
	metadata := map[string]interface{}{"source": "contracts-be"}
	metadataBytes, _ := json.Marshal(metadata)

	return &outbox.Event{
		AggregateID:   c.ID.String(),
		AggregateType: "contract",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}
