package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nataliadiasa/register/domain"
)

var ErrNotFound = errors.New("record not found")

type MemoryRepository struct {
	persons map[uuid.UUID]domain.Person
}

func New() *MemoryRepository {
	return &MemoryRepository{persons: make(map[uuid.UUID]domain.Person)}
}

func (mr *MemoryRepository) Create(ctx context.Context, person domain.Person) domain.Person {
	person.ID = uuid.New()
	mr.persons[person.ID] = person
	return person
}

func (mr *MemoryRepository) Update(ctx context.Context, person domain.Person, id uuid.UUID) error {
	if _, ok := mr.persons[id]; !ok {
		return ErrNotFound
	}

	person.ID = id
	mr.persons[id] = person
	return nil
}

func (mr *MemoryRepository) GetAll(ctx context.Context) []domain.Person {
	list := []domain.Person{}
	for _, v := range mr.persons {
		list = append(list, v)
	}

	return list
}

func (mr *MemoryRepository) Get(ctx context.Context, id uuid.UUID) (domain.Person, error) {
	if _, ok := mr.persons[id]; !ok {
		return domain.Person{}, ErrNotFound
	}

	return mr.persons[id], nil
}
