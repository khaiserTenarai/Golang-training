package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strings"
)

type contextKey string

const RequestIDKey contextKey = "request_id"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimSpace(r.Header.Get("X-Request-ID"))
		if id == "" || len(id) > 128 {
			buffer := make([]byte, 16)
			if _, err := rand.Read(buffer); err != nil {
				id = "request-id-unavailable"
			} else {
				id = hex.EncodeToString(buffer)
			}
		}
		w.Header().Set("X-Request-ID", id)
		ctx := context.WithValue(r.Context(), RequestIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
