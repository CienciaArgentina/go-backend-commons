package scope

import (
	"os"
	"strconv"
	"strings"
)

const (

	Scope = "SCOPE"
	IsCloudKey = "IS_CLOUD"

	Local = "local"
	Development = "development"
	Testing = "testing"
	Staging = "staging"
	Production = "production"
)

func GetScope() string {
	return os.Getenv(Scope)
}

func IsCloud() bool {
	isCloudVal := os.Getenv(IsCloudKey)
	cloud, err := strconv.ParseBool(isCloudVal)
	if err != nil {
		return false
	}
	return cloud
}

func IsProduction() bool {
	return strings.EqualFold(GetScope(), Production) && IsCloud()
}

func IsStaging() bool {
	return strings.EqualFold(GetScope(), Staging) && IsCloud()
}

func IsTesting() bool {
	return strings.EqualFold(GetScope(), Testing) && IsCloud()
}

func IsDevelopment() bool {
	return strings.EqualFold(GetScope(), Development) && IsCloud()
}

func IsLocal() bool {
	return !IsCloud() || strings.EqualFold(GetScope(), Local)
}
