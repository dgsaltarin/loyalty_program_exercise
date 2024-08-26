package routes

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/rest/gin/handlers"
	"github.com/gin-gonic/gin"
)

type campaignRoutes struct {
	group   *gin.RouterGroup
	handler *handlers.Handlers
}

func NewCampaignRoutes(group *gin.RouterGroup, handlers *handlers.Handlers) *campaignRoutes {
	routes := &campaignRoutes{
		group:   group,
		handler: handlers,
	}
	routes.register()
	return routes
}

func (r *campaignRoutes) register() {
	r.group.GET("/check", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "This is Campaign"})
	})
	r.group.POST("/", r.handler.CreateCampaign)
	r.group.GET("/:commerce_id", r.handler.GetCampaignsByCommerceID)
	r.group.PUT("/", r.handler.UpdateCampaign)
}
