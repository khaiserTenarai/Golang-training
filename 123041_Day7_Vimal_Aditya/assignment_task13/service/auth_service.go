package service

import "assignment_task13/model"

type AuthService interface {
	Register(user model.User) error

	Login(username, password string) error
}
