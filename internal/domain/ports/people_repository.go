package ports

import "github.com/jdcd9001/clean-architecture-template/internal/domain"

// PeopleRepository allow to manage a CRUD with people structs
type PeopleRepository interface {
	Create(people domain.People) (id string, err error)
	GetAll() (results []domain.People, err error)
}
