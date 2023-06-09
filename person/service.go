package person

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/nataliadiasa/register/domain"
)

type Service struct {
	repository *MemoryRepository
}

func NewService(r *MemoryRepository) *Service {
	return &Service{repository: r}
}

var ErrInvalidField = errors.New("invalid field")

// func validate(...) error -> ErrInvalidField + Campo que está invalido.
// func validate(...) string -> Campo que está invalido.+

func validate(person domain.Person) error {
	if len(person.Phone) != 11 {
		return fmt.Errorf("telefone tem que ter 11 digitos. %w", ErrInvalidField)
	}

	if person.Age < 0 {
		return fmt.Errorf("idade precisa ser maior que 0. %w", ErrInvalidField)
	}

	if person.Name == "" {
		return fmt.Errorf("nome nao foi inserido. %w", ErrInvalidField)
	}

	return nil
}

func (s Service) Create(person domain.Person) (domain.Person, error) {
	if err := validate(person); err != nil {
		return domain.Person{}, err
	}

	res := s.repository.Create(context.Background(), person)
	return res, nil
}

func (s Service) GetAll() []domain.Person {
	return s.repository.GetAll(context.Background())
}

func (s Service) Get(id uuid.UUID) (domain.Person, error) {
	return s.repository.Get(context.Background(), id)
}

func (s Service) Update(person domain.Person, id uuid.UUID) error {
	if err := validate(person); err != nil {
		return err
	}

	return s.repository.Update(context.Background(), person, id)
}
