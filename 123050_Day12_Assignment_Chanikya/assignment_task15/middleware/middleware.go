package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"time"
)

type contextKey string

const (
	RequestIDKey contextKey = "request_id"
	UserRoleKey  contextKey = "user_role"
)

// RequestID generates dynamic metadata trackers passed through context boundaries
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bytes := make([]byte, 8)
		_, _ = rand.Read(bytes)
		reqID := hex.EncodeToString(bytes)
		ctx := context.WithValue(r.Context(), RequestIDKey, reqID)
		w.Header().Set("X-Request-ID", reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Logger measures operations in accurate microsecond durations
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		reqID, _ := r.Context().Value(RequestIDKey).(string)
		
		next.ServeHTTP(w, r)
		
		duration := time.Since(start).Microseconds()
		log.Printf("[REQ_ID: %s] %s %s took %d microseconds", reqID, r.Method, r.URL.Path, duration)
	})
}

// AuthN mocks identity validation checking token authenticity
func AuthN(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		// For demo purposes, if missing or invalid, default to standard user or block
		if token == "" {
			http.Error(w, "Unauthorized token check failed", http.StatusUnauthorized)
			return
		}
		
		role := "user"
		if token == "Bearer admin-token" {
			role = "admin"
		}
		
		ctx := context.WithValue(r.Context(), UserRoleKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AuthZ validates scope access permissions matching assigned roles
func AuthZ(requiredRole string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, _ := r.Context().Value(UserRoleKey).(string)
		if role != requiredRole {
			http.Error(w, "Forbidden application tier permissions restriction", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
