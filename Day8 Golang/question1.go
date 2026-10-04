package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Employee represents the employee model
type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

func main() {
	// Database connection string
	connStr := "host=localhost port=5432 user=postgres password=pgadmin dbname=gotraining sslmode=disable"
	
	// Open a connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}
	fmt.Println("Successfully connected to database!")

	// 1. CREATE
	newID := createEmployee(db, "Alice Smith", "Software Engineer", 85000.50)
	fmt.Printf("Created employee with ID: %d\n", newID)

	// 2. READ (Single)
	emp := getEmployee(db, newID)
	fmt.Printf("Read: %+v\n", emp)

	// 3. UPDATE
	updateEmployee(db, newID, "Alice Smith", "Senior Software Engineer", 105000.00)
	fmt.Println("Updated employee successfully.")
	
	// Verify Update
	updatedEmp := getEmployee(db, newID)
	fmt.Printf("Read after update: %+v\n", updatedEmp)

	// 4. DELETE
	deleteEmployee(db, newID)
	fmt.Println("Deleted employee successfully.")
}

// createEmployee inserts a new record and returns the generated ID
func createEmployee(db *sql.DB, name, position string, salary float64) int {
	query := `INSERT INTO employees (name, position, salary) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := db.QueryRow(query, name, position, salary).Scan(&id)
	if err != nil {
		log.Fatalf("Error creating employee: %v", err)
	}
	return id
}

// getEmployee fetches a single employee by ID
func getEmployee(db *sql.DB, id int) Employee {
	var emp Employee
	query := `SELECT id, name, position, salary FROM employees WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No employee found with ID %d", id)
		}
		log.Fatalf("Error reading employee: %v", err)
	}
	return emp
}

// updateEmployee modifies an existing employee record
func updateEmployee(db *sql.DB, id int, name, position string, salary float64) {
	query := `UPDATE employees SET name = $1, position = $2, salary = $3 WHERE id = $4`
	_, err := db.Exec(query, name, position, salary, id)
	if err != nil {
		log.Fatalf("Error updating employee: %v", err)
	}
}

// deleteEmployee removes an employee record by ID
func deleteEmployee(db *sql.DB, id int) {
	query := `DELETE FROM employees WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatalf("Error deleting employee: %v", err)
	}
}