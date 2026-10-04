package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Models
type Department struct {
	ID   int
	Name string
}

type Employee struct {
	ID           int
	Name         string
	Position     string
	Salary       float64
	DepartmentID int // Foreign Key
}

func main() {
	connStr := "host=localhost port=5432 user=postgres password=pgadmin dbname=gotraining sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// --- DEPARTMENT CRUD ---
	
	// 1. Create Department
	deptID := createDepartment(db, "Engineering")
	fmt.Printf("Created Department with ID: %d\n", deptID)

	// 2. Read Department
	dept := getDepartment(db, deptID)
	fmt.Printf("Read Department: %+v\n", dept)

	// 3. Update Department
	updateDepartment(db, deptID, "Core Engineering")
	fmt.Println("Updated Department successfully.")

	// --- EMPLOYEE CRUD (Mapped to Department) ---

	// 1. Create Employee (Linked to deptID)
	empID := createEmployee(db, "David Chen", "Backend Developer", 95000.00, deptID)
	fmt.Printf("Created Employee mapped to Dept %d, ID: %d\n", deptID, empID)

	// 2. Read Employee
	emp := getEmployee(db, empID)
	fmt.Printf("Read Employee: %+v\n", emp)

	// --- CLEANUP ---
	// Deleting the department will also delete the employee due to ON DELETE CASCADE
	deleteDepartment(db, deptID)
	fmt.Println("Deleted Department (and cascaded to Employees) successfully.")
}

// ==============================
// DEPARTMENT OPERATIONS
// ==============================

func createDepartment(db *sql.DB, name string) int {
	var id int
	err := db.QueryRow(`INSERT INTO departments (name) VALUES ($1) RETURNING id`, name).Scan(&id)
	if err != nil {
		log.Fatalf("Error creating department: %v", err)
	}
	return id
}

func getDepartment(db *sql.DB, id int) Department {
	var dept Department
	err := db.QueryRow(`SELECT id, name FROM departments WHERE id = $1`, id).Scan(&dept.ID, &dept.Name)
	if err != nil {
		log.Fatalf("Error reading department: %v", err)
	}
	return dept
}

func updateDepartment(db *sql.DB, id int, newName string) {
	_, err := db.Exec(`UPDATE departments SET name = $1 WHERE id = $2`, newName, id)
	if err != nil {
		log.Fatalf("Error updating department: %v", err)
	}
}

func deleteDepartment(db *sql.DB, id int) {
	_, err := db.Exec(`DELETE FROM departments WHERE id = $1`, id)
	if err != nil {
		log.Fatalf("Error deleting department: %v", err)
	}
}

// ==============================
// EMPLOYEE OPERATIONS
// ==============================

func createEmployee(db *sql.DB, name, position string, salary float64, deptID int) int {
	var id int
	query := `INSERT INTO employees (name, position, salary, department_id) VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(query, name, position, salary, deptID).Scan(&id)
	if err != nil {
		log.Fatalf("Error creating employee: %v", err)
	}
	return id
}

func getEmployee(db *sql.DB, id int) Employee {
	var emp Employee
	query := `SELECT id, name, position, salary, department_id FROM employees WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary, &emp.DepartmentID)
	if err != nil {
		log.Fatalf("Error reading employee: %v", err)
	}
	return emp
}