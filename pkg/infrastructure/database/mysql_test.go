package database

import (
	"github.com/CienciaArgentina/go-backend-commons/config"
	"github.com/CienciaArgentina/go-backend-commons/pkg/scope"
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"
)

const (
	FilePath = "../../../config/config.development.yaml"
)

func SetEnvVariables() {
	os.Setenv("DB_USERNAME", "user")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_HOSTNAME", "hostname")
	os.Setenv("DB_PORT", "port")
	os.Setenv("DB_NAME", "name")
}

func TestNewShouldPanicIfConfigIsNil(t *testing.T) {
	require.Panics(t, func() {
		New(nil)
	})
}

func TestNewShouldReturnMockWhenScopeIsTesting(t *testing.T) {
	SetEnvVariables()
	tmpScope := scope.GetScope()
	tmpCloud := scope.IsCloud()
	os.Setenv(scope.Scope, scope.Testing)
	os.Setenv(scope.IsCloudKey, "true")
	cfg, _ := config.New()
	db, _ := New(cfg.Database)
	os.Setenv(scope.Scope, tmpScope)
	os.Setenv(scope.IsCloudKey, strconv.FormatBool(tmpCloud))
	require.NotNil(t, db.Mock)
}

func TestNewShouldPanicPingingDb(t *testing.T) {
	SetEnvVariables()
	cfg, _ := config.New()
	require.Panics(t, func() {
		New(cfg.Database)
	})
}
