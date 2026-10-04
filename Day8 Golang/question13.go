package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Username     string
	PasswordHash string
	Role         string
}

func main() {
	connStr := "host=localhost port=5432 user=postgres password=pgadmin dbname=Day7 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = registerUser(db, "admin_user", "supersecret", "admin")
	if err != nil {
		fmt.Println(err)
	}

	user, err := loginUser(db, "admin_user", "supersecret")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", user)
	}
}

func registerUser(db *sql.DB, username, password, role string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO users (username, password_hash, role) 
		VALUES ($1, $2, $3)`, username, string(hashedPassword), role)
	return err
}

func loginUser(db *sql.DB, username, password string) (*User, error) {
	user := &User{}
	err := db.QueryRow(`
		SELECT id, username, password_hash, role 
		FROM users WHERE username = $1`, username).
		Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Role)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}