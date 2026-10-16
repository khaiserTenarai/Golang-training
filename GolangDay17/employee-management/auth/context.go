package auth

import "context"

type ctxKey struct{}


func WithUser(ctx context.Context, c *Claims) context.Context {
	return context.WithValue(ctx, ctxKey{}, c)
}


func UserFrom(ctx context.Context) *Claims {
	if ctx == nil {
		return nil
	}
	c, _ := ctx.Value(ctxKey{}).(*Claims)
	return c
}


func RequireUser(ctx context.Context) (*Claims, error) {
	c := UserFrom(ctx)
	if c == nil {
		return nil, ErrUnauthenticated
	}
	return c, nil
}


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
