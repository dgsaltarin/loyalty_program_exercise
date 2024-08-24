package repository

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/domain/entity"
)

type CampaignRepository interface {
	CreateCampaign(campaign *entity.Campaign) (*entity.Campaign, error)
	GetCampaignByID(ID string) (*entity.Campaign, error)
	GetCampaigns() ([]*entity.Campaign, error)
	UpdateCampaign(campaign *entity.Campaign) (*entity.Campaign, error)
}
