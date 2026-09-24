// Package validator provides reusable input validation.
// It lives in pkg/ because other projects may safely import it.
package validator

import (
	"fmt"
	"net/mail"
)

// Email returns an error if addr is not a valid e-mail address.
func Email(addr string) error {
	if _, err := mail.ParseAddress(addr); err != nil {
		return fmt.Errorf("invalid email %q", addr)
	}
	return nil
}
