package middleware

import (
	"context"
	"net/http"
	"strings"
)

const roleKey contextKey = "role"

func RequireAuth(userToken, adminToken string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Fields(r.Header.Get("Authorization"))
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
			return
		}

		role := ""
		switch parts[1] {
		case adminToken:
			if adminToken != "" {
				role = "admin"
			}
		case userToken:
			if userToken != "" {
				role = "user"
			}
		}

		if role == "" {
			http.Error(w, `{"error":"invalid token"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), roleKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminWrites(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet || r.Method == http.MethodHead {
			next.ServeHTTP(w, r)
			return
		}

		role, _ := r.Context().Value(roleKey).(string)
		if role != "admin" {
			http.Error(w, `{"error":"admin role required"}`, http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
