// Tasks 7, 8: Authentication and Authorization Middleware
// Task 7 — Authentication: Validates Bearer token
// Task 8 — Authorization: Checks user role (admin vs user)
//
// Demo tokens (use in Authorization header as "Bearer <token>"):
//   admin-token-123  → role: admin  (full access)
//   user-token-456   → role: user   (read-only)
//   user-token-789   → role: user   (read-only)

package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

const (
	UserKey contextKey = "user"
	RoleKey contextKey = "role"
)

// Simulated token store (in production, use JWT or session DB)
var tokenStore = map[string]struct {
	UserID int
	Name   string
	Role   string
}{
	"admin-token-123": {1, "Admin", "admin"},
	"user-token-456":  {2, "Piyush", "user"},
	"user-token-789":  {3, "Alice", "user"},
}

// Task 7: Authentication — verifies the token is valid
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			sendAuthError(w, http.StatusUnauthorized, "Authorization header required")
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			sendAuthError(w, http.StatusUnauthorized, "Bearer token required")
			return
		}

		user, ok := tokenStore[token]
		if !ok {
			sendAuthError(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		// Store user info in context
		ctx := context.WithValue(r.Context(), UserKey, user.Name)
		ctx = context.WithValue(ctx, RoleKey, user.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Task 8: Authorization — only admin can write (POST/PUT/DELETE)
func AdminWriteOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// GET requests allowed for all authenticated users
		if r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
			return
		}

		// Write operations require admin role
		role, _ := r.Context().Value(RoleKey).(string)
		if role != "admin" {
			sendAuthError(w, http.StatusForbidden,
				"Admin access required for "+r.Method+" operations")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func sendAuthError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"message": msg,
	})
}
