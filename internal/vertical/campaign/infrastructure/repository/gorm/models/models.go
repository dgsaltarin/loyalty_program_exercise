package models

import (
	"time"

	"gorm.io/gorm"
)

type Campaign struct {
	ID                 string    `gorm:"primaryKey;colum:campaign_id"`
	Name               string    `gorm:"column:name"`
	CommerceID         string    `gorm:"column:commerce_id"`
	BranchID           string    `gorm:"column:branch_id"`
	StartDate          time.Time `gorm:"column:start_date"`
	EndDate            time.Time `gorm:"column:end_date"`
	PointsMultiplier   float64   `gorm:"column:points_multiplier"`
	CashbackPercentage float64   `gorm:"column:cashback_percentage"`
	CreatedAt          string    `gorm:"column:created_at"`
	UpdatedAt          string    `gorm:"column:updated_at"`
}

func (c *Campaign) TableName() string {
	return "campaigns"
}

func (c *Campaign) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *Campaign) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
