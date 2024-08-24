package mappers

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/domain/entity"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/rest/request"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/rest/response"
)

type Mapper struct{}

func NewMapper() *Mapper {
	return &Mapper{}
}

func (m *Mapper) MapCampaignRequestToCampaign(campaignRequest *request.CreateCampaignRequest) *entity.Campaign {
	campaign := entity.NewCampaign(campaignRequest.Name,
		campaignRequest.StartDate,
		campaignRequest.EndDate,
		campaignRequest.PointsMultiplier,
		campaignRequest.CashbackPercentage)
	return campaign
}

func (m *Mapper) MapCampaignToCampaignResponse(campaign *entity.Campaign) *response.CampaignResponse {
	campaignResponse := &response.CampaignResponse{
		ID:                 campaign.ID,
		Name:               campaign.Name,
		StartDate:          campaign.StartDate,
		EndDate:            campaign.EndDate,
		PointsMultiplier:   campaign.PointsMultiplier,
		CashbackPercentage: campaign.CashbackPercentage,
	}
	return campaignResponse
}
