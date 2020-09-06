package rest

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"runtime"
	"strconv"

	"github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
	"github.com/CienciaArgentina/go-backend-commons/pkg/auth"
	"github.com/CienciaArgentina/go-backend-commons/pkg/clog"
	"github.com/CienciaArgentina/go-backend-commons/pkg/scope"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	Gzip bool
}

func CustomCienciaArgentinaRouter(cfg RouterConfig) *gin.Engine {
	r := gin.New()
	r.Use(apiFilter())
	r.Use(recoverPanic(clog.GetOut()))

	if !scope.IsProductiveScope() {
		r.Use(gin.Logger())
	}

	if cfg.Gzip {
		r.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	r.NoRoute(noRouteHandler)
	return r
}

type customWritter struct {
	ctx          *gin.Context
	response     gin.ResponseWriter
	body         *bytes.Buffer
	writtenBytes int
}

func (c *customWritter) Header() http.Header {
	return c.response.Header()
}

func (c *customWritter) Write(b []byte) (int, error) {
	if c.Status() >= http.StatusInternalServerError {
		b = internalServerErrorHandler(c.ctx, b, c.Status())
	}

	size, err := c.body.Write(b)
	c.writtenBytes += size
	return size, err
}

func (c *customWritter) WriteHeader(statusCode int) {
	c.response.WriteHeader(statusCode)
}

func (c *customWritter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return c.response.(http.Hijacker).Hijack()
}

func (c *customWritter) Flush() {}

func (c *customWritter) CloseNotify() <-chan bool {
	return c.response.(http.CloseNotifier).CloseNotify()
}

func (c *customWritter) Status() int {
	return c.response.Status()
}

func (c *customWritter) Size() int {
	return c.writtenBytes
}

func (c *customWritter) WriteString(s string) (int, error) {
	return c.Write([]byte(s))
}

func (c *customWritter) Written() bool {
	return c.body.Len() != -1
}

func (c *customWritter) WriteHeaderNow() {}

func (c *customWritter) Pusher() http.Pusher {
	return c.response.Pusher()
}

func apiFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cw *customWritter

		if writter, ok := c.Writer.(gin.ResponseWriter); ok {
			cw = &customWritter{ctx: c, response: writter, body: &bytes.Buffer{}}
			c.Writer = cw
			c.Next()
		} else {
			c.Next()
			return
		}

		bodyBytes := cw.body.Bytes()
		if len(bodyBytes) > 0 {
			cw.body.Reset()
			cw.writtenBytes = len(bodyBytes)
			cw.response.Write(bodyBytes) // nolint
		}
		cw.Header().Set(HeaderContentLength, strconv.Itoa(cw.writtenBytes))
	}
}

func internalServerErrorHandler(c *gin.Context, body []byte, status int) []byte {
	if auth.IsPublicRequest(c.Request) {
		msg, _ := json.Marshal(apierror.NewInternalServerApiError("Ocurrió un error en el servidor, por favor intentar nuevamente", nil, "internal_server_error"))
		return msg
	}

	apierror, _ := apierror.NewErrorFromBytesWithStatus(body, status)
	parsedError, _ := json.Marshal(apierror)
	return parsedError
}

func recoverPanic(o io.Writer) gin.HandlerFunc {
	var logger *log.Logger
	if o != nil {
		logger = log.New(o, "\n\n\x1b[31m", log.LstdFlags)
	}
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var pInfo string
				if logger != nil {
					stack := stackTrace(3)
					request, _ := httputil.DumpRequest(c.Request, false)
					pInfo := fmt.Sprintf("[RECOVERY] Panic recovered: \n%s\n%s\n%s%s", string(request), err, stack, []byte{27, 91, 48, 109})
					logger.Printf(pInfo)
				}

				c.AbortWithStatusJSON(http.StatusInternalServerError, apierror.NewInternalServerApiError("Error interno del servidor, panic", errors.New(pInfo), "internal_server_error"))
			}
		}()
		c.Next()
	}
}

func stackTrace(skip int) []byte {
	buf := new(bytes.Buffer)
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return []byte{}
	}

	name := []byte(fn.Name())
	if lastslash := bytes.LastIndex(name, []byte{'/'}); lastslash >= 0 {
		name = name[lastslash+1:]
	}
	if period := bytes.Index(name, []byte{'.'}); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, []byte{'·'}, []byte{'.'}, -1)
	return name
}

func source(lines [][]byte, n int) []byte {
	n--
	if n < 0 || n >= len(lines) {
		return []byte{}
	}
	return bytes.TrimSpace(lines[n])
}

func noRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, apierror.NewNotFoundApiError(fmt.Sprintf("No se encontró el recurso %s", c.Request.URL.Path)))
}
