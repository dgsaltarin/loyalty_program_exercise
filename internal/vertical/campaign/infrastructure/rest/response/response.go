package response

import "time"

type CampaignResponse struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
	PointsMultiplier   float64   `json:"points_multiplier"`
	CashbackPercentage float64   `json:"cashback_percentage"`
	CommerceID         string    `json:"commerce_id"`
	BranchID           string    `json:"branch_id"`
}
