package repository

import (
	"context"

	"assignment_task13/model"

	"github.com/jackc/pgx/v5"
)

type UserRepositoryImpl struct {
	conn *pgx.Conn
}

func NewUserRepository(
	conn *pgx.Conn,
) UserRepository {

	return &UserRepositoryImpl{

		conn: conn,
	}

}

func (repo *UserRepositoryImpl) Register(
	user model.User,
) error {

	_, err := repo.conn.Exec(

		context.Background(),

		`
		INSERT INTO users
		(username,password,role)

		VALUES($1,$2,$3)
		`,

		user.Username,

		user.Password,

		user.Role,
	)

	return err

}

func (repo *UserRepositoryImpl) Login(
	username string,
) (model.User, error) {

	var user model.User

	err := repo.conn.QueryRow(

		context.Background(),

		`
		SELECT id,username,password,role

		FROM users

		WHERE username=$1
		`,

		username,
	).Scan(

		&user.ID,

		&user.Username,

		&user.Password,

		&user.Role,
	)

	return user, err

}
