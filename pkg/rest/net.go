package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/CienciaArgentina/go-backend-commons/pkg/clog"
)

func DoRequest(request *cienciaArgentinaRest) *response { //nolint
	if request != nil {
		response := new(response)
		var resp *http.Response
		request = createRequest(request)

		resp, err := request.client.Do(request.req)
		if resp.Body != nil {
			defer resp.Body.Close()
		}
		response.response = resp
		response.Err = err
		return response
	}

	return nil
}

func createRequest(request *cienciaArgentinaRest) *cienciaArgentinaRest {
	if request == nil {
		msg := "ciencia argentina request can't be nil"
		clog.Panic(msg, "create-request", errors.New(msg), nil)
	}

	request.resultURL = request.resource

	if request.queryStrings != nil && len(request.queryStrings) > 0 {
		request.resultURL = formatQueryString(request.resultURL, request.queryStrings)
	}

	request.resultURL = fmt.Sprintf("%s%s", request.baseURL, request.resultURL)

	if request.body != nil {
		request = parseBody(request)
	}

	request.WithHeader(ContentType, ContentTypeJSON)
	return request
}

func formatQueryString(path string, query map[string]string) string {
	q := make(url.Values)
	for k, v := range query {
		q.Add(k, v)
	}

	if len(q) > 0 {
		path = fmt.Sprintf("%s?%s", path, q.Encode())
	}

	return path
}

func parseBody(request *cienciaArgentinaRest) *cienciaArgentinaRest {
	b, err := json.Marshal(request.body)
	if err != nil {
		clog.Panic(err.Error(), "parse-body", err, nil)
		return nil
	}
	request.reqBody = b
	request.req.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	contentLen := strconv.Itoa(len(b))

	request.WithHeader(ContentLength, contentLen)
	return request
}

