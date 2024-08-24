package mappers

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/domain/entity"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/repository/gorm/models"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/rest/request"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/rest/response"
)

type Mapper struct{}

func NewMapper() *Mapper {
	return &Mapper{}
}

func (m *Mapper) MapCampaignRequestToCampaign(campaignRequest *request.CreateCampaignRequest) *entity.Campaign {
	campaign := entity.NewCampaign(campaignRequest.Name,
		campaignRequest.CommerceID,
		campaignRequest.BranchID,
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
		CommerceID:         campaign.CommerceID,
		BranchID:           campaign.BranchID,
		StartDate:          campaign.StartDate,
		EndDate:            campaign.EndDate,
		PointsMultiplier:   campaign.PointsMultiplier,
		CashbackPercentage: campaign.CashbackPercentage,
	}
	return campaignResponse
}

func (m *Mapper) MapCampaignToModelCampaign(campaign *entity.Campaign) *models.Campaign {
	return &models.Campaign{
		ID:                 campaign.ID,
		Name:               campaign.Name,
		CommerceID:         campaign.CommerceID,
		BranchID:           campaign.BranchID,
		StartDate:          campaign.StartDate,
		EndDate:            campaign.EndDate,
		PointsMultiplier:   campaign.PointsMultiplier,
		CashbackPercentage: campaign.CashbackPercentage,
	}
}

func (m *Mapper) MapModelCampaignToCampaign(modelCampaign *models.Campaign) *entity.Campaign {
	return &entity.Campaign{
		ID:                 modelCampaign.ID,
		Name:               modelCampaign.Name,
		CommerceID:         modelCampaign.CommerceID,
		BranchID:           modelCampaign.BranchID,
		StartDate:          modelCampaign.StartDate,
		EndDate:            modelCampaign.EndDate,
		PointsMultiplier:   modelCampaign.PointsMultiplier,
		CashbackPercentage: modelCampaign.CashbackPercentage,
	}
}
