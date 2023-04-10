package main

import (
	"fmt"
	"github.com/jdcd9001/clean-architecture-template/config"
	"github.com/jdcd9001/clean-architecture-template/internal/infrastructure/http/server"
	"os"

	"github.com/jdcd9001/clean-architecture-template/pkg"
)

func main() {
	appConfiguration := config.GetConfigurations()
	router := server.SetupRouter(config.GetRouterDependencies(appConfiguration))
	port := os.Getenv("PORT")

	if err := router.Run(); err != nil {
		errorDetail := fmt.Sprintf("unable to start app on the port: %s , %s", port, err.Error())
		pkg.ErrorLogger().Fatal(errorDetail)
	}
}
