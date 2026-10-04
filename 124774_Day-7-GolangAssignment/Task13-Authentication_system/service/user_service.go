package service

type UserService interface {
	Register(
		username string,
		password string,
		role string,
	) error

	Login(
		username string,
		password string,
	) error
}
