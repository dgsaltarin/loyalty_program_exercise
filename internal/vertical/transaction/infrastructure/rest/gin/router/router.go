package router

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/infrastructure/rest/gin/handlers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	group   *gin.RouterGroup
	handler *handlers.Handlers
}

func NewRouter(group *gin.RouterGroup, handlers *handlers.Handlers) *Router {
	router := &Router{
		group:   group,
		handler: handlers,
	}
	router.register()
	return router
}

func (r *Router) register() {
	r.group.POST("/", r.handler.CreateTransaction)
	r.group.GET("/", r.handler.GetTransactionsByUserDocument)
}
