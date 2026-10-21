package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

// Task 5: Request-ID generating middleware using context
// Context propagation passes request-scoped telemetry and deadlines down the stack.

type contextKey string
const RequestIDKey contextKey = "request_id"

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bytes := make([]byte, 8)
		_, _ = rand.Read(bytes)
		reqID := hex.EncodeToString(bytes)
		
		// Inject into context
		ctx := context.WithValue(r.Context(), RequestIDKey, reqID)
		w.Header().Set("X-Request-ID", reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {}
