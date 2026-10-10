// Package middleware contains role-based authorization middleware.
package middleware

import "net/http"

// AdminWrites allows GET requests to all authenticated users but restricts write methods to admins.
func AdminWrites(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
			return
		}
		role, _ := r.Context().Value(roleKey).(string)
		if role != "admin" {
			http.Error(w, `{"error":"forbidden: admin role required"}`, http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
