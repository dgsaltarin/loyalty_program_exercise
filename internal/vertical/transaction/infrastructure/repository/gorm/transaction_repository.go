package gorm

import (
	"time"

	modelsCampaign "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/repository/gorm/models"
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
	newTransaction := entity.NewTransaction(*transaction)
	modelTransaction := r.mapper.MapTransactionToTransactionModel(newTransaction)
	if err := r.gorm.Create(modelTransaction).Error; err != nil {
		return nil, err
	}
	return newTransaction, nil
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

// check if commerce has an active campaign
func (r *gormTransactionRepository) CheckCommerceHasActiveCampaign(commerceID string) (bool, error) {
	campaigns := []*modelsCampaign.Campaign{}
	if err := r.gorm.Where("commerce_id = ? AND start_date <= ? AND end_date >= ?", commerceID, time.Now(), time.Now()).Find(&campaigns).Error; err != nil {
		return false, err
	}
	return len(campaigns) > 0, nil
}

// check if branch has an active campaign
func (r *gormTransactionRepository) CheckBranchHasActiveCampaign(branchID string) (bool, error) {
	campaigns := []*modelsCampaign.Campaign{}
	if err := r.gorm.Where("branch_id = ? AND start_date <= ? AND end_date >= ?", branchID, time.Now(), time.Now()).Find(&campaigns).Error; err != nil {
		return false, err
	}
	return len(campaigns) > 0, nil
}
