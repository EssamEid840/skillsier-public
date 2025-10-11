package handlers

import (
	"net/http"
	"strconv"
	"jobs-be/internal/application/job"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type JobHandler struct {
	jobService *job.Service
}

func NewJobHandler(jobService *job.Service) *JobHandler {
	return &JobHandler{jobService: jobService}
}

// CreateJob handles POST /jobs
func (h *JobHandler) CreateJob(c *gin.Context) {
	clientID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	var dto job.CreateJobDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.jobService.CreateJob(c.Request.Context(), clientID, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// ListJobs handles GET /jobs with filters
func (h *JobHandler) ListJobs(c *gin.Context) {
	filters := &job.ListFilters{}
	
	// Parse query parameters
	if category := c.Query("category"); category != "" {
		filters.Category = &category
	}
	if budgetType := c.Query("budget_type"); budgetType != "" {
		bt := job.BudgetType(budgetType)
		filters.BudgetType = &bt
	}
	if status := c.Query("status"); status != "" {
		st := job.JobStatus(status)
		filters.Status = &st
	}
	if level := c.Query("experience_level"); level != "" {
		filters.ExperienceLevel = &level
	}
	if search := c.Query("search"); search != "" {
		filters.SearchTerm = &search
	}
	if minBudget := c.Query("min_budget"); minBudget != "" {
		if val, err := strconv.ParseFloat(minBudget, 64); err == nil {
			filters.MinBudget = &val
		}
	}
	if maxBudget := c.Query("max_budget"); maxBudget != "" {
		if val, err := strconv.ParseFloat(maxBudget, 64); err == nil {
			filters.MaxBudget = &val
		}
	}
	if skills := c.QueryArray("skills[]"); len(skills) > 0 {
		filters.Skills = skills
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	result, err := h.jobService.ListJobs(c.Request.Context(), filters, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetJob handles GET /jobs/:id
func (h *JobHandler) GetJob(c *gin.Context) {
	jobID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job ID"})
		return
	}

	result, err := h.jobService.GetJob(c.Request.Context(), jobID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == job.ErrJobNotFound {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// UpdateJob handles PATCH /jobs/:id
func (h *JobHandler) UpdateJob(c *gin.Context) {
	clientID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	jobID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job ID"})
		return
	}

	var dto job.UpdateJobDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.jobService.UpdateJob(c.Request.Context(), jobID, clientID, &dto)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == job.ErrJobNotFound {
			statusCode = http.StatusNotFound
		} else if err == job.ErrUnauthorized {
			statusCode = http.StatusForbidden
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeleteJob handles DELETE /jobs/:id
func (h *JobHandler) DeleteJob(c *gin.Context) {
	clientID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	jobID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job ID"})
		return
	}

	if err := h.jobService.DeleteJob(c.Request.Context(), jobID, clientID); err != nil {
		statusCode := http.StatusInternalServerError
		if err == job.ErrJobNotFound {
			statusCode = http.StatusNotFound
		} else if err == job.ErrUnauthorized {
			statusCode = http.StatusForbidden
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "job deleted successfully"})
}

// GetMyJobs handles GET /jobs/my-jobs
func (h *JobHandler) GetMyJobs(c *gin.Context) {
	clientID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	result, err := h.jobService.GetMyJobs(c.Request.Context(), clientID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}