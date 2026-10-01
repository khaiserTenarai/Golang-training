package repository

import "day7q13/model"

type UserRepository interface {
	Register(user model.User) error

	Login(username string) (model.User, error)
}
