package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"reviews-be/internal/application/review"
)

type ReviewHandler struct {
	reviewService *review.Service
}

func NewReviewHandler(reviewService *review.Service) *ReviewHandler {
	return &ReviewHandler{reviewService: reviewService}
}

func (h *ReviewHandler) Create(c *gin.Context) {
	userID, _ := c.Get("user_id")
	reviewerID, _ := uuid.Parse(userID.(string))

	var dto review.CreateReviewDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.reviewService.CreateReview(c.Request.Context(), reviewerID, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *ReviewHandler) GetReceivedReviews(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := uuid.Parse(userID.(string))

	result, err := h.reviewService.GetReceivedReviews(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": result})
}