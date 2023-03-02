package mock

import (
	"github.com/jdcd9001/clean-architecture-template/internal/domain"
	"github.com/stretchr/testify/mock"
)

type PeopleRepositoryMock struct {
	mock.Mock
}

func (m *PeopleRepositoryMock) Create(person domain.People) (string, error) {
	args := m.Called(person)
	return args.String(0), args.Error(1)
}

func (m *PeopleRepositoryMock) GetAll() ([]domain.People, error) {
	args := m.Called()
	return args.Get(0).([]domain.People), args.Error(1)
}
