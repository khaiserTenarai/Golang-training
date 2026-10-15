// Package apperror defines errors that carry a machine-readable code.
// graphql-go copies Extensions() into the "extensions" field of the GraphQL error,
// so clients can react to `code` instead of parsing messages.
package apperror

import (
	"errors"
	"log"
)

type Code string

const (
	CodeNotFound Code = "NOT_FOUND"
	CodeBadInput Code = "BAD_USER_INPUT"
	CodeConflict Code = "CONFLICT"
	CodeInternal Code = "INTERNAL"
)

type Error struct {
	Code    Code
	Message string
	Field   string // optional: which input field caused the problem
}

func (e *Error) Error() string { return e.Message }

// Extensions implements gqlerrors.ExtendedError.
func (e *Error) Extensions() map[string]interface{} {
	m := map[string]interface{}{"code": string(e.Code)}
	if e.Field != "" {
		m["field"] = e.Field
	}
	return m
}

func NotFound(msg string) *Error        { return &Error{Code: CodeNotFound, Message: msg} }
func Conflict(field, msg string) *Error { return &Error{Code: CodeConflict, Message: msg, Field: field} }
func BadInput(field, msg string) *Error { return &Error{Code: CodeBadInput, Message: msg, Field: field} }

// Wrap keeps our own errors as they are. Anything else (for example a database
// failure) is logged and replaced by a generic message, so internals never leak to clients.
func Wrap(err error) error {
	var ae *Error
	if errors.As(err, &ae) {
		return err
	}
	log.Printf("internal error: %v", err)
	return &Error{Code: CodeInternal, Message: "internal server error"}
}

// As returns the *Error inside err (after Wrap), for REST status mapping.
func As(err error) *Error {
	var ae *Error
	errors.As(Wrap(err), &ae)
	return ae
}
