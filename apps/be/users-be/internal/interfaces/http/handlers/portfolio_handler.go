package handlers

import (
	"net/http"
	"users-be/internal/application/portfolio"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PortfolioHandler struct {
	portfolioService *portfolio.Service
}

func NewPortfolioHandler(portfolioService *portfolio.Service) *PortfolioHandler {
	return &PortfolioHandler{portfolioService: portfolioService}
}

func (h *PortfolioHandler) Create(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := uuid.Parse(userID.(string))
	var dto portfolio.CreatePortfolioDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.portfolioService.CreatePortfolio(c.Request.Context(), uid, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (h *PortfolioHandler) GetAll(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := uuid.Parse(userID.(string))
	result, err := h.portfolioService.GetAllPortfolios(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *PortfolioHandler) Update(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := uuid.Parse(userID.(string))
	id, _ := uuid.Parse(c.Param("id"))
	var dto portfolio.UpdatePortfolioDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.portfolioService.UpdatePortfolio(c.Request.Context(), id, uid, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *PortfolioHandler) Delete(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := uuid.Parse(userID.(string))
	id, _ := uuid.Parse(c.Param("id"))
	if err := h.portfolioService.DeletePortfolio(c.Request.Context(), id, uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *PortfolioHandler) UploadImage(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := uuid.Parse(userID.(string))
	portfolioID, _ := uuid.Parse(c.Param("id"))
	var dto portfolio.UploadImageDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.portfolioService.UploadImage(c.Request.Context(), portfolioID, uid, &dto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "image uploaded"})
}

