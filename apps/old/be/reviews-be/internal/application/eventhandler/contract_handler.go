// ==========================================
// FILE: reviews-be/internal/application/eventhandler/contract_handler.go
// Handle events from contracts-be
// ==========================================
package eventhandler

import (
	"context"
	"encoding/json"
	"log"
)

type ContractEventHandler struct {
	// Add dependencies if needed
}

func NewContractEventHandler() *ContractEventHandler {
	return &ContractEventHandler{}
}

func (h *ContractEventHandler) HandleMessage(ctx context.Context, message []byte) error {
	var envelope struct {
		EventType string          `json:"event_type"`
		Payload   json.RawMessage `json:"payload"`
	}

	if err := json.Unmarshal(message, &envelope); err != nil {
		log.Printf("Failed to unmarshal event envelope: %v", err)
		return err
	}

	switch envelope.EventType {
	case "contract.completed":
		return h.handleContractCompleted(ctx, envelope.Payload)
	default:
		log.Printf("Unknown event type: %s", envelope.EventType)
	}

	return nil
}

func (h *ContractEventHandler) handleContractCompleted(ctx context.Context, payload json.RawMessage) error {
	var data struct {
		ContractID   string `json:"contract_id"`
		ClientID     string `json:"client_id"`
		FreelancerID string `json:"freelancer_id"`
	}

	if err := json.Unmarshal(payload, &data); err != nil {
		return err
	}

	log.Printf("Contract %s completed - reviews can now be created", data.ContractID)
	
	// TODO: Send notification to both parties to leave reviews
	// This could be done via a notifications service

	return nil
}
