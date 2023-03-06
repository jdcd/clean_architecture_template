package server

import (
	"github.com/gin-gonic/gin"
)

type RouterDependencies struct {
	CheckController  *PingController
	PeopleController *PeopleController
}

func SetupRouter(d *RouterDependencies) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")

	v1Check := v1.Group("/ping")
	v1Check.GET("", d.CheckController.GetPing)

	v1People := v1.Group("/people")
	v1People.POST("", d.PeopleController.CreatePerson)
	v1People.GET("", d.PeopleController.GetAllPeople)

	return router
}
