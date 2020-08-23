package rest

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/CienciaArgentina/go-backend-commons/pkg/clog"

	"github.com/CienciaArgentina/go-backend-commons/pkg/scope"
	"github.com/gin-gonic/gin"
)

var HeaderContentLength = http.CanonicalHeaderKey("Content-Length")

const (
	DevelopmentBaseURL = "dev.cienciaargentina.com"

	RequestIDHeader = "x-request-id"

	ContentTypeJSON = "application/json"

	ContentLength = "Content-Length"
	ContentType   = "Content-Type"
)

var defaultTimeout = 500 * time.Millisecond // nolint

type Client interface {
	NewRequest(c ...*gin.Context) CienciaArgentinaRest
}

type restClient struct{} // nolint

type CienciaArgentinaRest interface {
	WithBaseURL(baseurl string) CienciaArgentinaRest
	WithResource(resource string) CienciaArgentinaRest
	WithHeader(key, value string) CienciaArgentinaRest
	WithHeaders(header map[string]string) CienciaArgentinaRest
	WithTimeout(timeout time.Duration) CienciaArgentinaRest
	WithMethod(method string) CienciaArgentinaRest
	WithBody(body interface{}) CienciaArgentinaRest
	WithQueryString(key, value string) CienciaArgentinaRest
	WithQueryStrings(query map[string]string) CienciaArgentinaRest
	PerformRequest() CienciaArgentinaResponse
}

type cienciaArgentinaRest struct {
	baseURL      string
	resource     string
	resultURL    string // this is the result string after we concat base url, resource, query strings, and so on
	queryStrings map[string]string
	body         interface{}
	reqBody      []byte
	method       string
	req          *http.Request
	client       *http.Client
	context      context.Context // nolint
}

func (r *restClient) NewRequest(c ...*gin.Context) CienciaArgentinaRest {
	ca := &cienciaArgentinaRest{}

	if scope.IsLocal() {
		ca.baseURL = "localhost:8080"
	} else if scope.IsDevelopment() {
		ca.baseURL = DevelopmentBaseURL
	}

	ca.req = &http.Request{}
	ca.req.Header = http.Header{}
	ca.client = &http.Client{}
	ca.client.Timeout = defaultTimeout

	if len(c) > 0 && c[0] != nil {
		ca.context = c[0].Request.Context()
	} else {
		ca.context = StartNewContext(context.Background())
	}

	return ca
}

func (c *cienciaArgentinaRest) WithBaseURL(baseurl string) CienciaArgentinaRest {
	c.baseURL = baseurl
	return c
}

func (c *cienciaArgentinaRest) WithResource(resource string) CienciaArgentinaRest {
	c.resource = resource
	return c
}

func (c *cienciaArgentinaRest) WithHeader(key, value string) CienciaArgentinaRest {
	if c.req != nil {
		c.req.Header.Add(key, value)
	}

	return c
}

func (c *cienciaArgentinaRest) WithHeaders(header map[string]string) CienciaArgentinaRest {
	for key, value := range header {
		c.WithHeader(key, value)
	}

	return c
}

func (c *cienciaArgentinaRest) WithTimeout(timeout time.Duration) CienciaArgentinaRest {
	if c.client != nil {
		c.client.Timeout = timeout
	}

	return c
}

func (c *cienciaArgentinaRest) WithMethod(method string) CienciaArgentinaRest {
	if c.req != nil {
		c.req.Method = method
		c.method = method
	}
	return c
}

func (c *cienciaArgentinaRest) WithBody(body interface{}) CienciaArgentinaRest {
	if body != nil {
		c.body = body
	}
	return c
}

func (c *cienciaArgentinaRest) WithQueryString(key, value string) CienciaArgentinaRest {
	if c.queryStrings == nil {
		c.queryStrings = make(map[string]string)
	}

	c.queryStrings[key] = value
	return c
}

func (c *cienciaArgentinaRest) WithQueryStrings(query map[string]string) CienciaArgentinaRest {
	for key, value := range query {
		c.WithQueryString(key, value)
	}
	return c
}

func (c *cienciaArgentinaRest) PerformRequest() CienciaArgentinaResponse {
	if c.client == nil {
		msg := "client can't be nil"
		clog.Panic(msg, "perform-request", errors.New(msg), nil)
		return nil
	}

	if c.req == nil {
		msg := "request can't be nil"
		clog.Panic(msg, "perform-request", errors.New(msg), nil)
		return nil
	}

	if c.baseURL == "" {
		msg := "baseurl can't be nil"
		clog.Panic(msg, "perform-request", errors.New(msg), nil)
		return nil
	}

	if c.resource == "" {
		msg := "resource can't be nil"
		clog.Panic(msg, "perform-request", errors.New(msg), nil)
		return nil
	}

	if c.method == "" {
		msg := "method can't be nil"
		clog.Panic(msg, "perform-request", errors.New(msg), nil)
		return nil
	}

	requestTimer := time.Now()

	res := DoRequest(c)

	r := cienciaArgentinaResponse{response: res, executionTime: time.Since(requestTimer).Milliseconds()}

	return r
}
