package initial_entity_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"<module>/internal/application/initial_entity"
	domain "<module>/internal/domain/initial_entity"
)

// MockRepository is a mock implementation of the InitialEntity repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(ctx context.Context, entity *domain.InitialEntity) error {
	args := m.Called(ctx, entity)
	return args.Error(0)
}

func (m *MockRepository) CreateWithOutbox(ctx context.Context, entity *domain.InitialEntity, eventType string, eventPayload []byte, topic string) error {
	args := m.Called(ctx, entity, eventType, eventPayload, topic)
	return args.Error(0)
}

func (m *MockRepository) Update(ctx context.Context, entity *domain.InitialEntity) error {
	args := m.Called(ctx, entity)
	return args.Error(0)
}

func (m *MockRepository) UpdateWithOutbox(ctx context.Context, entity *domain.InitialEntity, eventType string, eventPayload []byte, topic string) error {
	args := m.Called(ctx, entity, eventType, eventPayload, topic)
	return args.Error(0)
}

func (m *MockRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockRepository) DeleteWithOutbox(ctx context.Context, id uuid.UUID, eventType string, eventPayload []byte, topic string) error {
	args := m.Called(ctx, id, eventType, eventPayload, topic)
	return args.Error(0)
}

func (m *MockRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.InitialEntity, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.InitialEntity), args.Error(1)
}

func (m *MockRepository) FindByIDWithDeleted(ctx context.Context, id uuid.UUID) (*domain.InitialEntity, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.InitialEntity), args.Error(1)
}

func (m *MockRepository) FindByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]*domain.InitialEntity, error) {
	args := m.Called(ctx, ownerID)
	return args.Get(0).([]*domain.InitialEntity), args.Error(1)
}

func (m *MockRepository) List(ctx context.Context, filter *domain.ListFilter) ([]*domain.InitialEntity, int64, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).([]*domain.InitialEntity), args.Get(1).(int64), args.Error(2)
}

func (m *MockRepository) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	args := m.Called(ctx, id)
	return args.Bool(0), args.Error(1)
}

func (m *MockRepository) Restore(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockRepository) HardDelete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockRepository) CountByStatus(ctx context.Context, status domain.Status) (int64, error) {
	args := m.Called(ctx, status)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockRepository) WithTx(tx interface{}) domain.Repository {
	args := m.Called(tx)
	return args.Get(0).(domain.Repository)
}

// TestCreateInitialEntity tests the Create method
func TestCreateInitialEntity(t *testing.T) {
	mockRepo := new(MockRepository)
	service := initial_entity.NewService(mockRepo, "test")

	dto := &initial_entity.CreateInitialEntityDTO{
		Name:        "Test Entity",
		Description: "Test Description",
		Status:      domain.StatusActive,
		OwnerID:     uuid.New(),
		Tags:        []string{"tag1", "tag2"},
	}

	// Setup mock expectations
	mockRepo.On("CreateWithOutbox", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Execute
	result, err := service.Create(context.Background(), dto)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, dto.Name, result.Name)
	assert.Equal(t, dto.Description, result.Description)
	mockRepo.AssertExpectations(t)
}

// TestCreateInitialEntity_ValidationError tests validation errors
func TestCreateInitialEntity_ValidationError(t *testing.T) {
	mockRepo := new(MockRepository)
	service := initial_entity.NewService(mockRepo, "test")

	dto := &initial_entity.CreateInitialEntityDTO{
		Name:        "AB", // Too short
		Description: "Test Description",
		OwnerID:     uuid.New(),
	}

	// Execute
	result, err := service.Create(context.Background(), dto)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, domain.ErrNameTooShort, err)
}

// TestGetInitialEntity tests the Get method
func TestGetInitialEntity(t *testing.T) {
	mockRepo := new(MockRepository)
	service := initial_entity.NewService(mockRepo, "test")

	entityID := uuid.New()
	expectedEntity := &domain.InitialEntity{
		ID:          entityID,
		Name:        "Test Entity",
		Description: "Test Description",
		Status:      domain.StatusActive,
		OwnerID:     uuid.New(),
	}

	// Setup mock expectations
	mockRepo.On("FindByID", mock.Anything, entityID).Return(expectedEntity, nil)

	// Execute
	result, err := service.Get(context.Background(), entityID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedEntity.ID, result.ID)
	assert.Equal(t, expectedEntity.Name, result.Name)
	mockRepo.AssertExpectations(t)
}

// TestGetInitialEntity_NotFound tests not found error
func TestGetInitialEntity_NotFound(t *testing.T) {
	mockRepo := new(MockRepository)
	service := initial_entity.NewService(mockRepo, "test")

	entityID := uuid.New()

	// Setup mock expectations
	mockRepo.On("FindByID", mock.Anything, entityID).Return(nil, domain.ErrNotFound)

	// Execute
	result, err := service.Get(context.Background(), entityID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, domain.ErrNotFound, err)
	mockRepo.AssertExpectations(t)
}

// TestUpdateInitialEntity tests the Update method
func TestUpdateInitialEntity(t *testing.T) {
	mockRepo := new(MockRepository)
	service := initial_entity.NewService(mockRepo, "test")

	entityID := uuid.New()
	existingEntity := &domain.InitialEntity{
		ID:          entityID,
		Name:        "Old Name",
		Description: "Old Description",
		Status:      domain.StatusActive,
		OwnerID:     uuid.New(),
	}

	newName := "New Name"
	dto := &initial_entity.UpdateInitialEntityDTO{
		Name: &newName,
	}

	// Setup mock expectations
	mockRepo.On("FindByID", mock.Anything, entityID).Return(existingEntity, nil)
	mockRepo.On("UpdateWithOutbox", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Execute
	result, err := service.Update(context.Background(), entityID, dto)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, newName, result.Name)
	mockRepo.AssertExpectations(t)
}

// TestDeleteInitialEntity tests the Delete method
func TestDeleteInitialEntity(t *testing.T) {
	mockRepo := new(MockRepository)
	service := initial_entity.NewService(mockRepo, "test")

	entityID := uuid.New()
	existingEntity := &domain.InitialEntity{
		ID:          entityID,
		Name:        "Test Entity",
		Description: "Test Description",
		Status:      domain.StatusActive,
		OwnerID:     uuid.New(),
	}

	// Setup mock expectations
	mockRepo.On("FindByID", mock.Anything, entityID).Return(existingEntity, nil)
	mockRepo.On("DeleteWithOutbox", mock.Anything, entityID, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Execute
	err := service.Delete(context.Background(), entityID)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}