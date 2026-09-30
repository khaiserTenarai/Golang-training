package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Department struct {
	ID   int
	Name string
}

type Employee struct {
	ID           int
	Name         string
	Email        string
	DepartmentID *int
}

func main() {
	connString := "postgres://postgres:password@localhost:5432/deptdb"
	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		fmt.Println("Connection failed:", err)
		os.Exit(1)
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)
	ctx := context.Background()

	for {
		fmt.Println("\n1. Create Department")
		fmt.Println("2. List Departments")
		fmt.Println("3. Update Department")
		fmt.Println("4. Delete Department")
		fmt.Println("5. Create Employee")
		fmt.Println("6. Assign Employee to Department")
		fmt.Println("7. List Employees by Department")
		fmt.Println("8. Exit")
		fmt.Print("Choose an option: ")

		switch readLine(reader) {
		case "1":
			createDepartment(ctx, db, reader)
		case "2":
			listDepartments(ctx, db)
		case "3":
			updateDepartment(ctx, db, reader)
		case "4":
			deleteDepartment(ctx, db, reader)
		case "5":
			createEmployee(ctx, db, reader)
		case "6":
			assignEmployee(ctx, db, reader)
		case "7":
			listEmployeesByDepartment(ctx, db, reader)
		case "8":
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

func createDepartment(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Department name: ")
	name := readLine(reader)

	var id int
	err := db.QueryRow(ctx, `INSERT INTO departments (name) VALUES ($1) RETURNING id`, name).Scan(&id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Created department with ID:", id)
}

func listDepartments(ctx context.Context, db *pgxpool.Pool) {
	rows, err := db.Query(ctx, `SELECT id, name FROM departments ORDER BY id`)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var d Department
		if err := rows.Scan(&d.ID, &d.Name); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("ID: %d | Name: %s\n", d.ID, d.Name)
	}
}

func updateDepartment(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Department ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	fmt.Print("New name: ")
	name := readLine(reader)

	result, err := db.Exec(ctx, `UPDATE departments SET name = $1 WHERE id = $2`, name, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if result.RowsAffected() == 0 {
		fmt.Println("Department not found")
		return
	}
	fmt.Println("Department updated")
}

func deleteDepartment(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Department ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	result, err := db.Exec(ctx, `DELETE FROM departments WHERE id = $1`, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if result.RowsAffected() == 0 {
		fmt.Println("Department not found")
		return
	}
	fmt.Println("Department deleted")
}

func createEmployee(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Name: ")
	name := readLine(reader)

	fmt.Print("Email: ")
	email := readLine(reader)

	var id int
	err := db.QueryRow(ctx, `INSERT INTO employees (name, email) VALUES ($1, $2) RETURNING id`, name, email).Scan(&id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Created employee with ID:", id)
}

func assignEmployee(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Employee ID: ")
	empID, _ := strconv.Atoi(readLine(reader))

	fmt.Print("Department ID: ")
	deptID, _ := strconv.Atoi(readLine(reader))

	result, err := db.Exec(ctx, `UPDATE employees SET department_id = $1 WHERE id = $2`, deptID, empID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if result.RowsAffected() == 0 {
		fmt.Println("Employee not found")
		return
	}
	fmt.Println("Employee assigned to department")
}

func listEmployeesByDepartment(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Department ID: ")
	deptID, _ := strconv.Atoi(readLine(reader))

	query := `SELECT e.id, e.name, e.email FROM employees e WHERE e.department_id = $1 ORDER BY e.id`
	rows, err := db.Query(ctx, query, deptID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		var id int
		var name, email string
		if err := rows.Scan(&id, &name, &email); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("ID: %d | Name: %s | Email: %s\n", id, name, email)
		found = true
	}
	if !found {
		fmt.Println("No employees in this department")
	}
}
