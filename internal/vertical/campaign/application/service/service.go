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

// getCampaigns retrieves all campaigns
func (c *campaignService) GetCampaigns() ([]*entity.Campaign, error) {
	return c.repository.GetCampaigns()
}

// getCampaignByID retrieves a campaign by its ID
func (c *campaignService) GetCampaignByID(id string) (*entity.Campaign, error) {
	return c.repository.GetCampaignByID(id)
}

// updateCampaign updates a campaign
func (c *campaignService) UpdateCampaign(campaign *entity.Campaign) (*entity.Campaign, error) {
	return c.repository.UpdateCampaign(campaign)
}
