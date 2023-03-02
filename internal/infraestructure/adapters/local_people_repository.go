package adapters

import (
	"github.com/google/uuid"
	"github.com/jdcd9001/clean-architecture-template/internal/domain"
)

type LocalPeopleRepository struct {
	Storage []domain.People
}

func (r *LocalPeopleRepository) Create(people domain.People) (id string, err error) {
	people.Id = uuid.New().String()
	r.Storage = append(r.Storage, people)
	return people.Id, nil
}

func (r *LocalPeopleRepository) GetAll() (results []domain.People, err error) {
	return r.Storage, nil
}
