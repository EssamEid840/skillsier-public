package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"<module>/internal/application/initial_entity"
	domain "<module>/internal/domain/initial_entity"
)

// InitialEntityHandler handles HTTP requests for InitialEntity
type InitialEntityHandler struct {
	service *initial_entity.Service
}

// NewInitialEntityHandler creates a new InitialEntityHandler
func NewInitialEntityHandler(service *initial_entity.Service) *InitialEntityHandler {
	return &InitialEntityHandler{
		service: service,
	}
}

// Create handles POST /v1/initial-entities
func (h *InitialEntityHandler) Create(c *gin.Context) {
	var dto initial_entity.CreateInitialEntityDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, initial_entity.ErrorResponse{
			Error:   "invalid_request",
			Message: "Invalid request body",
			Details: map[string]interface{}{"error": err.Error()},
		})
		return
	}

	// Create entity
	response, err := h.service.Create(c.Request.Context(), &dto)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

// Get handles GET /v1/initial-entities/:id
func (h *InitialEntityHandler) Get(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, initial_entity.ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid entity ID format",
		})
		return
	}

	response, err := h.service.Get(c.Request.Context(), id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// List handles GET /v1/initial-entities
func (h *InitialEntityHandler) List(c *gin.Context) {
	var dto initial_entity.ListInitialEntitiesDTO

	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(http.StatusBadRequest, initial_entity.ErrorResponse{
			Error:   "invalid_request",
			Message: "Invalid query parameters",
			Details: map[string]interface{}{"error": err.Error()},
		})
		return
	}

	response, err := h.service.List(c.Request.Context(), &dto)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// Update handles PUT /v1/initial-entities/:id
func (h *InitialEntityHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, initial_entity.ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid entity ID format",
		})
		return
	}

	var dto initial_entity.UpdateInitialEntityDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, initial_entity.ErrorResponse{
			Error:   "invalid_request",
			Message: "Invalid request body",
			Details: map[string]interface{}{"error": err.Error()},
		})
		return
	}

	response, err := h.service.Update(c.Request.Context(), id, &dto)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// Delete handles DELETE /v1/initial-entities/:id
func (h *InitialEntityHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, initial_entity.ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid entity ID format",
		})
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, initial_entity.SuccessResponse{
		Success: true,
		Message: "Entity deleted successfully",
	})
}

// Restore handles POST /v1/initial-entities/:id/restore
func (h *InitialEntityHandler) Restore(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, initial_entity.ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid entity ID format",
		})
		return
	}

	response, err := h.service.Restore(c.Request.Context(), id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetByOwner handles GET /v1/initial-entities/owner/:owner_id
func (h *InitialEntityHandler) GetByOwner(c *gin.Context) {
	ownerIDParam := c.Param("owner_id")
	ownerID, err := uuid.Parse(ownerIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, initial_entity.ErrorResponse{
			Error:   "invalid_owner_id",
			Message: "Invalid owner ID format",
		})
		return
	}

	response, err := h.service.GetByOwner(c.Request.Context(), ownerID)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": response,
		"count": len(response),
	})
}

// handleError handles errors and returns appropriate HTTP responses
func (h *InitialEntityHandler) handleError(c *gin.Context, err error) {
	// Check for domain errors
	if errors.Is(err, domain.ErrNotFound) {
		c.JSON(http.StatusNotFound, initial_entity.ErrorResponse{
			Error:   "not_found",
			Message: "Entity not found",
			Code:    domain.ErrCodeNotFound,
		})
		return
	}

	if errors.Is(err, domain.ErrAlreadyExists) {
		c.JSON(http.StatusConflict, initial_entity.ErrorResponse{
			Error:   "already_exists",
			Message: "Entity already exists",
			Code:    domain.ErrCodeAlreadyExists,
		})
		return
	}

	if errors.Is(err, domain.ErrInvalidStatusTransition) {
		c.JSON(http.StatusBadRequest, initial_entity.ErrorResponse{
			Error:   "invalid_status_transition",
			Message: "Invalid status transition",
			Code:    domain.ErrCodeInvalidTransition,
		})
		return
	}

	if errors.Is(err, domain.ErrCannotModifyDeleted) {
		c.JSON(http.StatusBadRequest, initial_entity.ErrorResponse{
			Error:   "cannot_modify_deleted",
			Message: "Cannot modify deleted entity",
			Code:    domain.ErrCodeDeleted,
		})
		return
	}

	if errors.Is(err, domain.ErrCannotModifyArchived) {
		c.JSON(http.StatusBadRequest, initial_entity.ErrorResponse{
			Error:   "cannot_modify_archived",
			Message: "Cannot modify archived entity",
			Code:    domain.ErrCodeDeleted,
		})
		return
	}

	// Check for validation errors
	if domain.IsValidationError(err) {
		c.JSON(http.StatusBadRequest, initial_entity.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
			Code:    domain.ErrCodeValidation,
		})
		return
	}

	// Generic error
	c.JSON(http.StatusInternalServerError, initial_entity.ErrorResponse{
		Error:   "internal_error",
		Message: "An internal error occurred",
		Code:    domain.ErrCodeInternal,
	})
}