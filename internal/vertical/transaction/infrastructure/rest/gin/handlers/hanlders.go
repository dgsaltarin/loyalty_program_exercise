package handlers

import (
	"net/http"

	services "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/application"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/infrastructure/mappers"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/infrastructure/rest/request"
	gin "github.com/gin-gonic/gin"
)

type Handlers struct {
	transactionService services.TransactionServices
	mappers            mappers.Mapper
}

func NewHandlers(transactionService services.TransactionServices, mappers mappers.Mapper) *Handlers {
	return &Handlers{
		transactionService: transactionService,
		mappers:            mappers,
	}
}

// @Summary Create a new Transaction
// @Description Create a new Transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Router /transactions [post]
func (h *Handlers) CreateTransaction(c *gin.Context) {
	var transactionRequest request.CreateTransactionRequest

	if err := c.BindJSON(&transactionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction := h.mappers.MapCrateTransactionRequestToTransaction(&transactionRequest)
	createdTransaction, err := h.transactionService.CreateTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, h.mappers.MapTransactionToTransactionResponse(createdTransaction))
}

func (h *Handlers) GetTransactionsByUserDocument(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "This is Transaction"})
}
