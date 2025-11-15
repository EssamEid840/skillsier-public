package initial_entity

import (
	"<module>/internal/domain/initial_entity"
)

// MapCreateDTOToEntity maps CreateInitialEntityDTO to InitialEntity
func MapCreateDTOToEntity(dto *CreateInitialEntityDTO) *initial_entity.InitialEntity {
	entity := &initial_entity.InitialEntity{
		Name:        dto.Name,
		Description: dto.Description,
		OwnerID:     dto.OwnerID,
		Metadata: initial_entity.Metadata{
			Tags:       dto.Tags,
			Properties: dto.Properties,
		},
	}

	// Set status if provided, otherwise use default
	if dto.Status != "" {
		entity.Status = dto.Status
	} else {
		entity.Status = initial_entity.StatusActive
	}

	return entity
}

// MapEntityToResponseDTO maps InitialEntity to InitialEntityResponseDTO
func MapEntityToResponseDTO(entity *initial_entity.InitialEntity) *InitialEntityResponseDTO {
	return &InitialEntityResponseDTO{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Status:      entity.Status,
		OwnerID:     entity.OwnerID,
		Tags:        entity.Metadata.Tags,
		Properties:  entity.Metadata.Properties,
		Version:     entity.Metadata.Version,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
	}
}

// MapListDTOToFilter maps ListInitialEntitiesDTO to domain ListFilter
func MapListDTOToFilter(dto *ListInitialEntitiesDTO) *initial_entity.ListFilter {
	return &initial_entity.ListFilter{
		Page:           dto.Page,
		PageSize:       dto.PageSize,
		Status:         dto.Status,
		OwnerID:        dto.OwnerID,
		Search:         dto.Search,
		Tags:           dto.Tags,
		SortBy:         dto.SortBy,
		SortOrder:      dto.SortOrder,
		IncludeDeleted: dto.IncludeDeleted,
	}
}