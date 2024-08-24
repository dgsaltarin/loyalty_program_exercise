package entity

import "time"

type Campaign struct {
	ID                 string
	Name               string
	StartDate          time.Time
	EndDate            time.Time
	PointsMultiplier   float64
	CashbackPercentage float64
}
