package main

import (
	"net/http"
)

// Task 8: Role-based Authorization Interceptor (AuthZ)
// Assumes AuthN has verified authentication status and assesses access privileges.

type UserCtxKey string
const UserRoleKey UserCtxKey = "user_role"

func AuthZMiddleware(requiredRole string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value(UserRoleKey).(string)
		if !ok || role != requiredRole {
			http.Error(w, "Forbidden: Insufficient privileges", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {}
