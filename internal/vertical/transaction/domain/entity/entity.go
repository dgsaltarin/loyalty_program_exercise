package entity

import (
	"time"

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
	TransactionDate time.Time
}

func NewTransaction(transaction Transaction) *Transaction {
	return &Transaction{
		ID:              uuid.New().String(),
		UserID:          transaction.UserID,
		CommerceID:      transaction.CommerceID,
		BranchID:        transaction.BranchID,
		CampaignID:      transaction.CampaignID,
		Amount:          transaction.Amount,
		PointsEarned:    transaction.PointsEarned,
		CashbackEarned:  transaction.CashbackEarned,
		TransactionDate: time.Now(),
	}
}
