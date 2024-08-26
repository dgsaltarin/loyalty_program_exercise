package entity

type Transaction struct {
	ID              string
	UserID          string
	CommerceID      string
	BranchID        string
	CampaignID      string
	Amount          float64
	PointsEarned    float64
	CashbackEarned  float64
	TransactionDate string
}
