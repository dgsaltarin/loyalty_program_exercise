package models

type Transaction struct {
	ID              string `gorm:"primaryKey"`
	Amount          int    `gorm:"not null"`
	CommerceID      string `gorm:"not null"`
	BranchID        string `gorm:"not null"`
	UserID          string `gorm:"not null"`
	PointsErned     int    `gorm:"not null"`
	CashbackEarned  int    `gorm:""`
	TransactionDate string `gorm:"not null"`
	CreatedAt       string `gorm:"autoCreateTime"`
}
