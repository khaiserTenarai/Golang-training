package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const connStr = "postgres://postgres:sasi2356@localhost:5432/assignment_db"

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Connection failed: %v\n", err)
	}
	defer conn.Close(ctx)

	// Create tables with Foreign Key relationship
	schema := `
	CREATE TABLE IF NOT EXISTS departments (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) UNIQUE NOT NULL
	);
	CREATE TABLE IF NOT EXISTS dept_employees (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		dept_id INT REFERENCES departments(id) ON DELETE SET NULL
	);`
	_, err = conn.Exec(ctx, schema)
	if err != nil {
		log.Fatalf("Schema setup failed: %v\n", err)
	}

	// Insert Department
	var deptID int
	err = conn.QueryRow(ctx, "INSERT INTO departments (name) VALUES ($1) ON CONFLICT (name) DO UPDATE SET name=EXCLUDED.name RETURNING id", "Engineering").Scan(&deptID)
	if err != nil {
		log.Fatalf("Insert dept failed: %v\n", err)
	}
	fmt.Printf("Department Created: ID = %d\n", deptID)

	// Insert Employee mapped to Department
	var empID int
	err = conn.QueryRow(ctx, "INSERT INTO dept_employees (name, dept_id) VALUES ($1, $2) RETURNING id", "Bob", deptID).Scan(&empID)
	if err != nil {
		log.Fatalf("Insert employee failed: %v\n", err)
	}
	fmt.Printf("Employee Created with Dept FK: ID = %d\n", empID)

	// Read joined employee-department data
	var empName, deptName string
	err = conn.QueryRow(ctx, `
		SELECT e.name, d.name 
		FROM dept_employees e 
		JOIN departments d ON e.dept_id = d.id 
		WHERE e.id = $1`, empID).Scan(&empName, &deptName)
	if err != nil {
		log.Fatalf("Query joined data failed: %v\n", err)
	}
	fmt.Printf("Employee %s works in Department %s\n", empName, deptName)
}