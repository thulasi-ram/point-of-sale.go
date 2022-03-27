package request_id

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

const (
	ContextKeyRequestID = "requestID"
	RequestIDHeader     = "X-Request-Id"
)

func SetRequestID(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, ContextKeyRequestID, reqID)
}

func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if reqID, ok := ctx.Value(ContextKeyRequestID).(string); ok {
		return reqID
	}
	return ""
}

func GetRequestIDFromRequest(r *http.Request) string {
	requestID := r.Header.Get(RequestIDHeader)
	if requestID == "" {
		requestID = uuid.NewString()
	}
	return requestID
}
