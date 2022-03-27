package middleware

import (
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-logr/logr"
	"net/http"
	"point-of-sale.go/v1/internal/logger"
	"time"
)

type loggerEntry struct {
	log logr.Logger
}

func (e *loggerEntry) Write(status, _ int, _ http.Header, elapsed time.Duration, _ interface{}) {
	e.log.Info(
		"Request Processed",
		"resp_status", status,
		//"resp_bytes_length", bytes,
		"duration_ms", float64(elapsed.Nanoseconds())/1000000.0,
		//"headers", header,
		//"extra", extra,
	)
}

func (e *loggerEntry) Panic(v interface{}, stack []byte) {
	e.log = e.log.WithValues(
		"stack", string(stack),
		"panic", fmt.Sprintf("%+v", v),
	)
}

func NewRequestLogFormatter(log logr.Logger) *responseLogger {
	return &responseLogger{logr: log}
}

type responseLogger struct {
	logr logr.Logger
}

func (l *responseLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &loggerEntry{log: l.logr}

	entry.log = logger.WithRequestID(r.Context(), entry.log)

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	entry.log.Info(
		"Request Initiated",
		"uri", fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI),
		//"user_agent", r.UserAgent(),
		//"http_scheme", scheme,
		//"http_proto", r.Proto,
		"http_method", r.Method,
		//"remote_addr", r.RemoteAddr,
		"timestamp", time.Now().UTC().Format(time.RFC3339Nano),
	)

	return entry
}
