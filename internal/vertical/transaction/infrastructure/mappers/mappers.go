package mappers

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/domain/entity"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/infrastructure/repository/gorm/models"
)

type Mapper struct{}

func NewMapper() *Mapper {
	return &Mapper{}
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
	return &entity.NewTransaction(
		transaction.ID,
		transaction.Amount,
		transaction.CommerceID,
		transaction.BranchID,
		transaction.UserID,
		transaction.PointsErned,
		transaction.CashbackEarned,
		transaction.TransactionDate,
	)
}
