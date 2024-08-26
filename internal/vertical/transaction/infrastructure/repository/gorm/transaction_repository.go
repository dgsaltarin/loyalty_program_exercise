package gorm

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/domain/entity"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/domain/repository"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/infrastructure/mappers"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/infrastructure/repository/gorm/models"
	"gorm.io/gorm"
)

type gormTransactionRepository struct {
	gorm   *gorm.DB
	mapper mappers.Mapper
}

// NewGormTransactionRepository creates a new instance of a gormTransactionRepository
func NewGormTransactionRepository(gorm *gorm.DB, mapper mappers.Mapper) repository.TransactionRepository {
	return &gormTransactionRepository{
		gorm:   gorm,
		mapper: mapper,
	}
}

func (r *gormTransactionRepository) CreateTransaction(transaction *entity.Transaction) (*entity.Transaction, error) {
	modelTransaction := r.mapper.MapTransactionToTransactionModel(transaction)
	if err := r.gorm.Create(modelTransaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *gormTransactionRepository) GetTransactionsByCommerceID(commerceID string) ([]*entity.Transaction, error) {
	transactions := []*models.Transaction{}
	if err := r.gorm.Where("commerce_id = ?", commerceID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	transactionList := []*entity.Transaction{}
	for _, modelTransaction := range transactions {
		transactionList = append(transactionList, r.mapper.MapTransactionModelToTransaction(modelTransaction))
	}
	return transactionList, nil
}

func (r *gormTransactionRepository) GetTransactionsByUserDocument(documentType string, documentNumber string) ([]*entity.Transaction, error) {
	transactions := []*models.Transaction{}
	if err := r.gorm.Where("document_type = ? AND document_number = ?", documentType, documentNumber).Find(&transactions).Error; err != nil {
		return nil, err
	}
	transactionList := []*entity.Transaction{}
	for _, modelTransaction := range transactions {
		transactionList = append(transactionList, r.mapper.MapTransactionModelToTransaction(modelTransaction))
	}
	return transactionList, nil
}
