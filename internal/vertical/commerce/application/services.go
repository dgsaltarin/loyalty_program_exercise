package application

import "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/domain/entity"

type CommerceServices interface {
	CreateCommerce(request *entity.Commerce) (*entity.Commerce, error)
	GetCommerceByID(id string) (*entity.Commerce, error)
	GetCommerces() ([]*entity.Commerce, error)
	UpdateCommerce(request *entity.Commerce) (*entity.Commerce, error)
}

type BranchServices interface {
	CreateBranch(request *entity.Branch) (*entity.Branch, error)
	GetBranchByID(id string) (*entity.Branch, error)
	GetBranchesByCommerceID(commerceID string) ([]*entity.Branch, error)
	UpdateBranch(request *entity.Branch) (*entity.Branch, error)
}
