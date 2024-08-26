package gorm

import (
	"time"

	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/domain/entity"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/domain/repository"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/mappers"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/repository/gorm/models"
	"gorm.io/gorm"
)

type gormCampaignRepository struct {
	db     *gorm.DB
	mapper mappers.Mapper
}

// NewGormCampaignRepository creates a new instance of a gormCampaignRepository
func NewGormCampaignRepository(db *gorm.DB, mapper mappers.Mapper) repository.CampaignRepository {
	return &gormCampaignRepository{
		db:     db,
		mapper: mapper,
	}
}

// CreateCampaign creates a new campaign in the database
func (r *gormCampaignRepository) CreateCampaign(campaign *entity.Campaign) (*entity.Campaign, error) {
	modelCampaign := r.mapper.MapCampaignToModelCampaign(campaign)
	modelCampaign.CreatedAt = time.Now().String()
	modelCampaign.UpdatedAt = time.Now().String()
	if err := r.db.Create(modelCampaign).Error; err != nil {
		return nil, err
	}
	return campaign, nil
}

// GetCampaignsByCommerceID retrieves all campaigns from the database by commerce id
func (r *gormCampaignRepository) GetCampaignsByCommerceID(commerceID string) ([]*entity.Campaign, error) {
	campaigns := []*models.Campaign{}
	if err := r.db.Where("commerce_id = ?", commerceID).Find(&campaigns).Error; err != nil {
		return nil, err
	}

	campaignList := []*entity.Campaign{}
	for _, modelCampaign := range campaigns {
		campaignList = append(campaignList, r.mapper.MapModelCampaignToCampaign(modelCampaign))
	}
	return campaignList, nil
}

// GetCampaignsByBranchID retrieves all campaigns from the database by branch id
func (r *gormCampaignRepository) GetCampaignsByBranchID(branchID string) ([]*entity.Campaign, error) {
	campaigns := []*models.Campaign{}
	if err := r.db.Where("branch_id = ?", branchID).Find(&campaigns).Error; err != nil {
		return nil, err
	}
	campaignList := []*entity.Campaign{}
	for _, modelCampaign := range campaigns {
		campaignList = append(campaignList, r.mapper.MapModelCampaignToCampaign(modelCampaign))
	}
	return campaignList, nil
}

// UpdateCampaign updates a campaign in the database
func (r *gormCampaignRepository) UpdateCampaign(campaign *entity.Campaign) (*entity.Campaign, error) {
	modelCampaign := r.mapper.MapCampaignToModelCampaign(campaign)
	modelCampaign.UpdatedAt = time.Now().String()
	if err := r.db.Save(modelCampaign).Error; err != nil {
		return nil, err
	}
	return campaign, nil
}
