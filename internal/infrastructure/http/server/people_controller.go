package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jdcd9001/clean-architecture-template/internal/application"
	"github.com/jdcd9001/clean-architecture-template/internal/domain"
	"github.com/jdcd9001/clean-architecture-template/internal/infrastructure/entity"
	"github.com/jdcd9001/clean-architecture-template/pkg"
)

const (
	EncodingJsonError        = "error reading your request"
	successfulCreationLogger = "person: with id %s creates successfully \n"
	successfulGetAllLogger   = "get all people query response successfully"
)

type PeopleController struct {
	App application.IPeopleApp
}

func (r *PeopleController) GetAllPeople(c *gin.Context) {
	result, err := r.App.GetAllPersons()
	if err != nil {
		apiError := pkg.MapApiError(err)
		c.IndentedJSON(apiError.Code, apiError)
		return
	}

	pkg.InfoLogger().Println(successfulGetAllLogger)
	c.IndentedJSON(http.StatusOK, result)
}

func (r *PeopleController) CreatePerson(c *gin.Context) {
	var person domain.People
	if err := c.BindJSON(&person); err != nil {
		formattedError := pkg.CreateFormatError(pkg.DataValidation, EncodingJsonError, err.Error())
		apiError := pkg.MapApiError(errors.New(formattedError))
		c.IndentedJSON(apiError.Code, apiError)
		pkg.ErrorLogger().Printf("%s, %s \n", EncodingJsonError, err.Error())
		return
	}

	if err := person.Validate(); err != nil {
		apiError := pkg.MapApiError(err)
		c.IndentedJSON(apiError.Code, apiError)
		return
	}

	id, err := r.App.CreatePerson(person)
	if err != nil {
		apiError := pkg.MapApiError(err)
		c.IndentedJSON(apiError.Code, apiError)
		return
	}

	responseStruct := entity.Id{Id: id}
	pkg.InfoLogger().Printf(successfulCreationLogger, id)
	c.IndentedJSON(http.StatusCreated, responseStruct)

}
