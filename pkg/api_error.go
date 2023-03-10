package pkg

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type ErrorType string

const (
	DataNotFound    ErrorType = "dataNotFound"
	ConnectionError ErrorType = "connectionError"
	DataValidation  ErrorType = "dataValidation"
	ThirdPart       ErrorType = "thirdPart"
	BusinessRule    ErrorType = "businessRule"
)

const (
	UnknownErrorDetail               = "unknown error"
	GenericInternalServerErrorDetail = "internal server error happens when try to process your request"
)

type ApiError struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"error_detail"`
	Message      string `json:"message"`
}

func CreateFormatError(errorType ErrorType, message string, errDetail string) string {
	return fmt.Sprintf("%s|%s|%s", errorType, message, errDetail)
}

func MapApiError(err error) ApiError {
	if err == nil {
		return getStandardApiError(errors.New(UnknownErrorDetail))
	}

	parts := strings.Split(err.Error(), "|")
	if len(parts) != 3 {
		return getStandardApiError(err)
	}

	errorType := ErrorType(parts[0])
	errorMessage := parts[1]
	errorDetail := parts[2]

	switch errorType {
	case ThirdPart, ConnectionError, DataValidation, DataNotFound, BusinessRule:
		return ApiError{
			Code:         selectStatusCode(errorType),
			Message:      errorMessage,
			ErrorMessage: errorDetail,
		}
	default:
		return getStandardApiError(errors.New(errorDetail))
	}
}

func selectStatusCode(errorType ErrorType) int {
	switch errorType {
	case ThirdPart, ConnectionError:
		return http.StatusInternalServerError
	case DataValidation:
		return http.StatusBadRequest
	case DataNotFound, BusinessRule:
		return http.StatusOK
	default:
		return http.StatusInternalServerError
	}
}

func getStandardApiError(err error) ApiError {
	return ApiError{
		Code:         http.StatusInternalServerError,
		Message:      GenericInternalServerErrorDetail,
		ErrorMessage: err.Error(),
	}
}
