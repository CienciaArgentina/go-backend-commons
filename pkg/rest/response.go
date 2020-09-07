package rest

import (
	"net/http"
)

type Response interface{}

type response struct {
	response *http.Response
	Err      error
}

type CienciaArgentinaResponse interface {
}

type cienciaArgentinaResponse struct {
	*response
	executionTime int64
}
