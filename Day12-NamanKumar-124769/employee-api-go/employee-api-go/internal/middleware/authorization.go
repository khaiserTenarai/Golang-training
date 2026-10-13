package middleware

import (
	"net/http"

	"employee-api-go/internal/auth"
)

func RequireRole(role string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(ClaimsKey).(auth.Claims)
		if !ok || claims.Role != role {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			_ = jsonEncode(w, map[string]string{"error": "insufficient permissions"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
