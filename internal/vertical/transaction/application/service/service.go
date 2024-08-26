package service

import (
	services "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/application"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/domain/entity"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/domain/repository"
)

type TransactionService struct {
	repository repository.TransactionRepository
}

func NewTransactionService(repository repository.TransactionRepository) services.TransactionServices {
	return &TransactionService{
		repository: repository,
	}
}

func (s *TransactionService) CreateTransaction(transaction *entity.Transaction) (*entity.Transaction, error) {
	return s.repository.CreateTransaction(transaction)
}

func (s *TransactionService) GetTransactionsByCommerceID(commerceID string) ([]*entity.Transaction, error) {
	return s.repository.GetTransactionsByCommerceID(commerceID)
}

func (s *TransactionService) GetTransactionsByUserDocument(documentType string, documentNumber string) ([]*entity.Transaction, error) {
	return s.repository.GetTransactionsByUserDocument(documentType, documentNumber)
}
