package basefilter

import (
	"bufio"
	"bytes"
	"github.com/CienciaArgentina/go-backend-commons/pkg/rest"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"strconv"
)

type baseWriter struct {
	ctx *gin.Context
	body *bytes.Buffer
	response gin.ResponseWriter
	size int
}

func BaseFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		var baseWriter *baseWriter

		// Check if we're using gin, otherwise we ignore it
		if wr, ok := c.Writer.(gin.ResponseWriter); ok {
			baseWriter.ctx = c
			baseWriter.body = &bytes.Buffer{}
			baseWriter.response = wr
			c.Writer = baseWriter
			c.Next()
		} else {
			c.Next()
			return
		}

		bodyBytes := baseWriter.body.Bytes()
		baseWriter.size = len(bodyBytes)
		baseWriter.response.Write(bodyBytes)

		baseWriter.Header().Set(rest.HeaderContentLength, strconv.Itoa(baseWriter.size))
	}
}

func (b *baseWriter) Flush() {}

func (b *baseWriter) WriteHeaderNow() {}

func (b *baseWriter) Header() http.Header {
	return b.response.Header()
}

func (b *baseWriter) Write([]byte) (int, error) {
	if b.Status() >= http.StatusInternalServerError {

	}
}

func (b *baseWriter) WriteHeader(statusCode int) {
	panic("implement me")
}

func (b *baseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	panic("implement me")
}

func (b *baseWriter) CloseNotify() <-chan bool {
	panic("implement me")
}

func (b *baseWriter) Status() int {
	panic("implement me")
}

func (b *baseWriter) Size() int {
	panic("implement me")
}

func (b *baseWriter) WriteString(string) (int, error) {
	panic("implement me")
}

func (b *baseWriter) Written() bool {
	panic("implement me")
}

func (b *baseWriter) Pusher() http.Pusher {
	panic("implement me")
}

func respondServerError(c *gin.Context, bodybytes []byte, status int) []byte {

}