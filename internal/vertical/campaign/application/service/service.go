package service

import (
	services "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/application"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/domain/entity"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/domain/repository"
)

type campaignService struct {
	repository repository.CampaignRepository
}

func NewCampaignService(repository repository.CampaignRepository) services.CampaignService {
	return &campaignService{
		repository: repository,
	}
}

// createCampaign creates a new campaign
func (c *campaignService) CreateCampaign(campaign *entity.Campaign) (*entity.Campaign, error) {
	return c.repository.CreateCampaign(campaign)
}

// getCampaignByCommerceID returns all campaigns by commerce id
func (c *campaignService) GetCampaignByCommerceID(commerceID string) ([]*entity.Campaign, error) {
	return c.repository.GetCampaignsByCommerceID(commerceID)
}

// getCampaignsByBranchID returns all campaigns by branch id
func (c *campaignService) GetCampaignsByBranchID(branchID string) ([]*entity.Campaign, error) {
	return c.repository.GetCampaignsByBranchID(branchID)
}

// updateCampaign updates a campaign
func (c *campaignService) UpdateCampaign(campaign *entity.Campaign) (*entity.Campaign, error) {
	return c.repository.UpdateCampaign(campaign)
}
