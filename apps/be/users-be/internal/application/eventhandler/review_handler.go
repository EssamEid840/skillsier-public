// ==========================================
// FILE: users-be/internal/application/eventhandler/review_handler.go
// Handle review events to update user ratings
// ==========================================
package eventhandler

import (
	"context"
	"encoding/json"
	"log"
	"users-be/internal/domain/freelancer"
	"github.com/google/uuid"
)

type ReviewEventHandler struct {
	freelancerRepo freelancer.Repository
}

func NewReviewEventHandler(freelancerRepo freelancer.Repository) *ReviewEventHandler {
	return &ReviewEventHandler{
		freelancerRepo: freelancerRepo,
	}
}

func (h *ReviewEventHandler) HandleMessage(ctx context.Context, message []byte) error {
	var envelope struct {
		EventType string          `json:"event_type"`
		Payload   json.RawMessage `json:"payload"`
	}

	if err := json.Unmarshal(message, &envelope); err != nil {
		log.Printf("Failed to unmarshal event envelope: %v", err)
		return err
	}

	switch envelope.EventType {
	case "review.created":
		return h.handleReviewCreated(ctx, envelope.Payload)
	default:
		log.Printf("Unknown event type: %s", envelope.EventType)
	}

	return nil
}

func (h *ReviewEventHandler) handleReviewCreated(ctx context.Context, payload json.RawMessage) error {
	var data struct {
		RevieweeID string `json:"reviewee_id"`
		Rating     int    `json:"rating"`
	}

	if err := json.Unmarshal(payload, &data); err != nil {
		return err
	}

	revieweeID, err := uuid.Parse(data.RevieweeID)
	if err != nil {
		return err
	}

	// This is a simplified version
	// In production, you would:
	// 1. Call reviews-be API to get user's average rating
	// 2. Update the freelancer profile with the new rating
	
	log.Printf("Review created for user %s with rating %d", revieweeID, data.Rating)
	
	// TODO: Implement rating update logic
	// This could involve calling reviews-be's GetUserRating endpoint
	// or maintaining a cached rating in the user service

	return nil
}
