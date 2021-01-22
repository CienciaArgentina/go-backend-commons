package performance

import (
	"fmt"
	"time"

	"github.com/CienciaArgentina/go-backend-commons/pkg/middleware"

	"github.com/CienciaArgentina/go-backend-commons/pkg/clog"
)

func TrackTime(start time.Time, functionName string, ctx *middleware.ContextInformation, f func()) {
	f()
	clog.Debug(fmt.Sprintf("Logueo de tiempo de %s", functionName), "log-time",
		map[string]string{"duration": fmt.Sprintf("%d ms", time.Since(start).Milliseconds()), "request-id": ctx.RequestID, "transaction": ctx.TransactionName})
}

