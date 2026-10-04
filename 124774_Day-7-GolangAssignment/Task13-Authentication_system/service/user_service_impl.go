package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"authentication/model"
	"authentication/repository"
	"authentication/utility"
)

type UserServiceImpl struct {
	repository repository.UserRepository
}

func NewUserService(
	repository repository.UserRepository,
) UserService {

	return &UserServiceImpl{
		repository: repository,
	}
}

func hashPassword(password string) string {

	hash := sha256.Sum256([]byte(password))

	return hex.EncodeToString(hash[:])
}

func (s *UserServiceImpl) Register(
	username string,
	password string,
	role string,
) error {

	// Validation

	if err := utility.ValidateUsername(username); err != nil {
		return err
	}

	if err := utility.ValidatePassword(password); err != nil {
		return err
	}

	if err := utility.ValidateRole(role); err != nil {
		return err
	}

	// Hash password

	hashedPassword := hashPassword(password)

	user := model.User{
		Username: username,
		Password: hashedPassword,
		Role:     role,
	}

	return s.repository.Register(user)
}

func (s *UserServiceImpl) Login(
	username string,
	password string,
) error {

	user, err := s.repository.FindByUsername(username)

	if err != nil {
		return err
	}

	// Hash entered password

	hashedPassword := hashPassword(password)

	// Compare passwords

	if hashedPassword != user.Password {
		return errors.New("invalid username or password")
	}

	return nil
}
