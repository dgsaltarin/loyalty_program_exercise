package postgresql

import (
	"fmt"
	"os"

	campaignModels "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/repository/gorm/models"
	commerceModels "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/commerce/infrastructure/repository/gorm/models"
	transactionModels "github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/transaction/infrastructure/repository/gorm/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewClient() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("POSTGRESS_HOST"),
		os.Getenv("POSTGRESS_PORT"),
		os.Getenv("POSTGRESS_USER"),
		os.Getenv("POSTGRESS_PASSWORD"),
		os.Getenv("POSTGRESS_DB"),
		os.Getenv("POSTGRESS_SSL"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&commerceModels.Commerce{}, commerceModels.Branch{}, &campaignModels.Campaign{}, &transactionModels.Transaction{})

	return db, nil
}
