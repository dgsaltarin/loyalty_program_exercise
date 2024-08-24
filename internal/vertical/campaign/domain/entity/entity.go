package entity

import (
	"time"

	"github.com/google/uuid"
)

type Campaign struct {
	ID                 string
	Name               string
	CommerceID         string
	BranchID           string
	StartDate          time.Time
	EndDate            time.Time
	PointsMultiplier   float64
	CashbackPercentage float64
}

func NewCampaign(name string, CommerceID string, BranchID string, startDate time.Time, endDate time.Time, pointsMultiplier float64, cashbackPercentage float64) *Campaign {
	return &Campaign{
		ID:                 uuid.New().String(),
		Name:               name,
		CommerceID:         CommerceID,
		BranchID:           BranchID,
		StartDate:          startDate,
		EndDate:            endDate,
		PointsMultiplier:   pointsMultiplier,
		CashbackPercentage: cashbackPercentage,
	}
}
