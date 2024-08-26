package repository

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/domain/entity"
)

type CampaignRepository interface {
	CreateCampaign(campaign *entity.Campaign) (*entity.Campaign, error)
	GetCampaignsByCommerceID(commerceID string) ([]*entity.Campaign, error)
	GetCampaignsByBranchID(branchID string) ([]*entity.Campaign, error)
	UpdateCampaign(campaign *entity.Campaign) (*entity.Campaign, error)
}
