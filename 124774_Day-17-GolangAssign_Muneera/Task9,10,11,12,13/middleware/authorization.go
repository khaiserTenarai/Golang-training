package middleware

import (
	"context"
	"errors"
)

func RequireAdmin(ctx context.Context) error {
	if RoleFromContext(ctx) != "ADMIN" {
		return errors.New("admin permission required")
	}

	return nil
}
