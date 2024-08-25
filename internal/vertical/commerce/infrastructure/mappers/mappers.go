package mappers

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/domain/entity"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/repository/gorm/models"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/rest/request"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/rest/response"
	"github.com/google/uuid"
)

type Mapper struct{}

func NewMapper() Mapper {
	return Mapper{}
}

// MapCreateCommerceRequestToCommerce maps a CreateCommerceRequest to a Commerce entity
func (m *Mapper) MapCreateCommerceRequestToCommerce(request *request.CreateCommerceRequest) *entity.Commerce {
	branches := make([]*entity.Branch, 0)
	for _, branch := range request.Branches {
		branches = append(branches, &entity.Branch{
			ID:         uuid.New().String(),
			CommerceID: "",
			Name:       branch.Name,
			Address:    branch.Address,
		})
	}
	return &entity.Commerce{
		ID:       uuid.New().String(),
		Name:     request.Name,
		Address:  request.Address,
		Branches: branches,
	}
}

// MapCommerceToCreateCommerceResponse maps a Commerce entity to a CreateCommerceResponse
func (m *Mapper) MapCommerceToCreateCommerceResponse(commerce *entity.Commerce) *response.CreateCommerceResponse {
	branches := make([]response.CreateBranchResponse, 0)
	for _, branch := range commerce.Branches {
		branches = append(branches, response.CreateBranchResponse{
			ID:         branch.ID,
			CommerceID: branch.CommerceID,
			Name:       branch.Name,
			Address:    branch.Address,
		})
	}
	return &response.CreateCommerceResponse{
		ID:       commerce.ID,
		Name:     commerce.Name,
		Address:  commerce.Address,
		Branches: branches,
	}
}

// MapCreateBranchRequestToBranch maps a CreateBranchRequest to a Branch entity
func (m *Mapper) MapCreateBranchRequestToBranch(request *request.CreateBranchRequest) *entity.Branch {
	return &entity.Branch{
		ID:         uuid.New().String(),
		CommerceID: request.CommerceID,
		Name:       request.Name,
		Address:    request.Address,
	}
}

// MapBranchToCreateBranchResponse maps a Branch entity to a CreateBranchResponse
func (m *Mapper) MapBranchToCreateBranchResponse(branch *entity.Branch) *response.CreateBranchResponse {
	return &response.CreateBranchResponse{
		ID:         branch.ID,
		CommerceID: branch.CommerceID,
		Name:       branch.Name,
		Address:    branch.Address,
	}
}

// MapCommerceToCommerModel maps a Commerce entity to a Commerce database model
func (m *Mapper) MapCommerceToCommerceModel(commerce *entity.Commerce) *models.Commerce {
	return &models.Commerce{
		ID:       commerce.ID,
		Name:     commerce.Name,
		Address:  commerce.Address,
		Branches: m.MapBranchesToBranchModel(commerce.Branches),
	}
}

// MapBranchesToBranchModel maps a slice of Branch entities to a slice of Branch database models
func (m *Mapper) MapBranchesToBranchModel(branches []*entity.Branch) []models.Branch {
	branchModels := make([]models.Branch, 0)
	for _, branch := range branches {
		branchModels = append(branchModels, models.Branch{
			ID:         branch.ID,
			CommerceID: branch.CommerceID,
			Name:       branch.Name,
			Address:    branch.Address,
		})
	}
	return branchModels
}

// MapCommerceModelToCommerce maps a Commerce database model to a Commerce entity
func (m *Mapper) MapCommerceModelToCommerce(commerceModel *models.Commerce) *entity.Commerce {
	return &entity.Commerce{
		ID:      commerceModel.ID,
		Name:    commerceModel.Name,
		Address: commerceModel.Address,
	}
}

// MapBranchModelsToBranches maps a slice of Branch database models to a slice of Branch entities
func (m *Mapper) MapBranchModelsToBranches(branchModels []*models.Branch) []*entity.Branch {
	branches := make([]*entity.Branch, 0)
	for _, branchModel := range branchModels {
		branches = append(branches, &entity.Branch{
			ID:         branchModel.ID,
			CommerceID: branchModel.CommerceID,
			Name:       branchModel.Name,
			Address:    branchModel.Address,
		})
	}
	return branches
}
