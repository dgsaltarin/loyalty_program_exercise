package main

import (
	"fmt"

	_ "github.com/dgsaltarin/loyalty_program_excersice/cmd/loyalty_program_excersice/docs"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/dependencies"
	campaignRouter "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/rest/gin/routes"
	commerceRouter "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/rest/gin/routes"
	transactionRouter "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/infrastructure/rest/gin/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"
)

// @title Loyalty Program api
// @version 1.0
// @description This is a simple loyalty program api
// @host localhost:8080
// @BasePath /api/
// @schemes http
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

	if err := InvokeDependencyInjectionTransaction(container, routerGroup); err != nil {
		panic(err)
	}

	ginInstance.Run(":8080")
}

func SetupGin() *gin.Engine {
	instance := gin.Default()

	instance.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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

func InvokeDependencyInjectionTransaction(container *dig.Container, routerGroup *gin.RouterGroup) error {
	return container.Invoke(func(h *dependencies.HandlersContainer) {
		transactionRouter.NewTransactionRouter(routerGroup.Group("/transactions"), h.TransactionHandler)
	})
}
