package auth

import (
	"net/http"
	"strconv"
	"strings"
)

var (
	IsPublic = http.CanonicalHeaderKey("Is-Public")
)

func IsPublicRequest(request *http.Request) bool {
	isPublic, _ :=  strconv.ParseBool(strings.ToLower(request.Header.Get(IsPublic)))
	return isPublic
}