package main

import (
	"context"
	"net/http"
)

type contextKey string

const RoleKey contextKey = "role"

func requireRole(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRole, ok := r.Context().Value(RoleKey).(string)
		if !ok || userRole != role {
			http.Error(w, "Forbidden: insufficient permissions", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}

func mockUserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), RoleKey, "viewer")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func deleteDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Resource deleted successfully"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /data", requireRole("admin", deleteDataHandler))

	http.ListenAndServe(":8080", mockUserContext(mux))
}