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
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)

	// 1. Create Table
	schema := `
	CREATE TABLE IF NOT EXISTS employees (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		salary NUMERIC(10,2) NOT NULL
	);`
	_, err = conn.Exec(ctx, schema)
	if err != nil {
		log.Fatalf("Failed to create table: %v\n", err)
	}

	// 2. CREATE (Insert)
	var id int
	err = conn.QueryRow(ctx, "INSERT INTO employees (name, email, salary) VALUES ($1, $2, $3) RETURNING id", "Sasi", "sasi@example.com", 65000.00).Scan(&id)
	if err != nil {
		log.Fatalf("Insert failed: %v\n", err)
	}
	fmt.Printf("[CREATE] Inserted Employee ID: %d\n", id)

	// 3. READ (Select)
	var name, email string
	var salary float64
	err = conn.QueryRow(ctx, "SELECT name, email, salary FROM employees WHERE id=$1", id).Scan(&name, &email, &salary)
	if err != nil {
		log.Fatalf("Select failed: %v\n", err)
	}
	fmt.Printf("[READ] ID: %d, Name: %s, Email: %s, Salary: %.2f\n", id, name, email, salary)

	// 4. UPDATE
	_, err = conn.Exec(ctx, "UPDATE employees SET salary=$1 WHERE id=$2", 72000.00, id)
	if err != nil {
		log.Fatalf("Update failed: %v\n", err)
	}
	fmt.Printf("[UPDATE] Updated Salary for Employee ID: %d\n", id)

	// 5. DELETE
	_, err = conn.Exec(ctx, "DELETE FROM employees WHERE id=$1", id)
	if err != nil {
		log.Fatalf("Delete failed: %v\n", err)
	}
	fmt.Printf("[DELETE] Deleted Employee ID: %d\n", id)
}