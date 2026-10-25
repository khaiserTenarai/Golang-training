package graphql

import (
	"net/http"
	"strings"

	"employee-management/auth"
	"employee-management/middleware"
)

// NewHTTPHandler exposes the GraphQL endpoint.
//
// POST requests require a valid JWT.
// GET requests serve the GraphiQL interface.
// GraphiQL must send the Authorization header with its GraphQL requests.
func NewHTTPHandler(server *Server, authService *auth.AuthService) http.Handler {
	graphQLHandler := newGraphQLHTTPHandler(server)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Serve the GraphiQL interface.
			// This does not authorize GraphQL data access.
			graphQLHandler.ServeHTTP(w, r)

		case http.MethodPost:
			// Reject missing or malformed bearer tokens.
			header := strings.TrimSpace(r.Header.Get("Authorization"))
			if header == "" {
				writeGraphQLError(w, http.StatusUnauthorized, "missing bearer token")
				return
			}

			// RequireAuth validates the token and adds the
			// authenticated username and role to the context.
			middleware.RequireAuth(authService, graphQLHandler).
				ServeHTTP(w, r)

		default:
			w.Header().Set("Allow", "GET, POST")
			writeGraphQLError(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	})
}

// writeGraphQLError writes a JSON error response.
func writeGraphQLError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(`{"error":"` + message + `"}`))
}
