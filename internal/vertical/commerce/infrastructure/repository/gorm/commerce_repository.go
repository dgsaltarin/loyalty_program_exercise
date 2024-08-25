package gorm

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/domain/entity"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/domain/repository"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/mappers"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/repository/gorm/models"
	"gorm.io/gorm"
)

type CommerceRepository struct {
	db     *gorm.DB
	mapper mappers.Mapper
}

// NewGormCommerceRepository creates a new instance of a CommerceRepository
func NewGormCommerceRepository(db *gorm.DB, mapper mappers.Mapper) repository.CommerceRepository {
	return &CommerceRepository{
		db:     db,
		mapper: mapper,
	}
}

// CreateCommerce creates a new commerce in the database
func (r *CommerceRepository) CreateCommerce(commerce *entity.Commerce) (*entity.Commerce, error) {
	modelCommerce := r.mapper.MapCommerceToCommerceModel(commerce)
	if err := r.db.Create(modelCommerce).Error; err != nil {
		return nil, err
	}
	return commerce, nil
}

// GetCommerceByID retrieves a commerce from the database by its ID
func (r *CommerceRepository) GetCommerceByID(ID string) (*entity.Commerce, error) {
	modelCommerce := &models.Commerce{}
	if err := r.db.Where("commerce_id = ?", ID).First(modelCommerce).Error; err != nil {
		return nil, err
	}
	return r.mapper.MapCommerceModelToCommerce(modelCommerce), nil
}

// GetCommerces retrieves all commerces from the database
func (r *CommerceRepository) GetCommerces() ([]*entity.Commerce, error) {
	modelCommerces := []*models.Commerce{}
	if err := r.db.Find(&modelCommerces).Error; err != nil {
		return nil, err
	}
	commerces := []*entity.Commerce{}
	for _, modelCommerce := range modelCommerces {
		commerces = append(commerces, r.mapper.MapCommerceModelToCommerce(modelCommerce))
	}
	return commerces, nil
}

// UpdateCommerce updates a commerce in the database
func (r *CommerceRepository) UpdateCommerce(commerce *entity.Commerce) (*entity.Commerce, error) {
	modelCommerce := r.mapper.MapCommerceToCommerceModel(commerce)
	if err := r.db.Save(modelCommerce).Error; err != nil {
		return nil, err
	}
	return commerce, nil
}
