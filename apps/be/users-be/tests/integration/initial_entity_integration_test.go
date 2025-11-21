package integration_test

import (
	"context"
	"os"
	"fmt"
	"testing"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"skillsier.dev/apps/be/users-be/internal/application/initial_entity"
	"skillsier.dev/apps/be/users-be/internal/config"
	domain "skillsier.dev/apps/be/users-be/internal/domain/initial_entity"
	"skillsier.dev/apps/be/users-be/internal/infrastructure/persistence/postgres"
)

// TestMain sets up and tears down the test environment
func TestMain(m *testing.M) {
	// Check if integration tests should run
	if os.Getenv("INTEGRATION_TEST") != "true" {
		println("⊗ Integration tests skipped (set INTEGRATION_TEST=true to run)")
		os.Exit(0)
	}

	println("→ Running integration tests...")
	code := m.Run()
	os.Exit(code)
}

// setupTestDB creates a test database connection
func setupTestDB(t *testing.T) *gorm.DB {
	// Load configuration
	cfg := config.Load()

	// Connect to database
	db, err := postgres.NewConnection(cfg.Database)
	require.NoError(t, err, "Failed to connect to test database")

	// Run migrations
	err = postgres.AutoMigrate(db, cfg.Database)
	require.NoError(t, err, "Failed to run migrations")

	return db
}

// cleanupTestData cleans up test data from the database
func cleanupTestData(t *testing.T, db *gorm.DB) {
	// Delete all test entities
	db.Exec("DELETE FROM initial_entities WHERE name LIKE 'Test%'")
	db.Exec("DELETE FROM outbox_events WHERE aggregate_type = 'initial_entity'")
}

// TestIntegration_CreateInitialEntity tests creating an entity with outbox
func TestIntegration_CreateInitialEntity(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestData(t, db)

	// Setup
	repo := postgres.NewInitialEntityRepository(db)
	service := initial_entity.NewService(repo, db)

	dto := &initial_entity.CreateInitialEntityDTO{
		Name:        "Test Integration Entity",
		Description: "Integration test description",
		Status:      domain.StatusActive,
		OwnerID:     uuid.New(),
		Tags:        []string{"integration", "test"},
		Properties: map[string]string{
			"test_key": "test_value",
		},
	}

	// Execute
	result, err := service.Create(context.Background(), dto)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, dto.Name, result.Name)
	assert.Equal(t, dto.Description, result.Description)
	assert.NotEqual(t, uuid.Nil, result.ID)

	// Verify entity was created in database
	entity, err := repo.FindByID(context.Background(), result.ID)
	require.NoError(t, err)
	assert.Equal(t, dto.Name, entity.Name)

	// Verify outbox event was created
	var outboxCount int64
	db.Table("outbox_events").
		Where("aggregate_id = ? AND event_type = ?", result.ID.String(), "initial_entity.created.v1").
		Count(&outboxCount)
	assert.Equal(t, int64(1), outboxCount, "Outbox event should be created")
}

// TestIntegration_UpdateInitialEntity tests updating an entity with outbox
func TestIntegration_UpdateInitialEntity(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestData(t, db)

	// Setup - create entity first
	repo := postgres.NewInitialEntityRepository(db)
	service := initial_entity.NewService(repo, db)

	createDTO := &initial_entity.CreateInitialEntityDTO{
		Name:        "Test Update Entity",
		Description: "Original description",
		Status:      domain.StatusActive,
		OwnerID:     uuid.New(),
	}

	created, err := service.Create(context.Background(), createDTO)
	require.NoError(t, err)

	// Execute - update entity
	newName := "Updated Test Entity"
	newDescription := "Updated description"
	updateDTO := &initial_entity.UpdateInitialEntityDTO{
		Name:        &newName,
		Description: &newDescription,
	}

	updated, err := service.Update(context.Background(), created.ID, updateDTO)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, newName, updated.Name)
	assert.Equal(t, newDescription, updated.Description)
	assert.Greater(t, updated.Version, created.Version, "Version should be incremented")

	// Verify updated event in outbox
	var outboxCount int64
	db.Table("outbox_events").
		Where("aggregate_id = ? AND event_type = ?", created.ID.String(), "initial_entity.updated.v1").
		Count(&outboxCount)
	assert.GreaterOrEqual(t, outboxCount, int64(1), "Outbox update event should be created")
}

// TestIntegration_DeleteInitialEntity tests soft deleting an entity
func TestIntegration_DeleteInitialEntity(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestData(t, db)

	// Setup - create entity first
	repo := postgres.NewInitialEntityRepository(db)
	service := initial_entity.NewService(repo, db)

	createDTO := &initial_entity.CreateInitialEntityDTO{
		Name:        "Test Delete Entity",
		Description: "To be deleted",
		Status:      domain.StatusActive,
		OwnerID:     uuid.New(),
	}

	created, err := service.Create(context.Background(), createDTO)
	require.NoError(t, err)

	// Execute - delete entity
	err = service.Delete(context.Background(), created.ID)

	// Assert
	require.NoError(t, err)

	// Verify entity is soft-deleted
	_, err = repo.FindByID(context.Background(), created.ID)
	assert.Error(t, err, "Should not find deleted entity")

	// Verify entity exists with deleted_at set
	entity, err := repo.FindByIDWithDeleted(context.Background(), created.ID)
	require.NoError(t, err)
	assert.NotNil(t, entity.DeletedAt, "DeletedAt should be set")

	// Verify delete event in outbox
	var outboxCount int64
	db.Table("outbox_events").
		Where("aggregate_id = ? AND event_type = ?", created.ID.String(), "initial_entity.deleted.v1").
		Count(&outboxCount)
	assert.Equal(t, int64(1), outboxCount, "Outbox delete event should be created")
}

// TestIntegration_ListInitialEntities tests listing entities with pagination
func TestIntegration_ListInitialEntities(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestData(t, db)

	// Setup - create multiple entities
	repo := postgres.NewInitialEntityRepository(db)
	service := initial_entity.NewService(repo, db)

	ownerID := uuid.New()

	for i := 0; i < 5; i++ {
		dto := &initial_entity.CreateInitialEntityDTO{
			Name:        fmt.Sprintf("Test List Entity %d", i),
			Description: "List test",
			Status:      domain.StatusActive,
			OwnerID:     ownerID,
		}
		_, err := service.Create(context.Background(), dto)
		require.NoError(t, err)
	}

	// Execute - list entities
	listDTO := &initial_entity.ListInitialEntitiesDTO{
		Page:     1,
		PageSize: 3,
		OwnerID:  &ownerID,
	}

	result, err := service.List(context.Background(), listDTO)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, 3, len(result.Items), "Should return 3 items (page size)")
	assert.GreaterOrEqual(t, result.Pagination.TotalItems, int64(5), "Should have at least 5 total items")
	assert.Equal(t, 1, result.Pagination.Page)
	assert.Equal(t, 3, result.Pagination.PageSize)
}