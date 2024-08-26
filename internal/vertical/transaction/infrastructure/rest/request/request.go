package request

type CreateTransactionRequest struct {
	Amount     float64 `json:"amount"`
	CommerceID string  `json:"commerce_id"`
	BranchID   string  `json:"branch_id"`
	UserID     string  `json:"user_id"`
}
