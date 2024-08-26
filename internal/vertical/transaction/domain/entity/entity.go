package entity

import (
	"github.com/google/uuid"
)

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

func NewTransaction(userID, commerceID, branchID, campaignID string, amount, pointsEarned, cashbackEarned float64, transactionDate string) *Transaction {
	return &Transaction{
		ID:              uuid.New().String(),
		UserID:          userID,
		CommerceID:      commerceID,
		BranchID:        branchID,
		CampaignID:      campaignID,
		Amount:          amount,
		PointsEarned:    pointsEarned,
		CashbackEarned:  cashbackEarned,
		TransactionDate: transactionDate,
	}
}
