package services

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/domain/entity"
)

type TransactionServices interface {
	CreateTransaction(transaction *entity.Transaction) (*entity.Transaction, error)
	GetTransactionsByCommerceID(commerceID string) ([]*entity.Transaction, error)
	GetTransactionsByUserDocument(documentType string, documentNumber string) ([]*entity.Transaction, error)
}
