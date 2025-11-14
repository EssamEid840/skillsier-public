package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HealthHandler handles health check requests
type HealthHandler struct {
	db *gorm.DB
}

// NewHealthHandler creates a new health handler
func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

// HealthCheck handles GET /health
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	// Check database connectivity
	sqlDB, err := h.db.DB()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unhealthy",
			"error":  "Failed to get database connection",
			"timestamp": time.Now().UTC(),
		})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unhealthy",
			"error":  "Database ping failed",
			"timestamp": time.Now().UTC(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"database":  "connected",
		"timestamp": time.Now().UTC(),
	})
}

// ReadinessCheck handles GET /ready
func (h *HealthHandler) ReadinessCheck(c *gin.Context) {
	// For Kubernetes readiness probes
	// Check if the service is ready to accept traffic
	c.JSON(http.StatusOK, gin.H{
		"status": "ready",
		"timestamp": time.Now().UTC(),
	})
}

// LivenessCheck handles GET /live
func (h *HealthHandler) LivenessCheck(c *gin.Context) {
	// For Kubernetes liveness probes
	// Check if the service is alive
	c.JSON(http.StatusOK, gin.H{
		"status": "alive",
		"timestamp": time.Now().UTC(),
	})
}