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

	_, _ = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS customers (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			phone VARCHAR(20) UNIQUE NOT NULL
		);
		INSERT INTO customers (name, phone) VALUES ('Eve', '1112223333'), ('Edward', '4445556666'), ('Ella', '7778889999')
		ON CONFLICT (phone) DO NOTHING;
	`)

	// Search and Paginate
	searchTerm := "E%"
	limit, offset := 2, 0

	rows, err := conn.Query(ctx, "SELECT id, name, phone FROM customers WHERE name LIKE $1 LIMIT $2 OFFSET $3", searchTerm, limit, offset)
	if err != nil {
		log.Fatalf("Search failed: %v\n", err)
	}
	defer rows.Close()

	fmt.Println("--- Customer Search Results ---")
	for rows.Next() {
		var id int
		var name, phone string
		_ = rows.Scan(&id, &name, &phone)
		fmt.Printf("Customer ID: %d | Name: %s | Phone: %s\n", id, name, phone)
	}
}