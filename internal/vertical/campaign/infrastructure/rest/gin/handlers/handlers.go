package handlers

import (
	"net/http"

	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/mappers"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/rest/request"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	mapper *mappers.Mapper
}

// NewHandlers creates a new Handlers instance
func NewHandlers(mapper *mappers.Mapper) *Handlers {
	return &Handlers{
		mapper: mapper,
	}
}

// CreateCampaign creates a new campaign
func (h *Handlers) CreateCampaign(c *gin.Context) {
	var campaignRequest request.CreateCampaignRequest

	if err := c.BindJSON(&campaignRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
