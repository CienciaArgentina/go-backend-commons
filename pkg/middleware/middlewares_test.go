package middleware

import (
	"errors"
	"testing"

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
