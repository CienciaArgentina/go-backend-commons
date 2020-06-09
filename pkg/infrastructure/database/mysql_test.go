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
	FilePath = "../../../config/config.development.yml"
)

func TestNewShouldPanicIfConfigIsNil(t *testing.T) {
	require.Panics(t, func() {
		New(nil)
	})
}

func TestNewShouldReturnMockWhenScopeIsTesting(t *testing.T) {
	tmpScope := scope.GetScope()
	tmpCloud := scope.IsCloud()
	os.Setenv(scope.Scope, scope.Testing)
	os.Setenv(scope.IsCloudKey, "true")
	cfg := config.New(&config.Options{
		FilePath: FilePath,
	})
	db, _ := New(&cfg.Database[0])
	os.Setenv(scope.Scope, tmpScope)
	os.Setenv(scope.IsCloudKey, strconv.FormatBool(tmpCloud))
	require.NotNil(t, db.Mock)
}

func TestNewShouldPanicWhenTheresNoPasswordInEnvironment(t *testing.T) {
	cfg := config.New(&config.Options{
		FilePath: FilePath,
	})
	require.Panics(t, func() {
		New(&cfg.Database[0])
	})
}

func TestNewShouldPanicPingingDb(t *testing.T) {
	os.Setenv("db_pass", "test")
	os.Setenv("db_host", "test")
	cfg := config.New(&config.Options{
		FilePath: FilePath,
	})
	require.Panics(t, func() {
		New(&cfg.Database[0])
	})
	os.Setenv("db_pass", "")
	os.Setenv("db_host", "")
}