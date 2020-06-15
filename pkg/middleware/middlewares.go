package middleware

import (
	"fmt"
	"net/http"
	"strings"

	apiErrors "github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
	"github.com/gin-gonic/gin"
)

// Controller API controller
type Controller func(c *gin.Context) error

const (
	// ResponseCodeKey Key that's search for in gin.Context to set response status code
	ResponseCodeKey = "responseCode"
	// ResponseBodyKey Key that's search for in gin.Context to set response body
	ResponseBodyKey = "responseBody"
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
		apiCode, ok = apiCode.(int)
		if !ok {
			// TODO: Should this panic? Or set an error in context instead?
			apiCode = http.StatusOK
		}
	}

	if response, exists := c.Get(ResponseBodyKey); exists {
		c.JSON(code, response)
		return
	}

	err := apiErrors.NewInternalServerApiError("Expected a response body", nil)
	c.Error(err)
	c.Abort()
}

// ErrorMiddleware Error handling middleware
func ErrorMiddleware(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		errorMsgs := []string{}

		for _, ginErr := range c.Errors {
			apiErr, ok := ginErr.Err.(apiErrors.ApiError)

			if ok {
				errorMsgs = append(errorMsgs, apiErr.Error())
				continue
			}

			errorMsgs = append(errorMsgs, ginErr.Err.Error())
		}

		msg := fmt.Sprintf("[%s]", strings.Join(errorMsgs, ", "))
		apiErr := apiErrors.NewInternalServerApiError(msg, nil)

		c.JSON(apiErr.Status(), apiErr)
		if !c.IsAborted() {
			c.AbortWithStatus(apiErr.Status())
		}
	}
}
