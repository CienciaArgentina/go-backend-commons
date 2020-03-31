package apierror

import (
	"github.com/CienciaArgentina/go-backend-commons/pkg/json"
	"net/http"
)

//type ErrorList []interface{}
//
//type ApiError interface {
//	Message() string
//	Code() string      // ex. invalid_user
//	Status() int       // http status code
//	Errors() ErrorList // detailed error list
//}
//
//type apiError struct {
//	ErrorMessage string    `json:"message"`
//	ErrorCode    string    `json:"code"`
//	ErrorStatus  int       `json:"status"`
//	ErrorList    ErrorList `json:"errors"`
//}
//
//func New(message, code string, status int, errors ErrorList) ApiError {
//	return apiError{message, code, status, errors}
//}
//
//func (e ErrorList) ToString() string {
//	return fmt.Sprint(e)
//}
//
//func (e apiError) Code() string {
//	return e.ErrorCode
//}
//
//func (e apiError) Status() int {
//	return e.ErrorStatus
//}
//
//func (e apiError) Message() string {
//	return e.ErrorMessage
//}
//
//func (e apiError) Errors() ErrorList {
//	return e.ErrorList
//}
//
//func NewNotFoundApiError(message string) ApiError {
//	return apiError{message, "not_found", http.StatusNotFound, ErrorList{}}
//}
//
//func NewBadRequestApiError(message string) ApiError {
//	return apiError{message, "bad_request", http.StatusBadRequest, ErrorList{}}
//}
//
//func NewMethodNotAllowedApiError() ApiError {
//	return apiError{"Method not allowed", "method_not_allowed", http.StatusMethodNotAllowed, ErrorList{}}
//}
//
//func NewInternalServerApiError(message string, err error) ApiError {
//	error := ErrorList{}
//	if err != nil {
//		error = append(error, err.Error())
//	}
//	return apiError{message, "internal_server_error", http.StatusInternalServerError, error}
//}
//
//func NewForbiddenApiError(message string) ApiError {
//	return apiError{message, "forbidden", http.StatusForbidden, ErrorList{}}
//}
//
//func NewUnauthorizedApiError(message string) ApiError {
//	return apiError{message, "unauthorized_scopes", http.StatusUnauthorized, ErrorList{}}
//}

type ErrorList []interface{}

type ApiError interface {
	Status() int
	Message() string
	Error() ErrorList
	AddError(message, code string)
}

type apiError struct {
	ErrStatus int `json:"status"`
	ErrMessage string `json:"message"`
	ErrError ErrorList `json:"error"`
}

type ErrorCause struct {
	Detail string `json:"detail"`
	Code string `json:"code"`
}

func New(status int, message string, error ErrorList) ApiError {
	return &apiError{ErrStatus:status, ErrMessage:message, ErrError:error}
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

func (a *apiError) Error() ErrorList {
	return a.ErrError
}

func (e ErrorList) ToString() string {
	str, _ := json.ToJSONString(e)
	return str
}

func (a *apiError) AddError(message, code string)  {
	a.ErrError = append(a.ErrError, ErrorCause{message, code})
}

func NewNotFoundApiError(message string) ApiError {
	return &apiError{http.StatusNotFound, message, NewErrorCause(message, "not_found")}
}

func NewBadRequestApiError(message string) ApiError {
	return &apiError{http.StatusBadRequest, message, NewErrorCause(message, "bad_request")}
}

func NewMethodNotAllowedApiError() ApiError {
	return &apiError{http.StatusMethodNotAllowed, "Method not allowed", NewErrorCause("Method not allowed","method_not_allowed")}
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