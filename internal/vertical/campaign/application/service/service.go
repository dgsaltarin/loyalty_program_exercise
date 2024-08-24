package service

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/domain/entity"
)

type CampaignService struct{}

func NewCampaignService() *CampaignService {
	return &CampaignService{}
}

// createCampaign creates a new campaign
func (c *CampaignService) CreateCampaign(campaign *entity.Campaign) (*entity.Campaign, error) {
	return nil, nil
}

// getCampaigns retrieves all campaigns
func (c *CampaignService) GetCampaigns() ([]*entity.Campaign, error) {
	return nil, nil
}

// getCampaignByID retrieves a campaign by its ID
func (c *CampaignService) GetCampaignByID(id string) (*entity.Campaign, error) {
	return nil, nil
}

// updateCampaign updates a campaign
func (c *CampaignService) UpdateCampaign(campaign *entity.Campaign) (*entity.Campaign, error) {
	return nil, nil
}
