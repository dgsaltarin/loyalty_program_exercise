package application

import "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/domain/entity"

type CampaignService interface {
	createCampaign(campaign *entity.Campaign) (*entity.Campaign, error)
	getCampaigns() ([]*entity.Campaign, error)
	getCampaignByID(id string) (*entity.Campaign, error)
	updateCampaign(campaign *entity.Campaign) (*entity.Campaign, error)
}
