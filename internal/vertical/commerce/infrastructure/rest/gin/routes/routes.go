package routes

import (
	"github.com/gin-gonic/gin"
)

type commerceRoutes struct {
	group *gin.RouterGroup
}

func NewCommerceRoutes(group *gin.RouterGroup) *commerceRoutes {
	return &commerceRoutes{group: group}
}

func (r *commerceRoutes) Register() {
	r.group.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})
}
