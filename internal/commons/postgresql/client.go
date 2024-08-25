package postgresql

import (
	"fmt"
	"os"

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

	return db, nil
}
