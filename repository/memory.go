package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nataliadiasa/register/domain"
)

var ErrNotFound = errors.New("record not found")

type MemoryRepository struct {
	persons []domain.Person
}

func New() *MemoryRepository {
	return &MemoryRepository{persons: []domain.Person{}}
}

func (mr *MemoryRepository) Create(ctx context.Context, person domain.Person) domain.Person {
	person.ID = uuid.New()
	mr.persons = append(mr.persons, person)
	return person
}

func (mr *MemoryRepository) Update(ctx context.Context, person domain.Person, id int) error {
	if id-1 < 0 || id-1 >= len(mr.persons) {
		return ErrNotFound
	}
	mr.persons[id-1] = person
	return nil
}

func (mr *MemoryRepository) GetAll(ctx context.Context) []domain.Person {
	return mr.persons
}

func (mr *MemoryRepository) Get(ctx context.Context, id int) (domain.Person, error) {
	if id-1 < 0 || id-1 >= len(mr.persons) {
		return domain.Person{}, ErrNotFound
	}
	return mr.persons[id-1], nil
}
