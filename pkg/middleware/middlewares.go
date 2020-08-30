package middleware

import (
	"net/http"

	apiErrors "github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
	"github.com/gin-gonic/gin"
)

// Controller API controller
type Controller func(c *gin.Context) error

const (
	// ResponseCodeKey Key that's searched for in gin.Context to set response status code
	ResponseCodeKey = "responseCode"
	// ResponseBodyKey Key that's searched for in gin.Context to set response body
	ResponseBodyKey = "responseBody"
	// codeNoResponseBody No response body code
	codeNoResponseBody = "no_response_body"
	// codeInternalError Internal error code
	codeInternalError = "internal_error"
)

// AdaptController Adapts controller to gin middleware
func AdaptController(handler Controller) func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := handler(c); err != nil {
			c.Error(err)
			c.Abort()
		}
	}
}

// ResponseMiddleware Response middleware
func ResponseMiddleware(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		return
	}

	code := http.StatusOK
	if apiCode, exists := c.Get(ResponseCodeKey); exists {
		var ok bool
		code, ok = apiCode.(int)
		if !ok {
			// TODO: Should this panic? Or set an error in context instead?
			apiCode = http.StatusOK
		}
	}

	if response, exists := c.Get(ResponseBodyKey); exists {
		c.JSON(code, response)
		return
	}

	err := apiErrors.NewInternalServerApiError("Expected a response body", nil, codeNoResponseBody)
	c.AbortWithError(err.Status(), err)
}

// ErrorMiddleware Error handling middleware
func ErrorMiddleware(c *gin.Context) {
	c.Next()
	if len(c.Errors) == 0 {
		return
	}

	ginErr := c.Errors.Last()
	apiErr := apiErrors.NewInternalServerApiError("Internal server error", ginErr, codeInternalError)
	if ginAPIErr, ok := ginErr.Err.(apiErrors.ApiError); ok {
		apiErr = ginAPIErr
	}

	c.JSON(apiErr.Status(), apiErr)
	if !c.IsAborted() {
		c.AbortWithStatusJSON(apiErr.Status(), apiErr)
	}
}
