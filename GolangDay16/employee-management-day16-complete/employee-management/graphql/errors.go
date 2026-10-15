package graphql

import (
	"employee-management/service"
	"errors"
	"log"
)

// apiError is a GraphQL error that carries an "extensions.code" value.
// graphql-go reads the Extensions() method automatically.
type apiError struct {
	message string
	code    string
}

func (e *apiError) Error() string { return e.message }
func (e *apiError) Extensions() map[string]interface{} {
	return map[string]interface{}{"code": e.code}
}
func newError(code, msg string) error { return &apiError{msg, code} }

// toGraphQLError converts service/database errors into safe GraphQL errors.
func toGraphQLError(err error, entity string) error {
	var v *service.ValidationError
	switch {
	case err == nil:
		return nil
	case service.IsNotFound(err):
		return newError("NOT_FOUND", entity+" not found")
	case errors.As(err, &v):
		return newError("BAD_USER_INPUT", v.Message)
	case service.IsDuplicate(err):
		return newError("CONFLICT", err.Error())
	default:
		log.Println("graphql internal error:", err)
		return newError("INTERNAL_SERVER_ERROR", "something went wrong")
	}
}
