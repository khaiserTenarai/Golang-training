package repository

import (
	"database/sql"
	"task13_authentication_system/models"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
	err := r.DB.QueryRow(
		`INSERT INTO auth_users (username, email, password_hash, role) VALUES ($1, $2, $3, $4)
		RETURNING id, username, email, role, created_at`,
		user.Username, user.Email, user.PasswordHash, user.Role).
		Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.CreatedAt)
	user.PasswordHash = ""
	return user, err
}

func (r *UserRepository) GetByUsername(username string) (models.User, error) {
	var user models.User
	err := r.DB.QueryRow(
		`SELECT id, username, email, password_hash, role, created_at FROM auth_users WHERE username=$1`,
		username).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role, &user.CreatedAt)
	return user, err
}

func (r *UserRepository) GetByID(id int) (models.User, error) {
	var user models.User
	err := r.DB.QueryRow(
		`SELECT id, username, email, role, created_at FROM auth_users WHERE id=$1`, id).
		Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.CreatedAt)
	return user, err
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	rows, err := r.DB.Query(`SELECT id, username, email, role, created_at FROM auth_users ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Role, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func (r *UserRepository) CreateSession(userID int, token string, expiresAt time.Time) error {
	_, err := r.DB.Exec(
		`INSERT INTO auth_sessions (user_id, token, expires_at) VALUES ($1, $2, $3)`,
		userID, token, expiresAt)
	return err
}

func (r *UserRepository) GetSession(token string) (int, string, error) {
	var userID int
	var expiresAt time.Time
	err := r.DB.QueryRow(
		`SELECT user_id, expires_at FROM auth_sessions WHERE token=$1`, token).
		Scan(&userID, &expiresAt)
	if err != nil {
		return 0, "", err
	}
	if time.Now().After(expiresAt) {
		r.DB.Exec(`DELETE FROM auth_sessions WHERE token=$1`, token)
		return 0, "", sql.ErrNoRows
	}
	var role string
	r.DB.QueryRow(`SELECT role FROM auth_users WHERE id=$1`, userID).Scan(&role)
	return userID, role, nil
}

func (r *UserRepository) DeleteSession(token string) error {
	_, err := r.DB.Exec(`DELETE FROM auth_sessions WHERE token=$1`, token)
	return err
}

func (r *UserRepository) DeleteUser(id int) error {
	_, err := r.DB.Exec(`DELETE FROM auth_users WHERE id=$1`, id)
	return err
}

func (r *UserRepository) UpdateRole(id int, role string) error {
	_, err := r.DB.Exec(`UPDATE auth_users SET role=$1 WHERE id=$2`, role, id)
	return err
}
