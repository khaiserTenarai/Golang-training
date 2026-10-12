package main

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const UserContextKey contextKey = "user"

func authnMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized: missing bearer token", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		// Simulate JWT parsing and validation
		if token != "valid-secret-token" {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, "admin_user")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /protected", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserContextKey).(string)
		w.Write([]byte("Hello, " + user))
	})

	http.ListenAndServe(":8080", authnMiddleware(mux))
}