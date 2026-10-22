package repositories

import "gin-layered-api/models"

type UserRepository interface {
	FindByUsername(username string) (*models.User, error)
}

type userRepository struct {
	users map[string]models.User
}

func NewUserRepository() UserRepository {
	return &userRepository{
		users: map[string]models.User{
			"alice": {ID: 1, Username: "alice", Password: "password123", Role: models.RoleUser},
			"bob":   {ID: 2, Username: "bob", Password: "adminpassword", Role: models.RoleAdmin},
		},
	}
}

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	user, exists := r.users[username]
	if !exists {
		return nil, nil
	}
	return &user, nil
}
