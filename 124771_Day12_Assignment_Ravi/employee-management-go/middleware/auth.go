package middleware

import (
	"context"

	"employee-management/auth"
	"net/http"
)

type contextKey string

const userKey contextKey = "authenticated-user"
const roleKey contextKey = "authenticated-role"

// RequireAuth requires a valid JWT bearer token.
func RequireAuth(a *auth.AuthService, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, role, err := a.Parse(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userKey, user)
		ctx = context.WithValue(ctx, roleKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequireRole permits only users with one of the supplied roles.
func RequireRole(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role, _ := r.Context().Value(roleKey).(string)
			for _, allowed := range roles {
				if role == allowed {
					next.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, `{"error":"forbidden"}`, http.StatusForbidden)
		})
	}
}
