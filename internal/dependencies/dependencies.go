package dependencies

import (
	"github.com/dgsaltarin/loyalty_program_excersice/internal/commons/postgresql"
	campaignService "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/application/service"
	campaignMapper "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/mappers"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/repository/gorm"
	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/rest/gin/handlers"
	campaignHandler "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/rest/gin/handlers"
	commerceService "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/application/service"
	commerMapper "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/mappers"
	gormCommerce "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/repository/gorm"
	commerceHandler "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/rest/gin/handlers"
	transactionService "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/application/service"
	transactionMapper "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/infrastructure/mappers"
	gormTransaction "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/infrastructure/repository/gorm"
	transactionHandler "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/infrastructure/rest/gin/handlers"
	"go.uber.org/dig"
)

type HandlersContainer struct {
	CampaignHandler    *handlers.Handlers
	CommerceHandler    *commerceHandler.Handlers
	TransactionHandler *transactionHandler.Handlers
}

func NewWire() *dig.Container {
	container := dig.New()

	// Inject database dependencies
	container.Provide(postgresql.NewClient)

	// Inject Campaign vertical
	container.Provide(campaignMapper.NewMapper)
	container.Provide(gorm.NewGormCampaignRepository)
	container.Provide(campaignService.NewCampaignService)
	container.Provide(campaignHandler.NewHandlers)

	// Inject Commerce vertical
	container.Provide(commerMapper.NewMapper)
	container.Provide(gormCommerce.NewGormCommerceRepository)
	container.Provide(commerceService.NewCommerceService)
	container.Provide(commerceHandler.NewHandlers)

	// Inject Dependencies for Transaction Vertical
	container.Provide(transactionMapper.NewMapper)
	container.Provide(gormTransaction.NewGormTransactionRepository)
	container.Provide(transactionService.NewTransactionService)
	container.Provide(transactionHandler.NewHandlers)

	container.Provide(func(campaignHandler *handlers.Handlers, commercehandler *commerceHandler.Handlers, transatransactionHandler *transactionHandler.Handlers) *HandlersContainer {
		return &HandlersContainer{
			CampaignHandler:    campaignHandler,
			CommerceHandler:    commercehandler,
			TransactionHandler: transatransactionHandler,
		}
	})
	return container
}
