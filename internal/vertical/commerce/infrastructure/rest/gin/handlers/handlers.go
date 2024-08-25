package handlers

import (
	"net/http"

	services "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/application"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/mappers"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/rest/request"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	mapper  mappers.Mapper
	service services.CommerceServices
}

func NewHandlers(mapper mappers.Mapper, service services.CommerceServices) *Handlers {
	return &Handlers{
		mapper:  mapper,
		service: service,
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
	commerce, err := h.service.CreateCommerce(commerce)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, h.mapper.MapCommerceToCreateCommerceResponse(commerce))
}

// CreateBranch creates a new branch
func (h *Handlers) CreateBranch(c *gin.Context) {
	request := &request.CreateBranchRequest{}

	if err := c.BindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	branch := h.mapper.MapCreateBranchRequestToBranch(request)
	c.JSON(http.StatusCreated, branch)
}
