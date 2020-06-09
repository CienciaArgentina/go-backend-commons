package scope

import (
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
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
	tmpCloud := IsCloud()
	os.Setenv(Scope, Production)
	os.Setenv(IsCloudKey, "TRUE")

	isProd := IsProduction()

	os.Setenv(Scope, tmp)
	os.Setenv(IsCloudKey, strconv.FormatBool(tmpCloud))
	require.True(t, isProd)
}

func TestIsStagingShouldReturnTrue(t *testing.T) {
	tmp := GetScope()
	tmpCloud := IsCloud()
	os.Setenv(Scope, Staging)
	os.Setenv(IsCloudKey, "TRUE")

	isStaging := IsStaging()

	os.Setenv(Scope, tmp)
	os.Setenv(IsCloudKey, strconv.FormatBool(tmpCloud))
	require.True(t, isStaging)
}

func TestIsTestingShouldReturnTrue(t *testing.T) {
	tmp := GetScope()
	tmpCloud := IsCloud()
	os.Setenv(Scope, Testing)
	os.Setenv(IsCloudKey, "TRUE")

	isTesting := IsTesting()

	os.Setenv(Scope, tmp)
	os.Setenv(IsCloudKey, strconv.FormatBool(tmpCloud))
	require.True(t, isTesting)
}

func TestIsDevelopmentShouldReturnTrue(t *testing.T) {
	tmp := GetScope()
	tmpCloud := IsCloud()
	os.Setenv(Scope, Development)
	os.Setenv(IsCloudKey, "TRUE")

	isDevelopment := IsDevelopment()

	os.Setenv(Scope, tmp)
	os.Setenv(IsCloudKey, strconv.FormatBool(tmpCloud))
	require.True(t, isDevelopment)
}

func TestIsLocalShouldReturnTrue(t *testing.T) {
	tmp := GetScope()
	os.Setenv(Scope, Local)

	isLocal := IsLocal()

	os.Setenv(Scope, tmp)
	require.True(t, isLocal)
}

func TestIsProductiveScope(t *testing.T) {
	require.False(t, IsProductiveScope())
}
