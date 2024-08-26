package models

import (
	"time"

	"github.com/dgsaltarin/loyalty_program_excersice/internal/vertical/campaign/infrastructure/repository/gorm/models"
)

type Commerce struct {
	ID        string            `gorm:"primaryKey"`
	Name      string            `gorm:"not null"`
	Address   string            `gorm:"not null"`
	CreatedAt time.Time         `gorm:"autoCreateTime"`
	UpdatedAt time.Time         `gorm:"autoUpdateTime"`
	Branches  []Branch          `gorm:"foreignKey:CommerceID"`
	Campaigns []models.Campaign `gorm:"foreignKey:CommerceID"`
}

type Branch struct {
	ID         string            `gorm:"primaryKey"`
	CommerceID string            `gorm:"not null"`
	Name       string            `gorm:"not null"`
	Address    string            `gorm:"not null"`
	CreatedAt  time.Time         `gorm:"autoCreateTime"`
	UpdatedAt  time.Time         `gorm:"autoUpdateTime"`
	Campaigns  []models.Campaign `gorm:"foreignKey:BranchID"`
}
