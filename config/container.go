package config

import (
	"github.com/jdcd9001/clean-architecture-template/config/factory/repository"
	"github.com/jdcd9001/clean-architecture-template/internal/application"
	"github.com/jdcd9001/clean-architecture-template/internal/infrastructure/http/server"
)

func GetRouterDependencies(config *AppConfiguration) *server.RouterDependencies {
	return &server.RouterDependencies{
		CheckController: &server.PingController{},
		PeopleController: &server.PeopleController{
			App: &application.PeopleApplication{
				Repository: repository.PeopleRepositoryFactory(config.Repository()),
			},
		},
	}
}
