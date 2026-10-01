package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

const connStr = "postgres://postgres:Info%402k26@localhost:5432/assignment_db"

func registerUser(ctx context.Context, conn *pgx.Conn, username, password, role string) error {
	// Hash password securely with Bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, "INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3)", username, string(hashedPassword), role)
	return err
}

func loginUser(ctx context.Context, conn *pgx.Conn, username, password string) (string, error) {
	var hash, role string
	err := conn.QueryRow(ctx, "SELECT password_hash, role FROM users WHERE username=$1", username).Scan(&hash, &role)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	// Compare raw password with stored Bcrypt hash
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid password")
	}

	return role, nil
}

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Connection failed: %v\n", err)
	}
	defer conn.Close(ctx)

	_, _ = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			role VARCHAR(20) NOT NULL
		);
	`)

	// 1. Register User
	err = registerUser(ctx, conn, "admin_user", "SecurePass123!", "ADMIN")
	if err != nil {
		fmt.Printf("Registration Note: %v\n", err)
	} else {
		fmt.Println("User 'admin_user' registered successfully with hashed password!")
	}

	// 2. Test Login (Success)
	role, err := loginUser(ctx, conn, "admin_user", "SecurePass123!")
	if err != nil {
		fmt.Printf("Login failed: %v\n", err)
	} else {
		fmt.Printf("Login successful! Welcome user with Role: %s\n", role)
	}

	// 3. Test Login (Wrong Password)
	_, err = loginUser(ctx, conn, "admin_user", "WrongPassword")
	if err != nil {
		fmt.Printf("Expected Failure: %v\n", err)
	}
}