package repository

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/domain/entity"
)

type CommerceRepository interface {
	CreateCommerce(entity *entity.Commerce) (*entity.Commerce, error)
	GetCommerceByID(id string) (*entity.Commerce, error)
	GetCommerces() ([]*entity.Commerce, error)
	UpdateCommerce(entity *entity.Commerce) (*entity.Commerce, error)
}

type BranchRepository interface {
	CreateBranch(entity *entity.Branch) (*entity.Branch, error)
	GetBranchByID(id string) (*entity.Branch, error)
	GetBranchesByCommerceID(commerceID string) ([]*entity.Branch, error)
	UpdateBranch(entity *entity.Branch) (*entity.Branch, error)
}
