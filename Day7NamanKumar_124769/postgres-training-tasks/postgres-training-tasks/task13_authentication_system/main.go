package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	connString := "postgres://postgres:password@localhost:5432/authdb"
	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		fmt.Println("Connection failed:", err)
		os.Exit(1)
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)
	ctx := context.Background()

	for {
		fmt.Println("\n1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		switch readLine(reader) {
		case "1":
			register(ctx, db, reader)
		case "2":
			login(ctx, db, reader)
		case "3":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func register(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Username: ")
	username := readLine(reader)

	fmt.Print("Password: ")
	password := readLine(reader)

	fmt.Print("Role (user/admin): ")
	role := readLine(reader)
	if role != "admin" {
		role = "user"
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = db.Exec(ctx, `INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3)`, username, string(hash), role)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Registered successfully")
}

func login(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Username: ")
	username := readLine(reader)

	fmt.Print("Password: ")
	password := readLine(reader)

	var storedHash, role string
	err := db.QueryRow(ctx, `SELECT password_hash, role FROM users WHERE username = $1`, username).Scan(&storedHash, &role)
	if errors.Is(err, pgx.ErrNoRows) {
		fmt.Println("Invalid username or password")
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password)); err != nil {
		fmt.Println("Invalid username or password")
		return
	}

	fmt.Printf("Login successful. Role: %s\n", role)
}
