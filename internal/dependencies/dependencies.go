package dependencies

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/commons/postgresql"
	campaignService "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/application/service"
	campaignMapper "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/mappers"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/repository/gorm"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/rest/gin/handlers"
	campaignHandler "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/rest/gin/handlers"
	"go.uber.org/dig"
)

type HandlersContainer struct {
	CampaignHandler *handlers.Handlers
}

func NewWire() *dig.Container {
	container := dig.New()

	// Inject database dependencies
	container.Provide(postgresql.NewClient)

	container.Provide(campaignMapper.NewMapper)
	container.Provide(gorm.NewGormCampaignRepository)
	container.Provide(campaignService.NewCampaignService)
	container.Provide(campaignHandler.NewHandlers)

	container.Provide(func(campaignHandler *handlers.Handlers) *HandlersContainer {
		return &HandlersContainer{
			CampaignHandler: campaignHandler,
		}
	})
	return container
}
