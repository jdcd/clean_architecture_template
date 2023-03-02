package application

import (
	"github.com/jdcd9001/clean-architecture-template/internal/domain"
	"github.com/jdcd9001/clean-architecture-template/internal/domain/ports"
)

type IPeopleApp interface {
	CreatePerson(person domain.People) (id string, err error)
	GetAllPersons() (results []domain.People, err error)
}

type PeopleApplication struct {
	Repository ports.PeopleRepository
}

func (r *PeopleApplication) CreatePerson(person domain.People) (string, error) {
	return r.Repository.Create(person)
}

func (r *PeopleApplication) GetAllPersons() ([]domain.People, error) {
	return r.Repository.GetAll()
}
