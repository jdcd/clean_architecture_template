package main

import (
	"fmt"
	"github.com/jdcd9001/clean-architecture-template/internal"
	"github.com/jdcd9001/clean-architecture-template/pkg"
	"os"
)

func main() {
	router := internal.SetupRouter(internal.GetRouterDependencies())
	port := os.Getenv("PORT")

	if err := router.Run(); err != nil {
		errorDetail := fmt.Sprintf("unable to start app on the port: %s , %s", port, err.Error())
		pkg.ErrorLogger().Fatal(errorDetail)
	}
}
