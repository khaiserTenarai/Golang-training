package middleware

import (
	"context"
	"log"
	"net/http"
	"time"
)

type ctxKey string
const RequestIDKey ctxKey = "request_id"

func TelemetryAndSecurityChain(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// AuthN Simulator
		token := r.Header.Get("Authorization")
		var userRole string
		if token == "Bearer admin-token" {
			userRole = "admin"
		} else {
			userRole = "guest"
		}

		// Inject Request ID and context values
		ctx := context.WithValue(r.Context(), RequestIDKey, "REQ-10015")
		r = r.WithContext(ctx)

		// Path Based AuthZ Enforcement Interceptor
		if r.URL.Path == "/employees/create" && userRole != "admin" {
			http.Error(w, "Forbidden - Administrator profile authorization failed", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
		
		duration := time.Since(start).Microseconds()
		log.Printf("[TELEMETRY] Method: %s | Path: %s | Latency: %d microseconds", r.Method, r.URL.Path, duration)
	})
}
