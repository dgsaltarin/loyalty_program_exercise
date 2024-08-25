package service

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/domain/entity"
)

type commerceService struct{}

type branchService struct{}

func NewCommerceService() *commerceService {
	return &commerceService{}
}

func NewBranchService() *branchService {
	return &branchService{}
}

func (s *commerceService) CreateCommerce(entity *entity.Commerce) (*entity.Commerce, error) {
	return nil, nil
}

func (s *commerceService) GetCommerceByID(id string) (*entity.Commerce, error) {
	return nil, nil
}

func (s *commerceService) GetCommerces() ([]*entity.Commerce, error) {
	return nil, nil
}

func (s *commerceService) UpdateCommerce(entity *entity.Commerce) (*entity.Commerce, error) {
	return nil, nil
}

func (s *branchService) CreateBranch(entity *entity.Branch) (*entity.Branch, error) {
	return nil, nil
}

func (s *branchService) GetBranchByID(id string) (*entity.Branch, error) {
	return nil, nil
}

func (s *branchService) GetBranchesByCommerceID(commerceID string) ([]*entity.Branch, error) {
	return nil, nil
}

func (s *branchService) UpdateBranch(entity *entity.Branch) (*entity.Branch, error) {
	return nil, nil
}
