package graphql

import (
	"github.com/graphql-go/handler"
	"net/http"
)

func NewHTTPHandler(s *Server) http.Handler {
	return handler.New(&handler.Config{Schema: &s.schema, Pretty: true, GraphiQL: true, Playground: true})
}
