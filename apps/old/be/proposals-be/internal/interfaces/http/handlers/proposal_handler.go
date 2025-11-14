package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"proposals-be/internal/application/proposal"
)

type ProposalHandler struct {
	proposalService *proposal.Service
}

func NewProposalHandler(proposalService *proposal.Service) *ProposalHandler {
	return &ProposalHandler{proposalService: proposalService}
}

func (h *ProposalHandler) Create(c *gin.Context) {
	userID, _ := c.Get("user_id")
	freelancerID, _ := uuid.Parse(userID.(string))

	var dto proposal.CreateProposalDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.proposalService.CreateProposal(c.Request.Context(), freelancerID, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *ProposalHandler) GetMyProposals(c *gin.Context) {
	userID, _ := c.Get("user_id")
	freelancerID, _ := uuid.Parse(userID.(string))

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	result, err := h.proposalService.GetMyProposals(c.Request.Context(), freelancerID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *ProposalHandler) Withdraw(c *gin.Context) {
	userID, _ := c.Get("user_id")
	freelancerID, _ := uuid.Parse(userID.(string))

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid proposal ID"})
		return
	}

	if err := h.proposalService.WithdrawProposal(c.Request.Context(), id, freelancerID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "proposal withdrawn"})
}