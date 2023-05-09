package service

import (
	"context"

	"github.com/nataliadiasa/register/domain"
	"github.com/nataliadiasa/register/repository"
)

type Service struct {
	repository *repository.MemoryRepository
}

func New(r *repository.MemoryRepository) *Service {
	return &Service{repository: r}
}

func (s Service) Create(person domain.Person) {
	s.repository.Create(context.Background(), person)
}

func (s Service) GetAll() []domain.Person {
	return s.repository.GetAll(context.Background())
}

func (s Service) Get(id int) (domain.Person, error) {
	return s.repository.Get(context.Background(), id)
}

func (s Service) Update(person domain.Person, id int) error {
	return s.repository.Update(context.Background(), person, id)
}
