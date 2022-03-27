package logger

import (
	"context"
	"github.com/bombsimon/logrusr/v2"
	"github.com/go-logr/logr"
	"github.com/sirupsen/logrus"
	"point-of-sale.go/v1/internal/environment"
	"point-of-sale.go/v1/internal/request_id"
)

const (
	logrRequestIDKey = "request_id"
)

func newLogrusLogr(f logrus.Formatter) logr.Logger {
	l := logrus.New()
	l.SetFormatter(f)
	return logrusr.New(l)
}

func NewLogr(env environment.Environment) logr.Logger {
	if env.IsStructuredLogging() {
		return newLogrusLogr(&logrus.JSONFormatter{})
	}
	return newLogrusLogr(&logrus.TextFormatter{})
}

func WithRequestID(ctx context.Context, l logr.Logger) logr.Logger {
	return l.WithValues(logrRequestIDKey, request_id.GetRequestID(ctx))
}
