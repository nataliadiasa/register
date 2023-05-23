package transaction

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	bankaccount "github.com/nataliadiasa/register/bank_account"
	"github.com/nataliadiasa/register/domain"
)

type Service struct {
	repository  *MemoryRepository
	bankService *bankaccount.Service
}

func NewService(r *MemoryRepository, bankService *bankaccount.Service) *Service {
	return &Service{
		repository:  r,
		bankService: bankService,
	}
}

func (s Service) Create(transaction domain.Transaction) (domain.Transaction, error) {
	_, err := s.bankService.Get(transaction.BankAccountID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return domain.Transaction{}, fmt.Errorf("failed to create transaction. %w", bankaccount.ErrPersonNeeded)
		}
		return domain.Transaction{}, err
	}

	return s.repository.Create(context.Background(), transaction)
}

func (s Service) Get(id uuid.UUID) (domain.Transaction, error) {
	return s.repository.Get(context.Background(), id)
}
