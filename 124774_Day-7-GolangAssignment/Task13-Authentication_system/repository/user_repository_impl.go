package repository

import (
	"context"
	"errors"

	"authentication/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewUserRepository(
	db *pgxpool.Pool,
) UserRepository {

	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) Register(
	user model.User,
) error {

	_, err := r.db.Exec(
		context.Background(),
		`INSERT INTO users
		(username, password, role)
		VALUES ($1, $2, $3)`,
		user.Username,
		user.Password,
		user.Role,
	)

	return err
}

func (r *UserRepositoryImpl) FindByUsername(
	username string,
) (model.User, error) {

	var user model.User

	err := r.db.QueryRow(
		context.Background(),
		`SELECT id, username, password, role
		FROM users
		WHERE username = $1`,
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return model.User{}, errors.New("user not found")
	}

	return user, err
}
