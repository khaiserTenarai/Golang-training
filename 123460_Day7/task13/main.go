package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time" // Imported and used correctly below

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

var reader = bufio.NewReader(os.Stdin)

func readInput(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	connString := "postgres://postgres:pgadmin@localhost:5432/employee_db"

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("PostgreSQL connected successfully")

	for {
		fmt.Println("\n==============================")
		fmt.Println("    AUTHENTICATION SYSTEM")
		fmt.Println("==============================")
		fmt.Println("1. Register User")
		fmt.Println("2. Login")
		fmt.Println("3. List All Users (Admin View)")
		fmt.Println("4. Exit")

		choiceStr := readInput("Choose an option: ")
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 4.")
			continue
		}

		switch choice {
		case 1:
			registerUser(conn)
		case 2:
			loginUser(conn)
		case 3:
			listUsers(conn)
		case 4:
			fmt.Println("Exiting application. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please select between 1 and 4.")
		}
	}
}

// HashPassword generates a bcrypt hash from a plain text password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a plain text password with a hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func registerUser(conn *pgx.Conn) {
	fmt.Println("\n----- USER REGISTRATION -----")
	username := readInput("Enter Username: ")
	password := readInput("Enter Password: ")
	
	fmt.Println("Roles available: admin, user")
	role := readInput("Enter Role (default: user): ")
	if role == "" {
		role = "user"
	}

	// 1. Hash the password
	hashedPassword, err := HashPassword(password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}

	// 2. Store the username, hash, and role in the database
	var id int
	err = conn.QueryRow(
		context.Background(),
		`INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3) RETURNING id`,
		username, hashedPassword, role,
	).Scan(&id)

	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			fmt.Println("Error: Username already exists.")
		} else {
			fmt.Println("Error registering user:", err)
		}
	} else {
		fmt.Printf("User '%s' registered successfully with ID: %d and Role: '%s'\n", username, id, role)
	}
}

func loginUser(conn *pgx.Conn) {
	fmt.Println("\n----- USER LOGIN -----")
	username := readInput("Username: ")
	password := readInput("Password: ")

	// 1. Fetch the stored hash and role from the database
	var storedHash, role string
	err := conn.QueryRow(
		context.Background(),
		`SELECT password_hash, role FROM users WHERE username = $1`,
		username,
	).Scan(&storedHash, &role)

	if err != nil {
		fmt.Println("Login Failed: Invalid username or password.")
		return
	}

	// 2. Compare the provided password with the stored hash
	if CheckPasswordHash(password, storedHash) {
		fmt.Println("\nLogin Successful!")
		fmt.Printf("Welcome, %s! Your access level is: %s\n", username, strings.ToUpper(role))
	} else {
		fmt.Println("\n❌ Login Failed: Invalid username or password.")
	}
}

func listUsers(conn *pgx.Conn) {
	fmt.Println("\n----- SYSTEM USERS (ADMIN VIEW) -----")
	
	rows, err := conn.Query(
		context.Background(),
		`SELECT id, username, role, created_at FROM users ORDER BY id`,
	)
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return
	}
	defer rows.Close()

	fmt.Printf("%-5s | %-15s | %-10s | %s\n", "ID", "Username", "Role", "Created At")
	fmt.Println("---------------------------------------------------------------")
	
	count := 0
	for rows.Next() {
		var id int
		var username, role string
		var createdAt time.Time 
		
		if err := rows.Scan(&id, &username, &role, &createdAt); err != nil {
			fmt.Println("Error reading row:", err) 
			continue
		}
		count++
		
		fmt.Printf("%-5d | %-15s | %-10s | %s\n", id, username, role, createdAt.Format("2006-01-02 15:04:05"))
	}

	if count == 0 {
		fmt.Println("No users found. Try registering one first!")
	}
}