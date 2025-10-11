package handlers

import (
	"net/http"
	"users-be/internal/application/certification"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CertificationHandler struct {
	certificationService *certification.Service
}

func NewCertificationHandler(certificationService *certification.Service) *CertificationHandler {
	return &CertificationHandler{certificationService: certificationService}
}

func (h *CertificationHandler) Create(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := uuid.Parse(userID.(string))
	var dto certification.CreateCertificationDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.certificationService.CreateCertification(c.Request.Context(), uid, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (h *CertificationHandler) GetAll(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := uuid.Parse(userID.(string))
	result, err := h.certificationService.GetAllCertifications(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *CertificationHandler) Update(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := uuid.Parse(userID.(string))
	id, _ := uuid.Parse(c.Param("id"))
	var dto certification.UpdateCertificationDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.certificationService.UpdateCertification(c.Request.Context(), id, uid, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *CertificationHandler) Delete(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := uuid.Parse(userID.(string))
	id, _ := uuid.Parse(c.Param("id"))
	if err := h.certificationService.DeleteCertification(c.Request.Context(), id, uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
