package middleware

import (
	"context"
	"fmt"
	"net/http"
	"sync/atomic"
)

type contextKey string

const RequestIDKey contextKey = "requestID"

var requestCounter atomic.Uint64

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := fmt.Sprintf("req-%d", requestCounter.Add(1))
		w.Header().Set("X-Request-ID", id)

		ctx := context.WithValue(r.Context(), RequestIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
