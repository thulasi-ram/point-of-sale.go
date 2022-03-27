package middleware

import (
	"net/http"
	"point-of-sale.go/v1/internal/request_id"
)

func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		reqID := request_id.GetRequestIDFromRequest(r)
		ctx := request_id.SetRequestID(r.Context(), reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
