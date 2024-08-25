package service

import (
	services "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/application"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/domain/entity"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/domain/repository"
)

type commerceService struct {
	repository repository.CommerceRepository
}

type branchService struct {
	repository repository.BranchRepository
}

func NewCommerceService(repository repository.CommerceRepository) services.CommerceServices {
	return &commerceService{
		repository: repository,
	}
}

func NewBranchService() *branchService {
	return &branchService{}
}

func (s *commerceService) CreateCommerce(entity *entity.Commerce) (*entity.Commerce, error) {
	return s.repository.CreateCommerce(entity)
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
