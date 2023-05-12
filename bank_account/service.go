package bankaccount

import (
	"context"
	"errors"
	"fmt"

	"github.com/nataliadiasa/register/domain"
	"github.com/nataliadiasa/register/person"
)

var ErrPersonNeeded = errors.New("person is needed")

type Service struct {
	repository    *MemoryRepository
	personService *person.Service
}

func NewService(r *MemoryRepository, personService *person.Service) *Service {
	return &Service{
		repository:    r,
		personService: personService,
	}
}

func (s Service) Create(bankAccount domain.BankAccount) (domain.BankAccount, error) {
	_, err := s.personService.Get(bankAccount.PersonID)
	if err != nil {
		if errors.Is(err, person.ErrNotFound) {
			return domain.BankAccount{}, fmt.Errorf("failed to create bank account. %w", ErrPersonNeeded)
		}
		return domain.BankAccount{}, err
	}

	return s.repository.Create(context.Background(), bankAccount)
}
