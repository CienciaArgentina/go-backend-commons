package scope

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGetScopeShouldReturnLocal(t *testing.T) {
	os.Setenv(Scope, Local)
	scope := GetScope()
	require.Equal(t, Local, scope)
}

func TestIsCloudShouldReturnFalse(t *testing.T) {
	os.Setenv(IsCloudKey, "")
	isCloud := IsCloud()
	require.False(t, isCloud)
}

func TestIsProductionShouldReturnTrue(t *testing.T) {
	tmp := GetScope()
	os.Setenv(Scope, Production)

	isProd := IsProduction()

	os.Setenv(Scope, tmp)
	require.True(t, isProd)
}

func TestIsStagingShouldReturnTrue(t *testing.T) {
	tmp := GetScope()
	os.Setenv(Scope, Staging)

	isStaging := IsStaging()

	os.Setenv(Scope, tmp)
	require.True(t, isStaging)
}

func TestIsTestingShouldReturnTrue(t *testing.T) {
	tmp := GetScope()
	os.Setenv(Scope, Testing)

	isTesting := IsTesting()

	os.Setenv(Scope, tmp)
	require.True(t, isTesting)
}

func TestIsDevelopmentShouldReturnTrue(t *testing.T) {
	tmp := GetScope()
	os.Setenv(Scope, Development)

	isDevelopment := IsDevelopment()

	os.Setenv(Scope, tmp)
	require.True(t, isDevelopment)
}

func TestIsLocalShouldReturnTrue(t *testing.T) {
	tmp := GetScope()
	os.Setenv(Scope, Local)

	isLocal := IsLocal()

	os.Setenv(Scope, tmp)
	require.True(t, isLocal)
}