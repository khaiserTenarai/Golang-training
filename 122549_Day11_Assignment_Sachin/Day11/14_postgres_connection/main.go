package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

func ask(reader *bufio.Reader, label, fallback string) string {
	fmt.Printf("%s [%s]: ", label, fallback)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "" {
		return fallback
	}
	return text
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	host := ask(reader, "Database host", "localhost")
	dbPort := ask(reader, "Database port", "5432")
	user := ask(reader, "Database user", "postgres")
	password := ask(reader, "Database password", "postgres")
	dbName := ask(reader, "Database name", "employee_db")
	serverPort := ask(reader, "Port for this server", "8080")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, dbPort, user, password, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Could not set up the connection:", err)
		return
	}
	defer db.Close()

	// sql.Open doesn't actually connect yet - Ping is what proves it works.
	if err := db.Ping(); err != nil {
		fmt.Println("Could not reach PostgreSQL:", err)
		return
	}
	fmt.Println("Connected to PostgreSQL!")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS employees (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(150) NOT NULL UNIQUE,
		department VARCHAR(100) NOT NULL,
		salary NUMERIC(12,2) NOT NULL
	)`)
	if err != nil {
		fmt.Println("Could not create the table:", err)
		return
	}

	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			rows, err := db.Query(`SELECT id, name, email, department, salary FROM employees ORDER BY id`)
			if err != nil {
				http.Error(w, "database error", http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			employees := []Employee{}
			for rows.Next() {
				var e Employee
				if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary); err != nil {
					http.Error(w, "database error", http.StatusInternalServerError)
					return
				}
				employees = append(employees, e)
			}
			json.NewEncoder(w).Encode(employees)

		case http.MethodPost:
			var e Employee
			if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
				http.Error(w, "could not read the JSON body", http.StatusBadRequest)
				return
			}
			err := db.QueryRow(
				`INSERT INTO employees (name, email, department, salary) VALUES ($1, $2, $3, $4) RETURNING id`,
				e.Name, e.Email, e.Department, e.Salary,
			).Scan(&e.ID)
			if err != nil {
				http.Error(w, "could not save the employee: "+err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(e)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running on port", serverPort)
	if err := http.ListenAndServe(":"+serverPort, nil); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
