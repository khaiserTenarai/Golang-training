package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"employee-management/auth"
)

type contextKey string

const (
	userKey contextKey = "authenticated-user"
	roleKey contextKey = "authenticated-role"
)

func UserFromContext(ctx context.Context) string {
	v, _ := ctx.Value(userKey).(string)
	return v
}

func RoleFromContext(ctx context.Context) string {
	v, _ := ctx.Value(roleKey).(string)
	return v
}

func writeAuthError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// RequireAuth verifies the bearer token before calling the protected handler.
func RequireAuth(a *auth.AuthService, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := a.ParseAuthorization(r.Header.Get("Authorization"))
		if err != nil {
			writeAuthError(w, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), userKey, claims.Subject)
		ctx = context.WithValue(ctx, roleKey, claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequireRole restricts the wrapped handler to the supplied roles.
func RequireRole(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role := RoleFromContext(r.Context())
			for _, allowed := range roles {
				if role == allowed {
					next.ServeHTTP(w, r)
					return
				}
			}
			writeAuthError(w, http.StatusForbidden, "forbidden")
		})
	}
}

// IsMutationRequest is useful when enforcing role checks on GraphQL operations.
func IsMutationRequest(query string) bool {
	return strings.HasPrefix(strings.TrimSpace(query), "mutation")
}
