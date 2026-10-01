package service

import "day7q13/model"

type AuthService interface {
	Register(user model.User) error

	Login(username, password string) error
}
