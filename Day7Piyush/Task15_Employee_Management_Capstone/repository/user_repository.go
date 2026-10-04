package repository

import (
	"database/sql"
	"task15_employee_management_capstone/models"
	"time"
)

// UserRepository interface
type UserRepository interface {
	Create(user models.User) (models.User, error)
	GetByUsername(username string) (models.User, error)
	GetByID(id int) (models.User, error)
	CreateSession(userID int, token string, expiresAt time.Time) error
	GetSession(token string) (int, string, error)
	DeleteSession(token string) error
}

// PostgresUserRepository implements UserRepository
type PostgresUserRepository struct {
	DB *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) UserRepository {
	return &PostgresUserRepository{DB: db}
}

func (r *PostgresUserRepository) Create(user models.User) (models.User, error) {
	err := r.DB.QueryRow(
		`INSERT INTO cap_users (username, email, password_hash, role) VALUES ($1, $2, $3, $4)
		RETURNING id, username, email, role, created_at`,
		user.Username, user.Email, user.PasswordHash, user.Role).
		Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.CreatedAt)
	user.PasswordHash = ""
	return user, err
}

func (r *PostgresUserRepository) GetByUsername(username string) (models.User, error) {
	var user models.User
	err := r.DB.QueryRow(
		`SELECT id, username, email, password_hash, role, created_at FROM cap_users WHERE username=$1`,
		username).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role, &user.CreatedAt)
	return user, err
}

func (r *PostgresUserRepository) GetByID(id int) (models.User, error) {
	var user models.User
	err := r.DB.QueryRow(
		`SELECT id, username, email, role, created_at FROM cap_users WHERE id=$1`, id).
		Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.CreatedAt)
	return user, err
}

func (r *PostgresUserRepository) CreateSession(userID int, token string, expiresAt time.Time) error {
	_, err := r.DB.Exec(
		`INSERT INTO cap_sessions (user_id, token, expires_at) VALUES ($1, $2, $3)`,
		userID, token, expiresAt)
	return err
}

func (r *PostgresUserRepository) GetSession(token string) (int, string, error) {
	var userID int
	var expiresAt time.Time
	err := r.DB.QueryRow(
		`SELECT user_id, expires_at FROM cap_sessions WHERE token=$1`, token).
		Scan(&userID, &expiresAt)
	if err != nil {
		return 0, "", err
	}
	if time.Now().After(expiresAt) {
		r.DB.Exec(`DELETE FROM cap_sessions WHERE token=$1`, token)
		return 0, "", sql.ErrNoRows
	}
	var role string
	r.DB.QueryRow(`SELECT role FROM cap_users WHERE id=$1`, userID).Scan(&role)
	return userID, role, nil
}

func (r *PostgresUserRepository) DeleteSession(token string) error {
	_, err := r.DB.Exec(`DELETE FROM cap_sessions WHERE token=$1`, token)
	return err
}
