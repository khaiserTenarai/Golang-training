package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
)

type contextKey string

const RequestIDKey contextKey = "requestID"

func generateID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = generateID()
		}

		// Attach to outgoing HTTP response headers
		w.Header().Set("X-Request-ID", reqID)

		// Inject request-id into request context
		ctx := context.WithValue(r.Context(), RequestIDKey, reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		reqID, _ := r.Context().Value(RequestIDKey).(string)
		w.Write([]byte("Request ID: " + reqID))
	})

	http.ListenAndServe(":8080", requestIDMiddleware(mux))
}