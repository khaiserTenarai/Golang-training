package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const connStr = "postgres://postgres:Indu%40123@localhost:5432/assignment_db"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func searchEmployees(ctx *context.Context, conn *pgx.Conn, nameFilter string, minSalary float64, sortBy string, limit, offset int) ([]Employee, error) {
	query := fmt.Sprintf(`
		SELECT id, name, salary 
		FROM employees 
		WHERE name ILIKE $1 AND salary >= $2
		ORDER BY %s
		LIMIT $3 OFFSET $4`, sortBy)

	rows, err := conn.Query(*ctx, query, "%"+nameFilter+"%", minSalary, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Employee
	for rows.Next() {
		var e Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Salary); err != nil {
			return nil, err
		}
		result = append(result, e)
	}
	return result, nil
}

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Connection failed: %v\n", err)
	}
	defer conn.Close(ctx)

	// Prepare data
	_, _ = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS employees (
			id SERIAL PRIMARY KEY, name VARCHAR(100), email VARCHAR(100) UNIQUE, salary NUMERIC(10,2)
		);
		INSERT INTO employees (name, email, salary) VALUES 
		('Charlie', 'c1@test.com', 50000), ('Charlene', 'c2@test.com', 75000), ('Charlotte', 'c3@test.com', 90000)
		ON CONFLICT (email) DO NOTHING;
	`)

	// Search parameters: page 1, 2 items per page, sorted by salary DESC
	page, pageSize := 1, 2
	offset := (page - 1) * pageSize

	emps, err := searchEmployees(&ctx, conn, "Char", 40000, "salary DESC", pageSize, offset)
	if err != nil {
		log.Fatalf("Search failed: %v\n", err)
	}

	fmt.Printf("--- Page %d Results ---\n", page)
	for _, e := range emps {
		fmt.Printf("ID: %d | Name: %s | Salary: %.2f\n", e.ID, e.Name, e.Salary)
	}
}