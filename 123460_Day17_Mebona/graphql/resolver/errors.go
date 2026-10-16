package resolver

import (
	"employee-management/auth"
	"employee-management/service"
	"errors"
	"log"
)

// apiError is a GraphQL error that also carries a machine-readable code
// (graphql-go puts it into  errors[].extensions.code  automatically).
type apiError struct {
	message string
	code    string
}

func (e *apiError) Error() string { return e.message }
func (e *apiError) Extensions() map[string]interface{} {
	return map[string]interface{}{"code": e.code}
}
func newError(code, msg string) error { return &apiError{msg, code} }

// toGraphQLError converts service / database errors into safe GraphQL errors.
func toGraphQLError(err error, entity string) error {
	var v *service.ValidationError
	switch {
	case err == nil:
		return nil
	case errors.Is(err, auth.ErrUnauthenticated), errors.Is(err, service.ErrInvalidCredentials):
		return newError("UNAUTHENTICATED", err.Error())
	case errors.Is(err, auth.ErrForbidden):
		return newError("FORBIDDEN", err.Error())
	case service.IsNotFound(err):
		return newError("NOT_FOUND", entity+" not found")
	case errors.As(err, &v):
		return newError("BAD_USER_INPUT", v.Message)
	case service.IsDuplicate(err):
		return newError("CONFLICT", err.Error())
	default:
		log.Println("graphql internal error:", err) // real cause stays in the server log
		return newError("INTERNAL_SERVER_ERROR", "something went wrong")
	}
}
