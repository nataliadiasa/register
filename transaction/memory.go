package transaction

import (
	"context"
	"errors"

	"github.com/google/uuid"
	bankaccount "github.com/nataliadiasa/register/bank_account"
	"github.com/nataliadiasa/register/domain"
)

var ErrNotFound = errors.New("record not found")

type MemoryRepository struct {
	transactions   map[uuid.UUID]domain.Transaction
	bankRepository *bankaccount.MemoryRepository
}

func NewRepository(bankRepository *bankaccount.MemoryRepository) *MemoryRepository {
	return &MemoryRepository{
		transactions:   make(map[uuid.UUID]domain.Transaction),
		bankRepository: bankRepository,
	}
}

func (mr *MemoryRepository) Create(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	_, err := mr.bankRepository.Get(ctx, transaction.BankAccountID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return domain.Transaction{}, ErrNotFound
		}
	}
	transaction.ID = uuid.New()
	mr.transactions[transaction.ID] = transaction
	return transaction, nil
}

func (mr *MemoryRepository) Get(ctx context.Context, id uuid.UUID) (domain.Transaction, error) {
	if _, ok := mr.transactions[id]; !ok {
		return domain.Transaction{}, ErrNotFound
	}

	return mr.transactions[id], nil
}
