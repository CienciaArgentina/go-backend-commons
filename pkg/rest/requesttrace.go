package rest

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type requestTrace struct{}

func StartNewContext(ctx context.Context) context.Context {
	headers := http.Header{}

	headers.Add(RequestIDHeader, NewRequestID())

	return context.WithValue(ctx, requestTrace{}, headers)
}

func NewRequestID() string {
	uuid := uuid.New()
	return uuid.String()
}

