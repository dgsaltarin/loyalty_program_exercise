package mappers

import (
	"testing"
	"time"

	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/domain/entity"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/rest/request"
	"github.com/stretchr/testify/assert"
)

func Test_Mapper_MapCampaignRequestToCampaign(t *testing.T) {
	assert := assert.New(t)

	mapper := NewMapper()
	request := &request.CreateCampaignRequest{
		Name:               "Campaign",
		StartDate:          time.Now(),
		EndDate:            time.Now().AddDate(0, 1, 0),
		PointsMultiplier:   2.0,
		CashbackPercentage: 0.1,
	}

	campaign := mapper.MapCampaignRequestToCampaign(request)

	assert.NotNil(campaign)
	assert.Equal(request.Name, campaign.Name)
	assert.Equal(request.StartDate, campaign.StartDate)
	assert.Equal(request.EndDate, campaign.EndDate)
	assert.Equal(request.PointsMultiplier, campaign.PointsMultiplier)
	assert.Equal(request.CashbackPercentage, campaign.CashbackPercentage)
}

func Test_Mapper_MapCampaignToCampaignResponse(t *testing.T) {
	assert := assert.New(t)

	mapper := NewMapper()
	campaign := &entity.Campaign{
		ID:                 "1",
		Name:               "Campaign",
		StartDate:          time.Now(),
		EndDate:            time.Now().AddDate(0, 1, 0),
		PointsMultiplier:   2.0,
		CashbackPercentage: 0.1,
	}

	campaignResponse := mapper.MapCampaignToCampaignResponse(campaign)

	assert.NotNil(campaignResponse)
	assert.Equal(campaign.ID, campaignResponse.ID)
	assert.Equal(campaign.Name, campaignResponse.Name)
	assert.Equal(campaign.StartDate, campaignResponse.StartDate)
	assert.Equal(campaign.EndDate, campaignResponse.EndDate)
	assert.Equal(campaign.PointsMultiplier, campaignResponse.PointsMultiplier)
	assert.Equal(campaign.CashbackPercentage, campaignResponse.CashbackPercentage)
}
