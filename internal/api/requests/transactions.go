package requests

import (
	"time"

	"github.com/christo-andrew/maybe-go/internal/models"
)

type CreateTransactionRequest struct {
	AccountID         int       `json:"account_id"`
	Amount            float64   `json:"amount"`
	Currency          string    `json:"currency"`
	Date              time.Time `json:"date"`
	Description       string    `json:"description"`
	CategoryID        int       `json:"category_id"`
	TransactionTypeID int       `json:"transaction_type_id"`
}

func (c *CreateTransactionRequest) Transaction() *models.Transaction {
	return &models.Transaction{
		AccountID:         c.AccountID,
		Amount:            c.Amount,
		Currency:          c.Currency,
		Date:              c.Date,
		Description:       c.Description,
		CategoryID:        c.CategoryID,
		TransactionTypeID: c.TransactionTypeID,
	}
}
