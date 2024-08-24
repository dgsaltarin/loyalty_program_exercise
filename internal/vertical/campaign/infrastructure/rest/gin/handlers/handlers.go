package handlers

import (
	"net/http"

	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/application/service"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/mappers"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/rest/request"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	mapper  *mappers.Mapper
	service *service.CampaignService
}

// NewHandlers creates a new Handlers instance
func NewHandlers(mapper *mappers.Mapper, service *service.CampaignService) *Handlers {
	return &Handlers{
		mapper:  mapper,
		service: service,
	}
}

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
