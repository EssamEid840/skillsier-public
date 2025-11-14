package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"contracts-be/internal/application/contract"
)

type ContractHandler struct {
	contractService *contract.Service
}

func NewContractHandler(contractService *contract.Service) *ContractHandler {
	return &ContractHandler{contractService: contractService}
}

func (h *ContractHandler) GetMyContracts(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := uuid.Parse(userID.(string))
	userType := c.DefaultQuery("user_type", "freelancer")

	result, err := h.contractService.GetMyContracts(c.Request.Context(), uid, userType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"contracts": result})
}

func (h *ContractHandler) SubmitMilestone(c *gin.Context) {
	userID, _ := c.Get("user_id")
	freelancerID, _ := uuid.Parse(userID.(string))

	contractID, _ := uuid.Parse(c.Param("id"))
	milestoneID, _ := uuid.Parse(c.Param("milestone_id"))

	if err := h.contractService.SubmitMilestone(c.Request.Context(), contractID, milestoneID, freelancerID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "milestone submitted"})
}

func (h *ContractHandler) ApproveMilestone(c *gin.Context) {
	userID, _ := c.Get("user_id")
	clientID, _ := uuid.Parse(userID.(string))

	contractID, _ := uuid.Parse(c.Param("id"))
	milestoneID, _ := uuid.Parse(c.Param("milestone_id"))

	if err := h.contractService.ApproveMilestone(c.Request.Context(), contractID, milestoneID, clientID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "milestone approved"})
}

func (h *ContractHandler) CompleteContract(c *gin.Context) {
	userID, _ := c.Get("user_id")
	clientID, _ := uuid.Parse(userID.(string))

	contractID, _ := uuid.Parse(c.Param("id"))

	if err := h.contractService.CompleteContract(c.Request.Context(), contractID, clientID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "contract completed"})
}