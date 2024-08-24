package main

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/dependencies"
	commerceRouter "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/rest/gin/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func main() {
	ginInstance := SetupGin()

	routerGroup := ginInstance.Group("/api")

	container := dependencies.NewWire()

	if err := InvokeDependencyInjectionCommerce(container, routerGroup); err != nil {
		panic(err)
	}

	ginInstance.Run(":8080")
}

func SetupGin() *gin.Engine {
	instance := gin.Default()

	return instance
}

func InvokeDependencyInjectionCommerce(container *dig.Container, routerGroup *gin.RouterGroup) error {
	return container.Invoke(func(h *dependencies.HandlersContainer) {
		commerceRouter.NewCommerceRoutes(routerGroup.Group("/commerce"))
	})
}
