package router

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/infrastructure/rest/gin/handlers"
	"github.com/gin-gonic/gin"
)

type transactionRouter struct {
	group   *gin.RouterGroup
	handler *handlers.Handlers
}

func NewTransactionRouter(group *gin.RouterGroup, handlers *handlers.Handlers) *transactionRouter {
	router := &transactionRouter{
		group:   group,
		handler: handlers,
	}
	router.register()
	return router
}

func (r *transactionRouter) register() {
	r.group.POST("/", r.handler.CreateTransaction)
	r.group.GET("/", r.handler.GetTransactionsByUserDocument)
}
