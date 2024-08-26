package models

type Transaction struct {
	ID              string  `gorm:"primaryKey"`
	Amount          float64 `gorm:"not null"`
	CommerceID      string  `gorm:"not null"`
	BranchID        string  `gorm:"not null"`
	UserID          string  `gorm:"not null"`
	PointsErned     float64 `gorm:"not null"`
	CashbackEarned  float64 `gorm:""`
	TransactionDate string  `gorm:"not null"`
	CreatedAt       string  `gorm:"autoCreateTime"`
}
