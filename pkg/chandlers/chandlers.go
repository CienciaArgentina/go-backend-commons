package chandlers

import (
	"fmt"
	"github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
	"github.com/CienciaArgentina/go-backend-commons/pkg/scope"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CienciaArgentinaRouter() *gin.Engine {
	r := gin.New()
	if !scope.IsProductiveScope() {
		r.Use(gin.Logger())
	}
	pprof.Register(r)
	r.NoRoute(noRouteHandler)
	return r
}

func noRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, apierror.NewNotFoundApiError(fmt.Sprintf("Resource %s not found", c.Request.URL.Path)))
}