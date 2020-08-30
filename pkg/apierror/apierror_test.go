package apierror

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewShouldReturnNewApiError(t *testing.T) {
	code := http.StatusBadRequest
	msg := "Bad request"
	err := New(code, msg, nil)
	require.NotNil(t, err)
	require.Equal(t, code, err.Status())
	require.Equal(t, msg, err.Message())
	require.IsType(t, &apiError{}, err)
}

func TestNewWithStatusShouldReturnNewApiErrorWithStatus(t *testing.T) {
	code := http.StatusBadRequest
	err := NewWithStatus(code)
	require.NotNil(t, err)
	require.Equal(t, code, err.Status())
	require.IsType(t, &apiError{}, err)
}

func TestNewErrorCauseShouldCreateNewErrorList(t *testing.T) {
	detail := "Testing"
	code := "testing_code"

	err := NewErrorCause(detail, code)

	require.NotNil(t, err)
	require.Equal(t, "[{\"detail\":\"Testing\",\"code\":\"testing_code\"}]", err.String())
}

func TestAddErrorShouldAddError(t *testing.T) {
	code := http.StatusBadRequest
	msg := "Bad request"
	err := New(code, msg, nil)
	errmsg := "AddError err"
	errcode := "adding_error_code"
	err.AddError(errmsg, errcode)
	require.NotNil(t, err)
	require.Equal(t, code, err.Status())
	require.Equal(t, msg, err.Message())
	require.Equal(t, "[{\"detail\":\"AddError err\",\"code\":\"adding_error_code\"}]", err.Error())
	require.IsType(t, &apiError{}, err)
}

func TestWithStatusShouldSetStatus(t *testing.T) {
	expected := &apiError{ErrStatus: http.StatusBadRequest}
	err := NewWithStatus(http.StatusInternalServerError)
	err.WithStatus(http.StatusBadRequest)
	require.Equal(t, expected, err)
}

func TestWithMessageShouldSetErrMsg(t *testing.T) {
	expected := &apiError{ErrMessage: "test", ErrStatus: http.StatusBadRequest}
	err := NewWithStatus(http.StatusBadRequest)
	err.WithMessage("test")
	require.Equal(t, expected, err)
}

func TestCommonApiErrors(t *testing.T) {
	msg := "Test msg"

	notFound := NewNotFoundApiError(msg)
	require.Equal(t, msg, notFound.Message())
	require.Equal(t, http.StatusNotFound, notFound.Status())

	badReq := NewBadRequestApiError(msg)
	require.Equal(t, msg, badReq.Message())
	require.Equal(t, http.StatusBadRequest, badReq.Status())
}
