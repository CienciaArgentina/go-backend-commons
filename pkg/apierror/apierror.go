package apierror

import (
	"fmt"
	"net/http"
)

type ErrorList []interface{}

type ApiError interface {
	Message() string
	Code() string // ex. invalid_user
	Status() int // http status code
	Errors() ErrorList // detailed error list
}

type apiError struct {
	ErrorMessage string    `json:"message"`
	ErrorCode    string    `json:"code"`
	ErrorStatus  int       `json:"status"`
	ErrorList       ErrorList `json:"errors"`
}

func New(message, code string,status int, errors ErrorList) ApiError {
	return apiError{message, code, status, errors}
}

func (e ErrorList) ToString() string {
	return fmt.Sprint(e)
}

func (e apiError) Code() string {
	return e.ErrorCode
}

func (e apiError) Status() int {
	return e.ErrorStatus
}

func (e apiError) Message() string {
	return e.ErrorMessage
}

func (e apiError) Errors() ErrorList {
	return e.ErrorList
}

func NewNotFoundApiError(message string) ApiError {
	return apiError{message, "not_found", http.StatusNotFound, ErrorList{}}
}

func NewBadRequestApiError(message string) ApiError {
	return apiError{message, "bad_request", http.StatusBadRequest, ErrorList{}}
}

func NewMethodNotAllowedApiError() ApiError {
	return apiError{"Method not allowed", "method_not_allowed", http.StatusMethodNotAllowed, ErrorList{}}
}

func NewInternalServerApiError(message string, err error) ApiError {
	error := ErrorList{}
	if err != nil {
		error = append(error, err.Error())
	}
	return apiError{message, "internal_server_error", http.StatusInternalServerError, error}
}

func NewForbiddenApiError(message string) ApiError {
	return apiError{message, "forbidden", http.StatusForbidden, ErrorList{}}
}

func NewUnauthorizedApiError(message string) ApiError {
	return apiError{message, "unauthorized_scopes", http.StatusUnauthorized, ErrorList{}}
}

