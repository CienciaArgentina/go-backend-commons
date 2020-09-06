package rest

import (
	"context"
	"net/http"

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
