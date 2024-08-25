package mappers

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/domain/entity"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/rest/request"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/rest/response"
	"github.com/google/uuid"
)

func mapCreateCommerceRequestToCommerce(request *request.CreateCommerceRequest) *entity.Commerce {
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

func mapCommerceToCreateCommerceResponse(commerce *entity.Commerce) *response.CreateCommerceResponse {
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

func mapCreateBranchRequestToBranch(request *request.CreateBranchRequest) *entity.Branch {
	return &entity.Branch{
		ID:         uuid.New().String(),
		CommerceID: request.CommerceID,
		Name:       request.Name,
		Address:    request.Address,
	}
}

func mapBranchToCreateBranchResponse(branch *entity.Branch) *response.CreateBranchResponse {
	return &response.CreateBranchResponse{
		ID:         branch.ID,
		CommerceID: branch.CommerceID,
		Name:       branch.Name,
		Address:    branch.Address,
	}
}
