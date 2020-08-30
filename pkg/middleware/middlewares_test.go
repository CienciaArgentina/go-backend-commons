package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	apiErrors "github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
	"github.com/gin-gonic/gin"
)

func TestAdaptControllerOk(t *testing.T) {
	ctx := gin.Context{}

	okHandler := func(c *gin.Context) error {
		return nil
	}
	adaptedOkHandler := AdaptController(okHandler)

	adaptedOkHandler(&ctx)
	if len(ctx.Errors) > 0 {
		t.Errorf("Context shouldn't have failed, returned %+v", ctx.Errors)
	}
}

func TestAdaptControllerFail(t *testing.T) {
	ctx := gin.Context{}

	failHandler := func(c *gin.Context) error {
		return errors.New("This should always fail")
	}
	adaptedFailHandler := AdaptController(failHandler)

	adaptedFailHandler(&ctx)

	if len(ctx.Errors) == 0 {
		t.Error("Context should have errors")
	}
}

func TestResponseMiddlewareCtxWithError(t *testing.T) {
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(ResponseMiddleware)
	r.GET("/test", func(c *gin.Context) {
		c.Error(errors.New("This should fail"))
	})

	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(resp, c.Request)

	if resp.Body.Bytes() != nil {
		t.Error("Middleware should not handle body when theres an error in ctx")
	}
}

func TestResponseMiddlewareStatusCode(t *testing.T) {
	resp := httptest.NewRecorder()
	expectedStatusCode := http.StatusTemporaryRedirect
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(ResponseMiddleware)
	r.GET("/test", func(c *gin.Context) {
		c.Set(ResponseCodeKey, expectedStatusCode)
		c.Set(ResponseBodyKey, map[string]string{"test": "test"})
	})

	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(resp, c.Request)

	if resp.Result().StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d but got %d", expectedStatusCode, resp.Result().StatusCode)
	}
}

func TestResponseMiddlewareStatusCodeWrong(t *testing.T) {
	resp := httptest.NewRecorder()
	expectedStatusCode := http.StatusOK
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(ResponseMiddleware)
	r.GET("/test", func(c *gin.Context) {
		c.Set(ResponseCodeKey, "This isn't an int!")
		c.Set(ResponseBodyKey, map[string]string{"test": "test"})
	})

	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(resp, c.Request)

	if resp.Result().StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d but got %d", expectedStatusCode, resp.Result().StatusCode)
	}
}

func TestResponseMiddlewareNoBody(t *testing.T) {
	resp := httptest.NewRecorder()
	expectedStatusCode := http.StatusInternalServerError
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(ResponseMiddleware)
	r.GET("/test", func(c *gin.Context) {
	})

	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(resp, c.Request)

	if resp.Result().StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d but got %d", expectedStatusCode, resp.Result().StatusCode)
	}
}

func TestErrorMiddlewareNoError(t *testing.T) {
	resp := httptest.NewRecorder()
	expectedStatusCode := http.StatusOK
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(ErrorMiddleware)
	r.GET("/test", func(c *gin.Context) {
	})

	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(resp, c.Request)

	if resp.Result().StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d but got %d", expectedStatusCode, resp.Result().StatusCode)
	}
}

func TestErrorMiddlewareGinError(t *testing.T) {
	resp := httptest.NewRecorder()
	expectedStatusCode := http.StatusInternalServerError
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(ErrorMiddleware)
	r.GET("/test", func(c *gin.Context) {
		c.Error(errors.New("This isn't an API error"))
	})

	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(resp, c.Request)

	if resp.Result().StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d but got %d", expectedStatusCode, resp.Result().StatusCode)
	}
}

func TestErrorMiddlewareAPIError(t *testing.T) {
	resp := httptest.NewRecorder()
	expectedStatusCode := http.StatusForbidden
	expectedBody := `{"status":403,"message":"You shall not pass!","error":[{"detail":"You shall not pass!","code":"forbidden"}]}`
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(ErrorMiddleware)
	r.GET("/test", func(c *gin.Context) {
		err := apiErrors.NewForbiddenApiError("You shall not pass!")
		c.Error(err)
	})

	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(resp, c.Request)

	if resp.Result().StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d but got %d", expectedStatusCode, resp.Result().StatusCode)
	}
	if reflect.DeepEqual(resp.Body.Bytes(), expectedBody) {
		t.Errorf("Expected body %s but got %+v", expectedBody, string(resp.Body.Bytes()))
	}
}
