package server_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jdcd9001/clean-architecture-template/internal/application"
	"github.com/jdcd9001/clean-architecture-template/internal/domain"
	"github.com/jdcd9001/clean-architecture-template/internal/domain/ports"
	"github.com/jdcd9001/clean-architecture-template/internal/infrastructure/entity"
	"github.com/jdcd9001/clean-architecture-template/internal/infrastructure/http/server"
	"github.com/jdcd9001/clean-architecture-template/pkg"
	"github.com/jdcd9001/clean-architecture-template/pkg/mock"
	"github.com/stretchr/testify/assert"
)

const (
	peopleURL = "/v1/people"
)

func Test_WhenPostPeopleRequestIsOkThenControllerShouldReturnSuccess201(t *testing.T) {
	repositoryMock := &mock.PeopleRepositoryMock{}
	router := getMockedRouter(repositoryMock)
	idExpected := "1233445323423"
	expectedStruct := entity.Id{Id: idExpected}
	requestObject := domain.People{Name: "cheems", Age: 9}
	body := bytes.NewReader([]byte(`{"name":"cheems", "age": 9}`))
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, peopleURL, body)
	var responseStruck entity.Id
	repositoryMock.On("Create", requestObject).Return(idExpected, nil).Once()

	router.ServeHTTP(recorder, request)
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseStruck)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, expectedStruct, responseStruck)
	repositoryMock.AssertExpectations(t)

}

func Test_WhenAppCannotCreatePersonThenControllerShouldReturn500Error(t *testing.T) {
	repositoryMock := &mock.PeopleRepositoryMock{}
	router := getMockedRouter(repositoryMock)
	errorExpected := pkg.CreateFormatError(pkg.ThirdPart, "some internal Error",
		"database down")
	apiErrorExpected := pkg.MapAPIError(errors.New(errorExpected))
	requestObject := domain.People{Name: "cheems", Age: 9}
	body := bytes.NewReader([]byte(`{"name":"cheems", "age": 9}`))
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, peopleURL, body)
	var responseStruck pkg.APIError
	repositoryMock.On("Create", requestObject).Return("", errors.New(errorExpected)).Once()

	router.ServeHTTP(recorder, request)
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseStruck)

	assert.Equal(t, apiErrorExpected, responseStruck)
	repositoryMock.AssertExpectations(t)

}

func Test_WhenPostPeopleRequestIsWrongThenControllerShouldReturnBadRequest401(t *testing.T) {
	router := getMockedRouter(nil)
	errorExpected := pkg.CreateFormatError(pkg.DataValidation, server.EncodingJsonError, "someError")
	apiErrorExpected := pkg.MapAPIError(errors.New(errorExpected))
	body := bytes.NewReader([]byte(`{"name":"cheems", "age": "fail"}`))
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, peopleURL, body)
	var responseStruck pkg.APIError

	router.ServeHTTP(recorder, request)
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseStruck)

	assert.Equal(t, apiErrorExpected.Code, recorder.Code)
	assert.Equal(t, apiErrorExpected.Message, responseStruck.Message)
}

func Test_WhenPostPeopleRequestHasInvalidDataThenControllerShouldReturnBadRequest401(t *testing.T) {
	router := getMockedRouter(nil)
	errorExpected := pkg.CreateFormatError(pkg.DataValidation, "the field name is invalid, please check it",
		"the field name cannot be empty")
	apiErrorExpected := pkg.MapAPIError(errors.New(errorExpected))
	body := bytes.NewReader([]byte(`{"name":"", "age": 8}`))
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, peopleURL, body)
	var responseStruck pkg.APIError

	router.ServeHTTP(recorder, request)
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseStruck)

	assert.Equal(t, apiErrorExpected.Code, recorder.Code)
	assert.Equal(t, apiErrorExpected.Message, responseStruck.Message)
}

func Test_WhenAppIsFineThenGetAllPeopleControllerShouldReturnDataAnd200(t *testing.T) {
	repositoryMock := &mock.PeopleRepositoryMock{}
	router := getMockedRouter(repositoryMock)
	expectedResult := []domain.People{{
		ID:   "1",
		Name: "cheems",
		Age:  6,
	}}
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, peopleURL, nil)
	var responseStruck []domain.People
	repositoryMock.On("GetAll").Return(expectedResult, nil).Once()

	router.ServeHTTP(recorder, request)
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseStruck)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expectedResult, responseStruck)
	repositoryMock.AssertExpectations(t)
}

func Test_WhenAppCannotReturnPeopleThenGetAllPeopleControllerShouldReturn500(t *testing.T) {
	repositoryMock := &mock.PeopleRepositoryMock{}
	router := getMockedRouter(repositoryMock)
	errorExpected := pkg.CreateFormatError(pkg.ThirdPart, "some internal Error",
		"database down")
	apiErrorExpected := pkg.MapAPIError(errors.New(errorExpected))
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, peopleURL, nil)
	var responseStruck pkg.APIError
	repositoryMock.On("GetAll").Return([]domain.People{}, errors.New(errorExpected)).Once()

	router.ServeHTTP(recorder, request)
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseStruck)

	assert.Equal(t, apiErrorExpected, responseStruck)
	repositoryMock.AssertExpectations(t)
}

func getMockedRouter(repository ports.PeopleRepository) *gin.Engine {
	appMock := &application.PeopleApplication{Repository: repository}
	router := server.SetupRouter(&server.RouterDependencies{
		PeopleController: &server.PeopleController{App: appMock},
		CheckController:  &server.PingController{},
	})
	return router
}
