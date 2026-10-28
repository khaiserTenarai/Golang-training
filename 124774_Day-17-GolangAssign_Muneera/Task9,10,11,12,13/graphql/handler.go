package graphql

import (
	"github.com/graphql-go/handler"
)

func (s *Server) NewHTTPHandler() *handler.Handler {
	return handler.New(&handler.Config{
		Schema:   &s.schema,
		Pretty:   true,
		GraphiQL: true,
	})
}
