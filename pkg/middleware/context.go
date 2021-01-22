package middleware

import (
	"github.com/CienciaArgentina/go-backend-commons/pkg/auth"
	"github.com/CienciaArgentina/go-backend-commons/pkg/rest"
	"github.com/gin-gonic/gin"
	"strings"
)

func SetContextInformation(c *gin.Context) {
	c.Writer.Header().Add(rest.RequestIDHeader, rest.NewRequestID())
}

func GetContextInformation(transaction string, c *gin.Context) *ContextInformation {
	auth := auth.Auth{JWT: getJWT(c)}
	return &ContextInformation{
		RequestID:       c.Writer.Header().Get(rest.RequestIDHeader),
		TransactionName: transaction,
		Auth: auth,
	}
}

type ContextInformation struct {
	RequestID       string
	TransactionName string
	Auth auth.Auth
}

func getJWT(c *gin.Context) string {
	header := c.Request.Header.Get("Authorization")
	if header == "" {
		return ""
	}
	split := strings.Split(header, "Bearer ")
	jwt := split[1]
	return jwt
}