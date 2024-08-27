package handlers

import (
	"net/http"

	services "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/application"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/mappers"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/rest/request"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/rest/response"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	mapper  mappers.Mapper
	service services.CampaignService
}

// NewHandlers creates a new Handlers instance
func NewHandlers(mapper mappers.Mapper, service services.CampaignService) *Handlers {
	return &Handlers{
		mapper:  mapper,
		service: service,
	}
}

// @Summary Create a new Campaign
// @Description Create a new Campaign
// @Tags Campaign
// @Accept json
// @Produce json
// @Router /campaigns [post]
// CreateCampaign creates a new campaign
func (h *Handlers) CreateCampaign(c *gin.Context) {
	var campaignRequest request.CreateCampaignRequest

	if err := c.BindJSON(&campaignRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	campaign := h.mapper.MapCampaignRequestToCampaign(&campaignRequest)
	createdCampaign, err := h.service.CreateCampaign(campaign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, h.mapper.MapCampaignToCampaignResponse(createdCampaign))
}

// GetCampaignsByCommerceID returns all campaigns by commerce id
// @Summary Get Campaigns by Commerce ID
// @Description Get Campaigns by Commerce ID
// @Tags Campaign
// @Accept json
// @Produce json
// @Param commerce_id path string true "Commerce ID"
// GetCampaignsByCommerceID returns all campaigns by commerce id
func (h *Handlers) GetCampaignsByCommerceID(c *gin.Context) {
	commerceID := c.Param("commerce_id")
	campaigns, err := h.service.GetCampaignByCommerceID(commerceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var campaignsResponse []*response.CampaignResponse
	for _, campaign := range campaigns {
		campaignsResponse = append(campaignsResponse, h.mapper.MapCampaignToCampaignResponse(campaign))
	}

	c.JSON(http.StatusOK, campaignsResponse)
}

// GetCampaignsByBranchID returns all campaigns by branch id
// @Summary Get Campaigns by Branch ID
// @Description Get Campaigns by Branch ID
// @Tags Campaign
// @Accept json
// @Produce json
// @Param branch_id path string true "Branch ID"
// GetCampaignsByBranchID returns all campaigns by branch id
func (h *Handlers) GetCampaignsByBranchID(c *gin.Context) {
	branchID := c.Param("branch_id")
	campaigns, err := h.service.GetCampaignsByBranchID(branchID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var campaignsResponse []*response.CampaignResponse
	for _, campaign := range campaigns {
		campaignsResponse = append(campaignsResponse, h.mapper.MapCampaignToCampaignResponse(campaign))
	}

	c.JSON(http.StatusOK, campaignsResponse)
}

// @Summary Update a Campaign
// @Description Update a Campaign
// @Tags Campaign
// @Accept json
// @Produce json
// @Router /campaigns [put]
// UpdateCampaign updates a campaigns
func (h *Handlers) UpdateCampaign(c *gin.Context) {
	var campaignRequest request.UpdateCampaignRequest

	if err := c.BindJSON(&campaignRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	campaign := h.mapper.MapUpdateCampaignRequestToCampaign(&campaignRequest)
	updatedCampaign, err := h.service.UpdateCampaign(campaign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.mapper.MapCampaignToCampaignResponse(updatedCampaign))
}
