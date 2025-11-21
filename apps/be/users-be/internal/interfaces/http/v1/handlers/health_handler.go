package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HealthHandler handles health check endpoints
type HealthHandler struct {
	db        *gorm.DB
	startTime time.Time
}

// NewHealthHandler creates a new HealthHandler
func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{
		db:        db,
		startTime: time.Now(),
	}
}

// Health handles GET /health - comprehensive health check
func (h *HealthHandler) Health(c *gin.Context) {
	health := gin.H{
		"status":    "ok",
		"timestamp": time.Now().Format(time.RFC3339),
		"uptime":    time.Since(h.startTime).String(),
	}

	// Check database connection
	sqlDB, err := h.db.DB()
	if err != nil {
		health["status"] = "degraded"
		health["database"] = "unavailable"
		c.JSON(http.StatusServiceUnavailable, health)
		return
	}

	if err := sqlDB.Ping(); err != nil {
		health["status"] = "degraded"
		health["database"] = "unhealthy"
		c.JSON(http.StatusServiceUnavailable, health)
		return
	}

	health["database"] = "healthy"

	// Get database stats
	stats := sqlDB.Stats()
	health["database_stats"] = gin.H{
		"open_connections": stats.OpenConnections,
		"in_use":          stats.InUse,
		"idle":            stats.Idle,
	}

	c.JSON(http.StatusOK, health)
}

// Ready handles GET /ready - Kubernetes readiness probe
func (h *HealthHandler) Ready(c *gin.Context) {
	// Check if service is ready to accept traffic
	sqlDB, err := h.db.DB()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"ready": false,
			"error": "database unavailable",
		})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"ready": false,
			"error": "database not ready",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ready": true,
	})
}

// Live handles GET /live - Kubernetes liveness probe
func (h *HealthHandler) Live(c *gin.Context) {
	// Check if service is alive (not deadlocked)
	c.JSON(http.StatusOK, gin.H{
		"alive": true,
	})
}