package repository

import (
	"context"
	"errors"
	"example.com/q13-authentication/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	Create(context.Context, model.User) error
	FindByUsername(context.Context, string) (model.User, error)
}
type PostgresUserRepository struct{ db *pgxpool.Pool }

func NewPostgresUserRepository(db *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}
func (r *PostgresUserRepository) Create(c context.Context, u model.User) error {
	_, e := r.db.Exec(c, `INSERT INTO users(username,password_hash,role) VALUES($1,$2,$3)`, u.Username, u.PasswordHash, u.Role)
	return e
}
func (r *PostgresUserRepository) FindByUsername(c context.Context, n string) (model.User, error) {
	var u model.User
	e := r.db.QueryRow(c, `SELECT id,username,password_hash,role FROM users WHERE username=$1`, n).Scan(&u.ID, &u.Username, &u.PasswordHash, &u.Role)
	if e != nil {
		return u, ErrUserNotFound
	}
	return u, nil
}
