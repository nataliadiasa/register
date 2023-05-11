package domain

import "github.com/google/uuid"

type BankAccount struct {
	ID           uuid.UUID `json:"id"`
	BranchNumber int       `json:"branch_number"`
	Number       int       `json:"number"`
}
