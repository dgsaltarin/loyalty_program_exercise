package routes

import (
	"github.com/gin-gonic/gin"
)

type campaignRoutes struct {
	group *gin.RouterGroup
}

func NewCampaignRoutes(group *gin.RouterGroup) *campaignRoutes {
	routes := &campaignRoutes{
		group: group,
	}
	routes.register()
	return routes
}

func (r *campaignRoutes) register() {
	r.group.GET("/check", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "This is Campaign"})
	})
}
