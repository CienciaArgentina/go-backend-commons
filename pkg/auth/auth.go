package auth

import (
	"net/http"
	"strconv"
	"strings"
)

type Auth struct {
	JWT string
}

var (
	IsPublic = http.CanonicalHeaderKey("Is-Public")
)

func IsPublicRequest(request *http.Request) bool {
	isPublic, _ :=  strconv.ParseBool(strings.ToLower(request.Header.Get(IsPublic)))
	return isPublic
}

type CheckClaimBody struct {
	JWT string `json:"jwt"`
	RequiredClaim string `json:"required_claim"`
}