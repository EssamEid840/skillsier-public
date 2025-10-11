// FILE: jobs-be/internal/application/eventhandler/proposal_handler_test.go

package eventhandler

import (
	"context"
	"encoding/json"
	"testing"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockJobRepository struct {
	mock.Mock
}

func (m *MockJobRepository) IncrementProposalCount(ctx context.Context, jobID uuid.UUID) error {
	args := m.Called(ctx, jobID)
	return args.Error(0)
}

func TestHandleProposalCreated(t *testing.T) {
	mockRepo := new(MockJobRepository)
	handler := NewProposalEventHandler(mockRepo)

	jobID := uuid.New()
	payload := map[string]interface{}{
		"job_id": jobID.String(),
	}
	payloadBytes, _ := json.Marshal(payload)

	mockRepo.On("IncrementProposalCount", mock.Anything, jobID).Return(nil)

	err := handler.handleProposalCreated(context.Background(), payloadBytes)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
*/