package response

import "time"

type CreateTransactionResponse struct {
	ID              string    `json:"id"`
	Amount          float64   `json:"amount"`
	CommerceID      string    `json:"commerce_id"`
	BranchID        string    `json:"branch_id"`
	UserID          string    `json:"user_id"`
	PointsEarned    float64   `json:"points_earned"`
	CashbackEarned  float64   `json:"cashback_earned"`
	TransactionDate time.Time `json:"transaction_date"`
}
