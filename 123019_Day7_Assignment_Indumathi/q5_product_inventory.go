package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const connStr = "postgres://postgres:Indu%40123@localhost:5432/assignment_db"

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Connection failed: %v\n", err)
	}
	defer conn.Close(ctx)

	schema := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		stock INT NOT NULL DEFAULT 0,
		price NUMERIC(10,2) NOT NULL
	);`
	_, _ = conn.Exec(ctx, schema)

	// Add product
	var prodID int
	_ = conn.QueryRow(ctx, "INSERT INTO products (name, stock, price) VALUES ($1, $2, $3) RETURNING id", "Laptop", 15, 999.99).Scan(&prodID)

	// Increase stock
	_, _ = conn.Exec(ctx, "UPDATE products SET stock = stock + $1 WHERE id = $2", 5, prodID)

	// Decrease stock
	_, _ = conn.Exec(ctx, "UPDATE products SET stock = stock - $1 WHERE id = $2", 12, prodID)

	// Search Low Stock Products (Stock < 10)
	rows, err := conn.Query(ctx, "SELECT id, name, stock FROM products WHERE stock < $1", 10)
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}
	defer rows.Close()

	fmt.Println("--- Low Stock Alert (< 10 units) ---")
	for rows.Next() {
		var id, stock int
		var name string
		_ = rows.Scan(&id, &name, &stock)
		fmt.Printf("Product ID: %d | Name: %s | Remaining Stock: %d\n", id, name, stock)
	}
}