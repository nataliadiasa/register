package bankaccount

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nataliadiasa/register/domain"
	"github.com/nataliadiasa/register/person"
)

var ErrForeignKeyViolation = errors.New("foreing key violation")

type MemoryRepository struct {
	bankAccounts     map[uuid.UUID]domain.BankAccount
	personRepository *person.MemoryRepository
}

func NewRepository(personRepository *person.MemoryRepository) *MemoryRepository {
	return &MemoryRepository{
		bankAccounts:     make(map[uuid.UUID]domain.BankAccount),
		personRepository: personRepository,
	}
}

func (mr *MemoryRepository) Create(ctx context.Context, bankAccount domain.BankAccount) (domain.BankAccount, error) {
	_, err := mr.personRepository.Get(ctx, bankAccount.PersonID)
	if err != nil {
		if errors.Is(err, person.ErrNotFound) {
			return domain.BankAccount{}, ErrForeignKeyViolation
		}
		return domain.BankAccount{}, err
	}

	bankAccount.ID = uuid.New()
	mr.bankAccounts[bankAccount.ID] = bankAccount
	return bankAccount, nil
}

func (mr *MemoryRepository) Get(ctx context.Context, id uuid.UUID) (domain.BankAccount, error) {
	if _, ok := mr.bankAccounts[id]; !ok {
		return domain.BankAccount{}, ErrForeignKeyViolation
	}

	return mr.bankAccounts[id], nil
}
