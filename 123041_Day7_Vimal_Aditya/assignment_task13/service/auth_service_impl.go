package service

import (
	"errors"

	"assignment_task13/model"

	"assignment_task13/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	repo repository.UserRepository
}

func NewAuthService(
	repo repository.UserRepository,
) AuthService {

	return &AuthServiceImpl{

		repo: repo,
	}

}

func (s *AuthServiceImpl) Register(
	user model.User,
) error {

	hash, err := bcrypt.GenerateFromPassword(

		[]byte(user.Password),

		bcrypt.DefaultCost,
	)

	if err != nil {

		return err

	}

	user.Password = string(hash)

	return s.repo.Register(user)

}

func (s *AuthServiceImpl) Login(
	username, password string,
) error {

	user, err := s.repo.Login(username)

	if err != nil {

		return err

	}

	err = bcrypt.CompareHashAndPassword(

		[]byte(user.Password),

		[]byte(password),
	)

	if err != nil {

		return errors.New("invalid password")

	}

	return nil

}
