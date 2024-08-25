package request

type CreateCampaignRequest struct {
	Name               string  `json:"name"`
	StartDate          string  `json:"start_date"`
	EndDate            string  `json:"end_date"`
	PointsMultiplier   float64 `json:"points_multiplier"`
	CashbackPercentage float64 `json:"cashback_percentage"`
	CommerceID         string  `json:"commerce_id"`
	BranchID           string  `json:"branch_id"`
}
