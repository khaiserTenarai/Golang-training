package repository

import "assignment_task13/model"

type UserRepository interface {
	Register(user model.User) error

	Login(username string) (model.User, error)
}
