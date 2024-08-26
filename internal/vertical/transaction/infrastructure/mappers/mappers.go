package mappers

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/domain/entity"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/infrastructure/repository/gorm/models"
	response "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/infrastructure/rest/reponse"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/infrastructure/rest/request"
)

type Mapper struct{}

func NewMapper() Mapper {
	return Mapper{}
}

// MapTransactionRequestToTransaction maps a Transaction to a Transaction repository model
func (m *Mapper) MapTransactionToTransactionModel(transaction *entity.Transaction) *models.Transaction {
	return &models.Transaction{
		ID:              transaction.ID,
		Amount:          transaction.Amount,
		CommerceID:      transaction.CommerceID,
		BranchID:        transaction.BranchID,
		UserID:          transaction.UserID,
		PointsErned:     transaction.PointsEarned,
		CashbackEarned:  transaction.CashbackEarned,
		TransactionDate: transaction.TransactionDate,
	}
}

// MapTransactionModelToTransaction maps a Transaction repository model to a Transaction entity
func (m *Mapper) MapTransactionModelToTransaction(transaction *models.Transaction) *entity.Transaction {
	return &entity.Transaction{
		ID:              transaction.ID,
		Amount:          transaction.Amount,
		CommerceID:      transaction.CommerceID,
		BranchID:        transaction.BranchID,
		UserID:          transaction.UserID,
		PointsEarned:    transaction.PointsErned,
		CashbackEarned:  transaction.CashbackEarned,
		TransactionDate: transaction.TransactionDate,
	}
}

func (m *Mapper) MapCrateTransactionRequestToTransaction(transaction *request.CreateTransactionRequest) *entity.Transaction {
	return &entity.Transaction{
		Amount:         transaction.Amount,
		CommerceID:     transaction.CommerceID,
		BranchID:       transaction.BranchID,
		UserID:         transaction.UserID,
		PointsEarned:   0,
		CashbackEarned: 0,
	}
}

func (m *Mapper) MapTransactionToTransactionResponse(transaction *entity.Transaction) *response.CreateTransactionResponse {
	return &response.CreateTransactionResponse{
		ID:              transaction.ID,
		Amount:          transaction.Amount,
		CommerceID:      transaction.CommerceID,
		BranchID:        transaction.BranchID,
		UserID:          transaction.UserID,
		PointsEarned:    transaction.PointsEarned,
		CashbackEarned:  transaction.CashbackEarned,
		TransactionDate: transaction.TransactionDate,
	}
}
