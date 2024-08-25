package main

import (
	"fmt"

	"github.com/dgsaltarin/loyalty_program_excersice/internal/dependencies"
	campaignRouter "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/rest/gin/routes"
	commerceRouter "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/rest/gin/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/dig"
)

func main() {
	ginInstance := SetupGin()

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	routerGroup := ginInstance.Group("/api")

	container := dependencies.NewWire()

	if err := InvokeDependencyInjectionCommerce(container, routerGroup); err != nil {
		panic(err)
	}

	if err := InvokeDependencyInjectionCampaign(container, routerGroup); err != nil {
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
		commerceRouter.NewCommerceRoutes(routerGroup.Group("/commerces"), h.CommerceHandler)
	})
}

func InvokeDependencyInjectionCampaign(container *dig.Container, routerGroup *gin.RouterGroup) error {
	return container.Invoke(func(h *dependencies.HandlersContainer) {
		campaignRouter.NewCampaignRoutes(routerGroup.Group("/campaigns"), h.CampaignHandler)
	})
}
