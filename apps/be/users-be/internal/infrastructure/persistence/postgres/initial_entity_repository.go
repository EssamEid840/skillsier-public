package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"reflect"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"skillsier.dev/apps/be/users-be/internal/domain/initial_entity"

	platformOutboxPostgres "skillsier.dev/platform-shared/outbox/postgres"

)

// initialEntityRepository implements the InitialEntity repository interface
type initialEntityRepository struct {
	db *gorm.DB
}

// NewInitialEntityRepository creates a new InitialEntity repository
func NewInitialEntityRepository(db *gorm.DB) *initialEntityRepository {
	return &initialEntityRepository{db: db}
}

// Create creates a new InitialEntity
func (r *initialEntityRepository) Create(ctx context.Context, entity *initial_entity.InitialEntity) error {
	if err := entity.Validate(); err != nil {
		return err
	}

	if err := r.db.WithContext(ctx).Create(entity).Error; err != nil {
		return fmt.Errorf("failed to create initial entity: %w", err)
	}

	return nil
}

// CreateWithOutbox creates a new InitialEntity and publishes an event to the outbox (atomic transaction)
func (r *initialEntityRepository) CreateWithOutbox(ctx context.Context, entity *initial_entity.InitialEntity, eventType string, eventPayload []byte, topic string) error {
	if err := entity.Validate(); err != nil {
		return err
	}

	// Execute within transaction
	return WithTxContext(ctx, r.db, func(ctx context.Context, tx *gorm.DB) error {
		// Create entity
		if err := tx.Create(entity).Error; err != nil {
			return fmt.Errorf("failed to create initial entity: %w", err)
		}

		// Create outbox event
		outboxEvent := &platformOutboxPostgres.OutboxEvent{
			ID:            uuid.New().String(),
			AggregateID:   entity.ID.String(),
			AggregateType: "initial_entity",
			EventType:     eventType,
			Payload:       eventPayload,
			Status:        "pending",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		if err := tx.Create(outboxEvent).Error; err != nil {
			return fmt.Errorf("failed to create outbox event: %w", err)
		}

		return nil
	})
}

// Update updates an existing InitialEntity
func (r *initialEntityRepository) Update(ctx context.Context, entity *initial_entity.InitialEntity) error {
	if err := entity.Validate(); err != nil {
		return err
	}

	result := r.db.WithContext(ctx).Save(entity)
	if result.Error != nil {
		return fmt.Errorf("failed to update initial entity: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return initial_entity.ErrNotFound
	}

	return nil
}

// UpdateWithOutbox updates an InitialEntity and publishes an event to the outbox (atomic transaction)
func (r *initialEntityRepository) UpdateWithOutbox(ctx context.Context, entity *initial_entity.InitialEntity, eventType string, eventPayload []byte, topic string) error {
	if err := entity.Validate(); err != nil {
		return err
	}

	return WithTxContext(ctx, r.db, func(ctx context.Context, tx *gorm.DB) error {
		// Update entity
		result := tx.Save(entity)
		if result.Error != nil {
			return fmt.Errorf("failed to update initial entity: %w", result.Error)
		}

		if result.RowsAffected == 0 {
			return initial_entity.ErrNotFound
		}

		// Create outbox event
		outboxEvent := &platformOutboxPostgres.OutboxEvent{
			ID:            uuid.New().String(),
			AggregateID:   entity.ID.String(),
			AggregateType: "initial_entity",
			EventType:     eventType,
			Payload:       eventPayload,
			Status:        "pending",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		if err := tx.Create(outboxEvent).Error; err != nil {
			return fmt.Errorf("failed to create outbox event: %w", err)
		}

		return nil
	})
}

// Delete soft-deletes an InitialEntity
func (r *initialEntityRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&initial_entity.InitialEntity{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete initial entity: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return initial_entity.ErrNotFound
	}

	return nil
}

// DeleteWithOutbox soft-deletes an InitialEntity and publishes an event to the outbox (atomic transaction)
func (r *initialEntityRepository) DeleteWithOutbox(ctx context.Context, id uuid.UUID, eventType string, eventPayload []byte, topic string) error {
	return WithTxContext(ctx, r.db, func(ctx context.Context, tx *gorm.DB) error {
		// Soft delete entity
		result := tx.Where("id = ?", id).Delete(&initial_entity.InitialEntity{})
		if result.Error != nil {
			return fmt.Errorf("failed to delete initial entity: %w", result.Error)
		}

		if result.RowsAffected == 0 {
			return initial_entity.ErrNotFound
		}

		// Create outbox event
		outboxEvent := &platformOutboxPostgres.OutboxEvent{
			ID:            uuid.New().String(),
			AggregateID:   id.String(),
			AggregateType: "initial_entity",
			EventType:     eventType,
			Payload:       eventPayload,
			Status:        "pending",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		if err := tx.Create(outboxEvent).Error; err != nil {
			return fmt.Errorf("failed to create outbox event: %w", err)
		}

		return nil
	})
}

// FindByID retrieves an InitialEntity by ID
func (r *initialEntityRepository) FindByID(ctx context.Context, id uuid.UUID) (*initial_entity.InitialEntity, error) {
	var entity initial_entity.InitialEntity
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&entity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, initial_entity.ErrNotFound
		}
		return nil, fmt.Errorf("failed to find initial entity: %w", err)
	}

	return &entity, nil
}

// FindByIDWithDeleted retrieves an InitialEntity by ID including soft-deleted entities
func (r *initialEntityRepository) FindByIDWithDeleted(ctx context.Context, id uuid.UUID) (*initial_entity.InitialEntity, error) {
	var entity initial_entity.InitialEntity
	err := r.db.WithContext(ctx).Unscoped().Where("id = ?", id).First(&entity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, initial_entity.ErrNotFound
		}
		return nil, fmt.Errorf("failed to find initial entity: %w", err)
	}

	return &entity, nil
}

// FindByOwnerID retrieves all InitialEntities owned by a specific user
func (r *initialEntityRepository) FindByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]*initial_entity.InitialEntity, error) {
	var entities []*initial_entity.InitialEntity
	err := r.db.WithContext(ctx).Where("owner_id = ?", ownerID).Find(&entities).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find entities by owner: %w", err)
	}

	return entities, nil
}

// List retrieves InitialEntities with pagination and filters
func (r *initialEntityRepository) List(ctx context.Context, filter *initial_entity.ListFilter) ([]*initial_entity.InitialEntity, int64, error) {
	// Ensure filter is not nil and provide sensible defaults instead of relying on a Validate method.
	if filter == nil {
		filter = &initial_entity.ListFilter{}
	}

	query := r.db.WithContext(ctx).Model(&initial_entity.InitialEntity{})

	// Apply filters
	if filter.Status != nil {
		query = query.Where("status = ?", *filter.Status)
	}

	if filter.OwnerID != nil {
		query = query.Where("owner_id = ?", *filter.OwnerID)
	}

	if filter.Search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+filter.Search+"%", "%"+filter.Search+"%")
	}
	if len(filter.Tags) > 0 {
		// Search for entities with any of the specified tags
		tagsJSON, _ := json.Marshal(filter.Tags)
		query = query.Where("metadata_tags ?| ARRAY[?]::text[]", tagsJSON)
	}

	// Support an optional IncludeDeleted boolean field on ListFilter without requiring it to exist.
	// Use reflection so this package does not depend on the field being present in the domain type.
	if filter != nil {
		v := reflect.ValueOf(filter)
		if v.Kind() == reflect.Ptr && !v.IsNil() {
			v = v.Elem()
			f := v.FieldByName("IncludeDeleted")
			if f.IsValid() && f.Kind() == reflect.Bool && f.Bool() {
				query = query.Unscoped()
			}
		}
	}

	// Count total
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count entities: %w", err)
	}

	// Apply sorting
	// Provide safe defaults and use reflection so this package doesn't depend on the domain having SortBy/SortOrder fields.
	sortBy := "created_at"
	sortOrder := "DESC"

	if filter != nil {
		v := reflect.ValueOf(filter)
		if v.Kind() == reflect.Ptr && !v.IsNil() {
			v = v.Elem()
			fBy := v.FieldByName("SortBy")
			if fBy.IsValid() && fBy.Kind() == reflect.String {
				if s := fBy.String(); s != "" {
					sortBy = s
				}
			}
			fOrder := v.FieldByName("SortOrder")
			if fOrder.IsValid() && fOrder.Kind() == reflect.String {
				if s := fOrder.String(); s != "" {
					sortOrder = s
				}
			}
		}
	}

	// normalize and validate sort order
	if sortOrder != "ASC" && sortOrder != "DESC" && sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "DESC"
	}

	orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)
	query = query.Order(orderClause)

	// Apply pagination
	// Determine offset and limit using reflection to avoid depending on concrete ListFilter methods.
	offset := 0
	limit := 20

	if filter != nil {
		v := reflect.ValueOf(filter)
		if v.IsValid() && !v.IsNil() {
			// Try methods first: GetOffset and GetLimit
			if m := v.MethodByName("GetOffset"); m.IsValid() && m.Type().NumIn() == 0 && m.Type().NumOut() == 1 {
				if outs := m.Call(nil); len(outs) == 1 {
					if out := outs[0]; out.IsValid() {
						switch out.Kind() {
						case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
							offset = int(out.Int())
						}
					}
				}
			}
			if m := v.MethodByName("GetLimit"); m.IsValid() && m.Type().NumIn() == 0 && m.Type().NumOut() == 1 {
				if outs := m.Call(nil); len(outs) == 1 {
					if out := outs[0]; out.IsValid() {
						switch out.Kind() {
						case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
							limit = int(out.Int())
						}
					}
				}
			}

			// Try fields: Offset, Limit, or Page/PageSize
			if v.Kind() == reflect.Ptr {
				ve := v.Elem()
				if ve.IsValid() {
					f := ve.FieldByName("Offset")
					if f.IsValid() && (f.Kind() == reflect.Int || f.Kind() == reflect.Int64) {
						offset = int(f.Int())
					}
					f = ve.FieldByName("Limit")
					if f.IsValid() && (f.Kind() == reflect.Int || f.Kind() == reflect.Int64) {
						limit = int(f.Int())
					}

					// Page & PageSize
					fp := ve.FieldByName("Page")
					fps := ve.FieldByName("PageSize")
					if fp.IsValid() && fps.IsValid() && (fp.Kind() == reflect.Int || fp.Kind() == reflect.Int64) && (fps.Kind() == reflect.Int || fps.Kind() == reflect.Int64) {
						page := int(fp.Int())
						pageSize := int(fps.Int())
						if page < 1 {
							page = 1
						}
						if pageSize > 0 {
							offset = (page - 1) * pageSize
							limit = pageSize
						}
					}
				}
			}
		}
	}

	query = query.Offset(offset).Limit(limit)

	// Execute query
	var entities []*initial_entity.InitialEntity
	if err := query.Find(&entities).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to list entities: %w", err)
	}

	return entities, total, nil
}

// Exists checks if an InitialEntity with the given ID exists
func (r *initialEntityRepository) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&initial_entity.InitialEntity{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("failed to check existence: %w", err)
	}

	return count > 0, nil
}

// Restore restores a soft-deleted InitialEntity
func (r *initialEntityRepository) Restore(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Model(&initial_entity.InitialEntity{}).Unscoped().
		Where("id = ? AND deleted_at IS NOT NULL", id).
		Update("deleted_at", nil)

	if result.Error != nil {
		return fmt.Errorf("failed to restore initial entity: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return initial_entity.ErrNotFound
	}

	return nil
}

// HardDelete permanently deletes an InitialEntity
func (r *initialEntityRepository) HardDelete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Unscoped().Where("id = ?", id).Delete(&initial_entity.InitialEntity{})
	if result.Error != nil {
		return fmt.Errorf("failed to hard delete initial entity: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return initial_entity.ErrNotFound
	}

	return nil
}

// CountByStatus counts entities by status
func (r *initialEntityRepository) CountByStatus(ctx context.Context, status initial_entity.Status) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&initial_entity.InitialEntity{}).Where("status = ?", status).Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("failed to count entities by status: %w", err)
	}

	return count, nil
}

// WithTx returns a new repository instance bound to the given transaction
func (r *initialEntityRepository) WithTx(tx *gorm.DB) initial_entity.Repository {
	return &initialEntityRepository{db: tx}
}