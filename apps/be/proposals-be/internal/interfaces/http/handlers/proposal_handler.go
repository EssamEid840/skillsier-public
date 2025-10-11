package handlers

import (
	"net/http"
	"strconv"
	"proposals-be/internal/application/proposal"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProposalHandler struct {
	proposalService *proposal.Service
}

func NewProposalHandler(proposalService *proposal.Service) *ProposalHandler {
	return &ProposalHandler{proposalService: proposalService}
}

func (h *ProposalHandler) CreateProposal(c *gin.Context) {
	freelancerID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	var dto proposal.CreateProposalDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.proposalService.CreateProposal(c.Request.Context(), freelancerID, &dto)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == proposal.ErrJobAlreadySubmitted {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *ProposalHandler) ListProposals(c *gin.Context) {
	filters := &proposal.ListFilters{}
	
	if jobID := c.Query("job_id"); jobID != "" {
		if id, err := uuid.Parse(jobID); err == nil {
			filters.JobID = &id
		}
	}
	if status := c.Query("status"); status != "" {
		st := proposal.ProposalStatus(status)
		filters.Status = &st
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	result, err := h.proposalService.ListProposals(c.Request.Context(), filters, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *ProposalHandler) GetProposal(c *gin.Context) {
	proposalID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid proposal ID"})
		return
	}

	result, err := h.proposalService.GetProposal(c.Request.Context(), proposalID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == proposal.ErrProposalNotFound {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *ProposalHandler) UpdateProposal(c *gin.Context) {
	freelancerID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	proposalID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid proposal ID"})
		return
	}

	var dto proposal.UpdateProposalDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.proposalService.UpdateProposal(c.Request.Context(), proposalID, freelancerID, &dto)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == proposal.ErrProposalNotFound {
			statusCode = http.StatusNotFound
		} else if err == proposal.ErrUnauthorized {
			statusCode = http.StatusForbidden
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *ProposalHandler) WithdrawProposal(c *gin.Context) {
	freelancerID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	proposalID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid proposal ID"})
		return
	}

	if err := h.proposalService.WithdrawProposal(c.Request.Context(), proposalID, freelancerID); err != nil {
		statusCode := http.StatusInternalServerError
		if err == proposal.ErrProposalNotFound {
			statusCode = http.StatusNotFound
		} else if err == proposal.ErrUnauthorized {
			statusCode = http.StatusForbidden
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "proposal withdrawn successfully"})
}

func (h *ProposalHandler) GetMyProposals(c *gin.Context) {
	freelancerID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	result, err := h.proposalService.GetMyProposals(c.Request.Context(), freelancerID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}


// Copy router.go and main.go from jobs-be and adapt for proposals-be
// Same pattern as jobs-be, just replace "job" with "proposal"