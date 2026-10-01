package service

import "context"

type AuthService interface {
	Register(context.Context, string, string, string) error
	Login(context.Context, string, string) (string, error)
}
