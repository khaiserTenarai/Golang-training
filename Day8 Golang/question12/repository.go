package main

import (
	"database/sql"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	Create(user *User) error
	GetActiveUsers(limit, offset int) ([]User, int, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *User) error {
	query := `
		INSERT INTO users (name, email, is_active, created_at) 
		VALUES ($1, $2, $3, NOW()) 
		RETURNING id, created_at`
	
	return r.db.QueryRow(query, user.Name, user.Email, user.IsActive).Scan(&user.ID, &user.CreatedAt)
}

func (r *userRepository) GetActiveUsers(limit, offset int) ([]User, int, error) {
	query := `
		SELECT id, name, email, is_active, created_at, COUNT(*) OVER() as total_count 
		FROM users 
		WHERE is_active = true 
		ORDER BY created_at DESC 
		LIMIT $1 OFFSET $2`
	
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []User
	var total int
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.IsActive, &u.CreatedAt, &total); err != nil {
			return nil, 0, err
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return users, total, nil
}