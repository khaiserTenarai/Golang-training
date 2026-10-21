package main

import (
	"context"
	"net/http"
)

// Task 7: Authentication Token Validator (AuthN)
// Middleware assessing identity validity and populating user status inside context.

type UserCtxKey string
const UserRoleKey UserCtxKey = "user_role"

func AuthNMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer secret-token-123" {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}
		// Validated role extracted
		ctx := context.WithValue(r.Context(), UserRoleKey, "admin")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {}
