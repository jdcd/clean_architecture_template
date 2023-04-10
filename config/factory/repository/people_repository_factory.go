package repository

import (
	"github.com/jdcd9001/clean-architecture-template/internal/domain"
	"github.com/jdcd9001/clean-architecture-template/internal/domain/ports"
	"github.com/jdcd9001/clean-architecture-template/internal/infrastructure/adapters"
	"github.com/jdcd9001/clean-architecture-template/pkg"
)

type AvailableRepo string

const (
	Local AvailableRepo = "LOCAL"
	MySQL AvailableRepo = "MYSQL"
	Mongo AvailableRepo = "MONGO"
)

func PeopleRepositoryFactory(opt AvailableRepo) ports.PeopleRepository {
	switch opt {
	case Local:
		return getLocalRepository()
	default:
		pkg.WarningLogger().Printf("{%s} is not a valid repository, switch to local...\n", opt)
		return getLocalRepository()
	}
}

func getLocalRepository() ports.PeopleRepository {
	var storageRepo []domain.People
	return &adapters.LocalPeopleRepository{Storage: storageRepo}
}
