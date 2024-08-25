package routes

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/rest/gin/handlers"
	"github.com/gin-gonic/gin"
)

type commerceRoutes struct {
	group   *gin.RouterGroup
	handler *handlers.Handlers
}

func NewCommerceRoutes(group *gin.RouterGroup, handler *handlers.Handlers) *commerceRoutes {
	routes := &commerceRoutes{
		group:   group,
		handler: handler,
	}
	routes.register()
	return routes
}

func (r *commerceRoutes) register() {
	r.group.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})
	r.group.POST("/", r.handler.CreateCommerce)
}
