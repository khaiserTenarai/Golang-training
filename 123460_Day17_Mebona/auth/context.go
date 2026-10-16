package auth

import "context"

type ctxKey struct{}

// WithUser stores the logged-in user inside the request context.
func WithUser(ctx context.Context, c *Claims) context.Context {
	return context.WithValue(ctx, ctxKey{}, c)
}

// UserFrom reads the logged-in user back. It returns nil when nobody is logged in.
func UserFrom(ctx context.Context) *Claims {
	if ctx == nil {
		return nil
	}
	c, _ := ctx.Value(ctxKey{}).(*Claims)
	return c
}

// RequireUser = authentication check ("are you logged in?")
func RequireUser(ctx context.Context) (*Claims, error) {
	c := UserFrom(ctx)
	if c == nil {
		return nil, ErrUnauthenticated
	}
	return c, nil
}

// RequireRole = authorization check ("are you allowed to do this?")
func RequireRole(ctx context.Context, role string) (*Claims, error) {
	c, err := RequireUser(ctx)
	if err != nil {
		return nil, err
	}
	if c.Role != role {
		return nil, ErrForbidden
	}
	return c, nil
}
