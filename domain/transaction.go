package domain

import (
	"time"

	"github.com/google/uuid"
)

type Category string

type Transaction struct {
	ID            uuid.UUID `json:"id"`
	Value         int       `json:"value"`
	Date          time.Time `json:"date"`
	Category      Category  `json:"category"`
	BankAccountID uuid.UUID `json:"bank_account_id"`
}
