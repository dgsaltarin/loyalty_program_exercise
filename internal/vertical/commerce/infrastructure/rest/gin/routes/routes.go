package routes

import (
	"github.com/gin-gonic/gin"
)

type commerceRoutes struct {
	group *gin.RouterGroup
}

func NewCommerceRoutes(group *gin.RouterGroup) *commerceRoutes {
	routes := &commerceRoutes{
		group: group,
	}
	routes.register()
	return routes
}

func (r *commerceRoutes) register() {
	r.group.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})
}
