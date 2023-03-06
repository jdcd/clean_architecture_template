package config

import (
	"github.com/jdcd9001/clean-architecture-template/internal/application"
	"github.com/jdcd9001/clean-architecture-template/internal/domain"
	"github.com/jdcd9001/clean-architecture-template/internal/domain/ports"
	"github.com/jdcd9001/clean-architecture-template/internal/infrastructure/adapters"
	"github.com/jdcd9001/clean-architecture-template/internal/infrastructure/http/server"
)

func GetRouterDependencies() *server.RouterDependencies {
	return &server.RouterDependencies{
		CheckController: &server.PingController{},
		PeopleController: &server.PeopleController{
			App: &application.PeopleApplication{
				Repository: getPeopleRepository(),
			},
		},
	}
}

func getPeopleRepository() ports.PeopleRepository {
	var storageRepo []domain.People
	return &adapters.LocalPeopleRepository{Storage: storageRepo}
}
