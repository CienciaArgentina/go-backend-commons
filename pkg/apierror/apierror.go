package apierror

import (
	"encoding/json"
	"net/http"

	j "github.com/CienciaArgentina/go-backend-commons/pkg/json"
)

type ErrorList []interface{}

type ApiError interface {
	Status() int
	Message() string
	Error() string
	Errors() ErrorList
	WithStatus(status int) *apiError
	WithMessage(message string) *apiError
	AddError(message, code string) *apiError
}

type apiError struct {
	ErrStatus  int       `json:"status"`
	ErrMessage string    `json:"message"`
	ErrError   ErrorList `json:"error"`
}

type ErrorCause struct {
	Detail string `json:"detail"`
	Code   string `json:"code"`
}

func New(status int, message string, error ErrorList) ApiError {
	return &apiError{ErrStatus: status, ErrMessage: message, ErrError: error}
}

func NewWithStatus(status int) ApiError {
	return &apiError{ErrStatus: status}
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
	return a.ErrStatus
}

func (a *apiError) Message() string {
	return a.ErrMessage
}

func (a *apiError) Error() string {
	return a.ErrError.String()
}

func (a *apiError) Errors() ErrorList {
	return a.ErrError
}

func (e ErrorList) String() string {
	str, _ := j.ToJSONString(e)
	return str
}

func (a *apiError) AddError(message, code string) *apiError {
	a.ErrError = append(a.ErrError, ErrorCause{message, code})
	return a
}

func (a *apiError) WithStatus(status int) *apiError {
	a.ErrStatus = status
	return a
}

func (a *apiError) WithMessage(message string) *apiError {
	a.ErrMessage = message
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
func NewInternalServerApiError(message string, err error, code string) ApiError {
	errL := ErrorList{}
	if err != nil {
		errL = NewErrorCause(err.Error(), code)
	}
	return &apiError{http.StatusInternalServerError, message, errL}
}

func NewForbiddenApiError(message string) ApiError {
	return &apiError{http.StatusForbidden, message, NewErrorCause(message, "forbidden")}
}

func NewUnauthorizedApiError(message string) ApiError {
	return &apiError{http.StatusUnauthorized, message, NewErrorCause(message, "unauthorized_scopes")}
}

func NewErrorFromBytesWithStatus(data []byte, status int) (ApiError, error) {
	var apierror apiError
	err := json.Unmarshal(data, &status)
	if apierror.ErrStatus == 0 {
		apierror.WithStatus(status)
	}
	return &apierror, err
}
