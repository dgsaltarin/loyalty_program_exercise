package handlers

import (
	"net/http"

	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/mappers"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/rest/request"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	mapper mappers.Mapper
}

func NewHandlers(mapper mappers.Mapper) *Handlers {
	return &Handlers{
		mapper: mapper,
	}
}

// CreateCommerce creates a new commerce
func (h *Handlers) CreateCommerce(c *gin.Context) {
	request := &request.CreateCommerceRequest{}

	if err := c.BindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commerce := h.mapper.MapCreateCommerceRequestToCommerce(request)
	c.JSON(http.StatusCreated, h.mapper.MapCommerceToCreateCommerceResponse(commerce))
}
