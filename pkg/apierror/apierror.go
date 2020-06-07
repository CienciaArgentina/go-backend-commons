package apierror

import (
	"github.com/CienciaArgentina/go-backend-commons/pkg/json"
	"net/http"
)

type ErrorList []interface{}

type ApiError interface {
	Status() int
	Message() string
	Error() ErrorList
	WithStatus(status int) *apiError
	WithMessage(message string) *apiError
	AddError(message, code string) *apiError
}

type apiError struct {
	errStatus  int       `json:"status"`
	errMessage string    `json:"message"`
	errError   ErrorList `json:"error"`
}

type ErrorCause struct {
	Detail string `json:"detail"`
	Code   string `json:"code"`
}

func New(status int, message string, error ErrorList) ApiError {
	return &apiError{errStatus: status, errMessage: message, errError: error}
}

func NewWithStatus(status int) ApiError {
	return &apiError{errStatus: status}
}

func NewErrorCause(detail, code string) ErrorList {
	error := ErrorList{}
	error = append(error, &ErrorCause{
		Detail: detail,
		Code:   code,
	})
	return error
}

func (a *apiError) Status() int {
	return a.errStatus
}

func (a *apiError) Message() string {
	return a.errMessage
}

func (a *apiError) Error() ErrorList {
	return a.errError
}

func (e ErrorList) ToString() string {
	str, _ := json.ToJSONString(e)
	return str
}

func (a *apiError) AddError(message, code string) *apiError {
	a.errError = append(a.errError, ErrorCause{message, code})
	return a
}

func (a *apiError) WithStatus(status int) *apiError {
	a.errStatus = status
	return a
}

func (a *apiError) WithMessage(message string) *apiError {
	a.errMessage = message
	return a
}

func NewNotFoundApiError(message string) ApiError {
	return &apiError{http.StatusNotFound, message, NewErrorCause(message, "not_found")}
}

func NewBadRequestApiError(message string) ApiError {
	return &apiError{http.StatusBadRequest, message, NewErrorCause(message, "bad_request")}
}

func NewMethodNotAllowedApiError() ApiError {
	return &apiError{http.StatusMethodNotAllowed, "Method not allowed", NewErrorCause("Method not allowed", "method_not_allowed")}
}

func NewInternalServerApiError(message string, err error) ApiError {
	error := ErrorList{}
	if err != nil {
		error = append(error, err.Error())
	}
	return &apiError{http.StatusInternalServerError, message, NewErrorCause(message, "internal_server_error")}
}

func NewForbiddenApiError(message string) ApiError {
	return &apiError{http.StatusForbidden, message, NewErrorCause(message, "forbidden")}
}

func NewUnauthorizedApiError(message string) ApiError {
	return &apiError{http.StatusUnauthorized, message, NewErrorCause(message, "unauthorized_scopes")}
}
