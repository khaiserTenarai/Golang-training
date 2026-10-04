package repository

import "authentication/model"

type UserRepository interface {
	Register(user model.User) error

	FindByUsername(username string) (model.User, error)
}
