package mappers

import (
	"testing"

	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/domain/entity"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/rest/request"
	"github.com/stretchr/testify/assert"
)

func Test_MapCommerceRequestToCommerce(t *testing.T) {
	assert := assert.New(t)
	commerceRequest := &request.CreateCommerceRequest{
		Name:    "Test",
		Address: "Test",
		Branches: []request.CreateBranchRequest{
			{
				Name:    "Test",
				Address: "Test",
			},
		},
	}
	mapper := NewMapper()
	commerce := mapper.MapCreateCommerceRequestToCommerce(commerceRequest)
	assert.NotNil(commerce)
	assert.Equal(commerce.Name, commerceRequest.Name)
	assert.Equal(commerce.Address, commerceRequest.Address)
	assert.Equal(commerce.Branches[0].Name, commerceRequest.Branches[0].Name)
	assert.Equal(commerce.Branches[0].Address, commerceRequest.Branches[0].Address)
}

func Test_MapCommerceToCreateCommerceResponse(t *testing.T) {
	assert := assert.New(t)
	commerce := &entity.Commerce{
		ID:      "Test",
		Name:    "Test",
		Address: "Test",
		Branches: []*entity.Branch{
			{
				ID:         "Test",
				CommerceID: "Test",
				Name:       "Test",
				Address:    "Test",
			},
		},
	}
	mapper := NewMapper()
	commerceResponse := mapper.MapCommerceToCreateCommerceResponse(commerce)
	assert.NotNil(commerceResponse)
	assert.Equal(commerceResponse.ID, commerce.ID)
	assert.Equal(commerceResponse.Name, commerce.Name)
	assert.Equal(commerceResponse.Address, commerce.Address)
	assert.Equal(commerceResponse.Branches[0].ID, commerce.Branches[0].ID)
	assert.Equal(commerceResponse.Branches[0].CommerceID, commerce.Branches[0].CommerceID)
	assert.Equal(commerceResponse.Branches[0].Name, commerce.Branches[0].Name)
	assert.Equal(commerceResponse.Branches[0].Address, commerce.Branches[0].Address)
}

func Test_MapCreateBranchRequestToBranch(t *testing.T) {
	assert := assert.New(t)
	branchRequest := &request.CreateBranchRequest{
		Name:    "Test",
		Address: "Test",
	}
	mapper := NewMapper()
	branch := mapper.MapCreateBranchRequestToBranch(branchRequest)
	assert.NotNil(branch)
	assert.Equal(branch.Name, branchRequest.Name)
	assert.Equal(branch.Address, branchRequest.Address)
}
