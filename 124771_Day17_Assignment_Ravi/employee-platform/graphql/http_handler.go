package graphql

import (
	"context"
	"net/http"
	"strings"

	"employee-platform/utils"
	gqlhandler "github.com/graphql-go/handler"
)

// NewHTTPHandler follows the same pattern as the user's previous project.
// GraphiQL and Playground are enabled by graphql-go/handler.
func NewHTTPHandler(s *Server) http.Handler {
	return gqlhandler.New(&gqlhandler.Config{
		Schema:     &s.schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})
}

// WithAuthContext attaches verified JWT claims to the request context without
// blocking the GraphiQL page itself. Resolvers require claims for protected operations.
func WithAuthContext(secret string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parts := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(parts) == 2 && strings.EqualFold(parts[0], "Bearer") {
			if claims, err := utils.ParseToken(parts[1], secret); err == nil {
				r = r.WithContext(context.WithValue(r.Context(), claimsContextKey{}, claims))
			}
		}
		next.ServeHTTP(w, r)
	})
}
