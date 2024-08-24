package gorm

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/domain/entity"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/mappers"
	"gorm.io/gorm"
)

type gormCampaignRepository struct {
	db     *gorm.DB
	mapper *mappers.Mapper
}

// NewGormCampaignRepository creates a new instance of a gormCampaignRepository
func NewGormCampaignRepository(db *gorm.DB, mapper *mappers.Mapper) *gormCampaignRepository {
	return &gormCampaignRepository{
		db:     db,
		mapper: mapper,
	}
}

// CreateCampaign creates a new campaign in the database
func (r *gormCampaignRepository) CreateCampaign(campaign *entity.Campaign) (campaign *entity.Campaign, error) {
	modelCampaign := r.mapper.MapCampaignToModelCampaign(campaign)
	if err := r.db.Create(modelCampaign).Error; err != nil {
    return nil, err
  }
	return result.Error
}

// GetCampaignByID retrieves a campaign from the database by its ID 
func (r *gormCampaignRepository) GetCampaignByID(ID string) (*entity.Campaign, error) {
  modelCampaign := &models.Campaign{}
  if err := r.db.Where("campaign_id = ?", ID).First(modelCampaign).Error; err != nil {
    return nil, err
  }
  return r.mapper.MapModelCampaignToCampaign(modelCampaign), nil
}

// GetCampaigns retrieves all campaigns from the database
func (r *gormCampaignRepository) GetCampaigns() ([]*entity.Campaign, error) {
  modelCampaigns := []*models.Campaign{}
  if err := r.db.Find(&modelCampaigns).Error; err != nil {
    return nil, err
  }
  campaigns := []*entity.Campaign{}
  for _, modelCampaign := range modelCampaigns {
    campaigns = append(campaigns, r.mapper.MapModelCampaignToCampaign(modelCampaign))
  }
  return campaigns, nil
}

// UpdateCampaign updates a campaign in the database
func (r *gormCampaignRepository) UpdateCampaign(campaign *entity.Campaign) (*entity.Campaign, error) {
  modelCampaign := r.mapper.MapCampaignToModelCampaign(campaign)
  if err := r.db.Save(modelCampaign).Error; err != nil {
    return nil, err
  }
  return campaign, nil
}
