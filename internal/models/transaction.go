package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID                int       `json:"id"`
	Amount            float64   `json:"amount"`
	Currency          string    `json:"currency"`
	Date              time.Time `json:"date"`
	Description       string    `json:"description"`
	AccountID         int       `json:"account_id"`
	Account           Account   `gorm:"foreignKey:AccountID"`
	CategoryID        int       `json:"category_id"`
	Category          Category  `gorm:"foreignKey:CategoryID"`
	TransactionTypeID int       `json:"transaction_type_id"`
	TransactionType   Category  `gorm:"foreignKey:TransactionTypeID"`
}

func TransactionTypeColors() map[string]string {
	return map[string]string{
		"debit":    "#FDA403",
		"credit":   "#898121",
		"deposit":  "#E5C287",
		"withdraw": "#E8751A",
	}
}
