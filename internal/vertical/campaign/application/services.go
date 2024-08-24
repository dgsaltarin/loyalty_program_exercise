package application

import "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/domain/entity"

type CampaignService interface {
	CreateCampaign(campaign *entity.Campaign) (*entity.Campaign, error)
	GetCampaigns() ([]*entity.Campaign, error)
	GetCampaignByID(id string) (*entity.Campaign, error)
	UpdateCampaign(campaign *entity.Campaign) (*entity.Campaign, error)
}
