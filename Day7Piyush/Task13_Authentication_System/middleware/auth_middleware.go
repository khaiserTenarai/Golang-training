package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"task13_authentication_system/models"
	"task13_authentication_system/service"
)

type contextKey string

const (
	UserIDKey contextKey = "userID"
	RoleKey   contextKey = "role"
)

func AuthMiddleware(authService *service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				sendJSON(w, http.StatusUnauthorized, models.Response{Message: "Authorization header required"})
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			if token == authHeader {
				sendJSON(w, http.StatusUnauthorized, models.Response{Message: "Bearer token required"})
				return
			}

			userID, role, err := authService.ValidateToken(token)
			if err != nil {
				sendJSON(w, http.StatusUnauthorized, models.Response{Message: "Invalid or expired token"})
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			ctx = context.WithValue(ctx, RoleKey, role)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value(RoleKey).(string)
		if !ok || role != "admin" {
			sendJSON(w, http.StatusForbidden, models.Response{Message: "Admin access required"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func sendJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
