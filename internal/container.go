package internal

import "github.com/jdcd9001/clean-architecture-template/internal/infraestructure/http/server"

func GetRouterDependencies() *RouterDependencies {
	return &RouterDependencies{
		CheckController: &server.PingController{},
	}
}
