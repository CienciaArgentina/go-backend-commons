package rest

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type requestTrace struct{}

func StartNewContext(ctx context.Context) context.Context {
	headers := http.Header{}

	headers.Add(RequestIDHeader, newRequestID())

	return context.WithValue(ctx, requestTrace{}, headers)
}

func newRequestID() string {
	uuid := uuid.New()
	return uuid.String()
}

func SetContextInformation(c *gin.Context) {
	c.Writer.Header().Add(RequestIDHeader, newRequestID())
}

func GetContextInformation(transaction string, c *gin.Context) *ContextInformation {
	return &ContextInformation{
		RequestID:       c.Writer.Header().Get(RequestIDHeader),
		TransactionName: transaction,
	}
}

type ContextInformation struct {
	RequestID       string
	TransactionName string
}
