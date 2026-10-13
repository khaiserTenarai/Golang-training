package middleware

import (
	"context"
	"net/http"
	"strings"

	"employee-api-go/internal/auth"
)

type ClaimsKeyType string

const ClaimsKey ClaimsKeyType = "claims"

func Authentication(manager *auth.TokenManager, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		parts := strings.Fields(header)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			writeAuthError(w, http.StatusUnauthorized, "missing or invalid bearer token")
			return
		}
		claims, err := manager.Verify(parts[1])
		if err != nil {
			writeAuthError(w, http.StatusUnauthorized, "invalid or expired token")
			return
		}
		ctx := context.WithValue(r.Context(), ClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func writeAuthError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = jsonEncode(w, map[string]string{"error": message})
}
