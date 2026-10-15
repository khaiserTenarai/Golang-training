package repository

import (
	"context"
	"database/sql"
	"errors"

	"secure-banking-api/model"

	"github.com/lib/pq"
)

type UserRepository interface {
	Create(ctx context.Context, u *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepository struct{ db *sql.DB }

func NewUserRepository(db *sql.DB) UserRepository { return &userRepository{db: db} }

func (r *userRepository) Create(ctx context.Context, u *model.User) error {
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO users (name, email, password_hash, role)
		 VALUES ($1, $2, $3, $4) RETURNING id, created_at`,
		u.Name, u.Email, u.PasswordHash, u.Role,
	).Scan(&u.ID, &u.CreatedAt)

	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "23505" {
		return model.ErrEmailExists
	}
	return err
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var u model.User
	err := r.db.QueryRowContext(ctx,
		`SELECT id, name, email, password_hash, role, created_at FROM users WHERE email = $1`, email,
	).Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.Role, &u.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, model.ErrNotFound
	}
	return &u, err
}
